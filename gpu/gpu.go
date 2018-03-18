package gpu

//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gpu
//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gpu
//#include <stdlib.h>
//#include "sdl_gpu_wrapper.h"
import "C"
import (
	"unsafe"
)

const GPUFalse = 0
const GPUTrue = 1

type Color struct {
	R, G, B, A uint8
}

type cColor C.SDL_Color

type Rect struct {
	X, Y, W, H float32
}

func (r *Rect) cptr() *C.GPU_Rect {
	return (*C.GPU_Rect)(unsafe.Pointer(r))
}

type cRect C.GPU_Rect

type RendererID C.GPU_RendererID

const RendererOrderMax int32 = 10

type RendererEnum C.Uint32

const (
	RendererUnknown     RendererEnum = 0
	RendererOpenGL1Base RendererEnum = 1
	RendererOpenGL1     RendererEnum = 2
	RendererOpenGL2     RendererEnum = 3
	RendererOpenGL3     RendererEnum = 4
	RendererOpenGL4     RendererEnum = 5
	RendererGLES1       RendererEnum = 11
	RendererGLES2       RendererEnum = 12
	RendererGlES3       RendererEnum = 13
	RendererD3D9        RendererEnum = 21
	RendererD3D10       RendererEnum = 22
	RendererD3D11       RendererEnum = 23
)
const (
	ComparisonNever    = C.GPU_NEVER
	ComparisonLess     = C.GPU_LESS
	ComparisonEqual    = C.GPU_EQUAL
	ComparisonLEqual   = C.GPU_LEQUAL
	ComparisonGreater  = C.GPU_GREATER
	ComparisonNotEqual = C.GPU_NOTEQUAL
	ComparisonGEqual   = C.GPU_GEQUAL
	ComparisonAlways   = C.GPU_ALWAYS
)

const (
	BlendFuncZero             = C.GPU_FUNC_ZERO
	BlendFuncOne              = C.GPU_FUNC_ONE
	BlendFuncSrcColor         = C.GPU_FUNC_SRC_COLOR
	BlendFuncDstColor         = C.GPU_FUNC_DST_COLOR
	BlendFuncOneMinusSrc      = C.GPU_FUNC_ONE_MINUS_SRC
	BlendFuncOneMinusDst      = C.GPU_FUNC_ONE_MINUS_DST
	BlendFuncSrcAlpha         = C.GPU_FUNC_SRC_ALPHA
	BlendFuncDstAlpha         = C.GPU_FUNC_DST_ALPHA
	BlendFuncOneMinusSrcAlpha = C.GPU_FUNC_ONE_MINUS_SRC_ALPHA
	BlendFuncOneMinusDstAlpha = C.GPU_FUNC_ONE_MINUS_DST_ALPHA
)

const (
	BlendEqAdd             = C.GPU_EQ_ADD
	BlendEqSubtract        = C.GPU_EQ_SUBTRACT
	BlendEqReverseSubtract = C.GPU_EQ_REVERSE_SUBTRACT
)

type BlendPresetEnum C.GPU_BlendPresetEnum

const (
	BlendPresetNormal             BlendPresetEnum = C.GPU_BLEND_NORMAL
	BlendPresetPremultipliedAlpha BlendPresetEnum = C.GPU_BLEND_PREMULTIPLIED_ALPHA
	BlendPresetMultiply           BlendPresetEnum = C.GPU_BLEND_MULTIPLY
	BlendPresetAdd                BlendPresetEnum = C.GPU_BLEND_ADD
	BlendPresetSubtract           BlendPresetEnum = C.GPU_BLEND_SUBTRACT
	BlendPresetModAlpha           BlendPresetEnum = C.GPU_BLEND_MOD_ALPHA
	BlendPresetSetAlpha           BlendPresetEnum = C.GPU_BLEND_SET_ALPHA
	BlendPresetSet                BlendPresetEnum = C.GPU_BLEND_SET
	BlendPresetNormalKeepAlpha    BlendPresetEnum = C.GPU_BLEND_NORMAL_KEEP_ALPHA
	BlendPresetAddAlpha           BlendPresetEnum = C.GPU_BLEND_NORMAL_ADD_ALPHA
	BlendPresetFactorAlpha        BlendPresetEnum = C.GPU_BLEND_NORMAL_FACTOR_ALPHA
)

