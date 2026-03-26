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
			Kind:       types.KindTextFile,
			Type:       subtype,
			Confidence: types.ConfidenceMedium,
		}
	}

	fallbackType := types.TypeUTF8Text
	if isASCII {
		fallbackType = types.TypeASCIIText
	}

	return &types.Metadata{
		Kind:       types.KindTextFile,
		Type:       fallbackType,
		Confidence: types.ConfidenceLow,
	}
}

func detectTextSubtype(data []byte) types.TypeID {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return types.TypeNone
	}

	// 1. Shebangs (Highly Reliable)
	if bytes.HasPrefix(trimmed, []byte("#!")) {
		lineEnd := bytes.IndexByte(trimmed, '\n')
		if lineEnd == -1 {
			lineEnd = len(trimmed)
		}

		shebang := trimmed[:lineEnd]

		if bytes.Contains(shebang, []byte("bash")) || bytes.Contains(shebang, []byte("sh")) {
			return types.TypeBashScript
		}

		if bytes.Contains(shebang, []byte("python")) {
			return types.TypePython
		}

		if bytes.Contains(shebang, []byte("node")) {
			return types.TypeJavaScript
		}

		if bytes.Contains(shebang, []byte("perl")) {
			return types.TypePerl
		}

		if bytes.Contains(shebang, []byte("ruby")) {
			return types.TypeRuby
		}

		if bytes.Contains(shebang, []byte("php")) {
			return types.TypePHP
		}

		if bytes.Contains(shebang, []byte("lua")) {
			return types.TypeLua
		}

		if bytes.Contains(shebang, []byte("Rscript")) {
			return types.TypeR
		}
	}

	// 2. Clear Prefixes
	if bytes.HasPrefix(trimmed, []byte("<?php")) {
		return types.TypePHP
	}

	if bytes.HasPrefix(trimmed, []byte("FROM ")) || bytes.HasPrefix(trimmed, []byte("# syntax=")) {
		return types.TypeDocker
	}

	if bytes.HasPrefix(trimmed, []byte("# ")) || bytes.HasPrefix(trimmed, []byte("## ")) {
		return types.TypeMarkdown
	}

	if bytes.HasPrefix(trimmed, []byte("---\n")) || bytes.HasPrefix(trimmed, []byte("---\r\n")) {
		if bytes.Contains(trimmed, []byte("title:")) || bytes.Contains(trimmed, []byte("layout:")) {
			return types.TypeMarkdown
		}

		return types.TypeYAML
	}

	if bytes.HasPrefix(trimmed, []byte("CREATE TABLE")) || bytes.HasPrefix(trimmed, []byte("SELECT ")) || bytes.HasPrefix(trimmed, []byte("INSERT INTO")) || bytes.HasPrefix(trimmed, []byte("-- SQL")) {
		return types.TypeSQL
	}

	if bytes.HasPrefix(trimmed, []byte("@echo off")) || bytes.HasPrefix(trimmed, []byte("@ECHO OFF")) || bytes.HasPrefix(trimmed, []byte("REM ")) || bytes.HasPrefix(trimmed, []byte("rem ")) {
		return types.TypeBatch
	}

	if bytes.HasPrefix(trimmed, []byte("cmake_minimum_required")) || bytes.HasPrefix(trimmed, []byte("project(")) {
		return types.TypeCMake
	}

	if bytes.HasPrefix(trimmed, []byte("body {")) || bytes.HasPrefix(trimmed, []byte("@import url(")) || bytes.HasPrefix(trimmed, []byte("@media ")) || bytes.HasPrefix(trimmed, []byte(":root {")) {
		return types.TypeCSS
	}

	if bytes.HasPrefix(trimmed, []byte("resource \"")) || bytes.HasPrefix(trimmed, []byte("variable \"")) || bytes.HasPrefix(trimmed, []byte("provider \"")) || bytes.HasPrefix(trimmed, []byte("terraform {")) {
		return types.TypeTerraform
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
			return types.TypeGo
		}

		if bytes.Contains(code, []byte("import java")) || bytes.Contains(code, []byte("public class ")) {
			return types.TypeJava
		}

		if bytes.Contains(code, []byte("import ")) || bytes.Contains(code, []byte("fun ")) || bytes.Contains(code, []byte("val ")) {
			return types.TypeKotlin
		}

		if bytes.Contains(code, []byte("import scala.")) || bytes.Contains(code, []byte("object ")) || bytes.Contains(code, []byte("trait ")) {
			return types.TypeScala
		}
	}

	if bytes.HasPrefix(code, []byte("#include <")) || bytes.HasPrefix(code, []byte("#include \"")) {
		if bytes.Contains(code, []byte("std::")) || bytes.Contains(code, []byte("class ")) || bytes.Contains(code, []byte("template <")) {
			return types.TypeCPP
		}

		return types.TypeC
	}

	if bytes.HasPrefix(code, []byte("#import <")) || bytes.HasPrefix(code, []byte("#import \"")) {
		return types.TypeObjectiveC
	}

	if bytes.HasPrefix(code, []byte("using System;")) || bytes.HasPrefix(code, []byte("namespace ")) {
		if bytes.Contains(code, []byte("public class ")) || bytes.Contains(code, []byte("public interface ")) {
			return types.TypeCSharp
		}
	}

	if bytes.HasPrefix(code, []byte("import Foundation")) || bytes.HasPrefix(code, []byte("import UIKit")) || bytes.HasPrefix(code, []byte("import SwiftUI")) {
		return types.TypeSwift
	}

	if bytes.HasPrefix(code, []byte("fn main()")) || bytes.HasPrefix(code, []byte("use std::")) || bytes.HasPrefix(code, []byte("pub fn ")) || bytes.HasPrefix(code, []byte("#![")) {
		return types.TypeRust
	}

	if bytes.HasPrefix(code, []byte("import ")) || bytes.HasPrefix(code, []byte("def ")) || bytes.HasPrefix(code, []byte("class ")) {
		if bytes.Contains(code, []byte("def ")) && bytes.Contains(code, []byte(":")) {
			return types.TypePython
		}
	}

	if bytes.HasPrefix(code, []byte("require '")) || bytes.HasPrefix(code, []byte("require \"")) || bytes.HasPrefix(code, []byte("require_relative ")) {
		if bytes.Contains(code, []byte("def ")) || bytes.Contains(code, []byte("class ")) || bytes.Contains(code, []byte("module ")) {
			return types.TypeRuby
		}
	}

	if bytes.HasPrefix(code, []byte("local ")) || bytes.HasPrefix(code, []byte("function ")) {
		if bytes.Contains(code, []byte("end")) && (bytes.Contains(code, []byte(" then")) || bytes.Contains(code, []byte(" do")) || bytes.Contains(code, []byte("require("))) {
			return types.TypeLua
		}
	}

	if bytes.HasPrefix(code, []byte("import 'package:")) || bytes.HasPrefix(code, []byte("import 'dart:")) {
		return types.TypeDart
	}

	if bytes.HasPrefix(code, []byte("library(")) || bytes.HasPrefix(code, []byte("require(")) {
		if bytes.Contains(code, []byte("<-")) || bytes.Contains(code, []byte("%>%")) {
			return types.TypeR
		}
	}

	if bytes.HasPrefix(code, []byte("const ")) || bytes.HasPrefix(code, []byte("let ")) || bytes.HasPrefix(code, []byte("var ")) || bytes.HasPrefix(code, []byte("import ")) {
		if bytes.Contains(code, []byte("interface ")) || bytes.Contains(code, []byte("type ")) || bytes.Contains(code, []byte(" as ")) {
			if bytes.Contains(code, []byte("=>")) || bytes.Contains(code, []byte("console.log")) || bytes.Contains(code, []byte("from '")) || bytes.Contains(code, []byte("from \"")) {
				return types.TypeTypeScript
			}
		}

		if bytes.Contains(code, []byte("=>")) || bytes.Contains(code, []byte("console.log")) || bytes.Contains(code, []byte("function(")) || bytes.Contains(code, []byte("function ")) {
			return types.TypeJavaScript
		}
	}

	if bytes.HasPrefix(code, []byte("[")) {
		lineEnd := bytes.IndexByte(code, '\n')
		if lineEnd != -1 && bytes.Contains(code[:lineEnd], []byte("]")) {
			if bytes.Contains(code, []byte("=\"")) || bytes.Contains(code, []byte("= \"")) || bytes.Contains(code, []byte("['")) {
				return types.TypeTOML
			}

			return types.TypeINI
		}
	}

	if bytes.HasPrefix(code, []byte("CFLAGS")) || bytes.HasPrefix(code, []byte("TARGET")) || bytes.HasPrefix(code, []byte("all:")) || bytes.HasPrefix(code, []byte(".PHONY:")) {
		return types.TypeMakefile
	}

	if bytes.Contains(code, []byte("Write-Host ")) || bytes.Contains(code, []byte("Get-")) || bytes.Contains(code, []byte("Set-")) || bytes.Contains(code, []byte("Out-Null")) {
		if bytes.Contains(code, []byte("$")) && (bytes.Contains(code, []byte("-eq")) || bytes.Contains(code, []byte("-ne")) || bytes.Contains(code, []byte("param("))) {
			return types.TypePowerShell
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
