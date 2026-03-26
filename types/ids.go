package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind1PasswordCloudKeychain
	Kind1PasswordKeychain
	Kind7ZipArchive
	KindAACAudio
	KindAccessDataFTK
	KindACEArchive
	KindAcrobatFormsDataFormat
	KindAcronisTrueImage
	KindAdobeFrameMaker
	KindAdobeInDesignDocument
	KindAdvancedForensicFormat
	KindAINArchive
	KindAllegroPackfile
	KindALZArchive
	KindAmigaDiskFile
	KindAmigaDiskMasher
	KindAmigaHardDiskFile
	KindAmigaHunkExecutable
	KindAmigaIcon
	KindAmigaOctaMED
	KindAMRAudio
	KindAMRWBAudio
	KindAnalogBox
	KindAndroidART
	KindAndroidBackup
	KindAndroidBinaryXML
	KindAndroidBootImage
	KindAndroidCompiledResources
	KindAndroidOAT
	KindAndroidODEX
	KindAndroidSparseImage
	KindAndroidVDEX
	KindAntennaData
	KindAOLARTImage
	KindAOLData
	KindAOLHistory
	KindAOLParameterInfo
	KindApacheArrowFile
	KindApacheParquet
	KindAPFSFilesystem
	KindAppImage
	KindAppleArchive
	KindAppleBillOfMaterials
	KindAppleBinaryCrashReport
	KindAppleBinaryPropertyList
	KindAppleDesktopServicesStore
	KindAppleDiskImage
	KindAppleDouble
	KindAppleISOHybridCDImage
	KindAppleiWorkDocument
	KindAppleKeychain
	KindAppleSingle
	KindAppleSystemLog
	KindAppleTexture
	KindAppleXMLPropertyList
	KindApproachIndex
	KindARArchive
	KindARCArchive
	KindARJArchive
	KindARRIRAWImage
	KindASFContainer
	KindASTCTexture
	KindAtari7800ROM
	KindAtariLynxROM
	KindAtomFeed
	KindAUAudio
	KindAudacityBlockFile
	KindAutoCADDrawing
	KindAutoCADDXF
	KindAVG6IntegrityDatabase
	KindAvroObjectContainer
	KindBAMData
	KindBase85
	KindBcachefsFilesystem
	KindBCFData
	KindBerkeleyDB
	KindBGBlitzDatabase
	KindBinHex
	KindBinkVideo
	KindBitcoinBlock
	KindBitcoinWallet
	KindBitLockerDiskEncryption
	KindBlackBerryBackup
	KindBlenderFile
	KindBlinkArchive
	KindBlizzardTexture
	KindBluetoothSnoop
	KindBMPImage
	KindBochsDiskImage
	KindBouncyCastleKeystore
	KindBPGImage
	KindBroadbandEBook
	KindBrotherEmbroidery
	KindBtrfsFilesystem
	KindBUFR
	KindBzip2Archive
	KindBzip3Archive
	KindCabinetArchive
	KindCAFAudio
	KindCalculuxProject
	KindCALSImage
	KindCanonRAWImage
	KindCBORData
	KindCDStomperLabel
	KindCerius2File
	KindCHMDocument
	KindChromaGraphBitmap
	KindChromiumPatch
	KindCinema4DModel
	KindCineonImage
	KindCloneCDControl
	KindComicBookArchive
	KindCommodore64Tape
	KindCommodore64TapeRAW
	KindCommodoreSID
	KindCompiledTerminfo
	KindCOMPlusCatalog
	KindCompressedSquareWave
	KindCorelBinaryMetafile
	KindCorelColorPalette
	KindCorelPhotoPaintImage
	KindCPIOArchive
	KindCRAMData
	KindCramFS
	KindCrushArchive
	KindCRXBrowserExtension
	KindCsoundMusic
	KindDalvikExecutable
	KindDAXArchive
	KindDB2Conversion
	KindDDSImage
	KindDebianPackage
	KindDeluxePaintAnimation
	KindDesignTools2D
	KindDeskMateDocument
	KindDeskMateWorksheet
	KindDialUpNetworking
	KindDICOMMedicalImage
	KindDigitalSpeechStandard
	KindDigitalWatchdogAudio
	KindDiracVideo
	KindDjVuDocument
	KindDolphinRVZ
	KindDOSExecutable
	KindDPXImage
	KindDreamcastAudio
	KindDSDIFFAudio
	KindDSFAudio
	KindDSTCompression
	KindDTSAudio
	KindDVDInfoFile
	KindDVRMSVideo
	KindDVRStudioStream
	KindEasyRecoverySavedState
	KindEasyStreetDraw
	KindEBMLContainer
	KindEFaxFile
	KindEGGArchive
	KindElitePlusCommander
	KindEncapsulatedPostScript
	KindEnCaseCaseFile
	KindEnCaseEvidenceV2
	KindEnCaseImage
	KindEndNoteLibrary
	KindEOTFont
	KindErlangBEAM
	KindESRIShapefile
	KindExecutableAndLinkableFormat
	KindExFATFilesystem
	KindExpertWitnessFormat
	KindExtensibleMusicFormat
	KindExtFilesystem
	KindF2FSFilesystem
	KindFamicomDiskSystemROM
	KindFarbfeldImage
	KindFastTracker2ExtendedInstrument
	KindFastTrackerModule
	KindFATFilesystem
	KindFBXModel
	KindFeatherData
	KindFiascoDatabase
	KindFilesystemEntry
	KindFirebirdDatabase
	KindFITActivity
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLIFImage
	KindFlightSimulatorConfig
	KindFLStudioProject
	KindFLVVideo
	KindFreeArcArchive
	KindFujifilmRAWImage
	KindFuzzyBitmap
	KindGameBoyAdvanceROM
	KindGameBoyROM
	KindGameBoySound
	KindGameCubeROM
	KindGDBMDatabase
	KindGEMRaster
	KindGenealogicalData
	KindGenetecVideo
	KindGHCiInterface
	KindGIFImage
	KindGIMPBrush
	KindGIMPPattern
	KindGIMPXCFImage
	KindGitBundle
	KindGitCommitGraph
	KindGitIndex
	KindGitPack
	KindGitPackBitmap
	KindGitPackIndex
	KindGitPackReverseIndex
	KindGlibcLocale
	KindGLTFBinary
	KindGlyphBitmapDistributionFormat
	KindGNOMEKeyring
	KindGNUGettextMachineCatalog
	KindGnuPGKeybox
	KindGodotPackage
	KindGoogleChromeDictionary
	KindGoogleDriveDrawing
	KindGoogleEarthETA
	KindGoogleEarthPlacemark
	KindGPGPublicKeyring
	KindGPSExchangeFormat
	KindGRIBData
	KindGUIDPartitionTable
	KindGzipArchive
	KindGzipData
	KindHadoopRCFile
	KindHadoopSequenceFile
	KindHamarsoftArchive
	KindHarvardGraphics
	KindHDF4Data
	KindHDF5Data
	KindHealthLevel7
	KindHFSPlusFilesystem
	KindHTMLDocument
	KindHuskygramPoem
	KindHusqvarnaDesigner
	KindICalendar
	KindICCProfile
	KindICNSIcon
	KindICOCURImage
	KindIEHistory
	KindIETPL
	KindIETrackingProtectionList
	KindIFFContainer
	KindImgSoftwareBitmap
	KindImpulseTrackerInstrument
	KindImpulseTrackerModule
	KindImpulseTrackerSample
	KindInnoSetupLog
	KindInstallShieldCAB
	KindInstallShieldScript
	KindISO9660Image
	KindISOBaseMedia
	KindIVFVideo
	KindJarArchive
	KindJARCSArchive
	KindJavaClass
	KindJavaCryptographyExtensionKeyStore
	KindJavaKeyStore
	KindJavaModule
	KindJavaPack200
	KindJavaSerialization
	KindJBIG2Image
	KindJeppesenFliteLog
	KindJFFS2Filesystem
	KindJNGImage
	KindJPEG2000Image
	KindJPEGImage
	KindJPEGLSImage
	KindJPEGXLImage
	KindJPEGXRImage
	KindJSONDocument
	KindKDEKWallet
	KindKeePassDatabase
	KindKeyboardDriver
	KindKeyholeMarkupLanguage
	KindKeyholeMarkupLanguageOverlay
	KindKGBArchive
	KindKorgAudio
	KindKTXTexture
	KindKyotoCabinetDatabase
	KindLevelDB
	KindLHAArchive
	KindLidarData
	KindLightWaveScene
	KindLinuxKernelImage
	KindLLVMBitcode
	KindLogicalFileEvidence
	KindLottieAnimation
	KindLotus123
	KindLotusAMIDocument
	KindLotusNotesDatabase
	KindLotusNotesDatabaseTemplate
	KindLotusWordPro
	KindLrzipArchive
	KindLuaBytecode
	KindLuceneIndex
	KindLUKSDiskEncryption
	KindLVM2PhysicalVolume
	KindLytroLightField
	KindLZ4Frame
	KindLZFArchive
	KindLZFSEData
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindLZXArchive
	KindM3U8Playlist
	KindMacBinary
	KindMachOBinary
	KindMachOUniversalBinary
	KindMacOSAlias
	KindMacOSXCodeSignature
	KindMacriumReflectImage
	KindMagicaVoxel
	KindMAMECHD
	KindMapInfoInterchange
	KindMapInfoSeaChart
	KindMArArchive
	KindMaterialExchangeFormat
	KindMATLABData
	KindMayaASCII
	KindMayaBinary
	KindMBOXEmailFolder
	KindMBOXTableOfContents
	KindMCAPCapture
	KindMeasurementDataFormat
	KindMediaDescriptor
	KindMerriamWebsterDictionary
	KindMetafileImage
	KindMicrografxDrawing
	KindMicrosoftAccessDatabase
	KindMicrosoftAgentCharacter
	KindMicrosoftAnswerWizard
	KindMicrosoftCodePageTranslation
	KindMicrosoftCompress
	KindMicrosoftCPlusPlusSymbols
	KindMicrosoftDeveloperStudioProject
	KindMicrosoftExchangeConfig
	KindMicrosoftFaxCover
	KindMicrosoftInfo
	KindMicrosoftJournal
	KindMicrosoftMoney
	KindMicrosoftNetworkMonitor
	KindMicrosoftOneNote
	KindMicrosoftOneNoteNote
	KindMicrosoftOutlookEmailFolder
	KindMicrosoftProgramDatabase
	KindMicrosoftReader
	KindMicrosoftSQLServerDatabase
	KindMicrosoftWinMobileNote
	KindMicrosoftWorksSpreadsheet
	KindMicrosoftWrite
	KindMicrosoftXbox360Profile
	KindMicrosoftXboxSave
	KindMIDISequence
	KindMilestonesProject
	KindMinoltaRAWImage
	KindMNGImage
	KindMOBIDocument
	KindMonkeysAudio
	KindMonochromePictureTIFF
	KindMoPaQArchive
	KindMozillaArchive
	KindMPEG2TransportStream
	KindMPEGAudio
	KindMPEGAudioFrame
	KindMPEGProgramStream
	KindMPEGTransportStream
	KindMPEGVideo
	KindMUGENSound
	KindMUGENSprite
	KindMultiBitBlockchain
	KindMultiBitInfo
	KindMultiBitWallet
	KindMusepackAudio
	KindMySQLMyISAM
	KindNationalImageryTransmission
	KindNationalTransferFormat
	KindNAVQuarantinedVirus
	KindNeoGeoPocketROM
	KindNeroCDCompilation
	KindNESROM
	KindNESSoundFormat
	KindNetCDFData
	KindNetpbmImage
	KindNetscapeMailFolder
	KindNIfTIMedicalImage
	KindNILFS2Filesystem
	KindNintendo3DSROM
	KindNintendo64ROM
	KindNintendoBCFNT
	KindNintendoBFLYT
	KindNintendoBRRES
	KindNintendoDSiROM
	KindNintendoDSROM
	KindNintendoRARC
	KindNintendoSwitchNRO
	KindNintendoSwitchNSO
	KindNintendoSwitchPackage
	KindNintendoSwitchROM
	KindNintendoTPL
	KindNintendoU8Archive
	KindNintendoWiiUExecutable
	KindNintendoYay0
	KindNintendoYaz0
	KindNISTSPHEREAudio
	KindNOAARasterChart
	KindNokiaPhoneBackup
	KindNortonDiskDoctorUndo
	KindNovellLANalyzerCapture
	KindNTFSFilesystem
	KindNullsoftVideo
	KindNumPyArray
	KindNVIDIASceneGraph
	KindOCamlObject
	KindOggContainer
	KindOLECompoundDocument
	KindOlympusRAWImage
	KindOnePasswordKeychain
	KindOpenEXRImage
	KindOpenSSHPrivateKey
	KindOpenTypeFont
	KindOptimFROGAudio
	KindORCColumnarData
	KindOutlookAddressBook
	KindOutlookAddressFile
	KindOutlookExpressAddressBook
	KindOutlookExpressDatabase
	KindPacketSnifferXCP
	KindPaintShopProImage
	KindPalmData
	KindPalmOSSuperMemo
	KindParallelsDiskImage
	KindParrotVideoEncapsulation
	KindPathWayMap
	KindPAXBitmap
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCFFont
	KindPCXDCXBitmap
	KindPCXImage
	KindPDFDocument
	KindPeaZipArchive
	KindPEMCertificate
	KindPEMPrivateKey
	KindPerfectOfficeDocument
	KindPestPatrolData
	KindPfaffEmbroidery
	KindPGFImage
	KindPGPDiskImage
	KindPGPMessage
	KindPGPPublicKeyring
	KindPGPSecretKeyring
	KindPhotoshopCustomShape
	KindPhotoshopDocument
	KindPicasaMovieProject
	KindPico8Cartridge
	KindPIMArchive
	KindPKCS12
	KindPKLITEArchive
	KindPKSFXArchive
	KindPlaylistFile
	KindPlayStation1Executable
	KindPlayStation2MemoryCard
	KindPlayStationExecutable
	KindPlayStationPortableExecutable
	KindPlayStationPortableISO
	KindPlayStationSoundFormat
	KindPLYModel
	KindPNGImage
	KindPortableExecutable
	KindPostgreSQLCustomDump
	KindPostScriptDocument
	KindPowerBASICDebugger
	KindPowerBuilderIDE
	KindPowerISODAA
	KindPRTGDatabase
	KindPSFFont
	KindPsiwtftabase
	KindPufferArchive
	KindPuttyPrivateKey
	KindPVRTexture
	KindPythonBytecode
	KindPythonPickle
	KindQCOWDiskImage
	KindQDBMDatabase
	KindQEMUQEDDiskImage
	KindQimageFilter
	KindQOIImage
	KindQuakePAK
	KindQuarkExpress
	KindQuattroPro
	KindQuickBooksBackup
	KindQuickenData
	KindQuickenQuickFinder
	KindQuickReport
	KindRadianceHDRImage
	KindRagTimeDocument
	KindRARArchive
	KindRData
	KindRealMedia
	KindRealMediaMetafile
	KindRealPlayerVideo
	KindRedisDatabase
	KindREDRAWImage
	KindReiserFSFilesystem
	KindRenPyArchive
	KindRhino3DModel
	KindRichTextFormatDocument
	KindRIFFContainer
	KindRKAudio
	KindRNCArchive
	KindROMFS
	KindROOTData
	KindROSBag
	KindRPGMakerArchive
	KindRPMPackage
	KindRSSFeed
	KindRubyGemPackage
	KindRubyMarshal
	KindRuntimeSoftwareImage
	KindRzipArchive
	KindSASData
	KindScreamTrackerModule
	KindSegaCDImage
	KindSegaDreamcastROM
	KindSegaMasterSystemROM
	KindSegaMegaDriveROM
	KindSegaNaomiROM
	KindSegaSaturnROM
	KindSGIImage
	KindShareazaThumbnail
	KindShebangScript
	KindShockwaveFlash
	KindShortenAudio
	KindShowPartnerGraphics
	KindSibeliusScore
	KindSietronicsDocument
	KindSigmaRAWImage
	KindSiliconGraphicsMovie
	KindSketchDocument
	KindSketchUpModel
	KindSkinCrafterSkin
	KindSkypeData
	KindSkypeLocalization
	KindSmackerVideo
	KindSmartDrawDrawing
	KindSnappyFramedData
	KindSNESSPC
	KindSnoopCapture
	KindSOAPMessage
	KindSonicFoundryAcid
	KindSonyCompressedVoice
	KindSonyOpenMG
	KindSonyWave64Audio
	KindSourceEngineBSP
	KindSpeedtouchFirmware
	KindSprintMusicStore
	KindSPSSData
	KindSPSSPortableData
	KindSPSSTemplate
	KindSQLite3SharedMemory
	KindSQLite3WriteAheadLog
	KindSQLiteDatabase
	KindSquashFSFilesystem
	KindSQXArchive
	KindSSHPublicKey
	KindStataData
	KindSteganosVirtualDrive
	KindSTEP3DModel
	KindSTLStereoLithography
	KindStorageCraftBackup
	KindStuffItArchive
	KindSubRipText
	KindSunRasterImage
	KindSuperCalcWorksheet
	KindSurfplanProject
	KindSVGImage
	KindSymantecGhostImage
	KindSymantecWiseInstallerLog
	KindSymbianExecutable
	KindSymbianInstallationFormat
	KindSystemdJournal
	KindSzipArchive
	KindTajimaEmbroidery
	KindTAKAudio
	KindTARArchive
	KindTargetExpress
	KindTensorFlowLiteModel
	KindTeraByteImage
	KindTextFile
	KindTheBatIndex
	KindThunderbirdMailSummary
	KindTIFFImage
	KindTimezoneData
	KindTokyoCabinetDatabase
	KindTomTomTraffic
	KindTorrentFile
	KindTransportNeutralEncapsulationFormat
	KindTrueTypeCollection
	KindTrueTypeFont
	KindTTAAudio
	KindU3DModel
	KindUBIFSFilesystem
	KindUBootImage
	KindUFAArchive
	KindUFOCaptureMap
	KindUHAArchive
	KindUndergroundAudio
	KindUnicodeExtensions
	KindUnityWebData
	KindUniversalDiskFormat
	KindUniversalSceneDescription
	KindUnixCompressArchive
	KindUnrealEnginePackage
	KindUUencodedArchive
	KindValvePak
	KindValveTextureFormat
	KindVBScript
	KindVBScriptEncoded
	KindVCard
	KindVeeamBackup
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVideoGameMusic
	KindVideoVCD
	KindVirtualBoxDiskImage
	KindVirtualBoxSavedState
	KindVisualBasicControl
	KindVisualCPreCompiledHeader
	KindVisualCWorkbenchInfo
	KindVisualStudioSolution
	KindVMapSourceWaypoint
	KindVMwareDiskImage
	KindVMwareNVRAM
	KindVMwareSnapshot
	KindVMwareSuspend
	KindVocaloidProject
	KindVocalTecMedia
	KindVOCAudio
	KindVRML3DModel
	KindVulkanSPIRV
	KindWADArchive
	KindWalkmanMP3
	KindWARCFile
	KindWarcraftIIIReplay
	KindWavPackAudio
	KindWavPackCorrection
	KindWebAssemblyModule
	KindWebexARF
	KindWebVTT
	KindWhereIsItCatalog
	KindWiiBackupFileSystem
	KindWiiGameCubeTHP
	KindWiiROM
	KindWiiUArchive
	KindWin95Password
	KindWin9xPassword
	KindWin9xPrinterSpool
	KindWin9xRegistryHive
	KindWindowsApplicationLog
	KindWindowsCalendar
	KindWindowsEventLog
	KindWindowsEventViewer
	KindWindowsHelp
	KindWindowsHibernation
	KindWindowsImagingFormat
	KindWindowsInternationalCodePage
	KindWindowsMemoryDump
	KindWindowsMinidump
	KindWindowsPrecompiledHeader
	KindWindowsPrefetch
	KindWindowsProgramManagerGroup
	KindWindowsRegistryHive
	KindWindowsResourceFile
	KindWindowsShortcut
	KindWindowsThumbnailCache
	KindWindowsTypeLibrary
	KindWindowsUserStateMigration
	KindWinNTNetmonCapture
	KindWinNTRegistryUndo
	KindWinPharoahCapture
	KindWinPharoahFilter
	KindWiredTigerDatabase
	KindWOFF2Font
	KindWOFFFont
	KindWord20Document
	KindWordPerfectDictionary
	KindWordPerfectGraphics
	KindWordPerfectText
	KindWordStarDocument
	KindWordStarWindows
	KindWTVVideo
	KindXARArchive
	KindXBMImage
	KindXbox360Executable
	KindXboxExecutable
	KindXboxISO
	KindXCOFFExecutable
	KindXFSFilesystem
	KindXilinxBitstream
	KindXMLDocument
	KindXPCOMLibraries
	KindXPMImage
	KindXZArchive
	KindYamahaSMAF
	KindYouTubeTimedText
	KindYUV4MPEG2Video
	KindZBackup
	KindZchunk
	KindZFSFilesystem
	KindZIPArchive
	KindZisofsArchive
	KindZlibData
	KindZOOArchive
	KindZoomBrowserIndex
	KindZPAQArchive
	KindZstandardArchive
	KindZstandardDictionary
	KindZXTape
)