type FilterEnum C.GPU_FilterEnum

const (
	FilterNearest      FilterEnum = C.GPU_FILTER_NEAREST
	FilterLinear       FilterEnum = C.GPU_FILTER_LINEAR
	FilterLinearMipmap FilterEnum = C.GPU_FILTER_LINEAR_MIPMAP
)

type SnapEnum C.GPU_SnapEnum

const (
	SnapNone                  SnapEnum = C.GPU_SNAP_NONE
	SnapPosition              SnapEnum = C.GPU_SNAP_POSITION
	SnapDimensions            SnapEnum = C.GPU_SNAP_DIMENSIONS
	SnapPositionAndDimensions SnapEnum = C.GPU_SNAP_POSITION_AND_DIMENSIONS
)

type WrapEnum C.GPU_WrapEnum

const (
	WrapNone     WrapEnum = C.GPU_WRAP_NONE
	WrapRepeat   WrapEnum = C.GPU_WRAP_REPEAT
	WrapMirrored WrapEnum = C.GPU_WRAP_MIRRORED
)

type FormatEnum C.GPU_FormatEnum

const (
	FormatLuminance      FormatEnum = C.GPU_FORMAT_LUMINANCE
	FormatLuminanceAlpha FormatEnum = C.GPU_FORMAT_LUMINANCE_ALPHA
	FormatRGB            FormatEnum = C.GPU_FORMAT_RGB
	FormatRGBA           FormatEnum = C.GPU_FORMAT_RGBA
	FormatAlpha          FormatEnum = C.GPU_FORMAT_ALPHA
	FormatRG             FormatEnum = C.GPU_FORMAT_RG
	FormatYCbCr422       FormatEnum = C.GPU_FORMAT_YCbCr422
	FormatYCbCr420P      FormatEnum = C.GPU_FORMAT_YCbCr420P
	FormatBGR            FormatEnum = C.GPU_FORMAT_BGR
	FormatBGRA           FormatEnum = C.GPU_FORMAT_BGRA
	FormatABGR           FormatEnum = C.GPU_FORMAT_ABGR
)

type FileFormatEnum C.GPU_FileFormatEnum

const (
	FileFormatAuto FileFormatEnum = C.GPU_FILE_AUTO
	FileFormatPNG  FileFormatEnum = C.GPU_FILE_PNG
	FileFormatBMP  FileFormatEnum = C.GPU_FILE_BMP
	FileFormatTGA  FileFormatEnum = C.GPU_FILE_TGA
)

// TODO: Resume at 210

// Make usable type
type Image C.GPU_Image

type Camera struct {
	X, Y, Z     float32
	Angle       float32
	Zoom        float32
	ZNear, ZFar float32
}

func (c *Camera) cptr() *C.GPU_Camera {
	return (*C.GPU_Camera)(unsafe.Pointer(c))
}

// Broken
// type TextureHandle c.GPU_TextureHandle

type Context C.GPU_Context

type Target C.GPU_Target

// TODO: L460 GPU_FeatureEnum
type FeatureEnum C.Uint32

const (
	FeatureNonPowerOfTwo          FeatureEnum = 0x1
	FeatureRenderTargets          FeatureEnum = 0x2
	FeatureBlendEquations         FeatureEnum = 0x4
	FeatureBlendFuncSeparate      FeatureEnum = 0x8
	FeatureBlendEquationsSeparate FeatureEnum = 0x10
	FeatureGLBGR                  FeatureEnum = 0x20
	FeatureGLBGRA                 FeatureEnum = 0x40
	FeatureGLABGR                 FeatureEnum = 0x80
	FeatureVertexShader           FeatureEnum = 0x100
	FeatureFragmentShader         FeatureEnum = 0x200
	FeaturePixelShader            FeatureEnum = 0x200
	FeatureGeometryShader         FeatureEnum = 0x400
	FeatureWrapRepeatMirrored     FeatureEnum = 0x800
	FeatureCoreFramebufferObjects FeatureEnum = 0x1000

	FeatureAllBase         = FeatureRenderTargets
	FeatureAllBlendPresets = FeatureBlendEquations | FeatureBlendFuncSeparate
	FeatureAllGLFormats    = FeatureGLBGR | FeatureGLBGRA | FeatureGLABGR
	FeatureBasicShaders    = FeatureFragmentShader | FeatureVertexShader

	// Why doesn't this include Pixel Shader? SDL_gpu.h doesn't have it either
	FeatureAllShaders = FeatureFragmentShader | FeatureVertexShader | FeatureGeometryShader
)

