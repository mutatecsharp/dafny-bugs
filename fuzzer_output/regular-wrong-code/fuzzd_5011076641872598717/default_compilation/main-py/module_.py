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
        return True

    @staticmethod
    def fm1(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(64, 565):
                d_0_v0_: int = compr_0_
                if ((64) <= (d_0_v0_)) and ((d_0_v0_) < (565)):
                    coll0_[(d_0_v0_) - (len(_dafny.Set({_dafny.CodePoint('u')})))] = True
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(556, 176):
                d_1_v1_: int = compr_1_
                if ((556) <= (d_1_v1_)) and ((d_1_v1_) < (176)):
                    coll1_[default__.safeDivisionInt(d_1_v1_, len(_dafny.SeqWithoutIsStrInference([992])))] = True
            return _dafny.Map(coll1_)
        return (((_dafny.MultiSet([699, 128])).intersection(_dafny.MultiSet([(_dafny.MultiSet([_dafny.CodePoint('y')])).cardinality, -261]))) - ((_dafny.MultiSet([len(_dafny.Map({not(True): len(iife0_()
        )}))])) - (_dafny.MultiSet([len(iife1_()
        ), len(_dafny.Map({402: 682}))])))).cardinality

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "be"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfixa")))

    @staticmethod
    def fm5(p0, p1, globalState):
        source0_ = D5_DC17(400, False, not(False), not(False), (_dafny.MultiSet([False, False])).cardinality)
        if source0_.is_DC17:
            d_2___mcc_h0_ = source0_.cf27
            d_3___mcc_h1_ = source0_.cf28
            d_4___mcc_h2_ = source0_.cf29
            d_5___mcc_h3_ = source0_.cf30
            d_6___mcc_h4_ = source0_.cf31
            d_7_cf31_ = d_6___mcc_h4_
            d_8_cf30_ = d_5___mcc_h3_
            d_9_cf29_ = d_4___mcc_h2_
            d_10_cf28_ = d_3___mcc_h1_
            d_11_cf27_ = d_2___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_12_i0_ in range(default__.abs(662))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clhe")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lh"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilyfkpfb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tygu"))]))
        elif source0_.is_DC18:
            d_13___mcc_h5_ = source0_.cf32
            d_14___mcc_h6_ = source0_.cf33
            d_15_cf33_ = d_14___mcc_h6_
            d_16_cf32_ = d_13___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfng"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jovnaak")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gepycipjg"))]))
        elif True:
            d_17___mcc_h7_ = source0_.cf26
            d_18_cf26_ = d_17___mcc_h7_
            return (_dafny.SeqWithoutIsStrInference([d_18_cf26_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_19_i1_ in range(default__.abs(-90))])])) + (_dafny.SeqWithoutIsStrInference([d_18_cf26_]))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({True}))) for d_20_i0_ in range(default__.abs(806))])), 625])).Elements:
                d_21_v0_: int = compr_2_
                if (d_21_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({True}))) for d_20_i0_ in range(default__.abs(806))])), 625])):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeDivisionInt(d_21_v0_, 237)]))
            return _dafny.Set(coll2_)
        if (_dafny.MultiSet([_dafny.MultiSet([2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdpoord")))])])).issubset(_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(iife2_()
        ), 547]))])):
            return _dafny.Set({_dafny.CodePoint('t'), _dafny.CodePoint('i')})
        elif True:
            def iife3_():
                def iife5_():
                    coll5_ = _dafny.Set()
                    compr_5_: str
                    for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('e')])).Elements:
                        d_24_v1_: str = compr_5_
                        if (d_24_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('e')])):
                            coll5_ = coll5_.union(_dafny.Set([d_24_v1_]))
                    return _dafny.Set(coll5_)
                coll3_ = _dafny.Set()
                def iife4_():
                    coll4_ = _dafny.Set()
                    compr_4_: str
                    for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('e')])).Elements:
                        d_22_v1_: str = compr_4_
                        if (d_22_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('e')])):
                            coll4_ = coll4_.union(_dafny.Set([d_22_v1_]))
                    return _dafny.Set(coll4_)
                compr_3_: str
                for compr_3_ in (iife4_()
                ).Elements:
                    d_23_v2_: str = compr_3_
                    if (d_23_v2_) in (iife5_()
                    ):
                        coll3_ = coll3_.union(_dafny.Set([d_23_v2_]))
                return _dafny.Set(coll3_)
            return (iife3_()
            ) - (_dafny.Set({_dafny.CodePoint('q')}))

    @staticmethod
    def fm9(p0, globalState):
        return D0_DC3((_dafny.MultiSet([True, True])).issubset(_dafny.MultiSet([False])))

    @staticmethod
    def fm10(globalState):
        return _dafny.SeqWithoutIsStrInference([(-108) + (len(_dafny.Set({False}))), (len(_dafny.SeqWithoutIsStrInference([False, not(not(False)), False]))) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "domjx")))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgwkdeg")))]))])

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([False, False])) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm12(p0, globalState):
        return (_dafny.MultiSet([not(False)])).intersection(_dafny.MultiSet([True]))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        if True:
            if True:
                def iife6_():
                    coll6_ = _dafny.Set()
                    compr_6_: str
                    for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('t')])).Elements:
                        d_25_v0_: str = compr_6_
                        if (d_25_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('t')])):
                            coll6_ = coll6_.union(_dafny.Set([d_25_v0_]))
                    return _dafny.Set(coll6_)
                return D0_DC0(len(iife6_()
))
            elif True:
                return D0_DC0(len(_dafny.Map({589: False})))
        elif not(True):
            return D0_DC0(len(_dafny.SeqWithoutIsStrInference([False, True, True])))
        elif True:
            return D0_DC0((_dafny.MultiSet([-76, -81])).cardinality)

    @staticmethod
    def fm14(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D0_DC2(False) for d_26_i0_ in range(default__.abs(98))])) + (_dafny.SeqWithoutIsStrInference([D0_DC2(True), D0_DC2(True), D0_DC2(False), D0_DC2(False)]))) + (_dafny.SeqWithoutIsStrInference([D0_DC2(True)]))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return (_dafny.Map({len(_dafny.Set({(_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_27_i0_ in range(default__.abs(392))]))])).cardinality])).cardinality})): _dafny.SeqWithoutIsStrInference([D5_DC18(20, 354), D5_DC18((0) - (len(_dafny.Map({True: False}))), 87)])})) | (_dafny.Map({748: _dafny.SeqWithoutIsStrInference([D5_DC18(73, len(_dafny.SeqWithoutIsStrInference([True]))), D5_DC18(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsqcv"))), len(_dafny.Map({False: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).cardinality}))), D5_DC18(len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_28_i1_ in range(default__.abs(14))]))})), len(_dafny.Map({-780: True})))])}))

    @staticmethod
    def fm16(globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_29_i0_ in range(default__.abs(354))])): (_dafny.SeqWithoutIsStrInference([339, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False]) for d_30_i1_ in range(default__.abs(-818))]))])) != (_dafny.SeqWithoutIsStrInference([740]))})

    @staticmethod
    def fm17(globalState):
        if (len(_dafny.Map({True: _dafny.CodePoint('l')}))) > (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([141]))]))):
            return D2_DC8(895, 72)
        elif True:
            return D2_DC8((0) - (len(_dafny.Map({len(_dafny.Set({False})): _dafny.CodePoint('b')}))), -93)

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.Set({False, True, True})) - (_dafny.Set({False, not(not(True))}))) - (_dafny.Set({True}))

    @staticmethod
    def fm19(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(200, 666):
                d_31_v0_: int = compr_7_
                if ((200) <= (d_31_v0_)) and ((d_31_v0_) < (666)):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeDivisionInt(d_31_v0_, 261)]))
            return _dafny.Set(coll7_)
        return ((iife7_()
        ) | (_dafny.Set({884}))) | (_dafny.Set({(_dafny.MultiSet([True, False, not(not(False)), True, False])).cardinality}))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return D6_DC20(_dafny.SeqWithoutIsStrInference([True]), not (False) or (not(not(True))))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return D8_DC25((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvyxx"))), len(_dafny.Map({not(True): _dafny.MultiSet([False])}))]))})) | (_dafny.Map({265: 982})))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return (((_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kromkpsyu"))))}))) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le")))}))) | (_dafny.Map({False: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))}))}))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        if (_dafny.MultiSet([not(False), True])).ispropersubset(_dafny.MultiSet([False, False])):
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True])])) + (_dafny.SeqWithoutIsStrInference([(D1_DC4(_dafny.MultiSet([False]))).cf4, _dafny.MultiSet([not(False), not(True)]), _dafny.MultiSet([False]), _dafny.MultiSet([False, False])]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, not(True)])])

    @staticmethod
    def m2(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: bool = False
        (globalState).f3 = default__.safeDivisionInt((0) - (default__.safeDivisionInt(p2, 102)), (0) - (p2))
        d_32_v0_: bool
        d_32_v0_ = False
        if d_32_v0_:
            d_33_v1_: _dafny.MultiSet
            d_33_v1_ = _dafny.MultiSet([416, p2])
            d_34_v2_: _dafny.Map
            d_34_v2_ = _dafny.Map({(d_33_v1_).cardinality: True})
            d_35_v3_: _dafny.Seq
            d_35_v3_ = _dafny.SeqWithoutIsStrInference([len(d_34_v2_), len(default__.fm16(globalState))])
            d_36_v4_: _dafny.Map
            d_36_v4_ = _dafny.Map({len(d_35_v3_): d_32_v0_})
            r3 = ((d_34_v2_)[len(_dafny.SeqWithoutIsStrInference([(p3)[default__.safeIndex(p2, len(p3))] for d_37_i0_ in range(default__.abs(12))]))] if (len(_dafny.SeqWithoutIsStrInference([(p3)[default__.safeIndex(p2, len(p3))] for d_38_i0_ in range(default__.abs(12))]))) in (d_34_v2_) else d_32_v0_)
            d_39_v5_: _dafny.Array
            nw0_ = _dafny.Array(None, 13)
            nw0_[int(0)] = p2
            nw0_[int(1)] = -94
            nw0_[int(2)] = p2
            nw0_[int(3)] = 119
            nw0_[int(4)] = 30
            nw0_[int(5)] = (0) - (p2)
            nw0_[int(6)] = p2
            nw0_[int(7)] = (p2) * (p2)
            nw0_[int(8)] = p2
            nw0_[int(9)] = (0) - (p2)
            nw0_[int(10)] = 340
            nw0_[int(11)] = (0) - (p2)
            nw0_[int(12)] = (d_35_v3_)[default__.safeIndex(p2, len(d_35_v3_))]
            d_39_v5_ = nw0_
            index0_ = default__.safeIndex(931, (d_39_v5_).length(0))
            (d_39_v5_)[index0_] = p2
            d_40_v6_: _dafny.Array
            def lambda0_(d_41_v0_):
                def lambda1_(d_42_i1_):
                    return (d_41_v0_) in (_dafny.Set({d_41_v0_}))

                return lambda1_

            init0_ = lambda0_(d_32_v0_)
            nw1_ = _dafny.Array(None, 4)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_40_v6_ = nw1_
            d_43_v7_: _dafny.Seq
            d_43_v7_ = _dafny.SeqWithoutIsStrInference([d_40_v6_, d_40_v6_])
            index1_ = default__.safeIndex(931, (d_39_v5_).length(0))
            rhs0_ = not ((d_32_v0_) and (d_32_v0_)) or (d_32_v0_)
            rhs1_ = p2
            rhs2_ = ((D3_DC11(p2)).cf14) * (p2)
            rhs3_ = d_40_v6_
            rhs4_ = (d_43_v7_)[default__.safeIndex(len(p3), len(d_43_v7_))]
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = d_39_v5_
            lhs3_ = default__.safeIndex(931, (d_39_v5_).length(0))
            lhs0_.f0 = rhs0_
            lhs1_.f3 = rhs1_
            lhs2_[lhs3_] = rhs2_
            d_40_v6_ = rhs3_
            d_40_v6_ = rhs4_
            (globalState).f1 = (d_39_v5_)[default__.safeIndex(931, (d_39_v5_).length(0))]
            d_44_v8_: _dafny.Array
            def lambda2_(d_45_v0_):
                def lambda3_(d_46_i2_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_45_v0_, False, d_45_v0_, d_45_v0_])), _dafny.MultiSet([d_45_v0_])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_45_v0_])]))

                return lambda3_

            init1_ = lambda2_(d_32_v0_)
            nw2_ = _dafny.Array(None, 17)
            for i0_1_ in range(nw2_.length(0)):
                nw2_[i0_1_] = init1_(i0_1_)
            d_44_v8_ = nw2_
            index2_ = default__.safeIndex(618, (d_44_v8_).length(0))
            (d_44_v8_)[index2_] = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_32_v0_, d_32_v0_, d_32_v0_]) for d_47_i3_ in range(default__.abs(337))])
            d_48_v9_: _dafny.Map
            d_48_v9_ = _dafny.Map({_dafny.CodePoint('q'): p3})
            d_49_v10_: str
            d_49_v10_ = _dafny.CodePoint('j')
            index3_ = default__.safeIndex(618, (d_44_v8_).length(0))
            (d_44_v8_)[index3_] = default__.fm23(49, _dafny.CodePoint('y'), (d_48_v9_).set(d_49_v10_, p3), p3, globalState)
            index4_ = default__.safeIndex(931, (d_39_v5_).length(0))
            (d_39_v5_)[index4_] = p2
        elif True:
            d_50_v11_: _dafny.Seq
            d_50_v11_ = _dafny.SeqWithoutIsStrInference([p2])
            d_51_v12_: _dafny.Map
            d_51_v12_ = _dafny.Map({d_50_v11_: p2})
            d_51_v12_ = (d_51_v12_).set(d_50_v11_, p2)
            d_52_v13_: D1
            d_52_v13_ = D1_DC5(d_32_v0_)
            source1_ = d_52_v13_
            if source1_.is_DC5:
                d_53___mcc_h0_ = source1_.cf5
                d_54_cf5_ = d_53___mcc_h0_
                (globalState).f1 = p2
                d_55_v14_: _dafny.Array
                nw3_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_55_v14_ = nw3_
                d_56_v15_: _dafny.Seq
                d_56_v15_ = _dafny.SeqWithoutIsStrInference([d_55_v14_])
                d_57_v16_: C0
                nw4_ = C0()
                nw4_.ctor__((d_56_v15_)[default__.safeIndex(-78, len(d_56_v15_))])
                d_57_v16_ = nw4_
                d_58_v17_: _dafny.MultiSet
                d_58_v17_ = _dafny.MultiSet([p2])
                d_59_v18_: D0
                d_59_v18_ = D0_DC3(d_32_v0_)
                d_60_v19_: str
                d_60_v19_ = _dafny.CodePoint('k')
                d_61_v20_: _dafny.Seq
                d_61_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_60_v19_: d_54_cf5_}), _dafny.Map({d_60_v19_: d_54_cf5_})])
                d_62_v21_: D3
                d_62_v21_ = D3_DC11(p2)
                d_63_v22_: _dafny.Seq
                d_63_v22_ = _dafny.SeqWithoutIsStrInference([d_62_v21_, d_62_v21_])
                d_64_v23_: _dafny.Array
                nw5_ = _dafny.Array(None, 26)
                nw5_[int(0)] = p2
                nw5_[int(1)] = p2
                nw5_[int(2)] = (0) - (p2)
                nw5_[int(3)] = p2
                nw5_[int(4)] = p2
                nw5_[int(5)] = (d_58_v17_).cardinality
                nw5_[int(6)] = len(default__.fm11(d_54_cf5_, d_32_v0_, d_59_v18_, globalState))
                nw5_[int(7)] = 473
                nw5_[int(8)] = p2
                nw5_[int(9)] = (0) - (p2)
                nw5_[int(10)] = p2
                nw5_[int(11)] = p2
                nw5_[int(12)] = p2
                nw5_[int(13)] = p2
                nw5_[int(14)] = -114
                nw5_[int(15)] = p2
                nw5_[int(16)] = (0) - (p2)
                nw5_[int(17)] = len((d_61_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p2, 697])), len(d_61_v20_))])
                nw5_[int(18)] = p2
                nw5_[int(19)] = p2
                nw5_[int(20)] = len(d_63_v22_)
                nw5_[int(21)] = p2
                nw5_[int(22)] = default__.fm1(len(p3), globalState)
                nw5_[int(23)] = (d_50_v11_)[default__.safeIndex(p2, len(d_50_v11_))]
                nw5_[int(24)] = p2
                nw5_[int(25)] = p2
                d_64_v23_ = nw5_
                d_65_v24_: _dafny.Seq
                d_65_v24_ = _dafny.SeqWithoutIsStrInference([d_64_v23_])
                d_65_v24_ = (d_65_v24_) + ((d_65_v24_) + (d_65_v24_))
                (globalState).f1 = (0) - ((427) + (p2))
            elif source1_.is_DC4:
                d_66___mcc_h1_ = source1_.cf4
                d_67_cf4_ = d_66___mcc_h1_
                d_68_v25_: _dafny.Array
                def lambda4_(d_69_p2_):
                    def lambda5_(d_70_i4_):
                        return (d_70_i4_) - ((D6_DC22(d_69_p2_, d_69_p2_)).cf42)

                    return lambda5_

                init2_ = lambda4_(p2)
                nw6_ = _dafny.Array(None, 22)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_68_v25_ = nw6_
                index5_ = default__.safeIndex(415, (d_68_v25_).length(0))
                (d_68_v25_)[index5_] = 183
                d_71_v26_: _dafny.Array
                nw7_ = _dafny.Array(None, 27)
                nw7_[int(0)] = d_32_v0_
                nw7_[int(1)] = d_32_v0_
                nw7_[int(2)] = d_32_v0_
                nw7_[int(3)] = False
                nw7_[int(4)] = d_32_v0_
                nw7_[int(5)] = d_32_v0_
                nw7_[int(6)] = d_32_v0_
                nw7_[int(7)] = not(d_32_v0_)
                nw7_[int(8)] = d_32_v0_
                nw7_[int(9)] = d_32_v0_
                nw7_[int(10)] = d_32_v0_
                nw7_[int(11)] = d_32_v0_
                nw7_[int(12)] = d_32_v0_
                nw7_[int(13)] = d_32_v0_
                nw7_[int(14)] = d_32_v0_
                nw7_[int(15)] = d_32_v0_
                nw7_[int(16)] = d_32_v0_
                nw7_[int(17)] = d_32_v0_
                nw7_[int(18)] = d_32_v0_
                nw7_[int(19)] = d_32_v0_
                nw7_[int(20)] = not(d_32_v0_)
                nw7_[int(21)] = False
                nw7_[int(22)] = d_32_v0_
                nw7_[int(23)] = d_32_v0_
                nw7_[int(24)] = d_32_v0_
                nw7_[int(25)] = d_32_v0_
                nw7_[int(26)] = d_32_v0_
                d_71_v26_ = nw7_
                d_72_v27_: _dafny.Map
                d_72_v27_ = _dafny.Map({d_71_v26_: p3})
                d_73_v28_: _dafny.Map
                d_73_v28_ = _dafny.Map({d_72_v27_: False})
                d_74_v29_: _dafny.Map
                d_74_v29_ = _dafny.Map({(d_50_v11_)[default__.safeIndex(p2, len(d_50_v11_))]: len(p3)})
                d_75_v30_: _dafny.Map
                d_75_v30_ = _dafny.Map({d_74_v29_: d_72_v27_})
                index6_ = default__.safeIndex(415, (d_68_v25_).length(0))
                rhs5_ = (p2 if d_32_v0_ else (len(p0)) - (p2))
                rhs6_ = p2
                rhs7_ = not(((d_73_v28_)[((d_75_v30_)[d_74_v29_] if (d_74_v29_) in (d_75_v30_) else _dafny.Map({d_71_v26_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft"))}))] if (((d_75_v30_)[d_74_v29_] if (d_74_v29_) in (d_75_v30_) else _dafny.Map({d_71_v26_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft"))}))) in (d_73_v28_) else True))
                rhs8_ = p2
                lhs4_ = globalState
                lhs5_ = globalState
                lhs6_ = d_68_v25_
                lhs7_ = default__.safeIndex(415, (d_68_v25_).length(0))
                lhs4_.f1 = rhs5_
                lhs5_.f3 = rhs6_
                r3 = rhs7_
                lhs6_[lhs7_] = rhs8_
                d_76_v31_: _dafny.Set
                d_76_v31_ = _dafny.Set({default__.safeModuloInt(347, default__.fm1(p2, globalState)), (d_68_v25_)[default__.safeIndex(415, (d_68_v25_).length(0))], (d_68_v25_)[default__.safeIndex(415, (d_68_v25_).length(0))]})
                d_77_v32_: _dafny.Set
                d_77_v32_ = _dafny.Set({d_32_v0_, d_32_v0_, d_32_v0_})
                rhs9_ = 362
                rhs10_ = (-108) - (p2)
                rhs11_ = p0
                rhs12_ = default__.safeModuloInt(len(_dafny.Map({True: p2})), (0) - (len(d_77_v32_)))
                lhs8_ = globalState
                r0 = rhs9_
                lhs8_.f3 = rhs10_
                d_76_v31_ = rhs11_
                r0 = rhs12_
                d_78_v33_: _dafny.Map
                d_78_v33_ = _dafny.Map({default__.fm1((0) - ((d_68_v25_)[default__.safeIndex(415, (d_68_v25_).length(0))]), globalState): (p3) + (p3)})
                d_79_v34_: _dafny.Seq
                d_79_v34_ = _dafny.SeqWithoutIsStrInference([p3])
                d_78_v33_ = (d_78_v33_).set((0) - ((-273) - ((0) - (p2))), (d_79_v34_)[default__.safeIndex(len(_dafny.Set({d_32_v0_, d_32_v0_, d_32_v0_, d_32_v0_, default__.fm0((d_68_v25_)[default__.safeIndex(415, (d_68_v25_).length(0))], globalState)})), len(d_79_v34_))])
                index7_ = default__.safeIndex(446, (d_71_v26_).length(0))
                (d_71_v26_)[index7_] = d_32_v0_
                index8_ = default__.safeIndex(446, (d_71_v26_).length(0))
                (d_71_v26_)[index8_] = not(d_32_v0_)
            elif True:
                d_80___mcc_h2_ = source1_.cf6
                d_81_cf6_ = d_80___mcc_h2_
                (globalState).f1 = (p2) - (p2)
                d_82_v35_: _dafny.Set
                d_82_v35_ = _dafny.Set({p2, p2, p2})
                d_82_v35_ = d_82_v35_
                r3 = d_32_v0_
                (globalState).f1 = p2
            d_83_v36_: C1
            nw8_ = C1()
            nw8_.ctor__()
            d_83_v36_ = nw8_
            d_84_v37_: str
            d_84_v37_ = _dafny.CodePoint('j')
            d_85_v38_: _dafny.Map
            d_85_v38_ = _dafny.Map({d_83_v36_: d_84_v37_})
            d_85_v38_ = (d_85_v38_).set(d_83_v36_, d_84_v37_)
            d_83_v36_ = d_83_v36_
            d_86_v39_: C1
            nw9_ = C1()
            nw9_.ctor__()
            d_86_v39_ = nw9_
        d_87_v40_: _dafny.Map
        d_87_v40_ = _dafny.Map({p2: True})
        d_88_v41_: _dafny.Map
        d_88_v41_ = _dafny.Map({d_32_v0_: (p3)[default__.safeIndex(p2, len(p3))]})
        d_89_v42_: _dafny.Seq
        d_89_v42_ = _dafny.SeqWithoutIsStrInference([d_32_v0_])
        d_90_v43_: _dafny.Set
        d_90_v43_ = _dafny.Set({False})
        d_91_v44_: _dafny.Array
        nw10_ = _dafny.Array(None, 11)
        nw10_[int(0)] = _dafny.Set({d_32_v0_})
        nw10_[int(1)] = (_dafny.Set({not(default__.fm0(len(p3), globalState)), d_32_v0_, ((d_87_v40_)[len(d_88_v41_)] if (len(d_88_v41_)) in (d_87_v40_) else d_32_v0_), d_32_v0_})) | (_dafny.Set({d_32_v0_, (d_89_v42_)[default__.safeIndex(p2, len(d_89_v42_))], d_32_v0_}))
        nw10_[int(2)] = (d_90_v43_) - (d_90_v43_)
        nw10_[int(3)] = (d_90_v43_) | (d_90_v43_)
        nw10_[int(4)] = _dafny.Set({d_32_v0_, d_32_v0_, d_32_v0_})
        nw10_[int(5)] = d_90_v43_
        nw10_[int(6)] = d_90_v43_
        nw10_[int(7)] = (_dafny.Set({d_32_v0_, d_32_v0_})).intersection(d_90_v43_)
        nw10_[int(8)] = default__.fm18(p2, globalState)
        nw10_[int(9)] = (_dafny.Set({d_32_v0_})) | (d_90_v43_)
        nw10_[int(10)] = d_90_v43_
        d_91_v44_ = nw10_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_91_v44_).length(0)):
            d_92_i5_: int = guard_loop_0_
            if (True) and (((0) <= (d_92_i5_)) and ((d_92_i5_) < ((d_91_v44_).length(0)))):
                (d_91_v44_)[(d_92_i5_)] = _dafny.Set({d_32_v0_})
        d_93_v45_: str
        d_93_v45_ = _dafny.CodePoint('i')
        d_94_v46_: D5
        d_94_v46_ = D5_DC16(((D6_DC21(p2, d_32_v0_, p3, d_32_v0_)).cf39).set(default__.safeIndex(len(p3), len((D6_DC21(p2, d_32_v0_, p3, d_32_v0_)).cf39)), d_93_v45_))
        d_94_v46_ = d_94_v46_
        (globalState).f3 = default__.fm1((p2) - (p2), globalState)
        d_95_i6_: int
        d_95_i6_ = 0
        with _dafny.label("0"):
            while default__.fm0(p2, globalState):
                with _dafny.c_label("0"):
                    if (d_95_i6_) >= (100):
                        raise _dafny.Break("0")
                    d_95_i6_ = (d_95_i6_) + (1)
                    d_96_v47_: _dafny.Seq
                    d_96_v47_ = _dafny.SeqWithoutIsStrInference([p2, len(default__.fm10(globalState)), p2])
                    d_97_v48_: D5
                    d_97_v48_ = D5_DC17((_dafny.MultiSet(((d_96_v47_).set(default__.safeIndex(p2, len(d_96_v47_)), len(d_90_v43_))) + (_dafny.SeqWithoutIsStrInference([p2 for d_98_i7_ in range(default__.abs(342))])))).cardinality, d_32_v0_, (not(d_32_v0_) if d_32_v0_ else d_32_v0_), (p2) == (p2), (p2) * (p2))
                    d_99_v49_: _dafny.Map
                    d_99_v49_ = _dafny.Map({p2: len(d_96_v47_)})
                    rhs13_ = p2
                    rhs14_ = d_97_v48_
                    rhs15_ = d_99_v49_
                    lhs9_ = globalState
                    lhs9_.f3 = rhs13_
                    d_97_v48_ = rhs14_
                    d_99_v49_ = rhs15_
                    if d_32_v0_:
                        d_100_v50_: _dafny.Array
                        nw11_ = _dafny.Array(False, 23)
                        d_100_v50_ = nw11_
                        d_100_v50_ = d_100_v50_
                        d_101_v51_: C1
                        nw12_ = C1()
                        nw12_.ctor__()
                        d_101_v51_ = nw12_
                        d_102_v52_: _dafny.Set
                        d_102_v52_ = _dafny.Set({d_101_v51_, d_101_v51_})
                        d_96_v47_ = (_dafny.SeqWithoutIsStrInference([p2 for d_103_i8_ in range(default__.abs(16))])) + ((_dafny.SeqWithoutIsStrInference([len(d_102_v52_)])) + (d_96_v47_))
                        d_104_v53_: _dafny.Map
                        d_104_v53_ = _dafny.Map({d_101_v51_: p2})
                        d_104_v53_ = (d_104_v53_).set(d_101_v51_, len(p3))
                        (globalState).f1 = (p2) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_105_i9_ in range(default__.abs(590))])))
                        r3 = d_32_v0_
                    elif True:
                        d_106_v54_: _dafny.Map
                        d_106_v54_ = _dafny.Map({-486: d_89_v42_})
                        d_107_v55_: D6
                        d_107_v55_ = D6_DC20(d_89_v42_, d_32_v0_)
                        d_108_v56_: _dafny.Seq
                        d_108_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_106_v54_)[p2] if (p2) in (d_106_v54_) else d_89_v42_), d_89_v42_, (d_107_v55_).cf35])])
                        (globalState).f1 = (0) - (len((p3).set(default__.safeIndex(len((d_108_v56_)[default__.safeIndex(p2, len(d_108_v56_))]), len(p3)), d_93_v45_)))
                        d_109_v57_: C1
                        nw13_ = C1()
                        nw13_.ctor__()
                        d_109_v57_ = nw13_
                        d_110_v58_: _dafny.Array
                        def lambda6_(d_111_v0_):
                            def lambda7_(d_112_i10_):
                                return d_111_v0_

                            return lambda7_

                        init3_ = lambda6_(d_32_v0_)
                        nw14_ = _dafny.Array(None, 15)
                        for i0_3_ in range(nw14_.length(0)):
                            nw14_[i0_3_] = init3_(i0_3_)
                        d_110_v58_ = nw14_
                        d_113_v59_: _dafny.Set
                        d_113_v59_ = _dafny.Set({d_110_v58_})
                        d_114_v60_: _dafny.Map
                        d_114_v60_ = _dafny.Map({d_32_v0_: -399})
                        d_115_v61_: D6
                        d_115_v61_ = D6_DC22(p2, p2)
                        rhs16_ = d_32_v0_
                        rhs17_ = (p3)[default__.safeIndex(len((d_113_v59_) | (d_113_v59_)), len(p3))]
                        rhs18_ = ((649) + (((d_114_v60_)[False] if (False) in (d_114_v60_) else p2))) < ((d_115_v61_).cf42)
                        r1 = rhs16_
                        d_93_v45_ = rhs17_
                        r1 = rhs18_
                        d_116_v62_: C1
                        nw15_ = C1()
                        nw15_.ctor__()
                        d_116_v62_ = nw15_
                        d_117_v63_: _dafny.Array
                        def lambda8_(d_118_p3_):
                            def lambda9_(d_119_i11_):
                                return d_118_p3_

                            return lambda9_

                        init4_ = lambda8_(p3)
                        nw16_ = _dafny.Array(None, 9)
                        for i0_4_ in range(nw16_.length(0)):
                            nw16_[i0_4_] = init4_(i0_4_)
                        d_117_v63_ = nw16_
                        d_120_v64_: C0
                        nw17_ = C0()
                        nw17_.ctor__(d_117_v63_)
                        d_120_v64_ = nw17_
                    (globalState).f1 = default__.fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywah"))), globalState)
                    if (d_90_v43_).ispropersubset((default__.fm18(p2, globalState)).intersection(d_90_v43_)):
                        r2 = (5) > (-131)
                        d_121_v65_: C1
                        nw18_ = C1()
                        nw18_.ctor__()
                        d_121_v65_ = nw18_
                        d_122_v66_: _dafny.Array
                        def lambda10_(d_123_v47_, d_124_p2_, d_125_p3_, d_126_v45_):
                            def lambda11_(d_127_i12_):
                                return (_dafny.MultiSet([448, d_124_p2_, len((d_125_p3_).set(default__.safeIndex(d_124_p2_, len(d_125_p3_)), d_126_v45_)), d_124_p2_, d_124_p2_])).issubset(_dafny.MultiSet(d_123_v47_))

                            return lambda11_

                        init5_ = lambda10_(d_96_v47_, p2, p3, d_93_v45_)
                        nw19_ = _dafny.Array(None, 5)
                        for i0_5_ in range(nw19_.length(0)):
                            nw19_[i0_5_] = init5_(i0_5_)
                        d_122_v66_ = nw19_
                        d_128_v67_: _dafny.MultiSet
                        d_128_v67_ = _dafny.MultiSet([p3, (p3) + (default__.fm4(True, p2, p2, _dafny.CodePoint('e'), globalState))])
                        d_129_v68_: _dafny.Array
                        nw20_ = _dafny.Array(None, 13)
                        d_129_v68_ = nw20_
                        index9_ = default__.safeIndex(952, (d_129_v68_).length(0))
                        (d_129_v68_)[index9_] = d_121_v65_
                        index10_ = default__.safeIndex(952, (d_129_v68_).length(0))
                        rhs19_ = d_122_v66_
                        rhs20_ = d_128_v67_
                        rhs21_ = d_121_v65_
                        lhs10_ = d_129_v68_
                        lhs11_ = default__.safeIndex(952, (d_129_v68_).length(0))
                        d_122_v66_ = rhs19_
                        d_128_v67_ = rhs20_
                        lhs10_[lhs11_] = rhs21_
                        (globalState).f0 = (p3) <= ((p3) + ((p3).set(default__.safeIndex(902, len(p3)), d_93_v45_)))
                        d_130_v69_: _dafny.Array
                        nw21_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                        d_130_v69_ = nw21_
                        d_131_v70_: C0
                        nw22_ = C0()
                        nw22_.ctor__(d_130_v69_)
                        d_131_v70_ = nw22_
                        d_132_v71_: _dafny.MultiSet
                        d_132_v71_ = _dafny.MultiSet([p2, (0) - (len((_dafny.SeqWithoutIsStrInference([d_131_v70_])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_131_v70_]))), d_131_v70_))), p2, p2])
                        d_133_v72_: D0
                        d_133_v72_ = D0_DC3(d_32_v0_)
                        d_134_v73_: int
                        out0_: int
                        out0_ = (d_121_v65_).m1((D0_DC3(d_32_v0_)).cf3, _dafny.Set({(d_121_v65_).fm2(d_132_v71_, globalState)}), (p3) + (p3), d_133_v72_, globalState)
                        d_134_v73_ = out0_
                    elif True:
                        d_135_v74_: _dafny.Array
                        nw23_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                        d_135_v74_ = nw23_
                        d_136_v75_: _dafny.Map
                        d_136_v75_ = _dafny.Map({d_135_v74_: (236) - (p2)})
                        d_136_v75_ = (d_136_v75_).set(d_135_v74_, p2)
                        d_137_v76_: _dafny.Array
                        nw24_ = _dafny.Array(int(0), 27)
                        d_137_v76_ = nw24_
                        index11_ = default__.safeIndex(186, (d_137_v76_).length(0))
                        (d_137_v76_)[index11_] = p2
                        d_138_v77_: D1
                        d_138_v77_ = D1_DC5(not(not(d_32_v0_)))
                        d_139_v78_: _dafny.Array
                        nw25_ = _dafny.Array(False, 15)
                        d_139_v78_ = nw25_
                        index12_ = default__.safeIndex(486, (d_139_v78_).length(0))
                        (d_139_v78_)[index12_] = d_32_v0_
                        index13_ = default__.safeIndex(186, (d_137_v76_).length(0))
                        index14_ = default__.safeIndex(486, (d_139_v78_).length(0))
                        rhs22_ = p2
                        rhs23_ = (273 if d_32_v0_ else -363)
                        rhs24_ = d_138_v77_
                        rhs25_ = not((p2) < (375))
                        lhs12_ = d_137_v76_
                        lhs13_ = default__.safeIndex(186, (d_137_v76_).length(0))
                        lhs14_ = globalState
                        lhs15_ = d_139_v78_
                        lhs16_ = default__.safeIndex(486, (d_139_v78_).length(0))
                        lhs12_[lhs13_] = rhs22_
                        lhs14_.f3 = rhs23_
                        d_138_v77_ = rhs24_
                        lhs15_[lhs16_] = rhs25_
                        d_140_v79_: D5
                        d_140_v79_ = D5_DC18(len(_dafny.SeqWithoutIsStrInference([d_93_v45_ for d_141_i13_ in range(default__.abs(994))])), len(p3))
                        d_142_v80_: _dafny.Map
                        d_142_v80_ = _dafny.Map({p3: d_140_v79_})
                        d_142_v80_ = (d_142_v80_).set(p3, d_140_v79_)
                        index15_ = default__.safeIndex(186, (d_137_v76_).length(0))
                        (d_137_v76_)[index15_] = default__.safeModuloInt(p2, ((d_137_v76_)[default__.safeIndex(186, (d_137_v76_).length(0))] if d_32_v0_ else p2))
                        d_143_v81_: C0
                        nw26_ = C0()
                        nw26_.ctor__(d_135_v74_)
                        d_143_v81_ = nw26_
                    pass
            pass
        d_144_v82_: C1
        nw27_ = C1()
        nw27_.ctor__()
        d_144_v82_ = nw27_
        d_145_v83_: _dafny.Set
        d_145_v83_ = _dafny.Set({d_144_v82_, d_144_v82_})
        r0 = len(d_145_v83_)
        d_146_v84_: _dafny.Seq
        d_146_v84_ = _dafny.SeqWithoutIsStrInference([d_87_v40_])
        d_147_v85_: _dafny.Seq
        d_147_v85_ = _dafny.SeqWithoutIsStrInference([len((p3).set(default__.safeIndex(len((d_146_v84_)[default__.safeIndex(718, len(d_146_v84_))]), len(p3)), d_93_v45_)), p2])
        d_148_v86_: _dafny.Set
        d_148_v86_ = _dafny.Set({d_147_v85_})
        d_149_v87_: _dafny.Seq
        d_149_v87_ = _dafny.SeqWithoutIsStrInference([d_147_v85_])
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: _dafny.Seq
            for compr_8_ in (d_149_v87_).Elements:
                d_151_v88_: _dafny.Seq = compr_8_
                if (d_151_v88_) in (d_149_v87_):
                    coll8_ = coll8_.union(_dafny.Set([d_151_v88_]))
            return _dafny.Set(coll8_)
        r1 = ((_dafny.Set({_dafny.SeqWithoutIsStrInference([len(d_88_v41_) for d_150_i14_ in range(default__.abs(73))])})) - (iife8_()
        )).issubset(d_148_v86_)
        r2 = d_32_v0_
        d_152_v89_: _dafny.Array
        nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
        d_152_v89_ = nw28_
        d_153_v90_: C0
        nw29_ = C0()
        nw29_.ctor__(d_152_v89_)
        d_153_v90_ = nw29_
        d_154_v91_: _dafny.Seq
        d_154_v91_ = _dafny.SeqWithoutIsStrInference([d_153_v90_])
        r3 = default__.fm0(default__.safeDivisionInt(p2, len(d_154_v91_)), globalState)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_155_v0_: _dafny.Array
        def lambda12_(d_156_i0_):
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))

        init6_ = lambda12_
        nw30_ = _dafny.Array(None, 13)
        for i0_6_ in range(nw30_.length(0)):
            nw30_[i0_6_] = init6_(i0_6_)
        d_155_v0_ = nw30_
        d_157_v1_: bool
        d_157_v1_ = False
        d_158_v2_: _dafny.Seq
        d_158_v2_ = _dafny.SeqWithoutIsStrInference([d_157_v1_, not(d_157_v1_)])
        d_159_globalState_: GlobalState
        nw31_ = GlobalState()
        nw31_.ctor__(True, 827, True, -651, d_155_v0_, (d_158_v2_) + (d_158_v2_))
        d_159_globalState_ = nw31_
        d_160_v3_: int
        d_160_v3_ = 272
        d_161_v4_: _dafny.Map
        d_161_v4_ = _dafny.Map({d_160_v3_: d_160_v3_})
        hi0_ = (d_160_v3_) - (((d_161_v4_)[(0) - (d_160_v3_)] if ((0) - (d_160_v3_)) in (d_161_v4_) else d_160_v3_))
        for d_162_i1_ in range(d_160_v3_, hi0_):
            d_163_v5_: str
            d_163_v5_ = _dafny.CodePoint('k')
            d_164_v6_: _dafny.MultiSet
            d_164_v6_ = _dafny.MultiSet([d_163_v5_, d_163_v5_])
            d_165_v7_: _dafny.Seq
            d_165_v7_ = _dafny.SeqWithoutIsStrInference([d_164_v6_, d_164_v6_, _dafny.MultiSet([_dafny.CodePoint('x')])])
            d_166_v8_: D0
            d_166_v8_ = D0_DC0(d_160_v3_)
            d_165_v7_ = (d_165_v7_) + ((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_163_v5_, d_163_v5_]) for d_167_i2_ in range(default__.abs(641))])).set(default__.safeIndex((d_166_v8_).cf0, len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_163_v5_, d_163_v5_]) for d_168_i2_ in range(default__.abs(641))]))), d_164_v6_))
            d_161_v4_ = (d_161_v4_).set(d_162_i1_, (d_162_i1_) - ((0) - (d_162_i1_)))
            d_169_v9_: _dafny.Set
            d_169_v9_ = _dafny.Set({default__.fm0(len(d_158_v2_), d_159_globalState_), d_157_v1_})
            d_170_v10_: _dafny.MultiSet
            d_170_v10_ = _dafny.MultiSet([d_169_v9_, d_169_v9_])
            d_171_v11_: D0
            d_171_v11_ = D0_DC2((d_170_v10_).ispropersubset(d_170_v10_))
            d_172_v12_: _dafny.Seq
            d_172_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_171_v11_ = (d_171_v11_ if (len((d_172_v12_).set(default__.safeIndex(d_160_v3_, len(d_172_v12_)), d_163_v5_))) != (d_160_v3_) else d_171_v11_)
            (d_159_globalState_).f3 = d_162_i1_
        d_173_v13_: _dafny.Map
        d_173_v13_ = _dafny.Map({d_157_v1_: (0) - (d_160_v3_)})
        d_174_v14_: D0
        d_174_v14_ = D0_DC2(True)
        d_173_v13_ = (d_173_v13_).set((d_174_v14_).cf2, d_160_v3_)
        d_157_v1_ = d_157_v1_
        d_175_v15_: _dafny.Array
        def lambda13_(d_176_v3_):
            def lambda14_(d_177_i3_):
                return (d_177_i3_) + (len(_dafny.Set({d_176_v3_})))

            return lambda14_

        init7_ = lambda13_(d_160_v3_)
        nw32_ = _dafny.Array(None, 20)
        for i0_7_ in range(nw32_.length(0)):
            nw32_[i0_7_] = init7_(i0_7_)
        d_175_v15_ = nw32_
        index16_ = default__.safeIndex(743, (d_175_v15_).length(0))
        (d_175_v15_)[index16_] = d_160_v3_
        index17_ = default__.safeIndex(743, (d_175_v15_).length(0))
        (d_175_v15_)[index17_] = d_160_v3_
        d_178_v16_: _dafny.Map
        d_178_v16_ = _dafny.Map({d_175_v15_: -954})
        hi1_ = len(d_178_v16_)
        for d_179_i4_ in range(default__.fm1(d_160_v3_, d_159_globalState_), hi1_):
            d_180_v17_: _dafny.Seq
            d_180_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldovmaed"))
            d_181_v18_: _dafny.Map
            d_181_v18_ = _dafny.Map({d_160_v3_: (d_180_v17_) + (d_180_v17_)})
            d_181_v18_ = (d_181_v18_).set(default__.safeDivisionInt(d_160_v3_, default__.fm1(d_160_v3_, d_159_globalState_)), d_180_v17_)
            d_182_v19_: _dafny.Map
            d_182_v19_ = _dafny.Map({len(d_161_v4_): d_157_v1_})
            d_183_v20_: _dafny.Set
            d_183_v20_ = _dafny.Set({d_157_v1_, d_157_v1_})
            d_157_v1_ = ((d_182_v19_)[d_179_i4_] if (d_179_i4_) in (d_182_v19_) else (d_183_v20_).ispropersubset(d_183_v20_))
            if False:
                d_184_v21_: C1
                nw33_ = C1()
                nw33_.ctor__()
                d_184_v21_ = nw33_
                d_185_v22_: _dafny.MultiSet
                d_185_v22_ = _dafny.MultiSet([((d_181_v18_)[(d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]] if ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]) in (d_181_v18_) else d_180_v17_)])
                d_186_v23_: _dafny.Array
                nw34_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_186_v23_ = nw34_
                d_187_v24_: D3
                d_187_v24_ = D3_DC12(d_160_v3_, d_185_v22_, len(d_158_v2_), d_186_v23_)
                d_188_v25_: _dafny.MultiSet
                d_188_v25_ = _dafny.MultiSet([d_187_v24_, d_187_v24_])
                d_189_v26_: _dafny.Array
                nw35_ = _dafny.Array(False, 14)
                d_189_v26_ = nw35_
                d_190_v27_: _dafny.Map
                d_190_v27_ = _dafny.Map({d_160_v3_: d_189_v26_})
                d_191_v28_: _dafny.Map
                d_191_v28_ = _dafny.Map({d_188_v25_: ((d_190_v27_)[d_160_v3_] if (d_160_v3_) in (d_190_v27_) else d_189_v26_)})
                d_191_v28_ = (d_191_v28_).set(d_188_v25_, d_189_v26_)
                d_192_v29_: _dafny.Map
                d_192_v29_ = _dafny.Map({((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]) - (d_160_v3_): (d_175_v15_ if d_157_v1_ else d_175_v15_)})
                d_175_v15_ = ((d_192_v29_)[d_179_i4_] if (d_179_i4_) in (d_192_v29_) else d_175_v15_)
                d_193_v30_: int
                d_194_v31_: _dafny.Array
                d_195_v32_: D0
                out1_: int
                out2_: _dafny.Array
                out3_: D0
                out1_, out2_, out3_ = (d_184_v21_).m0(d_159_globalState_)
                d_193_v30_ = out1_
                d_194_v31_ = out2_
                d_195_v32_ = out3_
                d_196_v33_: _dafny.Array
                def lambda15_(d_197_v17_):
                    def lambda16_(d_198_i5_):
                        return d_197_v17_

                    return lambda16_

                init8_ = lambda15_(d_180_v17_)
                nw36_ = _dafny.Array(None, 29)
                for i0_8_ in range(nw36_.length(0)):
                    nw36_[i0_8_] = init8_(i0_8_)
                d_196_v33_ = nw36_
                d_199_v34_: C0
                nw37_ = C0()
                nw37_.ctor__(d_196_v33_)
                d_199_v34_ = nw37_
                d_200_v35_: _dafny.Array
                nw38_ = _dafny.Array(None, 28)
                nw38_[int(0)] = d_199_v34_
                nw38_[int(1)] = d_199_v34_
                nw38_[int(2)] = d_199_v34_
                nw38_[int(3)] = d_199_v34_
                nw38_[int(4)] = d_199_v34_
                nw38_[int(5)] = d_199_v34_
                nw38_[int(6)] = d_199_v34_
                nw38_[int(7)] = d_199_v34_
                nw38_[int(8)] = d_199_v34_
                nw38_[int(9)] = d_199_v34_
                nw38_[int(10)] = d_199_v34_
                nw38_[int(11)] = d_199_v34_
                nw38_[int(12)] = d_199_v34_
                nw38_[int(13)] = d_199_v34_
                nw38_[int(14)] = d_199_v34_
                nw38_[int(15)] = d_199_v34_
                nw38_[int(16)] = d_199_v34_
                nw38_[int(17)] = d_199_v34_
                nw38_[int(18)] = (d_199_v34_ if d_157_v1_ else d_199_v34_)
                nw38_[int(19)] = d_199_v34_
                nw38_[int(20)] = d_199_v34_
                nw38_[int(21)] = d_199_v34_
                nw38_[int(22)] = d_199_v34_
                nw38_[int(23)] = d_199_v34_
                nw38_[int(24)] = d_199_v34_
                nw38_[int(25)] = d_199_v34_
                nw38_[int(26)] = d_199_v34_
                nw38_[int(27)] = d_199_v34_
                d_200_v35_ = nw38_
                d_200_v35_ = d_200_v35_
            elif True:
                (d_159_globalState_).f0 = not(not (d_157_v1_) or ((d_157_v1_ if True else d_157_v1_)))
                index18_ = default__.safeIndex(743, (d_175_v15_).length(0))
                (d_175_v15_)[index18_] = len((d_182_v19_) | (d_182_v19_))
                d_157_v1_ = not (d_157_v1_) or ((d_158_v2_) <= (d_158_v2_))
                (d_159_globalState_).f0 = default__.fm0((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_159_globalState_)
                d_201_v36_: _dafny.Array
                def lambda17_(d_202_v1_):
                    def lambda18_(d_203_i6_):
                        return d_202_v1_

                    return lambda18_

                init9_ = lambda17_(d_157_v1_)
                nw39_ = _dafny.Array(None, 13)
                for i0_9_ in range(nw39_.length(0)):
                    nw39_[i0_9_] = init9_(i0_9_)
                d_201_v36_ = nw39_
                d_201_v36_ = d_201_v36_
            d_204_v37_: _dafny.Seq
            d_204_v37_ = _dafny.SeqWithoutIsStrInference([d_161_v4_, (d_161_v4_).set(len(d_180_v17_), d_179_i4_)])
            d_205_v38_: D3
            d_205_v38_ = D3_DC9(d_204_v37_)
            d_206_v39_: _dafny.Seq
            d_206_v39_ = _dafny.SeqWithoutIsStrInference([d_205_v38_])
            d_207_v40_: _dafny.Map
            d_207_v40_ = _dafny.Map({d_206_v39_: d_180_v17_})
            rhs26_ = d_207_v40_
            rhs27_ = d_175_v15_
            d_207_v40_ = rhs26_
            d_175_v15_ = rhs27_
        d_208_v41_: _dafny.Array
        nw40_ = _dafny.Array(False, 22)
        d_208_v41_ = nw40_
        index19_ = default__.safeIndex(949, (d_208_v41_).length(0))
        (d_208_v41_)[index19_] = default__.fm0((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_159_globalState_)
        d_209_v42_: _dafny.MultiSet
        d_209_v42_ = _dafny.MultiSet([d_157_v1_])
        pat_let_tv0_ = d_160_v3_
        pat_let_tv1_ = d_175_v15_
        pat_let_tv2_ = d_175_v15_
        pat_let_tv3_ = d_209_v42_
        pat_let_tv4_ = d_160_v3_
        pat_let_tv5_ = d_157_v1_
        pat_let_tv6_ = d_160_v3_
        index20_ = default__.safeIndex(949, (d_208_v41_).length(0))
        def lambda19_(source2_):
            if source2_.is_DC5:
                d_210___mcc_h0_ = source2_.cf5
                d_211_cf5_ = d_210___mcc_h0_
                return (_dafny.MultiSet([pat_let_tv0_, (pat_let_tv2_)[default__.safeIndex(743, (pat_let_tv1_).length(0))], (pat_let_tv3_).cardinality])).isdisjoint(_dafny.MultiSet([pat_let_tv4_]))
            elif source2_.is_DC4:
                d_212___mcc_h1_ = source2_.cf4
                d_213_cf4_ = d_212___mcc_h1_
                return pat_let_tv5_
            elif True:
                d_214___mcc_h2_ = source2_.cf6
                d_215_cf6_ = d_214___mcc_h2_
                return (371) <= (pat_let_tv6_)

        (d_208_v41_)[index20_] = lambda19_(D1_DC4(d_209_v42_))
        d_216_i7_: int
        d_216_i7_ = 0
        with _dafny.label("1"):
            while (d_160_v3_) < (d_160_v3_):
                with _dafny.c_label("1"):
                    if (d_216_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_216_i7_ = (d_216_i7_) + (1)
                    d_217_v43_: _dafny.Array
                    nw41_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_217_v43_ = nw41_
                    index21_ = default__.safeIndex(203, (d_217_v43_).length(0))
                    (d_217_v43_)[index21_] = d_208_v41_
                    d_218_v44_: str
                    d_218_v44_ = _dafny.CodePoint('w')
                    d_219_v45_: _dafny.MultiSet
                    d_219_v45_ = _dafny.MultiSet([d_218_v44_, d_218_v44_, d_218_v44_, d_218_v44_])
                    index22_ = default__.safeIndex(203, (d_217_v43_).length(0))
                    nw42_ = _dafny.Array(None, 12)
                    nw42_[int(0)] = ((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] if d_157_v1_ else (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                    nw42_[int(1)] = d_157_v1_
                    nw42_[int(2)] = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                    nw42_[int(3)] = d_157_v1_
                    nw42_[int(4)] = (d_219_v45_).issubset(d_219_v45_)
                    nw42_[int(5)] = ((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] if d_157_v1_ else d_157_v1_)
                    nw42_[int(6)] = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                    nw42_[int(7)] = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                    nw42_[int(8)] = d_157_v1_
                    nw42_[int(9)] = d_157_v1_
                    nw42_[int(10)] = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                    nw42_[int(11)] = d_157_v1_
                    (d_217_v43_)[index22_] = nw42_
                    if False:
                        (d_159_globalState_).f0 = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                        d_220_v46_: D1
                        d_220_v46_ = D1_DC4(d_209_v42_)
                        d_221_v47_: _dafny.Seq
                        d_221_v47_ = _dafny.SeqWithoutIsStrInference([d_160_v3_, d_160_v3_, d_160_v3_, d_160_v3_])
                        index23_ = default__.safeIndex(743, (d_175_v15_).length(0))
                        rhs28_ = d_220_v46_
                        rhs29_ = default__.fm1((0) - (len(d_221_v47_)), d_159_globalState_)
                        rhs30_ = d_160_v3_
                        lhs17_ = d_159_globalState_
                        lhs18_ = d_175_v15_
                        lhs19_ = default__.safeIndex(743, (d_175_v15_).length(0))
                        d_220_v46_ = rhs28_
                        lhs17_.f3 = rhs29_
                        lhs18_[lhs19_] = rhs30_
                        d_222_v48_: _dafny.Seq
                        d_222_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tabouo"))
                        d_223_v49_: C1
                        nw43_ = C1()
                        nw43_.ctor__()
                        d_223_v49_ = nw43_
                        d_224_v50_: _dafny.MultiSet
                        d_224_v50_ = _dafny.MultiSet([d_223_v49_])
                        d_225_v51_: _dafny.Map
                        d_225_v51_ = _dafny.Map({len(d_222_v48_): (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]})
                        rhs31_ = (d_222_v48_) + (d_222_v48_)
                        rhs32_ = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                        rhs33_ = (d_160_v3_) > (((0) - (d_160_v3_)) * ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]))
                        rhs34_ = (default__.safeModuloInt(len(d_222_v48_), (d_224_v50_).cardinality)) in (d_225_v51_)
                        rhs35_ = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
                        lhs20_ = d_159_globalState_
                        lhs21_ = d_159_globalState_
                        d_222_v48_ = rhs31_
                        lhs20_.f0 = rhs32_
                        d_157_v1_ = rhs33_
                        d_157_v1_ = rhs34_
                        lhs21_.f1 = rhs35_
                        d_226_v52_: C1
                        nw44_ = C1()
                        nw44_.ctor__()
                        d_226_v52_ = nw44_
                        d_227_v53_: int
                        d_228_v54_: _dafny.Array
                        d_229_v55_: D0
                        out4_: int
                        out5_: _dafny.Array
                        out6_: D0
                        out4_, out5_, out6_ = (d_223_v49_).m0(d_159_globalState_)
                        d_227_v53_ = out4_
                        d_228_v54_ = out5_
                        d_229_v55_ = out6_
                    elif True:
                        d_230_v56_: _dafny.Set
                        d_230_v56_ = _dafny.Set({d_157_v1_})
                        d_231_v57_: _dafny.Set
                        d_231_v57_ = d_230_v56_
                        d_232_v58_: _dafny.Array
                        nw45_ = _dafny.Array(None, 9)
                        nw45_[int(0)] = d_230_v56_
                        nw45_[int(1)] = (d_231_v57_)
                        nw45_[int(2)] = d_230_v56_
                        nw45_[int(3)] = d_230_v56_
                        nw45_[int(4)] = default__.fm18(914, d_159_globalState_)
                        nw45_[int(5)] = d_230_v56_
                        nw45_[int(6)] = (_dafny.Set({not((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]), d_157_v1_})).intersection(_dafny.Set({d_157_v1_}))
                        nw45_[int(7)] = (d_230_v56_).intersection(d_230_v56_)
                        nw45_[int(8)] = d_230_v56_
                        d_232_v58_ = nw45_
                        index24_ = default__.safeIndex(606, (d_232_v58_).length(0))
                        (d_232_v58_)[index24_] = default__.fm18(359, d_159_globalState_)
                        index25_ = default__.safeIndex(606, (d_232_v58_).length(0))
                        (d_232_v58_)[index25_] = _dafny.Set({(d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]})
                        d_233_v59_: _dafny.Array
                        nw46_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                        d_233_v59_ = nw46_
                        index26_ = default__.safeIndex(403, (d_233_v59_).length(0))
                        (d_233_v59_)[index26_] = d_218_v44_
                        index27_ = default__.safeIndex(403, (d_233_v59_).length(0))
                        (d_233_v59_)[index27_] = d_218_v44_
                        d_234_v60_: _dafny.Set
                        d_234_v60_ = _dafny.Set({d_160_v3_, (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_160_v3_, 221})
                        d_235_v61_: _dafny.Array
                        def lambda20_(d_236_v2_, d_237_v1_):
                            def lambda21_(d_238_i8_):
                                return D6_DC20(d_236_v2_, d_237_v1_)

                            return lambda21_

                        init10_ = lambda20_(d_158_v2_, d_157_v1_)
                        nw47_ = _dafny.Array(None, 23)
                        for i0_10_ in range(nw47_.length(0)):
                            nw47_[i0_10_] = init10_(i0_10_)
                        d_235_v61_ = nw47_
                        d_239_v62_: _dafny.Seq
                        d_239_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdhmm"))
                        d_240_v63_: int
                        d_241_v64_: bool
                        d_242_v65_: bool
                        d_243_v66_: bool
                        out7_: int
                        out8_: bool
                        out9_: bool
                        out10_: bool
                        out7_, out8_, out9_, out10_ = default__.m2(d_234_v60_, d_235_v61_, (len((d_232_v58_)[default__.safeIndex(606, (d_232_v58_).length(0))]) if (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] else d_160_v3_), ((d_239_v62_) + (d_239_v62_)).set(default__.safeIndex((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], len((d_239_v62_) + (d_239_v62_))), (d_233_v59_)[default__.safeIndex(403, (d_233_v59_).length(0))]), d_159_globalState_)
                        d_240_v63_ = out7_
                        d_241_v64_ = out8_
                        d_242_v65_ = out9_
                        d_243_v66_ = out10_
                        d_244_v67_: int
                        d_245_v68_: bool
                        d_246_v69_: bool
                        d_247_v70_: bool
                        out11_: int
                        out12_: bool
                        out13_: bool
                        out14_: bool
                        out11_, out12_, out13_, out14_ = default__.m2((d_234_v60_) - (default__.fm19((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], d_159_globalState_)), d_235_v61_, ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]) - ((0) - (d_160_v3_)), d_239_v62_, d_159_globalState_)
                        d_244_v67_ = out11_
                        d_245_v68_ = out12_
                        d_246_v69_ = out13_
                        d_247_v70_ = out14_
                        d_247_v70_ = (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]
                    if d_157_v1_:
                        d_248_v71_: _dafny.Map
                        d_248_v71_ = _dafny.Map({d_160_v3_: not(d_157_v1_)})
                        d_248_v71_ = (d_248_v71_).set((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_157_v1_)
                        d_249_v72_: C1
                        nw48_ = C1()
                        nw48_.ctor__()
                        d_249_v72_ = nw48_
                        d_250_v73_: _dafny.Map
                        d_250_v73_ = _dafny.Map({len(_dafny.Map({(d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]: True})): d_249_v72_})
                        d_251_v74_: D6
                        d_251_v74_ = D6_DC20(d_158_v2_, d_157_v1_)
                        d_252_v75_: _dafny.Array
                        nw49_ = _dafny.Array(None, 24)
                        nw49_[int(0)] = d_251_v74_
                        nw49_[int(1)] = d_251_v74_
                        nw49_[int(2)] = d_251_v74_
                        nw49_[int(3)] = D6_DC20(d_158_v2_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                        nw49_[int(4)] = d_251_v74_
                        nw49_[int(5)] = d_251_v74_
                        nw49_[int(6)] = d_251_v74_
                        nw49_[int(7)] = d_251_v74_
                        nw49_[int(8)] = d_251_v74_
                        nw49_[int(9)] = d_251_v74_
                        nw49_[int(10)] = d_251_v74_
                        nw49_[int(11)] = d_251_v74_
                        nw49_[int(12)] = d_251_v74_
                        nw49_[int(13)] = d_251_v74_
                        nw49_[int(14)] = d_251_v74_
                        nw49_[int(15)] = d_251_v74_
                        nw49_[int(16)] = d_251_v74_
                        nw49_[int(17)] = d_251_v74_
                        nw49_[int(18)] = D6_DC20(d_158_v2_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                        nw49_[int(19)] = d_251_v74_
                        nw49_[int(20)] = d_251_v74_
                        nw49_[int(21)] = d_251_v74_
                        nw49_[int(22)] = D6_DC20(d_158_v2_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                        nw49_[int(23)] = d_251_v74_
                        d_252_v75_ = nw49_
                        d_253_v76_: _dafny.Seq
                        d_253_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quqguci"))
                        d_254_v77_: int
                        d_255_v78_: bool
                        d_256_v79_: bool
                        d_257_v80_: bool
                        out15_: int
                        out16_: bool
                        out17_: bool
                        out18_: bool
                        out15_, out16_, out17_, out18_ = default__.m2(_dafny.Set({len(d_250_v73_), (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_160_v3_, (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]}), (d_252_v75_ if (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] else d_252_v75_), d_160_v3_, d_253_v76_, d_159_globalState_)
                        d_254_v77_ = out15_
                        d_255_v78_ = out16_
                        d_256_v79_ = out17_
                        d_257_v80_ = out18_
                        d_258_v81_: _dafny.Seq
                        d_258_v81_ = _dafny.SeqWithoutIsStrInference([295, d_254_v77_, d_254_v77_])
                        d_256_v79_ = not(((d_160_v3_) + ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))])) not in (d_258_v81_))
                        d_259_v82_: _dafny.Set
                        d_259_v82_ = _dafny.Set({d_254_v77_, d_160_v3_})
                        d_260_v83_: _dafny.Map
                        d_260_v83_ = _dafny.Map({d_256_v79_: d_259_v82_})
                        nw50_ = _dafny.Array(None, 24)
                        nw50_[int(0)] = d_254_v77_
                        nw50_[int(1)] = d_254_v77_
                        nw50_[int(2)] = len(((d_260_v83_)[d_255_v78_] if (d_255_v78_) in (d_260_v83_) else default__.fm19((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], d_159_globalState_)))
                        nw50_[int(3)] = (d_160_v3_) + (d_160_v3_)
                        nw50_[int(4)] = d_254_v77_
                        nw50_[int(5)] = (338) - (39)
                        nw50_[int(6)] = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
                        nw50_[int(7)] = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
                        nw50_[int(8)] = (d_160_v3_) * (d_254_v77_)
                        nw50_[int(9)] = d_160_v3_
                        nw50_[int(10)] = (d_254_v77_) - (d_254_v77_)
                        nw50_[int(11)] = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
                        nw50_[int(12)] = d_160_v3_
                        nw50_[int(13)] = 525
                        nw50_[int(14)] = (d_160_v3_) - (d_254_v77_)
                        nw50_[int(15)] = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
                        nw50_[int(16)] = (d_254_v77_) - ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))])
                        nw50_[int(17)] = (0) - (d_160_v3_)
                        nw50_[int(18)] = ((d_161_v4_)[(d_209_v42_).cardinality] if ((d_209_v42_).cardinality) in (d_161_v4_) else (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))])
                        nw50_[int(19)] = len(_dafny.Map({d_255_v78_: d_255_v78_}))
                        nw50_[int(20)] = d_160_v3_
                        nw50_[int(21)] = d_254_v77_
                        nw50_[int(22)] = d_254_v77_
                        nw50_[int(23)] = len(d_158_v2_)
                        d_175_v15_ = nw50_
                        d_257_v80_ = not(default__.fm0(default__.fm1(len(_dafny.SeqWithoutIsStrInference([d_218_v44_ for d_261_i9_ in range(default__.abs(344))])), d_159_globalState_), d_159_globalState_))
                    elif True:
                        d_262_v84_: _dafny.MultiSet
                        d_262_v84_ = _dafny.MultiSet([d_208_v41_, d_208_v41_])
                        (d_159_globalState_).f0 = (d_262_v84_) != (d_262_v84_)
                        d_263_v85_: D2
                        d_263_v85_ = D2_DC8(default__.safeModuloInt(d_160_v3_, (_dafny.MultiSet([31])).cardinality), -850)
                        d_263_v85_ = d_263_v85_
                        d_264_v86_: D4
                        d_264_v86_ = D4_DC13((d_217_v43_)[default__.safeIndex(203, (d_217_v43_).length(0))])
                        d_265_v87_: _dafny.Map
                        d_265_v87_ = _dafny.Map({d_157_v1_: d_264_v86_})
                        d_266_v88_: _dafny.Array
                        def lambda22_(d_267_v44_):
                            def lambda23_(d_268_i10_):
                                return _dafny.SeqWithoutIsStrInference([d_267_v44_ for d_269_i11_ in range(default__.abs(571))])

                            return lambda23_

                        init11_ = lambda22_(d_218_v44_)
                        nw51_ = _dafny.Array(None, 13)
                        for i0_11_ in range(nw51_.length(0)):
                            nw51_[i0_11_] = init11_(i0_11_)
                        d_266_v88_ = nw51_
                        d_270_v89_: C0
                        nw52_ = C0()
                        nw52_.ctor__(d_266_v88_)
                        d_270_v89_ = nw52_
                        d_271_v90_: _dafny.Map
                        d_271_v90_ = _dafny.Map({d_265_v87_: d_270_v89_})
                        d_271_v90_ = (d_271_v90_).set((_dafny.Map({d_157_v1_: d_264_v86_})) | (d_265_v87_), d_270_v89_)
                        d_272_v91_: _dafny.Set
                        d_272_v91_ = _dafny.Set({(d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_160_v3_])), (0) - (d_160_v3_), (0) - (d_160_v3_)})
                        d_273_v92_: D6
                        d_273_v92_ = D6_DC20(d_158_v2_, d_157_v1_)
                        d_274_v93_: _dafny.Map
                        d_274_v93_ = _dafny.Map({d_218_v44_: (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]})
                        pat_let_tv7_ = d_208_v41_
                        pat_let_tv8_ = d_208_v41_
                        d_275_v94_: _dafny.Array
                        nw53_ = _dafny.Array(None, 28)
                        nw53_[int(0)] = d_273_v92_
                        nw53_[int(1)] = d_273_v92_
                        nw53_[int(2)] = d_273_v92_
                        nw53_[int(3)] = d_273_v92_
                        nw53_[int(4)] = d_273_v92_
                        nw53_[int(5)] = d_273_v92_
                        def iife9_(_pat_let0_0):
                            def iife10_(d_276_dt__update__tmp_h0_):
                                def iife11_(_pat_let1_0):
                                    def iife12_(d_277_dt__update_hcf35_h0_):
                                        return D6_DC20(d_277_dt__update_hcf35_h0_, (d_276_dt__update__tmp_h0_).cf36)
                                    return iife12_(_pat_let1_0)
                                return iife11_(_dafny.SeqWithoutIsStrInference([(pat_let_tv8_)[default__.safeIndex(949, (pat_let_tv7_).length(0))]]))
                            return iife10_(_pat_let0_0)
                        nw53_[int(6)] = iife9_(d_273_v92_)
                        nw53_[int(7)] = d_273_v92_
                        nw53_[int(8)] = d_273_v92_
                        nw53_[int(9)] = d_273_v92_
                        nw53_[int(10)] = d_273_v92_
                        nw53_[int(11)] = d_273_v92_
                        nw53_[int(12)] = d_273_v92_
                        nw53_[int(13)] = d_273_v92_
                        nw53_[int(14)] = d_273_v92_
                        nw53_[int(15)] = D6_DC20(d_158_v2_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                        nw53_[int(16)] = d_273_v92_
                        nw53_[int(17)] = d_273_v92_
                        nw53_[int(18)] = default__.fm20(d_274_v93_, d_157_v1_, (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_159_globalState_)
                        nw53_[int(19)] = d_273_v92_
                        nw53_[int(20)] = d_273_v92_
                        nw53_[int(21)] = d_273_v92_
                        nw53_[int(22)] = d_273_v92_
                        nw53_[int(23)] = D6_DC20(_dafny.SeqWithoutIsStrInference([d_157_v1_, False, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], (D5_DC17((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_157_v1_, d_157_v1_, False, (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))])).cf29, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]]), d_157_v1_)
                        nw53_[int(24)] = d_273_v92_
                        nw53_[int(25)] = d_273_v92_
                        nw53_[int(26)] = d_273_v92_
                        nw53_[int(27)] = D6_DC20(d_158_v2_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))])
                        d_275_v94_ = nw53_
                        d_278_v95_: _dafny.Seq
                        d_278_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbgug"))
                        d_279_v96_: int
                        d_280_v97_: bool
                        d_281_v98_: bool
                        d_282_v99_: bool
                        out19_: int
                        out20_: bool
                        out21_: bool
                        out22_: bool
                        out19_, out20_, out21_, out22_ = default__.m2(d_272_v91_, d_275_v94_, d_160_v3_, d_278_v95_, d_159_globalState_)
                        d_279_v96_ = out19_
                        d_280_v97_ = out20_
                        d_281_v98_ = out21_
                        d_282_v99_ = out22_
                        d_283_v100_: _dafny.MultiSet
                        d_283_v100_ = _dafny.MultiSet([default__.safeDivisionInt(d_160_v3_, d_160_v3_)])
                        d_283_v100_ = d_283_v100_
                    d_218_v44_ = d_218_v44_
                    pass
            pass
        d_284_v101_: C1
        nw54_ = C1()
        nw54_.ctor__()
        d_284_v101_ = nw54_
        d_285_v102_: _dafny.Seq
        d_285_v102_ = _dafny.SeqWithoutIsStrInference([d_284_v101_, d_284_v101_, d_284_v101_, d_284_v101_])
        d_284_v101_ = (d_285_v102_)[default__.safeIndex(((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]) * ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]), len(d_285_v102_))]
        d_286_v103_: _dafny.Map
        d_286_v103_ = _dafny.Map({d_160_v3_: d_158_v2_})
        d_287_v104_: _dafny.Seq
        d_287_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fb"))
        index28_ = default__.safeIndex(743, (d_175_v15_).length(0))
        (d_175_v15_)[index28_] = (len((d_286_v103_) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, False, d_157_v1_, False])): d_158_v2_}))) if d_157_v1_ else len(d_287_v104_))
        d_288_v105_: str
        d_288_v105_ = _dafny.CodePoint('d')
        d_289_v106_: _dafny.Set
        d_289_v106_ = _dafny.Set({(d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]})
        rhs36_ = (0) - (default__.safeDivisionInt(d_160_v3_, default__.fm1(-808, d_159_globalState_)))
        rhs37_ = (d_287_v104_).set(default__.safeIndex(d_160_v3_, len(d_287_v104_)), d_288_v105_)
        rhs38_ = (default__.fm21((d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_289_v106_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], d_159_globalState_)).cf45
        d_160_v3_ = rhs36_
        d_287_v104_ = rhs37_
        d_161_v4_ = rhs38_
        d_290_v107_: _dafny.MultiSet
        d_290_v107_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhuyl"))])
        d_291_v108_: _dafny.Seq
        d_291_v108_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_288_v105_ for d_292_i12_ in range(default__.abs(432))])])
        d_290_v107_ = ((_dafny.MultiSet([d_287_v104_, d_287_v104_])).intersection(d_290_v107_)) | (_dafny.MultiSet(d_291_v108_))
        d_293_v109_: int
        d_294_v110_: _dafny.Array
        d_295_v111_: D0
        out23_: int
        out24_: _dafny.Array
        out25_: D0
        out23_, out24_, out25_ = (d_284_v101_).m0(d_159_globalState_)
        d_293_v109_ = out23_
        d_294_v110_ = out24_
        d_295_v111_ = out25_
        d_296_v112_: int
        d_297_v113_: _dafny.Array
        d_298_v114_: D0
        out26_: int
        out27_: _dafny.Array
        out28_: D0
        out26_, out27_, out28_ = (d_284_v101_).m0(d_159_globalState_)
        d_296_v112_ = out26_
        d_297_v113_ = out27_
        d_298_v114_ = out28_
        if d_157_v1_:
            d_160_v3_ = default__.safeDivisionInt((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_160_v3_)
            d_284_v101_ = d_284_v101_
            d_299_v115_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Set({}), 28)
            d_299_v115_ = nw55_
            d_300_v116_: _dafny.Set
            d_300_v116_ = _dafny.Set({True})
            index29_ = default__.safeIndex(342, (d_299_v115_).length(0))
            (d_299_v115_)[index29_] = (d_300_v116_) - (d_300_v116_)
            index30_ = default__.safeIndex(342, (d_299_v115_).length(0))
            index31_ = default__.safeIndex(743, (d_175_v15_).length(0))
            rhs39_ = (default__.fm18(d_160_v3_, d_159_globalState_)).intersection(d_300_v116_)
            rhs40_ = default__.safeModuloInt((default__.fm1(default__.fm1((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))], d_159_globalState_), d_159_globalState_)) + ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]), (len(d_300_v116_)) - (len(d_173_v13_)))
            rhs41_ = default__.fm1(d_296_v112_, d_159_globalState_)
            lhs22_ = d_299_v115_
            lhs23_ = default__.safeIndex(342, (d_299_v115_).length(0))
            lhs24_ = d_175_v15_
            lhs25_ = default__.safeIndex(743, (d_175_v15_).length(0))
            lhs26_ = d_159_globalState_
            lhs22_[lhs23_] = rhs39_
            lhs24_[lhs25_] = rhs40_
            lhs26_.f3 = rhs41_
            d_301_v117_: _dafny.Array
            nw56_ = _dafny.Array(None, 17)
            nw56_[int(0)] = (d_173_v13_).set(d_157_v1_, len(_dafny.SeqWithoutIsStrInference([d_288_v105_ for d_302_i13_ in range(default__.abs(563))])))
            nw56_[int(1)] = d_173_v13_
            nw56_[int(2)] = default__.fm22(d_157_v1_, d_288_v105_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdnjcodfx")), not(not(d_157_v1_)), d_159_globalState_)
            nw56_[int(3)] = d_173_v13_
            nw56_[int(4)] = (d_173_v13_) | (d_173_v13_)
            nw56_[int(5)] = _dafny.Map({(d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]: (0) - (len(d_161_v4_))})
            nw56_[int(6)] = (_dafny.Map({d_157_v1_: (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]})) | (d_173_v13_)
            nw56_[int(7)] = d_173_v13_
            nw56_[int(8)] = d_173_v13_
            nw56_[int(9)] = (d_173_v13_ if (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] else d_173_v13_)
            nw56_[int(10)] = (d_173_v13_) | (d_173_v13_)
            nw56_[int(11)] = d_173_v13_
            nw56_[int(12)] = (d_173_v13_) | (d_173_v13_)
            nw56_[int(13)] = d_173_v13_
            nw56_[int(14)] = d_173_v13_
            nw56_[int(15)] = (d_173_v13_) | (d_173_v13_)
            nw56_[int(16)] = d_173_v13_
            d_301_v117_ = nw56_
            def iife13_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in (d_161_v4_).keys.Elements:
                    d_303_v118_: int = compr_9_
                    if (d_303_v118_) in (d_161_v4_):
                        coll9_[(d_303_v118_) * ((d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))])] = len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([d_160_v3_])).cardinality, d_296_v112_, d_293_v109_, d_160_v3_, (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]]))
                return _dafny.Map(coll9_)
            rhs42_ = ((iife13_()
            ) | (d_161_v4_)) | (d_161_v4_)
            rhs43_ = d_301_v117_
            d_161_v4_ = rhs42_
            d_301_v117_ = rhs43_
            d_160_v3_ = ((d_160_v3_) * (-212) if (d_293_v109_) <= (d_293_v109_) else d_296_v112_)
        elif True:
            d_304_v119_: D3
            d_304_v119_ = D3_DC11(174)
            d_305_v120_: D0
            d_305_v120_ = D0_DC0((d_304_v119_).cf14)
            d_306_v121_: _dafny.Seq
            d_306_v121_ = _dafny.SeqWithoutIsStrInference([d_296_v112_])
            d_307_v122_: _dafny.Seq
            d_307_v122_ = _dafny.SeqWithoutIsStrInference([d_306_v121_])
            d_308_v123_: D4
            d_308_v123_ = D4_DC14(d_296_v112_, d_305_v120_, (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))], d_157_v1_, (d_307_v122_)[default__.safeIndex(d_293_v109_, len(d_307_v122_))])
            d_309_v124_: D4
            d_309_v124_ = D4_DC15(D4_DC15(d_308_v123_))
            d_309_v124_ = d_309_v124_
            (d_159_globalState_).f1 = d_296_v112_
            d_310_v125_: _dafny.Seq
            d_310_v125_ = _dafny.SeqWithoutIsStrInference([d_175_v15_, (d_175_v15_ if True else d_175_v15_)])
            d_311_v126_: _dafny.MultiSet
            d_311_v126_ = _dafny.MultiSet([d_296_v112_])
            d_312_v127_: D6
            d_312_v127_ = D6_DC19(d_310_v125_)
            rhs44_ = d_284_v101_
            rhs45_ = default__.fm1(((d_311_v126_)[d_160_v3_] if (d_160_v3_) in (d_311_v126_) else d_296_v112_), d_159_globalState_)
            rhs46_ = ((d_310_v125_) + ((d_312_v127_).cf34)) + (_dafny.SeqWithoutIsStrInference([d_175_v15_, d_175_v15_]))
            rhs47_ = d_297_v113_
            d_284_v101_ = rhs44_
            d_293_v109_ = rhs45_
            d_310_v125_ = rhs46_
            d_297_v113_ = rhs47_
            d_313_v128_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
            d_313_v128_ = nw57_
            index32_ = default__.safeIndex(429, (d_313_v128_).length(0))
            (d_313_v128_)[index32_] = _dafny.SeqWithoutIsStrInference([(D0_DC1(d_288_v105_)).cf1 for d_314_i14_ in range(default__.abs(622))])
            index33_ = default__.safeIndex(429, (d_313_v128_).length(0))
            (d_313_v128_)[index33_] = ((d_287_v104_) + (d_287_v104_)) + (d_287_v104_)
            d_315_v129_: _dafny.Array
            nw58_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_315_v129_ = nw58_
            d_316_v130_: _dafny.Map
            d_316_v130_ = _dafny.Map({(d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]: d_315_v129_})
            rhs48_ = (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]
            rhs49_ = ((d_316_v130_)[((d_313_v128_)[default__.safeIndex(429, (d_313_v128_).length(0))]) <= ((d_313_v128_)[default__.safeIndex(429, (d_313_v128_).length(0))])] if (((d_313_v128_)[default__.safeIndex(429, (d_313_v128_).length(0))]) <= ((d_313_v128_)[default__.safeIndex(429, (d_313_v128_).length(0))])) in (d_316_v130_) else d_315_v129_)
            d_293_v109_ = rhs48_
            d_315_v129_ = rhs49_
        d_161_v4_ = (d_161_v4_).set(d_296_v112_, d_293_v109_)
        d_317_v131_: _dafny.Map
        d_317_v131_ = _dafny.Map({d_288_v105_: d_293_v109_})
        d_318_v132_: _dafny.Seq
        d_318_v132_ = _dafny.SeqWithoutIsStrInference([d_317_v131_])
        hi2_ = len((d_318_v132_)[default__.safeIndex(d_296_v112_, len(d_318_v132_))])
        for d_319_i15_ in range((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jywtmvy")))) - ((0) - (-98)), hi2_):
            d_320_v133_: _dafny.MultiSet
            d_320_v133_ = _dafny.MultiSet([len(d_289_v106_), len(d_287_v104_), d_319_i15_])
            index34_ = default__.safeIndex(949, (d_208_v41_).length(0))
            (d_208_v41_)[index34_] = (d_284_v101_).fm2(d_320_v133_, d_159_globalState_)
            d_321_v134_: _dafny.Array
            nw59_ = _dafny.Array(None, 5)
            nw59_[int(0)] = d_173_v13_
            nw59_[int(1)] = (d_173_v13_ if d_157_v1_ else _dafny.Map({d_157_v1_: (d_175_v15_)[default__.safeIndex(743, (d_175_v15_).length(0))]}))
            nw59_[int(2)] = _dafny.Map({(d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))]: d_293_v109_})
            nw59_[int(3)] = d_173_v13_
            nw59_[int(4)] = d_173_v13_
            d_321_v134_ = nw59_
            index35_ = default__.safeIndex(808, (d_321_v134_).length(0))
            (d_321_v134_)[index35_] = d_173_v13_
            index36_ = default__.safeIndex(808, (d_321_v134_).length(0))
            (d_321_v134_)[index36_] = d_173_v13_
            d_157_v1_ = not(not(d_157_v1_))
            d_322_v135_: _dafny.Array
            def lambda24_(d_323_v104_):
                def lambda25_(d_324_i16_):
                    return d_323_v104_

                return lambda25_

            init12_ = lambda24_(d_287_v104_)
            nw60_ = _dafny.Array(None, 24)
            for i0_12_ in range(nw60_.length(0)):
                nw60_[i0_12_] = init12_(i0_12_)
            d_322_v135_ = nw60_
            d_325_v136_: C0
            nw61_ = C0()
            nw61_.ctor__(d_322_v135_)
            d_325_v136_ = nw61_
            d_326_v137_: _dafny.Map
            d_326_v137_ = _dafny.Map({d_325_v136_: d_284_v101_})
            d_327_v138_: _dafny.Map
            d_327_v138_ = _dafny.Map({(d_325_v136_).f6: d_284_v101_})
            d_328_v139_: _dafny.Array
            nw62_ = _dafny.Array(None, 28)
            nw62_[int(0)] = d_284_v101_
            nw62_[int(1)] = d_284_v101_
            nw62_[int(2)] = d_284_v101_
            nw62_[int(3)] = d_284_v101_
            nw62_[int(4)] = d_284_v101_
            nw62_[int(5)] = (((d_326_v137_)[d_325_v136_] if (d_325_v136_) in (d_326_v137_) else d_284_v101_) if (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] else d_284_v101_)
            nw62_[int(6)] = ((d_327_v138_)[d_322_v135_] if (d_322_v135_) in (d_327_v138_) else d_284_v101_)
            nw62_[int(7)] = d_284_v101_
            nw62_[int(8)] = d_284_v101_
            nw62_[int(9)] = (d_285_v102_)[default__.safeIndex(442, len(d_285_v102_))]
            nw62_[int(10)] = d_284_v101_
            nw62_[int(11)] = d_284_v101_
            nw62_[int(12)] = d_284_v101_
            nw62_[int(13)] = d_284_v101_
            nw62_[int(14)] = d_284_v101_
            nw62_[int(15)] = d_284_v101_
            nw62_[int(16)] = d_284_v101_
            nw62_[int(17)] = d_284_v101_
            nw62_[int(18)] = d_284_v101_
            nw62_[int(19)] = d_284_v101_
            nw62_[int(20)] = d_284_v101_
            nw62_[int(21)] = d_284_v101_
            nw62_[int(22)] = d_284_v101_
            nw62_[int(23)] = d_284_v101_
            nw62_[int(24)] = d_284_v101_
            nw62_[int(25)] = (d_284_v101_ if (d_208_v41_)[default__.safeIndex(949, (d_208_v41_).length(0))] else d_284_v101_)
            nw62_[int(26)] = d_284_v101_
            nw62_[int(27)] = d_284_v101_
            d_328_v139_ = nw62_
            d_328_v139_ = d_328_v139_
        _dafny.print(_dafny.string_of(((d_155_v0_)[0]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[1]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[2]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[3]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[4]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[5]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[6]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[7]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[8]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[9]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[10]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[11]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_155_v0_)[12]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v2_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[0]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[1]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[2]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[3]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[4]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[5]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[6]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[7]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[8]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[9]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[10]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[11]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_.f4)[12]) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_159_globalState_).f5) == (_dafny.SeqWithoutIsStrInference([False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v4_) == (_dafny.Map({0: 5, 1: 2, 265: 982, 534: 534}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v13_) == (_dafny.Map({False: -272, True: 272}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v14_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v15_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_178_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v41_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v42_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_216_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_285_v102_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_286_v103_) == (_dafny.Map({272: _dafny.SeqWithoutIsStrInference([False, True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_287_v104_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_288_v105_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_289_v106_) == (_dafny.Set({2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_290_v107_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v108_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_293_v109_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v111_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_296_v112_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v114_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v131_) == (_dafny.Map({_dafny.CodePoint('d'): 534}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v132_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('d'): 534})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D1_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC5(D1, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC12(D3, NamedTuple('DC12', [('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14(int(0), D0.default()(), False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC14(D4, NamedTuple('DC14', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(int(0), False, False, False, int(0))
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

class D5_DC17(D5, NamedTuple('DC17', [('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({self.cf26.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC20(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)

class D6_DC20(D6, NamedTuple('DC20', [('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {self.cf39.VerbatimString(True)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC22(D6, NamedTuple('DC22', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)

class D7_DC24(D7, NamedTuple('DC24', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC26(False, _dafny.Array(None, 0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D8_DC27)

class D8_DC26(D8, NamedTuple('DC26', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC27(D8, NamedTuple('DC27', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC27({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC27) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC28(D9, NamedTuple('DC28', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f1: int = int(0)
        self.f3: int = int(0)
        self.f4: _dafny.Array = _dafny.Array(None, 0)
        self._f2: bool = False
        self._f5: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5

    @property
    def f2(self):
        return self._f2
    @property
    def f5(self):
        return self._f5

class C0:
    def  __init__(self):
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f6):
        (self)._f6 = f6

    def fm6(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uy"))

    def fm7(self, p0, globalState):
        return (725) * (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, False]): True})))

    @property
    def f6(self):
        return self._f6

class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm2(self, p0, globalState):
        def lambda26_(source3_):
            if source3_.is_DC1:
                d_329___mcc_h0_ = source3_.cf1
                d_330_cf1_ = d_329___mcc_h0_
                return False
            elif source3_.is_DC2:
                d_331___mcc_h1_ = source3_.cf2
                d_332_cf2_ = d_331___mcc_h1_
                return (_dafny.SeqWithoutIsStrInference([58, -449])) < (_dafny.SeqWithoutIsStrInference([3 for d_333_i0_ in range(default__.abs(773))]))
            elif source3_.is_DC3:
                d_334___mcc_h2_ = source3_.cf3
                d_335_cf3_ = d_334___mcc_h2_
                if d_335_cf3_:
                    return False
                elif True:
                    return d_335_cf3_
            elif True:
                d_336___mcc_h3_ = source3_.cf0
                d_337_cf0_ = d_336___mcc_h3_
                return True

        return not(lambda26_(D0_DC0((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hw")))))))

    def fm3(self, p0, globalState):
        return _dafny.CodePoint('c')

    def m0(self, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: D0 = D0.default()()
        d_338_v0_: _dafny.Array
        nw63_ = _dafny.Array(False, 29)
        d_338_v0_ = nw63_
        r1 = d_338_v0_
        r1 = d_338_v0_
        d_339_v1_: bool
        d_339_v1_ = False
        if d_339_v1_:
            d_340_v2_: _dafny.Array
            def lambda27_(d_341_i0_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pl"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lg")))

            init13_ = lambda27_
            nw64_ = _dafny.Array(None, 7)
            for i0_13_ in range(nw64_.length(0)):
                nw64_[i0_13_] = init13_(i0_13_)
            d_340_v2_ = nw64_
            d_342_v3_: int
            d_342_v3_ = 640
            d_343_v4_: _dafny.Array
            nw65_ = _dafny.Array(int(0), 9)
            d_343_v4_ = nw65_
            d_344_v5_: _dafny.Seq
            d_344_v5_ = _dafny.SeqWithoutIsStrInference([d_343_v4_, d_343_v4_, d_343_v4_])
            d_345_v6_: _dafny.Set
            d_345_v6_ = _dafny.Set({d_342_v3_, len(d_344_v5_)})
            d_346_v7_: str
            d_346_v7_ = _dafny.CodePoint('b')
            index37_ = default__.safeIndex(275, (d_340_v2_).length(0))
            (d_340_v2_)[index37_] = default__.fm4(d_339_v1_, d_342_v3_, len(_dafny.Map({d_345_v6_: d_339_v1_})), d_346_v7_, globalState)
            d_347_v8_: _dafny.Seq
            d_347_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eottg"))
            index38_ = default__.safeIndex(275, (d_340_v2_).length(0))
            (d_340_v2_)[index38_] = d_347_v8_
            index39_ = default__.safeIndex(33, (d_338_v0_).length(0))
            (d_338_v0_)[index39_] = d_339_v1_
            d_348_v9_: _dafny.Seq
            d_348_v9_ = _dafny.SeqWithoutIsStrInference([(d_347_v8_) + (d_347_v8_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srkfki")), (((d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]) + (d_347_v8_)).set(default__.safeIndex(len((d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]), len(((d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]) + (d_347_v8_))), d_346_v7_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_349_i1_ in range(default__.abs(-93))]), (d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]])
            d_350_v10_: D0
            d_350_v10_ = D0_DC1(d_346_v7_)
            index40_ = default__.safeIndex(33, (d_338_v0_).length(0))
            def iife14_(_pat_let2_0):
                def iife15_(d_351_dt__update__tmp_h0_):
                    def iife16_(_pat_let3_0):
                        def iife17_(d_352_dt__update_hcf1_h0_):
                            return D0_DC1(d_352_dt__update_hcf1_h0_)
                        return iife17_(_pat_let3_0)
                    return iife16_(_dafny.CodePoint('g'))
                return iife15_(_pat_let2_0)
            rhs50_ = d_339_v1_
            rhs51_ = ((_dafny.SeqWithoutIsStrInference([(d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))], default__.fm4(d_339_v1_, len(_dafny.Map({d_338_v0_: not(d_339_v1_)})), -828, d_346_v7_, globalState), (d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]])) + (((default__.fm5(d_342_v3_, d_342_v3_, globalState)).set(default__.safeIndex(d_342_v3_, len(default__.fm5(d_342_v3_, d_342_v3_, globalState))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkrtmgfeb")))).set(default__.safeIndex(d_342_v3_, len((default__.fm5(d_342_v3_, d_342_v3_, globalState)).set(default__.safeIndex(d_342_v3_, len(default__.fm5(d_342_v3_, d_342_v3_, globalState))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkrtmgfeb"))))), d_347_v8_))) + (d_348_v9_)
            rhs52_ = iife14_(d_350_v10_)
            lhs27_ = d_338_v0_
            lhs28_ = default__.safeIndex(33, (d_338_v0_).length(0))
            lhs27_[lhs28_] = rhs50_
            d_348_v9_ = rhs51_
            d_350_v10_ = rhs52_
            d_353_v11_: D0
            d_353_v11_ = D0_DC0(-402)
            d_354_v12_: _dafny.Seq
            d_354_v12_ = _dafny.SeqWithoutIsStrInference([d_353_v11_, d_353_v11_, d_353_v11_, D0_DC0(d_342_v3_)])
            (globalState).f3 = len(d_354_v12_)
            (globalState).f1 = default__.safeDivisionInt(272, d_342_v3_)
            d_355_v13_: _dafny.Map
            d_355_v13_ = _dafny.Map({d_346_v7_: (d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]})
            d_355_v13_ = (d_355_v13_).set(d_346_v7_, ((d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))]).set(default__.safeIndex(d_342_v3_, len((d_340_v2_)[default__.safeIndex(275, (d_340_v2_).length(0))])), d_346_v7_))
        elif True:
            d_356_v14_: int
            d_356_v14_ = 382
            (globalState).f1 = ((0) - (d_356_v14_)) * (d_356_v14_)
            d_357_v15_: _dafny.Seq
            d_357_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kp"))
            if not ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_358_i2_ in range(default__.abs(181))])) != (d_357_v15_)) or (d_339_v1_):
                (globalState).f1 = d_356_v14_
                d_359_v16_: str
                d_359_v16_ = _dafny.CodePoint('v')
                d_360_v17_: _dafny.Map
                d_360_v17_ = _dafny.Map({(D0_DC1(d_359_v16_)) in (_dafny.SeqWithoutIsStrInference([D0_DC1(d_359_v16_)])): d_356_v14_})
                d_361_v18_: _dafny.Map
                d_361_v18_ = _dafny.Map({d_339_v1_: _dafny.SeqWithoutIsStrInference([d_339_v1_])})
                d_362_v19_: _dafny.Seq
                d_362_v19_ = _dafny.SeqWithoutIsStrInference([d_339_v1_])
                d_360_v17_ = (d_360_v17_).set(d_339_v1_, (0) - (len((d_361_v18_).set(d_339_v1_, d_362_v19_))))
                d_363_v20_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_363_v20_ = nw66_
                d_364_v21_: C0
                nw67_ = C0()
                nw67_.ctor__(d_363_v20_)
                d_364_v21_ = nw67_
                d_365_v22_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_365_v22_ = nw68_
                index41_ = default__.safeIndex(685, (d_365_v22_).length(0))
                (d_365_v22_)[index41_] = (d_364_v21_).f6
                d_366_v23_: _dafny.Seq
                d_366_v23_ = _dafny.SeqWithoutIsStrInference([(d_364_v21_).f6])
                d_367_v24_: _dafny.Seq
                d_367_v24_ = _dafny.SeqWithoutIsStrInference([(d_364_v21_).f6, (d_364_v21_).f6, d_363_v20_, (d_366_v23_)[default__.safeIndex(len(d_362_v19_), len(d_366_v23_))], (d_364_v21_).f6])
                index42_ = default__.safeIndex(685, (d_365_v22_).length(0))
                (d_365_v22_)[index42_] = (d_366_v23_)[default__.safeIndex(d_356_v14_, len(d_366_v23_))]
                (globalState).f1 = d_356_v14_
            elif True:
                d_368_v25_: _dafny.Array
                nw69_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_368_v25_ = nw69_
                d_369_v26_: C0
                nw70_ = C0()
                nw70_.ctor__(d_368_v25_)
                d_369_v26_ = nw70_
                d_370_v27_: C0
                nw71_ = C0()
                nw71_.ctor__(d_368_v25_)
                d_370_v27_ = nw71_
                d_371_v28_: _dafny.Map
                d_371_v28_ = _dafny.Map({(default__.fm1(len(d_357_v15_), globalState)) >= (d_356_v14_): d_356_v14_})
                d_372_v29_: str
                d_372_v29_ = _dafny.CodePoint('o')
                d_371_v28_ = (d_371_v28_).set(not(not((default__.fm8(d_339_v1_, d_339_v1_, len(d_357_v15_), globalState)).ispropersubset(_dafny.Set({d_372_v29_})))), d_356_v14_)
                (globalState).f0 = d_339_v1_
                r0 = d_356_v14_
            (globalState).f3 = 240
            d_373_v30_: str
            d_373_v30_ = _dafny.CodePoint('n')
            d_374_v31_: _dafny.Array
            nw72_ = _dafny.Array(None, 8)
            nw72_[int(0)] = default__.fm4(d_339_v1_, d_356_v14_, d_356_v14_, d_373_v30_, globalState)
            nw72_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcokmtins"))
            nw72_[int(2)] = d_357_v15_
            nw72_[int(3)] = (d_357_v15_).set(default__.safeIndex(len(d_357_v15_), len(d_357_v15_)), d_373_v30_)
            nw72_[int(4)] = d_357_v15_
            nw72_[int(5)] = d_357_v15_
            nw72_[int(6)] = d_357_v15_
            nw72_[int(7)] = d_357_v15_
            d_374_v31_ = nw72_
            d_375_v32_: C0
            nw73_ = C0()
            nw73_.ctor__(d_374_v31_)
            d_375_v32_ = nw73_
            d_376_v33_: _dafny.Array
            nw74_ = _dafny.Array(None, 15)
            nw74_[int(0)] = d_375_v32_
            nw74_[int(1)] = d_375_v32_
            nw74_[int(2)] = d_375_v32_
            nw74_[int(3)] = d_375_v32_
            nw74_[int(4)] = d_375_v32_
            nw74_[int(5)] = d_375_v32_
            nw74_[int(6)] = d_375_v32_
            nw74_[int(7)] = d_375_v32_
            nw74_[int(8)] = d_375_v32_
            nw74_[int(9)] = d_375_v32_
            nw74_[int(10)] = d_375_v32_
            nw74_[int(11)] = d_375_v32_
            nw74_[int(12)] = d_375_v32_
            nw74_[int(13)] = d_375_v32_
            nw74_[int(14)] = d_375_v32_
            d_376_v33_ = nw74_
            d_377_v34_: _dafny.Set
            d_377_v34_ = _dafny.Set({d_376_v33_})
            d_377_v34_ = (d_377_v34_).intersection(d_377_v34_)
            d_378_v35_: _dafny.Array
            nw75_ = _dafny.Array(None, 1)
            nw75_[int(0)] = d_356_v14_
            d_378_v35_ = nw75_
            d_379_v36_: _dafny.Set
            d_379_v36_ = _dafny.Set({d_356_v14_})
            index43_ = default__.safeIndex(72, (d_378_v35_).length(0))
            (d_378_v35_)[index43_] = len((d_379_v36_) - (d_379_v36_))
            index44_ = default__.safeIndex(72, (d_378_v35_).length(0))
            (d_378_v35_)[index44_] = d_356_v14_
        d_380_v37_: int
        d_380_v37_ = 534
        d_381_v38_: _dafny.MultiSet
        d_381_v38_ = _dafny.MultiSet([d_339_v1_])
        d_382_v39_: _dafny.Set
        d_382_v39_ = _dafny.Set({D0_DC2(d_339_v1_), D0_DC2(d_339_v1_)})
        d_383_v40_: _dafny.Seq
        d_383_v40_ = _dafny.SeqWithoutIsStrInference([d_382_v39_, d_382_v39_])
        d_384_v41_: _dafny.Map
        d_384_v41_ = _dafny.Map({d_339_v1_: (0) - (d_380_v37_)})
        d_385_v42_: _dafny.Seq
        d_385_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqusjevmf"))
        d_386_v43_: _dafny.Array
        nw76_ = _dafny.Array(None, 20)
        nw76_[int(0)] = 195
        nw76_[int(1)] = d_380_v37_
        nw76_[int(2)] = d_380_v37_
        nw76_[int(3)] = d_380_v37_
        nw76_[int(4)] = default__.fm1(d_380_v37_, globalState)
        nw76_[int(5)] = d_380_v37_
        nw76_[int(6)] = len(_dafny.Map({d_381_v38_: d_380_v37_}))
        nw76_[int(7)] = d_380_v37_
        nw76_[int(8)] = d_380_v37_
        nw76_[int(9)] = d_380_v37_
        nw76_[int(10)] = len(d_383_v40_)
        nw76_[int(11)] = (d_380_v37_) + (d_380_v37_)
        nw76_[int(12)] = d_380_v37_
        nw76_[int(13)] = d_380_v37_
        nw76_[int(14)] = d_380_v37_
        nw76_[int(15)] = default__.safeDivisionInt(((d_384_v41_)[d_339_v1_] if (d_339_v1_) in (d_384_v41_) else d_380_v37_), len(d_385_v42_))
        nw76_[int(16)] = d_380_v37_
        nw76_[int(17)] = d_380_v37_
        nw76_[int(18)] = d_380_v37_
        nw76_[int(19)] = d_380_v37_
        d_386_v43_ = nw76_
        d_386_v43_ = d_386_v43_
        d_387_v44_: _dafny.Array
        def lambda28_(d_388_v1_):
            def lambda29_(d_389_i3_):
                return D0_DC3(d_388_v1_)

            return lambda29_

        init14_ = lambda28_(d_339_v1_)
        nw77_ = _dafny.Array(None, 22)
        for i0_14_ in range(nw77_.length(0)):
            nw77_[i0_14_] = init14_(i0_14_)
        d_387_v44_ = nw77_
        index45_ = default__.safeIndex(692, (d_387_v44_).length(0))
        (d_387_v44_)[index45_] = default__.fm9((_dafny.MultiSet([d_380_v37_, d_380_v37_, len(d_385_v42_), d_380_v37_, d_380_v37_])).cardinality, globalState)
        d_390_v45_: _dafny.Map
        d_390_v45_ = _dafny.Map({d_380_v37_: d_339_v1_})
        index46_ = default__.safeIndex(692, (d_387_v44_).length(0))
        rhs53_ = d_339_v1_
        rhs54_ = D0_DC3((d_339_v1_) or (d_339_v1_))
        rhs55_ = (0) - (d_380_v37_)
        rhs56_ = (0) - (len((_dafny.Map({342: d_339_v1_})) | (d_390_v45_)))
        lhs29_ = globalState
        lhs30_ = d_387_v44_
        lhs31_ = default__.safeIndex(692, (d_387_v44_).length(0))
        lhs32_ = globalState
        lhs33_ = globalState
        lhs29_.f0 = rhs53_
        lhs30_[lhs31_] = rhs54_
        lhs32_.f1 = rhs55_
        lhs33_.f1 = rhs56_
        d_380_v37_ = d_380_v37_
        r0 = d_380_v37_
        r1 = d_338_v0_
        r2 = (d_387_v44_)[default__.safeIndex(692, (d_387_v44_).length(0))]
        return r0, r1, r2

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_391_v0_: _dafny.Array
        nw78_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_391_v0_ = nw78_
        d_392_v1_: C0
        nw79_ = C0()
        nw79_.ctor__(d_391_v0_)
        d_392_v1_ = nw79_
        d_393_v2_: int
        d_393_v2_ = 652
        d_394_v3_: _dafny.Map
        d_394_v3_ = _dafny.Map({d_392_v1_: d_393_v2_})
        d_395_v4_: _dafny.Seq
        d_395_v4_ = _dafny.SeqWithoutIsStrInference([d_392_v1_])
        d_394_v3_ = (d_394_v3_).set(d_392_v1_, (824 if p0 else len(d_395_v4_)))
        d_396_i0_: int
        d_396_i0_ = 0
        with _dafny.label("2"):
            while p0:
                with _dafny.c_label("2"):
                    if (d_396_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_396_i0_ = (d_396_i0_) + (1)
                    d_397_v5_: _dafny.Map
                    d_397_v5_ = _dafny.Map({(0) - (d_393_v2_): D0_DC0(d_393_v2_)})
                    d_398_v6_: _dafny.Seq
                    d_398_v6_ = _dafny.SeqWithoutIsStrInference([d_397_v5_])
                    d_399_v7_: _dafny.Seq
                    d_399_v7_ = _dafny.SeqWithoutIsStrInference([d_393_v2_, d_393_v2_])
                    d_400_v8_: _dafny.MultiSet
                    d_400_v8_ = _dafny.MultiSet([(d_398_v6_)[default__.safeIndex(len(_dafny.Map({(d_399_v7_)[default__.safeIndex(d_393_v2_, len(d_399_v7_))]: d_393_v2_})), len(d_398_v6_))]])
                    d_400_v8_ = d_400_v8_
                    (globalState).f1 = (d_393_v2_) * (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([d_393_v2_ for d_401_i1_ in range(default__.abs(-537))]))), d_393_v2_))
                    d_402_v9_: C0
                    nw80_ = C0()
                    nw80_.ctor__(d_391_v0_)
                    d_402_v9_ = nw80_
                    d_403_v10_: _dafny.Array
                    nw81_ = _dafny.Array(int(0), 25)
                    d_403_v10_ = nw81_
                    d_404_v11_: _dafny.Seq
                    d_404_v11_ = _dafny.SeqWithoutIsStrInference([p0, p0, False, p0])
                    index47_ = default__.safeIndex(686, (d_403_v10_).length(0))
                    (d_403_v10_)[index47_] = len(d_404_v11_)
                    index48_ = default__.safeIndex(686, (d_403_v10_).length(0))
                    (d_403_v10_)[index48_] = len(d_404_v11_)
                    pass
            pass
        d_405_v12_: _dafny.Array
        def lambda30_(d_406_p0_):
            def lambda31_(d_407_i2_):
                return d_406_p0_

            return lambda31_

        init15_ = lambda30_(p0)
        nw82_ = _dafny.Array(None, 29)
        for i0_15_ in range(nw82_.length(0)):
            nw82_[i0_15_] = init15_(i0_15_)
        d_405_v12_ = nw82_
        d_408_v13_: _dafny.Map
        d_408_v13_ = _dafny.Map({len(default__.fm10(globalState)): d_405_v12_})
        d_408_v13_ = (d_408_v13_).set(d_393_v2_, (d_405_v12_ if p0 else d_405_v12_))
        d_409_v14_: str
        d_409_v14_ = _dafny.CodePoint('x')
        d_410_v15_: D0
        d_410_v15_ = D0_DC1(d_409_v14_)
        d_411_v16_: _dafny.Seq
        d_411_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm1((_dafny.MultiSet([D0_DC1(d_409_v14_), d_410_v15_])).cardinality, globalState)])
        if ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbijgu")))) * (len(p2))) in (d_411_v16_):
            d_412_v17_: _dafny.Seq
            d_412_v17_ = _dafny.SeqWithoutIsStrInference([p0, p0, False])
            d_412_v17_ = (d_412_v17_) + ((d_412_v17_ if not(p0) else d_412_v17_))
            d_413_v18_: D0
            d_413_v18_ = D0_DC2(p0)
            (globalState).f0 = (d_413_v18_).cf2
            source4_ = D0_DC3(default__.fm0(len(d_411_v16_), globalState))
            if source4_.is_DC1:
                d_414___mcc_h0_ = source4_.cf1
                d_415_cf1_ = d_414___mcc_h0_
                d_416_v19_: _dafny.Array
                def lambda32_(d_417_v2_):
                    def lambda33_(d_418_i3_):
                        return (d_418_i3_) + (d_417_v2_)

                    return lambda33_

                init16_ = lambda32_(d_393_v2_)
                nw83_ = _dafny.Array(None, 5)
                for i0_16_ in range(nw83_.length(0)):
                    nw83_[i0_16_] = init16_(i0_16_)
                d_416_v19_ = nw83_
                index49_ = default__.safeIndex(656, (d_416_v19_).length(0))
                (d_416_v19_)[index49_] = (0) - ((d_393_v2_) - (-624))
                d_419_v20_: _dafny.Seq
                d_419_v20_ = _dafny.SeqWithoutIsStrInference([d_416_v19_])
                index50_ = default__.safeIndex(656, (d_416_v19_).length(0))
                (d_416_v19_)[index50_] = len((_dafny.SeqWithoutIsStrInference([d_416_v19_])) + ((d_419_v20_) + (d_419_v20_)))
                d_420_v21_: _dafny.Map
                d_420_v21_ = _dafny.Map({d_394_v3_: p0})
                d_420_v21_ = _dafny.Map({d_394_v3_: p0})
                r0 = default__.safeModuloInt((d_416_v19_)[default__.safeIndex(656, (d_416_v19_).length(0))], d_393_v2_)
                index51_ = default__.safeIndex(265, (d_405_v12_).length(0))
                (d_405_v12_)[index51_] = p0
                index52_ = default__.safeIndex(265, (d_405_v12_).length(0))
                rhs57_ = (self).fm3(((d_416_v19_)[default__.safeIndex(656, (d_416_v19_).length(0))]) != ((d_416_v19_)[default__.safeIndex(656, (d_416_v19_).length(0))]), globalState)
                rhs58_ = p0
                rhs59_ = p0
                lhs34_ = globalState
                lhs35_ = d_405_v12_
                lhs36_ = default__.safeIndex(265, (d_405_v12_).length(0))
                d_415_cf1_ = rhs57_
                lhs34_.f0 = rhs58_
                lhs35_[lhs36_] = rhs59_
            elif source4_.is_DC2:
                d_421___mcc_h1_ = source4_.cf2
                d_422_cf2_ = d_421___mcc_h1_
                d_423_v22_: _dafny.Seq
                d_423_v22_ = _dafny.SeqWithoutIsStrInference([default__.fm11(p0, p0, D0_DC3(p0), globalState), d_412_v17_, d_412_v17_])
                d_424_v23_: _dafny.MultiSet
                d_424_v23_ = _dafny.MultiSet([len(d_423_v22_)])
                d_425_v24_: _dafny.Map
                d_425_v24_ = _dafny.Map({d_393_v2_: (d_424_v23_).cardinality})
                d_425_v24_ = (d_425_v24_).set(((0) - (d_393_v2_)) + (d_393_v2_), default__.fm1(d_393_v2_, globalState))
                d_426_v25_: C0
                nw84_ = C0()
                nw84_.ctor__((d_392_v1_).f6)
                d_426_v25_ = nw84_
                d_409_v14_ = _dafny.CodePoint('v')
                rhs60_ = d_393_v2_
                rhs61_ = (d_393_v2_) - ((d_393_v2_) * (d_393_v2_))
                lhs37_ = globalState
                lhs38_ = globalState
                lhs37_.f3 = rhs60_
                lhs38_.f1 = rhs61_
            elif source4_.is_DC3:
                d_427___mcc_h2_ = source4_.cf3
                d_428_cf3_ = d_427___mcc_h2_
                (globalState).f0 = p0
                d_429_v26_: _dafny.Set
                d_429_v26_ = _dafny.Set({d_392_v1_})
                d_430_v27_: _dafny.Set
                d_430_v27_ = _dafny.Set({d_429_v26_})
                d_431_v28_: _dafny.MultiSet
                d_431_v28_ = _dafny.MultiSet([d_428_cf3_])
                d_432_v29_: _dafny.Map
                d_432_v29_ = _dafny.Map({d_392_v1_: d_431_v28_})
                d_433_v30_: D1
                d_433_v30_ = D1_DC4(d_431_v28_)
                d_434_v31_: D2
                d_434_v31_ = D2_DC7(d_392_v1_)
                d_435_v32_: _dafny.Map
                d_435_v32_ = _dafny.Map({d_409_v14_: (d_434_v31_).cf7})
                d_436_v33_: _dafny.Array
                nw85_ = _dafny.Array(None, 27)
                nw85_[int(0)] = d_393_v2_
                nw85_[int(1)] = (d_393_v2_) + ((D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "reo"))))).cf0)
                nw85_[int(2)] = d_393_v2_
                nw85_[int(3)] = d_393_v2_
                nw85_[int(4)] = d_393_v2_
                nw85_[int(5)] = d_393_v2_
                nw85_[int(6)] = d_393_v2_
                nw85_[int(7)] = len((d_430_v27_).intersection(d_430_v27_))
                nw85_[int(8)] = d_393_v2_
                nw85_[int(9)] = d_393_v2_
                nw85_[int(10)] = default__.safeDivisionInt((0) - (d_393_v2_), d_393_v2_)
                nw85_[int(11)] = default__.safeDivisionInt(len(d_411_v16_), d_393_v2_)
                nw85_[int(12)] = ((d_411_v16_)[default__.safeIndex(d_393_v2_, len(d_411_v16_))]) - (d_393_v2_)
                nw85_[int(13)] = default__.safeModuloInt(d_393_v2_, ((d_431_v28_)[d_428_cf3_] if (d_428_cf3_) in (d_431_v28_) else d_393_v2_))
                nw85_[int(14)] = d_393_v2_
                nw85_[int(15)] = d_393_v2_
                nw85_[int(16)] = len(_dafny.Map({((d_432_v29_)[d_392_v1_] if (d_392_v1_) in (d_432_v29_) else (d_433_v30_).cf4): d_405_v12_}))
                nw85_[int(17)] = default__.fm1(d_393_v2_, globalState)
                nw85_[int(18)] = default__.safeDivisionInt(d_393_v2_, d_393_v2_)
                nw85_[int(19)] = d_393_v2_
                nw85_[int(20)] = d_393_v2_
                nw85_[int(21)] = len((d_435_v32_) | (d_435_v32_))
                nw85_[int(22)] = d_393_v2_
                nw85_[int(23)] = d_393_v2_
                nw85_[int(24)] = d_393_v2_
                nw85_[int(25)] = d_393_v2_
                nw85_[int(26)] = 814
                d_436_v33_ = nw85_
                index53_ = default__.safeIndex(294, (d_436_v33_).length(0))
                (d_436_v33_)[index53_] = d_393_v2_
                index54_ = default__.safeIndex(294, (d_436_v33_).length(0))
                (d_436_v33_)[index54_] = d_393_v2_
                d_437_v34_: _dafny.Seq
                d_437_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkcpig"))
                d_437_v34_ = (d_437_v34_ if (d_393_v2_) <= ((d_436_v33_)[default__.safeIndex(294, (d_436_v33_).length(0))]) else d_437_v34_)
                d_438_v35_: _dafny.MultiSet
                d_438_v35_ = _dafny.MultiSet([d_411_v16_, (default__.fm10(globalState)).set(default__.safeIndex((d_436_v33_)[default__.safeIndex(294, (d_436_v33_).length(0))], len(default__.fm10(globalState))), d_393_v2_), d_411_v16_])
                d_438_v35_ = ((d_438_v35_) - (d_438_v35_)).intersection((d_438_v35_ if p0 else _dafny.MultiSet([(d_411_v16_).set(default__.safeIndex(len(p1), len(d_411_v16_)), len(_dafny.SeqWithoutIsStrInference([d_428_cf3_]))), _dafny.SeqWithoutIsStrInference([(d_436_v33_)[default__.safeIndex(294, (d_436_v33_).length(0))], d_393_v2_]), d_411_v16_])))
            elif True:
                d_439___mcc_h3_ = source4_.cf0
                d_440_cf0_ = d_439___mcc_h3_
                d_441_v36_: _dafny.Set
                d_441_v36_ = _dafny.Set({d_440_cf0_, d_393_v2_})
                d_442_v37_: _dafny.Map
                d_442_v37_ = _dafny.Map({d_440_cf0_: len(d_441_v36_)})
                d_443_v38_: _dafny.Seq
                d_443_v38_ = _dafny.SeqWithoutIsStrInference([d_442_v37_])
                rhs62_ = default__.safeModuloInt(903, d_393_v2_)
                rhs63_ = (_dafny.SeqWithoutIsStrInference([d_392_v1_])) + (d_395_v4_)
                rhs64_ = (D3_DC9(d_443_v38_)).cf10
                rhs65_ = d_405_v12_
                lhs39_ = globalState
                lhs39_.f1 = rhs62_
                d_395_v4_ = rhs63_
                d_443_v38_ = rhs64_
                d_405_v12_ = rhs65_
                d_444_v39_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_444_v39_ = nw86_
                index55_ = default__.safeIndex(790, (d_444_v39_).length(0))
                (d_444_v39_)[index55_] = d_409_v14_
                d_445_v40_: _dafny.Map
                d_445_v40_ = _dafny.Map({d_392_v1_: d_409_v14_})
                d_446_v41_: D2
                d_446_v41_ = D2_DC7(d_392_v1_)
                index56_ = default__.safeIndex(790, (d_444_v39_).length(0))
                (d_444_v39_)[index56_] = ((d_445_v40_)[(d_446_v41_).cf7] if ((d_446_v41_).cf7) in (d_445_v40_) else d_409_v14_)
                (globalState).f1 = d_393_v2_
                d_447_v42_: _dafny.MultiSet
                d_447_v42_ = _dafny.MultiSet([p0, p0, not(p0)])
                d_448_v43_: D1
                d_448_v43_ = D1_DC4(_dafny.MultiSet(d_412_v17_))
                d_449_v44_: _dafny.Array
                nw87_ = _dafny.Array(None, 22)
                nw87_[int(0)] = (d_447_v42_) - (_dafny.MultiSet([False]))
                nw87_[int(1)] = (d_447_v42_) | (_dafny.MultiSet(d_412_v17_))
                nw87_[int(2)] = (d_447_v42_ if p0 else d_447_v42_)
                nw87_[int(3)] = _dafny.MultiSet([p0, p0, default__.fm0(d_393_v2_, globalState)])
                nw87_[int(4)] = d_447_v42_
                nw87_[int(5)] = (d_448_v43_).cf4
                nw87_[int(6)] = d_447_v42_
                nw87_[int(7)] = _dafny.MultiSet([p0])
                nw87_[int(8)] = d_447_v42_
                nw87_[int(9)] = (_dafny.MultiSet(d_412_v17_)).intersection(d_447_v42_)
                nw87_[int(10)] = d_447_v42_
                nw87_[int(11)] = d_447_v42_
                nw87_[int(12)] = _dafny.MultiSet([p0, p0])
                nw87_[int(13)] = d_447_v42_
                nw87_[int(14)] = (d_447_v42_).intersection(d_447_v42_)
                nw87_[int(15)] = (d_447_v42_) | (d_447_v42_)
                nw87_[int(16)] = (d_447_v42_) - (_dafny.MultiSet([True]))
                nw87_[int(17)] = default__.fm12(d_393_v2_, globalState)
                nw87_[int(18)] = _dafny.MultiSet([True])
                nw87_[int(19)] = d_447_v42_
                nw87_[int(20)] = (d_447_v42_) - (d_447_v42_)
                nw87_[int(21)] = (_dafny.MultiSet([False, not(p0)])) | (d_447_v42_)
                d_449_v44_ = nw87_
                d_450_v45_: _dafny.Map
                d_450_v45_ = _dafny.Map({d_392_v1_: d_447_v42_})
                index57_ = default__.safeIndex(99, (d_449_v44_).length(0))
                (d_449_v44_)[index57_] = ((d_450_v45_)[d_392_v1_] if (d_392_v1_) in (d_450_v45_) else d_447_v42_)
                d_451_v46_: D0
                d_451_v46_ = D0_DC0(d_440_cf0_)
                pat_let_tv9_ = d_440_cf0_
                d_452_v47_: _dafny.Array
                nw88_ = _dafny.Array(None, 11)
                nw88_[int(0)] = D0_DC0(d_393_v2_)
                nw88_[int(1)] = d_451_v46_
                def iife18_(_pat_let4_0):
                    def iife19_(d_453_dt__update__tmp_h0_):
                        def iife20_(_pat_let5_0):
                            def iife21_(d_454_dt__update_hcf0_h0_):
                                return D0_DC0(d_454_dt__update_hcf0_h0_)
                            return iife21_(_pat_let5_0)
                        return iife20_(pat_let_tv9_)
                    return iife19_(_pat_let4_0)
                nw88_[int(2)] = iife18_(d_451_v46_)
                nw88_[int(3)] = d_451_v46_
                nw88_[int(4)] = default__.fm13(False, p0, _dafny.Map({d_393_v2_: d_440_cf0_}), p0, globalState)
                nw88_[int(5)] = d_451_v46_
                nw88_[int(6)] = d_451_v46_
                nw88_[int(7)] = d_451_v46_
                nw88_[int(8)] = d_451_v46_
                nw88_[int(9)] = d_451_v46_
                nw88_[int(10)] = D0_DC0(d_393_v2_)
                d_452_v47_ = nw88_
                index58_ = default__.safeIndex(130, (d_452_v47_).length(0))
                (d_452_v47_)[index58_] = d_451_v46_
                d_455_v48_: D0
                d_455_v48_ = D0_DC3((d_393_v2_) > (d_440_cf0_))
                d_456_v49_: _dafny.Map
                d_456_v49_ = _dafny.Map({d_393_v2_: _dafny.MultiSet([not(p0), p0])})
                index59_ = default__.safeIndex(99, (d_449_v44_).length(0))
                index60_ = default__.safeIndex(130, (d_452_v47_).length(0))
                rhs66_ = ((d_456_v49_)[d_393_v2_] if (d_393_v2_) in (d_456_v49_) else (d_447_v42_) - (d_447_v42_))
                rhs67_ = d_440_cf0_
                rhs68_ = d_451_v46_
                rhs69_ = D0_DC3(p0)
                rhs70_ = default__.fm1(d_393_v2_, globalState)
                lhs40_ = d_449_v44_
                lhs41_ = default__.safeIndex(99, (d_449_v44_).length(0))
                lhs42_ = globalState
                lhs43_ = d_452_v47_
                lhs44_ = default__.safeIndex(130, (d_452_v47_).length(0))
                lhs40_[lhs41_] = rhs66_
                lhs42_.f3 = rhs67_
                lhs43_[lhs44_] = rhs68_
                d_455_v48_ = rhs69_
                r0 = rhs70_
            if (len(d_412_v17_)) >= (d_393_v2_):
                index61_ = default__.safeIndex(788, (d_405_v12_).length(0))
                (d_405_v12_)[index61_] = p0
                index62_ = default__.safeIndex(788, (d_405_v12_).length(0))
                (d_405_v12_)[index62_] = p0
                d_457_v50_: _dafny.Array
                def lambda34_(d_458_v2_):
                    def lambda35_(d_459_i4_):
                        return default__.safeDivisionInt(d_459_i4_, d_458_v2_)

                    return lambda35_

                init17_ = lambda34_(d_393_v2_)
                nw89_ = _dafny.Array(None, 3)
                for i0_17_ in range(nw89_.length(0)):
                    nw89_[i0_17_] = init17_(i0_17_)
                d_457_v50_ = nw89_
                index63_ = default__.safeIndex(199, (d_457_v50_).length(0))
                (d_457_v50_)[index63_] = d_393_v2_
                index64_ = default__.safeIndex(199, (d_457_v50_).length(0))
                (d_457_v50_)[index64_] = (d_393_v2_) - (d_393_v2_)
                (globalState).f0 = not(True)
                d_391_v0_ = (d_392_v1_).f6
                (globalState).f0 = p0
            elif True:
                (globalState).f0 = p0
                (globalState).f0 = ((d_412_v17_) + (d_412_v17_)) == (d_412_v17_)
                index65_ = default__.safeIndex(188, (d_405_v12_).length(0))
                (d_405_v12_)[index65_] = p0
                d_460_v51_: D4
                d_460_v51_ = D4_DC13(d_405_v12_)
                pat_let_tv10_ = d_460_v51_
                pat_let_tv11_ = d_405_v12_
                index66_ = default__.safeIndex(188, (d_405_v12_).length(0))
                def iife22_(_pat_let6_0):
                    def iife23_(d_461_dt__update__tmp_h1_):
                        def iife25_(_pat_let8_0):
                            def iife26_(d_462_dt__update__tmp_h2_):
                                def iife27_(_pat_let9_0):
                                    def iife28_(d_463_dt__update_hcf19_h1_):
                                        return D4_DC13(d_463_dt__update_hcf19_h1_)
                                    return iife28_(_pat_let9_0)
                                return iife27_(pat_let_tv11_)
                            return iife26_(_pat_let8_0)
                        def iife24_(_pat_let7_0):
                            def iife29_(d_464_dt__update_hcf19_h0_):
                                return D4_DC13(d_464_dt__update_hcf19_h0_)
                            return iife29_(_pat_let7_0)
                        return iife24_((iife25_(pat_let_tv10_)).cf19)
                    return iife23_(_pat_let6_0)
                rhs71_ = (iife22_(d_460_v51_)).cf19
                rhs72_ = (d_393_v2_) > (d_393_v2_)
                rhs73_ = p0
                lhs45_ = d_405_v12_
                lhs46_ = default__.safeIndex(188, (d_405_v12_).length(0))
                lhs47_ = globalState
                d_405_v12_ = rhs71_
                lhs45_[lhs46_] = rhs72_
                lhs47_.f0 = rhs73_
                index67_ = default__.safeIndex(188, (d_405_v12_).length(0))
                rhs74_ = d_391_v0_
                rhs75_ = (d_405_v12_)[default__.safeIndex(188, (d_405_v12_).length(0))]
                lhs48_ = d_405_v12_
                lhs49_ = default__.safeIndex(188, (d_405_v12_).length(0))
                d_391_v0_ = rhs74_
                lhs48_[lhs49_] = rhs75_
                (globalState).f0 = p0
            d_465_v52_: _dafny.Seq
            d_465_v52_ = _dafny.SeqWithoutIsStrInference([d_409_v14_])
            d_465_v52_ = _dafny.SeqWithoutIsStrInference([d_409_v14_])
        elif True:
            (globalState).f1 = ((0) - (d_393_v2_)) - (d_393_v2_)
            (globalState).f0 = p0
            d_466_v53_: _dafny.Map
            d_466_v53_ = _dafny.Map({d_393_v2_: 547})
            r0 = ((d_466_v53_)[d_393_v2_] if (d_393_v2_) in (d_466_v53_) else d_393_v2_)
            d_467_v54_: _dafny.Seq
            d_467_v54_ = _dafny.SeqWithoutIsStrInference([p0])
            d_468_v55_: _dafny.Map
            d_468_v55_ = _dafny.Map({d_467_v54_: default__.fm4(p0, d_393_v2_, d_393_v2_, d_409_v14_, globalState)})
            d_468_v55_ = (d_468_v55_).set(default__.fm11(p0, p0, p3, globalState), (p2) + (p2))
            index68_ = default__.safeIndex(154, (d_405_v12_).length(0))
            (d_405_v12_)[index68_] = True
            index69_ = default__.safeIndex(154, (d_405_v12_).length(0))
            (d_405_v12_)[index69_] = not(p0)
        d_469_v56_: D0
        d_469_v56_ = D0_DC2(p0)
        d_470_v57_: _dafny.Seq
        d_470_v57_ = _dafny.SeqWithoutIsStrInference([d_469_v56_, d_469_v56_])
        d_471_v58_: _dafny.Array
        nw90_ = _dafny.Array(None, 10)
        nw90_[int(0)] = d_470_v57_
        nw90_[int(1)] = (default__.fm14(p0, globalState)) + (_dafny.SeqWithoutIsStrInference([d_469_v56_]))
        nw90_[int(2)] = d_470_v57_
        nw90_[int(3)] = (d_470_v57_) + (d_470_v57_)
        nw90_[int(4)] = default__.fm14(p0, globalState)
        nw90_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_469_v56_ for d_472_i5_ in range(default__.abs(69))])) + (d_470_v57_)
        nw90_[int(6)] = d_470_v57_
        nw90_[int(7)] = (d_470_v57_) + (d_470_v57_)
        nw90_[int(8)] = _dafny.SeqWithoutIsStrInference([d_469_v56_, d_469_v56_])
        nw90_[int(9)] = (d_470_v57_) + (d_470_v57_)
        d_471_v58_ = nw90_
        index70_ = default__.safeIndex(127, (d_471_v58_).length(0))
        (d_471_v58_)[index70_] = d_470_v57_
        index71_ = default__.safeIndex(231, (d_405_v12_).length(0))
        (d_405_v12_)[index71_] = (len(default__.fm4(p0, d_393_v2_, d_393_v2_, d_409_v14_, globalState))) <= ((0) - (d_393_v2_))
        d_473_v59_: _dafny.Map
        d_473_v59_ = _dafny.Map({p0: d_470_v57_})
        index72_ = default__.safeIndex(127, (d_471_v58_).length(0))
        index73_ = default__.safeIndex(231, (d_405_v12_).length(0))
        rhs76_ = default__.safeModuloInt(d_393_v2_, d_393_v2_)
        rhs77_ = ((d_473_v59_)[not(p0)] if (not(p0)) in (d_473_v59_) else (_dafny.SeqWithoutIsStrInference([d_469_v56_ for d_474_i6_ in range(default__.abs(50))])) + (_dafny.SeqWithoutIsStrInference([d_469_v56_])))
        rhs78_ = p0
        rhs79_ = (p0) or ((p2) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))
        lhs50_ = globalState
        lhs51_ = d_471_v58_
        lhs52_ = default__.safeIndex(127, (d_471_v58_).length(0))
        lhs53_ = d_405_v12_
        lhs54_ = default__.safeIndex(231, (d_405_v12_).length(0))
        lhs55_ = globalState
        lhs50_.f3 = rhs76_
        lhs51_[lhs52_] = rhs77_
        lhs53_[lhs54_] = rhs78_
        lhs55_.f0 = rhs79_
        if (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]:
            d_475_v60_: _dafny.MultiSet
            d_475_v60_ = _dafny.MultiSet([p0])
            pat_let_tv12_ = d_393_v2_
            pat_let_tv13_ = d_475_v60_
            def iife30_(_pat_let10_0):
                def iife31_(d_476_dt__update__tmp_h3_):
                    def iife32_(_pat_let11_0):
                        def iife33_(d_477_dt__update_hcf2_h0_):
                            return D0_DC2(d_477_dt__update_hcf2_h0_)
                        return iife33_(_pat_let11_0)
                    return iife32_(not((pat_let_tv12_) >= ((pat_let_tv13_).cardinality)))
                return iife31_(_pat_let10_0)
            d_469_v56_ = iife30_(d_469_v56_)
            d_478_v61_: _dafny.Array
            def lambda36_(d_479_p0_):
                def lambda37_(d_480_i7_):
                    return _dafny.SeqWithoutIsStrInference([d_479_p0_, d_479_p0_])

                return lambda37_

            init18_ = lambda36_(p0)
            nw91_ = _dafny.Array(None, 9)
            for i0_18_ in range(nw91_.length(0)):
                nw91_[i0_18_] = init18_(i0_18_)
            d_478_v61_ = nw91_
            d_481_v62_: _dafny.Map
            d_481_v62_ = _dafny.Map({d_393_v2_: d_478_v61_})
            d_478_v61_ = (d_478_v61_ if (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))] else (((d_481_v62_)[d_393_v2_] if (d_393_v2_) in (d_481_v62_) else d_478_v61_) if (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))] else d_478_v61_))
            if (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]:
                (globalState).f1 = len(_dafny.SeqWithoutIsStrInference([(p2) + (p2)]))
                d_482_v63_: D2
                d_482_v63_ = D2_DC7(d_392_v1_)
                rhs80_ = d_392_v1_
                rhs81_ = d_482_v63_
                d_392_v1_ = rhs80_
                d_482_v63_ = rhs81_
                d_483_v64_: _dafny.Seq
                d_483_v64_ = _dafny.SeqWithoutIsStrInference([True, p0])
                d_484_v65_: _dafny.Map
                d_484_v65_ = _dafny.Map({len(d_483_v64_): (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]})
                d_485_v66_: _dafny.Array
                nw92_ = _dafny.Array(int(0), 15)
                d_485_v66_ = nw92_
                d_486_v67_: _dafny.Map
                d_486_v67_ = _dafny.Map({len(d_484_v65_): d_485_v66_})
                (globalState).f0 = (len(d_486_v67_)) < (993)
                d_487_v68_: D5
                d_487_v68_ = D5_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhrajd")))
                (globalState).f3 = ((d_393_v2_ if (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))] else len((d_487_v68_).cf26))) - (d_393_v2_)
                d_488_v69_: C0
                nw93_ = C0()
                nw93_.ctor__((d_392_v1_).f6)
                d_488_v69_ = nw93_
            elif True:
                (globalState).f3 = default__.safeDivisionInt(d_393_v2_, len(default__.fm4((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], (d_475_v60_).cardinality, d_393_v2_, d_409_v14_, globalState)))
                d_489_v70_: _dafny.Seq
                d_489_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcqa"))
                d_490_v71_: D5
                d_490_v71_ = D5_DC18(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcxcktchu"))), d_393_v2_)
                d_491_v72_: _dafny.Seq
                d_491_v72_ = _dafny.SeqWithoutIsStrInference([D5_DC18(d_393_v2_, d_393_v2_), d_490_v71_, d_490_v71_, D5_DC18(d_393_v2_, -180), d_490_v71_])
                d_492_v73_: _dafny.Map
                d_492_v73_ = _dafny.Map({(default__.fm1(default__.fm1(d_393_v2_, globalState), globalState)) - (d_393_v2_): d_491_v72_})
                d_493_v74_: _dafny.Seq
                d_493_v74_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_494_v75_: _dafny.MultiSet
                d_494_v75_ = _dafny.MultiSet([d_393_v2_])
                index74_ = default__.safeIndex(231, (d_405_v12_).length(0))
                rhs82_ = (d_392_v1_).fm7(_dafny.CodePoint('j'), globalState)
                rhs83_ = p2
                rhs84_ = 174
                rhs85_ = default__.fm15(((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]) == ((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]), ((d_493_v74_).set(default__.safeIndex((d_494_v75_).cardinality, len(d_493_v74_)), True)) + (_dafny.SeqWithoutIsStrInference([p0, p0, p0, (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], not(p0)])), (d_493_v74_) + (d_493_v74_), globalState)
                rhs86_ = p0
                lhs56_ = globalState
                lhs57_ = globalState
                lhs58_ = d_405_v12_
                lhs59_ = default__.safeIndex(231, (d_405_v12_).length(0))
                lhs56_.f3 = rhs82_
                d_489_v70_ = rhs83_
                lhs57_.f3 = rhs84_
                d_492_v73_ = rhs85_
                lhs58_[lhs59_] = rhs86_
                d_478_v61_ = d_478_v61_
                d_495_v76_: _dafny.Map
                d_495_v76_ = _dafny.Map({d_393_v2_: (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]})
                d_496_v77_: _dafny.Map
                d_496_v77_ = _dafny.Map({p0: len(d_495_v76_)})
                d_497_v78_: _dafny.Array
                nw94_ = _dafny.Array(None, 19)
                nw94_[int(0)] = len(d_496_v77_)
                nw94_[int(1)] = (0) - (d_393_v2_)
                nw94_[int(2)] = d_393_v2_
                nw94_[int(3)] = d_393_v2_
                nw94_[int(4)] = len(_dafny.SeqWithoutIsStrInference([p0]))
                nw94_[int(5)] = d_393_v2_
                nw94_[int(6)] = d_393_v2_
                nw94_[int(7)] = (d_494_v75_).cardinality
                nw94_[int(8)] = d_393_v2_
                nw94_[int(9)] = (0) - (d_393_v2_)
                nw94_[int(10)] = default__.fm1(len(default__.fm10(globalState)), globalState)
                nw94_[int(11)] = d_393_v2_
                nw94_[int(12)] = d_393_v2_
                nw94_[int(13)] = d_393_v2_
                nw94_[int(14)] = d_393_v2_
                nw94_[int(15)] = d_393_v2_
                nw94_[int(16)] = d_393_v2_
                nw94_[int(17)] = 655
                nw94_[int(18)] = d_393_v2_
                d_497_v78_ = nw94_
                d_498_v79_: _dafny.Seq
                d_498_v79_ = _dafny.SeqWithoutIsStrInference([d_497_v78_])
                d_499_v80_: _dafny.Set
                d_499_v80_ = _dafny.Set({_dafny.Set({(d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], not(p0)})})
                d_500_v81_: _dafny.Map
                d_500_v81_ = _dafny.Map({len(d_499_v80_): d_497_v78_})
                d_501_v82_: _dafny.Array
                nw95_ = _dafny.Array(None, 28)
                nw95_[int(0)] = d_498_v79_
                nw95_[int(1)] = d_498_v79_
                nw95_[int(2)] = d_498_v79_
                nw95_[int(3)] = d_498_v79_
                nw95_[int(4)] = d_498_v79_
                nw95_[int(5)] = d_498_v79_
                nw95_[int(6)] = d_498_v79_
                nw95_[int(7)] = _dafny.SeqWithoutIsStrInference([d_497_v78_])
                nw95_[int(8)] = d_498_v79_
                nw95_[int(9)] = d_498_v79_
                nw95_[int(10)] = d_498_v79_
                nw95_[int(11)] = d_498_v79_
                nw95_[int(12)] = d_498_v79_
                nw95_[int(13)] = d_498_v79_
                nw95_[int(14)] = _dafny.SeqWithoutIsStrInference([d_497_v78_, d_497_v78_])
                nw95_[int(15)] = d_498_v79_
                nw95_[int(16)] = d_498_v79_
                nw95_[int(17)] = (d_498_v79_) + (_dafny.SeqWithoutIsStrInference([d_497_v78_, d_497_v78_]))
                nw95_[int(18)] = d_498_v79_
                nw95_[int(19)] = d_498_v79_
                nw95_[int(20)] = ((d_498_v79_) + (d_498_v79_)).set(default__.safeIndex(len(default__.fm16(globalState)), len((d_498_v79_) + (d_498_v79_))), ((d_500_v81_)[992] if (992) in (d_500_v81_) else d_497_v78_))
                nw95_[int(21)] = (d_498_v79_).set(default__.safeIndex(d_393_v2_, len(d_498_v79_)), d_497_v78_)
                nw95_[int(22)] = ((d_498_v79_) + (d_498_v79_)).set(default__.safeIndex(len(d_493_v74_), len((d_498_v79_) + (d_498_v79_))), d_497_v78_)
                nw95_[int(23)] = _dafny.SeqWithoutIsStrInference([d_497_v78_, d_497_v78_])
                nw95_[int(24)] = d_498_v79_
                nw95_[int(25)] = d_498_v79_
                nw95_[int(26)] = (d_498_v79_) + (d_498_v79_)
                nw95_[int(27)] = d_498_v79_
                d_501_v82_ = nw95_
                d_502_v83_: D6
                d_502_v83_ = D6_DC19(d_498_v79_)
                index75_ = default__.safeIndex(463, (d_501_v82_).length(0))
                (d_501_v82_)[index75_] = (d_502_v83_).cf34
                index76_ = default__.safeIndex(463, (d_501_v82_).length(0))
                (d_501_v82_)[index76_] = (((d_498_v79_) + (d_498_v79_)) + (d_498_v79_)).set(default__.safeIndex(d_393_v2_, len(((d_498_v79_) + (d_498_v79_)) + (d_498_v79_))), d_497_v78_)
                d_496_v77_ = (d_496_v77_).set(not(p0), default__.safeModuloInt(d_393_v2_, d_393_v2_))
            d_503_v84_: _dafny.Seq
            d_503_v84_ = _dafny.SeqWithoutIsStrInference([p0, (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]])
            d_504_v85_: D6
            d_504_v85_ = D6_DC20(d_503_v84_, p0)
            source5_ = d_504_v85_
            if source5_.is_DC20:
                d_505___mcc_h4_ = source5_.cf35
                d_506___mcc_h5_ = source5_.cf36
                d_507_cf36_ = d_506___mcc_h5_
                d_508_cf35_ = d_505___mcc_h4_
                d_509_v86_: D6
                d_509_v86_ = D6_DC22((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nock"))) if d_507_cf36_ else d_393_v2_), default__.fm1(len(p2), globalState))
                d_510_v87_: _dafny.Set
                d_510_v87_ = _dafny.Set({d_392_v1_})
                rhs87_ = len(d_510_v87_)
                rhs88_ = (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]
                rhs89_ = D6_DC22(default__.safeDivisionInt(d_393_v2_, -337), (0) - ((len(d_408_v13_)) + (d_393_v2_)))
                d_393_v2_ = rhs87_
                d_507_cf36_ = rhs88_
                d_509_v86_ = rhs89_
                index77_ = default__.safeIndex(231, (d_405_v12_).length(0))
                (d_405_v12_)[index77_] = not((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))])
                d_393_v2_ = d_393_v2_
                (globalState).f1 = d_393_v2_
            elif source5_.is_DC21:
                d_511___mcc_h6_ = source5_.cf37
                d_512___mcc_h7_ = source5_.cf38
                d_513___mcc_h8_ = source5_.cf39
                d_514___mcc_h9_ = source5_.cf40
                d_515_cf40_ = d_514___mcc_h9_
                d_516_cf39_ = d_513___mcc_h8_
                d_517_cf38_ = d_512___mcc_h7_
                d_518_cf37_ = d_511___mcc_h6_
                index78_ = default__.safeIndex(231, (d_405_v12_).length(0))
                (d_405_v12_)[index78_] = d_515_cf40_
                index79_ = default__.safeIndex(231, (d_405_v12_).length(0))
                (d_405_v12_)[index79_] = d_515_cf40_
                d_515_cf40_ = p0
                d_519_v88_: _dafny.MultiSet
                d_519_v88_ = _dafny.MultiSet([d_516_cf39_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))).set(default__.safeIndex(d_393_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa")))), _dafny.CodePoint('o'))])
                index80_ = default__.safeIndex(231, (d_405_v12_).length(0))
                (d_405_v12_)[index80_] = ((d_519_v88_) | (d_519_v88_)).issubset(d_519_v88_)
            elif source5_.is_DC22:
                d_520___mcc_h10_ = source5_.cf41
                d_521___mcc_h11_ = source5_.cf42
                d_522_cf42_ = d_521___mcc_h11_
                d_523_cf41_ = d_520___mcc_h10_
                d_524_v89_: _dafny.Set
                d_524_v89_ = _dafny.Set({d_522_cf42_})
                d_525_v90_: _dafny.Array
                def lambda38_(d_526_v2_):
                    def lambda39_(d_527_i8_):
                        return (d_527_i8_) + (d_526_v2_)

                    return lambda39_

                init19_ = lambda38_(d_393_v2_)
                nw96_ = _dafny.Array(None, 15)
                for i0_19_ in range(nw96_.length(0)):
                    nw96_[i0_19_] = init19_(i0_19_)
                d_525_v90_ = nw96_
                d_528_v91_: _dafny.Set
                d_528_v91_ = _dafny.Set({d_525_v90_, d_525_v90_, d_525_v90_})
                d_529_v92_: _dafny.Map
                d_529_v92_ = _dafny.Map({d_524_v89_: len(d_528_v91_)})
                d_529_v92_ = (d_529_v92_).set((d_524_v89_).intersection(d_524_v89_), d_393_v2_)
                d_530_v93_: _dafny.Array
                nw97_ = _dafny.Array(None, 26)
                nw97_[int(0)] = d_394_v3_
                nw97_[int(1)] = (_dafny.Map({d_392_v1_: d_393_v2_})) | (_dafny.Map({d_392_v1_: d_393_v2_}))
                nw97_[int(2)] = (d_394_v3_).set(d_392_v1_, d_393_v2_)
                nw97_[int(3)] = d_394_v3_
                nw97_[int(4)] = d_394_v3_
                nw97_[int(5)] = d_394_v3_
                nw97_[int(6)] = d_394_v3_
                nw97_[int(7)] = d_394_v3_
                nw97_[int(8)] = (d_394_v3_) | (d_394_v3_)
                nw97_[int(9)] = d_394_v3_
                nw97_[int(10)] = d_394_v3_
                nw97_[int(11)] = d_394_v3_
                nw97_[int(12)] = d_394_v3_
                nw97_[int(13)] = _dafny.Map({(d_395_v4_)[default__.safeIndex(d_393_v2_, len(d_395_v4_))]: (0) - (len(d_411_v16_))})
                nw97_[int(14)] = (_dafny.Map({d_392_v1_: d_522_cf42_})) | ((d_394_v3_).set(d_392_v1_, d_522_cf42_))
                nw97_[int(15)] = d_394_v3_
                nw97_[int(16)] = (d_394_v3_).set(d_392_v1_, d_523_cf41_)
                nw97_[int(17)] = d_394_v3_
                nw97_[int(18)] = d_394_v3_
                nw97_[int(19)] = d_394_v3_
                nw97_[int(20)] = ((d_394_v3_).set(d_392_v1_, default__.fm1(d_393_v2_, globalState))) | (d_394_v3_)
                nw97_[int(21)] = d_394_v3_
                nw97_[int(22)] = _dafny.Map({d_392_v1_: len(p1)})
                nw97_[int(23)] = d_394_v3_
                nw97_[int(24)] = d_394_v3_
                nw97_[int(25)] = (d_394_v3_) | (d_394_v3_)
                d_530_v93_ = nw97_
                d_530_v93_ = d_530_v93_
                d_531_v94_: C0
                nw98_ = C0()
                nw98_.ctor__(((d_392_v1_).f6 if False else (d_392_v1_).f6))
                d_531_v94_ = nw98_
                index81_ = default__.safeIndex(589, (d_525_v90_).length(0))
                (d_525_v90_)[index81_] = (-10) * (d_523_cf41_)
                index82_ = default__.safeIndex(589, (d_525_v90_).length(0))
                (d_525_v90_)[index82_] = (0) - ((d_393_v2_ if (d_503_v84_)[default__.safeIndex(-978, len(d_503_v84_))] else d_393_v2_))
            elif source5_.is_DC19:
                d_532___mcc_h12_ = source5_.cf34
                d_533_cf34_ = d_532___mcc_h12_
                index83_ = default__.safeIndex(231, (d_405_v12_).length(0))
                (d_405_v12_)[index83_] = p0
                d_475_v60_ = d_475_v60_
                d_534_v95_: _dafny.Set
                d_534_v95_ = _dafny.Set({p1})
                (globalState).f1 = ((d_393_v2_) - (d_393_v2_)) * ((0) - ((d_393_v2_) + (len(d_534_v95_))))
                d_535_v96_: D2
                d_535_v96_ = D2_DC8((0) - (d_393_v2_), len(d_503_v84_))
                d_536_v97_: _dafny.Seq
                d_536_v97_ = _dafny.SeqWithoutIsStrInference([default__.fm17(globalState), D2_DC8(d_393_v2_, len(p2)), d_535_v96_])
                d_537_v98_: _dafny.Set
                d_537_v98_ = _dafny.Set({d_392_v1_})
                d_538_v99_: _dafny.Seq
                d_538_v99_ = _dafny.SeqWithoutIsStrInference([(d_536_v97_).set(default__.safeIndex(d_393_v2_, len(d_536_v97_)), d_535_v96_), _dafny.SeqWithoutIsStrInference([d_535_v96_, D2_DC8((0) - (len(d_537_v98_)), d_393_v2_), d_535_v96_]), (_dafny.SeqWithoutIsStrInference([d_535_v96_])) + (d_536_v97_), _dafny.SeqWithoutIsStrInference([d_535_v96_, d_535_v96_, d_535_v96_])])
                d_538_v99_ = _dafny.SeqWithoutIsStrInference([(d_536_v97_) + (d_536_v97_) for d_539_i9_ in range(default__.abs(297))])
            elif True:
                d_540___mcc_h13_ = source5_.cf43
                d_541_cf43_ = d_540___mcc_h13_
                (globalState).f3 = d_393_v2_
                d_542_v100_: C0
                nw99_ = C0()
                nw99_.ctor__(((d_392_v1_).f6 if False else (d_392_v1_).f6))
                d_542_v100_ = nw99_
                index84_ = default__.safeIndex(903, (d_478_v61_).length(0))
                (d_478_v61_)[index84_] = d_503_v84_
                index85_ = default__.safeIndex(903, (d_478_v61_).length(0))
                (d_478_v61_)[index85_] = (d_503_v84_) + ((default__.fm11(p0, (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], p3, globalState)).set(default__.safeIndex(d_393_v2_, len(default__.fm11(p0, (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], p3, globalState))), True))
                (globalState).f1 = d_393_v2_
            d_543_v101_: int
            d_544_v102_: _dafny.Array
            d_545_v103_: D0
            out29_: int
            out30_: _dafny.Array
            out31_: D0
            out29_, out30_, out31_ = (self).m0(globalState)
            d_543_v101_ = out29_
            d_544_v102_ = out30_
            d_545_v103_ = out31_
        elif True:
            d_546_v104_: _dafny.Map
            d_546_v104_ = _dafny.Map({not(p0): d_409_v14_})
            d_546_v104_ = (d_546_v104_).set((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))], d_409_v14_)
            d_409_v14_ = d_409_v14_
            d_547_v105_: _dafny.MultiSet
            d_547_v105_ = _dafny.MultiSet([not((d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]), (d_405_v12_)[default__.safeIndex(231, (d_405_v12_).length(0))]])
            d_547_v105_ = d_547_v105_
            index86_ = default__.safeIndex(231, (d_405_v12_).length(0))
            (d_405_v12_)[index86_] = default__.fm0(d_393_v2_, globalState)
            d_548_v106_: _dafny.Array
            nw100_ = _dafny.Array(_dafny.Seq({}), 28)
            d_548_v106_ = nw100_
            d_548_v106_ = d_548_v106_
        d_549_v107_: D3
        d_549_v107_ = D3_DC11(d_393_v2_)
        r0 = (d_393_v2_) + ((d_549_v107_).cf14)
        return r0

