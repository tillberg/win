// Borrowed from https://github.com/AllenDang/w32/blob/ad0a36d80adcd081d5c0dded8e97a009b486d1db/dwmapi.go et al
// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwm

import (
	"fmt"
	"syscall"
	"unsafe"
)

// DEFINED IN THE DWM API BUT NOT IMPLEMENTED BY MS:
// DwmAttachMilContent
// DwmDetachMilContent
// DwmEnableComposition
// DwmGetGraphicsStreamClient
// DwmGetGraphicsStreamTransformHint

type (
	ATOM            uint16
	BOOL            int32
	COLORREF        uint32
	DWM_FRAME_COUNT uint64
	DWORD           uint32
	HACCEL          HANDLE
	HANDLE          uintptr
	HBITMAP         HANDLE
	HBRUSH          HANDLE
	HCURSOR         HANDLE
	HDC             HANDLE
	HDROP           HANDLE
	HDWP            HANDLE
	HENHMETAFILE    HANDLE
	HFONT           HANDLE
	HGDIOBJ         HANDLE
	HGLOBAL         HANDLE
	HGLRC           HANDLE
	HHOOK           HANDLE
	HICON           HANDLE
	HIMAGELIST      HANDLE
	HINSTANCE       HANDLE
	HKEY            HANDLE
	HKL             HANDLE
	HMENU           HANDLE
	HMODULE         HANDLE
	HMONITOR        HANDLE
	HPEN            HANDLE
	HRESULT         int32
	HRGN            HANDLE
	HRSRC           HANDLE
	HTHUMBNAIL      HANDLE
	HWND            HANDLE
	LPARAM          uintptr
	LPCVOID         unsafe.Pointer
	LRESULT         uintptr
	PVOID           unsafe.Pointer
	QPC_TIME        uint64
	ULONG_PTR       uintptr
	WPARAM          uintptr
	TRACEHANDLE     uintptr
)

// Flags used by the DWM_BLURBEHIND structure to indicate
// which of its members contain valid information.
const (
	DWM_BB_ENABLE                = 0x00000001 //     A value for the fEnable member has been specified.
	DWM_BB_BLURREGION            = 0x00000002 //     A value for the hRgnBlur member has been specified.
	DWM_BB_TRANSITIONONMAXIMIZED = 0x00000004 //     A value for the fTransitionOnMaximized member has been specified.
)

// Flags used by the DwmEnableComposition  function
// to change the state of Desktop Window Manager (DWM) composition.
const (
	DWM_EC_DISABLECOMPOSITION = 0 //     Disable composition
	DWM_EC_ENABLECOMPOSITION  = 1 //     Enable composition
)

// enum-lite implementation for the following constant structure
type DWM_SHOWCONTACT int32

const (
	DWMSC_DOWN      = 0x00000001
	DWMSC_UP        = 0x00000002
	DWMSC_DRAG      = 0x00000004
	DWMSC_HOLD      = 0x00000008
	DWMSC_PENBARREL = 0x00000010
	DWMSC_NONE      = 0x00000000
	DWMSC_ALL       = 0xFFFFFFFF
)

// enum-lite implementation for the following constant structure
type DWM_SOURCE_FRAME_SAMPLING int32

// TODO: need to verify this construction
// Flags used by the DwmSetPresentParameters function
// to specify the frame sampling type
const (
	DWM_SOURCE_FRAME_SAMPLING_POINT = iota + 1
	DWM_SOURCE_FRAME_SAMPLING_COVERAGE
	DWM_SOURCE_FRAME_SAMPLING_LAST
)

// Flags used by the DWM_THUMBNAIL_PROPERTIES structure to
// indicate which of its members contain valid information.
const (
	DWM_TNP_RECTDESTINATION      = 0x00000001 //    A value for the rcDestination member has been specified
	DWM_TNP_RECTSOURCE           = 0x00000002 //    A value for the rcSource member has been specified
	DWM_TNP_OPACITY              = 0x00000004 //    A value for the opacity member has been specified
	DWM_TNP_VISIBLE              = 0x00000008 //    A value for the fVisible member has been specified
	DWM_TNP_SOURCECLIENTAREAONLY = 0x00000010 //    A value for the fSourceClientAreaOnly member has been specified
)

// enum-lite implementation for the following constant structure
type DWMFLIP3DWINDOWPOLICY int32

