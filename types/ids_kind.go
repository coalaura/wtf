package types

type KindID uint16

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
	KindAdobeColorBook
	KindAdobeFontMetrics
	KindAdobeFrameMakerDocument
	KindAdobeInDesignDocument
	KindAdobeSwatchExchange
	KindAdvancedForensicFormat
	KindAINArchive
	KindAlembicModel
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
	KindAMRWBPlus
	KindAmstradCPCDiskImage
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
	KindApplePartitionMap
	KindApplePICTImage
	KindAppleSingleArchive
	KindAppleSystemLog
	KindAppleSystemProfiler
	KindAppleTextureImage
	KindAppleXMLPropertyList
	KindApproachIndex
	KindARArchive
	KindARCArchive
	KindARJArchive
	KindARRIRAWImage
	KindASFContainer
	KindASTCTextureImage
	KindAsylumMusicFormat
	KindAtari7800ROM
	KindAtariJaguarROM
	KindAtariLynxROM
	KindAtariTOSExecutable
	KindAtomFeed
	KindAUAudio
	KindAudacityBlock
	KindAutoCADDrawing
	KindAutoCADDXFDrawing
	KindAV1Video
	KindAVG6IntegrityDatabase
	KindAvidDNxHD
	KindAvroObjectContainer
	KindB1Archive
	KindBadgerDB
	KindBAMData
	KindBase85
	KindBasisUniversalImage
	KindBcachefsFilesystem
	KindBCFData
	KindBeFSFilesystem
	KindBerkeleyDatabase
	KindBethesdaArchive
	KindBethesdaPlugin
	KindBGBlitzDatabase
	KindBinHexArchive
	KindBinkVideo
	KindBitcoinBlock
	KindBitcoinWallet
	KindBitLockerDiskEncryption
	KindBlackBerryBackup
	KindBlenderProject
	KindBlinkArchive
	KindBlizzardM2Model
	KindBlizzardTextureImage
	KindBluetoothSnoopCapture
	KindBMFont
	KindBMPImage
	KindBochsDiskImage
	KindBoltDB
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
	KindCellebriteUFED
	KindCHMDocument
	KindChromaGraphImage
	KindChromiumPatch
	KindCineFormVideo
	KindCinema4DModel
	KindCineonImage
	KindClamAVDatabase
	KindClipStudioPaintImage
	KindCloneCDControl
	KindColladaModel
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
	KindCRIHCAAudio
	KindCRIWAREAudioWaveBank
	KindCRIWARECPKArchive
	KindCRIWAREVideo
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
	KindDialogicADPCM
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
	KindECWImage
	KindEFax
	KindEGGArchive
	KindElectronicArtsBigArchive
	KindElectronicArtsDBPFArchive
	KindElitePlusCommander
	KindEncapsulatedPostScript
	KindEnCaseCase
	KindEnCaseEvidenceV2
	KindEnCaseImage
	KindEndNoteLibrary
	KindENVIRemoteSensingHeader
	KindEOTFont
	KindErlangBEAMBytecode
	KindESRIShapefile
	KindEuropeanDataFormat
	KindEVRCAudio
	KindExecutableAndLinkableFormat
	KindExFATFilesystem
	KindExpertWitnessFormat
	KindExtensibleMusicFormat
	KindExtFilesystem
	KindF2FSFilesystem
	KindFamicomDiskSystemROM
	KindFarandoleComposerModule
	KindFarbfeldImage
	KindFastTracker2ExtendedInstrument
	KindFastTrackerModule
	KindFATFilesystem
	KindFBXModel
	KindFeatherData
	KindFiascoDatabase
	KindFileMakerProDatabase
	KindFilesystemEntry
	KindFinaleNotationFile
	KindFirebirdDatabase
	KindFITActivity
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLICAnimation
	KindFLIFImage
	KindFlightSimulatorConfig
	KindFlowCytometryStandard
	KindFLStudioProject
	KindFLVVideo
	KindFMODEvent
	KindFMODSampleBank
	KindFreeArcArchive
	KindFujifilmRAWImage
	KindFuzzyBitmapImage
	KindGameBoyAdvanceROM
	KindGameBoyAudio
	KindGameBoyROM
	KindGamebryoModel
	KindGameCubeROM
	KindGameMakerData
	KindGDBMDatabase
	KindGDSII
	KindGEMRasterImage
	KindGenealogicalData
	KindGenetecVideo
	KindGeographyMarkupLanguage
	KindGerberFormat
	KindGHCiInterface
	KindGIFImage
	KindGIMPBrush
	KindGIMPPalette
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
	KindGuitarProTablature
	KindGzipArchive
	KindGzipData
	KindH264Video
	KindH265Video
	KindHadoopRC
	KindHadoopSequence
	KindHamarsoftArchive
	KindHarvardGraphicsImage
	KindHDF4Data
	KindHDF5Data
	KindHealthLevel7
	KindHFSPlusFilesystem
	KindHighEntropyData
	KindHTMLDocument
	KindHuskygramPoem
	KindHusqvarnaDesigner
	KindICalendar
	KindICCProfile
	KindICNSIcon
	KindICOCURImage
	KindIdRoQVideo
	KindIdSoftwareBSPMap
	KindIEHistory
	KindIFFContainer
	KindImgSoftwareImage
	KindImpulseTrackerInstrument
	KindImpulseTrackerModule
	KindImpulseTrackerSample
	KindInfluxDBTSM
	KindInnoSetupLog
	KindInstallShieldCabinet
	KindInstallShieldScript
	KindIntellivisionROM
	KindInterplayMVEVideo
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
	KindJFSFilesystem
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
	KindKodakPhotoCD
	KindKorgAudio
	KindKTXTextureImage
	KindKyotoCabinetDatabase
	KindLevelDB
	KindLHAArchive
	KindLidarData
	KindLightWaveScene
	KindLinuxKernelImage
	KindLinuxSwapSpace
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
	KindLUKS2DiskEncryption
	KindLUKSDiskEncryption
	KindLuraDocument
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
	KindMacOSEventsLog
	KindMacOSXCodeSignature
	KindMacPaintImage
	KindMacriumReflectDiskImage
	KindMadTracker2Module
	KindMagicaVoxelModel
	KindMAMECHDDiskImage
	KindMapInfoMIF
	KindMArArchive
	KindMarmosetToolbagScene
	KindMasterBootRecord
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
	KindMilesSoundSystem
	KindMilestonesProject
	KindMinixFilesystem
	KindMinoltaRAWImage
	KindMNGImage
	KindMOBIDocument
	KindMoneroWallet
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
	KindMrSIDImage
	KindMSDOSCOFFObject
	KindMSXCartridge
	KindMUGENAudio
	KindMUGENSprite
	KindMultiBitBlockchain
	KindMultiBitInfo
	KindMultiBitWallet
	KindMultiTrackerModule
	KindMusepackAudio
	KindMySQLBinlog
	KindMySQLMyISAMDatabase
	KindNASASPICEKernel
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
	KindNintendoContentArchive
	KindNintendoDSiROM
	KindNintendoDSROM
	KindNintendoRARCArchive
	KindNintendoSwitchKIPExecutable
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
	KindNRRDMedicalImage
	KindNTFSFilesystem
	KindNullsoftVideo
	KindNumPyArray
	KindNVIDIASceneGraph
	KindOCamlObject
	KindOggContainer
	KindOktalyzerModule
	KindOLECompoundDocument
	KindOpenEXRImage
	KindOpenSSHPrivateKey
	KindOpenStreetMapPBF
	KindOpenTypeFont
	KindOpenVDBModel
	KindOptimFROGAudio
	KindORCColumnarData
	KindOutlookAddress
	KindOutlookExpressAddressBook
	KindOutlookExpressDatabase
	KindPacketSnifferXCP
	KindPaintDotNetImage
	KindPaintShopProImage
	KindPaintToolSAIImage
	KindPalmData
	KindPalmOSApplication
	KindPanasonic3DOROM
	KindParallelsDiskImage
	KindParrotVideoEncapsulation
	KindPasswordSafeDatabase
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
	KindPlayStationPackage
	KindPlayStationPortableExecutable
	KindPlayStationPortableISO
	KindPLYModel
	KindPMarcArchive
	KindPMTiles
	KindPNGImage
	KindPolyTrackerModule
	KindPortableExecutable
	KindPortableFontResource
	KindPostgreSQLCustomDump
	KindPostScriptDocument
	KindPowerBASICDebugger
	KindPowerBuilderIDE
	KindPowerISODAA
	KindProteinDataBank
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
	KindQNX4Filesystem
	KindQNX6Filesystem
	KindQOIImage
	KindQuakePakArchive
	KindQuarkExpressDocument
	KindQuattroProSpreadsheet
	KindQuickBooksBackup
	KindQuickenData
	KindQuickenQuickFinder
	KindQuickReport
	KindRadianceHDRImage
	KindRagePackageFormat
	KindRagTimeDocument
	KindRARArchive
	KindRData
	KindRealMedia
	KindRealMediaMetafile
	KindRealPlayerVideo
	KindReaperProject
	KindRedisDatabase
	KindREDRAWImage
	KindReFSFilesystem
	KindReiserFSFilesystem
	KindRelicSGAArchive
	KindRenPyArchive
	KindRetroArchSavestate
	KindRhino3DModel
	KindRichTextFormatDocument
	KindRIFFContainer
	KindRiveAnimation
	KindRKAudio
	KindRNCArchive
	KindRobloxModel
	KindRolandSVQAudio
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
	KindSeattleFilmWorksImage
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
	KindSITXArchive
	KindSketchDocument
	KindSketchUpModel
	KindSkinCrafterSkin
	KindSkypeData
	KindSkypeLocalization
	KindSmackerVideo
	KindSmartDrawDrawing
	KindSnappyFramedData
	KindSnes9xSavestate
	KindSNESSPCAudio
	KindSnoopCapture
	KindSOAPMessage
	KindSoftimagePIC
	KindSonicFoundryAcid
	KindSonyCompressedVoice
	KindSonyOpenMG
	KindSonyWave64Audio
	KindSoundDesignerII
	KindSourceEngineBSPMap
	KindSourceEngineDemo
	KindSourceEngineModel
	KindSpeedtouchFirmware
	KindSPIFFImage
	KindSPSSData
	KindSPSSPortableData
	KindSPSSTemplate
	KindSQArchive
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
	KindTwinVQAudio
	KindU3DModel
	KindUBIFSFilesystem
	KindUBootImage
	KindUFAArchive
	KindUFSFilesystem
	KindUHAArchive
	KindUltraTrackerModule
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
	KindVectrexROM
	KindVeeamBackup
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVideoCD
	KindVideoGameAudio
	KindVirtualBoxDiskImage
	KindVirtualBoxSavedState
	KindVisualBoyAdvanceSavestate
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
	KindVxFSFilesystem
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
	KindWiiUCompressedDiskImage
	KindWiiUDiskImage
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
	KindWwisePackage
	KindWwiseSoundBank
	KindX3DModel
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
	KindXWindowDumpImage
	KindXZArchive
	KindYACArchive
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
	KindZSNESSavestate
	KindZstandardArchive
	KindZstandardDictionary
	KindZXTape
	MaxKinds
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
	KindAdobeColorBook:                      "Adobe Color Book",
	KindAdobeFontMetrics:                    "Adobe Font Metrics",
	KindAdobeFrameMakerDocument:             "Adobe FrameMaker Document",
	KindAdobeInDesignDocument:               "Adobe InDesign Document",
	KindAdobeSwatchExchange:                 "Adobe Swatch Exchange",
	KindAdvancedForensicFormat:              "Advanced Forensic Format",
	KindAINArchive:                          "AIN Archive",
	KindAlembicModel:                        "Alembic 3D Model",
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
	KindAMRWBPlus:                           "AMR-WB+ Audio",
	KindAmstradCPCDiskImage:                 "Amstrad CPC Disk Image",
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
	KindApplePartitionMap:                   "Apple Partition Map",
	KindApplePICTImage:                      "Apple PICT Image",
	KindAppleSingleArchive:                  "AppleSingle Archive",
	KindAppleSystemLog:                      "Apple System Log",
	KindAppleSystemProfiler:                 "Apple System Profiler",
	KindAppleTextureImage:                   "Apple Texture Image",
	KindAppleXMLPropertyList:                "Apple XML Property List",
	KindApproachIndex:                       "Approach Index",
	KindARArchive:                           "AR Archive",
	KindARCArchive:                          "ARC Archive",
	KindARJArchive:                          "ARJ Archive",
	KindARRIRAWImage:                        "ARRI RAW Image",
	KindASFContainer:                        "ASF Container",
	KindASTCTextureImage:                    "ASTC Texture Image",
	KindAsylumMusicFormat:                   "Asylum Music Format",
	KindAtari7800ROM:                        "Atari 7800 ROM",
	KindAtariJaguarROM:                      "Atari Jaguar ROM",
	KindAtariLynxROM:                        "Atari Lynx ROM",
	KindAtariTOSExecutable:                  "Atari TOS Executable",
	KindAtomFeed:                            "Atom Feed",
	KindAUAudio:                             "AU Audio",
	KindAudacityBlock:                       "Audacity Block",
	KindAutoCADDrawing:                      "AutoCAD Drawing",
	KindAutoCADDXFDrawing:                   "AutoCAD DXF Drawing",
	KindAV1Video:                            "AV1 Video",
	KindAVG6IntegrityDatabase:               "AVG6 Integrity Database",
	KindAvidDNxHD:                           "Avid DNxHD Video",
	KindAvroObjectContainer:                 "Avro Object Container",
	KindB1Archive:                           "B1 Archive",
	KindBadgerDB:                            "BadgerDB Database",
	KindBAMData:                             "BAM Data",
	KindBase85:                              "BASE85",
	KindBasisUniversalImage:                 "Basis Universal Image",
	KindBcachefsFilesystem:                  "Bcachefs Filesystem",
	KindBCFData:                             "Binary Call Format",
	KindBeFSFilesystem:                      "BeOS Filesystem",
	KindBerkeleyDatabase:                    "Berkeley Database",
	KindBethesdaArchive:                     "Bethesda Archive",
	KindBethesdaPlugin:                      "Bethesda Plugin",
	KindBGBlitzDatabase:                     "BGBlitz Database",
	KindBinHexArchive:                       "BinHex Archive",
	KindBinkVideo:                           "Bink Video",
	KindBitcoinBlock:                        "Bitcoin Block Data",
	KindBitcoinWallet:                       "Bitcoin Wallet",
	KindBitLockerDiskEncryption:             "BitLocker Disk Encryption",
	KindBlackBerryBackup:                    "BlackBerry Backup",
	KindBlenderProject:                      "Blender",
	KindBlinkArchive:                        "Blink Archive",
	KindBlizzardM2Model:                     "Blizzard M2 Model",
	KindBlizzardTextureImage:                "Blizzard Texture Image",
	KindBluetoothSnoopCapture:               "Bluetooth Snoop Capture",
	KindBMFont:                              "AngelCode Bitmap Font",
	KindBMPImage:                            "BMP Image",
	KindBochsDiskImage:                      "Bochs Disk Image",
	KindBoltDB:                              "BoltDB Database",
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
	KindCellebriteUFED:                      "Cellebrite UFED Extraction",
	KindCHMDocument:                         "Compiled HTML Help",
	KindChromaGraphImage:                    "ChromaGraph Image",
	KindChromiumPatch:                       "Chromium Patch",
	KindCineFormVideo:                       "CineForm Video",
	KindCinema4DModel:                       "Cinema 4D Model",
	KindCineonImage:                         "Cineon Image",
	KindClamAVDatabase:                      "ClamAV Database",
	KindClipStudioPaintImage:                "Clip Studio Paint Image",
	KindCloneCDControl:                      "CloneCD Control",
	KindColladaModel:                        "COLLADA Model",
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
	KindCRIHCAAudio:                         "CRI HCA Audio",
	KindCRIWAREAudioWaveBank:                "CRIWARE Audio Wave Bank",
	KindCRIWARECPKArchive:                   "CRIWARE CPK Archive",
	KindCRIWAREVideo:                        "CRIWARE Video",
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
	KindDialogicADPCM:                       "Dialogic ADPCM Audio",
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
	KindECWImage:                            "ECW Image",
	KindEFax:                                "eFax",
	KindEGGArchive:                          "EGG Archive",
	KindElectronicArtsBigArchive:            "Electronic Arts BIG Archive",
	KindElectronicArtsDBPFArchive:           "Electronic Arts DBPF Archive",
	KindElitePlusCommander:                  "Elite Plus Commander",
	KindEncapsulatedPostScript:              "Encapsulated PostScript",
	KindEnCaseCase:                          "EnCase Case",
	KindEnCaseEvidenceV2:                    "EnCase Evidence V2",
	KindEnCaseImage:                         "EnCase Image",
	KindEndNoteLibrary:                      "EndNote Library",
	KindENVIRemoteSensingHeader:             "ENVI Remote Sensing Header",
	KindEOTFont:                             "EOT Font",
	KindErlangBEAMBytecode:                  "Erlang BEAM Bytecode",
	KindESRIShapefile:                       "ESRI Shapefile",
	KindEuropeanDataFormat:                  "European Data Format",
	KindEVRCAudio:                           "EVRC Audio",
	KindExecutableAndLinkableFormat:         "Executable and Linkable Format",
	KindExFATFilesystem:                     "exFAT Filesystem",
	KindExpertWitnessFormat:                 "Expert Witness Format",
	KindExtensibleMusicFormat:               "Extensible Music Format",
	KindExtFilesystem:                       "ext Filesystem",
	KindF2FSFilesystem:                      "F2FS Filesystem",
	KindFamicomDiskSystemROM:                "Famicom Disk System ROM",
	KindFarandoleComposerModule:             "Farandole Composer Module",
	KindFarbfeldImage:                       "Farbfeld Image",
	KindFastTracker2ExtendedInstrument:      "FastTracker 2 Extended Instrument",
	KindFastTrackerModule:                   "FastTracker Module",
	KindFATFilesystem:                       "FAT Filesystem",
	KindFBXModel:                            "FBX Model",
	KindFeatherData:                         "Feather Data",
	KindFiascoDatabase:                      "Fiasco Database",
	KindFileMakerProDatabase:                "FileMaker Pro Database",
	KindFilesystemEntry:                     "Filesystem Entry",
	KindFinaleNotationFile:                  "Finale Notation File",
	KindFirebirdDatabase:                    "Firebird Database",
	KindFITActivity:                         "FIT Activity",
	KindFITSAstronomicalImage:               "FITS Astronomical Image",
	KindFLACAudio:                           "FLAC Audio",
	KindFLICAnimation:                       "FLIC Animation",
	KindFLIFImage:                           "FLIF Image",
	KindFlightSimulatorConfig:               "Flight Simulator Config",
	KindFlowCytometryStandard:               "Flow Cytometry Standard",
	KindFLStudioProject:                     "FL Studio Project",
	KindFLVVideo:                            "FLV Video",
	KindFMODEvent:                           "FMOD Event",
	KindFMODSampleBank:                      "FMOD Sample Bank",
	KindFreeArcArchive:                      "FreeArc Archive",
	KindFujifilmRAWImage:                    "Fujifilm RAW Image",
	KindFuzzyBitmapImage:                    "Fuzzy Bitmap Image",
	KindGameBoyAdvanceROM:                   "Game Boy Advance ROM",
	KindGameBoyAudio:                        "Game Boy Audio",
	KindGameBoyROM:                          "Game Boy ROM",
	KindGamebryoModel:                       "Gamebryo Model",
	KindGameCubeROM:                         "GameCube ROM",
	KindGameMakerData:                       "GameMaker Data",
	KindGDBMDatabase:                        "GDBM Database",
	KindGDSII:                               "GDSII IC Layout",
	KindGEMRasterImage:                      "GEM Raster Image",
	KindGenealogicalData:                    "Genealogical Data",
	KindGenetecVideo:                        "Genetec Video",
	KindGeographyMarkupLanguage:             "Geography Markup Language",
	KindGerberFormat:                        "Gerber Format",
	KindGHCiInterface:                       "GHCi Interface",
	KindGIFImage:                            "GIF Image",
	KindGIMPBrush:                           "GIMP Brush",
	KindGIMPPalette:                         "GIMP Palette",
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
	KindGuitarProTablature:                  "Guitar Pro Tablature",
	KindGzipArchive:                         "Gzip Archive",
	KindGzipData:                            "Gzip Data",
	KindH264Video:                           "H.264 Video",
	KindH265Video:                           "H.265 Video",
	KindHadoopRC:                            "Hadoop RC",
	KindHadoopSequence:                      "Hadoop Sequence",
	KindHamarsoftArchive:                    "Hamarsoft Archive",
	KindHarvardGraphicsImage:                "Harvard Graphics Image",
	KindHDF4Data:                            "HDF4 Data",
	KindHDF5Data:                            "HDF5 Data",
	KindHealthLevel7:                        "Health Level-7 Data",
	KindHFSPlusFilesystem:                   "HFS+ Filesystem",
	KindHighEntropyData:                     "High Entropy Data",
	KindHTMLDocument:                        "HTML Document",
	KindHuskygramPoem:                       "Huskygram Poem",
	KindHusqvarnaDesigner:                   "Husqvarna Designer",
	KindICalendar:                           "iCalendar",
	KindICCProfile:                          "ICC Profile",
	KindICNSIcon:                            "ICNS Icon",
	KindICOCURImage:                         "ICO/CUR Image",
	KindIdRoQVideo:                          "Id Software RoQ Video",
	KindIdSoftwareBSPMap:                    "Id Software BSP Map",
	KindIEHistory:                           "IE History",
	KindIFFContainer:                        "IFF Container",
	KindImgSoftwareImage:                    "Img Software Image",
	KindImpulseTrackerInstrument:            "Impulse Tracker Instrument",
	KindImpulseTrackerModule:                "Impulse Tracker Module",
	KindImpulseTrackerSample:                "Impulse Tracker Sample",
	KindInfluxDBTSM:                         "InfluxDB TSM",
	KindInnoSetupLog:                        "Inno Setup Log",
	KindInstallShieldCabinet:                "InstallShield Cabinet",
	KindInstallShieldScript:                 "InstallShield Script",
	KindIntellivisionROM:                    "Intellivision ROM",
	KindInterplayMVEVideo:                   "Interplay MVE Video",
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
	KindJFSFilesystem:                       "JFS Filesystem",
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
	KindKodakPhotoCD:                        "Kodak PhotoCD",
	KindKorgAudio:                           "Korg Audio",
	KindKTXTextureImage:                     "KTX Texture Image",
	KindKyotoCabinetDatabase:                "Kyoto Cabinet Database",
	KindLevelDB:                             "LevelDB / RocksDB",
	KindLHAArchive:                          "LHA Archive",
	KindLidarData:                           "LiDAR Data",
	KindLightWaveScene:                      "LightWave Scene",
	KindLinuxKernelImage:                    "Linux Kernel Image",
	KindLinuxSwapSpace:                      "Linux Swap Space",
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
	KindLUKS2DiskEncryption:                 "LUKS2 Disk Encryption",
	KindLUKSDiskEncryption:                  "LUKS Disk Encryption",
	KindLuraDocument:                        "LuraDocument",
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
	KindMacOSEventsLog:                      "macOS FSEvents Log",
	KindMacOSXCodeSignature:                 "macOS Code Signature",
	KindMacPaintImage:                       "MacPaint Image",
	KindMacriumReflectDiskImage:             "Macrium Reflect Disk Image",
	KindMadTracker2Module:                   "MadTracker 2 Module",
	KindMagicaVoxelModel:                    "MagicaVoxel Model",
	KindMAMECHDDiskImage:                    "MAME CHD Disk Image",
	KindMapInfoMIF:                          "MapInfo MIF",
	KindMArArchive:                          "MAr Archive",
	KindMarmosetToolbagScene:                "Marmoset Toolbag Scene",
	KindMasterBootRecord:                    "Master Boot Record",
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
	KindMilesSoundSystem:                    "Miles Sound System Audio",
	KindMilestonesProject:                   "Milestones Project",
	KindMinixFilesystem:                     "Minix Filesystem",
	KindMinoltaRAWImage:                     "Minolta RAW Image",
	KindMNGImage:                            "MNG Image",
	KindMOBIDocument:                        "MOBI Document",
	KindMoneroWallet:                        "Monero Wallet",
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
	KindMrSIDImage:                          "MrSID Image",
	KindMSDOSCOFFObject:                     "MS-DOS COFF Object",
	KindMSXCartridge:                        "MSX Cartridge",
	KindMUGENAudio:                          "M.U.G.E.N Audio",
	KindMUGENSprite:                         "M.U.G.E.N Sprite",
	KindMultiBitBlockchain:                  "MultiBit Blockchain",
	KindMultiBitInfo:                        "MultiBit Info",
	KindMultiBitWallet:                      "MultiBit Wallet",
	KindMultiTrackerModule:                  "MultiTracker Module",
	KindMusepackAudio:                       "Musepack Audio",
	KindMySQLBinlog:                         "MySQL Binary Log",
	KindMySQLMyISAMDatabase:                 "MySQL MyISAM Database",
	KindNASASPICEKernel:                     "NASA SPICE Kernel",
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
	KindNintendoContentArchive:              "Nintendo Content Archive",
	KindNintendoDSiROM:                      "Nintendo DSi ROM",
	KindNintendoDSROM:                       "Nintendo DS ROM",
	KindNintendoRARCArchive:                 "Nintendo RARC Archive",
	KindNintendoSwitchKIPExecutable:         "Nintendo Switch KIP Executable",
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
	KindNRRDMedicalImage:                    "NRRD Medical Image",
	KindNTFSFilesystem:                      "NTFS Filesystem",
	KindNullsoftVideo:                       "Nullsoft Video",
	KindNumPyArray:                          "NumPy Array",
	KindNVIDIASceneGraph:                    "NVIDIA Scene Graph",
	KindOCamlObject:                         "OCaml Object",
	KindOggContainer:                        "Ogg Container",
	KindOktalyzerModule:                     "Oktalyzer Module",
	KindOLECompoundDocument:                 "OLE Compound Document",
	KindOpenEXRImage:                        "OpenEXR Image",
	KindOpenSSHPrivateKey:                   "OpenSSH Private Key",
	KindOpenStreetMapPBF:                    "OpenStreetMap PBF",
	KindOpenTypeFont:                        "OpenType Font",
	KindOpenVDBModel:                        "OpenVDB Model",
	KindOptimFROGAudio:                      "OptimFROG Audio",
	KindORCColumnarData:                     "ORC Columnar Data",
	KindOutlookAddress:                      "Outlook Address",
	KindOutlookExpressAddressBook:           "Outlook Express Address Book",
	KindOutlookExpressDatabase:              "Outlook Express Database",
	KindPacketSnifferXCP:                    "Packet Sniffer XCP",
	KindPaintDotNetImage:                    "Paint.NET Image",
	KindPaintShopProImage:                   "PaintShop Pro Image",
	KindPaintToolSAIImage:                   "PaintTool SAI Image",
	KindPalmData:                            "Palm Data",
	KindPalmOSApplication:                   "Palm OS Application",
	KindPanasonic3DOROM:                     "Panasonic 3DO ROM",
	KindParallelsDiskImage:                  "Parallels Disk Image",
	KindParrotVideoEncapsulation:            "Parrot Video Encapsulation",
	KindPasswordSafeDatabase:                "Password Safe Database",
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
	KindPlayStationPackage:                  "PlayStation Package",
	KindPlayStationPortableExecutable:       "PlayStation Portable Executable",
	KindPlayStationPortableISO:              "PlayStation Portable ISO",
	KindPLYModel:                            "PLY Model",
	KindPMarcArchive:                        "PMarc Archive",
	KindPMTiles:                             "PMTiles",
	KindPNGImage:                            "PNG Image",
	KindPolyTrackerModule:                   "PolyTracker Module",
	KindPortableExecutable:                  "Portable Executable",
	KindPortableFontResource:                "Portable Font Resource",
	KindPostgreSQLCustomDump:                "PostgreSQL Custom Dump",
	KindPostScriptDocument:                  "PostScript Document",
	KindPowerBASICDebugger:                  "PowerBASIC Debugger",
	KindPowerBuilderIDE:                     "PowerBuilder IDE",
	KindPowerISODAA:                         "PowerISO DAA",
	KindProteinDataBank:                     "Protein Data Bank",
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
	KindQNX4Filesystem:                      "QNX4 Filesystem",
	KindQNX6Filesystem:                      "QNX6 Filesystem",
	KindQOIImage:                            "QOI Image",
	KindQuakePakArchive:                     "Quake PAK Archive",
	KindQuarkExpressDocument:                "Quark Express Document",
	KindQuattroProSpreadsheet:               "Quattro Pro Spreadsheet",
	KindQuickBooksBackup:                    "QuickBooks Backup",
	KindQuickenData:                         "Quicken Data",
	KindQuickenQuickFinder:                  "Quicken QuickFinder",
	KindQuickReport:                         "QuickReport",
	KindRadianceHDRImage:                    "Radiance HDR Image",
	KindRagePackageFormat:                   "RAGE Package Format",
	KindRagTimeDocument:                     "RagTime Document",
	KindRARArchive:                          "RAR Archive",
	KindRData:                               "RData",
	KindRealMedia:                           "RealMedia",
	KindRealMediaMetafile:                   "RealMedia Metafile",
	KindRealPlayerVideo:                     "RealPlayer Video",
	KindReaperProject:                       "Reaper Project",
	KindRedisDatabase:                       "Redis Database",
	KindREDRAWImage:                         "RED RAW Image",
	KindReFSFilesystem:                      "ReFS Filesystem",
	KindReiserFSFilesystem:                  "ReiserFS Filesystem",
	KindRelicSGAArchive:                     "Relic SGA Archive",
	KindRenPyArchive:                        "Ren'Py Archive",
	KindRetroArchSavestate:                  "RetroArch Savestate",
	KindRhino3DModel:                        "Rhino 3D Model",
	KindRichTextFormatDocument:              "Rich Text Format Document",
	KindRIFFContainer:                       "RIFF Container",
	KindRiveAnimation:                       "Rive Animation",
	KindRKAudio:                             "RKAudio",
	KindRNCArchive:                          "Rob Northen Compression Archive",
	KindRobloxModel:                         "Roblox Model",
	KindRolandSVQAudio:                      "Roland SVQ Audio",
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
	KindSeattleFilmWorksImage:               "Seattle FilmWorks Image",
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
	KindSITXArchive:                         "StuffIt X Archive",
	KindSketchDocument:                      "Sketch Document",
	KindSketchUpModel:                       "SketchUp Model",
	KindSkinCrafterSkin:                     "SkinCrafter Skin",
	KindSkypeData:                           "Skype Data",
	KindSkypeLocalization:                   "Skype Localization",
	KindSmackerVideo:                        "Smacker Video",
	KindSmartDrawDrawing:                    "SmartDraw Drawing",
	KindSnappyFramedData:                    "Snappy Framed Data",
	KindSnes9xSavestate:                     "Snes9x Savestate",
	KindSNESSPCAudio:                        "SNES SPC Audio",
	KindSnoopCapture:                        "Snoop Capture",
	KindSOAPMessage:                         "SOAP Message",
	KindSoftimagePIC:                        "Softimage PIC",
	KindSonicFoundryAcid:                    "Sonic Foundry Acid",
	KindSonyCompressedVoice:                 "Sony Compressed Voice",
	KindSonyOpenMG:                          "Sony OpenMG Audio",
	KindSonyWave64Audio:                     "Sony Wave64 Audio",
	KindSoundDesignerII:                     "Sound Designer II Audio",
	KindSourceEngineBSPMap:                  "Source Engine BSP Map",
	KindSourceEngineDemo:                    "Source Engine Demo",
	KindSourceEngineModel:                   "Source Engine Model",
	KindSpeedtouchFirmware:                  "Speedtouch Firmware",
	KindSPIFFImage:                          "SPIFF Image",
	KindSPSSData:                            "SPSS Data",
	KindSPSSPortableData:                    "SPSS Portable Data",
	KindSPSSTemplate:                        "SPSS Template",
	KindSQArchive:                           "Squeeze Archive",
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
	KindTwinVQAudio:                         "TwinVQ Audio",
	KindU3DModel:                            "U3D Model",
	KindUBIFSFilesystem:                     "UBIFS Filesystem",
	KindUBootImage:                          "U-Boot Image",
	KindUFAArchive:                          "UFA Archive",
	KindUFSFilesystem:                       "UFS Filesystem",
	KindUHAArchive:                          "UHA Archive",
	KindUltraTrackerModule:                  "UltraTracker Module",
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
	KindVectrexROM:                          "Vectrex ROM",
	KindVeeamBackup:                         "Veeam Backup",
	KindVHDDiskImage:                        "VHD Disk Image",
	KindVHDXDiskImage:                       "VHDX Disk Image",
	KindVideoCD:                             "Video VCD",
	KindVideoGameAudio:                      "Video Game Audio",
	KindVirtualBoxDiskImage:                 "VirtualBox Disk Image",
	KindVirtualBoxSavedState:                "VirtualBox Saved State",
	KindVisualBoyAdvanceSavestate:           "VisualBoyAdvance Savestate",
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
	KindVxFSFilesystem:                      "VxFS Filesystem",
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
	KindWiiUCompressedDiskImage:             "Wii U Compressed Disk Image",
	KindWiiUDiskImage:                       "Wii U Disk Image",
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
	KindWwisePackage:                        "Wwise Package",
	KindWwiseSoundBank:                      "Wwise SoundBank",
	KindX3DModel:                            "X3D Model",
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
	KindXWindowDumpImage:                    "X Window Dump Image",
	KindXZArchive:                           "XZ Archive",
	KindYACArchive:                          "YAC Archive",
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
	KindZSNESSavestate:                      "ZSNES Savestate",
	KindZstandardArchive:                    "Zstandard Archive",
	KindZstandardDictionary:                 "Zstandard Dictionary",
	KindZXTape:                              "ZX Spectrum Tape",
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