const (
	TypeNone TypeID = iota
	Type32Bit
	Type32BitBigEndian
	Type32BitLittleEndian
	Type3G2Media
	Type3GPPMedia
	Type3MFDocument
	Type64Bit
	Type64BitBigEndian
	Type64BitLittleEndian
	TypeAC3
	TypeAcronisB4
	TypeAcronisCE
	TypeAdobeDNGDNG
	TypeADTS
	TypeAfterEffectsProjectAEP
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAndroidAppBundleAAB
	TypeAndroidArchiveAAR
	TypeAndroidPackageAPK
	TypeAndroidPackageXAPK
	TypeAndroidSplitAPKS
	TypeAndroidSystemPackageAPEX
	TypeAPPXPackage
	TypeArchLinuxPackage
	TypeASCIIText
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeB4
	TypeBashScript
	TypeBatch
	TypeBGZF
	TypeBigEndian
	TypeBigTIFF
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBlackmagicRAW
	TypeBlockDevice
	TypeByteSwapped
	TypeC
	TypeCanonRAW3CR3
	TypeCanonRAWHE
	TypeCAT
	TypeCB7
	TypeCBR
	TypeCBT
	TypeCBZ
	TypeCDAAudio
	TypeCE
	TypeCharacterDevice
	TypeCMake
	TypeCodestream
	TypeCwtfPackage
	TypeContainer
	TypeCorelDRAWDocumentCDR
	TypeCOWD
	TypeCPP
	TypeCRXVersion2
	TypeCRXVersion3
	TypeCSharp
	TypeCSS
	TypeCubaseProjectCPR
	TypeDart
	TypeDEX035
	TypeDEX036
	TypeDEX037
	TypeDEX038
	TypeDEX039
	TypeDEX040
	TypeDEX041
	TypeDirectory
	TypeDocker
	TypeDownloadableSounds
	TypeDSAPrivateKey
	TypeEAC3
	TypeECPrivateKey
	TypeELF
	TypeELF32
	TypeELF32BigEndian
	TypeELF32LittleEndian
	TypeELF64
	TypeELF64BigEndian
	TypeELF64LittleEndian
	TypeEmpty
	TypeEnhancedMetafileEMF
	TypeEPUBDocument
	TypeExt2
	TypeExt3
	TypeExt4
	TypeF4VVideo
	TypeFabricMod
	TypeFAT12
	TypeFAT16
	TypeFAT32
	TypeFirefoxExtensionXPI
	TypeFLACAudio
	TypeForgeMod
	TypeGIF87a
	TypeGIF89a
	TypeGo
	TypeGraphQL
	TypeHE
	TypeHEIFImage
	TypeILBMImage
	TypeIMMM
	TypeINI
	TypeIOSApplicationArchiveIPA
	TypeIWAD
	TypeJava
	TypeJavaArchiveJAR
	TypeJavaEnterpriseArchiveEAR
	TypeJavaScript
	TypeJavaWebArchiveWAR
	TypeJMOD
	TypeKDB
	TypeKDBX
	TypeKDM
	TypeKerasModel
	TypeKMZArchive
	TypeKotlin
	TypeKritaDocumentKRA
	TypeKTX
	TypeKTX2
	TypeLIST
	TypeLittleEndian
	TypeLN
	TypeLOVEGame
	TypeLua
	TypeLZ4Legacy
	TypeLZMACompressed
	TypeM2TS
	TypeM2TSBDAV
	TypeM4VVideo
	TypeMakefile
	TypeMarkdown
	TypeMatroska
	TypeMicrosoftExcelAddInXLAM
	TypeMicrosoftExcelMacroEnabledTemplateXLTM
	TypeMicrosoftExcelMacroEnabledWorkbookXLSM
	TypeMicrosoftExcelSpreadsheetXLSX
	TypeMicrosoftExcelTemplateXLTX
	TypeMicrosoftExcelWorkbookXLS
	TypeMicrosoftInstallerMSI
	TypeMicrosoftOutlookMessageMSG
	TypeMicrosoftPowerPointAddInPPAM
	TypeMicrosoftPowerPointMacroEnabledPresentationPPTM
	TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM
	TypeMicrosoftPowerPointMacroEnabledTemplatePOTM
	TypeMicrosoftPowerPointPresentationPPT
	TypeMicrosoftPowerPointPresentationPPTX
	TypeMicrosoftPowerPointSlideshowPPSX
	TypeMicrosoftPowerPointTemplatePOTX
	TypeMicrosoftProjectDocumentMPP
	TypeMicrosoftPublisherDocumentPUB
	TypeMicrosoftVisioDrawingVSD
	TypeMicrosoftWordDocumentDOC
	TypeMicrosoftWordDocumentDOCX
	TypeMicrosoftWordMacroEnabledDocumentDOCM
	TypeMicrosoftWordMacroEnabledTemplateDOTM
	TypeMicrosoftWordTemplateDOTX
	TypeMinecraftResourcePack
	TypeMP3
	TypeMP3ID3Tagged
	TypeMP4Video
	TypeMPEG12Video
	TypeMPEG4AudioM4AFamily
	TypeMPEGLayer2
	TypeMPEGLayer3
	TypeMSIXPackage
	TypeNamedPipe
	TypeNanosecondBigEndian
	TypeNanosecondLittleEndian
	TypeNB
	TypeNewASCII
	TypeNewASCIIWithCRC
	TypeNikonRAWNEF
	TypeNpmPackageTarball
	TypeNuGetPackageNUPKG
	TypeObjectiveC
	TypeOCIImageLayoutTar
	TypeOldASCII
	TypeOpenDocumentChartODC
	TypeOpenDocumentDatabaseODB
	TypeOpenDocumentFormulaODF
	TypeOpenDocumentGraphicsODG
	TypeOpenDocumentImageODI
	TypeOpenDocumentPresentationODP
	TypeOpenDocumentSpreadsheetODS
	TypeOpenDocumentTextODT
	TypeOpenRasterImageORA
	TypeOpusAudio
	TypePAM
	TypePanasonicRAWRW2
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
	TypePentaxRAWPEF
	TypePerl
	TypePES
	TypePGMASCII
	TypePGMBinary
	TypePHP
	TypePowerShell
	TypePPMASCII
	TypePPMBinary
	TypePSB
	TypePSD
	TypePWAD
	TypePython
	TypePythonSourceDistributionSDist
	TypePythonWheelWHL
	TypePyTorchModel
	TypeQCOW1
	TypeQCOW2
	TypeQCPAudio
	TypeQuickTimeMovie
	TypeR
	TypeRAR4
	TypeRAR5
	TypeReturnPath
	TypeRIFFMIDI
	TypeRIFFPalette
	TypeRSAPrivateKey
	TypeRuby
	TypeRust
	TypeScala
	TypeSketchDocument
	TypeSlackwarePackage
	TypeSocket
	TypeSonyRAWARW
	TypeSonyRAWSR2
	TypeSoundFont2
	TypeSpanned
	TypeSpecial
	TypeSpeexAudio
	TypeSQL
	TypeStreamVersion7
	TypeStreamVersion8
	TypeSwift
	TypeSymbolicLink
	TypeSZ
	TypeTerraform
	TypeTheoraVideo
	TypeTOML
	TypeTrue
	TypeTS
	TypeTypeScript
	TypeUncompressed
	TypeUTF8Text
	TypeVagrantBox
	TypeVisualStudioExtensionVSIX
	TypeVMDK
	TypeVMDKDescription
	TypeVorbisAudio
	TypeWAVAudio
	TypeWebM
	TypeWebPImage
	TypeWindowsAnimatedCursor
	TypeWindowsCursor
	TypeWindowsIcon
	TypeWindowsLE
	TypeWindowsLX
	TypeWindowsMetafileWMF
	TypeWindowsNE
	TypeWrapper
	TypeXMLPaperSpecificationXPS
	TypeYAML
	TypeZlibCompressed
	TypeZstandardSkinnableFrame
)

