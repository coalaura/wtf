package types

type TypeID uint16

const (
	TypeNone TypeID = iota
	Type32Bit
	Type32BitBigEndian
	Type32BitLittleEndian
	Type3G2Video
	Type3GPPVideo
	Type3MFDocument
	Type64Bit
	Type64BitBigEndian
	Type64BitLittleEndian
	TypeAC3
	TypeAdobeDNG
	TypeADTS
	TypeAfterEffectsProject
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAndroidAppBundle
	TypeAndroidArchive
	TypeAndroidPackage
	TypeAndroidPackageX
	TypeAndroidSplitAPKSet
	TypeAndroidSystemPackage
	TypeAngularComponent
	TypeApacheLog
	TypeAPPXPackage
	TypeArchLinuxPackage
	TypeASCII
	TypeAsciiDoc
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeB4
	TypeBashScript
	TypeBatchScript
	TypeBGZF
	TypeBigEndian
	TypeBigTIFF
	TypeBinary
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBlackmagicRAWVideo
	TypeBlockDevice
	TypeByteSwapped
	TypeCanonRAW3Image
	TypeCAT
	TypeCDAAudio
	TypeCE
	TypeCharacterDevice
	TypeClojureScript
	TypeCMakeScript
	TypeCodestream
	TypeComicBook
	TypeCommitMessage
	TypeCondaPackage
	TypeContainer
	TypeCoqSource
	TypeCorelDRAWDocument
	TypeCOWD
	TypeCPPSource
	TypeCSharpSource
	TypeCSource
	TypeCSS
	TypeCSVData
	TypeCubaseProject
	TypeDartSource
	TypeDiffPatch
	TypeDirectory
	TypeDjangoTemplate
	TypeDockerComposeConfiguration
	TypeDockerfile
	TypeDownloadableSounds
	TypeDSA
	TypeEAC3
	TypeEC
	TypeElixirScript
	TypeEmacsLispScript
	TypeEmpty
	TypeEncrypted
	TypeEnhancedMetafile
	TypeEPUBDocument
	TypeExt2
	TypeExt3
	TypeExt4
	TypeF4VVideo
	TypeFabricMod
	TypeFAT12
	TypeFAT16
	TypeFAT32
	TypeFirefoxExtension
	TypeFishScript
	TypeFLACAudio
	TypeForgeMod
	TypeFSharpSource
	TypeGIF87a
	TypeGIF89a
	TypeGitConfig
	TypeGitDiff
	TypeGoSource
	TypeGoTemplate
	TypeGraphQL
	TypeGraphQLSchema
	TypeGuileScript
	TypeHandlebarsTemplate
	TypeHaskellSource
	TypeHDPhoto
	TypeHE
	TypeHEIFImage
	TypeHighEntropy
	TypeHTTPLog
	TypeIdrisSource
	TypeILBMImage
	TypeIMMM
	TypeINIConfiguration
	TypeIOSApplicationArchive
	TypeIWAD
	TypeJavaArchive
	TypeJavaEnterpriseArchive
	TypeJavaScriptSource
	TypeJavaSource
	TypeJavaWebArchive
	TypeJinjaTemplate
	TypeJSON5Source
	TypeJSONCSource
	TypeKDB
	TypeKDBX
	TypeKDM
	TypeKerasModel
	TypeKMZArchive
	TypeKotlinSource
	TypeKritaDocument
	TypeLaTeXDocument
	TypeLatin1
	TypeLeanSource
	TypeLegacy
	TypeLIST
	TypeLittleEndian
	TypeLogData
	TypeLOVEGame
	TypeLuaScript
	TypeLZ4Legacy
	TypeLZMACompressed
	TypeM2TSBDAV
	TypeM4VVideo
	TypeMakefile
	TypeMarkdownDocument
	TypeMatroskaVideo
	TypeMicrosoftExcelAddIn
	TypeMicrosoftExcelMacroEnabledTemplate
	TypeMicrosoftExcelMacroEnabledWorkbook
	TypeMicrosoftExcelTemplate
	TypeMicrosoftExcelWorkbook
	TypeMicrosoftInstaller
	TypeMicrosoftInstallerPatch
	TypeMicrosoftOutlookMessage
	TypeMicrosoftPowerPointAddIn
	TypeMicrosoftPowerPointMacroEnabledPresentation
	TypeMicrosoftPowerPointMacroEnabledSlideshow
	TypeMicrosoftPowerPointMacroEnabledTemplate
	TypeMicrosoftPowerPointPresentation
	TypeMicrosoftPowerPointSlideshow
	TypeMicrosoftPowerPointTemplate
	TypeMicrosoftProjectDocument
	TypeMicrosoftPublisherDocument
	TypeMicrosoftVisioDrawing
	TypeMicrosoftWordDocument
	TypeMicrosoftWordMacroEnabledDocument
	TypeMicrosoftWordMacroEnabledTemplate
	TypeMicrosoftWordTemplate
	TypeMIDIAudio
	TypeMinecraftResourcePack
	TypeMotionJPEG2000Video
	TypeMP3ID3Tagged
	TypeMP4Video
	TypeMPEG12Video
	TypeMPEG4Audio
	TypeMPEGLayer1
	TypeMPEGLayer2
	TypeMPEGLayer3
	TypeMSIXPackage
	TypeNamedPipe
	TypeNanosecondBigEndian
	TypeNanosecondLittleEndian
	TypeNewASCII
	TypeNewASCIIWithCRC
	TypeNginxLog
	TypeNikonRAW
	TypeNixExpression
	TypeNpmPackage
	TypeNuGetPackage
	TypeOBJ
	TypeObjectiveCSource
	TypeOCIImageLayout
	TypeOldASCII
	TypeOlympusRAW
	TypeOpenAPISpecification
	TypeOpenDocumentChart
	TypeOpenDocumentDatabase
	TypeOpenDocumentFormula
	TypeOpenDocumentGraphics
	TypeOpenDocumentImage
	TypeOpenDocumentPresentation
	TypeOpenDocumentSpreadsheet
	TypeOpenDocumentText
	TypeOpenRasterImage
	TypeOpusAudio
	TypeOrgMode
	TypePAM
	TypePanasonicRAW
	TypePBMASCII
	TypePBMBinary
	TypePE32ARM
	TypePE32ARM64
	TypePE32ARMv7
	TypePE32Itanium
	TypePE32PlusARM64
	TypePE32PlusUnknown
	TypePE32PlusX8664
	TypePE32Unknown
	TypePE32X86
	TypePE32X8664
	TypePEC
	TypePentaxRAW
	TypePerlScript
	TypePES
	TypePGMASCII
	TypePGMBinary
	TypePHPScript
	TypePowerShellScript
	TypePPMASCII
	TypePPMBinary
	TypePropertiesFile
	TypeProtocolBuffer
	TypePSB
	TypePSD
	TypePWAD
	TypePythonScript
	TypePythonSourceDistribution
	TypePythonWheel
	TypePyTorchModel
	TypeQCOW1
	TypeQCOW2
	TypeQCPAudio
	TypeQuickTimeMovie
	TypeReactComponent
	TypeReStructuredText
	TypeRF64Audio
	TypeRIFFPalette
	TypeRPF0
	TypeRPF2
	TypeRPF3
	TypeRPF4
	TypeRPF6
	TypeRPF7
	TypeRSA
	TypeRScript
	TypeRubyScript
	TypeRustSource
	TypeScalaSource
	TypeShellcheckDirective
	TypeSkinnableFrame
	TypeSlackwarePackage
	TypeSocket
	TypeSonyRAW
	TypeSonyRAWSR2
	TypeSoundFont2
	TypeSpecial
	TypeSpeexAudio
	TypeSQLScript
	TypeSSHConfig
	TypeSvelteComponent
	TypeSwiftSource
	TypeSymbolicLink
	TypeSystemVerilogSource
	TypeTerraformConfiguration
	TypeTerraformModule
	TypeTeXDocument
	TypeTheoraVideo
	TypeThriftInterface
	TypeTOMLConfiguration
	TypeTSVData
	TypeTypeScriptSource
	TypeUncompressed
	TypeUTF16BE
	TypeUTF16LE
	TypeUTF32BE
	TypeUTF32LE
	TypeUTF8
	TypeVagrantBox
	TypeVerilogSource
	TypeVersion035
	TypeVersion036
	TypeVersion037
	TypeVersion038
	TypeVersion039
	TypeVersion040
	TypeVersion041
	TypeVersion1
	TypeVersion2
	TypeVersion3
	TypeVersion5
	TypeVersion7
	TypeVersion8
	TypeVHDLSource
	TypeVimScript
	TypeVisualStudioExtension
	TypeVMDK
	TypeVMDKDescription
	TypeVorbisAudio
	TypeVueComponent
	TypeWAVAudio
	TypeWebMVideo
	TypeWebPImage
	TypeWindowsAnimatedCursor
	TypeWindowsCursor
	TypeWindowsIcon
	TypeWindowsLE
	TypeWindowsLX
	TypeWindowsMetafile
	TypeWindowsNE
	TypeWorkspace
	TypeWrapper
	TypeXMLPaperSpecification
	TypeXSLTStylesheet
	TypeYAMLConfiguration
	TypeZlibCompressed
	TypeZshScript
)

