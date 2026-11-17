package win32

//go:generate mkwinsyscall -output d2d1_functions.go d2d1.go

// Functions that set last error:
/*
 */

// Functions that do not set last error:
/*
//sys D2D1CreateFactory(factoryType e.D2D1_FACTORY_TYPE, riid *b.IID, pFactoryOptions *t.D2D1_FACTORY_OPTIONS, ppIFactory **i.ID2D1Factory) (ret b.HRESULT) = d2d1.D2D1CreateFactory
*/
