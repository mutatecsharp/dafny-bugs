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
    def fm4(p0, globalState):
        return (_dafny.Map({-771: not(False)})) | (_dafny.Map({523: False}))

    @staticmethod
    def fm5(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvxthx"))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.Set({_dafny.CodePoint('j')})).Elements:
                d_0_v0_: str = compr_0_
                if (d_0_v0_) in (_dafny.Set({_dafny.CodePoint('j')})):
                    coll0_[d_0_v0_] = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnemwohk"))), len(_dafny.SeqWithoutIsStrInference([762]))]))))
            return _dafny.Map(coll0_)
        return (((_dafny.MultiSet([693])).intersection(_dafny.MultiSet([533, len(_dafny.SeqWithoutIsStrInference([False])), len(iife0_()
        )]))) | (_dafny.MultiSet([len((D1_DC4(not((D1_DC4(False, True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rookixm")), 276, len(_dafny.SeqWithoutIsStrInference([False])))).cf9), True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kowjt")), 344, 452)).cf10), -249, 293, (D1_DC4(False, True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1_i0_ in range(default__.abs(299))]), 361, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2_i1_ in range(default__.abs(-739))])))).cf12]))).cardinality

    @staticmethod
    def fm9(p0, globalState):
        return D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khaefkx")), default__.safeModuloInt((_dafny.MultiSet([False, False, False, False, (D9_DC23(len(_dafny.Map({(0) - (len((D10_DC25(_dafny.Set({False, False}))).cf49)): False})), (0) - ((_dafny.MultiSet([_dafny.CodePoint('o')])).cardinality), True, _dafny.Map({934: _dafny.Set({True})}), -844)).cf45])).cardinality, 112), _dafny.Map({not(False): 117}), (_dafny.SeqWithoutIsStrInference([724])) + (_dafny.SeqWithoutIsStrInference([-729])))

    @staticmethod
    def fm10(p0, p1, globalState):
        return ((D5_DC15(_dafny.MultiSet([False, not(not(True)), False, not(not(False))]))).cf31).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False, False]))))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False}))])) + (_dafny.SeqWithoutIsStrInference([678]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucej")))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({66}))])))

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return ((_dafny.Map({not(False): len(_dafny.Set({True}))})) | (_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmj")))}))) | ((_dafny.Map({True: (0) - (len(_dafny.Map({True: 718})))})) | (_dafny.Map({True: -50})))

    @staticmethod
    def fm14(p0, globalState):
        if (False if True else True):
            return D1_DC7(D1_DC3(_dafny.Map({len(_dafny.Map({True: len(_dafny.Set({True, True}))})): False})))
        elif True:
            return D1_DC7(D1_DC5(True, False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gc"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))))

    @staticmethod
    def fm15(p0, globalState):
        return _dafny.Map({(-246) < (36): _dafny.SeqWithoutIsStrInference([-678])})

    @staticmethod
    def fm16(globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('t')])).isdisjoint(_dafny.MultiSet([_dafny.CodePoint('q')]))])

    @staticmethod
    def fm17(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]) for d_3_i0_ in range(default__.abs(631))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')]) for d_4_i1_ in range(default__.abs(814))])))

    @staticmethod
    def fm18(globalState):
        if (_dafny.MultiSet([True])).isdisjoint(_dafny.MultiSet([not(not(not(not(True)))), False, not(False), False, False])):
            return D6_DC17(_dafny.SeqWithoutIsStrInference([True]))
        elif True:
            return D6_DC17(_dafny.SeqWithoutIsStrInference([not(True)]))

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([208 for d_5_i0_ in range(default__.abs(786))])), _dafny.MultiSet([-582, 809]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_6_i1_ in range(default__.abs(982))]))])})) | ((_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "seyfsdjs"))), 136])), _dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwmmx")))}))]), _dafny.MultiSet([len(_dafny.Set({True, True}))])})).intersection(_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnlhtu")))]))})))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])).Elements:
                d_7_v0_: str = compr_1_
                if (d_7_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])):
                    coll1_[d_7_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "penvbw")))
            return _dafny.Map(coll1_)
        source0_ = D4_DC13((D4_DC13(iife1_()
)).cf30)
        if source0_.is_DC14:
            return D1_DC6(True, 19, -603, False)
        elif True:
            d_8___mcc_h0_ = source0_.cf30
            d_9_cf30_ = d_8___mcc_h0_
            return D1_DC6(False, (D1_DC5(False, not(True), len(_dafny.Map({True: -945})), 284)).cf15, 915, not(not(True)))

    @staticmethod
    def fm21(globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-357, 743):
                d_10_v0_: int = compr_2_
                if ((-357) <= (d_10_v0_)) and ((d_10_v0_) < (743)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_10_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: not(False)}), _dafny.Map({True: True})])))]))
            return _dafny.Set(coll2_)
        return iife2_()
        

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return D5_DC15((_dafny.MultiSet([not(not(True))])).intersection(_dafny.MultiSet([False])))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        if True:
            return D2_DC8(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-344])))
        elif True:
            return D2_DC8(_dafny.MultiSet([852]))

    @staticmethod
    def fm24(globalState):
        return (not((_dafny.Set({-500})).ispropersubset(_dafny.Set({458, (0) - (len(_dafny.Map({len(_dafny.Set({True})): False})))})))) == ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjjecxhfo")))) > (53))

    @staticmethod
    def fm25(globalState):
        return (_dafny.Map({_dafny.CodePoint('n'): 796})) | (_dafny.Map({_dafny.CodePoint('l'): len(_dafny.Map({not(False): len(_dafny.Set({True}))}))}))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        r3: int = int(0)
        d_11_v0_: bool
        d_11_v0_ = True
        d_12_v1_: _dafny.Seq
        d_12_v1_ = _dafny.SeqWithoutIsStrInference([d_11_v0_, d_11_v0_, d_11_v0_, d_11_v0_, d_11_v0_])
        hi0_ = len(d_12_v1_)
        for d_13_i0_ in range(p1, hi0_):
            d_14_v3_: _dafny.Set
            d_14_v3_ = _dafny.Set({-156, p0})
            d_15_v4_: _dafny.Map
            d_15_v4_ = _dafny.Map({d_13_i0_: 983})
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in (d_14_v3_).Elements:
                    d_16_v2_: int = compr_3_
                    if (d_16_v2_) in (d_14_v3_):
                        coll3_[default__.safeDivisionInt(d_16_v2_, d_13_i0_)] = p0
                return _dafny.Map(coll3_)
            d_11_v0_ = (iife3_()
            ) != (d_15_v4_)
            d_17_v5_: _dafny.Array
            def lambda0_(d_18_p0_):
                def lambda1_(d_19_i1_):
                    return default__.safeDivisionInt(d_19_i1_, d_18_p0_)

                return lambda1_

            init0_ = lambda0_(p0)
            nw0_ = _dafny.Array(None, 25)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_17_v5_ = nw0_
            index0_ = default__.safeIndex(500, (d_17_v5_).length(0))
            (d_17_v5_)[index0_] = p0
            d_20_v6_: _dafny.Seq
            d_20_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "riotpp"))
            d_21_v7_: D1
            d_21_v7_ = D1_DC4(d_11_v0_, d_11_v0_, d_20_v6_, p1, p0)
            d_22_v8_: _dafny.Seq
            d_22_v8_ = _dafny.SeqWithoutIsStrInference([d_12_v1_])
            d_23_v9_: C3
            nw1_ = C3()
            nw1_.ctor__(len(((d_22_v8_)[default__.safeIndex(p0, len(d_22_v8_))]).set(default__.safeIndex(364, len((d_22_v8_)[default__.safeIndex(p0, len(d_22_v8_))])), d_11_v0_)), d_13_i0_, d_11_v0_)
            d_23_v9_ = nw1_
            d_24_v10_: _dafny.Seq
            d_24_v10_ = _dafny.SeqWithoutIsStrInference([d_23_v9_, d_23_v9_, d_23_v9_])
            d_25_v11_: _dafny.MultiSet
            d_25_v11_ = _dafny.MultiSet([d_13_i0_, p0, (0) - ((0) - (len(d_24_v10_))), d_13_i0_, len(d_20_v6_)])
            index1_ = default__.safeIndex(500, (d_17_v5_).length(0))
            rhs0_ = (p1) - ((d_21_v7_).cf12)
            rhs1_ = ((D2_DC8(d_25_v11_)).cf22).cardinality
            rhs2_ = d_13_i0_
            lhs0_ = d_17_v5_
            lhs1_ = default__.safeIndex(500, (d_17_v5_).length(0))
            lhs0_[lhs1_] = rhs0_
            r2 = rhs1_
            r3 = rhs2_
            r2 = ((p1) * (len(d_20_v6_))) + (p0)
            d_26_v12_: C2
            nw2_ = C2()
            nw2_.ctor__(default__.safeDivisionInt(p1, (d_17_v5_)[default__.safeIndex(500, (d_17_v5_).length(0))]))
            d_26_v12_ = nw2_
        r2 = p0
        d_27_v13_: str
        d_27_v13_ = _dafny.CodePoint('q')
        def lambda2_(source1_):
            d_28___mcc_h0_ = source1_
            d_29_cf40_ = d_28___mcc_h0_
            return _dafny.CodePoint('c')

        d_27_v13_ = lambda2_(d_27_v13_)
        d_30_v14_: _dafny.Array
        def lambda3_(d_31_i2_):
            return (d_31_i2_) + (834)

        init1_ = lambda3_
        nw3_ = _dafny.Array(None, 4)
        for i0_1_ in range(nw3_.length(0)):
            nw3_[i0_1_] = init1_(i0_1_)
        d_30_v14_ = nw3_
        d_32_v15_: T0
        nw4_ = C3()
        nw4_.ctor__(p0, p0, d_11_v0_)
        d_32_v15_ = nw4_
        d_33_v16_: _dafny.MultiSet
        d_33_v16_ = _dafny.MultiSet([d_32_v15_])
        index2_ = default__.safeIndex(835, (d_30_v14_).length(0))
        (d_30_v14_)[index2_] = (d_33_v16_).cardinality
        index3_ = default__.safeIndex(835, (d_30_v14_).length(0))
        (d_30_v14_)[index3_] = p0
        d_34_v17_: D4
        d_34_v17_ = D4_DC13(default__.fm25(globalState))
        pat_let_tv0_ = d_32_v15_
        pat_let_tv1_ = d_11_v0_
        pat_let_tv2_ = d_32_v15_
        def lambda4_(source2_):
            if source2_.is_DC14:
                return (pat_let_tv0_).f6
            elif True:
                d_35___mcc_h1_ = source2_.cf30
                d_36_cf30_ = d_35___mcc_h1_
                return not (pat_let_tv1_) or ((pat_let_tv2_).f6)

        d_11_v0_ = lambda4_(d_34_v17_)
        d_37_v18_: _dafny.Array
        nw5_ = _dafny.Array(_dafny.Map({}), 4)
        d_37_v18_ = nw5_
        d_38_v19_: _dafny.MultiSet
        d_38_v19_ = _dafny.MultiSet([(d_32_v15_).f6, (d_32_v15_).f6])
        d_39_v20_: C0
        nw6_ = C0()
        nw6_.ctor__(867, ((d_38_v19_)[d_11_v0_] if (d_11_v0_) in (d_38_v19_) else (d_30_v14_)[default__.safeIndex(835, (d_30_v14_).length(0))]))
        d_39_v20_ = nw6_
        d_40_v21_: _dafny.Map
        d_40_v21_ = _dafny.Map({d_39_v20_: (d_32_v15_).f6})
        d_41_v22_: _dafny.Map
        d_41_v22_ = _dafny.Map({(d_37_v18_ if d_11_v0_ else d_37_v18_): (len((d_40_v21_).set(d_39_v20_, (d_32_v15_).f6))) <= (831)})
        d_41_v22_ = (d_41_v22_).set(d_37_v18_, d_11_v0_)
        d_42_v23_: _dafny.Seq
        d_42_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjkieshnd"))
        r0 = d_42_v23_
        d_43_v24_: _dafny.Map
        d_43_v24_ = _dafny.Map({(d_32_v15_).f6: d_11_v0_})
        d_44_v25_: _dafny.Seq
        d_44_v25_ = _dafny.SeqWithoutIsStrInference([p0, (d_38_v19_).cardinality, len(_dafny.SeqWithoutIsStrInference([d_42_v23_, d_42_v23_, d_42_v23_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eao")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iolhb"))])), len(((d_43_v24_).set((d_32_v15_).f6, d_11_v0_)).set(d_11_v0_, d_11_v0_))])
        nw7_ = _dafny.Array(None, 10)
        nw7_[int(0)] = (d_30_v14_)[default__.safeIndex(835, (d_30_v14_).length(0))]
        nw7_[int(1)] = (0) - (p1)
        nw7_[int(2)] = 112
        nw7_[int(3)] = p0
        nw7_[int(4)] = (p1) * (default__.fm6((d_39_v20_).f12, (d_32_v15_).f5, (d_32_v15_).f6, (d_39_v20_).f13, globalState))
        nw7_[int(5)] = p1
        nw7_[int(6)] = (d_44_v25_)[default__.safeIndex((d_39_v20_).f12, len(d_44_v25_))]
        nw7_[int(7)] = (d_39_v20_).f12
        nw7_[int(8)] = ((d_30_v14_)[default__.safeIndex(835, (d_30_v14_).length(0))]) + ((d_39_v20_).f13)
        nw7_[int(9)] = (d_39_v20_).f12
        r1 = nw7_
        r2 = (d_30_v14_)[default__.safeIndex(835, (d_30_v14_).length(0))]
        r3 = len(_dafny.SeqWithoutIsStrInference([d_30_v14_]))
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_45_v0_: bool
        d_45_v0_ = False
        d_46_v1_: int
        d_46_v1_ = 82
        d_47_v2_: str
        d_47_v2_ = _dafny.CodePoint('d')
        d_48_v3_: _dafny.MultiSet
        d_48_v3_ = _dafny.MultiSet([d_47_v2_])
        d_49_v4_: _dafny.Map
        d_49_v4_ = _dafny.Map({d_45_v0_: d_46_v1_})
        d_50_v5_: _dafny.Map
        d_50_v5_ = _dafny.Map({d_45_v0_: d_45_v0_})
        d_51_v6_: _dafny.Map
        d_51_v6_ = _dafny.Map({len(d_50_v5_): d_45_v0_})
        d_52_v7_: _dafny.Seq
        d_52_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxasa"))
        d_53_v8_: _dafny.Seq
        d_53_v8_ = _dafny.SeqWithoutIsStrInference([d_45_v0_])
        d_54_v9_: _dafny.Array
        nw8_ = _dafny.Array(None, 29)
        nw8_[int(0)] = d_46_v1_
        nw8_[int(1)] = ((d_48_v3_)[d_47_v2_] if (d_47_v2_) in (d_48_v3_) else 87)
        nw8_[int(2)] = 872
        nw8_[int(3)] = len(d_49_v4_)
        nw8_[int(4)] = d_46_v1_
        nw8_[int(5)] = d_46_v1_
        nw8_[int(6)] = d_46_v1_
        nw8_[int(7)] = (0) - (d_46_v1_)
        nw8_[int(8)] = len(d_51_v6_)
        nw8_[int(9)] = len(_dafny.SeqWithoutIsStrInference([d_46_v1_, d_46_v1_]))
        nw8_[int(10)] = (0) - (d_46_v1_)
        nw8_[int(11)] = len(d_52_v7_)
        nw8_[int(12)] = d_46_v1_
        nw8_[int(13)] = len(d_52_v7_)
        nw8_[int(14)] = len(_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_55_i0_ in range(default__.abs(325))]))
        nw8_[int(15)] = d_46_v1_
        nw8_[int(16)] = d_46_v1_
        nw8_[int(17)] = d_46_v1_
        nw8_[int(18)] = d_46_v1_
        nw8_[int(19)] = d_46_v1_
        nw8_[int(20)] = len(d_53_v8_)
        nw8_[int(21)] = d_46_v1_
        nw8_[int(22)] = len(_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_56_i1_ in range(default__.abs(47))]))
        nw8_[int(23)] = d_46_v1_
        nw8_[int(24)] = d_46_v1_
        nw8_[int(25)] = (0) - (d_46_v1_)
        nw8_[int(26)] = d_46_v1_
        nw8_[int(27)] = d_46_v1_
        nw8_[int(28)] = d_46_v1_
        d_54_v9_ = nw8_
        d_57_v10_: _dafny.Map
        d_57_v10_ = _dafny.Map({d_45_v0_: d_54_v9_})
        d_58_v11_: _dafny.Set
        d_58_v11_ = _dafny.Set({d_46_v1_})
        d_59_v12_: _dafny.Seq
        d_59_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_60_i2_ in range(default__.abs(475))])])
        d_61_globalState_: GlobalState
        nw9_ = GlobalState()
        nw9_.ctor__(766, d_57_v10_, _dafny.SeqWithoutIsStrInference([d_58_v11_]), d_59_v12_, True)
        d_61_globalState_ = nw9_
        d_62_v13_: _dafny.Seq
        d_63_v14_: _dafny.Array
        d_64_v15_: int
        d_65_v16_: int
        out0_: _dafny.Seq
        out1_: _dafny.Array
        out2_: int
        out3_: int
        out0_, out1_, out2_, out3_ = default__.m0((861) + (len(_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_66_i3_ in range(default__.abs(491))]))), d_46_v1_, d_61_globalState_)
        d_62_v13_ = out0_
        d_63_v14_ = out1_
        d_64_v15_ = out2_
        d_65_v16_ = out3_
        d_67_v17_: _dafny.MultiSet
        d_67_v17_ = _dafny.MultiSet([d_65_v16_, d_64_v15_, d_64_v15_])
        d_65_v16_ = default__.safeModuloInt(d_65_v16_, ((d_67_v17_).intersection(d_67_v17_)).cardinality)
        d_68_v18_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.MultiSet({}), 6)
        d_68_v18_ = nw10_
        index4_ = default__.safeIndex(816, (d_68_v18_).length(0))
        (d_68_v18_)[index4_] = d_67_v17_
        d_69_v19_: _dafny.Seq
        d_69_v19_ = _dafny.SeqWithoutIsStrInference([d_67_v17_])
        d_70_v20_: _dafny.Seq
        d_70_v20_ = _dafny.SeqWithoutIsStrInference([720])
        pat_let_tv3_ = d_52_v7_
        pat_let_tv4_ = d_49_v4_
        index5_ = default__.safeIndex(816, (d_68_v18_).length(0))
        def iife4_(_pat_let0_0):
            def iife5_(d_71_dt__update__tmp_h0_):
                def iife6_(_pat_let1_0):
                    def iife7_(d_72_dt__update_hcf1_h0_):
                        def iife8_(_pat_let2_0):
                            def iife9_(d_73_dt__update_hcf3_h0_):
                                return D0_DC1(d_72_dt__update_hcf1_h0_, (d_71_dt__update__tmp_h0_).cf2, d_73_dt__update_hcf3_h0_, (d_71_dt__update__tmp_h0_).cf4)
                            return iife9_(_pat_let2_0)
                        return iife8_(pat_let_tv4_)
                    return iife7_(_pat_let1_0)
                return iife6_(pat_let_tv3_)
            return iife5_(_pat_let0_0)
        (d_68_v18_)[index5_] = ((d_67_v17_).intersection((d_69_v19_)[default__.safeIndex(len(d_62_v13_), len(d_69_v19_))])).intersection((_dafny.MultiSet(d_70_v20_)) | (_dafny.MultiSet((iife4_(D0_DC1(d_62_v13_, d_64_v15_, d_49_v4_, (_dafny.SeqWithoutIsStrInference([len(d_58_v11_), d_46_v1_])).set(default__.safeIndex(d_64_v15_, len(_dafny.SeqWithoutIsStrInference([len(d_58_v11_), d_46_v1_]))), d_64_v15_)))).cf4)))
        d_74_v21_: _dafny.Map
        d_74_v21_ = _dafny.Map({d_65_v16_: default__.safeModuloInt(125, d_65_v16_)})
        d_75_v22_: D0
        d_75_v22_ = D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vewema")), d_46_v1_, d_49_v4_, d_70_v20_)
        d_74_v21_ = (d_74_v21_).set((d_64_v15_) - ((d_75_v22_).cf2), (0) - ((741) * (d_65_v16_)))
        d_64_v15_ = d_65_v16_
        d_76_v23_: _dafny.Array
        def lambda5_(d_77_v5_):
            def lambda6_(d_78_i4_):
                return d_77_v5_

            return lambda6_

        init2_ = lambda5_(d_50_v5_)
        nw11_ = _dafny.Array(None, 27)
        for i0_2_ in range(nw11_.length(0)):
            nw11_[i0_2_] = init2_(i0_2_)
        d_76_v23_ = nw11_
        index6_ = default__.safeIndex(758, (d_76_v23_).length(0))
        (d_76_v23_)[index6_] = (d_50_v5_) | (d_50_v5_)
        index7_ = default__.safeIndex(758, (d_76_v23_).length(0))
        (d_76_v23_)[index7_] = (d_50_v5_) | ((d_50_v5_) | (d_50_v5_))
        d_47_v2_ = _dafny.CodePoint('y')
        d_79_v24_: _dafny.Array
        nw12_ = _dafny.Array(False, 2)
        d_79_v24_ = nw12_
        d_80_v25_: _dafny.Map
        d_80_v25_ = _dafny.Map({d_79_v24_: not(d_45_v0_)})
        d_80_v25_ = (d_80_v25_).set(d_79_v24_, not(not (d_45_v0_) or (d_45_v0_)))
        if (len(d_58_v11_)) >= (d_46_v1_):
            d_81_v26_: _dafny.Map
            d_81_v26_ = _dafny.Map({d_45_v0_: d_70_v20_})
            d_82_v27_: D0
            d_82_v27_ = D0_DC2(d_81_v26_, not(d_45_v0_))
            source3_ = d_82_v27_
            if source3_.is_DC1:
                d_83___mcc_h0_ = source3_.cf1
                d_84___mcc_h1_ = source3_.cf2
                d_85___mcc_h2_ = source3_.cf3
                d_86___mcc_h3_ = source3_.cf4
                d_87_cf4_ = d_86___mcc_h3_
                d_88_cf3_ = d_85___mcc_h2_
                d_89_cf2_ = d_84___mcc_h1_
                d_90_cf1_ = d_83___mcc_h0_
                d_45_v0_ = ((((_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_91_i5_ in range(default__.abs(-964))])).set(default__.safeIndex(d_65_v16_, len(_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_92_i5_ in range(default__.abs(-964))]))), d_47_v2_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))).set(default__.safeIndex(d_64_v15_, len(((_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_93_i5_ in range(default__.abs(-964))])).set(default__.safeIndex(d_65_v16_, len(_dafny.SeqWithoutIsStrInference([d_47_v2_ for d_94_i5_ in range(default__.abs(-964))]))), d_47_v2_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))), d_47_v2_)) == (d_52_v7_)
                d_95_v28_: C0
                nw13_ = C0()
                nw13_.ctor__(d_89_cf2_, (0) - (d_46_v1_))
                d_95_v28_ = nw13_
                d_89_cf2_ = default__.safeDivisionInt((d_95_v28_).f13, d_89_cf2_)
                d_64_v15_ = (d_65_v16_) - ((d_95_v28_).f13)
            elif source3_.is_DC2:
                d_96___mcc_h4_ = source3_.cf5
                d_97___mcc_h5_ = source3_.cf6
                d_98_cf6_ = d_97___mcc_h5_
                d_99_cf5_ = d_96___mcc_h4_
                d_64_v15_ = d_65_v16_
                index8_ = default__.safeIndex(847, (d_79_v24_).length(0))
                (d_79_v24_)[index8_] = d_98_cf6_
                index9_ = default__.safeIndex(116, (d_54_v9_).length(0))
                (d_54_v9_)[index9_] = d_64_v15_
                d_100_v29_: C0
                nw14_ = C0()
                nw14_.ctor__(d_46_v1_, d_64_v15_)
                d_100_v29_ = nw14_
                d_101_v30_: _dafny.Map
                d_101_v30_ = _dafny.Map({d_100_v29_: True})
                d_102_v31_: _dafny.Map
                d_102_v31_ = _dafny.Map({len(d_101_v30_): d_49_v4_})
                index10_ = default__.safeIndex(187, (d_79_v24_).length(0))
                (d_79_v24_)[index10_] = (d_64_v15_) > (len(default__.fm5(d_49_v4_, d_102_v31_, d_61_globalState_)))
                d_103_v32_: _dafny.Set
                d_103_v32_ = _dafny.Set({True})
                index11_ = default__.safeIndex(847, (d_79_v24_).length(0))
                index12_ = default__.safeIndex(816, (d_68_v18_).length(0))
                index13_ = default__.safeIndex(116, (d_54_v9_).length(0))
                index14_ = default__.safeIndex(187, (d_79_v24_).length(0))
                rhs3_ = (d_45_v0_) == (False)
                rhs4_ = (d_67_v17_) | ((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))])
                rhs5_ = (d_100_v29_).f12
                rhs6_ = d_79_v24_
                rhs7_ = not(((d_103_v32_).issubset(d_103_v32_) if False else d_45_v0_))
                lhs2_ = d_79_v24_
                lhs3_ = default__.safeIndex(847, (d_79_v24_).length(0))
                lhs4_ = d_68_v18_
                lhs5_ = default__.safeIndex(816, (d_68_v18_).length(0))
                lhs6_ = d_54_v9_
                lhs7_ = default__.safeIndex(116, (d_54_v9_).length(0))
                lhs8_ = d_79_v24_
                lhs9_ = default__.safeIndex(187, (d_79_v24_).length(0))
                lhs2_[lhs3_] = rhs3_
                lhs4_[lhs5_] = rhs4_
                lhs6_[lhs7_] = rhs5_
                d_79_v24_ = rhs6_
                lhs8_[lhs9_] = rhs7_
                d_47_v2_ = default__.fm11(d_47_v2_, (not((d_79_v24_)[default__.safeIndex(847, (d_79_v24_).length(0))])) or ((d_79_v24_)[default__.safeIndex(847, (d_79_v24_).length(0))]), (d_100_v29_).fm8((d_79_v24_)[default__.safeIndex(847, (d_79_v24_).length(0))], d_52_v7_, d_47_v2_, d_61_globalState_), d_61_globalState_)
                d_104_v33_: T0
                nw15_ = C4()
                nw15_.ctor__((d_100_v29_).f12, (d_54_v9_)[default__.safeIndex(116, (d_54_v9_).length(0))], d_98_cf6_)
                d_104_v33_ = nw15_
                d_104_v33_ = d_104_v33_
            elif True:
                d_105___mcc_h6_ = source3_.cf0
                d_106_cf0_ = d_105___mcc_h6_
                d_107_v34_: _dafny.Seq
                d_108_v35_: _dafny.Array
                d_109_v36_: int
                d_110_v37_: int
                out4_: _dafny.Seq
                out5_: _dafny.Array
                out6_: int
                out7_: int
                out4_, out5_, out6_, out7_ = default__.m0(d_64_v15_, d_46_v1_, d_61_globalState_)
                d_107_v34_ = out4_
                d_108_v35_ = out5_
                d_109_v36_ = out6_
                d_110_v37_ = out7_
                d_111_v38_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_111_v38_ = nw16_
                index15_ = default__.safeIndex(373, (d_111_v38_).length(0))
                (d_111_v38_)[index15_] = d_62_v13_
                index16_ = default__.safeIndex(373, (d_111_v38_).length(0))
                (d_111_v38_)[index16_] = d_62_v13_
                d_112_v39_: _dafny.Array
                def lambda7_(d_113_i6_):
                    return _dafny.CodePoint('u')

                init3_ = lambda7_
                nw17_ = _dafny.Array(None, 20)
                for i0_3_ in range(nw17_.length(0)):
                    nw17_[i0_3_] = init3_(i0_3_)
                d_112_v39_ = nw17_
                index17_ = default__.safeIndex(60, (d_112_v39_).length(0))
                (d_112_v39_)[index17_] = d_47_v2_
                index18_ = default__.safeIndex(60, (d_112_v39_).length(0))
                (d_112_v39_)[index18_] = d_47_v2_
                d_64_v15_ = d_46_v1_
            d_114_v40_: str
            d_114_v40_ = d_47_v2_
            d_114_v40_ = d_47_v2_
            nw18_ = _dafny.Array(int(0), 20)
            d_63_v14_ = nw18_
            d_47_v2_ = default__.fm11(d_47_v2_, d_45_v0_, not (d_45_v0_) or (default__.fm24(d_61_globalState_)), d_61_globalState_)
            d_115_v41_: T0
            nw19_ = C1()
            nw19_.ctor__(d_64_v15_, d_64_v15_, d_64_v15_, d_45_v0_)
            d_115_v41_ = nw19_
            d_116_v42_: _dafny.Array
            nw20_ = _dafny.Array(None, 3)
            nw20_[int(0)] = (d_115_v41_ if d_45_v0_ else d_115_v41_)
            nw20_[int(1)] = d_115_v41_
            nw20_[int(2)] = d_115_v41_
            d_116_v42_ = nw20_
            index19_ = default__.safeIndex(472, (d_116_v42_).length(0))
            (d_116_v42_)[index19_] = d_115_v41_
            index20_ = default__.safeIndex(472, (d_116_v42_).length(0))
            (d_116_v42_)[index20_] = (d_115_v41_ if (d_115_v41_).f6 else d_115_v41_)
        elif True:
            d_117_v43_: C3
            nw21_ = C3()
            nw21_.ctor__(d_65_v16_, ((0) - (-469)) - (d_46_v1_), d_45_v0_)
            d_117_v43_ = nw21_
            d_65_v16_ = (d_117_v43_).f8
            d_118_v44_: _dafny.Set
            d_118_v44_ = _dafny.Set({d_45_v0_})
            d_119_v45_: _dafny.Map
            d_119_v45_ = _dafny.Map({d_118_v44_: d_52_v7_})
            d_120_v46_: T0
            nw22_ = C1()
            nw22_.ctor__((d_117_v43_).f8, 431, d_64_v15_, d_45_v0_)
            d_120_v46_ = nw22_
            d_121_v47_: _dafny.Set
            d_121_v47_ = _dafny.Set({d_120_v46_, d_120_v46_})
            d_45_v0_ = (d_47_v2_) not in ((((d_119_v45_)[d_118_v44_] if (d_118_v44_) in (d_119_v45_) else d_62_v13_)).set(default__.safeIndex(len(d_121_v47_), len(((d_119_v45_)[d_118_v44_] if (d_118_v44_) in (d_119_v45_) else d_62_v13_))), d_47_v2_))
            rhs8_ = (True if d_45_v0_ else (d_120_v46_).f6)
            rhs9_ = ((d_64_v15_) + (d_46_v1_)) * ((0) - ((len(d_53_v8_)) - (len(d_53_v8_))))
            d_45_v0_ = rhs8_
            d_65_v16_ = rhs9_
            rhs10_ = d_47_v2_
            rhs11_ = (0) - (d_65_v16_)
            d_47_v2_ = rhs10_
            d_65_v16_ = rhs11_
        d_46_v1_ = default__.safeModuloInt((d_70_v20_)[default__.safeIndex(d_64_v15_, len(d_70_v20_))], d_64_v15_)
        if default__.fm24(d_61_globalState_):
            d_45_v0_ = d_45_v0_
            d_122_v49_: D1
            def iife10_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in (_dafny.Set({d_64_v15_})).Elements:
                    d_123_v48_: int = compr_4_
                    if (d_123_v48_) in (_dafny.Set({d_64_v15_})):
                        coll4_[default__.safeModuloInt(d_123_v48_, 236)] = d_45_v0_
                return _dafny.Map(coll4_)
            d_122_v49_ = D1_DC3(iife10_()
)
            d_124_v50_: C3
            nw23_ = C3()
            nw23_.ctor__(-302, len(d_70_v20_), d_45_v0_)
            d_124_v50_ = nw23_
            d_125_v51_: T0
            nw24_ = C1()
            nw24_.ctor__(d_65_v16_, d_64_v15_, len(_dafny.Map({d_122_v49_: d_124_v50_})), d_45_v0_)
            d_125_v51_ = nw24_
            d_126_v52_: D5
            d_126_v52_ = D5_DC16(d_125_v51_, (d_125_v51_).f6, d_65_v16_)
            d_127_v53_: C0
            nw25_ = C0()
            nw25_.ctor__((default__.fm6((_dafny.MultiSet([d_45_v0_])).cardinality, d_46_v1_, d_45_v0_, d_46_v1_, d_61_globalState_)) + (default__.fm6(len(d_70_v20_), d_65_v16_, d_45_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puoy"))), d_61_globalState_)), default__.safeModuloInt((0) - (d_65_v16_), (d_126_v52_).cf34))
            d_127_v53_ = nw25_
            d_47_v2_ = default__.fm11(d_47_v2_, (d_45_v0_) and ((d_125_v51_).f6), (d_125_v51_).f6, d_61_globalState_)
            d_47_v2_ = d_47_v2_
            index21_ = default__.safeIndex(658, (d_54_v9_).length(0))
            (d_54_v9_)[index21_] = ((d_127_v53_).f13) - ((d_127_v53_).f13)
            index22_ = default__.safeIndex(658, (d_54_v9_).length(0))
            rhs12_ = d_62_v13_
            rhs13_ = ((d_127_v53_).f12) + (d_64_v15_)
            rhs14_ = d_46_v1_
            lhs10_ = d_54_v9_
            lhs11_ = default__.safeIndex(658, (d_54_v9_).length(0))
            d_62_v13_ = rhs12_
            d_46_v1_ = rhs13_
            lhs10_[lhs11_] = rhs14_
        elif True:
            d_128_v54_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.Seq({}), 4)
            d_128_v54_ = nw26_
            d_128_v54_ = d_128_v54_
            d_79_v24_ = d_79_v24_
            d_45_v0_ = d_45_v0_
            d_129_v55_: _dafny.Map
            d_129_v55_ = _dafny.Map({True: d_70_v20_})
            pat_let_tv5_ = d_45_v0_
            def iife11_(_pat_let3_0):
                def iife12_(d_130_dt__update__tmp_h1_):
                    def iife13_(_pat_let4_0):
                        def iife14_(d_131_dt__update_hcf28_h0_):
                            return D3_DC12(d_131_dt__update_hcf28_h0_, (d_130_dt__update__tmp_h1_).cf29)
                        return iife14_(_pat_let4_0)
                    return iife13_(pat_let_tv5_)
                return iife12_(_pat_let3_0)
            source4_ = iife11_(D3_DC12((D0_DC2(d_129_v55_, default__.fm24(d_61_globalState_))).cf6, (d_67_v17_).cardinality))
            if source4_.is_DC12:
                d_132___mcc_h7_ = source4_.cf28
                d_133___mcc_h8_ = source4_.cf29
                d_134_cf29_ = d_133___mcc_h8_
                d_135_cf28_ = d_132___mcc_h7_
                d_136_v56_: _dafny.Map
                d_136_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skjvslqs")): d_64_v15_})
                d_137_v57_: _dafny.Map
                d_137_v57_ = _dafny.Map({d_136_v56_: d_79_v24_})
                d_138_v58_: _dafny.Seq
                d_138_v58_ = _dafny.SeqWithoutIsStrInference([d_79_v24_])
                d_137_v57_ = (d_137_v57_).set(d_136_v56_, (d_138_v58_)[default__.safeIndex(d_134_cf29_, len(d_138_v58_))])
                d_139_v59_: _dafny.Seq
                d_140_v60_: _dafny.Array
                d_141_v61_: int
                d_142_v62_: int
                out8_: _dafny.Seq
                out9_: _dafny.Array
                out10_: int
                out11_: int
                out8_, out9_, out10_, out11_ = default__.m0(d_64_v15_, d_46_v1_, d_61_globalState_)
                d_139_v59_ = out8_
                d_140_v60_ = out9_
                d_141_v61_ = out10_
                d_142_v62_ = out11_
                d_143_v63_: C0
                nw27_ = C0()
                nw27_.ctor__(d_46_v1_, d_65_v16_)
                d_143_v63_ = nw27_
                d_144_v64_: _dafny.Map
                d_144_v64_ = _dafny.Map({d_143_v63_: ((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))]).cardinality})
                d_145_v65_: _dafny.MultiSet
                d_145_v65_ = _dafny.MultiSet([d_144_v64_])
                d_51_v6_ = ((d_51_v6_).set(d_134_cf29_, d_45_v0_) if default__.fm24(d_61_globalState_) else _dafny.Map({((d_145_v65_)[d_144_v64_] if (d_144_v64_) in (d_145_v65_) else d_141_v61_): d_45_v0_}))
                index23_ = default__.safeIndex(831, (d_54_v9_).length(0))
                (d_54_v9_)[index23_] = d_134_cf29_
                d_146_v66_: _dafny.Set
                d_146_v66_ = _dafny.Set({d_45_v0_})
                index24_ = default__.safeIndex(831, (d_54_v9_).length(0))
                (d_54_v9_)[index24_] = len(d_146_v66_)
            elif True:
                d_147___mcc_h9_ = source4_.cf27
                d_148_cf27_ = d_147___mcc_h9_
                d_149_v67_: _dafny.Map
                d_149_v67_ = _dafny.Map({(d_58_v11_).intersection(d_58_v11_): not(default__.fm24(d_61_globalState_))})
                d_149_v67_ = (d_149_v67_).set(d_58_v11_, d_45_v0_)
                d_46_v1_ = len(((d_75_v22_).cf3) | (default__.fm13(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptcfinu")), d_65_v16_, (d_69_v19_)[default__.safeIndex(d_64_v15_, len(d_69_v19_))], d_61_globalState_)))
                d_150_v68_: _dafny.Map
                d_150_v68_ = _dafny.Map({d_46_v1_: _dafny.Map({d_45_v0_: d_65_v16_})})
                d_151_v69_: C0
                nw28_ = C0()
                nw28_.ctor__((d_64_v15_) + (722), len(default__.fm5(d_49_v4_, d_150_v68_, d_61_globalState_)))
                d_151_v69_ = nw28_
                d_152_v70_: T0
                nw29_ = C3()
                nw29_.ctor__((d_151_v69_).f12, d_65_v16_, d_45_v0_)
                d_152_v70_ = nw29_
                rhs15_ = d_151_v69_
                rhs16_ = d_152_v70_
                d_151_v69_ = rhs15_
                d_152_v70_ = rhs16_
                d_153_v71_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_153_v71_ = nw30_
                index25_ = default__.safeIndex(314, (d_153_v71_).length(0))
                (d_153_v71_)[index25_] = (d_52_v7_) + (d_52_v7_)
                index26_ = default__.safeIndex(314, (d_153_v71_).length(0))
                (d_153_v71_)[index26_] = d_52_v7_
            d_154_v72_: _dafny.Seq
            d_155_v73_: _dafny.Array
            d_156_v74_: int
            d_157_v75_: int
            out12_: _dafny.Seq
            out13_: _dafny.Array
            out14_: int
            out15_: int
            out12_, out13_, out14_, out15_ = default__.m0(default__.safeModuloInt(d_46_v1_, d_64_v15_), d_65_v16_, d_61_globalState_)
            d_154_v72_ = out12_
            d_155_v73_ = out13_
            d_156_v74_ = out14_
            d_157_v75_ = out15_
        d_54_v9_ = d_63_v14_
        d_158_v76_: C4
        nw31_ = C4()
        nw31_.ctor__(d_65_v16_, d_64_v15_, d_45_v0_)
        d_158_v76_ = nw31_
        d_159_v77_: _dafny.Seq
        d_159_v77_ = _dafny.SeqWithoutIsStrInference([d_158_v76_])
        d_160_v78_: D9
        d_160_v78_ = D9_DC22(d_159_v77_)
        d_161_v79_: _dafny.MultiSet
        d_161_v79_ = _dafny.MultiSet([d_158_v76_])
        d_162_v81_: D2
        def iife15_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (d_52_v7_).Elements:
                d_163_v80_: str = compr_5_
                if (d_163_v80_) in (d_52_v7_):
                    coll5_[d_163_v80_] = not(d_45_v0_)
            return _dafny.Map(coll5_)
        d_162_v81_ = D2_DC10(D2_DC8((d_67_v17_).set(d_46_v1_, default__.abs((0) - (len(iife15_()
))))))
        d_164_v82_: D2
        d_164_v82_ = D2_DC8((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))])
        pat_let_tv6_ = d_159_v77_
        def iife16_(_pat_let5_0):
            def iife17_(d_165_dt__update__tmp_h2_):
                def iife18_(_pat_let6_0):
                    def iife19_(d_166_dt__update_hcf42_h0_):
                        return D9_DC22(d_166_dt__update_hcf42_h0_)
                    return iife19_(_pat_let6_0)
                return iife18_(pat_let_tv6_)
            return iife17_(_pat_let5_0)
        source5_ = (d_162_v81_ if (d_161_v79_).ispropersubset(_dafny.MultiSet((iife16_(d_160_v78_)).cf42)) else D2_DC10(d_164_v82_))
        if source5_.is_DC9:
            d_167___mcc_h10_ = source5_.cf23
            d_168___mcc_h11_ = source5_.cf24
            d_169___mcc_h12_ = source5_.cf25
            d_170_cf25_ = d_169___mcc_h12_
            d_171_cf24_ = d_168___mcc_h11_
            d_172_cf23_ = d_167___mcc_h10_
            d_46_v1_ = d_64_v15_
            d_62_v13_ = _dafny.SeqWithoutIsStrInference([(d_52_v7_)[default__.safeIndex(d_158_v76_.f7, len(d_52_v7_))] for d_173_i7_ in range(default__.abs(-45))])
            d_172_cf23_ = d_172_cf23_
            d_174_v83_: C0
            nw32_ = C0()
            nw32_.ctor__(d_46_v1_, d_64_v15_)
            d_174_v83_ = nw32_
            d_175_v84_: _dafny.Map
            d_175_v84_ = _dafny.Map({d_45_v0_: d_174_v83_})
            d_176_v85_: C2
            nw33_ = C2()
            nw33_.ctor__(len(d_175_v84_))
            d_176_v85_ = nw33_
        elif source5_.is_DC8:
            d_177___mcc_h13_ = source5_.cf22
            d_178_cf22_ = d_177___mcc_h13_
            index27_ = default__.safeIndex(5, (d_63_v14_).length(0))
            (d_63_v14_)[index27_] = (d_158_v76_.f7) + ((0) - (d_158_v76_.f7))
            index28_ = default__.safeIndex(5, (d_63_v14_).length(0))
            (d_63_v14_)[index28_] = d_158_v76_.f7
            (d_158_v76_).m1(d_158_v76_.f7, d_61_globalState_)
            d_179_v86_: _dafny.Array
            nw34_ = _dafny.Array(_dafny.Seq({}), 21)
            d_179_v86_ = nw34_
            index29_ = default__.safeIndex(224, (d_179_v86_).length(0))
            (d_179_v86_)[index29_] = (d_53_v8_).set(default__.safeIndex(len(d_53_v8_), len(d_53_v8_)), d_45_v0_)
            index30_ = default__.safeIndex(224, (d_179_v86_).length(0))
            (d_179_v86_)[index30_] = ((_dafny.SeqWithoutIsStrInference([d_45_v0_, True]) if d_45_v0_ else d_53_v8_)) + (d_53_v8_)
            d_180_v87_: C0
            nw35_ = C0()
            nw35_.ctor__(d_46_v1_, len(d_53_v8_))
            d_180_v87_ = nw35_
            d_181_v88_: _dafny.Map
            d_181_v88_ = _dafny.Map({_dafny.Set({d_180_v87_}): (d_180_v87_).f13})
            d_182_v89_: _dafny.Set
            d_182_v89_ = _dafny.Set({d_180_v87_})
            d_183_v90_: _dafny.Map
            d_183_v90_ = _dafny.Map({((d_181_v88_).set(d_182_v89_, len(d_74_v21_))).set(d_182_v89_, len(d_58_v11_)): not(not(d_45_v0_))})
            d_183_v90_ = (d_183_v90_).set(d_181_v88_, not((d_180_v87_).fm8(d_45_v0_, (d_75_v22_).cf1, d_47_v2_, d_61_globalState_)))
        elif True:
            d_184___mcc_h14_ = source5_.cf26
            d_185_cf26_ = d_184___mcc_h14_
            d_64_v15_ = default__.fm6(d_64_v15_, default__.safeModuloInt(d_158_v76_.f7, d_158_v76_.f7), not((d_158_v76_).fm0(False, d_61_globalState_)), ((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))]).cardinality, d_61_globalState_)
            d_186_v91_: _dafny.Array
            nw36_ = _dafny.Array(None, 4)
            nw36_[int(0)] = d_158_v76_
            nw36_[int(1)] = d_158_v76_
            nw36_[int(2)] = d_158_v76_
            nw36_[int(3)] = d_158_v76_
            d_186_v91_ = nw36_
            d_187_v92_: D3
            d_187_v92_ = D3_DC11(d_63_v14_)
            d_188_v93_: _dafny.Map
            d_188_v93_ = _dafny.Map({d_186_v91_: d_187_v92_})
            d_188_v93_ = (d_188_v93_).set(d_186_v91_, d_187_v92_)
            d_74_v21_ = (d_74_v21_).set(d_64_v15_, (d_158_v76_.f7 if True else d_64_v15_))
            d_189_v94_: _dafny.Map
            d_189_v94_ = _dafny.Map({d_45_v0_: d_79_v24_})
            d_190_v95_: T0
            nw37_ = C3()
            nw37_.ctor__(d_65_v16_, d_46_v1_, d_45_v0_)
            d_190_v95_ = nw37_
            d_191_v96_: _dafny.Set
            d_191_v96_ = _dafny.Set({d_190_v95_})
            d_192_v97_: _dafny.Map
            d_192_v97_ = _dafny.Map({d_189_v94_: d_191_v96_})
            d_192_v97_ = (d_192_v97_).set(d_189_v94_, d_191_v96_)
        d_193_i8_: int
        d_193_i8_ = 0
        with _dafny.label("0"):
            while (d_158_v76_).fm0((d_46_v1_) <= (383), d_61_globalState_):
                with _dafny.c_label("0"):
                    if (d_193_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_193_i8_ = (d_193_i8_) + (1)
                    d_194_v98_: _dafny.Set
                    d_194_v98_ = _dafny.Set({d_79_v24_, d_79_v24_})
                    d_195_v99_: _dafny.Seq
                    d_196_v100_: _dafny.Array
                    d_197_v101_: int
                    d_198_v102_: int
                    out16_: _dafny.Seq
                    out17_: _dafny.Array
                    out18_: int
                    out19_: int
                    out16_, out17_, out18_, out19_ = default__.m0(default__.safeModuloInt(len(d_194_v98_), d_158_v76_.f7), (d_64_v15_ if d_45_v0_ else d_158_v76_.f7), d_61_globalState_)
                    d_195_v99_ = out16_
                    d_196_v100_ = out17_
                    d_197_v101_ = out18_
                    d_198_v102_ = out19_
                    d_45_v0_ = (default__.fm24(d_61_globalState_) if (d_195_v99_) < (d_62_v13_) else d_45_v0_)
                    d_199_v103_: C0
                    nw38_ = C0()
                    nw38_.ctor__(d_197_v101_, d_64_v15_)
                    d_199_v103_ = nw38_
                    d_200_v104_: C1
                    nw39_ = C1()
                    nw39_.ctor__(d_64_v15_, default__.fm6(d_65_v16_, d_198_v102_, d_45_v0_, d_197_v101_, d_61_globalState_), (d_199_v103_).f12, d_45_v0_)
                    d_200_v104_ = nw39_
                    d_201_v105_: D4
                    d_201_v105_ = D4_DC14()
                    d_202_v106_: _dafny.Map
                    d_202_v106_ = _dafny.Map({d_200_v104_: d_201_v105_})
                    (d_158_v76_).f7 = default__.fm6(default__.safeDivisionInt(148, d_158_v76_.f7), len(d_202_v106_), (d_65_v16_) != (d_158_v76_.f7), (d_65_v16_) * (d_46_v1_), d_61_globalState_)
                    pass
            pass
        d_203_v107_: _dafny.Seq
        d_204_v108_: _dafny.Array
        d_205_v109_: int
        d_206_v110_: int
        out20_: _dafny.Seq
        out21_: _dafny.Array
        out22_: int
        out23_: int
        out20_, out21_, out22_, out23_ = default__.m0((((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))])[d_65_v16_] if (d_65_v16_) in ((d_68_v18_)[default__.safeIndex(816, (d_68_v18_).length(0))]) else d_65_v16_), d_64_v15_, d_61_globalState_)
        d_203_v107_ = out20_
        d_204_v108_ = out21_
        d_205_v109_ = out22_
        d_206_v110_ = out23_
        d_207_v111_: _dafny.MultiSet
        d_207_v111_ = _dafny.MultiSet([d_45_v0_, not(d_45_v0_)])
        d_45_v0_ = ((d_45_v0_) not in (d_207_v111_)) == (d_45_v0_)
        _dafny.print(_dafny.string_of(d_45_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_46_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_47_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_48_v3_) == (_dafny.MultiSet([_dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_49_v4_) == (_dafny.Map({False: 82}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_50_v5_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_51_v6_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_52_v7_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_53_v8_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_54_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_57_v10_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_58_v11_) == (_dafny.Set({82}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_59_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_61_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_61_globalState_).f1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_61_globalState_).f2) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({82})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_61_globalState_).f3) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d'), _dafny.CodePoint('d')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_61_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_62_v13_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_63_v14_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_64_v15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_65_v16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_67_v17_) == (_dafny.MultiSet([1, 1352, 1352]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_68_v18_)[0]) == (_dafny.MultiSet([1352]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_69_v19_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([1, 1352, 1352])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_70_v20_) == (_dafny.SeqWithoutIsStrInference([720]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_74_v21_) == (_dafny.Map({1: 0, 1270: -741, 4: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_75_v22_).cf1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_75_v22_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_v22_).cf3) == (_dafny.Map({False: 82}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_75_v22_).cf4) == (_dafny.SeqWithoutIsStrInference([720]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[0]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[1]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[2]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[3]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[4]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[5]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[6]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[7]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[8]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[9]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[10]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[11]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[12]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[13]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[14]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[15]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[16]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[17]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[18]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[19]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[20]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[21]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[22]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[23]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[24]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[25]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_76_v23_)[26]) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_80_v25_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_158_v76_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_159_v77_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_160_v78_).cf42)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v79_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_162_v81_).cf26).cf22) == (_dafny.MultiSet([1, 1352, 1352, 0, 0, 0, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_164_v82_).cf22) == (_dafny.MultiSet([1352]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_193_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_203_v107_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v108_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_205_v109_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_206_v110_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_207_v111_) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.Map({}), _dafny.Seq({}))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({self.cf1.VerbatimString(True)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
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
        return lambda: D1_DC4(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D1_DC7)

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {self.cf10.VerbatimString(True)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC7(D1, NamedTuple('DC7', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC7({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC7) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC9(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D2_DC10)

class D2_DC9(D2, NamedTuple('DC9', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC10(D2, NamedTuple('DC10', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC10({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC10) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC12(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D3_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC12(D3, NamedTuple('DC12', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC12({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC12) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC14()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC14(D4, NamedTuple('DC14', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(None, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)

class D6_DC18(D6, NamedTuple('DC18', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC20(D7, NamedTuple('DC20', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)

class D8_DC21(D8, NamedTuple('DC21', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(int(0), int(0), False, _dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC23(D9, NamedTuple('DC23', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC26(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)

class D10_DC26(D10, NamedTuple('DC26', [('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self._f0: int = int(0)
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.Seq({})
        self._f3: _dafny.Seq = _dafny.Seq({})
        self._f4: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4

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

class C0:
    def  __init__(self):
        self._f12: int = int(0)
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f12, f13):
        (self)._f12 = f12
        (self)._f13 = f13

    def fm7(self, p0, p1, p2, p3, globalState):
        def iife20_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(925, 394):
                d_208_v0_: int = compr_6_
                if ((925) <= (d_208_v0_)) and ((d_208_v0_) < (394)):
                    coll6_[(d_208_v0_) * ((self).f12)] = True
            return _dafny.Map(coll6_)
        def iife21_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(397, 116):
                d_209_v1_: int = compr_7_
                if ((397) <= (d_209_v1_)) and ((d_209_v1_) < (116)):
                    coll7_[(d_209_v1_) + ((self).f13)] = False
            return _dafny.Map(coll7_)
        return (iife20_()
        ) | ((_dafny.Map({(self).f12: True})) | ((D1_DC3(iife21_()
)).cf7))

    def fm8(self, p0, p1, p2, globalState):
        return (not(True)) or (False)

    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C1(T0):
    def  __init__(self):
        self._f5: int = int(0)
        self._f6: bool = False
        self.f10: int = int(0)
        self._f11: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f10, f11, f5, f6):
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f5 = f5
        (self)._f6 = f6

    def fm0(self, p0, globalState):
        return (self).f6

    def fm2(self, p0, p1, p2, globalState):
        return ((self).f11) == ((0) - (((self).f5) + ((self).f11)))

    def fm3(self, p0, p1, p2, globalState):
        return _dafny.Map({(self).f6: self.f10})

    def m1(self, p0, globalState):
        d_210_v0_: _dafny.Seq
        d_210_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju"))
        d_211_v1_: _dafny.MultiSet
        d_211_v1_ = _dafny.MultiSet([(d_210_v0_ if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcobgiw"))), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_212_i0_ in range(default__.abs(951))]), (d_210_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_213_i1_ in range(default__.abs(0))]))])
        (self).f10 = ((d_211_v1_)[d_210_v0_] if (d_210_v0_) in (d_211_v1_) else p0)
        d_214_v2_: _dafny.Array
        nw40_ = _dafny.Array(False, 15)
        d_214_v2_ = nw40_
        d_215_v3_: _dafny.Map
        d_215_v3_ = _dafny.Map({(self).f6: p0})
        d_216_v4_: _dafny.Map
        d_216_v4_ = _dafny.Map({len(d_215_v3_): self.f10})
        d_217_v5_: _dafny.Set
        d_217_v5_ = _dafny.Set({True, (self).f6})
        d_218_v6_: _dafny.Seq
        d_218_v6_ = _dafny.SeqWithoutIsStrInference([len(d_217_v5_)])
        d_219_v7_: _dafny.Map
        d_219_v7_ = _dafny.Map({(self).f6: d_218_v6_})
        d_220_v8_: D0
        d_220_v8_ = D0_DC2(d_219_v7_, not((self).f6))
        d_221_v9_: _dafny.Set
        d_221_v9_ = _dafny.Set({d_220_v8_})
        hi1_ = ((d_216_v4_)[len(d_210_v0_)] if (len(d_210_v0_)) in (d_216_v4_) else len(d_221_v9_))
        for d_222_i2_ in range(default__.safeDivisionInt((self).f5, len(default__.fm4(len(_dafny.SeqWithoutIsStrInference([d_214_v2_])), globalState))), hi1_):
            d_223_v10_: _dafny.Seq
            d_223_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f6: d_218_v6_})])
            d_224_v11_: _dafny.Map
            d_224_v11_ = _dafny.Map({(self).f5: False})
            d_215_v3_ = (d_215_v3_).set((D0_DC2((d_223_v10_)[default__.safeIndex(p0, len(d_223_v10_))], (self).f6)).cf6, len(d_224_v11_))
            d_225_v12_: _dafny.Map
            d_225_v12_ = _dafny.Map({d_222_i2_: d_210_v0_})
            d_226_v13_: str
            d_226_v13_ = _dafny.CodePoint('b')
            d_227_v14_: D0
            d_227_v14_ = D0_DC1((d_210_v0_).set(default__.safeIndex(self.f10, len(d_210_v0_)), d_226_v13_), (self).f5, d_215_v3_, d_218_v6_)
            d_228_v15_: _dafny.Seq
            d_228_v15_ = _dafny.SeqWithoutIsStrInference([d_210_v0_, d_210_v0_])
            d_229_v16_: _dafny.Map
            d_229_v16_ = _dafny.Map({652: d_215_v3_})
            d_230_v17_: _dafny.Array
            nw41_ = _dafny.Array(None, 24)
            nw41_[int(0)] = (d_210_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md")))
            nw41_[int(1)] = ((d_225_v12_)[p0] if (p0) in (d_225_v12_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_231_i3_ in range(default__.abs(70))]))
            nw41_[int(2)] = ((d_210_v0_).set(default__.safeIndex(p0, len(d_210_v0_)), d_226_v13_)) + (d_210_v0_)
            nw41_[int(3)] = ((d_227_v14_).cf1).set(default__.safeIndex((self).f11, len((d_227_v14_).cf1)), d_226_v13_)
            nw41_[int(4)] = d_210_v0_
            nw41_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
            nw41_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))
            nw41_[int(7)] = (d_228_v15_)[default__.safeIndex(self.f10, len(d_228_v15_))]
            nw41_[int(8)] = d_210_v0_
            nw41_[int(9)] = d_210_v0_
            nw41_[int(10)] = d_210_v0_
            nw41_[int(11)] = d_210_v0_
            nw41_[int(12)] = default__.fm5(d_215_v3_, d_229_v16_, globalState)
            nw41_[int(13)] = ((d_225_v12_)[(self).f11] if ((self).f11) in (d_225_v12_) else d_210_v0_)
            nw41_[int(14)] = d_210_v0_
            nw41_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chm"))) + (d_210_v0_)
            nw41_[int(16)] = d_210_v0_
            nw41_[int(17)] = d_210_v0_
            nw41_[int(18)] = (d_210_v0_).set(default__.safeIndex((self).f11, len(d_210_v0_)), d_226_v13_)
            nw41_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_210_v0_)
            nw41_[int(20)] = (D0_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylxgswe")), default__.fm6((self).f5, (self).f5, (self).f6, (self).f5, globalState), d_215_v3_, d_218_v6_)).cf1
            nw41_[int(21)] = (d_210_v0_) + (d_210_v0_)
            nw41_[int(22)] = (d_210_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "soi")))
            nw41_[int(23)] = d_210_v0_
            d_230_v17_ = nw41_
            index31_ = default__.safeIndex(567, (d_230_v17_).length(0))
            (d_230_v17_)[index31_] = d_210_v0_
            d_232_v18_: bool
            d_232_v18_ = False
            d_233_v19_: _dafny.Seq
            d_233_v19_ = _dafny.SeqWithoutIsStrInference([(self).fm0((self).fm0(True, globalState), globalState), d_232_v18_])
            d_234_v20_: _dafny.Seq
            d_234_v20_ = _dafny.SeqWithoutIsStrInference([d_233_v19_])
            index32_ = default__.safeIndex(567, (d_230_v17_).length(0))
            rhs17_ = (d_222_i2_) * ((self.f10) * (p0))
            rhs18_ = p0
            rhs19_ = d_210_v0_
            rhs20_ = (self.f10) < (len((d_234_v20_) + (d_234_v20_)))
            lhs12_ = self
            lhs13_ = self
            lhs14_ = d_230_v17_
            lhs15_ = default__.safeIndex(567, (d_230_v17_).length(0))
            lhs12_.f10 = rhs17_
            lhs13_.f10 = rhs18_
            lhs14_[lhs15_] = rhs19_
            d_232_v18_ = rhs20_
            (self).f10 = d_222_i2_
            (self).f10 = (0) - (((self).f11) - (p0))
        d_235_v22_: _dafny.Map
        def iife22_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(111, 211):
                d_236_v21_: int = compr_8_
                if ((111) <= (d_236_v21_)) and ((d_236_v21_) < (211)):
                    coll8_[(d_236_v21_) * ((self).f11)] = d_215_v3_
            return _dafny.Map(coll8_)
        d_235_v22_ = _dafny.Map({default__.fm5(d_215_v3_, iife22_()
        , globalState): (self).f11})
        (self).f10 = len(d_235_v22_)
        hi2_ = self.f10
        for d_237_i4_ in range((self).f5, hi2_):
            d_238_v23_: C0
            nw42_ = C0()
            nw42_.ctor__(default__.fm6((0) - (len(d_210_v0_)), d_237_i4_, (self).f6, (self).f5, globalState), (self).f11)
            d_238_v23_ = nw42_
            if ((self).f6 if (self).f6 else ((self).f6) == ((self).f6)):
                d_239_v24_: _dafny.Seq
                d_240_v25_: _dafny.Array
                d_241_v26_: int
                d_242_v27_: int
                out24_: _dafny.Seq
                out25_: _dafny.Array
                out26_: int
                out27_: int
                out24_, out25_, out26_, out27_ = default__.m0(self.f10, (d_218_v6_)[default__.safeIndex((d_238_v23_).f13, len(d_218_v6_))], globalState)
                d_239_v24_ = out24_
                d_240_v25_ = out25_
                d_241_v26_ = out26_
                d_242_v27_ = out27_
                d_243_v28_: _dafny.MultiSet
                d_243_v28_ = _dafny.MultiSet([(self).f6])
                d_244_v29_: _dafny.Map
                d_244_v29_ = _dafny.Map({d_243_v28_: (0) - ((d_238_v23_).f12)})
                d_245_v30_: _dafny.MultiSet
                d_245_v30_ = _dafny.MultiSet([(d_238_v23_).f12, p0, len(d_239_v24_), self.f10, 612])
                d_246_v31_: _dafny.Map
                d_246_v31_ = _dafny.Map({(self).f11: d_243_v28_})
                d_247_v32_: D2
                d_247_v32_ = D2_DC8(d_245_v30_)
                d_248_v33_: _dafny.Map
                d_248_v33_ = _dafny.Map({d_214_v2_: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))))})
                d_249_v34_: _dafny.Array
                nw43_ = _dafny.Array(None, 15)
                nw43_[int(0)] = d_245_v30_
                nw43_[int(1)] = (d_245_v30_) - (d_245_v30_)
                nw43_[int(2)] = _dafny.MultiSet(d_218_v6_)
                nw43_[int(3)] = d_245_v30_
                nw43_[int(4)] = d_245_v30_
                nw43_[int(5)] = d_245_v30_
                nw43_[int(6)] = d_245_v30_
                nw43_[int(7)] = (d_245_v30_) | (d_245_v30_)
                nw43_[int(8)] = d_245_v30_
                nw43_[int(9)] = (d_245_v30_) - (d_245_v30_)
                nw43_[int(10)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(d_238_v23_).f13 for d_250_i5_ in range(default__.abs(180))])) + ((default__.fm9((d_238_v23_).f13, globalState)).cf4))
                nw43_[int(11)] = _dafny.MultiSet(d_218_v6_)
                nw43_[int(12)] = d_245_v30_
                nw43_[int(13)] = _dafny.MultiSet((d_218_v6_) + (_dafny.SeqWithoutIsStrInference([(((d_246_v31_)[d_237_i4_] if (d_237_i4_) in (d_246_v31_) else d_243_v28_)).cardinality, (d_238_v23_).f13, len(d_210_v0_)])))
                nw43_[int(14)] = ((d_247_v32_).cf22).set(((d_248_v33_)[d_214_v2_] if (d_214_v2_) in (d_248_v33_) else len(d_215_v3_)), default__.abs(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_251_i6_ in range(default__.abs(457))]))))
                d_249_v34_ = nw43_
                index33_ = default__.safeIndex(573, (d_249_v34_).length(0))
                (d_249_v34_)[index33_] = d_245_v30_
                d_252_v35_: bool
                d_252_v35_ = True
                index34_ = default__.safeIndex(837, (d_214_v2_).length(0))
                (d_214_v2_)[index34_] = d_252_v35_
                d_253_v36_: _dafny.Seq
                d_253_v36_ = _dafny.SeqWithoutIsStrInference([True, (self).f6, False])
                d_254_v37_: D0
                d_254_v37_ = D0_DC1(d_239_v24_, d_242_v27_, d_215_v3_, _dafny.SeqWithoutIsStrInference([(0) - ((d_238_v23_).f13)]))
                d_255_v38_: str
                d_255_v38_ = _dafny.CodePoint('d')
                index35_ = default__.safeIndex(573, (d_249_v34_).length(0))
                index36_ = default__.safeIndex(837, (d_214_v2_).length(0))
                rhs21_ = d_244_v29_
                rhs22_ = d_245_v30_
                rhs23_ = ((d_253_v36_) + (_dafny.SeqWithoutIsStrInference([(self).f6, not((self).f6)]))) <= (_dafny.SeqWithoutIsStrInference([d_252_v35_]))
                rhs24_ = (d_238_v23_).fm8(d_252_v35_, (d_254_v37_).cf1, d_255_v38_, globalState)
                lhs16_ = d_249_v34_
                lhs17_ = default__.safeIndex(573, (d_249_v34_).length(0))
                lhs18_ = d_214_v2_
                lhs19_ = default__.safeIndex(837, (d_214_v2_).length(0))
                d_244_v29_ = rhs21_
                lhs16_[lhs17_] = rhs22_
                d_252_v35_ = rhs23_
                lhs18_[lhs19_] = rhs24_
                d_239_v24_ = d_239_v24_
                index37_ = default__.safeIndex(573, (d_249_v34_).length(0))
                (d_249_v34_)[index37_] = d_245_v30_
                index38_ = default__.safeIndex(837, (d_214_v2_).length(0))
                rhs25_ = default__.safeModuloInt((self).f5, (self).f11)
                rhs26_ = (d_239_v24_) < ((_dafny.SeqWithoutIsStrInference([d_255_v38_ for d_256_i7_ in range(default__.abs(110))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_257_i8_ in range(default__.abs(125))])).set(default__.safeIndex((d_238_v23_).f12, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_258_i8_ in range(default__.abs(125))]))), d_255_v38_)))
                rhs27_ = d_240_v25_
                rhs28_ = d_242_v27_
                rhs29_ = d_252_v35_
                lhs20_ = d_214_v2_
                lhs21_ = default__.safeIndex(837, (d_214_v2_).length(0))
                lhs22_ = self
                d_242_v27_ = rhs25_
                lhs20_[lhs21_] = rhs26_
                d_240_v25_ = rhs27_
                lhs22_.f10 = rhs28_
                d_252_v35_ = rhs29_
            elif True:
                d_259_v39_: str
                d_259_v39_ = _dafny.CodePoint('w')
                d_259_v39_ = d_259_v39_
                (self).f10 = (len(d_210_v0_) if not((self).f6) else (d_238_v23_).f12)
                d_260_v40_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_260_v40_ = nw44_
                d_261_v41_: _dafny.MultiSet
                d_261_v41_ = _dafny.MultiSet([(self).f6])
                d_262_v42_: _dafny.Seq
                d_262_v42_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
                d_263_v43_: _dafny.MultiSet
                d_263_v43_ = _dafny.MultiSet([len(d_210_v0_), p0, self.f10, len(d_262_v42_)])
                index39_ = default__.safeIndex(794, (d_260_v40_).length(0))
                (d_260_v40_)[index39_] = (d_261_v41_) - (default__.fm10(d_263_v43_, (d_238_v23_).f13, globalState))
                index40_ = default__.safeIndex(794, (d_260_v40_).length(0))
                (d_260_v40_)[index40_] = d_261_v41_
                d_259_v39_ = d_259_v39_
                d_264_v44_: bool
                d_264_v44_ = True
                d_264_v44_ = (len((d_210_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qchhltun"))))) < ((0) - ((d_237_i4_) + (self.f10)))
            if (self).f6:
                d_265_v45_: C0
                nw45_ = C0()
                nw45_.ctor__(401, (d_237_i4_ if (self).f6 else -2))
                d_265_v45_ = nw45_
                (self).f10 = (0) - ((self).f11)
                d_266_v46_: _dafny.Array
                nw46_ = _dafny.Array(int(0), 1)
                d_266_v46_ = nw46_
                d_267_v47_: _dafny.Set
                d_267_v47_ = _dafny.Set({d_266_v46_, d_266_v46_, d_266_v46_})
                index41_ = default__.safeIndex(131, (d_266_v46_).length(0))
                (d_266_v46_)[index41_] = (self).f5
                d_268_v48_: _dafny.Map
                d_268_v48_ = _dafny.Map({(d_265_v45_).f12: d_210_v0_})
                index42_ = default__.safeIndex(131, (d_266_v46_).length(0))
                rhs30_ = (default__.safeModuloInt(478, (d_265_v45_).f12)) + (len((d_268_v48_) | (d_268_v48_)))
                rhs31_ = d_267_v47_
                rhs32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihvp"))
                rhs33_ = default__.safeModuloInt((d_265_v45_).f13, (d_238_v23_).f12)
                lhs23_ = self
                lhs24_ = d_266_v46_
                lhs25_ = default__.safeIndex(131, (d_266_v46_).length(0))
                lhs23_.f10 = rhs30_
                d_267_v47_ = rhs31_
                d_210_v0_ = rhs32_
                lhs24_[lhs25_] = rhs33_
                d_214_v2_ = d_214_v2_
                d_269_v49_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_269_v49_ = nw47_
                d_270_v50_: _dafny.Seq
                d_270_v50_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                rhs34_ = (d_269_v49_ if ((d_266_v46_)[default__.safeIndex(131, (d_266_v46_).length(0))]) >= (115) else d_269_v49_)
                rhs35_ = ((d_270_v50_) + (d_270_v50_)) + (d_270_v50_)
                d_269_v49_ = rhs34_
                d_270_v50_ = rhs35_
            elif True:
                d_271_v51_: _dafny.Array
                def lambda8_(d_272_v23_):
                    def lambda9_(d_273_i9_):
                        return default__.safeDivisionInt(d_273_i9_, (d_272_v23_).f12)

                    return lambda9_

                init4_ = lambda8_(d_238_v23_)
                nw48_ = _dafny.Array(None, 3)
                for i0_4_ in range(nw48_.length(0)):
                    nw48_[i0_4_] = init4_(i0_4_)
                d_271_v51_ = nw48_
                index43_ = default__.safeIndex(153, (d_271_v51_).length(0))
                (d_271_v51_)[index43_] = p0
                d_274_v52_: _dafny.Seq
                d_274_v52_ = _dafny.SeqWithoutIsStrInference([not((self).f6)])
                d_275_v53_: str
                d_275_v53_ = _dafny.CodePoint('e')
                d_276_v54_: _dafny.Map
                d_276_v54_ = _dafny.Map({len(d_216_v4_): (self).f6})
                d_277_v55_: _dafny.Map
                d_277_v55_ = _dafny.Map({(d_238_v23_).f13: d_215_v3_})
                index44_ = default__.safeIndex(153, (d_271_v51_).length(0))
                (d_271_v51_)[index44_] = len((((d_210_v0_) + (d_210_v0_)).set(default__.safeIndex(len(d_274_v52_), len((d_210_v0_) + (d_210_v0_))), default__.fm11(d_275_v53_, True, (self).f6, globalState))) + ((default__.fm5((d_215_v3_).set((self).f6, len(d_276_v54_)), d_277_v55_, globalState)) + (d_210_v0_)))
                (self).f10 = (d_238_v23_).f12
                index45_ = default__.safeIndex(22, (d_214_v2_).length(0))
                (d_214_v2_)[index45_] = not ((self).f6) or ((self).f6)
                d_278_v56_: bool
                d_278_v56_ = True
                d_279_v57_: _dafny.Array
                nw49_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_279_v57_ = nw49_
                index46_ = default__.safeIndex(211, (d_279_v57_).length(0))
                (d_279_v57_)[index46_] = d_271_v51_
                d_280_v58_: _dafny.MultiSet
                d_280_v58_ = _dafny.MultiSet([d_278_v56_, (self).f6])
                d_281_v59_: _dafny.MultiSet
                d_281_v59_ = _dafny.MultiSet([(d_238_v23_).f12])
                index47_ = default__.safeIndex(22, (d_214_v2_).length(0))
                index48_ = default__.safeIndex(211, (d_279_v57_).length(0))
                rhs36_ = ((self).f6) == ((default__.fm10(d_281_v59_, 100, globalState)).ispropersubset(d_280_v58_))
                rhs37_ = d_278_v56_
                rhs38_ = d_271_v51_
                lhs26_ = d_214_v2_
                lhs27_ = default__.safeIndex(22, (d_214_v2_).length(0))
                lhs28_ = d_279_v57_
                lhs29_ = default__.safeIndex(211, (d_279_v57_).length(0))
                lhs26_[lhs27_] = rhs36_
                d_278_v56_ = rhs37_
                lhs28_[lhs29_] = rhs38_
                d_282_v60_: _dafny.Map
                d_282_v60_ = _dafny.Map({d_278_v56_: False})
                d_275_v53_ = (d_275_v53_ if not((len(d_282_v60_)) <= (294)) else d_275_v53_)
                d_278_v56_ = not((d_274_v52_) < (_dafny.SeqWithoutIsStrInference([(d_274_v52_)[default__.safeIndex((d_238_v23_).f13, len(d_274_v52_))]])))
            d_283_v61_: C0
            nw50_ = C0()
            nw50_.ctor__(default__.safeModuloInt(950, (self).f5), p0)
            d_283_v61_ = nw50_
        d_284_v62_: _dafny.Map
        d_284_v62_ = _dafny.Map({(self).f6: (self).f6})
        (self).f10 = (len(d_284_v62_)) * (len(((d_218_v6_).set(default__.safeIndex(p0, len(d_218_v6_)), self.f10)) + (_dafny.SeqWithoutIsStrInference([711 for d_285_i10_ in range(default__.abs(101))]))))
        d_286_v63_: bool
        d_286_v63_ = False
        d_286_v63_ = not((self).f6)

    def m4(self, p0, p1, p2, p3, globalState):
        d_287_v0_: _dafny.Seq
        d_287_v0_ = _dafny.SeqWithoutIsStrInference([True, (self).f6, (self).fm0((self).f6, globalState)])
        d_288_v1_: _dafny.Map
        d_288_v1_ = _dafny.Map({not((self).f6): d_287_v0_})
        d_288_v1_ = (d_288_v1_).set(not(False), d_287_v0_)
        hi3_ = (self).f5
        for d_289_i0_ in range(p3, hi3_):
            d_290_v2_: _dafny.MultiSet
            d_290_v2_ = _dafny.MultiSet([p2])
            d_291_v3_: _dafny.Seq
            d_291_v3_ = _dafny.SeqWithoutIsStrInference([(d_290_v2_).cardinality, default__.fm6(p3, (self).f5, (self).f6, -771, globalState)])
            d_292_v4_: _dafny.MultiSet
            d_292_v4_ = _dafny.MultiSet([self.f10, (d_291_v3_)[default__.safeIndex(p2, len(d_291_v3_))], 929])
            d_293_v5_: _dafny.Seq
            d_293_v5_ = _dafny.SeqWithoutIsStrInference([d_291_v3_])
            d_294_v6_: C0
            nw51_ = C0()
            nw51_.ctor__((self).f5, ((d_292_v4_)[len(d_293_v5_)] if (len(d_293_v5_)) in (d_292_v4_) else self.f10))
            d_294_v6_ = nw51_
            d_295_v7_: D0
            d_295_v7_ = D0_DC2(_dafny.Map({(self).f6: d_291_v3_}), (self).f6)
            d_296_v8_: _dafny.Map
            d_296_v8_ = _dafny.Map({(self).f6: _dafny.SeqWithoutIsStrInference([p3])})
            d_297_v9_: _dafny.MultiSet
            def iife23_(_pat_let7_0):
                def iife24_(d_298_dt__update__tmp_h0_):
                    def iife25_(_pat_let8_0):
                        def iife26_(d_299_dt__update_hcf6_h0_):
                            return D0_DC2((d_298_dt__update__tmp_h0_).cf5, d_299_dt__update_hcf6_h0_)
                        return iife26_(_pat_let8_0)
                    return iife25_((self).f6)
                return iife24_(_pat_let7_0)
            d_297_v9_ = _dafny.MultiSet([d_295_v7_, (d_295_v7_ if (self).f6 else iife23_(D0_DC2(d_296_v8_, False)))])
            d_297_v9_ = d_297_v9_
            d_300_v10_: bool
            d_300_v10_ = False
            d_300_v10_ = d_300_v10_
            d_301_v11_: D0
            d_301_v11_ = D0_DC0(_dafny.SeqWithoutIsStrInference([p3 for d_302_i1_ in range(default__.abs(455))]))
            source6_ = d_301_v11_
            if source6_.is_DC1:
                d_303___mcc_h0_ = source6_.cf1
                d_304___mcc_h1_ = source6_.cf2
                d_305___mcc_h2_ = source6_.cf3
                d_306___mcc_h3_ = source6_.cf4
                d_307_cf4_ = d_306___mcc_h3_
                d_308_cf3_ = d_305___mcc_h2_
                d_309_cf2_ = d_304___mcc_h1_
                d_310_cf1_ = d_303___mcc_h0_
                d_311_v12_: _dafny.Array
                nw52_ = _dafny.Array(int(0), 22)
                d_311_v12_ = nw52_
                d_311_v12_ = d_311_v12_
                rhs39_ = (self).f6
                rhs40_ = d_307_cf4_
                rhs41_ = d_300_v10_
                d_300_v10_ = rhs39_
                d_291_v3_ = rhs40_
                d_300_v10_ = rhs41_
                d_312_v13_: _dafny.MultiSet
                d_312_v13_ = _dafny.MultiSet([d_300_v10_])
                d_313_v14_: C0
                nw53_ = C0()
                nw53_.ctor__(default__.fm6(self.f10, ((d_312_v13_)[False] if (False) in (d_312_v13_) else p2), d_300_v10_, (d_294_v6_).f13, globalState), (d_294_v6_).f12)
                d_313_v14_ = nw53_
                index49_ = default__.safeIndex(446, (d_311_v12_).length(0))
                (d_311_v12_)[index49_] = (self).f5
                index50_ = default__.safeIndex(446, (d_311_v12_).length(0))
                (d_311_v12_)[index50_] = ((self).f11) * ((d_294_v6_).f13)
            elif source6_.is_DC2:
                d_314___mcc_h4_ = source6_.cf5
                d_315___mcc_h5_ = source6_.cf6
                d_316_cf6_ = d_315___mcc_h5_
                d_317_cf5_ = d_314___mcc_h4_
                d_318_v15_: _dafny.Set
                d_318_v15_ = _dafny.Set({d_316_cf6_})
                d_319_v16_: _dafny.Set
                d_319_v16_ = _dafny.Set({self.f10, p3, (d_294_v6_).f12, len(d_318_v15_), d_289_i0_})
                d_320_v17_: _dafny.Array
                nw54_ = _dafny.Array(None, 11)
                nw54_[int(0)] = len(d_319_v16_)
                nw54_[int(1)] = p3
                nw54_[int(2)] = len(d_291_v3_)
                nw54_[int(3)] = (self).f11
                nw54_[int(4)] = (d_294_v6_).f12
                nw54_[int(5)] = (d_294_v6_).f13
                nw54_[int(6)] = (self.f10) * ((d_294_v6_).f12)
                nw54_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wecoqytg")))
                nw54_[int(8)] = (self).f11
                nw54_[int(9)] = (d_294_v6_).f12
                nw54_[int(10)] = self.f10
                d_320_v17_ = nw54_
                d_320_v17_ = d_320_v17_
                d_321_v18_: _dafny.Array
                def lambda10_(d_322_v10_):
                    def lambda11_(d_323_i2_):
                        return d_322_v10_

                    return lambda11_

                init5_ = lambda10_(d_300_v10_)
                nw55_ = _dafny.Array(None, 21)
                for i0_5_ in range(nw55_.length(0)):
                    nw55_[i0_5_] = init5_(i0_5_)
                d_321_v18_ = nw55_
                d_324_v19_: _dafny.Array
                nw56_ = _dafny.Array(None, 29)
                nw56_[int(0)] = d_321_v18_
                nw56_[int(1)] = d_321_v18_
                nw56_[int(2)] = d_321_v18_
                nw56_[int(3)] = d_321_v18_
                nw56_[int(4)] = d_321_v18_
                nw56_[int(5)] = d_321_v18_
                nw56_[int(6)] = d_321_v18_
                nw56_[int(7)] = d_321_v18_
                nw56_[int(8)] = d_321_v18_
                nw56_[int(9)] = d_321_v18_
                nw56_[int(10)] = d_321_v18_
                nw56_[int(11)] = d_321_v18_
                nw56_[int(12)] = d_321_v18_
                nw56_[int(13)] = d_321_v18_
                nw56_[int(14)] = d_321_v18_
                nw56_[int(15)] = d_321_v18_
                nw56_[int(16)] = d_321_v18_
                nw56_[int(17)] = d_321_v18_
                nw56_[int(18)] = d_321_v18_
                nw56_[int(19)] = d_321_v18_
                nw56_[int(20)] = d_321_v18_
                nw56_[int(21)] = d_321_v18_
                nw56_[int(22)] = d_321_v18_
                nw56_[int(23)] = d_321_v18_
                nw56_[int(24)] = d_321_v18_
                nw56_[int(25)] = d_321_v18_
                nw56_[int(26)] = d_321_v18_
                nw56_[int(27)] = d_321_v18_
                nw56_[int(28)] = d_321_v18_
                d_324_v19_ = nw56_
                index51_ = default__.safeIndex(898, (d_324_v19_).length(0))
                (d_324_v19_)[index51_] = d_321_v18_
                index52_ = default__.safeIndex(898, (d_324_v19_).length(0))
                (d_324_v19_)[index52_] = d_321_v18_
                d_325_v20_: _dafny.Seq
                d_325_v20_ = _dafny.SeqWithoutIsStrInference([(d_324_v19_)[default__.safeIndex(898, (d_324_v19_).length(0))]])
                d_326_v21_: _dafny.Map
                d_326_v21_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_294_v6_).f12])): (d_325_v20_).set(default__.safeIndex((self).f11, len(d_325_v20_)), d_321_v18_)})
                d_327_v22_: _dafny.Seq
                d_328_v23_: _dafny.Array
                d_329_v24_: int
                d_330_v25_: int
                out28_: _dafny.Seq
                out29_: _dafny.Array
                out30_: int
                out31_: int
                out28_, out29_, out30_, out31_ = default__.m0((self.f10) - (len(((d_326_v21_)[(d_294_v6_).f13] if ((d_294_v6_).f13) in (d_326_v21_) else d_325_v20_))), 425, globalState)
                d_327_v22_ = out28_
                d_328_v23_ = out29_
                d_329_v24_ = out30_
                d_330_v25_ = out31_
                d_300_v10_ = (self).fm2(_dafny.Set({(self).f6, d_300_v10_, d_316_cf6_, d_316_cf6_, d_300_v10_}), (0) - (d_329_v24_), len(d_327_v22_), globalState)
            elif True:
                d_331___mcc_h6_ = source6_.cf0
                d_332_cf0_ = d_331___mcc_h6_
                d_333_v26_: _dafny.Map
                d_333_v26_ = _dafny.Map({d_294_v6_: d_300_v10_})
                d_334_v27_: _dafny.Set
                d_334_v27_ = _dafny.Set({d_333_v26_, d_333_v26_})
                d_335_v28_: str
                d_335_v28_ = _dafny.CodePoint('n')
                d_336_v29_: _dafny.Seq
                d_336_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcfy"))
                d_337_v30_: D1
                d_337_v30_ = D1_DC6((self).f6, len(d_334_v27_), (d_294_v6_).f12, (d_335_v28_) not in ((d_336_v29_).set(default__.safeIndex((d_294_v6_).f13, len(d_336_v29_)), d_335_v28_)))
                d_337_v30_ = d_337_v30_
                d_300_v10_ = d_300_v10_
                d_338_v31_: _dafny.MultiSet
                d_338_v31_ = _dafny.MultiSet([False])
                d_339_v32_: _dafny.Map
                d_339_v32_ = _dafny.Map({not ((self).f6) or (d_300_v10_): (d_338_v31_).ispropersubset(d_338_v31_)})
                d_339_v32_ = (d_339_v32_).set(d_300_v10_, d_300_v10_)
                (self).f10 = p2
        d_340_v33_: _dafny.Seq
        d_340_v33_ = _dafny.SeqWithoutIsStrInference([p2, self.f10, self.f10])
        d_341_v34_: _dafny.Set
        d_341_v34_ = _dafny.Set({(self).f11})
        d_342_v35_: _dafny.Map
        d_342_v35_ = _dafny.Map({(self).f6: (self).f6})
        d_343_v36_: _dafny.Array
        def lambda12_(d_344_p3_):
            def lambda13_(d_345_i3_):
                return default__.safeModuloInt(d_345_i3_, d_344_p3_)

            return lambda13_

        init6_ = lambda12_(p3)
        nw57_ = _dafny.Array(None, 5)
        for i0_6_ in range(nw57_.length(0)):
            nw57_[i0_6_] = init6_(i0_6_)
        d_343_v36_ = nw57_
        d_346_v37_: _dafny.Map
        d_346_v37_ = _dafny.Map({self.f10: d_343_v36_})
        d_347_v38_: D3
        d_347_v38_ = D3_DC11(d_343_v36_)
        d_340_v33_ = default__.fm12((len(d_341_v34_)) * (self.f10), not((len(_dafny.Map({len(d_342_v35_): ((d_346_v37_)[self.f10] if (self.f10) in (d_346_v37_) else (d_347_v38_).cf27)}))) == (431)), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_348_i4_ in range(default__.abs(-214))]), globalState)
        d_349_v39_: _dafny.Map
        d_349_v39_ = _dafny.Map({(self).f6: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbenawu")))})
        d_350_v40_: _dafny.MultiSet
        d_350_v40_ = _dafny.MultiSet([p2, (self).f5])
        d_351_v41_: _dafny.Map
        d_351_v41_ = _dafny.Map({len(d_349_v39_): (d_350_v40_).intersection(d_350_v40_)})
        d_352_v42_: _dafny.Map
        d_352_v42_ = _dafny.Map({(self).f6: d_340_v33_})
        d_353_v43_: D0
        d_353_v43_ = D0_DC2(d_352_v42_, (self).f6)
        d_354_v44_: _dafny.Set
        d_354_v44_ = _dafny.Set({not(True)})
        d_351_v41_ = (d_351_v41_).set(((self).f11 if (d_353_v43_).cf6 else default__.fm6(len(d_354_v44_), p3, False, 393, globalState)), (d_350_v40_).set((self).f5, default__.abs(515)))
        d_355_v45_: str
        d_355_v45_ = _dafny.CodePoint('i')
        d_356_v46_: _dafny.Map
        d_356_v46_ = _dafny.Map({d_355_v45_: p3})
        d_357_v47_: D4
        d_357_v47_ = D4_DC13(d_356_v46_)
        pat_let_tv7_ = d_356_v46_
        d_358_v48_: C0
        nw58_ = C0()
        def iife27_(_pat_let9_0):
            def iife28_(d_359_dt__update__tmp_h1_):
                def iife29_(_pat_let10_0):
                    def iife30_(d_360_dt__update_hcf30_h0_):
                        return D4_DC13(d_360_dt__update_hcf30_h0_)
                    return iife30_(_pat_let10_0)
                return iife29_(pat_let_tv7_)
            return iife28_(_pat_let9_0)
        nw58_.ctor__((self).f11, (len((iife27_(d_357_v47_)).cf30) if not(False) else p2))
        d_358_v48_ = nw58_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_343_v36_).length(0)):
            d_361_i5_: int = guard_loop_0_
            if (True) and (((0) <= (d_361_i5_)) and ((d_361_i5_) < ((d_343_v36_).length(0)))):
                (d_343_v36_)[(d_361_i5_)] = (d_361_i5_) + (default__.safeModuloInt((self).f11, (d_358_v48_).f12))

    @property
    def f11(self):
        return self._f11

class C2:
    def  __init__(self):
        self.f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f9):
        (self).f9 = f9

    def fm1(self, p0, p1, globalState):
        return ((0) - (len((_dafny.Map({_dafny.CodePoint('r'): True})) | (_dafny.Map({_dafny.CodePoint('t'): True}))))) < (self.f9)

    def m3(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        (self).f9 = self.f9
        d_362_v0_: bool
        d_362_v0_ = False
        d_363_i0_: int
        d_363_i0_ = 0
        with _dafny.label("1"):
            while (d_362_v0_) or (d_362_v0_):
                with _dafny.c_label("1"):
                    if (d_363_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_363_i0_ = (d_363_i0_) + (1)
                    d_364_v1_: T0
                    nw59_ = C1()
                    nw59_.ctor__(self.f9, len(_dafny.SeqWithoutIsStrInference([d_362_v0_, d_362_v0_, d_362_v0_, d_362_v0_])), self.f9, d_362_v0_)
                    d_364_v1_ = nw59_
                    d_365_v2_: _dafny.Set
                    d_365_v2_ = _dafny.Set({(d_364_v1_ if d_362_v0_ else d_364_v1_)})
                    d_366_v3_: _dafny.Map
                    d_366_v3_ = _dafny.Map({p0: _dafny.Set({d_364_v1_})})
                    d_367_v4_: _dafny.Seq
                    d_367_v4_ = _dafny.SeqWithoutIsStrInference([d_365_v2_])
                    d_365_v2_ = (((d_366_v3_)[p0] if (p0) in (d_366_v3_) else (d_367_v4_)[default__.safeIndex(249, len(d_367_v4_))])) - ((d_365_v2_) - (d_365_v2_))
                    d_368_v5_: _dafny.Map
                    d_368_v5_ = _dafny.Map({default__.safeDivisionInt((d_364_v1_).f5, 613): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_369_i1_ in range(default__.abs(193))])})
                    d_368_v5_ = (d_368_v5_).set((d_364_v1_).f5, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byj")) if (d_364_v1_).f6 else p0))
                    d_370_v6_: _dafny.Array
                    nw60_ = _dafny.Array(False, 2)
                    d_370_v6_ = nw60_
                    index53_ = default__.safeIndex(353, (d_370_v6_).length(0))
                    (d_370_v6_)[index53_] = d_362_v0_
                    index54_ = default__.safeIndex(353, (d_370_v6_).length(0))
                    (d_370_v6_)[index54_] = not(not(((d_364_v1_).f5) == (default__.fm6((d_364_v1_).f5, self.f9, d_362_v0_, (d_364_v1_).f5, globalState))))
                    d_371_v7_: _dafny.Map
                    d_371_v7_ = _dafny.Map({d_362_v0_: self.f9})
                    d_372_v8_: _dafny.Map
                    d_372_v8_ = _dafny.Map({(d_370_v6_)[default__.safeIndex(353, (d_370_v6_).length(0))]: d_371_v7_})
                    d_373_v9_: _dafny.Seq
                    d_373_v9_ = _dafny.SeqWithoutIsStrInference([(d_364_v1_).f5])
                    d_374_v10_: D0
                    d_374_v10_ = D0_DC1(p0, self.f9, d_371_v7_, d_373_v9_)
                    d_375_v11_: _dafny.Map
                    d_375_v11_ = _dafny.Map({self.f9: d_371_v7_})
                    d_376_v12_: _dafny.MultiSet
                    d_376_v12_ = _dafny.MultiSet([(d_364_v1_).f5])
                    d_377_v13_: D1
                    d_377_v13_ = D1_DC6(d_362_v0_, (d_376_v12_).cardinality, (d_373_v9_)[default__.safeIndex((d_364_v1_).f5, len(d_373_v9_))], (d_364_v1_).f6)
                    d_378_v14_: _dafny.Seq
                    d_378_v14_ = _dafny.SeqWithoutIsStrInference([(d_370_v6_)[default__.safeIndex(353, (d_370_v6_).length(0))], (d_370_v6_)[default__.safeIndex(353, (d_370_v6_).length(0))], (d_370_v6_)[default__.safeIndex(353, (d_370_v6_).length(0))], (d_364_v1_).f6, (d_377_v13_).cf20])
                    d_379_v15_: _dafny.MultiSet
                    d_379_v15_ = _dafny.MultiSet([len(d_378_v14_), (d_364_v1_).f5, self.f9, (d_364_v1_).f5, self.f9])
                    d_380_v16_: D2
                    d_380_v16_ = D2_DC8(d_379_v15_)
                    d_381_v17_: _dafny.Array
                    nw61_ = _dafny.Array(None, 19)
                    nw61_[int(0)] = ((d_372_v8_)[(d_364_v1_).f6] if ((d_364_v1_).f6) in (d_372_v8_) else _dafny.Map({(d_370_v6_)[default__.safeIndex(353, (d_370_v6_).length(0))]: (d_364_v1_).f5}))
                    nw61_[int(1)] = d_371_v7_
                    nw61_[int(2)] = d_371_v7_
                    nw61_[int(3)] = ((d_374_v10_).cf3) | (d_371_v7_)
                    nw61_[int(4)] = d_371_v7_
                    nw61_[int(5)] = d_371_v7_
                    nw61_[int(6)] = d_371_v7_
                    nw61_[int(7)] = d_371_v7_
                    nw61_[int(8)] = d_371_v7_
                    nw61_[int(9)] = d_371_v7_
                    nw61_[int(10)] = d_371_v7_
                    nw61_[int(11)] = d_371_v7_
                    nw61_[int(12)] = (d_371_v7_) | (((d_375_v11_)[(d_364_v1_).f5] if ((d_364_v1_).f5) in (d_375_v11_) else d_371_v7_))
                    nw61_[int(13)] = (d_371_v7_) | (d_371_v7_)
                    nw61_[int(14)] = d_371_v7_
                    nw61_[int(15)] = (d_371_v7_).set((d_364_v1_).f6, (d_364_v1_).f5)
                    nw61_[int(16)] = d_371_v7_
                    nw61_[int(17)] = default__.fm13(p0, ((d_379_v15_)[(0) - ((d_364_v1_).f5)] if ((0) - ((d_364_v1_).f5)) in (d_379_v15_) else 144), (d_380_v16_).cf22, globalState)
                    nw61_[int(18)] = d_371_v7_
                    d_381_v17_ = nw61_
                    d_381_v17_ = d_381_v17_
                    pass
            pass
        d_382_v18_: _dafny.Map
        d_382_v18_ = _dafny.Map({self.f9: self.f9})
        d_383_v19_: _dafny.Seq
        d_383_v19_ = _dafny.SeqWithoutIsStrInference([-537, default__.fm6(len(d_382_v18_), 210, d_362_v0_, self.f9, globalState)])
        d_384_i2_: int
        d_384_i2_ = 0
        with _dafny.label("2"):
            while (d_383_v19_) == ((d_383_v19_ if d_362_v0_ else _dafny.SeqWithoutIsStrInference([self.f9]))):
                with _dafny.c_label("2"):
                    if (d_384_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_384_i2_ = (d_384_i2_) + (1)
                    d_385_v20_: _dafny.Seq
                    d_385_v20_ = _dafny.SeqWithoutIsStrInference([d_362_v0_])
                    source7_ = default__.fm14((d_362_v0_) not in (d_385_v20_), globalState)
                    if source7_.is_DC4:
                        d_386___mcc_h0_ = source7_.cf8
                        d_387___mcc_h1_ = source7_.cf9
                        d_388___mcc_h2_ = source7_.cf10
                        d_389___mcc_h3_ = source7_.cf11
                        d_390___mcc_h4_ = source7_.cf12
                        d_391_cf12_ = d_390___mcc_h4_
                        d_392_cf11_ = d_389___mcc_h3_
                        d_393_cf10_ = d_388___mcc_h2_
                        d_394_cf9_ = d_387___mcc_h1_
                        d_395_cf8_ = d_386___mcc_h0_
                        rhs42_ = True
                        rhs43_ = d_394_cf9_
                        d_394_cf9_ = rhs42_
                        d_394_cf9_ = rhs43_
                        d_396_v21_: _dafny.MultiSet
                        d_396_v21_ = _dafny.MultiSet([d_362_v0_, d_394_cf9_, False, not(not (d_394_cf9_) or (d_395_cf8_))])
                        d_396_v21_ = d_396_v21_
                        d_394_cf9_ = d_395_cf8_
                        d_397_v22_: _dafny.Map
                        d_397_v22_ = _dafny.Map({d_395_cf8_: default__.fm12(d_392_cf11_, d_362_v0_, d_393_cf10_, globalState)})
                        d_394_cf9_ = (d_397_v22_) != ((default__.fm15(d_392_cf11_, globalState)) | (d_397_v22_))
                    elif source7_.is_DC5:
                        d_398___mcc_h5_ = source7_.cf13
                        d_399___mcc_h6_ = source7_.cf14
                        d_400___mcc_h7_ = source7_.cf15
                        d_401___mcc_h8_ = source7_.cf16
                        d_402_cf16_ = d_401___mcc_h8_
                        d_403_cf15_ = d_400___mcc_h7_
                        d_404_cf14_ = d_399___mcc_h6_
                        d_405_cf13_ = d_398___mcc_h5_
                        d_405_cf13_ = False
                        d_385_v20_ = default__.fm16(globalState)
                        d_406_v23_: _dafny.MultiSet
                        d_406_v23_ = _dafny.MultiSet([575])
                        d_407_v24_: _dafny.MultiSet
                        d_407_v24_ = _dafny.MultiSet([d_362_v0_, False, d_405_cf13_, d_362_v0_])
                        d_408_v25_: D2
                        d_408_v25_ = D2_DC8((d_406_v23_) | (_dafny.MultiSet([self.f9, ((d_407_v24_)[d_362_v0_] if (d_362_v0_) in (d_407_v24_) else d_402_cf16_), self.f9])))
                        pat_let_tv8_ = d_406_v23_
                        pat_let_tv9_ = d_402_cf16_
                        pat_let_tv10_ = d_403_cf15_
                        pat_let_tv11_ = d_402_cf16_
                        def iife32_(_pat_let12_0):
                            def iife33_(d_409_dt__update__tmp_h0_):
                                def iife34_(_pat_let13_0):
                                    def iife35_(d_410_dt__update_hcf22_h0_):
                                        return D2_DC8(d_410_dt__update_hcf22_h0_)
                                    return iife35_(_pat_let13_0)
                                return iife34_(pat_let_tv8_)
                            return iife33_(_pat_let12_0)
                        def iife31_(_pat_let11_0):
                            def iife36_(d_411_dt__update__tmp_h1_):
                                def iife37_(_pat_let14_0):
                                    def iife38_(d_412_dt__update_hcf22_h1_):
                                        return D2_DC8(d_412_dt__update_hcf22_h1_)
                                    return iife38_(_pat_let14_0)
                                return iife37_(_dafny.MultiSet([pat_let_tv9_, pat_let_tv10_, pat_let_tv11_]))
                            return iife36_(_pat_let11_0)
                        d_408_v25_ = iife31_(iife32_(d_408_v25_))
                        d_413_v26_: str
                        d_413_v26_ = _dafny.CodePoint('c')
                        d_413_v26_ = d_413_v26_
                    elif source7_.is_DC6:
                        d_414___mcc_h9_ = source7_.cf17
                        d_415___mcc_h10_ = source7_.cf18
                        d_416___mcc_h11_ = source7_.cf19
                        d_417___mcc_h12_ = source7_.cf20
                        d_418_cf20_ = d_417___mcc_h12_
                        d_419_cf19_ = d_416___mcc_h11_
                        d_420_cf18_ = d_415___mcc_h10_
                        d_421_cf17_ = d_414___mcc_h9_
                        d_362_v0_ = d_418_cf20_
                        d_422_v27_: _dafny.Seq
                        d_423_v28_: _dafny.Array
                        d_424_v29_: int
                        d_425_v30_: int
                        out32_: _dafny.Seq
                        out33_: _dafny.Array
                        out34_: int
                        out35_: int
                        out32_, out33_, out34_, out35_ = default__.m0(d_420_cf18_, default__.fm6(self.f9, self.f9, d_418_cf20_, self.f9, globalState), globalState)
                        d_422_v27_ = out32_
                        d_423_v28_ = out33_
                        d_424_v29_ = out34_
                        d_425_v30_ = out35_
                        d_426_v31_: _dafny.Map
                        d_426_v31_ = _dafny.Map({d_421_cf17_: d_362_v0_})
                        d_427_v32_: _dafny.Map
                        d_427_v32_ = _dafny.Map({647: d_426_v31_})
                        d_428_v33_: C1
                        nw62_ = C1()
                        nw62_.ctor__((d_425_v30_) * (len(((d_427_v32_)[984] if (984) in (d_427_v32_) else d_426_v31_))), 365, d_424_v29_, d_418_cf20_)
                        d_428_v33_ = nw62_
                        d_428_v33_ = d_428_v33_
                        d_429_v34_: D2
                        d_429_v34_ = D2_DC9(d_362_v0_, (d_428_v33_).f11, d_419_cf19_)
                        d_430_v35_: _dafny.Set
                        d_430_v35_ = _dafny.Set({d_362_v0_})
                        d_431_v36_: D1
                        d_431_v36_ = D1_DC6(d_421_cf17_, len(d_430_v35_), len(d_422_v27_), d_362_v0_)
                        d_432_v37_: _dafny.Map
                        d_432_v37_ = _dafny.Map({(d_431_v36_).cf17: d_383_v19_})
                        d_433_v38_: D0
                        d_433_v38_ = D0_DC2(d_432_v37_, d_418_cf20_)
                        d_434_v39_: _dafny.MultiSet
                        d_434_v39_ = _dafny.MultiSet([d_418_cf20_])
                        d_435_v40_: str
                        d_435_v40_ = _dafny.CodePoint('u')
                        d_436_v41_: _dafny.Array
                        nw63_ = _dafny.Array(None, 28)
                        nw63_[int(0)] = (not(d_362_v0_)) == (d_418_cf20_)
                        nw63_[int(1)] = d_362_v0_
                        nw63_[int(2)] = (D2_DC9((d_429_v34_).cf23, 809, len(p0))).cf23
                        nw63_[int(3)] = not(d_418_cf20_)
                        nw63_[int(4)] = d_362_v0_
                        nw63_[int(5)] = d_362_v0_
                        nw63_[int(6)] = not (d_362_v0_) or (d_421_cf17_)
                        nw63_[int(7)] = d_421_cf17_
                        nw63_[int(8)] = not(not(not (d_362_v0_) or (d_418_cf20_)))
                        nw63_[int(9)] = d_418_cf20_
                        nw63_[int(10)] = d_421_cf17_
                        nw63_[int(11)] = (d_362_v0_) not in (d_426_v31_)
                        nw63_[int(12)] = d_421_cf17_
                        nw63_[int(13)] = (d_429_v34_).cf23
                        nw63_[int(14)] = d_418_cf20_
                        nw63_[int(15)] = d_362_v0_
                        nw63_[int(16)] = d_421_cf17_
                        nw63_[int(17)] = d_362_v0_
                        nw63_[int(18)] = not ((d_433_v38_).cf6) or (d_362_v0_)
                        nw63_[int(19)] = d_418_cf20_
                        nw63_[int(20)] = (_dafny.MultiSet([not(d_418_cf20_)])).isdisjoint(d_434_v39_)
                        nw63_[int(21)] = d_418_cf20_
                        nw63_[int(22)] = (d_383_v19_) == (d_383_v19_)
                        nw63_[int(23)] = (d_428_v33_).fm2(d_430_v35_, (d_428_v33_).f11, d_420_cf18_, globalState)
                        nw63_[int(24)] = d_418_cf20_
                        nw63_[int(25)] = (self).fm1(d_428_v33_.f10, d_435_v40_, globalState)
                        nw63_[int(26)] = (d_385_v20_)[default__.safeIndex(d_425_v30_, len(d_385_v20_))]
                        nw63_[int(27)] = d_362_v0_
                        d_436_v41_ = nw63_
                        index55_ = default__.safeIndex(564, (d_436_v41_).length(0))
                        (d_436_v41_)[index55_] = d_421_cf17_
                        d_437_v42_: _dafny.Map
                        d_437_v42_ = _dafny.Map({d_362_v0_: len(d_422_v27_)})
                        d_438_v43_: _dafny.Seq
                        d_438_v43_ = _dafny.SeqWithoutIsStrInference([d_437_v42_])
                        d_439_v44_: D0
                        d_439_v44_ = D0_DC1(p0, d_424_v29_, (d_438_v43_)[default__.safeIndex(d_428_v33_.f10, len(d_438_v43_))], _dafny.SeqWithoutIsStrInference([d_425_v30_ for d_440_i3_ in range(default__.abs(557))]))
                        index56_ = default__.safeIndex(564, (d_436_v41_).length(0))
                        (d_436_v41_)[index56_] = (d_435_v40_) in ((d_439_v44_).cf1)
                    elif source7_.is_DC3:
                        d_441___mcc_h13_ = source7_.cf7
                        d_442_cf7_ = d_441___mcc_h13_
                        d_443_v45_: _dafny.Array
                        def lambda14_(d_444_i4_):
                            return (_dafny.CodePoint('w') if True else _dafny.CodePoint('l'))

                        init7_ = lambda14_
                        nw64_ = _dafny.Array(None, 19)
                        for i0_7_ in range(nw64_.length(0)):
                            nw64_[i0_7_] = init7_(i0_7_)
                        d_443_v45_ = nw64_
                        d_445_v46_: str
                        d_445_v46_ = _dafny.CodePoint('f')
                        index57_ = default__.safeIndex(407, (d_443_v45_).length(0))
                        (d_443_v45_)[index57_] = d_445_v46_
                        index58_ = default__.safeIndex(407, (d_443_v45_).length(0))
                        (d_443_v45_)[index58_] = d_445_v46_
                        d_446_v47_: _dafny.Array
                        nw65_ = _dafny.Array(False, 22)
                        d_446_v47_ = nw65_
                        index59_ = default__.safeIndex(588, (d_446_v47_).length(0))
                        (d_446_v47_)[index59_] = not(d_362_v0_)
                        index60_ = default__.safeIndex(588, (d_446_v47_).length(0))
                        (d_446_v47_)[index60_] = d_362_v0_
                        d_447_v48_: _dafny.MultiSet
                        d_447_v48_ = _dafny.MultiSet([(0) - (self.f9), self.f9, self.f9, self.f9])
                        d_448_v49_: D2
                        d_448_v49_ = D2_DC8(d_447_v48_)
                        pat_let_tv12_ = d_447_v48_
                        def iife39_(_pat_let15_0):
                            def iife40_(d_449_dt__update__tmp_h2_):
                                def iife41_(_pat_let16_0):
                                    def iife42_(d_450_dt__update_hcf22_h2_):
                                        return D2_DC8(d_450_dt__update_hcf22_h2_)
                                    return iife42_(_pat_let16_0)
                                return iife41_(pat_let_tv12_)
                            return iife40_(_pat_let15_0)
                        d_448_v49_ = iife39_(d_448_v49_)
                        d_451_v50_: _dafny.Array
                        nw66_ = _dafny.Array(_dafny.Array(None, 0), 5)
                        d_451_v50_ = nw66_
                        index61_ = default__.safeIndex(834, (d_451_v50_).length(0))
                        (d_451_v50_)[index61_] = d_446_v47_
                        index62_ = default__.safeIndex(834, (d_451_v50_).length(0))
                        nw67_ = _dafny.Array(False, 18)
                        (d_451_v50_)[index62_] = nw67_
                    elif True:
                        d_452___mcc_h14_ = source7_.cf21
                        d_453_cf21_ = d_452___mcc_h14_
                        d_454_v51_: _dafny.Map
                        d_454_v51_ = _dafny.Map({500: d_362_v0_})
                        d_455_v52_: D1
                        d_455_v52_ = D1_DC3(d_454_v51_)
                        d_456_v53_: D1
                        d_456_v53_ = D1_DC7(d_455_v52_)
                        d_456_v53_ = d_456_v53_
                        d_457_v54_: _dafny.Seq
                        d_458_v55_: _dafny.Array
                        d_459_v56_: int
                        d_460_v57_: int
                        out36_: _dafny.Seq
                        out37_: _dafny.Array
                        out38_: int
                        out39_: int
                        out36_, out37_, out38_, out39_ = default__.m0(self.f9, self.f9, globalState)
                        d_457_v54_ = out36_
                        d_458_v55_ = out37_
                        d_459_v56_ = out38_
                        d_460_v57_ = out39_
                        index63_ = default__.safeIndex(163, (d_458_v55_).length(0))
                        def iife43_():
                            coll9_ = _dafny.Map()
                            compr_9_: int
                            for compr_9_ in _dafny.IntegerRange(809, 186):
                                d_461_v58_: int = compr_9_
                                if ((809) <= (d_461_v58_)) and ((d_461_v58_) < (186)):
                                    coll9_[default__.safeModuloInt(d_461_v58_, self.f9)] = d_457_v54_
                            return _dafny.Map(coll9_)
                        (d_458_v55_)[index63_] = len(iife43_()
                        )
                        index64_ = default__.safeIndex(163, (d_458_v55_).length(0))
                        (d_458_v55_)[index64_] = d_460_v57_
                        d_462_v59_: _dafny.Set
                        d_462_v59_ = _dafny.Set({d_362_v0_, d_362_v0_})
                        d_463_v60_: D1
                        d_463_v60_ = D1_DC5(d_362_v0_, d_362_v0_, len(d_462_v59_), 90)
                        d_464_v61_: _dafny.Map
                        d_464_v61_ = _dafny.Map({d_463_v60_: d_383_v19_})
                        d_464_v61_ = (d_464_v61_).set(d_463_v60_, (_dafny.SeqWithoutIsStrInference([len(d_457_v54_)])) + (d_383_v19_))
                    d_465_v62_: _dafny.Array
                    nw68_ = _dafny.Array(False, 1)
                    d_465_v62_ = nw68_
                    index65_ = default__.safeIndex(252, (d_465_v62_).length(0))
                    (d_465_v62_)[index65_] = (d_385_v20_)[default__.safeIndex((0) - (len(p0)), len(d_385_v20_))]
                    index66_ = default__.safeIndex(252, (d_465_v62_).length(0))
                    (d_465_v62_)[index66_] = d_362_v0_
                    index67_ = default__.safeIndex(252, (d_465_v62_).length(0))
                    (d_465_v62_)[index67_] = ((d_465_v62_)[default__.safeIndex(252, (d_465_v62_).length(0))]) or ((d_465_v62_)[default__.safeIndex(252, (d_465_v62_).length(0))])
                    index68_ = default__.safeIndex(252, (d_465_v62_).length(0))
                    (d_465_v62_)[index68_] = (self.f9) == (self.f9)
                    pass
            pass
        d_362_v0_ = (d_362_v0_) == (d_362_v0_)
        d_466_v63_: D4
        d_466_v63_ = D4_DC14()
        d_466_v63_ = d_466_v63_
        d_467_v64_: _dafny.Map
        d_467_v64_ = _dafny.Map({self.f9: d_362_v0_})
        d_467_v64_ = (d_467_v64_).set(self.f9, d_362_v0_)
        d_468_v65_: str
        d_468_v65_ = _dafny.CodePoint('h')
        d_469_v66_: D0
        d_469_v66_ = D0_DC0(_dafny.SeqWithoutIsStrInference([-920]))
        d_470_v67_: _dafny.Map
        d_470_v67_ = _dafny.Map({d_468_v65_: d_469_v66_})
        d_471_v68_: _dafny.Set
        d_471_v68_ = _dafny.Set({d_362_v0_, d_362_v0_})
        d_472_v69_: _dafny.Map
        d_472_v69_ = _dafny.Map({d_470_v67_: len(d_471_v68_)})
        d_473_v71_: _dafny.Map
        d_473_v71_ = _dafny.Map({not(d_362_v0_): d_472_v69_})
        d_474_v74_: _dafny.MultiSet
        d_474_v74_ = _dafny.MultiSet([d_468_v65_])
        pat_let_tv13_ = d_383_v19_
        d_475_v75_: _dafny.Set
        def iife44_():
            coll10_ = _dafny.Map()
            compr_10_: str
            for compr_10_ in (d_474_v74_).Elements:
                d_476_v73_: str = compr_10_
                if (d_476_v73_) in (d_474_v74_):
                    coll10_[d_476_v73_] = d_469_v66_
            return _dafny.Map(coll10_)
        def iife45_(_pat_let17_0):
            def iife46_(d_477_dt__update__tmp_h3_):
                def iife47_(_pat_let18_0):
                    def iife48_(d_478_dt__update_hcf0_h0_):
                        return D0_DC0(d_478_dt__update_hcf0_h0_)
                    return iife48_(_pat_let18_0)
                return iife47_(pat_let_tv13_)
            return iife46_(_pat_let17_0)
        d_475_v75_ = _dafny.Set({iife44_()
        , _dafny.Map({d_468_v65_: iife45_(d_469_v66_)})})
        d_479_v76_: _dafny.Map
        d_479_v76_ = _dafny.Map({self.f9: d_475_v75_})
        d_480_v77_: _dafny.Seq
        def iife49_():
            def iife51_():
                coll13_ = _dafny.Map()
                compr_13_: _dafny.Map
                for compr_13_ in (((d_479_v76_)[self.f9] if (self.f9) in (d_479_v76_) else d_475_v75_)).Elements:
                    d_481_v72_: _dafny.Map = compr_13_
                    if (d_481_v72_) in (((d_479_v76_)[self.f9] if (self.f9) in (d_479_v76_) else d_475_v75_)):
                        coll13_[d_481_v72_] = self.f9
                return _dafny.Map(coll13_)
            coll11_ = _dafny.Map()
            def iife50_():
                coll12_ = _dafny.Map()
                compr_12_: _dafny.Map
                for compr_12_ in (((d_479_v76_)[self.f9] if (self.f9) in (d_479_v76_) else d_475_v75_)).Elements:
                    d_481_v72_: _dafny.Map = compr_12_
                    if (d_481_v72_) in (((d_479_v76_)[self.f9] if (self.f9) in (d_479_v76_) else d_475_v75_)):
                        coll12_[d_481_v72_] = self.f9
                return _dafny.Map(coll12_)
            compr_11_: _dafny.Map
            for compr_11_ in (((d_473_v71_)[d_362_v0_] if (d_362_v0_) in (d_473_v71_) else iife50_()
            )).keys.Elements:
                d_482_v70_: _dafny.Map = compr_11_
                if (d_482_v70_) in (((d_473_v71_)[d_362_v0_] if (d_362_v0_) in (d_473_v71_) else iife51_()
                )):
                    coll11_[d_482_v70_] = self.f9
            return _dafny.Map(coll11_)
        d_480_v77_ = _dafny.SeqWithoutIsStrInference([d_472_v69_, iife49_()
        , d_472_v69_])
        r0 = (d_480_v77_)[default__.safeIndex(self.f9, len(d_480_v77_))]
        return r0


class C3(T0):
    def  __init__(self):
        self._f5: int = int(0)
        self._f6: bool = False
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f8, f5, f6):
        (self)._f8 = f8
        (self)._f5 = f5
        (self)._f6 = f6

    def fm0(self, p0, globalState):
        return not(((self).f8) >= ((self).f5))

    def m1(self, p0, globalState):
        d_483_i0_: int
        d_483_i0_ = 0
        with _dafny.label("3"):
            while (self).f6:
                with _dafny.c_label("3"):
                    if (d_483_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_483_i0_ = (d_483_i0_) + (1)
                    d_484_v0_: bool
                    d_484_v0_ = False
                    d_484_v0_ = not(d_484_v0_)
                    d_485_v1_: _dafny.MultiSet
                    d_485_v1_ = _dafny.MultiSet([(self).f6, d_484_v0_])
                    d_486_v2_: D5
                    d_486_v2_ = D5_DC15(d_485_v1_)
                    d_487_v3_: _dafny.Seq
                    d_487_v3_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f6, (self).f6])).cardinality, (self).f8, ((d_486_v2_).cf31).cardinality, p0])
                    d_488_v4_: C2
                    nw69_ = C2()
                    nw69_.ctor__(((d_487_v3_)[default__.safeIndex((self).f5, len(d_487_v3_))]) + (p0))
                    d_488_v4_ = nw69_
                    d_489_v5_: _dafny.Set
                    d_489_v5_ = _dafny.Set({p0, (self).f5, p0, (self).f5})
                    d_490_v6_: _dafny.Map
                    d_490_v6_ = _dafny.Map({d_488_v4_.f9: 408})
                    d_491_v7_: _dafny.Map
                    d_491_v7_ = _dafny.Map({d_484_v0_: d_484_v0_})
                    d_492_v8_: _dafny.Array
                    nw70_ = _dafny.Array(None, 15)
                    nw70_[int(0)] = default__.fm6(56, d_488_v4_.f9, d_484_v0_, p0, globalState)
                    nw70_[int(1)] = (self).f8
                    nw70_[int(2)] = d_488_v4_.f9
                    nw70_[int(3)] = d_488_v4_.f9
                    nw70_[int(4)] = d_488_v4_.f9
                    nw70_[int(5)] = default__.safeModuloInt((self).f8, (self).f5)
                    nw70_[int(6)] = (0) - (d_488_v4_.f9)
                    nw70_[int(7)] = (d_488_v4_.f9) * (len(d_489_v5_))
                    nw70_[int(8)] = default__.fm6(d_488_v4_.f9, (self).f8, d_484_v0_, (self).f5, globalState)
                    nw70_[int(9)] = (self).f8
                    nw70_[int(10)] = ((d_490_v6_)[(self).f5] if ((self).f5) in (d_490_v6_) else (d_487_v3_)[default__.safeIndex((0) - (p0), len(d_487_v3_))])
                    nw70_[int(11)] = d_488_v4_.f9
                    nw70_[int(12)] = default__.safeModuloInt(d_488_v4_.f9, p0)
                    nw70_[int(13)] = len((_dafny.SeqWithoutIsStrInference([d_488_v4_.f9 for d_493_i1_ in range(default__.abs(889))])) + (d_487_v3_))
                    nw70_[int(14)] = ((0) - (len(d_491_v7_)) if d_484_v0_ else p0)
                    d_492_v8_ = nw70_
                    index69_ = default__.safeIndex(279, (d_492_v8_).length(0))
                    (d_492_v8_)[index69_] = d_488_v4_.f9
                    d_494_v9_: _dafny.Array
                    nw71_ = _dafny.Array(_dafny.Array(None, 0), 20)
                    d_494_v9_ = nw71_
                    d_495_v10_: _dafny.Array
                    nw72_ = _dafny.Array(False, 29)
                    d_495_v10_ = nw72_
                    index70_ = default__.safeIndex(888, (d_494_v9_).length(0))
                    (d_494_v9_)[index70_] = d_495_v10_
                    index71_ = default__.safeIndex(279, (d_492_v8_).length(0))
                    index72_ = default__.safeIndex(888, (d_494_v9_).length(0))
                    rhs44_ = (0) - ((0) - ((self).f5))
                    rhs45_ = d_495_v10_
                    rhs46_ = (self).f5
                    lhs30_ = d_492_v8_
                    lhs31_ = default__.safeIndex(279, (d_492_v8_).length(0))
                    lhs32_ = d_494_v9_
                    lhs33_ = default__.safeIndex(888, (d_494_v9_).length(0))
                    lhs34_ = d_488_v4_
                    lhs30_[lhs31_] = rhs44_
                    lhs32_[lhs33_] = rhs45_
                    lhs34_.f9 = rhs46_
                    arr0_ = (d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]
                    index73_ = default__.safeIndex(338, ((d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]).length(0))
                    arr0_[index73_] = (self).f6
                    d_496_v11_: T0
                    nw73_ = C1()
                    nw73_.ctor__((d_492_v8_)[default__.safeIndex(279, (d_492_v8_).length(0))], (self).f5, d_488_v4_.f9, (self).f6)
                    d_496_v11_ = nw73_
                    d_497_v12_: str
                    d_497_v12_ = _dafny.CodePoint('b')
                    d_498_v13_: _dafny.Seq
                    d_498_v13_ = _dafny.SeqWithoutIsStrInference([d_497_v12_])
                    d_499_v14_: D5
                    d_499_v14_ = D5_DC16(d_496_v11_, False, len(d_498_v13_))
                    pat_let_tv14_ = d_484_v0_
                    arr1_ = (d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]
                    index74_ = default__.safeIndex(338, ((d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]).length(0))
                    def iife52_(_pat_let19_0):
                        def iife53_(d_500_dt__update__tmp_h0_):
                            def iife54_(_pat_let20_0):
                                def iife55_(d_501_dt__update_hcf33_h0_):
                                    return D5_DC16((d_500_dt__update__tmp_h0_).cf32, d_501_dt__update_hcf33_h0_, (d_500_dt__update__tmp_h0_).cf34)
                                return iife55_(_pat_let20_0)
                            return iife54_(pat_let_tv14_)
                        return iife53_(_pat_let19_0)
                    rhs47_ = (iife52_(d_499_v14_)).cf33
                    rhs48_ = (self).f6
                    lhs35_ = (d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]
                    lhs36_ = default__.safeIndex(338, ((d_494_v9_)[default__.safeIndex(888, (d_494_v9_).length(0))]).length(0))
                    d_484_v0_ = rhs47_
                    lhs35_[lhs36_] = rhs48_
                    pass
            pass
        if (p0) <= (p0):
            d_502_v15_: str
            d_502_v15_ = _dafny.CodePoint('o')
            d_503_v16_: _dafny.Map
            d_503_v16_ = _dafny.Map({p0: d_502_v15_})
            d_504_v17_: _dafny.MultiSet
            d_504_v17_ = _dafny.MultiSet([(self).f6, (self).f6, (self).f6, (self).fm0((self).f6, globalState), (self).f6])
            d_505_v18_: C0
            nw74_ = C0()
            nw74_.ctor__((d_504_v17_).cardinality, (self).f8)
            d_505_v18_ = nw74_
            d_506_v19_: _dafny.Map
            d_506_v19_ = _dafny.Map({d_505_v18_: d_502_v15_})
            d_507_v20_: _dafny.MultiSet
            d_507_v20_ = _dafny.MultiSet([d_503_v16_, (d_503_v16_).set(len(d_506_v19_), default__.fm11(d_502_v15_, (self).f6, (self).f6, globalState))])
            d_508_v21_: _dafny.Seq
            d_508_v21_ = _dafny.SeqWithoutIsStrInference([d_507_v20_])
            d_509_v22_: _dafny.MultiSet
            d_509_v22_ = _dafny.MultiSet([(d_505_v18_).f13])
            d_510_v23_: _dafny.MultiSet
            d_510_v23_ = _dafny.MultiSet([(d_509_v22_).cardinality, default__.fm6((d_505_v18_).f12, (self).f5, (self).f6, (d_505_v18_).f12, globalState), default__.fm6((d_505_v18_).f13, (self).f8, (self).f6, (d_505_v18_).f13, globalState), (self).f5, (self).f8])
            d_507_v20_ = ((d_508_v21_)[default__.safeIndex((d_510_v23_).cardinality, len(d_508_v21_))] if (self).f6 else d_507_v20_)
            d_511_v24_: _dafny.Seq
            d_511_v24_ = _dafny.SeqWithoutIsStrInference([d_502_v15_])
            d_512_v25_: _dafny.Seq
            d_512_v25_ = _dafny.SeqWithoutIsStrInference([d_511_v24_])
            d_513_v26_: _dafny.Map
            d_513_v26_ = _dafny.Map({(self).f8: not((self).f6)})
            d_514_v27_: _dafny.Seq
            d_514_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm17(((d_513_v26_)[p0] if (p0) in (d_513_v26_) else (self).f6), globalState)])
            def iife56_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(119, 115):
                    d_515_v28_: int = compr_14_
                    if ((119) <= (d_515_v28_)) and ((d_515_v28_) < (115)):
                        coll14_[default__.safeDivisionInt(d_515_v28_, (d_505_v18_).f13)] = not((self).f6)
                return _dafny.Map(coll14_)
            def iife57_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(223, 441):
                    d_516_v29_: int = compr_15_
                    if ((223) <= (d_516_v29_)) and ((d_516_v29_) < (441)):
                        coll15_[(d_516_v29_) + ((d_505_v18_).f13)] = True
                return _dafny.Map(coll15_)
            d_512_v25_ = (d_514_v27_)[default__.safeIndex(len((iife56_()
            ) | (iife57_()
            )), len(d_514_v27_))]
            d_517_v30_: _dafny.Array
            nw75_ = _dafny.Array(None, 29)
            d_517_v30_ = nw75_
            d_518_v31_: _dafny.Set
            d_518_v31_ = _dafny.Set({d_502_v15_, d_502_v15_, d_502_v15_, d_502_v15_})
            d_519_v32_: C2
            nw76_ = C2()
            nw76_.ctor__(len(d_518_v31_))
            d_519_v32_ = nw76_
            index75_ = default__.safeIndex(190, (d_517_v30_).length(0))
            (d_517_v30_)[index75_] = d_519_v32_
            index76_ = default__.safeIndex(190, (d_517_v30_).length(0))
            (d_517_v30_)[index76_] = d_519_v32_
            d_520_v33_: bool
            d_520_v33_ = True
            d_520_v33_ = (self).f6
            d_502_v15_ = _dafny.CodePoint('r')
        elif True:
            d_521_v34_: _dafny.Map
            d_521_v34_ = _dafny.Map({not((self).f6): (self).f6})
            d_522_v35_: _dafny.Seq
            d_522_v35_ = _dafny.SeqWithoutIsStrInference([d_521_v34_, d_521_v34_])
            d_523_v36_: _dafny.Map
            d_523_v36_ = _dafny.Map({(self).f6: len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))})
            d_524_v37_: C1
            nw77_ = C1()
            nw77_.ctor__(len((_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f6: (self).f6}) for d_525_i2_ in range(default__.abs(274))])) + (d_522_v35_)), (self).f5, len(d_523_v36_), ((self).f6 if (self).f6 else (self).f6))
            d_524_v37_ = nw77_
            d_526_v38_: _dafny.Array
            nw78_ = _dafny.Array(_dafny.Seq({}), 10)
            d_526_v38_ = nw78_
            d_527_v39_: _dafny.Seq
            d_527_v39_ = _dafny.SeqWithoutIsStrInference([True])
            d_528_v40_: _dafny.Map
            d_528_v40_ = _dafny.Map({(self).f8: d_527_v39_})
            d_529_v41_: D6
            d_529_v41_ = D6_DC17(d_527_v39_)
            nw79_ = _dafny.Array(None, 28)
            nw79_[int(0)] = ((d_527_v39_) + (d_527_v39_)).set(default__.safeIndex(len(((d_528_v40_)[d_524_v37_.f10] if (d_524_v37_.f10) in (d_528_v40_) else d_527_v39_)), len((d_527_v39_) + (d_527_v39_))), (self).f6)
            nw79_[int(1)] = d_527_v39_
            nw79_[int(2)] = (d_527_v39_) + (d_527_v39_)
            nw79_[int(3)] = (D6_DC17(d_527_v39_)).cf35
            nw79_[int(4)] = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
            nw79_[int(5)] = default__.fm16(globalState)
            nw79_[int(6)] = default__.fm16(globalState)
            nw79_[int(7)] = (d_527_v39_) + (d_527_v39_)
            nw79_[int(8)] = d_527_v39_
            nw79_[int(9)] = d_527_v39_
            nw79_[int(10)] = (d_527_v39_) + (d_527_v39_)
            nw79_[int(11)] = (default__.fm18(globalState)).cf35
            nw79_[int(12)] = _dafny.SeqWithoutIsStrInference([True])
            nw79_[int(13)] = d_527_v39_
            nw79_[int(14)] = d_527_v39_
            nw79_[int(15)] = (d_527_v39_) + ((d_529_v41_).cf35)
            nw79_[int(16)] = (d_527_v39_) + (d_527_v39_)
            nw79_[int(17)] = d_527_v39_
            nw79_[int(18)] = _dafny.SeqWithoutIsStrInference([(self).f6])
            nw79_[int(19)] = _dafny.SeqWithoutIsStrInference([(self).f6])
            nw79_[int(20)] = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
            nw79_[int(21)] = d_527_v39_
            nw79_[int(22)] = ((d_527_v39_) + (d_527_v39_)).set(default__.safeIndex(d_524_v37_.f10, len((d_527_v39_) + (d_527_v39_))), False)
            nw79_[int(23)] = d_527_v39_
            nw79_[int(24)] = (d_527_v39_) + (d_527_v39_)
            nw79_[int(25)] = d_527_v39_
            nw79_[int(26)] = _dafny.SeqWithoutIsStrInference([True])
            nw79_[int(27)] = d_527_v39_
            d_526_v38_ = nw79_
            d_530_v42_: _dafny.Array
            def lambda15_(d_531_i3_):
                return (self).f6

            init8_ = lambda15_
            nw80_ = _dafny.Array(None, 18)
            for i0_8_ in range(nw80_.length(0)):
                nw80_[i0_8_] = init8_(i0_8_)
            d_530_v42_ = nw80_
            d_532_v43_: _dafny.Seq
            d_532_v43_ = _dafny.SeqWithoutIsStrInference([(self).f5])
            d_533_v44_: _dafny.Seq
            d_533_v44_ = _dafny.SeqWithoutIsStrInference([len(d_532_v43_)])
            index77_ = default__.safeIndex(947, (d_530_v42_).length(0))
            (d_530_v42_)[index77_] = (d_533_v44_) <= (d_532_v43_)
            d_534_v45_: _dafny.Array
            nw81_ = _dafny.Array(int(0), 11)
            d_534_v45_ = nw81_
            d_535_v46_: D3
            d_535_v46_ = D3_DC11(d_534_v45_)
            index78_ = default__.safeIndex(947, (d_530_v42_).length(0))
            rhs49_ = False
            rhs50_ = d_535_v46_
            lhs37_ = d_530_v42_
            lhs38_ = default__.safeIndex(947, (d_530_v42_).length(0))
            lhs37_[lhs38_] = rhs49_
            d_535_v46_ = rhs50_
            d_536_v47_: str
            d_536_v47_ = _dafny.CodePoint('n')
            d_536_v47_ = d_536_v47_
            d_537_v48_: _dafny.MultiSet
            d_537_v48_ = _dafny.MultiSet([622])
            d_538_v49_: _dafny.Seq
            d_538_v49_ = _dafny.SeqWithoutIsStrInference([d_537_v48_])
            d_539_v50_: _dafny.Seq
            d_539_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmwqpjm"))
            d_540_v51_: _dafny.Seq
            d_540_v51_ = _dafny.SeqWithoutIsStrInference([d_539_v50_])
            d_541_v52_: _dafny.Map
            d_541_v52_ = _dafny.Map({(d_538_v49_)[default__.safeIndex((self).f5, len(d_538_v49_))]: (d_540_v51_)[default__.safeIndex((d_524_v37_).f11, len(d_540_v51_))]})
            index79_ = default__.safeIndex(947, (d_530_v42_).length(0))
            (d_530_v42_)[index79_] = (not((d_524_v37_.f10) > (len(d_541_v52_))) if (d_524_v37_).fm0((self).f6, globalState) else (d_537_v48_).ispropersubset(d_537_v48_))
        d_542_v53_: _dafny.Seq
        d_542_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whxnlgjoh"))
        hi4_ = (0) - (len(d_542_v53_))
        for d_543_i4_ in range(((self).f5) - (default__.fm6(p0, (self).f5, (self).f6, p0, globalState)), hi4_):
            d_544_v54_: int
            d_544_v54_ = 352
            d_545_v55_: _dafny.Seq
            d_545_v55_ = _dafny.SeqWithoutIsStrInference([False])
            d_546_v56_: _dafny.Set
            d_546_v56_ = _dafny.Set({(self).f6})
            d_547_v57_: _dafny.Seq
            d_547_v57_ = _dafny.SeqWithoutIsStrInference([len(d_545_v55_), len(d_542_v53_), len(d_546_v56_)])
            d_544_v54_ = default__.safeDivisionInt(len((d_547_v57_) + (d_547_v57_)), d_544_v54_)
            d_548_v58_: bool
            d_548_v58_ = True
            d_549_v59_: _dafny.MultiSet
            d_549_v59_ = _dafny.MultiSet([len(d_542_v53_)])
            d_550_v60_: _dafny.Map
            d_550_v60_ = _dafny.Map({d_549_v59_: (self).f6})
            d_548_v58_ = (not((len(d_550_v60_)) >= (924)) if (self).f6 else (d_548_v58_) == ((self).f6))
            d_551_v61_: str
            d_551_v61_ = _dafny.CodePoint('t')
            d_552_v62_: _dafny.Array
            nw82_ = _dafny.Array(int(0), 15)
            d_552_v62_ = nw82_
            d_553_v63_: D6
            d_553_v63_ = D6_DC18((self).f6)
            rhs51_ = (self).fm0(((d_553_v63_).cf36) == ((self).f6), globalState)
            rhs52_ = p0
            rhs53_ = d_551_v61_
            rhs54_ = d_552_v62_
            rhs55_ = p0
            d_548_v58_ = rhs51_
            d_544_v54_ = rhs52_
            d_551_v61_ = rhs53_
            d_552_v62_ = rhs54_
            d_544_v54_ = rhs55_
            d_548_v58_ = (self).f6
        d_554_v64_: str
        d_554_v64_ = _dafny.CodePoint('o')
        d_555_v65_: _dafny.Array
        nw83_ = _dafny.Array(None, 4)
        nw83_[int(0)] = d_554_v64_
        nw83_[int(1)] = d_554_v64_
        nw83_[int(2)] = d_554_v64_
        nw83_[int(3)] = d_554_v64_
        d_555_v65_ = nw83_
        d_556_v66_: _dafny.Seq
        d_556_v66_ = _dafny.SeqWithoutIsStrInference([d_555_v65_])
        d_557_v67_: _dafny.MultiSet
        d_557_v67_ = _dafny.MultiSet([(self).f6, True])
        d_558_v68_: D5
        d_558_v68_ = D5_DC15(d_557_v67_)
        d_559_v69_: _dafny.Array
        def lambda16_(d_560_i5_):
            return (_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])) + (_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))

        init9_ = lambda16_
        nw84_ = _dafny.Array(None, 7)
        for i0_9_ in range(nw84_.length(0)):
            nw84_[i0_9_] = init9_(i0_9_)
        d_559_v69_ = nw84_
        d_561_v70_: _dafny.Seq
        d_561_v70_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
        index80_ = default__.safeIndex(417, (d_559_v69_).length(0))
        (d_559_v69_)[index80_] = (d_561_v70_) + (d_561_v70_)
        index81_ = default__.safeIndex(680, (d_559_v69_).length(0))
        (d_559_v69_)[index81_] = d_561_v70_
        d_562_v71_: bool
        d_562_v71_ = False
        d_563_v72_: D1
        d_563_v72_ = D1_DC7(D1_DC4(d_562_v71_, d_562_v71_, _dafny.SeqWithoutIsStrInference([d_554_v64_ for d_564_i6_ in range(default__.abs(508))]), p0, default__.fm6((self).f8, -992, d_562_v71_, 918, globalState)))
        d_565_v73_: _dafny.Map
        d_565_v73_ = _dafny.Map({d_563_v72_: d_557_v67_})
        pat_let_tv15_ = d_565_v73_
        pat_let_tv16_ = d_562_v71_
        pat_let_tv17_ = globalState
        pat_let_tv18_ = d_562_v71_
        pat_let_tv19_ = globalState
        pat_let_tv20_ = d_565_v73_
        pat_let_tv21_ = d_557_v67_
        index82_ = default__.safeIndex(417, (d_559_v69_).length(0))
        index83_ = default__.safeIndex(680, (d_559_v69_).length(0))
        def iife58_(_pat_let21_0):
            def iife59_(d_566_dt__update__tmp_h1_):
                def iife60_(_pat_let22_0):
                    def iife61_(d_567_dt__update_hcf31_h0_):
                        return D5_DC15(d_567_dt__update_hcf31_h0_)
                    return iife61_(_pat_let22_0)
                return iife60_(((pat_let_tv15_)[D1_DC7(default__.fm14(pat_let_tv16_, pat_let_tv17_))] if (D1_DC7(default__.fm14(pat_let_tv18_, pat_let_tv19_))) in (pat_let_tv20_) else pat_let_tv21_))
            return iife59_(_pat_let21_0)
        rhs56_ = d_556_v66_
        rhs57_ = iife58_(D5_DC15(d_557_v67_))
        rhs58_ = (d_561_v70_) + (d_561_v70_)
        rhs59_ = default__.fm16(globalState)
        rhs60_ = d_562_v71_
        lhs39_ = d_559_v69_
        lhs40_ = default__.safeIndex(417, (d_559_v69_).length(0))
        lhs41_ = d_559_v69_
        lhs42_ = default__.safeIndex(680, (d_559_v69_).length(0))
        d_556_v66_ = rhs56_
        d_558_v68_ = rhs57_
        lhs39_[lhs40_] = rhs58_
        lhs41_[lhs42_] = rhs59_
        d_562_v71_ = rhs60_
        d_568_v74_: _dafny.Array
        def lambda17_(d_569_p0_):
            def lambda18_(d_570_i8_):
                return (d_570_i8_) - (d_569_p0_)

            return lambda18_

        init10_ = lambda17_(p0)
        nw85_ = _dafny.Array(None, 25)
        for i0_10_ in range(nw85_.length(0)):
            nw85_[i0_10_] = init10_(i0_10_)
        d_568_v74_ = nw85_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_568_v74_).length(0)):
            d_571_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_571_i7_)) and ((d_571_i7_) < ((d_568_v74_).length(0)))):
                (d_568_v74_)[(d_571_i7_)] = default__.safeDivisionInt(d_571_i7_, p0)
        d_572_v75_: _dafny.Set
        d_572_v75_ = _dafny.Set({d_562_v71_, (self).f6})
        d_573_v76_: _dafny.Seq
        d_573_v76_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_572_v75_))])
        index84_ = default__.safeIndex(399, (d_568_v74_).length(0))
        (d_568_v74_)[index84_] = (d_573_v76_)[default__.safeIndex(len(d_542_v53_), len(d_573_v76_))]
        index85_ = default__.safeIndex(399, (d_568_v74_).length(0))
        (d_568_v74_)[index85_] = ((self).f5) + (len((d_572_v75_) - (d_572_v75_)))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: D0 = D0.default()()
        d_574_v0_: int
        d_574_v0_ = -393
        d_575_v1_: _dafny.Set
        d_575_v1_ = _dafny.Set({p2, 430})
        d_576_v2_: _dafny.MultiSet
        d_576_v2_ = _dafny.MultiSet([len(d_575_v1_), d_574_v0_])
        d_577_v3_: T0
        nw86_ = C1()
        nw86_.ctor__((d_576_v2_).cardinality, (self).f8, d_574_v0_, (self).f6)
        d_577_v3_ = nw86_
        d_578_v4_: D5
        d_578_v4_ = D5_DC16(d_577_v3_, (self).f6, p0)
        d_579_v5_: T0
        nw87_ = C1()
        nw87_.ctor__(-929, (d_578_v4_).cf34, p0, (d_577_v3_).f6)
        d_579_v5_ = nw87_
        d_580_v6_: _dafny.Seq
        d_580_v6_ = _dafny.SeqWithoutIsStrInference([len(p1), 968])
        d_581_v7_: _dafny.Set
        d_581_v7_ = _dafny.Set({(self).f6, (self).f6, (d_579_v5_).f6, (self).f6})
        d_582_v8_: _dafny.Seq
        d_582_v8_ = _dafny.SeqWithoutIsStrInference([(d_579_v5_).f6, True])
        d_583_v9_: _dafny.Map
        d_583_v9_ = _dafny.Map({len(d_581_v7_): len(d_582_v8_)})
        d_584_v10_: _dafny.Map
        d_584_v10_ = _dafny.Map({11: (d_577_v3_).f6})
        d_585_v11_: _dafny.Map
        d_585_v11_ = _dafny.Map({(d_579_v5_).f6: ((d_584_v10_)[(d_579_v5_).f5] if ((d_579_v5_).f5) in (d_584_v10_) else (d_579_v5_).f6)})
        d_586_v12_: _dafny.Array
        nw88_ = _dafny.Array(None, 9)
        nw88_[int(0)] = p0
        nw88_[int(1)] = len(d_580_v6_)
        nw88_[int(2)] = d_574_v0_
        nw88_[int(3)] = len((d_583_v9_) | (d_583_v9_))
        nw88_[int(4)] = (665) + (780)
        nw88_[int(5)] = (d_579_v5_).f5
        nw88_[int(6)] = default__.safeDivisionInt((d_577_v3_).f5, (self).f5)
        nw88_[int(7)] = (d_577_v3_).f5
        nw88_[int(8)] = default__.fm6(p0, len(d_585_v11_), (d_579_v5_).f6, default__.fm6(default__.fm6((d_577_v3_).f5, (self).f8, (d_577_v3_).f6, (d_579_v5_).f5, globalState), d_574_v0_, p3, p0, globalState), globalState)
        d_586_v12_ = nw88_
        index86_ = default__.safeIndex(109, (d_586_v12_).length(0))
        (d_586_v12_)[index86_] = default__.safeModuloInt((self).f8, (d_577_v3_).f5)
        d_587_v13_: D1
        d_587_v13_ = D1_DC6(((d_577_v3_).f6) or ((self).f6), (0) - ((d_577_v3_).f5), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))), (d_579_v5_).f6)
        d_588_v14_: _dafny.MultiSet
        d_588_v14_ = _dafny.MultiSet([not((d_577_v3_).f6)])
        d_589_v15_: _dafny.Map
        d_589_v15_ = _dafny.Map({not((d_577_v3_).f6): d_579_v5_})
        d_590_v16_: _dafny.Map
        d_590_v16_ = _dafny.Map({p3: d_580_v6_})
        index87_ = default__.safeIndex(109, (d_586_v12_).length(0))
        rhs61_ = default__.fm6(len(default__.fm19((d_588_v14_).cardinality, globalState)), (0) - ((len(p1)) * (p0)), p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "doyklvo"))), globalState)
        rhs62_ = ((d_589_v15_)[(self).f6] if ((self).f6) in (d_589_v15_) else d_577_v3_)
        rhs63_ = default__.fm6((d_577_v3_).f5, len((d_580_v6_) + (((d_590_v16_)[(d_577_v3_).f6] if ((d_577_v3_).f6) in (d_590_v16_) else d_580_v6_))), (d_577_v3_).fm0(p3, globalState), default__.safeModuloInt(d_574_v0_, p2), globalState)
        rhs64_ = (d_587_v13_ if (self).fm0(p3, globalState) else default__.fm20((0) - (len(p1)), d_588_v14_, (self).f6, (self).f6, globalState))
        rhs65_ = d_577_v3_
        lhs43_ = d_586_v12_
        lhs44_ = default__.safeIndex(109, (d_586_v12_).length(0))
        d_574_v0_ = rhs61_
        d_579_v5_ = rhs62_
        lhs43_[lhs44_] = rhs63_
        d_587_v13_ = rhs64_
        d_579_v5_ = rhs65_
        if ((d_585_v11_)[(d_577_v3_).f6] if ((d_577_v3_).f6) in (d_585_v11_) else (self).fm0((d_579_v5_).f6, globalState)):
            d_591_v17_: C0
            nw89_ = C0()
            nw89_.ctor__(p2, len(d_581_v7_))
            d_591_v17_ = nw89_
            if (d_588_v14_).issubset(d_588_v14_):
                index88_ = default__.safeIndex(109, (d_586_v12_).length(0))
                (d_586_v12_)[index88_] = (d_591_v17_).f12
                d_574_v0_ = default__.fm6((d_574_v0_) * ((d_579_v5_).f5), (0) - (p2), (d_577_v3_).f6, ((d_591_v17_).f13) * (p2), globalState)
                d_592_v18_: _dafny.Map
                d_592_v18_ = _dafny.Map({(d_591_v17_).f13: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_593_i0_ in range(default__.abs(737))])})
                d_594_v19_: _dafny.Array
                nw90_ = _dafny.Array(None, 17)
                nw90_[int(0)] = ((d_592_v18_)[(d_591_v17_).f12] if ((d_591_v17_).f12) in (d_592_v18_) else p1)
                nw90_[int(1)] = p1
                nw90_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvmci"))
                nw90_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_595_i1_ in range(default__.abs(962))])
                nw90_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_596_i2_ in range(default__.abs(908))])
                nw90_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ea"))
                nw90_[int(6)] = p1
                nw90_[int(7)] = p1
                nw90_[int(8)] = p1
                nw90_[int(9)] = p1
                nw90_[int(10)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_597_i3_ in range(default__.abs(852))])
                nw90_[int(11)] = p1
                nw90_[int(12)] = p1
                nw90_[int(13)] = p1
                nw90_[int(14)] = p1
                nw90_[int(15)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_598_i4_ in range(default__.abs(519))])
                nw90_[int(16)] = p1
                d_594_v19_ = nw90_
                d_599_v20_: _dafny.Map
                d_599_v20_ = _dafny.Map({(d_579_v5_).f5: d_594_v19_})
                d_599_v20_ = (d_599_v20_).set((0) - ((self).f8), d_594_v19_)
                index89_ = default__.safeIndex(980, (d_594_v19_).length(0))
                (d_594_v19_)[index89_] = p1
                d_600_v21_: _dafny.Array
                nw91_ = _dafny.Array(None, 28)
                d_600_v21_ = nw91_
                d_601_v22_: str
                d_601_v22_ = _dafny.CodePoint('g')
                d_602_v23_: str
                d_602_v23_ = d_601_v22_
                d_603_v24_: _dafny.Array
                nw92_ = _dafny.Array(None, 22)
                nw92_[int(0)] = (d_577_v3_).f6
                nw92_[int(1)] = (d_577_v3_).f6
                nw92_[int(2)] = (self).f6
                nw92_[int(3)] = not ((d_577_v3_).f6) or (not(((d_585_v11_)[(d_577_v3_).f6] if ((d_577_v3_).f6) in (d_585_v11_) else (d_577_v3_).f6)))
                nw92_[int(4)] = (d_577_v3_).f6
                nw92_[int(5)] = ((0) - ((d_579_v5_).f5)) != ((d_591_v17_).f12)
                nw92_[int(6)] = (not((d_577_v3_).f6) if (d_577_v3_).f6 else (d_577_v3_).f6)
                nw92_[int(7)] = (d_577_v3_).f6
                nw92_[int(8)] = not ((d_579_v5_).f6) or ((d_577_v3_).f6)
                nw92_[int(9)] = (d_577_v3_).f6
                nw92_[int(10)] = True
                nw92_[int(11)] = (d_577_v3_).f6
                nw92_[int(12)] = (d_577_v3_).f6
                nw92_[int(13)] = (d_577_v3_).f6
                nw92_[int(14)] = ((d_579_v5_).f6) and ((self).f6)
                nw92_[int(15)] = (d_577_v3_).f6
                nw92_[int(16)] = (d_579_v5_).fm0((d_579_v5_).f6, globalState)
                nw92_[int(17)] = (d_577_v3_).f6
                nw92_[int(18)] = (d_577_v3_).fm0((d_579_v5_).f6, globalState)
                nw92_[int(19)] = p3
                nw92_[int(20)] = (d_579_v5_).f6
                nw92_[int(21)] = not(((d_585_v11_)[(d_591_v17_).fm8(p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjkgdpeu")), (d_602_v23_), globalState)] if ((d_591_v17_).fm8(p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjkgdpeu")), (d_602_v23_), globalState)) in (d_585_v11_) else (d_577_v3_).f6))
                d_603_v24_ = nw92_
                d_604_v25_: _dafny.Map
                d_604_v25_ = _dafny.Map({129: d_603_v24_})
                index90_ = default__.safeIndex(980, (d_594_v19_).length(0))
                rhs66_ = (p1) + (p1)
                rhs67_ = 481
                rhs68_ = d_600_v21_
                rhs69_ = ((d_604_v25_)[d_574_v0_] if (d_574_v0_) in (d_604_v25_) else d_603_v24_)
                lhs45_ = d_594_v19_
                lhs46_ = default__.safeIndex(980, (d_594_v19_).length(0))
                lhs45_[lhs46_] = rhs66_
                d_574_v0_ = rhs67_
                d_600_v21_ = rhs68_
                d_603_v24_ = rhs69_
                d_605_v26_: D6
                d_605_v26_ = D6_DC17(d_582_v8_)
                pat_let_tv22_ = d_582_v8_
                pat_let_tv23_ = d_582_v8_
                def iife62_(_pat_let23_0):
                    def iife63_(d_606_dt__update__tmp_h0_):
                        def iife64_(_pat_let24_0):
                            def iife65_(d_607_dt__update_hcf35_h0_):
                                return D6_DC17(d_607_dt__update_hcf35_h0_)
                            return iife65_(_pat_let24_0)
                        return iife64_((pat_let_tv22_) + (pat_let_tv23_))
                    return iife63_(_pat_let23_0)
                d_605_v26_ = iife62_(d_605_v26_)
            elif True:
                d_608_v27_: D1
                d_608_v27_ = D1_DC4((d_579_v5_).f6, p3, (p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahcbwma"))), (d_579_v5_).f5, default__.safeDivisionInt(len(d_584_v10_), (d_591_v17_).f12))
                d_608_v27_ = d_608_v27_
                d_609_v29_: _dafny.Seq
                def iife66_():
                    coll16_ = _dafny.Set()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(626, 843):
                        d_610_v28_: int = compr_16_
                        if ((626) <= (d_610_v28_)) and ((d_610_v28_) < (843)):
                            coll16_ = coll16_.union(_dafny.Set([(d_610_v28_) + ((d_579_v5_).f5)]))
                    return _dafny.Set(coll16_)
                d_609_v29_ = _dafny.SeqWithoutIsStrInference([d_575_v1_, iife66_()
                , d_575_v1_])
                d_611_v31_: _dafny.Array
                nw93_ = _dafny.Array(None, 20)
                nw93_[int(0)] = (d_575_v1_) - (d_575_v1_)
                nw93_[int(1)] = d_575_v1_
                nw93_[int(2)] = d_575_v1_
                nw93_[int(3)] = default__.fm21(globalState)
                nw93_[int(4)] = d_575_v1_
                nw93_[int(5)] = (d_609_v29_)[default__.safeIndex(len(p1), len(d_609_v29_))]
                nw93_[int(6)] = (_dafny.Set({(d_577_v3_).f5})) | (d_575_v1_)
                nw93_[int(7)] = (d_575_v1_) - (_dafny.Set({-219}))
                nw93_[int(8)] = d_575_v1_
                nw93_[int(9)] = d_575_v1_
                nw93_[int(10)] = d_575_v1_
                def iife67_():
                    coll17_ = _dafny.Set()
                    compr_17_: int
                    for compr_17_ in _dafny.IntegerRange(-558, -764):
                        d_612_v30_: int = compr_17_
                        if ((-558) <= (d_612_v30_)) and ((d_612_v30_) < (-764)):
                            coll17_ = coll17_.union(_dafny.Set([default__.safeDivisionInt(d_612_v30_, (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))])]))
                    return _dafny.Set(coll17_)
                nw93_[int(11)] = iife67_()
                
                nw93_[int(12)] = (d_575_v1_) | (d_575_v1_)
                nw93_[int(13)] = d_575_v1_
                nw93_[int(14)] = d_575_v1_
                nw93_[int(15)] = d_575_v1_
                nw93_[int(16)] = d_575_v1_
                nw93_[int(17)] = d_575_v1_
                nw93_[int(18)] = d_575_v1_
                nw93_[int(19)] = d_575_v1_
                d_611_v31_ = nw93_
                index91_ = default__.safeIndex(347, (d_611_v31_).length(0))
                (d_611_v31_)[index91_] = d_575_v1_
                index92_ = default__.safeIndex(347, (d_611_v31_).length(0))
                (d_611_v31_)[index92_] = (d_575_v1_).intersection((d_575_v1_) | (d_575_v1_))
                d_613_v32_: _dafny.Array
                def lambda19_(d_614_v3_):
                    def lambda20_(d_615_i5_):
                        return (d_614_v3_).f6

                    return lambda20_

                init11_ = lambda19_(d_577_v3_)
                nw94_ = _dafny.Array(None, 5)
                for i0_11_ in range(nw94_.length(0)):
                    nw94_[i0_11_] = init11_(i0_11_)
                d_613_v32_ = nw94_
                index93_ = default__.safeIndex(721, (d_613_v32_).length(0))
                (d_613_v32_)[index93_] = (d_577_v3_).f6
                index94_ = default__.safeIndex(721, (d_613_v32_).length(0))
                (d_613_v32_)[index94_] = not (False) or ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_616_i6_ in range(default__.abs(-432))])) != (p1))
                index95_ = default__.safeIndex(721, (d_613_v32_).length(0))
                index96_ = default__.safeIndex(109, (d_586_v12_).length(0))
                rhs70_ = not (False) or ((d_613_v32_)[default__.safeIndex(721, (d_613_v32_).length(0))])
                rhs71_ = (default__.safeModuloInt(p0, 672)) + ((d_577_v3_).f5)
                lhs47_ = d_613_v32_
                lhs48_ = default__.safeIndex(721, (d_613_v32_).length(0))
                lhs49_ = d_586_v12_
                lhs50_ = default__.safeIndex(109, (d_586_v12_).length(0))
                lhs47_[lhs48_] = rhs70_
                lhs49_[lhs50_] = rhs71_
                d_617_v33_: D3
                d_617_v33_ = D3_DC11(d_586_v12_)
                d_617_v33_ = d_617_v33_
            d_618_v34_: _dafny.Array
            nw95_ = _dafny.Array(None, 9)
            nw95_[int(0)] = (d_577_v3_).f6
            nw95_[int(1)] = (d_577_v3_).f6
            nw95_[int(2)] = p3
            nw95_[int(3)] = p3
            nw95_[int(4)] = (d_577_v3_).f6
            nw95_[int(5)] = (d_577_v3_).f6
            nw95_[int(6)] = (d_577_v3_).f6
            nw95_[int(7)] = False
            nw95_[int(8)] = (d_579_v5_).f6
            d_618_v34_ = nw95_
            index97_ = default__.safeIndex(913, (d_618_v34_).length(0))
            (d_618_v34_)[index97_] = (d_579_v5_).f6
            index98_ = default__.safeIndex(913, (d_618_v34_).length(0))
            (d_618_v34_)[index98_] = p3
            index99_ = default__.safeIndex(913, (d_618_v34_).length(0))
            (d_618_v34_)[index99_] = (d_579_v5_).f6
            index100_ = default__.safeIndex(109, (d_586_v12_).length(0))
            (d_586_v12_)[index100_] = (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))]
        elif True:
            d_619_v35_: _dafny.Map
            d_619_v35_ = _dafny.Map({d_580_v6_: (d_577_v3_).f6})
            d_620_v36_: _dafny.Array
            nw96_ = _dafny.Array(None, 23)
            nw96_[int(0)] = True
            nw96_[int(1)] = (d_579_v5_).f6
            nw96_[int(2)] = (d_577_v3_).f6
            nw96_[int(3)] = (d_575_v1_).isdisjoint(d_575_v1_)
            nw96_[int(4)] = (d_579_v5_).f6
            nw96_[int(5)] = True
            nw96_[int(6)] = (d_577_v3_).f6
            nw96_[int(7)] = (d_579_v5_).f6
            nw96_[int(8)] = (d_577_v3_).f6
            nw96_[int(9)] = p3
            nw96_[int(10)] = (d_579_v5_).fm0((d_577_v3_).f6, globalState)
            nw96_[int(11)] = ((self).f8) < (d_574_v0_)
            nw96_[int(12)] = ((d_577_v3_).f6 if p3 else p3)
            nw96_[int(13)] = (d_579_v5_).f6
            nw96_[int(14)] = (d_579_v5_).f6
            nw96_[int(15)] = (self).f6
            nw96_[int(16)] = ((d_619_v35_)[_dafny.SeqWithoutIsStrInference([p2, (d_579_v5_).f5, (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))], (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))]])] if (_dafny.SeqWithoutIsStrInference([p2, (d_579_v5_).f5, (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))], (d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))]])) in (d_619_v35_) else (self).f6)
            nw96_[int(17)] = True
            nw96_[int(18)] = (d_577_v3_).f6
            nw96_[int(19)] = (self).f6
            nw96_[int(20)] = (self).f6
            nw96_[int(21)] = not ((d_579_v5_).f6) or (p3)
            nw96_[int(22)] = not ((d_582_v8_)[default__.safeIndex((0) - ((_dafny.MultiSet([(self).fm0((d_577_v3_).f6, globalState), (self).f6])).cardinality), len(d_582_v8_))]) or ((d_577_v3_).f6)
            d_620_v36_ = nw96_
            index101_ = default__.safeIndex(375, (d_620_v36_).length(0))
            (d_620_v36_)[index101_] = p3
            index102_ = default__.safeIndex(375, (d_620_v36_).length(0))
            (d_620_v36_)[index102_] = p3
            d_621_v37_: _dafny.Map
            d_621_v37_ = _dafny.Map({p2: d_620_v36_})
            d_621_v37_ = (d_621_v37_).set((d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))], d_620_v36_)
            d_582_v8_ = default__.fm16(globalState)
            index103_ = default__.safeIndex(375, (d_620_v36_).length(0))
            (d_620_v36_)[index103_] = (d_579_v5_).f6
            d_622_v38_: _dafny.Seq
            d_622_v38_ = _dafny.SeqWithoutIsStrInference([p1])
            d_623_v39_: D1
            d_623_v39_ = D1_DC4((p0) > (d_574_v0_), (d_577_v3_).fm0(p3, globalState), p1, ((d_577_v3_).f5) - (len((d_622_v38_)[default__.safeIndex((0) - ((d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))]), len(d_622_v38_))])), (d_577_v3_).f5)
            source8_ = d_623_v39_
            if source8_.is_DC4:
                d_624___mcc_h0_ = source8_.cf8
                d_625___mcc_h1_ = source8_.cf9
                d_626___mcc_h2_ = source8_.cf10
                d_627___mcc_h3_ = source8_.cf11
                d_628___mcc_h4_ = source8_.cf12
                d_629_cf12_ = d_628___mcc_h4_
                d_630_cf11_ = d_627___mcc_h3_
                d_631_cf10_ = d_626___mcc_h2_
                d_632_cf9_ = d_625___mcc_h1_
                d_633_cf8_ = d_624___mcc_h0_
                d_634_v40_: _dafny.Map
                d_634_v40_ = _dafny.Map({d_632_cf9_: (d_577_v3_).f5})
                d_635_v41_: _dafny.Map
                d_635_v41_ = _dafny.Map({d_585_v11_: ((0) - (default__.fm6(len(d_634_v40_), (d_577_v3_).f5, p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htbyp"))), globalState))) != ((d_577_v3_).f5)})
                d_635_v41_ = d_635_v41_
                d_636_v42_: str
                d_636_v42_ = _dafny.CodePoint('e')
                d_636_v42_ = d_636_v42_
                d_637_v43_: _dafny.Array
                nw97_ = _dafny.Array(None, 29)
                d_637_v43_ = nw97_
                index104_ = default__.safeIndex(199, (d_637_v43_).length(0))
                (d_637_v43_)[index104_] = d_579_v5_
                index105_ = default__.safeIndex(199, (d_637_v43_).length(0))
                (d_637_v43_)[index105_] = ((d_589_v15_)[p3] if (p3) in (d_589_v15_) else d_579_v5_)
                d_633_cf8_ = (d_577_v3_).f6
            elif source8_.is_DC5:
                d_638___mcc_h5_ = source8_.cf13
                d_639___mcc_h6_ = source8_.cf14
                d_640___mcc_h7_ = source8_.cf15
                d_641___mcc_h8_ = source8_.cf16
                d_642_cf16_ = d_641___mcc_h8_
                d_643_cf15_ = d_640___mcc_h7_
                d_644_cf14_ = d_639___mcc_h6_
                d_645_cf13_ = d_638___mcc_h5_
                d_643_cf15_ = (d_579_v5_).f5
                d_646_v44_: str
                d_646_v44_ = _dafny.CodePoint('e')
                rhs72_ = d_586_v12_
                rhs73_ = (D1_DC4((d_579_v5_).f6, ((d_585_v11_)[False] if (False) in (d_585_v11_) else p3), p1, p0, (d_579_v5_).f5)).cf8
                rhs74_ = (d_646_v44_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhrqm")))
                d_586_v12_ = rhs72_
                d_644_cf14_ = rhs73_
                d_645_cf13_ = rhs74_
                d_642_cf16_ = len(_dafny.SeqWithoutIsStrInference([(d_577_v3_).f5 for d_647_i7_ in range(default__.abs(138))]))
                d_644_cf14_ = (d_579_v5_).f6
            elif source8_.is_DC6:
                d_648___mcc_h9_ = source8_.cf17
                d_649___mcc_h10_ = source8_.cf18
                d_650___mcc_h11_ = source8_.cf19
                d_651___mcc_h12_ = source8_.cf20
                d_652_cf20_ = d_651___mcc_h12_
                d_653_cf19_ = d_650___mcc_h11_
                d_654_cf18_ = d_649___mcc_h10_
                d_655_cf17_ = d_648___mcc_h9_
                d_656_v45_: C2
                nw98_ = C2()
                nw98_.ctor__((d_577_v3_).f5)
                d_656_v45_ = nw98_
                d_657_v46_: _dafny.Map
                d_657_v46_ = _dafny.Map({d_586_v12_: d_656_v45_})
                d_658_v47_: _dafny.Seq
                d_658_v47_ = _dafny.SeqWithoutIsStrInference([d_586_v12_])
                d_659_v48_: _dafny.Map
                d_659_v48_ = _dafny.Map({((d_657_v46_)[(d_658_v47_)[default__.safeIndex((d_577_v3_).f5, len(d_658_v47_))]] if ((d_658_v47_)[default__.safeIndex((d_577_v3_).f5, len(d_658_v47_))]) in (d_657_v46_) else d_656_v45_): _dafny.Set({(d_577_v3_).f5, p2})})
                d_659_v48_ = d_659_v48_
                d_620_v36_ = d_620_v36_
                d_660_v49_: str
                d_660_v49_ = _dafny.CodePoint('k')
                d_652_cf20_ = not((d_656_v45_).fm1(p0, d_660_v49_, globalState))
                d_661_v50_: _dafny.Array
                nw99_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_661_v50_ = nw99_
                d_662_v52_: _dafny.Seq
                d_662_v52_ = _dafny.SeqWithoutIsStrInference([d_583_v9_])
                d_663_v53_: _dafny.Array
                nw100_ = _dafny.Array(None, 15)
                nw100_[int(0)] = d_583_v9_
                nw100_[int(1)] = d_583_v9_
                nw100_[int(2)] = d_583_v9_
                nw100_[int(3)] = d_583_v9_
                def iife68_():
                    coll18_ = _dafny.Map()
                    compr_18_: int
                    for compr_18_ in (d_583_v9_).keys.Elements:
                        d_664_v51_: int = compr_18_
                        if (d_664_v51_) in (d_583_v9_):
                            coll18_[(d_664_v51_) * (d_653_cf19_)] = d_653_cf19_
                    return _dafny.Map(coll18_)
                nw100_[int(4)] = iife68_()
                
                nw100_[int(5)] = (d_583_v9_).set((d_577_v3_).f5, (d_579_v5_).f5)
                nw100_[int(6)] = d_583_v9_
                nw100_[int(7)] = d_583_v9_
                nw100_[int(8)] = d_583_v9_
                nw100_[int(9)] = (d_662_v52_)[default__.safeIndex(-264, len(d_662_v52_))]
                nw100_[int(10)] = d_583_v9_
                nw100_[int(11)] = (d_583_v9_).set(967, (self).f8)
                nw100_[int(12)] = _dafny.Map({515: d_574_v0_})
                nw100_[int(13)] = d_583_v9_
                nw100_[int(14)] = d_583_v9_
                d_663_v53_ = nw100_
                index106_ = default__.safeIndex(473, (d_661_v50_).length(0))
                (d_661_v50_)[index106_] = d_663_v53_
                d_665_v54_: _dafny.Seq
                d_665_v54_ = _dafny.SeqWithoutIsStrInference([d_663_v53_, d_663_v53_, d_663_v53_, d_663_v53_])
                index107_ = default__.safeIndex(473, (d_661_v50_).length(0))
                (d_661_v50_)[index107_] = (d_665_v54_)[default__.safeIndex(-573, len(d_665_v54_))]
            elif source8_.is_DC3:
                d_666___mcc_h13_ = source8_.cf7
                d_667_cf7_ = d_666___mcc_h13_
                index108_ = default__.safeIndex(375, (d_620_v36_).length(0))
                (d_620_v36_)[index108_] = (d_579_v5_).f6
                index109_ = default__.safeIndex(109, (d_586_v12_).length(0))
                index110_ = default__.safeIndex(109, (d_586_v12_).length(0))
                rhs75_ = (d_579_v5_).f5
                rhs76_ = (d_577_v3_).f5
                lhs51_ = d_586_v12_
                lhs52_ = default__.safeIndex(109, (d_586_v12_).length(0))
                lhs53_ = d_586_v12_
                lhs54_ = default__.safeIndex(109, (d_586_v12_).length(0))
                lhs51_[lhs52_] = rhs75_
                lhs53_[lhs54_] = rhs76_
                d_668_v55_: D3
                d_668_v55_ = D3_DC11(d_586_v12_)
                d_669_v56_: _dafny.Map
                d_669_v56_ = _dafny.Map({d_668_v55_: (d_579_v5_).f5})
                d_669_v56_ = (d_669_v56_).set((d_668_v55_ if (self).f6 else D3_DC11(d_586_v12_)), (d_577_v3_).f5)
                d_670_v57_: _dafny.Map
                d_670_v57_ = _dafny.Map({(d_579_v5_).f5: p1})
                d_670_v57_ = (d_670_v57_).set(((d_579_v5_).f5) + (p0), p1)
            elif True:
                d_671___mcc_h14_ = source8_.cf21
                d_672_cf21_ = d_671___mcc_h14_
                index111_ = default__.safeIndex(375, (d_620_v36_).length(0))
                (d_620_v36_)[index111_] = (d_577_v3_).f6
                d_574_v0_ = p2
                d_673_v58_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_673_v58_ = nw101_
                d_674_v59_: _dafny.Seq
                d_674_v59_ = _dafny.SeqWithoutIsStrInference([d_620_v36_, d_620_v36_])
                d_675_v60_: _dafny.Map
                d_675_v60_ = _dafny.Map({True: (d_587_v13_).cf19})
                index112_ = default__.safeIndex(375, (d_620_v36_).length(0))
                index113_ = default__.safeIndex(109, (d_586_v12_).length(0))
                index114_ = default__.safeIndex(109, (d_586_v12_).length(0))
                rhs77_ = not((True) == ((d_620_v36_) not in (d_674_v59_)))
                rhs78_ = d_673_v58_
                rhs79_ = default__.safeModuloInt(((d_588_v14_) - (d_588_v14_)).cardinality, len(d_675_v60_))
                rhs80_ = default__.fm6(d_574_v0_, ((d_577_v3_).f5) + (p2), True, 197, globalState)
                lhs55_ = d_620_v36_
                lhs56_ = default__.safeIndex(375, (d_620_v36_).length(0))
                lhs57_ = d_586_v12_
                lhs58_ = default__.safeIndex(109, (d_586_v12_).length(0))
                lhs59_ = d_586_v12_
                lhs60_ = default__.safeIndex(109, (d_586_v12_).length(0))
                lhs55_[lhs56_] = rhs77_
                d_673_v58_ = rhs78_
                lhs57_[lhs58_] = rhs79_
                lhs59_[lhs60_] = rhs80_
                index115_ = default__.safeIndex(375, (d_620_v36_).length(0))
                (d_620_v36_)[index115_] = True
        d_676_v61_: C2
        nw102_ = C2()
        nw102_.ctor__(((d_586_v12_)[default__.safeIndex(109, (d_586_v12_).length(0))]) * ((0) - (d_574_v0_)))
        d_676_v61_ = nw102_
        d_677_i8_: int
        d_677_i8_ = 0
        with _dafny.label("4"):
            while not((d_579_v5_).f6):
                with _dafny.c_label("4"):
                    if (d_677_i8_) >= (100):
                        raise _dafny.Break("4")
                    d_677_i8_ = (d_677_i8_) + (1)
                    d_678_v62_: _dafny.Array
                    nw103_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                    d_678_v62_ = nw103_
                    d_679_v63_: _dafny.Map
                    d_679_v63_ = _dafny.Map({(p1 if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_680_i9_ in range(default__.abs(441))])): d_678_v62_})
                    d_679_v63_ = (d_679_v63_).set((p1) + (p1), d_678_v62_)
                    d_681_v64_: _dafny.Map
                    d_681_v64_ = _dafny.Map({(d_577_v3_).f6: (d_577_v3_).f5})
                    d_682_v65_: _dafny.Seq
                    d_682_v65_ = _dafny.SeqWithoutIsStrInference([d_580_v6_, _dafny.SeqWithoutIsStrInference([p2 for d_683_i10_ in range(default__.abs(141))])])
                    d_681_v64_ = (d_681_v64_).set(((d_682_v65_)[default__.safeIndex((d_579_v5_).f5, len(d_682_v65_))]) <= (d_580_v6_), (0) - ((d_579_v5_).f5))
                    d_684_v66_: C2
                    nw104_ = C2()
                    nw104_.ctor__(default__.safeModuloInt(585, (0) - ((self).f8)))
                    d_684_v66_ = nw104_
                    index116_ = default__.safeIndex(434, (d_678_v62_).length(0))
                    (d_678_v62_)[index116_] = p1
                    d_685_v67_: bool
                    d_685_v67_ = True
                    d_686_v68_: _dafny.Map
                    d_686_v68_ = _dafny.Map({(d_579_v5_).f5: d_681_v64_})
                    d_687_v69_: _dafny.Seq
                    d_687_v69_ = _dafny.SeqWithoutIsStrInference([default__.fm5(_dafny.Map({p3: len(_dafny.SeqWithoutIsStrInference([(d_579_v5_).f5]))}), d_686_v68_, globalState), p1])
                    index117_ = default__.safeIndex(434, (d_678_v62_).length(0))
                    rhs81_ = (p1) + ((d_687_v69_)[default__.safeIndex(len(d_582_v8_), len(d_687_v69_))])
                    rhs82_ = (self).f5
                    rhs83_ = p3
                    rhs84_ = not (((d_577_v3_).f5) <= ((d_577_v3_).f5)) or ((self).f6)
                    rhs85_ = (d_676_v61_.f9) != ((d_579_v5_).f5)
                    lhs61_ = d_678_v62_
                    lhs62_ = default__.safeIndex(434, (d_678_v62_).length(0))
                    lhs63_ = d_676_v61_
                    lhs61_[lhs62_] = rhs81_
                    lhs63_.f9 = rhs82_
                    d_685_v67_ = rhs83_
                    d_685_v67_ = rhs84_
                    d_685_v67_ = rhs85_
                    pass
            pass
        d_688_v70_: D3
        d_688_v70_ = D3_DC11(d_586_v12_)
        d_689_v71_: _dafny.Seq
        d_689_v71_ = _dafny.SeqWithoutIsStrInference([d_688_v70_])
        d_690_v72_: _dafny.MultiSet
        d_690_v72_ = _dafny.MultiSet([d_688_v70_])
        d_691_v73_: _dafny.Array
        nw105_ = _dafny.Array(None, 6)
        nw105_[int(0)] = _dafny.MultiSet((d_689_v71_) + (d_689_v71_))
        nw105_[int(1)] = (_dafny.MultiSet([d_688_v70_])) - (d_690_v72_)
        nw105_[int(2)] = d_690_v72_
        nw105_[int(3)] = d_690_v72_
        nw105_[int(4)] = d_690_v72_
        nw105_[int(5)] = (d_690_v72_) - (d_690_v72_)
        d_691_v73_ = nw105_
        d_691_v73_ = (d_691_v73_)
        d_692_v74_: _dafny.Array
        def lambda21_(d_693_v9_):
            def lambda22_(d_694_i11_):
                return d_693_v9_

            return lambda22_

        init12_ = lambda21_(d_583_v9_)
        nw106_ = _dafny.Array(None, 2)
        for i0_12_ in range(nw106_.length(0)):
            nw106_[i0_12_] = init12_(i0_12_)
        d_692_v74_ = nw106_
        d_692_v74_ = d_692_v74_
        d_695_v75_: _dafny.Array
        nw107_ = _dafny.Array(_dafny.Array(None, 0), 21)
        d_695_v75_ = nw107_
        d_696_v76_: _dafny.Seq
        d_696_v76_ = _dafny.SeqWithoutIsStrInference([d_695_v75_, d_695_v75_])
        r0 = (d_696_v76_)[default__.safeIndex(d_574_v0_, len(d_696_v76_))]
        d_697_v77_: D0
        d_697_v77_ = D0_DC0(d_580_v6_)
        r1 = d_697_v77_
        return r0, r1

    @property
    def f8(self):
        return self._f8

class C4(T0):
    def  __init__(self):
        self._f5: int = int(0)
        self._f6: bool = False
        self.f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f7, f5, f6):
        (self).f7 = f7
        (self)._f5 = f5
        (self)._f6 = f6

    def fm0(self, p0, globalState):
        source9_ = D0_DC0(_dafny.SeqWithoutIsStrInference([self.f7]))
        if source9_.is_DC1:
            d_698___mcc_h0_ = source9_.cf1
            d_699___mcc_h1_ = source9_.cf2
            d_700___mcc_h2_ = source9_.cf3
            d_701___mcc_h3_ = source9_.cf4
            d_702_cf4_ = d_701___mcc_h3_
            d_703_cf3_ = d_700___mcc_h2_
            d_704_cf2_ = d_699___mcc_h1_
            d_705_cf1_ = d_698___mcc_h0_
            return True
        elif source9_.is_DC2:
            d_706___mcc_h4_ = source9_.cf5
            d_707___mcc_h5_ = source9_.cf6
            d_708_cf6_ = d_707___mcc_h5_
            d_709_cf5_ = d_706___mcc_h4_
            return ((self).f6) in (_dafny.Set({(self).f6}))
        elif True:
            d_710___mcc_h6_ = source9_.cf0
            d_711_cf0_ = d_710___mcc_h6_
            return not((self).f6)

    def m1(self, p0, globalState):
        d_712_v0_: _dafny.Seq
        d_712_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
        d_713_v1_: _dafny.Seq
        d_713_v1_ = _dafny.SeqWithoutIsStrInference([self.f7, p0, p0, len(d_712_v0_), p0])
        d_714_v2_: D0
        d_714_v2_ = D0_DC2(_dafny.Map({False: d_713_v1_}), (self).f6)
        d_715_v3_: _dafny.Set
        d_715_v3_ = _dafny.Set({self.f7, self.f7})
        d_716_v4_: D1
        d_716_v4_ = D1_DC6((self).f6, self.f7, len(d_715_v3_), False)
        d_717_v5_: C0
        nw108_ = C0()
        nw108_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcehnr"))), (d_716_v4_).cf19)
        d_717_v5_ = nw108_
        d_718_v6_: _dafny.Map
        d_718_v6_ = _dafny.Map({d_717_v5_: (d_717_v5_).f12})
        d_719_v7_: _dafny.MultiSet
        d_719_v7_ = _dafny.MultiSet([d_718_v6_, d_718_v6_])
        d_720_v8_: _dafny.Map
        d_720_v8_ = _dafny.Map({-102: self.f7})
        d_721_v9_: C3
        nw109_ = C3()
        nw109_.ctor__((d_717_v5_).f12, (0) - (len(d_720_v8_)), (self).f6)
        d_721_v9_ = nw109_
        d_722_v10_: T0
        nw110_ = C3()
        nw110_.ctor__((d_719_v7_).cardinality, len(_dafny.Set({d_721_v9_, d_721_v9_})), (self).f6)
        d_722_v10_ = nw110_
        d_723_v11_: _dafny.Set
        d_723_v11_ = _dafny.Set({d_722_v10_})
        d_724_v12_: _dafny.Seq
        d_724_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhstj"))
        d_725_v13_: _dafny.Array
        nw111_ = _dafny.Array(None, 17)
        nw111_[int(0)] = (d_714_v2_).cf6
        nw111_[int(1)] = (56) > (p0)
        nw111_[int(2)] = (self).f6
        nw111_[int(3)] = (d_713_v1_) == (d_713_v1_)
        nw111_[int(4)] = (self).f6
        nw111_[int(5)] = (self).f6
        nw111_[int(6)] = (self).f6
        nw111_[int(7)] = (self.f7) not in (d_713_v1_)
        nw111_[int(8)] = (self).f6
        nw111_[int(9)] = not((self).f6)
        nw111_[int(10)] = (self).f6
        nw111_[int(11)] = (self).f6
        nw111_[int(12)] = (d_723_v11_).isdisjoint(d_723_v11_)
        nw111_[int(13)] = (d_724_v12_) == (d_724_v12_)
        nw111_[int(14)] = ((self).f6) and ((self).f6)
        nw111_[int(15)] = (d_722_v10_).f6
        nw111_[int(16)] = (self).f6
        d_725_v13_ = nw111_
        index118_ = default__.safeIndex(428, (d_725_v13_).length(0))
        (d_725_v13_)[index118_] = (d_722_v10_).f6
        index119_ = default__.safeIndex(428, (d_725_v13_).length(0))
        (d_725_v13_)[index119_] = (d_722_v10_).f6
        index120_ = default__.safeIndex(428, (d_725_v13_).length(0))
        (d_725_v13_)[index120_] = (d_722_v10_).f6
        if (self).f6:
            d_726_v14_: _dafny.Array
            nw112_ = _dafny.Array(_dafny.MultiSet({}), 17)
            d_726_v14_ = nw112_
            d_727_v15_: _dafny.Array
            d_727_v15_ = d_726_v14_
            index121_ = default__.safeIndex(428, (d_725_v13_).length(0))
            index122_ = default__.safeIndex(428, (d_725_v13_).length(0))
            rhs86_ = False
            rhs87_ = d_727_v15_
            rhs88_ = not((d_722_v10_).f6)
            rhs89_ = d_724_v12_
            rhs90_ = (d_721_v9_).f8
            lhs64_ = d_725_v13_
            lhs65_ = default__.safeIndex(428, (d_725_v13_).length(0))
            lhs66_ = d_725_v13_
            lhs67_ = default__.safeIndex(428, (d_725_v13_).length(0))
            lhs68_ = self
            lhs64_[lhs65_] = rhs86_
            d_727_v15_ = rhs87_
            lhs66_[lhs67_] = rhs88_
            d_724_v12_ = rhs89_
            lhs68_.f7 = rhs90_
            index123_ = default__.safeIndex(428, (d_725_v13_).length(0))
            (d_725_v13_)[index123_] = ((d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))] if (d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))] else (d_722_v10_).f6)
            d_728_v16_: _dafny.MultiSet
            d_728_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jijr"))])
            d_728_v16_ = (d_728_v16_) - (d_728_v16_)
            d_729_v17_: _dafny.Seq
            d_730_v18_: _dafny.Array
            d_731_v19_: int
            d_732_v20_: int
            out40_: _dafny.Seq
            out41_: _dafny.Array
            out42_: int
            out43_: int
            out40_, out41_, out42_, out43_ = default__.m0(((d_717_v5_).f12) * (self.f7), 253, globalState)
            d_729_v17_ = out40_
            d_730_v18_ = out41_
            d_731_v19_ = out42_
            d_732_v20_ = out43_
            d_731_v19_ = self.f7
        elif True:
            d_733_v21_: D5
            d_733_v21_ = D5_DC16(d_722_v10_, (d_722_v10_).f6, (0) - ((0) - ((d_722_v10_).f5)))
            d_734_v22_: _dafny.Map
            d_734_v22_ = _dafny.Map({(d_733_v21_).cf33: (d_713_v1_)[default__.safeIndex(self.f7, len(d_713_v1_))]})
            d_735_v23_: _dafny.Map
            d_735_v23_ = _dafny.Map({len(d_734_v22_): True})
            d_736_v24_: C3
            nw113_ = C3()
            nw113_.ctor__((d_721_v9_).f8, len(d_735_v23_), False)
            d_736_v24_ = nw113_
            d_737_v25_: C3
            nw114_ = C3()
            nw114_.ctor__((d_721_v9_).f8, (self.f7) * (((d_734_v22_)[(d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))]] if ((d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))]) in (d_734_v22_) else (_dafny.MultiSet([(d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))]])).cardinality)), ((d_736_v24_).f8) > ((d_736_v24_).f8))
            d_737_v25_ = nw114_
            d_738_v26_: str
            d_738_v26_ = _dafny.CodePoint('e')
            d_724_v12_ = ((d_724_v12_) + (d_724_v12_)) + ((d_724_v12_).set(default__.safeIndex(default__.fm6(-546, (d_717_v5_).f12, (d_722_v10_).f6, (d_717_v5_).f13, globalState), len(d_724_v12_)), d_738_v26_))
            d_739_v27_: C0
            nw115_ = C0()
            nw115_.ctor__((d_737_v25_).f8, (d_722_v10_).f5)
            d_739_v27_ = nw115_
            (self).f7 = default__.fm6((d_739_v27_).f12, self.f7, (d_722_v10_).f6, (self).f5, globalState)
        d_740_i0_: int
        d_740_i0_ = 0
        with _dafny.label("5"):
            while (d_722_v10_).f6:
                with _dafny.c_label("5"):
                    if (d_740_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_740_i0_ = (d_740_i0_) + (1)
                    d_741_v28_: _dafny.Array
                    def lambda23_(d_742_v0_):
                        def lambda24_(d_743_i1_):
                            return D5_DC15(_dafny.MultiSet(d_742_v0_))

                        return lambda24_

                    init13_ = lambda23_(d_712_v0_)
                    nw116_ = _dafny.Array(None, 29)
                    for i0_13_ in range(nw116_.length(0)):
                        nw116_[i0_13_] = init13_(i0_13_)
                    d_741_v28_ = nw116_
                    d_744_v29_: _dafny.MultiSet
                    d_744_v29_ = _dafny.MultiSet([(d_722_v10_).f6])
                    d_745_v30_: D5
                    d_745_v30_ = D5_DC15((d_744_v29_).set((self).f6, default__.abs(len(d_724_v12_))))
                    pat_let_tv24_ = d_717_v5_
                    pat_let_tv25_ = d_717_v5_
                    pat_let_tv26_ = globalState
                    index124_ = default__.safeIndex(784, (d_741_v28_).length(0))
                    def iife69_(_pat_let25_0):
                        def iife70_(d_746_dt__update__tmp_h0_):
                            def iife71_(_pat_let26_0):
                                def iife72_(d_747_dt__update_hcf31_h0_):
                                    return D5_DC15(d_747_dt__update_hcf31_h0_)
                                return iife72_(_pat_let26_0)
                            return iife71_(default__.fm10(_dafny.MultiSet([(pat_let_tv24_).f12, (pat_let_tv25_).f12, -863]), (self).f5, pat_let_tv26_))
                        return iife70_(_pat_let25_0)
                    (d_741_v28_)[index124_] = iife69_(d_745_v30_)
                    d_748_v31_: C1
                    nw117_ = C1()
                    nw117_.ctor__(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')])), (self).f5, (d_721_v9_).f8, False)
                    d_748_v31_ = nw117_
                    d_749_v32_: _dafny.Map
                    d_749_v32_ = _dafny.Map({(d_722_v10_).f6: d_748_v31_})
                    d_750_v33_: _dafny.Map
                    d_750_v33_ = _dafny.Map({(self).f6: self.f7})
                    index125_ = default__.safeIndex(784, (d_741_v28_).length(0))
                    (d_741_v28_)[index125_] = default__.fm22(len((d_749_v32_) | (d_749_v32_)), d_724_v12_, ((d_748_v31_).f11) - ((D0_DC1(d_724_v12_, (d_721_v9_).f8, d_750_v33_, d_713_v1_)).cf2), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_751_i2_ in range(default__.abs(126))]), globalState)
                    d_752_v34_: _dafny.MultiSet
                    d_752_v34_ = _dafny.MultiSet([(0) - ((d_717_v5_).f12), len(_dafny.Set({(self).f5}))])
                    d_753_v35_: D2
                    d_753_v35_ = D2_DC8(d_752_v34_)
                    d_753_v35_ = default__.fm23(self.f7, (len(d_724_v12_)) < (len(d_713_v1_)), 237, ((d_717_v5_).f12) + (len(d_724_v12_)), globalState)
                    d_754_v36_: _dafny.Map
                    d_754_v36_ = _dafny.Map({(d_721_v9_).f8: (d_724_v12_) != (d_724_v12_)})
                    d_755_v37_: D5
                    d_755_v37_ = D5_DC16(d_722_v10_, (self).f6, (d_717_v5_).f13)
                    d_754_v36_ = (d_754_v36_).set((d_755_v37_).cf34, (self).f6)
                    rhs91_ = (0) - ((d_717_v5_).f12)
                    rhs92_ = (_dafny.Map({(d_722_v10_).f5: (d_722_v10_).f6})).set(((d_722_v10_).f5) * ((d_717_v5_).f13), (d_725_v13_)[default__.safeIndex(428, (d_725_v13_).length(0))])
                    lhs69_ = self
                    lhs69_.f7 = rhs91_
                    d_754_v36_ = rhs92_
                    pass
            pass
        d_724_v12_ = d_724_v12_
        d_756_v38_: _dafny.Array
        nw118_ = _dafny.Array(None, 5)
        nw118_[int(0)] = (self).f5
        nw118_[int(1)] = self.f7
        nw118_[int(2)] = (d_717_v5_).f12
        nw118_[int(3)] = (d_722_v10_).f5
        nw118_[int(4)] = (self).f5
        d_756_v38_ = nw118_
        d_757_v39_: _dafny.MultiSet
        d_757_v39_ = _dafny.MultiSet([d_756_v38_])
        d_758_v40_: _dafny.Array
        nw119_ = _dafny.Array(None, 7)
        nw119_[int(0)] = d_757_v39_
        nw119_[int(1)] = (d_757_v39_).intersection(d_757_v39_)
        nw119_[int(2)] = _dafny.MultiSet([d_756_v38_, d_756_v38_, d_756_v38_, d_756_v38_, d_756_v38_])
        nw119_[int(3)] = (d_757_v39_) | (_dafny.MultiSet([d_756_v38_]))
        nw119_[int(4)] = d_757_v39_
        nw119_[int(5)] = _dafny.MultiSet([d_756_v38_, d_756_v38_])
        nw119_[int(6)] = d_757_v39_
        d_758_v40_ = nw119_
        index126_ = default__.safeIndex(860, (d_758_v40_).length(0))
        (d_758_v40_)[index126_] = d_757_v39_
        index127_ = default__.safeIndex(860, (d_758_v40_).length(0))
        (d_758_v40_)[index127_] = d_757_v39_

