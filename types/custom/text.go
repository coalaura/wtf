package custom

import (
	"bytes"
	"unicode/utf8"

	"github.com/coalaura/wtf/types"
)

func DetectText(b types.Buffer) *types.Metadata {
	if b.Len() == 0 {
		return nil
	}

	if !passesTextLengthGate(b) {
		return nil
	}

	var (
		isASCII  bool
		isUTF8   bool
		textData []byte
	)

	if b.Has(0, []byte{0xef, 0xbb, 0xbf}) {
		textData = b[3:]

		isUTF8 = isLikelyTextUTF8(textData)
	} else {
		textData = b

		isASCII = isLikelyASCIIText(textData)
		if !isASCII {
			isUTF8 = isLikelyTextUTF8(textData)
		}
	}

	if !isASCII && !isUTF8 {
		return nil
	}

	limit := min(len(textData), 4096)
	subtype := detectTextSubtype(textData[:limit])

	if subtype != types.TypeNone {
		return &types.Metadata{
			Kind:       types.KindText,
			Type:       subtype,
			Confidence: types.ConfidenceMedium,
		}
	}

	fallbackType := types.TypeUTF8
	if isASCII {
		fallbackType = types.TypeASCII
	}

	return &types.Metadata{
		Kind:       types.KindText,
		Type:       fallbackType,
		Confidence: types.ConfidenceLow,
	}
}

