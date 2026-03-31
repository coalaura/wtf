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
		content  []byte
		encoding types.TypeID
		isASCII  bool
		isUTF8   bool
		isLatin1 bool
	)

	detectedEncoding, hasEncoding := DetectTextEncoding(b)
	if hasEncoding {
		switch detectedEncoding {
		case types.TypeUTF16LE, types.TypeUTF16BE:
			decoded, _, ok := TryDecodeUTF16(b)
			if ok && len(decoded) > 0 {
				content = decoded
				encoding = detectedEncoding
			} else {
				content = b
			}
		case types.TypeUTF32LE, types.TypeUTF32BE:
			content = b
			encoding = detectedEncoding
		case types.TypeUTF8:
			if b.Len() > 3 {
				content = b[3:]
			} else {
				content = b
			}
			encoding = detectedEncoding
		default:
			content = b
		}
	} else {
		content = b

		isASCII = isLikelyASCIIText(content)
		if !isASCII {
			isUTF8 = isLikelyTextUTF8(content)
		}

		if !isASCII && !isUTF8 && isLikelyLatin1(b) {
			isLatin1 = true
		}
	}

	if encoding == types.TypeNone {
		if isASCII {
			encoding = types.TypeASCII
		} else if isUTF8 {
			encoding = types.TypeUTF8
		} else if isLatin1 {
			encoding = types.TypeLatin1
		} else {
			return nil
		}
	}

	limit := min(len(content), 4096)
	subtype := detectTextSubtype(content[:limit])

	if subtype != types.TypeNone {
		return &types.Metadata{
			Kind:       types.KindText,
			Type:       subtype,
			Confidence: types.ConfidenceMedium,
		}
	}

	return &types.Metadata{
		Kind:       types.KindText,
		Type:       encoding,
		Confidence: types.ConfidenceLow,
	}
}