type InitFlagEnum C.Uint32

const (
	DefaultInitFlags                     InitFlagEnum = 0
	InitEnableVsync                      InitFlagEnum = 0x1
	InitDisableVsync                     InitFlagEnum = 0x2
	InitDisableDoubleBuffer              InitFlagEnum = 0x4
	InitDisableAutoVirtualResolution     InitFlagEnum = 0x8
	InitRequestCmpatibilityProfile       InitFlagEnum = 0x10
	InitUseRowByRowTextureUploadFallback InitFlagEnum = 0x20
	InitUsecopyTextureUploadFallback     InitFlagEnum = 0x40
)

type PrimitiveEnum C.Uint32

const (
	PrimitivePoints        PrimitiveEnum = 0x0
	PrimitiveLines         PrimitiveEnum = 0x1
	PrimitiveLineLoop      PrimitiveEnum = 0x2
	PrimitiveLineStrip     PrimitiveEnum = 0x3
	PrimitiveTriangles     PrimitiveEnum = 0x4
	PrimitiveTriangleStrip PrimitiveEnum = 0x5
	PrimitiveTriangleFan   PrimitiveEnum = 0x6
)

type BatchFlagEnum C.Uint32

const (
	BatchXY    BatchFlagEnum = 0x1
	BatchXYZ   BatchFlagEnum = 0x2
	BatchST    BatchFlagEnum = 0x4
	BatchRGB   BatchFlagEnum = 0x8
	BatchRGBA  BatchFlagEnum = 0x10
	BatchRGB8  BatchFlagEnum = 0x20
	BatchRGBA8 BatchFlagEnum = 0x40

	BatchXYST       = BatchXY | BatchST
	BatchXYZST      = BatchXYZ | BatchST
	BatchXYRGB      = BatchXY | BatchRGB
	BatchXYZRGB     = BatchXYZ | BatchRGB
	BatchXYRGBA     = BatchXY | BatchRGBA
	BatchXYZRGBA    = BatchXYZ | BatchRGBA
	BatchXYSTRGBA   = BatchXY | BatchST | BatchRGBA
	BatchXYZSTRGBA  = BatchXYZ | BatchST | BatchRGBA
	BatchXYRGB8     = BatchXY | BatchRGB8
	BatchXYZRGB8    = BatchXYZ | BatchRGB8
	BatchXYRGBA8    = BatchXY | BatchRGBA8
	BatchXYZRGBA8   = BatchXYZ | BatchRGBA8
	BatchXYSTRGBA8  = BatchXY | BatchST | BatchRGBA8
	BatchXYZSTRGBA8 = BatchXYZ | BatchST | BatchRGBA8
)

type FlipEnum C.Uint32

const (
	FlipNone       FlipEnum = 0x0
	FlipHorizontal FlipEnum = 0x1
	FlipVertical   FlipEnum = 0x2
)

type TypeEnum C.Uint32

const (
	TypeByte          TypeEnum = 0x1400
	TypeUnsignedByte  TypeEnum = 0x1401
	TypeShort         TypeEnum = 0x1402
	TypeUnsignedShort TypeEnum = 0x1403
	TypeInt           TypeEnum = 0x1404
	TypeUnsignedInt   TypeEnum = 0x1405
	TypeFloat         TypeEnum = 0x1406
	TypeDoubel        TypeEnum = 0x140A
)