var kindNames = [...]string{
	KindUnknown:                             "Unknown",
	Kind1PasswordCloudKeychain:              "1Password Cloud Keychain",
	Kind1PasswordKeychain:                   "1Password Keychain",
	Kind7ZipArchive:                         "7-Zip Archive",
	KindAACAudio:                            "AAC Audio",
	KindAccessDataFTK:                       "AccessData FTK Evidence",
	KindACEArchive:                          "ACE Archive",
	KindAcrobatFormsDataFormat:              "Acrobat Forms Data Format",
	KindAcronisTrueImage:                    "Acronis True Image",
	KindAdobeFrameMaker:                     "Adobe FrameMaker",
	KindAdobeInDesignDocument:               "Adobe InDesign Document",
	KindAdvancedForensicFormat:              "Advanced Forensic Format",
	KindAINArchive:                          "AIN Archive",
	KindAllegroPackfile:                     "Allegro Packfile",
	KindALZArchive:                          "ALZ Archive",
	KindAmigaDiskFile:                       "Amiga Disk File",
	KindAmigaDiskMasher:                     "Amiga DiskMasher",
	KindAmigaHardDiskFile:                   "Amiga Hard Disk File",
	KindAmigaHunkExecutable:                 "Amiga Hunk Executable",
	KindAmigaIcon:                           "Amiga Icon",
	KindAmigaOctaMED:                        "Amiga OctaMED Audio",
	KindAMRAudio:                            "AMR Audio",
	KindAMRWBAudio:                          "AMR-WB Audio",
	KindAnalogBox:                           "Analog Box Circuit",
	KindAndroidART:                          "Android ART",
	KindAndroidBackup:                       "Android Backup",
	KindAndroidBinaryXML:                    "Android Binary XML",
	KindAndroidBootImage:                    "Android Boot Image",
	KindAndroidCompiledResources:            "Android Compiled Resources",
	KindAndroidOAT:                          "Android OAT",
	KindAndroidODEX:                         "Android ODEX",
	KindAndroidSparseImage:                  "Android Sparse Image",
	KindAndroidVDEX:                         "Android VDEX",
	KindAntennaData:                         "Antenna Data",
	KindAOLARTImage:                         "AOL ART Image",
	KindAOLData:                             "AOL Data",
	KindAOLHistory:                          "AOL History",
	KindAOLParameterInfo:                    "AOL Parameter Info",
	KindApacheArrowFile:                     "Apache Arrow File",
	KindApacheParquet:                       "Apache Parquet",
	KindAPFSFilesystem:                      "APFS Filesystem",
	KindAppImage:                            "AppImage",
	KindAppleArchive:                        "Apple Archive",
	KindAppleBillOfMaterials:                "Apple Bill of Materials",
	KindAppleBinaryCrashReport:              "Apple Binary Crash Report",
	KindAppleBinaryPropertyList:             "Apple Binary Property List",
	KindAppleDesktopServicesStore:           "Apple Desktop Services Store",
	KindAppleDiskImage:                      "Apple Disk Image",
	KindAppleDouble:                         "AppleDouble File",
	KindAppleISOHybridCDImage:               "Apple ISO Hybrid CD Image",
	KindAppleiWorkDocument:                  "Apple iWork Document",
	KindAppleKeychain:                       "Apple Keychain",
	KindAppleSingle:                         "AppleSingle File",
	KindAppleSystemLog:                      "Apple System Log",
	KindAppleTexture:                        "Apple Texture",
	KindAppleXMLPropertyList:                "Apple XML Property List",
	KindApproachIndex:                       "Approach Index",
	KindARArchive:                           "AR Archive",
	KindARCArchive:                          "ARC Archive",
	KindARJArchive:                          "ARJ Archive",
	KindARRIRAWImage:                        "ARRI RAW Image",
	KindASFContainer:                        "ASF Container",
	KindASTCTexture:                         "ASTC Texture",
	KindAtari7800ROM:                        "Atari 7800 ROM",
	KindAtariLynxROM:                        "Atari Lynx ROM",
	KindAtomFeed:                            "Atom Feed",
	KindAUAudio:                             "AU Audio",
	KindAudacityBlockFile:                   "Audacity Block File",
	KindAutoCADDrawing:                      "AutoCAD Drawing",
	KindAutoCADDXF:                          "AutoCAD DXF",
	KindAVG6IntegrityDatabase:               "AVG6 Integrity Database",
	KindAvroObjectContainer:                 "Avro Object Container",
	KindBAMData:                             "BAM Data",
	KindBase85:                              "BASE85 File",
	KindBcachefsFilesystem:                  "Bcachefs Filesystem",
	KindBCFData:                             "Binary Call Format (BCF)",
	KindBerkeleyDB:                          "Berkeley DB",
	KindBGBlitzDatabase:                     "BGBlitz Database",
	KindBinHex:                              "BinHex Archive",
	KindBinkVideo:                           "Bink Video",
	KindBitcoinBlock:                        "Bitcoin Block Data",
	KindBitcoinWallet:                       "Bitcoin Wallet",
	KindBitLockerDiskEncryption:             "BitLocker Disk Encryption",
	KindBlackBerryBackup:                    "BlackBerry Backup",
	KindBlenderFile:                         "Blender File",
	KindBlinkArchive:                        "Blink Archive",
	KindBlizzardTexture:                     "Blizzard Texture",
	KindBluetoothSnoop:                      "Bluetooth Snoop Capture",
	KindBMPImage:                            "BMP Image",
	KindBochsDiskImage:                      "Bochs Disk Image",
	KindBouncyCastleKeystore:                "BouncyCastle Keystore",
	KindBPGImage:                            "BPG Image",
	KindBroadbandEBook:                      "Broadband eBook",
	KindBrotherEmbroidery:                   "Brother Embroidery",
	KindBtrfsFilesystem:                     "Btrfs Filesystem",
	KindBUFR:                                "BUFR Data",
	KindBzip2Archive:                        "Bzip2 Archive",
	KindBzip3Archive:                        "Bzip3 Archive",
	KindCabinetArchive:                      "Cabinet Archive",
	KindCAFAudio:                            "CAF Audio",
	KindCalculuxProject:                     "Calculux Project",
	KindCALSImage:                           "CALS Raster Image",
	KindCanonRAWImage:                       "Canon RAW Image",
	KindCBORData:                            "CBOR Data",
	KindCDStomperLabel:                      "CD Stomper Label",
	KindCerius2File:                         "Cerius2 File",
	KindCHMDocument:                         "CHM Document",
	KindChromaGraphBitmap:                   "ChromaGraph Bitmap",
	KindChromiumPatch:                       "Chromium Patch",
	KindCinema4DModel:                       "Cinema 4D Model",
	KindCineonImage:                         "Cineon Image",
	KindCloneCDControl:                      "CloneCD Control File",
	KindComicBookArchive:                    "Comic Book Archive",
	KindCommodore64Tape:                     "Commodore 64 Tape",
	KindCommodore64TapeRAW:                  "Commodore 64 Tape RAW",
	KindCommodoreSID:                        "Commodore SID Audio",
	KindCompiledTerminfo:                    "Compiled Terminfo",
	KindCOMPlusCatalog:                      "COM+ Catalog",
	KindCompressedSquareWave:                "Compressed Square Wave",
	KindCorelBinaryMetafile:                 "Corel Binary Metafile",
	KindCorelColorPalette:                   "Corel Color Palette",
	KindCorelPhotoPaintImage:                "Corel Photo-Paint Image",
	KindCPIOArchive:                         "CPIO Archive",
	KindCRAMData:                            "CRAM Data",
	KindCramFS:                              "CramFS",
	KindCrushArchive:                        "Crush Archive",
	KindCRXBrowserExtension:                 "CRX Browser Extension",
	KindCsoundMusic:                         "Csound Music",
	KindDalvikExecutable:                    "Dalvik Executable",
	KindDAXArchive:                          "DAX Archive",
	KindDB2Conversion:                       "DB2 Conversion",
	KindDDSImage:                            "DDS Image",
	KindDebianPackage:                       "Debian Package",
	KindDeluxePaintAnimation:                "DeluxePaint Animation",
	KindDesignTools2D:                       "DesignTools 2D",
	KindDeskMateDocument:                    "DeskMate Document",
	KindDeskMateWorksheet:                   "DeskMate Worksheet",
	KindDialUpNetworking:                    "Dial-up Networking",
	KindDICOMMedicalImage:                   "DICOM Medical Image",
	KindDigitalSpeechStandard:               "Digital Speech Standard",
	KindDigitalWatchdogAudio:                "Digital Watchdog Audio",
	KindDiracVideo:                          "Dirac Video",
	KindDjVuDocument:                        "DjVu Document",
	KindDolphinRVZ:                          "Dolphin RVZ",
	KindDOSExecutable:                       "DOS Executable",
	KindDPXImage:                            "DPX Image",
	KindDreamcastAudio:                      "Dreamcast Audio",
	KindDSDIFFAudio:                         "DSDIFF Audio",
	KindDSFAudio:                            "DSF Audio",
	KindDSTCompression:                      "DST Compression",
	KindDTSAudio:                            "DTS Audio",
	KindDVDInfoFile:                         "DVD Info File",
	KindDVRMSVideo:                          "DVR-MS Video",
	KindDVRStudioStream:                     "DVR-Studio Stream",
	KindEasyRecoverySavedState:              "EasyRecovery Saved State",
	KindEasyStreetDraw:                      "Easy Street Draw",
	KindEBMLContainer:                       "EBML Container",
	KindEFaxFile:                            "eFax File",
	KindEGGArchive:                          "EGG Archive",
	KindElitePlusCommander:                  "Elite Plus Commander",
	KindEncapsulatedPostScript:              "Encapsulated PostScript",
	KindEnCaseCaseFile:                      "EnCase Case File",
	KindEnCaseEvidenceV2:                    "EnCase Evidence V2",
	KindEnCaseImage:                         "EnCase Image",
	KindEndNoteLibrary:                      "EndNote Library",
	KindEOTFont:                             "EOT Font",
	KindErlangBEAM:                          "Erlang BEAM Bytecode",
	KindESRIShapefile:                       "ESRI Shapefile",
	KindExecutableAndLinkableFormat:         "Executable and Linkable Format",
	KindExFATFilesystem:                     "exFAT Filesystem",
	KindExpertWitnessFormat:                 "Expert Witness Format",
	KindExtensibleMusicFormat:               "Extensible Music Format",
	KindExtFilesystem:                       "ext Filesystem",
	KindF2FSFilesystem:                      "F2FS Filesystem",
	KindFamicomDiskSystemROM:                "Famicom Disk System ROM",
	KindFarbfeldImage:                       "Farbfeld Image",
	KindFastTracker2ExtendedInstrument:      "FastTracker 2 Extended Instrument",
	KindFastTrackerModule:                   "FastTracker Module",
	KindFATFilesystem:                       "FAT Filesystem",
	KindFBXModel:                            "FBX Model",
	KindFeatherData:                         "Feather Data",
	KindFiascoDatabase:                      "Fiasco Database",
	KindFilesystemEntry:                     "Filesystem Entry",
	KindFirebirdDatabase:                    "Firebird Database",
	KindFITActivity:                         "FIT Activity",
	KindFITSAstronomicalImage:               "FITS Astronomical Image",
	KindFLACAudio:                           "FLAC Audio",
	KindFLIFImage:                           "FLIF Image",
	KindFlightSimulatorConfig:               "Flight Simulator Config",
	KindFLStudioProject:                     "FL Studio Project",
	KindFLVVideo:                            "FLV Video",
	KindFreeArcArchive:                      "FreeArc Archive",
	KindFujifilmRAWImage:                    "Fujifilm RAW Image",
	KindFuzzyBitmap:                         "Fuzzy Bitmap",
	KindGameBoyAdvanceROM:                   "Game Boy Advance ROM",
	KindGameBoyROM:                          "Game Boy ROM",
	KindGameBoySound:                        "Game Boy Sound",
	KindGameCubeROM:                         "GameCube ROM",
	KindGDBMDatabase:                        "GDBM Database",
	KindGEMRaster:                           "GEM Raster",
	KindGenealogicalData:                    "Genealogical Data",
	KindGenetecVideo:                        "Genetec Video",
	KindGHCiInterface:                       "GHCi Interface",
	KindGIFImage:                            "GIF Image",
	KindGIMPBrush:                           "GIMP Brush",
	KindGIMPPattern:                         "GIMP Pattern",
	KindGIMPXCFImage:                        "GIMP XCF Image",
	KindGitBundle:                           "Git Bundle",
	KindGitCommitGraph:                      "Git Commit Graph",
	KindGitIndex:                            "Git Index",
	KindGitPack:                             "Git Pack",
	KindGitPackBitmap:                       "Git Pack Bitmap",
	KindGitPackIndex:                        "Git Pack Index",
	KindGitPackReverseIndex:                 "Git Pack Reverse Index",
	KindGlibcLocale:                         "Glibc Locale",
	KindGLTFBinary:                          "glTF Binary",
	KindGlyphBitmapDistributionFormat:       "Glyph Bitmap Distribution Format (BDF)",
	KindGNOMEKeyring:                        "GNOME Keyring",
	KindGNUGettextMachineCatalog:            "GNU Gettext Machine Catalog",
	KindGnuPGKeybox:                         "GnuPG Keybox",
	KindGodotPackage:                        "Godot Engine Package",
	KindGoogleChromeDictionary:              "Google Chrome Dictionary",
	KindGoogleDriveDrawing:                  "Google Drive Drawing",
	KindGoogleEarthETA:                      "Google Earth ETA",
	KindGoogleEarthPlacemark:                "Google Earth Placemark",
	KindGPGPublicKeyring:                    "GPG Public Keyring",
	KindGPSExchangeFormat:                   "GPS Exchange Format",
	KindGRIBData:                            "GRIB Data",
	KindGUIDPartitionTable:                  "GUID Partition Table",
	KindGzipArchive:                         "Gzip Archive",
	KindGzipData:                            "Gzip Data",
	KindHadoopRCFile:                        "Hadoop RCFile",
	KindHadoopSequenceFile:                  "Hadoop SequenceFile",
	KindHamarsoftArchive:                    "Hamarsoft Archive",
	KindHarvardGraphics:                     "Harvard Graphics",
	KindHDF4Data:                            "HDF4 Data",
	KindHDF5Data:                            "HDF5 Data",
	KindHealthLevel7:                        "Health Level-7 Data",
	KindHFSPlusFilesystem:                   "HFS+ Filesystem",
	KindHTMLDocument:                        "HTML Document",
	KindHuskygramPoem:                       "Huskygram Poem",
	KindHusqvarnaDesigner:                   "Husqvarna Designer",
	KindICalendar:                           "iCalendar",
	KindICCProfile:                          "ICC Profile",
	KindICNSIcon:                            "ICNS Icon",
	KindICOCURImage:                         "ICO/CUR Image",
	KindIEHistory:                           "IE History",
	KindIETPL:                               "IE TPL",
	KindIETrackingProtectionList:            "IE Tracking Protection List",
	KindIFFContainer:                        "IFF Container",
	KindImgSoftwareBitmap:                   "Img Software Bitmap",
	KindImpulseTrackerInstrument:            "Impulse Tracker Instrument",
	KindImpulseTrackerModule:                "Impulse Tracker Module",
	KindImpulseTrackerSample:                "Impulse Tracker Sample",
	KindInnoSetupLog:                        "Inno Setup Log",
	KindInstallShieldCAB:                    "InstallShield Cabinet",
	KindInstallShieldScript:                 "InstallShield Script",
	KindISO9660Image:                        "ISO 9660 Image",
	KindISOBaseMedia:                        "ISO Base Media",
	KindIVFVideo:                            "IVF Video",
	KindJarArchive:                          "JAR Archive",
	KindJARCSArchive:                        "JARCS Archive",
	KindJavaClass:                           "Java Class",
	KindJavaCryptographyExtensionKeyStore:   "JCE KeyStore",
	KindJavaKeyStore:                        "Java KeyStore",
	KindJavaModule:                          "Java Module",
	KindJavaPack200:                         "Java Pack200 Archive",
	KindJavaSerialization:                   "Java Serialization Data",
	KindJBIG2Image:                          "JBIG2 Image",
	KindJeppesenFliteLog:                    "Jeppesen FliteLog",
	KindJFFS2Filesystem:                     "JFFS2 Filesystem",
	KindJNGImage:                            "JNG Image",
	KindJPEG2000Image:                       "JPEG 2000 Image",
	KindJPEGImage:                           "JPEG Image",
	KindJPEGLSImage:                         "JPEG-LS Image",
	KindJPEGXLImage:                         "JPEG XL Image",
	KindJPEGXRImage:                         "JPEG XR Image",
	KindJSONDocument:                        "JSON Document",
	KindKDEKWallet:                          "KDE KWallet",
	KindKeePassDatabase:                     "KeePass Database",
	KindKeyboardDriver:                      "Keyboard Driver",
	KindKeyholeMarkupLanguage:               "Keyhole Markup Language",
	KindKeyholeMarkupLanguageOverlay:        "Keyhole Markup Language Overlay",
	KindKGBArchive:                          "KGB Archive",
	KindKorgAudio:                           "Korg Audio",
	KindKTXTexture:                          "KTX Texture",
	KindKyotoCabinetDatabase:                "Kyoto Cabinet Database",
	KindLevelDB:                             "LevelDB/RocksDB",
	KindLHAArchive:                          "LHA Archive",
	KindLidarData:                           "LiDAR Data (LAS)",
	KindLightWaveScene:                      "LightWave Scene",
	KindLinuxKernelImage:                    "Linux Kernel Image",
	KindLLVMBitcode:                         "LLVM Bitcode",
	KindLogicalFileEvidence:                 "Logical File Evidence",
	KindLottieAnimation:                     "Lottie Animation",
	KindLotus123:                            "Lotus 1-2-3",
	KindLotusAMIDocument:                    "Lotus AMI Document",
	KindLotusNotesDatabase:                  "Lotus Notes Database",
	KindLotusNotesDatabaseTemplate:          "Lotus Notes Database Template",
	KindLotusWordPro:                        "Lotus WordPro",
	KindLrzipArchive:                        "Long Range ZIP Archive",
	KindLuaBytecode:                         "Lua Bytecode",
	KindLuceneIndex:                         "Lucene Index",
	KindLUKSDiskEncryption:                  "LUKS Disk Encryption",
	KindLVM2PhysicalVolume:                  "LVM2 Physical Volume",
	KindLytroLightField:                     "Lytro Light Field",
	KindLZ4Frame:                            "LZ4 Frame",
	KindLZFArchive:                          "LZF Archive",
	KindLZFSEData:                           "LZFSE Data",
	KindLZIPArchive:                         "LZIP Archive",
	KindLZMAData:                            "LZMA Data",
	KindLZOPArchive:                         "LZOP Archive",
	KindLZXArchive:                          "LZX Archive",
	KindM3U8Playlist:                        "M3U8 Playlist",
	KindMacBinary:                           "MacBinary",
	KindMachOBinary:                         "Mach-O Binary",
	KindMachOUniversalBinary:                "Mach-O Universal Binary",
	KindMacOSAlias:                          "macOS Alias",
	KindMacOSXCodeSignature:                 "macOS Code Signature",
	KindMacriumReflectImage:                 "Macrium Reflect Image",
	KindMagicaVoxel:                         "MagicaVoxel Model",
	KindMAMECHD:                             "MAME Compressed Hunks of Data",
	KindMapInfoInterchange:                  "MapInfo Interchange",
	KindMapInfoSeaChart:                     "MapInfo Sea Chart",
	KindMArArchive:                          "MAr Archive",
	KindMaterialExchangeFormat:              "Material Exchange Format",
	KindMATLABData:                          "MATLAB Data",
	KindMayaASCII:                           "Maya ASCII",
	KindMayaBinary:                          "Maya Binary",
	KindMBOXEmailFolder:                     "MBOX Email Folder",
	KindMBOXTableOfContents:                 "MBOX Table of Contents",
	KindMCAPCapture:                         "MCAP Capture",
	KindMeasurementDataFormat:               "Measurement Data Format",
	KindMediaDescriptor:                     "Media Descriptor",
	KindMerriamWebsterDictionary:            "Merriam-Webster Dictionary",
	KindMetafileImage:                       "Metafile Image",
	KindMicrografxDrawing:                   "Micrografx Drawing",
	KindMicrosoftAccessDatabase:             "Microsoft Access Database",
	KindMicrosoftAgentCharacter:             "Microsoft Agent Character",
	KindMicrosoftAnswerWizard:               "Microsoft Answer Wizard",
	KindMicrosoftCodePageTranslation:        "Microsoft Code Page Translation",
	KindMicrosoftCompress:                   "Microsoft Compress Archive",
	KindMicrosoftCPlusPlusSymbols:           "Microsoft C++ Symbols",
	KindMicrosoftDeveloperStudioProject:     "Microsoft Developer Studio Project",
	KindMicrosoftExchangeConfig:             "Microsoft Exchange Config",
	KindMicrosoftFaxCover:                   "Microsoft Fax Cover",
	KindMicrosoftInfo:                       "Microsoft Info",
	KindMicrosoftJournal:                    "Microsoft Journal",
	KindMicrosoftMoney:                      "Microsoft Money",
	KindMicrosoftNetworkMonitor:             "Microsoft Network Monitor Capture",
	KindMicrosoftOneNote:                    "Microsoft OneNote",
	KindMicrosoftOneNoteNote:                "Microsoft OneNote Note",
	KindMicrosoftOutlookEmailFolder:         "Microsoft Outlook Email Folder",
	KindMicrosoftProgramDatabase:            "Microsoft Program Database (PDB)",
	KindMicrosoftReader:                     "Microsoft Reader eBook",
	KindMicrosoftSQLServerDatabase:          "Microsoft SQL Server Database",
	KindMicrosoftWinMobileNote:              "Microsoft WinMobile Note",
	KindMicrosoftWorksSpreadsheet:           "Microsoft Works Spreadsheet",
	KindMicrosoftWrite:                      "Microsoft Write",
	KindMicrosoftXbox360Profile:             "Xbox 360 Profile",
	KindMicrosoftXboxSave:                   "Xbox Save",
	KindMIDISequence:                        "MIDI Sequence",
	KindMilestonesProject:                   "Milestones Project",
	KindMinoltaRAWImage:                     "Minolta RAW Image",
	KindMNGImage:                            "MNG Image",
	KindMOBIDocument:                        "MOBI Document",
	KindMonkeysAudio:                        "Monkey's Audio",
	KindMonochromePictureTIFF:               "Monochrome Picture TIFF",
	KindMoPaQArchive:                        "MoPaQ Archive",
	KindMozillaArchive:                      "Mozilla Archive (MAR)",
	KindMPEG2TransportStream:                "MPEG Transport Stream",
	KindMPEGAudio:                           "MPEG Audio",
	KindMPEGAudioFrame:                      "MPEG Audio",
	KindMPEGProgramStream:                   "MPEG Program Stream",
	KindMPEGTransportStream:                 "MPEG Transport Stream",
	KindMPEGVideo:                           "MPEG Video",
	KindMUGENSound:                          "M.U.G.E.N Sound",
	KindMUGENSprite:                         "M.U.G.E.N Sprite",
	KindMultiBitBlockchain:                  "MultiBit Blockchain",
	KindMultiBitInfo:                        "MultiBit Info",
	KindMultiBitWallet:                      "MultiBit Wallet",
	KindMusepackAudio:                       "Musepack Audio",
	KindMySQLMyISAM:                         "MySQL MyISAM",
	KindNationalImageryTransmission:         "National Imagery Transmission",
	KindNationalTransferFormat:              "National Transfer Format",
	KindNAVQuarantinedVirus:                 "NAV Quarantined Virus",
	KindNeoGeoPocketROM:                     "Neo Geo Pocket ROM",
	KindNeroCDCompilation:                   "Nero CD Compilation",
	KindNESROM:                              "NES ROM",
	KindNESSoundFormat:                      "NES Sound Format",
	KindNetCDFData:                          "NetCDF Data",
	KindNetpbmImage:                         "Netpbm Image",
	KindNetscapeMailFolder:                  "Netscape Mail Folder",
	KindNIfTIMedicalImage:                   "NIfTI Medical Image",
	KindNILFS2Filesystem:                    "NILFS2 Filesystem",
	KindNintendo3DSROM:                      "Nintendo 3DS ROM",
	KindNintendo64ROM:                       "Nintendo 64 ROM",
	KindNintendoBCFNT:                       "Nintendo BCFNT Font",
	KindNintendoBFLYT:                       "Nintendo BFLYT Layout",
	KindNintendoBRRES:                       "Nintendo BRRES Resource",
	KindNintendoDSiROM:                      "Nintendo DSi ROM",
	KindNintendoDSROM:                       "Nintendo DS ROM",
	KindNintendoRARC:                        "Nintendo RARC Archive",
	KindNintendoSwitchNRO:                   "Nintendo Switch NRO",
	KindNintendoSwitchNSO:                   "Nintendo Switch NSO",
	KindNintendoSwitchPackage:               "Nintendo Switch Package",
	KindNintendoSwitchROM:                   "Nintendo Switch ROM",
	KindNintendoTPL:                         "Nintendo TPL Image",
	KindNintendoU8Archive:                   "Nintendo U8 Archive",
	KindNintendoWiiUExecutable:              "Nintendo Wii U Executable",
	KindNintendoYay0:                        "Nintendo Yay0 Compressed",
	KindNintendoYaz0:                        "Nintendo Yaz0 Compressed",
	KindNISTSPHEREAudio:                     "NIST SPHERE Audio",
	KindNOAARasterChart:                     "NOAA Raster Chart",
	KindNokiaPhoneBackup:                    "Nokia Phone Backup",
	KindNortonDiskDoctorUndo:                "Norton Disk Doctor Undo",
	KindNovellLANalyzerCapture:              "Novell LANalyzer Capture",
	KindNTFSFilesystem:                      "NTFS Filesystem",
	KindNullsoftVideo:                       "Nullsoft Video",
	KindNumPyArray:                          "NumPy Array",
	KindNVIDIASceneGraph:                    "NVIDIA Scene Graph",
	KindOCamlObject:                         "OCaml Object",
	KindOggContainer:                        "Ogg Container",
	KindOLECompoundDocument:                 "OLE Compound Document",
	KindOlympusRAWImage:                     "Olympus RAW Image",
	KindOnePasswordKeychain:                 "1Password Keychain",
	KindOpenEXRImage:                        "OpenEXR Image",
	KindOpenSSHPrivateKey:                   "OpenSSH Private Key",
	KindOpenTypeFont:                        "OpenType Font",
	KindOptimFROGAudio:                      "OptimFROG Audio",
	KindORCColumnarData:                     "ORC Columnar Data",
	KindOutlookAddressBook:                  "Outlook Address Book",
	KindOutlookAddressFile:                  "Outlook Address File",
	KindOutlookExpressAddressBook:           "Outlook Express Address Book",
	KindOutlookExpressDatabase:              "Outlook Express Database",
	KindPacketSnifferXCP:                    "Packet Sniffer XCP",
	KindPaintShopProImage:                   "PaintShop Pro Image",
	KindPalmData:                            "Palm Data",
	KindPalmOSSuperMemo:                     "PalmOS SuperMemo",
	KindParallelsDiskImage:                  "Parallels Disk Image",
	KindParrotVideoEncapsulation:            "Parrot Video Encapsulation",
	KindPathWayMap:                          "PathWay Map",
	KindPAXBitmap:                           "PAX Bitmap",
	KindPCAPCapture:                         "PCAP Capture",
	KindPCAPNGCapture:                       "PCAPNG Capture",
	KindPCFFont:                             "Portable Compiled Format Font",
	KindPCXDCXBitmap:                        "PCX/DCX Bitmap",
	KindPCXImage:                            "PCX Image",
	KindPDFDocument:                         "PDF Document",
	KindPeaZipArchive:                       "PeaZip Archive",
	KindPEMCertificate:                      "PEM Certificate",
	KindPEMPrivateKey:                       "PEM Private Key",
	KindPerfectOfficeDocument:               "Perfect Office Document",
	KindPestPatrolData:                      "PestPatrol Data",
	KindPfaffEmbroidery:                     "Pfaff Embroidery",
	KindPGFImage:                            "Progressive Graphics File",
	KindPGPDiskImage:                        "PGP Disk Image",
	KindPGPMessage:                          "PGP Message",
	KindPGPPublicKeyring:                    "PGP Public Keyring",
	KindPGPSecretKeyring:                    "PGP Secret Keyring",
	KindPhotoshopCustomShape:                "Photoshop Custom Shape",
	KindPhotoshopDocument:                   "Photoshop Document",
	KindPicasaMovieProject:                  "Picasa Movie Project",
	KindPico8Cartridge:                      "Pico-8 Cartridge",
	KindPIMArchive:                          "PIM Archive",
	KindPKCS12:                              "PKCS#12",
	KindPKLITEArchive:                       "PKLITE Archive",
	KindPKSFXArchive:                        "PKSFX Archive",
	KindPlaylistFile:                        "Playlist File",
	KindPlayStation1Executable:              "PlayStation 1 Executable",
	KindPlayStation2MemoryCard:              "PlayStation 2 Memory Card",
	KindPlayStationExecutable:               "PlayStation Executable",
	KindPlayStationPortableExecutable:       "PlayStation Portable Executable",
	KindPlayStationPortableISO:              "PlayStation Portable ISO",
	KindPlayStationSoundFormat:              "PlayStation Sound Format",
	KindPLYModel:                            "PLY Model",
	KindPNGImage:                            "PNG Image",
	KindPortableExecutable:                  "Portable Executable",
	KindPostgreSQLCustomDump:                "PostgreSQL Custom Dump",
	KindPostScriptDocument:                  "PostScript Document",
	KindPowerBASICDebugger:                  "PowerBASIC Debugger",
	KindPowerBuilderIDE:                     "PowerBuilder IDE",
	KindPowerISODAA:                         "PowerISO DAA",
	KindPRTGDatabase:                        "PRTG Database",
	KindPSFFont:                             "PC Screen Font",
	KindPsiwtftabase:                       "Psion Database",
	KindPufferArchive:                       "Puffer Archive",
	KindPuttyPrivateKey:                     "PuTTY Private Key",
	KindPVRTexture:                          "PVR Texture",
	KindPythonBytecode:                      "Python Bytecode",
	KindPythonPickle:                        "Python Pickle",
	KindQCOWDiskImage:                       "QCOW Disk Image",
	KindQDBMDatabase:                        "QDBM Database",
	KindQEMUQEDDiskImage:                    "QEMU QED Disk Image",
	KindQimageFilter:                        "Qimage Filter",
	KindQOIImage:                            "QOI Image",
	KindQuakePAK:                            "Quake PAK Archive",
	KindQuarkExpress:                        "Quark Express",
	KindQuattroPro:                          "Quattro Pro",
	KindQuickBooksBackup:                    "QuickBooks Backup",
	KindQuickenData:                         "Quicken Data",
	KindQuickenQuickFinder:                  "Quicken QuickFinder",
	KindQuickReport:                         "QuickReport",
	KindRadianceHDRImage:                    "Radiance HDR Image",
	KindRagTimeDocument:                     "RagTime Document",
	KindRARArchive:                          "RAR Archive",
	KindRData:                               "RData",
	KindRealMedia:                           "RealMedia",
	KindRealMediaMetafile:                   "RealMedia Metafile",
	KindRealPlayerVideo:                     "RealPlayer Video",
	KindRedisDatabase:                       "Redis Database",
	KindREDRAWImage:                         "RED RAW Image",
	KindReiserFSFilesystem:                  "ReiserFS Filesystem",
	KindRenPyArchive:                        "Ren'Py Archive",
	KindRhino3DModel:                        "Rhino 3D Model",
	KindRichTextFormatDocument:              "Rich Text Format Document",
	KindRIFFContainer:                       "RIFF Container",
	KindRKAudio:                             "RKAudio",
	KindRNCArchive:                          "Rob Northen Compression Archive",
	KindROMFS:                               "ROMFS",
	KindROOTData:                            "ROOT Data",
	KindROSBag:                              "ROS Bag",
	KindRPGMakerArchive:                     "RPG Maker Archive",
	KindRPMPackage:                          "RPM Package",
	KindRSSFeed:                             "RSS Feed",
	KindRubyGemPackage:                      "RubyGem Package",
	KindRubyMarshal:                         "Ruby Marshal Data",
	KindRuntimeSoftwareImage:                "Runtime Software Image",
	KindRzipArchive:                         "rzip Archive",
	KindSASData:                             "SAS Data",
	KindScreamTrackerModule:                 "Scream Tracker Module",
	KindSegaCDImage:                         "Sega CD Image",
	KindSegaDreamcastROM:                    "Sega Dreamcast ROM",
	KindSegaMasterSystemROM:                 "Sega Master System ROM",
	KindSegaMegaDriveROM:                    "Sega Mega Drive ROM",
	KindSegaNaomiROM:                        "Sega NAOMI ROM",
	KindSegaSaturnROM:                       "Sega Saturn ROM",
	KindSGIImage:                            "SGI Image",
	KindShareazaThumbnail:                   "Shareaza Thumbnail",
	KindShebangScript:                       "Shebang Script",
	KindShockwaveFlash:                      "Shockwave Flash",
	KindShortenAudio:                        "Shorten Audio",
	KindShowPartnerGraphics:                 "Show Partner Graphics",
	KindSibeliusScore:                       "Sibelius Score",
	KindSietronicsDocument:                  "Sietronics Document",
	KindSigmaRAWImage:                       "Sigma RAW Image",
	KindSiliconGraphicsMovie:                "Silicon Graphics Movie",
	KindSketchDocument:                      "Sketch Document",
	KindSketchUpModel:                       "SketchUp Model",
	KindSkinCrafterSkin:                     "SkinCrafter Skin",
	KindSkypeData:                           "Skype Data",
	KindSkypeLocalization:                   "Skype Localization",
	KindSmackerVideo:                        "Smacker Video",
	KindSmartDrawDrawing:                    "SmartDraw Drawing",
	KindSnappyFramedData:                    "Snappy Framed Data",
	KindSNESSPC:                             "SNES SPC Audio",
	KindSnoopCapture:                        "Snoop Capture",
	KindSOAPMessage:                         "SOAP Message",
	KindSonicFoundryAcid:                    "Sonic Foundry Acid",
	KindSonyCompressedVoice:                 "Sony Compressed Voice",
	KindSonyOpenMG:                          "Sony OpenMG Audio",
	KindSonyWave64Audio:                     "Sony Wave64 Audio",
	KindSourceEngineBSP:                     "Source Engine BSP",
	KindSpeedtouchFirmware:                  "Speedtouch Firmware",
	KindSprintMusicStore:                    "Sprint Music Store",
	KindSPSSData:                            "SPSS Data",
	KindSPSSPortableData:                    "SPSS Portable Data",
	KindSPSSTemplate:                        "SPSS Template",
	KindSQLite3SharedMemory:                 "SQLite Shared Memory",
	KindSQLite3WriteAheadLog:                "SQLite Write-Ahead Log",
	KindSQLiteDatabase:                      "SQLite Database",
	KindSquashFSFilesystem:                  "SquashFS Filesystem",
	KindSQXArchive:                          "SQX Archive",
	KindSSHPublicKey:                        "SSH Public Key",
	KindStataData:                           "Stata Data",
	KindSteganosVirtualDrive:                "Steganos Virtual Drive",
	KindSTEP3DModel:                         "STEP 3D Model",
	KindSTLStereoLithography:                "STL StereoLithography",
	KindStorageCraftBackup:                  "StorageCraft Backup",
	KindStuffItArchive:                      "StuffIt Archive",
	KindSubRipText:                          "SubRip Text (SRT)",
	KindSunRasterImage:                      "Sun Raster Image",
	KindSuperCalcWorksheet:                  "SuperCalc Worksheet",
	KindSurfplanProject:                     "Surfplan Project",
	KindSVGImage:                            "SVG Image",
	KindSymantecGhostImage:                  "Symantec Ghost Image",
	KindSymantecWiseInstallerLog:            "Symantec Wise Installer Log",
	KindSymbianExecutable:                   "Symbian Executable",
	KindSymbianInstallationFormat:           "Symbian Installation Format",
	KindSystemdJournal:                      "Systemd Journal",
	KindSzipArchive:                         "Szip Archive",
	KindTajimaEmbroidery:                    "Tajima Embroidery",
	KindTAKAudio:                            "TAK Audio",
	KindTARArchive:                          "TAR Archive",
	KindTargetExpress:                       "TargetExpress",
	KindTensorFlowLiteModel:                 "TensorFlow Lite Model",
	KindTeraByteImage:                       "TeraByte Image",
	KindTextFile:                            "Text File",
	KindTheBatIndex:                         "The Bat! Index",
	KindThunderbirdMailSummary:              "Thunderbird Mail Summary",
	KindTIFFImage:                           "TIFF Image",
	KindTimezoneData:                        "Timezone Data",
	KindTokyoCabinetDatabase:                "Tokyo Cabinet Database",
	KindTomTomTraffic:                       "TomTom Traffic",
	KindTorrentFile:                         "Torrent File",
	KindTransportNeutralEncapsulationFormat: "Transport Neutral Encapsulation Format (TNEF)",
	KindTrueTypeCollection:                  "TrueType Collection",
	KindTrueTypeFont:                        "TrueType Font",
	KindTTAAudio:                            "TTA Audio",
	KindU3DModel:                            "U3D Model",
	KindUBIFSFilesystem:                     "UBIFS Filesystem",
	KindUBootImage:                          "U-Boot Image",
	KindUFAArchive:                          "UFA Archive",
	KindUFOCaptureMap:                       "UFO Capture Map",
	KindUHAArchive:                          "UHA Archive",
	KindUndergroundAudio:                    "Underground Audio",
	KindUnicodeExtensions:                   "Unicode Extensions",
	KindUnityWebData:                        "Unity Web Data",
	KindUniversalDiskFormat:                 "Universal Disk Format",
	KindUniversalSceneDescription:           "Universal Scene Description",
	KindUnixCompressArchive:                 "Unix Compress Archive",
	KindUnrealEnginePackage:                 "Unreal Engine Package",
	KindUUencodedArchive:                    "UUencoded Archive",
	KindValvePak:                            "Valve Pak",
	KindValveTextureFormat:                  "Valve Texture Format",
	KindVBScript:                            "VBScript",
	KindVBScriptEncoded:                     "VBScript Encoded",
	KindVCard:                               "vCard",
	KindVeeamBackup:                         "Veeam Backup",
	KindVHDDiskImage:                        "VHD Disk Image",
	KindVHDXDiskImage:                       "VHDX Disk Image",
	KindVideoGameMusic:                      "Video Game Music (VGM)",
	KindVideoVCD:                            "Video VCD",
	KindVirtualBoxDiskImage:                 "VirtualBox Disk Image",
	KindVirtualBoxSavedState:                "VirtualBox Saved State",
	KindVisualBasicControl:                  "Visual Basic Control",
	KindVisualCPreCompiledHeader:            "Visual C PreCompiled Header",
	KindVisualCWorkbenchInfo:                "Visual C++ Workbench Info",
	KindVisualStudioSolution:                "Visual Studio Solution",
	KindVMapSourceWaypoint:                  "VMapSource Waypoint",
	KindVMwareDiskImage:                     "VMware Disk Image",
	KindVMwareNVRAM:                         "VMware NVRAM",
	KindVMwareSnapshot:                      "VMware Snapshot",
	KindVMwareSuspend:                       "VMware Suspend",
	KindVocaloidProject:                     "Vocaloid Project",
	KindVocalTecMedia:                       "VocalTec Media",
	KindVOCAudio:                            "VOC Audio",
	KindVRML3DModel:                         "VRML 3D Model",
	KindVulkanSPIRV:                         "Vulkan SPIR-V",
	KindWADArchive:                          "WAD Archive",
	KindWalkmanMP3:                          "Walkman MP3",
	KindWARCFile:                            "WARC File",
	KindWarcraftIIIReplay:                   "Warcraft III Replay",
	KindWavPackAudio:                        "WavPack Audio",
	KindWavPackCorrection:                   "WavPack Correction",
	KindWebAssemblyModule:                   "WebAssembly Module",
	KindWebexARF:                            "Webex ARF",
	KindWebVTT:                              "WebVTT",
	KindWhereIsItCatalog:                    "WhereIsIt Catalog",
	KindWiiBackupFileSystem:                 "Wii Backup File System",
	KindWiiGameCubeTHP:                      "Wii GameCube THP",
	KindWiiROM:                              "Wii ROM",
	KindWiiUArchive:                         "Wii U Archive",
	KindWin95Password:                       "Windows 95 Password",
	KindWin9xPassword:                       "Win9x Password",
	KindWin9xPrinterSpool:                   "Win9x Printer Spool",
	KindWin9xRegistryHive:                   "Win9x Registry Hive",
	KindWindowsApplicationLog:               "Windows Application Log",
	KindWindowsCalendar:                     "Windows Calendar",
	KindWindowsEventLog:                     "Windows Event Log",
	KindWindowsEventViewer:                  "Windows Event Viewer",
	KindWindowsHelp:                         "Windows Help File",
	KindWindowsHibernation:                  "Windows Hibernation File",
	KindWindowsImagingFormat:                "Windows Imaging Format",
	KindWindowsInternationalCodePage:        "Windows International Code Page",
	KindWindowsMemoryDump:                   "Windows Memory Dump",
	KindWindowsMinidump:                     "Windows Minidump",
	KindWindowsPrecompiledHeader:            "Windows Precompiled Header",
	KindWindowsPrefetch:                     "Windows Prefetch",
	KindWindowsProgramManagerGroup:          "Windows Program Manager Group",
	KindWindowsRegistryHive:                 "Windows Registry Hive",
	KindWindowsResourceFile:                 "Windows Resource File",
	KindWindowsShortcut:                     "Windows Shortcut",
	KindWindowsThumbnailCache:               "Windows Thumbnail Cache",
	KindWindowsTypeLibrary:                  "Windows Type Library (TLB)",
	KindWindowsUserStateMigration:           "Windows User State Migration",
	KindWinNTNetmonCapture:                  "WinNT Netmon Capture",
	KindWinNTRegistryUndo:                   "WinNT Registry Undo",
	KindWinPharoahCapture:                   "WinPharoah Capture",
	KindWinPharoahFilter:                    "WinPharoah Filter",
	KindWiredTigerDatabase:                  "WiredTiger Database",
	KindWOFF2Font:                           "WOFF2 Font",
	KindWOFFFont:                            "WOFF Font",
	KindWord20Document:                      "Word 2.0 Document",
	KindWordPerfectDictionary:               "WordPerfect Dictionary",
	KindWordPerfectGraphics:                 "WordPerfect Graphics",
	KindWordPerfectText:                     "WordPerfect Text",
	KindWordStarDocument:                    "WordStar Document",
	KindWordStarWindows:                     "WordStar Windows",
	KindWTVVideo:                            "WTV Video",
	KindXARArchive:                          "XAR Archive",
	KindXBMImage:                            "X BitMap Image",
	KindXbox360Executable:                   "Xbox 360 Executable",
	KindXboxExecutable:                      "Xbox Executable",
	KindXboxISO:                             "Xbox ISO",
	KindXCOFFExecutable:                     "XCOFF Executable",
	KindXFSFilesystem:                       "XFS Filesystem",
	KindXilinxBitstream:                     "Xilinx Bitstream",
	KindXMLDocument:                         "XML Document",
	KindXPCOMLibraries:                      "XPCOM Libraries",
	KindXPMImage:                            "XPM Image",
	KindXZArchive:                           "XZ Archive",
	KindYamahaSMAF:                          "Yamaha SMAF",
	KindYouTubeTimedText:                    "YouTube Timed Text",
	KindYUV4MPEG2Video:                      "YUV4MPEG2 Video",
	KindZBackup:                             "ZBackup",
	KindZchunk:                              "Zchunk Archive",
	KindZFSFilesystem:                       "ZFS Filesystem",
	KindZIPArchive:                          "ZIP Archive",
	KindZisofsArchive:                       "zisofs Archive",
	KindZlibData:                            "Zlib Data",
	KindZOOArchive:                          "ZOO Archive",
	KindZoomBrowserIndex:                    "ZoomBrowser Index",
	KindZPAQArchive:                         "ZPAQ Archive",
	KindZstandardArchive:                    "Zstandard Archive",
	KindZstandardDictionary:                 "Zstandard Dictionary",
	KindZXTape:                              "ZX Spectrum Tape",
}

