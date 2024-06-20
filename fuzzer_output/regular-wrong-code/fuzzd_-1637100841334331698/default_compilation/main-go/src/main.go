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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("cfbosyfj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(990))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	})), _dafny.UnicodeSeqOfUtf8Bytes("w"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(890))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(424))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-981)), !(false))).Cardinality(), Companion_D1_.Create_DC3_(_dafny.SeqOf(!(true))))).Cardinality())).Cardinality()).Times((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SetOf(func() _dafny.Set {
				var _coll2 = _dafny.NewBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
					var _coll3 = _dafny.NewMapBuilder()
					_ = _coll3
					for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
						_compr_3, _ok3 := _iter3()
						if !_ok3 {
							break
						}
						var _3_v1 _dafny.Int
						_3_v1 = interface{}(_compr_3).(_dafny.Int)
						if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
							_coll3.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
						}
					}
					return _coll3.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _4_v2 _dafny.Int
					_4_v2 = interface{}(_compr_2).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
						var _coll4 = _dafny.NewMapBuilder()
						_ = _coll4
						for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
							_compr_4, _ok4 := _iter4()
							if !_ok4 {
								break
							}
							var _3_v1 _dafny.Int
							_3_v1 = interface{}(_compr_4).(_dafny.Int)
							if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
								_coll4.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
							}
						}
						return _coll4.ToMap()
					}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Contains(_4_v2) {
						_coll2.Add((_4_v2).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg3 _dafny.Int) interface{} {
								return coer3(arg3)
							}
						}(func(_5_i1 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg4 _dafny.Int) interface{} {
									return coer4(arg4)
								}
							}(func(_6_i2 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
							}))).Cardinality()))
						}))).Cardinality())))
					}
				}
				return _coll2.ToSet()
			}())).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _7_v0 _dafny.Set
				_7_v0 = interface{}(_compr_1).(_dafny.Set)
				if (_dafny.SetOf(func() _dafny.Set {
					var _coll5 = _dafny.NewBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
						var _coll6 = _dafny.NewMapBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _3_v1 _dafny.Int
							_3_v1 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
								_coll6.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
							}
						}
						return _coll6.ToMap()
					}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _8_v2 _dafny.Int
						_8_v2 = interface{}(_compr_5).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
							var _coll7 = _dafny.NewMapBuilder()
							_ = _coll7
							for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
								_compr_7, _ok7 := _iter7()
								if !_ok7 {
									break
								}
								var _3_v1 _dafny.Int
								_3_v1 = interface{}(_compr_7).(_dafny.Int)
								if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
									_coll7.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
								}
							}
							return _coll7.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Contains(_8_v2) {
							_coll5.Add((_8_v2).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg5 _dafny.Int) interface{} {
									return coer5(arg5)
								}
							}(func(_5_i1 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg6 _dafny.Int) interface{} {
										return coer6(arg6)
									}
								}(func(_6_i2 _dafny.Int) _dafny.Int {
									return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
								}))).Cardinality()))
							}))).Cardinality())))
						}
					}
					return _coll5.ToSet()
				}())).Contains(_7_v0) {
					_coll1.Add(_7_v0, _dafny.IntOfInt64(90))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality(), !(true))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _9_v3 _dafny.Int
			_9_v3 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((_dafny.SetOf(func() _dafny.Set {
					var _coll9 = _dafny.NewBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
						var _coll10 = _dafny.NewMapBuilder()
						_ = _coll10
						for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
							_compr_10, _ok10 := _iter10()
							if !_ok10 {
								break
							}
							var _3_v1 _dafny.Int
							_3_v1 = interface{}(_compr_10).(_dafny.Int)
							if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
								_coll10.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
							}
						}
						return _coll10.ToMap()
					}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _10_v2 _dafny.Int
						_10_v2 = interface{}(_compr_9).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
								_compr_11, _ok11 := _iter11()
								if !_ok11 {
									break
								}
								var _3_v1 _dafny.Int
								_3_v1 = interface{}(_compr_11).(_dafny.Int)
								if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
									_coll11.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
								}
							}
							return _coll11.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Contains(_10_v2) {
							_coll9.Add((_10_v2).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg7 _dafny.Int) interface{} {
									return coer7(arg7)
								}
							}(func(_5_i1 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg8 _dafny.Int) interface{} {
										return coer8(arg8)
									}
								}(func(_6_i2 _dafny.Int) _dafny.Int {
									return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
								}))).Cardinality()))
							}))).Cardinality())))
						}
					}
					return _coll9.ToSet()
				}())).Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _7_v0 _dafny.Set
					_7_v0 = interface{}(_compr_8).(_dafny.Set)
					if (_dafny.SetOf(func() _dafny.Set {
						var _coll12 = _dafny.NewBuilder()
						_ = _coll12
						for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
							var _coll13 = _dafny.NewMapBuilder()
							_ = _coll13
							for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
								_compr_13, _ok13 := _iter13()
								if !_ok13 {
									break
								}
								var _3_v1 _dafny.Int
								_3_v1 = interface{}(_compr_13).(_dafny.Int)
								if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
									_coll13.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
								}
							}
							return _coll13.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Keys().Elements()); ; {
							_compr_12, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _11_v2 _dafny.Int
							_11_v2 = interface{}(_compr_12).(_dafny.Int)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Map {
								var _coll14 = _dafny.NewMapBuilder()
								_ = _coll14
								for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(280), _dafny.IntOfInt64(120))); ; {
									_compr_14, _ok14 := _iter14()
									if !_ok14 {
										break
									}
									var _3_v1 _dafny.Int
									_3_v1 = interface{}(_compr_14).(_dafny.Int)
									if ((_dafny.IntOfInt64(280)).Cmp(_3_v1) <= 0) && ((_3_v1).Cmp(_dafny.IntOfInt64(120)) < 0) {
										_coll14.Add((_3_v1).Times(_dafny.IntOfInt64(-118)), (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cth")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(340), _dafny.IntOfInt64(90))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true, true)).Cardinality(), _dafny.IntOfInt64(622)))).Cardinality())
									}
								}
								return _coll14.ToMap()
							}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(-625))).Cardinality())).Cardinality())))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-234), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpmel")).Cardinality()))).Cardinality()))).Cardinality())).Contains(_11_v2) {
								_coll12.Add((_11_v2).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(751))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg9 _dafny.Int) interface{} {
										return coer9(arg9)
									}
								}(func(_5_i1 _dafny.Int) _dafny.Int {
									return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg10 _dafny.Int) interface{} {
											return coer10(arg10)
										}
									}(func(_6_i2 _dafny.Int) _dafny.Int {
										return _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())
									}))).Cardinality()))
								}))).Cardinality())))
							}
						}
						return _coll12.ToSet()
					}())).Contains(_7_v0) {
						_coll8.Add(_7_v0, _dafny.IntOfInt64(90))
					}
				}
				return _coll8.ToMap()
			}()).Cardinality(), !(true))).Contains(_9_v3) {
				_coll0.Add((_9_v3).Plus((_dafny.SetOf(_dafny.IntOfInt64(-920))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
		var _coll15 = _dafny.NewBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_12_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()), (_dafny.SetOf(false)).Cardinality())).Keys().Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _13_v4 _dafny.Int
			_13_v4 = interface{}(_compr_15).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_12_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))).Cardinality()), (_dafny.SetOf(false)).Cardinality())).Contains(_13_v4) {
				_coll15.Add((_13_v4).Times(_dafny.IntOfInt64(438)))
			}
		}
		return _coll15.ToSet()
	}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jiwdrv")).Cardinality()))).Cardinality()).Minus(_dafny.IntOfInt64(630)))
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	var _source0 _dafny.Set = _dafny.SetOf(false)
	_ = _source0
	var _14___mcc_h0 _dafny.Set = _source0
	_ = _14___mcc_h0
	var _15_cf23 _dafny.Set = _14___mcc_h0
	_ = _15_cf23
	return (_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 D0, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-621))))).Difference((_dafny.MultiSetOf(_dafny.IntOfInt64(545))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(352), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	var _source1 D14 = Companion_D14_.Create_DC32_(_dafny.MultiSetOf(!(true)), true, false, !(true), _dafny.IntOfInt64(542))
	_ = _source1
	if _source1.Is_DC32() {
		var _16___mcc_h0 _dafny.MultiSet = _source1.Get_().(D14_DC32).Cf56
		_ = _16___mcc_h0
		var _17___mcc_h1 bool = _source1.Get_().(D14_DC32).Cf57
		_ = _17___mcc_h1
		var _18___mcc_h2 bool = _source1.Get_().(D14_DC32).Cf58
		_ = _18___mcc_h2
		var _19___mcc_h3 bool = _source1.Get_().(D14_DC32).Cf59
		_ = _19___mcc_h3
		var _20___mcc_h4 _dafny.Int = _source1.Get_().(D14_DC32).Cf60
		_ = _20___mcc_h4
		var _21_cf60 _dafny.Int = _20___mcc_h4
		_ = _21_cf60
		var _22_cf59 bool = _19___mcc_h3
		_ = _22_cf59
		var _23_cf58 bool = _18___mcc_h2
		_ = _23_cf58
		var _24_cf57 bool = _17___mcc_h1
		_ = _24_cf57
		var _25_cf56 _dafny.MultiSet = _16___mcc_h0
		_ = _25_cf56
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_22_cf59), _dafny.SeqOf(_23_cf58))
	} else {
		var _26___mcc_h5 _dafny.Sequence = _source1.Get_().(D14_DC31).Cf55
		_ = _26___mcc_h5
		var _27_cf55 _dafny.Sequence = _26___mcc_h5
		_ = _27_cf55
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), true), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('d')
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(835))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_28_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(false)).Cardinality()
	})), _dafny.SeqOf(_dafny.IntOfInt64(215))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_29_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(600)
	})))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Set, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(!(false) || (!(false)))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_((_dafny.IntOfInt64(406)).Cmp(_dafny.IntOfInt64(-130)) == 0)
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC9_(_dafny.IntOfInt64(-36), _dafny.IntOfInt64(441), (false) && (false))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false, false, !(true))).Intersection(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("vvtodqg")
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	var _source2 D1 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(796))
	_ = _source2
	if _source2.Is_DC4() {
		var _30___mcc_h0 _dafny.Array = _source2.Get_().(D1_DC4).Cf7
		_ = _30___mcc_h0
		var _31___mcc_h1 _dafny.Int = _source2.Get_().(D1_DC4).Cf8
		_ = _31___mcc_h1
		var _32___mcc_h2 _dafny.Int = _source2.Get_().(D1_DC4).Cf9
		_ = _32___mcc_h2
		var _33___mcc_h3 _dafny.Sequence = _source2.Get_().(D1_DC4).Cf10
		_ = _33___mcc_h3
		var _34___mcc_h4 _dafny.Sequence = _source2.Get_().(D1_DC4).Cf11
		_ = _34___mcc_h4
		var _35_cf11 _dafny.Sequence = _34___mcc_h4
		_ = _35_cf11
		var _36_cf10 _dafny.Sequence = _33___mcc_h3
		_ = _36_cf10
		var _37_cf9 _dafny.Int = _32___mcc_h2
		_ = _37_cf9
		var _38_cf8 _dafny.Int = _31___mcc_h1
		_ = _38_cf8
		var _39_cf7 _dafny.Array = _30___mcc_h0
		_ = _39_cf7
		if !(true) {
			return _dafny.SeqOf(_38_cf8)
		} else {
			return _dafny.SeqOf((func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-531), _dafny.IntOfInt64(718))); ; {
					_compr_16, _ok16 := _iter16()
					if !_ok16 {
						break
					}
					var _40_v0 _dafny.Int
					_40_v0 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-531)).Cmp(_40_v0) <= 0) && ((_40_v0).Cmp(_dafny.IntOfInt64(718)) < 0) {
						_coll16.Add((_40_v0).Times(_37_cf9), true)
					}
				}
				return _coll16.ToMap()
			}()).Cardinality())
		}
	} else if _source2.Is_DC5() {
		var _41___mcc_h5 _dafny.Int = _source2.Get_().(D1_DC5).Cf12
		_ = _41___mcc_h5
		var _42_cf12 _dafny.Int = _41___mcc_h5
		_ = _42_cf12
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_42_cf12, _dafny.IntOfInt64(584), (_dafny.Zero).Minus(_42_cf12), _42_cf12), _dafny.SeqOf(_42_cf12))
	} else {
		var _43___mcc_h6 _dafny.Sequence = _source2.Get_().(D1_DC3).Cf6
		_ = _43___mcc_h6
		var _44_cf6 _dafny.Sequence = _43___mcc_h6
		_ = _44_cf6
		return _dafny.SeqOf(_dafny.IntOfInt64(107))
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('v')
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(244), _dafny.IntOfInt64(580)), _dafny.SeqOf(_dafny.IntOfInt64(-662), _dafny.IntOfUint32((_dafny.SeqOf(!(false), true, true, true)).Cardinality()))), true)
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('o'))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(327), true)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jsiwka")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vjgqe")).Cardinality())))).Intersection((_dafny.MultiSetFromSeq(_dafny.SeqOf(func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(334), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(false, true)).Cardinality())).Cardinality())).Keys().Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _45_v0 _dafny.Int
			_45_v0 = interface{}(_compr_17).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(334), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(false, true)).Cardinality())).Cardinality())).Contains(_45_v0) {
				_coll17.Add(Companion_Default___.SafeModuloInt(_45_v0, (_dafny.SetOf(true)).Cardinality()), _dafny.IntOfInt64(969))
			}
		}
		return _coll17.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(243), (_dafny.MultiSetOf((Companion_D22_.Create_DC54_(_dafny.CodePoint('u'), Companion_D9_.Create_DC22_(false), true, true, true)).Dtor_cf98(), _dafny.CodePoint('x'))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
		var _coll18 = _dafny.NewMapBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-42), _dafny.IntOfInt64(599))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _46_v1 _dafny.Int
			_46_v1 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(-42)).Cmp(_46_v1) <= 0) && ((_46_v1).Cmp(_dafny.IntOfInt64(599)) < 0) {
				_coll18.Add(Companion_Default___.SafeModuloInt(_46_v1, _dafny.Zero), _dafny.IntOfInt64(335))
			}
		}
		return _coll18.ToMap()
	}()).Cardinality())).Cardinality()))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(48), _dafny.IntOfInt64(789))))).Union(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(162), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bwfgndon")).Cardinality())))))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source3 D22 = Companion_D22_.Create_DC54_(_dafny.CodePoint('r'), Companion_D9_.Create_DC22_(true), false, true, !(false))
	_ = _source3
	if _source3.Is_DC54() {
		var _47___mcc_h0 _dafny.CodePoint = _source3.Get_().(D22_DC54).Cf98
		_ = _47___mcc_h0
		var _48___mcc_h1 D9 = _source3.Get_().(D22_DC54).Cf99
		_ = _48___mcc_h1
		var _49___mcc_h2 bool = _source3.Get_().(D22_DC54).Cf100
		_ = _49___mcc_h2
		var _50___mcc_h3 bool = _source3.Get_().(D22_DC54).Cf101
		_ = _50___mcc_h3
		var _51___mcc_h4 bool = _source3.Get_().(D22_DC54).Cf102
		_ = _51___mcc_h4
		var _52_cf102 bool = _51___mcc_h4
		_ = _52_cf102
		var _53_cf101 bool = _50___mcc_h3
		_ = _53_cf101
		var _54_cf100 bool = _49___mcc_h2
		_ = _54_cf100
		var _55_cf99 D9 = _48___mcc_h1
		_ = _55_cf99
		var _56_cf98 _dafny.CodePoint = _47___mcc_h0
		_ = _56_cf98
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rqm"), _54_cf100)
	} else {
		var _57___mcc_h5 _dafny.Set = _source3.Get_().(D22_DC53).Cf97
		_ = _57___mcc_h5
		var _58_cf97 _dafny.Set = _57___mcc_h5
		_ = _58_cf97
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("smnurn"), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_59_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), !(true)))
	}
}
func (_static *CompanionStruct_Default___) Fm28(globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC1_(true, true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true, false)).Cardinality())).Cardinality(), !(true)), Companion_D0_.Create_DC1_(!(false), !(false), _dafny.IntOfInt64(297), true), Companion_D0_.Create_DC1_(false, !(false), _dafny.IntOfInt64(106), true))).Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _60_v0 D0
			_60_v0 = interface{}(_compr_19).(D0)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC1_(true, true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true, false)).Cardinality())).Cardinality(), !(true)), Companion_D0_.Create_DC1_(!(false), !(false), _dafny.IntOfInt64(297), true), Companion_D0_.Create_DC1_(false, !(false), _dafny.IntOfInt64(106), true)), _60_v0) {
				_coll19.Add(_60_v0)
			}
		}
		return _coll19.ToSet()
	}()).Difference(_dafny.SetOf(Companion_D0_.Create_DC1_(false, !(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), false)))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(667), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jp")).Cardinality()))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('e'))).Cardinality()), _dafny.IntOfUint32(((Companion_D10_.Create_DC24_(_dafny.UnicodeSeqOfUtf8Bytes("lwtjp"), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(227))), true)).Dtor_cf44()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ojauypo")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-979)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D21_.Create_DC49_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(809), _dafny.IntOfInt64(745), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mf")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(326))))).Dtor_cf86()
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (!(false)) && (!(false)) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(true, true)))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC8_()
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ovj")).Cardinality()), _dafny.IntOfInt64(57)), _dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(362)))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vwg")).Cardinality()), _dafny.IntOfInt64(725))))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_61_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uxbg")).Cardinality())
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(364))).Uint32(), func(coer17 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_62_i1 _dafny.Int) D2 {
		return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jlhmrsu")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality())))
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(992))).Cardinality(), _dafny.IntOfInt64(-660))).Cardinality())).Cardinality(), true)).Cardinality(), false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-63))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_63_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	}))).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()), _dafny.SeqOf(Companion_D2_.Create_DC6_(func() _dafny.Map {
		var _coll20 = _dafny.NewMapBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(283))).Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _64_v0 _dafny.Int
			_64_v0 = interface{}(_compr_20).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfInt64(283))).Contains(_64_v0) {
				_coll20.Add((_64_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(980))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}(func(_65_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality()))), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-254))))
			}
		}
		return _coll20.ToMap()
	}())))))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if true {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer20 func(_dafny.Int) D5) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_66_i0 _dafny.Int) D5 {
			return Companion_D5_.Create_DC14_((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(false, true, true))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(161))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_67_i1 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()
			}))).Cardinality()))).Cardinality(), true)
		}))
	} else {
		return _dafny.SeqOf(Companion_D5_.Create_DC14_(_dafny.IntOfInt64(219), (_dafny.SetOf(_dafny.IntOfInt64(274))).Cardinality(), false))
	}
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-68), _dafny.IntOfInt64(-480)))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-255))).Uint32(), func(coer22 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
		return func(arg22 _dafny.Int) interface{} {
			return coer22(arg22)
		}
	}(func(_68_i0 _dafny.Int) D2 {
		return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-248)), _dafny.IntOfInt64(436)))
	})), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(282), _dafny.IntOfInt64(895))), Companion_D2_.Create_DC6_(func() _dafny.Map {
		var _coll21 = _dafny.NewMapBuilder()
		_ = _coll21
		for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(691), _dafny.IntOfInt64(981))); ; {
			_compr_21, _ok21 := _iter21()
			if !_ok21 {
				break
			}
			var _69_v0 _dafny.Int
			_69_v0 = interface{}(_compr_21).(_dafny.Int)
			if ((_dafny.IntOfInt64(691)).Cmp(_69_v0) <= 0) && ((_69_v0).Cmp(_dafny.IntOfInt64(981)) < 0) {
				_coll21.Add((_69_v0).Minus(_dafny.IntOfInt64(202)), _dafny.IntOfInt64(-438))
			}
		}
		return _coll21.ToMap()
	}()), Companion_D2_.Create_DC6_(func() _dafny.Map {
		var _coll22 = _dafny.NewMapBuilder()
		_ = _coll22
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-193), _dafny.IntOfInt64(618))); ; {
			_compr_22, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _70_v1 _dafny.Int
			_70_v1 = interface{}(_compr_22).(_dafny.Int)
			if ((_dafny.IntOfInt64(-193)).Cmp(_70_v1) <= 0) && ((_70_v1).Cmp(_dafny.IntOfInt64(618)) < 0) {
				_coll22.Add(Companion_Default___.SafeDivisionInt(_70_v1, _dafny.IntOfInt64(560)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ao")).Cardinality()))
			}
		}
		return _coll22.ToMap()
	}())), _dafny.SeqOf(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(-649))).Cardinality(), _dafny.IntOfInt64(384))))))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 D10, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("sqdccsvha")), _dafny.CodePoint('h'))).Merge((func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter23 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("nrjkdbhoe")), false)).Keys().Elements()); ; {
			_compr_23, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _71_v0 D6
			_71_v0 = interface{}(_compr_23).(D6)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("nrjkdbhoe")), false)).Contains(_71_v0) {
				_coll23.Add(_71_v0, _dafny.CodePoint('n'))
			}
		}
		return _coll23.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("mxdmgbo")), _dafny.CodePoint('d'))))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC16_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_72_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(135))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_73_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))))
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(479)
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source4 D14 = Companion_D14_.Create_DC31_(_dafny.SeqOf(_dafny.SetOf(!(true)), _dafny.SetOf(false), _dafny.SetOf(!(true)), _dafny.SetOf((Companion_D10_.Create_DC24_(_dafny.UnicodeSeqOfUtf8Bytes("hy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_74_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(855))
	})), true)).Dtor_cf46())))
	_ = _source4
	if _source4.Is_DC32() {
		var _75___mcc_h0 _dafny.MultiSet = _source4.Get_().(D14_DC32).Cf56
		_ = _75___mcc_h0
		var _76___mcc_h1 bool = _source4.Get_().(D14_DC32).Cf57
		_ = _76___mcc_h1
		var _77___mcc_h2 bool = _source4.Get_().(D14_DC32).Cf58
		_ = _77___mcc_h2
		var _78___mcc_h3 bool = _source4.Get_().(D14_DC32).Cf59
		_ = _78___mcc_h3
		var _79___mcc_h4 _dafny.Int = _source4.Get_().(D14_DC32).Cf60
		_ = _79___mcc_h4
		var _80_cf60 _dafny.Int = _79___mcc_h4
		_ = _80_cf60
		var _81_cf59 bool = _78___mcc_h3
		_ = _81_cf59
		var _82_cf58 bool = _77___mcc_h2
		_ = _82_cf58
		var _83_cf57 bool = _76___mcc_h1
		_ = _83_cf57
		var _84_cf56 _dafny.MultiSet = _75___mcc_h0
		_ = _84_cf56
		if _82_cf58 {
			return _dafny.CodePoint('v')
		} else {
			return _dafny.CodePoint('i')
		}
	} else {
		var _85___mcc_h5 _dafny.Sequence = _source4.Get_().(D14_DC31).Cf55
		_ = _85___mcc_h5
		var _86_cf55 _dafny.Sequence = _85___mcc_h5
		_ = _86_cf55
		if true {
			return _dafny.CodePoint('r')
		} else {
			return _dafny.CodePoint('q')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 D3, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SetOf(true), _dafny.SetOf(false, true), (_dafny.SetOf(true)).Difference(_dafny.SetOf(!(!(true)), false, true, false)), (_dafny.SetOf(true, false, false)).Intersection(_dafny.SetOf((Companion_D2_.Create_DC7_(true, true, _dafny.IntOfInt64(-255), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))).Cardinality())).Dtor_cf15(), true, true, !(false), false)), (func() _dafny.Set {
		if true {
			return _dafny.SetOf(true, false, true, false)
		}
		return _dafny.SetOf(!(!(true)))
	})())
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dagoucki")).Cardinality()))).Difference((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-490)))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('q')))).Cardinality(), _dafny.IntOfInt64(760)))))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(920))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_87_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()), _dafny.IntOfInt64(385))))
}
func (_static *CompanionStruct_Default___) Fm45(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((_dafny.IntOfInt64(-250)).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(528), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(117), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chdvwyj")).Cardinality()), (func() _dafny.Map {
		var _coll24 = _dafny.NewMapBuilder()
		_ = _coll24
		for _iter24 := _dafny.Iterate((_dafny.SeqOf(((Companion_D7_.Create_DC19_((_dafny.SetOf(_dafny.CodePoint('v'))).Cardinality(), _dafny.MultiSetOf(true), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(794)))).Dtor_cf37()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_88_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))).Cardinality()))).Elements()); ; {
			_compr_24, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _89_v0 _dafny.Int
			_89_v0 = interface{}(_compr_24).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(((Companion_D7_.Create_DC19_((_dafny.SetOf(_dafny.CodePoint('v'))).Cardinality(), _dafny.MultiSetOf(true), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(794)))).Dtor_cf37()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_88_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			}))).Cardinality())), _89_v0) {
				_coll24.Add((_89_v0).Minus(_dafny.IntOfInt64(423)), _dafny.IntOfInt64(909))
			}
		}
		return _coll24.ToMap()
	}()).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(612))).Cardinality()))).Cardinality())).Cardinality(), !(!(true)))))).Merge((func() _dafny.Map {
		var _coll25 = _dafny.NewMapBuilder()
		_ = _coll25
		for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(442), _dafny.IntOfInt64(367))); ; {
			_compr_25, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _90_v1 _dafny.Int
			_90_v1 = interface{}(_compr_25).(_dafny.Int)
			if ((_dafny.IntOfInt64(442)).Cmp(_90_v1) <= 0) && ((_90_v1).Cmp(_dafny.IntOfInt64(367)) < 0) {
				_coll25.Add((_90_v1).Plus(_dafny.IntOfInt64(-352)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(492), false))
			}
		}
		return _coll25.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(447))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg29 _dafny.Int) interface{} {
			return coer29(arg29)
		}
	}(func(_91_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(465)
	}))).Cardinality()), false)).Cardinality(), func() _dafny.Map {
		var _coll26 = _dafny.NewMapBuilder()
		_ = _coll26
		for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(812), _dafny.IntOfInt64(557))); ; {
			_compr_26, _ok26 := _iter26()
			if !_ok26 {
				break
			}
			var _92_v2 _dafny.Int
			_92_v2 = interface{}(_compr_26).(_dafny.Int)
			if ((_dafny.IntOfInt64(812)).Cmp(_92_v2) <= 0) && ((_92_v2).Cmp(_dafny.IntOfInt64(557)) < 0) {
				_coll26.Add((_92_v2).Times(_dafny.IntOfInt64(894)), true)
			}
		}
		return _coll26.ToMap()
	}())))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	var _source5 D16 = Companion_D16_.Create_DC37_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(610), true))
	_ = _source5
	if _source5.Is_DC38() {
		var _93___mcc_h0 _dafny.Int = _source5.Get_().(D16_DC38).Cf69
		_ = _93___mcc_h0
		var _94___mcc_h1 _dafny.Int = _source5.Get_().(D16_DC38).Cf70
		_ = _94___mcc_h1
		var _95___mcc_h2 bool = _source5.Get_().(D16_DC38).Cf71
		_ = _95___mcc_h2
		var _96_cf71 bool = _95___mcc_h2
		_ = _96_cf71
		var _97_cf70 _dafny.Int = _94___mcc_h1
		_ = _97_cf70
		var _98_cf69 _dafny.Int = _93___mcc_h0
		_ = _98_cf69
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_96_cf71)).Cardinality()))).Cardinality()), _96_cf71)
	} else if _source5.Is_DC37() {
		var _99___mcc_h3 _dafny.Map = _source5.Get_().(D16_DC37).Cf68
		_ = _99___mcc_h3
		var _100_cf68 _dafny.Map = _99___mcc_h3
		_ = _100_cf68
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vtfmeu"), false)).Cardinality(), true)).Merge(func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter27 := _dafny.Iterate((_100_cf68).Keys().Elements()); ; {
				_compr_27, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _101_v0 _dafny.Int
				_101_v0 = interface{}(_compr_27).(_dafny.Int)
				if (_100_cf68).Contains(_101_v0) {
					_coll27.Add(Companion_Default___.SafeModuloInt(_101_v0, _dafny.IntOfInt64(-772)), true)
				}
			}
			return _coll27.ToMap()
		}())
	} else {
		var _102___mcc_h4 D16 = _source5.Get_().(D16_DC39).Cf72
		_ = _102___mcc_h4
		var _103_cf72 D16 = _102___mcc_h4
		_ = _103_cf72
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tqifsibhn")).Cardinality()), true)
	}
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.CodePoint('o'))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(460))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg30 _dafny.Int) interface{} {
			return coer30(arg30)
		}
	}(func(_104_i0 _dafny.Int) _dafny.CodePoint {
		return (Companion_D3_.Create_DC11_(_dafny.CodePoint('t'))).Dtor_cf22()
	}))))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((func() _dafny.MultiSet {
		if true {
			return _dafny.MultiSetOf(_dafny.MultiSetOf((func() _dafny.Set {
				var _coll28 = _dafny.NewBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.IntOfInt64(536))).Keys().Elements()); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _105_v0 _dafny.CodePoint
					_105_v0 = interface{}(_compr_28).(_dafny.CodePoint)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.IntOfInt64(536))).Contains(_105_v0) {
						_coll28.Add(_105_v0)
					}
				}
				return _coll28.ToSet()
			}()).Cardinality(), (_dafny.SetOf(_dafny.IntOfInt64(923))).Cardinality()))
		}
		return _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(575), _dafny.IntOfInt64(305), _dafny.IntOfInt64(46))), _dafny.MultiSetOf(_dafny.IntOfInt64(-264), _dafny.IntOfInt64(693)), _dafny.MultiSetOf(_dafny.IntOfInt64(632)), _dafny.MultiSetOf(_dafny.IntOfInt64(170)))
	})()).Difference(_dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(178), !(false))).Cardinality())), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm50(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfInt64(-384)))).Difference((_dafny.MultiSetOf(func() _dafny.Set {
		var _coll29 = _dafny.NewBuilder()
		_ = _coll29
		for _iter29 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(339), _dafny.IntOfInt64(982))).Elements()); ; {
			_compr_29, _ok29 := _iter29()
			if !_ok29 {
				break
			}
			var _106_v0 _dafny.Int
			_106_v0 = interface{}(_compr_29).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(339), _dafny.IntOfInt64(982)), _106_v0) {
				_coll29.Add((_106_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
					var _coll30 = _dafny.NewMapBuilder()
					_ = _coll30
					for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-368), _dafny.IntOfInt64(100))); ; {
						_compr_30, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _107_v1 _dafny.Int
						_107_v1 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(-368)).Cmp(_107_v1) <= 0) && ((_107_v1).Cmp(_dafny.IntOfInt64(100)) < 0) {
							_coll30.Add(Companion_Default___.SafeDivisionInt(_107_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-175))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg31 _dafny.Int) interface{} {
									return coer31(arg31)
								}
							}(func(_108_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('i')
							}))).Cardinality()), (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(932), _dafny.MultiSetOf(_dafny.IntOfInt64(733)))).Cardinality())).Cardinality())).Cardinality()), _dafny.IntOfInt64(988))
						}
					}
					return _coll30.ToMap()
				}()).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xnb")).Cardinality()))).Cardinality())))
			}
		}
		return _coll29.ToSet()
	}())).Intersection(_dafny.MultiSetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(246))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Int {
		if false {
			return _dafny.IntOfInt64(443)
		}
		return _dafny.IntOfInt64(851)
	})()), _dafny.IntOfInt64(571), _dafny.IntOfInt64(829))
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(!(false))).Union(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Set, p3 _dafny.Map, globalState *GlobalState) D6 {
	var _source6 D10 = Companion_D10_.Create_DC23_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((func() _dafny.Map {
		var _coll31 = _dafny.NewMapBuilder()
		_ = _coll31
		for _iter31 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-916))).Cardinality(), _dafny.IntOfInt64(671)), _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), func() _dafny.Set {
			var _coll32 = _dafny.NewBuilder()
			_ = _coll32
			for _iter32 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Elements()); ; {
				_compr_32, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _109_v1 _dafny.Int
				_109_v1 = interface{}(_compr_32).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _109_v1) {
					_coll32.Add(Companion_Default___.SafeModuloInt(_109_v1, (_dafny.SetOf(!(false))).Cardinality()))
				}
			}
			return _coll32.ToSet()
		}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())))).Elements()); ; {
			_compr_31, _ok31 := _iter31()
			if !_ok31 {
				break
			}
			var _110_v0 _dafny.Set
			_110_v0 = interface{}(_compr_31).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-916))).Cardinality(), _dafny.IntOfInt64(671)), _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), func() _dafny.Set {
				var _coll33 = _dafny.NewBuilder()
				_ = _coll33
				for _iter33 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Elements()); ; {
					_compr_33, _ok33 := _iter33()
					if !_ok33 {
						break
					}
					var _111_v1 _dafny.Int
					_111_v1 = interface{}(_compr_33).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _111_v1) {
						_coll33.Add(Companion_Default___.SafeModuloInt(_111_v1, (_dafny.SetOf(!(false))).Cardinality()))
					}
				}
				return _coll33.ToSet()
			}(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()))), _110_v0) {
				_coll31.Add(_110_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(64))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_112_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality())))
			}
		}
		return _coll31.ToMap()
	}()).Cardinality(), _dafny.IntOfInt64(660)), _dafny.SeqOf(Companion_D2_.Create_DC6_(func() _dafny.Map {
		var _coll34 = _dafny.NewMapBuilder()
		_ = _coll34
		for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(874), _dafny.IntOfInt64(495))); ; {
			_compr_34, _ok34 := _iter34()
			if !_ok34 {
				break
			}
			var _113_v2 _dafny.Int
			_113_v2 = interface{}(_compr_34).(_dafny.Int)
			if ((_dafny.IntOfInt64(874)).Cmp(_113_v2) <= 0) && ((_113_v2).Cmp(_dafny.IntOfInt64(495)) < 0) {
				_coll34.Add((_113_v2).Times(_dafny.IntOfInt64(-116)), _dafny.IntOfInt64(823))
			}
		}
		return _coll34.ToMap()
	}()), Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(461), _dafny.IntOfInt64(387))))))
	_ = _source6
	if _source6.Is_DC24() {
		var _114___mcc_h0 _dafny.Sequence = _source6.Get_().(D10_DC24).Cf44
		_ = _114___mcc_h0
		var _115___mcc_h1 _dafny.Sequence = _source6.Get_().(D10_DC24).Cf45
		_ = _115___mcc_h1
		var _116___mcc_h2 bool = _source6.Get_().(D10_DC24).Cf46
		_ = _116___mcc_h2
		var _117_cf46 bool = _116___mcc_h2
		_ = _117_cf46
		var _118_cf45 _dafny.Sequence = _115___mcc_h1
		_ = _118_cf45
		var _119_cf44 _dafny.Sequence = _114___mcc_h0
		_ = _119_cf44
		return Companion_D6_.Create_DC17_(_dafny.IntOfInt64(-489), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_cf46, false)).Cardinality(), _dafny.IntOfInt64(973), _dafny.IntOfInt64(287), _dafny.IntOfInt64(272))
	} else if _source6.Is_DC25() {
		return Companion_D6_.Create_DC17_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tu")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), _dafny.IntOfInt64(-503), _dafny.IntOfInt64(26), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("grxj")).Cardinality())))
	} else {
		var _120___mcc_h3 _dafny.Map = _source6.Get_().(D10_DC23).Cf43
		_ = _120___mcc_h3
		var _121_cf43 _dafny.Map = _120___mcc_h3
		_ = _121_cf43
		if false {
			return Companion_D6_.Create_DC17_(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(-178), _dafny.IntOfInt64(-617), _dafny.IntOfInt64(884), (func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter35 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
					var _coll36 = _dafny.NewMapBuilder()
					_ = _coll36
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(656), _dafny.IntOfInt64(744))); ; {
						_compr_36, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _122_v4 _dafny.Int
						_122_v4 = interface{}(_compr_36).(_dafny.Int)
						if ((_dafny.IntOfInt64(656)).Cmp(_122_v4) <= 0) && ((_122_v4).Cmp(_dafny.IntOfInt64(744)) < 0) {
							_coll36.Add((_122_v4).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))), false)
						}
					}
					return _coll36.ToMap()
				}())).Elements()); ; {
					_compr_35, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _123_v3 _dafny.Map
					_123_v3 = interface{}(_compr_35).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
						var _coll37 = _dafny.NewMapBuilder()
						_ = _coll37
						for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(656), _dafny.IntOfInt64(744))); ; {
							_compr_37, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _122_v4 _dafny.Int
							_122_v4 = interface{}(_compr_37).(_dafny.Int)
							if ((_dafny.IntOfInt64(656)).Cmp(_122_v4) <= 0) && ((_122_v4).Cmp(_dafny.IntOfInt64(744)) < 0) {
								_coll37.Add((_122_v4).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))), false)
							}
						}
						return _coll37.ToMap()
					}()), _123_v3) {
						_coll35.Add(_123_v3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-843), !(false))).Cardinality(), _dafny.IntOfInt64(538)))
					}
				}
				return _coll35.ToMap()
			}()).Cardinality())
		} else {
			return Companion_D6_.Create_DC17_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(264))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_124_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			}))).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vghjbcr")).Cardinality()))), _dafny.IntOfInt64(-987), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rfqfeded")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('g'))).Cardinality()))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm54(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D16 {
	return Companion_D16_.Create_DC38_(((_dafny.Zero).Minus((func() _dafny.Map {
		var _coll38 = _dafny.NewMapBuilder()
		_ = _coll38
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-354), _dafny.IntOfInt64(574))); ; {
			_compr_38, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _125_v0 _dafny.Int
			_125_v0 = interface{}(_compr_38).(_dafny.Int)
			if ((_dafny.IntOfInt64(-354)).Cmp(_125_v0) <= 0) && ((_125_v0).Cmp(_dafny.IntOfInt64(574)) < 0) {
				_coll38.Add(Companion_Default___.SafeModuloInt(_125_v0, _dafny.IntOfInt64(-173)), false)
			}
		}
		return _coll38.ToMap()
	}()).Cardinality())).Minus(_dafny.IntOfInt64(443)), (func() _dafny.Int {
		if true {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(432), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(614), (_dafny.Zero).Minus((func() _dafny.Map {
				var _coll39 = _dafny.NewMapBuilder()
				_ = _coll39
				for _iter39 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(746), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wwtt")).Cardinality())))).Elements()); ; {
					_compr_39, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _126_v1 _dafny.Int
					_126_v1 = interface{}(_compr_39).(_dafny.Int)
					if (_dafny.MultiSetOf(_dafny.IntOfInt64(746), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wwtt")).Cardinality())))).Contains(_126_v1) {
						_coll39.Add((_126_v1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(730))).Cardinality())), _dafny.SeqOf(true))
					}
				}
				return _coll39.ToMap()
			}()).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()))).Cardinality())), _dafny.SeqOf(true, true, true))).Cardinality()
		}
		return (_dafny.Zero).Minus(_dafny.IntOfInt64(-784))
	})(), false)
}
func (_static *CompanionStruct_Default___) Fm55(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Set {
		var _coll40 = _dafny.NewBuilder()
		_ = _coll40
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(284), _dafny.IntOfInt64(-98))); ; {
			_compr_40, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _127_v0 _dafny.Int
			_127_v0 = interface{}(_compr_40).(_dafny.Int)
			if ((_dafny.IntOfInt64(284)).Cmp(_127_v0) <= 0) && ((_127_v0).Cmp(_dafny.IntOfInt64(-98)) < 0) {
				_coll40.Add(Companion_Default___.SafeModuloInt(_127_v0, _dafny.IntOfInt64(210)))
			}
		}
		return _coll40.ToSet()
	}()).Union(func() _dafny.Set {
		var _coll41 = _dafny.NewBuilder()
		_ = _coll41
		for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(314), _dafny.IntOfInt64(577))); ; {
			_compr_41, _ok41 := _iter41()
			if !_ok41 {
				break
			}
			var _128_v1 _dafny.Int
			_128_v1 = interface{}(_compr_41).(_dafny.Int)
			if ((_dafny.IntOfInt64(314)).Cmp(_128_v1) <= 0) && ((_128_v1).Cmp(_dafny.IntOfInt64(577)) < 0) {
				_coll41.Add(Companion_Default___.SafeDivisionInt(_128_v1, _dafny.IntOfInt64(-287)))
			}
		}
		return _coll41.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm56(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC11_((func() _dafny.CodePoint {
		if !(true) {
			return _dafny.CodePoint('j')
		}
		return _dafny.CodePoint('d')
	})())
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.CodePoint, _dafny.Sequence, bool, _dafny.Map) {
	var r0 _dafny.CodePoint = _dafny.CodePoint('D')
	_ = r0
	var r1 _dafny.Sequence = _dafny.EmptySeq
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 _dafny.Map = _dafny.EmptyMap
	_ = r3
	var _129_v0 bool
	_ = _129_v0
	_129_v0 = true
	var _130_v1 _dafny.Sequence
	_ = _130_v1
	_130_v1 = _dafny.SeqOf(_129_v0)
	var _131_v2 _dafny.Int
	_ = _131_v2
	_131_v2 = _dafny.IntOfInt64(762)
	if (_130_v1).Select((Companion_Default___.SafeIndex(_131_v2, _dafny.IntOfUint32((_130_v1).Cardinality()))).Uint32()).(bool) {
		var _132_v3 _dafny.MultiSet
		_ = _132_v3
		_132_v3 = _dafny.MultiSetOf(Companion_Default___.Fm39(globalState))
		var _133_v4 _dafny.Map
		_ = _133_v4
		_133_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v3, _129_v0)
		var _134_v5 D0
		_ = _134_v5
		_134_v5 = Companion_D0_.Create_DC0_(_129_v0)
		var _135_v6 _dafny.Sequence
		_ = _135_v6
		_135_v6 = _dafny.SeqOf(_134_v5, _134_v5)
		var _136_v7 D0
		_ = _136_v7
		_136_v7 = Companion_D0_.Create_DC2_((_135_v6).Select((Companion_Default___.SafeIndex(_131_v2, _dafny.IntOfUint32((_135_v6).Cardinality()))).Uint32()).(D0))
		_133_v4 = (_133_v4).Update(Companion_Default___.Fm10(_131_v2, _131_v2, _136_v7, globalState), _129_v0)
		(globalState).F15 = !(_129_v0)
		var _137_v8 _dafny.Array
		_ = _137_v8
		var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
		_ = _nw0
		_137_v8 = _nw0
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_137_v8), 0))
		_ = _index0
		(_137_v8).ArraySet1(Companion_Default___.SafeDivisionInt(_131_v2, _131_v2), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_137_v8), 0))
		_ = _index1
		(_137_v8).ArraySet1(Companion_Default___.Fm39(globalState), (_index1).Int())
		(globalState).F22 = _129_v0
		(globalState).F12 = _dafny.IntOfUint32((Companion_Default___.Fm22((_137_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_137_v8), 0))).Int()).(_dafny.Int), _129_v0, globalState)).Cardinality())
	} else {
		var _138_v9 _dafny.Map
		_ = _138_v9
		_138_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v2, true)
		var _139_v10 _dafny.Map
		_ = _139_v10
		_139_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, !(_129_v0))
		var _140_v11 _dafny.Map
		_ = _140_v11
		_140_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _131_v2)
		var _141_v12 _dafny.Sequence
		_ = _141_v12
		_141_v12 = _dafny.SeqOf((func() _dafny.Int {
			if (_140_v11).Contains(true) {
				return (_140_v11).Get(true).(_dafny.Int)
			}
			return _131_v2
		})())
		var _142_v13 _dafny.Sequence
		_ = _142_v13
		_142_v13 = _dafny.SeqOf(_141_v12, _141_v12, _dafny.SeqOf(_dafny.IntOfUint32((_130_v1).Cardinality()), _131_v2), _141_v12, _141_v12)
		var _143_v14 *C2
		_ = _143_v14
		var _nw1 *C2 = New_C2_()
		_ = _nw1
		_nw1.Ctor__(_142_v13, _129_v0, _129_v0)
		_143_v14 = _nw1
		var _144_v15 _dafny.Map
		_ = _144_v15
		_144_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_129_v0) == ((func() bool {
			if (_138_v9).Contains((_139_v10).Cardinality()) {
				return (_138_v9).Get((_139_v10).Cardinality()).(bool)
			}
			return _129_v0
		})()), _143_v14)
		var _145_v16 _dafny.MultiSet
		_ = _145_v16
		_145_v16 = _dafny.MultiSetOf(_129_v0, _129_v0)
		var _146_v17 *C6
		_ = _146_v17
		var _nw2 *C6 = New_C6_()
		_ = _nw2
		_nw2.Ctor__((_145_v16).Cardinality())
		_146_v17 = _nw2
		var _147_v18 _dafny.Map
		_ = _147_v18
		_147_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _146_v17)
		var _148_v19 D0
		_ = _148_v19
		_148_v19 = Companion_D0_.Create_DC1_(true, _129_v0, (_141_v12).Select((Companion_Default___.SafeIndex((_147_v18).Cardinality(), _dafny.IntOfUint32((_141_v12).Cardinality()))).Uint32()).(_dafny.Int), !(_129_v0))
		_144_v15 = (_144_v15).Update((_148_v19).Dtor_cf1(), _143_v14)
		(globalState).F2 = _129_v0
		var _149_v20 _dafny.Array
		_ = _149_v20
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_0
		var _nw3 _dafny.Array
		_ = _nw3
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw3 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) bool = func(_150_i0 _dafny.Int) bool {
				return true
			}
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw3 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw3).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw3).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_149_v20 = _nw3
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))
		_ = _index2
		(_149_v20).ArraySet1(_129_v0, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))
		_ = _index3
		(_149_v20).ArraySet1(false, (_index3).Int())
		var _151_v21 _dafny.Array
		_ = _151_v21
		var _nwElement0_0 _dafny.MultiSet = _145_v16
		_ = _nwElement0_0
		var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(15))
		_ = _nw4
		(_nw4).ArraySet1(_nwElement0_0, 0)
		(_nw4).ArraySet1(_145_v16, 1)
		(_nw4).ArraySet1((_145_v16).Union(_145_v16), 2)
		(_nw4).ArraySet1(_dafny.MultiSetFromSeq(_130_v1), 3)
		(_nw4).ArraySet1((_dafny.MultiSetOf(false)).Union((_145_v16).Update(_129_v0, Companion_Default___.Abs((_dafny.Zero).Minus(_131_v2)))), 4)
		(_nw4).ArraySet1(_145_v16, 5)
		(_nw4).ArraySet1(_dafny.MultiSetOf((_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)), 6)
		(_nw4).ArraySet1(Companion_Default___.Fm9(true, _146_v17.F40, _129_v0, globalState), 7)
		(_nw4).ArraySet1(Companion_Default___.Fm21(_146_v17.F40, globalState), 8)
		(_nw4).ArraySet1(_145_v16, 9)
		(_nw4).ArraySet1(_145_v16, 10)
		(_nw4).ArraySet1((_145_v16).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf((_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)))), 11)
		(_nw4).ArraySet1(_dafny.MultiSetOf(_129_v0, (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)), 12)
		(_nw4).ArraySet1((_145_v16).Difference(Companion_Default___.Fm9(_129_v0, _131_v2, (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool), globalState)), 13)
		(_nw4).ArraySet1(_145_v16, 14)
		_151_v21 = _nw4
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_151_v21), 0))
		_ = _index4
		(_151_v21).ArraySet1(_145_v16, (_index4).Int())
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(416), _dafny.ArrayLen((_151_v21), 0))
		_ = _index5
		(_151_v21).ArraySet1(_145_v16, (_index5).Int())
		var _152_v22 _dafny.Sequence
		_ = _152_v22
		_152_v22 = _dafny.SeqOf(p0)
		if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_152_v22, _dafny.SeqOf(p0))).Cardinality())).Cmp(_dafny.IntOfInt64(365)) == 0 {
			(globalState).F16 = (_dafny.IntOfInt64(200)).Plus(_131_v2)
			var _153_v23 D10
			_ = _153_v23
			_153_v23 = Companion_D10_.Create_DC25_()
			_153_v23 = Companion_D10_.Create_DC25_()
			var _154_v24 _dafny.Array
			_ = _154_v24
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_1
			var _nw5 _dafny.Array
			_ = _nw5
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw5 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = (func(_155_v17 *C6, _156_v2 _dafny.Int, _157_v20 _dafny.Array, _158_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_159_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_159_i1, (_dafny.SetOf(_155_v17.F40, (_dafny.Zero).Minus(_156_v2), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_157_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_157_v20), 0))).Int()).(bool), _158_v0)).Cardinality())).Cardinality())
					}
				})(_146_v17, _131_v2, _149_v20, _129_v0)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw5 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw5).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw5).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_154_v24 = _nw5
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_154_v24), 0))
			_ = _index6
			(_154_v24).ArraySet1(_146_v17.F40, (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_154_v24), 0))
			_ = _index7
			(_154_v24).ArraySet1(Companion_Default___.Fm39(globalState), (_index7).Int())
			var _160_v25 D9
			_ = _160_v25
			_160_v25 = Companion_D9_.Create_DC22_((_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool))
			var _161_v26 _dafny.Map
			_ = _161_v26
			_161_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_v1, _149_v20)
			var _162_v27 D9
			_ = _162_v27
			_162_v27 = Companion_D9_.Create_DC21_(_161_v26)
			var _163_v28 _dafny.Sequence
			_ = _163_v28
			_163_v28 = _dafny.SeqOf(_162_v27)
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_154_v24), 0))
			_ = _index8
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_154_v24), 0))
			_ = _index9
			var _rhs0 _dafny.Int = _146_v17.F40
			_ = _rhs0
			var _rhs1 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_141_v12, _141_v12)).Cardinality())
			_ = _rhs1
			var _rhs2 bool = (_160_v25).Dtor_cf42()
			_ = _rhs2
			var _rhs3 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool) {
					return _163_v28
				}
				return _163_v28
			})()).Cardinality())
			_ = _rhs3
			var _lhs0 *C6 = _146_v17
			_ = _lhs0
			var _lhs1 _dafny.Array = _154_v24
			_ = _lhs1
			var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_154_v24), 0))
			_ = _lhs2
			var _lhs3 _dafny.Array = _154_v24
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_154_v24), 0))
			_ = _lhs4
			_lhs0.F40 = _rhs0
			(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
			_129_v0 = _rhs2
			(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
			(globalState).F2 = !_dafny.Companion_Sequence_.Contains(_130_v1, (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool))
			_138_v9 = (_138_v9).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_146_v17.F40, (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)))
		} else {
			var _164_v29 _dafny.Set
			_ = _164_v29
			_164_v29 = _dafny.SetOf((_140_v11).Cardinality())
			var _165_v30 D9
			_ = _165_v30
			_165_v30 = Companion_D9_.Create_DC22_((func() bool {
				if (_143_v14).Fm16(_dafny.IntOfUint32((_152_v22).Cardinality()), _dafny.SeqOf(Companion_D2_.Create_DC9_(_131_v2, (_164_v29).Cardinality(), _129_v0)), _146_v17.F40, _146_v17.F40, globalState) {
					return (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)
				}
				return _129_v0
			})())
			_165_v30 = _165_v30
			r1 = _130_v1
			var _166_v31 D2
			_ = _166_v31
			_166_v31 = Companion_D2_.Create_DC9_(_146_v17.F40, _131_v2, (func() bool {
				if (_138_v9).Contains(_146_v17.F40) {
					return (_138_v9).Get(_146_v17.F40).(bool)
				}
				return (_149_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_149_v20), 0))).Int()).(bool)
			})())
			var _167_v32 _dafny.Sequence
			_ = _167_v32
			_167_v32 = _dafny.SeqOf(_166_v31)
			var _168_v33 _dafny.Map
			_ = _168_v33
			_168_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v2, _146_v17.F40)
			(globalState).F12 = (func() _dafny.Int {
				if (func() bool {
					if _129_v0 {
						return _129_v0
					}
					return (_143_v14).Fm16(_146_v17.F40, _167_v32, (_168_v33).Cardinality(), _146_v17.F40, globalState)
				})() {
					return _dafny.IntOfInt64(622)
				}
				return Companion_Default___.SafeModuloInt((_143_v14).Fm1(globalState), (_dafny.MultiSetOf((_146_v17).Fm2(_129_v0, _131_v2, globalState))).Cardinality())
			})()
			(globalState).F16 = (_146_v17).Fm1(globalState)
			(globalState).F22 = _dafny.Companion_Sequence_.Contains(_152_v22, p0)
		}
	}
	var _169_v34 _dafny.Array
	_ = _169_v34
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_2
	var _nw6 _dafny.Array
	_ = _nw6
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw6 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Set = (func(_170_v1 _dafny.Sequence, _171_v2 _dafny.Int, _172_v0 bool) func(_dafny.Int) _dafny.Set {
			return func(_173_i2 _dafny.Int) _dafny.Set {
				return (_dafny.SetOf(true, (_170_v1).Select((Companion_Default___.SafeIndex(_171_v2, _dafny.IntOfUint32((_170_v1).Cardinality()))).Uint32()).(bool))).Intersection(_dafny.SetOf(true, _172_v0))
			}
		})(_130_v1, _131_v2, _129_v0)
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
	_169_v34 = _nw6
	var _174_v35 _dafny.Set
	_ = _174_v35
	_174_v35 = _dafny.SetOf(_129_v0)
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_169_v34), 0))
	_ = _index10
	(_169_v34).ArraySet1((_174_v35).Union(_174_v35), (_index10).Int())
	var _175_v36 _dafny.Sequence
	_ = _175_v36
	_175_v36 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _131_v2))
	var _176_v37 D21
	_ = _176_v37
	_176_v37 = Companion_D21_.Create_DC52_(_dafny.IntOfUint32((_dafny.SeqOf(_129_v0)).Cardinality()), _175_v36)
	var _177_v38 _dafny.Map
	_ = _177_v38
	_177_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _176_v37)
	var _178_v39 _dafny.MultiSet
	_ = _178_v39
	_178_v39 = _dafny.MultiSetOf(_177_v38, _177_v38, _177_v38, (_177_v38).Update(_129_v0, Companion_D21_.Create_DC52_(_131_v2, _175_v36)))
	var _179_v40 _dafny.Map
	_ = _179_v40
	_179_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _174_v35)
	var _180_v41 _dafny.Sequence
	_ = _180_v41
	_180_v41 = _dafny.SeqOf(_177_v38)
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_169_v34), 0))
	_ = _index11
	var _rhs4 _dafny.Set = Companion_Default___.Fm18((func() _dafny.Set {
		if (_179_v40).Contains(_129_v0) {
			return (_179_v40).Get(_129_v0).(_dafny.Set)
		}
		return _174_v35
	})(), globalState)
	_ = _rhs4
	var _rhs5 _dafny.MultiSet = (_178_v39).Intersection(_dafny.MultiSetFromSeq(_180_v41))
	_ = _rhs5
	var _lhs5 _dafny.Array = _169_v34
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_169_v34), 0))
	_ = _lhs6
	(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
	_178_v39 = _rhs5
	var _181_v42 _dafny.Array
	_ = _181_v42
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
	_ = _nw7
	_181_v42 = _nw7
	var _182_v43 _dafny.MultiSet
	_ = _182_v43
	_182_v43 = _dafny.MultiSetOf(_131_v2)
	var _183_v44 _dafny.MultiSet
	_ = _183_v44
	_183_v44 = _dafny.MultiSetOf(_182_v43)
	var _rhs6 bool = _129_v0
	_ = _rhs6
	var _rhs7 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfInt64(812)).Minus(((func() _dafny.Int {
		if (_183_v44).Contains(_182_v43) {
			return (_183_v44).Multiplicity(_182_v43)
		}
		return _131_v2
	})()).Plus((Companion_Default___.Fm18(_174_v35, globalState)).Cardinality())))
	_ = _rhs7
	var _rhs8 _dafny.Int = (_131_v2).Plus(_131_v2)
	_ = _rhs8
	var _rhs9 _dafny.Array = _181_v42
	_ = _rhs9
	var _rhs10 _dafny.CodePoint = p0
	_ = _rhs10
	var _lhs7 *GlobalState = globalState
	_ = _lhs7
	var _lhs8 *GlobalState = globalState
	_ = _lhs8
	_lhs7.F2 = _rhs6
	_lhs8.F12 = _rhs7
	_131_v2 = _rhs8
	_181_v42 = _rhs9
	r0 = _rhs10
	(globalState).F16 = (_dafny.Zero).Minus(_131_v2)
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_169_v34), 0))
	_ = _index12
	(_169_v34).ArraySet1(Companion_Default___.Fm52(!(_129_v0) || (true), (_dafny.Zero).Minus((((_169_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(683), _dafny.ArrayLen((_169_v34), 0))).Int()).(_dafny.Set)).Cardinality()).Minus(_131_v2)), globalState), (_index12).Int())
	var _hi0 _dafny.Int = _131_v2
	_ = _hi0
	for _184_i3 := _131_v2; _184_i3.Cmp(_hi0) < 0; _184_i3 = _184_i3.Plus(_dafny.One) {
		(globalState).F2 = _129_v0
		_131_v2 = Companion_Default___.SafeModuloInt(_131_v2, (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(_184_i3)).Cardinality(), _131_v2)))
		var _185_v45 *C8
		_ = _185_v45
		var _nw8 *C8 = New_C8_()
		_ = _nw8
		_nw8.Ctor__(_182_v43)
		_185_v45 = _nw8
		var _186_v46 D25
		_ = _186_v46
		_186_v46 = Companion_D25_.Create_DC62_(_185_v45)
		var _source7 D25 = Companion_D25_.Create_DC64_(_186_v46)
		_ = _source7
		if _source7.Is_DC63() {
			var _187___mcc_h0 _dafny.Sequence = _source7.Get_().(D25_DC63).Cf116
			_ = _187___mcc_h0
			var _188___mcc_h1 bool = _source7.Get_().(D25_DC63).Cf117
			_ = _188___mcc_h1
			var _189___mcc_h2 _dafny.Int = _source7.Get_().(D25_DC63).Cf118
			_ = _189___mcc_h2
			var _190_cf118 _dafny.Int = _189___mcc_h2
			_ = _190_cf118
			var _191_cf117 bool = _188___mcc_h1
			_ = _191_cf117
			var _192_cf116 _dafny.Sequence = _187___mcc_h0
			_ = _192_cf116
			var _193_v47 D15
			_ = _193_v47
			_193_v47 = Companion_D15_.Create_DC35_(_129_v0)
			var _194_v48 *C5
			_ = _194_v48
			var _nw9 *C5 = New_C5_()
			_ = _nw9
			_nw9.Ctor__((_193_v47).Dtor_cf65(), _191_cf117)
			_194_v48 = _nw9
			var _195_v49 _dafny.Sequence
			_ = _195_v49
			_195_v49 = _dafny.SeqOf(_194_v48)
			var _196_v50 _dafny.Map
			_ = _196_v50
			_196_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v49, _129_v0)
			_196_v50 = (_196_v50).Update(_dafny.Companion_Sequence_.Update(_195_v49, (Companion_Default___.SafeIndex(_131_v2, _dafny.IntOfUint32((_195_v49).Cardinality()))).Uint32(), _194_v48), _129_v0)
			(globalState).F24 = (_129_v0) && (_129_v0)
			var _197_v51 _dafny.MultiSet
			_ = _197_v51
			_197_v51 = _dafny.MultiSetOf(true)
			var _198_v52 _dafny.Map
			_ = _198_v52
			_198_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_190_cf118).Plus(_184_i3), _197_v51)
			_198_v52 = (_198_v52).Update(_190_cf118, _dafny.MultiSetOf(_191_cf117, _129_v0))
			var _199_v53 _dafny.Map
			_ = _199_v53
			_199_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v43, _129_v0)
			(globalState).F15 = (_199_v53).Equals(_199_v53)
		} else if _source7.Is_DC62() {
			var _200___mcc_h3 *C8 = _source7.Get_().(D25_DC62).Cf115
			_ = _200___mcc_h3
			var _201_cf115 *C8 = _200___mcc_h3
			_ = _201_cf115
			var _202_v54 _dafny.Array
			_ = _202_v54
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw10
			_202_v54 = _nw10
			var _203_v55 D3
			_ = _203_v55
			_203_v55 = Companion_D3_.Create_DC10_(_202_v54)
			_202_v54 = (_203_v55).Dtor_cf21()
			(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm22(_184_i3, _129_v0, globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(219))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_204_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			})))
			var _205_v56 _dafny.Array
			_ = _205_v56
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_3
			var _nw11 _dafny.Array
			_ = _nw11
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw11 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_206_v0 bool) func(_dafny.Int) bool {
					return func(_207_i5 _dafny.Int) bool {
						return _206_v0
					}
				})(_129_v0)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw11 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw11).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw11).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_205_v56 = _nw11
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))
			_ = _index13
			(_205_v56).ArraySet1(_129_v0, (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))
			_ = _index14
			var _rhs11 bool = _129_v0
			_ = _rhs11
			var _rhs12 bool = _129_v0
			_ = _rhs12
			var _rhs13 bool = _129_v0
			_ = _rhs13
			var _rhs14 bool = _129_v0
			_ = _rhs14
			var _lhs9 _dafny.Array = _205_v56
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = globalState
			_ = _lhs11
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			var _lhs13 *GlobalState = globalState
			_ = _lhs13
			(_lhs9).ArraySet1(_rhs11, (_lhs10).Int())
			_lhs11.F2 = _rhs12
			_lhs12.F2 = _rhs13
			_lhs13.F15 = _rhs14
			var _208_v57 _dafny.MultiSet
			_ = _208_v57
			_208_v57 = _dafny.MultiSetOf(_202_v54, _202_v54)
			var _209_v58 _dafny.Map
			_ = _209_v58
			_209_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _208_v57)
			var _210_v59 D26
			_ = _210_v59
			_210_v59 = Companion_D26_.Create_DC65_(_208_v57)
			var _211_v60 _dafny.Map
			_ = _211_v60
			_211_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hmiq")).Cardinality()), ((_208_v57).Update(_202_v54, Companion_Default___.Abs(((func() _dafny.MultiSet {
				if (_209_v58).Contains((_205_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))).Int()).(bool)) {
					return (_209_v58).Get((_205_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))).Int()).(bool)).(_dafny.MultiSet)
				}
				return (_210_v59).Dtor_cf120()
			})()).Cardinality()))).Difference(_dafny.MultiSetOf(_202_v54, _202_v54, _202_v54)))
			var _212_v61 _dafny.Map
			_ = _212_v61
			_212_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_203_v55, _131_v2)
			var _213_v62 *C4
			_ = _213_v62
			var _nw12 *C4 = New_C4_()
			_ = _nw12
			_nw12.Ctor__(_212_v61, (_205_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_205_v56), 0))).Int()).(bool), _129_v0)
			_213_v62 = _nw12
			var _214_v63 D16
			_ = _214_v63
			_214_v63 = Companion_D16_.Create_DC38_((_dafny.Zero).Minus(_131_v2), _184_i3, _129_v0)
			_211_v60 = (_211_v60).Update(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v62, _129_v0)).Update(_213_v62, (_213_v62).Fm2(true, (_214_v63).Dtor_cf69(), globalState))).Cardinality(), _208_v57)
		} else {
			var _215___mcc_h4 D25 = _source7.Get_().(D25_DC64).Cf119
			_ = _215___mcc_h4
			var _216_cf119 D25 = _215___mcc_h4
			_ = _216_cf119
			var _217_v64 bool
			_ = _217_v64
			var _218_v65 bool
			_ = _218_v65
			var _out0 bool
			_ = _out0
			var _out1 bool
			_ = _out1
			_out0, _out1 = (_185_v45).M4(globalState)
			_217_v64 = _out0
			_218_v65 = _out1
			(globalState).F12 = _dafny.IntOfUint32((Companion_Default___.Fm22(_dafny.IntOfInt64(430), _218_v65, globalState)).Cardinality())
			(globalState).F16 = _184_i3
			_217_v64 = _217_v64
		}
		var _219_v66 _dafny.Array
		_ = _219_v66
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_4
		var _nw13 _dafny.Array
		_ = _nw13
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw13 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_220_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_221_i6 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_221_i6, _220_v2)
				}
			})(_131_v2)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw13 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw13).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw13).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_219_v66 = _nw13
		var _222_v67 _dafny.Map
		_ = _222_v67
		_222_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_129_v0, _dafny.IntOfInt64(668))
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_219_v66), 0))
		_ = _index15
		(_219_v66).ArraySet1((_131_v2).Plus((func() _dafny.Int {
			if (_222_v67).Contains(_129_v0) {
				return (_222_v67).Get(_129_v0).(_dafny.Int)
			}
			return _131_v2
		})()), (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_219_v66), 0))
		_ = _index16
		(_219_v66).ArraySet1(((_dafny.Zero).Minus(Companion_Default___.Fm39(globalState))).Plus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm39(globalState), _131_v2)), (_index16).Int())
	}
	r0 = p0
	r1 = Companion_Default___.Fm25(_129_v0, !(_129_v0) || (_129_v0), p0, globalState)
	r2 = (_131_v2).Cmp(_131_v2) > 0
	var _223_v68 _dafny.MultiSet
	_ = _223_v68
	_223_v68 = _dafny.MultiSetOf(true, _129_v0)
	r3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_223_v68).Union(_223_v68), !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_130_v1, (Companion_Default___.SafeIndex(_131_v2, _dafny.IntOfUint32((_130_v1).Cardinality()))).Uint32(), _129_v0), _130_v1)))
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _224_v0 _dafny.Int
	_ = _224_v0
	_224_v0 = _dafny.IntOfInt64(979)
	var _225_v1 _dafny.Sequence
	_ = _225_v1
	_225_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ovdhmu")
	var _226_v2 _dafny.Set
	_ = _226_v2
	_226_v2 = _dafny.SetOf(_224_v0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_225_v1).Cardinality())))
	var _227_v3 _dafny.MultiSet
	_ = _227_v3
	_227_v3 = _dafny.MultiSetOf(_225_v1)
	var _228_v4 _dafny.Array
	_ = _228_v4
	var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
	_ = _nw14
	_228_v4 = _nw14
	var _229_v5 _dafny.Array
	_ = _229_v5
	var _len0_5 _dafny.Int = _dafny.IntOfInt64(17)
	_ = _len0_5
	var _nw15 _dafny.Array
	_ = _nw15
	if _len0_5.Cmp(_dafny.Zero) == 0 {
		_nw15 = _dafny.NewArray(_len0_5)
	} else {
		var _init5 func(_dafny.Int) _dafny.Int = (func(_230_v2 _dafny.Set) func(_dafny.Int) _dafny.Int {
			return func(_231_i1 _dafny.Int) _dafny.Int {
				return (_231_i1).Minus((_230_v2).Cardinality())
			}
		})(_226_v2)
		_ = _init5
		var _element0_5 = _init5(_dafny.Zero)
		_ = _element0_5
		_nw15 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
		(_nw15).ArraySet1(_element0_5, 0)
		var _nativeLen0_5 = (_len0_5).Int()
		_ = _nativeLen0_5
		for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
			(_nw15).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
		}
	}
	_229_v5 = _nw15
	var _232_v6 _dafny.Map
	_ = _232_v6
	_232_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(434))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg35 _dafny.Int) interface{} {
			return coer35(arg35)
		}
	}(func(_233_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), _229_v5)
	var _234_v8 _dafny.CodePoint
	_ = _234_v8
	_234_v8 = _dafny.CodePoint('v')
	var _235_v9 _dafny.Set
	_ = _235_v9
	_235_v9 = _dafny.SetOf(_234_v8)
	var _236_v10 _dafny.Map
	_ = _236_v10
	_236_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v0, func() _dafny.Map {
		var _coll42 = _dafny.NewMapBuilder()
		_ = _coll42
		for _iter42 := _dafny.Iterate((_235_v9).Elements()); ; {
			_compr_42, _ok42 := _iter42()
			if !_ok42 {
				break
			}
			var _237_v7 _dafny.CodePoint
			_237_v7 = interface{}(_compr_42).(_dafny.CodePoint)
			if (_235_v9).Contains(_237_v7) {
				_coll42.Add(_237_v7, _dafny.IntOfInt64(-441))
			}
		}
		return _coll42.ToMap()
	}())
	var _238_v11 _dafny.Map
	_ = _238_v11
	_238_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v8, _224_v0)
	var _239_v12 _dafny.Map
	_ = _239_v12
	_239_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
		if (_236_v10).Contains(_224_v0) {
			return (_236_v10).Get(_224_v0).(_dafny.Map)
		}
		return _238_v11
	})(), _224_v0)
	var _240_globalState *GlobalState
	_ = _240_globalState
	var _nw16 *GlobalState = New_GlobalState_()
	_ = _nw16
	_nw16.Ctor__(_226_v2, _227_v3, true, true, _dafny.CodePoint('l'), _228_v4, _225_v1, _232_v6, false, true, _dafny.IntOfInt64(139), true, _dafny.IntOfInt64(66), _dafny.IntOfInt64(-610), _229_v5, false, _dafny.IntOfInt64(796), true, true, _dafny.CodePoint('w'), _225_v1, _dafny.IntOfInt64(942), true, _dafny.IntOfInt64(735), false, _239_v12)
	_240_globalState = _nw16
	var _241_v14 bool
	_ = _241_v14
	_241_v14 = false
	var _242_v15 _dafny.MultiSet
	_ = _242_v15
	_242_v15 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v14, _224_v0))
	var _243_v16 _dafny.Map
	_ = _243_v16
	_243_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _224_v0)
	var _244_v17 _dafny.Map
	_ = _244_v17
	_244_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_v16, _241_v14)
	(_240_globalState).F12 = ((func() _dafny.Map {
		var _coll43 = _dafny.NewMapBuilder()
		_ = _coll43
		for _iter43 := _dafny.Iterate((_242_v15).Elements()); ; {
			_compr_43, _ok43 := _iter43()
			if !_ok43 {
				break
			}
			var _245_v13 _dafny.Map
			_245_v13 = interface{}(_compr_43).(_dafny.Map)
			if (_242_v15).Contains(_245_v13) {
				_coll43.Add(_245_v13, (Companion_D0_.Create_DC1_(_241_v14, _241_v14, _224_v0, _241_v14)).Dtor_cf2())
			}
		}
		return _coll43.ToMap()
	}()).Merge(_244_v17)).Cardinality()
	if (_227_v3).IsProperSubsetOf(Companion_Default___.Fm0(false, _240_globalState)) {
		var _246_v18 _dafny.CodePoint
		_ = _246_v18
		var _247_v19 _dafny.Sequence
		_ = _247_v19
		var _248_v20 bool
		_ = _248_v20
		var _249_v21 _dafny.Map
		_ = _249_v21
		var _out2 _dafny.CodePoint
		_ = _out2
		var _out3 _dafny.Sequence
		_ = _out3
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Map
		_ = _out5
		_out2, _out3, _out4, _out5 = Companion_Default___.M0(_234_v8, _240_globalState)
		_246_v18 = _out2
		_247_v19 = _out3
		_248_v20 = _out4
		_249_v21 = _out5
		var _250_v22 D0
		_ = _250_v22
		_250_v22 = Companion_D0_.Create_DC0_(_248_v20)
		_250_v22 = _250_v22
		(_240_globalState).F12 = _224_v0
		var _251_v23 _dafny.MultiSet
		_ = _251_v23
		_251_v23 = _dafny.MultiSetOf(_dafny.IntOfInt64(882))
		var _252_v24 _dafny.Array
		_ = _252_v24
		var _nwElement0_1 _dafny.MultiSet = _dafny.MultiSetOf(_224_v0)
		_ = _nwElement0_1
		var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(2))
		_ = _nw17
		(_nw17).ArraySet1(_nwElement0_1, 0)
		(_nw17).ArraySet1(_251_v23, 1)
		_252_v24 = _nw17
		var _253_v25 D1
		_ = _253_v25
		_253_v25 = Companion_D1_.Create_DC3_(_247_v19)
		var _rhs15 bool = _241_v14
		_ = _rhs15
		var _rhs16 _dafny.Sequence = (_253_v25).Dtor_cf6()
		_ = _rhs16
		var _rhs17 _dafny.Array = _252_v24
		_ = _rhs17
		var _rhs18 _dafny.CodePoint = _246_v18
		_ = _rhs18
		var _lhs14 *GlobalState = _240_globalState
		_ = _lhs14
		_lhs14.F2 = _rhs15
		_247_v19 = _rhs16
		_252_v24 = _rhs17
		_246_v18 = _rhs18
		var _254_v26 _dafny.CodePoint
		_ = _254_v26
		var _255_v27 _dafny.Sequence
		_ = _255_v27
		var _256_v28 bool
		_ = _256_v28
		var _257_v29 _dafny.Map
		_ = _257_v29
		var _out6 _dafny.CodePoint
		_ = _out6
		var _out7 _dafny.Sequence
		_ = _out7
		var _out8 bool
		_ = _out8
		var _out9 _dafny.Map
		_ = _out9
		_out6, _out7, _out8, _out9 = Companion_Default___.M0(_234_v8, _240_globalState)
		_254_v26 = _out6
		_255_v27 = _out7
		_256_v28 = _out8
		_257_v29 = _out9
	} else {
		(_240_globalState).F15 = (_226_v2).IsDisjointFrom(_226_v2)
		(_240_globalState).F24 = (_241_v14) || (_241_v14)
		var _258_v30 *C3
		_ = _258_v30
		var _nw18 *C3 = New_C3_()
		_ = _nw18
		_nw18.Ctor__(_dafny.MultiSetOf(_dafny.SeqOf(_224_v0)), true, (_241_v14) && (_241_v14))
		_258_v30 = _nw18
		var _259_v31 _dafny.Map
		_ = _259_v31
		_259_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v14, true)
		var _260_v32 *C1
		_ = _260_v32
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__(_224_v0, (_259_v31).Cardinality(), _241_v14, false)
		_260_v32 = _nw19
		var _261_v33 _dafny.Map
		_ = _261_v33
		_261_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v32, _241_v14)
		var _262_v34 _dafny.Sequence
		_ = _262_v34
		_262_v34 = _dafny.SeqOf(_261_v33)
		var _263_v35 _dafny.MultiSet
		_ = _263_v35
		_263_v35 = _dafny.MultiSetOf(_dafny.IntOfUint32((_225_v1).Cardinality()))
		var _264_v36 T3
		_ = _264_v36
		var _nw20 *C8 = New_C8_()
		_ = _nw20
		_nw20.Ctor__(_263_v35)
		_264_v36 = _nw20
		var _265_v37 _dafny.MultiSet
		_ = _265_v37
		_265_v37 = _dafny.MultiSetOf(_264_v36, _264_v36, _264_v36, _264_v36)
		var _266_v38 _dafny.Map
		_ = _266_v38
		_266_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_224_v0).Plus(((_262_v34).Select((Companion_Default___.SafeIndex((_260_v32).F36(), _dafny.IntOfUint32((_262_v34).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), (_265_v37).Cardinality())
		_266_v38 = (_266_v38).Update(_224_v0, (_224_v0).Plus(_224_v0))
		var _267_v39 D3
		_ = _267_v39
		_267_v39 = Companion_D3_.Create_DC11_(_dafny.CodePoint('b'))
		var _268_v40 _dafny.Map
		_ = _268_v40
		_268_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_267_v39, _241_v14)
		_268_v40 = (_268_v40).Update(Companion_Default___.Fm56(_240_globalState), _241_v14)
	}
	var _269_v41 _dafny.Array
	_ = _269_v41
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(4)
	_ = _len0_6
	var _nw21 _dafny.Array
	_ = _nw21
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw21 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_270_v14 bool) func(_dafny.Int) bool {
			return func(_271_i2 _dafny.Int) bool {
				return _270_v14
			}
		})(_241_v14)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw21 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw21).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw21).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_269_v41 = _nw21
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
	_ = _index17
	(_269_v41).ArraySet1(_241_v14, (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
	_ = _index18
	(_269_v41).ArraySet1(_241_v14, (_index18).Int())
	var _272_v42 _dafny.Sequence
	_ = _272_v42
	_272_v42 = _dafny.SeqOf(_224_v0, _224_v0, _224_v0, _224_v0)
	var _273_v43 _dafny.Map
	_ = _273_v43
	_273_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v41, _241_v14)
	var _274_v44 _dafny.Map
	_ = _274_v44
	_274_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
	var _275_v45 _dafny.Sequence
	_ = _275_v45
	_275_v45 = _dafny.SeqOf(_241_v14)
	var _276_v46 _dafny.MultiSet
	_ = _276_v46
	_276_v46 = _dafny.MultiSetOf(_dafny.IntOfUint32((_275_v45).Cardinality()), _224_v0, _224_v0)
	var _277_v47 T2
	_ = _277_v47
	var _nw22 *C9 = New_C9_()
	_ = _nw22
	_nw22.Ctor__(_241_v14, (_274_v44).Cardinality(), !((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool)), !(_241_v14), _276_v46)
	_277_v47 = _nw22
	var _278_v48 _dafny.Sequence
	_ = _278_v48
	_278_v48 = _dafny.SeqOf(_277_v47)
	var _279_v49 _dafny.MultiSet
	_ = _279_v49
	_279_v49 = _dafny.MultiSetOf((_273_v43).Cardinality(), _224_v0, _dafny.IntOfUint32((_278_v48).Cardinality()))
	var _280_v50 *C8
	_ = _280_v50
	var _nw23 *C8 = New_C8_()
	_ = _nw23
	_nw23.Ctor__(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_272_v42, (Companion_Default___.SafeIndex((_279_v49).Cardinality(), _dafny.IntOfUint32((_272_v42).Cardinality()))).Uint32(), _dafny.IntOfInt64(-790))).Cardinality())))
	_280_v50 = _nw23
	var _281_v51 _dafny.Array
	_ = _281_v51
	var _nwElement0_2 *C8 = _280_v50
	_ = _nwElement0_2
	var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(2))
	_ = _nw24
	(_nw24).ArraySet1(_nwElement0_2, 0)
	(_nw24).ArraySet1(_280_v50, 1)
	_281_v51 = _nw24
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_281_v51), 0))
	_ = _index19
	(_281_v51).ArraySet1(_280_v50, (_index19).Int())
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_281_v51), 0))
	_ = _index20
	(_281_v51).ArraySet1(_280_v50, (_index20).Int())
	var _282_v52 _dafny.Map
	_ = _282_v52
	_282_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v0, _224_v0)
	var _283_v55 _dafny.Array
	_ = _283_v55
	var _nwElement0_3 _dafny.Map = _282_v52
	_ = _nwElement0_3
	var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(23))
	_ = _nw25
	(_nw25).ArraySet1(_nwElement0_3, 0)
	(_nw25).ArraySet1(_282_v52, 1)
	(_nw25).ArraySet1(_282_v52, 2)
	(_nw25).ArraySet1(_282_v52, 3)
	(_nw25).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_241_v14, _241_v14)).Cardinality(), _224_v0), 4)
	(_nw25).ArraySet1(_282_v52, 5)
	(_nw25).ArraySet1(_282_v52, 6)
	(_nw25).ArraySet1(func() _dafny.Map {
		var _coll44 = _dafny.NewMapBuilder()
		_ = _coll44
		for _iter44 := _dafny.Iterate((_272_v42).Elements()); ; {
			_compr_44, _ok44 := _iter44()
			if !_ok44 {
				break
			}
			var _284_v53 _dafny.Int
			_284_v53 = interface{}(_compr_44).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_272_v42, _284_v53) {
				_coll44.Add((_284_v53).Plus((func() _dafny.Int {
					if (_282_v52).Contains(_dafny.IntOfInt64(165)) {
						return (_282_v52).Get(_dafny.IntOfInt64(165)).(_dafny.Int)
					}
					return _224_v0
				})()), _224_v0)
			}
		}
		return _coll44.ToMap()
	}(), 7)
	(_nw25).ArraySet1((_282_v52).Update(_dafny.IntOfUint32((_225_v1).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_277_v47).F27(), _241_v14)).Cardinality())), 8)
	(_nw25).ArraySet1(_282_v52, 9)
	(_nw25).ArraySet1((_282_v52).Update(_224_v0, _dafny.IntOfUint32((_225_v1).Cardinality())), 10)
	(_nw25).ArraySet1(func() _dafny.Map {
		var _coll45 = _dafny.NewMapBuilder()
		_ = _coll45
		for _iter45 := _dafny.Iterate((_279_v49).Elements()); ; {
			_compr_45, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _285_v54 _dafny.Int
			_285_v54 = interface{}(_compr_45).(_dafny.Int)
			if (_279_v49).Contains(_285_v54) {
				_coll45.Add((_285_v54).Times((_243_v16).Cardinality()), _224_v0)
			}
		}
		return _coll45.ToMap()
	}(), 11)
	(_nw25).ArraySet1(_282_v52, 12)
	(_nw25).ArraySet1(_282_v52, 13)
	(_nw25).ArraySet1(_282_v52, 14)
	(_nw25).ArraySet1(_282_v52, 15)
	(_nw25).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("etqlwqsyc")).Cardinality()), _dafny.IntOfUint32((_275_v45).Cardinality())), 16)
	(_nw25).ArraySet1(_282_v52, 17)
	(_nw25).ArraySet1(_282_v52, 18)
	(_nw25).ArraySet1(_282_v52, 19)
	(_nw25).ArraySet1(_282_v52, 20)
	(_nw25).ArraySet1(_282_v52, 21)
	(_nw25).ArraySet1(_282_v52, 22)
	_283_v55 = _nw25
	var _286_v56 D6
	_ = _286_v56
	_286_v56 = Companion_D6_.Create_DC16_(_225_v1)
	var _287_v57 _dafny.Map
	_ = _287_v57
	_287_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v55, _dafny.IntOfUint32(((func() _dafny.Sequence {
		if (_277_v47).F27() {
			return (_286_v56).Dtor_cf29()
		}
		return _225_v1
	})()).Cardinality()))
	_224_v0 = (func() _dafny.Int {
		if (_287_v57).Contains(_283_v55) {
			return (_287_v57).Get(_283_v55).(_dafny.Int)
		}
		return _224_v0
	})()
	var _288_v58 T0
	_ = _288_v58
	var _nw26 *C6 = New_C6_()
	_ = _nw26
	_nw26.Ctor__(_224_v0)
	_288_v58 = _nw26
	var _289_v59 _dafny.CodePoint
	_ = _289_v59
	var _290_v60 _dafny.Int
	_ = _290_v60
	var _out10 _dafny.CodePoint
	_ = _out10
	var _out11 _dafny.Int
	_ = _out11
	_out10, _out11 = (_277_v47).M2((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool), _288_v58, (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool), _240_globalState)
	_289_v59 = _out10
	_290_v60 = _out11
	var _291_v61 _dafny.CodePoint
	_ = _291_v61
	var _292_v62 _dafny.Sequence
	_ = _292_v62
	var _293_v63 bool
	_ = _293_v63
	var _294_v64 _dafny.Map
	_ = _294_v64
	var _out12 _dafny.CodePoint
	_ = _out12
	var _out13 _dafny.Sequence
	_ = _out13
	var _out14 bool
	_ = _out14
	var _out15 _dafny.Map
	_ = _out15
	_out12, _out13, _out14, _out15 = Companion_Default___.M0(_289_v59, _240_globalState)
	_291_v61 = _out12
	_292_v62 = _out13
	_293_v63 = _out14
	_294_v64 = _out15
	var _295_v65 _dafny.MultiSet
	_ = _295_v65
	_295_v65 = _dafny.MultiSetOf((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
	if ((_dafny.MultiSetOf((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))).Difference(_295_v65)).IsDisjointFrom((_295_v65).Update(_293_v63, Companion_Default___.Abs(_290_v60))) {
		var _296_v66 _dafny.CodePoint
		_ = _296_v66
		var _297_v67 _dafny.Int
		_ = _297_v67
		var _out16 _dafny.CodePoint
		_ = _out16
		var _out17 _dafny.Int
		_ = _out17
		_out16, _out17 = (_277_v47).M2((_277_v47).F27(), _288_v58, (_277_v47).F27(), _240_globalState)
		_296_v66 = _out16
		_297_v67 = _out17
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_229_v5), 0))
		_ = _index21
		(_229_v5).ArraySet1(_224_v0, (_index21).Int())
		var _298_v68 _dafny.MultiSet
		_ = _298_v68
		_298_v68 = _dafny.MultiSetOf(_272_v42)
		var _299_v69 *C3
		_ = _299_v69
		var _nw27 *C3 = New_C3_()
		_ = _nw27
		_nw27.Ctor__((func() _dafny.MultiSet {
			if true {
				return _dafny.MultiSetOf(_272_v42, _dafny.SeqOf(_297_v67))
			}
			return _298_v68
		})(), (_288_v58).Fm2((_277_v47).F27(), _224_v0, _240_globalState), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
		_299_v69 = _nw27
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_229_v5), 0))
		_ = _index22
		var _rhs19 _dafny.Int = (_297_v67).Times((_224_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nhcidbx")).Cardinality())))
		_ = _rhs19
		var _rhs20 *C3 = _299_v69
		_ = _rhs20
		var _lhs15 _dafny.Array = _229_v5
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_229_v5), 0))
		_ = _lhs16
		(_lhs15).ArraySet1(_rhs19, (_lhs16).Int())
		_299_v69 = _rhs20
		_274_v44 = (_274_v44).Update(((_272_v42).Select((Companion_Default___.SafeIndex(_224_v0, _dafny.IntOfUint32((_272_v42).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_229_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_229_v5), 0))).Int()).(_dafny.Int))) == 0, (_277_v47).F26())
		var _300_v70 _dafny.Sequence
		_ = _300_v70
		_300_v70 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("oujjykne"), _225_v1)
		(_299_v69).M1((_277_v47).F26(), _dafny.Companion_Sequence_.Update(_300_v70, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_224_v0), _dafny.IntOfUint32((_300_v70).Cardinality()))).Uint32(), _225_v1), Companion_Default___.Fm25((_277_v47).F26(), (_277_v47).F26(), _296_v66, _240_globalState), _240_globalState)
		var _301_v71 _dafny.Map
		_ = _301_v71
		_301_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(370), _269_v41)
		_301_v71 = (_301_v71).Update((_229_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_229_v5), 0))).Int()).(_dafny.Int), _269_v41)
	} else {
		(_240_globalState).F2 = !((_277_v47).F26()) || ((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
		var _302_v72 _dafny.Map
		_ = _302_v72
		_302_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v60, _241_v14)
		var _303_v73 _dafny.Map
		_ = _303_v73
		_303_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v47).F26(), _302_v72)
		var _304_v75 _dafny.Set
		_ = _304_v75
		_304_v75 = _dafny.SetOf((_277_v47).F26(), (_277_v47).F26())
		var _305_v76 _dafny.Map
		_ = _305_v76
		_305_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_304_v75, (_277_v47).F26())
		var _306_v78 _dafny.Array
		_ = _306_v78
		var _nwElement0_4 _dafny.Map = _302_v72
		_ = _nwElement0_4
		var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(12))
		_ = _nw28
		(_nw28).ArraySet1(_nwElement0_4, 0)
		(_nw28).ArraySet1(_302_v72, 1)
		(_nw28).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v0, (_277_v47).F27()), 2)
		(_nw28).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_225_v1).Cardinality()), (_277_v47).F26())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_272_v42).Cardinality()), (_277_v47).Fm2((_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool), _290_v60, _240_globalState))), 3)
		(_nw28).ArraySet1((func() _dafny.Map {
			if (_303_v73).Contains(_293_v63) {
				return (_303_v73).Get(_293_v63).(_dafny.Map)
			}
			return _302_v72
		})(), 4)
		(_nw28).ArraySet1(func() _dafny.Map {
			var _coll46 = _dafny.NewMapBuilder()
			_ = _coll46
			for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(100), _dafny.IntOfInt64(8))); ; {
				_compr_46, _ok46 := _iter46()
				if !_ok46 {
					break
				}
				var _307_v74 _dafny.Int
				_307_v74 = interface{}(_compr_46).(_dafny.Int)
				if ((_dafny.IntOfInt64(100)).Cmp(_307_v74) <= 0) && ((_307_v74).Cmp(_dafny.IntOfInt64(8)) < 0) {
					_coll46.Add((_307_v74).Plus(_290_v60), (_277_v47).F27())
				}
			}
			return _coll46.ToMap()
		}(), 5)
		(_nw28).ArraySet1(_302_v72, 6)
		(_nw28).ArraySet1(_302_v72, 7)
		(_nw28).ArraySet1(_302_v72, 8)
		(_nw28).ArraySet1(_302_v72, 9)
		(_nw28).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v0, (_277_v47).F26())).Update((_305_v76).Cardinality(), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool)), 10)
		(_nw28).ArraySet1(func() _dafny.Map {
			var _coll47 = _dafny.NewMapBuilder()
			_ = _coll47
			for _iter47 := _dafny.Iterate((_279_v49).Elements()); ; {
				_compr_47, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _308_v77 _dafny.Int
				_308_v77 = interface{}(_compr_47).(_dafny.Int)
				if (_279_v49).Contains(_308_v77) {
					_coll47.Add(Companion_Default___.SafeModuloInt(_308_v77, _290_v60), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
				}
			}
			return _coll47.ToMap()
		}(), 11)
		_306_v78 = _nw28
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_306_v78), 0))
		_ = _index23
		(_306_v78).ArraySet1(_302_v72, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.ArrayLen((_306_v78), 0))
		_ = _index24
		(_306_v78).ArraySet1((_302_v72).Merge(_302_v72), (_index24).Int())
		(_277_v47).M1((_277_v47).F27(), _dafny.SeqOf(_225_v1, _225_v1), _292_v62, _240_globalState)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_229_v5), 0))
		_ = _index25
		(_229_v5).ArraySet1((_224_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rodtn")).Cardinality())), (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_229_v5), 0))
		_ = _index26
		(_229_v5).ArraySet1((_224_v0).Plus(_224_v0), (_index26).Int())
		var _309_v79 *C5
		_ = _309_v79
		var _nw29 *C5 = New_C5_()
		_ = _nw29
		_nw29.Ctor__((_277_v47).F27(), (_288_v58).Fm2(true, _dafny.IntOfInt64(144), _240_globalState))
		_309_v79 = _nw29
		var _310_v80 D25
		_ = _310_v80
		_310_v80 = Companion_D25_.Create_DC62_((_281_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_281_v51), 0))).Int()).(*C8))
		var _311_v81 _dafny.Array
		_ = _311_v81
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
		_ = _nw30
		_311_v81 = _nw30
		var _312_v82 D3
		_ = _312_v82
		_312_v82 = Companion_D3_.Create_DC10_(_311_v81)
		var _313_v83 _dafny.Map
		_ = _313_v83
		_313_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_312_v82, _224_v0)
		var _314_v84 T1
		_ = _314_v84
		var _nw31 *C4 = New_C4_()
		_ = _nw31
		_nw31.Ctor__(_313_v83, (_277_v47).F27(), (_277_v47).F27())
		_314_v84 = _nw31
		var _315_v85 D24
		_ = _315_v85
		_315_v85 = Companion_D24_.Create_DC60_(_314_v84, _dafny.IntOfUint32((_225_v1).Cardinality()), (_dafny.SetOf(_dafny.IntOfUint32((_292_v62).Cardinality()))).Cardinality(), (_314_v84).F27())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_229_v5), 0))
		_ = _index27
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_281_v51), 0))
		_ = _index28
		var _rhs21 _dafny.Int = (_dafny.IntOfInt64(-755)).Minus(_dafny.IntOfInt64(8))
		_ = _rhs21
		var _rhs22 *C5 = _309_v79
		_ = _rhs22
		var _rhs23 *C8 = (_310_v80).Dtor_cf115()
		_ = _rhs23
		var _rhs24 _dafny.Sequence = _292_v62
		_ = _rhs24
		var _rhs25 _dafny.Int = Companion_Default___.SafeDivisionInt(_224_v0, (_315_v85).Dtor_cf112())
		_ = _rhs25
		var _lhs17 _dafny.Array = _229_v5
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_229_v5), 0))
		_ = _lhs18
		var _lhs19 _dafny.Array = _281_v51
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_281_v51), 0))
		_ = _lhs20
		(_lhs17).ArraySet1(_rhs21, (_lhs18).Int())
		_309_v79 = _rhs22
		(_lhs19).ArraySet1(_rhs23, (_lhs20).Int())
		_292_v62 = _rhs24
		_224_v0 = _rhs25
	}
	(_240_globalState).F2 = (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool)
	var _316_v86 _dafny.Array
	_ = _316_v86
	var _nw32 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
	_ = _nw32
	_316_v86 = _nw32
	var _317_v87 *C5
	_ = _317_v87
	var _nw33 *C5 = New_C5_()
	_ = _nw33
	_nw33.Ctor__(_293_v63, _241_v14)
	_317_v87 = _nw33
	var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_316_v86), 0))
	_ = _index29
	(_316_v86).ArraySet1(_317_v87, (_index29).Int())
	var _318_v88 _dafny.Sequence
	_ = _318_v88
	_318_v88 = _dafny.SeqOf(_288_v58, _288_v58)
	var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_316_v86), 0))
	_ = _index30
	var _rhs26 _dafny.CodePoint = _289_v59
	_ = _rhs26
	var _rhs27 _dafny.Int = _224_v0
	_ = _rhs27
	var _rhs28 *C5 = _317_v87
	_ = _rhs28
	var _rhs29 bool = !(_241_v14)
	_ = _rhs29
	var _rhs30 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_288_v58, _288_v58, _288_v58, _288_v58, _288_v58), _dafny.Companion_Sequence_.Concatenate(_318_v88, _318_v88))
	_ = _rhs30
	var _lhs21 *GlobalState = _240_globalState
	_ = _lhs21
	var _lhs22 *GlobalState = _240_globalState
	_ = _lhs22
	var _lhs23 _dafny.Array = _316_v86
	_ = _lhs23
	var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_316_v86), 0))
	_ = _lhs24
	var _lhs25 *GlobalState = _240_globalState
	_ = _lhs25
	var _lhs26 *GlobalState = _240_globalState
	_ = _lhs26
	_lhs21.F19 = _rhs26
	_lhs22.F12 = _rhs27
	(_lhs23).ArraySet1(_rhs28, (_lhs24).Int())
	_lhs25.F15 = _rhs29
	_lhs26.F22 = _rhs30
	var _319_v89 _dafny.Sequence
	_ = _319_v89
	_319_v89 = _dafny.SeqOf(_229_v5, _229_v5, _229_v5)
	_229_v5 = (_319_v89).Select((Companion_Default___.SafeIndex((_243_v16).Cardinality(), _dafny.IntOfUint32((_319_v89).Cardinality()))).Uint32()).(_dafny.Array)
	var _source8 D7 = Companion_D7_.Create_DC18_(_277_v47)
	_ = _source8
	if _source8.Is_DC19() {
		var _320___mcc_h0 _dafny.Int = _source8.Get_().(D7_DC19).Cf36
		_ = _320___mcc_h0
		var _321___mcc_h1 _dafny.MultiSet = _source8.Get_().(D7_DC19).Cf37
		_ = _321___mcc_h1
		var _322___mcc_h2 bool = _source8.Get_().(D7_DC19).Cf38
		_ = _322___mcc_h2
		var _323___mcc_h3 _dafny.Map = _source8.Get_().(D7_DC19).Cf39
		_ = _323___mcc_h3
		var _324_cf39 _dafny.Map = _323___mcc_h3
		_ = _324_cf39
		var _325_cf38 bool = _322___mcc_h2
		_ = _325_cf38
		var _326_cf37 _dafny.MultiSet = _321___mcc_h1
		_ = _326_cf37
		var _327_cf36 _dafny.Int = _320___mcc_h0
		_ = _327_cf36
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
		_ = _index31
		(_269_v41).ArraySet1((_288_v58).Fm2((_277_v47).F26(), _327_cf36, _240_globalState), (_index31).Int())
		(_317_v87).M7((_276_v46).IsSubsetOf(_dafny.MultiSetFromSeq(_272_v42)), (_277_v47).F26(), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_327_cf36, _290_v60)), _276_v46, _240_globalState)
		_229_v5 = _229_v5
		var _328_v90 _dafny.Map
		_ = _328_v90
		_328_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v5, _229_v5)
		var _329_v91 _dafny.Map
		_ = _329_v91
		_329_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_328_v90, _290_v60)
		_329_v91 = (_329_v91).Update(_328_v90, _224_v0)
	} else {
		var _330___mcc_h4 T2 = _source8.Get_().(D7_DC18).Cf35
		_ = _330___mcc_h4
		var _331_cf35 T2 = _330___mcc_h4
		_ = _331_cf35
		var _332_v92 _dafny.Map
		_ = _332_v92
		_332_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v1, _dafny.IntOfUint32((_225_v1).Cardinality()))
		var _333_v93 _dafny.Map
		_ = _333_v93
		_333_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_282_v52).Cardinality(), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
		(_240_globalState).F16 = (func() _dafny.Int {
			if (_332_v92).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-115))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}((func(_334_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_335_i3 _dafny.Int) _dafny.CodePoint {
					return _334_v8
				}
			})(_234_v8)))) {
				return (_332_v92).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-115))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_336_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_337_i3 _dafny.Int) _dafny.CodePoint {
						return _336_v8
					}
				})(_234_v8)))).(_dafny.Int)
			}
			return (_333_v93).Cardinality()
		})()
		if !(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_225_v1, _225_v1))).Equals(_dafny.MultiSetFromSeq(_dafny.SeqOf(_291_v61))) {
			var _338_v94 D0
			_ = _338_v94
			_338_v94 = Companion_D0_.Create_DC0_((_331_cf35).F26())
			(_240_globalState).F16 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(753), (_277_v47).Fm4(_338_v94, _224_v0, _224_v0, _240_globalState)), _dafny.IntOfInt64(123)))
			(_240_globalState).F24 = (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
			_ = _index32
			(_269_v41).ArraySet1((func() bool {
				if (_331_cf35).F26() {
					return (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool)
				}
				return !_dafny.Companion_Sequence_.Contains(_225_v1, _dafny.CodePoint('s'))
			})(), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
			_ = _index33
			(_269_v41).ArraySet1((_331_cf35).F27(), (_index33).Int())
			var _pat_let_tv0 = _277_v47
			_ = _pat_let_tv0
			var _339_v95 bool
			_ = _339_v95
			var _out18 bool
			_ = _out18
			_out18 = (_317_v87).M8(func(_pat_let0_0 D0) D0 {
				return func(_340_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let1_0 bool) D0 {
						return func(_341_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_341_dt__update_hcf0_h0)
						}(_pat_let1_0)
					}((_pat_let_tv0).F26())
				}(_pat_let0_0)
			}(_338_v94), (_331_cf35).F27(), _295_v65, (_224_v0).Times(_290_v60), _240_globalState)
			_339_v95 = _out18
		} else {
			var _342_v96 _dafny.Map
			_ = _342_v96
			_342_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_331_cf35).F27(), _275_v45)
			_342_v96 = (_342_v96).Update((_277_v47).F27(), _275_v45)
			var _343_v97 _dafny.Sequence
			_ = _343_v97
			_343_v97 = _dafny.SeqOf(_225_v1, _225_v1, _225_v1, _225_v1)
			var _344_v98 *C9
			_ = _344_v98
			var _nw34 *C9 = New_C9_()
			_ = _nw34
			_nw34.Ctor__(false, _224_v0, (_277_v47).F27(), !_dafny.Companion_Sequence_.Equal((_343_v97).Select((Companion_Default___.SafeIndex(_224_v0, _dafny.IntOfUint32((_343_v97).Cardinality()))).Uint32()).(_dafny.Sequence), _225_v1), _279_v49)
			_344_v98 = _nw34
			var _345_v99 T1
			_ = _345_v99
			var _nw35 *C5 = New_C5_()
			_ = _nw35
			_nw35.Ctor__(!((_331_cf35).F26()), (_277_v47).F27())
			_345_v99 = _nw35
			_345_v99 = _345_v99
			(_240_globalState).F22 = (_288_v58).Fm2(!(_241_v14) || (!((_331_cf35).F26())), _290_v60, _240_globalState)
			var _346_v100 D0
			_ = _346_v100
			_346_v100 = Companion_D0_.Create_DC0_(false)
			(_240_globalState).F16 = (_331_cf35).Fm4(_346_v100, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_275_v45).Cardinality()), (_272_v42).Select((Companion_Default___.SafeIndex((_344_v98).F31(), _dafny.IntOfUint32((_272_v42).Cardinality()))).Uint32()).(_dafny.Int)), (func() _dafny.Int {
				if (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool) {
					return _224_v0
				}
				return _224_v0
			})(), _240_globalState)
		}
		var _347_v101 *C6
		_ = _347_v101
		var _nw36 *C6 = New_C6_()
		_ = _nw36
		_nw36.Ctor__((_290_v60).Times(_290_v60))
		_347_v101 = _nw36
		var _348_v102 D0
		_ = _348_v102
		_348_v102 = Companion_D0_.Create_DC0_(false)
		var _349_v103 D5
		_ = _349_v103
		_349_v103 = Companion_D5_.Create_DC14_(_224_v0, _347_v101.F40, (_348_v102).Dtor_cf0())
		var _350_v104 D5
		_ = _350_v104
		_350_v104 = Companion_D5_.Create_DC15_(_349_v103)
		var _pat_let_tv1 = _349_v103
		_ = _pat_let_tv1
		_350_v104 = (func() D5 {
			if (func() bool {
				if (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool) {
					return (_331_cf35).F27()
				}
				return true
			})() {
				return (func() D5 {
					if _293_v63 {
						return func(_pat_let2_0 D5) D5 {
							return func(_351_dt__update__tmp_h1 D5) D5 {
								return func(_pat_let3_0 D5) D5 {
									return func(_352_dt__update_hcf28_h0 D5) D5 {
										return Companion_D5_.Create_DC15_(_352_dt__update_hcf28_h0)
									}(_pat_let3_0)
								}(_pat_let_tv1)
							}(_pat_let2_0)
						}(_350_v104)
					}
					return _350_v104
				})()
			}
			return _350_v104
		})()
	}
	var _353_i4 _dafny.Int
	_ = _353_i4
	_353_i4 = _dafny.Zero
	{
		for (_295_v65).IsSubsetOf((_295_v65).Difference(_295_v65)) {
			{
				if (_353_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_353_i4 = (_353_i4).Plus(_dafny.One)
				var _354_v105 _dafny.Map
				_ = _354_v105
				_354_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(563), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
				var _355_v106 D16
				_ = _355_v106
				_355_v106 = Companion_D16_.Create_DC37_(_354_v105)
				var _356_v107 _dafny.MultiSet
				_ = _356_v107
				_356_v107 = _dafny.MultiSetOf((_355_v106).Dtor_cf68(), _354_v105)
				var _357_v108 _dafny.MultiSet
				_ = _357_v108
				_357_v108 = _356_v107
				var _source9 _dafny.MultiSet = _357_v108
				_ = _source9
				var _358___mcc_h5 _dafny.MultiSet = _source9
				_ = _358___mcc_h5
				var _359_cf82 _dafny.MultiSet = _358___mcc_h5
				_ = _359_cf82
				(_240_globalState).F2 = (_277_v47).F27()
				_243_v16 = _243_v16
				_290_v60 = Companion_Default___.SafeDivisionInt(_224_v0, (_288_v58).Fm1(_240_globalState))
				var _360_v109 _dafny.Map
				_ = _360_v109
				_360_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v5, (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
				var _rhs31 bool = (_277_v47).F27()
				_ = _rhs31
				var _rhs32 _dafny.Int = (((_360_v109).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v5, (_277_v47).F27()))).Merge(_360_v109)).Cardinality()
				_ = _rhs32
				var _lhs27 *GlobalState = _240_globalState
				_ = _lhs27
				_293_v63 = _rhs31
				_lhs27.F12 = _rhs32
				if (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool) {
					(_240_globalState).F22 = true
					var _361_v110 _dafny.Array
					_ = _361_v110
					var _len0_7 _dafny.Int = _dafny.One
					_ = _len0_7
					var _nw37 _dafny.Array
					_ = _nw37
					if _len0_7.Cmp(_dafny.Zero) == 0 {
						_nw37 = _dafny.NewArray(_len0_7)
					} else {
						var _init7 func(_dafny.Int) _dafny.Sequence = (func(_362_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_363_i5 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("glxoh"), _362_v1)
							}
						})(_225_v1)
						_ = _init7
						var _element0_7 = _init7(_dafny.Zero)
						_ = _element0_7
						_nw37 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
						(_nw37).ArraySet1(_element0_7, 0)
						var _nativeLen0_7 = (_len0_7).Int()
						_ = _nativeLen0_7
						for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
							(_nw37).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
						}
					}
					_361_v110 = _nw37
					var _364_v111 _dafny.Map
					_ = _364_v111
					_364_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v60, _225_v1)
					var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_361_v110), 0))
					_ = _index34
					(_361_v110).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-70))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg38 _dafny.Int) interface{} {
							return coer38(arg38)
						}
					}(func(_365_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('b')
					})), (func() _dafny.Sequence {
						if (_364_v111).Contains(_224_v0) {
							return (_364_v111).Get(_224_v0).(_dafny.Sequence)
						}
						return _225_v1
					})()), (_index34).Int())
					var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_361_v110), 0))
					_ = _index35
					(_361_v110).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_225_v1, _225_v1), (_index35).Int())
					var _366_v112 D3
					_ = _366_v112
					_366_v112 = Companion_D3_.Create_DC10_(_229_v5)
					var _367_v113 _dafny.Map
					_ = _367_v113
					_367_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_366_v112, _290_v60)
					var _368_v114 *C4
					_ = _368_v114
					var _nw38 *C4 = New_C4_()
					_ = _nw38
					_nw38.Ctor__(_367_v113, (_277_v47).F27(), (_277_v47).F26())
					_368_v114 = _nw38
					var _369_v115 _dafny.Map
					_ = _369_v115
					_369_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_v114, _290_v60)
					_369_v115 = _369_v115
					(_240_globalState).F6 = _225_v1
					(_240_globalState).F6 = (_361_v110).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_361_v110), 0))).Int()).(_dafny.Sequence)
				} else {
					(_240_globalState).F24 = (_277_v47).F27()
					var _370_v116 *C1
					_ = _370_v116
					var _nw39 *C1 = New_C1_()
					_ = _nw39
					_nw39.Ctor__(_dafny.IntOfUint32((_272_v42).Cardinality()), _290_v60, (_226_v2).Contains(_224_v0), (_277_v47).F26())
					_370_v116 = _nw39
					var _371_v117 _dafny.Array
					_ = _371_v117
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(8)
					_ = _len0_8
					var _nw40 _dafny.Array
					_ = _nw40
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw40 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.CodePoint = (func(_372_v61 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_373_i7 _dafny.Int) _dafny.CodePoint {
								return _372_v61
							}
						})(_291_v61)
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw40 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw40).ArraySet1CodePoint(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw40).ArraySet1CodePoint(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_371_v117 = _nw40
					var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_371_v117), 0))
					_ = _index36
					(_371_v117).ArraySet1CodePoint(_289_v59, (_index36).Int())
					var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_371_v117), 0))
					_ = _index37
					(_371_v117).ArraySet1CodePoint(_291_v61, (_index37).Int())
					var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
					_ = _index38
					(_269_v41).ArraySet1(_241_v14, (_index38).Int())
					var _374_v118 _dafny.Map
					_ = _374_v118
					_374_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_272_v42).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_370_v116.F37), _dafny.IntOfUint32((_272_v42).Cardinality()))).Uint32()).(_dafny.Int), _234_v8)
					_374_v118 = (_374_v118).Update(_290_v60, (_371_v117).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(346), _dafny.ArrayLen((_371_v117), 0))).Int()))
				}
				var _375_v119 _dafny.Map
				_ = _375_v119
				_375_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_272_v42, _224_v0)
				var _376_v120 _dafny.Set
				_ = _376_v120
				_376_v120 = _dafny.SetOf((_316_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_316_v86), 0))).Int()).(*C5))
				var _377_v121 _dafny.Sequence
				_ = _377_v121
				_377_v121 = _dafny.SeqOf(_376_v120, _376_v120, _376_v120, _376_v120)
				_375_v119 = (_375_v119).Update(_dafny.Companion_Sequence_.Concatenate(_272_v42, _272_v42), ((_377_v121).Select((Companion_Default___.SafeIndex(_224_v0, _dafny.IntOfUint32((_377_v121).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
				var _378_v122 _dafny.Array
				_ = _378_v122
				var _nw41 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
				_ = _nw41
				_378_v122 = _nw41
				_378_v122 = _378_v122
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _hi1 _dafny.Int = _290_v60
	_ = _hi1
	for _379_i8 := (_dafny.SetOf(false, (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool), (_277_v47).F26(), _241_v14)).Cardinality(); _379_i8.Cmp(_hi1) < 0; _379_i8 = _379_i8.Plus(_dafny.One) {
		_269_v41 = _269_v41
		var _380_v123 _dafny.Set
		_ = _380_v123
		_380_v123 = _dafny.SetOf((_277_v47).F26(), (_277_v47).F26(), (_269_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))).Int()).(bool))
		var _381_v124 D1
		_ = _381_v124
		_381_v124 = Companion_D1_.Create_DC5_((_dafny.Zero).Minus((_380_v123).Cardinality()))
		var _rhs33 _dafny.Int = (_379_i8).Minus((_381_v124).Dtor_cf12())
		_ = _rhs33
		var _rhs34 _dafny.Int = _224_v0
		_ = _rhs34
		var _rhs35 _dafny.Int = _dafny.IntOfInt64(855)
		_ = _rhs35
		var _rhs36 bool = (_226_v2).IsDisjointFrom(_dafny.SetOf(_379_i8))
		_ = _rhs36
		var _lhs28 *GlobalState = _240_globalState
		_ = _lhs28
		var _lhs29 *GlobalState = _240_globalState
		_ = _lhs29
		var _lhs30 *GlobalState = _240_globalState
		_ = _lhs30
		var _lhs31 *GlobalState = _240_globalState
		_ = _lhs31
		_lhs28.F16 = _rhs33
		_lhs29.F12 = _rhs34
		_lhs30.F12 = _rhs35
		_lhs31.F24 = _rhs36
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
		_ = _index39
		var _rhs37 _dafny.Int = _dafny.IntOfInt64(-455)
		_ = _rhs37
		var _rhs38 bool = (_277_v47).Fm2(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v60, (_dafny.Zero).Minus(_dafny.IntOfInt64(-932)))).Cardinality(), _240_globalState)
		_ = _rhs38
		var _lhs32 *GlobalState = _240_globalState
		_ = _lhs32
		var _lhs33 _dafny.Array = _269_v41
		_ = _lhs33
		var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
		_ = _lhs34
		_lhs32.F12 = _rhs37
		(_lhs33).ArraySet1(_rhs38, (_lhs34).Int())
		var _382_v125 _dafny.MultiSet
		_ = _382_v125
		_382_v125 = _dafny.MultiSetOf(_dafny.SeqOf(_224_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yrpw")).Cardinality()), _379_i8, (_dafny.Zero).Minus(_379_i8), _dafny.IntOfInt64(123)))
		var _383_v126 *C3
		_ = _383_v126
		var _nw42 *C3 = New_C3_()
		_ = _nw42
		_nw42.Ctor__(_382_v125, false, !((_277_v47).F27()))
		_383_v126 = _nw42
		var _384_v127 D23
		_ = _384_v127
		_384_v127 = Companion_D23_.Create_DC55_(_383_v126)
		_383_v126 = (_384_v127).Dtor_cf103()
	}
	var _385_v128 bool
	_ = _385_v128
	var _386_v129 bool
	_ = _386_v129
	var _out19 bool
	_ = _out19
	var _out20 bool
	_ = _out20
	_out19, _out20 = (_280_v50).M4(_240_globalState)
	_385_v128 = _out19
	_386_v129 = _out20
	var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_269_v41), 0))
	_ = _index40
	(_269_v41).ArraySet1(!(!(_241_v14)), (_index40).Int())
	_dafny.Print(_224_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_225_v1.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_226_v2).Equals(_dafny.SetOf(_dafny.IntOfInt64(979), _dafny.IntOfInt64(-6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_227_v3).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ovdhmu"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_232_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_234_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v9).Equals(_dafny.SetOf(_dafny.CodePoint('v'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(-441)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_239_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(-441)), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F0()).Equals(_dafny.SetOf(_dafny.IntOfInt64(979), _dafny.IntOfInt64(-6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F1()).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ovdhmu"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F6.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F7()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F14()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F15)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F16)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F20().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_240_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_240_globalState.F24)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_240_globalState).F25()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(-441)), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_241_v14)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v15).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(979)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_244_v17).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(979)), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v41).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v41).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v41).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_269_v41).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_272_v42, _dafny.SeqOf(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979), _dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_273_v43).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_v44).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_275_v45, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_276_v46).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v47).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v47).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_278_v48).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v49).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.One, _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_281_v51).ArrayGet1((_dafny.Zero).Int()).(*C8)).F29()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_281_v51).ArrayGet1((_dafny.One).Int()).(*C8)).F29()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(4))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v52).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(1958), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979)).UpdateUnsafe(_dafny.IntOfInt64(6), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(6))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(979)).UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_283_v55).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(979), _dafny.IntOfInt64(979))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_286_v56).Dtor_cf29().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_287_v57).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_289_v59)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_290_v60)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_291_v61)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_292_v62, _dafny.SeqOf(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_293_v63)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_294_v64).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, true, true, true), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v65).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_316_v86).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C5)).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_316_v86).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(*C5)).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_318_v88).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_319_v89).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_353_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_385_v128)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_386_v129)
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
	Cf1 bool
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool, Cf3 _dafny.Int, Cf4 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 bool
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 bool) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC2 struct {
	Cf5 D0
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 D0) D0 {
	return D0{D0_DC2{Cf5}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, false, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf5() D0 {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Equals(data2.Cf5)
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
	Cf7  _dafny.Array
	Cf8  _dafny.Int
	Cf9  _dafny.Int
	Cf10 _dafny.Sequence
	Cf11 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf7 _dafny.Array, Cf8 _dafny.Int, Cf9 _dafny.Int, Cf10 _dafny.Sequence, Cf11 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf7, Cf8, Cf9, Cf10, Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf12 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf12 _dafny.Int) D1 {
	return D1{D1_DC5{Cf12}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf6 _dafny.Sequence
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf6 _dafny.Sequence) D1 {
	return D1{D1_DC3{Cf6}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero, _dafny.Zero, _dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D1) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D1_DC4).Cf7
}

func (_this D1) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + data.Cf10.VerbatimString(true) + ", " + data.Cf11.VerbatimString(true) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Equals(data2.Cf11)
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0
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

type D2_DC7 struct {
	Cf14 bool
	Cf15 bool
	Cf16 _dafny.Int
	Cf17 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf14 bool, Cf15 bool, Cf16 _dafny.Int, Cf17 _dafny.Int) D2 {
	return D2{D2_DC7{Cf14, Cf15, Cf16, Cf17}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_() D2 {
	return D2{D2_DC8{}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf18 _dafny.Int
	Cf19 _dafny.Int
	Cf20 bool
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf18 _dafny.Int, Cf19 _dafny.Int, Cf20 bool) D2 {
	return D2{D2_DC9{Cf18, Cf19, Cf20}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC6 struct {
	Cf13 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf13 _dafny.Map) D2 {
	return D2{D2_DC6{Cf13}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, false, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf14() bool {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf17
}

func (_this D2) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf18
}

func (_this D2) Dtor_cf19() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf19
}

func (_this D2) Dtor_cf20() bool {
	return _this.Get_().(D2_DC9).Cf20
}

func (_this D2) Dtor_cf13() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf13
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC8:
		{
			_, ok := other.Get_().(D2_DC8)
			return ok
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Cmp(data2.Cf19) == 0 && data1.Cf20 == data2.Cf20
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf13.Equals(data2.Cf13)
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

type D3_DC11 struct {
	Cf22 _dafny.CodePoint
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf22 _dafny.CodePoint) D3 {
	return D3{D3_DC11{Cf22}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf21 _dafny.Array
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf21 _dafny.Array) D3 {
	return D3{D3_DC10{Cf21}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(_dafny.CodePoint('D'))
}

func (_this D3) Dtor_cf22() _dafny.CodePoint {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D3_DC10).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf22 == data2.Cf22
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf21 == data2.Cf21
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
	Cf23 _dafny.Set
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf23 _dafny.Set) D4 {
	return D4{D4_DC12{Cf23}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D4) Dtor_cf23() _dafny.Set {
	return _this.Get_().(D4_DC12).Cf23
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf23) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23)
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
	Cf25 _dafny.Int
	Cf26 _dafny.Int
	Cf27 bool
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 _dafny.Int, Cf26 _dafny.Int, Cf27 bool) D5 {
	return D5{D5_DC14{Cf25, Cf26, Cf27}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

type D5_DC13 struct {
	Cf24 _dafny.Map
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf24 _dafny.Map) D5 {
	return D5{D5_DC13{Cf24}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
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
	return Companion_D5_.Create_DC14_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D5) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D5_DC14).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC14).Cf27
}

func (_this D5) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf28() D5 {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf24) + ")"
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
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27 == data2.Cf27
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf24.Equals(data2.Cf24)
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
	Cf30 _dafny.Int
	Cf31 _dafny.Int
	Cf32 _dafny.Int
	Cf33 _dafny.Int
	Cf34 _dafny.Int
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf30 _dafny.Int, Cf31 _dafny.Int, Cf32 _dafny.Int, Cf33 _dafny.Int, Cf34 _dafny.Int) D6 {
	return D6{D6_DC17{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf29 _dafny.Sequence
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf29 _dafny.Sequence) D6 {
	return D6{D6_DC16{Cf29}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D6) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf30
}

func (_this D6) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf32
}

func (_this D6) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf33
}

func (_this D6) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D6_DC17).Cf34
}

func (_this D6) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D6_DC16).Cf29
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + data.Cf29.VerbatimString(true) + ")"
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
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32.Cmp(data2.Cf32) == 0 && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34.Cmp(data2.Cf34) == 0
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
	Cf36 _dafny.Int
	Cf37 _dafny.MultiSet
	Cf38 bool
	Cf39 _dafny.Map
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf36 _dafny.Int, Cf37 _dafny.MultiSet, Cf38 bool, Cf39 _dafny.Map) D7 {
	return D7{D7_DC19{Cf36, Cf37, Cf38, Cf39}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC18 struct {
	Cf35 T2
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf35 T2) D7 {
	return D7{D7_DC18{Cf35}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(_dafny.Zero, _dafny.EmptyMultiSet, false, _dafny.EmptyMap)
}

func (_this D7) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf36
}

func (_this D7) Dtor_cf37() _dafny.MultiSet {
	return _this.Get_().(D7_DC19).Cf37
}

func (_this D7) Dtor_cf38() bool {
	return _this.Get_().(D7_DC19).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D7_DC19).Cf39
}

func (_this D7) Dtor_cf35() T2 {
	return _this.Get_().(D7_DC18).Cf35
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37.Equals(data2.Cf37) && data1.Cf38 == data2.Cf38 && data1.Cf39.Equals(data2.Cf39)
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && _dafny.AreEqual(data1.Cf35, data2.Cf35)
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

type D8_DC20 struct {
	Cf40 _dafny.Sequence
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf40 _dafny.Sequence) D8 {
	return D8{D8_DC20{Cf40}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D8) Dtor_cf40() _dafny.Sequence {
	return _this.Get_().(D8_DC20).Cf40
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf40.Equals(data2.Cf40)
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

type D9_DC22 struct {
	Cf42 bool
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf42 bool) D9 {
	return D9{D9_DC22{Cf42}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC21 struct {
	Cf41 _dafny.Map
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf41 _dafny.Map) D9 {
	return D9{D9_DC21{Cf41}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC22_(false)
}

func (_this D9) Dtor_cf42() bool {
	return _this.Get_().(D9_DC22).Cf42
}

func (_this D9) Dtor_cf41() _dafny.Map {
	return _this.Get_().(D9_DC21).Cf41
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf42 == data2.Cf42
		}
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D10_DC24 struct {
	Cf44 _dafny.Sequence
	Cf45 _dafny.Sequence
	Cf46 bool
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf44 _dafny.Sequence, Cf45 _dafny.Sequence, Cf46 bool) D10 {
	return D10{D10_DC24{Cf44, Cf45, Cf46}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

type D10_DC25 struct {
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_() D10 {
	return D10{D10_DC25{}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

type D10_DC23 struct {
	Cf43 _dafny.Map
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf43 _dafny.Map) D10 {
	return D10{D10_DC23{Cf43}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC24_(_dafny.EmptySeq, _dafny.EmptySeq, false)
}

func (_this D10) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D10_DC24).Cf44
}

func (_this D10) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D10_DC24).Cf45
}

func (_this D10) Dtor_cf46() bool {
	return _this.Get_().(D10_DC24).Cf46
}

func (_this D10) Dtor_cf43() _dafny.Map {
	return _this.Get_().(D10_DC23).Cf43
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC24:
		{
			return "D10.DC24" + "(" + data.Cf44.VerbatimString(true) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D10_DC25:
		{
			return "D10.DC25"
		}
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf44.Equals(data2.Cf44) && data1.Cf45.Equals(data2.Cf45) && data1.Cf46 == data2.Cf46
		}
	case D10_DC25:
		{
			_, ok := other.Get_().(D10_DC25)
			return ok
		}
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf43.Equals(data2.Cf43)
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

type D11_DC26 struct {
	Cf47 _dafny.Array
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf47 _dafny.Array) D11 {
	return D11{D11_DC26{Cf47}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D11) Dtor_cf47() _dafny.Array {
	return _this.Get_().(D11_DC26).Cf47
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf47 == data2.Cf47
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

type D12_DC28 struct {
	Cf49 _dafny.Int
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf49 _dafny.Int) D12 {
	return D12{D12_DC28{Cf49}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC27 struct {
	Cf48 T0
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf48 T0) D12 {
	return D12{D12_DC27{Cf48}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC28_(_dafny.Zero)
}

func (_this D12) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D12_DC28).Cf49
}

func (_this D12) Dtor_cf48() T0 {
	return _this.Get_().(D12_DC27).Cf48
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf49.Cmp(data2.Cf49) == 0
		}
	case D12_DC27:
		{
			data2, ok := other.Get_().(D12_DC27)
			return ok && _dafny.AreEqual(data1.Cf48, data2.Cf48)
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

type D13_DC30 struct {
	Cf51 bool
	Cf52 _dafny.Int
	Cf53 bool
	Cf54 _dafny.Array
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf51 bool, Cf52 _dafny.Int, Cf53 bool, Cf54 _dafny.Array) D13 {
	return D13{D13_DC30{Cf51, Cf52, Cf53, Cf54}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

type D13_DC29 struct {
	Cf50 _dafny.Map
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf50 _dafny.Map) D13 {
	return D13{D13_DC29{Cf50}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC30_(false, _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D13) Dtor_cf51() bool {
	return _this.Get_().(D13_DC30).Cf51
}

func (_this D13) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D13_DC30).Cf52
}

func (_this D13) Dtor_cf53() bool {
	return _this.Get_().(D13_DC30).Cf53
}

func (_this D13) Dtor_cf54() _dafny.Array {
	return _this.Get_().(D13_DC30).Cf54
}

func (_this D13) Dtor_cf50() _dafny.Map {
	return _this.Get_().(D13_DC29).Cf50
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54
		}
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf50.Equals(data2.Cf50)
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

type D14_DC32 struct {
	Cf56 _dafny.MultiSet
	Cf57 bool
	Cf58 bool
	Cf59 bool
	Cf60 _dafny.Int
}

func (D14_DC32) isD14() {}

func (CompanionStruct_D14_) Create_DC32_(Cf56 _dafny.MultiSet, Cf57 bool, Cf58 bool, Cf59 bool, Cf60 _dafny.Int) D14 {
	return D14{D14_DC32{Cf56, Cf57, Cf58, Cf59, Cf60}}
}

func (_this D14) Is_DC32() bool {
	_, ok := _this.Get_().(D14_DC32)
	return ok
}

type D14_DC31 struct {
	Cf55 _dafny.Sequence
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf55 _dafny.Sequence) D14 {
	return D14{D14_DC31{Cf55}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC32_(_dafny.EmptyMultiSet, false, false, false, _dafny.Zero)
}

func (_this D14) Dtor_cf56() _dafny.MultiSet {
	return _this.Get_().(D14_DC32).Cf56
}

func (_this D14) Dtor_cf57() bool {
	return _this.Get_().(D14_DC32).Cf57
}

func (_this D14) Dtor_cf58() bool {
	return _this.Get_().(D14_DC32).Cf58
}

func (_this D14) Dtor_cf59() bool {
	return _this.Get_().(D14_DC32).Cf59
}

func (_this D14) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D14_DC32).Cf60
}

func (_this D14) Dtor_cf55() _dafny.Sequence {
	return _this.Get_().(D14_DC31).Cf55
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC32:
		{
			return "D14.DC32" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC32:
		{
			data2, ok := other.Get_().(D14_DC32)
			return ok && data1.Cf56.Equals(data2.Cf56) && data1.Cf57 == data2.Cf57 && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59 && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D15_DC34 struct {
	Cf62 _dafny.Int
	Cf63 bool
	Cf64 _dafny.Int
}

func (D15_DC34) isD15() {}

func (CompanionStruct_D15_) Create_DC34_(Cf62 _dafny.Int, Cf63 bool, Cf64 _dafny.Int) D15 {
	return D15{D15_DC34{Cf62, Cf63, Cf64}}
}

func (_this D15) Is_DC34() bool {
	_, ok := _this.Get_().(D15_DC34)
	return ok
}

type D15_DC35 struct {
	Cf65 bool
}

func (D15_DC35) isD15() {}

func (CompanionStruct_D15_) Create_DC35_(Cf65 bool) D15 {
	return D15{D15_DC35{Cf65}}
}

func (_this D15) Is_DC35() bool {
	_, ok := _this.Get_().(D15_DC35)
	return ok
}

type D15_DC36 struct {
	Cf66 bool
	Cf67 _dafny.Int
}

func (D15_DC36) isD15() {}

func (CompanionStruct_D15_) Create_DC36_(Cf66 bool, Cf67 _dafny.Int) D15 {
	return D15{D15_DC36{Cf66, Cf67}}
}

func (_this D15) Is_DC36() bool {
	_, ok := _this.Get_().(D15_DC36)
	return ok
}

type D15_DC33 struct {
	Cf61 _dafny.Map
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf61 _dafny.Map) D15 {
	return D15{D15_DC33{Cf61}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC34_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D15) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D15_DC34).Cf62
}

func (_this D15) Dtor_cf63() bool {
	return _this.Get_().(D15_DC34).Cf63
}

func (_this D15) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D15_DC34).Cf64
}

func (_this D15) Dtor_cf65() bool {
	return _this.Get_().(D15_DC35).Cf65
}

func (_this D15) Dtor_cf66() bool {
	return _this.Get_().(D15_DC36).Cf66
}

func (_this D15) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D15_DC36).Cf67
}

func (_this D15) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D15_DC33).Cf61
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC34:
		{
			return "D15.DC34" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D15_DC35:
		{
			return "D15.DC35" + "(" + _dafny.String(data.Cf65) + ")"
		}
	case D15_DC36:
		{
			return "D15.DC36" + "(" + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ")"
		}
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC34:
		{
			data2, ok := other.Get_().(D15_DC34)
			return ok && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63 == data2.Cf63 && data1.Cf64.Cmp(data2.Cf64) == 0
		}
	case D15_DC35:
		{
			data2, ok := other.Get_().(D15_DC35)
			return ok && data1.Cf65 == data2.Cf65
		}
	case D15_DC36:
		{
			data2, ok := other.Get_().(D15_DC36)
			return ok && data1.Cf66 == data2.Cf66 && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

type D16_DC38 struct {
	Cf69 _dafny.Int
	Cf70 _dafny.Int
	Cf71 bool
}

func (D16_DC38) isD16() {}

func (CompanionStruct_D16_) Create_DC38_(Cf69 _dafny.Int, Cf70 _dafny.Int, Cf71 bool) D16 {
	return D16{D16_DC38{Cf69, Cf70, Cf71}}
}

func (_this D16) Is_DC38() bool {
	_, ok := _this.Get_().(D16_DC38)
	return ok
}

type D16_DC37 struct {
	Cf68 _dafny.Map
}

func (D16_DC37) isD16() {}

func (CompanionStruct_D16_) Create_DC37_(Cf68 _dafny.Map) D16 {
	return D16{D16_DC37{Cf68}}
}

func (_this D16) Is_DC37() bool {
	_, ok := _this.Get_().(D16_DC37)
	return ok
}

type D16_DC39 struct {
	Cf72 D16
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf72 D16) D16 {
	return D16{D16_DC39{Cf72}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC38_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D16) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D16_DC38).Cf69
}

func (_this D16) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D16_DC38).Cf70
}

func (_this D16) Dtor_cf71() bool {
	return _this.Get_().(D16_DC38).Cf71
}

func (_this D16) Dtor_cf68() _dafny.Map {
	return _this.Get_().(D16_DC37).Cf68
}

func (_this D16) Dtor_cf72() D16 {
	return _this.Get_().(D16_DC39).Cf72
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC38:
		{
			return "D16.DC38" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D16_DC37:
		{
			return "D16.DC37" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf72) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC38:
		{
			data2, ok := other.Get_().(D16_DC38)
			return ok && data1.Cf69.Cmp(data2.Cf69) == 0 && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71 == data2.Cf71
		}
	case D16_DC37:
		{
			data2, ok := other.Get_().(D16_DC37)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf72.Equals(data2.Cf72)
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

// Definition of datatype D17
type D17 struct {
	Data_D17_
}

func (_this D17) Get_() Data_D17_ {
	return _this.Data_D17_
}

type Data_D17_ interface {
	isD17()
}

type CompanionStruct_D17_ struct {
}

var Companion_D17_ = CompanionStruct_D17_{}

type D17_DC41 struct {
	Cf74 bool
	Cf75 bool
	Cf76 _dafny.Int
}

func (D17_DC41) isD17() {}

func (CompanionStruct_D17_) Create_DC41_(Cf74 bool, Cf75 bool, Cf76 _dafny.Int) D17 {
	return D17{D17_DC41{Cf74, Cf75, Cf76}}
}

func (_this D17) Is_DC41() bool {
	_, ok := _this.Get_().(D17_DC41)
	return ok
}

type D17_DC42 struct {
	Cf77 _dafny.Int
	Cf78 _dafny.Int
	Cf79 _dafny.Int
}

func (D17_DC42) isD17() {}

func (CompanionStruct_D17_) Create_DC42_(Cf77 _dafny.Int, Cf78 _dafny.Int, Cf79 _dafny.Int) D17 {
	return D17{D17_DC42{Cf77, Cf78, Cf79}}
}

func (_this D17) Is_DC42() bool {
	_, ok := _this.Get_().(D17_DC42)
	return ok
}

type D17_DC43 struct {
	Cf80 _dafny.Int
	Cf81 bool
}

func (D17_DC43) isD17() {}

func (CompanionStruct_D17_) Create_DC43_(Cf80 _dafny.Int, Cf81 bool) D17 {
	return D17{D17_DC43{Cf80, Cf81}}
}

func (_this D17) Is_DC43() bool {
	_, ok := _this.Get_().(D17_DC43)
	return ok
}

type D17_DC40 struct {
	Cf73 _dafny.Sequence
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_(Cf73 _dafny.Sequence) D17 {
	return D17{D17_DC40{Cf73}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC41_(false, false, _dafny.Zero)
}

func (_this D17) Dtor_cf74() bool {
	return _this.Get_().(D17_DC41).Cf74
}

func (_this D17) Dtor_cf75() bool {
	return _this.Get_().(D17_DC41).Cf75
}

func (_this D17) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D17_DC41).Cf76
}

func (_this D17) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D17_DC42).Cf77
}

func (_this D17) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D17_DC42).Cf78
}

func (_this D17) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D17_DC42).Cf79
}

func (_this D17) Dtor_cf80() _dafny.Int {
	return _this.Get_().(D17_DC43).Cf80
}

func (_this D17) Dtor_cf81() bool {
	return _this.Get_().(D17_DC43).Cf81
}

func (_this D17) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D17_DC40).Cf73
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC41:
		{
			return "D17.DC41" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D17_DC42:
		{
			return "D17.DC42" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D17_DC43:
		{
			return "D17.DC43" + "(" + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D17_DC40:
		{
			return "D17.DC40" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC41:
		{
			data2, ok := other.Get_().(D17_DC41)
			return ok && data1.Cf74 == data2.Cf74 && data1.Cf75 == data2.Cf75 && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D17_DC42:
		{
			data2, ok := other.Get_().(D17_DC42)
			return ok && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D17_DC43:
		{
			data2, ok := other.Get_().(D17_DC43)
			return ok && data1.Cf80.Cmp(data2.Cf80) == 0 && data1.Cf81 == data2.Cf81
		}
	case D17_DC40:
		{
			data2, ok := other.Get_().(D17_DC40)
			return ok && data1.Cf73.Equals(data2.Cf73)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D17) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D17)
	return ok && _this.Equals(typed)
}

func Type_D17_() _dafny.TypeDescriptor {
	return type_D17_{}
}

type type_D17_ struct {
}

func (_this type_D17_) Default() interface{} {
	return Companion_D17_.Default()
}

func (_this type_D17_) String() string {
	return "main.D17"
}
func (_this D17) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D17{}

// End of datatype D17

// Definition of datatype D18
type D18 struct {
	Data_D18_
}

func (_this D18) Get_() Data_D18_ {
	return _this.Data_D18_
}

type Data_D18_ interface {
	isD18()
}

type CompanionStruct_D18_ struct {
}

var Companion_D18_ = CompanionStruct_D18_{}

type D18_DC44 struct {
	Cf82 _dafny.MultiSet
}

func (D18_DC44) isD18() {}

func (CompanionStruct_D18_) Create_DC44_(Cf82 _dafny.MultiSet) D18 {
	return D18{D18_DC44{Cf82}}
}

func (_this D18) Is_DC44() bool {
	_, ok := _this.Get_().(D18_DC44)
	return ok
}

func (CompanionStruct_D18_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D18) Dtor_cf82() _dafny.MultiSet {
	return _this.Get_().(D18_DC44).Cf82
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC44:
		{
			return "D18.DC44" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC44:
		{
			data2, ok := other.Get_().(D18_DC44)
			return ok && data1.Cf82.Equals(data2.Cf82)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D18) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D18)
	return ok && _this.Equals(typed)
}

func Type_D18_() _dafny.TypeDescriptor {
	return type_D18_{}
}

type type_D18_ struct {
}

func (_this type_D18_) Default() interface{} {
	return Companion_D18_.Default()
}

func (_this type_D18_) String() string {
	return "main.D18"
}
func (_this D18) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D18{}

// End of datatype D18

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC45 struct {
	Cf83 _dafny.Map
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf83 _dafny.Map) D19 {
	return D19{D19_DC45{Cf83}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

func (CompanionStruct_D19_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D19) Dtor_cf83() _dafny.Map {
	return _this.Get_().(D19_DC45).Cf83
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && data1.Cf83.Equals(data2.Cf83)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC47 struct {
}

func (D20_DC47) isD20() {}

func (CompanionStruct_D20_) Create_DC47_() D20 {
	return D20{D20_DC47{}}
}

func (_this D20) Is_DC47() bool {
	_, ok := _this.Get_().(D20_DC47)
	return ok
}

type D20_DC46 struct {
	Cf84 T1
}

func (D20_DC46) isD20() {}

func (CompanionStruct_D20_) Create_DC46_(Cf84 T1) D20 {
	return D20{D20_DC46{Cf84}}
}

func (_this D20) Is_DC46() bool {
	_, ok := _this.Get_().(D20_DC46)
	return ok
}

type D20_DC48 struct {
	Cf85 D20
}

func (D20_DC48) isD20() {}

func (CompanionStruct_D20_) Create_DC48_(Cf85 D20) D20 {
	return D20{D20_DC48{Cf85}}
}

func (_this D20) Is_DC48() bool {
	_, ok := _this.Get_().(D20_DC48)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC47_()
}

func (_this D20) Dtor_cf84() T1 {
	return _this.Get_().(D20_DC46).Cf84
}

func (_this D20) Dtor_cf85() D20 {
	return _this.Get_().(D20_DC48).Cf85
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC47:
		{
			return "D20.DC47"
		}
	case D20_DC46:
		{
			return "D20.DC46" + "(" + _dafny.String(data.Cf84) + ")"
		}
	case D20_DC48:
		{
			return "D20.DC48" + "(" + _dafny.String(data.Cf85) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC47:
		{
			_, ok := other.Get_().(D20_DC47)
			return ok
		}
	case D20_DC46:
		{
			data2, ok := other.Get_().(D20_DC46)
			return ok && _dafny.AreEqual(data1.Cf84, data2.Cf84)
		}
	case D20_DC48:
		{
			data2, ok := other.Get_().(D20_DC48)
			return ok && data1.Cf85.Equals(data2.Cf85)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC50 struct {
	Cf87 bool
	Cf88 _dafny.Int
	Cf89 bool
	Cf90 _dafny.Int
}

func (D21_DC50) isD21() {}

func (CompanionStruct_D21_) Create_DC50_(Cf87 bool, Cf88 _dafny.Int, Cf89 bool, Cf90 _dafny.Int) D21 {
	return D21{D21_DC50{Cf87, Cf88, Cf89, Cf90}}
}

func (_this D21) Is_DC50() bool {
	_, ok := _this.Get_().(D21_DC50)
	return ok
}

type D21_DC51 struct {
	Cf91 _dafny.Int
	Cf92 T1
	Cf93 _dafny.Int
	Cf94 _dafny.Int
}

func (D21_DC51) isD21() {}

func (CompanionStruct_D21_) Create_DC51_(Cf91 _dafny.Int, Cf92 T1, Cf93 _dafny.Int, Cf94 _dafny.Int) D21 {
	return D21{D21_DC51{Cf91, Cf92, Cf93, Cf94}}
}

func (_this D21) Is_DC51() bool {
	_, ok := _this.Get_().(D21_DC51)
	return ok
}

type D21_DC52 struct {
	Cf95 _dafny.Int
	Cf96 _dafny.Sequence
}

func (D21_DC52) isD21() {}

func (CompanionStruct_D21_) Create_DC52_(Cf95 _dafny.Int, Cf96 _dafny.Sequence) D21 {
	return D21{D21_DC52{Cf95, Cf96}}
}

func (_this D21) Is_DC52() bool {
	_, ok := _this.Get_().(D21_DC52)
	return ok
}

type D21_DC49 struct {
	Cf86 _dafny.Map
}

func (D21_DC49) isD21() {}

func (CompanionStruct_D21_) Create_DC49_(Cf86 _dafny.Map) D21 {
	return D21{D21_DC49{Cf86}}
}

func (_this D21) Is_DC49() bool {
	_, ok := _this.Get_().(D21_DC49)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC50_(false, _dafny.Zero, false, _dafny.Zero)
}

func (_this D21) Dtor_cf87() bool {
	return _this.Get_().(D21_DC50).Cf87
}

func (_this D21) Dtor_cf88() _dafny.Int {
	return _this.Get_().(D21_DC50).Cf88
}

func (_this D21) Dtor_cf89() bool {
	return _this.Get_().(D21_DC50).Cf89
}

func (_this D21) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D21_DC50).Cf90
}

func (_this D21) Dtor_cf91() _dafny.Int {
	return _this.Get_().(D21_DC51).Cf91
}

func (_this D21) Dtor_cf92() T1 {
	return _this.Get_().(D21_DC51).Cf92
}

func (_this D21) Dtor_cf93() _dafny.Int {
	return _this.Get_().(D21_DC51).Cf93
}

func (_this D21) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D21_DC51).Cf94
}

func (_this D21) Dtor_cf95() _dafny.Int {
	return _this.Get_().(D21_DC52).Cf95
}

func (_this D21) Dtor_cf96() _dafny.Sequence {
	return _this.Get_().(D21_DC52).Cf96
}

func (_this D21) Dtor_cf86() _dafny.Map {
	return _this.Get_().(D21_DC49).Cf86
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC50:
		{
			return "D21.DC50" + "(" + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ")"
		}
	case D21_DC51:
		{
			return "D21.DC51" + "(" + _dafny.String(data.Cf91) + ", " + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ")"
		}
	case D21_DC52:
		{
			return "D21.DC52" + "(" + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D21_DC49:
		{
			return "D21.DC49" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC50:
		{
			data2, ok := other.Get_().(D21_DC50)
			return ok && data1.Cf87 == data2.Cf87 && data1.Cf88.Cmp(data2.Cf88) == 0 && data1.Cf89 == data2.Cf89 && data1.Cf90.Cmp(data2.Cf90) == 0
		}
	case D21_DC51:
		{
			data2, ok := other.Get_().(D21_DC51)
			return ok && data1.Cf91.Cmp(data2.Cf91) == 0 && _dafny.AreEqual(data1.Cf92, data2.Cf92) && data1.Cf93.Cmp(data2.Cf93) == 0 && data1.Cf94.Cmp(data2.Cf94) == 0
		}
	case D21_DC52:
		{
			data2, ok := other.Get_().(D21_DC52)
			return ok && data1.Cf95.Cmp(data2.Cf95) == 0 && data1.Cf96.Equals(data2.Cf96)
		}
	case D21_DC49:
		{
			data2, ok := other.Get_().(D21_DC49)
			return ok && data1.Cf86.Equals(data2.Cf86)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC54 struct {
	Cf98  _dafny.CodePoint
	Cf99  D9
	Cf100 bool
	Cf101 bool
	Cf102 bool
}

func (D22_DC54) isD22() {}

func (CompanionStruct_D22_) Create_DC54_(Cf98 _dafny.CodePoint, Cf99 D9, Cf100 bool, Cf101 bool, Cf102 bool) D22 {
	return D22{D22_DC54{Cf98, Cf99, Cf100, Cf101, Cf102}}
}

func (_this D22) Is_DC54() bool {
	_, ok := _this.Get_().(D22_DC54)
	return ok
}

type D22_DC53 struct {
	Cf97 _dafny.Set
}

func (D22_DC53) isD22() {}

func (CompanionStruct_D22_) Create_DC53_(Cf97 _dafny.Set) D22 {
	return D22{D22_DC53{Cf97}}
}

func (_this D22) Is_DC53() bool {
	_, ok := _this.Get_().(D22_DC53)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC54_(_dafny.CodePoint('D'), Companion_D9_.Default(), false, false, false)
}

func (_this D22) Dtor_cf98() _dafny.CodePoint {
	return _this.Get_().(D22_DC54).Cf98
}

func (_this D22) Dtor_cf99() D9 {
	return _this.Get_().(D22_DC54).Cf99
}

func (_this D22) Dtor_cf100() bool {
	return _this.Get_().(D22_DC54).Cf100
}

func (_this D22) Dtor_cf101() bool {
	return _this.Get_().(D22_DC54).Cf101
}

func (_this D22) Dtor_cf102() bool {
	return _this.Get_().(D22_DC54).Cf102
}

func (_this D22) Dtor_cf97() _dafny.Set {
	return _this.Get_().(D22_DC53).Cf97
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC54:
		{
			return "D22.DC54" + "(" + _dafny.String(data.Cf98) + ", " + _dafny.String(data.Cf99) + ", " + _dafny.String(data.Cf100) + ", " + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ")"
		}
	case D22_DC53:
		{
			return "D22.DC53" + "(" + _dafny.String(data.Cf97) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC54:
		{
			data2, ok := other.Get_().(D22_DC54)
			return ok && data1.Cf98 == data2.Cf98 && data1.Cf99.Equals(data2.Cf99) && data1.Cf100 == data2.Cf100 && data1.Cf101 == data2.Cf101 && data1.Cf102 == data2.Cf102
		}
	case D22_DC53:
		{
			data2, ok := other.Get_().(D22_DC53)
			return ok && data1.Cf97.Equals(data2.Cf97)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC56 struct {
	Cf104 _dafny.Int
	Cf105 _dafny.Int
	Cf106 _dafny.Int
	Cf107 bool
}

func (D23_DC56) isD23() {}

func (CompanionStruct_D23_) Create_DC56_(Cf104 _dafny.Int, Cf105 _dafny.Int, Cf106 _dafny.Int, Cf107 bool) D23 {
	return D23{D23_DC56{Cf104, Cf105, Cf106, Cf107}}
}

func (_this D23) Is_DC56() bool {
	_, ok := _this.Get_().(D23_DC56)
	return ok
}

type D23_DC57 struct {
}

func (D23_DC57) isD23() {}

func (CompanionStruct_D23_) Create_DC57_() D23 {
	return D23{D23_DC57{}}
}

func (_this D23) Is_DC57() bool {
	_, ok := _this.Get_().(D23_DC57)
	return ok
}

type D23_DC55 struct {
	Cf103 *C3
}

func (D23_DC55) isD23() {}

func (CompanionStruct_D23_) Create_DC55_(Cf103 *C3) D23 {
	return D23{D23_DC55{Cf103}}
}

func (_this D23) Is_DC55() bool {
	_, ok := _this.Get_().(D23_DC55)
	return ok
}

type D23_DC58 struct {
	Cf108 D23
}

func (D23_DC58) isD23() {}

func (CompanionStruct_D23_) Create_DC58_(Cf108 D23) D23 {
	return D23{D23_DC58{Cf108}}
}

func (_this D23) Is_DC58() bool {
	_, ok := _this.Get_().(D23_DC58)
	return ok
}

func (CompanionStruct_D23_) Default() D23 {
	return Companion_D23_.Create_DC56_(_dafny.Zero, _dafny.Zero, _dafny.Zero, false)
}

func (_this D23) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D23_DC56).Cf104
}

func (_this D23) Dtor_cf105() _dafny.Int {
	return _this.Get_().(D23_DC56).Cf105
}

func (_this D23) Dtor_cf106() _dafny.Int {
	return _this.Get_().(D23_DC56).Cf106
}

func (_this D23) Dtor_cf107() bool {
	return _this.Get_().(D23_DC56).Cf107
}

func (_this D23) Dtor_cf103() *C3 {
	return _this.Get_().(D23_DC55).Cf103
}

func (_this D23) Dtor_cf108() D23 {
	return _this.Get_().(D23_DC58).Cf108
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC56:
		{
			return "D23.DC56" + "(" + _dafny.String(data.Cf104) + ", " + _dafny.String(data.Cf105) + ", " + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ")"
		}
	case D23_DC57:
		{
			return "D23.DC57"
		}
	case D23_DC55:
		{
			return "D23.DC55" + "(" + _dafny.String(data.Cf103) + ")"
		}
	case D23_DC58:
		{
			return "D23.DC58" + "(" + _dafny.String(data.Cf108) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC56:
		{
			data2, ok := other.Get_().(D23_DC56)
			return ok && data1.Cf104.Cmp(data2.Cf104) == 0 && data1.Cf105.Cmp(data2.Cf105) == 0 && data1.Cf106.Cmp(data2.Cf106) == 0 && data1.Cf107 == data2.Cf107
		}
	case D23_DC57:
		{
			_, ok := other.Get_().(D23_DC57)
			return ok
		}
	case D23_DC55:
		{
			data2, ok := other.Get_().(D23_DC55)
			return ok && data1.Cf103 == data2.Cf103
		}
	case D23_DC58:
		{
			data2, ok := other.Get_().(D23_DC58)
			return ok && data1.Cf108.Equals(data2.Cf108)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of datatype D24
type D24 struct {
	Data_D24_
}

func (_this D24) Get_() Data_D24_ {
	return _this.Data_D24_
}

type Data_D24_ interface {
	isD24()
}

type CompanionStruct_D24_ struct {
}

var Companion_D24_ = CompanionStruct_D24_{}

type D24_DC60 struct {
	Cf110 T1
	Cf111 _dafny.Int
	Cf112 _dafny.Int
	Cf113 bool
}

func (D24_DC60) isD24() {}

func (CompanionStruct_D24_) Create_DC60_(Cf110 T1, Cf111 _dafny.Int, Cf112 _dafny.Int, Cf113 bool) D24 {
	return D24{D24_DC60{Cf110, Cf111, Cf112, Cf113}}
}

func (_this D24) Is_DC60() bool {
	_, ok := _this.Get_().(D24_DC60)
	return ok
}

type D24_DC59 struct {
	Cf109 _dafny.MultiSet
}

func (D24_DC59) isD24() {}

func (CompanionStruct_D24_) Create_DC59_(Cf109 _dafny.MultiSet) D24 {
	return D24{D24_DC59{Cf109}}
}

func (_this D24) Is_DC59() bool {
	_, ok := _this.Get_().(D24_DC59)
	return ok
}

type D24_DC61 struct {
	Cf114 D24
}

func (D24_DC61) isD24() {}

func (CompanionStruct_D24_) Create_DC61_(Cf114 D24) D24 {
	return D24{D24_DC61{Cf114}}
}

func (_this D24) Is_DC61() bool {
	_, ok := _this.Get_().(D24_DC61)
	return ok
}

func (CompanionStruct_D24_) Default() D24 {
	return Companion_D24_.Create_DC60_((T1)(nil), _dafny.Zero, _dafny.Zero, false)
}

func (_this D24) Dtor_cf110() T1 {
	return _this.Get_().(D24_DC60).Cf110
}

func (_this D24) Dtor_cf111() _dafny.Int {
	return _this.Get_().(D24_DC60).Cf111
}

func (_this D24) Dtor_cf112() _dafny.Int {
	return _this.Get_().(D24_DC60).Cf112
}

func (_this D24) Dtor_cf113() bool {
	return _this.Get_().(D24_DC60).Cf113
}

func (_this D24) Dtor_cf109() _dafny.MultiSet {
	return _this.Get_().(D24_DC59).Cf109
}

func (_this D24) Dtor_cf114() D24 {
	return _this.Get_().(D24_DC61).Cf114
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC60:
		{
			return "D24.DC60" + "(" + _dafny.String(data.Cf110) + ", " + _dafny.String(data.Cf111) + ", " + _dafny.String(data.Cf112) + ", " + _dafny.String(data.Cf113) + ")"
		}
	case D24_DC59:
		{
			return "D24.DC59" + "(" + _dafny.String(data.Cf109) + ")"
		}
	case D24_DC61:
		{
			return "D24.DC61" + "(" + _dafny.String(data.Cf114) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC60:
		{
			data2, ok := other.Get_().(D24_DC60)
			return ok && _dafny.AreEqual(data1.Cf110, data2.Cf110) && data1.Cf111.Cmp(data2.Cf111) == 0 && data1.Cf112.Cmp(data2.Cf112) == 0 && data1.Cf113 == data2.Cf113
		}
	case D24_DC59:
		{
			data2, ok := other.Get_().(D24_DC59)
			return ok && data1.Cf109.Equals(data2.Cf109)
		}
	case D24_DC61:
		{
			data2, ok := other.Get_().(D24_DC61)
			return ok && data1.Cf114.Equals(data2.Cf114)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D24) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D24)
	return ok && _this.Equals(typed)
}

func Type_D24_() _dafny.TypeDescriptor {
	return type_D24_{}
}

type type_D24_ struct {
}

func (_this type_D24_) Default() interface{} {
	return Companion_D24_.Default()
}

func (_this type_D24_) String() string {
	return "main.D24"
}
func (_this D24) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D24{}

// End of datatype D24

// Definition of datatype D25
type D25 struct {
	Data_D25_
}

func (_this D25) Get_() Data_D25_ {
	return _this.Data_D25_
}

type Data_D25_ interface {
	isD25()
}

type CompanionStruct_D25_ struct {
}

var Companion_D25_ = CompanionStruct_D25_{}

type D25_DC63 struct {
	Cf116 _dafny.Sequence
	Cf117 bool
	Cf118 _dafny.Int
}

func (D25_DC63) isD25() {}

func (CompanionStruct_D25_) Create_DC63_(Cf116 _dafny.Sequence, Cf117 bool, Cf118 _dafny.Int) D25 {
	return D25{D25_DC63{Cf116, Cf117, Cf118}}
}

func (_this D25) Is_DC63() bool {
	_, ok := _this.Get_().(D25_DC63)
	return ok
}

type D25_DC62 struct {
	Cf115 *C8
}

func (D25_DC62) isD25() {}

func (CompanionStruct_D25_) Create_DC62_(Cf115 *C8) D25 {
	return D25{D25_DC62{Cf115}}
}

func (_this D25) Is_DC62() bool {
	_, ok := _this.Get_().(D25_DC62)
	return ok
}

type D25_DC64 struct {
	Cf119 D25
}

func (D25_DC64) isD25() {}

func (CompanionStruct_D25_) Create_DC64_(Cf119 D25) D25 {
	return D25{D25_DC64{Cf119}}
}

func (_this D25) Is_DC64() bool {
	_, ok := _this.Get_().(D25_DC64)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC63_(_dafny.EmptySeq, false, _dafny.Zero)
}

func (_this D25) Dtor_cf116() _dafny.Sequence {
	return _this.Get_().(D25_DC63).Cf116
}

func (_this D25) Dtor_cf117() bool {
	return _this.Get_().(D25_DC63).Cf117
}

func (_this D25) Dtor_cf118() _dafny.Int {
	return _this.Get_().(D25_DC63).Cf118
}

func (_this D25) Dtor_cf115() *C8 {
	return _this.Get_().(D25_DC62).Cf115
}

func (_this D25) Dtor_cf119() D25 {
	return _this.Get_().(D25_DC64).Cf119
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC63:
		{
			return "D25.DC63" + "(" + data.Cf116.VerbatimString(true) + ", " + _dafny.String(data.Cf117) + ", " + _dafny.String(data.Cf118) + ")"
		}
	case D25_DC62:
		{
			return "D25.DC62" + "(" + _dafny.String(data.Cf115) + ")"
		}
	case D25_DC64:
		{
			return "D25.DC64" + "(" + _dafny.String(data.Cf119) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC63:
		{
			data2, ok := other.Get_().(D25_DC63)
			return ok && data1.Cf116.Equals(data2.Cf116) && data1.Cf117 == data2.Cf117 && data1.Cf118.Cmp(data2.Cf118) == 0
		}
	case D25_DC62:
		{
			data2, ok := other.Get_().(D25_DC62)
			return ok && data1.Cf115 == data2.Cf115
		}
	case D25_DC64:
		{
			data2, ok := other.Get_().(D25_DC64)
			return ok && data1.Cf119.Equals(data2.Cf119)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D25) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D25)
	return ok && _this.Equals(typed)
}

func Type_D25_() _dafny.TypeDescriptor {
	return type_D25_{}
}

type type_D25_ struct {
}

func (_this type_D25_) Default() interface{} {
	return Companion_D25_.Default()
}

func (_this type_D25_) String() string {
	return "main.D25"
}
func (_this D25) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D25{}

// End of datatype D25

// Definition of datatype D26
type D26 struct {
	Data_D26_
}

func (_this D26) Get_() Data_D26_ {
	return _this.Data_D26_
}

type Data_D26_ interface {
	isD26()
}

type CompanionStruct_D26_ struct {
}

var Companion_D26_ = CompanionStruct_D26_{}

type D26_DC66 struct {
	Cf121 _dafny.Int
	Cf122 bool
	Cf123 *C3
}

func (D26_DC66) isD26() {}

func (CompanionStruct_D26_) Create_DC66_(Cf121 _dafny.Int, Cf122 bool, Cf123 *C3) D26 {
	return D26{D26_DC66{Cf121, Cf122, Cf123}}
}

func (_this D26) Is_DC66() bool {
	_, ok := _this.Get_().(D26_DC66)
	return ok
}

type D26_DC65 struct {
	Cf120 _dafny.MultiSet
}

func (D26_DC65) isD26() {}

func (CompanionStruct_D26_) Create_DC65_(Cf120 _dafny.MultiSet) D26 {
	return D26{D26_DC65{Cf120}}
}

func (_this D26) Is_DC65() bool {
	_, ok := _this.Get_().(D26_DC65)
	return ok
}

type D26_DC67 struct {
	Cf124 D26
}

func (D26_DC67) isD26() {}

func (CompanionStruct_D26_) Create_DC67_(Cf124 D26) D26 {
	return D26{D26_DC67{Cf124}}
}

func (_this D26) Is_DC67() bool {
	_, ok := _this.Get_().(D26_DC67)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC66_(_dafny.Zero, false, (*C3)(nil))
}

func (_this D26) Dtor_cf121() _dafny.Int {
	return _this.Get_().(D26_DC66).Cf121
}

func (_this D26) Dtor_cf122() bool {
	return _this.Get_().(D26_DC66).Cf122
}

func (_this D26) Dtor_cf123() *C3 {
	return _this.Get_().(D26_DC66).Cf123
}

func (_this D26) Dtor_cf120() _dafny.MultiSet {
	return _this.Get_().(D26_DC65).Cf120
}

func (_this D26) Dtor_cf124() D26 {
	return _this.Get_().(D26_DC67).Cf124
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC66:
		{
			return "D26.DC66" + "(" + _dafny.String(data.Cf121) + ", " + _dafny.String(data.Cf122) + ", " + _dafny.String(data.Cf123) + ")"
		}
	case D26_DC65:
		{
			return "D26.DC65" + "(" + _dafny.String(data.Cf120) + ")"
		}
	case D26_DC67:
		{
			return "D26.DC67" + "(" + _dafny.String(data.Cf124) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC66:
		{
			data2, ok := other.Get_().(D26_DC66)
			return ok && data1.Cf121.Cmp(data2.Cf121) == 0 && data1.Cf122 == data2.Cf122 && data1.Cf123 == data2.Cf123
		}
	case D26_DC65:
		{
			data2, ok := other.Get_().(D26_DC65)
			return ok && data1.Cf120.Equals(data2.Cf120)
		}
	case D26_DC67:
		{
			data2, ok := other.Get_().(D26_DC67)
			return ok && data1.Cf124.Equals(data2.Cf124)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D26) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D26)
	return ok && _this.Equals(typed)
}

func Type_D26_() _dafny.TypeDescriptor {
	return type_D26_{}
}

type type_D26_ struct {
}

func (_this type_D26_) Default() interface{} {
	return Companion_D26_.Default()
}

func (_this type_D26_) String() string {
	return "main.D26"
}
func (_this D26) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D26{}

// End of datatype D26

// Definition of trait T0
type T0 interface {
	String() string
	Fm1(globalState *GlobalState) _dafny.Int
	Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool
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
	Fm1(globalState *GlobalState) _dafny.Int
	Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool
	M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState)
	F26() bool
	F27() bool
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

// Definition of trait T2
type T2 interface {
	String() string
	Fm1(globalState *GlobalState) _dafny.Int
	Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool
	M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState)
	F26() bool
	F27() bool
	Fm3(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	Fm4(p0 D0, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int
	M2(p0 bool, p1 T0, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int)
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of trait T3
type T3 interface {
	String() string
	M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Array, T0)
	M4(globalState *GlobalState) (bool, bool)
	F29() _dafny.MultiSet
}
type CompanionStruct_T3_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T3_ = CompanionStruct_T3_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T3_) CastTo_(x interface{}) T3 {
	var t T3
	t, _ = x.(T3)
	return t
}

// End of trait T3

// Definition of class GlobalState
type GlobalState struct {
	F2   bool
	F5   _dafny.Array
	F6   _dafny.Sequence
	F12  _dafny.Int
	F15  bool
	F16  _dafny.Int
	F19  _dafny.CodePoint
	F22  bool
	F24  bool
	_f0  _dafny.Set
	_f1  _dafny.MultiSet
	_f3  bool
	_f4  _dafny.CodePoint
	_f7  _dafny.Map
	_f8  bool
	_f9  bool
	_f10 _dafny.Int
	_f11 bool
	_f13 _dafny.Int
	_f14 _dafny.Array
	_f17 bool
	_f18 bool
	_f20 _dafny.Sequence
	_f21 _dafny.Int
	_f23 _dafny.Int
	_f25 _dafny.Map
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = false
	_this.F5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F6 = _dafny.EmptySeq
	_this.F12 = _dafny.Zero
	_this.F15 = false
	_this.F16 = _dafny.Zero
	_this.F19 = _dafny.CodePoint('D')
	_this.F22 = false
	_this.F24 = false
	_this._f0 = _dafny.EmptySet
	_this._f1 = _dafny.EmptyMultiSet
	_this._f3 = false
	_this._f4 = _dafny.CodePoint('D')
	_this._f7 = _dafny.EmptyMap
	_this._f8 = false
	_this._f9 = false
	_this._f10 = _dafny.Zero
	_this._f11 = false
	_this._f13 = _dafny.Zero
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f17 = false
	_this._f18 = false
	_this._f20 = _dafny.EmptySeq
	_this._f21 = _dafny.Zero
	_this._f23 = _dafny.Zero
	_this._f25 = _dafny.EmptyMap
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

func (_this *GlobalState) Ctor__(f0 _dafny.Set, f1 _dafny.MultiSet, f2 bool, f3 bool, f4 _dafny.CodePoint, f5 _dafny.Array, f6 _dafny.Sequence, f7 _dafny.Map, f8 bool, f9 bool, f10 _dafny.Int, f11 bool, f12 _dafny.Int, f13 _dafny.Int, f14 _dafny.Array, f15 bool, f16 _dafny.Int, f17 bool, f18 bool, f19 _dafny.CodePoint, f20 _dafny.Sequence, f21 _dafny.Int, f22 bool, f23 _dafny.Int, f24 bool, f25 _dafny.Map) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f14 = f14
		(_this).F15 = f15
		(_this).F16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
		(_this)._f23 = f23
		(_this).F24 = f24
		(_this)._f25 = f25
	}
}
func (_this *GlobalState) F0() _dafny.Set {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.MultiSet {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.CodePoint {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F7() _dafny.Map {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() bool {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F13() _dafny.Int {
	{
		return _this._f13
	}
}
func (_this *GlobalState) F14() _dafny.Array {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F17() bool {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F20() _dafny.Sequence {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Int {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F23() _dafny.Int {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F25() _dafny.Map {
	{
		return _this._f25
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f32 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f32 = false
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

func (_this *C0) Ctor__(f32 bool) {
	{
		(_this)._f32 = f32
	}
}
func (_this *C0) Fm6(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) bool {
	{
		return (_this).F32()
	}
}
func (_this *C0) Fm7(p0 _dafny.Map, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C0) F32() bool {
	{
		return _this._f32
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f26 bool
	_f27 bool
	F37  _dafny.Int
	_f36 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f26 = false
	_this._f27 = false
	_this.F37 = _dafny.Zero
	_this._f36 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C1{}
var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F26() bool {
	return _this._f26
}
func (_this *C1) F27() bool {
	return _this._f27
}
func (_this *C1) Ctor__(f36 _dafny.Int, f37 _dafny.Int, f26 bool, f27 bool) {
	{
		(_this)._f36 = f36
		(_this).F37 = f37
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C1) Fm3(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26()), _dafny.SeqOf((_this).F26())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F27()), _dafny.SeqOf((_this).F27(), (_this).F26())))
	}
}
func (_this *C1) Fm4(p0 D0, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F36()
	}
}
func (_this *C1) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return _this.F37
	}
}
func (_this *C1) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F26()
	}
}
func (_this *C1) M2(p0 bool, p1 T0, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _387_i0 _dafny.Int
		_ = _387_i0
		_387_i0 = _dafny.Zero
		{
			for true {
				{
					if (_387_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_387_i0 = (_387_i0).Plus(_dafny.One)
					(globalState).F12 = _this.F37
					(globalState).F24 = !(p2) || ((_this).Fm2((_this).F26(), _this.F37, globalState))
					var _388_v0 _dafny.Array
					_ = _388_v0
					var _len0_9 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_9
					var _nw43 _dafny.Array
					_ = _nw43
					if _len0_9.Cmp(_dafny.Zero) == 0 {
						_nw43 = _dafny.NewArray(_len0_9)
					} else {
						var _init9 func(_dafny.Int) _dafny.Sequence = func(_389_i1 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf((_this).F36(), _this.F37)
						}
						_ = _init9
						var _element0_9 = _init9(_dafny.Zero)
						_ = _element0_9
						_nw43 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
						(_nw43).ArraySet1(_element0_9, 0)
						var _nativeLen0_9 = (_len0_9).Int()
						_ = _nativeLen0_9
						for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
							(_nw43).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
						}
					}
					_388_v0 = _nw43
					var _390_v1 _dafny.Sequence
					_ = _390_v1
					_390_v1 = _dafny.SeqOf(_this.F37, (_this).F36())
					var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_388_v0), 0))
					_ = _index41
					(_388_v0).ArraySet1(_390_v1, (_index41).Int())
					var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_388_v0), 0))
					_ = _index42
					(_388_v0).ArraySet1(_390_v1, (_index42).Int())
					(globalState).F2 = p2
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _391_v2 _dafny.Map
		_ = _391_v2
		_391_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F27()), !(p2) || (p2))
		var _392_v3 _dafny.Array
		_ = _392_v3
		var _nw44 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
		_ = _nw44
		_392_v3 = _nw44
		var _393_v4 *C0
		_ = _393_v4
		var _nw45 *C0 = New_C0_()
		_ = _nw45
		_nw45.Ctor__((_this).F27())
		_393_v4 = _nw45
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_392_v3), 0))
		_ = _index43
		(_392_v3).ArraySet1(_393_v4, (_index43).Int())
		var _394_v5 _dafny.Set
		_ = _394_v5
		_394_v5 = _dafny.SetOf(_this.F37, _dafny.IntOfInt64(975))
		var _395_v6 _dafny.Array
		_ = _395_v6
		var _nw46 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw46
		_395_v6 = _nw46
		var _396_v7 _dafny.MultiSet
		_ = _396_v7
		_396_v7 = _dafny.MultiSetOf((_this).F36())
		var _397_v8 _dafny.Map
		_ = _397_v8
		_397_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_395_v6, (func() _dafny.Int {
			if (_396_v7).Contains((_this).F36()) {
				return (_396_v7).Multiplicity((_this).F36())
			}
			return _this.F37
		})())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_392_v3), 0))
		_ = _index44
		var _rhs39 _dafny.Int = ((_dafny.IntOfInt64(-338)).Minus((_394_v5).Cardinality())).Plus((_this).Fm4(Companion_Default___.Fm19((_this).F36(), (_397_v8).Cardinality(), globalState), _dafny.IntOfInt64(-144), (_this).F36(), globalState))
		_ = _rhs39
		var _rhs40 _dafny.Map = _391_v2
		_ = _rhs40
		var _rhs41 *C0 = _393_v4
		_ = _rhs41
		var _lhs35 *C1 = _this
		_ = _lhs35
		var _lhs36 _dafny.Array = _392_v3
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_392_v3), 0))
		_ = _lhs37
		_lhs35.F37 = _rhs39
		_391_v2 = _rhs40
		(_lhs36).ArraySet1(_rhs41, (_lhs37).Int())
		var _398_v9 D2
		_ = _398_v9
		_398_v9 = Companion_D2_.Create_DC7_(p2, (_this).F27(), _dafny.IntOfInt64(-868), (_this).F36())
		var _399_v10 _dafny.Map
		_ = _399_v10
		_399_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F37, (_398_v9).Dtor_cf17())
		var _400_v11 D2
		_ = _400_v11
		_400_v11 = Companion_D2_.Create_DC6_(_399_v10)
		var _hi2 _dafny.Int = (_dafny.SetOf(_400_v11)).Cardinality()
		_ = _hi2
		for _401_i2 := _this.F37; _401_i2.Cmp(_hi2) < 0; _401_i2 = _401_i2.Plus(_dafny.One) {
			_394_v5 = _394_v5
			var _402_v12 _dafny.Array
			_ = _402_v12
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw47
			_402_v12 = _nw47
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_402_v12), 0))
			_ = _index45
			(_402_v12).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(796), (_this).F36()), (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_402_v12), 0))
			_ = _index46
			(_402_v12).ArraySet1(Companion_Default___.SafeModuloInt(_401_i2, _401_i2), (_index46).Int())
			var _403_v13 _dafny.CodePoint
			_ = _403_v13
			_403_v13 = _dafny.CodePoint('a')
			var _404_v14 _dafny.Array
			_ = _404_v14
			var _nwElement0_5 _dafny.CodePoint = _403_v13
			_ = _nwElement0_5
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(3))
			_ = _nw48
			(_nw48).ArraySet1CodePoint(_nwElement0_5, 0)
			(_nw48).ArraySet1CodePoint(_403_v13, 1)
			(_nw48).ArraySet1CodePoint(_403_v13, 2)
			_404_v14 = _nw48
			var _405_v15 _dafny.Map
			_ = _405_v15
			_405_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_402_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_402_v12), 0))).Int()).(_dafny.Int), _404_v14)
			_405_v15 = (_405_v15).Update(_401_i2, (func() _dafny.Array {
				if p0 {
					return _404_v14
				}
				return _404_v14
			})())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_395_v6), 0))
			_ = _index47
			(_395_v6).ArraySet1((_this).F26(), (_index47).Int())
			var _406_v16 _dafny.Sequence
			_ = _406_v16
			_406_v16 = _dafny.UnicodeSeqOfUtf8Bytes("tniumeoec")
			var _407_v17 _dafny.Map
			_ = _407_v17
			_407_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_406_v16).Cardinality())), _396_v7)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(955), _dafny.ArrayLen((_395_v6), 0))
			_ = _index48
			(_395_v6).ArraySet1(((_402_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_402_v12), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((((func() _dafny.MultiSet {
				if (_407_v17).Contains((_this).F36()) {
					return (_407_v17).Get((_this).F36()).(_dafny.MultiSet)
				}
				return _396_v7
			})()).Intersection(_396_v7)).Cardinality())) > 0, (_index48).Int())
		}
		(_this).F37 = _this.F37
		var _408_v18 _dafny.Sequence
		_ = _408_v18
		_408_v18 = _dafny.SeqOf((_this).F26())
		var _409_v19 _dafny.Sequence
		_ = _409_v19
		_409_v19 = _dafny.SeqOf(_dafny.IntOfUint32((_408_v18).Cardinality()))
		var _410_v20 _dafny.Sequence
		_ = _410_v20
		_410_v20 = _409_v19
		var _411_v21 _dafny.Map
		_ = _411_v21
		_411_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.Companion_Sequence_.Concatenate((_410_v20), _dafny.SeqOf(_dafny.IntOfUint32((_408_v18).Cardinality()))))
		var _412_v22 _dafny.Array
		_ = _412_v22
		var _len0_10 _dafny.Int = _dafny.One
		_ = _len0_10
		var _nw49 _dafny.Array
		_ = _nw49
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw49 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Sequence = (func(_413_v4 *C0, _414_v18 _dafny.Sequence, _415_p0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_416_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(Companion_D2_.Create_DC9_((_this).F36(), _this.F37, (_413_v4).F32())), _dafny.MultiSetOf(Companion_D2_.Create_DC9_((_this).F36(), (_this).F36(), true), Companion_D2_.Create_DC9_((_this).F36(), (_this).F36(), (_this).F27()), Companion_D2_.Create_DC9_(_this.F37, _this.F37, (_413_v4).F32()))), _dafny.SeqOf(_dafny.MultiSetOf(Companion_D2_.Create_DC9_((_this).F36(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_414_v18), (_413_v4).F32())).Cardinality())).Cardinality()), (_413_v4).F32()), Companion_D2_.Create_DC9_(_dafny.IntOfUint32((_414_v18).Cardinality()), _this.F37, _415_p0))))
				}
			})(_393_v4, _408_v18, p0)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw49 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw49).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw49).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_412_v22 = _nw49
		var _417_v23 D2
		_ = _417_v23
		_417_v23 = Companion_D2_.Create_DC9_((_this).F36(), _this.F37, (_this).F27())
		var _418_v24 _dafny.MultiSet
		_ = _418_v24
		_418_v24 = _dafny.MultiSetOf(_417_v23)
		var _419_v25 _dafny.MultiSet
		_ = _419_v25
		_419_v25 = _dafny.MultiSetOf(p2)
		var _420_v26 _dafny.Sequence
		_ = _420_v26
		_420_v26 = _dafny.SeqOf(_418_v24, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_417_v23), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_419_v25).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_417_v23)).Cardinality()))).Uint32(), Companion_Default___.Fm20((_this).F36(), p2, (_this).F36(), (_398_v9).Dtor_cf17(), globalState))))
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_412_v22), 0))
		_ = _index49
		(_412_v22).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_420_v26, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(62))).Uint32(), func(coer39 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}((func(_421_v24 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
			return func(_422_i4 _dafny.Int) _dafny.MultiSet {
				return _421_v24
			}
		})(_418_v24)))), (_index49).Int())
		var _423_v27 _dafny.Map
		_ = _423_v27
		_423_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F37, _418_v24)
		var _424_v28 D0
		_ = _424_v28
		_424_v28 = Companion_D0_.Create_DC0_((_this).F26())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_412_v22), 0))
		_ = _index50
		var _rhs42 _dafny.Map = _411_v21
		_ = _rhs42
		var _rhs43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_420_v26, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-716), _dafny.IntOfUint32((_420_v26).Cardinality()))).Uint32(), (func() _dafny.MultiSet {
			if (_423_v27).Contains(_this.F37) {
				return (_423_v27).Get(_this.F37).(_dafny.MultiSet)
			}
			return _418_v24
		})()), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_420_v26, (Companion_Default___.SafeIndex((_this).Fm4(_424_v28, _this.F37, (_this).F36(), globalState), _dafny.IntOfUint32((_420_v26).Cardinality()))).Uint32(), _418_v24), (Companion_Default___.SafeIndex((_this).F36(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_420_v26, (Companion_Default___.SafeIndex((_this).Fm4(_424_v28, _this.F37, (_this).F36(), globalState), _dafny.IntOfUint32((_420_v26).Cardinality()))).Uint32(), _418_v24)).Cardinality()))).Uint32(), _418_v24), _dafny.SeqOf(_418_v24, _418_v24)))
		_ = _rhs43
		var _lhs38 _dafny.Array = _412_v22
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_412_v22), 0))
		_ = _lhs39
		_411_v21 = _rhs42
		(_lhs38).ArraySet1(_rhs43, (_lhs39).Int())
		var _425_v29 _dafny.Array
		_ = _425_v29
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_11
		var _nw50 _dafny.Array
		_ = _nw50
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw50 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.MultiSet = (func(_426_v25 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_427_i5 _dafny.Int) _dafny.MultiSet {
					return _426_v25
				}
			})(_419_v25)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw50 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw50).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw50).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_425_v29 = _nw50
		_425_v29 = _425_v29
		var _428_v30 _dafny.Map
		_ = _428_v30
		_428_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.CodePoint('c'))
		var _429_v31 _dafny.CodePoint
		_ = _429_v31
		_429_v31 = _dafny.CodePoint('m')
		r0 = (func() _dafny.CodePoint {
			if (_428_v30).Contains((_393_v4).F32()) {
				return (_428_v30).Get((_393_v4).F32()).(_dafny.CodePoint)
			}
			return _429_v31
		})()
		var _430_v32 _dafny.Sequence
		_ = _430_v32
		_430_v32 = _dafny.SeqOf(_429_v31, _429_v31)
		var _431_v33 _dafny.Map
		_ = _431_v33
		_431_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_430_v32).Cardinality()), (_399_v10).Update(_this.F37, (_this).F36()))
		r1 = (_431_v33).Cardinality()
		return r0, r1
	}
}
func (_this *C1) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _432_v0 _dafny.Array
		_ = _432_v0
		var _nw51 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
		_ = _nw51
		_432_v0 = _nw51
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))
		_ = _index51
		(_432_v0).ArraySet1((_this).F26(), (_index51).Int())
		var _433_v1 _dafny.Array
		_ = _433_v1
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw52
		_433_v1 = _nw52
		var _434_v2 _dafny.Set
		_ = _434_v2
		_434_v2 = _dafny.SetOf((_this).F26())
		var _435_v3 _dafny.Sequence
		_ = _435_v3
		_435_v3 = _dafny.SeqOf(_434_v2)
		var _436_v4 _dafny.Map
		_ = _436_v4
		_436_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_435_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.IntOfUint32((_435_v3).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), (_this).F26())
		var _437_v5 _dafny.Map
		_ = _437_v5
		_437_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_dafny.Zero).Minus((_436_v4).Cardinality()))
		var _438_v6 _dafny.MultiSet
		_ = _438_v6
		_438_v6 = _dafny.MultiSetOf((func() _dafny.Int {
			if (_437_v5).Contains(p0) {
				return (_437_v5).Get(p0).(_dafny.Int)
			}
			return _this.F37
		})(), _this.F37)
		var _439_v7 _dafny.Sequence
		_ = _439_v7
		_439_v7 = _dafny.UnicodeSeqOfUtf8Bytes("k")
		var _440_v8 _dafny.CodePoint
		_ = _440_v8
		_440_v8 = _dafny.CodePoint('t')
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index52
		(_433_v1).ArraySet1((func() _dafny.Int {
			if (_438_v6).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_439_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.IntOfUint32((_439_v7).Cardinality()))).Uint32(), _440_v8)).Cardinality())) {
				return (_438_v6).Multiplicity(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_439_v7, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.IntOfUint32((_439_v7).Cardinality()))).Uint32(), _440_v8)).Cardinality()))
			}
			return (_this).Fm1(globalState)
		})(), (_index52).Int())
		var _441_v9 _dafny.Array
		_ = _441_v9
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_12
		var _nw53 _dafny.Array
		_ = _nw53
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw53 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.CodePoint = (func(_442_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_443_i0 _dafny.Int) _dafny.CodePoint {
					return _442_v8
				}
			})(_440_v8)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw53 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw53).ArraySet1CodePoint(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw53).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_441_v9 = _nw53
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))
		_ = _index53
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index54
		var _rhs44 bool = p0
		_ = _rhs44
		var _rhs45 _dafny.Int = _this.F37
		_ = _rhs45
		var _rhs46 _dafny.Array = _441_v9
		_ = _rhs46
		var _lhs40 _dafny.Array = _432_v0
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))
		_ = _lhs41
		var _lhs42 _dafny.Array = _433_v1
		_ = _lhs42
		var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _lhs43
		(_lhs40).ArraySet1(_rhs44, (_lhs41).Int())
		(_lhs42).ArraySet1(_rhs45, (_lhs43).Int())
		_441_v9 = _rhs46
		var _444_v10 _dafny.MultiSet
		_ = _444_v10
		_444_v10 = _dafny.MultiSetOf(p0, (_this).F27())
		var _445_v11 _dafny.Sequence
		_ = _445_v11
		_445_v11 = _dafny.SeqOf(_this.F37, _dafny.IntOfUint32((_439_v7).Cardinality()))
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index55
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index56
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index57
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))
		_ = _index58
		var _rhs47 _dafny.Int = _this.F37
		_ = _rhs47
		var _rhs48 _dafny.Int = (((_444_v10).Union(_444_v10)).Union(_444_v10)).Cardinality()
		_ = _rhs48
		var _rhs49 _dafny.Int = _dafny.IntOfInt64(-718)
		_ = _rhs49
		var _rhs50 bool = (func() bool {
			if (func() bool {
				if p0 {
					return p0
				}
				return !(p0)
			})() {
				return !_dafny.Companion_Sequence_.Equal(_445_v11, _445_v11)
			}
			return ((func() _dafny.Int {
				if (_437_v5).Contains((_this).F26()) {
					return (_437_v5).Get((_this).F26()).(_dafny.Int)
				}
				return (_437_v5).Cardinality()
			})()).Cmp(_dafny.IntOfInt64(349)) < 0
		})()
		_ = _rhs50
		var _rhs51 bool = (_this).F26()
		_ = _rhs51
		var _lhs44 _dafny.Array = _433_v1
		_ = _lhs44
		var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _lhs45
		var _lhs46 _dafny.Array = _433_v1
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _lhs47
		var _lhs48 _dafny.Array = _433_v1
		_ = _lhs48
		var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _lhs49
		var _lhs50 *GlobalState = globalState
		_ = _lhs50
		var _lhs51 _dafny.Array = _432_v0
		_ = _lhs51
		var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))
		_ = _lhs52
		(_lhs44).ArraySet1(_rhs47, (_lhs45).Int())
		(_lhs46).ArraySet1(_rhs48, (_lhs47).Int())
		(_lhs48).ArraySet1(_rhs49, (_lhs49).Int())
		_lhs50.F22 = _rhs50
		(_lhs51).ArraySet1(_rhs51, (_lhs52).Int())
		var _hi3 _dafny.Int = (_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int)
		_ = _hi3
		for _446_i1 := _this.F37; _446_i1.Cmp(_hi3) < 0; _446_i1 = _446_i1.Plus(_dafny.One) {
			var _447_v12 *C0
			_ = _447_v12
			var _nw54 *C0 = New_C0_()
			_ = _nw54
			_nw54.Ctor__((_444_v10).IsSubsetOf(Companion_Default___.Fm21(_this.F37, globalState)))
			_447_v12 = _nw54
			var _448_v13 _dafny.Set
			_ = _448_v13
			_448_v13 = _dafny.SetOf(_this.F37, (_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int), (_this).F36(), (_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int))
			var _449_v14 _dafny.Map
			_ = _449_v14
			_449_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_448_v13, _432_v0)
			_449_v14 = _449_v14
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
			_ = _index59
			(_433_v1).ArraySet1(_446_i1, (_index59).Int())
			(globalState).F15 = (p2).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_this.F37, (_434_v2).Cardinality()), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)
		}
		var _450_v15 _dafny.Map
		_ = _450_v15
		_450_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v2, !((_this).F26()))
		_450_v15 = (func() _dafny.Map {
			if p0 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v2, (_432_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(254), _dafny.ArrayLen((_432_v0), 0))).Int()).(bool))
			}
			return _450_v15
		})()
		var _hi4 _dafny.Int = (_this.F37).Times((_dafny.Zero).Minus((_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int)))
		_ = _hi4
		for _451_i2 := (_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int); _451_i2.Cmp(_hi4) < 0; _451_i2 = _451_i2.Plus(_dafny.One) {
			var _452_v16 *C0
			_ = _452_v16
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__((_this).F27())
			_452_v16 = _nw55
			(globalState).F12 = ((_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int)).Plus(_451_i2)
			var _453_v17 D1
			_ = _453_v17
			_453_v17 = Companion_D1_.Create_DC4_(_441_v9, (_433_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf(_440_v8, _440_v8)).Cardinality(), Companion_Default___.Fm22(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(344))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}((func(_454_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_455_i3 _dafny.Int) _dafny.CodePoint {
					return _454_v8
				}
			})(_440_v8)))).Cardinality()), (_this).F27(), globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(209))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_456_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_457_i4 _dafny.Int) _dafny.CodePoint {
					return _456_v8
				}
			})(_440_v8))))
			_453_v17 = (func() D1 {
				if (_this).F26() {
					return _453_v17
				}
				return Companion_D1_.Create_DC4_(_441_v9, _dafny.IntOfUint32(((_453_v17).Dtor_cf11()).Cardinality()), _this.F37, _439_v7, _439_v7)
			})()
			(globalState).F15 = false
		}
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_433_v1), 0))
		_ = _index60
		(_433_v1).ArraySet1(_this.F37, (_index60).Int())
	}
}
func (_this *C1) F36() _dafny.Int {
	{
		return _this._f36
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f26 bool
	_f27 bool
	_f35 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f26 = false
	_this._f27 = false
	_this._f35 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C2{}
var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F26() bool {
	return _this._f26
}
func (_this *C2) F27() bool {
	return _this._f27
}
func (_this *C2) Ctor__(f35 _dafny.Sequence, f26 bool, f27 bool) {
	{
		(_this)._f35 = f35
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C2) Fm3(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((func() _dafny.Set {
			var _coll48 = _dafny.NewBuilder()
			_ = _coll48
			for _iter48 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}(func(_458_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(589), true)).Cardinality())
			}))).Elements()); ; {
				_compr_48, _ok48 := _iter48()
				if !_ok48 {
					break
				}
				var _459_v0 _dafny.Int
				_459_v0 = interface{}(_compr_48).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}(func(_458_i0 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(589), true)).Cardinality())
				})), _459_v0) {
					_coll48.Add(Companion_Default___.SafeDivisionInt(_459_v0, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false, false)).Cardinality())).Cardinality())))
				}
			}
			return _coll48.ToSet()
		}()).Cardinality())).Cardinality(), (_this).F27())).Cardinality(), (_dafny.SetOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F26(), (_this).F27(), (_this).F26())).Cardinality()))).Cardinality())).Equals(_dafny.SetOf((Companion_D0_.Create_DC1_(false, (_this).F26(), _dafny.IntOfInt64(832), (_this).F26())).Dtor_cf3())) {
			return _dafny.SeqOf((_this).F26(), (_this).F27())
		} else {
			return _dafny.SeqOf((_this).F26(), (_this).F26())
		}
	}
}
func (_this *C2) Fm4(p0 D0, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (func(_source10 D2) _dafny.Map {
			if _source10.Is_DC7() {
				var _460___mcc_h0 bool = _source10.Get_().(D2_DC7).Cf14
				_ = _460___mcc_h0
				var _461___mcc_h1 bool = _source10.Get_().(D2_DC7).Cf15
				_ = _461___mcc_h1
				var _462___mcc_h2 _dafny.Int = _source10.Get_().(D2_DC7).Cf16
				_ = _462___mcc_h2
				var _463___mcc_h3 _dafny.Int = _source10.Get_().(D2_DC7).Cf17
				_ = _463___mcc_h3
				var _464_cf17 _dafny.Int = _463___mcc_h3
				_ = _464_cf17
				var _465_cf16 _dafny.Int = _462___mcc_h2
				_ = _465_cf16
				var _466_cf15 bool = _461___mcc_h1
				_ = _466_cf15
				var _467_cf14 bool = _460___mcc_h0
				_ = _467_cf14
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.MultiSetOf(_dafny.CodePoint('t'), _dafny.CodePoint('g')))
			} else if _source10.Is_DC8() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.MultiSetOf(_dafny.CodePoint('g')))
			} else if _source10.Is_DC9() {
				var _468___mcc_h4 _dafny.Int = _source10.Get_().(D2_DC9).Cf18
				_ = _468___mcc_h4
				var _469___mcc_h5 _dafny.Int = _source10.Get_().(D2_DC9).Cf19
				_ = _469___mcc_h5
				var _470___mcc_h6 bool = _source10.Get_().(D2_DC9).Cf20
				_ = _470___mcc_h6
				var _471_cf20 bool = _470___mcc_h6
				_ = _471_cf20
				var _472_cf19 _dafny.Int = _469___mcc_h5
				_ = _472_cf19
				var _473_cf18 _dafny.Int = _468___mcc_h4
				_ = _473_cf18
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_dafny.CodePoint('e'), _dafny.CodePoint('b')))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_dafny.CodePoint('s'))))
			} else {
				var _474___mcc_h7 _dafny.Map = _source10.Get_().(D2_DC6).Cf13
				_ = _474___mcc_h7
				var _475_cf13 _dafny.Map = _474___mcc_h7
				_ = _475_cf13
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.MultiSetOf(_dafny.CodePoint('x'), _dafny.CodePoint('b'), _dafny.CodePoint('r')))
			}
		}(Companion_D2_.Create_DC8_())).Cardinality()
	}
}
func (_this *C2) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetOf(_dafny.SeqOf((Companion_D7_.Create_DC19_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()), _dafny.MultiSetOf((_this).F27()), (_this).F27(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-748))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg44 _dafny.Int) interface{} {
				return coer44(arg44)
			}
		}(func(_476_i0 _dafny.Int) _dafny.Set {
			return _dafny.SetOf((_this).F26())
		}))).Cardinality())))).Dtor_cf36(), _dafny.IntOfInt64(-153)))).Cardinality()
	}
}
func (_this *C2) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C2) Fm16(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source11 D6 = Companion_D6_.Create_DC17_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uot")).Cardinality()), _dafny.IntOfInt64(884), (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F26()))).Cardinality(), _dafny.IntOfInt64(-382), (_dafny.SetOf((_this).F26(), (_this).F27(), true)).Cardinality())
		_ = _source11
		if _source11.Is_DC17() {
			var _477___mcc_h0 _dafny.Int = _source11.Get_().(D6_DC17).Cf30
			_ = _477___mcc_h0
			var _478___mcc_h1 _dafny.Int = _source11.Get_().(D6_DC17).Cf31
			_ = _478___mcc_h1
			var _479___mcc_h2 _dafny.Int = _source11.Get_().(D6_DC17).Cf32
			_ = _479___mcc_h2
			var _480___mcc_h3 _dafny.Int = _source11.Get_().(D6_DC17).Cf33
			_ = _480___mcc_h3
			var _481___mcc_h4 _dafny.Int = _source11.Get_().(D6_DC17).Cf34
			_ = _481___mcc_h4
			var _482_cf34 _dafny.Int = _481___mcc_h4
			_ = _482_cf34
			var _483_cf33 _dafny.Int = _480___mcc_h3
			_ = _483_cf33
			var _484_cf32 _dafny.Int = _479___mcc_h2
			_ = _484_cf32
			var _485_cf31 _dafny.Int = _478___mcc_h1
			_ = _485_cf31
			var _486_cf30 _dafny.Int = _477___mcc_h0
			_ = _486_cf30
			return (_this).F26()
		} else {
			var _487___mcc_h5 _dafny.Sequence = _source11.Get_().(D6_DC16).Cf29
			_ = _487___mcc_h5
			var _488_cf29 _dafny.Sequence = _487___mcc_h5
			_ = _488_cf29
			return !((_this).F26())
		}
	}
}
func (_this *C2) Fm17(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return (_this).F26()
	}
}
func (_this *C2) M2(p0 bool, p1 T0, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		r1 = (_this).Fm1(globalState)
		var _489_v0 _dafny.Set
		_ = _489_v0
		_489_v0 = _dafny.SetOf(p0)
		var _490_v1 _dafny.Int
		_ = _490_v1
		_490_v1 = _dafny.IntOfInt64(-337)
		var _491_v2 _dafny.Sequence
		_ = _491_v2
		_491_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ajxqysp")
		var _rhs52 _dafny.Int = (_dafny.Zero).Minus(((_490_v1).Times(_dafny.IntOfUint32((_491_v2).Cardinality()))).Minus(_490_v1))
		_ = _rhs52
		var _rhs53 _dafny.Set = (Companion_Default___.Fm18(_489_v0, globalState)).Difference(_489_v0)
		_ = _rhs53
		var _rhs54 _dafny.Int = _490_v1
		_ = _rhs54
		var _lhs53 *GlobalState = globalState
		_ = _lhs53
		_lhs53.F16 = _rhs52
		_489_v0 = _rhs53
		r1 = _rhs54
		var _492_v3 _dafny.Sequence
		_ = _492_v3
		_492_v3 = _dafny.SeqOf(p0)
		var _493_v4 _dafny.Sequence
		_ = _493_v4
		_493_v4 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_492_v3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_492_v3).Cardinality()), _dafny.IntOfUint32((_492_v3).Cardinality()))).Uint32(), p2), _dafny.Companion_Sequence_.Concatenate(_492_v3, _492_v3))
		_493_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_493_v4, _493_v4), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg45 _dafny.Int) interface{} {
				return coer45(arg45)
			}
		}((func(_494_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_495_i0 _dafny.Int) _dafny.Sequence {
				return _494_v3
			}
		})(_492_v3))), (Companion_Default___.SafeIndex(_490_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}((func(_496_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_497_i0 _dafny.Int) _dafny.Sequence {
				return _496_v3
			}
		})(_492_v3)))).Cardinality()))).Uint32(), _492_v3))
		var _498_v5 _dafny.Array
		_ = _498_v5
		var _nw56 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
		_ = _nw56
		_498_v5 = _nw56
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_498_v5), 0))
		_ = _index61
		(_498_v5).ArraySet1(true, (_index61).Int())
		var _499_v6 _dafny.Map
		_ = _499_v6
		_499_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_490_v1, _dafny.IntOfUint32((_491_v2).Cardinality()), _490_v1, _490_v1), (_this).F26())
		var _500_v7 _dafny.MultiSet
		_ = _500_v7
		_500_v7 = _dafny.MultiSetOf(_490_v1, _dafny.IntOfUint32((_491_v2).Cardinality()))
		var _501_v8 _dafny.Map
		_ = _501_v8
		_501_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_491_v2).Cardinality()))
		var _502_v9 _dafny.Map
		_ = _502_v9
		_502_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_501_v8).Cardinality(), !(false))
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_498_v5), 0))
		_ = _index62
		(_498_v5).ArraySet1((func() bool {
			if (_499_v6).Contains((_500_v7).Union(_dafny.MultiSetOf((_502_v9).Cardinality()))) {
				return (_499_v6).Get((_500_v7).Union(_dafny.MultiSetOf((_502_v9).Cardinality()))).(bool)
			}
			return (_this).F26()
		})(), (_index62).Int())
		var _503_v10 _dafny.CodePoint
		_ = _503_v10
		_503_v10 = _dafny.CodePoint('p')
		var _504_v11 _dafny.Map
		_ = _504_v11
		_504_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_491_v2).Select((Companion_Default___.SafeIndex(_490_v1, _dafny.IntOfUint32((_491_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), _503_v10)
		var _505_v12 _dafny.Sequence
		_ = _505_v12
		_505_v12 = _dafny.SeqOf(_504_v11)
		_505_v12 = _505_v12
		var _506_v13 _dafny.Sequence
		_ = _506_v13
		_506_v13 = _dafny.SeqOf(_dafny.IntOfInt64(-288), (_500_v7).Cardinality(), _dafny.IntOfInt64(180), _490_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flhymnjo")).Cardinality()))
		var _507_v15 _dafny.Array
		_ = _507_v15
		var _nwElement0_6 _dafny.Sequence = _506_v13
		_ = _nwElement0_6
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(18))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_6, 0)
		(_nw57).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(220))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg47 _dafny.Int) interface{} {
				return coer47(arg47)
			}
		}((func(_508_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_509_i1 _dafny.Int) _dafny.Int {
				return _508_v1
			}
		})(_490_v1))), 1)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_506_v13, _506_v13), 2)
		(_nw57).ArraySet1(_506_v13, 3)
		(_nw57).ArraySet1(_506_v13, 4)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if !((_498_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_498_v5), 0))).Int()).(bool)) {
				return _506_v13
			}
			return _506_v13
		})(), (Companion_Default___.SafeIndex(_490_v1, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if !((_498_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(800), _dafny.ArrayLen((_498_v5), 0))).Int()).(bool)) {
				return _506_v13
			}
			return _506_v13
		})()).Cardinality()))).Uint32(), _490_v1), 5)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_506_v13, _506_v13), 6)
		(_nw57).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}((func(_510_v7 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
			return func(_511_i2 _dafny.Int) _dafny.Int {
				return (_510_v7).Cardinality()
			}
		})(_500_v7))), 7)
		(_nw57).ArraySet1(_506_v13, 8)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_dafny.IntOfUint32((_491_v2).Cardinality())), (Companion_Default___.SafeIndex((func() _dafny.Set {
			var _coll49 = _dafny.NewBuilder()
			_ = _coll49
			for _iter49 := _dafny.Iterate((_504_v11).Keys().Elements()); ; {
				_compr_49, _ok49 := _iter49()
				if !_ok49 {
					break
				}
				var _512_v14 _dafny.CodePoint
				_512_v14 = interface{}(_compr_49).(_dafny.CodePoint)
				if (_504_v11).Contains(_512_v14) {
					_coll49.Add(_512_v14)
				}
			}
			return _coll49.ToSet()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_491_v2).Cardinality()))).Cardinality()))).Uint32(), _490_v1), 9)
		(_nw57).ArraySet1(_506_v13, 10)
		(_nw57).ArraySet1(_506_v13, 11)
		(_nw57).ArraySet1(_506_v13, 12)
		(_nw57).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(444))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}((func(_513_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_514_i3 _dafny.Int) _dafny.Int {
				return _513_v1
			}
		})(_490_v1))), 13)
		(_nw57).ArraySet1(_506_v13, 14)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Update(_506_v13, (Companion_Default___.SafeIndex(_490_v1, _dafny.IntOfUint32((_506_v13).Cardinality()))).Uint32(), _490_v1), 15)
		(_nw57).ArraySet1(_506_v13, 16)
		(_nw57).ArraySet1(_506_v13, 17)
		_507_v15 = _nw57
		var _515_v16 _dafny.Set
		_ = _515_v16
		_515_v16 = _dafny.SetOf((_500_v7).Cardinality())
		_507_v15 = (func() _dafny.Array {
			if (_515_v16).Contains((_dafny.Zero).Minus(_490_v1)) {
				return _507_v15
			}
			return _507_v15
		})()
		r0 = _503_v10
		r1 = ((_dafny.Zero).Minus(_490_v1)).Times((_dafny.Zero).Minus(_490_v1))
		return r0, r1
	}
}
func (_this *C2) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _516_v0 _dafny.Array
		_ = _516_v0
		var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw58
		_516_v0 = _nw58
		var _517_v1 _dafny.MultiSet
		_ = _517_v1
		_517_v1 = _dafny.MultiSetOf(_516_v0)
		var _518_v2 _dafny.Int
		_ = _518_v2
		_518_v2 = _dafny.IntOfInt64(240)
		var _519_v3 _dafny.Map
		_ = _519_v3
		_519_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v2, _516_v0)
		var _520_v4 D0
		_ = _520_v4
		_520_v4 = Companion_D0_.Create_DC0_(p0)
		var _521_v5 _dafny.CodePoint
		_ = _521_v5
		_521_v5 = _dafny.CodePoint('a')
		var _522_v6 _dafny.MultiSet
		_ = _522_v6
		_522_v6 = _dafny.MultiSetOf(_521_v5, _521_v5, _521_v5, _521_v5)
		var _523_v7 _dafny.Sequence
		_ = _523_v7
		_523_v7 = _dafny.UnicodeSeqOfUtf8Bytes("cudjrcpt")
		var _524_v8 _dafny.Map
		_ = _524_v8
		_524_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_523_v7).Cardinality()), _dafny.MultiSetOf(_516_v0, _516_v0, _516_v0, _516_v0, _516_v0))
		var _525_v9 _dafny.Array
		_ = _525_v9
		var _nwElement0_7 _dafny.MultiSet = _517_v1
		_ = _nwElement0_7
		var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(28))
		_ = _nw59
		(_nw59).ArraySet1(_nwElement0_7, 0)
		(_nw59).ArraySet1(_dafny.MultiSetOf((func() _dafny.Array {
			if (_519_v3).Contains((_this).Fm4(_520_v4, _518_v2, (_522_v6).Cardinality(), globalState)) {
				return (_519_v3).Get((_this).Fm4(_520_v4, _518_v2, (_522_v6).Cardinality(), globalState)).(_dafny.Array)
			}
			return _516_v0
		})()), 1)
		(_nw59).ArraySet1(_517_v1, 2)
		(_nw59).ArraySet1(_517_v1, 3)
		(_nw59).ArraySet1(_517_v1, 4)
		(_nw59).ArraySet1(_517_v1, 5)
		(_nw59).ArraySet1(_517_v1, 6)
		(_nw59).ArraySet1(_dafny.MultiSetOf(_516_v0, _516_v0), 7)
		(_nw59).ArraySet1(_517_v1, 8)
		(_nw59).ArraySet1(_517_v1, 9)
		(_nw59).ArraySet1(_517_v1, 10)
		(_nw59).ArraySet1(((_517_v1).Update(_516_v0, Companion_Default___.Abs(_dafny.IntOfInt64(892)))).Difference(_517_v1), 11)
		(_nw59).ArraySet1((_517_v1).Intersection(_517_v1), 12)
		(_nw59).ArraySet1(_517_v1, 13)
		(_nw59).ArraySet1(_517_v1, 14)
		(_nw59).ArraySet1(_517_v1, 15)
		(_nw59).ArraySet1(_517_v1, 16)
		(_nw59).ArraySet1((_517_v1).Union(_517_v1), 17)
		(_nw59).ArraySet1(_517_v1, 18)
		(_nw59).ArraySet1(_517_v1, 19)
		(_nw59).ArraySet1((_517_v1).Union(_dafny.MultiSetOf(_516_v0, _516_v0, _516_v0, _516_v0, _516_v0)), 20)
		(_nw59).ArraySet1(_517_v1, 21)
		(_nw59).ArraySet1((_517_v1).Difference(_517_v1), 22)
		(_nw59).ArraySet1(_517_v1, 23)
		(_nw59).ArraySet1(_517_v1, 24)
		(_nw59).ArraySet1(_517_v1, 25)
		(_nw59).ArraySet1((func() _dafny.MultiSet {
			if (_524_v8).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_518_v2))) {
				return (_524_v8).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(_518_v2))).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_516_v0, _516_v0, _516_v0)
		})(), 26)
		(_nw59).ArraySet1(_517_v1, 27)
		_525_v9 = _nw59
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_525_v9), 0))
		_ = _index63
		(_525_v9).ArraySet1(_517_v1, (_index63).Int())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_525_v9), 0))
		_ = _index64
		(_525_v9).ArraySet1(_517_v1, (_index64).Int())
		var _526_i0 _dafny.Int
		_ = _526_i0
		_526_i0 = _dafny.Zero
		{
			for !((_522_v6).IsProperSubsetOf((_dafny.MultiSetOf(_521_v5)).Union(_dafny.MultiSetOf(_521_v5)))) {
				{
					if (_526_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_526_i0 = (_526_i0).Plus(_dafny.One)
					var _527_v10 _dafny.Map
					_ = _527_v10
					_527_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _518_v2)
					_527_v10 = (_527_v10).Update((_this).F27(), _518_v2)
					var _528_v11 *C0
					_ = _528_v11
					var _nw60 *C0 = New_C0_()
					_ = _nw60
					_nw60.Ctor__(p0)
					_528_v11 = _nw60
					var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_516_v0), 0))
					_ = _index65
					(_516_v0).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).Fm1(globalState)), _518_v2)), (_index65).Int())
					var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_516_v0), 0))
					_ = _index66
					(_516_v0).ArraySet1(_518_v2, (_index66).Int())
					var _529_v12 _dafny.Array
					_ = _529_v12
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_13
					var _nw61 _dafny.Array
					_ = _nw61
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw61 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = (func(_530_v0 _dafny.Array) func(_dafny.Int) _dafny.Int {
							return func(_531_i1 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_531_i1, (_530_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_530_v0), 0))).Int()).(_dafny.Int))
							}
						})(_516_v0)
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw61 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw61).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw61).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_529_v12 = _nw61
					var _532_v13 D3
					_ = _532_v13
					_532_v13 = Companion_D3_.Create_DC10_(_529_v12)
					_532_v13 = Companion_D3_.Create_DC10_(_529_v12)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _533_v14 _dafny.Set
		_ = _533_v14
		_533_v14 = _dafny.SetOf(_518_v2)
		var _534_v15 _dafny.MultiSet
		_ = _534_v15
		_534_v15 = _dafny.MultiSetOf(_533_v14)
		var _535_v16 _dafny.MultiSet
		_ = _535_v16
		_535_v16 = _dafny.MultiSetOf((_this).Fm17(_dafny.IntOfUint32((_523_v7).Cardinality()), (_dafny.Zero).Minus(_518_v2), _518_v2, _534_v15, globalState), (_this).F27(), (_this).F26(), !(!(false)))
		(globalState).F12 = ((_dafny.IntOfUint32((_523_v7).Cardinality())).Times((func() _dafny.Int {
			if (_535_v16).Contains(!(p0)) {
				return (_535_v16).Multiplicity(!(p0))
			}
			return _518_v2
		})())).Minus((_this).Fm1(globalState))
		var _536_v17 _dafny.Map
		_ = _536_v17
		_536_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F26())
		(_this).M12(_518_v2, (_536_v17).Cardinality(), globalState)
		var _537_v18 T2
		_ = _537_v18
		var _nw62 *C1 = New_C1_()
		_ = _nw62
		_nw62.Ctor__(_518_v2, _dafny.IntOfUint32((_523_v7).Cardinality()), p0, (_this).F26())
		_537_v18 = _nw62
		var _538_v19 D7
		_ = _538_v19
		_538_v19 = Companion_D7_.Create_DC18_(_537_v18)
		var _source12 D7 = _538_v19
		_ = _source12
		if _source12.Is_DC19() {
			var _539___mcc_h0 _dafny.Int = _source12.Get_().(D7_DC19).Cf36
			_ = _539___mcc_h0
			var _540___mcc_h1 _dafny.MultiSet = _source12.Get_().(D7_DC19).Cf37
			_ = _540___mcc_h1
			var _541___mcc_h2 bool = _source12.Get_().(D7_DC19).Cf38
			_ = _541___mcc_h2
			var _542___mcc_h3 _dafny.Map = _source12.Get_().(D7_DC19).Cf39
			_ = _542___mcc_h3
			var _543_cf39 _dafny.Map = _542___mcc_h3
			_ = _543_cf39
			var _544_cf38 bool = _541___mcc_h2
			_ = _544_cf38
			var _545_cf37 _dafny.MultiSet = _540___mcc_h1
			_ = _545_cf37
			var _546_cf36 _dafny.Int = _539___mcc_h0
			_ = _546_cf36
			(globalState).F15 = (_537_v18).F26()
			(globalState).F12 = (_518_v2).Plus((_518_v2).Minus(_dafny.IntOfInt64(648)))
			var _547_v20 _dafny.Map
			_ = _547_v20
			_547_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()), _546_cf36)
			var _548_v21 _dafny.Array
			_ = _548_v21
			var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw63
			_548_v21 = _nw63
			var _549_v22 _dafny.Map
			_ = _549_v22
			_549_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_548_v21, (_537_v18).F27())
			var _550_v23 _dafny.Map
			_ = _550_v23
			_550_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_549_v22).Contains(_548_v21) {
					return (_549_v22).Get(_548_v21).(bool)
				}
				return (_537_v18).F27()
			})(), _536_v17)
			_546_cf36 = (_537_v18).Fm4(_520_v4, ((_547_v20).Cardinality()).Times(_518_v2), (_550_v23).Cardinality(), globalState)
			_521_v5 = _521_v5
		} else {
			var _551___mcc_h4 T2 = _source12.Get_().(D7_DC18).Cf35
			_ = _551___mcc_h4
			var _552_cf35 T2 = _551___mcc_h4
			_ = _552_cf35
			var _553_v24 _dafny.Array
			_ = _553_v24
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_14
			var _nw64 _dafny.Array
			_ = _nw64
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw64 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = (func(_554_v2 _dafny.Int, _555_v7 _dafny.Sequence, _556_v18 T2) func(_dafny.Int) bool {
					return func(_557_i2 _dafny.Int) bool {
						return (Companion_D5_.Create_DC14_(_554_v2, _dafny.IntOfUint32((_555_v7).Cardinality()), (_556_v18).F27())).Dtor_cf27()
					}
				})(_518_v2, _523_v7, _537_v18)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw64 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw64).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw64).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_553_v24 = _nw64
			var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("htwwcultw"), _dafny.UnicodeSeqOfUtf8Bytes("jlkne"))
			_ = _rhs55
			var _rhs56 _dafny.Array = _553_v24
			_ = _rhs56
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			_lhs54.F6 = _rhs55
			_553_v24 = _rhs56
			(globalState).F12 = (_518_v2).Times(_dafny.IntOfInt64(970))
			var _558_v25 _dafny.Array
			_ = _558_v25
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_15
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Set = (func(_559_v14 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_560_i3 _dafny.Int) _dafny.Set {
						return _559_v14
					}
				})(_533_v14)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw65 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw65).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw65).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_558_v25 = _nw65
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_558_v25), 0))
			_ = _index67
			(_558_v25).ArraySet1(func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter50 := _dafny.Iterate((_dafny.MultiSetFromSeq(Companion_Default___.Fm23(_521_v5, globalState))).Elements()); ; {
					_compr_50, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _561_v26 _dafny.Int
					_561_v26 = interface{}(_compr_50).(_dafny.Int)
					if (_dafny.MultiSetFromSeq(Companion_Default___.Fm23(_521_v5, globalState))).Contains(_561_v26) {
						_coll50.Add(Companion_Default___.SafeDivisionInt(_561_v26, (_dafny.SetOf(false)).Cardinality()))
					}
				}
				return _coll50.ToSet()
			}(), (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_558_v25), 0))
			_ = _index68
			(_558_v25).ArraySet1(_533_v14, (_index68).Int())
			(globalState).F16 = (_dafny.Zero).Minus(_518_v2)
		}
		_521_v5 = (func() _dafny.CodePoint {
			if (_537_v18).Fm2(p0, _518_v2, globalState) {
				return _521_v5
			}
			return (func() _dafny.CodePoint {
				if p0 {
					return _521_v5
				}
				return _521_v5
			})()
		})()
	}
}
func (_this *C2) M11(p0 _dafny.Sequence, p1 _dafny.Array, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _562_v0 D0
		_ = _562_v0
		_562_v0 = Companion_D0_.Create_DC0_((_this).F26())
		var _563_v1 _dafny.Sequence
		_ = _563_v1
		_563_v1 = _dafny.UnicodeSeqOfUtf8Bytes("euqbmr")
		var _564_v2 _dafny.Map
		_ = _564_v2
		_564_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v1, (_this).F27())
		var _565_v3 _dafny.Int
		_ = _565_v3
		_565_v3 = _dafny.IntOfInt64(671)
		var _566_v4 _dafny.Array
		_ = _566_v4
		var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw66
		_566_v4 = _nw66
		var _567_v5 _dafny.Map
		_ = _567_v5
		_567_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(func(_pat_let4_0 D0) D0 {
			return func(_568_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let5_0 bool) D0 {
					return func(_569_dt__update_hcf0_h0 bool) D0 {
						return Companion_D0_.Create_DC0_(_569_dt__update_hcf0_h0)
					}(_pat_let5_0)
				}((_this).F26())
			}(_pat_let4_0)
		}(_562_v0), (_564_v2).Cardinality(), _565_v3, globalState), _566_v4)
		_567_v5 = (_567_v5).Update(_565_v3, _566_v4)
		var _570_v6 _dafny.Map
		_ = _570_v6
		_570_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_v1, _565_v3)
		var _hi5 _dafny.Int = _565_v3
		_ = _hi5
		for _571_i0 := (_570_v6).Cardinality(); _571_i0.Cmp(_hi5) < 0; _571_i0 = _571_i0.Plus(_dafny.One) {
			var _572_v7 _dafny.Map
			_ = _572_v7
			_572_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm24(_563_v1, (_this).F26(), _565_v3, globalState), (_this).F27())
			_572_v7 = (_572_v7).Merge(_572_v7)
			var _573_v8 _dafny.CodePoint
			_ = _573_v8
			_573_v8 = _dafny.CodePoint('n')
			var _574_v9 _dafny.CodePoint
			_ = _574_v9
			var _575_v10 _dafny.Sequence
			_ = _575_v10
			var _576_v11 bool
			_ = _576_v11
			var _577_v12 _dafny.Map
			_ = _577_v12
			var _out21 _dafny.CodePoint
			_ = _out21
			var _out22 _dafny.Sequence
			_ = _out22
			var _out23 bool
			_ = _out23
			var _out24 _dafny.Map
			_ = _out24
			_out21, _out22, _out23, _out24 = Companion_Default___.M0(_573_v8, globalState)
			_574_v9 = _out21
			_575_v10 = _out22
			_576_v11 = _out23
			_577_v12 = _out24
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_566_v4), 0))
			_ = _index69
			(_566_v4).ArraySet1(_565_v3, (_index69).Int())
			var _578_v13 D6
			_ = _578_v13
			_578_v13 = Companion_D6_.Create_DC16_(_dafny.Companion_Sequence_.Concatenate(_563_v1, _563_v1))
			var _579_v14 D2
			_ = _579_v14
			_579_v14 = Companion_D2_.Create_DC9_(_571_i0, _571_i0, (_this).F27())
			var _580_v15 _dafny.MultiSet
			_ = _580_v15
			_580_v15 = _dafny.MultiSetOf(_dafny.SetOf(_565_v3, _dafny.IntOfInt64(201)))
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_566_v4), 0))
			_ = _index70
			var _rhs57 _dafny.Int = _571_i0
			_ = _rhs57
			var _rhs58 _dafny.Int = _571_i0
			_ = _rhs58
			var _rhs59 D6 = (func() D6 {
				if ((_579_v14).Dtor_cf19()).Cmp(_565_v3) != 0 {
					return _578_v13
				}
				return Companion_D6_.Create_DC16_(_563_v1)
			})()
			_ = _rhs59
			var _rhs60 bool = (_this).Fm17(_571_i0, (_dafny.Zero).Minus(_571_i0), _565_v3, _580_v15, globalState)
			_ = _rhs60
			var _lhs55 *GlobalState = globalState
			_ = _lhs55
			var _lhs56 _dafny.Array = _566_v4
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_566_v4), 0))
			_ = _lhs57
			var _lhs58 *GlobalState = globalState
			_ = _lhs58
			_lhs55.F12 = _rhs57
			(_lhs56).ArraySet1(_rhs58, (_lhs57).Int())
			_578_v13 = _rhs59
			_lhs58.F2 = _rhs60
			var _581_v16 _dafny.Array
			_ = _581_v16
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_16
			var _nw67 _dafny.Array
			_ = _nw67
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw67 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) bool = func(_582_i1 _dafny.Int) bool {
					return (_this).F27()
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw67 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw67).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw67).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_581_v16 = _nw67
			var _583_v17 _dafny.Map
			_ = _583_v17
			_583_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F27()), (_this).F27())
			_581_v16 = (func() _dafny.Array {
				if (func() bool {
					if (_583_v17).Contains(_576_v11) {
						return (_583_v17).Get(_576_v11).(bool)
					}
					return true
				})() {
					return _581_v16
				}
				return p1
			})()
		}
		var _584_i2 _dafny.Int
		_ = _584_i2
		_584_i2 = _dafny.Zero
		{
			for (false) == ((_this).F26()) {
				{
					if (_584_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_584_i2 = (_584_i2).Plus(_dafny.One)
					var _585_v18 _dafny.CodePoint
					_ = _585_v18
					_585_v18 = _dafny.CodePoint('g')
					var _586_v19 _dafny.Set
					_ = _586_v19
					_586_v19 = _dafny.SetOf((_this).F27())
					var _587_v20 _dafny.Map
					_ = _587_v20
					_587_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_585_v18, _585_v18), _563_v1)).Cardinality()), (_586_v19).Intersection(_586_v19))
					_587_v20 = _587_v20
					(globalState).F22 = (_this).F26()
					(globalState).F12 = _565_v3
					(globalState).F16 = _dafny.IntOfInt64(308)
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _588_v21 _dafny.MultiSet
		_ = _588_v21
		_588_v21 = _dafny.MultiSetOf(_565_v3)
		var _589_v22 _dafny.Set
		_ = _589_v22
		_589_v22 = _dafny.SetOf((_this).F27(), false)
		var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))
		_ = _index71
		(p1).ArraySet1((_588_v21).IsProperSubsetOf(_dafny.MultiSetOf((_589_v22).Cardinality(), _565_v3)), (_index71).Int())
		var _590_v23 _dafny.Map
		_ = _590_v23
		_590_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), false)
		var _591_v24 _dafny.MultiSet
		_ = _591_v24
		_591_v24 = _dafny.MultiSetOf((func() bool {
			if (_590_v23).Contains((_this).F27()) {
				return (_590_v23).Get((_this).F27()).(bool)
			}
			return false
		})())
		var _592_v25 _dafny.Array
		_ = _592_v25
		var _len0_17 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_17
		var _nw68 _dafny.Array
		_ = _nw68
		if _len0_17.Cmp(_dafny.Zero) == 0 {
			_nw68 = _dafny.NewArray(_len0_17)
		} else {
			var _init17 func(_dafny.Int) _dafny.Set = func(_593_i3 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F27())).Cardinality()))
			}
			_ = _init17
			var _element0_17 = _init17(_dafny.Zero)
			_ = _element0_17
			_nw68 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
			(_nw68).ArraySet1(_element0_17, 0)
			var _nativeLen0_17 = (_len0_17).Int()
			_ = _nativeLen0_17
			for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
				(_nw68).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
			}
		}
		_592_v25 = _nw68
		var _594_v26 _dafny.Map
		_ = _594_v26
		_594_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_565_v3, _591_v24)
		var _595_v27 _dafny.Map
		_ = _595_v27
		_595_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_589_v22, _592_v25)
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))
		_ = _index72
		var _rhs61 bool = (_this).F26()
		_ = _rhs61
		var _rhs62 _dafny.MultiSet = (func() _dafny.MultiSet {
			if (_594_v26).Contains(_565_v3) {
				return (_594_v26).Get(_565_v3).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(true, (_this).F27())
		})()
		_ = _rhs62
		var _rhs63 bool = ((_this).F27()) && ((_this).F26())
		_ = _rhs63
		var _rhs64 _dafny.Array = (func() _dafny.Array {
			if true {
				return _592_v25
			}
			return (func() _dafny.Array {
				if (_595_v27).Contains(_dafny.SetOf((_this).F27())) {
					return (_595_v27).Get(_dafny.SetOf((_this).F27())).(_dafny.Array)
				}
				return _592_v25
			})()
		})()
		_ = _rhs64
		var _rhs65 _dafny.Int = (_dafny.Zero).Minus(_565_v3)
		_ = _rhs65
		var _lhs59 _dafny.Array = p1
		_ = _lhs59
		var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))
		_ = _lhs60
		var _lhs61 *GlobalState = globalState
		_ = _lhs61
		(_lhs59).ArraySet1(_rhs61, (_lhs60).Int())
		_591_v24 = _rhs62
		_lhs61.F15 = _rhs63
		_592_v25 = _rhs64
		_565_v3 = _rhs65
		var _596_v28 _dafny.Set
		_ = _596_v28
		_596_v28 = _dafny.SetOf(_565_v3)
		var _597_i4 _dafny.Int
		_ = _597_i4
		_597_i4 = _dafny.Zero
		{
			for !(_596_v28).Contains((_565_v3).Plus(_565_v3)) {
				{
					if (_597_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_597_i4 = (_597_i4).Plus(_dafny.One)
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_566_v4), 0))
					_ = _index73
					(_566_v4).ArraySet1(_565_v3, (_index73).Int())
					var _598_v29 _dafny.Set
					_ = _598_v29
					_598_v29 = _589_v22
					var _599_v30 _dafny.CodePoint
					_ = _599_v30
					_599_v30 = _dafny.CodePoint('c')
					var _600_v31 _dafny.Sequence
					_ = _600_v31
					_600_v31 = _dafny.SeqOf((_this).F26(), false, true)
					var _601_v32 _dafny.Map
					_ = _601_v32
					_601_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _600_v31)
					var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_566_v4), 0))
					_ = _index74
					var _rhs66 _dafny.Int = ((_this).Fm4(_562_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool), _599_v30)).Cardinality(), (_601_v32).Cardinality(), globalState)).Minus(_565_v3)
					_ = _rhs66
					var _rhs67 _dafny.Int = _565_v3
					_ = _rhs67
					var _rhs68 _dafny.Set = _598_v29
					_ = _rhs68
					var _lhs62 *GlobalState = globalState
					_ = _lhs62
					var _lhs63 _dafny.Array = _566_v4
					_ = _lhs63
					var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_566_v4), 0))
					_ = _lhs64
					_lhs62.F16 = _rhs66
					(_lhs63).ArraySet1(_rhs67, (_lhs64).Int())
					_598_v29 = _rhs68
					var _602_v33 *C1
					_ = _602_v33
					var _nw69 *C1 = New_C1_()
					_ = _nw69
					_nw69.Ctor__(_565_v3, _565_v3, (_this).F27(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))
					_602_v33 = _nw69
					var _603_v34 _dafny.Sequence
					_ = _603_v34
					_603_v34 = _dafny.SeqOf(_602_v33, _602_v33, _602_v33)
					var _rhs69 bool = _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm22(_565_v3, (_this).F26(), globalState), _dafny.Companion_Sequence_.Concatenate(_563_v1, _563_v1))
					_ = _rhs69
					var _rhs70 *C1 = (_603_v34).Select((Companion_Default___.SafeIndex(_565_v3, _dafny.IntOfUint32((_603_v34).Cardinality()))).Uint32()).(*C1)
					_ = _rhs70
					var _rhs71 bool = _dafny.Companion_Sequence_.Contains(_563_v1, _599_v30)
					_ = _rhs71
					var _rhs72 _dafny.CodePoint = _599_v30
					_ = _rhs72
					var _lhs65 *GlobalState = globalState
					_ = _lhs65
					var _lhs66 *GlobalState = globalState
					_ = _lhs66
					_lhs65.F22 = _rhs69
					_602_v33 = _rhs70
					_lhs66.F22 = _rhs71
					_599_v30 = _rhs72
					var _604_v35 *C0
					_ = _604_v35
					var _nw70 *C0 = New_C0_()
					_ = _nw70
					_nw70.Ctor__((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))
					_604_v35 = _nw70
					var _605_v36 *C1
					_ = _605_v36
					var _nw71 *C1 = New_C1_()
					_ = _nw71
					_nw71.Ctor__((_602_v33).F36(), _565_v3, (_604_v35).F32(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))
					_605_v36 = _nw71
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _hi6 _dafny.Int = (_this).Fm1(globalState)
		_ = _hi6
		for _606_i5 := _565_v3; _606_i5.Cmp(_hi6) < 0; _606_i5 = _606_i5.Plus(_dafny.One) {
			(globalState).F19 = (_563_v1).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_this).F27() {
					return _565_v3
				}
				return _606_i5
			})(), _dafny.IntOfUint32((_563_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			var _607_v37 _dafny.CodePoint
			_ = _607_v37
			_607_v37 = _dafny.CodePoint('n')
			if (_dafny.IntOfInt64(-287)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("lbthv"), (Companion_Default___.SafeIndex(_606_i5, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lbthv")).Cardinality()))).Uint32(), _607_v37), _563_v1)).Cardinality())) <= 0 {
				var _608_v38 _dafny.Map
				_ = _608_v38
				_608_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_i5, (_this).F27())
				var _609_v39 D2
				_ = _609_v39
				_609_v39 = Companion_D2_.Create_DC9_((_608_v38).Cardinality(), _565_v3, false)
				var _610_v40 _dafny.Sequence
				_ = _610_v40
				_610_v40 = _dafny.SeqOf(_609_v39, _609_v39)
				var _611_v41 *C1
				_ = _611_v41
				var _nw72 *C1 = New_C1_()
				_ = _nw72
				_nw72.Ctor__(_dafny.IntOfInt64(133), _565_v3, (_this).Fm16(_606_i5, _610_v40, (_dafny.SetOf(_606_i5)).Cardinality(), _dafny.IntOfUint32((_563_v1).Cardinality()), globalState), (_this).F26())
				_611_v41 = _nw72
				var _612_v42 _dafny.Map
				_ = _612_v42
				_612_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_611_v41).F36(), _dafny.SeqOf(true, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool)))
				var _613_v43 _dafny.Sequence
				_ = _613_v43
				_613_v43 = _dafny.SeqOf(true)
				var _614_v44 _dafny.Map
				_ = _614_v44
				_614_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_612_v42).Contains((_611_v41).F36()) {
						return (_612_v42).Get((_611_v41).F36()).(_dafny.Sequence)
					}
					return _613_v43
				})(), p1)
				var _rhs73 _dafny.Map = ((Companion_D9_.Create_DC21_(_614_v44)).Dtor_cf41()).Merge(_614_v44)
				_ = _rhs73
				var _rhs74 _dafny.Int = ((func() _dafny.Int {
					if (_this).F26() {
						return _606_i5
					}
					return _565_v3
				})()).Times(Companion_Default___.SafeDivisionInt((_611_v41).F36(), _606_i5))
				_ = _rhs74
				var _lhs67 *GlobalState = globalState
				_ = _lhs67
				_614_v44 = _rhs73
				_lhs67.F12 = _rhs74
				var _615_v45 *C1
				_ = _615_v45
				var _nw73 *C1 = New_C1_()
				_ = _nw73
				_nw73.Ctor__(_611_v41.F37, (_611_v41.F37).Minus(_606_i5), !((_this).F27()), (false) && ((_this).F26()))
				_615_v45 = _nw73
				var _616_v46 _dafny.Map
				_ = _616_v46
				_616_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_611_v41).F36())
				_590_v23 = (_590_v23).Update(((_616_v46).Cardinality()).Cmp(_611_v41.F37) != 0, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))
				var _617_v47 _dafny.Array
				_ = _617_v47
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(21))
				_ = _nw74
				_617_v47 = _nw74
				var _618_v48 _dafny.Map
				_ = _618_v48
				_618_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_607_v37, _617_v47)
				_618_v48 = (_618_v48).Update(_607_v37, _617_v47)
			} else {
				var _619_v49 _dafny.MultiSet
				_ = _619_v49
				_619_v49 = _dafny.MultiSetOf(_596_v28)
				var _620_v50 _dafny.Sequence
				_ = _620_v50
				_620_v50 = _dafny.SeqOf((_this).Fm17(_565_v3, _606_i5, _565_v3, _619_v49, globalState))
				(globalState).F16 = (_this).Fm4(_562_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_this).Fm3(_563_v1, _620_v50, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_565_v3, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))).Cardinality(), globalState), _620_v50)).Cardinality()), _dafny.IntOfInt64(864), globalState)
				(globalState).F16 = (_dafny.Zero).Minus((_this).Fm4(_562_v0, _565_v3, _dafny.IntOfInt64(400), globalState))
				(globalState).F12 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Int {
					if (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool) {
						return _565_v3
					}
					return _565_v3
				})()), ((_596_v28).Intersection(_596_v28)).Cardinality())
				(globalState).F12 = _606_i5
				(globalState).F24 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool)
			}
			var _621_v51 _dafny.Sequence
			_ = _621_v51
			_621_v51 = _dafny.SeqOf(!(false), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool), true)
			var _622_v52 _dafny.Map
			_ = _622_v52
			_622_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_621_v51, _dafny.SeqOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((p1), 0))).Int()).(bool))), (_this).F26())
			_622_v52 = _622_v52
			_621_v51 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F27()), _dafny.Companion_Sequence_.Concatenate(_621_v51, _dafny.SeqOf(!((_this).F26()))))
		}
		r0 = _dafny.UnicodeSeqOfUtf8Bytes("bmiwf")
		return r0
	}
}
func (_this *C2) M12(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _623_v0 _dafny.MultiSet
		_ = _623_v0
		_623_v0 = _dafny.MultiSetOf(p1)
		var _624_v1 *C1
		_ = _624_v1
		var _nw75 *C1 = New_C1_()
		_ = _nw75
		_nw75.Ctor__(p1, p1, (func() bool {
			if (_this).F27() {
				return (_this).F27()
			}
			return (_this).F27()
		})(), (_dafny.MultiSetOf(p1, p1)).IsDisjointFrom(_623_v0))
		_624_v1 = _nw75
		var _625_v2 _dafny.Array
		_ = _625_v2
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_18
		var _nw76 _dafny.Array
		_ = _nw76
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw76 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) bool = func(_626_i0 _dafny.Int) bool {
				return (_this).F26()
			}
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw76 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw76).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw76).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_625_v2 = _nw76
		var _627_v3 _dafny.Sequence
		_ = _627_v3
		_627_v3 = _dafny.SeqOf((_this).F26(), (_this).F26(), (_this).F27())
		var _628_v4 _dafny.Map
		_ = _628_v4
		_628_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _627_v3)
		var _629_v5 D5
		_ = _629_v5
		_629_v5 = Companion_D5_.Create_DC13_(_628_v4)
		var _rhs75 bool = !(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_this).F27(), (_this).F27(), (_this).F27(), (_this).F27()), _dafny.SeqOf((_this).F27(), (_this).F26()))) || ((_this).F26())
		_ = _rhs75
		var _rhs76 _dafny.Array = _625_v2
		_ = _rhs76
		var _rhs77 bool = (_this).F27()
		_ = _rhs77
		var _rhs78 D5 = _629_v5
		_ = _rhs78
		var _lhs68 *GlobalState = globalState
		_ = _lhs68
		var _lhs69 *GlobalState = globalState
		_ = _lhs69
		_lhs68.F24 = _rhs75
		_625_v2 = _rhs76
		_lhs69.F22 = _rhs77
		_629_v5 = _rhs78
		var _630_v6 _dafny.Sequence
		_ = _630_v6
		_630_v6 = _dafny.UnicodeSeqOfUtf8Bytes("adlk")
		var _hi7 _dafny.Int = Companion_Default___.SafeDivisionInt(_624_v1.F37, (_dafny.Zero).Minus(_dafny.IntOfUint32((_627_v3).Cardinality())))
		_ = _hi7
		for _631_i1 := _dafny.IntOfUint32((_630_v6).Cardinality()); _631_i1.Cmp(_hi7) < 0; _631_i1 = _631_i1.Plus(_dafny.One) {
			(globalState).F12 = _631_i1
			var _632_v7 _dafny.Array
			_ = _632_v7
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_19
			var _nw77 _dafny.Array
			_ = _nw77
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw77 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_633_v1 *C1) func(_dafny.Int) _dafny.Int {
					return func(_634_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_634_i2, (_633_v1).F36())
					}
				})(_624_v1)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw77 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw77).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw77).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_632_v7 = _nw77
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_632_v7), 0))
			_ = _index75
			(_632_v7).ArraySet1(_624_v1.F37, (_index75).Int())
			var _635_v8 _dafny.CodePoint
			_ = _635_v8
			_635_v8 = _dafny.CodePoint('s')
			var _636_v9 _dafny.Sequence
			_ = _636_v9
			_636_v9 = _dafny.SeqOf(_625_v2, _625_v2)
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_632_v7), 0))
			_ = _index76
			var _rhs79 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_630_v6, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_627_v3).Cardinality()), _dafny.IntOfUint32((_630_v6).Cardinality()))).Uint32(), _635_v8)
			_ = _rhs79
			var _rhs80 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_636_v9).Cardinality())))
			_ = _rhs80
			var _rhs81 bool = (_this).F26()
			_ = _rhs81
			var _rhs82 bool = (_this).F27()
			_ = _rhs82
			var _rhs83 bool = !((_624_v1).Fm2((_this).F26(), (func() _dafny.Int {
				if true {
					return _624_v1.F37
				}
				return _624_v1.F37
			})(), globalState))
			_ = _rhs83
			var _lhs70 *GlobalState = globalState
			_ = _lhs70
			var _lhs71 _dafny.Array = _632_v7
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_632_v7), 0))
			_ = _lhs72
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			_lhs70.F6 = _rhs79
			(_lhs71).ArraySet1(_rhs80, (_lhs72).Int())
			_lhs73.F22 = _rhs81
			_lhs74.F24 = _rhs82
			_lhs75.F2 = _rhs83
			var _637_v10 _dafny.Set
			_ = _637_v10
			_637_v10 = _dafny.SetOf(!((_this).F27()), (_this).F27(), !(false))
			_637_v10 = ((_637_v10).Intersection(_637_v10)).Intersection((_637_v10).Union(_637_v10))
			var _638_v11 _dafny.Array
			_ = _638_v11
			var _nwElement0_8 _dafny.Array = _625_v2
			_ = _nwElement0_8
			var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(27))
			_ = _nw78
			(_nw78).ArraySet1(_nwElement0_8, 0)
			(_nw78).ArraySet1(_625_v2, 1)
			(_nw78).ArraySet1(_625_v2, 2)
			(_nw78).ArraySet1(_625_v2, 3)
			(_nw78).ArraySet1(_625_v2, 4)
			(_nw78).ArraySet1(_625_v2, 5)
			(_nw78).ArraySet1(_625_v2, 6)
			(_nw78).ArraySet1(_625_v2, 7)
			(_nw78).ArraySet1(_625_v2, 8)
			(_nw78).ArraySet1(_625_v2, 9)
			(_nw78).ArraySet1(_625_v2, 10)
			(_nw78).ArraySet1(_625_v2, 11)
			(_nw78).ArraySet1(_625_v2, 12)
			(_nw78).ArraySet1(_625_v2, 13)
			(_nw78).ArraySet1(_625_v2, 14)
			(_nw78).ArraySet1(_625_v2, 15)
			(_nw78).ArraySet1(_625_v2, 16)
			(_nw78).ArraySet1(_625_v2, 17)
			(_nw78).ArraySet1(_625_v2, 18)
			(_nw78).ArraySet1(_625_v2, 19)
			(_nw78).ArraySet1(_625_v2, 20)
			(_nw78).ArraySet1(_625_v2, 21)
			(_nw78).ArraySet1(_625_v2, 22)
			(_nw78).ArraySet1(_625_v2, 23)
			(_nw78).ArraySet1(_625_v2, 24)
			(_nw78).ArraySet1(_625_v2, 25)
			(_nw78).ArraySet1(_625_v2, 26)
			_638_v11 = _nw78
			var _639_v12 _dafny.Sequence
			_ = _639_v12
			_639_v12 = _dafny.SeqOf(_638_v11, _638_v11, _638_v11)
			var _rhs84 bool = ((_this).F27()) == ((_this).F26())
			_ = _rhs84
			var _rhs85 _dafny.Array = (_639_v12).Select((Companion_Default___.SafeIndex((_624_v1).F36(), _dafny.IntOfUint32((_639_v12).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs85
			var _lhs76 *GlobalState = globalState
			_ = _lhs76
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			_lhs76.F22 = _rhs84
			_lhs77.F5 = _rhs85
		}
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_625_v2), 0))
		_ = _index77
		(_625_v2).ArraySet1((_this).F27(), (_index77).Int())
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(575), _dafny.ArrayLen((_625_v2), 0))
		_ = _index78
		(_625_v2).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(815))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg50 _dafny.Int) interface{} {
				return coer50(arg50)
			}
		}(func(_640_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})), _630_v6), _630_v6), (_index78).Int())
		(globalState).F16 = (_dafny.Zero).Minus(p0)
		(_624_v1).F37 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_630_v6, _630_v6)).Cardinality())
	}
}
func (_this *C2) F35() _dafny.Sequence {
	{
		return _this._f35
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f26 bool
	_f27 bool
	_f34 _dafny.MultiSet
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f26 = false
	_this._f27 = false
	_this._f34 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F26() bool {
	return _this._f26
}
func (_this *C3) F27() bool {
	return _this._f27
}
func (_this *C3) Ctor__(f34 _dafny.MultiSet, f26 bool, f27 bool) {
	{
		(_this)._f34 = f34
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C3) Fm1(globalState *GlobalState) _dafny.Int {
	{
		if (_this).F26() {
			return (_dafny.IntOfUint32(((Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("yoqe"))).Dtor_cf29()).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jmfpgf")).Cardinality()))
		} else {
			return _dafny.IntOfUint32((_dafny.SeqOf((_this).F27(), (_this).F26(), (_this).F26())).Cardinality())
		}
	}
}
func (_this *C3) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C3) Fm12(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((func() _dafny.Int {
			if (_this).F27() {
				return _dafny.IntOfInt64(84)
			}
			return _dafny.IntOfInt64(719)
		})()).Plus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()), _dafny.IntOfInt64(676)))
	}
}
func (_this *C3) Fm13(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(109))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_641_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32(((Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("r"))).Dtor_cf29()).Cardinality())
		}))).Cardinality())).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(956), (_this).F26()))).Cardinality()).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('k'))).Cardinality())))
	}
}
func (_this *C3) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _642_v0 _dafny.Array
		_ = _642_v0
		var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw79
		_642_v0 = _nw79
		var _643_v1 _dafny.Int
		_ = _643_v1
		_643_v1 = _dafny.IntOfInt64(281)
		var _644_v2 _dafny.Map
		_ = _644_v2
		_644_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_643_v1, _642_v0)
		var _645_v3 _dafny.Array
		_ = _645_v3
		var _nwElement0_9 _dafny.Array = _642_v0
		_ = _nwElement0_9
		var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(20))
		_ = _nw80
		(_nw80).ArraySet1(_nwElement0_9, 0)
		(_nw80).ArraySet1(_642_v0, 1)
		(_nw80).ArraySet1(_642_v0, 2)
		(_nw80).ArraySet1(_642_v0, 3)
		(_nw80).ArraySet1(_642_v0, 4)
		(_nw80).ArraySet1(_642_v0, 5)
		(_nw80).ArraySet1(_642_v0, 6)
		(_nw80).ArraySet1(_642_v0, 7)
		(_nw80).ArraySet1(_642_v0, 8)
		(_nw80).ArraySet1(_642_v0, 9)
		(_nw80).ArraySet1(_642_v0, 10)
		(_nw80).ArraySet1(_642_v0, 11)
		(_nw80).ArraySet1(_642_v0, 12)
		(_nw80).ArraySet1((func() _dafny.Array {
			if (_644_v2).Contains(_643_v1) {
				return (_644_v2).Get(_643_v1).(_dafny.Array)
			}
			return _642_v0
		})(), 13)
		(_nw80).ArraySet1(_642_v0, 14)
		(_nw80).ArraySet1(_642_v0, 15)
		(_nw80).ArraySet1(_642_v0, 16)
		(_nw80).ArraySet1(_642_v0, 17)
		(_nw80).ArraySet1(_642_v0, 18)
		(_nw80).ArraySet1(_642_v0, 19)
		_645_v3 = _nw80
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_645_v3), 0))
		_ = _index79
		(_645_v3).ArraySet1(_642_v0, (_index79).Int())
		var _646_v4 _dafny.Sequence
		_ = _646_v4
		_646_v4 = _dafny.SeqOf(_642_v0, _642_v0)
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(142), _dafny.ArrayLen((_645_v3), 0))
		_ = _index80
		(_645_v3).ArraySet1((_646_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.IntOfUint32((_646_v4).Cardinality()))).Uint32()).(_dafny.Array), (_index80).Int())
		var _647_v5 _dafny.Map
		_ = _647_v5
		_647_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F26())
		var _648_v6 _dafny.Sequence
		_ = _648_v6
		_648_v6 = _dafny.SeqOf(_647_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F27()), _647_v5, _647_v5, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F26()))
		var _649_v7 _dafny.Sequence
		_ = _649_v7
		_649_v7 = _dafny.SeqOf(_643_v1, _643_v1, _643_v1)
		var _650_v8 _dafny.Map
		_ = _650_v8
		_650_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_648_v6, _649_v7)
		var _651_v9 _dafny.Sequence
		_ = _651_v9
		_651_v9 = _dafny.UnicodeSeqOfUtf8Bytes("qgla")
		var _652_v10 _dafny.Map
		_ = _652_v10
		_652_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_651_v9, _648_v6)
		_650_v8 = (_650_v8).Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_652_v10).Contains(_651_v9) {
				return (_652_v10).Get(_651_v9).(_dafny.Sequence)
			}
			return _648_v6
		})(), _dafny.SeqOf(_647_v5)), _dafny.Companion_Sequence_.Concatenate(_649_v7, _649_v7))
		var _653_v11 _dafny.Sequence
		_ = _653_v11
		_653_v11 = _dafny.SeqOf((_this).F27(), (_this).F26(), p0, (_643_v1).Cmp(_dafny.IntOfUint32((_651_v9).Cardinality())) == 0, p0)
		_653_v11 = _dafny.Companion_Sequence_.Concatenate(_653_v11, _653_v11)
		var _654_v12 _dafny.Sequence
		_ = _654_v12
		_654_v12 = _dafny.SeqOf(_dafny.MultiSetOf(_dafny.CodePoint('y')))
		_643_v1 = (_dafny.Zero).Minus((_this).Fm13(_dafny.Companion_Sequence_.Concatenate(_654_v12, _654_v12), (_651_v9).Select((Companion_Default___.SafeIndex(_643_v1, _dafny.IntOfUint32((_651_v9).Cardinality()))).Uint32()).(_dafny.CodePoint), (_643_v1).Plus(_643_v1), globalState))
		_651_v9 = _dafny.Companion_Sequence_.Concatenate(_651_v9, _651_v9)
		var _655_v13 _dafny.Set
		_ = _655_v13
		_655_v13 = _dafny.SetOf((_this).F27())
		var _656_v14 _dafny.Map
		_ = _656_v14
		_656_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_655_v13).Cardinality()), _643_v1)
		var _hi8 _dafny.Int = _643_v1
		_ = _hi8
		for _657_i0 := (_dafny.Zero).Minus((func() _dafny.Int {
			if (_656_v14).Contains(_643_v1) {
				return (_656_v14).Get(_643_v1).(_dafny.Int)
			}
			return _643_v1
		})()); _657_i0.Cmp(_hi8) < 0; _657_i0 = _657_i0.Plus(_dafny.One) {
			(globalState).F15 = (_657_i0).Cmp(Companion_Default___.SafeDivisionInt(_657_i0, _dafny.IntOfUint32((_653_v11).Cardinality()))) < 0
			var _658_v15 _dafny.Map
			_ = _658_v15
			_658_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_657_i0, (_this).F27())
			var _659_v16 _dafny.MultiSet
			_ = _659_v16
			_659_v16 = _dafny.MultiSetOf(_643_v1)
			var _660_v17 _dafny.Map
			_ = _660_v17
			_660_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_658_v15).Cardinality()).Plus((_659_v16).Cardinality()), _651_v9)
			var _661_v18 _dafny.CodePoint
			_ = _661_v18
			_661_v18 = _dafny.CodePoint('e')
			_660_v17 = (_660_v17).Update(_dafny.IntOfUint32((_dafny.SeqOf(Companion_Default___.Fm14((_this).F26(), _657_i0, _643_v1, globalState), _661_v18, _661_v18, _661_v18)).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(934))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_662_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), _651_v9))
			_643_v1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if p0 {
					return _dafny.SeqOf(_657_i0)
				}
				return _dafny.SeqOf(_657_i0)
			})(), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm15(globalState), _649_v7))).Cardinality())
			var _663_v19 T2
			_ = _663_v19
			var _nw81 *C1 = New_C1_()
			_ = _nw81
			_nw81.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ylv")).Cardinality()), _dafny.IntOfUint32((_651_v9).Cardinality()), p0, (_this).F27())
			_663_v19 = _nw81
			var _664_v20 _dafny.Map
			_ = _664_v20
			_664_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D7_.Create_DC18_(_663_v19)).Dtor_cf35(), _657_i0)
			var _665_v21 _dafny.Array
			_ = _665_v21
			var _nwElement0_10 _dafny.Map = _664_v20
			_ = _nwElement0_10
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(12))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_10, 0)
			(_nw82).ArraySet1((_664_v20).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v19, _dafny.IntOfUint32((_646_v4).Cardinality()))), 1)
			(_nw82).ArraySet1(_664_v20, 2)
			(_nw82).ArraySet1((_664_v20).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v19, _643_v1)), 3)
			(_nw82).ArraySet1(_664_v20, 4)
			(_nw82).ArraySet1(_664_v20, 5)
			(_nw82).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v19, (_659_v16).Cardinality())).Merge(_664_v20), 6)
			(_nw82).ArraySet1(_664_v20, 7)
			(_nw82).ArraySet1(_664_v20, 8)
			(_nw82).ArraySet1(_664_v20, 9)
			(_nw82).ArraySet1(_664_v20, 10)
			(_nw82).ArraySet1(_664_v20, 11)
			_665_v21 = _nw82
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_665_v21), 0))
			_ = _index81
			(_665_v21).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v19, _dafny.IntOfInt64(196)), (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(318), _dafny.ArrayLen((_665_v21), 0))
			_ = _index82
			(_665_v21).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_663_v19, _657_i0)).Merge((_664_v20).Merge(_664_v20)), (_index82).Int())
		}
	}
}
func (_this *C3) F34() _dafny.MultiSet {
	{
		return _this._f34
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f26 bool
	_f27 bool
	_f33 _dafny.Map
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f26 = false
	_this._f27 = false
	_this._f33 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F26() bool {
	return _this._f26
}
func (_this *C4) F27() bool {
	return _this._f27
}
func (_this *C4) Ctor__(f33 _dafny.Map, f26 bool, f27 bool) {
	{
		(_this)._f33 = f33
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C4) Fm1(globalState *GlobalState) _dafny.Int {
	{
		var _source13 _dafny.Set = _dafny.SetOf((_this).F27())
		_ = _source13
		var _666___mcc_h0 _dafny.Set = _source13
		_ = _666___mcc_h0
		var _667_cf23 _dafny.Set = _666___mcc_h0
		_ = _667_cf23
		return _dafny.IntOfUint32((_dafny.SeqOf(!((_this).F26()), (_this).F26(), (_this).F26())).Cardinality())
	}
}
func (_this *C4) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F27()
	}
}
func (_this *C4) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _668_v0 _dafny.Int
		_ = _668_v0
		_668_v0 = _dafny.IntOfInt64(32)
		(globalState).F12 = _668_v0
		var _hi9 _dafny.Int = _668_v0
		_ = _hi9
		for _669_i0 := _668_v0; _669_i0.Cmp(_hi9) < 0; _669_i0 = _669_i0.Plus(_dafny.One) {
			if true {
				var _670_v1 _dafny.MultiSet
				_ = _670_v1
				_670_v1 = _dafny.MultiSetOf((_dafny.Zero).Minus(_669_i0))
				var _671_v2 _dafny.MultiSet
				_ = _671_v2
				_671_v2 = _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((p2).Cardinality()))).IsProperSubsetOf(_670_v1), !((_this).F26()) || (p0))
				(globalState).F12 = (func() _dafny.Int {
					if (_671_v2).Contains((_this).F26()) {
						return (_671_v2).Multiplicity((_this).F26())
					}
					return (_669_i0).Plus(_668_v0)
				})()
				var _672_v3 _dafny.Array
				_ = _672_v3
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_20
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) D2 = (func(_673_v0 _dafny.Int) func(_dafny.Int) D2 {
						return func(_674_i1 _dafny.Int) D2 {
							return Companion_D2_.Create_DC7_((_this).F26(), (_this).F27(), _673_v0, _673_v0)
						}
					})(_668_v0)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw83 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw83).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw83).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_672_v3 = _nw83
				_672_v3 = _672_v3
				var _675_v4 _dafny.Map
				_ = _675_v4
				_675_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _dafny.IntOfInt64(136))
				(globalState).F12 = (_669_i0).Plus((func() _dafny.Int {
					if (_675_v4).Contains(_669_i0) {
						return (_675_v4).Get(_669_i0).(_dafny.Int)
					}
					return (_671_v2).Cardinality()
				})())
				var _676_v5 *C0
				_ = _676_v5
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__((_this).F27())
				_676_v5 = _nw84
				(globalState).F15 = !(!((_668_v0).Cmp(_669_i0) >= 0))
			} else {
				var _677_v6 _dafny.Map
				_ = _677_v6
				_677_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), p2)
				var _678_v7 D5
				_ = _678_v7
				_678_v7 = Companion_D5_.Create_DC13_(_677_v6)
				_677_v6 = ((_678_v7).Dtor_cf24()).Merge(_677_v6)
				var _679_v8 D1
				_ = _679_v8
				_679_v8 = Companion_D1_.Create_DC3_(p2)
				var _680_v9 *C0
				_ = _680_v9
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__(p0)
				_680_v9 = _nw85
				var _681_v10 _dafny.Sequence
				_ = _681_v10
				_681_v10 = _dafny.SeqOf(_680_v9)
				var _682_v11 _dafny.Map
				_ = _682_v11
				_682_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_v8, _dafny.Companion_Sequence_.Contains(_681_v10, _680_v9))
				var _683_v12 _dafny.Map
				_ = _683_v12
				_683_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F27())
				var _684_v13 _dafny.Sequence
				_ = _684_v13
				_684_v13 = _dafny.SeqOf(_682_v11)
				var _685_v15 _dafny.MultiSet
				_ = _685_v15
				_685_v15 = _dafny.MultiSetOf(_679_v8, Companion_D1_.Create_DC3_(p2))
				_682_v11 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_v8, (func() bool {
					if (_683_v12).Contains((_680_v9).F32()) {
						return (_683_v12).Get((_680_v9).F32()).(bool)
					}
					return (_this).F27()
				})())).Merge(((_684_v13).Select((Companion_Default___.SafeIndex(_669_i0, _dafny.IntOfUint32((_684_v13).Cardinality()))).Uint32()).(_dafny.Map)).Merge(func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter51 := _dafny.Iterate((_685_v15).Elements()); ; {
						_compr_51, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _686_v14 D1
						_686_v14 = interface{}(_compr_51).(D1)
						if (_685_v15).Contains(_686_v14) {
							_coll51.Add(_686_v14, (_680_v9).F32())
						}
					}
					return _coll51.ToMap()
				}()))
				var _rhs86 _dafny.Int = _668_v0
				_ = _rhs86
				var _rhs87 _dafny.Int = _668_v0
				_ = _rhs87
				var _lhs78 *GlobalState = globalState
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_lhs78.F12 = _rhs86
				_lhs79.F12 = _rhs87
				var _687_v16 D5
				_ = _687_v16
				_687_v16 = Companion_D5_.Create_DC14_(_dafny.IntOfUint32((p2).Cardinality()), _668_v0, true)
				var _688_v17 _dafny.Array
				_ = _688_v17
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
				_ = _nw86
				_688_v17 = _nw86
				var _689_v18 _dafny.Sequence
				_ = _689_v18
				_689_v18 = _dafny.SeqOf(_688_v17, _688_v17)
				(globalState).F12 = (((_687_v16).Dtor_cf26()).Plus(_dafny.IntOfUint32((_689_v18).Cardinality()))).Minus(_668_v0)
				var _690_v19 _dafny.Sequence
				_ = _690_v19
				_690_v19 = _dafny.UnicodeSeqOfUtf8Bytes("ldus")
				(globalState).F15 = (_669_i0).Cmp((_dafny.IntOfUint32((_690_v19).Cardinality())).Times(_669_i0)) < 0
			}
			var _691_v20 _dafny.Array
			_ = _691_v20
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_21
			var _nw87 _dafny.Array
			_ = _nw87
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw87 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_692_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_693_i2 _dafny.Int) _dafny.Int {
						return (_693_i2).Plus((Companion_D5_.Create_DC14_(_692_v0, _692_v0, true)).Dtor_cf26())
					}
				})(_668_v0)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw87 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw87).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw87).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_691_v20 = _nw87
			var _694_v21 _dafny.CodePoint
			_ = _694_v21
			_694_v21 = _dafny.CodePoint('o')
			var _695_v22 _dafny.Set
			_ = _695_v22
			_695_v22 = _dafny.SetOf(_694_v21)
			var _696_v23 _dafny.Map
			_ = _696_v23
			_696_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v22, _668_v0)
			var _697_v24 _dafny.Sequence
			_ = _697_v24
			_697_v24 = _dafny.UnicodeSeqOfUtf8Bytes("cdyf")
			var _698_v25 _dafny.Sequence
			_ = _698_v25
			_698_v25 = _dafny.SeqOf(_668_v0, _669_i0)
			var _699_v26 _dafny.MultiSet
			_ = _699_v26
			_699_v26 = _dafny.MultiSetOf(_698_v25, _698_v25)
			var _700_v27 T1
			_ = _700_v27
			var _nw88 *C3 = New_C3_()
			_ = _nw88
			_nw88.Ctor__(_699_v26, p0, p0)
			_700_v27 = _nw88
			var _701_v28 _dafny.Array
			_ = _701_v28
			var _nwElement0_11 _dafny.Int = (_669_i0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yghsv")).Cardinality()))
			_ = _nwElement0_11
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(14))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_11, 0)
			(_nw89).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v20, (_696_v23).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_691_v20, _669_i0))).Cardinality(), 1)
			(_nw89).ArraySet1(_668_v0, 2)
			(_nw89).ArraySet1(_668_v0, 3)
			(_nw89).ArraySet1((_dafny.Zero).Minus(_668_v0), 4)
			(_nw89).ArraySet1((_dafny.Zero).Minus(_669_i0), 5)
			(_nw89).ArraySet1(_669_i0, 6)
			(_nw89).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_697_v24, _697_v24)).Cardinality()), 7)
			(_nw89).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(546))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_702_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_703_i3 _dafny.Int) _dafny.Int {
					return _702_v0
				}
			})(_668_v0))), _700_v27)).Cardinality(), 8)
			(_nw89).ArraySet1(_669_i0, 9)
			(_nw89).ArraySet1(_668_v0, 10)
			(_nw89).ArraySet1(_669_i0, 11)
			(_nw89).ArraySet1(_669_i0, 12)
			(_nw89).ArraySet1(_668_v0, 13)
			_701_v28 = _nw89
			var _704_v29 D3
			_ = _704_v29
			_704_v29 = Companion_D3_.Create_DC10_(_701_v28)
			_691_v20 = (_704_v29).Dtor_cf21()
			var _705_v30 D7
			_ = _705_v30
			_705_v30 = Companion_D7_.Create_DC19_(_669_i0, _dafny.MultiSetOf(!(p0), (_this).F26()), false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _668_v0))
			var _706_v31 _dafny.MultiSet
			_ = _706_v31
			_706_v31 = _dafny.MultiSetOf((_700_v27).F27())
			var _pat_let_tv2 = _694_v21
			_ = _pat_let_tv2
			var _pat_let_tv3 = _706_v31
			_ = _pat_let_tv3
			var _707_v32 _dafny.Array
			_ = _707_v32
			var _nwElement0_12 D7 = (func() D7 {
				if (_this).F27() {
					return _705_v30
				}
				return _705_v30
			})()
			_ = _nwElement0_12
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(12))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_12, 0)
			(_nw90).ArraySet1(func(_pat_let6_0 D7) D7 {
				return func(_708_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let7_0 _dafny.Int) D7 {
						return func(_711_dt__update_hcf36_h0 _dafny.Int) D7 {
							return func(_pat_let8_0 bool) D7 {
								return func(_712_dt__update_hcf38_h0 bool) D7 {
									return func(_pat_let9_0 _dafny.MultiSet) D7 {
										return func(_713_dt__update_hcf37_h0 _dafny.MultiSet) D7 {
											return Companion_D7_.Create_DC19_(_711_dt__update_hcf36_h0, _713_dt__update_hcf37_h0, _712_dt__update_hcf38_h0, (_708_dt__update__tmp_h0).Dtor_cf39())
										}(_pat_let9_0)
									}(_pat_let_tv3)
								}(_pat_let8_0)
							}((_this).F26())
						}(_pat_let7_0)
					}(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}((func(_709_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_710_i4 _dafny.Int) _dafny.CodePoint {
							return _709_v21
						}
					})(_pat_let_tv2)))).Cardinality()))
				}(_pat_let6_0)
			}(_705_v30), 1)
			(_nw90).ArraySet1(_705_v30, 2)
			(_nw90).ArraySet1(_705_v30, 3)
			(_nw90).ArraySet1(_705_v30, 4)
			(_nw90).ArraySet1(_705_v30, 5)
			(_nw90).ArraySet1(_705_v30, 6)
			(_nw90).ArraySet1(_705_v30, 7)
			(_nw90).ArraySet1(_705_v30, 8)
			(_nw90).ArraySet1(_705_v30, 9)
			(_nw90).ArraySet1((func() D7 {
				if p0 {
					return _705_v30
				}
				return _705_v30
			})(), 10)
			(_nw90).ArraySet1(_705_v30, 11)
			_707_v32 = _nw90
			var _714_v33 _dafny.Map
			_ = _714_v33
			_714_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_700_v27).F26(), _dafny.IntOfInt64(637))
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_707_v32), 0))
			_ = _index83
			(_707_v32).ArraySet1(Companion_D7_.Create_DC19_(_668_v0, _706_v31, p0, _714_v33), (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_707_v32), 0))
			_ = _index84
			(_707_v32).ArraySet1(_705_v30, (_index84).Int())
			var _715_v34 _dafny.Array
			_ = _715_v34
			var _nwElement0_13 bool = (_700_v27).F26()
			_ = _nwElement0_13
			var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(4))
			_ = _nw91
			(_nw91).ArraySet1(_nwElement0_13, 0)
			(_nw91).ArraySet1((_700_v27).F27(), 1)
			(_nw91).ArraySet1(false, 2)
			(_nw91).ArraySet1((_700_v27).F26(), 3)
			_715_v34 = _nw91
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_715_v34), 0))
			_ = _index85
			(_715_v34).ArraySet1((func() bool {
				if true {
					return (_700_v27).Fm2((_700_v27).F26(), _669_i0, globalState)
				}
				return (_700_v27).F27()
			})(), (_index85).Int())
			var _716_v35 _dafny.Array
			_ = _716_v35
			var _nw92 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(29))
			_ = _nw92
			_716_v35 = _nw92
			var _717_v36 _dafny.Map
			_ = _717_v36
			_717_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _716_v35)
			var _718_v37 _dafny.Map
			_ = _718_v37
			_718_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_669_i0, _669_i0)
			var _719_v38 _dafny.Map
			_ = _719_v38
			_719_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v36, (_718_v37).Cardinality())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_715_v34), 0))
			_ = _index86
			var _rhs88 bool = !((_668_v0).Cmp(_669_i0) != 0) || ((_this).F26())
			_ = _rhs88
			var _rhs89 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fxf")
			_ = _rhs89
			var _rhs90 _dafny.Int = Companion_Default___.SafeModuloInt(_668_v0, _668_v0)
			_ = _rhs90
			var _rhs91 bool = ((_668_v0).Cmp((func() _dafny.Int {
				if (_719_v38).Contains(_717_v36) {
					return (_719_v38).Get(_717_v36).(_dafny.Int)
				}
				return (_698_v25).Select((Companion_Default___.SafeIndex(_669_i0, _dafny.IntOfUint32((_698_v25).Cardinality()))).Uint32()).(_dafny.Int)
			})()) != 0) || ((_669_i0).Cmp(_668_v0) <= 0)
			_ = _rhs91
			var _lhs80 _dafny.Array = _715_v34
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_715_v34), 0))
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			var _lhs83 *GlobalState = globalState
			_ = _lhs83
			var _lhs84 *GlobalState = globalState
			_ = _lhs84
			(_lhs80).ArraySet1(_rhs88, (_lhs81).Int())
			_lhs82.F6 = _rhs89
			_lhs83.F12 = _rhs90
			_lhs84.F22 = _rhs91
		}
		var _720_v39 _dafny.CodePoint
		_ = _720_v39
		_720_v39 = _dafny.CodePoint('q')
		var _721_v40 _dafny.Set
		_ = _721_v40
		_721_v40 = _dafny.SetOf((_this).F26())
		var _722_v41 _dafny.Array
		_ = _722_v41
		var _nwElement0_14 _dafny.CodePoint = _720_v39
		_ = _nwElement0_14
		var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(21))
		_ = _nw93
		(_nw93).ArraySet1CodePoint(_nwElement0_14, 0)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('f'), 1)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('l'), 2)
		(_nw93).ArraySet1CodePoint(Companion_Default___.Fm14(p0, _668_v0, (_721_v40).Cardinality(), globalState), 3)
		(_nw93).ArraySet1CodePoint(_720_v39, 4)
		(_nw93).ArraySet1CodePoint(_720_v39, 5)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('v'), 6)
		(_nw93).ArraySet1CodePoint(_720_v39, 7)
		(_nw93).ArraySet1CodePoint(_720_v39, 8)
		(_nw93).ArraySet1CodePoint(_720_v39, 9)
		(_nw93).ArraySet1CodePoint(_720_v39, 10)
		(_nw93).ArraySet1CodePoint(_720_v39, 11)
		(_nw93).ArraySet1CodePoint(_720_v39, 12)
		(_nw93).ArraySet1CodePoint(_720_v39, 13)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('u'), 14)
		(_nw93).ArraySet1CodePoint(_720_v39, 15)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('o'), 16)
		(_nw93).ArraySet1CodePoint(_720_v39, 17)
		(_nw93).ArraySet1CodePoint(_720_v39, 18)
		(_nw93).ArraySet1CodePoint(_dafny.CodePoint('m'), 19)
		(_nw93).ArraySet1CodePoint(_720_v39, 20)
		_722_v41 = _nw93
		var _723_v42 _dafny.Map
		_ = _723_v42
		_723_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, false)
		var _724_v43 _dafny.Sequence
		_ = _724_v43
		_724_v43 = _dafny.SeqOf(_723_v42, _723_v42)
		var _725_v44 D2
		_ = _725_v44
		_725_v44 = Companion_D2_.Create_DC7_(true, false, _668_v0, (_dafny.MultiSetFromSeq(_724_v43)).Cardinality())
		var _726_v45 _dafny.Map
		_ = _726_v45
		_726_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_725_v44, _725_v44), _668_v0)
		var _727_v46 _dafny.Sequence
		_ = _727_v46
		_727_v46 = _dafny.SeqOf(_725_v44, _725_v44)
		var _728_v47 _dafny.Sequence
		_ = _728_v47
		_728_v47 = _dafny.UnicodeSeqOfUtf8Bytes("baa")
		var _729_v48 D1
		_ = _729_v48
		_729_v48 = Companion_D1_.Create_DC4_(_722_v41, _668_v0, (func() _dafny.Int {
			if (_726_v45).Contains(_727_v46) {
				return (_726_v45).Get(_727_v46).(_dafny.Int)
			}
			return _668_v0
		})(), _728_v47, _dafny.Companion_Sequence_.Update(_728_v47, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _668_v0)).Cardinality()), _dafny.IntOfUint32((_728_v47).Cardinality()))).Uint32(), _720_v39))
		var _source14 D1 = _729_v48
		_ = _source14
		if _source14.Is_DC4() {
			var _730___mcc_h0 _dafny.Array = _source14.Get_().(D1_DC4).Cf7
			_ = _730___mcc_h0
			var _731___mcc_h1 _dafny.Int = _source14.Get_().(D1_DC4).Cf8
			_ = _731___mcc_h1
			var _732___mcc_h2 _dafny.Int = _source14.Get_().(D1_DC4).Cf9
			_ = _732___mcc_h2
			var _733___mcc_h3 _dafny.Sequence = _source14.Get_().(D1_DC4).Cf10
			_ = _733___mcc_h3
			var _734___mcc_h4 _dafny.Sequence = _source14.Get_().(D1_DC4).Cf11
			_ = _734___mcc_h4
			var _735_cf11 _dafny.Sequence = _734___mcc_h4
			_ = _735_cf11
			var _736_cf10 _dafny.Sequence = _733___mcc_h3
			_ = _736_cf10
			var _737_cf9 _dafny.Int = _732___mcc_h2
			_ = _737_cf9
			var _738_cf8 _dafny.Int = _731___mcc_h1
			_ = _738_cf8
			var _739_cf7 _dafny.Array = _730___mcc_h0
			_ = _739_cf7
			var _740_v49 _dafny.Sequence
			_ = _740_v49
			_740_v49 = _dafny.SeqOf(_668_v0, _738_cf8, _737_cf9)
			var _741_v50 _dafny.Sequence
			_ = _741_v50
			_741_v50 = _dafny.SeqOf(_740_v49, _740_v49, _dafny.Companion_Sequence_.Update(_740_v49, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_740_v49).Cardinality()), _dafny.IntOfUint32((_740_v49).Cardinality()))).Uint32(), _738_cf8), _dafny.SeqOf(_668_v0), _740_v49)
			var _742_v51 *C2
			_ = _742_v51
			var _nw94 *C2 = New_C2_()
			_ = _nw94
			_nw94.Ctor__(_741_v50, false, (_this).F26())
			_742_v51 = _nw94
			var _743_v52 _dafny.Sequence
			_ = _743_v52
			_743_v52 = _dafny.SeqOf(_742_v51)
			(globalState).F12 = (func() _dafny.Int {
				if (_this).F26() {
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).Fm1(globalState)), _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), !((_this).F26()))).Cardinality()), _737_cf9)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_743_v52, _743_v52)).Cardinality())
			})()
			var _744_v53 _dafny.Map
			_ = _744_v53
			_744_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), _738_cf8)
			var _745_v54 D0
			_ = _745_v54
			_745_v54 = Companion_D0_.Create_DC0_((_this).F26())
			_744_v53 = (_744_v53).Update(_738_cf8, (_742_v51).Fm4(_745_v54, _668_v0, _737_cf9, globalState))
			(globalState).F2 = !((_this).F26())
			if !((_this).F27()) {
				var _746_v55 _dafny.Array
				_ = _746_v55
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw95
				_746_v55 = _nw95
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_746_v55), 0))
				_ = _index87
				(_746_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm15(globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(191))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_747_cf8 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_748_i5 _dafny.Int) _dafny.Int {
						return _747_cf8
					}
				})(_738_cf8)))), (_index87).Int())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_746_v55), 0))
				_ = _index88
				var _rhs92 _dafny.Int = (_dafny.Zero).Minus((_742_v51).Fm4(_745_v54, _737_cf9, _668_v0, globalState))
				_ = _rhs92
				var _rhs93 bool = (p2).Select((Companion_Default___.SafeIndex(_668_v0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)
				_ = _rhs93
				var _rhs94 _dafny.Sequence = _736_cf10
				_ = _rhs94
				var _rhs95 _dafny.Sequence = Companion_Default___.Fm23(_720_v39, globalState)
				_ = _rhs95
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _746_v55
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_746_v55), 0))
				_ = _lhs88
				_737_cf9 = _rhs92
				_lhs85.F15 = _rhs93
				_lhs86.F6 = _rhs94
				(_lhs87).ArraySet1(_rhs95, (_lhs88).Int())
				(globalState).F19 = _720_v39
				(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_728_v47, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_668_v0), _dafny.IntOfUint32((_728_v47).Cardinality()))).Uint32(), _720_v39), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(472))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}((func(_749_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_750_i6 _dafny.Int) _dafny.CodePoint {
						return _749_v39
					}
				})(_720_v39))))
				var _751_v56 _dafny.Set
				_ = _751_v56
				_751_v56 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("riqaxgwa"), _dafny.UnicodeSeqOfUtf8Bytes("jqtgis"))
				(globalState).F24 = (_742_v51).Fm2((_751_v56).IsProperSubsetOf(_751_v56), _737_cf9, globalState)
				(globalState).F15 = (_this).F26()
			} else {
				var _752_v57 _dafny.Array
				_ = _752_v57
				var _nw96 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
				_ = _nw96
				_752_v57 = _nw96
				var _753_v58 _dafny.Array
				_ = _753_v58
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(11))
				_ = _nw97
				_753_v58 = _nw97
				var _754_v59 _dafny.Map
				_ = _754_v59
				_754_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F26()), _739_cf7)
				var _755_v60 _dafny.Sequence
				_ = _755_v60
				_755_v60 = _dafny.SeqOf(_739_cf7)
				var _756_v61 _dafny.Array
				_ = _756_v61
				var _nwElement0_15 _dafny.Array = _722_v41
				_ = _nwElement0_15
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(15))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_15, 0)
				(_nw98).ArraySet1(_739_cf7, 1)
				(_nw98).ArraySet1(_739_cf7, 2)
				(_nw98).ArraySet1(_739_cf7, 3)
				(_nw98).ArraySet1(_739_cf7, 4)
				(_nw98).ArraySet1(_722_v41, 5)
				(_nw98).ArraySet1(_722_v41, 6)
				(_nw98).ArraySet1(_739_cf7, 7)
				(_nw98).ArraySet1(_722_v41, 8)
				(_nw98).ArraySet1(_739_cf7, 9)
				(_nw98).ArraySet1((func() _dafny.Array {
					if true {
						return _739_cf7
					}
					return _722_v41
				})(), 10)
				(_nw98).ArraySet1(_739_cf7, 11)
				(_nw98).ArraySet1((func() _dafny.Array {
					if (_754_v59).Contains(true) {
						return (_754_v59).Get(true).(_dafny.Array)
					}
					return _739_cf7
				})(), 12)
				(_nw98).ArraySet1(_739_cf7, 13)
				(_nw98).ArraySet1((_755_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-535), _dafny.IntOfUint32((_755_v60).Cardinality()))).Uint32()).(_dafny.Array), 14)
				_756_v61 = _nw98
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_756_v61), 0))
				_ = _index89
				(_756_v61).ArraySet1(_722_v41, (_index89).Int())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_756_v61), 0))
				_ = _index90
				var _rhs96 _dafny.Array = (func() _dafny.Array {
					if p0 {
						return _752_v57
					}
					return _752_v57
				})()
				_ = _rhs96
				var _rhs97 _dafny.Array = _753_v58
				_ = _rhs97
				var _rhs98 _dafny.Array = _722_v41
				_ = _rhs98
				var _lhs89 _dafny.Array = _756_v61
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_756_v61), 0))
				_ = _lhs90
				_752_v57 = _rhs96
				_753_v58 = _rhs97
				(_lhs89).ArraySet1(_rhs98, (_lhs90).Int())
				_737_cf9 = (func() _dafny.Int {
					if (_this).F27() {
						return _737_cf9
					}
					return (_dafny.IntOfUint32((_736_cf10).Cardinality())).Times(_737_cf9)
				})()
				var _757_v62 _dafny.Set
				_ = _757_v62
				_757_v62 = _dafny.SetOf(_dafny.IntOfInt64(-955))
				var _758_v63 D2
				_ = _758_v63
				_758_v63 = Companion_D2_.Create_DC9_((_757_v62).Cardinality(), _668_v0, false)
				var _759_v64 _dafny.MultiSet
				_ = _759_v64
				_759_v64 = _dafny.MultiSetOf(!((_this).F27()), (_this).F26(), (_this).F26())
				var _760_v65 _dafny.Sequence
				_ = _760_v65
				_760_v65 = _dafny.SeqOf(p2)
				var _761_v66 _dafny.Array
				_ = _761_v66
				var _nwElement0_16 _dafny.Int = _738_cf8
				_ = _nwElement0_16
				var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(12))
				_ = _nw99
				(_nw99).ArraySet1(_nwElement0_16, 0)
				(_nw99).ArraySet1((_758_v63).Dtor_cf18(), 1)
				(_nw99).ArraySet1((_759_v64).Cardinality(), 2)
				(_nw99).ArraySet1((_dafny.Zero).Minus(_737_cf9), 3)
				(_nw99).ArraySet1((_dafny.MultiSetOf((_this).F27(), p0, p0, false)).Cardinality(), 4)
				(_nw99).ArraySet1(_668_v0, 5)
				(_nw99).ArraySet1(_738_cf8, 6)
				(_nw99).ArraySet1(_738_cf8, 7)
				(_nw99).ArraySet1(_dafny.IntOfUint32((_735_cf11).Cardinality()), 8)
				(_nw99).ArraySet1(_668_v0, 9)
				(_nw99).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p2, (_760_v65).Select((Companion_Default___.SafeIndex((_742_v51).Fm1(globalState), _dafny.IntOfUint32((_760_v65).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), 10)
				(_nw99).ArraySet1(_668_v0, 11)
				_761_v66 = _nw99
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_761_v66), 0))
				_ = _index91
				(_761_v66).ArraySet1((_668_v0).Times(_737_cf9), (_index91).Int())
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_761_v66), 0))
				_ = _index92
				var _rhs99 _dafny.Int = _737_cf9
				_ = _rhs99
				var _rhs100 _dafny.Array = _761_v66
				_ = _rhs100
				var _rhs101 bool = !(!((func() bool {
					if (_this).F26() {
						return p0
					}
					return (func() bool {
						if (_723_v42).Contains(_738_cf8) {
							return (_723_v42).Get(_738_cf8).(bool)
						}
						return p0
					})()
				})())) || (((_dafny.Zero).Minus(_737_cf9)).Cmp(_738_cf8) < 0)
				_ = _rhs101
				var _rhs102 bool = p0
				_ = _rhs102
				var _rhs103 _dafny.Int = Companion_Default___.SafeModuloInt(_737_cf9, _668_v0)
				_ = _rhs103
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 _dafny.Array = _761_v66
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_761_v66), 0))
				_ = _lhs95
				_lhs91.F16 = _rhs99
				_761_v66 = _rhs100
				_lhs92.F15 = _rhs101
				_lhs93.F2 = _rhs102
				(_lhs94).ArraySet1(_rhs103, (_lhs95).Int())
				var _762_v67 _dafny.Sequence
				_ = _762_v67
				_762_v67 = _dafny.SeqOf(p0)
				_762_v67 = _762_v67
				var _763_v68 D7
				_ = _763_v68
				_763_v68 = Companion_D7_.Create_DC19_(_668_v0, (_759_v64).Update((_this).F27(), Companion_Default___.Abs(_737_cf9)), (_this).F26(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _737_cf9))
				var _764_v69 _dafny.MultiSet
				_ = _764_v69
				_764_v69 = _dafny.MultiSetOf(_738_cf8, (_dafny.Zero).Minus(_737_cf9))
				var _765_v70 _dafny.Map
				_ = _765_v70
				_765_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_742_v51).Fm2(p0, _668_v0, globalState))
				var _766_v71 _dafny.Array
				_ = _766_v71
				var _nwElement0_17 bool = false
				_ = _nwElement0_17
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(26))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_17, 0)
				(_nw100).ArraySet1(false, 1)
				(_nw100).ArraySet1((_dafny.IntOfUint32((Companion_Default___.Fm22((func() _dafny.Int {
					if (_744_v53).Contains(_738_cf8) {
						return (_744_v53).Get(_738_cf8).(_dafny.Int)
					}
					return _737_cf9
				})(), (_this).F27(), globalState)).Cardinality())).Cmp(_738_cf8) > 0, 2)
				(_nw100).ArraySet1((_this).F27(), 3)
				(_nw100).ArraySet1((_this).F27(), 4)
				(_nw100).ArraySet1((_763_v68).Dtor_cf38(), 5)
				(_nw100).ArraySet1((_this).F27(), 6)
				(_nw100).ArraySet1(p0, 7)
				(_nw100).ArraySet1((_this).F27(), 8)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_728_v47, _728_v47), 9)
				(_nw100).ArraySet1((_dafny.MultiSetFromSeq(_740_v49)).IsDisjointFrom(_764_v69), 10)
				(_nw100).ArraySet1((_this).F27(), 11)
				(_nw100).ArraySet1((_this).F27(), 12)
				(_nw100).ArraySet1(true, 13)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_736_cf10, _dafny.UnicodeSeqOfUtf8Bytes("h")), 14)
				(_nw100).ArraySet1(((_this).F26()) && (p0), 15)
				(_nw100).ArraySet1((_this).F26(), 16)
				(_nw100).ArraySet1((_this).F27(), 17)
				(_nw100).ArraySet1((_this).F26(), 18)
				(_nw100).ArraySet1((func() bool {
					if (_765_v70).Contains((_this).F26()) {
						return (_765_v70).Get((_this).F26()).(bool)
					}
					return (_this).F27()
				})(), 19)
				(_nw100).ArraySet1(true, 20)
				(_nw100).ArraySet1(false, 21)
				(_nw100).ArraySet1((_this).F27(), 22)
				(_nw100).ArraySet1((_this).F27(), 23)
				(_nw100).ArraySet1(!((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F27())).Equals(_765_v70)), 24)
				(_nw100).ArraySet1(!(p0), 25)
				_766_v71 = _nw100
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_766_v71), 0))
				_ = _index93
				(_766_v71).ArraySet1(false, (_index93).Int())
				var _767_v72 _dafny.Map
				_ = _767_v72
				_767_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_761_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_761_v66), 0))).Int()).(_dafny.Int))
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_766_v71), 0))
				_ = _index94
				var _rhs104 bool = p0
				_ = _rhs104
				var _rhs105 _dafny.Int = (_this).Fm1(globalState)
				_ = _rhs105
				var _rhs106 _dafny.CodePoint = _720_v39
				_ = _rhs106
				var _rhs107 bool = ((func() bool {
					if (_742_v51).Fm2(!((_this).F27()), (_761_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_761_v66), 0))).Int()).(_dafny.Int), globalState) {
						return p0
					}
					return (_this).F27()
				})()) || (((_767_v72).Cardinality()).Cmp(_668_v0) != 0)
				_ = _rhs107
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 _dafny.Array = _766_v71
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_766_v71), 0))
				_ = _lhs100
				_lhs96.F22 = _rhs104
				_lhs97.F16 = _rhs105
				_lhs98.F19 = _rhs106
				(_lhs99).ArraySet1(_rhs107, (_lhs100).Int())
			}
		} else if _source14.Is_DC5() {
			var _768___mcc_h5 _dafny.Int = _source14.Get_().(D1_DC5).Cf12
			_ = _768___mcc_h5
			var _769_cf12 _dafny.Int = _768___mcc_h5
			_ = _769_cf12
			if (_this).F26() {
				var _770_v73 _dafny.Map
				_ = _770_v73
				_770_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _720_v39)
				var _771_v74 _dafny.Sequence
				_ = _771_v74
				_771_v74 = _dafny.SeqOf(_770_v73)
				(globalState).F15 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_771_v74, _771_v74)).Cardinality())).Cmp(_668_v0) >= 0
				var _772_v75 _dafny.MultiSet
				_ = _772_v75
				_772_v75 = _dafny.MultiSetOf(_dafny.SeqOf(_769_cf12))
				var _773_v76 *C3
				_ = _773_v76
				var _nw101 *C3 = New_C3_()
				_ = _nw101
				_nw101.Ctor__(_772_v75, (p2).Select((Companion_Default___.SafeIndex(_769_cf12, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), (_this).F27())
				_773_v76 = _nw101
				var _774_v77 _dafny.Sequence
				_ = _774_v77
				_774_v77 = _dafny.SeqOf(_769_cf12, _668_v0)
				var _775_v78 _dafny.Map
				_ = _775_v78
				_775_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-971))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_776_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_777_i7 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_dafny.SeqOf(_776_v0, _776_v0, (_dafny.Zero).Minus(_776_v0))).Cardinality())
					}
				})(_668_v0)))).Cardinality()), (_774_v77).Select((Companion_Default___.SafeIndex((_773_v76).Fm12(_dafny.IntOfUint32((_728_v47).Cardinality()), globalState), _dafny.IntOfUint32((_774_v77).Cardinality()))).Uint32()).(_dafny.Int))
				var _778_v79 _dafny.Map
				_ = _778_v79
				_778_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_v76, (_775_v78).Cardinality())
				var _779_v80 _dafny.Map
				_ = _779_v80
				_779_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_778_v79, _720_v39)
				var _780_v81 _dafny.Map
				_ = _780_v81
				_780_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _775_v78)
				var _781_v82 _dafny.Array
				_ = _781_v82
				var _nwElement0_18 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p2, _dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_769_cf12, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F26()))
				_ = _nwElement0_18
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(19))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_18, 0)
				(_nw102).ArraySet1(p2, 1)
				(_nw102).ArraySet1(_dafny.SeqOf(!(p0)), 2)
				(_nw102).ArraySet1(p2, 3)
				(_nw102).ArraySet1(Companion_Default___.Fm25(true, (_this).F26(), _720_v39, globalState), 4)
				(_nw102).ArraySet1(p2, 5)
				(_nw102).ArraySet1(p2, 6)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, p2), 7)
				(_nw102).ArraySet1(Companion_Default___.Fm25((_this).F27(), (_this).F26(), (func() _dafny.CodePoint {
					if (_779_v80).Contains(_778_v79) {
						return (_779_v80).Get(_778_v79).(_dafny.CodePoint)
					}
					return _720_v39
				})(), globalState), 8)
				(_nw102).ArraySet1(p2, 9)
				(_nw102).ArraySet1(p2, 10)
				(_nw102).ArraySet1(p2, 11)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_this).F26() {
						return p2
					}
					return _dafny.SeqOf(false)
				})(), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_773_v76).Fm12(_769_cf12, globalState), (func() _dafny.Map {
					if (_780_v81).Contains((_dafny.MultiSetFromSeq(Companion_Default___.Fm15(globalState))).Cardinality()) {
						return (_780_v81).Get((_dafny.MultiSetFromSeq(Companion_Default___.Fm15(globalState))).Cardinality()).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_769_cf12, _668_v0)
				})())).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_this).F26() {
						return p2
					}
					return _dafny.SeqOf(false)
				})()).Cardinality()))).Uint32(), (_this).F27()), 12)
				(_nw102).ArraySet1(p2, 13)
				(_nw102).ArraySet1(p2, 14)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_769_cf12, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F27()), 15)
				(_nw102).ArraySet1(p2, 16)
				(_nw102).ArraySet1(p2, 17)
				(_nw102).ArraySet1(Companion_Default___.Fm25((_this).F26(), (_this).F27(), _720_v39, globalState), 18)
				_781_v82 = _nw102
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_781_v82), 0))
				_ = _index95
				(_781_v82).ArraySet1(p2, (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_781_v82), 0))
				_ = _index96
				(_781_v82).ArraySet1(p2, (_index96).Int())
				var _782_v83 _dafny.Array
				_ = _782_v83
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw103
				_782_v83 = _nw103
				var _783_v84 D3
				_ = _783_v84
				_783_v84 = Companion_D3_.Create_DC10_(_782_v83)
				var _784_v85 _dafny.Sequence
				_ = _784_v85
				_784_v85 = _dafny.SeqOf((_783_v84).Dtor_cf21())
				var _785_v86 D2
				_ = _785_v86
				_785_v86 = Companion_D2_.Create_DC9_(_dafny.IntOfInt64(121), _668_v0, false)
				var _786_v87 _dafny.Array
				_ = _786_v87
				var _nwElement0_19 _dafny.Array = _782_v83
				_ = _nwElement0_19
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(27))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_19, 0)
				(_nw104).ArraySet1((_784_v85).Select((Companion_Default___.SafeIndex(_668_v0, _dafny.IntOfUint32((_784_v85).Cardinality()))).Uint32()).(_dafny.Array), 1)
				(_nw104).ArraySet1(_782_v83, 2)
				(_nw104).ArraySet1(_782_v83, 3)
				(_nw104).ArraySet1(_782_v83, 4)
				(_nw104).ArraySet1(_782_v83, 5)
				(_nw104).ArraySet1(_782_v83, 6)
				(_nw104).ArraySet1(_782_v83, 7)
				(_nw104).ArraySet1(_782_v83, 8)
				(_nw104).ArraySet1(_782_v83, 9)
				(_nw104).ArraySet1(_782_v83, 10)
				(_nw104).ArraySet1((func() _dafny.Array {
					if (_785_v86).Dtor_cf20() {
						return _782_v83
					}
					return _782_v83
				})(), 11)
				(_nw104).ArraySet1(_782_v83, 12)
				(_nw104).ArraySet1(_782_v83, 13)
				(_nw104).ArraySet1(_782_v83, 14)
				(_nw104).ArraySet1(_782_v83, 15)
				(_nw104).ArraySet1((func() _dafny.Array {
					if (_this).F26() {
						return _782_v83
					}
					return _782_v83
				})(), 16)
				(_nw104).ArraySet1(_782_v83, 17)
				(_nw104).ArraySet1(_782_v83, 18)
				(_nw104).ArraySet1(_782_v83, 19)
				(_nw104).ArraySet1(_782_v83, 20)
				(_nw104).ArraySet1(_782_v83, 21)
				(_nw104).ArraySet1(_782_v83, 22)
				(_nw104).ArraySet1((_783_v84).Dtor_cf21(), 23)
				(_nw104).ArraySet1(_782_v83, 24)
				(_nw104).ArraySet1(_782_v83, 25)
				(_nw104).ArraySet1(_782_v83, 26)
				_786_v87 = _nw104
				_786_v87 = _786_v87
				var _787_v88 _dafny.MultiSet
				_ = _787_v88
				_787_v88 = _dafny.MultiSetOf((_this).F27(), (_this).F27())
				var _788_v89 _dafny.Array
				_ = _788_v89
				var _nwElement0_20 bool = !(!(!(p0))) || (!(p0))
				_ = _nwElement0_20
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(25))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_20, 0)
				(_nw105).ArraySet1((_this).F27(), 1)
				(_nw105).ArraySet1((_this).F26(), 2)
				(_nw105).ArraySet1(p0, 3)
				(_nw105).ArraySet1(!(p0), 4)
				(_nw105).ArraySet1(p0, 5)
				(_nw105).ArraySet1((_this).F27(), 6)
				(_nw105).ArraySet1((_this).F26(), 7)
				(_nw105).ArraySet1(p0, 8)
				(_nw105).ArraySet1(false, 9)
				(_nw105).ArraySet1((_this).F27(), 10)
				(_nw105).ArraySet1((_this).F27(), 11)
				(_nw105).ArraySet1((_this).F27(), 12)
				(_nw105).ArraySet1(!((_787_v88).IsDisjointFrom(_dafny.MultiSetOf((_this).Fm2(p0, _668_v0, globalState)))), 13)
				(_nw105).ArraySet1(p0, 14)
				(_nw105).ArraySet1((_this).F26(), 15)
				(_nw105).ArraySet1(p0, 16)
				(_nw105).ArraySet1(p0, 17)
				(_nw105).ArraySet1(!((_this).F26()), 18)
				(_nw105).ArraySet1((_this).F27(), 19)
				(_nw105).ArraySet1((_this).F27(), 20)
				(_nw105).ArraySet1((_this).F27(), 21)
				(_nw105).ArraySet1((p0) && (false), 22)
				(_nw105).ArraySet1((_this).F27(), 23)
				(_nw105).ArraySet1((_dafny.IntOfUint32(((_781_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_781_v82), 0))).Int()).(_dafny.Sequence)).Cardinality())).Cmp(_668_v0) < 0, 24)
				_788_v89 = _nw105
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_788_v89), 0))
				_ = _index97
				(_788_v89).ArraySet1((_this).Fm2(p0, (_dafny.Zero).Minus(_769_cf12), globalState), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_788_v89), 0))
				_ = _index98
				(_788_v89).ArraySet1(p0, (_index98).Int())
				var _789_v90 _dafny.Sequence
				_ = _789_v90
				_789_v90 = _dafny.SeqOf(_dafny.SeqOf((_this).Fm2(p0, _668_v0, globalState), (_788_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_788_v89), 0))).Int()).(bool)), p2, (_781_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_781_v82), 0))).Int()).(_dafny.Sequence))
				var _790_v91 _dafny.Array
				_ = _790_v91
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw106
				_790_v91 = _nw106
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_790_v91), 0))
				_ = _index99
				(_790_v91).ArraySet1(_770_v73, (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_788_v89), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_790_v91), 0))
				_ = _index101
				var _rhs108 _dafny.Sequence = _789_v90
				_ = _rhs108
				var _rhs109 bool = true
				_ = _rhs109
				var _rhs110 _dafny.Int = (func() _dafny.Int {
					if (_775_v78).Contains(_769_cf12) {
						return (_775_v78).Get(_769_cf12).(_dafny.Int)
					}
					return _dafny.IntOfInt64(953)
				})()
				_ = _rhs110
				var _rhs111 _dafny.Map = _770_v73
				_ = _rhs111
				var _lhs101 _dafny.Array = _788_v89
				_ = _lhs101
				var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_788_v89), 0))
				_ = _lhs102
				var _lhs103 *GlobalState = globalState
				_ = _lhs103
				var _lhs104 _dafny.Array = _790_v91
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_790_v91), 0))
				_ = _lhs105
				_789_v90 = _rhs108
				(_lhs101).ArraySet1(_rhs109, (_lhs102).Int())
				_lhs103.F12 = _rhs110
				(_lhs104).ArraySet1(_rhs111, (_lhs105).Int())
			} else {
				var _791_v92 _dafny.Map
				_ = _791_v92
				_791_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _668_v0)
				var _792_v93 _dafny.Sequence
				_ = _792_v93
				_792_v93 = _dafny.SeqOf(_791_v92)
				var _793_v94 _dafny.Array
				_ = _793_v94
				var _nwElement0_21 bool = (_this).F27()
				_ = _nwElement0_21
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(26))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_21, 0)
				(_nw107).ArraySet1(p0, 1)
				(_nw107).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_728_v47, (Companion_Default___.SafeIndex(_769_cf12, _dafny.IntOfUint32((_728_v47).Cardinality()))).Uint32(), _dafny.CodePoint('e')), _dafny.UnicodeSeqOfUtf8Bytes("urnk")), 2)
				(_nw107).ArraySet1(!(p0) || ((_this).Fm2((_this).F27(), _668_v0, globalState)), 3)
				(_nw107).ArraySet1((_this).F27(), 4)
				(_nw107).ArraySet1((_this).Fm2(p0, _769_cf12, globalState), 5)
				(_nw107).ArraySet1((_this).F27(), 6)
				(_nw107).ArraySet1(!((_this).F27()), 7)
				(_nw107).ArraySet1(!((Companion_D5_.Create_DC14_(_668_v0, _769_cf12, (_this).F26())).Dtor_cf27()), 8)
				(_nw107).ArraySet1(p0, 9)
				(_nw107).ArraySet1(((_dafny.Zero).Minus(_769_cf12)).Cmp(_dafny.IntOfUint32((p2).Cardinality())) >= 0, 10)
				(_nw107).ArraySet1(p0, 11)
				(_nw107).ArraySet1((_this).F26(), 12)
				(_nw107).ArraySet1((_this).F27(), 13)
				(_nw107).ArraySet1(p0, 14)
				(_nw107).ArraySet1(p0, 15)
				(_nw107).ArraySet1((_this).F26(), 16)
				(_nw107).ArraySet1((_this).F27(), 17)
				(_nw107).ArraySet1((p2).Select((Companion_Default___.SafeIndex(_668_v0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), 18)
				(_nw107).ArraySet1((_this).F26(), 19)
				(_nw107).ArraySet1(p0, 20)
				(_nw107).ArraySet1((_this).F26(), 21)
				(_nw107).ArraySet1(true, 22)
				(_nw107).ArraySet1(p0, 23)
				(_nw107).ArraySet1((_this).F27(), 24)
				(_nw107).ArraySet1(!_dafny.Companion_Sequence_.Equal(_792_v93, _792_v93), 25)
				_793_v94 = _nw107
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))
				_ = _index102
				(_793_v94).ArraySet1((p2).Select((Companion_Default___.SafeIndex(_668_v0, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))
				_ = _index103
				(_793_v94).ArraySet1((_this).F27(), (_index103).Int())
				var _794_v95 _dafny.Array
				_ = _794_v95
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw108
				_794_v95 = _nw108
				var _795_v96 _dafny.Map
				_ = _795_v96
				_795_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_794_v95, (_793_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))).Int()).(bool))
				(globalState).F22 = (((_795_v96).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_794_v95, p0))).Cardinality()).Cmp(_769_cf12) < 0
				var _796_v97 *C1
				_ = _796_v97
				var _nw109 *C1 = New_C1_()
				_ = _nw109
				_nw109.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_728_v47, (Companion_Default___.SafeIndex(_668_v0, _dafny.IntOfUint32((_728_v47).Cardinality()))).Uint32(), _720_v39)).Cardinality()), _668_v0), (_this).Fm1(globalState), (_793_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))).Int()).(bool), (_this).F27())
				_796_v97 = _nw109
				var _797_v98 D3
				_ = _797_v98
				_797_v98 = Companion_D3_.Create_DC11_(_720_v39)
				var _798_v99 _dafny.Sequence
				_ = _798_v99
				_798_v99 = _dafny.SeqOf((_721_v40).Cardinality(), (_796_v97).F36())
				var _799_v100 _dafny.Map
				_ = _799_v100
				_799_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_793_v94, _dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm23((_797_v98).Dtor_cf22(), globalState), _798_v99))
				_799_v100 = (_799_v100).Update(_793_v94, !(false))
				var _800_v101 T2
				_ = _800_v101
				var _nw110 *C1 = New_C1_()
				_ = _nw110
				_nw110.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(false, (_this).F26())).Cardinality()), _796_v97.F37, (_793_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))).Int()).(bool), p0)
				_800_v101 = _nw110
				var _801_v102 D7
				_ = _801_v102
				_801_v102 = Companion_D7_.Create_DC18_(_800_v101)
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_794_v95), 0))
				_ = _index104
				(_794_v95).ArraySet1(_668_v0, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_794_v95), 0))
				_ = _index105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))
				_ = _index106
				var _rhs112 D7 = _801_v102
				_ = _rhs112
				var _rhs113 _dafny.Int = _796_v97.F37
				_ = _rhs113
				var _rhs114 bool = (_this).F26()
				_ = _rhs114
				var _rhs115 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_668_v0), _798_v99)
				_ = _rhs115
				var _rhs116 T2 = _800_v101
				_ = _rhs116
				var _lhs106 _dafny.Array = _794_v95
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_794_v95), 0))
				_ = _lhs107
				var _lhs108 _dafny.Array = _793_v94
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_793_v94), 0))
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				_801_v102 = _rhs112
				(_lhs106).ArraySet1(_rhs113, (_lhs107).Int())
				(_lhs108).ArraySet1(_rhs114, (_lhs109).Int())
				_lhs110.F2 = _rhs115
				_800_v101 = _rhs116
			}
			if !(!(true) || (p0)) {
				(globalState).F24 = !_dafny.Companion_Sequence_.Contains(p2, (_this).F27())
				(globalState).F22 = (_this).F26()
				var _rhs117 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_668_v0))
				_ = _rhs117
				var _rhs118 _dafny.CodePoint = _dafny.CodePoint('y')
				_ = _rhs118
				var _rhs119 _dafny.Int = _dafny.IntOfInt64(687)
				_ = _rhs119
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_lhs111.F16 = _rhs117
				_720_v39 = _rhs118
				_lhs112.F16 = _rhs119
				(globalState).F15 = ((_this).F26()) || (p0)
				var _802_v103 _dafny.Array
				_ = _802_v103
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_22
				var _nw111 _dafny.Array
				_ = _nw111
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw111 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.MultiSet = (func(_803_cf12 _dafny.Int, _804_v0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
						return func(_805_i8 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_803_cf12, _804_v0))
						}
					})(_769_cf12, _668_v0)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw111 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw111).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw111).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_802_v103 = _nw111
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_802_v103), 0))
				_ = _index107
				(_802_v103).ArraySet1(Companion_Default___.Fm26((_this).Fm2((_this).F26(), _769_cf12, globalState), p0, p0, globalState), (_index107).Int())
				var _806_v104 _dafny.Map
				_ = _806_v104
				_806_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _769_cf12)
				var _807_v105 _dafny.MultiSet
				_ = _807_v105
				_807_v105 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_769_cf12, _668_v0), (_806_v104).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-261), _769_cf12)), _806_v104)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_802_v103), 0))
				_ = _index108
				(_802_v103).ArraySet1(_807_v105, (_index108).Int())
			} else {
				var _808_v106 _dafny.MultiSet
				_ = _808_v106
				_808_v106 = _dafny.MultiSetOf(!((_this).Fm2(p0, (_this).Fm1(globalState), globalState)))
				var _rhs120 bool = !(!(!((_this).F27()))) || ((_668_v0).Cmp(_769_cf12) == 0)
				_ = _rhs120
				var _rhs121 _dafny.MultiSet = _808_v106
				_ = _rhs121
				var _lhs113 *GlobalState = globalState
				_ = _lhs113
				_lhs113.F2 = _rhs120
				_808_v106 = _rhs121
				var _809_v107 _dafny.Array
				_ = _809_v107
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_23
				var _nw112 _dafny.Array
				_ = _nw112
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw112 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Sequence = (func(_810_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_811_i9 _dafny.Int) _dafny.Sequence {
							return _810_p2
						}
					})(p2)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw112 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw112).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw112).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_809_v107 = _nw112
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_809_v107), 0))
				_ = _index109
				(_809_v107).ArraySet1(p2, (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_809_v107), 0))
				_ = _index110
				var _rhs122 _dafny.Int = _769_cf12
				_ = _rhs122
				var _rhs123 bool = (_769_cf12).Cmp(_668_v0) != 0
				_ = _rhs123
				var _rhs124 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if true {
						return p2
					}
					return p2
				})(), p2)
				_ = _rhs124
				var _lhs114 *GlobalState = globalState
				_ = _lhs114
				var _lhs115 _dafny.Array = _809_v107
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_809_v107), 0))
				_ = _lhs116
				_668_v0 = _rhs122
				_lhs114.F24 = _rhs123
				(_lhs115).ArraySet1(_rhs124, (_lhs116).Int())
				var _812_v108 D2
				_ = _812_v108
				_812_v108 = Companion_D2_.Create_DC8_()
				var _813_v109 _dafny.Set
				_ = _813_v109
				_813_v109 = _dafny.SetOf(_812_v108)
				_813_v109 = _813_v109
				(globalState).F15 = (_this).F26()
				var _814_v110 _dafny.Array
				_ = _814_v110
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_24
				var _nw113 _dafny.Array
				_ = _nw113
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw113 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) bool = (func(_815_p0 bool) func(_dafny.Int) bool {
						return func(_816_i10 _dafny.Int) bool {
							return _815_p0
						}
					})(p0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw113 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw113).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw113).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_814_v110 = _nw113
				var _817_v111 _dafny.Map
				_ = _817_v111
				_817_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_809_v107).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(74), _dafny.ArrayLen((_809_v107), 0))).Int()).(_dafny.Sequence), _814_v110)
				var _818_v112 D9
				_ = _818_v112
				_818_v112 = Companion_D9_.Create_DC21_(_817_v111)
				var _819_v113 _dafny.Sequence
				_ = _819_v113
				_819_v113 = _dafny.SeqOf(_818_v112)
				(globalState).F24 = _dafny.Companion_Sequence_.Contains(_819_v113, _818_v112)
			}
			var _820_v115 _dafny.Array
			_ = _820_v115
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_25
			var _nw114 _dafny.Array
			_ = _nw114
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw114 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Set = (func(_821_cf12 _dafny.Int, _822_v0 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_823_i11 _dafny.Int) _dafny.Set {
						return (func() _dafny.Set {
							var _coll52 = _dafny.NewBuilder()
							_ = _coll52
							for _iter52 := _dafny.Iterate((_dafny.SetOf(_822_v0)).Elements()); ; {
								_compr_52, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _824_v114 _dafny.Int
								_824_v114 = interface{}(_compr_52).(_dafny.Int)
								if (_dafny.SetOf(_822_v0)).Contains(_824_v114) {
									_coll52.Add((_824_v114).Plus((_dafny.Zero).Minus(_dafny.IntOfInt64(-311))))
								}
							}
							return _coll52.ToSet()
						}()).Intersection(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_822_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_822_v0, _821_cf12))).Cardinality()))
					}
				})(_769_cf12, _668_v0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw114 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw114).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw114).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_820_v115 = _nw114
			var _825_v116 D1
			_ = _825_v116
			_825_v116 = Companion_D1_.Create_DC3_((Companion_D1_.Create_DC3_(_dafny.SeqOf(p0))).Dtor_cf6())
			var _826_v117 _dafny.MultiSet
			_ = _826_v117
			_826_v117 = _dafny.MultiSetOf((Companion_Default___.Fm18(_721_v40, globalState)).Cardinality())
			var _827_v118 _dafny.Sequence
			_ = _827_v118
			_827_v118 = _dafny.SeqOf(_dafny.IntOfInt64(336))
			var _pat_let_tv4 = p0
			_ = _pat_let_tv4
			var _pat_let_tv5 = _826_v117
			_ = _pat_let_tv5
			var _rhs125 bool = ((_this).F27()) && ((_this).F26())
			_ = _rhs125
			var _rhs126 _dafny.Array = _820_v115
			_ = _rhs126
			var _rhs127 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((func(_pat_let10_0 D2) D2 {
				return func(_828_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let11_0 bool) D2 {
						return func(_829_dt__update_hcf15_h0 bool) D2 {
							return func(_pat_let12_0 _dafny.Int) D2 {
								return func(_830_dt__update_hcf16_h0 _dafny.Int) D2 {
									return Companion_D2_.Create_DC7_((_828_dt__update__tmp_h1).Dtor_cf14(), _829_dt__update_hcf15_h0, _830_dt__update_hcf16_h0, (_828_dt__update__tmp_h1).Dtor_cf17())
								}(_pat_let12_0)
							}((_pat_let_tv5).Cardinality())
						}(_pat_let11_0)
					}(_pat_let_tv4)
				}(_pat_let10_0)
			}(_725_v44)).Dtor_cf17(), (_827_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-276), _dafny.IntOfUint32((_827_v118).Cardinality()))).Uint32()).(_dafny.Int)))
			_ = _rhs127
			var _rhs128 D1 = _825_v116
			_ = _rhs128
			var _lhs117 *GlobalState = globalState
			_ = _lhs117
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			_lhs117.F15 = _rhs125
			_820_v115 = _rhs126
			_lhs118.F12 = _rhs127
			_825_v116 = _rhs128
			var _831_v119 _dafny.Map
			_ = _831_v119
			_831_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F26())
			_831_v119 = (_831_v119).Update(p0, p0)
		} else {
			var _832___mcc_h6 _dafny.Sequence = _source14.Get_().(D1_DC3).Cf6
			_ = _832___mcc_h6
			var _833_cf6 _dafny.Sequence = _832___mcc_h6
			_ = _833_cf6
			(globalState).F19 = _720_v39
			var _834_v120 _dafny.Map
			_ = _834_v120
			_834_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q"), _668_v0)
			(globalState).F12 = (_834_v120).Cardinality()
			var _835_v121 _dafny.CodePoint
			_ = _835_v121
			var _836_v122 _dafny.Sequence
			_ = _836_v122
			var _837_v123 bool
			_ = _837_v123
			var _838_v124 _dafny.Map
			_ = _838_v124
			var _out25 _dafny.CodePoint
			_ = _out25
			var _out26 _dafny.Sequence
			_ = _out26
			var _out27 bool
			_ = _out27
			var _out28 _dafny.Map
			_ = _out28
			_out25, _out26, _out27, _out28 = Companion_Default___.M0(_720_v39, globalState)
			_835_v121 = _out25
			_836_v122 = _out26
			_837_v123 = _out27
			_838_v124 = _out28
			if (func() bool {
				if !((_this).F26()) {
					return (func() bool {
						if p0 {
							return false
						}
						return (_this).Fm2((_this).F26(), _668_v0, globalState)
					})()
				}
				return p0
			})() {
				var _839_v125 _dafny.Map
				_ = _839_v125
				_839_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm1(globalState))
				(globalState).F15 = (_dafny.IntOfInt64(954)).Cmp((_668_v0).Times((func() _dafny.Int {
					if (_839_v125).Contains(_837_v123) {
						return (_839_v125).Get(_837_v123).(_dafny.Int)
					}
					return _dafny.IntOfInt64(869)
				})())) >= 0
				var _840_v126 _dafny.Map
				_ = _840_v126
				_840_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _668_v0)
				_840_v126 = (_840_v126).Update(_668_v0, _dafny.IntOfUint32((_836_v122).Cardinality()))
				_837_v123 = _837_v123
				var _841_v127 _dafny.Array
				_ = _841_v127
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_26
				var _nw115 _dafny.Array
				_ = _nw115
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw115 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) bool = func(_842_i12 _dafny.Int) bool {
						return (_this).F26()
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw115 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw115).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw115).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_841_v127 = _nw115
				var _843_v128 _dafny.Map
				_ = _843_v128
				_843_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(243), _841_v127)
				var _844_v129 _dafny.Array
				_ = _844_v129
				var _nwElement0_22 _dafny.Array = _841_v127
				_ = _nwElement0_22
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(23))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_22, 0)
				(_nw116).ArraySet1((func() _dafny.Array {
					if (_843_v128).Contains(_668_v0) {
						return (_843_v128).Get(_668_v0).(_dafny.Array)
					}
					return _841_v127
				})(), 1)
				(_nw116).ArraySet1(_841_v127, 2)
				(_nw116).ArraySet1(_841_v127, 3)
				(_nw116).ArraySet1(_841_v127, 4)
				(_nw116).ArraySet1(_841_v127, 5)
				(_nw116).ArraySet1(_841_v127, 6)
				(_nw116).ArraySet1(_841_v127, 7)
				(_nw116).ArraySet1(_841_v127, 8)
				(_nw116).ArraySet1(_841_v127, 9)
				(_nw116).ArraySet1(_841_v127, 10)
				(_nw116).ArraySet1(_841_v127, 11)
				(_nw116).ArraySet1(_841_v127, 12)
				(_nw116).ArraySet1(_841_v127, 13)
				(_nw116).ArraySet1(_841_v127, 14)
				(_nw116).ArraySet1(_841_v127, 15)
				(_nw116).ArraySet1(_841_v127, 16)
				(_nw116).ArraySet1(_841_v127, 17)
				(_nw116).ArraySet1(_841_v127, 18)
				(_nw116).ArraySet1(_841_v127, 19)
				(_nw116).ArraySet1(_841_v127, 20)
				(_nw116).ArraySet1(_841_v127, 21)
				(_nw116).ArraySet1(_841_v127, 22)
				_844_v129 = _nw116
				(globalState).F5 = _844_v129
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_841_v127), 0))
				_ = _index111
				(_841_v127).ArraySet1(p0, (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_841_v127), 0))
				_ = _index112
				var _rhs129 _dafny.CodePoint = _835_v121
				_ = _rhs129
				var _rhs130 bool = !(p0)
				_ = _rhs130
				var _rhs131 _dafny.Int = (_668_v0).Times((_dafny.Zero).Minus(_668_v0))
				_ = _rhs131
				var _rhs132 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_836_v122, _dafny.Companion_Sequence_.Concatenate(p2, _836_v122))
				_ = _rhs132
				var _rhs133 bool = (_this).F26()
				_ = _rhs133
				var _lhs119 *GlobalState = globalState
				_ = _lhs119
				var _lhs120 _dafny.Array = _841_v127
				_ = _lhs120
				var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_841_v127), 0))
				_ = _lhs121
				var _lhs122 *GlobalState = globalState
				_ = _lhs122
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				_lhs119.F19 = _rhs129
				(_lhs120).ArraySet1(_rhs130, (_lhs121).Int())
				_lhs122.F12 = _rhs131
				_836_v122 = _rhs132
				_lhs123.F15 = _rhs133
			} else {
				var _845_v130 _dafny.Set
				_ = _845_v130
				_845_v130 = _dafny.SetOf(_668_v0, _668_v0)
				var _846_v131 _dafny.Map
				_ = _846_v131
				_846_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, (_845_v130).Cardinality())
				var _847_v132 _dafny.Array
				_ = _847_v132
				var _nwElement0_23 _dafny.Int = _668_v0
				_ = _nwElement0_23
				var _nw117 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(13))
				_ = _nw117
				(_nw117).ArraySet1(_nwElement0_23, 0)
				(_nw117).ArraySet1(_668_v0, 1)
				(_nw117).ArraySet1(_668_v0, 2)
				(_nw117).ArraySet1(_668_v0, 3)
				(_nw117).ArraySet1(_668_v0, 4)
				(_nw117).ArraySet1(_dafny.IntOfInt64(250), 5)
				(_nw117).ArraySet1(_668_v0, 6)
				(_nw117).ArraySet1((_846_v131).Cardinality(), 7)
				(_nw117).ArraySet1(_dafny.IntOfInt64(492), 8)
				(_nw117).ArraySet1((_dafny.Zero).Minus(_668_v0), 9)
				(_nw117).ArraySet1(_668_v0, 10)
				(_nw117).ArraySet1(_668_v0, 11)
				(_nw117).ArraySet1(_668_v0, 12)
				_847_v132 = _nw117
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_847_v132), 0))
				_ = _index113
				(_847_v132).ArraySet1(_668_v0, (_index113).Int())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_847_v132), 0))
				_ = _index114
				(_847_v132).ArraySet1((_dafny.Zero).Minus((_668_v0).Plus(_668_v0)), (_index114).Int())
				(globalState).F24 = false
				var _848_v133 D7
				_ = _848_v133
				_848_v133 = Companion_D7_.Create_DC19_(_668_v0, Companion_Default___.Fm21((_847_v132).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_847_v132), 0))).Int()).(_dafny.Int), globalState), true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _668_v0))
				(globalState).F22 = !((func() bool {
					if (_this).F26() {
						return (_dafny.IntOfUint32((_728_v47).Cardinality())).Cmp(_dafny.IntOfUint32((_728_v47).Cardinality())) < 0
					}
					return (_848_v133).Dtor_cf38()
				})())
				(globalState).F2 = (false) || ((_this).F27())
				var _849_v134 _dafny.Map
				_ = _849_v134
				_849_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_728_v47, (_this).F26())
				var _850_v135 _dafny.Map
				_ = _850_v135
				_850_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_721_v40).Cardinality())
				_849_v134 = (Companion_Default___.Fm27((_850_v135).Cardinality(), _668_v0, globalState)).Merge((_849_v134).Update(_dafny.UnicodeSeqOfUtf8Bytes("bd"), _837_v123))
			}
		}
		var _851_v136 _dafny.Map
		_ = _851_v136
		_851_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_668_v0, _dafny.UnicodeSeqOfUtf8Bytes("ga"))
		var _852_v137 D0
		_ = _852_v137
		_852_v137 = Companion_D0_.Create_DC1_((_this).F26(), (_this).F26(), (_851_v136).Cardinality(), p0)
		var _853_i13 _dafny.Int
		_ = _853_i13
		_853_i13 = _dafny.Zero
		{
			for (Companion_Default___.Fm28(globalState)).Equals(_dafny.SetOf(_852_v137)) {
				{
					if (_853_i13).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_853_i13 = (_853_i13).Plus(_dafny.One)
					(globalState).F12 = _668_v0
					(globalState).F12 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_668_v0, _668_v0), _dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_this).F27() {
							return p2
						}
						return p2
					})()).Cardinality()))
					var _854_v138 _dafny.Array
					_ = _854_v138
					var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
					_ = _nw118
					_854_v138 = _nw118
					var _855_v139 _dafny.Sequence
					_ = _855_v139
					_855_v139 = _dafny.SeqOf(_854_v138)
					var _rhs134 _dafny.Int = Companion_Default___.SafeDivisionInt(_668_v0, _dafny.IntOfInt64(-344))
					_ = _rhs134
					var _rhs135 bool = ((_this).F27()) && (_dafny.Companion_Sequence_.Equal(_855_v139, _855_v139))
					_ = _rhs135
					var _lhs124 *GlobalState = globalState
					_ = _lhs124
					var _lhs125 *GlobalState = globalState
					_ = _lhs125
					_lhs124.F12 = _rhs134
					_lhs125.F15 = _rhs135
					var _856_v140 D5
					_ = _856_v140
					_856_v140 = Companion_D5_.Create_DC14_(_dafny.IntOfInt64(803), _668_v0, (_this).F27())
					var _857_v141 _dafny.Sequence
					_ = _857_v141
					_857_v141 = _dafny.SeqOf((_856_v140).Dtor_cf26(), (_this).Fm1(globalState))
					_668_v0 = (_857_v141).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(291), _dafny.IntOfUint32((_857_v141).Cardinality()))).Uint32()).(_dafny.Int)
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _858_v142 _dafny.Array
		_ = _858_v142
		var _nw119 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw119
		_858_v142 = _nw119
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_858_v142), 0))
		_ = _index115
		(_858_v142).ArraySet1((_this).F27(), (_index115).Int())
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_858_v142), 0))
		_ = _index116
		(_858_v142).ArraySet1((_668_v0).Cmp(_668_v0) >= 0, (_index116).Int())
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_858_v142), 0))
		_ = _index117
		var _rhs136 _dafny.Sequence = _728_v47
		_ = _rhs136
		var _rhs137 _dafny.Int = (_668_v0).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-402), (_721_v40).Cardinality()))
		_ = _rhs137
		var _rhs138 bool = p0
		_ = _rhs138
		var _rhs139 bool = !((_this).F26())
		_ = _rhs139
		var _lhs126 *GlobalState = globalState
		_ = _lhs126
		var _lhs127 *GlobalState = globalState
		_ = _lhs127
		var _lhs128 _dafny.Array = _858_v142
		_ = _lhs128
		var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_858_v142), 0))
		_ = _lhs129
		var _lhs130 *GlobalState = globalState
		_ = _lhs130
		_lhs126.F6 = _rhs136
		_lhs127.F16 = _rhs137
		(_lhs128).ArraySet1(_rhs138, (_lhs129).Int())
		_lhs130.F24 = _rhs139
	}
}
func (_this *C4) M9(p0 bool, p1 _dafny.Array, p2 _dafny.Map, globalState *GlobalState) (bool, bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 bool = false
		_ = r3
		var _859_v0 _dafny.Array
		_ = _859_v0
		var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
		_ = _nw120
		_859_v0 = _nw120
		var _860_v1 _dafny.Int
		_ = _860_v1
		_860_v1 = _dafny.IntOfInt64(-964)
		var _861_v2 _dafny.Map
		_ = _861_v2
		_861_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_860_v1, _860_v1)
		var _862_v3 _dafny.Map
		_ = _862_v3
		_862_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_860_v1, _861_v2)
		var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_859_v0), 0))
		_ = _index118
		(_859_v0).ArraySet1(_862_v3, (_index118).Int())
		var _863_v4 _dafny.Sequence
		_ = _863_v4
		_863_v4 = _dafny.SeqOf(_860_v1, _860_v1)
		var _864_v5 _dafny.MultiSet
		_ = _864_v5
		_864_v5 = _dafny.MultiSetOf((_this).Fm2(p0, _dafny.IntOfUint32((_863_v4).Cardinality()), globalState), (_this).F26())
		var _865_v6 _dafny.Sequence
		_ = _865_v6
		_865_v6 = _dafny.SeqOf((_864_v5).IsSubsetOf(_864_v5))
		var _866_v7 _dafny.Array
		_ = _866_v7
		var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
		_ = _nw121
		_866_v7 = _nw121
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_866_v7), 0))
		_ = _index119
		(_866_v7).ArraySet1(p1, (_index119).Int())
		var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_859_v0), 0))
		_ = _index120
		var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_866_v7), 0))
		_ = _index121
		var _rhs140 _dafny.Map = _862_v3
		_ = _rhs140
		var _rhs141 _dafny.Sequence = _865_v6
		_ = _rhs141
		var _rhs142 _dafny.Array = p1
		_ = _rhs142
		var _lhs131 _dafny.Array = _859_v0
		_ = _lhs131
		var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_859_v0), 0))
		_ = _lhs132
		var _lhs133 _dafny.Array = _866_v7
		_ = _lhs133
		var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_866_v7), 0))
		_ = _lhs134
		(_lhs131).ArraySet1(_rhs140, (_lhs132).Int())
		_865_v6 = _rhs141
		(_lhs133).ArraySet1(_rhs142, (_lhs134).Int())
		_860_v1 = _860_v1
		(globalState).F2 = p0
		r0 = (_this).F26()
		var _867_v8 _dafny.Array
		_ = _867_v8
		var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw122
		_867_v8 = _nw122
		var _868_v9 _dafny.Array
		_ = _868_v9
		var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw123
		_868_v9 = _nw123
		var _869_v10 _dafny.Map
		_ = _869_v10
		_869_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_860_v1), _868_v9)
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_867_v8), 0))
		_ = _index122
		(_867_v8).ArraySet1((func() _dafny.Array {
			if (_869_v10).Contains(_860_v1) {
				return (_869_v10).Get(_860_v1).(_dafny.Array)
			}
			return _868_v9
		})(), (_index122).Int())
		var _870_v11 _dafny.Array
		_ = _870_v11
		var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw124
		_870_v11 = _nw124
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((p1), 0))
		_ = _index123
		(p1).ArraySet1((_this).F27(), (_index123).Int())
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_867_v8), 0))
		_ = _index124
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((p1), 0))
		_ = _index125
		var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_866_v7), 0))
		_ = _index126
		var _rhs143 _dafny.Array = _868_v9
		_ = _rhs143
		var _rhs144 _dafny.Array = _870_v11
		_ = _rhs144
		var _rhs145 bool = p0
		_ = _rhs145
		var _rhs146 _dafny.Array = p1
		_ = _rhs146
		var _lhs135 _dafny.Array = _867_v8
		_ = _lhs135
		var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_867_v8), 0))
		_ = _lhs136
		var _lhs137 _dafny.Array = p1
		_ = _lhs137
		var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((p1), 0))
		_ = _lhs138
		var _lhs139 _dafny.Array = _866_v7
		_ = _lhs139
		var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_866_v7), 0))
		_ = _lhs140
		(_lhs135).ArraySet1(_rhs143, (_lhs136).Int())
		_870_v11 = _rhs144
		(_lhs137).ArraySet1(_rhs145, (_lhs138).Int())
		(_lhs139).ArraySet1(_rhs146, (_lhs140).Int())
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((p1), 0))
		_ = _index127
		(p1).ArraySet1(!((_this).F27()), (_index127).Int())
		var _871_v12 _dafny.CodePoint
		_ = _871_v12
		_871_v12 = _dafny.CodePoint('e')
		var _872_v13 _dafny.Map
		_ = _872_v13
		_872_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_860_v1, _871_v12)
		var _873_v14 _dafny.Set
		_ = _873_v14
		_873_v14 = _dafny.SetOf(_872_v13)
		r0 = !(((_873_v14).Difference(func() _dafny.Set {
			var _coll53 = _dafny.NewBuilder()
			_ = _coll53
			for _iter53 := _dafny.Iterate((_873_v14).Elements()); ; {
				_compr_53, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _874_v15 _dafny.Map
				_874_v15 = interface{}(_compr_53).(_dafny.Map)
				if (_873_v14).Contains(_874_v15) {
					_coll53.Add(_874_v15)
				}
			}
			return _coll53.ToSet()
		}())).IsSubsetOf(_873_v14))
		r1 = !(!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(530))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}((func(_875_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_876_i0 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_875_v1)
			}
		})(_860_v1))), _863_v4))
		var _877_v16 D9
		_ = _877_v16
		_877_v16 = Companion_D9_.Create_DC22_((_this).F27())
		r2 = (func() _dafny.Int {
			if (_877_v16).Dtor_cf42() {
				return (_dafny.Zero).Minus(_860_v1)
			}
			return _dafny.IntOfUint32((Companion_Default___.Fm15(globalState)).Cardinality())
		})()
		var _878_v17 D2
		_ = _878_v17
		_878_v17 = Companion_D2_.Create_DC8_()
		var _pat_let_tv6 = p1
		_ = _pat_let_tv6
		var _pat_let_tv7 = p1
		_ = _pat_let_tv7
		var _pat_let_tv8 = p1
		_ = _pat_let_tv8
		var _pat_let_tv9 = p1
		_ = _pat_let_tv9
		var _pat_let_tv10 = p2
		_ = _pat_let_tv10
		var _pat_let_tv11 = _865_v6
		_ = _pat_let_tv11
		var _pat_let_tv12 = _865_v6
		_ = _pat_let_tv12
		var _pat_let_tv13 = p2
		_ = _pat_let_tv13
		var _pat_let_tv14 = _865_v6
		_ = _pat_let_tv14
		var _pat_let_tv15 = _865_v6
		_ = _pat_let_tv15
		var _pat_let_tv16 = _860_v1
		_ = _pat_let_tv16
		var _pat_let_tv17 = _860_v1
		_ = _pat_let_tv17
		var _pat_let_tv18 = _860_v1
		_ = _pat_let_tv18
		r3 = func(_source15 D2) bool {
			if _source15.Is_DC7() {
				var _879___mcc_h0 bool = _source15.Get_().(D2_DC7).Cf14
				_ = _879___mcc_h0
				var _880___mcc_h1 bool = _source15.Get_().(D2_DC7).Cf15
				_ = _880___mcc_h1
				var _881___mcc_h2 _dafny.Int = _source15.Get_().(D2_DC7).Cf16
				_ = _881___mcc_h2
				var _882___mcc_h3 _dafny.Int = _source15.Get_().(D2_DC7).Cf17
				_ = _882___mcc_h3
				var _883_cf17 _dafny.Int = _882___mcc_h3
				_ = _883_cf17
				var _884_cf16 _dafny.Int = _881___mcc_h2
				_ = _884_cf16
				var _885_cf15 bool = _880___mcc_h1
				_ = _885_cf15
				var _886_cf14 bool = _879___mcc_h0
				_ = _886_cf14
				return (_pat_let_tv7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_pat_let_tv6), 0))).Int()).(bool)
			} else if _source15.Is_DC8() {
				return (_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(427), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(bool)
			} else if _source15.Is_DC9() {
				var _887___mcc_h4 _dafny.Int = _source15.Get_().(D2_DC9).Cf18
				_ = _887___mcc_h4
				var _888___mcc_h5 _dafny.Int = _source15.Get_().(D2_DC9).Cf19
				_ = _888___mcc_h5
				var _889___mcc_h6 bool = _source15.Get_().(D2_DC9).Cf20
				_ = _889___mcc_h6
				var _890_cf20 bool = _889___mcc_h6
				_ = _890_cf20
				var _891_cf19 _dafny.Int = _888___mcc_h5
				_ = _891_cf19
				var _892_cf18 _dafny.Int = _887___mcc_h4
				_ = _892_cf18
				if (_pat_let_tv10).Contains((_pat_let_tv11).Select((Companion_Default___.SafeIndex(_892_cf18, _dafny.IntOfUint32((_pat_let_tv12).Cardinality()))).Uint32()).(bool)) {
					return (_pat_let_tv13).Get((_pat_let_tv14).Select((Companion_Default___.SafeIndex(_892_cf18, _dafny.IntOfUint32((_pat_let_tv15).Cardinality()))).Uint32()).(bool)).(bool)
				} else {
					return (Companion_D5_.Create_DC14_(_dafny.IntOfInt64(752), _dafny.IntOfInt64(463), _890_cf20)).Dtor_cf27()
				}
			} else {
				var _893___mcc_h7 _dafny.Map = _source15.Get_().(D2_DC6).Cf13
				_ = _893___mcc_h7
				var _894_cf13 _dafny.Map = _893___mcc_h7
				_ = _894_cf13
				return (_dafny.SetOf(_pat_let_tv16, _pat_let_tv17)).Equals(func() _dafny.Set {
					var _coll54 = _dafny.NewBuilder()
					_ = _coll54
					for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-111), _dafny.IntOfInt64(-734))); ; {
						_compr_54, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _895_v18 _dafny.Int
						_895_v18 = interface{}(_compr_54).(_dafny.Int)
						if ((_dafny.IntOfInt64(-111)).Cmp(_895_v18) <= 0) && ((_895_v18).Cmp(_dafny.IntOfInt64(-734)) < 0) {
							_coll54.Add((_895_v18).Plus(_pat_let_tv18))
						}
					}
					return _coll54.ToSet()
				}())
			}
		}(_878_v17)
		return r0, r1, r2, r3
	}
}
func (_this *C4) M10(p0 bool, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r3
		var _896_v0 _dafny.Map
		_ = _896_v0
		_896_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p1), !(p0))
		var _897_v1 _dafny.Sequence
		_ = _897_v1
		_897_v1 = _dafny.SeqOf(p0, (_this).Fm2((_this).F27(), _dafny.IntOfInt64(771), globalState), p0)
		var _hi10 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_897_v1, _897_v1)).Cardinality())
		_ = _hi10
		for _898_i0 := (_896_v0).Cardinality(); _898_i0.Cmp(_hi10) < 0; _898_i0 = _898_i0.Plus(_dafny.One) {
			var _899_v2 _dafny.Array
			_ = _899_v2
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_27
			var _nw125 _dafny.Array
			_ = _nw125
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw125 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Sequence = (func(_900_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_901_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_900_v1, _900_v1, _900_v1)
					}
				})(_897_v1)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw125 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw125).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw125).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_899_v2 = _nw125
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_899_v2), 0))
			_ = _index128
			(_899_v2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_897_v1), _dafny.SeqOf(_897_v1)), (_index128).Int())
			var _902_v3 _dafny.Sequence
			_ = _902_v3
			_902_v3 = _dafny.SeqOf(_897_v1, _897_v1)
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_899_v2), 0))
			_ = _index129
			(_899_v2).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_902_v3, _dafny.SeqOf(_dafny.SeqOf(p0), _897_v1)), (_index129).Int())
			var _903_v4 _dafny.Sequence
			_ = _903_v4
			_903_v4 = _dafny.UnicodeSeqOfUtf8Bytes("usmgaxt")
			var _904_v5 _dafny.MultiSet
			_ = _904_v5
			_904_v5 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}((func(_905_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_906_i2 _dafny.Int) _dafny.Int {
					return _905_p1
				}
			})(p1)))).Cardinality()), _dafny.IntOfInt64(123), p1, _898_i0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(886))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_907_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			}))).Cardinality()))
			var _908_v6 _dafny.Set
			_ = _908_v6
			_908_v6 = _dafny.SetOf(p1, (_904_v5).Cardinality())
			var _909_v7 _dafny.Map
			_ = _909_v7
			_909_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_908_v6, (_this).F27())
			var _910_v8 _dafny.Sequence
			_ = _910_v8
			_910_v8 = _dafny.SeqOf(_dafny.IntOfInt64(-252), (_908_v6).Cardinality())
			var _911_v9 _dafny.MultiSet
			_ = _911_v9
			_911_v9 = _dafny.MultiSetOf((_this).F26())
			var _912_v10 _dafny.Array
			_ = _912_v10
			var _nwElement0_24 _dafny.Int = (Companion_Default___.Fm29(false, _898_i0, globalState)).Cardinality()
			_ = _nwElement0_24
			var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(20))
			_ = _nw126
			(_nw126).ArraySet1(_nwElement0_24, 0)
			(_nw126).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm22(_dafny.IntOfUint32((_903_v4).Cardinality()), (_897_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_897_v1).Cardinality()))).Uint32()).(bool), globalState)).Cardinality()), _898_i0), 1)
			(_nw126).ArraySet1((_908_v6).Cardinality(), 2)
			(_nw126).ArraySet1(((_908_v6).Difference(_908_v6)).Cardinality(), 3)
			(_nw126).ArraySet1(_898_i0, 4)
			(_nw126).ArraySet1((p1).Minus(p1), 5)
			(_nw126).ArraySet1((_dafny.Zero).Minus(((_909_v7).Merge(_909_v7)).Cardinality()), 6)
			(_nw126).ArraySet1(Companion_Default___.SafeDivisionInt(p1, _898_i0), 7)
			(_nw126).ArraySet1((_898_i0).Plus(p1), 8)
			(_nw126).ArraySet1(Companion_Default___.SafeDivisionInt((Companion_Default___.Fm30(_898_i0, p0, globalState)).Cardinality(), _dafny.IntOfUint32((_897_v1).Cardinality())), 9)
			(_nw126).ArraySet1(_898_i0, 10)
			(_nw126).ArraySet1((_this).Fm1(globalState), 11)
			(_nw126).ArraySet1(_898_i0, 12)
			(_nw126).ArraySet1(_898_i0, 13)
			(_nw126).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F27() {
					return p1
				}
				return _898_i0
			})()), 14)
			(_nw126).ArraySet1(p1, 15)
			(_nw126).ArraySet1(Companion_Default___.SafeModuloInt(_898_i0, (_910_v8).Select((Companion_Default___.SafeIndex(_898_i0, _dafny.IntOfUint32((_910_v8).Cardinality()))).Uint32()).(_dafny.Int)), 16)
			(_nw126).ArraySet1(p1, 17)
			(_nw126).ArraySet1(_dafny.IntOfUint32((_903_v4).Cardinality()), 18)
			(_nw126).ArraySet1((_911_v9).Cardinality(), 19)
			_912_v10 = _nw126
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_912_v10), 0))
			_ = _index130
			(_912_v10).ArraySet1(p1, (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(625), _dafny.ArrayLen((_912_v10), 0))
			_ = _index131
			(_912_v10).ArraySet1((p1).Minus(_898_i0), (_index131).Int())
			var _913_v11 *C1
			_ = _913_v11
			var _nw127 *C1 = New_C1_()
			_ = _nw127
			_nw127.Ctor__((_this).Fm1(globalState), _dafny.IntOfUint32((_897_v1).Cardinality()), (_this).F26(), p0)
			_913_v11 = _nw127
			(_913_v11).F37 = ((_911_v9).Difference(_911_v9)).Cardinality()
		}
		var _914_v12 _dafny.MultiSet
		_ = _914_v12
		_914_v12 = _dafny.MultiSetOf((_this).F26(), (_this).F27(), (_this).Fm2((_this).F27(), p1, globalState))
		(globalState).F22 = (_897_v1).Select((Companion_Default___.SafeIndex((_914_v12).Cardinality(), _dafny.IntOfUint32((_897_v1).Cardinality()))).Uint32()).(bool)
		var _915_v13 _dafny.Array
		_ = _915_v13
		var _len0_28 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_28
		var _nw128 _dafny.Array
		_ = _nw128
		if _len0_28.Cmp(_dafny.Zero) == 0 {
			_nw128 = _dafny.NewArray(_len0_28)
		} else {
			var _init28 func(_dafny.Int) _dafny.Sequence = (func(_916_p1 _dafny.Int, _917_p0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_918_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_dafny.Zero).Minus((Companion_D5_.Create_DC14_(_916_p1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-342), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}(func(_919_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					}))).Cardinality()))).Cardinality()), _917_p0)).Dtor_cf26()))
				}
			})(p1, p0)
			_ = _init28
			var _element0_28 = _init28(_dafny.Zero)
			_ = _element0_28
			_nw128 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
			(_nw128).ArraySet1(_element0_28, 0)
			var _nativeLen0_28 = (_len0_28).Int()
			_ = _nativeLen0_28
			for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
				(_nw128).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
			}
		}
		_915_v13 = _nw128
		_915_v13 = _915_v13
		var _920_v14 _dafny.Sequence
		_ = _920_v14
		_920_v14 = _dafny.UnicodeSeqOfUtf8Bytes("vgkdg")
		var _921_v16 _dafny.Map
		_ = _921_v16
		_921_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(602), _dafny.SetOf((_this).F26()))
		if !((func() _dafny.Map {
			if p0 {
				return func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-648), _dafny.IntOfInt64(519))); ; {
						_compr_55, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _922_v15 _dafny.Int
						_922_v15 = interface{}(_compr_55).(_dafny.Int)
						if ((_dafny.IntOfInt64(-648)).Cmp(_922_v15) <= 0) && ((_922_v15).Cmp(_dafny.IntOfInt64(519)) < 0) {
							_coll55.Add((_922_v15).Times(p1), _dafny.SetOf((_this).F26()))
						}
					}
					return _coll55.ToMap()
				}()
			}
			return _921_v16
		})()).Contains(_dafny.IntOfUint32((_920_v14).Cardinality())) {
			var _923_v17 _dafny.Array
			_ = _923_v17
			var _nw129 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw129
			_923_v17 = _nw129
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_923_v17), 0))
			_ = _index132
			(_923_v17).ArraySet1((_this).F26(), (_index132).Int())
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_923_v17), 0))
			_ = _index133
			(_923_v17).ArraySet1(!(true), (_index133).Int())
			var _924_v18 _dafny.Sequence
			_ = _924_v18
			_924_v18 = _dafny.SeqOf(_897_v1)
			_924_v18 = _924_v18
			var _925_v19 *C0
			_ = _925_v19
			var _nw130 *C0 = New_C0_()
			_ = _nw130
			_nw130.Ctor__((_this).F27())
			_925_v19 = _nw130
			var _926_v20 _dafny.Map
			_ = _926_v20
			_926_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ettqukfnq"), _925_v19)
			var _927_v21 _dafny.Array
			_ = _927_v21
			var _nwElement0_25 *C0 = _925_v19
			_ = _nwElement0_25
			var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(16))
			_ = _nw131
			(_nw131).ArraySet1(_nwElement0_25, 0)
			(_nw131).ArraySet1(_925_v19, 1)
			(_nw131).ArraySet1(_925_v19, 2)
			(_nw131).ArraySet1(_925_v19, 3)
			(_nw131).ArraySet1(_925_v19, 4)
			(_nw131).ArraySet1(_925_v19, 5)
			(_nw131).ArraySet1(_925_v19, 6)
			(_nw131).ArraySet1(_925_v19, 7)
			(_nw131).ArraySet1(_925_v19, 8)
			(_nw131).ArraySet1((func() *C0 {
				if (_923_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_923_v17), 0))).Int()).(bool) {
					return _925_v19
				}
				return _925_v19
			})(), 9)
			(_nw131).ArraySet1(_925_v19, 10)
			(_nw131).ArraySet1(_925_v19, 11)
			(_nw131).ArraySet1((func() *C0 {
				if (_926_v20).Contains(_920_v14) {
					return (_926_v20).Get(_920_v14).(*C0)
				}
				return _925_v19
			})(), 12)
			(_nw131).ArraySet1(_925_v19, 13)
			(_nw131).ArraySet1(_925_v19, 14)
			(_nw131).ArraySet1(_925_v19, 15)
			_927_v21 = _nw131
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_927_v21), 0))
			_ = _index134
			(_927_v21).ArraySet1(_925_v19, (_index134).Int())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_927_v21), 0))
			_ = _index135
			(_927_v21).ArraySet1(_925_v19, (_index135).Int())
			var _928_v22 _dafny.Array
			_ = _928_v22
			var _nw132 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(4))
			_ = _nw132
			_928_v22 = _nw132
			var _929_v23 _dafny.Sequence
			_ = _929_v23
			_929_v23 = _dafny.SeqOf(p1, p1)
			var _930_v24 _dafny.MultiSet
			_ = _930_v24
			_930_v24 = _dafny.MultiSetOf(_929_v23, _929_v23, _929_v23)
			var _931_v25 *C3
			_ = _931_v25
			var _nw133 *C3 = New_C3_()
			_ = _nw133
			_nw133.Ctor__(_930_v24, (_923_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen((_923_v17), 0))).Int()).(bool), (_this).F26())
			_931_v25 = _nw133
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_928_v22), 0))
			_ = _index136
			(_928_v22).ArraySet1(_931_v25, (_index136).Int())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_928_v22), 0))
			_ = _index137
			(_928_v22).ArraySet1(_931_v25, (_index137).Int())
			var _932_v26 _dafny.Map
			_ = _932_v26
			_932_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			_932_v26 = (_932_v26).Update(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_897_v1), _924_v18), (_this).F26())
		} else {
			var _933_v27 _dafny.Map
			_ = _933_v27
			_933_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F27()), p1)
			var _934_v28 _dafny.Map
			_ = _934_v28
			_934_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_933_v27).Cardinality(), p1)
			var _935_v29 _dafny.Sequence
			_ = _935_v29
			_935_v29 = _dafny.SeqOf(_915_v13, _915_v13)
			var _rhs147 _dafny.Int = p1
			_ = _rhs147
			var _rhs148 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_897_v1, _dafny.Companion_Sequence_.Concatenate(_897_v1, _897_v1))).Cardinality())
			_ = _rhs148
			var _rhs149 bool = (_this).F26()
			_ = _rhs149
			var _rhs150 _dafny.Int = (func() _dafny.Int {
				if false {
					return p1
				}
				return (func() _dafny.Int {
					if (_934_v28).Contains(p1) {
						return (_934_v28).Get(p1).(_dafny.Int)
					}
					return p1
				})()
			})()
			_ = _rhs150
			var _rhs151 _dafny.Array = (_935_v29).Select((Companion_Default___.SafeIndex((_this).Fm1(globalState), _dafny.IntOfUint32((_935_v29).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs151
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			r2 = _rhs147
			r2 = _rhs148
			_lhs141.F22 = _rhs149
			r2 = _rhs150
			_915_v13 = _rhs151
			var _936_v30 _dafny.Array
			_ = _936_v30
			var _nw134 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw134
			_936_v30 = _nw134
			var _937_v31 _dafny.CodePoint
			_ = _937_v31
			_937_v31 = _dafny.CodePoint('s')
			var _938_v32 _dafny.Map
			_ = _938_v32
			_938_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_936_v30, _937_v31)
			var _939_v33 _dafny.CodePoint
			_ = _939_v33
			var _940_v34 _dafny.Sequence
			_ = _940_v34
			var _941_v35 bool
			_ = _941_v35
			var _942_v36 _dafny.Map
			_ = _942_v36
			var _out29 _dafny.CodePoint
			_ = _out29
			var _out30 _dafny.Sequence
			_ = _out30
			var _out31 bool
			_ = _out31
			var _out32 _dafny.Map
			_ = _out32
			_out29, _out30, _out31, _out32 = Companion_Default___.M0((func() _dafny.CodePoint {
				if (_938_v32).Contains(_936_v30) {
					return (_938_v32).Get(_936_v30).(_dafny.CodePoint)
				}
				return _937_v31
			})(), globalState)
			_939_v33 = _out29
			_940_v34 = _out30
			_941_v35 = _out31
			_942_v36 = _out32
			var _943_v37 _dafny.Map
			_ = _943_v37
			_943_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _940_v34)
			var _944_v38 *C1
			_ = _944_v38
			var _nw135 *C1 = New_C1_()
			_ = _nw135
			_nw135.Ctor__(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_943_v37).Contains(p1) {
					return (_943_v37).Get(p1).(_dafny.Sequence)
				}
				return _940_v34
			})()).Cardinality()), (_this).F26(), (_this).F26())
			_944_v38 = _nw135
			var _945_v39 _dafny.CodePoint
			_ = _945_v39
			var _946_v40 _dafny.Sequence
			_ = _946_v40
			var _947_v41 bool
			_ = _947_v41
			var _948_v42 _dafny.Map
			_ = _948_v42
			var _out33 _dafny.CodePoint
			_ = _out33
			var _out34 _dafny.Sequence
			_ = _out34
			var _out35 bool
			_ = _out35
			var _out36 _dafny.Map
			_ = _out36
			_out33, _out34, _out35, _out36 = Companion_Default___.M0(_dafny.CodePoint('l'), globalState)
			_945_v39 = _out33
			_946_v40 = _out34
			_947_v41 = _out35
			_948_v42 = _out36
			if (Companion_D5_.Create_DC14_(_944_v38.F37, (_914_v12).Cardinality(), _947_v41)).Dtor_cf27() {
				var _949_v43 _dafny.Sequence
				_ = _949_v43
				_949_v43 = _dafny.SeqOf((_dafny.Zero).Minus(_944_v38.F37), _944_v38.F37)
				var _950_v44 _dafny.Sequence
				_ = _950_v44
				_950_v44 = _dafny.SeqOf(_949_v43, _949_v43)
				var _951_v45 D0
				_ = _951_v45
				_951_v45 = Companion_D0_.Create_DC0_((_this).F26())
				var _952_v46 _dafny.Set
				_ = _952_v46
				_952_v46 = _dafny.SetOf((_944_v38).Fm4(_951_v45, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uc")).Cardinality()), _dafny.IntOfUint32((_946_v40).Cardinality()), globalState))
				var _953_v47 T2
				_ = _953_v47
				var _nw136 *C2 = New_C2_()
				_ = _nw136
				_nw136.Ctor__(_950_v44, (func() bool {
					if _941_v35 {
						return p0
					}
					return _947_v41
				})(), (_952_v46).IsProperSubsetOf(_952_v46))
				_953_v47 = _nw136
				var _nw137 *C2 = New_C2_()
				_ = _nw137
				_nw137.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-120))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_954_v43 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_955_i6 _dafny.Int) _dafny.Sequence {
						return _954_v43
					}
				})(_949_v43))), _dafny.Companion_Sequence_.IsPrefixOf(_920_v14, _920_v14), (_953_v47).F26())
				_953_v47 = _nw137
				_936_v30 = _936_v30
				(globalState).F2 = (_953_v47).F26()
				var _956_v48 _dafny.Set
				_ = _956_v48
				_956_v48 = _dafny.SetOf((_this).Fm2(!((_953_v47).F27()), _944_v38.F37, globalState))
				var _957_v49 _dafny.Map
				_ = _957_v49
				_957_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v33, _944_v38.F37)
				var _958_v50 _dafny.Map
				_ = _958_v50
				_958_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_944_v38).F36(), _944_v38)
				var _959_v51 _dafny.Map
				_ = _959_v51
				_959_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_944_v38).F36(), _920_v14)
				var _960_v52 _dafny.Array
				_ = _960_v52
				var _nwElement0_26 _dafny.Int = (_956_v48).Cardinality()
				_ = _nwElement0_26
				var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(17))
				_ = _nw138
				(_nw138).ArraySet1(_nwElement0_26, 0)
				(_nw138).ArraySet1((_957_v49).Cardinality(), 1)
				(_nw138).ArraySet1((_944_v38).F36(), 2)
				(_nw138).ArraySet1((_958_v50).Cardinality(), 3)
				(_nw138).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_959_v51).Contains((_944_v38).F36()) {
						return (_959_v51).Get((_944_v38).F36()).(_dafny.Sequence)
					}
					return _920_v14
				})()).Cardinality()), 4)
				(_nw138).ArraySet1(_944_v38.F37, 5)
				(_nw138).ArraySet1((_944_v38).F36(), 6)
				(_nw138).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("umscmdh")).Cardinality()), 7)
				(_nw138).ArraySet1((_944_v38).F36(), 8)
				(_nw138).ArraySet1((_944_v38).F36(), 9)
				(_nw138).ArraySet1(_944_v38.F37, 10)
				(_nw138).ArraySet1((Companion_Default___.Fm31(p1, globalState)).Cardinality(), 11)
				(_nw138).ArraySet1((_dafny.Zero).Minus((_944_v38).F36()), 12)
				(_nw138).ArraySet1((_944_v38).F36(), 13)
				(_nw138).ArraySet1(_944_v38.F37, 14)
				(_nw138).ArraySet1(_944_v38.F37, 15)
				(_nw138).ArraySet1((_944_v38).F36(), 16)
				_960_v52 = _nw138
				var _961_v53 _dafny.Map
				_ = _961_v53
				_961_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v52, _945_v39)
				_961_v53 = (_961_v53).Update(_960_v52, _939_v33)
				r0 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if _947_v41 {
						return (_944_v38).F36()
					}
					return _dafny.IntOfUint32(((_953_v47).Fm3(_920_v14, _897_v1, (_944_v38).F36(), globalState)).Cardinality())
				})(), (_944_v38).F36())
			} else {
				var _962_v54 _dafny.Array
				_ = _962_v54
				var _nwElement0_27 _dafny.Int = Companion_Default___.SafeDivisionInt((_944_v38).F36(), _944_v38.F37)
				_ = _nwElement0_27
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(7))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_27, 0)
				(_nw139).ArraySet1((_944_v38).F36(), 1)
				(_nw139).ArraySet1(_944_v38.F37, 2)
				(_nw139).ArraySet1(p1, 3)
				(_nw139).ArraySet1(_944_v38.F37, 4)
				(_nw139).ArraySet1((_944_v38).F36(), 5)
				(_nw139).ArraySet1((_944_v38.F37).Plus((_944_v38).F36()), 6)
				_962_v54 = _nw139
				var _963_v55 _dafny.Array
				_ = _963_v55
				var _nwElement0_28 _dafny.Array = _962_v54
				_ = _nwElement0_28
				var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(4))
				_ = _nw140
				(_nw140).ArraySet1(_nwElement0_28, 0)
				(_nw140).ArraySet1(_962_v54, 1)
				(_nw140).ArraySet1(_962_v54, 2)
				(_nw140).ArraySet1(_962_v54, 3)
				_963_v55 = _nw140
				var _rhs152 bool = _941_v35
				_ = _rhs152
				var _rhs153 _dafny.Array = _962_v54
				_ = _rhs153
				var _rhs154 _dafny.Int = _dafny.IntOfInt64(89)
				_ = _rhs154
				var _rhs155 _dafny.Array = _963_v55
				_ = _rhs155
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				var _lhs143 *C1 = _944_v38
				_ = _lhs143
				_lhs142.F2 = _rhs152
				_962_v54 = _rhs153
				_lhs143.F37 = _rhs154
				_963_v55 = _rhs155
				var _964_v56 D0
				_ = _964_v56
				_964_v56 = Companion_D0_.Create_DC0_((_this).F26())
				(globalState).F12 = (_944_v38).Fm4(_964_v56, (_dafny.Zero).Minus(((_944_v38).F36()).Times(_944_v38.F37)), Companion_Default___.SafeDivisionInt(_944_v38.F37, (_944_v38).F36()), globalState)
				(globalState).F22 = !(true)
				var _965_v57 _dafny.Array
				_ = _965_v57
				var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
				_ = _nw141
				_965_v57 = _nw141
				var _966_v58 _dafny.Array
				_ = _966_v58
				var _nwElement0_29 *C1 = _944_v38
				_ = _nwElement0_29
				var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(5))
				_ = _nw142
				(_nw142).ArraySet1(_nwElement0_29, 0)
				(_nw142).ArraySet1(_944_v38, 1)
				(_nw142).ArraySet1(_944_v38, 2)
				(_nw142).ArraySet1(_944_v38, 3)
				(_nw142).ArraySet1(_944_v38, 4)
				_966_v58 = _nw142
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_966_v58), 0))
				_ = _index138
				(_966_v58).ArraySet1(_944_v38, (_index138).Int())
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_966_v58), 0))
				_ = _index139
				var _rhs156 _dafny.Array = _965_v57
				_ = _rhs156
				var _rhs157 bool = (_this).F27()
				_ = _rhs157
				var _rhs158 _dafny.Int = (_dafny.IntOfInt64(826)).Times((_944_v38).F36())
				_ = _rhs158
				var _rhs159 _dafny.Int = (_944_v38).F36()
				_ = _rhs159
				var _rhs160 *C1 = _944_v38
				_ = _rhs160
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 _dafny.Array = _966_v58
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_966_v58), 0))
				_ = _lhs148
				_965_v57 = _rhs156
				_lhs144.F22 = _rhs157
				_lhs145.F16 = _rhs158
				_lhs146.F16 = _rhs159
				(_lhs147).ArraySet1(_rhs160, (_lhs148).Int())
				(globalState).F22 = (_this).F26()
			}
		}
		var _967_v59 _dafny.Array
		_ = _967_v59
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_29
		var _nw143 _dafny.Array
		_ = _nw143
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw143 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) bool = func(_968_i7 _dafny.Int) bool {
				return (_this).F27()
			}
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw143 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw143).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw143).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_967_v59 = _nw143
		var _969_v60 _dafny.Map
		_ = _969_v60
		_969_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F27())
		var _970_v61 bool
		_ = _970_v61
		var _971_v62 bool
		_ = _971_v62
		var _972_v63 _dafny.Int
		_ = _972_v63
		var _973_v64 bool
		_ = _973_v64
		var _out37 bool
		_ = _out37
		var _out38 bool
		_ = _out38
		var _out39 _dafny.Int
		_ = _out39
		var _out40 bool
		_ = _out40
		_out37, _out38, _out39, _out40 = (_this).M9((_this).F27(), _967_v59, (func() _dafny.Map {
			if (_this).F26() {
				return _969_v60
			}
			return _969_v60
		})(), globalState)
		_970_v61 = _out37
		_971_v62 = _out38
		_972_v63 = _out39
		_973_v64 = _out40
		var _974_v65 _dafny.Sequence
		_ = _974_v65
		_974_v65 = _dafny.SeqOf(_897_v1)
		_974_v65 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_897_v1), _974_v65), _974_v65)
		r0 = p1
		r1 = _972_v63
		r2 = (_dafny.Zero).Minus(p1)
		r3 = Companion_Default___.Fm24(_920_v14, (_this).F26(), p1, globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C4) F33() _dafny.Map {
	{
		return _this._f33
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f26 bool
	_f27 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f26 = false
	_this._f27 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F26() bool {
	return _this._f26
}
func (_this *C5) F27() bool {
	return _this._f27
}
func (_this *C5) Ctor__(f26 bool, f27 bool) {
	{
		(_this)._f26 = f26
		(_this)._f27 = f27
	}
}
func (_this *C5) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-356)
	}
}
func (_this *C5) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source16 D0 = (func() D0 {
			if (_this).F27() {
				return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC1_((_this).F27(), (_this).F27(), _dafny.IntOfInt64(756), !(true)))
			}
			return Companion_D0_.Create_DC2_(Companion_D0_.Create_DC0_((_this).F26()))
		})()
		_ = _source16
		if _source16.Is_DC1() {
			var _975___mcc_h0 bool = _source16.Get_().(D0_DC1).Cf1
			_ = _975___mcc_h0
			var _976___mcc_h1 bool = _source16.Get_().(D0_DC1).Cf2
			_ = _976___mcc_h1
			var _977___mcc_h2 _dafny.Int = _source16.Get_().(D0_DC1).Cf3
			_ = _977___mcc_h2
			var _978___mcc_h3 bool = _source16.Get_().(D0_DC1).Cf4
			_ = _978___mcc_h3
			var _979_cf4 bool = _978___mcc_h3
			_ = _979_cf4
			var _980_cf3 _dafny.Int = _977___mcc_h2
			_ = _980_cf3
			var _981_cf2 bool = _976___mcc_h1
			_ = _981_cf2
			var _982_cf1 bool = _975___mcc_h0
			_ = _982_cf1
			return (_this).F27()
		} else if _source16.Is_DC0() {
			var _983___mcc_h4 bool = _source16.Get_().(D0_DC0).Cf0
			_ = _983___mcc_h4
			var _984_cf0 bool = _983___mcc_h4
			_ = _984_cf0
			return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wrscocei")).Cardinality())).Cmp(_dafny.IntOfInt64(468)) == 0
		} else {
			var _985___mcc_h5 D0 = _source16.Get_().(D0_DC2).Cf5
			_ = _985___mcc_h5
			var _986_cf5 D0 = _985___mcc_h5
			_ = _986_cf5
			return (_this).F26()
		}
	}
}
func (_this *C5) Fm5(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 D0, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		return _dafny.CodePoint('g')
	}
}
func (_this *C5) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _987_v0 *C0
		_ = _987_v0
		var _nw144 *C0 = New_C0_()
		_ = _nw144
		_nw144.Ctor__((_this).F26())
		_987_v0 = _nw144
		var _988_v2 _dafny.Map
		_ = _988_v2
		_988_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _989_v3 _dafny.Map
		_ = _989_v3
		_989_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
			var _coll56 = _dafny.NewMapBuilder()
			_ = _coll56
			for _iter56 := _dafny.Iterate((_988_v2).Keys().Elements()); ; {
				_compr_56, _ok56 := _iter56()
				if !_ok56 {
					break
				}
				var _990_v1 _dafny.Sequence
				_990_v1 = interface{}(_compr_56).(_dafny.Sequence)
				if (_988_v2).Contains(_990_v1) {
					_coll56.Add(_990_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fmtrj")).Cardinality()))
				}
			}
			return _coll56.ToMap()
		}(), p0)
		var _991_v4 _dafny.Int
		_ = _991_v4
		_991_v4 = _dafny.IntOfInt64(-494)
		var _992_v5 _dafny.Map
		_ = _992_v5
		_992_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _991_v4)
		var _993_v7 D2
		_ = _993_v7
		_993_v7 = Companion_D2_.Create_DC6_(func() _dafny.Map {
			var _coll57 = _dafny.NewMapBuilder()
			_ = _coll57
			for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(828), _dafny.IntOfInt64(982))); ; {
				_compr_57, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _994_v6 _dafny.Int
				_994_v6 = interface{}(_compr_57).(_dafny.Int)
				if ((_dafny.IntOfInt64(828)).Cmp(_994_v6) <= 0) && ((_994_v6).Cmp(_dafny.IntOfInt64(982)) < 0) {
					_coll57.Add((_994_v6).Minus(_991_v4), _991_v4)
				}
			}
			return _coll57.ToMap()
		}())
		_989_v3 = (_989_v3).Update(_992_v5, !((((_993_v7).Dtor_cf13()).Cardinality()).Cmp(_991_v4) > 0))
		(globalState).F6 = _dafny.UnicodeSeqOfUtf8Bytes("nbjowp")
		var _995_v8 *C0
		_ = _995_v8
		var _nw145 *C0 = New_C0_()
		_ = _nw145
		_nw145.Ctor__(!(p0) || ((_this).F27()))
		_995_v8 = _nw145
		var _996_v9 D2
		_ = _996_v9
		_996_v9 = Companion_D2_.Create_DC9_(_dafny.IntOfInt64(530), _dafny.IntOfUint32((p2).Cardinality()), (_this).F26())
		var _source17 D2 = _996_v9
		_ = _source17
		if _source17.Is_DC7() {
			var _997___mcc_h0 bool = _source17.Get_().(D2_DC7).Cf14
			_ = _997___mcc_h0
			var _998___mcc_h1 bool = _source17.Get_().(D2_DC7).Cf15
			_ = _998___mcc_h1
			var _999___mcc_h2 _dafny.Int = _source17.Get_().(D2_DC7).Cf16
			_ = _999___mcc_h2
			var _1000___mcc_h3 _dafny.Int = _source17.Get_().(D2_DC7).Cf17
			_ = _1000___mcc_h3
			var _1001_cf17 _dafny.Int = _1000___mcc_h3
			_ = _1001_cf17
			var _1002_cf16 _dafny.Int = _999___mcc_h2
			_ = _1002_cf16
			var _1003_cf15 bool = _998___mcc_h1
			_ = _1003_cf15
			var _1004_cf14 bool = _997___mcc_h0
			_ = _1004_cf14
			var _1005_v10 _dafny.Array
			_ = _1005_v10
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw146
			_1005_v10 = _nw146
			var _1006_v11 _dafny.Sequence
			_ = _1006_v11
			_1006_v11 = _dafny.UnicodeSeqOfUtf8Bytes("m")
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1005_v10), 0))
			_ = _index140
			(_1005_v10).ArraySet1(_1006_v11, (_index140).Int())
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1005_v10), 0))
			_ = _index141
			(_1005_v10).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1006_v11, _1006_v11), (_index141).Int())
			(globalState).F22 = (_this).Fm2(_1004_cf14, (_dafny.Zero).Minus(_991_v4), globalState)
			var _1007_v12 _dafny.Array
			_ = _1007_v12
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
			_ = _nw147
			_1007_v12 = _nw147
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1007_v12), 0))
			_ = _index142
			(_1007_v12).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf((_1005_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1005_v10), 0))).Int()).(_dafny.Sequence), (_1005_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_1005_v10), 0))).Int()).(_dafny.Sequence)), (_index142).Int())
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_1007_v12), 0))
			_ = _index143
			(_1007_v12).ArraySet1(false, (_index143).Int())
			(globalState).F2 = (_this).F27()
		} else if _source17.Is_DC8() {
			(globalState).F15 = (_987_v0).F32()
			var _1008_v13 D0
			_ = _1008_v13
			_1008_v13 = Companion_D0_.Create_DC1_((_987_v0).F32(), (_995_v8).F32(), _991_v4, (_995_v8).F32())
			var _1009_v14 _dafny.Map
			_ = _1009_v14
			_1009_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, (_1008_v13).Dtor_cf3())
			var _1010_v15 _dafny.MultiSet
			_ = _1010_v15
			_1010_v15 = _dafny.MultiSetOf(Companion_Default___.Fm8(_dafny.IntOfInt64(40), _991_v4, (_987_v0).F32(), _991_v4, globalState), _1009_v14, _1009_v14, _1009_v14, _1009_v14)
			var _1011_v16 _dafny.CodePoint
			_ = _1011_v16
			_1011_v16 = _dafny.CodePoint('o')
			var _1012_v17 _dafny.Map
			_ = _1012_v17
			_1012_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1011_v16)
			var _1013_v18 _dafny.Map
			_ = _1013_v18
			_1013_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1012_v17).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1011_v16))).Cardinality(), _dafny.MultiSetOf(_1009_v14))
			_1010_v15 = (func() _dafny.MultiSet {
				if (_1013_v18).Contains(_991_v4) {
					return (_1013_v18).Get(_991_v4).(_dafny.MultiSet)
				}
				return (_1010_v15).Union(_1010_v15)
			})()
			var _1014_v19 _dafny.Sequence
			_ = _1014_v19
			_1014_v19 = _dafny.UnicodeSeqOfUtf8Bytes("sksvlyg")
			(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1014_v19, (Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_1014_v19).Cardinality()))).Uint32(), _1011_v16), _1014_v19)
			var _1015_v20 _dafny.MultiSet
			_ = _1015_v20
			_1015_v20 = _dafny.MultiSetOf((_987_v0).F32(), (_987_v0).F32())
			_1015_v20 = Companion_Default___.Fm9(!((_995_v8).F32()), _991_v4, (_987_v0).F32(), globalState)
		} else if _source17.Is_DC9() {
			var _1016___mcc_h4 _dafny.Int = _source17.Get_().(D2_DC9).Cf18
			_ = _1016___mcc_h4
			var _1017___mcc_h5 _dafny.Int = _source17.Get_().(D2_DC9).Cf19
			_ = _1017___mcc_h5
			var _1018___mcc_h6 bool = _source17.Get_().(D2_DC9).Cf20
			_ = _1018___mcc_h6
			var _1019_cf20 bool = _1018___mcc_h6
			_ = _1019_cf20
			var _1020_cf19 _dafny.Int = _1017___mcc_h5
			_ = _1020_cf19
			var _1021_cf18 _dafny.Int = _1016___mcc_h4
			_ = _1021_cf18
			var _1022_v21 _dafny.Sequence
			_ = _1022_v21
			_1022_v21 = _dafny.UnicodeSeqOfUtf8Bytes("bfsobu")
			var _1023_v22 _dafny.MultiSet
			_ = _1023_v22
			_1023_v22 = _dafny.MultiSetOf(true)
			var _1024_v23 _dafny.Map
			_ = _1024_v23
			_1024_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _1020_cf19)
			var _1025_v24 _dafny.MultiSet
			_ = _1025_v24
			_1025_v24 = _dafny.MultiSetOf((_1023_v22).Cardinality(), (_1024_v23).Cardinality())
			var _1026_v25 _dafny.CodePoint
			_ = _1026_v25
			_1026_v25 = _dafny.CodePoint('l')
			var _1027_v26 _dafny.Map
			_ = _1027_v26
			_1027_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1022_v21, (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1025_v24).Contains((_this).Fm1(globalState)) {
					return (_1025_v24).Multiplicity((_this).Fm1(globalState))
				}
				return _1020_cf19
			})(), _dafny.IntOfUint32((_1022_v21).Cardinality()))).Uint32(), _1026_v25), _1026_v25)
			_1027_v26 = (_1027_v26).Update(_dafny.UnicodeSeqOfUtf8Bytes("ja"), _1026_v25)
			(globalState).F16 = _991_v4
			_1021_cf18 = (_this).Fm1(globalState)
			var _1028_v27 _dafny.Map
			_ = _1028_v27
			_1028_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1021_cf18, (_987_v0).F32())
			var _1029_v28 _dafny.Map
			_ = _1029_v28
			_1029_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1028_v27).Update(_1021_cf18, (_this).F26()), _991_v4)
			_1021_cf18 = (_1029_v28).Cardinality()
		} else {
			var _1030___mcc_h7 _dafny.Map = _source17.Get_().(D2_DC6).Cf13
			_ = _1030___mcc_h7
			var _1031_cf13 _dafny.Map = _1030___mcc_h7
			_ = _1031_cf13
			(globalState).F16 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(535))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}(func(_1032_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Cardinality())
			var _1033_v29 _dafny.Sequence
			_ = _1033_v29
			_1033_v29 = _dafny.UnicodeSeqOfUtf8Bytes("cc")
			var _1034_v30 _dafny.Map
			_ = _1034_v30
			_1034_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), _991_v4)
			var _1035_v31 _dafny.Sequence
			_ = _1035_v31
			_1035_v31 = _dafny.SeqOf(_991_v4, _991_v4, _991_v4)
			var _1036_v32 _dafny.Map
			_ = _1036_v32
			_1036_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, (_995_v8).F32())
			var _1037_v33 D1
			_ = _1037_v33
			_1037_v33 = Companion_D1_.Create_DC5_(_991_v4)
			var _1038_v34 _dafny.Array
			_ = _1038_v34
			var _nwElement0_30 _dafny.Int = _991_v4
			_ = _nwElement0_30
			var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(21))
			_ = _nw148
			(_nw148).ArraySet1(_nwElement0_30, 0)
			(_nw148).ArraySet1(_dafny.IntOfInt64(41), 1)
			(_nw148).ArraySet1(_991_v4, 2)
			(_nw148).ArraySet1(_991_v4, 3)
			(_nw148).ArraySet1(_991_v4, 4)
			(_nw148).ArraySet1(_991_v4, 5)
			(_nw148).ArraySet1(_991_v4, 6)
			(_nw148).ArraySet1(_dafny.IntOfInt64(887), 7)
			(_nw148).ArraySet1(_991_v4, 8)
			(_nw148).ArraySet1(_991_v4, 9)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_1033_v29).Cardinality()), 10)
			(_nw148).ArraySet1(_991_v4, 11)
			(_nw148).ArraySet1(_991_v4, 12)
			(_nw148).ArraySet1(_991_v4, 13)
			(_nw148).ArraySet1((func() _dafny.Int {
				if (_1034_v30).Contains((_987_v0).F32()) {
					return (_1034_v30).Get((_987_v0).F32()).(_dafny.Int)
				}
				return _991_v4
			})(), 14)
			(_nw148).ArraySet1(_991_v4, 15)
			(_nw148).ArraySet1((_1035_v31).Select((Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_1035_v31).Cardinality()))).Uint32()).(_dafny.Int), 16)
			(_nw148).ArraySet1(_dafny.IntOfUint32((_1033_v29).Cardinality()), 17)
			(_nw148).ArraySet1(_991_v4, 18)
			(_nw148).ArraySet1((_1036_v32).Cardinality(), 19)
			(_nw148).ArraySet1((_1037_v33).Dtor_cf12(), 20)
			_1038_v34 = _nw148
			var _1039_v35 D3
			_ = _1039_v35
			_1039_v35 = Companion_D3_.Create_DC10_(_1038_v34)
			var _1040_v36 _dafny.Array
			_ = _1040_v36
			var _nwElement0_31 _dafny.Array = _1038_v34
			_ = _nwElement0_31
			var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(25))
			_ = _nw149
			(_nw149).ArraySet1(_nwElement0_31, 0)
			(_nw149).ArraySet1(_1038_v34, 1)
			(_nw149).ArraySet1(_1038_v34, 2)
			(_nw149).ArraySet1(_1038_v34, 3)
			(_nw149).ArraySet1((_1039_v35).Dtor_cf21(), 4)
			(_nw149).ArraySet1(_1038_v34, 5)
			(_nw149).ArraySet1(_1038_v34, 6)
			(_nw149).ArraySet1(_1038_v34, 7)
			(_nw149).ArraySet1(_1038_v34, 8)
			(_nw149).ArraySet1(_1038_v34, 9)
			(_nw149).ArraySet1(_1038_v34, 10)
			(_nw149).ArraySet1(_1038_v34, 11)
			(_nw149).ArraySet1(_1038_v34, 12)
			(_nw149).ArraySet1(_1038_v34, 13)
			(_nw149).ArraySet1(_1038_v34, 14)
			(_nw149).ArraySet1(_1038_v34, 15)
			(_nw149).ArraySet1(_1038_v34, 16)
			(_nw149).ArraySet1(_1038_v34, 17)
			(_nw149).ArraySet1(_1038_v34, 18)
			(_nw149).ArraySet1(_1038_v34, 19)
			(_nw149).ArraySet1(_1038_v34, 20)
			(_nw149).ArraySet1(_1038_v34, 21)
			(_nw149).ArraySet1(_1038_v34, 22)
			(_nw149).ArraySet1(_1038_v34, 23)
			(_nw149).ArraySet1(_1038_v34, 24)
			_1040_v36 = _nw149
			_1040_v36 = _1040_v36
			var _1041_v37 _dafny.Map
			_ = _1041_v37
			_1041_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v29, _991_v4)
			var _1042_v38 _dafny.MultiSet
			_ = _1042_v38
			_1042_v38 = _dafny.MultiSetOf(p0, (_987_v0).Fm6(_dafny.CodePoint('j'), (_995_v8).F32(), globalState), (_987_v0).F32(), (p2).Select((Companion_Default___.SafeIndex(((_1041_v37).Update(_1033_v29, _991_v4)).Cardinality(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), (_995_v8).F32())
			var _1043_v39 _dafny.MultiSet
			_ = _1043_v39
			_1043_v39 = _dafny.MultiSetOf(((_1042_v38).Cardinality()).Plus(_991_v4), _991_v4)
			var _1044_v40 _dafny.Array
			_ = _1044_v40
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
			_ = _nw150
			_1044_v40 = _nw150
			var _1045_v41 D1
			_ = _1045_v41
			_1045_v41 = Companion_D1_.Create_DC4_(_1044_v40, _991_v4, _991_v4, _1033_v29, _1033_v29)
			var _1046_v42 _dafny.Map
			_ = _1046_v42
			_1046_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_987_v0).F32())
			var _pat_let_tv19 = _1046_v42
			_ = _pat_let_tv19
			var _1047_v43 _dafny.Map
			_ = _1047_v43
			_1047_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).F32(), (func(_pat_let13_0 D1) D1 {
				return func(_1048_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let14_0 _dafny.Int) D1 {
						return func(_1049_dt__update_hcf8_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_((_1048_dt__update__tmp_h0).Dtor_cf7(), _1049_dt__update_hcf8_h0, (_1048_dt__update__tmp_h0).Dtor_cf9(), (_1048_dt__update__tmp_h0).Dtor_cf10(), (_1048_dt__update__tmp_h0).Dtor_cf11())
						}(_pat_let14_0)
					}((_pat_let_tv19).Cardinality())
				}(_pat_let13_0)
			}(_1045_v41)).Dtor_cf10())
			var _1050_v44 D0
			_ = _1050_v44
			_1050_v44 = Companion_D0_.Create_DC1_((_987_v0).F32(), (_this).F26(), _991_v4, !((_this).F26()))
			_1043_v39 = (Companion_Default___.Fm10(_991_v4, (_1047_v43).Cardinality(), Companion_D0_.Create_DC2_(_1050_v44), globalState)).Union(_1043_v39)
			if (_this).F26() {
				(globalState).F16 = (((_dafny.MultiSetOf((_this).F26())).Union(_1042_v38)).Difference(_1042_v38)).Cardinality()
				(globalState).F16 = _991_v4
				var _1051_v45 _dafny.Array
				_ = _1051_v45
				var _nwElement0_32 bool = p0
				_ = _nwElement0_32
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(3))
				_ = _nw151
				(_nw151).ArraySet1(_nwElement0_32, 0)
				(_nw151).ArraySet1((_987_v0).F32(), 1)
				(_nw151).ArraySet1((_987_v0).F32(), 2)
				_1051_v45 = _nw151
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1051_v45), 0))
				_ = _index144
				(_1051_v45).ArraySet1((_this).F26(), (_index144).Int())
				var _1052_v46 D0
				_ = _1052_v46
				_1052_v46 = Companion_D0_.Create_DC0_((_987_v0).F32())
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1051_v45), 0))
				_ = _index145
				var _rhs161 bool = !(p0)
				_ = _rhs161
				var _rhs162 bool = (func() bool {
					if (_1052_v46).Dtor_cf0() {
						return (_987_v0).F32()
					}
					return !((_this).F27())
				})()
				_ = _rhs162
				var _lhs149 _dafny.Array = _1051_v45
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1051_v45), 0))
				_ = _lhs150
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				(_lhs149).ArraySet1(_rhs161, (_lhs150).Int())
				_lhs151.F24 = _rhs162
				_1038_v34 = _1038_v34
				(globalState).F12 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_991_v4), _991_v4)
			} else {
				var _1053_v47 _dafny.Array
				_ = _1053_v47
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw152
				_1053_v47 = _nw152
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_1053_v47), 0))
				_ = _index146
				(_1053_v47).ArraySet1((_1043_v39).IsProperSubsetOf(_1043_v39), (_index146).Int())
				var _1054_v48 _dafny.CodePoint
				_ = _1054_v48
				_1054_v48 = _dafny.CodePoint('d')
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_1053_v47), 0))
				_ = _index147
				(_1053_v47).ArraySet1((_995_v8).Fm6(_1054_v48, (_995_v8).F32(), globalState), (_index147).Int())
				var _1055_v49 D2
				_ = _1055_v49
				_1055_v49 = Companion_D2_.Create_DC8_()
				var _rhs163 D2 = _1055_v49
				_ = _rhs163
				var _rhs164 bool = true
				_ = _rhs164
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				_1055_v49 = _rhs163
				_lhs152.F24 = _rhs164
				var _1056_v50 D0
				_ = _1056_v50
				_1056_v50 = Companion_D0_.Create_DC0_(true)
				var _1057_v51 bool
				_ = _1057_v51
				var _out41 bool
				_ = _out41
				_out41 = (_this).M8(_1056_v50, (_this).F27(), (_dafny.MultiSetOf((_987_v0).F32())).Update((_987_v0).F32(), Companion_Default___.Abs(_991_v4)), _991_v4, globalState)
				_1057_v51 = _out41
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_1038_v34), 0))
				_ = _index148
				(_1038_v34).ArraySet1(_991_v4, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(616), _dafny.ArrayLen((_1038_v34), 0))
				_ = _index149
				(_1038_v34).ArraySet1(((func() _dafny.Int {
					if (_1053_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(358), _dafny.ArrayLen((_1053_v47), 0))).Int()).(bool) {
						return _991_v4
					}
					return _991_v4
				})()).Minus(_dafny.IntOfUint32((_1033_v29).Cardinality())), (_index149).Int())
				_1054_v48 = _1054_v48
			}
		}
		var _1058_v52 _dafny.CodePoint
		_ = _1058_v52
		_1058_v52 = _dafny.CodePoint('e')
		var _1059_v53 D1
		_ = _1059_v53
		_1059_v53 = Companion_D1_.Create_DC3_(Companion_Default___.Fm11((_995_v8).F32(), (_987_v0).F32(), _1058_v52, globalState))
		var _pat_let_tv20 = p2
		_ = _pat_let_tv20
		var _pat_let_tv21 = _991_v4
		_ = _pat_let_tv21
		var _pat_let_tv22 = p2
		_ = _pat_let_tv22
		var _pat_let_tv23 = _987_v0
		_ = _pat_let_tv23
		var _source18 D1 = func(_pat_let15_0 D1) D1 {
			return func(_1060_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let16_0 _dafny.Sequence) D1 {
					return func(_1061_dt__update_hcf6_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC3_(_1061_dt__update_hcf6_h0)
					}(_pat_let16_0)
				}(_dafny.Companion_Sequence_.Update(_pat_let_tv20, (Companion_Default___.SafeIndex(_pat_let_tv21, _dafny.IntOfUint32((_pat_let_tv22).Cardinality()))).Uint32(), (_pat_let_tv23).F32()))
			}(_pat_let15_0)
		}(_1059_v53)
		_ = _source18
		if _source18.Is_DC4() {
			var _1062___mcc_h8 _dafny.Array = _source18.Get_().(D1_DC4).Cf7
			_ = _1062___mcc_h8
			var _1063___mcc_h9 _dafny.Int = _source18.Get_().(D1_DC4).Cf8
			_ = _1063___mcc_h9
			var _1064___mcc_h10 _dafny.Int = _source18.Get_().(D1_DC4).Cf9
			_ = _1064___mcc_h10
			var _1065___mcc_h11 _dafny.Sequence = _source18.Get_().(D1_DC4).Cf10
			_ = _1065___mcc_h11
			var _1066___mcc_h12 _dafny.Sequence = _source18.Get_().(D1_DC4).Cf11
			_ = _1066___mcc_h12
			var _1067_cf11 _dafny.Sequence = _1066___mcc_h12
			_ = _1067_cf11
			var _1068_cf10 _dafny.Sequence = _1065___mcc_h11
			_ = _1068_cf10
			var _1069_cf9 _dafny.Int = _1064___mcc_h10
			_ = _1069_cf9
			var _1070_cf8 _dafny.Int = _1063___mcc_h9
			_ = _1070_cf8
			var _1071_cf7 _dafny.Array = _1062___mcc_h8
			_ = _1071_cf7
			var _1072_v54 _dafny.Array
			_ = _1072_v54
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw153
			_1072_v54 = _nw153
			var _1073_v55 _dafny.Map
			_ = _1073_v55
			_1073_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1072_v54, (_this).Fm1(globalState))
			var _1074_v56 _dafny.Map
			_ = _1074_v56
			_1074_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(true, (_987_v0).F32(), (_this).F27()))
			_1073_v55 = (_1073_v55).Update(_1072_v54, ((_1074_v56).Cardinality()).Plus(_1070_cf8))
			var _1075_v57 _dafny.Set
			_ = _1075_v57
			_1075_v57 = _dafny.SetOf(!((_995_v8).F32()), (_this).F27())
			var _1076_v58 _dafny.Set
			_ = _1076_v58
			_1076_v58 = _1075_v57
			(globalState).F2 = (_1075_v57).IsProperSubsetOf((_1076_v58))
			_991_v4 = _1069_cf9
			var _1077_v59 D2
			_ = _1077_v59
			_1077_v59 = Companion_D2_.Create_DC8_()
			var _source19 D2 = _1077_v59
			_ = _source19
			if _source19.Is_DC7() {
				var _1078___mcc_h15 bool = _source19.Get_().(D2_DC7).Cf14
				_ = _1078___mcc_h15
				var _1079___mcc_h16 bool = _source19.Get_().(D2_DC7).Cf15
				_ = _1079___mcc_h16
				var _1080___mcc_h17 _dafny.Int = _source19.Get_().(D2_DC7).Cf16
				_ = _1080___mcc_h17
				var _1081___mcc_h18 _dafny.Int = _source19.Get_().(D2_DC7).Cf17
				_ = _1081___mcc_h18
				var _1082_cf17 _dafny.Int = _1081___mcc_h18
				_ = _1082_cf17
				var _1083_cf16 _dafny.Int = _1080___mcc_h17
				_ = _1083_cf16
				var _1084_cf15 bool = _1079___mcc_h16
				_ = _1084_cf15
				var _1085_cf14 bool = _1078___mcc_h15
				_ = _1085_cf14
				var _1086_v60 _dafny.Array
				_ = _1086_v60
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_30
				var _nw154 _dafny.Array
				_ = _nw154
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw154 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Int = (func(_1087_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1088_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1088_i1, _dafny.IntOfUint32((_1087_p2).Cardinality()))
						}
					})(p2)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw154 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw154).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw154).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1086_v60 = _nw154
				var _1089_v61 _dafny.Map
				_ = _1089_v61
				_1089_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1086_v60, (_this).Fm1(globalState))
				_1089_v61 = (_1089_v61).Update(_1086_v60, (_1082_cf17).Minus(_dafny.IntOfUint32((_1068_cf10).Cardinality())))
				(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(45))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1090_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1091_i2 _dafny.Int) _dafny.CodePoint {
						return _1090_v52
					}
				})(_1058_v52))), _1067_cf11)
				_995_v8 = _995_v8
				(globalState).F6 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(830))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_1092_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1093_i3 _dafny.Int) _dafny.CodePoint {
						return _1092_v52
					}
				})(_1058_v52)))
			} else if _source19.Is_DC8() {
				var _1094_v62 _dafny.Sequence
				_ = _1094_v62
				_1094_v62 = _dafny.SeqOf(((_dafny.MultiSetOf(_1058_v52)).Update(_1058_v52, Companion_Default___.Abs(_dafny.IntOfInt64(204)))).Cardinality())
				var _1095_v63 _dafny.Array
				_ = _1095_v63
				var _nw155 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw155
				_1095_v63 = _nw155
				var _1096_v64 _dafny.Map
				_ = _1096_v64
				_1096_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1095_v63, _dafny.Companion_Sequence_.Update(_1094_v62, (Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_1094_v62).Cardinality()))).Uint32(), _1069_cf9))
				var _1097_v65 _dafny.Map
				_ = _1097_v65
				_1097_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1094_v62, (func() _dafny.Sequence {
					if (_1096_v64).Contains(_1095_v63) {
						return (_1096_v64).Get(_1095_v63).(_dafny.Sequence)
					}
					return _1094_v62
				})())
				(globalState).F15 = !(!(_1097_v65).Equals(_1097_v65))
				var _1098_v66 _dafny.Map
				_ = _1098_v66
				_1098_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1069_cf9, _1094_v62)
				_1098_v66 = (_1098_v66).Update(_991_v4, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}((func(_1099_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1100_i4 _dafny.Int) _dafny.Int {
						return _1099_cf9
					}
				})(_1069_cf9))), _1094_v62), (Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1101_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1102_i4 _dafny.Int) _dafny.Int {
						return _1101_cf9
					}
				})(_1069_cf9))), _1094_v62)).Cardinality()))).Uint32(), _991_v4))
				var _1103_v67 _dafny.Map
				_ = _1103_v67
				_1103_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, _1070_cf8)
				_1103_v67 = (_1103_v67).Update(_dafny.IntOfInt64(-944), _991_v4)
				(globalState).F2 = (_987_v0).F32()
			} else if _source19.Is_DC9() {
				var _1104___mcc_h19 _dafny.Int = _source19.Get_().(D2_DC9).Cf18
				_ = _1104___mcc_h19
				var _1105___mcc_h20 _dafny.Int = _source19.Get_().(D2_DC9).Cf19
				_ = _1105___mcc_h20
				var _1106___mcc_h21 bool = _source19.Get_().(D2_DC9).Cf20
				_ = _1106___mcc_h21
				var _1107_cf20 bool = _1106___mcc_h21
				_ = _1107_cf20
				var _1108_cf19 _dafny.Int = _1105___mcc_h20
				_ = _1108_cf19
				var _1109_cf18 _dafny.Int = _1104___mcc_h19
				_ = _1109_cf18
				(globalState).F2 = !(_1107_cf20) || ((_1109_cf18).Cmp(_1070_cf8) < 0)
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1071_cf7), 0))
				_ = _index150
				(_1071_cf7).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_995_v8).F32() {
						return _1058_v52
					}
					return _1058_v52
				})(), (_index150).Int())
				var _1110_v68 _dafny.Map
				_ = _1110_v68
				_1110_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1107_cf20, (_987_v0).F32())
				var _1111_v69 _dafny.MultiSet
				_ = _1111_v69
				_1111_v69 = _dafny.MultiSetOf(_991_v4, (_dafny.Zero).Minus((_1070_cf8).Minus(_991_v4)), (_dafny.Zero).Minus((func() _dafny.Int {
					if false {
						return (_1110_v68).Cardinality()
					}
					return _1109_cf18
				})()), (_1069_cf9).Times(_dafny.IntOfInt64(194)))
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1071_cf7), 0))
				_ = _index151
				var _rhs165 _dafny.CodePoint = _1058_v52
				_ = _rhs165
				var _rhs166 _dafny.MultiSet = ((_1111_v69).Intersection(_1111_v69)).Union((_1111_v69).Intersection(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, p0)).Cardinality())))
				_ = _rhs166
				var _rhs167 _dafny.Int = _1109_cf18
				_ = _rhs167
				var _lhs153 _dafny.Array = _1071_cf7
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1071_cf7), 0))
				_ = _lhs154
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				(_lhs153).ArraySet1CodePoint(_rhs165, (_lhs154).Int())
				_1111_v69 = _rhs166
				_lhs155.F12 = _rhs167
				var _1112_v70 _dafny.Array
				_ = _1112_v70
				var _len0_31 _dafny.Int = _dafny.One
				_ = _len0_31
				var _nw156 _dafny.Array
				_ = _nw156
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw156 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Int = (func(_1113_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1114_i5 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1114_i5, _1113_v4)
						}
					})(_991_v4)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw156 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw156).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw156).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1112_v70 = _nw156
				var _1115_v71 D3
				_ = _1115_v71
				_1115_v71 = Companion_D3_.Create_DC10_(_1112_v70)
				var _1116_v72 _dafny.Map
				_ = _1116_v72
				_1116_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1115_v71, (_this).Fm1(globalState))
				var _1117_v73 T1
				_ = _1117_v73
				var _nw157 *C4 = New_C4_()
				_ = _nw157
				_nw157.Ctor__(_1116_v72, (_995_v8).F32(), (_this).F27())
				_1117_v73 = _nw157
				var _1118_v74 _dafny.Map
				_ = _1118_v74
				_1118_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1070_cf8, _1117_v73)
				var _1119_v75 _dafny.Array
				_ = _1119_v75
				var _nwElement0_33 _dafny.Int = _1109_cf18
				_ = _nwElement0_33
				var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(15))
				_ = _nw158
				(_nw158).ArraySet1(_nwElement0_33, 0)
				(_nw158).ArraySet1((_this).Fm1(globalState), 1)
				(_nw158).ArraySet1((_1075_v57).Cardinality(), 2)
				(_nw158).ArraySet1(_1108_cf19, 3)
				(_nw158).ArraySet1(_dafny.IntOfInt64(507), 4)
				(_nw158).ArraySet1(_dafny.IntOfUint32((_1068_cf10).Cardinality()), 5)
				(_nw158).ArraySet1(_1108_cf19, 6)
				(_nw158).ArraySet1((_dafny.Zero).Minus(_1070_cf8), 7)
				(_nw158).ArraySet1(_1069_cf9, 8)
				(_nw158).ArraySet1((Companion_D0_.Create_DC1_((_987_v0).F32(), (_995_v8).F32(), _dafny.IntOfUint32((_1068_cf10).Cardinality()), (_995_v8).F32())).Dtor_cf3(), 9)
				(_nw158).ArraySet1(_1108_cf19, 10)
				(_nw158).ArraySet1((_1118_v74).Cardinality(), 11)
				(_nw158).ArraySet1(_1108_cf19, 12)
				(_nw158).ArraySet1((_dafny.Zero).Minus(_1070_cf8), 13)
				(_nw158).ArraySet1(_1108_cf19, 14)
				_1119_v75 = _nw158
				var _1120_v76 _dafny.MultiSet
				_ = _1120_v76
				_1120_v76 = _dafny.MultiSetOf(_1119_v75, _1119_v75, _1112_v70)
				(globalState).F16 = (_1120_v76).Cardinality()
				var _1121_v77 *C1
				_ = _1121_v77
				var _nw159 *C1 = New_C1_()
				_ = _nw159
				_nw159.Ctor__((_1070_cf8).Plus(_1069_cf9), (_1069_cf9).Minus(_1069_cf9), (_1070_cf8).Cmp(_991_v4) >= 0, (!((_995_v8).F32())) || (_1107_cf20))
				_1121_v77 = _nw159
			} else {
				var _1122___mcc_h22 _dafny.Map = _source19.Get_().(D2_DC6).Cf13
				_ = _1122___mcc_h22
				var _1123_cf13 _dafny.Map = _1122___mcc_h22
				_ = _1123_cf13
				var _1124_v78 *C1
				_ = _1124_v78
				var _nw160 *C1 = New_C1_()
				_ = _nw160
				_nw160.Ctor__(_dafny.IntOfUint32((p2).Cardinality()), _1069_cf9, p0, (_this).F27())
				_1124_v78 = _nw160
				var _1125_v79 _dafny.Set
				_ = _1125_v79
				_1125_v79 = _dafny.SetOf(_1124_v78, _1124_v78, _1124_v78)
				_1125_v79 = (_1125_v79).Difference(_1125_v79)
				(globalState).F22 = (true) && ((_995_v8).F32())
				(globalState).F12 = (_1124_v78).F36()
				(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_1067_cf11, (func() _dafny.Sequence {
					if (_987_v0).F32() {
						return _1068_cf10
					}
					return _1067_cf11
				})())
			}
		} else if _source18.Is_DC5() {
			var _1126___mcc_h13 _dafny.Int = _source18.Get_().(D1_DC5).Cf12
			_ = _1126___mcc_h13
			var _1127_cf12 _dafny.Int = _1126___mcc_h13
			_ = _1127_cf12
			var _1128_v80 _dafny.Map
			_ = _1128_v80
			_1128_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), _dafny.SeqOf((_987_v0).F32(), (_987_v0).F32()))
			var _1129_v81 D5
			_ = _1129_v81
			_1129_v81 = Companion_D5_.Create_DC13_((func() _dafny.Map {
				if true {
					return Companion_Default___.Fm32((_this).Fm1(globalState), _1127_cf12, globalState)
				}
				return (_1128_v80).Update(p0, p2)
			})())
			var _source20 D5 = _1129_v81
			_ = _source20
			if _source20.Is_DC14() {
				var _1130___mcc_h23 _dafny.Int = _source20.Get_().(D5_DC14).Cf25
				_ = _1130___mcc_h23
				var _1131___mcc_h24 _dafny.Int = _source20.Get_().(D5_DC14).Cf26
				_ = _1131___mcc_h24
				var _1132___mcc_h25 bool = _source20.Get_().(D5_DC14).Cf27
				_ = _1132___mcc_h25
				var _1133_cf27 bool = _1132___mcc_h25
				_ = _1133_cf27
				var _1134_cf26 _dafny.Int = _1131___mcc_h24
				_ = _1134_cf26
				var _1135_cf25 _dafny.Int = _1130___mcc_h23
				_ = _1135_cf25
				(globalState).F16 = (_dafny.Zero).Minus(_1134_cf26)
				var _rhs168 _dafny.Int = _1134_cf26
				_ = _rhs168
				var _rhs169 _dafny.Int = _dafny.IntOfInt64(294)
				_ = _rhs169
				var _rhs170 bool = (_987_v0).F32()
				_ = _rhs170
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				_lhs156.F16 = _rhs168
				_lhs157.F12 = _rhs169
				_lhs158.F24 = _rhs170
				var _1136_v82 _dafny.Array
				_ = _1136_v82
				var _len0_32 _dafny.Int = _dafny.One
				_ = _len0_32
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) D2 = func(_1137_i6 _dafny.Int) D2 {
						return Companion_D2_.Create_DC8_()
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw161 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw161).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw161).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1136_v82 = _nw161
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1136_v82), 0))
				_ = _index152
				(_1136_v82).ArraySet1(Companion_D2_.Create_DC8_(), (_index152).Int())
				var _1138_v83 _dafny.Sequence
				_ = _1138_v83
				_1138_v83 = _dafny.SeqOf(_1058_v52)
				var _1139_v84 _dafny.Set
				_ = _1139_v84
				_1139_v84 = _dafny.SetOf(p0, false)
				var _1140_v85 _dafny.Map
				_ = _1140_v85
				_1140_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1138_v83, (_1139_v84).Cardinality())
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_1136_v82), 0))
				_ = _index153
				(_1136_v82).ArraySet1(Companion_Default___.Fm33((_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F26() {
						return _dafny.IntOfInt64(743)
					}
					return _1135_cf25
				})()), ((_1140_v85).Merge((_1140_v85).Update(_1138_v83, _1127_cf12))).Cardinality(), globalState), (_index153).Int())
				(globalState).F16 = (_dafny.Zero).Minus(_1135_cf25)
			} else if _source20.Is_DC13() {
				var _1141___mcc_h26 _dafny.Map = _source20.Get_().(D5_DC13).Cf24
				_ = _1141___mcc_h26
				var _1142_cf24 _dafny.Map = _1141___mcc_h26
				_ = _1142_cf24
				var _1143_v86 _dafny.Array
				_ = _1143_v86
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw162
				_1143_v86 = _nw162
				var _1144_v87 _dafny.Map
				_ = _1144_v87
				_1144_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-46), _1127_cf12)
				var _1145_v88 _dafny.Sequence
				_ = _1145_v88
				_1145_v88 = _dafny.SeqOf(_1144_v87)
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1143_v86), 0))
				_ = _index154
				(_1143_v86).ArraySet1(_1145_v88, (_index154).Int())
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_1143_v86), 0))
				_ = _index155
				(_1143_v86).ArraySet1(_1145_v88, (_index155).Int())
				var _1146_v89 _dafny.Array
				_ = _1146_v89
				var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw163
				_1146_v89 = _nw163
				_1146_v89 = _1146_v89
				_1127_cf12 = _991_v4
				var _1147_v90 _dafny.Sequence
				_ = _1147_v90
				_1147_v90 = _dafny.UnicodeSeqOfUtf8Bytes("priuoxrge")
				var _1148_v91 _dafny.MultiSet
				_ = _1148_v91
				_1148_v91 = _dafny.MultiSetOf((p1).Select((Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Sequence), _1147_v90)
				_1148_v91 = _1148_v91
			} else {
				var _1149___mcc_h27 D5 = _source20.Get_().(D5_DC15).Cf28
				_ = _1149___mcc_h27
				var _1150_cf28 D5 = _1149___mcc_h27
				_ = _1150_cf28
				var _1151_v92 _dafny.MultiSet
				_ = _1151_v92
				_1151_v92 = _dafny.MultiSetOf(_dafny.SeqOf((_995_v8).F32()))
				var _1152_v93 _dafny.Array
				_ = _1152_v93
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(13)
				_ = _len0_33
				var _nw164 _dafny.Array
				_ = _nw164
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw164 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1153_v0 *C0) func(_dafny.Int) bool {
						return func(_1154_i7 _dafny.Int) bool {
							return (_1153_v0).F32()
						}
					})(_987_v0)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw164 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw164).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw164).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1152_v93 = _nw164
				var _1155_v94 _dafny.MultiSet
				_ = _1155_v94
				_1155_v94 = _dafny.MultiSetOf(_1152_v93, _1152_v93, _1152_v93)
				var _1156_v95 _dafny.Map
				_ = _1156_v95
				_1156_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _991_v4)
				var _1157_v96 _dafny.Sequence
				_ = _1157_v96
				_1157_v96 = _dafny.UnicodeSeqOfUtf8Bytes("dq")
				var _1158_v97 _dafny.Sequence
				_ = _1158_v97
				_1158_v97 = _dafny.SeqOf(_dafny.IntOfUint32((_1157_v96).Cardinality()))
				var _1159_v98 _dafny.Set
				_ = _1159_v98
				_1159_v98 = _dafny.SetOf((_987_v0).F32(), (_this).F26(), false, (_dafny.IntOfInt64(609)).Cmp(((_1151_v92).Update(p2, Companion_Default___.Abs((func() _dafny.Int {
					if (_1155_v94).Contains(_1152_v93) {
						return (_1155_v94).Multiplicity(_1152_v93)
					}
					return _991_v4
				})()))).Cardinality()) == 0, (_987_v0).Fm7(_1156_v95, _1158_v97, (_this).Fm2((_this).F26(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1157_v96).Cardinality())), globalState), globalState))
				var _rhs171 _dafny.Set = _1159_v98
				_ = _rhs171
				var _rhs172 D2 = _996_v9
				_ = _rhs172
				var _rhs173 _dafny.Sequence = _1158_v97
				_ = _rhs173
				_1159_v98 = _rhs171
				_996_v9 = _rhs172
				_1158_v97 = _rhs173
				var _1160_v99 _dafny.MultiSet
				_ = _1160_v99
				_1160_v99 = _dafny.MultiSetOf(_1158_v97, _1158_v97, _1158_v97)
				var _1161_v100 *C3
				_ = _1161_v100
				var _nw165 *C3 = New_C3_()
				_ = _nw165
				_nw165.Ctor__((_1160_v99).Update(_1158_v97, Companion_Default___.Abs(_991_v4)), (_991_v4).Cmp(_1127_cf12) <= 0, ((_995_v8).F32()) == (!((_995_v8).F32())))
				_1161_v100 = _nw165
				_1158_v97 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1158_v97, _dafny.SeqOf(_1127_cf12, _1127_cf12)), _1158_v97)
				(globalState).F22 = (_995_v8).F32()
			}
			var _source21 D1 = Companion_D1_.Create_DC3_(_dafny.SeqOf(true, (_this).F26(), (_987_v0).F32()))
			_ = _source21
			if _source21.Is_DC4() {
				var _1162___mcc_h28 _dafny.Array = _source21.Get_().(D1_DC4).Cf7
				_ = _1162___mcc_h28
				var _1163___mcc_h29 _dafny.Int = _source21.Get_().(D1_DC4).Cf8
				_ = _1163___mcc_h29
				var _1164___mcc_h30 _dafny.Int = _source21.Get_().(D1_DC4).Cf9
				_ = _1164___mcc_h30
				var _1165___mcc_h31 _dafny.Sequence = _source21.Get_().(D1_DC4).Cf10
				_ = _1165___mcc_h31
				var _1166___mcc_h32 _dafny.Sequence = _source21.Get_().(D1_DC4).Cf11
				_ = _1166___mcc_h32
				var _1167_cf11 _dafny.Sequence = _1166___mcc_h32
				_ = _1167_cf11
				var _1168_cf10 _dafny.Sequence = _1165___mcc_h31
				_ = _1168_cf10
				var _1169_cf9 _dafny.Int = _1164___mcc_h30
				_ = _1169_cf9
				var _1170_cf8 _dafny.Int = _1163___mcc_h29
				_ = _1170_cf8
				var _1171_cf7 _dafny.Array = _1162___mcc_h28
				_ = _1171_cf7
				var _1172_v101 _dafny.Array
				_ = _1172_v101
				var _nwElement0_34 _dafny.Int = _1170_cf8
				_ = _nwElement0_34
				var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(11))
				_ = _nw166
				(_nw166).ArraySet1(_nwElement0_34, 0)
				(_nw166).ArraySet1(_dafny.IntOfInt64(-155), 1)
				(_nw166).ArraySet1(_1169_cf9, 2)
				(_nw166).ArraySet1(_1170_cf8, 3)
				(_nw166).ArraySet1(_dafny.IntOfInt64(272), 4)
				(_nw166).ArraySet1(_dafny.IntOfInt64(842), 5)
				(_nw166).ArraySet1(_991_v4, 6)
				(_nw166).ArraySet1(_991_v4, 7)
				(_nw166).ArraySet1(_1169_cf9, 8)
				(_nw166).ArraySet1(_dafny.IntOfUint32((_1167_cf11).Cardinality()), 9)
				(_nw166).ArraySet1(_991_v4, 10)
				_1172_v101 = _nw166
				var _1173_v102 D3
				_ = _1173_v102
				_1173_v102 = Companion_D3_.Create_DC10_(_1172_v101)
				var _1174_v103 _dafny.Map
				_ = _1174_v103
				_1174_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1173_v102, _991_v4)
				var _1175_v104 *C4
				_ = _1175_v104
				var _nw167 *C4 = New_C4_()
				_ = _nw167
				_nw167.Ctor__(_1174_v103, (_987_v0).F32(), (_995_v8).F32())
				_1175_v104 = _nw167
				var _1176_v105 _dafny.Sequence
				_ = _1176_v105
				_1176_v105 = _dafny.SeqOf(_1175_v104)
				var _1177_v106 _dafny.Set
				_ = _1177_v106
				_1177_v106 = _dafny.SetOf((_987_v0).F32())
				var _1178_v107 _dafny.Set
				_ = _1178_v107
				_1178_v107 = _dafny.SetOf((_1177_v106).Cardinality())
				var _1179_v108 _dafny.MultiSet
				_ = _1179_v108
				_1179_v108 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(231))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_1180_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1181_i8 _dafny.Int) _dafny.Int {
						return _1180_v4
					}
				})(_991_v4))))
				var _1182_v109 *C3
				_ = _1182_v109
				var _nw168 *C3 = New_C3_()
				_ = _nw168
				_nw168.Ctor__(_1179_v108, (_995_v8).F32(), !(p0))
				_1182_v109 = _nw168
				var _1183_v110 _dafny.Map
				_ = _1183_v110
				_1183_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_1176_v105).Cardinality()), _1170_cf8)).IsSubsetOf(_1178_v107), _1182_v109)
				_1183_v110 = _1183_v110
				(globalState).F22 = !((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1167_cf11, _1167_cf11)).Cardinality())).Cmp((func() _dafny.Int {
					if (_995_v8).F32() {
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(77))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg69 _dafny.Int) interface{} {
								return coer69(arg69)
							}
						}((func(_1184_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1185_i9 _dafny.Int) _dafny.CodePoint {
								return _1184_v52
							}
						})(_1058_v52)))).Cardinality()))
					}
					return _991_v4
				})()) == 0)
				var _1186_v111 _dafny.Array
				_ = _1186_v111
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_34
				var _nw169 _dafny.Array
				_ = _nw169
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw169 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1187_p2 _dafny.Sequence, _1188_v0 *C0) func(_dafny.Int) _dafny.Sequence {
						return func(_1189_i10 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_1187_p2, _dafny.SeqOf((_1188_v0).F32(), (_this).F27()))
						}
					})(p2, _987_v0)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw169 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw169).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw169).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1186_v111 = _nw169
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_1186_v111), 0))
				_ = _index156
				(_1186_v111).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, p2), (_index156).Int())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_1186_v111), 0))
				_ = _index157
				(_1186_v111).ArraySet1(p2, (_index157).Int())
				var _1190_v112 _dafny.Map
				_ = _1190_v112
				_1190_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(120))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1191_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1192_i11 _dafny.Int) _dafny.CodePoint {
						return _1191_v52
					}
				})(_1058_v52))), _1170_cf8)
				_1190_v112 = (_1190_v112).Update(_dafny.Companion_Sequence_.Concatenate(_1167_cf11, _1168_cf10), _991_v4)
			} else if _source21.Is_DC5() {
				var _1193___mcc_h33 _dafny.Int = _source21.Get_().(D1_DC5).Cf12
				_ = _1193___mcc_h33
				var _1194_cf12 _dafny.Int = _1193___mcc_h33
				_ = _1194_cf12
				var _1195_v113 _dafny.Sequence
				_ = _1195_v113
				_1195_v113 = _dafny.SeqOf(_991_v4, _1194_cf12)
				var _1196_v114 _dafny.Sequence
				_ = _1196_v114
				_1196_v114 = _dafny.UnicodeSeqOfUtf8Bytes("nhtyej")
				var _1197_v115 _dafny.MultiSet
				_ = _1197_v115
				_1197_v115 = _dafny.MultiSetOf(_1196_v114, _1196_v114)
				var _1198_v116 _dafny.Sequence
				_ = _1198_v116
				_1198_v116 = _dafny.SeqOf(_993_v7, _993_v7, _993_v7)
				var _1199_v117 _dafny.Map
				_ = _1199_v117
				_1199_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1195_v113, (Companion_Default___.SafeIndex((_1197_v115).Cardinality(), _dafny.IntOfUint32((_1195_v113).Cardinality()))).Uint32(), _1194_cf12), _1198_v116)
				var _1200_v118 _dafny.Map
				_ = _1200_v118
				_1200_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v113, _1198_v116))
				var _1201_v119 D10
				_ = _1201_v119
				_1201_v119 = Companion_D10_.Create_DC23_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v113, _1198_v116))
				var _1202_v120 _dafny.Array
				_ = _1202_v120
				var _nwElement0_35 _dafny.Map = _1199_v117
				_ = _nwElement0_35
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(23))
				_ = _nw170
				(_nw170).ArraySet1(_nwElement0_35, 0)
				(_nw170).ArraySet1((func() _dafny.Map {
					if (_1200_v118).Contains(true) {
						return (_1200_v118).Get(true).(_dafny.Map)
					}
					return _1199_v117
				})(), 1)
				(_nw170).ArraySet1(_1199_v117, 2)
				(_nw170).ArraySet1(_1199_v117, 3)
				(_nw170).ArraySet1(_1199_v117, 4)
				(_nw170).ArraySet1(_1199_v117, 5)
				(_nw170).ArraySet1((_1199_v117).Update(Companion_Default___.Fm15(globalState), _dafny.SeqOf(_993_v7)), 6)
				(_nw170).ArraySet1((_1199_v117).Merge(_1199_v117), 7)
				(_nw170).ArraySet1((Companion_Default___.Fm34(globalState)).Update(_dafny.SeqOf(_991_v4), _1198_v116), 8)
				(_nw170).ArraySet1((_1199_v117).Merge(_1199_v117), 9)
				(_nw170).ArraySet1(_1199_v117, 10)
				(_nw170).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v113, _1198_v116), 11)
				(_nw170).ArraySet1(_1199_v117, 12)
				(_nw170).ArraySet1((_1199_v117).Update(_1195_v113, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(721))).Uint32(), func(coer71 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1203_v7 D2) func(_dafny.Int) D2 {
					return func(_1204_i12 _dafny.Int) D2 {
						return _1203_v7
					}
				})(_993_v7)))), 13)
				(_nw170).ArraySet1(_1199_v117, 14)
				(_nw170).ArraySet1((_1201_v119).Dtor_cf43(), 15)
				(_nw170).ArraySet1(_1199_v117, 16)
				(_nw170).ArraySet1(_1199_v117, 17)
				(_nw170).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1205_cf12 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1206_i13 _dafny.Int) _dafny.Int {
						return _1205_cf12
					}
				})(_1194_cf12))), _1198_v116), 18)
				(_nw170).ArraySet1((_1199_v117).Merge(_1199_v117), 19)
				(_nw170).ArraySet1((_1199_v117).Update(_dafny.SeqOf(_1194_cf12), _1198_v116), 20)
				(_nw170).ArraySet1(_1199_v117, 21)
				(_nw170).ArraySet1((_1199_v117).Merge(_1199_v117), 22)
				_1202_v120 = _nw170
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1202_v120), 0))
				_ = _index158
				(_1202_v120).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v113, _1198_v116), (_index158).Int())
				var _1207_v122 _dafny.Map
				_ = _1207_v122
				_1207_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1195_v113, _991_v4)
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1202_v120), 0))
				_ = _index159
				(_1202_v120).ArraySet1(((_1199_v117).Merge((func() _dafny.Map {
					var _coll58 = _dafny.NewMapBuilder()
					_ = _coll58
					for _iter58 := _dafny.Iterate((_1207_v122).Keys().Elements()); ; {
						_compr_58, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1208_v121 _dafny.Sequence
						_1208_v121 = interface{}(_compr_58).(_dafny.Sequence)
						if (_1207_v122).Contains(_1208_v121) {
							_coll58.Add(_1208_v121, _1198_v116)
						}
					}
					return _coll58.ToMap()
				}()).Update(_dafny.SeqOf(_1194_cf12), _1198_v116))).Merge(_1199_v117), (_index159).Int())
				var _rhs174 _dafny.Int = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ertaynatq")).Cardinality())).Minus(_dafny.IntOfInt64(893))
				_ = _rhs174
				var _rhs175 _dafny.Int = _1194_cf12
				_ = _rhs175
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				var _lhs160 *GlobalState = globalState
				_ = _lhs160
				_lhs159.F16 = _rhs174
				_lhs160.F16 = _rhs175
				(globalState).F15 = (_987_v0).F32()
				var _1209_v123 _dafny.Array
				_ = _1209_v123
				var _nw171 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
				_ = _nw171
				_1209_v123 = _nw171
				var _1210_v124 _dafny.MultiSet
				_ = _1210_v124
				_1210_v124 = _dafny.MultiSetOf(_1209_v123, _1209_v123)
				var _rhs176 _dafny.MultiSet = _1210_v124
				_ = _rhs176
				var _rhs177 _dafny.Int = (_996_v9).Dtor_cf19()
				_ = _rhs177
				var _lhs161 *GlobalState = globalState
				_ = _lhs161
				_1210_v124 = _rhs176
				_lhs161.F12 = _rhs177
			} else {
				var _1211___mcc_h34 _dafny.Sequence = _source21.Get_().(D1_DC3).Cf6
				_ = _1211___mcc_h34
				var _1212_cf6 _dafny.Sequence = _1211___mcc_h34
				_ = _1212_cf6
				var _1213_v125 _dafny.Sequence
				_ = _1213_v125
				_1213_v125 = _dafny.UnicodeSeqOfUtf8Bytes("wsdmh")
				var _1214_v126 _dafny.Sequence
				_ = _1214_v126
				_1214_v126 = _dafny.SeqOf((_this).Fm1(globalState))
				var _1215_v127 _dafny.Map
				_ = _1215_v127
				_1215_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).F32(), _991_v4)
				var _1216_v128 _dafny.Map
				_ = _1216_v128
				_1216_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
				var _1217_v129 _dafny.Array
				_ = _1217_v129
				var _nwElement0_36 _dafny.Int = (_1214_v126).Select((Companion_Default___.SafeIndex(_1127_cf12, _dafny.IntOfUint32((_1214_v126).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _nwElement0_36
				var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(15))
				_ = _nw172
				(_nw172).ArraySet1(_nwElement0_36, 0)
				(_nw172).ArraySet1((_dafny.Zero).Minus(_1127_cf12), 1)
				(_nw172).ArraySet1((_dafny.Zero).Minus(_991_v4), 2)
				(_nw172).ArraySet1((_this).Fm1(globalState), 3)
				(_nw172).ArraySet1(_1127_cf12, 4)
				(_nw172).ArraySet1((_1215_v127).Cardinality(), 5)
				(_nw172).ArraySet1(_dafny.IntOfInt64(698), 6)
				(_nw172).ArraySet1(_991_v4, 7)
				(_nw172).ArraySet1((_1216_v128).Cardinality(), 8)
				(_nw172).ArraySet1(_991_v4, 9)
				(_nw172).ArraySet1(_1127_cf12, 10)
				(_nw172).ArraySet1((_dafny.Zero).Minus(_991_v4), 11)
				(_nw172).ArraySet1(_1127_cf12, 12)
				(_nw172).ArraySet1(_991_v4, 13)
				(_nw172).ArraySet1(_1127_cf12, 14)
				_1217_v129 = _nw172
				var _1218_v130 _dafny.Map
				_ = _1218_v130
				_1218_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1217_v129, _1213_v125)
				var _1219_v131 *C1
				_ = _1219_v131
				var _nw173 *C1 = New_C1_()
				_ = _nw173
				_nw173.Ctor__(_dafny.IntOfUint32((_1213_v125).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1214_v126).Cardinality())), (_1218_v130).Contains(_1217_v129), (_995_v8).F32())
				_1219_v131 = _nw173
				var _1220_v132 D0
				_ = _1220_v132
				_1220_v132 = Companion_D0_.Create_DC0_((_this).F26())
				(globalState).F22 = !((_987_v0).F32()) || ((_this).Fm2((_this).F26(), (_1219_v131).Fm4(_1220_v132, (_1219_v131).F36(), _991_v4, globalState), globalState))
				var _1221_v133 _dafny.Map
				_ = _1221_v133
				_1221_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1213_v125, _dafny.UnicodeSeqOfUtf8Bytes("uxenwl")), _1058_v52)
				_1221_v133 = (_1221_v133).Update(_1213_v125, _1058_v52)
				_1215_v127 = (_1215_v127).Update((_995_v8).Fm6(_1058_v52, (_987_v0).F32(), globalState), _dafny.IntOfInt64(487))
			}
			if (_995_v8).F32() {
				var _1222_v134 _dafny.Sequence
				_ = _1222_v134
				_1222_v134 = _dafny.UnicodeSeqOfUtf8Bytes("ck")
				var _1223_v135 _dafny.MultiSet
				_ = _1223_v135
				_1223_v135 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1222_v134).Cardinality()))
				var _1224_v136 _dafny.Map
				_ = _1224_v136
				_1224_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1223_v135, (_987_v0).F32())
				var _1225_v137 _dafny.Set
				_ = _1225_v137
				_1225_v137 = _dafny.SetOf((_995_v8).F32(), (_987_v0).F32())
				var _1226_v138 _dafny.Map
				_ = _1226_v138
				_1226_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, _1225_v137)
				var _1227_v140 _dafny.Array
				_ = _1227_v140
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_35
				var _nw174 _dafny.Array
				_ = _nw174
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw174 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) bool = func(_1228_i15 _dafny.Int) bool {
						return (_this).F26()
					}
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw174 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw174).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw174).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1227_v140 = _nw174
				var _1229_v141 _dafny.MultiSet
				_ = _1229_v141
				_1229_v141 = _dafny.MultiSetOf(_1227_v140)
				var _1230_v142 _dafny.Map
				_ = _1230_v142
				_1230_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, (func() _dafny.Int {
					if (_1229_v141).Contains(_1227_v140) {
						return (_1229_v141).Multiplicity(_1227_v140)
					}
					return _1127_cf12
				})())
				var _1231_v143 _dafny.Map
				_ = _1231_v143
				_1231_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1230_v142, _dafny.IntOfInt64(53))
				var _1232_v144 _dafny.MultiSet
				_ = _1232_v144
				_1232_v144 = _dafny.MultiSetOf((_987_v0).F32())
				var _1233_v145 _dafny.Array
				_ = _1233_v145
				var _nwElement0_37 _dafny.Int = _1127_cf12
				_ = _nwElement0_37
				var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(24))
				_ = _nw175
				(_nw175).ArraySet1(_nwElement0_37, 0)
				(_nw175).ArraySet1((_1224_v136).Cardinality(), 1)
				(_nw175).ArraySet1((_991_v4).Times(_991_v4), 2)
				(_nw175).ArraySet1(((func() _dafny.Set {
					if (_1226_v138).Contains(_dafny.IntOfInt64(255)) {
						return (_1226_v138).Get(_dafny.IntOfInt64(255)).(_dafny.Set)
					}
					return _dafny.SetOf((_995_v8).F32())
				})()).Cardinality(), 3)
				(_nw175).ArraySet1((func() _dafny.Set {
					var _coll59 = _dafny.NewBuilder()
					_ = _coll59
					for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(158), _dafny.IntOfInt64(746))); ; {
						_compr_59, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1234_v139 _dafny.Int
						_1234_v139 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(158)).Cmp(_1234_v139) <= 0) && ((_1234_v139).Cmp(_dafny.IntOfInt64(746)) < 0) {
							_coll59.Add(Companion_Default___.SafeModuloInt(_1234_v139, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(858))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg73 _dafny.Int) interface{} {
									return coer73(arg73)
								}
							}(func(_1235_i14 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cuhflyu")).Cardinality()))
							}))).Cardinality())))
						}
					}
					return _coll59.ToSet()
				}()).Cardinality(), 4)
				(_nw175).ArraySet1((_dafny.Zero).Minus(_991_v4), 5)
				(_nw175).ArraySet1(_1127_cf12, 6)
				(_nw175).ArraySet1(_991_v4, 7)
				(_nw175).ArraySet1(_1127_cf12, 8)
				(_nw175).ArraySet1(_1127_cf12, 9)
				(_nw175).ArraySet1(_991_v4, 10)
				(_nw175).ArraySet1((_991_v4).Minus(_991_v4), 11)
				(_nw175).ArraySet1(_991_v4, 12)
				(_nw175).ArraySet1(Companion_Default___.SafeDivisionInt(_1127_cf12, _1127_cf12), 13)
				(_nw175).ArraySet1(_1127_cf12, 14)
				(_nw175).ArraySet1(_1127_cf12, 15)
				(_nw175).ArraySet1((_1127_cf12).Minus((_1231_v143).Cardinality()), 16)
				(_nw175).ArraySet1(_1127_cf12, 17)
				(_nw175).ArraySet1((_1232_v144).Cardinality(), 18)
				(_nw175).ArraySet1(_991_v4, 19)
				(_nw175).ArraySet1((_1230_v142).Cardinality(), 20)
				(_nw175).ArraySet1(_991_v4, 21)
				(_nw175).ArraySet1(((_this).Fm1(globalState)).Times(_991_v4), 22)
				(_nw175).ArraySet1(((_dafny.Zero).Minus(_991_v4)).Times(_1127_cf12), 23)
				_1233_v145 = _nw175
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))
				_ = _index160
				(_1233_v145).ArraySet1(_1127_cf12, (_index160).Int())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))
				_ = _index161
				(_1233_v145).ArraySet1((func() _dafny.Int {
					if !(((_this).F27()) == ((_this).F27())) {
						return _dafny.IntOfUint32((_1222_v134).Cardinality())
					}
					return _991_v4
				})(), (_index161).Int())
				(globalState).F16 = _991_v4
				var _1236_v146 _dafny.Map
				_ = _1236_v146
				_1236_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).F32(), _1127_cf12)
				var _1237_v147 D5
				_ = _1237_v147
				_1237_v147 = Companion_D5_.Create_DC14_((_1233_v145).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))).Int()).(_dafny.Int), _1127_cf12, !((_987_v0).F32()))
				var _1238_v148 _dafny.Sequence
				_ = _1238_v148
				_1238_v148 = _dafny.SeqOf(_1237_v147, _1237_v147)
				var _1239_v149 _dafny.Map
				_ = _1239_v149
				_1239_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).F32(), (_995_v8).F32())
				var _1240_v150 _dafny.Sequence
				_ = _1240_v150
				_1240_v150 = _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm35((func() _dafny.Int {
					if (_1236_v146).Contains((_987_v0).F32()) {
						return (_1236_v146).Get((_987_v0).F32()).(_dafny.Int)
					}
					return (_1233_v145).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))).Int()).(_dafny.Int)
				})(), globalState), (Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((Companion_Default___.Fm35((func() _dafny.Int {
					if (_1236_v146).Contains((_987_v0).F32()) {
						return (_1236_v146).Get((_987_v0).F32()).(_dafny.Int)
					}
					return (_1233_v145).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))).Int()).(_dafny.Int)
				})(), globalState)).Cardinality()))).Uint32(), _1237_v147), _1238_v148), !(_1239_v149).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_987_v0).F32())), (_995_v8).F32(), (_995_v8).F32())
				_1240_v150 = _dafny.Companion_Sequence_.Concatenate(p2, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_987_v0).F32(), (_995_v8).F32()), (_1059_v53).Dtor_cf6()))
				var _1241_v151 _dafny.MultiSet
				_ = _1241_v151
				_1241_v151 = _dafny.MultiSetOf(_1222_v134)
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))
				_ = _index162
				(_1233_v145).ArraySet1((func() _dafny.Int {
					if (_1241_v151).Contains(_1222_v134) {
						return (_1241_v151).Multiplicity(_1222_v134)
					}
					return (_1233_v145).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))).Int()).(_dafny.Int)
				})(), (_index162).Int())
				var _1242_v152 *C1
				_ = _1242_v152
				var _nw176 *C1 = New_C1_()
				_ = _nw176
				_nw176.Ctor__((_dafny.Zero).Minus((_1233_v145).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.ArrayLen((_1233_v145), 0))).Int()).(_dafny.Int)), _1127_cf12, (_987_v0).F32(), (_this).F27())
				_1242_v152 = _nw176
			} else {
				_1127_cf12 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_1127_cf12, _1127_cf12))
				var _1243_v153 _dafny.MultiSet
				_ = _1243_v153
				_1243_v153 = _dafny.MultiSetOf(_dafny.IntOfInt64(878), _991_v4)
				var _1244_v154 bool
				_ = _1244_v154
				var _out42 bool
				_ = _out42
				_out42 = (_this).M8(Companion_D0_.Create_DC0_((_987_v0).F32()), (_995_v8).F32(), _dafny.MultiSetOf((_987_v0).F32(), (_987_v0).F32(), p0), Companion_Default___.SafeModuloInt(_1127_cf12, ((_1243_v153).Update(_1127_cf12, Companion_Default___.Abs(_991_v4))).Cardinality()), globalState)
				_1244_v154 = _out42
				var _1245_v155 _dafny.Array
				_ = _1245_v155
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_36
				var _nw177 _dafny.Array
				_ = _nw177
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw177 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = func(_1246_i16 _dafny.Int) _dafny.Int {
						return (_1246_i16).Plus(_dafny.IntOfInt64(121))
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw177 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw177).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw177).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1245_v155 = _nw177
				var _1247_v156 D3
				_ = _1247_v156
				_1247_v156 = Companion_D3_.Create_DC10_(_1245_v155)
				var _1248_v157 _dafny.Map
				_ = _1248_v157
				_1248_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1247_v156, (_dafny.Zero).Minus((_dafny.Zero).Minus(_991_v4)))
				var _1249_v158 *C4
				_ = _1249_v158
				var _nw178 *C4 = New_C4_()
				_ = _nw178
				_nw178.Ctor__((func() _dafny.Map {
					if (_987_v0).F32() {
						return _1248_v157
					}
					return _1248_v157
				})(), (_987_v0).F32(), (_this).F27())
				_1249_v158 = _nw178
				(globalState).F2 = (_995_v8).F32()
				var _1250_v159 _dafny.Array
				_ = _1250_v159
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw179
				_1250_v159 = _nw179
				var _1251_v160 _dafny.Sequence
				_ = _1251_v160
				_1251_v160 = _dafny.UnicodeSeqOfUtf8Bytes("hnfdufwb")
				var _1252_v161 _dafny.Map
				_ = _1252_v161
				_1252_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1127_cf12, _1251_v160)
				var _1253_v162 _dafny.Sequence
				_ = _1253_v162
				_1253_v162 = _dafny.SeqOf(_1252_v161, _1252_v161, _1252_v161, _1252_v161)
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1250_v159), 0))
				_ = _index163
				(_1250_v159).ArraySet1(((_1253_v162).Select((Companion_Default___.SafeIndex(_1127_cf12, _dafny.IntOfUint32((_1253_v162).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1252_v161), (_index163).Int())
				var _1254_v163 D5
				_ = _1254_v163
				_1254_v163 = Companion_D5_.Create_DC13_((_1128_v80).Update((_this).F26(), _dafny.SeqOf((_987_v0).F32())))
				var _1255_v164 D5
				_ = _1255_v164
				_1255_v164 = Companion_D5_.Create_DC15_(_1254_v163)
				var _1256_v165 D5
				_ = _1256_v165
				_1256_v165 = Companion_D5_.Create_DC15_(Companion_D5_.Create_DC15_(_1255_v164))
				var _1257_v166 D5
				_ = _1257_v166
				_1257_v166 = Companion_D5_.Create_DC15_(_1255_v164)
				var _1258_v167 _dafny.Array
				_ = _1258_v167
				var _nwElement0_38 D5 = Companion_D5_.Create_DC15_(_1254_v163)
				_ = _nwElement0_38
				var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(6))
				_ = _nw180
				(_nw180).ArraySet1(_nwElement0_38, 0)
				(_nw180).ArraySet1(_1257_v166, 1)
				(_nw180).ArraySet1(_1257_v166, 2)
				(_nw180).ArraySet1(_1257_v166, 3)
				(_nw180).ArraySet1(_1257_v166, 4)
				(_nw180).ArraySet1(_1257_v166, 5)
				_1258_v167 = _nw180
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1250_v159), 0))
				_ = _index164
				var _rhs178 _dafny.Map = (func() _dafny.Map {
					var _coll60 = _dafny.NewMapBuilder()
					_ = _coll60
					for _iter60 := _dafny.Iterate(((_993_v7).Dtor_cf13()).Keys().Elements()); ; {
						_compr_60, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _1259_v168 _dafny.Int
						_1259_v168 = interface{}(_compr_60).(_dafny.Int)
						if ((_993_v7).Dtor_cf13()).Contains(_1259_v168) {
							_coll60.Add((_1259_v168).Plus(_dafny.IntOfInt64(725)), _1251_v160)
						}
					}
					return _coll60.ToMap()
				}()).Merge((_1252_v161).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_991_v4), _1251_v160)))
				_ = _rhs178
				var _rhs179 bool = p0
				_ = _rhs179
				var _rhs180 _dafny.Array = _1258_v167
				_ = _rhs180
				var _rhs181 bool = !((_987_v0).F32())
				_ = _rhs181
				var _lhs162 _dafny.Array = _1250_v159
				_ = _lhs162
				var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_1250_v159), 0))
				_ = _lhs163
				var _lhs164 *GlobalState = globalState
				_ = _lhs164
				(_lhs162).ArraySet1(_rhs178, (_lhs163).Int())
				_1244_v154 = _rhs179
				_1258_v167 = _rhs180
				_lhs164.F2 = _rhs181
			}
			if (_this).F26() {
				var _1260_v169 _dafny.Map
				_ = _1260_v169
				_1260_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.Zero).Minus(_1127_cf12)).Plus(_1127_cf12), (_995_v8).F32())
				var _1261_v170 _dafny.Map
				_ = _1261_v170
				_1261_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), _991_v4)
				_1260_v169 = (_1260_v169).Update((_1261_v170).Cardinality(), (_987_v0).F32())
				var _1262_v171 _dafny.Array
				_ = _1262_v171
				var _nw181 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(16))
				_ = _nw181
				_1262_v171 = _nw181
				var _1263_v172 D0
				_ = _1263_v172
				_1263_v172 = Companion_D0_.Create_DC0_((_this).F27())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1262_v171), 0))
				_ = _index165
				(_1262_v171).ArraySet1(_1263_v172, (_index165).Int())
				var _1264_v173 _dafny.Map
				_ = _1264_v173
				_1264_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_995_v8).F32())
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1262_v171), 0))
				_ = _index166
				var _rhs182 D0 = Companion_D0_.Create_DC0_((func() bool {
					if (_1264_v173).Contains((_995_v8).F32()) {
						return (_1264_v173).Get((_995_v8).F32()).(bool)
					}
					return (_987_v0).F32()
				})())
				_ = _rhs182
				var _rhs183 _dafny.Int = _991_v4
				_ = _rhs183
				var _lhs165 _dafny.Array = _1262_v171
				_ = _lhs165
				var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1262_v171), 0))
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				(_lhs165).ArraySet1(_rhs182, (_lhs166).Int())
				_lhs167.F16 = _rhs183
				var _1265_v174 _dafny.Sequence
				_ = _1265_v174
				_1265_v174 = _dafny.SeqOf(_1127_cf12, _991_v4)
				var _1266_v175 _dafny.Sequence
				_ = _1266_v175
				_1266_v175 = _1265_v174
				var _1267_v176 _dafny.Sequence
				_ = _1267_v176
				_1267_v176 = _dafny.SeqOf(_1266_v175, _1266_v175)
				var _1268_v177 _dafny.Sequence
				_ = _1268_v177
				_1268_v177 = _dafny.UnicodeSeqOfUtf8Bytes("iywdtfdjg")
				var _1269_v178 D6
				_ = _1269_v178
				_1269_v178 = Companion_D6_.Create_DC17_(_1127_cf12, _1127_cf12, _991_v4, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1268_v177, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-75), _dafny.IntOfUint32((_1268_v177).Cardinality()))).Uint32(), _1058_v52)).Cardinality()), _991_v4)
				var _1270_v179 _dafny.Map
				_ = _1270_v179
				_1270_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1267_v176, _1269_v178)
				_1270_v179 = _1270_v179
				var _1271_v180 D0
				_ = _1271_v180
				_1271_v180 = Companion_D0_.Create_DC1_((_987_v0).F32(), (_995_v8).F32(), _dafny.IntOfUint32((_1268_v177).Cardinality()), p0)
				var _1272_v181 *C1
				_ = _1272_v181
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__(_991_v4, _991_v4, (_this).F26(), (_this).F27())
				_1272_v181 = _nw182
				var _1273_v182 _dafny.MultiSet
				_ = _1273_v182
				_1273_v182 = _dafny.MultiSetOf(_1272_v181)
				var _1274_v183 _dafny.Set
				_ = _1274_v183
				_1274_v183 = _dafny.SetOf((_this).F26(), (_995_v8).F32())
				var _1275_v184 _dafny.Map
				_ = _1275_v184
				_1275_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1058_v52)
				var _1276_v185 D5
				_ = _1276_v185
				_1276_v185 = Companion_D5_.Create_DC14_(_1272_v181.F37, _dafny.IntOfUint32((_1268_v177).Cardinality()), (_995_v8).F32())
				var _1277_v186 _dafny.Array
				_ = _1277_v186
				var _nwElement0_39 bool = !(!((_dafny.IntOfInt64(-994)).Cmp(_991_v4) != 0))
				_ = _nwElement0_39
				var _nw183 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(20))
				_ = _nw183
				(_nw183).ArraySet1(_nwElement0_39, 0)
				(_nw183).ArraySet1((_995_v8).F32(), 1)
				(_nw183).ArraySet1((_this).Fm2(true, (_this).Fm1(globalState), globalState), 2)
				(_nw183).ArraySet1((_995_v8).F32(), 3)
				(_nw183).ArraySet1((_995_v8).F32(), 4)
				(_nw183).ArraySet1(!((_987_v0).Fm7(_1261_v170, _1265_v174, (_this).F27(), globalState)) || (true), 5)
				(_nw183).ArraySet1(!((_987_v0).F32()) || ((_1271_v180).Dtor_cf1()), 6)
				(_nw183).ArraySet1((_this).F26(), 7)
				(_nw183).ArraySet1((_this).F27(), 8)
				(_nw183).ArraySet1((_995_v8).F32(), 9)
				(_nw183).ArraySet1(((_987_v0).Fm7(_1261_v170, _1265_v174, !(p0), globalState)) && ((_this).F26()), 10)
				(_nw183).ArraySet1(((func() _dafny.Int {
					if (_1273_v182).Contains(_1272_v181) {
						return (_1273_v182).Multiplicity(_1272_v181)
					}
					return _dafny.IntOfUint32((_1265_v174).Cardinality())
				})()).Cmp((_1274_v183).Cardinality()) < 0, 11)
				(_nw183).ArraySet1((_987_v0).Fm6(_1058_v52, !((_987_v0).Fm7(_1261_v170, _dafny.SeqOf((_1275_v184).Cardinality()), (_987_v0).F32(), globalState)), globalState), 12)
				(_nw183).ArraySet1(true, 13)
				(_nw183).ArraySet1(!(!((_this).F27())) || ((_987_v0).F32()), 14)
				(_nw183).ArraySet1((_1276_v185).Dtor_cf27(), 15)
				(_nw183).ArraySet1((func() bool {
					if p0 {
						return (_995_v8).F32()
					}
					return p0
				})(), 16)
				(_nw183).ArraySet1((!(!(p0))) == ((_987_v0).F32()), 17)
				(_nw183).ArraySet1((_995_v8).F32(), 18)
				(_nw183).ArraySet1(_dafny.Companion_Sequence_.Equal(p2, _dafny.SeqOf(p0, p0, (_995_v8).F32())), 19)
				_1277_v186 = _nw183
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_1277_v186), 0))
				_ = _index167
				(_1277_v186).ArraySet1((_995_v8).F32(), (_index167).Int())
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(947), _dafny.ArrayLen((_1277_v186), 0))
				_ = _index168
				(_1277_v186).ArraySet1(!((_987_v0).F32()), (_index168).Int())
				var _1278_v187 _dafny.Array
				_ = _1278_v187
				var _nw184 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
				_ = _nw184
				_1278_v187 = _nw184
				var _1279_v188 _dafny.Sequence
				_ = _1279_v188
				_1279_v188 = _dafny.SeqOf(_1278_v187)
				(globalState).F22 = !_dafny.Companion_Sequence_.Equal(_1279_v188, _1279_v188)
			} else {
				(globalState).F12 = _1127_cf12
				var _1280_v189 _dafny.Sequence
				_ = _1280_v189
				_1280_v189 = _dafny.SeqOf(_dafny.IntOfInt64(911), _dafny.IntOfUint32((p2).Cardinality()))
				var _1281_v190 _dafny.Sequence
				_ = _1281_v190
				_1281_v190 = _dafny.SeqOf(_1280_v189)
				var _1282_v191 T2
				_ = _1282_v191
				var _nw185 *C2 = New_C2_()
				_ = _nw185
				_nw185.Ctor__(_1281_v190, (_987_v0).F32(), (p2).Select((Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))
				_1282_v191 = _nw185
				var _1283_v192 D7
				_ = _1283_v192
				_1283_v192 = Companion_D7_.Create_DC18_(_1282_v191)
				var _1284_v193 _dafny.MultiSet
				_ = _1284_v193
				_1284_v193 = _dafny.MultiSetOf(_1283_v192)
				var _1285_v194 _dafny.Sequence
				_ = _1285_v194
				_1285_v194 = _dafny.SeqOf(_1284_v193)
				var _1286_v195 _dafny.Sequence
				_ = _1286_v195
				_1286_v195 = _dafny.SeqOf(_1283_v192, _1283_v192)
				(globalState).F24 = !(((_dafny.MultiSetFromSeq(_1286_v195)).Update(_1283_v192, Companion_Default___.Abs(_1127_cf12))).IsSubsetOf((_1285_v194).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(38), _dafny.IntOfUint32((_1285_v194).Cardinality()))).Uint32()).(_dafny.MultiSet)))
				var _1287_v196 _dafny.Array
				_ = _1287_v196
				var _nw186 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw186
				_1287_v196 = _nw186
				var _rhs184 _dafny.Array = _1287_v196
				_ = _rhs184
				var _rhs185 bool = !((_this).F27()) || ((_995_v8).F32())
				_ = _rhs185
				var _rhs186 _dafny.Int = (_this).Fm1(globalState)
				_ = _rhs186
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				var _lhs169 *GlobalState = globalState
				_ = _lhs169
				_1287_v196 = _rhs184
				_lhs168.F2 = _rhs185
				_lhs169.F16 = _rhs186
				var _1288_v197 _dafny.Array
				_ = _1288_v197
				var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw187
				_1288_v197 = _nw187
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_1288_v197), 0))
				_ = _index169
				(_1288_v197).ArraySet1(_1127_cf12, (_index169).Int())
				var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(581), _dafny.ArrayLen((_1288_v197), 0))
				_ = _index170
				(_1288_v197).ArraySet1(_1127_cf12, (_index170).Int())
				var _1289_v198 _dafny.Set
				_ = _1289_v198
				_1289_v198 = _dafny.SetOf(_1288_v197)
				_1289_v198 = (_dafny.SetOf(_1288_v197, _1288_v197, _1288_v197, _1288_v197)).Intersection(_dafny.SetOf(_1288_v197, _1288_v197, _1288_v197))
			}
		} else {
			var _1290___mcc_h14 _dafny.Sequence = _source18.Get_().(D1_DC3).Cf6
			_ = _1290___mcc_h14
			var _1291_cf6 _dafny.Sequence = _1290___mcc_h14
			_ = _1291_cf6
			var _1292_v199 _dafny.Map
			_ = _1292_v199
			_1292_v199 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), p0)
			(globalState).F15 = (func() bool {
				if (_1292_v199).Contains((_this).F27()) {
					return (_1292_v199).Get((_this).F27()).(bool)
				}
				return (p2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)
			})()
			var _1293_v200 _dafny.Map
			_ = _1293_v200
			_1293_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_995_v8).F32(), _991_v4)
			var _1294_v201 _dafny.Sequence
			_ = _1294_v201
			_1294_v201 = _dafny.SeqOf(_991_v4)
			var _1295_v202 _dafny.Map
			_ = _1295_v202
			_1295_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).Fm7(_1293_v200, _1294_v201, (p2).Select((Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, _dafny.IntOfInt64(-45)))
			var _1296_v203 _dafny.Map
			_ = _1296_v203
			_1296_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_991_v4, _dafny.IntOfInt64(-349))
			var _1297_v204 _dafny.Map
			_ = _1297_v204
			_1297_v204 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1296_v203).Contains(_991_v4) {
					return (_1296_v203).Get(_991_v4).(_dafny.Int)
				}
				return _dafny.IntOfInt64(912)
			})(), _dafny.IntOfInt64(2))
			var _1298_v205 _dafny.Sequence
			_ = _1298_v205
			_1298_v205 = _dafny.UnicodeSeqOfUtf8Bytes("gwqodhw")
			var _1299_v206 _dafny.Map
			_ = _1299_v206
			_1299_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rphevkuq")).Cardinality()), (_987_v0).Fm6(_1058_v52, (_987_v0).F32(), globalState))
			var _1300_v207 _dafny.MultiSet
			_ = _1300_v207
			_1300_v207 = _dafny.MultiSetOf(_1294_v201, _1294_v201)
			var _1301_v208 *C3
			_ = _1301_v208
			var _nw188 *C3 = New_C3_()
			_ = _nw188
			_nw188.Ctor__(_1300_v207, (_this).F27(), (_987_v0).F32())
			_1301_v208 = _nw188
			var _1302_v209 _dafny.Array
			_ = _1302_v209
			var _nwElement0_40 _dafny.Int = _991_v4
			_ = _nwElement0_40
			var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(22))
			_ = _nw189
			(_nw189).ArraySet1(_nwElement0_40, 0)
			(_nw189).ArraySet1(_991_v4, 1)
			(_nw189).ArraySet1((_1292_v199).Cardinality(), 2)
			(_nw189).ArraySet1(_991_v4, 3)
			(_nw189).ArraySet1(_991_v4, 4)
			(_nw189).ArraySet1((_1299_v206).Cardinality(), 5)
			(_nw189).ArraySet1(_dafny.IntOfUint32((_1298_v205).Cardinality()), 6)
			(_nw189).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_987_v0).F32())).Cardinality()), 7)
			(_nw189).ArraySet1(_991_v4, 8)
			(_nw189).ArraySet1(_991_v4, 9)
			(_nw189).ArraySet1(_991_v4, 10)
			(_nw189).ArraySet1(_991_v4, 11)
			(_nw189).ArraySet1(_991_v4, 12)
			(_nw189).ArraySet1(_dafny.IntOfInt64(138), 13)
			(_nw189).ArraySet1(_991_v4, 14)
			(_nw189).ArraySet1(_991_v4, 15)
			(_nw189).ArraySet1(_991_v4, 16)
			(_nw189).ArraySet1(_dafny.IntOfInt64(61), 17)
			(_nw189).ArraySet1(_dafny.IntOfInt64(540), 18)
			(_nw189).ArraySet1(_991_v4, 19)
			(_nw189).ArraySet1(_991_v4, 20)
			(_nw189).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1301_v208)).Cardinality()), 21)
			_1302_v209 = _nw189
			var _1303_v210 D3
			_ = _1303_v210
			_1303_v210 = Companion_D3_.Create_DC10_(_1302_v209)
			var _1304_v211 *C4
			_ = _1304_v211
			var _nw190 *C4 = New_C4_()
			_ = _nw190
			_nw190.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1303_v210, _dafny.IntOfUint32((_1298_v205).Cardinality())), (_this).F27(), (_995_v8).F32())
			_1304_v211 = _nw190
			var _1305_v212 _dafny.Map
			_ = _1305_v212
			_1305_v212 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1304_v211, (_this).F27())
			var _1306_v213 _dafny.Map
			_ = _1306_v213
			_1306_v213 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1294_v201, _991_v4)
			var _1307_v214 _dafny.Array
			_ = _1307_v214
			var _nwElement0_41 _dafny.Int = _991_v4
			_ = _nwElement0_41
			var _nw191 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(29))
			_ = _nw191
			(_nw191).ArraySet1(_nwElement0_41, 0)
			(_nw191).ArraySet1(_dafny.IntOfInt64(-872), 1)
			(_nw191).ArraySet1(_991_v4, 2)
			(_nw191).ArraySet1(((_1295_v202).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_987_v0).F32(), _1296_v203))).Cardinality(), 3)
			(_nw191).ArraySet1(_991_v4, 4)
			(_nw191).ArraySet1(_dafny.IntOfInt64(-557), 5)
			(_nw191).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-248), _991_v4), 6)
			(_nw191).ArraySet1(_991_v4, 7)
			(_nw191).ArraySet1(_991_v4, 8)
			(_nw191).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_this).F26() {
					return _dafny.Companion_Sequence_.Update(_1298_v205, (Companion_Default___.SafeIndex((_1292_v199).Cardinality(), _dafny.IntOfUint32((_1298_v205).Cardinality()))).Uint32(), _1058_v52)
				}
				return _1298_v205
			})()).Cardinality())), 9)
			(_nw191).ArraySet1(((_dafny.Zero).Minus(_991_v4)).Minus(_991_v4), 10)
			(_nw191).ArraySet1((_dafny.IntOfInt64(354)).Minus((_1294_v201).Select((Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_1294_v201).Cardinality()))).Uint32()).(_dafny.Int)), 11)
			(_nw191).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_991_v4, (_1299_v206).Cardinality())), 12)
			(_nw191).ArraySet1(_991_v4, 13)
			(_nw191).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1298_v205, (Companion_Default___.SafeIndex(_991_v4, _dafny.IntOfUint32((_1298_v205).Cardinality()))).Uint32(), _1058_v52)).Cardinality()), 14)
			(_nw191).ArraySet1(_dafny.IntOfInt64(850), 15)
			(_nw191).ArraySet1(_dafny.IntOfInt64(942), 16)
			(_nw191).ArraySet1(_991_v4, 17)
			(_nw191).ArraySet1(_991_v4, 18)
			(_nw191).ArraySet1((_dafny.Zero).Minus(((_1305_v212).Merge(_1305_v212)).Cardinality()), 19)
			(_nw191).ArraySet1(_991_v4, 20)
			(_nw191).ArraySet1((_1306_v213).Cardinality(), 21)
			(_nw191).ArraySet1((func() _dafny.Int {
				if (_995_v8).F32() {
					return _dafny.IntOfInt64(196)
				}
				return _dafny.IntOfInt64(664)
			})(), 22)
			(_nw191).ArraySet1(Companion_Default___.SafeModuloInt(_991_v4, _991_v4), 23)
			(_nw191).ArraySet1(_991_v4, 24)
			(_nw191).ArraySet1(_991_v4, 25)
			(_nw191).ArraySet1(_991_v4, 26)
			(_nw191).ArraySet1(_991_v4, 27)
			(_nw191).ArraySet1(_dafny.IntOfInt64(480), 28)
			_1307_v214 = _nw191
			_1307_v214 = _1302_v209
			_1305_v212 = _1305_v212
			var _1308_v215 _dafny.Array
			_ = _1308_v215
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_37
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Sequence = func(_1309_i17 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf((_this).F26(), (_this).F27())
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw192 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw192).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw192).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1308_v215 = _nw192
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1308_v215), 0))
			_ = _index171
			(_1308_v215).ArraySet1(_1291_cf6, (_index171).Int())
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen((_1308_v215), 0))
			_ = _index172
			(_1308_v215).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p2, _1291_cf6), (_index172).Int())
		}
	}
}
func (_this *C5) M7(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) {
	{
		var _1310_i0 _dafny.Int
		_ = _1310_i0
		_1310_i0 = _dafny.Zero
		{
			for (_this).F27() {
				{
					if (_1310_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1310_i0 = (_1310_i0).Plus(_dafny.One)
					var _1311_v0 *C1
					_ = _1311_v0
					var _nw193 *C1 = New_C1_()
					_ = _nw193
					_nw193.Ctor__(p2, p2, (_this).F26(), p1)
					_1311_v0 = _nw193
					var _1312_v1 _dafny.Array
					_ = _1312_v1
					var _nwElement0_42 *C1 = _1311_v0
					_ = _nwElement0_42
					var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(3))
					_ = _nw194
					(_nw194).ArraySet1(_nwElement0_42, 0)
					(_nw194).ArraySet1(_1311_v0, 1)
					(_nw194).ArraySet1(_1311_v0, 2)
					_1312_v1 = _nw194
					var _1313_v2 _dafny.Sequence
					_ = _1313_v2
					_1313_v2 = _dafny.SeqOf(_1312_v1)
					(globalState).F2 = _dafny.Companion_Sequence_.IsPrefixOf(_1313_v2, _1313_v2)
					var _1314_v3 _dafny.Sequence
					_ = _1314_v3
					_1314_v3 = _dafny.UnicodeSeqOfUtf8Bytes("jn")
					var _1315_v4 _dafny.Set
					_ = _1315_v4
					_1315_v4 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(825))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}(func(_1316_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('c')
					})), _dafny.UnicodeSeqOfUtf8Bytes("yknlqwi")), _1314_v3)
					(_1311_v0).F37 = (_1315_v4).Cardinality()
					var _1317_v5 _dafny.Sequence
					_ = _1317_v5
					_1317_v5 = _dafny.SeqOf(_1314_v3)
					_1317_v5 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1317_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(330))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg75 _dafny.Int) interface{} {
							return coer75(arg75)
						}
					}((func(_1318_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1319_i2 _dafny.Int) _dafny.Sequence {
							return _1318_v3
						}
					})(_1314_v3)))), _1317_v5), (Companion_Default___.SafeIndex(_1311_v0.F37, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1317_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(330))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg76 _dafny.Int) interface{} {
							return coer76(arg76)
						}
					}((func(_1320_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1321_i2 _dafny.Int) _dafny.Sequence {
							return _1320_v3
						}
					})(_1314_v3)))), _1317_v5)).Cardinality()))).Uint32(), _1314_v3)
					var _1322_v6 _dafny.Array
					_ = _1322_v6
					var _nw195 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
					_ = _nw195
					_1322_v6 = _nw195
					var _1323_v7 _dafny.Array
					_ = _1323_v7
					var _nw196 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
					_ = _nw196
					_1323_v7 = _nw196
					var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1322_v6), 0))
					_ = _index173
					(_1322_v6).ArraySet1(_1323_v7, (_index173).Int())
					var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(961), _dafny.ArrayLen((_1322_v6), 0))
					_ = _index174
					(_1322_v6).ArraySet1(_1323_v7, (_index174).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1324_v8 _dafny.Array
		_ = _1324_v8
		var _len0_38 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_38
		var _nw197 _dafny.Array
		_ = _nw197
		if _len0_38.Cmp(_dafny.Zero) == 0 {
			_nw197 = _dafny.NewArray(_len0_38)
		} else {
			var _init38 func(_dafny.Int) bool = func(_1325_i3 _dafny.Int) bool {
				return (_this).F26()
			}
			_ = _init38
			var _element0_38 = _init38(_dafny.Zero)
			_ = _element0_38
			_nw197 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
			(_nw197).ArraySet1(_element0_38, 0)
			var _nativeLen0_38 = (_len0_38).Int()
			_ = _nativeLen0_38
			for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
				(_nw197).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
			}
		}
		_1324_v8 = _nw197
		var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
		_ = _index175
		(_1324_v8).ArraySet1((_this).F26(), (_index175).Int())
		var _1326_v9 _dafny.Array
		_ = _1326_v9
		var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw198
		_1326_v9 = _nw198
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1324_v8), 0))
		_ = _index176
		(_1324_v8).ArraySet1((_this).F26(), (_index176).Int())
		var _1327_v10 _dafny.Sequence
		_ = _1327_v10
		_1327_v10 = _dafny.SeqOf(p2)
		var _1328_v11 _dafny.MultiSet
		_ = _1328_v11
		_1328_v11 = _dafny.MultiSetOf(_1327_v10, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-396))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg77 _dafny.Int) interface{} {
				return coer77(arg77)
			}
		}((func(_1329_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1330_i4 _dafny.Int) _dafny.Int {
				return _1329_p2
			}
		})(p2))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-396))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg78 _dafny.Int) interface{} {
				return coer78(arg78)
			}
		}((func(_1331_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1332_i4 _dafny.Int) _dafny.Int {
				return _1331_p2
			}
		})(p2)))).Cardinality()))).Uint32(), p2))
		var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
		_ = _index177
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1324_v8), 0))
		_ = _index178
		var _rhs187 bool = false
		_ = _rhs187
		var _rhs188 bool = (_this).F27()
		_ = _rhs188
		var _rhs189 _dafny.Array = _1326_v9
		_ = _rhs189
		var _rhs190 bool = p1
		_ = _rhs190
		var _rhs191 bool = ((_1328_v11).Union(_1328_v11)).IsSubsetOf(_1328_v11)
		_ = _rhs191
		var _lhs170 *GlobalState = globalState
		_ = _lhs170
		var _lhs171 _dafny.Array = _1324_v8
		_ = _lhs171
		var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
		_ = _lhs172
		var _lhs173 _dafny.Array = _1324_v8
		_ = _lhs173
		var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen((_1324_v8), 0))
		_ = _lhs174
		var _lhs175 *GlobalState = globalState
		_ = _lhs175
		_lhs170.F15 = _rhs187
		(_lhs171).ArraySet1(_rhs188, (_lhs172).Int())
		_1326_v9 = _rhs189
		(_lhs173).ArraySet1(_rhs190, (_lhs174).Int())
		_lhs175.F2 = _rhs191
		var _hi11 _dafny.Int = p2
		_ = _hi11
		for _1333_i5 := Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).Fm1(globalState)), p2); _1333_i5.Cmp(_hi11) < 0; _1333_i5 = _1333_i5.Plus(_dafny.One) {
			(globalState).F12 = p2
			var _1334_v12 _dafny.Map
			_ = _1334_v12
			_1334_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2((_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool), (_this).Fm1(globalState), globalState), p2)
			(globalState).F16 = (func() _dafny.Int {
				if (_1333_i5).Cmp(p2) <= 0 {
					return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(463))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg79 _dafny.Int) interface{} {
							return coer79(arg79)
						}
					}(func(_1335_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('i')
					}))).Cardinality())).Plus((_1334_v12).Cardinality())
				}
				return _dafny.IntOfInt64(887)
			})()
			var _1336_v13 _dafny.Sequence
			_ = _1336_v13
			_1336_v13 = _dafny.SeqOf(!(p1), (_this).F26())
			(globalState).F15 = (func() bool {
				if !(!(true)) {
					return (_1336_v13).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1336_v13).Cardinality()))).Uint32()).(bool)
				}
				return p0
			})()
			if p1 {
				var _1337_v14 _dafny.Map
				_ = _1337_v14
				_1337_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1338_v15 _dafny.Map
				_ = _1338_v15
				_1338_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1337_v14).Cardinality())
				var _1339_v16 D2
				_ = _1339_v16
				_1339_v16 = Companion_D2_.Create_DC6_(_1338_v15)
				var _1340_v17 _dafny.Sequence
				_ = _1340_v17
				_1340_v17 = _dafny.SeqOf(_1339_v16, Companion_D2_.Create_DC6_(_1338_v15))
				var _1341_v18 _dafny.MultiSet
				_ = _1341_v18
				_1341_v18 = _dafny.MultiSetOf(_1340_v17, _1340_v17, _1340_v17, _1340_v17, _1340_v17)
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
				_ = _index179
				(_1324_v8).ArraySet1((_1341_v18).IsSubsetOf((func() _dafny.MultiSet {
					if (_this).F26() {
						return _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_1340_v17, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.IntOfUint32((_1340_v17).Cardinality()))).Uint32(), _1339_v16))
					}
					return Companion_Default___.Fm36(p2, (_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool), !((_this).Fm2(p1, p2, globalState)), (_this).F26(), globalState)
				})()), (_index179).Int())
				var _1342_v19 D3
				_ = _1342_v19
				_1342_v19 = Companion_D3_.Create_DC10_(_1326_v9)
				var _1343_v20 _dafny.Map
				_ = _1343_v20
				_1343_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1324_v8)
				var _1344_v21 _dafny.Map
				_ = _1344_v21
				_1344_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1342_v19, (_1343_v20).Cardinality())
				var _1345_v22 *C4
				_ = _1345_v22
				var _nw199 *C4 = New_C4_()
				_ = _nw199
				_nw199.Ctor__(_1344_v21, p1, p1)
				_1345_v22 = _nw199
				var _1346_v23 _dafny.Sequence
				_ = _1346_v23
				_1346_v23 = _dafny.UnicodeSeqOfUtf8Bytes("gegvu")
				(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_1346_v23, _1346_v23)
				var _1347_v24 _dafny.Sequence
				_ = _1347_v24
				_1347_v24 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(117))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg80 _dafny.Int) interface{} {
						return coer80(arg80)
					}
				}((func(_1348_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1349_i7 _dafny.Int) _dafny.Int {
						return _1348_p2
					}
				})(p2))), _dafny.SeqOf(p2), _dafny.SeqOf((_dafny.Zero).Minus((_1338_v15).Cardinality()), _1333_i5), _1327_v10, _dafny.SeqOf(_1333_i5, _1333_i5))
				var _1350_v25 *C2
				_ = _1350_v25
				var _nw200 *C2 = New_C2_()
				_ = _nw200
				_nw200.Ctor__(_1347_v24, (_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool), (_this).F27())
				_1350_v25 = _nw200
				var _1351_v26 _dafny.Map
				_ = _1351_v26
				_1351_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1327_v10, _1350_v25)
				var _1352_v27 _dafny.Array
				_ = _1352_v27
				var _nwElement0_43 _dafny.Map = (_1351_v26).Merge(_1351_v26)
				_ = _nwElement0_43
				var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(3))
				_ = _nw201
				(_nw201).ArraySet1(_nwElement0_43, 0)
				(_nw201).ArraySet1(_1351_v26, 1)
				(_nw201).ArraySet1(_1351_v26, 2)
				_1352_v27 = _nw201
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1352_v27), 0))
				_ = _index180
				(_1352_v27).ArraySet1(_1351_v26, (_index180).Int())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1352_v27), 0))
				_ = _index181
				(_1352_v27).ArraySet1(_1351_v26, (_index181).Int())
				var _1353_v28 D0
				_ = _1353_v28
				_1353_v28 = Companion_D0_.Create_DC0_((_this).F26())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1326_v9), 0))
				_ = _index182
				(_1326_v9).ArraySet1(((_1350_v25).Fm4(_1353_v28, _1333_i5, p2, globalState)).Minus(_1333_i5), (_index182).Int())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1326_v9), 0))
				_ = _index183
				var _rhs192 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1336_v13).Select((Companion_Default___.SafeIndex(_1333_i5, _dafny.IntOfUint32((_1336_v13).Cardinality()))).Uint32()).(bool) {
						return _dafny.IntOfInt64(606)
					}
					return (_1333_i5).Minus(p2)
				})())
				_ = _rhs192
				var _rhs193 _dafny.Int = _1333_i5
				_ = _rhs193
				var _lhs176 *GlobalState = globalState
				_ = _lhs176
				var _lhs177 _dafny.Array = _1326_v9
				_ = _lhs177
				var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1326_v9), 0))
				_ = _lhs178
				_lhs176.F12 = _rhs192
				(_lhs177).ArraySet1(_rhs193, (_lhs178).Int())
			} else {
				var _1354_v29 _dafny.Array
				_ = _1354_v29
				var _nwElement0_44 _dafny.Array = _1326_v9
				_ = _nwElement0_44
				var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(6))
				_ = _nw202
				(_nw202).ArraySet1(_nwElement0_44, 0)
				(_nw202).ArraySet1(_1326_v9, 1)
				(_nw202).ArraySet1(_1326_v9, 2)
				(_nw202).ArraySet1(_1326_v9, 3)
				(_nw202).ArraySet1(_1326_v9, 4)
				(_nw202).ArraySet1(_1326_v9, 5)
				_1354_v29 = _nw202
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1354_v29), 0))
				_ = _index184
				(_1354_v29).ArraySet1(_1326_v9, (_index184).Int())
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1354_v29), 0))
				_ = _index185
				(_1354_v29).ArraySet1(_1326_v9, (_index185).Int())
				var _1355_v30 _dafny.Sequence
				_ = _1355_v30
				_1355_v30 = _dafny.UnicodeSeqOfUtf8Bytes("ctfdse")
				var _1356_v31 _dafny.Sequence
				_ = _1356_v31
				_1356_v31 = _dafny.SeqOf(_dafny.ArrayCastTo((_1354_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1354_v29), 0))).Int())), _1326_v9, _1326_v9)
				var _rhs194 _dafny.Array = _dafny.ArrayCastTo((_1354_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1354_v29), 0))).Int()))
				_ = _rhs194
				var _rhs195 _dafny.Sequence = _1355_v30
				_ = _rhs195
				var _rhs196 _dafny.Int = _dafny.IntOfInt64(812)
				_ = _rhs196
				var _rhs197 _dafny.Array = (_1356_v31).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1356_v31).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs197
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				_1326_v9 = _rhs194
				_lhs179.F6 = _rhs195
				_lhs180.F16 = _rhs196
				_1326_v9 = _rhs197
				(globalState).F16 = (_this).Fm1(globalState)
				var _1357_v32 _dafny.Array
				_ = _1357_v32
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(4))
				_ = _nw203
				_1357_v32 = _nw203
				var _1358_v33 _dafny.Array
				_ = _1358_v33
				var _nw204 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(4))
				_ = _nw204
				_1358_v33 = _nw204
				var _1359_v34 _dafny.Sequence
				_ = _1359_v34
				_1359_v34 = _dafny.SeqOf(_1358_v33, _1358_v33)
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1357_v32), 0))
				_ = _index186
				(_1357_v32).ArraySet1(_1359_v34, (_index186).Int())
				var _1360_v35 _dafny.MultiSet
				_ = _1360_v35
				_1360_v35 = _dafny.MultiSetOf(p3, p3)
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1357_v32), 0))
				_ = _index187
				var _rhs198 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1359_v34, _1359_v34)
				_ = _rhs198
				var _rhs199 _dafny.Int = _1333_i5
				_ = _rhs199
				var _rhs200 bool = (_1360_v35).IsSubsetOf(_1360_v35)
				_ = _rhs200
				var _lhs181 _dafny.Array = _1357_v32
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1357_v32), 0))
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				(_lhs181).ArraySet1(_rhs198, (_lhs182).Int())
				_lhs183.F16 = _rhs199
				_lhs184.F15 = _rhs200
				var _1361_v36 *C1
				_ = _1361_v36
				var _nw205 *C1 = New_C1_()
				_ = _nw205
				_nw205.Ctor__(p2, _1333_i5, ((_dafny.Zero).Minus(_dafny.IntOfInt64(-125))).Cmp((p3).Cardinality()) == 0, (_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool))
				_1361_v36 = _nw205
			}
		}
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
		_ = _index188
		(_1324_v8).ArraySet1((_this).F26(), (_index188).Int())
		var _1362_v37 _dafny.Map
		_ = _1362_v37
		_1362_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(946), true)
		(globalState).F24 = ((_1362_v37).Cardinality()).Cmp(p2) == 0
		if (_this).F26() {
			var _1363_v38 _dafny.Map
			_ = _1363_v38
			_1363_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3).Cardinality(), p2)
			var _1364_v39 _dafny.MultiSet
			_ = _1364_v39
			_1364_v39 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(p2, (_1363_v38).Cardinality()), p2)
			_1364_v39 = (_1364_v39).Difference((_1364_v39).Difference(_1364_v39))
			(globalState).F16 = _dafny.IntOfUint32((_1327_v10).Cardinality())
			var _1365_v40 _dafny.Sequence
			_ = _1365_v40
			_1365_v40 = _dafny.SeqOf((_this).F26())
			_1327_v10 = _dafny.SeqOf(_dafny.IntOfUint32((_1365_v40).Cardinality()), p2, (_dafny.IntOfInt64(599)).Minus(p2))
			(globalState).F22 = p1
			var _1366_v41 _dafny.Array
			_ = _1366_v41
			var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw206
			_1366_v41 = _nw206
			_1366_v41 = _1366_v41
		} else {
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))
			_ = _index189
			(_1326_v9).ArraySet1(p2, (_index189).Int())
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))
			_ = _index190
			(_1326_v9).ArraySet1(p2, (_index190).Int())
			var _1367_v42 D0
			_ = _1367_v42
			_1367_v42 = Companion_D0_.Create_DC0_(p0)
			var _1368_v43 _dafny.Sequence
			_ = _1368_v43
			_1368_v43 = _dafny.SeqOf(true)
			var _1369_v44 _dafny.MultiSet
			_ = _1369_v44
			_1369_v44 = _dafny.MultiSetOf(p1)
			var _1370_v45 bool
			_ = _1370_v45
			var _out43 bool
			_ = _out43
			_out43 = (_this).M8(_1367_v42, (_1368_v43).Select((Companion_Default___.SafeIndex((_this).Fm1(globalState), _dafny.IntOfUint32((_1368_v43).Cardinality()))).Uint32()).(bool), _1369_v44, (_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), globalState)
			_1370_v45 = _out43
			var _1371_v46 _dafny.MultiSet
			_ = _1371_v46
			_1371_v46 = _dafny.MultiSetOf(_1324_v8, _1324_v8)
			var _1372_v47 _dafny.Sequence
			_ = _1372_v47
			_1372_v47 = _dafny.UnicodeSeqOfUtf8Bytes("nhfpskf")
			var _1373_v48 _dafny.Map
			_ = _1373_v48
			_1373_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1371_v46).Intersection(_1371_v46)).Cardinality(), (func() _dafny.Sequence {
				if (_this).F27() {
					return _1372_v47
				}
				return _1372_v47
			})())
			_1373_v48 = (_1373_v48).Update((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("t"), Companion_Default___.Fm22(p2, (_this).Fm2(p0, (_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), globalState), globalState)))
			var _1374_v49 _dafny.Map
			_ = _1374_v49
			_1374_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int))
			var _1375_v50 D7
			_ = _1375_v50
			_1375_v50 = Companion_D7_.Create_DC19_((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _1369_v44, (_this).F27(), _1374_v49)
			if (_1375_v50).Dtor_cf38() {
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))
				_ = _index191
				(_1326_v9).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), (_index191).Int())
				(globalState).F12 = p2
				var _1376_v51 _dafny.Map
				_ = _1376_v51
				_1376_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _dafny.SeqOf((_1375_v50).Dtor_cf38()))
				var _1377_v52 _dafny.CodePoint
				_ = _1377_v52
				_1377_v52 = _dafny.CodePoint('h')
				_1376_v51 = (_1376_v51).Update((_this).Fm1(globalState), Companion_Default___.Fm11((_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool), (_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool), _1377_v52, globalState))
				var _1378_v53 _dafny.Array
				_ = _1378_v53
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw207
				_1378_v53 = _nw207
				var _1379_v54 _dafny.Array
				_ = _1379_v54
				var _nwElement0_45 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("fbnprm")
				_ = _nwElement0_45
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(25))
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_45, 0)
				(_nw208).ArraySet1(_1372_v47, 1)
				(_nw208).ArraySet1(_1372_v47, 2)
				(_nw208).ArraySet1(_1372_v47, 3)
				(_nw208).ArraySet1(_1372_v47, 4)
				(_nw208).ArraySet1(_1372_v47, 5)
				(_nw208).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("c"), 6)
				(_nw208).ArraySet1(_1372_v47, 7)
				(_nw208).ArraySet1(_1372_v47, 8)
				(_nw208).ArraySet1(_1372_v47, 9)
				(_nw208).ArraySet1(_1372_v47, 10)
				(_nw208).ArraySet1(Companion_Default___.Fm22(p2, _1370_v45, globalState), 11)
				(_nw208).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hluwry"), 12)
				(_nw208).ArraySet1(_1372_v47, 13)
				(_nw208).ArraySet1(_1372_v47, 14)
				(_nw208).ArraySet1(_1372_v47, 15)
				(_nw208).ArraySet1(_1372_v47, 16)
				(_nw208).ArraySet1(_1372_v47, 17)
				(_nw208).ArraySet1(_1372_v47, 18)
				(_nw208).ArraySet1(_1372_v47, 19)
				(_nw208).ArraySet1(_1372_v47, 20)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("pkxwgoox"), (Companion_Default___.SafeIndex((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pkxwgoox")).Cardinality()))).Uint32(), _1377_v52), 21)
				(_nw208).ArraySet1(_1372_v47, 22)
				(_nw208).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(773))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg81 _dafny.Int) interface{} {
						return coer81(arg81)
					}
				}((func(_1380_v52 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1381_i8 _dafny.Int) _dafny.CodePoint {
						return _1380_v52
					}
				})(_1377_v52))), 23)
				(_nw208).ArraySet1(_1372_v47, 24)
				_1379_v54 = _nw208
				var _1382_v55 _dafny.Sequence
				_ = _1382_v55
				_1382_v55 = _dafny.SeqOf(_1379_v54, _1379_v54)
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1378_v53), 0))
				_ = _index192
				(_1378_v53).ArraySet1((func() _dafny.Array {
					if (_this).Fm2(true, (_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), globalState) {
						return _1379_v54
					}
					return (_1382_v55).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1382_v55).Cardinality()))).Uint32()).(_dafny.Array)
				})(), (_index192).Int())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1378_v53), 0))
				_ = _index193
				(_1378_v53).ArraySet1(_1379_v54, (_index193).Int())
				var _1383_v56 _dafny.Map
				_ = _1383_v56
				_1383_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.SeqOf((_this).F27(), _1370_v45))
				var _1384_v57 D5
				_ = _1384_v57
				_1384_v57 = Companion_D5_.Create_DC13_(_1383_v56)
				_1384_v57 = _1384_v57
			} else {
				var _1385_v58 _dafny.CodePoint
				_ = _1385_v58
				_1385_v58 = _dafny.CodePoint('a')
				var _1386_v59 _dafny.Array
				_ = _1386_v59
				var _nwElement0_46 _dafny.Int = p2
				_ = _nwElement0_46
				var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(25))
				_ = _nw209
				(_nw209).ArraySet1(_nwElement0_46, 0)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 1)
				(_nw209).ArraySet1(p2, 2)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 3)
				(_nw209).ArraySet1(p2, 4)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 5)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 6)
				(_nw209).ArraySet1((_dafny.Zero).Minus((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int)), 7)
				(_nw209).ArraySet1(_dafny.IntOfUint32((_1327_v10).Cardinality()), 8)
				(_nw209).ArraySet1(p2, 9)
				(_nw209).ArraySet1(p2, 10)
				(_nw209).ArraySet1(p2, 11)
				(_nw209).ArraySet1(_dafny.IntOfInt64(606), 12)
				(_nw209).ArraySet1(p2, 13)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 14)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 15)
				(_nw209).ArraySet1(p2, 16)
				(_nw209).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1385_v58)).Cardinality()), 17)
				(_nw209).ArraySet1(p2, 18)
				(_nw209).ArraySet1(p2, 19)
				(_nw209).ArraySet1((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), 20)
				(_nw209).ArraySet1(p2, 21)
				(_nw209).ArraySet1(_dafny.IntOfUint32((_1372_v47).Cardinality()), 22)
				(_nw209).ArraySet1(p2, 23)
				(_nw209).ArraySet1(p2, 24)
				_1386_v59 = _nw209
				var _1387_v60 _dafny.Map
				_ = _1387_v60
				_1387_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1386_v59, _dafny.IntOfUint32((_1372_v47).Cardinality()))
				var _1388_v61 _dafny.Map
				_ = _1388_v61
				_1388_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1369_v44, _dafny.SetOf((_this).F27()))
				var _1389_v62 D5
				_ = _1389_v62
				_1389_v62 = Companion_D5_.Create_DC14_(((func() _dafny.Set {
					if (_1388_v61).Contains(_1369_v44) {
						return (_1388_v61).Get(_1369_v44).(_dafny.Set)
					}
					return _dafny.SetOf((_this).F26(), (_this).F26())
				})()).Cardinality(), p2, (_1368_v43).Select((Companion_Default___.SafeIndex((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1368_v43).Cardinality()))).Uint32()).(bool))
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))
				_ = _index194
				var _rhs201 _dafny.Int = (func() _dafny.Int {
					if (_this).F27() {
						return ((_1387_v60).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1386_v59, (_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int)))).Cardinality()
					}
					return p2
				})()
				_ = _rhs201
				var _rhs202 _dafny.Int = (_1389_v62).Dtor_cf26()
				_ = _rhs202
				var _rhs203 bool = (func() bool {
					if !((p2).Cmp(p2) <= 0) {
						return _1370_v45
					}
					return (_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool)
				})()
				_ = _rhs203
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				var _lhs186 _dafny.Array = _1326_v9
				_ = _lhs186
				var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))
				_ = _lhs187
				var _lhs188 *GlobalState = globalState
				_ = _lhs188
				_lhs185.F12 = _rhs201
				(_lhs186).ArraySet1(_rhs202, (_lhs187).Int())
				_lhs188.F24 = _rhs203
				var _1390_v63 *C0
				_ = _1390_v63
				var _nw210 *C0 = New_C0_()
				_ = _nw210
				_nw210.Ctor__((_this).F26())
				_1390_v63 = _nw210
				var _1391_v64 _dafny.Set
				_ = _1391_v64
				_1391_v64 = _dafny.SetOf((_this).F26())
				var _1392_v65 _dafny.Map
				_ = _1392_v65
				_1392_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(262))).Uint32(), func(coer82 func(_dafny.Int) D3) func(_dafny.Int) interface{} {
					return func(arg82 _dafny.Int) interface{} {
						return coer82(arg82)
					}
				}((func(_1393_v58 _dafny.CodePoint) func(_dafny.Int) D3 {
					return func(_1394_i9 _dafny.Int) D3 {
						return Companion_D3_.Create_DC11_(_1393_v58)
					}
				})(_1385_v58))), _1390_v63)
				var _1395_v66 D3
				_ = _1395_v66
				_1395_v66 = Companion_D3_.Create_DC11_(_1385_v58)
				var _rhs204 bool = ((func() _dafny.Set {
					if (_this).F27() {
						return (_1391_v64)
					}
					return _1391_v64
				})()).Contains((_1390_v63).F32())
				_ = _rhs204
				var _rhs205 *C0 = (func() *C0 {
					if (_1392_v65).Contains(_dafny.SeqOf(_1395_v66, Companion_D3_.Create_DC11_(_1385_v58), _1395_v66, _1395_v66)) {
						return (_1392_v65).Get(_dafny.SeqOf(_1395_v66, Companion_D3_.Create_DC11_(_1385_v58), _1395_v66, _1395_v66)).(*C0)
					}
					return _1390_v63
				})()
				_ = _rhs205
				var _lhs189 *GlobalState = globalState
				_ = _lhs189
				_lhs189.F2 = _rhs204
				_1390_v63 = _rhs205
				(globalState).F6 = _1372_v47
				var _1396_v67 _dafny.Array
				_ = _1396_v67
				var _nwElement0_47 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mxwll"), _1372_v47)
				_ = _nwElement0_47
				var _nw211 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(18))
				_ = _nw211
				(_nw211).ArraySet1(_nwElement0_47, 0)
				(_nw211).ArraySet1(_1372_v47, 1)
				(_nw211).ArraySet1(_1372_v47, 2)
				(_nw211).ArraySet1(_dafny.Companion_Sequence_.Update(_1372_v47, (Companion_Default___.SafeIndex((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1372_v47).Cardinality()))).Uint32(), _1385_v58), 3)
				(_nw211).ArraySet1(_1372_v47, 4)
				(_nw211).ArraySet1(_1372_v47, 5)
				(_nw211).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1372_v47, _1372_v47), 6)
				(_nw211).ArraySet1(_1372_v47, 7)
				(_nw211).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lawmo"), 8)
				(_nw211).ArraySet1(_1372_v47, 9)
				(_nw211).ArraySet1(Companion_Default___.Fm22((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), p0, globalState), 10)
				(_nw211).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1397_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1398_i10 _dafny.Int) _dafny.CodePoint {
						return _1397_v58
					}
				})(_1385_v58))), 11)
				(_nw211).ArraySet1(_1372_v47, 12)
				(_nw211).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}(func(_1399_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				})), 13)
				(_nw211).ArraySet1(_1372_v47, 14)
				(_nw211).ArraySet1(_1372_v47, 15)
				(_nw211).ArraySet1(_1372_v47, 16)
				(_nw211).ArraySet1(_1372_v47, 17)
				_1396_v67 = _nw211
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_39
				var _nw212 _dafny.Array
				_ = _nw212
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw212 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.Sequence = (func(_1400_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
						return func(_1401_i12 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg85 _dafny.Int) interface{} {
									return coer85(arg85)
								}
							}((func(_1402_v58 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_1403_i13 _dafny.Int) _dafny.CodePoint {
									return _1402_v58
								}
							})(_1400_v58)))
						}
					})(_1385_v58)
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw212 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw212).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw212).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1396_v67 = _nw212
				var _1404_v68 _dafny.Array
				_ = _1404_v68
				_1404_v68 = _1324_v8
				var _1405_v69 _dafny.Sequence
				_ = _1405_v69
				_1405_v69 = _dafny.SeqOf(_1324_v8, (_1404_v68), _1324_v8, _1324_v8)
				var _1406_v70 _dafny.Map
				_ = _1406_v70
				_1406_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("hnkb"), _dafny.SeqOf(_1324_v8))
				var _1407_v71 _dafny.Array
				_ = _1407_v71
				var _nwElement0_48 _dafny.Sequence = _1405_v69
				_ = _nwElement0_48
				var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(5))
				_ = _nw213
				(_nw213).ArraySet1(_nwElement0_48, 0)
				(_nw213).ArraySet1(_1405_v69, 1)
				(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1405_v69, _1405_v69), 2)
				(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1405_v69, _1405_v69), 3)
				(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1405_v69, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_1406_v70).Contains(Companion_Default___.Fm22((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), (_1390_v63).Fm6(_1385_v58, (_1390_v63).F32(), globalState), globalState)) {
						return (_1406_v70).Get(Companion_Default___.Fm22((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), (_1390_v63).Fm6(_1385_v58, (_1390_v63).F32(), globalState), globalState)).(_dafny.Sequence)
					}
					return _1405_v69
				})(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1406_v70).Contains(Companion_Default___.Fm22((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), (_1390_v63).Fm6(_1385_v58, (_1390_v63).F32(), globalState), globalState)) {
						return (_1406_v70).Get(Companion_Default___.Fm22((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int), (_1390_v63).Fm6(_1385_v58, (_1390_v63).F32(), globalState), globalState)).(_dafny.Sequence)
					}
					return _1405_v69
				})()).Cardinality()))).Uint32(), _1324_v8)), 4)
				_1407_v71 = _nw213
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1407_v71), 0))
				_ = _index195
				(_1407_v71).ArraySet1(_1405_v69, (_index195).Int())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1407_v71), 0))
				_ = _index196
				(_1407_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1405_v69, _dafny.Companion_Sequence_.Concatenate(_1405_v69, _1405_v69)), (_index196).Int())
			}
			if p1 {
				var _1408_v72 _dafny.Map
				_ = _1408_v72
				_1408_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1372_v47, _1324_v8)
				_1408_v72 = (_1408_v72).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}(func(_1409_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})), _1324_v8)
				var _1410_v73 *C3
				_ = _1410_v73
				var _nw214 *C3 = New_C3_()
				_ = _nw214
				_nw214.Ctor__(_1328_v11, (p1) == (_1370_v45), !(_1370_v45) || (_1370_v45))
				_1410_v73 = _nw214
				var _1411_v74 bool
				_ = _1411_v74
				var _out44 bool
				_ = _out44
				_out44 = (_this).M8(_1367_v42, (_this).F26(), _dafny.MultiSetOf((_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool)), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(553), _dafny.IntOfUint32((_1372_v47).Cardinality())), globalState)
				_1411_v74 = _out44
				(globalState).F12 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int)), p2)
				var _1412_v75 _dafny.CodePoint
				_ = _1412_v75
				_1412_v75 = _dafny.CodePoint('h')
				(globalState).F19 = _1412_v75
			} else {
				var _1413_v76 _dafny.Map
				_ = _1413_v76
				_1413_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1362_v37, p1)
				var _1414_v77 *C1
				_ = _1414_v77
				var _nw215 *C1 = New_C1_()
				_ = _nw215
				_nw215.Ctor__((_1413_v76).Cardinality(), p2, p0, false)
				_1414_v77 = _nw215
				var _1415_v78 _dafny.Map
				_ = _1415_v78
				_1415_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1414_v77, p3)
				var _1416_v79 _dafny.Map
				_ = _1416_v79
				_1416_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() _dafny.MultiSet {
					if (_1415_v78).Contains(_1414_v77) {
						return (_1415_v78).Get(_1414_v77).(_dafny.MultiSet)
					}
					return _dafny.MultiSetOf((_1414_v77).F36())
				})())
				var _1417_v80 *C0
				_ = _1417_v80
				var _nw216 *C0 = New_C0_()
				_ = _nw216
				_nw216.Ctor__(true)
				_1417_v80 = _nw216
				var _1418_v81 _dafny.Sequence
				_ = _1418_v81
				_1418_v81 = _dafny.SeqOf(_1417_v80, _1417_v80)
				var _1419_v82 _dafny.Map
				_ = _1419_v82
				_1419_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1416_v79, _dafny.Companion_Sequence_.Concatenate(_1418_v81, _1418_v81))
				_1419_v82 = (_1419_v82).Update(_1416_v79, _1418_v81)
				_1327_v10 = _dafny.Companion_Sequence_.Concatenate(_1327_v10, _1327_v10)
				var _1420_v83 _dafny.Array
				_ = _1420_v83
				var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw217
				_1420_v83 = _nw217
				var _1421_v84 D6
				_ = _1421_v84
				_1421_v84 = Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("drgpsi"))
				var _1422_v85 _dafny.CodePoint
				_ = _1422_v85
				_1422_v85 = _dafny.CodePoint('c')
				var _1423_v86 _dafny.Map
				_ = _1423_v86
				_1423_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1421_v84, _1422_v85)
				var _1424_v87 _dafny.Set
				_ = _1424_v87
				_1424_v87 = _dafny.SetOf(_1423_v86)
				var _1425_v89 D3
				_ = _1425_v89
				_1425_v89 = Companion_D3_.Create_DC10_(_1420_v83)
				var _1426_v90 *C4
				_ = _1426_v90
				var _nw218 *C4 = New_C4_()
				_ = _nw218
				_nw218.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1425_v89, (_dafny.Zero).Minus((_1326_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1326_v9), 0))).Int()).(_dafny.Int))), p0, true)
				_1426_v90 = _nw218
				var _1427_v91 _dafny.Sequence
				_ = _1427_v91
				_1427_v91 = _dafny.SeqOf(_1426_v90, _1426_v90, _1426_v90)
				var _1428_v92 _dafny.Map
				_ = _1428_v92
				_1428_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1427_v91).Cardinality()), _1414_v77.F37)
				var _1429_v93 D2
				_ = _1429_v93
				_1429_v93 = Companion_D2_.Create_DC6_(func() _dafny.Map {
					var _coll61 = _dafny.NewMapBuilder()
					_ = _coll61
					for _iter61 := _dafny.Iterate((_1428_v92).Keys().Elements()); ; {
						_compr_61, _ok61 := _iter61()
						if !_ok61 {
							break
						}
						var _1430_v88 _dafny.Int
						_1430_v88 = interface{}(_compr_61).(_dafny.Int)
						if (_1428_v92).Contains(_1430_v88) {
							_coll61.Add((_1430_v88).Times(_dafny.IntOfInt64(-800)), p2)
						}
					}
					return _coll61.ToMap()
				}())
				var _1431_v94 _dafny.Map
				_ = _1431_v94
				_1431_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1327_v10, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(719))).Uint32(), func(coer87 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1432_v77 *C1, _1433_v9 _dafny.Array) func(_dafny.Int) D2 {
					return func(_1434_i15 _dafny.Int) D2 {
						return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1432_v77.F37, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_1433_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1433_v9), 0))).Int()).(_dafny.Int)))).Cardinality())))
					}
				})(_1414_v77, _1326_v9))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1372_v47).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(719))).Uint32(), func(coer88 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}((func(_1435_v77 *C1, _1436_v9 _dafny.Array) func(_dafny.Int) D2 {
					return func(_1437_i15 _dafny.Int) D2 {
						return Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1435_v77.F37, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_1436_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_1436_v9), 0))).Int()).(_dafny.Int)))).Cardinality())))
					}
				})(_1414_v77, _1326_v9)))).Cardinality()))).Uint32(), _1429_v93))
				var _1438_v95 D10
				_ = _1438_v95
				_1438_v95 = Companion_D10_.Create_DC23_(_1431_v94)
				var _rhs206 _dafny.Int = _1414_v77.F37
				_ = _rhs206
				var _rhs207 _dafny.Array = _1420_v83
				_ = _rhs207
				var _rhs208 _dafny.Array = _1324_v8
				_ = _rhs208
				var _rhs209 _dafny.Set = _dafny.SetOf((Companion_Default___.Fm37(_1370_v45, (_1414_v77).Fm1(globalState), p2, _1438_v95, globalState)).Merge(_1423_v86))
				_ = _rhs209
				var _rhs210 _dafny.Int = _dafny.IntOfInt64(569)
				_ = _rhs210
				var _lhs190 *GlobalState = globalState
				_ = _lhs190
				var _lhs191 *GlobalState = globalState
				_ = _lhs191
				_lhs190.F16 = _rhs206
				_1420_v83 = _rhs207
				_1324_v8 = _rhs208
				_1424_v87 = _rhs209
				_lhs191.F16 = _rhs210
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
				_ = _index197
				(_1324_v8).ArraySet1(((_1426_v90).Fm1(globalState)).Cmp((_dafny.IntOfUint32((_dafny.SeqOf((_1324_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))).Int()).(bool))).Cardinality())).Minus((_1414_v77).F36())) <= 0, (_index197).Int())
				var _1439_v96 _dafny.Set
				_ = _1439_v96
				_1439_v96 = _dafny.SetOf(_dafny.IntOfUint32((_1372_v47).Cardinality()))
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1324_v8), 0))
				_ = _index198
				(_1324_v8).ArraySet1((((_1439_v96).Cardinality()).Minus(p2)).Cmp(p2) < 0, (_index198).Int())
			}
		}
	}
}
func (_this *C5) M8(p0 D0, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1440_v0 _dafny.Array
		_ = _1440_v0
		var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
		_ = _nw219
		_1440_v0 = _nw219
		var _1441_v1 _dafny.CodePoint
		_ = _1441_v1
		_1441_v1 = _dafny.CodePoint('h')
		var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))
		_ = _index199
		(_1440_v0).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qsnm"), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qsnm")).Cardinality()))).Uint32(), _1441_v1), (_index199).Int())
		var _1442_v2 _dafny.Sequence
		_ = _1442_v2
		_1442_v2 = _dafny.UnicodeSeqOfUtf8Bytes("yxuu")
		var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))
		_ = _index200
		(_1440_v0).ArraySet1(_1442_v2, (_index200).Int())
		var _1443_v3 _dafny.Sequence
		_ = _1443_v3
		_1443_v3 = _dafny.SeqOf((_this).Fm2((_this).F26(), p3, globalState), p1)
		(globalState).F12 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26()), _1443_v3)).Cardinality())
		(globalState).F16 = (_dafny.Zero).Minus((_dafny.Zero).Minus(p3))
		var _1444_v4 _dafny.Map
		_ = _1444_v4
		_1444_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p3, p3)).Cardinality()), _dafny.IntOfInt64(427))
		var _1445_v5 D2
		_ = _1445_v5
		_1445_v5 = Companion_D2_.Create_DC6_(_1444_v4)
		var _1446_v7 _dafny.Set
		_ = _1446_v7
		_1446_v7 = _dafny.SetOf(p3)
		var _pat_let_tv24 = _1446_v7
		_ = _pat_let_tv24
		var _pat_let_tv25 = _1446_v7
		_ = _pat_let_tv25
		var _pat_let_tv26 = p3
		_ = _pat_let_tv26
		var _1447_v8 _dafny.Array
		_ = _1447_v8
		var _nwElement0_49 D2 = _1445_v5
		_ = _nwElement0_49
		var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(18))
		_ = _nw220
		(_nw220).ArraySet1(_nwElement0_49, 0)
		(_nw220).ArraySet1(_1445_v5, 1)
		(_nw220).ArraySet1(_1445_v5, 2)
		(_nw220).ArraySet1(_1445_v5, 3)
		(_nw220).ArraySet1(func(_pat_let17_0 D2) D2 {
			return func(_1448_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let18_0 _dafny.Map) D2 {
					return func(_1450_dt__update_hcf13_h0 _dafny.Map) D2 {
						return Companion_D2_.Create_DC6_(_1450_dt__update_hcf13_h0)
					}(_pat_let18_0)
				}(func() _dafny.Map {
					var _coll62 = _dafny.NewMapBuilder()
					_ = _coll62
					for _iter62 := _dafny.Iterate((_pat_let_tv24).Elements()); ; {
						_compr_62, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1449_v6 _dafny.Int
						_1449_v6 = interface{}(_compr_62).(_dafny.Int)
						if (_pat_let_tv25).Contains(_1449_v6) {
							_coll62.Add(Companion_Default___.SafeDivisionInt(_1449_v6, _dafny.IntOfInt64(270)), _pat_let_tv26)
						}
					}
					return _coll62.ToMap()
				}())
			}(_pat_let17_0)
		}(_1445_v5), 4)
		(_nw220).ArraySet1(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)), 5)
		(_nw220).ArraySet1(Companion_D2_.Create_DC6_(_1444_v4), 6)
		(_nw220).ArraySet1(_1445_v5, 7)
		(_nw220).ArraySet1(_1445_v5, 8)
		(_nw220).ArraySet1(_1445_v5, 9)
		(_nw220).ArraySet1(_1445_v5, 10)
		(_nw220).ArraySet1(Companion_D2_.Create_DC6_(_1444_v4), 11)
		(_nw220).ArraySet1(_1445_v5, 12)
		(_nw220).ArraySet1(_1445_v5, 13)
		(_nw220).ArraySet1(_1445_v5, 14)
		(_nw220).ArraySet1(Companion_D2_.Create_DC6_(_1444_v4), 15)
		(_nw220).ArraySet1(_1445_v5, 16)
		(_nw220).ArraySet1(Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)), 17)
		_1447_v8 = _nw220
		var _1451_v9 _dafny.Map
		_ = _1451_v9
		_1451_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(999))).Uint32(), func(coer89 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg89 _dafny.Int) interface{} {
				return coer89(arg89)
			}
		}(func(_1452_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), (_this).Fm5(_1443_v3, _1443_v3, p0, !(false), globalState)), _1447_v8)
		_1451_v9 = (_1451_v9).Update((_this).F26(), _1447_v8)
		var _hi12 _dafny.Int = (_dafny.Zero).Minus(p3)
		_ = _hi12
		for _1453_i1 := p3; _1453_i1.Cmp(_hi12) < 0; _1453_i1 = _1453_i1.Plus(_dafny.One) {
			var _1454_v10 _dafny.Array
			_ = _1454_v10
			var _nw221 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw221
			_1454_v10 = _nw221
			var _1455_v11 _dafny.Map
			_ = _1455_v11
			_1455_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1446_v7)
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))
			_ = _index201
			(_1454_v10).ArraySet1(!((_dafny.SetOf(p3, _1453_i1)).IsSubsetOf((func() _dafny.Set {
				if (_1455_v11).Contains((_this).F27()) {
					return (_1455_v11).Get((_this).F27()).(_dafny.Set)
				}
				return _1446_v7
			})())), (_index201).Int())
			var _1456_v12 _dafny.Map
			_ = _1456_v12
			_1456_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
			var _1457_v13 _dafny.MultiSet
			_ = _1457_v13
			_1457_v13 = _dafny.MultiSetOf(_1454_v10)
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))
			_ = _index202
			var _rhs211 bool = (Companion_Default___.SafeModuloInt((_1456_v12).Cardinality(), _1453_i1)).Cmp(_dafny.IntOfInt64(297)) != 0
			_ = _rhs211
			var _rhs212 bool = ((_1457_v13).Difference(_dafny.MultiSetOf(_1454_v10, _1454_v10))).IsProperSubsetOf(_1457_v13)
			_ = _rhs212
			var _lhs192 *GlobalState = globalState
			_ = _lhs192
			var _lhs193 _dafny.Array = _1454_v10
			_ = _lhs193
			var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))
			_ = _lhs194
			_lhs192.F22 = _rhs211
			(_lhs193).ArraySet1(_rhs212, (_lhs194).Int())
			(globalState).F16 = p3
			var _1458_v14 _dafny.MultiSet
			_ = _1458_v14
			_1458_v14 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm25((_1454_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))).Int()).(bool), (_1454_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))).Int()).(bool), _dafny.CodePoint('q'), globalState), _1443_v3))
			_1458_v14 = (_1458_v14).Union(_dafny.MultiSetOf(_1443_v3, _1443_v3, _dafny.SeqOf((_1454_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))).Int()).(bool), (_this).Fm2(p1, p3, globalState)), _dafny.SeqOf((_this).F26(), (_1454_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1454_v10), 0))).Int()).(bool), (_this).F26())))
			(globalState).F16 = _1453_i1
		}
		var _1459_v15 _dafny.Sequence
		_ = _1459_v15
		_1459_v15 = _dafny.SeqOf((_this).Fm1(globalState))
		var _1460_i2 _dafny.Int
		_ = _1460_i2
		_1460_i2 = _dafny.Zero
		{
			for (func() bool {
				if !((_1443_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1459_v15, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1459_v15).Cardinality()))).Uint32(), _dafny.IntOfInt64(60))).Cardinality()), _dafny.IntOfUint32((_1443_v3).Cardinality()))).Uint32()).(bool)) {
					return (_this).F26()
				}
				return false
			})() {
				{
					if (_1460_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1460_i2 = (_1460_i2).Plus(_dafny.One)
					var _1461_v16 D0
					_ = _1461_v16
					_1461_v16 = Companion_D0_.Create_DC1_((_this).F27(), (_this).F26(), _dafny.IntOfUint32((_1443_v3).Cardinality()), p1)
					var _1462_v17 D0
					_ = _1462_v17
					_1462_v17 = Companion_D0_.Create_DC2_(_1461_v16)
					var _1463_v18 D0
					_ = _1463_v18
					_1463_v18 = Companion_D0_.Create_DC2_(_1462_v17)
					var _source22 D0 = Companion_D0_.Create_DC2_(_1462_v17)
					_ = _source22
					if _source22.Is_DC1() {
						var _1464___mcc_h0 bool = _source22.Get_().(D0_DC1).Cf1
						_ = _1464___mcc_h0
						var _1465___mcc_h1 bool = _source22.Get_().(D0_DC1).Cf2
						_ = _1465___mcc_h1
						var _1466___mcc_h2 _dafny.Int = _source22.Get_().(D0_DC1).Cf3
						_ = _1466___mcc_h2
						var _1467___mcc_h3 bool = _source22.Get_().(D0_DC1).Cf4
						_ = _1467___mcc_h3
						var _1468_cf4 bool = _1467___mcc_h3
						_ = _1468_cf4
						var _1469_cf3 _dafny.Int = _1466___mcc_h2
						_ = _1469_cf3
						var _1470_cf2 bool = _1465___mcc_h1
						_ = _1470_cf2
						var _1471_cf1 bool = _1464___mcc_h0
						_ = _1471_cf1
						var _1472_v19 _dafny.Array
						_ = _1472_v19
						var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
						_ = _nw222
						_1472_v19 = _nw222
						var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_1472_v19), 0))
						_ = _index203
						(_1472_v19).ArraySet1CodePoint(_1441_v1, (_index203).Int())
						var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(419), _dafny.ArrayLen((_1472_v19), 0))
						_ = _index204
						(_1472_v19).ArraySet1CodePoint(_1441_v1, (_index204).Int())
						var _1473_v20 _dafny.Map
						_ = _1473_v20
						_1473_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1469_cf3, !(_1470_cf2) || (true))
						_1473_v20 = (_1473_v20).Update(_1469_cf3, (_1469_cf3).Cmp(p3) <= 0)
						(globalState).F24 = (func() bool {
							if (_this).Fm2(p1, p3, globalState) {
								return _1468_cf4
							}
							return true
						})()
						var _1474_v21 _dafny.Array
						_ = _1474_v21
						var _len0_40 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_40
						var _nw223 _dafny.Array
						_ = _nw223
						if _len0_40.Cmp(_dafny.Zero) == 0 {
							_nw223 = _dafny.NewArray(_len0_40)
						} else {
							var _init40 func(_dafny.Int) bool = (func(_1475_cf2 bool) func(_dafny.Int) bool {
								return func(_1476_i3 _dafny.Int) bool {
									return _1475_cf2
								}
							})(_1470_cf2)
							_ = _init40
							var _element0_40 = _init40(_dafny.Zero)
							_ = _element0_40
							_nw223 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
							(_nw223).ArraySet1(_element0_40, 0)
							var _nativeLen0_40 = (_len0_40).Int()
							_ = _nativeLen0_40
							for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
								(_nw223).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
							}
						}
						_1474_v21 = _nw223
						var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1474_v21), 0))
						_ = _index205
						(_1474_v21).ArraySet1(_1471_cf1, (_index205).Int())
						var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1474_v21), 0))
						_ = _index206
						(_1474_v21).ArraySet1(!(_1471_cf1), (_index206).Int())
					} else if _source22.Is_DC0() {
						var _1477___mcc_h4 bool = _source22.Get_().(D0_DC0).Cf0
						_ = _1477___mcc_h4
						var _1478_cf0 bool = _1477___mcc_h4
						_ = _1478_cf0
						var _1479_v22 _dafny.Array
						_ = _1479_v22
						var _len0_41 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_41
						var _nw224 _dafny.Array
						_ = _nw224
						if _len0_41.Cmp(_dafny.Zero) == 0 {
							_nw224 = _dafny.NewArray(_len0_41)
						} else {
							var _init41 func(_dafny.Int) _dafny.Int = (func(_1480_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1481_i4 _dafny.Int) _dafny.Int {
									return Companion_Default___.SafeModuloInt(_1481_i4, _1480_p3)
								}
							})(p3)
							_ = _init41
							var _element0_41 = _init41(_dafny.Zero)
							_ = _element0_41
							_nw224 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
							(_nw224).ArraySet1(_element0_41, 0)
							var _nativeLen0_41 = (_len0_41).Int()
							_ = _nativeLen0_41
							for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
								(_nw224).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
							}
						}
						_1479_v22 = _nw224
						var _1482_v23 _dafny.Map
						_ = _1482_v23
						_1482_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _1479_v22)
						var _1483_v24 _dafny.Array
						_ = _1483_v24
						var _nwElement0_50 _dafny.Array = _1479_v22
						_ = _nwElement0_50
						var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(4))
						_ = _nw225
						(_nw225).ArraySet1(_nwElement0_50, 0)
						(_nw225).ArraySet1(_1479_v22, 1)
						(_nw225).ArraySet1((func() _dafny.Array {
							if (_1482_v23).Contains(!(true)) {
								return (_1482_v23).Get(!(true)).(_dafny.Array)
							}
							return _1479_v22
						})(), 2)
						(_nw225).ArraySet1(_1479_v22, 3)
						_1483_v24 = _nw225
						var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_1483_v24), 0))
						_ = _index207
						(_1483_v24).ArraySet1(_1479_v22, (_index207).Int())
						var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(175), _dafny.ArrayLen((_1483_v24), 0))
						_ = _index208
						(_1483_v24).ArraySet1(_1479_v22, (_index208).Int())
						var _1484_v25 _dafny.MultiSet
						_ = _1484_v25
						_1484_v25 = _dafny.MultiSetOf(p3, p3, p3, p3, p3)
						var _1485_v26 _dafny.Array
						_ = _1485_v26
						var _nwElement0_51 bool = (_dafny.IntOfInt64(728)).Cmp(p3) > 0
						_ = _nwElement0_51
						var _nw226 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(11))
						_ = _nw226
						(_nw226).ArraySet1(_nwElement0_51, 0)
						(_nw226).ArraySet1(_1478_cf0, 1)
						(_nw226).ArraySet1(_1478_cf0, 2)
						(_nw226).ArraySet1(_1478_cf0, 3)
						(_nw226).ArraySet1(!(!_dafny.Companion_Sequence_.Equal((_1440_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))).Int()).(_dafny.Sequence), _1442_v2)), 4)
						(_nw226).ArraySet1(!((_1484_v25).IsSubsetOf(_1484_v25)), 5)
						(_nw226).ArraySet1((_this).Fm2(false, p3, globalState), 6)
						(_nw226).ArraySet1((_this).F26(), 7)
						(_nw226).ArraySet1((_this).F27(), 8)
						(_nw226).ArraySet1(p1, 9)
						(_nw226).ArraySet1(true, 10)
						_1485_v26 = _nw226
						var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1485_v26), 0))
						_ = _index209
						(_1485_v26).ArraySet1(p1, (_index209).Int())
						var _1486_v27 D6
						_ = _1486_v27
						_1486_v27 = Companion_D6_.Create_DC16_(_dafny.UnicodeSeqOfUtf8Bytes("oqncb"))
						var _pat_let_tv27 = _1442_v2
						_ = _pat_let_tv27
						var _pat_let_tv28 = _1442_v2
						_ = _pat_let_tv28
						var _1487_v28 _dafny.Array
						_ = _1487_v28
						var _nwElement0_52 D6 = Companion_D6_.Create_DC16_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg90 _dafny.Int) interface{} {
								return coer90(arg90)
							}
						}(func(_1488_i5 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('b')
						})))
						_ = _nwElement0_52
						var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(28))
						_ = _nw227
						(_nw227).ArraySet1(_nwElement0_52, 0)
						(_nw227).ArraySet1(_1486_v27, 1)
						(_nw227).ArraySet1(_1486_v27, 2)
						(_nw227).ArraySet1(Companion_Default___.Fm38(true, p3, (_1440_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))).Int()).(_dafny.Sequence), globalState), 3)
						(_nw227).ArraySet1(_1486_v27, 4)
						(_nw227).ArraySet1(_1486_v27, 5)
						(_nw227).ArraySet1(_1486_v27, 6)
						(_nw227).ArraySet1(_1486_v27, 7)
						(_nw227).ArraySet1(_1486_v27, 8)
						(_nw227).ArraySet1(_1486_v27, 9)
						(_nw227).ArraySet1(func(_pat_let19_0 D6) D6 {
							return func(_1489_dt__update__tmp_h1 D6) D6 {
								return func(_pat_let20_0 _dafny.Sequence) D6 {
									return func(_1490_dt__update_hcf29_h0 _dafny.Sequence) D6 {
										return Companion_D6_.Create_DC16_(_1490_dt__update_hcf29_h0)
									}(_pat_let20_0)
								}(_pat_let_tv27)
							}(_pat_let19_0)
						}(Companion_D6_.Create_DC16_(_1442_v2)), 10)
						(_nw227).ArraySet1(_1486_v27, 11)
						(_nw227).ArraySet1(Companion_D6_.Create_DC16_(_1442_v2), 12)
						(_nw227).ArraySet1(func(_pat_let21_0 D6) D6 {
							return func(_1491_dt__update__tmp_h2 D6) D6 {
								return func(_pat_let22_0 _dafny.Sequence) D6 {
									return func(_1492_dt__update_hcf29_h1 _dafny.Sequence) D6 {
										return Companion_D6_.Create_DC16_(_1492_dt__update_hcf29_h1)
									}(_pat_let22_0)
								}(_pat_let_tv28)
							}(_pat_let21_0)
						}(_1486_v27), 13)
						(_nw227).ArraySet1(_1486_v27, 14)
						(_nw227).ArraySet1(_1486_v27, 15)
						(_nw227).ArraySet1(_1486_v27, 16)
						(_nw227).ArraySet1(_1486_v27, 17)
						(_nw227).ArraySet1(Companion_D6_.Create_DC16_(_1442_v2), 18)
						(_nw227).ArraySet1(_1486_v27, 19)
						(_nw227).ArraySet1(_1486_v27, 20)
						(_nw227).ArraySet1(_1486_v27, 21)
						(_nw227).ArraySet1(_1486_v27, 22)
						(_nw227).ArraySet1(_1486_v27, 23)
						(_nw227).ArraySet1(_1486_v27, 24)
						(_nw227).ArraySet1(Companion_D6_.Create_DC16_((_1440_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))).Int()).(_dafny.Sequence)), 25)
						(_nw227).ArraySet1(_1486_v27, 26)
						(_nw227).ArraySet1(_1486_v27, 27)
						_1487_v28 = _nw227
						var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1487_v28), 0))
						_ = _index210
						(_1487_v28).ArraySet1(_1486_v27, (_index210).Int())
						var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1485_v26), 0))
						_ = _index211
						var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1487_v28), 0))
						_ = _index212
						var _rhs213 bool = !((_this).F27())
						_ = _rhs213
						var _rhs214 D6 = _1486_v27
						_ = _rhs214
						var _lhs195 _dafny.Array = _1485_v26
						_ = _lhs195
						var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_1485_v26), 0))
						_ = _lhs196
						var _lhs197 _dafny.Array = _1487_v28
						_ = _lhs197
						var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1487_v28), 0))
						_ = _lhs198
						(_lhs195).ArraySet1(_rhs213, (_lhs196).Int())
						(_lhs197).ArraySet1(_rhs214, (_lhs198).Int())
						var _1493_v29 _dafny.Map
						_ = _1493_v29
						_1493_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(p3), _1441_v1)
						var _1494_v30 _dafny.Sequence
						_ = _1494_v30
						_1494_v30 = _dafny.SeqOf(_1493_v29)
						var _1495_v31 _dafny.Map
						_ = _1495_v31
						_1495_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1494_v30, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(68))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
							return func(arg91 _dafny.Int) interface{} {
								return coer91(arg91)
							}
						}((func(_1496_p3 _dafny.Int, _1497_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
							return func(_1498_i6 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(_1496_p3), _1497_v1)
							}
						})(p3, _1441_v1)))), _1443_v3)
						_1495_v31 = (_1495_v31).Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(func() _dafny.Map {
							var _coll63 = _dafny.NewMapBuilder()
							_ = _coll63
							for _iter63 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(p3), p3)).Keys().Elements()); ; {
								_compr_63, _ok63 := _iter63()
								if !_ok63 {
									break
								}
								var _1499_v32 D1
								_1499_v32 = interface{}(_compr_63).(D1)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC5_(p3), p3)).Contains(_1499_v32) {
									_coll63.Add(_1499_v32, _1441_v1)
								}
							}
							return _coll63.ToMap()
						}()), _1494_v30), _1443_v3)
						var _pat_let_tv29 = _1440_v0
						_ = _pat_let_tv29
						var _pat_let_tv30 = _1440_v0
						_ = _pat_let_tv30
						(globalState).F6 = (func(_pat_let23_0 D6) D6 {
							return func(_1500_dt__update__tmp_h3 D6) D6 {
								return func(_pat_let24_0 _dafny.Sequence) D6 {
									return func(_1501_dt__update_hcf29_h2 _dafny.Sequence) D6 {
										return Companion_D6_.Create_DC16_(_1501_dt__update_hcf29_h2)
									}(_pat_let24_0)
								}((_pat_let_tv30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_pat_let_tv29), 0))).Int()).(_dafny.Sequence))
							}(_pat_let23_0)
						}(_1486_v27)).Dtor_cf29()
					} else {
						var _1502___mcc_h5 D0 = _source22.Get_().(D0_DC2).Cf5
						_ = _1502___mcc_h5
						var _1503_cf5 D0 = _1502___mcc_h5
						_ = _1503_cf5
						var _1504_v33 _dafny.Array
						_ = _1504_v33
						var _len0_42 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_42
						var _nw228 _dafny.Array
						_ = _nw228
						if _len0_42.Cmp(_dafny.Zero) == 0 {
							_nw228 = _dafny.NewArray(_len0_42)
						} else {
							var _init42 func(_dafny.Int) bool = func(_1505_i7 _dafny.Int) bool {
								return (_this).F26()
							}
							_ = _init42
							var _element0_42 = _init42(_dafny.Zero)
							_ = _element0_42
							_nw228 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
							(_nw228).ArraySet1(_element0_42, 0)
							var _nativeLen0_42 = (_len0_42).Int()
							_ = _nativeLen0_42
							for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
								(_nw228).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
							}
						}
						_1504_v33 = _nw228
						var _1506_v34 _dafny.Map
						_ = _1506_v34
						_1506_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v33, p3)
						_1506_v34 = (_1506_v34).Update(_1504_v33, p3)
						var _1507_v35 _dafny.Map
						_ = _1507_v35
						_1507_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfUint32(((_1440_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()))
						var _1508_v36 _dafny.MultiSet
						_ = _1508_v36
						_1508_v36 = _dafny.MultiSetOf(_1507_v35, _1507_v35)
						var _rhs215 _dafny.Sequence = (_1440_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1440_v0), 0))).Int()).(_dafny.Sequence)
						_ = _rhs215
						var _rhs216 bool = p1
						_ = _rhs216
						var _rhs217 bool = ((_1508_v36).Cardinality()).Cmp(p3) > 0
						_ = _rhs217
						var _rhs218 bool = (_this).F27()
						_ = _rhs218
						var _rhs219 _dafny.CodePoint = _1441_v1
						_ = _rhs219
						var _lhs199 *GlobalState = globalState
						_ = _lhs199
						var _lhs200 *GlobalState = globalState
						_ = _lhs200
						var _lhs201 *GlobalState = globalState
						_ = _lhs201
						var _lhs202 *GlobalState = globalState
						_ = _lhs202
						var _lhs203 *GlobalState = globalState
						_ = _lhs203
						_lhs199.F6 = _rhs215
						_lhs200.F24 = _rhs216
						_lhs201.F15 = _rhs217
						_lhs202.F24 = _rhs218
						_lhs203.F19 = _rhs219
						(globalState).F2 = p1
						(globalState).F16 = p3
					}
					var _1509_v37 _dafny.CodePoint
					_ = _1509_v37
					var _1510_v38 _dafny.Sequence
					_ = _1510_v38
					var _1511_v39 bool
					_ = _1511_v39
					var _1512_v40 _dafny.Map
					_ = _1512_v40
					var _out45 _dafny.CodePoint
					_ = _out45
					var _out46 _dafny.Sequence
					_ = _out46
					var _out47 bool
					_ = _out47
					var _out48 _dafny.Map
					_ = _out48
					_out45, _out46, _out47, _out48 = Companion_Default___.M0(_1441_v1, globalState)
					_1509_v37 = _out45
					_1510_v38 = _out46
					_1511_v39 = _out47
					_1512_v40 = _out48
					var _1513_v41 _dafny.Array
					_ = _1513_v41
					var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
					_ = _nw229
					_1513_v41 = _nw229
					var _1514_v42 _dafny.Array
					_ = _1514_v42
					var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
					_ = _nw230
					_1514_v42 = _nw230
					var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1513_v41), 0))
					_ = _index213
					(_1513_v41).ArraySet1(_1514_v42, (_index213).Int())
					var _1515_v43 T2
					_ = _1515_v43
					var _nw231 *C1 = New_C1_()
					_ = _nw231
					_nw231.Ctor__(p3, p3, true, false)
					_1515_v43 = _nw231
					var _1516_v44 _dafny.Sequence
					_ = _1516_v44
					_1516_v44 = _dafny.SeqOf(_1515_v43)
					var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1513_v41), 0))
					_ = _index214
					var _rhs220 _dafny.Array = (func() _dafny.Array {
						if ((_this).F27()) || (p1) {
							return _1514_v42
						}
						return _1514_v42
					})()
					_ = _rhs220
					var _rhs221 _dafny.Sequence = _1516_v44
					_ = _rhs221
					var _rhs222 _dafny.Int = (_dafny.Zero).Minus((p3).Plus(p3))
					_ = _rhs222
					var _rhs223 bool = (_1515_v43).F26()
					_ = _rhs223
					var _lhs204 _dafny.Array = _1513_v41
					_ = _lhs204
					var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1513_v41), 0))
					_ = _lhs205
					var _lhs206 *GlobalState = globalState
					_ = _lhs206
					var _lhs207 *GlobalState = globalState
					_ = _lhs207
					(_lhs204).ArraySet1(_rhs220, (_lhs205).Int())
					_1516_v44 = _rhs221
					_lhs206.F12 = _rhs222
					_lhs207.F24 = _rhs223
					var _1517_v45 D3
					_ = _1517_v45
					_1517_v45 = Companion_D3_.Create_DC10_(_dafny.ArrayCastTo((_1513_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_1513_v41), 0))).Int())))
					var _1518_v46 _dafny.Map
					_ = _1518_v46
					_1518_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1515_v43).F26(), (_this).F26())
					var _1519_v47 _dafny.Map
					_ = _1519_v47
					_1519_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1517_v45, (_1518_v46).Cardinality())
					var _1520_v48 T1
					_ = _1520_v48
					var _nw232 *C4 = New_C4_()
					_ = _nw232
					_nw232.Ctor__(_1519_v47, (_1515_v43).F26(), (_this).F26())
					_1520_v48 = _nw232
					var _1521_v49 _dafny.Map
					_ = _1521_v49
					_1521_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1520_v48, !(false))
					_1521_v49 = (_1521_v49).Update(_1520_v48, true)
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		r0 = !((_1443_v3).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_1443_v3).Cardinality()))).Uint32()).(bool)) || ((_this).F27())
		return r0
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	F40 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this.F40 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__(f40 _dafny.Int) {
	{
		(_this).F40 = f40
	}
}
func (_this *C6) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return ((_this.F40).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(113))).Uint32(), func(coer92 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}(func(_1522_i0 _dafny.Int) _dafny.Int {
			return _this.F40
		}))).Cardinality()))).Minus(_this.F40)
	}
}
func (_this *C6) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Set {
			var _coll64 = _dafny.NewBuilder()
			_ = _coll64
			for _iter64 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F40, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Cardinality(), (_dafny.Zero).Minus(_this.F40)))).Elements()); ; {
				_compr_64, _ok64 := _iter64()
				if !_ok64 {
					break
				}
				var _1523_v0 _dafny.Int
				_1523_v0 = interface{}(_compr_64).(_dafny.Int)
				if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F40, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))).Cardinality(), (_dafny.Zero).Minus(_this.F40)))).Contains(_1523_v0) {
					_coll64.Add(Companion_Default___.SafeModuloInt(_1523_v0, (_dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), false, false, !(false), true))).Cardinality()))
				}
			}
			return _coll64.ToSet()
		}()).Difference(_dafny.SetOf((_dafny.Zero).Minus(_this.F40), (_dafny.MultiSetOf((_dafny.MultiSetOf(_this.F40)).Cardinality())).Cardinality(), (_dafny.MultiSetOf(true, !(true))).Cardinality(), _this.F40))).Equals((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(400))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_1524_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()), (_dafny.Zero).Minus(_this.F40), _this.F40)).Difference(_dafny.SetOf(_this.F40, _dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Set {
			var _coll65 = _dafny.NewBuilder()
			_ = _coll65
			for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(178), _dafny.IntOfInt64(967))); ; {
				_compr_65, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1525_v1 _dafny.Int
				_1525_v1 = interface{}(_compr_65).(_dafny.Int)
				if ((_dafny.IntOfInt64(178)).Cmp(_1525_v1) <= 0) && ((_1525_v1).Cmp(_dafny.IntOfInt64(967)) < 0) {
					_coll65.Add((_1525_v1).Times(_this.F40))
				}
			}
			return _coll65.ToSet()
		}())).Cardinality()))))
	}
}
func (_this *C6) Fm44(globalState *GlobalState) bool {
	{
		return (!(!(false))) && (!((!(false)) == (true)))
	}
}
func (_this *C6) M15(p0 _dafny.Int, p1 D14, p2 bool, p3 bool, globalState *GlobalState) (D1, _dafny.Map, _dafny.Int) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Map = _dafny.EmptyMap
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		(globalState).F12 = (_this).Fm1(globalState)
		var _1526_v0 _dafny.MultiSet
		_ = _1526_v0
		_1526_v0 = _dafny.MultiSetOf(p2)
		var _1527_v1 _dafny.Map
		_ = _1527_v1
		_1527_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1526_v0).Cardinality()).Plus(_dafny.IntOfInt64(866)), p2)
		var _1528_v2 D16
		_ = _1528_v2
		_1528_v2 = Companion_D16_.Create_DC37_(_1527_v1)
		_1527_v1 = (_1527_v1).Merge((_1528_v2).Dtor_cf68())
		(globalState).F15 = (p3) && (p2)
		var _1529_v3 _dafny.Sequence
		_ = _1529_v3
		_1529_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ymvn")
		var _1530_v4 _dafny.Map
		_ = _1530_v4
		_1530_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v3, p2)
		(globalState).F12 = (_dafny.Zero).Minus((func() _dafny.Int {
			if (_1526_v0).Contains(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v3, p2)).Equals(_1530_v4)) {
				return (_1526_v0).Multiplicity(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1529_v3, p2)).Equals(_1530_v4))
			}
			return p0
		})())
		var _1531_i0 _dafny.Int
		_ = _1531_i0
		_1531_i0 = _dafny.Zero
		{
			for p2 {
				{
					if (_1531_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1531_i0 = (_1531_i0).Plus(_dafny.One)
					var _1532_v5 D15
					_ = _1532_v5
					_1532_v5 = Companion_D15_.Create_DC35_(p2)
					_1532_v5 = _1532_v5
					var _1533_v6 _dafny.Array
					_ = _1533_v6
					var _len0_43 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_43
					var _nw233 _dafny.Array
					_ = _nw233
					if _len0_43.Cmp(_dafny.Zero) == 0 {
						_nw233 = _dafny.NewArray(_len0_43)
					} else {
						var _init43 func(_dafny.Int) _dafny.Int = func(_1534_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1534_i1, _dafny.IntOfInt64(669))
						}
						_ = _init43
						var _element0_43 = _init43(_dafny.Zero)
						_ = _element0_43
						_nw233 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
						(_nw233).ArraySet1(_element0_43, 0)
						var _nativeLen0_43 = (_len0_43).Int()
						_ = _nativeLen0_43
						for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
							(_nw233).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
						}
					}
					_1533_v6 = _nw233
					var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1533_v6), 0))
					_ = _index215
					(_1533_v6).ArraySet1(_this.F40, (_index215).Int())
					var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1533_v6), 0))
					_ = _index216
					(_1533_v6).ArraySet1(_dafny.IntOfInt64(134), (_index216).Int())
					var _1535_v7 T2
					_ = _1535_v7
					var _nw234 *C1 = New_C1_()
					_ = _nw234
					_nw234.Ctor__(p0, _this.F40, p2, p2)
					_1535_v7 = _nw234
					var _1536_v8 D7
					_ = _1536_v8
					_1536_v8 = Companion_D7_.Create_DC18_(_1535_v7)
					var _pat_let_tv31 = _1535_v7
					_ = _pat_let_tv31
					var _pat_let_tv32 = _1535_v7
					_ = _pat_let_tv32
					var _1537_v9 _dafny.Array
					_ = _1537_v9
					var _nwElement0_53 D7 = _1536_v8
					_ = _nwElement0_53
					var _nw235 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(18))
					_ = _nw235
					(_nw235).ArraySet1(_nwElement0_53, 0)
					(_nw235).ArraySet1(func(_pat_let25_0 D7) D7 {
						return func(_1538_dt__update__tmp_h0 D7) D7 {
							return func(_pat_let26_0 T2) D7 {
								return func(_1539_dt__update_hcf35_h0 T2) D7 {
									return Companion_D7_.Create_DC18_(_1539_dt__update_hcf35_h0)
								}(_pat_let26_0)
							}(_pat_let_tv31)
						}(_pat_let25_0)
					}(_1536_v8), 1)
					(_nw235).ArraySet1(_1536_v8, 2)
					(_nw235).ArraySet1(_1536_v8, 3)
					(_nw235).ArraySet1(_1536_v8, 4)
					(_nw235).ArraySet1(_1536_v8, 5)
					(_nw235).ArraySet1(_1536_v8, 6)
					(_nw235).ArraySet1(func(_pat_let27_0 D7) D7 {
						return func(_1540_dt__update__tmp_h1 D7) D7 {
							return func(_pat_let28_0 T2) D7 {
								return func(_1541_dt__update_hcf35_h1 T2) D7 {
									return Companion_D7_.Create_DC18_(_1541_dt__update_hcf35_h1)
								}(_pat_let28_0)
							}(_pat_let_tv32)
						}(_pat_let27_0)
					}(Companion_D7_.Create_DC18_(_1535_v7)), 7)
					(_nw235).ArraySet1(_1536_v8, 8)
					(_nw235).ArraySet1(_1536_v8, 9)
					(_nw235).ArraySet1(_1536_v8, 10)
					(_nw235).ArraySet1(_1536_v8, 11)
					(_nw235).ArraySet1(_1536_v8, 12)
					(_nw235).ArraySet1(_1536_v8, 13)
					(_nw235).ArraySet1(_1536_v8, 14)
					(_nw235).ArraySet1(Companion_D7_.Create_DC18_(_1535_v7), 15)
					(_nw235).ArraySet1(_1536_v8, 16)
					(_nw235).ArraySet1(_1536_v8, 17)
					_1537_v9 = _nw235
					var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1537_v9), 0))
					_ = _index217
					(_1537_v9).ArraySet1(_1536_v8, (_index217).Int())
					var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1537_v9), 0))
					_ = _index218
					(_1537_v9).ArraySet1(_1536_v8, (_index218).Int())
					var _1542_v10 _dafny.Array
					_ = _1542_v10
					var _nw236 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
					_ = _nw236
					_1542_v10 = _nw236
					var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1542_v10), 0))
					_ = _index219
					(_1542_v10).ArraySet1((_1535_v7).F26(), (_index219).Int())
					var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(263), _dafny.ArrayLen((_1542_v10), 0))
					_ = _index220
					(_1542_v10).ArraySet1(p2, (_index220).Int())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1543_i2 _dafny.Int
		_ = _1543_i2
		_1543_i2 = _dafny.Zero
		{
			for p2 {
				{
					if (_1543_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1543_i2 = (_1543_i2).Plus(_dafny.One)
					var _1544_v11 _dafny.Sequence
					_ = _1544_v11
					_1544_v11 = _dafny.SeqOf(p3)
					var _1545_v12 _dafny.Set
					_ = _1545_v12
					_1545_v12 = _dafny.SetOf(_this.F40, _this.F40, p0, p0, p0)
					var _1546_v13 _dafny.Array
					_ = _1546_v13
					var _nwElement0_54 bool = p3
					_ = _nwElement0_54
					var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(8))
					_ = _nw237
					(_nw237).ArraySet1(_nwElement0_54, 0)
					(_nw237).ArraySet1((_1544_v11).Select((Companion_Default___.SafeIndex((_1545_v12).Cardinality(), _dafny.IntOfUint32((_1544_v11).Cardinality()))).Uint32()).(bool), 1)
					(_nw237).ArraySet1(true, 2)
					(_nw237).ArraySet1((_this).Fm2(p2, p0, globalState), 3)
					(_nw237).ArraySet1(true, 4)
					(_nw237).ArraySet1(p3, 5)
					(_nw237).ArraySet1(p2, 6)
					(_nw237).ArraySet1(p2, 7)
					_1546_v13 = _nw237
					var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_1546_v13), 0))
					_ = _index221
					(_1546_v13).ArraySet1((_this).Fm44(globalState), (_index221).Int())
					var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_1546_v13), 0))
					_ = _index222
					(_1546_v13).ArraySet1((_this.F40).Cmp(((_this).Fm1(globalState)).Plus((_dafny.Zero).Minus(p0))) > 0, (_index222).Int())
					r2 = p0
					var _rhs224 _dafny.Int = (_this).Fm1(globalState)
					_ = _rhs224
					var _rhs225 _dafny.Int = (_this).Fm1(globalState)
					_ = _rhs225
					var _rhs226 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-197), (_dafny.Zero).Minus((_this).Fm1(globalState)), (_1545_v12).Cardinality(), ((_dafny.Zero).Minus(_this.F40)).Plus(_this.F40), _dafny.IntOfInt64(452))).Cardinality())
					_ = _rhs226
					var _rhs227 _dafny.Sequence = _1529_v3
					_ = _rhs227
					var _lhs208 *GlobalState = globalState
					_ = _lhs208
					var _lhs209 *GlobalState = globalState
					_ = _lhs209
					var _lhs210 *GlobalState = globalState
					_ = _lhs210
					var _lhs211 *GlobalState = globalState
					_ = _lhs211
					_lhs208.F16 = _rhs224
					_lhs209.F12 = _rhs225
					_lhs210.F12 = _rhs226
					_lhs211.F6 = _rhs227
					var _1547_v14 _dafny.Sequence
					_ = _1547_v14
					_1547_v14 = _dafny.SeqOf(_dafny.IntOfUint32((_1529_v3).Cardinality()))
					var _1548_v15 _dafny.MultiSet
					_ = _1548_v15
					_1548_v15 = _dafny.MultiSetOf(_dafny.SeqOf(p0, p0), _1547_v14, _1547_v14)
					var _1549_v16 *C3
					_ = _1549_v16
					var _nw238 *C3 = New_C3_()
					_ = _nw238
					_nw238.Ctor__(_1548_v15, (func() bool {
						if (_1546_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_1546_v13), 0))).Int()).(bool) {
							return p3
						}
						return true
					})(), (_1546_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_1546_v13), 0))).Int()).(bool))
					_1549_v16 = _nw238
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		r0 = Companion_D1_.Create_DC5_(Companion_Default___.SafeModuloInt(p0, p0))
		var _1550_v17 _dafny.Sequence
		_ = _1550_v17
		_1550_v17 = _dafny.SeqOf(p3)
		var _1551_v18 _dafny.Map
		_ = _1551_v18
		_1551_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1526_v0, (_1550_v17).Select((Companion_Default___.SafeIndex(_this.F40, _dafny.IntOfUint32((_1550_v17).Cardinality()))).Uint32()).(bool))
		r1 = _1551_v18
		r2 = (_this.F40).Times(_this.F40)
		return r0, r1, r2
	}
}
func (_this *C6) M16(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var _1552_v0 _dafny.Sequence
		_ = _1552_v0
		_1552_v0 = _dafny.SeqOf(_this.F40)
		var _1553_v1 _dafny.Sequence
		_ = _1553_v1
		_1553_v1 = _dafny.SeqOf(_1552_v0)
		var _1554_v2 bool
		_ = _1554_v2
		_1554_v2 = false
		var _rhs228 _dafny.Int = p0
		_ = _rhs228
		var _rhs229 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm45(_1554_v2, p0, _dafny.IntOfInt64(-164), _1554_v2, globalState), _1552_v0, _1552_v0), _1553_v1)
		_ = _rhs229
		var _lhs212 *GlobalState = globalState
		_ = _lhs212
		_lhs212.F16 = _rhs228
		_1553_v1 = _rhs229
		var _rhs230 _dafny.Int = p0
		_ = _rhs230
		var _rhs231 _dafny.Int = _this.F40
		_ = _rhs231
		var _lhs213 *GlobalState = globalState
		_ = _lhs213
		var _lhs214 *GlobalState = globalState
		_ = _lhs214
		_lhs213.F16 = _rhs230
		_lhs214.F12 = _rhs231
		var _1555_v3 _dafny.Array
		_ = _1555_v3
		var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw239
		_1555_v3 = _nw239
		var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))
		_ = _index223
		(_1555_v3).ArraySet1(p0, (_index223).Int())
		var _1556_v4 _dafny.Map
		_ = _1556_v4
		_1556_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F40, Companion_Default___.Fm47(_1554_v2, _1554_v2, _1554_v2, !(_1554_v2), globalState))
		var _1557_v5 _dafny.Sequence
		_ = _1557_v5
		_1557_v5 = _dafny.UnicodeSeqOfUtf8Bytes("mswelsh")
		var _1558_v6 _dafny.Map
		_ = _1558_v6
		_1558_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1557_v5).Cardinality()), _1554_v2)
		var _1559_v7 _dafny.MultiSet
		_ = _1559_v7
		_1559_v7 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0))
		var _1560_v8 _dafny.Map
		_ = _1560_v8
		_1560_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1559_v7, _this.F40)
		var _1561_v9 _dafny.Map
		_ = _1561_v9
		_1561_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, _this.F40), (_1560_v8).Cardinality())
		var _1562_v10 _dafny.Array
		_ = _1562_v10
		var _nwElement0_55 _dafny.Map = Companion_Default___.Fm46(_1554_v2, p0, p0, _1554_v2, globalState)
		_ = _nwElement0_55
		var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(4))
		_ = _nw240
		(_nw240).ArraySet1(_nwElement0_55, 0)
		(_nw240).ArraySet1(_1556_v4, 1)
		(_nw240).ArraySet1((_1556_v4).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1558_v6)), 2)
		(_nw240).ArraySet1(Companion_Default___.Fm46(_1554_v2, (func() _dafny.Int {
			if (_1561_v9).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxbdjm")).Cardinality()))) {
				return (_1561_v9).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxbdjm")).Cardinality()))).(_dafny.Int)
			}
			return _this.F40
		})(), p0, _1554_v2, globalState), 3)
		_1562_v10 = _nw240
		var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1562_v10), 0))
		_ = _index224
		(_1562_v10).ArraySet1(_1556_v4, (_index224).Int())
		var _1563_v11 _dafny.Array
		_ = _1563_v11
		var _nw241 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
		_ = _nw241
		_1563_v11 = _nw241
		var _1564_v12 _dafny.Sequence
		_ = _1564_v12
		_1564_v12 = _dafny.SeqOf(_1563_v11, _1563_v11)
		var _1565_v13 _dafny.Sequence
		_ = _1565_v13
		_1565_v13 = _dafny.SeqOf(_1564_v12)
		var _1566_v14 _dafny.Array
		_ = _1566_v14
		var _nwElement0_56 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1564_v12, _1564_v12)
		_ = _nwElement0_56
		var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(25))
		_ = _nw242
		(_nw242).ArraySet1(_nwElement0_56, 0)
		(_nw242).ArraySet1(_1564_v12, 1)
		(_nw242).ArraySet1(_1564_v12, 2)
		(_nw242).ArraySet1(_1564_v12, 3)
		(_nw242).ArraySet1(_1564_v12, 4)
		(_nw242).ArraySet1(_1564_v12, 5)
		(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1564_v12, (_1565_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(490), _dafny.IntOfUint32((_1565_v13).Cardinality()))).Uint32()).(_dafny.Sequence)), 6)
		(_nw242).ArraySet1(_1564_v12, 7)
		(_nw242).ArraySet1(_dafny.SeqOf(_1563_v11, _1563_v11), 8)
		(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1563_v11), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajrpp")).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_1563_v11)).Cardinality()))).Uint32(), _1563_v11), _1564_v12), 9)
		(_nw242).ArraySet1(_1564_v12, 10)
		(_nw242).ArraySet1(_1564_v12, 11)
		(_nw242).ArraySet1(_1564_v12, 12)
		(_nw242).ArraySet1(_1564_v12, 13)
		(_nw242).ArraySet1(_1564_v12, 14)
		(_nw242).ArraySet1(_dafny.SeqOf(_1563_v11, _1563_v11, _1563_v11), 15)
		(_nw242).ArraySet1(_1564_v12, 16)
		(_nw242).ArraySet1(_dafny.Companion_Sequence_.Update(_1564_v12, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1564_v12).Cardinality()))).Uint32(), _1563_v11), 17)
		(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1564_v12, _dafny.Companion_Sequence_.Update(_1564_v12, (Companion_Default___.SafeIndex(_this.F40, _dafny.IntOfUint32((_1564_v12).Cardinality()))).Uint32(), _1563_v11)), 18)
		(_nw242).ArraySet1(_1564_v12, 19)
		(_nw242).ArraySet1(_1564_v12, 20)
		(_nw242).ArraySet1(_1564_v12, 21)
		(_nw242).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1564_v12, _1564_v12), 22)
		(_nw242).ArraySet1(_1564_v12, 23)
		(_nw242).ArraySet1(_1564_v12, 24)
		_1566_v14 = _nw242
		var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1566_v14), 0))
		_ = _index225
		(_1566_v14).ArraySet1(_1564_v12, (_index225).Int())
		var _1567_v15 _dafny.CodePoint
		_ = _1567_v15
		_1567_v15 = _dafny.CodePoint('k')
		var _1568_v16 _dafny.Map
		_ = _1568_v16
		_1568_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(p0), _1567_v15)
		var _1569_v17 _dafny.Map
		_ = _1569_v17
		_1569_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
			if (_1568_v16).Contains(_1559_v7) {
				return (_1568_v16).Get(_1559_v7).(_dafny.CodePoint)
			}
			return (_1557_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1557_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
		})(), _1554_v2)
		var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))
		_ = _index226
		var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1562_v10), 0))
		_ = _index227
		var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1566_v14), 0))
		_ = _index228
		var _rhs232 _dafny.Int = p0
		_ = _rhs232
		var _rhs233 bool = (func() bool {
			if _1554_v2 {
				return _1554_v2
			}
			return (func() bool {
				if !(false) {
					return _1554_v2
				}
				return _1554_v2
			})()
		})()
		_ = _rhs233
		var _rhs234 _dafny.Map = _1556_v4
		_ = _rhs234
		var _rhs235 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1564_v12, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1564_v12).Cardinality()))).Uint32(), _1563_v11)
		_ = _rhs235
		var _rhs236 _dafny.Map = _1569_v17
		_ = _rhs236
		var _lhs215 _dafny.Array = _1555_v3
		_ = _lhs215
		var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))
		_ = _lhs216
		var _lhs217 *GlobalState = globalState
		_ = _lhs217
		var _lhs218 _dafny.Array = _1562_v10
		_ = _lhs218
		var _lhs219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_1562_v10), 0))
		_ = _lhs219
		var _lhs220 _dafny.Array = _1566_v14
		_ = _lhs220
		var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1566_v14), 0))
		_ = _lhs221
		(_lhs215).ArraySet1(_rhs232, (_lhs216).Int())
		_lhs217.F24 = _rhs233
		(_lhs218).ArraySet1(_rhs234, (_lhs219).Int())
		(_lhs220).ArraySet1(_rhs235, (_lhs221).Int())
		_1569_v17 = _rhs236
		(globalState).F24 = !(!(_1554_v2))
		var _1570_v18 D6
		_ = _1570_v18
		_1570_v18 = Companion_D6_.Create_DC16_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(977))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}((func(_1571_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1572_i0 _dafny.Int) _dafny.CodePoint {
				return _1571_v15
			}
		})(_1567_v15))))
		var _pat_let_tv33 = _1554_v2
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1554_v2
		_ = _pat_let_tv34
		if func(_source23 D6) bool {
			if _source23.Is_DC17() {
				var _1573___mcc_h0 _dafny.Int = _source23.Get_().(D6_DC17).Cf30
				_ = _1573___mcc_h0
				var _1574___mcc_h1 _dafny.Int = _source23.Get_().(D6_DC17).Cf31
				_ = _1574___mcc_h1
				var _1575___mcc_h2 _dafny.Int = _source23.Get_().(D6_DC17).Cf32
				_ = _1575___mcc_h2
				var _1576___mcc_h3 _dafny.Int = _source23.Get_().(D6_DC17).Cf33
				_ = _1576___mcc_h3
				var _1577___mcc_h4 _dafny.Int = _source23.Get_().(D6_DC17).Cf34
				_ = _1577___mcc_h4
				var _1578_cf34 _dafny.Int = _1577___mcc_h4
				_ = _1578_cf34
				var _1579_cf33 _dafny.Int = _1576___mcc_h3
				_ = _1579_cf33
				var _1580_cf32 _dafny.Int = _1575___mcc_h2
				_ = _1580_cf32
				var _1581_cf31 _dafny.Int = _1574___mcc_h1
				_ = _1581_cf31
				var _1582_cf30 _dafny.Int = _1573___mcc_h0
				_ = _1582_cf30
				return _pat_let_tv33
			} else {
				var _1583___mcc_h5 _dafny.Sequence = _source23.Get_().(D6_DC16).Cf29
				_ = _1583___mcc_h5
				var _1584_cf29 _dafny.Sequence = _1583___mcc_h5
				_ = _1584_cf29
				return _pat_let_tv34
			}
		}((func() D6 {
			if _1554_v2 {
				return _1570_v18
			}
			return _1570_v18
		})()) {
			var _1585_v19 _dafny.Sequence
			_ = _1585_v19
			_1585_v19 = _dafny.SeqOf(_1554_v2)
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))
			_ = _index229
			(_1555_v3).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1585_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1585_v19, (Companion_Default___.SafeIndex(_this.F40, _dafny.IntOfUint32((_1585_v19).Cardinality()))).Uint32(), _1554_v2), _1585_v19))).Cardinality()), (_index229).Int())
			(globalState).F24 = _1554_v2
			_1555_v3 = _1555_v3
			_1555_v3 = _1555_v3
			var _1586_v20 D1
			_ = _1586_v20
			_1586_v20 = Companion_D1_.Create_DC5_((_dafny.Zero).Minus(_this.F40))
			var _1587_v21 _dafny.Set
			_ = _1587_v21
			_1587_v21 = _dafny.SetOf(_1586_v20)
			var _1588_v22 _dafny.Map
			_ = _1588_v22
			_1588_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yeue"), _1557_v5)).Cardinality()))
			var _rhs237 _dafny.Set = (_dafny.SetOf(_1586_v20, _1586_v20, _1586_v20)).Difference(_1587_v21)
			_ = _rhs237
			var _rhs238 _dafny.Map = _1588_v22
			_ = _rhs238
			var _rhs239 _dafny.Array = _1563_v11
			_ = _rhs239
			_1587_v21 = _rhs237
			_1588_v22 = _rhs238
			_1563_v11 = _rhs239
		} else {
			var _1589_v23 _dafny.Sequence
			_ = _1589_v23
			_1589_v23 = _dafny.SeqOf(_1557_v5)
			var _1590_v24 _dafny.Array
			_ = _1590_v24
			var _nwElement0_57 _dafny.CodePoint = _1567_v15
			_ = _nwElement0_57
			var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(2))
			_ = _nw243
			(_nw243).ArraySet1CodePoint(_nwElement0_57, 0)
			(_nw243).ArraySet1CodePoint(_1567_v15, 1)
			_1590_v24 = _nw243
			var _1591_v25 D1
			_ = _1591_v25
			_1591_v25 = Companion_D1_.Create_DC4_(_1590_v24, p0, p0, _1557_v5, _dafny.UnicodeSeqOfUtf8Bytes("cjntvhbdw"))
			var _1592_v26 _dafny.Array
			_ = _1592_v26
			var _nwElement0_58 _dafny.Sequence = _1557_v5
			_ = _nwElement0_58
			var _nw244 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(20))
			_ = _nw244
			(_nw244).ArraySet1(_nwElement0_58, 0)
			(_nw244).ArraySet1(_1557_v5, 1)
			(_nw244).ArraySet1(_1557_v5, 2)
			(_nw244).ArraySet1(_1557_v5, 3)
			(_nw244).ArraySet1((_1589_v23).Select((Companion_Default___.SafeIndex(_this.F40, _dafny.IntOfUint32((_1589_v23).Cardinality()))).Uint32()).(_dafny.Sequence), 4)
			(_nw244).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("p"), 5)
			(_nw244).ArraySet1(_1557_v5, 6)
			(_nw244).ArraySet1(_1557_v5, 7)
			(_nw244).ArraySet1(_1557_v5, 8)
			(_nw244).ArraySet1(_1557_v5, 9)
			(_nw244).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("agk"), 10)
			(_nw244).ArraySet1((_1591_v25).Dtor_cf10(), 11)
			(_nw244).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1557_v5, _1557_v5), 12)
			(_nw244).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(566))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg95 _dafny.Int) interface{} {
					return coer95(arg95)
				}
			}((func(_1593_v15 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1594_i1 _dafny.Int) _dafny.CodePoint {
					return _1593_v15
				}
			})(_1567_v15))), 13)
			(_nw244).ArraySet1(_1557_v5, 14)
			(_nw244).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lwu"), 15)
			(_nw244).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1557_v5, _1557_v5), 16)
			(_nw244).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gdid"), 17)
			(_nw244).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1557_v5, _1557_v5), 18)
			(_nw244).ArraySet1(_1557_v5, 19)
			_1592_v26 = _nw244
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1592_v26), 0))
			_ = _index230
			(_1592_v26).ArraySet1(_1557_v5, (_index230).Int())
			var _1595_v27 _dafny.Set
			_ = _1595_v27
			_1595_v27 = _dafny.SetOf(_1554_v2)
			var _1596_v28 _dafny.MultiSet
			_ = _1596_v28
			_1596_v28 = _dafny.MultiSetOf(true)
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1592_v26), 0))
			_ = _index231
			(_1592_v26).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1557_v5, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if true {
					return _1557_v5
				}
				return _dafny.Companion_Sequence_.Update(_1557_v5, (Companion_Default___.SafeIndex((_1595_v27).Cardinality(), _dafny.IntOfUint32((_1557_v5).Cardinality()))).Uint32(), _1567_v15)
			})(), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1596_v28).Contains(false) {
					return (_1596_v28).Multiplicity(false)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1559_v7).Cardinality(), _1554_v2)).Cardinality()
			})(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if true {
					return _1557_v5
				}
				return _dafny.Companion_Sequence_.Update(_1557_v5, (Companion_Default___.SafeIndex((_1595_v27).Cardinality(), _dafny.IntOfUint32((_1557_v5).Cardinality()))).Uint32(), _1567_v15)
			})()).Cardinality()))).Uint32(), _1567_v15)), (_index231).Int())
			var _1597_v29 _dafny.Map
			_ = _1597_v29
			_1597_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, _1554_v2)
			(globalState).F22 = !((_1554_v2) || ((_this.F40).Cmp((_1597_v29).Cardinality()) <= 0))
			var _1598_v30 D5
			_ = _1598_v30
			_1598_v30 = Companion_D5_.Create_DC14_(p0, _this.F40, _1554_v2)
			var _1599_v31 _dafny.Array
			_ = _1599_v31
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw245
			_1599_v31 = _nw245
			var _1600_v32 _dafny.Map
			_ = _1600_v32
			_1600_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1599_v31, (func() _dafny.Array {
				if _1554_v2 {
					return _1555_v3
				}
				return _1555_v3
			})())
			var _1601_v33 _dafny.Map
			_ = _1601_v33
			_1601_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, (_dafny.Zero).Minus(_this.F40))
			var _1602_v34 D7
			_ = _1602_v34
			_1602_v34 = Companion_D7_.Create_DC19_((_dafny.Zero).Minus((_1555_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))).Int()).(_dafny.Int)), _1596_v28, _1554_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, (_1555_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))).Int()).(_dafny.Int)))
			var _1603_v35 D7
			_ = _1603_v35
			_1603_v35 = Companion_D7_.Create_DC19_(_this.F40, _1596_v28, _1554_v2, (_1602_v34).Dtor_cf39())
			var _pat_let_tv35 = globalState
			_ = _pat_let_tv35
			var _rhs240 _dafny.Array = (func() _dafny.Array {
				if (_1600_v32).Contains(_1599_v31) {
					return (_1600_v32).Get(_1599_v31).(_dafny.Array)
				}
				return _1555_v3
			})()
			_ = _rhs240
			var _rhs241 bool = ((func() _dafny.Int {
				if (_1601_v33).Contains(_1554_v2) {
					return (_1601_v33).Get(_1554_v2).(_dafny.Int)
				}
				return (_1601_v33).Cardinality()
			})()).Cmp((_1603_v35).Dtor_cf36()) <= 0
			_ = _rhs241
			var _rhs242 D5 = func(_pat_let29_0 D5) D5 {
				return func(_1604_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let30_0 _dafny.Int) D5 {
						return func(_1605_dt__update_hcf26_h0 _dafny.Int) D5 {
							return func(_pat_let31_0 _dafny.Int) D5 {
								return func(_1606_dt__update_hcf25_h0 _dafny.Int) D5 {
									return Companion_D5_.Create_DC14_(_1606_dt__update_hcf25_h0, _1605_dt__update_hcf26_h0, (_1604_dt__update__tmp_h0).Dtor_cf27())
								}(_pat_let31_0)
							}(_dafny.IntOfInt64(176))
						}(_pat_let30_0)
					}((_this).Fm1(_pat_let_tv35))
				}(_pat_let29_0)
			}(_1598_v30)
			_ = _rhs242
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			_1555_v3 = _rhs240
			_lhs222.F24 = _rhs241
			_1598_v30 = _rhs242
			_1564_v12 = _dafny.Companion_Sequence_.Concatenate(_1564_v12, (_1566_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1566_v14), 0))).Int()).(_dafny.Sequence))
			var _1607_v36 _dafny.Array
			_ = _1607_v36
			var _nwElement0_59 _dafny.Array = _1555_v3
			_ = _nwElement0_59
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(2))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_59, 0)
			(_nw246).ArraySet1(_1555_v3, 1)
			_1607_v36 = _nw246
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1607_v36), 0))
			_ = _index232
			(_1607_v36).ArraySet1(_1555_v3, (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1607_v36), 0))
			_ = _index233
			(_1607_v36).ArraySet1((func() _dafny.Array {
				if true {
					return _1555_v3
				}
				return _1555_v3
			})(), (_index233).Int())
		}
		var _1608_v37 _dafny.Set
		_ = _1608_v37
		_1608_v37 = _dafny.SetOf((_1555_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1557_v5).Cardinality()), p0, (_1555_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_1555_v3), 0))).Int()).(_dafny.Int))
		var _1609_v38 _dafny.Map
		_ = _1609_v38
		_1609_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1554_v2, _1608_v37)
		_1609_v38 = (_1609_v38).Update(_dafny.Companion_Sequence_.IsPrefixOf(_1557_v5, _dafny.UnicodeSeqOfUtf8Bytes("krmbocy")), _1608_v37)
		var _1610_v39 _dafny.Sequence
		_ = _1610_v39
		_1610_v39 = _dafny.SeqOf(_1554_v2, _1554_v2, _1554_v2, _1554_v2, _1554_v2)
		r0 = _dafny.Companion_Sequence_.Concatenate(_1610_v39, _dafny.SeqOf(_1554_v2, _1554_v2, false, _1554_v2))
		return r0
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	F38 bool
	F39 _dafny.Map
}

func New_C7_() *C7 {
	_this := C7{}

	_this.F38 = false
	_this.F39 = _dafny.EmptyMap
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

func (_this *C7) Ctor__(f38 bool, f39 _dafny.Map) {
	{
		(_this).F38 = f38
		(_this).F39 = f39
	}
}
func (_this *C7) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return (func(_source24 _dafny.Sequence) _dafny.Map {
			var _1611___mcc_h0 _dafny.Sequence = _source24
			_ = _1611___mcc_h0
			var _1612_cf40 _dafny.Sequence = _1611___mcc_h0
			_ = _1612_cf40
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_this.F38, false, _this.F38, _this.F38, true)).Cardinality()), _dafny.IntOfInt64(428))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(818), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (Companion_D2_.Create_DC9_(_dafny.IntOfInt64(795), _dafny.IntOfInt64(-873), _this.F38)).Dtor_cf20())).Cardinality()))
		}(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-746))).Cardinality(), _this.F38)).Cardinality()))).Cardinality()
	}
}
func (_this *C7) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (func() _dafny.Set {
			var _coll66 = _dafny.NewBuilder()
			_ = _coll66
			for _iter66 := _dafny.Iterate((_dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-201)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("imnioove")).Cardinality()), _dafny.IntOfInt64(112)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(373))), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetOf(_this.F38, _this.F38)).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(987)))).Elements()); ; {
				_compr_66, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1613_v0 _dafny.MultiSet
				_1613_v0 = interface{}(_compr_66).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-201)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("imnioove")).Cardinality()), _dafny.IntOfInt64(112)), _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(373))), _dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetOf(_this.F38, _this.F38)).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfInt64(987))), _1613_v0) {
					_coll66.Add(_1613_v0)
				}
			}
			return _coll66.ToSet()
		}()).IsProperSubsetOf((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(699), _dafny.IntOfInt64(-34), _dafny.IntOfInt64(-43), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqgjwd")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F38), true)).Cardinality())), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jhuvf")).Cardinality())))).Difference(_dafny.SetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihdf")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-567))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg96 _dafny.Int) interface{} {
				return coer96(arg96)
			}
		}(func(_1614_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("ro")
		}))).Cardinality()), _dafny.IntOfInt64(871))).Cardinality())))))
	}
}
func (_this *C7) M13(globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1615_v0 _dafny.Int
		_ = _1615_v0
		_1615_v0 = _dafny.IntOfInt64(616)
		var _1616_v1 D12
		_ = _1616_v1
		_1616_v1 = Companion_D12_.Create_DC28_(_1615_v0)
		var _1617_v2 _dafny.Map
		_ = _1617_v2
		_1617_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(!(_this.F38), _1615_v0, globalState), _1615_v0)
		var _pat_let_tv38 = _1615_v0
		_ = _pat_let_tv38
		var _pat_let_tv39 = _1615_v0
		_ = _pat_let_tv39
		var _hi13 _dafny.Int = (func() _dafny.Int {
			if (_1617_v2).Contains(_this.F38) {
				return (_1617_v2).Get(_this.F38).(_dafny.Int)
			}
			return _1615_v0
		})()
		_ = _hi13
		for _1618_i0 := _dafny.IntOfUint32((_dafny.SeqOf(_1616_v1, _1616_v1, _1616_v1, func(_pat_let36_0 D12) D12 {
			return func(_1648_dt__update__tmp_h1 D12) D12 {
				return func(_pat_let39_0 _dafny.Int) D12 {
					return func(_1649_dt__update_hcf49_h1 _dafny.Int) D12 {
						return Companion_D12_.Create_DC28_(_1649_dt__update_hcf49_h1)
					}(_pat_let39_0)
				}(_pat_let_tv39)
			}(_pat_let36_0)
		}(func(_pat_let37_0 D12) D12 {
			return func(_1646_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let38_0 _dafny.Int) D12 {
					return func(_1647_dt__update_hcf49_h0 _dafny.Int) D12 {
						return Companion_D12_.Create_DC28_(_1647_dt__update_hcf49_h0)
					}(_pat_let38_0)
				}(_pat_let_tv38)
			}(_pat_let37_0)
		}(_1616_v1)), _1616_v1)).Cardinality()); _1618_i0.Cmp(_hi13) < 0; _1618_i0 = _1618_i0.Plus(_dafny.One) {
			var _1619_v3 _dafny.Map
			_ = _1619_v3
			_1619_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1615_v0, _1618_i0)
			var _1620_v4 D0
			_ = _1620_v4
			_1620_v4 = Companion_D0_.Create_DC1_(_this.F38, _this.F38, (_1619_v3).Cardinality(), _this.F38)
			_1620_v4 = _1620_v4
			var _1621_v5 D5
			_ = _1621_v5
			_1621_v5 = Companion_D5_.Create_DC15_(Companion_D5_.Create_DC14_(_1618_i0, (_this).Fm1(globalState), (_this).Fm2(_this.F38, _1615_v0, globalState)))
			var _1622_v6 D5
			_ = _1622_v6
			_1622_v6 = Companion_D5_.Create_DC15_(_1621_v5)
			var _1623_v7 _dafny.Array
			_ = _1623_v7
			var _nwElement0_60 D5 = _1622_v6
			_ = _nwElement0_60
			var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(8))
			_ = _nw247
			(_nw247).ArraySet1(_nwElement0_60, 0)
			(_nw247).ArraySet1(Companion_D5_.Create_DC15_(_1621_v5), 1)
			(_nw247).ArraySet1(_1622_v6, 2)
			(_nw247).ArraySet1(_1622_v6, 3)
			(_nw247).ArraySet1(_1622_v6, 4)
			(_nw247).ArraySet1(_1622_v6, 5)
			(_nw247).ArraySet1(_1622_v6, 6)
			(_nw247).ArraySet1(_1622_v6, 7)
			_1623_v7 = _nw247
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1623_v7), 0))
			_ = _index234
			(_1623_v7).ArraySet1((func() D5 {
				if !(!(_this.F38)) {
					return _1622_v6
				}
				return _1622_v6
			})(), (_index234).Int())
			var _pat_let_tv36 = _1621_v5
			_ = _pat_let_tv36
			var _pat_let_tv37 = _1621_v5
			_ = _pat_let_tv37
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(53), _dafny.ArrayLen((_1623_v7), 0))
			_ = _index235
			(_1623_v7).ArraySet1(func(_pat_let32_0 D5) D5 {
				return func(_1626_dt__update__tmp_h3 D5) D5 {
					return func(_pat_let35_0 D5) D5 {
						return func(_1627_dt__update_hcf28_h1 D5) D5 {
							return Companion_D5_.Create_DC15_(_1627_dt__update_hcf28_h1)
						}(_pat_let35_0)
					}(_pat_let_tv37)
				}(_pat_let32_0)
			}(func(_pat_let33_0 D5) D5 {
				return func(_1624_dt__update__tmp_h2 D5) D5 {
					return func(_pat_let34_0 D5) D5 {
						return func(_1625_dt__update_hcf28_h0 D5) D5 {
							return Companion_D5_.Create_DC15_(_1625_dt__update_hcf28_h0)
						}(_pat_let34_0)
					}(_pat_let_tv36)
				}(_pat_let33_0)
			}(Companion_D5_.Create_DC15_(Companion_D5_.Create_DC14_((_dafny.Zero).Minus(_1615_v0), _1618_i0, _this.F38)))), (_index235).Int())
			var _1628_v8 _dafny.Sequence
			_ = _1628_v8
			_1628_v8 = _dafny.SeqOf(_1615_v0)
			var _1629_v9 _dafny.Sequence
			_ = _1629_v9
			_1629_v9 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1628_v8, (Companion_Default___.SafeIndex(_1618_i0, _dafny.IntOfUint32((_1628_v8).Cardinality()))).Uint32(), _1615_v0))
			var _1630_v10 *C2
			_ = _1630_v10
			var _nw248 *C2 = New_C2_()
			_ = _nw248
			_nw248.Ctor__(_dafny.Companion_Sequence_.Update(_1629_v9, (Companion_Default___.SafeIndex(_1615_v0, _dafny.IntOfUint32((_1629_v9).Cardinality()))).Uint32(), _dafny.SeqOf(_1615_v0, _dafny.IntOfInt64(250), _1615_v0, _1618_i0)), !(_this.F38), _this.F38)
			_1630_v10 = _nw248
			var _1631_v11 _dafny.MultiSet
			_ = _1631_v11
			_1631_v11 = _dafny.MultiSetOf(_dafny.IntOfInt64(-248))
			var _1632_v12 _dafny.Sequence
			_ = _1632_v12
			_1632_v12 = _dafny.SeqOf(_this.F38, _this.F38)
			var _1633_v13 _dafny.Array
			_ = _1633_v13
			var _nwElement0_61 _dafny.Int = _1618_i0
			_ = _nwElement0_61
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(19))
			_ = _nw249
			(_nw249).ArraySet1(_nwElement0_61, 0)
			(_nw249).ArraySet1((_1631_v11).Cardinality(), 1)
			(_nw249).ArraySet1(_1618_i0, 2)
			(_nw249).ArraySet1(_dafny.IntOfInt64(812), 3)
			(_nw249).ArraySet1(_dafny.IntOfInt64(661), 4)
			(_nw249).ArraySet1(_1618_i0, 5)
			(_nw249).ArraySet1(_1618_i0, 6)
			(_nw249).ArraySet1(_1618_i0, 7)
			(_nw249).ArraySet1(_dafny.IntOfUint32((_1632_v12).Cardinality()), 8)
			(_nw249).ArraySet1(_1618_i0, 9)
			(_nw249).ArraySet1(_1618_i0, 10)
			(_nw249).ArraySet1(_1615_v0, 11)
			(_nw249).ArraySet1(_1615_v0, 12)
			(_nw249).ArraySet1((_1628_v8).Select((Companion_Default___.SafeIndex(_1618_i0, _dafny.IntOfUint32((_1628_v8).Cardinality()))).Uint32()).(_dafny.Int), 13)
			(_nw249).ArraySet1(_1615_v0, 14)
			(_nw249).ArraySet1(_1618_i0, 15)
			(_nw249).ArraySet1(_1618_i0, 16)
			(_nw249).ArraySet1(_1618_i0, 17)
			(_nw249).ArraySet1(_1618_i0, 18)
			_1633_v13 = _nw249
			var _1634_v14 _dafny.Map
			_ = _1634_v14
			_1634_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1633_v13, false)
			var _1635_v15 D5
			_ = _1635_v15
			_1635_v15 = Companion_D5_.Create_DC14_(_1618_i0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1632_v12).Cardinality())), _this.F38)
			var _1636_v16 D2
			_ = _1636_v16
			_1636_v16 = Companion_D2_.Create_DC7_(_this.F38, _this.F38, (_1634_v14).Cardinality(), (_1635_v15).Dtor_cf25())
			var _1637_v17 _dafny.Array
			_ = _1637_v17
			var _nw250 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw250
			_1637_v17 = _nw250
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1637_v17), 0))
			_ = _index236
			(_1637_v17).ArraySet1(_this.F38, (_index236).Int())
			var _1638_v18 _dafny.Sequence
			_ = _1638_v18
			_1638_v18 = _dafny.UnicodeSeqOfUtf8Bytes("mhjfmmfl")
			var _1639_v19 D3
			_ = _1639_v19
			_1639_v19 = Companion_D3_.Create_DC11_((_1638_v18).Select((Companion_Default___.SafeIndex(_1618_i0, _dafny.IntOfUint32((_1638_v18).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _1640_v20 _dafny.Map
			_ = _1640_v20
			_1640_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1639_v19)
			var _1641_v21 _dafny.Map
			_ = _1641_v21
			_1641_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1628_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer97 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
				return func(arg97 _dafny.Int) interface{} {
					return coer97(arg97)
				}
			}((func(_1642_v3 _dafny.Map) func(_dafny.Int) D2 {
				return func(_1643_i1 _dafny.Int) D2 {
					return Companion_D2_.Create_DC6_(_1642_v3)
				}
			})(_1619_v3))))
			var _1644_v22 D10
			_ = _1644_v22
			_1644_v22 = Companion_D10_.Create_DC23_(_1641_v21)
			var _1645_v23 _dafny.Map
			_ = _1645_v23
			_1645_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1644_v22, _this.F38)
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1637_v17), 0))
			_ = _index237
			var _rhs243 *C2 = (func() *C2 {
				if _this.F38 {
					return _1630_v10
				}
				return _1630_v10
			})()
			_ = _rhs243
			var _rhs244 D2 = _1636_v16
			_ = _rhs244
			var _rhs245 _dafny.Int = (_1640_v20).Cardinality()
			_ = _rhs245
			var _rhs246 _dafny.Int = _1618_i0
			_ = _rhs246
			var _rhs247 bool = (func() bool {
				if (_1645_v23).Contains(_1644_v22) {
					return (_1645_v23).Get(_1644_v22).(bool)
				}
				return _this.F38
			})()
			_ = _rhs247
			var _lhs223 *GlobalState = globalState
			_ = _lhs223
			var _lhs224 *GlobalState = globalState
			_ = _lhs224
			var _lhs225 _dafny.Array = _1637_v17
			_ = _lhs225
			var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1637_v17), 0))
			_ = _lhs226
			_1630_v10 = _rhs243
			_1636_v16 = _rhs244
			_lhs223.F16 = _rhs245
			_lhs224.F12 = _rhs246
			(_lhs225).ArraySet1(_rhs247, (_lhs226).Int())
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1637_v17), 0))
			_ = _index238
			(_1637_v17).ArraySet1((_1637_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(687), _dafny.ArrayLen((_1637_v17), 0))).Int()).(bool), (_index238).Int())
		}
		var _1650_v24 _dafny.CodePoint
		_ = _1650_v24
		_1650_v24 = _dafny.CodePoint('l')
		var _1651_v25 _dafny.CodePoint
		_ = _1651_v25
		var _1652_v26 _dafny.Sequence
		_ = _1652_v26
		var _1653_v27 bool
		_ = _1653_v27
		var _1654_v28 _dafny.Map
		_ = _1654_v28
		var _out49 _dafny.CodePoint
		_ = _out49
		var _out50 _dafny.Sequence
		_ = _out50
		var _out51 bool
		_ = _out51
		var _out52 _dafny.Map
		_ = _out52
		_out49, _out50, _out51, _out52 = Companion_Default___.M0(_1650_v24, globalState)
		_1651_v25 = _out49
		_1652_v26 = _out50
		_1653_v27 = _out51
		_1654_v28 = _out52
		var _1655_v29 _dafny.Map
		_ = _1655_v29
		_1655_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1651_v25, _dafny.IntOfInt64(369))
		var _rhs248 bool = !((Companion_Default___.SafeDivisionInt(_1615_v0, (_this).Fm1(globalState))).Cmp(_1615_v0) >= 0)
		_ = _rhs248
		var _rhs249 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(809), _dafny.IntOfInt64(846))).Times((func() _dafny.Int {
			if (_1655_v29).Contains(Companion_Default___.Fm40(_1615_v0, (_this).Fm2(_1653_v27, _1615_v0, globalState), _1653_v27, globalState)) {
				return (_1655_v29).Get(Companion_Default___.Fm40(_1615_v0, (_this).Fm2(_1653_v27, _1615_v0, globalState), _1653_v27, globalState)).(_dafny.Int)
			}
			return _1615_v0
		})())
		_ = _rhs249
		var _lhs227 *GlobalState = globalState
		_ = _lhs227
		_lhs227.F24 = _rhs248
		r1 = _rhs249
		var _1656_v30 _dafny.MultiSet
		_ = _1656_v30
		_1656_v30 = _dafny.MultiSetOf(_1615_v0, _1615_v0)
		var _1657_v31 _dafny.Map
		_ = _1657_v31
		_1657_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v30, _this.F38)
		var _1658_v32 _dafny.Map
		_ = _1658_v32
		_1658_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1615_v0, (_1657_v31).Cardinality())
		_1658_v32 = (_1658_v32).Update((_dafny.Zero).Minus(_1615_v0), _1615_v0)
		(_this).F38 = _1653_v27
		var _1659_v33 D10
		_ = _1659_v33
		_1659_v33 = Companion_D10_.Create_DC25_()
		var _pat_let_tv40 = _1656_v30
		_ = _pat_let_tv40
		var _pat_let_tv41 = _1656_v30
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1615_v0
		_ = _pat_let_tv42
		var _pat_let_tv43 = _1656_v30
		_ = _pat_let_tv43
		_1656_v30 = func(_source25 D10) _dafny.MultiSet {
			if _source25.Is_DC24() {
				var _1660___mcc_h0 _dafny.Sequence = _source25.Get_().(D10_DC24).Cf44
				_ = _1660___mcc_h0
				var _1661___mcc_h1 _dafny.Sequence = _source25.Get_().(D10_DC24).Cf45
				_ = _1661___mcc_h1
				var _1662___mcc_h2 bool = _source25.Get_().(D10_DC24).Cf46
				_ = _1662___mcc_h2
				var _1663_cf46 bool = _1662___mcc_h2
				_ = _1663_cf46
				var _1664_cf45 _dafny.Sequence = _1661___mcc_h1
				_ = _1664_cf45
				var _1665_cf44 _dafny.Sequence = _1660___mcc_h0
				_ = _1665_cf44
				return _pat_let_tv40
			} else if _source25.Is_DC25() {
				return (_pat_let_tv41).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iowvgu")).Cardinality()), Companion_Default___.Abs(_pat_let_tv42))
			} else {
				var _1666___mcc_h3 _dafny.Map = _source25.Get_().(D10_DC23).Cf43
				_ = _1666___mcc_h3
				var _1667_cf43 _dafny.Map = _1666___mcc_h3
				_ = _1667_cf43
				return _pat_let_tv43
			}
		}(_1659_v33)
		r0 = _1653_v27
		r1 = (_1615_v0).Times(_1615_v0)
		r2 = _1653_v27
		return r0, r1, r2
	}
}
func (_this *C7) M14(p0 bool, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) (bool, _dafny.MultiSet, _dafny.Map, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _1668_v0 _dafny.Int
		_ = _1668_v0
		_1668_v0 = _dafny.IntOfInt64(-497)
		var _1669_v1 _dafny.Sequence
		_ = _1669_v1
		_1669_v1 = _dafny.SeqOf(_1668_v0)
		var _1670_v2 *C2
		_ = _1670_v2
		var _nw251 *C2 = New_C2_()
		_ = _nw251
		_nw251.Ctor__(_dafny.SeqOf(_1669_v1), !(p0), _this.F38)
		_1670_v2 = _nw251
		var _1671_v3 _dafny.Sequence
		_ = _1671_v3
		_1671_v3 = _dafny.SeqOf((_1670_v2).Fm2(p1, _1668_v0, globalState), p0)
		var _1672_v4 D1
		_ = _1672_v4
		_1672_v4 = Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1), _1671_v3))
		var _source26 D1 = _1672_v4
		_ = _source26
		if _source26.Is_DC4() {
			var _1673___mcc_h0 _dafny.Array = _source26.Get_().(D1_DC4).Cf7
			_ = _1673___mcc_h0
			var _1674___mcc_h1 _dafny.Int = _source26.Get_().(D1_DC4).Cf8
			_ = _1674___mcc_h1
			var _1675___mcc_h2 _dafny.Int = _source26.Get_().(D1_DC4).Cf9
			_ = _1675___mcc_h2
			var _1676___mcc_h3 _dafny.Sequence = _source26.Get_().(D1_DC4).Cf10
			_ = _1676___mcc_h3
			var _1677___mcc_h4 _dafny.Sequence = _source26.Get_().(D1_DC4).Cf11
			_ = _1677___mcc_h4
			var _1678_cf11 _dafny.Sequence = _1677___mcc_h4
			_ = _1678_cf11
			var _1679_cf10 _dafny.Sequence = _1676___mcc_h3
			_ = _1679_cf10
			var _1680_cf9 _dafny.Int = _1675___mcc_h2
			_ = _1680_cf9
			var _1681_cf8 _dafny.Int = _1674___mcc_h1
			_ = _1681_cf8
			var _1682_cf7 _dafny.Array = _1673___mcc_h0
			_ = _1682_cf7
			var _1683_v5 _dafny.Map
			_ = _1683_v5
			_1683_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).Fm2(p1, _1681_cf8, globalState))
			var _1684_v6 _dafny.Set
			_ = _1684_v6
			_1684_v6 = _dafny.SetOf(_this.F38, (func() bool {
				if (_1683_v5).Contains(_this.F38) {
					return (_1683_v5).Get(_this.F38).(bool)
				}
				return p1
			})(), !(_this.F38), _this.F38, p0)
			var _1685_v7 _dafny.Map
			_ = _1685_v7
			_1685_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				if p1 {
					return _1684_v6
				}
				return _1684_v6
			})(), p1)
			var _1686_v8 _dafny.Set
			_ = _1686_v8
			_1686_v8 = _1684_v6
			var _rhs250 bool = ((_1684_v6).Difference(_1684_v6)).IsSubsetOf((_1684_v6).Intersection(_dafny.SetOf(true)))
			_ = _rhs250
			var _rhs251 _dafny.Sequence = _1671_v3
			_ = _rhs251
			var _rhs252 _dafny.Sequence = _1671_v3
			_ = _rhs252
			var _rhs253 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1684_v6).Cardinality(), _1681_cf8))
			_ = _rhs253
			var _rhs254 bool = (func() bool {
				if (_1685_v7).Contains(_1686_v8) {
					return (_1685_v7).Get(_1686_v8).(bool)
				}
				return _this.F38
			})()
			_ = _rhs254
			var _lhs228 *GlobalState = globalState
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			_lhs228.F22 = _rhs250
			_1671_v3 = _rhs251
			_1671_v3 = _rhs252
			_1680_cf9 = _rhs253
			_lhs229.F22 = _rhs254
			var _1687_v9 _dafny.Sequence
			_ = _1687_v9
			_1687_v9 = _dafny.SeqOf(_1684_v6)
			var _1688_v10 D14
			_ = _1688_v10
			_1688_v10 = Companion_D14_.Create_DC31_(_dafny.SeqOf(_1684_v6))
			var _1689_v11 _dafny.Map
			_ = _1689_v11
			_1689_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SeqOf(_1684_v6))
			var _1690_v12 _dafny.Array
			_ = _1690_v12
			var _nwElement0_62 _dafny.Sequence = _1687_v9
			_ = _nwElement0_62
			var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(18))
			_ = _nw252
			(_nw252).ArraySet1(_nwElement0_62, 0)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1687_v9, _dafny.Companion_Sequence_.Update(_1687_v9, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.IntOfUint32((_1687_v9).Cardinality()))).Uint32(), _dafny.SetOf(p0, true))), 1)
			(_nw252).ArraySet1(_1687_v9, 2)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1687_v9, (Companion_Default___.SafeIndex(_1681_cf8, _dafny.IntOfUint32((_1687_v9).Cardinality()))).Uint32(), _1684_v6), _1687_v9), 3)
			(_nw252).ArraySet1(_1687_v9, 4)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1687_v9, _1687_v9), 5)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1684_v6, _1684_v6, _1684_v6), _1687_v9), 6)
			(_nw252).ArraySet1(_dafny.SeqOf(_1684_v6), 7)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1687_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(799))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg98 _dafny.Int) interface{} {
					return coer98(arg98)
				}
			}((func(_1691_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_1692_i0 _dafny.Int) _dafny.Set {
					return _1691_v6
				}
			})(_1684_v6)))), 8)
			(_nw252).ArraySet1(_1687_v9, 9)
			(_nw252).ArraySet1(_1687_v9, 10)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Update(_1687_v9, (Companion_Default___.SafeIndex(_1668_v0, _dafny.IntOfUint32((_1687_v9).Cardinality()))).Uint32(), _1684_v6), 11)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1688_v10).Dtor_cf55(), _1687_v9), 12)
			(_nw252).ArraySet1((func() _dafny.Sequence {
				if (_1689_v11).Contains(p0) {
					return (_1689_v11).Get(p0).(_dafny.Sequence)
				}
				return _1687_v9
			})(), 13)
			(_nw252).ArraySet1((_1688_v10).Dtor_cf55(), 14)
			(_nw252).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(785))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1693_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
				return func(_1694_i1 _dafny.Int) _dafny.Set {
					return _1693_v6
				}
			})(_1684_v6))), 15)
			(_nw252).ArraySet1((func() _dafny.Sequence {
				if _this.F38 {
					return _1687_v9
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1695_v6 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_1696_i2 _dafny.Int) _dafny.Set {
						return _1695_v6
					}
				})(_1684_v6)))
			})(), 16)
			(_nw252).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1687_v9, _1687_v9), 17)
			_1690_v12 = _nw252
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1690_v12), 0))
			_ = _index239
			(_1690_v12).ArraySet1(_1687_v9, (_index239).Int())
			var _1697_v13 D3
			_ = _1697_v13
			_1697_v13 = Companion_D3_.Create_DC11_(_dafny.CodePoint('a'))
			var _1698_v14 _dafny.Sequence
			_ = _1698_v14
			_1698_v14 = _dafny.SeqOf(Companion_Default___.Fm41(!(_this.F38), _1697_v13, _1680_cf9, p0, globalState))
			var _1699_v15 *C5
			_ = _1699_v15
			var _nw253 *C5 = New_C5_()
			_ = _nw253
			_nw253.Ctor__((_1671_v3).Select((Companion_Default___.SafeIndex(_1680_cf9, _dafny.IntOfUint32((_1671_v3).Cardinality()))).Uint32()).(bool), true)
			_1699_v15 = _nw253
			var _1700_v16 _dafny.Map
			_ = _1700_v16
			_1700_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1699_v15, p1)
			var _1701_v17 D15
			_ = _1701_v17
			_1701_v17 = Companion_D15_.Create_DC33_(_1700_v16)
			var _pat_let_tv44 = _1700_v16
			_ = _pat_let_tv44
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1690_v12), 0))
			_ = _index240
			var _rhs255 _dafny.Int = _1681_cf8
			_ = _rhs255
			var _rhs256 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1687_v9, (_1698_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1679_cf10, (Companion_Default___.SafeIndex((_1670_v2).Fm4(Companion_D0_.Create_DC0_(p1), _1668_v0, ((func(_pat_let40_0 D15) D15 {
				return func(_1702_dt__update__tmp_h0 D15) D15 {
					return func(_pat_let41_0 _dafny.Map) D15 {
						return func(_1703_dt__update_hcf61_h0 _dafny.Map) D15 {
							return Companion_D15_.Create_DC33_(_1703_dt__update_hcf61_h0)
						}(_pat_let41_0)
					}(_pat_let_tv44)
				}(_pat_let40_0)
			}(_1701_v17)).Dtor_cf61()).Cardinality(), globalState), _dafny.IntOfUint32((_1679_cf10).Cardinality()))).Uint32(), p2)).Cardinality()), _dafny.IntOfUint32((_1698_v14).Cardinality()))).Uint32()).(_dafny.Sequence))
			_ = _rhs256
			var _lhs230 *GlobalState = globalState
			_ = _lhs230
			var _lhs231 _dafny.Array = _1690_v12
			_ = _lhs231
			var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1690_v12), 0))
			_ = _lhs232
			_lhs230.F16 = _rhs255
			(_lhs231).ArraySet1(_rhs256, (_lhs232).Int())
			var _1704_v18 _dafny.Map
			_ = _1704_v18
			_1704_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1668_v0, _1680_cf9)
			if ((_dafny.Zero).Minus((func() _dafny.Int {
				if p0 {
					return _1681_cf8
				}
				return _1668_v0
			})())).Cmp(((_1704_v18).Update(_1668_v0, _1681_cf8)).Cardinality()) == 0 {
				_1704_v18 = (_1704_v18).Update(_dafny.IntOfInt64(567), _1681_cf8)
				var _1705_v19 _dafny.Array
				_ = _1705_v19
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_44
				var _nw254 _dafny.Array
				_ = _nw254
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw254 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) _dafny.Int = (func(_1706_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1707_i3 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1707_i3, (_dafny.Zero).Minus(_1706_cf9))
						}
					})(_1680_cf9)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw254 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw254).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw254).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1705_v19 = _nw254
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1705_v19), 0))
				_ = _index241
				(_1705_v19).ArraySet1((_1699_v15).Fm1(globalState), (_index241).Int())
				var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1705_v19), 0))
				_ = _index242
				(_1705_v19).ArraySet1((_1669_v1).Select((Companion_Default___.SafeIndex((_1699_v15).Fm1(globalState), _dafny.IntOfUint32((_1669_v1).Cardinality()))).Uint32()).(_dafny.Int), (_index242).Int())
				var _1708_v20 *C0
				_ = _1708_v20
				var _nw255 *C0 = New_C0_()
				_ = _nw255
				_nw255.Ctor__((_1681_cf8).Cmp((_1705_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1705_v19), 0))).Int()).(_dafny.Int)) <= 0)
				_1708_v20 = _nw255
				var _rhs257 *C0 = _1708_v20
				_ = _rhs257
				var _rhs258 _dafny.Array = _1705_v19
				_ = _rhs258
				var _rhs259 _dafny.Sequence = ((_1670_v2).F35()).Select((Companion_Default___.SafeIndex(_1668_v0, _dafny.IntOfUint32(((_1670_v2).F35()).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs259
				_1708_v20 = _rhs257
				_1705_v19 = _rhs258
				_1669_v1 = _rhs259
				var _1709_v21 D15
				_ = _1709_v21
				_1709_v21 = Companion_D15_.Create_DC34_(_1681_cf8, p0, _dafny.IntOfUint32((_1678_cf11).Cardinality()))
				_1709_v21 = Companion_D15_.Create_DC34_((_1705_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(658), _dafny.ArrayLen((_1705_v19), 0))).Int()).(_dafny.Int), false, Companion_Default___.SafeDivisionInt(_1668_v0, _1680_cf9))
				var _1710_v22 *C5
				_ = _1710_v22
				var _nw256 *C5 = New_C5_()
				_ = _nw256
				_nw256.Ctor__(true, (true) || (true))
				_1710_v22 = _nw256
			} else {
				var _1711_v23 _dafny.Array
				_ = _1711_v23
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw257
				_1711_v23 = _nw257
				_1711_v23 = _1711_v23
				var _1712_v24 _dafny.Array
				_ = _1712_v24
				var _nwElement0_63 _dafny.Int = Companion_Default___.SafeDivisionInt(_1680_cf9, _1681_cf8)
				_ = _nwElement0_63
				var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(9))
				_ = _nw258
				(_nw258).ArraySet1(_nwElement0_63, 0)
				(_nw258).ArraySet1((func() _dafny.Int {
					if p1 {
						return _1680_cf9
					}
					return _1680_cf9
				})(), 1)
				(_nw258).ArraySet1(_1680_cf9, 2)
				(_nw258).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1671_v3).Cardinality()), _1680_cf9)), 3)
				(_nw258).ArraySet1(_dafny.IntOfInt64(836), 4)
				(_nw258).ArraySet1(_1681_cf8, 5)
				(_nw258).ArraySet1(_1668_v0, 6)
				(_nw258).ArraySet1(((_dafny.Zero).Minus(_1681_cf8)).Times(_dafny.IntOfUint32((_1679_cf10).Cardinality())), 7)
				(_nw258).ArraySet1(_1680_cf9, 8)
				_1712_v24 = _nw258
				var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1712_v24), 0))
				_ = _index243
				(_1712_v24).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1671_v3, (Companion_Default___.SafeIndex(_1668_v0, _dafny.IntOfUint32((_1671_v3).Cardinality()))).Uint32(), _this.F38)).Cardinality()), (_index243).Int())
				var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1712_v24), 0))
				_ = _index244
				(_1712_v24).ArraySet1(_1668_v0, (_index244).Int())
				var _1713_v25 _dafny.Map
				_ = _1713_v25
				_1713_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1711_v23, _1680_cf9)
				var _1714_v26 _dafny.Map
				_ = _1714_v26
				_1714_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1713_v25).Cardinality())
				(globalState).F12 = ((_1714_v26).Merge((func() _dafny.Map {
					if p1 {
						return _1714_v26
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1712_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1712_v24), 0))).Int()).(_dafny.Int))
				})())).Cardinality()
				var _1715_v27 _dafny.MultiSet
				_ = _1715_v27
				_1715_v27 = _dafny.MultiSetOf(p1, !(p1), _this.F38)
				var _1716_v28 _dafny.Set
				_ = _1716_v28
				_1716_v28 = _dafny.SetOf((_1712_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1712_v24), 0))).Int()).(_dafny.Int))
				var _1717_v29 D2
				_ = _1717_v29
				_1717_v29 = Companion_D2_.Create_DC9_((_1668_v0).Minus((_dafny.Zero).Minus(_1668_v0)), (_dafny.Zero).Minus(_1668_v0), (_1716_v28).IsSubsetOf(_dafny.SetOf((_1715_v27).Cardinality())))
				_1717_v29 = _1717_v29
				(globalState).F16 = _1680_cf9
			}
			(globalState).F22 = p0
		} else if _source26.Is_DC5() {
			var _1718___mcc_h5 _dafny.Int = _source26.Get_().(D1_DC5).Cf12
			_ = _1718___mcc_h5
			var _1719_cf12 _dafny.Int = _1718___mcc_h5
			_ = _1719_cf12
			r3 = p1
			var _1720_v30 _dafny.Set
			_ = _1720_v30
			_1720_v30 = _dafny.SetOf(_1719_cf12, _1668_v0)
			var _1721_v31 _dafny.Array
			_ = _1721_v31
			var _nwElement0_64 _dafny.CodePoint = p2
			_ = _nwElement0_64
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(2))
			_ = _nw259
			(_nw259).ArraySet1CodePoint(_nwElement0_64, 0)
			(_nw259).ArraySet1CodePoint(p2, 1)
			_1721_v31 = _nw259
			var _1722_v32 _dafny.Sequence
			_ = _1722_v32
			_1722_v32 = _dafny.UnicodeSeqOfUtf8Bytes("ryndpso")
			var _1723_v33 D1
			_ = _1723_v33
			_1723_v33 = Companion_D1_.Create_DC4_(_1721_v31, _1719_cf12, _1668_v0, _1722_v32, _1722_v32)
			var _1724_v34 _dafny.Map
			_ = _1724_v34
			_1724_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1725_v35 _dafny.Map
			_ = _1725_v35
			_1725_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1719_cf12, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vi")).Cardinality())))
			var _1726_v36 _dafny.Set
			_ = _1726_v36
			_1726_v36 = _dafny.SetOf(_this.F38, p0, true, p1)
			var _1727_v37 _dafny.MultiSet
			_ = _1727_v37
			_1727_v37 = _dafny.MultiSetOf(_1668_v0, _1719_cf12, _1719_cf12, _1719_cf12)
			var _1728_v38 _dafny.Array
			_ = _1728_v38
			var _nwElement0_65 _dafny.Int = (_dafny.Zero).Minus(_1668_v0)
			_ = _nwElement0_65
			var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(24))
			_ = _nw260
			(_nw260).ArraySet1(_nwElement0_65, 0)
			(_nw260).ArraySet1(_1719_cf12, 1)
			(_nw260).ArraySet1((_1723_v33).Dtor_cf8(), 2)
			(_nw260).ArraySet1(_1719_cf12, 3)
			(_nw260).ArraySet1((_dafny.Zero).Minus(_1668_v0), 4)
			(_nw260).ArraySet1(_1668_v0, 5)
			(_nw260).ArraySet1(((_1724_v34).Merge(_1724_v34)).Cardinality(), 6)
			(_nw260).ArraySet1(_1719_cf12, 7)
			(_nw260).ArraySet1(_1719_cf12, 8)
			(_nw260).ArraySet1(_1719_cf12, 9)
			(_nw260).ArraySet1((_1670_v2).Fm1(globalState), 10)
			(_nw260).ArraySet1(Companion_Default___.SafeModuloInt(_1719_cf12, _1719_cf12), 11)
			(_nw260).ArraySet1(_1668_v0, 12)
			(_nw260).ArraySet1((Companion_Default___.Fm42(_1668_v0, globalState)).Cardinality(), 13)
			(_nw260).ArraySet1(_1719_cf12, 14)
			(_nw260).ArraySet1(_1668_v0, 15)
			(_nw260).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()), 16)
			(_nw260).ArraySet1(_1719_cf12, 17)
			(_nw260).ArraySet1(Companion_Default___.SafeModuloInt((_1725_v35).Cardinality(), (_1726_v36).Cardinality()), 18)
			(_nw260).ArraySet1((_dafny.MultiSetOf(_1719_cf12, _1668_v0)).Cardinality(), 19)
			(_nw260).ArraySet1(_1668_v0, 20)
			(_nw260).ArraySet1((_1668_v0).Times((_dafny.Zero).Minus((_1727_v37).Cardinality())), 21)
			(_nw260).ArraySet1(_1668_v0, 22)
			(_nw260).ArraySet1(_1668_v0, 23)
			_1728_v38 = _nw260
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1728_v38), 0))
			_ = _index245
			(_1728_v38).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1669_v1, _1669_v1)).Cardinality()), (_index245).Int())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1728_v38), 0))
			_ = _index246
			var _rhs260 _dafny.Set = (_1720_v30).Difference(_1720_v30)
			_ = _rhs260
			var _rhs261 _dafny.Sequence = Companion_Default___.Fm43(_1668_v0, _dafny.IntOfInt64(909), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1671_v3, _1671_v3)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2), globalState)
			_ = _rhs261
			var _rhs262 bool = (p0) && (!(p0))
			_ = _rhs262
			var _rhs263 _dafny.Int = _1668_v0
			_ = _rhs263
			var _rhs264 _dafny.Int = (_1719_cf12).Plus(_1719_cf12)
			_ = _rhs264
			var _lhs233 *GlobalState = globalState
			_ = _lhs233
			var _lhs234 *GlobalState = globalState
			_ = _lhs234
			var _lhs235 _dafny.Array = _1728_v38
			_ = _lhs235
			var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen((_1728_v38), 0))
			_ = _lhs236
			_1720_v30 = _rhs260
			_1669_v1 = _rhs261
			_lhs233.F15 = _rhs262
			_lhs234.F16 = _rhs263
			(_lhs235).ArraySet1(_rhs264, (_lhs236).Int())
			var _1729_v39 _dafny.Array
			_ = _1729_v39
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(8))
			_ = _nw261
			_1729_v39 = _nw261
			_1729_v39 = _1729_v39
			var _1730_v40 D0
			_ = _1730_v40
			_1730_v40 = Companion_D0_.Create_DC0_(_this.F38)
			var _pat_let_tv45 = p0
			_ = _pat_let_tv45
			_1719_cf12 = (_1670_v2).Fm4(func(_pat_let42_0 D0) D0 {
				return func(_1731_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let43_0 bool) D0 {
						return func(_1732_dt__update_hcf0_h0 bool) D0 {
							return Companion_D0_.Create_DC0_(_1732_dt__update_hcf0_h0)
						}(_pat_let43_0)
					}(_pat_let_tv45)
				}(_pat_let42_0)
			}(_1730_v40), Companion_Default___.SafeModuloInt(_1668_v0, _1719_cf12), (_1719_cf12).Minus(_1719_cf12), globalState)
		} else {
			var _1733___mcc_h6 _dafny.Sequence = _source26.Get_().(D1_DC3).Cf6
			_ = _1733___mcc_h6
			var _1734_cf6 _dafny.Sequence = _1733___mcc_h6
			_ = _1734_cf6
			var _1735_v41 T0
			_ = _1735_v41
			var _nw262 *C6 = New_C6_()
			_ = _nw262
			_nw262.Ctor__(_1668_v0)
			_1735_v41 = _nw262
			var _1736_v42 _dafny.CodePoint
			_ = _1736_v42
			var _1737_v43 _dafny.Int
			_ = _1737_v43
			var _out53 _dafny.CodePoint
			_ = _out53
			var _out54 _dafny.Int
			_ = _out54
			_out53, _out54 = (_1670_v2).M2((_1668_v0).Cmp(_1668_v0) != 0, _1735_v41, _this.F38, globalState)
			_1736_v42 = _out53
			_1737_v43 = _out54
			if p1 {
				(globalState).F15 = p1
				(globalState).F22 = (_1735_v41).Fm2(p1, _1668_v0, globalState)
				var _1738_v44 _dafny.Map
				_ = _1738_v44
				_1738_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _1739_v45 D13
				_ = _1739_v45
				_1739_v45 = Companion_D13_.Create_DC29_(_1738_v44)
				var _1740_v46 _dafny.Set
				_ = _1740_v46
				_1740_v46 = _dafny.SetOf(_1739_v45, _1739_v45)
				(globalState).F15 = (_1740_v46).IsDisjointFrom(_1740_v46)
				var _1741_v47 _dafny.MultiSet
				_ = _1741_v47
				_1741_v47 = _dafny.MultiSetOf(p0)
				var _1742_v48 _dafny.Sequence
				_ = _1742_v48
				_1742_v48 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1738_v44).Cardinality(), false))
				(globalState).F24 = (((_dafny.Zero).Minus((_1741_v47).Cardinality())).Cmp(_dafny.IntOfUint32((_1742_v48).Cardinality())) == 0) == ((_dafny.IntOfUint32((_1734_cf6).Cardinality())).Cmp(_dafny.IntOfInt64(505)) == 0)
				_1668_v0 = _1737_v43
			} else {
				var _1743_v49 _dafny.Set
				_ = _1743_v49
				_1743_v49 = _dafny.SetOf(_1737_v43)
				var _1744_v50 _dafny.Map
				_ = _1744_v50
				_1744_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1668_v0, p1)
				var _1745_v51 _dafny.Sequence
				_ = _1745_v51
				_1745_v51 = _dafny.UnicodeSeqOfUtf8Bytes("bam")
				var _1746_v52 D5
				_ = _1746_v52
				_1746_v52 = Companion_D5_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1671_v3))
				var _1747_v53 D5
				_ = _1747_v53
				_1747_v53 = Companion_D5_.Create_DC15_(_1746_v52)
				var _1748_v54 _dafny.Map
				_ = _1748_v54
				_1748_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1747_v53, _1737_v43)
				var _1749_v55 _dafny.Array
				_ = _1749_v55
				var _nwElement0_66 _dafny.Int = _1737_v43
				_ = _nwElement0_66
				var _nw263 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(25))
				_ = _nw263
				(_nw263).ArraySet1(_nwElement0_66, 0)
				(_nw263).ArraySet1((_1743_v49).Cardinality(), 1)
				(_nw263).ArraySet1(_1668_v0, 2)
				(_nw263).ArraySet1(_1668_v0, 3)
				(_nw263).ArraySet1(_dafny.IntOfInt64(-377), 4)
				(_nw263).ArraySet1(_1668_v0, 5)
				(_nw263).ArraySet1(_1737_v43, 6)
				(_nw263).ArraySet1(_1668_v0, 7)
				(_nw263).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((func() bool {
					if (_1744_v50).Contains(_1737_v43) {
						return (_1744_v50).Get(_1737_v43).(bool)
					}
					return _this.F38
				})()), _1737_v43)).Cardinality(), 8)
				(_nw263).ArraySet1((_1735_v41).Fm1(globalState), 9)
				(_nw263).ArraySet1((_dafny.Zero).Minus(_1668_v0), 10)
				(_nw263).ArraySet1(_1737_v43, 11)
				(_nw263).ArraySet1(_1668_v0, 12)
				(_nw263).ArraySet1(_1668_v0, 13)
				(_nw263).ArraySet1((_dafny.Zero).Minus(_1737_v43), 14)
				(_nw263).ArraySet1(_1737_v43, 15)
				(_nw263).ArraySet1(_dafny.IntOfUint32((_1745_v51).Cardinality()), 16)
				(_nw263).ArraySet1((_1735_v41).Fm1(globalState), 17)
				(_nw263).ArraySet1((_1748_v54).Cardinality(), 18)
				(_nw263).ArraySet1(_1737_v43, 19)
				(_nw263).ArraySet1((_dafny.Zero).Minus(_1668_v0), 20)
				(_nw263).ArraySet1(_dafny.IntOfUint32((_1745_v51).Cardinality()), 21)
				(_nw263).ArraySet1(_1668_v0, 22)
				(_nw263).ArraySet1(_1737_v43, 23)
				(_nw263).ArraySet1(_1737_v43, 24)
				_1749_v55 = _nw263
				var _1750_v56 _dafny.MultiSet
				_ = _1750_v56
				_1750_v56 = _dafny.MultiSetOf(_this.F38, false, false)
				var _1751_v57 _dafny.Map
				_ = _1751_v57
				_1751_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1749_v55, _1750_v56)
				_1751_v57 = (_1751_v57).Merge(_1751_v57)
				(globalState).F22 = p0
				(globalState).F6 = _dafny.UnicodeSeqOfUtf8Bytes("qx")
				var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_1749_v55), 0))
				_ = _index247
				(_1749_v55).ArraySet1(_1668_v0, (_index247).Int())
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_1749_v55), 0))
				_ = _index248
				(_1749_v55).ArraySet1((Companion_D1_.Create_DC5_((_dafny.Zero).Minus(_1737_v43))).Dtor_cf12(), (_index248).Int())
				(globalState).F15 = _dafny.Companion_Sequence_.IsPrefixOf(_1745_v51, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(657))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}(func(_1752_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				})))
			}
			var _1753_v58 T2
			_ = _1753_v58
			var _nw264 *C1 = New_C1_()
			_ = _nw264
			_nw264.Ctor__(_1668_v0, _1668_v0, p1, p0)
			_1753_v58 = _nw264
			var _1754_v59 _dafny.MultiSet
			_ = _1754_v59
			_1754_v59 = _dafny.MultiSetOf(_1753_v58, _1753_v58, _1753_v58, _1753_v58)
			var _1755_v60 D15
			_ = _1755_v60
			_1755_v60 = Companion_D15_.Create_DC34_(_dafny.IntOfInt64(-234), p1, _1668_v0)
			var _1756_v61 _dafny.Map
			_ = _1756_v61
			_1756_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(944), (_1755_v60).Dtor_cf63())).Cardinality())
			var _1757_v62 _dafny.Set
			_ = _1757_v62
			_1757_v62 = _dafny.SetOf(true, p1, !((_1753_v58).F27()), (_1753_v58).F26(), (_1753_v58).F26())
			var _1758_v63 _dafny.MultiSet
			_ = _1758_v63
			_1758_v63 = _dafny.MultiSetOf(p2)
			var _1759_v64 _dafny.Array
			_ = _1759_v64
			var _nwElement0_67 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return _dafny.IntOfUint32((_1669_v1).Cardinality())
				}
				return _1737_v43
			})()
			_ = _nwElement0_67
			var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(8))
			_ = _nw265
			(_nw265).ArraySet1(_nwElement0_67, 0)
			(_nw265).ArraySet1(_1737_v43, 1)
			(_nw265).ArraySet1((_1754_v59).Cardinality(), 2)
			(_nw265).ArraySet1(_1737_v43, 3)
			(_nw265).ArraySet1(_dafny.IntOfUint32((_1734_cf6).Cardinality()), 4)
			(_nw265).ArraySet1(_dafny.IntOfInt64(405), 5)
			(_nw265).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_1756_v61).Contains((_1753_v58).F26()) {
					return (_1756_v61).Get((_1753_v58).F26()).(_dafny.Int)
				}
				return (_1757_v62).Cardinality()
			})(), _1668_v0), 6)
			(_nw265).ArraySet1(((Companion_Default___.Fm48(_dafny.IntOfUint32((_dafny.SeqOf(_1668_v0)).Cardinality()), globalState)).Union(_1758_v63)).Cardinality(), 7)
			_1759_v64 = _nw265
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1759_v64), 0))
			_ = _index249
			(_1759_v64).ArraySet1(_1737_v43, (_index249).Int())
			var _1760_v65 D0
			_ = _1760_v65
			_1760_v65 = Companion_D0_.Create_DC0_((_1734_cf6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}((func(_1761_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1762_i5 _dafny.Int) _dafny.CodePoint {
					return _1761_p2
				}
			})(p2)))).Cardinality()), _dafny.IntOfUint32((_1734_cf6).Cardinality()))).Uint32()).(bool))
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1759_v64), 0))
			_ = _index250
			var _rhs265 _dafny.Int = (_dafny.Zero).Minus((_1670_v2).Fm4(_1760_v65, _1737_v43, _1737_v43, globalState))
			_ = _rhs265
			var _rhs266 _dafny.Int = _dafny.IntOfUint32((_1734_cf6).Cardinality())
			_ = _rhs266
			var _lhs237 _dafny.Array = _1759_v64
			_ = _lhs237
			var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1759_v64), 0))
			_ = _lhs238
			(_lhs237).ArraySet1(_rhs265, (_lhs238).Int())
			_1737_v43 = _rhs266
			var _1763_v66 _dafny.MultiSet
			_ = _1763_v66
			_1763_v66 = _dafny.MultiSetOf(_dafny.IntOfInt64(-658))
			var _1764_v67 _dafny.MultiSet
			_ = _1764_v67
			_1764_v67 = _dafny.MultiSetOf((_1763_v66).Difference(_1763_v66))
			_1764_v67 = Companion_Default___.Fm49((_1763_v66).Update(_1737_v43, Companion_Default___.Abs(_dafny.IntOfInt64(-948))), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1734_cf6, (Companion_Default___.SafeIndex(_1668_v0, _dafny.IntOfUint32((_1734_cf6).Cardinality()))).Uint32(), p1)).Cardinality()), (_1753_v58).F27(), globalState)
		}
		var _hi14 _dafny.Int = _1668_v0
		_ = _hi14
		for _1765_i6 := (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(206), _1668_v0)).Cardinality(); _1765_i6.Cmp(_hi14) < 0; _1765_i6 = _1765_i6.Plus(_dafny.One) {
			var _1766_v68 _dafny.MultiSet
			_ = _1766_v68
			_1766_v68 = _dafny.MultiSetOf(_1668_v0)
			var _1767_v69 _dafny.Set
			_ = _1767_v69
			_1767_v69 = _dafny.SetOf(p1, _this.F38)
			var _1768_v70 _dafny.Map
			_ = _1768_v70
			_1768_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1668_v0), p1)
			var _1769_v71 _dafny.MultiSet
			_ = _1769_v71
			_1769_v71 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1766_v68).Cardinality(), p1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1767_v69).Cardinality(), p0), _1768_v70, _1768_v70)
			var _1770_v72 _dafny.Map
			_ = _1770_v72
			_1770_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F38), _1769_v71)
			r3 = !(((_1769_v71).Intersection((func() _dafny.MultiSet {
				if (_1770_v72).Contains(p0) {
					return (_1770_v72).Get(p0).(_dafny.MultiSet)
				}
				return _1769_v71
			})())).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1668_v0, p1))))
			var _1771_v73 _dafny.MultiSet
			_ = _1771_v73
			_1771_v73 = _dafny.MultiSetOf(true)
			var _1772_v74 _dafny.Map
			_ = _1772_v74
			_1772_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, p0)
			var _1773_v75 _dafny.Array
			_ = _1773_v75
			var _nwElement0_68 bool = p0
			_ = _nwElement0_68
			var _nw266 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(21))
			_ = _nw266
			(_nw266).ArraySet1(_nwElement0_68, 0)
			(_nw266).ArraySet1(p0, 1)
			(_nw266).ArraySet1(p1, 2)
			(_nw266).ArraySet1(true, 3)
			(_nw266).ArraySet1(p0, 4)
			(_nw266).ArraySet1(p0, 5)
			(_nw266).ArraySet1((_1671_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1771_v73).Contains(_this.F38) {
					return (_1771_v73).Multiplicity(_this.F38)
				}
				return _1668_v0
			})(), _dafny.IntOfUint32((_1671_v3).Cardinality()))).Uint32()).(bool), 6)
			(_nw266).ArraySet1(_this.F38, 7)
			(_nw266).ArraySet1(_this.F38, 8)
			(_nw266).ArraySet1(p0, 9)
			(_nw266).ArraySet1(!(true), 10)
			(_nw266).ArraySet1(p1, 11)
			(_nw266).ArraySet1(_this.F38, 12)
			(_nw266).ArraySet1(_this.F38, 13)
			(_nw266).ArraySet1((_1670_v2).Fm2((func() bool {
				if (_1772_v74).Contains(_this.F38) {
					return (_1772_v74).Get(_this.F38).(bool)
				}
				return _this.F38
			})(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-511))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}((func(_1774_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1775_i7 _dafny.Int) _dafny.CodePoint {
					return _1774_p2
				}
			})(p2)))).Cardinality()), globalState), 14)
			(_nw266).ArraySet1(p1, 15)
			(_nw266).ArraySet1(p1, 16)
			(_nw266).ArraySet1(p1, 17)
			(_nw266).ArraySet1(p0, 18)
			(_nw266).ArraySet1(true, 19)
			(_nw266).ArraySet1(!((Companion_D2_.Create_DC7_(p1, true, _1668_v0, _1668_v0)).Dtor_cf14()), 20)
			_1773_v75 = _nw266
			var _1776_v76 D14
			_ = _1776_v76
			_1776_v76 = Companion_D14_.Create_DC32_(_1771_v73, p1, p0, p1, (_1767_v69).Cardinality())
			var _1777_v77 _dafny.Map
			_ = _1777_v77
			_1777_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1776_v76, _this.F38)
			var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1773_v75), 0))
			_ = _index251
			(_1773_v75).ArraySet1((_1670_v2).Fm17(_1668_v0, _1765_i6, (_1777_v77).Cardinality(), Companion_Default___.Fm50(globalState), globalState), (_index251).Int())
			var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1773_v75), 0))
			_ = _index252
			var _rhs267 bool = p1
			_ = _rhs267
			var _rhs268 _dafny.Int = _1765_i6
			_ = _rhs268
			var _lhs239 _dafny.Array = _1773_v75
			_ = _lhs239
			var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1773_v75), 0))
			_ = _lhs240
			var _lhs241 *GlobalState = globalState
			_ = _lhs241
			(_lhs239).ArraySet1(_rhs267, (_lhs240).Int())
			_lhs241.F16 = _rhs268
			var _1778_v78 T0
			_ = _1778_v78
			var _nw267 *C6 = New_C6_()
			_ = _nw267
			_nw267.Ctor__((_1766_v68).Cardinality())
			_1778_v78 = _nw267
			var _1779_v79 _dafny.CodePoint
			_ = _1779_v79
			var _1780_v80 _dafny.Int
			_ = _1780_v80
			var _out55 _dafny.CodePoint
			_ = _out55
			var _out56 _dafny.Int
			_ = _out56
			_out55, _out56 = (_1670_v2).M2(p0, _1778_v78, !((func() bool {
				if p1 {
					return _this.F38
				}
				return p0
			})()), globalState)
			_1779_v79 = _out55
			_1780_v80 = _out56
			var _1781_v81 _dafny.Map
			_ = _1781_v81
			_1781_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1668_v0)
			var _1782_v82 D7
			_ = _1782_v82
			_1782_v82 = Companion_D7_.Create_DC19_(_1668_v0, _dafny.MultiSetOf(p1, p0), (_1773_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1773_v75), 0))).Int()).(bool), _1781_v81)
			var _1783_v83 _dafny.Map
			_ = _1783_v83
			_1783_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F38), (_1782_v82).Dtor_cf36())
			var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_1773_v75), 0))
			_ = _index253
			(_1773_v75).ArraySet1((func() bool {
				if (_1768_v70).Contains(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(641), (_dafny.Zero).Minus((_1783_v83).Cardinality()))) {
					return (_1768_v70).Get(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(641), (_dafny.Zero).Minus((_1783_v83).Cardinality()))).(bool)
				}
				return p0
			})(), (_index253).Int())
		}
		(globalState).F16 = _1668_v0
		var _1784_v84 _dafny.Array
		_ = _1784_v84
		var _nw268 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
		_ = _nw268
		_1784_v84 = _nw268
		_1784_v84 = _1784_v84
		var _1785_v85 _dafny.Sequence
		_ = _1785_v85
		_1785_v85 = _dafny.UnicodeSeqOfUtf8Bytes("hcjrgaf")
		if _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("tixwcfv"), _1785_v85) {
			var _1786_v86 _dafny.Array
			_ = _1786_v86
			var _nw269 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw269
			_1786_v86 = _nw269
			var _1787_v87 D3
			_ = _1787_v87
			_1787_v87 = Companion_D3_.Create_DC10_(_1786_v86)
			var _1788_v88 _dafny.Map
			_ = _1788_v88
			_1788_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1787_v87, _dafny.IntOfUint32((_dafny.SeqOf(_1668_v0)).Cardinality()))
			var _1789_v89 *C4
			_ = _1789_v89
			var _nw270 *C4 = New_C4_()
			_ = _nw270
			_nw270.Ctor__(_1788_v88, p1, _this.F38)
			_1789_v89 = _nw270
			var _1790_v90 _dafny.Sequence
			_ = _1790_v90
			_1790_v90 = _dafny.SeqOf(_1785_v85, _1785_v85, _1785_v85, _1785_v85)
			(_1670_v2).M1((func() bool {
				if !(!(false)) {
					return true
				}
				return _this.F38
			})(), _1790_v90, _dafny.Companion_Sequence_.Concatenate(_1671_v3, _1671_v3), globalState)
			var _1791_v91 _dafny.Map
			_ = _1791_v91
			_1791_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}(func(_1792_i8 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})))
			var _rhs269 _dafny.Map = _1791_v91
			_ = _rhs269
			var _rhs270 bool = true
			_ = _rhs270
			var _rhs271 _dafny.Int = _1668_v0
			_ = _rhs271
			var _rhs272 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1785_v85, _1785_v85)
			_ = _rhs272
			var _lhs242 *C7 = _this
			_ = _lhs242
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			var _lhs244 *GlobalState = globalState
			_ = _lhs244
			_1791_v91 = _rhs269
			_lhs242.F38 = _rhs270
			_lhs243.F12 = _rhs271
			_lhs244.F6 = _rhs272
			r0 = false
			var _1793_v92 _dafny.Array
			_ = _1793_v92
			var _nwElement0_69 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("pnu")
			_ = _nwElement0_69
			var _nw271 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(9))
			_ = _nw271
			(_nw271).ArraySet1(_nwElement0_69, 0)
			(_nw271).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1785_v85, _dafny.UnicodeSeqOfUtf8Bytes("uxsmk")), 1)
			(_nw271).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1785_v85, _1785_v85), 2)
			(_nw271).ArraySet1(_1785_v85, 3)
			(_nw271).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("m"), 4)
			(_nw271).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fws"), 5)
			(_nw271).ArraySet1(_1785_v85, 6)
			(_nw271).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cobamhx"), 7)
			(_nw271).ArraySet1(_1785_v85, 8)
			_1793_v92 = _nw271
			_1793_v92 = _1793_v92
		} else {
			var _1794_v93 _dafny.Map
			_ = _1794_v93
			_1794_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1668_v0, _dafny.IntOfInt64(223))
			var _1795_v94 _dafny.Map
			_ = _1795_v94
			_1795_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1784_v84, (_dafny.MultiSetOf(false, _this.F38, _this.F38, p1, p1)).Cardinality())
			var _1796_v95 _dafny.Set
			_ = _1796_v95
			_1796_v95 = _dafny.SetOf(_this.F38, !(true), p0, true, false)
			var _1797_v96 _dafny.Array
			_ = _1797_v96
			var _nwElement0_70 _dafny.Int = _1668_v0
			_ = _nwElement0_70
			var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(18))
			_ = _nw272
			(_nw272).ArraySet1(_nwElement0_70, 0)
			(_nw272).ArraySet1(((_1794_v93).Cardinality()).Minus(_dafny.IntOfInt64(647)), 1)
			(_nw272).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1785_v85).Cardinality()))), 2)
			(_nw272).ArraySet1(_1668_v0, 3)
			(_nw272).ArraySet1(_1668_v0, 4)
			(_nw272).ArraySet1(Companion_Default___.SafeDivisionInt(_1668_v0, _1668_v0), 5)
			(_nw272).ArraySet1(_dafny.IntOfInt64(834), 6)
			(_nw272).ArraySet1(_1668_v0, 7)
			(_nw272).ArraySet1((_1668_v0).Plus(_1668_v0), 8)
			(_nw272).ArraySet1((func() _dafny.Int {
				if (_1795_v94).Contains(_1784_v84) {
					return (_1795_v94).Get(_1784_v84).(_dafny.Int)
				}
				return _1668_v0
			})(), 9)
			(_nw272).ArraySet1(_1668_v0, 10)
			(_nw272).ArraySet1(((Companion_D16_.Create_DC38_(_1668_v0, _1668_v0, _this.F38)).Dtor_cf70()).Minus((_dafny.Zero).Minus(_1668_v0)), 11)
			(_nw272).ArraySet1(_dafny.IntOfInt64(174), 12)
			(_nw272).ArraySet1(_1668_v0, 13)
			(_nw272).ArraySet1((_1796_v95).Cardinality(), 14)
			(_nw272).ArraySet1(_1668_v0, 15)
			(_nw272).ArraySet1(_1668_v0, 16)
			(_nw272).ArraySet1((func() _dafny.Int {
				if !(_this.F38) {
					return _1668_v0
				}
				return _1668_v0
			})(), 17)
			_1797_v96 = _nw272
			var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))
			_ = _index254
			(_1797_v96).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1668_v0), _1668_v0), (_index254).Int())
			var _1798_v97 D0
			_ = _1798_v97
			_1798_v97 = Companion_D0_.Create_DC0_(p0)
			var _1799_v98 D0
			_ = _1799_v98
			_1799_v98 = Companion_D0_.Create_DC2_(_1798_v97)
			var _1800_v99 _dafny.Sequence
			_ = _1800_v99
			_1800_v99 = _dafny.SeqOf(_1799_v98)
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))
			_ = _index255
			var _rhs273 _dafny.Int = _1668_v0
			_ = _rhs273
			var _rhs274 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1668_v0), _dafny.IntOfInt64(-14)), _1668_v0)
			_ = _rhs274
			var _rhs275 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1785_v85, _1785_v85)).Cardinality()))).Times((_1670_v2).Fm1(globalState))
			_ = _rhs275
			var _rhs276 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(63))).Uint32(), func(coer105 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1801_v98 D0) func(_dafny.Int) D0 {
				return func(_1802_i9 _dafny.Int) D0 {
					return _1801_v98
				}
			})(_1799_v98)))
			_ = _rhs276
			var _rhs277 _dafny.Int = _1668_v0
			_ = _rhs277
			var _lhs245 *GlobalState = globalState
			_ = _lhs245
			var _lhs246 _dafny.Array = _1797_v96
			_ = _lhs246
			var _lhs247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))
			_ = _lhs247
			var _lhs248 *GlobalState = globalState
			_ = _lhs248
			_1668_v0 = _rhs273
			_lhs245.F12 = _rhs274
			(_lhs246).ArraySet1(_rhs275, (_lhs247).Int())
			_1800_v99 = _rhs276
			_lhs248.F12 = _rhs277
			(globalState).F22 = p1
			var _1803_v100 _dafny.Array
			_ = _1803_v100
			var _nw273 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw273
			_1803_v100 = _nw273
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1803_v100), 0))
			_ = _index256
			(_1803_v100).ArraySet1(_1797_v96, (_index256).Int())
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1803_v100), 0))
			_ = _index257
			(_1803_v100).ArraySet1(_1797_v96, (_index257).Int())
			_1796_v95 = (_1796_v95).Union(_dafny.SetOf(p1))
			var _1804_v101 _dafny.Map
			_ = _1804_v101
			_1804_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int))
			if (func() bool {
				if true {
					return (_1804_v101).Equals(_1804_v101)
				}
				return _this.F38
			})() {
				(globalState).F24 = ((_this.F38) && (false)) || (p1)
				(globalState).F15 = _this.F38
				var _1805_v102 _dafny.Map
				_ = _1805_v102
				_1805_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
				var _1806_v103 _dafny.Map
				_ = _1806_v103
				_1806_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1803_v100).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1803_v100), 0))).Int())), p0)
				_1805_v102 = (_1805_v102).Update((func() bool {
					if (_1806_v103).Contains(_1797_v96) {
						return (_1806_v103).Get(_1797_v96).(bool)
					}
					return p0
				})(), _this.F38)
				_1784_v84 = _1784_v84
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))
				_ = _index258
				(_1797_v96).ArraySet1((func() _dafny.Int {
					if true {
						return (_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int)
					}
					return (_1668_v0).Minus(_1668_v0)
				})(), (_index258).Int())
			} else {
				var _1807_v104 _dafny.Set
				_ = _1807_v104
				_1807_v104 = _dafny.SetOf(((_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int)).Times((_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int)), _1668_v0, _dafny.IntOfInt64(177), _dafny.IntOfInt64(986), (_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int))
				var _1808_v105 _dafny.Sequence
				_ = _1808_v105
				_1808_v105 = _dafny.SeqOf(_1785_v85, _1785_v85, _1785_v85)
				var _1809_v106 _dafny.Map
				_ = _1809_v106
				_1809_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1808_v105, _1808_v105)), _dafny.SetOf((_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int), _1668_v0))
				var _1810_v107 _dafny.MultiSet
				_ = _1810_v107
				_1810_v107 = _dafny.MultiSetOf(_1785_v85)
				_1807_v104 = (func() _dafny.Set {
					if (_1809_v106).Contains((_1810_v107).Update(_1785_v85, Companion_Default___.Abs(_1668_v0))) {
						return (_1809_v106).Get((_1810_v107).Update(_1785_v85, Companion_Default___.Abs(_1668_v0))).(_dafny.Set)
					}
					return _1807_v104
				})()
				var _1811_v108 D17
				_ = _1811_v108
				_1811_v108 = Companion_D17_.Create_DC40_((_1670_v2).F35())
				var _1812_v109 _dafny.Map
				_ = _1812_v109
				_1812_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqOf((_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int), _1668_v0))
				var _1813_v110 _dafny.MultiSet
				_ = _1813_v110
				_1813_v110 = _dafny.MultiSetOf(_1669_v1)
				var _1814_v111 _dafny.Sequence
				_ = _1814_v111
				_1814_v111 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update((_1811_v108).Dtor_cf73(), (Companion_Default___.SafeIndex((_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_1811_v108).Dtor_cf73()).Cardinality()))).Uint32(), _dafny.SeqOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1812_v109).Contains(p2) {
						return (_1812_v109).Get(p2).(_dafny.Sequence)
					}
					return _1669_v1
				})()).Cardinality())))), _1813_v110, _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg106 _dafny.Int) interface{} {
						return coer106(arg106)
					}
				}((func(_1815_v2 *C2, _1816_v0 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1817_i10 _dafny.Int) _dafny.Sequence {
						return ((_1815_v2).F35()).Select((Companion_Default___.SafeIndex(_1816_v0, _dafny.IntOfUint32(((_1815_v2).F35()).Cardinality()))).Uint32()).(_dafny.Sequence)
					}
				})(_1670_v2, _1668_v0)))), _1813_v110, _1813_v110)
				var _1818_v112 *C3
				_ = _1818_v112
				var _nw274 *C3 = New_C3_()
				_ = _nw274
				_nw274.Ctor__((_1814_v111).Select((Companion_Default___.SafeIndex(_1668_v0, _dafny.IntOfUint32((_1814_v111).Cardinality()))).Uint32()).(_dafny.MultiSet), _this.F38, p1)
				_1818_v112 = _nw274
				var _1819_v113 _dafny.Sequence
				_ = _1819_v113
				_1819_v113 = _dafny.SeqOf(_1794_v93, _1794_v93)
				_1819_v113 = _1819_v113
				(globalState).F16 = Companion_Default___.SafeDivisionInt(_1668_v0, (_1797_v96).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(907), _dafny.ArrayLen((_1797_v96), 0))).Int()).(_dafny.Int))
				var _1820_v114 *C2
				_ = _1820_v114
				var _nw275 *C2 = New_C2_()
				_ = _nw275
				_nw275.Ctor__((_1670_v2).F35(), false, false)
				_1820_v114 = _nw275
			}
		}
		r0 = p1
		var _1821_v115 _dafny.Map
		_ = _1821_v115
		_1821_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, _1668_v0)
		var _1822_v116 _dafny.MultiSet
		_ = _1822_v116
		_1822_v116 = _dafny.MultiSetOf(_1821_v115, _1821_v115)
		var _1823_v117 D1
		_ = _1823_v117
		_1823_v117 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(226))
		r1 = ((_1822_v116).Update(_1821_v115, Companion_Default___.Abs((_1823_v117).Dtor_cf12()))).Intersection((_dafny.MultiSetOf(_1821_v115)).Difference(_dafny.MultiSetOf(_1821_v115, _1821_v115)))
		var _1824_v118 _dafny.Map
		_ = _1824_v118
		_1824_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F38, p2)
		var _1825_v119 _dafny.Map
		_ = _1825_v119
		_1825_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1671_v3, _1671_v3), ((func() _dafny.Map {
			if p1 {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			}
			return _1824_v118
		})()).Cardinality())
		r2 = _1825_v119
		r3 = p1
		return r0, r1, r2, r3
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f29 _dafny.MultiSet
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f29 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_}
}

var _ T3 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F29() _dafny.MultiSet {
	return _this._f29
}
func (_this *C8) Ctor__(f29 _dafny.MultiSet) {
	{
		(_this)._f29 = f29
	}
}
func (_this *C8) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Array, T0) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 T0 = (T0)(nil)
		_ = r3
		(globalState).F16 = p0
		var _hi15 _dafny.Int = p0
		_ = _hi15
		for _1826_i0 := p0; _1826_i0.Cmp(_hi15) < 0; _1826_i0 = _1826_i0.Plus(_dafny.One) {
			var _1827_v0 _dafny.CodePoint
			_ = _1827_v0
			_1827_v0 = _dafny.CodePoint('y')
			if (Companion_Default___.SafeDivisionInt(p3, _dafny.IntOfInt64(-589))).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), _1827_v0)).Cardinality()))) == 0 {
				var _1828_v1 bool
				_ = _1828_v1
				_1828_v1 = false
				var _1829_v3 _dafny.Map
				_ = _1829_v3
				_1829_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1826_i0)
				var _1830_v4 _dafny.Set
				_ = _1830_v4
				_1830_v4 = _dafny.SetOf(_dafny.IntOfInt64(-207), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Map {
					var _coll67 = _dafny.NewMapBuilder()
					_ = _coll67
					for _iter67 := _dafny.Iterate((_1829_v3).Keys().Elements()); ; {
						_compr_67, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1831_v2 _dafny.Int
						_1831_v2 = interface{}(_compr_67).(_dafny.Int)
						if (_1829_v3).Contains(_1831_v2) {
							_coll67.Add(Companion_Default___.SafeDivisionInt(_1831_v2, _1826_i0), _1826_i0)
						}
					}
					return _coll67.ToMap()
				}()).Cardinality())).Cardinality(), p3)
				var _1832_v5 _dafny.Map
				_ = _1832_v5
				_1832_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1828_v1, _1830_v4)
				var _1833_v6 _dafny.Map
				_ = _1833_v6
				_1833_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1826_i0).Times(p0), (_1832_v5).Merge(_1832_v5))
				_1833_v6 = (_1833_v6).Update(_dafny.IntOfUint32((p2).Cardinality()), (_1832_v5).Merge(_1832_v5))
				_1829_v3 = (_1829_v3).Update(Companion_Default___.Fm39(globalState), p1)
				var _1834_v7 _dafny.Array
				_ = _1834_v7
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_45
				var _nw276 _dafny.Array
				_ = _nw276
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw276 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.Int = (func(_1835_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1836_i1 _dafny.Int) _dafny.Int {
							return (_1836_i1).Plus(_1835_p1)
						}
					})(p1)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw276 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw276).ArraySet1(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw276).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1834_v7 = _nw276
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_1834_v7), 0))
				_ = _index259
				(_1834_v7).ArraySet1((_dafny.Zero).Minus(p3), (_index259).Int())
				var _1837_v8 _dafny.Array
				_ = _1837_v8
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_46
				var _nw277 _dafny.Array
				_ = _nw277
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw277 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) bool = (func(_1838_v1 bool) func(_dafny.Int) bool {
						return func(_1839_i2 _dafny.Int) bool {
							return _1838_v1
						}
					})(_1828_v1)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw277 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw277).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw277).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1837_v8 = _nw277
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1837_v8), 0))
				_ = _index260
				(_1837_v8).ArraySet1(!((_1828_v1) && (_1828_v1)), (_index260).Int())
				var _1840_v9 _dafny.Map
				_ = _1840_v9
				_1840_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)
				var _1841_v10 D13
				_ = _1841_v10
				_1841_v10 = Companion_D13_.Create_DC29_(_1840_v9)
				var _1842_v11 _dafny.Sequence
				_ = _1842_v11
				_1842_v11 = _dafny.SeqOf(_1826_i0, _dafny.IntOfUint32((p2).Cardinality()), _1826_i0, p3, p0)
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_1834_v7), 0))
				_ = _index261
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1837_v8), 0))
				_ = _index262
				var _rhs278 _dafny.Int = (_dafny.IntOfInt64(109)).Times((((_1841_v10).Dtor_cf50()).Cardinality()).Times(_1826_i0))
				_ = _rhs278
				var _rhs279 bool = _1828_v1
				_ = _rhs279
				var _rhs280 _dafny.Int = (p0).Times(p0)
				_ = _rhs280
				var _rhs281 bool = (func() bool {
					if !(_1828_v1) {
						return _dafny.Companion_Sequence_.Contains(_1842_v11, p3)
					}
					return (_1828_v1) && (_1828_v1)
				})()
				_ = _rhs281
				var _rhs282 _dafny.Int = ((_1826_i0).Plus(p1)).Plus((func() _dafny.Int {
					if _1828_v1 {
						return _dafny.IntOfInt64(578)
					}
					return _1826_i0
				})())
				_ = _rhs282
				var _lhs249 _dafny.Array = _1834_v7
				_ = _lhs249
				var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_1834_v7), 0))
				_ = _lhs250
				var _lhs251 _dafny.Array = _1837_v8
				_ = _lhs251
				var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1837_v8), 0))
				_ = _lhs252
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				var _lhs254 *GlobalState = globalState
				_ = _lhs254
				var _lhs255 *GlobalState = globalState
				_ = _lhs255
				(_lhs249).ArraySet1(_rhs278, (_lhs250).Int())
				(_lhs251).ArraySet1(_rhs279, (_lhs252).Int())
				_lhs253.F16 = _rhs280
				_lhs254.F15 = _rhs281
				_lhs255.F12 = _rhs282
				(globalState).F22 = (_1837_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1837_v8), 0))).Int()).(bool)
				var _1843_v12 *C0
				_ = _1843_v12
				var _nw278 *C0 = New_C0_()
				_ = _nw278
				_nw278.Ctor__(!((_1837_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1837_v8), 0))).Int()).(bool)))
				_1843_v12 = _nw278
			} else {
				var _1844_v13 bool
				_ = _1844_v13
				_1844_v13 = true
				(globalState).F22 = !((_1844_v13) && (_1844_v13))
				(globalState).F16 = p1
				var _1845_v14 _dafny.Set
				_ = _1845_v14
				_1845_v14 = _dafny.SetOf(_1827_v0)
				var _1846_v15 *C1
				_ = _1846_v15
				var _nw279 *C1 = New_C1_()
				_ = _nw279
				_nw279.Ctor__(p3, _1826_i0, _1844_v13, (_dafny.IntOfUint32((p2).Cardinality())).Cmp((_1845_v14).Cardinality()) == 0)
				_1846_v15 = _nw279
				var _1847_v16 _dafny.Array
				_ = _1847_v16
				var _nw280 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
				_ = _nw280
				_1847_v16 = _nw280
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1847_v16), 0))
				_ = _index263
				(_1847_v16).ArraySet1(true, (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_1847_v16), 0))
				_ = _index264
				(_1847_v16).ArraySet1(!(_1844_v13), (_index264).Int())
				(globalState).F16 = _1846_v15.F37
			}
			(globalState).F12 = _1826_i0
			var _1848_v17 _dafny.Sequence
			_ = _1848_v17
			_1848_v17 = _dafny.SeqOf(_1826_i0, Companion_Default___.Fm39(globalState))
			(globalState).F16 = _dafny.IntOfUint32((_1848_v17).Cardinality())
			(globalState).F19 = _1827_v0
		}
		var _1849_v18 _dafny.Array
		_ = _1849_v18
		var _len0_47 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_47
		var _nw281 _dafny.Array
		_ = _nw281
		if _len0_47.Cmp(_dafny.Zero) == 0 {
			_nw281 = _dafny.NewArray(_len0_47)
		} else {
			var _init47 func(_dafny.Int) _dafny.Int = (func(_1850_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1851_i3 _dafny.Int) _dafny.Int {
					return (_1851_i3).Times(_1850_p1)
				}
			})(p1)
			_ = _init47
			var _element0_47 = _init47(_dafny.Zero)
			_ = _element0_47
			_nw281 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
			(_nw281).ArraySet1(_element0_47, 0)
			var _nativeLen0_47 = (_len0_47).Int()
			_ = _nativeLen0_47
			for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
				(_nw281).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
			}
		}
		_1849_v18 = _nw281
		var _1852_v19 _dafny.Sequence
		_ = _1852_v19
		_1852_v19 = _dafny.SeqOf(_1849_v18, _1849_v18, _1849_v18)
		var _1853_v20 D3
		_ = _1853_v20
		_1853_v20 = Companion_D3_.Create_DC10_((_1852_v19).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1852_v19).Cardinality()))).Uint32()).(_dafny.Array))
		var _pat_let_tv46 = _1849_v18
		_ = _pat_let_tv46
		var _1854_v21 _dafny.Map
		_ = _1854_v21
		_1854_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let44_0 D3) D3 {
			return func(_1855_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let45_0 _dafny.Array) D3 {
					return func(_1856_dt__update_hcf21_h0 _dafny.Array) D3 {
						return Companion_D3_.Create_DC10_(_1856_dt__update_hcf21_h0)
					}(_pat_let45_0)
				}(_pat_let_tv46)
			}(_pat_let44_0)
		}(_1853_v20), p1)
		var _1857_v22 _dafny.Set
		_ = _1857_v22
		_1857_v22 = _dafny.SetOf(_1849_v18)
		var _1858_v23 bool
		_ = _1858_v23
		_1858_v23 = true
		var _1859_v24 *C4
		_ = _1859_v24
		var _nw282 *C4 = New_C4_()
		_ = _nw282
		_nw282.Ctor__(_1854_v21, (_1857_v22).IsDisjointFrom(_dafny.SetOf(_1849_v18, _1849_v18, _1849_v18, _1849_v18)), _1858_v23)
		_1859_v24 = _nw282
		var _1860_v25 _dafny.CodePoint
		_ = _1860_v25
		_1860_v25 = _dafny.CodePoint('f')
		var _1861_v26 _dafny.CodePoint
		_ = _1861_v26
		var _1862_v27 _dafny.Sequence
		_ = _1862_v27
		var _1863_v28 bool
		_ = _1863_v28
		var _1864_v29 _dafny.Map
		_ = _1864_v29
		var _out57 _dafny.CodePoint
		_ = _out57
		var _out58 _dafny.Sequence
		_ = _out58
		var _out59 bool
		_ = _out59
		var _out60 _dafny.Map
		_ = _out60
		_out57, _out58, _out59, _out60 = Companion_Default___.M0(_1860_v25, globalState)
		_1861_v26 = _out57
		_1862_v27 = _out58
		_1863_v28 = _out59
		_1864_v29 = _out60
		var _1865_v30 _dafny.Map
		_ = _1865_v30
		_1865_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1858_v23, p1)
		var _hi16 _dafny.Int = (_1865_v30).Cardinality()
		_ = _hi16
		for _1866_i4 := (_dafny.IntOfUint32((p2).Cardinality())).Minus(p1); _1866_i4.Cmp(_hi16) < 0; _1866_i4 = _1866_i4.Plus(_dafny.One) {
			(globalState).F16 = _1866_i4
			(globalState).F16 = (_1859_v24).Fm1(globalState)
			_1862_v27 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1858_v23), _1862_v27), _1862_v27)
			var _1867_v31 _dafny.Sequence
			_ = _1867_v31
			_1867_v31 = _dafny.SeqOf(p3)
			var _1868_v32 _dafny.MultiSet
			_ = _1868_v32
			_1868_v32 = _dafny.MultiSetOf(_1867_v31)
			var _1869_v33 T1
			_ = _1869_v33
			var _nw283 *C3 = New_C3_()
			_ = _nw283
			_nw283.Ctor__(_1868_v32, _1863_v28, _1858_v23)
			_1869_v33 = _nw283
			var _1870_v34 _dafny.Map
			_ = _1870_v34
			_1870_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1869_v33, (_this).F29())
			var _1871_v35 _dafny.Map
			_ = _1871_v35
			_1871_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1870_v34, (p1).Cmp(_dafny.IntOfInt64(76)) < 0)
			_1871_v35 = (_1871_v35).Update((_1870_v34).Update(_1869_v33, (_this).F29()), (_1862_v27).Select((Companion_Default___.SafeIndex((_1869_v33).Fm1(globalState), _dafny.IntOfUint32((_1862_v27).Cardinality()))).Uint32()).(bool))
		}
		var _1872_v36 _dafny.Set
		_ = _1872_v36
		_1872_v36 = _dafny.SetOf(_1863_v28)
		var _rhs283 bool = ((_1872_v36).Intersection(_1872_v36)).IsProperSubsetOf(_1872_v36)
		_ = _rhs283
		var _rhs284 _dafny.Int = p0
		_ = _rhs284
		var _lhs256 *GlobalState = globalState
		_ = _lhs256
		r0 = _rhs283
		_lhs256.F16 = _rhs284
		r0 = !(_1858_v23) || (_1858_v23)
		r1 = !((_1858_v23) || (_1858_v23))
		r2 = (_1852_v19).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1852_v19).Cardinality()))).Uint32()).(_dafny.Array)
		var _1873_v38 T0
		_ = _1873_v38
		var _nw284 *C7 = New_C7_()
		_ = _nw284
		_nw284.Ctor__(_1858_v23, func() _dafny.Map {
			var _coll68 = _dafny.NewMapBuilder()
			_ = _coll68
			for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-329), _dafny.IntOfInt64(305))); ; {
				_compr_68, _ok68 := _iter68()
				if !_ok68 {
					break
				}
				var _1874_v37 _dafny.Int
				_1874_v37 = interface{}(_compr_68).(_dafny.Int)
				if ((_dafny.IntOfInt64(-329)).Cmp(_1874_v37) <= 0) && ((_1874_v37).Cmp(_dafny.IntOfInt64(305)) < 0) {
					_coll68.Add((_1874_v37).Plus((_dafny.MultiSetFromSeq(_1862_v27)).Cardinality()), _1865_v30)
				}
			}
			return _coll68.ToMap()
		}())
		_1873_v38 = _nw284
		r3 = _1873_v38
		return r0, r1, r2, r3
	}
}
func (_this *C8) M4(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _1875_v0 _dafny.Int
		_ = _1875_v0
		_1875_v0 = _dafny.IntOfInt64(290)
		var _1876_v1 _dafny.Sequence
		_ = _1876_v1
		_1876_v1 = _dafny.SeqOf(_1875_v0)
		var _1877_v2 *C6
		_ = _1877_v2
		var _nw285 *C6 = New_C6_()
		_ = _nw285
		_nw285.Ctor__(_1875_v0)
		_1877_v2 = _nw285
		var _1878_v3 _dafny.Map
		_ = _1878_v3
		_1878_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1877_v2, _1876_v1)
		var _1879_v4 _dafny.MultiSet
		_ = _1879_v4
		_1879_v4 = _dafny.MultiSetOf(_1876_v1, (func() _dafny.Sequence {
			if (_1878_v3).Contains(_1877_v2) {
				return (_1878_v3).Get(_1877_v2).(_dafny.Sequence)
			}
			return _1876_v1
		})())
		var _1880_v5 bool
		_ = _1880_v5
		_1880_v5 = false
		var _1881_v6 *C3
		_ = _1881_v6
		var _nw286 *C3 = New_C3_()
		_ = _nw286
		_nw286.Ctor__(_1879_v4, false, _1880_v5)
		_1881_v6 = _nw286
		var _1882_v7 _dafny.MultiSet
		_ = _1882_v7
		_1882_v7 = _dafny.MultiSetOf(_1881_v6, _1881_v6)
		var _1883_v8 _dafny.Sequence
		_ = _1883_v8
		_1883_v8 = _dafny.SeqOf(_1880_v5, false)
		var _1884_v9 D6
		_ = _1884_v9
		_1884_v9 = Companion_D6_.Create_DC17_((func() _dafny.Int {
			if (_1882_v7).Contains(_1881_v6) {
				return (_1882_v7).Multiplicity(_1881_v6)
			}
			return _1877_v2.F40
		})(), _1877_v2.F40, _dafny.IntOfInt64(51), (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1883_v8, (Companion_Default___.SafeIndex(_1875_v0, _dafny.IntOfUint32((_1883_v8).Cardinality()))).Uint32(), _1880_v5))).Cardinality(), _1875_v0)
		var _1885_v10 _dafny.CodePoint
		_ = _1885_v10
		_1885_v10 = _dafny.CodePoint('v')
		var _1886_v11 _dafny.MultiSet
		_ = _1886_v11
		_1886_v11 = _dafny.MultiSetOf(_1885_v10)
		var _1887_v12 _dafny.Sequence
		_ = _1887_v12
		_1887_v12 = _dafny.SeqOf(_1886_v11)
		var _pat_let_tv47 = _1881_v6
		_ = _pat_let_tv47
		var _pat_let_tv48 = _1887_v12
		_ = _pat_let_tv48
		var _pat_let_tv49 = _1885_v10
		_ = _pat_let_tv49
		var _pat_let_tv50 = globalState
		_ = _pat_let_tv50
		var _pat_let_tv51 = _1875_v0
		_ = _pat_let_tv51
		var _pat_let_tv52 = _1877_v2
		_ = _pat_let_tv52
		var _pat_let_tv53 = _1877_v2
		_ = _pat_let_tv53
		var _pat_let_tv54 = _1875_v0
		_ = _pat_let_tv54
		var _1888_v13 _dafny.Array
		_ = _1888_v13
		var _nwElement0_71 D6 = _1884_v9
		_ = _nwElement0_71
		var _nw287 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(23))
		_ = _nw287
		(_nw287).ArraySet1(_nwElement0_71, 0)
		(_nw287).ArraySet1(_1884_v9, 1)
		(_nw287).ArraySet1(_1884_v9, 2)
		(_nw287).ArraySet1(_1884_v9, 3)
		(_nw287).ArraySet1(_1884_v9, 4)
		(_nw287).ArraySet1(_1884_v9, 5)
		(_nw287).ArraySet1(Companion_D6_.Create_DC17_(_1875_v0, _1875_v0, (_dafny.Zero).Minus(_1877_v2.F40), _1877_v2.F40, _1877_v2.F40), 6)
		(_nw287).ArraySet1(_1884_v9, 7)
		(_nw287).ArraySet1(_1884_v9, 8)
		(_nw287).ArraySet1(Companion_D6_.Create_DC17_(_dafny.IntOfInt64(195), _1875_v0, _1877_v2.F40, _1875_v0, _1875_v0), 9)
		(_nw287).ArraySet1(_1884_v9, 10)
		(_nw287).ArraySet1(_1884_v9, 11)
		(_nw287).ArraySet1(func(_pat_let46_0 D6) D6 {
			return func(_1889_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let47_0 _dafny.Int) D6 {
					return func(_1890_dt__update_hcf30_h0 _dafny.Int) D6 {
						return func(_pat_let48_0 _dafny.Int) D6 {
							return func(_1891_dt__update_hcf31_h0 _dafny.Int) D6 {
								return Companion_D6_.Create_DC17_(_1890_dt__update_hcf30_h0, _1891_dt__update_hcf31_h0, (_1889_dt__update__tmp_h0).Dtor_cf32(), (_1889_dt__update__tmp_h0).Dtor_cf33(), (_1889_dt__update__tmp_h0).Dtor_cf34())
							}(_pat_let48_0)
						}(_dafny.IntOfInt64(87))
					}(_pat_let47_0)
				}((_pat_let_tv47).Fm13(_pat_let_tv48, _pat_let_tv49, ((_this).F29()).Cardinality(), _pat_let_tv50))
			}(_pat_let46_0)
		}(_1884_v9), 12)
		(_nw287).ArraySet1(_1884_v9, 13)
		(_nw287).ArraySet1(_1884_v9, 14)
		(_nw287).ArraySet1(_1884_v9, 15)
		(_nw287).ArraySet1(func(_pat_let49_0 D6) D6 {
			return func(_1892_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let50_0 _dafny.Int) D6 {
					return func(_1893_dt__update_hcf34_h0 _dafny.Int) D6 {
						return func(_pat_let51_0 _dafny.Int) D6 {
							return func(_1894_dt__update_hcf33_h0 _dafny.Int) D6 {
								return Companion_D6_.Create_DC17_((_1892_dt__update__tmp_h1).Dtor_cf30(), (_1892_dt__update__tmp_h1).Dtor_cf31(), (_1892_dt__update__tmp_h1).Dtor_cf32(), _1894_dt__update_hcf33_h0, _1893_dt__update_hcf34_h0)
							}(_pat_let51_0)
						}((_dafny.Zero).Minus(_pat_let_tv52.F40))
					}(_pat_let50_0)
				}(_pat_let_tv51)
			}(_pat_let49_0)
		}(Companion_D6_.Create_DC17_(_1875_v0, (Companion_Default___.Fm51(_1875_v0, _1875_v0, _1880_v5, _dafny.IntOfInt64(597), globalState)).Cardinality(), _1877_v2.F40, _1875_v0, _1877_v2.F40)), 16)
		(_nw287).ArraySet1(_1884_v9, 17)
		(_nw287).ArraySet1(Companion_D6_.Create_DC17_(_1877_v2.F40, (_1881_v6).Fm13(_1887_v12, _1885_v10, _1877_v2.F40, globalState), _1875_v0, ((_this).F29()).Cardinality(), _1875_v0), 18)
		(_nw287).ArraySet1(_1884_v9, 19)
		(_nw287).ArraySet1(func(_pat_let52_0 D6) D6 {
			return func(_1895_dt__update__tmp_h2 D6) D6 {
				return func(_pat_let53_0 _dafny.Int) D6 {
					return func(_1896_dt__update_hcf33_h1 _dafny.Int) D6 {
						return func(_pat_let54_0 _dafny.Int) D6 {
							return func(_1897_dt__update_hcf30_h1 _dafny.Int) D6 {
								return Companion_D6_.Create_DC17_(_1897_dt__update_hcf30_h1, (_1895_dt__update__tmp_h2).Dtor_cf31(), (_1895_dt__update__tmp_h2).Dtor_cf32(), _1896_dt__update_hcf33_h1, (_1895_dt__update__tmp_h2).Dtor_cf34())
							}(_pat_let54_0)
						}(_pat_let_tv54)
					}(_pat_let53_0)
				}(_pat_let_tv53.F40)
			}(_pat_let52_0)
		}(_1884_v9), 20)
		(_nw287).ArraySet1(_1884_v9, 21)
		(_nw287).ArraySet1(_1884_v9, 22)
		_1888_v13 = _nw287
		var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1888_v13), 0))
		_ = _index265
		(_1888_v13).ArraySet1(Companion_D6_.Create_DC17_(_1875_v0, _1875_v0, _1875_v0, _dafny.IntOfInt64(66), _1875_v0), (_index265).Int())
		var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_1888_v13), 0))
		_ = _index266
		(_1888_v13).ArraySet1(Companion_D6_.Create_DC17_((_dafny.IntOfUint32((_dafny.SeqOf(_1885_v10, _1885_v10)).Cardinality())).Plus(_1877_v2.F40), _1877_v2.F40, _1877_v2.F40, _1875_v0, (_dafny.Zero).Minus(_1877_v2.F40)), (_index266).Int())
		(globalState).F12 = _1875_v0
		var _1898_v14 _dafny.Map
		_ = _1898_v14
		_1898_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1880_v5), _1875_v0)
		var _1899_v15 _dafny.Map
		_ = _1899_v15
		_1899_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1880_v5), _1898_v14)
		_1899_v15 = (_1899_v15).Update(_1880_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v5, _1875_v0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v5, _1875_v0)))
		if !(!(_dafny.Companion_Sequence_.IsPrefixOf(_1883_v8, _dafny.SeqOf(_1880_v5)))) {
			var _1900_v16 _dafny.Sequence
			_ = _1900_v16
			_1900_v16 = _dafny.UnicodeSeqOfUtf8Bytes("kyvdj")
			var _1901_v17 _dafny.MultiSet
			_ = _1901_v17
			_1901_v17 = _dafny.MultiSetOf(_1880_v5)
			var _1902_v18 _dafny.Array
			_ = _1902_v18
			var _nwElement0_72 _dafny.Int = _1877_v2.F40
			_ = _nwElement0_72
			var _nw288 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(19))
			_ = _nw288
			(_nw288).ArraySet1(_nwElement0_72, 0)
			(_nw288).ArraySet1(_1875_v0, 1)
			(_nw288).ArraySet1(_dafny.IntOfInt64(158), 2)
			(_nw288).ArraySet1(_1877_v2.F40, 3)
			(_nw288).ArraySet1((_1877_v2).Fm1(globalState), 4)
			(_nw288).ArraySet1(Companion_Default___.Fm39(globalState), 5)
			(_nw288).ArraySet1(_dafny.IntOfUint32((_1900_v16).Cardinality()), 6)
			(_nw288).ArraySet1(_1877_v2.F40, 7)
			(_nw288).ArraySet1(Companion_Default___.SafeModuloInt(_1877_v2.F40, (_1901_v17).Cardinality()), 8)
			(_nw288).ArraySet1((_1877_v2.F40).Plus(_1877_v2.F40), 9)
			(_nw288).ArraySet1(((_1877_v2).Fm1(globalState)).Times(_1875_v0), 10)
			(_nw288).ArraySet1(_1875_v0, 11)
			(_nw288).ArraySet1(_1877_v2.F40, 12)
			(_nw288).ArraySet1((_1877_v2.F40).Plus((_1898_v14).Cardinality()), 13)
			(_nw288).ArraySet1(Companion_Default___.SafeDivisionInt(_1875_v0, _1875_v0), 14)
			(_nw288).ArraySet1(_dafny.IntOfInt64(-336), 15)
			(_nw288).ArraySet1(((_1901_v17).Union(_dafny.MultiSetOf(_1880_v5))).Cardinality(), 16)
			(_nw288).ArraySet1(_dafny.IntOfInt64(416), 17)
			(_nw288).ArraySet1(Companion_Default___.SafeDivisionInt((_1876_v1).Select((Companion_Default___.SafeIndex(((_this).F29()).Cardinality(), _dafny.IntOfUint32((_1876_v1).Cardinality()))).Uint32()).(_dafny.Int), _1875_v0), 18)
			_1902_v18 = _nw288
			_1902_v18 = _1902_v18
			var _1903_v19 _dafny.Map
			_ = _1903_v19
			_1903_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1881_v6).Fm13(_1887_v12, _1885_v10, _1877_v2.F40, globalState), _1900_v16)
			var _1904_v20 _dafny.Set
			_ = _1904_v20
			_1904_v20 = _dafny.SetOf(_1875_v0, _1877_v2.F40)
			var _1905_v21 _dafny.Map
			_ = _1905_v21
			_1905_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_1903_v19).Contains((_1904_v20).Cardinality()) {
					return (_1903_v19).Get((_1904_v20).Cardinality()).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("xicujy")
			})(), (func() _dafny.Int {
				if _1880_v5 {
					return (_dafny.MultiSetFromSeq(_1900_v16)).Cardinality()
				}
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(_1877_v2.F40))
			})())
			_1905_v21 = (_1905_v21).Update(_1900_v16, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("akwbgonr"), _1900_v16)).Cardinality()))
			var _1906_v22 D3
			_ = _1906_v22
			_1906_v22 = Companion_D3_.Create_DC10_(_1902_v18)
			var _1907_v23 _dafny.Map
			_ = _1907_v23
			_1907_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v5, (_1906_v22).Dtor_cf21())
			var _1908_v24 _dafny.Map
			_ = _1908_v24
			_1908_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_1907_v23).Contains(_1880_v5) {
					return (_1907_v23).Get(_1880_v5).(_dafny.Array)
				}
				return _1902_v18
			})(), !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1900_v16, (Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1900_v16).Cardinality()))).Uint32(), (_1900_v16).Select((Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1900_v16).Cardinality()))).Uint32()).(_dafny.CodePoint)), Companion_Default___.Fm40(_dafny.IntOfInt64(373), _1880_v5, true, globalState)))
			_1908_v24 = (_1908_v24).Update(_1902_v18, (_1877_v2).Fm44(globalState))
			var _1909_v25 D14
			_ = _1909_v25
			_1909_v25 = Companion_D14_.Create_DC32_(_1901_v17, (_1883_v8).Select((Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1883_v8).Cardinality()))).Uint32()).(bool), _1880_v5, _1880_v5, _1877_v2.F40)
			var _1910_v26 D3
			_ = _1910_v26
			_1910_v26 = Companion_D3_.Create_DC11_(_1885_v10)
			var _1911_v27 _dafny.Array
			_ = _1911_v27
			var _nwElement0_73 _dafny.CodePoint = _1885_v10
			_ = _nwElement0_73
			var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(27))
			_ = _nw289
			(_nw289).ArraySet1CodePoint(_nwElement0_73, 0)
			(_nw289).ArraySet1CodePoint(_1885_v10, 1)
			(_nw289).ArraySet1CodePoint(_1885_v10, 2)
			(_nw289).ArraySet1CodePoint(_1885_v10, 3)
			(_nw289).ArraySet1CodePoint(_1885_v10, 4)
			(_nw289).ArraySet1CodePoint(_1885_v10, 5)
			(_nw289).ArraySet1CodePoint(_1885_v10, 6)
			(_nw289).ArraySet1CodePoint(_1885_v10, 7)
			(_nw289).ArraySet1CodePoint((_1910_v26).Dtor_cf22(), 8)
			(_nw289).ArraySet1CodePoint(_1885_v10, 9)
			(_nw289).ArraySet1CodePoint(_1885_v10, 10)
			(_nw289).ArraySet1CodePoint(_1885_v10, 11)
			(_nw289).ArraySet1CodePoint(_1885_v10, 12)
			(_nw289).ArraySet1CodePoint(_1885_v10, 13)
			(_nw289).ArraySet1CodePoint(_1885_v10, 14)
			(_nw289).ArraySet1CodePoint(_1885_v10, 15)
			(_nw289).ArraySet1CodePoint(_dafny.CodePoint('v'), 16)
			(_nw289).ArraySet1CodePoint(_1885_v10, 17)
			(_nw289).ArraySet1CodePoint(_1885_v10, 18)
			(_nw289).ArraySet1CodePoint(_1885_v10, 19)
			(_nw289).ArraySet1CodePoint(_1885_v10, 20)
			(_nw289).ArraySet1CodePoint(_1885_v10, 21)
			(_nw289).ArraySet1CodePoint(_1885_v10, 22)
			(_nw289).ArraySet1CodePoint(_1885_v10, 23)
			(_nw289).ArraySet1CodePoint(_1885_v10, 24)
			(_nw289).ArraySet1CodePoint(_1885_v10, 25)
			(_nw289).ArraySet1CodePoint(_dafny.CodePoint('c'), 26)
			_1911_v27 = _nw289
			var _1912_v28 D1
			_ = _1912_v28
			_1912_v28 = Companion_D1_.Create_DC4_(_1911_v27, _1877_v2.F40, _1877_v2.F40, _1900_v16, _1900_v16)
			var _1913_v29 _dafny.Map
			_ = _1913_v29
			_1913_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1875_v0, (_1904_v20).Cardinality())
			var _1914_v30 _dafny.Sequence
			_ = _1914_v30
			_1914_v30 = _dafny.SeqOf(_dafny.SetOf((func() _dafny.Int {
				if (_1913_v29).Contains(_1877_v2.F40) {
					return (_1913_v29).Get(_1877_v2.F40).(_dafny.Int)
				}
				return (_1913_v29).Cardinality()
			})(), _1877_v2.F40))
			var _1915_v31 _dafny.Array
			_ = _1915_v31
			var _nwElement0_74 bool = _1880_v5
			_ = _nwElement0_74
			var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(22))
			_ = _nw290
			(_nw290).ArraySet1(_nwElement0_74, 0)
			(_nw290).ArraySet1(_1880_v5, 1)
			(_nw290).ArraySet1(_1880_v5, 2)
			(_nw290).ArraySet1(_1880_v5, 3)
			(_nw290).ArraySet1(false, 4)
			(_nw290).ArraySet1(!(_1880_v5), 5)
			(_nw290).ArraySet1((func() bool {
				if (_1909_v25).Dtor_cf58() {
					return _1880_v5
				}
				return false
			})(), 6)
			(_nw290).ArraySet1(false, 7)
			(_nw290).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf((_1912_v28).Dtor_cf10(), _dafny.UnicodeSeqOfUtf8Bytes("kyjjbux")), 8)
			(_nw290).ArraySet1(_1880_v5, 9)
			(_nw290).ArraySet1(_1880_v5, 10)
			(_nw290).ArraySet1(_1880_v5, 11)
			(_nw290).ArraySet1(_1880_v5, 12)
			(_nw290).ArraySet1((_1880_v5) || (_1880_v5), 13)
			(_nw290).ArraySet1(_1880_v5, 14)
			(_nw290).ArraySet1((_1904_v20).IsSubsetOf((_1914_v30).Select((Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1914_v30).Cardinality()))).Uint32()).(_dafny.Set)), 15)
			(_nw290).ArraySet1((!(true)) && (_1880_v5), 16)
			(_nw290).ArraySet1(_1880_v5, 17)
			(_nw290).ArraySet1((_1880_v5) && (_1880_v5), 18)
			(_nw290).ArraySet1(false, 19)
			(_nw290).ArraySet1(false, 20)
			(_nw290).ArraySet1(!(_1880_v5), 21)
			_1915_v31 = _nw290
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_1915_v31), 0))
			_ = _index267
			(_1915_v31).ArraySet1(_1880_v5, (_index267).Int())
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_1915_v31), 0))
			_ = _index268
			(_1915_v31).ArraySet1(!((_1877_v2.F40).Cmp((func() _dafny.Int {
				if _1880_v5 {
					return (_dafny.Zero).Minus(_1875_v0)
				}
				return _1875_v0
			})()) != 0), (_index268).Int())
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1902_v18), 0))
			_ = _index269
			(_1902_v18).ArraySet1(_1877_v2.F40, (_index269).Int())
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_1902_v18), 0))
			_ = _index270
			(_1902_v18).ArraySet1(_1877_v2.F40, (_index270).Int())
		} else {
			r1 = ((func() _dafny.Int {
				if (_1882_v7).Contains(_1881_v6) {
					return (_1882_v7).Multiplicity(_1881_v6)
				}
				return _1875_v0
			})()).Cmp(_1875_v0) < 0
			var _1916_v32 _dafny.Sequence
			_ = _1916_v32
			_1916_v32 = _dafny.UnicodeSeqOfUtf8Bytes("vaiuecm")
			(globalState).F6 = _1916_v32
			var _1917_v33 _dafny.Array
			_ = _1917_v33
			var _nw291 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw291
			_1917_v33 = _nw291
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1917_v33), 0))
			_ = _index271
			(_1917_v33).ArraySet1(_dafny.IntOfInt64(-728), (_index271).Int())
			var _1918_v34 _dafny.Map
			_ = _1918_v34
			_1918_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(480), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1880_v5, _1875_v0))
			var _1919_v35 T0
			_ = _1919_v35
			var _nw292 *C7 = New_C7_()
			_ = _nw292
			_nw292.Ctor__(_1880_v5, _1918_v34)
			_1919_v35 = _nw292
			var _1920_v36 _dafny.Set
			_ = _1920_v36
			_1920_v36 = _dafny.SetOf(_1919_v35)
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1917_v33), 0))
			_ = _index272
			(_1917_v33).ArraySet1(((func() _dafny.Set {
				if !((_1880_v5) && (_1880_v5)) {
					return (_1920_v36).Union(_1920_v36)
				}
				return (_1920_v36).Difference(_1920_v36)
			})()).Cardinality(), (_index272).Int())
			var _1921_v37 D3
			_ = _1921_v37
			_1921_v37 = Companion_D3_.Create_DC10_(_1917_v33)
			var _1922_v38 _dafny.Map
			_ = _1922_v38
			_1922_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v37, _1877_v2.F40)
			var _1923_v39 _dafny.MultiSet
			_ = _1923_v39
			_1923_v39 = _dafny.MultiSetOf(_1880_v5, _1880_v5)
			var _1924_v40 D14
			_ = _1924_v40
			_1924_v40 = Companion_D14_.Create_DC32_(_1923_v39, _1880_v5, !(_1880_v5), _1880_v5, _1877_v2.F40)
			var _1925_v41 *C1
			_ = _1925_v41
			var _nw293 *C1 = New_C1_()
			_ = _nw293
			_nw293.Ctor__(((_this).F29()).Cardinality(), (Companion_Default___.Fm52(_1880_v5, (_1917_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1917_v33), 0))).Int()).(_dafny.Int), globalState)).Cardinality(), (_1924_v40).Dtor_cf57(), _1880_v5)
			_1925_v41 = _nw293
			var _1926_v42 _dafny.Sequence
			_ = _1926_v42
			_1926_v42 = _dafny.SeqOf(_1925_v41)
			var _1927_v43 *C4
			_ = _1927_v43
			var _nw294 *C4 = New_C4_()
			_ = _nw294
			_nw294.Ctor__(_1922_v38, _1880_v5, !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(_1926_v42, (Companion_Default___.SafeIndex((_1925_v41).F36(), _dafny.IntOfUint32((_1926_v42).Cardinality()))).Uint32(), _1925_v41), _dafny.SeqOf(_1925_v41, _1925_v41)))
			_1927_v43 = _nw294
			var _1928_v44 _dafny.Sequence
			_ = _1928_v44
			_1928_v44 = _dafny.SeqOf(_1916_v32)
			var _1929_v45 _dafny.Array
			_ = _1929_v45
			var _nwElement0_75 _dafny.Sequence = _1916_v32
			_ = _nwElement0_75
			var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(25))
			_ = _nw295
			(_nw295).ArraySet1(_nwElement0_75, 0)
			(_nw295).ArraySet1(_1916_v32, 1)
			(_nw295).ArraySet1(_1916_v32, 2)
			(_nw295).ArraySet1((_1928_v44).Select((Companion_Default___.SafeIndex((_1917_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1917_v33), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1928_v44).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
			(_nw295).ArraySet1(_1916_v32, 4)
			(_nw295).ArraySet1(_1916_v32, 5)
			(_nw295).ArraySet1(_1916_v32, 6)
			(_nw295).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hiqyfxr"), 7)
			(_nw295).ArraySet1(_1916_v32, 8)
			(_nw295).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1916_v32, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-400), _dafny.IntOfUint32((_1916_v32).Cardinality()))).Uint32(), _1885_v10), _dafny.UnicodeSeqOfUtf8Bytes("mfc")), 9)
			(_nw295).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("bsbqonwkb"), 10)
			(_nw295).ArraySet1(_1916_v32, 11)
			(_nw295).ArraySet1(_1916_v32, 12)
			(_nw295).ArraySet1(_1916_v32, 13)
			(_nw295).ArraySet1(_1916_v32, 14)
			(_nw295).ArraySet1(_dafny.Companion_Sequence_.Update(_1916_v32, (Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1916_v32).Cardinality()))).Uint32(), _1885_v10), 15)
			(_nw295).ArraySet1(_1916_v32, 16)
			(_nw295).ArraySet1(_1916_v32, 17)
			(_nw295).ArraySet1(_1916_v32, 18)
			(_nw295).ArraySet1(_1916_v32, 19)
			(_nw295).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("kasxjomfh"), 20)
			(_nw295).ArraySet1(_1916_v32, 21)
			(_nw295).ArraySet1((_1928_v44).Select((Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1928_v44).Cardinality()))).Uint32()).(_dafny.Sequence), 22)
			(_nw295).ArraySet1(_1916_v32, 23)
			(_nw295).ArraySet1(_1916_v32, 24)
			_1929_v45 = _nw295
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1929_v45), 0))
			_ = _index273
			(_1929_v45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1916_v32, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(283))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}(func(_1930_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			}))), (_index273).Int())
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1929_v45), 0))
			_ = _index274
			(_1929_v45).ArraySet1(_1916_v32, (_index274).Int())
		}
		var _1931_v46 _dafny.Sequence
		_ = _1931_v46
		_1931_v46 = _dafny.UnicodeSeqOfUtf8Bytes("tgaaxgqsq")
		var _1932_v47 D6
		_ = _1932_v47
		_1932_v47 = Companion_D6_.Create_DC16_(_dafny.Companion_Sequence_.Concatenate(_1931_v46, _1931_v46))
		var _pat_let_tv55 = _1931_v46
		_ = _pat_let_tv55
		_1932_v47 = (func() D6 {
			if _1880_v5 {
				return _1932_v47
			}
			return func(_pat_let55_0 D6) D6 {
				return func(_1933_dt__update__tmp_h3 D6) D6 {
					return func(_pat_let56_0 _dafny.Sequence) D6 {
						return func(_1934_dt__update_hcf29_h0 _dafny.Sequence) D6 {
							return Companion_D6_.Create_DC16_(_1934_dt__update_hcf29_h0)
						}(_pat_let56_0)
					}(_pat_let_tv55)
				}(_pat_let55_0)
			}(_1932_v47)
		})()
		var _hi17 _dafny.Int = (_1877_v2.F40).Minus(_1875_v0)
		_ = _hi17
		for _1935_i1 := (_dafny.IntOfInt64(-597)).Times(_1875_v0); _1935_i1.Cmp(_hi17) < 0; _1935_i1 = _1935_i1.Plus(_dafny.One) {
			_1885_v10 = _1885_v10
			var _1936_v48 *C0
			_ = _1936_v48
			var _nw296 *C0 = New_C0_()
			_ = _nw296
			_nw296.Ctor__(true)
			_1936_v48 = _nw296
			(globalState).F22 = !((func() bool {
				if _1880_v5 {
					return _1880_v5
				}
				return (_1936_v48).F32()
			})())
			var _1937_v49 D0
			_ = _1937_v49
			_1937_v49 = Companion_D0_.Create_DC1_(true, (_1936_v48).F32(), _1877_v2.F40, (_1936_v48).F32())
			var _1938_v50 _dafny.Set
			_ = _1938_v50
			_1938_v50 = _dafny.SetOf(_1937_v49)
			_1938_v50 = (_1938_v50).Difference((_1938_v50).Union(_dafny.SetOf(_1937_v49)))
		}
		r0 = _1880_v5
		r1 = (_1883_v8).Select((Companion_Default___.SafeIndex(_1877_v2.F40, _dafny.IntOfUint32((_1883_v8).Cardinality()))).Uint32()).(bool)
		return r0, r1
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	_f26 bool
	_f27 bool
	_f29 _dafny.MultiSet
	F30  bool
	_f31 _dafny.Int
}

func New_C9_() *C9 {
	_this := C9{}

	_this._f26 = false
	_this._f27 = false
	_this._f29 = _dafny.EmptyMultiSet
	_this.F30 = false
	_this._f31 = _dafny.Zero
	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C9{}
var _ T3 = &C9{}
var _ T1 = &C9{}
var _ T0 = &C9{}
var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) F26() bool {
	return _this._f26
}
func (_this *C9) F27() bool {
	return _this._f27
}
func (_this *C9) F29() _dafny.MultiSet {
	return _this._f29
}
func (_this *C9) Ctor__(f30 bool, f31 _dafny.Int, f26 bool, f27 bool, f29 _dafny.MultiSet) {
	{
		(_this).F30 = f30
		(_this)._f31 = f31
		(_this)._f26 = f26
		(_this)._f27 = f27
		(_this)._f29 = f29
	}
}
func (_this *C9) Fm3(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26(), (_this).F26()), _dafny.SeqOf((_this).F27())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F26()), _dafny.SeqOf((_this).F27())))
	}
}
func (_this *C9) Fm4(p0 D0, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_this).F31()).Minus((_this).F31())
	}
}
func (_this *C9) Fm1(globalState *GlobalState) _dafny.Int {
	{
		return (_this).F31()
	}
}
func (_this *C9) Fm2(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_this).F27()) == ((_this).F27())
	}
}
func (_this *C9) M2(p0 bool, p1 T0, p2 bool, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var _1939_v0 _dafny.Array
		_ = _1939_v0
		var _nw297 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
		_ = _nw297
		_1939_v0 = _nw297
		var _1940_v1 _dafny.Map
		_ = _1940_v1
		_1940_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _this.F30)
		var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))
		_ = _index275
		(_1939_v0).ArraySet1(_1940_v1, (_index275).Int())
		var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))
		_ = _index276
		var _rhs285 _dafny.Int = (func() _dafny.Int {
			if _this.F30 {
				return (_dafny.Zero).Minus((_this).F31())
			}
			return Companion_Default___.SafeDivisionInt((_this).F31(), (_this).F31())
		})()
		_ = _rhs285
		var _rhs286 _dafny.Map = (_1940_v1).Merge((_1940_v1).Merge(_1940_v1))
		_ = _rhs286
		var _lhs257 *GlobalState = globalState
		_ = _lhs257
		var _lhs258 _dafny.Array = _1939_v0
		_ = _lhs258
		var _lhs259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))
		_ = _lhs259
		_lhs257.F12 = _rhs285
		(_lhs258).ArraySet1(_rhs286, (_lhs259).Int())
		var _1941_v2 _dafny.Sequence
		_ = _1941_v2
		_1941_v2 = _dafny.SeqOf((func() bool {
			if ((_1939_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))).Int()).(_dafny.Map)).Contains((_this).F27()) {
				return ((_1939_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))).Int()).(_dafny.Map)).Get((_this).F27()).(bool)
			}
			return p0
		})(), !(_this.F30), (_this).Fm2(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F31())).Cardinality(), globalState))
		var _1942_v3 _dafny.CodePoint
		_ = _1942_v3
		_1942_v3 = _dafny.CodePoint('v')
		var _1943_v4 _dafny.Map
		_ = _1943_v4
		_1943_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1941_v2).Cardinality()), _1942_v3)
		var _1944_v5 _dafny.Map
		_ = _1944_v5
		_1944_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _this.F30)
		var _1945_v6 _dafny.Set
		_ = _1945_v6
		_1945_v6 = _dafny.SetOf((func() _dafny.CodePoint {
			if (_1943_v4).Contains((_1944_v5).Cardinality()) {
				return (_1943_v4).Get((_1944_v5).Cardinality()).(_dafny.CodePoint)
			}
			return _1942_v3
		})(), _1942_v3)
		var _1946_i0 _dafny.Int
		_ = _1946_i0
		_1946_i0 = _dafny.Zero
		{
			for (_1945_v6).IsDisjointFrom(_1945_v6) {
				{
					if (_1946_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1946_i0 = (_1946_i0).Plus(_dafny.One)
					var _1947_v7 D0
					_ = _1947_v7
					_1947_v7 = Companion_D0_.Create_DC0_(p2)
					var _1948_v8 _dafny.Map
					_ = _1948_v8
					_1948_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F31())
					var _pat_let_tv56 = globalState
					_ = _pat_let_tv56
					var _pat_let_tv57 = _1948_v8
					_ = _pat_let_tv57
					var _source27 D0 = func(_pat_let57_0 D0) D0 {
						return func(_1949_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let58_0 bool) D0 {
								return func(_1950_dt__update_hcf0_h0 bool) D0 {
									return Companion_D0_.Create_DC0_(_1950_dt__update_hcf0_h0)
								}(_pat_let58_0)
							}(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm1(_pat_let_tv56), (_this).F31())).Equals(_pat_let_tv57))
						}(_pat_let57_0)
					}(_1947_v7)
					_ = _source27
					if _source27.Is_DC1() {
						var _1951___mcc_h0 bool = _source27.Get_().(D0_DC1).Cf1
						_ = _1951___mcc_h0
						var _1952___mcc_h1 bool = _source27.Get_().(D0_DC1).Cf2
						_ = _1952___mcc_h1
						var _1953___mcc_h2 _dafny.Int = _source27.Get_().(D0_DC1).Cf3
						_ = _1953___mcc_h2
						var _1954___mcc_h3 bool = _source27.Get_().(D0_DC1).Cf4
						_ = _1954___mcc_h3
						var _1955_cf4 bool = _1954___mcc_h3
						_ = _1955_cf4
						var _1956_cf3 _dafny.Int = _1953___mcc_h2
						_ = _1956_cf3
						var _1957_cf2 bool = _1952___mcc_h1
						_ = _1957_cf2
						var _1958_cf1 bool = _1951___mcc_h0
						_ = _1958_cf1
						(globalState).F22 = !((_1947_v7).Dtor_cf0())
						_1943_v4 = _1943_v4
						var _1959_v9 _dafny.Array
						_ = _1959_v9
						var _len0_48 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_48
						var _nw298 _dafny.Array
						_ = _nw298
						if _len0_48.Cmp(_dafny.Zero) == 0 {
							_nw298 = _dafny.NewArray(_len0_48)
						} else {
							var _init48 func(_dafny.Int) bool = (func(_1960_p2 bool) func(_dafny.Int) bool {
								return func(_1961_i1 _dafny.Int) bool {
									return _1960_p2
								}
							})(p2)
							_ = _init48
							var _element0_48 = _init48(_dafny.Zero)
							_ = _element0_48
							_nw298 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
							(_nw298).ArraySet1(_element0_48, 0)
							var _nativeLen0_48 = (_len0_48).Int()
							_ = _nativeLen0_48
							for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
								(_nw298).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
							}
						}
						_1959_v9 = _nw298
						var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1959_v9), 0))
						_ = _index277
						(_1959_v9).ArraySet1((func() bool {
							if true {
								return (_this).F26()
							}
							return _this.F30
						})(), (_index277).Int())
						var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(924), _dafny.ArrayLen((_1959_v9), 0))
						_ = _index278
						(_1959_v9).ArraySet1(_1955_cf4, (_index278).Int())
						(globalState).F12 = (func() _dafny.Int {
							if ((_this).F29()).Contains(_1956_cf3) {
								return ((_this).F29()).Multiplicity(_1956_cf3)
							}
							return (_this).F31()
						})()
					} else if _source27.Is_DC0() {
						var _1962___mcc_h4 bool = _source27.Get_().(D0_DC0).Cf0
						_ = _1962___mcc_h4
						var _1963_cf0 bool = _1962___mcc_h4
						_ = _1963_cf0
						var _1964_v10 _dafny.Array
						_ = _1964_v10
						var _len0_49 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_49
						var _nw299 _dafny.Array
						_ = _nw299
						if _len0_49.Cmp(_dafny.Zero) == 0 {
							_nw299 = _dafny.NewArray(_len0_49)
						} else {
							var _init49 func(_dafny.Int) bool = (func(_1965_p0 bool) func(_dafny.Int) bool {
								return func(_1966_i2 _dafny.Int) bool {
									return _1965_p0
								}
							})(p0)
							_ = _init49
							var _element0_49 = _init49(_dafny.Zero)
							_ = _element0_49
							_nw299 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
							(_nw299).ArraySet1(_element0_49, 0)
							var _nativeLen0_49 = (_len0_49).Int()
							_ = _nativeLen0_49
							for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
								(_nw299).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
							}
						}
						_1964_v10 = _nw299
						var _1967_v11 _dafny.Sequence
						_ = _1967_v11
						_1967_v11 = _dafny.SeqOf(_1964_v10, _1964_v10, _1964_v10, _1964_v10, _1964_v10)
						(globalState).F12 = ((func() _dafny.Int {
							if p0 {
								return _dafny.IntOfInt64(692)
							}
							return (_this).F31()
						})()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1967_v11, _1967_v11)).Cardinality()))
						var _1968_v12 _dafny.Sequence
						_ = _1968_v12
						_1968_v12 = _dafny.UnicodeSeqOfUtf8Bytes("bex")
						var _1969_v13 _dafny.Array
						_ = _1969_v13
						var _nwElement0_76 _dafny.Int = (_this).F31()
						_ = _nwElement0_76
						var _nw300 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(24))
						_ = _nw300
						(_nw300).ArraySet1(_nwElement0_76, 0)
						(_nw300).ArraySet1(_dafny.IntOfInt64(-773), 1)
						(_nw300).ArraySet1(_dafny.IntOfUint32((_1968_v12).Cardinality()), 2)
						(_nw300).ArraySet1((_this).F31(), 3)
						(_nw300).ArraySet1((_this).F31(), 4)
						(_nw300).ArraySet1((_this).F31(), 5)
						(_nw300).ArraySet1((_this).F31(), 6)
						(_nw300).ArraySet1((_this).F31(), 7)
						(_nw300).ArraySet1((_this).F31(), 8)
						(_nw300).ArraySet1((_this).F31(), 9)
						(_nw300).ArraySet1((_this).F31(), 10)
						(_nw300).ArraySet1((_this).F31(), 11)
						(_nw300).ArraySet1((_this).F31(), 12)
						(_nw300).ArraySet1((_this).F31(), 13)
						(_nw300).ArraySet1((_this).F31(), 14)
						(_nw300).ArraySet1((_this).F31(), 15)
						(_nw300).ArraySet1((_this).F31(), 16)
						(_nw300).ArraySet1(_dafny.IntOfUint32((_1941_v2).Cardinality()), 17)
						(_nw300).ArraySet1((_this).F31(), 18)
						(_nw300).ArraySet1((_this).F31(), 19)
						(_nw300).ArraySet1(_dafny.IntOfUint32((_1968_v12).Cardinality()), 20)
						(_nw300).ArraySet1((_this).F31(), 21)
						(_nw300).ArraySet1((_this).F31(), 22)
						(_nw300).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1941_v2, _1941_v2)).Cardinality()), 23)
						_1969_v13 = _nw300
						var _1970_v14 D3
						_ = _1970_v14
						_1970_v14 = Companion_D3_.Create_DC10_(_1969_v13)
						var _1971_v15 _dafny.Map
						_ = _1971_v15
						_1971_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1970_v14, (_this).F31())
						var _1972_v16 *C4
						_ = _1972_v16
						var _nw301 *C4 = New_C4_()
						_ = _nw301
						_nw301.Ctor__(_1971_v15, _1963_cf0, !((_this).F26()))
						_1972_v16 = _nw301
						(globalState).F2 = true
						var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1969_v13), 0))
						_ = _index279
						(_1969_v13).ArraySet1((_this).F31(), (_index279).Int())
						var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1969_v13), 0))
						_ = _index280
						(_1969_v13).ArraySet1((_this).F31(), (_index280).Int())
					} else {
						var _1973___mcc_h5 D0 = _source27.Get_().(D0_DC2).Cf5
						_ = _1973___mcc_h5
						var _1974_cf5 D0 = _1973___mcc_h5
						_ = _1974_cf5
						var _1975_v17 _dafny.Array
						_ = _1975_v17
						var _nw302 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
						_ = _nw302
						_1975_v17 = _nw302
						var _1976_v18 _dafny.Sequence
						_ = _1976_v18
						_1976_v18 = _dafny.UnicodeSeqOfUtf8Bytes("dbooou")
						var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1975_v17), 0))
						_ = _index281
						(_1975_v17).ArraySet1(_1976_v18, (_index281).Int())
						var _1977_v19 _dafny.Set
						_ = _1977_v19
						_1977_v19 = _dafny.SetOf(_this.F30)
						var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1975_v17), 0))
						_ = _index282
						(_1975_v17).ArraySet1((func() _dafny.Sequence {
							if (_1977_v19).IsProperSubsetOf(_1977_v19) {
								return _1976_v18
							}
							return _dafny.Companion_Sequence_.Concatenate(_1976_v18, _1976_v18)
						})(), (_index282).Int())
						var _1978_v20 _dafny.Set
						_ = _1978_v20
						_1978_v20 = _dafny.SetOf((_this).F31(), (_this).F31())
						var _1979_v21 _dafny.Array
						_ = _1979_v21
						var _nw303 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
						_ = _nw303
						_1979_v21 = _nw303
						var _1980_v22 D1
						_ = _1980_v22
						_1980_v22 = Companion_D1_.Create_DC4_(_1979_v21, _dafny.IntOfInt64(731), _dafny.IntOfInt64(279), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(60))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg108 _dafny.Int) interface{} {
								return coer108(arg108)
							}
						}((func(_1981_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1982_i3 _dafny.Int) _dafny.CodePoint {
								return _1981_v3
							}
						})(_1942_v3))), _dafny.UnicodeSeqOfUtf8Bytes("l"))
						(globalState).F12 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1976_v18).Cardinality()), (_1978_v20).Cardinality())), (_1980_v22).Dtor_cf8())
						var _1983_v23 _dafny.Map
						_ = _1983_v23
						_1983_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F31())
						(globalState).F2 = (func() bool {
							if (_this).F27() {
								return false
							}
							return (p1).Fm2((_this).Fm2(_this.F30, (_1983_v23).Cardinality(), globalState), (_this).F31(), globalState)
						})()
						(globalState).F2 = (_this).Fm2(((_this).F31()).Cmp((_this).F31()) >= 0, ((_this).F31()).Plus(_dafny.IntOfInt64(111)), globalState)
					}
					var _1984_v24 _dafny.Sequence
					_ = _1984_v24
					_1984_v24 = _dafny.SeqOf(_1942_v3, _1942_v3)
					var _1985_v25 _dafny.Sequence
					_ = _1985_v25
					_1985_v25 = _dafny.SeqOf(_1984_v24, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg109 _dafny.Int) interface{} {
							return coer109(arg109)
						}
					}((func(_1986_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1987_i4 _dafny.Int) _dafny.CodePoint {
							return _1986_v3
						}
					})(_1942_v3))))
					_1985_v25 = _1985_v25
					if _this.F30 {
						var _1988_v26 _dafny.Array
						_ = _1988_v26
						var _nw304 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
						_ = _nw304
						_1988_v26 = _nw304
						var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1988_v26), 0))
						_ = _index283
						(_1988_v26).ArraySet1(!((p1).Fm2((_this).F26(), (_this).F31(), globalState)), (_index283).Int())
						var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1988_v26), 0))
						_ = _index284
						(_1988_v26).ArraySet1((_this).Fm2((_this).F26(), (_this).F31(), globalState), (_index284).Int())
						var _1989_v27 _dafny.Sequence
						_ = _1989_v27
						_1989_v27 = _dafny.SeqOf(_1988_v26)
						var _1990_v28 _dafny.Array
						_ = _1990_v28
						var _len0_50 _dafny.Int = _dafny.IntOfInt64(17)
						_ = _len0_50
						var _nw305 _dafny.Array
						_ = _nw305
						if _len0_50.Cmp(_dafny.Zero) == 0 {
							_nw305 = _dafny.NewArray(_len0_50)
						} else {
							var _init50 func(_dafny.Int) _dafny.Int = func(_1991_i5 _dafny.Int) _dafny.Int {
								return (_1991_i5).Plus((_this).F31())
							}
							_ = _init50
							var _element0_50 = _init50(_dafny.Zero)
							_ = _element0_50
							_nw305 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
							(_nw305).ArraySet1(_element0_50, 0)
							var _nativeLen0_50 = (_len0_50).Int()
							_ = _nativeLen0_50
							for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
								(_nw305).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
							}
						}
						_1990_v28 = _nw305
						var _rhs287 _dafny.Sequence = _1989_v27
						_ = _rhs287
						var _rhs288 bool = (p1).Fm2((_1988_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(357), _dafny.ArrayLen((_1988_v26), 0))).Int()).(bool), (_this).F31(), globalState)
						_ = _rhs288
						var _rhs289 _dafny.Array = _1990_v28
						_ = _rhs289
						var _rhs290 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F31(), (_this).F31())
						_ = _rhs290
						var _lhs260 *GlobalState = globalState
						_ = _lhs260
						var _lhs261 *GlobalState = globalState
						_ = _lhs261
						_1989_v27 = _rhs287
						_lhs260.F24 = _rhs288
						_1990_v28 = _rhs289
						_lhs261.F16 = _rhs290
						var _1992_v29 _dafny.Sequence
						_ = _1992_v29
						_1992_v29 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F31())).Cardinality()))
						var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1990_v28), 0))
						_ = _index285
						(_1990_v28).ArraySet1((func() _dafny.Int {
							if p0 {
								return (_this).F31()
							}
							return _dafny.IntOfUint32((_1992_v29).Cardinality())
						})(), (_index285).Int())
						var _1993_v30 _dafny.Sequence
						_ = _1993_v30
						_1993_v30 = _dafny.SeqOf(_1989_v27, _1989_v27)
						var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1990_v28), 0))
						_ = _index286
						var _rhs291 _dafny.Int = (_dafny.Zero).Minus((_this).F31())
						_ = _rhs291
						var _rhs292 bool = (p1).Fm2(_dafny.Companion_Sequence_.Contains((_1993_v30).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.IntOfUint32((_1993_v30).Cardinality()))).Uint32()).(_dafny.Sequence), _1988_v26), _dafny.IntOfInt64(22), globalState)
						_ = _rhs292
						var _rhs293 _dafny.Int = (_this).F31()
						_ = _rhs293
						var _rhs294 bool = _this.F30
						_ = _rhs294
						var _lhs262 *GlobalState = globalState
						_ = _lhs262
						var _lhs263 *C9 = _this
						_ = _lhs263
						var _lhs264 _dafny.Array = _1990_v28
						_ = _lhs264
						var _lhs265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1990_v28), 0))
						_ = _lhs265
						var _lhs266 *GlobalState = globalState
						_ = _lhs266
						_lhs262.F12 = _rhs291
						_lhs263.F30 = _rhs292
						(_lhs264).ArraySet1(_rhs293, (_lhs265).Int())
						_lhs266.F15 = _rhs294
						_1988_v26 = _1988_v26
						var _1994_v31 _dafny.Map
						_ = _1994_v31
						_1994_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC10_(_1990_v28), (_1990_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1990_v28), 0))).Int()).(_dafny.Int))
						var _1995_v32 *C4
						_ = _1995_v32
						var _nw306 *C4 = New_C4_()
						_ = _nw306
						_nw306.Ctor__(_1994_v31, p2, !(p0))
						_1995_v32 = _nw306
					} else {
						var _1996_v33 _dafny.Map
						_ = _1996_v33
						_1996_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F31())
						var _1997_v34 _dafny.Int
						_ = _1997_v34
						var _1998_v35 _dafny.Int
						_ = _1998_v35
						var _1999_v36 bool
						_ = _1999_v36
						var _out61 _dafny.Int
						_ = _out61
						var _out62 _dafny.Int
						_ = _out62
						var _out63 bool
						_ = _out63
						_out61, _out62, _out63 = (_this).M6(_1996_v33, (_dafny.Zero).Minus((_this).F31()), _1942_v3, (_this).F31(), globalState)
						_1997_v34 = _out61
						_1998_v35 = _out62
						_1999_v36 = _out63
						_1942_v3 = _dafny.CodePoint('o')
						var _2000_v37 _dafny.Array
						_ = _2000_v37
						var _len0_51 _dafny.Int = _dafny.IntOfInt64(3)
						_ = _len0_51
						var _nw307 _dafny.Array
						_ = _nw307
						if _len0_51.Cmp(_dafny.Zero) == 0 {
							_nw307 = _dafny.NewArray(_len0_51)
						} else {
							var _init51 func(_dafny.Int) bool = func(_2001_i6 _dafny.Int) bool {
								return (_this).F26()
							}
							_ = _init51
							var _element0_51 = _init51(_dafny.Zero)
							_ = _element0_51
							_nw307 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
							(_nw307).ArraySet1(_element0_51, 0)
							var _nativeLen0_51 = (_len0_51).Int()
							_ = _nativeLen0_51
							for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
								(_nw307).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
							}
						}
						_2000_v37 = _nw307
						_2000_v37 = _2000_v37
						var _2002_v38 _dafny.Map
						_ = _2002_v38
						_2002_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
						var _2003_v39 D12
						_ = _2003_v39
						_2003_v39 = Companion_D12_.Create_DC27_(p1)
						(globalState).F24 = (func() bool {
							if (_2002_v38).Contains((_2003_v39).Dtor_cf48()) {
								return (_2002_v38).Get((_2003_v39).Dtor_cf48()).(bool)
							}
							return p0
						})()
						(globalState).F24 = (_this).F26()
					}
					var _2004_v40 _dafny.Set
					_ = _2004_v40
					_2004_v40 = _dafny.SetOf(_1940_v1, ((_1939_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1939_v0), 0))).Int()).(_dafny.Map)).Update(p0, (_this).F26()), _1940_v1)
					_2004_v40 = (_2004_v40).Union(_2004_v40)
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _2005_v41 D12
		_ = _2005_v41
		_2005_v41 = Companion_D12_.Create_DC28_(_dafny.IntOfInt64(229))
		var _source28 D12 = _2005_v41
		_ = _source28
		if _source28.Is_DC28() {
			var _2006___mcc_h6 _dafny.Int = _source28.Get_().(D12_DC28).Cf49
			_ = _2006___mcc_h6
			var _2007_cf49 _dafny.Int = _2006___mcc_h6
			_ = _2007_cf49
			(globalState).F22 = (func() bool {
				if _this.F30 {
					return (func() bool {
						if _this.F30 {
							return _this.F30
						}
						return p2
					})()
				}
				return _this.F30
			})()
			(globalState).F2 = (_1941_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(237), _dafny.IntOfUint32((_1941_v2).Cardinality()))).Uint32()).(bool)
			var _2008_v42 *C5
			_ = _2008_v42
			var _nw308 *C5 = New_C5_()
			_ = _nw308
			_nw308.Ctor__(false, false)
			_2008_v42 = _nw308
			(globalState).F12 = (_this).F31()
		} else {
			var _2009___mcc_h7 T0 = _source28.Get_().(D12_DC27).Cf48
			_ = _2009___mcc_h7
			var _2010_cf48 T0 = _2009___mcc_h7
			_ = _2010_cf48
			var _2011_v43 _dafny.Array
			_ = _2011_v43
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_52
			var _nw309 _dafny.Array
			_ = _nw309
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw309 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Int = func(_2012_i7 _dafny.Int) _dafny.Int {
					return (_2012_i7).Plus((_this).F31())
				}
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw309 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw309).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw309).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_2011_v43 = _nw309
			var _2013_v44 D3
			_ = _2013_v44
			_2013_v44 = Companion_D3_.Create_DC10_(_2011_v43)
			var _2014_v45 _dafny.Map
			_ = _2014_v45
			_2014_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2013_v44, _dafny.IntOfUint32((_dafny.SeqOf(_this.F30)).Cardinality()))
			var _2015_v46 *C4
			_ = _2015_v46
			var _nw310 *C4 = New_C4_()
			_ = _nw310
			_nw310.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2013_v44, (_this).F31())).Merge(_2014_v45), p2, (func() bool {
				if p2 {
					return p0
				}
				return (_this).F27()
			})())
			_2015_v46 = _nw310
			if (_this).F26() {
				var _2016_v47 _dafny.Set
				_ = _2016_v47
				_2016_v47 = _dafny.SetOf((_2015_v46).Fm1(globalState), (_this).F31())
				_2016_v47 = _2016_v47
				var _2017_v48 _dafny.Sequence
				_ = _2017_v48
				_2017_v48 = _dafny.UnicodeSeqOfUtf8Bytes("vsjg")
				(globalState).F6 = _2017_v48
				var _2018_v49 _dafny.Set
				_ = _2018_v49
				_2018_v49 = _dafny.SetOf(_2017_v48, _2017_v48, _2017_v48)
				var _2019_v50 _dafny.MultiSet
				_ = _2019_v50
				_2019_v50 = _dafny.MultiSetOf(_1942_v3)
				var _2020_v51 D0
				_ = _2020_v51
				_2020_v51 = Companion_D0_.Create_DC0_(p2)
				var _2021_v52 _dafny.Set
				_ = _2021_v52
				_2021_v52 = _dafny.SetOf((_1941_v2).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_1941_v2).Cardinality()))).Uint32()).(bool))
				var _2022_v53 T2
				_ = _2022_v53
				var _nw311 *C1 = New_C1_()
				_ = _nw311
				_nw311.Ctor__(_dafny.IntOfUint32((_2017_v48).Cardinality()), _dafny.IntOfInt64(800), (_this).F27(), p0)
				_2022_v53 = _nw311
				var _2023_v54 _dafny.Sequence
				_ = _2023_v54
				_2023_v54 = _dafny.SeqOf(_2022_v53, _2022_v53)
				var _2024_v55 _dafny.Map
				_ = _2024_v55
				_2024_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2023_v54, (_this).F31())
				var _2025_v56 _dafny.Array
				_ = _2025_v56
				var _nwElement0_77 bool = (_2018_v49).IsProperSubsetOf(_2018_v49)
				_ = _nwElement0_77
				var _nw312 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(25))
				_ = _nw312
				(_nw312).ArraySet1(_nwElement0_77, 0)
				(_nw312).ArraySet1((_this).F27(), 1)
				(_nw312).ArraySet1((_this).F26(), 2)
				(_nw312).ArraySet1((_2019_v50).IsSubsetOf(_2019_v50), 3)
				(_nw312).ArraySet1((_2020_v51).Dtor_cf0(), 4)
				(_nw312).ArraySet1(((_this).F27()) || (_this.F30), 5)
				(_nw312).ArraySet1(((_this).F31()).Cmp((_this).F31()) <= 0, 6)
				(_nw312).ArraySet1(p0, 7)
				(_nw312).ArraySet1(!((_this).F26()), 8)
				(_nw312).ArraySet1(!((_dafny.SetOf(_this.F30, p2)).IsProperSubsetOf(_2021_v52)), 9)
				(_nw312).ArraySet1((_2010_cf48).Fm2((_this).F27(), (_this).F31(), globalState), 10)
				(_nw312).ArraySet1(!(!((_this.F30) || ((_this).F27()))), 11)
				(_nw312).ArraySet1(p0, 12)
				(_nw312).ArraySet1(p0, 13)
				(_nw312).ArraySet1(_this.F30, 14)
				(_nw312).ArraySet1((_this).Fm2((_this).F27(), (_2024_v55).Cardinality(), globalState), 15)
				(_nw312).ArraySet1((_2022_v53).F26(), 16)
				(_nw312).ArraySet1(!(((_this).F27()) == ((_this).F27())), 17)
				(_nw312).ArraySet1(!((_dafny.SetOf(_dafny.IntOfInt64(379))).IsDisjointFrom(_2016_v47)), 18)
				(_nw312).ArraySet1(p2, 19)
				(_nw312).ArraySet1(!((_this).Fm2((_2022_v53).F26(), (_this).F31(), globalState)) || ((_this).F27()), 20)
				(_nw312).ArraySet1((_this).F27(), 21)
				(_nw312).ArraySet1(_this.F30, 22)
				(_nw312).ArraySet1((_2022_v53).F26(), 23)
				(_nw312).ArraySet1((_1941_v2).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_1941_v2).Cardinality()))).Uint32()).(bool), 24)
				_2025_v56 = _nw312
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_2025_v56), 0))
				_ = _index287
				(_2025_v56).ArraySet1(!((_2015_v46).Fm2((_2022_v53).F27(), (_this).F31(), globalState)), (_index287).Int())
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_2025_v56), 0))
				_ = _index288
				(_2025_v56).ArraySet1((_2022_v53).F26(), (_index288).Int())
				(globalState).F16 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2017_v48, _dafny.Companion_Sequence_.Concatenate(_2017_v48, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-868))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}((func(_2026_v3 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2027_i8 _dafny.Int) _dafny.CodePoint {
						return _2026_v3
					}
				})(_1942_v3)))))).Cardinality())
				var _2028_v57 _dafny.Sequence
				_ = _2028_v57
				_2028_v57 = _dafny.SeqOf((_this).F31())
				(globalState).F2 = (_1941_v2).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_2028_v57).Cardinality())).Plus((_this).F31()), _dafny.IntOfUint32((_1941_v2).Cardinality()))).Uint32()).(bool)
			} else {
				var _2029_v58 _dafny.Map
				_ = _2029_v58
				_2029_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F31())
				var _2030_v59 _dafny.Sequence
				_ = _2030_v59
				_2030_v59 = _dafny.SeqOf((_2029_v58).Cardinality(), (_this).F31())
				var _2031_v60 T3
				_ = _2031_v60
				var _nw313 *C8 = New_C8_()
				_ = _nw313
				_nw313.Ctor__(_dafny.MultiSetFromSeq(_2030_v59))
				_2031_v60 = _nw313
				var _2032_v61 _dafny.Sequence
				_ = _2032_v61
				_2032_v61 = _dafny.UnicodeSeqOfUtf8Bytes("qtfvyyvhs")
				var _2033_v62 _dafny.Map
				_ = _2033_v62
				_2033_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2032_v61, _1941_v2)
				var _2034_v63 _dafny.Map
				_ = _2034_v63
				_2034_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T3 {
					if (_this).F27() {
						return _2031_v60
					}
					return _2031_v60
				})(), _2033_v62)
				_2034_v63 = (_2034_v63).Update(_2031_v60, _2033_v62)
				var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))
				_ = _index289
				(_2011_v43).ArraySet1((_this).F31(), (_index289).Int())
				var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))
				_ = _index290
				var _rhs295 _dafny.Int = (_this).F31()
				_ = _rhs295
				var _rhs296 bool = true
				_ = _rhs296
				var _rhs297 _dafny.Int = (_dafny.IntOfInt64(771)).Minus((_this).F31())
				_ = _rhs297
				var _rhs298 _dafny.Int = (_this).F31()
				_ = _rhs298
				var _rhs299 bool = (_this).F27()
				_ = _rhs299
				var _lhs267 _dafny.Array = _2011_v43
				_ = _lhs267
				var _lhs268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))
				_ = _lhs268
				var _lhs269 *GlobalState = globalState
				_ = _lhs269
				var _lhs270 *GlobalState = globalState
				_ = _lhs270
				var _lhs271 *GlobalState = globalState
				_ = _lhs271
				var _lhs272 *GlobalState = globalState
				_ = _lhs272
				(_lhs267).ArraySet1(_rhs295, (_lhs268).Int())
				_lhs269.F22 = _rhs296
				_lhs270.F12 = _rhs297
				_lhs271.F12 = _rhs298
				_lhs272.F22 = _rhs299
				var _2035_v64 _dafny.MultiSet
				_ = _2035_v64
				_2035_v64 = _dafny.MultiSetOf((_2011_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))).Int()).(_dafny.Int))
				_2035_v64 = _2035_v64
				var _2036_v65 _dafny.MultiSet
				_ = _2036_v65
				_2036_v65 = _dafny.MultiSetOf(_1944_v5)
				var _2037_v66 _dafny.MultiSet
				_ = _2037_v66
				_2037_v66 = _2036_v65
				var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))
				_ = _index291
				var _rhs300 _dafny.Int = (_2015_v46).Fm1(globalState)
				_ = _rhs300
				var _rhs301 _dafny.Int = ((_2037_v66).Cardinality()).Times((func() _dafny.Int {
					if _this.F30 {
						return (_this).F31()
					}
					return _dafny.IntOfUint32((_2030_v59).Cardinality())
				})())
				_ = _rhs301
				var _rhs302 bool = p0
				_ = _rhs302
				var _rhs303 _dafny.MultiSet = ((_2031_v60).F29()).Intersection(((_2031_v60).F29()).Difference(_2035_v64))
				_ = _rhs303
				var _lhs273 *GlobalState = globalState
				_ = _lhs273
				var _lhs274 _dafny.Array = _2011_v43
				_ = _lhs274
				var _lhs275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2011_v43), 0))
				_ = _lhs275
				var _lhs276 *GlobalState = globalState
				_ = _lhs276
				_lhs273.F12 = _rhs300
				(_lhs274).ArraySet1(_rhs301, (_lhs275).Int())
				_lhs276.F15 = _rhs302
				_2035_v64 = _rhs303
				var _2038_v67 _dafny.MultiSet
				_ = _2038_v67
				_2038_v67 = _dafny.MultiSetOf(_1942_v3, _1942_v3, _dafny.CodePoint('u'), _1942_v3)
				var _2039_v68 _dafny.Array
				_ = _2039_v68
				var _nwElement0_78 bool = (_dafny.MultiSetOf(_1942_v3)).IsSubsetOf(_2038_v67)
				_ = _nwElement0_78
				var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(7))
				_ = _nw314
				(_nw314).ArraySet1(_nwElement0_78, 0)
				(_nw314).ArraySet1(p0, 1)
				(_nw314).ArraySet1(!(true) || (p2), 2)
				(_nw314).ArraySet1(!(!((_this).F26())), 3)
				(_nw314).ArraySet1(p2, 4)
				(_nw314).ArraySet1(_this.F30, 5)
				(_nw314).ArraySet1((p0) || ((_this).F27()), 6)
				_2039_v68 = _nw314
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2039_v68), 0))
				_ = _index292
				(_2039_v68).ArraySet1((_this).F26(), (_index292).Int())
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(355), _dafny.ArrayLen((_2039_v68), 0))
				_ = _index293
				(_2039_v68).ArraySet1((_this).F27(), (_index293).Int())
			}
			var _2040_v69 D2
			_ = _2040_v69
			_2040_v69 = Companion_D2_.Create_DC7_(p2, false, (_this).F31(), (((_this).F29()).Cardinality()).Minus((_this).F31()))
			var _rhs304 bool = (func() bool {
				if !(p2) {
					return p2
				}
				return p0
			})()
			_ = _rhs304
			var _rhs305 D2 = _2040_v69
			_ = _rhs305
			var _lhs277 *GlobalState = globalState
			_ = _lhs277
			_lhs277.F15 = _rhs304
			_2040_v69 = _rhs305
			var _2041_v70 _dafny.Array
			_ = _2041_v70
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_53
			var _nw315 _dafny.Array
			_ = _nw315
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw315 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) bool = func(_2042_i9 _dafny.Int) bool {
					return (_this).F26()
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw315 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw315).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw315).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_2041_v70 = _nw315
			var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_2041_v70), 0))
			_ = _index294
			(_2041_v70).ArraySet1((_2010_cf48).Fm2((_this).F26(), (_this).F31(), globalState), (_index294).Int())
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_2041_v70), 0))
			_ = _index295
			(_2041_v70).ArraySet1(p0, (_index295).Int())
		}
		var _2043_v71 _dafny.Array
		_ = _2043_v71
		var _nw316 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
		_ = _nw316
		_2043_v71 = _nw316
		for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2043_v71), 0))); ; {
			_guard_loop_0, _ok69 := _iter69()
			if !_ok69 {
				break
			}
			var _2044_i10 _dafny.Int
			_2044_i10 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_2044_i10).Sign() != -1) && ((_2044_i10).Cmp(_dafny.ArrayLen((_2043_v71), 0)) < 0)) {
				(_2043_v71).ArraySet1((Companion_D16_.Create_DC38_((_this).F31(), _dafny.IntOfInt64(928), _this.F30)).Dtor_cf71(), (_2044_i10).Int())
			}
		}
		var _hi18 _dafny.Int = (_this).F31()
		_ = _hi18
		for _2045_i11 := (_this).F31(); _2045_i11.Cmp(_hi18) < 0; _2045_i11 = _2045_i11.Plus(_dafny.One) {
			var _2046_v72 _dafny.Sequence
			_ = _2046_v72
			_2046_v72 = _dafny.UnicodeSeqOfUtf8Bytes("eobm")
			(globalState).F6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2046_v72, (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2046_v72).Cardinality()))).Uint32(), _dafny.CodePoint('e')), _2046_v72), (Companion_Default___.SafeIndex(_2045_i11, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2046_v72, (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2046_v72).Cardinality()))).Uint32(), _dafny.CodePoint('e')), _2046_v72)).Cardinality()))).Uint32(), _1942_v3), _2046_v72)
			(globalState).F6 = _2046_v72
			var _2047_v73 _dafny.Map
			_ = _2047_v73
			_2047_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _dafny.UnicodeSeqOfUtf8Bytes("yxokab"))
			var _2048_v74 _dafny.Sequence
			_ = _2048_v74
			_2048_v74 = _dafny.SeqOf(_2045_i11, (_this).F31())
			var _2049_v75 _dafny.Sequence
			_ = _2049_v75
			_2049_v75 = _dafny.SeqOf(_2048_v74)
			var _2050_v76 T2
			_ = _2050_v76
			var _nw317 *C2 = New_C2_()
			_ = _nw317
			_nw317.Ctor__(_2049_v75, _this.F30, p2)
			_2050_v76 = _nw317
			var _2051_v77 _dafny.Array
			_ = _2051_v77
			var _nwElement0_79 T2 = _2050_v76
			_ = _nwElement0_79
			var _nw318 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(3))
			_ = _nw318
			(_nw318).ArraySet1(_nwElement0_79, 0)
			(_nw318).ArraySet1(_2050_v76, 1)
			(_nw318).ArraySet1(_2050_v76, 2)
			_2051_v77 = _nw318
			var _2052_v78 _dafny.Sequence
			_ = _2052_v78
			_2052_v78 = _dafny.SeqOf(_2051_v77, _2051_v77)
			var _2053_v79 _dafny.Set
			_ = _2053_v79
			_2053_v79 = _dafny.SetOf((_this).F31(), _2045_i11, _dafny.IntOfUint32((_2046_v72).Cardinality()), (_2047_v73).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2052_v78, _dafny.SeqOf(_2051_v77, _2051_v77))).Cardinality()))
			_2053_v79 = (func() _dafny.Set {
				var _coll69 = _dafny.NewBuilder()
				_ = _coll69
				for _iter70 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(809), _dafny.IntOfInt64(689))); ; {
					_compr_69, _ok70 := _iter70()
					if !_ok70 {
						break
					}
					var _2054_v80 _dafny.Int
					_2054_v80 = interface{}(_compr_69).(_dafny.Int)
					if ((_dafny.IntOfInt64(809)).Cmp(_2054_v80) <= 0) && ((_2054_v80).Cmp(_dafny.IntOfInt64(689)) < 0) {
						_coll69.Add((_2054_v80).Plus((_this).F31()))
					}
				}
				return _coll69.ToSet()
			}()).Intersection(_2053_v79)
			(globalState).F12 = _2045_i11
		}
		(globalState).F12 = (_this).F31()
		r0 = _1942_v3
		r1 = (_this).F31()
		return r0, r1
	}
}
func (_this *C9) M1(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _2055_v0 D0
		_ = _2055_v0
		_2055_v0 = Companion_D0_.Create_DC0_((_this).F27())
		var _2056_v1 _dafny.Sequence
		_ = _2056_v1
		_2056_v1 = _dafny.SeqOf((_this).Fm4(_2055_v0, (_this).F31(), (_this).F31(), globalState))
		var _2057_v2 _dafny.MultiSet
		_ = _2057_v2
		_2057_v2 = _dafny.MultiSetOf(_2056_v1, _2056_v1)
		var _2058_v3 *C3
		_ = _2058_v3
		var _nw319 *C3 = New_C3_()
		_ = _nw319
		_nw319.Ctor__(_2057_v2, _this.F30, _this.F30)
		_2058_v3 = _nw319
		_2058_v3 = _2058_v3
		var _2059_v4 _dafny.CodePoint
		_ = _2059_v4
		_2059_v4 = _dafny.CodePoint('s')
		var _2060_v5 _dafny.Map
		_ = _2060_v5
		_2060_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2056_v1)
		var _2061_v6 _dafny.Map
		_ = _2061_v6
		_2061_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2059_v4, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_2060_v5).Contains((_this).F27()) {
				return (_2060_v5).Get((_this).F27()).(_dafny.Sequence)
			}
			return _2056_v1
		})()).Cardinality()))
		var _2062_v7 _dafny.Array
		_ = _2062_v7
		var _len0_54 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_54
		var _nw320 _dafny.Array
		_ = _nw320
		if _len0_54.Cmp(_dafny.Zero) == 0 {
			_nw320 = _dafny.NewArray(_len0_54)
		} else {
			var _init54 func(_dafny.Int) _dafny.CodePoint = (func(_2063_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2064_i1 _dafny.Int) _dafny.CodePoint {
					return _2063_v4
				}
			})(_2059_v4)
			_ = _init54
			var _element0_54 = _init54(_dafny.Zero)
			_ = _element0_54
			_nw320 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
			(_nw320).ArraySet1CodePoint(_element0_54, 0)
			var _nativeLen0_54 = (_len0_54).Int()
			_ = _nativeLen0_54
			for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
				(_nw320).ArraySet1CodePoint(_init54(_dafny.IntOf(_i0_54)), _i0_54)
			}
		}
		_2062_v7 = _nw320
		var _2065_v8 _dafny.Sequence
		_ = _2065_v8
		_2065_v8 = _dafny.UnicodeSeqOfUtf8Bytes("kbutbaoar")
		var _2066_v9 D1
		_ = _2066_v9
		_2066_v9 = Companion_D1_.Create_DC4_(_2062_v7, (_this).F31(), (_this).F31(), _2065_v8, _2065_v8)
		var _hi19 _dafny.Int = (_2066_v9).Dtor_cf8()
		_ = _hi19
		for _2067_i0 := ((_2058_v3).Fm12((_2061_v6).Cardinality(), globalState)).Plus((_this).F31()); _2067_i0.Cmp(_hi19) < 0; _2067_i0 = _2067_i0.Plus(_dafny.One) {
			_2065_v8 = _dafny.Companion_Sequence_.Concatenate(_2065_v8, _2065_v8)
			var _2068_v10 _dafny.Array
			_ = _2068_v10
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_55
			var _nw321 _dafny.Array
			_ = _nw321
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw321 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) _dafny.Int = (func(_2069_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_2070_i2 _dafny.Int) _dafny.Int {
						return (_2070_i2).Times((_dafny.SetOf((_this).F27(), (_this).F27(), _2069_p0)).Cardinality())
					}
				})(p0)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw321 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw321).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw321).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_2068_v10 = _nw321
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))
			_ = _index296
			(_2068_v10).ArraySet1(((_this).F31()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sw")).Cardinality())), (_index296).Int())
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))
			_ = _index297
			(_2068_v10).ArraySet1((_this).F31(), (_index297).Int())
			var _2071_v11 _dafny.Map
			_ = _2071_v11
			_2071_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, (_this).Fm3(_2065_v8, p2, (_this).F31(), globalState))
			var _2072_v12 D5
			_ = _2072_v12
			_2072_v12 = Companion_D5_.Create_DC13_((_2071_v11).Merge(_2071_v11))
			var _source29 D5 = _2072_v12
			_ = _source29
			if _source29.Is_DC14() {
				var _2073___mcc_h0 _dafny.Int = _source29.Get_().(D5_DC14).Cf25
				_ = _2073___mcc_h0
				var _2074___mcc_h1 _dafny.Int = _source29.Get_().(D5_DC14).Cf26
				_ = _2074___mcc_h1
				var _2075___mcc_h2 bool = _source29.Get_().(D5_DC14).Cf27
				_ = _2075___mcc_h2
				var _2076_cf27 bool = _2075___mcc_h2
				_ = _2076_cf27
				var _2077_cf26 _dafny.Int = _2074___mcc_h1
				_ = _2077_cf26
				var _2078_cf25 _dafny.Int = _2073___mcc_h0
				_ = _2078_cf25
				var _2079_v13 _dafny.Map
				_ = _2079_v13
				_2079_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("hhvqkm"))
				var _2080_v14 _dafny.Map
				_ = _2080_v14
				_2080_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2078_cf25)
				var _2081_v15 _dafny.Sequence
				_ = _2081_v15
				_2081_v15 = _dafny.SeqOf(_2080_v14, _2080_v14)
				var _2082_v16 D10
				_ = _2082_v16
				_2082_v16 = Companion_D10_.Create_DC24_((func() _dafny.Sequence {
					if (_2079_v13).Contains(_2067_i0) {
						return (_2079_v13).Get(_2067_i0).(_dafny.Sequence)
					}
					return _2065_v8
				})(), _2081_v15, (_this).F26())
				_2082_v16 = _2082_v16
				(globalState).F12 = Companion_Default___.SafeModuloInt((_this).F31(), _2077_cf26)
				var _2083_v17 _dafny.Sequence
				_ = _2083_v17
				_2083_v17 = _dafny.SeqOf(_2056_v1, _2056_v1)
				var _2084_v18 *C2
				_ = _2084_v18
				var _nw322 *C2 = New_C2_()
				_ = _nw322
				_nw322.Ctor__(_2083_v17, true, _this.F30)
				_2084_v18 = _nw322
				var _2085_v19 _dafny.Map
				_ = _2085_v19
				_2085_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2058_v3).Fm1(globalState), p0)
				var _2086_v20 _dafny.MultiSet
				_ = _2086_v20
				_2086_v20 = _dafny.MultiSetOf(_2085_v19)
				var _2087_v21 _dafny.MultiSet
				_ = _2087_v21
				_2087_v21 = _2086_v20
				_2087_v21 = _2086_v20
			} else if _source29.Is_DC13() {
				var _2088___mcc_h3 _dafny.Map = _source29.Get_().(D5_DC13).Cf24
				_ = _2088___mcc_h3
				var _2089_cf24 _dafny.Map = _2088___mcc_h3
				_ = _2089_cf24
				(globalState).F15 = (_this).F27()
				var _2090_v22 _dafny.Map
				_ = _2090_v22
				_2090_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2056_v1, (_dafny.Zero).Minus((_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int)))
				_2090_v22 = (_2090_v22).Update(_2056_v1, (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int))
				_2059_v4 = _2059_v4
				var _2091_v23 D6
				_ = _2091_v23
				_2091_v23 = Companion_D6_.Create_DC16_(_2065_v8)
				var _2092_v24 _dafny.Map
				_ = _2092_v24
				_2092_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int))
				var _2093_v25 _dafny.Map
				_ = _2093_v25
				_2093_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2067_i0), _2056_v1), (Companion_Default___.SafeIndex((_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2067_i0), _2056_v1)).Cardinality()))).Uint32(), _dafny.IntOfUint32(((_2091_v23).Dtor_cf29()).Cardinality())), !(_2092_v24).Equals(_2092_v24))
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))
				_ = _index298
				(_2068_v10).ArraySet1((_2093_v25).Cardinality(), (_index298).Int())
			} else {
				var _2094___mcc_h4 D5 = _source29.Get_().(D5_DC15).Cf28
				_ = _2094___mcc_h4
				var _2095_cf28 D5 = _2094___mcc_h4
				_ = _2095_cf28
				var _2096_v26 _dafny.Map
				_ = _2096_v26
				_2096_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2067_i0, (_this).Fm2((p2).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool), _2067_i0, globalState))
				var _2097_v27 D16
				_ = _2097_v27
				_2097_v27 = Companion_D16_.Create_DC37_((_2096_v26).Update((_this).F31(), p0))
				var _2098_v28 _dafny.Map
				_ = _2098_v28
				_2098_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F31()), (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int))
				var _2099_v29 _dafny.Map
				_ = _2099_v29
				_2099_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2097_v27).Dtor_cf68(), _2098_v28)
				var _2100_v31 _dafny.Map
				_ = _2100_v31
				_2100_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(103), (_this).F26()), (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int))
				var _2101_v32 _dafny.Map
				_ = _2101_v32
				_2101_v32 = _2100_v31
				_2099_v29 = (_2099_v29).Merge(func() _dafny.Map {
					var _coll70 = _dafny.NewMapBuilder()
					_ = _coll70
					for _iter71 := _dafny.Iterate((_2101_v32).Keys().Elements()); ; {
						_compr_70, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _2102_v30 _dafny.Map
						_2102_v30 = interface{}(_compr_70).(_dafny.Map)
						if (_2101_v32).Contains(_2102_v30) {
							_coll70.Add(_2102_v30, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F26())).Cardinality(), (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int)))
						}
					}
					return _coll70.ToMap()
				}())
				var _2103_v33 _dafny.Map
				_ = _2103_v33
				_2103_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2065_v8).Cardinality()), p2)
				var _2104_v34 _dafny.Map
				_ = _2104_v34
				_2104_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_2065_v8).Cardinality()))
				var _2105_v35 D15
				_ = _2105_v35
				_2105_v35 = Companion_D15_.Create_DC34_((_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int), (_this).F26(), (_this).F31())
				var _2106_v36 _dafny.MultiSet
				_ = _2106_v36
				_2106_v36 = _dafny.MultiSetOf(_2105_v35)
				var _2107_v37 _dafny.Map
				_ = _2107_v37
				_2107_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if (_2103_v33).Contains((_2104_v34).Cardinality()) {
						return (_2103_v33).Get((_2104_v34).Cardinality()).(_dafny.Sequence)
					}
					return p2
				})(), (_2106_v36).Union(_2106_v36))
				_2107_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex((_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F27()), _2106_v36)
				var _2108_v38 *C0
				_ = _2108_v38
				var _nw323 *C0 = New_C0_()
				_ = _nw323
				_nw323.Ctor__(true)
				_2108_v38 = _nw323
				var _2109_v39 _dafny.Map
				_ = _2109_v39
				_2109_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2108_v38).F32(), (_this).F29())
				var _2110_v40 _dafny.MultiSet
				_ = _2110_v40
				_2110_v40 = _dafny.MultiSetOf(p0)
				var _rhs306 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _this.F30 {
						return _2059_v4
					}
					return _2059_v4
				})()
				_ = _rhs306
				var _rhs307 *C0 = _2108_v38
				_ = _rhs307
				var _rhs308 _dafny.Int = (((_2109_v39).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(p0, (_2068_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(745), _dafny.ArrayLen((_2068_v10), 0))).Int()).(_dafny.Int), globalState), (_this).F29()))).Cardinality()).Times(_dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()))
				_ = _rhs308
				var _rhs309 _dafny.Sequence = _2065_v8
				_ = _rhs309
				var _rhs310 bool = !((_2110_v40).IsSubsetOf(_2110_v40))
				_ = _rhs310
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				var _lhs280 *GlobalState = globalState
				_ = _lhs280
				var _lhs281 *GlobalState = globalState
				_ = _lhs281
				_lhs278.F19 = _rhs306
				_2108_v38 = _rhs307
				_lhs279.F16 = _rhs308
				_lhs280.F6 = _rhs309
				_lhs281.F15 = _rhs310
				var _2111_v41 _dafny.Array
				_ = _2111_v41
				var _nw324 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
				_ = _nw324
				_2111_v41 = _nw324
				var _2112_v42 _dafny.Map
				_ = _2112_v42
				_2112_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_this).F26())
				var _2113_v43 D13
				_ = _2113_v43
				_2113_v43 = Companion_D13_.Create_DC29_(_2112_v42)
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2111_v41), 0))
				_ = _index299
				(_2111_v41).ArraySet1(((_2113_v43).Dtor_cf50()).Merge(_2112_v42), (_index299).Int())
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_2111_v41), 0))
				_ = _index300
				(_2111_v41).ArraySet1(_2112_v42, (_index300).Int())
			}
			var _2114_v44 _dafny.Map
			_ = _2114_v44
			_2114_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2068_v10, (p2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))
			_2114_v44 = (func() _dafny.Map {
				if _this.F30 {
					return _2114_v44
				}
				return _2114_v44
			})()
		}
		var _2115_v45 _dafny.MultiSet
		_ = _2115_v45
		_2115_v45 = _dafny.MultiSetOf((_this).F26())
		var _2116_v46 _dafny.Set
		_ = _2116_v46
		_2116_v46 = _dafny.SetOf(false, (_this).F27())
		var _2117_v47 _dafny.Map
		_ = _2117_v47
		_2117_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2115_v45, (_2116_v46).Cardinality())
		var _2118_v48 D0
		_ = _2118_v48
		_2118_v48 = Companion_D0_.Create_DC1_(p0, false, _dafny.IntOfUint32((_2065_v8).Cardinality()), (p2).Select((Companion_Default___.SafeIndex((_2117_v47).Cardinality(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool))
		var _2119_v49 D0
		_ = _2119_v49
		_2119_v49 = Companion_D0_.Create_DC2_(_2118_v48)
		(globalState).F12 = (_dafny.Zero).Minus(func(_source30 D0) _dafny.Int {
			if _source30.Is_DC1() {
				var _2120___mcc_h5 bool = _source30.Get_().(D0_DC1).Cf1
				_ = _2120___mcc_h5
				var _2121___mcc_h6 bool = _source30.Get_().(D0_DC1).Cf2
				_ = _2121___mcc_h6
				var _2122___mcc_h7 _dafny.Int = _source30.Get_().(D0_DC1).Cf3
				_ = _2122___mcc_h7
				var _2123___mcc_h8 bool = _source30.Get_().(D0_DC1).Cf4
				_ = _2123___mcc_h8
				var _2124_cf4 bool = _2123___mcc_h8
				_ = _2124_cf4
				var _2125_cf3 _dafny.Int = _2122___mcc_h7
				_ = _2125_cf3
				var _2126_cf2 bool = _2121___mcc_h6
				_ = _2126_cf2
				var _2127_cf1 bool = _2120___mcc_h5
				_ = _2127_cf1
				return _2125_cf3
			} else if _source30.Is_DC0() {
				var _2128___mcc_h9 bool = _source30.Get_().(D0_DC0).Cf0
				_ = _2128___mcc_h9
				var _2129_cf0 bool = _2128___mcc_h9
				_ = _2129_cf0
				return _dafny.IntOfInt64(691)
			} else {
				var _2130___mcc_h10 D0 = _source30.Get_().(D0_DC2).Cf5
				_ = _2130___mcc_h10
				var _2131_cf5 D0 = _2130___mcc_h10
				_ = _2131_cf5
				return (_this).F31()
			}
		}(_2119_v49))
		var _pat_let_tv58 = _2115_v45
		_ = _pat_let_tv58
		var _2132_v50 _dafny.Map
		_ = _2132_v50
		_2132_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let59_0 D7) D7 {
			return func(_2135_dt__update__tmp_h0 D7) D7 {
				return func(_pat_let60_0 _dafny.MultiSet) D7 {
					return func(_2136_dt__update_hcf37_h0 _dafny.MultiSet) D7 {
						return func(_pat_let61_0 _dafny.Int) D7 {
							return func(_2137_dt__update_hcf36_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC19_(_2137_dt__update_hcf36_h0, _2136_dt__update_hcf37_h0, (_2135_dt__update__tmp_h0).Dtor_cf38(), (_2135_dt__update__tmp_h0).Dtor_cf39())
							}(_pat_let61_0)
						}((_this).F31())
					}(_pat_let60_0)
				}(_pat_let_tv58)
			}(_pat_let59_0)
		}(Companion_D7_.Create_DC19_((_this).F31(), _2115_v45, (_2058_v3).Fm2(false, (_this).F31(), globalState), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(359))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg111 _dafny.Int) interface{} {
				return coer111(arg111)
			}
		}((func(_2133_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2134_i4 _dafny.Int) _dafny.CodePoint {
				return _2133_v4
			}
		})(_2059_v4)))).Cardinality()))))).Dtor_cf38(), (_this).F27())
		var _2138_i3 _dafny.Int
		_ = _2138_i3
		_2138_i3 = _dafny.Zero
		{
			for (func() bool {
				if (_2132_v50).Contains(p0) {
					return (_2132_v50).Get(p0).(bool)
				}
				return (_this).F27()
			})() {
				{
					if (_2138_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_2138_i3 = (_2138_i3).Plus(_dafny.One)
					var _2139_v51 _dafny.Sequence
					_ = _2139_v51
					_2139_v51 = _dafny.SeqOf((_this).F29())
					var _2140_v52 _dafny.Map
					_ = _2140_v52
					_2140_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfUint32((_2065_v8).Cardinality()))
					(globalState).F2 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(612))).Uint32(), func(coer112 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg112 _dafny.Int) interface{} {
							return coer112(arg112)
						}
					}(func(_2141_i5 _dafny.Int) _dafny.Int {
						return (_this).F31()
					}))).Cardinality()), (_dafny.Zero).Minus((_this).F31()))).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
						if (_this).F27() {
							return (_dafny.SetOf((_this).F29(), (_2139_v51).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2139_v51).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality()
						}
						return (func() _dafny.Int {
							if (_2140_v52).Contains((_this).F26()) {
								return (_2140_v52).Get((_this).F26()).(_dafny.Int)
							}
							return (_this).F31()
						})()
					})())) > 0
					var _2142_v53 T1
					_ = _2142_v53
					var _nw325 *C3 = New_C3_()
					_ = _nw325
					_nw325.Ctor__(_dafny.MultiSetOf(_dafny.SeqOf((_this).F31()), _dafny.Companion_Sequence_.Update(_2056_v1, (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2056_v1).Cardinality()))).Uint32(), _dafny.IntOfInt64(-473))), false, p0)
					_2142_v53 = _nw325
					var _2143_v54 D20
					_ = _2143_v54
					_2143_v54 = Companion_D20_.Create_DC46_(_2142_v53)
					var _2144_v55 _dafny.Array
					_ = _2144_v55
					var _nwElement0_80 T1 = _2142_v53
					_ = _nwElement0_80
					var _nw326 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(28))
					_ = _nw326
					(_nw326).ArraySet1(_nwElement0_80, 0)
					(_nw326).ArraySet1(_2142_v53, 1)
					(_nw326).ArraySet1(_2142_v53, 2)
					(_nw326).ArraySet1(_2142_v53, 3)
					(_nw326).ArraySet1(_2142_v53, 4)
					(_nw326).ArraySet1(_2142_v53, 5)
					(_nw326).ArraySet1(_2142_v53, 6)
					(_nw326).ArraySet1(_2142_v53, 7)
					(_nw326).ArraySet1(_2142_v53, 8)
					(_nw326).ArraySet1(_2142_v53, 9)
					(_nw326).ArraySet1(_2142_v53, 10)
					(_nw326).ArraySet1(_2142_v53, 11)
					(_nw326).ArraySet1(_2142_v53, 12)
					(_nw326).ArraySet1(_2142_v53, 13)
					(_nw326).ArraySet1(_2142_v53, 14)
					(_nw326).ArraySet1((_2143_v54).Dtor_cf84(), 15)
					(_nw326).ArraySet1(_2142_v53, 16)
					(_nw326).ArraySet1(_2142_v53, 17)
					(_nw326).ArraySet1(_2142_v53, 18)
					(_nw326).ArraySet1(_2142_v53, 19)
					(_nw326).ArraySet1(_2142_v53, 20)
					(_nw326).ArraySet1(_2142_v53, 21)
					(_nw326).ArraySet1(_2142_v53, 22)
					(_nw326).ArraySet1(_2142_v53, 23)
					(_nw326).ArraySet1(_2142_v53, 24)
					(_nw326).ArraySet1(_2142_v53, 25)
					(_nw326).ArraySet1(_2142_v53, 26)
					(_nw326).ArraySet1(_2142_v53, 27)
					_2144_v55 = _nw326
					var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_2144_v55), 0))
					_ = _index301
					(_2144_v55).ArraySet1(_2142_v53, (_index301).Int())
					var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(309), _dafny.ArrayLen((_2144_v55), 0))
					_ = _index302
					(_2144_v55).ArraySet1(_2142_v53, (_index302).Int())
					var _2145_v56 _dafny.Map
					_ = _2145_v56
					_2145_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_2142_v53).F27())
					var _2146_v57 _dafny.Array
					_ = _2146_v57
					var _nwElement0_81 bool = (func() bool {
						if (_2145_v56).Contains((_this).F31()) {
							return (_2145_v56).Get((_this).F31()).(bool)
						}
						return (_2142_v53).F27()
					})()
					_ = _nwElement0_81
					var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(3))
					_ = _nw327
					(_nw327).ArraySet1(_nwElement0_81, 0)
					(_nw327).ArraySet1((_this).F26(), 1)
					(_nw327).ArraySet1((_this).F27(), 2)
					_2146_v57 = _nw327
					var _2147_v58 D9
					_ = _2147_v58
					_2147_v58 = Companion_D9_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2146_v57))
					var _2148_v59 _dafny.Array
					_ = _2148_v59
					var _len0_56 _dafny.Int = _dafny.IntOfInt64(20)
					_ = _len0_56
					var _nw328 _dafny.Array
					_ = _nw328
					if _len0_56.Cmp(_dafny.Zero) == 0 {
						_nw328 = _dafny.NewArray(_len0_56)
					} else {
						var _init56 func(_dafny.Int) _dafny.Map = (func(_2149_v52 _dafny.Map) func(_dafny.Int) _dafny.Map {
							return func(_2150_i6 _dafny.Int) _dafny.Map {
								return _2149_v52
							}
						})(_2140_v52)
						_ = _init56
						var _element0_56 = _init56(_dafny.Zero)
						_ = _element0_56
						_nw328 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
						(_nw328).ArraySet1(_element0_56, 0)
						var _nativeLen0_56 = (_len0_56).Int()
						_ = _nativeLen0_56
						for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
							(_nw328).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
						}
					}
					_2148_v59 = _nw328
					var _2151_v60 _dafny.Map
					_ = _2151_v60
					_2151_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((p2).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((p2).Cardinality()))).Uint32()).(bool)), _2146_v57)
					var _rhs311 *C3 = _2058_v3
					_ = _rhs311
					var _rhs312 D9 = Companion_D9_.Create_DC21_(_2151_v60)
					_ = _rhs312
					var _rhs313 _dafny.Array = _2148_v59
					_ = _rhs313
					var _rhs314 _dafny.Int = _dafny.IntOfInt64(110)
					_ = _rhs314
					var _lhs282 *GlobalState = globalState
					_ = _lhs282
					_2058_v3 = _rhs311
					_2147_v58 = _rhs312
					_2148_v59 = _rhs313
					_lhs282.F12 = _rhs314
					var _2152_v61 D1
					_ = _2152_v61
					_2152_v61 = Companion_D1_.Create_DC5_((_this).F31())
					_2152_v61 = _2152_v61
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _2153_v62 _dafny.Array
		_ = _2153_v62
		var _len0_57 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_57
		var _nw329 _dafny.Array
		_ = _nw329
		if _len0_57.Cmp(_dafny.Zero) == 0 {
			_nw329 = _dafny.NewArray(_len0_57)
		} else {
			var _init57 func(_dafny.Int) _dafny.Int = func(_2154_i7 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_2154_i7, (_this).F31())
			}
			_ = _init57
			var _element0_57 = _init57(_dafny.Zero)
			_ = _element0_57
			_nw329 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
			(_nw329).ArraySet1(_element0_57, 0)
			var _nativeLen0_57 = (_len0_57).Int()
			_ = _nativeLen0_57
			for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
				(_nw329).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
			}
		}
		_2153_v62 = _nw329
		var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_2153_v62), 0))
		_ = _index303
		(_2153_v62).ArraySet1((_this).F31(), (_index303).Int())
		var _2155_v63 _dafny.Array
		_ = _2155_v63
		var _nw330 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
		_ = _nw330
		_2155_v63 = _nw330
		var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_2155_v63), 0))
		_ = _index304
		(_2155_v63).ArraySet1(_2153_v62, (_index304).Int())
		var _2156_v64 _dafny.Sequence
		_ = _2156_v64
		_2156_v64 = _dafny.SeqOf(_dafny.SetOf((_this).F27()))
		var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_2153_v62), 0))
		_ = _index305
		var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_2155_v63), 0))
		_ = _index306
		var _rhs315 _dafny.Int = (((_2156_v64).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2156_v64).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()).Minus((_this).F31())
		_ = _rhs315
		var _rhs316 bool = (_this).F26()
		_ = _rhs316
		var _rhs317 _dafny.Array = _2153_v62
		_ = _rhs317
		var _rhs318 _dafny.Array = _2153_v62
		_ = _rhs318
		var _lhs283 _dafny.Array = _2153_v62
		_ = _lhs283
		var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_2153_v62), 0))
		_ = _lhs284
		var _lhs285 *GlobalState = globalState
		_ = _lhs285
		var _lhs286 _dafny.Array = _2155_v63
		_ = _lhs286
		var _lhs287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_2155_v63), 0))
		_ = _lhs287
		(_lhs283).ArraySet1(_rhs315, (_lhs284).Int())
		_lhs285.F15 = _rhs316
		(_lhs286).ArraySet1(_rhs317, (_lhs287).Int())
		_2153_v62 = _rhs318
		(globalState).F15 = (_this).F27()
	}
}
func (_this *C9) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Array, T0) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 T0 = (T0)(nil)
		_ = r3
		var _2157_v0 _dafny.Sequence
		_ = _2157_v0
		_2157_v0 = _dafny.SeqOf(p3)
		(globalState).F16 = (_2157_v0).Select((Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32()).(_dafny.Int)
		var _hi20 _dafny.Int = _dafny.IntOfInt64(-364)
		_ = _hi20
		for _2158_i0 := p3; _2158_i0.Cmp(_hi20) < 0; _2158_i0 = _2158_i0.Plus(_dafny.One) {
			var _2159_v1 _dafny.Array
			_ = _2159_v1
			var _nw331 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(29))
			_ = _nw331
			_2159_v1 = _nw331
			_2159_v1 = _2159_v1
			var _2160_v2 _dafny.Map
			_ = _2160_v2
			_2160_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F31()), (_this).F27())
			_2160_v2 = (_2160_v2).Update(p0, _this.F30)
			var _2161_v3 _dafny.Array
			_ = _2161_v3
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_58
			var _nw332 _dafny.Array
			_ = _nw332
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw332 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) bool = func(_2162_i1 _dafny.Int) bool {
					return false
				}
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw332 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw332).ArraySet1(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw332).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_2161_v3 = _nw332
			var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2161_v3), 0))
			_ = _index307
			(_2161_v3).ArraySet1((_this).F27(), (_index307).Int())
			var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2161_v3), 0))
			_ = _index308
			var _rhs319 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-29))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg113 _dafny.Int) interface{} {
					return coer113(arg113)
				}
			}(func(_2163_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			})), p2)
			_ = _rhs319
			var _rhs320 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(p2, Companion_Default___.Fm22((_dafny.Zero).Minus((_this).F31()), _this.F30, globalState)), p2)
			_ = _rhs320
			var _rhs321 bool = (_this).F26()
			_ = _rhs321
			var _lhs288 *GlobalState = globalState
			_ = _lhs288
			var _lhs289 *GlobalState = globalState
			_ = _lhs289
			var _lhs290 _dafny.Array = _2161_v3
			_ = _lhs290
			var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2161_v3), 0))
			_ = _lhs291
			_lhs288.F6 = _rhs319
			_lhs289.F2 = _rhs320
			(_lhs290).ArraySet1(_rhs321, (_lhs291).Int())
			(globalState).F16 = _dafny.IntOfInt64(763)
		}
		var _2164_i3 _dafny.Int
		_ = _2164_i3
		_2164_i3 = _dafny.Zero
		{
			for !((p3).Cmp(_dafny.IntOfInt64(-399)) <= 0) {
				{
					if (_2164_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_2164_i3 = (_2164_i3).Plus(_dafny.One)
					var _2165_v4 _dafny.Map
					_ = _2165_v4
					_2165_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F30)
					var _2166_v5 _dafny.Map
					_ = _2166_v5
					_2166_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, p1)
					var _2167_v6 _dafny.Map
					_ = _2167_v6
					_2167_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2165_v4).Cardinality(), _2166_v5)
					var _2168_v7 _dafny.MultiSet
					_ = _2168_v7
					_2168_v7 = _dafny.MultiSetOf((_this).F26())
					var _2169_v8 _dafny.CodePoint
					_ = _2169_v8
					_2169_v8 = _dafny.CodePoint('o')
					var _2170_v9 _dafny.Int
					_ = _2170_v9
					var _2171_v10 _dafny.Int
					_ = _2171_v10
					var _2172_v11 bool
					_ = _2172_v11
					var _out64 _dafny.Int
					_ = _out64
					var _out65 _dafny.Int
					_ = _out65
					var _out66 bool
					_ = _out66
					_out64, _out65, _out66 = (_this).M6((func() _dafny.Map {
						if (_2167_v6).Contains(p1) {
							return (_2167_v6).Get(p1).(_dafny.Map)
						}
						return Companion_Default___.Fm30((_2168_v7).Cardinality(), true, globalState)
					})(), (Companion_Default___.Fm39(globalState)).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), p0)).Cardinality()), _2169_v8, p1, globalState)
					_2170_v9 = _out64
					_2171_v10 = _out65
					_2172_v11 = _out66
					var _2173_v12 _dafny.Array
					_ = _2173_v12
					var _len0_59 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_59
					var _nw333 _dafny.Array
					_ = _nw333
					if _len0_59.Cmp(_dafny.Zero) == 0 {
						_nw333 = _dafny.NewArray(_len0_59)
					} else {
						var _init59 func(_dafny.Int) _dafny.Int = (func(_2174_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2175_i4 _dafny.Int) _dafny.Int {
								return (_2175_i4).Times((_dafny.MultiSetOf((_this).F31(), _2174_p0)).Cardinality())
							}
						})(p0)
						_ = _init59
						var _element0_59 = _init59(_dafny.Zero)
						_ = _element0_59
						_nw333 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
						(_nw333).ArraySet1(_element0_59, 0)
						var _nativeLen0_59 = (_len0_59).Int()
						_ = _nativeLen0_59
						for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
							(_nw333).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
						}
					}
					_2173_v12 = _nw333
					var _2176_v13 T3
					_ = _2176_v13
					var _nw334 *C8 = New_C8_()
					_ = _nw334
					_nw334.Ctor__((_this).F29())
					_2176_v13 = _nw334
					var _2177_v14 _dafny.MultiSet
					_ = _2177_v14
					_2177_v14 = _dafny.MultiSetOf(_2176_v13)
					var _2178_v15 _dafny.Array
					_ = _2178_v15
					var _nwElement0_82 bool = _this.F30
					_ = _nwElement0_82
					var _nw335 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(10))
					_ = _nw335
					(_nw335).ArraySet1(_nwElement0_82, 0)
					(_nw335).ArraySet1((_this).F26(), 1)
					(_nw335).ArraySet1((_this).F26(), 2)
					(_nw335).ArraySet1(true, 3)
					(_nw335).ArraySet1((_this).Fm2((_this).F26(), (_2177_v14).Cardinality(), globalState), 4)
					(_nw335).ArraySet1(_2172_v11, 5)
					(_nw335).ArraySet1(_2172_v11, 6)
					(_nw335).ArraySet1(_this.F30, 7)
					(_nw335).ArraySet1(_this.F30, 8)
					(_nw335).ArraySet1((_this).F26(), 9)
					_2178_v15 = _nw335
					var _2179_v16 _dafny.Sequence
					_ = _2179_v16
					_2179_v16 = _dafny.SeqOf(_2178_v15, _2178_v15)
					var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_2173_v12), 0))
					_ = _index309
					(_2173_v12).ArraySet1((func() _dafny.Int {
						if (_2168_v7).Contains(_2172_v11) {
							return (_2168_v7).Multiplicity(_2172_v11)
						}
						return _dafny.IntOfUint32((_2179_v16).Cardinality())
					})(), (_index309).Int())
					var _2180_v17 _dafny.Map
					_ = _2180_v17
					_2180_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2166_v5).Cardinality(), _dafny.IntOfUint32((p2).Cardinality()))
					var _2181_v18 _dafny.Map
					_ = _2181_v18
					_2181_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2172_v11, _2180_v17)
					var _2182_v19 D21
					_ = _2182_v19
					_2182_v19 = Companion_D21_.Create_DC49_(_2181_v18)
					var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_2173_v12), 0))
					_ = _index310
					(_2173_v12).ArraySet1((Companion_Default___.SafeModuloInt(_2170_v9, p0)).Minus(Companion_Default___.SafeDivisionInt(_2171_v10, ((_2182_v19).Dtor_cf86()).Cardinality())), (_index310).Int())
					var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_2178_v15), 0))
					_ = _index311
					(_2178_v15).ArraySet1(_2172_v11, (_index311).Int())
					var _2183_v20 _dafny.Map
					_ = _2183_v20
					_2183_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _2169_v8)
					var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_2178_v15), 0))
					_ = _index312
					(_2178_v15).ArraySet1(((_2183_v20).Cardinality()).Cmp(p1) > 0, (_index312).Int())
					var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(586), _dafny.ArrayLen((_2178_v15), 0))
					_ = _index313
					(_2178_v15).ArraySet1(!(_dafny.MultiSetOf((_this).F27(), (_this).Fm2((_this).F27(), (_2173_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_2173_v12), 0))).Int()).(_dafny.Int), globalState), false)).Contains((_this).F27()), (_index313).Int())
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		if (_this).F27() {
			var _2184_v21 _dafny.Array
			_ = _2184_v21
			var _nwElement0_83 bool = (_this).F26()
			_ = _nwElement0_83
			var _nw336 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.One)
			_ = _nw336
			(_nw336).ArraySet1(_nwElement0_83, 0)
			_2184_v21 = _nw336
			var _2185_v22 D13
			_ = _2185_v22
			_2185_v22 = Companion_D13_.Create_DC30_(_this.F30, (_this).F31(), _this.F30, _2184_v21)
			var _2186_v23 _dafny.Sequence
			_ = _2186_v23
			_2186_v23 = _dafny.SeqOf(_this.F30, _this.F30)
			var _2187_v24 T0
			_ = _2187_v24
			var _nw337 *C6 = New_C6_()
			_ = _nw337
			_nw337.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_2186_v23).Cardinality()))).Cardinality())
			_2187_v24 = _nw337
			var _2188_v25 _dafny.Map
			_ = _2188_v25
			_2188_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _2187_v24)
			var _2189_v26 _dafny.Set
			_ = _2189_v26
			_2189_v26 = _dafny.SetOf((func() T0 {
				if (_2188_v25).Contains((_this).F26()) {
					return (_2188_v25).Get((_this).F26()).(T0)
				}
				return _2187_v24
			})())
			var _2190_v27 D22
			_ = _2190_v27
			_2190_v27 = Companion_D22_.Create_DC53_(_2189_v26)
			var _2191_v28 _dafny.MultiSet
			_ = _2191_v28
			_2191_v28 = _dafny.MultiSetOf((_this).F27())
			var _2192_v29 _dafny.Array
			_ = _2192_v29
			var _nwElement0_84 bool = (p3).Cmp(p0) >= 0
			_ = _nwElement0_84
			var _nw338 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(18))
			_ = _nw338
			(_nw338).ArraySet1(_nwElement0_84, 0)
			(_nw338).ArraySet1((_this).F27(), 1)
			(_nw338).ArraySet1(!((_this).F27()) || ((_this).F27()), 2)
			(_nw338).ArraySet1(!((_2185_v22).Dtor_cf51()), 3)
			(_nw338).ArraySet1(_this.F30, 4)
			(_nw338).ArraySet1((_this).F27(), 5)
			(_nw338).ArraySet1(!((_2190_v27).Dtor_cf97()).Equals(_2189_v26), 6)
			(_nw338).ArraySet1((_this).F27(), 7)
			(_nw338).ArraySet1(!((_this).F27()) || (_this.F30), 8)
			(_nw338).ArraySet1(_this.F30, 9)
			(_nw338).ArraySet1((_this).F26(), 10)
			(_nw338).ArraySet1((_this).F26(), 11)
			(_nw338).ArraySet1(((_this).F26()) && ((_this).F27()), 12)
			(_nw338).ArraySet1(!((_2191_v28).IsDisjointFrom(_dafny.MultiSetFromSeq(_2186_v23))), 13)
			(_nw338).ArraySet1(_this.F30, 14)
			(_nw338).ArraySet1(((func() _dafny.Int {
				if (_2191_v28).Contains(_this.F30) {
					return (_2191_v28).Multiplicity(_this.F30)
				}
				return p3
			})()).Cmp(_dafny.IntOfInt64(-253)) < 0, 15)
			(_nw338).ArraySet1((_this).Fm2((_this).F26(), _dafny.IntOfUint32((p2).Cardinality()), globalState), 16)
			(_nw338).ArraySet1((_this).F27(), 17)
			_2192_v29 = _nw338
			_2192_v29 = _2184_v21
			var _2193_v30 _dafny.Array
			_ = _2193_v30
			var _len0_60 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_60
			var _nw339 _dafny.Array
			_ = _nw339
			if _len0_60.Cmp(_dafny.Zero) == 0 {
				_nw339 = _dafny.NewArray(_len0_60)
			} else {
				var _init60 func(_dafny.Int) _dafny.Int = func(_2194_i5 _dafny.Int) _dafny.Int {
					return (_2194_i5).Minus((_this).F31())
				}
				_ = _init60
				var _element0_60 = _init60(_dafny.Zero)
				_ = _element0_60
				_nw339 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
				(_nw339).ArraySet1(_element0_60, 0)
				var _nativeLen0_60 = (_len0_60).Int()
				_ = _nativeLen0_60
				for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
					(_nw339).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
				}
			}
			_2193_v30 = _nw339
			var _2195_v31 _dafny.Map
			_ = _2195_v31
			_2195_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F30)
			var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
			_ = _index314
			(_2193_v30).ArraySet1((_dafny.IntOfInt64(952)).Plus((_2195_v31).Cardinality()), (_index314).Int())
			var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
			_ = _index315
			(_2193_v30).ArraySet1(p0, (_index315).Int())
			var _2196_v32 _dafny.Array
			_ = _2196_v32
			var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
			_ = _nw340
			_2196_v32 = _nw340
			var _2197_v33 _dafny.Map
			_ = _2197_v33
			_2197_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_2193_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))).Int()).(_dafny.Int))
			var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2196_v32), 0))
			_ = _index316
			(_2196_v32).ArraySet1((_2197_v33).Merge(_2197_v33), (_index316).Int())
			var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))
			_ = _index317
			(_2184_v21).ArraySet1(!(_this.F30), (_index317).Int())
			var _2198_v34 D0
			_ = _2198_v34
			_2198_v34 = Companion_D0_.Create_DC1_((_2186_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F26())).Cardinality()), _dafny.IntOfUint32((_2186_v23).Cardinality()))).Uint32()).(bool), true, p1, (_this).F26())
			var _2199_v35 D0
			_ = _2199_v35
			_2199_v35 = Companion_D0_.Create_DC2_(_2198_v34)
			var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2196_v32), 0))
			_ = _index318
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))
			_ = _index319
			var _rhs322 _dafny.Map = _2197_v33
			_ = _rhs322
			var _rhs323 bool = (_this).F27()
			_ = _rhs323
			var _rhs324 bool = ((_2193_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
				if (_2191_v28).Contains((_this).F26()) {
					return (_2191_v28).Multiplicity((_this).F26())
				}
				return (_2197_v33).Cardinality()
			})()) <= 0
			_ = _rhs324
			var _rhs325 _dafny.Int = (_this).F31()
			_ = _rhs325
			var _rhs326 bool = !((_dafny.MultiSetOf(p3, p3)).IsSubsetOf(Companion_Default___.Fm10((_this).F31(), p3, Companion_D0_.Create_DC2_(_2199_v35), globalState)))
			_ = _rhs326
			var _lhs292 _dafny.Array = _2196_v32
			_ = _lhs292
			var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2196_v32), 0))
			_ = _lhs293
			var _lhs294 _dafny.Array = _2184_v21
			_ = _lhs294
			var _lhs295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))
			_ = _lhs295
			var _lhs296 *GlobalState = globalState
			_ = _lhs296
			var _lhs297 *GlobalState = globalState
			_ = _lhs297
			(_lhs292).ArraySet1(_rhs322, (_lhs293).Int())
			r0 = _rhs323
			(_lhs294).ArraySet1(_rhs324, (_lhs295).Int())
			_lhs296.F12 = _rhs325
			_lhs297.F22 = _rhs326
			var _2200_v36 D3
			_ = _2200_v36
			_2200_v36 = Companion_D3_.Create_DC10_(_2193_v30)
			var _2201_v37 _dafny.Map
			_ = _2201_v37
			_2201_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2200_v36, p0)
			var _2202_v38 *C4
			_ = _2202_v38
			var _nw341 *C4 = New_C4_()
			_ = _nw341
			_nw341.Ctor__(_2201_v37, (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), (_this).F27())
			_2202_v38 = _nw341
			var _2203_v39 _dafny.MultiSet
			_ = _2203_v39
			_2203_v39 = _dafny.MultiSetOf(_2193_v30)
			var _rhs327 bool = (_2203_v39).IsSubsetOf(_2203_v39)
			_ = _rhs327
			var _rhs328 _dafny.Int = (_2157_v0).Select((Companion_Default___.SafeIndex((_2187_v24).Fm1(globalState), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs328
			var _rhs329 *C4 = (func() *C4 {
				if (_this).F27() {
					return _2202_v38
				}
				return _2202_v38
			})()
			_ = _rhs329
			var _rhs330 _dafny.Int = p1
			_ = _rhs330
			var _lhs298 *C9 = _this
			_ = _lhs298
			var _lhs299 *GlobalState = globalState
			_ = _lhs299
			var _lhs300 *GlobalState = globalState
			_ = _lhs300
			_lhs298.F30 = _rhs327
			_lhs299.F12 = _rhs328
			_2202_v38 = _rhs329
			_lhs300.F12 = _rhs330
			var _2204_v40 _dafny.Set
			_ = _2204_v40
			_2204_v40 = _dafny.SetOf((_2187_v24).Fm1(globalState))
			var _2205_v41 _dafny.Set
			_ = _2205_v41
			_2205_v41 = _dafny.SetOf(!((_this).F26()), (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool))
			var _2206_v42 _dafny.Map
			_ = _2206_v42
			_2206_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2205_v41)
			var _2207_v43 D21
			_ = _2207_v43
			_2207_v43 = Companion_D21_.Create_DC50_((_this).F26(), (_2204_v40).Cardinality(), (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), (_2206_v42).Cardinality())
			var _source31 D21 = _2207_v43
			_ = _source31
			if _source31.Is_DC50() {
				var _2208___mcc_h0 bool = _source31.Get_().(D21_DC50).Cf87
				_ = _2208___mcc_h0
				var _2209___mcc_h1 _dafny.Int = _source31.Get_().(D21_DC50).Cf88
				_ = _2209___mcc_h1
				var _2210___mcc_h2 bool = _source31.Get_().(D21_DC50).Cf89
				_ = _2210___mcc_h2
				var _2211___mcc_h3 _dafny.Int = _source31.Get_().(D21_DC50).Cf90
				_ = _2211___mcc_h3
				var _2212_cf90 _dafny.Int = _2211___mcc_h3
				_ = _2212_cf90
				var _2213_cf89 bool = _2210___mcc_h2
				_ = _2213_cf89
				var _2214_cf88 _dafny.Int = _2209___mcc_h1
				_ = _2214_cf88
				var _2215_cf87 bool = _2208___mcc_h0
				_ = _2215_cf87
				var _2216_v44 _dafny.Array
				_ = _2216_v44
				var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
				_ = _nw342
				_2216_v44 = _nw342
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_2216_v44), 0))
				_ = _index320
				(_2216_v44).ArraySet1(_2184_v21, (_index320).Int())
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))
				_ = _index321
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_2216_v44), 0))
				_ = _index322
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _index323
				var _rhs331 bool = _2213_cf89
				_ = _rhs331
				var _rhs332 _dafny.Array = _2192_v29
				_ = _rhs332
				var _rhs333 _dafny.Int = _2214_cf88
				_ = _rhs333
				var _rhs334 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p2, p2)
				_ = _rhs334
				var _rhs335 _dafny.Sequence = p2
				_ = _rhs335
				var _lhs301 _dafny.Array = _2184_v21
				_ = _lhs301
				var _lhs302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))
				_ = _lhs302
				var _lhs303 _dafny.Array = _2216_v44
				_ = _lhs303
				var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.ArrayLen((_2216_v44), 0))
				_ = _lhs304
				var _lhs305 _dafny.Array = _2193_v30
				_ = _lhs305
				var _lhs306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _lhs306
				var _lhs307 *GlobalState = globalState
				_ = _lhs307
				var _lhs308 *GlobalState = globalState
				_ = _lhs308
				(_lhs301).ArraySet1(_rhs331, (_lhs302).Int())
				(_lhs303).ArraySet1(_rhs332, (_lhs304).Int())
				(_lhs305).ArraySet1(_rhs333, (_lhs306).Int())
				_lhs307.F6 = _rhs334
				_lhs308.F6 = _rhs335
				var _2217_v45 _dafny.Array
				_ = _2217_v45
				var _nw343 _dafny.Array = _dafny.NewArray(_dafny.One)
				_ = _nw343
				_2217_v45 = _nw343
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_2217_v45), 0))
				_ = _index324
				(_2217_v45).ArraySet1(_2202_v38, (_index324).Int())
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _index325
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_2217_v45), 0))
				_ = _index326
				var _rhs336 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(p2, p2))
				_ = _rhs336
				var _rhs337 _dafny.Int = (_this).F31()
				_ = _rhs337
				var _rhs338 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("icsw")
				_ = _rhs338
				var _rhs339 *C4 = _2202_v38
				_ = _rhs339
				var _lhs309 _dafny.Array = _2193_v30
				_ = _lhs309
				var _lhs310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _lhs310
				var _lhs311 *GlobalState = globalState
				_ = _lhs311
				var _lhs312 _dafny.Array = _2217_v45
				_ = _lhs312
				var _lhs313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_2217_v45), 0))
				_ = _lhs313
				r0 = _rhs336
				(_lhs309).ArraySet1(_rhs337, (_lhs310).Int())
				_lhs311.F6 = _rhs338
				(_lhs312).ArraySet1(_rhs339, (_lhs313).Int())
				var _2218_v46 _dafny.Array
				_ = _2218_v46
				var _len0_61 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_61
				var _nw344 _dafny.Array
				_ = _nw344
				if _len0_61.Cmp(_dafny.Zero) == 0 {
					_nw344 = _dafny.NewArray(_len0_61)
				} else {
					var _init61 func(_dafny.Int) _dafny.Sequence = (func(_2219_p3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_2220_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_2219_p3)
						}
					})(p3)
					_ = _init61
					var _element0_61 = _init61(_dafny.Zero)
					_ = _element0_61
					_nw344 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
					(_nw344).ArraySet1(_element0_61, 0)
					var _nativeLen0_61 = (_len0_61).Int()
					_ = _nativeLen0_61
					for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
						(_nw344).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
					}
				}
				_2218_v46 = _nw344
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_2218_v46), 0))
				_ = _index327
				(_2218_v46).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(827))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}((func(_2221_p2 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_2222_i7 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_2221_p2).Cardinality()))
					}
				})(p2))), (_index327).Int())
				var _2223_v47 D12
				_ = _2223_v47
				_2223_v47 = Companion_D12_.Create_DC28_(p3)
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_2218_v46), 0))
				_ = _index328
				(_2218_v46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2157_v0, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0), (Companion_Default___.SafeIndex((_2157_v0).Select((Companion_Default___.SafeIndex((_2223_v47).Dtor_cf49(), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()))).Uint32(), p3)), (_index328).Int())
				var _2224_v48 _dafny.Map
				_ = _2224_v48
				_2224_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))
				var _2225_v50 D17
				_ = _2225_v50
				_2225_v50 = Companion_D17_.Create_DC41_(!((func() _dafny.Set {
					var _coll71 = _dafny.NewBuilder()
					_ = _coll71
					for _iter72 := _dafny.Iterate((_2204_v40).Elements()); ; {
						_compr_71, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _2226_v49 _dafny.Int
						_2226_v49 = interface{}(_compr_71).(_dafny.Int)
						if (_2204_v40).Contains(_2226_v49) {
							_coll71.Add((_2226_v49).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xxuvuil")).Cardinality())))
						}
					}
					return _coll71.ToSet()
				}()).IsProperSubsetOf(_dafny.SetOf((_2224_v48).Cardinality(), p3))), (_this).F26(), p0)
				_2225_v50 = _2225_v50
			} else if _source31.Is_DC51() {
				var _2227___mcc_h4 _dafny.Int = _source31.Get_().(D21_DC51).Cf91
				_ = _2227___mcc_h4
				var _2228___mcc_h5 T1 = _source31.Get_().(D21_DC51).Cf92
				_ = _2228___mcc_h5
				var _2229___mcc_h6 _dafny.Int = _source31.Get_().(D21_DC51).Cf93
				_ = _2229___mcc_h6
				var _2230___mcc_h7 _dafny.Int = _source31.Get_().(D21_DC51).Cf94
				_ = _2230___mcc_h7
				var _2231_cf94 _dafny.Int = _2230___mcc_h7
				_ = _2231_cf94
				var _2232_cf93 _dafny.Int = _2229___mcc_h6
				_ = _2232_cf93
				var _2233_cf92 T1 = _2228___mcc_h5
				_ = _2233_cf92
				var _2234_cf91 _dafny.Int = _2227___mcc_h4
				_ = _2234_cf91
				var _2235_v51 _dafny.Map
				_ = _2235_v51
				_2235_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), _dafny.IntOfInt64(-449))
				var _2236_v52 *C7
				_ = _2236_v52
				var _nw345 *C7 = New_C7_()
				_ = _nw345
				_nw345.Ctor__((_2233_cf92).F27(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2234_cf91, _2235_v51))
				_2236_v52 = _nw345
				var _2237_v53 _dafny.Map
				_ = _2237_v53
				_2237_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2236_v52, p3)
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _index329
				(_2193_v30).ArraySet1((Companion_D17_.Create_DC42_((_2237_v53).Cardinality(), p1, p3)).Dtor_cf79(), (_index329).Int())
				(globalState).F6 = (func() _dafny.Sequence {
					if _dafny.Companion_Sequence_.IsPrefixOf(p2, p2) {
						return (func() _dafny.Sequence {
							if !(_this.F30) {
								return p2
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("hsr")
						})()
					}
					return p2
				})()
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _index330
				(_2193_v30).ArraySet1((p1).Plus(_2231_cf94), (_index330).Int())
				var _2238_v54 *C1
				_ = _2238_v54
				var _nw346 *C1 = New_C1_()
				_ = _nw346
				_nw346.Ctor__(Companion_Default___.SafeModuloInt(p3, (_this).F31()), _2232_cf93, (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), (p0).Cmp(_2232_cf93) == 0)
				_2238_v54 = _nw346
				var _nw347 *C1 = New_C1_()
				_ = _nw347
				_nw347.Ctor__((_2238_v54.F37).Times(p3), p0, !((_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool)), (_2233_cf92).F26())
				_2238_v54 = _nw347
			} else if _source31.Is_DC52() {
				var _2239___mcc_h8 _dafny.Int = _source31.Get_().(D21_DC52).Cf95
				_ = _2239___mcc_h8
				var _2240___mcc_h9 _dafny.Sequence = _source31.Get_().(D21_DC52).Cf96
				_ = _2240___mcc_h9
				var _2241_cf96 _dafny.Sequence = _2240___mcc_h9
				_ = _2241_cf96
				var _2242_cf95 _dafny.Int = _2239___mcc_h8
				_ = _2242_cf95
				(globalState).F12 = p1
				_2157_v0 = _dafny.Companion_Sequence_.Concatenate(_2157_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2157_v0, _dafny.Companion_Sequence_.Update(_2157_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32(), p0)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2157_v0, _dafny.Companion_Sequence_.Update(_2157_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32(), p0))).Cardinality()))).Uint32(), (_this).F31()))
				var _2243_v55 _dafny.MultiSet
				_ = _2243_v55
				_2243_v55 = _dafny.MultiSetOf(_2157_v0)
				var _2244_v56 *C3
				_ = _2244_v56
				var _nw348 *C3 = New_C3_()
				_ = _nw348
				_nw348.Ctor__(_2243_v55, true, (_this).F27())
				_2244_v56 = _nw348
				var _2245_v57 _dafny.Array
				_ = _2245_v57
				var _nw349 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw349
				_2245_v57 = _nw349
				var _2246_v58 _dafny.Array
				_ = _2246_v58
				var _len0_62 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_62
				var _nw350 _dafny.Array
				_ = _nw350
				if _len0_62.Cmp(_dafny.Zero) == 0 {
					_nw350 = _dafny.NewArray(_len0_62)
				} else {
					var _init62 func(_dafny.Int) _dafny.MultiSet = (func(_2247_v28 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
						return func(_2248_i8 _dafny.Int) _dafny.MultiSet {
							return _2247_v28
						}
					})(_2191_v28)
					_ = _init62
					var _element0_62 = _init62(_dafny.Zero)
					_ = _element0_62
					_nw350 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
					(_nw350).ArraySet1(_element0_62, 0)
					var _nativeLen0_62 = (_len0_62).Int()
					_ = _nativeLen0_62
					for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
						(_nw350).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
					}
				}
				_2246_v58 = _nw350
				var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_2245_v57), 0))
				_ = _index331
				(_2245_v57).ArraySet1(_2246_v58, (_index331).Int())
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_2245_v57), 0))
				_ = _index332
				var _rhs340 _dafny.Int = (((_2157_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32()).(_dafny.Int)).Plus(p0)).Minus((_dafny.Zero).Minus((_2157_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-797), _dafny.IntOfUint32((_2157_v0).Cardinality()))).Uint32()).(_dafny.Int)))
				_ = _rhs340
				var _rhs341 _dafny.Array = _2246_v58
				_ = _rhs341
				var _rhs342 _dafny.Array = _2193_v30
				_ = _rhs342
				var _rhs343 _dafny.Int = ((_dafny.SetOf((_2191_v28).Cardinality())).Intersection(_dafny.SetOf(p1, (_this).F31(), p3))).Cardinality()
				_ = _rhs343
				var _lhs314 *GlobalState = globalState
				_ = _lhs314
				var _lhs315 _dafny.Array = _2245_v57
				_ = _lhs315
				var _lhs316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_2245_v57), 0))
				_ = _lhs316
				var _lhs317 *GlobalState = globalState
				_ = _lhs317
				_lhs314.F16 = _rhs340
				(_lhs315).ArraySet1(_rhs341, (_lhs316).Int())
				_2193_v30 = _rhs342
				_lhs317.F16 = _rhs343
			} else {
				var _2249___mcc_h10 _dafny.Map = _source31.Get_().(D21_DC49).Cf86
				_ = _2249___mcc_h10
				var _2250_cf86 _dafny.Map = _2249___mcc_h10
				_ = _2250_cf86
				var _2251_v59 *C4
				_ = _2251_v59
				var _nw351 *C4 = New_C4_()
				_ = _nw351
				_nw351.Ctor__(((_2202_v38).F33()).Merge((_2202_v38).F33()), !_dafny.Companion_Sequence_.Equal(p2, p2), _this.F30)
				_2251_v59 = _nw351
				var _2252_v60 _dafny.Map
				_ = _2252_v60
				_2252_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool))
				var _2253_v61 *C4
				_ = _2253_v61
				var _nw352 *C4 = New_C4_()
				_ = _nw352
				_nw352.Ctor__((func() _dafny.Map {
					if _this.F30 {
						return (_2202_v38).F33()
					}
					return (_2251_v59).F33()
				})(), (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), (func() bool {
					if (_2252_v60).Contains(!((_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool))) {
						return (_2252_v60).Get(!((_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool))).(bool)
					}
					return (_this).F26()
				})())
				_2253_v61 = _nw352
				var _2254_v62 T1
				_ = _2254_v62
				var _nw353 *C4 = New_C4_()
				_ = _nw353
				_nw353.Ctor__((_2251_v59).F33(), (_this).F27(), (_this).F26())
				_2254_v62 = _nw353
				var _2255_v63 D21
				_ = _2255_v63
				_2255_v63 = Companion_D21_.Create_DC51_((_this).F31(), _2254_v62, _dafny.IntOfInt64(875), (_this).F31())
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(247), _dafny.ArrayLen((_2193_v30), 0))
				_ = _index333
				(_2193_v30).ArraySet1((_2255_v63).Dtor_cf91(), (_index333).Int())
				var _2256_v64 _dafny.CodePoint
				_ = _2256_v64
				_2256_v64 = _dafny.CodePoint('r')
				var _2257_v65 _dafny.Set
				_ = _2257_v65
				_2257_v65 = _dafny.SetOf(_2256_v64, _2256_v64)
				var _2258_v66 *C6
				_ = _2258_v66
				var _nw354 *C6 = New_C6_()
				_ = _nw354
				_nw354.Ctor__((_2257_v65).Cardinality())
				_2258_v66 = _nw354
				var _2259_v67 _dafny.MultiSet
				_ = _2259_v67
				_2259_v67 = _dafny.MultiSetOf(_2258_v66)
				var _2260_v68 _dafny.Map
				_ = _2260_v68
				_2260_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2259_v67, p3)
				var _2261_v69 D16
				_ = _2261_v69
				_2261_v69 = Companion_D16_.Create_DC37_(_2195_v31)
				var _2262_v71 _dafny.Map
				_ = _2262_v71
				_2262_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _2256_v64)
				var _2263_v72 D14
				_ = _2263_v72
				_2263_v72 = Companion_D14_.Create_DC32_(_dafny.MultiSetOf((_2254_v62).F27()), (_this).F26(), (_this).F27(), (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool), ((_2262_v71).Update((_2254_v62).F27(), _2256_v64)).Cardinality())
				var _2264_v73 _dafny.Map
				_ = _2264_v73
				_2264_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2263_v72, (_2254_v62).Fm2(_this.F30, p1, globalState))
				var _2265_v74 _dafny.Map
				_ = _2265_v74
				_2265_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _2264_v73)
				var _2266_v75 _dafny.Array
				_ = _2266_v75
				var _nwElement0_85 _dafny.Map = (_2195_v31).Update((_2260_v68).Cardinality(), (_this).F26())
				_ = _nwElement0_85
				var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(6))
				_ = _nw355
				(_nw355).ArraySet1(_nwElement0_85, 0)
				(_nw355).ArraySet1(_2195_v31, 1)
				(_nw355).ArraySet1(_2195_v31, 2)
				(_nw355).ArraySet1((_2195_v31).Merge((_2261_v69).Dtor_cf68()), 3)
				(_nw355).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2251_v59).Fm1(globalState), false), 4)
				(_nw355).ArraySet1(func() _dafny.Map {
					var _coll72 = _dafny.NewMapBuilder()
					_ = _coll72
					for _iter73 := _dafny.Iterate((_2265_v74).Keys().Elements()); ; {
						_compr_72, _ok73 := _iter73()
						if !_ok73 {
							break
						}
						var _2267_v70 _dafny.Int
						_2267_v70 = interface{}(_compr_72).(_dafny.Int)
						if (_2265_v74).Contains(_2267_v70) {
							_coll72.Add((_2267_v70).Plus(p1), (_2254_v62).F26())
						}
					}
					return _coll72.ToMap()
				}(), 5)
				_2266_v75 = _nw355
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_2266_v75), 0))
				_ = _index334
				(_2266_v75).ArraySet1(_2195_v31, (_index334).Int())
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_2266_v75), 0))
				_ = _index335
				(_2266_v75).ArraySet1((func() _dafny.Map {
					if (_2184_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_2184_v21), 0))).Int()).(bool) {
						return _2195_v31
					}
					return (_2195_v31).Merge(_2195_v31)
				})(), (_index335).Int())
			}
		} else {
			var _2268_v76 _dafny.Set
			_ = _2268_v76
			_2268_v76 = _dafny.SetOf(_this.F30)
			var _2269_v77 D6
			_ = _2269_v77
			_2269_v77 = Companion_D6_.Create_DC17_((_2268_v76).Cardinality(), p0, p3, p1, (_dafny.Zero).Minus(p0))
			var _2270_v78 _dafny.Map
			_ = _2270_v78
			_2270_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F26())
			var _2271_v79 D16
			_ = _2271_v79
			_2271_v79 = Companion_D16_.Create_DC38_((_2270_v78).Cardinality(), p1, _this.F30)
			var _2272_v80 _dafny.Map
			_ = _2272_v80
			_2272_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, (_dafny.SetOf(_2271_v79)).Cardinality())
			var _2273_v81 D0
			_ = _2273_v81
			_2273_v81 = Companion_D0_.Create_DC0_(_this.F30)
			var _2274_v82 _dafny.Map
			_ = _2274_v82
			_2274_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(_2273_v81, p0, (_dafny.Zero).Minus(p3), globalState), _dafny.UnicodeSeqOfUtf8Bytes("k"))
			var _2275_v83 _dafny.MultiSet
			_ = _2275_v83
			_2275_v83 = _dafny.MultiSetOf((_this).F27())
			var _2276_v84 _dafny.Set
			_ = _2276_v84
			_2276_v84 = _dafny.SetOf((_this).F31(), (func() _dafny.Int {
				if (_2275_v83).Contains((_this).Fm2((func(_pat_let62_0 D7) D7 {
					return func(_2277_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let63_0 _dafny.Int) D7 {
							return func(_2278_dt__update_hcf36_h0 _dafny.Int) D7 {
								return Companion_D7_.Create_DC19_(_2278_dt__update_hcf36_h0, (_2277_dt__update__tmp_h0).Dtor_cf37(), (_2277_dt__update__tmp_h0).Dtor_cf38(), (_2277_dt__update__tmp_h0).Dtor_cf39())
							}(_pat_let63_0)
						}((_this).F31())
					}(_pat_let62_0)
				}(Companion_D7_.Create_DC19_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("smafhihgt")).Cardinality()), _2275_v83, (_this).F26(), _2272_v80))).Dtor_cf38(), p0, globalState)) {
					return (_2275_v83).Multiplicity((_this).Fm2((func(_pat_let64_0 D7) D7 {
						return func(_2279_dt__update__tmp_h1 D7) D7 {
							return func(_pat_let65_0 _dafny.Int) D7 {
								return func(_2280_dt__update_hcf36_h1 _dafny.Int) D7 {
									return Companion_D7_.Create_DC19_(_2280_dt__update_hcf36_h1, (_2279_dt__update__tmp_h1).Dtor_cf37(), (_2279_dt__update__tmp_h1).Dtor_cf38(), (_2279_dt__update__tmp_h1).Dtor_cf39())
								}(_pat_let65_0)
							}((_this).F31())
						}(_pat_let64_0)
					}(Companion_D7_.Create_DC19_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("smafhihgt")).Cardinality()), _2275_v83, (_this).F26(), _2272_v80))).Dtor_cf38(), p0, globalState))
				}
				return p0
			})(), p0)
			var _2281_v85 _dafny.Map
			_ = _2281_v85
			_2281_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F31())
			_2269_v77 = Companion_Default___.Fm53((func() _dafny.Int {
				if (_this).F27() {
					return (_dafny.MultiSetOf(p3, p3, (_2272_v80).Cardinality())).Cardinality()
				}
				return _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2274_v82).Contains(p0) {
						return (_2274_v82).Get(p0).(_dafny.Sequence)
					}
					return p2
				})()).Cardinality())
			})(), p1, _2276_v84, _2281_v85, globalState)
			(globalState).F24 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(p2, p2), _dafny.Companion_Sequence_.Concatenate(p2, p2))
			(globalState).F15 = (_this).F27()
			(globalState).F16 = (p3).Plus(p1)
			var _2282_v86 _dafny.Sequence
			_ = _2282_v86
			_2282_v86 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(704))).Uint32(), func(coer115 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg115 _dafny.Int) interface{} {
					return coer115(arg115)
				}
			}((func(_2283_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2284_i9 _dafny.Int) _dafny.Int {
					return _2283_p1
				}
			})(p1))))
			var _2285_v87 *C2
			_ = _2285_v87
			var _nw356 *C2 = New_C2_()
			_ = _nw356
			_nw356.Ctor__(_dafny.Companion_Sequence_.Update(_2282_v86, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F31()), _dafny.IntOfUint32((_2282_v86).Cardinality()))).Uint32(), _dafny.SeqOf(p0)), (_this).F27(), (_this).F27())
			_2285_v87 = _nw356
		}
		var _2286_v88 _dafny.Array
		_ = _2286_v88
		var _nw357 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
		_ = _nw357
		_2286_v88 = _nw357
		for _iter74 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2286_v88), 0))); ; {
			_guard_loop_1, _ok74 := _iter74()
			if !_ok74 {
				break
			}
			var _2287_i10 _dafny.Int
			_2287_i10 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_2287_i10).Sign() != -1) && ((_2287_i10).Cmp(_dafny.ArrayLen((_2286_v88), 0)) < 0)) {
				(_2286_v88).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(680))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}(func(_2288_i11 _dafny.Int) _dafny.CodePoint {
					return (func() _dafny.CodePoint {
						if _this.F30 {
							return _dafny.CodePoint('u')
						}
						return _dafny.CodePoint('e')
					})()
				})), (_2287_i10).Int())
			}
		}
		var _2289_v89 D0
		_ = _2289_v89
		_2289_v89 = Companion_D0_.Create_DC0_(_this.F30)
		var _2290_v90 _dafny.Sequence
		_ = _2290_v90
		_2290_v90 = _dafny.SeqOf(!(!(false)))
		var _2291_v91 _dafny.Map
		_ = _2291_v91
		_2291_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, (_this).F31())
		var _2292_v92 _dafny.Array
		_ = _2292_v92
		var _nwElement0_86 bool = true
		_ = _nwElement0_86
		var _nw358 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(26))
		_ = _nw358
		(_nw358).ArraySet1(_nwElement0_86, 0)
		(_nw358).ArraySet1((_this).Fm2((_this).F26(), (_this).F31(), globalState), 1)
		(_nw358).ArraySet1(_this.F30, 2)
		(_nw358).ArraySet1((_dafny.IntOfInt64(-543)).Cmp(p1) > 0, 3)
		(_nw358).ArraySet1(!((_this).F27()) || (_this.F30), 4)
		(_nw358).ArraySet1((_2289_v89).Dtor_cf0(), 5)
		(_nw358).ArraySet1((_2290_v90).Select((Companion_Default___.SafeIndex((_2291_v91).Cardinality(), _dafny.IntOfUint32((_2290_v90).Cardinality()))).Uint32()).(bool), 6)
		(_nw358).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_2157_v0, _2157_v0), 7)
		(_nw358).ArraySet1(((_this).Fm1(globalState)).Cmp((_dafny.Zero).Minus(p1)) <= 0, 8)
		(_nw358).ArraySet1((_this).F27(), 9)
		(_nw358).ArraySet1((_this).F27(), 10)
		(_nw358).ArraySet1((_this).F26(), 11)
		(_nw358).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("o"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg117 _dafny.Int) interface{} {
				return coer117(arg117)
			}
		}(func(_2293_i13 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))), 12)
		(_nw358).ArraySet1(_this.F30, 13)
		(_nw358).ArraySet1(true, 14)
		(_nw358).ArraySet1((_this).F27(), 15)
		(_nw358).ArraySet1((_this).F27(), 16)
		(_nw358).ArraySet1((_this).F27(), 17)
		(_nw358).ArraySet1((_this).F26(), 18)
		(_nw358).ArraySet1(_this.F30, 19)
		(_nw358).ArraySet1((_this).F26(), 20)
		(_nw358).ArraySet1(!((_this).F27()) || (_this.F30), 21)
		(_nw358).ArraySet1(_this.F30, 22)
		(_nw358).ArraySet1(_this.F30, 23)
		(_nw358).ArraySet1(_this.F30, 24)
		(_nw358).ArraySet1(false, 25)
		_2292_v92 = _nw358
		for _iter75 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2292_v92), 0))); ; {
			_guard_loop_2, _ok75 := _iter75()
			if !_ok75 {
				break
			}
			var _2294_i12 _dafny.Int
			_2294_i12 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_2294_i12).Sign() != -1) && ((_2294_i12).Cmp(_dafny.ArrayLen((_2292_v92), 0)) < 0)) {
				(_2292_v92).ArraySet1(!((func() bool {
					if (_this).F27() {
						return (_this).F26()
					}
					return false
				})()), (_2294_i12).Int())
			}
		}
		r0 = ((_this).Fm2((_this).F27(), p0, globalState)) == ((_dafny.MultiSetFromSeq(_2157_v0)).Equals((_this).F29()))
		var _2295_v93 *C0
		_ = _2295_v93
		var _nw359 *C0 = New_C0_()
		_ = _nw359
		_nw359.Ctor__((_this).F27())
		_2295_v93 = _nw359
		var _2296_v94 _dafny.Map
		_ = _2296_v94
		_2296_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, _2295_v93)
		r1 = ((_2296_v94).Cardinality()).Cmp((_this).F31()) >= 0
		var _2297_v95 _dafny.Array
		_ = _2297_v95
		var _nwElement0_87 _dafny.Int = p3
		_ = _nwElement0_87
		var _nw360 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(5))
		_ = _nw360
		(_nw360).ArraySet1(_nwElement0_87, 0)
		(_nw360).ArraySet1((_this).F31(), 1)
		(_nw360).ArraySet1((_this).F31(), 2)
		(_nw360).ArraySet1((func() _dafny.Int {
			if _this.F30 {
				return (_this).F31()
			}
			return (_this).F31()
		})(), 3)
		(_nw360).ArraySet1((_this).F31(), 4)
		_2297_v95 = _nw360
		r2 = _2297_v95
		var _2298_v96 T0
		_ = _2298_v96
		var _nw361 *C6 = New_C6_()
		_ = _nw361
		_nw361.Ctor__(p0)
		_2298_v96 = _nw361
		r3 = _2298_v96
		return r0, r1, r2, r3
	}
}
func (_this *C9) M4(globalState *GlobalState) (bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var _2299_v0 D6
		_ = _2299_v0
		_2299_v0 = Companion_D6_.Create_DC16_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-172))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg118 _dafny.Int) interface{} {
				return coer118(arg118)
			}
		}(func(_2300_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))
		var _2301_i0 _dafny.Int
		_ = _2301_i0
		_2301_i0 = _dafny.Zero
		{
			for func(_source32 D6) bool {
				if _source32.Is_DC17() {
					var _2313___mcc_h0 _dafny.Int = _source32.Get_().(D6_DC17).Cf30
					_ = _2313___mcc_h0
					var _2314___mcc_h1 _dafny.Int = _source32.Get_().(D6_DC17).Cf31
					_ = _2314___mcc_h1
					var _2315___mcc_h2 _dafny.Int = _source32.Get_().(D6_DC17).Cf32
					_ = _2315___mcc_h2
					var _2316___mcc_h3 _dafny.Int = _source32.Get_().(D6_DC17).Cf33
					_ = _2316___mcc_h3
					var _2317___mcc_h4 _dafny.Int = _source32.Get_().(D6_DC17).Cf34
					_ = _2317___mcc_h4
					var _2318_cf34 _dafny.Int = _2317___mcc_h4
					_ = _2318_cf34
					var _2319_cf33 _dafny.Int = _2316___mcc_h3
					_ = _2319_cf33
					var _2320_cf32 _dafny.Int = _2315___mcc_h2
					_ = _2320_cf32
					var _2321_cf31 _dafny.Int = _2314___mcc_h1
					_ = _2321_cf31
					var _2322_cf30 _dafny.Int = _2313___mcc_h0
					_ = _2322_cf30
					return true
				} else {
					var _2323___mcc_h5 _dafny.Sequence = _source32.Get_().(D6_DC16).Cf29
					_ = _2323___mcc_h5
					var _2324_cf29 _dafny.Sequence = _2323___mcc_h5
					_ = _2324_cf29
					return !(true)
				}
			}(_2299_v0) {
				{
					if (_2301_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_2301_i0 = (_2301_i0).Plus(_dafny.One)
					var _2302_v1 _dafny.Map
					_ = _2302_v1
					_2302_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), (_this).F27())
					var _2303_v2 _dafny.Sequence
					_ = _2303_v2
					_2303_v2 = _dafny.SeqOf((_2302_v1).Merge(_2302_v1), _2302_v1)
					_2303_v2 = _2303_v2
					var _2304_v3 _dafny.Set
					_ = _2304_v3
					_2304_v3 = _dafny.SetOf((_this).F31())
					var _2305_v4 _dafny.Sequence
					_ = _2305_v4
					_2305_v4 = _dafny.SeqOf(_dafny.IntOfInt64(-469), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2304_v3).Cardinality(), (_this).F26())).Cardinality(), (_this).F31())
					var _2306_v5 _dafny.Sequence
					_ = _2306_v5
					_2306_v5 = _dafny.SeqOf(_2305_v4, _2305_v4, _2305_v4)
					var _2307_v6 *C2
					_ = _2307_v6
					var _nw362 *C2 = New_C2_()
					_ = _nw362
					_nw362.Ctor__(_2306_v5, _this.F30, (_this).F26())
					_2307_v6 = _nw362
					(globalState).F2 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_2305_v4, _2305_v4))).IsSubsetOf(((_this).F29()).Union((_this).F29()))
					var _2308_v7 _dafny.Sequence
					_ = _2308_v7
					_2308_v7 = _dafny.SeqOf(_this.F30)
					var _2309_v8 _dafny.Map
					_ = _2309_v8
					_2309_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_2307_v6).Fm1(globalState))
					var _2310_v9 _dafny.Map
					_ = _2310_v9
					_2310_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F27(), (_2309_v8).Cardinality())
					var _2311_v10 _dafny.Map
					_ = _2311_v10
					_2311_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2308_v7).Cardinality()), _2310_v9)
					var _2312_v11 *C7
					_ = _2312_v11
					var _nw363 *C7 = New_C7_()
					_ = _nw363
					_nw363.Ctor__(_this.F30, _2311_v10)
					_2312_v11 = _nw363
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _2325_v12 *C6
		_ = _2325_v12
		var _nw364 *C6 = New_C6_()
		_ = _nw364
		_nw364.Ctor__((_this).F31())
		_2325_v12 = _nw364
		var _2326_v13 _dafny.Set
		_ = _2326_v13
		_2326_v13 = _dafny.SetOf(_this.F30)
		_2326_v13 = _2326_v13
		if (_this).F27() {
			var _2327_v14 D16
			_ = _2327_v14
			_2327_v14 = Companion_D16_.Create_DC39_(Companion_Default___.Fm54(true, (_this).F31(), _2325_v12.F40, globalState))
			var _source33 D16 = _2327_v14
			_ = _source33
			if _source33.Is_DC38() {
				var _2328___mcc_h6 _dafny.Int = _source33.Get_().(D16_DC38).Cf69
				_ = _2328___mcc_h6
				var _2329___mcc_h7 _dafny.Int = _source33.Get_().(D16_DC38).Cf70
				_ = _2329___mcc_h7
				var _2330___mcc_h8 bool = _source33.Get_().(D16_DC38).Cf71
				_ = _2330___mcc_h8
				var _2331_cf71 bool = _2330___mcc_h8
				_ = _2331_cf71
				var _2332_cf70 _dafny.Int = _2329___mcc_h7
				_ = _2332_cf70
				var _2333_cf69 _dafny.Int = _2328___mcc_h6
				_ = _2333_cf69
				var _2334_v15 _dafny.Sequence
				_ = _2334_v15
				_2334_v15 = _dafny.UnicodeSeqOfUtf8Bytes("nq")
				var _2335_v16 D3
				_ = _2335_v16
				_2335_v16 = Companion_D3_.Create_DC11_(Companion_Default___.Fm24(_2334_v15, (_this).F26(), _2332_cf70, globalState))
				var _2336_v17 _dafny.Array
				_ = _2336_v17
				var _len0_63 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_63
				var _nw365 _dafny.Array
				_ = _nw365
				if _len0_63.Cmp(_dafny.Zero) == 0 {
					_nw365 = _dafny.NewArray(_len0_63)
				} else {
					var _init63 func(_dafny.Int) bool = func(_2337_i2 _dafny.Int) bool {
						return !(false)
					}
					_ = _init63
					var _element0_63 = _init63(_dafny.Zero)
					_ = _element0_63
					_nw365 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
					(_nw365).ArraySet1(_element0_63, 0)
					var _nativeLen0_63 = (_len0_63).Int()
					_ = _nativeLen0_63
					for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
						(_nw365).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
					}
				}
				_2336_v17 = _nw365
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_2336_v17), 0))
				_ = _index336
				(_2336_v17).ArraySet1(_this.F30, (_index336).Int())
				var _2338_v18 _dafny.CodePoint
				_ = _2338_v18
				_2338_v18 = _dafny.CodePoint('f')
				var _2339_v19 _dafny.Map
				_ = _2339_v19
				_2339_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2333_cf69, true)
				var _2340_v20 _dafny.Array
				_ = _2340_v20
				var _nwElement0_88 _dafny.Int = _2325_v12.F40
				_ = _nwElement0_88
				var _nw366 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(19))
				_ = _nw366
				(_nw366).ArraySet1(_nwElement0_88, 0)
				(_nw366).ArraySet1((_dafny.Zero).Minus(_2325_v12.F40), 1)
				(_nw366).ArraySet1(_2325_v12.F40, 2)
				(_nw366).ArraySet1((_2339_v19).Cardinality(), 3)
				(_nw366).ArraySet1((_this).F31(), 4)
				(_nw366).ArraySet1(_2325_v12.F40, 5)
				(_nw366).ArraySet1(_dafny.IntOfInt64(-305), 6)
				(_nw366).ArraySet1(_dafny.IntOfInt64(635), 7)
				(_nw366).ArraySet1(_2333_cf69, 8)
				(_nw366).ArraySet1(_2332_cf70, 9)
				(_nw366).ArraySet1(_2325_v12.F40, 10)
				(_nw366).ArraySet1(_2332_cf70, 11)
				(_nw366).ArraySet1(_2332_cf70, 12)
				(_nw366).ArraySet1(_2325_v12.F40, 13)
				(_nw366).ArraySet1(_2333_cf69, 14)
				(_nw366).ArraySet1(_2325_v12.F40, 15)
				(_nw366).ArraySet1((_this).F31(), 16)
				(_nw366).ArraySet1(_dafny.IntOfUint32((_2334_v15).Cardinality()), 17)
				(_nw366).ArraySet1((_this).F31(), 18)
				_2340_v20 = _nw366
				var _2341_v21 _dafny.MultiSet
				_ = _2341_v21
				_2341_v21 = _dafny.MultiSetOf(_2340_v20, _2340_v20)
				var _2342_v22 _dafny.Set
				_ = _2342_v22
				_2342_v22 = _dafny.SetOf(_2333_cf69, (func() _dafny.Int {
					if (_2341_v21).Contains(_2340_v20) {
						return (_2341_v21).Multiplicity(_2340_v20)
					}
					return _2325_v12.F40
				})(), _2325_v12.F40)
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_2336_v17), 0))
				_ = _index337
				var _rhs344 D3 = Companion_D3_.Create_DC11_(_2338_v18)
				_ = _rhs344
				var _rhs345 bool = (_2326_v13).IsProperSubsetOf((_2326_v13).Difference(_dafny.SetOf(_this.F30, _this.F30, (_this).F26())))
				_ = _rhs345
				var _rhs346 _dafny.Int = _2325_v12.F40
				_ = _rhs346
				var _rhs347 _dafny.Int = (_2342_v22).Cardinality()
				_ = _rhs347
				var _lhs318 _dafny.Array = _2336_v17
				_ = _lhs318
				var _lhs319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_2336_v17), 0))
				_ = _lhs319
				var _lhs320 *GlobalState = globalState
				_ = _lhs320
				_2335_v16 = _rhs344
				(_lhs318).ArraySet1(_rhs345, (_lhs319).Int())
				_2333_cf69 = _rhs346
				_lhs320.F16 = _rhs347
				(_2325_v12).F40 = _2325_v12.F40
				(globalState).F16 = _2333_cf69
				var _2343_v23 _dafny.Sequence
				_ = _2343_v23
				_2343_v23 = _dafny.SeqOf((_this).F31())
				(globalState).F15 = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_2343_v23).Cardinality())), _2333_cf69)
			} else if _source33.Is_DC37() {
				var _2344___mcc_h9 _dafny.Map = _source33.Get_().(D16_DC37).Cf68
				_ = _2344___mcc_h9
				var _2345_cf68 _dafny.Map = _2344___mcc_h9
				_ = _2345_cf68
				var _2346_v24 _dafny.Array
				_ = _2346_v24
				var _len0_64 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_64
				var _nw367 _dafny.Array
				_ = _nw367
				if _len0_64.Cmp(_dafny.Zero) == 0 {
					_nw367 = _dafny.NewArray(_len0_64)
				} else {
					var _init64 func(_dafny.Int) bool = func(_2347_i3 _dafny.Int) bool {
						return (_this).F26()
					}
					_ = _init64
					var _element0_64 = _init64(_dafny.Zero)
					_ = _element0_64
					_nw367 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
					(_nw367).ArraySet1(_element0_64, 0)
					var _nativeLen0_64 = (_len0_64).Int()
					_ = _nativeLen0_64
					for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
						(_nw367).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
					}
				}
				_2346_v24 = _nw367
				var _2348_v25 _dafny.Set
				_ = _2348_v25
				_2348_v25 = _dafny.SetOf((_this).F31())
				var _2349_v26 D12
				_ = _2349_v26
				_2349_v26 = Companion_D12_.Create_DC28_(_2325_v12.F40)
				var _2350_v27 _dafny.Array
				_ = _2350_v27
				var _nwElement0_89 _dafny.Int = (_2348_v25).Cardinality()
				_ = _nwElement0_89
				var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.IntOfInt64(2))
				_ = _nw368
				(_nw368).ArraySet1(_nwElement0_89, 0)
				(_nw368).ArraySet1((_2349_v26).Dtor_cf49(), 1)
				_2350_v27 = _nw368
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2346_v24), 0))
				_ = _index338
				(_2346_v24).ArraySet1(!((_dafny.IntOfUint32((_dafny.SeqOf(_2350_v27, _2350_v27)).Cardinality())).Cmp((_this).F31()) <= 0), (_index338).Int())
				var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2346_v24), 0))
				_ = _index339
				(_2346_v24).ArraySet1((_this).F27(), (_index339).Int())
				(globalState).F2 = !((_2325_v12).Fm44(globalState))
				var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_2350_v27), 0))
				_ = _index340
				(_2350_v27).ArraySet1(Companion_Default___.SafeModuloInt((_this).F31(), _dafny.IntOfInt64(-769)), (_index340).Int())
				var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_2350_v27), 0))
				_ = _index341
				(_2350_v27).ArraySet1(_2325_v12.F40, (_index341).Int())
				(globalState).F2 = ((_2346_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.One, _dafny.ArrayLen((_2346_v24), 0))).Int()).(bool)) == (!(_this.F30))
			} else {
				var _2351___mcc_h10 D16 = _source33.Get_().(D16_DC39).Cf72
				_ = _2351___mcc_h10
				var _2352_cf72 D16 = _2351___mcc_h10
				_ = _2352_cf72
				var _2353_v28 _dafny.Sequence
				_ = _2353_v28
				_2353_v28 = _dafny.UnicodeSeqOfUtf8Bytes("veta")
				var _2354_v29 _dafny.Map
				_ = _2354_v29
				_2354_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _dafny.IntOfUint32((_2353_v28).Cardinality()))
				var _2355_v30 _dafny.CodePoint
				_ = _2355_v30
				_2355_v30 = _dafny.CodePoint('t')
				var _2356_v31 _dafny.MultiSet
				_ = _2356_v31
				_2356_v31 = _dafny.MultiSetOf(_dafny.CodePoint('q'), _2355_v30)
				_2354_v29 = (_2354_v29).Update((_2356_v31).IsSubsetOf(_dafny.MultiSetOf(_2355_v30, _2355_v30, _2355_v30)), Companion_Default___.SafeModuloInt(_2325_v12.F40, (_dafny.Zero).Minus(_2325_v12.F40)))
				r1 = !((_2325_v12).Fm44(globalState))
				var _2357_v32 _dafny.Array
				_ = _2357_v32
				var _len0_65 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_65
				var _nw369 _dafny.Array
				_ = _nw369
				if _len0_65.Cmp(_dafny.Zero) == 0 {
					_nw369 = _dafny.NewArray(_len0_65)
				} else {
					var _init65 func(_dafny.Int) _dafny.Int = (func(_2358_v12 *C6) func(_dafny.Int) _dafny.Int {
						return func(_2359_i4 _dafny.Int) _dafny.Int {
							return (_2359_i4).Plus(_2358_v12.F40)
						}
					})(_2325_v12)
					_ = _init65
					var _element0_65 = _init65(_dafny.Zero)
					_ = _element0_65
					_nw369 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
					(_nw369).ArraySet1(_element0_65, 0)
					var _nativeLen0_65 = (_len0_65).Int()
					_ = _nativeLen0_65
					for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
						(_nw369).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
					}
				}
				_2357_v32 = _nw369
				var _2360_v33 _dafny.Sequence
				_ = _2360_v33
				_2360_v33 = _dafny.SeqOf(true, (_this).F26(), _this.F30, (_this).F26())
				var _2361_v34 _dafny.Sequence
				_ = _2361_v34
				_2361_v34 = _dafny.SeqOf(_dafny.IntOfUint32((_2360_v33).Cardinality()), _2325_v12.F40)
				var _2362_v35 _dafny.Set
				_ = _2362_v35
				_2362_v35 = _dafny.SetOf(_2325_v12.F40, (_this).Fm1(globalState))
				var _nwElement0_90 _dafny.Int = (_dafny.Zero).Minus((_2361_v34).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm39(globalState), _dafny.IntOfUint32((_2361_v34).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _nwElement0_90
				var _nw370 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(4))
				_ = _nw370
				(_nw370).ArraySet1(_nwElement0_90, 0)
				(_nw370).ArraySet1(((_dafny.Zero).Minus((_2362_v35).Cardinality())).Minus((_this).F31()), 1)
				(_nw370).ArraySet1(_dafny.IntOfUint32((_2353_v28).Cardinality()), 2)
				(_nw370).ArraySet1((_this).F31(), 3)
				_2357_v32 = _nw370
				var _2363_v36 _dafny.Array
				_ = _2363_v36
				var _nw371 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw371
				_2363_v36 = _nw371
				var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_2363_v36), 0))
				_ = _index342
				(_2363_v36).ArraySet1(_this.F30, (_index342).Int())
				var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_2363_v36), 0))
				_ = _index343
				(_2363_v36).ArraySet1((func() bool {
					if (_this).F26() {
						return !((_this).F27())
					}
					return (_this).F27()
				})(), (_index343).Int())
			}
			var _2364_v37 _dafny.Sequence
			_ = _2364_v37
			_2364_v37 = _dafny.SeqOf(_2325_v12.F40, Companion_Default___.SafeModuloInt(_2325_v12.F40, ((_this).F29()).Cardinality()))
			var _2365_v38 _dafny.Sequence
			_ = _2365_v38
			_2365_v38 = _dafny.UnicodeSeqOfUtf8Bytes("ivqetbs")
			(globalState).F12 = (_2364_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2365_v38).Cardinality()), _dafny.IntOfUint32((_2364_v37).Cardinality()))).Uint32()).(_dafny.Int)
			var _2366_v39 _dafny.Array
			_ = _2366_v39
			var _len0_66 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_66
			var _nw372 _dafny.Array
			_ = _nw372
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw372 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) bool = (func(_2367_v38 _dafny.Sequence) func(_dafny.Int) bool {
					return func(_2368_i5 _dafny.Int) bool {
						return (_dafny.IntOfUint32((_2367_v38).Cardinality())).Cmp((_this).F31()) <= 0
					}
				})(_2365_v38)
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw372 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw372).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw372).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2366_v39 = _nw372
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2366_v39), 0))
			_ = _index344
			(_2366_v39).ArraySet1((_2325_v12.F40).Cmp(_dafny.IntOfInt64(788)) <= 0, (_index344).Int())
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(191), _dafny.ArrayLen((_2366_v39), 0))
			_ = _index345
			(_2366_v39).ArraySet1((func() bool {
				if ((_this).F31()).Cmp(((_this).F29()).Cardinality()) > 0 {
					return _this.F30
				}
				return (_this).F26()
			})(), (_index345).Int())
			var _2369_v40 _dafny.Map
			_ = _2369_v40
			_2369_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2325_v12.F40, !(_this.F30))
			_2369_v40 = _2369_v40
			(globalState).F15 = (_2325_v12).Fm44(globalState)
		} else {
			var _2370_v41 *C5
			_ = _2370_v41
			var _nw373 *C5 = New_C5_()
			_ = _nw373
			_nw373.Ctor__((_this).F27(), !((_this).F26()))
			_2370_v41 = _nw373
			_2370_v41 = _2370_v41
			(globalState).F16 = Companion_Default___.SafeDivisionInt(_2325_v12.F40, (_2325_v12.F40).Plus(_dafny.IntOfInt64(844)))
			var _2371_v42 D7
			_ = _2371_v42
			_2371_v42 = Companion_D7_.Create_DC19_(_2325_v12.F40, _dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F30, (_this).F26())), (Companion_D0_.Create_DC1_((_this).F27(), (_this).F26(), _dafny.IntOfInt64(344), _this.F30)).Dtor_cf4(), Companion_Default___.Fm30(_2325_v12.F40, (_this).F27(), globalState))
			var _2372_v43 _dafny.Map
			_ = _2372_v43
			_2372_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2325_v12).Fm44(globalState), (_this).F31())
			var _pat_let_tv59 = _2372_v43
			_ = _pat_let_tv59
			var _2373_v44 _dafny.Map
			_ = _2373_v44
			_2373_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let66_0 D7) D7 {
				return func(_2374_dt__update__tmp_h0 D7) D7 {
					return func(_pat_let67_0 _dafny.Map) D7 {
						return func(_2375_dt__update_hcf39_h0 _dafny.Map) D7 {
							return Companion_D7_.Create_DC19_((_2374_dt__update__tmp_h0).Dtor_cf36(), (_2374_dt__update__tmp_h0).Dtor_cf37(), (_2374_dt__update__tmp_h0).Dtor_cf38(), _2375_dt__update_hcf39_h0)
						}(_pat_let67_0)
					}(_pat_let_tv59)
				}(_pat_let66_0)
			}(_2371_v42)).Dtor_cf39(), ((_this).F27()) || (_this.F30))
			_2373_v44 = (_2373_v44).Update((_2372_v43).Update(_this.F30, _dafny.IntOfInt64(990)), !(true))
			if (_this).F26() {
				var _2376_v45 _dafny.Array
				_ = _2376_v45
				var _nw374 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw374
				_2376_v45 = _nw374
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2376_v45), 0))
				_ = _index346
				(_2376_v45).ArraySet1((_this).F31(), (_index346).Int())
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2376_v45), 0))
				_ = _index347
				(_2376_v45).ArraySet1(_dafny.IntOfInt64(537), (_index347).Int())
				var _2377_v46 _dafny.Sequence
				_ = _2377_v46
				_2377_v46 = _dafny.UnicodeSeqOfUtf8Bytes("ohnwk")
				var _2378_v47 _dafny.Map
				_ = _2378_v47
				_2378_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-558), _2377_v46)
				var _2379_v48 *C1
				_ = _2379_v48
				var _nw375 *C1 = New_C1_()
				_ = _nw375
				_nw375.Ctor__(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2378_v47).Contains(_2325_v12.F40) {
						return (_2378_v47).Get(_2325_v12.F40).(_dafny.Sequence)
					}
					return _2377_v46
				})()).Cardinality()), _2325_v12.F40, (_this).F27(), (_this).F27())
				_2379_v48 = _nw375
				var _2380_v49 _dafny.Map
				_ = _2380_v49
				_2380_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2379_v48).F36(), _2379_v48)
				var _2381_v50 _dafny.Sequence
				_ = _2381_v50
				_2381_v50 = _dafny.SeqOf(_2379_v48, _2379_v48, (func() *C1 {
					if (_2380_v49).Contains((_this).F31()) {
						return (_2380_v49).Get((_this).F31()).(*C1)
					}
					return _2379_v48
				})(), _2379_v48)
				var _2382_v51 _dafny.Map
				_ = _2382_v51
				_2382_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2381_v50, _2379_v48.F37)
				_2382_v51 = _2382_v51
				var _2383_v52 _dafny.Sequence
				_ = _2383_v52
				_2383_v52 = _dafny.SeqOf(_dafny.IntOfUint32((_2377_v46).Cardinality()))
				var _2384_v53 _dafny.Sequence
				_ = _2384_v53
				_2384_v53 = _dafny.SeqOf(_dafny.IntOfUint32((_2383_v52).Cardinality()))
				var _2385_v54 D0
				_ = _2385_v54
				_2385_v54 = Companion_D0_.Create_DC0_(!(true))
				var _2386_v55 _dafny.MultiSet
				_ = _2386_v55
				_2386_v55 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_2383_v52, (Companion_Default___.SafeIndex((_this).Fm4(_2385_v54, _dafny.IntOfInt64(713), _dafny.IntOfInt64(453), globalState), _dafny.IntOfUint32((_2383_v52).Cardinality()))).Uint32(), _2325_v12.F40), _dafny.SeqOf(_dafny.IntOfInt64(-911)), _2383_v52)
				var _2387_v56 *C3
				_ = _2387_v56
				var _nw376 *C3 = New_C3_()
				_ = _nw376
				_nw376.Ctor__(_2386_v55, (_this).F27(), (_this).F26())
				_2387_v56 = _nw376
				var _2388_v57 _dafny.Map
				_ = _2388_v57
				_2388_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _2387_v56)
				var _2389_v58 _dafny.CodePoint
				_ = _2389_v58
				_2389_v58 = _dafny.CodePoint('s')
				var _2390_v59 D23
				_ = _2390_v59
				_2390_v59 = Companion_D23_.Create_DC55_(_2387_v56)
				_2387_v56 = (func() *C3 {
					if (_2388_v57).Contains(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_2377_v46, (Companion_Default___.SafeIndex(_2325_v12.F40, _dafny.IntOfUint32((_2377_v46).Cardinality()))).Uint32(), _2389_v58), _2389_v58)) {
						return (_2388_v57).Get(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_2377_v46, (Companion_Default___.SafeIndex(_2325_v12.F40, _dafny.IntOfUint32((_2377_v46).Cardinality()))).Uint32(), _2389_v58), _2389_v58)).(*C3)
					}
					return (func() *C3 {
						if (_2388_v57).Contains((_this).F26()) {
							return (_2388_v57).Get((_this).F26()).(*C3)
						}
						return (_2390_v59).Dtor_cf103()
					})()
				})()
				(globalState).F2 = ((_this).F27()) && (_this.F30)
				(globalState).F12 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_2376_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_2376_v45), 0))).Int()).(_dafny.Int), _2325_v12.F40), _2325_v12.F40)
			} else {
				var _2391_v60 _dafny.MultiSet
				_ = _2391_v60
				_2391_v60 = _dafny.MultiSetOf(_2325_v12.F40)
				_2391_v60 = _2391_v60
				var _2392_v61 _dafny.MultiSet
				_ = _2392_v61
				_2392_v61 = _dafny.MultiSetOf(Companion_D1_.Create_DC3_(_dafny.SeqOf((_this).F26())))
				var _2393_v62 _dafny.Sequence
				_ = _2393_v62
				_2393_v62 = _dafny.SeqOf((_this).F26(), (_this).F27())
				var _2394_v63 D1
				_ = _2394_v63
				_2394_v63 = Companion_D1_.Create_DC3_(_2393_v62)
				var _pat_let_tv60 = _2393_v62
				_ = _pat_let_tv60
				var _2395_v64 _dafny.Sequence
				_ = _2395_v64
				_2395_v64 = _dafny.SeqOf(func(_pat_let68_0 D1) D1 {
					return func(_2396_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let69_0 _dafny.Sequence) D1 {
							return func(_2397_dt__update_hcf6_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_2397_dt__update_hcf6_h0)
							}(_pat_let69_0)
						}(_pat_let_tv60)
					}(_pat_let68_0)
				}(_2394_v63), _2394_v63)
				var _2398_v65 _dafny.Set
				_ = _2398_v65
				_2398_v65 = _dafny.SetOf(_dafny.IntOfUint32((_2393_v62).Cardinality()))
				var _2399_v66 _dafny.Map
				_ = _2399_v66
				_2399_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2398_v65, _dafny.IntOfInt64(979))
				var _2400_v67 D24
				_ = _2400_v67
				_2400_v67 = Companion_D24_.Create_DC59_(_2392_v61)
				var _pat_let_tv61 = _2393_v62
				_ = _pat_let_tv61
				var _pat_let_tv62 = _2393_v62
				_ = _pat_let_tv62
				var _2401_v68 _dafny.Array
				_ = _2401_v68
				var _nwElement0_91 _dafny.MultiSet = _2392_v61
				_ = _nwElement0_91
				var _nw377 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(27))
				_ = _nw377
				(_nw377).ArraySet1(_nwElement0_91, 0)
				(_nw377).ArraySet1((_2392_v61).Union(_2392_v61), 1)
				(_nw377).ArraySet1(_2392_v61, 2)
				(_nw377).ArraySet1((_2392_v61).Update(_2394_v63, Companion_Default___.Abs(_dafny.IntOfUint32((_2393_v62).Cardinality()))), 3)
				(_nw377).ArraySet1(((_2392_v61).Update(_2394_v63, Companion_Default___.Abs(_2325_v12.F40))).Intersection(_2392_v61), 4)
				(_nw377).ArraySet1(_2392_v61, 5)
				(_nw377).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf(_2394_v63, _2394_v63, _2394_v63))).Union(_2392_v61), 6)
				(_nw377).ArraySet1(_2392_v61, 7)
				(_nw377).ArraySet1((_dafny.MultiSetOf(func(_pat_let70_0 D1) D1 {
					return func(_2402_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let71_0 _dafny.Sequence) D1 {
							return func(_2403_dt__update_hcf6_h1 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_2403_dt__update_hcf6_h1)
							}(_pat_let71_0)
						}(_pat_let_tv61)
					}(_pat_let70_0)
				}(_2394_v63))).Union(_2392_v61), 8)
				(_nw377).ArraySet1(_dafny.MultiSetFromSeq(_2395_v64), 9)
				(_nw377).ArraySet1((_2392_v61).Union(_2392_v61), 10)
				(_nw377).ArraySet1(_dafny.MultiSetOf(_2394_v63, func(_pat_let72_0 D1) D1 {
					return func(_2404_dt__update__tmp_h3 D1) D1 {
						return func(_pat_let73_0 _dafny.Sequence) D1 {
							return func(_2405_dt__update_hcf6_h2 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC3_(_2405_dt__update_hcf6_h2)
							}(_pat_let73_0)
						}(_pat_let_tv62)
					}(_pat_let72_0)
				}(_2394_v63)), 11)
				(_nw377).ArraySet1((_dafny.MultiSetFromSeq(_2395_v64)).Union(_2392_v61), 12)
				(_nw377).ArraySet1(_2392_v61, 13)
				(_nw377).ArraySet1((_2392_v61).Update(_2394_v63, Companion_Default___.Abs((_2399_v66).Cardinality())), 14)
				(_nw377).ArraySet1((_2392_v61).Difference((_2400_v67).Dtor_cf109()), 15)
				(_nw377).ArraySet1(_2392_v61, 16)
				(_nw377).ArraySet1(_2392_v61, 17)
				(_nw377).ArraySet1(_2392_v61, 18)
				(_nw377).ArraySet1((_dafny.MultiSetOf(_2394_v63)).Update(_2394_v63, Companion_Default___.Abs((_this).F31())), 19)
				(_nw377).ArraySet1(_dafny.MultiSetOf(_2394_v63, _2394_v63, _2394_v63, _2394_v63), 20)
				(_nw377).ArraySet1((_2392_v61).Intersection(_2392_v61), 21)
				(_nw377).ArraySet1(_2392_v61, 22)
				(_nw377).ArraySet1(_dafny.MultiSetOf(_2394_v63, Companion_D1_.Create_DC3_(_dafny.Companion_Sequence_.Update(_2393_v62, (Companion_Default___.SafeIndex((_this).F31(), _dafny.IntOfUint32((_2393_v62).Cardinality()))).Uint32(), (_this).F26()))), 23)
				(_nw377).ArraySet1(_2392_v61, 24)
				(_nw377).ArraySet1(_2392_v61, 25)
				(_nw377).ArraySet1(_2392_v61, 26)
				_2401_v68 = _nw377
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_2401_v68), 0))
				_ = _index348
				(_2401_v68).ArraySet1(_dafny.MultiSetFromSeq(_2395_v64), (_index348).Int())
				var _2406_v69 _dafny.Map
				_ = _2406_v69
				_2406_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2325_v12.F40, _this.F30)).Cardinality(), (_this).F26())
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_2401_v68), 0))
				_ = _index349
				var _rhs348 _dafny.MultiSet = (_2392_v61).Update(Companion_D1_.Create_DC3_(_2393_v62), Companion_Default___.Abs((func() _dafny.Int {
					if (_2370_v41).Fm2((_this).F27(), (_2406_v69).Cardinality(), globalState) {
						return _2325_v12.F40
					}
					return _dafny.IntOfInt64(120)
				})()))
				_ = _rhs348
				var _rhs349 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_2325_v12.F40, _2325_v12.F40))
				_ = _rhs349
				var _rhs350 _dafny.Int = _2325_v12.F40
				_ = _rhs350
				var _rhs351 _dafny.Int = (_2325_v12.F40).Plus(_2325_v12.F40)
				_ = _rhs351
				var _rhs352 _dafny.Int = (func() _dafny.Int {
					if (_this).F27() {
						return Companion_Default___.Fm39(globalState)
					}
					return _2325_v12.F40
				})()
				_ = _rhs352
				var _lhs321 _dafny.Array = _2401_v68
				_ = _lhs321
				var _lhs322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_2401_v68), 0))
				_ = _lhs322
				var _lhs323 *GlobalState = globalState
				_ = _lhs323
				var _lhs324 *GlobalState = globalState
				_ = _lhs324
				var _lhs325 *GlobalState = globalState
				_ = _lhs325
				var _lhs326 *C6 = _2325_v12
				_ = _lhs326
				(_lhs321).ArraySet1(_rhs348, (_lhs322).Int())
				_lhs323.F12 = _rhs349
				_lhs324.F16 = _rhs350
				_lhs325.F16 = _rhs351
				_lhs326.F40 = _rhs352
				var _2407_v70 _dafny.Array
				_ = _2407_v70
				var _nw378 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
				_ = _nw378
				_2407_v70 = _nw378
				(globalState).F5 = (func() _dafny.Array {
					if (_this).F27() {
						return _2407_v70
					}
					return _2407_v70
				})()
				_2393_v62 = (func() _dafny.Sequence {
					if (_this).F27() {
						return _2393_v62
					}
					return _dafny.SeqOf(_this.F30)
				})()
				var _2408_v71 _dafny.MultiSet
				_ = _2408_v71
				_2408_v71 = _dafny.MultiSetOf((_this).F26())
				var _2409_v72 _dafny.Sequence
				_ = _2409_v72
				_2409_v72 = _dafny.SeqOf(_dafny.MultiSetOf((_this).F27(), true), _2408_v71, _dafny.MultiSetOf(false))
				var _2410_v73 _dafny.Map
				_ = _2410_v73
				_2410_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F26(), _2408_v71)
				var _2411_v74 D14
				_ = _2411_v74
				_2411_v74 = Companion_D14_.Create_DC32_(_2408_v71, (_this).F27(), (_this).F26(), _this.F30, _dafny.IntOfInt64(407))
				var _2412_v75 _dafny.Array
				_ = _2412_v75
				var _nwElement0_92 _dafny.MultiSet = _2408_v71
				_ = _nwElement0_92
				var _nw379 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(15))
				_ = _nw379
				(_nw379).ArraySet1(_nwElement0_92, 0)
				(_nw379).ArraySet1(_2408_v71, 1)
				(_nw379).ArraySet1(_2408_v71, 2)
				(_nw379).ArraySet1(Companion_Default___.Fm21((_this).F31(), globalState), 3)
				(_nw379).ArraySet1((_2409_v72).Select((Companion_Default___.SafeIndex(_2325_v12.F40, _dafny.IntOfUint32((_2409_v72).Cardinality()))).Uint32()).(_dafny.MultiSet), 4)
				(_nw379).ArraySet1((_2408_v71).Difference(_2408_v71), 5)
				(_nw379).ArraySet1((func() _dafny.MultiSet {
					if (_2410_v73).Contains(_this.F30) {
						return (_2410_v73).Get(_this.F30).(_dafny.MultiSet)
					}
					return _2408_v71
				})(), 6)
				(_nw379).ArraySet1((_2411_v74).Dtor_cf56(), 7)
				(_nw379).ArraySet1(_2408_v71, 8)
				(_nw379).ArraySet1(_2408_v71, 9)
				(_nw379).ArraySet1(_dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if (_this).F27() {
						return _2393_v62
					}
					return _2393_v62
				})()), 10)
				(_nw379).ArraySet1(_2408_v71, 11)
				(_nw379).ArraySet1((_2409_v72).Select((Companion_Default___.SafeIndex(_2325_v12.F40, _dafny.IntOfUint32((_2409_v72).Cardinality()))).Uint32()).(_dafny.MultiSet), 12)
				(_nw379).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).Fm2((_this).F27(), _dafny.IntOfInt64(49), globalState))), 13)
				(_nw379).ArraySet1(_2408_v71, 14)
				_2412_v75 = _nw379
				var _2413_v76 _dafny.Array
				_ = _2413_v76
				var _len0_67 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_67
				var _nw380 _dafny.Array
				_ = _nw380
				if _len0_67.Cmp(_dafny.Zero) == 0 {
					_nw380 = _dafny.NewArray(_len0_67)
				} else {
					var _init67 func(_dafny.Int) _dafny.Int = (func(_2414_v12 *C6) func(_dafny.Int) _dafny.Int {
						return func(_2415_i6 _dafny.Int) _dafny.Int {
							return (_2415_i6).Times((_dafny.Zero).Minus(_2414_v12.F40))
						}
					})(_2325_v12)
					_ = _init67
					var _element0_67 = _init67(_dafny.Zero)
					_ = _element0_67
					_nw380 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
					(_nw380).ArraySet1(_element0_67, 0)
					var _nativeLen0_67 = (_len0_67).Int()
					_ = _nativeLen0_67
					for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
						(_nw380).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
					}
				}
				_2413_v76 = _nw380
				var _2416_v77 _dafny.Sequence
				_ = _2416_v77
				_2416_v77 = _dafny.SeqOf(_2413_v76)
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_2412_v75), 0))
				_ = _index350
				(_2412_v75).ArraySet1(Companion_Default___.Fm21(_dafny.IntOfUint32((_2416_v77).Cardinality()), globalState), (_index350).Int())
				var _2417_v78 _dafny.Sequence
				_ = _2417_v78
				_2417_v78 = _dafny.SeqOf((_2325_v12.F40).Plus(_2325_v12.F40))
				var _2418_v79 _dafny.Map
				_ = _2418_v79
				_2418_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2325_v12.F40, _2325_v12.F40)
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_2412_v75), 0))
				_ = _index351
				var _rhs353 _dafny.Int = (func() _dafny.Int {
					if _this.F30 {
						return _2325_v12.F40
					}
					return Companion_Default___.SafeModuloInt((_this).F31(), _2325_v12.F40)
				})()
				_ = _rhs353
				var _rhs354 _dafny.Int = (_2417_v78).Select((Companion_Default___.SafeIndex((_this).Fm1(globalState), _dafny.IntOfUint32((_2417_v78).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs354
				var _rhs355 _dafny.MultiSet = _2408_v71
				_ = _rhs355
				var _rhs356 bool = (_this).Fm2((_2325_v12).Fm2((_this).F27(), _2325_v12.F40, globalState), Companion_Default___.SafeDivisionInt(_2325_v12.F40, (_2418_v79).Cardinality()), globalState)
				_ = _rhs356
				var _lhs327 *GlobalState = globalState
				_ = _lhs327
				var _lhs328 *C6 = _2325_v12
				_ = _lhs328
				var _lhs329 _dafny.Array = _2412_v75
				_ = _lhs329
				var _lhs330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(434), _dafny.ArrayLen((_2412_v75), 0))
				_ = _lhs330
				var _lhs331 *GlobalState = globalState
				_ = _lhs331
				_lhs327.F16 = _rhs353
				_lhs328.F40 = _rhs354
				(_lhs329).ArraySet1(_rhs355, (_lhs330).Int())
				_lhs331.F22 = _rhs356
			}
			var _2419_v80 _dafny.Sequence
			_ = _2419_v80
			_2419_v80 = _dafny.SeqOf(_dafny.SeqOf((_this).F31(), (_this).F31()))
			var _2420_v81 _dafny.Sequence
			_ = _2420_v81
			_2420_v81 = _dafny.SeqOf(_2419_v80)
			var _2421_v82 T2
			_ = _2421_v82
			var _nw381 *C2 = New_C2_()
			_ = _nw381
			_nw381.Ctor__((func() _dafny.Sequence {
				if _this.F30 {
					return (_2420_v81).Select((Companion_Default___.SafeIndex(_2325_v12.F40, _dafny.IntOfUint32((_2420_v81).Cardinality()))).Uint32()).(_dafny.Sequence)
				}
				return _2419_v80
			})(), !(true) || (!((_this).F26())), (_this).F27())
			_2421_v82 = _nw381
			_2421_v82 = _2421_v82
		}
		var _2422_v83 _dafny.Sequence
		_ = _2422_v83
		_2422_v83 = _dafny.UnicodeSeqOfUtf8Bytes("uuuanc")
		var _2423_v84 _dafny.Sequence
		_ = _2423_v84
		_2423_v84 = _dafny.SeqOf((_dafny.IntOfInt64(884)).Plus(_dafny.IntOfUint32((_2422_v83).Cardinality())), _2325_v12.F40)
		var _2424_v85 _dafny.CodePoint
		_ = _2424_v85
		_2424_v85 = _dafny.CodePoint('i')
		var _2425_v86 _dafny.Map
		_ = _2425_v86
		_2425_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, _2424_v85)
		_2423_v84 = _dafny.Companion_Sequence_.Concatenate(_2423_v84, _dafny.Companion_Sequence_.Concatenate(_2423_v84, Companion_Default___.Fm43((_this).F31(), _2325_v12.F40, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yjmphn")).Cardinality()), _2425_v86, globalState)))
		var _2426_v87 _dafny.Array
		_ = _2426_v87
		var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw382
		_2426_v87 = _nw382
		var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_2426_v87), 0))
		_ = _index352
		(_2426_v87).ArraySet1((_this).F31(), (_index352).Int())
		var _2427_v88 _dafny.Array
		_ = _2427_v88
		var _len0_68 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_68
		var _nw383 _dafny.Array
		_ = _nw383
		if _len0_68.Cmp(_dafny.Zero) == 0 {
			_nw383 = _dafny.NewArray(_len0_68)
		} else {
			var _init68 func(_dafny.Int) bool = func(_2428_i7 _dafny.Int) bool {
				return (_this).F27()
			}
			_ = _init68
			var _element0_68 = _init68(_dafny.Zero)
			_ = _element0_68
			_nw383 = _dafny.NewArrayFromExample(_element0_68, nil, _len0_68)
			(_nw383).ArraySet1(_element0_68, 0)
			var _nativeLen0_68 = (_len0_68).Int()
			_ = _nativeLen0_68
			for _i0_68 := 1; _i0_68 < _nativeLen0_68; _i0_68++ {
				(_nw383).ArraySet1(_init68(_dafny.IntOf(_i0_68)), _i0_68)
			}
		}
		_2427_v88 = _nw383
		var _2429_v89 _dafny.Set
		_ = _2429_v89
		_2429_v89 = _dafny.SetOf(_2427_v88)
		var _2430_v90 _dafny.Sequence
		_ = _2430_v90
		_2430_v90 = _dafny.SeqOf(_2429_v89, _2429_v89, _2429_v89)
		var _2431_v91 _dafny.Map
		_ = _2431_v91
		_2431_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2325_v12.F40, _2429_v89)
		var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_2426_v87), 0))
		_ = _index353
		(_2426_v87).ArraySet1((((_2430_v90).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_2325_v12.F40), _dafny.IntOfUint32((_2430_v90).Cardinality()))).Uint32()).(_dafny.Set)).Union((func() _dafny.Set {
			if (_2431_v91).Contains(_2325_v12.F40) {
				return (_2431_v91).Get(_2325_v12.F40).(_dafny.Set)
			}
			return _2429_v89
		})())).Cardinality(), (_index353).Int())
		var _2432_v92 _dafny.Sequence
		_ = _2432_v92
		_2432_v92 = _dafny.SeqOf((_this).F26())
		var _2433_v93 D16
		_ = _2433_v93
		_2433_v93 = Companion_D16_.Create_DC38_(_dafny.IntOfInt64(718), _dafny.IntOfUint32((_2432_v92).Cardinality()), (_this).F26())
		var _pat_let_tv63 = _2426_v87
		_ = _pat_let_tv63
		var _pat_let_tv64 = _2426_v87
		_ = _pat_let_tv64
		var _pat_let_tv65 = _2432_v92
		_ = _pat_let_tv65
		var _pat_let_tv66 = _2326_v13
		_ = _pat_let_tv66
		r0 = func(_source34 D16) bool {
			if _source34.Is_DC38() {
				var _2434___mcc_h11 _dafny.Int = _source34.Get_().(D16_DC38).Cf69
				_ = _2434___mcc_h11
				var _2435___mcc_h12 _dafny.Int = _source34.Get_().(D16_DC38).Cf70
				_ = _2435___mcc_h12
				var _2436___mcc_h13 bool = _source34.Get_().(D16_DC38).Cf71
				_ = _2436___mcc_h13
				var _2437_cf71 bool = _2436___mcc_h13
				_ = _2437_cf71
				var _2438_cf70 _dafny.Int = _2435___mcc_h12
				_ = _2438_cf70
				var _2439_cf69 _dafny.Int = _2434___mcc_h11
				_ = _2439_cf69
				return (_this).F26()
			} else if _source34.Is_DC37() {
				var _2440___mcc_h14 _dafny.Map = _source34.Get_().(D16_DC37).Cf68
				_ = _2440___mcc_h14
				var _2441_cf68 _dafny.Map = _2440___mcc_h14
				_ = _2441_cf68
				return !(((_pat_let_tv64).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_pat_let_tv63), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_pat_let_tv65).Cardinality())) < 0)
			} else {
				var _2442___mcc_h15 D16 = _source34.Get_().(D16_DC39).Cf72
				_ = _2442___mcc_h15
				var _2443_cf72 D16 = _2442___mcc_h15
				_ = _2443_cf72
				return !(!(_pat_let_tv66).Equals((_dafny.SetOf(_this.F30, true))))
			}
		}(_2433_v93)
		r1 = (_this).F27()
		return r0, r1
	}
}
func (_this *C9) M5(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _2444_v0 _dafny.Sequence
		_ = _2444_v0
		_2444_v0 = _dafny.UnicodeSeqOfUtf8Bytes("flsqn")
		var _2445_v1 _dafny.Sequence
		_ = _2445_v1
		_2445_v1 = _dafny.SeqOf(_dafny.IntOfUint32((_2444_v0).Cardinality()), p1)
		var _2446_v2 _dafny.Array
		_ = _2446_v2
		var _nwElement0_93 _dafny.Int = p1
		_ = _nwElement0_93
		var _nw384 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_93, nil, _dafny.IntOfInt64(14))
		_ = _nw384
		(_nw384).ArraySet1(_nwElement0_93, 0)
		(_nw384).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_this.F30)).Cardinality()), 1)
		(_nw384).ArraySet1((_this).F31(), 2)
		(_nw384).ArraySet1((_this).F31(), 3)
		(_nw384).ArraySet1((_this).F31(), 4)
		(_nw384).ArraySet1(_dafny.IntOfUint32((_2445_v1).Cardinality()), 5)
		(_nw384).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p1)).Cardinality(), 6)
		(_nw384).ArraySet1((_this).F31(), 7)
		(_nw384).ArraySet1(p1, 8)
		(_nw384).ArraySet1((_this).F31(), 9)
		(_nw384).ArraySet1(p0, 10)
		(_nw384).ArraySet1(p1, 11)
		(_nw384).ArraySet1(p1, 12)
		(_nw384).ArraySet1(p0, 13)
		_2446_v2 = _nw384
		var _2447_v3 _dafny.Map
		_ = _2447_v3
		_2447_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D3_.Create_DC10_(_2446_v2)).Dtor_cf21(), p1)
		_2447_v3 = (_2447_v3).Update(_2446_v2, (_dafny.Zero).Minus(p0))
		r1 = Companion_Default___.Fm39(globalState)
		var _2448_v4 _dafny.MultiSet
		_ = _2448_v4
		_2448_v4 = _dafny.MultiSetOf(p2)
		var _2449_v5 _dafny.Set
		_ = _2449_v5
		_2449_v5 = _dafny.SetOf(_dafny.IntOfUint32((_2445_v1).Cardinality()), p1, p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ujlp")).Cardinality()), (_dafny.Zero).Minus((_2448_v4).Cardinality()))
		var _2450_v6 _dafny.Map
		_ = _2450_v6
		_2450_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2449_v5)
		var _2451_v7 _dafny.Sequence
		_ = _2451_v7
		_2451_v7 = _dafny.SeqOf(_dafny.SetOf(p1, _dafny.IntOfInt64(-51), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), (_this).F27())).Cardinality()), _dafny.SetOf(p0), (func() _dafny.Set {
			if (_2450_v6).Contains((_this).F27()) {
				return (_2450_v6).Get((_this).F27()).(_dafny.Set)
			}
			return _2449_v5
		})())
		var _2452_v8 _dafny.Map
		_ = _2452_v8
		_2452_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2451_v7, (_this).F26())
		var _2453_i0 _dafny.Int
		_ = _2453_i0
		_2453_i0 = _dafny.Zero
		{
			for !((func() bool {
				if (_2452_v8).Contains(_2451_v7) {
					return (_2452_v8).Get(_2451_v7).(bool)
				}
				return _this.F30
			})()) {
				{
					if (_2453_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_2453_i0 = (_2453_i0).Plus(_dafny.One)
					var _2454_v9 _dafny.Sequence
					_ = _2454_v9
					_2454_v9 = _dafny.SeqOf(_this.F30, true)
					var _2455_v10 _dafny.Map
					_ = _2455_v10
					_2455_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_2454_v9, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2454_v9).Cardinality()))).Uint32(), true), p0)
					var _2456_v11 *C5
					_ = _2456_v11
					var _nw385 *C5 = New_C5_()
					_ = _nw385
					_nw385.Ctor__((_this).F26(), _this.F30)
					_2456_v11 = _nw385
					var _2457_v12 _dafny.Map
					_ = _2457_v12
					_2457_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), _2456_v11)
					var _2458_v13 _dafny.Set
					_ = _2458_v13
					_2458_v13 = _dafny.SetOf(_this.F30)
					var _2459_v14 _dafny.Set
					_ = _2459_v14
					_2459_v14 = _dafny.SetOf(_2458_v13)
					var _2460_v15 _dafny.Map
					_ = _2460_v15
					_2460_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2459_v14).Cardinality(), (_this).F26())
					var _2461_v16 D16
					_ = _2461_v16
					_2461_v16 = Companion_D16_.Create_DC38_((_2457_v12).Cardinality(), (_2460_v15).Cardinality(), (_this).F27())
					var _2462_v17 _dafny.Sequence
					_ = _2462_v17
					_2462_v17 = _dafny.SeqOf(_2461_v16, _2461_v16)
					r1 = (_this).Fm4(Companion_D0_.Create_DC0_((_this).F26()), (func() _dafny.Int {
						if (_2455_v10).Contains(_dafny.SeqOf((_this).F27(), _this.F30)) {
							return (_2455_v10).Get(_dafny.SeqOf((_this).F27(), _this.F30)).(_dafny.Int)
						}
						return p1
					})(), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2454_v9).Cardinality()), _dafny.IntOfUint32((_2462_v17).Cardinality())), globalState)
					(globalState).F16 = _dafny.IntOfInt64(154)
					var _2463_v18 _dafny.Map
					_ = _2463_v18
					_2463_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if !((_this).F27()) {
							return (_this).F26()
						}
						return p2
					})(), !((_this).F26()) || (_this.F30))
					_2463_v18 = (_2463_v18).Update(p2, _this.F30)
					var _2464_v19 _dafny.MultiSet
					_ = _2464_v19
					_2464_v19 = _dafny.MultiSetOf(_2445_v1)
					var _2465_v20 *C3
					_ = _2465_v20
					var _nw386 *C3 = New_C3_()
					_ = _nw386
					_nw386.Ctor__(_2464_v19, true, (_dafny.MultiSetOf(_this.F30)).Contains(p2))
					_2465_v20 = _nw386
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		r1 = (_this).F31()
		(globalState).F15 = false
		if (_this).Fm2(p2, p1, globalState) {
			var _2466_v21 _dafny.Array
			_ = _2466_v21
			var _nw387 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
			_ = _nw387
			_2466_v21 = _nw387
			var _2467_v22 _dafny.Array
			_ = _2467_v22
			var _len0_69 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_69
			var _nw388 _dafny.Array
			_ = _nw388
			if _len0_69.Cmp(_dafny.Zero) == 0 {
				_nw388 = _dafny.NewArray(_len0_69)
			} else {
				var _init69 func(_dafny.Int) D20 = func(_2468_i1 _dafny.Int) D20 {
					return (func() D20 {
						if _this.F30 {
							return Companion_D20_.Create_DC47_()
						}
						return Companion_D20_.Create_DC47_()
					})()
				}
				_ = _init69
				var _element0_69 = _init69(_dafny.Zero)
				_ = _element0_69
				_nw388 = _dafny.NewArrayFromExample(_element0_69, nil, _len0_69)
				(_nw388).ArraySet1(_element0_69, 0)
				var _nativeLen0_69 = (_len0_69).Int()
				_ = _nativeLen0_69
				for _i0_69 := 1; _i0_69 < _nativeLen0_69; _i0_69++ {
					(_nw388).ArraySet1(_init69(_dafny.IntOf(_i0_69)), _i0_69)
				}
			}
			_2467_v22 = _nw388
			var _2469_v23 D20
			_ = _2469_v23
			_2469_v23 = Companion_D20_.Create_DC47_()
			var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_2467_v22), 0))
			_ = _index354
			(_2467_v22).ArraySet1(_2469_v23, (_index354).Int())
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_2467_v22), 0))
			_ = _index355
			var _rhs357 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2444_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-395))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg119 _dafny.Int) interface{} {
					return coer119(arg119)
				}
			}(func(_2470_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))
			_ = _rhs357
			var _rhs358 _dafny.Array = _2466_v21
			_ = _rhs358
			var _rhs359 _dafny.Array = _2446_v2
			_ = _rhs359
			var _rhs360 D20 = _2469_v23
			_ = _rhs360
			var _lhs332 *GlobalState = globalState
			_ = _lhs332
			var _lhs333 _dafny.Array = _2467_v22
			_ = _lhs333
			var _lhs334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(120), _dafny.ArrayLen((_2467_v22), 0))
			_ = _lhs334
			_lhs332.F6 = _rhs357
			_2466_v21 = _rhs358
			_2446_v2 = _rhs359
			(_lhs333).ArraySet1(_rhs360, (_lhs334).Int())
			(globalState).F16 = (((_2448_v4).Union(_2448_v4)).Difference(_2448_v4)).Cardinality()
			r1 = (_dafny.Zero).Minus(p0)
			var _2471_v24 _dafny.Map
			_ = _2471_v24
			_2471_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2449_v5).Cardinality(), (_this).F26())
			var _2472_v25 D0
			_ = _2472_v25
			_2472_v25 = Companion_D0_.Create_DC0_((_this).F27())
			var _2473_v26 _dafny.Sequence
			_ = _2473_v26
			_2473_v26 = _dafny.SeqOf((_this).F26(), (_this).F26())
			if (((func() _dafny.Map {
				if p2 {
					return _2471_v24
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F26())
			})()).Cardinality()).Cmp((_this).Fm4(_2472_v25, _dafny.IntOfUint32((_2473_v26).Cardinality()), p0, globalState)) > 0 {
				var _2474_v27 _dafny.Map
				_ = _2474_v27
				_2474_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				(globalState).F12 = (((_2474_v27).Update(p0, p0)).Merge(_2474_v27)).Cardinality()
				var _2475_v28 _dafny.Array
				_ = _2475_v28
				var _nw389 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
				_ = _nw389
				_2475_v28 = _nw389
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2475_v28), 0))
				_ = _index356
				(_2475_v28).ArraySet1((_this).F29(), (_index356).Int())
				var _2476_v29 _dafny.CodePoint
				_ = _2476_v29
				_2476_v29 = _dafny.CodePoint('p')
				var _2477_v30 D6
				_ = _2477_v30
				_2477_v30 = Companion_D6_.Create_DC16_(_2444_v0)
				var _2478_v31 _dafny.Array
				_ = _2478_v31
				var _nwElement0_94 _dafny.Sequence = _2444_v0
				_ = _nwElement0_94
				var _nw390 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_94, nil, _dafny.IntOfInt64(16))
				_ = _nw390
				(_nw390).ArraySet1(_nwElement0_94, 0)
				(_nw390).ArraySet1(_2444_v0, 1)
				(_nw390).ArraySet1(_2444_v0, 2)
				(_nw390).ArraySet1(_2444_v0, 3)
				(_nw390).ArraySet1(_2444_v0, 4)
				(_nw390).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(881))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg120 _dafny.Int) interface{} {
						return coer120(arg120)
					}
				}(func(_2479_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), _2444_v0), 5)
				(_nw390).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gixiulo"), 6)
				(_nw390).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(521))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg121 _dafny.Int) interface{} {
						return coer121(arg121)
					}
				}(func(_2480_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				})), 7)
				(_nw390).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gxxceu"), 8)
				(_nw390).ArraySet1((func() _dafny.Sequence {
					if _this.F30 {
						return _dafny.UnicodeSeqOfUtf8Bytes("qkrwt")
					}
					return _2444_v0
				})(), 9)
				(_nw390).ArraySet1(_dafny.Companion_Sequence_.Update(_2444_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2444_v0).Cardinality()), _dafny.IntOfUint32((_2444_v0).Cardinality()))).Uint32(), _2476_v29), 10)
				(_nw390).ArraySet1(_2444_v0, 11)
				(_nw390).ArraySet1(_2444_v0, 12)
				(_nw390).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2444_v0, _2444_v0), 13)
				(_nw390).ArraySet1(_2444_v0, 14)
				(_nw390).ArraySet1((_2477_v30).Dtor_cf29(), 15)
				_2478_v31 = _nw390
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2478_v31), 0))
				_ = _index357
				(_2478_v31).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2444_v0, Companion_Default___.Fm22(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aj")).Cardinality()), _this.F30, globalState)), (_index357).Int())
				var _2481_v32 _dafny.Sequence
				_ = _2481_v32
				_2481_v32 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_2473_v26), ((_2448_v4).Update(true, Companion_Default___.Abs((_2445_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.IntOfUint32((_2445_v1).Cardinality()))).Uint32()).(_dafny.Int)))).Update((_this).F27(), Companion_Default___.Abs(p1)))
				var _2482_v33 _dafny.Sequence
				_ = _2482_v33
				_2482_v33 = _dafny.SeqOf(Companion_Default___.Fm22(p0, (_this).F26(), globalState))
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2475_v28), 0))
				_ = _index358
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2478_v31), 0))
				_ = _index359
				var _rhs361 _dafny.MultiSet = (_2481_v32).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_2448_v4).Contains((_this).F26()) {
						return (_2448_v4).Multiplicity((_this).F26())
					}
					return p1
				})(), _dafny.IntOfUint32((_2481_v32).Cardinality()))).Uint32()).(_dafny.MultiSet)
				_ = _rhs361
				var _rhs362 _dafny.MultiSet = (((_this).F29()).Difference(Companion_Default___.Fm42(((_this).F29()).Cardinality(), globalState))).Difference((_this).F29())
				_ = _rhs362
				var _rhs363 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pyblpod"), _dafny.Companion_Sequence_.Concatenate((_2482_v33).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2482_v33).Cardinality()))).Uint32()).(_dafny.Sequence), _2444_v0))
				_ = _rhs363
				var _lhs335 _dafny.Array = _2475_v28
				_ = _lhs335
				var _lhs336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_2475_v28), 0))
				_ = _lhs336
				var _lhs337 _dafny.Array = _2478_v31
				_ = _lhs337
				var _lhs338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_2478_v31), 0))
				_ = _lhs338
				_2448_v4 = _rhs361
				(_lhs335).ArraySet1(_rhs362, (_lhs336).Int())
				(_lhs337).ArraySet1(_rhs363, (_lhs338).Int())
				(globalState).F16 = p1
				var _2483_v34 *C0
				_ = _2483_v34
				var _nw391 *C0 = New_C0_()
				_ = _nw391
				_nw391.Ctor__(_this.F30)
				_2483_v34 = _nw391
				_2483_v34 = _2483_v34
				var _2484_v35 _dafny.Array
				_ = _2484_v35
				var _len0_70 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_70
				var _nw392 _dafny.Array
				_ = _nw392
				if _len0_70.Cmp(_dafny.Zero) == 0 {
					_nw392 = _dafny.NewArray(_len0_70)
				} else {
					var _init70 func(_dafny.Int) bool = func(_2485_i5 _dafny.Int) bool {
						return (_this).F27()
					}
					_ = _init70
					var _element0_70 = _init70(_dafny.Zero)
					_ = _element0_70
					_nw392 = _dafny.NewArrayFromExample(_element0_70, nil, _len0_70)
					(_nw392).ArraySet1(_element0_70, 0)
					var _nativeLen0_70 = (_len0_70).Int()
					_ = _nativeLen0_70
					for _i0_70 := 1; _i0_70 < _nativeLen0_70; _i0_70++ {
						(_nw392).ArraySet1(_init70(_dafny.IntOf(_i0_70)), _i0_70)
					}
				}
				_2484_v35 = _nw392
				var _2486_v36 D5
				_ = _2486_v36
				_2486_v36 = Companion_D5_.Create_DC14_(p1, p0, (_2483_v34).F32())
				var _2487_v37 *C1
				_ = _2487_v37
				var _nw393 *C1 = New_C1_()
				_ = _nw393
				_nw393.Ctor__(p0, p0, (_2486_v36).Dtor_cf27(), (_2474_v27).Equals(_2474_v27))
				_2487_v37 = _nw393
				var _2488_v38 _dafny.Map
				_ = _2488_v38
				_2488_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2483_v34).F32(), _dafny.IntOfInt64(-689))
				var _rhs364 _dafny.Array = _2484_v35
				_ = _rhs364
				var _rhs365 *C1 = _2487_v37
				_ = _rhs365
				var _rhs366 _dafny.Sequence = _dafny.SeqOf((_2483_v34).Fm7(_2488_v38, _2445_v1, p2, globalState))
				_ = _rhs366
				var _rhs367 bool = !(_this.F30)
				_ = _rhs367
				var _rhs368 bool = !(_this.F30) || ((false) || (p2))
				_ = _rhs368
				var _lhs339 *GlobalState = globalState
				_ = _lhs339
				var _lhs340 *GlobalState = globalState
				_ = _lhs340
				_2484_v35 = _rhs364
				_2487_v37 = _rhs365
				_2473_v26 = _rhs366
				_lhs339.F15 = _rhs367
				_lhs340.F15 = _rhs368
			} else {
				var _2489_v39 _dafny.Map
				_ = _2489_v39
				_2489_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F31(), p1)
				var _2490_v40 _dafny.Map
				_ = _2490_v40
				_2490_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F26()), _2489_v39)
				_2490_v40 = _2490_v40
				var _2491_v41 _dafny.CodePoint
				_ = _2491_v41
				_2491_v41 = _dafny.CodePoint('t')
				var _2492_v42 _dafny.CodePoint
				_ = _2492_v42
				var _2493_v43 _dafny.Sequence
				_ = _2493_v43
				var _2494_v44 bool
				_ = _2494_v44
				var _2495_v45 _dafny.Map
				_ = _2495_v45
				var _out67 _dafny.CodePoint
				_ = _out67
				var _out68 _dafny.Sequence
				_ = _out68
				var _out69 bool
				_ = _out69
				var _out70 _dafny.Map
				_ = _out70
				_out67, _out68, _out69, _out70 = Companion_Default___.M0(_2491_v41, globalState)
				_2492_v42 = _out67
				_2493_v43 = _out68
				_2494_v44 = _out69
				_2495_v45 = _out70
				var _2496_v46 *C0
				_ = _2496_v46
				var _nw394 *C0 = New_C0_()
				_ = _nw394
				_nw394.Ctor__((p0).Cmp((_this).F31()) < 0)
				_2496_v46 = _nw394
				var _2497_v47 *C2
				_ = _2497_v47
				var _nw395 *C2 = New_C2_()
				_ = _nw395
				_nw395.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(437))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg122 _dafny.Int) interface{} {
						return coer122(arg122)
					}
				}((func(_2498_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2499_i6 _dafny.Int) _dafny.Sequence {
						return _2498_v1
					}
				})(_2445_v1))), (_this).F26(), (_2496_v46).F32())
				_2497_v47 = _nw395
				var _2500_v48 *C5
				_ = _2500_v48
				var _nw396 *C5 = New_C5_()
				_ = _nw396
				_nw396.Ctor__((_this).F26(), (_this).F27())
				_2500_v48 = _nw396
				var _2501_v49 _dafny.Map
				_ = _2501_v49
				_2501_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_2444_v0).Cardinality())).Plus(p0), _2500_v48)
				var _2502_v50 _dafny.Map
				_ = _2502_v50
				_2502_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2496_v46).F32(), (_this).F27())
				_2501_v49 = (_2501_v49).Update(((_2502_v50).Merge(_2502_v50)).Cardinality(), _2500_v48)
			}
			var _2503_v51 _dafny.Set
			_ = _2503_v51
			_2503_v51 = _dafny.SetOf((_this).F26(), (_this).F27())
			(globalState).F16 = (_2503_v51).Cardinality()
		} else {
			var _2504_v52 *C5
			_ = _2504_v52
			var _nw397 *C5 = New_C5_()
			_ = _nw397
			_nw397.Ctor__((_dafny.IntOfInt64(126)).Cmp(p0) <= 0, (func() bool {
				if p2 {
					return true
				}
				return (_this).F27()
			})())
			_2504_v52 = _nw397
			var _2505_v53 _dafny.Array
			_ = _2505_v53
			var _len0_71 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_71
			var _nw398 _dafny.Array
			_ = _nw398
			if _len0_71.Cmp(_dafny.Zero) == 0 {
				_nw398 = _dafny.NewArray(_len0_71)
			} else {
				var _init71 func(_dafny.Int) D6 = (func(_2506_v0 _dafny.Sequence) func(_dafny.Int) D6 {
					return func(_2507_i7 _dafny.Int) D6 {
						return Companion_D6_.Create_DC16_(_2506_v0)
					}
				})(_2444_v0)
				_ = _init71
				var _element0_71 = _init71(_dafny.Zero)
				_ = _element0_71
				_nw398 = _dafny.NewArrayFromExample(_element0_71, nil, _len0_71)
				(_nw398).ArraySet1(_element0_71, 0)
				var _nativeLen0_71 = (_len0_71).Int()
				_ = _nativeLen0_71
				for _i0_71 := 1; _i0_71 < _nativeLen0_71; _i0_71++ {
					(_nw398).ArraySet1(_init71(_dafny.IntOf(_i0_71)), _i0_71)
				}
			}
			_2505_v53 = _nw398
			var _2508_v54 _dafny.Array
			_ = _2508_v54
			var _len0_72 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_72
			var _nw399 _dafny.Array
			_ = _nw399
			if _len0_72.Cmp(_dafny.Zero) == 0 {
				_nw399 = _dafny.NewArray(_len0_72)
			} else {
				var _init72 func(_dafny.Int) _dafny.CodePoint = func(_2509_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				}
				_ = _init72
				var _element0_72 = _init72(_dafny.Zero)
				_ = _element0_72
				_nw399 = _dafny.NewArrayFromExample(_element0_72, nil, _len0_72)
				(_nw399).ArraySet1CodePoint(_element0_72, 0)
				var _nativeLen0_72 = (_len0_72).Int()
				_ = _nativeLen0_72
				for _i0_72 := 1; _i0_72 < _nativeLen0_72; _i0_72++ {
					(_nw399).ArraySet1CodePoint(_init72(_dafny.IntOf(_i0_72)), _i0_72)
				}
			}
			_2508_v54 = _nw399
			var _2510_v55 D1
			_ = _2510_v55
			_2510_v55 = Companion_D1_.Create_DC4_(_2508_v54, p1, (_2448_v4).Cardinality(), _2444_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-896))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg123 _dafny.Int) interface{} {
					return coer123(arg123)
				}
			}(func(_2511_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})))
			var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_2505_v53), 0))
			_ = _index360
			(_2505_v53).ArraySet1(Companion_D6_.Create_DC16_((_2510_v55).Dtor_cf11()), (_index360).Int())
			var _2512_v56 _dafny.Array
			_ = _2512_v56
			var _len0_73 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_73
			var _nw400 _dafny.Array
			_ = _nw400
			if _len0_73.Cmp(_dafny.Zero) == 0 {
				_nw400 = _dafny.NewArray(_len0_73)
			} else {
				var _init73 func(_dafny.Int) _dafny.Sequence = (func(_2513_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2514_i10 _dafny.Int) _dafny.Sequence {
						return _2513_v0
					}
				})(_2444_v0)
				_ = _init73
				var _element0_73 = _init73(_dafny.Zero)
				_ = _element0_73
				_nw400 = _dafny.NewArrayFromExample(_element0_73, nil, _len0_73)
				(_nw400).ArraySet1(_element0_73, 0)
				var _nativeLen0_73 = (_len0_73).Int()
				_ = _nativeLen0_73
				for _i0_73 := 1; _i0_73 < _nativeLen0_73; _i0_73++ {
					(_nw400).ArraySet1(_init73(_dafny.IntOf(_i0_73)), _i0_73)
				}
			}
			_2512_v56 = _nw400
			var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))
			_ = _index361
			(_2512_v56).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(704))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg124 _dafny.Int) interface{} {
					return coer124(arg124)
				}
			}(func(_2515_i11 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), (_index361).Int())
			var _2516_v57 D21
			_ = _2516_v57
			_2516_v57 = Companion_D21_.Create_DC50_(_this.F30, (_dafny.Zero).Minus((_this).F31()), _this.F30, (_this).F31())
			var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_2505_v53), 0))
			_ = _index362
			var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))
			_ = _index363
			var _rhs369 *C5 = _2504_v52
			_ = _rhs369
			var _rhs370 bool = (p1).Cmp((p1).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-898))))) == 0
			_ = _rhs370
			var _rhs371 bool = _this.F30
			_ = _rhs371
			var _rhs372 D6 = Companion_Default___.Fm38((_2516_v57).Dtor_cf89(), (_dafny.Zero).Minus((_this).Fm1(globalState)), _2444_v0, globalState)
			_ = _rhs372
			var _rhs373 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(632))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg125 _dafny.Int) interface{} {
					return coer125(arg125)
				}
			}(func(_2517_i12 _dafny.Int) _dafny.CodePoint {
				return (func() _dafny.CodePoint {
					if false {
						return _dafny.CodePoint('f')
					}
					return _dafny.CodePoint('e')
				})()
			}))
			_ = _rhs373
			var _lhs341 *GlobalState = globalState
			_ = _lhs341
			var _lhs342 *GlobalState = globalState
			_ = _lhs342
			var _lhs343 _dafny.Array = _2505_v53
			_ = _lhs343
			var _lhs344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_2505_v53), 0))
			_ = _lhs344
			var _lhs345 _dafny.Array = _2512_v56
			_ = _lhs345
			var _lhs346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))
			_ = _lhs346
			_2504_v52 = _rhs369
			_lhs341.F2 = _rhs370
			_lhs342.F22 = _rhs371
			(_lhs343).ArraySet1(_rhs372, (_lhs344).Int())
			(_lhs345).ArraySet1(_rhs373, (_lhs346).Int())
			var _2518_v58 _dafny.Set
			_ = _2518_v58
			_2518_v58 = _dafny.SetOf((_this).F26())
			var _2519_v59 _dafny.Sequence
			_ = _2519_v59
			_2519_v59 = _dafny.SeqOf(_2518_v58, _2518_v58)
			var _2520_v60 _dafny.Map
			_ = _2520_v60
			_2520_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2519_v59, _2444_v0)
			(globalState).F6 = (func() _dafny.Sequence {
				if (_2520_v60).Contains(_2519_v59) {
					return (_2520_v60).Get(_2519_v59).(_dafny.Sequence)
				}
				return _2444_v0
			})()
			r0 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ofrmiv"), (_2512_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))).Int()).(_dafny.Sequence))
			(globalState).F16 = (_this).F31()
			var _2521_v61 _dafny.Set
			_ = _2521_v61
			_2521_v61 = _dafny.SetOf(Companion_Default___.Fm22(p0, _this.F30, globalState), (_2512_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))).Int()).(_dafny.Sequence), (func() _dafny.Sequence {
				if p2 {
					return (_2512_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))).Int()).(_dafny.Sequence)
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("x")
			})())
			var _2522_v62 _dafny.Map
			_ = _2522_v62
			_2522_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2512_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))).Int()).(_dafny.Sequence), p0)
			var _rhs374 _dafny.Set = (_dafny.SetOf((_2512_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_2512_v56), 0))).Int()).(_dafny.Sequence))).Union(func() _dafny.Set {
				var _coll73 = _dafny.NewBuilder()
				_ = _coll73
				for _iter76 := _dafny.Iterate((_2522_v62).Keys().Elements()); ; {
					_compr_73, _ok76 := _iter76()
					if !_ok76 {
						break
					}
					var _2523_v63 _dafny.Sequence
					_2523_v63 = interface{}(_compr_73).(_dafny.Sequence)
					if (_2522_v62).Contains(_2523_v63) {
						_coll73.Add(_2523_v63)
					}
				}
				return _coll73.ToSet()
			}())
			_ = _rhs374
			var _rhs375 bool = (_this).F26()
			_ = _rhs375
			var _lhs347 *GlobalState = globalState
			_ = _lhs347
			_2521_v61 = _rhs374
			_lhs347.F2 = _rhs375
		}
		r0 = Companion_Default___.Fm22(p1, !((p2) == (_this.F30)), globalState)
		r1 = (_dafny.Zero).Minus((_this).Fm4(Companion_D0_.Create_DC0_((_this).F27()), (p0).Plus(p1), p1, globalState))
		r2 = p2
		return r0, r1, r2
	}
}
func (_this *C9) M6(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.CodePoint, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _2524_v0 _dafny.Sequence
		_ = _2524_v0
		_2524_v0 = _dafny.SeqOf((_this).F27(), (_this).F26())
		var _hi21 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2524_v0, _2524_v0)).Cardinality())
		_ = _hi21
		for _2525_i0 := p1; _2525_i0.Cmp(_hi21) < 0; _2525_i0 = _2525_i0.Plus(_dafny.One) {
			var _2526_v1 _dafny.Array
			_ = _2526_v1
			var _len0_74 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_74
			var _nw401 _dafny.Array
			_ = _nw401
			if _len0_74.Cmp(_dafny.Zero) == 0 {
				_nw401 = _dafny.NewArray(_len0_74)
			} else {
				var _init74 func(_dafny.Int) bool = (func(_2527_v0 _dafny.Sequence, _2528_i0 _dafny.Int, _2529_p1 _dafny.Int, _2530_p3 _dafny.Int) func(_dafny.Int) bool {
					return func(_2531_i1 _dafny.Int) bool {
						return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.IntOfUint32((_2527_v0).Cardinality()), _2528_i0), _dafny.SeqOf(_2529_p1, (_this).F31(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer126 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg126 _dafny.Int) interface{} {
								return coer126(arg126)
							}
						}((func(_2532_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2533_i2 _dafny.Int) _dafny.Int {
								return _2532_p3
							}
						})(_2530_p3)))).Cardinality())))
					}
				})(_2524_v0, _2525_i0, p1, p3)
				_ = _init74
				var _element0_74 = _init74(_dafny.Zero)
				_ = _element0_74
				_nw401 = _dafny.NewArrayFromExample(_element0_74, nil, _len0_74)
				(_nw401).ArraySet1(_element0_74, 0)
				var _nativeLen0_74 = (_len0_74).Int()
				_ = _nativeLen0_74
				for _i0_74 := 1; _i0_74 < _nativeLen0_74; _i0_74++ {
					(_nw401).ArraySet1(_init74(_dafny.IntOf(_i0_74)), _i0_74)
				}
			}
			_2526_v1 = _nw401
			var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))
			_ = _index364
			(_2526_v1).ArraySet1((func() bool {
				if !((_this).F26()) {
					return _this.F30
				}
				return (_this).F26()
			})(), (_index364).Int())
			var _2534_v2 _dafny.Map
			_ = _2534_v2
			_2534_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !((_this).F26()))
			var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))
			_ = _index365
			(_2526_v1).ArraySet1((func() bool {
				if (_2534_v2).Contains(p3) {
					return (_2534_v2).Get(p3).(bool)
				}
				return true
			})(), (_index365).Int())
			var _2535_v3 D1
			_ = _2535_v3
			_2535_v3 = Companion_D1_.Create_DC3_(_dafny.SeqOf(!(_this.F30)))
			var _2536_v5 _dafny.Sequence
			_ = _2536_v5
			_2536_v5 = _dafny.SeqOf(func() _dafny.Set {
				var _coll74 = _dafny.NewBuilder()
				_ = _coll74
				for _iter77 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(435), _dafny.IntOfInt64(179))); ; {
					_compr_74, _ok77 := _iter77()
					if !_ok77 {
						break
					}
					var _2537_v4 _dafny.Int
					_2537_v4 = interface{}(_compr_74).(_dafny.Int)
					if ((_dafny.IntOfInt64(435)).Cmp(_2537_v4) <= 0) && ((_2537_v4).Cmp(_dafny.IntOfInt64(179)) < 0) {
						_coll74.Add(Companion_Default___.SafeModuloInt(_2537_v4, (p0).Cardinality()))
					}
				}
				return _coll74.ToSet()
			}())
			var _2538_v6 _dafny.Set
			_ = _2538_v6
			_2538_v6 = _dafny.SetOf(_dafny.IntOfInt64(545), _dafny.IntOfInt64(-539), (_this).Fm1(globalState), _dafny.IntOfInt64(747), p3)
			var _2539_v7 _dafny.Set
			_ = _2539_v7
			_2539_v7 = _dafny.SetOf((_this).F26(), true, (_this).Fm2((_2526_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))).Int()).(bool), _2525_i0, globalState))
			var _2540_v8 D6
			_ = _2540_v8
			_2540_v8 = Companion_D6_.Create_DC17_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_2526_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))).Int()).(bool), (_this).F27()), (_this).F31())).Cardinality(), (_this).F31(), p3, _dafny.IntOfInt64(-624), p1)
			var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))
			_ = _index366
			var _rhs376 bool = !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm55(globalState), _dafny.Companion_Sequence_.Update(_2536_v5, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_2536_v5).Cardinality()))).Uint32(), _2538_v6))
			_ = _rhs376
			var _rhs377 bool = (((_2539_v7).Cardinality()).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_2526_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))).Int()).(bool))).Cardinality()))).Cmp((_this).F31()) < 0
			_ = _rhs377
			var _rhs378 bool = _this.F30
			_ = _rhs378
			var _rhs379 _dafny.Int = (_2540_v8).Dtor_cf34()
			_ = _rhs379
			var _rhs380 D1 = _2535_v3
			_ = _rhs380
			var _lhs348 *GlobalState = globalState
			_ = _lhs348
			var _lhs349 *GlobalState = globalState
			_ = _lhs349
			var _lhs350 _dafny.Array = _2526_v1
			_ = _lhs350
			var _lhs351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(908), _dafny.ArrayLen((_2526_v1), 0))
			_ = _lhs351
			var _lhs352 *GlobalState = globalState
			_ = _lhs352
			_lhs348.F24 = _rhs376
			_lhs349.F15 = _rhs377
			(_lhs350).ArraySet1(_rhs378, (_lhs351).Int())
			_lhs352.F12 = _rhs379
			_2535_v3 = _rhs380
			(globalState).F12 = Companion_Default___.Fm39(globalState)
			(globalState).F12 = _2525_i0
		}
		var _2541_v9 _dafny.Map
		_ = _2541_v9
		_2541_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !((_this).Fm2(_this.F30, (_this).F31(), globalState)))
		var _2542_v10 _dafny.Map
		_ = _2542_v10
		_2542_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2541_v9).Cardinality(), !((_this).F26()))
		var _2543_v11 _dafny.Sequence
		_ = _2543_v11
		_2543_v11 = _dafny.UnicodeSeqOfUtf8Bytes("tvcq")
		var _2544_v12 _dafny.Map
		_ = _2544_v12
		_2544_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() bool {
			if (_2541_v9).Contains(_dafny.IntOfInt64(-801)) {
				return (_2541_v9).Get(_dafny.IntOfInt64(-801)).(bool)
			}
			return (_this).Fm2(_this.F30, p1, globalState)
		})()) || (_this.F30), _2543_v11)
		(globalState).F6 = (func() _dafny.Sequence {
			if (_2544_v12).Contains((_this).F27()) {
				return (_2544_v12).Get((_this).F27()).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qntndfgm"), _dafny.UnicodeSeqOfUtf8Bytes("jkyhfrdrq"))
		})()
		var _2545_v13 D1
		_ = _2545_v13
		_2545_v13 = Companion_D1_.Create_DC3_(_2524_v0)
		if func(_source35 D1) bool {
			if _source35.Is_DC4() {
				var _2546___mcc_h0 _dafny.Array = _source35.Get_().(D1_DC4).Cf7
				_ = _2546___mcc_h0
				var _2547___mcc_h1 _dafny.Int = _source35.Get_().(D1_DC4).Cf8
				_ = _2547___mcc_h1
				var _2548___mcc_h2 _dafny.Int = _source35.Get_().(D1_DC4).Cf9
				_ = _2548___mcc_h2
				var _2549___mcc_h3 _dafny.Sequence = _source35.Get_().(D1_DC4).Cf10
				_ = _2549___mcc_h3
				var _2550___mcc_h4 _dafny.Sequence = _source35.Get_().(D1_DC4).Cf11
				_ = _2550___mcc_h4
				var _2551_cf11 _dafny.Sequence = _2550___mcc_h4
				_ = _2551_cf11
				var _2552_cf10 _dafny.Sequence = _2549___mcc_h3
				_ = _2552_cf10
				var _2553_cf9 _dafny.Int = _2548___mcc_h2
				_ = _2553_cf9
				var _2554_cf8 _dafny.Int = _2547___mcc_h1
				_ = _2554_cf8
				var _2555_cf7 _dafny.Array = _2546___mcc_h0
				_ = _2555_cf7
				return ((_this).F27()) == (_this.F30)
			} else if _source35.Is_DC5() {
				var _2556___mcc_h5 _dafny.Int = _source35.Get_().(D1_DC5).Cf12
				_ = _2556___mcc_h5
				var _2557_cf12 _dafny.Int = _2556___mcc_h5
				_ = _2557_cf12
				return (_this).F27()
			} else {
				var _2558___mcc_h6 _dafny.Sequence = _source35.Get_().(D1_DC3).Cf6
				_ = _2558___mcc_h6
				var _2559_cf6 _dafny.Sequence = _2558___mcc_h6
				_ = _2559_cf6
				return (_this).F26()
			}
		}(_2545_v13) {
			if !(((p3).Times(_dafny.IntOfInt64(737))).Cmp((_dafny.Zero).Minus(p1)) > 0) {
				var _2560_v14 _dafny.Array
				_ = _2560_v14
				var _len0_75 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_75
				var _nw402 _dafny.Array
				_ = _nw402
				if _len0_75.Cmp(_dafny.Zero) == 0 {
					_nw402 = _dafny.NewArray(_len0_75)
				} else {
					var _init75 func(_dafny.Int) bool = func(_2561_i3 _dafny.Int) bool {
						return _this.F30
					}
					_ = _init75
					var _element0_75 = _init75(_dafny.Zero)
					_ = _element0_75
					_nw402 = _dafny.NewArrayFromExample(_element0_75, nil, _len0_75)
					(_nw402).ArraySet1(_element0_75, 0)
					var _nativeLen0_75 = (_len0_75).Int()
					_ = _nativeLen0_75
					for _i0_75 := 1; _i0_75 < _nativeLen0_75; _i0_75++ {
						(_nw402).ArraySet1(_init75(_dafny.IntOf(_i0_75)), _i0_75)
					}
				}
				_2560_v14 = _nw402
				var _2562_v15 _dafny.Array
				_ = _2562_v15
				_2562_v15 = _2560_v14
				_2562_v15 = _2562_v15
				(globalState).F15 = !(((_this).F26()) && (_this.F30)) || (((_this).F27()) && (_this.F30))
				var _2563_v16 _dafny.Map
				_ = _2563_v16
				_2563_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-67), p3), p2)
				_2563_v16 = (_2563_v16).Update(p1, p2)
				var _2564_v17 _dafny.MultiSet
				_ = _2564_v17
				_2564_v17 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(p3, (_this).F31()), p1)
				_2564_v17 = (_this).F29()
				var _rhs381 _dafny.Sequence = _2543_v11
				_ = _rhs381
				var _rhs382 _dafny.Sequence = _2543_v11
				_ = _rhs382
				var _rhs383 _dafny.Int = (_this).F31()
				_ = _rhs383
				var _rhs384 _dafny.CodePoint = p2
				_ = _rhs384
				var _rhs385 _dafny.Int = p3
				_ = _rhs385
				var _lhs353 *GlobalState = globalState
				_ = _lhs353
				var _lhs354 *GlobalState = globalState
				_ = _lhs354
				var _lhs355 *GlobalState = globalState
				_ = _lhs355
				var _lhs356 *GlobalState = globalState
				_ = _lhs356
				var _lhs357 *GlobalState = globalState
				_ = _lhs357
				_lhs353.F6 = _rhs381
				_lhs354.F6 = _rhs382
				_lhs355.F16 = _rhs383
				_lhs356.F19 = _rhs384
				_lhs357.F12 = _rhs385
			} else {
				var _2565_v18 _dafny.MultiSet
				_ = _2565_v18
				_2565_v18 = _dafny.MultiSetOf(_dafny.SeqOf(p1))
				var _2566_v19 *C3
				_ = _2566_v19
				var _nw403 *C3 = New_C3_()
				_ = _nw403
				_nw403.Ctor__(_2565_v18, (_this).F26(), _this.F30)
				_2566_v19 = _nw403
				var _2567_v20 _dafny.Map
				_ = _2567_v20
				_2567_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2566_v19)
				var _2568_v21 _dafny.Set
				_ = _2568_v21
				_2568_v21 = _dafny.SetOf(_2567_v20, _2567_v20, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2566_v19)).Merge(_2567_v20), _2567_v20)
				var _2569_v22 _dafny.Array
				_ = _2569_v22
				var _nw404 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw404
				_2569_v22 = _nw404
				var _2570_v23 _dafny.Array
				_ = _2570_v23
				var _nw405 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw405
				_2570_v23 = _nw405
				var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_2569_v22), 0))
				_ = _index367
				(_2569_v22).ArraySet1(_2570_v23, (_index367).Int())
				var _2571_v24 _dafny.Sequence
				_ = _2571_v24
				_2571_v24 = _dafny.SeqOf(p0, p0)
				var _2572_v25 D10
				_ = _2572_v25
				_2572_v25 = Companion_D10_.Create_DC24_(_2543_v11, _2571_v24, _this.F30)
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_2569_v22), 0))
				_ = _index368
				var _rhs386 _dafny.Set = _2568_v21
				_ = _rhs386
				var _rhs387 _dafny.Int = (Companion_Default___.SafeDivisionInt((_this).F31(), (_this).F31())).Minus((p3).Times(p1))
				_ = _rhs387
				var _rhs388 bool = _dafny.Companion_Sequence_.Contains((_2572_v25).Dtor_cf44(), p2)
				_ = _rhs388
				var _rhs389 bool = (_this).F26()
				_ = _rhs389
				var _rhs390 _dafny.Array = _2570_v23
				_ = _rhs390
				var _lhs358 *GlobalState = globalState
				_ = _lhs358
				var _lhs359 *GlobalState = globalState
				_ = _lhs359
				var _lhs360 *GlobalState = globalState
				_ = _lhs360
				var _lhs361 _dafny.Array = _2569_v22
				_ = _lhs361
				var _lhs362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_2569_v22), 0))
				_ = _lhs362
				_2568_v21 = _rhs386
				_lhs358.F12 = _rhs387
				_lhs359.F15 = _rhs388
				_lhs360.F15 = _rhs389
				(_lhs361).ArraySet1(_rhs390, (_lhs362).Int())
				var _2573_v26 _dafny.Map
				_ = _2573_v26
				_2573_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F27())
				_2573_v26 = (_2573_v26).Update(p2, _this.F30)
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_2569_v22), 0))
				_ = _index369
				var _rhs391 _dafny.Int = (_this).F31()
				_ = _rhs391
				var _rhs392 bool = (_this).F26()
				_ = _rhs392
				var _rhs393 _dafny.Array = _2570_v23
				_ = _rhs393
				var _lhs363 *GlobalState = globalState
				_ = _lhs363
				var _lhs364 *GlobalState = globalState
				_ = _lhs364
				var _lhs365 _dafny.Array = _2569_v22
				_ = _lhs365
				var _lhs366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(780), _dafny.ArrayLen((_2569_v22), 0))
				_ = _lhs366
				_lhs363.F12 = _rhs391
				_lhs364.F2 = _rhs392
				(_lhs365).ArraySet1(_rhs393, (_lhs366).Int())
				(globalState).F15 = !(false)
				(globalState).F16 = p3
			}
			var _2574_v27 _dafny.Array
			_ = _2574_v27
			var _nw406 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
			_ = _nw406
			_2574_v27 = _nw406
			_2574_v27 = _2574_v27
			var _2575_v28 _dafny.Map
			_ = _2575_v28
			_2575_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2543_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(932))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg127 _dafny.Int) interface{} {
					return coer127(arg127)
				}
			}((func(_2576_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2577_i4 _dafny.Int) _dafny.CodePoint {
					return _2576_p2
				}
			})(p2))))
			(globalState).F12 = ((_2575_v28).Merge((_2575_v28).Update(_2543_v11, _2543_v11))).Cardinality()
			var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_2574_v27), 0))
			_ = _index370
			(_2574_v27).ArraySet1(false, (_index370).Int())
			var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_2574_v27), 0))
			_ = _index371
			(_2574_v27).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ioovj"), _2543_v11), (_index371).Int())
			var _2578_v29 _dafny.Sequence
			_ = _2578_v29
			_2578_v29 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F31()), (_this).F31(), p1)
			var _2579_v30 _dafny.Map
			_ = _2579_v30
			_2579_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2578_v29, p3)
			var _2580_v32 _dafny.Sequence
			_ = _2580_v32
			_2580_v32 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer128 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg128 _dafny.Int) interface{} {
					return coer128(arg128)
				}
			}((func(_2581_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2582_i5 _dafny.Int) _dafny.Int {
					return _2581_p1
				}
			})(p1)))
			var _2583_v33 _dafny.Sequence
			_ = _2583_v33
			_2583_v33 = _dafny.SeqOf(_2578_v29, (_2580_v32))
			_2579_v30 = func() _dafny.Map {
				var _coll75 = _dafny.NewMapBuilder()
				_ = _coll75
				for _iter78 := _dafny.Iterate((_2583_v33).Elements()); ; {
					_compr_75, _ok78 := _iter78()
					if !_ok78 {
						break
					}
					var _2584_v31 _dafny.Sequence
					_2584_v31 = interface{}(_compr_75).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_2583_v33, _2584_v31) {
						_coll75.Add(_2584_v31, p3)
					}
				}
				return _coll75.ToMap()
			}()
		} else {
			(globalState).F24 = (_this).F27()
			var _2585_v34 _dafny.MultiSet
			_ = _2585_v34
			_2585_v34 = _dafny.MultiSetOf(_this.F30)
			var _2586_v35 _dafny.Sequence
			_ = _2586_v35
			_2586_v35 = _dafny.SeqOf(p1, _dafny.IntOfInt64(72))
			var _2587_v36 _dafny.Map
			_ = _2587_v36
			_2587_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2586_v35).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_2586_v35).Cardinality()))).Uint32()).(_dafny.Int), p0)
			var _2588_v37 T0
			_ = _2588_v37
			var _nw407 *C7 = New_C7_()
			_ = _nw407
			_nw407.Ctor__(true, _2587_v36)
			_2588_v37 = _nw407
			var _2589_v38 _dafny.Set
			_ = _2589_v38
			_2589_v38 = _dafny.SetOf(p1)
			var _rhs394 _dafny.MultiSet = (_2585_v34).Update((_this).F26(), Companion_Default___.Abs((_2589_v38).Cardinality()))
			_ = _rhs394
			var _rhs395 bool = _dafny.Companion_Sequence_.IsPrefixOf(_2586_v35, _2586_v35)
			_ = _rhs395
			var _rhs396 T0 = _2588_v37
			_ = _rhs396
			var _lhs367 *GlobalState = globalState
			_ = _lhs367
			_2585_v34 = _rhs394
			_lhs367.F22 = _rhs395
			_2588_v37 = _rhs396
			var _2590_v39 _dafny.Sequence
			_ = _2590_v39
			_2590_v39 = _dafny.SeqOf(p0)
			var _2591_v40 _dafny.Map
			_ = _2591_v40
			_2591_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, ((_dafny.Zero).Minus(p3)).Minus((Companion_D21_.Create_DC52_((_this).F31(), _2590_v39)).Dtor_cf95()))
			var _2592_v41 D23
			_ = _2592_v41
			_2592_v41 = Companion_D23_.Create_DC57_()
			var _2593_v42 _dafny.Map
			_ = _2593_v42
			_2593_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D23_.Create_DC58_(_2592_v41), (_this).F27())
			var _2594_v43 D23
			_ = _2594_v43
			_2594_v43 = Companion_D23_.Create_DC58_(_2592_v41)
			var _2595_v44 D10
			_ = _2595_v44
			_2595_v44 = Companion_D10_.Create_DC24_(_2543_v11, _2590_v39, false)
			_2591_v40 = (_2591_v40).Update((func() bool {
				if (_this).F26() {
					return (_this).F27()
				}
				return (func() bool {
					if (_2593_v42).Contains(_2594_v43) {
						return (_2593_v42).Get(_2594_v43).(bool)
					}
					return (_this).F26()
				})()
			})(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_2543_v11).Cardinality()), _dafny.IntOfUint32((_2543_v11).Cardinality()), _dafny.IntOfUint32(((_2595_v44).Dtor_cf44()).Cardinality()), p1, _dafny.IntOfUint32((_2524_v0).Cardinality()))).Cardinality())
			(globalState).F2 = (_2589_v38).IsProperSubsetOf(_dafny.SetOf((_this).F31(), _dafny.IntOfInt64(-14), (_this).F31(), _dafny.IntOfInt64(-503), (_this).F31()))
			(globalState).F16 = (_this).F31()
		}
		var _2596_v45 *C8
		_ = _2596_v45
		var _nw408 *C8 = New_C8_()
		_ = _nw408
		_nw408.Ctor__(_dafny.MultiSetOf(p3))
		_2596_v45 = _nw408
		var _2597_v46 _dafny.Array
		_ = _2597_v46
		var _len0_76 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_76
		var _nw409 _dafny.Array
		_ = _nw409
		if _len0_76.Cmp(_dafny.Zero) == 0 {
			_nw409 = _dafny.NewArray(_len0_76)
		} else {
			var _init76 func(_dafny.Int) bool = func(_2598_i7 _dafny.Int) bool {
				return _this.F30
			}
			_ = _init76
			var _element0_76 = _init76(_dafny.Zero)
			_ = _element0_76
			_nw409 = _dafny.NewArrayFromExample(_element0_76, nil, _len0_76)
			(_nw409).ArraySet1(_element0_76, 0)
			var _nativeLen0_76 = (_len0_76).Int()
			_ = _nativeLen0_76
			for _i0_76 := 1; _i0_76 < _nativeLen0_76; _i0_76++ {
				(_nw409).ArraySet1(_init76(_dafny.IntOf(_i0_76)), _i0_76)
			}
		}
		_2597_v46 = _nw409
		for _iter79 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2597_v46), 0))); ; {
			_guard_loop_3, _ok79 := _iter79()
			if !_ok79 {
				break
			}
			var _2599_i6 _dafny.Int
			_2599_i6 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_2599_i6).Sign() != -1) && ((_2599_i6).Cmp(_dafny.ArrayLen((_2597_v46), 0)) < 0)) {
				(_2597_v46).ArraySet1((_this.F30) && ((func() bool {
					if (_2542_v10).Contains(_dafny.IntOfUint32((_2543_v11).Cardinality())) {
						return (_2542_v10).Get(_dafny.IntOfUint32((_2543_v11).Cardinality())).(bool)
					}
					return (_this).F26()
				})()), (_2599_i6).Int())
			}
		}
		var _2600_v47 _dafny.Sequence
		_ = _2600_v47
		_2600_v47 = _dafny.SeqOf(p3)
		var _2601_v48 _dafny.Set
		_ = _2601_v48
		_2601_v48 = _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(922), (_this).Fm1(globalState)), _dafny.Companion_Sequence_.Concatenate(_2600_v47, _2600_v47))
		_2601_v48 = _2601_v48
		r0 = p3
		r1 = (_this).Fm1(globalState)
		r2 = !((p1).Cmp(p3) == 0)
		return r0, r1, r2
	}
}
func (_this *C9) F31() _dafny.Int {
	{
		return _this._f31
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