func detectTextSubtype(data []byte) types.TypeID {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return types.TypeNone
	}

	// 1. Shebangs (Highly Reliable) - must be at byte 0, not after whitespace
	if bytes.HasPrefix(data, []byte("#!")) {
		lineEnd := bytes.IndexByte(data, '\n')
		if lineEnd == -1 {
			lineEnd = len(data)
		}

		shebang := data[:lineEnd]

		if bytes.Contains(shebang, []byte("bash")) || bytes.Contains(shebang, []byte("sh")) {
			return types.TypeBashScript
		}

		if bytes.Contains(shebang, []byte("python")) {
			return types.TypePythonScript
		}

		if bytes.Contains(shebang, []byte("node")) {
			return types.TypeJavaScriptSource
		}

		if bytes.Contains(shebang, []byte("perl")) {
			return types.TypePerlScript
		}

		if bytes.Contains(shebang, []byte("ruby")) {
			return types.TypeRubyScript
		}

		if bytes.Contains(shebang, []byte("php")) {
			return types.TypePHPScript
		}

		if bytes.Contains(shebang, []byte("lua")) {
			return types.TypeLuaScript
		}

		if bytes.Contains(shebang, []byte("Rscript")) {
			return types.TypeRScript
		}
	}

	// 2. Clear Prefixes
	if bytes.HasPrefix(trimmed, []byte("<?php")) {
		return types.TypePHPScript
	}

	if bytes.HasPrefix(trimmed, []byte("FROM ")) || bytes.HasPrefix(trimmed, []byte("# syntax=")) {
		return types.TypeDockerfile
	}

	if bytes.HasPrefix(trimmed, []byte("# ")) || bytes.HasPrefix(trimmed, []byte("## ")) {
		return types.TypeMarkdownDocument
	}

	if bytes.HasPrefix(trimmed, []byte("---\n")) || bytes.HasPrefix(trimmed, []byte("---\r\n")) {
		if bytes.Contains(trimmed, []byte("title:")) || bytes.Contains(trimmed, []byte("layout:")) {
			return types.TypeMarkdownDocument
		}

		return types.TypeYAMLConfiguration
	}

	if bytes.HasPrefix(trimmed, []byte("CREATE TABLE")) || bytes.HasPrefix(trimmed, []byte("SELECT ")) || bytes.HasPrefix(trimmed, []byte("INSERT INTO")) || bytes.HasPrefix(trimmed, []byte("-- SQL")) {
		return types.TypeSQLScript
	}

	if bytes.HasPrefix(trimmed, []byte("@echo off")) || bytes.HasPrefix(trimmed, []byte("@ECHO OFF")) || bytes.HasPrefix(trimmed, []byte("REM ")) || bytes.HasPrefix(trimmed, []byte("rem ")) {
		return types.TypeBatchScript
	}

	if bytes.HasPrefix(trimmed, []byte("cmake_minimum_required")) || bytes.HasPrefix(trimmed, []byte("project(")) {
		return types.TypeCMakeScript
	}

	if bytes.HasPrefix(trimmed, []byte("body {")) || bytes.HasPrefix(trimmed, []byte("@import url(")) || bytes.HasPrefix(trimmed, []byte("@media ")) || bytes.HasPrefix(trimmed, []byte(":root {")) {
		return types.TypeCSS
	}

	if bytes.HasPrefix(trimmed, []byte("resource \"")) || bytes.HasPrefix(trimmed, []byte("variable \"")) || bytes.HasPrefix(trimmed, []byte("provider \"")) || bytes.HasPrefix(trimmed, []byte("terraform {")) {
		return types.TypeTerraformConfiguration
	}

	if bytes.HasPrefix(trimmed, []byte("query {")) || bytes.HasPrefix(trimmed, []byte("mutation {")) || bytes.HasPrefix(trimmed, []byte("fragment ")) {
		return types.TypeGraphQL
	}

	// 3. Skip C-style comments for code detection (Go, C, C++, Java, Rust, JS, etc.)
	code := skipCStyleComments(trimmed)
	if len(code) == 0 {
		return types.TypeNone
	}

	if bytes.HasPrefix(code, []byte("package ")) {
		if bytes.Contains(code, []byte("import (")) || bytes.Contains(code, []byte("import \"")) || bytes.Contains(code, []byte("func ")) {
			return types.TypeGoSource
		}

		if bytes.Contains(code, []byte("import java")) || bytes.Contains(code, []byte("public class ")) {
			return types.TypeJavaSource
		}

		if bytes.Contains(code, []byte("import ")) || bytes.Contains(code, []byte("fun ")) || bytes.Contains(code, []byte("val ")) {
			return types.TypeKotlinSource
		}

		if bytes.Contains(code, []byte("import scala.")) || bytes.Contains(code, []byte("object ")) || bytes.Contains(code, []byte("trait ")) {
			return types.TypeScalaSource
		}
	}

	if bytes.HasPrefix(code, []byte("#include <")) || bytes.HasPrefix(code, []byte("#include \"")) {
		if bytes.Contains(code, []byte("std::")) || bytes.Contains(code, []byte("class ")) || bytes.Contains(code, []byte("template <")) {
			return types.TypeCPPSource
		}

		return types.TypeCSource
	}

	if bytes.HasPrefix(code, []byte("#import <")) || bytes.HasPrefix(code, []byte("#import \"")) {
		return types.TypeObjectiveCSource
	}

	if bytes.HasPrefix(code, []byte("using System;")) || bytes.HasPrefix(code, []byte("namespace ")) {
		if bytes.Contains(code, []byte("public class ")) || bytes.Contains(code, []byte("public interface ")) {
			return types.TypeCSharpSource
		}
	}

	if bytes.HasPrefix(code, []byte("import Foundation")) || bytes.HasPrefix(code, []byte("import UIKit")) || bytes.HasPrefix(code, []byte("import SwiftUI")) {
		return types.TypeSwiftSource
	}

	if bytes.HasPrefix(code, []byte("fn main()")) || bytes.HasPrefix(code, []byte("use std::")) || bytes.HasPrefix(code, []byte("pub fn ")) || bytes.HasPrefix(code, []byte("#![")) || bytes.HasPrefix(code, []byte("#[derive(")) || bytes.HasPrefix(code, []byte("#[allow(")) {
		return types.TypeRustSource
	}

	if bytes.HasPrefix(code, []byte("import ")) || bytes.HasPrefix(code, []byte("def ")) || bytes.HasPrefix(code, []byte("class ")) {
		if bytes.Contains(code, []byte("def ")) && bytes.Contains(code, []byte(":")) {
			return types.TypePythonScript
		}
	}

	if bytes.HasPrefix(code, []byte("require '")) || bytes.HasPrefix(code, []byte("require \"")) || bytes.HasPrefix(code, []byte("require_relative ")) {
		if bytes.Contains(code, []byte("def ")) || bytes.Contains(code, []byte("class ")) || bytes.Contains(code, []byte("module ")) {
			return types.TypeRubyScript
		}
	}

	if bytes.HasPrefix(code, []byte("local ")) || bytes.HasPrefix(code, []byte("function ")) || bytes.HasPrefix(code, []byte("require ")) || bytes.HasPrefix(code, []byte("require(\"")) {
		if bytes.Contains(code, []byte("end")) || bytes.Contains(code, []byte(" then")) || bytes.Contains(code, []byte(" do")) || bytes.Contains(code, []byte(".lua")) {
			return types.TypeLuaScript
		}
	}

	if bytes.HasPrefix(code, []byte("import 'package:")) || bytes.HasPrefix(code, []byte("import 'dart:")) {
		return types.TypeDartSource
	}

	if bytes.HasPrefix(code, []byte("library(")) || bytes.HasPrefix(code, []byte("require(")) {
		if bytes.Contains(code, []byte("<-")) || bytes.Contains(code, []byte("%>%")) {
			return types.TypeRScript
		}
	}

	if bytes.HasPrefix(code, []byte("const ")) || bytes.HasPrefix(code, []byte("let ")) || bytes.HasPrefix(code, []byte("var ")) || bytes.HasPrefix(code, []byte("import ")) {
		if bytes.Contains(code, []byte("interface ")) || bytes.Contains(code, []byte("type ")) || bytes.Contains(code, []byte(" as ")) || bytes.Contains(code, []byte(": string")) || bytes.Contains(code, []byte(": number")) || bytes.Contains(code, []byte(": boolean")) {
			if bytes.Contains(code, []byte("=>")) || bytes.Contains(code, []byte("console.log")) || bytes.Contains(code, []byte("from '")) || bytes.Contains(code, []byte("from \"")) {
				return types.TypeTypeScriptSource
			}
		}

		if bytes.Contains(code, []byte("=>")) || bytes.Contains(code, []byte("console.log")) || bytes.Contains(code, []byte("function(")) || bytes.Contains(code, []byte("function ")) {
			return types.TypeJavaScriptSource
		}
	}

	if bytes.HasPrefix(code, []byte("[")) {
		lineEnd := bytes.IndexByte(code, '\n')
		if lineEnd != -1 && bytes.Contains(code[:lineEnd], []byte("]")) {
			if bytes.Contains(code, []byte("=\"")) || bytes.Contains(code, []byte("= \"")) || bytes.Contains(code, []byte("['")) {
				return types.TypeTOMLConfiguration
			}

			return types.TypeINIConfiguration
		}
	}

	if bytes.HasPrefix(code, []byte("CFLAGS")) || bytes.HasPrefix(code, []byte("TARGET")) || bytes.HasPrefix(code, []byte("all:")) || bytes.HasPrefix(code, []byte(".PHONY:")) {
		return types.TypeMakefile
	}

	if bytes.Contains(code, []byte("Write-Host ")) || bytes.Contains(code, []byte("Get-")) || bytes.Contains(code, []byte("Set-")) || bytes.Contains(code, []byte("Out-Null")) {
		if bytes.Contains(code, []byte("$")) && (bytes.Contains(code, []byte("-eq")) || bytes.Contains(code, []byte("-ne")) || bytes.Contains(code, []byte("param("))) {
			return types.TypePowerShellScript
		}
	}

	return types.TypeNone
}