// TODO: need to verify this construction
// Flags used by the DwmSetWindowAttribute function
// to specify the Flip3D window policy
const (
	DWMFLIP3D_DEFAULT = iota + 1
	DWMFLIP3D_EXCLUDEBELOW
	DWMFLIP3D_EXCLUDEABOVE
	DWMFLIP3D_LAST
)

// enum-lite implementation for the following constant structure
type DWMNCRENDERINGPOLICY int32

// TODO: need to verify this construction
// Flags used by the DwmSetWindowAttribute function
// to specify the non-client area rendering policy
const (
	DWMNCRP_USEWINDOWSTYLE = iota + 1
	DWMNCRP_DISABLED
	DWMNCRP_ENABLED
	DWMNCRP_LAST
)

// enum-lite implementation for the following constant structure
type DWMTRANSITION_OWNEDWINDOW_TARGET int32

const (
	DWMTRANSITION_OWNEDWINDOW_NULL       = -1
	DWMTRANSITION_OWNEDWINDOW_REPOSITION = 0
)

// enum-lite implementation for the following constant structure
type DWMWINDOWATTRIBUTE int32

// TODO: need to verify this construction
// Flags used by the DwmGetWindowAttribute and DwmSetWindowAttribute functions
// to specify window attributes for non-client rendering
const (
	DWMWA_NCRENDERING_ENABLED = iota + 1
	DWMWA_NCRENDERING_POLICY
	DWMWA_TRANSITIONS_FORCEDISABLED
	DWMWA_ALLOW_NCPAINT
	DWMWA_CAPTION_BUTTON_BOUNDS
	DWMWA_NONCLIENT_RTL_LAYOUT
	DWMWA_FORCE_ICONIC_REPRESENTATION
	DWMWA_FLIP3D_POLICY
	DWMWA_EXTENDED_FRAME_BOUNDS
	DWMWA_HAS_ICONIC_BITMAP
	DWMWA_DISALLOW_PEEK
	DWMWA_EXCLUDED_FROM_PEEK
	DWMWA_CLOAK
	DWMWA_CLOAKED
	DWMWA_FREEZE_REPRESENTATION
	DWMWA_LAST
)

// enum-lite implementation for the following constant structure
type GESTURE_TYPE int32