const (
	VertexShader   = C.GPU_VERTEX_SHADER
	FragmentShader = C.GPU_FRAGMENT_SHADER
	PixelShader    = C.GPU_PIXEL_SHADER
	GeometryShader = C.GPU_GEOMETRY_SHADER
)

const (
	LanguageNone        = C.GPU_LANGUAGE_NONE
	LanguageARBAssembly = C.GPU_LANGUAGE_ARB_ASSEMBLY
	LanguageGLSL        = C.GPU_LANGUAGE_GLSL
	LanguageGLSLES      = C.GPU_LANGUAGE_GLSLES
	LanguageHLSL        = C.GPU_LANGUAGE_HLSL
	LanguageCG          = C.GPU_LANGUAGE_CG
)

type AttributeFormat C.GPU_AttributeFormat

type Attribute C.GPU_Attribute

type AttributeSource C.GPU_AttributeSource

const (
	ErrorNone                = C.GPU_ERROR_NONE
	ErrorBackendError        = C.GPU_ERROR_BACKEND_ERROR
	ErrorDataError           = C.GPU_ERROR_DATA_ERROR
	ErrorUserError           = C.GPU_ERROR_USER_ERROR
	ErrorUnsupportedFunction = C.GPU_ERROR_UNSUPPORTED_FUNCTION
	ErrorNullArgument        = C.GPU_ERROR_NULL_ARGUMENT
	ErrorFileNotFound        = C.GPU_ERROR_FILE_NOT_FOUND
)

type ErrorObject C.GPU_ErrorObject

const (
	DebugLevel0   = C.GPU_DEBUG_LEVEL_0
	DebugLevel1   = C.GPU_DEBUG_LEVEL_1
	DebugLevel2   = C.GPU_DEBUG_LEVEL_2
	DebugLevel3   = C.GPU_DEBUG_LEVEL_3
	DebugLevelMax = C.GPU_DEBUG_LEVEL_MAX
)

const (
	LogInfo    = C.GPU_LOG_INFO
	LogWarning = C.GPU_LOG_WARNING
	LogError   = C.GPU_LOG_ERROR
)

type Renderer C.GPU_Renderer

// Initialization Begin

func Init(w, h uint16, flags InitFlagEnum) *Target {
	return (*Target)(unsafe.Pointer(C.GPU_Init(C.Uint16(w),
		C.Uint16(h),
		C.Uint32(flags))))
}

func InitRenderer(renderer RendererEnum, w, h uint16, windowFlags uint32) *Target {
	return (*Target)(unsafe.Pointer(C.GPU_InitRenderer(C.GPU_RendererEnum(renderer),
		C.Uint16(w),
		C.Uint16(h),
		C.Uint32(windowFlags))))
}

func InitRendererByID(rendererID RendererID, w, h uint16, windowFlags uint32) *Target {
	return (*Target)(unsafe.Pointer(C.GPU_InitRendererByID(C.GPU_RendererID(rendererID),
		C.Uint16(w),
		C.Uint16(h),
		C.Uint32(windowFlags))))
}

func SetInitWindow(windowID uint32) {
	C.GPU_SetInitWindow(C.Uint32(windowID))
}

func GetInitWindow() uint32 {
	return uint32(C.GPU_GetInitWindow())
}

func SetPreInitFlags(flags InitFlagEnum) {
	C.GPU_SetPreInitFlags(C.Uint32(flags))
}

func IsFeatureEnabled(features FeatureEnum) bool {
	return bool(C.GPU_IsFeatureEnabled(C.Uint32(features)))
}

func SetRequiredFeatures(features FeatureEnum) {
	C.GPU_SetRequiredFeatures(C.Uint32(features))
}

func GetRequiredFeatures() FeatureEnum {
	return FeatureEnum(C.GPU_GetRequiredFeatures())
}

func GetDefaultRendererOrder() []RendererID {
	var orderSize int
	var order []RendererID = make([]RendererID, RendererOrderMax)

	C.GPU_GetDefaultRendererOrder(
		(*C.int)(unsafe.Pointer(&orderSize)),
		(*C.GPU_RendererID)(unsafe.Pointer(&order[0])),
	)

	return order[:orderSize]
}

