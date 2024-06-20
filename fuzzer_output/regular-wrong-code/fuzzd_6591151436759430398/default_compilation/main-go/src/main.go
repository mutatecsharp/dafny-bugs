// Dafny program main.dfy compiled into Go
package main

import (
	_System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_module.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Abs(x _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return (_dafny.IntOfInt64(-1)).Times(x)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeIndex(x _dafny.Int, length _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return _dafny.Zero
	} else if (x).Cmp(length) >= 0 {
		return (x).Modulo(length)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeDivisionInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).DivBy(x2)
	}
}
func (_static *CompanionStruct_Default___) SafeModuloInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).Modulo(x2)
	}
}
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, globalState *GlobalState) bool {
	return (true) == ((!(false)) && (!(true)))
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("sxninyjx")
}
func (_static *CompanionStruct_Default___) Fm2(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('s')
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(332)
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.MultiSetOf(false, false)), Companion_D1_.Create_DC3_(_dafny.MultiSetOf(true))), _dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(false))))), Companion_D1_.Create_DC3_(_dafny.MultiSetOf(true, false)), Companion_D1_.Create_DC3_(_dafny.MultiSetFromSeq(_dafny.SeqOf(true)))))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SeqOf(true, !(!(true))), _dafny.SeqOf(true))).Difference(_dafny.MultiSetOf(_dafny.SeqOf(true), _dafny.SeqOf(false, false, true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(960))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.Int {
		return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lheve")).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
	}))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("psg")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("wufuwmsuw")
	})), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("pfu"))))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(146)).Times((_dafny.MultiSetOf(false, !(false))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm20(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(!(false))), true), _dafny.SeqOf(false, false))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC12_(true, !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-894), _dafny.SeqOf(_dafny.IntOfInt64(283)))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(328), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality())))), ((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(878), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("du")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(241))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _3_v0 _dafny.Int
			_3_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(241)), _3_v0) {
				_coll0.Add((_3_v0).Plus(_dafny.IntOfInt64(-99)))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality())).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pnlbg")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC20_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(false))), _dafny.SeqOf(true)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.MultiSet, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D11_.Create_DC31_(false, _dafny.IntOfInt64(948), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-910))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), true)).Dtor_cf56(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(!(true))))))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC23_()
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetFromSeq((Companion_D14_.Create_DC38_(_dafny.SeqOf(Companion_D8_.Create_DC23_()))).Dtor_cf70())).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 D8
			_5_v0 = interface{}(_compr_1).(D8)
			if (_dafny.MultiSetFromSeq((Companion_D14_.Create_DC38_(_dafny.SeqOf(Companion_D8_.Create_DC23_()))).Dtor_cf70())).Contains(_5_v0) {
				_coll1.Add(_5_v0)
			}
		}
		return _coll1.ToSet()
	}()).Union(_dafny.SetOf(Companion_D8_.Create_DC23_())), _dafny.SetOf(Companion_D8_.Create_DC23_()), (_dafny.SetOf(Companion_D8_.Create_DC23_(), Companion_D8_.Create_DC23_(), Companion_D8_.Create_DC23_())).Union(_dafny.SetOf(Companion_D8_.Create_DC23_(), Companion_D8_.Create_DC23_(), Companion_D8_.Create_DC23_())))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) D8 {
	if (_dafny.SetOf(_dafny.SeqOf(true))).Contains(_dafny.SeqOf(!(true))) {
		return Companion_D8_.Create_DC21_((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(324))).Cardinality()))).Cardinality(), true, !(true), true, _dafny.IntOfInt64(194))
	} else {
		return Companion_D8_.Create_DC21_((_dafny.SetOf(_dafny.IntOfInt64(-729), (_dafny.Zero).Minus(_dafny.IntOfInt64(-109)))).Cardinality(), false, true, true, _dafny.IntOfInt64(-579))
	}
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('n'))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 _dafny.CodePoint
			_6_v0 = interface{}(_compr_2).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('n')), _6_v0) {
				_coll2.Add(_6_v0)
			}
		}
		return _coll2.ToSet()
	}()).Union((_dafny.SetOf(_dafny.CodePoint('v'))).Difference(_dafny.SetOf(_dafny.CodePoint('n'))))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(708), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bnk")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(802), false)).Cardinality())).Cardinality()), _dafny.IntOfInt64(475)), func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(-250))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _7_v0 _dafny.Int
			_7_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfInt64(-250))).Contains(_7_v0) {
				_coll3.Add((_7_v0).Plus((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('l'))).Cardinality()))).Cardinality()))
			}
		}
		return _coll3.ToSet()
	}())).Union((_dafny.SetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(769))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(223), _dafny.IntOfInt64(890))); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _9_v1 _dafny.Int
				_9_v1 = interface{}(_compr_4).(_dafny.Int)
				if ((_dafny.IntOfInt64(223)).Cmp(_9_v1) <= 0) && ((_9_v1).Cmp(_dafny.IntOfInt64(890)) < 0) {
					_coll4.Add(Companion_Default___.SafeDivisionInt(_9_v1, _dafny.IntOfInt64(291)))
				}
			}
			return _coll4.ToSet()
		}()).Cardinality())
	}))).Cardinality())))).Union(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-278), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(436))).Cardinality())), _dafny.SetOf((func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()), true)).Keys().Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _10_v2 _dafny.Int
				_10_v2 = interface{}(_compr_6).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()), true)).Contains(_10_v2) {
					_coll6.Add((_10_v2).Times(_dafny.IntOfInt64(15)), true)
				}
			}
			return _coll6.ToMap()
		}()).Cardinality()))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _11_v3 _dafny.Set
			_11_v3 = interface{}(_compr_5).(_dafny.Set)
			if (_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-278), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(436))).Cardinality())), _dafny.SetOf((func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()), true)).Keys().Elements()); ; {
					_compr_7, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _10_v2 _dafny.Int
					_10_v2 = interface{}(_compr_7).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()), true)).Contains(_10_v2) {
						_coll7.Add((_10_v2).Times(_dafny.IntOfInt64(15)), true)
					}
				}
				return _coll7.ToMap()
			}()).Cardinality()))).Contains(_11_v3) {
				_coll5.Add(_11_v3)
			}
		}
		return _coll5.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	var _source0 D9 = Companion_D9_.Create_DC27_(false, _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'), _dafny.CodePoint('v'))).Cardinality())))
	_ = _source0
	if _source0.Is_DC26() {
		var _12___mcc_h0 _dafny.Int = _source0.Get_().(D9_DC26).Cf45
		_ = _12___mcc_h0
		var _13___mcc_h1 *C0 = _source0.Get_().(D9_DC26).Cf46
		_ = _13___mcc_h1
		var _14_cf46 *C0 = _13___mcc_h1
		_ = _14_cf46
		var _15_cf45 _dafny.Int = _12___mcc_h0
		_ = _15_cf45
		return _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efkgpwsh")).Cardinality())), func() _dafny.Set {
			var _coll8 = _dafny.NewBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(809), _dafny.IntOfInt64(957))); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _16_v0 _dafny.Int
				_16_v0 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(809)).Cmp(_16_v0) <= 0) && ((_16_v0).Cmp(_dafny.IntOfInt64(957)) < 0) {
					_coll8.Add(Companion_Default___.SafeModuloInt(_16_v0, (_14_cf46).F26()))
				}
			}
			return _coll8.ToSet()
		}())
	} else if _source0.Is_DC27() {
		var _17___mcc_h2 bool = _source0.Get_().(D9_DC27).Cf47
		_ = _17___mcc_h2
		var _18___mcc_h3 _dafny.Set = _source0.Get_().(D9_DC27).Cf48
		_ = _18___mcc_h3
		var _19_cf48 _dafny.Set = _18___mcc_h3
		_ = _19_cf48
		var _20_cf47 bool = _17___mcc_h2
		_ = _20_cf47
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_19_cf48), _dafny.SeqOf(_19_cf48))
	} else {
		var _21___mcc_h4 _dafny.Map = _source0.Get_().(D9_DC25).Cf44
		_ = _21___mcc_h4
		var _22_cf44 _dafny.Map = _21___mcc_h4
		_ = _22_cf44
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_23_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality()))).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _24_v1 _dafny.Int
				_24_v1 = interface{}(_compr_9).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_23_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality())), _24_v1) {
					_coll9.Add((_24_v1).Minus(_dafny.IntOfInt64(-706)))
				}
			}
			return _coll9.ToSet()
		}()), _dafny.SeqOf(func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(836), _dafny.IntOfInt64(-360))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _25_v2 _dafny.Int
				_25_v2 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(836)).Cmp(_25_v2) <= 0) && ((_25_v2).Cmp(_dafny.IntOfInt64(-360)) < 0) {
					_coll10.Add(Companion_Default___.SafeModuloInt(_25_v2, _dafny.IntOfInt64(13)))
				}
			}
			return _coll10.ToSet()
		}()))
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 _dafny.Sequence, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-865))).Uint32(), func(coer7 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}(func(_26_i0 _dafny.Int) D1 {
				return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(320), true)
			}))
		}
		return _dafny.SeqOf(Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dwmmy")).Cardinality()), false), Companion_D1_.Create_DC5_(_dafny.IntOfInt64(301), true))
	})())).Difference(_dafny.MultiSetOf(Companion_D1_.Create_DC5_(_dafny.IntOfInt64(-624), false), Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(154))).Difference(func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(530), false)).Keys().Elements()); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _27_v0 _dafny.Int
			_27_v0 = interface{}(_compr_11).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(530), false)).Contains(_27_v0) {
				_coll11.Add(Companion_Default___.SafeDivisionInt(_27_v0, (func() _dafny.Set {
					var _coll12 = _dafny.NewBuilder()
					_ = _coll12
					for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.UnicodeSeqOfUtf8Bytes("bnkgmf"))).Keys().Elements()); ; {
						_compr_12, _ok12 := _iter12()
						if !_ok12 {
							break
						}
						var _28_v1 _dafny.CodePoint
						_28_v1 = interface{}(_compr_12).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.UnicodeSeqOfUtf8Bytes("bnkgmf"))).Contains(_28_v1) {
							_coll12.Add(_28_v1)
						}
					}
					return _coll12.ToSet()
				}()).Cardinality()))
			}
		}
		return _coll11.ToSet()
	}())).Union(_dafny.SetOf(_dafny.IntOfInt64(881)))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	if !(!(!(false) || (true))) {
		return _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ehocjrafq")).Cardinality())))
	} else {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(790))).Uint32(), func(coer8 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_29_i0 _dafny.Int) _dafny.MultiSet {
			return _dafny.MultiSetOf((_dafny.MultiSetOf(true, (Companion_D8_.Create_DC21_((_dafny.MultiSetOf(_dafny.IntOfInt64(387))).Cardinality(), true, true, !(true), _dafny.IntOfInt64(990))).Dtor_cf36(), true)).Cardinality())
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Sequence, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_((_dafny.IntOfInt64(-221)).Cmp((Companion_D11_.Create_DC30_(false, _dafny.IntOfInt64(-315))).Dtor_cf52()) != 0, (_dafny.IntOfInt64(-347)).Plus((_dafny.MultiSetOf(true)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_((_dafny.MultiSetOf(false)).Cardinality(), !(false))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll13 = _dafny.NewBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-241), _dafny.IntOfInt64(396))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _30_v0 _dafny.Int
				_30_v0 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(-241)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(396)) < 0) {
					_coll14.Add(Companion_Default___.SafeDivisionInt(_30_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xwfabehw")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality())))
				}
			}
			return _coll14.ToMap()
		}()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-954)), func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-481), _dafny.IntOfInt64(89))); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _32_v1 _dafny.Int
				_32_v1 = interface{}(_compr_15).(_dafny.Int)
				if ((_dafny.IntOfInt64(-481)).Cmp(_32_v1) <= 0) && ((_32_v1).Cmp(_dafny.IntOfInt64(89)) < 0) {
					_coll15.Add((_32_v1).Times((_dafny.SetOf(_dafny.IntOfInt64(741))).Cardinality()), _dafny.IntOfInt64(-324))
				}
			}
			return _coll15.ToMap()
		}()))).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _33_v2 _dafny.Map
			_33_v2 = interface{}(_compr_13).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-241), _dafny.IntOfInt64(396))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _30_v0 _dafny.Int
					_30_v0 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-241)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(396)) < 0) {
						_coll16.Add(Companion_Default___.SafeDivisionInt(_30_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xwfabehw")).Cardinality())), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg10 _dafny.Int) interface{} {
								return coer10(arg10)
							}
						}(func(_31_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('l')
						}))).Cardinality())))
					}
				}
				return _coll16.ToMap()
			}()), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-954)), func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-481), _dafny.IntOfInt64(89))); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _32_v1 _dafny.Int
					_32_v1 = interface{}(_compr_17).(_dafny.Int)
					if ((_dafny.IntOfInt64(-481)).Cmp(_32_v1) <= 0) && ((_32_v1).Cmp(_dafny.IntOfInt64(89)) < 0) {
						_coll17.Add((_32_v1).Times((_dafny.SetOf(_dafny.IntOfInt64(741))).Cardinality()), _dafny.IntOfInt64(-324))
					}
				}
				return _coll17.ToMap()
			}())), _33_v2) {
				_coll13.Add(_33_v2)
			}
		}
		return _coll13.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(!(false)))).Union(_dafny.MultiSetOf(true, !(false)))).Difference(_dafny.MultiSetOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(893))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_34_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('r')
	})))
}
func (_static *CompanionStruct_Default___) Fm38(p0 D9, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), true)).Contains(_dafny.SetOf(false)) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_35_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(412)
		}))).Cardinality()), !(true))).Merge(func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(918), _dafny.IntOfInt64(-354))); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _36_v0 _dafny.Int
				_36_v0 = interface{}(_compr_18).(_dafny.Int)
				if ((_dafny.IntOfInt64(918)).Cmp(_36_v0) <= 0) && ((_36_v0).Cmp(_dafny.IntOfInt64(-354)) < 0) {
					_coll18.Add(Companion_Default___.SafeModuloInt(_36_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg13 _dafny.Int) interface{} {
							return coer13(arg13)
						}
					}(func(_37_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(232))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg14 _dafny.Int) interface{} {
								return coer14(arg14)
							}
						}(func(_38_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}))).Cardinality())
					}))).Cardinality())), !(true))
				}
			}
			return _coll18.ToMap()
		}())
	} else {
		return (Companion_D15_.Create_DC40_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC12_(true, false, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-207))).Cardinality()), _dafny.IntOfInt64(640))).Cardinality())).Dtor_cf23(), true))).Dtor_cf73()
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, p1 _dafny.Sequence, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("usfsg")).Cardinality())), (Companion_D16_.Create_DC44_(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-290))).Cardinality()))).Dtor_cf82(), (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(352), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(4))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(884), _dafny.IntOfInt64(762))).Cardinality())))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-430)))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(-57), (_dafny.Zero).Minus(_dafny.IntOfInt64(-119)))), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-953), _dafny.IntOfInt64(886))).Cardinality(), Companion_D15_.Create_DC41_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(483))).Cardinality())).Cardinality(), _dafny.SeqOf(_dafny.IntOfInt64(296)), _dafny.SeqOf(true), _dafny.UnicodeSeqOfUtf8Bytes("gp")))).Cardinality(), (_dafny.SetOf(_dafny.SetOf(true), _dafny.SetOf(false, (Companion_D6_.Create_DC16_(true, true)).Dtor_cf30()), _dafny.SetOf(false, true), _dafny.SetOf(false))).Cardinality(), _dafny.IntOfInt64(580), _dafny.IntOfInt64(128), (_dafny.MultiSetOf(_dafny.IntOfInt64(816), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-119)))).Cardinality()), _dafny.IntOfInt64(131))).Cardinality()))).Union(_dafny.MultiSetOf((func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(937), _dafny.IntOfInt64(-278))).Keys().Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _39_v0 _dafny.Int
			_39_v0 = interface{}(_compr_19).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(937), _dafny.IntOfInt64(-278))).Contains(_39_v0) {
				_coll19.Add((_39_v0).Minus(_dafny.IntOfInt64(313)))
			}
		}
		return _coll19.ToSet()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(204), false)).Cardinality(), _dafny.IntOfInt64(707), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(!(true)))).Cardinality())), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.CodePoint('l'), _dafny.CodePoint('g'))).Union((_dafny.MultiSetOf(_dafny.CodePoint('l'), _dafny.CodePoint('r'), _dafny.CodePoint('v'), _dafny.CodePoint('f'), _dafny.CodePoint('t'))).Union(_dafny.MultiSetOf(_dafny.CodePoint('m'), _dafny.CodePoint('o'))))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) {
	var _40_v0 bool
	_ = _40_v0
	_40_v0 = false
	if _40_v0 {
		var _41_v1 _dafny.Int
		_ = _41_v1
		_41_v1 = _dafny.IntOfInt64(9)
		var _42_v2 _dafny.Map
		_ = _42_v2
		_42_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(Companion_Default___.Fm3(globalState)), _dafny.SeqOf(_41_v1, p2)), _40_v0)
		_41_v1 = (_42_v2).Cardinality()
		var _43_v3 _dafny.Map
		_ = _43_v3
		_43_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_41_v1, _dafny.IntOfInt64(469)), !(_40_v0))
		var _44_v4 _dafny.Sequence
		_ = _44_v4
		_44_v4 = _dafny.SeqOf((_43_v3).Cardinality(), _41_v1, (_dafny.Zero).Minus(_41_v1))
		var _45_v5 _dafny.Set
		_ = _45_v5
		_45_v5 = _dafny.SetOf(_44_v4, _dafny.SeqOf(_41_v1), _44_v4)
		_45_v5 = _45_v5
		var _46_v6 _dafny.MultiSet
		_ = _46_v6
		_46_v6 = _dafny.MultiSetOf(p1, p1)
		var _47_v7 _dafny.CodePoint
		_ = _47_v7
		_47_v7 = _dafny.CodePoint('m')
		_46_v6 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})), _dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _47_v7))
		var _49_v8 _dafny.Array
		_ = _49_v8
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_0
		var _nw0 _dafny.Array
		_ = _nw0
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw0 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = (func(_50_v0 bool) func(_dafny.Int) bool {
				return func(_51_i1 _dafny.Int) bool {
					return _50_v0
				}
			})(_40_v0)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw0).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw0).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_49_v8 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_49_v8), 0))
		_ = _index0
		(_49_v8).ArraySet1(!(Companion_Default___.Fm0(_40_v0, _40_v0, globalState)), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_49_v8), 0))
		_ = _index1
		var _rhs0 bool = _40_v0
		_ = _rhs0
		var _rhs1 bool = true
		_ = _rhs1
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 _dafny.Array = _49_v8
		_ = _lhs1
		var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_49_v8), 0))
		_ = _lhs2
		_lhs0.F2 = _rhs0
		(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_49_v8), 0))
		_ = _index2
		(_49_v8).ArraySet1(Companion_Default___.Fm0(_dafny.Companion_Sequence_.IsPrefixOf(p1, p1), true, globalState), (_index2).Int())
	} else {
		var _52_v9 _dafny.Map
		_ = _52_v9
		_52_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-2))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_53_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), p2)
		if !(_52_v9).Equals(_52_v9) {
			var _54_v10 _dafny.Int
			_ = _54_v10
			_54_v10 = _dafny.IntOfInt64(844)
			var _55_v11 _dafny.Map
			_ = _55_v11
			_55_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, _40_v0)
			var _56_v12 D1
			_ = _56_v12
			_56_v12 = Companion_D1_.Create_DC4_(_40_v0, p2, (_55_v11).Update(_40_v0, _40_v0))
			var _57_v13 _dafny.MultiSet
			_ = _57_v13
			_57_v13 = _dafny.MultiSetOf((_56_v12).Dtor_cf7(), _40_v0)
			_54_v10 = (Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(323))).Minus(Companion_Default___.SafeModuloInt((_57_v13).Cardinality(), p0))
			var _58_v14 _dafny.Set
			_ = _58_v14
			_58_v14 = _dafny.SetOf(_40_v0)
			var _59_v15 _dafny.Sequence
			_ = _59_v15
			_59_v15 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()))
			var _60_v16 *C8
			_ = _60_v16
			var _nw1 *C8 = New_C8_()
			_ = _nw1
			_nw1.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(909))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}(func(_61_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			})), _58_v14, _40_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_59_v15, _59_v15), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_59_v15, _59_v15)).Cardinality()))).Uint32(), p3))
			_60_v16 = _nw1
			var _62_v17 _dafny.Array
			_ = _62_v17
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_1
			var _nw2 _dafny.Array
			_ = _nw2
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw2 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_63_v0 bool) func(_dafny.Int) bool {
					return func(_64_i4 _dafny.Int) bool {
						return _63_v0
					}
				})(_40_v0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw2 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw2).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw2).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_62_v17 = _nw2
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_62_v17), 0))
			_ = _index3
			(_62_v17).ArraySet1(_40_v0, (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(392), _dafny.ArrayLen((_62_v17), 0))
			_ = _index4
			(_62_v17).ArraySet1(_40_v0, (_index4).Int())
			var _65_v18 _dafny.Array
			_ = _65_v18
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw3
			_65_v18 = _nw3
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_65_v18), 0))
			_ = _index5
			(_65_v18).ArraySet1(p0, (_index5).Int())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(816), _dafny.ArrayLen((_65_v18), 0))
			_ = _index6
			(_65_v18).ArraySet1((Companion_Default___.Fm3(globalState)).Minus(p3), (_index6).Int())
			var _66_v19 *C2
			_ = _66_v19
			var _nw4 *C2 = New_C2_()
			_ = _nw4
			_nw4.Ctor__(_dafny.IntOfUint32((_59_v15).Cardinality()))
			_66_v19 = _nw4
		} else {
			var _67_v20 _dafny.Array
			_ = _67_v20
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(18))
			_ = _nw5
			_67_v20 = _nw5
			var _68_v21 _dafny.Array
			_ = _68_v21
			_68_v21 = _67_v20
			_68_v21 = (func() _dafny.Array {
				if _40_v0 {
					return _67_v20
				}
				return _67_v20
			})()
			var _69_v22 _dafny.Sequence
			_ = _69_v22
			_69_v22 = _dafny.SeqOf(_dafny.IntOfInt64(672))
			var _70_v23 _dafny.Array
			_ = _70_v23
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_2
			var _nw6 _dafny.Array
			_ = _nw6
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw6 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_71_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_72_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_72_i6, _dafny.IntOfUint32((_71_p1).Cardinality()))
					}
				})(p1)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw6 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw6).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw6).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_70_v23 = _nw6
			var _73_v24 _dafny.Sequence
			_ = _73_v24
			_73_v24 = _dafny.SeqOf(_70_v23)
			var _74_v25 _dafny.Array
			_ = _74_v25
			var _nwElement0_0 _dafny.Sequence = _69_v22
			_ = _nwElement0_0
			var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(13))
			_ = _nw7
			(_nw7).ArraySet1(_nwElement0_0, 0)
			(_nw7).ArraySet1(_69_v22, 1)
			(_nw7).ArraySet1(_69_v22, 2)
			(_nw7).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(102))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_75_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_76_i5 _dafny.Int) _dafny.Int {
					return _75_p0
				}
			})(p0))), 3)
			(_nw7).ArraySet1(_69_v22, 4)
			(_nw7).ArraySet1(_69_v22, 5)
			(_nw7).ArraySet1(_dafny.SeqOf(p2, _dafny.IntOfUint32((_73_v24).Cardinality()), p2, p2, p0), 6)
			(_nw7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_69_v22, Companion_Default___.Fm17(globalState)), 7)
			(_nw7).ArraySet1(Companion_Default___.Fm17(globalState), 8)
			(_nw7).ArraySet1(_69_v22, 9)
			(_nw7).ArraySet1(_69_v22, 10)
			(_nw7).ArraySet1(_69_v22, 11)
			(_nw7).ArraySet1(_69_v22, 12)
			_74_v25 = _nw7
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_74_v25), 0))
			_ = _index7
			(_74_v25).ArraySet1(_dafny.SeqOf((_dafny.SetOf(_40_v0)).Cardinality()), (_index7).Int())
			var _77_v26 _dafny.Map
			_ = _77_v26
			_77_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_40_v0), _69_v22)
			var _78_v27 _dafny.Sequence
			_ = _78_v27
			_78_v27 = _dafny.SeqOf((func() _dafny.Sequence {
				if (_77_v26).Contains(_40_v0) {
					return (_77_v26).Get(_40_v0).(_dafny.Sequence)
				}
				return _dafny.SeqOf(p2, p0, p2)
			})())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_74_v25), 0))
			_ = _index8
			(_74_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm17(globalState), (_78_v27).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_78_v27).Cardinality()))).Uint32()).(_dafny.Sequence)), (_index8).Int())
			var _79_v28 _dafny.Int
			_ = _79_v28
			_79_v28 = _dafny.IntOfInt64(907)
			_79_v28 = p2
			var _80_v29 _dafny.MultiSet
			_ = _80_v29
			_80_v29 = _dafny.MultiSetOf(_79_v28)
			(globalState).F2 = ((_80_v29).Difference(_dafny.MultiSetOf(p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(327))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_81_p2 _dafny.Int, _82_v0 bool) func(_dafny.Int) _dafny.Int {
				return func(_83_i7 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_p2, _82_v0)).Cardinality()
				}
			})(p2, _40_v0)))).Cardinality())))).IsSubsetOf(_dafny.MultiSetFromSeq((_74_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_74_v25), 0))).Int()).(_dafny.Sequence)))
			_79_v28 = p0
		}
		var _84_v30 _dafny.Array
		_ = _84_v30
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_3
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_85_v0 bool) func(_dafny.Int) bool {
				return func(_86_i8 _dafny.Int) bool {
					return !(false) || (_85_v0)
				}
			})(_40_v0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw8 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw8).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw8).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_84_v30 = _nw8
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_84_v30), 0))
		_ = _index9
		(_84_v30).ArraySet1(_40_v0, (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_84_v30), 0))
		_ = _index10
		(_84_v30).ArraySet1(_40_v0, (_index10).Int())
		var _87_v31 T0
		_ = _87_v31
		var _nw9 *C5 = New_C5_()
		_ = _nw9
		_nw9.Ctor__(_40_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_88_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_89_i9 _dafny.Int) _dafny.Int {
				return _88_p0
			}
		})(p0))))
		_87_v31 = _nw9
		var _90_v32 _dafny.CodePoint
		_ = _90_v32
		_90_v32 = _dafny.CodePoint('m')
		var _91_v33 _dafny.Map
		_ = _91_v33
		_91_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_87_v31, _90_v32)
		(globalState).F2 = !(_91_v33).Contains(_87_v31)
		var _92_v34 *C5
		_ = _92_v34
		var _nw10 *C5 = New_C5_()
		_ = _nw10
		_nw10.Ctor__(!((_84_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_84_v30), 0))).Int()).(bool)) || (_87_v31.F16()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg21 _dafny.Int) interface{} {
				return coer21(arg21)
			}
		}((func(_93_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_94_i10 _dafny.Int) _dafny.Int {
				return _93_p0
			}
		})(p0))))
		_92_v34 = _nw10
		_90_v32 = _90_v32
	}
	var _95_v35 _dafny.Int
	_ = _95_v35
	_95_v35 = _dafny.IntOfInt64(-579)
	var _96_v36 D4
	_ = _96_v36
	_96_v36 = Companion_D4_.Create_DC12_(true, false, p2)
	var _pat_let_tv0 = p3
	_ = _pat_let_tv0
	var _pat_let_tv1 = p3
	_ = _pat_let_tv1
	var _pat_let_tv2 = _40_v0
	_ = _pat_let_tv2
	var _pat_let_tv3 = _40_v0
	_ = _pat_let_tv3
	var _pat_let_tv4 = _40_v0
	_ = _pat_let_tv4
	var _pat_let_tv5 = _40_v0
	_ = _pat_let_tv5
	var _pat_let_tv6 = _40_v0
	_ = _pat_let_tv6
	_95_v35 = func(_source1 D4) _dafny.Int {
		if _source1.Is_DC12() {
			var _97___mcc_h0 bool = _source1.Get_().(D4_DC12).Cf21
			_ = _97___mcc_h0
			var _98___mcc_h1 bool = _source1.Get_().(D4_DC12).Cf22
			_ = _98___mcc_h1
			var _99___mcc_h2 _dafny.Int = _source1.Get_().(D4_DC12).Cf23
			_ = _99___mcc_h2
			var _100_cf23 _dafny.Int = _99___mcc_h2
			_ = _100_cf23
			var _101_cf22 bool = _98___mcc_h1
			_ = _101_cf22
			var _102_cf21 bool = _97___mcc_h0
			_ = _102_cf21
			return _pat_let_tv0
		} else if _source1.Is_DC13() {
			var _103___mcc_h3 _dafny.Int = _source1.Get_().(D4_DC13).Cf24
			_ = _103___mcc_h3
			var _104___mcc_h4 _dafny.Int = _source1.Get_().(D4_DC13).Cf25
			_ = _104___mcc_h4
			var _105___mcc_h5 _dafny.Int = _source1.Get_().(D4_DC13).Cf26
			_ = _105___mcc_h5
			var _106_cf26 _dafny.Int = _105___mcc_h5
			_ = _106_cf26
			var _107_cf25 _dafny.Int = _104___mcc_h4
			_ = _107_cf25
			var _108_cf24 _dafny.Int = _103___mcc_h3
			_ = _108_cf24
			return (_dafny.Zero).Minus(_pat_let_tv1)
		} else {
			var _109___mcc_h6 _dafny.MultiSet = _source1.Get_().(D4_DC11).Cf20
			_ = _109___mcc_h6
			var _110_cf20 _dafny.MultiSet = _109___mcc_h6
			_ = _110_cf20
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_pat_let_tv2, !(_pat_let_tv3), _pat_let_tv4), _dafny.SeqOf(_pat_let_tv5, _pat_let_tv6))).Cardinality()))
		}
	}(_96_v36)
	var _111_v37 _dafny.CodePoint
	_ = _111_v37
	_111_v37 = _dafny.CodePoint('j')
	var _112_v38 _dafny.Sequence
	_ = _112_v38
	_112_v38 = _dafny.SeqOf(_40_v0)
	var _113_v39 _dafny.Map
	_ = _113_v39
	_113_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v37, _112_v38)
	var _114_v40 _dafny.Sequence
	_ = _114_v40
	_114_v40 = _dafny.SeqOf(_112_v38)
	_113_v39 = (_113_v39).Update(_111_v37, _dafny.Companion_Sequence_.Concatenate(_112_v38, (_114_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dsxwg")).Cardinality()), _dafny.IntOfUint32((_114_v40).Cardinality()))).Uint32()).(_dafny.Sequence)))
	var _115_v41 _dafny.Map
	_ = _115_v41
	_115_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
	var _116_v42 _dafny.Sequence
	_ = _116_v42
	_116_v42 = _dafny.SeqOf(_95_v35, p3)
	var _117_v43 *C3
	_ = _117_v43
	var _nw11 *C3 = New_C3_()
	_ = _nw11
	_nw11.Ctor__(p1, _40_v0, _116_v42)
	_117_v43 = _nw11
	var _118_v44 D12
	_ = _118_v44
	_118_v44 = Companion_D12_.Create_DC35_(_40_v0, _40_v0, _115_v41, true, _117_v43)
	var _119_v45 _dafny.Array
	_ = _119_v45
	var _nwElement0_1 bool = _40_v0
	_ = _nwElement0_1
	var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(12))
	_ = _nw12
	(_nw12).ArraySet1(_nwElement0_1, 0)
	(_nw12).ArraySet1((_118_v44).Dtor_cf64(), 1)
	(_nw12).ArraySet1(_40_v0, 2)
	(_nw12).ArraySet1(_40_v0, 3)
	(_nw12).ArraySet1(_40_v0, 4)
	(_nw12).ArraySet1(_40_v0, 5)
	(_nw12).ArraySet1(_40_v0, 6)
	(_nw12).ArraySet1(_40_v0, 7)
	(_nw12).ArraySet1(_40_v0, 8)
	(_nw12).ArraySet1(_40_v0, 9)
	(_nw12).ArraySet1(_40_v0, 10)
	(_nw12).ArraySet1(false, 11)
	_119_v45 = _nw12
	var _120_v46 D0
	_ = _120_v46
	_120_v46 = Companion_D0_.Create_DC2_((_dafny.MultiSetFromSeq(_112_v38)).Cardinality(), p0, p3, _119_v45)
	var _121_v47 *C6
	_ = _121_v47
	var _nw13 *C6 = New_C6_()
	_ = _nw13
	_nw13.Ctor__(func(_pat_let0_0 D0) D0 {
		return func(_122_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Int) D0 {
				return func(_123_dt__update_hcf3_h0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC2_((_122_dt__update__tmp_h0).Dtor_cf2(), _123_dt__update_hcf3_h0, (_122_dt__update__tmp_h0).Dtor_cf4(), (_122_dt__update__tmp_h0).Dtor_cf5())
				}(_pat_let1_0)
			}(_dafny.IntOfInt64(-516))
		}(_pat_let0_0)
	}(_120_v46))
	_121_v47 = _nw13
	var _124_v48 _dafny.Set
	_ = _124_v48
	_124_v48 = _dafny.SetOf(_121_v47)
	var _hi0 _dafny.Int = (_124_v48).Cardinality()
	_ = _hi0
	for _125_i11 := _95_v35; _125_i11.Cmp(_hi0) < 0; _125_i11 = _125_i11.Plus(_dafny.One) {
		var _126_v49 _dafny.Array
		_ = _126_v49
		var _nwElement0_2 _dafny.Array = _119_v45
		_ = _nwElement0_2
		var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(5))
		_ = _nw14
		(_nw14).ArraySet1(_nwElement0_2, 0)
		(_nw14).ArraySet1(_119_v45, 1)
		(_nw14).ArraySet1(_119_v45, 2)
		(_nw14).ArraySet1((func() _dafny.Array {
			if _40_v0 {
				return _119_v45
			}
			return _119_v45
		})(), 3)
		(_nw14).ArraySet1(_119_v45, 4)
		_126_v49 = _nw14
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_126_v49), 0))
		_ = _index11
		(_126_v49).ArraySet1(_119_v45, (_index11).Int())
		var _127_v50 _dafny.Map
		_ = _127_v50
		_127_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, _dafny.IntOfInt64(260))
		var _128_v51 _dafny.Array
		_ = _128_v51
		var _nwElement0_3 _dafny.Sequence = _116_v42
		_ = _nwElement0_3
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(3))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_3, 0)
		(_nw15).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}((func(_129_i11 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_130_i12 _dafny.Int) _dafny.Int {
				return _129_i11
			}
		})(_125_i11))), 1)
		(_nw15).ArraySet1(_116_v42, 2)
		_128_v51 = _nw15
		var _131_v52 D2
		_ = _131_v52
		_131_v52 = Companion_D2_.Create_DC8_(_127_v50, _128_v51, _40_v0, _40_v0, _125_i11)
		var _132_v53 D6
		_ = _132_v53
		_132_v53 = Companion_D6_.Create_DC17_(_131_v52)
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_126_v49), 0))
		_ = _index12
		var _rhs2 _dafny.Array = _119_v45
		_ = _rhs2
		var _rhs3 D6 = _132_v53
		_ = _rhs3
		var _lhs3 _dafny.Array = _126_v49
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(987), _dafny.ArrayLen((_126_v49), 0))
		_ = _lhs4
		(_lhs3).ArraySet1(_rhs2, (_lhs4).Int())
		_132_v53 = _rhs3
		_111_v37 = _dafny.CodePoint('v')
		var _133_v54 _dafny.Array
		_ = _133_v54
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_4
		var _nw16 _dafny.Array
		_ = _nw16
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw16 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Sequence = (func(_134_v37 _dafny.CodePoint, _135_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_136_i13 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg23 _dafny.Int) interface{} {
							return coer23(arg23)
						}
					}((func(_137_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_138_i14 _dafny.Int) _dafny.CodePoint {
							return _137_v37
						}
					})(_134_v37))), _135_p1)
				}
			})(_111_v37, p1)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw16 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw16).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw16).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_133_v54 = _nw16
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_133_v54), 0))
		_ = _index13
		(_133_v54).ArraySet1(p1, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_133_v54), 0))
		_ = _index14
		(_133_v54).ArraySet1((_117_v43).F24(), (_index14).Int())
		var _rhs4 bool = !(_40_v0)
		_ = _rhs4
		var _rhs5 bool = (_112_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_116_v42).Cardinality()), _dafny.IntOfUint32((_112_v38).Cardinality()))).Uint32()).(bool)
		_ = _rhs5
		var _lhs5 *GlobalState = globalState
		_ = _lhs5
		var _lhs6 *GlobalState = globalState
		_ = _lhs6
		_lhs5.F2 = _rhs4
		_lhs6.F2 = _rhs5
	}
	var _139_v55 _dafny.Set
	_ = _139_v55
	_139_v55 = _dafny.SetOf(Companion_Default___.Fm0(_40_v0, _40_v0, globalState), _40_v0)
	var _140_v56 _dafny.Sequence
	_ = _140_v56
	_140_v56 = _dafny.SeqOf(_139_v55, _139_v55, _139_v55)
	if !(((_140_v56).Select((Companion_Default___.SafeIndex((_116_v42).Select((Companion_Default___.SafeIndex(_95_v35, _dafny.IntOfUint32((_116_v42).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_140_v56).Cardinality()))).Uint32()).(_dafny.Set)).IsSubsetOf(_139_v55)) || (false) {
		var _141_v57 _dafny.Map
		_ = _141_v57
		_141_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, p3)
		var _142_v58 *C2
		_ = _142_v58
		var _nw17 *C2 = New_C2_()
		_ = _nw17
		_nw17.Ctor__((_dafny.MultiSetOf(_40_v0)).Cardinality())
		_142_v58 = _nw17
		var _143_v59 _dafny.Map
		_ = _143_v59
		_143_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _142_v58)
		var _144_v60 _dafny.Sequence
		_ = _144_v60
		_144_v60 = _dafny.SeqOf(_143_v59, _143_v59)
		var _145_v61 T0
		_ = _145_v61
		var _nw18 *C5 = New_C5_()
		_ = _nw18
		_nw18.Ctor__(_40_v0, (_121_v47).Fm6(_95_v35, _40_v0, _111_v37, p1, globalState))
		_145_v61 = _nw18
		var _146_v62 _dafny.Map
		_ = _146_v62
		_146_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _40_v0)
		var _147_v63 _dafny.Map
		_ = _147_v63
		_147_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v61, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, (_146_v62).Cardinality()))
		var _148_v64 _dafny.Array
		_ = _148_v64
		var _nwElement0_4 _dafny.Map = _141_v57
		_ = _nwElement0_4
		var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(23))
		_ = _nw19
		(_nw19).ArraySet1(_nwElement0_4, 0)
		(_nw19).ArraySet1((_141_v57).Update(_40_v0, _dafny.IntOfInt64(194)), 1)
		(_nw19).ArraySet1(_141_v57, 2)
		(_nw19).ArraySet1((_141_v57).Merge(_141_v57), 3)
		(_nw19).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, p2), 4)
		(_nw19).ArraySet1(_141_v57, 5)
		(_nw19).ArraySet1((_141_v57).Merge(_141_v57), 6)
		(_nw19).ArraySet1(_141_v57, 7)
		(_nw19).ArraySet1(_141_v57, 8)
		(_nw19).ArraySet1(_141_v57, 9)
		(_nw19).ArraySet1(_141_v57, 10)
		(_nw19).ArraySet1(_141_v57, 11)
		(_nw19).ArraySet1((_141_v57).Merge(_141_v57), 12)
		(_nw19).ArraySet1(_141_v57, 13)
		(_nw19).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, _dafny.IntOfUint32((_116_v42).Cardinality())), 14)
		(_nw19).ArraySet1((_141_v57).Merge(_141_v57), 15)
		(_nw19).ArraySet1(_141_v57, 16)
		(_nw19).ArraySet1(_141_v57, 17)
		(_nw19).ArraySet1(_141_v57, 18)
		(_nw19).ArraySet1((_141_v57).Merge((_141_v57).Update(_40_v0, _dafny.IntOfUint32((_144_v60).Cardinality()))), 19)
		(_nw19).ArraySet1((func() _dafny.Map {
			if (_147_v63).Contains(_145_v61) {
				return (_147_v63).Get(_145_v61).(_dafny.Map)
			}
			return _141_v57
		})(), 20)
		(_nw19).ArraySet1(_141_v57, 21)
		(_nw19).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v61.F16(), p3), 22)
		_148_v64 = _nw19
		var _149_v65 _dafny.Map
		_ = _149_v65
		_149_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, (_141_v57).Update(_40_v0, (_142_v58).F25()))
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_148_v64), 0))
		_ = _index15
		(_148_v64).ArraySet1((func() _dafny.Map {
			if (_149_v65).Contains(_40_v0) {
				return (_149_v65).Get(_40_v0).(_dafny.Map)
			}
			return _141_v57
		})(), (_index15).Int())
		var _150_v66 *C0
		_ = _150_v66
		var _nw20 *C0 = New_C0_()
		_ = _nw20
		_nw20.Ctor__(p3, _119_v45)
		_150_v66 = _nw20
		var _151_v67 T1
		_ = _151_v67
		var _nw21 *C1 = New_C1_()
		_ = _nw21
		_nw21.Ctor__(((_142_v58).F25()).Minus(_dafny.IntOfInt64(903)), (_dafny.SetOf(_dafny.Companion_Sequence_.Update(_112_v38, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_112_v38).Cardinality()))).Uint32(), false))).Cardinality(), (_112_v38).Select((Companion_Default___.SafeIndex((Companion_D9_.Create_DC26_((_142_v58).F25(), _150_v66)).Dtor_cf45(), _dafny.IntOfUint32((_112_v38).Cardinality()))).Uint32()).(bool), _116_v42, _95_v35)
		_151_v67 = _nw21
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_148_v64), 0))
		_ = _index16
		var _rhs6 _dafny.Map = _141_v57
		_ = _rhs6
		var _rhs7 _dafny.Int = (_142_v58).F25()
		_ = _rhs7
		var _rhs8 T1 = _151_v67
		_ = _rhs8
		var _lhs7 _dafny.Array = _148_v64
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(396), _dafny.ArrayLen((_148_v64), 0))
		_ = _lhs8
		(_lhs7).ArraySet1(_rhs6, (_lhs8).Int())
		_95_v35 = _rhs7
		_151_v67 = _rhs8
		var _152_v68 _dafny.Array
		_ = _152_v68
		var _nwElement0_5 _dafny.Int = (_151_v67).Fm13(_151_v67.F16(), globalState)
		_ = _nwElement0_5
		var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(9))
		_ = _nw22
		(_nw22).ArraySet1(_nwElement0_5, 0)
		(_nw22).ArraySet1(p3, 1)
		(_nw22).ArraySet1((_151_v67).F28(), 2)
		(_nw22).ArraySet1((_142_v58).F25(), 3)
		(_nw22).ArraySet1(p0, 4)
		(_nw22).ArraySet1(p2, 5)
		(_nw22).ArraySet1(p2, 6)
		(_nw22).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 7)
		(_nw22).ArraySet1((_142_v58).F25(), 8)
		_152_v68 = _nw22
		var _153_v69 _dafny.Array
		_ = _153_v69
		var _nwElement0_6 _dafny.Array = (func() _dafny.Array {
			if _151_v67.F16() {
				return _152_v68
			}
			return _152_v68
		})()
		_ = _nwElement0_6
		var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw23
		(_nw23).ArraySet1(_nwElement0_6, 0)
		(_nw23).ArraySet1(_152_v68, 1)
		(_nw23).ArraySet1(_152_v68, 2)
		(_nw23).ArraySet1(_152_v68, 3)
		(_nw23).ArraySet1(_152_v68, 4)
		(_nw23).ArraySet1(_152_v68, 5)
		(_nw23).ArraySet1(_152_v68, 6)
		(_nw23).ArraySet1(_152_v68, 7)
		(_nw23).ArraySet1(_152_v68, 8)
		(_nw23).ArraySet1(_152_v68, 9)
		(_nw23).ArraySet1(_152_v68, 10)
		(_nw23).ArraySet1(_152_v68, 11)
		(_nw23).ArraySet1(_152_v68, 12)
		_153_v69 = _nw23
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_153_v69), 0))
		_ = _index17
		(_153_v69).ArraySet1(_152_v68, (_index17).Int())
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_153_v69), 0))
		_ = _index18
		(_153_v69).ArraySet1(_152_v68, (_index18).Int())
		if _145_v61.F16() {
			_95_v35 = _95_v35
			(globalState).F5 = _40_v0
			var _154_v70 _dafny.Array
			_ = _154_v70
			var _nwElement0_7 _dafny.Sequence = _dafny.SeqOf((_150_v66).F26())
			_ = _nwElement0_7
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(27))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_7, 0)
			(_nw24).ArraySet1((_151_v67).F17(), 1)
			(_nw24).ArraySet1((_145_v61).F17(), 2)
			(_nw24).ArraySet1((_151_v67).F17(), 3)
			(_nw24).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-381))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_155_v67 T1) func(_dafny.Int) _dafny.Int {
				return func(_156_i15 _dafny.Int) _dafny.Int {
					return (_155_v67).F28()
				}
			})(_151_v67))), 4)
			(_nw24).ArraySet1((_145_v61).F17(), 5)
			(_nw24).ArraySet1((_151_v67).F17(), 6)
			(_nw24).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_157_v61 T0) func(_dafny.Int) _dafny.Int {
				return func(_158_i16 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32(((_157_v61).F17()).Cardinality())
				}
			})(_145_v61))), 7)
			(_nw24).ArraySet1((_145_v61).F17(), 8)
			(_nw24).ArraySet1((_151_v67).F17(), 9)
			(_nw24).ArraySet1((_145_v61).F17(), 10)
			(_nw24).ArraySet1((_151_v67).F17(), 11)
			(_nw24).ArraySet1((_145_v61).F17(), 12)
			(_nw24).ArraySet1(_116_v42, 13)
			(_nw24).ArraySet1((_145_v61).F17(), 14)
			(_nw24).ArraySet1((_145_v61).F17(), 15)
			(_nw24).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(503))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_159_v67 T1) func(_dafny.Int) _dafny.Int {
				return func(_160_i17 _dafny.Int) _dafny.Int {
					return (_159_v67).F28()
				}
			})(_151_v67))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(503))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}((func(_161_v67 T1) func(_dafny.Int) _dafny.Int {
				return func(_162_i17 _dafny.Int) _dafny.Int {
					return (_161_v67).F28()
				}
			})(_151_v67)))).Cardinality()))).Uint32(), _dafny.IntOfInt64(-749)), 16)
			(_nw24).ArraySet1((_151_v67).F17(), 17)
			(_nw24).ArraySet1(_dafny.SeqOf(p2), 18)
			(_nw24).ArraySet1((_145_v61).F17(), 19)
			(_nw24).ArraySet1((_145_v61).F17(), 20)
			(_nw24).ArraySet1(_116_v42, 21)
			(_nw24).ArraySet1((_151_v67).F17(), 22)
			(_nw24).ArraySet1((_151_v67).F17(), 23)
			(_nw24).ArraySet1((_145_v61).F17(), 24)
			(_nw24).ArraySet1(_116_v42, 25)
			(_nw24).ArraySet1((_151_v67).F17(), 26)
			_154_v70 = _nw24
			var _163_v71 D2
			_ = _163_v71
			_163_v71 = Companion_D2_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_151_v67.F16(), _95_v35), _154_v70, _40_v0, _145_v61.F16(), _95_v35)
			var _164_v72 D6
			_ = _164_v72
			_164_v72 = Companion_D6_.Create_DC17_(_163_v71)
			var _pat_let_tv7 = _163_v71
			_ = _pat_let_tv7
			var _165_v73 _dafny.MultiSet
			_ = _165_v73
			_165_v73 = _dafny.MultiSetOf(func(_pat_let2_0 D6) D6 {
				return func(_166_dt__update__tmp_h1 D6) D6 {
					return func(_pat_let3_0 D2) D6 {
						return func(_167_dt__update_hcf31_h0 D2) D6 {
							return Companion_D6_.Create_DC17_(_167_dt__update_hcf31_h0)
						}(_pat_let3_0)
					}(_pat_let_tv7)
				}(_pat_let2_0)
			}(_164_v72), _164_v72, _164_v72)
			var _168_v74 _dafny.Sequence
			_ = _168_v74
			_168_v74 = _dafny.SeqOf(_165_v73)
			var _169_v75 _dafny.Sequence
			_ = _169_v75
			_169_v75 = _168_v74
			var _170_v76 _dafny.Set
			_ = _170_v76
			_170_v76 = _dafny.SetOf(_169_v75)
			var _171_v77 _dafny.MultiSet
			_ = _171_v77
			_171_v77 = _dafny.MultiSetOf((_170_v76).IsSubsetOf(_170_v76))
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_119_v45), 0))
			_ = _index19
			(_119_v45).ArraySet1(_151_v67.F16(), (_index19).Int())
			var _172_v78 _dafny.Array
			_ = _172_v78
			var _nw25 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
			_ = _nw25
			_172_v78 = _nw25
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_172_v78), 0))
			_ = _index20
			(_172_v78).ArraySet1((func() _dafny.Sequence {
				if _151_v67.F16() {
					return _112_v38
				}
				return _112_v38
			})(), (_index20).Int())
			var _173_v79 _dafny.Map
			_ = _173_v79
			_173_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _151_v67)
			var _174_v80 _dafny.Map
			_ = _174_v80
			_174_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_173_v79).Cardinality(), _150_v66)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_119_v45), 0))
			_ = _index21
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_172_v78), 0))
			_ = _index22
			var _rhs9 bool = _145_v61.F16()
			_ = _rhs9
			var _rhs10 _dafny.MultiSet = _171_v77
			_ = _rhs10
			var _rhs11 bool = true
			_ = _rhs11
			var _rhs12 _dafny.Sequence = (_114_v40).Select((Companion_Default___.SafeIndex(_95_v35, _dafny.IntOfUint32((_114_v40).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs12
			var _rhs13 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(globalState), (_174_v80).Cardinality()))
			_ = _rhs13
			var _lhs9 T0 = _145_v61
			_ = _lhs9
			var _lhs10 _dafny.Array = _119_v45
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_119_v45), 0))
			_ = _lhs11
			var _lhs12 _dafny.Array = _172_v78
			_ = _lhs12
			var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_172_v78), 0))
			_ = _lhs13
			_lhs9.F16_set_(_rhs9)
			_171_v77 = _rhs10
			(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
			(_lhs12).ArraySet1(_rhs12, (_lhs13).Int())
			_95_v35 = _rhs13
			(_150_v66).F27 = _150_v66.F27
			(globalState).F5 = (_40_v0) || ((_119_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_119_v45), 0))).Int()).(bool))
		} else {
			(_145_v61).F16_set_(!(_145_v61.F16()))
			var _175_v81 D12
			_ = _175_v81
			_175_v81 = Companion_D12_.Create_DC34_(_119_v45, p3, _152_v68, _111_v37)
			var _176_v82 _dafny.Sequence
			_ = _176_v82
			_176_v82 = _dafny.SeqOf(_150_v66.F27)
			var _177_v83 _dafny.Array
			_ = _177_v83
			var _nwElement0_8 _dafny.Array = _150_v66.F27
			_ = _nwElement0_8
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(10))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_8, 0)
			(_nw26).ArraySet1(_150_v66.F27, 1)
			(_nw26).ArraySet1((_175_v81).Dtor_cf59(), 2)
			(_nw26).ArraySet1(_150_v66.F27, 3)
			(_nw26).ArraySet1(_150_v66.F27, 4)
			(_nw26).ArraySet1(_150_v66.F27, 5)
			(_nw26).ArraySet1(_150_v66.F27, 6)
			(_nw26).ArraySet1(_150_v66.F27, 7)
			(_nw26).ArraySet1((_176_v82).Select((Companion_Default___.SafeIndex((_142_v58).F25(), _dafny.IntOfUint32((_176_v82).Cardinality()))).Uint32()).(_dafny.Array), 8)
			(_nw26).ArraySet1(_119_v45, 9)
			_177_v83 = _nw26
			var _178_v84 _dafny.Map
			_ = _178_v84
			_178_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v37, _177_v83)
			_178_v84 = (_178_v84).Update(_111_v37, _177_v83)
			var _179_v85 _dafny.MultiSet
			_ = _179_v85
			_179_v85 = _dafny.MultiSetOf(_145_v61.F16())
			(globalState).F2 = (_179_v85).IsDisjointFrom(_179_v85)
			var _180_v86 *C7
			_ = _180_v86
			var _nw27 *C7 = New_C7_()
			_ = _nw27
			_nw27.Ctor__(_152_v68, _151_v67.F16(), (_145_v61).F17())
			_180_v86 = _nw27
			_180_v86 = _180_v86
			var _181_v87 *C5
			_ = _181_v87
			var _nw28 *C5 = New_C5_()
			_ = _nw28
			_nw28.Ctor__(_151_v67.F16(), (_145_v61).F17())
			_181_v87 = _nw28
			var _182_v88 _dafny.Map
			_ = _182_v88
			_182_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_181_v87, _40_v0)
			var _183_v89 _dafny.Map
			_ = _183_v89
			_183_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v61.F16(), _145_v61.F16())
			var _184_v90 D1
			_ = _184_v90
			_184_v90 = Companion_D1_.Create_DC4_(_145_v61.F16(), (_182_v88).Cardinality(), (func() _dafny.Map {
				if _145_v61.F16() {
					return _183_v89
				}
				return _183_v89
			})())
			var _rhs14 bool = (func() bool {
				if _145_v61.F16() {
					return _dafny.Companion_Sequence_.Equal(_112_v38, _dafny.SeqOf(_151_v67.F16(), _145_v61.F16()))
				}
				return _40_v0
			})()
			_ = _rhs14
			var _rhs15 D1 = (func() D1 {
				if (((_121_v47).F21()).Dtor_cf2()).Cmp(p3) <= 0 {
					return _184_v90
				}
				return _184_v90
			})()
			_ = _rhs15
			var _lhs14 T1 = _151_v67
			_ = _lhs14
			_lhs14.F16_set_(_rhs14)
			_184_v90 = _rhs15
		}
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_119_v45), 0))
		_ = _index23
		(_119_v45).ArraySet1(_151_v67.F16(), (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_119_v45), 0))
		_ = _index24
		(_119_v45).ArraySet1(_40_v0, (_index24).Int())
		(globalState).F2 = !(_40_v0)
	} else {
		var _185_v91 _dafny.MultiSet
		_ = _185_v91
		_185_v91 = _dafny.MultiSetOf(_40_v0, _40_v0, _40_v0)
		var _186_v92 _dafny.Map
		_ = _186_v92
		_186_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_40_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_116_v42).Cardinality()), _40_v0)).Cardinality())
		var _187_v93 _dafny.MultiSet
		_ = _187_v93
		_187_v93 = _dafny.MultiSetOf(_185_v91, Companion_Default___.Fm36((_186_v92).Cardinality(), _40_v0, globalState), _dafny.MultiSetOf(_40_v0))
		(globalState).F2 = !((_187_v93).Intersection(_dafny.MultiSetOf(_dafny.MultiSetOf(_40_v0)))).Contains((_185_v91).Update(_40_v0, Companion_Default___.Abs(_95_v35)))
		_112_v38 = Companion_Default___.Fm20(_40_v0, _40_v0, !(_40_v0), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(577), p2), globalState)
		(globalState).F5 = (_40_v0) == (_40_v0)
		var _188_v94 _dafny.Array
		_ = _188_v94
		var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
		_ = _nw29
		_188_v94 = _nw29
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_188_v94), 0))
		_ = _index25
		(_188_v94).ArraySet1CodePoint(_111_v37, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(336), _dafny.ArrayLen((_188_v94), 0))
		_ = _index26
		(_188_v94).ArraySet1CodePoint(_111_v37, (_index26).Int())
		_40_v0 = (_112_v38).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if _40_v0 {
				return (func() _dafny.Int {
					if (_115_v41).Contains(_dafny.IntOfUint32((p1).Cardinality())) {
						return (_115_v41).Get(_dafny.IntOfUint32((p1).Cardinality())).(_dafny.Int)
					}
					return _dafny.IntOfInt64(487)
				})()
			}
			return p0
		})(), _dafny.IntOfUint32((_112_v38).Cardinality()))).Uint32()).(bool)
	}
	if (func() bool {
		if (!(_40_v0)) && (_40_v0) {
			return _dafny.Companion_Sequence_.Contains(_112_v38, _40_v0)
		}
		return !(!_dafny.Companion_Sequence_.Equal(p1, _dafny.UnicodeSeqOfUtf8Bytes("ykfd")))
	})() {
		(globalState).F5 = (p2).Cmp(p0) != 0
		_95_v35 = _95_v35
		var _189_v95 _dafny.Sequence
		_ = _189_v95
		_189_v95 = _dafny.UnicodeSeqOfUtf8Bytes("vwqu")
		_189_v95 = _189_v95
		var _190_v96 *C1
		_ = _190_v96
		var _nw30 *C1 = New_C1_()
		_ = _nw30
		_nw30.Ctor__(_dafny.IntOfInt64(527), (func() _dafny.Int {
			if _40_v0 {
				return p3
			}
			return p0
		})(), false, _116_v42, p0)
		_190_v96 = _nw30
		var _191_v97 _dafny.MultiSet
		_ = _191_v97
		_191_v97 = _dafny.MultiSetOf(_111_v37, _111_v37)
		var _rhs16 _dafny.Int = Companion_Default___.SafeModuloInt(((func() _dafny.MultiSet {
			if _40_v0 {
				return _191_v97
			}
			return (Companion_Default___.Fm40(p3, _40_v0, _139_v55, _dafny.IntOfInt64(582), globalState)).Update(_111_v37, Companion_Default___.Abs(p0))
		})()).Cardinality(), Companion_Default___.SafeModuloInt(p2, p2))
		_ = _rhs16
		var _rhs17 *C1 = _190_v96
		_ = _rhs17
		_95_v35 = _rhs16
		_190_v96 = _rhs17
		var _192_v98 _dafny.MultiSet
		_ = _192_v98
		_192_v98 = _dafny.MultiSetOf(_112_v38)
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_119_v45), 0))
		_ = _index27
		(_119_v45).ArraySet1((_dafny.MultiSetOf(_112_v38)).IsProperSubsetOf(_192_v98), (_index27).Int())
		var _193_v99 _dafny.Set
		_ = _193_v99
		_193_v99 = _dafny.SetOf((_117_v43).Fm11(_112_v38, _40_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-325)), _40_v0, globalState), (_117_v43).F24(), _dafny.UnicodeSeqOfUtf8Bytes("vgwwcrgf"))
		var _194_v100 _dafny.MultiSet
		_ = _194_v100
		_194_v100 = _dafny.MultiSetOf(p0, p3)
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_119_v45), 0))
		_ = _index28
		(_119_v45).ArraySet1((_194_v100).IsSubsetOf(_194_v100), (_index28).Int())
		var _195_v101 *C8
		_ = _195_v101
		var _nw31 *C8 = New_C8_()
		_ = _nw31
		_nw31.Ctor__(_189_v95, _139_v55, _40_v0, _dafny.Companion_Sequence_.Update(_116_v42, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_116_v42).Cardinality()))).Uint32(), (_190_v96).F30()))
		_195_v101 = _nw31
		var _196_v102 _dafny.Set
		_ = _196_v102
		_196_v102 = _dafny.SetOf(_195_v101, _195_v101)
		var _197_v103 _dafny.Map
		_ = _197_v103
		_197_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-662))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_198_v37 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_199_i18 _dafny.Int) _dafny.CodePoint {
				return _198_v37
			}
		})(_111_v37))), _dafny.IntOfUint32((_dafny.SeqOf(!(_40_v0), false, _40_v0)).Cardinality()))
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_119_v45), 0))
		_ = _index29
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_119_v45), 0))
		_ = _index30
		var _rhs18 bool = (_196_v102).IsDisjointFrom(_196_v102)
		_ = _rhs18
		var _rhs19 _dafny.Set = func() _dafny.Set {
			var _coll20 = _dafny.NewBuilder()
			_ = _coll20
			for _iter20 := _dafny.Iterate((_197_v103).Keys().Elements()); ; {
				_compr_20, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _200_v104 _dafny.Sequence
				_200_v104 = interface{}(_compr_20).(_dafny.Sequence)
				if (_197_v103).Contains(_200_v104) {
					_coll20.Add(_200_v104)
				}
			}
			return _coll20.ToSet()
		}()
		_ = _rhs19
		var _rhs20 bool = !(_40_v0)
		_ = _rhs20
		var _rhs21 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p2))
		_ = _rhs21
		var _lhs15 _dafny.Array = _119_v45
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(94), _dafny.ArrayLen((_119_v45), 0))
		_ = _lhs16
		var _lhs17 _dafny.Array = _119_v45
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(208), _dafny.ArrayLen((_119_v45), 0))
		_ = _lhs18
		(_lhs15).ArraySet1(_rhs18, (_lhs16).Int())
		_193_v99 = _rhs19
		(_lhs17).ArraySet1(_rhs20, (_lhs18).Int())
		_95_v35 = _rhs21
	} else {
		var _201_v105 T1
		_ = _201_v105
		var _nw32 *C1 = New_C1_()
		_ = _nw32
		_nw32.Ctor__(_95_v35, p3, _40_v0, _116_v42, p2)
		_201_v105 = _nw32
		var _202_v106 _dafny.Map
		_ = _202_v106
		_202_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_201_v105, _119_v45)
		var _203_v107 _dafny.Map
		_ = _203_v107
		_203_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
			if _40_v0 {
				return _119_v45
			}
			return (func() _dafny.Array {
				if (_202_v106).Contains(_201_v105) {
					return (_202_v106).Get(_201_v105).(_dafny.Array)
				}
				return _119_v45
			})()
		})(), _201_v105.F16())
		_203_v107 = (_203_v107).Merge(_203_v107)
		var _204_v108 _dafny.Array
		_ = _204_v108
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_5
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) D1 = (func(_205_v105 T1) func(_dafny.Int) D1 {
				return func(_206_i19 _dafny.Int) D1 {
					return Companion_D1_.Create_DC5_((_205_v105).F28(), _205_v105.F16())
				}
			})(_201_v105)
			_ = _init5
			var _element0_5 = _init5(_dafny.Zero)
			_ = _element0_5
			_nw33 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
			(_nw33).ArraySet1(_element0_5, 0)
			var _nativeLen0_5 = (_len0_5).Int()
			_ = _nativeLen0_5
			for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
				(_nw33).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
			}
		}
		_204_v108 = _nw33
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_204_v108), 0))
		_ = _index31
		(_204_v108).ArraySet1(Companion_Default___.Fm34(_dafny.IntOfInt64(13), _dafny.IntOfInt64(795), _139_v55, _dafny.IntOfUint32((_112_v38).Cardinality()), globalState), (_index31).Int())
		var _207_v109 D1
		_ = _207_v109
		_207_v109 = Companion_D1_.Create_DC5_((_201_v105).F28(), true)
		var _pat_let_tv8 = _40_v0
		_ = _pat_let_tv8
		var _pat_let_tv9 = _40_v0
		_ = _pat_let_tv9
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_204_v108), 0))
		_ = _index32
		var _rhs22 D1 = func(_pat_let4_0 D1) D1 {
			return func(_208_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let5_0 bool) D1 {
					return func(_209_dt__update_hcf11_h0 bool) D1 {
						return Companion_D1_.Create_DC5_((_208_dt__update__tmp_h2).Dtor_cf10(), _209_dt__update_hcf11_h0)
					}(_pat_let5_0)
				}((func() bool {
					if _pat_let_tv8 {
						return true
					}
					return _pat_let_tv9
				})())
			}(_pat_let4_0)
		}(_207_v109)
		_ = _rhs22
		var _rhs23 bool = (p2).Cmp(p0) > 0
		_ = _rhs23
		var _lhs19 _dafny.Array = _204_v108
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_204_v108), 0))
		_ = _lhs20
		(_lhs19).ArraySet1(_rhs22, (_lhs20).Int())
		_40_v0 = _rhs23
		var _210_v110 _dafny.Array
		_ = _210_v110
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
		_ = _nw34
		_210_v110 = _nw34
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_210_v110), 0))
		_ = _index33
		(_210_v110).ArraySet1CodePoint(_111_v37, (_index33).Int())
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_210_v110), 0))
		_ = _index34
		(_210_v110).ArraySet1CodePoint(_111_v37, (_index34).Int())
		var _211_v111 D2
		_ = _211_v111
		_211_v111 = Companion_D2_.Create_DC7_(_139_v55)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_119_v45), 0))
		_ = _index35
		(_119_v45).ArraySet1(_201_v105.F16(), (_index35).Int())
		var _212_v113 _dafny.MultiSet
		_ = _212_v113
		_212_v113 = _dafny.MultiSetOf((_210_v110).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_210_v110), 0))).Int()), (_210_v110).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_210_v110), 0))).Int()), _111_v37)
		var _213_v114 D9
		_ = _213_v114
		_213_v114 = Companion_D9_.Create_DC25_(func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((_212_v113).Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _214_v112 _dafny.CodePoint
				_214_v112 = interface{}(_compr_21).(_dafny.CodePoint)
				if (_212_v113).Contains(_214_v112) {
					_coll21.Add(_214_v112, true)
				}
			}
			return _coll21.ToMap()
		}())
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_119_v45), 0))
		_ = _index36
		var _rhs24 D2 = _211_v111
		_ = _rhs24
		var _rhs25 _dafny.CodePoint = ((_117_v43).F24()).Select((Companion_Default___.SafeIndex(_95_v35, _dafny.IntOfUint32(((_117_v43).F24()).Cardinality()))).Uint32()).(_dafny.CodePoint)
		_ = _rhs25
		var _rhs26 bool = (func() bool {
			if _201_v105.F16() {
				return true
			}
			return _201_v105.F16()
		})()
		_ = _rhs26
		var _rhs27 bool = !(_40_v0)
		_ = _rhs27
		var _rhs28 _dafny.Map = Companion_Default___.Fm38(_213_v114, (_201_v105).F28(), globalState)
		_ = _rhs28
		var _lhs21 *GlobalState = globalState
		_ = _lhs21
		var _lhs22 _dafny.Array = _119_v45
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_119_v45), 0))
		_ = _lhs23
		var _lhs24 T1 = _201_v105
		_ = _lhs24
		var _lhs25 *GlobalState = globalState
		_ = _lhs25
		_211_v111 = _rhs24
		_lhs21.F15 = _rhs25
		(_lhs22).ArraySet1(_rhs26, (_lhs23).Int())
		_lhs24.F16_set_(_rhs27)
		_lhs25.F8 = _rhs28
		(_201_v105).F16_set_(!(_40_v0) || (!(true)))
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _215_v0 _dafny.Sequence
	_ = _215_v0
	_215_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dg")
	var _216_v1 _dafny.Int
	_ = _216_v1
	_216_v1 = _dafny.IntOfInt64(879)
	var _217_v2 _dafny.CodePoint
	_ = _217_v2
	_217_v2 = _dafny.CodePoint('t')
	var _218_v3 _dafny.Array
	_ = _218_v3
	var _nwElement0_9 _dafny.Sequence = _215_v0
	_ = _nwElement0_9
	var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(12))
	_ = _nw35
	(_nw35).ArraySet1(_nwElement0_9, 0)
	(_nw35).ArraySet1(_215_v0, 1)
	(_nw35).ArraySet1(_215_v0, 2)
	(_nw35).ArraySet1(_215_v0, 3)
	(_nw35).ArraySet1(_215_v0, 4)
	(_nw35).ArraySet1(_215_v0, 5)
	(_nw35).ArraySet1(_215_v0, 6)
	(_nw35).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("k"), 7)
	(_nw35).ArraySet1(_dafny.Companion_Sequence_.Update(_215_v0, (Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_215_v0).Cardinality()))).Uint32(), _217_v2), 8)
	(_nw35).ArraySet1(_215_v0, 9)
	(_nw35).ArraySet1(_dafny.Companion_Sequence_.Update(_215_v0, (Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_215_v0).Cardinality()))).Uint32(), _217_v2), 10)
	(_nw35).ArraySet1(_215_v0, 11)
	_218_v3 = _nw35
	var _219_v4 _dafny.Array
	_ = _219_v4
	var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
	_ = _nw36
	_219_v4 = _nw36
	var _220_v5 _dafny.Array
	_ = _220_v5
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
	_ = _len0_6
	var _nw37 _dafny.Array
	_ = _nw37
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw37 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = func(_221_i0 _dafny.Int) bool {
			return true
		}
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw37 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw37).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw37).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_220_v5 = _nw37
	var _222_v6 _dafny.Sequence
	_ = _222_v6
	_222_v6 = _dafny.SeqOf(_220_v5)
	var _223_v7 bool
	_ = _223_v7
	_223_v7 = true
	var _224_v8 _dafny.Map
	_ = _224_v8
	_224_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v7, _223_v7)
	var _225_v9 _dafny.Map
	_ = _225_v9
	_225_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v1, (func() bool {
		if (_224_v8).Contains(false) {
			return (_224_v8).Get(false).(bool)
		}
		return _223_v7
	})())
	var _226_v10 _dafny.Sequence
	_ = _226_v10
	_226_v10 = _dafny.SeqOf(_223_v7)
	var _227_v11 _dafny.Map
	_ = _227_v11
	_227_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_v10, _223_v7)
	var _228_v12 _dafny.Sequence
	_ = _228_v12
	_228_v12 = _dafny.SeqOf(_dafny.MultiSetOf(_216_v1, _dafny.IntOfUint32((_215_v0).Cardinality()), _216_v1, _216_v1, _216_v1))
	var _229_v13 _dafny.Sequence
	_ = _229_v13
	_229_v13 = _dafny.SeqOf(_216_v1, _216_v1)
	var _230_globalState *GlobalState
	_ = _230_globalState
	var _nw38 *GlobalState = New_GlobalState_()
	_ = _nw38
	_nw38.Ctor__(_218_v3, _219_v4, false, _dafny.IntOfInt64(526), (_222_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.IntOfUint32((_222_v6).Cardinality()))).Uint32()).(_dafny.Array), true, true, _dafny.IntOfInt64(72), _225_v9, (_227_v11).Merge(_227_v11), _228_v12, false, _229_v13, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}((func(_231_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_232_i1 _dafny.Int) _dafny.CodePoint {
			return _231_v2
		}
	})(_217_v2))), true, _dafny.CodePoint('w'))
	_230_globalState = _nw38
	var _233_v14 _dafny.Map
	_ = _233_v14
	_233_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(Companion_Default___.Fm0(_223_v7, _223_v7, _230_globalState)), _220_v5)
	_216_v1 = (((_233_v14).Merge(_233_v14)).Update(_223_v7, _220_v5)).Cardinality()
	var _234_v15 _dafny.Array
	_ = _234_v15
	var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
	_ = _len0_7
	var _nw39 _dafny.Array
	_ = _nw39
	if _len0_7.Cmp(_dafny.Zero) == 0 {
		_nw39 = _dafny.NewArray(_len0_7)
	} else {
		var _init7 func(_dafny.Int) _dafny.Map = (func(_235_v1 _dafny.Int) func(_dafny.Int) _dafny.Map {
			return func(_236_i2 _dafny.Int) _dafny.Map {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v1, _dafny.IntOfInt64(-637))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_235_v1, _235_v1))
			}
		})(_216_v1)
		_ = _init7
		var _element0_7 = _init7(_dafny.Zero)
		_ = _element0_7
		_nw39 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
		(_nw39).ArraySet1(_element0_7, 0)
		var _nativeLen0_7 = (_len0_7).Int()
		_ = _nativeLen0_7
		for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
			(_nw39).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
		}
	}
	_234_v15 = _nw39
	var _237_v16 _dafny.Map
	_ = _237_v16
	_237_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v1, _216_v1)
	var _238_v17 _dafny.Map
	_ = _238_v17
	_238_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfUint32((_215_v0).Cardinality()))
	var _239_v18 _dafny.Map
	_ = _239_v18
	_239_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
		if (_237_v16).Contains(_216_v1) {
			return (_237_v16).Get(_216_v1).(_dafny.Int)
		}
		return _dafny.IntOfUint32((Companion_Default___.Fm1(_230_globalState)).Cardinality())
	})(), (_238_v17).Cardinality())
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_234_v15), 0))
	_ = _index37
	(_234_v15).ArraySet1(_239_v18, (_index37).Int())
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_234_v15), 0))
	_ = _index38
	(_234_v15).ArraySet1(_239_v18, (_index38).Int())
	_216_v1 = (_216_v1).Times(_216_v1)
	var _240_v19 _dafny.Array
	_ = _240_v19
	var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
	_ = _nw40
	_240_v19 = _nw40
	var _241_v20 _dafny.Map
	_ = _241_v20
	_241_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v19, _239_v18)
	var _242_v21 _dafny.Array
	_ = _242_v21
	var _nwElement0_10 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_226_v10, _226_v10)).Cardinality())
	_ = _nwElement0_10
	var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
	_ = _nw41
	(_nw41).ArraySet1(_nwElement0_10, 0)
	(_nw41).ArraySet1((_216_v1).Minus(_dafny.IntOfInt64(948)), 1)
	(_nw41).ArraySet1((_dafny.Zero).Minus(_216_v1), 2)
	(_nw41).ArraySet1(_216_v1, 3)
	(_nw41).ArraySet1(_216_v1, 4)
	(_nw41).ArraySet1((_dafny.Zero).Minus(((_241_v20).Merge(_241_v20)).Cardinality()), 5)
	(_nw41).ArraySet1((_dafny.SetOf(_223_v7, _223_v7)).Cardinality(), 6)
	(_nw41).ArraySet1(_216_v1, 7)
	(_nw41).ArraySet1(_216_v1, 8)
	(_nw41).ArraySet1(_216_v1, 9)
	(_nw41).ArraySet1(_216_v1, 10)
	(_nw41).ArraySet1(_216_v1, 11)
	(_nw41).ArraySet1(_216_v1, 12)
	_242_v21 = _nw41
	_240_v19 = _240_v19
	var _243_i3 _dafny.Int
	_ = _243_i3
	_243_i3 = _dafny.Zero
	{
		for (_223_v7) && (_223_v7) {
			{
				if (_243_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_243_i3 = (_243_i3).Plus(_dafny.One)
				_216_v1 = _216_v1
				Companion_Default___.M0(_216_v1, _dafny.Companion_Sequence_.Update(_215_v0, (Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_215_v0).Cardinality()))).Uint32(), _217_v2), _dafny.IntOfUint32((_dafny.SeqOf(_216_v1, (_dafny.Zero).Minus((_dafny.Zero).Minus(_216_v1)))).Cardinality()), (_216_v1).Minus(_dafny.IntOfInt64(132)), _230_globalState)
				var _244_v22 _dafny.Map
				_ = _244_v22
				_244_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v6, _223_v7)
				_244_v22 = (_244_v22).Update(_dafny.Companion_Sequence_.Concatenate(_222_v6, _222_v6), (_226_v10).Select((Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_226_v10).Cardinality()))).Uint32()).(bool))
				Companion_Default___.M0(_216_v1, _215_v0, (_216_v1).Plus(_216_v1), _216_v1, _230_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	Companion_Default___.M0(_216_v1, _dafny.UnicodeSeqOfUtf8Bytes("casqliqs"), Companion_Default___.SafeDivisionInt(_216_v1, _dafny.IntOfUint32((_215_v0).Cardinality())), (func() _dafny.Map {
		var _coll22 = _dafny.NewMapBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(327), _dafny.IntOfInt64(229))); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _245_v23 _dafny.Int
			_245_v23 = interface{}(_compr_22).(_dafny.Int)
			if ((_dafny.IntOfInt64(327)).Cmp(_245_v23) <= 0) && ((_245_v23).Cmp(_dafny.IntOfInt64(229)) < 0) {
				_coll22.Add(Companion_Default___.SafeModuloInt(_245_v23, _216_v1), _223_v7)
			}
		}
		return _coll22.ToMap()
	}()).Cardinality(), _230_globalState)
	_216_v1 = _216_v1
	var _246_i4 _dafny.Int
	_ = _246_i4
	_246_i4 = _dafny.Zero
	{
		for _223_v7 {
			{
				if (_246_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_246_i4 = (_246_i4).Plus(_dafny.One)
				var _247_v24 _dafny.Map
				_ = _247_v24
				_247_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wqs"), _216_v1)
				var _248_v26 _dafny.Map
				_ = _248_v26
				_248_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_215_v0, _dafny.SetOf(_223_v7))
				var _249_v28 _dafny.Sequence
				_ = _249_v28
				_249_v28 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-396))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_250_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_251_i5 _dafny.Int) _dafny.CodePoint {
						return _250_v2
					}
				})(_217_v2))))
				var _252_v29 _dafny.Sequence
				_ = _252_v29
				_252_v29 = _dafny.SeqOf(func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter23 := _dafny.Iterate((_248_v26).Keys().Elements()); ; {
						_compr_23, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _253_v25 _dafny.Sequence
						_253_v25 = interface{}(_compr_23).(_dafny.Sequence)
						if (_248_v26).Contains(_253_v25) {
							_coll23.Add(_253_v25, _216_v1)
						}
					}
					return _coll23.ToMap()
				}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rsb"), _216_v1), (func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter24 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_249_v28, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_216_v1), _dafny.IntOfUint32((_249_v28).Cardinality()))).Uint32(), _215_v0)).Elements()); ; {
						_compr_24, _ok24 := _iter24()
						if !_ok24 {
							break
						}
						var _254_v27 _dafny.Sequence
						_254_v27 = interface{}(_compr_24).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_249_v28, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_216_v1), _dafny.IntOfUint32((_249_v28).Cardinality()))).Uint32(), _215_v0), _254_v27) {
							_coll24.Add(_254_v27, _216_v1)
						}
					}
					return _coll24.ToMap()
				}()).Update(_215_v0, _dafny.IntOfUint32((_215_v0).Cardinality())))
				_247_v24 = ((_252_v29).Select((Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_252_v29).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_247_v24)
				var _nw42 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw42
				_220_v5 = _nw42
				_216_v1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_216_v1))
				(_230_globalState).F2 = ((_239_v18).Cardinality()).Cmp(_216_v1) <= 0
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))
	_ = _index39
	(_218_v3).ArraySet1(_215_v0, (_index39).Int())
	var _255_v30 _dafny.Map
	_ = _255_v30
	_255_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v7, _217_v2)
	var _256_v31 _dafny.Map
	_ = _256_v31
	_256_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_223_v7), _255_v30)
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
	_ = _index40
	(_240_v19).ArraySet1((func() _dafny.Int {
		if (func() bool {
			if (_224_v8).Contains(_223_v7) {
				return (_224_v8).Get(_223_v7).(bool)
			}
			return _223_v7
		})() {
			return (_256_v31).Cardinality()
		}
		return _216_v1
	})(), (_index40).Int())
	var _257_v32 _dafny.MultiSet
	_ = _257_v32
	_257_v32 = _dafny.MultiSetOf(_217_v2, _217_v2, Companion_Default___.Fm2(_230_globalState))
	var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))
	_ = _index41
	var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
	_ = _index42
	var _rhs29 _dafny.Sequence = (Companion_D0_.Create_DC0_(Companion_Default___.Fm1(_230_globalState))).Dtor_cf0()
	_ = _rhs29
	var _rhs30 _dafny.Int = (_257_v32).Cardinality()
	_ = _rhs30
	var _rhs31 _dafny.Array = _220_v5
	_ = _rhs31
	var _lhs26 _dafny.Array = _218_v3
	_ = _lhs26
	var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))
	_ = _lhs27
	var _lhs28 _dafny.Array = _240_v19
	_ = _lhs28
	var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
	_ = _lhs29
	(_lhs26).ArraySet1(_rhs29, (_lhs27).Int())
	(_lhs28).ArraySet1(_rhs30, (_lhs29).Int())
	_220_v5 = _rhs31
	for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_220_v5), 0))); ; {
		_guard_loop_0, _ok25 := _iter25()
		if !_ok25 {
			break
		}
		var _258_i6 _dafny.Int
		_258_i6 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_258_i6).Sign() != -1) && ((_258_i6).Cmp(_dafny.ArrayLen((_220_v5), 0)) < 0)) {
			(_220_v5).ArraySet1((_dafny.MultiSetOf(_223_v7)).IsDisjointFrom((_dafny.MultiSetOf(_223_v7)).Intersection(_dafny.MultiSetOf(_223_v7))), (_258_i6).Int())
		}
	}
	for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_220_v5), 0))); ; {
		_guard_loop_1, _ok26 := _iter26()
		if !_ok26 {
			break
		}
		var _259_i7 _dafny.Int
		_259_i7 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_259_i7).Sign() != -1) && ((_259_i7).Cmp(_dafny.ArrayLen((_220_v5), 0)) < 0)) {
			(_220_v5).ArraySet1(!(_223_v7), (_259_i7).Int())
		}
	}
	if _223_v7 {
		var _260_v34 _dafny.Map
		_ = _260_v34
		_260_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v7, func() _dafny.Map {
			var _coll25 = _dafny.NewMapBuilder()
			_ = _coll25
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(333), _dafny.IntOfInt64(468))); ; {
				_compr_25, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _261_v33 _dafny.Int
				_261_v33 = interface{}(_compr_25).(_dafny.Int)
				if ((_dafny.IntOfInt64(333)).Cmp(_261_v33) <= 0) && ((_261_v33).Cmp(_dafny.IntOfInt64(468)) < 0) {
					_coll25.Add(Companion_Default___.SafeDivisionInt(_261_v33, _216_v1), _223_v7)
				}
			}
			return _coll25.ToMap()
		}())
		_216_v1 = (func() _dafny.Int {
			if _223_v7 {
				return ((_260_v34).Merge(_260_v34)).Cardinality()
			}
			return Companion_Default___.Fm3(_230_globalState)
		})()
		var _262_v35 D0
		_ = _262_v35
		_262_v35 = Companion_D0_.Create_DC1_(Companion_Default___.Fm1(_230_globalState))
		var _source2 D0 = _262_v35
		_ = _source2
		if _source2.Is_DC1() {
			var _263___mcc_h0 _dafny.Sequence = _source2.Get_().(D0_DC1).Cf1
			_ = _263___mcc_h0
			var _264_cf1 _dafny.Sequence = _263___mcc_h0
			_ = _264_cf1
			var _265_v36 _dafny.MultiSet
			_ = _265_v36
			_265_v36 = _dafny.MultiSetOf(_223_v7, _223_v7, _223_v7)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_219_v4), 0))
			_ = _index43
			(_219_v4).ArraySet1(_265_v36, (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_220_v5), 0))
			_ = _index44
			(_220_v5).ArraySet1(_223_v7, (_index44).Int())
			var _266_v37 _dafny.Sequence
			_ = _266_v37
			_266_v37 = _dafny.SeqOf(_265_v36)
			var _267_v38 D1
			_ = _267_v38
			_267_v38 = Companion_D1_.Create_DC3_((_266_v37).Select((Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_266_v37).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _268_v39 D1
			_ = _268_v39
			_268_v39 = Companion_D1_.Create_DC4_(_223_v7, _216_v1, _224_v8)
			var _pat_let_tv10 = _265_v36
			_ = _pat_let_tv10
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_219_v4), 0))
			_ = _index45
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_220_v5), 0))
			_ = _index46
			var _rhs32 _dafny.MultiSet = (func(_pat_let6_0 D1) D1 {
				return func(_269_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let7_0 _dafny.MultiSet) D1 {
						return func(_270_dt__update_hcf6_h0 _dafny.MultiSet) D1 {
							return Companion_D1_.Create_DC3_(_270_dt__update_hcf6_h0)
						}(_pat_let7_0)
					}(_pat_let_tv10)
				}(_pat_let6_0)
			}(_267_v38)).Dtor_cf6()
			_ = _rhs32
			var _rhs33 bool = (_268_v39).Dtor_cf7()
			_ = _rhs33
			var _rhs34 bool = !(_223_v7)
			_ = _rhs34
			var _rhs35 bool = !(!(!(_223_v7)))
			_ = _rhs35
			var _lhs30 _dafny.Array = _219_v4
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen((_219_v4), 0))
			_ = _lhs31
			var _lhs32 *GlobalState = _230_globalState
			_ = _lhs32
			var _lhs33 _dafny.Array = _220_v5
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_220_v5), 0))
			_ = _lhs34
			var _lhs35 *GlobalState = _230_globalState
			_ = _lhs35
			(_lhs30).ArraySet1(_rhs32, (_lhs31).Int())
			_lhs32.F2 = _rhs33
			(_lhs33).ArraySet1(_rhs34, (_lhs34).Int())
			_lhs35.F5 = _rhs35
			_216_v1 = (_dafny.IntOfUint32((_215_v0).Cardinality())).Minus(_216_v1)
			var _271_v41 _dafny.Sequence
			_ = _271_v41
			_271_v41 = _dafny.SeqOf(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-26), _dafny.IntOfInt64(365))); ; {
					_compr_26, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _272_v40 _dafny.Int
					_272_v40 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(-26)).Cmp(_272_v40) <= 0) && ((_272_v40).Cmp(_dafny.IntOfInt64(365)) < 0) {
						_coll26.Add((_272_v40).Times(_216_v1), false)
					}
				}
				return _coll26.ToMap()
			}(), (_225_v9).Merge(_225_v9))
			_271_v41 = (func() _dafny.Sequence {
				if (_220_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_220_v5), 0))).Int()).(bool) {
					return _dafny.Companion_Sequence_.Update(_271_v41, (Companion_Default___.SafeIndex(_216_v1, _dafny.IntOfUint32((_271_v41).Cardinality()))).Uint32(), _225_v9)
				}
				return _dafny.Companion_Sequence_.Concatenate(_271_v41, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_216_v1, (_220_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_220_v5), 0))).Int()).(bool)), _225_v9, _225_v9))
			})()
			_225_v9 = (_225_v9).Update((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), !_dafny.Companion_Sequence_.Equal(_264_cf1, _215_v0))
		} else if _source2.Is_DC2() {
			var _273___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC2).Cf2
			_ = _273___mcc_h1
			var _274___mcc_h2 _dafny.Int = _source2.Get_().(D0_DC2).Cf3
			_ = _274___mcc_h2
			var _275___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC2).Cf4
			_ = _275___mcc_h3
			var _276___mcc_h4 _dafny.Array = _source2.Get_().(D0_DC2).Cf5
			_ = _276___mcc_h4
			var _277_cf5 _dafny.Array = _276___mcc_h4
			_ = _277_cf5
			var _278_cf4 _dafny.Int = _275___mcc_h3
			_ = _278_cf4
			var _279_cf3 _dafny.Int = _274___mcc_h2
			_ = _279_cf3
			var _280_cf2 _dafny.Int = _273___mcc_h1
			_ = _280_cf2
			_277_cf5 = _220_v5
			Companion_Default___.M0(_278_cf4, _dafny.Companion_Sequence_.Concatenate(_215_v0, (_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence)), _280_cf2, (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _230_globalState)
			var _281_v42 _dafny.MultiSet
			_ = _281_v42
			_281_v42 = _dafny.MultiSetOf(_223_v7)
			var _282_v43 _dafny.Sequence
			_ = _282_v43
			_282_v43 = _dafny.SeqOf(Companion_D1_.Create_DC3_(_281_v42))
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_220_v5), 0))
			_ = _index47
			(_220_v5).ArraySet1(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm4(_223_v7, _dafny.IntOfInt64(-367), _230_globalState), _282_v43), (_index47).Int())
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_220_v5), 0))
			_ = _index49
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index50
			var _rhs36 _dafny.Int = _278_cf4
			_ = _rhs36
			var _rhs37 bool = _223_v7
			_ = _rhs37
			var _rhs38 bool = _223_v7
			_ = _rhs38
			var _rhs39 _dafny.Int = _216_v1
			_ = _rhs39
			var _rhs40 _dafny.Int = _279_cf3
			_ = _rhs40
			var _lhs36 _dafny.Array = _240_v19
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _220_v5
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_220_v5), 0))
			_ = _lhs39
			var _lhs40 *GlobalState = _230_globalState
			_ = _lhs40
			var _lhs41 _dafny.Array = _240_v19
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs42
			(_lhs36).ArraySet1(_rhs36, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs37, (_lhs39).Int())
			_lhs40.F5 = _rhs38
			_279_cf3 = _rhs39
			(_lhs41).ArraySet1(_rhs40, (_lhs42).Int())
			(_230_globalState).F2 = (_220_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_220_v5), 0))).Int()).(bool)
		} else {
			var _283___mcc_h5 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
			_ = _283___mcc_h5
			var _284_cf0 _dafny.Sequence = _283___mcc_h5
			_ = _284_cf0
			var _285_v44 _dafny.MultiSet
			_ = _285_v44
			_285_v44 = _dafny.MultiSetOf(_223_v7, false)
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_219_v4), 0))
			_ = _index51
			(_219_v4).ArraySet1((_285_v44).Update(_223_v7, Companion_Default___.Abs(_216_v1)), (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_219_v4), 0))
			_ = _index52
			var _rhs41 _dafny.MultiSet = _285_v44
			_ = _rhs41
			var _rhs42 bool = _223_v7
			_ = _rhs42
			var _lhs43 _dafny.Array = _219_v4
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_219_v4), 0))
			_ = _lhs44
			var _lhs45 *GlobalState = _230_globalState
			_ = _lhs45
			(_lhs43).ArraySet1(_rhs41, (_lhs44).Int())
			_lhs45.F5 = _rhs42
			Companion_Default___.M0((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("wxynt"), (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(255)), _230_globalState)
			var _286_v45 _dafny.MultiSet
			_ = _286_v45
			_286_v45 = _dafny.MultiSetOf(((_219_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_219_v4), 0))).Int()).(_dafny.MultiSet)).Cardinality())
			Companion_Default___.M0(_216_v1, (_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence), (_286_v45).Cardinality(), (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _230_globalState)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_220_v5), 0))
			_ = _index53
			(_220_v5).ArraySet1(_223_v7, (_index53).Int())
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_220_v5), 0))
			_ = _index54
			(_220_v5).ArraySet1(_dafny.Companion_Sequence_.Equal(_284_cf0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-422))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}((func(_287_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_288_i8 _dafny.Int) _dafny.CodePoint {
					return _287_v2
				}
			})(_217_v2)))), (_index54).Int())
		}
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _index55
		(_240_v19).ArraySet1(_216_v1, (_index55).Int())
		var _289_v46 _dafny.Array
		_ = _289_v46
		var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw43
		_289_v46 = _nw43
		var _290_v47 _dafny.Sequence
		_ = _290_v47
		_290_v47 = _dafny.SeqOf(_dafny.SeqOf(_223_v7, _223_v7, _223_v7, _223_v7, _223_v7), _226_v10, _226_v10, _dafny.Companion_Sequence_.Update(_226_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_223_v7)).Cardinality()), _dafny.IntOfUint32((_226_v10).Cardinality()))).Uint32(), _223_v7), _dafny.SeqOf(false))
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_289_v46), 0))
		_ = _index56
		(_289_v46).ArraySet1(_290_v47, (_index56).Int())
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_289_v46), 0))
		_ = _index57
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _index58
		var _rhs43 _dafny.Sequence = _290_v47
		_ = _rhs43
		var _rhs44 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_229_v13, _229_v13)).Cardinality())
		_ = _rhs44
		var _rhs45 bool = _223_v7
		_ = _rhs45
		var _lhs46 _dafny.Array = _289_v46
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_289_v46), 0))
		_ = _lhs47
		var _lhs48 _dafny.Array = _240_v19
		_ = _lhs48
		var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _lhs49
		var _lhs50 *GlobalState = _230_globalState
		_ = _lhs50
		(_lhs46).ArraySet1(_rhs43, (_lhs47).Int())
		(_lhs48).ArraySet1(_rhs44, (_lhs49).Int())
		_lhs50.F5 = _rhs45
		var _291_v48 _dafny.Array
		_ = _291_v48
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw44
		_291_v48 = _nw44
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_291_v48), 0))
		_ = _index59
		(_291_v48).ArraySet1(_220_v5, (_index59).Int())
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_291_v48), 0))
		_ = _index60
		(_291_v48).ArraySet1(_220_v5, (_index60).Int())
	} else {
		var _292_v49 *C7
		_ = _292_v49
		var _nw45 *C7 = New_C7_()
		_ = _nw45
		_nw45.Ctor__(_242_v21, (!(_223_v7)) == (_223_v7), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _216_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_215_v0).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _216_v1)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xse")).Cardinality())))
		_292_v49 = _nw45
		var _293_v50 *C5
		_ = _293_v50
		var _nw46 *C5 = New_C5_()
		_ = _nw46
		_nw46.Ctor__(_223_v7, _229_v13)
		_293_v50 = _nw46
		_216_v1 = _dafny.IntOfInt64(-693)
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _index61
		var _rhs46 bool = !(_223_v7)
		_ = _rhs46
		var _rhs47 _dafny.Int = (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)
		_ = _rhs47
		var _lhs51 *GlobalState = _230_globalState
		_ = _lhs51
		var _lhs52 _dafny.Array = _240_v19
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _lhs53
		_lhs51.F5 = _rhs46
		(_lhs52).ArraySet1(_rhs47, (_lhs53).Int())
		var _294_v51 _dafny.Array
		_ = _294_v51
		var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
		_ = _nw47
		_294_v51 = _nw47
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _index62
		var _rhs48 _dafny.Array = (func() _dafny.Array {
			if Companion_Default___.Fm0(_223_v7, _223_v7, _230_globalState) {
				return _294_v51
			}
			return _294_v51
		})()
		_ = _rhs48
		var _rhs49 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_223_v7, _223_v7, _223_v7, _223_v7, _223_v7), _dafny.SeqOf(_223_v7)), _226_v10)
		_ = _rhs49
		var _rhs50 _dafny.Int = Companion_Default___.Fm3(_230_globalState)
		_ = _rhs50
		var _lhs54 _dafny.Array = _240_v19
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
		_ = _lhs55
		_294_v51 = _rhs48
		_223_v7 = _rhs49
		(_lhs54).ArraySet1(_rhs50, (_lhs55).Int())
	}
	var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
	_ = _index63
	(_240_v19).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_226_v10).Cardinality())), (_index63).Int())
	var _295_v52 _dafny.MultiSet
	_ = _295_v52
	_295_v52 = _dafny.MultiSetOf(_223_v7)
	var _296_v53 *C7
	_ = _296_v53
	var _nw48 *C7 = New_C7_()
	_ = _nw48
	_nw48.Ctor__(_242_v21, ((_295_v52).Update(_223_v7, Companion_Default___.Abs(_216_v1))).Equals(_295_v52), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(384))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer32(arg32)
		}
	}((func(_297_v19 _dafny.Array) func(_dafny.Int) _dafny.Int {
		return func(_298_i9 _dafny.Int) _dafny.Int {
			return (_297_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_297_v19), 0))).Int()).(_dafny.Int)
		}
	})(_240_v19))))
	_296_v53 = _nw48
	var _299_v54 D3
	_ = _299_v54
	_299_v54 = Companion_D3_.Create_DC10_(_223_v7)
	_299_v54 = _299_v54
	var _300_v55 _dafny.Set
	_ = _300_v55
	_300_v55 = _dafny.SetOf(_223_v7, !(_223_v7))
	var _hi1 _dafny.Int = (func() _dafny.Int {
		if (_295_v52).Contains(_223_v7) {
			return (_295_v52).Multiplicity(_223_v7)
		}
		return (_300_v55).Cardinality()
	})()
	_ = _hi1
	for _301_i10 := _216_v1; _301_i10.Cmp(_hi1) < 0; _301_i10 = _301_i10.Plus(_dafny.One) {
		if _223_v7 {
			var _302_v56 _dafny.Array
			_ = _302_v56
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(13))
			_ = _nw49
			_302_v56 = _nw49
			var _303_v57 *C8
			_ = _303_v57
			var _nw50 *C8 = New_C8_()
			_ = _nw50
			_nw50.Ctor__(_215_v0, _300_v55, _223_v7, _dafny.SeqOf(_dafny.IntOfInt64(529), _216_v1))
			_303_v57 = _nw50
			var _304_v58 _dafny.Map
			_ = _304_v58
			_304_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v57, _223_v7)
			var _305_v59 _dafny.Map
			_ = _305_v59
			_305_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v58, _302_v56)
			var _306_v60 _dafny.Array
			_ = _306_v60
			var _nwElement0_11 _dafny.Array = _302_v56
			_ = _nwElement0_11
			var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(8))
			_ = _nw51
			(_nw51).ArraySet1(_nwElement0_11, 0)
			(_nw51).ArraySet1(_302_v56, 1)
			(_nw51).ArraySet1(_302_v56, 2)
			(_nw51).ArraySet1((func() _dafny.Array {
				if (_305_v59).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v57, _223_v7)) {
					return (_305_v59).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_303_v57, _223_v7)).(_dafny.Array)
				}
				return _302_v56
			})(), 3)
			(_nw51).ArraySet1(_302_v56, 4)
			(_nw51).ArraySet1(_302_v56, 5)
			(_nw51).ArraySet1(_302_v56, 6)
			(_nw51).ArraySet1(_302_v56, 7)
			_306_v60 = _nw51
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_306_v60), 0))
			_ = _index64
			(_306_v60).ArraySet1(_302_v56, (_index64).Int())
			var _307_v61 _dafny.Map
			_ = _307_v61
			_307_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v7, (_dafny.Zero).Minus((_dafny.Zero).Minus((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int))))
			var _308_v62 _dafny.Map
			_ = _308_v62
			_308_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_216_v1).Times(_216_v1), _dafny.Companion_Sequence_.Concatenate(_226_v10, _226_v10))
			var _309_v63 _dafny.MultiSet
			_ = _309_v63
			_309_v63 = _dafny.MultiSetOf(_dafny.MultiSetOf(true), _295_v52, _295_v52)
			var _310_v64 T1
			_ = _310_v64
			var _nw52 *C1 = New_C1_()
			_ = _nw52
			_nw52.Ctor__((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _216_v1, Companion_Default___.Fm0(_223_v7, _223_v7, _230_globalState), _229_v13, _216_v1)
			_310_v64 = _nw52
			var _311_v65 T0
			_ = _311_v65
			var _nw53 *C1 = New_C1_()
			_ = _nw53
			_nw53.Ctor__((_307_v61).Cardinality(), (_dafny.MultiSetOf(_310_v64, _310_v64)).Cardinality(), _223_v7, (_310_v64).F17(), (_310_v64).F28())
			_311_v65 = _nw53
			var _312_v66 _dafny.Map
			_ = _312_v66
			_312_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _311_v65)
			var _313_v67 *C3
			_ = _313_v67
			var _nw54 *C3 = New_C3_()
			_ = _nw54
			_nw54.Ctor__(_215_v0, (Companion_D1_.Create_DC5_((_303_v57.F19).Cardinality(), _311_v65.F16())).Dtor_cf11(), (_311_v65).F17())
			_313_v67 = _nw54
			var _314_v68 _dafny.Map
			_ = _314_v68
			_314_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_313_v67, (Companion_Default___.Fm31(_310_v64.F16(), _230_globalState)).Cardinality())
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_306_v60), 0))
			_ = _index65
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index66
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index67
			var _rhs51 _dafny.Array = _302_v56
			_ = _rhs51
			var _rhs52 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_309_v63).Contains(_295_v52), (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int))
			_ = _rhs52
			var _rhs53 _dafny.Int = (_301_i10).Minus((func() _dafny.Int {
				if _223_v7 {
					return (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)
				}
				return (_312_v66).Cardinality()
			})())
			_ = _rhs53
			var _rhs54 _dafny.Map = _308_v62
			_ = _rhs54
			var _rhs55 _dafny.Int = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_229_v13).Cardinality()), ((_295_v52).Update(false, Companion_Default___.Abs((func() _dafny.Int {
				if (_314_v68).Contains(_313_v67) {
					return (_314_v68).Get(_313_v67).(_dafny.Int)
				}
				return (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)
			})()))).Cardinality())).Times(Companion_Default___.SafeDivisionInt(_301_i10, _301_i10))
			_ = _rhs55
			var _lhs56 _dafny.Array = _306_v60
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(448), _dafny.ArrayLen((_306_v60), 0))
			_ = _lhs57
			var _lhs58 _dafny.Array = _240_v19
			_ = _lhs58
			var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs59
			var _lhs60 _dafny.Array = _240_v19
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs61
			(_lhs56).ArraySet1(_rhs51, (_lhs57).Int())
			_307_v61 = _rhs52
			(_lhs58).ArraySet1(_rhs53, (_lhs59).Int())
			_308_v62 = _rhs54
			(_lhs60).ArraySet1(_rhs55, (_lhs61).Int())
			var _315_v69 _dafny.Map
			_ = _315_v69
			_315_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_217_v2, _296_v53), (_dafny.Zero).Minus((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)))
			var _316_v70 _dafny.Map
			_ = _316_v70
			_316_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_301_i10, _296_v53)
			var _rhs56 _dafny.Map = _315_v69
			_ = _rhs56
			var _rhs57 bool = !((_316_v70).Merge(_316_v70)).Equals((_316_v70).Merge(_316_v70))
			_ = _rhs57
			_315_v69 = _rhs56
			_223_v7 = _rhs57
			var _317_v71 _dafny.Sequence
			_ = _317_v71
			_317_v71 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_318_v8 _dafny.Map) func(_dafny.Int) _dafny.Int {
				return func(_319_i11 _dafny.Int) _dafny.Int {
					return (_318_v8).Cardinality()
				}
			})(_224_v8))), (_311_v65).F17(), (_311_v65).F17(), (_311_v65).F17(), (_310_v64).F17())
			var _320_v72 *C1
			_ = _320_v72
			var _nw55 *C1 = New_C1_()
			_ = _nw55
			_nw55.Ctor__(_216_v1, _216_v1, false, _dafny.Companion_Sequence_.Concatenate((_317_v71).Select((Companion_Default___.SafeIndex((_310_v64).F28(), _dafny.IntOfUint32((_317_v71).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.SeqOf(_216_v1, _dafny.IntOfInt64(427))), _301_i10)
			_320_v72 = _nw55
			var _321_v75 _dafny.Map
			_ = _321_v75
			_321_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_300_v55, (Companion_Default___.Fm26(_230_globalState)).Dtor_cf35())
			var _322_v76 _dafny.Set
			_ = _322_v76
			_322_v76 = _dafny.SetOf((func() _dafny.Map {
				var _coll27 = _dafny.NewMapBuilder()
				_ = _coll27
				for _iter29 := _dafny.Iterate((func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter30 := _dafny.Iterate((_321_v75).Keys().Elements()); ; {
						_compr_28, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _323_v74 _dafny.Set
						_323_v74 = interface{}(_compr_28).(_dafny.Set)
						if (_321_v75).Contains(_323_v74) {
							_coll28.Add(_323_v74, _217_v2)
						}
					}
					return _coll28.ToMap()
				}()).Keys().Elements()); ; {
					_compr_27, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _324_v73 _dafny.Set
					_324_v73 = interface{}(_compr_27).(_dafny.Set)
					if (func() _dafny.Map {
						var _coll29 = _dafny.NewMapBuilder()
						_ = _coll29
						for _iter31 := _dafny.Iterate((_321_v75).Keys().Elements()); ; {
							_compr_29, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _323_v74 _dafny.Set
							_323_v74 = interface{}(_compr_29).(_dafny.Set)
							if (_321_v75).Contains(_323_v74) {
								_coll29.Add(_323_v74, _217_v2)
							}
						}
						return _coll29.ToMap()
					}()).Contains(_324_v73) {
						_coll27.Add(_324_v73, _310_v64.F16())
					}
				}
				return _coll27.ToMap()
			}()).Cardinality(), _dafny.IntOfInt64(-809))
			var _325_v77 _dafny.Set
			_ = _325_v77
			_325_v77 = _dafny.SetOf(_301_i10, (_322_v76).Cardinality(), (_320_v72).F30(), Companion_Default___.SafeModuloInt((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality())))
			_322_v76 = func() _dafny.Set {
				var _coll30 = _dafny.NewBuilder()
				_ = _coll30
				for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-708), _dafny.IntOfInt64(256))); ; {
					_compr_30, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _326_v78 _dafny.Int
					_326_v78 = interface{}(_compr_30).(_dafny.Int)
					if ((_dafny.IntOfInt64(-708)).Cmp(_326_v78) <= 0) && ((_326_v78).Cmp(_dafny.IntOfInt64(256)) < 0) {
						_coll30.Add((_326_v78).Minus((_320_v72).F30()))
					}
				}
				return _coll30.ToSet()
			}()
			var _327_v79 _dafny.Array
			_ = _327_v79
			var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw56
			_327_v79 = _nw56
			var _328_v81 _dafny.Sequence
			_ = _328_v81
			_328_v81 = _dafny.SeqOf(_226_v10)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index68
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index69
			var _rhs58 _dafny.Set = (_325_v77).Intersection(func() _dafny.Set {
				var _coll31 = _dafny.NewBuilder()
				_ = _coll31
				for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(76), _dafny.IntOfInt64(-151))); ; {
					_compr_31, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _329_v80 _dafny.Int
					_329_v80 = interface{}(_compr_31).(_dafny.Int)
					if ((_dafny.IntOfInt64(76)).Cmp(_329_v80) <= 0) && ((_329_v80).Cmp(_dafny.IntOfInt64(-151)) < 0) {
						_coll31.Add((_329_v80).Times((_320_v72).F30()))
					}
				}
				return _coll31.ToSet()
			}())
			_ = _rhs58
			var _rhs59 _dafny.Int = ((_320_v72).F30()).Plus((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int))
			_ = _rhs59
			var _rhs60 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_328_v81).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_226_v10).Cardinality()), _dafny.IntOfUint32((_328_v81).Cardinality()))).Uint32()).(_dafny.Sequence), _226_v10)
			_ = _rhs60
			var _rhs61 _dafny.Array = _327_v79
			_ = _rhs61
			var _rhs62 _dafny.Int = _301_i10
			_ = _rhs62
			var _lhs62 _dafny.Array = _240_v19
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs63
			var _lhs64 _dafny.Array = _240_v19
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs65
			_322_v76 = _rhs58
			(_lhs62).ArraySet1(_rhs59, (_lhs63).Int())
			_226_v10 = _rhs60
			_327_v79 = _rhs61
			(_lhs64).ArraySet1(_rhs62, (_lhs65).Int())
		} else {
			_218_v3 = _218_v3
			var _330_v82 *C3
			_ = _330_v82
			var _nw57 *C3 = New_C3_()
			_ = _nw57
			_nw57.Ctor__(_dafny.Companion_Sequence_.Concatenate((_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm1(_230_globalState)), _223_v7, _229_v13)
			_330_v82 = _nw57
			(_230_globalState).F5 = _223_v7
			var _331_v83 *C0
			_ = _331_v83
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__((_dafny.Zero).Minus((_dafny.Zero).Minus(_216_v1)), _220_v5)
			_331_v83 = _nw58
			var _332_v84 *C3
			_ = _332_v84
			var _nw59 *C3 = New_C3_()
			_ = _nw59
			_nw59.Ctor__((_330_v82).F24(), _223_v7, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_229_v13, (Companion_Default___.SafeIndex((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_229_v13).Cardinality()))).Uint32(), _216_v1), _229_v13))
			_332_v84 = _nw59
			_332_v84 = _330_v82
		}
		_223_v7 = true
		var _nwElement0_12 bool = true
		_ = _nwElement0_12
		var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(2))
		_ = _nw60
		(_nw60).ArraySet1(_nwElement0_12, 0)
		(_nw60).ArraySet1(_223_v7, 1)
		_220_v5 = _nw60
		if _223_v7 {
			var _333_v85 T0
			_ = _333_v85
			var _nw61 *C5 = New_C5_()
			_ = _nw61
			_nw61.Ctor__(_223_v7, _229_v13)
			_333_v85 = _nw61
			var _334_v86 _dafny.Map
			_ = _334_v86
			_334_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v85, _301_i10)
			var _335_v87 _dafny.MultiSet
			_ = _335_v87
			_335_v87 = _dafny.MultiSetOf(_334_v86)
			var _336_v88 _dafny.Map
			_ = _336_v88
			_336_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v7, (func() _dafny.Int {
				if (_335_v87).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v85, _216_v1)) {
					return (_335_v87).Multiplicity(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_333_v85, _216_v1))
				}
				return _301_i10
			})())
			_336_v88 = (_336_v88).Update(false, _216_v1)
			var _337_v89 T1
			_ = _337_v89
			var _nw62 *C1 = New_C1_()
			_ = _nw62
			_nw62.Ctor__(_dafny.IntOfInt64(55), _dafny.IntOfUint32(((_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence)).Cardinality()), _223_v7, (_333_v85).F17(), (_dafny.Zero).Minus((_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)))
			_337_v89 = _nw62
			var _338_v90 _dafny.Map
			_ = _338_v90
			_338_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_230_globalState), _337_v89)
			var _rhs63 bool = (func() bool {
				if (func() bool {
					if _333_v85.F16() {
						return !(_337_v89.F16())
					}
					return _223_v7
				})() {
					return Companion_Default___.Fm0(_223_v7, _333_v85.F16(), _230_globalState)
				}
				return _333_v85.F16()
			})()
			_ = _rhs63
			var _rhs64 T1 = (func() T1 {
				if (_338_v90).Contains(_dafny.CodePoint('v')) {
					return (_338_v90).Get(_dafny.CodePoint('v')).(T1)
				}
				return _337_v89
			})()
			_ = _rhs64
			var _rhs65 _dafny.MultiSet = _295_v52
			_ = _rhs65
			var _lhs66 *GlobalState = _230_globalState
			_ = _lhs66
			_lhs66.F2 = _rhs63
			_337_v89 = _rhs64
			_295_v52 = _rhs65
			var _339_v91 D11
			_ = _339_v91
			_339_v91 = Companion_D11_.Create_DC29_(_217_v2)
			_339_v91 = _339_v91
			var _340_v92 _dafny.Set
			_ = _340_v92
			_340_v92 = _dafny.SetOf((_337_v89).F28(), (_337_v89).F28())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index70
			(_240_v19).ArraySet1(((_340_v92).Intersection(_340_v92)).Cardinality(), (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index71
			(_240_v19).ArraySet1(Companion_Default___.Fm3(_230_globalState), (_index71).Int())
		} else {
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))
			_ = _index72
			(_218_v3).ArraySet1((_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence), (_index72).Int())
			_224_v8 = (_224_v8).Update(_223_v7, !(_223_v7))
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index73
			var _rhs66 _dafny.Int = (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)
			_ = _rhs66
			var _rhs67 bool = _223_v7
			_ = _rhs67
			var _rhs68 _dafny.Int = (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)
			_ = _rhs68
			var _rhs69 _dafny.Int = _301_i10
			_ = _rhs69
			var _lhs67 *GlobalState = _230_globalState
			_ = _lhs67
			var _lhs68 _dafny.Array = _240_v19
			_ = _lhs68
			var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs69
			_216_v1 = _rhs66
			_lhs67.F5 = _rhs67
			_216_v1 = _rhs68
			(_lhs68).ArraySet1(_rhs69, (_lhs69).Int())
			var _341_v93 *C4
			_ = _341_v93
			var _nw63 *C4 = New_C4_()
			_ = _nw63
			_nw63.Ctor__(_dafny.IntOfInt64(578), Companion_Default___.Fm0((func() bool {
				if (_224_v8).Contains(_223_v7) {
					return (_224_v8).Get(_223_v7).(bool)
				}
				return _223_v7
			})(), _223_v7, _230_globalState))
			_341_v93 = _nw63
			var _342_v94 _dafny.Sequence
			_ = _342_v94
			_342_v94 = _dafny.SeqOf(_226_v10)
			var _rhs70 bool = _dafny.Companion_Sequence_.Equal(_226_v10, _dafny.Companion_Sequence_.Concatenate(_226_v10, (_342_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.IntOfUint32((_342_v94).Cardinality()))).Uint32()).(_dafny.Sequence)))
			_ = _rhs70
			var _rhs71 _dafny.CodePoint = Companion_Default___.Fm2(_230_globalState)
			_ = _rhs71
			var _rhs72 *C4 = _341_v93
			_ = _rhs72
			_223_v7 = _rhs70
			_217_v2 = _rhs71
			_341_v93 = _rhs72
			var _343_v95 *C3
			_ = _343_v95
			var _nw64 *C3 = New_C3_()
			_ = _nw64
			_nw64.Ctor__((_218_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(595), _dafny.ArrayLen((_218_v3), 0))).Int()).(_dafny.Sequence), _223_v7, _dafny.SeqOf(_dafny.IntOfInt64(-573), (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int)))
			_343_v95 = _nw64
			var _344_v96 _dafny.Array
			_ = _344_v96
			var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_8
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_8.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_8)
			} else {
				var _init8 func(_dafny.Int) _dafny.CodePoint = (func(_345_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_346_i12 _dafny.Int) _dafny.CodePoint {
						return _345_v2
					}
				})(_217_v2)
				_ = _init8
				var _element0_8 = _init8(_dafny.Zero)
				_ = _element0_8
				_nw65 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
				(_nw65).ArraySet1CodePoint(_element0_8, 0)
				var _nativeLen0_8 = (_len0_8).Int()
				_ = _nativeLen0_8
				for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
					(_nw65).ArraySet1CodePoint(_init8(_dafny.IntOf(_i0_8)), _i0_8)
				}
			}
			_344_v96 = _nw65
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_344_v96), 0))
			_ = _index74
			(_344_v96).ArraySet1CodePoint(_217_v2, (_index74).Int())
			var _347_v97 *C8
			_ = _347_v97
			var _nw66 *C8 = New_C8_()
			_ = _nw66
			_nw66.Ctor__(_215_v0, _300_v55, _223_v7, _dafny.SeqOf((_229_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(814), _dafny.IntOfUint32((_229_v13).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm17(_230_globalState)).Cardinality()), (_240_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))).Int()).(_dafny.Int), (_229_v13).Select((Companion_Default___.SafeIndex(_301_i10, _dafny.IntOfUint32((_229_v13).Cardinality()))).Uint32()).(_dafny.Int)))
			_347_v97 = _nw66
			var _348_v98 _dafny.Map
			_ = _348_v98
			_348_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_347_v97, _223_v7)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _index75
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_344_v96), 0))
			_ = _index76
			var _rhs73 bool = (func() bool {
				if (_348_v98).Contains(_347_v97) {
					return (_348_v98).Get(_347_v97).(bool)
				}
				return _223_v7
			})()
			_ = _rhs73
			var _rhs74 _dafny.Int = (_341_v93).F22()
			_ = _rhs74
			var _rhs75 *C3 = _343_v95
			_ = _rhs75
			var _rhs76 _dafny.CodePoint = _dafny.CodePoint('i')
			_ = _rhs76
			var _lhs70 *C4 = _341_v93
			_ = _lhs70
			var _lhs71 _dafny.Array = _240_v19
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_240_v19), 0))
			_ = _lhs72
			var _lhs73 _dafny.Array = _344_v96
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_344_v96), 0))
			_ = _lhs74
			_lhs70.F23 = _rhs73
			(_lhs71).ArraySet1(_rhs74, (_lhs72).Int())
			_343_v95 = _rhs75
			(_lhs73).ArraySet1CodePoint(_rhs76, (_lhs74).Int())
		}
	}
	_dafny.Print(_215_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_216_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_217_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_219_v4).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v5).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_222_v6).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_223_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_225_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(879), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_226_v10, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_228_v12, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_229_v13, _dafny.SeqOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(879))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F0()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_230_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_230_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F4()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_230_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState.F8).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(318), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_230_globalState).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_230_globalState).F10(), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(879), _dafny.IntOfInt64(2)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_230_globalState).F12(), _dafny.SeqOf(_dafny.IntOfInt64(879), _dafny.IntOfInt64(879))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_230_globalState).F13(), _dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('t'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_230_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v14).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_234_v15).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_237_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v18).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_v19).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_241_v20).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_243_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v31).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('t')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v32).Equals(_dafny.MultiSetOf(_dafny.CodePoint('t'), _dafny.CodePoint('t'), _dafny.CodePoint('s'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v52).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_296_v53).F20()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_299_v54).Dtor_cf19())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_300_v55).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
}

// End of class Default__

// Definition of datatype D0
type D0 struct {
	Data_D0_
}

func (_this D0) Get_() Data_D0_ {
	return _this.Data_D0_
}

type Data_D0_ interface {
	isD0()
}

type CompanionStruct_D0_ struct {
}

var Companion_D0_ = CompanionStruct_D0_{}

type D0_DC1 struct {
	Cf1 _dafny.Sequence
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Sequence) D0 {
	return D0{D0_DC1{Cf1}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf2 _dafny.Int
	Cf3 _dafny.Int
	Cf4 _dafny.Int
	Cf5 _dafny.Array
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.Int, Cf5 _dafny.Array) D0 {
	return D0{D0_DC2{Cf2, Cf3, Cf4, Cf5}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.EmptySeq)
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + data.Cf1.VerbatimString(true) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D0) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D0)
	return ok && _this.Equals(typed)
}

func Type_D0_() _dafny.TypeDescriptor {
	return type_D0_{}
}

type type_D0_ struct {
}

func (_this type_D0_) Default() interface{} {
	return Companion_D0_.Default()
}

func (_this type_D0_) String() string {
	return "main.D0"
}
func (_this D0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D0{}

// End of datatype D0

// Definition of datatype D1
type D1 struct {
	Data_D1_
}

func (_this D1) Get_() Data_D1_ {
	return _this.Data_D1_
}

type Data_D1_ interface {
	isD1()
}

type CompanionStruct_D1_ struct {
}

var Companion_D1_ = CompanionStruct_D1_{}

type D1_DC4 struct {
	Cf7 bool
	Cf8 _dafny.Int
	Cf9 _dafny.Map
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 bool, Cf8 _dafny.Int, Cf9 _dafny.Map) D1 {
	return D1{D1_DC4{Cf7, Cf8, Cf9}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf10 _dafny.Int
	Cf11 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf10 _dafny.Int, Cf11 bool) D1 {
	return D1{D1_DC5{Cf10, Cf11}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_() D1 {
	return D1{D1_DC6{}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.MultiSet
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.MultiSet) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(false, _dafny.Zero, _dafny.EmptyMap)
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf6() _dafny.MultiSet {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf6) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Equals(data2.Cf9)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11 == data2.Cf11
		}
	case D1_DC6:
		{
			_, ok := other.Get_().(D1_DC6)
			return ok
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D1)
	return ok && _this.Equals(typed)
}

func Type_D1_() _dafny.TypeDescriptor {
	return type_D1_{}
}

type type_D1_ struct {
}

func (_this type_D1_) Default() interface{} {
	return Companion_D1_.Default()
}

func (_this type_D1_) String() string {
	return "main.D1"
}
func (_this D1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D1{}

// End of datatype D1

// Definition of datatype D2
type D2 struct {
	Data_D2_
}

func (_this D2) Get_() Data_D2_ {
	return _this.Data_D2_
}

type Data_D2_ interface {
	isD2()
}

type CompanionStruct_D2_ struct {
}

var Companion_D2_ = CompanionStruct_D2_{}

type D2_DC8 struct {
	Cf13 _dafny.Map
	Cf14 _dafny.Array
	Cf15 bool
	Cf16 bool
	Cf17 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf13 _dafny.Map, Cf14 _dafny.Array, Cf15 bool, Cf16 bool, Cf17 _dafny.Int) D2 {
	return D2{D2_DC8{Cf13, Cf14, Cf15, Cf16, Cf17}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf12 _dafny.Set
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf12 _dafny.Set) D2 {
	return D2{D2_DC7{Cf12}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false, _dafny.Zero)
}

func (_this D2) Dtor_cf13() _dafny.Map {
	return _this.Get_().(D2_DC8).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D2_DC8).Cf14
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC8).Cf15
}

func (_this D2) Dtor_cf16() bool {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf12() _dafny.Set {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf13.Equals(data2.Cf13) && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf12.Equals(data2.Cf12)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D2) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D2)
	return ok && _this.Equals(typed)
}

func Type_D2_() _dafny.TypeDescriptor {
	return type_D2_{}
}

type type_D2_ struct {
}

func (_this type_D2_) Default() interface{} {
	return Companion_D2_.Default()
}

func (_this type_D2_) String() string {
	return "main.D2"
}
func (_this D2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D2{}

// End of datatype D2

// Definition of datatype D3
type D3 struct {
	Data_D3_
}

func (_this D3) Get_() Data_D3_ {
	return _this.Data_D3_
}

type Data_D3_ interface {
	isD3()
}

type CompanionStruct_D3_ struct {
}

var Companion_D3_ = CompanionStruct_D3_{}

type D3_DC10 struct {
	Cf19 bool
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf19 bool) D3 {
	return D3{D3_DC10{Cf19}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC9 struct {
	Cf18 _dafny.Sequence
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf18 _dafny.Sequence) D3 {
	return D3{D3_DC9{Cf18}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false)
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D3_DC9).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf19 == data2.Cf19
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D3) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D3)
	return ok && _this.Equals(typed)
}

func Type_D3_() _dafny.TypeDescriptor {
	return type_D3_{}
}

type type_D3_ struct {
}

func (_this type_D3_) Default() interface{} {
	return Companion_D3_.Default()
}

func (_this type_D3_) String() string {
	return "main.D3"
}
func (_this D3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D3{}

// End of datatype D3

// Definition of datatype D4
type D4 struct {
	Data_D4_
}

func (_this D4) Get_() Data_D4_ {
	return _this.Data_D4_
}

type Data_D4_ interface {
	isD4()
}

type CompanionStruct_D4_ struct {
}

var Companion_D4_ = CompanionStruct_D4_{}

type D4_DC12 struct {
	Cf21 bool
	Cf22 bool
	Cf23 _dafny.Int
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf21 bool, Cf22 bool, Cf23 _dafny.Int) D4 {
	return D4{D4_DC12{Cf21, Cf22, Cf23}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf24 _dafny.Int
	Cf25 _dafny.Int
	Cf26 _dafny.Int
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf24 _dafny.Int, Cf25 _dafny.Int, Cf26 _dafny.Int) D4 {
	return D4{D4_DC13{Cf24, Cf25, Cf26}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC11 struct {
	Cf20 _dafny.MultiSet
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf20 _dafny.MultiSet) D4 {
	return D4{D4_DC11{Cf20}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(false, false, _dafny.Zero)
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf22() bool {
	return _this.Get_().(D4_DC12).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf25
}

func (_this D4) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf26
}

func (_this D4) Dtor_cf20() _dafny.MultiSet {
	return _this.Get_().(D4_DC11).Cf20
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Cmp(data2.Cf26) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf20.Equals(data2.Cf20)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D4) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D4)
	return ok && _this.Equals(typed)
}

func Type_D4_() _dafny.TypeDescriptor {
	return type_D4_{}
}

type type_D4_ struct {
}

func (_this type_D4_) Default() interface{} {
	return Companion_D4_.Default()
}

func (_this type_D4_) String() string {
	return "main.D4"
}
func (_this D4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D4{}

// End of datatype D4

// Definition of datatype D5
type D5 struct {
	Data_D5_
}

func (_this D5) Get_() Data_D5_ {
	return _this.Data_D5_
}

type Data_D5_ interface {
	isD5()
}

type CompanionStruct_D5_ struct {
}

var Companion_D5_ = CompanionStruct_D5_{}

type D5_DC14 struct {
	Cf27 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf27 _dafny.Array) D5 {
	return D5{D5_DC14{Cf27}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf27() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf27 == data2.Cf27
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D5) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D5)
	return ok && _this.Equals(typed)
}

func Type_D5_() _dafny.TypeDescriptor {
	return type_D5_{}
}

type type_D5_ struct {
}

func (_this type_D5_) Default() interface{} {
	return Companion_D5_.Default()
}

func (_this type_D5_) String() string {
	return "main.D5"
}
func (_this D5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D5{}

// End of datatype D5

// Definition of datatype D6
type D6 struct {
	Data_D6_
}

func (_this D6) Get_() Data_D6_ {
	return _this.Data_D6_
}

type Data_D6_ interface {
	isD6()
}

type CompanionStruct_D6_ struct {
}

var Companion_D6_ = CompanionStruct_D6_{}

type D6_DC16 struct {
	Cf29 bool
	Cf30 bool
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf29 bool, Cf30 bool) D6 {
	return D6{D6_DC16{Cf29, Cf30}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC17 struct {
	Cf31 D2
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf31 D2) D6 {
	return D6{D6_DC17{Cf31}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf32 _dafny.Array
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf32 _dafny.Array) D6 {
	return D6{D6_DC18{Cf32}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC15 struct {
	Cf28 _dafny.Array
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf28 _dafny.Array) D6 {
	return D6{D6_DC15{Cf28}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(false, false)
}

func (_this D6) Dtor_cf29() bool {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) Dtor_cf30() bool {
	return _this.Get_().(D6_DC16).Cf30
}

func (_this D6) Dtor_cf31() D2 {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Array {
	return _this.Get_().(D6_DC18).Cf32
}

func (_this D6) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D6_DC15).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf31) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf31.Equals(data2.Cf31)
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf32 == data2.Cf32
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf28 == data2.Cf28
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D6) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D6)
	return ok && _this.Equals(typed)
}

func Type_D6_() _dafny.TypeDescriptor {
	return type_D6_{}
}

type type_D6_ struct {
}

func (_this type_D6_) Default() interface{} {
	return Companion_D6_.Default()
}

func (_this type_D6_) String() string {
	return "main.D6"
}
func (_this D6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D6{}

// End of datatype D6

// Definition of datatype D7
type D7 struct {
	Data_D7_
}

func (_this D7) Get_() Data_D7_ {
	return _this.Data_D7_
}

type Data_D7_ interface {
	isD7()
}

type CompanionStruct_D7_ struct {
}

var Companion_D7_ = CompanionStruct_D7_{}

type D7_DC19 struct {
	Cf33 _dafny.Sequence
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf33 _dafny.Sequence) D7 {
	return D7{D7_DC19{Cf33}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D7) Dtor_cf33() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf33.Equals(data2.Cf33)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D7) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D7)
	return ok && _this.Equals(typed)
}

func Type_D7_() _dafny.TypeDescriptor {
	return type_D7_{}
}

type type_D7_ struct {
}

func (_this type_D7_) Default() interface{} {
	return Companion_D7_.Default()
}

func (_this type_D7_) String() string {
	return "main.D7"
}
func (_this D7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D7{}

// End of datatype D7

// Definition of datatype D8
type D8 struct {
	Data_D8_
}

func (_this D8) Get_() Data_D8_ {
	return _this.Data_D8_
}

type Data_D8_ interface {
	isD8()
}

type CompanionStruct_D8_ struct {
}

var Companion_D8_ = CompanionStruct_D8_{}

type D8_DC21 struct {
	Cf35 _dafny.Int
	Cf36 bool
	Cf37 bool
	Cf38 bool
	Cf39 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf35 _dafny.Int, Cf36 bool, Cf37 bool, Cf38 bool, Cf39 _dafny.Int) D8 {
	return D8{D8_DC21{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC22 struct {
	Cf40 bool
	Cf41 T0
	Cf42 bool
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf40 bool, Cf41 T0, Cf42 bool) D8 {
	return D8{D8_DC22{Cf40, Cf41, Cf42}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC23 struct {
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_() D8 {
	return D8{D8_DC23{}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC20 struct {
	Cf34 _dafny.Sequence
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf34 _dafny.Sequence) D8 {
	return D8{D8_DC20{Cf34}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC24 struct {
	Cf43 D8
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf43 D8) D8 {
	return D8{D8_DC24{Cf43}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.Zero, false, false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf35
}

func (_this D8) Dtor_cf36() bool {
	return _this.Get_().(D8_DC21).Cf36
}

func (_this D8) Dtor_cf37() bool {
	return _this.Get_().(D8_DC21).Cf37
}

func (_this D8) Dtor_cf38() bool {
	return _this.Get_().(D8_DC21).Cf38
}

func (_this D8) Dtor_cf39() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf39
}

func (_this D8) Dtor_cf40() bool {
	return _this.Get_().(D8_DC22).Cf40
}

func (_this D8) Dtor_cf41() T0 {
	return _this.Get_().(D8_DC22).Cf41
}

func (_this D8) Dtor_cf42() bool {
	return _this.Get_().(D8_DC22).Cf42
}

func (_this D8) Dtor_cf34() _dafny.Sequence {
	return _this.Get_().(D8_DC20).Cf34
}

func (_this D8) Dtor_cf43() D8 {
	return _this.Get_().(D8_DC24).Cf43
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38 && data1.Cf39.Cmp(data2.Cf39) == 0
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf40 == data2.Cf40 && _dafny.AreEqual(data1.Cf41, data2.Cf41) && data1.Cf42 == data2.Cf42
		}
	case D8_DC23:
		{
			_, ok := other.Get_().(D8_DC23)
			return ok
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf34.Equals(data2.Cf34)
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D8) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D8)
	return ok && _this.Equals(typed)
}

func Type_D8_() _dafny.TypeDescriptor {
	return type_D8_{}
}

type type_D8_ struct {
}

func (_this type_D8_) Default() interface{} {
	return Companion_D8_.Default()
}

func (_this type_D8_) String() string {
	return "main.D8"
}
func (_this D8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D8{}

// End of datatype D8

// Definition of datatype D9
type D9 struct {
	Data_D9_
}

func (_this D9) Get_() Data_D9_ {
	return _this.Data_D9_
}

type Data_D9_ interface {
	isD9()
}

type CompanionStruct_D9_ struct {
}

var Companion_D9_ = CompanionStruct_D9_{}

type D9_DC26 struct {
	Cf45 _dafny.Int
	Cf46 *C0
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf45 _dafny.Int, Cf46 *C0) D9 {
	return D9{D9_DC26{Cf45, Cf46}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC27 struct {
	Cf47 bool
	Cf48 _dafny.Set
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf47 bool, Cf48 _dafny.Set) D9 {
	return D9{D9_DC27{Cf47, Cf48}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC25 struct {
	Cf44 _dafny.Map
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf44 _dafny.Map) D9 {
	return D9{D9_DC25{Cf44}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC26_(_dafny.Zero, (*C0)(nil))
}

func (_this D9) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf45
}

func (_this D9) Dtor_cf46() *C0 {
	return _this.Get_().(D9_DC26).Cf46
}

func (_this D9) Dtor_cf47() bool {
	return _this.Get_().(D9_DC27).Cf47
}

func (_this D9) Dtor_cf48() _dafny.Set {
	return _this.Get_().(D9_DC27).Cf48
}

func (_this D9) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D9_DC25).Cf44
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46 == data2.Cf46
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf47 == data2.Cf47 && data1.Cf48.Equals(data2.Cf48)
		}
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf44.Equals(data2.Cf44)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D9) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D9)
	return ok && _this.Equals(typed)
}

func Type_D9_() _dafny.TypeDescriptor {
	return type_D9_{}
}

type type_D9_ struct {
}

func (_this type_D9_) Default() interface{} {
	return Companion_D9_.Default()
}

func (_this type_D9_) String() string {
	return "main.D9"
}
func (_this D9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D9{}

// End of datatype D9

// Definition of datatype D10
type D10 struct {
	Data_D10_
}

func (_this D10) Get_() Data_D10_ {
	return _this.Data_D10_
}

type Data_D10_ interface {
	isD10()
}

type CompanionStruct_D10_ struct {
}

var Companion_D10_ = CompanionStruct_D10_{}

type D10_DC28 struct {
	Cf49 _dafny.Array
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf49 _dafny.Array) D10 {
	return D10{D10_DC28{Cf49}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf49() _dafny.Array {
	return _this.Get_().(D10_DC28).Cf49
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf49 == data2.Cf49
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D10) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D10)
	return ok && _this.Equals(typed)
}

func Type_D10_() _dafny.TypeDescriptor {
	return type_D10_{}
}

type type_D10_ struct {
}

func (_this type_D10_) Default() interface{} {
	return Companion_D10_.Default()
}

func (_this type_D10_) String() string {
	return "main.D10"
}
func (_this D10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D10{}

// End of datatype D10

// Definition of datatype D11
type D11 struct {
	Data_D11_
}

func (_this D11) Get_() Data_D11_ {
	return _this.Data_D11_
}

type Data_D11_ interface {
	isD11()
}

type CompanionStruct_D11_ struct {
}

var Companion_D11_ = CompanionStruct_D11_{}

type D11_DC30 struct {
	Cf51 bool
	Cf52 _dafny.Int
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf51 bool, Cf52 _dafny.Int) D11 {
	return D11{D11_DC30{Cf51, Cf52}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC31 struct {
	Cf53 bool
	Cf54 _dafny.Int
	Cf55 _dafny.Int
	Cf56 bool
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf53 bool, Cf54 _dafny.Int, Cf55 _dafny.Int, Cf56 bool) D11 {
	return D11{D11_DC31{Cf53, Cf54, Cf55, Cf56}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

type D11_DC29 struct {
	Cf50 _dafny.CodePoint
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf50 _dafny.CodePoint) D11 {
	return D11{D11_DC29{Cf50}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

type D11_DC32 struct {
	Cf57 D11
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf57 D11) D11 {
	return D11{D11_DC32{Cf57}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(false, _dafny.Zero)
}

func (_this D11) Dtor_cf51() bool {
	return _this.Get_().(D11_DC30).Cf51
}

func (_this D11) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D11_DC30).Cf52
}

func (_this D11) Dtor_cf53() bool {
	return _this.Get_().(D11_DC31).Cf53
}

func (_this D11) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf54
}

func (_this D11) Dtor_cf55() _dafny.Int {
	return _this.Get_().(D11_DC31).Cf55
}

func (_this D11) Dtor_cf56() bool {
	return _this.Get_().(D11_DC31).Cf56
}

func (_this D11) Dtor_cf50() _dafny.CodePoint {
	return _this.Get_().(D11_DC29).Cf50
}

func (_this D11) Dtor_cf57() D11 {
	return _this.Get_().(D11_DC32).Cf57
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55.Cmp(data2.Cf55) == 0 && data1.Cf56 == data2.Cf56
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf50 == data2.Cf50
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D11) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D11)
	return ok && _this.Equals(typed)
}

func Type_D11_() _dafny.TypeDescriptor {
	return type_D11_{}
}

type type_D11_ struct {
}

func (_this type_D11_) Default() interface{} {
	return Companion_D11_.Default()
}

func (_this type_D11_) String() string {
	return "main.D11"
}
func (_this D11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D11{}

// End of datatype D11

// Definition of datatype D12
type D12 struct {
	Data_D12_
}

func (_this D12) Get_() Data_D12_ {
	return _this.Data_D12_
}

type Data_D12_ interface {
	isD12()
}

type CompanionStruct_D12_ struct {
}

var Companion_D12_ = CompanionStruct_D12_{}

type D12_DC34 struct {
	Cf59 _dafny.Array
	Cf60 _dafny.Int
	Cf61 _dafny.Array
	Cf62 _dafny.CodePoint
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf59 _dafny.Array, Cf60 _dafny.Int, Cf61 _dafny.Array, Cf62 _dafny.CodePoint) D12 {
	return D12{D12_DC34{Cf59, Cf60, Cf61, Cf62}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC35 struct {
	Cf63 bool
	Cf64 bool
	Cf65 _dafny.Map
	Cf66 bool
	Cf67 *C3
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf63 bool, Cf64 bool, Cf65 _dafny.Map, Cf66 bool, Cf67 *C3) D12 {
	return D12{D12_DC35{Cf63, Cf64, Cf65, Cf66, Cf67}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC33 struct {
	Cf58 _dafny.Array
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf58 _dafny.Array) D12 {
	return D12{D12_DC33{Cf58}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC36 struct {
	Cf68 D12
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf68 D12) D12 {
	return D12{D12_DC36{Cf68}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC34_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.CodePoint('D'))
}

func (_this D12) Dtor_cf59() _dafny.Array {
	return _this.Get_().(D12_DC34).Cf59
}

func (_this D12) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D12_DC34).Cf60
}

func (_this D12) Dtor_cf61() _dafny.Array {
	return _this.Get_().(D12_DC34).Cf61
}

func (_this D12) Dtor_cf62() _dafny.CodePoint {
	return _this.Get_().(D12_DC34).Cf62
}

func (_this D12) Dtor_cf63() bool {
	return _this.Get_().(D12_DC35).Cf63
}

func (_this D12) Dtor_cf64() bool {
	return _this.Get_().(D12_DC35).Cf64
}

func (_this D12) Dtor_cf65() _dafny.Map {
	return _this.Get_().(D12_DC35).Cf65
}

func (_this D12) Dtor_cf66() bool {
	return _this.Get_().(D12_DC35).Cf66
}

func (_this D12) Dtor_cf67() *C3 {
	return _this.Get_().(D12_DC35).Cf67
}

func (_this D12) Dtor_cf58() _dafny.Array {
	return _this.Get_().(D12_DC33).Cf58
}

func (_this D12) Dtor_cf68() D12 {
	return _this.Get_().(D12_DC36).Cf68
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf59 == data2.Cf59 && data1.Cf60.Cmp(data2.Cf60) == 0 && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62
		}
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf63 == data2.Cf63 && data1.Cf64 == data2.Cf64 && data1.Cf65.Equals(data2.Cf65) && data1.Cf66 == data2.Cf66 && data1.Cf67 == data2.Cf67
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf58 == data2.Cf58
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D12) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D12)
	return ok && _this.Equals(typed)
}

func Type_D12_() _dafny.TypeDescriptor {
	return type_D12_{}
}

type type_D12_ struct {
}

func (_this type_D12_) Default() interface{} {
	return Companion_D12_.Default()
}

func (_this type_D12_) String() string {
	return "main.D12"
}
func (_this D12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D12{}

// End of datatype D12

// Definition of datatype D13
type D13 struct {
	Data_D13_
}

func (_this D13) Get_() Data_D13_ {
	return _this.Data_D13_
}

type Data_D13_ interface {
	isD13()
}

type CompanionStruct_D13_ struct {
}

var Companion_D13_ = CompanionStruct_D13_{}

type D13_DC37 struct {
	Cf69 _dafny.Map
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf69 _dafny.Map) D13 {
	return D13{D13_DC37{Cf69}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf69() _dafny.Map {
	return _this.Get_().(D13_DC37).Cf69
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf69) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf69.Equals(data2.Cf69)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D13) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D13)
	return ok && _this.Equals(typed)
}

func Type_D13_() _dafny.TypeDescriptor {
	return type_D13_{}
}

type type_D13_ struct {
}

func (_this type_D13_) Default() interface{} {
	return Companion_D13_.Default()
}

func (_this type_D13_) String() string {
	return "main.D13"
}
func (_this D13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D13{}

// End of datatype D13

// Definition of datatype D14
type D14 struct {
	Data_D14_
}

func (_this D14) Get_() Data_D14_ {
	return _this.Data_D14_
}

type Data_D14_ interface {
	isD14()
}

type CompanionStruct_D14_ struct {
}

var Companion_D14_ = CompanionStruct_D14_{}

type D14_DC39 struct {
	Cf71 bool
	Cf72 _dafny.Int
}

func (D14_DC39) isD14() {}

func (CompanionStruct_D14_) Create_DC39_(Cf71 bool, Cf72 _dafny.Int) D14 {
	return D14{D14_DC39{Cf71, Cf72}}
}

func (_this D14) Is_DC39() bool {
	_, ok := _this.Get_().(D14_DC39)
	return ok
}

type D14_DC38 struct {
	Cf70 _dafny.Sequence
}

func (D14_DC38) isD14() {}

func (CompanionStruct_D14_) Create_DC38_(Cf70 _dafny.Sequence) D14 {
	return D14{D14_DC38{Cf70}}
}

func (_this D14) Is_DC38() bool {
	_, ok := _this.Get_().(D14_DC38)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC39_(false, _dafny.Zero)
}

func (_this D14) Dtor_cf71() bool {
	return _this.Get_().(D14_DC39).Cf71
}

func (_this D14) Dtor_cf72() _dafny.Int {
	return _this.Get_().(D14_DC39).Cf72
}

func (_this D14) Dtor_cf70() _dafny.Sequence {
	return _this.Get_().(D14_DC38).Cf70
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC39:
		{
			return "D14.DC39" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D14_DC38:
		{
			return "D14.DC38" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC39:
		{
			data2, ok := other.Get_().(D14_DC39)
			return ok && data1.Cf71 == data2.Cf71 && data1.Cf72.Cmp(data2.Cf72) == 0
		}
	case D14_DC38:
		{
			data2, ok := other.Get_().(D14_DC38)
			return ok && data1.Cf70.Equals(data2.Cf70)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D14) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D14)
	return ok && _this.Equals(typed)
}

func Type_D14_() _dafny.TypeDescriptor {
	return type_D14_{}
}

type type_D14_ struct {
}

func (_this type_D14_) Default() interface{} {
	return Companion_D14_.Default()
}

func (_this type_D14_) String() string {
	return "main.D14"
}
func (_this D14) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D14{}

// End of datatype D14

// Definition of datatype D15
type D15 struct {
	Data_D15_
}

func (_this D15) Get_() Data_D15_ {
	return _this.Data_D15_
}

type Data_D15_ interface {
	isD15()
}

type CompanionStruct_D15_ struct {
}

var Companion_D15_ = CompanionStruct_D15_{}

type D15_DC41 struct {
	Cf74 _dafny.Int
	Cf75 _dafny.Sequence
	Cf76 _dafny.Sequence
	Cf77 _dafny.Sequence
}

func (D15_DC41) isD15() {}

func (CompanionStruct_D15_) Create_DC41_(Cf74 _dafny.Int, Cf75 _dafny.Sequence, Cf76 _dafny.Sequence, Cf77 _dafny.Sequence) D15 {
	return D15{D15_DC41{Cf74, Cf75, Cf76, Cf77}}
}

func (_this D15) Is_DC41() bool {
	_, ok := _this.Get_().(D15_DC41)
	return ok
}

type D15_DC42 struct {
	Cf78 _dafny.Sequence
	Cf79 bool
	Cf80 D9
}

func (D15_DC42) isD15() {}

func (CompanionStruct_D15_) Create_DC42_(Cf78 _dafny.Sequence, Cf79 bool, Cf80 D9) D15 {
	return D15{D15_DC42{Cf78, Cf79, Cf80}}
}

func (_this D15) Is_DC42() bool {
	_, ok := _this.Get_().(D15_DC42)
	return ok
}

type D15_DC40 struct {
	Cf73 _dafny.Map
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf73 _dafny.Map) D15 {
	return D15{D15_DC40{Cf73}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

type D15_DC43 struct {
	Cf81 D15
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf81 D15) D15 {
	return D15{D15_DC43{Cf81}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC41_(_dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D15) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D15_DC41).Cf74
}

func (_this D15) Dtor_cf75() _dafny.Sequence {
	return _this.Get_().(D15_DC41).Cf75
}

func (_this D15) Dtor_cf76() _dafny.Sequence {
	return _this.Get_().(D15_DC41).Cf76
}

func (_this D15) Dtor_cf77() _dafny.Sequence {
	return _this.Get_().(D15_DC41).Cf77
}

func (_this D15) Dtor_cf78() _dafny.Sequence {
	return _this.Get_().(D15_DC42).Cf78
}

func (_this D15) Dtor_cf79() bool {
	return _this.Get_().(D15_DC42).Cf79
}

func (_this D15) Dtor_cf80() D9 {
	return _this.Get_().(D15_DC42).Cf80
}

func (_this D15) Dtor_cf73() _dafny.Map {
	return _this.Get_().(D15_DC40).Cf73
}

func (_this D15) Dtor_cf81() D15 {
	return _this.Get_().(D15_DC43).Cf81
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC41:
		{
			return "D15.DC41" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ", " + data.Cf77.VerbatimString(true) + ")"
		}
	case D15_DC42:
		{
			return "D15.DC42" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + _dafny.String(data.Cf80) + ")"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf73) + ")"
		}
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf81) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC41:
		{
			data2, ok := other.Get_().(D15_DC41)
			return ok && data1.Cf74.Cmp(data2.Cf74) == 0 && data1.Cf75.Equals(data2.Cf75) && data1.Cf76.Equals(data2.Cf76) && data1.Cf77.Equals(data2.Cf77)
		}
	case D15_DC42:
		{
			data2, ok := other.Get_().(D15_DC42)
			return ok && data1.Cf78.Equals(data2.Cf78) && data1.Cf79 == data2.Cf79 && data1.Cf80.Equals(data2.Cf80)
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf73.Equals(data2.Cf73)
		}
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf81.Equals(data2.Cf81)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D15) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D15)
	return ok && _this.Equals(typed)
}

func Type_D15_() _dafny.TypeDescriptor {
	return type_D15_{}
}

type type_D15_ struct {
}

func (_this type_D15_) Default() interface{} {
	return Companion_D15_.Default()
}

func (_this type_D15_) String() string {
	return "main.D15"
}
func (_this D15) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D15{}

// End of datatype D15

// Definition of datatype D16
type D16 struct {
	Data_D16_
}

func (_this D16) Get_() Data_D16_ {
	return _this.Data_D16_
}

type Data_D16_ interface {
	isD16()
}

type CompanionStruct_D16_ struct {
}

var Companion_D16_ = CompanionStruct_D16_{}

type D16_DC45 struct {
	Cf83 bool
	Cf84 _dafny.Int
	Cf85 _dafny.Int
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf83 bool, Cf84 _dafny.Int, Cf85 _dafny.Int) D16 {
	return D16{D16_DC45{Cf83, Cf84, Cf85}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC46 struct {
	Cf86 _dafny.Map
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf86 _dafny.Map) D16 {
	return D16{D16_DC46{Cf86}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

type D16_DC44 struct {
	Cf82 _dafny.MultiSet
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf82 _dafny.MultiSet) D16 {
	return D16{D16_DC44{Cf82}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

type D16_DC47 struct {
	Cf87 D16
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf87 D16) D16 {
	return D16{D16_DC47{Cf87}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC45_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf83() bool {
	return _this.Get_().(D16_DC45).Cf83
}

func (_this D16) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D16_DC45).Cf84
}

func (_this D16) Dtor_cf85() _dafny.Int {
	return _this.Get_().(D16_DC45).Cf85
}

func (_this D16) Dtor_cf86() _dafny.Map {
	return _this.Get_().(D16_DC46).Cf86
}

func (_this D16) Dtor_cf82() _dafny.MultiSet {
	return _this.Get_().(D16_DC44).Cf82
}

func (_this D16) Dtor_cf87() D16 {
	return _this.Get_().(D16_DC47).Cf87
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D16_DC46:
		{
			return "D16.DC46" + "(" + _dafny.String(data.Cf86) + ")"
		}
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf87) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf83 == data2.Cf83 && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85.Cmp(data2.Cf85) == 0
		}
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf86.Equals(data2.Cf86)
		}
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf82.Equals(data2.Cf82)
		}
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf87.Equals(data2.Cf87)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D16) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D16)
	return ok && _this.Equals(typed)
}

func Type_D16_() _dafny.TypeDescriptor {
	return type_D16_{}
}

type type_D16_ struct {
}

func (_this type_D16_) Default() interface{} {
	return Companion_D16_.Default()
}

func (_this type_D16_) String() string {
	return "main.D16"
}
func (_this D16) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D16{}

// End of datatype D16

// Definition of trait T0
type T0 interface {
	String() string
	F16() bool
	F16_set_(value bool)
	M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	F17() _dafny.Sequence
}
type CompanionStruct_T0_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T0_ = CompanionStruct_T0_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T0_) CastTo_(x interface{}) T0 {
	var t T0
	t, _ = x.(T0)
	return t
}

// End of trait T0

// Definition of trait T1
type T1 interface {
	String() string
	F16() bool
	F16_set_(value bool)
	M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	F17() _dafny.Sequence
	Fm13(p0 bool, globalState *GlobalState) _dafny.Int
	Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	M11(globalState *GlobalState) (_dafny.CodePoint, bool, _dafny.Int)
	F28() _dafny.Int
}
type CompanionStruct_T1_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T1_ = CompanionStruct_T1_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T1_) CastTo_(x interface{}) T1 {
	var t T1
	t, _ = x.(T1)
	return t
}

// End of trait T1

// Definition of class GlobalState
type GlobalState struct {
	F2   bool
	F5   bool
	F8   _dafny.Map
	F15  _dafny.CodePoint
	_f0  _dafny.Array
	_f1  _dafny.Array
	_f3  _dafny.Int
	_f4  _dafny.Array
	_f6  bool
	_f7  _dafny.Int
	_f9  _dafny.Map
	_f10 _dafny.Sequence
	_f11 bool
	_f12 _dafny.Sequence
	_f13 _dafny.Sequence
	_f14 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this.F5 = false
	_this.F8 = _dafny.EmptyMap
	_this.F15 = _dafny.CodePoint('D')
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f6 = false
	_this._f7 = _dafny.Zero
	_this._f9 = _dafny.EmptyMap
	_this._f10 = _dafny.EmptySeq
	_this._f11 = false
	_this._f12 = _dafny.EmptySeq
	_this._f13 = _dafny.EmptySeq
	_this._f14 = false
	return &_this
}

type CompanionStruct_GlobalState_ struct {
}

var Companion_GlobalState_ = CompanionStruct_GlobalState_{}

func (_this *GlobalState) Equals(other *GlobalState) bool {
	return _this == other
}

func (_this *GlobalState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*GlobalState)
	return ok && _this.Equals(other)
}

func (*GlobalState) String() string {
	return "_module.GlobalState"
}

func Type_GlobalState_() _dafny.TypeDescriptor {
	return type_GlobalState_{}
}

type type_GlobalState_ struct {
}

func (_this type_GlobalState_) Default() interface{} {
	return (*GlobalState)(nil)
}

func (_this type_GlobalState_) String() string {
	return "main.GlobalState"
}
func (_this *GlobalState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &GlobalState{}

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 _dafny.Array, f2 bool, f3 _dafny.Int, f4 _dafny.Array, f5 bool, f6 bool, f7 _dafny.Int, f8 _dafny.Map, f9 _dafny.Map, f10 _dafny.Sequence, f11 bool, f12 _dafny.Sequence, f13 _dafny.Sequence, f14 bool, f15 _dafny.CodePoint) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Array {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() bool {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F7() _dafny.Int {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F9() _dafny.Map {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Sequence {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Sequence {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F13() _dafny.Sequence {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F27  _dafny.Array
	_f26 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F27 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f26 = _dafny.Zero
	return &_this
}

type CompanionStruct_C0_ struct {
}

var Companion_C0_ = CompanionStruct_C0_{}

func (_this *C0) Equals(other *C0) bool {
	return _this == other
}

func (_this *C0) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C0)
	return ok && _this.Equals(other)
}

func (*C0) String() string {
	return "_module.C0"
}

func Type_C0_() _dafny.TypeDescriptor {
	return type_C0_{}
}

type type_C0_ struct {
}

func (_this type_C0_) Default() interface{} {
	return (*C0)(nil)
}

func (_this type_C0_) String() string {
	return "main.C0"
}
func (_this *C0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f26 _dafny.Int, f27 _dafny.Array) {
	{
		(_this)._f26 = f26
		(_this).F27 = f27
	}
}
func (_this *C0) F26() _dafny.Int {
	{
		return _this._f26
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f16 bool
	_f17 _dafny.Sequence
	_f28 _dafny.Int
	_f29 _dafny.Int
	_f30 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	_this._f28 = _dafny.Zero
	_this._f29 = _dafny.Zero
	_this._f30 = _dafny.Zero
	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C1{}
var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F16() bool {
	return _this._f16
}
func (_this *C1) F16_set_(value bool) {
	_this._f16 = value
}
func (_this *C1) F17() _dafny.Sequence {
	return _this._f17
}
func (_this *C1) F28() _dafny.Int {
	return _this._f28
}
func (_this *C1) Ctor__(f29 _dafny.Int, f30 _dafny.Int, f16 bool, f17 _dafny.Sequence, f28 _dafny.Int) {
	{
		(_this)._f29 = f29
		(_this)._f30 = f30
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f28 = f28
	}
}
func (_this *C1) Fm13(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(519)
	}
}
func (_this *C1) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt(((_this).F29()).Minus((_this).F29()), (func() _dafny.Set {
			var _coll32 = _dafny.NewBuilder()
			_ = _coll32
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(847), _dafny.IntOfInt64(900))); ; {
				_compr_32, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _349_v0 _dafny.Int
				_349_v0 = interface{}(_compr_32).(_dafny.Int)
				if ((_dafny.IntOfInt64(847)).Cmp(_349_v0) <= 0) && ((_349_v0).Cmp(_dafny.IntOfInt64(900)) < 0) {
					_coll32.Add((_349_v0).Plus((_this).F29()))
				}
			}
			return _coll32.ToSet()
		}()).Cardinality())
	}
}
func (_this *C1) Fm15(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_D4_.Create_DC12_(_this.F16(), _this.F16(), (_this).F28())).Dtor_cf23()
	}
}
func (_this *C1) M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _350_v0 _dafny.Int
		_ = _350_v0
		_350_v0 = _dafny.IntOfInt64(138)
		var _rhs77 bool = ((func() bool {
			if _this.F16() {
				return p1
			}
			return _this.F16()
		})()) || (p2)
		_ = _rhs77
		var _rhs78 _dafny.Int = Companion_Default___.Fm3(globalState)
		_ = _rhs78
		var _lhs75 *C1 = _this
		_ = _lhs75
		_lhs75.F16_set_(_rhs77)
		_350_v0 = _rhs78
		var _351_v1 _dafny.Map
		_ = _351_v1
		_351_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F29()).Minus(_dafny.IntOfInt64(253)), p1)
		_351_v1 = (_351_v1).Update((_this).F30(), _this.F16())
		var _352_v2 _dafny.Array
		_ = _352_v2
		var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw67
		_352_v2 = _nw67
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))
		_ = _index77
		(_352_v2).ArraySet1(_this.F16(), (_index77).Int())
		var _353_v3 _dafny.CodePoint
		_ = _353_v3
		_353_v3 = _dafny.CodePoint('r')
		var _354_v4 _dafny.Sequence
		_ = _354_v4
		_354_v4 = _dafny.SeqOf(((_dafny.Zero).Minus((_this).Fm14((_this).F28(), _dafny.IntOfUint32((p0).Cardinality()), _353_v3, globalState))).Minus(_350_v0))
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))
		_ = _index78
		var _rhs79 bool = !(p1)
		_ = _rhs79
		var _rhs80 bool = _this.F16()
		_ = _rhs80
		var _rhs81 bool = (_350_v0).Cmp((_this).F29()) != 0
		_ = _rhs81
		var _rhs82 _dafny.Sequence = (_this).F17()
		_ = _rhs82
		var _lhs76 *GlobalState = globalState
		_ = _lhs76
		var _lhs77 *GlobalState = globalState
		_ = _lhs77
		var _lhs78 _dafny.Array = _352_v2
		_ = _lhs78
		var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))
		_ = _lhs79
		_lhs76.F2 = _rhs79
		_lhs77.F5 = _rhs80
		(_lhs78).ArraySet1(_rhs81, (_lhs79).Int())
		_354_v4 = _rhs82
		var _355_v5 *C0
		_ = _355_v5
		var _nw68 *C0 = New_C0_()
		_ = _nw68
		_nw68.Ctor__((Companion_Default___.Fm16((_this).F30(), globalState)).Cardinality(), _352_v2)
		_355_v5 = _nw68
		var _356_v6 D0
		_ = _356_v6
		_356_v6 = Companion_D0_.Create_DC0_(p0)
		var _pat_let_tv11 = globalState
		_ = _pat_let_tv11
		var _source3 D0 = func(_pat_let8_0 D0) D0 {
			return func(_357_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let9_0 _dafny.Sequence) D0 {
					return func(_358_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_358_dt__update_hcf0_h0)
					}(_pat_let9_0)
				}(Companion_Default___.Fm1(_pat_let_tv11))
			}(_pat_let8_0)
		}(_356_v6)
		_ = _source3
		if _source3.Is_DC1() {
			var _359___mcc_h0 _dafny.Sequence = _source3.Get_().(D0_DC1).Cf1
			_ = _359___mcc_h0
			var _360_cf1 _dafny.Sequence = _359___mcc_h0
			_ = _360_cf1
			var _361_v7 _dafny.Sequence
			_ = _361_v7
			var _362_v8 bool
			_ = _362_v8
			var _363_v9 bool
			_ = _363_v9
			var _out0 _dafny.Sequence
			_ = _out0
			var _out1 bool
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = (_this).M12(globalState)
			_361_v7 = _out0
			_362_v8 = _out1
			_363_v9 = _out2
			var _364_v10 _dafny.Set
			_ = _364_v10
			_364_v10 = _dafny.SetOf(_dafny.IntOfUint32((_360_cf1).Cardinality()))
			var _365_v11 *C0
			_ = _365_v11
			var _nw69 *C0 = New_C0_()
			_ = _nw69
			_nw69.Ctor__(Companion_Default___.SafeDivisionInt(_350_v0, (_364_v10).Cardinality()), _352_v2)
			_365_v11 = _nw69
			(globalState).F2 = _363_v9
			_350_v0 = Companion_Default___.SafeModuloInt(((_355_v5).F26()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mtmijpy")).Cardinality())), p3)
		} else if _source3.Is_DC2() {
			var _366___mcc_h1 _dafny.Int = _source3.Get_().(D0_DC2).Cf2
			_ = _366___mcc_h1
			var _367___mcc_h2 _dafny.Int = _source3.Get_().(D0_DC2).Cf3
			_ = _367___mcc_h2
			var _368___mcc_h3 _dafny.Int = _source3.Get_().(D0_DC2).Cf4
			_ = _368___mcc_h3
			var _369___mcc_h4 _dafny.Array = _source3.Get_().(D0_DC2).Cf5
			_ = _369___mcc_h4
			var _370_cf5 _dafny.Array = _369___mcc_h4
			_ = _370_cf5
			var _371_cf4 _dafny.Int = _368___mcc_h3
			_ = _371_cf4
			var _372_cf3 _dafny.Int = _367___mcc_h2
			_ = _372_cf3
			var _373_cf2 _dafny.Int = _366___mcc_h1
			_ = _373_cf2
			var _374_v12 _dafny.Array
			_ = _374_v12
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw70
			_374_v12 = _nw70
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_374_v12), 0))
			_ = _index79
			(_374_v12).ArraySet1(_350_v0, (_index79).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(871), _dafny.ArrayLen((_374_v12), 0))
			_ = _index80
			(_374_v12).ArraySet1((_dafny.Zero).Minus((_this).F28()), (_index80).Int())
			var _375_v13 _dafny.Set
			_ = _375_v13
			_375_v13 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(691))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_376_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_377_i0 _dafny.Int) _dafny.CodePoint {
					return _376_v3
				}
			})(_353_v3))), p0)
			(globalState).F5 = Companion_Default___.Fm0((_375_v13).Equals(_dafny.SetOf(p0, p0)), p2, globalState)
			var _378_v14 *C0
			_ = _378_v14
			var _nw71 *C0 = New_C0_()
			_ = _nw71
			_nw71.Ctor__((_355_v5).F26(), (func() _dafny.Array {
				if p2 {
					return _370_cf5
				}
				return _355_v5.F27
			})())
			_378_v14 = _nw71
			_371_cf4 = (_378_v14).F26()
		} else {
			var _379___mcc_h5 _dafny.Sequence = _source3.Get_().(D0_DC0).Cf0
			_ = _379___mcc_h5
			var _380_cf0 _dafny.Sequence = _379___mcc_h5
			_ = _380_cf0
			var _381_v15 _dafny.Map
			_ = _381_v15
			_381_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_355_v5).F26(), _352_v2)
			var _382_v16 _dafny.Map
			_ = _382_v16
			_382_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_350_v0, (func() _dafny.Array {
				if (_381_v15).Contains((_this).F30()) {
					return (_381_v15).Get((_this).F30()).(_dafny.Array)
				}
				return _355_v5.F27
			})())
			_382_v16 = (_382_v16).Update(_dafny.IntOfInt64(297), _355_v5.F27)
			var _383_v17 _dafny.Map
			_ = _383_v17
			_383_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_352_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))).Int()).(bool))
			var _384_v18 D1
			_ = _384_v18
			_384_v18 = Companion_D1_.Create_DC4_(p2, (_355_v5).F26(), (_383_v17).Merge(_383_v17))
			var _385_v19 _dafny.Map
			_ = _385_v19
			_385_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_v3, _dafny.IntOfUint32(((_this).F17()).Cardinality()))
			var _386_v20 _dafny.Map
			_ = _386_v20
			_386_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _385_v19)
			_384_v18 = Companion_D1_.Create_DC4_(false, (_386_v20).Cardinality(), _383_v17)
			if (false) && (((_355_v5).F26()).Cmp((_355_v5).F26()) == 0) {
				_350_v0 = Companion_Default___.SafeDivisionInt((_this).F28(), (func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter35 := _dafny.Iterate(((_this).F17()).Elements()); ; {
						_compr_33, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _387_v21 _dafny.Int
						_387_v21 = interface{}(_compr_33).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains((_this).F17(), _387_v21) {
							_coll33.Add(Companion_Default___.SafeModuloInt(_387_v21, _350_v0), (_this).F28())
						}
					}
					return _coll33.ToMap()
				}()).Cardinality())
				var _388_v22 *C0
				_ = _388_v22
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__(_350_v0, _352_v2)
				_388_v22 = _nw72
				var _389_v23 _dafny.Map
				_ = _389_v23
				_389_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_352_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))).Int()).(bool), _dafny.IntOfInt64(443))
				(globalState).F5 = !(_389_v23).Contains(p2)
				var _390_v24 _dafny.Array
				_ = _390_v24
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(9))
				_ = _nw73
				_390_v24 = _nw73
				var _391_v25 _dafny.Array
				_ = _391_v25
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_9
				var _nw74 _dafny.Array
				_ = _nw74
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw74 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Sequence = (func(_392_cf0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_393_i1 _dafny.Int) _dafny.Sequence {
							return _392_cf0
						}
					})(_380_cf0)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw74 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw74).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw74).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_391_v25 = _nw74
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_391_v25), 0))
				_ = _index81
				(_391_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_380_cf0, p0), (_index81).Int())
				var _394_v26 _dafny.Array
				_ = _394_v26
				_394_v26 = _390_v24
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_391_v25), 0))
				_ = _index82
				var _rhs83 bool = (func() bool {
					if _this.F16() {
						return p1
					}
					return _this.F16()
				})()
				_ = _rhs83
				var _rhs84 _dafny.Array = (_394_v26)
				_ = _rhs84
				var _rhs85 _dafny.CodePoint = _353_v3
				_ = _rhs85
				var _rhs86 _dafny.Sequence = _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm17(globalState), _354_v4)).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _353_v3)
				_ = _rhs86
				var _rhs87 bool = p1
				_ = _rhs87
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				var _lhs81 *GlobalState = globalState
				_ = _lhs81
				var _lhs82 _dafny.Array = _391_v25
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_391_v25), 0))
				_ = _lhs83
				var _lhs84 *C1 = _this
				_ = _lhs84
				_lhs80.F5 = _rhs83
				_390_v24 = _rhs84
				_lhs81.F15 = _rhs85
				(_lhs82).ArraySet1(_rhs86, (_lhs83).Int())
				_lhs84.F16_set_(_rhs87)
				var _395_v27 _dafny.Sequence
				_ = _395_v27
				_395_v27 = _dafny.SeqOf(_351_v1, _351_v1, _351_v1, _351_v1)
				var _396_v28 _dafny.Array
				_ = _396_v28
				var _nwElement0_13 _dafny.Map = (_395_v27).Select((Companion_Default___.SafeIndex((_388_v22).F26(), _dafny.IntOfUint32((_395_v27).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _nwElement0_13
				var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
				_ = _nw75
				(_nw75).ArraySet1(_nwElement0_13, 0)
				_396_v28 = _nw75
				var _rhs88 bool = (Companion_Default___.SafeModuloInt(p3, _dafny.IntOfInt64(151))).Cmp((_this).F30()) < 0
				_ = _rhs88
				var _rhs89 bool = ((_this).Fm14(p3, (_355_v5).F26(), _353_v3, globalState)).Cmp(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_383_v17).Cardinality(), (_352_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_352_v2), 0))).Int()).(bool))).Merge(func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter36 := _dafny.Iterate(((_this).F17()).Elements()); ; {
						_compr_34, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _397_v29 _dafny.Int
						_397_v29 = interface{}(_compr_34).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains((_this).F17(), _397_v29) {
							_coll34.Add((_397_v29).Plus((_355_v5).F26()), p1)
						}
					}
					return _coll34.ToMap()
				}())).Cardinality()) > 0
				_ = _rhs89
				var _rhs90 _dafny.Array = _396_v28
				_ = _rhs90
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				_lhs85.F2 = _rhs88
				_lhs86.F5 = _rhs89
				_396_v28 = _rhs90
			} else {
				(globalState).F2 = (((_355_v5).F26()).Times((_this).F29())).Cmp(((_this).F17()).Select((Companion_Default___.SafeIndex(_350_v0, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int)) > 0
				(globalState).F15 = _353_v3
				var _398_v30 *C0
				_ = _398_v30
				var _nw76 *C0 = New_C0_()
				_ = _nw76
				_nw76.Ctor__((_this).Fm13(p1, globalState), _355_v5.F27)
				_398_v30 = _nw76
				var _399_v31 _dafny.MultiSet
				_ = _399_v31
				_399_v31 = _dafny.MultiSetOf((_398_v30).F26())
				_399_v31 = (_dafny.MultiSetFromSeq((_this).F17())).Intersection(_399_v31)
				(globalState).F5 = _dafny.Companion_Sequence_.IsPrefixOf((_this).F17(), _354_v4)
			}
			var _400_v32 _dafny.Array
			_ = _400_v32
			var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw77
			_400_v32 = _nw77
			var _401_v33 _dafny.MultiSet
			_ = _401_v33
			_401_v33 = _dafny.MultiSetOf(_this.F16(), Companion_Default___.Fm0(!(p1), Companion_Default___.Fm0(p2, _this.F16(), globalState), globalState), true)
			var _402_v34 _dafny.Map
			_ = _402_v34
			_402_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_400_v32, (_dafny.Zero).Minus((_401_v33).Cardinality()))
			_402_v34 = _402_v34
		}
		var _403_v35 _dafny.Array
		_ = _403_v35
		var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
		_ = _nw78
		_403_v35 = _nw78
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_403_v35), 0))
		_ = _index83
		(_403_v35).ArraySet1(_352_v2, (_index83).Int())
		var _404_v36 _dafny.Sequence
		_ = _404_v36
		_404_v36 = _dafny.SeqOf(_this.F16())
		var _pat_let_tv12 = _355_v5
		_ = _pat_let_tv12
		var _pat_let_tv13 = _355_v5
		_ = _pat_let_tv13
		var _pat_let_tv14 = _404_v36
		_ = _pat_let_tv14
		var _pat_let_tv15 = p2
		_ = _pat_let_tv15
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_403_v35), 0))
		_ = _index84
		var _rhs91 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_404_v36, _404_v36), _404_v36)
		_ = _rhs91
		var _rhs92 _dafny.Int = func(_source4 D4) _dafny.Int {
			if _source4.Is_DC12() {
				var _405___mcc_h6 bool = _source4.Get_().(D4_DC12).Cf21
				_ = _405___mcc_h6
				var _406___mcc_h7 bool = _source4.Get_().(D4_DC12).Cf22
				_ = _406___mcc_h7
				var _407___mcc_h8 _dafny.Int = _source4.Get_().(D4_DC12).Cf23
				_ = _407___mcc_h8
				var _408_cf23 _dafny.Int = _407___mcc_h8
				_ = _408_cf23
				var _409_cf22 bool = _406___mcc_h7
				_ = _409_cf22
				var _410_cf21 bool = _405___mcc_h6
				_ = _410_cf21
				return (_pat_let_tv12).F26()
			} else if _source4.Is_DC13() {
				var _411___mcc_h9 _dafny.Int = _source4.Get_().(D4_DC13).Cf24
				_ = _411___mcc_h9
				var _412___mcc_h10 _dafny.Int = _source4.Get_().(D4_DC13).Cf25
				_ = _412___mcc_h10
				var _413___mcc_h11 _dafny.Int = _source4.Get_().(D4_DC13).Cf26
				_ = _413___mcc_h11
				var _414_cf26 _dafny.Int = _413___mcc_h11
				_ = _414_cf26
				var _415_cf25 _dafny.Int = _412___mcc_h10
				_ = _415_cf25
				var _416_cf24 _dafny.Int = _411___mcc_h9
				_ = _416_cf24
				return Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv13).F26(), _this.F16())).Cardinality(), _dafny.IntOfUint32((_pat_let_tv14).Cardinality()))
			} else {
				var _417___mcc_h12 _dafny.MultiSet = _source4.Get_().(D4_DC11).Cf20
				_ = _417___mcc_h12
				var _418_cf20 _dafny.MultiSet = _417___mcc_h12
				_ = _418_cf20
				return (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.SetOf(_pat_let_tv15))).Cardinality())
			}
		}(Companion_D4_.Create_DC13_((_this).F28(), (_this).F30(), (_this).F30()))
		_ = _rhs92
		var _rhs93 _dafny.Int = (_this).F29()
		_ = _rhs93
		var _rhs94 _dafny.Int = p3
		_ = _rhs94
		var _rhs95 _dafny.Array = _355_v5.F27
		_ = _rhs95
		var _lhs87 _dafny.Array = _403_v35
		_ = _lhs87
		var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_403_v35), 0))
		_ = _lhs88
		r0 = _rhs91
		_350_v0 = _rhs92
		_350_v0 = _rhs93
		_350_v0 = _rhs94
		(_lhs87).ArraySet1(_rhs95, (_lhs88).Int())
		r0 = _dafny.SeqOf(p2)
		return r0
	}
}
func (_this *C1) M11(globalState *GlobalState) (_dafny.CodePoint, bool, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _hi2 _dafny.Int = ((_this).F29()).Plus((_this).F30())
		_ = _hi2
		for _419_i0 := (_this).F28(); _419_i0.Cmp(_hi2) < 0; _419_i0 = _419_i0.Plus(_dafny.One) {
			r2 = Companion_Default___.Fm3(globalState)
			r0 = Companion_Default___.Fm2(globalState)
			var _420_v0 _dafny.Array
			_ = _420_v0
			var _nw79 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
			_ = _nw79
			_420_v0 = _nw79
			var _421_v1 *C0
			_ = _421_v1
			var _nw80 *C0 = New_C0_()
			_ = _nw80
			_nw80.Ctor__(_419_i0, _420_v0)
			_421_v1 = _nw80
			var _422_v2 _dafny.Sequence
			_ = _422_v2
			_422_v2 = _dafny.UnicodeSeqOfUtf8Bytes("anddf")
			var _423_v3 _dafny.MultiSet
			_ = _423_v3
			_423_v3 = _dafny.MultiSetOf(_420_v0)
			var _424_v4 _dafny.Map
			_ = _424_v4
			_424_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_422_v2, (_423_v3).Difference(_423_v3))
			_424_v4 = (_424_v4).Update((func() _dafny.Sequence {
				if _this.F16() {
					return _422_v2
				}
				return _422_v2
			})(), _423_v3)
		}
		var _hi3 _dafny.Int = _dafny.IntOfUint32(((_this).F17()).Cardinality())
		_ = _hi3
		for _425_i1 := (_this).F28(); _425_i1.Cmp(_hi3) < 0; _425_i1 = _425_i1.Plus(_dafny.One) {
			var _426_v5 _dafny.Array
			_ = _426_v5
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw81
			_426_v5 = _nw81
			var _427_v6 _dafny.Map
			_ = _427_v6
			_427_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F30())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
			_ = _index85
			(_426_v5).ArraySet1((func() _dafny.Int {
				if (_427_v6).Contains((_this).F17()) {
					return (_427_v6).Get((_this).F17()).(_dafny.Int)
				}
				return _425_i1
			})(), (_index85).Int())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_426_v5), 0))
			_ = _index86
			(_426_v5).ArraySet1(_dafny.IntOfInt64(759), (_index86).Int())
			var _428_v7 _dafny.Map
			_ = _428_v7
			_428_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _dafny.IntOfInt64(13))
			var _429_v8 _dafny.Sequence
			_ = _429_v8
			_429_v8 = _dafny.UnicodeSeqOfUtf8Bytes("lw")
			var _430_v9 _dafny.MultiSet
			_ = _430_v9
			_430_v9 = _dafny.MultiSetOf(_428_v7, _428_v7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), Companion_Default___.Fm3(globalState))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _dafny.IntOfUint32((_429_v8).Cardinality()))))
			var _431_v10 _dafny.Set
			_ = _431_v10
			_431_v10 = _dafny.SetOf(_this.F16())
			var _432_v11 _dafny.MultiSet
			_ = _432_v11
			_432_v11 = _dafny.MultiSetOf((_this).F30(), _dafny.IntOfInt64(393))
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
			_ = _index87
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_426_v5), 0))
			_ = _index88
			var _rhs96 _dafny.Int = (_430_v9).Cardinality()
			_ = _rhs96
			var _rhs97 _dafny.Int = _dafny.IntOfUint32((_429_v8).Cardinality())
			_ = _rhs97
			var _rhs98 bool = (!((_431_v10).IsSubsetOf(_431_v10))) && (!(_dafny.Companion_Sequence_.Equal((_this).F17(), (_this).F17())))
			_ = _rhs98
			var _rhs99 bool = _this.F16()
			_ = _rhs99
			var _rhs100 bool = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _432_v11)).Contains(!(_this.F16()) || (Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState)))
			_ = _rhs100
			var _lhs89 _dafny.Array = _426_v5
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
			_ = _lhs90
			var _lhs91 _dafny.Array = _426_v5
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_426_v5), 0))
			_ = _lhs92
			var _lhs93 *GlobalState = globalState
			_ = _lhs93
			var _lhs94 *GlobalState = globalState
			_ = _lhs94
			var _lhs95 *GlobalState = globalState
			_ = _lhs95
			(_lhs89).ArraySet1(_rhs96, (_lhs90).Int())
			(_lhs91).ArraySet1(_rhs97, (_lhs92).Int())
			_lhs93.F2 = _rhs98
			_lhs94.F5 = _rhs99
			_lhs95.F5 = _rhs100
			var _433_v12 D1
			_ = _433_v12
			_433_v12 = Companion_D1_.Create_DC4_(_this.F16(), (_dafny.IntOfInt64(43)).Plus((_this).F29()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16()))
			var _source5 D1 = _433_v12
			_ = _source5
			if _source5.Is_DC4() {
				var _434___mcc_h0 bool = _source5.Get_().(D1_DC4).Cf7
				_ = _434___mcc_h0
				var _435___mcc_h1 _dafny.Int = _source5.Get_().(D1_DC4).Cf8
				_ = _435___mcc_h1
				var _436___mcc_h2 _dafny.Map = _source5.Get_().(D1_DC4).Cf9
				_ = _436___mcc_h2
				var _437_cf9 _dafny.Map = _436___mcc_h2
				_ = _437_cf9
				var _438_cf8 _dafny.Int = _435___mcc_h1
				_ = _438_cf8
				var _439_cf7 bool = _434___mcc_h0
				_ = _439_cf7
				var _440_v13 _dafny.Sequence
				_ = _440_v13
				_440_v13 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("hndbi"))
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
				_ = _index89
				(_426_v5).ArraySet1((_425_i1).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_440_v13).Cardinality()), (_this).F30())), (_index89).Int())
				var _441_v14 _dafny.Sequence
				_ = _441_v14
				_441_v14 = _dafny.SeqOf(false)
				_438_cf8 = (_dafny.IntOfUint32((_441_v14).Cardinality())).Plus((_this).F29())
				var _442_v15 _dafny.Sequence
				_ = _442_v15
				_442_v15 = _dafny.SeqOf(_426_v5)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
				_ = _index90
				(_426_v5).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_442_v15, _442_v15)).Cardinality()), (_index90).Int())
				var _443_v16 _dafny.Array
				_ = _443_v16
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_10
				var _nw82 _dafny.Array
				_ = _nw82
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw82 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = func(_444_i2 _dafny.Int) bool {
						return false
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw82 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw82).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw82).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_443_v16 = _nw82
				var _445_v17 *C0
				_ = _445_v17
				var _nw83 *C0 = New_C0_()
				_ = _nw83
				_nw83.Ctor__((_426_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))).Int()).(_dafny.Int), _443_v16)
				_445_v17 = _nw83
			} else if _source5.Is_DC5() {
				var _446___mcc_h3 _dafny.Int = _source5.Get_().(D1_DC5).Cf10
				_ = _446___mcc_h3
				var _447___mcc_h4 bool = _source5.Get_().(D1_DC5).Cf11
				_ = _447___mcc_h4
				var _448_cf11 bool = _447___mcc_h4
				_ = _448_cf11
				var _449_cf10 _dafny.Int = _446___mcc_h3
				_ = _449_cf10
				var _450_v18 _dafny.MultiSet
				_ = _450_v18
				_450_v18 = _dafny.MultiSetOf(_448_cf11, !(_this.F16()))
				var _451_v19 D4
				_ = _451_v19
				_451_v19 = Companion_D4_.Create_DC12_(_this.F16(), (func() bool {
					if _448_cf11 {
						return _this.F16()
					}
					return _448_cf11
				})(), (_450_v18).Cardinality())
				_451_v19 = _451_v19
				var _452_v20 _dafny.Array
				_ = _452_v20
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw84
				_452_v20 = _nw84
				_452_v20 = _452_v20
				(_this).F16_set_(_this.F16())
				var _453_v21 _dafny.Array
				_ = _453_v21
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_11
				var _nw85 _dafny.Array
				_ = _nw85
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw85 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = (func(_454_cf11 bool) func(_dafny.Int) bool {
						return func(_455_i3 _dafny.Int) bool {
							return _454_cf11
						}
					})(_448_cf11)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw85 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw85).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw85).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_453_v21 = _nw85
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_453_v21), 0))
				_ = _index91
				(_453_v21).ArraySet1(true, (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_453_v21), 0))
				_ = _index92
				(_453_v21).ArraySet1(_this.F16(), (_index92).Int())
			} else if _source5.Is_DC6() {
				var _456_v22 _dafny.Array
				_ = _456_v22
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw86
				_456_v22 = _nw86
				var _457_v23 *C0
				_ = _457_v23
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__((_this).F29(), _456_v22)
				_457_v23 = _nw87
				var _458_v24 _dafny.MultiSet
				_ = _458_v24
				_458_v24 = _dafny.MultiSetOf(_this.F16(), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_this).F30()), _dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), _425_i1)), _this.F16(), _this.F16(), true)
				var _459_v25 _dafny.Set
				_ = _459_v25
				_459_v25 = _dafny.SetOf((_this).F29(), ((_428_v7).Merge(_428_v7)).Cardinality())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
				_ = _index93
				var _rhs101 _dafny.Int = (_dafny.SetOf(_429_v8, _dafny.Companion_Sequence_.Concatenate(_429_v8, Companion_Default___.Fm1(globalState)), _429_v8, _dafny.Companion_Sequence_.Concatenate(_429_v8, _429_v8))).Cardinality()
				_ = _rhs101
				var _rhs102 _dafny.MultiSet = (_458_v24).Difference(_dafny.MultiSetOf(_this.F16()))
				_ = _rhs102
				var _rhs103 _dafny.Set = (_459_v25).Intersection(_459_v25)
				_ = _rhs103
				var _lhs96 _dafny.Array = _426_v5
				_ = _lhs96
				var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
				_ = _lhs97
				(_lhs96).ArraySet1(_rhs101, (_lhs97).Int())
				_458_v24 = _rhs102
				_459_v25 = _rhs103
				var _460_v26 _dafny.Map
				_ = _460_v26
				_460_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_426_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))).Int()).(_dafny.Int), _this.F16())
				var _461_v27 *C0
				_ = _461_v27
				var _nw88 *C0 = New_C0_()
				_ = _nw88
				_nw88.Ctor__(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_460_v26).Cardinality())), (_this).F29()), _457_v23.F27)
				_461_v27 = _nw88
				var _462_v28 _dafny.CodePoint
				_ = _462_v28
				_462_v28 = _dafny.CodePoint('y')
				(globalState).F15 = _462_v28
			} else {
				var _463___mcc_h5 _dafny.MultiSet = _source5.Get_().(D1_DC3).Cf6
				_ = _463___mcc_h5
				var _464_cf6 _dafny.MultiSet = _463___mcc_h5
				_ = _464_cf6
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))
				_ = _index94
				(_426_v5).ArraySet1(_dafny.IntOfUint32((_429_v8).Cardinality()), (_index94).Int())
				(globalState).F5 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), _426_v5)).Cardinality()).Cmp((_this).F30()) < 0
				var _465_v29 _dafny.Sequence
				_ = _465_v29
				var _466_v30 bool
				_ = _466_v30
				var _467_v31 bool
				_ = _467_v31
				var _out3 _dafny.Sequence
				_ = _out3
				var _out4 bool
				_ = _out4
				var _out5 bool
				_ = _out5
				_out3, _out4, _out5 = (_this).M12(globalState)
				_465_v29 = _out3
				_466_v30 = _out4
				_467_v31 = _out5
				var _468_v32 _dafny.Map
				_ = _468_v32
				_468_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_464_cf6), _dafny.Companion_Sequence_.Concatenate(_465_v29, _465_v29))
				var _469_v33 D1
				_ = _469_v33
				_469_v33 = Companion_D1_.Create_DC3_(_464_cf6)
				var _470_v34 _dafny.CodePoint
				_ = _470_v34
				_470_v34 = _dafny.CodePoint('q')
				var _rhs104 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_468_v32).Contains(_469_v33) {
						return (_468_v32).Get(_469_v33).(_dafny.Sequence)
					}
					return _429_v8
				})(), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_468_v32).Contains(_469_v33) {
						return (_468_v32).Get(_469_v33).(_dafny.Sequence)
					}
					return _429_v8
				})()).Cardinality()))).Uint32(), _470_v34), (Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_468_v32).Contains(_469_v33) {
						return (_468_v32).Get(_469_v33).(_dafny.Sequence)
					}
					return _429_v8
				})(), (Companion_Default___.SafeIndex((_this).F30(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_468_v32).Contains(_469_v33) {
						return (_468_v32).Get(_469_v33).(_dafny.Sequence)
					}
					return _429_v8
				})()).Cardinality()))).Uint32(), _470_v34)).Cardinality()))).Uint32(), _dafny.CodePoint('i'))
				_ = _rhs104
				var _rhs105 bool = _this.F16()
				_ = _rhs105
				var _rhs106 bool = Companion_Default___.Fm0(_467_v31, _466_v30, globalState)
				_ = _rhs106
				var _rhs107 bool = _466_v30
				_ = _rhs107
				var _rhs108 bool = _467_v31
				_ = _rhs108
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				_465_v29 = _rhs104
				_lhs98.F2 = _rhs105
				r1 = _rhs106
				r1 = _rhs107
				_lhs99.F5 = _rhs108
			}
			var _471_v36 _dafny.Map
			_ = _471_v36
			_471_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _425_i1)
			var _472_v38 _dafny.Sequence
			_ = _472_v38
			_472_v38 = _dafny.SeqOf(_this.F16())
			var _473_v39 _dafny.Sequence
			_ = _473_v39
			_473_v39 = _dafny.SeqOf(_472_v38)
			var _474_v40 _dafny.Sequence
			_ = _474_v40
			_474_v40 = _dafny.SeqOf((_471_v36).Update(_dafny.IntOfUint32(((_473_v39).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_472_v38).Cardinality()), _dafny.IntOfUint32((_473_v39).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_this).F30()))
			var _475_v42 _dafny.Array
			_ = _475_v42
			var _nwElement0_14 _dafny.Map = func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter37 := _dafny.Iterate((_dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), (_this).F29())).Elements()); ; {
					_compr_35, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _476_v35 _dafny.Int
					_476_v35 = interface{}(_compr_35).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex((_this).F29(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), (_this).F29()), _476_v35) {
						_coll35.Add(Companion_Default___.SafeModuloInt(_476_v35, _425_i1), (_426_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))).Int()).(_dafny.Int))
					}
				}
				return _coll35.ToMap()
			}()
			_ = _nwElement0_14
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(12))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_14, 0)
			(_nw89).ArraySet1((_471_v36).Merge(_471_v36), 1)
			(_nw89).ArraySet1((_471_v36).Merge(_471_v36), 2)
			(_nw89).ArraySet1(func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter38 := _dafny.Iterate(((_this).F17()).Elements()); ; {
					_compr_36, _ok38 := _iter38()
					if !_ok38 {
						break
					}
					var _477_v37 _dafny.Int
					_477_v37 = interface{}(_compr_36).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((_this).F17(), _477_v37) {
						_coll36.Add(Companion_Default___.SafeModuloInt(_477_v37, (_426_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_426_v5), 0))).Int()).(_dafny.Int)), (_dafny.Zero).Minus((_this).F30()))
					}
				}
				return _coll36.ToMap()
			}(), 3)
			(_nw89).ArraySet1(_471_v36, 4)
			(_nw89).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_i1, (_this).F30()), 5)
			(_nw89).ArraySet1(_471_v36, 6)
			(_nw89).ArraySet1((_474_v40).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_428_v7).Contains(_this.F16()) {
					return (_428_v7).Get(_this.F16()).(_dafny.Int)
				}
				return _425_i1
			})(), _dafny.IntOfUint32((_474_v40).Cardinality()))).Uint32()).(_dafny.Map), 7)
			(_nw89).ArraySet1(func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter39 := _dafny.Iterate((_432_v11).Elements()); ; {
					_compr_37, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _478_v41 _dafny.Int
					_478_v41 = interface{}(_compr_37).(_dafny.Int)
					if (_432_v11).Contains(_478_v41) {
						_coll37.Add((_478_v41).Plus((_this).F29()), _dafny.IntOfUint32((_472_v38).Cardinality()))
					}
				}
				return _coll37.ToMap()
			}(), 8)
			(_nw89).ArraySet1(_471_v36, 9)
			(_nw89).ArraySet1(_471_v36, 10)
			(_nw89).ArraySet1((_471_v36).Update((_this).F29(), (_this).F29()), 11)
			_475_v42 = _nw89
			_475_v42 = _475_v42
			r2 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(60), (_425_i1).Times(_425_i1))
		}
		var _479_v43 _dafny.Array
		_ = _479_v43
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw90
		_479_v43 = _nw90
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
		_ = _index95
		(_479_v43).ArraySet1((_this).F29(), (_index95).Int())
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
		_ = _index96
		var _rhs109 _dafny.Int = (_this).F30()
		_ = _rhs109
		var _rhs110 _dafny.Int = (_this).F28()
		_ = _rhs110
		var _lhs100 _dafny.Array = _479_v43
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
		_ = _lhs101
		(_lhs100).ArraySet1(_rhs109, (_lhs101).Int())
		r2 = _rhs110
		var _480_v44 _dafny.Array
		_ = _480_v44
		var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw91
		_480_v44 = _nw91
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_480_v44), 0))); ; {
			_guard_loop_2, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _481_i4 _dafny.Int
			_481_i4 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_481_i4).Sign() != -1) && ((_481_i4).Cmp(_dafny.ArrayLen((_480_v44), 0)) < 0)) {
				(_480_v44).ArraySet1(_this.F16(), (_481_i4).Int())
			}
		}
		var _482_v45 D4
		_ = _482_v45
		_482_v45 = Companion_D4_.Create_DC13_((_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F16()), _480_v44)).Cardinality(), (_this).F28()), _dafny.IntOfInt64(-224))
		var _source6 D4 = _482_v45
		_ = _source6
		if _source6.Is_DC12() {
			var _483___mcc_h6 bool = _source6.Get_().(D4_DC12).Cf21
			_ = _483___mcc_h6
			var _484___mcc_h7 bool = _source6.Get_().(D4_DC12).Cf22
			_ = _484___mcc_h7
			var _485___mcc_h8 _dafny.Int = _source6.Get_().(D4_DC12).Cf23
			_ = _485___mcc_h8
			var _486_cf23 _dafny.Int = _485___mcc_h8
			_ = _486_cf23
			var _487_cf22 bool = _484___mcc_h7
			_ = _487_cf22
			var _488_cf21 bool = _483___mcc_h6
			_ = _488_cf21
			var _489_v46 _dafny.Sequence
			_ = _489_v46
			_489_v46 = _dafny.UnicodeSeqOfUtf8Bytes("wivplsm")
			var _490_v47 _dafny.Sequence
			_ = _490_v47
			_490_v47 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wugmj"), _dafny.UnicodeSeqOfUtf8Bytes("cgl"), _dafny.UnicodeSeqOfUtf8Bytes("kamk"), _489_v46, _489_v46)
			var _491_v48 _dafny.MultiSet
			_ = _491_v48
			_491_v48 = _dafny.MultiSetOf(_this.F16(), false)
			var _492_v49 _dafny.Map
			_ = _492_v49
			_492_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_490_v47, (func() _dafny.MultiSet {
				if _this.F16() {
					return _dafny.MultiSetOf(!(false), _488_cf21, _this.F16())
				}
				return _491_v48
			})())
			_492_v49 = (_492_v49).Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18((_dafny.Zero).Minus(_486_cf23), globalState), _490_v47), _491_v48)
			(_this).F16_set_(_this.F16())
			if _488_cf21 {
				var _493_v50 _dafny.Set
				_ = _493_v50
				_493_v50 = _dafny.SetOf(((_this).F28()).Plus(_dafny.IntOfInt64(748)), (_this).F28(), _dafny.IntOfInt64(811))
				var _494_v51 _dafny.Array
				_ = _494_v51
				var _nw92 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
				_ = _nw92
				_494_v51 = _nw92
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_494_v51), 0))
				_ = _index97
				(_494_v51).ArraySet1(_dafny.SeqOf((_this).Fm13(_488_cf21, globalState)), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_494_v51), 0))
				_ = _index98
				var _rhs111 _dafny.Set = _493_v50
				_ = _rhs111
				var _rhs112 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17())
				_ = _rhs112
				var _lhs102 _dafny.Array = _494_v51
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_494_v51), 0))
				_ = _lhs103
				_493_v50 = _rhs111
				(_lhs102).ArraySet1(_rhs112, (_lhs103).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_480_v44), 0))
				_ = _index99
				(_480_v44).ArraySet1(_this.F16(), (_index99).Int())
				var _495_v52 _dafny.Array
				_ = _495_v52
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_12
				var _nw93 _dafny.Array
				_ = _nw93
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw93 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.CodePoint = (func(_496_cf22 bool) func(_dafny.Int) _dafny.CodePoint {
						return func(_497_i5 _dafny.Int) _dafny.CodePoint {
							return (func() _dafny.CodePoint {
								if _496_cf22 {
									return _dafny.CodePoint('f')
								}
								return _dafny.CodePoint('d')
							})()
						}
					})(_487_cf22)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw93 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw93).ArraySet1CodePoint(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw93).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_495_v52 = _nw93
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_495_v52), 0))
				_ = _index100
				(_495_v52).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index100).Int())
				var _498_v53 _dafny.CodePoint
				_ = _498_v53
				_498_v53 = _dafny.CodePoint('q')
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_480_v44), 0))
				_ = _index101
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_495_v52), 0))
				_ = _index102
				var _rhs113 bool = ((_this).F29()).Cmp((_this).F29()) != 0
				_ = _rhs113
				var _rhs114 bool = !(((_this).F29()).Cmp(Companion_Default___.SafeDivisionInt((_this).F29(), (_this).F29())) < 0)
				_ = _rhs114
				var _rhs115 _dafny.CodePoint = _498_v53
				_ = _rhs115
				var _lhs104 _dafny.Array = _480_v44
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_480_v44), 0))
				_ = _lhs105
				var _lhs106 _dafny.Array = _495_v52
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_495_v52), 0))
				_ = _lhs107
				(_lhs104).ArraySet1(_rhs113, (_lhs105).Int())
				r1 = _rhs114
				(_lhs106).ArraySet1CodePoint(_rhs115, (_lhs107).Int())
				_498_v53 = _dafny.CodePoint('i')
				_493_v50 = _493_v50
				r1 = ((_dafny.Zero).Minus(_486_cf23)).Cmp(((_494_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_494_v51), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(786), _dafny.IntOfUint32(((_494_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_494_v51), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int)) != 0
			} else {
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
				_ = _index103
				(_479_v43).ArraySet1((_this).F29(), (_index103).Int())
				var _499_v54 _dafny.Map
				_ = _499_v54
				_499_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_this).F29())
				var _500_v55 _dafny.Array
				_ = _500_v55
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_13
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = func(_501_i6 _dafny.Int) _dafny.Sequence {
						return (_this).F17()
					}
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw94 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw94).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw94).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_500_v55 = _nw94
				var _502_v56 _dafny.Sequence
				_ = _502_v56
				_502_v56 = _dafny.SeqOf(true)
				var _503_v57 _dafny.Map
				_ = _503_v57
				_503_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_488_cf21, Companion_D2_.Create_DC8_(_499_v54, _500_v55, _487_cf22, _487_cf22, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_502_v56).Cardinality()), true)).Cardinality()))
				var _504_v58 D2
				_ = _504_v58
				_504_v58 = Companion_D2_.Create_DC8_(_499_v54, _500_v55, _488_cf21, Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState), _486_cf23)
				_503_v57 = (_503_v57).Update(Companion_Default___.Fm0(_this.F16(), _487_cf22, globalState), _504_v58)
				_489_v46 = Companion_Default___.Fm1(globalState)
				r2 = (_this).Fm13(_487_cf22, globalState)
				(globalState).F5 = _487_cf22
			}
			if _488_cf21 {
				var _505_v59 _dafny.Array
				_ = _505_v59
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(22))
				_ = _nw95
				_505_v59 = _nw95
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_505_v59), 0))
				_ = _index104
				(_505_v59).ArraySet1(_dafny.SetOf(_489_v46), (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(76), _dafny.ArrayLen((_505_v59), 0))
				_ = _index105
				(_505_v59).ArraySet1(_dafny.SetOf(_489_v46), (_index105).Int())
				_487_cf22 = _488_cf21
				var _506_v60 _dafny.CodePoint
				_ = _506_v60
				_506_v60 = _dafny.CodePoint('p')
				var _507_v61 _dafny.MultiSet
				_ = _507_v61
				_507_v61 = _dafny.MultiSetOf(_506_v60, _dafny.CodePoint('s'))
				var _508_v62 _dafny.Sequence
				_ = _508_v62
				_508_v62 = _dafny.SeqOf(_507_v61, (_507_v61).Update(_506_v60, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sey")).Cardinality()))), _dafny.MultiSetOf(_506_v60, _506_v60, _506_v60, _506_v60, _506_v60), _507_v61)
				var _509_v63 _dafny.Map
				_ = _509_v63
				_509_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_508_v62).Cardinality()), _486_cf23)
				var _rhs116 _dafny.Map = _509_v63
				_ = _rhs116
				var _rhs117 bool = _487_cf22
				_ = _rhs117
				var _lhs108 *C1 = _this
				_ = _lhs108
				_509_v63 = _rhs116
				_lhs108.F16_set_(_rhs117)
				var _510_v64 D0
				_ = _510_v64
				_510_v64 = Companion_D0_.Create_DC0_(_489_v46)
				_510_v64 = _510_v64
				_509_v63 = (_509_v63).Update((_this).F28(), (_this).F29())
			} else {
				var _511_v65 D1
				_ = _511_v65
				_511_v65 = Companion_D1_.Create_DC3_((_491_v48).Difference(_491_v48))
				_511_v65 = _511_v65
				var _512_v66 *C0
				_ = _512_v66
				var _nw96 *C0 = New_C0_()
				_ = _nw96
				_nw96.Ctor__((_this).F29(), (Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_489_v46).Cardinality()), (_this).F28(), (_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int), _480_v44)).Dtor_cf5())
				_512_v66 = _nw96
				var _513_v67 _dafny.Set
				_ = _513_v67
				_513_v67 = _dafny.SetOf(_488_cf21, _487_cf22)
				var _514_v68 D2
				_ = _514_v68
				_514_v68 = Companion_D2_.Create_DC7_(_513_v67)
				var _515_v69 _dafny.Array
				_ = _515_v69
				var _nwElement0_15 D2 = _514_v68
				_ = _nwElement0_15
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(14))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_15, 0)
				(_nw97).ArraySet1(_514_v68, 1)
				(_nw97).ArraySet1(_514_v68, 2)
				(_nw97).ArraySet1(Companion_D2_.Create_DC7_(_513_v67), 3)
				(_nw97).ArraySet1(_514_v68, 4)
				(_nw97).ArraySet1(_514_v68, 5)
				(_nw97).ArraySet1(_514_v68, 6)
				(_nw97).ArraySet1(_514_v68, 7)
				(_nw97).ArraySet1(_514_v68, 8)
				(_nw97).ArraySet1(_514_v68, 9)
				(_nw97).ArraySet1(_514_v68, 10)
				(_nw97).ArraySet1(Companion_D2_.Create_DC7_(_513_v67), 11)
				(_nw97).ArraySet1(_514_v68, 12)
				(_nw97).ArraySet1(_514_v68, 13)
				_515_v69 = _nw97
				var _516_v70 _dafny.CodePoint
				_ = _516_v70
				_516_v70 = _dafny.CodePoint('a')
				var _517_v71 _dafny.Map
				_ = _517_v71
				_517_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_515_v69, Companion_Default___.SafeDivisionInt((_512_v66).F26(), (_this).Fm14((_512_v66).F26(), (_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int), _516_v70, globalState)))
				_517_v71 = (_517_v71).Update(_515_v69, _dafny.IntOfUint32(((_this).F17()).Cardinality()))
				_488_cf21 = _487_cf22
				var _518_v72 _dafny.Array
				_ = _518_v72
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
				_ = _nw98
				_518_v72 = _nw98
				var _519_v73 _dafny.Map
				_ = _519_v73
				_519_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _488_cf21)
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_518_v72), 0))
				_ = _index106
				(_518_v72).ArraySet1(_519_v73, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_518_v72), 0))
				_ = _index107
				(_518_v72).ArraySet1(_519_v73, (_index107).Int())
			}
		} else if _source6.Is_DC13() {
			var _520___mcc_h9 _dafny.Int = _source6.Get_().(D4_DC13).Cf24
			_ = _520___mcc_h9
			var _521___mcc_h10 _dafny.Int = _source6.Get_().(D4_DC13).Cf25
			_ = _521___mcc_h10
			var _522___mcc_h11 _dafny.Int = _source6.Get_().(D4_DC13).Cf26
			_ = _522___mcc_h11
			var _523_cf26 _dafny.Int = _522___mcc_h11
			_ = _523_cf26
			var _524_cf25 _dafny.Int = _521___mcc_h10
			_ = _524_cf25
			var _525_cf24 _dafny.Int = _520___mcc_h9
			_ = _525_cf24
			var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
			_ = _index108
			(_479_v43).ArraySet1((_this).F28(), (_index108).Int())
			var _526_v74 _dafny.Sequence
			_ = _526_v74
			_526_v74 = _dafny.SeqOf(_this.F16())
			_526_v74 = _dafny.Companion_Sequence_.Concatenate(_526_v74, _526_v74)
			var _527_v75 _dafny.Set
			_ = _527_v75
			_527_v75 = _dafny.SetOf(_479_v43)
			_527_v75 = ((_527_v75).Intersection(_527_v75)).Difference(_527_v75)
			var _528_v76 _dafny.Sequence
			_ = _528_v76
			_528_v76 = _dafny.SeqOf((_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int), (_527_v75).Cardinality(), (_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int))
			var _529_v77 _dafny.Sequence
			_ = _529_v77
			_529_v77 = _dafny.UnicodeSeqOfUtf8Bytes("js")
			var _530_v78 _dafny.MultiSet
			_ = _530_v78
			_530_v78 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ixeqyo"), _529_v77), _this.F16(), (_this.F16()) && (_this.F16()), _this.F16())
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
			_ = _index109
			var _rhs118 _dafny.Sequence = _526_v74
			_ = _rhs118
			var _rhs119 _dafny.Sequence = _528_v76
			_ = _rhs119
			var _rhs120 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_526_v74)).Union(_530_v78)
			_ = _rhs120
			var _rhs121 _dafny.Int = (_this).F30()
			_ = _rhs121
			var _lhs109 _dafny.Array = _479_v43
			_ = _lhs109
			var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))
			_ = _lhs110
			_526_v74 = _rhs118
			_528_v76 = _rhs119
			_530_v78 = _rhs120
			(_lhs109).ArraySet1(_rhs121, (_lhs110).Int())
		} else {
			var _531___mcc_h12 _dafny.MultiSet = _source6.Get_().(D4_DC11).Cf20
			_ = _531___mcc_h12
			var _532_cf20 _dafny.MultiSet = _531___mcc_h12
			_ = _532_cf20
			var _533_v79 _dafny.Array
			_ = _533_v79
			var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
			_ = _nw99
			_533_v79 = _nw99
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_533_v79), 0))
			_ = _index110
			(_533_v79).ArraySet1((_this).F17(), (_index110).Int())
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_533_v79), 0))
			_ = _index111
			(_533_v79).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17()), (_index111).Int())
			var _534_v80 _dafny.Sequence
			_ = _534_v80
			_534_v80 = _dafny.UnicodeSeqOfUtf8Bytes("lju")
			_534_v80 = _534_v80
			var _535_v81 D0
			_ = _535_v81
			_535_v81 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_534_v80, _534_v80))
			_535_v81 = _535_v81
			var _536_v82 _dafny.Set
			_ = _536_v82
			_536_v82 = _dafny.SetOf(false, Companion_Default___.Fm0(!(false), _this.F16(), globalState), true, _this.F16(), ((_this).F28()).Cmp((_this).F28()) != 0)
			_536_v82 = _536_v82
		}
		r1 = (func() bool {
			if _this.F16() {
				return _this.F16()
			}
			return _this.F16()
		})()
		var _537_v83 _dafny.CodePoint
		_ = _537_v83
		_537_v83 = _dafny.CodePoint('c')
		r0 = _537_v83
		r1 = _this.F16()
		r2 = ((_479_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_479_v43), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-271)))
		return r0, r1, r2
	}
}
func (_this *C1) M12(globalState *GlobalState) (_dafny.Sequence, bool, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _538_v0 _dafny.Sequence
		_ = _538_v0
		_538_v0 = _dafny.SeqOf(_this.F16(), _this.F16())
		var _539_v1 _dafny.Array
		_ = _539_v1
		var _nwElement0_16 bool = false
		_ = _nwElement0_16
		var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(5))
		_ = _nw100
		(_nw100).ArraySet1(_nwElement0_16, 0)
		(_nw100).ArraySet1((_dafny.MultiSetOf(_this.F16())).IsProperSubsetOf(_dafny.MultiSetFromSeq(_538_v0)), 1)
		(_nw100).ArraySet1((func() bool {
			if _this.F16() {
				return _this.F16()
			}
			return !(true)
		})(), 2)
		(_nw100).ArraySet1(((_this).F29()).Cmp((_this).F29()) <= 0, 3)
		(_nw100).ArraySet1(!((_dafny.IntOfUint32((_dafny.SeqOf((_this).F30())).Cardinality())).Cmp(_dafny.IntOfUint32(((_this).F17()).Cardinality())) >= 0), 4)
		_539_v1 = _nw100
		for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_539_v1), 0))); ; {
			_guard_loop_3, _ok41 := _iter41()
			if !_ok41 {
				break
			}
			var _540_i0 _dafny.Int
			_540_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_540_i0).Sign() != -1) && ((_540_i0).Cmp(_dafny.ArrayLen((_539_v1), 0)) < 0)) {
				(_539_v1).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ucfop"), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F29(), _this.F16()), _dafny.SetOf((_dafny.MultiSetOf(_this.F16())).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ucfop")).Cardinality()))).Uint32(), _dafny.CodePoint('b')), _dafny.UnicodeSeqOfUtf8Bytes("lpiib")), (_540_i0).Int())
			}
		}
		var _hi4 _dafny.Int = (_this).F30()
		_ = _hi4
		for _541_i1 := ((_this).F28()).Times(_dafny.IntOfInt64(566)); _541_i1.Cmp(_hi4) < 0; _541_i1 = _541_i1.Plus(_dafny.One) {
			if Companion_Default___.Fm0(_dafny.Companion_Sequence_.IsProperPrefixOf(_538_v0, _538_v0), _this.F16(), globalState) {
				var _542_v2 _dafny.Int
				_ = _542_v2
				_542_v2 = _dafny.IntOfInt64(460)
				_542_v2 = (_this).F30()
				var _543_v3 _dafny.Sequence
				_ = _543_v3
				_543_v3 = _dafny.UnicodeSeqOfUtf8Bytes("g")
				var _544_v4 _dafny.Map
				_ = _544_v4
				_544_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_543_v3).Cardinality()), false)
				var _545_v5 _dafny.CodePoint
				_ = _545_v5
				_545_v5 = _dafny.CodePoint('p')
				var _546_v6 _dafny.Map
				_ = _546_v6
				_546_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v4, _545_v5)
				_546_v6 = (_546_v6).Update(_544_v4, _545_v5)
				(globalState).F5 = false
				_542_v2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F28(), (_this).F28()))
				var _547_v7 D1
				_ = _547_v7
				_547_v7 = Companion_D1_.Create_DC4_(_this.F16(), _542_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16()))
				var _rhs122 _dafny.Int = (_547_v7).Dtor_cf8()
				_ = _rhs122
				var _rhs123 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("rdpvkliok")
				_ = _rhs123
				var _rhs124 _dafny.Int = _542_v2
				_ = _rhs124
				_542_v2 = _rhs122
				_543_v3 = _rhs123
				_542_v2 = _rhs124
			} else {
				var _548_v8 _dafny.Map
				_ = _548_v8
				_548_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F28())))
				var _549_v9 _dafny.Array
				_ = _549_v9
				var _len0_14 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_14
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_14.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_14)
				} else {
					var _init14 func(_dafny.Int) _dafny.Sequence = func(_550_i2 _dafny.Int) _dafny.Sequence {
						return (_this).F17()
					}
					_ = _init14
					var _element0_14 = _init14(_dafny.Zero)
					_ = _element0_14
					_nw101 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
					(_nw101).ArraySet1(_element0_14, 0)
					var _nativeLen0_14 = (_len0_14).Int()
					_ = _nativeLen0_14
					for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
						(_nw101).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
					}
				}
				_549_v9 = _nw101
				var _pat_let_tv16 = _549_v9
				_ = _pat_let_tv16
				_548_v8 = ((_548_v8).Merge(_548_v8)).Merge((func(_pat_let10_0 D2) D2 {
					return func(_551_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let11_0 _dafny.Array) D2 {
							return func(_552_dt__update_hcf14_h0 _dafny.Array) D2 {
								return func(_pat_let12_0 bool) D2 {
									return func(_553_dt__update_hcf16_h0 bool) D2 {
										return Companion_D2_.Create_DC8_((_551_dt__update__tmp_h0).Dtor_cf13(), _552_dt__update_hcf14_h0, (_551_dt__update__tmp_h0).Dtor_cf15(), _553_dt__update_hcf16_h0, (_551_dt__update__tmp_h0).Dtor_cf17())
									}(_pat_let12_0)
								}(true)
							}(_pat_let11_0)
						}(_pat_let_tv16)
					}(_pat_let10_0)
				}(Companion_D2_.Create_DC8_(_548_v8, _549_v9, _this.F16(), _this.F16(), (_this).F29()))).Dtor_cf13())
				var _554_v10 _dafny.Int
				_ = _554_v10
				_554_v10 = _dafny.IntOfInt64(704)
				var _555_v11 _dafny.CodePoint
				_ = _555_v11
				_555_v11 = _dafny.CodePoint('k')
				var _rhs125 _dafny.Int = _541_i1
				_ = _rhs125
				var _rhs126 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ufg"), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_541_i1, (_this).F30()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ufg")).Cardinality()))).Uint32(), _555_v11)
				_ = _rhs126
				var _rhs127 bool = (func() bool {
					if Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState) {
						return _this.F16()
					}
					return true
				})()
				_ = _rhs127
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				_554_v10 = _rhs125
				r0 = _rhs126
				_lhs111.F2 = _rhs127
				_554_v10 = ((_this).F28()).Minus((_this).F28())
				var _556_v12 _dafny.Array
				_ = _556_v12
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_15
				var _nw102 _dafny.Array
				_ = _nw102
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw102 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) _dafny.CodePoint = (func(_557_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_558_i3 _dafny.Int) _dafny.CodePoint {
							return _557_v11
						}
					})(_555_v11)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw102 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw102).ArraySet1CodePoint(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw102).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_556_v12 = _nw102
				var _559_v13 _dafny.Map
				_ = _559_v13
				_559_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _556_v12)
				var _560_v14 _dafny.Sequence
				_ = _560_v14
				_560_v14 = _dafny.UnicodeSeqOfUtf8Bytes("fj")
				var _561_v15 _dafny.MultiSet
				_ = _561_v15
				_561_v15 = _dafny.MultiSetOf(_554_v10, (_this).F30(), _dafny.IntOfUint32((_560_v14).Cardinality()))
				var _562_v16 _dafny.Array
				_ = _562_v16
				var _nwElement0_17 _dafny.Int = Companion_Default___.SafeModuloInt((_this).Fm14((_this).F29(), (_this).F29(), _dafny.CodePoint('h'), globalState), _dafny.IntOfInt64(-34))
				_ = _nwElement0_17
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(17))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_17, 0)
				(_nw103).ArraySet1((_this).F30(), 1)
				(_nw103).ArraySet1(Companion_Default___.Fm3(globalState), 2)
				(_nw103).ArraySet1(_541_i1, 3)
				(_nw103).ArraySet1(_554_v10, 4)
				(_nw103).ArraySet1(((_this).F29()).Times((_548_v8).Cardinality()), 5)
				(_nw103).ArraySet1((_this).F28(), 6)
				(_nw103).ArraySet1(_dafny.IntOfInt64(84), 7)
				(_nw103).ArraySet1((_this).F30(), 8)
				(_nw103).ArraySet1((_this).F30(), 9)
				(_nw103).ArraySet1((_this).F29(), 10)
				(_nw103).ArraySet1((_this).F28(), 11)
				(_nw103).ArraySet1((_559_v13).Cardinality(), 12)
				(_nw103).ArraySet1(((_561_v15).Difference(_561_v15)).Cardinality(), 13)
				(_nw103).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_554_v10), _dafny.IntOfInt64(210)), 14)
				(_nw103).ArraySet1(_dafny.IntOfUint32((_560_v14).Cardinality()), 15)
				(_nw103).ArraySet1((_this).F30(), 16)
				_562_v16 = _nw103
				var _rhs128 _dafny.Array = _562_v16
				_ = _rhs128
				var _rhs129 bool = _this.F16()
				_ = _rhs129
				var _rhs130 bool = _this.F16()
				_ = _rhs130
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				var _lhs113 *GlobalState = globalState
				_ = _lhs113
				_562_v16 = _rhs128
				_lhs112.F2 = _rhs129
				_lhs113.F5 = _rhs130
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_539_v1), 0))
				_ = _index112
				(_539_v1).ArraySet1(Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState), (_index112).Int())
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_539_v1), 0))
				_ = _index113
				var _rhs131 bool = !(_this.F16())
				_ = _rhs131
				var _rhs132 bool = _this.F16()
				_ = _rhs132
				var _rhs133 bool = Companion_Default___.Fm0(((_this).F28()).Cmp((_this).F30()) != 0, (func() bool {
					if _this.F16() {
						return _this.F16()
					}
					return _this.F16()
				})(), globalState)
				_ = _rhs133
				var _rhs134 bool = (_538_v0).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if _this.F16() {
						return Companion_Default___.Fm3(globalState)
					}
					return (_this).F28()
				})(), _dafny.IntOfUint32((_538_v0).Cardinality()))).Uint32()).(bool)
				_ = _rhs134
				var _lhs114 _dafny.Array = _539_v1
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_539_v1), 0))
				_ = _lhs115
				var _lhs116 *GlobalState = globalState
				_ = _lhs116
				r2 = _rhs131
				(_lhs114).ArraySet1(_rhs132, (_lhs115).Int())
				_lhs116.F5 = _rhs133
				r2 = _rhs134
			}
			(globalState).F2 = _this.F16()
			var _563_v17 *C0
			_ = _563_v17
			var _nw104 *C0 = New_C0_()
			_ = _nw104
			_nw104.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _this.F16())).Cardinality()).Times((_this).F28()), _539_v1)
			_563_v17 = _nw104
			var _564_v18 _dafny.Map
			_ = _564_v18
			_564_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_539_v1, _539_v1)
			var _565_v19 _dafny.Array
			_ = _565_v19
			var _nwElement0_18 _dafny.Array = _563_v17.F27
			_ = _nwElement0_18
			var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(12))
			_ = _nw105
			(_nw105).ArraySet1(_nwElement0_18, 0)
			(_nw105).ArraySet1(_539_v1, 1)
			(_nw105).ArraySet1(_563_v17.F27, 2)
			(_nw105).ArraySet1(_539_v1, 3)
			(_nw105).ArraySet1(_539_v1, 4)
			(_nw105).ArraySet1(_563_v17.F27, 5)
			(_nw105).ArraySet1((func() _dafny.Array {
				if _this.F16() {
					return _539_v1
				}
				return _563_v17.F27
			})(), 6)
			(_nw105).ArraySet1(_539_v1, 7)
			(_nw105).ArraySet1(_563_v17.F27, 8)
			(_nw105).ArraySet1((func() _dafny.Array {
				if (_564_v18).Contains(_563_v17.F27) {
					return (_564_v18).Get(_563_v17.F27).(_dafny.Array)
				}
				return _563_v17.F27
			})(), 9)
			(_nw105).ArraySet1(_539_v1, 10)
			(_nw105).ArraySet1(_539_v1, 11)
			_565_v19 = _nw105
			var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw106
			_565_v19 = _nw106
		}
		var _566_v20 _dafny.Int
		_ = _566_v20
		_566_v20 = _dafny.IntOfInt64(354)
		_566_v20 = (_this).F28()
		var _567_v21 _dafny.Set
		_ = _567_v21
		_567_v21 = _dafny.SetOf(_539_v1)
		var _568_v22 _dafny.Set
		_ = _568_v22
		_568_v22 = _dafny.SetOf(_this.F16(), _this.F16())
		var _569_v23 _dafny.Map
		_ = _569_v23
		_569_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _568_v22)
		r2 = (func() bool {
			if (_567_v21).IsProperSubsetOf(_567_v21) {
				return _this.F16()
			}
			return (_569_v23).Equals(_569_v23)
		})()
		_566_v20 = _dafny.IntOfInt64(497)
		var _570_v24 _dafny.Sequence
		_ = _570_v24
		_570_v24 = _dafny.UnicodeSeqOfUtf8Bytes("ldcs")
		var _571_v25 _dafny.Map
		_ = _571_v25
		_571_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_570_v24).Cardinality()), _this.F16())
		(globalState).F8 = (_571_v25).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), !(_this.F16())))
		r0 = Companion_Default___.Fm1(globalState)
		var _572_v26 _dafny.MultiSet
		_ = _572_v26
		_572_v26 = _dafny.MultiSetOf((_this).F28())
		r1 = ((_572_v26).Difference(Companion_Default___.Fm19((_this).F28(), globalState))).IsDisjointFrom(_572_v26)
		r2 = _this.F16()
		return r0, r1, r2
	}
}
func (_this *C1) M13(p0 _dafny.Array, p1 bool, p2 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _573_v0 _dafny.Array
		_ = _573_v0
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_16
		var _nw107 _dafny.Array
		_ = _nw107
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw107 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) bool = (func(_574_p2 bool) func(_dafny.Int) bool {
				return func(_575_i0 _dafny.Int) bool {
					return _574_p2
				}
			})(p2)
			_ = _init16
			var _element0_16 = _init16(_dafny.Zero)
			_ = _element0_16
			_nw107 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
			(_nw107).ArraySet1(_element0_16, 0)
			var _nativeLen0_16 = (_len0_16).Int()
			_ = _nativeLen0_16
			for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
				(_nw107).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
			}
		}
		_573_v0 = _nw107
		var _576_v1 *C0
		_ = _576_v1
		var _nw108 *C0 = New_C0_()
		_ = _nw108
		_nw108.Ctor__(_dafny.IntOfInt64(876), _573_v0)
		_576_v1 = _nw108
		var _577_v2 _dafny.Array
		_ = _577_v2
		var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw109
		_577_v2 = _nw109
		_577_v2 = _577_v2
		var _578_i1 _dafny.Int
		_ = _578_i1
		_578_i1 = _dafny.Zero
		{
			for (func() bool {
				if _this.F16() {
					return !(p2)
				}
				return ((_this).F28()).Cmp((_this).F30()) <= 0
			})() {
				{
					if (_578_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_578_i1 = (_578_i1).Plus(_dafny.One)
					r0 = (_this).F28()
					r0 = _dafny.IntOfInt64(965)
					(globalState).F2 = (func() bool {
						if _this.F16() {
							return p2
						}
						return _this.F16()
					})()
					var _579_v3 _dafny.Set
					_ = _579_v3
					_579_v3 = _dafny.SetOf(p1, _this.F16(), false, p2, _this.F16())
					var _580_v4 D2
					_ = _580_v4
					_580_v4 = Companion_D2_.Create_DC7_((_dafny.SetOf(_this.F16())).Union(_579_v3))
					_580_v4 = _580_v4
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _581_v5 _dafny.Set
		_ = _581_v5
		_581_v5 = _dafny.SetOf(p1, p2)
		var _582_v6 _dafny.Sequence
		_ = _582_v6
		_582_v6 = _dafny.UnicodeSeqOfUtf8Bytes("curdxgmyi")
		var _583_v7 _dafny.Map
		_ = _583_v7
		_583_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_581_v5).Cardinality(), _dafny.IntOfUint32((_582_v6).Cardinality()))
		var _584_v8 _dafny.Sequence
		_ = _584_v8
		_584_v8 = _dafny.SeqOf(_583_v7, _583_v7)
		var _585_v10 _dafny.CodePoint
		_ = _585_v10
		_585_v10 = _dafny.CodePoint('p')
		var _586_v11 _dafny.Map
		_ = _586_v11
		_586_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v10, p1)
		var _587_v13 _dafny.MultiSet
		_ = _587_v13
		_587_v13 = _dafny.MultiSetOf((_584_v8).Select((Companion_Default___.SafeIndex((_this).F28(), _dafny.IntOfUint32((_584_v8).Cardinality()))).Uint32()).(_dafny.Map), (func() _dafny.Map {
			var _coll38 = _dafny.NewMapBuilder()
			_ = _coll38
			for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(357), _dafny.IntOfInt64(697))); ; {
				_compr_38, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _588_v9 _dafny.Int
				_588_v9 = interface{}(_compr_38).(_dafny.Int)
				if ((_dafny.IntOfInt64(357)).Cmp(_588_v9) <= 0) && ((_588_v9).Cmp(_dafny.IntOfInt64(697)) < 0) {
					_coll38.Add((_588_v9).Plus(_dafny.IntOfInt64(196)), (_576_v1).F26())
				}
			}
			return _coll38.ToMap()
		}()).Update((_this).F30(), (_586_v11).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F28(), (_576_v1).F26()), func() _dafny.Map {
			var _coll39 = _dafny.NewMapBuilder()
			_ = _coll39
			for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(983), _dafny.IntOfInt64(580))); ; {
				_compr_39, _ok43 := _iter43()
				if !_ok43 {
					break
				}
				var _589_v12 _dafny.Int
				_589_v12 = interface{}(_compr_39).(_dafny.Int)
				if ((_dafny.IntOfInt64(983)).Cmp(_589_v12) <= 0) && ((_589_v12).Cmp(_dafny.IntOfInt64(580)) < 0) {
					_coll39.Add((_589_v12).Times(_dafny.IntOfInt64(323)), (_576_v1).F26())
				}
			}
			return _coll39.ToMap()
		}(), _583_v7)
		(globalState).F5 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_584_v8, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F30(), _dafny.IntOfInt64(-324)), _583_v7, _583_v7)))).IsDisjointFrom(_587_v13)
		var _590_v15 _dafny.Array
		_ = _590_v15
		var _nwElement0_19 _dafny.Int = (_this).F28()
		_ = _nwElement0_19
		var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(17))
		_ = _nw110
		(_nw110).ArraySet1(_nwElement0_19, 0)
		(_nw110).ArraySet1((_this).F30(), 1)
		(_nw110).ArraySet1((_this).F30(), 2)
		(_nw110).ArraySet1(Companion_Default___.SafeModuloInt((_576_v1).F26(), (_this).F28()), 3)
		(_nw110).ArraySet1((_576_v1).F26(), 4)
		(_nw110).ArraySet1((_576_v1).F26(), 5)
		(_nw110).ArraySet1(((_this).F28()).Plus((_dafny.Zero).Minus((_576_v1).F26())), 6)
		(_nw110).ArraySet1((func() _dafny.Int {
			if p1 {
				return (_576_v1).F26()
			}
			return _dafny.IntOfInt64(700)
		})(), 7)
		(_nw110).ArraySet1((_this).F30(), 8)
		(_nw110).ArraySet1(_dafny.IntOfInt64(-614), 9)
		(_nw110).ArraySet1((_this).F28(), 10)
		(_nw110).ArraySet1((_this).Fm13(_this.F16(), globalState), 11)
		(_nw110).ArraySet1((_this).F30(), 12)
		(_nw110).ArraySet1((_this).F28(), 13)
		(_nw110).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Set {
			var _coll40 = _dafny.NewBuilder()
			_ = _coll40
			for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(965), _dafny.IntOfInt64(-1))); ; {
				_compr_40, _ok44 := _iter44()
				if !_ok44 {
					break
				}
				var _591_v14 _dafny.Int
				_591_v14 = interface{}(_compr_40).(_dafny.Int)
				if ((_dafny.IntOfInt64(965)).Cmp(_591_v14) <= 0) && ((_591_v14).Cmp(_dafny.IntOfInt64(-1)) < 0) {
					_coll40.Add((_591_v14).Minus((_576_v1).F26()))
				}
			}
			return _coll40.ToSet()
		}()).Cardinality(), _dafny.IntOfInt64(40)), 14)
		(_nw110).ArraySet1((_this).F28(), 15)
		(_nw110).ArraySet1((_this).F28(), 16)
		_590_v15 = _nw110
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_590_v15), 0))
		_ = _index114
		(_590_v15).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((_582_v6).Cardinality()))).Plus((_this).F29()), (_index114).Int())
		var _592_v16 _dafny.Array
		_ = _592_v16
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_17
		var _nw111 _dafny.Array
		_ = _nw111
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw111 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.MultiSet = (func(_593_v1 *C0) func(_dafny.Int) _dafny.MultiSet {
				return func(_594_i2 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf(_dafny.MultiSetOf((_593_v1).F26()))
				}
			})(_576_v1)
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw111 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw111).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw111).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_592_v16 = _nw111
		var _595_v17 _dafny.MultiSet
		_ = _595_v17
		_595_v17 = _dafny.MultiSetOf((_this).F29())
		var _596_v18 _dafny.MultiSet
		_ = _596_v18
		_596_v18 = _dafny.MultiSetOf(_595_v17, _595_v17)
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_592_v16), 0))
		_ = _index115
		(_592_v16).ArraySet1((_596_v18).Union(_596_v18), (_index115).Int())
		var _597_v19 _dafny.Array
		_ = _597_v19
		var _nw112 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw112
		_597_v19 = _nw112
		var _598_v20 _dafny.Array
		_ = _598_v20
		var _nwElement0_20 _dafny.CodePoint = _585_v10
		_ = _nwElement0_20
		var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.One)
		_ = _nw113
		(_nw113).ArraySet1CodePoint(_nwElement0_20, 0)
		_598_v20 = _nw113
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_597_v19), 0))
		_ = _index116
		(_597_v19).ArraySet1(_598_v20, (_index116).Int())
		var _599_v21 _dafny.Sequence
		_ = _599_v21
		_599_v21 = _dafny.SeqOf(_595_v17, _595_v17)
		var _600_v22 _dafny.Sequence
		_ = _600_v22
		_600_v22 = _dafny.SeqOf(p2)
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_590_v15), 0))
		_ = _index117
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_592_v16), 0))
		_ = _index118
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_597_v19), 0))
		_ = _index119
		var _rhs135 _dafny.Int = _dafny.IntOfInt64(-533)
		_ = _rhs135
		var _rhs136 _dafny.Int = (_this).F29()
		_ = _rhs136
		var _rhs137 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_599_v21, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(463), (_dafny.Zero).Minus((_576_v1).F26())))))).Difference(_596_v18)
		_ = _rhs137
		var _rhs138 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_600_v22, Companion_Default___.Fm20(Companion_Default___.Fm0(p1, p1, globalState), _this.F16(), p1, Companion_Default___.Fm3(globalState), globalState))).Cardinality())
		_ = _rhs138
		var _rhs139 _dafny.Array = (func() _dafny.Array {
			if p2 {
				return _598_v20
			}
			return _598_v20
		})()
		_ = _rhs139
		var _lhs117 _dafny.Array = _590_v15
		_ = _lhs117
		var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_590_v15), 0))
		_ = _lhs118
		var _lhs119 _dafny.Array = _592_v16
		_ = _lhs119
		var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_592_v16), 0))
		_ = _lhs120
		var _lhs121 _dafny.Array = _597_v19
		_ = _lhs121
		var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_597_v19), 0))
		_ = _lhs122
		(_lhs117).ArraySet1(_rhs135, (_lhs118).Int())
		r1 = _rhs136
		(_lhs119).ArraySet1(_rhs137, (_lhs120).Int())
		r0 = _rhs138
		(_lhs121).ArraySet1(_rhs139, (_lhs122).Int())
		var _601_v23 _dafny.Map
		_ = _601_v23
		_601_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(globalState), p1)
		var _602_v24 *C0
		_ = _602_v24
		var _nw114 *C0 = New_C0_()
		_ = _nw114
		_nw114.Ctor__((func() _dafny.Int {
			if p1 {
				return (_601_v23).Cardinality()
			}
			return (_576_v1).F26()
		})(), _576_v1.F27)
		_602_v24 = _nw114
		r0 = (func() _dafny.Int {
			if !(p1) {
				return ((_dafny.Zero).Minus((_this).F28())).Plus((_this).F28())
			}
			return _dafny.IntOfUint32(((_this).F17()).Cardinality())
		})()
		var _603_v25 _dafny.MultiSet
		_ = _603_v25
		_603_v25 = _dafny.MultiSetOf(_600_v22)
		var _604_v26 _dafny.MultiSet
		_ = _604_v26
		_604_v26 = _dafny.MultiSetOf((_this).F17())
		var _605_v27 _dafny.Sequence
		_ = _605_v27
		_605_v27 = _dafny.SeqOf((_this).F17())
		r1 = ((_603_v25).Update(_600_v22, Companion_Default___.Abs((func() _dafny.Int {
			if (_604_v26).Contains((_605_v27).Select((Companion_Default___.SafeIndex(((_this).F17()).Select((Companion_Default___.SafeIndex((_576_v1).F26(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_605_v27).Cardinality()))).Uint32()).(_dafny.Sequence)) {
				return (_604_v26).Multiplicity((_605_v27).Select((Companion_Default___.SafeIndex(((_this).F17()).Select((Companion_Default___.SafeIndex((_576_v1).F26(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_605_v27).Cardinality()))).Uint32()).(_dafny.Sequence))
			}
			return (_this).Fm15((_this).F17(), _dafny.IntOfUint32((_600_v22).Cardinality()), _dafny.IntOfUint32((_582_v6).Cardinality()), _dafny.CodePoint('i'), globalState)
		})()))).Cardinality()
		return r0, r1
	}
}
func (_this *C1) F29() _dafny.Int {
	{
		return _this._f29
	}
}
func (_this *C1) F30() _dafny.Int {
	{
		return _this._f30
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f25 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f25 = _dafny.Zero
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f25 _dafny.Int) {
	{
		(_this)._f25 = f25
	}
}
func (_this *C2) Fm12(globalState *GlobalState) _dafny.Set {
	{
		return (_dafny.SetOf(_dafny.SeqOf(true, false))).Difference((_dafny.SetOf(_dafny.SeqOf(true, true), _dafny.SeqOf(!(false), false), _dafny.SeqOf(true), _dafny.SeqOf(!(true), true, false, true, false), _dafny.SeqOf(false, true))).Difference(_dafny.SetOf(_dafny.SeqOf(true, false))))
	}
}
func (_this *C2) M9(p0 _dafny.Array, p1 _dafny.Array, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _606_v0 _dafny.Set
		_ = _606_v0
		_606_v0 = _dafny.SetOf((_this).F25())
		var _607_i0 _dafny.Int
		_ = _607_i0
		_607_i0 = _dafny.Zero
		{
			for !((Companion_Default___.SafeDivisionInt((_this).F25(), _dafny.IntOfInt64(873))).Cmp(((_this).F25()).Times((_606_v0).Cardinality())) >= 0) {
				{
					if (_607_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_607_i0 = (_607_i0).Plus(_dafny.One)
					var _608_v1 _dafny.MultiSet
					_ = _608_v1
					_608_v1 = _dafny.MultiSetOf((_dafny.IntOfInt64(389)).Times((_this).F25()), ((_this).F25()).Plus((_this).F25()))
					var _609_v2 _dafny.CodePoint
					_ = _609_v2
					_609_v2 = _dafny.CodePoint('r')
					var _610_v3 _dafny.MultiSet
					_ = _610_v3
					_610_v3 = _dafny.MultiSetOf(_609_v2, _609_v2, _609_v2)
					var _611_v4 bool
					_ = _611_v4
					_611_v4 = false
					var _612_v5 _dafny.Set
					_ = _612_v5
					_612_v5 = _dafny.SetOf(_611_v4)
					_608_v1 = (_dafny.MultiSetOf((_this).F25(), ((_610_v3).Update(_dafny.CodePoint('c'), Companion_Default___.Abs((_612_v5).Cardinality()))).Cardinality(), (_this).F25())).Intersection((_608_v1).Difference(p2))
					var _613_v6 _dafny.Map
					_ = _613_v6
					_613_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v4, (_this).F25())
					var _614_v7 _dafny.MultiSet
					_ = _614_v7
					_614_v7 = _dafny.MultiSetOf((_613_v6).Update(_611_v4, (_this).F25()))
					var _615_v8 D4
					_ = _615_v8
					_615_v8 = Companion_D4_.Create_DC11_((_614_v7).Difference(_614_v7))
					var _source7 D4 = _615_v8
					_ = _source7
					if _source7.Is_DC12() {
						var _616___mcc_h0 bool = _source7.Get_().(D4_DC12).Cf21
						_ = _616___mcc_h0
						var _617___mcc_h1 bool = _source7.Get_().(D4_DC12).Cf22
						_ = _617___mcc_h1
						var _618___mcc_h2 _dafny.Int = _source7.Get_().(D4_DC12).Cf23
						_ = _618___mcc_h2
						var _619_cf23 _dafny.Int = _618___mcc_h2
						_ = _619_cf23
						var _620_cf22 bool = _617___mcc_h1
						_ = _620_cf22
						var _621_cf21 bool = _616___mcc_h0
						_ = _621_cf21
						var _622_v9 _dafny.Array
						_ = _622_v9
						var _nw115 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
						_ = _nw115
						_622_v9 = _nw115
						var _623_v10 *C0
						_ = _623_v10
						var _nw116 *C0 = New_C0_()
						_ = _nw116
						_nw116.Ctor__(Companion_Default___.SafeModuloInt((_this).F25(), (_this).F25()), _622_v9)
						_623_v10 = _nw116
						(globalState).F5 = _611_v4
						var _624_v11 _dafny.Sequence
						_ = _624_v11
						_624_v11 = _dafny.UnicodeSeqOfUtf8Bytes("rlvehw")
						var _625_v12 _dafny.Sequence
						_ = _625_v12
						_625_v12 = _dafny.SeqOf(false, _611_v4, _621_cf21, _620_cf22)
						var _rhs140 bool = (_625_v12).Select((Companion_Default___.SafeIndex(_619_cf23, _dafny.IntOfUint32((_625_v12).Cardinality()))).Uint32()).(bool)
						_ = _rhs140
						var _rhs141 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_624_v11, Companion_Default___.Fm1(globalState))
						_ = _rhs141
						var _lhs123 *GlobalState = globalState
						_ = _lhs123
						_lhs123.F5 = _rhs140
						_624_v11 = _rhs141
						var _626_v13 _dafny.Array
						_ = _626_v13
						var _len0_18 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_18
						var _nw117 _dafny.Array
						_ = _nw117
						if _len0_18.Cmp(_dafny.Zero) == 0 {
							_nw117 = _dafny.NewArray(_len0_18)
						} else {
							var _init18 func(_dafny.Int) _dafny.Int = func(_627_i1 _dafny.Int) _dafny.Int {
								return (_627_i1).Plus((_dafny.Zero).Minus((_this).F25()))
							}
							_ = _init18
							var _element0_18 = _init18(_dafny.Zero)
							_ = _element0_18
							_nw117 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
							(_nw117).ArraySet1(_element0_18, 0)
							var _nativeLen0_18 = (_len0_18).Int()
							_ = _nativeLen0_18
							for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
								(_nw117).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
							}
						}
						_626_v13 = _nw117
						var _628_v14 _dafny.MultiSet
						_ = _628_v14
						_628_v14 = _dafny.MultiSetOf(_dafny.SeqOf(_620_cf22), _625_v12)
						var _629_v15 _dafny.Map
						_ = _629_v15
						_629_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_623_v10).F26()).Plus((_628_v14).Cardinality()), _626_v13)
						var _630_v16 _dafny.Map
						_ = _630_v16
						_630_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v4, p0)
						_626_v13 = (func() _dafny.Array {
							if (_629_v15).Contains((_dafny.Zero).Minus(((_this).F25()).Minus(_619_cf23))) {
								return (_629_v15).Get((_dafny.Zero).Minus(((_this).F25()).Minus(_619_cf23))).(_dafny.Array)
							}
							return (func() _dafny.Array {
								if (_630_v16).Contains(false) {
									return (_630_v16).Get(false).(_dafny.Array)
								}
								return p0
							})()
						})()
					} else if _source7.Is_DC13() {
						var _631___mcc_h3 _dafny.Int = _source7.Get_().(D4_DC13).Cf24
						_ = _631___mcc_h3
						var _632___mcc_h4 _dafny.Int = _source7.Get_().(D4_DC13).Cf25
						_ = _632___mcc_h4
						var _633___mcc_h5 _dafny.Int = _source7.Get_().(D4_DC13).Cf26
						_ = _633___mcc_h5
						var _634_cf26 _dafny.Int = _633___mcc_h5
						_ = _634_cf26
						var _635_cf25 _dafny.Int = _632___mcc_h4
						_ = _635_cf25
						var _636_cf24 _dafny.Int = _631___mcc_h3
						_ = _636_cf24
						var _637_v17 _dafny.Array
						_ = _637_v17
						var _nwElement0_21 bool = _611_v4
						_ = _nwElement0_21
						var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(2))
						_ = _nw118
						(_nw118).ArraySet1(_nwElement0_21, 0)
						(_nw118).ArraySet1(true, 1)
						_637_v17 = _nw118
						var _638_v18 D0
						_ = _638_v18
						_638_v18 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus((_this).F25()), (_this).F25(), _634_cf26, _637_v17)
						var _pat_let_tv17 = _635_cf25
						_ = _pat_let_tv17
						var _pat_let_tv18 = _636_cf24
						_ = _pat_let_tv18
						_635_cf25 = (func(_pat_let13_0 D0) D0 {
							return func(_639_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let14_0 _dafny.Int) D0 {
									return func(_640_dt__update_hcf4_h0 _dafny.Int) D0 {
										return func(_pat_let15_0 _dafny.Int) D0 {
											return func(_641_dt__update_hcf2_h0 _dafny.Int) D0 {
												return Companion_D0_.Create_DC2_(_641_dt__update_hcf2_h0, (_639_dt__update__tmp_h0).Dtor_cf3(), _640_dt__update_hcf4_h0, (_639_dt__update__tmp_h0).Dtor_cf5())
											}(_pat_let15_0)
										}(_pat_let_tv18)
									}(_pat_let14_0)
								}(_pat_let_tv17)
							}(_pat_let13_0)
						}(_638_v18)).Dtor_cf4()
						(globalState).F2 = _611_v4
						var _642_v19 _dafny.Map
						_ = _642_v19
						_642_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_634_cf26), true)
						(globalState).F5 = (func() bool {
							if (_642_v19).Contains(p2) {
								return (_642_v19).Get(p2).(bool)
							}
							return (_635_cf25).Cmp(_dafny.IntOfInt64(543)) != 0
						})()
						var _643_v20 _dafny.Sequence
						_ = _643_v20
						_643_v20 = _dafny.UnicodeSeqOfUtf8Bytes("rsh")
						var _644_v21 D4
						_ = _644_v21
						_644_v21 = Companion_D4_.Create_DC12_(_611_v4, true, _dafny.IntOfUint32((_643_v20).Cardinality()))
						var _rhs142 _dafny.CodePoint = _609_v2
						_ = _rhs142
						var _rhs143 D4 = _644_v21
						_ = _rhs143
						var _lhs124 *GlobalState = globalState
						_ = _lhs124
						_lhs124.F15 = _rhs142
						_644_v21 = _rhs143
					} else {
						var _645___mcc_h6 _dafny.MultiSet = _source7.Get_().(D4_DC11).Cf20
						_ = _645___mcc_h6
						var _646_cf20 _dafny.MultiSet = _645___mcc_h6
						_ = _646_cf20
						r0 = (_this).F25()
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((p0), 0))
						_ = _index120
						(p0).ArraySet1((_this).F25(), (_index120).Int())
						var _647_v22 _dafny.Sequence
						_ = _647_v22
						_647_v22 = _dafny.SeqOf(_611_v4)
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((p0), 0))
						_ = _index121
						(p0).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(((_this).F25()).Minus((func() _dafny.Int {
							if (_608_v1).Contains((_this).F25()) {
								return (_608_v1).Multiplicity((_this).F25())
							}
							return _dafny.IntOfUint32((_647_v22).Cardinality())
						})()))), (_index121).Int())
						var _648_v23 _dafny.Array
						_ = _648_v23
						var _nw119 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(23))
						_ = _nw119
						_648_v23 = _nw119
						var _649_v24 _dafny.Map
						_ = _649_v24
						_649_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)), (_this).F25())
						var _650_v25 _dafny.Map
						_ = _650_v25
						_650_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_649_v24, _611_v4)
						var _651_v26 _dafny.Map
						_ = _651_v26
						_651_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_648_v23, (func() bool {
							if (_650_v25).Contains(_649_v24) {
								return (_650_v25).Get(_649_v24).(bool)
							}
							return _611_v4
						})())
						_651_v26 = (_651_v26).Update(_648_v23, _611_v4)
						var _652_v27 _dafny.Array
						_ = _652_v27
						var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
						_ = _nw120
						_652_v27 = _nw120
						var _653_v28 _dafny.Map
						_ = _653_v28
						_653_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v4, _611_v4)
						var _654_v29 _dafny.Array
						_ = _654_v29
						var _nwElement0_22 bool = (func() bool {
							if (_653_v28).Contains(_611_v4) {
								return (_653_v28).Get(_611_v4).(bool)
							}
							return _611_v4
						})()
						_ = _nwElement0_22
						var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(28))
						_ = _nw121
						(_nw121).ArraySet1(_nwElement0_22, 0)
						(_nw121).ArraySet1(_611_v4, 1)
						(_nw121).ArraySet1(false, 2)
						(_nw121).ArraySet1(_611_v4, 3)
						(_nw121).ArraySet1(_611_v4, 4)
						(_nw121).ArraySet1(_611_v4, 5)
						(_nw121).ArraySet1(_611_v4, 6)
						(_nw121).ArraySet1(_611_v4, 7)
						(_nw121).ArraySet1(_611_v4, 8)
						(_nw121).ArraySet1(_611_v4, 9)
						(_nw121).ArraySet1(_611_v4, 10)
						(_nw121).ArraySet1(_611_v4, 11)
						(_nw121).ArraySet1(_611_v4, 12)
						(_nw121).ArraySet1(_611_v4, 13)
						(_nw121).ArraySet1(_611_v4, 14)
						(_nw121).ArraySet1(_611_v4, 15)
						(_nw121).ArraySet1(_611_v4, 16)
						(_nw121).ArraySet1(_611_v4, 17)
						(_nw121).ArraySet1(_611_v4, 18)
						(_nw121).ArraySet1(!(!(_611_v4)), 19)
						(_nw121).ArraySet1(_611_v4, 20)
						(_nw121).ArraySet1(_611_v4, 21)
						(_nw121).ArraySet1(_611_v4, 22)
						(_nw121).ArraySet1(_611_v4, 23)
						(_nw121).ArraySet1(_611_v4, 24)
						(_nw121).ArraySet1(_611_v4, 25)
						(_nw121).ArraySet1(true, 26)
						(_nw121).ArraySet1(_611_v4, 27)
						_654_v29 = _nw121
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_652_v27), 0))
						_ = _index122
						(_652_v27).ArraySet1(_654_v29, (_index122).Int())
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(316), _dafny.ArrayLen((_652_v27), 0))
						_ = _index123
						(_652_v27).ArraySet1(_654_v29, (_index123).Int())
					}
					var _655_v30 _dafny.Array
					_ = _655_v30
					var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
					_ = _nw122
					_655_v30 = _nw122
					var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))
					_ = _index124
					(_655_v30).ArraySet1(_611_v4, (_index124).Int())
					var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))
					_ = _index125
					(_655_v30).ArraySet1(_611_v4, (_index125).Int())
					if !(!(_611_v4)) || ((Companion_Default___.Fm3(globalState)).Cmp((_this).F25()) < 0) {
						var _656_v31 D3
						_ = _656_v31
						_656_v31 = Companion_D3_.Create_DC10_((_655_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))).Int()).(bool))
						(globalState).F5 = (_656_v31).Dtor_cf19()
						(globalState).F2 = false
						var _657_v32 _dafny.Map
						_ = _657_v32
						_657_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_611_v4, _612_v5)
						_612_v5 = (((func() _dafny.Set {
							if (_657_v32).Contains((_655_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))).Int()).(bool)) {
								return (_657_v32).Get((_655_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))).Int()).(bool)).(_dafny.Set)
							}
							return _612_v5
						})()).Union(_612_v5)).Intersection(_dafny.SetOf(false))
						(globalState).F5 = (_655_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))).Int()).(bool)
						var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))
						_ = _index126
						(_655_v30).ArraySet1(((_this).F25()).Cmp(((_dafny.Zero).Minus((_this).F25())).Plus((_dafny.Zero).Minus((p2).Cardinality()))) == 0, (_index126).Int())
					} else {
						var _658_v34 _dafny.Map
						_ = _658_v34
						_658_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_613_v6, (_this).F25())
						var _659_v35 *C0
						_ = _659_v35
						var _nw123 *C0 = New_C0_()
						_ = _nw123
						_nw123.Ctor__((func() _dafny.Map {
							var _coll41 = _dafny.NewMapBuilder()
							_ = _coll41
							for _iter45 := _dafny.Iterate((_658_v34).Keys().Elements()); ; {
								_compr_41, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _660_v33 _dafny.Map
								_660_v33 = interface{}(_compr_41).(_dafny.Map)
								if (_658_v34).Contains(_660_v33) {
									_coll41.Add(_660_v33, (_this).F25())
								}
							}
							return _coll41.ToMap()
						}()).Cardinality(), (func() _dafny.Array {
							if _611_v4 {
								return _655_v30
							}
							return _655_v30
						})())
						_659_v35 = _nw123
						var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))
						_ = _index127
						var _rhs144 _dafny.Int = (_659_v35).F26()
						_ = _rhs144
						var _rhs145 bool = (_655_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))).Int()).(bool)
						_ = _rhs145
						var _lhs125 _dafny.Array = _655_v30
						_ = _lhs125
						var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_655_v30), 0))
						_ = _lhs126
						r0 = _rhs144
						(_lhs125).ArraySet1(_rhs145, (_lhs126).Int())
						r0 = ((_659_v35).F26()).Times((_659_v35).F26())
						var _661_v36 _dafny.Map
						_ = _661_v36
						_661_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v2, Companion_D1_.Create_DC6_())
						var _662_v37 D1
						_ = _662_v37
						_662_v37 = Companion_D1_.Create_DC6_()
						_661_v36 = (_661_v36).Update(_609_v2, _662_v37)
						_655_v30 = _659_v35.F27
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _663_v38 bool
		_ = _663_v38
		_663_v38 = false
		var _664_i2 _dafny.Int
		_ = _664_i2
		_664_i2 = _dafny.Zero
		{
			for _663_v38 {
				{
					if (_664_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_664_i2 = (_664_i2).Plus(_dafny.One)
					var _665_v39 _dafny.Array
					_ = _665_v39
					var _nw124 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(22))
					_ = _nw124
					_665_v39 = _nw124
					var _666_v40 _dafny.Array
					_ = _666_v40
					var _nw125 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
					_ = _nw125
					_666_v40 = _nw125
					var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_665_v39), 0))
					_ = _index128
					(_665_v39).ArraySet1(Companion_D0_.Create_DC2_((_this).F25(), (_this).F25(), (_this).F25(), _666_v40), (_index128).Int())
					var _667_v41 D0
					_ = _667_v41
					_667_v41 = Companion_D0_.Create_DC2_((_this).F25(), (func() _dafny.Int {
						if (p2).Contains((_this).F25()) {
							return (p2).Multiplicity((_this).F25())
						}
						return (_this).F25()
					})(), ((_this).F25()).Plus(_dafny.IntOfInt64(730)), _666_v40)
					var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_665_v39), 0))
					_ = _index129
					var _rhs146 D0 = _667_v41
					_ = _rhs146
					var _rhs147 _dafny.Int = _dafny.IntOfInt64(-984)
					_ = _rhs147
					var _lhs127 _dafny.Array = _665_v39
					_ = _lhs127
					var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_665_v39), 0))
					_ = _lhs128
					(_lhs127).ArraySet1(_rhs146, (_lhs128).Int())
					r0 = _rhs147
					var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))
					_ = _index130
					(p0).ArraySet1(Companion_Default___.Fm3(globalState), (_index130).Int())
					var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))
					_ = _index131
					(p0).ArraySet1(_dafny.IntOfInt64(-683), (_index131).Int())
					var _668_v42 _dafny.MultiSet
					_ = _668_v42
					_668_v42 = _dafny.MultiSetOf(_663_v38)
					var _669_v43 D1
					_ = _669_v43
					_669_v43 = Companion_D1_.Create_DC3_(_668_v42)
					_669_v43 = _669_v43
					if _663_v38 {
						var _670_v44 _dafny.Array
						_ = _670_v44
						var _len0_19 _dafny.Int = _dafny.IntOfInt64(23)
						_ = _len0_19
						var _nw126 _dafny.Array
						_ = _nw126
						if _len0_19.Cmp(_dafny.Zero) == 0 {
							_nw126 = _dafny.NewArray(_len0_19)
						} else {
							var _init19 func(_dafny.Int) _dafny.Int = (func(_671_p0 _dafny.Array) func(_dafny.Int) _dafny.Int {
								return func(_672_i3 _dafny.Int) _dafny.Int {
									return (_672_i3).Plus((_671_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_671_p0), 0))).Int()).(_dafny.Int))
								}
							})(p0)
							_ = _init19
							var _element0_19 = _init19(_dafny.Zero)
							_ = _element0_19
							_nw126 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
							(_nw126).ArraySet1(_element0_19, 0)
							var _nativeLen0_19 = (_len0_19).Int()
							_ = _nativeLen0_19
							for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
								(_nw126).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
							}
						}
						_670_v44 = _nw126
						_670_v44 = _670_v44
						var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))
						_ = _index132
						(p0).ArraySet1((((p2).Cardinality()).Times(_dafny.IntOfInt64(257))).Plus((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)), (_index132).Int())
						var _673_v45 _dafny.Array
						_ = _673_v45
						var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
						_ = _nw127
						_673_v45 = _nw127
						var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_673_v45), 0))
						_ = _index133
						(_673_v45).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(372))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}(func(_674_i4 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('o')
						})), (_index133).Int())
						var _675_v46 _dafny.Sequence
						_ = _675_v46
						_675_v46 = _dafny.UnicodeSeqOfUtf8Bytes("hmr")
						var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_673_v45), 0))
						_ = _index134
						(_673_v45).ArraySet1(_675_v46, (_index134).Int())
						var _676_v47 _dafny.Sequence
						_ = _676_v47
						_676_v47 = _dafny.SeqOf((_this).F25())
						var _677_v48 T0
						_ = _677_v48
						var _nw128 *C1 = New_C1_()
						_ = _nw128
						_nw128.Ctor__((_this).F25(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _663_v38, _676_v47, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
						_677_v48 = _nw128
						var _678_v49 _dafny.Map
						_ = _678_v49
						_678_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_663_v38), _677_v48)
						var _rhs148 bool = _663_v38
						_ = _rhs148
						var _rhs149 _dafny.Map = _678_v49
						_ = _rhs149
						var _lhs129 *GlobalState = globalState
						_ = _lhs129
						_lhs129.F2 = _rhs148
						_678_v49 = _rhs149
						r0 = (func() _dafny.Int {
							if _677_v48.F16() {
								return (func() _dafny.Int {
									if (_668_v42).Contains(_677_v48.F16()) {
										return (_668_v42).Multiplicity(_677_v48.F16())
									}
									return Companion_Default___.Fm3(globalState)
								})()
							}
							return ((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Times((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
						})()
					} else {
						_663_v38 = !(_663_v38)
						var _679_v50 _dafny.Sequence
						_ = _679_v50
						_679_v50 = _dafny.SeqOf(!(_663_v38))
						var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))
						_ = _index135
						(p0).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F25(), _dafny.IntOfUint32((_679_v50).Cardinality())), (_index135).Int())
						var _680_v51 _dafny.Map
						_ = _680_v51
						_680_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v38, _663_v38)
						var _681_v52 D1
						_ = _681_v52
						_681_v52 = Companion_D1_.Create_DC4_(!(_663_v38), (_this).F25(), _680_v51)
						var _682_v53 _dafny.Map
						_ = _682_v53
						_682_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_681_v52, _606_v0)
						_682_v53 = (_682_v53).Update(_681_v52, _dafny.SetOf((_this).F25()))
						var _683_v54 _dafny.Array
						_ = _683_v54
						var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
						_ = _nw129
						_683_v54 = _nw129
						var _684_v55 D6
						_ = _684_v55
						_684_v55 = Companion_D6_.Create_DC15_(_683_v54)
						_683_v54 = (_684_v55).Dtor_cf28()
						var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((p0), 0))
						_ = _index136
						(p0).ArraySet1((_this).F25(), (_index136).Int())
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (func() _dafny.Int {
			if _663_v38 {
				return (_this).F25()
			}
			return Companion_Default___.SafeDivisionInt((_this).F25(), (_this).F25())
		})()
		var _685_v56 _dafny.Array
		_ = _685_v56
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw130
		_685_v56 = _nw130
		var _686_v57 *C0
		_ = _686_v57
		var _nw131 *C0 = New_C0_()
		_ = _nw131
		_nw131.Ctor__((_this).F25(), _685_v56)
		_686_v57 = _nw131
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_685_v56), 0))
		_ = _index137
		(_685_v56).ArraySet1(_663_v38, (_index137).Int())
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_685_v56), 0))
		_ = _index138
		(_685_v56).ArraySet1((((_this).F25()).Times(_dafny.IntOfInt64(-768))).Cmp((_this).F25()) <= 0, (_index138).Int())
		var _687_v58 _dafny.CodePoint
		_ = _687_v58
		_687_v58 = _dafny.CodePoint('a')
		(globalState).F15 = _687_v58
		var _688_v59 _dafny.Array
		_ = _688_v59
		var _nw132 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(20))
		_ = _nw132
		_688_v59 = _nw132
		var _689_v60 _dafny.Map
		_ = _689_v60
		_689_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_688_v59, (_685_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_685_v56), 0))).Int()).(bool))
		var _690_v61 _dafny.Array
		_ = _690_v61
		_690_v61 = _688_v59
		r0 = (_dafny.SetOf((_this).F25(), ((_689_v60).Merge((_689_v60).Update(_690_v61, (_685_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_685_v56), 0))).Int()).(bool)))).Cardinality())).Cardinality()
		return r0
	}
}
func (_this *C2) M10(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _691_v0 _dafny.CodePoint
		_ = _691_v0
		_691_v0 = _dafny.CodePoint('k')
		(globalState).F15 = _691_v0
		r0 = Companion_Default___.Fm3(globalState)
		var _692_v1 D6
		_ = _692_v1
		_692_v1 = Companion_D6_.Create_DC16_(!(!(p0)), p0)
		r0 = func(_source8 D6) _dafny.Int {
			if _source8.Is_DC16() {
				var _693___mcc_h0 bool = _source8.Get_().(D6_DC16).Cf29
				_ = _693___mcc_h0
				var _694___mcc_h1 bool = _source8.Get_().(D6_DC16).Cf30
				_ = _694___mcc_h1
				var _695_cf30 bool = _694___mcc_h1
				_ = _695_cf30
				var _696_cf29 bool = _693___mcc_h0
				_ = _696_cf29
				return (_this).F25()
			} else if _source8.Is_DC17() {
				var _697___mcc_h2 D2 = _source8.Get_().(D6_DC17).Cf31
				_ = _697___mcc_h2
				var _698_cf31 D2 = _697___mcc_h2
				_ = _698_cf31
				return (_this).F25()
			} else if _source8.Is_DC18() {
				var _699___mcc_h3 _dafny.Array = _source8.Get_().(D6_DC18).Cf32
				_ = _699___mcc_h3
				var _700_cf32 _dafny.Array = _699___mcc_h3
				_ = _700_cf32
				return ((_this).F25()).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}(func(_701_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()))
			} else {
				var _702___mcc_h4 _dafny.Array = _source8.Get_().(D6_DC15).Cf28
				_ = _702___mcc_h4
				var _703_cf28 _dafny.Array = _702___mcc_h4
				_ = _703_cf28
				return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(584), (_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F25())))
			}
		}((func() D6 {
			if p0 {
				return _692_v1
			}
			return _692_v1
		})())
		var _hi5 _dafny.Int = (_this).F25()
		_ = _hi5
		for _704_i1 := (_this).F25(); _704_i1.Cmp(_hi5) < 0; _704_i1 = _704_i1.Plus(_dafny.One) {
			var _705_v2 _dafny.Sequence
			_ = _705_v2
			_705_v2 = _dafny.SeqOf(_dafny.IntOfInt64(459))
			var _706_v3 *C1
			_ = _706_v3
			var _nw133 *C1 = New_C1_()
			_ = _nw133
			_nw133.Ctor__((_this).F25(), (_this).F25(), p0, _705_v2, _704_i1)
			_706_v3 = _nw133
			(globalState).F2 = false
			var _707_v4 _dafny.Sequence
			_ = _707_v4
			_707_v4 = _dafny.UnicodeSeqOfUtf8Bytes("vbabos")
			var _708_v5 _dafny.Map
			_ = _708_v5
			_708_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), _707_v4)).Cardinality())
			var _709_v6 _dafny.Array
			_ = _709_v6
			var _nw134 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw134
			_709_v6 = _nw134
			var _710_v7 _dafny.Set
			_ = _710_v7
			_710_v7 = _dafny.SetOf((_706_v3).F30(), (_this).F25())
			var _711_v8 _dafny.Map
			_ = _711_v8
			_711_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_706_v3, p0)
			var _712_v9 D1
			_ = _712_v9
			_712_v9 = Companion_D1_.Create_DC5_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_710_v7).Cardinality(), _dafny.IntOfInt64(99))).Cardinality(), (func() bool {
				if (_711_v8).Contains(_706_v3) {
					return (_711_v8).Get(_706_v3).(bool)
				}
				return p0
			})())
			var _713_v10 D2
			_ = _713_v10
			_713_v10 = Companion_D2_.Create_DC8_(_708_v5, _709_v6, p0, (_712_v9).Dtor_cf11(), (_dafny.Zero).Minus((_this).F25()))
			var _714_v11 _dafny.MultiSet
			_ = _714_v11
			_714_v11 = _dafny.MultiSetOf(((_706_v3).F30()).Plus((_706_v3).F29()), (_706_v3).Fm14((_this).F25(), (_706_v3).F30(), _691_v0, globalState), (_713_v10).Dtor_cf17())
			_714_v11 = _714_v11
			var _715_v12 D1
			_ = _715_v12
			_715_v12 = Companion_D1_.Create_DC6_()
			_715_v12 = _715_v12
		}
		var _716_v13 _dafny.Sequence
		_ = _716_v13
		_716_v13 = _dafny.UnicodeSeqOfUtf8Bytes("lskhib")
		_716_v13 = _dafny.Companion_Sequence_.Concatenate(_716_v13, _dafny.UnicodeSeqOfUtf8Bytes("otcfiuf"))
		r0 = (_dafny.IntOfUint32((_716_v13).Cardinality())).Times((_this).F25())
		r0 = (_this).F25()
		var _717_v14 _dafny.Array
		_ = _717_v14
		var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
		_ = _nw135
		_717_v14 = _nw135
		r1 = (func() _dafny.Array {
			if (p0) && (!(p0)) {
				return _717_v14
			}
			return _717_v14
		})()
		return r0, r1
	}
}
func (_this *C2) F25() _dafny.Int {
	{
		return _this._f25
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f16 bool
	_f17 _dafny.Sequence
	_f24 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	_this._f24 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F16() bool {
	return _this._f16
}
func (_this *C3) F16_set_(value bool) {
	_this._f16 = value
}
func (_this *C3) F17() _dafny.Sequence {
	return _this._f17
}
func (_this *C3) Ctor__(f24 _dafny.Sequence, f16 bool, f17 _dafny.Sequence) {
	{
		(_this)._f24 = f24
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C3) Fm11(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return (_this).F24()
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _718_v1 _dafny.Array
		_ = _718_v1
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_20
		var _nw136 _dafny.Array
		_ = _nw136
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw136 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Set = func(_719_i0 _dafny.Int) _dafny.Set {
				return func() _dafny.Set {
					var _coll42 = _dafny.NewBuilder()
					_ = _coll42
					for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(757), _dafny.IntOfInt64(81))); ; {
						_compr_42, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _720_v0 _dafny.Int
						_720_v0 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(757)).Cmp(_720_v0) <= 0) && ((_720_v0).Cmp(_dafny.IntOfInt64(81)) < 0) {
							_coll42.Add((_720_v0).Plus((_dafny.MultiSetOf(_this.F16())).Cardinality()))
						}
					}
					return _coll42.ToSet()
				}()
			}
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw136 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw136).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw136).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_718_v1 = _nw136
		var _721_v2 _dafny.Set
		_ = _721_v2
		_721_v2 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F16())).Cardinality())
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_718_v1), 0))
		_ = _index139
		(_718_v1).ArraySet1(_721_v2, (_index139).Int())
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_718_v1), 0))
		_ = _index140
		var _rhs150 _dafny.Set = _721_v2
		_ = _rhs150
		var _rhs151 bool = _this.F16()
		_ = _rhs151
		var _rhs152 bool = (func() bool {
			if (_721_v2).IsProperSubsetOf(_721_v2) {
				return (p3).Cmp(p3) < 0
			}
			return _this.F16()
		})()
		_ = _rhs152
		var _lhs130 _dafny.Array = _718_v1
		_ = _lhs130
		var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_718_v1), 0))
		_ = _lhs131
		var _lhs132 *GlobalState = globalState
		_ = _lhs132
		var _lhs133 *GlobalState = globalState
		_ = _lhs133
		(_lhs130).ArraySet1(_rhs150, (_lhs131).Int())
		_lhs132.F5 = _rhs151
		_lhs133.F2 = _rhs152
		(_this).F16_set_(p1)
		var _722_v3 _dafny.Array
		_ = _722_v3
		var _nw137 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw137
		_722_v3 = _nw137
		var _723_v4 _dafny.Set
		_ = _723_v4
		_723_v4 = _dafny.SetOf(_722_v3, _722_v3)
		var _724_i1 _dafny.Int
		_ = _724_i1
		_724_i1 = _dafny.Zero
		{
			for ((_723_v4).Difference(_723_v4)).Contains(_722_v3) {
				{
					if (_724_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_724_i1 = (_724_i1).Plus(_dafny.One)
					var _725_v5 D3
					_ = _725_v5
					_725_v5 = Companion_D3_.Create_DC9_((_this).F17())
					var _726_v6 _dafny.Map
					_ = _726_v6
					_726_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_722_v3, p3)
					var _727_v7 _dafny.Map
					_ = _727_v7
					_727_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_725_v5, ((_726_v6).Merge(_726_v6)).Cardinality())
					_727_v7 = (_727_v7).Update(_725_v5, p3)
					var _728_v8 _dafny.Array
					_ = _728_v8
					var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw138
					_728_v8 = _nw138
					var _729_v9 _dafny.Sequence
					_ = _729_v9
					_729_v9 = _dafny.SeqOf(_this.F16(), p1)
					var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_728_v8), 0))
					_ = _index141
					(_728_v8).ArraySet1(_dafny.IntOfUint32((_729_v9).Cardinality()), (_index141).Int())
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_728_v8), 0))
					_ = _index142
					(_728_v8).ArraySet1(p3, (_index142).Int())
					var _730_v10 _dafny.Map
					_ = _730_v10
					_730_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(false))
					_730_v10 = (_730_v10).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-629))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}(func(_731_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('s')
					})), _this.F16())
					var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_728_v8), 0))
					_ = _index143
					(_728_v8).ArraySet1((_728_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_728_v8), 0))).Int()).(_dafny.Int), (_index143).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _732_v11 _dafny.Map
		_ = _732_v11
		_732_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F24())
		var _733_v12 _dafny.Map
		_ = _733_v12
		_733_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16())
		var _734_v13 _dafny.Sequence
		_ = _734_v13
		_734_v13 = _dafny.SeqOf(_733_v12)
		var _735_v14 _dafny.Map
		_ = _735_v14
		_735_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_v11, (func() _dafny.Sequence {
			if _this.F16() {
				return _734_v13
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(607))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_736_v12 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_737_i3 _dafny.Int) _dafny.Map {
					return _736_v12
				}
			})(_733_v12)))
		})())
		_735_v14 = (_735_v14).Update(_732_v11, _734_v13)
		var _738_v15 _dafny.Int
		_ = _738_v15
		_738_v15 = _dafny.IntOfInt64(-688)
		_738_v15 = (_dafny.Zero).Minus((_738_v15).Times(_738_v15))
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_722_v3), 0))
		_ = _index144
		(_722_v3).ArraySet1(_this.F16(), (_index144).Int())
		var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_722_v3), 0))
		_ = _index145
		var _rhs153 bool = _this.F16()
		_ = _rhs153
		var _rhs154 bool = !(Companion_Default___.Fm0(p1, _this.F16(), globalState))
		_ = _rhs154
		var _lhs134 *GlobalState = globalState
		_ = _lhs134
		var _lhs135 _dafny.Array = _722_v3
		_ = _lhs135
		var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_722_v3), 0))
		_ = _lhs136
		_lhs134.F5 = _rhs153
		(_lhs135).ArraySet1(_rhs154, (_lhs136).Int())
		var _739_v16 _dafny.Sequence
		_ = _739_v16
		_739_v16 = _dafny.SeqOf(p1)
		r0 = _dafny.Companion_Sequence_.Concatenate(_739_v16, _739_v16)
		return r0
	}
}
func (_this *C3) M8(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Map) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var _740_v0 _dafny.Int
		_ = _740_v0
		_740_v0 = _dafny.IntOfInt64(69)
		var _741_v1 _dafny.Sequence
		_ = _741_v1
		_741_v1 = _dafny.SeqOf(_this.F16(), !(_this.F16()))
		var _742_v2 _dafny.Map
		_ = _742_v2
		_742_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-279), Companion_Default___.SafeDivisionInt(_740_v0, _dafny.IntOfUint32((_741_v1).Cardinality())))
		var _743_v3 _dafny.Map
		_ = _743_v3
		_743_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _740_v0)
		_742_v2 = (_742_v2).Update(((_this).F17()).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), (_740_v0).Times((_743_v3).Cardinality()))
		var _744_i0 _dafny.Int
		_ = _744_i0
		_744_i0 = _dafny.Zero
		{
			for Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState) {
				{
					if (_744_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_744_i0 = (_744_i0).Plus(_dafny.One)
					if (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_this).F24()).Cardinality()), (_dafny.Zero).Minus(_740_v0))).Cmp(_dafny.IntOfInt64(878)) == 0 {
						(globalState).F2 = _this.F16()
						var _745_v4 _dafny.Array
						_ = _745_v4
						var _nw139 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
						_ = _nw139
						_745_v4 = _nw139
						var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))
						_ = _index146
						(_745_v4).ArraySet1(!(_this.F16()), (_index146).Int())
						var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))
						_ = _index147
						var _rhs155 _dafny.Map = (func() _dafny.Map {
							if !(_this.F16()) {
								return _742_v2
							}
							return _742_v2
						})()
						_ = _rhs155
						var _rhs156 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F24()).Cardinality())), _740_v0)
						_ = _rhs156
						var _rhs157 bool = _this.F16()
						_ = _rhs157
						var _lhs137 _dafny.Array = _745_v4
						_ = _lhs137
						var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))
						_ = _lhs138
						_742_v2 = _rhs155
						_740_v0 = _rhs156
						(_lhs137).ArraySet1(_rhs157, (_lhs138).Int())
						var _746_v5 _dafny.Array
						_ = _746_v5
						var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
						_ = _nw140
						_746_v5 = _nw140
						var _747_v6 _dafny.CodePoint
						_ = _747_v6
						_747_v6 = _dafny.CodePoint('w')
						var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_746_v5), 0))
						_ = _index148
						(_746_v5).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_v0, _747_v6), (_index148).Int())
						var _748_v7 D3
						_ = _748_v7
						_748_v7 = Companion_D3_.Create_DC10_((_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool))
						var _749_v8 _dafny.Set
						_ = _749_v8
						_749_v8 = _dafny.SetOf(_this.F16(), _this.F16())
						var _750_v9 _dafny.Map
						_ = _750_v9
						_750_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_740_v0, _747_v6)
						var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_746_v5), 0))
						_ = _index149
						var _rhs158 bool = ((func() bool {
							if (_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool) {
								return (_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool)
							}
							return !((_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool))
						})()) == ((_749_v8).Contains((_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool)))
						_ = _rhs158
						var _rhs159 _dafny.Map = (func() _dafny.Map {
							if Companion_Default___.Fm0(!(_this.F16()), _this.F16(), globalState) {
								return _750_v9
							}
							return (_750_v9).Merge(_750_v9)
						})()
						_ = _rhs159
						var _rhs160 D3 = _748_v7
						_ = _rhs160
						var _lhs139 *GlobalState = globalState
						_ = _lhs139
						var _lhs140 _dafny.Array = _746_v5
						_ = _lhs140
						var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_746_v5), 0))
						_ = _lhs141
						_lhs139.F5 = _rhs158
						(_lhs140).ArraySet1(_rhs159, (_lhs141).Int())
						_748_v7 = _rhs160
						r1 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_740_v0).Times(_740_v0)))
						var _751_v10 _dafny.Map
						_ = _751_v10
						_751_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_745_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_745_v4), 0))).Int()).(bool))
						var _752_v11 D1
						_ = _752_v11
						_752_v11 = Companion_D1_.Create_DC4_(false, (_dafny.Zero).Minus(_740_v0), _751_v10)
						var _753_v12 _dafny.MultiSet
						_ = _753_v12
						_753_v12 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(832)), (_743_v3).Merge(_743_v3), _743_v3)
						var _754_v13 _dafny.Sequence
						_ = _754_v13
						_754_v13 = _dafny.SeqOf(_dafny.CodePoint('t'))
						var _755_v14 D4
						_ = _755_v14
						_755_v14 = Companion_D4_.Create_DC11_(_753_v12)
						var _rhs161 D1 = _752_v11
						_ = _rhs161
						var _rhs162 _dafny.MultiSet = (_dafny.MultiSetOf(_743_v3, _743_v3)).Intersection((_755_v14).Dtor_cf20())
						_ = _rhs162
						var _rhs163 _dafny.Int = (_740_v0).Minus(_740_v0)
						_ = _rhs163
						var _rhs164 _dafny.Int = _740_v0
						_ = _rhs164
						var _rhs165 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(((_this).F24()).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32(((_this).F24()).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.CodePoint('a')), (Companion_Default___.SafeIndex((_740_v0).Times(_740_v0), _dafny.IntOfUint32((_dafny.SeqOf(((_this).F24()).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32(((_this).F24()).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.CodePoint('a'))).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
							if _this.F16() {
								return _747_v6
							}
							return (func() _dafny.CodePoint {
								if (_750_v9).Contains(_740_v0) {
									return (_750_v9).Get(_740_v0).(_dafny.CodePoint)
								}
								return _747_v6
							})()
						})())
						_ = _rhs165
						_752_v11 = _rhs161
						_753_v12 = _rhs162
						_740_v0 = _rhs163
						_740_v0 = _rhs164
						_754_v13 = _rhs165
					} else {
						var _rhs166 _dafny.Int = _740_v0
						_ = _rhs166
						var _rhs167 _dafny.Int = _dafny.IntOfInt64(-554)
						_ = _rhs167
						_740_v0 = _rhs166
						r1 = _rhs167
						(globalState).F2 = _this.F16()
						var _756_v15 _dafny.Map
						_ = _756_v15
						_756_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16())
						_741_v1 = (func() _dafny.Sequence {
							if (func() bool {
								if (_756_v15).Contains(_this.F16()) {
									return (_756_v15).Get(_this.F16()).(bool)
								}
								return _this.F16()
							})() {
								return _741_v1
							}
							return _741_v1
						})()
						var _757_v16 _dafny.Set
						_ = _757_v16
						_757_v16 = _dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Int {
							if (_743_v3).Contains(_this.F16()) {
								return (_743_v3).Get(_this.F16()).(_dafny.Int)
							}
							return _740_v0
						})()), _740_v0)
						var _758_v17 _dafny.MultiSet
						_ = _758_v17
						_758_v17 = _dafny.MultiSetOf(_757_v16)
						_758_v17 = _758_v17
						var _759_v18 _dafny.Array
						_ = _759_v18
						var _nw141 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
						_ = _nw141
						_759_v18 = _nw141
						var _760_v19 *C0
						_ = _760_v19
						var _nw142 *C0 = New_C0_()
						_ = _nw142
						_nw142.Ctor__(_740_v0, _759_v18)
						_760_v19 = _nw142
					}
					var _761_v20 D4
					_ = _761_v20
					_761_v20 = Companion_D4_.Create_DC12_(_this.F16(), Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState), _740_v0)
					var _pat_let_tv19 = _740_v0
					_ = _pat_let_tv19
					var _pat_let_tv20 = _740_v0
					_ = _pat_let_tv20
					var _pat_let_tv21 = _761_v20
					_ = _pat_let_tv21
					var _pat_let_tv22 = globalState
					_ = _pat_let_tv22
					var _762_v21 _dafny.Array
					_ = _762_v21
					var _nwElement0_23 D4 = _761_v20
					_ = _nwElement0_23
					var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(17))
					_ = _nw143
					(_nw143).ArraySet1(_nwElement0_23, 0)
					(_nw143).ArraySet1(Companion_D4_.Create_DC12_(_this.F16(), _this.F16(), Companion_Default___.Fm3(globalState)), 1)
					(_nw143).ArraySet1(Companion_Default___.Fm21(_740_v0, _740_v0, globalState), 2)
					(_nw143).ArraySet1(_761_v20, 3)
					(_nw143).ArraySet1(_761_v20, 4)
					(_nw143).ArraySet1(_761_v20, 5)
					(_nw143).ArraySet1(func(_pat_let16_0 D4) D4 {
						return func(_763_dt__update__tmp_h0 D4) D4 {
							return func(_pat_let17_0 _dafny.Int) D4 {
								return func(_764_dt__update_hcf23_h0 _dafny.Int) D4 {
									return Companion_D4_.Create_DC12_((_763_dt__update__tmp_h0).Dtor_cf21(), (_763_dt__update__tmp_h0).Dtor_cf22(), _764_dt__update_hcf23_h0)
								}(_pat_let17_0)
							}(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv19)).Cardinality()))
						}(_pat_let16_0)
					}(_761_v20), 6)
					(_nw143).ArraySet1(Companion_D4_.Create_DC12_(_this.F16(), _this.F16(), _740_v0), 7)
					(_nw143).ArraySet1(_761_v20, 8)
					(_nw143).ArraySet1(func(_pat_let18_0 D4) D4 {
						return func(_765_dt__update__tmp_h1 D4) D4 {
							return func(_pat_let19_0 _dafny.Int) D4 {
								return func(_766_dt__update_hcf23_h1 _dafny.Int) D4 {
									return Companion_D4_.Create_DC12_((_765_dt__update__tmp_h1).Dtor_cf21(), (_765_dt__update__tmp_h1).Dtor_cf22(), _766_dt__update_hcf23_h1)
								}(_pat_let19_0)
							}(_pat_let_tv20)
						}(_pat_let18_0)
					}(Companion_Default___.Fm21(_740_v0, _740_v0, globalState)), 9)
					(_nw143).ArraySet1(Companion_D4_.Create_DC12_(_this.F16(), _this.F16(), (Companion_Default___.Fm21(_740_v0, _740_v0, globalState)).Dtor_cf23()), 10)
					(_nw143).ArraySet1(func(_pat_let20_0 D4) D4 {
						return func(_767_dt__update__tmp_h2 D4) D4 {
							return func(_pat_let21_0 bool) D4 {
								return func(_768_dt__update_hcf22_h0 bool) D4 {
									return func(_pat_let22_0 bool) D4 {
										return func(_771_dt__update_hcf21_h0 bool) D4 {
											return Companion_D4_.Create_DC12_(_771_dt__update_hcf21_h0, _768_dt__update_hcf22_h0, (_767_dt__update__tmp_h2).Dtor_cf23())
										}(_pat_let22_0)
									}(Companion_Default___.Fm0(_this.F16(), (func(_pat_let23_0 D4) D4 {
										return func(_769_dt__update__tmp_h3 D4) D4 {
											return func(_pat_let24_0 bool) D4 {
												return func(_770_dt__update_hcf21_h1 bool) D4 {
													return Companion_D4_.Create_DC12_(_770_dt__update_hcf21_h1, (_769_dt__update__tmp_h3).Dtor_cf22(), (_769_dt__update__tmp_h3).Dtor_cf23())
												}(_pat_let24_0)
											}(!(_this.F16()))
										}(_pat_let23_0)
									}(_pat_let_tv21)).Dtor_cf21(), _pat_let_tv22))
								}(_pat_let21_0)
							}(!(_this.F16()))
						}(_pat_let20_0)
					}(_761_v20), 11)
					(_nw143).ArraySet1(_761_v20, 12)
					(_nw143).ArraySet1(_761_v20, 13)
					(_nw143).ArraySet1(_761_v20, 14)
					(_nw143).ArraySet1(_761_v20, 15)
					(_nw143).ArraySet1(_761_v20, 16)
					_762_v21 = _nw143
					var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_762_v21), 0))
					_ = _index150
					(_762_v21).ArraySet1(_761_v20, (_index150).Int())
					var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_762_v21), 0))
					_ = _index151
					(_762_v21).ArraySet1(Companion_D4_.Create_DC12_(false, Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState), Companion_Default___.SafeModuloInt(_740_v0, _740_v0)), (_index151).Int())
					var _772_v22 _dafny.Map
					_ = _772_v22
					_772_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_740_v0), _740_v0), !(_this.F16()))
					_772_v22 = (_772_v22).Update(_740_v0, _this.F16())
					var _773_v23 _dafny.Map
					_ = _773_v23
					_773_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if _this.F16() {
							return _this.F16()
						}
						return _this.F16()
					})(), (_740_v0).Cmp((func() _dafny.Int {
						if (_743_v3).Contains(true) {
							return (_743_v3).Get(true).(_dafny.Int)
						}
						return _740_v0
					})()) <= 0)
					var _774_v24 _dafny.Sequence
					_ = _774_v24
					_774_v24 = _dafny.SeqOf((_773_v23).Update(_this.F16(), _this.F16()))
					var _775_v25 D4
					_ = _775_v25
					_775_v25 = Companion_D4_.Create_DC13_((_742_v2).Cardinality(), _740_v0, _740_v0)
					_773_v23 = (_773_v23).Merge((_774_v24).Select((Companion_Default___.SafeIndex((_775_v25).Dtor_cf24(), _dafny.IntOfUint32((_774_v24).Cardinality()))).Uint32()).(_dafny.Map))
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		(_this).F16_set_(_this.F16())
		var _776_v26 _dafny.Array
		_ = _776_v26
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw144
		_776_v26 = _nw144
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
		_ = _index152
		(_776_v26).ArraySet1(_this.F16(), (_index152).Int())
		var _777_v27 D0
		_ = _777_v27
		_777_v27 = Companion_D0_.Create_DC2_(_740_v0, _740_v0, _740_v0, _776_v26)
		var _778_v28 _dafny.MultiSet
		_ = _778_v28
		_778_v28 = _dafny.MultiSetOf(Companion_D0_.Create_DC2_(_740_v0, _740_v0, _740_v0, _776_v26))
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
		_ = _index153
		(_776_v26).ArraySet1(!(_778_v28).Contains(_777_v27), (_index153).Int())
		(globalState).F5 = !((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool))
		if !((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)) || (_this.F16()) {
			_740_v0 = _740_v0
			var _779_v29 _dafny.MultiSet
			_ = _779_v29
			_779_v29 = _dafny.MultiSetOf((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool))
			if !(_779_v29).Contains((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)) {
				var _780_v30 _dafny.CodePoint
				_ = _780_v30
				_780_v30 = _dafny.CodePoint('w')
				(globalState).F15 = (func() _dafny.CodePoint {
					if _this.F16() {
						return _dafny.CodePoint('m')
					}
					return _780_v30
				})()
				var _781_v31 _dafny.Sequence
				_ = _781_v31
				_781_v31 = _dafny.UnicodeSeqOfUtf8Bytes("dita")
				_781_v31 = _dafny.Companion_Sequence_.Update((_this).F24(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_740_v0), _dafny.IntOfUint32(((_this).F24()).Cardinality()))).Uint32(), _780_v30)
				var _782_v32 _dafny.Array
				_ = _782_v32
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
				_ = _nw145
				_782_v32 = _nw145
				var _783_v33 D2
				_ = _783_v33
				_783_v33 = Companion_D2_.Create_DC8_(_743_v3, _782_v32, false, (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool), _740_v0)
				var _784_v34 D2
				_ = _784_v34
				_784_v34 = Companion_D2_.Create_DC8_(_743_v3, (_783_v33).Dtor_cf14(), _this.F16(), (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool), (_743_v3).Cardinality())
				var _785_v35 D6
				_ = _785_v35
				_785_v35 = Companion_D6_.Create_DC17_(_784_v34)
				var _786_v36 _dafny.MultiSet
				_ = _786_v36
				_786_v36 = _dafny.MultiSetOf(_785_v35)
				var _787_v37 _dafny.Sequence
				_ = _787_v37
				_787_v37 = _dafny.SeqOf(_786_v36)
				var _788_v38 _dafny.Sequence
				_ = _788_v38
				_788_v38 = _787_v37
				var _rhs168 _dafny.Int = Companion_Default___.Fm3(globalState)
				_ = _rhs168
				var _rhs169 _dafny.Sequence = (_787_v37)
				_ = _rhs169
				r0 = _rhs168
				_787_v37 = _rhs169
				(globalState).F2 = _this.F16()
				var _789_v39 _dafny.Array
				_ = _789_v39
				var _nwElement0_24 _dafny.CodePoint = _780_v30
				_ = _nwElement0_24
				var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(24))
				_ = _nw146
				(_nw146).ArraySet1CodePoint(_nwElement0_24, 0)
				(_nw146).ArraySet1CodePoint(_780_v30, 1)
				(_nw146).ArraySet1CodePoint(_780_v30, 2)
				(_nw146).ArraySet1CodePoint(_780_v30, 3)
				(_nw146).ArraySet1CodePoint(_780_v30, 4)
				(_nw146).ArraySet1CodePoint(_780_v30, 5)
				(_nw146).ArraySet1CodePoint(_780_v30, 6)
				(_nw146).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool) {
						return (_781_v31).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32((_781_v31).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
					return _780_v30
				})(), 7)
				(_nw146).ArraySet1CodePoint(_780_v30, 8)
				(_nw146).ArraySet1CodePoint(_780_v30, 9)
				(_nw146).ArraySet1CodePoint(_780_v30, 10)
				(_nw146).ArraySet1CodePoint(_780_v30, 11)
				(_nw146).ArraySet1CodePoint(_780_v30, 12)
				(_nw146).ArraySet1CodePoint(_dafny.CodePoint('g'), 13)
				(_nw146).ArraySet1CodePoint(_780_v30, 14)
				(_nw146).ArraySet1CodePoint(_780_v30, 15)
				(_nw146).ArraySet1CodePoint(_780_v30, 16)
				(_nw146).ArraySet1CodePoint(_780_v30, 17)
				(_nw146).ArraySet1CodePoint(Companion_Default___.Fm2(globalState), 18)
				(_nw146).ArraySet1CodePoint(_780_v30, 19)
				(_nw146).ArraySet1CodePoint(_780_v30, 20)
				(_nw146).ArraySet1CodePoint(_780_v30, 21)
				(_nw146).ArraySet1CodePoint(_780_v30, 22)
				(_nw146).ArraySet1CodePoint(_780_v30, 23)
				_789_v39 = _nw146
				_789_v39 = _789_v39
			} else {
				var _790_v40 _dafny.Array
				_ = _790_v40
				var _nw147 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(20))
				_ = _nw147
				_790_v40 = _nw147
				var _791_v41 D6
				_ = _791_v41
				_791_v41 = Companion_D6_.Create_DC16_(true, _this.F16())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_790_v40), 0))
				_ = _index154
				(_790_v40).ArraySet1(_791_v41, (_index154).Int())
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_790_v40), 0))
				_ = _index155
				(_790_v40).ArraySet1(Companion_D6_.Create_DC16_((_740_v0).Cmp(_dafny.IntOfUint32(((_this).F24()).Cardinality())) < 0, _this.F16()), (_index155).Int())
				var _792_v42 _dafny.Array
				_ = _792_v42
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_21
				var _nw148 _dafny.Array
				_ = _nw148
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw148 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Sequence = func(_793_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17())
					}
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw148 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw148).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw148).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_792_v42 = _nw148
				var _794_v43 _dafny.Map
				_ = _794_v43
				_794_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_741_v1, _740_v0)
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_792_v42), 0))
				_ = _index156
				(_792_v42).ArraySet1(_dafny.SeqOf(_740_v0, (_794_v43).Cardinality(), _740_v0), (_index156).Int())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_792_v42), 0))
				_ = _index157
				(_792_v42).ArraySet1(Companion_Default___.Fm17(globalState), (_index157).Int())
				var _795_v44 _dafny.Set
				_ = _795_v44
				_795_v44 = _dafny.SetOf((_dafny.Zero).Minus(_740_v0), _740_v0, _740_v0, (_dafny.Zero).Minus(_740_v0), _740_v0)
				(_this).F16_set_(!(!((_795_v44).IsProperSubsetOf(_795_v44))))
				var _796_v45 _dafny.Array
				_ = _796_v45
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
				_ = _nw149
				_796_v45 = _nw149
				var _797_v46 _dafny.Set
				_ = _797_v46
				_797_v46 = _dafny.SetOf(_741_v1, _741_v1)
				var _798_v47 _dafny.Map
				_ = _798_v47
				_798_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v45, _797_v46)
				_798_v47 = (_798_v47).Update(_796_v45, _dafny.SetOf(_741_v1, (Companion_Default___.Fm22(_this.F16(), _dafny.MultiSetFromSeq((_792_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_792_v42), 0))).Int()).(_dafny.Sequence)), (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool), globalState)).Dtor_cf34()))
				var _799_v48 _dafny.Array
				_ = _799_v48
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw150
				_799_v48 = _nw150
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_799_v48), 0))
				_ = _index158
				(_799_v48).ArraySet1(_740_v0, (_index158).Int())
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_799_v48), 0))
				_ = _index159
				var _rhs170 _dafny.Int = _740_v0
				_ = _rhs170
				var _rhs171 _dafny.MultiSet = _779_v29
				_ = _rhs171
				var _rhs172 _dafny.Int = _dafny.IntOfInt64(-58)
				_ = _rhs172
				var _rhs173 D0 = _777_v27
				_ = _rhs173
				var _lhs142 _dafny.Array = _799_v48
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_799_v48), 0))
				_ = _lhs143
				r1 = _rhs170
				_779_v29 = _rhs171
				(_lhs142).ArraySet1(_rhs172, (_lhs143).Int())
				_777_v27 = _rhs173
			}
			if !((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)) {
				var _800_v49 _dafny.Array
				_ = _800_v49
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw151
				_800_v49 = _nw151
				var _rhs174 _dafny.Int = Companion_Default___.Fm3(globalState)
				_ = _rhs174
				var _rhs175 _dafny.Array = _800_v49
				_ = _rhs175
				var _rhs176 _dafny.Int = _740_v0
				_ = _rhs176
				var _rhs177 _dafny.Int = ((func() _dafny.Int {
					if _this.F16() {
						return (_dafny.Zero).Minus(_740_v0)
					}
					return _740_v0
				})()).Plus(_740_v0)
				_ = _rhs177
				_740_v0 = _rhs174
				_800_v49 = _rhs175
				r0 = _rhs176
				_740_v0 = _rhs177
				r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F24(), (_this).F24()), (Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F24(), (_this).F24())).Cardinality()))).Uint32(), _dafny.CodePoint('x')), (_this).F24())).Cardinality())
				r0 = _dafny.IntOfInt64(960)
				var _801_v50 _dafny.Array
				_ = _801_v50
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw152
				_801_v50 = _nw152
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_801_v50), 0))
				_ = _index160
				(_801_v50).ArraySet1(_800_v49, (_index160).Int())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_801_v50), 0))
				_ = _index161
				(_801_v50).ArraySet1(_800_v49, (_index161).Int())
				var _802_v51 _dafny.Map
				_ = _802_v51
				_802_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool), (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool))
				var _803_v52 _dafny.Map
				_ = _803_v52
				_803_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v51, (func() _dafny.Array {
					if _this.F16() {
						return _800_v49
					}
					return _dafny.ArrayCastTo((_801_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(834), _dafny.ArrayLen((_801_v50), 0))).Int()))
				})())
				var _804_v53 _dafny.Array
				_ = _804_v53
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_22
				var _nw153 _dafny.Array
				_ = _nw153
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw153 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Map = (func(_805_v51 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_806_i2 _dafny.Int) _dafny.Map {
							return (_805_v51).Update(_this.F16(), _this.F16())
						}
					})(_802_v51)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw153 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw153).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw153).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_804_v53 = _nw153
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_804_v53), 0))
				_ = _index162
				(_804_v53).ArraySet1(_802_v51, (_index162).Int())
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_804_v53), 0))
				_ = _index163
				var _rhs178 _dafny.Map = _803_v52
				_ = _rhs178
				var _rhs179 _dafny.Map = ((_802_v51).Merge(_802_v51)).Merge((Companion_Default___.Fm23(_dafny.MultiSetOf((_dafny.Zero).Minus(_740_v0), _740_v0, _740_v0), !(!((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool))), globalState)).Update(_this.F16(), (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)))
				_ = _rhs179
				var _lhs144 _dafny.Array = _804_v53
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_804_v53), 0))
				_ = _lhs145
				_803_v52 = _rhs178
				(_lhs144).ArraySet1(_rhs179, (_lhs145).Int())
			} else {
				var _rhs180 _dafny.Map = func() _dafny.Map {
					var _coll43 = _dafny.NewMapBuilder()
					_ = _coll43
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(147), _dafny.IntOfInt64(327))); ; {
						_compr_43, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _807_v54 _dafny.Int
						_807_v54 = interface{}(_compr_43).(_dafny.Int)
						if ((_dafny.IntOfInt64(147)).Cmp(_807_v54) <= 0) && ((_807_v54).Cmp(_dafny.IntOfInt64(327)) < 0) {
							_coll43.Add(Companion_Default___.SafeModuloInt(_807_v54, (_779_v29).Cardinality()), _dafny.IntOfUint32(((_this).F24()).Cardinality()))
						}
					}
					return _coll43.ToMap()
				}()
				_ = _rhs180
				var _rhs181 _dafny.Int = _740_v0
				_ = _rhs181
				_742_v2 = _rhs180
				r1 = _rhs181
				var _808_v55 _dafny.Sequence
				_ = _808_v55
				_808_v55 = _dafny.UnicodeSeqOfUtf8Bytes("he")
				_808_v55 = _808_v55
				var _809_v56 _dafny.Array
				_ = _809_v56
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_23
				var _nw154 _dafny.Array
				_ = _nw154
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw154 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_810_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_811_i3 _dafny.Int) _dafny.Int {
							return (_811_i3).Plus(_810_v0)
						}
					})(_740_v0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw154 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw154).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw154).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_809_v56 = _nw154
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_24
				var _nw155 _dafny.Array
				_ = _nw155
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw155 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = (func(_812_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_813_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_813_i4, _812_v0)
						}
					})(_740_v0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw155 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw155).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw155).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_809_v56 = _nw155
				r1 = _740_v0
				_743_v3 = (_743_v3).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)), _740_v0))
			}
			var _814_v57 D8
			_ = _814_v57
			_814_v57 = Companion_D8_.Create_DC23_()
			var _815_v58 _dafny.Set
			_ = _815_v58
			_815_v58 = _dafny.SetOf(_814_v57, _814_v57, _814_v57)
			var _816_v59 _dafny.CodePoint
			_ = _816_v59
			_816_v59 = _dafny.CodePoint('s')
			var _817_v60 _dafny.Map
			_ = _817_v60
			_817_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_816_v59, _dafny.IntOfUint32(((_this).F24()).Cardinality()))
			var _818_v61 _dafny.MultiSet
			_ = _818_v61
			_818_v61 = _dafny.MultiSetOf((_815_v58).Union(_815_v58), _dafny.SetOf(_814_v57, Companion_Default___.Fm24(((_this).F17()).Select((Companion_Default___.SafeIndex((_817_v60).Cardinality(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), _816_v59, globalState)), _815_v58, _815_v58, _dafny.SetOf(_814_v57))
			var _819_v62 _dafny.Sequence
			_ = _819_v62
			_819_v62 = _dafny.SeqOf((_dafny.IntOfUint32(((_this).F17()).Cardinality())).Minus(_740_v0))
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
			_ = _index164
			var _rhs182 _dafny.MultiSet = Companion_Default___.Fm25(true, _740_v0, globalState)
			_ = _rhs182
			var _rhs183 _dafny.Int = _dafny.IntOfInt64(675)
			_ = _rhs183
			var _rhs184 bool = (func() bool {
				if (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool) {
					return !(_this.F16())
				}
				return _this.F16()
			})()
			_ = _rhs184
			var _rhs185 _dafny.Sequence = _819_v62
			_ = _rhs185
			var _rhs186 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F16(), _this.F16()), _741_v1), _741_v1)
			_ = _rhs186
			var _lhs146 _dafny.Array = _776_v26
			_ = _lhs146
			var _lhs147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
			_ = _lhs147
			_818_v61 = _rhs182
			_740_v0 = _rhs183
			(_lhs146).ArraySet1(_rhs184, (_lhs147).Int())
			_819_v62 = _rhs185
			_741_v1 = _rhs186
			var _820_v63 _dafny.Array
			_ = _820_v63
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_25
			var _nw156 _dafny.Array
			_ = _nw156
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw156 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = (func(_821_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_822_i5 _dafny.Int) _dafny.Int {
						return (_822_i5).Times(_821_v0)
					}
				})(_740_v0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw156 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw156).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw156).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_820_v63 = _nw156
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_820_v63), 0))
			_ = _index165
			(_820_v63).ArraySet1(_740_v0, (_index165).Int())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_820_v63), 0))
			_ = _index166
			(_820_v63).ArraySet1((_dafny.Zero).Minus(_740_v0), (_index166).Int())
		} else {
			var _823_v64 _dafny.Set
			_ = _823_v64
			_823_v64 = _dafny.SetOf(_dafny.SeqOf(Companion_Default___.Fm3(globalState)), _dafny.SeqOf((_dafny.Zero).Minus(_740_v0)), _dafny.SeqOf(_740_v0))
			if (_823_v64).IsProperSubsetOf(_823_v64) {
				var _824_v65 _dafny.Array
				_ = _824_v65
				var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw157
				_824_v65 = _nw157
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_824_v65), 0))
				_ = _index167
				(_824_v65).ArraySet1((_dafny.Zero).Minus(_740_v0), (_index167).Int())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_824_v65), 0))
				_ = _index168
				(_824_v65).ArraySet1((func() _dafny.Int {
					if (_743_v3).Contains((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)) {
						return (_743_v3).Get((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)).(_dafny.Int)
					}
					return (_740_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("isodoen")).Cardinality()))
				})(), (_index168).Int())
				(globalState).F2 = (((_824_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_824_v65), 0))).Int()).(_dafny.Int)).Cmp(_740_v0) > 0) || (false)
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
				_ = _index169
				(_776_v26).ArraySet1(!((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)), (_index169).Int())
				_743_v3 = (_743_v3).Update(_this.F16(), (_824_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_824_v65), 0))).Int()).(_dafny.Int))
				Companion_Default___.M0((_824_v65).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_824_v65), 0))).Int()).(_dafny.Int), (_this).F24(), (func(_pat_let25_0 D8) D8 {
					return func(_825_dt__update__tmp_h5 D8) D8 {
						return func(_pat_let26_0 _dafny.Int) D8 {
							return func(_826_dt__update_hcf39_h0 _dafny.Int) D8 {
								return Companion_D8_.Create_DC21_((_825_dt__update__tmp_h5).Dtor_cf35(), (_825_dt__update__tmp_h5).Dtor_cf36(), (_825_dt__update__tmp_h5).Dtor_cf37(), (_825_dt__update__tmp_h5).Dtor_cf38(), _826_dt__update_hcf39_h0)
							}(_pat_let26_0)
						}(_dafny.IntOfInt64(856))
					}(_pat_let25_0)
				}(Companion_Default___.Fm26(globalState))).Dtor_cf35(), _740_v0, globalState)
			} else {
				var _827_v66 _dafny.Map
				_ = _827_v66
				_827_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16())
				var _828_v67 *C1
				_ = _828_v67
				var _nw158 *C1 = New_C1_()
				_ = _nw158
				_nw158.Ctor__((_740_v0).Times(_740_v0), _740_v0, Companion_Default___.Fm0((func() bool {
					if (_827_v66).Contains((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)) {
						return (_827_v66).Get((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)).(bool)
					}
					return _this.F16()
				})(), false, globalState), (_this).F17(), _740_v0)
				_828_v67 = _nw158
				_740_v0 = _dafny.IntOfInt64(853)
				r0 = (_828_v67).F30()
				var _829_v68 _dafny.Array
				_ = _829_v68
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_26
				var _nw159 _dafny.Array
				_ = _nw159
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw159 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Int = (func(_830_v67 *C1) func(_dafny.Int) _dafny.Int {
						return func(_831_i6 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_831_i6, (_830_v67).F30())
						}
					})(_828_v67)
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw159 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw159).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw159).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_829_v68 = _nw159
				var _832_v69 D6
				_ = _832_v69
				_832_v69 = Companion_D6_.Create_DC18_(_829_v68)
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
				_ = _index170
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
				_ = _index171
				var _rhs187 _dafny.Array = (_832_v69).Dtor_cf32()
				_ = _rhs187
				var _rhs188 bool = !(!((_dafny.IntOfUint32((_741_v1).Cardinality())).Cmp(_740_v0) <= 0))
				_ = _rhs188
				var _rhs189 bool = (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)
				_ = _rhs189
				var _rhs190 bool = (_741_v1).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32((_741_v1).Cardinality()))).Uint32()).(bool)
				_ = _rhs190
				var _rhs191 _dafny.Array = _829_v68
				_ = _rhs191
				var _lhs148 _dafny.Array = _776_v26
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
				_ = _lhs149
				var _lhs150 *GlobalState = globalState
				_ = _lhs150
				var _lhs151 _dafny.Array = _776_v26
				_ = _lhs151
				var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))
				_ = _lhs152
				_829_v68 = _rhs187
				(_lhs148).ArraySet1(_rhs188, (_lhs149).Int())
				_lhs150.F5 = _rhs189
				(_lhs151).ArraySet1(_rhs190, (_lhs152).Int())
				_829_v68 = _rhs191
				var _833_v70 *C2
				_ = _833_v70
				var _nw160 *C2 = New_C2_()
				_ = _nw160
				_nw160.Ctor__(_740_v0)
				_833_v70 = _nw160
				_833_v70 = (func() *C2 {
					if _this.F16() {
						return _833_v70
					}
					return _833_v70
				})()
			}
			var _834_v71 _dafny.Sequence
			_ = _834_v71
			_834_v71 = _dafny.UnicodeSeqOfUtf8Bytes("wmq")
			var _835_v72 _dafny.CodePoint
			_ = _835_v72
			_835_v72 = _dafny.CodePoint('i')
			_834_v71 = _dafny.Companion_Sequence_.Update(_834_v71, (Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32((_834_v71).Cardinality()))).Uint32(), _835_v72)
			_740_v0 = (_740_v0).Minus(_740_v0)
			var _836_v73 _dafny.Map
			_ = _836_v73
			_836_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool))
			var _837_v74 _dafny.Map
			_ = _837_v74
			_837_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool), _776_v26)
			var _838_v75 _dafny.MultiSet
			_ = _838_v75
			_838_v75 = _dafny.MultiSetOf((_837_v74).Cardinality(), _740_v0)
			(globalState).F5 = (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_740_v0, _dafny.IntOfUint32((_741_v1).Cardinality())), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_740_v0), _dafny.IntOfUint32((_dafny.SeqOf(_740_v0, _dafny.IntOfUint32((_741_v1).Cardinality()))).Cardinality()))).Uint32(), (_836_v73).Cardinality())).Cardinality()), _740_v0, (_dafny.Zero).Minus(_740_v0), _740_v0)).IsProperSubsetOf((func() _dafny.MultiSet {
				if (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool) {
					return _838_v75
				}
				return _dafny.MultiSetOf(_dafny.IntOfUint32((_834_v71).Cardinality()))
			})())
			(globalState).F2 = (_776_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_776_v26), 0))).Int()).(bool)
		}
		r0 = _740_v0
		r1 = _740_v0
		var _839_v76 _dafny.Sequence
		_ = _839_v76
		_839_v76 = _dafny.SeqOf(Companion_Default___.Fm17(globalState), _dafny.SeqOf(_740_v0, _740_v0), (_this).F17(), (_this).F17())
		var _840_v77 _dafny.Map
		_ = _840_v77
		_840_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_839_v76).Select((Companion_Default___.SafeIndex(_740_v0, _dafny.IntOfUint32((_839_v76).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), false)
		r2 = _840_v77
		return r0, r1, r2
	}
}
func (_this *C3) F24() _dafny.Sequence {
	{
		return _this._f24
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	F23  bool
	_f22 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this.F23 = false
	_this._f22 = _dafny.Zero
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f22 _dafny.Int, f23 bool) {
	{
		(_this)._f22 = f22
		(_this).F23 = f23
	}
}
func (_this *C4) Fm9(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(533), _dafny.IntOfInt64(795)), _dafny.SeqOf((_dafny.MultiSetOf((_this).F22(), (_this).F22())).Cardinality())), _dafny.SeqOf(_dafny.SeqOf((_this).F22(), (_this).F22(), (_this).F22(), (_this).F22(), (_this).F22()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_841_i0 _dafny.Int) _dafny.Int {
			return (_this).F22()
		})))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-128))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}(func(_842_i1 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((_this).F22())
		})))
	}
}
func (_this *C4) Fm10(p0 D1, p1 bool, globalState *GlobalState) _dafny.Map {
	{
		if (Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F23, false)).Cardinality()), false)).Dtor_cf11() {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _dafny.MultiSetOf((_this).F22(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D2_.Create_DC7_(_dafny.SetOf(_this.F23)))).Cardinality()), (_this).F22())).Cardinality()))
		} else if _this.F23 {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false, _this.F23, _this.F23, _this.F23, _this.F23)).Cardinality()), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _this.F23)).Cardinality(), _dafny.IntOfInt64(702), (_dafny.Zero).Minus((_this).F22()), (_this).F22(), _dafny.IntOfInt64(62)))
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(508), _dafny.MultiSetOf((_this).F22()))
		}
	}
}
func (_this *C4) M6(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Set, bool, _dafny.Int) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _843_v0 _dafny.CodePoint
		_ = _843_v0
		_843_v0 = _dafny.CodePoint('j')
		var _844_v1 _dafny.Sequence
		_ = _844_v1
		_844_v1 = _dafny.UnicodeSeqOfUtf8Bytes("hvsewr")
		(globalState).F2 = !_dafny.Companion_Sequence_.Contains(_844_v1, _843_v0)
		if _this.F23 {
			var _845_v2 _dafny.MultiSet
			_ = _845_v2
			_845_v2 = _dafny.MultiSetOf(p1)
			var _846_v3 D1
			_ = _846_v3
			_846_v3 = Companion_D1_.Create_DC3_(_845_v2)
			var _source9 D1 = _846_v3
			_ = _source9
			if _source9.Is_DC4() {
				var _847___mcc_h0 bool = _source9.Get_().(D1_DC4).Cf7
				_ = _847___mcc_h0
				var _848___mcc_h1 _dafny.Int = _source9.Get_().(D1_DC4).Cf8
				_ = _848___mcc_h1
				var _849___mcc_h2 _dafny.Map = _source9.Get_().(D1_DC4).Cf9
				_ = _849___mcc_h2
				var _850_cf9 _dafny.Map = _849___mcc_h2
				_ = _850_cf9
				var _851_cf8 _dafny.Int = _848___mcc_h1
				_ = _851_cf8
				var _852_cf7 bool = _847___mcc_h0
				_ = _852_cf7
				var _853_v4 _dafny.Array
				_ = _853_v4
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_27
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) _dafny.Int = (func(_854_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_855_i0 _dafny.Int) _dafny.Int {
							return (_855_i0).Plus(_854_p2)
						}
					})(p2)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw161 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw161).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw161).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_853_v4 = _nw161
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))
				_ = _index172
				(_853_v4).ArraySet1((func() _dafny.Int {
					if (_this).Fm9(p2, globalState) {
						return _dafny.IntOfInt64(368)
					}
					return Companion_Default___.Fm3(globalState)
				})(), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))
				_ = _index173
				var _rhs192 _dafny.Int = _851_cf8
				_ = _rhs192
				var _rhs193 _dafny.Int = p2
				_ = _rhs193
				var _rhs194 _dafny.Int = p2
				_ = _rhs194
				var _rhs195 _dafny.Int = _851_cf8
				_ = _rhs195
				var _rhs196 _dafny.Int = (_dafny.IntOfInt64(-974)).Times((_this).F22())
				_ = _rhs196
				var _lhs153 _dafny.Array = _853_v4
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))
				_ = _lhs154
				(_lhs153).ArraySet1(_rhs192, (_lhs154).Int())
				r2 = _rhs193
				_851_cf8 = _rhs194
				r2 = _rhs195
				r2 = _rhs196
				_850_cf9 = (_850_cf9).Update(_852_cf7, _this.F23)
				var _856_v5 _dafny.Map
				_ = _856_v5
				_856_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(772), (_this).F22())
				var _857_v6 _dafny.Sequence
				_ = _857_v6
				_857_v6 = _dafny.SeqOf(_852_cf7)
				_856_v5 = (_856_v5).Update((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_857_v6).Cardinality()), p0)).Cardinality())), _851_cf8)
				var _858_v7 _dafny.Map
				_ = _858_v7
				_858_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v0, _this.F23)
				var _859_v8 D9
				_ = _859_v8
				_859_v8 = Companion_D9_.Create_DC25_(_858_v7)
				var _860_v9 _dafny.Map
				_ = _860_v9
				_860_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_859_v8).Dtor_cf44()).Cardinality(), false)
				var _861_v10 _dafny.Array
				_ = _861_v10
				var _nwElement0_25 bool = (func() bool {
					if (_860_v9).Contains((_853_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))).Int()).(_dafny.Int)) {
						return (_860_v9).Get((_853_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))).Int()).(_dafny.Int)).(bool)
					}
					return p1
				})()
				_ = _nwElement0_25
				var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(16))
				_ = _nw162
				(_nw162).ArraySet1(_nwElement0_25, 0)
				(_nw162).ArraySet1(!(_this.F23), 1)
				(_nw162).ArraySet1(_852_cf7, 2)
				(_nw162).ArraySet1(p0, 3)
				(_nw162).ArraySet1(p0, 4)
				(_nw162).ArraySet1(_this.F23, 5)
				(_nw162).ArraySet1(_852_cf7, 6)
				(_nw162).ArraySet1(_this.F23, 7)
				(_nw162).ArraySet1(p0, 8)
				(_nw162).ArraySet1(p1, 9)
				(_nw162).ArraySet1(p1, 10)
				(_nw162).ArraySet1(false, 11)
				(_nw162).ArraySet1(false, 12)
				(_nw162).ArraySet1(p1, 13)
				(_nw162).ArraySet1(_this.F23, 14)
				(_nw162).ArraySet1(p1, 15)
				_861_v10 = _nw162
				var _862_v11 *C0
				_ = _862_v11
				var _nw163 *C0 = New_C0_()
				_ = _nw163
				_nw163.Ctor__((_853_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_853_v4), 0))).Int()).(_dafny.Int), _861_v10)
				_862_v11 = _nw163
			} else if _source9.Is_DC5() {
				var _863___mcc_h3 _dafny.Int = _source9.Get_().(D1_DC5).Cf10
				_ = _863___mcc_h3
				var _864___mcc_h4 bool = _source9.Get_().(D1_DC5).Cf11
				_ = _864___mcc_h4
				var _865_cf11 bool = _864___mcc_h4
				_ = _865_cf11
				var _866_cf10 _dafny.Int = _863___mcc_h3
				_ = _866_cf10
				var _867_v12 _dafny.Sequence
				_ = _867_v12
				_867_v12 = _dafny.SeqOf(false, _this.F23)
				var _868_v13 _dafny.Map
				_ = _868_v13
				_868_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
				var _869_v14 _dafny.Set
				_ = _869_v14
				_869_v14 = _dafny.SetOf((_this).F22())
				var _870_v15 _dafny.Map
				_ = _870_v15
				_870_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_869_v14, (func() bool {
					if (_868_v13).Contains(p1) {
						return (_868_v13).Get(p1).(bool)
					}
					return !(false)
				})())
				var _871_v16 _dafny.Sequence
				_ = _871_v16
				_871_v16 = _dafny.SeqOf(p2, (_870_v15).Cardinality())
				var _872_v17 *C1
				_ = _872_v17
				var _nw164 *C1 = New_C1_()
				_ = _nw164
				_nw164.Ctor__(_dafny.IntOfUint32((_867_v12).Cardinality()), (_868_v13).Cardinality(), !(p1), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm3(globalState)), _871_v16), (_this).F22())
				_872_v17 = _nw164
				(globalState).F2 = !(_865_cf11)
				var _873_v18 _dafny.Map
				_ = _873_v18
				_873_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_872_v17).F30(), p0)
				var _874_v19 _dafny.Array
				_ = _874_v19
				var _nwElement0_26 bool = p1
				_ = _nwElement0_26
				var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(8))
				_ = _nw165
				(_nw165).ArraySet1(_nwElement0_26, 0)
				(_nw165).ArraySet1(_865_cf11, 1)
				(_nw165).ArraySet1(_865_cf11, 2)
				(_nw165).ArraySet1(p0, 3)
				(_nw165).ArraySet1(false, 4)
				(_nw165).ArraySet1(_865_cf11, 5)
				(_nw165).ArraySet1((_867_v12).Select((Companion_Default___.SafeIndex((_872_v17).F30(), _dafny.IntOfUint32((_867_v12).Cardinality()))).Uint32()).(bool), 6)
				(_nw165).ArraySet1((func() bool {
					if (_873_v18).Contains(_dafny.IntOfUint32((_867_v12).Cardinality())) {
						return (_873_v18).Get(_dafny.IntOfUint32((_867_v12).Cardinality())).(bool)
					}
					return Companion_Default___.Fm0(p0, p1, globalState)
				})(), 7)
				_874_v19 = _nw165
				_874_v19 = _874_v19
				var _875_v20 *C0
				_ = _875_v20
				var _nw166 *C0 = New_C0_()
				_ = _nw166
				_nw166.Ctor__((_872_v17).Fm13(_this.F23, globalState), _874_v19)
				_875_v20 = _nw166
			} else if _source9.Is_DC6() {
				r2 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(447))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_876_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality())
				var _877_v21 _dafny.Array
				_ = _877_v21
				var _nw167 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw167
				_877_v21 = _nw167
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_877_v21), 0))
				_ = _index174
				(_877_v21).ArraySet1((p0) == (p0), (_index174).Int())
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(330), _dafny.ArrayLen((_877_v21), 0))
				_ = _index175
				(_877_v21).ArraySet1(!(true), (_index175).Int())
				(globalState).F2 = p0
				var _878_v22 _dafny.Array
				_ = _878_v22
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_28
				var _nw168 _dafny.Array
				_ = _nw168
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw168 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Map = (func(_879_p2 _dafny.Int) func(_dafny.Int) _dafny.Map {
						return func(_880_i2 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _879_p2)
						}
					})(p2)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw168 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw168).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw168).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_878_v22 = _nw168
				var _881_v23 _dafny.Map
				_ = _881_v23
				_881_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v0, _878_v22)
				_881_v23 = (_881_v23).Update(Companion_Default___.Fm2(globalState), _878_v22)
			} else {
				var _882___mcc_h5 _dafny.MultiSet = _source9.Get_().(D1_DC3).Cf6
				_ = _882___mcc_h5
				var _883_cf6 _dafny.MultiSet = _882___mcc_h5
				_ = _883_cf6
				var _884_v24 _dafny.MultiSet
				_ = _884_v24
				_884_v24 = _dafny.MultiSetOf(_844_v1, _844_v1)
				var _885_v25 _dafny.Sequence
				_ = _885_v25
				_885_v25 = _dafny.SeqOf((_884_v24).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_886_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_887_i3 _dafny.Int) _dafny.CodePoint {
						return _886_v0
					}
				})(_843_v0)))).Cardinality()))
				var _888_v26 D3
				_ = _888_v26
				_888_v26 = Companion_D3_.Create_DC10_(p1)
				var _889_v27 _dafny.Set
				_ = _889_v27
				_889_v27 = _dafny.SetOf(false, !((_888_v26).Dtor_cf19()))
				var _890_v28 *C1
				_ = _890_v28
				var _nw169 *C1 = New_C1_()
				_ = _nw169
				_nw169.Ctor__((_dafny.Zero).Minus((_this).F22()), p2, p0, _885_v25, (_889_v27).Cardinality())
				_890_v28 = _nw169
				var _891_v29 _dafny.MultiSet
				_ = _891_v29
				_891_v29 = _dafny.MultiSetOf(_890_v28)
				r2 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_891_v29).Contains(_890_v28) {
						return (_891_v29).Multiplicity(_890_v28)
					}
					return (_890_v28).Fm14(p2, (_890_v28).F30(), _843_v0, globalState)
				})(), Companion_Default___.SafeModuloInt((_890_v28).F29(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_890_v28).F29(), p2)).Update(p2, (_890_v28).F29())).Cardinality()))
				var _892_v30 D8
				_ = _892_v30
				_892_v30 = Companion_D8_.Create_DC23_()
				var _893_v31 D8
				_ = _893_v31
				_893_v31 = Companion_D8_.Create_DC24_((func() D8 {
					if p0 {
						return _892_v30
					}
					return _892_v30
				})())
				_893_v31 = _893_v31
				var _894_v32 _dafny.Map
				_ = _894_v32
				_894_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_844_v1, (_this).Fm9(p2, globalState))
				_894_v32 = (_894_v32).Update(_dafny.UnicodeSeqOfUtf8Bytes("weptditry"), _this.F23)
				var _895_v33 _dafny.Map
				_ = _895_v33
				_895_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), p0)
				var _896_v34 _dafny.Map
				_ = _896_v34
				_896_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_895_v33).Update(p1, p1)).Cardinality(), p2)
				(globalState).F5 = ((_890_v28).Fm15(_dafny.Companion_Sequence_.Update(_885_v25, (Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_885_v25).Cardinality()))).Uint32(), (_890_v28).F30()), (_889_v27).Cardinality(), (_890_v28).F30(), _843_v0, globalState)).Cmp((func() _dafny.Int {
					if (_896_v34).Contains(p2) {
						return (_896_v34).Get(p2).(_dafny.Int)
					}
					return p2
				})()) > 0
			}
			var _897_v35 _dafny.Array
			_ = _897_v35
			var _nw170 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw170
			_897_v35 = _nw170
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
			_ = _index176
			(_897_v35).ArraySet1(p0, (_index176).Int())
			var _898_v36 _dafny.Array
			_ = _898_v36
			var _nwElement0_27 _dafny.Array = _897_v35
			_ = _nwElement0_27
			var _nw171 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(24))
			_ = _nw171
			(_nw171).ArraySet1(_nwElement0_27, 0)
			(_nw171).ArraySet1(_897_v35, 1)
			(_nw171).ArraySet1(_897_v35, 2)
			(_nw171).ArraySet1(_897_v35, 3)
			(_nw171).ArraySet1(_897_v35, 4)
			(_nw171).ArraySet1(_897_v35, 5)
			(_nw171).ArraySet1(_897_v35, 6)
			(_nw171).ArraySet1(_897_v35, 7)
			(_nw171).ArraySet1(_897_v35, 8)
			(_nw171).ArraySet1(_897_v35, 9)
			(_nw171).ArraySet1(_897_v35, 10)
			(_nw171).ArraySet1(_897_v35, 11)
			(_nw171).ArraySet1(_897_v35, 12)
			(_nw171).ArraySet1(_897_v35, 13)
			(_nw171).ArraySet1(_897_v35, 14)
			(_nw171).ArraySet1((func() _dafny.Array {
				if p0 {
					return _897_v35
				}
				return _897_v35
			})(), 15)
			(_nw171).ArraySet1(_897_v35, 16)
			(_nw171).ArraySet1(_897_v35, 17)
			(_nw171).ArraySet1(_897_v35, 18)
			(_nw171).ArraySet1(_897_v35, 19)
			(_nw171).ArraySet1(_897_v35, 20)
			(_nw171).ArraySet1(_897_v35, 21)
			(_nw171).ArraySet1((func() _dafny.Array {
				if false {
					return _897_v35
				}
				return _897_v35
			})(), 22)
			(_nw171).ArraySet1(_897_v35, 23)
			_898_v36 = _nw171
			var _899_v37 _dafny.Sequence
			_ = _899_v37
			_899_v37 = _dafny.SeqOf(p2)
			var _900_v38 _dafny.Map
			_ = _900_v38
			_900_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(569))
			var _901_v39 _dafny.Map
			_ = _901_v39
			_901_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_899_v37).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_899_v37).Cardinality()))).Uint32()).(_dafny.Int))
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
			_ = _index177
			var _rhs197 bool = (((_dafny.Zero).Minus((_this).F22())).Cmp(p2) > 0) || (p1)
			_ = _rhs197
			var _rhs198 _dafny.Int = p2
			_ = _rhs198
			var _rhs199 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_899_v37, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F22(), (_dafny.SetOf(_dafny.IntOfInt64(996), _dafny.IntOfUint32((_844_v1).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_899_v37).Cardinality()))).Uint32(), (func() _dafny.Int {
				if (_900_v38).Contains(p2) {
					return (_900_v38).Get(p2).(_dafny.Int)
				}
				return Companion_Default___.Fm3(globalState)
			})())).Cardinality())
			_ = _rhs199
			var _rhs200 bool = (Companion_Default___.Fm3(globalState)).Cmp((_901_v39).Cardinality()) < 0
			_ = _rhs200
			var _rhs201 _dafny.Array = _898_v36
			_ = _rhs201
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			var _lhs156 _dafny.Array = _897_v35
			_ = _lhs156
			var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
			_ = _lhs157
			_lhs155.F2 = _rhs197
			r2 = _rhs198
			r2 = _rhs199
			(_lhs156).ArraySet1(_rhs200, (_lhs157).Int())
			_898_v36 = _rhs201
			if _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_844_v1, Companion_Default___.Fm1(globalState)), _844_v1) {
				var _902_v40 *C3
				_ = _902_v40
				var _nw172 *C3 = New_C3_()
				_ = _nw172
				_nw172.Ctor__(_844_v1, _this.F23, Companion_Default___.Fm17(globalState))
				_902_v40 = _nw172
				var _903_v41 _dafny.Map
				_ = _903_v41
				_903_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v0, _this.F23)
				var _904_v42 _dafny.Map
				_ = _904_v42
				_904_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC25_(_903_v41), (p2).Minus((_this).F22()))
				var _905_v44 _dafny.Map
				_ = _905_v44
				_905_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v0, p2)
				var _906_v45 D9
				_ = _906_v45
				_906_v45 = Companion_D9_.Create_DC25_(func() _dafny.Map {
					var _coll44 = _dafny.NewMapBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate((_905_v44).Keys().Elements()); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _907_v43 _dafny.CodePoint
						_907_v43 = interface{}(_compr_44).(_dafny.CodePoint)
						if (_905_v44).Contains(_907_v43) {
							_coll44.Add(_907_v43, p0)
						}
					}
					return _coll44.ToMap()
				}())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _index178
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _index179
				var _rhs202 bool = p0
				_ = _rhs202
				var _rhs203 bool = p0
				_ = _rhs203
				var _rhs204 _dafny.Map = (_904_v42).Update(_906_v45, _dafny.IntOfInt64(539))
				_ = _rhs204
				var _rhs205 _dafny.Int = (_this).F22()
				_ = _rhs205
				var _lhs158 _dafny.Array = _897_v35
				_ = _lhs158
				var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _lhs159
				var _lhs160 _dafny.Array = _897_v35
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _lhs161
				(_lhs158).ArraySet1(_rhs202, (_lhs159).Int())
				(_lhs160).ArraySet1(_rhs203, (_lhs161).Int())
				_904_v42 = _rhs204
				r2 = _rhs205
				var _908_v46 _dafny.Array
				_ = _908_v46
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw173
				_908_v46 = _nw173
				var _909_v47 _dafny.Array
				_ = _909_v47
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
				_ = _nw174
				_909_v47 = _nw174
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_908_v46), 0))
				_ = _index180
				(_908_v46).ArraySet1(_909_v47, (_index180).Int())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_908_v46), 0))
				_ = _index181
				(_908_v46).ArraySet1(_909_v47, (_index181).Int())
				var _910_v48 _dafny.Sequence
				_ = _910_v48
				_910_v48 = _dafny.SeqOf((_897_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))).Int()).(bool), p0, false, (_dafny.SetOf((_this).F22())).Contains((_this).F22()))
				var _911_v51 _dafny.Array
				_ = _911_v51
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_29
				var _nw175 _dafny.Array
				_ = _nw175
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw175 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Set = (func(_912_v37 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
						return func(_913_i4 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf(func() _dafny.Set {
								var _coll45 = _dafny.NewBuilder()
								_ = _coll45
								for _iter49 := _dafny.Iterate((_912_v37).Elements()); ; {
									_compr_45, _ok49 := _iter49()
									if !_ok49 {
										break
									}
									var _914_v49 _dafny.Int
									_914_v49 = interface{}(_compr_45).(_dafny.Int)
									if _dafny.Companion_Sequence_.Contains(_912_v37, _914_v49) {
										_coll45.Add((_914_v49).Plus((_dafny.Zero).Minus(_dafny.IntOfInt64(-821))))
									}
								}
								return _coll45.ToSet()
							}(), func() _dafny.Set {
								var _coll46 = _dafny.NewBuilder()
								_ = _coll46
								for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(375), _dafny.IntOfInt64(298))); ; {
									_compr_46, _ok50 := _iter50()
									if !_ok50 {
										break
									}
									var _915_v50 _dafny.Int
									_915_v50 = interface{}(_compr_46).(_dafny.Int)
									if ((_dafny.IntOfInt64(375)).Cmp(_915_v50) <= 0) && ((_915_v50).Cmp(_dafny.IntOfInt64(298)) < 0) {
										_coll46.Add((_915_v50).Minus((_this).F22()))
									}
								}
								return _coll46.ToSet()
							}())).Difference(_dafny.SetOf(_dafny.SetOf((_this).F22())))
						}
					})(_899_v37)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw175 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw175).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw175).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_911_v51 = _nw175
				var _916_v52 _dafny.Set
				_ = _916_v52
				_916_v52 = _dafny.SetOf((_dafny.Zero).Minus(p2))
				var _917_v53 _dafny.Set
				_ = _917_v53
				_917_v53 = _dafny.SetOf(_916_v52, _916_v52)
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_911_v51), 0))
				_ = _index182
				(_911_v51).ArraySet1(_917_v53, (_index182).Int())
				var _918_v54 _dafny.Map
				_ = _918_v54
				_918_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_844_v1, _dafny.CodePoint('g'))
				var _919_v56 _dafny.Sequence
				_ = _919_v56
				_919_v56 = _dafny.SeqOf(_844_v1)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_908_v46), 0))
				_ = _index183
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_911_v51), 0))
				_ = _index184
				var _rhs206 _dafny.Array = _909_v47
				_ = _rhs206
				var _rhs207 _dafny.Sequence = _dafny.SeqOf(false)
				_ = _rhs207
				var _rhs208 _dafny.Set = _917_v53
				_ = _rhs208
				var _rhs209 _dafny.Int = _dafny.IntOfInt64(-22)
				_ = _rhs209
				var _rhs210 _dafny.Map = (func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter51 := _dafny.Iterate((_919_v56).Elements()); ; {
						_compr_47, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _920_v55 _dafny.Sequence
						_920_v55 = interface{}(_compr_47).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_919_v56, _920_v55) {
							_coll47.Add(_920_v55, _dafny.CodePoint('e'))
						}
					}
					return _coll47.ToMap()
				}()).Merge(_918_v54)
				_ = _rhs210
				var _lhs162 _dafny.Array = _908_v46
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_908_v46), 0))
				_ = _lhs163
				var _lhs164 _dafny.Array = _911_v51
				_ = _lhs164
				var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_911_v51), 0))
				_ = _lhs165
				(_lhs162).ArraySet1(_rhs206, (_lhs163).Int())
				_910_v48 = _rhs207
				(_lhs164).ArraySet1(_rhs208, (_lhs165).Int())
				r2 = _rhs209
				_918_v54 = _rhs210
				var _921_v57 _dafny.Sequence
				_ = _921_v57
				_921_v57 = _dafny.SeqOf(_909_v47)
				var _922_v58 _dafny.Array
				_ = _922_v58
				var _nwElement0_28 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_921_v57, _921_v57)
				_ = _nwElement0_28
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(16))
				_ = _nw176
				(_nw176).ArraySet1(_nwElement0_28, 0)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_921_v57, _dafny.SeqOf(_909_v47)), 1)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_921_v57, _921_v57), 2)
				(_nw176).ArraySet1(_dafny.SeqOf(_909_v47), 3)
				(_nw176).ArraySet1(_921_v57, 4)
				(_nw176).ArraySet1((func() _dafny.Sequence {
					if p1 {
						return _921_v57
					}
					return _921_v57
				})(), 5)
				(_nw176).ArraySet1(_921_v57, 6)
				(_nw176).ArraySet1(_921_v57, 7)
				(_nw176).ArraySet1(_921_v57, 8)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_921_v57, _921_v57), 9)
				(_nw176).ArraySet1(_921_v57, 10)
				(_nw176).ArraySet1(_921_v57, 11)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_909_v47), _921_v57), 12)
				(_nw176).ArraySet1(_dafny.SeqOf(_909_v47), 13)
				(_nw176).ArraySet1(_dafny.SeqOf(_dafny.ArrayCastTo((_908_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_908_v46), 0))).Int())), _909_v47), 14)
				(_nw176).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_921_v57, _921_v57), 15)
				_922_v58 = _nw176
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_922_v58), 0))
				_ = _index185
				(_922_v58).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_921_v57, _921_v57), (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_922_v58), 0))
				_ = _index186
				(_922_v58).ArraySet1(_921_v57, (_index186).Int())
			} else {
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _index187
				var _rhs211 _dafny.Int = (p2).Minus(p2)
				_ = _rhs211
				var _rhs212 bool = !(!_dafny.Companion_Sequence_.Contains(_899_v37, ((_this).F22()).Minus((_this).F22())))
				_ = _rhs212
				var _lhs166 _dafny.Array = _897_v35
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))
				_ = _lhs167
				r2 = _rhs211
				(_lhs166).ArraySet1(_rhs212, (_lhs167).Int())
				(globalState).F2 = ((_897_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))).Int()).(bool)) && ((_dafny.MultiSetOf(p0)).IsSubsetOf(_845_v2))
				var _923_v59 _dafny.Map
				_ = _923_v59
				_923_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _this.F23)
				var _924_v60 _dafny.Map
				_ = _924_v60
				_924_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_897_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen((_897_v35), 0))).Int()).(bool)) && (_this.F23), Companion_D1_.Create_DC4_(_this.F23, p2, _923_v59))
				_924_v60 = (_924_v60).Merge(_924_v60)
				_897_v35 = _897_v35
				var _925_v61 _dafny.Set
				_ = _925_v61
				_925_v61 = _dafny.SetOf(_843_v0)
				var _926_v62 _dafny.Map
				_ = _926_v62
				_926_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_843_v0, false)
				var _927_v64 _dafny.Map
				_ = _927_v64
				_927_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_925_v61).Difference(func() _dafny.Set {
					var _coll48 = _dafny.NewBuilder()
					_ = _coll48
					for _iter52 := _dafny.Iterate((_926_v62).Keys().Elements()); ; {
						_compr_48, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _928_v63 _dafny.CodePoint
						_928_v63 = interface{}(_compr_48).(_dafny.CodePoint)
						if (_926_v62).Contains(_928_v63) {
							_coll48.Add(_928_v63)
						}
					}
					return _coll48.ToSet()
				}()), _897_v35)
				var _rhs213 _dafny.Int = (p2).Minus(p2)
				_ = _rhs213
				var _rhs214 _dafny.Int = p2
				_ = _rhs214
				var _rhs215 _dafny.Int = (_this).F22()
				_ = _rhs215
				var _rhs216 _dafny.Array = (func() _dafny.Array {
					if (_927_v64).Contains(Companion_Default___.Fm27(globalState)) {
						return (_927_v64).Get(Companion_Default___.Fm27(globalState)).(_dafny.Array)
					}
					return _897_v35
				})()
				_ = _rhs216
				var _rhs217 _dafny.Int = (_this).F22()
				_ = _rhs217
				r2 = _rhs213
				r2 = _rhs214
				r2 = _rhs215
				_897_v35 = _rhs216
				r2 = _rhs217
			}
			r2 = Companion_Default___.Fm3(globalState)
			var _929_v65 _dafny.Sequence
			_ = _929_v65
			_929_v65 = _dafny.SeqOf(false)
			_929_v65 = _dafny.Companion_Sequence_.Concatenate(_929_v65, _929_v65)
		} else {
			var _930_v66 _dafny.Set
			_ = _930_v66
			_930_v66 = _dafny.SetOf(p2)
			var _931_v67 _dafny.Set
			_ = _931_v67
			_931_v67 = _dafny.SetOf((_930_v66).Cardinality(), p2)
			var _932_v68 _dafny.Set
			_ = _932_v68
			_932_v68 = _dafny.SetOf(_931_v67, _931_v67, _930_v66)
			var _933_v69 _dafny.Sequence
			_ = _933_v69
			_933_v69 = _dafny.SeqOf(_dafny.SetOf(_930_v66))
			var _934_v70 _dafny.Sequence
			_ = _934_v70
			_934_v70 = _dafny.SeqOf(_dafny.IntOfInt64(-236), p2)
			var _935_v71 *C1
			_ = _935_v71
			var _nw177 *C1 = New_C1_()
			_ = _nw177
			_nw177.Ctor__(p2, p2, _this.F23, _934_v70, _dafny.IntOfUint32((_844_v1).Cardinality()))
			_935_v71 = _nw177
			var _936_v72 _dafny.Map
			_ = _936_v72
			_936_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_935_v71, (_this).F22())
			var _937_v73 _dafny.Map
			_ = _937_v73
			_937_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_936_v72, _dafny.SetOf(_dafny.SetOf((_935_v71).F30(), (_935_v71).F30(), (_935_v71).F30())))
			var _938_v76 _dafny.Sequence
			_ = _938_v76
			_938_v76 = _dafny.SeqOf(_930_v66)
			var _939_v79 _dafny.Array
			_ = _939_v79
			var _nwElement0_29 _dafny.Set = _932_v68
			_ = _nwElement0_29
			var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(24))
			_ = _nw178
			(_nw178).ArraySet1(_nwElement0_29, 0)
			(_nw178).ArraySet1(_dafny.SetOf(_931_v67), 1)
			(_nw178).ArraySet1((_933_v69).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_933_v69).Cardinality()))).Uint32()).(_dafny.Set), 2)
			(_nw178).ArraySet1((_932_v68).Union(_932_v68), 3)
			(_nw178).ArraySet1((_932_v68).Union(_932_v68), 4)
			(_nw178).ArraySet1(Companion_Default___.Fm28(p1, (_this).F22(), p0, _843_v0, globalState), 5)
			(_nw178).ArraySet1(_932_v68, 6)
			(_nw178).ArraySet1(_932_v68, 7)
			(_nw178).ArraySet1((func() _dafny.Set {
				if (_937_v73).Contains(_936_v72) {
					return (_937_v73).Get(_936_v72).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll49 = _dafny.NewBuilder()
					_ = _coll49
					for _iter53 := _dafny.Iterate((Companion_Default___.Fm29(p1, p2, p0, globalState)).Elements()); ; {
						_compr_49, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _940_v74 _dafny.Set
						_940_v74 = interface{}(_compr_49).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm29(p1, p2, p0, globalState), _940_v74) {
							_coll49.Add(_940_v74)
						}
					}
					return _coll49.ToSet()
				}()
			})(), 8)
			(_nw178).ArraySet1(func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter54 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v67, (_935_v71).F30())).Keys().Elements()); ; {
					_compr_50, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _941_v75 _dafny.Set
					_941_v75 = interface{}(_compr_50).(_dafny.Set)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v67, (_935_v71).F30())).Contains(_941_v75) {
						_coll50.Add(_941_v75)
					}
				}
				return _coll50.ToSet()
			}(), 9)
			(_nw178).ArraySet1(_932_v68, 10)
			(_nw178).ArraySet1((func() _dafny.Set {
				if p0 {
					return _932_v68
				}
				return _932_v68
			})(), 11)
			(_nw178).ArraySet1(_932_v68, 12)
			(_nw178).ArraySet1((_932_v68).Difference(_dafny.SetOf(_931_v67, _dafny.SetOf(p2), _931_v67, _930_v66, _930_v66)), 13)
			(_nw178).ArraySet1(_932_v68, 14)
			(_nw178).ArraySet1(_932_v68, 15)
			(_nw178).ArraySet1(_932_v68, 16)
			(_nw178).ArraySet1(func() _dafny.Set {
				var _coll51 = _dafny.NewBuilder()
				_ = _coll51
				for _iter55 := _dafny.Iterate((_938_v76).Elements()); ; {
					_compr_51, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _942_v77 _dafny.Set
					_942_v77 = interface{}(_compr_51).(_dafny.Set)
					if _dafny.Companion_Sequence_.Contains(_938_v76, _942_v77) {
						_coll51.Add(_942_v77)
					}
				}
				return _coll51.ToSet()
			}(), 17)
			(_nw178).ArraySet1(_932_v68, 18)
			(_nw178).ArraySet1(_932_v68, 19)
			(_nw178).ArraySet1(_932_v68, 20)
			(_nw178).ArraySet1(_932_v68, 21)
			(_nw178).ArraySet1(func() _dafny.Set {
				var _coll52 = _dafny.NewBuilder()
				_ = _coll52
				for _iter56 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v67, p0)).Keys().Elements()); ; {
					_compr_52, _ok56 := _iter56()
					if !_ok56 {
						break
					}
					var _943_v78 _dafny.Set
					_943_v78 = interface{}(_compr_52).(_dafny.Set)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v67, p0)).Contains(_943_v78) {
						_coll52.Add(_943_v78)
					}
				}
				return _coll52.ToSet()
			}(), 22)
			(_nw178).ArraySet1((_dafny.SetOf(_930_v66)).Difference(_932_v68), 23)
			_939_v79 = _nw178
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_939_v79), 0))
			_ = _index188
			(_939_v79).ArraySet1(_dafny.SetOf(_931_v67), (_index188).Int())
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_939_v79), 0))
			_ = _index189
			(_939_v79).ArraySet1(Companion_Default___.Fm28(_this.F23, (_935_v71).F29(), !(p0) || (p1), _843_v0, globalState), (_index189).Int())
			var _944_v80 _dafny.Map
			_ = _944_v80
			_944_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F22())
			var _945_v81 _dafny.Array
			_ = _945_v81
			var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
			_ = _nw179
			_945_v81 = _nw179
			var _946_v82 D2
			_ = _946_v82
			_946_v82 = Companion_D2_.Create_DC8_(_944_v80, _945_v81, _this.F23, p1, (_935_v71).F30())
			var _947_v83 D6
			_ = _947_v83
			_947_v83 = Companion_D6_.Create_DC17_(_946_v82)
			var _948_v84 _dafny.MultiSet
			_ = _948_v84
			_948_v84 = _dafny.MultiSetOf(_947_v83, Companion_D6_.Create_DC17_(_946_v82))
			var _949_v85 _dafny.Sequence
			_ = _949_v85
			_949_v85 = _dafny.SeqOf(_948_v84)
			var _950_v86 _dafny.Sequence
			_ = _950_v86
			_950_v86 = _949_v85
			var _951_v87 _dafny.Map
			_ = _951_v87
			_951_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_950_v86, (func() _dafny.CodePoint {
				if p0 {
					return _843_v0
				}
				return _843_v0
			})())
			var _rhs218 _dafny.Int = (_934_v70).Select((Companion_Default___.SafeIndex((_this).F22(), _dafny.IntOfUint32((_934_v70).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs218
			var _rhs219 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_951_v87).Contains(_950_v86) {
					return (_951_v87).Get(_950_v86).(_dafny.CodePoint)
				}
				return _843_v0
			})()
			_ = _rhs219
			var _lhs168 *GlobalState = globalState
			_ = _lhs168
			r2 = _rhs218
			_lhs168.F15 = _rhs219
			_843_v0 = _dafny.CodePoint('t')
			_843_v0 = _843_v0
			var _952_v88 *C1
			_ = _952_v88
			var _nw180 *C1 = New_C1_()
			_ = _nw180
			_nw180.Ctor__((_935_v71).F29(), p2, p1, _934_v70, (_this).F22())
			_952_v88 = _nw180
		}
		(globalState).F2 = !(_this.F23)
		var _953_v89 _dafny.Map
		_ = _953_v89
		_953_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _843_v0)
		var _954_v90 _dafny.Set
		_ = _954_v90
		_954_v90 = _dafny.SetOf(p0, !(Companion_Default___.Fm0(false, false, globalState)))
		var _955_v91 _dafny.Array
		_ = _955_v91
		var _nwElement0_30 _dafny.Int = Companion_Default___.SafeDivisionInt((_953_v89).Cardinality(), (_this).F22())
		_ = _nwElement0_30
		var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(3))
		_ = _nw181
		(_nw181).ArraySet1(_nwElement0_30, 0)
		(_nw181).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 1)
		(_nw181).ArraySet1((_954_v90).Cardinality(), 2)
		_955_v91 = _nw181
		var _956_v92 _dafny.Array
		_ = _956_v92
		var _nw182 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw182
		_956_v92 = _nw182
		var _957_v93 _dafny.Map
		_ = _957_v93
		_957_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_956_v92, Companion_Default___.Fm3(globalState))
		var _958_v94 _dafny.Sequence
		_ = _958_v94
		_958_v94 = _dafny.SeqOf(p1, p1)
		var _959_v95 _dafny.Map
		_ = _959_v95
		_959_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(647), _dafny.SeqOf((_958_v94).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_958_v94).Cardinality()))).Uint32()).(bool), p1))
		var _960_v96 _dafny.Map
		_ = _960_v96
		_960_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((func() _dafny.Int {
			if (_957_v93).Contains(_956_v92) {
				return (_957_v93).Get(_956_v92).(_dafny.Int)
			}
			return (_959_v95).Cardinality()
		})()), (_this).F22())
		var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))
		_ = _index190
		(_955_v91).ArraySet1((_960_v96).Cardinality(), (_index190).Int())
		var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))
		_ = _index191
		(_955_v91).ArraySet1(p2, (_index191).Int())
		var _961_i5 _dafny.Int
		_ = _961_i5
		_961_i5 = _dafny.Zero
		{
			for p1 {
				{
					if (_961_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_961_i5 = (_961_i5).Plus(_dafny.One)
					_956_v92 = _956_v92
					var _962_v97 _dafny.MultiSet
					_ = _962_v97
					_962_v97 = _dafny.MultiSetOf(_this.F23)
					_962_v97 = (_962_v97).Union(_962_v97)
					var _963_v98 _dafny.MultiSet
					_ = _963_v98
					_963_v98 = _dafny.MultiSetOf((_955_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_958_v94).Cardinality()))
					var _964_v99 _dafny.Map
					_ = _964_v99
					_964_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), p1)
					var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_956_v92), 0))
					_ = _index192
					(_956_v92).ArraySet1(((_963_v98).Update((_964_v99).Cardinality(), Companion_Default___.Abs((_dafny.Zero).Minus((_955_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))).Int()).(_dafny.Int))))).IsDisjointFrom(_963_v98), (_index192).Int())
					var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_956_v92), 0))
					_ = _index193
					(_956_v92).ArraySet1(((_dafny.IntOfInt64(391)).Minus((_955_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))).Int()).(_dafny.Int))).Cmp(Companion_Default___.SafeModuloInt((_this).F22(), (_this).F22())) < 0, (_index193).Int())
					var _965_v100 D4
					_ = _965_v100
					_965_v100 = Companion_D4_.Create_DC13_(p2, (_dafny.Zero).Minus((_955_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))).Int()).(_dafny.Int)), (_this).F22())
					var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_955_v91), 0))
					_ = _index194
					(_955_v91).ArraySet1((_965_v100).Dtor_cf25(), (_index194).Int())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_956_v92), 0))); ; {
			_guard_loop_4, _ok57 := _iter57()
			if !_ok57 {
				break
			}
			var _966_i6 _dafny.Int
			_966_i6 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_966_i6).Sign() != -1) && ((_966_i6).Cmp(_dafny.ArrayLen((_956_v92), 0)) < 0)) {
				(_956_v92).ArraySet1(p0, (_966_i6).Int())
			}
		}
		var _967_v101 _dafny.MultiSet
		_ = _967_v101
		_967_v101 = _dafny.MultiSetOf((_this).F22(), _dafny.IntOfInt64(628), _dafny.IntOfInt64(437), (_dafny.Zero).Minus(p2), _dafny.IntOfUint32((_844_v1).Cardinality()))
		var _968_v102 _dafny.Set
		_ = _968_v102
		_968_v102 = _dafny.SetOf(_967_v101)
		r0 = (func() _dafny.Set {
			var _coll53 = _dafny.NewBuilder()
			_ = _coll53
			for _iter58 := _dafny.Iterate((_968_v102).Elements()); ; {
				_compr_53, _ok58 := _iter58()
				if !_ok58 {
					break
				}
				var _969_v103 _dafny.MultiSet
				_969_v103 = interface{}(_compr_53).(_dafny.MultiSet)
				if (_968_v102).Contains(_969_v103) {
					_coll53.Add(_969_v103)
				}
			}
			return _coll53.ToSet()
		}()).Union((func() _dafny.Set {
			if _this.F23 {
				return _968_v102
			}
			return _968_v102
		})())
		r1 = p1
		r2 = (_this).F22()
		return r0, r1, r2
	}
}
func (_this *C4) M7(globalState *GlobalState) (_dafny.Int, _dafny.Sequence, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _970_v0 _dafny.Map
		_ = _970_v0
		_970_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _this.F23)
		var _971_v1 D1
		_ = _971_v1
		_971_v1 = Companion_D1_.Create_DC4_(_this.F23, (_this).F22(), _970_v0)
		var _972_v2 _dafny.Sequence
		_ = _972_v2
		_972_v2 = _dafny.SeqOf(_971_v1, _971_v1, _971_v1)
		_972_v2 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_971_v1), _dafny.Companion_Sequence_.Concatenate(_972_v2, _972_v2))
		var _973_v3 D8
		_ = _973_v3
		_973_v3 = Companion_D8_.Create_DC21_((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F22())), _this.F23, _this.F23, !(_this.F23), (_this).F22())
		var _974_i0 _dafny.Int
		_ = _974_i0
		_974_i0 = _dafny.Zero
		{
			for func(_source10 D8) bool {
				if _source10.Is_DC21() {
					var _979___mcc_h0 _dafny.Int = _source10.Get_().(D8_DC21).Cf35
					_ = _979___mcc_h0
					var _980___mcc_h1 bool = _source10.Get_().(D8_DC21).Cf36
					_ = _980___mcc_h1
					var _981___mcc_h2 bool = _source10.Get_().(D8_DC21).Cf37
					_ = _981___mcc_h2
					var _982___mcc_h3 bool = _source10.Get_().(D8_DC21).Cf38
					_ = _982___mcc_h3
					var _983___mcc_h4 _dafny.Int = _source10.Get_().(D8_DC21).Cf39
					_ = _983___mcc_h4
					var _984_cf39 _dafny.Int = _983___mcc_h4
					_ = _984_cf39
					var _985_cf38 bool = _982___mcc_h3
					_ = _985_cf38
					var _986_cf37 bool = _981___mcc_h2
					_ = _986_cf37
					var _987_cf36 bool = _980___mcc_h1
					_ = _987_cf36
					var _988_cf35 _dafny.Int = _979___mcc_h0
					_ = _988_cf35
					return _986_cf37
				} else if _source10.Is_DC22() {
					var _989___mcc_h5 bool = _source10.Get_().(D8_DC22).Cf40
					_ = _989___mcc_h5
					var _990___mcc_h6 T0 = _source10.Get_().(D8_DC22).Cf41
					_ = _990___mcc_h6
					var _991___mcc_h7 bool = _source10.Get_().(D8_DC22).Cf42
					_ = _991___mcc_h7
					var _992_cf42 bool = _991___mcc_h7
					_ = _992_cf42
					var _993_cf41 T0 = _990___mcc_h6
					_ = _993_cf41
					var _994_cf40 bool = _989___mcc_h5
					_ = _994_cf40
					return !(_993_cf41.F16())
				} else if _source10.Is_DC23() {
					return _this.F23
				} else if _source10.Is_DC20() {
					var _995___mcc_h8 _dafny.Sequence = _source10.Get_().(D8_DC20).Cf34
					_ = _995___mcc_h8
					var _996_cf34 _dafny.Sequence = _995___mcc_h8
					_ = _996_cf34
					return _this.F23
				} else {
					var _997___mcc_h9 D8 = _source10.Get_().(D8_DC24).Cf43
					_ = _997___mcc_h9
					var _998_cf43 D8 = _997___mcc_h9
					_ = _998_cf43
					return _dafny.Companion_Sequence_.IsProperPrefixOf((Companion_D3_.Create_DC9_(_dafny.SeqOf((_this).F22()))).Dtor_cf18(), _dafny.SeqOf((_this).F22(), (_dafny.Zero).Minus((_this).F22())))
				}
			}(_973_v3) {
				{
					if (_974_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_974_i0 = (_974_i0).Plus(_dafny.One)
					(_this).F23 = ((_this).F22()).Cmp((_this).F22()) >= 0
					var _975_v4 _dafny.Array
					_ = _975_v4
					var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
					_ = _nw183
					_975_v4 = _nw183
					var _976_v5 _dafny.Map
					_ = _976_v5
					_976_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_this).F22())
					var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_975_v4), 0))
					_ = _index195
					(_975_v4).ArraySet1(((_976_v5).Cardinality()).Minus((_this).F22()), (_index195).Int())
					var _977_v6 _dafny.Array
					_ = _977_v6
					var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
					_ = _nw184
					_977_v6 = _nw184
					var _978_v7 _dafny.Array
					_ = _978_v7
					var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
					_ = _nw185
					_978_v7 = _nw185
					var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_977_v6), 0))
					_ = _index196
					(_977_v6).ArraySet1(_978_v7, (_index196).Int())
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_975_v4), 0))
					_ = _index197
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_977_v6), 0))
					_ = _index198
					var _rhs220 _dafny.Int = (_dafny.Zero).Minus((_this).F22())
					_ = _rhs220
					var _rhs221 bool = _this.F23
					_ = _rhs221
					var _rhs222 bool = _this.F23
					_ = _rhs222
					var _rhs223 bool = ((_this).F22()).Cmp((_this).F22()) == 0
					_ = _rhs223
					var _rhs224 _dafny.Array = _978_v7
					_ = _rhs224
					var _lhs169 _dafny.Array = _975_v4
					_ = _lhs169
					var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_975_v4), 0))
					_ = _lhs170
					var _lhs171 *GlobalState = globalState
					_ = _lhs171
					var _lhs172 *GlobalState = globalState
					_ = _lhs172
					var _lhs173 *GlobalState = globalState
					_ = _lhs173
					var _lhs174 _dafny.Array = _977_v6
					_ = _lhs174
					var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_977_v6), 0))
					_ = _lhs175
					(_lhs169).ArraySet1(_rhs220, (_lhs170).Int())
					_lhs171.F5 = _rhs221
					_lhs172.F2 = _rhs222
					_lhs173.F2 = _rhs223
					(_lhs174).ArraySet1(_rhs224, (_lhs175).Int())
					(globalState).F2 = (Companion_Default___.SafeDivisionInt((_this).F22(), (_this).F22())).Cmp((_975_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_975_v4), 0))).Int()).(_dafny.Int)) >= 0
					(globalState).F15 = _dafny.CodePoint('f')
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		r0 = Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, false)).Cardinality(), ((_this).F22()).Times(_dafny.IntOfInt64(858)))
		var _hi6 _dafny.Int = (_this).F22()
		_ = _hi6
		for _999_i1 := (_this).F22(); _999_i1.Cmp(_hi6) < 0; _999_i1 = _999_i1.Plus(_dafny.One) {
			r2 = !(_this.F23) || (_this.F23)
			r2 = _this.F23
			var _1000_v8 _dafny.Sequence
			_ = _1000_v8
			_1000_v8 = _dafny.SeqOf(_this.F23)
			_1000_v8 = _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf(_1000_v8, _dafny.Companion_Sequence_.Update(_1000_v8, (Companion_Default___.SafeIndex(Companion_Default___.Fm3(globalState), _dafny.IntOfUint32((_1000_v8).Cardinality()))).Uint32(), _this.F23)))
			var _1001_v9 _dafny.Array
			_ = _1001_v9
			var _nwElement0_31 D8 = _973_v3
			_ = _nwElement0_31
			var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(7))
			_ = _nw186
			(_nw186).ArraySet1(_nwElement0_31, 0)
			(_nw186).ArraySet1(_973_v3, 1)
			(_nw186).ArraySet1(Companion_D8_.Create_DC21_((_this).F22(), _this.F23, _this.F23, true, _999_i1), 2)
			(_nw186).ArraySet1(_973_v3, 3)
			(_nw186).ArraySet1(_973_v3, 4)
			(_nw186).ArraySet1(_973_v3, 5)
			(_nw186).ArraySet1(_973_v3, 6)
			_1001_v9 = _nw186
			var _1002_v10 *C2
			_ = _1002_v10
			var _nw187 *C2 = New_C2_()
			_ = _nw187
			_nw187.Ctor__((_this).F22())
			_1002_v10 = _nw187
			var _1003_v11 _dafny.Map
			_ = _1003_v11
			_1003_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _1002_v10)
			var _pat_let_tv23 = _1003_v11
			_ = _pat_let_tv23
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1001_v9), 0))
			_ = _index199
			(_1001_v9).ArraySet1((func() D8 {
				if true {
					return func(_pat_let27_0 D8) D8 {
						return func(_1004_dt__update__tmp_h0 D8) D8 {
							return func(_pat_let28_0 _dafny.Int) D8 {
								return func(_1005_dt__update_hcf35_h0 _dafny.Int) D8 {
									return Companion_D8_.Create_DC21_(_1005_dt__update_hcf35_h0, (_1004_dt__update__tmp_h0).Dtor_cf36(), (_1004_dt__update__tmp_h0).Dtor_cf37(), (_1004_dt__update__tmp_h0).Dtor_cf38(), (_1004_dt__update__tmp_h0).Dtor_cf39())
								}(_pat_let28_0)
							}((_pat_let_tv23).Cardinality())
						}(_pat_let27_0)
					}(_973_v3)
				}
				return Companion_Default___.Fm26(globalState)
			})(), (_index199).Int())
			var _1006_v12 _dafny.Array
			_ = _1006_v12
			var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw188
			_1006_v12 = _nw188
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1006_v12), 0))
			_ = _index200
			(_1006_v12).ArraySet1((func() _dafny.Int {
				if _this.F23 {
					return (_1002_v10).F25()
				}
				return (_dafny.Zero).Minus((_1002_v10).F25())
			})(), (_index200).Int())
			var _1007_v13 _dafny.MultiSet
			_ = _1007_v13
			_1007_v13 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F23), _1000_v8))
			var _1008_v14 _dafny.Map
			_ = _1008_v14
			_1008_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _dafny.IntOfInt64(491))
			var _1009_v15 _dafny.Map
			_ = _1009_v15
			_1009_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_999_i1, _this.F23)
			var _1010_v16 _dafny.Array
			_ = _1010_v16
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw189
			_1010_v16 = _nw189
			var _1011_v17 D2
			_ = _1011_v17
			_1011_v17 = Companion_D2_.Create_DC8_(_1008_v14, _1010_v16, _this.F23, _this.F23, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qjqdbdskh")).Cardinality())))
			var _1012_v18 _dafny.Array
			_ = _1012_v18
			var _nwElement0_32 bool = _this.F23
			_ = _nwElement0_32
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_32, 0)
			(_nw190).ArraySet1(_this.F23, 1)
			(_nw190).ArraySet1(_this.F23, 2)
			(_nw190).ArraySet1(_this.F23, 3)
			(_nw190).ArraySet1(_this.F23, 4)
			(_nw190).ArraySet1(!(_this.F23), 5)
			(_nw190).ArraySet1(!(_1008_v14).Equals(_1008_v14), 6)
			(_nw190).ArraySet1(false, 7)
			(_nw190).ArraySet1((func() bool {
				if (_1000_v8).Select((Companion_Default___.SafeIndex((_1002_v10).F25(), _dafny.IntOfUint32((_1000_v8).Cardinality()))).Uint32()).(bool) {
					return _this.F23
				}
				return (func() bool {
					if (_1009_v15).Contains(_999_i1) {
						return (_1009_v15).Get(_999_i1).(bool)
					}
					return (_this).Fm9(_999_i1, globalState)
				})()
			})(), 8)
			(_nw190).ArraySet1(_this.F23, 9)
			(_nw190).ArraySet1(false, 10)
			(_nw190).ArraySet1(_this.F23, 11)
			(_nw190).ArraySet1((_dafny.IntOfInt64(499)).Cmp(_999_i1) != 0, 12)
			(_nw190).ArraySet1(_this.F23, 13)
			(_nw190).ArraySet1(_this.F23, 14)
			(_nw190).ArraySet1(!(_this.F23) || (_this.F23), 15)
			(_nw190).ArraySet1(_this.F23, 16)
			(_nw190).ArraySet1(true, 17)
			(_nw190).ArraySet1(!((_999_i1).Cmp(_dafny.IntOfInt64(59)) > 0), 18)
			(_nw190).ArraySet1(_this.F23, 19)
			(_nw190).ArraySet1((func() bool {
				if _this.F23 {
					return (_1011_v17).Dtor_cf15()
				}
				return _this.F23
			})(), 20)
			(_nw190).ArraySet1(_this.F23, 21)
			(_nw190).ArraySet1(!(!(_this.F23)), 22)
			(_nw190).ArraySet1(_this.F23, 23)
			(_nw190).ArraySet1(_this.F23, 24)
			_1012_v18 = _nw190
			var _1013_v19 _dafny.MultiSet
			_ = _1013_v19
			_1013_v19 = _dafny.MultiSetOf(_999_i1)
			var _1014_v20 _dafny.Map
			_ = _1014_v20
			_1014_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v18, _1012_v18)
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1001_v9), 0))
			_ = _index201
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1006_v12), 0))
			_ = _index202
			var _rhs225 D8 = Companion_D8_.Create_DC21_((_dafny.Zero).Minus((func() _dafny.Int {
				if true {
					return (_dafny.Zero).Minus((_this).F22())
				}
				return (_dafny.Zero).Minus(_999_i1)
			})()), !(_1013_v19).Equals(_1013_v19), _this.F23, _this.F23, _999_i1)
			_ = _rhs225
			var _rhs226 _dafny.Int = (_this).F22()
			_ = _rhs226
			var _rhs227 _dafny.Int = _999_i1
			_ = _rhs227
			var _rhs228 _dafny.MultiSet = _1007_v13
			_ = _rhs228
			var _rhs229 _dafny.Array = (func() _dafny.Array {
				if (_1014_v20).Contains(_1012_v18) {
					return (_1014_v20).Get(_1012_v18).(_dafny.Array)
				}
				return _1012_v18
			})()
			_ = _rhs229
			var _lhs176 _dafny.Array = _1001_v9
			_ = _lhs176
			var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(383), _dafny.ArrayLen((_1001_v9), 0))
			_ = _lhs177
			var _lhs178 _dafny.Array = _1006_v12
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1006_v12), 0))
			_ = _lhs179
			(_lhs176).ArraySet1(_rhs225, (_lhs177).Int())
			r0 = _rhs226
			(_lhs178).ArraySet1(_rhs227, (_lhs179).Int())
			_1007_v13 = _rhs228
			_1012_v18 = _rhs229
		}
		var _1015_v21 D4
		_ = _1015_v21
		_1015_v21 = Companion_D4_.Create_DC12_(_this.F23, _this.F23, (_this).F22())
		var _source11 D4 = _1015_v21
		_ = _source11
		if _source11.Is_DC12() {
			var _1016___mcc_h10 bool = _source11.Get_().(D4_DC12).Cf21
			_ = _1016___mcc_h10
			var _1017___mcc_h11 bool = _source11.Get_().(D4_DC12).Cf22
			_ = _1017___mcc_h11
			var _1018___mcc_h12 _dafny.Int = _source11.Get_().(D4_DC12).Cf23
			_ = _1018___mcc_h12
			var _1019_cf23 _dafny.Int = _1018___mcc_h12
			_ = _1019_cf23
			var _1020_cf22 bool = _1017___mcc_h11
			_ = _1020_cf22
			var _1021_cf21 bool = _1016___mcc_h10
			_ = _1021_cf21
			var _1022_v22 _dafny.Sequence
			_ = _1022_v22
			_1022_v22 = _dafny.UnicodeSeqOfUtf8Bytes("eh")
			r1 = _1022_v22
			var _1023_v23 _dafny.Array
			_ = _1023_v23
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw191
			_1023_v23 = _nw191
			var _1024_v24 _dafny.Sequence
			_ = _1024_v24
			_1024_v24 = _dafny.SeqOf(!(_1021_cf21))
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_1023_v23), 0))
			_ = _index203
			(_1023_v23).ArraySet1(_1024_v24, (_index203).Int())
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_1023_v23), 0))
			_ = _index204
			(_1023_v23).ArraySet1(_1024_v24, (_index204).Int())
			r2 = !_dafny.Companion_Sequence_.Equal(_1022_v22, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("x"), _1022_v22))
			var _1025_v25 _dafny.Map
			_ = _1025_v25
			_1025_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F22())
			_1025_v25 = (_1025_v25).Update((_1021_cf21) && (true), (_this).F22())
		} else if _source11.Is_DC13() {
			var _1026___mcc_h13 _dafny.Int = _source11.Get_().(D4_DC13).Cf24
			_ = _1026___mcc_h13
			var _1027___mcc_h14 _dafny.Int = _source11.Get_().(D4_DC13).Cf25
			_ = _1027___mcc_h14
			var _1028___mcc_h15 _dafny.Int = _source11.Get_().(D4_DC13).Cf26
			_ = _1028___mcc_h15
			var _1029_cf26 _dafny.Int = _1028___mcc_h15
			_ = _1029_cf26
			var _1030_cf25 _dafny.Int = _1027___mcc_h14
			_ = _1030_cf25
			var _1031_cf24 _dafny.Int = _1026___mcc_h13
			_ = _1031_cf24
			var _1032_v26 _dafny.Map
			_ = _1032_v26
			_1032_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1030_cf25, false)
			var _1033_v27 _dafny.MultiSet
			_ = _1033_v27
			_1033_v27 = _dafny.MultiSetOf(_1031_cf24, (_1032_v26).Cardinality())
			var _1034_v28 _dafny.Sequence
			_ = _1034_v28
			_1034_v28 = _dafny.SeqOf((Companion_Default___.Fm23(_1033_v27, _this.F23, globalState)).Cardinality(), (_1032_v26).Cardinality())
			var _1035_v29 _dafny.Array
			_ = _1035_v29
			var _nw192 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw192
			_1035_v29 = _nw192
			var _1036_v30 *C0
			_ = _1036_v30
			var _nw193 *C0 = New_C0_()
			_ = _nw193
			_nw193.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1034_v28, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_1037_cf25 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1038_i2 _dafny.Int) _dafny.Int {
					return _1037_cf25
				}
			})(_1030_cf25))))).Cardinality()), _1035_v29)
			_1036_v30 = _nw193
			(globalState).F2 = _this.F23
			(globalState).F2 = _this.F23
			if _this.F23 {
				var _1039_v31 _dafny.Map
				_ = _1039_v31
				_1039_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, _1031_cf24)
				var _1040_v32 D3
				_ = _1040_v32
				_1040_v32 = Companion_D3_.Create_DC9_(_dafny.SeqOf((_1039_v31).Cardinality(), (_1036_v30).F26(), (_1036_v30).F26()))
				var _pat_let_tv24 = _1034_v28
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1034_v28
				_ = _pat_let_tv25
				var _pat_let_tv26 = _1034_v28
				_ = _pat_let_tv26
				var _1041_v33 _dafny.Array
				_ = _1041_v33
				var _nwElement0_33 D3 = _1040_v32
				_ = _nwElement0_33
				var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(15))
				_ = _nw194
				(_nw194).ArraySet1(_nwElement0_33, 0)
				(_nw194).ArraySet1(_1040_v32, 1)
				(_nw194).ArraySet1(_1040_v32, 2)
				(_nw194).ArraySet1(func(_pat_let29_0 D3) D3 {
					return func(_1042_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let30_0 _dafny.Sequence) D3 {
							return func(_1043_dt__update_hcf18_h0 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC9_(_1043_dt__update_hcf18_h0)
							}(_pat_let30_0)
						}(_pat_let_tv24)
					}(_pat_let29_0)
				}(_1040_v32), 3)
				(_nw194).ArraySet1(_1040_v32, 4)
				(_nw194).ArraySet1(func(_pat_let31_0 D3) D3 {
					return func(_1044_dt__update__tmp_h2 D3) D3 {
						return func(_pat_let32_0 _dafny.Sequence) D3 {
							return func(_1045_dt__update_hcf18_h1 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC9_(_1045_dt__update_hcf18_h1)
							}(_pat_let32_0)
						}(_pat_let_tv25)
					}(_pat_let31_0)
				}(_1040_v32), 5)
				(_nw194).ArraySet1(func(_pat_let33_0 D3) D3 {
					return func(_1046_dt__update__tmp_h3 D3) D3 {
						return func(_pat_let34_0 _dafny.Sequence) D3 {
							return func(_1047_dt__update_hcf18_h2 _dafny.Sequence) D3 {
								return Companion_D3_.Create_DC9_(_1047_dt__update_hcf18_h2)
							}(_pat_let34_0)
						}(_pat_let_tv26)
					}(_pat_let33_0)
				}(_1040_v32), 6)
				(_nw194).ArraySet1((func() D3 {
					if _this.F23 {
						return _1040_v32
					}
					return _1040_v32
				})(), 7)
				(_nw194).ArraySet1((func() D3 {
					if _this.F23 {
						return _1040_v32
					}
					return _1040_v32
				})(), 8)
				(_nw194).ArraySet1(_1040_v32, 9)
				(_nw194).ArraySet1(_1040_v32, 10)
				(_nw194).ArraySet1(_1040_v32, 11)
				(_nw194).ArraySet1(_1040_v32, 12)
				(_nw194).ArraySet1(_1040_v32, 13)
				(_nw194).ArraySet1(_1040_v32, 14)
				_1041_v33 = _nw194
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1041_v33), 0))
				_ = _index205
				(_1041_v33).ArraySet1(_1040_v32, (_index205).Int())
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_1041_v33), 0))
				_ = _index206
				(_1041_v33).ArraySet1(_1040_v32, (_index206).Int())
				(_1036_v30).F27 = (func() _dafny.Array {
					if _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm17(globalState), _1034_v28) {
						return _1035_v29
					}
					return _1035_v29
				})()
				var _1048_v34 _dafny.Sequence
				_ = _1048_v34
				_1048_v34 = _dafny.UnicodeSeqOfUtf8Bytes("bglc")
				var _1049_v35 *C3
				_ = _1049_v35
				var _nw195 *C3 = New_C3_()
				_ = _nw195
				_nw195.Ctor__(_1048_v34, _this.F23, _dafny.Companion_Sequence_.Concatenate(_1034_v28, _1034_v28))
				_1049_v35 = _nw195
				_1049_v35 = _1049_v35
				(globalState).F2 = _this.F23
				var _1050_v36 *C0
				_ = _1050_v36
				var _nw196 *C0 = New_C0_()
				_ = _nw196
				_nw196.Ctor__(_1031_cf24, _1035_v29)
				_1050_v36 = _nw196
			} else {
				r3 = (_1030_cf25).Minus((_1036_v30).F26())
				r1 = Companion_Default___.Fm1(globalState)
				var _1051_v37 *C2
				_ = _1051_v37
				var _nw197 *C2 = New_C2_()
				_ = _nw197
				_nw197.Ctor__((_1036_v30).F26())
				_1051_v37 = _nw197
				_1031_cf24 = (_1036_v30).F26()
				var _1052_v38 _dafny.Array
				_ = _1052_v38
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw198
				_1052_v38 = _nw198
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1052_v38), 0))
				_ = _index207
				(_1052_v38).ArraySet1(_1034_v28, (_index207).Int())
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen((_1052_v38), 0))
				_ = _index208
				(_1052_v38).ArraySet1(_1034_v28, (_index208).Int())
			}
		} else {
			var _1053___mcc_h16 _dafny.MultiSet = _source11.Get_().(D4_DC11).Cf20
			_ = _1053___mcc_h16
			var _1054_cf20 _dafny.MultiSet = _1053___mcc_h16
			_ = _1054_cf20
			(globalState).F2 = _this.F23
			var _1055_v39 _dafny.Set
			_ = _1055_v39
			_1055_v39 = _dafny.SetOf(_this.F23)
			if ((_dafny.SetOf(_this.F23, _this.F23)).Union(_1055_v39)).IsSubsetOf(_1055_v39) {
				var _1056_v40 _dafny.Array
				_ = _1056_v40
				var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw199
				_1056_v40 = _nw199
				var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw200
				_1056_v40 = _nw200
				var _1057_v41 _dafny.Map
				_ = _1057_v41
				_1057_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F22())
				var _1058_v42 _dafny.Sequence
				_ = _1058_v42
				_1058_v42 = _dafny.UnicodeSeqOfUtf8Bytes("ap")
				_1057_v41 = (_1057_v41).Update((_this).Fm9((_this).F22(), globalState), _dafny.IntOfUint32((_1058_v42).Cardinality()))
				r0 = (_this).F22()
				var _rhs230 bool = true
				_ = _rhs230
				var _rhs231 _dafny.Array = _1056_v40
				_ = _rhs231
				var _rhs232 _dafny.Int = (_this).F22()
				_ = _rhs232
				var _rhs233 bool = _this.F23
				_ = _rhs233
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 *GlobalState = globalState
				_ = _lhs181
				_lhs180.F5 = _rhs230
				_1056_v40 = _rhs231
				r0 = _rhs232
				_lhs181.F2 = _rhs233
				r2 = _this.F23
			} else {
				r2 = _this.F23
				var _1059_v43 D0
				_ = _1059_v43
				_1059_v43 = Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("o"))
				var _1060_v44 _dafny.Map
				_ = _1060_v44
				_1060_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1059_v43, _this.F23)
				var _1061_v46 _dafny.Set
				_ = _1061_v46
				_1061_v46 = _dafny.SetOf(_1059_v43)
				var _1062_v47 _dafny.Set
				_ = _1062_v47
				_1062_v47 = _dafny.SetOf(func() _dafny.Map {
					var _coll54 = _dafny.NewMapBuilder()
					_ = _coll54
					for _iter59 := _dafny.Iterate((_1061_v46).Elements()); ; {
						_compr_54, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1063_v45 D0
						_1063_v45 = interface{}(_compr_54).(D0)
						if (_1061_v46).Contains(_1063_v45) {
							_coll54.Add(_1063_v45, _this.F23)
						}
					}
					return _coll54.ToMap()
				}())
				(_this).F23 = !(_1062_v47).Contains(_1060_v44)
				var _1064_v48 D1
				_ = _1064_v48
				_1064_v48 = Companion_D1_.Create_DC5_((_dafny.IntOfInt64(35)).Minus(_dafny.IntOfInt64(174)), _this.F23)
				var _1065_v49 _dafny.MultiSet
				_ = _1065_v49
				_1065_v49 = _dafny.MultiSetOf((func() bool {
					if true {
						return _this.F23
					}
					return _this.F23
				})())
				var _1066_v50 _dafny.Map
				_ = _1066_v50
				_1066_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(222), _this.F23)
				var _1067_v51 _dafny.Map
				_ = _1067_v51
				_1067_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, (_this).F22())
				var _rhs234 bool = Companion_Default___.Fm0(false, !(_1066_v50).Equals((_1066_v50).Update((_this).F22(), _this.F23)), globalState)
				_ = _rhs234
				var _rhs235 D1 = _1064_v48
				_ = _rhs235
				var _rhs236 _dafny.MultiSet = (_1065_v49).Update(_this.F23, Companion_Default___.Abs((func() _dafny.Int {
					if (_1067_v51).Contains(_this.F23) {
						return (_1067_v51).Get(_this.F23).(_dafny.Int)
					}
					return (_this).F22()
				})()))
				_ = _rhs236
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				_lhs182.F2 = _rhs234
				_1064_v48 = _rhs235
				_1065_v49 = _rhs236
				var _1068_v52 _dafny.Array
				_ = _1068_v52
				var _nwElement0_34 bool = _this.F23
				_ = _nwElement0_34
				var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(3))
				_ = _nw201
				(_nw201).ArraySet1(_nwElement0_34, 0)
				(_nw201).ArraySet1(_this.F23, 1)
				(_nw201).ArraySet1((_this.F23) && (_this.F23), 2)
				_1068_v52 = _nw201
				_1068_v52 = (func() _dafny.Array {
					if _this.F23 {
						return _1068_v52
					}
					return _1068_v52
				})()
				var _1069_v53 _dafny.Array
				_ = _1069_v53
				var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw202
				_1069_v53 = _nw202
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_1069_v53), 0))
				_ = _index209
				(_1069_v53).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F23)).Cardinality(), (_index209).Int())
				var _1070_v54 *C2
				_ = _1070_v54
				var _nw203 *C2 = New_C2_()
				_ = _nw203
				_nw203.Ctor__((_970_v0).Cardinality())
				_1070_v54 = _nw203
				var _1071_v55 _dafny.Array
				_ = _1071_v55
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
				_ = _nw204
				_1071_v55 = _nw204
				var _1072_v56 _dafny.MultiSet
				_ = _1072_v56
				_1072_v56 = _dafny.MultiSetOf(Companion_D1_.Create_DC5_((_1070_v54).F25(), _this.F23), _1064_v48)
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1071_v55), 0))
				_ = _index210
				(_1071_v55).ArraySet1(_1072_v56, (_index210).Int())
				var _1073_v57 _dafny.Sequence
				_ = _1073_v57
				_1073_v57 = _dafny.UnicodeSeqOfUtf8Bytes("fs")
				var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_1069_v53), 0))
				_ = _index211
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1071_v55), 0))
				_ = _index212
				var _rhs237 _dafny.Int = (_this).F22()
				_ = _rhs237
				var _rhs238 *C2 = _1070_v54
				_ = _rhs238
				var _rhs239 _dafny.MultiSet = (Companion_Default___.Fm30(_this.F23, _dafny.UnicodeSeqOfUtf8Bytes("ut"), true, _this.F23, globalState)).Difference(_1072_v56)
				_ = _rhs239
				var _rhs240 bool = _this.F23
				_ = _rhs240
				var _rhs241 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1073_v57).Cardinality()), ((_dafny.Zero).Minus((_1070_v54).F25())).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F23, !(_this.F23))).Cardinality())))
				_ = _rhs241
				var _lhs183 _dafny.Array = _1069_v53
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_1069_v53), 0))
				_ = _lhs184
				var _lhs185 _dafny.Array = _1071_v55
				_ = _lhs185
				var _lhs186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_1071_v55), 0))
				_ = _lhs186
				var _lhs187 *C4 = _this
				_ = _lhs187
				(_lhs183).ArraySet1(_rhs237, (_lhs184).Int())
				_1070_v54 = _rhs238
				(_lhs185).ArraySet1(_rhs239, (_lhs186).Int())
				_lhs187.F23 = _rhs240
				r0 = _rhs241
			}
			var _1074_v58 _dafny.Sequence
			_ = _1074_v58
			_1074_v58 = _dafny.UnicodeSeqOfUtf8Bytes("gao")
			var _1075_v59 _dafny.CodePoint
			_ = _1075_v59
			_1075_v59 = _dafny.CodePoint('x')
			Companion_Default___.M0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if _this.F23 {
					return _1074_v58
				}
				return _1074_v58
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F22()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _this.F23 {
					return _1074_v58
				}
				return _1074_v58
			})()).Cardinality()))).Uint32(), _1075_v59)).Cardinality()), _1074_v58, (_this).F22(), Companion_Default___.SafeModuloInt((_dafny.SetOf(_this.F23)).Cardinality(), _dafny.IntOfInt64(992)), globalState)
			var _1076_v60 _dafny.Sequence
			_ = _1076_v60
			_1076_v60 = _dafny.SeqOf((_this).F22())
			var _rhs242 _dafny.Int = (_this).F22()
			_ = _rhs242
			var _rhs243 bool = _this.F23
			_ = _rhs243
			var _rhs244 _dafny.CodePoint = _dafny.CodePoint('w')
			_ = _rhs244
			var _rhs245 _dafny.Int = (func() _dafny.Int {
				if _this.F23 {
					return (_this).F22()
				}
				return (_this).F22()
			})()
			_ = _rhs245
			var _rhs246 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1076_v60, _dafny.Companion_Sequence_.Concatenate(_1076_v60, _1076_v60))
			_ = _rhs246
			var _lhs188 *GlobalState = globalState
			_ = _lhs188
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			r0 = _rhs242
			_lhs188.F2 = _rhs243
			_lhs189.F15 = _rhs244
			r0 = _rhs245
			_1076_v60 = _rhs246
		}
		var _1077_v61 _dafny.Array
		_ = _1077_v61
		var _nw205 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(29))
		_ = _nw205
		_1077_v61 = _nw205
		var _1078_v62 _dafny.Map
		_ = _1078_v62
		_1078_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F23, (_this).F22())
		var _1079_v63 _dafny.Array
		_ = _1079_v63
		var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
		_ = _nw206
		_1079_v63 = _nw206
		var _1080_v64 D2
		_ = _1080_v64
		_1080_v64 = Companion_D2_.Create_DC8_(_1078_v62, _1079_v63, _this.F23, _this.F23, (_this).F22())
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1077_v61), 0))
		_ = _index213
		(_1077_v61).ArraySet1(_1080_v64, (_index213).Int())
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1077_v61), 0))
		_ = _index214
		(_1077_v61).ArraySet1(_1080_v64, (_index214).Int())
		r0 = _dafny.IntOfInt64(-398)
		var _1081_v65 _dafny.Sequence
		_ = _1081_v65
		_1081_v65 = _dafny.UnicodeSeqOfUtf8Bytes("m")
		r1 = _1081_v65
		r2 = _this.F23
		var _1082_v66 _dafny.Sequence
		_ = _1082_v66
		_1082_v66 = _dafny.SeqOf(_this.F23, _this.F23, _this.F23, (_this.F23) == (_this.F23), (func() bool {
			if (_970_v0).Contains(false) {
				return (_970_v0).Get(false).(bool)
			}
			return _this.F23
		})())
		r3 = _dafny.IntOfUint32((_1082_v66).Cardinality())
		return r0, r1, r2, r3
	}
}
func (_this *C4) F22() _dafny.Int {
	{
		return _this._f22
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f16 bool
	_f17 _dafny.Sequence
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C5_ struct {
}

var Companion_C5_ = CompanionStruct_C5_{}

func (_this *C5) Equals(other *C5) bool {
	return _this == other
}

func (_this *C5) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C5)
	return ok && _this.Equals(other)
}

func (*C5) String() string {
	return "_module.C5"
}

func Type_C5_() _dafny.TypeDescriptor {
	return type_C5_{}
}

type type_C5_ struct {
}

func (_this type_C5_) Default() interface{} {
	return (*C5)(nil)
}

func (_this type_C5_) String() string {
	return "main.C5"
}
func (_this *C5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F16() bool {
	return _this._f16
}
func (_this *C5) F16_set_(value bool) {
	_this._f16 = value
}
func (_this *C5) F17() _dafny.Sequence {
	return _this._f17
}
func (_this *C5) Ctor__(f16 bool, f17 _dafny.Sequence) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C5) Fm7(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		if _this.F16() {
			return _dafny.SeqOf(_dafny.CodePoint('y'))
		} else {
			return (Companion_D0_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(117))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}(func(_1083_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})))).Dtor_cf1()
		}
	}
}
func (_this *C5) Fm8(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf((_this).F17(), (Companion_D3_.Create_DC9_((_this).F17())).Dtor_cf18())
	}
}
func (_this *C5) M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		(_this).F16_set_(!(p2) || (Companion_Default___.Fm0(p2, p1, globalState)))
		var _1084_v0 _dafny.Array
		_ = _1084_v0
		var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw207
		_1084_v0 = _nw207
		for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1084_v0), 0))); ; {
			_guard_loop_5, _ok60 := _iter60()
			if !_ok60 {
				break
			}
			var _1085_i0 _dafny.Int
			_1085_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1085_i0).Sign() != -1) && ((_1085_i0).Cmp(_dafny.ArrayLen((_1084_v0), 0)) < 0)) {
				(_1084_v0).ArraySet1((_1085_i0).Plus(_dafny.IntOfInt64(8)), (_1085_i0).Int())
			}
		}
		var _1086_v1 *C4
		_ = _1086_v1
		var _nw208 *C4 = New_C4_()
		_ = _nw208
		_nw208.Ctor__(p3, true)
		_1086_v1 = _nw208
		var _1087_v2 _dafny.Array
		_ = _1087_v2
		var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
		_ = _nw209
		_1087_v2 = _nw209
		for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1087_v2), 0))); ; {
			_guard_loop_6, _ok61 := _iter61()
			if !_ok61 {
				break
			}
			var _1088_i1 _dafny.Int
			_1088_i1 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1088_i1).Sign() != -1) && ((_1088_i1).Cmp(_dafny.ArrayLen((_1087_v2), 0)) < 0)) {
				(_1087_v2).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-728))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}(func(_1089_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(59)
				})), (_1088_i1).Int())
			}
		}
		var _1090_v3 _dafny.Array
		_ = _1090_v3
		var _nw210 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(11))
		_ = _nw210
		_1090_v3 = _nw210
		for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1090_v3), 0))); ; {
			_guard_loop_7, _ok62 := _iter62()
			if !_ok62 {
				break
			}
			var _1091_i3 _dafny.Int
			_1091_i3 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1091_i3).Sign() != -1) && ((_1091_i3).Cmp(_dafny.ArrayLen((_1090_v3), 0)) < 0)) {
				(_1090_v3).ArraySet1(Companion_D0_.Create_DC0_(p0), (_1091_i3).Int())
			}
		}
		var _hi7 _dafny.Int = (_1086_v1).F22()
		_ = _hi7
		for _1092_i4 := p3; _1092_i4.Cmp(_hi7) < 0; _1092_i4 = _1092_i4.Plus(_dafny.One) {
			var _1093_v4 _dafny.Int
			_ = _1093_v4
			_1093_v4 = _dafny.IntOfInt64(577)
			_1093_v4 = p3
			var _1094_v5 _dafny.Int
			_ = _1094_v5
			var _1095_v6 _dafny.Sequence
			_ = _1095_v6
			var _1096_v7 bool
			_ = _1096_v7
			var _1097_v8 _dafny.Int
			_ = _1097_v8
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Sequence
			_ = _out7
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			_out6, _out7, _out8, _out9 = (_1086_v1).M7(globalState)
			_1094_v5 = _out6
			_1095_v6 = _out7
			_1096_v7 = _out8
			_1097_v8 = _out9
			var _1098_v9 _dafny.Array
			_ = _1098_v9
			var _nw211 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw211
			_1098_v9 = _nw211
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1098_v9), 0))
			_ = _index215
			(_1098_v9).ArraySet1(p2, (_index215).Int())
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_1084_v0), 0))
			_ = _index216
			(_1084_v0).ArraySet1((_1086_v1).F22(), (_index216).Int())
			var _1099_v10 _dafny.Set
			_ = _1099_v10
			_1099_v10 = _dafny.SetOf(_1092_i4)
			var _1100_v11 _dafny.Array
			_ = _1100_v11
			var _nwElement0_35 _dafny.Set = _1099_v10
			_ = _nwElement0_35
			var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(6))
			_ = _nw212
			(_nw212).ArraySet1(_nwElement0_35, 0)
			(_nw212).ArraySet1(_1099_v10, 1)
			(_nw212).ArraySet1(Companion_Default___.Fm31(_this.F16(), globalState), 2)
			(_nw212).ArraySet1(Companion_Default___.Fm31(p2, globalState), 3)
			(_nw212).ArraySet1(Companion_Default___.Fm31(_1086_v1.F23, globalState), 4)
			(_nw212).ArraySet1((_1099_v10).Intersection(_1099_v10), 5)
			_1100_v11 = _nw212
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1100_v11), 0))
			_ = _index217
			(_1100_v11).ArraySet1(_1099_v10, (_index217).Int())
			var _1101_v12 _dafny.CodePoint
			_ = _1101_v12
			_1101_v12 = _dafny.CodePoint('h')
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1098_v9), 0))
			_ = _index218
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_1084_v0), 0))
			_ = _index219
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1100_v11), 0))
			_ = _index220
			var _rhs247 bool = _this.F16()
			_ = _rhs247
			var _rhs248 _dafny.CodePoint = _1101_v12
			_ = _rhs248
			var _rhs249 _dafny.Int = _dafny.IntOfInt64(716)
			_ = _rhs249
			var _rhs250 _dafny.Set = _1099_v10
			_ = _rhs250
			var _lhs190 _dafny.Array = _1098_v9
			_ = _lhs190
			var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1098_v9), 0))
			_ = _lhs191
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			var _lhs193 _dafny.Array = _1084_v0
			_ = _lhs193
			var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_1084_v0), 0))
			_ = _lhs194
			var _lhs195 _dafny.Array = _1100_v11
			_ = _lhs195
			var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_1100_v11), 0))
			_ = _lhs196
			(_lhs190).ArraySet1(_rhs247, (_lhs191).Int())
			_lhs192.F15 = _rhs248
			(_lhs193).ArraySet1(_rhs249, (_lhs194).Int())
			(_lhs195).ArraySet1(_rhs250, (_lhs196).Int())
			var _1102_v13 _dafny.Map
			_ = _1102_v13
			_1102_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1086_v1).F22(), _1094_v5)
			var _1103_v14 _dafny.Array
			_ = _1103_v14
			var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw213
			_1103_v14 = _nw213
			var _1104_v15 _dafny.Sequence
			_ = _1104_v15
			_1104_v15 = _dafny.SeqOf(_this.F16())
			var _1105_v16 _dafny.Map
			_ = _1105_v16
			_1105_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F16()), _1103_v14)
			var _1106_v17 _dafny.Map
			_ = _1106_v17
			_1106_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (func() _dafny.Array {
				if (_1105_v16).Contains(p1) {
					return (_1105_v16).Get(p1).(_dafny.Array)
				}
				return _1103_v14
			})())
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1098_v9), 0))
			_ = _index221
			var _rhs251 bool = ((_1104_v15).Select((Companion_Default___.SafeIndex(_1097_v8, _dafny.IntOfUint32((_1104_v15).Cardinality()))).Uint32()).(bool)) && ((_this).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("vjxld"), globalState))
			_ = _rhs251
			var _rhs252 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1097_v8, (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1094_v5, _1093_v4))))
			_ = _rhs252
			var _rhs253 _dafny.Array = (func() _dafny.Array {
				if (_1106_v17).Contains(_1092_i4) {
					return (_1106_v17).Get(_1092_i4).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if p2 {
						return _1103_v14
					}
					return _1103_v14
				})()
			})()
			_ = _rhs253
			var _lhs197 _dafny.Array = _1098_v9
			_ = _lhs197
			var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_1098_v9), 0))
			_ = _lhs198
			(_lhs197).ArraySet1(_rhs251, (_lhs198).Int())
			_1102_v13 = _rhs252
			_1103_v14 = _rhs253
		}
		var _1107_v18 _dafny.Sequence
		_ = _1107_v18
		_1107_v18 = _dafny.SeqOf(_this.F16())
		r0 = _1107_v18
		return r0
	}
}
func (_this *C5) M5(globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		if !(_this.F16()) || (_this.F16()) {
			if _this.F16() {
				var _1108_v0 _dafny.Sequence
				_ = _1108_v0
				_1108_v0 = _dafny.SeqOf(Companion_Default___.Fm19(_dafny.IntOfInt64(125), globalState))
				var _1109_v1 _dafny.Int
				_ = _1109_v1
				_1109_v1 = _dafny.IntOfInt64(42)
				var _1110_v2 _dafny.Sequence
				_ = _1110_v2
				_1110_v2 = _dafny.SeqOf(_this.F16(), !(_this.F16()), _this.F16(), _this.F16())
				var _1111_v3 _dafny.MultiSet
				_ = _1111_v3
				_1111_v3 = _dafny.MultiSetOf(_this.F16(), _this.F16())
				var _1112_v4 _dafny.Map
				_ = _1112_v4
				_1112_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1111_v3).Cardinality(), !(_this.F16()))
				var _1113_v5 _dafny.CodePoint
				_ = _1113_v5
				_1113_v5 = _dafny.CodePoint('y')
				var _1114_v6 _dafny.MultiSet
				_ = _1114_v6
				_1114_v6 = _dafny.MultiSetOf(_1113_v5)
				var _1115_v7 _dafny.Map
				_ = _1115_v7
				_1115_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1109_v1)
				var _1116_v8 _dafny.Array
				_ = _1116_v8
				var _nwElement0_36 _dafny.Int = _1109_v1
				_ = _nwElement0_36
				var _nw214 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(28))
				_ = _nw214
				(_nw214).ArraySet1(_nwElement0_36, 0)
				(_nw214).ArraySet1(_dafny.IntOfInt64(203), 1)
				(_nw214).ArraySet1(Companion_Default___.SafeModuloInt(_1109_v1, _1109_v1), 2)
				(_nw214).ArraySet1((_1109_v1).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1109_v1)).Cardinality()))), 3)
				(_nw214).ArraySet1(_1109_v1, 4)
				(_nw214).ArraySet1(_1109_v1, 5)
				(_nw214).ArraySet1(_1109_v1, 6)
				(_nw214).ArraySet1(_1109_v1, 7)
				(_nw214).ArraySet1((_1109_v1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F16(), _this.F16())).Cardinality())), 8)
				(_nw214).ArraySet1(_1109_v1, 9)
				(_nw214).ArraySet1(_dafny.IntOfUint32((_1110_v2).Cardinality()), 10)
				(_nw214).ArraySet1((_dafny.IntOfInt64(187)).Plus(_1109_v1), 11)
				(_nw214).ArraySet1(_1109_v1, 12)
				(_nw214).ArraySet1(_1109_v1, 13)
				(_nw214).ArraySet1(_1109_v1, 14)
				(_nw214).ArraySet1(((_1112_v4).Cardinality()).Times(_dafny.IntOfInt64(-267)), 15)
				(_nw214).ArraySet1(_1109_v1, 16)
				(_nw214).ArraySet1(_1109_v1, 17)
				(_nw214).ArraySet1((_1111_v3).Cardinality(), 18)
				(_nw214).ArraySet1((_1109_v1).Minus(_1109_v1), 19)
				(_nw214).ArraySet1((_1109_v1).Minus(((_this).F17()).Select((Companion_Default___.SafeIndex(_1109_v1, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int)), 20)
				(_nw214).ArraySet1(_dafny.IntOfInt64(-718), 21)
				(_nw214).ArraySet1(_1109_v1, 22)
				(_nw214).ArraySet1(_1109_v1, 23)
				(_nw214).ArraySet1(_1109_v1, 24)
				(_nw214).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_1114_v6).Contains(_1113_v5) {
						return (_1114_v6).Multiplicity(_1113_v5)
					}
					return _1109_v1
				})(), (_1115_v7).Cardinality()), 25)
				(_nw214).ArraySet1((_1109_v1).Times(_dafny.IntOfUint32(((_this).F17()).Cardinality())), 26)
				(_nw214).ArraySet1(_1109_v1, 27)
				_1116_v8 = _nw214
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _index222
				(_1116_v8).ArraySet1(_1109_v1, (_index222).Int())
				var _1117_v9 _dafny.Map
				_ = _1117_v9
				_1117_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _1110_v2)
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _index223
				var _rhs254 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _this.F16() {
						return _1108_v0
					}
					return _1108_v0
				})(), Companion_Default___.Fm32(_dafny.IntOfInt64(102), false, _this.F16(), _1113_v5, globalState))
				_ = _rhs254
				var _rhs255 _dafny.Int = (((_1117_v9).Cardinality()).Times(_1109_v1)).Minus(Companion_Default___.Fm3(globalState))
				_ = _rhs255
				var _lhs199 _dafny.Array = _1116_v8
				_ = _lhs199
				var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _lhs200
				_1108_v0 = _rhs254
				(_lhs199).ArraySet1(_rhs255, (_lhs200).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _index224
				(_1116_v8).ArraySet1((_1116_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))).Int()).(_dafny.Int), (_index224).Int())
				var _1118_v10 _dafny.Sequence
				_ = _1118_v10
				_1118_v10 = _dafny.UnicodeSeqOfUtf8Bytes("klmwlqbvh")
				Companion_Default___.M0(_1109_v1, _1118_v10, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nvw"), _1118_v10)).Cardinality()), Companion_Default___.SafeDivisionInt((_1116_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))).Int()).(_dafny.Int), (_1116_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))).Int()).(_dafny.Int)), globalState)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _index225
				(_1116_v8).ArraySet1(((_1116_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))).Int()).(_dafny.Int)).Times((_1109_v1).Times((_1116_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))).Int()).(_dafny.Int))), (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1116_v8), 0))
				_ = _index226
				(_1116_v8).ArraySet1((_dafny.Zero).Minus(_1109_v1), (_index226).Int())
			} else {
				var _1119_v11 _dafny.Int
				_ = _1119_v11
				_1119_v11 = _dafny.Zero
				var _1120_v12 _dafny.Sequence
				_ = _1120_v12
				_1120_v12 = _dafny.UnicodeSeqOfUtf8Bytes("sf")
				_1119_v11 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1120_v12, Companion_Default___.Fm1(globalState)), _1120_v12)).Cardinality())
				_1119_v11 = _1119_v11
				_1119_v11 = _1119_v11
				var _1121_v13 D0
				_ = _1121_v13
				_1121_v13 = Companion_D0_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("xvkvrqxb"))
				var _1122_v14 _dafny.CodePoint
				_ = _1122_v14
				_1122_v14 = _dafny.CodePoint('s')
				var _1123_v15 _dafny.MultiSet
				_ = _1123_v15
				_1123_v15 = _dafny.MultiSetOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if _this.F16() {
						return _1120_v12
					}
					return _dafny.Companion_Sequence_.Update((_1121_v13).Dtor_cf1(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.IntOfUint32(((_1121_v13).Dtor_cf1()).Cardinality()))).Uint32(), _1122_v14)
				})()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(185))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_1124_v11 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1125_i0 _dafny.Int) _dafny.Int {
						return _1124_v11
					}
				})(_1119_v11)))).Cardinality()))
				_1123_v15 = _dafny.MultiSetOf((_1119_v11).Times(_1119_v11), _1119_v11, _1119_v11)
				var _1126_v16 _dafny.Array
				_ = _1126_v16
				var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw215
				_1126_v16 = _nw215
				_1126_v16 = (func() _dafny.Array {
					if _this.F16() {
						return _1126_v16
					}
					return _1126_v16
				})()
			}
			var _source12 D1 = Companion_Default___.Fm33((_this).F17(), globalState)
			_ = _source12
			if _source12.Is_DC4() {
				var _1127___mcc_h0 bool = _source12.Get_().(D1_DC4).Cf7
				_ = _1127___mcc_h0
				var _1128___mcc_h1 _dafny.Int = _source12.Get_().(D1_DC4).Cf8
				_ = _1128___mcc_h1
				var _1129___mcc_h2 _dafny.Map = _source12.Get_().(D1_DC4).Cf9
				_ = _1129___mcc_h2
				var _1130_cf9 _dafny.Map = _1129___mcc_h2
				_ = _1130_cf9
				var _1131_cf8 _dafny.Int = _1128___mcc_h1
				_ = _1131_cf8
				var _1132_cf7 bool = _1127___mcc_h0
				_ = _1132_cf7
				var _1133_v17 _dafny.Map
				_ = _1133_v17
				_1133_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1132_cf7, _dafny.IntOfInt64(-752))
				_1133_v17 = _1133_v17
				var _1134_v18 _dafny.Sequence
				_ = _1134_v18
				_1134_v18 = _dafny.UnicodeSeqOfUtf8Bytes("m")
				_1131_cf8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1134_v18, _1134_v18), _dafny.UnicodeSeqOfUtf8Bytes("xlh"))).Cardinality())
				var _1135_v19 _dafny.MultiSet
				_ = _1135_v19
				_1135_v19 = _dafny.MultiSetOf(_1131_cf8)
				var _1136_v20 *C1
				_ = _1136_v20
				var _nw216 *C1 = New_C1_()
				_ = _nw216
				_nw216.Ctor__(_1131_cf8, _1131_cf8, (_dafny.MultiSetOf((_dafny.Zero).Minus(_1131_cf8))).IsDisjointFrom(_1135_v19), _dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17()), Companion_Default___.SafeDivisionInt(_1131_cf8, _1131_cf8))
				_1136_v20 = _nw216
				var _1137_v21 _dafny.Array
				_ = _1137_v21
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_30
				var _nw217 _dafny.Array
				_ = _nw217
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw217 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1138_v18 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1139_i1 _dafny.Int) _dafny.Sequence {
							return _1138_v18
						}
					})(_1134_v18)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw217 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw217).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw217).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1137_v21 = _nw217
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1137_v21), 0))
				_ = _index227
				(_1137_v21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1134_v18, _1134_v18), (_index227).Int())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1137_v21), 0))
				_ = _index228
				var _rhs256 *C1 = _1136_v20
				_ = _rhs256
				var _rhs257 _dafny.Sequence = _1134_v18
				_ = _rhs257
				var _lhs201 _dafny.Array = _1137_v21
				_ = _lhs201
				var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1137_v21), 0))
				_ = _lhs202
				_1136_v20 = _rhs256
				(_lhs201).ArraySet1(_rhs257, (_lhs202).Int())
				var _1140_v22 D1
				_ = _1140_v22
				_1140_v22 = Companion_D1_.Create_DC5_((func() _dafny.Int {
					if (_1133_v17).Contains(_1132_cf7) {
						return (_1133_v17).Get(_1132_cf7).(_dafny.Int)
					}
					return _dafny.IntOfUint32(((_1137_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1137_v21), 0))).Int()).(_dafny.Sequence)).Cardinality())
				})(), _1132_cf7)
				Companion_Default___.M0(_1131_cf8, _1134_v18, Companion_Default___.SafeDivisionInt((_1140_v22).Dtor_cf10(), (_1136_v20).F30()), _1131_cf8, globalState)
			} else if _source12.Is_DC5() {
				var _1141___mcc_h3 _dafny.Int = _source12.Get_().(D1_DC5).Cf10
				_ = _1141___mcc_h3
				var _1142___mcc_h4 bool = _source12.Get_().(D1_DC5).Cf11
				_ = _1142___mcc_h4
				var _1143_cf11 bool = _1142___mcc_h4
				_ = _1143_cf11
				var _1144_cf10 _dafny.Int = _1141___mcc_h3
				_ = _1144_cf10
				var _1145_v23 _dafny.Sequence
				_ = _1145_v23
				_1145_v23 = _dafny.SeqOf(_1143_cf11)
				var _1146_v24 _dafny.Sequence
				_ = _1146_v24
				_1146_v24 = _dafny.UnicodeSeqOfUtf8Bytes("rnjvqpd")
				var _1147_v25 *C3
				_ = _1147_v25
				var _nw218 *C3 = New_C3_()
				_ = _nw218
				_nw218.Ctor__(_1146_v24, _this.F16(), (_this).F17())
				_1147_v25 = _nw218
				var _1148_v26 _dafny.MultiSet
				_ = _1148_v26
				_1148_v26 = _dafny.MultiSetOf(_1144_cf10, (_dafny.SetOf(_1147_v25)).Cardinality())
				var _1149_v27 _dafny.Set
				_ = _1149_v27
				_1149_v27 = _dafny.SetOf(Companion_Default___.Fm3(globalState))
				var _1150_v28 _dafny.Map
				_ = _1150_v28
				_1150_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1143_cf11, _1144_cf10)
				var _1151_v29 _dafny.Array
				_ = _1151_v29
				var _nwElement0_37 _dafny.Int = _1144_cf10
				_ = _nwElement0_37
				var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(21))
				_ = _nw219
				(_nw219).ArraySet1(_nwElement0_37, 0)
				(_nw219).ArraySet1((_dafny.Zero).Minus(_1144_cf10), 1)
				(_nw219).ArraySet1(_dafny.IntOfUint32((_1145_v23).Cardinality()), 2)
				(_nw219).ArraySet1(_dafny.IntOfUint32((_1146_v24).Cardinality()), 3)
				(_nw219).ArraySet1(_dafny.IntOfInt64(575), 4)
				(_nw219).ArraySet1((_dafny.MultiSetOf(_1144_cf10, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Cardinality(), 5)
				(_nw219).ArraySet1(_1144_cf10, 6)
				(_nw219).ArraySet1(_1144_cf10, 7)
				(_nw219).ArraySet1(_1144_cf10, 8)
				(_nw219).ArraySet1(_1144_cf10, 9)
				(_nw219).ArraySet1(_1144_cf10, 10)
				(_nw219).ArraySet1(_1144_cf10, 11)
				(_nw219).ArraySet1(_dafny.IntOfUint32((_1145_v23).Cardinality()), 12)
				(_nw219).ArraySet1(_1144_cf10, 13)
				(_nw219).ArraySet1(_1144_cf10, 14)
				(_nw219).ArraySet1(_1144_cf10, 15)
				(_nw219).ArraySet1(_1144_cf10, 16)
				(_nw219).ArraySet1((func() _dafny.Int {
					if (_1148_v26).Contains(_dafny.IntOfInt64(396)) {
						return (_1148_v26).Multiplicity(_dafny.IntOfInt64(396))
					}
					return (_1149_v27).Cardinality()
				})(), 17)
				(_nw219).ArraySet1(_dafny.IntOfUint32((_1146_v24).Cardinality()), 18)
				(_nw219).ArraySet1(_1144_cf10, 19)
				(_nw219).ArraySet1((func() _dafny.Int {
					if (_1150_v28).Contains(_this.F16()) {
						return (_1150_v28).Get(_this.F16()).(_dafny.Int)
					}
					return _1144_cf10
				})(), 20)
				_1151_v29 = _nw219
				var _1152_v30 _dafny.MultiSet
				_ = _1152_v30
				_1152_v30 = _dafny.MultiSetOf((func() _dafny.Array {
					if _1143_cf11 {
						return _1151_v29
					}
					return _1151_v29
				})())
				var _1153_v31 _dafny.Sequence
				_ = _1153_v31
				_1153_v31 = _dafny.SeqOf(_1152_v30)
				var _1154_v32 _dafny.Sequence
				_ = _1154_v32
				_1154_v32 = _dafny.SeqOf(_1146_v24)
				_1152_v30 = (_1153_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1154_v32, Companion_Default___.Fm18(_1144_cf10, globalState))).Cardinality()), _dafny.IntOfUint32((_1153_v31).Cardinality()))).Uint32()).(_dafny.MultiSet)
				var _1155_v33 _dafny.Array
				_ = _1155_v33
				var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw220
				_1155_v33 = _nw220
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1155_v33), 0))
				_ = _index229
				(_1155_v33).ArraySet1(_1146_v24, (_index229).Int())
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1155_v33), 0))
				_ = _index230
				var _rhs258 _dafny.Sequence = _1146_v24
				_ = _rhs258
				var _rhs259 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1146_v24, (_1147_v25).F24())
				_ = _rhs259
				var _rhs260 bool = _1143_cf11
				_ = _rhs260
				var _lhs203 _dafny.Array = _1155_v33
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_1155_v33), 0))
				_ = _lhs204
				var _lhs205 *GlobalState = globalState
				_ = _lhs205
				_1146_v24 = _rhs258
				(_lhs203).ArraySet1(_rhs259, (_lhs204).Int())
				_lhs205.F5 = _rhs260
				_1144_cf10 = _1144_cf10
				(globalState).F2 = !(_this.F16())
			} else if _source12.Is_DC6() {
				var _1156_v34 _dafny.Map
				_ = _1156_v34
				_1156_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _this.F16())).Cardinality())
				var _1157_v35 _dafny.Sequence
				_ = _1157_v35
				_1157_v35 = _dafny.SeqOf(_this.F16())
				_1156_v34 = (_1156_v34).Update(!(_this.F16()), _dafny.IntOfUint32((_1157_v35).Cardinality()))
				var _1158_v36 _dafny.Int
				_ = _1158_v36
				_1158_v36 = _dafny.IntOfInt64(39)
				var _1159_v37 D8
				_ = _1159_v37
				_1159_v37 = Companion_D8_.Create_DC21_(_1158_v36, _this.F16(), _this.F16(), false, _1158_v36)
				var _1160_v38 D4
				_ = _1160_v38
				_1160_v38 = Companion_D4_.Create_DC12_(_this.F16(), _this.F16(), (_1159_v37).Dtor_cf39())
				var _1161_v39 _dafny.Array
				_ = _1161_v39
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_31
				var _nw221 _dafny.Array
				_ = _nw221
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw221 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) bool = func(_1162_i2 _dafny.Int) bool {
						return _this.F16()
					}
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw221 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw221).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw221).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1161_v39 = _nw221
				var _1163_v40 *C0
				_ = _1163_v40
				var _nw222 *C0 = New_C0_()
				_ = _nw222
				_nw222.Ctor__((_1158_v36).Times((_1160_v38).Dtor_cf23()), _1161_v39)
				_1163_v40 = _nw222
				var _1164_v41 _dafny.Map
				_ = _1164_v41
				_1164_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1163_v40).F26()), _this.F16())
				var _1165_v42 _dafny.CodePoint
				_ = _1165_v42
				_1165_v42 = _dafny.CodePoint('x')
				(_this).F16_set_(!(((_this.F16()) || ((func() bool {
					if (_1164_v41).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _1165_v42)).Cardinality()) {
						return (_1164_v41).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _1165_v42)).Cardinality()).(bool)
					}
					return _this.F16()
				})())) == (_this.F16())))
				var _rhs261 bool = !((_1158_v36).Cmp((_1163_v40).F26()) <= 0)
				_ = _rhs261
				var _rhs262 _dafny.Int = (_dafny.Zero).Minus(_1158_v36)
				_ = _rhs262
				var _lhs206 *GlobalState = globalState
				_ = _lhs206
				_lhs206.F2 = _rhs261
				_1158_v36 = _rhs262
			} else {
				var _1166___mcc_h5 _dafny.MultiSet = _source12.Get_().(D1_DC3).Cf6
				_ = _1166___mcc_h5
				var _1167_cf6 _dafny.MultiSet = _1166___mcc_h5
				_ = _1167_cf6
				var _1168_v43 _dafny.Array
				_ = _1168_v43
				var _nw223 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw223
				_1168_v43 = _nw223
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_32
				var _nw224 _dafny.Array
				_ = _nw224
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw224 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = func(_1169_i3 _dafny.Int) bool {
						return _this.F16()
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw224 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw224).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw224).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1168_v43 = _nw224
				var _1170_v44 _dafny.Int
				_ = _1170_v44
				_1170_v44 = _dafny.IntOfInt64(-578)
				var _1171_v45 _dafny.Sequence
				_ = _1171_v45
				_1171_v45 = _dafny.UnicodeSeqOfUtf8Bytes("lbnnl")
				Companion_Default___.M0(_1170_v44, _1171_v45, _1170_v44, (func() _dafny.Int {
					if false {
						return _1170_v44
					}
					return _1170_v44
				})(), globalState)
				var _1172_v46 *C4
				_ = _1172_v46
				var _nw225 *C4 = New_C4_()
				_ = _nw225
				_nw225.Ctor__(_1170_v44, _this.F16())
				_1172_v46 = _nw225
				var _1173_v47 _dafny.MultiSet
				_ = _1173_v47
				_1173_v47 = _dafny.MultiSetOf(_1172_v46, _1172_v46, _1172_v46)
				var _1174_v49 _dafny.Sequence
				_ = _1174_v49
				_1174_v49 = _dafny.SeqOf(Companion_Default___.Fm17(globalState))
				var _1175_v50 _dafny.Map
				_ = _1175_v50
				_1175_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1167_cf6).Cardinality(), _1174_v49)
				var _1176_v51 _dafny.Map
				_ = _1176_v51
				_1176_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1173_v47, (_dafny.Zero).Minus((func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter63 := _dafny.Iterate(((func() _dafny.Sequence {
						if (_1175_v50).Contains((_1172_v46).F22()) {
							return (_1175_v50).Get((_1172_v46).F22()).(_dafny.Sequence)
						}
						return _1174_v49
					})()).Elements()); ; {
						_compr_55, _ok63 := _iter63()
						if !_ok63 {
							break
						}
						var _1177_v48 _dafny.Sequence
						_1177_v48 = interface{}(_compr_55).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
							if (_1175_v50).Contains((_1172_v46).F22()) {
								return (_1175_v50).Get((_1172_v46).F22()).(_dafny.Sequence)
							}
							return _1174_v49
						})(), _1177_v48) {
							_coll55.Add(_1177_v48, _1172_v46.F23)
						}
					}
					return _coll55.ToMap()
				}()).Cardinality()))
				var _1178_v52 _dafny.Sequence
				_ = _1178_v52
				_1178_v52 = _dafny.SeqOf(_1172_v46)
				_1176_v51 = (_1176_v51).Update(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1178_v52, _1178_v52)), Companion_Default___.Fm3(globalState))
				var _1179_v53 D1
				_ = _1179_v53
				_1179_v53 = Companion_D1_.Create_DC5_((_1172_v46).F22(), _1172_v46.F23)
				(_this).F16_set_(((_1179_v53).Dtor_cf10()).Cmp(((_1172_v46).F22()).Minus((_1172_v46).F22())) <= 0)
			}
			var _1180_v54 _dafny.Int
			_ = _1180_v54
			_1180_v54 = _dafny.IntOfInt64(649)
			_1180_v54 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.Fm3(globalState)), (_dafny.Zero).Minus(_1180_v54))
			var _1181_v55 *C2
			_ = _1181_v55
			var _nw226 *C2 = New_C2_()
			_ = _nw226
			_nw226.Ctor__(_1180_v54)
			_1181_v55 = _nw226
			(globalState).F2 = !(!((func() bool {
				if _this.F16() {
					return _this.F16()
				}
				return _this.F16()
			})()))
		} else {
			var _1182_v56 _dafny.Int
			_ = _1182_v56
			_1182_v56 = _dafny.IntOfInt64(119)
			_1182_v56 = Companion_Default___.SafeDivisionInt(_1182_v56, _1182_v56)
			var _1183_v57 _dafny.Set
			_ = _1183_v57
			_1183_v57 = _dafny.SetOf(Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState))
			var _1184_v58 _dafny.Sequence
			_ = _1184_v58
			_1184_v58 = _dafny.SeqOf(true, (_1183_v57).IsDisjointFrom(_1183_v57), _this.F16(), (_1182_v56).Cmp((_dafny.Zero).Minus(_1182_v56)) > 0)
			(_this).F16_set_((_1184_v58).Select((Companion_Default___.SafeIndex(_1182_v56, _dafny.IntOfUint32((_1184_v58).Cardinality()))).Uint32()).(bool))
			var _1185_v59 _dafny.MultiSet
			_ = _1185_v59
			_1185_v59 = _dafny.MultiSetOf(_1182_v56)
			var _1186_v60 _dafny.Map
			_ = _1186_v60
			_1186_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1185_v59, _this.F16())
			_1186_v60 = (_1186_v60).Update(((_dafny.MultiSetOf(_1182_v56)).Update(_1182_v56, Companion_Default___.Abs(_1182_v56))).Update(_1182_v56, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_1182_v56)), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), _1182_v56)).Cardinality()))), !(!(_this.F16())))
			var _1187_v61 _dafny.Sequence
			_ = _1187_v61
			_1187_v61 = _dafny.UnicodeSeqOfUtf8Bytes("wbufmbxty")
			var _source13 D1 = Companion_Default___.Fm34(_1182_v56, _1182_v56, _1183_v57, _dafny.IntOfUint32((_1187_v61).Cardinality()), globalState)
			_ = _source13
			if _source13.Is_DC4() {
				var _1188___mcc_h6 bool = _source13.Get_().(D1_DC4).Cf7
				_ = _1188___mcc_h6
				var _1189___mcc_h7 _dafny.Int = _source13.Get_().(D1_DC4).Cf8
				_ = _1189___mcc_h7
				var _1190___mcc_h8 _dafny.Map = _source13.Get_().(D1_DC4).Cf9
				_ = _1190___mcc_h8
				var _1191_cf9 _dafny.Map = _1190___mcc_h8
				_ = _1191_cf9
				var _1192_cf8 _dafny.Int = _1189___mcc_h7
				_ = _1192_cf8
				var _1193_cf7 bool = _1188___mcc_h6
				_ = _1193_cf7
				var _1194_v62 _dafny.CodePoint
				_ = _1194_v62
				_1194_v62 = _dafny.CodePoint('s')
				var _1195_v63 _dafny.Map
				_ = _1195_v63
				_1195_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1182_v56), _1182_v56), _1194_v62)
				_1195_v63 = (_1195_v63).Update(_1182_v56, _1194_v62)
				var _1196_v64 *C2
				_ = _1196_v64
				var _nw227 *C2 = New_C2_()
				_ = _nw227
				_nw227.Ctor__(Companion_Default___.Fm3(globalState))
				_1196_v64 = _nw227
				_1187_v61 = _dafny.Companion_Sequence_.Concatenate(_1187_v61, _1187_v61)
				var _1197_v65 _dafny.Array
				_ = _1197_v65
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw228
				_1197_v65 = _nw228
				var _1198_v66 _dafny.Array
				_ = _1198_v66
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_33
				var _nw229 _dafny.Array
				_ = _nw229
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw229 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1199_v64 *C2) func(_dafny.Int) _dafny.Int {
						return func(_1200_i4 _dafny.Int) _dafny.Int {
							return (_1200_i4).Times((_1199_v64).F25())
						}
					})(_1196_v64)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw229 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw229).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw229).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1198_v66 = _nw229
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1197_v65), 0))
				_ = _index231
				(_1197_v65).ArraySet1(_1198_v66, (_index231).Int())
				var _1201_v67 _dafny.MultiSet
				_ = _1201_v67
				_1201_v67 = _dafny.MultiSetOf(_1193_cf7, _1193_cf7)
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1197_v65), 0))
				_ = _index232
				var _rhs263 _dafny.Int = (_1182_v56).Plus((_dafny.Zero).Minus((func() _dafny.Int {
					if _this.F16() {
						return (_dafny.Zero).Minus(_1182_v56)
					}
					return (_1196_v64).F25()
				})()))
				_ = _rhs263
				var _rhs264 _dafny.Array = _1198_v66
				_ = _rhs264
				var _rhs265 bool = (_this).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-637))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_1202_v62 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1203_i5 _dafny.Int) _dafny.CodePoint {
						return _1202_v62
					}
				})(_1194_v62))), globalState)
				_ = _rhs265
				var _rhs266 _dafny.Sequence = _dafny.SeqOf((_this.F16()) && (_this.F16()), _1193_cf7, ((Companion_D8_.Create_DC21_(_dafny.IntOfInt64(62), !(_this.F16()), _this.F16(), _1193_cf7, (_1201_v67).Cardinality())).Dtor_cf35()).Cmp(_1192_cf8) == 0, true)
				_ = _rhs266
				var _rhs267 bool = (_1201_v67).IsSubsetOf((_1201_v67).Union(_dafny.MultiSetFromSeq(_1184_v58)))
				_ = _rhs267
				var _lhs207 _dafny.Array = _1197_v65
				_ = _lhs207
				var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1197_v65), 0))
				_ = _lhs208
				var _lhs209 *C5 = _this
				_ = _lhs209
				var _lhs210 *GlobalState = globalState
				_ = _lhs210
				_1192_cf8 = _rhs263
				(_lhs207).ArraySet1(_rhs264, (_lhs208).Int())
				_lhs209.F16_set_(_rhs265)
				_1184_v58 = _rhs266
				_lhs210.F2 = _rhs267
			} else if _source13.Is_DC5() {
				var _1204___mcc_h9 _dafny.Int = _source13.Get_().(D1_DC5).Cf10
				_ = _1204___mcc_h9
				var _1205___mcc_h10 bool = _source13.Get_().(D1_DC5).Cf11
				_ = _1205___mcc_h10
				var _1206_cf11 bool = _1205___mcc_h10
				_ = _1206_cf11
				var _1207_cf10 _dafny.Int = _1204___mcc_h9
				_ = _1207_cf10
				var _1208_v68 _dafny.Map
				_ = _1208_v68
				_1208_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_cf10, (_this).F17())
				var _1209_v69 _dafny.CodePoint
				_ = _1209_v69
				_1209_v69 = _dafny.CodePoint('d')
				var _1210_v70 _dafny.Map
				_ = _1210_v70
				_1210_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1209_v69, _dafny.IntOfInt64(483))
				_1208_v68 = (_1208_v68).Update((func() _dafny.Set {
					var _coll56 = _dafny.NewBuilder()
					_ = _coll56
					for _iter64 := _dafny.Iterate((_1210_v70).Keys().Elements()); ; {
						_compr_56, _ok64 := _iter64()
						if !_ok64 {
							break
						}
						var _1211_v71 _dafny.CodePoint
						_1211_v71 = interface{}(_compr_56).(_dafny.CodePoint)
						if (_1210_v70).Contains(_1211_v71) {
							_coll56.Add(_1211_v71)
						}
					}
					return _coll56.ToSet()
				}()).Cardinality(), (_this).F17())
				_1206_cf11 = _this.F16()
				var _1212_v72 _dafny.Array
				_ = _1212_v72
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_34
				var _nw230 _dafny.Array
				_ = _nw230
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw230 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Int = (func(_1213_cf10 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1214_i6 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1214_i6, _1213_cf10)
						}
					})(_1207_cf10)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw230 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw230).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw230).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1212_v72 = _nw230
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1212_v72), 0))
				_ = _index233
				(_1212_v72).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hfgkuwyr")).Cardinality()), (_index233).Int())
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1212_v72), 0))
				_ = _index234
				var _rhs268 _dafny.Int = ((_1182_v56).Minus(_1182_v56)).Minus(_dafny.IntOfInt64(-640))
				_ = _rhs268
				var _rhs269 _dafny.Int = ((_this).F17()).Select((Companion_Default___.SafeIndex(_1182_v56, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs269
				var _rhs270 _dafny.Sequence = _1187_v61
				_ = _rhs270
				var _rhs271 _dafny.Int = (func() _dafny.Int {
					if _this.F16() {
						return _1182_v56
					}
					return _1182_v56
				})()
				_ = _rhs271
				var _lhs211 _dafny.Array = _1212_v72
				_ = _lhs211
				var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1212_v72), 0))
				_ = _lhs212
				(_lhs211).ArraySet1(_rhs268, (_lhs212).Int())
				_1207_cf10 = _rhs269
				_1187_v61 = _rhs270
				_1207_cf10 = _rhs271
				_1182_v56 = (func() _dafny.Int {
					if _this.F16() {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1187_v61, _1187_v61)).Cardinality())
					}
					return (_dafny.IntOfUint32((_1187_v61).Cardinality())).Times(_dafny.IntOfInt64(-682))
				})()
			} else if _source13.Is_DC6() {
				var _1215_v73 _dafny.CodePoint
				_ = _1215_v73
				_1215_v73 = _dafny.CodePoint('w')
				(globalState).F15 = _1215_v73
				_1182_v56 = _1182_v56
				var _1216_v74 _dafny.Map
				_ = _1216_v74
				_1216_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), true)
				var _1217_v75 _dafny.Map
				_ = _1217_v75
				_1217_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1182_v56, false)
				_1216_v74 = (_1216_v74).Update((func() bool {
					if (_1217_v75).Contains(_1182_v56) {
						return (_1217_v75).Get(_1182_v56).(bool)
					}
					return _this.F16()
				})(), Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState))
				(globalState).F2 = Companion_Default___.Fm0(!(_this.F16()), _this.F16(), globalState)
			} else {
				var _1218___mcc_h11 _dafny.MultiSet = _source13.Get_().(D1_DC3).Cf6
				_ = _1218___mcc_h11
				var _1219_cf6 _dafny.MultiSet = _1218___mcc_h11
				_ = _1219_cf6
				var _rhs272 _dafny.Int = _1182_v56
				_ = _rhs272
				var _rhs273 _dafny.Sequence = _1187_v61
				_ = _rhs273
				_1182_v56 = _rhs272
				_1187_v61 = _rhs273
				var _1220_v76 _dafny.Set
				_ = _1220_v76
				_1220_v76 = _dafny.SetOf(_1182_v56)
				var _1221_v77 _dafny.Map
				_ = _1221_v77
				_1221_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1219_cf6).Cardinality(), (_1220_v76).Cardinality())
				_1182_v56 = (_1182_v56).Plus(((_1221_v77).Cardinality()).Plus(_1182_v56))
				_1182_v56 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1182_v56, _1182_v56))
				_1182_v56 = (((_1219_cf6).Intersection(_1219_cf6)).Difference(_1219_cf6)).Cardinality()
			}
			var _1222_v78 _dafny.Map
			_ = _1222_v78
			_1222_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F16()) || (_this.F16()), _this.F16())
			_1222_v78 = (_1222_v78).Update(_this.F16(), _this.F16())
		}
		var _1223_v79 _dafny.Array
		_ = _1223_v79
		var _len0_35 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_35
		var _nw231 _dafny.Array
		_ = _nw231
		if _len0_35.Cmp(_dafny.Zero) == 0 {
			_nw231 = _dafny.NewArray(_len0_35)
		} else {
			var _init35 func(_dafny.Int) bool = func(_1224_i8 _dafny.Int) bool {
				return _this.F16()
			}
			_ = _init35
			var _element0_35 = _init35(_dafny.Zero)
			_ = _element0_35
			_nw231 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
			(_nw231).ArraySet1(_element0_35, 0)
			var _nativeLen0_35 = (_len0_35).Int()
			_ = _nativeLen0_35
			for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
				(_nw231).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
			}
		}
		_1223_v79 = _nw231
		for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1223_v79), 0))); ; {
			_guard_loop_8, _ok65 := _iter65()
			if !_ok65 {
				break
			}
			var _1225_i7 _dafny.Int
			_1225_i7 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_1225_i7).Sign() != -1) && ((_1225_i7).Cmp(_dafny.ArrayLen((_1223_v79), 0)) < 0)) {
				(_1223_v79).ArraySet1((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _dafny.IntOfInt64(-381))).Cardinality()).Minus(_dafny.IntOfInt64(836))).Cmp((func() _dafny.Int {
					if true {
						return _dafny.IntOfInt64(319)
					}
					return _dafny.IntOfInt64(-786)
				})()) == 0, (_1225_i7).Int())
			}
		}
		_1223_v79 = _1223_v79
		var _1226_v80 _dafny.CodePoint
		_ = _1226_v80
		_1226_v80 = _dafny.CodePoint('n')
		var _1227_v81 _dafny.Sequence
		_ = _1227_v81
		_1227_v81 = _dafny.SeqOf((func() _dafny.CodePoint {
			if _this.F16() {
				return _1226_v80
			}
			return _1226_v80
		})(), _1226_v80, _1226_v80)
		var _1228_v82 _dafny.Int
		_ = _1228_v82
		_1228_v82 = _dafny.IntOfInt64(516)
		(globalState).F15 = (_1227_v81).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_1228_v82, _1228_v82), _dafny.IntOfUint32((_1227_v81).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _1229_v83 _dafny.Sequence
		_ = _1229_v83
		_1229_v83 = _dafny.SeqOf(_this.F16())
		var _1230_v84 D0
		_ = _1230_v84
		_1230_v84 = Companion_D0_.Create_DC1_(_1227_v81)
		var _1231_v85 _dafny.Set
		_ = _1231_v85
		_1231_v85 = _dafny.SetOf((_1229_v83).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1230_v84).Dtor_cf1()).Cardinality()), _dafny.IntOfUint32((_1229_v83).Cardinality()))).Uint32()).(bool), !(!(true)), !(_this.F16()))
		if (_1231_v85).IsProperSubsetOf(_1231_v85) {
			var _1232_v86 _dafny.Array
			_ = _1232_v86
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_36
			var _nw232 _dafny.Array
			_ = _nw232
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw232 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = (func(_1233_v82 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1234_i9 _dafny.Int) _dafny.Int {
						return (_1234_i9).Plus(_1233_v82)
					}
				})(_1228_v82)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw232 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw232).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw232).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1232_v86 = _nw232
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))
			_ = _index235
			(_1232_v86).ArraySet1(_1228_v82, (_index235).Int())
			var _1235_v87 _dafny.MultiSet
			_ = _1235_v87
			_1235_v87 = _dafny.MultiSetOf(_1226_v80, (_1227_v81).Select((Companion_Default___.SafeIndex(_1228_v82, _dafny.IntOfUint32((_1227_v81).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.CodePoint('j'))
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))
			_ = _index236
			(_1232_v86).ArraySet1((_1235_v87).Cardinality(), (_index236).Int())
			var _1236_v88 _dafny.Set
			_ = _1236_v88
			_1236_v88 = _dafny.SetOf(_1228_v82, (_1232_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))).Int()).(_dafny.Int))
			var _1237_v90 D9
			_ = _1237_v90
			_1237_v90 = Companion_D9_.Create_DC27_(_this.F16(), _1236_v88)
			var _1238_v91 _dafny.Sequence
			_ = _1238_v91
			_1238_v91 = _dafny.SeqOf((_1237_v90).Dtor_cf48(), _dafny.SetOf(_1228_v82))
			var _1239_v92 _dafny.MultiSet
			_ = _1239_v92
			_1239_v92 = _dafny.MultiSetOf(_1236_v88, func() _dafny.Set {
				var _coll57 = _dafny.NewBuilder()
				_ = _coll57
				for _iter66 := _dafny.Iterate(((_this).F17()).Elements()); ; {
					_compr_57, _ok66 := _iter66()
					if !_ok66 {
						break
					}
					var _1240_v89 _dafny.Int
					_1240_v89 = interface{}(_compr_57).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((_this).F17(), _1240_v89) {
						_coll57.Add(Companion_Default___.SafeModuloInt(_1240_v89, _dafny.IntOfInt64(564)))
					}
				}
				return _coll57.ToSet()
			}(), (_1236_v88).Difference(_1236_v88), (_1238_v91).Select((Companion_Default___.SafeIndex(_1228_v82, _dafny.IntOfUint32((_1238_v91).Cardinality()))).Uint32()).(_dafny.Set), _1236_v88)
			_1239_v92 = _1239_v92
			var _1241_v93 _dafny.Array
			_ = _1241_v93
			var _nwElement0_38 _dafny.CodePoint = _1226_v80
			_ = _nwElement0_38
			var _nw233 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(10))
			_ = _nw233
			(_nw233).ArraySet1CodePoint(_nwElement0_38, 0)
			(_nw233).ArraySet1CodePoint(_1226_v80, 1)
			(_nw233).ArraySet1CodePoint(_1226_v80, 2)
			(_nw233).ArraySet1CodePoint(_1226_v80, 3)
			(_nw233).ArraySet1CodePoint(_1226_v80, 4)
			(_nw233).ArraySet1CodePoint(_dafny.CodePoint('k'), 5)
			(_nw233).ArraySet1CodePoint(_1226_v80, 6)
			(_nw233).ArraySet1CodePoint(_1226_v80, 7)
			(_nw233).ArraySet1CodePoint(_1226_v80, 8)
			(_nw233).ArraySet1CodePoint(_1226_v80, 9)
			_1241_v93 = _nw233
			_1241_v93 = _1241_v93
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))
			_ = _index237
			var _rhs274 bool = _this.F16()
			_ = _rhs274
			var _rhs275 _dafny.Int = (_1232_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))).Int()).(_dafny.Int)
			_ = _rhs275
			var _rhs276 bool = (Companion_Default___.Fm3(globalState)).Cmp((_1232_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))).Int()).(_dafny.Int)) <= 0
			_ = _rhs276
			var _lhs213 *GlobalState = globalState
			_ = _lhs213
			var _lhs214 _dafny.Array = _1232_v86
			_ = _lhs214
			var _lhs215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))
			_ = _lhs215
			var _lhs216 *GlobalState = globalState
			_ = _lhs216
			_lhs213.F5 = _rhs274
			(_lhs214).ArraySet1(_rhs275, (_lhs215).Int())
			_lhs216.F2 = _rhs276
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))
			_ = _index238
			(_1232_v86).ArraySet1((_1232_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1232_v86), 0))).Int()).(_dafny.Int), (_index238).Int())
		} else {
			var _1242_v94 D0
			_ = _1242_v94
			_1242_v94 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dlth"), _dafny.UnicodeSeqOfUtf8Bytes("ilunfi")))
			_1242_v94 = _1242_v94
			var _1243_v95 _dafny.Array
			_ = _1243_v95
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw234
			_1243_v95 = _nw234
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))
			_ = _index239
			(_1223_v79).ArraySet1((_this).Fm8(_1227_v81, globalState), (_index239).Int())
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))
			_ = _index240
			var _rhs277 _dafny.CodePoint = _1226_v80
			_ = _rhs277
			var _rhs278 _dafny.Array = _1243_v95
			_ = _rhs278
			var _rhs279 bool = _this.F16()
			_ = _rhs279
			var _lhs217 *GlobalState = globalState
			_ = _lhs217
			var _lhs218 _dafny.Array = _1223_v79
			_ = _lhs218
			var _lhs219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))
			_ = _lhs219
			_lhs217.F15 = _rhs277
			_1243_v95 = _rhs278
			(_lhs218).ArraySet1(_rhs279, (_lhs219).Int())
			if (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool) {
				var _1244_v96 _dafny.Array
				_ = _1244_v96
				var _nw235 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw235
				_1244_v96 = _nw235
				var _1245_v97 D0
				_ = _1245_v97
				_1245_v97 = Companion_D0_.Create_DC2_(_1228_v82, _1228_v82, (_dafny.SetOf(_1228_v82)).Cardinality(), _1244_v96)
				var _1246_v98 *C0
				_ = _1246_v98
				var _nw236 *C0 = New_C0_()
				_ = _nw236
				_nw236.Ctor__(_1228_v82, (_1245_v97).Dtor_cf5())
				_1246_v98 = _nw236
				var _1247_v99 _dafny.Sequence
				_ = _1247_v99
				_1247_v99 = _dafny.SeqOf(_1244_v96, _1244_v96)
				var _1248_v100 _dafny.Map
				_ = _1248_v100
				_1248_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1229_v83).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1246_v98).F26()), _dafny.IntOfUint32((_1229_v83).Cardinality()))).Uint32()).(bool), (_1246_v98).F26())
				var _1249_v101 _dafny.Map
				_ = _1249_v101
				_1249_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _1244_v96)
				var _1250_v102 _dafny.Array
				_ = _1250_v102
				var _nwElement0_39 _dafny.Array = (func() _dafny.Array {
					if (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool) {
						return (_1247_v99).Select((Companion_Default___.SafeIndex((_1248_v100).Cardinality(), _dafny.IntOfUint32((_1247_v99).Cardinality()))).Uint32()).(_dafny.Array)
					}
					return _1244_v96
				})()
				_ = _nwElement0_39
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(18))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_39, 0)
				(_nw237).ArraySet1(_1246_v98.F27, 1)
				(_nw237).ArraySet1(_1246_v98.F27, 2)
				(_nw237).ArraySet1((func() _dafny.Array {
					if (_1249_v101).Contains(false) {
						return (_1249_v101).Get(false).(_dafny.Array)
					}
					return _1246_v98.F27
				})(), 3)
				(_nw237).ArraySet1(_1244_v96, 4)
				(_nw237).ArraySet1(_1246_v98.F27, 5)
				(_nw237).ArraySet1(_1246_v98.F27, 6)
				(_nw237).ArraySet1(_1244_v96, 7)
				(_nw237).ArraySet1(_1244_v96, 8)
				(_nw237).ArraySet1(_1244_v96, 9)
				(_nw237).ArraySet1(_1246_v98.F27, 10)
				(_nw237).ArraySet1(_1244_v96, 11)
				(_nw237).ArraySet1(_1246_v98.F27, 12)
				(_nw237).ArraySet1(_1246_v98.F27, 13)
				(_nw237).ArraySet1(_1244_v96, 14)
				(_nw237).ArraySet1(_1246_v98.F27, 15)
				(_nw237).ArraySet1(_1246_v98.F27, 16)
				(_nw237).ArraySet1(_1246_v98.F27, 17)
				_1250_v102 = _nw237
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1250_v102), 0))
				_ = _index241
				(_1250_v102).ArraySet1(_1244_v96, (_index241).Int())
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1250_v102), 0))
				_ = _index242
				var _rhs280 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf(_this.F16())).Cardinality())).Plus((_dafny.Zero).Minus(_1228_v82))
				_ = _rhs280
				var _rhs281 _dafny.CodePoint = _1226_v80
				_ = _rhs281
				var _rhs282 _dafny.Array = _1246_v98.F27
				_ = _rhs282
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				var _lhs221 _dafny.Array = _1250_v102
				_ = _lhs221
				var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_1250_v102), 0))
				_ = _lhs222
				_1228_v82 = _rhs280
				_lhs220.F15 = _rhs281
				(_lhs221).ArraySet1(_rhs282, (_lhs222).Int())
				var _1251_v104 _dafny.Map
				_ = _1251_v104
				_1251_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v82, (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool))
				var _1252_v106 _dafny.Map
				_ = _1252_v106
				_1252_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					var _coll58 = _dafny.NewMapBuilder()
					_ = _coll58
					for _iter67 := _dafny.Iterate((_1251_v104).Keys().Elements()); ; {
						_compr_58, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1253_v103 _dafny.Int
						_1253_v103 = interface{}(_compr_58).(_dafny.Int)
						if (_1251_v104).Contains(_1253_v103) {
							_coll58.Add((_1253_v103).Times(_dafny.IntOfInt64(586)), (Companion_D6_.Create_DC16_(_this.F16(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool))).Dtor_cf29())
						}
					}
					return _coll58.ToMap()
				}()).Merge(func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(709), _dafny.IntOfInt64(931))); ; {
						_compr_59, _ok68 := _iter68()
						if !_ok68 {
							break
						}
						var _1254_v105 _dafny.Int
						_1254_v105 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(709)).Cmp(_1254_v105) <= 0) && ((_1254_v105).Cmp(_dafny.IntOfInt64(931)) < 0) {
							_coll59.Add(Companion_Default___.SafeDivisionInt(_1254_v105, (_1246_v98).F26()), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool))
						}
					}
					return _coll59.ToMap()
				}()), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(748), (_1246_v98).F26()))
				_1252_v106 = (_1252_v106).Update((_1251_v104).Update((_1246_v98).F26(), false), (_1246_v98).F26())
				var _1255_v107 _dafny.Set
				_ = _1255_v107
				_1255_v107 = _dafny.SetOf(_1231_v85, _1231_v85, _1231_v85, _1231_v85, _1231_v85)
				var _1256_v108 _dafny.Array
				_ = _1256_v108
				var _len0_37 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_37
				var _nw238 _dafny.Array
				_ = _nw238
				if _len0_37.Cmp(_dafny.Zero) == 0 {
					_nw238 = _dafny.NewArray(_len0_37)
				} else {
					var _init37 func(_dafny.Int) _dafny.CodePoint = (func(_1257_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1258_i10 _dafny.Int) _dafny.CodePoint {
							return _1257_v80
						}
					})(_1226_v80)
					_ = _init37
					var _element0_37 = _init37(_dafny.Zero)
					_ = _element0_37
					_nw238 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
					(_nw238).ArraySet1CodePoint(_element0_37, 0)
					var _nativeLen0_37 = (_len0_37).Int()
					_ = _nativeLen0_37
					for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
						(_nw238).ArraySet1CodePoint(_init37(_dafny.IntOf(_i0_37)), _i0_37)
					}
				}
				_1256_v108 = _nw238
				var _1259_v109 _dafny.Array
				_ = _1259_v109
				_1259_v109 = _1256_v108
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))
				_ = _index243
				var _rhs283 bool = Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState)
				_ = _rhs283
				var _rhs284 bool = !(_1255_v107).Equals(_1255_v107)
				_ = _rhs284
				var _rhs285 _dafny.Int = ((func() _dafny.Int {
					if !((_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool)) {
						return _dafny.IntOfUint32((_dafny.SeqOf(_1256_v108, _1256_v108, _1256_v108, (_1259_v109), _1256_v108)).Cardinality())
					}
					return _1228_v82
				})()).Plus(_1228_v82)
				_ = _rhs285
				var _rhs286 bool = _this.F16()
				_ = _rhs286
				var _rhs287 bool = ((_1246_v98).F26()).Cmp(_1228_v82) != 0
				_ = _rhs287
				var _lhs223 _dafny.Array = _1223_v79
				_ = _lhs223
				var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))
				_ = _lhs224
				var _lhs225 *GlobalState = globalState
				_ = _lhs225
				var _lhs226 *C5 = _this
				_ = _lhs226
				var _lhs227 *GlobalState = globalState
				_ = _lhs227
				(_lhs223).ArraySet1(_rhs283, (_lhs224).Int())
				_lhs225.F5 = _rhs284
				_1228_v82 = _rhs285
				_lhs226.F16_set_(_rhs286)
				_lhs227.F2 = _rhs287
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_1243_v95), 0))
				_ = _index244
				(_1243_v95).ArraySet1(_dafny.IntOfInt64(724), (_index244).Int())
				var _1260_v110 _dafny.Array
				_ = _1260_v110
				var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
				_ = _nw239
				_1260_v110 = _nw239
				var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_1260_v110), 0))
				_ = _index245
				(_1260_v110).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(775))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_1261_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1262_i11 _dafny.Int) _dafny.CodePoint {
						return _1261_v80
					}
				})(_1226_v80))), _1227_v81), (_index245).Int())
				var _1263_v111 _dafny.Set
				_ = _1263_v111
				_1263_v111 = _dafny.SetOf(Companion_Default___.Fm3(globalState), (_1246_v98).F26())
				var _1264_v112 D8
				_ = _1264_v112
				_1264_v112 = Companion_D8_.Create_DC21_(_1228_v82, _this.F16(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), (_1246_v98).F26())
				var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1243_v95), 0))
				_ = _index246
				(_1243_v95).ArraySet1((func() _dafny.Int {
					if (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool) {
						return (_1263_v111).Cardinality()
					}
					return (_1264_v112).Dtor_cf35()
				})(), (_index246).Int())
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_1243_v95), 0))
				_ = _index247
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_1260_v110), 0))
				_ = _index248
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1243_v95), 0))
				_ = _index249
				var _rhs288 _dafny.Int = (_1246_v98).F26()
				_ = _rhs288
				var _rhs289 _dafny.Sequence = _1227_v81
				_ = _rhs289
				var _rhs290 _dafny.Map = _1248_v100
				_ = _rhs290
				var _rhs291 _dafny.Int = (_1228_v82).Times((_1246_v98).F26())
				_ = _rhs291
				var _lhs228 _dafny.Array = _1243_v95
				_ = _lhs228
				var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_1243_v95), 0))
				_ = _lhs229
				var _lhs230 _dafny.Array = _1260_v110
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_1260_v110), 0))
				_ = _lhs231
				var _lhs232 _dafny.Array = _1243_v95
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1243_v95), 0))
				_ = _lhs233
				(_lhs228).ArraySet1(_rhs288, (_lhs229).Int())
				(_lhs230).ArraySet1(_rhs289, (_lhs231).Int())
				_1248_v100 = _rhs290
				(_lhs232).ArraySet1(_rhs291, (_lhs233).Int())
			} else {
				var _1265_v113 _dafny.Array
				_ = _1265_v113
				var _nwElement0_40 _dafny.Sequence = _1227_v81
				_ = _nwElement0_40
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(20))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_40, 0)
				(_nw240).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_1266_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1267_i12 _dafny.Int) _dafny.CodePoint {
						return _1266_v80
					}
				})(_1226_v80))), 1)
				(_nw240).ArraySet1(_1227_v81, 2)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tgjcs"), 3)
				(_nw240).ArraySet1(_1227_v81, 4)
				(_nw240).ArraySet1(_1227_v81, 5)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1227_v81, _1227_v81), 6)
				(_nw240).ArraySet1(_1227_v81, 7)
				(_nw240).ArraySet1(_1227_v81, 8)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fyruluxc"), 9)
				(_nw240).ArraySet1(_1227_v81, 10)
				(_nw240).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_1268_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1269_i13 _dafny.Int) _dafny.CodePoint {
						return _1268_v80
					}
				})(_1226_v80))), 11)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1227_v81, _1227_v81), 12)
				(_nw240).ArraySet1(_1227_v81, 13)
				(_nw240).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(229))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}((func(_1270_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1271_i14 _dafny.Int) _dafny.CodePoint {
						return _1270_v80
					}
				})(_1226_v80))), 14)
				(_nw240).ArraySet1(_1227_v81, 15)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hfvxlqw"), 16)
				(_nw240).ArraySet1(_1227_v81, 17)
				(_nw240).ArraySet1(_1227_v81, 18)
				(_nw240).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("evndkdyhw"), 19)
				_1265_v113 = _nw240
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1265_v113), 0))
				_ = _index250
				(_1265_v113).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1227_v81, _1227_v81), (_index250).Int())
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1265_v113), 0))
				_ = _index251
				var _rhs292 _dafny.Int = Companion_Default___.Fm3(globalState)
				_ = _rhs292
				var _rhs293 _dafny.Sequence = _1229_v83
				_ = _rhs293
				var _rhs294 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1227_v81, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(205))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1272_v80 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1273_i15 _dafny.Int) _dafny.CodePoint {
						return _1272_v80
					}
				})(_1226_v80))))
				_ = _rhs294
				var _lhs234 _dafny.Array = _1265_v113
				_ = _lhs234
				var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1265_v113), 0))
				_ = _lhs235
				_1228_v82 = _rhs292
				r0 = _rhs293
				(_lhs234).ArraySet1(_rhs294, (_lhs235).Int())
				var _1274_v114 _dafny.Map
				_ = _1274_v114
				_1274_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(944)).Plus(_1228_v82), _this.F16())
				var _1275_v115 _dafny.Array
				_ = _1275_v115
				var _nw241 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw241
				_1275_v115 = _nw241
				var _1276_v116 *C0
				_ = _1276_v116
				var _nw242 *C0 = New_C0_()
				_ = _nw242
				_nw242.Ctor__(_dafny.IntOfUint32(((_1265_v113).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_1265_v113), 0))).Int()).(_dafny.Sequence)).Cardinality()), _1275_v115)
				_1276_v116 = _nw242
				var _1277_v117 D9
				_ = _1277_v117
				_1277_v117 = Companion_D9_.Create_DC26_(_1228_v82, _1276_v116)
				_1274_v114 = (_1274_v114).Update((_1277_v117).Dtor_cf45(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool))
				var _1278_v118 _dafny.Array
				_ = _1278_v118
				var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
				_ = _nw243
				_1278_v118 = _nw243
				var _1279_v119 _dafny.Map
				_ = _1279_v119
				_1279_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1276_v116).F26()).Plus((_1276_v116).F26()), _1278_v118)
				_1279_v119 = (_1279_v119).Update(_1228_v82, _1278_v118)
				_1228_v82 = (_1228_v82).Plus((_1276_v116).F26())
				var _1280_v120 D8
				_ = _1280_v120
				_1280_v120 = Companion_D8_.Create_DC21_((_1276_v116).F26(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), _this.F16(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), (_1276_v116).F26())
				var _1281_v121 _dafny.MultiSet
				_ = _1281_v121
				_1281_v121 = _dafny.MultiSetOf((_1280_v120).Dtor_cf35(), _dafny.IntOfInt64(141))
				var _1282_v122 _dafny.MultiSet
				_ = _1282_v122
				_1282_v122 = _dafny.MultiSetOf(_1226_v80, _1226_v80, _1226_v80)
				_1228_v82 = ((_1281_v121).Cardinality()).Times((_1282_v122).Cardinality())
			}
			if (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool) {
				var _1283_v123 *C3
				_ = _1283_v123
				var _nw244 *C3 = New_C3_()
				_ = _nw244
				_nw244.Ctor__(_1227_v81, _this.F16(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1228_v82), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1284_v83 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1285_i16 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_1284_v83).Cardinality())
					}
				})(_1229_v83)))))
				_1283_v123 = _nw244
				var _1286_v124 _dafny.Map
				_ = _1286_v124
				_1286_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v82, _1228_v82)
				var _1287_v125 _dafny.Set
				_ = _1287_v125
				_1287_v125 = _dafny.SetOf(_1286_v124)
				var _1288_v126 _dafny.Array
				_ = _1288_v126
				var _nw245 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw245
				_1288_v126 = _nw245
				var _1289_v127 _dafny.Map
				_ = _1289_v127
				_1289_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1287_v125).IsProperSubsetOf(Companion_Default___.Fm35(_1228_v82, _dafny.CodePoint('o'), globalState)), _1288_v126)
				_1289_v127 = (_1289_v127).Update((_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), _1288_v126)
				(_this).F16_set_((_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool))
				_1228_v82 = _1228_v82
				_1228_v82 = _1228_v82
			} else {
				var _1290_v128 _dafny.Map
				_ = _1290_v128
				_1290_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ikbah"), Companion_Default___.Fm0(_this.F16(), (_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), globalState))
				var _1291_v129 _dafny.Map
				_ = _1291_v129
				_1291_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F16()) || (_this.F16()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("fhohi"), _this.F16())).Merge(_1290_v128))
				_1291_v129 = (_1291_v129).Update((_1223_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1223_v79), 0))).Int()).(bool), _1290_v128)
				var _1292_v130 _dafny.MultiSet
				_ = _1292_v130
				_1292_v130 = _dafny.MultiSetOf(_1227_v81, _1227_v81, _1227_v81)
				var _1293_v131 _dafny.Sequence
				_ = _1293_v131
				_1293_v131 = _dafny.SeqOf(_1227_v81, _1227_v81)
				var _rhs295 bool = _this.F16()
				_ = _rhs295
				var _rhs296 _dafny.Int = _1228_v82
				_ = _rhs296
				var _rhs297 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_1293_v131)).Update(_dafny.UnicodeSeqOfUtf8Bytes("ljveyjw"), Companion_Default___.Abs(Companion_Default___.SafeModuloInt(_1228_v82, _1228_v82)))
				_ = _rhs297
				var _rhs298 _dafny.Int = _1228_v82
				_ = _rhs298
				var _lhs236 *GlobalState = globalState
				_ = _lhs236
				_lhs236.F5 = _rhs295
				_1228_v82 = _rhs296
				_1292_v130 = _rhs297
				_1228_v82 = _rhs298
				(_this).F16_set_(false)
				_1228_v82 = (_1228_v82).Times((_dafny.Zero).Minus(_1228_v82))
				var _1294_v132 _dafny.MultiSet
				_ = _1294_v132
				_1294_v132 = _dafny.MultiSetOf(_1228_v82)
				(globalState).F2 = (_1294_v132).Contains(_1228_v82)
			}
			var _rhs299 _dafny.Int = _1228_v82
			_ = _rhs299
			var _rhs300 bool = ((_dafny.Zero).Minus(_1228_v82)).Cmp(_1228_v82) != 0
			_ = _rhs300
			var _lhs237 *C5 = _this
			_ = _lhs237
			_1228_v82 = _rhs299
			_lhs237.F16_set_(_rhs300)
		}
		var _1295_v133 D8
		_ = _1295_v133
		_1295_v133 = Companion_D8_.Create_DC24_(Companion_D8_.Create_DC23_())
		_1295_v133 = _1295_v133
		r0 = _dafny.SeqOf(_this.F16())
		return r0
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f21 D0
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f21 = Companion_D0_.Default()
	return &_this
}

type CompanionStruct_C6_ struct {
}

var Companion_C6_ = CompanionStruct_C6_{}

func (_this *C6) Equals(other *C6) bool {
	return _this == other
}

func (_this *C6) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C6)
	return ok && _this.Equals(other)
}

func (*C6) String() string {
	return "_module.C6"
}

func Type_C6_() _dafny.TypeDescriptor {
	return type_C6_{}
}

type type_C6_ struct {
}

func (_this type_C6_) Default() interface{} {
	return (*C6)(nil)
}

func (_this type_C6_) String() string {
	return "main.C6"
}
func (_this *C6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__(f21 D0) {
	{
		(_this)._f21 = f21
	}
}
func (_this *C6) Fm6(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.SeqOf(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lc")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg54 _dafny.Int) interface{} {
				return coer54(arg54)
			}
		}(func(_1296_i0 _dafny.Int) _dafny.Int {
			return (_dafny.SetOf(!(!(true)), false, !(false), false)).Cardinality()
		}))).Cardinality())))
	}
}
func (_this *C6) M3(p0 bool, globalState *GlobalState) {
	{
		var _1297_v0 _dafny.Sequence
		_ = _1297_v0
		_1297_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qpaxlupea")
		var _1298_v1 _dafny.Int
		_ = _1298_v1
		_1298_v1 = _dafny.IntOfInt64(76)
		var _hi8 _dafny.Int = (_1298_v1).Times(_1298_v1)
		_ = _hi8
		for _1299_i0 := _dafny.IntOfUint32((_1297_v0).Cardinality()); _1299_i0.Cmp(_hi8) < 0; _1299_i0 = _1299_i0.Plus(_dafny.One) {
			var _1300_v2 _dafny.Array
			_ = _1300_v2
			var _nw246 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw246
			_1300_v2 = _nw246
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1300_v2), 0))
			_ = _index252
			(_1300_v2).ArraySet1((_1299_i0).Plus(_1299_i0), (_index252).Int())
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1300_v2), 0))
			_ = _index253
			(_1300_v2).ArraySet1(_1298_v1, (_index253).Int())
			var _1301_v3 _dafny.Set
			_ = _1301_v3
			_1301_v3 = _dafny.SetOf(p0)
			var _1302_v4 D2
			_ = _1302_v4
			_1302_v4 = Companion_D2_.Create_DC7_(_1301_v3)
			var _1303_v5 _dafny.MultiSet
			_ = _1303_v5
			_1303_v5 = _dafny.MultiSetOf((((_1302_v4).Dtor_cf12()).Union(_1301_v3)).Cardinality(), _dafny.IntOfInt64(441))
			var _1304_v6 _dafny.CodePoint
			_ = _1304_v6
			_1304_v6 = _dafny.CodePoint('g')
			var _1305_v7 _dafny.Set
			_ = _1305_v7
			_1305_v7 = _dafny.SetOf(_1304_v6, _1304_v6, Companion_Default___.Fm2(globalState))
			_1298_v1 = (func() _dafny.Int {
				if (_1303_v5).Contains((_1305_v7).Cardinality()) {
					return (_1303_v5).Multiplicity((_1305_v7).Cardinality())
				}
				return (_1298_v1).Times(_1299_i0)
			})()
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_1300_v2), 0))
			_ = _index254
			(_1300_v2).ArraySet1((_dafny.Zero).Minus(_1299_i0), (_index254).Int())
			var _1306_v8 D1
			_ = _1306_v8
			_1306_v8 = Companion_D1_.Create_DC6_()
			_1306_v8 = _1306_v8
		}
		var _hi9 _dafny.Int = _dafny.IntOfInt64(169)
		_ = _hi9
		for _1307_i1 := _1298_v1; _1307_i1.Cmp(_hi9) < 0; _1307_i1 = _1307_i1.Plus(_dafny.One) {
			var _1308_v9 _dafny.MultiSet
			_ = _1308_v9
			_1308_v9 = _dafny.MultiSetOf(p0)
			var _1309_v10 D1
			_ = _1309_v10
			_1309_v10 = Companion_D1_.Create_DC3_(_1308_v9)
			var _pat_let_tv27 = _1308_v9
			_ = _pat_let_tv27
			var _source14 D1 = func(_pat_let35_0 D1) D1 {
				return func(_1310_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let36_0 _dafny.MultiSet) D1 {
						return func(_1311_dt__update_hcf6_h0 _dafny.MultiSet) D1 {
							return Companion_D1_.Create_DC3_(_1311_dt__update_hcf6_h0)
						}(_pat_let36_0)
					}(_pat_let_tv27)
				}(_pat_let35_0)
			}(_1309_v10)
			_ = _source14
			if _source14.Is_DC4() {
				var _1312___mcc_h0 bool = _source14.Get_().(D1_DC4).Cf7
				_ = _1312___mcc_h0
				var _1313___mcc_h1 _dafny.Int = _source14.Get_().(D1_DC4).Cf8
				_ = _1313___mcc_h1
				var _1314___mcc_h2 _dafny.Map = _source14.Get_().(D1_DC4).Cf9
				_ = _1314___mcc_h2
				var _1315_cf9 _dafny.Map = _1314___mcc_h2
				_ = _1315_cf9
				var _1316_cf8 _dafny.Int = _1313___mcc_h1
				_ = _1316_cf8
				var _1317_cf7 bool = _1312___mcc_h0
				_ = _1317_cf7
				var _1318_v11 _dafny.Set
				_ = _1318_v11
				_1318_v11 = _dafny.SetOf(_1298_v1)
				var _1319_v12 _dafny.Set
				_ = _1319_v12
				_1319_v12 = _dafny.SetOf(_1318_v11, _1318_v11)
				var _1320_v13 _dafny.Map
				_ = _1320_v13
				_1320_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(175), _1307_i1)
				(globalState).F2 = ((_dafny.SetOf(_1318_v11)).Intersection(_dafny.SetOf(func() _dafny.Set {
					var _coll60 = _dafny.NewBuilder()
					_ = _coll60
					for _iter69 := _dafny.Iterate((_1320_v13).Keys().Elements()); ; {
						_compr_60, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _1321_v14 _dafny.Int
						_1321_v14 = interface{}(_compr_60).(_dafny.Int)
						if (_1320_v13).Contains(_1321_v14) {
							_coll60.Add(Companion_Default___.SafeDivisionInt(_1321_v14, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mgmhqd")).Cardinality())))
						}
					}
					return _coll60.ToSet()
				}()))).IsProperSubsetOf(_1319_v12)
				var _1322_v15 bool
				_ = _1322_v15
				var _1323_v16 bool
				_ = _1323_v16
				var _1324_v17 bool
				_ = _1324_v17
				var _1325_v18 _dafny.Int
				_ = _1325_v18
				var _out10 bool
				_ = _out10
				var _out11 bool
				_ = _out11
				var _out12 bool
				_ = _out12
				var _out13 _dafny.Int
				_ = _out13
				_out10, _out11, _out12, _out13 = (_this).M4(globalState)
				_1322_v15 = _out10
				_1323_v16 = _out11
				_1324_v17 = _out12
				_1325_v18 = _out13
				var _1326_v19 _dafny.Sequence
				_ = _1326_v19
				_1326_v19 = _dafny.SeqOf(_1297_v0, _dafny.UnicodeSeqOfUtf8Bytes("sluvhgu"))
				var _1327_v20 _dafny.CodePoint
				_ = _1327_v20
				_1327_v20 = _dafny.CodePoint('w')
				var _1328_v21 _dafny.MultiSet
				_ = _1328_v21
				_1328_v21 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update((_1326_v19).Select((Companion_Default___.SafeIndex(_1316_cf8, _dafny.IntOfUint32((_1326_v19).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1297_v0).Cardinality()), _dafny.IntOfUint32(((_1326_v19).Select((Companion_Default___.SafeIndex(_1316_cf8, _dafny.IntOfUint32((_1326_v19).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1327_v20), _1297_v0, _1297_v0)
				_1328_v21 = _1328_v21
				var _1329_v22 _dafny.Sequence
				_ = _1329_v22
				_1329_v22 = _dafny.SeqOf(_dafny.IntOfInt64(-967))
				var _1330_v23 T0
				_ = _1330_v23
				var _nw247 *C1 = New_C1_()
				_ = _nw247
				_nw247.Ctor__(_dafny.IntOfInt64(919), _dafny.IntOfUint32((_1297_v0).Cardinality()), _1317_cf7, _1329_v22, _1316_cf8)
				_1330_v23 = _nw247
				var _1331_v24 _dafny.Map
				_ = _1331_v24
				_1331_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1323_v16, _1330_v23)
				var _1332_v25 D8
				_ = _1332_v25
				_1332_v25 = Companion_D8_.Create_DC22_(_1322_v15, _1330_v23, _1324_v17)
				var _1333_v26 _dafny.Sequence
				_ = _1333_v26
				_1333_v26 = _dafny.SeqOf(_1330_v23, (_1332_v25).Dtor_cf41())
				var _1334_v27 _dafny.Map
				_ = _1334_v27
				_1334_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1322_v15, _dafny.IntOfUint32((_1329_v22).Cardinality()))
				var _1335_v28 _dafny.Map
				_ = _1335_v28
				_1335_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T0 {
					if (_1331_v24).Contains(_1317_cf7) {
						return (_1331_v24).Get(_1317_cf7).(T0)
					}
					return (_1333_v26).Select((Companion_Default___.SafeIndex((_1334_v27).Cardinality(), _dafny.IntOfUint32((_1333_v26).Cardinality()))).Uint32()).(T0)
				})(), _1330_v23.F16())
				_1335_v28 = (_1335_v28).Update(_1330_v23, _1323_v16)
			} else if _source14.Is_DC5() {
				var _1336___mcc_h3 _dafny.Int = _source14.Get_().(D1_DC5).Cf10
				_ = _1336___mcc_h3
				var _1337___mcc_h4 bool = _source14.Get_().(D1_DC5).Cf11
				_ = _1337___mcc_h4
				var _1338_cf11 bool = _1337___mcc_h4
				_ = _1338_cf11
				var _1339_cf10 _dafny.Int = _1336___mcc_h3
				_ = _1339_cf10
				var _1340_v29 _dafny.Array
				_ = _1340_v29
				var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw248
				_1340_v29 = _nw248
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1340_v29), 0))
				_ = _index255
				(_1340_v29).ArraySet1(_1307_i1, (_index255).Int())
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1340_v29), 0))
				_ = _index256
				(_1340_v29).ArraySet1(Companion_Default___.SafeModuloInt(_1339_cf10, _1339_cf10), (_index256).Int())
				(globalState).F2 = (func() bool {
					if p0 {
						return p0
					}
					return !(p0)
				})()
				var _1341_v30 _dafny.Sequence
				_ = _1341_v30
				_1341_v30 = _dafny.SeqOf(_1338_cf11)
				var _1342_v31 _dafny.Set
				_ = _1342_v31
				_1342_v31 = _dafny.SetOf(p0, false)
				_1341_v30 = _dafny.Companion_Sequence_.Concatenate(_1341_v30, (func() _dafny.Sequence {
					if _1338_cf11 {
						return _1341_v30
					}
					return _dafny.Companion_Sequence_.Update(_1341_v30, (Companion_Default___.SafeIndex((_1342_v31).Cardinality(), _dafny.IntOfUint32((_1341_v30).Cardinality()))).Uint32(), _1338_cf11)
				})())
				var _1343_v32 _dafny.Set
				_ = _1343_v32
				_1343_v32 = _dafny.SetOf(_1298_v1)
				(globalState).F2 = Companion_Default___.Fm0((_1343_v32).IsProperSubsetOf(_1343_v32), Companion_Default___.Fm0(Companion_Default___.Fm0(p0, p0, globalState), p0, globalState), globalState)
			} else if _source14.Is_DC6() {
				var _1344_v33 _dafny.CodePoint
				_ = _1344_v33
				_1344_v33 = _dafny.CodePoint('a')
				var _pat_let_tv28 = _1344_v33
				_ = _pat_let_tv28
				(globalState).F5 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_1297_v0, (Companion_Default___.SafeIndex(_1298_v1, _dafny.IntOfUint32((_1297_v0).Cardinality()))).Uint32(), (func(_pat_let37_0 D11) D11 {
					return func(_1345_dt__update__tmp_h1 D11) D11 {
						return func(_pat_let38_0 _dafny.CodePoint) D11 {
							return func(_1346_dt__update_hcf50_h0 _dafny.CodePoint) D11 {
								return Companion_D11_.Create_DC29_(_1346_dt__update_hcf50_h0)
							}(_pat_let38_0)
						}(_pat_let_tv28)
					}(_pat_let37_0)
				}(Companion_D11_.Create_DC29_(_1344_v33))).Dtor_cf50()), _1297_v0)
				_1298_v1 = _1307_i1
				var _1347_v34 _dafny.Set
				_ = _1347_v34
				_1347_v34 = _dafny.SetOf(_1344_v33, _dafny.CodePoint('u'))
				_1347_v34 = _1347_v34
				var _1348_v35 _dafny.Array
				_ = _1348_v35
				var _nw249 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
				_ = _nw249
				_1348_v35 = _nw249
				var _1349_v36 _dafny.Sequence
				_ = _1349_v36
				_1349_v36 = _dafny.SeqOf(p0)
				var _1350_v37 _dafny.Array
				_ = _1350_v37
				var _nwElement0_41 _dafny.Int = _dafny.IntOfUint32((_1349_v36).Cardinality())
				_ = _nwElement0_41
				var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(2))
				_ = _nw250
				(_nw250).ArraySet1(_nwElement0_41, 0)
				(_nw250).ArraySet1((_dafny.Zero).Minus(_1307_i1), 1)
				_1350_v37 = _nw250
				var _1351_v38 _dafny.Map
				_ = _1351_v38
				_1351_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if p0 {
						return _1348_v35
					}
					return _1348_v35
				})(), _1350_v37)
				_1351_v38 = (_1351_v38).Update(_1348_v35, _1350_v37)
			} else {
				var _1352___mcc_h5 _dafny.MultiSet = _source14.Get_().(D1_DC3).Cf6
				_ = _1352___mcc_h5
				var _1353_cf6 _dafny.MultiSet = _1352___mcc_h5
				_ = _1353_cf6
				var _1354_v39 _dafny.Map
				_ = _1354_v39
				_1354_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1307_i1)
				var _1355_v40 _dafny.CodePoint
				_ = _1355_v40
				_1355_v40 = _dafny.CodePoint('t')
				var _1356_v41 _dafny.Map
				_ = _1356_v41
				_1356_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v1, p0)
				var _1357_v42 _dafny.Sequence
				_ = _1357_v42
				_1357_v42 = _dafny.SeqOf((_1356_v41).Cardinality())
				var _1358_v43 _dafny.Map
				_ = _1358_v43
				_1358_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1355_v40, _1357_v42)
				var _1359_v44 _dafny.Sequence
				_ = _1359_v44
				_1359_v44 = _dafny.SeqOf(_1357_v42)
				var _1360_v45 _dafny.Array
				_ = _1360_v45
				var _nwElement0_42 _dafny.Sequence = _dafny.SeqOf(_1298_v1)
				_ = _nwElement0_42
				var _nw251 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(18))
				_ = _nw251
				(_nw251).ArraySet1(_nwElement0_42, 0)
				(_nw251).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(888))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}(func(_1361_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))).Cardinality())), 1)
				(_nw251).ArraySet1(_dafny.SeqOf(_1298_v1), 2)
				(_nw251).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(925))).Uint32(), func(coer56 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_1362_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1363_i3 _dafny.Int) _dafny.Int {
						return _1362_v1
					}
				})(_1298_v1))), 3)
				(_nw251).ArraySet1(_dafny.SeqOf(_1307_i1), 4)
				(_nw251).ArraySet1((func() _dafny.Sequence {
					if (_1358_v43).Contains(_1355_v40) {
						return (_1358_v43).Get(_1355_v40).(_dafny.Sequence)
					}
					return _1357_v42
				})(), 5)
				(_nw251).ArraySet1(_1357_v42, 6)
				(_nw251).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(668)), 7)
				(_nw251).ArraySet1(_1357_v42, 8)
				(_nw251).ArraySet1(_1357_v42, 9)
				(_nw251).ArraySet1(_1357_v42, 10)
				(_nw251).ArraySet1(_1357_v42, 11)
				(_nw251).ArraySet1(_1357_v42, 12)
				(_nw251).ArraySet1((_1359_v44).Select((Companion_Default___.SafeIndex(_1298_v1, _dafny.IntOfUint32((_1359_v44).Cardinality()))).Uint32()).(_dafny.Sequence), 13)
				(_nw251).ArraySet1(_1357_v42, 14)
				(_nw251).ArraySet1(_1357_v42, 15)
				(_nw251).ArraySet1(_1357_v42, 16)
				(_nw251).ArraySet1((Companion_D3_.Create_DC9_(_1357_v42)).Dtor_cf18(), 17)
				_1360_v45 = _nw251
				var _1364_v46 D2
				_ = _1364_v46
				_1364_v46 = Companion_D2_.Create_DC8_(_1354_v39, _1360_v45, p0, false, _1307_i1)
				var _1365_v47 _dafny.Map
				_ = _1365_v47
				_1365_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1364_v46).Dtor_cf15(), _1297_v0)
				_1298_v1 = (_1307_i1).Minus((_1365_v47).Cardinality())
				var _1366_v48 _dafny.Array
				_ = _1366_v48
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_38
				var _nw252 _dafny.Array
				_ = _nw252
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw252 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Sequence = (func(_1367_p0 bool) func(_dafny.Int) _dafny.Sequence {
						return func(_1368_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_1367_p0, _1367_p0)
						}
					})(p0)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw252 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw252).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw252).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1366_v48 = _nw252
				var _1369_v49 _dafny.Sequence
				_ = _1369_v49
				_1369_v49 = _dafny.SeqOf(false, p0)
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1366_v48), 0))
				_ = _index257
				(_1366_v48).ArraySet1(_1369_v49, (_index257).Int())
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1366_v48), 0))
				_ = _index258
				(_1366_v48).ArraySet1(_1369_v49, (_index258).Int())
				var _1370_v50 _dafny.Array
				_ = _1370_v50
				var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw253
				_1370_v50 = _nw253
				var _1371_v51 _dafny.Array
				_ = _1371_v51
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw254
				_1371_v51 = _nw254
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1370_v50), 0))
				_ = _index259
				(_1370_v50).ArraySet1(_1371_v51, (_index259).Int())
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_1370_v50), 0))
				_ = _index260
				(_1370_v50).ArraySet1(_1371_v51, (_index260).Int())
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1366_v48), 0))
				_ = _index261
				(_1366_v48).ArraySet1(_1369_v49, (_index261).Int())
			}
			var _1372_v52 _dafny.MultiSet
			_ = _1372_v52
			_1372_v52 = _dafny.MultiSetOf(_1298_v1, (func() _dafny.Int {
				if p0 {
					return _1307_i1
				}
				return _1298_v1
			})(), (_dafny.IntOfInt64(-937)).Plus(_1307_i1), _1298_v1)
			var _1373_v53 _dafny.Sequence
			_ = _1373_v53
			_1373_v53 = _dafny.SeqOf(_dafny.IntOfInt64(-710), _1298_v1)
			var _1374_v54 _dafny.Map
			_ = _1374_v54
			_1374_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1373_v53)).Cardinality(), p0)
			var _1375_v55 _dafny.Map
			_ = _1375_v55
			_1375_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v54, _1307_i1)
			var _1376_v56 _dafny.Sequence
			_ = _1376_v56
			_1376_v56 = _dafny.SeqOf(_1297_v0)
			_1372_v52 = (_1372_v52).Update((_1375_v55).Cardinality(), Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1376_v56).Select((Companion_Default___.SafeIndex(_1298_v1, _dafny.IntOfUint32((_1376_v56).Cardinality()))).Uint32()).(_dafny.Sequence), _1297_v0)).Cardinality()))))
			var _1377_v57 _dafny.Array
			_ = _1377_v57
			var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw255
			_1377_v57 = _nw255
			var _1378_v58 _dafny.Map
			_ = _1378_v58
			_1378_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1377_v57, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}(func(_1379_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))
			var _1380_v59 _dafny.Map
			_ = _1380_v59
			_1380_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v1, _1377_v57)
			var _1381_v60 _dafny.Sequence
			_ = _1381_v60
			_1381_v60 = _dafny.SeqOf(p0)
			var _1382_v61 D12
			_ = _1382_v61
			_1382_v61 = Companion_D12_.Create_DC33_((func() _dafny.Array {
				if (_1380_v59).Contains(_dafny.IntOfUint32((_1381_v60).Cardinality())) {
					return (_1380_v59).Get(_dafny.IntOfUint32((_1381_v60).Cardinality())).(_dafny.Array)
				}
				return _1377_v57
			})())
			_1297_v0 = (func() _dafny.Sequence {
				if (_1378_v58).Contains((_1382_v61).Dtor_cf58()) {
					return (_1378_v58).Get((_1382_v61).Dtor_cf58()).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("utfdflun")
			})()
			var _1383_v62 _dafny.Set
			_ = _1383_v62
			_1383_v62 = _dafny.SetOf(_1307_i1)
			var _1384_v63 _dafny.Array
			_ = _1384_v63
			var _nwElement0_43 bool = p0
			_ = _nwElement0_43
			var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(20))
			_ = _nw256
			(_nw256).ArraySet1(_nwElement0_43, 0)
			(_nw256).ArraySet1(p0, 1)
			(_nw256).ArraySet1((p0) == (p0), 2)
			(_nw256).ArraySet1((_1308_v9).IsDisjointFrom(_1308_v9), 3)
			(_nw256).ArraySet1(p0, 4)
			(_nw256).ArraySet1(p0, 5)
			(_nw256).ArraySet1(!(p0), 6)
			(_nw256).ArraySet1(p0, 7)
			(_nw256).ArraySet1(p0, 8)
			(_nw256).ArraySet1(p0, 9)
			(_nw256).ArraySet1(!(!((func() bool {
				if p0 {
					return p0
				}
				return p0
			})())), 10)
			(_nw256).ArraySet1(false, 11)
			(_nw256).ArraySet1((_1298_v1).Cmp(_1298_v1) < 0, 12)
			(_nw256).ArraySet1(false, 13)
			(_nw256).ArraySet1(p0, 14)
			(_nw256).ArraySet1((p0) == (p0), 15)
			(_nw256).ArraySet1(p0, 16)
			(_nw256).ArraySet1((_1383_v62).IsProperSubsetOf(Companion_Default___.Fm31(p0, globalState)), 17)
			(_nw256).ArraySet1(p0, 18)
			(_nw256).ArraySet1(((_1373_v53).Select((Companion_Default___.SafeIndex(_1307_i1, _dafny.IntOfUint32((_1373_v53).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_1381_v60).Cardinality())) >= 0, 19)
			_1384_v63 = _nw256
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1384_v63), 0))
			_ = _index262
			(_1384_v63).ArraySet1(p0, (_index262).Int())
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1384_v63), 0))
			_ = _index263
			(_1384_v63).ArraySet1((_1372_v52).Contains(_dafny.IntOfInt64(-628)), (_index263).Int())
		}
		var _hi10 _dafny.Int = _1298_v1
		_ = _hi10
		for _1385_i6 := _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfmtbfldy")).Cardinality()); _1385_i6.Cmp(_hi10) < 0; _1385_i6 = _1385_i6.Plus(_dafny.One) {
			(globalState).F5 = (func() bool {
				if true {
					return p0
				}
				return false
			})()
			var _1386_v64 _dafny.Array
			_ = _1386_v64
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_39
			var _nw257 _dafny.Array
			_ = _nw257
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw257 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Int = (func(_1387_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1388_i7 _dafny.Int) _dafny.Int {
						return (_1388_i7).Plus(_1387_v1)
					}
				})(_1298_v1)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw257 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw257).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw257).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1386_v64 = _nw257
			var _1389_v65 _dafny.Sequence
			_ = _1389_v65
			_1389_v65 = _dafny.SeqOf(_1386_v64, _1386_v64)
			_1389_v65 = _1389_v65
			var _1390_v66 _dafny.MultiSet
			_ = _1390_v66
			_1390_v66 = _dafny.MultiSetOf(p0, p0, p0)
			var _1391_v67 D11
			_ = _1391_v67
			_1391_v67 = Companion_D11_.Create_DC30_(p0, _1385_i6)
			_1298_v1 = (func() _dafny.Int {
				if (_1390_v66).Contains((_1391_v67).Dtor_cf51()) {
					return (_1390_v66).Multiplicity((_1391_v67).Dtor_cf51())
				}
				return _1298_v1
			})()
			_1298_v1 = _1385_i6
		}
		var _hi11 _dafny.Int = Companion_Default___.SafeDivisionInt(_1298_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mkxb")).Cardinality()))
		_ = _hi11
		for _1392_i8 := _1298_v1; _1392_i8.Cmp(_hi11) < 0; _1392_i8 = _1392_i8.Plus(_dafny.One) {
			var _1393_v68 _dafny.CodePoint
			_ = _1393_v68
			_1393_v68 = _dafny.CodePoint('o')
			(globalState).F2 = !_dafny.Companion_Sequence_.Contains(_1297_v0, _1393_v68)
			var _1394_v69 _dafny.Sequence
			_ = _1394_v69
			_1394_v69 = _dafny.SeqOf(_1298_v1, _1392_i8)
			var _1395_v70 D3
			_ = _1395_v70
			_1395_v70 = Companion_D3_.Create_DC9_(_1394_v69)
			var _1396_v71 _dafny.Map
			_ = _1396_v71
			_1396_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1395_v70).Dtor_cf18())
			_1298_v1 = (_dafny.MultiSetOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1396_v71).Contains(p0) {
					return (_1396_v71).Get(p0).(_dafny.Sequence)
				}
				return (_this).Fm6(Companion_Default___.Fm3(globalState), p0, _1393_v68, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(96))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_1397_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('y')
				})), globalState)
			})()).Cardinality()))).Cardinality()
			var _1398_v72 _dafny.Map
			_ = _1398_v72
			_1398_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1392_i8)
			var _1399_v73 _dafny.Array
			_ = _1399_v73
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw258
			_1399_v73 = _nw258
			var _1400_v74 D2
			_ = _1400_v74
			_1400_v74 = Companion_D2_.Create_DC8_(_1398_v72, _1399_v73, p0, p0, _1298_v1)
			var _pat_let_tv29 = p0
			_ = _pat_let_tv29
			var _pat_let_tv30 = _1399_v73
			_ = _pat_let_tv30
			var _1401_v75 D6
			_ = _1401_v75
			_1401_v75 = Companion_D6_.Create_DC17_(func(_pat_let39_0 D2) D2 {
				return func(_1402_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let40_0 bool) D2 {
						return func(_1403_dt__update_hcf15_h0 bool) D2 {
							return func(_pat_let41_0 _dafny.Array) D2 {
								return func(_1404_dt__update_hcf14_h0 _dafny.Array) D2 {
									return Companion_D2_.Create_DC8_((_1402_dt__update__tmp_h2).Dtor_cf13(), _1404_dt__update_hcf14_h0, _1403_dt__update_hcf15_h0, (_1402_dt__update__tmp_h2).Dtor_cf16(), (_1402_dt__update__tmp_h2).Dtor_cf17())
								}(_pat_let41_0)
							}(_pat_let_tv30)
						}(_pat_let40_0)
					}(_pat_let_tv29)
				}(_pat_let39_0)
			}(_1400_v74))
			var _1405_v76 _dafny.MultiSet
			_ = _1405_v76
			_1405_v76 = _dafny.MultiSetOf(_1401_v75, _1401_v75)
			var _1406_v77 _dafny.Sequence
			_ = _1406_v77
			_1406_v77 = _dafny.SeqOf((_1405_v76).Update(_1401_v75, Companion_Default___.Abs(_1298_v1)))
			var _1407_v78 _dafny.Sequence
			_ = _1407_v78
			_1407_v78 = _1406_v77
			_1407_v78 = _1407_v78
			var _1408_v79 *C2
			_ = _1408_v79
			var _nw259 *C2 = New_C2_()
			_ = _nw259
			_nw259.Ctor__(_1298_v1)
			_1408_v79 = _nw259
			var _1409_v80 _dafny.Set
			_ = _1409_v80
			_1409_v80 = _dafny.SetOf(_1408_v79, _1408_v79, _1408_v79, _1408_v79)
			_1409_v80 = _1409_v80
		}
		var _hi12 _dafny.Int = _1298_v1
		_ = _hi12
		for _1410_i10 := _1298_v1; _1410_i10.Cmp(_hi12) < 0; _1410_i10 = _1410_i10.Plus(_dafny.One) {
			(globalState).F5 = !(false)
			if (_1298_v1).Cmp((_1298_v1).Plus(_1410_i10)) <= 0 {
				var _1411_v81 _dafny.Set
				_ = _1411_v81
				_1411_v81 = _dafny.SetOf(Companion_Default___.Fm0(p0, p0, globalState), false)
				var _1412_v82 _dafny.MultiSet
				_ = _1412_v82
				_1412_v82 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfw")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_1411_v81).Cardinality())).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm20(p0, !(p0), p0, _1410_i10, globalState)).Cardinality()), _dafny.IntOfInt64(-685), _1410_i10)
				var _1413_v83 D1
				_ = _1413_v83
				_1413_v83 = Companion_D1_.Create_DC5_(_1298_v1, false)
				var _1414_v84 *C1
				_ = _1414_v84
				var _nw260 *C1 = New_C1_()
				_ = _nw260
				_nw260.Ctor__(_1410_i10, (func() _dafny.Int {
					if (_1412_v82).Contains((_1413_v83).Dtor_cf10()) {
						return (_1412_v82).Multiplicity((_1413_v83).Dtor_cf10())
					}
					return _1410_i10
				})(), (func() bool {
					if !(p0) {
						return p0
					}
					return p0
				})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1410_i10), _dafny.SeqOf(_1298_v1, _1298_v1, (_dafny.MultiSetOf(_dafny.IntOfInt64(787))).Cardinality(), _1410_i10)), _1410_i10)
				_1414_v84 = _nw260
				var _1415_v85 _dafny.Array
				_ = _1415_v85
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_40
				var _nw261 _dafny.Array
				_ = _nw261
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw261 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Sequence = func(_1416_i11 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(483))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}(func(_1417_i12 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('p')
						}))
					}
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw261 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw261).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw261).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1415_v85 = _nw261
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1415_v85), 0))
				_ = _index264
				(_1415_v85).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1297_v0, _1297_v0), (_index264).Int())
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1415_v85), 0))
				_ = _index265
				var _rhs301 _dafny.Int = _1298_v1
				_ = _rhs301
				var _rhs302 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vbdbgrpij"), _1297_v0)
				_ = _rhs302
				var _lhs238 _dafny.Array = _1415_v85
				_ = _lhs238
				var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_1415_v85), 0))
				_ = _lhs239
				_1298_v1 = _rhs301
				(_lhs238).ArraySet1(_rhs302, (_lhs239).Int())
				_1298_v1 = _1410_i10
				var _1418_v86 _dafny.Sequence
				_ = _1418_v86
				_1418_v86 = _dafny.SeqOf((_1414_v84).F29())
				_1298_v1 = (func() _dafny.Int {
					if (_1412_v82).Contains((_1414_v84).F30()) {
						return (_1412_v82).Multiplicity((_1414_v84).F30())
					}
					return Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1418_v86, p0)).Cardinality(), _dafny.IntOfInt64(-373))
				})()
				var _1419_v87 _dafny.Array
				_ = _1419_v87
				var _nw262 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw262
				_1419_v87 = _nw262
				var _1420_v88 *C0
				_ = _1420_v88
				var _nw263 *C0 = New_C0_()
				_ = _nw263
				_nw263.Ctor__((_1414_v84).F30(), _1419_v87)
				_1420_v88 = _nw263
			} else {
				(globalState).F15 = Companion_Default___.Fm2(globalState)
				var _1421_v89 _dafny.Array
				_ = _1421_v89
				var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw264
				_1421_v89 = _nw264
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1421_v89), 0))
				_ = _index266
				(_1421_v89).ArraySet1(_1297_v0, (_index266).Int())
				var _1422_v90 _dafny.CodePoint
				_ = _1422_v90
				_1422_v90 = _dafny.CodePoint('a')
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_1421_v89), 0))
				_ = _index267
				(_1421_v89).ArraySet1(_dafny.Companion_Sequence_.Update(_1297_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-864), _dafny.IntOfUint32((_1297_v0).Cardinality()))).Uint32(), _1422_v90), (_index267).Int())
				var _1423_v91 _dafny.Set
				_ = _1423_v91
				_1423_v91 = _dafny.SetOf((_dafny.SetOf(_1298_v1, Companion_Default___.Fm3(globalState))).Cardinality())
				var _1424_v92 _dafny.Sequence
				_ = _1424_v92
				_1424_v92 = _dafny.SeqOf(p0)
				var _1425_v93 _dafny.Map
				_ = _1425_v93
				_1425_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("prrne"), _1410_i10)
				var _1426_v94 _dafny.Map
				_ = _1426_v94
				_1426_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1424_v92).Select((Companion_Default___.SafeIndex(_1298_v1, _dafny.IntOfUint32((_1424_v92).Cardinality()))).Uint32()).(bool), ((_1425_v93).Update(_dafny.UnicodeSeqOfUtf8Bytes("pn"), _1410_i10)).Cardinality())
				var _1427_v95 _dafny.Array
				_ = _1427_v95
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_41
				var _nw265 _dafny.Array
				_ = _nw265
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw265 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) _dafny.Sequence = (func(_1428_i10 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1429_i13 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_1428_i10)
						}
					})(_1410_i10)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw265 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw265).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw265).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1427_v95 = _nw265
				var _1430_v96 D2
				_ = _1430_v96
				_1430_v96 = Companion_D2_.Create_DC8_(_1426_v94, _1427_v95, p0, p0, _1410_i10)
				var _1431_v97 _dafny.Set
				_ = _1431_v97
				_1431_v97 = _dafny.SetOf(_1430_v96, _1430_v96)
				var _1432_v98 D9
				_ = _1432_v98
				_1432_v98 = Companion_D9_.Create_DC27_(p0, (func() _dafny.Set {
					if p0 {
						return _1423_v91
					}
					return _dafny.SetOf((_1431_v97).Cardinality())
				})())
				var _1433_v99 _dafny.MultiSet
				_ = _1433_v99
				_1433_v99 = _dafny.MultiSetOf(p0)
				var _rhs303 bool = ((_1433_v99).Cardinality()).Cmp(_dafny.IntOfInt64(214)) >= 0
				_ = _rhs303
				var _rhs304 D9 = _1432_v98
				_ = _rhs304
				var _rhs305 _dafny.Int = _1410_i10
				_ = _rhs305
				var _lhs240 *GlobalState = globalState
				_ = _lhs240
				_lhs240.F2 = _rhs303
				_1432_v98 = _rhs304
				_1298_v1 = _rhs305
				(globalState).F5 = p0
				_1422_v90 = _1422_v90
			}
			var _1434_v100 _dafny.Map
			_ = _1434_v100
			_1434_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), (_dafny.Zero).Minus(_1298_v1))
			_1298_v1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1297_v0).Cardinality()), (_1434_v100).Cardinality()), Companion_Default___.SafeModuloInt(_1298_v1, _1410_i10)))
			var _1435_v101 _dafny.Array
			_ = _1435_v101
			var _nwElement0_44 bool = p0
			_ = _nwElement0_44
			var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(8))
			_ = _nw266
			(_nw266).ArraySet1(_nwElement0_44, 0)
			(_nw266).ArraySet1(p0, 1)
			(_nw266).ArraySet1(p0, 2)
			(_nw266).ArraySet1(p0, 3)
			(_nw266).ArraySet1(p0, 4)
			(_nw266).ArraySet1(p0, 5)
			(_nw266).ArraySet1(false, 6)
			(_nw266).ArraySet1(p0, 7)
			_1435_v101 = _nw266
			var _1436_v102 _dafny.Array
			_ = _1436_v102
			var _nwElement0_45 _dafny.Array = _1435_v101
			_ = _nwElement0_45
			var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(6))
			_ = _nw267
			(_nw267).ArraySet1(_nwElement0_45, 0)
			(_nw267).ArraySet1(((_this).F21()).Dtor_cf5(), 1)
			(_nw267).ArraySet1(_1435_v101, 2)
			(_nw267).ArraySet1(_1435_v101, 3)
			(_nw267).ArraySet1(_1435_v101, 4)
			(_nw267).ArraySet1(_1435_v101, 5)
			_1436_v102 = _nw267
			var _1437_v103 D6
			_ = _1437_v103
			_1437_v103 = Companion_D6_.Create_DC15_(_1436_v102)
			var _pat_let_tv31 = _1436_v102
			_ = _pat_let_tv31
			var _1438_v104 _dafny.Map
			_ = _1438_v104
			_1438_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(705), func(_pat_let42_0 D6) D6 {
				return func(_1439_dt__update__tmp_h3 D6) D6 {
					return func(_pat_let43_0 _dafny.Array) D6 {
						return func(_1440_dt__update_hcf28_h0 _dafny.Array) D6 {
							return Companion_D6_.Create_DC15_(_1440_dt__update_hcf28_h0)
						}(_pat_let43_0)
					}(_pat_let_tv31)
				}(_pat_let42_0)
			}(_1437_v103))
			_1438_v104 = (_1438_v104).Update((_1410_i10).Minus(_1410_i10), _1437_v103)
		}
		var _1441_v105 _dafny.MultiSet
		_ = _1441_v105
		_1441_v105 = _dafny.MultiSetOf(_1298_v1)
		var _1442_v106 D4
		_ = _1442_v106
		_1442_v106 = Companion_D4_.Create_DC12_(p0, !(p0), _1298_v1)
		var _1443_v107 _dafny.MultiSet
		_ = _1443_v107
		_1443_v107 = _dafny.MultiSetOf(p0)
		var _1444_v108 _dafny.Sequence
		_ = _1444_v108
		_1444_v108 = _dafny.SeqOf(_1298_v1)
		var _1445_v109 *C3
		_ = _1445_v109
		var _nw268 *C3 = New_C3_()
		_ = _nw268
		_nw268.Ctor__(_1297_v0, p0, _1444_v108)
		_1445_v109 = _nw268
		var _1446_v110 _dafny.Map
		_ = _1446_v110
		_1446_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_v109, p0)
		var _1447_v111 D8
		_ = _1447_v111
		_1447_v111 = Companion_D8_.Create_DC21_(_1298_v1, p0, true, p0, _1298_v1)
		var _1448_v112 _dafny.Set
		_ = _1448_v112
		_1448_v112 = _dafny.SetOf(_1298_v1)
		var _1449_v113 _dafny.Map
		_ = _1449_v113
		_1449_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1298_v1)
		var _1450_v114 D11
		_ = _1450_v114
		_1450_v114 = Companion_D11_.Create_DC31_(p0, _1298_v1, _1298_v1, p0)
		var _1451_v115 _dafny.Sequence
		_ = _1451_v115
		_1451_v115 = _dafny.SeqOf((_1450_v114).Dtor_cf53())
		var _1452_v116 _dafny.Map
		_ = _1452_v116
		_1452_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(505), _dafny.IntOfInt64(617))
		var _1453_v117 _dafny.Array
		_ = _1453_v117
		var _nwElement0_46 _dafny.Int = (_1298_v1).Plus(_dafny.IntOfUint32((_1297_v0).Cardinality()))
		_ = _nwElement0_46
		var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(23))
		_ = _nw269
		(_nw269).ArraySet1(_nwElement0_46, 0)
		(_nw269).ArraySet1(_1298_v1, 1)
		(_nw269).ArraySet1(_1298_v1, 2)
		(_nw269).ArraySet1(_1298_v1, 3)
		(_nw269).ArraySet1(_1298_v1, 4)
		(_nw269).ArraySet1((func() _dafny.Int {
			if (_1441_v105).Contains(_1298_v1) {
				return (_1441_v105).Multiplicity(_1298_v1)
			}
			return (_1442_v106).Dtor_cf23()
		})(), 5)
		(_nw269).ArraySet1((_1298_v1).Plus((_1443_v107).Cardinality()), 6)
		(_nw269).ArraySet1(_1298_v1, 7)
		(_nw269).ArraySet1(_1298_v1, 8)
		(_nw269).ArraySet1((_dafny.Zero).Minus(_1298_v1), 9)
		(_nw269).ArraySet1(_1298_v1, 10)
		(_nw269).ArraySet1((_dafny.Zero).Minus(_1298_v1), 11)
		(_nw269).ArraySet1(((_1446_v110).Merge((_1446_v110).Update(_1445_v109, p0))).Cardinality(), 12)
		(_nw269).ArraySet1((_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1447_v111).Dtor_cf37(), (_1448_v112).Cardinality())).Merge(_1449_v113)).Cardinality()), 13)
		(_nw269).ArraySet1((_dafny.Zero).Minus(_1298_v1), 14)
		(_nw269).ArraySet1(_1298_v1, 15)
		(_nw269).ArraySet1(((Companion_Default___.Fm36(_1298_v1, p0, globalState)).Intersection(_1443_v107)).Cardinality(), 16)
		(_nw269).ArraySet1((func() _dafny.Int {
			if p0 {
				return _1298_v1
			}
			return _1298_v1
		})(), 17)
		(_nw269).ArraySet1((_1298_v1).Plus(_dafny.IntOfUint32((_1451_v115).Cardinality())), 18)
		(_nw269).ArraySet1(_1298_v1, 19)
		(_nw269).ArraySet1((_1452_v116).Cardinality(), 20)
		(_nw269).ArraySet1(Companion_Default___.SafeDivisionInt(_1298_v1, _dafny.IntOfInt64(153)), 21)
		(_nw269).ArraySet1(_1298_v1, 22)
		_1453_v117 = _nw269
		var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1453_v117), 0))
		_ = _index268
		(_1453_v117).ArraySet1(_1298_v1, (_index268).Int())
		var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_1453_v117), 0))
		_ = _index269
		(_1453_v117).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm3(globalState)), (_index269).Int())
	}
}
func (_this *C6) M4(globalState *GlobalState) (bool, bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1454_v0 bool
		_ = _1454_v0
		_1454_v0 = false
		if _1454_v0 {
			var _source15 D0 = (_this).F21()
			_ = _source15
			if _source15.Is_DC1() {
				var _1455___mcc_h0 _dafny.Sequence = _source15.Get_().(D0_DC1).Cf1
				_ = _1455___mcc_h0
				var _1456_cf1 _dafny.Sequence = _1455___mcc_h0
				_ = _1456_cf1
				r3 = Companion_Default___.Fm3(globalState)
				var _1457_v1 _dafny.Int
				_ = _1457_v1
				_1457_v1 = _dafny.IntOfInt64(408)
				r3 = (_1457_v1).Times(_1457_v1)
				var _1458_v2 _dafny.Map
				_ = _1458_v2
				_1458_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(18), _1457_v1)
				var _1459_v3 _dafny.Map
				_ = _1459_v3
				_1459_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v0, (_1458_v2).Cardinality())
				var _1460_v4 _dafny.MultiSet
				_ = _1460_v4
				_1460_v4 = _dafny.MultiSetOf(_1459_v3)
				var _1461_v5 _dafny.Sequence
				_ = _1461_v5
				_1461_v5 = _dafny.SeqOf(_1460_v4, _dafny.MultiSetOf(_1459_v3))
				var _1462_v6 D4
				_ = _1462_v6
				_1462_v6 = Companion_D4_.Create_DC11_((_1461_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.IntOfUint32((_1461_v5).Cardinality()))).Uint32()).(_dafny.MultiSet))
				_1462_v6 = _1462_v6
				var _1463_v7 _dafny.Sequence
				_ = _1463_v7
				_1463_v7 = _dafny.SeqOf(_1457_v1)
				var _1464_v8 _dafny.CodePoint
				_ = _1464_v8
				_1464_v8 = _dafny.CodePoint('i')
				var _1465_v9 D0
				_ = _1465_v9
				_1465_v9 = Companion_D0_.Create_DC1_(_dafny.Companion_Sequence_.Update(_1456_cf1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("mmnwsjn"), (Companion_Default___.SafeIndex(_1457_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mmnwsjn")).Cardinality()))).Uint32(), _1464_v8)).Cardinality()), _dafny.IntOfUint32((_1456_cf1).Cardinality()))).Uint32(), _1464_v8))
				var _1466_v10 _dafny.Map
				_ = _1466_v10
				_1466_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v1, _1465_v9)
				var _1467_v11 *C1
				_ = _1467_v11
				var _nw270 *C1 = New_C1_()
				_ = _nw270
				_nw270.Ctor__(_dafny.IntOfInt64(931), _1457_v1, _1454_v0, _1463_v7, (_1466_v10).Cardinality())
				_1467_v11 = _nw270
				var _1468_v12 _dafny.Map
				_ = _1468_v12
				_1468_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1457_v1, _1467_v11)
				_1468_v12 = (_1468_v12).Update((_1458_v2).Cardinality(), _1467_v11)
			} else if _source15.Is_DC2() {
				var _1469___mcc_h1 _dafny.Int = _source15.Get_().(D0_DC2).Cf2
				_ = _1469___mcc_h1
				var _1470___mcc_h2 _dafny.Int = _source15.Get_().(D0_DC2).Cf3
				_ = _1470___mcc_h2
				var _1471___mcc_h3 _dafny.Int = _source15.Get_().(D0_DC2).Cf4
				_ = _1471___mcc_h3
				var _1472___mcc_h4 _dafny.Array = _source15.Get_().(D0_DC2).Cf5
				_ = _1472___mcc_h4
				var _1473_cf5 _dafny.Array = _1472___mcc_h4
				_ = _1473_cf5
				var _1474_cf4 _dafny.Int = _1471___mcc_h3
				_ = _1474_cf4
				var _1475_cf3 _dafny.Int = _1470___mcc_h2
				_ = _1475_cf3
				var _1476_cf2 _dafny.Int = _1469___mcc_h1
				_ = _1476_cf2
				var _1477_v13 _dafny.Map
				_ = _1477_v13
				_1477_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1475_cf3, _1454_v0)
				var _1478_v14 _dafny.Sequence
				_ = _1478_v14
				_1478_v14 = _dafny.UnicodeSeqOfUtf8Bytes("vxaubieu")
				r2 = (func() bool {
					if (_1477_v13).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((_1476_cf2).Plus(_dafny.IntOfUint32((_1478_v14).Cardinality()))))) {
						return (_1477_v13).Get((_dafny.Zero).Minus((_dafny.Zero).Minus((_1476_cf2).Plus(_dafny.IntOfUint32((_1478_v14).Cardinality()))))).(bool)
					}
					return !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-21))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg60 _dafny.Int) interface{} {
							return coer60(arg60)
						}
					}((func(_1479_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1480_i0 _dafny.Int) _dafny.Int {
							return _1479_cf4
						}
					})(_1474_cf4))), _dafny.SeqOf(_1476_cf2))
				})()
				var _1481_v15 _dafny.Map
				_ = _1481_v15
				_1481_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v0, _1478_v14)
				var _1482_v16 _dafny.Map
				_ = _1482_v16
				_1482_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1481_v15).Cardinality(), _1475_cf3)
				var _1483_v17 _dafny.Set
				_ = _1483_v17
				_1483_v17 = _dafny.SetOf(_1482_v16)
				_1483_v17 = (_1483_v17).Union(_1483_v17)
				_1477_v13 = _1477_v13
				var _1484_v18 _dafny.Sequence
				_ = _1484_v18
				_1484_v18 = _dafny.SeqOf(_1454_v0, !(_1454_v0))
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_1473_cf5), 0))
				_ = _index270
				(_1473_cf5).ArraySet1((_1484_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1484_v18).Cardinality()), _dafny.IntOfUint32((_1484_v18).Cardinality()))).Uint32()).(bool), (_index270).Int())
				var _1485_v19 _dafny.CodePoint
				_ = _1485_v19
				_1485_v19 = _dafny.CodePoint('m')
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_1473_cf5), 0))
				_ = _index271
				var _rhs306 _dafny.CodePoint = _1485_v19
				_ = _rhs306
				var _rhs307 bool = _1454_v0
				_ = _rhs307
				var _rhs308 _dafny.Int = ((_dafny.Zero).Minus(_1476_cf2)).Times(Companion_Default___.SafeModuloInt(_1474_cf4, _dafny.IntOfUint32((_1484_v18).Cardinality())))
				_ = _rhs308
				var _rhs309 _dafny.CodePoint = _1485_v19
				_ = _rhs309
				var _lhs241 *GlobalState = globalState
				_ = _lhs241
				var _lhs242 _dafny.Array = _1473_cf5
				_ = _lhs242
				var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(41), _dafny.ArrayLen((_1473_cf5), 0))
				_ = _lhs243
				var _lhs244 *GlobalState = globalState
				_ = _lhs244
				_lhs241.F15 = _rhs306
				(_lhs242).ArraySet1(_rhs307, (_lhs243).Int())
				_1475_cf3 = _rhs308
				_lhs244.F15 = _rhs309
			} else {
				var _1486___mcc_h5 _dafny.Sequence = _source15.Get_().(D0_DC0).Cf0
				_ = _1486___mcc_h5
				var _1487_cf0 _dafny.Sequence = _1486___mcc_h5
				_ = _1487_cf0
				(globalState).F5 = _1454_v0
				_1487_cf0 = _1487_cf0
				r3 = _dafny.IntOfUint32((_1487_cf0).Cardinality())
				var _1488_v20 _dafny.Array
				_ = _1488_v20
				var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(14))
				_ = _nw271
				_1488_v20 = _nw271
				var _1489_v21 _dafny.Int
				_ = _1489_v21
				_1489_v21 = _dafny.IntOfInt64(-675)
				var _1490_v22 _dafny.MultiSet
				_ = _1490_v22
				_1490_v22 = _dafny.MultiSetOf(_1489_v21, _1489_v21)
				var _1491_v23 _dafny.Sequence
				_ = _1491_v23
				_1491_v23 = _dafny.SeqOf(_1489_v21, _1489_v21, _1489_v21)
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1488_v20), 0))
				_ = _index272
				(_1488_v20).ArraySet1((_dafny.MultiSetOf((_1490_v22).Cardinality())).Intersection(_dafny.MultiSetFromSeq(_1491_v23)), (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_1488_v20), 0))
				_ = _index273
				(_1488_v20).ArraySet1(_1490_v22, (_index273).Int())
			}
			var _1492_v24 _dafny.Array
			_ = _1492_v24
			var _nw272 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
			_ = _nw272
			_1492_v24 = _nw272
			var _1493_v25 _dafny.Int
			_ = _1493_v25
			_1493_v25 = _dafny.IntOfInt64(645)
			var _1494_v26 *C4
			_ = _1494_v26
			var _nw273 *C4 = New_C4_()
			_ = _nw273
			_nw273.Ctor__(_dafny.IntOfInt64(52), _1454_v0)
			_1494_v26 = _nw273
			var _1495_v27 _dafny.CodePoint
			_ = _1495_v27
			_1495_v27 = _dafny.CodePoint('u')
			var _1496_v28 _dafny.Map
			_ = _1496_v28
			_1496_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1494_v26, _dafny.CodePoint('m'))
			var _1497_v29 _dafny.Sequence
			_ = _1497_v29
			_1497_v29 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1494_v26, _1495_v27), _1496_v28, _1496_v28)
			var _1498_v30 _dafny.Sequence
			_ = _1498_v30
			_1498_v30 = _dafny.SeqOf(((_1497_v29).Select((Companion_Default___.SafeIndex((_1494_v26).F22(), _dafny.IntOfUint32((_1497_v29).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
			var _1499_v31 *C1
			_ = _1499_v31
			var _nw274 *C1 = New_C1_()
			_ = _nw274
			_nw274.Ctor__(_1493_v25, _1493_v25, true, _1498_v30, _dafny.IntOfUint32((_1498_v30).Cardinality()))
			_1499_v31 = _nw274
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_1492_v24), 0))
			_ = _index274
			(_1492_v24).ArraySet1(_1499_v31, (_index274).Int())
			var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_1492_v24), 0))
			_ = _index275
			(_1492_v24).ArraySet1(_1499_v31, (_index275).Int())
			var _1500_v32 _dafny.Array
			_ = _1500_v32
			var _nw275 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw275
			_1500_v32 = _nw275
			var _1501_v33 _dafny.MultiSet
			_ = _1501_v33
			_1501_v33 = _dafny.MultiSetOf(_1500_v32, _1500_v32)
			var _1502_v34 _dafny.Array
			_ = _1502_v34
			var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw276
			_1502_v34 = _nw276
			var _1503_v35 _dafny.Set
			_ = _1503_v35
			_1503_v35 = _dafny.SetOf(_1502_v34)
			var _1504_v36 D11
			_ = _1504_v36
			_1504_v36 = Companion_D11_.Create_DC30_(_1494_v26.F23, (_1494_v26).F22())
			var _1505_v37 _dafny.MultiSet
			_ = _1505_v37
			_1505_v37 = _dafny.MultiSetOf(!((_1504_v36).Dtor_cf51()), _1494_v26.F23, _1494_v26.F23)
			var _1506_v38 _dafny.Set
			_ = _1506_v38
			_1506_v38 = _dafny.SetOf(_1494_v26.F23, !(_1454_v0))
			var _1507_v39 _dafny.Map
			_ = _1507_v39
			_1507_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1494_v26.F23, (_1494_v26).F22())
			var _1508_v40 _dafny.Array
			_ = _1508_v40
			var _nwElement0_47 _dafny.Int = _1493_v25
			_ = _nwElement0_47
			var _nw277 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(16))
			_ = _nw277
			(_nw277).ArraySet1(_nwElement0_47, 0)
			(_nw277).ArraySet1((_1499_v31).F30(), 1)
			(_nw277).ArraySet1((_1501_v33).Cardinality(), 2)
			(_nw277).ArraySet1((_1499_v31).Fm15(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(324))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}((func(_1509_v31 *C1) func(_dafny.Int) _dafny.Int {
				return func(_1510_i1 _dafny.Int) _dafny.Int {
					return (_1509_v31).F30()
				}
			})(_1499_v31))), (_1503_v35).Cardinality(), _1493_v25, _1495_v27, globalState), 3)
			(_nw277).ArraySet1(_1493_v25, 4)
			(_nw277).ArraySet1((_1499_v31).F30(), 5)
			(_nw277).ArraySet1(_1493_v25, 6)
			(_nw277).ArraySet1((_1505_v37).Cardinality(), 7)
			(_nw277).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if _1494_v26.F23 {
					return _1493_v25
				}
				return (_1506_v38).Cardinality()
			})()), 8)
			(_nw277).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1500_v32, _1454_v0)).Cardinality(), 9)
			(_nw277).ArraySet1(_dafny.IntOfInt64(32), 10)
			(_nw277).ArraySet1((_1499_v31).F30(), 11)
			(_nw277).ArraySet1(_dafny.IntOfInt64(-898), 12)
			(_nw277).ArraySet1((_1499_v31).F29(), 13)
			(_nw277).ArraySet1((_1494_v26).F22(), 14)
			(_nw277).ArraySet1((func() _dafny.Int {
				if (_1507_v39).Contains(_1454_v0) {
					return (_1507_v39).Get(_1454_v0).(_dafny.Int)
				}
				return _dafny.IntOfInt64(116)
			})(), 15)
			_1508_v40 = _nw277
			_1502_v34 = _1502_v34
			var _1511_v41 _dafny.MultiSet
			_ = _1511_v41
			_1511_v41 = _dafny.MultiSetOf((_1499_v31).F30(), (_1499_v31).F29())
			var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1502_v34), 0))
			_ = _index276
			(_1502_v34).ArraySet1(((_1511_v41).Intersection(_1511_v41)).Cardinality(), (_index276).Int())
			var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1502_v34), 0))
			_ = _index277
			(_1502_v34).ArraySet1((_1499_v31).F29(), (_index277).Int())
			var _1512_v42 _dafny.Sequence
			_ = _1512_v42
			_1512_v42 = _dafny.UnicodeSeqOfUtf8Bytes("rhpqpdgw")
			_1493_v25 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1512_v42).Cardinality()))
		} else {
			var _1513_v43 _dafny.Sequence
			_ = _1513_v43
			_1513_v43 = _dafny.UnicodeSeqOfUtf8Bytes("aebefuri")
			var _1514_v44 D0
			_ = _1514_v44
			_1514_v44 = Companion_D0_.Create_DC0_(_1513_v43)
			_1514_v44 = _1514_v44
			var _1515_v45 _dafny.Int
			_ = _1515_v45
			_1515_v45 = _dafny.IntOfInt64(700)
			var _1516_v46 _dafny.Sequence
			_ = _1516_v46
			_1516_v46 = _dafny.SeqOf(_1515_v45)
			var _1517_v47 *C4
			_ = _1517_v47
			var _nw278 *C4 = New_C4_()
			_ = _nw278
			_nw278.Ctor__((_dafny.MultiSetOf(_dafny.IntOfUint32((_1516_v46).Cardinality()))).Cardinality(), (func() bool {
				if _1454_v0 {
					return _1454_v0
				}
				return Companion_Default___.Fm0(_1454_v0, _1454_v0, globalState)
			})())
			_1517_v47 = _nw278
			var _1518_v48 *C2
			_ = _1518_v48
			var _nw279 *C2 = New_C2_()
			_ = _nw279
			_nw279.Ctor__(_dafny.IntOfUint32((_1516_v46).Cardinality()))
			_1518_v48 = _nw279
			var _1519_v49 _dafny.Map
			_ = _1519_v49
			_1519_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v0, _1517_v47.F23)
			var _1520_v50 _dafny.Sequence
			_ = _1520_v50
			_1520_v50 = _dafny.SeqOf((func() bool {
				if (_1519_v49).Contains(_1454_v0) {
					return (_1519_v49).Get(_1454_v0).(bool)
				}
				return _1454_v0
			})(), _1454_v0)
			var _1521_v51 _dafny.Map
			_ = _1521_v51
			_1521_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1520_v50).Cardinality()), _1518_v48)
			var _1522_v52 _dafny.Map
			_ = _1522_v52
			_1522_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v0, (_1521_v51).Cardinality())
			var _1523_v53 _dafny.CodePoint
			_ = _1523_v53
			_1523_v53 = _dafny.CodePoint('l')
			var _1524_v54 _dafny.Map
			_ = _1524_v54
			_1524_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1523_v53, _1454_v0)
			var _1525_v55 *C1
			_ = _1525_v55
			var _nw280 *C1 = New_C1_()
			_ = _nw280
			_nw280.Ctor__((_1522_v52).Cardinality(), (_1518_v48).F25(), (func() bool {
				if (_1524_v54).Contains(_1523_v53) {
					return (_1524_v54).Get(_1523_v53).(bool)
				}
				return _1517_v47.F23
			})(), _1516_v46, (_1518_v48).F25())
			_1525_v55 = _nw280
			var _1526_v56 *C3
			_ = _1526_v56
			var _nw281 *C3 = New_C3_()
			_ = _nw281
			_nw281.Ctor__(_1513_v43, _1517_v47.F23, (func() _dafny.Sequence {
				if _1454_v0 {
					return _1516_v46
				}
				return (_this).Fm6(_dafny.IntOfInt64(822), !(!(_1517_v47.F23)), _1523_v53, _1513_v43, globalState)
			})())
			_1526_v56 = _nw281
		}
		var _1527_v57 _dafny.CodePoint
		_ = _1527_v57
		_1527_v57 = _dafny.CodePoint('h')
		var _1528_v58 _dafny.Int
		_ = _1528_v58
		_1528_v58 = _dafny.IntOfInt64(461)
		var _1529_v59 _dafny.Sequence
		_ = _1529_v59
		_1529_v59 = _dafny.SeqOf(_1454_v0)
		var _1530_v60 T0
		_ = _1530_v60
		var _nw282 *C1 = New_C1_()
		_ = _nw282
		_nw282.Ctor__(_1528_v58, _1528_v58, _1454_v0, _dafny.SeqOf((_dafny.MultiSetFromSeq(_1529_v59)).Cardinality(), _1528_v58, _dafny.IntOfInt64(-126)), _1528_v58)
		_1530_v60 = _nw282
		var _1531_v61 D8
		_ = _1531_v61
		_1531_v61 = Companion_D8_.Create_DC22_(_1454_v0, _1530_v60, _1454_v0)
		var _1532_v62 _dafny.Map
		_ = _1532_v62
		_1532_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1527_v57, (_1531_v61).Dtor_cf40())
		var _1533_v63 D9
		_ = _1533_v63
		_1533_v63 = Companion_D9_.Create_DC25_(_1532_v62)
		var _1534_v64 _dafny.Set
		_ = _1534_v64
		_1534_v64 = _dafny.SetOf(_1533_v63, _1533_v63, _1533_v63)
		var _1535_v65 _dafny.Map
		_ = _1535_v65
		_1535_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1533_v63, Companion_Default___.Fm3(globalState))
		var _1536_v67 *C5
		_ = _1536_v67
		var _nw283 *C5 = New_C5_()
		_ = _nw283
		_nw283.Ctor__(!(_1534_v64).Equals(func() _dafny.Set {
			var _coll61 = _dafny.NewBuilder()
			_ = _coll61
			for _iter70 := _dafny.Iterate((_1535_v65).Keys().Elements()); ; {
				_compr_61, _ok70 := _iter70()
				if !_ok70 {
					break
				}
				var _1537_v66 D9
				_1537_v66 = interface{}(_compr_61).(D9)
				if (_1535_v65).Contains(_1537_v66) {
					_coll61.Add(_1537_v66)
				}
			}
			return _coll61.ToSet()
		}()), _dafny.SeqOf(_1528_v58, _1528_v58, _1528_v58, _1528_v58))
		_1536_v67 = _nw283
		_1528_v58 = Companion_Default___.SafeDivisionInt(_1528_v58, _1528_v58)
		var _1538_v68 _dafny.Sequence
		_ = _1538_v68
		_1538_v68 = _dafny.UnicodeSeqOfUtf8Bytes("kp")
		var _1539_v69 D0
		_ = _1539_v69
		_1539_v69 = Companion_D0_.Create_DC1_(_1538_v68)
		var _1540_v70 _dafny.MultiSet
		_ = _1540_v70
		_1540_v70 = _dafny.MultiSetOf(_dafny.IntOfInt64(987))
		_1539_v69 = Companion_Default___.Fm37(_1454_v0, (_1540_v70).Cardinality(), _1454_v0, globalState)
		var _hi13 _dafny.Int = _1528_v58
		_ = _hi13
		for _1541_i2 := Companion_Default___.Fm3(globalState); _1541_i2.Cmp(_hi13) < 0; _1541_i2 = _1541_i2.Plus(_dafny.One) {
			var _1542_v71 D3
			_ = _1542_v71
			_1542_v71 = Companion_D3_.Create_DC10_((_1529_v59).Select((Companion_Default___.SafeIndex(_1541_i2, _dafny.IntOfUint32((_1529_v59).Cardinality()))).Uint32()).(bool))
			var _source16 D3 = _1542_v71
			_ = _source16
			if _source16.Is_DC10() {
				var _1543___mcc_h6 bool = _source16.Get_().(D3_DC10).Cf19
				_ = _1543___mcc_h6
				var _1544_cf19 bool = _1543___mcc_h6
				_ = _1544_cf19
				var _1545_v72 _dafny.Map
				_ = _1545_v72
				_1545_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1528_v58, Companion_Default___.SafeDivisionInt(_1528_v58, (_dafny.Zero).Minus(_1541_i2)))
				_1545_v72 = (_1545_v72).Update(_1541_i2, ((_1530_v60).F17()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1538_v68).Cardinality()), _dafny.IntOfUint32(((_1530_v60).F17()).Cardinality()))).Uint32()).(_dafny.Int))
				r3 = _1528_v58
				r3 = _1528_v58
				var _1546_v73 _dafny.Set
				_ = _1546_v73
				_1546_v73 = _dafny.SetOf(_1528_v58)
				var _1547_v74 _dafny.Set
				_ = _1547_v74
				_1547_v74 = _dafny.SetOf(_1546_v73)
				var _1548_v75 _dafny.Map
				_ = _1548_v75
				_1548_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1527_v57, _1527_v57)
				var _1549_v76 _dafny.Array
				_ = _1549_v76
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_42
				var _nw284 _dafny.Array
				_ = _nw284
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw284 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) bool = (func(_1550_v60 T0) func(_dafny.Int) bool {
						return func(_1551_i3 _dafny.Int) bool {
							return !(_1550_v60.F16())
						}
					})(_1530_v60)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw284 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw284).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw284).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1549_v76 = _nw284
				var _1552_v77 _dafny.Map
				_ = _1552_v77
				_1552_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1541_i2, _1549_v76)
				var _1553_v78 _dafny.Array
				_ = _1553_v78
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_43
				var _nw285 _dafny.Array
				_ = _nw285
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw285 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Int = (func(_1554_i2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1555_i4 _dafny.Int) _dafny.Int {
							return (_1555_i4).Times(_1554_i2)
						}
					})(_1541_i2)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw285 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw285).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw285).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1553_v78 = _nw285
				var _1556_v79 D12
				_ = _1556_v79
				_1556_v79 = Companion_D12_.Create_DC34_((func() _dafny.Array {
					if (_1552_v77).Contains(_dafny.IntOfInt64(63)) {
						return (_1552_v77).Get(_dafny.IntOfInt64(63)).(_dafny.Array)
					}
					return _1549_v76
				})(), _1541_i2, _1553_v78, _1527_v57)
				var _1557_v80 _dafny.Array
				_ = _1557_v80
				var _nwElement0_48 bool = _1530_v60.F16()
				_ = _nwElement0_48
				var _nw286 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(26))
				_ = _nw286
				(_nw286).ArraySet1(_nwElement0_48, 0)
				(_nw286).ArraySet1(_1544_cf19, 1)
				(_nw286).ArraySet1(_1530_v60.F16(), 2)
				(_nw286).ArraySet1(!(_1530_v60.F16()), 3)
				(_nw286).ArraySet1((_1541_i2).Cmp(Companion_Default___.Fm3(globalState)) > 0, 4)
				(_nw286).ArraySet1((_1547_v74).Equals(_1547_v74), 5)
				(_nw286).ArraySet1(Companion_Default___.Fm0(false, true, globalState), 6)
				(_nw286).ArraySet1(Companion_Default___.Fm0(_1454_v0, false, globalState), 7)
				(_nw286).ArraySet1((_1528_v58).Cmp(_1541_i2) >= 0, 8)
				(_nw286).ArraySet1(false, 9)
				(_nw286).ArraySet1(_1544_cf19, 10)
				(_nw286).ArraySet1(_1544_cf19, 11)
				(_nw286).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_1527_v57, (func() _dafny.CodePoint {
					if (_1548_v75).Contains((_1556_v79).Dtor_cf62()) {
						return (_1548_v75).Get((_1556_v79).Dtor_cf62()).(_dafny.CodePoint)
					}
					return Companion_Default___.Fm2(globalState)
				})(), _1527_v57, _1527_v57, _1527_v57), _1538_v68), 12)
				(_nw286).ArraySet1(!(_1454_v0) || (!(_1454_v0)), 13)
				(_nw286).ArraySet1((func() bool {
					if _1530_v60.F16() {
						return _1530_v60.F16()
					}
					return _1530_v60.F16()
				})(), 14)
				(_nw286).ArraySet1(_1530_v60.F16(), 15)
				(_nw286).ArraySet1(_1530_v60.F16(), 16)
				(_nw286).ArraySet1(_1530_v60.F16(), 17)
				(_nw286).ArraySet1(_1544_cf19, 18)
				(_nw286).ArraySet1(!(_1530_v60.F16()) || (_1530_v60.F16()), 19)
				(_nw286).ArraySet1(_1454_v0, 20)
				(_nw286).ArraySet1((_1536_v67).Fm8(_1538_v68, globalState), 21)
				(_nw286).ArraySet1(false, 22)
				(_nw286).ArraySet1(false, 23)
				(_nw286).ArraySet1(false, 24)
				(_nw286).ArraySet1((func() bool {
					if !(!(_1454_v0)) {
						return _1530_v60.F16()
					}
					return _1454_v0
				})(), 25)
				_1557_v80 = _nw286
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1549_v76), 0))
				_ = _index278
				(_1549_v76).ArraySet1(_1530_v60.F16(), (_index278).Int())
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1549_v76), 0))
				_ = _index279
				(_1549_v76).ArraySet1(_1544_cf19, (_index279).Int())
			} else {
				var _1558___mcc_h7 _dafny.Sequence = _source16.Get_().(D3_DC9).Cf18
				_ = _1558___mcc_h7
				var _1559_cf18 _dafny.Sequence = _1558___mcc_h7
				_ = _1559_cf18
				(globalState).F15 = _1527_v57
				_1540_v70 = _1540_v70
				var _1560_v81 _dafny.MultiSet
				_ = _1560_v81
				_1560_v81 = _dafny.MultiSetOf(_1454_v0, _1454_v0, false)
				var _1561_v82 D1
				_ = _1561_v82
				_1561_v82 = Companion_D1_.Create_DC3_((_dafny.MultiSetFromSeq(_1529_v59)).Intersection(_1560_v81))
				_1561_v82 = (func() D1 {
					if _1530_v60.F16() {
						return _1561_v82
					}
					return _1561_v82
				})()
				var _1562_v83 _dafny.Map
				_ = _1562_v83
				_1562_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1528_v58, _1454_v0)
				var _1563_v84 _dafny.Map
				_ = _1563_v84
				_1563_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1527_v57, (_1562_v83).Merge(_1562_v83))
				_1563_v84 = (_1563_v84).Update(_1527_v57, (Companion_Default___.Fm38(_1533_v63, _1541_i2, globalState)).Merge(_1562_v83))
			}
			var _1564_v85 *C1
			_ = _1564_v85
			var _nw287 *C1 = New_C1_()
			_ = _nw287
			_nw287.Ctor__(_1541_i2, _1528_v58, _1530_v60.F16(), (_1530_v60).F17(), _1528_v58)
			_1564_v85 = _nw287
			var _1565_v86 D1
			_ = _1565_v86
			_1565_v86 = Companion_D1_.Create_DC5_(_1528_v58, _1454_v0)
			_1454_v0 = ((_1565_v86).Dtor_cf10()).Cmp(Companion_Default___.SafeDivisionInt((_1564_v85).F30(), (_dafny.SetOf((_1536_v67).Fm8(_1538_v68, globalState), true)).Cardinality())) > 0
			var _1566_v87 _dafny.Set
			_ = _1566_v87
			_1566_v87 = _dafny.SetOf(_1454_v0, _1530_v60.F16(), _1530_v60.F16())
			var _1567_v88 D11
			_ = _1567_v88
			_1567_v88 = Companion_D11_.Create_DC31_(_1454_v0, (_1564_v85).F29(), (_dafny.Zero).Minus((_1564_v85).Fm14(_1541_i2, _1528_v58, _1527_v57, globalState)), _1454_v0)
			var _1568_v89 _dafny.Set
			_ = _1568_v89
			_1568_v89 = _dafny.SetOf(_1567_v88)
			var _1569_v90 _dafny.Sequence
			_ = _1569_v90
			_1569_v90 = _dafny.SeqOf(_1568_v89)
			var _1570_v91 _dafny.Map
			_ = _1570_v91
			_1570_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1564_v85).F29(), _1530_v60.F16())
			var _1571_v92 _dafny.Array
			_ = _1571_v92
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_44
			var _nw288 _dafny.Array
			_ = _nw288
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw288 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) bool = (func(_1572_v60 T0) func(_dafny.Int) bool {
					return func(_1573_i5 _dafny.Int) bool {
						return _1572_v60.F16()
					}
				})(_1530_v60)
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw288 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw288).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw288).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_1571_v92 = _nw288
			var _1574_v93 *C0
			_ = _1574_v93
			var _nw289 *C0 = New_C0_()
			_ = _nw289
			_nw289.Ctor__(_1528_v58, _1571_v92)
			_1574_v93 = _nw289
			var _1575_v94 _dafny.Array
			_ = _1575_v94
			var _nwElement0_49 _dafny.Int = (_1566_v87).Cardinality()
			_ = _nwElement0_49
			var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(9))
			_ = _nw290
			(_nw290).ArraySet1(_nwElement0_49, 0)
			(_nw290).ArraySet1(_1541_i2, 1)
			(_nw290).ArraySet1(_1541_i2, 2)
			(_nw290).ArraySet1(((_1569_v90).Select((Companion_Default___.SafeIndex((_1564_v85).F29(), _dafny.IntOfUint32((_1569_v90).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), 3)
			(_nw290).ArraySet1(_1528_v58, 4)
			(_nw290).ArraySet1(_1528_v58, 5)
			(_nw290).ArraySet1(_1528_v58, 6)
			(_nw290).ArraySet1((_1564_v85).F29(), 7)
			(_nw290).ArraySet1((Companion_D9_.Create_DC26_((_1570_v91).Cardinality(), _1574_v93)).Dtor_cf45(), 8)
			_1575_v94 = _nw290
			var _1576_v95 D6
			_ = _1576_v95
			_1576_v95 = Companion_D6_.Create_DC18_(_1575_v94)
			var _source17 D6 = _1576_v95
			_ = _source17
			if _source17.Is_DC16() {
				var _1577___mcc_h8 bool = _source17.Get_().(D6_DC16).Cf29
				_ = _1577___mcc_h8
				var _1578___mcc_h9 bool = _source17.Get_().(D6_DC16).Cf30
				_ = _1578___mcc_h9
				var _1579_cf30 bool = _1578___mcc_h9
				_ = _1579_cf30
				var _1580_cf29 bool = _1577___mcc_h8
				_ = _1580_cf29
				var _1581_v96 *C4
				_ = _1581_v96
				var _nw291 *C4 = New_C4_()
				_ = _nw291
				_nw291.Ctor__((_1564_v85).F29(), _1579_cf30)
				_1581_v96 = _nw291
				r2 = _1579_cf30
				var _1582_v97 _dafny.Map
				_ = _1582_v97
				_1582_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), (_dafny.Zero).Minus(_1541_i2))
				var _1583_v98 D11
				_ = _1583_v98
				_1583_v98 = Companion_D11_.Create_DC29_(_1527_v57)
				_1582_v97 = (_1582_v97).Update((_1583_v98).Dtor_cf50(), (_dafny.Zero).Minus((_1564_v85).F29()))
				_1528_v58 = (Companion_D1_.Create_DC5_((_1564_v85).F29(), false)).Dtor_cf10()
			} else if _source17.Is_DC17() {
				var _1584___mcc_h10 D2 = _source17.Get_().(D6_DC17).Cf31
				_ = _1584___mcc_h10
				var _1585_cf31 D2 = _1584___mcc_h10
				_ = _1585_cf31
				_1528_v58 = (_1564_v85).F30()
				var _1586_v99 _dafny.Array
				_ = _1586_v99
				var _nw292 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
				_ = _nw292
				_1586_v99 = _nw292
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_1586_v99), 0))
				_ = _index280
				(_1586_v99).ArraySet1(_1566_v87, (_index280).Int())
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_1586_v99), 0))
				_ = _index281
				(_1586_v99).ArraySet1(_1566_v87, (_index281).Int())
				_1528_v58 = Companion_Default___.SafeDivisionInt((_1564_v85).F29(), (_1564_v85).F30())
				r3 = (((func() _dafny.Map {
					var _coll62 = _dafny.NewMapBuilder()
					_ = _coll62
					for _iter71 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(119), _dafny.IntOfInt64(403))); ; {
						_compr_62, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _1587_v100 _dafny.Int
						_1587_v100 = interface{}(_compr_62).(_dafny.Int)
						if ((_dafny.IntOfInt64(119)).Cmp(_1587_v100) <= 0) && ((_1587_v100).Cmp(_dafny.IntOfInt64(403)) < 0) {
							_coll62.Add((_1587_v100).Times((_1564_v85).F29()), _1530_v60.F16())
						}
					}
					return _coll62.ToMap()
				}()).Merge((_1570_v91).Update((_1574_v93).F26(), _1454_v0))).Cardinality()).Plus((_1564_v85).F29())
			} else if _source17.Is_DC18() {
				var _1588___mcc_h11 _dafny.Array = _source17.Get_().(D6_DC18).Cf32
				_ = _1588___mcc_h11
				var _1589_cf32 _dafny.Array = _1588___mcc_h11
				_ = _1589_cf32
				var _1590_v101 _dafny.Map
				_ = _1590_v101
				_1590_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1564_v85).F30(), _dafny.CodePoint('m'))
				var _1591_v102 _dafny.Array
				_ = _1591_v102
				var _nwElement0_50 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (_1590_v101).Contains((_1564_v85).F29()) {
						return (_1590_v101).Get((_1564_v85).F29()).(_dafny.CodePoint)
					}
					return _1527_v57
				})()
				_ = _nwElement0_50
				var _nw293 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(16))
				_ = _nw293
				(_nw293).ArraySet1CodePoint(_nwElement0_50, 0)
				(_nw293).ArraySet1CodePoint(_1527_v57, 1)
				(_nw293).ArraySet1CodePoint(_1527_v57, 2)
				(_nw293).ArraySet1CodePoint(_1527_v57, 3)
				(_nw293).ArraySet1CodePoint(_1527_v57, 4)
				(_nw293).ArraySet1CodePoint(_1527_v57, 5)
				(_nw293).ArraySet1CodePoint(_1527_v57, 6)
				(_nw293).ArraySet1CodePoint(_1527_v57, 7)
				(_nw293).ArraySet1CodePoint(_1527_v57, 8)
				(_nw293).ArraySet1CodePoint(_1527_v57, 9)
				(_nw293).ArraySet1CodePoint(_1527_v57, 10)
				(_nw293).ArraySet1CodePoint(_1527_v57, 11)
				(_nw293).ArraySet1CodePoint(_dafny.CodePoint('h'), 12)
				(_nw293).ArraySet1CodePoint(_1527_v57, 13)
				(_nw293).ArraySet1CodePoint(_1527_v57, 14)
				(_nw293).ArraySet1CodePoint(_1527_v57, 15)
				_1591_v102 = _nw293
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_1591_v102), 0))
				_ = _index282
				(_1591_v102).ArraySet1CodePoint(_1527_v57, (_index282).Int())
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_1591_v102), 0))
				_ = _index283
				(_1591_v102).ArraySet1CodePoint(_1527_v57, (_index283).Int())
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1575_v94), 0))
				_ = _index284
				(_1575_v94).ArraySet1(((_1574_v93).F26()).Times(_1541_i2), (_index284).Int())
				var _1592_v103 _dafny.Array
				_ = _1592_v103
				var _nw294 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw294
				_1592_v103 = _nw294
				var _1593_v104 _dafny.Map
				_ = _1593_v104
				_1593_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1564_v85).F30(), (_1530_v60).F17())
				var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1575_v94), 0))
				_ = _index285
				var _rhs310 _dafny.Int = (_dafny.Zero).Minus((_1564_v85).Fm15((func() _dafny.Sequence {
					if (_1593_v104).Contains((_1540_v70).Cardinality()) {
						return (_1593_v104).Get((_1540_v70).Cardinality()).(_dafny.Sequence)
					}
					return (_1530_v60).F17()
				})(), (_1528_v58).Minus((_1574_v93).F26()), (_1574_v93).F26(), _1527_v57, globalState))
				_ = _rhs310
				var _rhs311 _dafny.Array = _1592_v103
				_ = _rhs311
				var _rhs312 _dafny.Int = (_1564_v85).F30()
				_ = _rhs312
				var _lhs245 _dafny.Array = _1575_v94
				_ = _lhs245
				var _lhs246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1575_v94), 0))
				_ = _lhs246
				(_lhs245).ArraySet1(_rhs310, (_lhs246).Int())
				_1592_v103 = _rhs311
				_1528_v58 = _rhs312
				var _1594_v105 *C4
				_ = _1594_v105
				var _nw295 *C4 = New_C4_()
				_ = _nw295
				_nw295.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(241))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1595_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1596_i6 _dafny.Int) _dafny.CodePoint {
						return _1595_v57
					}
				})(_1527_v57))), _dafny.Companion_Sequence_.Update(_1538_v68, (Companion_Default___.SafeIndex((_1564_v85).F29(), _dafny.IntOfUint32((_1538_v68).Cardinality()))).Uint32(), _1527_v57))).Cardinality()), (func() bool {
					if _1530_v60.F16() {
						return false
					}
					return _1454_v0
				})())
				_1594_v105 = _nw295
				_1528_v58 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_1575_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.ArrayLen((_1575_v94), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(100)), (_1564_v85).F29())
			} else {
				var _1597___mcc_h12 _dafny.Array = _source17.Get_().(D6_DC15).Cf28
				_ = _1597___mcc_h12
				var _1598_cf28 _dafny.Array = _1597___mcc_h12
				_ = _1598_cf28
				var _1599_v106 _dafny.MultiSet
				_ = _1599_v106
				_1599_v106 = _dafny.MultiSetOf(_1530_v60.F16())
				var _1600_v107 D0
				_ = _1600_v107
				_1600_v107 = Companion_D0_.Create_DC2_((_1564_v85).F29(), (_dafny.IntOfUint32((Companion_Default___.Fm20(!(_1454_v0), _1530_v60.F16(), !(_1530_v60.F16()), (Companion_D4_.Create_DC12_(_1454_v0, true, (_1599_v106).Cardinality())).Dtor_cf23(), globalState)).Cardinality())).Times((_1564_v85).F30()), _dafny.IntOfUint32((_1529_v59).Cardinality()), _1574_v93.F27)
				_1600_v107 = _1600_v107
				var _1601_v108 _dafny.Array
				_ = _1601_v108
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_45
				var _nw296 _dafny.Array
				_ = _nw296
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw296 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.CodePoint = (func(_1602_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1603_i7 _dafny.Int) _dafny.CodePoint {
							return _1602_v57
						}
					})(_1527_v57)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw296 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw296).ArraySet1CodePoint(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw296).ArraySet1CodePoint(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1601_v108 = _nw296
				var _1604_v109 T1
				_ = _1604_v109
				var _nw297 *C1 = New_C1_()
				_ = _nw297
				_nw297.Ctor__((_1574_v93).F26(), (_1574_v93).F26(), _1530_v60.F16(), _dafny.Companion_Sequence_.Update((_1530_v60).F17(), (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_1601_v108, _1601_v108)).Cardinality(), _dafny.IntOfUint32(((_1530_v60).F17()).Cardinality()))).Uint32(), _dafny.IntOfInt64(875)), (_1564_v85).F30())
				_1604_v109 = _nw297
				var _1605_v110 _dafny.Map
				_ = _1605_v110
				_1605_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1604_v109, !((func() bool {
					if _1530_v60.F16() {
						return _1604_v109.F16()
					}
					return _1604_v109.F16()
				})()))
				var _1606_v111 _dafny.Sequence
				_ = _1606_v111
				_1606_v111 = _dafny.SeqOf(_1604_v109)
				_1605_v110 = (_1605_v110).Update((_1606_v111).Select((Companion_Default___.SafeIndex((_1564_v85).F29(), _dafny.IntOfUint32((_1606_v111).Cardinality()))).Uint32()).(T1), _1604_v109.F16())
				(_1530_v60).F16_set_(_1530_v60.F16())
				var _1607_v112 _dafny.Sequence
				_ = _1607_v112
				_1607_v112 = _dafny.SeqOf(_1538_v68)
				_1529_v59 = _dafny.Companion_Sequence_.Update(_1529_v59, (Companion_Default___.SafeIndex((_1564_v85).F30(), _dafny.IntOfUint32((_1529_v59).Cardinality()))).Uint32(), ((_1564_v85).F29()).Cmp(_dafny.IntOfUint32((_1607_v112).Cardinality())) != 0)
			}
		}
		var _1608_v113 _dafny.Map
		_ = _1608_v113
		_1608_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1528_v58, _1528_v58)
		var _hi14 _dafny.Int = _1528_v58
		_ = _hi14
		for _1609_i8 := ((_1608_v113).Cardinality()).Times(_dafny.IntOfInt64(-870)); _1609_i8.Cmp(_hi14) < 0; _1609_i8 = _1609_i8.Plus(_dafny.One) {
			var _1610_v114 D8
			_ = _1610_v114
			_1610_v114 = Companion_D8_.Create_DC21_(_1528_v58, (_1609_i8).Cmp(_1609_i8) <= 0, (_1528_v58).Cmp(_1528_v58) != 0, _dafny.Companion_Sequence_.Contains(_1529_v59, _1454_v0), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1538_v68).Cardinality())))
			var _source18 D8 = _1610_v114
			_ = _source18
			if _source18.Is_DC21() {
				var _1611___mcc_h13 _dafny.Int = _source18.Get_().(D8_DC21).Cf35
				_ = _1611___mcc_h13
				var _1612___mcc_h14 bool = _source18.Get_().(D8_DC21).Cf36
				_ = _1612___mcc_h14
				var _1613___mcc_h15 bool = _source18.Get_().(D8_DC21).Cf37
				_ = _1613___mcc_h15
				var _1614___mcc_h16 bool = _source18.Get_().(D8_DC21).Cf38
				_ = _1614___mcc_h16
				var _1615___mcc_h17 _dafny.Int = _source18.Get_().(D8_DC21).Cf39
				_ = _1615___mcc_h17
				var _1616_cf39 _dafny.Int = _1615___mcc_h17
				_ = _1616_cf39
				var _1617_cf38 bool = _1614___mcc_h16
				_ = _1617_cf38
				var _1618_cf37 bool = _1613___mcc_h15
				_ = _1618_cf37
				var _1619_cf36 bool = _1612___mcc_h14
				_ = _1619_cf36
				var _1620_cf35 _dafny.Int = _1611___mcc_h13
				_ = _1620_cf35
				r3 = _1616_cf39
				_1616_cf39 = _1616_cf39
				var _1621_v115 _dafny.Map
				_ = _1621_v115
				_1621_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _1620_cf35)
				var _1622_v116 _dafny.Sequence
				_ = _1622_v116
				_1622_v116 = _dafny.SeqOf(_1621_v115)
				_1622_v116 = _1622_v116
				(_1530_v60).F16_set_(_1530_v60.F16())
			} else if _source18.Is_DC22() {
				var _1623___mcc_h18 bool = _source18.Get_().(D8_DC22).Cf40
				_ = _1623___mcc_h18
				var _1624___mcc_h19 T0 = _source18.Get_().(D8_DC22).Cf41
				_ = _1624___mcc_h19
				var _1625___mcc_h20 bool = _source18.Get_().(D8_DC22).Cf42
				_ = _1625___mcc_h20
				var _1626_cf42 bool = _1625___mcc_h20
				_ = _1626_cf42
				var _1627_cf41 T0 = _1624___mcc_h19
				_ = _1627_cf41
				var _1628_cf40 bool = _1623___mcc_h18
				_ = _1628_cf40
				(globalState).F5 = _1454_v0
				var _1629_v117 _dafny.Map
				_ = _1629_v117
				_1629_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_1628_cf40, _1628_cf40, globalState), _1538_v68)
				var _1630_v118 _dafny.Sequence
				_ = _1630_v118
				_1630_v118 = _dafny.SeqOf((_1629_v117).Cardinality(), (_1528_v58).Times(_1528_v58), _1528_v58)
				_1630_v118 = (func() _dafny.Sequence {
					if false {
						return _dafny.Companion_Sequence_.Update((_1627_cf41).F17(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1627_cf41).F17()).Cardinality())), _dafny.IntOfUint32(((_1627_cf41).F17()).Cardinality()))).Uint32(), _1609_i8)
					}
					return Companion_Default___.Fm17(globalState)
				})()
				var _1631_v119 *C3
				_ = _1631_v119
				var _nw298 *C3 = New_C3_()
				_ = _nw298
				_nw298.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("gj"), (_1536_v67).Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(248))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1632_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1633_i9 _dafny.Int) _dafny.CodePoint {
						return _1632_v57
					}
				})(_1527_v57))), globalState), (_1530_v60).F17())
				_1631_v119 = _nw298
				var _1634_v120 _dafny.Map
				_ = _1634_v120
				_1634_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1528_v58, _1530_v60.F16())
				var _1635_v121 *C1
				_ = _1635_v121
				var _nw299 *C1 = New_C1_()
				_ = _nw299
				_nw299.Ctor__(_1609_i8, _dafny.IntOfUint32(((_1631_v119).F24()).Cardinality()), _1530_v60.F16(), _1630_v118, Companion_Default___.SafeModuloInt((_1634_v120).Cardinality(), (_1540_v70).Cardinality()))
				_1635_v121 = _nw299
			} else if _source18.Is_DC23() {
				var _1636_v122 _dafny.Map
				_ = _1636_v122
				_1636_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1609_i8, _dafny.UnicodeSeqOfUtf8Bytes("jvrpwreug"))
				var _1637_v123 _dafny.Map
				_ = _1637_v123
				_1637_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-794), _1530_v60.F16())
				var _1638_v125 _dafny.Sequence
				_ = _1638_v125
				_1638_v125 = _dafny.SeqOf(_1608_v113, func() _dafny.Map {
					var _coll63 = _dafny.NewMapBuilder()
					_ = _coll63
					for _iter72 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(918), _dafny.IntOfInt64(620))); ; {
						_compr_63, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _1639_v124 _dafny.Int
						_1639_v124 = interface{}(_compr_63).(_dafny.Int)
						if ((_dafny.IntOfInt64(918)).Cmp(_1639_v124) <= 0) && ((_1639_v124).Cmp(_dafny.IntOfInt64(620)) < 0) {
							_coll63.Add((_1639_v124).Times(_1609_i8), _1609_i8)
						}
					}
					return _coll63.ToMap()
				}())
				var _1640_v126 *C3
				_ = _1640_v126
				var _nw300 *C3 = New_C3_()
				_ = _nw300
				_nw300.Ctor__(_1538_v68, _1530_v60.F16(), (_1530_v60).F17())
				_1640_v126 = _nw300
				var _1641_v127 D12
				_ = _1641_v127
				_1641_v127 = Companion_D12_.Create_DC35_(false, _1530_v60.F16(), (_1638_v125).Select((Companion_Default___.SafeIndex(_1528_v58, _dafny.IntOfUint32((_1638_v125).Cardinality()))).Uint32()).(_dafny.Map), _1530_v60.F16(), _1640_v126)
				_1636_v122 = (_1636_v122).Update((_dafny.Zero).Minus(((_1637_v123).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1528_v58, (_1641_v127).Dtor_cf63()))).Cardinality()), (_1640_v126).F24())
				var _1642_v128 _dafny.Sequence
				_ = _1642_v128
				var _out14 _dafny.Sequence
				_ = _out14
				_out14 = (_1536_v67).M5(globalState)
				_1642_v128 = _out14
				_1538_v68 = Companion_Default___.Fm1(globalState)
				var _1643_v129 _dafny.Array
				_ = _1643_v129
				var _nw301 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw301
				_1643_v129 = _nw301
				var _1644_v130 D6
				_ = _1644_v130
				_1644_v130 = Companion_D6_.Create_DC15_(_1643_v129)
				var _1645_v131 _dafny.Array
				_ = _1645_v131
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_46
				var _nw302 _dafny.Array
				_ = _nw302
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw302 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) bool = (func(_1646_v60 T0) func(_dafny.Int) bool {
						return func(_1647_i10 _dafny.Int) bool {
							return (_1646_v60.F16()) || (_1646_v60.F16())
						}
					})(_1530_v60)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw302 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw302).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw302).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1645_v131 = _nw302
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1645_v131), 0))
				_ = _index286
				(_1645_v131).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((_1640_v126).F24(), _dafny.UnicodeSeqOfUtf8Bytes("uncqkk")), (_index286).Int())
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_1645_v131), 0))
				_ = _index287
				(_1645_v131).ArraySet1(Companion_Default___.Fm0(_1530_v60.F16(), true, globalState), (_index287).Int())
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1645_v131), 0))
				_ = _index288
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_1645_v131), 0))
				_ = _index289
				var _rhs313 D6 = _1644_v130
				_ = _rhs313
				var _rhs314 bool = _1530_v60.F16()
				_ = _rhs314
				var _rhs315 bool = !((((_dafny.Zero).Minus(_1528_v58)).Cmp(_dafny.IntOfUint32((_1642_v128).Cardinality())) == 0) && (_1454_v0))
				_ = _rhs315
				var _lhs247 _dafny.Array = _1645_v131
				_ = _lhs247
				var _lhs248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1645_v131), 0))
				_ = _lhs248
				var _lhs249 _dafny.Array = _1645_v131
				_ = _lhs249
				var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_1645_v131), 0))
				_ = _lhs250
				_1644_v130 = _rhs313
				(_lhs247).ArraySet1(_rhs314, (_lhs248).Int())
				(_lhs249).ArraySet1(_rhs315, (_lhs250).Int())
			} else if _source18.Is_DC20() {
				var _1648___mcc_h21 _dafny.Sequence = _source18.Get_().(D8_DC20).Cf34
				_ = _1648___mcc_h21
				var _1649_cf34 _dafny.Sequence = _1648___mcc_h21
				_ = _1649_cf34
				var _1650_v132 _dafny.Map
				_ = _1650_v132
				_1650_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1609_i8).Plus(_1609_i8)), !(!((func() bool {
					if _1530_v60.F16() {
						return _1530_v60.F16()
					}
					return (_1536_v67).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("pwy"), globalState)
				})())))
				_1650_v132 = (_1650_v132).Update(_1609_i8, true)
				r3 = (_dafny.IntOfInt64(794)).Plus(_1609_i8)
				var _1651_v133 _dafny.Array
				_ = _1651_v133
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_47
				var _nw303 _dafny.Array
				_ = _nw303
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw303 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Sequence = (func(_1652_v60 T0, _1653_v68 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1654_i11 _dafny.Int) _dafny.Sequence {
							return (func() _dafny.Sequence {
								if _1652_v60.F16() {
									return _1653_v68
								}
								return (Companion_D0_.Create_DC0_(_1653_v68)).Dtor_cf0()
							})()
						}
					})(_1530_v60, _1538_v68)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw303 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw303).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw303).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1651_v133 = _nw303
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))
				_ = _index290
				(_1651_v133).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1538_v68, _1538_v68), (_index290).Int())
				var _1655_v134 _dafny.Set
				_ = _1655_v134
				_1655_v134 = _dafny.SetOf((func() bool {
					if (_1650_v132).Contains(_dafny.IntOfInt64(836)) {
						return (_1650_v132).Get(_dafny.IntOfInt64(836)).(bool)
					}
					return _1454_v0
				})(), _1530_v60.F16(), _1530_v60.F16())
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))
				_ = _index291
				var _rhs316 bool = _1530_v60.F16()
				_ = _rhs316
				var _rhs317 _dafny.Sequence = _1538_v68
				_ = _rhs317
				var _rhs318 _dafny.Set = _1655_v134
				_ = _rhs318
				var _rhs319 _dafny.Int = _1528_v58
				_ = _rhs319
				var _lhs251 *GlobalState = globalState
				_ = _lhs251
				var _lhs252 _dafny.Array = _1651_v133
				_ = _lhs252
				var _lhs253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))
				_ = _lhs253
				_lhs251.F5 = _rhs316
				(_lhs252).ArraySet1(_rhs317, (_lhs253).Int())
				_1655_v134 = _rhs318
				r3 = _rhs319
				var _1656_v135 _dafny.Array
				_ = _1656_v135
				var _len0_48 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_48
				var _nw304 _dafny.Array
				_ = _nw304
				if _len0_48.Cmp(_dafny.Zero) == 0 {
					_nw304 = _dafny.NewArray(_len0_48)
				} else {
					var _init48 func(_dafny.Int) _dafny.Int = (func(_1657_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1658_i12 _dafny.Int) _dafny.Int {
							return (_1658_i12).Times(_1657_i8)
						}
					})(_1609_i8)
					_ = _init48
					var _element0_48 = _init48(_dafny.Zero)
					_ = _element0_48
					_nw304 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
					(_nw304).ArraySet1(_element0_48, 0)
					var _nativeLen0_48 = (_len0_48).Int()
					_ = _nativeLen0_48
					for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
						(_nw304).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
					}
				}
				_1656_v135 = _nw304
				var _1659_v136 _dafny.Map
				_ = _1659_v136
				_1659_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1609_i8)
				var _1660_v137 D3
				_ = _1660_v137
				_1660_v137 = Companion_D3_.Create_DC9_(Companion_Default___.Fm17(globalState))
				var _1661_v138 _dafny.Map
				_ = _1661_v138
				_1661_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1609_i8, _1527_v57)
				var _1662_v139 _dafny.Map
				_ = _1662_v139
				_1662_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1660_v137, (_1661_v138).Cardinality())
				var _1663_v141 _dafny.Array
				_ = _1663_v141
				var _nwElement0_51 _dafny.Sequence = (_1530_v60).F17()
				_ = _nwElement0_51
				var _nw305 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(25))
				_ = _nw305
				(_nw305).ArraySet1(_nwElement0_51, 0)
				(_nw305).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1664_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1665_i13 _dafny.Int) _dafny.Int {
						return _1664_i8
					}
				})(_1609_i8))), 1)
				(_nw305).ArraySet1(_dafny.SeqOf(_1609_i8), 2)
				(_nw305).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32(((_1651_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))).Int()).(_dafny.Sequence)).Cardinality())), 3)
				(_nw305).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_1666_i8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1667_i14 _dafny.Int) _dafny.Int {
						return _1666_i8
					}
				})(_1609_i8))), 4)
				(_nw305).ArraySet1((_1530_v60).F17(), 5)
				(_nw305).ArraySet1((_1530_v60).F17(), 6)
				(_nw305).ArraySet1((_1530_v60).F17(), 7)
				(_nw305).ArraySet1((_1530_v60).F17(), 8)
				(_nw305).ArraySet1((_1530_v60).F17(), 9)
				(_nw305).ArraySet1((_1530_v60).F17(), 10)
				(_nw305).ArraySet1((_1530_v60).F17(), 11)
				(_nw305).ArraySet1(_dafny.SeqOf(_1528_v58, _1528_v58, _1609_i8), 12)
				(_nw305).ArraySet1((_1530_v60).F17(), 13)
				(_nw305).ArraySet1((_1530_v60).F17(), 14)
				(_nw305).ArraySet1(_dafny.SeqOf(_1609_i8, _1528_v58, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(320))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_1668_v57 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1669_i15 _dafny.Int) _dafny.CodePoint {
						return _1668_v57
					}
				})(_1527_v57)))).Cardinality())), 15)
				(_nw305).ArraySet1((_1530_v60).F17(), 16)
				(_nw305).ArraySet1(_dafny.Companion_Sequence_.Update((_1530_v60).F17(), (Companion_Default___.SafeIndex(_1609_i8, _dafny.IntOfUint32(((_1530_v60).F17()).Cardinality()))).Uint32(), (_1662_v139).Cardinality()), 17)
				(_nw305).ArraySet1((_1530_v60).F17(), 18)
				(_nw305).ArraySet1(_dafny.SeqOf((func() _dafny.Set {
					var _coll64 = _dafny.NewBuilder()
					_ = _coll64
					for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(263), _dafny.IntOfInt64(597))); ; {
						_compr_64, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _1670_v140 _dafny.Int
						_1670_v140 = interface{}(_compr_64).(_dafny.Int)
						if ((_dafny.IntOfInt64(263)).Cmp(_1670_v140) <= 0) && ((_1670_v140).Cmp(_dafny.IntOfInt64(597)) < 0) {
							_coll64.Add((_1670_v140).Times(_1609_i8))
						}
					}
					return _coll64.ToSet()
				}()).Cardinality()), 19)
				(_nw305).ArraySet1(_dafny.SeqOf(_1609_i8), 20)
				(_nw305).ArraySet1((_1530_v60).F17(), 21)
				(_nw305).ArraySet1((_1530_v60).F17(), 22)
				(_nw305).ArraySet1((_1530_v60).F17(), 23)
				(_nw305).ArraySet1((_1530_v60).F17(), 24)
				_1663_v141 = _nw305
				var _1671_v142 _dafny.Sequence
				_ = _1671_v142
				_1671_v142 = _dafny.SeqOf(_1663_v141, _1663_v141, _1663_v141)
				var _1672_v143 D2
				_ = _1672_v143
				_1672_v143 = Companion_D2_.Create_DC8_(_1659_v136, (_1671_v142).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1538_v68, (Companion_Default___.SafeIndex(_1528_v58, _dafny.IntOfUint32((_1538_v68).Cardinality()))).Uint32(), _1527_v57)).Cardinality()), _dafny.IntOfUint32((_1671_v142).Cardinality()))).Uint32()).(_dafny.Array), _1530_v60.F16(), !(_1454_v0), _1609_i8)
				var _1673_v144 _dafny.Map
				_ = _1673_v144
				_1673_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _1530_v60.F16())
				var _1674_v145 _dafny.Map
				_ = _1674_v145
				_1674_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1672_v143).Dtor_cf16(), (_1673_v144).Merge(_1673_v144))
				var _1675_v146 _dafny.Map
				_ = _1675_v146
				_1675_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1538_v68, _1656_v135)
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))
				_ = _index292
				var _rhs320 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_1651_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))).Int()).(_dafny.Sequence), _1538_v68), (Companion_Default___.SafeIndex(_1609_i8, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1651_v133).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))).Int()).(_dafny.Sequence), _1538_v68)).Cardinality()))).Uint32(), _1527_v57)
				_ = _rhs320
				var _rhs321 _dafny.Array = (func() _dafny.Array {
					if (_1675_v146).Contains(_1538_v68) {
						return (_1675_v146).Get(_1538_v68).(_dafny.Array)
					}
					return _1656_v135
				})()
				_ = _rhs321
				var _rhs322 _dafny.Map = (((_1674_v145).Update(!(_1530_v60.F16()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _1530_v60.F16()))).Merge(_1674_v145)).Merge(_1674_v145)
				_ = _rhs322
				var _lhs254 _dafny.Array = _1651_v133
				_ = _lhs254
				var _lhs255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_1651_v133), 0))
				_ = _lhs255
				(_lhs254).ArraySet1(_rhs320, (_lhs255).Int())
				_1656_v135 = _rhs321
				_1674_v145 = _rhs322
			} else {
				var _1676___mcc_h22 D8 = _source18.Get_().(D8_DC24).Cf43
				_ = _1676___mcc_h22
				var _1677_cf43 D8 = _1676___mcc_h22
				_ = _1677_cf43
				var _1678_v147 _dafny.Array
				_ = _1678_v147
				var _nw306 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw306
				_1678_v147 = _nw306
				var _1679_v148 _dafny.Map
				_ = _1679_v148
				_1679_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1454_v0, _1530_v60.F16())
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1678_v147), 0))
				_ = _index293
				(_1678_v147).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((func() bool {
					if (_1679_v148).Contains(_1530_v60.F16()) {
						return (_1679_v148).Get(_1530_v60.F16()).(bool)
					}
					return _1530_v60.F16()
				})()), _1528_v58), (_index293).Int())
				var _1680_v149 _dafny.Map
				_ = _1680_v149
				_1680_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("phnuuwd")).Cardinality()))
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((_1678_v147), 0))
				_ = _index294
				(_1678_v147).ArraySet1((_1680_v149).Merge(_1680_v149), (_index294).Int())
				var _1681_v150 _dafny.Map
				_ = _1681_v150
				_1681_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1536_v67, _1530_v60.F16())
				_1681_v150 = (_1681_v150).Update(_1536_v67, _1530_v60.F16())
				var _1682_v151 _dafny.Array
				_ = _1682_v151
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_49
				var _nw307 _dafny.Array
				_ = _nw307
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw307 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Int = (func(_1683_v68 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1684_i16 _dafny.Int) _dafny.Int {
							return (_1684_i16).Times(_dafny.IntOfUint32((_1683_v68).Cardinality()))
						}
					})(_1538_v68)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw307 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw307).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw307).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1682_v151 = _nw307
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1682_v151), 0))
				_ = _index295
				(_1682_v151).ArraySet1(_1528_v58, (_index295).Int())
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(934), _dafny.ArrayLen((_1682_v151), 0))
				_ = _index296
				(_1682_v151).ArraySet1(_1609_i8, (_index296).Int())
				(_1530_v60).F16_set_(true)
			}
			var _1685_v152 _dafny.Set
			_ = _1685_v152
			_1685_v152 = _dafny.SetOf(_1530_v60.F16(), (func() bool {
				if _1454_v0 {
					return _1530_v60.F16()
				}
				return _1530_v60.F16()
			})())
			_1685_v152 = (_1685_v152).Union(_1685_v152)
			_1528_v58 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1454_v0) || (!(_1454_v0)) {
					return (_1530_v60).F17()
				}
				return (_1530_v60).F17()
			})()).Cardinality())
			_1528_v58 = (_1609_i8).Plus(_1528_v58)
		}
		var _1686_v153 _dafny.Map
		_ = _1686_v153
		_1686_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _1528_v58)
		var _1687_v154 _dafny.Map
		_ = _1687_v154
		_1687_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1686_v153).Cardinality(), _1454_v0)
		var _1688_v155 _dafny.Sequence
		_ = _1688_v155
		_1688_v155 = _dafny.SeqOf(_1687_v154, _1687_v154)
		r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_1688_v155, _dafny.Companion_Sequence_.Concatenate(_1688_v155, _1688_v155))
		var _1689_v156 _dafny.MultiSet
		_ = _1689_v156
		_1689_v156 = _dafny.MultiSetOf(_1686_v153, _1686_v153, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1530_v60.F16(), _1528_v58))
		var _pat_let_tv32 = _1540_v70
		_ = _pat_let_tv32
		var _pat_let_tv33 = _1540_v70
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1530_v60
		_ = _pat_let_tv34
		var _pat_let_tv35 = _1530_v60
		_ = _pat_let_tv35
		r1 = func(_source19 D4) bool {
			if _source19.Is_DC12() {
				var _1690___mcc_h23 bool = _source19.Get_().(D4_DC12).Cf21
				_ = _1690___mcc_h23
				var _1691___mcc_h24 bool = _source19.Get_().(D4_DC12).Cf22
				_ = _1691___mcc_h24
				var _1692___mcc_h25 _dafny.Int = _source19.Get_().(D4_DC12).Cf23
				_ = _1692___mcc_h25
				var _1693_cf23 _dafny.Int = _1692___mcc_h25
				_ = _1693_cf23
				var _1694_cf22 bool = _1691___mcc_h24
				_ = _1694_cf22
				var _1695_cf21 bool = _1690___mcc_h23
				_ = _1695_cf21
				return (_pat_let_tv32).IsSubsetOf(_pat_let_tv33)
			} else if _source19.Is_DC13() {
				var _1696___mcc_h26 _dafny.Int = _source19.Get_().(D4_DC13).Cf24
				_ = _1696___mcc_h26
				var _1697___mcc_h27 _dafny.Int = _source19.Get_().(D4_DC13).Cf25
				_ = _1697___mcc_h27
				var _1698___mcc_h28 _dafny.Int = _source19.Get_().(D4_DC13).Cf26
				_ = _1698___mcc_h28
				var _1699_cf26 _dafny.Int = _1698___mcc_h28
				_ = _1699_cf26
				var _1700_cf25 _dafny.Int = _1697___mcc_h27
				_ = _1700_cf25
				var _1701_cf24 _dafny.Int = _1696___mcc_h26
				_ = _1701_cf24
				return _pat_let_tv34.F16()
			} else {
				var _1702___mcc_h29 _dafny.MultiSet = _source19.Get_().(D4_DC11).Cf20
				_ = _1702___mcc_h29
				var _1703_cf20 _dafny.MultiSet = _1702___mcc_h29
				_ = _1703_cf20
				return _pat_let_tv35.F16()
			}
		}(Companion_D4_.Create_DC11_(_1689_v156))
		r2 = !(_1530_v60.F16())
		r3 = Companion_Default___.Fm3(globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C6) F21() D0 {
	{
		return _this._f21
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f16 bool
	_f17 _dafny.Sequence
	_f20 _dafny.Array
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	_this._f20 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	return &_this
}

type CompanionStruct_C7_ struct {
}

var Companion_C7_ = CompanionStruct_C7_{}

func (_this *C7) Equals(other *C7) bool {
	return _this == other
}

func (_this *C7) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C7)
	return ok && _this.Equals(other)
}

func (*C7) String() string {
	return "_module.C7"
}

func Type_C7_() _dafny.TypeDescriptor {
	return type_C7_{}
}

type type_C7_ struct {
}

func (_this type_C7_) Default() interface{} {
	return (*C7)(nil)
}

func (_this type_C7_) String() string {
	return "main.C7"
}
func (_this *C7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F16() bool {
	return _this._f16
}
func (_this *C7) F16_set_(value bool) {
	_this._f16 = value
}
func (_this *C7) F17() _dafny.Sequence {
	return _this._f17
}
func (_this *C7) Ctor__(f20 _dafny.Array, f16 bool, f17 _dafny.Sequence) {
	{
		(_this)._f20 = f20
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C7) M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _1704_v0 _dafny.Int
		_ = _1704_v0
		_1704_v0 = _dafny.IntOfInt64(33)
		_1704_v0 = Companion_Default___.Fm3(globalState)
		(globalState).F5 = !(_this.F16()) || ((p3).Cmp((_dafny.MultiSetFromSeq(_dafny.SeqOf(_1704_v0))).Cardinality()) == 0)
		var _1705_v1 _dafny.Array
		_ = _1705_v1
		var _nw308 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
		_ = _nw308
		_1705_v1 = _nw308
		var _1706_v2 D0
		_ = _1706_v2
		_1706_v2 = Companion_D0_.Create_DC2_(_1704_v0, _1704_v0, _1704_v0, _1705_v1)
		var _hi15 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(929), (_1706_v2).Dtor_cf3())
		_ = _hi15
		for _1707_i0 := (p3).Plus(_1704_v0); _1707_i0.Cmp(_hi15) < 0; _1707_i0 = _1707_i0.Plus(_dafny.One) {
			var _1708_v3 _dafny.Sequence
			_ = _1708_v3
			_1708_v3 = _dafny.SeqOf(p0)
			var _1709_v4 _dafny.Array
			_ = _1709_v4
			var _nwElement0_52 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("w")
			_ = _nwElement0_52
			var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(24))
			_ = _nw309
			(_nw309).ArraySet1(_nwElement0_52, 0)
			(_nw309).ArraySet1(p0, 1)
			(_nw309).ArraySet1(p0, 2)
			(_nw309).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(802))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}(func(_1710_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), 3)
			(_nw309).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-237))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg68 _dafny.Int) interface{} {
					return coer68(arg68)
				}
			}(func(_1711_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			})), 4)
			(_nw309).ArraySet1(p0, 5)
			(_nw309).ArraySet1(p0, 6)
			(_nw309).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("magnsasim"), 7)
			(_nw309).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("edeabhm"), 8)
			(_nw309).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xiq"), 9)
			(_nw309).ArraySet1(p0, 10)
			(_nw309).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("mul"), 11)
			(_nw309).ArraySet1(p0, 12)
			(_nw309).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 13)
			(_nw309).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1712_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), 14)
			(_nw309).ArraySet1((func() _dafny.Sequence {
				if false {
					return _dafny.UnicodeSeqOfUtf8Bytes("wbhpqphi")
				}
				return p0
			})(), 15)
			(_nw309).ArraySet1(p0, 16)
			(_nw309).ArraySet1(p0, 17)
			(_nw309).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, Companion_Default___.Fm1(globalState)), 18)
			(_nw309).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, p0), 19)
			(_nw309).ArraySet1(p0, 20)
			(_nw309).ArraySet1((_1708_v3).Select((Companion_Default___.SafeIndex(_1704_v0, _dafny.IntOfUint32((_1708_v3).Cardinality()))).Uint32()).(_dafny.Sequence), 21)
			(_nw309).ArraySet1(p0, 22)
			(_nw309).ArraySet1(p0, 23)
			_1709_v4 = _nw309
			_1709_v4 = _1709_v4
			var _1713_v5 _dafny.Map
			_ = _1713_v5
			_1713_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F16())
			_1713_v5 = (_1713_v5).Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), _this.F16())
			var _1714_v6 *C3
			_ = _1714_v6
			var _nw310 *C3 = New_C3_()
			_ = _nw310
			_nw310.Ctor__(p0, p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}((func(_1715_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1716_i4 _dafny.Int) _dafny.Int {
					return _1715_i0
				}
			})(_1707_i0))))
			_1714_v6 = _nw310
			var _1717_v7 D11
			_ = _1717_v7
			_1717_v7 = Companion_D11_.Create_DC30_(p2, (_dafny.Zero).Minus(_1704_v0))
			_1704_v0 = ((_1717_v7).Dtor_cf52()).Minus(p3)
		}
		if p1 {
			var _1718_v8 _dafny.CodePoint
			_ = _1718_v8
			_1718_v8 = _dafny.CodePoint('o')
			var _1719_v9 _dafny.Map
			_ = _1719_v9
			_1719_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1718_v8, p3)
			_1704_v0 = (_1719_v9).Cardinality()
			_1705_v1 = _1705_v1
			var _1720_v10 _dafny.Sequence
			_ = _1720_v10
			_1720_v10 = _dafny.SeqOf(_1704_v0)
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen(((_this).F20()), 0))
			_ = _index297
			((_this).F20()).ArraySet1(_dafny.IntOfInt64(252), (_index297).Int())
			var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen(((_this).F20()), 0))
			_ = _index298
			var _rhs323 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_this).F17(), (_this).F17()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}((func(_1721_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1722_i5 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F16())).Cardinality())), _dafny.SeqOf(_dafny.IntOfUint32((_1721_p0).Cardinality())))).Cardinality()
				}
			})(p0))))
			_ = _rhs323
			var _rhs324 _dafny.Int = Companion_Default___.Fm3(globalState)
			_ = _rhs324
			var _rhs325 _dafny.Int = p3
			_ = _rhs325
			var _lhs256 _dafny.Array = (_this).F20()
			_ = _lhs256
			var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen(((_this).F20()), 0))
			_ = _lhs257
			_1720_v10 = _rhs323
			_1704_v0 = _rhs324
			(_lhs256).ArraySet1(_rhs325, (_lhs257).Int())
			var _1723_v11 _dafny.Set
			_ = _1723_v11
			_1723_v11 = _dafny.SetOf(Companion_Default___.Fm0(p2, p1, globalState))
			var _1724_v12 _dafny.Set
			_ = _1724_v12
			_1724_v12 = _dafny.SetOf(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), p3, p3, Companion_Default___.Fm3(globalState))
			var _1725_v13 _dafny.Map
			_ = _1725_v13
			_1725_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1723_v11, _1724_v12)
			_1725_v13 = (_1725_v13).Update(_1723_v11, _1724_v12)
			var _1726_v14 _dafny.Array
			_ = _1726_v14
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_50
			var _nw311 _dafny.Array
			_ = _nw311
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw311 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Int = func(_1727_i6 _dafny.Int) _dafny.Int {
					return (_1727_i6).Times(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int))
				}
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw311 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw311).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw311).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_1726_v14 = _nw311
			_1726_v14 = _1726_v14
		} else {
			var _1728_v15 _dafny.Array
			_ = _1728_v15
			var _nw312 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw312
			_1728_v15 = _nw312
			var _1729_v16 _dafny.CodePoint
			_ = _1729_v16
			_1729_v16 = _dafny.CodePoint('q')
			var _1730_v17 _dafny.Map
			_ = _1730_v17
			_1730_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1729_v16, p3)
			var _1731_v18 _dafny.Set
			_ = _1731_v18
			_1731_v18 = _dafny.SetOf(p3)
			var _nwElement0_53 _dafny.Int = _1704_v0
			_ = _nwElement0_53
			var _nw313 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(12))
			_ = _nw313
			(_nw313).ArraySet1(_nwElement0_53, 0)
			(_nw313).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if p2 {
					return _dafny.UnicodeSeqOfUtf8Bytes("eio")
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}(func(_1732_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))
			})(), (Companion_Default___.SafeIndex(_1704_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p2 {
					return _dafny.UnicodeSeqOfUtf8Bytes("eio")
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1733_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))
			})()).Cardinality()))).Uint32(), _1729_v16)).Cardinality()), 1)
			(_nw313).ArraySet1((func() _dafny.Int {
				if (_1730_v17).Contains(_1729_v16) {
					return (_1730_v17).Get(_1729_v16).(_dafny.Int)
				}
				return _1704_v0
			})(), 2)
			(_nw313).ArraySet1(_1704_v0, 3)
			(_nw313).ArraySet1(Companion_Default___.SafeModuloInt(p3, _dafny.IntOfUint32((p0).Cardinality())), 4)
			(_nw313).ArraySet1((func() _dafny.Int {
				if p2 {
					return _1704_v0
				}
				return p3
			})(), 5)
			(_nw313).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), 6)
			(_nw313).ArraySet1(Companion_Default___.SafeDivisionInt(_1704_v0, _1704_v0), 7)
			(_nw313).ArraySet1(p3, 8)
			(_nw313).ArraySet1((_1731_v18).Cardinality(), 9)
			(_nw313).ArraySet1(p3, 10)
			(_nw313).ArraySet1((func() _dafny.Int {
				if p2 {
					return _dafny.IntOfInt64(956)
				}
				return (_1731_v18).Cardinality()
			})(), 11)
			_1728_v15 = _nw313
			var _1734_v19 _dafny.Map
			_ = _1734_v19
			_1734_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("hwadxtdi"), p1)
			var _1735_v20 _dafny.Sequence
			_ = _1735_v20
			_1735_v20 = _dafny.SeqOf(_this.F16())
			var _1736_v21 D4
			_ = _1736_v21
			_1736_v21 = Companion_D4_.Create_DC12_(p2, Companion_Default___.Fm0(p1, p2, globalState), _dafny.IntOfUint32((_1735_v20).Cardinality()))
			_1734_v19 = (_1734_v19).Update(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_1736_v21).Dtor_cf21())
			var _1737_v22 _dafny.Sequence
			_ = _1737_v22
			_1737_v22 = _dafny.UnicodeSeqOfUtf8Bytes("gmrqcudhj")
			_1737_v22 = _1737_v22
			var _1738_v23 *C1
			_ = _1738_v23
			var _nw314 *C1 = New_C1_()
			_ = _nw314
			_nw314.Ctor__(p3, Companion_Default___.SafeDivisionInt(p3, (_1731_v18).Cardinality()), p2, (_this).F17(), _1704_v0)
			_1738_v23 = _nw314
			var _1739_v24 _dafny.Sequence
			_ = _1739_v24
			_1739_v24 = _dafny.SeqOf(_1704_v0)
			var _1740_v25 _dafny.Map
			_ = _1740_v25
			_1740_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1738_v23).F29(), !(_this.F16()))
			var _1741_v26 *C0
			_ = _1741_v26
			var _nw315 *C0 = New_C0_()
			_ = _nw315
			_nw315.Ctor__((_1738_v23).F29(), _1705_v1)
			_1741_v26 = _nw315
			var _1742_v27 _dafny.MultiSet
			_ = _1742_v27
			_1742_v27 = _dafny.MultiSetOf(p1)
			var _rhs326 _dafny.Sequence = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm17(globalState), (Companion_Default___.SafeIndex((Companion_D9_.Create_DC26_((_1740_v25).Cardinality(), _1741_v26)).Dtor_cf45(), _dafny.IntOfUint32((Companion_Default___.Fm17(globalState)).Cardinality()))).Uint32(), (_1704_v0).Minus((_1738_v23).F30()))
			_ = _rhs326
			var _rhs327 _dafny.Int = (_dafny.Zero).Minus(((_1741_v26).F26()).Minus(_1704_v0))
			_ = _rhs327
			var _rhs328 bool = Companion_Default___.Fm0((func() bool {
				if p2 {
					return !(_this.F16())
				}
				return p1
			})(), (_1731_v18).IsProperSubsetOf(_1731_v18), globalState)
			_ = _rhs328
			var _rhs329 bool = (_1742_v27).IsSubsetOf(_1742_v27)
			_ = _rhs329
			var _rhs330 _dafny.Map = _1734_v19
			_ = _rhs330
			var _lhs258 *GlobalState = globalState
			_ = _lhs258
			var _lhs259 *GlobalState = globalState
			_ = _lhs259
			_1739_v24 = _rhs326
			_1704_v0 = _rhs327
			_lhs258.F5 = _rhs328
			_lhs259.F2 = _rhs329
			_1734_v19 = _rhs330
		}
		var _1743_v28 D0
		_ = _1743_v28
		_1743_v28 = Companion_D0_.Create_DC0_(_dafny.UnicodeSeqOfUtf8Bytes("egdqwsb"))
		if _dafny.Companion_Sequence_.IsProperPrefixOf((_1743_v28).Dtor_cf0(), p0) {
			var _1744_v29 _dafny.Map
			_ = _1744_v29
			_1744_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v1, _this.F16())
			var _1745_v30 _dafny.Map
			_ = _1745_v30
			_1745_v30 = _1744_v29
			var _1746_v31 _dafny.Array
			_ = _1746_v31
			var _nwElement0_54 _dafny.Map = _1744_v29
			_ = _nwElement0_54
			var _nw316 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(21))
			_ = _nw316
			(_nw316).ArraySet1(_nwElement0_54, 0)
			(_nw316).ArraySet1(_1744_v29, 1)
			(_nw316).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v1, p1), 2)
			(_nw316).ArraySet1(_1744_v29, 3)
			(_nw316).ArraySet1(_1744_v29, 4)
			(_nw316).ArraySet1(_1744_v29, 5)
			(_nw316).ArraySet1(_1744_v29, 6)
			(_nw316).ArraySet1((_1744_v29).Merge(_1744_v29), 7)
			(_nw316).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v1, p1), 8)
			(_nw316).ArraySet1((_1744_v29).Merge(_1744_v29), 9)
			(_nw316).ArraySet1(_1744_v29, 10)
			(_nw316).ArraySet1((_1745_v30), 11)
			(_nw316).ArraySet1(_1744_v29, 12)
			(_nw316).ArraySet1(_1744_v29, 13)
			(_nw316).ArraySet1(_1744_v29, 14)
			(_nw316).ArraySet1(_1744_v29, 15)
			(_nw316).ArraySet1((_1744_v29).Update(_1705_v1, p2), 16)
			(_nw316).ArraySet1(_1744_v29, 17)
			(_nw316).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v1, !(_this.F16())), 18)
			(_nw316).ArraySet1((_1744_v29).Update(_1705_v1, p2), 19)
			(_nw316).ArraySet1(_1744_v29, 20)
			_1746_v31 = _nw316
			var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1746_v31), 0))
			_ = _index299
			(_1746_v31).ArraySet1((_1744_v29).Merge(_1744_v29), (_index299).Int())
			var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1746_v31), 0))
			_ = _index300
			(_1746_v31).ArraySet1(_1744_v29, (_index300).Int())
			var _1747_v32 _dafny.Set
			_ = _1747_v32
			_1747_v32 = _dafny.SetOf((_this).F20(), (_this).F20())
			var _1748_v33 *C5
			_ = _1748_v33
			var _nw317 *C5 = New_C5_()
			_ = _nw317
			_nw317.Ctor__(_this.F16(), _dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex((_1747_v32).Cardinality(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), p3))
			_1748_v33 = _nw317
			var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
			_ = _index301
			(_1705_v1).ArraySet1(p2, (_index301).Int())
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
			_ = _index302
			(_1705_v1).ArraySet1(_this.F16(), (_index302).Int())
			var _1749_v34 _dafny.Map
			_ = _1749_v34
			_1749_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1704_v0)
			var _1750_v35 _dafny.Array
			_ = _1750_v35
			var _len0_51 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_51
			var _nw318 _dafny.Array
			_ = _nw318
			if _len0_51.Cmp(_dafny.Zero) == 0 {
				_nw318 = _dafny.NewArray(_len0_51)
			} else {
				var _init51 func(_dafny.Int) _dafny.Sequence = func(_1751_i8 _dafny.Int) _dafny.Sequence {
					return (_this).F17()
				}
				_ = _init51
				var _element0_51 = _init51(_dafny.Zero)
				_ = _element0_51
				_nw318 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
				(_nw318).ArraySet1(_element0_51, 0)
				var _nativeLen0_51 = (_len0_51).Int()
				_ = _nativeLen0_51
				for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
					(_nw318).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
				}
			}
			_1750_v35 = _nw318
			var _1752_v36 D2
			_ = _1752_v36
			_1752_v36 = Companion_D2_.Create_DC8_(_1749_v34, _1750_v35, !(_this.F16()), p1, p3)
			var _source20 D6 = Companion_D6_.Create_DC17_(_1752_v36)
			_ = _source20
			if _source20.Is_DC16() {
				var _1753___mcc_h0 bool = _source20.Get_().(D6_DC16).Cf29
				_ = _1753___mcc_h0
				var _1754___mcc_h1 bool = _source20.Get_().(D6_DC16).Cf30
				_ = _1754___mcc_h1
				var _1755_cf30 bool = _1754___mcc_h1
				_ = _1755_cf30
				var _1756_cf29 bool = _1753___mcc_h0
				_ = _1756_cf29
				var _1757_v37 _dafny.Sequence
				_ = _1757_v37
				_1757_v37 = _dafny.UnicodeSeqOfUtf8Bytes("lutqh")
				var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index303
				var _rhs331 bool = _1755_cf30
				_ = _rhs331
				var _rhs332 _dafny.Int = (_1704_v0).Plus(_1704_v0)
				_ = _rhs332
				var _rhs333 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, _1757_v37)).Cardinality())).Cmp(_1704_v0) <= 0
				_ = _rhs333
				var _rhs334 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1748_v33).Fm7(_dafny.SeqOf(_1756_cf29, _1756_cf29), globalState), _dafny.UnicodeSeqOfUtf8Bytes("soa"))
				_ = _rhs334
				var _lhs260 *GlobalState = globalState
				_ = _lhs260
				var _lhs261 _dafny.Array = _1705_v1
				_ = _lhs261
				var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _lhs262
				_lhs260.F2 = _rhs331
				_1704_v0 = _rhs332
				(_lhs261).ArraySet1(_rhs333, (_lhs262).Int())
				_1757_v37 = _rhs334
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index304
				(_1705_v1).ArraySet1(p1, (_index304).Int())
				var _1758_v38 _dafny.Map
				_ = _1758_v38
				_1758_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(globalState), _dafny.SeqOf((_this).F20()))
				var _1759_v39 _dafny.Sequence
				_ = _1759_v39
				_1759_v39 = _dafny.SeqOf((_this).F20())
				_1758_v38 = (_1758_v38).Update(p3, _dafny.Companion_Sequence_.Concatenate(_1759_v39, _1759_v39))
				_1755_cf30 = p1
			} else if _source20.Is_DC17() {
				var _1760___mcc_h2 D2 = _source20.Get_().(D6_DC17).Cf31
				_ = _1760___mcc_h2
				var _1761_cf31 D2 = _1760___mcc_h2
				_ = _1761_cf31
				_1704_v0 = ((_dafny.IntOfInt64(677)).Minus((_dafny.Zero).Minus(p3))).Minus((_1704_v0).Plus(p3))
				var _1762_v40 _dafny.Map
				_ = _1762_v40
				_1762_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_this).F20())
				var _1763_v41 _dafny.Sequence
				_ = _1763_v41
				_1763_v41 = _dafny.SeqOf((_this).F20(), (_this).F20(), (_this).F20(), (_this).F20(), (_this).F20())
				var _1764_v42 _dafny.MultiSet
				_ = _1764_v42
				_1764_v42 = _dafny.MultiSetOf((_this).F20(), (_this).F20(), (_this).F20(), (func() _dafny.Array {
					if (_1762_v40).Contains((_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool)) {
						return (_1762_v40).Get((_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool)).(_dafny.Array)
					}
					return (_this).F20()
				})(), (_1763_v41).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1763_v41).Cardinality()))).Uint32()).(_dafny.Array))
				_1764_v42 = ((_1764_v42).Union(_1764_v42)).Union(_1764_v42)
				var _1765_v43 _dafny.Map
				_ = _1765_v43
				_1765_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				_1762_v40 = (_1762_v40).Update(!((func() bool {
					if (_1765_v43).Contains(p0) {
						return (_1765_v43).Get(p0).(bool)
					}
					return p2
				})()), (_this).F20())
				var _1766_v44 _dafny.Array
				_ = _1766_v44
				var _nw319 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw319
				_1766_v44 = _nw319
				var _1767_v45 *C0
				_ = _1767_v45
				var _nw320 *C0 = New_C0_()
				_ = _nw320
				_nw320.Ctor__(p3, _1766_v44)
				_1767_v45 = _nw320
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index305
				var _rhs335 bool = (true) && (p2)
				_ = _rhs335
				var _rhs336 *C0 = _1767_v45
				_ = _rhs336
				var _lhs263 _dafny.Array = _1705_v1
				_ = _lhs263
				var _lhs264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _lhs264
				(_lhs263).ArraySet1(_rhs335, (_lhs264).Int())
				_1767_v45 = _rhs336
			} else if _source20.Is_DC18() {
				var _1768___mcc_h3 _dafny.Array = _source20.Get_().(D6_DC18).Cf32
				_ = _1768___mcc_h3
				var _1769_cf32 _dafny.Array = _1768___mcc_h3
				_ = _1769_cf32
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1769_cf32), 0))
				_ = _index306
				(_1769_cf32).ArraySet1(p3, (_index306).Int())
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(629), _dafny.ArrayLen((_1769_cf32), 0))
				_ = _index307
				(_1769_cf32).ArraySet1(p3, (_index307).Int())
				(globalState).F2 = p1
				var _1770_v46 *C4
				_ = _1770_v46
				var _nw321 *C4 = New_C4_()
				_ = _nw321
				_nw321.Ctor__(_dafny.IntOfInt64(235), _this.F16())
				_1770_v46 = _nw321
				(globalState).F5 = (_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool)
			} else {
				var _1771___mcc_h4 _dafny.Array = _source20.Get_().(D6_DC15).Cf28
				_ = _1771___mcc_h4
				var _1772_cf28 _dafny.Array = _1771___mcc_h4
				_ = _1772_cf28
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index308
				(_1705_v1).ArraySet1((_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool), (_index308).Int())
				var _1773_v47 D0
				_ = _1773_v47
				_1773_v47 = Companion_D0_.Create_DC1_(p0)
				(globalState).F5 = _dafny.Companion_Sequence_.IsPrefixOf(p0, (func() _dafny.Sequence {
					if p2 {
						return (_1773_v47).Dtor_cf1()
					}
					return p0
				})())
				var _1774_v48 _dafny.MultiSet
				_ = _1774_v48
				_1774_v48 = _dafny.MultiSetOf(_1704_v0, p3)
				var _1775_v49 _dafny.MultiSet
				_ = _1775_v49
				_1775_v49 = _dafny.MultiSetOf((_1774_v48).Intersection(_1774_v48), Companion_Default___.Fm19(_1704_v0, globalState), _dafny.MultiSetFromSeq(Companion_Default___.Fm17(globalState)))
				_1775_v49 = Companion_Default___.Fm39((_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(498))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}((func(_1776_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1777_i9 _dafny.Int) _dafny.Int {
						return _1776_p3
					}
				})(p3))), (p0).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState)
				(_this).M2(globalState)
			}
			var _1778_v50 _dafny.Map
			_ = _1778_v50
			_1778_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _rhs337 bool = !(!(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("iuu"), p0), p0))))
			_ = _rhs337
			var _rhs338 _dafny.Int = (func() _dafny.Int {
				if (_1778_v50).Contains(_1704_v0) {
					return (_1778_v50).Get(_1704_v0).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(489), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool) {
						return _1704_v0
					}
					return _1704_v0
				})()
			})()
			_ = _rhs338
			var _rhs339 _dafny.Int = p3
			_ = _rhs339
			var _lhs265 *GlobalState = globalState
			_ = _lhs265
			_lhs265.F5 = _rhs337
			_1704_v0 = _rhs338
			_1704_v0 = _rhs339
		} else {
			var _1779_v51 _dafny.CodePoint
			_ = _1779_v51
			_1779_v51 = _dafny.CodePoint('q')
			if ((_1704_v0).Times((Companion_D12_.Create_DC34_(_1705_v1, p3, (_this).F20(), _1779_v51)).Dtor_cf60())).Cmp(p3) >= 0 {
				var _1780_v52 D11
				_ = _1780_v52
				_1780_v52 = Companion_D11_.Create_DC31_(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(721))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1781_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1782_i10 _dafny.Int) _dafny.CodePoint {
						return _1781_v51
					}
				})(_1779_v51))), p0)).Cardinality()), _1704_v0, p2)
				_1780_v52 = _1780_v52
				var _1783_v53 _dafny.MultiSet
				_ = _1783_v53
				_1783_v53 = _dafny.MultiSetOf(_1704_v0)
				Companion_Default___.M0((func() _dafny.Int {
					if (_1783_v53).Contains(p3) {
						return (_1783_v53).Multiplicity(p3)
					}
					return (_dafny.Zero).Minus(Companion_Default___.Fm3(globalState))
				})(), _dafny.UnicodeSeqOfUtf8Bytes("re"), ((_this).F17()).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm3(globalState), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), _1704_v0, globalState)
				(globalState).F2 = _this.F16()
				var _1784_v54 _dafny.Sequence
				_ = _1784_v54
				_1784_v54 = _dafny.UnicodeSeqOfUtf8Bytes("ykvixfv")
				_1784_v54 = _dafny.Companion_Sequence_.Concatenate(_1784_v54, _dafny.UnicodeSeqOfUtf8Bytes("ir"))
				var _1785_v55 _dafny.Set
				_ = _1785_v55
				_1785_v55 = _dafny.SetOf(p2)
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index309
				((_this).F20()).ArraySet1(((_1785_v55).Intersection(_1785_v55)).Cardinality(), (_index309).Int())
				var _1786_v56 _dafny.Sequence
				_ = _1786_v56
				_1786_v56 = _dafny.SeqOf((_this).F17(), Companion_Default___.Fm17(globalState), (_this).F17(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-271))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1787_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1788_i11 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1787_v0)
					}
				})(_1704_v0))), (_this).F17())
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index310
				var _rhs340 _dafny.Int = _dafny.IntOfUint32((p0).Cardinality())
				_ = _rhs340
				var _rhs341 _dafny.Sequence = _1786_v56
				_ = _rhs341
				var _rhs342 _dafny.Sequence = _1784_v54
				_ = _rhs342
				var _lhs266 _dafny.Array = (_this).F20()
				_ = _lhs266
				var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _lhs267
				(_lhs266).ArraySet1(_rhs340, (_lhs267).Int())
				_1786_v56 = _rhs341
				_1784_v54 = _rhs342
			} else {
				(globalState).F5 = !(p1)
				(globalState).F5 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(914))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1789_v51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1790_i12 _dafny.Int) _dafny.CodePoint {
						return _1789_v51
					}
				})(_1779_v51)))), p0)
				_1704_v0 = p3
				var _1791_v57 _dafny.Array
				_ = _1791_v57
				var _nwElement0_55 _dafny.Array = _1705_v1
				_ = _nwElement0_55
				var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(3))
				_ = _nw322
				(_nw322).ArraySet1(_nwElement0_55, 0)
				(_nw322).ArraySet1((func() _dafny.Array {
					if p1 {
						return _1705_v1
					}
					return _1705_v1
				})(), 1)
				(_nw322).ArraySet1(_1705_v1, 2)
				_1791_v57 = _nw322
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_1791_v57), 0))
				_ = _index311
				(_1791_v57).ArraySet1(_1705_v1, (_index311).Int())
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_1791_v57), 0))
				_ = _index312
				(_1791_v57).ArraySet1(_1705_v1, (_index312).Int())
				_1704_v0 = _1704_v0
			}
			if (_1704_v0).Cmp(_1704_v0) > 0 {
				var _1792_v58 D2
				_ = _1792_v58
				_1792_v58 = Companion_D2_.Create_DC7_(_dafny.SetOf(_this.F16(), _this.F16(), p1))
				var _1793_v59 _dafny.Sequence
				_ = _1793_v59
				_1793_v59 = _dafny.SeqOf(_1792_v58)
				var _1794_v60 _dafny.Sequence
				_ = _1794_v60
				_1794_v60 = _dafny.SeqOf(_1793_v59, _1793_v59)
				var _1795_v61 _dafny.Map
				_ = _1795_v61
				_1795_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq((_1794_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1794_v60).Cardinality()))).Uint32()).(_dafny.Sequence)), p3)
				var _1796_v62 _dafny.Set
				_ = _1796_v62
				_1796_v62 = _dafny.SetOf(p2, _this.F16())
				var _1797_v63 _dafny.MultiSet
				_ = _1797_v63
				_1797_v63 = _dafny.MultiSetOf(Companion_D2_.Create_DC7_(_1796_v62))
				_1795_v61 = (_1795_v61).Update(_1797_v63, _1704_v0)
				var _1798_v64 _dafny.Sequence
				_ = _1798_v64
				_1798_v64 = _dafny.SeqOf((_this).F17())
				var _1799_v65 T0
				_ = _1799_v65
				var _nw323 *C1 = New_C1_()
				_ = _nw323
				_nw323.Ctor__((_dafny.Zero).Minus(p3), _1704_v0, p1, (_1798_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((_1798_v64).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.IntOfInt64(154))
				_1799_v65 = _nw323
				_1799_v65 = _1799_v65
				(_this).M2(globalState)
				(globalState).F2 = _1799_v65.F16()
				var _1800_v66 _dafny.Sequence
				_ = _1800_v66
				_1800_v66 = _dafny.SeqOf((_1704_v0).Plus(p3))
				var _1801_v68 _dafny.MultiSet
				_ = _1801_v68
				_1801_v68 = _dafny.MultiSetOf(_1779_v51)
				var _rhs343 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1800_v66, _1800_v66), (_1799_v65).F17())
				_ = _rhs343
				var _rhs344 _dafny.Int = ((func() _dafny.Map {
					var _coll65 = _dafny.NewMapBuilder()
					_ = _coll65
					for _iter74 := _dafny.Iterate((_1801_v68).Elements()); ; {
						_compr_65, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _1802_v67 _dafny.CodePoint
						_1802_v67 = interface{}(_compr_65).(_dafny.CodePoint)
						if (_1801_v68).Contains(_1802_v67) {
							_coll65.Add(_1802_v67, _dafny.IntOfUint32((_1800_v66).Cardinality()))
						}
					}
					return _coll65.ToMap()
				}()).Cardinality()).Plus(Companion_Default___.SafeModuloInt(_1704_v0, _dafny.IntOfUint32((p0).Cardinality())))
				_ = _rhs344
				var _rhs345 bool = ((false) || (true)) == (_this.F16())
				_ = _rhs345
				var _rhs346 _dafny.Int = p3
				_ = _rhs346
				var _rhs347 bool = true
				_ = _rhs347
				var _lhs268 *GlobalState = globalState
				_ = _lhs268
				var _lhs269 *GlobalState = globalState
				_ = _lhs269
				_1800_v66 = _rhs343
				_1704_v0 = _rhs344
				_lhs268.F5 = _rhs345
				_1704_v0 = _rhs346
				_lhs269.F5 = _rhs347
			} else {
				var _1803_v69 _dafny.Map
				_ = _1803_v69
				_1803_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _1804_v70 D1
				_ = _1804_v70
				_1804_v70 = Companion_D1_.Create_DC4_(!(_this.F16()), _1704_v0, _1803_v69)
				var _1805_v71 _dafny.Sequence
				_ = _1805_v71
				_1805_v71 = _dafny.SeqOf((_1804_v70).Dtor_cf9(), _1803_v69, _1803_v69)
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index313
				(_1705_v1).ArraySet1(!_dafny.Companion_Sequence_.Equal(_1805_v71, _1805_v71), (_index313).Int())
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index314
				(_1705_v1).ArraySet1(p2, (_index314).Int())
				var _1806_v72 _dafny.Map
				_ = _1806_v72
				_1806_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1704_v0, p3)
				var _1807_v73 *C3
				_ = _1807_v73
				var _nw324 *C3 = New_C3_()
				_ = _nw324
				_nw324.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("t"), p1, _dafny.SeqOf(_dafny.IntOfUint32(((_this).F17()).Cardinality())))
				_1807_v73 = _nw324
				var _1808_v74 D12
				_ = _1808_v74
				_1808_v74 = Companion_D12_.Create_DC35_(p1, (p3).Cmp((_dafny.Zero).Minus(_1704_v0)) < 0, _1806_v72, (_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool), _1807_v73)
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index315
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _index316
				var _rhs348 bool = p1
				_ = _rhs348
				var _rhs349 bool = (_1705_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))).Int()).(bool)
				_ = _rhs349
				var _rhs350 bool = _this.F16()
				_ = _rhs350
				var _rhs351 D12 = _1808_v74
				_ = _rhs351
				var _lhs270 *GlobalState = globalState
				_ = _lhs270
				var _lhs271 _dafny.Array = _1705_v1
				_ = _lhs271
				var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _lhs272
				var _lhs273 _dafny.Array = _1705_v1
				_ = _lhs273
				var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1705_v1), 0))
				_ = _lhs274
				_lhs270.F5 = _rhs348
				(_lhs271).ArraySet1(_rhs349, (_lhs272).Int())
				(_lhs273).ArraySet1(_rhs350, (_lhs274).Int())
				_1808_v74 = _rhs351
				var _1809_v75 _dafny.Set
				_ = _1809_v75
				_1809_v75 = _dafny.SetOf((_dafny.Zero).Minus(_1704_v0), _dafny.IntOfInt64(406))
				var _rhs352 bool = p2
				_ = _rhs352
				var _rhs353 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((p3).Times((_1809_v75).Cardinality()), (_1704_v0).Times(p3)))
				_ = _rhs353
				var _lhs275 *C7 = _this
				_ = _lhs275
				_lhs275.F16_set_(_rhs352)
				_1704_v0 = _rhs353
				var _1810_v76 _dafny.Int
				_ = _1810_v76
				var _1811_v77 _dafny.Int
				_ = _1811_v77
				var _1812_v78 _dafny.Map
				_ = _1812_v78
				var _out15 _dafny.Int
				_ = _out15
				var _out16 _dafny.Int
				_ = _out16
				var _out17 _dafny.Map
				_ = _out17
				_out15, _out16, _out17 = (_1807_v73).M8(globalState)
				_1810_v76 = _out15
				_1811_v77 = _out16
				_1812_v78 = _out17
				var _1813_v79 *C2
				_ = _1813_v79
				var _nw325 *C2 = New_C2_()
				_ = _nw325
				_nw325.Ctor__(_1704_v0)
				_1813_v79 = _nw325
			}
			var _1814_v80 _dafny.Map
			_ = _1814_v80
			_1814_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v1, Companion_Default___.Fm0(p2, _this.F16(), globalState))
			var _1815_v81 _dafny.Sequence
			_ = _1815_v81
			_1815_v81 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1814_v80, false))
			var _1816_v82 *C1
			_ = _1816_v82
			var _nw326 *C1 = New_C1_()
			_ = _nw326
			_nw326.Ctor__(_dafny.IntOfUint32((_1815_v81).Cardinality()), _1704_v0, p2, (_this).F17(), _1704_v0)
			_1816_v82 = _nw326
			var _1817_v83 _dafny.Sequence
			_ = _1817_v83
			_1817_v83 = _dafny.SeqOf(_1816_v82, _1816_v82)
			var _1818_v84 _dafny.Sequence
			_ = _1818_v84
			_1818_v84 = _dafny.SeqOf((_1817_v83).Select((Companion_Default___.SafeIndex((_1816_v82).F30(), _dafny.IntOfUint32((_1817_v83).Cardinality()))).Uint32()).(*C1), _1816_v82, _1816_v82)
			var _1819_v85 _dafny.MultiSet
			_ = _1819_v85
			_1819_v85 = _dafny.MultiSetOf((_1816_v82).F30(), _dafny.IntOfInt64(128))
			var _1820_v86 _dafny.Array
			_ = _1820_v86
			var _nwElement0_56 *C1 = _1816_v82
			_ = _nwElement0_56
			var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(14))
			_ = _nw327
			(_nw327).ArraySet1(_nwElement0_56, 0)
			(_nw327).ArraySet1(_1816_v82, 1)
			(_nw327).ArraySet1(_1816_v82, 2)
			(_nw327).ArraySet1(_1816_v82, 3)
			(_nw327).ArraySet1((_1818_v84).Select((Companion_Default___.SafeIndex((_1819_v85).Cardinality(), _dafny.IntOfUint32((_1818_v84).Cardinality()))).Uint32()).(*C1), 4)
			(_nw327).ArraySet1(_1816_v82, 5)
			(_nw327).ArraySet1(_1816_v82, 6)
			(_nw327).ArraySet1(_1816_v82, 7)
			(_nw327).ArraySet1(_1816_v82, 8)
			(_nw327).ArraySet1(_1816_v82, 9)
			(_nw327).ArraySet1(_1816_v82, 10)
			(_nw327).ArraySet1(_1816_v82, 11)
			(_nw327).ArraySet1(_1816_v82, 12)
			(_nw327).ArraySet1((_1818_v84).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.IntOfUint32((_1818_v84).Cardinality()))).Uint32()).(*C1), 13)
			_1820_v86 = _nw327
			var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1820_v86), 0))
			_ = _index317
			(_1820_v86).ArraySet1(_1816_v82, (_index317).Int())
			var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1820_v86), 0))
			_ = _index318
			(_1820_v86).ArraySet1(_1816_v82, (_index318).Int())
			var _1821_v87 T0
			_ = _1821_v87
			var _nw328 *C1 = New_C1_()
			_ = _nw328
			_nw328.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F17()).Cardinality()), _dafny.IntOfInt64(-478)), (_dafny.Zero).Minus(((_this).F17()).Select((Companion_Default___.SafeIndex(_1704_v0, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int)), false, _dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jstytmew")).Cardinality()), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), (_1816_v82).F30()), (_1816_v82).F29())
			_1821_v87 = _nw328
			_1821_v87 = _1821_v87
			var _1822_v88 _dafny.Array
			_ = _1822_v88
			var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(26))
			_ = _nw329
			_1822_v88 = _nw329
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_1822_v88), 0))
			_ = _index319
			(_1822_v88).ArraySet1CodePoint(_1779_v51, (_index319).Int())
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_1822_v88), 0))
			_ = _index320
			(_1822_v88).ArraySet1CodePoint(_1779_v51, (_index320).Int())
		}
		var _hi16 _dafny.Int = _dafny.IntOfUint32(((_this).F17()).Cardinality())
		_ = _hi16
		for _1823_i13 := _1704_v0; _1823_i13.Cmp(_hi16) < 0; _1823_i13 = _1823_i13.Plus(_dafny.One) {
			var _1824_v89 _dafny.Array
			_ = _1824_v89
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_52
			var _nw330 _dafny.Array
			_ = _nw330
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw330 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Map = (func(_1825_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_1826_i14 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1825_v0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-813)))
					}
				})(_1704_v0)
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw330 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw330).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw330).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1824_v89 = _nw330
			var _1827_v90 _dafny.Set
			_ = _1827_v90
			_1827_v90 = _dafny.SetOf(_dafny.IntOfInt64(-759))
			var _1828_v91 _dafny.Set
			_ = _1828_v91
			_1828_v91 = _dafny.SetOf(p3, (_1827_v90).Cardinality(), p3)
			var _1829_v92 _dafny.MultiSet
			_ = _1829_v92
			_1829_v92 = _dafny.MultiSetOf(_1704_v0)
			var _1830_v93 _dafny.Map
			_ = _1830_v93
			_1830_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1704_v0, _1829_v92)
			var _rhs354 bool = !(_1827_v90).Contains(_dafny.IntOfUint32((p0).Cardinality()))
			_ = _rhs354
			var _rhs355 _dafny.Int = Companion_Default___.SafeDivisionInt((_1830_v93).Cardinality(), _1823_i13)
			_ = _rhs355
			var _rhs356 _dafny.Array = _1824_v89
			_ = _rhs356
			var _rhs357 bool = p1
			_ = _rhs357
			var _lhs276 *C7 = _this
			_ = _lhs276
			var _lhs277 *GlobalState = globalState
			_ = _lhs277
			_lhs276.F16_set_(_rhs354)
			_1704_v0 = _rhs355
			_1824_v89 = _rhs356
			_lhs277.F2 = _rhs357
			(globalState).F5 = _this.F16()
			var _1831_v94 _dafny.Map
			_ = _1831_v94
			_1831_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), p1)
			var _1832_v95 _dafny.Map
			_ = _1832_v95
			_1832_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), (_1831_v94).Cardinality())
			_1704_v0 = (func() _dafny.Int {
				if (_1832_v95).Contains(Companion_Default___.Fm3(globalState)) {
					return (_1832_v95).Get(Companion_Default___.Fm3(globalState)).(_dafny.Int)
				}
				return p3
			})()
			if false {
				var _1833_v96 D1
				_ = _1833_v96
				_1833_v96 = Companion_D1_.Create_DC3_(_dafny.MultiSetOf(p1))
				var _1834_v97 _dafny.Map
				_ = _1834_v97
				_1834_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1704_v0).Minus(_dafny.IntOfUint32((p0).Cardinality())), _1833_v96)
				_1834_v97 = (_1834_v97).Update(_1823_i13, _1833_v96)
				var _1835_v98 _dafny.MultiSet
				_ = _1835_v98
				_1835_v98 = _dafny.MultiSetOf(_1705_v1, _1705_v1, _1705_v1, _1705_v1, _1705_v1)
				var _rhs358 bool = !((func() bool {
					if (_1831_v94).Contains((func() bool {
						if p2 {
							return _this.F16()
						}
						return p2
					})()) {
						return (_1831_v94).Get((func() bool {
							if p2 {
								return _this.F16()
							}
							return p2
						})()).(bool)
					}
					return _this.F16()
				})())
				_ = _rhs358
				var _rhs359 _dafny.MultiSet = _1835_v98
				_ = _rhs359
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				_lhs278.F5 = _rhs358
				_1835_v98 = _rhs359
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index321
				((_this).F20()).ArraySet1(Companion_Default___.SafeModuloInt(_1704_v0, _1704_v0), (_index321).Int())
				var _1836_v99 _dafny.CodePoint
				_ = _1836_v99
				_1836_v99 = _dafny.CodePoint('g')
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index322
				((_this).F20()).ArraySet1((((_this).F17()).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_1836_v99, _1836_v99, _1836_v99)).Cardinality(), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int)).Minus(p3), (_index322).Int())
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index323
				((_this).F20()).ArraySet1(_1823_i13, (_index323).Int())
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen(((_this).F20()), 0))
				_ = _index324
				((_this).F20()).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), (_index324).Int())
			} else {
				(globalState).F5 = p2
				var _1837_v100 _dafny.Map
				_ = _1837_v100
				_1837_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1704_v0)
				_1837_v100 = _1837_v100
				var _pat_let_tv36 = _1704_v0
				_ = _pat_let_tv36
				var _pat_let_tv37 = _1705_v1
				_ = _pat_let_tv37
				var _pat_let_tv38 = _1705_v1
				_ = _pat_let_tv38
				var _1838_v101 *C6
				_ = _1838_v101
				var _nw331 *C6 = New_C6_()
				_ = _nw331
				_nw331.Ctor__(func(_pat_let44_0 D0) D0 {
					return func(_1843_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let49_0 _dafny.Array) D0 {
							return func(_1844_dt__update_hcf5_h1 _dafny.Array) D0 {
								return func(_pat_let50_0 _dafny.Int) D0 {
									return func(_1845_dt__update_hcf3_h1 _dafny.Int) D0 {
										return Companion_D0_.Create_DC2_((_1843_dt__update__tmp_h1).Dtor_cf2(), _1845_dt__update_hcf3_h1, (_1843_dt__update__tmp_h1).Dtor_cf4(), _1844_dt__update_hcf5_h1)
									}(_pat_let50_0)
								}(_1823_i13)
							}(_pat_let49_0)
						}(_pat_let_tv38)
					}(_pat_let44_0)
				}(func(_pat_let45_0 D0) D0 {
					return func(_1839_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let46_0 _dafny.Int) D0 {
							return func(_1840_dt__update_hcf3_h0 _dafny.Int) D0 {
								return func(_pat_let47_0 _dafny.Int) D0 {
									return func(_1841_dt__update_hcf2_h0 _dafny.Int) D0 {
										return func(_pat_let48_0 _dafny.Array) D0 {
											return func(_1842_dt__update_hcf5_h0 _dafny.Array) D0 {
												return Companion_D0_.Create_DC2_(_1841_dt__update_hcf2_h0, _1840_dt__update_hcf3_h0, (_1839_dt__update__tmp_h0).Dtor_cf4(), _1842_dt__update_hcf5_h0)
											}(_pat_let48_0)
										}(_pat_let_tv37)
									}(_pat_let47_0)
								}((_dafny.Zero).Minus(_dafny.IntOfInt64(-269)))
							}(_pat_let46_0)
						}(_pat_let_tv36)
					}(_pat_let45_0)
				}(_1706_v2)))
				_1838_v101 = _nw331
				var _1846_v102 _dafny.Map
				_ = _1846_v102
				_1846_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if false {
						return _dafny.IntOfInt64(-131)
					}
					return p3
				})(), (func() *C6 {
					if false {
						return _1838_v101
					}
					return _1838_v101
				})())
				var _1847_v103 _dafny.Sequence
				_ = _1847_v103
				_1847_v103 = _dafny.SeqOf(_1838_v101)
				_1838_v101 = (func() *C6 {
					if (_1846_v102).Contains(_dafny.IntOfInt64(153)) {
						return (_1846_v102).Get(_dafny.IntOfInt64(153)).(*C6)
					}
					return (_1847_v103).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_1832_v95).Contains(p3) {
							return (_1832_v95).Get(p3).(_dafny.Int)
						}
						return (_dafny.Zero).Minus((_1837_v100).Cardinality())
					})(), _dafny.IntOfUint32((_1847_v103).Cardinality()))).Uint32()).(*C6)
				})()
				var _1848_v104 _dafny.CodePoint
				_ = _1848_v104
				_1848_v104 = _dafny.CodePoint('p')
				var _1849_v105 _dafny.MultiSet
				_ = _1849_v105
				_1849_v105 = _dafny.MultiSetOf(!(!_dafny.Companion_Sequence_.Contains(p0, _1848_v104)))
				_1849_v105 = (_1849_v105).Intersection((_1849_v105).Update(_this.F16(), Companion_Default___.Abs(Companion_Default___.Fm3(globalState))))
				var _1850_v106 _dafny.Sequence
				_ = _1850_v106
				_1850_v106 = _dafny.SeqOf(p2)
				var _1851_v107 _dafny.Map
				_ = _1851_v107
				_1851_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _dafny.Companion_Sequence_.Concatenate(_1850_v106, _1850_v106))
				_1851_v107 = (_1851_v107).Update(_dafny.Companion_Sequence_.Equal(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(77))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg78 _dafny.Int) interface{} {
						return coer78(arg78)
					}
				}((func(_1852_v104 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1853_i15 _dafny.Int) _dafny.CodePoint {
						return _1852_v104
					}
				})(_1848_v104)))), _1850_v106)
			}
		}
		var _1854_v108 _dafny.Sequence
		_ = _1854_v108
		_1854_v108 = _dafny.SeqOf(false)
		r0 = _1854_v108
		return r0
	}
}
func (_this *C7) M2(globalState *GlobalState) {
	{
		var _1855_v0 _dafny.Int
		_ = _1855_v0
		_1855_v0 = _dafny.IntOfInt64(411)
		var _1856_v1 _dafny.Sequence
		_ = _1856_v1
		_1856_v1 = _dafny.SeqOf(_this.F16())
		var _1857_v2 _dafny.MultiSet
		_ = _1857_v2
		_1857_v2 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1856_v1).Cardinality())), _1855_v0, _1855_v0, _1855_v0, _1855_v0)
		var _1858_v3 _dafny.Array
		_ = _1858_v3
		var _nwElement0_57 bool = _this.F16()
		_ = _nwElement0_57
		var _nw332 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(25))
		_ = _nw332
		(_nw332).ArraySet1(_nwElement0_57, 0)
		(_nw332).ArraySet1(_this.F16(), 1)
		(_nw332).ArraySet1(_this.F16(), 2)
		(_nw332).ArraySet1(_this.F16(), 3)
		(_nw332).ArraySet1(_this.F16(), 4)
		(_nw332).ArraySet1(_this.F16(), 5)
		(_nw332).ArraySet1(false, 6)
		(_nw332).ArraySet1((_1857_v2).IsSubsetOf(_1857_v2), 7)
		(_nw332).ArraySet1((_1855_v0).Cmp(_1855_v0) > 0, 8)
		(_nw332).ArraySet1(_this.F16(), 9)
		(_nw332).ArraySet1(_this.F16(), 10)
		(_nw332).ArraySet1(_this.F16(), 11)
		(_nw332).ArraySet1(_this.F16(), 12)
		(_nw332).ArraySet1(_this.F16(), 13)
		(_nw332).ArraySet1(Companion_Default___.Fm0(!(_this.F16()), _this.F16(), globalState), 14)
		(_nw332).ArraySet1((_1857_v2).IsSubsetOf(_1857_v2), 15)
		(_nw332).ArraySet1(_this.F16(), 16)
		(_nw332).ArraySet1((_this.F16()) == (_this.F16()), 17)
		(_nw332).ArraySet1((_1856_v1).Select((Companion_Default___.SafeIndex(_1855_v0, _dafny.IntOfUint32((_1856_v1).Cardinality()))).Uint32()).(bool), 18)
		(_nw332).ArraySet1(_this.F16(), 19)
		(_nw332).ArraySet1(false, 20)
		(_nw332).ArraySet1(_this.F16(), 21)
		(_nw332).ArraySet1(_this.F16(), 22)
		(_nw332).ArraySet1(_this.F16(), 23)
		(_nw332).ArraySet1(_this.F16(), 24)
		_1858_v3 = _nw332
		var _1859_v4 D1
		_ = _1859_v4
		_1859_v4 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(520), false)
		var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))
		_ = _index325
		(_1858_v3).ArraySet1((func() bool {
			if _this.F16() {
				return _this.F16()
			}
			return (_1859_v4).Dtor_cf11()
		})(), (_index325).Int())
		var _1860_v5 D0
		_ = _1860_v5
		_1860_v5 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(184), (_1857_v2).Cardinality(), _1855_v0, _1858_v3)
		var _1861_v6 _dafny.Sequence
		_ = _1861_v6
		_1861_v6 = _dafny.UnicodeSeqOfUtf8Bytes("bv")
		var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))
		_ = _index326
		var _rhs360 _dafny.Int = (_1860_v5).Dtor_cf2()
		_ = _rhs360
		var _rhs361 bool = ((_1855_v0).Minus(_1855_v0)).Cmp(_dafny.IntOfUint32((_1861_v6).Cardinality())) == 0
		_ = _rhs361
		var _lhs279 _dafny.Array = _1858_v3
		_ = _lhs279
		var _lhs280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))
		_ = _lhs280
		_1855_v0 = _rhs360
		(_lhs279).ArraySet1(_rhs361, (_lhs280).Int())
		var _1862_i0 _dafny.Int
		_ = _1862_i0
		_1862_i0 = _dafny.Zero
		{
			for _this.F16() {
				{
					if (_1862_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1862_i0 = (_1862_i0).Plus(_dafny.One)
					var _1863_v7 _dafny.CodePoint
					_ = _1863_v7
					_1863_v7 = _dafny.CodePoint('r')
					var _1864_v8 _dafny.Map
					_ = _1864_v8
					_1864_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pdxlfemr"), _1863_v7)
					_1864_v8 = (_1864_v8).Update(_1861_v6, _1863_v7)
					_1858_v3 = _1858_v3
					if (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool) {
						var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _index327
						((_this).F20()).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(533))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg79 _dafny.Int) interface{} {
								return coer79(arg79)
							}
						}((func(_1865_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1866_i1 _dafny.Int) _dafny.Int {
								return _1865_v0
							}
						})(_1855_v0)))).Cardinality()), (_index327).Int())
						var _1867_v9 _dafny.Array
						_ = _1867_v9
						var _len0_53 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_53
						var _nw333 _dafny.Array
						_ = _nw333
						if _len0_53.Cmp(_dafny.Zero) == 0 {
							_nw333 = _dafny.NewArray(_len0_53)
						} else {
							var _init53 func(_dafny.Int) _dafny.Sequence = (func(_1868_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1869_i2 _dafny.Int) _dafny.Sequence {
									return _1868_v6
								}
							})(_1861_v6)
							_ = _init53
							var _element0_53 = _init53(_dafny.Zero)
							_ = _element0_53
							_nw333 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
							(_nw333).ArraySet1(_element0_53, 0)
							var _nativeLen0_53 = (_len0_53).Int()
							_ = _nativeLen0_53
							for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
								(_nw333).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
							}
						}
						_1867_v9 = _nw333
						var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_1867_v9), 0))
						_ = _index328
						(_1867_v9).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xmofc"), (_index328).Int())
						var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _index329
						((_this).F20()).ArraySet1(_1855_v0, (_index329).Int())
						var _1870_v10 _dafny.Sequence
						_ = _1870_v10
						_1870_v10 = _dafny.SeqOf(_1861_v6, _1861_v6)
						var _1871_v11 _dafny.Set
						_ = _1871_v11
						_1871_v11 = _dafny.SetOf(_this.F16(), (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool), _this.F16())
						var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _index330
						var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_1867_v9), 0))
						_ = _index331
						var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _index332
						var _rhs362 _dafny.Int = _1855_v0
						_ = _rhs362
						var _rhs363 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1870_v10).Select((Companion_Default___.SafeIndex(_1855_v0, _dafny.IntOfUint32((_1870_v10).Cardinality()))).Uint32()).(_dafny.Sequence), _1861_v6)
						_ = _rhs363
						var _rhs364 _dafny.Int = (_1871_v11).Cardinality()
						_ = _rhs364
						var _rhs365 bool = (!((_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)) || ((_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool))) && ((_1855_v0).Cmp(_1855_v0) >= 0)
						_ = _rhs365
						var _rhs366 _dafny.Int = (_1855_v0).Minus(((_dafny.Zero).Minus(_1855_v0)).Times(_1855_v0))
						_ = _rhs366
						var _lhs281 _dafny.Array = (_this).F20()
						_ = _lhs281
						var _lhs282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _lhs282
						var _lhs283 _dafny.Array = _1867_v9
						_ = _lhs283
						var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_1867_v9), 0))
						_ = _lhs284
						var _lhs285 *GlobalState = globalState
						_ = _lhs285
						var _lhs286 _dafny.Array = (_this).F20()
						_ = _lhs286
						var _lhs287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _lhs287
						(_lhs281).ArraySet1(_rhs362, (_lhs282).Int())
						(_lhs283).ArraySet1(_rhs363, (_lhs284).Int())
						_1855_v0 = _rhs364
						_lhs285.F2 = _rhs365
						(_lhs286).ArraySet1(_rhs366, (_lhs287).Int())
						_1855_v0 = (func() _dafny.Int {
							if (func() bool {
								if _this.F16() {
									return true
								}
								return (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)
							})() {
								return (func() _dafny.Int {
									if (_1857_v2).Contains(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int)) {
										return (_1857_v2).Multiplicity(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int))
									}
									return ((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int)
								})()
							}
							return _1855_v0
						})()
						var _1872_v12 *C0
						_ = _1872_v12
						var _nw334 *C0 = New_C0_()
						_ = _nw334
						_nw334.Ctor__(_dafny.IntOfInt64(742), _1858_v3)
						_1872_v12 = _nw334
						var _1873_v13 _dafny.Map
						_ = _1873_v13
						_1873_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1872_v12, ((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int))
						var _1874_v14 _dafny.Array
						_ = _1874_v14
						var _nwElement0_58 _dafny.Int = ((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int)
						_ = _nwElement0_58
						var _nw335 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(21))
						_ = _nw335
						(_nw335).ArraySet1(_nwElement0_58, 0)
						(_nw335).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), 1)
						(_nw335).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-983))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg80 _dafny.Int) interface{} {
								return coer80(arg80)
							}
						}(func(_1875_i3 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						}))).Cardinality())), 2)
						(_nw335).ArraySet1((_dafny.Zero).Minus(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int)), 3)
						(_nw335).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool))).Cardinality()), 4)
						(_nw335).ArraySet1((_dafny.MultiSetOf(_1855_v0)).Cardinality(), 5)
						(_nw335).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), 6)
						(_nw335).ArraySet1(_1855_v0, 7)
						(_nw335).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), 8)
						(_nw335).ArraySet1(_dafny.IntOfInt64(420), 9)
						(_nw335).ArraySet1((_1873_v13).Cardinality(), 10)
						(_nw335).ArraySet1(_1855_v0, 11)
						(_nw335).ArraySet1(_1855_v0, 12)
						(_nw335).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), 13)
						(_nw335).ArraySet1((_1872_v12).F26(), 14)
						(_nw335).ArraySet1(_1855_v0, 15)
						(_nw335).ArraySet1(_dafny.IntOfInt64(-696), 16)
						(_nw335).ArraySet1(Companion_Default___.Fm3(globalState), 17)
						(_nw335).ArraySet1(_1855_v0, 18)
						(_nw335).ArraySet1((_1872_v12).F26(), 19)
						(_nw335).ArraySet1(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), 20)
						_1874_v14 = _nw335
						var _1876_v15 _dafny.Map
						_ = _1876_v15
						_1876_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1874_v14)
						_1876_v15 = (_1876_v15).Update(_this.F16(), _1874_v14)
						(globalState).F5 = (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)
						var _1877_v17 _dafny.Set
						_ = _1877_v17
						_1877_v17 = _dafny.SetOf((_1872_v12).F26(), (_1872_v12).F26(), (func() _dafny.Set {
							var _coll66 = _dafny.NewBuilder()
							_ = _coll66
							for _iter75 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-397), _dafny.IntOfInt64(-384))); ; {
								_compr_66, _ok75 := _iter75()
								if !_ok75 {
									break
								}
								var _1878_v16 _dafny.Int
								_1878_v16 = interface{}(_compr_66).(_dafny.Int)
								if ((_dafny.IntOfInt64(-397)).Cmp(_1878_v16) <= 0) && ((_1878_v16).Cmp(_dafny.IntOfInt64(-384)) < 0) {
									_coll66.Add(Companion_Default___.SafeModuloInt(_1878_v16, (_1872_v12).F26()))
								}
							}
							return _coll66.ToSet()
						}()).Cardinality())
						var _1879_v18 D9
						_ = _1879_v18
						_1879_v18 = Companion_D9_.Create_DC26_(_dafny.IntOfUint32(((_1870_v10).Select((Companion_Default___.SafeIndex((_1872_v12).F26(), _dafny.IntOfUint32((_1870_v10).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _1872_v12)
						var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _index333
						var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))
						_ = _index334
						var _rhs367 *C0 = (func() *C0 {
							if (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool) {
								return _1872_v12
							}
							return (_1879_v18).Dtor_cf46()
						})()
						_ = _rhs367
						var _rhs368 _dafny.Int = (_1872_v12).F26()
						_ = _rhs368
						var _rhs369 bool = (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)
						_ = _rhs369
						var _rhs370 _dafny.Set = (_1877_v17).Difference(_dafny.SetOf(((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), ((_this).F20()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))).Int()).(_dafny.Int), _1855_v0, (_1872_v12).F26(), _1855_v0))
						_ = _rhs370
						var _lhs288 _dafny.Array = (_this).F20()
						_ = _lhs288
						var _lhs289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen(((_this).F20()), 0))
						_ = _lhs289
						var _lhs290 _dafny.Array = _1858_v3
						_ = _lhs290
						var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))
						_ = _lhs291
						_1872_v12 = _rhs367
						(_lhs288).ArraySet1(_rhs368, (_lhs289).Int())
						(_lhs290).ArraySet1(_rhs369, (_lhs291).Int())
						_1877_v17 = _rhs370
					} else {
						_1861_v6 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm2(globalState), _1863_v7), _dafny.Companion_Sequence_.Concatenate(_1861_v6, _1861_v6))
						var _1880_v19 _dafny.Sequence
						_ = _1880_v19
						_1880_v19 = _dafny.SeqOf(_1861_v6)
						var _1881_v20 _dafny.Map
						_ = _1881_v20
						_1881_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v19, Companion_Default___.SafeModuloInt(_1855_v0, _1855_v0))
						_1881_v20 = (_1881_v20).Update(_1880_v19, _1855_v0)
						_1861_v6 = (func() _dafny.Sequence {
							if true {
								return _dafny.UnicodeSeqOfUtf8Bytes("yerr")
							}
							return _1861_v6
						})()
						var _1882_v21 _dafny.Map
						_ = _1882_v21
						_1882_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool) {
								return (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)
							}
							return _this.F16()
						})(), ((_dafny.SetOf((_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool), (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool))).Cardinality()).Minus(_dafny.IntOfInt64(37)))
						_1882_v21 = (_1882_v21).Update(!(_this.F16()), Companion_Default___.Fm3(globalState))
						(globalState).F5 = (_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)
					}
					var _1883_v22 _dafny.Set
					_ = _1883_v22
					_1883_v22 = _dafny.SetOf(_1855_v0, _1855_v0, (_dafny.Zero).Minus(_1855_v0))
					var _1884_v23 D9
					_ = _1884_v23
					_1884_v23 = Companion_D9_.Create_DC27_(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("yxw"), _1861_v6), _1883_v22)
					_1884_v23 = _1884_v23
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1885_v24 *C2
		_ = _1885_v24
		var _nw336 *C2 = New_C2_()
		_ = _nw336
		_nw336.Ctor__((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1855_v0, _1855_v0)))
		_1885_v24 = _nw336
		_1858_v3 = _1858_v3
		var _hi17 _dafny.Int = _1855_v0
		_ = _hi17
		for _1886_i4 := _dafny.IntOfInt64(676); _1886_i4.Cmp(_hi17) < 0; _1886_i4 = _1886_i4.Plus(_dafny.One) {
			_1855_v0 = Companion_Default___.SafeModuloInt(_1886_i4, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1885_v24).F25(), _1886_i4)).Cardinality())
			var _1887_v25 _dafny.MultiSet
			_ = _1887_v25
			_1887_v25 = _dafny.MultiSetOf(_this.F16())
			var _1888_v26 _dafny.Sequence
			_ = _1888_v26
			_1888_v26 = _dafny.SeqOf(_dafny.MultiSetOf((_1858_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(847), _dafny.ArrayLen((_1858_v3), 0))).Int()).(bool)), _1887_v25)
			_1887_v25 = (_1888_v26).Select((Companion_Default___.SafeIndex((_1885_v24).F25(), _dafny.IntOfUint32((_1888_v26).Cardinality()))).Uint32()).(_dafny.MultiSet)
			var _1889_v27 _dafny.CodePoint
			_ = _1889_v27
			_1889_v27 = _dafny.CodePoint('l')
			var _1890_v28 _dafny.Set
			_ = _1890_v28
			_1890_v28 = _dafny.SetOf(_1889_v27)
			_1890_v28 = Companion_Default___.Fm27(globalState)
			var _1891_v29 *C6
			_ = _1891_v29
			var _nw337 *C6 = New_C6_()
			_ = _nw337
			_nw337.Ctor__(_1860_v5)
			_1891_v29 = _nw337
		}
		_1855_v0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F17(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-893))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1892_i5 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(109)
		})), (_this).F17()))).Cardinality())
	}
}
func (_this *C7) F20() _dafny.Array {
	{
		return _this._f20
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f16 bool
	_f17 _dafny.Sequence
	F19  _dafny.Set
	_f18 _dafny.Sequence
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f16 = false
	_this._f17 = _dafny.EmptySeq
	_this.F19 = _dafny.EmptySet
	_this._f18 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C8_ struct {
}

var Companion_C8_ = CompanionStruct_C8_{}

func (_this *C8) Equals(other *C8) bool {
	return _this == other
}

func (_this *C8) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C8)
	return ok && _this.Equals(other)
}

func (*C8) String() string {
	return "_module.C8"
}

func Type_C8_() _dafny.TypeDescriptor {
	return type_C8_{}
}

type type_C8_ struct {
}

func (_this type_C8_) Default() interface{} {
	return (*C8)(nil)
}

func (_this type_C8_) String() string {
	return "main.C8"
}
func (_this *C8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F16() bool {
	return _this._f16
}
func (_this *C8) F16_set_(value bool) {
	_this._f16 = value
}
func (_this *C8) F17() _dafny.Sequence {
	return _this._f17
}
func (_this *C8) Ctor__(f18 _dafny.Sequence, f19 _dafny.Set, f16 bool, f17 _dafny.Sequence) {
	{
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this)._f16 = f16
		(_this)._f17 = f17
	}
}
func (_this *C8) Fm5(globalState *GlobalState) bool {
	{
		return _this.F16()
	}
}
func (_this *C8) M1(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _hi18 _dafny.Int = p3
		_ = _hi18
		for _1893_i0 := p3; _1893_i0.Cmp(_hi18) < 0; _1893_i0 = _1893_i0.Plus(_dafny.One) {
			var _1894_v0 _dafny.Map
			_ = _1894_v0
			_1894_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2)
			_1894_v0 = (_1894_v0).Update(_1893_i0, !(p1))
			(globalState).F2 = p2
			var _1895_v1 _dafny.Int
			_ = _1895_v1
			_1895_v1 = _dafny.IntOfInt64(960)
			_1895_v1 = _1893_i0
			var _1896_v2 _dafny.Array
			_ = _1896_v2
			var _nw338 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw338
			_1896_v2 = _nw338
			var _1897_v3 _dafny.Map
			_ = _1897_v3
			_1897_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), _1896_v2)
			_1897_v3 = (_1897_v3).Update(_this.F16(), _1896_v2)
		}
		var _hi19 _dafny.Int = p3
		_ = _hi19
		for _1898_i1 := p3; _1898_i1.Cmp(_hi19) < 0; _1898_i1 = _1898_i1.Plus(_dafny.One) {
			var _1899_v4 _dafny.Int
			_ = _1899_v4
			_1899_v4 = _dafny.IntOfInt64(284)
			var _1900_v5 _dafny.Array
			_ = _1900_v5
			var _nw339 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw339
			_1900_v5 = _nw339
			var _1901_v6 _dafny.Array
			_ = _1901_v6
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_54
			var _nw340 _dafny.Array
			_ = _nw340
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw340 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) bool = (func(_1902_p2 bool) func(_dafny.Int) bool {
					return func(_1903_i2 _dafny.Int) bool {
						return _1902_p2
					}
				})(p2)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw340 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw340).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw340).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_1901_v6 = _nw340
			var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_1901_v6), 0))
			_ = _index335
			(_1901_v6).ArraySet1((_dafny.SetOf(p1, p2)).IsDisjointFrom(_this.F19), (_index335).Int())
			var _1904_v7 D0
			_ = _1904_v7
			_1904_v7 = Companion_D0_.Create_DC0_(Companion_Default___.Fm1(globalState))
			var _1905_v8 _dafny.Map
			_ = _1905_v8
			_1905_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1904_v7)
			var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_1901_v6), 0))
			_ = _index336
			var _rhs371 _dafny.Int = (_1905_v8).Cardinality()
			_ = _rhs371
			var _rhs372 _dafny.Array = _1900_v5
			_ = _rhs372
			var _rhs373 bool = !(p2)
			_ = _rhs373
			var _lhs292 _dafny.Array = _1901_v6
			_ = _lhs292
			var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_1901_v6), 0))
			_ = _lhs293
			_1899_v4 = _rhs371
			_1900_v5 = _rhs372
			(_lhs292).ArraySet1(_rhs373, (_lhs293).Int())
			_1899_v4 = _dafny.IntOfInt64(-278)
			var _1906_v9 *C6
			_ = _1906_v9
			var _nw341 *C6 = New_C6_()
			_ = _nw341
			_nw341.Ctor__(Companion_D0_.Create_DC2_(((_this).F17()).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32()).(_dafny.Int), p3, _1899_v4, _1901_v6))
			_1906_v9 = _nw341
			var _1907_v10 _dafny.Sequence
			_ = _1907_v10
			_1907_v10 = _dafny.SeqOf(_1901_v6)
			(globalState).F5 = _dafny.Companion_Sequence_.Contains(_1907_v10, _1901_v6)
		}
		var _1908_v11 _dafny.Array
		_ = _1908_v11
		var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw342
		_1908_v11 = _nw342
		var _1909_v12 _dafny.Map
		_ = _1909_v12
		_1909_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1908_v11)
		var _hi20 _dafny.Int = (_dafny.Zero).Minus(((_1909_v12).Update(_this.F16(), _1908_v11)).Cardinality())
		_ = _hi20
		for _1910_i3 := p3; _1910_i3.Cmp(_hi20) < 0; _1910_i3 = _1910_i3.Plus(_dafny.One) {
			var _1911_v13 _dafny.Sequence
			_ = _1911_v13
			_1911_v13 = _dafny.SeqOf(p3)
			_1911_v13 = (_this).F17()
			(globalState).F5 = true
			var _1912_v14 T1
			_ = _1912_v14
			var _nw343 *C1 = New_C1_()
			_ = _nw343
			_nw343.Ctor__((_this.F19).Cardinality(), p3, p1, (_this).F17(), p3)
			_1912_v14 = _nw343
			var _1913_v15 _dafny.Map
			_ = _1913_v15
			_1913_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F19).Cardinality(), _1912_v14)
			_1913_v15 = (_1913_v15).Update(p3, _1912_v14)
			if _this.F16() {
				var _1914_v16 D3
				_ = _1914_v16
				_1914_v16 = Companion_D3_.Create_DC10_(true)
				var _1915_v17 _dafny.Sequence
				_ = _1915_v17
				_1915_v17 = _dafny.SeqOf(_1912_v14.F16(), _1912_v14.F16(), (_1914_v16).Dtor_cf19())
				r0 = _dafny.Companion_Sequence_.Update(_1915_v17, (Companion_Default___.SafeIndex((p3).Minus(_dafny.IntOfInt64(835)), _dafny.IntOfUint32((_1915_v17).Cardinality()))).Uint32(), p2)
				var _1916_v18 _dafny.Array
				_ = _1916_v18
				var _nw344 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw344
				_1916_v18 = _nw344
				var _1917_v19 _dafny.CodePoint
				_ = _1917_v19
				_1917_v19 = _dafny.CodePoint('q')
				var _1918_v20 _dafny.MultiSet
				_ = _1918_v20
				_1918_v20 = _dafny.MultiSetOf(_1917_v19)
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1916_v18), 0))
				_ = _index337
				(_1916_v18).ArraySet1((_dafny.MultiSetOf(_1917_v19)).Equals(_1918_v20), (_index337).Int())
				var _1919_v21 _dafny.Sequence
				_ = _1919_v21
				_1919_v21 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_1912_v14).F28(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _1917_v19), (_this).F18(), p0)
				var _1920_v22 D0
				_ = _1920_v22
				_1920_v22 = Companion_D0_.Create_DC1_(p0)
				var _1921_v23 _dafny.Set
				_ = _1921_v23
				_1921_v23 = _dafny.SetOf((_this).F18(), p0, (_1919_v21).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_dafny.SetOf((_1912_v14).F28())).Cardinality())).Cardinality()), _dafny.IntOfUint32((_1919_v21).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("apnbo"), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_this).F18(), (_1920_v22).Dtor_cf1()), (Companion_Default___.SafeIndex((_1912_v14).F28(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).F18(), (_1920_v22).Dtor_cf1())).Cardinality()))).Uint32(), _1917_v19))
				var _1922_v24 _dafny.Int
				_ = _1922_v24
				_1922_v24 = _dafny.IntOfInt64(320)
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1916_v18), 0))
				_ = _index338
				var _rhs374 bool = _1912_v14.F16()
				_ = _rhs374
				var _rhs375 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ktslxunay"), (_this).F18())
				_ = _rhs375
				var _rhs376 _dafny.Set = _1921_v23
				_ = _rhs376
				var _rhs377 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1915_v17, _1915_v17)
				_ = _rhs377
				var _rhs378 _dafny.Int = _1910_i3
				_ = _rhs378
				var _lhs294 _dafny.Array = _1916_v18
				_ = _lhs294
				var _lhs295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1916_v18), 0))
				_ = _lhs295
				var _lhs296 *GlobalState = globalState
				_ = _lhs296
				(_lhs294).ArraySet1(_rhs374, (_lhs295).Int())
				_lhs296.F2 = _rhs375
				_1921_v23 = _rhs376
				_1915_v17 = _rhs377
				_1922_v24 = _rhs378
				_1911_v13 = (_1912_v14).F17()
				var _1923_v25 _dafny.Sequence
				_ = _1923_v25
				_1923_v25 = _dafny.UnicodeSeqOfUtf8Bytes("jco")
				var _1924_v26 _dafny.Set
				_ = _1924_v26
				_1924_v26 = _dafny.SetOf(_1908_v11)
				var _1925_v27 _dafny.MultiSet
				_ = _1925_v27
				_1925_v27 = _dafny.MultiSetOf(p2)
				var _1926_v28 _dafny.MultiSet
				_ = _1926_v28
				_1926_v28 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_1912_v14).F17()).Cardinality()), (_1925_v27).Cardinality(), _1910_i3)
				var _rhs379 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("xfcilcoa")
				_ = _rhs379
				var _rhs380 _dafny.Set = _1924_v26
				_ = _rhs380
				var _rhs381 bool = ((_dafny.MultiSetOf(_1922_v24)).Update(_dafny.IntOfInt64(-137), Companion_Default___.Abs((_dafny.Zero).Minus(p3)))).Equals(_1926_v28)
				_ = _rhs381
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				_1923_v25 = _rhs379
				_1924_v26 = _rhs380
				_lhs297.F2 = _rhs381
				var _1927_v29 _dafny.Map
				_ = _1927_v29
				_1927_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).Fm5(globalState))
				_1927_v29 = (_1927_v29).Update(p2, true)
			} else {
				var _1928_v30 _dafny.CodePoint
				_ = _1928_v30
				_1928_v30 = _dafny.CodePoint('b')
				(globalState).F15 = _1928_v30
				var _1929_v31 _dafny.Int
				_ = _1929_v31
				_1929_v31 = _dafny.IntOfInt64(418)
				var _rhs382 _dafny.Int = (_1912_v14).F28()
				_ = _rhs382
				var _rhs383 _dafny.Int = Companion_Default___.Fm3(globalState)
				_ = _rhs383
				_1929_v31 = _rhs382
				_1929_v31 = _rhs383
				var _1930_v32 _dafny.Array
				_ = _1930_v32
				var _nw345 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw345
				_1930_v32 = _nw345
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_1930_v32), 0))
				_ = _index339
				(_1930_v32).ArraySet1(((_1912_v14).F28()).Cmp((_1912_v14).F28()) < 0, (_index339).Int())
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_1930_v32), 0))
				_ = _index340
				(_1930_v32).ArraySet1(p2, (_index340).Int())
				var _rhs384 _dafny.Int = (_dafny.IntOfUint32(((_this).F18()).Cardinality())).Minus((p3).Times(_1910_i3))
				_ = _rhs384
				var _rhs385 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lvfjtwdb")).Cardinality()), ((_1912_v14).F28()).Plus(_1910_i3))
				_ = _rhs385
				var _rhs386 bool = (!(_1912_v14.F16())) == ((p1) && (!((_1930_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(252), _dafny.ArrayLen((_1930_v32), 0))).Int()).(bool))))
				_ = _rhs386
				var _rhs387 _dafny.Int = (_dafny.Zero).Minus(_1929_v31)
				_ = _rhs387
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				_1929_v31 = _rhs384
				_1929_v31 = _rhs385
				_lhs298.F2 = _rhs386
				_1929_v31 = _rhs387
				var _1931_v33 *C3
				_ = _1931_v33
				var _nw346 *C3 = New_C3_()
				_ = _nw346
				_nw346.Ctor__((_this).F18(), _1912_v14.F16(), (_1912_v14).F17())
				_1931_v33 = _nw346
			}
		}
		if _this.F16() {
			var _1932_v34 _dafny.Map
			_ = _1932_v34
			_1932_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1908_v11, p0)
			var _1933_v35 _dafny.Map
			_ = _1933_v35
			_1933_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1934_v36 _dafny.Map
			_ = _1934_v36
			_1934_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (func() _dafny.Sequence {
				if (_1933_v35).Contains(p1) {
					return (_1933_v35).Get(p1).(_dafny.Sequence)
				}
				return p0
			})())
			_1932_v34 = (_1932_v34).Update(_1908_v11, (func() _dafny.Sequence {
				if (_1933_v35).Contains(p2) {
					return (_1933_v35).Get(p2).(_dafny.Sequence)
				}
				return p0
			})())
			var _1935_v37 _dafny.Sequence
			_ = _1935_v37
			_1935_v37 = _dafny.SeqOf(_this.F16())
			var _1936_v38 _dafny.Set
			_ = _1936_v38
			_1936_v38 = _dafny.SetOf(_dafny.IntOfUint32((_1935_v37).Cardinality()), (_dafny.Zero).Minus(p3), _dafny.IntOfInt64(-403))
			_1936_v38 = (_1936_v38).Union(Companion_Default___.Fm31(p2, globalState))
			var _1937_v39 _dafny.Array
			_ = _1937_v39
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_55
			var _nw347 _dafny.Array
			_ = _nw347
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw347 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) _dafny.CodePoint = func(_1938_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw347 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw347).ArraySet1CodePoint(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw347).ArraySet1CodePoint(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_1937_v39 = _nw347
			var _1939_v40 _dafny.Sequence
			_ = _1939_v40
			_1939_v40 = _dafny.SeqOf(_1937_v39, _1937_v39, _1937_v39, _1937_v39)
			_1939_v40 = _dafny.Companion_Sequence_.Concatenate(_1939_v40, _1939_v40)
			var _1940_v41 _dafny.Array
			_ = _1940_v41
			var _nw348 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw348
			_1940_v41 = _nw348
			var _1941_v42 D11
			_ = _1941_v42
			_1941_v42 = Companion_D11_.Create_DC31_(p1, _dafny.IntOfInt64(919), p3, p1)
			var _1942_v43 _dafny.Map
			_ = _1942_v43
			_1942_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D11_.Create_DC32_(_1941_v42), _1940_v41)
			var _1943_v44 D11
			_ = _1943_v44
			_1943_v44 = Companion_D11_.Create_DC32_(_1941_v42)
			var _1944_v45 _dafny.MultiSet
			_ = _1944_v45
			_1944_v45 = _dafny.MultiSetOf(_dafny.IntOfInt64(762))
			var _1945_v46 D0
			_ = _1945_v46
			_1945_v46 = Companion_D0_.Create_DC2_((_1944_v45).Cardinality(), p3, p3, _1940_v41)
			var _rhs388 _dafny.Array = (func() _dafny.Array {
				if (_1942_v43).Contains(_1943_v44) {
					return (_1942_v43).Get(_1943_v44).(_dafny.Array)
				}
				return (_1945_v46).Dtor_cf5()
			})()
			_ = _rhs388
			var _rhs389 _dafny.Array = _1908_v11
			_ = _rhs389
			_1940_v41 = _rhs388
			_1908_v11 = _rhs389
			var _1946_v47 _dafny.Int
			_ = _1946_v47
			_1946_v47 = _dafny.IntOfInt64(-83)
			var _1947_v48 D11
			_ = _1947_v48
			_1947_v48 = Companion_D11_.Create_DC29_(Companion_Default___.Fm2(globalState))
			var _pat_let_tv39 = _1941_v42
			_ = _pat_let_tv39
			var _rhs390 _dafny.Int = _1946_v47
			_ = _rhs390
			var _rhs391 D11 = _1947_v48
			_ = _rhs391
			var _rhs392 _dafny.Int = _1946_v47
			_ = _rhs392
			var _rhs393 D11 = func(_pat_let51_0 D11) D11 {
				return func(_1948_dt__update__tmp_h0 D11) D11 {
					return func(_pat_let52_0 D11) D11 {
						return func(_1949_dt__update_hcf57_h0 D11) D11 {
							return Companion_D11_.Create_DC32_(_1949_dt__update_hcf57_h0)
						}(_pat_let52_0)
					}(_pat_let_tv39)
				}(_pat_let51_0)
			}(_1943_v44)
			_ = _rhs393
			var _rhs394 bool = true
			_ = _rhs394
			var _lhs299 *C8 = _this
			_ = _lhs299
			_1946_v47 = _rhs390
			_1947_v48 = _rhs391
			_1946_v47 = _rhs392
			_1943_v44 = _rhs393
			_lhs299.F16_set_(_rhs394)
		} else {
			if p2 {
				var _1950_v49 _dafny.Sequence
				_ = _1950_v49
				_1950_v49 = _dafny.UnicodeSeqOfUtf8Bytes("ebxvgp")
				var _1951_v50 _dafny.Map
				_ = _1951_v50
				_1951_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1950_v49)
				var _1952_v51 _dafny.Sequence
				_ = _1952_v51
				_1952_v51 = _dafny.SeqOf(_1950_v49)
				var _1953_v52 _dafny.MultiSet
				_ = _1953_v52
				_1953_v52 = _dafny.MultiSetOf(p1)
				_1950_v49 = (func() _dafny.Sequence {
					if (_1951_v50).Contains(Companion_Default___.SafeModuloInt(p3, Companion_Default___.Fm3(globalState))) {
						return (_1951_v50).Get(Companion_Default___.SafeModuloInt(p3, Companion_Default___.Fm3(globalState))).(_dafny.Sequence)
					}
					return (func() _dafny.Sequence {
						if _this.F16() {
							return (_1952_v51).Select((Companion_Default___.SafeIndex((_1953_v52).Cardinality(), _dafny.IntOfUint32((_1952_v51).Cardinality()))).Uint32()).(_dafny.Sequence)
						}
						return (_1952_v51).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1952_v51).Cardinality()))).Uint32()).(_dafny.Sequence)
					})()
				})()
				var _1954_v53 _dafny.Array
				_ = _1954_v53
				var _nw349 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw349
				_1954_v53 = _nw349
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1954_v53), 0))
				_ = _index341
				(_1954_v53).ArraySet1((false) && (p2), (_index341).Int())
				var _1955_v54 _dafny.Map
				_ = _1955_v54
				_1955_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), p1)
				var _1956_v55 D1
				_ = _1956_v55
				_1956_v55 = Companion_D1_.Create_DC5_(((_1955_v54).Update(p2, p1)).Cardinality(), p2)
				var _1957_v56 _dafny.Sequence
				_ = _1957_v56
				_1957_v56 = _dafny.SeqOf(_this.F16(), Companion_Default___.Fm0(_this.F16(), _this.F16(), globalState), _this.F16(), p2)
				var _1958_v57 *C5
				_ = _1958_v57
				var _nw350 *C5 = New_C5_()
				_ = _nw350
				_nw350.Ctor__(_this.F16(), (_this).F17())
				_1958_v57 = _nw350
				var _1959_v58 _dafny.Map
				_ = _1959_v58
				_1959_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1958_v57, p3)
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1954_v53), 0))
				_ = _index342
				(_1954_v53).ArraySet1((func() bool {
					if (_1956_v55).Dtor_cf11() {
						return (_1957_v56).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_1959_v58).Contains(_1958_v57) {
								return (_1959_v58).Get(_1958_v57).(_dafny.Int)
							}
							return p3
						})(), _dafny.IntOfUint32((_1957_v56).Cardinality()))).Uint32()).(bool)
					}
					return p1
				})(), (_index342).Int())
				var _1960_v59 _dafny.Map
				_ = _1960_v59
				_1960_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32(((_this).F18()).Cardinality()))
				var _1961_v60 _dafny.Array
				_ = _1961_v60
				var _nwElement0_59 _dafny.Sequence = _dafny.SeqOf(p3, (_dafny.Zero).Minus(p3), p3, p3, p3)
				_ = _nwElement0_59
				var _nw351 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(20))
				_ = _nw351
				(_nw351).ArraySet1(_nwElement0_59, 0)
				(_nw351).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(-740)), 1)
				(_nw351).ArraySet1((_this).F17(), 2)
				(_nw351).ArraySet1((_this).F17(), 3)
				(_nw351).ArraySet1((_this).F17(), 4)
				(_nw351).ArraySet1((_this).F17(), 5)
				(_nw351).ArraySet1((_this).F17(), 6)
				(_nw351).ArraySet1((_this).F17(), 7)
				(_nw351).ArraySet1((_this).F17(), 8)
				(_nw351).ArraySet1(_dafny.SeqOf(p3, p3), 9)
				(_nw351).ArraySet1((_this).F17(), 10)
				(_nw351).ArraySet1((_this).F17(), 11)
				(_nw351).ArraySet1((_this).F17(), 12)
				(_nw351).ArraySet1(_dafny.SeqOf(p3), 13)
				(_nw351).ArraySet1((_this).F17(), 14)
				(_nw351).ArraySet1(_dafny.Companion_Sequence_.Update((_this).F17(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(474), _dafny.IntOfUint32(((_this).F17()).Cardinality()))).Uint32(), p3), 15)
				(_nw351).ArraySet1((_this).F17(), 16)
				(_nw351).ArraySet1((_this).F17(), 17)
				(_nw351).ArraySet1((_this).F17(), 18)
				(_nw351).ArraySet1((_this).F17(), 19)
				_1961_v60 = _nw351
				var _1962_v61 D2
				_ = _1962_v61
				_1962_v61 = Companion_D2_.Create_DC8_(_1960_v59, _1961_v60, p1, _this.F16(), p3)
				var _1963_v62 D6
				_ = _1963_v62
				_1963_v62 = Companion_D6_.Create_DC17_(_1962_v61)
				_1963_v62 = _1963_v62
				var _1964_v63 _dafny.Set
				_ = _1964_v63
				_1964_v63 = _dafny.SetOf(p3)
				var _1965_v64 _dafny.Map
				_ = _1965_v64
				_1965_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1964_v63, _1908_v11)
				_1965_v64 = (_1965_v64).Update(func() _dafny.Set {
					var _coll67 = _dafny.NewBuilder()
					_ = _coll67
					for _iter76 := _dafny.Iterate((_1964_v63).Elements()); ; {
						_compr_67, _ok76 := _iter76()
						if !_ok76 {
							break
						}
						var _1966_v65 _dafny.Int
						_1966_v65 = interface{}(_compr_67).(_dafny.Int)
						if (_1964_v63).Contains(_1966_v65) {
							_coll67.Add(Companion_Default___.SafeDivisionInt(_1966_v65, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(439), !(true)))).Cardinality(), false))).Cardinality())))
						}
					}
					return _coll67.ToSet()
				}(), _1908_v11)
				var _1967_v66 _dafny.Array
				_ = _1967_v66
				var _nwElement0_60 _dafny.Set = _this.F19
				_ = _nwElement0_60
				var _nw352 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.One)
				_ = _nw352
				(_nw352).ArraySet1(_nwElement0_60, 0)
				_1967_v66 = _nw352
				_1967_v66 = _1967_v66
			} else {
				(globalState).F2 = _dafny.Companion_Sequence_.IsProperPrefixOf(p0, (_this).F18())
				var _1968_v67 D11
				_ = _1968_v67
				_1968_v67 = Companion_D11_.Create_DC31_(_this.F16(), p3, p3, true)
				_1968_v67 = Companion_D11_.Create_DC31_(_this.F16(), p3, p3, _this.F16())
				var _1969_v68 _dafny.Int
				_ = _1969_v68
				_1969_v68 = _dafny.IntOfInt64(248)
				_1969_v68 = _1969_v68
				_1969_v68 = (func() _dafny.Int {
					if !(_this.F16()) {
						return Companion_Default___.Fm3(globalState)
					}
					return _dafny.IntOfInt64(6)
				})()
				(globalState).F15 = ((_this).F18()).Select((Companion_Default___.SafeIndex(_1969_v68, _dafny.IntOfUint32(((_this).F18()).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
			var _1970_v69 _dafny.Array
			_ = _1970_v69
			var _len0_56 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_56
			var _nw353 _dafny.Array
			_ = _nw353
			if _len0_56.Cmp(_dafny.Zero) == 0 {
				_nw353 = _dafny.NewArray(_len0_56)
			} else {
				var _init56 func(_dafny.Int) bool = func(_1971_i5 _dafny.Int) bool {
					return false
				}
				_ = _init56
				var _element0_56 = _init56(_dafny.Zero)
				_ = _element0_56
				_nw353 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
				(_nw353).ArraySet1(_element0_56, 0)
				var _nativeLen0_56 = (_len0_56).Int()
				_ = _nativeLen0_56
				for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
					(_nw353).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
				}
			}
			_1970_v69 = _nw353
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1970_v69), 0))
			_ = _index343
			(_1970_v69).ArraySet1(Companion_Default___.Fm0(p2, p2, globalState), (_index343).Int())
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_1970_v69), 0))
			_ = _index344
			(_1970_v69).ArraySet1(p2, (_index344).Int())
			var _1972_v70 _dafny.Int
			_ = _1972_v70
			_1972_v70 = _dafny.IntOfInt64(-337)
			var _rhs395 _dafny.Int = ((_dafny.Zero).Minus(p3)).Times(_1972_v70)
			_ = _rhs395
			var _rhs396 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(15), p3)
			_ = _rhs396
			_1972_v70 = _rhs395
			_1972_v70 = _rhs396
			var _1973_v71 _dafny.Sequence
			_ = _1973_v71
			_1973_v71 = _dafny.UnicodeSeqOfUtf8Bytes("dlvn")
			_1973_v71 = _dafny.Companion_Sequence_.Concatenate((_this).F18(), (_this).F18())
			_1972_v70 = (_dafny.IntOfUint32(((_this).F17()).Cardinality())).Plus(p3)
		}
		for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1908_v11), 0))); ; {
			_guard_loop_9, _ok77 := _iter77()
			if !_ok77 {
				break
			}
			var _1974_i6 _dafny.Int
			_1974_i6 = interface{}(_guard_loop_9).(_dafny.Int)
			if (true) && (((_1974_i6).Sign() != -1) && ((_1974_i6).Cmp(_dafny.ArrayLen((_1908_v11), 0)) < 0)) {
				(_1908_v11).ArraySet1((_1974_i6).Times(p3), (_1974_i6).Int())
			}
		}
		var _1975_v72 _dafny.Int
		_ = _1975_v72
		_1975_v72 = _dafny.IntOfInt64(-430)
		var _1976_v73 _dafny.Sequence
		_ = _1976_v73
		_1976_v73 = _dafny.UnicodeSeqOfUtf8Bytes("dmsoc")
		var _rhs397 _dafny.Array = (func() _dafny.Array {
			if _this.F16() {
				return _1908_v11
			}
			return _1908_v11
		})()
		_ = _rhs397
		var _rhs398 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm20(_this.F16(), !(p2), true, _1975_v72, globalState), _dafny.SeqOf(p1, _this.F16()))).Cardinality())
		_ = _rhs398
		var _rhs399 bool = p2
		_ = _rhs399
		var _rhs400 bool = !(_this.F16())
		_ = _rhs400
		var _rhs401 _dafny.Sequence = p0
		_ = _rhs401
		var _lhs300 *GlobalState = globalState
		_ = _lhs300
		var _lhs301 *GlobalState = globalState
		_ = _lhs301
		_1908_v11 = _rhs397
		_1975_v72 = _rhs398
		_lhs300.F5 = _rhs399
		_lhs301.F5 = _rhs400
		_1976_v73 = _rhs401
		var _1977_v74 _dafny.Sequence
		_ = _1977_v74
		_1977_v74 = _dafny.SeqOf(p2, _dafny.Companion_Sequence_.IsPrefixOf(p0, _dafny.UnicodeSeqOfUtf8Bytes("c")), p1)
		r0 = _1977_v74
		return r0
	}
}
func (_this *C8) F18() _dafny.Sequence {
	{
		return _this._f18
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