// TODO: use iota?
// Identifies the gesture type
const (
	GT_PEN_TAP                 = 0
	GT_PEN_DOUBLETAP           = 1
	GT_PEN_RIGHTTAP            = 2
	GT_PEN_PRESSANDHOLD        = 3
	GT_PEN_PRESSANDHOLDABORT   = 4
	GT_TOUCH_TAP               = 5
	GT_TOUCH_DOUBLETAP         = 6
	GT_TOUCH_RIGHTTAP          = 7
	GT_TOUCH_PRESSANDHOLD      = 8
	GT_TOUCH_PRESSANDHOLDABORT = 9
	GT_TOUCH_PRESSANDTAP       = 10
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa969505.aspx
type UNSIGNED_RATIO struct {
	uiNumerator   uint32
	uiDenominator uint32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145106.aspx
type SIZE struct {
	CX, CY int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162897.aspx
type RECT struct {
	Left, Top, Right, Bottom int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/bb773244.aspx
type MARGINS struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa969500.aspx
type DWM_BLURBEHIND struct {
	DwFlags                uint32
	fEnable                BOOL
	hRgnBlur               HRGN
	fTransitionOnMaximized BOOL
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa969501.aspx
type DWM_PRESENT_PARAMETERS struct {
	cbSize             uint32
	fQueue             BOOL
	cRefreshStart      DWM_FRAME_COUNT
	cBuffer            uint32
	fUseSourceRate     BOOL
	rateSource         UNSIGNED_RATIO
	cRefreshesPerFrame uint32
	eSampling          DWM_SOURCE_FRAME_SAMPLING
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa969502.aspx
type DWM_THUMBNAIL_PROPERTIES struct {
	DWFlags               uint32
	RCDestination         RECT
	RCSource              RECT
	Opacity               byte
	FVisible              BOOL
	FSourceClientAreaOnly BOOL
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa969503.aspx
type DWM_TIMING_INFO struct {
	cbSize                 uint32
	rateRefresh            UNSIGNED_RATIO
	qpcRefreshPeriod       QPC_TIME
	rateCompose            UNSIGNED_RATIO
	qpcVBlank              QPC_TIME
	cRefresh               DWM_FRAME_COUNT
	cDXRefresh             uint32
	qpcCompose             QPC_TIME
	cFrame                 DWM_FRAME_COUNT
	cDXPresent             uint32
	cRefreshFrame          DWM_FRAME_COUNT
	cFrameSubmitted        DWM_FRAME_COUNT
	cDXPresentSubmitted    uint32
	cFrameConfirmed        DWM_FRAME_COUNT
	cDXPresentConfirmed    uint32
	cRefreshConfirmed      DWM_FRAME_COUNT
	cDXRefreshConfirmed    uint32
	cFramesLate            DWM_FRAME_COUNT
	cFramesOutstanding     uint32
	cFrameDisplayed        DWM_FRAME_COUNT
	qpcFrameDisplayed      QPC_TIME
	cRefreshFrameDisplayed DWM_FRAME_COUNT
	cFrameComplete         DWM_FRAME_COUNT
	qpcFrameComplete       QPC_TIME
	cFramePending          DWM_FRAME_COUNT
	qpcFramePending        QPC_TIME
	cFramesDisplayed       DWM_FRAME_COUNT
	cFramesComplete        DWM_FRAME_COUNT
	cFramesPending         DWM_FRAME_COUNT
	cFramesAvailable       DWM_FRAME_COUNT
	cFramesDropped         DWM_FRAME_COUNT
	cFramesMissed          DWM_FRAME_COUNT
	cRefreshNextDisplayed  DWM_FRAME_COUNT
	cRefreshNextPresented  DWM_FRAME_COUNT
	cRefreshesDisplayed    DWM_FRAME_COUNT
	cRefreshesPresented    DWM_FRAME_COUNT
	cRefreshStarted        DWM_FRAME_COUNT
	cPixelsReceived        uint64
	cPixelsDrawn           uint64
	cBuffersEmpty          DWM_FRAME_COUNT
}

var (
	moddwmapi = syscall.NewLazyDLL("dwmapi.dll")

	procDwmDefWindowProc                 = moddwmapi.NewProc("DwmDefWindowProc")
	procDwmEnableBlurBehindWindow        = moddwmapi.NewProc("DwmEnableBlurBehindWindow")
	procDwmEnableMMCSS                   = moddwmapi.NewProc("DwmEnableMMCSS")
	procDwmExtendFrameIntoClientArea     = moddwmapi.NewProc("DwmExtendFrameIntoClientArea")
	procDwmFlush                         = moddwmapi.NewProc("DwmFlush")
	procDwmGetColorizationColor          = moddwmapi.NewProc("DwmGetColorizationColor")
	procDwmGetCompositionTimingInfo      = moddwmapi.NewProc("DwmGetCompositionTimingInfo")
	procDwmGetTransportAttributes        = moddwmapi.NewProc("DwmGetTransportAttributes")
	procDwmGetWindowAttribute            = moddwmapi.NewProc("DwmGetWindowAttribute")
	procDwmInvalidateIconicBitmaps       = moddwmapi.NewProc("DwmInvalidateIconicBitmaps")
	procDwmIsCompositionEnabled          = moddwmapi.NewProc("DwmIsCompositionEnabled")
	procDwmModifyPreviousDxFrameDuration = moddwmapi.NewProc("DwmModifyPreviousDxFrameDuration")
	procDwmQueryThumbnailSourceSize      = moddwmapi.NewProc("DwmQueryThumbnailSourceSize")
	procDwmRegisterThumbnail             = moddwmapi.NewProc("DwmRegisterThumbnail")
	procDwmRenderGesture                 = moddwmapi.NewProc("DwmRenderGesture")
	procDwmSetDxFrameDuration            = moddwmapi.NewProc("DwmSetDxFrameDuration")
	procDwmSetIconicLivePreviewBitmap    = moddwmapi.NewProc("DwmSetIconicLivePreviewBitmap")
	procDwmSetIconicThumbnail            = moddwmapi.NewProc("DwmSetIconicThumbnail")
	procDwmSetPresentParameters          = moddwmapi.NewProc("DwmSetPresentParameters")
	procDwmSetWindowAttribute            = moddwmapi.NewProc("DwmSetWindowAttribute")
	procDwmShowContact                   = moddwmapi.NewProc("DwmShowContact")
	procDwmTetherContact                 = moddwmapi.NewProc("DwmTetherContact")
	procDwmTransitionOwnedWindow         = moddwmapi.NewProc("DwmTransitionOwnedWindow")
	procDwmUnregisterThumbnail           = moddwmapi.NewProc("DwmUnregisterThumbnail")
	procDwmUpdateThumbnailProperties     = moddwmapi.NewProc("DwmUpdateThumbnailProperties")
)

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}
	return 0
}

func DwmDefWindowProc(hWnd HWND, msg uint, wParam, lParam uintptr) (bool, uint) {
	var result uint
	ret, _, _ := procDwmDefWindowProc.Call(
		uintptr(hWnd),
		uintptr(msg),
		wParam,
		lParam,
		uintptr(unsafe.Pointer(&result)))
	return ret != 0, result
}

func DwmEnableBlurBehindWindow(hWnd HWND, pBlurBehind *DWM_BLURBEHIND) HRESULT {
	ret, _, _ := procDwmEnableBlurBehindWindow.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pBlurBehind)))
	return HRESULT(ret)
}

func DwmEnableMMCSS(fEnableMMCSS bool) HRESULT {
	ret, _, _ := procDwmEnableMMCSS.Call(
		uintptr(BoolToBOOL(fEnableMMCSS)))
	return HRESULT(ret)
}

func DwmExtendFrameIntoClientArea(hWnd HWND, pMarInset *MARGINS) HRESULT {
	ret, _, _ := procDwmExtendFrameIntoClientArea.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pMarInset)))
	return HRESULT(ret)
}

