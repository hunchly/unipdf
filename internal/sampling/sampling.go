//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_c "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/imageutil";_d "io";);func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _ec []uint32 ;_bbf :=bitsPerOutputSample ;var _bd uint32 ;var _db uint32 ;_cfg :=0;_bcd :=0;_da :=0;for _da < len (data ){if _cfg > 0{_ef :=_cfg ;if _bbf < _ef {_ef =_bbf ;};_bd =(_bd <<uint (_ef ))|uint32 (_db >>uint (bitsPerInputSample -_ef ));_cfg -=_ef ;if _cfg > 0{_db =_db <<uint (_ef );}else {_db =0;};_bbf -=_ef ;if _bbf ==0{_ec =append (_ec ,_bd );_bbf =bitsPerOutputSample ;_bd =0;_bcd ++;};}else {_fbf :=data [_da ];_da ++;_cec :=bitsPerInputSample ;if _bbf < _cec {_cec =_bbf ;};_cfg =bitsPerInputSample -_cec ;_bd =(_bd <<uint (_cec ))|uint32 (_fbf >>uint (_cfg ));if _cec < bitsPerInputSample {_db =_fbf <<uint (_cec );};_bbf -=_cec ;if _bbf ==0{_ec =append (_ec ,_bd );_bbf =bitsPerOutputSample ;_bd =0;_bcd ++;};};};for _cfg >=bitsPerOutputSample {_fa :=_cfg ;if _bbf < _fa {_fa =_bbf ;};_bd =(_bd <<uint (_fa ))|uint32 (_db >>uint (bitsPerInputSample -_fa ));_cfg -=_fa ;if _cfg > 0{_db =_db <<uint (_fa );}else {_db =0;};_bbf -=_fa ;if _bbf ==0{_ec =append (_ec ,_bd );_bbf =bitsPerOutputSample ;_bd =0;_bcd ++;};};if _bbf > 0&&_bbf < bitsPerOutputSample {_bd <<=uint (_bbf );_ec =append (_ec ,_bd );};return _ec ;};type Writer struct{_gf _a .ImageBase ;_dff *_c .Writer ;_agbe ,_eg int ;_cg bool ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_gc []uint32 )error ;};type SampleWriter interface{WriteSample (_bg uint32 )error ;WriteSamples (_ac []uint32 )error ;};func NewReader (img _a .ImageBase )*Reader {return &Reader {_f :_c .NewReader (img .Data ),_ad :img ,_fg :img .ColorComponents ,_ag :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _e []uint32 ;_cf :=bitsPerSample ;var _fb uint32 ;var _dfd byte ;_cdg :=0;_ed :=0;_cc :=0;for _cc < len (data ){if _cdg > 0{_gd :=_cdg ;if _cf < _gd {_gd =_cf ;};_fb =(_fb <<uint (_gd ))|uint32 (_dfd >>uint (8-_gd ));_cdg -=_gd ;if _cdg > 0{_dfd =_dfd <<uint (_gd );}else {_dfd =0;};_cf -=_gd ;if _cf ==0{_e =append (_e ,_fb );_cf =bitsPerSample ;_fb =0;_ed ++;};}else {_eb :=data [_cc ];_cc ++;_ff :=8;if _cf < _ff {_ff =_cf ;};_cdg =8-_ff ;_fb =(_fb <<uint (_ff ))|uint32 (_eb >>uint (_cdg ));if _ff < 8{_dfd =_eb <<uint (_ff );};_cf -=_ff ;if _cf ==0{_e =append (_e ,_fb );_cf =bitsPerSample ;_fb =0;_ed ++;};};};for _cdg >=bitsPerSample {_cea :=_cdg ;if _cf < _cea {_cea =_cf ;};_fb =(_fb <<uint (_cea ))|uint32 (_dfd >>uint (8-_cea ));_cdg -=_cea ;if _cdg > 0{_dfd =_dfd <<uint (_cea );}else {_dfd =0;};_cf -=_cea ;if _cf ==0{_e =append (_e ,_fb );_cf =bitsPerSample ;_fb =0;_ed ++;};};return _e ;};func (_dbg *Writer )WriteSample (sample uint32 )error {if _ ,_ea :=_dbg ._dff .WriteBits (uint64 (sample ),_dbg ._gf .BitsPerComponent );_ea !=nil {return _ea ;};_dbg ._eg --;if _dbg ._eg ==0{_dbg ._eg =_dbg ._gf .ColorComponents ;_dbg ._agbe ++;};if _dbg ._agbe ==_dbg ._gf .Width {if _dbg ._cg {_dbg ._dff .FinishByte ();};_dbg ._agbe =0;};return nil ;};func (_fc *Writer )WriteSamples (samples []uint32 )error {for _fe :=0;_fe < len (samples );_fe ++{if _ffd :=_fc .WriteSample (samples [_fe ]);_ffd !=nil {return _ffd ;};};return nil ;};func (_af *Reader )ReadSample ()(uint32 ,error ){if _af ._cd ==_af ._ad .Height {return 0,_d .EOF ;};_df ,_bb :=_af ._f .ReadBits (byte (_af ._ad .BitsPerComponent ));if _bb !=nil {return 0,_bb ;};_af ._fg --;if _af ._fg ==0{_af ._fg =_af ._ad .ColorComponents ;_af ._b ++;};if _af ._b ==_af ._ad .Width {if _af ._ag {_af ._f .ConsumeRemainingBits ();};_af ._b =0;_af ._cd ++;};return uint32 (_df ),nil ;};type Reader struct{_ad _a .ImageBase ;_f *_c .Reader ;_b ,_cd ,_fg int ;_ag bool ;};func (_bc *Reader )ReadSamples (samples []uint32 )(_agb error ){for _ce :=0;_ce < len (samples );_ce ++{samples [_ce ],_agb =_bc .ReadSample ();if _agb !=nil {return _agb ;};};return nil ;};func NewWriter (img _a .ImageBase )*Writer {return &Writer {_dff :_c .NewWriterMSB (img .Data ),_gf :img ,_eg :img .ColorComponents ,_cg :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};