func SetRendererOrder(renderers []RendererID) {
	var _renderers *C.GPU_RendererID
	if len(renderers) > 0 {
		_renderers = (*C.GPU_RendererID)(unsafe.Pointer(&renderers[0]))
	}
	_size := C.int(len(renderers))

	C.GPU_SetRendererOrder(_size, _renderers)
}

func CloseCurrentRenderer() {
	C.GPU_CloseCurrentRenderer()
}

func Quit() {
	C.GPU_Quit()
}

// Initialization End

// Begin Renderer Setup

func MakeRendererID(name string, renderer RendererEnum, majorVersion, minorVersion int) RendererID {
	// Do I need to free this? Looks like it gets stashed in the returned RendererID
	_name := C.CString(name)
	_renderer := C.GPU_RendererEnum(renderer)
	_major := C.int(majorVersion)
	_minor := C.int(minorVersion)

	return RendererID(C.GPU_MakeRendererID(_name, _renderer, _major, _minor))
}

func GetRendererID(renderer RendererEnum) RendererID {
	return RendererID(C.GPU_GetRendererID(C.GPU_RendererEnum(renderer)))
}

func GetNumRegisteredRenderers() int {
	return int(C.GPU_GetNumRegisteredRenderers())
}

func GetRegisteredRenderList() []RendererID {
	numRenderers := GetNumRegisteredRenderers()
	var renderers []RendererID = make([]RendererID, numRenderers)

	C.GPU_GetRegisteredRendererList(
		(*C.GPU_RendererID)(unsafe.Pointer(&renderers[0])),
	)

	return renderers
}

func RegisterRenderer(id RendererID) {
	// TODO: I don't know how to implement this
	// DECLSPEC void SDLCALL GPU_RegisterRenderer(GPU_RendererID id, GPU_Renderer* (SDLCALL *create_renderer)(GPU_RendererID request), void (SDLCALL *free_renderer)(GPU_Renderer* renderer));
	panic("Not Implemented")
}

// End Renderer Setup

// Begin Target
func LoadTarget(i *Image) *Target {
	return (*Target)(C.GPU_LoadTarget(i.cptr()))
}

func GetDefaultCamera() Camera {
	camera := C.GPU_GetDefaultCamera()
	return Camera{
		X:     float32(camera.x),
		Y:     float32(camera.y),
		Z:     float32(camera.z),
		Angle: float32(camera.angle),
		Zoom:  float32(camera.zoom),
		ZNear: float32(camera.z_near),
		ZFar:  float32(camera.z_far),
	}
}

func (t *Target) GetCamera() Camera {
	camera := C.GPU_GetCamera(t.cptr())
	return Camera{
		X:     float32(camera.x),
		Y:     float32(camera.y),
		Z:     float32(camera.z),
		Angle: float32(camera.angle),
		Zoom:  float32(camera.zoom),
		ZNear: float32(camera.z_near),
		ZFar:  float32(camera.z_far),
	}
}

func (t *Target) SetCamera(camera *Camera) {
	C.GPU_SetCamera(t.cptr(), camera.cptr())
}

func (t *Target) EnableCamera(enable bool) {
	C.GPU_EnableCamera(t.cptr(), C.bool(enable))
}

func (t *Target) IsCameraEnabled() bool {
	return bool(C.GPU_IsCameraEnabled(t.cptr()))
}

func (t *Target) GetPixel(x, y int16) Color {
	_color := C.GPU_GetPixel(t.cptr(), C.Sint16(x), C.Sint16(y))

	return Color{uint8(_color.r), uint8(_color.g), uint8(_color.b), uint8(_color.a)}
}

func (t *Target) SetClipRect(rect Rect) {
	_r := C.GPU_Rect{C.float(rect.X), C.float(rect.Y), C.float(rect.W), C.float(rect.H)}
	C.GPU_SetClipRect(t.cptr(), _r)
}

func (t *Target) SetClip(x, y int16, w, h uint16) {
	C.GPU_SetClip(t.cptr(), C.Sint16(x), C.Sint16(y), C.Uint16(w), C.Uint16(h))
}

func (t *Target) UnsetClip() {
	C.GPU_UnsetClip(t.cptr())
}