func DwmFlush() HRESULT {
	ret, _, _ := procDwmFlush.Call()
	return HRESULT(ret)
}

func DwmGetColorizationColor(pcrColorization *uint32, pfOpaqueBlend *BOOL) HRESULT {
	ret, _, _ := procDwmGetColorizationColor.Call(
		uintptr(unsafe.Pointer(pcrColorization)),
		uintptr(unsafe.Pointer(pfOpaqueBlend)))
	return HRESULT(ret)
}

func DwmGetCompositionTimingInfo(hWnd HWND, pTimingInfo *DWM_TIMING_INFO) HRESULT {
	ret, _, _ := procDwmGetCompositionTimingInfo.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pTimingInfo)))
	return HRESULT(ret)
}

func DwmGetTransportAttributes(pfIsRemoting *BOOL, pfIsConnected *BOOL, pDwGeneration *uint32) HRESULT {
	ret, _, _ := procDwmGetTransportAttributes.Call(
		uintptr(unsafe.Pointer(pfIsRemoting)),
		uintptr(unsafe.Pointer(pfIsConnected)),
		uintptr(unsafe.Pointer(pDwGeneration)))
	return HRESULT(ret)
}

// TODO: verify handling of variable arguments
func DwmGetWindowAttribute(hWnd HWND, dwAttribute uint32) (pAttribute interface{}, result HRESULT) {
	var pvAttribute, pvAttrSize uintptr
	switch dwAttribute {
	case DWMWA_NCRENDERING_ENABLED:
		v := new(BOOL)
		pAttribute = v
		pvAttribute = uintptr(unsafe.Pointer(v))
		pvAttrSize = unsafe.Sizeof(*v)
	case DWMWA_CAPTION_BUTTON_BOUNDS, DWMWA_EXTENDED_FRAME_BOUNDS:
		v := new(RECT)
		pAttribute = v
		pvAttribute = uintptr(unsafe.Pointer(v))
		pvAttrSize = unsafe.Sizeof(*v)
	case DWMWA_CLOAKED:
		panic(fmt.Sprintf("DwmGetWindowAttribute(%d) is not currently supported.", dwAttribute))
	default:
		panic(fmt.Sprintf("DwmGetWindowAttribute(%d) is not valid.", dwAttribute))
	}

	ret, _, _ := procDwmGetWindowAttribute.Call(
		uintptr(hWnd),
		uintptr(dwAttribute),
		pvAttribute,
		pvAttrSize)
	result = HRESULT(ret)
	return
}

func DwmInvalidateIconicBitmaps(hWnd HWND) HRESULT {
	ret, _, _ := procDwmInvalidateIconicBitmaps.Call(
		uintptr(hWnd))
	return HRESULT(ret)
}

func DwmIsCompositionEnabled(pfEnabled *BOOL) HRESULT {
	ret, _, _ := procDwmIsCompositionEnabled.Call(
		uintptr(unsafe.Pointer(pfEnabled)))
	return HRESULT(ret)
}