func skipCStyleComments(b []byte) []byte {
	for len(b) > 0 {
		b = bytes.TrimSpace(b)
		if bytes.HasPrefix(b, []byte("//")) {
			idx := bytes.IndexByte(b, '\n')
			if idx == -1 {
				return nil
			}

			b = b[idx:]
		} else if bytes.HasPrefix(b, []byte("/*")) {
			idx := bytes.Index(b, []byte("*/"))
			if idx == -1 {
				return nil
			}

			b = b[idx+2:]
		} else {
			break
		}
	}

	return b
}

func passesTextLengthGate(b types.Buffer) bool {
	if b.Len() >= 16 {
		return true
	}

	return bytes.ContainsAny(b, " \t\r\n")
}

func isLikelyASCIIText(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit == 0 {
		return false
	}

	var printable int

	target := b[:limit]

	for _, c := range target {
		if c == 0 {
			return false
		}

		if c == '\n' || c == '\r' || c == '\t' || (c >= 32 && c <= 126) {
			printable++
		}
	}

	return printable*100/limit >= 95
}

func isLikelyTextUTF8(b types.Buffer) bool {
	limit := min(b.Len(), 4096)
	if limit == 0 {
		return false
	}

	target := b[:limit]
	if !utf8.Valid(target) {
		return false
	}

	var printable int

	for _, c := range target {
		if c == 0 {
			return false
		}

		if c == '\n' || c == '\r' || c == '\t' || c >= 32 {
			printable++
		}
	}

	return printable*100/limit >= 90
}