func detectTextSubtype(data []byte) types.TypeID {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return types.TypeNone
	}

	// 1. Shebangs (Highly Reliable)
	if bytes.HasPrefix(data, []byte("#!")) {
		lineEnd := bytes.IndexByte(data, '\n')
		if lineEnd == -1 {
			lineEnd = len(data)
		}

		shebang := data[:lineEnd]

		if bytes.Contains(shebang, []byte("bash")) || bytes.Contains(shebang, []byte("/sh")) {
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

		if bytes.Contains(shebang, []byte("Rscript")) || bytes.Contains(shebang, []byte("R ")) {
			return types.TypeRScript
		}

		if bytes.Contains(shebang, []byte("fish")) {
			return types.TypeFishScript
		}

		if bytes.Contains(shebang, []byte("zsh")) {
			return types.TypeZshScript
		}

		if bytes.Contains(shebang, []byte("guile")) || bytes.Contains(shebang, []byte("scheme")) {
			return types.TypeGuileScript
		}

		if bytes.Contains(shebang, []byte("emacs")) || bytes.Contains(shebang, []byte("elisp")) {
			return types.TypeEmacsLispScript
		}

		if bytes.Contains(shebang, []byte("clojure")) || bytes.Contains(shebang, []byte("clj")) {
			return types.TypeClojureScript
		}

		if bytes.Contains(shebang, []byte("elixir")) {
			return types.TypeElixirScript
		}

		if bytes.Contains(shebang, []byte("haskell")) || bytes.Contains(shebang, []byte("runhaskell")) {
			return types.TypeHaskellSource
		}

		if bytes.Contains(shebang, []byte("nix")) {
			return types.TypeNixExpression
		}

		if bytes.Contains(shebang, []byte("fsharpi")) || bytes.Contains(shebang, []byte("fsharp")) {
			return types.TypeFSharpSource
		}
	}

	// TopoJSON / GeoJSON
	if bytes.HasPrefix(trimmed, []byte(`{"type":`)) || bytes.HasPrefix(trimmed, []byte(`{"type": `)) {
		if bytes.Contains(trimmed, []byte(`"Topology"`)) || bytes.Contains(trimmed, []byte(`"transform"`)) {
			return types.TypeTopoJSON
		}
		if bytes.Contains(trimmed, []byte(`"FeatureCollection"`)) || bytes.Contains(trimmed, []byte(`"Feature"`)) {
			return types.TypeGeoJSON
		}
	}

	// 2. Clear Prefixes
	if bytes.HasPrefix(trimmed, []byte("<?php")) {
		return types.TypePHPScript
	}

	if bytes.HasPrefix(trimmed, []byte("FROM ")) || bytes.HasPrefix(trimmed, []byte("# syntax=")) {
		return types.TypeDockerfile
	}

	// Docker Compose
	if bytes.Contains(trimmed, []byte("version:")) && bytes.Contains(trimmed, []byte("services:")) {
		return types.TypeDockerComposeConfiguration
	}

	if bytes.HasPrefix(trimmed, []byte("# ")) || bytes.HasPrefix(trimmed, []byte("## ")) {
		// Check for specific markdown patterns
		if bytes.Contains(trimmed, []byte("- [ ]")) || bytes.Contains(trimmed, []byte("- [x]")) {
			return types.TypeMarkdownDocument
		}

		if bytes.Contains(trimmed, []byte("```")) {
			return types.TypeMarkdownDocument
		}

		if bytes.Contains(trimmed, []byte("| ")) && bytes.Contains(trimmed, []byte("|---")) {
			return types.TypeMarkdownDocument
		}

		return types.TypeMarkdownDocument
	}

	if bytes.HasPrefix(trimmed, []byte("---\n")) || bytes.HasPrefix(trimmed, []byte("---\r\n")) {
		if bytes.Contains(trimmed, []byte("title:")) || bytes.Contains(trimmed, []byte("layout:")) {
			return types.TypeMarkdownDocument
		}

		// Check for YAML with common keys
		if bytes.Contains(trimmed, []byte("apiVersion:")) || bytes.Contains(trimmed, []byte("kind:")) {
			return types.TypeYAMLConfiguration
		}

		if bytes.Contains(trimmed, []byte("services:")) || bytes.Contains(trimmed, []byte("version:")) {
			return types.TypeDockerComposeConfiguration
		}

		return types.TypeYAMLConfiguration
	}

	if bytes.HasPrefix(trimmed, []byte("CREATE TABLE")) || bytes.HasPrefix(trimmed, []byte("SELECT ")) || bytes.HasPrefix(trimmed, []byte("INSERT INTO")) || bytes.HasPrefix(trimmed, []byte("-- SQL")) || bytes.HasPrefix(trimmed, []byte("ALTER TABLE")) || bytes.HasPrefix(trimmed, []byte("DROP TABLE")) {
		return types.TypeSQLScript
	}

	if bytes.HasPrefix(trimmed, []byte("@echo off")) || bytes.HasPrefix(trimmed, []byte("@ECHO OFF")) || bytes.HasPrefix(trimmed, []byte("REM ")) || bytes.HasPrefix(trimmed, []byte("rem ")) {
		return types.TypeBatchScript
	}

	if bytes.HasPrefix(trimmed, []byte("cmake_minimum_required")) || bytes.HasPrefix(trimmed, []byte("project(")) {
		return types.TypeCMakeScript
	}

	if bytes.HasPrefix(trimmed, []byte("body {")) || bytes.HasPrefix(trimmed, []byte("@import url(")) || bytes.HasPrefix(trimmed, []byte("@media ")) || bytes.HasPrefix(trimmed, []byte(":root {")) || bytes.HasPrefix(trimmed, []byte("@keyframes")) {
		return types.TypeCSS
	}

	if bytes.HasPrefix(trimmed, []byte("resource \"")) || bytes.HasPrefix(trimmed, []byte("variable \"")) || bytes.HasPrefix(trimmed, []byte("provider \"")) || bytes.HasPrefix(trimmed, []byte("terraform {")) || bytes.HasPrefix(trimmed, []byte("module \"")) {
		if bytes.Contains(trimmed, []byte("module \"")) {
			return types.TypeTerraformModule
		}

		return types.TypeTerraformConfiguration
	}

	if bytes.HasPrefix(trimmed, []byte("schema {")) || bytes.HasPrefix(trimmed, []byte("type ")) && bytes.Contains(trimmed, []byte("{")) && bytes.Contains(trimmed, []byte("}")) {
		return types.TypeGraphQLSchema
	}

	if bytes.HasPrefix(trimmed, []byte("query {")) || bytes.HasPrefix(trimmed, []byte("mutation {")) || bytes.HasPrefix(trimmed, []byte("fragment ")) {
		return types.TypeGraphQL
	}

	// Protocol Buffers
	if bytes.HasPrefix(trimmed, []byte("syntax = \"proto")) || bytes.HasPrefix(trimmed, []byte("package ")) && bytes.Contains(trimmed, []byte("message ")) {
		return types.TypeProtocolBuffer
	}

	// Thrift
	if bytes.HasPrefix(trimmed, []byte("namespace ")) && (bytes.Contains(trimmed, []byte("service ")) || bytes.Contains(trimmed, []byte("struct "))) {
		return types.TypeThriftInterface
	}

	// OpenAPI
	if bytes.Contains(trimmed, []byte("openapi:")) || bytes.Contains(trimmed, []byte("\"openapi\"")) || bytes.Contains(trimmed, []byte("swagger:")) {
		return types.TypeOpenAPISpecification
	}

	// Shellcheck directives
	if bytes.HasPrefix(trimmed, []byte("# shellcheck")) {
		return types.TypeShellcheckDirective
	}

	// Vim scripts
	if bytes.HasPrefix(trimmed, []byte("set ")) || bytes.HasPrefix(trimmed, []byte("let ")) || bytes.HasPrefix(trimmed, []byte("function! ")) || bytes.HasPrefix(trimmed, []byte("call ")) {
		if bytes.Contains(trimmed, []byte("nmap ")) || bytes.Contains(trimmed, []byte("imap ")) || bytes.Contains(trimmed, []byte("vmap ")) || bytes.Contains(trimmed, []byte("syntax ")) {
			return types.TypeVimScript
		}
	}

	// 3. Skip C-style comments for code detection
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

	if bytes.HasPrefix(code, []byte("defmodule ")) {
		return types.TypeElixirScript
	}

	if bytes.HasPrefix(code, []byte("module ")) && (bytes.Contains(code, []byte("where ")) || bytes.Contains(code, []byte("import "))) {
		return types.TypeHaskellSource
	}

	if bytes.HasPrefix(code, []byte("(ns ")) {
		return types.TypeClojureScript
	}

	if bytes.HasPrefix(code, []byte("#include <")) || bytes.HasPrefix(code, []byte("#include \"")) {
		if bytes.Contains(code, []byte("std::")) || bytes.Contains(code, []byte("class ")) || bytes.Contains(code, []byte("template <")) {
			return types.TypeCPPSource
		}

		// Check for Verilog/SystemVerilog
		if bytes.Contains(code, []byte("module ")) && bytes.Contains(code, []byte("endmodule")) {
			if bytes.Contains(code, []byte("logic ")) || bytes.Contains(code, []byte("always_comb")) || bytes.Contains(code, []byte("always_ff")) {
				return types.TypeSystemVerilogSource
			}

			return types.TypeVerilogSource
		}

		// Check for VHDL
		if bytes.Contains(code, []byte("entity ")) && bytes.Contains(code, []byte("end ")) {
			return types.TypeVHDLSource
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

	// Zig
	if bytes.HasPrefix(code, []byte("const std = @import(")) || bytes.Contains(code, []byte("pub fn ")) {
		return types.TypeZigSource
	}

	// Nim
	if bytes.Contains(code, []byte("proc ")) && bytes.Contains(code, []byte("=")) && bytes.Contains(code, []byte("import ")) {
		return types.TypeNimSource
	}

	// Solidity
	if bytes.HasPrefix(code, []byte("pragma solidity")) {
		return types.TypeSoliditySource
	}

	// F#
	if bytes.HasPrefix(code, []byte("open ")) || bytes.HasPrefix(code, []byte("let ")) {
		if bytes.Contains(code, []byte("type ")) && bytes.Contains(code, []byte("member ")) {
			return types.TypeFSharpSource
		}
	}

	// Nix
	if bytes.HasPrefix(code, []byte("{ config")) || bytes.HasPrefix(code, []byte("with import")) || bytes.HasPrefix(code, []byte("let ")) && bytes.Contains(code, []byte("in ")) {
		if bytes.Contains(code, []byte("pkgs.")) || bytes.Contains(code, []byte("stdenv")) || bytes.Contains(code, []byte("fetchurl")) {
			return types.TypeNixExpression
		}
	}

	// Coq, Lean, Idris
	if bytes.HasPrefix(code, []byte("Theorem ")) || bytes.HasPrefix(code, []byte("Lemma ")) || bytes.HasPrefix(code, []byte("Definition ")) {
		if bytes.Contains(code, []byte("Proof.")) {
			return types.TypeCoqSource
		}
	}

	if bytes.HasPrefix(code, []byte("theorem ")) || bytes.HasPrefix(code, []byte("lemma ")) || bytes.HasPrefix(code, []byte("def ")) {
		if bytes.Contains(code, []byte("by ")) {
			return types.TypeLeanSource
		}
	}

	if bytes.HasPrefix(code, []byte("module ")) && bytes.Contains(code, []byte("where ")) {
		if bytes.Contains(code, []byte("data ")) || bytes.Contains(code, []byte("total ")) {
			return types.TypeIdrisSource
		}
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

	// React/Vue/Angular/Svelte components
	if bytes.HasPrefix(code, []byte("import React")) || bytes.HasPrefix(code, []byte("from 'react")) || bytes.HasPrefix(code, []byte("from \"react")) {
		if bytes.Contains(code, []byte("export default function")) || bytes.Contains(code, []byte("const ")) && bytes.Contains(code, []byte("= ()")) {
			return types.TypeReactComponent
		}
	}

	if bytes.HasPrefix(code, []byte("<template>")) || (bytes.HasPrefix(code, []byte("<script")) && bytes.Contains(code, []byte("</template>"))) {
		if bytes.Contains(code, []byte("export default")) {
			return types.TypeVueComponent
		}
	}

	if bytes.HasPrefix(code, []byte("@Component(")) || bytes.HasPrefix(code, []byte("@NgModule(")) {
		return types.TypeAngularComponent
	}

	if bytes.HasPrefix(code, []byte("<script>")) || bytes.HasPrefix(code, []byte("<svelte:")) {
		return types.TypeSvelteComponent
	}

	// Templates
	if bytes.HasPrefix(code, []byte("{% extends")) || bytes.HasPrefix(code, []byte("{% block")) || bytes.HasPrefix(code, []byte("{% load")) {
		return types.TypeDjangoTemplate
	}

	if bytes.HasPrefix(code, []byte("{# ")) || bytes.HasPrefix(code, []byte("{% ")) || bytes.HasPrefix(code, []byte("{{ ")) {
		return types.TypeJinjaTemplate
	}

	if bytes.HasPrefix(code, []byte("{{define ")) || bytes.HasPrefix(code, []byte("{{if ")) || bytes.HasPrefix(code, []byte("{{range ")) || bytes.HasPrefix(code, []byte("{{template ")) {
		return types.TypeGoTemplate
	}

	if bytes.HasPrefix(code, []byte("{{#each ")) || bytes.HasPrefix(code, []byte("{{#if ")) || bytes.HasPrefix(code, []byte("{{#with ")) {
		return types.TypeHandlebarsTemplate
	}

	// LaTeX/TeX
	if bytes.HasPrefix(code, []byte("\\documentclass")) || bytes.HasPrefix(code, []byte("\\begin{document}")) {
		return types.TypeLaTeXDocument
	}

	if bytes.HasPrefix(code, []byte("\\input ")) || bytes.HasPrefix(code, []byte("\\include ")) {
		return types.TypeTeXDocument
	}

	// reStructuredText
	if bytes.HasPrefix(code, []byte(".. ")) || bytes.HasPrefix(code, []byte("=======")) || bytes.HasPrefix(code, []byte("-------")) {
		return types.TypeReStructuredText
	}

	// AsciiDoc
	if bytes.HasPrefix(code, []byte("= ")) || bytes.HasPrefix(code, []byte("== ")) || bytes.HasPrefix(code, []byte(":")) && bytes.Contains(code, []byte(": ")) {
		return types.TypeAsciiDoc
	}

	// Org Mode
	if bytes.HasPrefix(code, []byte("#+TITLE:")) || bytes.HasPrefix(code, []byte("#+BEGIN_")) || bytes.HasPrefix(code, []byte("* ")) {
		return types.TypeOrgMode
	}

	// JSONC (JSON with comments)
	if (trimmed[0] == '{' || trimmed[0] == '[') && (bytes.Contains(trimmed, []byte("//")) || bytes.Contains(trimmed, []byte("/*"))) {
		if bytes.Contains(trimmed, []byte(`":`)) || bytes.Contains(trimmed, []byte(`": `)) {
			return types.TypeJSONCSource
		}
	}

	// JSON5
	if (trimmed[0] == '{' || trimmed[0] == '[') && bytes.Contains(trimmed, []byte("':")) {
		return types.TypeJSON5Source
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

	// Config file patterns
	if bytes.HasPrefix(code, []byte("[")) {
		lineEnd := bytes.IndexByte(code, '\n')
		if lineEnd != -1 && bytes.Contains(code[:lineEnd], []byte("]")) {
			if bytes.Contains(code, []byte("=\"")) || bytes.Contains(code, []byte("= \"")) || bytes.Contains(code, []byte("['")) {
				return types.TypeTOMLConfiguration
			}

			// Git config
			if bytes.Contains(code, []byte("[remote")) || bytes.Contains(code, []byte("[branch")) || bytes.Contains(code, []byte("[core")) || bytes.Contains(code, []byte("[user")) {
				return types.TypeGitConfig
			}

			// SSH config
			if bytes.Contains(code, []byte("Host ")) || bytes.Contains(code, []byte("[host")) {
				return types.TypeSSHConfig
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

	// Properties file (key=value or key: value)
	if bytes.HasPrefix(code, []byte("#")) && bytes.Contains(trimmed, []byte("=")) && !bytes.Contains(trimmed, []byte("==")) {
		return types.TypePropertiesFile
	}

	// CSV detection
	if bytes.Count(trimmed, []byte(",")) > 3 && bytes.Count(trimmed, []byte("\n")) > 1 {
		lines := bytes.Split(trimmed, []byte("\n"))
		if len(lines) > 1 {
			firstLineCommas := bytes.Count(lines[0], []byte(","))
			secondLineCommas := bytes.Count(lines[1], []byte(","))

			if firstLineCommas > 0 && firstLineCommas == secondLineCommas {
				return types.TypeCSVData
			}
		}
	}

	// TSV detection
	if bytes.Count(trimmed, []byte("\t")) > 5 && bytes.Count(trimmed, []byte("\n")) > 2 {
		lines := bytes.Split(trimmed, []byte("\n"))
		if len(lines) > 2 {
			firstLineTabs := bytes.Count(lines[0], []byte("\t"))
			secondLineTabs := bytes.Count(lines[1], []byte("\t"))

			if firstLineTabs >= 2 && firstLineTabs == secondLineTabs {
				return types.TypeTSVData
			}
		}
	}

	// Log file detection
	if bytes.Contains(trimmed, []byte(" [INFO] ")) || bytes.Contains(trimmed, []byte(" [WARN] ")) || bytes.Contains(trimmed, []byte(" [ERROR] ")) || bytes.Contains(trimmed, []byte(" [DEBUG] ")) {
		return types.TypeLogData
	}

	// Apache log
	if bytes.Contains(trimmed, []byte("\"GET ")) || bytes.Contains(trimmed, []byte("\"POST ")) || bytes.Contains(trimmed, []byte("\"PUT ")) {
		if bytes.Contains(trimmed, []byte("HTTP/")) {
			return types.TypeApacheLog
		}
	}

	// Nginx log
	if bytes.Contains(trimmed, []byte(" - - [")) && bytes.Contains(trimmed, []byte("] \"")) {
		return types.TypeNginxLog
	}

	// HTTP log
	if bytes.Contains(trimmed, []byte("HTTP/1.")) && bytes.Contains(trimmed, []byte(" 200 ")) || bytes.Contains(trimmed, []byte(" 301 ")) || bytes.Contains(trimmed, []byte(" 404 ")) {
		return types.TypeHTTPLog
	}

	// Diff/Git diff
	if bytes.HasPrefix(trimmed, []byte("diff --git")) || bytes.HasPrefix(trimmed, []byte("index ")) || bytes.HasPrefix(trimmed, []byte("--- a/")) || bytes.HasPrefix(trimmed, []byte("+++ b/")) {
		return types.TypeGitDiff
	}

	if bytes.HasPrefix(trimmed, []byte("--- ")) || bytes.HasPrefix(trimmed, []byte("+++ ")) || bytes.HasPrefix(trimmed, []byte("@@ ")) {
		return types.TypeDiffPatch
	}

	// Commit message
	if bytes.HasPrefix(trimmed, []byte("Merge ")) || (bytes.Contains(trimmed, []byte("\n\n")) && !bytes.Contains(trimmed, []byte("{")) && !bytes.Contains(trimmed, []byte("("))) {
		lines := bytes.Split(trimmed, []byte("\n"))
		if len(lines) > 0 && len(lines[0]) < 80 && !bytes.Contains(lines[0], []byte(":")) && !bytes.Contains(lines[0], []byte("=")) {
			return types.TypeCommitMessage
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
