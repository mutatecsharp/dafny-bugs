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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) bool {
	return ((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(855), _dafny.CodePoint('c'))).Cardinality(), _dafny.IntOfInt64(-525), (_dafny.Zero).Minus(_dafny.IntOfInt64(-448)))).Union(_dafny.SetOf(_dafny.IntOfInt64(289), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false, !(true), false, true)).Cardinality()))).Cardinality()))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dfanxb")).Cardinality()), _dafny.IntOfInt64(-110), _dafny.IntOfInt64(-828)))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('i')
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), true)).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.CodePoint
			_0_v0 = interface{}(_compr_0).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), true)).Contains(_0_v0) {
				_coll0.Add(_0_v0, false)
			}
		}
		return _coll0.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), true))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(406), _dafny.IntOfInt64(244), _dafny.IntOfInt64(694), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tudpflvdl")).Cardinality()))).Cardinality())).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(329), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nmmeb")).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdcukv")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_dafny.SetOf(_dafny.IntOfInt64(831), (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(251))).Cardinality())).Cardinality())).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("famu")).Cardinality()), _dafny.IntOfUint32(((Companion_D3_.Create_DC8_(_dafny.UnicodeSeqOfUtf8Bytes("fsdvonkfs"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality())
	})))).Dtor_cf13()).Cardinality())))), true)
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(!(false))).Union(_dafny.MultiSetOf(!(true), false, !(true)))).Union((_dafny.MultiSetOf(true, true, false, true, false)).Intersection(_dafny.MultiSetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(_dafny.UnicodeSeqOfUtf8Bytes("jucvf"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-165))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(604), _dafny.IntOfInt64(866)), _dafny.CodePoint('u'))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_3_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())
	}))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 D2, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	})), _dafny.UnicodeSeqOfUtf8Bytes("bxf"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("epp"), _dafny.UnicodeSeqOfUtf8Bytes("kjb")))).Merge((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("b"), !(false))).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v0 _dafny.Sequence
			_5_v0 = interface{}(_compr_1).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("b"), !(false))).Contains(_5_v0) {
				_coll1.Add(_5_v0, _dafny.UnicodeSeqOfUtf8Bytes("lcmexjuub"))
			}
		}
		return _coll1.ToMap()
	}()).Merge((Companion_D6_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-223))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_6_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.UnicodeSeqOfUtf8Bytes("lsqfg")))).Dtor_cf29()))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("obosvty"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))), false)
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC7_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-538))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(616))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_9_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})), _dafny.UnicodeSeqOfUtf8Bytes("fn")))
}
func (_static *CompanionStruct_Default___) Fm11(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))), (func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(-84)
		}
		return _dafny.IntOfInt64(-734)
	})())
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Int, globalState *GlobalState) {
	var _hi0 _dafny.Int = (p3).Plus(p0)
	_ = _hi0
	for _10_i0 := p3; _10_i0.Cmp(_hi0) < 0; _10_i0 = _10_i0.Plus(_dafny.One) {
		var _source0 D2 = Companion_D2_.Create_DC6_((_10_i0).Cmp(p0) != 0)
		_ = _source0
		if _source0.Is_DC6() {
			var _11___mcc_h0 bool = _source0.Get_().(D2_DC6).Cf10
			_ = _11___mcc_h0
			var _12_cf10 bool = _11___mcc_h0
			_ = _12_cf10
			var _13_v0 *C0
			_ = _13_v0
			var _nw0 *C0 = New_C0_()
			_ = _nw0
			_nw0.Ctor__(p0)
			_13_v0 = _nw0
			var _14_v1 _dafny.Array
			_ = _14_v1
			var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw1
			_14_v1 = _nw1
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_14_v1), 0))
			_ = _index0
			(_14_v1).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm10(!(p2), globalState)).Cardinality()), (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_14_v1), 0))
			_ = _index1
			(_14_v1).ArraySet1(Companion_Default___.SafeModuloInt((p0).Plus(_10_i0), p3), (_index1).Int())
			var _15_v2 _dafny.Sequence
			_ = _15_v2
			_15_v2 = _dafny.SeqOf(p2)
			_15_v2 = _15_v2
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_14_v1), 0))
			_ = _index2
			(_14_v1).ArraySet1(_dafny.IntOfInt64(-844), (_index2).Int())
		} else {
			var _16___mcc_h1 _dafny.Sequence = _source0.Get_().(D2_DC5).Cf9
			_ = _16___mcc_h1
			var _17_cf9 _dafny.Sequence = _16___mcc_h1
			_ = _17_cf9
			var _18_v3 _dafny.Array
			_ = _18_v3
			var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw2
			_18_v3 = _nw2
			var _19_v4 _dafny.Array
			_ = _19_v4
			var _nw3 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw3
			_19_v4 = _nw3
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_18_v3), 0))
			_ = _index3
			(_18_v3).ArraySet1(_19_v4, (_index3).Int())
			var _20_v5 _dafny.Array
			_ = _20_v5
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_0
			var _nw4 _dafny.Array
			_ = _nw4
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw4 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = func(_21_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_21_i1, _dafny.IntOfInt64(805))
				}
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw4 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw4).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw4).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_20_v5 = _nw4
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_20_v5), 0))
			_ = _index4
			(_20_v5).ArraySet1(_10_i0, (_index4).Int())
			var _22_v6 _dafny.Sequence
			_ = _22_v6
			_22_v6 = _dafny.UnicodeSeqOfUtf8Bytes("cfaaiqx")
			var _23_v7 _dafny.Map
			_ = _23_v7
			_23_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_17_cf9).Select((Companion_Default___.SafeIndex(_10_i0, _dafny.IntOfUint32((_17_cf9).Cardinality()))).Uint32()).(bool) {
					return p2
				}
				return p2
			})(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ma"), _22_v6))
			var _24_v8 _dafny.Map
			_ = _24_v8
			_24_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _25_v9 _dafny.CodePoint
			_ = _25_v9
			_25_v9 = _dafny.CodePoint('p')
			var _26_v10 _dafny.Set
			_ = _26_v10
			_26_v10 = _dafny.SetOf(p2)
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_18_v3), 0))
			_ = _index5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_20_v5), 0))
			_ = _index6
			var _rhs0 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_23_v7).Contains((func() bool {
					if (_24_v8).Contains(p2) {
						return (_24_v8).Get(p2).(bool)
					}
					return p2
				})()) {
					return (_23_v7).Get((func() bool {
						if (_24_v8).Contains(p2) {
							return (_24_v8).Get(p2).(bool)
						}
						return p2
					})()).(_dafny.Sequence)
				}
				return _22_v6
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_23_v7).Contains((func() bool {
					if (_24_v8).Contains(p2) {
						return (_24_v8).Get(p2).(bool)
					}
					return p2
				})()) {
					return (_23_v7).Get((func() bool {
						if (_24_v8).Contains(p2) {
							return (_24_v8).Get(p2).(bool)
						}
						return p2
					})()).(_dafny.Sequence)
				}
				return _22_v6
			})()).Cardinality()))).Uint32(), _25_v9), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(globalState), _dafny.IntOfUint32((_22_v6).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_23_v7).Contains((func() bool {
					if (_24_v8).Contains(p2) {
						return (_24_v8).Get(p2).(bool)
					}
					return p2
				})()) {
					return (_23_v7).Get((func() bool {
						if (_24_v8).Contains(p2) {
							return (_24_v8).Get(p2).(bool)
						}
						return p2
					})()).(_dafny.Sequence)
				}
				return _22_v6
			})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(p3)), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_23_v7).Contains((func() bool {
					if (_24_v8).Contains(p2) {
						return (_24_v8).Get(p2).(bool)
					}
					return p2
				})()) {
					return (_23_v7).Get((func() bool {
						if (_24_v8).Contains(p2) {
							return (_24_v8).Get(p2).(bool)
						}
						return p2
					})()).(_dafny.Sequence)
				}
				return _22_v6
			})()).Cardinality()))).Uint32(), _25_v9)).Cardinality()))).Uint32(), _25_v9)
			_ = _rhs0
			var _rhs1 _dafny.Array = _19_v4
			_ = _rhs1
			var _rhs2 _dafny.Int = _10_i0
			_ = _rhs2
			var _rhs3 _dafny.Int = (func() _dafny.Int {
				if p2 {
					return (func() _dafny.Int {
						if p2 {
							return p3
						}
						return (_26_v10).Cardinality()
					})()
				}
				return _10_i0
			})()
			_ = _rhs3
			var _lhs0 *GlobalState = globalState
			_ = _lhs0
			var _lhs1 _dafny.Array = _18_v3
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_18_v3), 0))
			_ = _lhs2
			var _lhs3 _dafny.Array = _20_v5
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_20_v5), 0))
			_ = _lhs4
			var _lhs5 *GlobalState = globalState
			_ = _lhs5
			_lhs0.F4 = _rhs0
			(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
			(_lhs3).ArraySet1(_rhs2, (_lhs4).Int())
			_lhs5.F7 = _rhs3
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_19_v4), 0))
			_ = _index7
			(_19_v4).ArraySet1(p2, (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_19_v4), 0))
			_ = _index8
			(_19_v4).ArraySet1((func() bool {
				if false {
					return p2
				}
				return p2
			})(), (_index8).Int())
			var _27_v11 _dafny.Map
			_ = _27_v11
			_27_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v9, p3)
			var _28_v12 D0
			_ = _28_v12
			_28_v12 = Companion_D0_.Create_DC2_(p0)
			var _29_v13 _dafny.Map
			_ = _29_v13
			_29_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _30_v14 _dafny.Set
			_ = _30_v14
			_30_v14 = _dafny.SetOf(_28_v12, _28_v12, Companion_D0_.Create_DC2_((_20_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_20_v5), 0))).Int()).(_dafny.Int)), Companion_D0_.Create_DC2_((_29_v13).Cardinality()), _28_v12)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_19_v4), 0))
			_ = _index9
			var _rhs4 _dafny.Int = Companion_Default___.SafeModuloInt((_27_v11).Cardinality(), ((_30_v14).Cardinality()).Plus(p0))
			_ = _rhs4
			var _rhs5 bool = p2
			_ = _rhs5
			var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_22_v6, _22_v6)
			_ = _rhs6
			var _lhs6 *GlobalState = globalState
			_ = _lhs6
			var _lhs7 _dafny.Array = _19_v4
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_19_v4), 0))
			_ = _lhs8
			var _lhs9 *GlobalState = globalState
			_ = _lhs9
			_lhs6.F7 = _rhs4
			(_lhs7).ArraySet1(_rhs5, (_lhs8).Int())
			_lhs9.F4 = _rhs6
			var _31_v15 _dafny.Map
			_ = _31_v15
			_31_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_19_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_19_v4), 0))).Int()).(bool)), _dafny.ArrayCastTo((_18_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_18_v3), 0))).Int())))
			var _32_v16 _dafny.Map
			_ = _32_v16
			_32_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_31_v15).Merge(_31_v15), Companion_Default___.SafeModuloInt((_20_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_20_v5), 0))).Int()).(_dafny.Int), p3))
			_32_v16 = (_32_v16).Update((_31_v15).Update(false, _dafny.ArrayCastTo((_18_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_18_v3), 0))).Int()))), p3)
		}
		var _33_v17 D0
		_ = _33_v17
		_33_v17 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(993))
		var _source1 D0 = func(_pat_let0_0 D0) D0 {
			return func(_34_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let1_0 _dafny.Int) D0 {
					return func(_36_dt__update_hcf4_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC2_(_36_dt__update_hcf4_h0)
					}(_pat_let1_0)
				}((func() _dafny.Map {
					var _coll2 = _dafny.NewMapBuilder()
					_ = _coll2
					for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-54), _dafny.IntOfInt64(138))); ; {
						_compr_2, _ok2 := _iter2()
						if !_ok2 {
							break
						}
						var _35_v18 _dafny.Int
						_35_v18 = interface{}(_compr_2).(_dafny.Int)
						if ((_dafny.IntOfInt64(-54)).Cmp(_35_v18) <= 0) && ((_35_v18).Cmp(_dafny.IntOfInt64(138)) < 0) {
							_coll2.Add((_35_v18).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fqe")).Cardinality())), _dafny.IntOfInt64(776))
						}
					}
					return _coll2.ToMap()
				}()).Cardinality())
			}(_pat_let0_0)
		}(_33_v17)
		_ = _source1
		if _source1.Is_DC1() {
			var _37___mcc_h2 _dafny.Int = _source1.Get_().(D0_DC1).Cf1
			_ = _37___mcc_h2
			var _38___mcc_h3 _dafny.Int = _source1.Get_().(D0_DC1).Cf2
			_ = _38___mcc_h3
			var _39___mcc_h4 _dafny.Int = _source1.Get_().(D0_DC1).Cf3
			_ = _39___mcc_h4
			var _40_cf3 _dafny.Int = _39___mcc_h4
			_ = _40_cf3
			var _41_cf2 _dafny.Int = _38___mcc_h3
			_ = _41_cf2
			var _42_cf1 _dafny.Int = _37___mcc_h2
			_ = _42_cf1
			var _43_v19 _dafny.Array
			_ = _43_v19
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw5
			_43_v19 = _nw5
			var _44_v20 *C0
			_ = _44_v20
			var _nw6 *C0 = New_C0_()
			_ = _nw6
			_nw6.Ctor__(p3)
			_44_v20 = _nw6
			var _45_v21 _dafny.Sequence
			_ = _45_v21
			_45_v21 = _dafny.SeqOf(_44_v20, _44_v20)
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_43_v19), 0))
			_ = _index10
			(_43_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_45_v21, _45_v21), (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_43_v19), 0))
			_ = _index11
			(_43_v19).ArraySet1(_dafny.SeqOf(_44_v20), (_index11).Int())
			_40_cf3 = Companion_Default___.SafeModuloInt(p0, _40_cf3)
			_40_cf3 = Companion_Default___.SafeModuloInt(_42_cf1, _10_i0)
			var _46_v22 _dafny.Sequence
			_ = _46_v22
			_46_v22 = _dafny.UnicodeSeqOfUtf8Bytes("s")
			var _47_v23 _dafny.CodePoint
			_ = _47_v23
			_47_v23 = _dafny.CodePoint('u')
			(globalState).F4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_46_v22, _46_v22), (Companion_Default___.SafeIndex(_42_cf1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_46_v22, _46_v22)).Cardinality()))).Uint32(), _47_v23), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ihfuqk"), _dafny.UnicodeSeqOfUtf8Bytes("unigk")))
		} else if _source1.Is_DC2() {
			var _48___mcc_h5 _dafny.Int = _source1.Get_().(D0_DC2).Cf4
			_ = _48___mcc_h5
			var _49_cf4 _dafny.Int = _48___mcc_h5
			_ = _49_cf4
			var _50_v24 _dafny.Array
			_ = _50_v24
			var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw7
			_50_v24 = _nw7
			var _51_v25 _dafny.Sequence
			_ = _51_v25
			_51_v25 = _dafny.UnicodeSeqOfUtf8Bytes("vp")
			var _52_v26 _dafny.Sequence
			_ = _52_v26
			_52_v26 = _dafny.SeqOf(_51_v25)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(444), _dafny.ArrayLen((_50_v24), 0))
			_ = _index12
			(_50_v24).ArraySet1(_52_v26, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(444), _dafny.ArrayLen((_50_v24), 0))
			_ = _index13
			(_50_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}((func(_53_v25 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_54_i2 _dafny.Int) _dafny.Sequence {
					return _53_v25
				}
			})(_51_v25))), _dafny.Companion_Sequence_.Concatenate(_52_v26, _52_v26)), (_index13).Int())
			var _55_v27 _dafny.Array
			_ = _55_v27
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_1
			var _nw8 _dafny.Array
			_ = _nw8
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw8 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_56_cf4 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_57_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_57_i3, _56_cf4)
					}
				})(_49_cf4)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw8 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw8).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw8).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_55_v27 = _nw8
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_55_v27), 0))
			_ = _index14
			(_55_v27).ArraySet1(p0, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_55_v27), 0))
			_ = _index15
			(_55_v27).ArraySet1(_10_i0, (_index15).Int())
			var _58_v28 bool
			_ = _58_v28
			_58_v28 = false
			_58_v28 = _58_v28
			_58_v28 = _58_v28
		} else if _source1.Is_DC3() {
			var _59___mcc_h6 _dafny.Map = _source1.Get_().(D0_DC3).Cf5
			_ = _59___mcc_h6
			var _60___mcc_h7 bool = _source1.Get_().(D0_DC3).Cf6
			_ = _60___mcc_h7
			var _61___mcc_h8 _dafny.Int = _source1.Get_().(D0_DC3).Cf7
			_ = _61___mcc_h8
			var _62_cf7 _dafny.Int = _61___mcc_h8
			_ = _62_cf7
			var _63_cf6 bool = _60___mcc_h7
			_ = _63_cf6
			var _64_cf5 _dafny.Map = _59___mcc_h6
			_ = _64_cf5
			(globalState).F7 = p3
			_63_cf6 = _63_cf6
			var _65_v29 *C0
			_ = _65_v29
			var _nw9 *C0 = New_C0_()
			_ = _nw9
			_nw9.Ctor__(_dafny.IntOfInt64(-122))
			_65_v29 = _nw9
			(_65_v29).F8 = _62_cf7
		} else {
			var _66___mcc_h9 _dafny.Sequence = _source1.Get_().(D0_DC0).Cf0
			_ = _66___mcc_h9
			var _67_cf0 _dafny.Sequence = _66___mcc_h9
			_ = _67_cf0
			var _68_v30 *C0
			_ = _68_v30
			var _nw10 *C0 = New_C0_()
			_ = _nw10
			_nw10.Ctor__(p0)
			_68_v30 = _nw10
			var _69_v31 _dafny.Array
			_ = _69_v31
			var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw11
			_69_v31 = _nw11
			var _70_v32 _dafny.Sequence
			_ = _70_v32
			_70_v32 = _dafny.SeqOf(_69_v31, _69_v31, _69_v31)
			_70_v32 = _70_v32
			var _71_v33 _dafny.Map
			_ = _71_v33
			_71_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_68_v30.F8, (p3).Cmp(_10_i0) != 0)
			_71_v33 = (_71_v33).Update(p3, p2)
			var _72_v34 bool
			_ = _72_v34
			_72_v34 = true
			_72_v34 = _72_v34
		}
		var _73_v35 bool
		_ = _73_v35
		_73_v35 = false
		var _74_v36 _dafny.Sequence
		_ = _74_v36
		_74_v36 = _dafny.UnicodeSeqOfUtf8Bytes("c")
		_73_v35 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_74_v36, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(180))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_75_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})))).Cardinality())).Cmp(p0) <= 0
		var _76_v37 _dafny.Array
		_ = _76_v37
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_2
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = func(_77_i5 _dafny.Int) _dafny.Int {
				return (_77_i5).Minus(_dafny.IntOfInt64(-109))
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw12 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw12).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw12).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_76_v37 = _nw12
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_76_v37), 0))
		_ = _index16
		(_76_v37).ArraySet1((p3).Minus(p3), (_index16).Int())
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_76_v37), 0))
		_ = _index17
		(_76_v37).ArraySet1(p0, (_index17).Int())
	}
	var _78_i6 _dafny.Int
	_ = _78_i6
	_78_i6 = _dafny.Zero
	{
		for p2 {
			{
				if (_78_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_78_i6 = (_78_i6).Plus(_dafny.One)
				var _79_v38 _dafny.CodePoint
				_ = _79_v38
				_79_v38 = _dafny.CodePoint('k')
				var _80_v39 _dafny.Map
				_ = _80_v39
				_80_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_79_v38, p2)
				var _81_v40 _dafny.MultiSet
				_ = _81_v40
				_81_v40 = _dafny.MultiSetOf(_80_v39, _80_v39)
				var _82_v41 D0
				_ = _82_v41
				_82_v41 = Companion_D0_.Create_DC3_(Companion_Default___.Fm11(globalState), Companion_Default___.Fm0(!(p2), p2, _81_v40, p0, globalState), p0)
				(globalState).F7 = (_82_v41).Dtor_cf7()
				var _83_v42 *C0
				_ = _83_v42
				var _nw13 *C0 = New_C0_()
				_ = _nw13
				_nw13.Ctor__(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg10 _dafny.Int) interface{} {
						return coer10(arg10)
					}
				}((func(_84_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_85_i7 _dafny.Int) _dafny.CodePoint {
						return _84_v38
					}
				})(_79_v38)))).Cardinality()))
				_83_v42 = _nw13
				var _86_v43 D4
				_ = _86_v43
				_86_v43 = Companion_D4_.Create_DC10_(p2, _83_v42)
				var _pat_let_tv0 = p2
				_ = _pat_let_tv0
				var _pat_let_tv1 = _83_v42
				_ = _pat_let_tv1
				var _87_v44 _dafny.Array
				_ = _87_v44
				var _nwElement0_0 D4 = _86_v43
				_ = _nwElement0_0
				var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(29))
				_ = _nw14
				(_nw14).ArraySet1(_nwElement0_0, 0)
				(_nw14).ArraySet1(func(_pat_let2_0 D4) D4 {
					return func(_88_dt__update__tmp_h1 D4) D4 {
						return func(_pat_let3_0 bool) D4 {
							return func(_89_dt__update_hcf15_h0 bool) D4 {
								return Companion_D4_.Create_DC10_(_89_dt__update_hcf15_h0, (_88_dt__update__tmp_h1).Dtor_cf16())
							}(_pat_let3_0)
						}(_pat_let_tv0)
					}(_pat_let2_0)
				}(_86_v43), 1)
				(_nw14).ArraySet1(_86_v43, 2)
				(_nw14).ArraySet1(_86_v43, 3)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(p2, _83_v42), 4)
				(_nw14).ArraySet1(_86_v43, 5)
				(_nw14).ArraySet1(func(_pat_let4_0 D4) D4 {
					return func(_90_dt__update__tmp_h2 D4) D4 {
						return func(_pat_let5_0 *C0) D4 {
							return func(_91_dt__update_hcf16_h0 *C0) D4 {
								return Companion_D4_.Create_DC10_((_90_dt__update__tmp_h2).Dtor_cf15(), _91_dt__update_hcf16_h0)
							}(_pat_let5_0)
						}(_pat_let_tv1)
					}(_pat_let4_0)
				}(_86_v43), 6)
				(_nw14).ArraySet1(_86_v43, 7)
				(_nw14).ArraySet1(_86_v43, 8)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(p2, _83_v42), 9)
				(_nw14).ArraySet1(_86_v43, 10)
				(_nw14).ArraySet1(_86_v43, 11)
				(_nw14).ArraySet1(_86_v43, 12)
				(_nw14).ArraySet1(_86_v43, 13)
				(_nw14).ArraySet1(_86_v43, 14)
				(_nw14).ArraySet1(_86_v43, 15)
				(_nw14).ArraySet1(_86_v43, 16)
				(_nw14).ArraySet1(_86_v43, 17)
				(_nw14).ArraySet1(_86_v43, 18)
				(_nw14).ArraySet1(_86_v43, 19)
				(_nw14).ArraySet1(_86_v43, 20)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(p2, _83_v42), 21)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(p2, _83_v42), 22)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(!(p2), _83_v42), 23)
				(_nw14).ArraySet1(_86_v43, 24)
				(_nw14).ArraySet1(_86_v43, 25)
				(_nw14).ArraySet1(Companion_D4_.Create_DC10_(p2, _83_v42), 26)
				(_nw14).ArraySet1(_86_v43, 27)
				(_nw14).ArraySet1(_86_v43, 28)
				_87_v44 = _nw14
				var _92_v45 _dafny.Array
				_ = _92_v45
				var _nwElement0_1 _dafny.Array = _87_v44
				_ = _nwElement0_1
				var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(17))
				_ = _nw15
				(_nw15).ArraySet1(_nwElement0_1, 0)
				(_nw15).ArraySet1(_87_v44, 1)
				(_nw15).ArraySet1(_87_v44, 2)
				(_nw15).ArraySet1(_87_v44, 3)
				(_nw15).ArraySet1(_87_v44, 4)
				(_nw15).ArraySet1(_87_v44, 5)
				(_nw15).ArraySet1(_87_v44, 6)
				(_nw15).ArraySet1(_87_v44, 7)
				(_nw15).ArraySet1(_87_v44, 8)
				(_nw15).ArraySet1(_87_v44, 9)
				(_nw15).ArraySet1(_87_v44, 10)
				(_nw15).ArraySet1(_87_v44, 11)
				(_nw15).ArraySet1(_87_v44, 12)
				(_nw15).ArraySet1(_87_v44, 13)
				(_nw15).ArraySet1(_87_v44, 14)
				(_nw15).ArraySet1((func() _dafny.Array {
					if p2 {
						return _87_v44
					}
					return _87_v44
				})(), 15)
				(_nw15).ArraySet1(_87_v44, 16)
				_92_v45 = _nw15
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_92_v45), 0))
				_ = _index18
				(_92_v45).ArraySet1(_87_v44, (_index18).Int())
				var _pat_let_tv2 = p2
				_ = _pat_let_tv2
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_92_v45), 0))
				_ = _index19
				var _nwElement0_2 D4 = func(_pat_let6_0 D4) D4 {
					return func(_93_dt__update__tmp_h3 D4) D4 {
						return func(_pat_let7_0 bool) D4 {
							return func(_94_dt__update_hcf15_h1 bool) D4 {
								return Companion_D4_.Create_DC10_(_94_dt__update_hcf15_h1, (_93_dt__update__tmp_h3).Dtor_cf16())
							}(_pat_let7_0)
						}(_pat_let_tv2)
					}(_pat_let6_0)
				}(_86_v43)
				_ = _nwElement0_2
				var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
				_ = _nw16
				(_nw16).ArraySet1(_nwElement0_2, 0)
				(_nw16).ArraySet1(_86_v43, 1)
				(_92_v45).ArraySet1(_nw16, (_index19).Int())
				var _95_v46 *C0
				_ = _95_v46
				var _nw17 *C0 = New_C0_()
				_ = _nw17
				_nw17.Ctor__(p0)
				_95_v46 = _nw17
				var _96_v47 _dafny.Sequence
				_ = _96_v47
				_96_v47 = _dafny.UnicodeSeqOfUtf8Bytes("ogpsjgmip")
				var _97_v48 _dafny.Map
				_ = _97_v48
				_97_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v42, _96_v47)
				_97_v48 = (_97_v48).Update(_95_v46, _96_v47)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _98_v49 *C0
	_ = _98_v49
	var _nw18 *C0 = New_C0_()
	_ = _nw18
	_nw18.Ctor__(p3)
	_98_v49 = _nw18
	var _99_v50 _dafny.CodePoint
	_ = _99_v50
	_99_v50 = _dafny.CodePoint('s')
	var _100_v51 _dafny.Map
	_ = _100_v51
	_100_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v49, _99_v50)
	var _101_v53 _dafny.Array
	_ = _101_v53
	var _len0_3 _dafny.Int = _dafny.IntOfInt64(13)
	_ = _len0_3
	var _nw19 _dafny.Array
	_ = _nw19
	if _len0_3.Cmp(_dafny.Zero) == 0 {
		_nw19 = _dafny.NewArray(_len0_3)
	} else {
		var _init3 func(_dafny.Int) _dafny.Int = (func(_102_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_103_i8 _dafny.Int) _dafny.Int {
				return (_103_i8).Times((_dafny.Zero).Minus((func() _dafny.Set {
					var _coll3 = _dafny.NewBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate((_dafny.SetOf(_102_p0)).Elements()); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _104_v52 _dafny.Int
						_104_v52 = interface{}(_compr_3).(_dafny.Int)
						if (_dafny.SetOf(_102_p0)).Contains(_104_v52) {
							_coll3.Add((_104_v52).Times(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
						}
					}
					return _coll3.ToSet()
				}()).Cardinality()))
			}
		})(p0)
		_ = _init3
		var _element0_3 = _init3(_dafny.Zero)
		_ = _element0_3
		_nw19 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
		(_nw19).ArraySet1(_element0_3, 0)
		var _nativeLen0_3 = (_len0_3).Int()
		_ = _nativeLen0_3
		for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
			(_nw19).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
		}
	}
	_101_v53 = _nw19
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_101_v53), 0))
	_ = _index20
	(_101_v53).ArraySet1((_98_v49.F8).Minus((_dafny.Zero).Minus(p3)), (_index20).Int())
	var _105_v54 _dafny.Map
	_ = _105_v54
	_105_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), p2)
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_101_v53), 0))
	_ = _index21
	var _rhs7 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v49, _99_v50)
	_ = _rhs7
	var _rhs8 _dafny.Int = p0
	_ = _rhs8
	var _rhs9 _dafny.Int = ((_105_v54).Cardinality()).Times((p3).Minus(_98_v49.F8))
	_ = _rhs9
	var _rhs10 *C0 = _98_v49
	_ = _rhs10
	var _lhs10 _dafny.Array = _101_v53
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_101_v53), 0))
	_ = _lhs11
	var _lhs12 *GlobalState = globalState
	_ = _lhs12
	_100_v51 = _rhs7
	(_lhs10).ArraySet1(_rhs8, (_lhs11).Int())
	_lhs12.F7 = _rhs9
	_98_v49 = _rhs10
	var _106_v55 _dafny.Array
	_ = _106_v55
	var _nw20 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
	_ = _nw20
	_106_v55 = _nw20
	_106_v55 = _106_v55
	var _107_v56 _dafny.Sequence
	_ = _107_v56
	_107_v56 = _dafny.UnicodeSeqOfUtf8Bytes("skwswvs")
	var _108_v57 _dafny.Map
	_ = _108_v57
	_108_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(973), _107_v56)
	var _109_v58 _dafny.Map
	_ = _109_v58
	_109_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_107_v56, (func() _dafny.Sequence {
		if (_108_v57).Contains(_98_v49.F8) {
			return (_108_v57).Get(_98_v49.F8).(_dafny.Sequence)
		}
		return _dafny.UnicodeSeqOfUtf8Bytes("kefocyf")
	})())).Cardinality()))
	var _110_v59 _dafny.Map
	_ = _110_v59
	_110_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v50, p2)
	var _111_v60 _dafny.MultiSet
	_ = _111_v60
	_111_v60 = _dafny.MultiSetOf(_110_v59)
	var _112_v61 _dafny.Set
	_ = _112_v61
	_112_v61 = _dafny.SetOf(p2, p2, p2, p2, p2)
	var _113_v62 _dafny.Sequence
	_ = _113_v62
	_113_v62 = _dafny.SeqOf(_dafny.SetOf(Companion_Default___.Fm0(p2, p2, _111_v60, p0, globalState)), _112_v61)
	var _114_v63 _dafny.Sequence
	_ = _114_v63
	_114_v63 = _dafny.SeqOf(_dafny.IntOfInt64(-8), _98_v49.F8)
	_109_v58 = (_109_v58).Update(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_112_v61), (_113_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_114_v63).Cardinality()), _dafny.IntOfUint32((_113_v62).Cardinality()))).Uint32()).(_dafny.Set)), (_101_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_101_v53), 0))).Int()).(_dafny.Int))
	for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_106_v55), 0))); ; {
		_guard_loop_0, _ok4 := _iter4()
		if !_ok4 {
			break
		}
		var _115_i9 _dafny.Int
		_115_i9 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_115_i9).Sign() != -1) && ((_115_i9).Cmp(_dafny.ArrayLen((_106_v55), 0)) < 0)) {
			(_106_v55).ArraySet1((_112_v61).IsSubsetOf(_112_v61), (_115_i9).Int())
		}
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _116_v0 bool
	_ = _116_v0
	_116_v0 = false
	var _117_v1 _dafny.Map
	_ = _117_v1
	_117_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gmoircrjr")).Cardinality()))
	var _118_v2 _dafny.Array
	_ = _118_v2
	var _len0_4 _dafny.Int = _dafny.One
	_ = _len0_4
	var _nw21 _dafny.Array
	_ = _nw21
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw21 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = func(_119_i0 _dafny.Int) _dafny.Int {
			return Companion_Default___.SafeModuloInt(_119_i0, _dafny.IntOfInt64(448))
		}
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw21 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw21).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw21).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_118_v2 = _nw21
	var _120_v3 _dafny.Array
	_ = _120_v3
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_5
	var _nw22 _dafny.Array
	_ = _nw22
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw22 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) bool = func(_121_i1 _dafny.Int) bool {
			return false
		}
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw22 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw22).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw22).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_120_v3 = _nw22
	var _122_v4 _dafny.Sequence
	_ = _122_v4
	_122_v4 = _dafny.UnicodeSeqOfUtf8Bytes("pnncmt")
	var _123_globalState *GlobalState
	_ = _123_globalState
	var _nw23 *GlobalState = New_GlobalState_()
	_ = _nw23
	_nw23.Ctor__(_117_v1, _118_v2, _120_v3, _dafny.IntOfInt64(287), _122_v4, false, _122_v4, _dafny.IntOfInt64(632))
	_123_globalState = _nw23
	var _124_v5 _dafny.Int
	_ = _124_v5
	_124_v5 = _dafny.IntOfInt64(681)
	var _125_v6 _dafny.Map
	_ = _125_v6
	_125_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_116_v0, _124_v5, _124_v5, _123_globalState), _116_v0)
	var _126_v7 _dafny.Sequence
	_ = _126_v7
	_126_v7 = _dafny.SeqOf(_125_v6)
	var _127_v8 _dafny.MultiSet
	_ = _127_v8
	_127_v8 = _dafny.MultiSetOf(_125_v6, _125_v6, _125_v6)
	_116_v0 = Companion_Default___.Fm0(Companion_Default___.Fm0(_116_v0, _116_v0, _dafny.MultiSetFromSeq(_126_v7), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lcjlcyyhn")).Cardinality()), _123_globalState), Companion_Default___.Fm0(_116_v0, _116_v0, _dafny.MultiSetOf(_125_v6, _125_v6, Companion_Default___.Fm2(!(_116_v0), _123_globalState)), _124_v5, _123_globalState), (_127_v8).Difference(_dafny.MultiSetFromSeq(_126_v7)), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_122_v4).Cardinality()))).Times(_124_v5), _123_globalState)
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
	_ = _index22
	(_120_v3).ArraySet1(true, (_index22).Int())
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
	_ = _index23
	(_120_v3).ArraySet1((_124_v5).Cmp(_124_v5) >= 0, (_index23).Int())
	var _128_v9 _dafny.Sequence
	_ = _128_v9
	_128_v9 = _dafny.SeqOf((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _116_v0)
	var _129_v10 _dafny.Map
	_ = _129_v10
	_129_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v0, _118_v2)
	var _130_i2 _dafny.Int
	_ = _130_i2
	_130_i2 = _dafny.Zero
	{
		for (_128_v9).Select((Companion_Default___.SafeIndex(((_129_v10).Merge(_129_v10)).Cardinality(), _dafny.IntOfUint32((_128_v9).Cardinality()))).Uint32()).(bool) {
			{
				if (_130_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_130_i2 = (_130_i2).Plus(_dafny.One)
				var _131_v11 _dafny.Sequence
				_ = _131_v11
				_131_v11 = _dafny.SeqOf(_124_v5)
				var _132_v12 _dafny.CodePoint
				_ = _132_v12
				_132_v12 = _dafny.CodePoint('i')
				var _133_v13 _dafny.Map
				_ = _133_v13
				_133_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v4, _132_v12)
				var _134_v14 _dafny.MultiSet
				_ = _134_v14
				_134_v14 = _dafny.MultiSetOf((_131_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_131_v11).Cardinality()), _dafny.IntOfUint32((_131_v11).Cardinality()))).Uint32()).(_dafny.Int), _124_v5, (_dafny.Zero).Minus((_133_v13).Cardinality()), _124_v5)
				Companion_Default___.M0((_dafny.Zero).Minus(_124_v5), _134_v14, ((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)) == ((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)), (_dafny.Zero).Minus(_124_v5), _123_globalState)
				var _135_v15 _dafny.Map
				_ = _135_v15
				_135_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_122_v4, _122_v4), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-825), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_122_v4, _122_v4)).Cardinality()))).Uint32(), _132_v12), (func() bool {
					if (_128_v9).Select((Companion_Default___.SafeIndex(_124_v5, _dafny.IntOfUint32((_128_v9).Cardinality()))).Uint32()).(bool) {
						return !(_116_v0)
					}
					return (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)
				})())
				var _136_v17 _dafny.MultiSet
				_ = _136_v17
				_136_v17 = _dafny.MultiSetOf(_122_v4, _dafny.UnicodeSeqOfUtf8Bytes("n"), _dafny.UnicodeSeqOfUtf8Bytes("npfn"))
				var _rhs11 bool = (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)
				_ = _rhs11
				var _rhs12 _dafny.Int = (_dafny.Zero).Minus((_124_v5).Minus(_124_v5))
				_ = _rhs12
				var _rhs13 _dafny.Map = func() _dafny.Map {
					var _coll4 = _dafny.NewMapBuilder()
					_ = _coll4
					for _iter5 := _dafny.Iterate((_136_v17).Elements()); ; {
						_compr_4, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _137_v16 _dafny.Sequence
						_137_v16 = interface{}(_compr_4).(_dafny.Sequence)
						if (_136_v17).Contains(_137_v16) {
							_coll4.Add(_137_v16, _116_v0)
						}
					}
					return _coll4.ToMap()
				}()
				_ = _rhs13
				var _lhs13 *GlobalState = _123_globalState
				_ = _lhs13
				_116_v0 = _rhs11
				_lhs13.F7 = _rhs12
				_135_v15 = _rhs13
				if !(!((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)) || (_116_v0)) || ((_134_v14).Contains(_124_v5)) {
					var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_118_v2), 0))
					_ = _index24
					(_118_v2).ArraySet1(_124_v5, (_index24).Int())
					var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_118_v2), 0))
					_ = _index25
					(_118_v2).ArraySet1(_124_v5, (_index25).Int())
					var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_118_v2), 0))
					_ = _index26
					(_118_v2).ArraySet1(_124_v5, (_index26).Int())
					(_123_globalState).F7 = ((_118_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_118_v2), 0))).Int()).(_dafny.Int)).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(_123_globalState), (_dafny.Zero).Minus(_124_v5)))
					var _138_v18 _dafny.Array
					_ = _138_v18
					var _len0_6 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_6
					var _nw24 _dafny.Array
					_ = _nw24
					if _len0_6.Cmp(_dafny.Zero) == 0 {
						_nw24 = _dafny.NewArray(_len0_6)
					} else {
						var _init6 func(_dafny.Int) _dafny.Sequence = (func(_139_v0 bool) func(_dafny.Int) _dafny.Sequence {
							return func(_140_i3 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(_139_v0)
							}
						})(_116_v0)
						_ = _init6
						var _element0_6 = _init6(_dafny.Zero)
						_ = _element0_6
						_nw24 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
						(_nw24).ArraySet1(_element0_6, 0)
						var _nativeLen0_6 = (_len0_6).Int()
						_ = _nativeLen0_6
						for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
							(_nw24).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
						}
					}
					_138_v18 = _nw24
					var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v18), 0))
					_ = _index27
					(_138_v18).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_128_v9, _128_v9), (_index27).Int())
					var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v18), 0))
					_ = _index28
					(_138_v18).ArraySet1(_128_v9, (_index28).Int())
					_116_v0 = ((_138_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v18), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf((_138_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v18), 0))).Int()).(_dafny.Sequence))).Cardinality(), _dafny.IntOfUint32(((_138_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_138_v18), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(bool)
				} else {
					var _141_v19 *C0
					_ = _141_v19
					var _nw25 *C0 = New_C0_()
					_ = _nw25
					_nw25.Ctor__(_124_v5)
					_141_v19 = _nw25
					var _142_v20 bool
					_ = _142_v20
					var _143_v21 bool
					_ = _143_v21
					var _144_v22 _dafny.Int
					_ = _144_v22
					var _145_v23 _dafny.CodePoint
					_ = _145_v23
					var _out0 bool
					_ = _out0
					var _out1 bool
					_ = _out1
					var _out2 _dafny.Int
					_ = _out2
					var _out3 _dafny.CodePoint
					_ = _out3
					_out0, _out1, _out2, _out3 = (_141_v19).M1(_123_globalState)
					_142_v20 = _out0
					_143_v21 = _out1
					_144_v22 = _out2
					_145_v23 = _out3
					_131_v11 = _dafny.Companion_Sequence_.Concatenate(_131_v11, _131_v11)
					var _146_v24 _dafny.Map
					_ = _146_v24
					_146_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v19, _141_v19)
					var _147_v25 _dafny.Map
					_ = _147_v25
					_147_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(587), _146_v24)
					var _148_v26 _dafny.MultiSet
					_ = _148_v26
					_148_v26 = _dafny.MultiSetOf(_141_v19, _141_v19)
					Companion_Default___.M0(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_122_v4, _122_v4)).Cardinality()), _134_v14, (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_122_v4, (Companion_Default___.SafeIndex(_141_v19.F8, _dafny.IntOfUint32((_122_v4).Cardinality()))).Uint32(), _dafny.CodePoint('c'))).Cardinality())).Cmp(((func() _dafny.Map {
						if (_147_v25).Contains(_141_v19.F8) {
							return (_147_v25).Get(_141_v19.F8).(_dafny.Map)
						}
						return _146_v24
					})()).Cardinality()) != 0, (_148_v26).Cardinality(), _123_globalState)
					var _149_v27 D2
					_ = _149_v27
					_149_v27 = Companion_D2_.Create_DC5_(_128_v9)
					var _150_v28 _dafny.Set
					_ = _150_v28
					_150_v28 = _dafny.SetOf(_143_v21, !(_143_v21) || (false), (_124_v5).Cmp((func() _dafny.Int {
						if (_117_v1).Contains(_116_v0) {
							return (_117_v1).Get(_116_v0).(_dafny.Int)
						}
						return _dafny.IntOfUint32(((_149_v27).Dtor_cf9()).Cardinality())
					})()) != 0)
					_150_v28 = _150_v28
				}
				Companion_Default___.M0(_124_v5, _134_v14, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _124_v5, _123_globalState)
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _151_v29 _dafny.Array
	_ = _151_v29
	var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
	_ = _nw26
	_151_v29 = _nw26
	var _152_v30 _dafny.Set
	_ = _152_v30
	_152_v30 = _dafny.SetOf(_151_v29, _151_v29, _151_v29, _151_v29, _151_v29)
	var _153_v31 _dafny.MultiSet
	_ = _153_v31
	_153_v31 = _dafny.MultiSetOf(_124_v5)
	var _154_v32 D0
	_ = _154_v32
	_154_v32 = Companion_D0_.Create_DC3_(_117_v1, _116_v0, _124_v5)
	Companion_Default___.M0((_152_v30).Cardinality(), _153_v31, (_154_v32).Dtor_cf6(), _124_v5, _123_globalState)
	var _155_v33 _dafny.Map
	_ = _155_v33
	_155_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v5, _118_v2)
	_155_v33 = (_155_v33).Update(_124_v5, _118_v2)
	var _156_v34 D0
	_ = _156_v34
	_156_v34 = Companion_D0_.Create_DC1_(_124_v5, (_124_v5).Minus(_124_v5), _124_v5)
	_156_v34 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(322), _124_v5, _124_v5)
	var _157_v35 D0
	_ = _157_v35
	_157_v35 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-986))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}((func(_158_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_159_i4 _dafny.Int) _dafny.Int {
			return _158_v5
		}
	})(_124_v5)))).Cardinality()))
	var _160_v36 _dafny.MultiSet
	_ = _160_v36
	_160_v36 = _dafny.MultiSetOf(_157_v35, Companion_D0_.Create_DC2_(_124_v5))
	_160_v36 = (_160_v36).Update(_157_v35, Companion_Default___.Abs(_124_v5))
	var _161_v37 _dafny.Map
	_ = _161_v37
	_161_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v0, _116_v0)
	var _rhs14 _dafny.Int = (func() _dafny.Int {
		if false {
			return ((_161_v37).Merge(_161_v37)).Cardinality()
		}
		return _124_v5
	})()
	_ = _rhs14
	var _rhs15 _dafny.Int = _124_v5
	_ = _rhs15
	var _lhs14 *GlobalState = _123_globalState
	_ = _lhs14
	_124_v5 = _rhs14
	_lhs14.F7 = _rhs15
	var _162_v38 _dafny.Map
	_ = _162_v38
	_162_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v2, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
	var _163_v39 _dafny.CodePoint
	_ = _163_v39
	_163_v39 = _dafny.CodePoint('c')
	if (func() bool {
		if (_162_v38).Contains(_118_v2) {
			return (_162_v38).Get(_118_v2).(bool)
		}
		return Companion_Default___.Fm0(!(_116_v0), Companion_Default___.Fm0(false, false, _dafny.MultiSetOf((_125_v6).Update(_163_v39, _116_v0)), _124_v5, _123_globalState), (_dafny.MultiSetFromSeq(_126_v7)).Update(_125_v6, Companion_Default___.Abs(_124_v5)), _124_v5, _123_globalState)
	})() {
		var _164_v40 _dafny.Map
		_ = _164_v40
		_164_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v0, _120_v3)
		_164_v40 = (_164_v40).Update(_116_v0, _120_v3)
		var _165_v41 D3
		_ = _165_v41
		_165_v41 = Companion_D3_.Create_DC7_(_161_v37)
		_116_v0 = (((_165_v41).Dtor_cf11()).Merge(_161_v37)).Equals((_161_v37).Merge(Companion_Default___.Fm4(!(_116_v0), _124_v5, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _123_globalState)))
		var _166_v42 _dafny.MultiSet
		_ = _166_v42
		_166_v42 = _dafny.MultiSetOf((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _116_v0, !(!(!(_116_v0))) || ((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)), _116_v0, _dafny.Companion_Sequence_.IsProperPrefixOf(_122_v4, _122_v4))
		var _167_v43 _dafny.Set
		_ = _167_v43
		_167_v43 = _dafny.SetOf(!(_116_v0))
		var _168_v44 _dafny.Map
		_ = _168_v44
		_168_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(604))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}((func(_169_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_170_i5 _dafny.Int) _dafny.Int {
				return _169_v5
			}
		})(_124_v5)))).Cardinality()), !(_116_v0))
		(_123_globalState).F7 = (func() _dafny.Int {
			if (_166_v42).Contains(Companion_Default___.Fm0(_116_v0, _116_v0, _127_v8, (_167_v43).Cardinality(), _123_globalState)) {
				return (_166_v42).Multiplicity(Companion_Default___.Fm0(_116_v0, _116_v0, _127_v8, (_167_v43).Cardinality(), _123_globalState))
			}
			return Companion_Default___.SafeDivisionInt((Companion_Default___.Fm5(_dafny.IntOfInt64(17), _116_v0, (func() bool {
				if (_168_v44).Contains((_153_v31).Cardinality()) {
					return (_168_v44).Get((_153_v31).Cardinality()).(bool)
				}
				return true
			})(), _123_globalState)).Cardinality(), _dafny.IntOfInt64(136))
		})()
		var _171_v46 _dafny.Map
		_ = _171_v46
		_171_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v5, Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_172_i7 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hwbs")).Cardinality())
		})))))
		var _173_v47 _dafny.Sequence
		_ = _173_v47
		_173_v47 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_124_v5, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))).Cardinality())
		var _174_v48 _dafny.Sequence
		_ = _174_v48
		_174_v48 = _dafny.SeqOf(_173_v47)
		var _175_v49 D0
		_ = _175_v49
		_175_v49 = Companion_D0_.Create_DC0_(_174_v48)
		Companion_Default___.M0(((func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(954), _dafny.IntOfInt64(623))); ; {
				_compr_5, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _176_v45 _dafny.Int
				_176_v45 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(954)).Cmp(_176_v45) <= 0) && ((_176_v45).Cmp(_dafny.IntOfInt64(623)) < 0) {
					_coll5.Add((_176_v45).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_124_v5)).Cardinality())), Companion_D0_.Create_DC0_(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-263))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg14 _dafny.Int) interface{} {
							return coer14(arg14)
						}
					}((func(_177_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_178_i6 _dafny.Int) _dafny.Int {
							return _177_v5
						}
					})(_124_v5))), _dafny.SeqOf(_124_v5, _dafny.IntOfUint32((_122_v4).Cardinality())))))
				}
			}
			return _coll5.ToMap()
		}()).Merge((_171_v46).Update(_dafny.IntOfUint32((_122_v4).Cardinality()), _175_v49))).Cardinality(), _153_v31, _116_v0, _124_v5, _123_globalState)
		_124_v5 = _124_v5
	} else {
		var _179_v50 *C0
		_ = _179_v50
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__(Companion_Default___.SafeModuloInt(_124_v5, _dafny.IntOfInt64(-119)))
		_179_v50 = _nw27
		var _180_v51 _dafny.Set
		_ = _180_v51
		_180_v51 = _dafny.SetOf(_161_v37)
		var _181_v52 _dafny.Set
		_ = _181_v52
		_181_v52 = _180_v51
		var _182_v53 _dafny.Map
		_ = _182_v53
		_182_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-303), _181_v52)
		var _183_v54 _dafny.Sequence
		_ = _183_v54
		_183_v54 = _dafny.SeqOf(_124_v5, _dafny.IntOfInt64(801), _179_v50.F8, _179_v50.F8, _dafny.IntOfInt64(333))
		_182_v53 = (_182_v53).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v54, _122_v4)).Cardinality(), _181_v52)
		(_123_globalState).F7 = (_156_v34).Dtor_cf2()
		(_123_globalState).F7 = _179_v50.F8
		var _184_v55 _dafny.Array
		_ = _184_v55
		var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(22))
		_ = _nw28
		_184_v55 = _nw28
		var _185_v56 _dafny.Map
		_ = _185_v56
		_185_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v50, _184_v55)
		var _rhs16 D0 = _157_v35
		_ = _rhs16
		var _rhs17 D0 = _154_v32
		_ = _rhs17
		var _rhs18 _dafny.Int = _179_v50.F8
		_ = _rhs18
		var _rhs19 _dafny.Array = (func() _dafny.Array {
			if _116_v0 {
				return (func() _dafny.Array {
					if (_185_v56).Contains(_179_v50) {
						return (_185_v56).Get(_179_v50).(_dafny.Array)
					}
					return _184_v55
				})()
			}
			return (func() _dafny.Array {
				if _116_v0 {
					return _184_v55
				}
				return _184_v55
			})()
		})()
		_ = _rhs19
		var _lhs15 *C0 = _179_v50
		_ = _lhs15
		_157_v35 = _rhs16
		_154_v32 = _rhs17
		_lhs15.F8 = _rhs18
		_184_v55 = _rhs19
	}
	_124_v5 = _124_v5
	var _186_v57 *C0
	_ = _186_v57
	var _nw29 *C0 = New_C0_()
	_ = _nw29
	_nw29.Ctor__(_dafny.IntOfInt64(530))
	_186_v57 = _nw29
	_118_v2 = _118_v2
	var _hi1 _dafny.Int = _124_v5
	_ = _hi1
	for _187_i8 := _186_v57.F8; _187_i8.Cmp(_hi1) < 0; _187_i8 = _187_i8.Plus(_dafny.One) {
		var _188_v58 _dafny.Map
		_ = _188_v58
		_188_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v57, false)
		_188_v58 = (_188_v58).Update(_186_v57, !((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)))
		var _189_v59 _dafny.Array
		_ = _189_v59
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_7
		var _nw30 _dafny.Array
		_ = _nw30
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw30 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = (func(_190_v4 _dafny.Sequence, _191_v5 _dafny.Int, _192_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_193_i9 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_190_v4, (Companion_Default___.SafeIndex(_191_v5, _dafny.IntOfUint32((_190_v4).Cardinality()))).Uint32(), _192_v39)), _dafny.SeqOf(_190_v4, _dafny.UnicodeSeqOfUtf8Bytes("fkc")))
				}
			})(_122_v4, _124_v5, _163_v39)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw30 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw30).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw30).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_189_v59 = _nw30
		var _194_v60 _dafny.Sequence
		_ = _194_v60
		_194_v60 = _dafny.SeqOf(_122_v4, (Companion_Default___.Fm6(_187_i8, _186_v57.F8, _116_v0, _dafny.CodePoint('f'), _123_globalState)).Dtor_cf12(), _122_v4, _122_v4)
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_189_v59), 0))
		_ = _index29
		(_189_v59).ArraySet1((func() _dafny.Sequence {
			if (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool) {
				return _194_v60
			}
			return _194_v60
		})(), (_index29).Int())
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_189_v59), 0))
		_ = _index30
		(_189_v59).ArraySet1(_194_v60, (_index30).Int())
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
		_ = _index31
		(_120_v3).ArraySet1(_116_v0, (_index31).Int())
		var _195_v61 _dafny.Map
		_ = _195_v61
		_195_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v2, _163_v39), _116_v0)
		var _196_v62 _dafny.Map
		_ = _196_v62
		_196_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_118_v2, _163_v39)
		_116_v0 = !((_128_v9).Select((Companion_Default___.SafeIndex(_186_v57.F8, _dafny.IntOfUint32((_128_v9).Cardinality()))).Uint32()).(bool)) || ((func() bool {
			if (_195_v61).Contains(_196_v62) {
				return (_195_v61).Get(_196_v62).(bool)
			}
			return (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)
		})())
	}
	var _hi2 _dafny.Int = _186_v57.F8
	_ = _hi2
	for _197_i10 := _186_v57.F8; _197_i10.Cmp(_hi2) < 0; _197_i10 = _197_i10.Plus(_dafny.One) {
		var _198_v63 _dafny.Sequence
		_ = _198_v63
		_198_v63 = _dafny.SeqOf(_161_v37)
		var _199_v64 _dafny.Map
		_ = _199_v64
		_199_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_198_v63).Cardinality()), _116_v0)
		var _200_v65 _dafny.Map
		_ = _200_v65
		_200_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v57.F8, _124_v5)
		_199_v64 = (_199_v64).Update(Companion_Default___.SafeDivisionInt((_200_v65).Cardinality(), _197_i10), _116_v0)
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
		_ = _index32
		(_120_v3).ArraySet1(Companion_Default___.Fm0((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _127_v8, _186_v57.F8, _123_globalState), (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
		_ = _index33
		(_120_v3).ArraySet1((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_index33).Int())
		_157_v35 = _157_v35
	}
	var _201_v66 _dafny.Array
	_ = _201_v66
	var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
	_ = _nw31
	_201_v66 = _nw31
	var _202_v67 _dafny.Sequence
	_ = _202_v67
	_202_v67 = _dafny.SeqOf(_124_v5)
	var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_201_v66), 0))
	_ = _index34
	(_201_v66).ArraySet1(_202_v67, (_index34).Int())
	var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_201_v66), 0))
	_ = _index35
	(_201_v66).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(466))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}((func(_203_v57 *C0) func(_dafny.Int) _dafny.Int {
		return func(_204_i11 _dafny.Int) _dafny.Int {
			return _203_v57.F8
		}
	})(_186_v57))), _202_v67), (_index35).Int())
	var _hi3 _dafny.Int = (_124_v5).Plus(_dafny.IntOfInt64(-16))
	_ = _hi3
	for _205_i12 := (_dafny.Zero).Minus(_124_v5); _205_i12.Cmp(_hi3) < 0; _205_i12 = _205_i12.Plus(_dafny.One) {
		if !(!_dafny.Companion_Sequence_.Contains(_122_v4, _163_v39)) {
			var _206_v68 _dafny.Map
			_ = _206_v68
			_206_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v0, _186_v57)
			var _207_v69 _dafny.Map
			_ = _207_v69
			_207_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _206_v68)
			_207_v69 = (_207_v69).Update(false, (_206_v68).Update((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _186_v57))
			var _208_v70 _dafny.MultiSet
			_ = _208_v70
			_208_v70 = _dafny.MultiSetOf(_116_v0, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _116_v0)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
			_ = _index36
			var _rhs20 bool = !((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
			_ = _rhs20
			var _rhs21 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(454), _186_v57.F8)
			_ = _rhs21
			var _rhs22 _dafny.Int = _dafny.IntOfInt64(-651)
			_ = _rhs22
			var _rhs23 _dafny.Int = ((_186_v57.F8).Minus(_124_v5)).Minus(_186_v57.F8)
			_ = _rhs23
			var _rhs24 bool = (func() bool {
				if true {
					return (func() bool {
						if (_161_v37).Contains(_116_v0) {
							return (_161_v37).Get(_116_v0).(bool)
						}
						return Companion_Default___.Fm0(_116_v0, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _127_v8, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(109))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg16 _dafny.Int) interface{} {
								return coer16(arg16)
							}
						}((func(_209_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_210_i13 _dafny.Int) _dafny.CodePoint {
								return _209_v39
							}
						})(_163_v39)))).Cardinality()), _123_globalState)
					})()
				}
				return (_208_v70).IsSubsetOf(_208_v70)
			})()
			_ = _rhs24
			var _lhs16 *C0 = _186_v57
			_ = _lhs16
			var _lhs17 *C0 = _186_v57
			_ = _lhs17
			var _lhs18 *C0 = _186_v57
			_ = _lhs18
			var _lhs19 _dafny.Array = _120_v3
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
			_ = _lhs20
			_116_v0 = _rhs20
			_lhs16.F8 = _rhs21
			_lhs17.F8 = _rhs22
			_lhs18.F8 = _rhs23
			(_lhs19).ArraySet1(_rhs24, (_lhs20).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
			_ = _index37
			(_120_v3).ArraySet1((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_index37).Int())
			var _211_v71 D2
			_ = _211_v71
			_211_v71 = Companion_D2_.Create_DC5_(_dafny.SeqOf((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)))
			var _212_v72 _dafny.Map
			_ = _212_v72
			_212_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v39, _211_v71)
			_212_v72 = (_212_v72).Update(_163_v39, _211_v71)
			_124_v5 = _dafny.IntOfInt64(-669)
		} else {
			var _213_v74 D2
			_ = _213_v74
			_213_v74 = Companion_D2_.Create_DC5_(_128_v9)
			var _214_v75 _dafny.Map
			_ = _214_v75
			_214_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v4, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
			var _215_v76 D4
			_ = _215_v76
			_215_v76 = Companion_D4_.Create_DC9_(_214_v75)
			var _216_v78 _dafny.Set
			_ = _216_v78
			_216_v78 = _dafny.SetOf(_122_v4)
			var _217_v79 _dafny.Sequence
			_ = _217_v79
			_217_v79 = _dafny.SeqOf(func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter7 := _dafny.Iterate((_216_v78).Elements()); ; {
					_compr_6, _ok7 := _iter7()
					if !_ok7 {
						break
					}
					var _218_v77 _dafny.Sequence
					_218_v77 = interface{}(_compr_6).(_dafny.Sequence)
					if (_216_v78).Contains(_218_v77) {
						_coll6.Add(_218_v77, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
					}
				}
				return _coll6.ToMap()
			}(), _214_v75)
			var _219_v80 _dafny.Array
			_ = _219_v80
			var _nwElement0_3 _dafny.Map = (func() _dafny.Map {
				var _coll7 = _dafny.NewMapBuilder()
				_ = _coll7
				for _iter8 := _dafny.Iterate((Companion_Default___.Fm7(_116_v0, _213_v74, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _123_globalState)).Keys().Elements()); ; {
					_compr_7, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _220_v73 _dafny.Sequence
					_220_v73 = interface{}(_compr_7).(_dafny.Sequence)
					if (Companion_Default___.Fm7(_116_v0, _213_v74, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _123_globalState)).Contains(_220_v73) {
						_coll7.Add(_220_v73, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
					}
				}
				return _coll7.ToMap()
			}()).Merge(_214_v75)
			_ = _nwElement0_3
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_3, 0)
			(_nw32).ArraySet1(_214_v75, 1)
			(_nw32).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v4, false), 2)
			(_nw32).ArraySet1(_214_v75, 3)
			(_nw32).ArraySet1(_214_v75, 4)
			(_nw32).ArraySet1(_214_v75, 5)
			(_nw32).ArraySet1(_214_v75, 6)
			(_nw32).ArraySet1(_214_v75, 7)
			(_nw32).ArraySet1(_214_v75, 8)
			(_nw32).ArraySet1((_214_v75).Merge(_214_v75), 9)
			(_nw32).ArraySet1(_214_v75, 10)
			(_nw32).ArraySet1((_214_v75).Merge(_214_v75), 11)
			(_nw32).ArraySet1((_215_v76).Dtor_cf14(), 12)
			(_nw32).ArraySet1(_214_v75, 13)
			(_nw32).ArraySet1((_217_v79).Select((Companion_Default___.SafeIndex(_124_v5, _dafny.IntOfUint32((_217_v79).Cardinality()))).Uint32()).(_dafny.Map), 14)
			(_nw32).ArraySet1(_214_v75, 15)
			(_nw32).ArraySet1(_214_v75, 16)
			_219_v80 = _nw32
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_219_v80), 0))
			_ = _index38
			(_219_v80).ArraySet1(Companion_Default___.Fm8(Companion_Default___.Fm0(false, false, _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v39, _116_v0), _125_v6), _dafny.IntOfInt64(10), _123_globalState), _123_globalState), (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_219_v80), 0))
			_ = _index39
			(_219_v80).ArraySet1((_214_v75).Merge(_214_v75), (_index39).Int())
			_122_v4 = _dafny.Companion_Sequence_.Concatenate(_122_v4, _122_v4)
			var _221_v81 D3
			_ = _221_v81
			_221_v81 = Companion_D3_.Create_DC7_(_161_v37)
			var _222_v82 D3
			_ = _222_v82
			_222_v82 = Companion_D3_.Create_DC8_(_122_v4, _202_v67)
			var _223_v83 _dafny.MultiSet
			_ = _223_v83
			_223_v83 = _dafny.MultiSetOf(_222_v82, _222_v82)
			_221_v81 = Companion_Default___.Fm9(_186_v57.F8, (_223_v83).Cardinality(), _123_globalState)
			(_123_globalState).F7 = (_dafny.Zero).Minus(_186_v57.F8)
			(_186_v57).F8 = _186_v57.F8
		}
		(_186_v57).F8 = _186_v57.F8
		var _224_v84 _dafny.Map
		_ = _224_v84
		_224_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_186_v57, !(_116_v0) || (_116_v0))
		_224_v84 = (_224_v84).Update(_186_v57, _116_v0)
		var _225_v85 _dafny.Array
		_ = _225_v85
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_8
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Set = (func(_226_v37 _dafny.Map) func(_dafny.Int) _dafny.Set {
				return func(_227_i14 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_226_v37)
				}
			})(_161_v37)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw33 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw33).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw33).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_225_v85 = _nw33
		var _228_v86 D5
		_ = _228_v86
		_228_v86 = Companion_D5_.Create_DC11_(_163_v39)
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_151_v29), 0))
		_ = _index40
		(_151_v29).ArraySet1CodePoint((_228_v86).Dtor_cf17(), (_index40).Int())
		var _229_v87 _dafny.MultiSet
		_ = _229_v87
		_229_v87 = _dafny.MultiSetOf(Companion_Default___.Fm0(_116_v0, (_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool), _127_v8, _dafny.IntOfUint32((_202_v67).Cardinality()), _123_globalState))
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_151_v29), 0))
		_ = _index41
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
		_ = _index42
		var _rhs25 _dafny.Int = (_124_v5).Times((func() _dafny.Int {
			if (_229_v87).Contains((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool)) {
				return (_229_v87).Multiplicity((_120_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))).Int()).(bool))
			}
			return _205_i12
		})())
		_ = _rhs25
		var _rhs26 _dafny.Array = _225_v85
		_ = _rhs26
		var _rhs27 _dafny.CodePoint = _dafny.CodePoint('r')
		_ = _rhs27
		var _rhs28 _dafny.Int = _186_v57.F8
		_ = _rhs28
		var _rhs29 bool = true
		_ = _rhs29
		var _lhs21 *C0 = _186_v57
		_ = _lhs21
		var _lhs22 _dafny.Array = _151_v29
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_151_v29), 0))
		_ = _lhs23
		var _lhs24 *GlobalState = _123_globalState
		_ = _lhs24
		var _lhs25 _dafny.Array = _120_v3
		_ = _lhs25
		var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(619), _dafny.ArrayLen((_120_v3), 0))
		_ = _lhs26
		_lhs21.F8 = _rhs25
		_225_v85 = _rhs26
		(_lhs22).ArraySet1CodePoint(_rhs27, (_lhs23).Int())
		_lhs24.F7 = _rhs28
		(_lhs25).ArraySet1(_rhs29, (_lhs26).Int())
	}
	_dafny.Print(_116_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_117_v1).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v2).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v3).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_122_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_123_globalState).F2()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_123_globalState.F4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_globalState).F6().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_123_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_124_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_125_v6).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_126_v7, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_v8).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_128_v9, _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v10).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_130_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v29).ArrayGet1CodePoint((_dafny.IntOfInt64(18)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_152_v30).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_153_v31).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(681))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_154_v32).Dtor_cf5()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(9))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v32).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v32).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_155_v33).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v34).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v34).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v34).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_v35).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v36).Equals(_dafny.MultiSetOf(Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(986)), Companion_D0_.Create_DC2_(_dafny.IntOfInt64(681)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v37).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_162_v38).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_163_v39)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_186_v57.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_201_v66).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(530), _dafny.IntOfInt64(681))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_202_v67, _dafny.SeqOf(_dafny.IntOfInt64(681))))
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
	Cf1 _dafny.Int
	Cf2 _dafny.Int
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Int) D0 {
	return D0{D0_DC2{Cf4}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf5 _dafny.Map
	Cf6 bool
	Cf7 _dafny.Int
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf5 _dafny.Map, Cf6 bool, Cf7 _dafny.Int) D0 {
	return D0{D0_DC3{Cf5, Cf6, Cf7}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Map {
	return _this.Get_().(D0_DC3).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC3).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC3).Cf7
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0
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
	Cf8 _dafny.Set
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Set) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D1) Dtor_cf8() _dafny.Set {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf8.Equals(data2.Cf8)
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

type D2_DC6 struct {
	Cf10 bool
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 bool) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC5 struct {
	Cf9 _dafny.Sequence
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC5{Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(false)
}

func (_this D2) Dtor_cf10() bool {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10 == data2.Cf10
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D3_DC8 struct {
	Cf12 _dafny.Sequence
	Cf13 _dafny.Sequence
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf12 _dafny.Sequence, Cf13 _dafny.Sequence) D3 {
	return D3{D3_DC8{Cf12, Cf13}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf11 _dafny.Map
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf11 _dafny.Map) D3 {
	return D3{D3_DC7{Cf11}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D3) Dtor_cf12() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D3_DC8).Cf13
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC7).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + data.Cf12.VerbatimString(true) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf12.Equals(data2.Cf12) && data1.Cf13.Equals(data2.Cf13)
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf11.Equals(data2.Cf11)
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

type D4_DC10 struct {
	Cf15 bool
	Cf16 *C0
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf15 bool, Cf16 *C0) D4 {
	return D4{D4_DC10{Cf15, Cf16}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

type D4_DC9 struct {
	Cf14 _dafny.Map
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf14 _dafny.Map) D4 {
	return D4{D4_DC9{Cf14}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC10_(false, (*C0)(nil))
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC10).Cf15
}

func (_this D4) Dtor_cf16() *C0 {
	return _this.Get_().(D4_DC10).Cf16
}

func (_this D4) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf15 == data2.Cf15 && data1.Cf16 == data2.Cf16
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf14.Equals(data2.Cf14)
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

type D5_DC12 struct {
	Cf18 bool
	Cf19 _dafny.Int
	Cf20 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf18 bool, Cf19 _dafny.Int, Cf20 _dafny.Int) D5 {
	return D5{D5_DC12{Cf18, Cf19, Cf20}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
	Cf21 _dafny.Sequence
	Cf22 bool
	Cf23 _dafny.Int
	Cf24 bool
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf21 _dafny.Sequence, Cf22 bool, Cf23 _dafny.Int, Cf24 bool) D5 {
	return D5{D5_DC13{Cf21, Cf22, Cf23, Cf24}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC14 struct {
	Cf25 _dafny.Int
	Cf26 bool
	Cf27 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 _dafny.Int, Cf26 bool, Cf27 bool) D5 {
	return D5{D5_DC14{Cf25, Cf26, Cf27}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC11 struct {
	Cf17 _dafny.CodePoint
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf17 _dafny.CodePoint) D5 {
	return D5{D5_DC11{Cf17}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC15 struct {
	Cf28 D5
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf28 D5) D5 {
	return D5{D5_DC15{Cf28}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D5) Dtor_cf18() bool {
	return _this.Get_().(D5_DC12).Cf18
}

func (_this D5) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf20
}

func (_this D5) Dtor_cf21() _dafny.Sequence {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf22() bool {
	return _this.Get_().(D5_DC13).Cf22
}

func (_this D5) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) Dtor_cf24() bool {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) Dtor_cf26() bool {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) Dtor_cf17() _dafny.CodePoint {
	return _this.Get_().(D5_DC11).Cf17
}

func (_this D5) Dtor_cf28() D5 {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf21.Equals(data2.Cf21) && data1.Cf22 == data2.Cf22 && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf28.Equals(data2.Cf28)
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

type D6_DC17 struct {
	Cf30 bool
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 bool) D6 {
	return D6{D6_DC17{Cf30}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC18 struct {
	Cf31 bool
	Cf32 bool
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf31 bool, Cf32 bool) D6 {
	return D6{D6_DC18{Cf31, Cf32}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC19 struct {
}

func (D6_DC19) isD6() {}

func (CompanionStruct_D6_) Create_DC19_() D6 {
	return D6{D6_DC19{}}
}

func (_this D6) Is_DC19() bool {
	_, ok := _this.Get_().(D6_DC19)
	return ok
}

type D6_DC16 struct {
	Cf29 _dafny.Map
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf29 _dafny.Map) D6 {
	return D6{D6_DC16{Cf29}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(false)
}

func (_this D6) Dtor_cf30() bool {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf31() bool {
	return _this.Get_().(D6_DC18).Cf31
}

func (_this D6) Dtor_cf32() bool {
	return _this.Get_().(D6_DC18).Cf32
}

func (_this D6) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D6_DC19:
		{
			return "D6.DC19"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf30 == data2.Cf30
		}
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf31 == data2.Cf31 && data1.Cf32 == data2.Cf32
		}
	case D6_DC19:
		{
			_, ok := other.Get_().(D6_DC19)
			return ok
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf29.Equals(data2.Cf29)
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

// Definition of class GlobalState
type GlobalState struct {
	F4  _dafny.Sequence
	F7  _dafny.Int
	_f0 _dafny.Map
	_f1 _dafny.Array
	_f2 _dafny.Array
	_f3 _dafny.Int
	_f5 bool
	_f6 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.EmptySeq
	_this.F7 = _dafny.Zero
	_this._f0 = _dafny.EmptyMap
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f3 = _dafny.Zero
	_this._f5 = false
	_this._f6 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Array, f2 _dafny.Array, f3 _dafny.Int, f4 _dafny.Sequence, f5 bool, f6 _dafny.Sequence, f7 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
	}
}
func (_this *GlobalState) F0() _dafny.Map {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Array {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Sequence {
	{
		return _this._f6
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F8 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F8 = _dafny.Zero
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

func (_this *C0) Ctor__(f8 _dafny.Int) {
	{
		(_this).F8 = f8
	}
}
func (_this *C0) M1(globalState *GlobalState) (bool, bool, _dafny.Int, _dafny.CodePoint) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r3
		var _230_v0 _dafny.Sequence
		_ = _230_v0
		_230_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wdjxni")
		var _231_v1 bool
		_ = _231_v1
		_231_v1 = true
		var _232_v2 _dafny.Map
		_ = _232_v2
		_232_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_231_v1, _this.F8)
		var _233_v3 _dafny.Array
		_ = _233_v3
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw34
		_233_v3 = _nw34
		var _234_v4 _dafny.Map
		_ = _234_v4
		_234_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v3, _this.F8)
		var _235_v5 _dafny.Sequence
		_ = _235_v5
		_235_v5 = _dafny.SeqOf((_232_v2).Cardinality(), (_234_v4).Cardinality())
		var _236_v6 _dafny.Array
		_ = _236_v6
		var _nwElement0_4 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_230_v0, _230_v0)
		_ = _nwElement0_4
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(4))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_4, 0)
		(_nw35).ArraySet1(!_dafny.Companion_Sequence_.Contains((Companion_D0_.Create_DC0_(_dafny.SeqOf(_235_v5))).Dtor_cf0(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(117))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_237_i0 _dafny.Int) _dafny.Int {
			return _this.F8
		}))), 1)
		(_nw35).ArraySet1(_231_v1, 2)
		(_nw35).ArraySet1(_231_v1, 3)
		_236_v6 = _nw35
		var _238_v7 _dafny.Map
		_ = _238_v7
		_238_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_231_v1), _231_v1)
		var _239_v8 _dafny.Set
		_ = _239_v8
		_239_v8 = _dafny.SetOf(_238_v7, ((_238_v7).Update(_231_v1, false)).Update(_231_v1, true), _238_v7)
		var _240_v9 _dafny.CodePoint
		_ = _240_v9
		_240_v9 = _dafny.CodePoint('f')
		var _241_v11 _dafny.Set
		_ = _241_v11
		_241_v11 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _231_v1), _238_v7)
		var _rhs30 _dafny.Int = (func() _dafny.Set {
			var _coll8 = _dafny.NewBuilder()
			_ = _coll8
			for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v9, !(_231_v1)), _231_v1)).Keys().Elements()); ; {
				_compr_8, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _242_v10 _dafny.Map
				_242_v10 = interface{}(_compr_8).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v9, !(_231_v1)), _231_v1)).Contains(_242_v10) {
					_coll8.Add(_242_v10)
				}
			}
			return _coll8.ToSet()
		}()).Cardinality()
		_ = _rhs30
		var _rhs31 _dafny.Array = _236_v6
		_ = _rhs31
		var _rhs32 _dafny.Int = Companion_Default___.SafeDivisionInt((_this.F8).Times(_this.F8), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_231_v1)).Cardinality())))
		_ = _rhs32
		var _rhs33 _dafny.Int = (_this.F8).Plus(_this.F8)
		_ = _rhs33
		var _rhs34 _dafny.Set = (_241_v11)
		_ = _rhs34
		var _lhs27 *GlobalState = globalState
		_ = _lhs27
		var _lhs28 *GlobalState = globalState
		_ = _lhs28
		var _lhs29 *GlobalState = globalState
		_ = _lhs29
		_lhs27.F7 = _rhs30
		_236_v6 = _rhs31
		_lhs28.F7 = _rhs32
		_lhs29.F7 = _rhs33
		_239_v8 = _rhs34
		var _243_v12 _dafny.Sequence
		_ = _243_v12
		_243_v12 = _dafny.SeqOf(_231_v1)
		var _244_v13 _dafny.Sequence
		_ = _244_v13
		_244_v13 = _dafny.SeqOf(_243_v12)
		_243_v12 = _dafny.Companion_Sequence_.Concatenate(_243_v12, (_244_v13).Select((Companion_Default___.SafeIndex(_this.F8, _dafny.IntOfUint32((_244_v13).Cardinality()))).Uint32()).(_dafny.Sequence))
		(globalState).F7 = _this.F8
		(globalState).F7 = ((_dafny.Zero).Minus(_this.F8)).Times(_dafny.IntOfInt64(616))
		(_this).F8 = _this.F8
		var _245_v14 _dafny.MultiSet
		_ = _245_v14
		_245_v14 = _dafny.MultiSetOf(_this.F8)
		Companion_Default___.M0(_this.F8, _245_v14, _231_v1, _this.F8, globalState)
		var _246_v15 _dafny.Map
		_ = _246_v15
		_246_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v9, _231_v1)
		var _247_v16 _dafny.MultiSet
		_ = _247_v16
		_247_v16 = _dafny.MultiSetOf(_246_v15)
		r0 = Companion_Default___.Fm0(Companion_Default___.Fm0(true, _231_v1, _247_v16, (_239_v8).Cardinality(), globalState), _231_v1, ((_247_v16).Update(func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter10 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_248_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_249_i1 _dafny.Int) _dafny.CodePoint {
					return _248_v9
				}
			})(_240_v9)))).Elements()); ; {
				_compr_9, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _250_v17 _dafny.CodePoint
				_250_v17 = interface{}(_compr_9).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(605))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_251_v9 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_249_i1 _dafny.Int) _dafny.CodePoint {
						return _251_v9
					}
				})(_240_v9))), _250_v17) {
					_coll9.Add(_250_v17, !(_231_v1))
				}
			}
			return _coll9.ToMap()
		}(), Companion_Default___.Abs(_this.F8))).Union(_247_v16), _this.F8, globalState)
		r1 = _231_v1
		var _252_v18 _dafny.Set
		_ = _252_v18
		_252_v18 = _dafny.SetOf(_231_v1, true)
		r2 = Companion_Default___.SafeDivisionInt((_252_v18).Cardinality(), ((_dafny.Zero).Minus(_this.F8)).Minus(_this.F8))
		var _253_v19 D0
		_ = _253_v19
		_253_v19 = Companion_D0_.Create_DC3_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(!(_231_v1), false, _231_v1, !(_231_v1))).Cardinality()))).Update(_231_v1, _this.F8), _231_v1, _this.F8)
		var _pat_let_tv3 = _240_v9
		_ = _pat_let_tv3
		var _pat_let_tv4 = _240_v9
		_ = _pat_let_tv4
		var _pat_let_tv5 = _230_v0
		_ = _pat_let_tv5
		var _pat_let_tv6 = _230_v0
		_ = _pat_let_tv6
		r3 = func(_source2 D0) _dafny.CodePoint {
			if _source2.Is_DC1() {
				var _254___mcc_h0 _dafny.Int = _source2.Get_().(D0_DC1).Cf1
				_ = _254___mcc_h0
				var _255___mcc_h1 _dafny.Int = _source2.Get_().(D0_DC1).Cf2
				_ = _255___mcc_h1
				var _256___mcc_h2 _dafny.Int = _source2.Get_().(D0_DC1).Cf3
				_ = _256___mcc_h2
				var _257_cf3 _dafny.Int = _256___mcc_h2
				_ = _257_cf3
				var _258_cf2 _dafny.Int = _255___mcc_h1
				_ = _258_cf2
				var _259_cf1 _dafny.Int = _254___mcc_h0
				_ = _259_cf1
				return _pat_let_tv3
			} else if _source2.Is_DC2() {
				var _260___mcc_h3 _dafny.Int = _source2.Get_().(D0_DC2).Cf4
				_ = _260___mcc_h3
				var _261_cf4 _dafny.Int = _260___mcc_h3
				_ = _261_cf4
				return _dafny.CodePoint('d')
			} else if _source2.Is_DC3() {
				var _262___mcc_h4 _dafny.Map = _source2.Get_().(D0_DC3).Cf5
				_ = _262___mcc_h4
				var _263___mcc_h5 bool = _source2.Get_().(D0_DC3).Cf6
				_ = _263___mcc_h5
				var _264___mcc_h6 _dafny.Int = _source2.Get_().(D0_DC3).Cf7
				_ = _264___mcc_h6
				var _265_cf7 _dafny.Int = _264___mcc_h6
				_ = _265_cf7
				var _266_cf6 bool = _263___mcc_h5
				_ = _266_cf6
				var _267_cf5 _dafny.Map = _262___mcc_h4
				_ = _267_cf5
				return _pat_let_tv4
			} else {
				var _268___mcc_h7 _dafny.Sequence = _source2.Get_().(D0_DC0).Cf0
				_ = _268___mcc_h7
				var _269_cf0 _dafny.Sequence = _268___mcc_h7
				_ = _269_cf0
				return (_pat_let_tv5).Select((Companion_Default___.SafeIndex(_this.F8, _dafny.IntOfUint32((_pat_let_tv6).Cardinality()))).Uint32()).(_dafny.CodePoint)
			}
		}(_253_v19)
		return r0, r1, r2, r3
	}
}

// End of class C0
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
