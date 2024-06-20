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
        return default__.safeModuloInt(856, -444)

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return (False) and ((-520) > (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_0_i0_ in range(default__.abs(-618))])))))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-158, 228):
                d_1_v0_: int = compr_0_
                if ((-158) <= (d_1_v0_)) and ((d_1_v0_) < (228)):
                    coll0_[default__.safeModuloInt(d_1_v0_, -841)] = not (True) or (False)
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if False else _dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True, True])))

    @staticmethod
    def fm10(p0, globalState):
        return ((_dafny.Map({968: True}) if True else _dafny.Map({-115: True}))) | (_dafny.Map({-981: not(False)}))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return _dafny.Map({(0) - ((D2_DC7(644)).cf10): 76})

    @staticmethod
    def fm13(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(507, 23):
                d_2_v0_: int = compr_1_
                if ((507) <= (d_2_v0_)) and ((d_2_v0_) < (23)):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeModuloInt(d_2_v0_, len(_dafny.SeqWithoutIsStrInference([True, False])))]))
            return _dafny.Set(coll1_)
        return (iife1_()
        ) | ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gj"))), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False, False]))}))})).intersection(_dafny.Set({(0) - (len(_dafny.Map({True: _dafny.CodePoint('q')}))), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_3_i0_ in range(default__.abs(383))]): True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyck"))), len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))})), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_4_i1_ in range(default__.abs(-66))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_5_i2_ in range(default__.abs(-114))])])).cardinality]))}))})))

    @staticmethod
    def fm14(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(535, 768):
                d_7_v0_: int = compr_2_
                if ((535) <= (d_7_v0_)) and ((d_7_v0_) < (768)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_7_v0_, 455)]))
            return _dafny.Set(coll2_)
        return _dafny.Set({True, (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_6_i0_ in range(default__.abs(572))])), len(iife2_()
        )})).ispropersubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqbhaywk"))), 160}))})

    @staticmethod
    def fm15(p0, p1, globalState):
        return _dafny.CodePoint('e')

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([True, True, False, False]), _dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([False])])) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(D8_DC24(_dafny.Set({not(True), True}), _dafny.SeqWithoutIsStrInference([not(False)]))).cf31, _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True, True])]) if not(True) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, True, True, False]), _dafny.SeqWithoutIsStrInference([True, True])]))))

    @staticmethod
    def fm17(p0, p1, globalState):
        source0_ = D6_DC18(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqs")))])))
        if source0_.is_DC16:
            d_8___mcc_h0_ = source0_.cf22
            d_9___mcc_h1_ = source0_.cf23
            d_10_cf23_ = d_9___mcc_h1_
            d_11_cf22_ = d_8___mcc_h0_
            return D2_DC8(False, len((D12_DC36(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtpervwnl"))}))).cf52), False)
        elif source0_.is_DC17:
            d_12___mcc_h2_ = source0_.cf24
            d_13_cf24_ = d_12___mcc_h2_
            return D2_DC8(d_13_cf24_, 877, d_13_cf24_)
        elif source0_.is_DC18:
            d_14___mcc_h3_ = source0_.cf25
            d_15_cf25_ = d_14___mcc_h3_
            return D2_DC8(False, d_15_cf25_, not(not(not(True))))
        elif source0_.is_DC15:
            d_16___mcc_h4_ = source0_.cf21
            d_17_cf21_ = d_16___mcc_h4_
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: _dafny.Set
                for compr_3_ in (_dafny.Set({_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_18_i0_ in range(default__.abs(647))])))}), _dafny.Set({-146, 401, -293, len(d_17_cf21_), (_dafny.MultiSet([(_dafny.MultiSet([(0) - ((_dafny.MultiSet([471])).cardinality), 386])).cardinality, len(_dafny.Map({False: False}))])).cardinality})})).Elements:
                    d_19_v0_: _dafny.Set = compr_3_
                    if (d_19_v0_) in (_dafny.Set({_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_18_i0_ in range(default__.abs(647))])))}), _dafny.Set({-146, 401, -293, len(d_17_cf21_), (_dafny.MultiSet([(_dafny.MultiSet([(0) - ((_dafny.MultiSet([471])).cardinality), 386])).cardinality, len(_dafny.Map({False: False}))])).cardinality})})):
                        coll3_[d_19_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_20_i1_ in range(default__.abs(708))]))
                return _dafny.Map(coll3_)
            return D2_DC8(False, len(iife3_()
), False)
        elif True:
            d_21___mcc_h5_ = source0_.cf26
            d_22_cf26_ = d_21___mcc_h5_
            return D2_DC8(not(True), -990, False)

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        return _dafny.Map({(not(not(False))) or (False): False})

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.MultiSet([False])) | (_dafny.MultiSet([True]))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.CodePoint('v'), _dafny.CodePoint('u'), ((D9_DC25(_dafny.CodePoint('k'))).cf32 if True else _dafny.CodePoint('w'))})

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        source1_ = D8_DC23(663)
        if source1_.is_DC23:
            d_23___mcc_h0_ = source1_.cf29
            d_24_cf29_ = d_23___mcc_h0_
            return (_dafny.MultiSet([d_24_cf29_, d_24_cf29_, len(_dafny.SeqWithoutIsStrInference([True, False]))])) | (_dafny.MultiSet([240]))
        elif source1_.is_DC24:
            d_25___mcc_h1_ = source1_.cf30
            d_26___mcc_h2_ = source1_.cf31
            d_27_cf31_ = d_26___mcc_h2_
            d_28_cf30_ = d_25___mcc_h1_
            return _dafny.MultiSet([267])
        elif True:
            d_29___mcc_h3_ = source1_.cf28
            d_30_cf28_ = d_29___mcc_h3_
            return (_dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('i'): 992}))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgetcs")))])))

    @staticmethod
    def fm22(p0, globalState):
        return D8_DC24(_dafny.Set({False, True}), (_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_31_i1_ in range(default__.abs(220))])) for d_32_i0_ in range(default__.abs(459))]) if not(False) else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({-19: False}): not(False)})), 49]))) + (_dafny.SeqWithoutIsStrInference([-289 for d_33_i2_ in range(default__.abs(-953))]))

    @staticmethod
    def fm24(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D5_DC14(True), D5_DC14(True)])) + (_dafny.SeqWithoutIsStrInference([D5_DC14(True), D5_DC14(False)]))) + (_dafny.SeqWithoutIsStrInference([D5_DC14(True), D5_DC14(True)]))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return D5_DC14(not(False))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D5_DC13(True, False) for d_34_i0_ in range(default__.abs(739))])

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        return D9_DC26()

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        if p3:
            d_35_v0_: _dafny.Seq
            d_35_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnxtuopx"))
            d_36_v1_: _dafny.MultiSet
            d_36_v1_ = _dafny.MultiSet([p3])
            d_37_v2_: _dafny.Array
            def lambda0_(d_38_v1_):
                def lambda1_(d_39_i2_):
                    return D2_DC7((_dafny.MultiSet([d_38_v1_])).cardinality)

                return lambda1_

            init0_ = lambda0_(d_36_v1_)
            nw0_ = _dafny.Array(None, 22)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_37_v2_ = nw0_
            d_40_v3_: T2
            nw1_ = C3()
            nw1_.ctor__(d_37_v2_, _dafny.CodePoint('v'), 419)
            d_40_v3_ = nw1_
            d_41_v4_: _dafny.Map
            d_41_v4_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([-194 for d_42_i1_ in range(default__.abs(668))]): d_40_v3_})
            d_43_v5_: _dafny.Map
            d_43_v5_ = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_44_i3_ in range(default__.abs(-89))])})
            d_45_v6_: _dafny.Array
            nw2_ = _dafny.Array(None, 13)
            nw2_[int(0)] = d_35_v0_
            nw2_[int(1)] = default__.fm11(((d_36_v1_)[p3] if (p3) in (d_36_v1_) else len(d_35_v0_)), p2, p3, globalState)
            nw2_[int(2)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_46_i0_ in range(default__.abs(120))])
            nw2_[int(3)] = d_35_v0_
            nw2_[int(4)] = (d_35_v0_) + (default__.fm11(len(d_41_v4_), p0, p3, globalState))
            nw2_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exf"))
            nw2_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivgkwls"))
            nw2_[int(7)] = ((d_43_v5_)[p3] if (p3) in (d_43_v5_) else d_35_v0_)
            nw2_[int(8)] = d_35_v0_
            nw2_[int(9)] = d_35_v0_
            nw2_[int(10)] = d_35_v0_
            nw2_[int(11)] = d_35_v0_
            nw2_[int(12)] = d_35_v0_
            d_45_v6_ = nw2_
            index0_ = default__.safeIndex(44, (d_45_v6_).length(0))
            (d_45_v6_)[index0_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_47_i4_ in range(default__.abs(473))])
            d_48_v7_: str
            d_48_v7_ = _dafny.CodePoint('c')
            index1_ = default__.safeIndex(44, (d_45_v6_).length(0))
            rhs0_ = (p1) * (p0)
            rhs1_ = (d_35_v0_).set(default__.safeIndex(d_40_v3_.f12, len(d_35_v0_)), d_48_v7_)
            lhs0_ = globalState
            lhs1_ = d_45_v6_
            lhs2_ = default__.safeIndex(44, (d_45_v6_).length(0))
            lhs0_.f5 = rhs0_
            lhs1_[lhs2_] = rhs1_
            d_49_v8_: _dafny.Seq
            d_49_v8_ = _dafny.SeqWithoutIsStrInference([d_40_v3_.f12])
            d_50_v9_: _dafny.Map
            d_50_v9_ = _dafny.Map({len((d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))]): p0})
            d_51_v10_: _dafny.Array
            nw3_ = _dafny.Array(None, 13)
            nw3_[int(0)] = d_49_v8_
            nw3_[int(1)] = (d_49_v8_) + (d_49_v8_)
            nw3_[int(2)] = d_49_v8_
            nw3_[int(3)] = d_49_v8_
            nw3_[int(4)] = d_49_v8_
            nw3_[int(5)] = (d_49_v8_) + (_dafny.SeqWithoutIsStrInference([p1]))
            nw3_[int(6)] = d_49_v8_
            nw3_[int(7)] = d_49_v8_
            nw3_[int(8)] = default__.fm23(p3, 16, -473, globalState)
            nw3_[int(9)] = d_49_v8_
            nw3_[int(10)] = _dafny.SeqWithoutIsStrInference([len(d_50_v9_)])
            nw3_[int(11)] = (d_49_v8_) + (d_49_v8_)
            nw3_[int(12)] = _dafny.SeqWithoutIsStrInference([p2])
            d_51_v10_ = nw3_
            index2_ = default__.safeIndex(191, (d_51_v10_).length(0))
            (d_51_v10_)[index2_] = d_49_v8_
            d_52_v11_: _dafny.Map
            d_52_v11_ = _dafny.Map({p3: d_49_v8_})
            index3_ = default__.safeIndex(191, (d_51_v10_).length(0))
            (d_51_v10_)[index3_] = ((d_52_v11_)[True] if (True) in (d_52_v11_) else d_49_v8_)
            d_53_v12_: D0
            d_53_v12_ = D0_DC0(p3)
            source2_ = d_53_v12_
            if source2_.is_DC1:
                d_54___mcc_h0_ = source2_.cf1
                d_55_cf1_ = d_54___mcc_h0_
                d_56_v13_: _dafny.Map
                d_56_v13_ = _dafny.Map({d_40_v3_.f12: p3})
                d_56_v13_ = (d_56_v13_).set(p2, p3)
                (d_40_v3_).f12 = p1
                d_57_v14_: D5
                d_57_v14_ = D5_DC14(False)
                d_58_v15_: _dafny.Seq
                d_58_v15_ = _dafny.SeqWithoutIsStrInference([D5_DC14(p3), d_57_v14_])
                d_59_v16_: _dafny.Seq
                d_59_v16_ = _dafny.SeqWithoutIsStrInference([d_58_v15_, d_58_v15_])
                d_60_v17_: _dafny.Seq
                d_60_v17_ = _dafny.SeqWithoutIsStrInference([d_37_v2_, d_37_v2_])
                d_61_v18_: T1
                nw4_ = C3()
                nw4_.ctor__((d_60_v17_)[default__.safeIndex((_dafny.MultiSet([(0) - (p2)])).cardinality, len(d_60_v17_))], _dafny.CodePoint('s'), (_dafny.MultiSet([p3, p3])).cardinality)
                d_61_v18_ = nw4_
                d_62_v19_: _dafny.MultiSet
                d_62_v19_ = _dafny.MultiSet([d_61_v18_, d_61_v18_])
                d_63_v20_: _dafny.Map
                d_63_v20_ = _dafny.Map({(d_62_v19_).set(d_61_v18_, default__.abs(d_40_v3_.f12)): d_48_v7_})
                d_64_v21_: _dafny.MultiSet
                d_64_v21_ = _dafny.MultiSet([p2, d_40_v3_.f12])
                d_65_v22_: _dafny.Map
                d_65_v22_ = _dafny.Map({(((d_59_v16_)[default__.safeIndex(len(d_63_v20_), len(d_59_v16_))]) + (default__.fm24(p3, d_40_v3_.f12, globalState))).set(default__.safeIndex((d_64_v21_).cardinality, len(((d_59_v16_)[default__.safeIndex(len(d_63_v20_), len(d_59_v16_))]) + (default__.fm24(p3, d_40_v3_.f12, globalState)))), D5_DC14(p3)): d_40_v3_})
                d_65_v22_ = (d_65_v22_).set((d_58_v15_) + ((d_58_v15_).set(default__.safeIndex(d_61_v18_.f12, len(d_58_v15_)), d_57_v14_)), d_40_v3_)
                d_66_v23_: bool
                d_66_v23_ = False
                d_66_v23_ = p3
            elif source2_.is_DC2:
                d_67___mcc_h1_ = source2_.cf2
                d_68___mcc_h2_ = source2_.cf3
                d_69___mcc_h3_ = source2_.cf4
                d_70___mcc_h4_ = source2_.cf5
                d_71_cf5_ = d_70___mcc_h4_
                d_72_cf4_ = d_69___mcc_h3_
                d_73_cf3_ = d_68___mcc_h2_
                d_74_cf2_ = d_67___mcc_h1_
                d_49_v8_ = (_dafny.SeqWithoutIsStrInference([p2, (0) - (p0), 973])) + (_dafny.SeqWithoutIsStrInference([len(d_35_v0_)]))
                d_75_v24_: _dafny.Map
                d_75_v24_ = _dafny.Map({d_48_v7_: p3})
                d_76_v25_: _dafny.Array
                nw5_ = _dafny.Array(None, 18)
                nw5_[int(0)] = (d_40_v3_.f12) - (p2)
                nw5_[int(1)] = len((d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))])
                nw5_[int(2)] = (0) - (default__.safeDivisionInt(d_40_v3_.f12, p1))
                nw5_[int(3)] = p0
                nw5_[int(4)] = d_40_v3_.f12
                nw5_[int(5)] = (p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "die"))))
                nw5_[int(6)] = (0) - ((default__.fm0(493, len(d_75_v24_), globalState)) * (685))
                nw5_[int(7)] = p1
                nw5_[int(8)] = len((d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))])
                nw5_[int(9)] = (0) - (d_40_v3_.f12)
                nw5_[int(10)] = p2
                nw5_[int(11)] = p1
                nw5_[int(12)] = p1
                nw5_[int(13)] = p1
                nw5_[int(14)] = len(d_35_v0_)
                nw5_[int(15)] = (len(_dafny.SeqWithoutIsStrInference([len(d_49_v8_) for d_77_i5_ in range(default__.abs(540))]))) - (p0)
                nw5_[int(16)] = d_40_v3_.f12
                nw5_[int(17)] = d_40_v3_.f12
                d_76_v25_ = nw5_
                index4_ = default__.safeIndex(571, (d_76_v25_).length(0))
                (d_76_v25_)[index4_] = (d_40_v3_.f12) * (d_40_v3_.f12)
                d_78_v26_: _dafny.Set
                d_78_v26_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scckxsky")), (d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))]})
                index5_ = default__.safeIndex(571, (d_76_v25_).length(0))
                (d_76_v25_)[index5_] = (p2 if d_71_cf5_ else len(d_78_v26_))
                d_48_v7_ = d_48_v7_
                d_71_cf5_ = not((p2) < (p0))
            elif source2_.is_DC0:
                d_79___mcc_h5_ = source2_.cf0
                d_80_cf0_ = d_79___mcc_h5_
                d_81_v28_: _dafny.Set
                d_81_v28_ = _dafny.Set({p1})
                nw6_ = C0()
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: int
                    for compr_4_ in (default__.fm23(d_80_cf0_, d_40_v3_.f12, p2, globalState)).Elements:
                        d_82_v27_: int = compr_4_
                        if (d_82_v27_) in (default__.fm23(d_80_cf0_, d_40_v3_.f12, p2, globalState)):
                            coll4_[(d_82_v27_) * (d_40_v3_.f12)] = d_80_cf0_
                    return _dafny.Map(coll4_)
                nw6_.ctor__(iife4_()
                , len(d_81_v28_))
                d_40_v3_ = nw6_
                d_83_v29_: _dafny.Array
                nw7_ = _dafny.Array(False, 22)
                d_83_v29_ = nw7_
                index6_ = default__.safeIndex(609, (d_83_v29_).length(0))
                (d_83_v29_)[index6_] = d_80_cf0_
                d_84_v30_: D5
                d_84_v30_ = D5_DC14(False)
                d_85_v31_: _dafny.Seq
                d_85_v31_ = _dafny.SeqWithoutIsStrInference([(d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))]])
                index7_ = default__.safeIndex(609, (d_83_v29_).length(0))
                rhs2_ = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")), d_35_v0_])) < (d_85_v31_)
                rhs3_ = default__.fm25(d_80_cf0_, d_80_cf0_, default__.safeModuloInt(default__.fm0(p2, (0) - (p2), globalState), d_40_v3_.f12), globalState)
                rhs4_ = d_80_cf0_
                lhs3_ = d_83_v29_
                lhs4_ = default__.safeIndex(609, (d_83_v29_).length(0))
                lhs3_[lhs4_] = rhs2_
                d_84_v30_ = rhs3_
                d_80_cf0_ = rhs4_
                index8_ = default__.safeIndex(609, (d_83_v29_).length(0))
                (d_83_v29_)[index8_] = d_80_cf0_
                rhs5_ = d_40_v3_.f12
                rhs6_ = (len((d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))]) if (d_83_v29_)[default__.safeIndex(609, (d_83_v29_).length(0))] else ((d_51_v10_)[default__.safeIndex(191, (d_51_v10_).length(0))])[default__.safeIndex(p2, len((d_51_v10_)[default__.safeIndex(191, (d_51_v10_).length(0))]))])
                rhs7_ = (0) - (d_40_v3_.f12)
                rhs8_ = d_48_v7_
                lhs5_ = globalState
                lhs6_ = globalState
                lhs7_ = globalState
                lhs5_.f5 = rhs5_
                lhs6_.f5 = rhs6_
                lhs7_.f5 = rhs7_
                d_48_v7_ = rhs8_
            elif True:
                d_86___mcc_h6_ = source2_.cf6
                d_87_cf6_ = d_86___mcc_h6_
                d_88_v32_: _dafny.Array
                nw8_ = _dafny.Array(None, 14)
                d_88_v32_ = nw8_
                d_89_v33_: C3
                nw9_ = C3()
                nw9_.ctor__(d_37_v2_, d_48_v7_, p1)
                d_89_v33_ = nw9_
                index9_ = default__.safeIndex(0, (d_88_v32_).length(0))
                (d_88_v32_)[index9_] = d_89_v33_
                index10_ = default__.safeIndex(0, (d_88_v32_).length(0))
                nw10_ = C3()
                nw10_.ctor__(d_37_v2_, d_48_v7_, (0) - (p1))
                (d_88_v32_)[index10_] = nw10_
                d_90_v34_: _dafny.Array
                nw11_ = _dafny.Array(int(0), 25)
                d_90_v34_ = nw11_
                d_91_v35_: _dafny.MultiSet
                d_91_v35_ = _dafny.MultiSet([d_90_v34_, d_90_v34_])
                (d_40_v3_).f12 = ((d_91_v35_)[d_90_v34_] if (d_90_v34_) in (d_91_v35_) else 432)
                d_92_v36_: _dafny.Map
                d_92_v36_ = _dafny.Map({(d_40_v3_.f12) * (d_40_v3_.f12): p3})
                d_92_v36_ = (d_92_v36_).set((d_40_v3_.f12) * ((d_36_v1_).cardinality), p3)
                d_93_v37_: _dafny.Set
                d_93_v37_ = _dafny.Set({d_40_v3_.f12})
                d_94_v38_: C4
                nw12_ = C4()
                nw12_.ctor__(len((d_93_v37_) | (d_93_v37_)), (d_40_v3_.f12 if default__.fm1(-148, p3, p0, p3, globalState) else d_40_v3_.f12))
                d_94_v38_ = nw12_
            d_95_v39_: _dafny.Seq
            d_95_v39_ = _dafny.SeqWithoutIsStrInference([p3])
            if (_dafny.SeqWithoutIsStrInference([p3])) != ((d_95_v39_) + (d_95_v39_)):
                (d_40_v3_).f12 = ((d_50_v9_)[p0] if (p0) in (d_50_v9_) else p0)
                d_96_v40_: _dafny.Map
                d_96_v40_ = _dafny.Map({p3: p1})
                d_97_v41_: _dafny.Array
                nw13_ = _dafny.Array(None, 15)
                nw13_[int(0)] = default__.fm1(p0, False, (d_49_v8_)[default__.safeIndex((d_36_v1_).cardinality, len(d_49_v8_))], p3, globalState)
                nw13_[int(1)] = p3
                nw13_[int(2)] = p3
                nw13_[int(3)] = p3
                nw13_[int(4)] = p3
                nw13_[int(5)] = p3
                nw13_[int(6)] = (p3) or (not(False))
                nw13_[int(7)] = False
                nw13_[int(8)] = p3
                nw13_[int(9)] = not((default__.fm23(p3, default__.fm0(len(d_96_v40_), p2, globalState), p1, globalState)) <= ((d_51_v10_)[default__.safeIndex(191, (d_51_v10_).length(0))]))
                nw13_[int(10)] = p3
                nw13_[int(11)] = p3
                nw13_[int(12)] = p3
                nw13_[int(13)] = p3
                nw13_[int(14)] = (d_95_v39_)[default__.safeIndex(489, len(d_95_v39_))]
                d_97_v41_ = nw13_
                index11_ = default__.safeIndex(757, (d_97_v41_).length(0))
                (d_97_v41_)[index11_] = p3
                index12_ = default__.safeIndex(757, (d_97_v41_).length(0))
                (d_97_v41_)[index12_] = p3
                index13_ = default__.safeIndex(757, (d_97_v41_).length(0))
                (d_97_v41_)[index13_] = (d_97_v41_)[default__.safeIndex(757, (d_97_v41_).length(0))]
                index14_ = default__.safeIndex(757, (d_97_v41_).length(0))
                (d_97_v41_)[index14_] = p3
                (d_40_v3_).f12 = ((d_96_v40_)[(d_97_v41_)[default__.safeIndex(757, (d_97_v41_).length(0))]] if ((d_97_v41_)[default__.safeIndex(757, (d_97_v41_).length(0))]) in (d_96_v40_) else p2)
            elif True:
                d_98_v42_: _dafny.Array
                nw14_ = _dafny.Array(None, 14)
                d_98_v42_ = nw14_
                d_99_v43_: _dafny.Array
                def lambda2_(d_100_p3_):
                    def lambda3_(d_101_i6_):
                        return d_100_p3_

                    return lambda3_

                init1_ = lambda2_(p3)
                nw15_ = _dafny.Array(None, 4)
                for i0_1_ in range(nw15_.length(0)):
                    nw15_[i0_1_] = init1_(i0_1_)
                d_99_v43_ = nw15_
                d_102_v44_: _dafny.Map
                d_102_v44_ = _dafny.Map({d_98_v42_: d_99_v43_})
                d_102_v44_ = d_102_v44_
                d_103_v45_: D5
                d_103_v45_ = D5_DC13(p3, p3)
                d_104_v46_: _dafny.Seq
                d_104_v46_ = _dafny.SeqWithoutIsStrInference([d_103_v45_, d_103_v45_, d_103_v45_, d_103_v45_])
                d_105_v47_: _dafny.Map
                d_105_v47_ = _dafny.Map({p3: p3})
                d_106_v48_: _dafny.Set
                d_106_v48_ = _dafny.Set({d_104_v46_, (_dafny.SeqWithoutIsStrInference([d_103_v45_, d_103_v45_])) + (d_104_v46_), default__.fm26(p0, default__.fm15(d_105_v47_, d_40_v3_.f12, globalState), p3, d_48_v7_, globalState)})
                d_106_v48_ = d_106_v48_
                d_107_v49_: bool
                d_107_v49_ = False
                d_108_v50_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Map({}), 12)
                d_108_v50_ = nw16_
                d_109_v51_: _dafny.Map
                d_109_v51_ = _dafny.Map({p1: d_99_v43_})
                index15_ = default__.safeIndex(178, (d_108_v50_).length(0))
                (d_108_v50_)[index15_] = d_109_v51_
                index16_ = default__.safeIndex(178, (d_108_v50_).length(0))
                rhs9_ = default__.fm1(p0, d_107_v49_, (p0) + (-415), p3, globalState)
                rhs10_ = d_109_v51_
                lhs8_ = d_108_v50_
                lhs9_ = default__.safeIndex(178, (d_108_v50_).length(0))
                d_107_v49_ = rhs9_
                lhs8_[lhs9_] = rhs10_
                d_110_v52_: D10
                d_110_v52_ = D10_DC29(d_40_v3_)
                d_111_v53_: _dafny.Array
                nw17_ = _dafny.Array(None, 26)
                nw17_[int(0)] = d_40_v3_
                nw17_[int(1)] = d_40_v3_
                nw17_[int(2)] = d_40_v3_
                nw17_[int(3)] = d_40_v3_
                nw17_[int(4)] = d_40_v3_
                nw17_[int(5)] = d_40_v3_
                nw17_[int(6)] = d_40_v3_
                nw17_[int(7)] = (d_40_v3_ if not(True) else d_40_v3_)
                nw17_[int(8)] = d_40_v3_
                nw17_[int(9)] = d_40_v3_
                nw17_[int(10)] = d_40_v3_
                nw17_[int(11)] = d_40_v3_
                nw17_[int(12)] = d_40_v3_
                nw17_[int(13)] = d_40_v3_
                nw17_[int(14)] = d_40_v3_
                nw17_[int(15)] = d_40_v3_
                nw17_[int(16)] = d_40_v3_
                nw17_[int(17)] = d_40_v3_
                nw17_[int(18)] = d_40_v3_
                nw17_[int(19)] = d_40_v3_
                nw17_[int(20)] = d_40_v3_
                nw17_[int(21)] = (d_40_v3_ if not(p3) else d_40_v3_)
                nw17_[int(22)] = d_40_v3_
                nw17_[int(23)] = d_40_v3_
                nw17_[int(24)] = ((d_110_v52_).cf39 if d_107_v49_ else d_40_v3_)
                nw17_[int(25)] = d_40_v3_
                d_111_v53_ = nw17_
                index17_ = default__.safeIndex(84, (d_111_v53_).length(0))
                (d_111_v53_)[index17_] = d_40_v3_
                d_112_v54_: D8
                d_112_v54_ = D8_DC23(d_40_v3_.f12)
                d_113_v55_: D9
                d_113_v55_ = D9_DC26()
                d_114_v56_: C1
                nw18_ = C1()
                nw18_.ctor__(_dafny.CodePoint('g'), d_48_v7_, 696)
                d_114_v56_ = nw18_
                d_115_v57_: _dafny.Set
                d_115_v57_ = _dafny.Set({d_114_v56_, d_114_v56_})
                d_116_v58_: _dafny.Map
                d_116_v58_ = _dafny.Map({d_40_v3_.f12: p3})
                index18_ = default__.safeIndex(84, (d_111_v53_).length(0))
                rhs11_ = d_40_v3_
                rhs12_ = (d_40_v3_ if d_107_v49_ else d_40_v3_)
                rhs13_ = d_48_v7_
                rhs14_ = D8_DC23(default__.safeModuloInt(p2, len(d_115_v57_)))
                rhs15_ = default__.fm27(d_116_v58_, p3, (d_35_v0_) + ((d_45_v6_)[default__.safeIndex(44, (d_45_v6_).length(0))]), globalState)
                lhs10_ = d_111_v53_
                lhs11_ = default__.safeIndex(84, (d_111_v53_).length(0))
                d_40_v3_ = rhs11_
                lhs10_[lhs11_] = rhs12_
                d_48_v7_ = rhs13_
                d_112_v54_ = rhs14_
                d_113_v55_ = rhs15_
                d_35_v0_ = _dafny.SeqWithoutIsStrInference([(d_114_v56_).f15 for d_117_i7_ in range(default__.abs(49))])
            d_118_v59_: _dafny.MultiSet
            d_118_v59_ = _dafny.MultiSet([p2])
            d_119_v60_: _dafny.Map
            d_119_v60_ = _dafny.Map({p3: p1})
            r0 = default__.safeDivisionInt(default__.safeDivisionInt(d_40_v3_.f12, -35), ((d_118_v59_)[(0) - (p1)] if ((0) - (p1)) in (d_118_v59_) else len(d_119_v60_)))
        elif True:
            d_120_v61_: bool
            d_120_v61_ = True
            d_121_v62_: _dafny.MultiSet
            d_121_v62_ = _dafny.MultiSet([p3])
            d_122_v63_: _dafny.MultiSet
            d_122_v63_ = _dafny.MultiSet([p1, (d_121_v62_).cardinality])
            d_120_v61_ = (_dafny.MultiSet([default__.fm0(864, p1, globalState)])) == ((d_122_v63_) | (d_122_v63_))
            d_123_v64_: D10
            d_123_v64_ = D10_DC30(p0)
            (globalState).f5 = ((d_123_v64_).cf40) - ((p0) * (p1))
            d_124_v65_: _dafny.Array
            nw19_ = _dafny.Array(None, 24)
            nw19_[int(0)] = d_120_v61_
            nw19_[int(1)] = d_120_v61_
            nw19_[int(2)] = p3
            nw19_[int(3)] = d_120_v61_
            nw19_[int(4)] = d_120_v61_
            nw19_[int(5)] = True
            nw19_[int(6)] = False
            nw19_[int(7)] = not(d_120_v61_)
            nw19_[int(8)] = not(d_120_v61_)
            nw19_[int(9)] = p3
            nw19_[int(10)] = d_120_v61_
            nw19_[int(11)] = d_120_v61_
            nw19_[int(12)] = True
            nw19_[int(13)] = p3
            nw19_[int(14)] = p3
            nw19_[int(15)] = p3
            nw19_[int(16)] = p3
            nw19_[int(17)] = d_120_v61_
            nw19_[int(18)] = p3
            nw19_[int(19)] = False
            nw19_[int(20)] = p3
            nw19_[int(21)] = d_120_v61_
            nw19_[int(22)] = (D5_DC13(d_120_v61_, p3)).cf19
            nw19_[int(23)] = p3
            d_124_v65_ = nw19_
            d_125_v66_: _dafny.Map
            d_125_v66_ = _dafny.Map({d_120_v61_: d_124_v65_})
            d_126_v67_: D11
            d_126_v67_ = D11_DC32(d_125_v66_)
            d_125_v66_ = ((d_126_v67_).cf42) | ((d_125_v66_) | (d_125_v66_))
            (globalState).f5 = (p2 if p3 else (p2) * (p1))
            d_127_v68_: C4
            nw20_ = C4()
            nw20_.ctor__(p1, p0)
            d_127_v68_ = nw20_
        d_128_v69_: _dafny.Set
        d_128_v69_ = _dafny.Set({p3})
        d_129_v70_: _dafny.Seq
        d_129_v70_ = _dafny.SeqWithoutIsStrInference([p2, p0])
        d_130_v71_: _dafny.Array
        nw21_ = _dafny.Array(None, 17)
        nw21_[int(0)] = (646) - (p1)
        nw21_[int(1)] = p2
        nw21_[int(2)] = (p1) + (len(d_128_v69_))
        nw21_[int(3)] = (p1 if False else len(_dafny.SeqWithoutIsStrInference([p0])))
        nw21_[int(4)] = p0
        nw21_[int(5)] = p2
        nw21_[int(6)] = p1
        nw21_[int(7)] = (p1) + (p1)
        nw21_[int(8)] = p1
        nw21_[int(9)] = len(d_129_v70_)
        nw21_[int(10)] = p1
        nw21_[int(11)] = len(_dafny.SeqWithoutIsStrInference([p3, p3, p3]))
        nw21_[int(12)] = default__.safeDivisionInt(652, p0)
        nw21_[int(13)] = p0
        nw21_[int(14)] = p0
        nw21_[int(15)] = p2
        nw21_[int(16)] = default__.safeModuloInt(p2, p2)
        d_130_v71_ = nw21_
        d_130_v71_ = d_130_v71_
        d_131_v72_: bool
        d_131_v72_ = False
        d_132_v73_: str
        d_132_v73_ = _dafny.CodePoint('y')
        d_133_v74_: C1
        nw22_ = C1()
        nw22_.ctor__(d_132_v73_, d_132_v73_, p2)
        d_133_v74_ = nw22_
        d_134_v75_: _dafny.Seq
        d_134_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        rhs16_ = (d_131_v72_ if d_131_v72_ else ((0) - (len(d_134_v75_))) > (p1))
        rhs17_ = (d_131_v72_ if p3 else d_131_v72_)
        rhs18_ = d_131_v72_
        rhs19_ = d_133_v74_
        rhs20_ = p2
        lhs12_ = globalState
        d_131_v72_ = rhs16_
        d_131_v72_ = rhs17_
        d_131_v72_ = rhs18_
        d_133_v74_ = rhs19_
        lhs12_.f5 = rhs20_
        d_135_v76_: _dafny.Array
        def lambda4_(d_136_v75_, d_137_p1_, d_138_v74_):
            def lambda5_(d_139_i8_):
                return (-476) <= (len((d_136_v75_).set(default__.safeIndex((0) - (d_137_p1_), len(d_136_v75_)), (d_138_v74_).f15)))

            return lambda5_

        init2_ = lambda4_(d_134_v75_, p1, d_133_v74_)
        nw23_ = _dafny.Array(None, 21)
        for i0_2_ in range(nw23_.length(0)):
            nw23_[i0_2_] = init2_(i0_2_)
        d_135_v76_ = nw23_
        index19_ = default__.safeIndex(94, (d_135_v76_).length(0))
        (d_135_v76_)[index19_] = p3
        d_140_v77_: _dafny.Map
        d_140_v77_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_133_v74_).f15 for d_141_i9_ in range(default__.abs(-649))])): d_134_v75_})
        d_142_v78_: D1
        d_142_v78_ = D1_DC4(d_130_v71_)
        d_143_v79_: _dafny.Set
        d_143_v79_ = _dafny.Set({(d_142_v78_).cf7, d_130_v71_})
        index20_ = default__.safeIndex(94, (d_135_v76_).length(0))
        rhs21_ = default__.fm1(p0, d_131_v72_, (d_129_v70_)[default__.safeIndex(p0, len(d_129_v70_))], p3, globalState)
        rhs22_ = default__.fm0(default__.safeModuloInt(len(((d_140_v77_)[-120] if (-120) in (d_140_v77_) else d_134_v75_)), p2), len((d_143_v79_ if p3 else d_143_v79_)), globalState)
        lhs13_ = d_135_v76_
        lhs14_ = default__.safeIndex(94, (d_135_v76_).length(0))
        lhs13_[lhs14_] = rhs21_
        r0 = rhs22_
        d_144_v80_: C4
        nw24_ = C4()
        nw24_.ctor__((-263) + (p0), p2)
        d_144_v80_ = nw24_
        d_131_v72_ = True
        r0 = default__.safeModuloInt(d_144_v80_.f11, (d_144_v80_).f10)
        d_145_v81_: _dafny.Map
        d_145_v81_ = _dafny.Map({p3: 821})
        r1 = d_145_v81_
        d_146_v82_: _dafny.Map
        d_146_v82_ = _dafny.Map({d_144_v80_.f11: p0})
        d_147_v83_: _dafny.Seq
        d_147_v83_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_148_i10_ in range(default__.abs(347))]), d_134_v75_, (((d_134_v75_).set(default__.safeIndex((0) - (p1), len(d_134_v75_)), d_132_v73_)) + (d_134_v75_)).set(default__.safeIndex(len(d_146_v82_), len(((d_134_v75_).set(default__.safeIndex((0) - (p1), len(d_134_v75_)), d_132_v73_)) + (d_134_v75_))), d_132_v73_)])
        r2 = (d_147_v83_)[default__.safeIndex(-890, len(d_147_v83_))]
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_149_v0_: int
        d_149_v0_ = 172
        d_150_v1_: _dafny.Seq
        d_150_v1_ = _dafny.SeqWithoutIsStrInference([d_149_v0_])
        d_151_v2_: _dafny.Map
        d_151_v2_ = _dafny.Map({d_149_v0_: d_150_v1_})
        d_152_v3_: _dafny.MultiSet
        d_152_v3_ = _dafny.MultiSet([d_149_v0_])
        d_153_v4_: str
        d_153_v4_ = _dafny.CodePoint('s')
        d_154_v5_: _dafny.Seq
        d_154_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
        d_155_v6_: _dafny.Array
        nw25_ = _dafny.Array(None, 11)
        nw25_[int(0)] = d_153_v4_
        nw25_[int(1)] = d_153_v4_
        nw25_[int(2)] = d_153_v4_
        nw25_[int(3)] = d_153_v4_
        nw25_[int(4)] = d_153_v4_
        nw25_[int(5)] = d_153_v4_
        nw25_[int(6)] = d_153_v4_
        nw25_[int(7)] = d_153_v4_
        nw25_[int(8)] = d_153_v4_
        nw25_[int(9)] = (d_154_v5_)[default__.safeIndex(d_149_v0_, len(d_154_v5_))]
        nw25_[int(10)] = d_153_v4_
        d_155_v6_ = nw25_
        d_156_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__((_dafny.MultiSet(((d_151_v2_)[len(d_150_v1_)] if (len(d_150_v1_)) in (d_151_v2_) else _dafny.SeqWithoutIsStrInference([138])))) - (d_152_v3_), True, False, True, 841, 338, True, False, -192, d_155_v6_)
        d_156_globalState_ = nw26_
        d_157_v7_: _dafny.Map
        d_157_v7_ = _dafny.Map({d_149_v0_: 447})
        d_158_v8_: bool
        d_158_v8_ = True
        d_159_v9_: int
        d_160_v10_: _dafny.Map
        d_161_v11_: _dafny.Seq
        out0_: int
        out1_: _dafny.Map
        out2_: _dafny.Seq
        out0_, out1_, out2_ = default__.m0(((d_157_v7_)[49] if (49) in (d_157_v7_) else d_149_v0_), d_149_v0_, d_149_v0_, d_158_v8_, d_156_globalState_)
        d_159_v9_ = out0_
        d_160_v10_ = out1_
        d_161_v11_ = out2_
        d_158_v8_ = d_158_v8_
        hi0_ = d_149_v0_
        for d_162_i0_ in range(d_159_v9_, hi0_):
            d_163_v12_: _dafny.Map
            d_163_v12_ = _dafny.Map({(-269) - (default__.fm0(d_159_v9_, d_159_v9_, d_156_globalState_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytxvk"))})
            d_154_v5_ = ((d_163_v12_)[(0) - (d_159_v9_)] if ((0) - (d_159_v9_)) in (d_163_v12_) else d_161_v11_)
            d_158_v8_ = d_158_v8_
            d_164_v13_: _dafny.Seq
            d_164_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrlcvxj")), d_154_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enoyjic"))])
            d_165_v14_: _dafny.Seq
            d_165_v14_ = _dafny.SeqWithoutIsStrInference([d_158_v8_])
            d_166_v15_: D0
            d_166_v15_ = D0_DC0(d_158_v8_)
            d_167_v16_: D0
            d_167_v16_ = D0_DC0((d_166_v15_).cf0)
            rhs23_ = (d_164_v13_) + ((d_164_v13_) + (d_164_v13_))
            rhs24_ = not(d_158_v8_)
            rhs25_ = not ((len(d_165_v14_)) > (d_162_i0_)) or (d_158_v8_)
            rhs26_ = d_162_i0_
            rhs27_ = (d_166_v15_).cf0
            d_164_v13_ = rhs23_
            d_158_v8_ = rhs24_
            d_158_v8_ = rhs25_
            d_149_v0_ = rhs26_
            d_158_v8_ = rhs27_
            d_168_v17_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Set({}), 29)
            d_168_v17_ = nw27_
            d_168_v17_ = d_168_v17_
        d_169_v18_: _dafny.Array
        nw28_ = _dafny.Array(int(0), 20)
        d_169_v18_ = nw28_
        d_170_v19_: _dafny.Array
        nw29_ = _dafny.Array(None, 14)
        nw29_[int(0)] = d_169_v18_
        nw29_[int(1)] = d_169_v18_
        nw29_[int(2)] = (D1_DC4(d_169_v18_)).cf7
        nw29_[int(3)] = d_169_v18_
        nw29_[int(4)] = d_169_v18_
        nw29_[int(5)] = d_169_v18_
        nw29_[int(6)] = d_169_v18_
        nw29_[int(7)] = d_169_v18_
        nw29_[int(8)] = d_169_v18_
        nw29_[int(9)] = d_169_v18_
        nw29_[int(10)] = d_169_v18_
        nw29_[int(11)] = d_169_v18_
        nw29_[int(12)] = d_169_v18_
        nw29_[int(13)] = d_169_v18_
        d_170_v19_ = nw29_
        index21_ = default__.safeIndex(329, (d_170_v19_).length(0))
        (d_170_v19_)[index21_] = d_169_v18_
        index22_ = default__.safeIndex(329, (d_170_v19_).length(0))
        (d_170_v19_)[index22_] = d_169_v18_
        d_171_v20_: _dafny.Set
        d_171_v20_ = _dafny.Set({-144})
        d_172_v21_: int
        d_173_v22_: _dafny.Map
        d_174_v23_: _dafny.Seq
        out3_: int
        out4_: _dafny.Map
        out5_: _dafny.Seq
        out3_, out4_, out5_ = default__.m0(d_149_v0_, d_159_v9_, (0) - (((d_150_v1_)[default__.safeIndex((0) - (len(d_171_v20_)), len(d_150_v1_))] if True else len(d_154_v5_))), default__.fm1((0) - (d_149_v0_), d_158_v8_, -267, not(True), d_156_globalState_), d_156_globalState_)
        d_172_v21_ = out3_
        d_173_v22_ = out4_
        d_174_v23_ = out5_
        d_175_v24_: _dafny.Array
        def lambda6_(d_176_v0_):
            def lambda7_(d_177_i1_):
                return D2_DC7(d_176_v0_)

            return lambda7_

        init3_ = lambda6_(d_149_v0_)
        nw30_ = _dafny.Array(None, 5)
        for i0_3_ in range(nw30_.length(0)):
            nw30_[i0_3_] = init3_(i0_3_)
        d_175_v24_ = nw30_
        d_178_v25_: C3
        nw31_ = C3()
        nw31_.ctor__(d_175_v24_, _dafny.CodePoint('x'), (0) - (default__.safeDivisionInt(d_172_v21_, d_172_v21_)))
        d_178_v25_ = nw31_
        d_179_v26_: _dafny.Array
        nw32_ = _dafny.Array(None, 23)
        d_179_v26_ = nw32_
        d_180_v27_: C1
        nw33_ = C1()
        nw33_.ctor__(d_153_v4_, d_153_v4_, d_159_v9_)
        d_180_v27_ = nw33_
        d_181_v28_: _dafny.Map
        d_181_v28_ = _dafny.Map({d_149_v0_: d_158_v8_})
        d_182_v29_: _dafny.Map
        d_182_v29_ = _dafny.Map({d_181_v28_: (d_180_v27_).f15})
        d_183_v30_: D9
        d_183_v30_ = D9_DC27(d_180_v27_, (d_152_v3_).cardinality, len(d_154_v5_), -208, d_182_v29_)
        index23_ = default__.safeIndex(747, (d_179_v26_).length(0))
        (d_179_v26_)[index23_] = (d_183_v30_).cf33
        index24_ = default__.safeIndex(747, (d_179_v26_).length(0))
        nw34_ = C1()
        nw34_.ctor__(d_153_v4_, (d_180_v27_).f15, (len(d_161_v11_)) - (d_149_v0_))
        (d_179_v26_)[index24_] = nw34_
        index25_ = default__.safeIndex(190, (d_169_v18_).length(0))
        (d_169_v18_)[index25_] = d_149_v0_
        index26_ = default__.safeIndex(190, (d_169_v18_).length(0))
        (d_169_v18_)[index26_] = (0) - ((d_178_v25_).fm4(d_156_globalState_))
        d_184_v31_: _dafny.Array
        nw35_ = _dafny.Array(None, 24)
        d_184_v31_ = nw35_
        index27_ = default__.safeIndex(244, (d_184_v31_).length(0))
        (d_184_v31_)[index27_] = d_178_v25_
        index28_ = default__.safeIndex(244, (d_184_v31_).length(0))
        (d_184_v31_)[index28_] = (d_178_v25_ if d_158_v8_ else d_178_v25_)
        d_185_v32_: bool
        d_186_v33_: _dafny.Array
        d_187_v34_: _dafny.Map
        d_188_v35_: int
        out6_: bool
        out7_: _dafny.Array
        out8_: _dafny.Map
        out9_: int
        out6_, out7_, out8_, out9_ = (d_180_v27_).m2(d_156_globalState_)
        d_185_v32_ = out6_
        d_186_v33_ = out7_
        d_187_v34_ = out8_
        d_188_v35_ = out9_
        d_189_v36_: D5
        d_189_v36_ = D5_DC14(d_158_v8_)
        if (d_189_v36_).cf20:
            d_190_v37_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Set({}), 12)
            d_190_v37_ = nw36_
            d_191_v38_: T0
            nw37_ = C2()
            nw37_.ctor__(d_185_v32_, (d_178_v25_).fm7(d_156_globalState_))
            d_191_v38_ = nw37_
            d_192_v39_: _dafny.Map
            d_192_v39_ = _dafny.Map({d_190_v37_: d_191_v38_})
            d_192_v39_ = (d_192_v39_).set(d_190_v37_, d_191_v38_)
            d_193_v40_: _dafny.Array
            nw38_ = _dafny.Array(None, 5)
            nw38_[int(0)] = not(d_185_v32_)
            nw38_[int(1)] = not(d_158_v8_)
            nw38_[int(2)] = d_185_v32_
            nw38_[int(3)] = ((_dafny.MultiSet([d_188_v35_])).cardinality) < (d_172_v21_)
            nw38_[int(4)] = d_158_v8_
            d_193_v40_ = nw38_
            index29_ = default__.safeIndex(953, (d_193_v40_).length(0))
            (d_193_v40_)[index29_] = d_158_v8_
            d_194_v41_: D2
            d_194_v41_ = D2_DC8(((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]) >= (d_159_v9_), ((d_150_v1_)[default__.safeIndex(d_172_v21_, len(d_150_v1_))] if not(d_185_v32_) else (d_152_v3_).cardinality), (_dafny.SeqWithoutIsStrInference([(d_180_v27_).f15, _dafny.CodePoint('n')])) <= (d_174_v23_))
            d_195_v42_: _dafny.Map
            d_195_v42_ = _dafny.Map({D1_DC4((d_170_v19_)[default__.safeIndex(329, (d_170_v19_).length(0))]): d_186_v33_})
            d_196_v43_: _dafny.Map
            d_196_v43_ = _dafny.Map({d_185_v32_: _dafny.Set({-880, d_172_v21_})})
            d_197_v44_: _dafny.Array
            nw39_ = _dafny.Array(None, 12)
            nw39_[int(0)] = (_dafny.Set({d_188_v35_})) - (d_171_v20_)
            nw39_[int(1)] = default__.fm13(False, d_158_v8_, d_156_globalState_)
            nw39_[int(2)] = d_171_v20_
            nw39_[int(3)] = (_dafny.Set({(0) - (d_172_v21_)})).intersection(d_171_v20_)
            nw39_[int(4)] = _dafny.Set({d_188_v35_})
            nw39_[int(5)] = d_171_v20_
            nw39_[int(6)] = (_dafny.Set({len(d_195_v42_)})) - (((d_196_v43_)[d_185_v32_] if (d_185_v32_) in (d_196_v43_) else d_171_v20_))
            nw39_[int(7)] = _dafny.Set({d_191_v38_.f12})
            nw39_[int(8)] = (d_171_v20_) | (d_171_v20_)
            nw39_[int(9)] = (d_171_v20_) - (d_171_v20_)
            nw39_[int(10)] = _dafny.Set({(d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]})
            nw39_[int(11)] = d_171_v20_
            d_197_v44_ = nw39_
            index30_ = default__.safeIndex(177, (d_197_v44_).length(0))
            (d_197_v44_)[index30_] = _dafny.Set({(d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))], (0) - (d_149_v0_)})
            index31_ = default__.safeIndex(953, (d_193_v40_).length(0))
            index32_ = default__.safeIndex(177, (d_197_v44_).length(0))
            def iife5_(_pat_let0_0):
                def iife6_(d_198_dt__update__tmp_h0_):
                    def iife7_(_pat_let1_0):
                        def iife8_(d_199_dt__update_hcf13_h0_):
                            def iife9_(_pat_let2_0):
                                def iife10_(d_200_dt__update_hcf11_h0_):
                                    return D2_DC8(d_200_dt__update_hcf11_h0_, (d_198_dt__update__tmp_h0_).cf12, d_199_dt__update_hcf13_h0_)
                                return iife10_(_pat_let2_0)
                            return iife9_(False)
                        return iife8_(_pat_let1_0)
                    return iife7_(True)
                return iife6_(_pat_let0_0)
            rhs28_ = d_158_v8_
            rhs29_ = (d_178_v25_).fm7(d_156_globalState_)
            rhs30_ = iife5_(d_194_v41_)
            rhs31_ = d_171_v20_
            lhs15_ = d_193_v40_
            lhs16_ = default__.safeIndex(953, (d_193_v40_).length(0))
            lhs17_ = d_197_v44_
            lhs18_ = default__.safeIndex(177, (d_197_v44_).length(0))
            lhs15_[lhs16_] = rhs28_
            d_159_v9_ = rhs29_
            d_194_v41_ = rhs30_
            lhs17_[lhs18_] = rhs31_
            d_201_v45_: _dafny.Map
            d_201_v45_ = _dafny.Map({d_158_v8_: True})
            d_202_v46_: _dafny.MultiSet
            d_202_v46_ = _dafny.MultiSet([d_158_v8_])
            d_203_v47_: _dafny.Seq
            d_203_v47_ = _dafny.SeqWithoutIsStrInference([d_202_v46_])
            d_204_v48_: D5
            d_204_v48_ = D5_DC12(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_185_v32_, not(d_185_v32_)])]))
            d_205_v49_: _dafny.Seq
            d_205_v49_ = _dafny.SeqWithoutIsStrInference([D5_DC12(d_203_v47_), D5_DC12(d_203_v47_), d_204_v48_, d_204_v48_, D5_DC12(d_203_v47_)])
            d_206_v50_: D6
            d_206_v50_ = D6_DC15((d_180_v27_).fm5(((d_201_v45_)[not(False)] if (not(False)) in (d_201_v45_) else (d_193_v40_)[default__.safeIndex(953, (d_193_v40_).length(0))]), d_158_v8_, d_185_v32_, default__.fm1(d_188_v35_, True, (d_150_v1_)[default__.safeIndex(len(d_205_v49_), len(d_150_v1_))], d_185_v32_, d_156_globalState_), d_156_globalState_))
            pat_let_tv0_ = d_150_v1_
            def iife11_(_pat_let3_0):
                def iife12_(d_207_dt__update__tmp_h1_):
                    def iife13_(_pat_let4_0):
                        def iife14_(d_208_dt__update_hcf21_h0_):
                            return D6_DC15(d_208_dt__update_hcf21_h0_)
                        return iife14_(_pat_let4_0)
                    return iife13_(pat_let_tv0_)
                return iife12_(_pat_let3_0)
            d_206_v50_ = iife11_(D6_DC15(d_150_v1_))
            d_209_v51_: _dafny.Array
            def lambda8_(d_210_v40_, d_211_v11_, d_212_v23_):
                def lambda9_(d_213_i2_):
                    return (d_211_v11_ if (d_210_v40_)[default__.safeIndex(953, (d_210_v40_).length(0))] else d_212_v23_)

                return lambda9_

            init4_ = lambda8_(d_193_v40_, d_161_v11_, d_174_v23_)
            nw40_ = _dafny.Array(None, 24)
            for i0_4_ in range(nw40_.length(0)):
                nw40_[i0_4_] = init4_(i0_4_)
            d_209_v51_ = nw40_
            d_209_v51_ = d_209_v51_
            d_214_v52_: _dafny.Set
            d_214_v52_ = _dafny.Set({True, (d_193_v40_)[default__.safeIndex(953, (d_193_v40_).length(0))], d_185_v32_})
            if default__.fm1((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))], default__.fm1(d_159_v9_, ((d_201_v45_)[d_185_v32_] if (d_185_v32_) in (d_201_v45_) else d_185_v32_), d_191_v38_.f12, default__.fm1(125, (d_193_v40_)[default__.safeIndex(953, (d_193_v40_).length(0))], len(d_214_v52_), d_185_v32_, d_156_globalState_), d_156_globalState_), (d_178_v25_).fm7(d_156_globalState_), d_185_v32_, d_156_globalState_):
                index33_ = default__.safeIndex(953, (d_193_v40_).length(0))
                (d_193_v40_)[index33_] = not (d_158_v8_) or ((_dafny.MultiSet(d_150_v1_)).issubset(default__.fm21(d_149_v0_, d_172_v21_, d_153_v4_, (0) - (d_188_v35_), d_156_globalState_)))
                d_158_v8_ = d_185_v32_
                d_215_v53_: bool
                d_216_v54_: _dafny.Array
                d_217_v55_: _dafny.Map
                d_218_v56_: int
                out10_: bool
                out11_: _dafny.Array
                out12_: _dafny.Map
                out13_: int
                out10_, out11_, out12_, out13_ = (d_178_v25_).m2(d_156_globalState_)
                d_215_v53_ = out10_
                d_216_v54_ = out11_
                d_217_v55_ = out12_
                d_218_v56_ = out13_
                d_219_v57_: _dafny.Map
                d_219_v57_ = _dafny.Map({d_191_v38_: d_158_v8_})
                d_220_v58_: C2
                nw41_ = C2()
                nw41_.ctor__(d_158_v8_, len((d_219_v57_) | (d_219_v57_)))
                d_220_v58_ = nw41_
                d_193_v40_ = d_193_v40_
            elif True:
                (d_156_globalState_).f5 = (0) - ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))])
                d_221_v59_: _dafny.Map
                d_221_v59_ = _dafny.Map({(d_178_v25_).fm7(d_156_globalState_): d_154_v5_})
                d_222_v60_: _dafny.Set
                d_222_v60_ = _dafny.Set({d_180_v27_})
                d_221_v59_ = (d_221_v59_).set((len(d_222_v60_)) * (627), d_161_v11_)
                d_185_v32_ = not ((d_193_v40_)[default__.safeIndex(953, (d_193_v40_).length(0))]) or ((d_193_v40_)[default__.safeIndex(953, (d_193_v40_).length(0))])
                index34_ = default__.safeIndex(190, (d_169_v18_).length(0))
                (d_169_v18_)[index34_] = default__.safeModuloInt((0) - (default__.safeDivisionInt(d_149_v0_, 241)), default__.fm0(((d_157_v7_)[(d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]] if ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]) in (d_157_v7_) else d_149_v0_), 658, d_156_globalState_))
                d_174_v23_ = _dafny.SeqWithoutIsStrInference([(d_161_v11_)[default__.safeIndex(d_159_v9_, len(d_161_v11_))] for d_223_i3_ in range(default__.abs(-635))])
        elif True:
            d_224_v61_: C3
            nw42_ = C3()
            nw42_.ctor__(d_175_v24_, _dafny.CodePoint('j'), d_149_v0_)
            d_224_v61_ = nw42_
            d_225_v62_: _dafny.Seq
            d_225_v62_ = _dafny.SeqWithoutIsStrInference([d_185_v32_])
            d_225_v62_ = (d_225_v62_).set(default__.safeIndex(d_172_v21_, len(d_225_v62_)), (d_185_v32_ if True else default__.fm1(d_149_v0_, d_158_v8_, d_149_v0_, d_185_v32_, d_156_globalState_)))
            d_226_v63_: _dafny.Map
            d_226_v63_ = _dafny.Map({d_186_v33_: d_172_v21_})
            if ((d_226_v63_).set(d_169_v18_, d_172_v21_)) != (d_226_v63_):
                d_227_v64_: _dafny.Map
                d_227_v64_ = _dafny.Map({d_185_v32_: d_179_v26_})
                d_228_v65_: _dafny.Array
                nw43_ = _dafny.Array(None, 3)
                nw43_[int(0)] = d_171_v20_
                nw43_[int(1)] = d_171_v20_
                nw43_[int(2)] = _dafny.Set({(default__.fm19(d_149_v0_, d_156_globalState_)).cardinality, (0) - ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]), d_172_v21_})
                d_228_v65_ = nw43_
                index35_ = default__.safeIndex(588, (d_228_v65_).length(0))
                (d_228_v65_)[index35_] = d_171_v20_
                d_229_v66_: D2
                d_229_v66_ = D2_DC9(d_149_v0_)
                index36_ = default__.safeIndex(588, (d_228_v65_).length(0))
                rhs32_ = (d_227_v64_).set((d_152_v3_).issubset(d_152_v3_), d_179_v26_)
                rhs33_ = _dafny.Set({(len((d_224_v61_).fm5(d_185_v32_, d_185_v32_, d_185_v32_, d_158_v8_, d_156_globalState_))) + ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]), (d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]})
                rhs34_ = d_229_v66_
                rhs35_ = d_174_v23_
                lhs19_ = d_228_v65_
                lhs20_ = default__.safeIndex(588, (d_228_v65_).length(0))
                d_227_v64_ = rhs32_
                lhs19_[lhs20_] = rhs33_
                d_229_v66_ = rhs34_
                d_174_v23_ = rhs35_
                d_230_v67_: C2
                nw44_ = C2()
                nw44_.ctor__(d_158_v8_, d_159_v9_)
                d_230_v67_ = nw44_
                d_231_v68_: _dafny.Array
                nw45_ = _dafny.Array(None, 11)
                d_231_v68_ = nw45_
                d_232_v69_: T2
                nw46_ = C3()
                nw46_.ctor__((d_224_v61_).f14, _dafny.CodePoint('r'), 268)
                d_232_v69_ = nw46_
                index37_ = default__.safeIndex(729, (d_231_v68_).length(0))
                (d_231_v68_)[index37_] = d_232_v69_
                index38_ = default__.safeIndex(729, (d_231_v68_).length(0))
                rhs36_ = ((d_225_v62_)[default__.safeIndex(d_232_v69_.f12, len(d_225_v62_))]) == (d_185_v32_)
                rhs37_ = (d_232_v69_ if d_158_v8_ else d_232_v69_)
                lhs21_ = d_231_v68_
                lhs22_ = default__.safeIndex(729, (d_231_v68_).length(0))
                d_158_v8_ = rhs36_
                lhs21_[lhs22_] = rhs37_
                d_233_v70_: bool
                d_234_v71_: _dafny.Array
                d_235_v72_: _dafny.Map
                d_236_v73_: int
                out14_: bool
                out15_: _dafny.Array
                out16_: _dafny.Map
                out17_: int
                out14_, out15_, out16_, out17_ = (d_178_v25_).m2(d_156_globalState_)
                d_233_v70_ = out14_
                d_234_v71_ = out15_
                d_235_v72_ = out16_
                d_236_v73_ = out17_
                d_174_v23_ = d_154_v5_
            elif True:
                d_237_v74_: C4
                nw47_ = C4()
                nw47_.ctor__((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))], d_159_v9_)
                d_237_v74_ = nw47_
                d_238_v75_: _dafny.Seq
                d_238_v75_ = _dafny.SeqWithoutIsStrInference([d_237_v74_, d_237_v74_, d_237_v74_])
                d_239_v76_: _dafny.Map
                d_239_v76_ = _dafny.Map({len(d_173_v22_): d_238_v75_})
                d_172_v21_ = default__.safeDivisionInt(len(((d_239_v76_)[(d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]] if ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]) in (d_239_v76_) else d_238_v75_)), d_188_v35_)
                d_240_v77_: C0
                nw48_ = C0()
                nw48_.ctor__(d_181_v28_, d_188_v35_)
                d_240_v77_ = nw48_
                d_241_v78_: _dafny.Map
                d_241_v78_ = _dafny.Map({d_240_v77_: (d_154_v5_).set(default__.safeIndex(d_159_v9_, len(d_154_v5_)), (d_180_v27_).f15)})
                d_188_v35_ = len(((d_241_v78_) | (d_241_v78_) if False else (d_241_v78_) | (d_241_v78_)))
                d_158_v8_ = d_185_v32_
                d_158_v8_ = False
                index39_ = default__.safeIndex(329, (d_170_v19_).length(0))
                def lambda10_(d_242_v74_):
                    def lambda11_(d_243_i4_):
                        return default__.safeModuloInt(d_243_i4_, (d_242_v74_).f10)

                    return lambda11_

                init5_ = lambda10_(d_237_v74_)
                nw49_ = _dafny.Array(None, 6)
                for i0_5_ in range(nw49_.length(0)):
                    nw49_[i0_5_] = init5_(i0_5_)
                (d_170_v19_)[index39_] = nw49_
            d_244_v79_: D9
            d_244_v79_ = D9_DC26()
            d_244_v79_ = (d_244_v79_ if not(d_185_v32_) else d_244_v79_)
            d_245_v80_: int
            d_246_v81_: bool
            d_247_v82_: int
            out18_: int
            out19_: bool
            out20_: int
            out18_, out19_, out20_ = (d_178_v25_).m3(_dafny.MultiSet([False, not(d_185_v32_)]), d_156_globalState_)
            d_245_v80_ = out18_
            d_246_v81_ = out19_
            d_247_v82_ = out20_
        d_248_v83_: _dafny.MultiSet
        d_248_v83_ = _dafny.MultiSet([d_158_v8_])
        if default__.fm1(len(_dafny.Map({d_185_v32_: d_248_v83_})), d_158_v8_, 844, (d_172_v21_) < ((d_248_v83_).cardinality), d_156_globalState_):
            d_188_v35_ = 214
            d_181_v28_ = (d_181_v28_).set(((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))] if d_158_v8_ else (0) - ((d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))])), False)
            d_249_v84_: _dafny.Map
            d_249_v84_ = _dafny.Map({d_158_v8_: d_184_v31_})
            d_250_v85_: _dafny.Map
            d_250_v85_ = _dafny.Map({(d_181_v28_) | (default__.fm3(d_185_v32_, d_158_v8_, (d_178_v25_).fm7(d_156_globalState_), d_156_globalState_)): ((d_249_v84_)[d_158_v8_] if (d_158_v8_) in (d_249_v84_) else d_184_v31_)})
            d_251_v87_: _dafny.Map
            d_251_v87_ = _dafny.Map({default__.fm1(46, d_185_v32_, 959, d_158_v8_, d_156_globalState_): (d_180_v27_).f15})
            d_252_v88_: D0
            def iife15_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(601, 587):
                    d_253_v86_: int = compr_5_
                    if ((601) <= (d_253_v86_)) and ((d_253_v86_) < (587)):
                        coll5_[default__.safeDivisionInt(d_253_v86_, len(d_154_v5_))] = d_185_v32_
                return _dafny.Map(coll5_)
            d_252_v88_ = D0_DC2(d_158_v8_, iife15_()
, d_251_v87_, d_158_v8_)
            d_254_v89_: D0
            d_254_v89_ = D0_DC2(d_158_v8_, (d_252_v88_).cf3, (d_251_v87_).set(d_158_v8_, (d_180_v27_).f15), d_158_v8_)
            d_255_v90_: _dafny.Map
            d_255_v90_ = _dafny.Map({d_188_v35_: d_184_v31_})
            d_250_v85_ = _dafny.Map({(d_252_v88_).cf3: ((d_255_v90_)[702] if (702) in (d_255_v90_) else d_184_v31_)})
            d_158_v8_ = d_158_v8_
            d_256_v91_: _dafny.Set
            d_256_v91_ = _dafny.Set({d_185_v32_})
            d_257_v92_: _dafny.Seq
            d_257_v92_ = _dafny.SeqWithoutIsStrInference([d_185_v32_])
            d_258_v93_: _dafny.Seq
            d_258_v93_ = _dafny.SeqWithoutIsStrInference([d_257_v92_, _dafny.SeqWithoutIsStrInference([d_158_v8_, d_185_v32_, d_185_v32_]), d_257_v92_])
            d_259_v94_: D8
            d_259_v94_ = D8_DC24(d_256_v91_, (d_258_v93_)[default__.safeIndex(d_172_v21_, len(d_258_v93_))])
            d_260_v95_: _dafny.Set
            d_260_v95_ = _dafny.Set({default__.fm22(d_185_v32_, d_156_globalState_), d_259_v94_})
            d_261_v96_: _dafny.Map
            d_261_v96_ = _dafny.Map({(d_260_v95_).intersection(d_260_v95_): (len(d_150_v1_) if d_185_v32_ else d_188_v35_)})
            index40_ = default__.safeIndex(244, (d_184_v31_).length(0))
            rhs38_ = (d_261_v96_) | (d_261_v96_)
            rhs39_ = d_178_v25_
            rhs40_ = True
            lhs23_ = d_184_v31_
            lhs24_ = default__.safeIndex(244, (d_184_v31_).length(0))
            d_261_v96_ = rhs38_
            lhs23_[lhs24_] = rhs39_
            d_158_v8_ = rhs40_
        elif True:
            d_185_v32_ = d_158_v8_
            d_153_v4_ = d_153_v4_
            d_153_v4_ = (_dafny.CodePoint('m') if not(d_185_v32_) else (d_161_v11_)[default__.safeIndex(d_172_v21_, len(d_161_v11_))])
            (d_156_globalState_).f5 = d_149_v0_
            (d_156_globalState_).f5 = 578
        d_262_v97_: bool
        d_263_v98_: _dafny.Array
        d_264_v99_: _dafny.Map
        d_265_v100_: int
        out21_: bool
        out22_: _dafny.Array
        out23_: _dafny.Map
        out24_: int
        out21_, out22_, out23_, out24_ = (d_180_v27_).m2(d_156_globalState_)
        d_262_v97_ = out21_
        d_263_v98_ = out22_
        d_264_v99_ = out23_
        d_265_v100_ = out24_
        d_266_v101_: C3
        nw50_ = C3()
        nw50_.ctor__((d_178_v25_).f14, d_153_v4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))
        d_266_v101_ = nw50_
        d_267_i5_: int
        d_267_i5_ = 0
        with _dafny.label("0"):
            while True:
                with _dafny.c_label("0"):
                    if (d_267_i5_) >= (100):
                        raise _dafny.Break("0")
                    d_267_i5_ = (d_267_i5_) + (1)
                    (d_156_globalState_).f5 = d_172_v21_
                    if not (d_185_v32_) or (d_262_v97_):
                        d_268_v102_: bool
                        d_269_v103_: _dafny.Array
                        d_270_v104_: _dafny.Map
                        d_271_v105_: int
                        out25_: bool
                        out26_: _dafny.Array
                        out27_: _dafny.Map
                        out28_: int
                        out25_, out26_, out27_, out28_ = (d_180_v27_).m2(d_156_globalState_)
                        d_268_v102_ = out25_
                        d_269_v103_ = out26_
                        d_270_v104_ = out27_
                        d_271_v105_ = out28_
                        d_188_v35_ = ((d_178_v25_).fm4(d_156_globalState_) if d_262_v97_ else (d_159_v9_) - (d_159_v9_))
                        d_272_v106_: _dafny.Array
                        def lambda12_(d_273_v1_):
                            def lambda13_(d_274_i6_):
                                return (d_273_v1_) < (d_273_v1_)

                            return lambda13_

                        init6_ = lambda12_(d_150_v1_)
                        nw51_ = _dafny.Array(None, 4)
                        for i0_6_ in range(nw51_.length(0)):
                            nw51_[i0_6_] = init6_(i0_6_)
                        d_272_v106_ = nw51_
                        index41_ = default__.safeIndex(299, (d_272_v106_).length(0))
                        (d_272_v106_)[index41_] = d_262_v97_
                        index42_ = default__.safeIndex(299, (d_272_v106_).length(0))
                        rhs41_ = (d_170_v19_)[default__.safeIndex(329, (d_170_v19_).length(0))]
                        rhs42_ = not(d_262_v97_)
                        lhs25_ = d_272_v106_
                        lhs26_ = default__.safeIndex(299, (d_272_v106_).length(0))
                        d_269_v103_ = rhs41_
                        lhs25_[lhs26_] = rhs42_
                        index43_ = default__.safeIndex(190, (d_169_v18_).length(0))
                        rhs43_ = d_154_v5_
                        rhs44_ = d_188_v35_
                        rhs45_ = (d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]
                        rhs46_ = True
                        lhs27_ = d_169_v18_
                        lhs28_ = default__.safeIndex(190, (d_169_v18_).length(0))
                        d_174_v23_ = rhs43_
                        d_172_v21_ = rhs44_
                        lhs27_[lhs28_] = rhs45_
                        d_262_v97_ = rhs46_
                        d_275_v107_: C3
                        nw52_ = C3()
                        nw52_.ctor__((d_178_v25_).f14, (d_180_v27_).f15, (0) - ((d_271_v105_) - (d_188_v35_)))
                        d_275_v107_ = nw52_
                    elif True:
                        d_161_v11_ = ((d_174_v23_) + (d_154_v5_) if d_158_v8_ else (d_154_v5_) + (d_174_v23_))
                        index44_ = default__.safeIndex(190, (d_169_v18_).length(0))
                        (d_169_v18_)[index44_] = d_188_v35_
                        d_276_v108_: _dafny.Set
                        d_276_v108_ = _dafny.Set({d_153_v4_})
                        rhs47_ = ((_dafny.Set({(d_180_v27_).f15, d_153_v4_})) | (_dafny.Set({d_153_v4_}))) | (d_276_v108_)
                        rhs48_ = ((d_265_v100_) - (d_149_v0_)) != ((d_150_v1_)[default__.safeIndex(d_159_v9_, len(d_150_v1_))])
                        rhs49_ = d_159_v9_
                        rhs50_ = d_150_v1_
                        lhs29_ = d_156_globalState_
                        d_276_v108_ = rhs47_
                        d_262_v97_ = rhs48_
                        lhs29_.f5 = rhs49_
                        d_150_v1_ = rhs50_
                        d_277_v109_: _dafny.Seq
                        d_277_v109_ = (d_154_v5_) + (d_154_v5_)
                        d_277_v109_ = d_277_v109_
                        d_278_v110_: bool
                        d_279_v111_: _dafny.Array
                        d_280_v112_: _dafny.Map
                        d_281_v113_: int
                        out29_: bool
                        out30_: _dafny.Array
                        out31_: _dafny.Map
                        out32_: int
                        out29_, out30_, out31_, out32_ = (d_180_v27_).m2(d_156_globalState_)
                        d_278_v110_ = out29_
                        d_279_v111_ = out30_
                        d_280_v112_ = out31_
                        d_281_v113_ = out32_
                    if d_158_v8_:
                        d_282_v114_: D6
                        d_282_v114_ = D6_DC15((d_150_v1_).set(default__.safeIndex(d_188_v35_, len(d_150_v1_)), d_172_v21_))
                        d_282_v114_ = (d_282_v114_ if not(d_158_v8_) else d_282_v114_)
                        d_150_v1_ = (_dafny.SeqWithoutIsStrInference([len(d_171_v20_)])) + (d_150_v1_)
                        d_150_v1_ = (_dafny.SeqWithoutIsStrInference([d_265_v100_, (d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))], d_159_v9_, d_149_v0_])) + (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_149_v0_))]))
                        index45_ = default__.safeIndex(190, (d_169_v18_).length(0))
                        (d_169_v18_)[index45_] = 399
                        d_149_v0_ = d_172_v21_
                    elif True:
                        d_283_v115_: C4
                        nw53_ = C4()
                        nw53_.ctor__((d_159_v9_) + (d_188_v35_), default__.safeModuloInt(len(d_173_v22_), d_172_v21_))
                        d_283_v115_ = nw53_
                        d_284_v116_: bool
                        d_285_v117_: bool
                        d_286_v118_: int
                        out33_: bool
                        out34_: bool
                        out35_: int
                        out33_, out34_, out35_ = (d_266_v101_).m4(d_159_v9_, d_158_v8_, (d_154_v5_) + (_dafny.SeqWithoutIsStrInference([(d_180_v27_).f15 for d_287_i7_ in range(default__.abs(-669))])), d_156_globalState_)
                        d_284_v116_ = out33_
                        d_285_v117_ = out34_
                        d_286_v118_ = out35_
                        d_288_v119_: _dafny.Array
                        nw54_ = _dafny.Array(False, 16)
                        d_288_v119_ = nw54_
                        index46_ = default__.safeIndex(8, (d_288_v119_).length(0))
                        (d_288_v119_)[index46_] = (len(d_171_v20_)) != ((d_283_v115_).f10)
                        d_289_v120_: _dafny.Map
                        d_289_v120_ = _dafny.Map({d_149_v0_: _dafny.Set({d_172_v21_})})
                        index47_ = default__.safeIndex(8, (d_288_v119_).length(0))
                        (d_288_v119_)[index47_] = (d_171_v20_).issubset(((d_289_v120_)[(0) - (len(d_174_v23_))] if ((0) - (len(d_174_v23_))) in (d_289_v120_) else d_171_v20_))
                        d_290_v121_: int
                        d_291_v122_: bool
                        d_292_v123_: int
                        out36_: int
                        out37_: bool
                        out38_: int
                        out36_, out37_, out38_ = (d_266_v101_).m3(d_248_v83_, d_156_globalState_)
                        d_290_v121_ = out36_
                        d_291_v122_ = out37_
                        d_292_v123_ = out38_
                        index48_ = default__.safeIndex(8, (d_288_v119_).length(0))
                        (d_288_v119_)[index48_] = False
                    d_149_v0_ = (d_169_v18_)[default__.safeIndex(190, (d_169_v18_).length(0))]
                    pass
            pass
        rhs51_ = (0) - (d_159_v9_)
        rhs52_ = default__.fm1(len(d_160_v10_), d_262_v97_, d_265_v100_, d_262_v97_, d_156_globalState_)
        d_188_v35_ = rhs51_
        d_185_v32_ = rhs52_
        _dafny.print(_dafny.string_of(d_149_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v1_) == (_dafny.SeqWithoutIsStrInference([6, 399, 583, 399, 399]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v2_) == (_dafny.Map({172: _dafny.SeqWithoutIsStrInference([172])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v3_) == (_dafny.MultiSet([172]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_153_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_154_v5_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f0) == (_dafny.MultiSet([138]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_156_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_globalState_).f9)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v7_) == (_dafny.Map({172: 447}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_158_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_159_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v10_) == (_dafny.Map({True: 821}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v11_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v18_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[1])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[2])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[3])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[4])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[5])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[6])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[7])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[8])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[9])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[10])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[11])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[12])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_170_v19_)[13])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v20_) == (_dafny.Set({-144}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_172_v21_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v22_) == (_dafny.Map({False: 821}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v23_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v24_)[0]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v24_)[1]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v24_)[2]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v24_)[3]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_175_v24_)[4]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_178_v25_).f14)[0]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_178_v25_).f14)[1]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_178_v25_).f14)[2]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_178_v25_).f14)[3]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_178_v25_).f14)[4]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v26_)[11]).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v26_)[11]).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v26_)[11].f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v27_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_v28_) == (_dafny.Map({171: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v29_) == (_dafny.Map({_dafny.Map({171: True}): _dafny.CodePoint('s')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v30_).cf33).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v30_).cf33).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v30_).cf33.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v30_).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v30_).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v30_).cf36))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_183_v30_).cf37) == (_dafny.Map({_dafny.Map({171: True}): _dafny.CodePoint('s')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_184_v31_)[4]).f14)[0]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_184_v31_)[4]).f14)[1]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_184_v31_)[4]).f14)[2]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_184_v31_)[4]).f14)[3]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_184_v31_)[4]).f14)[4]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_184_v31_)[4]).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v31_)[4].f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_185_v32_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_186_v33_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_187_v34_) == (_dafny.Map({_dafny.CodePoint('b'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_188_v35_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_189_v36_).cf20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v83_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_262_v97_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v98_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v99_) == (_dafny.Map({_dafny.CodePoint('b'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_265_v100_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_266_v101_).f14)[0]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_266_v101_).f14)[1]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_266_v101_).f14)[2]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_266_v101_).f14)[3]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_266_v101_).f14)[4]).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_267_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf6 == __o.cf6
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

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC10(D3, NamedTuple('DC10', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({self.cf15.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(_dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC21(D7, NamedTuple('DC21', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC26(D9, NamedTuple('DC26', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf33', Any), ('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC28(D9, NamedTuple('DC28', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC30(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D10_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D10_DC31)

class D10_DC30(D10, NamedTuple('DC30', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC31(D10, NamedTuple('DC31', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC31({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC31) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC33(_dafny.Map({}), int(0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D11_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D11_DC35)

class D11_DC33(D11, NamedTuple('DC33', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC34(D11, NamedTuple('DC34', [('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC34({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC34) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC35(D11, NamedTuple('DC35', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC35({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC35) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC37()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D12_DC39)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)

class D12_DC37(D12, NamedTuple('DC37', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC38(D12, NamedTuple('DC38', [('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC39(D12, NamedTuple('DC39', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC39({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC39) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value

class T1:
    pass
    def m2(self, globalState):
        pass


class T2:
    pass

class GlobalState:
    def  __init__(self):
        self.f5: int = int(0)
        self._f0: _dafny.MultiSet = _dafny.MultiSet({})
        self._f1: bool = False
        self._f2: bool = False
        self._f3: bool = False
        self._f4: int = int(0)
        self._f6: bool = False
        self._f7: bool = False
        self._f8: int = int(0)
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9

    @property
    def f0(self):
        return self._f0
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

class C0(T2, T0):
    def  __init__(self):
        self._f12: int = int(0)
        self.f16: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f16, f12):
        (self).f16 = f16
        (self).f12 = f12

    def fm6(self, globalState):
        return D0_DC0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxlwa"))) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_293_i0_ in range(default__.abs(370))])))

    def fm8(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnadqpvyl"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yutm")) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slatllus"))))


class C1(T1, T0):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: str = _dafny.CodePoint('D')
        self._f15: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f15, f13, f12):
        (self)._f15 = f15
        (self)._f13 = f13
        (self).f12 = f12

    def fm4(self, globalState):
        return self.f12

    def fm5(self, p0, p1, p2, p3, globalState):
        if False:
            return _dafny.SeqWithoutIsStrInference([self.f12])
        elif True:
            return _dafny.SeqWithoutIsStrInference([self.f12 for d_294_i0_ in range(default__.abs(155))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([234, self.f12])

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_295_v0_: _dafny.Array
        def lambda14_(d_296_i0_):
            return (d_296_i0_) * (self.f12)

        init7_ = lambda14_
        nw55_ = _dafny.Array(None, 12)
        for i0_7_ in range(nw55_.length(0)):
            nw55_[i0_7_] = init7_(i0_7_)
        d_295_v0_ = nw55_
        index49_ = default__.safeIndex(216, (d_295_v0_).length(0))
        (d_295_v0_)[index49_] = self.f12
        index50_ = default__.safeIndex(216, (d_295_v0_).length(0))
        (d_295_v0_)[index50_] = self.f12
        d_297_v1_: bool
        d_297_v1_ = False
        d_298_v2_: D1
        d_298_v2_ = D1_DC5(d_297_v1_)
        d_298_v2_ = D1_DC5(d_297_v1_)
        if (d_297_v1_ if d_297_v1_ else d_297_v1_):
            d_299_v3_: _dafny.Set
            d_299_v3_ = _dafny.Set({d_297_v1_, d_297_v1_})
            d_300_v4_: _dafny.Map
            d_300_v4_ = _dafny.Map({len(d_299_v3_): False})
            d_301_v5_: T2
            nw56_ = C0()
            nw56_.ctor__(d_300_v4_, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])
            d_301_v5_ = nw56_
            d_302_v6_: _dafny.Seq
            d_302_v6_ = _dafny.SeqWithoutIsStrInference([d_297_v1_, d_297_v1_, d_297_v1_])
            d_303_v7_: _dafny.Seq
            d_303_v7_ = _dafny.SeqWithoutIsStrInference([d_295_v0_])
            rhs53_ = True
            rhs54_ = not (d_297_v1_) or ((d_301_v5_) in (_dafny.SeqWithoutIsStrInference([d_301_v5_])))
            rhs55_ = not (d_297_v1_) or (not((d_302_v6_)[default__.safeIndex((0) - (d_301_v5_.f12), len(d_302_v6_))]))
            rhs56_ = len((((d_303_v7_) + (d_303_v7_)) + (d_303_v7_)).set(default__.safeIndex(self.f12, len(((d_303_v7_) + (d_303_v7_)) + (d_303_v7_))), d_295_v0_))
            lhs30_ = globalState
            r0 = rhs53_
            r0 = rhs54_
            r0 = rhs55_
            lhs30_.f5 = rhs56_
            d_304_v8_: _dafny.Seq
            d_304_v8_ = _dafny.SeqWithoutIsStrInference([self.f12])
            d_305_v9_: _dafny.MultiSet
            d_305_v9_ = _dafny.MultiSet([d_301_v5_.f12])
            d_306_v10_: _dafny.Map
            d_306_v10_ = _dafny.Map({d_297_v1_: d_305_v9_})
            d_307_v11_: _dafny.Map
            d_307_v11_ = _dafny.Map({d_306_v10_: d_297_v1_})
            d_308_v12_: _dafny.Map
            d_308_v12_ = _dafny.Map({d_297_v1_: (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]})
            d_309_v13_: D2
            d_309_v13_ = D2_DC7(len(_dafny.Map({((d_305_v9_)[self.f12] if (self.f12) in (d_305_v9_) else (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]): ((d_308_v12_)[d_297_v1_] if (d_297_v1_) in (d_308_v12_) else self.f12)})))
            rhs57_ = d_297_v1_
            rhs58_ = d_301_v5_.f12
            rhs59_ = ((d_307_v11_)[d_306_v10_] if (d_306_v10_) in (d_307_v11_) else not((d_297_v1_) or (d_297_v1_)))
            rhs60_ = (d_304_v8_).set(default__.safeIndex(((d_309_v13_).cf10 if d_297_v1_ else -345), len(d_304_v8_)), d_301_v5_.f12)
            lhs31_ = globalState
            d_297_v1_ = rhs57_
            lhs31_.f5 = rhs58_
            r0 = rhs59_
            d_304_v8_ = rhs60_
            if False:
                rhs61_ = (d_297_v1_ if d_297_v1_ else (d_297_v1_) not in (d_302_v6_))
                rhs62_ = d_297_v1_
                d_297_v1_ = rhs61_
                r0 = rhs62_
                d_310_v14_: C0
                nw57_ = C0()
                nw57_.ctor__(d_300_v4_, (D2_DC9(-834)).cf14)
                d_310_v14_ = nw57_
                d_311_v15_: str
                d_311_v15_ = _dafny.CodePoint('n')
                d_311_v15_ = (self).f13
                d_312_v16_: _dafny.Seq
                d_312_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiotodfiq"))
                d_312_v16_ = ((d_312_v16_) + (d_312_v16_)) + (d_312_v16_)
                d_302_v6_ = (default__.fm9(not(d_297_v1_), globalState)) + (_dafny.SeqWithoutIsStrInference([d_297_v1_, True, d_297_v1_]))
            elif True:
                d_313_v17_: _dafny.Seq
                d_313_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
                d_314_v18_: _dafny.Seq
                d_314_v18_ = _dafny.SeqWithoutIsStrInference([d_313_v17_, _dafny.SeqWithoutIsStrInference([(self).f15 for d_315_i1_ in range(default__.abs(718))]), (d_313_v17_) + (d_313_v17_)])
                d_316_v19_: _dafny.Map
                d_316_v19_ = _dafny.Map({(0) - (d_301_v5_.f12): 210})
                d_313_v17_ = (d_314_v18_)[default__.safeIndex((440) + (((d_316_v19_)[len(d_313_v17_)] if (len(d_313_v17_)) in (d_316_v19_) else (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])), len(d_314_v18_))]
                d_317_v20_: C0
                nw58_ = C0()
                nw58_.ctor__(d_300_v4_, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])
                d_317_v20_ = nw58_
                d_317_v20_ = d_317_v20_
                d_318_v21_: _dafny.Array
                nw59_ = _dafny.Array(None, 22)
                d_318_v21_ = nw59_
                d_318_v21_ = d_318_v21_
                d_319_v22_: _dafny.MultiSet
                d_319_v22_ = _dafny.MultiSet([d_302_v6_])
                d_320_v23_: _dafny.Seq
                d_320_v23_ = _dafny.SeqWithoutIsStrInference([d_302_v6_])
                d_319_v22_ = _dafny.MultiSet(d_320_v23_)
                (d_317_v20_).f16 = d_317_v20_.f16
            d_297_v1_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbyqqyufh")))
            d_321_v24_: C0
            nw60_ = C0()
            nw60_.ctor__(_dafny.Map({self.f12: d_297_v1_}), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])
            d_321_v24_ = nw60_
            d_322_v25_: _dafny.MultiSet
            d_322_v25_ = _dafny.MultiSet([d_321_v24_, d_321_v24_, d_321_v24_])
            d_323_v26_: _dafny.Map
            d_323_v26_ = _dafny.Map({d_297_v1_: (self).f15})
            d_324_v27_: D0
            d_324_v27_ = D0_DC2(d_297_v1_, default__.fm10((d_322_v25_).cardinality, globalState), d_323_v26_, True)
            d_325_v28_: _dafny.Map
            d_325_v28_ = _dafny.Map({d_324_v27_: self.f12})
            (globalState).f5 = default__.safeDivisionInt((((d_325_v28_)[d_324_v27_] if (d_324_v27_) in (d_325_v28_) else (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) if True else d_301_v5_.f12), (d_301_v5_.f12) - ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]))
        elif True:
            if (d_297_v1_ if d_297_v1_ else False):
                d_326_v29_: _dafny.Map
                d_326_v29_ = _dafny.Map({_dafny.CodePoint('n'): d_297_v1_})
                d_327_v30_: _dafny.Seq
                d_327_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynbd"))
                d_297_v1_ = not (((d_326_v29_)[_dafny.CodePoint('w')] if (_dafny.CodePoint('w')) in (d_326_v29_) else d_297_v1_)) or (default__.fm1(len(d_327_v30_), default__.fm1(self.f12, not(d_297_v1_), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], d_297_v1_, globalState), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], d_297_v1_, globalState))
                d_328_v31_: _dafny.Map
                d_328_v31_ = _dafny.Map({-503: d_297_v1_})
                d_329_v32_: C0
                nw61_ = C0()
                nw61_.ctor__((d_328_v31_) | (d_328_v31_), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])
                d_329_v32_ = nw61_
                d_327_v30_ = d_327_v30_
                r0 = d_297_v1_
                r0 = default__.fm1(len((d_327_v30_) + ((_dafny.SeqWithoutIsStrInference([(self).f13 for d_330_i2_ in range(default__.abs(198))])).set(default__.safeIndex((_dafny.MultiSet([self.f12])).cardinality, len(_dafny.SeqWithoutIsStrInference([(self).f13 for d_331_i2_ in range(default__.abs(198))]))), _dafny.CodePoint('q')))), (self.f12) < ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], (d_297_v1_) and (d_297_v1_), globalState)
            elif True:
                d_332_v33_: _dafny.Seq
                d_332_v33_ = _dafny.SeqWithoutIsStrInference([d_297_v1_, d_297_v1_, d_297_v1_])
                d_333_v34_: _dafny.Seq
                d_333_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqtlrd"))
                d_334_v35_: _dafny.Seq
                d_334_v35_ = _dafny.SeqWithoutIsStrInference([d_333_v34_])
                d_335_v36_: D0
                d_335_v36_ = D0_DC0(True)
                d_336_v37_: D0
                d_336_v37_ = D0_DC3(d_335_v36_)
                d_337_v38_: _dafny.Array
                nw62_ = _dafny.Array(None, 23)
                nw62_[int(0)] = d_297_v1_
                nw62_[int(1)] = d_297_v1_
                nw62_[int(2)] = False
                nw62_[int(3)] = d_297_v1_
                nw62_[int(4)] = default__.fm1(self.f12, d_297_v1_, -989, d_297_v1_, globalState)
                nw62_[int(5)] = d_297_v1_
                nw62_[int(6)] = d_297_v1_
                nw62_[int(7)] = d_297_v1_
                nw62_[int(8)] = d_297_v1_
                nw62_[int(9)] = d_297_v1_
                nw62_[int(10)] = d_297_v1_
                nw62_[int(11)] = d_297_v1_
                nw62_[int(12)] = False
                nw62_[int(13)] = d_297_v1_
                nw62_[int(14)] = d_297_v1_
                nw62_[int(15)] = d_297_v1_
                nw62_[int(16)] = False
                nw62_[int(17)] = d_297_v1_
                nw62_[int(18)] = d_297_v1_
                nw62_[int(19)] = d_297_v1_
                nw62_[int(20)] = d_297_v1_
                nw62_[int(21)] = d_297_v1_
                nw62_[int(22)] = d_297_v1_
                d_337_v38_ = nw62_
                d_338_v39_: _dafny.Array
                nw63_ = _dafny.Array(None, 19)
                nw63_[int(0)] = d_297_v1_
                nw63_[int(1)] = d_297_v1_
                nw63_[int(2)] = d_297_v1_
                nw63_[int(3)] = (d_332_v33_)[default__.safeIndex(self.f12, len(d_332_v33_))]
                nw63_[int(4)] = d_297_v1_
                nw63_[int(5)] = default__.fm1(self.f12, True, 716, d_297_v1_, globalState)
                nw63_[int(6)] = not (d_297_v1_) or (d_297_v1_)
                nw63_[int(7)] = not (d_297_v1_) or (d_297_v1_)
                nw63_[int(8)] = d_297_v1_
                nw63_[int(9)] = d_297_v1_
                nw63_[int(10)] = (default__.fm11((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], self.f12, d_297_v1_, globalState)) <= (d_333_v34_)
                nw63_[int(11)] = (d_333_v34_) != ((d_334_v35_)[default__.safeIndex(len((d_332_v33_).set(default__.safeIndex((0) - ((0) - ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])), len(d_332_v33_)), False)), len(d_334_v35_))])
                nw63_[int(12)] = d_297_v1_
                nw63_[int(13)] = d_297_v1_
                nw63_[int(14)] = (d_336_v37_) in (_dafny.Map({d_336_v37_: d_337_v38_}))
                nw63_[int(15)] = d_297_v1_
                nw63_[int(16)] = d_297_v1_
                nw63_[int(17)] = not(d_297_v1_)
                nw63_[int(18)] = d_297_v1_
                d_338_v39_ = nw63_
                rhs63_ = d_297_v1_
                rhs64_ = -754
                rhs65_ = ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) * ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])
                rhs66_ = d_338_v39_
                rhs67_ = not(not (not(d_297_v1_)) or (True))
                lhs32_ = globalState
                lhs33_ = globalState
                r0 = rhs63_
                lhs32_.f5 = rhs64_
                lhs33_.f5 = rhs65_
                d_338_v39_ = rhs66_
                r0 = rhs67_
                d_339_v40_: _dafny.MultiSet
                d_339_v40_ = _dafny.MultiSet([self.f12, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]])
                d_340_v41_: _dafny.Map
                d_340_v41_ = _dafny.Map({d_297_v1_: d_297_v1_})
                d_341_v42_: _dafny.Map
                d_341_v42_ = _dafny.Map({((d_339_v40_)[len(d_340_v41_)] if (len(d_340_v41_)) in (d_339_v40_) else (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]): False})
                d_342_v43_: _dafny.Seq
                d_342_v43_ = default__.fm11(default__.fm0(self.f12, self.f12, globalState), (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], d_297_v1_, globalState)
                d_343_v44_: _dafny.Map
                d_343_v44_ = _dafny.Map({d_297_v1_: len((d_342_v43_))})
                d_344_v45_: C0
                nw64_ = C0()
                nw64_.ctor__(d_341_v42_, len(d_343_v44_))
                d_344_v45_ = nw64_
                d_345_v46_: _dafny.Array
                def lambda15_(d_346_v41_):
                    def lambda16_(d_347_i3_):
                        return d_346_v41_

                    return lambda16_

                init8_ = lambda15_(d_340_v41_)
                nw65_ = _dafny.Array(None, 7)
                for i0_8_ in range(nw65_.length(0)):
                    nw65_[i0_8_] = init8_(i0_8_)
                d_345_v46_ = nw65_
                index51_ = default__.safeIndex(873, (d_345_v46_).length(0))
                (d_345_v46_)[index51_] = (d_340_v41_) | (_dafny.Map({d_297_v1_: d_297_v1_}))
                index52_ = default__.safeIndex(873, (d_345_v46_).length(0))
                (d_345_v46_)[index52_] = d_340_v41_
                d_348_v47_: _dafny.Map
                d_348_v47_ = _dafny.Map({d_332_v33_: self.f12})
                d_348_v47_ = (d_348_v47_).set(_dafny.SeqWithoutIsStrInference([default__.fm1((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], d_297_v1_, self.f12, d_297_v1_, globalState)]), self.f12)
                d_349_v48_: _dafny.Set
                d_349_v48_ = _dafny.Set({True})
                d_350_v49_: _dafny.Map
                d_350_v49_ = _dafny.Map({d_349_v48_: ((d_339_v40_)[(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]] if ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) in (d_339_v40_) else len(d_333_v34_))})
                d_351_v50_: str
                d_351_v50_ = _dafny.CodePoint('s')
                d_352_v51_: _dafny.Array
                def lambda17_(d_353_v44_):
                    def lambda18_(d_354_i4_):
                        return d_353_v44_

                    return lambda18_

                init9_ = lambda17_(d_343_v44_)
                nw66_ = _dafny.Array(None, 2)
                for i0_9_ in range(nw66_.length(0)):
                    nw66_[i0_9_] = init9_(i0_9_)
                d_352_v51_ = nw66_
                index53_ = default__.safeIndex(647, (d_352_v51_).length(0))
                (d_352_v51_)[index53_] = d_343_v44_
                d_355_v53_: _dafny.Seq
                d_355_v53_ = _dafny.SeqWithoutIsStrInference([d_349_v48_])
                d_356_v54_: _dafny.Set
                d_356_v54_ = d_349_v48_
                d_357_v55_: D2
                d_357_v55_ = D2_DC8(d_297_v1_, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], d_297_v1_)
                index54_ = default__.safeIndex(647, (d_352_v51_).length(0))
                def iife16_():
                    coll6_ = _dafny.Map()
                    compr_6_: _dafny.Set
                    for compr_6_ in (d_355_v53_).Elements:
                        d_358_v52_: _dafny.Set = compr_6_
                        if (d_358_v52_) in (d_355_v53_):
                            coll6_[d_358_v52_] = (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]
                    return _dafny.Map(coll6_)
                rhs68_ = ((iife16_()
                ).set((d_356_v54_), (d_357_v55_).cf12)).set(_dafny.Set({d_297_v1_, d_297_v1_, d_297_v1_}), default__.fm0(self.f12, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], globalState))
                rhs69_ = d_351_v50_
                rhs70_ = d_343_v44_
                lhs34_ = d_352_v51_
                lhs35_ = default__.safeIndex(647, (d_352_v51_).length(0))
                d_350_v49_ = rhs68_
                d_351_v50_ = rhs69_
                lhs34_[lhs35_] = rhs70_
            d_359_v56_: _dafny.Seq
            d_359_v56_ = _dafny.SeqWithoutIsStrInference([d_297_v1_])
            d_360_v57_: D2
            d_360_v57_ = D2_DC7(706)
            d_361_v58_: _dafny.MultiSet
            d_361_v58_ = _dafny.MultiSet([len((_dafny.SeqWithoutIsStrInference([d_359_v56_, _dafny.SeqWithoutIsStrInference([False])])) + (_dafny.SeqWithoutIsStrInference([d_359_v56_ for d_362_i5_ in range(default__.abs(97))]))), self.f12, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], (d_360_v57_).cf10, (self.f12) - (self.f12)])
            d_361_v58_ = d_361_v58_
            d_297_v1_ = ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) <= (self.f12)
            d_363_v59_: _dafny.Map
            d_363_v59_ = _dafny.Map({(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]: d_297_v1_})
            d_364_v60_: _dafny.MultiSet
            d_364_v60_ = _dafny.MultiSet([d_297_v1_, False, d_297_v1_, d_297_v1_, d_297_v1_])
            d_365_v63_: _dafny.Array
            nw67_ = _dafny.Array(None, 13)
            nw67_[int(0)] = (d_363_v59_) | (_dafny.Map({(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]: d_297_v1_}))
            nw67_[int(1)] = _dafny.Map({(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]: d_297_v1_})
            nw67_[int(2)] = (d_363_v59_) | (d_363_v59_)
            nw67_[int(3)] = d_363_v59_
            nw67_[int(4)] = d_363_v59_
            nw67_[int(5)] = default__.fm10((d_364_v60_).cardinality, globalState)
            nw67_[int(6)] = d_363_v59_
            nw67_[int(7)] = (d_363_v59_) | (d_363_v59_)
            nw67_[int(8)] = d_363_v59_
            nw67_[int(9)] = d_363_v59_
            def iife17_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(14, 690):
                    d_366_v61_: int = compr_7_
                    if ((14) <= (d_366_v61_)) and ((d_366_v61_) < (690)):
                        coll7_[default__.safeModuloInt(d_366_v61_, self.f12)] = not(d_297_v1_)
                return _dafny.Map(coll7_)
            nw67_[int(10)] = iife17_()
            
            nw67_[int(11)] = d_363_v59_
            def iife18_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in _dafny.IntegerRange(382, 85):
                    d_367_v62_: int = compr_8_
                    if ((382) <= (d_367_v62_)) and ((d_367_v62_) < (85)):
                        coll8_[(d_367_v62_) + (self.f12)] = d_297_v1_
                return _dafny.Map(coll8_)
            nw67_[int(12)] = iife18_()
            
            d_365_v63_ = nw67_
            index55_ = default__.safeIndex(948, (d_365_v63_).length(0))
            (d_365_v63_)[index55_] = (d_363_v59_).set(self.f12, d_297_v1_)
            index56_ = default__.safeIndex(948, (d_365_v63_).length(0))
            (d_365_v63_)[index56_] = d_363_v59_
            d_368_v64_: _dafny.Array
            def lambda19_(d_369_i6_):
                return D2_DC9(524)

            init10_ = lambda19_
            nw68_ = _dafny.Array(None, 4)
            for i0_10_ in range(nw68_.length(0)):
                nw68_[i0_10_] = init10_(i0_10_)
            d_368_v64_ = nw68_
            d_370_v65_: _dafny.Map
            d_370_v65_ = _dafny.Map({d_361_v58_: d_368_v64_})
            d_370_v65_ = (d_370_v65_).set(_dafny.MultiSet([self.f12]), d_368_v64_)
        d_371_v66_: _dafny.Map
        d_371_v66_ = _dafny.Map({self.f12: not(not(True))})
        d_372_v67_: _dafny.MultiSet
        d_372_v67_ = _dafny.MultiSet([d_297_v1_, d_297_v1_])
        d_373_v68_: _dafny.Map
        d_373_v68_ = _dafny.Map({20: (d_372_v67_).cardinality})
        d_374_v71_: _dafny.Seq
        def iife19_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (d_373_v68_).keys.Elements:
                d_375_v69_: int = compr_9_
                if (d_375_v69_) in (d_373_v68_):
                    def iife20_():
                        coll10_ = _dafny.Map()
                        compr_10_: int
                        for compr_10_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({242: False})), 192])).Elements:
                            d_376_v70_: int = compr_10_
                            if (d_376_v70_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({242: False})), 192])):
                                coll10_[(d_376_v70_) + (851)] = not(True)
                        return _dafny.Map(coll10_)
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_375_v69_, len(_dafny.Map({(D0_DC2(not(True), iife20_()
, _dafny.Map({True: _dafny.CodePoint('f')}), True)).cf2: True})))]))
            return _dafny.Set(coll9_)
        d_374_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (len(_dafny.Set({d_371_v66_})))}), iife19_()
        ])
        d_377_v72_: _dafny.Seq
        d_377_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvadps"))
        d_378_v73_: _dafny.Set
        d_378_v73_ = _dafny.Set({self.f12})
        if ((d_374_v71_)[default__.safeIndex(len(d_377_v72_), len(d_374_v71_))]) == ((d_378_v73_) - (d_378_v73_)):
            d_379_v74_: _dafny.Seq
            d_379_v74_ = _dafny.SeqWithoutIsStrInference([(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]])
            r0 = default__.fm1((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], False, (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) > (len(d_379_v74_)), globalState)
            d_380_v75_: _dafny.Array
            nw69_ = _dafny.Array(False, 14)
            d_380_v75_ = nw69_
            d_381_v76_: D0
            d_381_v76_ = D0_DC0(d_297_v1_)
            index57_ = default__.safeIndex(89, (d_380_v75_).length(0))
            (d_380_v75_)[index57_] = (d_381_v76_).cf0
            d_382_v77_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
            d_382_v77_ = nw70_
            d_383_v78_: _dafny.Set
            d_383_v78_ = _dafny.Set({d_382_v77_, d_382_v77_})
            index58_ = default__.safeIndex(216, (d_295_v0_).length(0))
            index59_ = default__.safeIndex(89, (d_380_v75_).length(0))
            rhs71_ = len(d_383_v78_)
            rhs72_ = ((0) - ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])) <= (default__.safeModuloInt((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]))
            rhs73_ = self.f12
            rhs74_ = (d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]
            lhs36_ = d_295_v0_
            lhs37_ = default__.safeIndex(216, (d_295_v0_).length(0))
            lhs38_ = d_380_v75_
            lhs39_ = default__.safeIndex(89, (d_380_v75_).length(0))
            lhs36_[lhs37_] = rhs71_
            lhs38_[lhs39_] = rhs72_
            r3 = rhs73_
            r3 = rhs74_
            d_384_v79_: C0
            nw71_ = C0()
            nw71_.ctor__(_dafny.Map({550: True}), 688)
            d_384_v79_ = nw71_
            d_385_v80_: _dafny.Array
            nw72_ = _dafny.Array(_dafny.Seq({}), 29)
            d_385_v80_ = nw72_
            d_386_v81_: _dafny.Seq
            d_386_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_380_v75_)[default__.safeIndex(89, (d_380_v75_).length(0))]]), _dafny.SeqWithoutIsStrInference([(d_380_v75_)[default__.safeIndex(89, (d_380_v75_).length(0))], d_297_v1_])])
            index60_ = default__.safeIndex(563, (d_385_v80_).length(0))
            (d_385_v80_)[index60_] = d_386_v81_
            index61_ = default__.safeIndex(563, (d_385_v80_).length(0))
            (d_385_v80_)[index61_] = d_386_v81_
            d_380_v75_ = d_380_v75_
        elif True:
            d_387_v82_: C0
            nw73_ = C0()
            nw73_.ctor__(_dafny.Map({len(d_377_v72_): True}), self.f12)
            d_387_v82_ = nw73_
            d_388_v83_: C0
            nw74_ = C0()
            nw74_.ctor__(d_371_v66_, self.f12)
            d_388_v83_ = nw74_
            d_388_v83_ = d_388_v83_
            d_389_v84_: _dafny.Array
            def lambda20_(d_390_v72_):
                def lambda21_(d_391_i7_):
                    return (d_390_v72_) + (d_390_v72_)

                return lambda21_

            init11_ = lambda20_(d_377_v72_)
            nw75_ = _dafny.Array(None, 11)
            for i0_11_ in range(nw75_.length(0)):
                nw75_[i0_11_] = init11_(i0_11_)
            d_389_v84_ = nw75_
            index62_ = default__.safeIndex(857, (d_389_v84_).length(0))
            (d_389_v84_)[index62_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            d_392_v85_: _dafny.Seq
            d_392_v85_ = d_377_v72_
            index63_ = default__.safeIndex(857, (d_389_v84_).length(0))
            (d_389_v84_)[index63_] = ((d_377_v72_ if d_297_v1_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hu")))) + ((d_392_v85_))
            index64_ = default__.safeIndex(857, (d_389_v84_).length(0))
            (d_389_v84_)[index64_] = ((d_377_v72_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsdwgeit")))
        d_393_v86_: _dafny.Seq
        d_393_v86_ = _dafny.SeqWithoutIsStrInference([d_297_v1_])
        d_394_v87_: D2
        d_394_v87_ = D2_DC7(len(d_393_v86_))
        d_395_v88_: _dafny.Seq
        d_395_v88_ = _dafny.SeqWithoutIsStrInference([(d_394_v87_).cf10])
        d_396_v89_: D1
        d_396_v89_ = D1_DC4(d_295_v0_)
        d_397_v90_: _dafny.Map
        d_397_v90_ = _dafny.Map({(d_395_v88_)[default__.safeIndex(self.f12, len(d_395_v88_))]: (d_396_v89_).cf7})
        d_397_v90_ = (d_397_v90_).set(((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]) + ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]), d_295_v0_)
        hi1_ = self.f12
        for d_398_i8_ in range((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], hi1_):
            d_399_v91_: _dafny.Set
            d_399_v91_ = _dafny.Set({d_393_v86_})
            d_400_v92_: _dafny.Map
            d_400_v92_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_297_v1_, d_297_v1_]): _dafny.MultiSet([d_297_v1_, d_297_v1_, False, default__.fm1(len(d_378_v73_), d_297_v1_, self.f12, d_297_v1_, globalState)])})
            def iife21_():
                coll11_ = _dafny.Set()
                compr_11_: _dafny.Seq
                for compr_11_ in (d_400_v92_).keys.Elements:
                    d_401_v93_: _dafny.Seq = compr_11_
                    if (d_401_v93_) in (d_400_v92_):
                        coll11_ = coll11_.union(_dafny.Set([d_401_v93_]))
                return _dafny.Set(coll11_)
            if ((iife21_()
            ).ispropersubset(d_399_v91_) if True else not(not((d_297_v1_ if d_297_v1_ else d_297_v1_)))):
                d_377_v72_ = (d_377_v72_) + (_dafny.SeqWithoutIsStrInference([(self).f15 for d_402_i9_ in range(default__.abs(-898))]))
                d_377_v72_ = d_377_v72_
                r0 = d_297_v1_
                d_403_v94_: str
                d_403_v94_ = _dafny.CodePoint('p')
                d_403_v94_ = (self).f13
                d_404_v95_: _dafny.Array
                nw76_ = _dafny.Array(D0.default()(), 18)
                d_404_v95_ = nw76_
                d_404_v95_ = d_404_v95_
            elif True:
                d_405_v96_: _dafny.Array
                nw77_ = _dafny.Array(_dafny.Map({}), 7)
                d_405_v96_ = nw77_
                d_406_v97_: _dafny.Map
                d_406_v97_ = _dafny.Map({d_405_v96_: d_297_v1_})
                d_406_v97_ = (d_406_v97_).set(d_405_v96_, not (default__.fm1(self.f12, not(d_297_v1_), d_398_i8_, d_297_v1_, globalState)) or (d_297_v1_))
                d_407_v98_: T2
                nw78_ = C0()
                nw78_.ctor__(_dafny.Map({len(d_395_v88_): d_297_v1_}), self.f12)
                d_407_v98_ = nw78_
                d_408_v99_: _dafny.Seq
                d_408_v99_ = _dafny.SeqWithoutIsStrInference([d_407_v98_, d_407_v98_, d_407_v98_, d_407_v98_, d_407_v98_])
                d_409_v100_: _dafny.Map
                d_409_v100_ = _dafny.Map({d_408_v99_: d_407_v98_.f12})
                d_409_v100_ = (d_409_v100_).set(_dafny.SeqWithoutIsStrInference([d_407_v98_, d_407_v98_, d_407_v98_, d_407_v98_]), (self.f12) * ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]))
                index65_ = default__.safeIndex(216, (d_295_v0_).length(0))
                (d_295_v0_)[index65_] = len((d_377_v72_).set(default__.safeIndex((self).fm4(globalState), len(d_377_v72_)), (self).f15))
                d_373_v68_ = default__.fm12((0) - ((self.f12) - ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))])), d_407_v98_.f12, self.f12, (default__.fm0(len(d_393_v86_), len(d_373_v68_), globalState)) + ((d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))]), globalState)
                d_410_v101_: _dafny.Set
                d_410_v101_ = _dafny.Set({d_373_v68_, d_373_v68_, d_373_v68_})
                d_411_v102_: _dafny.Array
                nw79_ = _dafny.Array(None, 13)
                nw79_[int(0)] = d_297_v1_
                nw79_[int(1)] = d_297_v1_
                nw79_[int(2)] = d_297_v1_
                nw79_[int(3)] = d_297_v1_
                nw79_[int(4)] = d_297_v1_
                nw79_[int(5)] = d_297_v1_
                nw79_[int(6)] = True
                nw79_[int(7)] = not(d_297_v1_)
                nw79_[int(8)] = d_297_v1_
                nw79_[int(9)] = (d_297_v1_) in (d_393_v86_)
                nw79_[int(10)] = d_297_v1_
                nw79_[int(11)] = d_297_v1_
                nw79_[int(12)] = (d_410_v101_).issubset(d_410_v101_)
                d_411_v102_ = nw79_
                d_411_v102_ = (d_411_v102_ if ((d_371_v66_)[len(d_373_v68_)] if (len(d_373_v68_)) in (d_371_v66_) else d_297_v1_) else d_411_v102_)
            d_395_v88_ = d_395_v88_
            d_297_v1_ = not(True)
            d_412_v103_: D0
            d_412_v103_ = D0_DC0(d_297_v1_)
            d_297_v1_ = not((d_412_v103_).cf0)
        r0 = not(default__.fm1(self.f12, d_297_v1_, (self.f12) * (self.f12), False, globalState))
        r1 = d_295_v0_
        d_413_v104_: _dafny.Set
        d_413_v104_ = _dafny.Set({d_395_v88_})
        d_414_v105_: _dafny.MultiSet
        d_414_v105_ = _dafny.MultiSet([(d_295_v0_)[default__.safeIndex(216, (d_295_v0_).length(0))], self.f12])
        d_415_v106_: _dafny.Map
        d_415_v106_ = _dafny.Map({_dafny.CodePoint('b'): not(((d_371_v66_)[len(d_413_v104_)] if (len(d_413_v104_)) in (d_371_v66_) else default__.fm1(self.f12, d_297_v1_, (d_414_v105_).cardinality, d_297_v1_, globalState)))})
        r2 = (d_415_v106_) | (d_415_v106_)
        r3 = len(d_377_v72_)
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Map = _dafny.Map({})
        d_416_v0_: _dafny.Set
        d_416_v0_ = _dafny.Set({self.f12, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqyamhoax")))})
        r1 = (d_416_v0_).issubset(d_416_v0_)
        def lambda22_(source3_):
            if source3_.is_DC8:
                d_417___mcc_h0_ = source3_.cf11
                d_418___mcc_h1_ = source3_.cf12
                d_419___mcc_h2_ = source3_.cf13
                d_420_cf13_ = d_419___mcc_h2_
                d_421_cf12_ = d_418___mcc_h1_
                d_422_cf11_ = d_417___mcc_h0_
                return d_421_cf12_
            elif source3_.is_DC9:
                d_423___mcc_h3_ = source3_.cf14
                d_424_cf14_ = d_423___mcc_h3_
                return d_424_cf14_
            elif True:
                d_425___mcc_h4_ = source3_.cf10
                d_426_cf10_ = d_425___mcc_h4_
                return d_426_cf10_

        (self).f12 = lambda22_(D2_DC9(self.f12))
        d_427_v1_: bool
        d_427_v1_ = True
        d_428_i0_: int
        d_428_i0_ = 0
        with _dafny.label("1"):
            while d_427_v1_:
                with _dafny.c_label("1"):
                    if (d_428_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_428_i0_ = (d_428_i0_) + (1)
                    d_429_v2_: _dafny.Seq
                    d_429_v2_ = _dafny.SeqWithoutIsStrInference([d_427_v1_, d_427_v1_])
                    d_430_v3_: _dafny.Seq
                    d_430_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iy"))
                    d_431_v4_: _dafny.Map
                    d_431_v4_ = _dafny.Map({(self.f12 if d_427_v1_ else self.f12): (default__.fm0(len(d_429_v2_), len(d_430_v3_), globalState)) <= (self.f12)})
                    d_431_v4_ = (d_431_v4_).set(self.f12, not(d_427_v1_))
                    d_432_v5_: _dafny.Array
                    nw80_ = _dafny.Array(_dafny.Array(None, 0), 10)
                    d_432_v5_ = nw80_
                    d_433_v6_: D2
                    d_433_v6_ = D2_DC8(d_427_v1_, self.f12, d_427_v1_)
                    d_434_v7_: _dafny.Array
                    nw81_ = _dafny.Array(None, 5)
                    nw81_[int(0)] = (d_433_v6_).cf11
                    nw81_[int(1)] = d_427_v1_
                    nw81_[int(2)] = d_427_v1_
                    nw81_[int(3)] = d_427_v1_
                    nw81_[int(4)] = d_427_v1_
                    d_434_v7_ = nw81_
                    index66_ = default__.safeIndex(717, (d_432_v5_).length(0))
                    (d_432_v5_)[index66_] = d_434_v7_
                    index67_ = default__.safeIndex(717, (d_432_v5_).length(0))
                    rhs75_ = not ((self.f12) in (default__.fm13(d_427_v1_, d_427_v1_, globalState))) or (not(d_427_v1_))
                    rhs76_ = d_434_v7_
                    lhs40_ = d_432_v5_
                    lhs41_ = default__.safeIndex(717, (d_432_v5_).length(0))
                    d_427_v1_ = rhs75_
                    lhs40_[lhs41_] = rhs76_
                    d_435_v8_: _dafny.Set
                    d_435_v8_ = _dafny.Set({d_427_v1_})
                    d_436_v9_: D2
                    d_436_v9_ = D2_DC9(self.f12)
                    d_435_v8_ = default__.fm14(self.f12, d_436_v9_, globalState)
                    d_437_v10_: _dafny.Seq
                    d_437_v10_ = _dafny.SeqWithoutIsStrInference([self.f12])
                    d_438_v11_: _dafny.MultiSet
                    d_438_v11_ = _dafny.MultiSet([_dafny.CodePoint('j'), (self).f15, (self).f15])
                    (globalState).f5 = default__.safeModuloInt(self.f12, (d_437_v10_)[default__.safeIndex(((d_438_v11_)[(self).f15] if ((self).f15) in (d_438_v11_) else self.f12), len(d_437_v10_))])
                    pass
            pass
        r1 = ((d_427_v1_ if d_427_v1_ else d_427_v1_) if d_427_v1_ else not(d_427_v1_))
        d_439_v12_: _dafny.Seq
        d_439_v12_ = _dafny.SeqWithoutIsStrInference([self.f12, default__.safeDivisionInt(164, self.f12), self.f12])
        d_440_v13_: _dafny.Seq
        d_440_v13_ = _dafny.SeqWithoutIsStrInference([True, d_427_v1_])
        d_441_v14_: _dafny.Seq
        d_441_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iiunb"))
        d_439_v12_ = ((self).fm5((d_440_v13_) <= (d_440_v13_), d_427_v1_, d_427_v1_, True, globalState)).set(default__.safeIndex((len(d_441_v14_)) - (self.f12), len((self).fm5((d_440_v13_) <= (d_440_v13_), d_427_v1_, d_427_v1_, True, globalState))), (0) - ((self.f12) * (self.f12)))
        d_442_v15_: _dafny.Array
        nw82_ = _dafny.Array(_dafny.Array(None, 0), 26)
        d_442_v15_ = nw82_
        d_443_v16_: _dafny.Array
        def lambda23_(d_444_v0_):
            def lambda24_(d_445_i1_):
                return d_444_v0_

            return lambda24_

        init12_ = lambda23_(d_416_v0_)
        nw83_ = _dafny.Array(None, 25)
        for i0_12_ in range(nw83_.length(0)):
            nw83_[i0_12_] = init12_(i0_12_)
        d_443_v16_ = nw83_
        index68_ = default__.safeIndex(52, (d_442_v15_).length(0))
        (d_442_v15_)[index68_] = d_443_v16_
        d_446_v17_: _dafny.Map
        d_446_v17_ = _dafny.Map({d_427_v1_: d_443_v16_})
        index69_ = default__.safeIndex(52, (d_442_v15_).length(0))
        (d_442_v15_)[index69_] = ((d_446_v17_)[default__.fm1((0) - (len(d_441_v14_)), d_427_v1_, self.f12, d_427_v1_, globalState)] if (default__.fm1((0) - (len(d_441_v14_)), d_427_v1_, self.f12, d_427_v1_, globalState)) in (d_446_v17_) else d_443_v16_)
        def lambda25_(source4_):
            if source4_.is_DC1:
                d_447___mcc_h5_ = source4_.cf1
                d_448_cf1_ = d_447___mcc_h5_
                return default__.safeDivisionInt(592, self.f12)
            elif source4_.is_DC2:
                d_449___mcc_h6_ = source4_.cf2
                d_450___mcc_h7_ = source4_.cf3
                d_451___mcc_h8_ = source4_.cf4
                d_452___mcc_h9_ = source4_.cf5
                d_453_cf5_ = d_452___mcc_h9_
                d_454_cf4_ = d_451___mcc_h8_
                d_455_cf3_ = d_450___mcc_h7_
                d_456_cf2_ = d_449___mcc_h6_
                return (self.f12) + (self.f12)
            elif source4_.is_DC0:
                d_457___mcc_h10_ = source4_.cf0
                d_458_cf0_ = d_457___mcc_h10_
                return self.f12
            elif True:
                d_459___mcc_h11_ = source4_.cf6
                d_460_cf6_ = d_459___mcc_h11_
                return 103

        r0 = (0) - (lambda25_(D0_DC0(default__.fm1(self.f12, not(d_427_v1_), self.f12, d_427_v1_, globalState))))
        d_461_v18_: D1
        d_461_v18_ = D1_DC5(d_427_v1_)
        d_462_v19_: _dafny.Map
        d_462_v19_ = _dafny.Map({(d_461_v18_).cf8: 623})
        r1 = (d_427_v1_ if not(d_427_v1_) else default__.fm1(self.f12, not(d_427_v1_), (0) - (len(d_462_v19_)), d_427_v1_, globalState))
        d_463_v20_: _dafny.Set
        d_463_v20_ = _dafny.Set({d_427_v1_, False, d_427_v1_, d_427_v1_})
        d_464_v21_: _dafny.Map
        d_464_v21_ = _dafny.Map({self.f12: d_427_v1_})
        d_465_v22_: _dafny.Map
        d_465_v22_ = _dafny.Map({True: ((d_464_v21_)[421] if (421) in (d_464_v21_) else not(d_427_v1_))})
        d_466_v23_: _dafny.MultiSet
        d_466_v23_ = _dafny.MultiSet([self.f12, len(d_465_v22_)])
        d_467_v24_: _dafny.Map
        d_467_v24_ = _dafny.Map({_dafny.Set({d_427_v1_}): d_466_v23_})
        d_468_v26_: _dafny.Seq
        d_468_v26_ = _dafny.SeqWithoutIsStrInference([d_463_v20_])
        def iife22_():
            coll12_ = _dafny.Map()
            compr_12_: _dafny.Set
            for compr_12_ in (d_468_v26_).Elements:
                d_469_v25_: _dafny.Set = compr_12_
                if (d_469_v25_) in (d_468_v26_):
                    coll12_[d_469_v25_] = d_466_v23_
            return _dafny.Map(coll12_)
        r2 = (_dafny.Map({d_463_v20_: d_466_v23_})) | ((d_467_v24_) | (iife22_()
        ))
        return r0, r1, r2

    @property
    def f15(self):
        return self._f15

class C2(T0):
    def  __init__(self):
        self._f12: int = int(0)
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f17, f12):
        (self)._f17 = f17
        (self).f12 = f12

    @property
    def f17(self):
        return self._f17

class C3(T1, T2, T0):
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: str = _dafny.CodePoint('D')
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f14, f13, f12):
        (self)._f14 = f14
        (self)._f13 = f13
        (self).f12 = f12

    def fm4(self, globalState):
        return default__.safeDivisionInt(self.f12, default__.safeModuloInt(632, (_dafny.MultiSet([True])).cardinality))

    def fm5(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([self.f12])) + (_dafny.SeqWithoutIsStrInference([self.f12]))

    def fm6(self, globalState):
        return D0_DC0(False)

    def fm7(self, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdgs")) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xerenpimg")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sugwvjq"))) + (_dafny.SeqWithoutIsStrInference([(self).f13 for d_470_i0_ in range(default__.abs(567))]))))

    def m2(self, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_471_v0_: _dafny.Array
        def lambda26_(d_472_i0_):
            return (d_472_i0_) + (self.f12)

        init13_ = lambda26_
        nw84_ = _dafny.Array(None, 4)
        for i0_13_ in range(nw84_.length(0)):
            nw84_[i0_13_] = init13_(i0_13_)
        d_471_v0_ = nw84_
        index70_ = default__.safeIndex(341, (d_471_v0_).length(0))
        (d_471_v0_)[index70_] = default__.safeModuloInt(self.f12, self.f12)
        d_473_v1_: bool
        d_473_v1_ = True
        d_474_v2_: D2
        d_474_v2_ = D2_DC8(d_473_v1_, self.f12, d_473_v1_)
        index71_ = default__.safeIndex(341, (d_471_v0_).length(0))
        def lambda27_(source5_):
            if source5_.is_DC8:
                d_475___mcc_h0_ = source5_.cf11
                d_476___mcc_h1_ = source5_.cf12
                d_477___mcc_h2_ = source5_.cf13
                d_478_cf13_ = d_477___mcc_h2_
                d_479_cf12_ = d_476___mcc_h1_
                d_480_cf11_ = d_475___mcc_h0_
                return (d_479_cf12_) + (len(_dafny.SeqWithoutIsStrInference([True])))
            elif source5_.is_DC9:
                d_481___mcc_h3_ = source5_.cf14
                d_482_cf14_ = d_481___mcc_h3_
                return self.f12
            elif True:
                d_483___mcc_h4_ = source5_.cf10
                d_484_cf10_ = d_483___mcc_h4_
                return self.f12

        (d_471_v0_)[index71_] = lambda27_(d_474_v2_)
        d_485_i1_: int
        d_485_i1_ = 0
        with _dafny.label("2"):
            while d_473_v1_:
                with _dafny.c_label("2"):
                    if (d_485_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_485_i1_ = (d_485_i1_) + (1)
                    d_486_v3_: _dafny.Seq
                    d_486_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dii"))
                    (globalState).f5 = (default__.fm0(len(d_486_v3_), (d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))], globalState)) + (len((d_486_v3_).set(default__.safeIndex((d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))], len(d_486_v3_)), (self).f13)))
                    d_487_v4_: _dafny.Map
                    d_487_v4_ = _dafny.Map({858: True})
                    d_488_v5_: C0
                    nw85_ = C0()
                    nw85_.ctor__(d_487_v4_, default__.safeDivisionInt(default__.fm0((d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))], (d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))], globalState), (d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))]))
                    d_488_v5_ = nw85_
                    d_489_v6_: _dafny.Seq
                    d_489_v6_ = _dafny.SeqWithoutIsStrInference([d_473_v1_])
                    d_489_v6_ = (d_489_v6_ if d_473_v1_ else (_dafny.SeqWithoutIsStrInference([d_473_v1_, d_473_v1_, d_473_v1_, d_473_v1_, d_473_v1_])).set(default__.safeIndex((0) - (self.f12), len(_dafny.SeqWithoutIsStrInference([d_473_v1_, d_473_v1_, d_473_v1_, d_473_v1_, d_473_v1_]))), d_473_v1_))
                    d_490_v7_: _dafny.Array
                    nw86_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_490_v7_ = nw86_
                    d_491_v8_: _dafny.Array
                    nw87_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                    d_491_v8_ = nw87_
                    index72_ = default__.safeIndex(302, (d_490_v7_).length(0))
                    (d_490_v7_)[index72_] = d_491_v8_
                    d_492_v9_: _dafny.Seq
                    d_492_v9_ = _dafny.SeqWithoutIsStrInference([d_491_v8_])
                    d_493_v10_: _dafny.MultiSet
                    d_493_v10_ = _dafny.MultiSet([(self).f13, (self).f13, (self).f13, (self).f13])
                    index73_ = default__.safeIndex(302, (d_490_v7_).length(0))
                    (d_490_v7_)[index73_] = (d_492_v9_)[default__.safeIndex(((d_493_v10_)[(self).f13] if ((self).f13) in (d_493_v10_) else self.f12), len(d_492_v9_))]
                    pass
            pass
        d_494_v11_: _dafny.Map
        d_494_v11_ = _dafny.Map({(d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))]: d_471_v0_})
        d_494_v11_ = (d_494_v11_).set((0) - ((497) - ((d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))])), d_471_v0_)
        r3 = (d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))]
        d_473_v1_ = ((d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))]) < ((d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))])
        r0 = d_473_v1_
        d_495_v12_: _dafny.MultiSet
        d_495_v12_ = _dafny.MultiSet([d_473_v1_, d_473_v1_])
        d_496_v13_: D5
        d_496_v13_ = D5_DC12(_dafny.SeqWithoutIsStrInference([d_495_v12_, d_495_v12_]))
        r0 = (((d_496_v13_).cf17) + (_dafny.SeqWithoutIsStrInference([d_495_v12_ for d_497_i2_ in range(default__.abs(605))]))) <= (_dafny.SeqWithoutIsStrInference([d_495_v12_ for d_498_i3_ in range(default__.abs(924))]))
        r1 = d_471_v0_
        d_499_v14_: _dafny.Map
        d_499_v14_ = _dafny.Map({(self).f13: d_473_v1_})
        r2 = (d_499_v14_) | (d_499_v14_)
        d_500_v15_: _dafny.Set
        d_500_v15_ = _dafny.Set({d_473_v1_, default__.fm1(self.f12, d_473_v1_, (d_471_v0_)[default__.safeIndex(341, (d_471_v0_).length(0))], d_473_v1_, globalState)})
        d_501_v16_: _dafny.Set
        d_501_v16_ = _dafny.Set({_dafny.Set({d_473_v1_, d_473_v1_, d_473_v1_, False}), _dafny.Set({d_473_v1_, False, d_473_v1_, d_473_v1_}), d_500_v15_})
        r3 = len(d_501_v16_)
        return r0, r1, r2, r3

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        d_502_v0_: bool
        d_502_v0_ = True
        d_503_i0_: int
        d_503_i0_ = 0
        with _dafny.label("3"):
            while d_502_v0_:
                with _dafny.c_label("3"):
                    if (d_503_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_503_i0_ = (d_503_i0_) + (1)
                    d_504_v1_: _dafny.Map
                    d_504_v1_ = _dafny.Map({902: self.f12})
                    d_505_v2_: _dafny.Array
                    nw88_ = _dafny.Array(_dafny.Array(None, 0), 7)
                    d_505_v2_ = nw88_
                    d_506_v3_: _dafny.Seq
                    d_506_v3_ = _dafny.SeqWithoutIsStrInference([self.f12, 842])
                    d_507_v4_: _dafny.Seq
                    d_507_v4_ = _dafny.SeqWithoutIsStrInference([self.f12, len(d_506_v3_)])
                    d_508_v5_: _dafny.Array
                    nw89_ = _dafny.Array(None, 22)
                    nw89_[int(0)] = d_502_v0_
                    nw89_[int(1)] = True
                    nw89_[int(2)] = d_502_v0_
                    nw89_[int(3)] = d_502_v0_
                    nw89_[int(4)] = d_502_v0_
                    nw89_[int(5)] = default__.fm1(self.f12, d_502_v0_, self.f12, d_502_v0_, globalState)
                    nw89_[int(6)] = d_502_v0_
                    nw89_[int(7)] = d_502_v0_
                    nw89_[int(8)] = d_502_v0_
                    nw89_[int(9)] = d_502_v0_
                    nw89_[int(10)] = d_502_v0_
                    nw89_[int(11)] = d_502_v0_
                    nw89_[int(12)] = False
                    nw89_[int(13)] = d_502_v0_
                    nw89_[int(14)] = d_502_v0_
                    nw89_[int(15)] = default__.fm1(len(d_506_v3_), d_502_v0_, 129, d_502_v0_, globalState)
                    nw89_[int(16)] = default__.fm1(self.f12, d_502_v0_, self.f12, d_502_v0_, globalState)
                    nw89_[int(17)] = d_502_v0_
                    nw89_[int(18)] = d_502_v0_
                    nw89_[int(19)] = d_502_v0_
                    nw89_[int(20)] = d_502_v0_
                    nw89_[int(21)] = d_502_v0_
                    d_508_v5_ = nw89_
                    index74_ = default__.safeIndex(919, (d_505_v2_).length(0))
                    (d_505_v2_)[index74_] = d_508_v5_
                    index75_ = default__.safeIndex(919, (d_505_v2_).length(0))
                    rhs77_ = d_504_v1_
                    rhs78_ = (d_508_v5_ if d_502_v0_ else d_508_v5_)
                    lhs42_ = d_505_v2_
                    lhs43_ = default__.safeIndex(919, (d_505_v2_).length(0))
                    d_504_v1_ = rhs77_
                    lhs42_[lhs43_] = rhs78_
                    d_509_v6_: _dafny.Seq
                    d_509_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doktuoubt"))
                    if (default__.safeModuloInt(len(_dafny.Map({len(_dafny.Map({(0) - (len(d_509_v6_)): self.f12})): d_502_v0_})), 11)) == (len((d_509_v6_) + (_dafny.SeqWithoutIsStrInference([(d_509_v6_)[default__.safeIndex(self.f12, len(d_509_v6_))] for d_510_i1_ in range(default__.abs(94))])))):
                        d_511_v7_: _dafny.Map
                        d_511_v7_ = _dafny.Map({d_502_v0_: self.f12})
                        (globalState).f5 = (default__.safeDivisionInt(len(d_511_v7_), self.f12)) * (self.f12)
                        (self).f12 = self.f12
                        d_512_v8_: _dafny.Array
                        nw90_ = _dafny.Array(int(0), 27)
                        d_512_v8_ = nw90_
                        index76_ = default__.safeIndex(331, (d_512_v8_).length(0))
                        (d_512_v8_)[index76_] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_512_v8_])), self.f12)
                        index77_ = default__.safeIndex(331, (d_512_v8_).length(0))
                        (d_512_v8_)[index77_] = (0) - ((0) - (self.f12))
                        d_513_v9_: _dafny.Seq
                        d_513_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f13 for d_514_i2_ in range(default__.abs(-655))])])
                        d_515_v10_: D2
                        d_515_v10_ = D2_DC8(d_502_v0_, (d_512_v8_)[default__.safeIndex(331, (d_512_v8_).length(0))], d_502_v0_)
                        d_516_v11_: _dafny.Map
                        d_516_v11_ = _dafny.Map({len((d_513_v9_)[default__.safeIndex((d_512_v8_)[default__.safeIndex(331, (d_512_v8_).length(0))], len(d_513_v9_))]): d_515_v10_})
                        d_517_v12_: _dafny.Seq
                        d_517_v12_ = _dafny.SeqWithoutIsStrInference([d_502_v0_])
                        pat_let_tv1_ = d_502_v0_
                        index78_ = default__.safeIndex(331, (d_512_v8_).length(0))
                        def iife23_(_pat_let5_0):
                            def iife24_(d_518_dt__update__tmp_h0_):
                                def iife25_(_pat_let6_0):
                                    def iife26_(d_519_dt__update_hcf11_h0_):
                                        return D2_DC8(d_519_dt__update_hcf11_h0_, (d_518_dt__update__tmp_h0_).cf12, (d_518_dt__update__tmp_h0_).cf13)
                                    return iife26_(_pat_let6_0)
                                return iife25_(pat_let_tv1_)
                            return iife24_(_pat_let5_0)
                        (d_512_v8_)[index78_] = len((((d_516_v11_).set(self.f12, iife23_(d_515_v10_))).set(len(d_509_v6_), d_515_v10_)).set((d_512_v8_)[default__.safeIndex(331, (d_512_v8_).length(0))], D2_DC8((d_517_v12_)[default__.safeIndex(self.f12, len(d_517_v12_))], (d_512_v8_)[default__.safeIndex(331, (d_512_v8_).length(0))], not(not(d_502_v0_)))))
                        d_520_v13_: _dafny.Set
                        d_520_v13_ = _dafny.Set({(self).f13, (self).f13, (d_509_v6_)[default__.safeIndex(363, len(d_509_v6_))], _dafny.CodePoint('x')})
                        d_521_v14_: _dafny.Set
                        d_521_v14_ = _dafny.Set({d_502_v0_})
                        d_522_v15_: _dafny.Seq
                        d_522_v15_ = _dafny.SeqWithoutIsStrInference([d_521_v14_, d_521_v14_])
                        d_523_v16_: D6
                        d_523_v16_ = D6_DC15(d_507_v4_)
                        d_524_v17_: _dafny.Map
                        d_524_v17_ = _dafny.Map({d_520_v13_: (_dafny.SeqWithoutIsStrInference([len((d_522_v15_)[default__.safeIndex(self.f12, len(d_522_v15_))]), (0) - (default__.fm0((d_512_v8_)[default__.safeIndex(331, (d_512_v8_).length(0))], self.f12, globalState))]) if d_502_v0_ else (d_523_v16_).cf21)})
                        d_524_v17_ = _dafny.Map({(d_520_v13_) | (d_520_v13_): d_507_v4_})
                    elif True:
                        d_525_v18_: _dafny.Set
                        d_525_v18_ = _dafny.Set({d_506_v3_})
                        d_526_v19_: _dafny.Map
                        d_526_v19_ = _dafny.Map({d_525_v18_: d_509_v6_})
                        d_502_v0_ = not(not (d_502_v0_) or ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvvjigs"))) < (((d_526_v19_)[d_525_v18_] if (d_525_v18_) in (d_526_v19_) else d_509_v6_))))
                        d_527_v20_: _dafny.Array
                        nw91_ = _dafny.Array(int(0), 11)
                        d_527_v20_ = nw91_
                        index79_ = default__.safeIndex(737, (d_527_v20_).length(0))
                        (d_527_v20_)[index79_] = self.f12
                        d_528_v21_: D0
                        d_528_v21_ = D0_DC0(not(d_502_v0_))
                        d_529_v22_: D0
                        d_529_v22_ = D0_DC3(d_528_v21_)
                        index80_ = default__.safeIndex(663, (d_527_v20_).length(0))
                        (d_527_v20_)[index80_] = 894
                        index81_ = default__.safeIndex(928, (d_527_v20_).length(0))
                        (d_527_v20_)[index81_] = self.f12
                        index82_ = default__.safeIndex(66, (d_527_v20_).length(0))
                        (d_527_v20_)[index82_] = (0) - (((d_504_v1_)[self.f12] if (self.f12) in (d_504_v1_) else len(d_509_v6_)))
                        pat_let_tv2_ = d_528_v21_
                        index83_ = default__.safeIndex(737, (d_527_v20_).length(0))
                        index84_ = default__.safeIndex(663, (d_527_v20_).length(0))
                        index85_ = default__.safeIndex(928, (d_527_v20_).length(0))
                        index86_ = default__.safeIndex(66, (d_527_v20_).length(0))
                        def iife27_(_pat_let7_0):
                            def iife28_(d_530_dt__update__tmp_h1_):
                                def iife29_(_pat_let8_0):
                                    def iife30_(d_531_dt__update_hcf6_h0_):
                                        return D0_DC3(d_531_dt__update_hcf6_h0_)
                                    return iife30_(_pat_let8_0)
                                return iife29_(D0_DC3(pat_let_tv2_))
                            return iife28_(_pat_let7_0)
                        rhs79_ = self.f12
                        rhs80_ = iife27_(d_529_v22_)
                        rhs81_ = self.f12
                        rhs82_ = self.f12
                        rhs83_ = ((p0 if (d_502_v0_) or (not(d_502_v0_)) else p0)).cardinality
                        lhs44_ = d_527_v20_
                        lhs45_ = default__.safeIndex(737, (d_527_v20_).length(0))
                        lhs46_ = d_527_v20_
                        lhs47_ = default__.safeIndex(663, (d_527_v20_).length(0))
                        lhs48_ = d_527_v20_
                        lhs49_ = default__.safeIndex(928, (d_527_v20_).length(0))
                        lhs50_ = d_527_v20_
                        lhs51_ = default__.safeIndex(66, (d_527_v20_).length(0))
                        lhs44_[lhs45_] = rhs79_
                        d_529_v22_ = rhs80_
                        lhs46_[lhs47_] = rhs81_
                        lhs48_[lhs49_] = rhs82_
                        lhs50_[lhs51_] = rhs83_
                        r1 = d_502_v0_
                        d_532_v23_: _dafny.Set
                        d_532_v23_ = _dafny.Set({d_502_v0_, d_502_v0_})
                        (self).f12 = ((d_527_v20_)[default__.safeIndex(737, (d_527_v20_).length(0))] if d_502_v0_ else len((d_507_v4_) + ((self).fm5(d_502_v0_, d_502_v0_, d_502_v0_, default__.fm1(((d_504_v1_)[self.f12] if (self.f12) in (d_504_v1_) else len(d_532_v23_)), d_502_v0_, (d_527_v20_)[default__.safeIndex(737, (d_527_v20_).length(0))], True, globalState), globalState))))
                        r1 = not(d_502_v0_)
                    d_502_v0_ = False
                    d_533_v24_: D5
                    d_533_v24_ = D5_DC14(d_502_v0_)
                    d_534_v25_: _dafny.Set
                    d_534_v25_ = _dafny.Set({d_502_v0_})
                    d_535_v26_: _dafny.MultiSet
                    d_535_v26_ = _dafny.MultiSet([self.f12, len(d_534_v25_)])
                    pat_let_tv3_ = d_535_v26_
                    pat_let_tv4_ = d_535_v26_
                    def iife31_(_pat_let9_0):
                        def iife32_(d_536_dt__update__tmp_h2_):
                            def iife33_(_pat_let10_0):
                                def iife34_(d_537_dt__update_hcf20_h0_):
                                    return D5_DC14(d_537_dt__update_hcf20_h0_)
                                return iife34_(_pat_let10_0)
                            return iife33_((pat_let_tv3_).issubset(pat_let_tv4_))
                        return iife32_(_pat_let9_0)
                    source6_ = iife31_(d_533_v24_)
                    if source6_.is_DC13:
                        d_538___mcc_h0_ = source6_.cf18
                        d_539___mcc_h1_ = source6_.cf19
                        d_540_cf19_ = d_539___mcc_h1_
                        d_541_cf18_ = d_538___mcc_h0_
                        d_542_v27_: _dafny.Array
                        nw92_ = _dafny.Array(None, 1)
                        nw92_[int(0)] = (self.f12) * ((d_507_v4_)[default__.safeIndex(self.f12, len(d_507_v4_))])
                        d_542_v27_ = nw92_
                        index87_ = default__.safeIndex(237, (d_542_v27_).length(0))
                        (d_542_v27_)[index87_] = self.f12
                        index88_ = default__.safeIndex(237, (d_542_v27_).length(0))
                        (d_542_v27_)[index88_] = self.f12
                        d_509_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbfyp"))
                        (globalState).f5 = (0) - (self.f12)
                        rhs84_ = d_541_cf18_
                        rhs85_ = (0) - (self.f12)
                        r1 = rhs84_
                        r2 = rhs85_
                    elif source6_.is_DC14:
                        d_543___mcc_h2_ = source6_.cf20
                        d_544_cf20_ = d_543___mcc_h2_
                        d_545_v28_: _dafny.Seq
                        d_545_v28_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('n')) in (d_509_v6_)])
                        d_502_v0_ = (d_545_v28_)[default__.safeIndex(self.f12, len(d_545_v28_))]
                        d_546_v29_: _dafny.MultiSet
                        d_546_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbjwfs")), d_509_v6_])
                        r1 = default__.fm1(self.f12, not(default__.fm1(self.f12, d_544_cf20_, self.f12, d_502_v0_, globalState)), ((d_546_v29_).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f13 for d_547_i3_ in range(default__.abs(653))]), _dafny.SeqWithoutIsStrInference([(self).f13 for d_548_i4_ in range(default__.abs(-737))])]))).cardinality, d_544_cf20_, globalState)
                        d_549_v30_: _dafny.Map
                        d_549_v30_ = _dafny.Map({False: d_502_v0_})
                        d_550_v31_: _dafny.Seq
                        d_550_v31_ = _dafny.SeqWithoutIsStrInference([d_534_v25_, d_534_v25_, d_534_v25_, d_534_v25_, _dafny.Set({True, d_544_cf20_, d_502_v0_})])
                        d_551_v32_: _dafny.Map
                        d_551_v32_ = _dafny.Map({d_550_v31_: (self).f13})
                        d_552_v33_: _dafny.Array
                        nw93_ = _dafny.Array(None, 10)
                        nw93_[int(0)] = (self).f13
                        nw93_[int(1)] = default__.fm15(d_549_v30_, len(d_509_v6_), globalState)
                        nw93_[int(2)] = ((d_551_v32_)[d_550_v31_] if (d_550_v31_) in (d_551_v32_) else (self).f13)
                        nw93_[int(3)] = (self).f13
                        nw93_[int(4)] = (self).f13
                        nw93_[int(5)] = (self).f13
                        nw93_[int(6)] = (self).f13
                        nw93_[int(7)] = (self).f13
                        nw93_[int(8)] = (self).f13
                        nw93_[int(9)] = (self).f13
                        d_552_v33_ = nw93_
                        rhs86_ = d_552_v33_
                        rhs87_ = ((0) - (self.f12)) - (self.f12)
                        lhs52_ = globalState
                        d_552_v33_ = rhs86_
                        lhs52_.f5 = rhs87_
                        d_502_v0_ = d_544_cf20_
                    elif True:
                        d_553___mcc_h3_ = source6_.cf17
                        d_554_cf17_ = d_553___mcc_h3_
                        d_509_v6_ = (d_509_v6_).set(default__.safeIndex(self.f12, len(d_509_v6_)), (self).f13)
                        d_555_v34_: _dafny.Array
                        def lambda28_(d_556_i5_):
                            return (d_556_i5_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([(self).f13 for d_557_i6_ in range(default__.abs(491))]))))

                        init14_ = lambda28_
                        nw94_ = _dafny.Array(None, 1)
                        for i0_14_ in range(nw94_.length(0)):
                            nw94_[i0_14_] = init14_(i0_14_)
                        d_555_v34_ = nw94_
                        d_555_v34_ = (d_555_v34_ if (d_502_v0_) == (d_502_v0_) else (d_555_v34_ if d_502_v0_ else d_555_v34_))
                        r1 = not(not(d_502_v0_))
                        d_558_v35_: C1
                        nw95_ = C1()
                        nw95_.ctor__((self).f13, (self).f13, self.f12)
                        d_558_v35_ = nw95_
                    pass
            pass
        d_559_v36_: _dafny.Map
        d_559_v36_ = _dafny.Map({d_502_v0_: d_502_v0_})
        d_560_v37_: _dafny.Seq
        d_560_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpstgyprm"))
        d_561_v38_: _dafny.Seq
        d_561_v38_ = _dafny.SeqWithoutIsStrInference([default__.fm1(len(d_559_v36_), d_502_v0_, len(d_560_v37_), d_502_v0_, globalState)])
        d_562_v39_: _dafny.Set
        d_562_v39_ = _dafny.Set({default__.fm1(self.f12, False, len(d_561_v38_), d_502_v0_, globalState), d_502_v0_, ((d_559_v36_)[d_502_v0_] if (d_502_v0_) in (d_559_v36_) else d_502_v0_), d_502_v0_})
        source7_ = d_562_v39_
        d_563___mcc_h4_ = source7_
        d_564_cf16_ = d_563___mcc_h4_
        d_565_v40_: _dafny.Array
        nw96_ = _dafny.Array(None, 8)
        nw96_[int(0)] = not(default__.fm1(783, d_502_v0_, self.f12, d_502_v0_, globalState))
        nw96_[int(1)] = (d_560_v37_) <= (d_560_v37_)
        nw96_[int(2)] = True
        nw96_[int(3)] = d_502_v0_
        nw96_[int(4)] = d_502_v0_
        nw96_[int(5)] = (self.f12) > (self.f12)
        nw96_[int(6)] = not((not(d_502_v0_)) or (d_502_v0_))
        nw96_[int(7)] = d_502_v0_
        d_565_v40_ = nw96_
        index89_ = default__.safeIndex(553, (d_565_v40_).length(0))
        (d_565_v40_)[index89_] = d_502_v0_
        d_566_v41_: _dafny.Seq
        d_566_v41_ = _dafny.SeqWithoutIsStrInference([len(d_560_v37_)])
        d_567_v42_: _dafny.Map
        d_567_v42_ = _dafny.Map({d_502_v0_: (self).f13})
        d_568_v44_: _dafny.Set
        d_568_v44_ = _dafny.Set({(self).f13})
        index90_ = default__.safeIndex(553, (d_565_v40_).length(0))
        def iife35_():
            coll13_ = _dafny.Map()
            compr_13_: str
            for compr_13_ in (d_568_v44_).Elements:
                d_569_v43_: str = compr_13_
                if (d_569_v43_) in (d_568_v44_):
                    coll13_[d_569_v43_] = d_502_v0_
            return _dafny.Map(coll13_)
        (d_565_v40_)[index90_] = ((len(d_566_v41_)) == (self.f12) if d_502_v0_ else (((d_567_v42_)[d_502_v0_] if (d_502_v0_) in (d_567_v42_) else (self).f13)) not in (iife35_()
        ))
        d_570_v46_: _dafny.Map
        d_570_v46_ = _dafny.Map({self.f12: self.f12})
        d_571_v47_: _dafny.Map
        d_571_v47_ = _dafny.Map({self.f12: len(d_570_v46_)})
        d_572_v48_: _dafny.Map
        d_572_v48_ = _dafny.Map({len(d_571_v47_): d_502_v0_})
        d_573_v49_: C0
        nw97_ = C0()
        def iife36_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in (d_566_v41_).Elements:
                d_574_v45_: int = compr_14_
                if (d_574_v45_) in (d_566_v41_):
                    coll14_[(d_574_v45_) - ((0) - (-485))] = not((d_565_v40_)[default__.safeIndex(553, (d_565_v40_).length(0))])
            return _dafny.Map(coll14_)
        nw97_.ctor__((iife36_()
        ) | (d_572_v48_), self.f12)
        d_573_v49_ = nw97_
        d_575_v50_: _dafny.Array
        def lambda29_(d_576_i7_):
            return (d_576_i7_) - (self.f12)

        init15_ = lambda29_
        nw98_ = _dafny.Array(None, 22)
        for i0_15_ in range(nw98_.length(0)):
            nw98_[i0_15_] = init15_(i0_15_)
        d_575_v50_ = nw98_
        index91_ = default__.safeIndex(919, (d_575_v50_).length(0))
        (d_575_v50_)[index91_] = self.f12
        index92_ = default__.safeIndex(919, (d_575_v50_).length(0))
        (d_575_v50_)[index92_] = (self.f12) + (len(d_564_cf16_))
        (d_573_v49_).f16 = (d_573_v49_.f16).set(437, (self.f12) not in ((self).fm5(not(d_502_v0_), d_502_v0_, d_502_v0_, d_502_v0_, globalState)))
        d_577_v51_: _dafny.Array
        nw99_ = _dafny.Array(False, 27)
        d_577_v51_ = nw99_
        index93_ = default__.safeIndex(348, (d_577_v51_).length(0))
        (d_577_v51_)[index93_] = d_502_v0_
        d_578_v52_: _dafny.Seq
        d_578_v52_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_579_v53_: D5
        d_579_v53_ = D5_DC12(d_578_v52_)
        d_580_v54_: _dafny.Seq
        d_580_v54_ = _dafny.SeqWithoutIsStrInference([d_579_v53_])
        index94_ = default__.safeIndex(348, (d_577_v51_).length(0))
        (d_577_v51_)[index94_] = ((_dafny.SeqWithoutIsStrInference([d_579_v53_])) + ((d_580_v54_).set(default__.safeIndex(self.f12, len(d_580_v54_)), d_579_v53_))) <= (d_580_v54_)
        d_581_v55_: _dafny.MultiSet
        d_581_v55_ = _dafny.MultiSet([d_561_v38_, _dafny.SeqWithoutIsStrInference([(d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]]), d_561_v38_])
        d_582_v56_: _dafny.Set
        d_582_v56_ = _dafny.Set({self.f12, self.f12})
        d_583_v57_: _dafny.Seq
        d_583_v57_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f12, len(d_582_v56_), self.f12])
        d_584_v58_: D6
        d_584_v58_ = D6_DC15(d_583_v57_)
        d_502_v0_ = ((d_581_v55_).intersection(d_581_v55_)).issubset((d_581_v55_) | (default__.fm16(d_502_v0_, self.f12, d_584_v58_, self.f12, globalState)))
        if ((d_560_v37_ if True else d_560_v37_)) != (d_560_v37_):
            d_585_v59_: D2
            d_585_v59_ = D2_DC8((d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))], self.f12, d_502_v0_)
            pat_let_tv5_ = d_502_v0_
            d_586_v60_: _dafny.Array
            nw100_ = _dafny.Array(None, 29)
            nw100_[int(0)] = d_585_v59_
            nw100_[int(1)] = d_585_v59_
            nw100_[int(2)] = d_585_v59_
            nw100_[int(3)] = d_585_v59_
            nw100_[int(4)] = d_585_v59_
            nw100_[int(5)] = default__.fm17(d_502_v0_, default__.fm1(self.f12, d_502_v0_, self.f12, d_502_v0_, globalState), globalState)
            nw100_[int(6)] = d_585_v59_
            nw100_[int(7)] = d_585_v59_
            nw100_[int(8)] = d_585_v59_
            def iife37_(_pat_let11_0):
                def iife38_(d_587_dt__update__tmp_h3_):
                    def iife39_(_pat_let12_0):
                        def iife40_(d_588_dt__update_hcf13_h0_):
                            return D2_DC8((d_587_dt__update__tmp_h3_).cf11, (d_587_dt__update__tmp_h3_).cf12, d_588_dt__update_hcf13_h0_)
                        return iife40_(_pat_let12_0)
                    return iife39_(pat_let_tv5_)
                return iife38_(_pat_let11_0)
            nw100_[int(9)] = iife37_(d_585_v59_)
            nw100_[int(10)] = d_585_v59_
            nw100_[int(11)] = d_585_v59_
            nw100_[int(12)] = d_585_v59_
            nw100_[int(13)] = D2_DC8(False, len(d_561_v38_), d_502_v0_)
            nw100_[int(14)] = d_585_v59_
            nw100_[int(15)] = d_585_v59_
            nw100_[int(16)] = d_585_v59_
            nw100_[int(17)] = d_585_v59_
            nw100_[int(18)] = d_585_v59_
            nw100_[int(19)] = d_585_v59_
            nw100_[int(20)] = d_585_v59_
            nw100_[int(21)] = d_585_v59_
            nw100_[int(22)] = d_585_v59_
            nw100_[int(23)] = d_585_v59_
            nw100_[int(24)] = d_585_v59_
            nw100_[int(25)] = default__.fm17(d_502_v0_, d_502_v0_, globalState)
            nw100_[int(26)] = d_585_v59_
            nw100_[int(27)] = d_585_v59_
            nw100_[int(28)] = d_585_v59_
            d_586_v60_ = nw100_
            d_589_v61_: _dafny.Seq
            d_589_v61_ = _dafny.SeqWithoutIsStrInference([d_586_v60_, d_586_v60_])
            d_589_v61_ = (d_589_v61_ if (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))] else d_589_v61_)
            d_590_v62_: int
            d_591_v63_: _dafny.Map
            d_592_v64_: _dafny.Seq
            out39_: int
            out40_: _dafny.Map
            out41_: _dafny.Seq
            out39_, out40_, out41_ = default__.m0(self.f12, (113) - ((0) - (self.f12)), self.f12, (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))], globalState)
            d_590_v62_ = out39_
            d_591_v63_ = out40_
            d_592_v64_ = out41_
            d_502_v0_ = (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]
            d_593_v65_: _dafny.Map
            d_593_v65_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([not(d_502_v0_)]): d_502_v0_})
            d_593_v65_ = (d_593_v65_).set(d_561_v38_, (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))])
            index95_ = default__.safeIndex(348, (d_577_v51_).length(0))
            (d_577_v51_)[index95_] = d_502_v0_
        elif True:
            index96_ = default__.safeIndex(348, (d_577_v51_).length(0))
            (d_577_v51_)[index96_] = (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]
            d_594_v66_: _dafny.Array
            def lambda30_(d_595_i8_):
                return _dafny.Map({self.f12: (0) - (self.f12)})

            init16_ = lambda30_
            nw101_ = _dafny.Array(None, 27)
            for i0_16_ in range(nw101_.length(0)):
                nw101_[i0_16_] = init16_(i0_16_)
            d_594_v66_ = nw101_
            d_596_v67_: _dafny.Map
            d_596_v67_ = _dafny.Map({self.f12: False})
            d_597_v68_: _dafny.Seq
            d_597_v68_ = _dafny.SeqWithoutIsStrInference([d_561_v38_, _dafny.SeqWithoutIsStrInference([(d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))], False, False, d_502_v0_])])
            d_598_v69_: _dafny.Map
            d_598_v69_ = _dafny.Map({len(d_596_v67_): (D2_DC9(default__.fm0(self.f12, (_dafny.MultiSet((d_597_v68_)[default__.safeIndex(408, len(d_597_v68_))])).cardinality, globalState))).cf14})
            index97_ = default__.safeIndex(218, (d_594_v66_).length(0))
            (d_594_v66_)[index97_] = (d_598_v69_) | (_dafny.Map({self.f12: self.f12}))
            d_599_v71_: _dafny.MultiSet
            d_599_v71_ = _dafny.MultiSet([(self).f13])
            index98_ = default__.safeIndex(218, (d_594_v66_).length(0))
            def iife41_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in (d_583_v57_).Elements:
                    d_600_v70_: int = compr_15_
                    if (d_600_v70_) in (d_583_v57_):
                        coll15_[(d_600_v70_) * (901)] = self.f12
                return _dafny.Map(coll15_)
            rhs88_ = iife41_()

            rhs89_ = d_560_v37_
            rhs90_ = (_dafny.MultiSet(default__.fm11(self.f12, (_dafny.MultiSet(d_583_v57_)).cardinality, (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))], globalState))).issubset((_dafny.MultiSet([(self).f13, _dafny.CodePoint('f'), (self).f13])) - (d_599_v71_))
            rhs91_ = d_560_v37_
            lhs53_ = d_594_v66_
            lhs54_ = default__.safeIndex(218, (d_594_v66_).length(0))
            lhs53_[lhs54_] = rhs88_
            d_560_v37_ = rhs89_
            d_502_v0_ = rhs90_
            d_560_v37_ = rhs91_
            d_601_v72_: _dafny.Array
            nw102_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_601_v72_ = nw102_
            d_602_v73_: _dafny.Map
            d_602_v73_ = _dafny.Map({default__.fm1(self.f12, False, self.f12, False, globalState): d_601_v72_})
            d_602_v73_ = (d_602_v73_).set(d_502_v0_, ((d_602_v73_)[d_502_v0_] if (d_502_v0_) in (d_602_v73_) else d_601_v72_))
            d_603_v74_: str
            d_603_v74_ = _dafny.CodePoint('o')
            d_603_v74_ = default__.fm15(d_559_v36_, (d_583_v57_)[default__.safeIndex(len(d_560_v37_), len(d_583_v57_))], globalState)
            d_582_v56_ = d_582_v56_
        d_604_i9_: int
        d_604_i9_ = 0
        with _dafny.label("4"):
            while (d_561_v38_)[default__.safeIndex(((0) - (self.f12)) + (self.f12), len(d_561_v38_))]:
                with _dafny.c_label("4"):
                    if (d_604_i9_) >= (100):
                        raise _dafny.Break("4")
                    d_604_i9_ = (d_604_i9_) + (1)
                    (globalState).f5 = 479
                    d_605_v75_: _dafny.Map
                    d_605_v75_ = _dafny.Map({self.f12: (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]})
                    d_605_v75_ = (d_605_v75_).set(self.f12, not ((d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]) or ((d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]))
                    d_606_v76_: str
                    d_606_v76_ = _dafny.CodePoint('u')
                    d_606_v76_ = d_606_v76_
                    d_502_v0_ = (((d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]) not in (_dafny.Set({d_502_v0_, d_502_v0_})) if (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))] else (self.f12) == ((0) - (self.f12)))
                    pass
            pass
        r0 = (832) * (self.f12)
        r1 = (d_577_v51_)[default__.safeIndex(348, (d_577_v51_).length(0))]
        d_607_v77_: _dafny.MultiSet
        d_607_v77_ = _dafny.MultiSet([(d_560_v37_).set(default__.safeIndex(self.f12, len(d_560_v37_)), (self).f13), d_560_v37_])
        r2 = default__.safeModuloInt(((d_607_v77_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptyrk"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptyrk"))) in (d_607_v77_) else self.f12), (self.f12) * (self.f12))
        return r0, r1, r2

    def m4(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        hi2_ = (self.f12) - (self.f12)
        for d_608_i0_ in range(p0, hi2_):
            d_609_v0_: _dafny.Map
            d_609_v0_ = _dafny.Map({False: self.f12})
            d_610_v1_: _dafny.MultiSet
            d_610_v1_ = _dafny.MultiSet([d_609_v0_, d_609_v0_])
            d_610_v1_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_609_v0_]))
            d_611_v2_: _dafny.Array
            nw103_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_611_v2_ = nw103_
            d_612_v3_: D2
            d_612_v3_ = D2_DC8(p1, -238, p1)
            d_613_v4_: _dafny.Set
            d_613_v4_ = _dafny.Set({(d_612_v3_).cf11})
            d_614_v5_: _dafny.Map
            d_614_v5_ = _dafny.Map({p0: p1})
            d_615_v6_: T0
            nw104_ = C2()
            nw104_.ctor__(True, (0) - (d_608_i0_))
            d_615_v6_ = nw104_
            d_616_v7_: _dafny.Seq
            d_616_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, len(_dafny.Map({p0: d_615_v6_}))])
            d_617_v8_: _dafny.Map
            d_617_v8_ = _dafny.Map({(D2_DC8(p1, p0, p1)).cf12: self.f12})
            d_618_v9_: _dafny.Array
            nw105_ = _dafny.Array(None, 25)
            nw105_[int(0)] = self.f12
            nw105_[int(1)] = p0
            nw105_[int(2)] = len(default__.fm18(len(p2), 421, p1, globalState))
            nw105_[int(3)] = self.f12
            nw105_[int(4)] = self.f12
            nw105_[int(5)] = p0
            nw105_[int(6)] = len(d_613_v4_)
            nw105_[int(7)] = len(p2)
            nw105_[int(8)] = self.f12
            nw105_[int(9)] = self.f12
            nw105_[int(10)] = d_608_i0_
            nw105_[int(11)] = 203
            nw105_[int(12)] = p0
            nw105_[int(13)] = d_608_i0_
            nw105_[int(14)] = p0
            nw105_[int(15)] = d_608_i0_
            nw105_[int(16)] = (0) - (len(p2))
            nw105_[int(17)] = 968
            nw105_[int(18)] = self.f12
            nw105_[int(19)] = d_608_i0_
            nw105_[int(20)] = len((d_614_v5_).set(len(d_616_v7_), p1))
            nw105_[int(21)] = p0
            nw105_[int(22)] = p0
            nw105_[int(23)] = len(d_617_v8_)
            nw105_[int(24)] = d_608_i0_
            d_618_v9_ = nw105_
            index99_ = default__.safeIndex(47, (d_611_v2_).length(0))
            (d_611_v2_)[index99_] = d_618_v9_
            d_619_v10_: D6
            d_619_v10_ = D6_DC15((d_616_v7_ if p1 else d_616_v7_))
            index100_ = default__.safeIndex(47, (d_611_v2_).length(0))
            rhs92_ = d_615_v6_.f12
            rhs93_ = d_618_v9_
            rhs94_ = p1
            rhs95_ = p1
            rhs96_ = D6_DC15(_dafny.SeqWithoutIsStrInference([len(p2) for d_620_i1_ in range(default__.abs(684))]))
            lhs55_ = d_611_v2_
            lhs56_ = default__.safeIndex(47, (d_611_v2_).length(0))
            r2 = rhs92_
            lhs55_[lhs56_] = rhs93_
            r1 = rhs94_
            r1 = rhs95_
            d_619_v10_ = rhs96_
            index101_ = default__.safeIndex(47, (d_611_v2_).length(0))
            nw106_ = _dafny.Array(int(0), 8)
            (d_611_v2_)[index101_] = nw106_
            d_621_v11_: _dafny.Seq
            d_621_v11_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_622_v12_: _dafny.Array
            nw107_ = _dafny.Array(None, 16)
            nw107_[int(0)] = p1
            nw107_[int(1)] = default__.fm1(p0, p1, -195, p1, globalState)
            nw107_[int(2)] = p1
            nw107_[int(3)] = p1
            nw107_[int(4)] = p1
            nw107_[int(5)] = p1
            nw107_[int(6)] = p1
            nw107_[int(7)] = p1
            nw107_[int(8)] = False
            nw107_[int(9)] = p1
            nw107_[int(10)] = (d_621_v11_)[default__.safeIndex(p0, len(d_621_v11_))]
            nw107_[int(11)] = p1
            nw107_[int(12)] = p1
            nw107_[int(13)] = p1
            nw107_[int(14)] = p1
            nw107_[int(15)] = p1
            d_622_v12_ = nw107_
            d_623_v13_: _dafny.Map
            d_623_v13_ = _dafny.Map({d_608_i0_: d_622_v12_})
            d_624_v14_: _dafny.Seq
            d_624_v14_ = _dafny.SeqWithoutIsStrInference([d_622_v12_])
            d_623_v13_ = (d_623_v13_).set(self.f12, (d_624_v14_)[default__.safeIndex((d_616_v7_)[default__.safeIndex(self.f12, len(d_616_v7_))], len(d_624_v14_))])
        if p1:
            d_625_v15_: _dafny.Set
            d_625_v15_ = _dafny.Set({self.f12})
            d_626_v16_: _dafny.Map
            d_626_v16_ = _dafny.Map({p1: self.f12})
            d_627_v17_: _dafny.Map
            d_627_v17_ = _dafny.Map({self.f12: p1})
            d_628_v18_: T2
            nw108_ = C0()
            nw108_.ctor__(d_627_v17_, self.f12)
            d_628_v18_ = nw108_
            d_629_v19_: _dafny.Seq
            d_629_v19_ = _dafny.SeqWithoutIsStrInference([d_628_v18_])
            d_630_v20_: _dafny.MultiSet
            d_630_v20_ = _dafny.MultiSet([False])
            d_631_v21_: _dafny.Array
            nw109_ = _dafny.Array(None, 29)
            nw109_[int(0)] = p0
            nw109_[int(1)] = self.f12
            nw109_[int(2)] = (0) - (self.f12)
            nw109_[int(3)] = self.f12
            nw109_[int(4)] = len(d_625_v15_)
            nw109_[int(5)] = p0
            nw109_[int(6)] = p0
            nw109_[int(7)] = p0
            nw109_[int(8)] = self.f12
            nw109_[int(9)] = self.f12
            nw109_[int(10)] = p0
            nw109_[int(11)] = self.f12
            nw109_[int(12)] = ((d_626_v16_)[p1] if (p1) in (d_626_v16_) else p0)
            nw109_[int(13)] = p0
            nw109_[int(14)] = len((d_629_v19_).set(default__.safeIndex((d_630_v20_).cardinality, len(d_629_v19_)), d_628_v18_))
            nw109_[int(15)] = p0
            nw109_[int(16)] = p0
            nw109_[int(17)] = (0) - (p0)
            nw109_[int(18)] = d_628_v18_.f12
            nw109_[int(19)] = self.f12
            nw109_[int(20)] = p0
            nw109_[int(21)] = d_628_v18_.f12
            nw109_[int(22)] = self.f12
            nw109_[int(23)] = (0) - (p0)
            nw109_[int(24)] = self.f12
            nw109_[int(25)] = p0
            nw109_[int(26)] = p0
            nw109_[int(27)] = self.f12
            nw109_[int(28)] = d_628_v18_.f12
            d_631_v21_ = nw109_
            d_632_v22_: D1
            d_632_v22_ = D1_DC4(d_631_v21_)
            d_633_v23_: _dafny.Map
            d_633_v23_ = _dafny.Map({_dafny.Map({p1: (d_632_v22_).cf7}): (p1 if p1 else p1)})
            d_634_v24_: _dafny.Map
            d_634_v24_ = _dafny.Map({p1: d_631_v21_})
            d_635_v25_: D7
            d_635_v25_ = D7_DC20(d_634_v24_)
            d_636_v26_: _dafny.Seq
            d_636_v26_ = _dafny.SeqWithoutIsStrInference([not(p1)])
            d_633_v23_ = (d_633_v23_).set((d_635_v25_).cf27, not(default__.fm1((_dafny.MultiSet(d_636_v26_)).cardinality, False, p0, p1, globalState)))
            r2 = self.f12
            r0 = p1
            d_637_v27_: str
            d_637_v27_ = _dafny.CodePoint('b')
            d_637_v27_ = (self).f13
            d_638_v28_: _dafny.Array
            nw110_ = _dafny.Array(False, 4)
            d_638_v28_ = nw110_
            index102_ = default__.safeIndex(175, (d_638_v28_).length(0))
            (d_638_v28_)[index102_] = p1
            index103_ = default__.safeIndex(175, (d_638_v28_).length(0))
            (d_638_v28_)[index103_] = p1
        elif True:
            d_639_v29_: _dafny.Seq
            d_639_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cggpcap"))
            d_639_v29_ = d_639_v29_
            d_640_v30_: C0
            nw111_ = C0()
            nw111_.ctor__((_dafny.Map({self.f12: True})).set(25, p1), self.f12)
            d_640_v30_ = nw111_
            d_640_v30_ = d_640_v30_
            if p1:
                d_641_v31_: str
                d_641_v31_ = _dafny.CodePoint('e')
                d_641_v31_ = (self).f13
                d_642_v32_: _dafny.Seq
                d_642_v32_ = _dafny.SeqWithoutIsStrInference([not(p1)])
                d_643_v33_: _dafny.Map
                d_643_v33_ = _dafny.Map({(d_642_v32_)[default__.safeIndex(515, len(d_642_v32_))]: p1})
                d_644_v34_: _dafny.Set
                d_644_v34_ = _dafny.Set({p1, False, p1})
                d_645_v35_: _dafny.Array
                nw112_ = _dafny.Array(None, 22)
                nw112_[int(0)] = p1
                nw112_[int(1)] = (_dafny.Map({p1: p1})) == (d_643_v33_)
                nw112_[int(2)] = default__.fm1(self.f12, p1, p0, p1, globalState)
                nw112_[int(3)] = p1
                nw112_[int(4)] = p1
                nw112_[int(5)] = (p0) >= (p0)
                nw112_[int(6)] = not (p1) or (p1)
                nw112_[int(7)] = p1
                nw112_[int(8)] = p1
                nw112_[int(9)] = p1
                nw112_[int(10)] = p1
                nw112_[int(11)] = p1
                nw112_[int(12)] = p1
                nw112_[int(13)] = p1
                nw112_[int(14)] = p1
                nw112_[int(15)] = (p1 if p1 else p1)
                nw112_[int(16)] = p1
                nw112_[int(17)] = (d_644_v34_).isdisjoint(d_644_v34_)
                nw112_[int(18)] = True
                nw112_[int(19)] = p1
                nw112_[int(20)] = p1
                nw112_[int(21)] = (p1 if p1 else p1)
                d_645_v35_ = nw112_
                index104_ = default__.safeIndex(581, (d_645_v35_).length(0))
                (d_645_v35_)[index104_] = p1
                d_646_v36_: _dafny.Seq
                d_646_v36_ = _dafny.SeqWithoutIsStrInference([-803])
                index105_ = default__.safeIndex(581, (d_645_v35_).length(0))
                (d_645_v35_)[index105_] = (len(_dafny.Set({d_645_v35_}))) != (len(d_646_v36_))
                d_647_v37_: _dafny.Array
                def lambda31_(d_648_v29_, d_649_p2_):
                    def lambda32_(d_650_i2_):
                        return (d_648_v29_) + (d_649_p2_)

                    return lambda32_

                init17_ = lambda31_(d_639_v29_, p2)
                nw113_ = _dafny.Array(None, 5)
                for i0_17_ in range(nw113_.length(0)):
                    nw113_[i0_17_] = init17_(i0_17_)
                d_647_v37_ = nw113_
                index106_ = default__.safeIndex(252, (d_647_v37_).length(0))
                (d_647_v37_)[index106_] = p2
                index107_ = default__.safeIndex(252, (d_647_v37_).length(0))
                index108_ = default__.safeIndex(581, (d_645_v35_).length(0))
                rhs97_ = p1
                rhs98_ = (p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
                rhs99_ = _dafny.CodePoint('d')
                rhs100_ = p1
                lhs57_ = d_647_v37_
                lhs58_ = default__.safeIndex(252, (d_647_v37_).length(0))
                lhs59_ = d_645_v35_
                lhs60_ = default__.safeIndex(581, (d_645_v35_).length(0))
                r0 = rhs97_
                lhs57_[lhs58_] = rhs98_
                d_641_v31_ = rhs99_
                lhs59_[lhs60_] = rhs100_
                d_651_v40_: C0
                nw114_ = C0()
                def iife42_():
                    def iife44_():
                        coll18_ = _dafny.Set()
                        compr_18_: int
                        for compr_18_ in _dafny.IntegerRange(307, 123):
                            d_654_v39_: int = compr_18_
                            if ((307) <= (d_654_v39_)) and ((d_654_v39_) < (123)):
                                coll18_ = coll18_.union(_dafny.Set([default__.safeDivisionInt(d_654_v39_, self.f12)]))
                        return _dafny.Set(coll18_)
                    coll16_ = _dafny.Map()
                    def iife43_():
                        coll17_ = _dafny.Set()
                        compr_17_: int
                        for compr_17_ in _dafny.IntegerRange(307, 123):
                            d_652_v39_: int = compr_17_
                            if ((307) <= (d_652_v39_)) and ((d_652_v39_) < (123)):
                                coll17_ = coll17_.union(_dafny.Set([default__.safeDivisionInt(d_652_v39_, self.f12)]))
                        return _dafny.Set(coll17_)
                    compr_16_: int
                    for compr_16_ in (iife43_()
                    ).Elements:
                        d_653_v38_: int = compr_16_
                        if (d_653_v38_) in (iife44_()
                        ):
                            coll16_[(d_653_v38_) - (p0)] = (d_645_v35_)[default__.safeIndex(581, (d_645_v35_).length(0))]
                    return _dafny.Map(coll16_)
                nw114_.ctor__((d_640_v30_.f16) | (iife42_()
                ), self.f12)
                d_651_v40_ = nw114_
                d_646_v36_ = _dafny.SeqWithoutIsStrInference([self.f12, default__.safeModuloInt(self.f12, 476), p0])
            elif True:
                d_655_v41_: _dafny.Map
                d_655_v41_ = _dafny.Map({True: p1})
                d_656_v42_: _dafny.MultiSet
                d_656_v42_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f13 for d_657_i3_ in range(default__.abs(653))])), len(d_655_v41_)])
                d_658_v43_: _dafny.Seq
                d_658_v43_ = _dafny.SeqWithoutIsStrInference([(0) - (((_dafny.MultiSet([d_640_v30_.f16])).set(d_640_v30_.f16, default__.abs((0) - (p0)))).cardinality), len(_dafny.SeqWithoutIsStrInference([p1, p1]))])
                d_659_v44_: _dafny.Set
                d_659_v44_ = _dafny.Set({d_658_v43_})
                d_660_v45_: _dafny.Map
                d_660_v45_ = _dafny.Map({d_656_v42_: len(d_659_v44_)})
                d_660_v45_ = _dafny.Map({d_656_v42_: self.f12})
                (globalState).f5 = p0
                d_639_v29_ = (p2).set(default__.safeIndex(p0, len(p2)), (self).f13)
                d_661_v46_: _dafny.Array
                nw115_ = _dafny.Array(int(0), 16)
                d_661_v46_ = nw115_
                index109_ = default__.safeIndex(47, (d_661_v46_).length(0))
                (d_661_v46_)[index109_] = p0
                index110_ = default__.safeIndex(47, (d_661_v46_).length(0))
                (d_661_v46_)[index110_] = (0) - (len(d_639_v29_))
                d_662_v47_: _dafny.Map
                d_662_v47_ = _dafny.Map({p1: (self).f13})
                d_663_v48_: _dafny.Array
                nw116_ = _dafny.Array(None, 29)
                nw116_[int(0)] = d_662_v47_
                nw116_[int(1)] = _dafny.Map({p1: _dafny.CodePoint('v')})
                nw116_[int(2)] = d_662_v47_
                nw116_[int(3)] = _dafny.Map({p1: (self).f13})
                nw116_[int(4)] = ((d_662_v47_).set(not(p1), (self).f13)) | (d_662_v47_)
                nw116_[int(5)] = (d_662_v47_) | (d_662_v47_)
                nw116_[int(6)] = d_662_v47_
                nw116_[int(7)] = d_662_v47_
                nw116_[int(8)] = (d_662_v47_) | (d_662_v47_)
                nw116_[int(9)] = d_662_v47_
                nw116_[int(10)] = d_662_v47_
                nw116_[int(11)] = (d_662_v47_) | (d_662_v47_)
                nw116_[int(12)] = (d_662_v47_ if p1 else _dafny.Map({p1: (self).f13}))
                nw116_[int(13)] = (d_662_v47_) | (d_662_v47_)
                nw116_[int(14)] = d_662_v47_
                nw116_[int(15)] = _dafny.Map({p1: (self).f13})
                nw116_[int(16)] = d_662_v47_
                nw116_[int(17)] = (((d_662_v47_).set(p1, (self).f13)).set(p1, _dafny.CodePoint('t'))) | (d_662_v47_)
                nw116_[int(18)] = d_662_v47_
                nw116_[int(19)] = d_662_v47_
                nw116_[int(20)] = d_662_v47_
                nw116_[int(21)] = d_662_v47_
                nw116_[int(22)] = d_662_v47_
                nw116_[int(23)] = d_662_v47_
                nw116_[int(24)] = d_662_v47_
                nw116_[int(25)] = (d_662_v47_).set(False, (self).f13)
                nw116_[int(26)] = d_662_v47_
                nw116_[int(27)] = (d_662_v47_) | (d_662_v47_)
                nw116_[int(28)] = d_662_v47_
                d_663_v48_ = nw116_
                index111_ = default__.safeIndex(706, (d_663_v48_).length(0))
                (d_663_v48_)[index111_] = (_dafny.Map({p1: (self).f13})).set(p1, (self).f13)
                index112_ = default__.safeIndex(706, (d_663_v48_).length(0))
                (d_663_v48_)[index112_] = d_662_v47_
            d_664_v49_: _dafny.Array
            nw117_ = _dafny.Array(D6.default()(), 6)
            d_664_v49_ = nw117_
            d_665_v50_: _dafny.Map
            d_665_v50_ = _dafny.Map({p1: p1})
            d_666_v51_: _dafny.Array
            nw118_ = _dafny.Array(None, 28)
            nw118_[int(0)] = p1
            nw118_[int(1)] = p1
            nw118_[int(2)] = p1
            nw118_[int(3)] = False
            nw118_[int(4)] = False
            nw118_[int(5)] = p1
            nw118_[int(6)] = False
            nw118_[int(7)] = p1
            nw118_[int(8)] = p1
            nw118_[int(9)] = p1
            nw118_[int(10)] = p1
            nw118_[int(11)] = p1
            nw118_[int(12)] = True
            nw118_[int(13)] = p1
            nw118_[int(14)] = p1
            nw118_[int(15)] = p1
            nw118_[int(16)] = p1
            nw118_[int(17)] = p1
            nw118_[int(18)] = p1
            nw118_[int(19)] = p1
            nw118_[int(20)] = ((d_665_v50_)[p1] if (p1) in (d_665_v50_) else p1)
            nw118_[int(21)] = p1
            nw118_[int(22)] = p1
            nw118_[int(23)] = p1
            nw118_[int(24)] = p1
            nw118_[int(25)] = False
            nw118_[int(26)] = not(False)
            nw118_[int(27)] = p1
            d_666_v51_ = nw118_
            d_667_v52_: D6
            d_667_v52_ = D6_DC16(d_666_v51_, self.f12)
            d_668_v53_: D6
            d_668_v53_ = D6_DC19(d_667_v52_)
            index113_ = default__.safeIndex(651, (d_664_v49_).length(0))
            (d_664_v49_)[index113_] = d_668_v53_
            d_669_v54_: _dafny.Array
            def lambda33_(d_670_i4_):
                return (self).f13

            init18_ = lambda33_
            nw119_ = _dafny.Array(None, 10)
            for i0_18_ in range(nw119_.length(0)):
                nw119_[i0_18_] = init18_(i0_18_)
            d_669_v54_ = nw119_
            index114_ = default__.safeIndex(233, (d_669_v54_).length(0))
            (d_669_v54_)[index114_] = _dafny.CodePoint('l')
            d_671_v55_: _dafny.Map
            d_671_v55_ = _dafny.Map({p0: 864})
            index115_ = default__.safeIndex(651, (d_664_v49_).length(0))
            index116_ = default__.safeIndex(233, (d_669_v54_).length(0))
            rhs101_ = p1
            rhs102_ = d_668_v53_
            rhs103_ = _dafny.CodePoint('k')
            rhs104_ = (d_671_v55_) | (d_671_v55_)
            lhs61_ = d_664_v49_
            lhs62_ = default__.safeIndex(651, (d_664_v49_).length(0))
            lhs63_ = d_669_v54_
            lhs64_ = default__.safeIndex(233, (d_669_v54_).length(0))
            r1 = rhs101_
            lhs61_[lhs62_] = rhs102_
            lhs63_[lhs64_] = rhs103_
            d_671_v55_ = rhs104_
            d_672_v56_: _dafny.Seq
            d_672_v56_ = _dafny.SeqWithoutIsStrInference([p1])
            d_673_v57_: _dafny.Seq
            d_673_v57_ = _dafny.SeqWithoutIsStrInference([d_672_v56_])
            d_674_v58_: _dafny.Array
            nw120_ = _dafny.Array(None, 9)
            nw120_[int(0)] = d_672_v56_
            nw120_[int(1)] = d_672_v56_
            nw120_[int(2)] = _dafny.SeqWithoutIsStrInference([p1])
            nw120_[int(3)] = d_672_v56_
            nw120_[int(4)] = d_672_v56_
            nw120_[int(5)] = d_672_v56_
            nw120_[int(6)] = (d_673_v57_)[default__.safeIndex(len((d_673_v57_).set(default__.safeIndex(p0, len(d_673_v57_)), default__.fm9(p1, globalState))), len(d_673_v57_))]
            nw120_[int(7)] = d_672_v56_
            nw120_[int(8)] = d_672_v56_
            d_674_v58_ = nw120_
            d_675_v59_: _dafny.Seq
            d_675_v59_ = _dafny.SeqWithoutIsStrInference([d_674_v58_])
            d_676_v60_: _dafny.Array
            nw121_ = _dafny.Array(None, 6)
            nw121_[int(0)] = d_674_v58_
            nw121_[int(1)] = d_674_v58_
            nw121_[int(2)] = d_674_v58_
            nw121_[int(3)] = d_674_v58_
            nw121_[int(4)] = d_674_v58_
            nw121_[int(5)] = (d_675_v59_)[default__.safeIndex((0) - (self.f12), len(d_675_v59_))]
            d_676_v60_ = nw121_
            index117_ = default__.safeIndex(871, (d_676_v60_).length(0))
            (d_676_v60_)[index117_] = d_674_v58_
            index118_ = default__.safeIndex(871, (d_676_v60_).length(0))
            (d_676_v60_)[index118_] = d_674_v58_
        d_677_v61_: _dafny.MultiSet
        d_677_v61_ = _dafny.MultiSet([p0])
        d_678_i5_: int
        d_678_i5_ = 0
        with _dafny.label("5"):
            while (p0) not in (d_677_v61_):
                with _dafny.c_label("5"):
                    if (d_678_i5_) >= (100):
                        raise _dafny.Break("5")
                    d_678_i5_ = (d_678_i5_) + (1)
                    rhs105_ = (0) - (self.f12)
                    rhs106_ = self.f12
                    lhs65_ = globalState
                    lhs65_.f5 = rhs105_
                    r2 = rhs106_
                    d_679_v62_: D5
                    d_679_v62_ = D5_DC14(p1)
                    source8_ = d_679_v62_
                    if source8_.is_DC13:
                        d_680___mcc_h0_ = source8_.cf18
                        d_681___mcc_h1_ = source8_.cf19
                        d_682_cf19_ = d_681___mcc_h1_
                        d_683_cf18_ = d_680___mcc_h0_
                        d_684_v63_: _dafny.Map
                        d_684_v63_ = _dafny.Map({self.f12: not(p1)})
                        d_683_cf18_ = ((d_684_v63_)[self.f12] if (self.f12) in (d_684_v63_) else not(d_682_cf19_))
                        r1 = not(not(p1))
                        (globalState).f5 = (507 if d_682_cf19_ else self.f12)
                        d_685_v64_: _dafny.Array
                        nw122_ = _dafny.Array(int(0), 25)
                        d_685_v64_ = nw122_
                        index119_ = default__.safeIndex(197, (d_685_v64_).length(0))
                        (d_685_v64_)[index119_] = default__.safeModuloInt(p0, len(p2))
                        index120_ = default__.safeIndex(197, (d_685_v64_).length(0))
                        (d_685_v64_)[index120_] = (0) - (p0)
                    elif source8_.is_DC14:
                        d_686___mcc_h2_ = source8_.cf20
                        d_687_cf20_ = d_686___mcc_h2_
                        d_688_v65_: _dafny.Seq
                        d_688_v65_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_689_v66_: _dafny.Array
                        nw123_ = _dafny.Array(None, 21)
                        nw123_[int(0)] = not(d_687_cf20_)
                        nw123_[int(1)] = p1
                        nw123_[int(2)] = d_687_cf20_
                        nw123_[int(3)] = p1
                        nw123_[int(4)] = not(p1)
                        nw123_[int(5)] = d_687_cf20_
                        nw123_[int(6)] = p1
                        nw123_[int(7)] = not((p0) < (p0))
                        nw123_[int(8)] = p1
                        nw123_[int(9)] = not(d_687_cf20_)
                        nw123_[int(10)] = p1
                        nw123_[int(11)] = d_687_cf20_
                        nw123_[int(12)] = (d_688_v65_)[default__.safeIndex(p0, len(d_688_v65_))]
                        nw123_[int(13)] = d_687_cf20_
                        nw123_[int(14)] = not (p1) or (p1)
                        nw123_[int(15)] = (False if p1 else d_687_cf20_)
                        nw123_[int(16)] = d_687_cf20_
                        nw123_[int(17)] = not((d_687_cf20_ if d_687_cf20_ else p1))
                        nw123_[int(18)] = (self.f12) >= (self.f12)
                        nw123_[int(19)] = (p1) == (d_687_cf20_)
                        nw123_[int(20)] = d_687_cf20_
                        d_689_v66_ = nw123_
                        index121_ = default__.safeIndex(866, (d_689_v66_).length(0))
                        (d_689_v66_)[index121_] = (p0) == (p0)
                        d_690_v67_: _dafny.Seq
                        d_690_v67_ = _dafny.SeqWithoutIsStrInference([850])
                        index122_ = default__.safeIndex(866, (d_689_v66_).length(0))
                        (d_689_v66_)[index122_] = (_dafny.SeqWithoutIsStrInference([self.f12])) < ((d_690_v67_ if p1 else (self).fm5(p1, d_687_cf20_, d_687_cf20_, not(p1), globalState)))
                        (globalState).f5 = p0
                        d_691_v68_: D2
                        d_691_v68_ = D2_DC8(d_687_cf20_, p0, d_687_cf20_)
                        index123_ = default__.safeIndex(866, (d_689_v66_).length(0))
                        (d_689_v66_)[index123_] = not(default__.fm1(757, default__.fm1(self.f12, (d_689_v66_)[default__.safeIndex(866, (d_689_v66_).length(0))], self.f12, (d_689_v66_)[default__.safeIndex(866, (d_689_v66_).length(0))], globalState), len(_dafny.Map({(0) - (-788): p1})), (d_691_v68_).cf13, globalState))
                        d_692_v69_: _dafny.Set
                        d_692_v69_ = _dafny.Set({_dafny.Map({d_687_cf20_: not(default__.fm1(-745, (d_689_v66_)[default__.safeIndex(866, (d_689_v66_).length(0))], p0, d_687_cf20_, globalState))})})
                        d_693_v71_: _dafny.MultiSet
                        d_693_v71_ = _dafny.MultiSet([_dafny.CodePoint('j')])
                        d_694_v72_: _dafny.Array
                        nw124_ = _dafny.Array(None, 17)
                        nw124_[int(0)] = 176
                        nw124_[int(1)] = self.f12
                        nw124_[int(2)] = default__.safeModuloInt(p0, (0) - (self.f12))
                        nw124_[int(3)] = ((0) - (len(d_692_v69_)) if False else self.f12)
                        nw124_[int(4)] = default__.safeModuloInt(p0, 440)
                        nw124_[int(5)] = p0
                        nw124_[int(6)] = self.f12
                        nw124_[int(7)] = p0
                        nw124_[int(8)] = self.f12
                        def iife45_():
                            coll19_ = _dafny.Map()
                            compr_19_: str
                            for compr_19_ in (d_693_v71_).Elements:
                                d_695_v70_: str = compr_19_
                                if (d_695_v70_) in (d_693_v71_):
                                    coll19_[d_695_v70_] = d_687_cf20_
                            return _dafny.Map(coll19_)
                        nw124_[int(9)] = len(iife45_()
                        )
                        nw124_[int(10)] = p0
                        nw124_[int(11)] = (0) - (-239)
                        nw124_[int(12)] = (self.f12) + ((_dafny.MultiSet(d_688_v65_)).cardinality)
                        nw124_[int(13)] = (self.f12) + (p0)
                        nw124_[int(14)] = self.f12
                        nw124_[int(15)] = 536
                        nw124_[int(16)] = self.f12
                        d_694_v72_ = nw124_
                        d_694_v72_ = d_694_v72_
                    elif True:
                        d_696___mcc_h3_ = source8_.cf17
                        d_697_cf17_ = d_696___mcc_h3_
                        d_698_v73_: C1
                        nw125_ = C1()
                        nw125_.ctor__((self).f13, (self).f13, self.f12)
                        d_698_v73_ = nw125_
                        d_699_v74_: _dafny.Array
                        def lambda34_(d_700_p2_):
                            def lambda35_(d_701_i6_):
                                return default__.safeDivisionInt(d_701_i6_, len(_dafny.Set({len(d_700_p2_)})))

                            return lambda35_

                        init19_ = lambda34_(p2)
                        nw126_ = _dafny.Array(None, 28)
                        for i0_19_ in range(nw126_.length(0)):
                            nw126_[i0_19_] = init19_(i0_19_)
                        d_699_v74_ = nw126_
                        index124_ = default__.safeIndex(985, (d_699_v74_).length(0))
                        (d_699_v74_)[index124_] = 299
                        d_702_v75_: str
                        d_702_v75_ = _dafny.CodePoint('m')
                        index125_ = default__.safeIndex(985, (d_699_v74_).length(0))
                        rhs107_ = d_698_v73_
                        rhs108_ = p0
                        rhs109_ = (d_698_v73_).f15
                        lhs66_ = d_699_v74_
                        lhs67_ = default__.safeIndex(985, (d_699_v74_).length(0))
                        d_698_v73_ = rhs107_
                        lhs66_[lhs67_] = rhs108_
                        d_702_v75_ = rhs109_
                        d_702_v75_ = (d_698_v73_).f15
                        (globalState).f5 = (d_698_v73_).fm4(globalState)
                        d_703_v76_: int
                        d_704_v77_: _dafny.Map
                        d_705_v78_: _dafny.Seq
                        out42_: int
                        out43_: _dafny.Map
                        out44_: _dafny.Seq
                        out42_, out43_, out44_ = default__.m0((d_699_v74_)[default__.safeIndex(985, (d_699_v74_).length(0))], ((d_699_v74_)[default__.safeIndex(985, (d_699_v74_).length(0))]) * (self.f12), ((d_699_v74_)[default__.safeIndex(985, (d_699_v74_).length(0))]) + ((d_699_v74_)[default__.safeIndex(985, (d_699_v74_).length(0))]), True, globalState)
                        d_703_v76_ = out42_
                        d_704_v77_ = out43_
                        d_705_v78_ = out44_
                    d_706_v79_: _dafny.Map
                    d_706_v79_ = _dafny.Map({(0) - (self.f12): p2})
                    d_707_v80_: _dafny.Seq
                    d_707_v80_ = _dafny.SeqWithoutIsStrInference([True])
                    d_706_v79_ = (d_706_v79_).set(len(d_707_v80_), (p2) + (p2))
                    d_708_v81_: _dafny.Array
                    nw127_ = _dafny.Array(int(0), 3)
                    d_708_v81_ = nw127_
                    rhs110_ = d_708_v81_
                    rhs111_ = (0) - (default__.fm0(p0, self.f12, globalState))
                    lhs68_ = self
                    d_708_v81_ = rhs110_
                    lhs68_.f12 = rhs111_
                    pass
            pass
        hi3_ = self.f12
        for d_709_i7_ in range((-779) * (357), hi3_):
            r0 = (_dafny.MultiSet([p0])).issubset(d_677_v61_)
            d_710_v82_: C0
            nw128_ = C0()
            nw128_.ctor__(_dafny.Map({d_709_i7_: p1}), d_709_i7_)
            d_710_v82_ = nw128_
            d_711_v83_: _dafny.Map
            d_711_v83_ = _dafny.Map({p1: p1})
            d_712_v84_: _dafny.Seq
            d_712_v84_ = _dafny.SeqWithoutIsStrInference([p1, ((d_711_v83_)[not(p1)] if (not(p1)) in (d_711_v83_) else p1), p1, p1, p1])
            d_713_v85_: _dafny.Seq
            d_713_v85_ = _dafny.SeqWithoutIsStrInference([d_712_v84_])
            if (len(d_713_v85_)) >= (self.f12):
                d_714_v86_: _dafny.Array
                nw129_ = _dafny.Array(False, 22)
                d_714_v86_ = nw129_
                index126_ = default__.safeIndex(844, (d_714_v86_).length(0))
                (d_714_v86_)[index126_] = p1
                index127_ = default__.safeIndex(844, (d_714_v86_).length(0))
                (d_714_v86_)[index127_] = (p2) != (p2)
                d_715_v87_: _dafny.Array
                nw130_ = _dafny.Array(int(0), 10)
                d_715_v87_ = nw130_
                index128_ = default__.safeIndex(923, (d_715_v87_).length(0))
                (d_715_v87_)[index128_] = d_709_i7_
                d_716_v88_: _dafny.Set
                d_716_v88_ = _dafny.Set({False, (d_714_v86_)[default__.safeIndex(844, (d_714_v86_).length(0))]})
                d_717_v89_: _dafny.Map
                d_717_v89_ = _dafny.Map({p1: d_716_v88_})
                index129_ = default__.safeIndex(923, (d_715_v87_).length(0))
                (d_715_v87_)[index129_] = len(((d_717_v89_)[(self.f12) == (d_709_i7_)] if ((self.f12) == (d_709_i7_)) in (d_717_v89_) else d_716_v88_))
                index130_ = default__.safeIndex(844, (d_714_v86_).length(0))
                (d_714_v86_)[index130_] = ((D0_DC0(p1)).cf0) == ((d_714_v86_)[default__.safeIndex(844, (d_714_v86_).length(0))])
                r1 = (d_714_v86_)[default__.safeIndex(844, (d_714_v86_).length(0))]
                d_718_v90_: _dafny.Seq
                d_718_v90_ = _dafny.SeqWithoutIsStrInference([796])
                d_719_v91_: _dafny.MultiSet
                d_719_v91_ = _dafny.MultiSet([d_718_v90_])
                index131_ = default__.safeIndex(923, (d_715_v87_).length(0))
                index132_ = default__.safeIndex(923, (d_715_v87_).length(0))
                rhs112_ = d_709_i7_
                rhs113_ = ((d_719_v91_)[(d_718_v90_) + (d_718_v90_)] if ((d_718_v90_) + (d_718_v90_)) in (d_719_v91_) else p0)
                rhs114_ = d_709_i7_
                lhs69_ = d_715_v87_
                lhs70_ = default__.safeIndex(923, (d_715_v87_).length(0))
                lhs71_ = globalState
                lhs72_ = d_715_v87_
                lhs73_ = default__.safeIndex(923, (d_715_v87_).length(0))
                lhs69_[lhs70_] = rhs112_
                lhs71_.f5 = rhs113_
                lhs72_[lhs73_] = rhs114_
            elif True:
                (globalState).f5 = (0) - (d_709_i7_)
                d_720_v92_: _dafny.MultiSet
                d_720_v92_ = _dafny.MultiSet([p1, p1])
                d_720_v92_ = _dafny.MultiSet([p1])
                d_721_v93_: _dafny.Array
                nw131_ = _dafny.Array(False, 22)
                d_721_v93_ = nw131_
                index133_ = default__.safeIndex(5, (d_721_v93_).length(0))
                (d_721_v93_)[index133_] = default__.fm1(self.f12, p1, d_709_i7_, False, globalState)
                index134_ = default__.safeIndex(5, (d_721_v93_).length(0))
                (d_721_v93_)[index134_] = p1
                d_722_v94_: str
                d_722_v94_ = _dafny.CodePoint('j')
                d_722_v94_ = (self).f13
                d_723_v95_: int
                d_724_v96_: _dafny.Map
                d_725_v97_: _dafny.Seq
                out45_: int
                out46_: _dafny.Map
                out47_: _dafny.Seq
                out45_, out46_, out47_ = default__.m0(self.f12, (0) - (default__.fm0(p0, p0, globalState)), 406, (d_721_v93_)[default__.safeIndex(5, (d_721_v93_).length(0))], globalState)
                d_723_v95_ = out45_
                d_724_v96_ = out46_
                d_725_v97_ = out47_
            d_726_v98_: _dafny.Seq
            d_726_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_726_v98_ = (d_726_v98_) + ((_dafny.SeqWithoutIsStrInference([(self).f13 for d_727_i8_ in range(default__.abs(831))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxb"))))
        d_728_v99_: _dafny.Array
        def lambda36_(d_729_p1_):
            def lambda37_(d_730_i9_):
                return d_729_p1_

            return lambda37_

        init20_ = lambda36_(p1)
        nw132_ = _dafny.Array(None, 16)
        for i0_20_ in range(nw132_.length(0)):
            nw132_[i0_20_] = init20_(i0_20_)
        d_728_v99_ = nw132_
        d_731_v100_: D0
        d_731_v100_ = D0_DC1(d_728_v99_)
        source9_ = d_731_v100_
        if source9_.is_DC1:
            d_732___mcc_h4_ = source9_.cf1
            d_733_cf1_ = d_732___mcc_h4_
            index135_ = default__.safeIndex(979, (d_728_v99_).length(0))
            (d_728_v99_)[index135_] = not(not(p1))
            index136_ = default__.safeIndex(979, (d_728_v99_).length(0))
            (d_728_v99_)[index136_] = p1
            d_734_v101_: _dafny.Seq
            d_734_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uadjarwl"))
            d_734_v101_ = d_734_v101_
            d_735_v102_: D2
            d_735_v102_ = D2_DC9(p0)
            d_736_v103_: _dafny.Set
            d_736_v103_ = default__.fm14(-662, d_735_v102_, globalState)
            source10_ = d_736_v103_
            d_737___mcc_h11_ = source10_
            d_738_cf16_ = d_737___mcc_h11_
            d_739_v104_: _dafny.Map
            d_739_v104_ = _dafny.Map({not(p1): True})
            d_739_v104_ = (d_739_v104_).set(default__.fm1(p0, p1, p0, p1, globalState), (d_728_v99_)[default__.safeIndex(979, (d_728_v99_).length(0))])
            d_740_v105_: _dafny.Seq
            d_740_v105_ = _dafny.SeqWithoutIsStrInference([self.f12, p0, p0, (self.f12) * (p0)])
            d_740_v105_ = d_740_v105_
            d_677_v61_ = d_677_v61_
            d_741_v106_: T1
            nw133_ = C1()
            nw133_.ctor__(_dafny.CodePoint('l'), (self).f13, 738)
            d_741_v106_ = nw133_
            d_741_v106_ = d_741_v106_
            index137_ = default__.safeIndex(979, (d_728_v99_).length(0))
            rhs115_ = (len(d_734_v101_)) * (p0)
            rhs116_ = p1
            rhs117_ = default__.safeDivisionInt(self.f12, 979)
            lhs74_ = d_728_v99_
            lhs75_ = default__.safeIndex(979, (d_728_v99_).length(0))
            lhs76_ = self
            r2 = rhs115_
            lhs74_[lhs75_] = rhs116_
            lhs76_.f12 = rhs117_
        elif source9_.is_DC2:
            d_742___mcc_h5_ = source9_.cf2
            d_743___mcc_h6_ = source9_.cf3
            d_744___mcc_h7_ = source9_.cf4
            d_745___mcc_h8_ = source9_.cf5
            d_746_cf5_ = d_745___mcc_h8_
            d_747_cf4_ = d_744___mcc_h7_
            d_748_cf3_ = d_743___mcc_h6_
            d_749_cf2_ = d_742___mcc_h5_
            r0 = (d_677_v61_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f12, p0, self.f12])), (0) - (len(p2))])))
            d_750_v107_: T0
            nw134_ = C2()
            nw134_.ctor__(d_749_cf2_, 369)
            d_750_v107_ = nw134_
            rhs118_ = d_750_v107_
            rhs119_ = d_749_cf2_
            d_750_v107_ = rhs118_
            d_749_cf2_ = rhs119_
            d_751_v108_: C1
            nw135_ = C1()
            nw135_.ctor__((self).f13, (self).f13, self.f12)
            d_751_v108_ = nw135_
            if not(True):
                index138_ = default__.safeIndex(100, (d_728_v99_).length(0))
                (d_728_v99_)[index138_] = d_749_cf2_
                index139_ = default__.safeIndex(100, (d_728_v99_).length(0))
                (d_728_v99_)[index139_] = not(True)
                (globalState).f5 = len(p2)
                d_752_v109_: _dafny.Set
                d_752_v109_ = _dafny.Set({d_749_cf2_, default__.fm1(p0, d_746_cf5_, p0, (d_728_v99_)[default__.safeIndex(100, (d_728_v99_).length(0))], globalState)})
                d_753_v110_: _dafny.Set
                d_753_v110_ = d_752_v109_
                rhs120_ = p0
                rhs121_ = d_753_v110_
                lhs77_ = d_750_v107_
                lhs77_.f12 = rhs120_
                d_753_v110_ = rhs121_
                d_754_v111_: C0
                nw136_ = C0()
                nw136_.ctor__((d_748_cf3_) | (_dafny.Map({len(p2): (d_728_v99_)[default__.safeIndex(100, (d_728_v99_).length(0))]})), p0)
                d_754_v111_ = nw136_
                d_755_v112_: _dafny.Array
                def lambda38_(d_756_cf5_):
                    def lambda39_(d_757_i10_):
                        return d_756_cf5_

                    return lambda39_

                init21_ = lambda38_(d_746_cf5_)
                nw137_ = _dafny.Array(None, 28)
                for i0_21_ in range(nw137_.length(0)):
                    nw137_[i0_21_] = init21_(i0_21_)
                d_755_v112_ = nw137_
                d_758_v113_: D6
                d_758_v113_ = D6_DC16(d_755_v112_, d_750_v107_.f12)
                d_759_v115_: _dafny.Seq
                d_759_v115_ = _dafny.SeqWithoutIsStrInference([False])
                d_760_v116_: _dafny.Map
                d_760_v116_ = _dafny.Map({(d_728_v99_)[default__.safeIndex(100, (d_728_v99_).length(0))]: p0})
                d_761_v117_: _dafny.Seq
                d_761_v117_ = _dafny.SeqWithoutIsStrInference([d_750_v107_.f12, 471, self.f12])
                index140_ = default__.safeIndex(100, (d_728_v99_).length(0))
                index141_ = default__.safeIndex(100, (d_728_v99_).length(0))
                def iife46_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(945, -119):
                        d_762_v114_: int = compr_20_
                        if ((945) <= (d_762_v114_)) and ((d_762_v114_) < (-119)):
                            coll20_[default__.safeModuloInt(d_762_v114_, d_750_v107_.f12)] = d_750_v107_.f12
                    return _dafny.Map(coll20_)
                rhs122_ = (d_758_v113_).cf23
                rhs123_ = ((0) - ((0) - (len(iife46_()
                )))) < (d_750_v107_.f12)
                rhs124_ = (d_749_cf2_ if not ((d_759_v115_)[default__.safeIndex(len(d_760_v116_), len(d_759_v115_))]) or (False) else default__.fm1(len(d_761_v117_), p1, self.f12, default__.fm1(610, (d_728_v99_)[default__.safeIndex(100, (d_728_v99_).length(0))], d_750_v107_.f12, d_749_cf2_, globalState), globalState))
                lhs78_ = d_750_v107_
                lhs79_ = d_728_v99_
                lhs80_ = default__.safeIndex(100, (d_728_v99_).length(0))
                lhs81_ = d_728_v99_
                lhs82_ = default__.safeIndex(100, (d_728_v99_).length(0))
                lhs78_.f12 = rhs122_
                lhs79_[lhs80_] = rhs123_
                lhs81_[lhs82_] = rhs124_
            elif True:
                d_763_v118_: _dafny.Array
                nw138_ = _dafny.Array(int(0), 11)
                d_763_v118_ = nw138_
                d_764_v119_: _dafny.Set
                d_764_v119_ = _dafny.Set({d_748_cf3_})
                rhs125_ = d_763_v118_
                rhs126_ = default__.safeModuloInt(len(d_764_v119_), 808)
                lhs83_ = d_750_v107_
                d_763_v118_ = rhs125_
                lhs83_.f12 = rhs126_
                r2 = d_750_v107_.f12
                r1 = not(not(d_749_cf2_))
                d_765_v120_: _dafny.Seq
                d_765_v120_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbjki"))
                d_765_v120_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))) + (p2)
                d_766_v121_: _dafny.Seq
                d_766_v121_ = _dafny.SeqWithoutIsStrInference([d_750_v107_.f12])
                d_767_v122_: _dafny.Map
                d_767_v122_ = _dafny.Map({len(d_766_v121_): (d_751_v108_).f15})
                d_768_v123_: _dafny.MultiSet
                d_768_v123_ = _dafny.MultiSet([_dafny.Map({(0) - (p0): (self).f13}), d_767_v122_, (d_767_v122_).set(self.f12, _dafny.CodePoint('h')), (d_767_v122_) | (d_767_v122_), (d_767_v122_) | (d_767_v122_)])
                d_768_v123_ = d_768_v123_
        elif source9_.is_DC0:
            d_769___mcc_h9_ = source9_.cf0
            d_770_cf0_ = d_769___mcc_h9_
            d_771_v124_: _dafny.MultiSet
            d_771_v124_ = _dafny.MultiSet([d_770_cf0_])
            d_772_v125_: _dafny.Set
            d_772_v125_ = _dafny.Set({(d_771_v124_).issubset(d_771_v124_)})
            d_772_v125_ = d_772_v125_
            index142_ = default__.safeIndex(537, (d_728_v99_).length(0))
            (d_728_v99_)[index142_] = d_770_cf0_
            index143_ = default__.safeIndex(537, (d_728_v99_).length(0))
            (d_728_v99_)[index143_] = d_770_cf0_
            d_773_v126_: _dafny.Array
            nw139_ = _dafny.Array(False, 19)
            d_773_v126_ = nw139_
            d_774_v127_: D6
            d_774_v127_ = D6_DC16(d_773_v126_, p0)
            source11_ = d_774_v127_
            if source11_.is_DC16:
                d_775___mcc_h12_ = source11_.cf22
                d_776___mcc_h13_ = source11_.cf23
                d_777_cf23_ = d_776___mcc_h13_
                d_778_cf22_ = d_775___mcc_h12_
                d_779_v128_: _dafny.Array
                nw140_ = _dafny.Array(int(0), 15)
                d_779_v128_ = nw140_
                index144_ = default__.safeIndex(718, (d_779_v128_).length(0))
                (d_779_v128_)[index144_] = d_777_cf23_
                d_780_v129_: _dafny.Seq
                d_780_v129_ = _dafny.SeqWithoutIsStrInference([d_778_cf22_])
                index145_ = default__.safeIndex(718, (d_779_v128_).length(0))
                (d_779_v128_)[index145_] = (len((d_780_v129_) + (d_780_v129_))) * (p0)
                d_781_v130_: _dafny.Map
                d_781_v130_ = _dafny.Map({(d_779_v128_)[default__.safeIndex(718, (d_779_v128_).length(0))]: default__.safeDivisionInt(p0, p0)})
                index146_ = default__.safeIndex(718, (d_779_v128_).length(0))
                (d_779_v128_)[index146_] = len(d_781_v130_)
                d_782_v131_: D2
                d_782_v131_ = D2_DC8(d_770_cf0_, 288, d_770_cf0_)
                (self).f12 = default__.fm0(default__.safeModuloInt((d_779_v128_)[default__.safeIndex(718, (d_779_v128_).length(0))], (d_782_v131_).cf12), 37, globalState)
                index147_ = default__.safeIndex(537, (d_728_v99_).length(0))
                (d_728_v99_)[index147_] = (_dafny.SeqWithoutIsStrInference([d_778_cf22_, d_773_v126_])) <= (d_780_v129_)
            elif source11_.is_DC17:
                d_783___mcc_h14_ = source11_.cf24
                d_784_cf24_ = d_783___mcc_h14_
                r0 = (self.f12) > (self.f12)
                d_785_v132_: T0
                nw141_ = C2()
                nw141_.ctor__(p1, self.f12)
                d_785_v132_ = nw141_
                d_786_v133_: _dafny.Map
                d_786_v133_ = _dafny.Map({d_785_v132_: 919})
                d_787_v135_: _dafny.MultiSet
                d_787_v135_ = _dafny.MultiSet([(self).f13, (self).f13])
                d_788_v136_: C2
                nw142_ = C2()
                def iife47_():
                    coll21_ = _dafny.Map()
                    compr_21_: str
                    for compr_21_ in (d_787_v135_).Elements:
                        d_789_v134_: str = compr_21_
                        if (d_789_v134_) in (d_787_v135_):
                            coll21_[d_789_v134_] = p0
                    return _dafny.Map(coll21_)
                nw142_.ctor__(False, len(iife47_()
                ))
                d_788_v136_ = nw142_
                d_786_v133_ = (d_786_v133_).set(d_785_v132_, default__.safeDivisionInt(len(_dafny.Map({-507: d_788_v136_})), p0))
                d_790_v137_: _dafny.Seq
                d_790_v137_ = _dafny.SeqWithoutIsStrInference([p1])
                d_791_v138_: int
                d_792_v139_: bool
                d_793_v140_: int
                out48_: int
                out49_: bool
                out50_: int
                out48_, out49_, out50_ = (self).m3((_dafny.MultiSet(d_790_v137_) if d_770_cf0_ else d_771_v124_), globalState)
                d_791_v138_ = out48_
                d_792_v139_ = out49_
                d_793_v140_ = out50_
                d_794_v141_: _dafny.Seq
                d_794_v141_ = _dafny.SeqWithoutIsStrInference([p0, self.f12])
                d_795_v142_: _dafny.MultiSet
                d_795_v142_ = _dafny.MultiSet([d_794_v141_, d_794_v141_, d_794_v141_])
                index148_ = default__.safeIndex(537, (d_728_v99_).length(0))
                (d_728_v99_)[index148_] = ((d_795_v142_).intersection(d_795_v142_)) == (d_795_v142_)
            elif source11_.is_DC18:
                d_796___mcc_h15_ = source11_.cf25
                d_797_cf25_ = d_796___mcc_h15_
                d_797_cf25_ = (0) - (d_797_cf25_)
                d_798_v143_: _dafny.Seq
                d_798_v143_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                d_798_v143_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kj"))
                index149_ = default__.safeIndex(537, (d_728_v99_).length(0))
                (d_728_v99_)[index149_] = d_770_cf0_
                d_799_v144_: str
                d_799_v144_ = _dafny.CodePoint('l')
                d_800_v145_: _dafny.Map
                d_800_v145_ = _dafny.Map({self.f12: d_770_cf0_})
                d_801_v146_: _dafny.Map
                d_801_v146_ = _dafny.Map({d_770_cf0_: d_797_cf25_})
                d_802_v147_: C0
                nw143_ = C0()
                nw143_.ctor__(d_800_v145_, (95) * (len(d_801_v146_)))
                d_802_v147_ = nw143_
                d_803_v148_: D6
                d_803_v148_ = D6_DC18(default__.safeModuloInt(self.f12, ((d_771_v124_)[(d_728_v99_)[default__.safeIndex(537, (d_728_v99_).length(0))]] if ((d_728_v99_)[default__.safeIndex(537, (d_728_v99_).length(0))]) in (d_771_v124_) else d_797_cf25_)))
                rhs127_ = ((d_677_v61_)[900] if (900) in (d_677_v61_) else self.f12)
                rhs128_ = d_773_v126_
                rhs129_ = d_799_v144_
                rhs130_ = d_802_v147_
                rhs131_ = d_803_v148_
                lhs84_ = self
                lhs84_.f12 = rhs127_
                d_773_v126_ = rhs128_
                d_799_v144_ = rhs129_
                d_802_v147_ = rhs130_
                d_803_v148_ = rhs131_
            elif source11_.is_DC15:
                d_804___mcc_h16_ = source11_.cf21
                d_805_cf21_ = d_804___mcc_h16_
                d_806_v149_: _dafny.MultiSet
                d_806_v149_ = _dafny.MultiSet([p2, p2])
                d_770_cf0_ = (d_806_v149_).isdisjoint((d_806_v149_) - (d_806_v149_))
                d_807_v150_: C1
                nw144_ = C1()
                nw144_.ctor__((self).f13, (self).f13, -857)
                d_807_v150_ = nw144_
                d_808_v151_: D6
                d_808_v151_ = D6_DC15((d_805_cf21_) + (d_805_cf21_))
                d_808_v151_ = D6_DC15(_dafny.SeqWithoutIsStrInference([len(d_805_cf21_) for d_809_i11_ in range(default__.abs(872))]))
                d_810_v152_: _dafny.Array
                def lambda40_(d_811_v99_):
                    def lambda41_(d_812_i12_):
                        return _dafny.SeqWithoutIsStrInference([(d_811_v99_)[default__.safeIndex(537, (d_811_v99_).length(0))]])

                    return lambda41_

                init22_ = lambda40_(d_728_v99_)
                nw145_ = _dafny.Array(None, 11)
                for i0_22_ in range(nw145_.length(0)):
                    nw145_[i0_22_] = init22_(i0_22_)
                d_810_v152_ = nw145_
                d_813_v153_: _dafny.Map
                d_813_v153_ = _dafny.Map({(0) - (self.f12): self.f12})
                d_814_v154_: _dafny.Set
                d_814_v154_ = _dafny.Set({(0) - (len(d_813_v153_))})
                index150_ = default__.safeIndex(537, (d_728_v99_).length(0))
                rhs132_ = (d_814_v154_).issubset(d_814_v154_)
                rhs133_ = (d_728_v99_)[default__.safeIndex(537, (d_728_v99_).length(0))]
                rhs134_ = (0) - (p0)
                rhs135_ = d_810_v152_
                rhs136_ = False
                lhs85_ = d_728_v99_
                lhs86_ = default__.safeIndex(537, (d_728_v99_).length(0))
                lhs87_ = globalState
                r0 = rhs132_
                lhs85_[lhs86_] = rhs133_
                lhs87_.f5 = rhs134_
                d_810_v152_ = rhs135_
                d_770_cf0_ = rhs136_
            elif True:
                d_815___mcc_h17_ = source11_.cf26
                d_816_cf26_ = d_815___mcc_h17_
                (globalState).f5 = p0
                d_817_v155_: _dafny.Map
                d_817_v155_ = _dafny.Map({d_771_v124_: self.f12})
                d_818_v156_: D2
                d_818_v156_ = D2_DC8(d_770_cf0_, (0) - (p0), p1)
                d_817_v155_ = (d_817_v155_).set(default__.fm19((d_818_v156_).cf12, globalState), p0)
                index151_ = default__.safeIndex(537, (d_728_v99_).length(0))
                (d_728_v99_)[index151_] = d_770_cf0_
                d_819_v157_: _dafny.Array
                nw146_ = _dafny.Array(int(0), 16)
                d_819_v157_ = nw146_
                d_820_v158_: _dafny.Map
                d_820_v158_ = _dafny.Map({(d_728_v99_)[default__.safeIndex(537, (d_728_v99_).length(0))]: d_819_v157_})
                d_821_v159_: _dafny.Array
                nw147_ = _dafny.Array(None, 5)
                nw147_[int(0)] = d_819_v157_
                nw147_[int(1)] = d_819_v157_
                nw147_[int(2)] = d_819_v157_
                nw147_[int(3)] = ((d_820_v158_)[d_770_cf0_] if (d_770_cf0_) in (d_820_v158_) else d_819_v157_)
                nw147_[int(4)] = d_819_v157_
                d_821_v159_ = nw147_
                index152_ = default__.safeIndex(407, (d_821_v159_).length(0))
                (d_821_v159_)[index152_] = d_819_v157_
                index153_ = default__.safeIndex(407, (d_821_v159_).length(0))
                (d_821_v159_)[index153_] = d_819_v157_
            r1 = not(p1)
        elif True:
            d_822___mcc_h10_ = source9_.cf6
            d_823_cf6_ = d_822___mcc_h10_
            d_824_v160_: _dafny.Array
            nw148_ = _dafny.Array(int(0), 19)
            d_824_v160_ = nw148_
            d_825_v161_: _dafny.Set
            d_825_v161_ = _dafny.Set({d_824_v160_})
            d_826_v162_: _dafny.Map
            d_826_v162_ = _dafny.Map({p0: len(d_825_v161_)})
            d_827_v163_: _dafny.Map
            d_827_v163_ = _dafny.Map({p0: d_826_v162_})
            d_827_v163_ = (d_827_v163_).set(len(p2), d_826_v162_)
            if p1:
                d_828_v164_: _dafny.Set
                d_828_v164_ = _dafny.Set({p1})
                d_829_v165_: D6
                d_829_v165_ = D6_DC17(p1)
                d_830_v166_: _dafny.Seq
                d_830_v166_ = _dafny.SeqWithoutIsStrInference([(d_828_v164_).issubset(d_828_v164_), (p1 if p1 else p1), (188) < (p0), True, (d_829_v165_).cf24])
                r1 = (d_830_v166_)[default__.safeIndex(self.f12, len(d_830_v166_))]
                d_824_v160_ = d_824_v160_
                d_831_v167_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Map({}), 8)
                d_831_v167_ = nw149_
                d_832_v168_: _dafny.Map
                d_832_v168_ = _dafny.Map({p1: p1})
                index154_ = default__.safeIndex(274, (d_831_v167_).length(0))
                (d_831_v167_)[index154_] = d_832_v168_
                index155_ = default__.safeIndex(274, (d_831_v167_).length(0))
                (d_831_v167_)[index155_] = (d_832_v168_) | ((default__.fm18(self.f12, self.f12, p1, globalState)) | (d_832_v168_))
                r2 = ((0) - (p0)) * (self.f12)
                d_833_v169_: _dafny.MultiSet
                d_833_v169_ = _dafny.MultiSet([p1, True])
                d_834_v170_: int
                d_835_v171_: bool
                d_836_v172_: int
                out51_: int
                out52_: bool
                out53_: int
                out51_, out52_, out53_ = (self).m3((d_833_v169_) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p1]))), p1))), globalState)
                d_834_v170_ = out51_
                d_835_v171_ = out52_
                d_836_v172_ = out53_
            elif True:
                index156_ = default__.safeIndex(731, (d_824_v160_).length(0))
                (d_824_v160_)[index156_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "va")))
                index157_ = default__.safeIndex(731, (d_824_v160_).length(0))
                (d_824_v160_)[index157_] = len(p2)
                index158_ = default__.safeIndex(731, (d_824_v160_).length(0))
                (d_824_v160_)[index158_] = self.f12
                d_837_v173_: _dafny.Seq
                d_837_v173_ = _dafny.SeqWithoutIsStrInference([p0])
                d_838_v174_: _dafny.Map
                d_838_v174_ = _dafny.Map({_dafny.Map({((d_677_v61_)[self.f12] if (self.f12) in (d_677_v61_) else (d_837_v173_)[default__.safeIndex(len(p2), len(d_837_v173_))]): len(p2)}): (d_824_v160_)[default__.safeIndex(731, (d_824_v160_).length(0))]})
                d_838_v174_ = (d_838_v174_ if p1 else d_838_v174_)
                d_839_v175_: _dafny.Array
                nw150_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                d_839_v175_ = nw150_
                index159_ = default__.safeIndex(232, (d_839_v175_).length(0))
                (d_839_v175_)[index159_] = (p2) + (p2)
                index160_ = default__.safeIndex(232, (d_839_v175_).length(0))
                (d_839_v175_)[index160_] = (p2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kemlgdi")))
                d_840_v176_: _dafny.Set
                d_840_v176_ = _dafny.Set({p1})
                d_841_v177_: _dafny.Seq
                d_841_v177_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1, not(True)}), d_840_v176_, d_840_v176_])
                d_842_v178_: _dafny.Set
                d_842_v178_ = _dafny.Set({(d_841_v177_)[default__.safeIndex((d_824_v160_)[default__.safeIndex(731, (d_824_v160_).length(0))], len(d_841_v177_))], d_840_v176_, d_840_v176_})
                d_843_v179_: D6
                d_843_v179_ = D6_DC17(True)
                d_844_v180_: _dafny.Seq
                d_844_v180_ = _dafny.SeqWithoutIsStrInference([d_728_v99_, d_728_v99_, d_728_v99_])
                rhs137_ = d_842_v178_
                rhs138_ = d_843_v179_
                rhs139_ = ((d_844_v180_) + (d_844_v180_)).set(default__.safeIndex((d_837_v173_)[default__.safeIndex(self.f12, len(d_837_v173_))], len((d_844_v180_) + (d_844_v180_))), d_728_v99_)
                rhs140_ = default__.safeModuloInt((0) - ((0) - (self.f12)), self.f12)
                rhs141_ = -338
                lhs88_ = globalState
                lhs89_ = globalState
                d_842_v178_ = rhs137_
                d_843_v179_ = rhs138_
                d_844_v180_ = rhs139_
                lhs88_.f5 = rhs140_
                lhs89_.f5 = rhs141_
            d_845_v181_: _dafny.Array
            def lambda42_(d_846_p2_):
                def lambda43_(d_847_i13_):
                    return d_846_p2_

                return lambda43_

            init23_ = lambda42_(p2)
            nw151_ = _dafny.Array(None, 3)
            for i0_23_ in range(nw151_.length(0)):
                nw151_[i0_23_] = init23_(i0_23_)
            d_845_v181_ = nw151_
            d_845_v181_ = (d_845_v181_ if p1 else d_845_v181_)
            d_848_v182_: _dafny.Seq
            d_848_v182_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_849_v183_: _dafny.Set
            d_849_v183_ = _dafny.Set({p1})
            d_850_v184_: _dafny.Seq
            d_850_v184_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_848_v182_ = default__.fm9((_dafny.SeqWithoutIsStrInference([len((d_826_v162_).set(self.f12, 105)), p0, (0) - (len(d_849_v183_))])) <= (d_850_v184_), globalState)
        d_851_v185_: _dafny.MultiSet
        d_851_v185_ = _dafny.MultiSet([p1])
        d_852_v186_: _dafny.Map
        d_852_v186_ = _dafny.Map({not(p1): d_851_v185_})
        d_853_v187_: _dafny.Set
        d_853_v187_ = _dafny.Set({True, p1})
        r0 = not(default__.fm1(self.f12, (d_851_v185_).issubset(((d_852_v186_)[p1] if (p1) in (d_852_v186_) else d_851_v185_)), len(d_853_v187_), False, globalState))
        r0 = ((d_677_v61_).cardinality) > (p0)
        r1 = not ((p1) == (p1)) or (not (p1) or (p1))
        r2 = default__.safeDivisionInt(p0, self.f12)
        return r0, r1, r2

    @property
    def f14(self):
        return self._f14