var typeNames = [...]string{
	TypeNone:                               "",
	Type32Bit:                              "32-bit",
	Type32BitBigEndian:                     "32-bit Big-Endian",
	Type32BitLittleEndian:                  "32-bit Little-Endian",
	Type3G2Video:                           "3G2 Video",
	Type3GPPVideo:                          "3GPP Video",
	Type3MFDocument:                        "3MF Document",
	Type64Bit:                              "64-bit",
	Type64BitBigEndian:                     "64-bit Big-Endian",
	Type64BitLittleEndian:                  "64-bit Little-Endian",
	TypeAC3:                                "AC-3",
	TypeAdobeDNG:                           "Adobe DNG",
	TypeADTS:                               "ADTS",
	TypeAfterEffectsProject:                "After Effects Project",
	TypeAIFCAudio:                          "AIFC Audio",
	TypeAIFFAudio:                          "AIFF Audio",
	TypeAndroidAppBundle:                   "AAB Package",
	TypeAndroidArchive:                     "AAR Package",
	TypeAndroidPackage:                     "APK Package",
	TypeAndroidPackageX:                    "XAPK Package",
	TypeAndroidSplitAPKSet:                 "APKS Package",
	TypeAndroidSystemPackage:               "Android System Package",
	TypeAngularComponent:                   "Angular Component",
	TypeApacheLog:                          "Apache Log",
	TypeAPPXPackage:                        "APPX Package",
	TypeArchLinuxPackage:                   "Arch Linux Package",
	TypeASCII:                              "ASCII",
	TypeAsciiDoc:                           "AsciiDoc",
	TypeAVIFImage:                          "AVIF Image",
	TypeAVIFImageSequence:                  "AVIF Image Sequence",
	TypeAVIVideo:                           "AVI Video",
	TypeB4:                                 "B4",
	TypeBashScript:                         "Bash Script",
	TypeBatchScript:                        "Batch Script",
	TypeBGZF:                               "BGZF",
	TypeBigEndian:                          "Big-Endian",
	TypeBigTIFF:                            "BigTIFF",
	TypeBinary:                             "Binary",
	TypeBinaryBigEndian:                    "Binary Big-Endian",
	TypeBinaryLittleEndian:                 "Binary Little-Endian",
	TypeBlackmagicRAWVideo:                 "Blackmagic RAW Video",
	TypeBlockDevice:                        "Block Device",
	TypeByteSwapped:                        "Byte-Swapped",
	TypeCanonRAW3Image:                     "Canon RAW 3 Image",
	TypeCAT:                                "CAT",
	TypeCDAAudio:                           "CD Audio",
	TypeCE:                                 "CE",
	TypeCharacterDevice:                    "Character Device",
	TypeClojureScript:                      "Clojure Script",
	TypeCMakeScript:                        "CMake Script",
	TypeCodestream:                         "Codestream",
	TypeComicBook:                          "Comic Book",
	TypeCommitMessage:                      "Commit Message",
	TypeCondaPackage:                       "Conda Package",
	TypeContainer:                          "Container",
	TypeCoqSource:                          "Coq Source",
	TypeCorelDRAWDocument:                  "CorelDRAW Document",
	TypeCOWD:                               "COWD",
	TypeCPPSource:                          "C++ Source",
	TypeCSharpSource:                       "C# Source",
	TypeCSource:                            "C Source",
	TypeCSS:                                "Cascading Style Sheets",
	TypeCSVData:                            "CSV Data",
	TypeCubaseProject:                      "Cubase Project",
	TypeDartSource:                         "Dart Source",
	TypeDiffPatch:                          "Diff/Patch",
	TypeDirectory:                          "Directory",
	TypeDjangoTemplate:                     "Django Template",
	TypeDockerComposeConfiguration:         "Docker Compose Configuration",
	TypeDockerfile:                         "Dockerfile",
	TypeDownloadableSounds:                 "Downloadable Sounds",
	TypeDSA:                                "DSA",
	TypeEAC3:                               "E-AC-3",
	TypeEC:                                 "EC",
	TypeElixirScript:                       "Elixir Script",
	TypeEmacsLispScript:                    "Emacs Lisp Script",
	TypeEmpty:                              "Empty",
	TypeEncrypted:                          "Encrypted",
	TypeEnhancedMetafile:                   "Enhanced Metafile",
	TypeEPUBDocument:                       "EPUB Document",
	TypeExt2:                               "ext2",
	TypeExt3:                               "ext3",
	TypeExt4:                               "ext4",
	TypeF4VVideo:                           "F4V Video",
	TypeFabricMod:                          "Fabric Mod",
	TypeFAT12:                              "FAT12",
	TypeFAT16:                              "FAT16",
	TypeFAT32:                              "FAT32",
	TypeFirefoxExtension:                   "Firefox Extension",
	TypeFishScript:                         "Fish Script",
	TypeFLACAudio:                          "FLAC Audio",
	TypeForgeMod:                           "Forge Mod",
	TypeFSharpSource:                       "F# Source",
	TypeGIF87a:                             "GIF87a",
	TypeGIF89a:                             "GIF89a",
	TypeGitConfig:                          "Git Configuration",
	TypeGitDiff:                            "Git Diff",
	TypeGoSource:                           "Go Source",
	TypeGoTemplate:                         "Go Template",
	TypeGraphQL:                            "GraphQL",
	TypeGraphQLSchema:                      "GraphQL Schema",
	TypeGuileScript:                        "Guile Script",
	TypeHandlebarsTemplate:                 "Handlebars Template",
	TypeHaskellSource:                      "Haskell Source",
	TypeHDPhoto:                            "HD Photo",
	TypeHE:                                 "HE",
	TypeHEIFImage:                          "HEIF Image",
	TypeHighEntropy:                        "High Entropy",
	TypeHTTPLog:                            "HTTP Log",
	TypeIdrisSource:                        "Idris Source",
	TypeILBMImage:                          "ILBM Image",
	TypeIMMM:                               "IMMM",
	TypeINIConfiguration:                   "INI Configuration",
	TypeIOSApplicationArchive:              "iOS Application Archive",
	TypeIWAD:                               "IWAD",
	TypeJavaArchive:                        "Java Archive",
	TypeJavaEnterpriseArchive:              "Java Enterprise Archive",
	TypeJavaScriptSource:                   "JavaScript Source",
	TypeJavaSource:                         "Java Source",
	TypeJavaWebArchive:                     "Java Web Archive",
	TypeJinjaTemplate:                      "Jinja Template",
	TypeJSON5Source:                        "JSON5 Source",
	TypeJSONCSource:                        "JSONC Source",
	TypeKDB:                                "KDB",
	TypeKDBX:                               "KDBX",
	TypeKDM:                                "KDM",
	TypeKerasModel:                         "Keras Model",
	TypeKMZArchive:                         "KMZ Archive",
	TypeKotlinSource:                       "Kotlin Source",
	TypeKritaDocument:                      "Krita Document",
	TypeLaTeXDocument:                      "LaTeX Document",
	TypeLatin1:                             "Latin-1",
	TypeLeanSource:                         "Lean Source",
	TypeLegacy:                             "Legacy",
	TypeLIST:                               "LIST",
	TypeLittleEndian:                       "Little-Endian",
	TypeLogData:                            "Log Data",
	TypeLOVEGame:                           "LÖVE Game",
	TypeLuaScript:                          "Lua Script",
	TypeLZ4Legacy:                          "LZ4 Legacy",
	TypeLZMACompressed:                     "LZMA Compressed",
	TypeM2TSBDAV:                           "M2TS/BDAV",
	TypeM4VVideo:                           "M4V Video",
	TypeMakefile:                           "Makefile",
	TypeMarkdownDocument:                   "Markdown Document",
	TypeMatroskaVideo:                      "Matroska Video",
	TypeMicrosoftExcelAddIn:                "Microsoft Excel Add-In",
	TypeMicrosoftExcelMacroEnabledTemplate: "Microsoft Excel Macro-Enabled Template",
	TypeMicrosoftExcelMacroEnabledWorkbook: "Microsoft Excel Macro-Enabled Workbook",
	TypeMicrosoftExcelTemplate:             "Microsoft Excel Template",
	TypeMicrosoftExcelWorkbook:             "Microsoft Excel Workbook",
	TypeMicrosoftInstaller:                 "Microsoft Installer",
	TypeMicrosoftInstallerPatch:            "Microsoft Installer Patch",
	TypeMicrosoftOutlookMessage:            "Microsoft Outlook Message",
	TypeMicrosoftPowerPointAddIn:           "Microsoft PowerPoint Add-In",
	TypeMicrosoftPowerPointMacroEnabledPresentation: "Microsoft PowerPoint Macro-Enabled Presentation",
	TypeMicrosoftPowerPointMacroEnabledSlideshow:    "Microsoft PowerPoint Macro-Enabled Slideshow",
	TypeMicrosoftPowerPointMacroEnabledTemplate:     "Microsoft PowerPoint Macro-Enabled Template",
	TypeMicrosoftPowerPointPresentation:             "Microsoft PowerPoint Presentation",
	TypeMicrosoftPowerPointSlideshow:                "Microsoft PowerPoint Slideshow",
	TypeMicrosoftPowerPointTemplate:                 "Microsoft PowerPoint Template",
	TypeMicrosoftProjectDocument:                    "Microsoft Project Document",
	TypeMicrosoftPublisherDocument:                  "Microsoft Publisher Document",
	TypeMicrosoftVisioDrawing:                       "Microsoft Visio Drawing",
	TypeMicrosoftWordDocument:                       "Microsoft Word Document",
	TypeMicrosoftWordMacroEnabledDocument:           "Microsoft Word Macro-Enabled Document",
	TypeMicrosoftWordMacroEnabledTemplate:           "Microsoft Word Macro-Enabled Template",
	TypeMicrosoftWordTemplate:                       "Microsoft Word Template",
	TypeMIDIAudio:                                   "MIDI Audio",
	TypeMinecraftResourcePack:                       "Minecraft Resource Pack",
	TypeMotionJPEG2000Video:                         "Motion JPEG 2000 Video",
	TypeMP3ID3Tagged:                                "MP3 (ID3 Tagged)",
	TypeMP4Video:                                    "MP4 Video",
	TypeMPEG12Video:                                 "MPEG-1/2 Video",
	TypeMPEG4Audio:                                  "MPEG-4 Audio",
	TypeMPEGLayer1:                                  "MPEG Layer I",
	TypeMPEGLayer2:                                  "MPEG Layer II",
	TypeMPEGLayer3:                                  "MPEG Layer III",
	TypeMSIXPackage:                                 "MSIX Package",
	TypeNamedPipe:                                   "Named Pipe",
	TypeNanosecondBigEndian:                         "Nanosecond Big-Endian",
	TypeNanosecondLittleEndian:                      "Nanosecond Little-Endian",
	TypeNewASCII:                                    "New ASCII",
	TypeNewASCIIWithCRC:                             "New ASCII with CRC",
	TypeNginxLog:                                    "Nginx Log",
	TypeNikonRAW:                                    "Nikon RAW",
	TypeNixExpression:                               "Nix Expression",
	TypeNpmPackage:                                  "npm Package",
	TypeNuGetPackage:                                "NuGet Package",
	TypeOBJ:                                         "COFF Object",
	TypeObjectiveCSource:                            "Objective-C Source",
	TypeOCIImageLayout:                              "OCI Image Layout",
	TypeOldASCII:                                    "Old ASCII",
	TypeOlympusRAW:                                  "Olympus RAW",
	TypeOpenAPISpecification:                        "OpenAPI Specification",
	TypeOpenDocumentChart:                           "OpenDocument Chart",
	TypeOpenDocumentDatabase:                        "OpenDocument Database",
	TypeOpenDocumentFormula:                         "OpenDocument Formula",
	TypeOpenDocumentGraphics:                        "OpenDocument Graphics",
	TypeOpenDocumentImage:                           "OpenDocument Image",
	TypeOpenDocumentPresentation:                    "OpenDocument Presentation",
	TypeOpenDocumentSpreadsheet:                     "OpenDocument Spreadsheet",
	TypeOpenDocumentText:                            "OpenDocument Text",
	TypeOpenRasterImage:                             "OpenRaster Image",
	TypeOpusAudio:                                   "Opus Audio",
	TypeOrgMode:                                     "Org Mode",
	TypePAM:                                         "PAM",
	TypePanasonicRAW:                                "Panasonic RAW",
	TypePBMASCII:                                    "PBM ASCII",
	TypePBMBinary:                                   "PBM binary",
	TypePE32ARM:                                     "PE32 ARM",
	TypePE32ARM64:                                   "PE32 ARM64",
	TypePE32ARMv7:                                   "PE32 ARMv7",
	TypePE32Itanium:                                 "PE32 Itanium",
	TypePE32PlusARM64:                               "PE32+ ARM64",
	TypePE32PlusUnknown:                             "PE32+ Unknown",
	TypePE32PlusX8664:                               "PE32+ x86-64",
	TypePE32Unknown:                                 "PE32 Unknown",
	TypePE32X86:                                     "PE32 x86",
	TypePE32X8664:                                   "PE32 x86-64",
	TypePEC:                                         "PEC",
	TypePentaxRAW:                                   "Pentax RAW",
	TypePerlScript:                                  "Perl Script",
	TypePES:                                         "PES",
	TypePGMASCII:                                    "PGM ASCII",
	TypePGMBinary:                                   "PGM binary",
	TypePHPScript:                                   "PHP Script",
	TypePowerShellScript:                            "PowerShell Script",
	TypePPMASCII:                                    "PPM ASCII",
	TypePPMBinary:                                   "PPM binary",
	TypePropertiesFile:                              "Properties File",
	TypeProtocolBuffer:                              "Protocol Buffer",
	TypePSB:                                         "PSB",
	TypePSD:                                         "PSD",
	TypePWAD:                                        "PWAD",
	TypePythonScript:                                "Python Script",
	TypePythonSourceDistribution:                    "Python Source Distribution",
	TypePythonWheel:                                 "Python Wheel",
	TypePyTorchModel:                                "PyTorch Model",
	TypeQCOW1:                                       "QCOW1",
	TypeQCOW2:                                       "QCOW2",
	TypeQCPAudio:                                    "QCP Audio",
	TypeQuickTimeMovie:                              "QuickTime Movie",
	TypeReactComponent:                              "React Component",
	TypeReStructuredText:                            "reStructuredText",
	TypeRF64Audio:                                   "RF64 Audio",
	TypeRIFFPalette:                                 "Palette",
	TypeRPF0:                                        "RPF0",
	TypeRPF2:                                        "RPF2",
	TypeRPF3:                                        "RPF3",
	TypeRPF4:                                        "RPF4",
	TypeRPF6:                                        "RPF6",
	TypeRPF7:                                        "RPF7",
	TypeRSA:                                         "RSA",
	TypeRScript:                                     "R Script",
	TypeRubyScript:                                  "Ruby Script",
	TypeRustSource:                                  "Rust Source",
	TypeScalaSource:                                 "Scala Source",
	TypeShellcheckDirective:                         "Shellcheck Directive",
	TypeSkinnableFrame:                              "Skinnable Frame",
	TypeSlackwarePackage:                            "Slackware Package",
	TypeSocket:                                      "Socket",
	TypeSonyRAW:                                     "Sony RAW",
	TypeSonyRAWSR2:                                  "Sony SR2 RAW",
	TypeSoundFont2:                                  "SoundFont 2",
	TypeSpecial:                                     "Special",
	TypeSpeexAudio:                                  "Speex Audio",
	TypeSQLScript:                                   "SQL Script",
	TypeSSHConfig:                                   "SSH Configuration",
	TypeSvelteComponent:                             "Svelte Component",
	TypeSwiftSource:                                 "Swift Source",
	TypeSymbolicLink:                                "Symbolic Link",
	TypeSystemVerilogSource:                         "SystemVerilog Source",
	TypeTerraformConfiguration:                      "Terraform Configuration",
	TypeTerraformModule:                             "Terraform Module",
	TypeTeXDocument:                                 "TeX Document",
	TypeTheoraVideo:                                 "Theora Video",
	TypeThriftInterface:                             "Thrift Interface",
	TypeTOMLConfiguration:                           "TOML Configuration",
	TypeTSVData:                                     "TSV Data",
	TypeTypeScriptSource:                            "TypeScript Source",
	TypeUncompressed:                                "Uncompressed",
	TypeUTF16BE:                                     "UTF-16 Big-Endian",
	TypeUTF16LE:                                     "UTF-16 Little-Endian",
	TypeUTF32BE:                                     "UTF-32 Big-Endian",
	TypeUTF32LE:                                     "UTF-32 Little-Endian",
	TypeUTF8:                                        "UTF-8",
	TypeVagrantBox:                                  "Vagrant Box",
	TypeVerilogSource:                               "Verilog Source",
	TypeVersion035:                                  "Version 035",
	TypeVersion036:                                  "Version 036",
	TypeVersion037:                                  "Version 037",
	TypeVersion038:                                  "Version 038",
	TypeVersion039:                                  "Version 039",
	TypeVersion040:                                  "Version 040",
	TypeVersion041:                                  "Version 041",
	TypeVersion1:                                    "Version 1",
	TypeVersion2:                                    "Version 2",
	TypeVersion3:                                    "Version 3",
	TypeVersion5:                                    "Version 5",
	TypeVersion7:                                    "Version 7",
	TypeVersion8:                                    "Version 8",
	TypeVHDLSource:                                  "VHDL Source",
	TypeVimScript:                                   "Vim Script",
	TypeVisualStudioExtension:                       "Visual Studio Extension",
	TypeVMDK:                                        "VMDK",
	TypeVMDKDescription:                             "VMDK Description",
	TypeVorbisAudio:                                 "Vorbis Audio",
	TypeVueComponent:                                "Vue Component",
	TypeWAVAudio:                                    "WAV Audio",
	TypeWebMVideo:                                   "WebM Video",
	TypeWebPImage:                                   "WebP Image",
	TypeWindowsAnimatedCursor:                       "Windows Animated Cursor",
	TypeWindowsCursor:                               "Windows Cursor",
	TypeWindowsIcon:                                 "Windows Icon",
	TypeWindowsLE:                                   "Linear Executable (LE)",
	TypeWindowsLX:                                   "OS/2 Linear Executable (LX)",
	TypeWindowsMetafile:                             "Windows Metafile",
	TypeWindowsNE:                                   "16-bit New Executable (NE)",
	TypeWorkspace:                                   "Workspace",
	TypeWrapper:                                     "Wrapper",
	TypeXMLPaperSpecification:                       "XML Paper Specification",
	TypeXSLTStylesheet:                              "XSLT Stylesheet",
	TypeYAMLConfiguration:                           "YAML Configuration",
	TypeZlibCompressed:                              "Zlib Compressed",
	TypeZshScript:                                   "Zsh Script",
}

func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