func IntersectRect(r1 Rect, r2 Rect) (Rect, bool) {
	// Do I need to allocate?
	var r3 Rect

	_r1 := C.GPU_Rect{C.float(r1.X), C.float(r1.Y), C.float(r1.W), C.float(r1.H)}
	_r2 := C.GPU_Rect{C.float(r2.X), C.float(r2.Y), C.float(r2.W), C.float(r2.H)}

	ok := bool(C.GPU_IntersectRect(_r1, _r2, (&r3).cptr()))

	return r3, ok
}

func (t *Target) IntersectClipRect(r1 Rect) (Rect, bool) {
	// Do I need to allocate?
	var r2 Rect

	_r1 := C.GPU_Rect{C.float(r1.X), C.float(r1.Y), C.float(r1.W), C.float(r1.H)}

	ok := bool(C.GPU_IntersectClipRect(t.cptr(), _r1, (&r2).cptr()))

	return r2, ok
}

// Not sure if it should hit the c lib, or just do this
func MakeRect(x, y, w, h float32) Rect {
	return Rect{x, y, w, h}
}

// Not sure if it should hit the c lib, or just do this
func MakeColor(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}

func (t *Target) cptr() *C.GPU_Target {
	return (*C.GPU_Target)(unsafe.Pointer(t))
}

func (t *Target) CreateAlias() *Target {
	return (*Target)(C.GPU_CreateAliasTarget(t.cptr()))
}

func (t *Target) Free() {
	C.GPU_FreeTarget(t.cptr())
}

func (t *Target) SetVirtualResolution(w, h uint16) {
	C.GPU_SetVirtualResolution(t.cptr(), C.Uint16(w), C.Uint16(h))
}

func (t *Target) GetVirtualResolution() (width, height uint16) {
	C.GPU_GetVirtualResolution(
		t.cptr(), (*C.Uint16)(&width), (*C.Uint16)(&height),
	)
	return width, height
}

func (t *Target) GetVirtualCoords(displayX, displayY float32) (x, y float32) {
	C.GPU_GetVirtualCoords(
		t.cptr(), (*C.float)(&x), (*C.float)(&y), C.float(displayX), C.float(displayY),
	)
	return x, y
}

func (t *Target) UnsetVirtualResolution() {
	C.GPU_UnsetVirtualResolution(t.cptr())
}

func (t *Target) SetViewport(viewport Rect) {
	_r := C.GPU_Rect{C.float(viewport.X), C.float(viewport.Y), C.float(viewport.W), C.float(viewport.H)}
	C.GPU_SetViewport(t.cptr(), _r)
}

func (t *Target) UnsetViewport() {
	C.GPU_UnsetViewport(t.cptr())
}

func (t *Target) SetColor(color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_SetTargetColor(t.cptr(), _c)
}

func (t *Target) SetRGB(r, g, b uint8) {
	C.GPU_SetTargetRGB(t.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b))
}

func (t *Target) SetRGBA(r, g, b, a uint8) {
	C.GPU_SetTargetRGBA(t.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a))
}

func (t *Target) UnsetColor() {
	C.GPU_UnsetTargetColor(t.cptr())
}

// End Target

// Begin Surface
type Surface C.SDL_Surface
type RWOps C.SDL_RWops

func (s *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(s))
}

func LoadSurface(fileName string) *Surface {
	_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(_fileName))

	return (*Surface)(C.GPU_LoadSurface(_fileName))
}

func (s *Surface) SaveSurface(fileName string, format FileFormatEnum) bool {
	_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(_fileName))

	return bool(C.GPU_SaveSurface(s.cptr(), _fileName, C.GPU_FileFormatEnum(format)))
}

// TODO: GPU_LoadSurface_RW, GPU_SaveSurface_RW

// End Surface

// Begin Image
func (i *Image) cptr() *C.GPU_Image {
	return (*C.GPU_Image)(unsafe.Pointer(i))
}

// End Image

// Begin Rendering
func (t *Target) Clear() {
	C.GPU_Clear(t.cptr())
}

func (t *Target) ClearColor(color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_ClearColor(t.cptr(), _c)
}

