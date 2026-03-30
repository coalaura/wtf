package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind1PasswordCloudKeychain
	Kind1PasswordKeychain
	Kind7ZipArchive
	KindAACAudio
	KindAccessDataFTKEvidence
	KindACEArchive
	KindAcrobatFormsDataFormat
	KindAcronisTrueImage
	KindAdobeFrameMakerDocument
	KindAdobeInDesignDocument
	KindAdvancedForensicFormat
	KindAINArchive
	KindAllegroPackfile
	KindALZArchive
	KindAmigaDiskImage
	KindAmigaDiskMasherArchive
	KindAmigaHardDiskImage
	KindAmigaHunkExecutable
	KindAmigaIcon
	KindAmigaOctaMED
	KindAMRAudio
	KindAMRWBAudio
	KindAnalogBoxCircuit
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
	KindApacheArrow
	KindApacheParquet
	KindAPFSFilesystem
	KindAppImage
	KindAppleArchive
	KindAppleBillOfMaterials
	KindAppleBinaryCrashReport
	KindAppleBinaryPropertyList
	KindAppleDesktopServicesStore
	KindAppleDiskImage
	KindAppleDoubleArchive
	KindAppleISOHybridDiskImage
	KindAppleiWorkDocument
	KindAppleKeychain
	KindAppleSingleArchive
	KindAppleSystemLog
	KindAppleTextureImage
	KindAppleXMLPropertyList
	KindApproachIndex
	KindARArchive
	KindARCArchive
	KindARJArchive
	KindARRIRAWImage
	KindASFContainer
	KindASTCTextureImage
	KindAtari7800ROM
	KindAtariLynxROM
	KindAtomFeed
	KindAUAudio
	KindAudacityBlock
	KindAutoCADDrawing
	KindAutoCADDXFDrawing
	KindAVG6IntegrityDatabase
	KindAvroObjectContainer
	KindBAMData
	KindBase85
	KindBcachefsFilesystem
	KindBCFData
	KindBerkeleyDatabase
	KindBGBlitzDatabase
	KindBinHexArchive
	KindBinkVideo
	KindBitcoinBlock
	KindBitcoinWallet
	KindBitLockerDiskEncryption
	KindBlackBerryBackup
	KindBlenderProject
	KindBlinkArchive
	KindBlizzardTextureImage
	KindBluetoothSnoopCapture
	KindBMPImage
	KindBochsDiskImage
	KindBouncyCastleKeystore
	KindBPGImage
	KindBroadbandEBookDocument
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
	KindCHMDocument
	KindChromaGraphImage
	KindChromiumPatch
	KindCinema4DModel
	KindCineonImage
	KindCloneCDControl
	KindCommodore64TapeImage
	KindCommodore64TapeRAWImage
	KindCommodoreSIDAudio
	KindCompiledTerminfo
	KindCOMPlusCatalog
	KindCompressedSquareWave
	KindCorelBinaryMetafile
	KindCorelColorPalette
	KindCorelPhotoPaintImage
	KindCPIOArchive
	KindCRAMData
	KindCramFSFilesystem
	KindCrushArchive
	KindCRXBrowserExtension
	KindCSOArchive
	KindCsoundAudio
	KindDalvikExecutable
	KindDAXArchive
	KindDB2Conversion
	KindDDSImage
	KindDebianPackage
	KindDeluxePaintAnimation
	KindDesignTools2DDrawing
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
	KindDVRMSVideo
	KindEasyStreetDrawDrawing
	KindEBMLContainer
	KindEFax
	KindEGGArchive
	KindElitePlusCommander
	KindEncapsulatedPostScript
	KindEnCaseCase
	KindEnCaseEvidenceV2
	KindEnCaseImage
	KindEndNoteLibrary
	KindEOTFont
	KindErlangBEAMBytecode
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
	KindFLICAnimation
	KindFLIFImage
	KindFlightSimulatorConfig
	KindFLStudioProject
	KindFLVVideo
	KindFreeArcArchive
	KindFujifilmRAWImage
	KindFuzzyBitmapImage
	KindGameBoyAdvanceROM
	KindGameBoyAudio
	KindGameBoyROM
	KindGameCubeROM
	KindGDBMDatabase
	KindGEMRasterImage
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
	KindGLTFBinaryModel
	KindGlyphBitmapDistributionFormat
	KindGNOMEKeyring
	KindGNUGettextMachineCatalog
	KindGnuPGKeybox
	KindGodotPackage
	KindGoogleChromeDictionary
	KindGoogleDriveDrawing
	KindGoogleEarthPlacemark
	KindGPSExchangeFormat
	KindGRIBData
	KindGUIDPartitionTable
	KindGzipArchive
	KindGzipData
	KindHadoopRC
	KindHadoopSequence
	KindHamarsoftArchive
	KindHarvardGraphicsImage
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
	KindIFFContainer
	KindImgSoftwareImage
	KindImpulseTrackerInstrument
	KindImpulseTrackerModule
	KindImpulseTrackerSample
	KindInnoSetupLog
	KindInstallShieldCabinet
	KindInstallShieldScript
	KindISO9660DiskImage
	KindISOBaseMedia
	KindIVFVideo
	KindJarArchive
	KindJARCSArchive
	KindJavaClassBytecode
	KindJavaCryptographyExtensionKeyStore
	KindJavaKeyStore
	KindJavaModule
	KindJavaPack200Archive
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
	KindKTXTextureImage
	KindKyotoCabinetDatabase
	KindLHAArchive
	KindLidarData
	KindLightWaveScene
	KindLinuxKernelImage
	KindLLVMBitcode
	KindLogicalFileEvidence
	KindLottieAnimation
	KindLotus123Spreadsheet
	KindLotusAMIDocument
	KindLotusNotesDatabase
	KindLotusNotesDatabaseTemplate
	KindLotusWordProDocument
	KindLrzipArchive
	KindLuaBytecode
	KindLuceneIndex
	KindLUKSDiskEncryption
	KindLVM2PhysicalVolume
	KindLytroLightFieldImage
	KindLZ4Frame
	KindLZFArchive
	KindLZFSEData
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindLZXArchive
	KindM3U8Playlist
	KindMacBinaryArchive
	KindMachOBinary
	KindMachOUniversalBinary
	KindMacOSAlias
	KindMacOSXCodeSignature
	KindMacriumReflectDiskImage
	KindMagicaVoxelModel
	KindMAMECHDDiskImage
	KindMArArchive
	KindMaterialExchangeFormat
	KindMATLABData
	KindMayaASCIIModel
	KindMayaBinaryModel
	KindMBOXEmailFolder
	KindMBOXTableOfContents
	KindMCAPCapture
	KindMeasurementDataFormat
	KindMediaDescriptor
	KindMerriamWebsterDictionary
	KindMetafileImage
	KindMicrografxDrawing
	KindMicrosoftAccessDatabase
	KindMicrosoftAnswerWizard
	KindMicrosoftCodePageTranslation
	KindMicrosoftCompress
	KindMicrosoftDeveloperStudioProject
	KindMicrosoftExchangeConfig
	KindMicrosoftInfo
	KindMicrosoftJournalDocument
	KindMicrosoftMoney
	KindMicrosoftNetworkMonitorCapture
	KindMicrosoftOneNoteNote
	KindMicrosoftOutlookEmailFolder
	KindMicrosoftProgramDatabase
	KindMicrosoftReaderDocument
	KindMicrosoftSQLServerDatabase
	KindMicrosoftWinMobileNoteDocument
	KindMicrosoftWorksSpreadsheet
	KindMicrosoftWriteDocument
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
	KindMSDOSCOFFObject
	KindMUGENAudio
	KindMUGENSprite
	KindMultiBitBlockchain
	KindMultiBitInfo
	KindMultiBitWallet
	KindMusepackAudio
	KindMySQLMyISAMDatabase
	KindNationalImageryTransmission
	KindNationalTransferFormat
	KindNAVQuarantinedVirus
	KindNeoGeoPocketROM
	KindNeroCDCompilation
	KindNESAudio
	KindNESROM
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
	KindNintendoRARCArchive
	KindNintendoSwitchNROExecutable
	KindNintendoSwitchNSOExecutable
	KindNintendoSwitchPackage
	KindNintendoSwitchROM
	KindNintendoTPL
	KindNintendoU8Archive
	KindNintendoWiiUExecutable
	KindNintendoYay0Archive
	KindNintendoYaz0Archive
	KindNISTSPHEREAudio
	KindNOAARasterChart
	KindNokiaPhoneBackup
	KindNovellLANalyzerCapture
	KindNTFSFilesystem
	KindNullsoftVideo
	KindNumPyArray
	KindNVIDIASceneGraph
	KindOCamlObject
	KindOggContainer
	KindOLECompoundDocument
	KindOpenEXRImage
	KindOpenSSHPrivateKey
	KindOpenTypeFont
	KindOptimFROGAudio
	KindORCColumnarData
	KindOutlookAddress
	KindOutlookExpressAddressBook
	KindOutlookExpressDatabase
	KindPacketSnifferXCP
	KindPaintShopProImage
	KindPalmData
	KindParallelsDiskImage
	KindParrotVideoEncapsulation
	KindPathWayMap
	KindPAXImage
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCFFont
	KindPCXDCXImage
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
	KindPlayStation1Executable
	KindPlayStation2MemoryCard
	KindPlayStationAudio
	KindPlayStationExecutable
	KindPlayStationPortableExecutable
	KindPlayStationPortableISO
	KindPLYModel
	KindPNGImage
	KindPortableExecutable
	KindPostgreSQLCustomDump
	KindPostScriptDocument
	KindPowerBASICDebugger
	KindPowerBuilderIDE
	KindPowerISODAA
	KindProTrackerModule
	KindPRTGDatabase
	KindPSFFont
	KindPsionDatabase
	KindPufferArchive
	KindPuttyPrivateKey
	KindPVRTextureImage
	KindPythonBytecode
	KindPythonPickleData
	KindQCOWDiskImage
	KindQDBMDatabase
	KindQEMUQEDDiskImage
	KindQimageFilter
	KindQOIImage
	KindQuarkExpressDocument
	KindQuattroProSpreadsheet
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
	KindROMFSFilesystem
	KindROOTData
	KindROSBag
	KindRPGMakerArchive
	KindRPMPackage
	KindRSSFeed
	KindRubyGemPackage
	KindRubyMarshalData
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
	KindShareazaThumbnailImage
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
	KindSNESSPCAudio
	KindSnoopCapture
	KindSOAPMessage
	KindSonicFoundryAcid
	KindSonyCompressedVoice
	KindSonyOpenMG
	KindSonyWave64Audio
	KindSourceEngineBSPMap
	KindSpeedtouchFirmware
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
	KindSuperCalcSpreadsheet
	KindSurfplanProject
	KindSVGImage
	KindSymantecGhostDiskImage
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
	KindTeraByteDiskImage
	KindText
	KindTheBatIndex
	KindThunderbirdMailSummary
	KindTIFFImage
	KindTimezoneData
	KindTokyoCabinetDatabase
	KindTomTomTrafficData
	KindTorrent
	KindTransportNeutralEncapsulationFormat
	KindTrueTypeCollection
	KindTrueTypeFont
	KindTTAAudio
	KindU3DModel
	KindUBIFSFilesystem
	KindUBootImage
	KindUFAArchive
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
	KindValveTextureFormatImage
	KindVBScriptEncoded
	KindVCard
	KindVeeamBackup
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVideoCD
	KindVideoGameAudio
	KindVirtualBoxDiskImage
	KindVirtualBoxSavedState
	KindVisualCWorkbenchInfo
	KindVisualStudioSolution
	KindVMapSourceWaypoint
	KindVMwareDiskImage
	KindVMwareNVRAM
	KindVMwareSnapshot
	KindVMwareSuspend
	KindVocaloidProject
	KindVocalTecAudio
	KindVOCAudio
	KindVRML3DModel
	KindVulkanSPIRV
	KindWADArchive
	KindWalkmanMP3
	KindWARC
	KindWarcraftIIIReplay
	KindWavPackAudio
	KindWavPackCorrection
	KindWebAssemblyModule
	KindWebexARFVideo
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
	KindWinAmpPlaylist
	KindWindowsApplicationLog
	KindWindowsCalendar
	KindWindowsEventLog
	KindWindowsEventViewer
	KindWindowsHelp
	KindWindowsHibernationData
	KindWindowsImagingFormat
	KindWindowsInternationalCodePage
	KindWindowsMemoryDump
	KindWindowsMinidump
	KindWindowsPrecompiledHeader
	KindWindowsPrefetch
	KindWindowsProgramManagerGroup
	KindWindowsRegistryHive
	KindWindowsResource
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
	KindWordPerfectGraphicsImage
	KindWordPerfectTextDocument
	KindWordStarDocument
	KindWordStarWindows
	KindWTVVideo
	KindXARArchive
	KindXBitMapImage
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
	KindZchunkArchive
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
	TypeAPPXPackage
	TypeArchLinuxPackage
	TypeASCII
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeB4
	TypeBashScript
	TypeBatchScript
	TypeBGZF
	TypeBigEndian
	TypeBigTIFF
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
	TypeCMakeScript
	TypeCodestream
	TypeComicBook
	TypeCondaPackage
	TypeContainer
	TypeCorelDRAWDocument
	TypeCOWD
	TypeCPPSource
	TypeCSharpSource
	TypeCSource
	TypeCSS
	TypeCubaseProject
	TypeDartSource
	TypeDirectory
	TypeDockerfile
	TypeDownloadableSounds
	TypeDSA
	TypeEAC3
	TypeEC
	TypeEmpty
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
	TypeFLACAudio
	TypeForgeMod
	TypeGIF87a
	TypeGIF89a
	TypeGoSource
	TypeGraphQL
	TypeHDPhoto
	TypeHE
	TypeHEIFImage
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
	TypeKDB
	TypeKDBX
	TypeKDM
	TypeKerasModel
	TypeKMZArchive
	TypeKotlinSource
	TypeKritaDocument
	TypeLegacy
	TypeLIST
	TypeLittleEndian
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
	TypeNikonRAW
	TypeNpmPackage
	TypeNuGetPackage
	TypeOBJ
	TypeObjectiveCSource
	TypeOCIImageLayout
	TypeOldASCII
	TypeOlympusRAW
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
	TypeRIFFPalette
	TypeRSA
	TypeRScript
	TypeRubyScript
	TypeRustSource
	TypeScalaSource
	TypeSkinnableFrame
	TypeSlackwarePackage
	TypeSocket
	TypeSonyRAW
	TypeSonyRAWSR2
	TypeSoundFont2
	TypeSpecial
	TypeSpeexAudio
	TypeSQLScript
	TypeSwiftSource
	TypeSymbolicLink
	TypeTerraformConfiguration
	TypeTheoraVideo
	TypeTOMLConfiguration
	TypeTypeScriptSource
	TypeUncompressed
	TypeUTF8
	TypeVagrantBox
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
	TypeVisualStudioExtension
	TypeVMDK
	TypeVMDKDescription
	TypeVorbisAudio
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
)