func DwmModifyPreviousDxFrameDuration(hWnd HWND, cRefreshes int, fRelative bool) HRESULT {
	ret, _, _ := procDwmModifyPreviousDxFrameDuration.Call(
		uintptr(hWnd),
		uintptr(cRefreshes),
		uintptr(BoolToBOOL(fRelative)))
	return HRESULT(ret)
}

func DwmQueryThumbnailSourceSize(hThumbnail HTHUMBNAIL, pSize *SIZE) HRESULT {
	ret, _, _ := procDwmQueryThumbnailSourceSize.Call(
		uintptr(hThumbnail),
		uintptr(unsafe.Pointer(pSize)))
	return HRESULT(ret)
}

func DwmRegisterThumbnail(hWndDestination HWND, hWndSource HWND, phThumbnailId *HTHUMBNAIL) HRESULT {
	ret, _, _ := procDwmRegisterThumbnail.Call(
		uintptr(hWndDestination),
		uintptr(hWndSource),
		uintptr(unsafe.Pointer(phThumbnailId)))
	return HRESULT(ret)
}

func DwmRenderGesture(gt GESTURE_TYPE, cContacts uint, pdwPointerID *uint32, pPoints *POINT) {
	procDwmRenderGesture.Call(
		uintptr(gt),
		uintptr(cContacts),
		uintptr(unsafe.Pointer(pdwPointerID)),
		uintptr(unsafe.Pointer(pPoints)))
	return
}

func DwmSetDxFrameDuration(hWnd HWND, cRefreshes int) HRESULT {
	ret, _, _ := procDwmSetDxFrameDuration.Call(
		uintptr(hWnd),
		uintptr(cRefreshes))
	return HRESULT(ret)
}

func DwmSetIconicLivePreviewBitmap(hWnd HWND, hbmp HBITMAP, pptClient *POINT, dwSITFlags uint32) HRESULT {
	ret, _, _ := procDwmSetIconicLivePreviewBitmap.Call(
		uintptr(hWnd),
		uintptr(hbmp),
		uintptr(unsafe.Pointer(pptClient)),
		uintptr(dwSITFlags))
	return HRESULT(ret)
}

func DwmSetIconicThumbnail(hWnd HWND, hbmp HBITMAP, dwSITFlags uint32) HRESULT {
	ret, _, _ := procDwmSetIconicThumbnail.Call(
		uintptr(hWnd),
		uintptr(hbmp),
		uintptr(dwSITFlags))
	return HRESULT(ret)
}

func DwmSetPresentParameters(hWnd HWND, pPresentParams *DWM_PRESENT_PARAMETERS) HRESULT {
	ret, _, _ := procDwmSetPresentParameters.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(pPresentParams)))
	return HRESULT(ret)
}

func DwmSetWindowAttribute(hWnd HWND, dwAttribute uint32, pvAttribute LPCVOID, cbAttribute uint32) HRESULT {
	ret, _, _ := procDwmSetWindowAttribute.Call(
		uintptr(hWnd),
		uintptr(dwAttribute),
		uintptr(pvAttribute),
		uintptr(cbAttribute))
	return HRESULT(ret)
}

func DwmShowContact(dwPointerID uint32, eShowContact DWM_SHOWCONTACT) {
	procDwmShowContact.Call(
		uintptr(dwPointerID),
		uintptr(eShowContact))
	return
}

func DwmTetherContact(dwPointerID uint32, fEnable bool, ptTether POINT) {
	procDwmTetherContact.Call(
		uintptr(dwPointerID),
		uintptr(BoolToBOOL(fEnable)),
		uintptr(unsafe.Pointer(&ptTether)))
	return
}

func DwmTransitionOwnedWindow(hWnd HWND, target DWMTRANSITION_OWNEDWINDOW_TARGET) {
	procDwmTransitionOwnedWindow.Call(
		uintptr(hWnd),
		uintptr(target))
	return
}

func DwmUnregisterThumbnail(hThumbnailId HTHUMBNAIL) HRESULT {
	ret, _, _ := procDwmUnregisterThumbnail.Call(
		uintptr(hThumbnailId))
	return HRESULT(ret)
}

func DwmUpdateThumbnailProperties(hThumbnailId HTHUMBNAIL, ptnProperties *DWM_THUMBNAIL_PROPERTIES) HRESULT {
	ret, _, _ := procDwmUpdateThumbnailProperties.Call(
		uintptr(hThumbnailId),
		uintptr(unsafe.Pointer(ptnProperties)))
	return HRESULT(ret)
}