func (t *Target) ClearRGB(r, g, b uint8) {
	C.GPU_ClearRGB(t.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b))
}

func (t *Target) ClearRGBA(r, g, b, a uint8) {
	C.GPU_ClearRGBA(t.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a))
}

func (t *Target) Blit(image *Image, source *Rect, x, y float32) {
	C.GPU_Blit(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		C.float(x),
		C.float(y),
	)
}

func (t *Target) BlitRotate(image *Image, source *Rect, x, y, degrees float32) {
	C.GPU_BlitRotate(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		C.float(x),
		C.float(y),
		C.float(degrees),
	)
}

func (t *Target) BlitScale(image *Image, source *Rect, x, y, scaleX, scaleY float32) {
	C.GPU_BlitScale(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		C.float(x),
		C.float(y),
		C.float(scaleX),
		C.float(scaleY),
	)
}

func (t *Target) BlitTransform(image *Image, source *Rect, x, y, degrees, scaleX, scaleY float32) {
	C.GPU_BlitTransform(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		C.float(x),
		C.float(y),
		C.float(degrees),
		C.float(scaleX),
		C.float(scaleY),
	)
}

func (t *Target) BlitTransformX(image *Image, source *Rect, x, y, pivotX, pivotY, degrees, scaleX, scaleY float32) {
	C.GPU_BlitTransformX(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		C.float(x),
		C.float(y),
		C.float(pivotX),
		C.float(pivotY),
		C.float(degrees),
		C.float(scaleX),
		C.float(scaleY),
	)
}

func (t *Target) BlitRect(image *Image, source, dest *Rect) {
	C.GPU_BlitRect(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(dest)),
	)
}

func (t *Target) BlitRectX(image *Image, source, dest *Rect, degrees, pivotX, pivotY float32, flip FlipEnum) {
	C.GPU_BlitRectX(image.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(source)),
		t.cptr(),
		(*C.GPU_Rect)(unsafe.Pointer(dest)),
		C.float(degrees),
		C.float(pivotX),
		C.float(pivotY),
		C.GPU_FlipEnum(flip),
	)
}

// TODO: Double Check this
func (t *Target) TriangleBatch(image *Image, numVertices uint8, vertices []float32, numIndices uint, indices []uint8, batchFlags BatchFlagEnum) {
	var _verts *C.float
	var _indices *C.ushort

	if len(vertices) > 0 {
		_verts = (*C.float)(unsafe.Pointer(&vertices[0]))
	}
	if len(indices) > 0 {
		_indices = (*C.ushort)(unsafe.Pointer(&indices[0]))
	}

	C.GPU_TriangleBatch(image.cptr(),
		t.cptr(),
		C.ushort(numVertices),
		_verts,
		C.uint(numIndices),
		_indices,
		C.GPU_BatchFlagEnum(batchFlags),
	)
}

// TODO: Double Check this
func (t *Target) TriangleBatchX(image *Image, numVertices uint8, vertices []interface{}, numIndices uint, indices []uint8, batchFlags BatchFlagEnum) {
	var _verts unsafe.Pointer
	var _indices *C.ushort

	if len(vertices) > 0 {
		_verts = unsafe.Pointer(&vertices[0])
	}
	if len(indices) > 0 {
		_indices = (*C.ushort)(unsafe.Pointer(&indices[0]))
	}

	C.GPU_TriangleBatchX(image.cptr(),
		t.cptr(),
		C.ushort(numVertices),
		_verts,
		C.uint(numIndices),
		_indices,
		C.GPU_BatchFlagEnum(batchFlags),
	)
}

func FlushBlitBuffer() {
	C.GPU_FlushBlitBuffer()
}

func (t *Target) Flip() {
	C.GPU_Flip(t.cptr())
}

// End Rendering

// Shapes Begin

func (t *Target) Pixel(x, y float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Pixel(t.cptr(),
		C.float(x),
		C.float(y),
		_c)
}

func (t *Target) Line(x1, y1, x2, y2 float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Line(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		_c)
}

