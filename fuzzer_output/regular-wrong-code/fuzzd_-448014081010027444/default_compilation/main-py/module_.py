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
    def fm0(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, True])) + ((_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([not(True), True, not(True)])))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        if False:
            return True
        elif True:
            return (not(not(not(not(not(not(not(False)))))))) and (True)

    @staticmethod
    def fm2(p0, p1, globalState):
        return ((_dafny.Map({(0) - (len(_dafny.Set({True, True, False}))): False})) | (_dafny.Map({463: False}))) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([308, len(_dafny.Set({930, 305})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kie"))), 941, (_dafny.MultiSet([D1_DC2(len(_dafny.Map({False: _dafny.MultiSet([True])})), True, True)])).cardinality])): True})) | (_dafny.Map({-918: True})))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([True, True])) - (_dafny.MultiSet([False]))).intersection(_dafny.MultiSet([True]))

    @staticmethod
    def fm4(globalState):
        if False:
            return -885
        elif True:
            return len(_dafny.SeqWithoutIsStrInference([True, not(False), True, False]))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return (((D6_DC18(_dafny.Set({False}))).cf27) | (_dafny.Set({False, True}))) - ((_dafny.Set({False})).intersection(_dafny.Set({not(True), not(not(False))})))

    @staticmethod
    def fm6(p0, p1, globalState):
        source0_ = D2_DC7(True, 462)
        if source0_.is_DC7:
            d_0___mcc_h0_ = source0_.cf11
            d_1___mcc_h1_ = source0_.cf12
            d_2_cf12_ = d_1___mcc_h1_
            d_3_cf11_ = d_0___mcc_h0_
            return _dafny.Map({D1_DC2(127, True, d_3_cf11_): d_3_cf11_})
        elif source0_.is_DC6:
            d_4___mcc_h2_ = source0_.cf10
            d_5_cf10_ = d_4___mcc_h2_
            def iife0_():
                coll0_ = _dafny.Map()
                compr_0_: int
                for compr_0_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdogpmi"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_6_i0_ in range(default__.abs(713))]))})).keys.Elements:
                    d_7_v0_: int = compr_0_
                    if (d_7_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdogpmi"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_6_i0_ in range(default__.abs(713))]))})):
                        coll0_[default__.safeModuloInt(d_7_v0_, 857)] = True
                return _dafny.Map(coll0_)
            return _dafny.Map({D1_DC2(len(_dafny.Map({_dafny.Set({len(iife0_()
)}): 969})), False, not(False)): False})
        elif True:
            d_8___mcc_h3_ = source0_.cf13
            d_9_cf13_ = d_8___mcc_h3_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: D1
                for compr_1_ in (_dafny.Set({D1_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogqm"))), True, True), D1_DC2(-484, not(False), True)})).Elements:
                    d_10_v1_: D1 = compr_1_
                    if (d_10_v1_) in (_dafny.Set({D1_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogqm"))), True, True), D1_DC2(-484, not(False), True)})):
                        coll1_[d_10_v1_] = True
                return _dafny.Map(coll1_)
            return iife1_()
            

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([589, (0) - (len(_dafny.Set({645}))), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_11_i0_ in range(default__.abs(-720))])))])) + (_dafny.SeqWithoutIsStrInference([44 for d_12_i1_ in range(default__.abs(-525))]))) + ((_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([False, not(not(not(True)))])).cardinality)])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqtumvw")))])))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrikrvux"))) + ((D4_DC13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ur")))).cf20)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hytv")))

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.CodePoint('t')

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return (D9_DC24(_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({False: 878}))]))).cf38

    @staticmethod
    def fm11(globalState):
        return _dafny.Set({_dafny.CodePoint('k'), _dafny.CodePoint('x'), _dafny.CodePoint('n'), _dafny.CodePoint('x'), _dafny.CodePoint('l')})

    @staticmethod
    def fm12(globalState):
        return D4_DC13((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urylnfh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mimbr"))))

    @staticmethod
    def fm13(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-208, 303):
                d_13_v0_: int = compr_2_
                if ((-208) <= (d_13_v0_)) and ((d_13_v0_) < (303)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeDivisionInt(d_13_v0_, 869)]))
            return _dafny.Set(coll2_)
        return iife2_()
        

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_14_v0_: _dafny.MultiSet
        d_14_v0_ = _dafny.MultiSet([True])
        d_15_v1_: _dafny.Seq
        d_15_v1_ = _dafny.SeqWithoutIsStrInference([d_14_v0_])
        if (((d_15_v1_)[default__.safeIndex((0) - (p1), len(d_15_v1_))]) | (d_14_v0_)).ispropersubset((_dafny.MultiSet([p2, not(not(p2)), p2, False])).intersection(d_14_v0_)):
            if (default__.safeDivisionInt(p1, p1)) not in (_dafny.MultiSet([p1])):
                d_16_v2_: _dafny.Set
                d_16_v2_ = _dafny.Set({not(p0), p0})
                r1 = default__.safeDivisionInt((p1) + (p1), len(d_16_v2_))
                d_17_v3_: _dafny.Array
                nw0_ = _dafny.Array(int(0), 15)
                d_17_v3_ = nw0_
                index0_ = default__.safeIndex(1, (d_17_v3_).length(0))
                (d_17_v3_)[index0_] = p1
                index1_ = default__.safeIndex(1, (d_17_v3_).length(0))
                (d_17_v3_)[index1_] = p1
                index2_ = default__.safeIndex(1, (d_17_v3_).length(0))
                (d_17_v3_)[index2_] = p1
                d_18_v4_: _dafny.Array
                def lambda0_(d_19_p0_):
                    def lambda1_(d_20_i0_):
                        return d_19_p0_

                    return lambda1_

                init0_ = lambda0_(p0)
                nw1_ = _dafny.Array(None, 10)
                for i0_0_ in range(nw1_.length(0)):
                    nw1_[i0_0_] = init0_(i0_0_)
                d_18_v4_ = nw1_
                d_21_v5_: C0
                nw2_ = C0()
                nw2_.ctor__(d_18_v4_)
                d_21_v5_ = nw2_
                r2 = p2
            elif True:
                d_22_v6_: _dafny.Seq
                d_22_v6_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_23_v7_: _dafny.Seq
                d_23_v7_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_22_v6_)).cardinality])
                d_24_v8_: D3
                d_24_v8_ = D3_DC11(_dafny.SeqWithoutIsStrInference([p1 for d_25_i1_ in range(default__.abs(835))]), p2)
                d_26_v9_: _dafny.Seq
                d_26_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwnjkq"))
                d_27_v10_: _dafny.Seq
                d_27_v10_ = _dafny.SeqWithoutIsStrInference([d_23_v7_])
                d_28_v11_: str
                d_28_v11_ = _dafny.CodePoint('r')
                d_29_v12_: _dafny.Array
                nw3_ = _dafny.Array(None, 18)
                nw3_[int(0)] = d_23_v7_
                nw3_[int(1)] = (d_24_v8_).cf17
                nw3_[int(2)] = (_dafny.SeqWithoutIsStrInference([p1 for d_30_i2_ in range(default__.abs(-640))])) + (d_23_v7_)
                nw3_[int(3)] = d_23_v7_
                nw3_[int(4)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1, p1])
                nw3_[int(5)] = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])
                nw3_[int(6)] = d_23_v7_
                nw3_[int(7)] = (_dafny.SeqWithoutIsStrInference([p1 for d_31_i3_ in range(default__.abs(880))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([p1 for d_32_i3_ in range(default__.abs(880))]))), p1)
                nw3_[int(8)] = (d_23_v7_).set(default__.safeIndex(p1, len(d_23_v7_)), len(d_26_v9_))
                nw3_[int(9)] = (d_27_v10_)[default__.safeIndex(p1, len(d_27_v10_))]
                nw3_[int(10)] = (d_23_v7_).set(default__.safeIndex(default__.fm4(globalState), len(d_23_v7_)), -246)
                nw3_[int(11)] = _dafny.SeqWithoutIsStrInference([882 for d_33_i4_ in range(default__.abs(374))])
                nw3_[int(12)] = d_23_v7_
                nw3_[int(13)] = d_23_v7_
                nw3_[int(14)] = d_23_v7_
                nw3_[int(15)] = d_23_v7_
                nw3_[int(16)] = (d_23_v7_) + (d_23_v7_)
                nw3_[int(17)] = default__.fm7(p2, d_28_v11_, 349, globalState)
                d_29_v12_ = nw3_
                index3_ = default__.safeIndex(938, (d_29_v12_).length(0))
                (d_29_v12_)[index3_] = d_23_v7_
                index4_ = default__.safeIndex(938, (d_29_v12_).length(0))
                (d_29_v12_)[index4_] = d_23_v7_
                d_34_v13_: D3
                d_34_v13_ = D3_DC9((d_29_v12_)[default__.safeIndex(938, (d_29_v12_).length(0))])
                d_34_v13_ = d_34_v13_
                d_35_v14_: _dafny.Array
                nw4_ = _dafny.Array(None, 8)
                nw4_[int(0)] = default__.fm9(p2, globalState)
                nw4_[int(1)] = d_28_v11_
                nw4_[int(2)] = d_28_v11_
                nw4_[int(3)] = (d_28_v11_ if p2 else _dafny.CodePoint('v'))
                nw4_[int(4)] = d_28_v11_
                nw4_[int(5)] = d_28_v11_
                nw4_[int(6)] = d_28_v11_
                nw4_[int(7)] = d_28_v11_
                d_35_v14_ = nw4_
                index5_ = default__.safeIndex(896, (d_35_v14_).length(0))
                (d_35_v14_)[index5_] = d_28_v11_
                d_36_v15_: _dafny.Array
                def lambda2_(d_37_p1_):
                    def lambda3_(d_38_i5_):
                        return default__.safeModuloInt(d_38_i5_, d_37_p1_)

                    return lambda3_

                init1_ = lambda2_(p1)
                nw5_ = _dafny.Array(None, 19)
                for i0_1_ in range(nw5_.length(0)):
                    nw5_[i0_1_] = init1_(i0_1_)
                d_36_v15_ = nw5_
                index6_ = default__.safeIndex(896, (d_35_v14_).length(0))
                rhs0_ = p0
                rhs1_ = d_28_v11_
                rhs2_ = d_36_v15_
                lhs0_ = d_35_v14_
                lhs1_ = default__.safeIndex(896, (d_35_v14_).length(0))
                r2 = rhs0_
                lhs0_[lhs1_] = rhs1_
                d_36_v15_ = rhs2_
                d_39_v16_: _dafny.Array
                def lambda4_(d_40_v0_):
                    def lambda5_(d_41_i6_):
                        return d_40_v0_

                    return lambda5_

                init2_ = lambda4_(d_14_v0_)
                nw6_ = _dafny.Array(None, 29)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_39_v16_ = nw6_
                index7_ = default__.safeIndex(952, (d_39_v16_).length(0))
                (d_39_v16_)[index7_] = d_14_v0_
                index8_ = default__.safeIndex(952, (d_39_v16_).length(0))
                (d_39_v16_)[index8_] = d_14_v0_
                d_42_v17_: _dafny.Array
                nw7_ = _dafny.Array(False, 16)
                d_42_v17_ = nw7_
                d_43_v18_: C0
                nw8_ = C0()
                nw8_.ctor__(d_42_v17_)
                d_43_v18_ = nw8_
            d_44_v19_: _dafny.Map
            d_44_v19_ = _dafny.Map({(0) - (p1): p0})
            d_45_v20_: _dafny.Seq
            d_45_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxdguiwrc"))
            d_46_v21_: D4
            d_46_v21_ = D4_DC13(d_45_v20_)
            d_47_v22_: _dafny.Seq
            d_47_v22_ = _dafny.SeqWithoutIsStrInference([p0, False, p0])
            d_48_v23_: _dafny.Map
            d_48_v23_ = _dafny.Map({d_47_v22_: p0})
            d_49_v24_: D3
            d_49_v24_ = D3_DC11(_dafny.SeqWithoutIsStrInference([p1 for d_50_i7_ in range(default__.abs(24))]), p0)
            d_51_v25_: _dafny.Array
            nw9_ = _dafny.Array(None, 29)
            nw9_[int(0)] = False
            nw9_[int(1)] = p0
            nw9_[int(2)] = default__.fm1(d_44_v19_, (d_46_v21_).cf20, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lenrsoqb")), globalState)
            nw9_[int(3)] = p2
            nw9_[int(4)] = p0
            nw9_[int(5)] = p0
            nw9_[int(6)] = p0
            nw9_[int(7)] = p2
            nw9_[int(8)] = p0
            nw9_[int(9)] = p0
            nw9_[int(10)] = p2
            nw9_[int(11)] = not(not(p2))
            nw9_[int(12)] = True
            nw9_[int(13)] = not(p0)
            nw9_[int(14)] = False
            nw9_[int(15)] = p2
            nw9_[int(16)] = p2
            nw9_[int(17)] = p0
            nw9_[int(18)] = ((d_48_v23_)[d_47_v22_] if (d_47_v22_) in (d_48_v23_) else not(p0))
            nw9_[int(19)] = True
            nw9_[int(20)] = p2
            nw9_[int(21)] = not(p0)
            nw9_[int(22)] = (d_49_v24_).cf18
            nw9_[int(23)] = p0
            nw9_[int(24)] = p2
            nw9_[int(25)] = p0
            nw9_[int(26)] = p2
            nw9_[int(27)] = not(p2)
            nw9_[int(28)] = p0
            d_51_v25_ = nw9_
            d_52_v26_: C0
            nw10_ = C0()
            nw10_.ctor__(d_51_v25_)
            d_52_v26_ = nw10_
            d_53_v27_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Seq({}), 13)
            d_53_v27_ = nw11_
            d_54_v28_: _dafny.Array
            def lambda6_(d_55_p1_):
                def lambda7_(d_56_i8_):
                    return (d_56_i8_) + (d_55_p1_)

                return lambda7_

            init3_ = lambda6_(p1)
            nw12_ = _dafny.Array(None, 21)
            for i0_3_ in range(nw12_.length(0)):
                nw12_[i0_3_] = init3_(i0_3_)
            d_54_v28_ = nw12_
            d_57_v29_: _dafny.Seq
            d_57_v29_ = _dafny.SeqWithoutIsStrInference([d_54_v28_])
            index9_ = default__.safeIndex(305, (d_53_v27_).length(0))
            (d_53_v27_)[index9_] = d_57_v29_
            index10_ = default__.safeIndex(305, (d_53_v27_).length(0))
            (d_53_v27_)[index10_] = d_57_v29_
            d_58_v30_: C0
            nw13_ = C0()
            nw13_.ctor__((d_52_v26_).f15)
            d_58_v30_ = nw13_
            (globalState).f0 = default__.safeModuloInt(p1, (0) - (p1))
        elif True:
            if not(p0):
                d_59_v31_: _dafny.Array
                nw14_ = _dafny.Array(False, 16)
                d_59_v31_ = nw14_
                index11_ = default__.safeIndex(359, (d_59_v31_).length(0))
                (d_59_v31_)[index11_] = p2
                index12_ = default__.safeIndex(359, (d_59_v31_).length(0))
                (d_59_v31_)[index12_] = p2
                d_60_v32_: C0
                nw15_ = C0()
                nw15_.ctor__(d_59_v31_)
                d_60_v32_ = nw15_
                d_60_v32_ = d_60_v32_
                d_61_v33_: C0
                nw16_ = C0()
                nw16_.ctor__((d_60_v32_).f15)
                d_61_v33_ = nw16_
                d_62_v34_: _dafny.Set
                d_62_v34_ = _dafny.Set({p0, p0, p2, p2, (d_59_v31_)[default__.safeIndex(359, (d_59_v31_).length(0))]})
                d_63_v35_: _dafny.MultiSet
                d_63_v35_ = _dafny.MultiSet([d_62_v34_, d_62_v34_, d_62_v34_])
                d_64_v36_: D5
                d_64_v36_ = D5_DC15(d_63_v35_)
                index13_ = default__.safeIndex(359, (d_59_v31_).length(0))
                (d_59_v31_)[index13_] = ((d_64_v36_).cf23).isdisjoint(_dafny.MultiSet([d_62_v34_]))
                d_65_v37_: _dafny.Array
                nw17_ = _dafny.Array(_dafny.Set({}), 5)
                d_65_v37_ = nw17_
                index14_ = default__.safeIndex(540, (d_65_v37_).length(0))
                (d_65_v37_)[index14_] = default__.fm11(globalState)
                d_66_v38_: _dafny.Set
                d_66_v38_ = _dafny.Set({_dafny.CodePoint('f')})
                d_67_v39_: str
                d_67_v39_ = _dafny.CodePoint('v')
                index15_ = default__.safeIndex(540, (d_65_v37_).length(0))
                (d_65_v37_)[index15_] = ((d_66_v38_).intersection(_dafny.Set({_dafny.CodePoint('y'), d_67_v39_, d_67_v39_}))) | (d_66_v38_)
            elif True:
                d_68_v40_: D5
                d_68_v40_ = D5_DC17(p2)
                pat_let_tv0_ = p0
                def iife3_(_pat_let0_0):
                    def iife4_(d_69_dt__update__tmp_h0_):
                        def iife5_(_pat_let1_0):
                            def iife6_(d_70_dt__update_hcf26_h0_):
                                return D5_DC17(d_70_dt__update_hcf26_h0_)
                            return iife6_(_pat_let1_0)
                        return iife5_(pat_let_tv0_)
                    return iife4_(_pat_let0_0)
                d_68_v40_ = (iife3_(d_68_v40_) if True else d_68_v40_)
                d_71_v41_: _dafny.Array
                nw18_ = _dafny.Array(int(0), 10)
                d_71_v41_ = nw18_
                index16_ = default__.safeIndex(913, (d_71_v41_).length(0))
                (d_71_v41_)[index16_] = p1
                index17_ = default__.safeIndex(913, (d_71_v41_).length(0))
                (d_71_v41_)[index17_] = ((0) - (p1)) - (-577)
                d_72_v42_: _dafny.Array
                nw19_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_72_v42_ = nw19_
                d_73_v43_: _dafny.Seq
                d_73_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctnu"))
                index18_ = default__.safeIndex(658, (d_72_v42_).length(0))
                (d_72_v42_)[index18_] = d_73_v43_
                index19_ = default__.safeIndex(658, (d_72_v42_).length(0))
                (d_72_v42_)[index19_] = d_73_v43_
                d_74_v44_: _dafny.Array
                def lambda8_(d_75_p0_):
                    def lambda9_(d_76_i9_):
                        return d_75_p0_

                    return lambda9_

                init4_ = lambda8_(p0)
                nw20_ = _dafny.Array(None, 24)
                for i0_4_ in range(nw20_.length(0)):
                    nw20_[i0_4_] = init4_(i0_4_)
                d_74_v44_ = nw20_
                d_77_v45_: C0
                nw21_ = C0()
                nw21_.ctor__(d_74_v44_)
                d_77_v45_ = nw21_
                d_78_v46_: str
                d_78_v46_ = _dafny.CodePoint('r')
                rhs3_ = default__.fm8(not (p0) or (p0), d_78_v46_, d_78_v46_, globalState)
                rhs4_ = p1
                lhs2_ = globalState
                lhs3_ = globalState
                lhs2_.f5 = rhs3_
                lhs3_.f0 = rhs4_
            d_79_v47_: _dafny.Map
            d_79_v47_ = _dafny.Map({p1: p2})
            d_80_v48_: _dafny.Array
            nw22_ = _dafny.Array(None, 4)
            nw22_[int(0)] = p0
            nw22_[int(1)] = p2
            nw22_[int(2)] = p2
            nw22_[int(3)] = ((d_79_v47_)[p1] if (p1) in (d_79_v47_) else p2)
            d_80_v48_ = nw22_
            d_81_v49_: C0
            nw23_ = C0()
            nw23_.ctor__(d_80_v48_)
            d_81_v49_ = nw23_
            d_82_v50_: _dafny.Map
            d_82_v50_ = _dafny.Map({p1: (d_79_v47_) | (d_79_v47_)})
            d_83_v51_: _dafny.MultiSet
            d_83_v51_ = _dafny.MultiSet([d_81_v49_, d_81_v49_])
            d_82_v50_ = (d_82_v50_).set(((d_83_v51_).cardinality) * (p1), ((d_82_v50_)[70] if (70) in (d_82_v50_) else d_79_v47_))
            d_84_v52_: C0
            nw24_ = C0()
            nw24_.ctor__((d_81_v49_).f15)
            d_84_v52_ = nw24_
            d_84_v52_ = d_84_v52_
            index20_ = default__.safeIndex(186, ((d_84_v52_).f15).length(0))
            ((d_84_v52_).f15)[index20_] = p2
            d_85_v53_: _dafny.Seq
            d_85_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuolsr"))
            d_86_v54_: _dafny.MultiSet
            d_86_v54_ = _dafny.MultiSet([d_85_v53_])
            index21_ = default__.safeIndex(186, ((d_84_v52_).f15).length(0))
            ((d_84_v52_).f15)[index21_] = (d_86_v54_) != (d_86_v54_)
        if True:
            r1 = default__.fm4(globalState)
            source1_ = default__.fm12(globalState)
            if source1_.is_DC14:
                d_87___mcc_h0_ = source1_.cf21
                d_88___mcc_h1_ = source1_.cf22
                d_89_cf22_ = d_88___mcc_h1_
                d_90_cf21_ = d_87___mcc_h0_
                r2 = p2
                r2 = p2
                d_91_v55_: D1
                d_91_v55_ = D1_DC1((0) - (p1))
                d_92_v56_: D2
                d_92_v56_ = D2_DC7(p2, (d_91_v55_).cf1)
                d_92_v56_ = d_92_v56_
                d_93_v57_: _dafny.Map
                d_93_v57_ = _dafny.Map({p2: p2})
                d_94_v58_: _dafny.MultiSet
                d_94_v58_ = _dafny.MultiSet([_dafny.Set({521})])
                d_95_v59_: _dafny.Set
                d_95_v59_ = _dafny.Set({p1, -743})
                d_93_v57_ = (d_93_v57_).set(True, (d_94_v58_).isdisjoint(_dafny.MultiSet([d_95_v59_, d_95_v59_])))
            elif True:
                d_96___mcc_h2_ = source1_.cf20
                d_97_cf20_ = d_96___mcc_h2_
                d_98_v60_: _dafny.Map
                d_98_v60_ = _dafny.Map({p1: len(_dafny.Map({316: ((_dafny.MultiSet([False])).set(p0, default__.abs(p1))).set(p0, default__.abs(884))}))})
                d_99_v61_: str
                d_99_v61_ = _dafny.CodePoint('d')
                d_100_v62_: _dafny.Map
                d_100_v62_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyww")): p0})
                d_101_v63_: _dafny.Seq
                d_101_v63_ = _dafny.SeqWithoutIsStrInference([not(p0), p2, default__.fm1(_dafny.Map({len(d_98_v60_): p0}), (d_97_cf20_).set(default__.safeIndex((0) - (p1), len(d_97_cf20_)), d_99_v61_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rswrhg")), globalState), ((d_100_v62_)[(d_97_cf20_).set(default__.safeIndex(p1, len(d_97_cf20_)), d_99_v61_)] if ((d_97_cf20_).set(default__.safeIndex(p1, len(d_97_cf20_)), d_99_v61_)) in (d_100_v62_) else p0)])
                d_102_v64_: _dafny.Map
                d_102_v64_ = _dafny.Map({p0: d_101_v63_})
                d_102_v64_ = d_102_v64_
                d_103_v65_: _dafny.Array
                nw25_ = _dafny.Array(_dafny.Map({}), 20)
                d_103_v65_ = nw25_
                index22_ = default__.safeIndex(313, (d_103_v65_).length(0))
                (d_103_v65_)[index22_] = _dafny.Map({p0: p2})
                d_104_v66_: _dafny.Map
                d_104_v66_ = _dafny.Map({not(p0): (p0 if True else p2)})
                index23_ = default__.safeIndex(313, (d_103_v65_).length(0))
                (d_103_v65_)[index23_] = d_104_v66_
                d_105_v67_: _dafny.Array
                nw26_ = _dafny.Array(False, 23)
                d_105_v67_ = nw26_
                index24_ = default__.safeIndex(603, (d_105_v67_).length(0))
                (d_105_v67_)[index24_] = p0
                index25_ = default__.safeIndex(603, (d_105_v67_).length(0))
                (d_105_v67_)[index25_] = (989) != (p1)
                d_106_v68_: _dafny.Map
                d_106_v68_ = _dafny.Map({(d_105_v67_)[default__.safeIndex(603, (d_105_v67_).length(0))]: p1})
                d_107_v69_: D2
                d_107_v69_ = D2_DC6(d_106_v68_)
                d_108_v70_: _dafny.Seq
                d_108_v70_ = _dafny.SeqWithoutIsStrInference([d_107_v69_, d_107_v69_, D2_DC6(d_106_v68_)])
                index26_ = default__.safeIndex(603, (d_105_v67_).length(0))
                (d_105_v67_)[index26_] = (d_107_v69_) in (d_108_v70_)
            d_109_v71_: _dafny.Array
            def lambda10_(d_110_p0_):
                def lambda11_(d_111_i10_):
                    return d_110_p0_

                return lambda11_

            init5_ = lambda10_(p0)
            nw27_ = _dafny.Array(None, 13)
            for i0_5_ in range(nw27_.length(0)):
                nw27_[i0_5_] = init5_(i0_5_)
            d_109_v71_ = nw27_
            d_112_v72_: C0
            nw28_ = C0()
            nw28_.ctor__(d_109_v71_)
            d_112_v72_ = nw28_
            if p0:
                r0 = p1
                d_113_v73_: _dafny.Set
                d_113_v73_ = _dafny.Set({True, not(True)})
                d_114_v74_: D6
                d_114_v74_ = D6_DC18(d_113_v73_)
                rhs5_ = (0) - (p1)
                rhs6_ = (d_114_v74_).cf27
                r1 = rhs5_
                d_113_v73_ = rhs6_
                d_115_v75_: _dafny.Array
                nw29_ = _dafny.Array(int(0), 13)
                d_115_v75_ = nw29_
                d_116_v76_: D7
                d_116_v76_ = D7_DC20(d_115_v75_)
                d_117_v77_: _dafny.Array
                nw30_ = _dafny.Array(None, 18)
                nw30_[int(0)] = d_115_v75_
                nw30_[int(1)] = d_115_v75_
                nw30_[int(2)] = d_115_v75_
                nw30_[int(3)] = d_115_v75_
                nw30_[int(4)] = d_115_v75_
                nw30_[int(5)] = d_115_v75_
                nw30_[int(6)] = d_115_v75_
                nw30_[int(7)] = d_115_v75_
                nw30_[int(8)] = d_115_v75_
                nw30_[int(9)] = d_115_v75_
                nw30_[int(10)] = (d_116_v76_).cf31
                nw30_[int(11)] = (d_116_v76_).cf31
                nw30_[int(12)] = d_115_v75_
                nw30_[int(13)] = d_115_v75_
                nw30_[int(14)] = d_115_v75_
                nw30_[int(15)] = d_115_v75_
                nw30_[int(16)] = d_115_v75_
                nw30_[int(17)] = d_115_v75_
                d_117_v77_ = nw30_
                d_118_v78_: _dafny.Map
                d_118_v78_ = _dafny.Map({p2: d_117_v77_})
                d_118_v78_ = (d_118_v78_).set(p0, d_117_v77_)
                (globalState).f0 = p1
                d_119_v79_: _dafny.Seq
                d_119_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcyapl"))
                r2 = not((not(p2) if (d_119_v79_) != (d_119_v79_) else False))
            elif True:
                d_120_v80_: _dafny.Map
                d_120_v80_ = _dafny.Map({p1: p1})
                d_121_v81_: _dafny.Seq
                d_121_v81_ = _dafny.SeqWithoutIsStrInference([len(d_120_v80_), len(default__.fm0(p0, globalState)), p1])
                d_122_v82_: D3
                d_122_v82_ = D3_DC11(d_121_v81_, p2)
                d_122_v82_ = d_122_v82_
                r1 = p1
                d_123_v83_: _dafny.Array
                def lambda12_(d_124_p1_):
                    def lambda13_(d_125_i11_):
                        return (d_125_i11_) * (d_124_p1_)

                    return lambda13_

                init6_ = lambda12_(p1)
                nw31_ = _dafny.Array(None, 6)
                for i0_6_ in range(nw31_.length(0)):
                    nw31_[i0_6_] = init6_(i0_6_)
                d_123_v83_ = nw31_
                index27_ = default__.safeIndex(926, (d_123_v83_).length(0))
                (d_123_v83_)[index27_] = (0) - (p1)
                index28_ = default__.safeIndex(926, (d_123_v83_).length(0))
                (d_123_v83_)[index28_] = (((d_14_v0_)[True] if (True) in (d_14_v0_) else p1)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eskgngpo"))))
                d_126_v84_: C0
                nw32_ = C0()
                nw32_.ctor__((d_112_v72_).f15)
                d_126_v84_ = nw32_
                index29_ = default__.safeIndex(112, ((d_112_v72_).f15).length(0))
                ((d_112_v72_).f15)[index29_] = p0
                d_127_v85_: _dafny.Set
                d_127_v85_ = _dafny.Set({(d_123_v83_)[default__.safeIndex(926, (d_123_v83_).length(0))], p1})
                d_128_v86_: _dafny.Seq
                d_128_v86_ = _dafny.SeqWithoutIsStrInference([p0])
                d_129_v87_: str
                d_129_v87_ = _dafny.CodePoint('f')
                d_130_v88_: _dafny.Map
                d_130_v88_ = _dafny.Map({(d_128_v86_)[default__.safeIndex(len(d_121_v81_), len(d_128_v86_))]: d_129_v87_})
                index30_ = default__.safeIndex(112, ((d_112_v72_).f15).length(0))
                ((d_112_v72_).f15)[index30_] = ((_dafny.Set({p1, len(_dafny.SeqWithoutIsStrInference([d_130_v88_]))})) - (d_127_v85_)).ispropersubset(d_127_v85_)
            d_131_v89_: _dafny.Map
            d_131_v89_ = _dafny.Map({not(p2): (d_112_v72_).f15})
            d_131_v89_ = (d_131_v89_).set(p2, d_109_v71_)
        elif True:
            d_132_v90_: _dafny.Seq
            d_132_v90_ = _dafny.SeqWithoutIsStrInference([(p1) * (p1), p1])
            d_133_v91_: _dafny.Array
            nw33_ = _dafny.Array(False, 23)
            d_133_v91_ = nw33_
            index31_ = default__.safeIndex(921, (d_133_v91_).length(0))
            (d_133_v91_)[index31_] = p0
            d_134_v92_: _dafny.Seq
            d_134_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
            d_135_v93_: str
            d_135_v93_ = _dafny.CodePoint('p')
            index32_ = default__.safeIndex(921, (d_133_v91_).length(0))
            rhs7_ = d_132_v90_
            rhs8_ = ((p1) + (len((d_134_v92_).set(default__.safeIndex(p1, len(d_134_v92_)), d_135_v93_)))) < ((d_132_v90_)[default__.safeIndex(p1, len(d_132_v90_))])
            lhs4_ = d_133_v91_
            lhs5_ = default__.safeIndex(921, (d_133_v91_).length(0))
            d_132_v90_ = rhs7_
            lhs4_[lhs5_] = rhs8_
            d_136_v94_: C0
            nw34_ = C0()
            nw34_.ctor__(d_133_v91_)
            d_136_v94_ = nw34_
            d_137_v95_: _dafny.Seq
            d_137_v95_ = _dafny.SeqWithoutIsStrInference([d_136_v94_])
            d_138_v96_: _dafny.MultiSet
            d_138_v96_ = _dafny.MultiSet([d_136_v94_, d_136_v94_, (d_137_v95_)[default__.safeIndex(-526, len(d_137_v95_))], d_136_v94_])
            d_138_v96_ = d_138_v96_
            d_139_v97_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.MultiSet({}), 26)
            d_139_v97_ = nw35_
            d_140_v98_: _dafny.Map
            d_140_v98_ = _dafny.Map({(0) - (p1): p1})
            d_141_v99_: _dafny.MultiSet
            d_141_v99_ = _dafny.MultiSet([d_140_v98_])
            d_142_v100_: _dafny.MultiSet
            d_142_v100_ = d_141_v99_
            index33_ = default__.safeIndex(596, (d_139_v97_).length(0))
            (d_139_v97_)[index33_] = (d_142_v100_)
            index34_ = default__.safeIndex(596, (d_139_v97_).length(0))
            (d_139_v97_)[index34_] = d_141_v99_
            d_143_v101_: _dafny.Map
            d_143_v101_ = _dafny.Map({p2: p0})
            d_144_v102_: _dafny.Map
            d_144_v102_ = _dafny.Map({d_143_v101_: (d_133_v91_)[default__.safeIndex(921, (d_133_v91_).length(0))]})
            d_144_v102_ = (d_144_v102_).set(d_143_v101_, (len(d_134_v92_)) == (p1))
            r0 = p1
        d_145_v103_: _dafny.MultiSet
        d_145_v103_ = _dafny.MultiSet([p1])
        d_146_v105_: str
        d_146_v105_ = _dafny.CodePoint('a')
        d_147_v106_: _dafny.Seq
        def iife7_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.Map({d_146_v105_: p2})).keys.Elements:
                d_148_v104_: str = compr_3_
                if (d_148_v104_) in (_dafny.Map({d_146_v105_: p2})):
                    coll3_[d_148_v104_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spyhcfqae")))
            return _dafny.Map(coll3_)
        d_147_v106_ = _dafny.SeqWithoutIsStrInference([iife7_()
        ])
        d_149_v107_: _dafny.Map
        d_149_v107_ = _dafny.Map({d_146_v105_: 844})
        d_150_v108_: _dafny.Set
        d_150_v108_ = _dafny.Set({p1, (0) - (p1), (d_145_v103_).cardinality, len((d_147_v106_).set(default__.safeIndex(default__.fm4(globalState), len(d_147_v106_)), d_149_v107_)), p1})
        d_151_v109_: _dafny.Seq
        d_151_v109_ = _dafny.SeqWithoutIsStrInference([default__.fm13(503, globalState)])
        d_152_v110_: _dafny.Set
        d_152_v110_ = _dafny.Set({p0})
        d_153_v111_: _dafny.MultiSet
        d_153_v111_ = _dafny.MultiSet([d_152_v110_, _dafny.Set({True, p2}), d_152_v110_, d_152_v110_])
        d_154_v112_: D5
        d_154_v112_ = D5_DC15(d_153_v111_)
        d_155_v113_: _dafny.Map
        d_155_v113_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.Set({p1}), d_150_v108_, (d_151_v109_)[default__.safeIndex(p1, len(d_151_v109_))], d_150_v108_, d_150_v108_]): d_154_v112_})
        d_155_v113_ = (d_155_v113_).set(d_151_v109_, d_154_v112_)
        d_156_v114_: _dafny.Array
        def lambda14_(d_157_p2_):
            def lambda15_(d_158_i13_):
                return d_157_p2_

            return lambda15_

        init7_ = lambda14_(p2)
        nw36_ = _dafny.Array(None, 22)
        for i0_7_ in range(nw36_.length(0)):
            nw36_[i0_7_] = init7_(i0_7_)
        d_156_v114_ = nw36_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_156_v114_).length(0)):
            d_159_i12_: int = guard_loop_0_
            if (True) and (((0) <= (d_159_i12_)) and ((d_159_i12_) < ((d_156_v114_).length(0)))):
                (d_156_v114_)[(d_159_i12_)] = (not(p2) if p0 else (p1) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldv")))))
        d_160_i14_: int
        d_160_i14_ = 0
        with _dafny.label("0"):
            while (d_14_v0_).ispropersubset(((d_14_v0_).set(p2, default__.abs((0) - (p1)))).set(False, default__.abs(p1))):
                with _dafny.c_label("0"):
                    if (d_160_i14_) >= (100):
                        raise _dafny.Break("0")
                    d_160_i14_ = (d_160_i14_) + (1)
                    r2 = (p0 if p2 else not((p2 if p0 else p2)))
                    r2 = (False if p2 else p2)
                    index35_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    (d_156_v114_)[index35_] = not(p0)
                    index36_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    (d_156_v114_)[index36_] = False
                    d_161_v115_: D5
                    d_161_v115_ = D5_DC16((len(_dafny.Map({p1: p0}))) * (p1), p1)
                    d_162_v116_: _dafny.Seq
                    d_162_v116_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsrufvirc"))
                    index37_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    index38_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    def iife8_(_pat_let2_0):
                        def iife9_(d_163_dt__update__tmp_h1_):
                            def iife10_(_pat_let3_0):
                                def iife11_(d_164_dt__update_hcf24_h0_):
                                    return D5_DC16(d_164_dt__update_hcf24_h0_, (d_163_dt__update__tmp_h1_).cf25)
                                return iife11_(_pat_let3_0)
                            return iife10_(590)
                        return iife9_(_pat_let2_0)
                    rhs9_ = iife8_(D5_DC16(p1, p1))
                    rhs10_ = len(d_162_v116_)
                    rhs11_ = not(p2)
                    rhs12_ = (d_156_v114_)[default__.safeIndex(117, (d_156_v114_).length(0))]
                    lhs6_ = globalState
                    lhs7_ = d_156_v114_
                    lhs8_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    lhs9_ = d_156_v114_
                    lhs10_ = default__.safeIndex(117, (d_156_v114_).length(0))
                    d_161_v115_ = rhs9_
                    lhs6_.f0 = rhs10_
                    lhs7_[lhs8_] = rhs11_
                    lhs9_[lhs10_] = rhs12_
                    pass
            pass
        d_165_v117_: C0
        nw37_ = C0()
        nw37_.ctor__(d_156_v114_)
        d_165_v117_ = nw37_
        r0 = p1
        r1 = (0) - ((937) + (default__.fm4(globalState)))
        r2 = not(p0)
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_166_v0_: bool
        d_166_v0_ = True
        d_167_v1_: _dafny.Seq
        d_167_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer"))
        d_168_v2_: _dafny.Array
        nw38_ = _dafny.Array(_dafny.Set({}), 28)
        d_168_v2_ = nw38_
        d_169_globalState_: GlobalState
        nw39_ = GlobalState()
        nw39_.ctor__(-982, _dafny.SeqWithoutIsStrInference([d_166_v0_]), False, (d_167_v1_) + (d_167_v1_), 745, d_167_v1_, 292, _dafny.CodePoint('x'), 420, (d_167_v1_) + (d_167_v1_), 835, False, -89, 488, d_168_v2_)
        d_169_globalState_ = nw39_
        d_170_v3_: _dafny.Seq
        d_170_v3_ = _dafny.SeqWithoutIsStrInference([d_166_v0_, d_166_v0_])
        d_170_v3_ = (d_170_v3_) + (default__.fm0(d_166_v0_, d_169_globalState_))
        d_171_v4_: _dafny.Seq
        d_171_v4_ = _dafny.SeqWithoutIsStrInference([d_170_v3_])
        d_172_v5_: _dafny.Map
        d_172_v5_ = _dafny.Map({(True): d_166_v0_})
        d_173_v6_: _dafny.Map
        d_173_v6_ = _dafny.Map({((d_171_v4_)[default__.safeIndex(len(d_172_v5_), len(d_171_v4_))]) != (d_170_v3_): d_167_v1_})
        d_173_v6_ = (d_173_v6_).set(default__.fm1(default__.fm2(d_166_v0_, d_166_v0_, d_169_globalState_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqfstui")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_174_i0_ in range(default__.abs(820))]), d_169_globalState_), (d_167_v1_) + (d_167_v1_))
        d_175_i1_: int
        d_175_i1_ = 0
        with _dafny.label("1"):
            while (d_166_v0_ if d_166_v0_ else d_166_v0_):
                with _dafny.c_label("1"):
                    if (d_175_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_175_i1_ = (d_175_i1_) + (1)
                    d_176_v7_: int
                    d_176_v7_ = 795
                    d_177_v8_: int
                    d_178_v9_: int
                    d_179_v10_: bool
                    out0_: int
                    out1_: int
                    out2_: bool
                    out0_, out1_, out2_ = default__.m0(not(d_166_v0_), (0) - (d_176_v7_), d_166_v0_, d_169_globalState_)
                    d_177_v8_ = out0_
                    d_178_v9_ = out1_
                    d_179_v10_ = out2_
                    d_180_v11_: _dafny.Array
                    nw40_ = _dafny.Array(_dafny.Array(None, 0), 29)
                    d_180_v11_ = nw40_
                    d_181_v12_: _dafny.Array
                    nw41_ = _dafny.Array(False, 20)
                    d_181_v12_ = nw41_
                    nw42_ = _dafny.Array(None, 27)
                    nw42_[int(0)] = d_181_v12_
                    nw42_[int(1)] = d_181_v12_
                    nw42_[int(2)] = d_181_v12_
                    nw42_[int(3)] = d_181_v12_
                    nw42_[int(4)] = d_181_v12_
                    nw42_[int(5)] = d_181_v12_
                    nw42_[int(6)] = d_181_v12_
                    nw42_[int(7)] = d_181_v12_
                    nw42_[int(8)] = d_181_v12_
                    nw42_[int(9)] = d_181_v12_
                    nw42_[int(10)] = d_181_v12_
                    nw42_[int(11)] = d_181_v12_
                    nw42_[int(12)] = d_181_v12_
                    nw42_[int(13)] = d_181_v12_
                    nw42_[int(14)] = d_181_v12_
                    nw42_[int(15)] = d_181_v12_
                    nw42_[int(16)] = d_181_v12_
                    nw42_[int(17)] = d_181_v12_
                    nw42_[int(18)] = d_181_v12_
                    nw42_[int(19)] = d_181_v12_
                    nw42_[int(20)] = d_181_v12_
                    nw42_[int(21)] = d_181_v12_
                    nw42_[int(22)] = d_181_v12_
                    nw42_[int(23)] = d_181_v12_
                    nw42_[int(24)] = d_181_v12_
                    nw42_[int(25)] = d_181_v12_
                    nw42_[int(26)] = d_181_v12_
                    d_180_v11_ = nw42_
                    d_182_v13_: _dafny.Array
                    nw43_ = _dafny.Array(False, 15)
                    d_182_v13_ = nw43_
                    d_183_v14_: bool
                    d_183_v14_ = d_166_v0_
                    index39_ = default__.safeIndex(349, (d_182_v13_).length(0))
                    (d_182_v13_)[index39_] = d_183_v14_
                    index40_ = default__.safeIndex(349, (d_182_v13_).length(0))
                    (d_182_v13_)[index40_] = d_183_v14_
                    rhs13_ = d_176_v7_
                    rhs14_ = d_179_v10_
                    d_176_v7_ = rhs13_
                    d_179_v10_ = rhs14_
                    pass
            pass
        d_184_v15_: int
        d_184_v15_ = 261
        (d_169_globalState_).f0 = (d_184_v15_) * (d_184_v15_)
        d_185_v16_: bool
        d_185_v16_ = not(d_166_v0_)
        d_185_v16_ = d_185_v16_
        (d_169_globalState_).f0 = default__.safeDivisionInt((0) - ((0) - ((d_184_v15_ if d_166_v0_ else d_184_v15_))), (d_184_v15_) - (d_184_v15_))
        d_186_i2_: int
        d_186_i2_ = 0
        with _dafny.label("2"):
            while (d_166_v0_):
                with _dafny.c_label("2"):
                    if (d_186_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_186_i2_ = (d_186_i2_) + (1)
                    d_187_v17_: _dafny.Array
                    def lambda16_(d_188_v0_):
                        def lambda17_(d_189_i3_):
                            return d_188_v0_

                        return lambda17_

                    init8_ = lambda16_(d_166_v0_)
                    nw44_ = _dafny.Array(None, 28)
                    for i0_8_ in range(nw44_.length(0)):
                        nw44_[i0_8_] = init8_(i0_8_)
                    d_187_v17_ = nw44_
                    index41_ = default__.safeIndex(902, (d_187_v17_).length(0))
                    (d_187_v17_)[index41_] = d_166_v0_
                    d_190_v18_: _dafny.Array
                    nw45_ = _dafny.Array(None, 1)
                    nw45_[int(0)] = d_185_v16_
                    d_190_v18_ = nw45_
                    d_191_v19_: _dafny.Map
                    d_191_v19_ = _dafny.Map({d_190_v18_: d_184_v15_})
                    index42_ = default__.safeIndex(902, (d_187_v17_).length(0))
                    (d_187_v17_)[index42_] = (d_191_v19_) == (d_191_v19_)
                    d_192_v20_: str
                    d_192_v20_ = _dafny.CodePoint('b')
                    rhs15_ = d_184_v15_
                    rhs16_ = default__.safeModuloInt(d_184_v15_, d_184_v15_)
                    rhs17_ = not((d_187_v17_)[default__.safeIndex(902, (d_187_v17_).length(0))])
                    rhs18_ = d_167_v1_
                    rhs19_ = _dafny.CodePoint('c')
                    lhs11_ = d_169_globalState_
                    lhs12_ = d_169_globalState_
                    lhs11_.f0 = rhs15_
                    d_184_v15_ = rhs16_
                    d_166_v0_ = rhs17_
                    lhs12_.f5 = rhs18_
                    d_192_v20_ = rhs19_
                    d_193_v21_: _dafny.Set
                    d_193_v21_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_170_v3_)[default__.safeIndex((default__.fm3(d_184_v15_, d_184_v15_, d_184_v15_, d_169_globalState_)).cardinality, len(d_170_v3_))]]), d_170_v3_})
                    d_194_v22_: _dafny.Map
                    d_194_v22_ = _dafny.Map({default__.safeDivisionInt(len(d_193_v21_), 426): d_184_v15_})
                    d_194_v22_ = ((d_194_v22_) | (_dafny.Map({d_184_v15_: d_184_v15_}))) | (d_194_v22_)
                    d_195_v23_: _dafny.MultiSet
                    d_195_v23_ = _dafny.MultiSet([d_184_v15_])
                    d_196_v24_: _dafny.Map
                    d_196_v24_ = _dafny.Map({(d_195_v23_).set(d_184_v15_, default__.abs(d_184_v15_)): d_184_v15_})
                    d_170_v3_ = ((d_170_v3_).set(default__.safeIndex(len(d_196_v24_), len(d_170_v3_)), d_166_v0_)) + (d_170_v3_)
                    pass
            pass
        d_197_v25_: _dafny.Array
        nw46_ = _dafny.Array(False, 28)
        d_197_v25_ = nw46_
        d_198_v26_: _dafny.Set
        d_198_v26_ = _dafny.Set({d_197_v25_, d_197_v25_})
        d_199_v27_: _dafny.Map
        d_199_v27_ = _dafny.Map({default__.fm4(d_169_globalState_): d_198_v26_})
        d_200_v28_: _dafny.Array
        nw47_ = _dafny.Array(None, 25)
        nw47_[int(0)] = d_198_v26_
        nw47_[int(1)] = _dafny.Set({d_197_v25_, d_197_v25_})
        nw47_[int(2)] = d_198_v26_
        nw47_[int(3)] = _dafny.Set({d_197_v25_})
        nw47_[int(4)] = (d_198_v26_) | (d_198_v26_)
        nw47_[int(5)] = _dafny.Set({d_197_v25_})
        nw47_[int(6)] = d_198_v26_
        nw47_[int(7)] = d_198_v26_
        nw47_[int(8)] = _dafny.Set({d_197_v25_, d_197_v25_, d_197_v25_})
        nw47_[int(9)] = d_198_v26_
        nw47_[int(10)] = (d_198_v26_).intersection(d_198_v26_)
        nw47_[int(11)] = d_198_v26_
        nw47_[int(12)] = d_198_v26_
        nw47_[int(13)] = d_198_v26_
        nw47_[int(14)] = _dafny.Set({d_197_v25_})
        nw47_[int(15)] = d_198_v26_
        nw47_[int(16)] = d_198_v26_
        nw47_[int(17)] = d_198_v26_
        nw47_[int(18)] = (_dafny.Set({d_197_v25_, d_197_v25_})) | (d_198_v26_)
        nw47_[int(19)] = d_198_v26_
        nw47_[int(20)] = d_198_v26_
        nw47_[int(21)] = (((d_199_v27_)[d_184_v15_] if (d_184_v15_) in (d_199_v27_) else d_198_v26_)) | (d_198_v26_)
        nw47_[int(22)] = d_198_v26_
        nw47_[int(23)] = (d_198_v26_ if (d_170_v3_)[default__.safeIndex(d_184_v15_, len(d_170_v3_))] else d_198_v26_)
        nw47_[int(24)] = d_198_v26_
        d_200_v28_ = nw47_
        index43_ = default__.safeIndex(693, (d_200_v28_).length(0))
        (d_200_v28_)[index43_] = d_198_v26_
        index44_ = default__.safeIndex(693, (d_200_v28_).length(0))
        (d_200_v28_)[index44_] = d_198_v26_
        source2_ = d_185_v16_
        d_201___mcc_h0_ = source2_
        d_202_cf0_ = d_201___mcc_h0_
        d_173_v6_ = (d_173_v6_).set((not(d_166_v0_)) == (d_166_v0_), (d_167_v1_) + (d_167_v1_))
        d_203_v29_: _dafny.Set
        d_203_v29_ = _dafny.Set({d_202_cf0_})
        d_204_v30_: _dafny.Array
        nw48_ = _dafny.Array(int(0), 3)
        d_204_v30_ = nw48_
        index45_ = default__.safeIndex(733, (d_204_v30_).length(0))
        (d_204_v30_)[index45_] = d_184_v15_
        d_205_v31_: _dafny.Set
        d_205_v31_ = _dafny.Set({(0) - (d_184_v15_)})
        d_206_v32_: _dafny.MultiSet
        d_206_v32_ = _dafny.MultiSet([d_167_v1_])
        index46_ = default__.safeIndex(733, (d_204_v30_).length(0))
        rhs20_ = default__.fm5(d_184_v15_, (len(d_205_v31_)) < (d_184_v15_), (d_167_v1_) + (d_167_v1_), (d_206_v32_) | (d_206_v32_), d_169_globalState_)
        rhs21_ = d_166_v0_
        rhs22_ = (d_184_v15_) + (default__.fm4(d_169_globalState_))
        lhs13_ = d_204_v30_
        lhs14_ = default__.safeIndex(733, (d_204_v30_).length(0))
        d_203_v29_ = rhs20_
        d_202_cf0_ = rhs21_
        lhs13_[lhs14_] = rhs22_
        d_207_v33_: _dafny.Map
        d_207_v33_ = _dafny.Map({(d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))]: d_166_v0_})
        d_208_v34_: str
        d_208_v34_ = _dafny.CodePoint('f')
        index47_ = default__.safeIndex(733, (d_204_v30_).length(0))
        rhs23_ = 373
        rhs24_ = ((_dafny.Set({d_202_cf0_})).intersection(d_203_v29_)).ispropersubset(d_203_v29_)
        rhs25_ = default__.fm1(d_207_v33_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtng")), d_167_v1_, d_169_globalState_)
        rhs26_ = (0) - (((d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))]) + ((len((d_167_v1_).set(default__.safeIndex((d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))], len(d_167_v1_)), d_208_v34_)) if d_202_cf0_ else (d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))])))
        rhs27_ = (d_184_v15_) >= (len(d_167_v1_))
        lhs15_ = d_169_globalState_
        lhs16_ = d_204_v30_
        lhs17_ = default__.safeIndex(733, (d_204_v30_).length(0))
        lhs15_.f0 = rhs23_
        d_202_cf0_ = rhs24_
        d_166_v0_ = rhs25_
        lhs16_[lhs17_] = rhs26_
        d_202_cf0_ = rhs27_
        if not(False):
            d_209_v35_: C0
            nw49_ = C0()
            nw49_.ctor__(d_197_v25_)
            d_209_v35_ = nw49_
            d_166_v0_ = (default__.fm4(d_169_globalState_)) >= ((len(d_170_v3_)) * ((0) - ((d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))])))
            d_210_v36_: int
            d_211_v37_: int
            d_212_v38_: bool
            out3_: int
            out4_: int
            out5_: bool
            out3_, out4_, out5_ = default__.m0(d_202_cf0_, (d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))], d_202_cf0_, d_169_globalState_)
            d_210_v36_ = out3_
            d_211_v37_ = out4_
            d_212_v38_ = out5_
            d_172_v5_ = (d_172_v5_) | ((d_172_v5_) | (_dafny.Map({not(d_202_cf0_): d_166_v0_})))
            d_213_v39_: int
            d_214_v40_: int
            d_215_v41_: bool
            out6_: int
            out7_: int
            out8_: bool
            out6_, out7_, out8_ = default__.m0((default__.fm4(d_169_globalState_)) > (d_210_v36_), (0) - ((d_211_v37_) + ((D1_DC1(len(_dafny.SeqWithoutIsStrInference([d_208_v34_ for d_216_i4_ in range(default__.abs(783))])))).cf1)), d_202_cf0_, d_169_globalState_)
            d_213_v39_ = out6_
            d_214_v40_ = out7_
            d_215_v41_ = out8_
        elif True:
            d_217_v42_: _dafny.Map
            d_217_v42_ = _dafny.Map({len(d_207_v33_): (d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))]})
            d_218_v43_: _dafny.Map
            d_218_v43_ = _dafny.Map({d_202_cf0_: len(d_217_v42_)})
            d_219_v44_: _dafny.Map
            d_219_v44_ = _dafny.Map({d_184_v15_: d_218_v43_})
            d_220_v45_: _dafny.Seq
            d_220_v45_ = _dafny.SeqWithoutIsStrInference([(d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))], d_184_v15_, (d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))], len(((d_219_v44_)[d_184_v15_] if (d_184_v15_) in (d_219_v44_) else (d_218_v43_).set(False, d_184_v15_))), len(d_205_v31_)])
            d_221_v46_: int
            d_222_v47_: int
            d_223_v48_: bool
            out9_: int
            out10_: int
            out11_: bool
            out9_, out10_, out11_ = default__.m0(d_166_v0_, (len(d_220_v45_) if d_202_cf0_ else (d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))]), not((len(d_172_v5_)) != ((d_204_v30_)[default__.safeIndex(733, (d_204_v30_).length(0))])), d_169_globalState_)
            d_221_v46_ = out9_
            d_222_v47_ = out10_
            d_223_v48_ = out11_
            d_173_v6_ = (d_173_v6_).set(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhphth")))
            d_224_v49_: int
            d_225_v50_: int
            d_226_v51_: bool
            out12_: int
            out13_: int
            out14_: bool
            out12_, out13_, out14_ = default__.m0((True) and (d_202_cf0_), 825, default__.fm1(default__.fm2(d_223_v48_, d_166_v0_, d_169_globalState_), d_167_v1_, d_167_v1_, d_169_globalState_), d_169_globalState_)
            d_224_v49_ = out12_
            d_225_v50_ = out13_
            d_226_v51_ = out14_
            (d_169_globalState_).f5 = d_167_v1_
            (d_169_globalState_).f5 = (d_167_v1_) + ((d_167_v1_) + (_dafny.SeqWithoutIsStrInference([d_208_v34_ for d_227_i5_ in range(default__.abs(72))])))
        d_228_v52_: C0
        nw50_ = C0()
        nw50_.ctor__(d_197_v25_)
        d_228_v52_ = nw50_
        d_229_v53_: _dafny.Set
        d_229_v53_ = _dafny.Set({d_166_v0_})
        d_230_v54_: _dafny.MultiSet
        d_230_v54_ = _dafny.MultiSet([d_167_v1_, d_167_v1_])
        d_231_v55_: D1
        d_231_v55_ = D1_DC4((default__.fm5(867, not(False), d_167_v1_, d_230_v54_, d_169_globalState_)).issubset(d_229_v53_), d_166_v0_, d_197_v25_, d_166_v0_)
        source3_ = d_231_v55_
        if source3_.is_DC2:
            d_232___mcc_h1_ = source3_.cf2
            d_233___mcc_h2_ = source3_.cf3
            d_234___mcc_h3_ = source3_.cf4
            d_235_cf4_ = d_234___mcc_h3_
            d_236_cf3_ = d_233___mcc_h2_
            d_237_cf2_ = d_232___mcc_h1_
            d_238_v56_: D1
            d_238_v56_ = D1_DC2(d_237_cf2_, d_166_v0_, d_166_v0_)
            d_239_v57_: _dafny.Map
            d_239_v57_ = _dafny.Map({d_238_v56_: d_236_cf3_})
            d_240_v58_: _dafny.MultiSet
            d_240_v58_ = _dafny.MultiSet([d_239_v57_, _dafny.Map({D1_DC2(d_237_cf2_, False, d_185_v16_): d_236_cf3_})])
            d_241_v59_: _dafny.Seq
            d_241_v59_ = _dafny.SeqWithoutIsStrInference([357, ((d_240_v58_)[default__.fm6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhshk")), len(d_170_v3_), d_169_globalState_)] if (default__.fm6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhshk")), len(d_170_v3_), d_169_globalState_)) in (d_240_v58_) else d_237_cf2_), d_237_cf2_, d_237_cf2_, d_184_v15_])
            d_242_v60_: str
            d_242_v60_ = _dafny.CodePoint('d')
            d_243_v61_: _dafny.Array
            nw51_ = _dafny.Array(None, 28)
            d_243_v61_ = nw51_
            d_244_v62_: _dafny.Map
            d_244_v62_ = _dafny.Map({d_184_v15_: d_243_v61_})
            d_245_v63_: _dafny.Array
            nw52_ = _dafny.Array(None, 19)
            nw52_[int(0)] = d_241_v59_
            nw52_[int(1)] = d_241_v59_
            nw52_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_184_v15_])) + (d_241_v59_)
            nw52_[int(3)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_184_v15_: True})) for d_246_i6_ in range(default__.abs(285))])) + (_dafny.SeqWithoutIsStrInference([d_184_v15_ for d_247_i7_ in range(default__.abs(-511))]))
            nw52_[int(4)] = _dafny.SeqWithoutIsStrInference([d_184_v15_ for d_248_i8_ in range(default__.abs(363))])
            nw52_[int(5)] = (d_241_v59_) + (d_241_v59_)
            nw52_[int(6)] = d_241_v59_
            nw52_[int(7)] = d_241_v59_
            nw52_[int(8)] = (d_241_v59_) + (d_241_v59_)
            nw52_[int(9)] = _dafny.SeqWithoutIsStrInference([(0) - (d_184_v15_) for d_249_i9_ in range(default__.abs(637))])
            nw52_[int(10)] = _dafny.SeqWithoutIsStrInference([d_237_cf2_ for d_250_i10_ in range(default__.abs(-277))])
            nw52_[int(11)] = (d_241_v59_) + (d_241_v59_)
            nw52_[int(12)] = (d_241_v59_) + (d_241_v59_)
            nw52_[int(13)] = d_241_v59_
            nw52_[int(14)] = ((d_241_v59_ if True else default__.fm7(d_236_cf3_, d_242_v60_, d_237_cf2_, d_169_globalState_))).set(default__.safeIndex(len((d_244_v62_).set(d_184_v15_, d_243_v61_)), len((d_241_v59_ if True else default__.fm7(d_236_cf3_, d_242_v60_, d_237_cf2_, d_169_globalState_)))), d_237_cf2_)
            nw52_[int(15)] = d_241_v59_
            nw52_[int(16)] = d_241_v59_
            nw52_[int(17)] = d_241_v59_
            nw52_[int(18)] = _dafny.SeqWithoutIsStrInference([len(d_170_v3_) for d_251_i11_ in range(default__.abs(674))])
            d_245_v63_ = nw52_
            index48_ = default__.safeIndex(598, (d_245_v63_).length(0))
            (d_245_v63_)[index48_] = d_241_v59_
            index49_ = default__.safeIndex(598, (d_245_v63_).length(0))
            (d_245_v63_)[index49_] = ((d_241_v59_) + (d_241_v59_)).set(default__.safeIndex((d_237_cf2_) * (25), len((d_241_v59_) + (d_241_v59_))), len(d_167_v1_))
            rhs28_ = not(not((d_236_cf3_)))
            rhs29_ = d_228_v52_
            rhs30_ = d_166_v0_
            rhs31_ = d_237_cf2_
            rhs32_ = d_184_v15_
            d_236_cf3_ = rhs28_
            d_228_v52_ = rhs29_
            d_166_v0_ = rhs30_
            d_184_v15_ = rhs31_
            d_184_v15_ = rhs32_
            d_252_v64_: int
            d_253_v65_: int
            d_254_v66_: bool
            out15_: int
            out16_: int
            out17_: bool
            out15_, out16_, out17_ = default__.m0((d_170_v3_)[default__.safeIndex((0) - (d_237_cf2_), len(d_170_v3_))], d_184_v15_, d_166_v0_, d_169_globalState_)
            d_252_v64_ = out15_
            d_253_v65_ = out16_
            d_254_v66_ = out17_
            d_255_v67_: _dafny.Map
            d_255_v67_ = _dafny.Map({d_254_v66_: d_170_v3_})
            d_184_v15_ = (len((d_255_v67_) | (d_255_v67_))) * (214)
        elif source3_.is_DC3:
            (d_169_globalState_).f0 = (0) - (d_184_v15_)
            index50_ = default__.safeIndex(441, ((d_228_v52_).f15).length(0))
            ((d_228_v52_).f15)[index50_] = d_166_v0_
            index51_ = default__.safeIndex(441, ((d_228_v52_).f15).length(0))
            ((d_228_v52_).f15)[index51_] = d_166_v0_
            index52_ = default__.safeIndex(441, ((d_228_v52_).f15).length(0))
            ((d_228_v52_).f15)[index52_] = d_166_v0_
            d_256_v69_: _dafny.Array
            def lambda18_(d_257_v15_, d_258_v52_):
                def lambda19_(d_259_i12_):
                    def iife12_():
                        coll4_ = _dafny.Map()
                        compr_4_: int
                        for compr_4_ in (_dafny.SeqWithoutIsStrInference([d_257_v15_])).Elements:
                            d_260_v68_: int = compr_4_
                            if (d_260_v68_) in (_dafny.SeqWithoutIsStrInference([d_257_v15_])):
                                coll4_[default__.safeDivisionInt(d_260_v68_, d_257_v15_)] = ((d_258_v52_).f15)[default__.safeIndex(441, ((d_258_v52_).f15).length(0))]
                        return _dafny.Map(coll4_)
                    return default__.safeModuloInt(d_259_i12_, len(iife12_()
                    ))

                return lambda19_

            init9_ = lambda18_(d_184_v15_, d_228_v52_)
            nw53_ = _dafny.Array(None, 11)
            for i0_9_ in range(nw53_.length(0)):
                nw53_[i0_9_] = init9_(i0_9_)
            d_256_v69_ = nw53_
            d_256_v69_ = d_256_v69_
        elif source3_.is_DC4:
            d_261___mcc_h4_ = source3_.cf5
            d_262___mcc_h5_ = source3_.cf6
            d_263___mcc_h6_ = source3_.cf7
            d_264___mcc_h7_ = source3_.cf8
            d_265_cf8_ = d_264___mcc_h7_
            d_266_cf7_ = d_263___mcc_h6_
            d_267_cf6_ = d_262___mcc_h5_
            d_268_cf5_ = d_261___mcc_h4_
            d_269_v70_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
            d_269_v70_ = nw54_
            index53_ = default__.safeIndex(675, (d_269_v70_).length(0))
            (d_269_v70_)[index53_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cccpixh"))
            index54_ = default__.safeIndex(675, (d_269_v70_).length(0))
            (d_269_v70_)[index54_] = d_167_v1_
            if d_265_cf8_:
                (d_169_globalState_).f0 = d_184_v15_
                d_270_v71_: _dafny.Array
                nw55_ = _dafny.Array(int(0), 22)
                d_270_v71_ = nw55_
                index55_ = default__.safeIndex(181, (d_270_v71_).length(0))
                (d_270_v71_)[index55_] = default__.fm4(d_169_globalState_)
                index56_ = default__.safeIndex(181, (d_270_v71_).length(0))
                (d_270_v71_)[index56_] = default__.fm4(d_169_globalState_)
                d_271_v72_: _dafny.MultiSet
                d_271_v72_ = _dafny.MultiSet([True, d_268_cf5_])
                rhs33_ = not(not (((d_270_v71_)[default__.safeIndex(181, (d_270_v71_).length(0))]) == (len(_dafny.SeqWithoutIsStrInference([d_184_v15_ for d_272_i13_ in range(default__.abs(550))])))) or ((d_231_v55_).cf5))
                rhs34_ = (d_271_v72_).ispropersubset(d_271_v72_)
                rhs35_ = (d_269_v70_)[default__.safeIndex(675, (d_269_v70_).length(0))]
                d_166_v0_ = rhs33_
                d_267_cf6_ = rhs34_
                d_167_v1_ = rhs35_
                d_273_v73_: _dafny.Map
                d_273_v73_ = _dafny.Map({(d_270_v71_)[default__.safeIndex(181, (d_270_v71_).length(0))]: d_267_cf6_})
                d_274_v74_: str
                d_274_v74_ = _dafny.CodePoint('b')
                d_265_cf8_ = default__.fm1(d_273_v73_, d_167_v1_, default__.fm8(d_265_cf8_, d_274_v74_, default__.fm9(d_268_cf5_, d_169_globalState_), d_169_globalState_), d_169_globalState_)
                d_275_v75_: _dafny.Array
                nw56_ = _dafny.Array(None, 15)
                d_275_v75_ = nw56_
                d_276_v76_: _dafny.Seq
                d_276_v76_ = _dafny.SeqWithoutIsStrInference([d_275_v75_, d_275_v75_, d_275_v75_, d_275_v75_, d_275_v75_])
                d_276_v76_ = d_276_v76_
            elif True:
                (d_169_globalState_).f0 = d_184_v15_
                (d_169_globalState_).f0 = (d_184_v15_) + (d_184_v15_)
                d_277_v77_: int
                d_278_v78_: int
                d_279_v79_: bool
                out18_: int
                out19_: int
                out20_: bool
                out18_, out19_, out20_ = default__.m0((d_184_v15_) != (d_184_v15_), (0) - (d_184_v15_), d_265_cf8_, d_169_globalState_)
                d_277_v77_ = out18_
                d_278_v78_ = out19_
                d_279_v79_ = out20_
                d_280_v80_: C0
                nw57_ = C0()
                nw57_.ctor__((d_228_v52_).f15)
                d_280_v80_ = nw57_
                index57_ = default__.safeIndex(295, ((d_228_v52_).f15).length(0))
                ((d_228_v52_).f15)[index57_] = d_166_v0_
                index58_ = default__.safeIndex(295, ((d_228_v52_).f15).length(0))
                ((d_228_v52_).f15)[index58_] = d_268_cf5_
            d_281_v81_: _dafny.Map
            d_281_v81_ = _dafny.Map({d_184_v15_: d_265_cf8_})
            d_281_v81_ = (d_281_v81_).set(d_184_v15_, d_166_v0_)
            (d_169_globalState_).f5 = d_167_v1_
        elif source3_.is_DC1:
            d_282___mcc_h8_ = source3_.cf1
            d_283_cf1_ = d_282___mcc_h8_
            (d_169_globalState_).f0 = default__.safeDivisionInt(d_283_cf1_, d_283_cf1_)
            d_284_v82_: _dafny.Map
            d_284_v82_ = _dafny.Map({d_184_v15_: d_166_v0_})
            d_285_v83_: _dafny.Map
            d_285_v83_ = _dafny.Map({d_184_v15_: len(d_284_v82_)})
            d_286_v84_: D1
            d_286_v84_ = D1_DC1(d_283_cf1_)
            (d_169_globalState_).f0 = (d_283_cf1_ if True else len((d_285_v83_).set((0) - ((d_286_v84_).cf1), d_283_cf1_)))
            d_166_v0_ = d_166_v0_
            (d_169_globalState_).f0 = d_184_v15_
        elif True:
            d_287___mcc_h9_ = source3_.cf9
            d_288_cf9_ = d_287___mcc_h9_
            d_289_v85_: _dafny.Map
            d_289_v85_ = _dafny.Map({d_184_v15_: d_184_v15_})
            (d_169_globalState_).f0 = (0) - ((default__.safeModuloInt(d_184_v15_, d_184_v15_)) - (((d_289_v85_)[d_184_v15_] if (d_184_v15_) in (d_289_v85_) else d_184_v15_)))
            d_290_v86_: C0
            nw58_ = C0()
            nw58_.ctor__(d_197_v25_)
            d_290_v86_ = nw58_
            d_291_v87_: _dafny.Map
            d_291_v87_ = _dafny.Map({((d_172_v5_)[d_166_v0_] if (d_166_v0_) in (d_172_v5_) else d_166_v0_): d_184_v15_})
            d_292_v88_: D2
            d_292_v88_ = D2_DC6(d_291_v87_)
            d_291_v87_ = (((d_292_v88_).cf10) | (_dafny.Map({d_166_v0_: d_184_v15_}))) | (d_291_v87_)
            d_166_v0_ = d_166_v0_
        d_293_v89_: _dafny.Map
        d_293_v89_ = _dafny.Map({d_184_v15_: False})
        d_294_v90_: _dafny.Map
        d_294_v90_ = _dafny.Map({default__.fm1(d_293_v89_, d_167_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvidmhm")), d_169_globalState_): default__.fm7(not(d_166_v0_), _dafny.CodePoint('p'), d_184_v15_, d_169_globalState_)})
        d_295_v91_: _dafny.Seq
        d_295_v91_ = _dafny.SeqWithoutIsStrInference([len(d_229_v53_), default__.fm4(d_169_globalState_), d_184_v15_])
        d_294_v90_ = (d_294_v90_).set(d_166_v0_, d_295_v91_)
        d_296_v92_: _dafny.Map
        d_296_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xndmxbthy")): d_167_v1_})
        d_297_v93_: _dafny.Map
        d_297_v93_ = _dafny.Map({d_296_v92_: d_166_v0_})
        def iife13_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (d_230_v54_).Elements:
                d_298_v94_: _dafny.Seq = compr_5_
                if (d_298_v94_) in (d_230_v54_):
                    coll5_[d_298_v94_] = d_167_v1_
            return _dafny.Map(coll5_)
        d_297_v93_ = (d_297_v93_).set(iife13_()
        , (d_166_v0_ if d_166_v0_ else d_166_v0_))
        d_299_v95_: str
        d_299_v95_ = _dafny.CodePoint('p')
        d_300_v96_: _dafny.MultiSet
        d_300_v96_ = _dafny.MultiSet([d_166_v0_, d_166_v0_, d_166_v0_])
        d_301_v97_: _dafny.Seq
        d_301_v97_ = _dafny.SeqWithoutIsStrInference([d_300_v96_])
        d_302_v98_: _dafny.Map
        d_302_v98_ = _dafny.Map({default__.fm10(d_299_v95_, d_166_v0_, len(d_301_v97_), d_169_globalState_): (d_228_v52_).f15})
        d_303_v99_: _dafny.Map
        d_303_v99_ = _dafny.Map({d_166_v0_: d_184_v15_})
        d_304_v100_: D2
        d_304_v100_ = D2_DC6(d_303_v99_)
        d_305_v101_: _dafny.Seq
        d_305_v101_ = _dafny.SeqWithoutIsStrInference([d_304_v100_, d_304_v100_])
        d_302_v98_ = (d_302_v98_).set(d_305_v101_, (d_228_v52_).f15)
        d_306_v102_: int
        d_307_v103_: int
        d_308_v104_: bool
        out21_: int
        out22_: int
        out23_: bool
        out21_, out22_, out23_ = default__.m0(d_166_v0_, len(d_170_v3_), (d_166_v0_ if d_166_v0_ else False), d_169_globalState_)
        d_306_v102_ = out21_
        d_307_v103_ = out22_
        d_308_v104_ = out23_
        d_308_v104_ = not(d_166_v0_)
        _dafny.print(_dafny.string_of(d_166_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_167_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_169_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_169_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_169_globalState_).f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_169_globalState_.f5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_169_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v3_) == (_dafny.SeqWithoutIsStrInference([True, False, True, True, False, False, True, False, True, True, True, True, False, False, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v4_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, True, True, False, False, True, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v5_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v6_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqerfyhqxeqer")), True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqerfyhqxeqer"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_175_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_184_v15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_186_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v25_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_198_v26_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_199_v27_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[0])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[1])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[2])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[3])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[4])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[5])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[6])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[7])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[8])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[9])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[10])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[11])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[12])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[13])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[14])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[15])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[16])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[17])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[18])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[19])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[20])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[21])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[22])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[23])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_200_v28_)[24])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_228_v52_).f15)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v53_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v54_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v55_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v55_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_231_v55_).cf7)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v55_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v89_) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v90_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([589, -1, -720, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, 44, -2, 7]), False: _dafny.SeqWithoutIsStrInference([1, 4, 0])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v91_) == (_dafny.SeqWithoutIsStrInference([1, 4, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v92_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xndmxbthy")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_297_v93_) == (_dafny.Map({_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xndmxbthy")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer"))}): False, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyhqxeqer"))}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_299_v95_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v96_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v97_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_302_v98_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_303_v99_) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_304_v100_).cf10) == (_dafny.Map({False: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v101_) == (_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({False: 0})), D2_DC6(_dafny.Map({False: 0}))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_306_v102_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_307_v103_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_308_v104_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)

class D3_DC10(D3, NamedTuple('DC10', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(None, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({self.cf20.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC19(_dafny.CodePoint('D'), False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)

class D6_DC19(D6, NamedTuple('DC19', [('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC21(D7, NamedTuple('DC21', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)

class D8_DC23(D8, NamedTuple('DC23', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC25(D9, NamedTuple('DC25', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f14: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Seq = _dafny.Seq({})
        self._f2: bool = False
        self._f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f6: int = int(0)
        self._f7: str = _dafny.CodePoint('D')
        self._f8: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: int = int(0)
        self._f11: bool = False
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self).f14 = f14

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
    def f4(self):
        return self._f4
    @property
    def f6(self):
        return self._f6
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
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C0:
    def  __init__(self):
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f15):
        (self)._f15 = f15

    @property
    def f15(self):
        return self._f15