var typeNames = [...]string{
	TypeNone:                     "",
	Type32Bit:                    "32-bit",
	Type32BitBigEndian:           "32-bit Big-Endian",
	Type32BitLittleEndian:        "32-bit Little-Endian",
	Type3G2Media:                 "3G2 Media",
	Type3GPPMedia:                "3GPP Media",
	Type3MFDocument:              "3MF Document",
	Type64Bit:                    "64-bit",
	Type64BitBigEndian:           "64-bit Big-Endian",
	Type64BitLittleEndian:        "64-bit Little-Endian",
	TypeAC3:                      "AC-3",
	TypeAcronisB4:                "Acronis B4",
	TypeAcronisCE:                "Acronis CE",
	TypeAdobeDNGDNG:              "Adobe DNG (DNG)",
	TypeADTS:                     "ADTS",
	TypeAfterEffectsProjectAEP:   "After Effects Project (AEP)",
	TypeAIFCAudio:                "AIFC Audio",
	TypeAIFFAudio:                "AIFF Audio",
	TypeAndroidAppBundleAAB:      "Android App Bundle (AAB)",
	TypeAndroidArchiveAAR:        "Android Archive (AAR)",
	TypeAndroidPackageAPK:        "Android Package (APK)",
	TypeAndroidPackageXAPK:       "Android Package (XAPK)",
	TypeAndroidSplitAPKS:         "Android Split APK Set (APKS)",
	TypeAndroidSystemPackageAPEX: "Android System Package (APEX)",
	TypeAPPXPackage:              "APPX Package",
	TypeArchLinuxPackage:         "Arch Linux Package",
	TypeASCIIText:                "ASCII Text",
	TypeAVIFImage:                "AVIF Image",
	TypeAVIFImageSequence:        "AVIF Image Sequence",
	TypeAVIVideo:                 "AVI Video",
	TypeB4:                       "B4",
	TypeBashScript:               "Bash Script",
	TypeBatch:                    "Batch Script",
	TypeBGZF:                     "BGZF",
	TypeBigEndian:                "Big-Endian",
	TypeBigTIFF:                  "BigTIFF",
	TypeBinaryBigEndian:          "Binary Big-Endian",
	TypeBinaryLittleEndian:       "Binary Little-Endian",
	TypeBlackmagicRAW:            "Blackmagic RAW",
	TypeBlockDevice:              "Block Device",
	TypeByteSwapped:              "Byte-Swapped",
	TypeC:                        "C Source",
	TypeCanonRAW3CR3:             "Canon RAW 3 (CR3)",
	TypeCanonRAWHE:               "Canon RAW HE",
	TypeCAT:                      "CAT",
	TypeCB7:                      "CB7",
	TypeCBR:                      "CBR",
	TypeCBT:                      "CBT",
	TypeCBZ:                      "CBZ",
	TypeCDAAudio:                 "CD Audio",
	TypeCE:                       "CE",
	TypeCharacterDevice:          "Character Device",
	TypeCMake:                    "CMake Script",
	TypeCodestream:               "Codestream",
	TypeCwtfPackage:             "Cwtf Package",
	TypeContainer:                "Container",
	TypeCorelDRAWDocumentCDR:     "CorelDRAW Document (CDR)",
	TypeCOWD:                     "COWD",
	TypeCPP:                      "C++ Source",
	TypeCRXVersion2:              "Version 2",
	TypeCRXVersion3:              "Version 3",
	TypeCSharp:                   "C# Source",
	TypeCSS:                      "Cascading Style Sheets (CSS)",
	TypeCubaseProjectCPR:         "Cubase Project (CPR)",
	TypeDart:                     "Dart Source",
	TypeDEX035:                   "DEX 035",
	TypeDEX036:                   "DEX 036",
	TypeDEX037:                   "DEX 037",
	TypeDEX038:                   "DEX 038",
	TypeDEX039:                   "DEX 039",
	TypeDEX040:                   "DEX 040",
	TypeDEX041:                   "DEX 041",
	TypeDirectory:                "Directory",
	TypeDocker:                   "Dockerfile",
	TypeDownloadableSounds:       "Downloadable Sounds",
	TypeDSAPrivateKey:            "DSA Private Key",
	TypeEAC3:                     "E-AC-3",
	TypeECPrivateKey:             "EC Private Key",
	TypeELF:                      "ELF",
	TypeELF32:                    "ELF32",
	TypeELF32BigEndian:           "ELF32 Big-Endian",
	TypeELF32LittleEndian:        "ELF32 Little-Endian",
	TypeELF64:                    "ELF64",
	TypeELF64BigEndian:           "ELF64 Big-Endian",
	TypeELF64LittleEndian:        "ELF64 Little-Endian",
	TypeEmpty:                    "Empty",
	TypeEnhancedMetafileEMF:      "Enhanced Metafile (EMF)",
	TypeEPUBDocument:             "EPUB Document",
	TypeExt2:                     "ext2",
	TypeExt3:                     "ext3",
	TypeExt4:                     "ext4",
	TypeF4VVideo:                 "F4V Video",
	TypeFabricMod:                "Fabric Mod",
	TypeFAT12:                    "FAT12",
	TypeFAT16:                    "FAT16",
	TypeFAT32:                    "FAT32",
	TypeFirefoxExtensionXPI:      "Firefox Extension (XPI)",
	TypeFLACAudio:                "FLAC Audio",
	TypeForgeMod:                 "Forge Mod",
	TypeGIF87a:                   "GIF87a",
	TypeGIF89a:                   "GIF89a",
	TypeGo:                       "Go Source",
	TypeGraphQL:                  "GraphQL",
	TypeHE:                       "HE",
	TypeHEIFImage:                "HEIF Image",
	TypeILBMImage:                "ILBM Image",
	TypeIMMM:                     "IMMM",
	TypeINI:                      "INI Configuration",
	TypeIOSApplicationArchiveIPA: "iOS Application Archive (IPA)",
	TypeIWAD:                     "IWAD",
	TypeJava:                     "Java Source",
	TypeJavaArchiveJAR:           "Java Archive (JAR)",
	TypeJavaEnterpriseArchiveEAR: "Java Enterprise Archive (EAR)",
	TypeJavaScript:               "JavaScript Source",
	TypeJavaWebArchiveWAR:        "Java Web Archive (WAR)",
	TypeJMOD:                     "JMOD",
	TypeKDB:                      "KDB",
	TypeKDBX:                     "KDBX",
	TypeKDM:                      "KDM",
	TypeKerasModel:               "Keras Model",
	TypeKMZArchive:               "KMZ Archive",
	TypeKotlin:                   "Kotlin Source",
	TypeKritaDocumentKRA:         "Krita Document (KRA)",
	TypeKTX:                      "KTX",
	TypeKTX2:                     "KTX2",
	TypeLIST:                     "LIST",
	TypeLittleEndian:             "Little-Endian",
	TypeLN:                       "LN",
	TypeLOVEGame:                 "LÖVE Game",
	TypeLua:                      "Lua Script",
	TypeLZ4Legacy:                "LZ4 Legacy",
	TypeLZMACompressed:           "LZMA Compressed",
	TypeM2TS:                     "M2TS",
	TypeM2TSBDAV:                 "M2TS/BDAV",
	TypeM4VVideo:                 "M4V Video",
	TypeMakefile:                 "Makefile",
	TypeMarkdown:                 "Markdown Document",
	TypeMatroska:                 "Matroska",
	TypeMicrosoftExcelAddInXLAM:  "Microsoft Excel Add-In (XLAM)",
	TypeMicrosoftExcelMacroEnabledTemplateXLTM:          "Microsoft Excel Macro-Enabled Template (XLTM)",
	TypeMicrosoftExcelMacroEnabledWorkbookXLSM:          "Microsoft Excel Macro-Enabled Workbook (XLSM)",
	TypeMicrosoftExcelSpreadsheetXLSX:                   "Microsoft Excel Spreadsheet (XLSX)",
	TypeMicrosoftExcelTemplateXLTX:                      "Microsoft Excel Template (XLTX)",
	TypeMicrosoftExcelWorkbookXLS:                       "Microsoft Excel Workbook (XLS)",
	TypeMicrosoftInstallerMSI:                           "Microsoft Installer (MSI)",
	TypeMicrosoftOutlookMessageMSG:                      "Microsoft Outlook Message (MSG)",
	TypeMicrosoftPowerPointAddInPPAM:                    "Microsoft PowerPoint Add-In (PPAM)",
	TypeMicrosoftPowerPointMacroEnabledPresentationPPTM: "Microsoft PowerPoint Macro-Enabled Presentation (PPTM)",
	TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM:    "Microsoft PowerPoint Macro-Enabled Slideshow (PPSM)",
	TypeMicrosoftPowerPointMacroEnabledTemplatePOTM:     "Microsoft PowerPoint Macro-Enabled Template (POTM)",
	TypeMicrosoftPowerPointPresentationPPT:              "Microsoft PowerPoint Presentation (PPT)",
	TypeMicrosoftPowerPointPresentationPPTX:             "Microsoft PowerPoint Presentation (PPTX)",
	TypeMicrosoftPowerPointSlideshowPPSX:                "Microsoft PowerPoint Slideshow (PPSX)",
	TypeMicrosoftPowerPointTemplatePOTX:                 "Microsoft PowerPoint Template (POTX)",
	TypeMicrosoftProjectDocumentMPP:                     "Microsoft Project Document (MPP)",
	TypeMicrosoftPublisherDocumentPUB:                   "Microsoft Publisher Document (PUB)",
	TypeMicrosoftVisioDrawingVSD:                        "Microsoft Visio Drawing (VSD)",
	TypeMicrosoftWordDocumentDOC:                        "Microsoft Word Document (DOC)",
	TypeMicrosoftWordDocumentDOCX:                       "Microsoft Word Document (DOCX)",
	TypeMicrosoftWordMacroEnabledDocumentDOCM:           "Microsoft Word Macro-Enabled Document (DOCM)",
	TypeMicrosoftWordMacroEnabledTemplateDOTM:           "Microsoft Word Macro-Enabled Template (DOTM)",
	TypeMicrosoftWordTemplateDOTX:                       "Microsoft Word Template (DOTX)",
	TypeMinecraftResourcePack:                           "Minecraft Resource Pack",
	TypeMP3:                                             "MP3",
	TypeMP3ID3Tagged:                                    "MP3 (ID3 Tagged)",
	TypeMP4Video:                                        "MP4 Video",
	TypeMPEG12Video:                                     "MPEG-1/2 Video",
	TypeMPEG4AudioM4AFamily:                             "MPEG-4 Audio (M4A Family)",
	TypeMPEGLayer2:                                      "MPEG Layer II",
	TypeMPEGLayer3:                                      "MPEG Layer III",
	TypeMSIXPackage:                                     "MSIX Package",
	TypeNamedPipe:                                       "Named Pipe",
	TypeNanosecondBigEndian:                             "Nanosecond Big-Endian",
	TypeNanosecondLittleEndian:                          "Nanosecond Little-Endian",
	TypeNB:                                              "NB",
	TypeNewASCII:                                        "New ASCII",
	TypeNewASCIIWithCRC:                                 "New ASCII with CRC",
	TypeNikonRAWNEF:                                     "Nikon RAW (NEF)",
	TypeNpmPackageTarball:                               "npm Package Tarball",
	TypeNuGetPackageNUPKG:                               "NuGet Package (NUPKG)",
	TypeObjectiveC:                                      "Objective-C Source",
	TypeOCIImageLayoutTar:                               "OCI Image Layout (TAR)",
	TypeOldASCII:                                        "Old ASCII",
	TypeOpenDocumentChartODC:                            "OpenDocument Chart (ODC)",
	TypeOpenDocumentDatabaseODB:                         "OpenDocument Database (ODB)",
	TypeOpenDocumentFormulaODF:                          "OpenDocument Formula (ODF)",
	TypeOpenDocumentGraphicsODG:                         "OpenDocument Graphics (ODG)",
	TypeOpenDocumentImageODI:                            "OpenDocument Image (ODI)",
	TypeOpenDocumentPresentationODP:                     "OpenDocument Presentation (ODP)",
	TypeOpenDocumentSpreadsheetODS:                      "OpenDocument Spreadsheet (ODS)",
	TypeOpenDocumentTextODT:                             "OpenDocument Text (ODT)",
	TypeOpenRasterImageORA:                              "OpenRaster Image (ORA)",
	TypeOpusAudio:                                       "Opus Audio",
	TypePAM:                                             "PAM",
	TypePanasonicRAWRW2:                                 "Panasonic RAW (RW2)",
	TypePBMASCII:                                        "PBM ASCII",
	TypePBMBinary:                                       "PBM binary",
	TypePE32ARM:                                         "PE32 ARM",
	TypePE32ARM64:                                       "PE32 ARM64",
	TypePE32ARMv7:                                       "PE32 ARMv7",
	TypePE32Itanium:                                     "PE32 Itanium",
	TypePE32PlusARM64:                                   "PE32+ ARM64",
	TypePE32PlusUnknown:                                 "PE32+ Unknown",
	TypePE32PlusX8664:                                   "PE32+ x86-64",
	TypePE32Unknown:                                     "PE32 Unknown",
	TypePE32X86:                                         "PE32 x86",
	TypePE32X8664:                                       "PE32 x86-64",
	TypePEC:                                             "PEC",
	TypePentaxRAWPEF:                                    "Pentax RAW (PEF)",
	TypePerl:                                            "Perl Script",
	TypePES:                                             "PES",
	TypePGMASCII:                                        "PGM ASCII",
	TypePGMBinary:                                       "PGM binary",
	TypePHP:                                             "PHP Script",
	TypePowerShell:                                      "PowerShell Script",
	TypePPMASCII:                                        "PPM ASCII",
	TypePPMBinary:                                       "PPM binary",
	TypePSB:                                             "PSB",
	TypePSD:                                             "PSD",
	TypePWAD:                                            "PWAD",
	TypePython:                                          "Python Script",
	TypePythonSourceDistributionSDist:                   "Python Source Distribution (sdist)",
	TypePythonWheelWHL:                                  "Python Wheel (WHL)",
	TypePyTorchModel:                                    "PyTorch Model",
	TypeQCOW1:                                           "QCOW1",
	TypeQCOW2:                                           "QCOW2",
	TypeQCPAudio:                                        "QCP Audio",
	TypeQuickTimeMovie:                                  "QuickTime Movie",
	TypeR:                                               "R Script",
	TypeRAR4:                                            "RAR4",
	TypeRAR5:                                            "RAR5",
	TypeReturnPath:                                      "Return Path",
	TypeRIFFMIDI:                                        "MIDI",
	TypeRIFFPalette:                                     "Palette",
	TypeRSAPrivateKey:                                   "RSA Private Key",
	TypeRuby:                                            "Ruby Script",
	TypeRust:                                            "Rust Source",
	TypeScala:                                           "Scala Source",
	TypeSketchDocument:                                  "Sketch Document",
	TypeSlackwarePackage:                                "Slackware Package",
	TypeSocket:                                          "Socket",
	TypeSonyRAWARW:                                      "Sony RAW (ARW)",
	TypeSonyRAWSR2:                                      "Sony RAW (SR2)",
	TypeSoundFont2:                                      "SoundFont 2",
	TypeSpanned:                                         "Spanned",
	TypeSpecial:                                         "Special",
	TypeSpeexAudio:                                      "Speex Audio",
	TypeSQL:                                             "SQL Script",
	TypeStreamVersion7:                                  "Stream Version 7",
	TypeStreamVersion8:                                  "Stream Version 8",
	TypeSwift:                                           "Swift Source",
	TypeSymbolicLink:                                    "Symbolic Link",
	TypeSZ:                                              "SZ",
	TypeTerraform:                                       "Terraform Configuration",
	TypeTheoraVideo:                                     "Theora Video",
	TypeTOML:                                            "TOML Configuration",
	TypeTrue:                                            "True",
	TypeTS:                                              "TS",
	TypeTypeScript:                                      "TypeScript Source",
	TypeUncompressed:                                    "Uncompressed",
	TypeUTF8Text:                                        "UTF-8 Text",
	TypeVagrantBox:                                      "Vagrant Box",
	TypeVisualStudioExtensionVSIX:                       "Visual Studio Extension (VSIX)",
	TypeVMDK:                                            "VMDK",
	TypeVMDKDescription:                                 "VMDK Description",
	TypeVorbisAudio:                                     "Vorbis Audio",
	TypeWAVAudio:                                        "WAV Audio",
	TypeWebM:                                            "WebM",
	TypeWebPImage:                                       "WebP Image",
	TypeWindowsAnimatedCursor:                           "Windows Animated Cursor",
	TypeWindowsCursor:                                   "Windows Cursor",
	TypeWindowsIcon:                                     "Windows Icon",
	TypeWindowsLE:                                       "Linear Executable (LE)",
	TypeWindowsLX:                                       "OS/2 Linear Executable (LX)",
	TypeWindowsMetafileWMF:                              "Windows Metafile (WMF)",
	TypeWindowsNE:                                       "16-bit New Executable (NE)",
	TypeWrapper:                                         "Wrapper",
	TypeXMLPaperSpecificationXPS:                        "XML Paper Specification (XPS)",
	TypeYAML:                                            "YAML Configuration",
	TypeZlibCompressed:                                  "Zlib Compressed",
	TypeZstandardSkinnableFrame:                         "Zstandard Skinnable Frame",
}

func (k KindID) String() string {
	if int(k) >= 0 && int(k) < len(kindNames) {
		name := kindNames[k]
		if name != "" {
			return name
		}
	}

	return "Unknown"
}

func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