func (t *Target) Arc(x, y, radius, startAngle, endAngle float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Arc(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(radius),
		C.float(startAngle),
		C.float(endAngle),
		_c)
}

func (t *Target) ArcFilled(x, y, radius, startAngle, endAngle float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_ArcFilled(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(radius),
		C.float(startAngle),
		C.float(endAngle),
		_c)
}

func (t *Target) Circle(x, y, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Circle(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(radius),
		_c)
}

func (t *Target) CircleFilled(x, y, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_CircleFilled(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(radius),
		_c)
}

func (t *Target) Ellipse(x, y, rX, rY, degrees float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Ellipse(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(rX),
		C.float(rY),
		C.float(degrees),
		_c)
}

func (t *Target) EllipseFilled(x, y, rX, rY, degrees float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_EllipseFilled(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(rX),
		C.float(rY),
		C.float(degrees),
		_c)
}

func (t *Target) Sector(x, y, innerRadius, outerRadius, startAngle, endAngle float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Sector(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(innerRadius),
		C.float(outerRadius),
		C.float(startAngle),
		C.float(endAngle),
		_c)
}

func (t *Target) SectorFilled(x, y, innerRadius, outerRadius, startAngle, endAngle float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_SectorFilled(t.cptr(),
		C.float(x),
		C.float(y),
		C.float(innerRadius),
		C.float(outerRadius),
		C.float(startAngle),
		C.float(endAngle),
		_c)
}

func (t *Target) Tri(x1, y1, x2, y2, x3, y3 float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Tri(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		C.float(x3),
		C.float(y3),
		_c)
}

func (t *Target) TriFilled(x1, y1, x2, y2, x3, y3 float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_TriFilled(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		C.float(x3),
		C.float(y3),
		_c)
}

func (t *Target) Rectangle(x1, y1, x2, y2 float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_Rectangle(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		_c)
}

func (t *Target) Rectangle2(rect Rect, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	_r := C.GPU_Rect{C.float(rect.X), C.float(rect.Y), C.float(rect.W), C.float(rect.H)}
	C.GPU_Rectangle2(t.cptr(),
		_r,
		_c)
}

func (t *Target) RectangleFilled(x1, y1, x2, y2 float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_RectangleFilled(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		_c)
}

func (t *Target) RectangleFilled2(rect Rect, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	_r := C.GPU_Rect{C.float(rect.X), C.float(rect.Y), C.float(rect.W), C.float(rect.H)}
	C.GPU_RectangleFilled2(t.cptr(),
		_r,
		_c)
}

func (t *Target) RectangleRound(x1, y1, x2, y2, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_RectangleRound(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		C.float(radius),
		_c)
}

func (t *Target) RectangleRound2(rect Rect, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	_r := C.GPU_Rect{C.float(rect.X), C.float(rect.Y), C.float(rect.W), C.float(rect.H)}
	C.GPU_RectangleRound2(t.cptr(),
		_r,
		C.float(radius),
		_c)
}

func (t *Target) RectangleRoundFilled(x1, y1, x2, y2, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	C.GPU_RectangleRoundFilled(t.cptr(),
		C.float(x1),
		C.float(y1),
		C.float(x2),
		C.float(y2),
		C.float(radius),
		_c)
}

func (t *Target) RectangleRoundFilled2(rect Rect, radius float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	_r := C.GPU_Rect{C.float(rect.X), C.float(rect.Y), C.float(rect.W), C.float(rect.H)}
	C.GPU_RectangleRoundFilled2(t.cptr(),
		_r,
		C.float(radius),
		_c)
}

func (t *Target) Polygon(vertices []float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	numVert := len(vertices) / 2
	_vert := (*C.float)(unsafe.Pointer(&vertices[0]))
	C.GPU_Polygon(t.cptr(),
		C.Uint32(numVert),
		_vert,
		_c)
}

func (t *Target) PolygonFilled(vertices []float32, color Color) {
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	numVert := len(vertices) / 2
	_vert := (*C.float)(unsafe.Pointer(&vertices[0]))
	C.GPU_PolygonFilled(t.cptr(),
		C.Uint32(numVert),
		_vert,
		_c)
}

// Shapes End
