import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(p0, p1, globalState):
        return default__.safeDivisionInt(len(_dafny.Map({(0) - (-836): _dafny.CodePoint('b')})), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmxs"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrusd")))))

    @staticmethod
    def fm1(p0, globalState):
        return ((_dafny.Map({911: _dafny.MultiSet([-611, 2])})) | (_dafny.Map({(D0_DC2(780, _dafny.MultiSet([160]), True)).cf6: _dafny.MultiSet([253])}))) | (_dafny.Map({235: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([True])))})): _dafny.CodePoint('b')}))]))}))

    @staticmethod
    def fm2(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgrolje"))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        source0_ = D2_DC10()
        if source0_.is_DC10:
            return not(False)
        elif source0_.is_DC11:
            d_0___mcc_h0_ = source0_.cf17
            d_1___mcc_h1_ = source0_.cf18
            d_2_cf18_ = d_1___mcc_h1_
            d_3_cf17_ = d_0___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([30])) < (_dafny.SeqWithoutIsStrInference([-447]))
        elif True:
            d_4___mcc_h2_ = source0_.cf16
            d_5_cf16_ = d_4___mcc_h2_
            return (not((D0_DC2(len(d_5_cf16_), _dafny.MultiSet([-769]), True)).cf8)) == (False)

    @staticmethod
    def fm5(globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if True else _dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return ((_dafny.Set({False, True}) if False else _dafny.Set({not(False), False, False}))).intersection((_dafny.Set({False, True}) if False else _dafny.Set({False})))

    @staticmethod
    def fm7(globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([True, (7) < (211), (False if True else True)])

    @staticmethod
    def fm9(p0, globalState):
        source1_ = D1_DC6(_dafny.Set({False}))
        if source1_.is_DC5:
            d_6___mcc_h0_ = source1_.cf11
            d_7_cf11_ = d_6___mcc_h0_
            return (_dafny.Map({d_7_cf11_: d_7_cf11_})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtisl"))): d_7_cf11_}))
        elif source1_.is_DC6:
            d_8___mcc_h1_ = source1_.cf12
            d_9_cf12_ = d_8___mcc_h1_
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wf"))): 589})
        elif source1_.is_DC7:
            d_10___mcc_h2_ = source1_.cf13
            d_11___mcc_h3_ = source1_.cf14
            d_12_cf14_ = d_11___mcc_h3_
            d_13_cf13_ = d_10___mcc_h2_
            return (_dafny.Map({d_12_cf14_: d_13_cf13_})) | (_dafny.Map({d_13_cf13_: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))).cardinality}))
        elif source1_.is_DC4:
            d_14___mcc_h4_ = source1_.cf10
            d_15_cf10_ = d_14___mcc_h4_
            return (_dafny.Map({len((D5_DC16(_dafny.Map({_dafny.SeqWithoutIsStrInference([262]): False}))).cf25): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv"))), 521]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.MultiSet([-972]): len(_dafny.Map({-630: 558}))})), 78])): 682}))
        elif True:
            d_16___mcc_h5_ = source1_.cf15
            d_17_cf15_ = d_16___mcc_h5_
            return _dafny.Map({len(_dafny.Map({not(not(True)): False})): 592})

    @staticmethod
    def fm10(p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttrqiyhpy"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twandbl"))):
            return D1_DC6(_dafny.Set({False}))
        elif True:
            return D1_DC6(_dafny.Set({False, not(True), True}))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwuiwydjh")))])) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwoc"))), len(_dafny.Set({not(False)}))])) | (_dafny.MultiSet([(0) - ((_dafny.MultiSet([66, 605])).cardinality)])))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        d_18_v0_: int
        d_18_v0_ = 326
        d_19_v1_: _dafny.Seq
        d_19_v1_ = _dafny.SeqWithoutIsStrInference([d_18_v0_, d_18_v0_])
        r0 = default__.safeDivisionInt(default__.safeDivisionInt(222, d_18_v0_), default__.safeDivisionInt(len(d_19_v1_), d_18_v0_))
        d_20_v2_: D2
        d_20_v2_ = D2_DC10()
        source2_ = d_20_v2_
        if source2_.is_DC10:
            d_21_v3_: _dafny.Array
            nw0_ = _dafny.Array(D0.default()(), 10)
            d_21_v3_ = nw0_
            d_22_v4_: _dafny.Seq
            d_22_v4_ = _dafny.SeqWithoutIsStrInference([d_21_v3_, d_21_v3_, d_21_v3_])
            d_23_v5_: _dafny.Array
            nw1_ = _dafny.Array(None, 27)
            nw1_[int(0)] = d_21_v3_
            nw1_[int(1)] = d_21_v3_
            nw1_[int(2)] = d_21_v3_
            nw1_[int(3)] = d_21_v3_
            nw1_[int(4)] = d_21_v3_
            nw1_[int(5)] = d_21_v3_
            nw1_[int(6)] = d_21_v3_
            nw1_[int(7)] = d_21_v3_
            nw1_[int(8)] = d_21_v3_
            nw1_[int(9)] = d_21_v3_
            nw1_[int(10)] = d_21_v3_
            nw1_[int(11)] = d_21_v3_
            nw1_[int(12)] = d_21_v3_
            nw1_[int(13)] = d_21_v3_
            nw1_[int(14)] = d_21_v3_
            nw1_[int(15)] = d_21_v3_
            nw1_[int(16)] = d_21_v3_
            nw1_[int(17)] = d_21_v3_
            nw1_[int(18)] = d_21_v3_
            nw1_[int(19)] = d_21_v3_
            nw1_[int(20)] = d_21_v3_
            nw1_[int(21)] = d_21_v3_
            nw1_[int(22)] = d_21_v3_
            nw1_[int(23)] = (d_22_v4_)[default__.safeIndex(d_18_v0_, len(d_22_v4_))]
            nw1_[int(24)] = d_21_v3_
            nw1_[int(25)] = d_21_v3_
            nw1_[int(26)] = d_21_v3_
            d_23_v5_ = nw1_
            index0_ = default__.safeIndex(430, (d_23_v5_).length(0))
            (d_23_v5_)[index0_] = d_21_v3_
            index1_ = default__.safeIndex(430, (d_23_v5_).length(0))
            (d_23_v5_)[index1_] = d_21_v3_
            d_24_v6_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.Set({}), 7)
            d_24_v6_ = nw2_
            d_25_v7_: _dafny.Seq
            d_25_v7_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            d_26_v8_: _dafny.Set
            d_26_v8_ = _dafny.Set({len(d_25_v7_)})
            index2_ = default__.safeIndex(530, (d_24_v6_).length(0))
            (d_24_v6_)[index2_] = (_dafny.Set({d_18_v0_})) | (d_26_v8_)
            d_27_v9_: str
            d_27_v9_ = _dafny.CodePoint('v')
            d_28_v10_: _dafny.Set
            d_28_v10_ = _dafny.Set({d_27_v9_, d_27_v9_})
            d_29_v11_: _dafny.Map
            d_29_v11_ = _dafny.Map({p1: len(d_25_v7_)})
            d_30_v12_: _dafny.Map
            d_30_v12_ = _dafny.Map({False: d_29_v11_})
            d_31_v13_: _dafny.Map
            d_31_v13_ = _dafny.Map({d_27_v9_: len(d_19_v1_)})
            d_32_v14_: _dafny.MultiSet
            d_32_v14_ = _dafny.MultiSet([len(p0), d_18_v0_, len(_dafny.Set({d_27_v9_})), len(d_31_v13_), d_18_v0_])
            d_33_v15_: _dafny.Array
            nw3_ = _dafny.Array(None, 24)
            nw3_[int(0)] = 56
            nw3_[int(1)] = d_18_v0_
            nw3_[int(2)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))) + ((p0).set(default__.safeIndex(d_18_v0_, len(p0)), d_27_v9_)))
            nw3_[int(3)] = d_18_v0_
            nw3_[int(4)] = len(d_25_v7_)
            nw3_[int(5)] = (0) - (d_18_v0_)
            nw3_[int(6)] = default__.safeDivisionInt(d_18_v0_, d_18_v0_)
            nw3_[int(7)] = len(d_26_v8_)
            nw3_[int(8)] = default__.safeModuloInt(d_18_v0_, d_18_v0_)
            nw3_[int(9)] = len(p0)
            nw3_[int(10)] = d_18_v0_
            nw3_[int(11)] = d_18_v0_
            nw3_[int(12)] = d_18_v0_
            nw3_[int(13)] = (len(d_25_v7_)) - (d_18_v0_)
            nw3_[int(14)] = d_18_v0_
            nw3_[int(15)] = len((d_28_v10_ if p1 else _dafny.Set({d_27_v9_})))
            nw3_[int(16)] = (d_18_v0_) * (d_18_v0_)
            nw3_[int(17)] = (len(((d_30_v12_)[p1] if (p1) in (d_30_v12_) else d_29_v11_))) + (d_18_v0_)
            nw3_[int(18)] = (d_32_v14_).cardinality
            nw3_[int(19)] = (0) - ((d_18_v0_) + (d_18_v0_))
            nw3_[int(20)] = d_18_v0_
            nw3_[int(21)] = d_18_v0_
            nw3_[int(22)] = (0) - (d_18_v0_)
            nw3_[int(23)] = d_18_v0_
            d_33_v15_ = nw3_
            index3_ = default__.safeIndex(838, (d_33_v15_).length(0))
            (d_33_v15_)[index3_] = d_18_v0_
            d_34_v16_: D0
            d_34_v16_ = D0_DC0(d_26_v8_, p0)
            index4_ = default__.safeIndex(530, (d_24_v6_).length(0))
            index5_ = default__.safeIndex(838, (d_33_v15_).length(0))
            def iife0_(_pat_let0_0):
                def iife1_(d_35_dt__update__tmp_h0_):
                    def iife2_(_pat_let1_0):
                        def iife3_(d_36_dt__update_hcf1_h0_):
                            return D0_DC0((d_35_dt__update__tmp_h0_).cf0, d_36_dt__update_hcf1_h0_)
                        return iife3_(_pat_let1_0)
                    return iife2_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqexxyx")))
                return iife1_(_pat_let0_0)
            rhs0_ = (iife0_(d_34_v16_)).cf0
            rhs1_ = (0) - (d_18_v0_)
            rhs2_ = d_27_v9_
            rhs3_ = (0) - ((d_32_v14_).cardinality)
            lhs0_ = d_24_v6_
            lhs1_ = default__.safeIndex(530, (d_24_v6_).length(0))
            lhs2_ = globalState
            lhs3_ = d_33_v15_
            lhs4_ = default__.safeIndex(838, (d_33_v15_).length(0))
            lhs0_[lhs1_] = rhs0_
            lhs2_.f16 = rhs1_
            d_27_v9_ = rhs2_
            lhs3_[lhs4_] = rhs3_
            (globalState).f16 = d_18_v0_
            d_37_v17_: D0
            d_37_v17_ = D0_DC2(len(_dafny.SeqWithoutIsStrInference([(d_33_v15_)[default__.safeIndex(838, (d_33_v15_).length(0))], 906, d_18_v0_, (d_33_v15_)[default__.safeIndex(838, (d_33_v15_).length(0))], (d_33_v15_)[default__.safeIndex(838, (d_33_v15_).length(0))]])), d_32_v14_, p1)
            (globalState).f10 = (0) - ((len(_dafny.Set({p1, p1}))) - ((d_18_v0_) * ((d_37_v17_).cf6)))
        elif source2_.is_DC11:
            d_38___mcc_h0_ = source2_.cf17
            d_39___mcc_h1_ = source2_.cf18
            d_40_cf18_ = d_39___mcc_h1_
            d_41_cf17_ = d_38___mcc_h0_
            d_42_v18_: D1
            d_42_v18_ = D1_DC5(d_18_v0_)
            source3_ = d_42_v18_
            if source3_.is_DC5:
                d_43___mcc_h3_ = source3_.cf11
                d_44_cf11_ = d_43___mcc_h3_
                d_45_v19_: C0
                nw4_ = C0()
                nw4_.ctor__(d_18_v0_, p1)
                d_45_v19_ = nw4_
                d_46_v20_: _dafny.Map
                d_46_v20_ = _dafny.Map({p1: p1})
                d_47_v21_: D1
                d_47_v21_ = D1_DC7(d_18_v0_, len(d_46_v20_))
                d_48_v22_: _dafny.Map
                d_48_v22_ = _dafny.Map({(_dafny.Map({d_44_cf11_: p1})).set(len(p0), p1): default__.fm3(p1, d_40_cf18_, len(p0), False, globalState)})
                d_49_v23_: _dafny.MultiSet
                d_49_v23_ = _dafny.MultiSet([(d_48_v22_) | (d_48_v22_), d_48_v22_])
                d_50_v24_: _dafny.Map
                d_50_v24_ = _dafny.Map({d_40_cf18_: p1})
                d_51_v25_: _dafny.Set
                d_51_v25_ = _dafny.Set({d_40_cf18_, len(d_50_v24_)})
                d_52_v26_: _dafny.Map
                d_52_v26_ = _dafny.Map({len(d_51_v25_): p1})
                rhs4_ = d_47_v21_
                rhs5_ = ((d_49_v23_)[_dafny.Map({d_50_v24_: p1})] if (_dafny.Map({d_50_v24_: p1})) in (d_49_v23_) else (d_18_v0_) * (888))
                lhs5_ = globalState
                d_47_v21_ = rhs4_
                lhs5_.f10 = rhs5_
                d_18_v0_ = (d_44_cf11_) * (d_18_v0_)
                d_53_v27_: _dafny.Seq
                d_53_v27_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_54_v28_: T0
                nw5_ = C0()
                nw5_.ctor__(d_40_cf18_, (d_53_v27_) == ((d_53_v27_).set(default__.safeIndex(d_18_v0_, len(d_53_v27_)), p1)))
                d_54_v28_ = nw5_
                d_54_v28_ = d_54_v28_
            elif source3_.is_DC6:
                d_55___mcc_h4_ = source3_.cf12
                d_56_cf12_ = d_55___mcc_h4_
                d_57_v29_: D1
                d_57_v29_ = D1_DC7(default__.safeDivisionInt(d_18_v0_, d_18_v0_), d_40_cf18_)
                pat_let_tv0_ = p0
                def iife4_(_pat_let2_0):
                    def iife5_(d_58_dt__update__tmp_h1_):
                        def iife6_(_pat_let3_0):
                            def iife7_(d_59_dt__update_hcf14_h0_):
                                return D1_DC7((d_58_dt__update__tmp_h1_).cf13, d_59_dt__update_hcf14_h0_)
                            return iife7_(_pat_let3_0)
                        return iife6_((0) - (len(pat_let_tv0_)))
                    return iife5_(_pat_let2_0)
                d_57_v29_ = iife4_(d_57_v29_)
                d_60_v30_: _dafny.MultiSet
                d_60_v30_ = _dafny.MultiSet([d_18_v0_, len(d_56_cf12_), d_18_v0_])
                d_61_v31_: _dafny.Seq
                d_61_v31_ = _dafny.SeqWithoutIsStrInference([d_60_v30_])
                d_62_v32_: _dafny.Map
                d_62_v32_ = _dafny.Map({d_40_cf18_: p1})
                d_63_v33_: _dafny.Seq
                d_63_v33_ = _dafny.SeqWithoutIsStrInference([p1])
                d_64_v34_: _dafny.Array
                nw6_ = _dafny.Array(None, 20)
                nw6_[int(0)] = p1
                nw6_[int(1)] = (d_40_cf18_) != (d_40_cf18_)
                nw6_[int(2)] = default__.fm3(p1, d_18_v0_, d_40_cf18_, p1, globalState)
                nw6_[int(3)] = (not(p1)) and (p1)
                nw6_[int(4)] = p1
                nw6_[int(5)] = p1
                nw6_[int(6)] = ((d_61_v31_)[default__.safeIndex(d_18_v0_, len(d_61_v31_))]).ispropersubset(_dafny.MultiSet([(0) - (d_18_v0_)]))
                nw6_[int(7)] = not(((d_62_v32_)[d_18_v0_] if (d_18_v0_) in (d_62_v32_) else p1))
                nw6_[int(8)] = p1
                nw6_[int(9)] = p1
                nw6_[int(10)] = p1
                nw6_[int(11)] = (d_63_v33_)[default__.safeIndex((D1_DC7(d_18_v0_, -298)).cf13, len(d_63_v33_))]
                nw6_[int(12)] = p1
                nw6_[int(13)] = default__.fm3(p1, d_18_v0_, len(d_19_v1_), p1, globalState)
                nw6_[int(14)] = (((d_62_v32_)[d_40_cf18_] if (d_40_cf18_) in (d_62_v32_) else p1)) and (p1)
                nw6_[int(15)] = p1
                nw6_[int(16)] = p1
                nw6_[int(17)] = (p0) == (p0)
                nw6_[int(18)] = not(False)
                nw6_[int(19)] = p1
                d_64_v34_ = nw6_
                d_65_v35_: D3
                d_65_v35_ = D3_DC12(d_64_v34_)
                d_64_v34_ = (D3_DC12((d_65_v35_).cf19)).cf19
                d_66_v36_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_66_v36_ = nw7_
                index6_ = default__.safeIndex(900, (d_64_v34_).length(0))
                (d_64_v34_)[index6_] = p1
                d_67_v37_: _dafny.Map
                d_67_v37_ = _dafny.Map({p1: p1})
                d_68_v38_: _dafny.MultiSet
                d_68_v38_ = _dafny.MultiSet([p1, p1])
                d_69_v39_: C0
                nw8_ = C0()
                nw8_.ctor__((d_68_v38_).cardinality, p1)
                d_69_v39_ = nw8_
                d_70_v40_: D0
                d_70_v40_ = D0_DC1(p1, d_67_v37_, d_69_v39_, len(d_63_v33_))
                pat_let_tv1_ = d_67_v37_
                index7_ = default__.safeIndex(900, (d_64_v34_).length(0))
                def iife8_(_pat_let4_0):
                    def iife9_(d_71_dt__update__tmp_h2_):
                        def iife10_(_pat_let5_0):
                            def iife11_(d_72_dt__update_hcf3_h0_):
                                return D0_DC1((d_71_dt__update__tmp_h2_).cf2, d_72_dt__update_hcf3_h0_, (d_71_dt__update__tmp_h2_).cf4, (d_71_dt__update__tmp_h2_).cf5)
                            return iife11_(_pat_let5_0)
                        return iife10_(pat_let_tv1_)
                    return iife9_(_pat_let4_0)
                rhs6_ = d_66_v36_
                rhs7_ = (iife8_(d_70_v40_)).cf5
                rhs8_ = not((d_62_v32_) == (_dafny.Map({d_40_cf18_: False})))
                rhs9_ = p1
                lhs6_ = globalState
                lhs7_ = d_64_v34_
                lhs8_ = default__.safeIndex(900, (d_64_v34_).length(0))
                lhs9_ = globalState
                d_66_v36_ = rhs6_
                lhs6_.f16 = rhs7_
                lhs7_[lhs8_] = rhs8_
                lhs9_.f4 = rhs9_
                d_73_v41_: C0
                nw9_ = C0()
                nw9_.ctor__(d_18_v0_, (p1) or (p1))
                d_73_v41_ = nw9_
            elif source3_.is_DC7:
                d_74___mcc_h5_ = source3_.cf13
                d_75___mcc_h6_ = source3_.cf14
                d_76_cf14_ = d_75___mcc_h6_
                d_77_cf13_ = d_74___mcc_h5_
                (globalState).f16 = d_77_cf13_
                d_78_v42_: C0
                nw10_ = C0()
                nw10_.ctor__(d_18_v0_, p1)
                d_78_v42_ = nw10_
                d_79_v43_: _dafny.Map
                d_79_v43_ = _dafny.Map({True: d_78_v42_})
                d_80_v44_: _dafny.Array
                nw11_ = _dafny.Array(None, 17)
                nw11_[int(0)] = (d_78_v42_ if p1 else d_78_v42_)
                nw11_[int(1)] = d_78_v42_
                nw11_[int(2)] = d_78_v42_
                nw11_[int(3)] = d_78_v42_
                nw11_[int(4)] = d_78_v42_
                nw11_[int(5)] = d_78_v42_
                nw11_[int(6)] = d_78_v42_
                nw11_[int(7)] = d_78_v42_
                nw11_[int(8)] = d_78_v42_
                nw11_[int(9)] = d_78_v42_
                nw11_[int(10)] = d_78_v42_
                nw11_[int(11)] = d_78_v42_
                nw11_[int(12)] = d_78_v42_
                nw11_[int(13)] = d_78_v42_
                nw11_[int(14)] = ((d_79_v43_)[p1] if (p1) in (d_79_v43_) else d_78_v42_)
                nw11_[int(15)] = d_78_v42_
                nw11_[int(16)] = d_78_v42_
                d_80_v44_ = nw11_
                index8_ = default__.safeIndex(716, (d_80_v44_).length(0))
                (d_80_v44_)[index8_] = d_78_v42_
                index9_ = default__.safeIndex(716, (d_80_v44_).length(0))
                (d_80_v44_)[index9_] = d_78_v42_
                d_78_v42_ = d_78_v42_
                d_81_v45_: D2
                d_81_v45_ = D2_DC11(d_41_cf17_, -243)
                d_82_v46_: _dafny.Map
                d_82_v46_ = _dafny.Map({d_78_v42_: (d_81_v45_).cf18})
                d_83_v47_: _dafny.Seq
                d_83_v47_ = _dafny.SeqWithoutIsStrInference([p0])
                d_82_v46_ = (d_82_v46_).set((d_80_v44_)[default__.safeIndex(716, (d_80_v44_).length(0))], len((p0) + ((d_83_v47_)[default__.safeIndex(d_40_cf18_, len(d_83_v47_))])))
            elif source3_.is_DC4:
                d_84___mcc_h7_ = source3_.cf10
                d_85_cf10_ = d_84___mcc_h7_
                (globalState).f4 = p1
                d_86_v48_: _dafny.Map
                d_86_v48_ = _dafny.Map({p1: p1})
                d_87_v49_: C0
                nw12_ = C0()
                nw12_.ctor__(d_40_cf18_, p1)
                d_87_v49_ = nw12_
                d_88_v50_: _dafny.Set
                d_88_v50_ = _dafny.Set({d_40_cf18_})
                d_89_v51_: D0
                d_89_v51_ = D0_DC1(p1, d_86_v48_, d_87_v49_, len(d_88_v50_))
                d_90_v52_: _dafny.Seq
                d_90_v52_ = _dafny.SeqWithoutIsStrInference([True, p1, p1])
                d_91_v53_: _dafny.Array
                nw13_ = _dafny.Array(None, 27)
                nw13_[int(0)] = p1
                nw13_[int(1)] = p1
                nw13_[int(2)] = (p1 if True else p1)
                nw13_[int(3)] = p1
                nw13_[int(4)] = True
                nw13_[int(5)] = p1
                nw13_[int(6)] = p1
                nw13_[int(7)] = p1
                nw13_[int(8)] = p1
                nw13_[int(9)] = not(not((d_89_v51_).cf2))
                nw13_[int(10)] = p1
                nw13_[int(11)] = (p0) < (p0)
                nw13_[int(12)] = (d_89_v51_).cf2
                nw13_[int(13)] = not(p1)
                nw13_[int(14)] = not(p1)
                nw13_[int(15)] = p1
                nw13_[int(16)] = (d_90_v52_)[default__.safeIndex((d_19_v1_)[default__.safeIndex(d_40_cf18_, len(d_19_v1_))], len(d_90_v52_))]
                nw13_[int(17)] = p1
                nw13_[int(18)] = p1
                nw13_[int(19)] = not((757) == (d_40_cf18_))
                nw13_[int(20)] = p1
                nw13_[int(21)] = True
                nw13_[int(22)] = p1
                nw13_[int(23)] = p1
                nw13_[int(24)] = not (p1) or (True)
                nw13_[int(25)] = p1
                nw13_[int(26)] = (p1) and (p1)
                d_91_v53_ = nw13_
                index10_ = default__.safeIndex(369, (d_91_v53_).length(0))
                (d_91_v53_)[index10_] = p1
                index11_ = default__.safeIndex(455, (d_85_cf10_).length(0))
                (d_85_cf10_)[index11_] = d_40_cf18_
                d_92_v54_: _dafny.Set
                d_92_v54_ = _dafny.Set({p1})
                d_93_v55_: D1
                d_93_v55_ = D1_DC6(d_92_v54_)
                d_94_v56_: D1
                d_94_v56_ = D1_DC8(d_93_v55_)
                d_95_v57_: D3
                d_95_v57_ = D3_DC14(D3_DC14(D3_DC13(p1, p1, d_94_v56_)))
                d_96_v58_: _dafny.Map
                d_96_v58_ = _dafny.Map({d_41_cf17_: d_95_v57_})
                index12_ = default__.safeIndex(137, (d_91_v53_).length(0))
                (d_91_v53_)[index12_] = (_dafny.Map({_dafny.CodePoint('g'): d_95_v57_})) != (d_96_v58_)
                d_97_v59_: D0
                d_97_v59_ = D0_DC0(_dafny.Set({d_40_cf18_}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsto")))
                d_98_v60_: _dafny.Map
                d_98_v60_ = _dafny.Map({d_97_v59_: (not(p1) if p1 else not(p1))})
                d_99_v61_: D1
                d_99_v61_ = D1_DC6(_dafny.Set({p1}))
                pat_let_tv2_ = d_90_v52_
                pat_let_tv3_ = d_18_v0_
                pat_let_tv4_ = d_90_v52_
                pat_let_tv5_ = p1
                d_100_v62_: _dafny.MultiSet
                def iife12_(_pat_let6_0):
                    def iife13_(d_101_dt__update__tmp_h3_):
                        def iife14_(_pat_let7_0):
                            def iife15_(d_102_dt__update_hcf12_h0_):
                                return D1_DC6(d_102_dt__update_hcf12_h0_)
                            return iife15_(_pat_let7_0)
                        return iife14_(_dafny.Set({(pat_let_tv2_)[default__.safeIndex(pat_let_tv3_, len(pat_let_tv4_))], pat_let_tv5_}))
                    return iife13_(_pat_let6_0)
                d_100_v62_ = _dafny.MultiSet([d_99_v61_, d_99_v61_, d_99_v61_, d_99_v61_, iife12_(d_99_v61_)])
                d_103_v63_: _dafny.Set
                d_103_v63_ = _dafny.Set({p0})
                d_104_v64_: _dafny.Seq
                d_104_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_97_v59_: p1})])
                d_105_v66_: _dafny.Map
                d_105_v66_ = _dafny.Map({d_97_v59_: d_41_cf17_})
                index13_ = default__.safeIndex(369, (d_91_v53_).length(0))
                index14_ = default__.safeIndex(455, (d_85_cf10_).length(0))
                index15_ = default__.safeIndex(137, (d_91_v53_).length(0))
                def iife16_():
                    coll0_ = _dafny.Map()
                    compr_0_: D0
                    for compr_0_ in (d_105_v66_).keys.Elements:
                        d_106_v65_: D0 = compr_0_
                        if (d_106_v65_) in (d_105_v66_):
                            coll0_[d_106_v65_] = p1
                    return _dafny.Map(coll0_)
                rhs10_ = (d_100_v62_) != (d_100_v62_)
                rhs11_ = default__.safeDivisionInt(d_18_v0_, len(d_103_v63_))
                rhs12_ = p1
                rhs13_ = D2_DC10()
                rhs14_ = ((d_104_v64_)[default__.safeIndex((0) - (d_18_v0_), len(d_104_v64_))]) | (iife16_()
                )
                lhs10_ = d_91_v53_
                lhs11_ = default__.safeIndex(369, (d_91_v53_).length(0))
                lhs12_ = d_85_cf10_
                lhs13_ = default__.safeIndex(455, (d_85_cf10_).length(0))
                lhs14_ = d_91_v53_
                lhs15_ = default__.safeIndex(137, (d_91_v53_).length(0))
                lhs10_[lhs11_] = rhs10_
                lhs12_[lhs13_] = rhs11_
                lhs14_[lhs15_] = rhs12_
                d_20_v2_ = rhs13_
                d_98_v60_ = rhs14_
                (globalState).f4 = (d_89_v51_).cf2
                r0 = d_40_cf18_
            elif True:
                d_107___mcc_h8_ = source3_.cf15
                d_108_cf15_ = d_107___mcc_h8_
                d_109_v67_: _dafny.Seq
                d_109_v67_ = _dafny.SeqWithoutIsStrInference([default__.fm2(p1, p1, globalState), (p0) + (p0), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_110_i0_ in range(default__.abs(413))]), p0])
                d_109_v67_ = d_109_v67_
                d_111_v68_: _dafny.Map
                d_111_v68_ = _dafny.Map({p1: default__.fm7(globalState)})
                (globalState).f4 = ((len(d_111_v68_)) - (len(_dafny.SeqWithoutIsStrInference([d_41_cf17_ for d_112_i1_ in range(default__.abs(242))])))) < ((d_18_v0_ if p1 else d_40_cf18_))
                d_113_v69_: T0
                nw14_ = C0()
                nw14_.ctor__(339, p1)
                d_113_v69_ = nw14_
                d_114_v70_: _dafny.Array
                nw15_ = _dafny.Array(int(0), 18)
                d_114_v70_ = nw15_
                d_115_v71_: D1
                d_115_v71_ = D1_DC4(d_114_v70_)
                d_116_v72_: D1
                d_116_v72_ = D1_DC8(D1_DC8(d_115_v71_))
                d_117_v73_: _dafny.Map
                d_117_v73_ = _dafny.Map({d_113_v69_: D3_DC14(D3_DC13((d_113_v69_).f19, p1, d_116_v72_))})
                d_118_v74_: _dafny.Array
                nw16_ = _dafny.Array(None, 10)
                nw16_[int(0)] = p1
                nw16_[int(1)] = (d_113_v69_).f19
                nw16_[int(2)] = p1
                nw16_[int(3)] = p1
                nw16_[int(4)] = (d_113_v69_).f19
                nw16_[int(5)] = p1
                nw16_[int(6)] = p1
                nw16_[int(7)] = (d_113_v69_).f19
                nw16_[int(8)] = p1
                nw16_[int(9)] = False
                d_118_v74_ = nw16_
                d_119_v75_: D3
                d_119_v75_ = D3_DC12(d_118_v74_)
                d_120_v76_: D3
                d_120_v76_ = D3_DC14(d_119_v75_)
                d_121_v77_: _dafny.Seq
                d_121_v77_ = _dafny.SeqWithoutIsStrInference([d_117_v73_, _dafny.Map({d_113_v69_: d_120_v76_})])
                d_122_v78_: _dafny.MultiSet
                d_122_v78_ = _dafny.MultiSet([p1, (d_113_v69_).f19])
                d_123_v79_: _dafny.Seq
                d_123_v79_ = _dafny.SeqWithoutIsStrInference([(d_121_v77_)[default__.safeIndex((d_122_v78_).cardinality, len(d_121_v77_))]])
                d_124_v80_: _dafny.Set
                d_124_v80_ = _dafny.Set({(d_113_v69_).f19, (d_113_v69_).f19, p1})
                d_125_v81_: _dafny.Map
                d_125_v81_ = _dafny.Map({d_113_v69_.f18: d_124_v80_})
                d_126_v82_: C0
                nw17_ = C0()
                nw17_.ctor__(len(p0), p1)
                d_126_v82_ = nw17_
                d_127_v83_: _dafny.Set
                d_127_v83_ = _dafny.Set({d_126_v82_})
                d_128_v84_: _dafny.Map
                d_128_v84_ = _dafny.Map({(d_113_v69_).f19: (d_113_v69_).f19})
                d_129_v85_: D2
                d_129_v85_ = D2_DC11(d_41_cf17_, d_40_cf18_)
                d_130_v86_: _dafny.Map
                d_130_v86_ = _dafny.Map({-886: p1})
                d_131_v87_: _dafny.Array
                nw18_ = _dafny.Array(None, 5)
                nw18_[int(0)] = d_129_v85_
                nw18_[int(1)] = d_129_v85_
                nw18_[int(2)] = d_129_v85_
                nw18_[int(3)] = D2_DC11(_dafny.CodePoint('k'), len(d_130_v86_))
                nw18_[int(4)] = d_129_v85_
                d_131_v87_ = nw18_
                d_132_v88_: _dafny.MultiSet
                d_132_v88_ = _dafny.MultiSet([(d_126_v82_).fm4(d_113_v69_.f18, (d_113_v69_).f19, p1, globalState), d_40_cf18_])
                d_133_v89_: _dafny.Array
                nw19_ = _dafny.Array(None, 28)
                nw19_[int(0)] = d_40_cf18_
                nw19_[int(1)] = (d_18_v0_) * (d_18_v0_)
                nw19_[int(2)] = default__.safeModuloInt(503, d_40_cf18_)
                nw19_[int(3)] = len((d_123_v79_)[default__.safeIndex(d_18_v0_, len(d_123_v79_))])
                nw19_[int(4)] = d_40_cf18_
                nw19_[int(5)] = (746) * (default__.fm0(len(((d_125_v81_)[d_18_v0_] if (d_18_v0_) in (d_125_v81_) else d_124_v80_)), True, globalState))
                nw19_[int(6)] = d_40_cf18_
                nw19_[int(7)] = len(d_127_v83_)
                nw19_[int(8)] = d_40_cf18_
                nw19_[int(9)] = d_113_v69_.f18
                nw19_[int(10)] = len(p0)
                nw19_[int(11)] = d_113_v69_.f18
                nw19_[int(12)] = (0) - (len(p0))
                nw19_[int(13)] = len((d_128_v84_) | (d_128_v84_))
                nw19_[int(14)] = d_113_v69_.f18
                nw19_[int(15)] = 594
                nw19_[int(16)] = (0) - ((d_113_v69_.f18) + (d_18_v0_))
                nw19_[int(17)] = d_18_v0_
                nw19_[int(18)] = default__.safeModuloInt((0) - (d_113_v69_.f18), (0) - (d_113_v69_.f18))
                nw19_[int(19)] = (len(_dafny.SeqWithoutIsStrInference([p1, not(p1), p1, (d_113_v69_).f19, (d_113_v69_).f19])) if default__.fm3(False, d_40_cf18_, d_113_v69_.f18, (d_113_v69_).f19, globalState) else len(p0))
                nw19_[int(20)] = len(_dafny.Map({(0) - ((0) - (d_113_v69_.f18)): d_131_v87_}))
                nw19_[int(21)] = (d_18_v0_ if False else d_113_v69_.f18)
                nw19_[int(22)] = default__.safeModuloInt(d_18_v0_, ((d_132_v88_)[(0) - (d_18_v0_)] if ((0) - (d_18_v0_)) in (d_132_v88_) else d_113_v69_.f18))
                nw19_[int(23)] = (d_113_v69_.f18) * (d_40_cf18_)
                nw19_[int(24)] = (d_113_v69_.f18) - (d_18_v0_)
                nw19_[int(25)] = default__.safeModuloInt(d_113_v69_.f18, d_18_v0_)
                nw19_[int(26)] = default__.safeModuloInt(d_113_v69_.f18, d_113_v69_.f18)
                nw19_[int(27)] = 713
                d_133_v89_ = nw19_
                index16_ = default__.safeIndex(727, (d_114_v70_).length(0))
                (d_114_v70_)[index16_] = d_113_v69_.f18
                index17_ = default__.safeIndex(727, (d_114_v70_).length(0))
                (d_114_v70_)[index17_] = ((0) - (default__.safeModuloInt(d_113_v69_.f18, d_18_v0_))) - (default__.safeDivisionInt(d_113_v69_.f18, d_18_v0_))
                d_134_v90_: _dafny.Map
                d_134_v90_ = _dafny.Map({d_113_v69_: (d_113_v69_).f19})
                d_135_v91_: _dafny.Map
                d_135_v91_ = _dafny.Map({p1: len(d_134_v90_)})
                d_136_v92_: _dafny.Set
                d_136_v92_ = _dafny.Set({len((d_135_v91_).set(True, len(p0))), (0) - (d_113_v69_.f18)})
                (globalState).f4 = (len(d_136_v92_)) != (d_113_v69_.f18)
            d_137_v93_: T0
            nw20_ = C0()
            nw20_.ctor__(d_18_v0_, not(p1))
            d_137_v93_ = nw20_
            d_138_v94_: _dafny.Map
            d_138_v94_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not(p1)]): d_137_v93_})
            d_139_v95_: _dafny.Seq
            d_139_v95_ = _dafny.SeqWithoutIsStrInference([p1])
            d_138_v94_ = (d_138_v94_).set(d_139_v95_, d_137_v93_)
            d_140_v96_: _dafny.Set
            d_140_v96_ = _dafny.Set({(d_137_v93_).f19})
            d_141_v97_: _dafny.Map
            d_141_v97_ = _dafny.Map({(_dafny.Map({d_18_v0_: (d_19_v1_)[default__.safeIndex(d_18_v0_, len(d_19_v1_))]})).set(len(d_140_v96_), d_137_v93_.f18): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyjjr"))})
            d_142_v98_: _dafny.Map
            d_142_v98_ = _dafny.Map({d_18_v0_: (_dafny.MultiSet([d_137_v93_.f18])).cardinality})
            d_141_v97_ = (d_141_v97_).set((d_142_v98_) | (d_142_v98_), (p0) + (p0))
            d_143_v99_: _dafny.Map
            d_143_v99_ = _dafny.Map({p1: d_137_v93_.f18})
            if ((((d_143_v99_)[(d_137_v93_).f19] if ((d_137_v93_).f19) in (d_143_v99_) else d_137_v93_.f18) if p1 else d_137_v93_.f18)) < (d_40_cf18_):
                d_144_v100_: _dafny.MultiSet
                d_144_v100_ = _dafny.MultiSet([(d_137_v93_).f19])
                d_145_v101_: _dafny.Array
                nw21_ = _dafny.Array(None, 14)
                nw21_[int(0)] = d_144_v100_
                nw21_[int(1)] = (d_144_v100_).intersection(_dafny.MultiSet([False, p1]))
                nw21_[int(2)] = (d_144_v100_).set((d_137_v93_).f19, default__.abs(576))
                nw21_[int(3)] = _dafny.MultiSet(d_139_v95_)
                nw21_[int(4)] = _dafny.MultiSet(d_139_v95_)
                nw21_[int(5)] = d_144_v100_
                nw21_[int(6)] = d_144_v100_
                nw21_[int(7)] = (d_144_v100_).intersection(d_144_v100_)
                nw21_[int(8)] = d_144_v100_
                nw21_[int(9)] = _dafny.MultiSet([(d_137_v93_).f19, not((d_137_v93_).f19)])
                nw21_[int(10)] = (d_144_v100_) - (d_144_v100_)
                nw21_[int(11)] = (d_144_v100_).set(p1, default__.abs(len(p0)))
                nw21_[int(12)] = d_144_v100_
                nw21_[int(13)] = d_144_v100_
                d_145_v101_ = nw21_
                index18_ = default__.safeIndex(578, (d_145_v101_).length(0))
                (d_145_v101_)[index18_] = d_144_v100_
                d_146_v102_: _dafny.MultiSet
                d_146_v102_ = d_144_v100_
                d_147_v103_: C0
                nw22_ = C0()
                nw22_.ctor__(263, (_dafny.MultiSet(d_139_v95_)).ispropersubset((d_146_v102_)))
                d_147_v103_ = nw22_
                d_148_v104_: _dafny.Array
                nw23_ = _dafny.Array(None, 16)
                nw23_[int(0)] = (d_137_v93_).f19
                nw23_[int(1)] = p1
                nw23_[int(2)] = p1
                nw23_[int(3)] = (d_137_v93_).f19
                nw23_[int(4)] = not(p1)
                nw23_[int(5)] = (d_137_v93_).f19
                nw23_[int(6)] = (d_137_v93_).f19
                nw23_[int(7)] = (d_137_v93_).f19
                nw23_[int(8)] = (d_137_v93_).f19
                nw23_[int(9)] = not(p1)
                nw23_[int(10)] = p1
                nw23_[int(11)] = p1
                nw23_[int(12)] = (d_137_v93_).f19
                nw23_[int(13)] = (d_137_v93_).f19
                nw23_[int(14)] = (d_137_v93_).f19
                nw23_[int(15)] = (d_137_v93_).f19
                d_148_v104_ = nw23_
                d_149_v105_: D3
                d_149_v105_ = D3_DC12(d_148_v104_)
                d_150_v106_: _dafny.Set
                d_150_v106_ = _dafny.Set({len(d_140_v96_)})
                d_151_v107_: D0
                d_151_v107_ = D0_DC0(d_150_v106_, p0)
                index19_ = default__.safeIndex(578, (d_145_v101_).length(0))
                rhs15_ = (d_144_v100_ if p1 else d_144_v100_)
                rhs16_ = d_147_v103_
                rhs17_ = d_149_v105_
                rhs18_ = d_151_v107_
                lhs16_ = d_145_v101_
                lhs17_ = default__.safeIndex(578, (d_145_v101_).length(0))
                lhs16_[lhs17_] = rhs15_
                d_147_v103_ = rhs16_
                d_149_v105_ = rhs17_
                d_151_v107_ = rhs18_
                d_146_v102_ = d_146_v102_
                d_142_v98_ = (d_142_v98_).set(default__.safeModuloInt(d_137_v93_.f18, len(p0)), d_18_v0_)
                (globalState).f4 = (d_40_cf18_) > (d_40_cf18_)
                d_152_v108_: _dafny.Set
                d_152_v108_ = _dafny.Set({d_147_v103_, d_147_v103_})
                (globalState).f16 = default__.safeModuloInt((175) - (d_18_v0_), len((d_152_v108_).intersection(d_152_v108_)))
            elif True:
                (globalState).f4 = default__.fm3(p1, (d_40_cf18_ if (d_137_v93_).f19 else d_137_v93_.f18), d_40_cf18_, p1, globalState)
                (globalState).f0 = (((0) - (d_137_v93_.f18)) + (default__.fm0(d_137_v93_.f18, (d_137_v93_).f19, globalState))) * (-271)
                d_153_v109_: _dafny.Array
                nw24_ = _dafny.Array(_dafny.Map({}), 9)
                d_153_v109_ = nw24_
                d_154_v110_: _dafny.Map
                d_154_v110_ = _dafny.Map({d_137_v93_: (d_137_v93_).f19})
                d_155_v111_: _dafny.Set
                d_155_v111_ = _dafny.Set({d_154_v110_, d_154_v110_, d_154_v110_, d_154_v110_, d_154_v110_})
                d_156_v112_: D0
                d_156_v112_ = D0_DC2(len(d_155_v111_), _dafny.MultiSet([d_137_v93_.f18, d_137_v93_.f18, d_18_v0_]), (d_137_v93_).f19)
                d_157_v113_: _dafny.Map
                d_157_v113_ = _dafny.Map({d_156_v112_: False})
                index20_ = default__.safeIndex(165, (d_153_v109_).length(0))
                (d_153_v109_)[index20_] = d_157_v113_
                index21_ = default__.safeIndex(165, (d_153_v109_).length(0))
                (d_153_v109_)[index21_] = d_157_v113_
                d_158_v114_: _dafny.Set
                d_158_v114_ = _dafny.Set({len(d_142_v98_), d_40_cf18_, d_18_v0_})
                d_159_v115_: C0
                nw25_ = C0()
                nw25_.ctor__((0) - (d_137_v93_.f18), (d_137_v93_).f19)
                d_159_v115_ = nw25_
                d_160_v116_: _dafny.Seq
                d_160_v116_ = _dafny.SeqWithoutIsStrInference([d_159_v115_, d_159_v115_])
                d_161_v117_: _dafny.Map
                d_161_v117_ = _dafny.Map({len(d_158_v114_): p1})
                d_162_v118_: _dafny.MultiSet
                d_162_v118_ = _dafny.MultiSet([p1, (d_137_v93_).f19, p1, (d_137_v93_).f19, (d_137_v93_).f19])
                d_163_v119_: _dafny.MultiSet
                d_163_v119_ = _dafny.MultiSet([d_137_v93_])
                d_164_v120_: D2
                d_164_v120_ = D2_DC11(d_41_cf17_, d_40_cf18_)
                d_165_v121_: _dafny.Array
                nw26_ = _dafny.Array(None, 29)
                nw26_[int(0)] = len(p0)
                nw26_[int(1)] = (d_19_v1_)[default__.safeIndex(539, len(d_19_v1_))]
                nw26_[int(2)] = d_137_v93_.f18
                nw26_[int(3)] = d_137_v93_.f18
                nw26_[int(4)] = (default__.fm0(d_137_v93_.f18, (d_137_v93_).f19, globalState) if p1 else d_18_v0_)
                nw26_[int(5)] = default__.safeModuloInt(d_18_v0_, d_137_v93_.f18)
                nw26_[int(6)] = d_18_v0_
                nw26_[int(7)] = (d_19_v1_)[default__.safeIndex(len(d_158_v114_), len(d_19_v1_))]
                nw26_[int(8)] = (d_40_cf18_) * (len(d_160_v116_))
                nw26_[int(9)] = d_137_v93_.f18
                nw26_[int(10)] = d_137_v93_.f18
                nw26_[int(11)] = (d_40_cf18_) + (d_137_v93_.f18)
                nw26_[int(12)] = 195
                nw26_[int(13)] = 930
                nw26_[int(14)] = d_137_v93_.f18
                nw26_[int(15)] = len(d_139_v95_)
                nw26_[int(16)] = (d_137_v93_.f18 if (d_137_v93_).f19 else len(d_161_v117_))
                nw26_[int(17)] = d_40_cf18_
                nw26_[int(18)] = d_137_v93_.f18
                nw26_[int(19)] = ((d_162_v118_)[default__.fm3((d_137_v93_).f19, d_18_v0_, 104, p1, globalState)] if (default__.fm3((d_137_v93_).f19, d_18_v0_, 104, p1, globalState)) in (d_162_v118_) else (d_19_v1_)[default__.safeIndex(d_40_cf18_, len(d_19_v1_))])
                nw26_[int(20)] = d_137_v93_.f18
                nw26_[int(21)] = d_40_cf18_
                nw26_[int(22)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcwsjh")))
                nw26_[int(23)] = (d_163_v119_).cardinality
                nw26_[int(24)] = len(p0)
                nw26_[int(25)] = (d_164_v120_).cf18
                nw26_[int(26)] = (d_159_v115_).fm4(d_18_v0_, (d_137_v93_).f19, (d_137_v93_).f19, globalState)
                nw26_[int(27)] = d_18_v0_
                nw26_[int(28)] = d_18_v0_
                d_165_v121_ = nw26_
                index22_ = default__.safeIndex(481, (d_165_v121_).length(0))
                (d_165_v121_)[index22_] = default__.safeModuloInt(d_18_v0_, d_40_cf18_)
                index23_ = default__.safeIndex(481, (d_165_v121_).length(0))
                (d_165_v121_)[index23_] = d_137_v93_.f18
                d_166_v122_: _dafny.Array
                nw27_ = _dafny.Array(None, 20)
                nw27_[int(0)] = (d_140_v96_) - (d_140_v96_)
                nw27_[int(1)] = d_140_v96_
                nw27_[int(2)] = d_140_v96_
                nw27_[int(3)] = _dafny.Set({p1})
                nw27_[int(4)] = d_140_v96_
                nw27_[int(5)] = d_140_v96_
                nw27_[int(6)] = d_140_v96_
                nw27_[int(7)] = (d_140_v96_) - (d_140_v96_)
                nw27_[int(8)] = (d_140_v96_) | (d_140_v96_)
                nw27_[int(9)] = (default__.fm6((d_165_v121_)[default__.safeIndex(481, (d_165_v121_).length(0))], (p0).set(default__.safeIndex(d_137_v93_.f18, len(p0)), d_41_cf17_), not((d_137_v93_).f19), globalState)) | (d_140_v96_)
                nw27_[int(10)] = d_140_v96_
                nw27_[int(11)] = (d_140_v96_).intersection(d_140_v96_)
                nw27_[int(12)] = d_140_v96_
                nw27_[int(13)] = (d_140_v96_) | (d_140_v96_)
                nw27_[int(14)] = (d_140_v96_).intersection(d_140_v96_)
                nw27_[int(15)] = d_140_v96_
                nw27_[int(16)] = d_140_v96_
                nw27_[int(17)] = _dafny.Set({(d_137_v93_).f19, (d_137_v93_).f19})
                nw27_[int(18)] = (d_140_v96_) - (d_140_v96_)
                nw27_[int(19)] = d_140_v96_
                d_166_v122_ = nw27_
                index24_ = default__.safeIndex(150, (d_166_v122_).length(0))
                (d_166_v122_)[index24_] = d_140_v96_
                index25_ = default__.safeIndex(150, (d_166_v122_).length(0))
                (d_166_v122_)[index25_] = (_dafny.Set({(d_137_v93_).f19, not(True)})) - (d_140_v96_)
        elif True:
            d_167___mcc_h2_ = source2_.cf16
            d_168_cf16_ = d_167___mcc_h2_
            d_169_v123_: _dafny.MultiSet
            d_169_v123_ = _dafny.MultiSet([((_dafny.MultiSet([p1, p1])).set(p1, default__.abs(d_18_v0_))).cardinality])
            d_170_v124_: T0
            nw28_ = C0()
            nw28_.ctor__(d_18_v0_, p1)
            d_170_v124_ = nw28_
            d_171_v125_: _dafny.Map
            d_171_v125_ = _dafny.Map({d_170_v124_: p1})
            (globalState).f4 = default__.fm3(p1, (d_19_v1_)[default__.safeIndex(((d_169_v123_)[default__.fm0(len(d_171_v125_), p1, globalState)] if (default__.fm0(len(d_171_v125_), p1, globalState)) in (d_169_v123_) else d_170_v124_.f18), len(d_19_v1_))], d_18_v0_, (d_170_v124_).f19, globalState)
            d_172_v126_: _dafny.Map
            d_172_v126_ = _dafny.Map({p1: not(p1)})
            d_173_v127_: _dafny.MultiSet
            d_173_v127_ = _dafny.MultiSet([p1, (d_170_v124_).f19, (d_170_v124_).f19])
            d_174_v128_: _dafny.Map
            d_174_v128_ = _dafny.Map({d_173_v127_: (d_172_v126_) | ((_dafny.Map({p1: p1})).set(p1, True))})
            d_175_v129_: _dafny.Seq
            d_175_v129_ = _dafny.SeqWithoutIsStrInference([p1])
            d_176_v130_: D2
            d_176_v130_ = D2_DC9(d_168_cf16_)
            pat_let_tv6_ = d_168_cf16_
            pat_let_tv7_ = d_168_cf16_
            def iife17_(_pat_let8_0):
                def iife18_(d_177_dt__update__tmp_h5_):
                    def iife19_(_pat_let9_0):
                        def iife20_(d_178_dt__update_hcf16_h1_):
                            return D2_DC9(d_178_dt__update_hcf16_h1_)
                        return iife20_(_pat_let9_0)
                    return iife19_(pat_let_tv6_)
                return iife18_(_pat_let8_0)
            def iife21_(_pat_let10_0):
                def iife22_(d_179_dt__update__tmp_h4_):
                    def iife23_(_pat_let11_0):
                        def iife24_(d_180_dt__update_hcf16_h0_):
                            return D2_DC9(d_180_dt__update_hcf16_h0_)
                        return iife24_(_pat_let11_0)
                    return iife23_(pat_let_tv7_)
                return iife22_(_pat_let10_0)
            rhs19_ = ((d_174_v128_)[(_dafny.MultiSet(d_175_v129_)) - (default__.fm8(default__.fm0((0) - (d_170_v124_.f18), (d_170_v124_).f19, globalState), iife17_(d_176_v130_), not(p1), p0, globalState))] if ((_dafny.MultiSet(d_175_v129_)) - (default__.fm8(default__.fm0((0) - (d_170_v124_.f18), (d_170_v124_).f19, globalState), iife21_(d_176_v130_), not(p1), p0, globalState))) in (d_174_v128_) else d_172_v126_)
            rhs20_ = True
            rhs21_ = (d_170_v124_).f19
            lhs18_ = globalState
            lhs19_ = globalState
            d_172_v126_ = rhs19_
            lhs18_.f4 = rhs20_
            lhs19_.f4 = rhs21_
            d_181_v131_: str
            d_181_v131_ = _dafny.CodePoint('p')
            d_181_v131_ = d_181_v131_
            (globalState).f4 = (d_181_v131_) not in (p0)
        (globalState).f16 = default__.fm0(d_18_v0_, p1, globalState)
        if not(not(p1)):
            (globalState).f16 = (d_19_v1_)[default__.safeIndex(d_18_v0_, len(d_19_v1_))]
            d_182_v132_: _dafny.Map
            d_182_v132_ = _dafny.Map({p1: len(_dafny.Map({d_19_v1_: p1}))})
            d_183_v133_: C0
            nw29_ = C0()
            nw29_.ctor__(len(d_182_v132_), p1)
            d_183_v133_ = nw29_
            rhs22_ = d_183_v133_
            rhs23_ = (d_18_v0_ if p1 else d_18_v0_)
            lhs20_ = globalState
            d_183_v133_ = rhs22_
            lhs20_.f0 = rhs23_
            d_184_v134_: _dafny.Map
            d_184_v134_ = _dafny.Map({default__.safeDivisionInt((d_19_v1_)[default__.safeIndex(d_18_v0_, len(d_19_v1_))], (0) - ((0) - (default__.fm0(d_18_v0_, default__.fm3(p1, d_18_v0_, 932, p1, globalState), globalState)))): p1})
            d_184_v134_ = (d_184_v134_).set(d_18_v0_, p1)
            (globalState).f0 = 165
            d_185_v135_: _dafny.Array
            nw30_ = _dafny.Array(False, 27)
            d_185_v135_ = nw30_
            d_185_v135_ = d_185_v135_
        elif True:
            d_186_v136_: _dafny.Array
            nw31_ = _dafny.Array(None, 24)
            nw31_[int(0)] = d_18_v0_
            nw31_[int(1)] = len(p0)
            nw31_[int(2)] = default__.safeModuloInt(d_18_v0_, d_18_v0_)
            nw31_[int(3)] = d_18_v0_
            nw31_[int(4)] = (d_18_v0_) * (d_18_v0_)
            nw31_[int(5)] = d_18_v0_
            nw31_[int(6)] = d_18_v0_
            nw31_[int(7)] = d_18_v0_
            nw31_[int(8)] = 521
            nw31_[int(9)] = d_18_v0_
            nw31_[int(10)] = d_18_v0_
            nw31_[int(11)] = d_18_v0_
            nw31_[int(12)] = 999
            nw31_[int(13)] = d_18_v0_
            nw31_[int(14)] = d_18_v0_
            nw31_[int(15)] = ((0) - (d_18_v0_) if p1 else d_18_v0_)
            nw31_[int(16)] = (0) - ((0) - ((d_18_v0_) - (d_18_v0_)))
            nw31_[int(17)] = d_18_v0_
            nw31_[int(18)] = d_18_v0_
            nw31_[int(19)] = 681
            nw31_[int(20)] = d_18_v0_
            nw31_[int(21)] = d_18_v0_
            nw31_[int(22)] = d_18_v0_
            nw31_[int(23)] = default__.fm0(d_18_v0_, False, globalState)
            d_186_v136_ = nw31_
            index26_ = default__.safeIndex(773, (d_186_v136_).length(0))
            (d_186_v136_)[index26_] = (d_18_v0_ if p1 else 592)
            index27_ = default__.safeIndex(773, (d_186_v136_).length(0))
            (d_186_v136_)[index27_] = d_18_v0_
            (globalState).f4 = (p0) < (p0)
            d_187_v137_: _dafny.MultiSet
            d_187_v137_ = _dafny.MultiSet([d_18_v0_])
            index28_ = default__.safeIndex(773, (d_186_v136_).length(0))
            (d_186_v136_)[index28_] = (default__.safeModuloInt((d_187_v137_).cardinality, d_18_v0_) if p1 else (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))])
            d_188_v138_: _dafny.Set
            d_188_v138_ = _dafny.Set({not(p1)})
            pat_let_tv8_ = d_188_v138_
            def iife25_(_pat_let12_0):
                def iife26_(d_189_dt__update__tmp_h6_):
                    def iife27_(_pat_let13_0):
                        def iife28_(d_190_dt__update_hcf12_h1_):
                            return D1_DC6(d_190_dt__update_hcf12_h1_)
                        return iife28_(_pat_let13_0)
                    return iife27_(pat_let_tv8_)
                return iife26_(_pat_let12_0)
            source4_ = iife25_(default__.fm10(p1, (0) - ((d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]), globalState))
            if source4_.is_DC5:
                d_191___mcc_h9_ = source4_.cf11
                d_192_cf11_ = d_191___mcc_h9_
                d_193_v139_: C0
                nw32_ = C0()
                nw32_.ctor__((d_187_v137_).cardinality, p1)
                d_193_v139_ = nw32_
                d_194_v140_: T0
                nw33_ = C0()
                nw33_.ctor__((d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))], (not(not(p1)) if p1 else p1))
                d_194_v140_ = nw33_
                nw34_ = C0()
                nw34_.ctor__(default__.safeDivisionInt(len(d_19_v1_), (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]), p1)
                d_194_v140_ = nw34_
                d_195_v141_: _dafny.Map
                d_195_v141_ = _dafny.Map({(d_194_v140_).f19: d_18_v0_})
                (globalState).f16 = (633) * (((d_195_v141_)[p1] if (p1) in (d_195_v141_) else (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]))
                d_196_v142_: _dafny.Map
                d_196_v142_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_194_v140_.f18]): d_194_v140_.f18})
                d_197_v143_: _dafny.MultiSet
                d_197_v143_ = _dafny.MultiSet([d_187_v137_, default__.fm11(d_196_v142_, p1, d_188_v138_, globalState), (d_187_v137_).set(d_18_v0_, default__.abs(d_18_v0_)), d_187_v137_])
                d_198_v144_: _dafny.Seq
                d_198_v144_ = _dafny.SeqWithoutIsStrInference([d_187_v137_])
                d_197_v143_ = (d_197_v143_) | (_dafny.MultiSet(d_198_v144_))
            elif source4_.is_DC6:
                d_199___mcc_h10_ = source4_.cf12
                d_200_cf12_ = d_199___mcc_h10_
                d_201_v146_: _dafny.Array
                def lambda0_(d_202_v0_, d_203_p0_, d_204_v136_, d_205_p1_):
                    def lambda1_(d_206_i2_):
                        def iife29_():
                            coll1_ = _dafny.Set()
                            compr_1_: _dafny.Map
                            for compr_1_ in (_dafny.Map({_dafny.Map({len(d_203_p0_): (d_204_v136_)[default__.safeIndex(773, (d_204_v136_).length(0))]}): d_205_p1_})).keys.Elements:
                                d_207_v145_: _dafny.Map = compr_1_
                                if (d_207_v145_) in (_dafny.Map({_dafny.Map({len(d_203_p0_): (d_204_v136_)[default__.safeIndex(773, (d_204_v136_).length(0))]}): d_205_p1_})):
                                    coll1_ = coll1_.union(_dafny.Set([d_207_v145_]))
                            return _dafny.Set(coll1_)
                        return (iife29_()
                        ) | (_dafny.Set({_dafny.Map({165: d_202_v0_}), _dafny.Map({d_202_v0_: len(_dafny.SeqWithoutIsStrInference([d_205_p1_]))})}))

                    return lambda1_

                init0_ = lambda0_(d_18_v0_, p0, d_186_v136_, p1)
                nw35_ = _dafny.Array(None, 16)
                for i0_0_ in range(nw35_.length(0)):
                    nw35_[i0_0_] = init0_(i0_0_)
                d_201_v146_ = nw35_
                d_208_v147_: _dafny.Map
                d_208_v147_ = _dafny.Map({(d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]: d_18_v0_})
                d_209_v148_: _dafny.Set
                d_209_v148_ = _dafny.Set({d_208_v147_, d_208_v147_})
                index29_ = default__.safeIndex(673, (d_201_v146_).length(0))
                (d_201_v146_)[index29_] = d_209_v148_
                d_210_v149_: _dafny.Seq
                d_210_v149_ = _dafny.SeqWithoutIsStrInference([d_208_v147_, d_208_v147_])
                index30_ = default__.safeIndex(673, (d_201_v146_).length(0))
                def iife30_():
                    coll2_ = _dafny.Set()
                    compr_2_: _dafny.Map
                    for compr_2_ in (d_210_v149_).Elements:
                        d_211_v150_: _dafny.Map = compr_2_
                        if (d_211_v150_) in (d_210_v149_):
                            coll2_ = coll2_.union(_dafny.Set([d_211_v150_]))
                    return _dafny.Set(coll2_)
                (d_201_v146_)[index30_] = (d_209_v148_) | (iife30_()
                )
                d_212_v151_: _dafny.Map
                d_212_v151_ = _dafny.Map({(0) - (d_18_v0_): p1})
                d_212_v151_ = (d_212_v151_) | (d_212_v151_)
                d_213_v152_: _dafny.Array
                nw36_ = _dafny.Array(None, 10)
                d_213_v152_ = nw36_
                d_214_v153_: _dafny.Map
                d_214_v153_ = _dafny.Map({d_18_v0_: d_213_v152_})
                def iife31_():
                    def iife33_():
                        coll5_ = _dafny.Set()
                        compr_5_: int
                        for compr_5_ in _dafny.IntegerRange(-154, -722):
                            d_217_v155_: int = compr_5_
                            if ((-154) <= (d_217_v155_)) and ((d_217_v155_) < (-722)):
                                coll5_ = coll5_.union(_dafny.Set([(d_217_v155_) * (d_18_v0_)]))
                        return _dafny.Set(coll5_)
                    coll3_ = _dafny.Map()
                    def iife32_():
                        coll4_ = _dafny.Set()
                        compr_4_: int
                        for compr_4_ in _dafny.IntegerRange(-154, -722):
                            d_215_v155_: int = compr_4_
                            if ((-154) <= (d_215_v155_)) and ((d_215_v155_) < (-722)):
                                coll4_ = coll4_.union(_dafny.Set([(d_215_v155_) * (d_18_v0_)]))
                        return _dafny.Set(coll4_)
                    compr_3_: int
                    for compr_3_ in (iife32_()
                    ).Elements:
                        d_216_v154_: int = compr_3_
                        if (d_216_v154_) in (iife33_()
                        ):
                            coll3_[(d_216_v154_) - (d_18_v0_)] = _dafny.CodePoint('g')
                    return _dafny.Map(coll3_)
                d_214_v153_ = (d_214_v153_).set((len(iife31_()
                )) + ((d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]), d_213_v152_)
                (globalState).f4 = not(p1)
            elif source4_.is_DC7:
                d_218___mcc_h11_ = source4_.cf13
                d_219___mcc_h12_ = source4_.cf14
                d_220_cf14_ = d_219___mcc_h12_
                d_221_cf13_ = d_218___mcc_h11_
                d_222_v156_: _dafny.Array
                def lambda2_(d_223_i3_):
                    return False

                init1_ = lambda2_
                nw37_ = _dafny.Array(None, 3)
                for i0_1_ in range(nw37_.length(0)):
                    nw37_[i0_1_] = init1_(i0_1_)
                d_222_v156_ = nw37_
                index31_ = default__.safeIndex(752, (d_222_v156_).length(0))
                (d_222_v156_)[index31_] = p1
                d_224_v157_: _dafny.Map
                d_224_v157_ = _dafny.Map({p1: ((d_19_v1_)[default__.safeIndex((0) - (d_220_cf14_), len(d_19_v1_))]) not in (d_19_v1_)})
                d_225_v158_: C0
                nw38_ = C0()
                nw38_.ctor__(d_220_cf14_, p1)
                d_225_v158_ = nw38_
                d_226_v159_: D0
                d_226_v159_ = D0_DC1(p1, d_224_v157_, d_225_v158_, (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))])
                index32_ = default__.safeIndex(752, (d_222_v156_).length(0))
                (d_222_v156_)[index32_] = not(not(((d_224_v157_)[(d_226_v159_).cf2] if ((d_226_v159_).cf2) in (d_224_v157_) else p1)))
                d_227_v160_: D3
                d_227_v160_ = D3_DC12(d_222_v156_)
                d_228_v161_: _dafny.Seq
                d_228_v161_ = _dafny.SeqWithoutIsStrInference([d_227_v160_])
                d_229_v162_: _dafny.Map
                d_229_v162_ = _dafny.Map({((d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))] if default__.fm3((d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))], d_221_cf13_, (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))], False, globalState) else (d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))]): (d_228_v161_).set(default__.safeIndex(len(p0), len(d_228_v161_)), d_227_v160_)})
                d_229_v162_ = (d_229_v162_).set((d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))], d_228_v161_)
                d_230_v163_: D1
                d_230_v163_ = D1_DC4(d_186_v136_)
                d_231_v164_: _dafny.Seq
                d_231_v164_ = _dafny.SeqWithoutIsStrInference([(d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))], (d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))]])
                rhs24_ = d_230_v163_
                rhs25_ = d_225_v158_
                rhs26_ = (d_231_v164_ if (d_222_v156_)[default__.safeIndex(752, (d_222_v156_).length(0))] else d_231_v164_)
                lhs21_ = globalState
                d_230_v163_ = rhs24_
                d_225_v158_ = rhs25_
                lhs21_.f6 = rhs26_
                d_220_cf14_ = (d_19_v1_)[default__.safeIndex(d_220_cf14_, len(d_19_v1_))]
            elif source4_.is_DC4:
                d_232___mcc_h13_ = source4_.cf10
                d_233_cf10_ = d_232___mcc_h13_
                (globalState).f4 = p1
                (globalState).f16 = (d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]
                (globalState).f4 = ((d_18_v0_) * (d_18_v0_)) < ((d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))])
                d_234_v165_: C0
                nw39_ = C0()
                nw39_.ctor__((d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))], p1)
                d_234_v165_ = nw39_
                d_234_v165_ = d_234_v165_
            elif True:
                d_235___mcc_h14_ = source4_.cf15
                d_236_cf15_ = d_235___mcc_h14_
                d_237_v166_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Set({}), 1)
                d_237_v166_ = nw40_
                d_238_v167_: C0
                nw41_ = C0()
                nw41_.ctor__(2, True)
                d_238_v167_ = nw41_
                d_239_v168_: _dafny.Set
                d_239_v168_ = _dafny.Set({d_238_v167_, d_238_v167_})
                index33_ = default__.safeIndex(247, (d_237_v166_).length(0))
                (d_237_v166_)[index33_] = d_239_v168_
                index34_ = default__.safeIndex(247, (d_237_v166_).length(0))
                (d_237_v166_)[index34_] = (d_239_v168_ if p1 else d_239_v168_)
                d_18_v0_ = 99
                d_240_v169_: _dafny.Array
                nw42_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_240_v169_ = nw42_
                d_241_v170_: _dafny.Map
                d_241_v170_ = _dafny.Map({d_240_v169_: p1})
                d_241_v170_ = (d_241_v170_).set(d_240_v169_, p1)
                d_242_v171_: _dafny.Set
                d_242_v171_ = _dafny.Set({len(p0)})
                (globalState).f4 = default__.fm3(p1, default__.safeModuloInt(d_18_v0_, len(p0)), len(d_242_v171_), p1, globalState)
            d_243_v172_: _dafny.Array
            nw43_ = _dafny.Array(None, 1)
            nw43_[int(0)] = p1
            d_243_v172_ = nw43_
            d_244_v173_: _dafny.Seq
            d_244_v173_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlhubva"))])
            index35_ = default__.safeIndex(925, (d_243_v172_).length(0))
            (d_243_v172_)[index35_] = (p0) in (d_244_v173_)
            d_245_v174_: _dafny.Set
            d_245_v174_ = _dafny.Set({(d_186_v136_)[default__.safeIndex(773, (d_186_v136_).length(0))]})
            index36_ = default__.safeIndex(925, (d_243_v172_).length(0))
            (d_243_v172_)[index36_] = ((d_245_v174_) | (d_245_v174_)).isdisjoint(d_245_v174_)
        (globalState).f4 = p1
        d_246_v175_: _dafny.Seq
        d_246_v175_ = _dafny.SeqWithoutIsStrInference([p1, not(p1), p1])
        hi0_ = d_18_v0_
        for d_247_i4_ in range(default__.fm0(len(d_246_v175_), p1, globalState), hi0_):
            d_248_v176_: C0
            nw44_ = C0()
            nw44_.ctor__(d_247_i4_, p1)
            d_248_v176_ = nw44_
            d_248_v176_ = d_248_v176_
            r0 = d_247_i4_
            d_249_v177_: _dafny.Array
            def lambda3_(d_250_p1_):
                def lambda4_(d_251_i5_):
                    return (False) or (d_250_p1_)

                return lambda4_

            init2_ = lambda3_(p1)
            nw45_ = _dafny.Array(None, 4)
            for i0_2_ in range(nw45_.length(0)):
                nw45_[i0_2_] = init2_(i0_2_)
            d_249_v177_ = nw45_
            index37_ = default__.safeIndex(718, (d_249_v177_).length(0))
            (d_249_v177_)[index37_] = p1
            d_252_v178_: D0
            d_252_v178_ = D0_DC1(p1, _dafny.Map({p1: p1}), d_248_v176_, d_247_i4_)
            d_253_v179_: _dafny.MultiSet
            d_253_v179_ = _dafny.MultiSet([p1])
            d_254_v180_: _dafny.Map
            d_254_v180_ = _dafny.Map({d_252_v178_: (d_253_v179_).cardinality})
            index38_ = default__.safeIndex(718, (d_249_v177_).length(0))
            (d_249_v177_)[index38_] = ((d_19_v1_)[default__.safeIndex(len((d_254_v180_).set(d_252_v178_, d_247_i4_)), len(d_19_v1_))]) > ((d_247_i4_) - (258))
            if (d_249_v177_)[default__.safeIndex(718, (d_249_v177_).length(0))]:
                d_255_v181_: C0
                nw46_ = C0()
                nw46_.ctor__(d_18_v0_, p1)
                d_255_v181_ = nw46_
                (globalState).f16 = d_18_v0_
                d_256_v182_: str
                d_256_v182_ = _dafny.CodePoint('f')
                d_257_v183_: _dafny.Array
                nw47_ = _dafny.Array(None, 25)
                nw47_[int(0)] = (p0) + (p0)
                nw47_[int(1)] = (p0) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_258_i6_ in range(default__.abs(147))]))
                nw47_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_259_i7_ in range(default__.abs(290))])
                nw47_[int(3)] = p0
                nw47_[int(4)] = (p0) + ((p0).set(default__.safeIndex(d_18_v0_, len(p0)), d_256_v182_))
                nw47_[int(5)] = default__.fm2(True, p1, globalState)
                nw47_[int(6)] = default__.fm2(p1, (d_249_v177_)[default__.safeIndex(718, (d_249_v177_).length(0))], globalState)
                nw47_[int(7)] = _dafny.SeqWithoutIsStrInference([d_256_v182_ for d_260_i8_ in range(default__.abs(106))])
                nw47_[int(8)] = p0
                nw47_[int(9)] = p0
                nw47_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lo"))
                nw47_[int(11)] = p0
                nw47_[int(12)] = p0
                nw47_[int(13)] = p0
                nw47_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyu"))) + (p0)
                nw47_[int(15)] = p0
                nw47_[int(16)] = ((p0) + (p0)).set(default__.safeIndex(d_247_i4_, len((p0) + (p0))), d_256_v182_)
                nw47_[int(17)] = p0
                nw47_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujrexjxx"))
                nw47_[int(19)] = _dafny.SeqWithoutIsStrInference([d_256_v182_ for d_261_i9_ in range(default__.abs(-616))])
                nw47_[int(20)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekt"))).set(default__.safeIndex(d_18_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekt")))), d_256_v182_)) + (p0)
                nw47_[int(21)] = p0
                nw47_[int(22)] = default__.fm2(p1, (d_249_v177_)[default__.safeIndex(718, (d_249_v177_).length(0))], globalState)
                nw47_[int(23)] = default__.fm2(p1, (d_249_v177_)[default__.safeIndex(718, (d_249_v177_).length(0))], globalState)
                nw47_[int(24)] = p0
                d_257_v183_ = nw47_
                index39_ = default__.safeIndex(606, (d_257_v183_).length(0))
                (d_257_v183_)[index39_] = p0
                index40_ = default__.safeIndex(606, (d_257_v183_).length(0))
                (d_257_v183_)[index40_] = p0
                (globalState).f4 = p1
                index41_ = default__.safeIndex(606, (d_257_v183_).length(0))
                rhs27_ = (d_256_v182_) not in ((p0) + ((d_257_v183_)[default__.safeIndex(606, (d_257_v183_).length(0))]))
                rhs28_ = (_dafny.SeqWithoutIsStrInference([d_256_v182_ for d_262_i10_ in range(default__.abs(284))])) + (((d_257_v183_)[default__.safeIndex(606, (d_257_v183_).length(0))]) + (_dafny.SeqWithoutIsStrInference([d_256_v182_ for d_263_i11_ in range(default__.abs(272))])))
                rhs29_ = default__.fm7(globalState)
                rhs30_ = d_247_i4_
                lhs22_ = globalState
                lhs23_ = d_257_v183_
                lhs24_ = default__.safeIndex(606, (d_257_v183_).length(0))
                lhs25_ = globalState
                lhs22_.f4 = rhs27_
                lhs23_[lhs24_] = rhs28_
                d_256_v182_ = rhs29_
                lhs25_.f16 = rhs30_
            elif True:
                (globalState).f4 = True
                d_264_v184_: _dafny.Seq
                d_264_v184_ = _dafny.SeqWithoutIsStrInference([d_19_v1_])
                d_265_v185_: _dafny.Array
                nw48_ = _dafny.Array(None, 12)
                nw48_[int(0)] = (d_19_v1_).set(default__.safeIndex(d_18_v0_, len(d_19_v1_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))
                nw48_[int(1)] = d_19_v1_
                nw48_[int(2)] = (_dafny.SeqWithoutIsStrInference([331, d_247_i4_])) + (d_19_v1_)
                nw48_[int(3)] = (d_19_v1_) + (d_19_v1_)
                nw48_[int(4)] = d_19_v1_
                nw48_[int(5)] = _dafny.SeqWithoutIsStrInference([d_18_v0_])
                nw48_[int(6)] = ((d_264_v184_)[default__.safeIndex((d_19_v1_)[default__.safeIndex(d_18_v0_, len(d_19_v1_))], len(d_264_v184_))]).set(default__.safeIndex(d_247_i4_, len((d_264_v184_)[default__.safeIndex((d_19_v1_)[default__.safeIndex(d_18_v0_, len(d_19_v1_))], len(d_264_v184_))])), d_18_v0_)
                nw48_[int(7)] = (d_19_v1_) + (_dafny.SeqWithoutIsStrInference([(d_253_v179_).cardinality for d_266_i12_ in range(default__.abs(737))]))
                nw48_[int(8)] = d_19_v1_
                nw48_[int(9)] = d_19_v1_
                nw48_[int(10)] = _dafny.SeqWithoutIsStrInference([(d_253_v179_).cardinality, d_18_v0_])
                nw48_[int(11)] = (d_19_v1_) + (d_19_v1_)
                d_265_v185_ = nw48_
                index42_ = default__.safeIndex(838, (d_265_v185_).length(0))
                (d_265_v185_)[index42_] = d_19_v1_
                d_267_v186_: _dafny.MultiSet
                d_267_v186_ = _dafny.MultiSet([d_248_v176_])
                d_268_v187_: _dafny.MultiSet
                d_268_v187_ = _dafny.MultiSet([d_267_v186_, d_267_v186_, d_267_v186_])
                d_269_v188_: _dafny.Seq
                d_269_v188_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhknlfumr"))
                d_270_v189_: _dafny.Set
                d_270_v189_ = _dafny.Set({p1})
                d_271_v190_: _dafny.Map
                d_271_v190_ = _dafny.Map({d_18_v0_: d_247_i4_})
                d_272_v191_: _dafny.Map
                d_272_v191_ = _dafny.Map({d_271_v190_: d_270_v189_})
                index43_ = default__.safeIndex(838, (d_265_v185_).length(0))
                index44_ = default__.safeIndex(718, (d_249_v177_).length(0))
                index45_ = default__.safeIndex(718, (d_249_v177_).length(0))
                rhs31_ = d_19_v1_
                rhs32_ = p1
                rhs33_ = (d_270_v189_).ispropersubset((d_270_v189_).intersection(((d_272_v191_)[d_271_v190_] if (d_271_v190_) in (d_272_v191_) else d_270_v189_)))
                rhs34_ = d_268_v187_
                rhs35_ = (d_269_v188_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jduntgx")))
                lhs26_ = d_265_v185_
                lhs27_ = default__.safeIndex(838, (d_265_v185_).length(0))
                lhs28_ = d_249_v177_
                lhs29_ = default__.safeIndex(718, (d_249_v177_).length(0))
                lhs30_ = d_249_v177_
                lhs31_ = default__.safeIndex(718, (d_249_v177_).length(0))
                lhs26_[lhs27_] = rhs31_
                lhs28_[lhs29_] = rhs32_
                lhs30_[lhs31_] = rhs33_
                d_268_v187_ = rhs34_
                d_269_v188_ = rhs35_
                d_273_v192_: C0
                nw49_ = C0()
                nw49_.ctor__(len(d_270_v189_), False)
                d_273_v192_ = nw49_
                (globalState).f16 = len((d_270_v189_).intersection((d_270_v189_) | (d_270_v189_)))
                (globalState).f16 = default__.safeDivisionInt(default__.safeModuloInt(d_18_v0_, d_247_i4_), default__.safeDivisionInt(len(d_271_v190_), d_18_v0_))
        r0 = default__.safeDivisionInt(859, 457)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_274_v0_: bool
        d_274_v0_ = True
        d_275_v1_: _dafny.Seq
        d_275_v1_ = _dafny.SeqWithoutIsStrInference([d_274_v0_, d_274_v0_, d_274_v0_])
        d_276_v2_: _dafny.Array
        nw50_ = _dafny.Array(int(0), 15)
        d_276_v2_ = nw50_
        d_277_v3_: int
        d_277_v3_ = -235
        d_278_v4_: _dafny.Map
        d_278_v4_ = _dafny.Map({d_274_v0_: d_277_v3_})
        d_279_v5_: _dafny.Seq
        d_279_v5_ = _dafny.SeqWithoutIsStrInference([d_276_v2_])
        d_280_v6_: _dafny.Array
        nw51_ = _dafny.Array(None, 8)
        nw51_[int(0)] = not(d_274_v0_)
        nw51_[int(1)] = not(d_274_v0_)
        nw51_[int(2)] = d_274_v0_
        nw51_[int(3)] = d_274_v0_
        nw51_[int(4)] = True
        nw51_[int(5)] = d_274_v0_
        nw51_[int(6)] = d_274_v0_
        nw51_[int(7)] = d_274_v0_
        d_280_v6_ = nw51_
        d_281_v7_: _dafny.MultiSet
        d_281_v7_ = _dafny.MultiSet([d_280_v6_])
        d_282_globalState_: GlobalState
        nw52_ = GlobalState()
        nw52_.ctor__(485, -62, 525, False, False, 646, d_275_v1_, 462, d_276_v2_, 789, 610, d_278_v4_, False, (d_279_v5_)[default__.safeIndex(d_277_v3_, len(d_279_v5_))], d_280_v6_, (d_281_v7_) - (d_281_v7_), 789, _dafny.CodePoint('d'))
        d_282_globalState_ = nw52_
        d_283_v8_: _dafny.Map
        d_283_v8_ = _dafny.Map({d_274_v0_: d_274_v0_})
        d_284_v9_: _dafny.Set
        d_284_v9_ = _dafny.Set({d_274_v0_, ((d_283_v8_)[d_274_v0_] if (d_274_v0_) in (d_283_v8_) else False)})
        d_285_i0_: int
        d_285_i0_ = 0
        with _dafny.label("0"):
            while (d_275_v1_)[default__.safeIndex((0) - (len((d_284_v9_) - (d_284_v9_))), len(d_275_v1_))]:
                with _dafny.c_label("0"):
                    if (d_285_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_285_i0_ = (d_285_i0_) + (1)
                    d_286_v10_: _dafny.Array
                    nw53_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                    d_286_v10_ = nw53_
                    d_287_v11_: str
                    d_287_v11_ = _dafny.CodePoint('n')
                    index46_ = default__.safeIndex(53, (d_286_v10_).length(0))
                    (d_286_v10_)[index46_] = d_287_v11_
                    index47_ = default__.safeIndex(53, (d_286_v10_).length(0))
                    (d_286_v10_)[index47_] = d_287_v11_
                    d_288_v12_: _dafny.Seq
                    d_288_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jx"))
                    d_289_v13_: int
                    out0_: int
                    out0_ = default__.m0(d_288_v12_, not(not(not(d_274_v0_))), d_282_globalState_)
                    d_289_v13_ = out0_
                    d_290_v14_: int
                    out1_: int
                    out1_ = default__.m0(d_288_v12_, d_274_v0_, d_282_globalState_)
                    d_290_v14_ = out1_
                    (d_282_globalState_).f10 = d_290_v14_
                    pass
            pass
        index48_ = default__.safeIndex(460, (d_280_v6_).length(0))
        (d_280_v6_)[index48_] = not(not(d_274_v0_))
        index49_ = default__.safeIndex(460, (d_280_v6_).length(0))
        (d_280_v6_)[index49_] = not((d_274_v0_ if d_274_v0_ else d_274_v0_))
        if ((d_283_v8_)[False] if (False) in (d_283_v8_) else (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]):
            d_291_v15_: _dafny.Seq
            d_291_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
            (d_282_globalState_).f16 = len(d_291_v15_)
            d_292_v16_: int
            out2_: int
            out2_ = default__.m0(d_291_v15_, not(((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] if False else d_274_v0_)), d_282_globalState_)
            d_292_v16_ = out2_
            d_293_v17_: _dafny.Map
            d_293_v17_ = _dafny.Map({d_274_v0_: d_275_v1_})
            d_294_v18_: _dafny.Seq
            d_294_v18_ = _dafny.SeqWithoutIsStrInference([d_293_v17_, d_293_v17_])
            d_295_v19_: _dafny.Map
            d_295_v19_ = _dafny.Map({d_277_v3_: _dafny.Map({(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]: _dafny.SeqWithoutIsStrInference([True, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])})})
            index50_ = default__.safeIndex(460, (d_280_v6_).length(0))
            (d_280_v6_)[index50_] = ((d_293_v17_) | (d_293_v17_)) == ((((d_294_v18_)[default__.safeIndex((0) - (d_292_v16_), len(d_294_v18_))]).set(d_274_v0_, d_275_v1_)) | (((d_295_v19_)[default__.fm0(d_292_v16_, d_274_v0_, d_282_globalState_)] if (default__.fm0(d_292_v16_, d_274_v0_, d_282_globalState_)) in (d_295_v19_) else d_293_v17_)))
            d_296_v20_: str
            d_296_v20_ = _dafny.CodePoint('b')
            d_297_v21_: _dafny.Map
            d_297_v21_ = _dafny.Map({default__.fm1(d_296_v20_, d_282_globalState_): ((0) - (d_277_v3_)) < (d_292_v16_)})
            d_298_v22_: _dafny.MultiSet
            d_298_v22_ = _dafny.MultiSet([len(_dafny.Map({d_274_v0_: d_296_v20_}))])
            d_299_v23_: _dafny.Map
            d_299_v23_ = _dafny.Map({-695: d_298_v22_})
            d_274_v0_ = ((d_297_v21_)[d_299_v23_] if (d_299_v23_) in (d_297_v21_) else (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
            d_300_v24_: _dafny.Array
            def lambda5_(d_301_v0_):
                def lambda6_(d_302_i1_):
                    return _dafny.SeqWithoutIsStrInference([d_301_v0_])

                return lambda6_

            init3_ = lambda5_(d_274_v0_)
            nw54_ = _dafny.Array(None, 28)
            for i0_3_ in range(nw54_.length(0)):
                nw54_[i0_3_] = init3_(i0_3_)
            d_300_v24_ = nw54_
            d_303_v27_: _dafny.Array
            def lambda7_(d_304_v3_, d_305_v9_, d_306_v16_, d_307_v20_):
                def lambda8_(d_308_i2_):
                    def iife34_():
                        def iife35_():
                            coll7_ = _dafny.Map()
                            compr_7_: int
                            for compr_7_ in _dafny.IntegerRange(149, 819):
                                d_310_v26_: int = compr_7_
                                if ((149) <= (d_310_v26_)) and ((d_310_v26_) < (819)):
                                    coll7_[default__.safeDivisionInt(d_310_v26_, d_306_v16_)] = d_307_v20_
                            return _dafny.Map(coll7_)
                        coll6_ = _dafny.Map()
                        compr_6_: int
                        for compr_6_ in _dafny.IntegerRange(175, -248):
                            d_309_v25_: int = compr_6_
                            if ((175) <= (d_309_v25_)) and ((d_309_v25_) < (-248)):
                                coll6_[(d_309_v25_) + (len(iife35_()
                                ))] = d_305_v9_
                        return _dafny.Map(coll6_)
                    return (iife34_()
                    ) | (_dafny.Map({d_304_v3_: d_305_v9_}))

                return lambda8_

            init4_ = lambda7_(d_277_v3_, d_284_v9_, d_292_v16_, d_296_v20_)
            nw55_ = _dafny.Array(None, 1)
            for i0_4_ in range(nw55_.length(0)):
                nw55_[i0_4_] = init4_(i0_4_)
            d_303_v27_ = nw55_
            d_311_v29_: _dafny.Set
            d_311_v29_ = _dafny.Set({d_292_v16_})
            index51_ = default__.safeIndex(3, (d_303_v27_).length(0))
            def iife36_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (d_311_v29_).Elements:
                    d_312_v28_: int = compr_8_
                    if (d_312_v28_) in (d_311_v29_):
                        coll8_[default__.safeModuloInt(d_312_v28_, (0) - (d_292_v16_))] = _dafny.Set({d_274_v0_})
                return _dafny.Map(coll8_)
            (d_303_v27_)[index51_] = iife36_()
            
            d_313_v31_: _dafny.Map
            d_313_v31_ = _dafny.Map({d_292_v16_: d_284_v9_})
            index52_ = default__.safeIndex(460, (d_280_v6_).length(0))
            index53_ = default__.safeIndex(3, (d_303_v27_).length(0))
            def iife37_():
                coll9_ = _dafny.Set()
                compr_9_: int
                for compr_9_ in _dafny.IntegerRange(182, 539):
                    d_314_v30_: int = compr_9_
                    if ((182) <= (d_314_v30_)) and ((d_314_v30_) < (539)):
                        coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_314_v30_, d_292_v16_)]))
                return _dafny.Set(coll9_)
            rhs36_ = d_296_v20_
            rhs37_ = ((d_311_v29_ if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] else d_311_v29_)).issubset(iife37_()
            )
            rhs38_ = d_300_v24_
            rhs39_ = (d_313_v31_) | (d_313_v31_)
            rhs40_ = d_276_v2_
            lhs32_ = d_280_v6_
            lhs33_ = default__.safeIndex(460, (d_280_v6_).length(0))
            lhs34_ = d_303_v27_
            lhs35_ = default__.safeIndex(3, (d_303_v27_).length(0))
            d_296_v20_ = rhs36_
            lhs32_[lhs33_] = rhs37_
            d_300_v24_ = rhs38_
            lhs34_[lhs35_] = rhs39_
            d_276_v2_ = rhs40_
        elif True:
            d_315_v32_: _dafny.Seq
            d_315_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkkvwrsmk"))
            d_316_v33_: _dafny.Set
            d_316_v33_ = _dafny.Set({d_278_v4_})
            d_317_v34_: int
            out3_: int
            out3_ = default__.m0(d_315_v32_, (d_316_v33_).ispropersubset(d_316_v33_), d_282_globalState_)
            d_317_v34_ = out3_
            d_274_v0_ = True
            d_315_v32_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_318_i3_ in range(default__.abs(176))])) + (default__.fm2((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_274_v0_, d_282_globalState_))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_319_i4_ in range(default__.abs(-631))]))
            d_277_v3_ = (len(default__.fm2(d_274_v0_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_))) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_320_i5_ in range(default__.abs(62))])))
            (d_282_globalState_).f4 = default__.fm3(d_274_v0_, d_317_v34_, d_317_v34_, d_274_v0_, d_282_globalState_)
        d_321_v35_: _dafny.Seq
        d_321_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgyxvjy"))
        d_322_v36_: int
        out4_: int
        out4_ = default__.m0(d_321_v35_, (d_275_v1_)[default__.safeIndex(d_277_v3_, len(d_275_v1_))], d_282_globalState_)
        d_322_v36_ = out4_
        d_323_v37_: _dafny.Seq
        d_323_v37_ = _dafny.SeqWithoutIsStrInference([d_322_v36_, default__.fm0(d_322_v36_, not(False), d_282_globalState_), len(d_321_v35_), 135])
        index54_ = default__.safeIndex(460, (d_280_v6_).length(0))
        index55_ = default__.safeIndex(460, (d_280_v6_).length(0))
        rhs41_ = (d_321_v35_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryj")))
        rhs42_ = d_280_v6_
        rhs43_ = d_322_v36_
        rhs44_ = (default__.fm3(d_274_v0_, d_277_v3_, (0) - ((0) - (d_277_v3_)), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)) or ((default__.fm0(d_322_v36_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)) in (d_323_v37_))
        rhs45_ = (d_275_v1_)[default__.safeIndex(d_322_v36_, len(d_275_v1_))]
        lhs36_ = d_282_globalState_
        lhs37_ = d_280_v6_
        lhs38_ = default__.safeIndex(460, (d_280_v6_).length(0))
        lhs39_ = d_280_v6_
        lhs40_ = default__.safeIndex(460, (d_280_v6_).length(0))
        d_321_v35_ = rhs41_
        d_280_v6_ = rhs42_
        lhs36_.f0 = rhs43_
        lhs37_[lhs38_] = rhs44_
        lhs39_[lhs40_] = rhs45_
        if (d_321_v35_) <= (d_321_v35_):
            d_324_v38_: C0
            nw56_ = C0()
            nw56_.ctor__(606, d_274_v0_)
            d_324_v38_ = nw56_
            d_325_v39_: _dafny.Map
            d_325_v39_ = _dafny.Map({d_321_v35_: not(False)})
            d_326_v40_: _dafny.Map
            d_326_v40_ = _dafny.Map({d_322_v36_: (d_324_v38_).fm4(d_322_v36_, ((d_325_v39_)[d_321_v35_] if (d_321_v35_) in (d_325_v39_) else (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)})
            d_326_v40_ = (d_326_v40_).set(d_322_v36_, d_277_v3_)
            d_327_v41_: _dafny.Seq
            d_327_v41_ = _dafny.SeqWithoutIsStrInference([d_275_v1_])
            d_328_v42_: _dafny.Map
            d_328_v42_ = _dafny.Map({d_277_v3_: d_327_v41_})
            rhs46_ = (((((d_328_v42_)[d_322_v36_] if (d_322_v36_) in (d_328_v42_) else d_327_v41_)) + (d_327_v41_)).set(default__.safeIndex(-929, len((((d_328_v42_)[d_322_v36_] if (d_322_v36_) in (d_328_v42_) else d_327_v41_)) + (d_327_v41_))), default__.fm5(d_282_globalState_))).set(default__.safeIndex(d_277_v3_, len(((((d_328_v42_)[d_322_v36_] if (d_322_v36_) in (d_328_v42_) else d_327_v41_)) + (d_327_v41_)).set(default__.safeIndex(-929, len((((d_328_v42_)[d_322_v36_] if (d_322_v36_) in (d_328_v42_) else d_327_v41_)) + (d_327_v41_))), default__.fm5(d_282_globalState_)))), _dafny.SeqWithoutIsStrInference([d_274_v0_]))
            rhs47_ = d_280_v6_
            rhs48_ = (default__.fm0(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_329_i6_ in range(default__.abs(512))])), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)) >= (len((d_321_v35_) + (d_321_v35_)))
            lhs41_ = d_282_globalState_
            d_327_v41_ = rhs46_
            d_280_v6_ = rhs47_
            lhs41_.f4 = rhs48_
            d_330_v43_: _dafny.Map
            d_330_v43_ = _dafny.Map({d_323_v37_: (0) - (d_277_v3_)})
            d_331_v44_: _dafny.Map
            d_331_v44_ = _dafny.Map({default__.safeDivisionInt((0) - (len(d_323_v37_)), d_322_v36_): d_330_v43_})
            d_332_v45_: _dafny.MultiSet
            d_332_v45_ = _dafny.MultiSet([(d_323_v37_)[default__.safeIndex(d_277_v3_, len(d_323_v37_))], d_322_v36_, (d_323_v37_)[default__.safeIndex(95, len(d_323_v37_))], (0) - ((0) - (d_322_v36_))])
            d_331_v44_ = (d_331_v44_).set(d_322_v36_, (d_330_v43_).set(d_323_v37_, ((d_332_v45_)[d_277_v3_] if (d_277_v3_) in (d_332_v45_) else len(d_284_v9_))))
            source5_ = D0_DC1(d_274_v0_, _dafny.Map({d_274_v0_: (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]}), d_324_v38_, d_322_v36_)
            if source5_.is_DC0:
                d_333___mcc_h0_ = source5_.cf0
                d_334___mcc_h1_ = source5_.cf1
                d_335_cf1_ = d_334___mcc_h1_
                d_336_cf0_ = d_333___mcc_h0_
                d_324_v38_ = d_324_v38_
                (d_282_globalState_).f4 = (_dafny.SeqWithoutIsStrInference([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])) < (d_275_v1_)
                d_337_v46_: _dafny.Set
                d_337_v46_ = _dafny.Set({d_280_v6_, d_280_v6_, d_280_v6_, d_280_v6_})
                d_337_v46_ = _dafny.Set({d_280_v6_, d_280_v6_, d_280_v6_, d_280_v6_, d_280_v6_})
                d_338_v47_: _dafny.Map
                d_338_v47_ = _dafny.Map({d_274_v0_: d_279_v5_})
                d_339_v48_: D1
                d_339_v48_ = D1_DC4(d_276_v2_)
                d_338_v47_ = (d_338_v47_).set((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], _dafny.SeqWithoutIsStrInference([(d_339_v48_).cf10]))
            elif source5_.is_DC1:
                d_340___mcc_h2_ = source5_.cf2
                d_341___mcc_h3_ = source5_.cf3
                d_342___mcc_h4_ = source5_.cf4
                d_343___mcc_h5_ = source5_.cf5
                d_344_cf5_ = d_343___mcc_h5_
                d_345_cf4_ = d_342___mcc_h4_
                d_346_cf3_ = d_341___mcc_h3_
                d_347_cf2_ = d_340___mcc_h2_
                d_345_cf4_ = d_345_cf4_
                d_348_v49_: _dafny.Seq
                d_348_v49_ = _dafny.SeqWithoutIsStrInference([d_324_v38_, d_345_cf4_])
                d_348_v49_ = (d_348_v49_) + (_dafny.SeqWithoutIsStrInference([d_324_v38_]))
                d_283_v8_ = (d_283_v8_).set(d_347_cf2_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
                d_349_v50_: _dafny.Map
                d_349_v50_ = _dafny.Map({len(d_323_v37_): d_347_cf2_})
                index56_ = default__.safeIndex(460, (d_280_v6_).length(0))
                rhs49_ = ((d_323_v37_)[default__.safeIndex(248, len(d_323_v37_))]) <= (default__.safeModuloInt(d_322_v36_, d_344_cf5_))
                rhs50_ = ((d_349_v50_)[d_277_v3_] if (d_277_v3_) in (d_349_v50_) else (d_277_v3_) <= ((d_323_v37_)[default__.safeIndex(462, len(d_323_v37_))]))
                rhs51_ = True
                lhs42_ = d_280_v6_
                lhs43_ = default__.safeIndex(460, (d_280_v6_).length(0))
                lhs44_ = d_282_globalState_
                lhs42_[lhs43_] = rhs49_
                lhs44_.f4 = rhs50_
                d_347_cf2_ = rhs51_
            elif source5_.is_DC2:
                d_350___mcc_h6_ = source5_.cf6
                d_351___mcc_h7_ = source5_.cf7
                d_352___mcc_h8_ = source5_.cf8
                d_353_cf8_ = d_352___mcc_h8_
                d_354_cf7_ = d_351___mcc_h7_
                d_355_cf6_ = d_350___mcc_h6_
                d_356_v51_: int
                out5_: int
                out5_ = default__.m0(d_321_v35_, d_274_v0_, d_282_globalState_)
                d_356_v51_ = out5_
                (d_282_globalState_).f4 = (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]
                d_357_v52_: D0
                d_357_v52_ = D0_DC1(d_274_v0_, d_283_v8_, d_324_v38_, len(d_284_v9_))
                (d_282_globalState_).f4 = (d_357_v52_).cf2
                (d_282_globalState_).f10 = d_322_v36_
            elif True:
                d_358___mcc_h9_ = source5_.cf9
                d_359_cf9_ = d_358___mcc_h9_
                d_360_v53_: int
                out6_: int
                out6_ = default__.m0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tiftisjra")), default__.fm3(True, -562, d_277_v3_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_), d_282_globalState_)
                d_360_v53_ = out6_
                d_361_v54_: _dafny.Map
                d_361_v54_ = _dafny.Map({d_324_v38_: d_322_v36_})
                (d_282_globalState_).f10 = (d_322_v36_) * (len((d_361_v54_).set(d_324_v38_, d_360_v53_)))
                d_360_v53_ = 293
                (d_282_globalState_).f16 = d_360_v53_
        elif True:
            d_362_v56_: _dafny.Map
            def iife38_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(-75, 511):
                    d_363_v55_: int = compr_10_
                    if ((-75) <= (d_363_v55_)) and ((d_363_v55_) < (511)):
                        coll10_ = coll10_.union(_dafny.Set([default__.safeDivisionInt(d_363_v55_, len(d_323_v37_))]))
                return _dafny.Set(coll10_)
            d_362_v56_ = _dafny.Map({len(iife38_()
            ): d_322_v36_})
            d_364_v57_: D2
            d_364_v57_ = D2_DC9(d_362_v56_)
            d_365_v58_: T0
            nw57_ = C0()
            nw57_.ctor__(len((d_364_v57_).cf16), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
            d_365_v58_ = nw57_
            index57_ = default__.safeIndex(460, (d_280_v6_).length(0))
            rhs52_ = (not((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])) == ((d_362_v56_) != (d_362_v56_))
            rhs53_ = d_365_v58_
            rhs54_ = default__.safeModuloInt((d_277_v3_) + (d_365_v58_.f18), d_365_v58_.f18)
            rhs55_ = d_274_v0_
            rhs56_ = d_274_v0_
            lhs45_ = d_282_globalState_
            lhs46_ = d_282_globalState_
            lhs47_ = d_280_v6_
            lhs48_ = default__.safeIndex(460, (d_280_v6_).length(0))
            lhs45_.f4 = rhs52_
            d_365_v58_ = rhs53_
            d_277_v3_ = rhs54_
            lhs46_.f4 = rhs55_
            lhs47_[lhs48_] = rhs56_
            d_366_v59_: C0
            nw58_ = C0()
            nw58_.ctor__(len(d_275_v1_), (d_365_v58_).f19)
            d_366_v59_ = nw58_
            d_367_v60_: _dafny.Map
            d_367_v60_ = _dafny.Map({_dafny.CodePoint('m'): d_366_v59_})
            d_367_v60_ = d_367_v60_
            d_368_v61_: _dafny.Set
            d_368_v61_ = _dafny.Set({d_322_v36_})
            d_369_v62_: D0
            d_369_v62_ = D0_DC2(len(d_368_v61_), _dafny.MultiSet([(0) - (d_277_v3_), d_322_v36_, d_365_v58_.f18, 473, d_365_v58_.f18]), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
            rhs57_ = (d_369_v62_).cf8
            rhs58_ = d_323_v37_
            lhs49_ = d_282_globalState_
            lhs49_.f4 = rhs57_
            d_323_v37_ = rhs58_
            d_322_v36_ = default__.fm0(d_277_v3_, default__.fm3((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], len(d_321_v35_), len(d_284_v9_), (d_365_v58_).f19, d_282_globalState_), d_282_globalState_)
            d_370_v63_: D1
            d_370_v63_ = D1_DC5(800)
            index58_ = default__.safeIndex(460, (d_280_v6_).length(0))
            (d_280_v6_)[index58_] = default__.fm3(not(d_274_v0_), default__.fm0(len(d_278_v4_), d_274_v0_, d_282_globalState_), (d_370_v63_).cf11, not(False), d_282_globalState_)
        d_275_v1_ = ((d_275_v1_) + (d_275_v1_)) + ((_dafny.SeqWithoutIsStrInference([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])).set(default__.safeIndex(d_322_v36_, len(_dafny.SeqWithoutIsStrInference([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]]))), False))
        d_371_v64_: _dafny.Array
        nw59_ = _dafny.Array(D2.default()(), 7)
        d_371_v64_ = nw59_
        d_372_v65_: _dafny.Map
        d_372_v65_ = _dafny.Map({d_274_v0_: d_371_v64_})
        if ((d_372_v65_).set((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_371_v64_)) != ((d_372_v65_) | (d_372_v65_)):
            d_373_v66_: _dafny.Set
            d_373_v66_ = _dafny.Set({default__.fm6(-990, d_321_v35_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_), d_284_v9_})
            d_322_v36_ = len((d_373_v66_) | ((d_373_v66_).intersection(d_373_v66_)))
            d_323_v37_ = (d_323_v37_) + ((_dafny.SeqWithoutIsStrInference([250, d_322_v36_])) + (d_323_v37_))
            d_274_v0_ = d_274_v0_
            if not(((d_323_v37_) + (_dafny.SeqWithoutIsStrInference([408, d_322_v36_]))) <= (d_323_v37_)):
                (d_282_globalState_).f0 = ((d_278_v4_)[d_274_v0_] if (d_274_v0_) in (d_278_v4_) else (d_277_v3_) + (d_277_v3_))
                d_374_v67_: D1
                d_374_v67_ = D1_DC4(d_276_v2_)
                d_375_v68_: _dafny.Map
                d_375_v68_ = _dafny.Map({(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]: d_374_v67_})
                d_376_v69_: C0
                nw60_ = C0()
                nw60_.ctor__(d_322_v36_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
                d_376_v69_ = nw60_
                d_377_v70_: _dafny.Seq
                d_377_v70_ = _dafny.SeqWithoutIsStrInference([d_376_v69_])
                d_375_v68_ = (d_375_v68_).set((d_275_v1_)[default__.safeIndex(len(d_377_v70_), len(d_275_v1_))], D1_DC4(d_276_v2_))
                d_378_v71_: str
                d_378_v71_ = _dafny.CodePoint('o')
                d_379_v72_: _dafny.Map
                d_379_v72_ = _dafny.Map({d_277_v3_: d_378_v71_})
                d_380_v73_: _dafny.MultiSet
                d_380_v73_ = _dafny.MultiSet([d_277_v3_, (0) - ((0) - (len(d_379_v72_))), d_322_v36_, d_322_v36_, d_322_v36_])
                (d_282_globalState_).f0 = ((d_380_v73_).set(d_277_v3_, default__.abs(len(d_321_v35_)))).cardinality
                d_381_v74_: D1
                d_381_v74_ = D1_DC5(d_322_v36_)
                d_382_v75_: C0
                nw61_ = C0()
                nw61_.ctor__((0) - (((d_381_v74_).cf11) - (d_277_v3_)), d_274_v0_)
                d_382_v75_ = nw61_
                d_383_v76_: _dafny.Map
                d_383_v76_ = _dafny.Map({d_382_v75_: d_274_v0_})
                d_383_v76_ = (d_383_v76_).set((d_382_v75_ if d_274_v0_ else d_382_v75_), not(((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] else (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])))
            elif True:
                d_384_v77_: str
                d_384_v77_ = _dafny.CodePoint('o')
                (d_282_globalState_).f4 = (d_322_v36_) <= (default__.fm0(len((d_321_v35_).set(default__.safeIndex(-268, len(d_321_v35_)), d_384_v77_)), d_274_v0_, d_282_globalState_))
                d_385_v78_: _dafny.MultiSet
                d_385_v78_ = _dafny.MultiSet([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])
                d_386_v79_: _dafny.Set
                d_386_v79_ = _dafny.Set({d_277_v3_, (d_385_v78_).cardinality})
                d_387_v80_: C0
                nw62_ = C0()
                nw62_.ctor__((d_277_v3_) - ((D1_DC7(len(d_386_v79_), d_322_v36_)).cf13), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
                d_387_v80_ = nw62_
                d_388_v81_: _dafny.Array
                nw63_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_388_v81_ = nw63_
                d_389_v82_: _dafny.MultiSet
                d_389_v82_ = _dafny.MultiSet([d_277_v3_, d_322_v36_])
                d_390_v83_: _dafny.Map
                d_390_v83_ = _dafny.Map({((d_389_v82_)[d_277_v3_] if (d_277_v3_) in (d_389_v82_) else 125): d_389_v82_})
                d_391_v84_: _dafny.Map
                d_391_v84_ = _dafny.Map({(d_390_v83_) | (d_390_v83_): d_388_v81_})
                d_388_v81_ = ((d_391_v84_)[d_390_v83_] if (d_390_v83_) in (d_391_v84_) else d_388_v81_)
                d_392_v85_: _dafny.Array
                nw64_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
                d_392_v85_ = nw64_
                d_392_v85_ = d_392_v85_
                (d_282_globalState_).f10 = d_277_v3_
            d_393_v86_: _dafny.Array
            nw65_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
            d_393_v86_ = nw65_
            index59_ = default__.safeIndex(265, (d_393_v86_).length(0))
            (d_393_v86_)[index59_] = d_321_v35_
            d_394_v87_: C0
            nw66_ = C0()
            nw66_.ctor__(d_277_v3_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
            d_394_v87_ = nw66_
            d_395_v88_: _dafny.Seq
            d_395_v88_ = _dafny.SeqWithoutIsStrInference([d_321_v35_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hl"))) + (d_321_v35_)])
            d_396_v89_: _dafny.Seq
            d_396_v89_ = _dafny.SeqWithoutIsStrInference([d_394_v87_, d_394_v87_, d_394_v87_, d_394_v87_])
            index60_ = default__.safeIndex(265, (d_393_v86_).length(0))
            rhs59_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_397_i7_ in range(default__.abs(-560))])
            rhs60_ = (d_395_v88_)[default__.safeIndex((0) - (d_322_v36_), len(d_395_v88_))]
            rhs61_ = (d_396_v89_)[default__.safeIndex((d_322_v36_) + (len(d_323_v37_)), len(d_396_v89_))]
            rhs62_ = (0) - (d_322_v36_)
            lhs50_ = d_393_v86_
            lhs51_ = default__.safeIndex(265, (d_393_v86_).length(0))
            d_321_v35_ = rhs59_
            lhs50_[lhs51_] = rhs60_
            d_394_v87_ = rhs61_
            d_277_v3_ = rhs62_
        elif True:
            d_398_v90_: str
            d_398_v90_ = _dafny.CodePoint('n')
            (d_282_globalState_).f4 = (d_398_v90_) not in ((d_321_v35_).set(default__.safeIndex(786, len(d_321_v35_)), d_398_v90_))
            if not((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]):
                (d_282_globalState_).f16 = (d_322_v36_) + (d_277_v3_)
                d_399_v91_: _dafny.Array
                nw67_ = _dafny.Array(_dafny.Seq({}), 22)
                d_399_v91_ = nw67_
                index61_ = default__.safeIndex(585, (d_399_v91_).length(0))
                (d_399_v91_)[index61_] = (default__.fm5(d_282_globalState_)) + (d_275_v1_)
                index62_ = default__.safeIndex(585, (d_399_v91_).length(0))
                (d_399_v91_)[index62_] = _dafny.SeqWithoutIsStrInference([False, (d_321_v35_) == (d_321_v35_), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])
                d_400_v92_: C0
                nw68_ = C0()
                nw68_.ctor__(-238, ((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]) == (default__.fm3(False, d_322_v36_, d_322_v36_, d_274_v0_, d_282_globalState_)))
                d_400_v92_ = nw68_
                (d_282_globalState_).f10 = (d_400_v92_).fm4(d_322_v36_, d_274_v0_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)
                d_401_v93_: D1
                d_401_v93_ = D1_DC4(d_276_v2_)
                d_402_v94_: _dafny.Map
                d_402_v94_ = _dafny.Map({d_401_v93_: d_280_v6_})
                d_402_v94_ = (d_402_v94_).set(d_401_v93_, d_280_v6_)
            elif True:
                d_403_v95_: _dafny.MultiSet
                d_403_v95_ = _dafny.MultiSet([d_277_v3_])
                d_404_v96_: _dafny.MultiSet
                d_404_v96_ = _dafny.MultiSet([not(not((d_403_v95_).isdisjoint(d_403_v95_)))])
                d_404_v96_ = (d_404_v96_).intersection(d_404_v96_)
                d_398_v90_ = _dafny.CodePoint('k')
                d_405_v97_: _dafny.Array
                nw69_ = _dafny.Array(None, 22)
                nw69_[int(0)] = d_398_v90_
                nw69_[int(1)] = d_398_v90_
                nw69_[int(2)] = _dafny.CodePoint('t')
                nw69_[int(3)] = _dafny.CodePoint('k')
                nw69_[int(4)] = d_398_v90_
                nw69_[int(5)] = d_398_v90_
                nw69_[int(6)] = default__.fm7(d_282_globalState_)
                nw69_[int(7)] = d_398_v90_
                nw69_[int(8)] = d_398_v90_
                nw69_[int(9)] = d_398_v90_
                nw69_[int(10)] = d_398_v90_
                nw69_[int(11)] = d_398_v90_
                nw69_[int(12)] = d_398_v90_
                nw69_[int(13)] = d_398_v90_
                nw69_[int(14)] = d_398_v90_
                nw69_[int(15)] = d_398_v90_
                nw69_[int(16)] = d_398_v90_
                nw69_[int(17)] = d_398_v90_
                nw69_[int(18)] = _dafny.CodePoint('n')
                nw69_[int(19)] = d_398_v90_
                nw69_[int(20)] = d_398_v90_
                nw69_[int(21)] = d_398_v90_
                d_405_v97_ = nw69_
                d_406_v98_: _dafny.Map
                d_406_v98_ = _dafny.Map({d_274_v0_: d_405_v97_})
                d_406_v98_ = d_406_v98_
                index63_ = default__.safeIndex(460, (d_280_v6_).length(0))
                (d_280_v6_)[index63_] = (d_398_v90_) in ((_dafny.SeqWithoutIsStrInference([d_398_v90_ for d_407_i8_ in range(default__.abs(-809))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhsmdmwu"))))
                rhs63_ = True
                rhs64_ = ((d_403_v95_).intersection(d_403_v95_)).issubset(d_403_v95_)
                rhs65_ = d_280_v6_
                lhs52_ = d_282_globalState_
                d_274_v0_ = rhs63_
                lhs52_.f4 = rhs64_
                d_280_v6_ = rhs65_
            d_408_v99_: C0
            nw70_ = C0()
            nw70_.ctor__(651, (d_275_v1_) <= (_dafny.SeqWithoutIsStrInference([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])))
            d_408_v99_ = nw70_
            d_408_v99_ = d_408_v99_
            (d_282_globalState_).f10 = d_322_v36_
            d_409_v100_: _dafny.Map
            d_409_v100_ = _dafny.Map({d_322_v36_: d_321_v35_})
            d_410_v101_: _dafny.Seq
            d_410_v101_ = _dafny.SeqWithoutIsStrInference([d_321_v35_, d_321_v35_, _dafny.SeqWithoutIsStrInference([d_398_v90_ for d_411_i11_ in range(default__.abs(84))])])
            d_412_v102_: _dafny.Array
            nw71_ = _dafny.Array(None, 17)
            nw71_[int(0)] = d_321_v35_
            nw71_[int(1)] = (d_321_v35_) + (_dafny.SeqWithoutIsStrInference([d_398_v90_ for d_413_i9_ in range(default__.abs(957))]))
            nw71_[int(2)] = d_321_v35_
            nw71_[int(3)] = d_321_v35_
            nw71_[int(4)] = (((d_409_v100_)[d_322_v36_] if (d_322_v36_) in (d_409_v100_) else d_321_v35_) if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] else (d_321_v35_).set(default__.safeIndex(d_322_v36_, len(d_321_v35_)), d_398_v90_))
            nw71_[int(5)] = d_321_v35_
            nw71_[int(6)] = d_321_v35_
            nw71_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgyyvl"))) + (d_321_v35_)
            nw71_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llc"))
            nw71_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqphv"))
            nw71_[int(10)] = (d_321_v35_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npufrcjrg")))
            nw71_[int(11)] = (d_321_v35_) + (d_321_v35_)
            nw71_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sa"))
            nw71_[int(13)] = (d_321_v35_) + ((d_321_v35_).set(default__.safeIndex(d_277_v3_, len(d_321_v35_)), d_398_v90_))
            nw71_[int(14)] = (d_321_v35_ if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] else _dafny.SeqWithoutIsStrInference([d_398_v90_ for d_414_i10_ in range(default__.abs(923))]))
            nw71_[int(15)] = default__.fm2(not(not((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)
            nw71_[int(16)] = (d_410_v101_)[default__.safeIndex(d_277_v3_, len(d_410_v101_))]
            d_412_v102_ = nw71_
            index64_ = default__.safeIndex(353, (d_412_v102_).length(0))
            (d_412_v102_)[index64_] = d_321_v35_
            index65_ = default__.safeIndex(353, (d_412_v102_).length(0))
            (d_412_v102_)[index65_] = (d_321_v35_ if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))] else (d_321_v35_).set(default__.safeIndex(d_277_v3_, len(d_321_v35_)), d_398_v90_))
        d_415_v103_: _dafny.MultiSet
        d_415_v103_ = _dafny.MultiSet([(d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]])
        d_416_v104_: _dafny.Map
        d_416_v104_ = _dafny.Map({d_415_v103_: d_321_v35_})
        d_417_v105_: _dafny.Map
        d_417_v105_ = _dafny.Map({d_274_v0_: d_321_v35_})
        d_416_v104_ = ((d_416_v104_).set(d_415_v103_, d_321_v35_)).set(default__.fm8(len(d_283_v8_), D2_DC9(default__.fm9(d_274_v0_, d_282_globalState_)), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_321_v35_, d_282_globalState_), ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xifslf"))).set(default__.safeIndex(d_277_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xifslf")))), (d_321_v35_)[default__.safeIndex(d_322_v36_, len(d_321_v35_))])) + (((d_417_v105_)[d_274_v0_] if (d_274_v0_) in (d_417_v105_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))))
        d_418_i12_: int
        d_418_i12_ = 0
        with _dafny.label("1"):
            while d_274_v0_:
                with _dafny.c_label("1"):
                    if (d_418_i12_) >= (100):
                        raise _dafny.Break("1")
                    d_418_i12_ = (d_418_i12_) + (1)
                    d_419_v106_: _dafny.Array
                    nw72_ = _dafny.Array(_dafny.Map({}), 18)
                    d_419_v106_ = nw72_
                    d_420_v107_: _dafny.Array
                    def lambda9_(d_421_i13_):
                        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tof"))

                    init5_ = lambda9_
                    nw73_ = _dafny.Array(None, 6)
                    for i0_5_ in range(nw73_.length(0)):
                        nw73_[i0_5_] = init5_(i0_5_)
                    d_420_v107_ = nw73_
                    d_422_v108_: _dafny.Map
                    d_422_v108_ = _dafny.Map({d_419_v106_: d_420_v107_})
                    d_422_v108_ = (d_422_v108_).set(d_419_v106_, d_420_v107_)
                    d_423_v109_: str
                    d_423_v109_ = _dafny.CodePoint('r')
                    d_424_v110_: D2
                    d_424_v110_ = D2_DC11(d_423_v109_, d_322_v36_)
                    d_425_v111_: _dafny.Map
                    d_425_v111_ = _dafny.Map({d_424_v110_: d_321_v35_})
                    d_425_v111_ = (d_425_v111_).set(d_424_v110_, d_321_v35_)
                    d_322_v36_ = d_277_v3_
                    d_426_v112_: _dafny.Set
                    d_426_v112_ = _dafny.Set({_dafny.MultiSet(d_323_v37_)})
                    d_427_v113_: _dafny.MultiSet
                    d_427_v113_ = _dafny.MultiSet([d_322_v36_])
                    if (_dafny.Set({d_427_v113_, d_427_v113_})).ispropersubset(d_426_v112_):
                        d_423_v109_ = d_423_v109_
                        d_274_v0_ = (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]
                        d_428_v114_: _dafny.Map
                        d_428_v114_ = _dafny.Map({d_322_v36_: 802})
                        d_429_v115_: D0
                        d_429_v115_ = D0_DC2((0) - (((d_428_v114_)[d_277_v3_] if (d_277_v3_) in (d_428_v114_) else d_322_v36_)), d_427_v113_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
                        d_430_v116_: _dafny.Set
                        d_430_v116_ = _dafny.Set({D0_DC2(d_322_v36_, d_427_v113_, d_274_v0_)})
                        d_431_v117_: int
                        out7_: int
                        out7_ = default__.m0(d_321_v35_, not(not((_dafny.Set({d_429_v115_})).isdisjoint(d_430_v116_))), d_282_globalState_)
                        d_431_v117_ = out7_
                        d_432_v118_: _dafny.Map
                        d_432_v118_ = _dafny.Map({d_431_v117_: d_274_v0_})
                        d_433_v119_: C0
                        nw74_ = C0()
                        nw74_.ctor__((0) - (d_322_v36_), d_274_v0_)
                        d_433_v119_ = nw74_
                        d_434_v120_: _dafny.MultiSet
                        d_434_v120_ = _dafny.MultiSet([(d_433_v119_ if ((d_432_v118_)[d_431_v117_] if (d_431_v117_) in (d_432_v118_) else d_274_v0_) else d_433_v119_), d_433_v119_])
                        d_435_v121_: _dafny.Seq
                        d_435_v121_ = _dafny.SeqWithoutIsStrInference([d_433_v119_])
                        d_434_v120_ = (d_434_v120_).intersection(_dafny.MultiSet(d_435_v121_))
                        (d_282_globalState_).f16 = d_277_v3_
                    elif True:
                        d_436_v122_: C0
                        nw75_ = C0()
                        nw75_.ctor__(d_322_v36_, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
                        d_436_v122_ = nw75_
                        d_436_v122_ = d_436_v122_
                        (d_282_globalState_).f4 = d_274_v0_
                        (d_282_globalState_).f16 = d_277_v3_
                        d_437_v123_: _dafny.Array
                        nw76_ = _dafny.Array(None, 16)
                        d_437_v123_ = nw76_
                        d_438_v124_: D2
                        d_438_v124_ = D2_DC10()
                        d_439_v125_: _dafny.Map
                        d_439_v125_ = _dafny.Map({d_437_v123_: d_438_v124_})
                        (d_282_globalState_).f16 = len((d_439_v125_).set(d_437_v123_, d_438_v124_))
                        d_278_v4_ = (d_278_v4_).set((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_322_v36_)
                    pass
            pass
        (d_282_globalState_).f4 = d_274_v0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_280_v6_).length(0)):
            d_440_i14_: int = guard_loop_0_
            if (True) and (((0) <= (d_440_i14_)) and ((d_440_i14_) < ((d_280_v6_).length(0)))):
                (d_280_v6_)[(d_440_i14_)] = (len((d_321_v35_) + (d_321_v35_))) in ((_dafny.Map({d_322_v36_: False})) | (_dafny.Map({d_277_v3_: d_274_v0_})))
        index66_ = default__.safeIndex(460, (d_280_v6_).length(0))
        (d_280_v6_)[index66_] = ((d_277_v3_) == (len(((d_417_v105_)[False] if (False) in (d_417_v105_) else d_321_v35_)))) and (d_274_v0_)
        if (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]:
            d_281_v7_ = ((d_281_v7_ if d_274_v0_ else _dafny.MultiSet([d_280_v6_, d_280_v6_, d_280_v6_, d_280_v6_, d_280_v6_]))) - (d_281_v7_)
            if (d_277_v3_) != (-187):
                d_441_v126_: int
                out8_: int
                out8_ = default__.m0((d_321_v35_) + (d_321_v35_), (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))], d_282_globalState_)
                d_441_v126_ = out8_
                (d_282_globalState_).f6 = ((d_275_v1_) + (d_275_v1_)) + (_dafny.SeqWithoutIsStrInference([d_274_v0_]))
                (d_282_globalState_).f4 = d_274_v0_
                (d_282_globalState_).f4 = ((False if d_274_v0_ else (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])) and (not ((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]) or (d_274_v0_))
                d_442_v127_: C0
                nw77_ = C0()
                nw77_.ctor__(d_322_v36_, d_274_v0_)
                d_442_v127_ = nw77_
                d_443_v128_: _dafny.MultiSet
                d_443_v128_ = _dafny.MultiSet([563, (D0_DC1(d_274_v0_, d_283_v8_, d_442_v127_, d_277_v3_)).cf5])
                (d_282_globalState_).f4 = (d_443_v128_) != ((d_443_v128_) | (_dafny.MultiSet([d_322_v36_])))
            elif True:
                d_444_v129_: _dafny.Map
                d_444_v129_ = _dafny.Map({394: (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]})
                d_444_v129_ = d_444_v129_
                index67_ = default__.safeIndex(460, (d_280_v6_).length(0))
                (d_280_v6_)[index67_] = d_274_v0_
                index68_ = default__.safeIndex(460, (d_280_v6_).length(0))
                (d_280_v6_)[index68_] = (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]
                d_445_v130_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_445_v130_ = nw78_
                d_446_v131_: str
                d_446_v131_ = _dafny.CodePoint('b')
                d_447_v132_: _dafny.Seq
                d_447_v132_ = _dafny.SeqWithoutIsStrInference([d_445_v130_, d_445_v130_, d_445_v130_])
                d_448_v133_: C0
                nw79_ = C0()
                nw79_.ctor__(d_277_v3_, d_274_v0_)
                d_448_v133_ = nw79_
                d_449_v134_: _dafny.MultiSet
                d_449_v134_ = _dafny.MultiSet([d_448_v133_])
                index69_ = default__.safeIndex(460, (d_280_v6_).length(0))
                rhs66_ = (d_447_v132_)[default__.safeIndex((d_449_v134_).cardinality, len(d_447_v132_))]
                rhs67_ = (d_321_v35_) + (d_321_v35_)
                rhs68_ = True
                rhs69_ = len((_dafny.Set({not((d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))]), d_274_v0_}) if d_274_v0_ else d_284_v9_))
                rhs70_ = d_446_v131_
                lhs53_ = d_280_v6_
                lhs54_ = default__.safeIndex(460, (d_280_v6_).length(0))
                lhs55_ = d_282_globalState_
                d_445_v130_ = rhs66_
                d_321_v35_ = rhs67_
                lhs53_[lhs54_] = rhs68_
                lhs55_.f0 = rhs69_
                d_446_v131_ = rhs70_
                (d_282_globalState_).f0 = (len(d_275_v1_)) * (default__.safeModuloInt(len(d_321_v35_), d_277_v3_))
            d_450_v135_: C0
            nw80_ = C0()
            nw80_.ctor__(default__.safeDivisionInt(d_322_v36_, d_322_v36_), (d_277_v3_) > (d_277_v3_))
            d_450_v135_ = nw80_
            d_451_v136_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.MultiSet({}), 29)
            d_451_v136_ = nw81_
            index70_ = default__.safeIndex(74, (d_451_v136_).length(0))
            (d_451_v136_)[index70_] = d_415_v103_
            index71_ = default__.safeIndex(74, (d_451_v136_).length(0))
            (d_451_v136_)[index71_] = ((d_415_v103_).intersection(d_415_v103_)) - (d_415_v103_)
            d_452_v137_: _dafny.Array
            nw82_ = _dafny.Array(_dafny.Set({}), 4)
            d_452_v137_ = nw82_
            d_452_v137_ = d_452_v137_
        elif True:
            d_453_v138_: C0
            nw83_ = C0()
            nw83_.ctor__(268, (d_280_v6_)[default__.safeIndex(460, (d_280_v6_).length(0))])
            d_453_v138_ = nw83_
            index72_ = default__.safeIndex(460, (d_280_v6_).length(0))
            (d_280_v6_)[index72_] = d_274_v0_
            d_454_v139_: _dafny.MultiSet
            d_454_v139_ = _dafny.MultiSet([826])
            d_455_v140_: T0
            nw84_ = C0()
            nw84_.ctor__((d_454_v139_).cardinality, True)
            d_455_v140_ = nw84_
            d_456_v141_: str
            d_456_v141_ = _dafny.CodePoint('v')
            index73_ = default__.safeIndex(460, (d_280_v6_).length(0))
            rhs71_ = not(not(d_274_v0_))
            rhs72_ = d_455_v140_
            rhs73_ = d_456_v141_
            rhs74_ = not(not(d_274_v0_))
            rhs75_ = default__.safeDivisionInt(default__.fm0(default__.fm0(d_455_v140_.f18, (d_455_v140_).f19, d_282_globalState_), (d_455_v140_).f19, d_282_globalState_), len(_dafny.SeqWithoutIsStrInference([d_456_v141_, d_456_v141_, d_456_v141_, d_456_v141_])))
            lhs56_ = d_280_v6_
            lhs57_ = default__.safeIndex(460, (d_280_v6_).length(0))
            lhs58_ = d_282_globalState_
            lhs59_ = d_282_globalState_
            lhs56_[lhs57_] = rhs71_
            d_455_v140_ = rhs72_
            d_456_v141_ = rhs73_
            lhs58_.f4 = rhs74_
            lhs59_.f16 = rhs75_
            d_457_v142_: C0
            nw85_ = C0()
            nw85_.ctor__(d_322_v36_, d_274_v0_)
            d_457_v142_ = nw85_
            index74_ = default__.safeIndex(460, (d_280_v6_).length(0))
            (d_280_v6_)[index74_] = (d_321_v35_) == (d_321_v35_)
        d_321_v35_ = d_321_v35_
        d_458_v143_: _dafny.Seq
        d_458_v143_ = _dafny.SeqWithoutIsStrInference([d_321_v35_])
        d_459_v144_: int
        out9_: int
        out9_ = default__.m0((d_458_v143_)[default__.safeIndex(d_322_v36_, len(d_458_v143_))], (d_275_v1_) != (d_275_v1_), d_282_globalState_)
        d_459_v144_ = out9_
        _dafny.print(_dafny.string_of(d_274_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v1_) == (_dafny.SeqWithoutIsStrInference([True, True, True, True, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v4_) == (_dafny.Map({True: -235, False: -2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_279_v5_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_281_v7_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_282_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_282_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_.f6) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_282_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_.f11) == (_dafny.Map({True: -235}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f14)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_globalState_).f15).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_282_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v8_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_284_v9_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_285_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v35_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_322_v36_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_323_v37_) == (_dafny.SeqWithoutIsStrInference([1, 0, 7, 135, 250, 2, 1, 0, 7, 135]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_372_v65_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_415_v103_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_416_v104_) == (_dafny.Map({_dafny.MultiSet([False]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]), _dafny.MultiSet([True, True, False]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jifslfjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_417_v105_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_418_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_458_v143_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_459_v144_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Set({}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {self.cf1.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC5(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D1_DC8)

class D1_DC5(D1, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC8(D1, NamedTuple('DC8', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC8({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC8) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC10()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D2_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)

class D2_DC10(D2, NamedTuple('DC10', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC11(D2, NamedTuple('DC11', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC13(False, False, D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D3_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D3_DC14)

class D3_DC13(D3, NamedTuple('DC13', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC13({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC13) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC14(D3, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC15(D4, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(D2.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)

class D5_DC17(D5, NamedTuple('DC17', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value

class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f4: bool = False
        self.f6: _dafny.Seq = _dafny.Seq({})
        self.f10: int = int(0)
        self.f11: _dafny.Map = _dafny.Map({})
        self.f13: _dafny.Array = _dafny.Array(None, 0)
        self.f16: int = int(0)
        self._f1: int = int(0)
        self._f2: int = int(0)
        self._f3: bool = False
        self._f5: int = int(0)
        self._f7: int = int(0)
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: int = int(0)
        self._f12: bool = False
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f17: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f5(self):
        return self._f5
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17

class C0(T0):
    def  __init__(self):
        self._f18: int = int(0)
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f18(self):
        return self._f18
    @f18.setter
    def f18(self, value):
        self._f18 = value
    @property
    def f19(self):
        return self._f19
    def ctor__(self, f18, f19):
        (self).f18 = f18
        (self)._f19 = f19

    def fm4(self, p0, p1, p2, globalState):
        return (0) - (self.f18)