class C4:
    def  __init__(self):
        self.f11: int = int(0)
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f10, f11):
        (self)._f10 = f10
        (self).f11 = f11

    def fm2(self, p0, globalState):
        return (self).f10

    def m1(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Set = _dafny.Set({})
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_854_v1_: _dafny.Set
        d_854_v1_ = _dafny.Set({999, (self).f10, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stclb"))), (0) - ((self).f10), self.f11})
        d_855_v2_: D0
        def iife48_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in (d_854_v1_).Elements:
                d_856_v0_: int = compr_22_
                if (d_856_v0_) in (d_854_v1_):
                    coll22_[(d_856_v0_) * ((self).f10)] = True
            return _dafny.Map(coll22_)
        d_855_v2_ = D0_DC2(p3, iife48_()
, (_dafny.Map({p2: p1})).set(p3, p1), p2)
        source12_ = d_855_v2_
        if source12_.is_DC1:
            d_857___mcc_h0_ = source12_.cf1
            d_858_cf1_ = d_857___mcc_h0_
            r1 = d_858_cf1_
            if (p3 if p3 else p2):
                d_859_v3_: _dafny.MultiSet
                d_859_v3_ = _dafny.MultiSet([self.f11])
                d_860_v4_: D2
                d_860_v4_ = D2_DC8(p3, (d_859_v3_).cardinality, p2)
                d_861_v5_: _dafny.Map
                d_861_v5_ = _dafny.Map({(self).f10: p3})
                d_862_v6_: _dafny.Map
                d_862_v6_ = _dafny.Map({p2: self.f11})
                d_863_v7_: _dafny.Seq
                d_863_v7_ = _dafny.SeqWithoutIsStrInference([self.f11])
                d_864_v8_: _dafny.Map
                d_864_v8_ = _dafny.Map({self.f11: p1})
                d_865_v9_: _dafny.Array
                nw152_ = _dafny.Array(None, 27)
                nw152_[int(0)] = self.f11
                nw152_[int(1)] = default__.fm0((self).f10, (self).f10, globalState)
                nw152_[int(2)] = (self).f10
                nw152_[int(3)] = default__.safeDivisionInt(len(p0), len(p0))
                nw152_[int(4)] = 528
                nw152_[int(5)] = (self).f10
                nw152_[int(6)] = (0) - ((d_860_v4_).cf12)
                nw152_[int(7)] = (self).f10
                nw152_[int(8)] = len((d_861_v5_ if p2 else default__.fm3(p2, p3, len(d_862_v6_), globalState)))
                nw152_[int(9)] = self.f11
                nw152_[int(10)] = (self).f10
                nw152_[int(11)] = 171
                nw152_[int(12)] = (self).f10
                nw152_[int(13)] = -520
                nw152_[int(14)] = (self).f10
                nw152_[int(15)] = (d_863_v7_)[default__.safeIndex(658, len(d_863_v7_))]
                nw152_[int(16)] = len((p0) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tup"))))
                nw152_[int(17)] = len(d_864_v8_)
                nw152_[int(18)] = default__.safeDivisionInt(106, self.f11)
                nw152_[int(19)] = ((self).f10) - (self.f11)
                nw152_[int(20)] = self.f11
                nw152_[int(21)] = default__.safeDivisionInt(self.f11, self.f11)
                nw152_[int(22)] = (self).f10
                nw152_[int(23)] = (self.f11 if p3 else self.f11)
                nw152_[int(24)] = (self).f10
                nw152_[int(25)] = self.f11
                nw152_[int(26)] = self.f11
                d_865_v9_ = nw152_
                d_866_v10_: D2
                d_866_v10_ = D2_DC9(779)
                index161_ = default__.safeIndex(128, (d_865_v9_).length(0))
                (d_865_v9_)[index161_] = (d_866_v10_).cf14
                index162_ = default__.safeIndex(128, (d_865_v9_).length(0))
                (d_865_v9_)[index162_] = self.f11
                d_867_v11_: bool
                d_867_v11_ = True
                d_867_v11_ = False
                d_868_v12_: _dafny.Set
                d_868_v12_ = _dafny.Set({d_859_v3_, d_859_v3_, _dafny.MultiSet([(self).f10])})
                rhs142_ = d_868_v12_
                rhs143_ = False
                rhs144_ = default__.safeModuloInt((self).f10, self.f11)
                rhs145_ = not (p2) or (d_867_v11_)
                d_868_v12_ = rhs142_
                d_867_v11_ = rhs143_
                r0 = rhs144_
                d_867_v11_ = rhs145_
                d_869_v13_: _dafny.Set
                d_869_v13_ = _dafny.Set({p2})
                d_870_v14_: C2
                nw153_ = C2()
                nw153_.ctor__((d_869_v13_).issubset(d_869_v13_), (d_865_v9_)[default__.safeIndex(128, (d_865_v9_).length(0))])
                d_870_v14_ = nw153_
                (globalState).f5 = (d_865_v9_)[default__.safeIndex(128, (d_865_v9_).length(0))]
            elif True:
                d_871_v15_: bool
                d_871_v15_ = False
                d_872_v16_: _dafny.MultiSet
                d_872_v16_ = _dafny.MultiSet([d_858_cf1_, d_858_cf1_])
                d_871_v15_ = not ((d_872_v16_).isdisjoint(_dafny.MultiSet([d_858_cf1_, d_858_cf1_, d_858_cf1_]))) or (p2)
                d_873_v17_: _dafny.Array
                nw154_ = _dafny.Array(int(0), 3)
                d_873_v17_ = nw154_
                index163_ = default__.safeIndex(625, (d_873_v17_).length(0))
                (d_873_v17_)[index163_] = self.f11
                index164_ = default__.safeIndex(625, (d_873_v17_).length(0))
                (d_873_v17_)[index164_] = self.f11
                d_871_v15_ = ((self.f11 if d_871_v15_ else (self).f10)) < (len(p0))
                (globalState).f5 = default__.safeModuloInt((d_873_v17_)[default__.safeIndex(625, (d_873_v17_).length(0))], self.f11)
                d_871_v15_ = (p2) or ((p0) != (p0))
            d_874_v18_: D5
            d_874_v18_ = D5_DC13(True, False)
            d_875_v19_: _dafny.MultiSet
            d_875_v19_ = _dafny.MultiSet([p3, (d_874_v18_).cf19])
            d_876_v20_: _dafny.Seq
            d_876_v20_ = _dafny.SeqWithoutIsStrInference([681, self.f11])
            d_877_v21_: _dafny.Seq
            d_877_v21_ = _dafny.SeqWithoutIsStrInference([d_875_v19_, d_875_v19_, (d_875_v19_).set(p2, default__.abs(len(d_876_v20_))), d_875_v19_, d_875_v19_])
            source13_ = D5_DC12(d_877_v21_)
            if source13_.is_DC13:
                d_878___mcc_h7_ = source13_.cf18
                d_879___mcc_h8_ = source13_.cf19
                d_880_cf19_ = d_879___mcc_h8_
                d_881_cf18_ = d_878___mcc_h7_
                d_882_v22_: C1
                nw155_ = C1()
                nw155_.ctor__(p1, p1, (self).f10)
                d_882_v22_ = nw155_
                rhs146_ = d_882_v22_
                rhs147_ = d_880_cf19_
                d_882_v22_ = rhs146_
                d_881_cf18_ = rhs147_
                (globalState).f5 = 115
                d_883_v23_: _dafny.MultiSet
                d_883_v23_ = _dafny.MultiSet([self.f11, len(p0)])
                d_884_v24_: D5
                d_884_v24_ = D5_DC12(_dafny.SeqWithoutIsStrInference([d_875_v19_, d_875_v19_]))
                (self).f11 = len(default__.fm20((self).fm2(d_881_cf18_, globalState), self.f11, (d_883_v23_).cardinality, d_884_v24_, globalState))
                (globalState).f5 = (d_875_v19_).cardinality
            elif source13_.is_DC14:
                d_885___mcc_h9_ = source13_.cf20
                d_886_cf20_ = d_885___mcc_h9_
                d_887_v25_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                d_887_v25_ = nw156_
                index165_ = default__.safeIndex(37, (d_887_v25_).length(0))
                (d_887_v25_)[index165_] = p0
                index166_ = default__.safeIndex(37, (d_887_v25_).length(0))
                (d_887_v25_)[index166_] = p0
                d_888_v26_: C1
                nw157_ = C1()
                nw157_.ctor__(p1, p1, len((d_887_v25_)[default__.safeIndex(37, (d_887_v25_).length(0))]))
                d_888_v26_ = nw157_
                d_888_v26_ = (d_888_v26_ if False else d_888_v26_)
                d_889_v27_: _dafny.Array
                nw158_ = _dafny.Array(int(0), 7)
                d_889_v27_ = nw158_
                r3 = d_889_v27_
                d_890_v28_: _dafny.MultiSet
                d_890_v28_ = _dafny.MultiSet([p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ed"))])
                d_891_v29_: _dafny.Map
                d_891_v29_ = _dafny.Map({d_858_cf1_: (p0) not in (d_890_v28_)})
                d_891_v29_ = (d_891_v29_).set(d_858_cf1_, default__.fm1((self).f10, not(d_886_cf20_), (self).f10, p3, globalState))
            elif True:
                d_892___mcc_h10_ = source13_.cf17
                d_893_cf17_ = d_892___mcc_h10_
                d_894_v30_: C1
                nw159_ = C1()
                nw159_.ctor__(p1, p1, self.f11)
                d_894_v30_ = nw159_
                r0 = self.f11
                d_895_v31_: bool
                d_895_v31_ = True
                d_895_v31_ = p2
                d_896_v32_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_896_v32_ = nw160_
                index167_ = default__.safeIndex(252, (d_896_v32_).length(0))
                (d_896_v32_)[index167_] = p0
                index168_ = default__.safeIndex(252, (d_896_v32_).length(0))
                (d_896_v32_)[index168_] = (p0) + (p0)
            d_897_v33_: _dafny.Array
            nw161_ = _dafny.Array(int(0), 1)
            d_897_v33_ = nw161_
            index169_ = default__.safeIndex(558, (d_897_v33_).length(0))
            (d_897_v33_)[index169_] = (self).f10
            index170_ = default__.safeIndex(558, (d_897_v33_).length(0))
            (d_897_v33_)[index170_] = self.f11
        elif source12_.is_DC2:
            d_898___mcc_h1_ = source12_.cf2
            d_899___mcc_h2_ = source12_.cf3
            d_900___mcc_h3_ = source12_.cf4
            d_901___mcc_h4_ = source12_.cf5
            d_902_cf5_ = d_901___mcc_h4_
            d_903_cf4_ = d_900___mcc_h3_
            d_904_cf3_ = d_899___mcc_h2_
            d_905_cf2_ = d_898___mcc_h1_
            d_906_v34_: D7
            d_906_v34_ = D7_DC21()
            d_906_v34_ = D7_DC21()
            if False:
                d_902_cf5_ = d_902_cf5_
                d_907_v35_: _dafny.Seq
                d_907_v35_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11, len(_dafny.Map({p2: (D2_DC7((self).f10)).cf10})), 533])
                d_908_v36_: C2
                nw162_ = C2()
                nw162_.ctor__(p3, (0) - (default__.safeModuloInt(len(d_907_v35_), (0) - (self.f11))))
                d_908_v36_ = nw162_
                d_909_v37_: _dafny.Map
                d_909_v37_ = _dafny.Map({p1: (d_908_v36_).f17})
                d_910_v38_: _dafny.MultiSet
                d_910_v38_ = _dafny.MultiSet([((d_909_v37_)[_dafny.CodePoint('a')] if (_dafny.CodePoint('a')) in (d_909_v37_) else not(p2)), d_902_cf5_])
                (globalState).f5 = ((d_910_v38_)[((self).f10) < (self.f11)] if (((self).f10) < (self.f11)) in (d_910_v38_) else len(d_904_cf3_))
                d_911_v39_: _dafny.Array
                nw163_ = _dafny.Array(None, 28)
                nw163_[int(0)] = d_902_cf5_
                nw163_[int(1)] = p3
                nw163_[int(2)] = d_905_cf2_
                nw163_[int(3)] = (d_908_v36_).f17
                nw163_[int(4)] = p3
                nw163_[int(5)] = d_902_cf5_
                nw163_[int(6)] = (d_908_v36_).f17
                nw163_[int(7)] = d_902_cf5_
                nw163_[int(8)] = p3
                nw163_[int(9)] = (d_908_v36_).f17
                nw163_[int(10)] = (d_908_v36_).f17
                nw163_[int(11)] = d_905_cf2_
                nw163_[int(12)] = d_902_cf5_
                nw163_[int(13)] = p2
                nw163_[int(14)] = True
                nw163_[int(15)] = p2
                nw163_[int(16)] = d_905_cf2_
                nw163_[int(17)] = d_902_cf5_
                nw163_[int(18)] = d_902_cf5_
                nw163_[int(19)] = p3
                nw163_[int(20)] = True
                nw163_[int(21)] = p2
                nw163_[int(22)] = p3
                nw163_[int(23)] = False
                nw163_[int(24)] = p3
                nw163_[int(25)] = p2
                nw163_[int(26)] = p2
                nw163_[int(27)] = p2
                d_911_v39_ = nw163_
                d_912_v40_: T1
                nw164_ = C1()
                nw164_.ctor__(p1, _dafny.CodePoint('l'), (D6_DC16(d_911_v39_, (self).f10)).cf23)
                d_912_v40_ = nw164_
                d_913_v41_: _dafny.Array
                def lambda44_(d_914_p0_):
                    def lambda45_(d_915_i0_):
                        return d_914_p0_

                    return lambda45_

                init24_ = lambda44_(p0)
                nw165_ = _dafny.Array(None, 6)
                for i0_24_ in range(nw165_.length(0)):
                    nw165_[i0_24_] = init24_(i0_24_)
                d_913_v41_ = nw165_
                d_916_v42_: _dafny.Seq
                d_916_v42_ = _dafny.SeqWithoutIsStrInference([d_912_v40_, d_912_v40_])
                rhs148_ = (d_916_v42_)[default__.safeIndex(-414, len(d_916_v42_))]
                rhs149_ = ((_dafny.SeqWithoutIsStrInference([(d_912_v40_).f13 for d_917_i1_ in range(default__.abs(644))])) + (p0)) not in ((_dafny.SeqWithoutIsStrInference([p0 for d_918_i2_ in range(default__.abs(302))])) + (_dafny.SeqWithoutIsStrInference([p0, p0])))
                rhs150_ = d_913_v41_
                rhs151_ = (0) - ((self).f10)
                lhs90_ = d_912_v40_
                d_912_v40_ = rhs148_
                d_902_cf5_ = rhs149_
                d_913_v41_ = rhs150_
                lhs90_.f12 = rhs151_
                d_919_v43_: _dafny.MultiSet
                d_919_v43_ = _dafny.MultiSet([default__.fm11(d_912_v40_.f12, d_912_v40_.f12, p2, globalState)])
                d_920_v44_: _dafny.Array
                nw166_ = _dafny.Array(D2.default()(), 23)
                d_920_v44_ = nw166_
                d_921_v45_: _dafny.Set
                d_921_v45_ = _dafny.Set({(d_912_v40_).f13})
                d_922_v46_: _dafny.Map
                d_922_v46_ = _dafny.Map({d_905_cf2_: default__.fm0(len(d_921_v45_), self.f11, globalState)})
                d_923_v47_: C3
                nw167_ = C3()
                nw167_.ctor__(d_920_v44_, p1, ((d_922_v46_)[p2] if (p2) in (d_922_v46_) else self.f11))
                d_923_v47_ = nw167_
                d_924_v48_: _dafny.Map
                d_924_v48_ = _dafny.Map({(D8_DC22(d_919_v43_)).cf28: d_923_v47_})
                d_924_v48_ = (d_924_v48_).set(_dafny.MultiSet([p0, p0, p0, p0, p0]), d_923_v47_)
            elif True:
                d_925_v49_: _dafny.Array
                nw168_ = _dafny.Array(_dafny.MultiSet({}), 8)
                d_925_v49_ = nw168_
                d_926_v50_: _dafny.Set
                d_926_v50_ = _dafny.Set({p2, d_902_cf5_})
                d_927_v51_: _dafny.MultiSet
                d_927_v51_ = _dafny.MultiSet([len(d_926_v50_)])
                index171_ = default__.safeIndex(693, (d_925_v49_).length(0))
                (d_925_v49_)[index171_] = d_927_v51_
                index172_ = default__.safeIndex(693, (d_925_v49_).length(0))
                (d_925_v49_)[index172_] = d_927_v51_
                d_928_v52_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 10)
                d_928_v52_ = nw169_
                d_929_v53_: _dafny.Map
                d_929_v53_ = _dafny.Map({p3: self.f11})
                index173_ = default__.safeIndex(478, (d_928_v52_).length(0))
                (d_928_v52_)[index173_] = len((d_929_v53_).set(p2, (0) - (self.f11)))
                d_930_v54_: _dafny.Map
                d_930_v54_ = _dafny.Map({len(p0): (self).f10})
                index174_ = default__.safeIndex(478, (d_928_v52_).length(0))
                (d_928_v52_)[index174_] = ((self.f11) - (len(d_930_v54_))) * (self.f11)
                d_905_cf2_ = d_905_cf2_
                d_905_cf2_ = default__.fm1(self.f11, p3, (self).f10, d_902_cf5_, globalState)
                d_931_v55_: _dafny.Set
                d_931_v55_ = _dafny.Set({d_928_v52_, d_928_v52_, d_928_v52_})
                d_932_v56_: D5
                d_932_v56_ = D5_DC13(True, ((0) - (self.f11)) != (len(d_931_v55_)))
                d_932_v56_ = d_932_v56_
            (globalState).f5 = len(default__.fm10(default__.safeModuloInt(-926, self.f11), globalState))
            d_933_v57_: _dafny.Seq
            d_933_v57_ = _dafny.SeqWithoutIsStrInference([d_905_cf2_])
            d_934_v58_: _dafny.Map
            d_934_v58_ = _dafny.Map({d_933_v57_: _dafny.Set({p3})})
            d_935_v59_: _dafny.Set
            d_935_v59_ = _dafny.Set({p2, p3})
            d_902_cf5_ = (d_935_v59_).ispropersubset((((d_934_v58_)[d_933_v57_] if (d_933_v57_) in (d_934_v58_) else d_935_v59_)) - (_dafny.Set({p3, p3})))
        elif source12_.is_DC0:
            d_936___mcc_h5_ = source12_.cf0
            d_937_cf0_ = d_936___mcc_h5_
            d_938_v60_: _dafny.MultiSet
            d_938_v60_ = _dafny.MultiSet([d_854_v1_])
            d_937_cf0_ = (d_938_v60_).ispropersubset(d_938_v60_)
            d_939_v61_: _dafny.Array
            nw170_ = _dafny.Array(None, 8)
            nw170_[int(0)] = self.f11
            nw170_[int(1)] = self.f11
            nw170_[int(2)] = (0) - (self.f11)
            nw170_[int(3)] = len(d_854_v1_)
            nw170_[int(4)] = (self).f10
            nw170_[int(5)] = self.f11
            nw170_[int(6)] = (self).f10
            nw170_[int(7)] = len(_dafny.Map({self.f11: self.f11}))
            d_939_v61_ = nw170_
            d_940_v62_: _dafny.Array
            nw171_ = _dafny.Array(None, 16)
            nw171_[int(0)] = d_937_cf0_
            nw171_[int(1)] = False
            nw171_[int(2)] = p3
            nw171_[int(3)] = p3
            nw171_[int(4)] = p2
            nw171_[int(5)] = p2
            nw171_[int(6)] = p2
            nw171_[int(7)] = p3
            nw171_[int(8)] = p2
            nw171_[int(9)] = default__.fm1((self).f10, p3, (self).f10, p2, globalState)
            nw171_[int(10)] = p3
            nw171_[int(11)] = p2
            nw171_[int(12)] = d_937_cf0_
            nw171_[int(13)] = d_937_cf0_
            nw171_[int(14)] = p3
            nw171_[int(15)] = d_937_cf0_
            d_940_v62_ = nw171_
            d_941_v63_: _dafny.Map
            d_941_v63_ = _dafny.Map({d_939_v61_: d_940_v62_})
            index175_ = default__.safeIndex(731, (d_940_v62_).length(0))
            (d_940_v62_)[index175_] = p2
            index176_ = default__.safeIndex(731, (d_940_v62_).length(0))
            rhs152_ = ((d_941_v63_).set(d_939_v61_, d_940_v62_)) | (d_941_v63_)
            rhs153_ = p2
            lhs91_ = d_940_v62_
            lhs92_ = default__.safeIndex(731, (d_940_v62_).length(0))
            d_941_v63_ = rhs152_
            lhs91_[lhs92_] = rhs153_
            d_942_v64_: _dafny.MultiSet
            d_942_v64_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hs")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yn"))])
            d_943_v65_: D8
            d_943_v65_ = D8_DC22(d_942_v64_)
            source14_ = d_943_v65_
            if source14_.is_DC23:
                d_944___mcc_h11_ = source14_.cf29
                d_945_cf29_ = d_944___mcc_h11_
                d_946_v66_: _dafny.Map
                d_946_v66_ = _dafny.Map({p2: d_939_v61_})
                d_947_v67_: T1
                nw172_ = C1()
                nw172_.ctor__(p1, p1, d_945_cf29_)
                d_947_v67_ = nw172_
                pat_let_tv6_ = d_946_v66_
                d_948_v68_: _dafny.Map
                def iife49_(_pat_let13_0):
                    def iife50_(d_949_dt__update__tmp_h0_):
                        def iife51_(_pat_let14_0):
                            def iife52_(d_950_dt__update_hcf27_h0_):
                                return D7_DC20(d_950_dt__update_hcf27_h0_)
                            return iife52_(_pat_let14_0)
                        return iife51_(pat_let_tv6_)
                    return iife50_(_pat_let13_0)
                d_948_v68_ = _dafny.Map({iife49_(D7_DC20(_dafny.Map({(d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))]: d_939_v61_}))): d_947_v67_})
                d_951_v69_: _dafny.Map
                d_951_v69_ = _dafny.Map({d_948_v68_: not(p3)})
                d_952_v70_: _dafny.Map
                d_952_v70_ = _dafny.Map({d_945_cf29_: d_951_v69_})
                d_953_v71_: D2
                d_953_v71_ = D2_DC7(d_947_v67_.f12)
                d_954_v72_: _dafny.Map
                d_954_v72_ = _dafny.Map({(d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))]: (d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))]})
                pat_let_tv7_ = d_945_cf29_
                d_955_v73_: _dafny.Array
                nw173_ = _dafny.Array(None, 15)
                nw173_[int(0)] = d_953_v71_
                def iife53_(_pat_let15_0):
                    def iife54_(d_956_dt__update__tmp_h1_):
                        def iife55_(_pat_let16_0):
                            def iife56_(d_957_dt__update_hcf10_h0_):
                                return D2_DC7(d_957_dt__update_hcf10_h0_)
                            return iife56_(_pat_let16_0)
                        return iife55_(pat_let_tv7_)
                    return iife54_(_pat_let15_0)
                nw173_[int(1)] = iife53_(d_953_v71_)
                nw173_[int(2)] = d_953_v71_
                nw173_[int(3)] = D2_DC7(len(d_954_v72_))
                nw173_[int(4)] = d_953_v71_
                nw173_[int(5)] = d_953_v71_
                nw173_[int(6)] = D2_DC7((self).f10)
                nw173_[int(7)] = d_953_v71_
                nw173_[int(8)] = d_953_v71_
                nw173_[int(9)] = d_953_v71_
                nw173_[int(10)] = d_953_v71_
                nw173_[int(11)] = d_953_v71_
                nw173_[int(12)] = D2_DC7(539)
                nw173_[int(13)] = d_953_v71_
                nw173_[int(14)] = d_953_v71_
                d_955_v73_ = nw173_
                d_958_v74_: C3
                nw174_ = C3()
                nw174_.ctor__(d_955_v73_, p1, -596)
                d_958_v74_ = nw174_
                d_959_v75_: _dafny.Seq
                d_959_v75_ = _dafny.SeqWithoutIsStrInference([d_958_v74_])
                d_952_v70_ = (d_952_v70_).set(default__.safeDivisionInt(d_945_cf29_, len(d_959_v75_)), d_951_v69_)
                d_960_v76_: D7
                d_960_v76_ = D7_DC20(d_946_v66_)
                d_961_v77_: _dafny.Seq
                d_961_v77_ = _dafny.SeqWithoutIsStrInference([d_960_v76_])
                d_961_v77_ = (((d_961_v77_) + (d_961_v77_)) + (d_961_v77_)).set(default__.safeIndex(d_947_v67_.f12, len(((d_961_v77_) + (d_961_v77_)) + (d_961_v77_))), d_960_v76_)
                d_962_v78_: _dafny.Array
                def lambda46_(d_963_p3_):
                    def lambda47_(d_964_i3_):
                        return _dafny.Set({d_963_p3_})

                    return lambda47_

                init25_ = lambda46_(p3)
                nw175_ = _dafny.Array(None, 29)
                for i0_25_ in range(nw175_.length(0)):
                    nw175_[i0_25_] = init25_(i0_25_)
                d_962_v78_ = nw175_
                d_965_v79_: C3
                nw176_ = C3()
                nw176_.ctor__((d_958_v74_).f14, p1, (_dafny.MultiSet([d_962_v78_, d_962_v78_])).cardinality)
                d_965_v79_ = nw176_
                d_966_v80_: str
                d_966_v80_ = _dafny.CodePoint('u')
                d_967_v81_: _dafny.Map
                d_967_v81_ = _dafny.Map({p2: (d_947_v67_).f13})
                d_966_v80_ = ((d_967_v81_)[(p2 if (d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))] else p2)] if ((p2 if (d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))] else p2)) in (d_967_v81_) else p1)
            elif source14_.is_DC24:
                d_968___mcc_h12_ = source14_.cf30
                d_969___mcc_h13_ = source14_.cf31
                d_970_cf31_ = d_969___mcc_h13_
                d_971_cf30_ = d_968___mcc_h12_
                d_972_v82_: D2
                d_972_v82_ = D2_DC7(self.f11)
                d_973_v83_: _dafny.Array
                nw177_ = _dafny.Array(None, 13)
                nw177_[int(0)] = d_972_v82_
                nw177_[int(1)] = d_972_v82_
                nw177_[int(2)] = d_972_v82_
                nw177_[int(3)] = d_972_v82_
                nw177_[int(4)] = d_972_v82_
                nw177_[int(5)] = d_972_v82_
                nw177_[int(6)] = d_972_v82_
                nw177_[int(7)] = d_972_v82_
                nw177_[int(8)] = d_972_v82_
                nw177_[int(9)] = d_972_v82_
                nw177_[int(10)] = d_972_v82_
                nw177_[int(11)] = d_972_v82_
                nw177_[int(12)] = d_972_v82_
                d_973_v83_ = nw177_
                d_974_v84_: D9
                d_974_v84_ = D9_DC25(p1)
                d_975_v85_: C3
                nw178_ = C3()
                nw178_.ctor__(d_973_v83_, (d_974_v84_).cf32, len(_dafny.Set({d_937_cf0_})))
                d_975_v85_ = nw178_
                index177_ = default__.safeIndex(731, (d_940_v62_).length(0))
                (d_940_v62_)[index177_] = True
                d_976_v86_: _dafny.Map
                d_976_v86_ = _dafny.Map({True: not(not(p3))})
                d_977_v87_: _dafny.Seq
                d_977_v87_ = _dafny.SeqWithoutIsStrInference([(len(d_976_v86_) if not((d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))]) else self.f11)])
                (globalState).f5 = (d_977_v87_)[default__.safeIndex((0) - ((0) - ((self).f10)), len(d_977_v87_))]
                d_978_v88_: D6
                d_978_v88_ = D6_DC18(len(_dafny.Map({not(d_937_cf0_): p2})))
                index178_ = default__.safeIndex(731, (d_940_v62_).length(0))
                index179_ = default__.safeIndex(731, (d_940_v62_).length(0))
                index180_ = default__.safeIndex(731, (d_940_v62_).length(0))
                rhs154_ = (d_940_v62_)[default__.safeIndex(731, (d_940_v62_).length(0))]
                rhs155_ = (d_978_v88_).cf25
                rhs156_ = default__.fm1((self).f10, d_937_cf0_, (0) - (default__.safeModuloInt(self.f11, 250)), p2, globalState)
                rhs157_ = p2
                lhs93_ = d_940_v62_
                lhs94_ = default__.safeIndex(731, (d_940_v62_).length(0))
                lhs95_ = self
                lhs96_ = d_940_v62_
                lhs97_ = default__.safeIndex(731, (d_940_v62_).length(0))
                lhs98_ = d_940_v62_
                lhs99_ = default__.safeIndex(731, (d_940_v62_).length(0))
                lhs93_[lhs94_] = rhs154_
                lhs95_.f11 = rhs155_
                lhs96_[lhs97_] = rhs156_
                lhs98_[lhs99_] = rhs157_
            elif True:
                d_979___mcc_h14_ = source14_.cf28
                d_980_cf28_ = d_979___mcc_h14_
                d_981_v89_: _dafny.Map
                d_981_v89_ = _dafny.Map({(self).f10: self.f11})
                d_982_v90_: _dafny.Seq
                d_982_v90_ = _dafny.SeqWithoutIsStrInference([858])
                d_981_v89_ = (d_981_v89_).set(self.f11, (d_982_v90_)[default__.safeIndex(len(p0), len(d_982_v90_))])
                d_983_v91_: T1
                nw179_ = C1()
                nw179_.ctor__((p0)[default__.safeIndex(self.f11, len(p0))], p1, (self).f10)
                d_983_v91_ = nw179_
                d_984_v92_: _dafny.Array
                def lambda48_(d_985_v91_, d_986_p0_):
                    def lambda49_(d_987_i4_):
                        return D2_DC7(len(_dafny.Set({d_985_v91_.f12, (self).f10, d_985_v91_.f12, len(d_986_p0_)})))

                    return lambda49_

                init26_ = lambda48_(d_983_v91_, p0)
                nw180_ = _dafny.Array(None, 12)
                for i0_26_ in range(nw180_.length(0)):
                    nw180_[i0_26_] = init26_(i0_26_)
                d_984_v92_ = nw180_
                nw181_ = C3()
                nw181_.ctor__(d_984_v92_, (d_983_v91_).f13, (d_982_v90_)[default__.safeIndex((self).f10, len(d_982_v90_))])
                d_983_v91_ = nw181_
                d_988_v93_: str
                d_988_v93_ = _dafny.CodePoint('n')
                d_988_v93_ = _dafny.CodePoint('a')
                d_854_v1_ = d_854_v1_
            d_940_v62_ = d_940_v62_
        elif True:
            d_989___mcc_h6_ = source12_.cf6
            d_990_cf6_ = d_989___mcc_h6_
            d_991_v94_: _dafny.Array
            nw182_ = _dafny.Array(False, 21)
            d_991_v94_ = nw182_
            d_992_v95_: D0
            d_992_v95_ = D0_DC1(d_991_v94_)
            d_993_v97_: _dafny.Seq
            def iife57_():
                coll23_ = _dafny.Map()
                compr_23_: str
                for compr_23_ in (p0).Elements:
                    d_994_v96_: str = compr_23_
                    if (d_994_v96_) in (p0):
                        coll23_[d_994_v96_] = p2
                return _dafny.Map(coll23_)
            d_993_v97_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife57_()
            )])), 257])
            d_995_v98_: _dafny.Map
            d_995_v98_ = _dafny.Map({(d_992_v95_).cf1: len((d_993_v97_) + (_dafny.SeqWithoutIsStrInference([self.f11])))})
            index181_ = default__.safeIndex(125, (d_991_v94_).length(0))
            (d_991_v94_)[index181_] = p3
            d_996_v99_: _dafny.MultiSet
            d_996_v99_ = _dafny.MultiSet([p0, (p0) + (_dafny.SeqWithoutIsStrInference([p1 for d_997_i5_ in range(default__.abs(447))]))])
            d_998_v100_: _dafny.MultiSet
            d_998_v100_ = _dafny.MultiSet([self.f11])
            d_999_v101_: _dafny.Set
            d_999_v101_ = _dafny.Set({_dafny.MultiSet(d_993_v97_), d_998_v100_, d_998_v100_})
            index182_ = default__.safeIndex(125, (d_991_v94_).length(0))
            rhs158_ = d_995_v98_
            rhs159_ = self.f11
            rhs160_ = (d_999_v101_) != ((d_999_v101_).intersection(d_999_v101_))
            rhs161_ = d_996_v99_
            lhs100_ = globalState
            lhs101_ = d_991_v94_
            lhs102_ = default__.safeIndex(125, (d_991_v94_).length(0))
            d_995_v98_ = rhs158_
            lhs100_.f5 = rhs159_
            lhs101_[lhs102_] = rhs160_
            d_996_v99_ = rhs161_
            d_1000_v102_: str
            d_1000_v102_ = _dafny.CodePoint('u')
            d_1000_v102_ = d_1000_v102_
            d_1001_v103_: _dafny.Map
            d_1001_v103_ = _dafny.Map({p0: p1})
            d_1002_v104_: _dafny.Map
            d_1002_v104_ = _dafny.Map({d_1001_v103_: self.f11})
            d_1002_v104_ = (d_1002_v104_).set(d_1001_v103_, (self).f10)
            index183_ = default__.safeIndex(125, (d_991_v94_).length(0))
            (d_991_v94_)[index183_] = p2
        d_1003_v105_: C1
        nw183_ = C1()
        nw183_.ctor__(p1, p1, 554)
        d_1003_v105_ = nw183_
        d_1004_v106_: _dafny.MultiSet
        d_1004_v106_ = _dafny.MultiSet([-902])
        d_1005_v107_: _dafny.Map
        d_1005_v107_ = _dafny.Map({146: d_1004_v106_})
        if ((d_1004_v106_).set(self.f11, default__.abs((self).f10))).issubset((d_1004_v106_ if p3 else ((d_1005_v107_)[self.f11] if (self.f11) in (d_1005_v107_) else d_1004_v106_))):
            d_1006_v108_: _dafny.Map
            d_1006_v108_ = _dafny.Map({p2: (self).f10})
            d_1007_v109_: _dafny.Map
            d_1007_v109_ = _dafny.Map({p2: len(d_1006_v108_)})
            d_1008_v110_: T0
            nw184_ = C2()
            nw184_.ctor__(p3, ((d_1007_v109_)[p2] if (p2) in (d_1007_v109_) else self.f11))
            d_1008_v110_ = nw184_
            d_1009_v111_: _dafny.Array
            def lambda50_(d_1010_v110_):
                def lambda51_(d_1011_i6_):
                    return (d_1011_i6_) - (d_1010_v110_.f12)

                return lambda51_

            init27_ = lambda50_(d_1008_v110_)
            nw185_ = _dafny.Array(None, 9)
            for i0_27_ in range(nw185_.length(0)):
                nw185_[i0_27_] = init27_(i0_27_)
            d_1009_v111_ = nw185_
            d_1012_v112_: _dafny.Map
            d_1012_v112_ = _dafny.Map({d_1008_v110_.f12: d_1009_v111_})
            d_1013_v113_: _dafny.Seq
            d_1013_v113_ = _dafny.SeqWithoutIsStrInference([d_1009_v111_])
            rhs162_ = (((d_1004_v106_).intersection(d_1004_v106_)) - ((d_1004_v106_).set((self).f10, default__.abs((self).f10)))).cardinality
            rhs163_ = len(((_dafny.Map({(self).f10: d_1009_v111_})) | ((d_1012_v112_).set(self.f11, d_1009_v111_))).set(self.f11, (d_1013_v113_)[default__.safeIndex(self.f11, len(d_1013_v113_))]))
            rhs164_ = d_1008_v110_
            rhs165_ = self.f11
            lhs103_ = self
            lhs104_ = globalState
            lhs105_ = d_1008_v110_
            lhs103_.f11 = rhs162_
            lhs104_.f5 = rhs163_
            d_1008_v110_ = rhs164_
            lhs105_.f12 = rhs165_
            d_1014_v114_: _dafny.Array
            nw186_ = _dafny.Array(None, 25)
            nw186_[int(0)] = d_1009_v111_
            nw186_[int(1)] = d_1009_v111_
            nw186_[int(2)] = d_1009_v111_
            nw186_[int(3)] = d_1009_v111_
            nw186_[int(4)] = d_1009_v111_
            nw186_[int(5)] = d_1009_v111_
            nw186_[int(6)] = d_1009_v111_
            nw186_[int(7)] = d_1009_v111_
            nw186_[int(8)] = d_1009_v111_
            nw186_[int(9)] = d_1009_v111_
            nw186_[int(10)] = d_1009_v111_
            nw186_[int(11)] = d_1009_v111_
            nw186_[int(12)] = d_1009_v111_
            nw186_[int(13)] = d_1009_v111_
            nw186_[int(14)] = d_1009_v111_
            nw186_[int(15)] = d_1009_v111_
            nw186_[int(16)] = d_1009_v111_
            nw186_[int(17)] = d_1009_v111_
            nw186_[int(18)] = d_1009_v111_
            nw186_[int(19)] = d_1009_v111_
            nw186_[int(20)] = d_1009_v111_
            nw186_[int(21)] = d_1009_v111_
            nw186_[int(22)] = d_1009_v111_
            nw186_[int(23)] = d_1009_v111_
            nw186_[int(24)] = d_1009_v111_
            d_1014_v114_ = nw186_
            index184_ = default__.safeIndex(501, (d_1014_v114_).length(0))
            (d_1014_v114_)[index184_] = d_1009_v111_
            d_1015_v115_: _dafny.Seq
            d_1015_v115_ = _dafny.SeqWithoutIsStrInference([p3, not(True)])
            d_1016_v116_: _dafny.Array
            def lambda52_(d_1017_p3_):
                def lambda53_(d_1018_i7_):
                    return d_1017_p3_

                return lambda53_

            init28_ = lambda52_(p3)
            nw187_ = _dafny.Array(None, 17)
            for i0_28_ in range(nw187_.length(0)):
                nw187_[i0_28_] = init28_(i0_28_)
            d_1016_v116_ = nw187_
            d_1019_v117_: _dafny.Map
            d_1019_v117_ = _dafny.Map({d_1008_v110_.f12: d_1016_v116_})
            index185_ = default__.safeIndex(501, (d_1014_v114_).length(0))
            nw188_ = _dafny.Array(None, 5)
            nw188_[int(0)] = len(d_1015_v115_)
            nw188_[int(1)] = self.f11
            nw188_[int(2)] = len((d_1019_v117_).set(self.f11, d_1016_v116_))
            nw188_[int(3)] = d_1008_v110_.f12
            nw188_[int(4)] = (self).f10
            (d_1014_v114_)[index185_] = nw188_
            (d_1008_v110_).f12 = len((p0) + ((p0) + (p0)))
            d_1020_v118_: _dafny.Array
            def lambda54_(d_1021_v105_):
                def lambda55_(d_1022_i8_):
                    return (d_1021_v105_).f15

                return lambda55_

            init29_ = lambda54_(d_1003_v105_)
            nw189_ = _dafny.Array(None, 10)
            for i0_29_ in range(nw189_.length(0)):
                nw189_[i0_29_] = init29_(i0_29_)
            d_1020_v118_ = nw189_
            index186_ = default__.safeIndex(310, (d_1020_v118_).length(0))
            (d_1020_v118_)[index186_] = (d_1003_v105_).f15
            index187_ = default__.safeIndex(310, (d_1020_v118_).length(0))
            (d_1020_v118_)[index187_] = p1
            index188_ = default__.safeIndex(824, (d_1016_v116_).length(0))
            (d_1016_v116_)[index188_] = not(p2)
            d_1023_v119_: _dafny.Map
            d_1023_v119_ = _dafny.Map({p3: p2})
            index189_ = default__.safeIndex(824, (d_1016_v116_).length(0))
            (d_1016_v116_)[index189_] = (((d_1023_v119_)[default__.fm1((d_1003_v105_).fm4(globalState), p3, self.f11, not(False), globalState)] if (default__.fm1((d_1003_v105_).fm4(globalState), p3, self.f11, not(False), globalState)) in (d_1023_v119_) else p2) if p3 else p3)
        elif True:
            d_1024_v120_: _dafny.Array
            def lambda56_(d_1025_i9_):
                return default__.safeModuloInt(d_1025_i9_, self.f11)

            init30_ = lambda56_
            nw190_ = _dafny.Array(None, 1)
            for i0_30_ in range(nw190_.length(0)):
                nw190_[i0_30_] = init30_(i0_30_)
            d_1024_v120_ = nw190_
            d_1026_v121_: D1
            d_1026_v121_ = D1_DC4(d_1024_v120_)
            d_1027_v122_: _dafny.Set
            d_1027_v122_ = _dafny.Set({D1_DC4(d_1024_v120_), d_1026_v121_, d_1026_v121_})
            d_1028_v123_: _dafny.Map
            d_1028_v123_ = _dafny.Map({(d_1003_v105_).f15: d_1027_v122_})
            d_1029_v124_: _dafny.Map
            d_1029_v124_ = _dafny.Map({D7_DC20(_dafny.Map({p3: d_1024_v120_})): self.f11})
            d_1030_v125_: _dafny.Seq
            d_1030_v125_ = _dafny.SeqWithoutIsStrInference([(self).f10])
            d_1031_v126_: _dafny.Map
            d_1031_v126_ = _dafny.Map({p2: d_1030_v125_})
            d_1032_v127_: _dafny.Seq
            d_1032_v127_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({d_1026_v121_})) == (((d_1028_v123_)[(d_1003_v105_).f15] if ((d_1003_v105_).f15) in (d_1028_v123_) else d_1027_v122_)), ((0) - (((d_1029_v124_)[D7_DC20(_dafny.Map({p3: d_1024_v120_}))] if (D7_DC20(_dafny.Map({p3: d_1024_v120_}))) in (d_1029_v124_) else len(d_1031_v126_)))) == (self.f11), p3, ((self).f10) > (self.f11)])
            d_1032_v127_ = default__.fm9(p2, globalState)
            d_1033_v128_: bool
            d_1033_v128_ = False
            d_1033_v128_ = p2
            d_1034_v129_: D1
            d_1034_v129_ = D1_DC5(p2)
            d_1035_v130_: D1
            d_1035_v130_ = D1_DC6(d_1034_v129_)
            source15_ = d_1035_v130_
            if source15_.is_DC5:
                d_1036___mcc_h15_ = source15_.cf8
                d_1037_cf8_ = d_1036___mcc_h15_
                d_1038_v131_: _dafny.Map
                d_1038_v131_ = _dafny.Map({not(d_1033_v128_): p0})
                d_1039_v132_: _dafny.Seq
                d_1039_v132_ = _dafny.SeqWithoutIsStrInference([((d_1038_v131_)[p3] if (p3) in (d_1038_v131_) else p0)])
                d_1040_v133_: _dafny.Map
                d_1040_v133_ = _dafny.Map({self.f11: d_1039_v132_})
                d_1041_v134_: _dafny.Map
                d_1041_v134_ = _dafny.Map({self.f11: ((d_1040_v133_)[self.f11] if (self.f11) in (d_1040_v133_) else d_1039_v132_)})
                d_1039_v132_ = ((d_1041_v134_)[((self).f10) + (len(d_1032_v127_))] if (((self).f10) + (len(d_1032_v127_))) in (d_1041_v134_) else d_1039_v132_)
                d_1037_cf8_ = p2
                d_1042_v136_: C0
                nw191_ = C0()
                def iife58_():
                    coll24_ = _dafny.Map()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(-138, 227):
                        d_1043_v135_: int = compr_24_
                        if ((-138) <= (d_1043_v135_)) and ((d_1043_v135_) < (227)):
                            coll24_[(d_1043_v135_) + (self.f11)] = d_1037_cf8_
                    return _dafny.Map(coll24_)
                nw191_.ctor__(iife58_()
                , self.f11)
                d_1042_v136_ = nw191_
                d_1044_v137_: D1
                d_1044_v137_ = D1_DC5(d_1033_v128_)
                d_1045_v138_: C2
                nw192_ = C2()
                nw192_.ctor__((d_1044_v137_).cf8, (self).f10)
                d_1045_v138_ = nw192_
            elif source15_.is_DC4:
                d_1046___mcc_h16_ = source15_.cf7
                d_1047_cf7_ = d_1046___mcc_h16_
                index190_ = default__.safeIndex(549, (d_1047_cf7_).length(0))
                (d_1047_cf7_)[index190_] = self.f11
                index191_ = default__.safeIndex(549, (d_1047_cf7_).length(0))
                (d_1047_cf7_)[index191_] = self.f11
                d_1048_v140_: _dafny.Map
                d_1048_v140_ = _dafny.Map({(d_1047_cf7_)[default__.safeIndex(549, (d_1047_cf7_).length(0))]: (d_1003_v105_).f15})
                d_1049_v141_: _dafny.Map
                d_1049_v141_ = _dafny.Map({(self).f10: True})
                d_1050_v142_: C0
                nw193_ = C0()
                def iife59_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in (d_1048_v140_).keys.Elements:
                        d_1051_v139_: int = compr_25_
                        if (d_1051_v139_) in (d_1048_v140_):
                            coll25_[(d_1051_v139_) + (self.f11)] = p2
                    return _dafny.Map(coll25_)
                nw193_.ctor__((iife59_()
                ) | (d_1049_v141_), (self).f10)
                d_1050_v142_ = nw193_
                d_1052_v143_: _dafny.Seq
                d_1052_v143_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puri"))
                d_1052_v143_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))) + (_dafny.SeqWithoutIsStrInference([(d_1003_v105_).f15 for d_1053_i10_ in range(default__.abs(651))]))) + (d_1052_v143_)
                d_1054_v144_: _dafny.Seq
                d_1054_v144_ = d_1052_v143_
                d_1052_v143_ = ((d_1054_v144_)) + (_dafny.SeqWithoutIsStrInference([(d_1003_v105_).f15 for d_1055_i11_ in range(default__.abs(-971))]))
            elif True:
                d_1056___mcc_h17_ = source15_.cf9
                d_1057_cf9_ = d_1056___mcc_h17_
                d_1033_v128_ = p3
                d_1058_v145_: D5
                d_1058_v145_ = D5_DC14(p2)
                d_1059_v146_: _dafny.Array
                nw194_ = _dafny.Array(None, 11)
                nw194_[int(0)] = p3
                nw194_[int(1)] = d_1033_v128_
                nw194_[int(2)] = (_dafny.SeqWithoutIsStrInference([(self).f10])) == (d_1030_v125_)
                nw194_[int(3)] = d_1033_v128_
                nw194_[int(4)] = False
                nw194_[int(5)] = d_1033_v128_
                nw194_[int(6)] = d_1033_v128_
                nw194_[int(7)] = p2
                nw194_[int(8)] = p3
                nw194_[int(9)] = not(default__.fm1((self).f10, not(d_1033_v128_), (self).f10, d_1033_v128_, globalState))
                nw194_[int(10)] = (d_1058_v145_).cf20
                d_1059_v146_ = nw194_
                d_1060_v147_: _dafny.Set
                d_1060_v147_ = _dafny.Set({p3, not(p3)})
                index192_ = default__.safeIndex(69, (d_1059_v146_).length(0))
                (d_1059_v146_)[index192_] = default__.fm1(len(d_1060_v147_), p3, (self).f10, p2, globalState)
                index193_ = default__.safeIndex(69, (d_1059_v146_).length(0))
                (d_1059_v146_)[index193_] = True
                d_1061_v148_: _dafny.Array
                def lambda57_(d_1062_p0_):
                    def lambda58_(d_1063_i12_):
                        return (_dafny.SeqWithoutIsStrInference([len(d_1062_p0_)])) + (_dafny.SeqWithoutIsStrInference([self.f11]))

                    return lambda58_

                init31_ = lambda57_(p0)
                nw195_ = _dafny.Array(None, 3)
                for i0_31_ in range(nw195_.length(0)):
                    nw195_[i0_31_] = init31_(i0_31_)
                d_1061_v148_ = nw195_
                index194_ = default__.safeIndex(38, (d_1061_v148_).length(0))
                (d_1061_v148_)[index194_] = d_1030_v125_
                index195_ = default__.safeIndex(38, (d_1061_v148_).length(0))
                (d_1061_v148_)[index195_] = (d_1030_v125_).set(default__.safeIndex(self.f11, len(d_1030_v125_)), (self).f10)
                d_1033_v128_ = d_1033_v128_
            d_1033_v128_ = not(default__.fm1(self.f11, d_1033_v128_, (self).f10, (p3) == (p3), globalState))
            d_1064_v149_: _dafny.Map
            d_1064_v149_ = _dafny.Map({p2: d_1033_v128_})
            (globalState).f5 = (len(d_1064_v149_) if default__.fm1(self.f11, d_1033_v128_, (self).f10, d_1033_v128_, globalState) else (self).f10)
        d_1065_i13_: int
        d_1065_i13_ = 0
        with _dafny.label("6"):
            while not(not (not(p3)) or (True)):
                with _dafny.c_label("6"):
                    if (d_1065_i13_) >= (100):
                        raise _dafny.Break("6")
                    d_1065_i13_ = (d_1065_i13_) + (1)
                    d_1066_v150_: bool
                    d_1066_v150_ = False
                    d_1067_v151_: _dafny.Array
                    def lambda59_(d_1068_v150_):
                        def lambda60_(d_1069_i14_):
                            return d_1068_v150_

                        return lambda60_

                    init32_ = lambda59_(d_1066_v150_)
                    nw196_ = _dafny.Array(None, 22)
                    for i0_32_ in range(nw196_.length(0)):
                        nw196_[i0_32_] = init32_(i0_32_)
                    d_1067_v151_ = nw196_
                    d_1070_v152_: _dafny.Set
                    d_1070_v152_ = _dafny.Set({d_1067_v151_})
                    d_1066_v150_ = (d_1070_v152_).ispropersubset((d_1070_v152_) - (d_1070_v152_))
                    d_1071_v154_: _dafny.MultiSet
                    d_1071_v154_ = _dafny.MultiSet([False, False])
                    d_1072_v155_: _dafny.Map
                    d_1072_v155_ = _dafny.Map({(d_1071_v154_).cardinality: True})
                    d_1073_v156_: D9
                    def iife60_():
                        coll26_ = _dafny.Map()
                        compr_26_: _dafny.Map
                        for compr_26_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1072_v155_]))).Elements:
                            d_1074_v153_: _dafny.Map = compr_26_
                            if (d_1074_v153_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1072_v155_]))):
                                coll26_[d_1074_v153_] = (d_1003_v105_).f15
                        return _dafny.Map(coll26_)
                    d_1073_v156_ = D9_DC27(d_1003_v105_, self.f11, self.f11, (self).f10, iife60_()
)
                    d_1066_v150_ = default__.fm1((self).f10, p3, len((_dafny.Set({(self).f10, self.f11, (d_1073_v156_).cf34, (self).f10})) | (d_854_v1_)), not(not((not(p3) if default__.fm1((0) - (self.f11), p3, (self).f10, False, globalState) else d_1066_v150_))), globalState)
                    d_1075_v157_: _dafny.Seq
                    d_1075_v157_ = _dafny.SeqWithoutIsStrInference([self.f11])
                    d_1075_v157_ = d_1075_v157_
                    if d_1066_v150_:
                        d_1076_v158_: _dafny.Array
                        nw197_ = _dafny.Array(None, 5)
                        nw197_[int(0)] = d_1075_v157_
                        nw197_[int(1)] = d_1075_v157_
                        nw197_[int(2)] = d_1075_v157_
                        nw197_[int(3)] = d_1075_v157_
                        nw197_[int(4)] = d_1075_v157_
                        d_1076_v158_ = nw197_
                        index196_ = default__.safeIndex(474, (d_1076_v158_).length(0))
                        (d_1076_v158_)[index196_] = _dafny.SeqWithoutIsStrInference([(d_1075_v157_)[default__.safeIndex(len(p0), len(d_1075_v157_))], (self).f10])
                        d_1077_v159_: _dafny.Array
                        def lambda61_(d_1078_i15_):
                            return D2_DC7((self).f10)

                        init33_ = lambda61_
                        nw198_ = _dafny.Array(None, 12)
                        for i0_33_ in range(nw198_.length(0)):
                            nw198_[i0_33_] = init33_(i0_33_)
                        d_1077_v159_ = nw198_
                        d_1079_v160_: T1
                        nw199_ = C3()
                        nw199_.ctor__(d_1077_v159_, (p0)[default__.safeIndex((self).f10, len(p0))], (self).f10)
                        d_1079_v160_ = nw199_
                        index197_ = default__.safeIndex(474, (d_1076_v158_).length(0))
                        rhs166_ = (_dafny.SeqWithoutIsStrInference([-973, d_1079_v160_.f12, -185])) + (d_1075_v157_)
                        rhs167_ = -939
                        rhs168_ = ((self).fm2(False, globalState)) * (d_1079_v160_.f12)
                        rhs169_ = d_1079_v160_
                        rhs170_ = default__.fm1(599, d_1066_v150_, self.f11, d_1066_v150_, globalState)
                        lhs106_ = d_1076_v158_
                        lhs107_ = default__.safeIndex(474, (d_1076_v158_).length(0))
                        lhs108_ = globalState
                        lhs106_[lhs107_] = rhs166_
                        lhs108_.f5 = rhs167_
                        r0 = rhs168_
                        d_1079_v160_ = rhs169_
                        d_1066_v150_ = rhs170_
                        d_1080_v161_: _dafny.Map
                        d_1080_v161_ = _dafny.Map({d_1066_v150_: (self).f10})
                        d_1081_v162_: _dafny.Seq
                        d_1081_v162_ = _dafny.SeqWithoutIsStrInference([True])
                        rhs171_ = (d_1079_v160_.f12) <= (len((d_1080_v161_) | (d_1080_v161_)))
                        rhs172_ = default__.fm1(len(d_1072_v155_), (False) in (d_1081_v162_), ((self).f10 if d_1066_v150_ else (0) - (len(p0))), p3, globalState)
                        rhs173_ = (d_1079_v160_.f12) == ((self).f10)
                        d_1066_v150_ = rhs171_
                        d_1066_v150_ = rhs172_
                        d_1066_v150_ = rhs173_
                        index198_ = default__.safeIndex(155, (d_1067_v151_).length(0))
                        (d_1067_v151_)[index198_] = not(p2)
                        index199_ = default__.safeIndex(155, (d_1067_v151_).length(0))
                        (d_1067_v151_)[index199_] = p3
                        d_1082_v163_: _dafny.Seq
                        d_1082_v163_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({True: (self).f10})])
                        d_1083_v164_: T0
                        nw200_ = C2()
                        nw200_.ctor__(p3, (0) - (default__.fm0(len((d_1082_v163_)[default__.safeIndex((self).f10, len(d_1082_v163_))]), len(d_1080_v161_), globalState)))
                        d_1083_v164_ = nw200_
                        d_1083_v164_ = d_1083_v164_
                        index200_ = default__.safeIndex(155, (d_1067_v151_).length(0))
                        (d_1067_v151_)[index200_] = p3
                    elif True:
                        d_1066_v150_ = p3
                        d_1084_v165_: _dafny.Array
                        def lambda62_(d_1085_v150_, d_1086_p2_):
                            def lambda63_(d_1087_i16_):
                                return D5_DC13(not(d_1085_v150_), not(d_1086_p2_))

                            return lambda63_

                        init34_ = lambda62_(d_1066_v150_, p2)
                        nw201_ = _dafny.Array(None, 5)
                        for i0_34_ in range(nw201_.length(0)):
                            nw201_[i0_34_] = init34_(i0_34_)
                        d_1084_v165_ = nw201_
                        d_1088_v166_: D5
                        d_1088_v166_ = D5_DC13(not(p3), p3)
                        index201_ = default__.safeIndex(792, (d_1084_v165_).length(0))
                        (d_1084_v165_)[index201_] = d_1088_v166_
                        index202_ = default__.safeIndex(792, (d_1084_v165_).length(0))
                        (d_1084_v165_)[index202_] = d_1088_v166_
                        d_1089_v167_: C1
                        nw202_ = C1()
                        nw202_.ctor__((d_1003_v105_).f15, (d_1003_v105_).f15, self.f11)
                        d_1089_v167_ = nw202_
                        d_1090_v168_: C1
                        nw203_ = C1()
                        nw203_.ctor__((d_1089_v167_).f15, (d_1089_v167_).f15, (self).f10)
                        d_1090_v168_ = nw203_
                        d_1091_v169_: _dafny.Array
                        nw204_ = _dafny.Array(int(0), 5)
                        d_1091_v169_ = nw204_
                        d_1092_v170_: _dafny.Set
                        d_1092_v170_ = _dafny.Set({p2, True, d_1066_v150_})
                        index203_ = default__.safeIndex(924, (d_1091_v169_).length(0))
                        (d_1091_v169_)[index203_] = len(d_1092_v170_)
                        index204_ = default__.safeIndex(924, (d_1091_v169_).length(0))
                        (d_1091_v169_)[index204_] = (self).f10
                    pass
            pass
        if (D2_DC8(p3, self.f11, p2)).cf11:
            d_1093_v171_: _dafny.Map
            d_1093_v171_ = _dafny.Map({(self).f10: p2})
            d_1094_v172_: C0
            nw205_ = C0()
            nw205_.ctor__(d_1093_v171_, (self).f10)
            d_1094_v172_ = nw205_
            d_1095_v173_: _dafny.Map
            d_1095_v173_ = _dafny.Map({(d_1003_v105_).f15: d_1094_v172_})
            d_1095_v173_ = d_1095_v173_
            if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eubmv")))) > ((self).f10):
                d_1096_v174_: int
                d_1097_v175_: bool
                d_1098_v176_: _dafny.Map
                out54_: int
                out55_: bool
                out56_: _dafny.Map
                out54_, out55_, out56_ = (d_1003_v105_).m5(globalState)
                d_1096_v174_ = out54_
                d_1097_v175_ = out55_
                d_1098_v176_ = out56_
                d_1099_v177_: _dafny.Set
                d_1099_v177_ = _dafny.Set({d_1097_v175_})
                d_1100_v178_: D0
                d_1100_v178_ = D0_DC0(True)
                d_1101_v179_: _dafny.Array
                nw206_ = _dafny.Array(None, 28)
                nw206_[int(0)] = d_1099_v177_
                nw206_[int(1)] = d_1099_v177_
                nw206_[int(2)] = _dafny.Set({p2})
                nw206_[int(3)] = (d_1099_v177_) | (d_1099_v177_)
                nw206_[int(4)] = _dafny.Set({True})
                nw206_[int(5)] = _dafny.Set({d_1097_v175_})
                nw206_[int(6)] = d_1099_v177_
                nw206_[int(7)] = (_dafny.Set({p3, p3})) - (d_1099_v177_)
                nw206_[int(8)] = d_1099_v177_
                nw206_[int(9)] = _dafny.Set({(d_1100_v178_).cf0, p3, not(d_1097_v175_), default__.fm1(818, p3, d_1096_v174_, p2, globalState)})
                nw206_[int(10)] = d_1099_v177_
                nw206_[int(11)] = d_1099_v177_
                nw206_[int(12)] = _dafny.Set({p2, p2})
                nw206_[int(13)] = _dafny.Set({p3})
                nw206_[int(14)] = d_1099_v177_
                nw206_[int(15)] = d_1099_v177_
                nw206_[int(16)] = d_1099_v177_
                nw206_[int(17)] = d_1099_v177_
                nw206_[int(18)] = _dafny.Set({d_1097_v175_, d_1097_v175_, p3})
                nw206_[int(19)] = (d_1099_v177_) - (d_1099_v177_)
                nw206_[int(20)] = d_1099_v177_
                nw206_[int(21)] = (d_1099_v177_) | (_dafny.Set({p3}))
                nw206_[int(22)] = d_1099_v177_
                nw206_[int(23)] = d_1099_v177_
                nw206_[int(24)] = (d_1099_v177_) - (d_1099_v177_)
                nw206_[int(25)] = d_1099_v177_
                nw206_[int(26)] = (d_1099_v177_).intersection(_dafny.Set({p3}))
                nw206_[int(27)] = d_1099_v177_
                d_1101_v179_ = nw206_
                d_1102_v180_: _dafny.Set
                d_1102_v180_ = d_1099_v177_
                index205_ = default__.safeIndex(504, (d_1101_v179_).length(0))
                (d_1101_v179_)[index205_] = ((d_1102_v180_)).intersection((d_1102_v180_))
                index206_ = default__.safeIndex(504, (d_1101_v179_).length(0))
                rhs174_ = False
                rhs175_ = (d_1099_v177_) - ((d_1099_v177_) - (d_1099_v177_))
                lhs109_ = d_1101_v179_
                lhs110_ = default__.safeIndex(504, (d_1101_v179_).length(0))
                d_1097_v175_ = rhs174_
                lhs109_[lhs110_] = rhs175_
                d_1103_v181_: _dafny.Array
                def lambda64_(d_1104_p0_):
                    def lambda65_(d_1105_i17_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhjkitey"))) + (d_1104_p0_)

                    return lambda65_

                init35_ = lambda64_(p0)
                nw207_ = _dafny.Array(None, 8)
                for i0_35_ in range(nw207_.length(0)):
                    nw207_[i0_35_] = init35_(i0_35_)
                d_1103_v181_ = nw207_
                index207_ = default__.safeIndex(207, (d_1103_v181_).length(0))
                (d_1103_v181_)[index207_] = (p0) + (p0)
                d_1106_v182_: _dafny.Seq
                d_1106_v182_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_1096_v174_, False, (self).f10, p3, globalState)])
                d_1107_v183_: D8
                d_1107_v183_ = D8_DC24(d_1099_v177_, d_1106_v182_)
                pat_let_tv8_ = d_1101_v179_
                pat_let_tv9_ = d_1101_v179_
                d_1108_v184_: _dafny.MultiSet
                def iife61_(_pat_let17_0):
                    def iife62_(d_1109_dt__update__tmp_h2_):
                        def iife63_(_pat_let18_0):
                            def iife64_(d_1110_dt__update_hcf30_h0_):
                                return D8_DC24(d_1110_dt__update_hcf30_h0_, (d_1109_dt__update__tmp_h2_).cf31)
                            return iife64_(_pat_let18_0)
                        return iife63_((pat_let_tv9_)[default__.safeIndex(504, (pat_let_tv8_).length(0))])
                    return iife62_(_pat_let17_0)
                d_1108_v184_ = _dafny.MultiSet([iife61_(d_1107_v183_)])
                index208_ = default__.safeIndex(207, (d_1103_v181_).length(0))
                (d_1103_v181_)[index208_] = (p0 if (d_1108_v184_).issubset(d_1108_v184_) else p0)
                d_1111_v185_: _dafny.Array
                def lambda66_(d_1112_v106_):
                    def lambda67_(d_1113_i18_):
                        return (d_1113_i18_) + ((d_1112_v106_).cardinality)

                    return lambda67_

                init36_ = lambda66_(d_1004_v106_)
                nw208_ = _dafny.Array(None, 4)
                for i0_36_ in range(nw208_.length(0)):
                    nw208_[i0_36_] = init36_(i0_36_)
                d_1111_v185_ = nw208_
                d_1114_v186_: _dafny.Array
                def lambda68_(d_1115_v1_):
                    def lambda69_(d_1116_i19_):
                        return D2_DC7(len(d_1115_v1_))

                    return lambda69_

                init37_ = lambda68_(d_854_v1_)
                nw209_ = _dafny.Array(None, 2)
                for i0_37_ in range(nw209_.length(0)):
                    nw209_[i0_37_] = init37_(i0_37_)
                d_1114_v186_ = nw209_
                d_1117_v187_: T1
                nw210_ = C3()
                nw210_.ctor__(d_1114_v186_, (d_1003_v105_).f15, (self).f10)
                d_1117_v187_ = nw210_
                d_1118_v188_: _dafny.Seq
                d_1118_v188_ = _dafny.SeqWithoutIsStrInference([d_1117_v187_])
                d_1119_v189_: _dafny.Map
                d_1119_v189_ = _dafny.Map({d_1111_v185_: d_1118_v188_})
                d_1119_v189_ = (d_1119_v189_).set((d_1111_v185_ if False else d_1111_v185_), (d_1118_v188_) + (d_1118_v188_))
                (globalState).f5 = default__.safeDivisionInt((self).f10, (self).f10)
            elif True:
                d_1120_v190_: bool
                d_1120_v190_ = False
                d_1120_v190_ = p3
                d_1121_v191_: _dafny.Map
                d_1121_v191_ = _dafny.Map({(d_1003_v105_).f15: _dafny.MultiSet((d_1003_v105_).fm5(p2, not(p2), False, not(p3), globalState))})
                d_1122_v192_: _dafny.Seq
                d_1122_v192_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_1123_v193_: D0
                d_1123_v193_ = D0_DC0(p2)
                d_1124_v194_: _dafny.Seq
                d_1124_v194_ = _dafny.SeqWithoutIsStrInference([d_1120_v190_])
                d_1125_v196_: _dafny.Array
                nw211_ = _dafny.Array(None, 20)
                nw211_[int(0)] = False
                nw211_[int(1)] = (default__.fm21((self).f10, (self).f10, _dafny.CodePoint('h'), (self).f10, globalState)).issubset((((d_1121_v191_)[(d_1003_v105_).f15] if ((d_1003_v105_).f15) in (d_1121_v191_) else d_1004_v106_)).set((self).f10, default__.abs(-600)))
                nw211_[int(2)] = p2
                nw211_[int(3)] = p3
                nw211_[int(4)] = p2
                nw211_[int(5)] = p2
                nw211_[int(6)] = not(p2)
                nw211_[int(7)] = (len(d_1122_v192_)) <= (((d_1004_v106_)[self.f11] if (self.f11) in (d_1004_v106_) else self.f11))
                nw211_[int(8)] = (d_1123_v193_).cf0
                nw211_[int(9)] = p3
                nw211_[int(10)] = default__.fm1(len(d_1122_v192_), default__.fm1(len(_dafny.Set({(self).f10})), p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wma"))), (d_1124_v194_)[default__.safeIndex((self).f10, len(d_1124_v194_))], globalState), (0) - ((0) - ((self).f10)), d_1120_v190_, globalState)
                nw211_[int(11)] = True
                nw211_[int(12)] = True
                nw211_[int(13)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1126_i20_ in range(default__.abs(61))])) < (p0)
                nw211_[int(14)] = p2
                def iife65_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(612, -261):
                        d_1127_v195_: int = compr_27_
                        if ((612) <= (d_1127_v195_)) and ((d_1127_v195_) < (-261)):
                            coll27_ = coll27_.union(_dafny.Set([default__.safeDivisionInt(d_1127_v195_, (self).f10)]))
                    return _dafny.Set(coll27_)
                nw211_[int(15)] = not((iife65_()
                ).issubset(d_854_v1_))
                nw211_[int(16)] = d_1120_v190_
                nw211_[int(17)] = d_1120_v190_
                nw211_[int(18)] = ((self).f10) >= (self.f11)
                nw211_[int(19)] = False
                d_1125_v196_ = nw211_
                index209_ = default__.safeIndex(763, (d_1125_v196_).length(0))
                (d_1125_v196_)[index209_] = p2
                d_1128_v197_: _dafny.Set
                d_1128_v197_ = _dafny.Set({False})
                d_1129_v198_: _dafny.Set
                d_1129_v198_ = d_1128_v197_
                d_1130_v199_: _dafny.Seq
                d_1130_v199_ = _dafny.SeqWithoutIsStrInference([d_1094_v172_])
                d_1131_v200_: _dafny.Array
                nw212_ = _dafny.Array(int(0), 16)
                d_1131_v200_ = nw212_
                d_1132_v201_: _dafny.Seq
                d_1132_v201_ = _dafny.SeqWithoutIsStrInference([d_1131_v200_, d_1131_v200_, d_1131_v200_])
                index210_ = default__.safeIndex(763, (d_1125_v196_).length(0))
                rhs176_ = default__.safeModuloInt(len(d_854_v1_), len(d_1130_v199_))
                rhs177_ = ((self).f10) * ((self).f10)
                rhs178_ = ((d_1004_v106_).intersection(d_1004_v106_)).ispropersubset(d_1004_v106_)
                rhs179_ = d_1129_v198_
                rhs180_ = (d_1132_v201_)[default__.safeIndex(self.f11, len(d_1132_v201_))]
                lhs111_ = globalState
                lhs112_ = globalState
                lhs113_ = d_1125_v196_
                lhs114_ = default__.safeIndex(763, (d_1125_v196_).length(0))
                lhs111_.f5 = rhs176_
                lhs112_.f5 = rhs177_
                lhs113_[lhs114_] = rhs178_
                d_1129_v198_ = rhs179_
                r3 = rhs180_
                d_1133_v202_: D1
                d_1133_v202_ = D1_DC5(p3)
                d_1120_v190_ = (d_1133_v202_).cf8
                nw213_ = C0()
                def iife66_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(361, 634):
                        d_1134_v203_: int = compr_28_
                        if ((361) <= (d_1134_v203_)) and ((d_1134_v203_) < (634)):
                            coll28_[default__.safeDivisionInt(d_1134_v203_, (self).f10)] = True
                    return _dafny.Map(coll28_)
                nw213_.ctor__(iife66_()
                , self.f11)
                d_1094_v172_ = nw213_
                d_1135_v204_: _dafny.Map
                d_1135_v204_ = _dafny.Map({(d_1122_v192_) + (d_1122_v192_): d_1124_v194_})
                d_1124_v194_ = ((d_1135_v204_)[d_1122_v192_] if (d_1122_v192_) in (d_1135_v204_) else _dafny.SeqWithoutIsStrInference([d_1120_v190_]))
            d_1136_v205_: _dafny.Map
            d_1136_v205_ = _dafny.Map({not(p3): (p2) and (p3)})
            d_1136_v205_ = (d_1136_v205_).set(p3, p3)
            d_1137_v206_: _dafny.Seq
            d_1137_v206_ = _dafny.SeqWithoutIsStrInference([self.f11])
            (globalState).f5 = (d_1137_v206_)[default__.safeIndex((d_1004_v106_).cardinality, len(d_1137_v206_))]
            d_1138_v207_: _dafny.Array
            nw214_ = _dafny.Array(False, 21)
            d_1138_v207_ = nw214_
            index211_ = default__.safeIndex(232, (d_1138_v207_).length(0))
            (d_1138_v207_)[index211_] = (p2 if p3 else p3)
            index212_ = default__.safeIndex(232, (d_1138_v207_).length(0))
            (d_1138_v207_)[index212_] = p3
        elif True:
            d_1139_v208_: _dafny.Map
            d_1139_v208_ = _dafny.Map({p2: p3})
            d_1139_v208_ = (d_1139_v208_).set(p3, default__.fm1((self).f10, p2, self.f11, not(p2), globalState))
            d_1140_v209_: _dafny.Array
            def lambda70_(d_1141_p3_, d_1142_p2_):
                def lambda71_(d_1143_i21_):
                    return _dafny.Set({d_1141_p3_, d_1142_p2_, d_1142_p2_, not(True), d_1142_p2_})

                return lambda71_

            init38_ = lambda70_(p3, p2)
            nw215_ = _dafny.Array(None, 22)
            for i0_38_ in range(nw215_.length(0)):
                nw215_[i0_38_] = init38_(i0_38_)
            d_1140_v209_ = nw215_
            d_1140_v209_ = d_1140_v209_
            d_1144_v210_: D2
            d_1144_v210_ = D2_DC7(self.f11)
            d_1145_v211_: _dafny.Array
            nw216_ = _dafny.Array(None, 3)
            nw216_[int(0)] = d_1144_v210_
            nw216_[int(1)] = D2_DC7(self.f11)
            nw216_[int(2)] = d_1144_v210_
            d_1145_v211_ = nw216_
            d_1146_v212_: C3
            nw217_ = C3()
            nw217_.ctor__(d_1145_v211_, (d_1003_v105_).f15, self.f11)
            d_1146_v212_ = nw217_
            d_1147_v213_: _dafny.Map
            d_1147_v213_ = _dafny.Map({d_1146_v212_: self.f11})
            d_1148_v214_: _dafny.MultiSet
            d_1148_v214_ = _dafny.MultiSet([d_1147_v213_, d_1147_v213_])
            d_1149_v215_: _dafny.MultiSet
            d_1149_v215_ = _dafny.MultiSet([(d_1148_v214_).isdisjoint(d_1148_v214_)])
            d_1149_v215_ = d_1149_v215_
            d_1150_v216_: _dafny.Array
            def lambda72_(d_1151_i22_):
                return (d_1151_i22_) - ((self).f10)

            init39_ = lambda72_
            nw218_ = _dafny.Array(None, 10)
            for i0_39_ in range(nw218_.length(0)):
                nw218_[i0_39_] = init39_(i0_39_)
            d_1150_v216_ = nw218_
            index213_ = default__.safeIndex(546, (d_1150_v216_).length(0))
            (d_1150_v216_)[index213_] = (0) - ((self).f10)
            index214_ = default__.safeIndex(546, (d_1150_v216_).length(0))
            (d_1150_v216_)[index214_] = len(p0)
            d_1152_v217_: _dafny.Array
            nw219_ = _dafny.Array(False, 25)
            d_1152_v217_ = nw219_
            d_1153_v218_: D0
            d_1153_v218_ = D0_DC1(d_1152_v217_)
            d_1153_v218_ = d_1153_v218_
        d_854_v1_ = d_854_v1_
        r0 = -907
        d_1154_v219_: _dafny.MultiSet
        d_1154_v219_ = _dafny.MultiSet([p2])
        d_1155_v220_: _dafny.Array
        nw220_ = _dafny.Array(None, 5)
        nw220_[int(0)] = (p0) < (p0)
        nw220_[int(1)] = (self.f11) == (len(_dafny.SeqWithoutIsStrInference([(d_1003_v105_).f15 for d_1156_i23_ in range(default__.abs(302))])))
        nw220_[int(2)] = not(p3)
        nw220_[int(3)] = not(default__.fm1((self).f10, p2, (d_1154_v219_).cardinality, p2, globalState))
        nw220_[int(4)] = p3
        d_1155_v220_ = nw220_
        r1 = d_1155_v220_
        d_1157_v221_: _dafny.Array
        def lambda73_(d_1158_v2_):
            def lambda74_(d_1159_i24_):
                return d_1158_v2_

            return lambda74_

        init40_ = lambda73_(d_855_v2_)
        nw221_ = _dafny.Array(None, 3)
        for i0_40_ in range(nw221_.length(0)):
            nw221_[i0_40_] = init40_(i0_40_)
        d_1157_v221_ = nw221_
        d_1160_v222_: _dafny.Set
        d_1160_v222_ = _dafny.Set({d_1157_v221_})
        r2 = ((d_1160_v222_).intersection(d_1160_v222_)) | (d_1160_v222_)
        d_1161_v223_: _dafny.Array
        nw222_ = _dafny.Array(int(0), 1)
        d_1161_v223_ = nw222_
        r3 = d_1161_v223_
        return r0, r1, r2, r3

    @property
    def f10(self):
        return self._f10