var kindNames = [...]string{
	KindUnknown:                             "Unknown",
	Kind1PasswordCloudKeychain:              "1Password Cloud Keychain",
	Kind1PasswordKeychain:                   "1Password Keychain",
	Kind7ZipArchive:                         "7-Zip Archive",
	KindAACAudio:                            "AAC Audio",
	KindAccessDataFTKEvidence:               "AccessData FTK Evidence",
	KindACEArchive:                          "ACE Archive",
	KindAcrobatFormsDataFormat:              "Acrobat Forms Data Format",
	KindAcronisTrueImage:                    "Acronis True Image",
	KindAdobeFrameMakerDocument:             "Adobe FrameMaker Document",
	KindAdobeInDesignDocument:               "Adobe InDesign Document",
	KindAdvancedForensicFormat:              "Advanced Forensic Format",
	KindAINArchive:                          "AIN Archive",
	KindAllegroPackfile:                     "Allegro Packfile",
	KindALZArchive:                          "ALZ Archive",
	KindAmigaDiskImage:                      "Amiga Disk Image",
	KindAmigaDiskMasherArchive:              "Amiga DiskMasher Archive",
	KindAmigaHardDiskImage:                  "Amiga Hard Disk Image",
	KindAmigaHunkExecutable:                 "Amiga Hunk Executable",
	KindAmigaIcon:                           "Amiga Icon",
	KindAmigaOctaMED:                        "Amiga OctaMED Audio",
	KindAMRAudio:                            "AMR Audio",
	KindAMRWBAudio:                          "AMR-WB Audio",
	KindAnalogBoxCircuit:                    "Analog Box Circuit",
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
	KindApacheArrow:                         "Apache Arrow",
	KindApacheParquet:                       "Apache Parquet",
	KindAPFSFilesystem:                      "APFS Filesystem",
	KindAppImage:                            "AppImage",
	KindAppleArchive:                        "Apple Archive",
	KindAppleBillOfMaterials:                "Apple Bill of Materials",
	KindAppleBinaryCrashReport:              "Apple Binary Crash Report",
	KindAppleBinaryPropertyList:             "Apple Binary Property List",
	KindAppleDesktopServicesStore:           "Apple Desktop Services Store",
	KindAppleDiskImage:                      "Apple Disk Image",
	KindAppleDoubleArchive:                  "AppleDouble Archive",
	KindAppleISOHybridDiskImage:             "Apple ISO Hybrid CD Image",
	KindAppleiWorkDocument:                  "Apple iWork Document",
	KindAppleKeychain:                       "Apple Keychain",
	KindAppleSingleArchive:                  "AppleSingle Archive",
	KindAppleSystemLog:                      "Apple System Log",
	KindAppleTextureImage:                   "Apple Texture Image",
	KindAppleXMLPropertyList:                "Apple XML Property List",
	KindApproachIndex:                       "Approach Index",
	KindARArchive:                           "AR Archive",
	KindARCArchive:                          "ARC Archive",
	KindARJArchive:                          "ARJ Archive",
	KindARRIRAWImage:                        "ARRI RAW Image",
	KindASFContainer:                        "ASF Container",
	KindASTCTextureImage:                    "ASTC Texture Image",
	KindAtari7800ROM:                        "Atari 7800 ROM",
	KindAtariLynxROM:                        "Atari Lynx ROM",
	KindAtomFeed:                            "Atom Feed",
	KindAUAudio:                             "AU Audio",
	KindAudacityBlock:                       "Audacity Block",
	KindAutoCADDrawing:                      "AutoCAD Drawing",
	KindAutoCADDXFDrawing:                   "AutoCAD DXF Drawing",
	KindAVG6IntegrityDatabase:               "AVG6 Integrity Database",
	KindAvroObjectContainer:                 "Avro Object Container",
	KindBAMData:                             "BAM Data",
	KindBase85:                              "BASE85",
	KindBcachefsFilesystem:                  "Bcachefs Filesystem",
	KindBCFData:                             "Binary Call Format",
	KindBerkeleyDatabase:                    "Berkeley Database",
	KindBGBlitzDatabase:                     "BGBlitz Database",
	KindBinHexArchive:                       "BinHex Archive",
	KindBinkVideo:                           "Bink Video",
	KindBitcoinBlock:                        "Bitcoin Block Data",
	KindBitcoinWallet:                       "Bitcoin Wallet",
	KindBitLockerDiskEncryption:             "BitLocker Disk Encryption",
	KindBlackBerryBackup:                    "BlackBerry Backup",
	KindBlenderProject:                      "Blender",
	KindBlinkArchive:                        "Blink Archive",
	KindBlizzardTextureImage:                "Blizzard Texture Image",
	KindBluetoothSnoopCapture:               "Bluetooth Snoop Capture",
	KindBMPImage:                            "BMP Image",
	KindBochsDiskImage:                      "Bochs Disk Image",
	KindBouncyCastleKeystore:                "BouncyCastle Keystore",
	KindBPGImage:                            "BPG Image",
	KindBroadbandEBookDocument:              "Broadband eBook Document",
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
	KindCHMDocument:                         "Compiled HTML Help",
	KindChromaGraphImage:                    "ChromaGraph Image",
	KindChromiumPatch:                       "Chromium Patch",
	KindCinema4DModel:                       "Cinema 4D Model",
	KindCineonImage:                         "Cineon Image",
	KindCloneCDControl:                      "CloneCD Control",
	KindCommodore64TapeImage:                "Commodore 64 Tape Image",
	KindCommodore64TapeRAWImage:             "Commodore 64 Tape RAW Image",
	KindCommodoreSIDAudio:                   "Commodore SID Audio",
	KindCompiledTerminfo:                    "Compiled Terminfo",
	KindCOMPlusCatalog:                      "COM+ Catalog",
	KindCompressedSquareWave:                "Compressed Square Wave",
	KindCorelBinaryMetafile:                 "Corel Binary Metafile",
	KindCorelColorPalette:                   "Corel Color Palette",
	KindCorelPhotoPaintImage:                "Corel Photo-Paint Image",
	KindCPIOArchive:                         "CPIO Archive",
	KindCRAMData:                            "CRAM Data",
	KindCramFSFilesystem:                    "CramFS Filesystem",
	KindCrushArchive:                        "Crush Archive",
	KindCRXBrowserExtension:                 "CRX Browser Extension",
	KindCSOArchive:                          "CSO Compressed ISO",
	KindCsoundAudio:                         "Csound Audio",
	KindDalvikExecutable:                    "Dalvik Executable",
	KindDAXArchive:                          "DAX Archive",
	KindDB2Conversion:                       "DB2 Conversion",
	KindDDSImage:                            "DDS Image",
	KindDebianPackage:                       "Debian Package",
	KindDeluxePaintAnimation:                "DeluxePaint Animation",
	KindDesignTools2DDrawing:                "DesignTools 2D Drawing",
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
	KindDVRMSVideo:                          "DVR-MS Video",
	KindEasyStreetDrawDrawing:               "Easy Street Draw Drawing",
	KindEBMLContainer:                       "EBML Container",
	KindEFax:                                "eFax",
	KindEGGArchive:                          "EGG Archive",
	KindElitePlusCommander:                  "Elite Plus Commander",
	KindEncapsulatedPostScript:              "Encapsulated PostScript",
	KindEnCaseCase:                          "EnCase Case",
	KindEnCaseEvidenceV2:                    "EnCase Evidence V2",
	KindEnCaseImage:                         "EnCase Image",
	KindEndNoteLibrary:                      "EndNote Library",
	KindEOTFont:                             "EOT Font",
	KindErlangBEAMBytecode:                  "Erlang BEAM Bytecode",
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
	KindFLICAnimation:                       "FLIC Animation",
	KindFLIFImage:                           "FLIF Image",
	KindFlightSimulatorConfig:               "Flight Simulator Config",
	KindFLStudioProject:                     "FL Studio Project",
	KindFLVVideo:                            "FLV Video",
	KindFreeArcArchive:                      "FreeArc Archive",
	KindFujifilmRAWImage:                    "Fujifilm RAW Image",
	KindFuzzyBitmapImage:                    "Fuzzy Bitmap Image",
	KindGameBoyAdvanceROM:                   "Game Boy Advance ROM",
	KindGameBoyAudio:                        "Game Boy Audio",
	KindGameBoyROM:                          "Game Boy ROM",
	KindGameCubeROM:                         "GameCube ROM",
	KindGDBMDatabase:                        "GDBM Database",
	KindGEMRasterImage:                      "GEM Raster Image",
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
	KindGLTFBinaryModel:                     "glTF Binary Model",
	KindGlyphBitmapDistributionFormat:       "Glyph Bitmap Distribution Format",
	KindGNOMEKeyring:                        "GNOME Keyring",
	KindGNUGettextMachineCatalog:            "GNU Gettext Machine Catalog",
	KindGnuPGKeybox:                         "GnuPG Keybox",
	KindGodotPackage:                        "Godot Engine Package",
	KindGoogleChromeDictionary:              "Google Chrome Dictionary",
	KindGoogleDriveDrawing:                  "Google Drive Drawing",
	KindGoogleEarthPlacemark:                "Google Earth Placemark",
	KindGPSExchangeFormat:                   "GPS Exchange Format",
	KindGRIBData:                            "GRIB Data",
	KindGUIDPartitionTable:                  "GUID Partition Table",
	KindGzipArchive:                         "Gzip Archive",
	KindGzipData:                            "Gzip Data",
	KindHadoopRC:                            "Hadoop RC",
	KindHadoopSequence:                      "Hadoop Sequence",
	KindHamarsoftArchive:                    "Hamarsoft Archive",
	KindHarvardGraphicsImage:                "Harvard Graphics Image",
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
	KindIFFContainer:                        "IFF Container",
	KindImgSoftwareImage:                    "Img Software Image",
	KindImpulseTrackerInstrument:            "Impulse Tracker Instrument",
	KindImpulseTrackerModule:                "Impulse Tracker Module",
	KindImpulseTrackerSample:                "Impulse Tracker Sample",
	KindInnoSetupLog:                        "Inno Setup Log",
	KindInstallShieldCabinet:                "InstallShield Cabinet",
	KindInstallShieldScript:                 "InstallShield Script",
	KindISO9660DiskImage:                    "ISO 9660 Disk Image",
	KindISOBaseMedia:                        "ISO Base Media",
	KindIVFVideo:                            "IVF Video",
	KindJarArchive:                          "JAR Archive",
	KindJARCSArchive:                        "JARCS Archive",
	KindJavaClassBytecode:                   "Java Class Bytecode",
	KindJavaCryptographyExtensionKeyStore:   "JCE KeyStore",
	KindJavaKeyStore:                        "Java KeyStore",
	KindJavaModule:                          "Java Module",
	KindJavaPack200Archive:                  "Java Pack200 Archive",
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
	KindKTXTextureImage:                     "KTX Texture Image",
	KindKyotoCabinetDatabase:                "Kyoto Cabinet Database",
	KindLHAArchive:                          "LHA Archive",
	KindLidarData:                           "LiDAR Data",
	KindLightWaveScene:                      "LightWave Scene",
	KindLinuxKernelImage:                    "Linux Kernel Image",
	KindLLVMBitcode:                         "LLVM Bitcode",
	KindLogicalFileEvidence:                 "Logical File Evidence",
	KindLottieAnimation:                     "Lottie Animation",
	KindLotus123Spreadsheet:                 "Lotus 1-2-3 Spreadsheet",
	KindLotusAMIDocument:                    "Lotus AMI Document",
	KindLotusNotesDatabase:                  "Lotus Notes Database",
	KindLotusNotesDatabaseTemplate:          "Lotus Notes Database Template",
	KindLotusWordProDocument:                "Lotus WordPro Document",
	KindLrzipArchive:                        "Long Range ZIP Archive",
	KindLuaBytecode:                         "Lua Bytecode",
	KindLuceneIndex:                         "Lucene Index",
	KindLUKSDiskEncryption:                  "LUKS Disk Encryption",
	KindLVM2PhysicalVolume:                  "LVM2 Physical Volume",
	KindLytroLightFieldImage:                "Lytro Light Field",
	KindLZ4Frame:                            "LZ4 Frame",
	KindLZFArchive:                          "LZF Archive",
	KindLZFSEData:                           "LZFSE Data",
	KindLZIPArchive:                         "LZIP Archive",
	KindLZMAData:                            "LZMA Data",
	KindLZOPArchive:                         "LZOP Archive",
	KindLZXArchive:                          "LZX Archive",
	KindM3U8Playlist:                        "M3U8 Playlist",
	KindMacBinaryArchive:                    "MacBinary Archive",
	KindMachOBinary:                         "Mach-O Binary",
	KindMachOUniversalBinary:                "Mach-O Universal Binary",
	KindMacOSAlias:                          "macOS Alias",
	KindMacOSXCodeSignature:                 "macOS Code Signature",
	KindMacriumReflectDiskImage:             "Macrium Reflect Disk Image",
	KindMagicaVoxelModel:                    "MagicaVoxel Model",
	KindMAMECHDDiskImage:                    "MAME CHD Disk Image",
	KindMArArchive:                          "MAr Archive",
	KindMaterialExchangeFormat:              "Material Exchange Format",
	KindMATLABData:                          "MATLAB Data",
	KindMayaASCIIModel:                      "Maya ASCII Model",
	KindMayaBinaryModel:                     "Maya Binary Model",
	KindMBOXEmailFolder:                     "MBOX Email Folder",
	KindMBOXTableOfContents:                 "MBOX Table of Contents",
	KindMCAPCapture:                         "MCAP Capture",
	KindMeasurementDataFormat:               "Measurement Data Format",
	KindMediaDescriptor:                     "Media Descriptor",
	KindMerriamWebsterDictionary:            "Merriam-Webster Dictionary",
	KindMetafileImage:                       "Metafile Image",
	KindMicrografxDrawing:                   "Micrografx Drawing",
	KindMicrosoftAccessDatabase:             "Microsoft Access Database",
	KindMicrosoftAnswerWizard:               "Microsoft Answer Wizard",
	KindMicrosoftCodePageTranslation:        "Microsoft Code Page Translation",
	KindMicrosoftCompress:                   "Microsoft Compress Archive",
	KindMicrosoftDeveloperStudioProject:     "Microsoft Developer Studio Project",
	KindMicrosoftExchangeConfig:             "Microsoft Exchange Config",
	KindMicrosoftInfo:                       "Microsoft Info",
	KindMicrosoftJournalDocument:            "Microsoft Journal Document",
	KindMicrosoftMoney:                      "Microsoft Money",
	KindMicrosoftNetworkMonitorCapture:      "Microsoft Network Monitor Capture",
	KindMicrosoftOneNoteNote:                "Microsoft OneNote Note",
	KindMicrosoftOutlookEmailFolder:         "Microsoft Outlook Email Folder",
	KindMicrosoftProgramDatabase:            "Microsoft Program Database",
	KindMicrosoftReaderDocument:             "Microsoft Reader Document",
	KindMicrosoftSQLServerDatabase:          "Microsoft SQL Server Database",
	KindMicrosoftWinMobileNoteDocument:      "Microsoft WinMobile Note Document",
	KindMicrosoftWorksSpreadsheet:           "Microsoft Works Spreadsheet",
	KindMicrosoftWriteDocument:              "Microsoft Write Document",
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
	KindMozillaArchive:                      "Mozilla Archive",
	KindMPEG2TransportStream:                "MPEG-2 Transport Stream",
	KindMPEGAudio:                           "MPEG Audio",
	KindMPEGAudioFrame:                      "MPEG Audio Frame",
	KindMPEGProgramStream:                   "MPEG Program Stream",
	KindMPEGTransportStream:                 "MPEG Transport Stream",
	KindMPEGVideo:                           "MPEG Video",
	KindMSDOSCOFFObject:                     "MS-DOS COFF Object",
	KindMUGENAudio:                          "M.U.G.E.N Audio",
	KindMUGENSprite:                         "M.U.G.E.N Sprite",
	KindMultiBitBlockchain:                  "MultiBit Blockchain",
	KindMultiBitInfo:                        "MultiBit Info",
	KindMultiBitWallet:                      "MultiBit Wallet",
	KindMusepackAudio:                       "Musepack Audio",
	KindMySQLMyISAMDatabase:                 "MySQL MyISAM Database",
	KindNationalImageryTransmission:         "National Imagery Transmission",
	KindNationalTransferFormat:              "National Transfer Format",
	KindNAVQuarantinedVirus:                 "NAV Quarantined Virus",
	KindNeoGeoPocketROM:                     "Neo Geo Pocket ROM",
	KindNeroCDCompilation:                   "Nero CD Compilation",
	KindNESAudio:                            "NES Audio",
	KindNESROM:                              "NES ROM",
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
	KindNintendoRARCArchive:                 "Nintendo RARC Archive",
	KindNintendoSwitchNROExecutable:         "Nintendo Switch NRO Executable",
	KindNintendoSwitchNSOExecutable:         "Nintendo Switch NSO Executable",
	KindNintendoSwitchPackage:               "Nintendo Switch Package",
	KindNintendoSwitchROM:                   "Nintendo Switch ROM",
	KindNintendoTPL:                         "Nintendo TPL Image",
	KindNintendoU8Archive:                   "Nintendo U8 Archive",
	KindNintendoWiiUExecutable:              "Nintendo Wii U Executable",
	KindNintendoYay0Archive:                 "Nintendo Yay0 Archive",
	KindNintendoYaz0Archive:                 "Nintendo Yaz0 Archive",
	KindNISTSPHEREAudio:                     "NIST SPHERE Audio",
	KindNOAARasterChart:                     "NOAA Raster Chart",
	KindNokiaPhoneBackup:                    "Nokia Phone Backup",
	KindNovellLANalyzerCapture:              "Novell LANalyzer Capture",
	KindNTFSFilesystem:                      "NTFS Filesystem",
	KindNullsoftVideo:                       "Nullsoft Video",
	KindNumPyArray:                          "NumPy Array",
	KindNVIDIASceneGraph:                    "NVIDIA Scene Graph",
	KindOCamlObject:                         "OCaml Object",
	KindOggContainer:                        "Ogg Container",
	KindOLECompoundDocument:                 "OLE Compound Document",
	KindOpenEXRImage:                        "OpenEXR Image",
	KindOpenSSHPrivateKey:                   "OpenSSH Private Key",
	KindOpenTypeFont:                        "OpenType Font",
	KindOptimFROGAudio:                      "OptimFROG Audio",
	KindORCColumnarData:                     "ORC Columnar Data",
	KindOutlookAddress:                      "Outlook Address",
	KindOutlookExpressAddressBook:           "Outlook Express Address Book",
	KindOutlookExpressDatabase:              "Outlook Express Database",
	KindPacketSnifferXCP:                    "Packet Sniffer XCP",
	KindPaintShopProImage:                   "PaintShop Pro Image",
	KindPalmData:                            "Palm Data",
	KindParallelsDiskImage:                  "Parallels Disk Image",
	KindParrotVideoEncapsulation:            "Parrot Video Encapsulation",
	KindPathWayMap:                          "PathWay Map",
	KindPAXImage:                            "PAX Image",
	KindPCAPCapture:                         "PCAP Capture",
	KindPCAPNGCapture:                       "PCAPNG Capture",
	KindPCFFont:                             "Portable Compiled Format Font",
	KindPCXDCXImage:                         "PCX/DCX Image",
	KindPCXImage:                            "PCX Image",
	KindPDFDocument:                         "PDF Document",
	KindPeaZipArchive:                       "PeaZip Archive",
	KindPEMCertificate:                      "PEM Certificate",
	KindPEMPrivateKey:                       "PEM Private Key",
	KindPerfectOfficeDocument:               "Perfect Office Document",
	KindPestPatrolData:                      "PestPatrol Data",
	KindPfaffEmbroidery:                     "Pfaff Embroidery",
	KindPGFImage:                            "PGF Image",
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
	KindPlayStation1Executable:              "PlayStation 1 Executable",
	KindPlayStation2MemoryCard:              "PlayStation 2 Memory Card",
	KindPlayStationAudio:                    "PlayStation Audio",
	KindPlayStationExecutable:               "PlayStation Executable",
	KindPlayStationPortableExecutable:       "PlayStation Portable Executable",
	KindPlayStationPortableISO:              "PlayStation Portable ISO",
	KindPLYModel:                            "PLY Model",
	KindPNGImage:                            "PNG Image",
	KindPortableExecutable:                  "Portable Executable",
	KindPostgreSQLCustomDump:                "PostgreSQL Custom Dump",
	KindPostScriptDocument:                  "PostScript Document",
	KindPowerBASICDebugger:                  "PowerBASIC Debugger",
	KindPowerBuilderIDE:                     "PowerBuilder IDE",
	KindPowerISODAA:                         "PowerISO DAA",
	KindProTrackerModule:                    "ProTracker Module",
	KindPRTGDatabase:                        "PRTG Database",
	KindPSFFont:                             "PC Screen Font",
	KindPsionDatabase:                       "Psion Database",
	KindPufferArchive:                       "Puffer Archive",
	KindPuttyPrivateKey:                     "PuTTY Private Key",
	KindPVRTextureImage:                     "PVR Texture Image",
	KindPythonBytecode:                      "Python Bytecode",
	KindPythonPickleData:                    "Python Pickle",
	KindQCOWDiskImage:                       "QCOW Disk Image",
	KindQDBMDatabase:                        "QDBM Database",
	KindQEMUQEDDiskImage:                    "QEMU QED Disk Image",
	KindQimageFilter:                        "Qimage Filter",
	KindQOIImage:                            "QOI Image",
	KindQuarkExpressDocument:                "Quark Express Document",
	KindQuattroProSpreadsheet:               "Quattro Pro Spreadsheet",
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
	KindROMFSFilesystem:                     "ROMFS Filesystem",
	KindROOTData:                            "ROOT Data",
	KindROSBag:                              "ROS Bag",
	KindRPGMakerArchive:                     "RPG Maker Archive",
	KindRPMPackage:                          "RPM Package",
	KindRSSFeed:                             "RSS Feed",
	KindRubyGemPackage:                      "RubyGem Package",
	KindRubyMarshalData:                     "Ruby Marshal Data",
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
	KindShareazaThumbnailImage:              "Shareaza Thumbnail",
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
	KindSNESSPCAudio:                        "SNES SPC Audio",
	KindSnoopCapture:                        "Snoop Capture",
	KindSOAPMessage:                         "SOAP Message",
	KindSonicFoundryAcid:                    "Sonic Foundry Acid",
	KindSonyCompressedVoice:                 "Sony Compressed Voice",
	KindSonyOpenMG:                          "Sony OpenMG Audio",
	KindSonyWave64Audio:                     "Sony Wave64 Audio",
	KindSourceEngineBSPMap:                  "Source Engine BSP Map",
	KindSpeedtouchFirmware:                  "Speedtouch Firmware",
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
	KindSubRipText:                          "SubRip Text",
	KindSunRasterImage:                      "Sun Raster Image",
	KindSuperCalcSpreadsheet:                "SuperCalc Spreadsheet",
	KindSurfplanProject:                     "Surfplan Project",
	KindSVGImage:                            "SVG Image",
	KindSymantecGhostDiskImage:              "Symantec Ghost Disk Image",
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
	KindTeraByteDiskImage:                   "TeraByte Disk Image",
	KindText:                                "Text",
	KindTheBatIndex:                         "The Bat! Index",
	KindThunderbirdMailSummary:              "Thunderbird Mail Summary",
	KindTIFFImage:                           "TIFF Image",
	KindTimezoneData:                        "Timezone Data",
	KindTokyoCabinetDatabase:                "Tokyo Cabinet Database",
	KindTomTomTrafficData:                   "TomTom Traffic",
	KindTorrent:                             "Torrent",
	KindTransportNeutralEncapsulationFormat: "Transport Neutral Encapsulation Format",
	KindTrueTypeCollection:                  "TrueType Collection",
	KindTrueTypeFont:                        "TrueType Font",
	KindTTAAudio:                            "TTA Audio",
	KindU3DModel:                            "U3D Model",
	KindUBIFSFilesystem:                     "UBIFS Filesystem",
	KindUBootImage:                          "U-Boot Image",
	KindUFAArchive:                          "UFA Archive",
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
	KindValveTextureFormatImage:             "Valve Texture Format Image",
	KindVBScriptEncoded:                     "VBScript Encoded",
	KindVCard:                               "vCard",
	KindVeeamBackup:                         "Veeam Backup",
	KindVHDDiskImage:                        "VHD Disk Image",
	KindVHDXDiskImage:                       "VHDX Disk Image",
	KindVideoCD:                             "Video VCD",
	KindVideoGameAudio:                      "Video Game Audio",
	KindVirtualBoxDiskImage:                 "VirtualBox Disk Image",
	KindVirtualBoxSavedState:                "VirtualBox Saved State",
	KindVisualCWorkbenchInfo:                "Visual C++ Workbench Info",
	KindVisualStudioSolution:                "Visual Studio Solution",
	KindVMapSourceWaypoint:                  "VMapSource Waypoint",
	KindVMwareDiskImage:                     "VMware Disk Image",
	KindVMwareNVRAM:                         "VMware NVRAM",
	KindVMwareSnapshot:                      "VMware Snapshot",
	KindVMwareSuspend:                       "VMware Suspend",
	KindVocaloidProject:                     "Vocaloid Project",
	KindVocalTecAudio:                       "VocalTec Audio",
	KindVOCAudio:                            "VOC Audio",
	KindVRML3DModel:                         "VRML 3D Model",
	KindVulkanSPIRV:                         "Vulkan SPIR-V",
	KindWADArchive:                          "WAD Archive",
	KindWalkmanMP3:                          "Walkman MP3",
	KindWARC:                                "WARC",
	KindWarcraftIIIReplay:                   "Warcraft III Replay",
	KindWavPackAudio:                        "WavPack Audio",
	KindWavPackCorrection:                   "WavPack Correction",
	KindWebAssemblyModule:                   "WebAssembly Module",
	KindWebexARFVideo:                       "Webex ARF",
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
	KindWinAmpPlaylist:                      "WinAmp Playlist",
	KindWindowsApplicationLog:               "Windows Application Log",
	KindWindowsCalendar:                     "Windows Calendar",
	KindWindowsEventLog:                     "Windows Event Log",
	KindWindowsEventViewer:                  "Windows Event Viewer",
	KindWindowsHelp:                         "Windows Help",
	KindWindowsHibernationData:              "Windows Hibernation Data",
	KindWindowsImagingFormat:                "Windows Imaging Format",
	KindWindowsInternationalCodePage:        "Windows International Code Page",
	KindWindowsMemoryDump:                   "Windows Memory Dump",
	KindWindowsMinidump:                     "Windows Minidump",
	KindWindowsPrecompiledHeader:            "Windows Precompiled Header",
	KindWindowsPrefetch:                     "Windows Prefetch",
	KindWindowsProgramManagerGroup:          "Windows Program Manager Group",
	KindWindowsRegistryHive:                 "Windows Registry Hive",
	KindWindowsResource:                     "Windows Resource",
	KindWindowsShortcut:                     "Windows Shortcut",
	KindWindowsThumbnailCache:               "Windows Thumbnail Cache",
	KindWindowsTypeLibrary:                  "Windows Type Library",
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
	KindWordPerfectGraphicsImage:            "WordPerfect Graphics Image",
	KindWordPerfectTextDocument:             "WordPerfect Text Document",
	KindWordStarDocument:                    "WordStar Document",
	KindWordStarWindows:                     "WordStar Windows",
	KindWTVVideo:                            "WTV Video",
	KindXARArchive:                          "XAR Archive",
	KindXBitMapImage:                        "X BitMap Image",
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
	KindZchunkArchive:                       "Zchunk Archive",
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
	TypeAPPXPackage:                        "APPX Package",
	TypeArchLinuxPackage:                   "Arch Linux Package",
	TypeASCII:                              "ASCII",
	TypeAVIFImage:                          "AVIF Image",
	TypeAVIFImageSequence:                  "AVIF Image Sequence",
	TypeAVIVideo:                           "AVI Video",
	TypeB4:                                 "B4",
	TypeBashScript:                         "Bash Script",
	TypeBatchScript:                        "Batch Script",
	TypeBGZF:                               "BGZF",
	TypeBigEndian:                          "Big-Endian",
	TypeBigTIFF:                            "BigTIFF",
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
	TypeCMakeScript:                        "CMake Script",
	TypeCodestream:                         "Codestream",
	TypeComicBook:                          "Comic Book",
	TypeCondaPackage:                       "Conda Package",
	TypeContainer:                          "Container",
	TypeCorelDRAWDocument:                  "CorelDRAW Document",
	TypeCOWD:                               "COWD",
	TypeCPPSource:                          "C++ Source",
	TypeCSharpSource:                       "C# Source",
	TypeCSource:                            "C Source",
	TypeCSS:                                "Cascading Style Sheets",
	TypeCubaseProject:                      "Cubase Project",
	TypeDartSource:                         "Dart Source",
	TypeDirectory:                          "Directory",
	TypeDockerfile:                         "Dockerfile",
	TypeDownloadableSounds:                 "Downloadable Sounds",
	TypeDSA:                                "DSA",
	TypeEAC3:                               "E-AC-3",
	TypeEC:                                 "EC",
	TypeEmpty:                              "Empty",
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
	TypeFLACAudio:                          "FLAC Audio",
	TypeForgeMod:                           "Forge Mod",
	TypeGIF87a:                             "GIF87a",
	TypeGIF89a:                             "GIF89a",
	TypeGoSource:                           "Go Source",
	TypeGraphQL:                            "GraphQL",
	TypeHDPhoto:                            "HD Photo",
	TypeHE:                                 "HE",
	TypeHEIFImage:                          "HEIF Image",
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
	TypeKDB:                                "KDB",
	TypeKDBX:                               "KDBX",
	TypeKDM:                                "KDM",
	TypeKerasModel:                         "Keras Model",
	TypeKMZArchive:                         "KMZ Archive",
	TypeKotlinSource:                       "Kotlin Source",
	TypeKritaDocument:                      "Krita Document",
	TypeLegacy:                             "Legacy",
	TypeLIST:                               "LIST",
	TypeLittleEndian:                       "Little-Endian",
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
	TypeNikonRAW:                                    "Nikon RAW",
	TypeNpmPackage:                                  "npm Package",
	TypeNuGetPackage:                                "NuGet Package",
	TypeOBJ:                                         "COFF Object",
	TypeObjectiveCSource:                            "Objective-C Source",
	TypeOCIImageLayout:                              "OCI Image Layout",
	TypeOldASCII:                                    "Old ASCII",
	TypeOlympusRAW:                                  "Olympus RAW",
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
	TypeRIFFPalette:                                 "Palette",
	TypeRSA:                                         "RSA",
	TypeRScript:                                     "R Script",
	TypeRubyScript:                                  "Ruby Script",
	TypeRustSource:                                  "Rust Source",
	TypeScalaSource:                                 "Scala Source",
	TypeSkinnableFrame:                              "Skinnable Frame",
	TypeSlackwarePackage:                            "Slackware Package",
	TypeSocket:                                      "Socket",
	TypeSonyRAW:                                     "Sony RAW",
	TypeSonyRAWSR2:                                  "Sony SR2 RAW",
	TypeSoundFont2:                                  "SoundFont 2",
	TypeSpecial:                                     "Special",
	TypeSpeexAudio:                                  "Speex Audio",
	TypeSQLScript:                                   "SQL Script",
	TypeSwiftSource:                                 "Swift Source",
	TypeSymbolicLink:                                "Symbolic Link",
	TypeTerraformConfiguration:                      "Terraform Configuration",
	TypeTheoraVideo:                                 "Theora Video",
	TypeTOMLConfiguration:                           "TOML Configuration",
	TypeTypeScriptSource:                            "TypeScript Source",
	TypeUncompressed:                                "Uncompressed",
	TypeUTF8:                                        "UTF-8",
	TypeVagrantBox:                                  "Vagrant Box",
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
	TypeVisualStudioExtension:                       "Visual Studio Extension",
	TypeVMDK:                                        "VMDK",
	TypeVMDKDescription:                             "VMDK Description",
	TypeVorbisAudio:                                 "Vorbis Audio",
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
