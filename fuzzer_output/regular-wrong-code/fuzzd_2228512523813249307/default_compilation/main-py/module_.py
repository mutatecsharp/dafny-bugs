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
    def fm0(p0, p1, p2, globalState):
        return False

    @staticmethod
    def fm1(p0, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkbuevdy"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_0_i0_ in range(default__.abs(10))])))) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1_i1_ in range(default__.abs(895))])))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irien"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2_i0_ in range(default__.abs(552))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwfh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hc"))))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        if not (False) or (False):
            return D0_DC0(_dafny.Set({not(not(True)), False, not(True), True}))
        elif True:
            return D0_DC0(_dafny.Set({True}))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([-33, 87])) + ((_dafny.SeqWithoutIsStrInference([67])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjhywi"))), len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference([False]))])))

    @staticmethod
    def fm6(globalState):
        return _dafny.MultiSet([(_dafny.Map({_dafny.Map({len(_dafny.Set({True})): 977}): False})) == (_dafny.Map({_dafny.Map({88: 786}): False}))])

    @staticmethod
    def fm7(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: D0
            for compr_0_ in (_dafny.Map({D0_DC0(_dafny.Set({not(False)})): False})).keys.Elements:
                d_4_v0_: D0 = compr_0_
                if (d_4_v0_) in (_dafny.Map({D0_DC0(_dafny.Set({not(False)})): False})):
                    coll0_[d_4_v0_] = False
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: D0
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.Set({False})), D0_DC0(_dafny.Set({True}))])).Elements:
                d_5_v1_: D0 = compr_1_
                if (d_5_v1_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.Set({False})), D0_DC0(_dafny.Set({True}))])):
                    coll1_[d_5_v1_] = True
            return _dafny.Map(coll1_)
        return D2_DC6(D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({not(False)})): not(False)})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jm")), len((_dafny.SeqWithoutIsStrInference([D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({False})): True})) for d_3_i0_ in range(default__.abs(469))])) + (_dafny.SeqWithoutIsStrInference([D1_DC2(iife0_()
), D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({False})): False})), D1_DC2(iife1_()
), D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({True, False, False, False})): not(True)})), D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({True, True})): False}))]))))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        if (False if not(not(False)) else not(False)):
            return _dafny.CodePoint('f')
        elif True:
            return _dafny.CodePoint('n')

    @staticmethod
    def fm9(p0, globalState):
        return ((_dafny.Set({False})).intersection(_dafny.Set({False, False, False, True}))) - (_dafny.Set({False}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: D0
            for compr_2_ in (_dafny.Map({D0_DC0(_dafny.Set({False})): not(False)})).keys.Elements:
                d_6_v0_: D0 = compr_2_
                if (d_6_v0_) in (_dafny.Map({D0_DC0(_dafny.Set({False})): not(False)})):
                    coll2_[d_6_v0_] = False
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(162, 62):
                d_7_v1_: int = compr_3_
                if ((162) <= (d_7_v1_)) and ((d_7_v1_) < (62)):
                    coll3_[default__.safeModuloInt(d_7_v1_, 859)] = 979
            return _dafny.Map(coll3_)
        return ((_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([-398, len(_dafny.Map({True: len(_dafny.Set({D8_DC21(D2_DC6(D1_DC2(iife2_()
), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gjxpwm")), 833), _dafny.Map({_dafny.Map({True: 689}): True}))}))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))), len(_dafny.Set({False, True, True})), len(_dafny.Map({90: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, False]))).cardinality}))]))): 505})) | (_dafny.Map({793: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwa"))))}))) | ((_dafny.Map({len(_dafny.Map({False: False})): len(_dafny.Set({False}))})) | (iife3_()
        ))

    @staticmethod
    def fm13(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))) - (_dafny.MultiSet([True, not(not(False))]))) | ((_dafny.MultiSet([False, True, False, not(not(True))])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        if not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdpuwaavw"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyc")))):
            return _dafny.CodePoint('o')
        elif True:
            return _dafny.CodePoint('j')

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return (_dafny.Map({-171: 923})) | (_dafny.Map({451: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([444 for d_8_i0_ in range(default__.abs(400))])), len(_dafny.Map({True: 585}))])).cardinality]))}))

    @staticmethod
    def fm18(p0, p1, globalState):
        return (_dafny.Set({True, True, True})).intersection(_dafny.Set({False}))

    @staticmethod
    def fm19(p0, globalState):
        return _dafny.MultiSet([not(True)])

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm21(globalState):
        source0_ = D9_DC22(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(True)])).cardinality for d_9_i0_ in range(default__.abs(210))]))
        if source0_.is_DC23:
            d_10___mcc_h0_ = source0_.cf46
            d_11___mcc_h1_ = source0_.cf47
            d_12___mcc_h2_ = source0_.cf48
            d_13___mcc_h3_ = source0_.cf49
            d_14___mcc_h4_ = source0_.cf50
            d_15_cf50_ = d_14___mcc_h4_
            d_16_cf49_ = d_13___mcc_h3_
            d_17_cf48_ = d_12___mcc_h2_
            d_18_cf47_ = d_11___mcc_h1_
            d_19_cf46_ = d_10___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([d_16_cf49_, d_17_cf48_, not(d_19_cf46_)])) + (_dafny.SeqWithoutIsStrInference([d_19_cf46_]))
        elif True:
            d_20___mcc_h5_ = source0_.cf45
            d_21_cf45_ = d_20___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([(D1_DC4(True)).cf8])

    @staticmethod
    def fm22(p0, globalState):
        return _dafny.MultiSet([not(False), (_dafny.Set({460})).issubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wda")))})), not(((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbv")))])).cardinality) > (-888))])

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({True})): True}))

    @staticmethod
    def fm25(globalState):
        source1_ = D7_DC16(_dafny.Map({_dafny.CodePoint('d'): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrvjgmae")))}))
        if source1_.is_DC17:
            d_22___mcc_h0_ = source1_.cf31
            d_23___mcc_h1_ = source1_.cf32
            d_24_cf32_ = d_23___mcc_h1_
            d_25_cf31_ = d_22___mcc_h0_
            return (d_25_cf31_)[default__.safeIndex(137, len(d_25_cf31_))]
        elif True:
            d_26___mcc_h2_ = source1_.cf30
            d_27_cf30_ = d_26___mcc_h2_
            return _dafny.CodePoint('a')

    @staticmethod
    def fm26(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Set({-724})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Map({True: not(True)}))}) for d_28_i0_ in range(default__.abs(843))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjmi")))), (_dafny.MultiSet([53])).cardinality, len(_dafny.Map({756: not(True)})), len(_dafny.SeqWithoutIsStrInference([True, False])), -95}) for d_29_i1_ in range(default__.abs(298))]))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        source2_ = D12_DC30(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekdl")) for d_30_i0_ in range(default__.abs(713))]))
        if source2_.is_DC31:
            d_31___mcc_h0_ = source2_.cf60
            d_32_cf60_ = d_31___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([not(False)])
        elif True:
            d_33___mcc_h1_ = source2_.cf59
            d_34_cf59_ = d_33___mcc_h1_
            return (_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([False, not(False)]))

    @staticmethod
    def fm28(p0, p1, globalState):
        return D5_DC12(806)

    @staticmethod
    def fm29(globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_35_i0_ in range(default__.abs(750))])])).Elements:
                d_36_v0_: _dafny.Seq = compr_4_
                if (d_36_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True})) for d_35_i0_ in range(default__.abs(750))])])):
                    coll4_[d_36_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvfumpis")))
            return _dafny.Map(coll4_)
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(231, -903):
                d_37_v1_: int = compr_5_
                if ((231) <= (d_37_v1_)) and ((d_37_v1_) < (-903)):
                    coll5_[(d_37_v1_) + ((0) - (len(_dafny.Set({True}))))] = (0) - (len(_dafny.Map({len(_dafny.Map({True: True})): True})))
            return _dafny.Map(coll5_)
        return ((iife4_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({680: not(True)})), (_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True, False]))).cardinality])).cardinality]): len(iife5_()
        )}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([864]): len(_dafny.Map({_dafny.CodePoint('y'): True}))}))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_38_i0_ in range(default__.abs(-996))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsneoe"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmfqkanh")))})

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({True})) - (_dafny.Set({(D10_DC26(True)).cf52}))).intersection((_dafny.Set({not(False)}) if False else _dafny.Set({False, False})))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return D9_DC22(_dafny.SeqWithoutIsStrInference([998]))

    @staticmethod
    def fm33(globalState):
        return (_dafny.MultiSet([True, False, not(False), False, True])) | (_dafny.MultiSet([False]))

    @staticmethod
    def fm34(globalState):
        def iife6_():
            coll6_ = _dafny.Set()
            compr_6_: int
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])), -504])).Elements:
                d_39_v0_: int = compr_6_
                if (d_39_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])), -504])):
                    coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_39_v0_, 968)]))
            return _dafny.Set(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(-491, -155):
                d_41_v1_: int = compr_7_
                if ((-491) <= (d_41_v1_)) and ((d_41_v1_) < (-155)):
                    coll7_[default__.safeModuloInt(d_41_v1_, 778)] = _dafny.SeqWithoutIsStrInference([True, True])
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: _dafny.Set
            for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({731})])).Elements:
                d_42_v2_: _dafny.Set = compr_8_
                if (d_42_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({731})])):
                    coll8_ = coll8_.union(_dafny.Set([d_42_v2_]))
            return _dafny.Set(coll8_)
        return ((_dafny.Set({_dafny.Set({279, 407}), iife6_()
        })).intersection(_dafny.Set({_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True})) for d_40_i0_ in range(default__.abs(-84))])), len(iife7_()
        )})}))) - (iife8_()
        )

    @staticmethod
    def fm35(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - ((0) - (-469)), 655])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emsbo"))), 330, 554, len(_dafny.SeqWithoutIsStrInference([True]))]))) + (_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([len(_dafny.Map({132: True})), (_dafny.MultiSet([not(True)])).cardinality])).cardinality)]))

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return _dafny.Set({(_dafny.MultiSet([41, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqwwqd"))), 79, 710, 416])).cardinality, (len(_dafny.Set({False, False, True}))) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference([False, True, not(not(False)), False, not(not(False))])), len(_dafny.Map({True: False}))]))), default__.safeModuloInt(679, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_43_i0_ in range(default__.abs(188))])))})

    @staticmethod
    def fm37(p0, p1, globalState):
        return _dafny.Map({_dafny.CodePoint('x'): 875})

    @staticmethod
    def fm38(globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeomwdehc"))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_44_i1_ in range(default__.abs(561))]) for d_45_i0_ in range(default__.abs(-764))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unkma")) for d_46_i2_ in range(default__.abs(169))])))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(935, 433):
                    d_47_v0_: int = compr_11_
                    if ((935) <= (d_47_v0_)) and ((d_47_v0_) < (433)):
                        coll11_[default__.safeDivisionInt(d_47_v0_, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_48_i0_ in range(default__.abs(1))]): True})))] = True
                return _dafny.Map(coll11_)
            coll9_ = _dafny.Set()
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(935, 433):
                    d_47_v0_: int = compr_10_
                    if ((935) <= (d_47_v0_)) and ((d_47_v0_) < (433)):
                        coll10_[default__.safeDivisionInt(d_47_v0_, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_48_i0_ in range(default__.abs(1))]): True})))] = True
                return _dafny.Map(coll10_)
            compr_9_: int
            for compr_9_ in (_dafny.Map({len(iife10_()
            ): False})).keys.Elements:
                d_49_v1_: int = compr_9_
                if (d_49_v1_) in (_dafny.Map({len(iife11_()
                ): False})):
                    coll9_ = coll9_.union(_dafny.Set([(d_49_v1_) * (979)]))
            return _dafny.Set(coll9_)
        if (iife9_()
        ).isdisjoint(_dafny.Set({737})):
            return D10_DC24(_dafny.CodePoint('o'))
        elif True:
            return D10_DC24(_dafny.CodePoint('l'))

    @staticmethod
    def fm40(p0, p1, globalState):
        return _dafny.Map({len((_dafny.Map({not(False): True})) | (_dafny.Map({not(True): not(True)}))): _dafny.Map({691: not(not(True))})})

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r1 = False
        d_50_v0_: bool
        d_50_v0_ = False
        d_51_v1_: _dafny.Map
        d_51_v1_ = _dafny.Map({p1: p1})
        d_52_v2_: D8
        d_52_v2_ = D8_DC20(d_50_v0_, default__.fm0(d_50_v0_, (d_51_v1_).set(p1, p1), p1, globalState), d_50_v0_, (False) in (_dafny.Map({default__.fm0(d_50_v0_, d_51_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), globalState): 515})))
        d_53_v3_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 16)
        d_53_v3_ = nw0_
        d_54_v4_: _dafny.MultiSet
        d_54_v4_ = _dafny.MultiSet([d_53_v3_])
        d_52_v2_ = D8_DC20(d_50_v0_, (-148) <= (p1), True, (d_54_v4_).issubset(d_54_v4_))
        d_55_v5_: str
        d_55_v5_ = _dafny.CodePoint('d')
        d_56_v6_: T0
        nw1_ = C1()
        nw1_.ctor__(d_55_v5_, p1, len(_dafny.Map({d_50_v0_: d_50_v0_})))
        d_56_v6_ = nw1_
        d_57_v7_: D8
        d_57_v7_ = D8_DC18(d_56_v6_)
        pat_let_tv0_ = d_56_v6_
        pat_let_tv1_ = d_56_v6_
        def iife12_(_pat_let0_0):
            def iife13_(d_58_dt__update__tmp_h0_):
                def iife14_(_pat_let1_0):
                    def iife15_(d_59_dt__update_hcf33_h0_):
                        return D8_DC18(d_59_dt__update_hcf33_h0_)
                    return iife15_(_pat_let1_0)
                return iife14_((pat_let_tv0_ if False else pat_let_tv1_))
            return iife13_(_pat_let0_0)
        source3_ = iife12_(d_57_v7_)
        if source3_.is_DC19:
            d_60___mcc_h0_ = source3_.cf34
            d_61___mcc_h1_ = source3_.cf35
            d_62___mcc_h2_ = source3_.cf36
            d_63___mcc_h3_ = source3_.cf37
            d_64___mcc_h4_ = source3_.cf38
            d_65_cf38_ = d_64___mcc_h4_
            d_66_cf37_ = d_63___mcc_h3_
            d_67_cf36_ = d_62___mcc_h2_
            d_68_cf35_ = d_61___mcc_h1_
            d_69_cf34_ = d_60___mcc_h0_
            d_70_v8_: _dafny.Array
            nw2_ = _dafny.Array(None, 22)
            nw2_[int(0)] = d_69_cf34_
            nw2_[int(1)] = d_68_cf35_
            nw2_[int(2)] = d_50_v0_
            nw2_[int(3)] = d_68_cf35_
            nw2_[int(4)] = d_68_cf35_
            nw2_[int(5)] = d_50_v0_
            nw2_[int(6)] = d_69_cf34_
            nw2_[int(7)] = d_69_cf34_
            nw2_[int(8)] = d_50_v0_
            nw2_[int(9)] = d_68_cf35_
            nw2_[int(10)] = d_69_cf34_
            nw2_[int(11)] = False
            nw2_[int(12)] = d_69_cf34_
            nw2_[int(13)] = d_50_v0_
            nw2_[int(14)] = d_69_cf34_
            nw2_[int(15)] = d_50_v0_
            nw2_[int(16)] = d_68_cf35_
            nw2_[int(17)] = d_69_cf34_
            nw2_[int(18)] = d_50_v0_
            nw2_[int(19)] = d_69_cf34_
            nw2_[int(20)] = d_50_v0_
            nw2_[int(21)] = False
            d_70_v8_ = nw2_
            d_71_v9_: _dafny.Map
            d_71_v9_ = _dafny.Map({d_69_cf34_: d_66_cf37_})
            d_72_v10_: _dafny.Map
            d_72_v10_ = _dafny.Map({len(d_71_v9_): d_50_v0_})
            d_73_v11_: _dafny.Map
            d_73_v11_ = _dafny.Map({d_70_v8_: d_72_v10_})
            d_74_v12_: C4
            nw3_ = C4()
            nw3_.ctor__(d_69_cf34_, (d_73_v11_) | (_dafny.Map({d_70_v8_: d_72_v10_})), (p1) * (d_56_v6_.f7), (0) - ((d_56_v6_).f6))
            d_74_v12_ = nw3_
            (globalState).f1 = (p1) - ((d_56_v6_).f6)
            (d_56_v6_).f7 = (d_56_v6_).f6
            if d_50_v0_:
                d_75_v13_: _dafny.Seq
                d_75_v13_ = _dafny.SeqWithoutIsStrInference([d_66_cf37_])
                d_71_v9_ = (d_71_v9_).set((d_75_v13_) < (_dafny.SeqWithoutIsStrInference([93])), d_65_cf38_)
                d_65_cf38_ = ((d_56_v6_).f6) * (d_56_v6_.f7)
                d_76_v14_: _dafny.MultiSet
                d_76_v14_ = _dafny.MultiSet([(d_56_v6_).f6, (d_56_v6_).f6])
                d_77_v15_: _dafny.Seq
                d_77_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehcuiqm"))
                d_78_v16_: C3
                nw4_ = C3()
                nw4_.ctor__(d_56_v6_.f7, ((d_76_v14_).cardinality) - (len(d_77_v15_)))
                d_78_v16_ = nw4_
                d_79_v17_: _dafny.Seq
                d_79_v17_ = _dafny.SeqWithoutIsStrInference([(d_56_v6_ if d_50_v0_ else d_56_v6_)])
                rhs0_ = d_78_v16_
                rhs1_ = d_55_v5_
                rhs2_ = d_79_v17_
                rhs3_ = d_56_v6_.f7
                lhs0_ = d_56_v6_
                d_78_v16_ = rhs0_
                d_55_v5_ = rhs1_
                d_79_v17_ = rhs2_
                lhs0_.f7 = rhs3_
                d_80_v18_: C0
                nw5_ = C0()
                nw5_.ctor__((d_56_v6_.f7) < ((d_56_v6_).f6))
                d_80_v18_ = nw5_
                rhs4_ = d_77_v15_
                rhs5_ = 591
                rhs6_ = d_80_v18_
                rhs7_ = (d_56_v6_).f6
                rhs8_ = d_56_v6_
                lhs1_ = d_56_v6_
                lhs2_ = globalState
                d_77_v15_ = rhs4_
                lhs1_.f7 = rhs5_
                d_80_v18_ = rhs6_
                lhs2_.f1 = rhs7_
                d_56_v6_ = rhs8_
                d_55_v5_ = d_55_v5_
            elif True:
                d_81_v19_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_81_v19_ = nw6_
                d_81_v19_ = d_81_v19_
                d_82_v20_: _dafny.Map
                d_82_v20_ = _dafny.Map({_dafny.CodePoint('r'): d_71_v9_})
                d_82_v20_ = (d_82_v20_).set(d_55_v5_, d_71_v9_)
                (globalState).f1 = d_56_v6_.f7
                d_83_v21_: D13
                d_83_v21_ = D13_DC32(d_51_v1_)
                d_84_v22_: _dafny.Set
                d_84_v22_ = _dafny.Set({False, False})
                d_85_v23_: _dafny.Map
                d_85_v23_ = _dafny.Map({default__.fm0((d_74_v12_).f11, (d_83_v21_).cf61, 764, globalState): d_84_v22_})
                d_86_v24_: _dafny.Map
                d_86_v24_ = _dafny.Map({(d_74_v12_).f11: (d_74_v12_).f11})
                d_85_v23_ = (d_85_v23_).set(((d_86_v24_)[((d_72_v10_)[d_56_v6_.f7] if (d_56_v6_.f7) in (d_72_v10_) else d_69_cf34_)] if (((d_72_v10_)[d_56_v6_.f7] if (d_56_v6_.f7) in (d_72_v10_) else d_69_cf34_)) in (d_86_v24_) else d_50_v0_), d_84_v22_)
                d_68_cf35_ = d_69_cf34_
        elif source3_.is_DC20:
            d_87___mcc_h5_ = source3_.cf39
            d_88___mcc_h6_ = source3_.cf40
            d_89___mcc_h7_ = source3_.cf41
            d_90___mcc_h8_ = source3_.cf42
            d_91_cf42_ = d_90___mcc_h8_
            d_92_cf41_ = d_89___mcc_h7_
            d_93_cf40_ = d_88___mcc_h6_
            d_94_cf39_ = d_87___mcc_h5_
            d_95_v25_: _dafny.Array
            nw7_ = _dafny.Array(None, 19)
            d_95_v25_ = nw7_
            d_96_v26_: C0
            nw8_ = C0()
            nw8_.ctor__(d_94_cf39_)
            d_96_v26_ = nw8_
            index0_ = default__.safeIndex(5, (d_95_v25_).length(0))
            (d_95_v25_)[index0_] = d_96_v26_
            index1_ = default__.safeIndex(5, (d_95_v25_).length(0))
            (d_95_v25_)[index1_] = d_96_v26_
            if ((d_56_v6_).f6) <= ((0) - (d_56_v6_.f7)):
                d_97_v27_: _dafny.Seq
                d_97_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rcspe"))
                (d_96_v26_).f10 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjfovcio"))).set(default__.safeIndex(d_56_v6_.f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjfovcio")))), d_55_v5_)) < ((d_97_v27_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmouvrgh"))))
                (d_56_v6_).f7 = (d_56_v6_.f7) + ((d_56_v6_).f6)
                d_98_v28_: _dafny.MultiSet
                d_98_v28_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_55_v5_ for d_99_i0_ in range(default__.abs(799))])])
                d_100_v29_: _dafny.Seq
                d_100_v29_ = _dafny.SeqWithoutIsStrInference([d_91_cf42_, d_92_cf41_])
                d_101_v30_: _dafny.Seq
                d_101_v30_ = _dafny.SeqWithoutIsStrInference([(d_56_v6_).f6])
                rhs9_ = d_96_v26_.f10
                rhs10_ = ((d_56_v6_).f6) * (len((default__.fm35(len(d_100_v29_), (d_56_v6_).f6, globalState)) + (d_101_v30_)))
                rhs11_ = d_98_v28_
                rhs12_ = (d_56_v6_.f7) + (((d_56_v6_).f6) + (d_56_v6_.f7))
                lhs3_ = globalState
                d_93_cf40_ = rhs9_
                lhs3_.f1 = rhs10_
                d_98_v28_ = rhs11_
                r0 = rhs12_
                d_102_v31_: _dafny.Array
                def lambda0_(d_103_v6_):
                    def lambda1_(d_104_i1_):
                        return _dafny.Set({d_103_v6_.f7, (d_103_v6_).f6})

                    return lambda1_

                init0_ = lambda0_(d_56_v6_)
                nw9_ = _dafny.Array(None, 2)
                for i0_0_ in range(nw9_.length(0)):
                    nw9_[i0_0_] = init0_(i0_0_)
                d_102_v31_ = nw9_
                d_105_v32_: _dafny.Map
                d_105_v32_ = _dafny.Map({p0: (D11_DC28(d_97_v27_, d_102_v31_, d_93_cf40_, not(d_50_v0_))).cf54})
                d_105_v32_ = d_105_v32_
                d_106_v33_: _dafny.Seq
                d_106_v33_ = _dafny.SeqWithoutIsStrInference([d_53_v3_, d_53_v3_, d_53_v3_])
                d_53_v3_ = (d_106_v33_)[default__.safeIndex(d_56_v6_.f7, len(d_106_v33_))]
            elif True:
                d_107_v34_: _dafny.Seq
                d_107_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tobbymm"))
                d_108_v35_: _dafny.Seq
                d_108_v35_ = _dafny.SeqWithoutIsStrInference([len(d_107_v34_)])
                d_109_v36_: _dafny.Seq
                d_109_v36_ = _dafny.SeqWithoutIsStrInference([d_91_cf42_, d_50_v0_, d_50_v0_, ((d_108_v35_)[default__.safeIndex(d_56_v6_.f7, len(d_108_v35_))]) >= (len(_dafny.SeqWithoutIsStrInference([d_94_cf39_]))), (d_55_v5_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlngpj")))])
                rhs13_ = (d_57_v7_).cf33
                rhs14_ = not((d_109_v36_)[default__.safeIndex(412, len(d_109_v36_))])
                rhs15_ = d_107_v34_
                lhs4_ = d_96_v26_
                d_56_v6_ = rhs13_
                lhs4_.f10 = rhs14_
                d_107_v34_ = rhs15_
                d_110_v37_: _dafny.Array
                nw10_ = _dafny.Array(False, 4)
                d_110_v37_ = nw10_
                d_111_v38_: _dafny.Map
                d_111_v38_ = _dafny.Map({d_56_v6_: d_96_v26_.f10})
                index2_ = default__.safeIndex(654, (d_110_v37_).length(0))
                (d_110_v37_)[index2_] = default__.fm0(d_50_v0_, (d_51_v1_).set(len(d_111_v38_), d_56_v6_.f7), (0) - (len(_dafny.Map({_dafny.CodePoint('f'): d_92_cf41_}))), globalState)
                index3_ = default__.safeIndex(654, (d_110_v37_).length(0))
                (d_110_v37_)[index3_] = d_93_cf40_
                d_107_v34_ = d_107_v34_
                d_112_v39_: _dafny.Map
                d_112_v39_ = _dafny.Map({p1: d_91_cf42_})
                d_113_v40_: _dafny.Map
                d_113_v40_ = _dafny.Map({d_110_v37_: d_112_v39_})
                d_114_v41_: C4
                nw11_ = C4()
                nw11_.ctor__(d_96_v26_.f10, d_113_v40_, default__.safeDivisionInt(len(d_107_v34_), (0) - (d_56_v6_.f7)), default__.fm1(d_91_cf42_, globalState))
                d_114_v41_ = nw11_
                d_114_v41_ = d_114_v41_
                (d_56_v6_).f7 = d_56_v6_.f7
            d_115_v42_: _dafny.Map
            d_115_v42_ = _dafny.Map({not(not(d_93_cf40_)): d_53_v3_})
            d_116_v43_: D1
            d_116_v43_ = D1_DC3(d_93_cf40_, p1, len(d_115_v42_))
            d_117_v44_: _dafny.Map
            d_117_v44_ = _dafny.Map({d_92_cf41_: (d_56_v6_).f6})
            d_118_v45_: _dafny.Map
            d_118_v45_ = _dafny.Map({d_96_v26_.f10: d_91_cf42_})
            d_119_v46_: _dafny.Map
            d_119_v46_ = _dafny.Map({d_116_v43_: ((d_117_v44_)[((d_118_v45_)[d_50_v0_] if (d_50_v0_) in (d_118_v45_) else d_92_cf41_)] if (((d_118_v45_)[d_50_v0_] if (d_50_v0_) in (d_118_v45_) else d_92_cf41_)) in (d_117_v44_) else d_56_v6_.f7)})
            d_119_v46_ = (d_119_v46_).set(d_116_v43_, (d_56_v6_.f7) * ((d_56_v6_).f6))
            d_120_v47_: _dafny.Map
            d_120_v47_ = _dafny.Map({d_56_v6_: d_53_v3_})
            d_121_v48_: _dafny.Array
            nw12_ = _dafny.Array(None, 26)
            nw12_[int(0)] = d_53_v3_
            nw12_[int(1)] = d_53_v3_
            nw12_[int(2)] = d_53_v3_
            nw12_[int(3)] = d_53_v3_
            nw12_[int(4)] = d_53_v3_
            nw12_[int(5)] = d_53_v3_
            nw12_[int(6)] = d_53_v3_
            nw12_[int(7)] = d_53_v3_
            nw12_[int(8)] = d_53_v3_
            nw12_[int(9)] = d_53_v3_
            nw12_[int(10)] = d_53_v3_
            nw12_[int(11)] = d_53_v3_
            nw12_[int(12)] = d_53_v3_
            nw12_[int(13)] = ((d_120_v47_)[d_56_v6_] if (d_56_v6_) in (d_120_v47_) else d_53_v3_)
            nw12_[int(14)] = d_53_v3_
            nw12_[int(15)] = d_53_v3_
            nw12_[int(16)] = d_53_v3_
            nw12_[int(17)] = d_53_v3_
            nw12_[int(18)] = d_53_v3_
            nw12_[int(19)] = d_53_v3_
            nw12_[int(20)] = d_53_v3_
            nw12_[int(21)] = d_53_v3_
            nw12_[int(22)] = d_53_v3_
            nw12_[int(23)] = d_53_v3_
            nw12_[int(24)] = d_53_v3_
            nw12_[int(25)] = d_53_v3_
            d_121_v48_ = nw12_
            d_122_v49_: _dafny.Seq
            d_122_v49_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            d_123_v50_: C5
            nw13_ = C5()
            nw13_.ctor__(d_121_v48_, d_122_v49_, (d_56_v6_).f6, (d_56_v6_).f6)
            d_123_v50_ = nw13_
        elif source3_.is_DC21:
            d_124___mcc_h9_ = source3_.cf43
            d_125___mcc_h10_ = source3_.cf44
            d_126_cf44_ = d_125___mcc_h10_
            d_127_cf43_ = d_124___mcc_h9_
            d_128_v51_: _dafny.MultiSet
            d_128_v51_ = _dafny.MultiSet([d_50_v0_])
            r1 = ((d_128_v51_).intersection(d_128_v51_)).ispropersubset(d_128_v51_)
            if d_50_v0_:
                d_129_v52_: C2
                nw14_ = C2()
                nw14_.ctor__((d_56_v6_).f6, (d_56_v6_).f6)
                d_129_v52_ = nw14_
                d_130_v53_: _dafny.Seq
                d_130_v53_ = _dafny.SeqWithoutIsStrInference([p1])
                d_131_v54_: D9
                d_131_v54_ = D9_DC22(d_130_v53_)
                d_131_v54_ = d_131_v54_
                d_132_v55_: _dafny.Map
                d_132_v55_ = _dafny.Map({_dafny.Set({d_50_v0_, d_50_v0_, d_50_v0_}): (d_56_v6_).f6})
                d_133_v56_: _dafny.Set
                d_133_v56_ = _dafny.Set({d_50_v0_})
                d_132_v55_ = (d_132_v55_).set(d_133_v56_, (d_56_v6_).f6)
                d_134_v57_: C3
                nw15_ = C3()
                nw15_.ctor__(len(_dafny.SeqWithoutIsStrInference([d_56_v6_.f7])), (0) - ((d_56_v6_).f6))
                d_134_v57_ = nw15_
                d_135_v58_: _dafny.Seq
                d_135_v58_ = _dafny.SeqWithoutIsStrInference([d_134_v57_, d_134_v57_, d_134_v57_])
                d_136_v59_: _dafny.Seq
                d_136_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxyo"))
                index4_ = default__.safeIndex(952, (d_53_v3_).length(0))
                (d_53_v3_)[index4_] = len(d_136_v59_)
                d_137_v60_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Seq({}), 3)
                d_137_v60_ = nw16_
                index5_ = default__.safeIndex(952, (d_53_v3_).length(0))
                rhs16_ = d_135_v58_
                rhs17_ = default__.safeDivisionInt((d_56_v6_).f6, d_56_v6_.f7)
                rhs18_ = (d_137_v60_ if (d_56_v6_.f7) <= (d_56_v6_.f7) else d_137_v60_)
                lhs5_ = d_53_v3_
                lhs6_ = default__.safeIndex(952, (d_53_v3_).length(0))
                d_135_v58_ = rhs16_
                lhs5_[lhs6_] = rhs17_
                d_137_v60_ = rhs18_
                d_138_v61_: _dafny.Map
                d_138_v61_ = _dafny.Map({d_50_v0_: d_50_v0_})
                d_139_v62_: D0
                d_139_v62_ = D0_DC0(d_133_v56_)
                d_140_v63_: _dafny.Map
                d_140_v63_ = _dafny.Map({d_139_v62_: False})
                d_141_v64_: D1
                d_141_v64_ = D1_DC2(d_140_v63_)
                (d_129_v52_).m8(d_138_v61_, D2_DC6(d_141_v64_, d_136_v59_, -757), d_136_v59_, globalState)
            elif True:
                d_142_v65_: _dafny.Array
                nw17_ = _dafny.Array(None, 1)
                d_142_v65_ = nw17_
                nw18_ = _dafny.Array(None, 9)
                d_142_v65_ = nw18_
                d_143_v66_: _dafny.Seq
                d_143_v66_ = _dafny.SeqWithoutIsStrInference([(d_56_v6_).f6])
                d_144_v67_: _dafny.Map
                d_144_v67_ = _dafny.Map({d_51_v1_: d_143_v66_})
                d_143_v66_ = ((d_144_v67_)[(d_51_v1_) | (d_51_v1_)] if ((d_51_v1_) | (d_51_v1_)) in (d_144_v67_) else d_143_v66_)
                d_145_v68_: _dafny.Seq
                d_145_v68_ = _dafny.SeqWithoutIsStrInference([d_50_v0_, d_50_v0_, d_50_v0_])
                (d_56_v6_).f7 = len(default__.fm2(d_143_v66_, d_56_v6_.f7, d_56_v6_.f7, (_dafny.MultiSet(d_145_v68_)).ispropersubset(d_128_v51_), globalState))
                d_50_v0_ = False
                r1 = d_50_v0_
            index6_ = default__.safeIndex(329, (d_53_v3_).length(0))
            (d_53_v3_)[index6_] = (d_56_v6_).f6
            d_146_v69_: _dafny.Seq
            d_146_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fummh"))
            d_147_v70_: _dafny.Seq
            d_147_v70_ = _dafny.SeqWithoutIsStrInference([(d_56_v6_).f6, len(_dafny.Map({d_56_v6_.f7: d_50_v0_}))])
            d_148_v71_: _dafny.Seq
            d_148_v71_ = _dafny.SeqWithoutIsStrInference([len(d_147_v70_)])
            index7_ = default__.safeIndex(329, (d_53_v3_).length(0))
            rhs19_ = (d_56_v6_).f6
            rhs20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pba"))
            rhs21_ = d_53_v3_
            rhs22_ = d_147_v70_
            rhs23_ = d_50_v0_
            lhs7_ = d_53_v3_
            lhs8_ = default__.safeIndex(329, (d_53_v3_).length(0))
            lhs7_[lhs8_] = rhs19_
            d_146_v69_ = rhs20_
            d_53_v3_ = rhs21_
            d_148_v71_ = rhs22_
            r1 = rhs23_
            d_149_v72_: _dafny.Array
            def lambda2_(d_150_v0_):
                def lambda3_(d_151_i2_):
                    return d_150_v0_

                return lambda3_

            init1_ = lambda2_(d_50_v0_)
            nw19_ = _dafny.Array(None, 18)
            for i0_1_ in range(nw19_.length(0)):
                nw19_[i0_1_] = init1_(i0_1_)
            d_149_v72_ = nw19_
            index8_ = default__.safeIndex(894, (d_149_v72_).length(0))
            (d_149_v72_)[index8_] = not(d_50_v0_)
            index9_ = default__.safeIndex(894, (d_149_v72_).length(0))
            (d_149_v72_)[index9_] = d_50_v0_
        elif True:
            d_152___mcc_h11_ = source3_.cf33
            d_153_cf33_ = d_152___mcc_h11_
            d_154_v73_: D10
            d_154_v73_ = D10_DC24(d_55_v5_)
            pat_let_tv2_ = d_55_v5_
            d_155_v74_: _dafny.Map
            def iife16_(_pat_let2_0):
                def iife17_(d_156_dt__update__tmp_h1_):
                    def iife18_(_pat_let3_0):
                        def iife19_(d_157_dt__update_hcf51_h0_):
                            return D10_DC24(d_157_dt__update_hcf51_h0_)
                        return iife19_(_pat_let3_0)
                    return iife18_(pat_let_tv2_)
                return iife17_(_pat_let2_0)
            d_155_v74_ = _dafny.Map({(iife16_(d_154_v73_)).cf51: d_153_cf33_})
            d_158_v75_: D8
            d_158_v75_ = D8_DC19(True, (d_55_v5_) not in (d_155_v74_), _dafny.Set({d_53_v3_, d_53_v3_}), (d_153_cf33_.f7) - (p1), default__.safeDivisionInt(d_56_v6_.f7, (d_56_v6_).f6))
            source4_ = d_158_v75_
            if source4_.is_DC19:
                d_159___mcc_h12_ = source4_.cf34
                d_160___mcc_h13_ = source4_.cf35
                d_161___mcc_h14_ = source4_.cf36
                d_162___mcc_h15_ = source4_.cf37
                d_163___mcc_h16_ = source4_.cf38
                d_164_cf38_ = d_163___mcc_h16_
                d_165_cf37_ = d_162___mcc_h15_
                d_166_cf36_ = d_161___mcc_h14_
                d_167_cf35_ = d_160___mcc_h13_
                d_168_cf34_ = d_159___mcc_h12_
                d_169_v76_: _dafny.Seq
                d_169_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvmlujfv"))
                d_170_v77_: _dafny.MultiSet
                d_170_v77_ = _dafny.MultiSet([d_169_v76_])
                d_170_v77_ = d_170_v77_
                d_171_v78_: _dafny.Array
                def lambda4_(d_172_v0_):
                    def lambda5_(d_173_i3_):
                        return d_172_v0_

                    return lambda5_

                init2_ = lambda4_(d_50_v0_)
                nw20_ = _dafny.Array(None, 18)
                for i0_2_ in range(nw20_.length(0)):
                    nw20_[i0_2_] = init2_(i0_2_)
                d_171_v78_ = nw20_
                d_174_v79_: _dafny.Map
                d_174_v79_ = _dafny.Map({(d_56_v6_).f6: not(d_50_v0_)})
                d_175_v80_: _dafny.Array
                nw21_ = _dafny.Array(None, 13)
                nw21_[int(0)] = d_171_v78_
                nw21_[int(1)] = d_171_v78_
                nw21_[int(2)] = d_171_v78_
                nw21_[int(3)] = d_171_v78_
                nw21_[int(4)] = d_171_v78_
                nw21_[int(5)] = d_171_v78_
                nw21_[int(6)] = (d_171_v78_ if ((d_174_v79_)[(d_153_cf33_).f6] if ((d_153_cf33_).f6) in (d_174_v79_) else d_168_cf34_) else d_171_v78_)
                nw21_[int(7)] = d_171_v78_
                nw21_[int(8)] = d_171_v78_
                nw21_[int(9)] = d_171_v78_
                nw21_[int(10)] = d_171_v78_
                nw21_[int(11)] = d_171_v78_
                nw21_[int(12)] = d_171_v78_
                d_175_v80_ = nw21_
                index10_ = default__.safeIndex(481, (d_175_v80_).length(0))
                (d_175_v80_)[index10_] = d_171_v78_
                index11_ = default__.safeIndex(481, (d_175_v80_).length(0))
                (d_175_v80_)[index11_] = d_171_v78_
                index12_ = default__.safeIndex(177, (d_53_v3_).length(0))
                (d_53_v3_)[index12_] = default__.safeModuloInt(d_56_v6_.f7, p1)
                index13_ = default__.safeIndex(177, (d_53_v3_).length(0))
                (d_53_v3_)[index13_] = d_56_v6_.f7
                d_176_v81_: _dafny.Map
                d_176_v81_ = _dafny.Map({d_55_v5_: (113) - (len(d_169_v76_))})
                (d_153_cf33_).f7 = len(d_176_v81_)
            elif source4_.is_DC20:
                d_177___mcc_h17_ = source4_.cf39
                d_178___mcc_h18_ = source4_.cf40
                d_179___mcc_h19_ = source4_.cf41
                d_180___mcc_h20_ = source4_.cf42
                d_181_cf42_ = d_180___mcc_h20_
                d_182_cf41_ = d_179___mcc_h19_
                d_183_cf40_ = d_178___mcc_h18_
                d_184_cf39_ = d_177___mcc_h17_
                d_185_v82_: _dafny.Map
                d_185_v82_ = _dafny.Map({d_181_cf42_: (d_56_v6_).f6})
                d_186_v83_: _dafny.Map
                d_186_v83_ = _dafny.Map({d_153_cf33_.f7: d_184_cf39_})
                d_185_v82_ = (d_185_v82_).set((d_56_v6_.f7) > ((d_153_cf33_).f6), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([((d_186_v83_)[224] if (224) in (d_186_v83_) else False), d_181_cf42_, True])), (d_153_cf33_).f6))
                d_51_v1_ = (d_51_v1_) | (_dafny.Map({(d_56_v6_).f6: d_56_v6_.f7}))
                (d_56_v6_).f7 = (d_56_v6_).f6
                d_187_v84_: _dafny.Seq
                d_187_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiffwqiab"))
                d_188_v85_: _dafny.MultiSet
                d_188_v85_ = _dafny.MultiSet([len((d_51_v1_).set((d_153_cf33_).f6, (d_56_v6_).f6)), len(d_185_v82_), (d_56_v6_).f6])
                d_189_v86_: _dafny.Map
                d_189_v86_ = _dafny.Map({d_188_v85_: d_187_v84_})
                d_190_v87_: _dafny.Array
                nw22_ = _dafny.Array(None, 22)
                nw22_[int(0)] = d_187_v84_
                nw22_[int(1)] = d_187_v84_
                nw22_[int(2)] = _dafny.SeqWithoutIsStrInference([d_55_v5_ for d_191_i4_ in range(default__.abs(397))])
                nw22_[int(3)] = d_187_v84_
                nw22_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stfj"))
                nw22_[int(5)] = d_187_v84_
                nw22_[int(6)] = d_187_v84_
                nw22_[int(7)] = d_187_v84_
                nw22_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iausxosc"))
                nw22_[int(9)] = d_187_v84_
                nw22_[int(10)] = d_187_v84_
                nw22_[int(11)] = d_187_v84_
                nw22_[int(12)] = d_187_v84_
                nw22_[int(13)] = d_187_v84_
                nw22_[int(14)] = d_187_v84_
                nw22_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjnked"))
                nw22_[int(16)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_192_i5_ in range(default__.abs(514))])).set(default__.safeIndex(d_56_v6_.f7, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_193_i5_ in range(default__.abs(514))]))), d_55_v5_)
                nw22_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laj"))
                nw22_[int(18)] = d_187_v84_
                nw22_[int(19)] = d_187_v84_
                nw22_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xt"))
                nw22_[int(21)] = ((d_189_v86_)[d_188_v85_] if (d_188_v85_) in (d_189_v86_) else d_187_v84_)
                d_190_v87_ = nw22_
                d_194_v88_: D13
                d_194_v88_ = D13_DC33(d_182_cf41_, (d_153_cf33_).f6)
                d_195_v89_: _dafny.Array
                nw23_ = _dafny.Array(None, 22)
                nw23_[int(0)] = d_190_v87_
                nw23_[int(1)] = d_190_v87_
                nw23_[int(2)] = d_190_v87_
                nw23_[int(3)] = d_190_v87_
                nw23_[int(4)] = d_190_v87_
                nw23_[int(5)] = (d_190_v87_ if d_181_cf42_ else d_190_v87_)
                nw23_[int(6)] = d_190_v87_
                nw23_[int(7)] = d_190_v87_
                nw23_[int(8)] = d_190_v87_
                nw23_[int(9)] = d_190_v87_
                nw23_[int(10)] = d_190_v87_
                nw23_[int(11)] = d_190_v87_
                nw23_[int(12)] = d_190_v87_
                nw23_[int(13)] = d_190_v87_
                nw23_[int(14)] = d_190_v87_
                nw23_[int(15)] = d_190_v87_
                nw23_[int(16)] = d_190_v87_
                nw23_[int(17)] = (d_190_v87_ if (d_194_v88_).cf62 else d_190_v87_)
                nw23_[int(18)] = d_190_v87_
                nw23_[int(19)] = d_190_v87_
                nw23_[int(20)] = (d_190_v87_ if d_183_cf40_ else d_190_v87_)
                nw23_[int(21)] = d_190_v87_
                d_195_v89_ = nw23_
                index14_ = default__.safeIndex(458, (d_195_v89_).length(0))
                (d_195_v89_)[index14_] = d_190_v87_
                index15_ = default__.safeIndex(458, (d_195_v89_).length(0))
                (d_195_v89_)[index15_] = d_190_v87_
            elif source4_.is_DC21:
                d_196___mcc_h21_ = source4_.cf43
                d_197___mcc_h22_ = source4_.cf44
                d_198_cf44_ = d_197___mcc_h22_
                d_199_cf43_ = d_196___mcc_h21_
                d_55_v5_ = d_55_v5_
                d_200_v90_: D10
                d_200_v90_ = D10_DC25()
                d_200_v90_ = (d_200_v90_ if d_50_v0_ else D10_DC25())
                d_201_v91_: _dafny.Array
                def lambda6_(d_202_v5_):
                    def lambda7_(d_203_i6_):
                        return d_202_v5_

                    return lambda7_

                init3_ = lambda6_(d_55_v5_)
                nw24_ = _dafny.Array(None, 23)
                for i0_3_ in range(nw24_.length(0)):
                    nw24_[i0_3_] = init3_(i0_3_)
                d_201_v91_ = nw24_
                index16_ = default__.safeIndex(433, (d_201_v91_).length(0))
                (d_201_v91_)[index16_] = (d_55_v5_ if d_50_v0_ else d_55_v5_)
                index17_ = default__.safeIndex(433, (d_201_v91_).length(0))
                (d_201_v91_)[index17_] = d_55_v5_
                r1 = (d_50_v0_) or (False)
            elif True:
                d_204___mcc_h23_ = source4_.cf33
                d_205_cf33_ = d_204___mcc_h23_
                d_206_v92_: D10
                d_206_v92_ = D10_DC26(d_50_v0_)
                d_206_v92_ = D10_DC26(d_50_v0_)
                d_207_v93_: _dafny.Seq
                d_207_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpmp"))
                d_208_v94_: _dafny.Set
                d_208_v94_ = _dafny.Set({(d_56_v6_).f6, (d_205_cf33_).f6, (d_56_v6_).f6, (0) - (d_56_v6_.f7)})
                d_209_v95_: _dafny.Array
                nw25_ = _dafny.Array(None, 4)
                nw25_[int(0)] = d_208_v94_
                nw25_[int(1)] = d_208_v94_
                nw25_[int(2)] = d_208_v94_
                nw25_[int(3)] = d_208_v94_
                d_209_v95_ = nw25_
                d_210_v96_: D11
                d_210_v96_ = D11_DC28(d_207_v93_, d_209_v95_, d_50_v0_, not(d_50_v0_))
                d_211_v97_: _dafny.Set
                d_211_v97_ = _dafny.Set({d_50_v0_, (d_210_v96_).cf56})
                d_211_v97_ = _dafny.Set({d_50_v0_})
                d_212_v98_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_212_v98_ = nw26_
                index18_ = default__.safeIndex(150, (d_212_v98_).length(0))
                (d_212_v98_)[index18_] = d_53_v3_
                index19_ = default__.safeIndex(150, (d_212_v98_).length(0))
                rhs24_ = d_53_v3_
                rhs25_ = d_207_v93_
                rhs26_ = (not (d_50_v0_) or (d_50_v0_)) and (not (d_50_v0_) or (d_50_v0_))
                lhs9_ = d_212_v98_
                lhs10_ = default__.safeIndex(150, (d_212_v98_).length(0))
                lhs9_[lhs10_] = rhs24_
                d_207_v93_ = rhs25_
                d_50_v0_ = rhs26_
                d_213_v99_: _dafny.Seq
                d_213_v99_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_205_cf33_).f6)])
                d_214_v100_: C5
                nw27_ = C5()
                nw27_.ctor__(d_212_v98_, d_213_v99_, (d_56_v6_).f6, (d_56_v6_).f6)
                d_214_v100_ = nw27_
                d_214_v100_ = d_214_v100_
            d_215_v101_: C3
            nw28_ = C3()
            nw28_.ctor__(p1, (d_153_cf33_).f6)
            d_215_v101_ = nw28_
            index20_ = default__.safeIndex(771, (d_53_v3_).length(0))
            (d_53_v3_)[index20_] = (d_56_v6_.f7 if False else (d_56_v6_).f6)
            d_216_v102_: _dafny.Array
            nw29_ = _dafny.Array(None, 13)
            nw29_[int(0)] = d_50_v0_
            nw29_[int(1)] = d_50_v0_
            nw29_[int(2)] = d_50_v0_
            nw29_[int(3)] = d_50_v0_
            nw29_[int(4)] = d_50_v0_
            nw29_[int(5)] = (d_158_v75_).cf35
            nw29_[int(6)] = d_50_v0_
            nw29_[int(7)] = d_50_v0_
            nw29_[int(8)] = d_50_v0_
            nw29_[int(9)] = d_50_v0_
            nw29_[int(10)] = d_50_v0_
            nw29_[int(11)] = d_50_v0_
            nw29_[int(12)] = not((d_215_v101_).fm16(globalState))
            d_216_v102_ = nw29_
            d_217_v103_: _dafny.Set
            d_217_v103_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_215_v101_])})
            d_218_v104_: _dafny.Map
            d_218_v104_ = _dafny.Map({len(d_217_v103_): d_50_v0_})
            d_219_v105_: _dafny.Map
            d_219_v105_ = _dafny.Map({d_216_v102_: d_218_v104_})
            d_220_v106_: C4
            nw30_ = C4()
            nw30_.ctor__(d_50_v0_, d_219_v105_, default__.fm1(d_50_v0_, globalState), (d_56_v6_).f6)
            d_220_v106_ = nw30_
            d_221_v107_: _dafny.Map
            d_221_v107_ = _dafny.Map({d_220_v106_: True})
            d_222_v108_: _dafny.Seq
            d_222_v108_ = _dafny.SeqWithoutIsStrInference([d_56_v6_])
            index21_ = default__.safeIndex(771, (d_53_v3_).length(0))
            rhs27_ = ((d_56_v6_.f7) == (len(d_221_v107_)) if d_50_v0_ else (d_220_v106_).f11)
            rhs28_ = (d_220_v106_).f11
            rhs29_ = ((d_153_cf33_).f6) * (default__.safeModuloInt(len(d_222_v108_), p1))
            rhs30_ = d_153_cf33_.f7
            lhs11_ = d_153_cf33_
            lhs12_ = d_53_v3_
            lhs13_ = default__.safeIndex(771, (d_53_v3_).length(0))
            d_50_v0_ = rhs27_
            d_50_v0_ = rhs28_
            lhs11_.f7 = rhs29_
            lhs12_[lhs13_] = rhs30_
            index22_ = default__.safeIndex(256, (d_216_v102_).length(0))
            (d_216_v102_)[index22_] = (d_220_v106_).f11
            d_223_v109_: _dafny.Array
            nw31_ = _dafny.Array(None, 1)
            d_223_v109_ = nw31_
            index23_ = default__.safeIndex(343, (d_223_v109_).length(0))
            (d_223_v109_)[index23_] = d_220_v106_
            index24_ = default__.safeIndex(256, (d_216_v102_).length(0))
            index25_ = default__.safeIndex(343, (d_223_v109_).length(0))
            rhs31_ = (d_220_v106_).f11
            rhs32_ = d_220_v106_
            rhs33_ = (d_53_v3_)[default__.safeIndex(771, (d_53_v3_).length(0))]
            rhs34_ = default__.safeModuloInt(p1, (0) - ((d_56_v6_).f6))
            lhs14_ = d_216_v102_
            lhs15_ = default__.safeIndex(256, (d_216_v102_).length(0))
            lhs16_ = d_223_v109_
            lhs17_ = default__.safeIndex(343, (d_223_v109_).length(0))
            lhs18_ = d_153_cf33_
            lhs19_ = d_56_v6_
            lhs14_[lhs15_] = rhs31_
            lhs16_[lhs17_] = rhs32_
            lhs18_.f7 = rhs33_
            lhs19_.f7 = rhs34_
        hi0_ = (0) - (default__.safeModuloInt(p1, d_56_v6_.f7))
        for d_224_i7_ in range(p1, hi0_):
            index26_ = default__.safeIndex(482, (d_53_v3_).length(0))
            (d_53_v3_)[index26_] = d_224_i7_
            d_225_v110_: _dafny.Seq
            d_225_v110_ = _dafny.SeqWithoutIsStrInference([(d_56_v6_).f6])
            index27_ = default__.safeIndex(482, (d_53_v3_).length(0))
            (d_53_v3_)[index27_] = (p1) - (len(d_225_v110_))
            index28_ = default__.safeIndex(482, (d_53_v3_).length(0))
            (d_53_v3_)[index28_] = d_56_v6_.f7
            d_226_v111_: _dafny.Array
            nw32_ = _dafny.Array(None, 13)
            nw32_[int(0)] = d_56_v6_
            nw32_[int(1)] = d_56_v6_
            nw32_[int(2)] = d_56_v6_
            nw32_[int(3)] = d_56_v6_
            nw32_[int(4)] = d_56_v6_
            nw32_[int(5)] = d_56_v6_
            nw32_[int(6)] = d_56_v6_
            nw32_[int(7)] = (d_56_v6_ if d_50_v0_ else d_56_v6_)
            nw32_[int(8)] = d_56_v6_
            nw32_[int(9)] = d_56_v6_
            nw32_[int(10)] = d_56_v6_
            nw32_[int(11)] = d_56_v6_
            nw32_[int(12)] = d_56_v6_
            d_226_v111_ = nw32_
            index29_ = default__.safeIndex(974, (d_226_v111_).length(0))
            (d_226_v111_)[index29_] = d_56_v6_
            d_227_v112_: _dafny.Array
            nw33_ = _dafny.Array(False, 28)
            d_227_v112_ = nw33_
            d_228_v113_: _dafny.Map
            d_228_v113_ = _dafny.Map({d_227_v112_: _dafny.Map({d_224_i7_: not(not(d_50_v0_))})})
            index30_ = default__.safeIndex(974, (d_226_v111_).length(0))
            nw34_ = C4()
            nw34_.ctor__(d_50_v0_, d_228_v113_, (p1) * ((d_56_v6_).f6), default__.safeModuloInt((d_56_v6_).f6, default__.fm1(d_50_v0_, globalState)))
            (d_226_v111_)[index30_] = nw34_
            if d_50_v0_:
                d_229_v115_: _dafny.Set
                d_229_v115_ = _dafny.Set({(d_56_v6_).f6, d_56_v6_.f7})
                d_230_v116_: _dafny.Set
                def iife20_():
                    coll12_ = _dafny.Set()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(798, 348):
                        d_231_v114_: int = compr_12_
                        if ((798) <= (d_231_v114_)) and ((d_231_v114_) < (348)):
                            coll12_ = coll12_.union(_dafny.Set([(d_231_v114_) + ((d_53_v3_)[default__.safeIndex(482, (d_53_v3_).length(0))])]))
                    return _dafny.Set(coll12_)
                d_230_v116_ = _dafny.Set({not((iife20_()
                ).ispropersubset(d_229_v115_))})
                index31_ = default__.safeIndex(482, (d_53_v3_).length(0))
                (d_53_v3_)[index31_] = len(d_230_v116_)
                d_50_v0_ = (True) or (d_50_v0_)
                d_232_v117_: _dafny.Seq
                d_232_v117_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbbdvvot"))
                d_233_v118_: C0
                nw35_ = C0()
                nw35_.ctor__(d_50_v0_)
                d_233_v118_ = nw35_
                d_234_v119_: _dafny.Map
                d_234_v119_ = _dafny.Map({len(d_229_v115_): d_233_v118_})
                d_235_v120_: C1
                nw36_ = C1()
                nw36_.ctor__(default__.fm20(_dafny.SeqWithoutIsStrInference([d_55_v5_ for d_236_i8_ in range(default__.abs(831))]), False, default__.fm1(d_50_v0_, globalState), globalState), (len(d_232_v117_)) + (len(d_234_v119_)), (d_56_v6_).f6)
                d_235_v120_ = nw36_
                d_237_v121_: _dafny.Array
                nw37_ = _dafny.Array(None, 4)
                nw37_[int(0)] = (d_235_v120_).f13
                nw37_[int(1)] = (d_55_v5_ if d_50_v0_ else d_55_v5_)
                nw37_[int(2)] = (d_235_v120_).f13
                nw37_[int(3)] = (d_235_v120_).f13
                d_237_v121_ = nw37_
                d_238_v122_: _dafny.Array
                def lambda8_(d_239_v6_):
                    def lambda9_(d_240_i9_):
                        return (d_240_i9_) * ((d_239_v6_).f6)

                    return lambda9_

                init4_ = lambda8_(d_56_v6_)
                nw38_ = _dafny.Array(None, 7)
                for i0_4_ in range(nw38_.length(0)):
                    nw38_[i0_4_] = init4_(i0_4_)
                d_238_v122_ = nw38_
                d_241_v123_: _dafny.Set
                d_241_v123_ = _dafny.Set({default__.fm14(831, p1, d_56_v6_.f7, globalState), (d_55_v5_ if d_233_v118_.f10 else d_55_v5_), (d_235_v120_).f13, (d_235_v120_).f13})
                d_242_v124_: D2
                d_242_v124_ = D2_DC5(d_238_v122_)
                rhs35_ = d_237_v121_
                rhs36_ = d_50_v0_
                rhs37_ = (d_242_v124_).cf9
                rhs38_ = (d_241_v123_) - (d_241_v123_)
                d_237_v121_ = rhs35_
                d_50_v0_ = rhs36_
                d_238_v122_ = rhs37_
                d_241_v123_ = rhs38_
                d_243_v125_: C2
                nw39_ = C2()
                nw39_.ctor__(d_56_v6_.f7, len(d_232_v117_))
                d_243_v125_ = nw39_
            elif True:
                d_244_v126_: C0
                nw40_ = C0()
                nw40_.ctor__(d_50_v0_)
                d_244_v126_ = nw40_
                d_245_v127_: _dafny.Array
                nw41_ = _dafny.Array(None, 22)
                nw41_[int(0)] = d_244_v126_
                nw41_[int(1)] = d_244_v126_
                nw41_[int(2)] = d_244_v126_
                nw41_[int(3)] = d_244_v126_
                nw41_[int(4)] = d_244_v126_
                nw41_[int(5)] = d_244_v126_
                nw41_[int(6)] = d_244_v126_
                nw41_[int(7)] = d_244_v126_
                nw41_[int(8)] = d_244_v126_
                nw41_[int(9)] = d_244_v126_
                nw41_[int(10)] = d_244_v126_
                nw41_[int(11)] = d_244_v126_
                nw41_[int(12)] = d_244_v126_
                nw41_[int(13)] = d_244_v126_
                nw41_[int(14)] = d_244_v126_
                nw41_[int(15)] = d_244_v126_
                nw41_[int(16)] = d_244_v126_
                nw41_[int(17)] = d_244_v126_
                nw41_[int(18)] = d_244_v126_
                nw41_[int(19)] = d_244_v126_
                nw41_[int(20)] = d_244_v126_
                nw41_[int(21)] = d_244_v126_
                d_245_v127_ = nw41_
                d_246_v128_: _dafny.Seq
                d_246_v128_ = _dafny.SeqWithoutIsStrInference([d_244_v126_, d_244_v126_])
                d_247_v129_: _dafny.Seq
                d_247_v129_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrbqmsdid"))
                d_248_v130_: C0
                d_248_v130_ = (d_246_v128_)[default__.safeIndex(len(d_247_v129_), len(d_246_v128_))]
                index32_ = default__.safeIndex(26, (d_245_v127_).length(0))
                (d_245_v127_)[index32_] = (d_248_v130_)
                index33_ = default__.safeIndex(26, (d_245_v127_).length(0))
                (d_245_v127_)[index33_] = d_244_v126_
                r1 = not(d_50_v0_)
                d_249_v131_: _dafny.Array
                def lambda10_(d_250_v6_):
                    def lambda11_(d_251_i10_):
                        return default__.safeDivisionInt(d_251_i10_, (d_250_v6_).f6)

                    return lambda11_

                init5_ = lambda10_(d_56_v6_)
                nw42_ = _dafny.Array(None, 23)
                for i0_5_ in range(nw42_.length(0)):
                    nw42_[i0_5_] = init5_(i0_5_)
                d_249_v131_ = nw42_
                d_252_v132_: _dafny.Set
                d_252_v132_ = _dafny.Set({d_249_v131_})
                d_253_v133_: D8
                d_253_v133_ = D8_DC19(d_50_v0_, d_50_v0_, d_252_v132_, d_56_v6_.f7, 130)
                d_254_v134_: _dafny.Map
                d_254_v134_ = _dafny.Map({d_56_v6_: not((d_253_v133_).cf34)})
                d_255_v135_: D1
                d_255_v135_ = D1_DC3(d_50_v0_, default__.fm1(d_244_v126_.f10, globalState), (d_56_v6_).f6)
                index34_ = default__.safeIndex(482, (d_53_v3_).length(0))
                rhs39_ = (d_225_v110_)[default__.safeIndex((d_53_v3_)[default__.safeIndex(482, (d_53_v3_).length(0))], len(d_225_v110_))]
                rhs40_ = d_244_v126_.f10
                rhs41_ = (len(_dafny.SeqWithoutIsStrInference([(d_255_v135_).cf5])) if (d_50_v0_ if ((d_254_v134_)[d_56_v6_] if (d_56_v6_) in (d_254_v134_) else d_50_v0_) else d_50_v0_) else (d_56_v6_).f6)
                rhs42_ = default__.fm1(d_244_v126_.f10, globalState)
                lhs20_ = d_53_v3_
                lhs21_ = default__.safeIndex(482, (d_53_v3_).length(0))
                r0 = rhs39_
                r1 = rhs40_
                r0 = rhs41_
                lhs20_[lhs21_] = rhs42_
                nw43_ = _dafny.Array(int(0), 20)
                d_249_v131_ = nw43_
                (globalState).f1 = default__.safeModuloInt((-750) - ((d_56_v6_).f6), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wunqg")) if d_50_v0_ else d_247_v129_)))
        hi1_ = d_56_v6_.f7
        for d_256_i11_ in range((d_56_v6_.f7) - (d_56_v6_.f7), hi1_):
            d_257_v136_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_257_v136_ = nw44_
            d_258_v137_: _dafny.Seq
            d_258_v137_ = _dafny.SeqWithoutIsStrInference([d_256_i11_])
            d_259_v138_: _dafny.Map
            d_259_v138_ = _dafny.Map({d_56_v6_.f7: d_50_v0_})
            d_260_v139_: C5
            nw45_ = C5()
            nw45_.ctor__(d_257_v136_, d_258_v137_, (p1 if d_50_v0_ else p1), len(d_259_v138_))
            d_260_v139_ = nw45_
            d_260_v139_ = d_260_v139_
            d_261_v140_: C0
            nw46_ = C0()
            nw46_.ctor__(d_50_v0_)
            d_261_v140_ = nw46_
            d_262_v141_: _dafny.Array
            nw47_ = _dafny.Array(False, 2)
            d_262_v141_ = nw47_
            (d_260_v139_).m3(d_262_v141_, (d_56_v6_.f7) - (p1), globalState)
            d_263_v142_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.Set({}), 18)
            d_263_v142_ = nw48_
            index35_ = default__.safeIndex(603, (d_263_v142_).length(0))
            (d_263_v142_)[index35_] = (_dafny.Set({d_261_v140_.f10, d_50_v0_})).intersection(_dafny.Set({False, ((d_259_v138_)[d_256_i11_] if (d_256_i11_) in (d_259_v138_) else False)}))
            d_264_v143_: _dafny.Set
            d_264_v143_ = _dafny.Set({d_50_v0_, (d_50_v0_) or (True), (d_261_v140_.f10) and (d_50_v0_)})
            index36_ = default__.safeIndex(603, (d_263_v142_).length(0))
            (d_263_v142_)[index36_] = d_264_v143_
        d_265_v144_: _dafny.Seq
        d_265_v144_ = _dafny.SeqWithoutIsStrInference([False, d_50_v0_, d_50_v0_])
        d_266_v145_: _dafny.Map
        d_266_v145_ = _dafny.Map({d_265_v144_: not(d_50_v0_)})
        d_267_v146_: D5
        d_267_v146_ = D5_DC11(d_265_v144_)
        d_266_v145_ = (d_266_v145_).set((d_267_v146_).cf20, d_50_v0_)
        d_268_v147_: _dafny.Set
        d_268_v147_ = _dafny.Set({d_56_v6_.f7, d_56_v6_.f7})
        d_269_v148_: _dafny.Seq
        d_269_v148_ = _dafny.SeqWithoutIsStrInference([len(d_268_v147_), d_56_v6_.f7])
        r0 = default__.safeDivisionInt(d_56_v6_.f7, ((d_269_v148_)[default__.safeIndex((d_56_v6_).f6, len(d_269_v148_))]) - (p1))
        r1 = d_50_v0_
        d_270_v149_: _dafny.Array
        nw49_ = _dafny.Array(_dafny.Map({}), 8)
        d_270_v149_ = nw49_
        r2 = d_270_v149_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_271_globalState_: GlobalState
        nw50_ = GlobalState()
        nw50_.ctor__(True, 952, 374, 575, 555, False)
        d_271_globalState_ = nw50_
        d_272_v0_: int
        d_272_v0_ = -379
        (d_271_globalState_).f1 = default__.safeDivisionInt(d_272_v0_, (d_272_v0_) - (d_272_v0_))
        d_273_v1_: bool
        d_273_v1_ = False
        (d_271_globalState_).f1 = ((d_272_v0_) * (d_272_v0_) if d_273_v1_ else d_272_v0_)
        d_274_v2_: _dafny.Array
        def lambda12_(d_275_i0_):
            return not(True)

        init6_ = lambda12_
        nw51_ = _dafny.Array(None, 17)
        for i0_6_ in range(nw51_.length(0)):
            nw51_[i0_6_] = init6_(i0_6_)
        d_274_v2_ = nw51_
        index37_ = default__.safeIndex(673, (d_274_v2_).length(0))
        (d_274_v2_)[index37_] = d_273_v1_
        d_276_v3_: _dafny.Set
        d_276_v3_ = _dafny.Set({d_273_v1_})
        d_277_v4_: _dafny.Map
        d_277_v4_ = _dafny.Map({(d_276_v3_) | (d_276_v3_): d_273_v1_})
        d_278_v5_: _dafny.Array
        nw52_ = _dafny.Array(int(0), 21)
        d_278_v5_ = nw52_
        d_279_v6_: _dafny.Set
        d_279_v6_ = _dafny.Set({d_278_v5_, d_278_v5_})
        d_280_v7_: _dafny.Seq
        d_280_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klhmhabdk"))
        d_281_v8_: _dafny.MultiSet
        d_281_v8_ = _dafny.MultiSet([d_280_v7_])
        d_282_v9_: _dafny.Map
        d_282_v9_ = _dafny.Map({d_281_v8_: d_279_v6_})
        index38_ = default__.safeIndex(673, (d_274_v2_).length(0))
        rhs43_ = d_273_v1_
        rhs44_ = d_273_v1_
        rhs45_ = _dafny.Map({_dafny.Set({d_273_v1_, d_273_v1_}): d_273_v1_})
        rhs46_ = (d_279_v6_).isdisjoint(((d_282_v9_)[_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_280_v7_, d_280_v7_, d_280_v7_, d_280_v7_]))] if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_280_v7_, d_280_v7_, d_280_v7_, d_280_v7_]))) in (d_282_v9_) else d_279_v6_))
        lhs22_ = d_274_v2_
        lhs23_ = default__.safeIndex(673, (d_274_v2_).length(0))
        d_273_v1_ = rhs43_
        lhs22_[lhs23_] = rhs44_
        d_277_v4_ = rhs45_
        d_273_v1_ = rhs46_
        hi2_ = d_272_v0_
        for d_283_i1_ in range((d_272_v0_) + (d_272_v0_), hi2_):
            index39_ = default__.safeIndex(990, (d_278_v5_).length(0))
            (d_278_v5_)[index39_] = d_283_i1_
            index40_ = default__.safeIndex(990, (d_278_v5_).length(0))
            (d_278_v5_)[index40_] = d_283_i1_
            d_284_v10_: _dafny.Map
            d_284_v10_ = _dafny.Map({(d_278_v5_)[default__.safeIndex(990, (d_278_v5_).length(0))]: d_283_i1_})
            d_285_v11_: _dafny.Map
            d_285_v11_ = _dafny.Map({False: default__.fm0((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_284_v10_, (0) - (d_283_i1_), d_271_globalState_)})
            d_285_v11_ = (d_285_v11_).set(d_273_v1_, ((d_278_v5_)[default__.safeIndex(990, (d_278_v5_).length(0))]) > (d_283_i1_))
            index41_ = default__.safeIndex(990, (d_278_v5_).length(0))
            (d_278_v5_)[index41_] = d_272_v0_
            (d_271_globalState_).f1 = d_272_v0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_278_v5_).length(0)):
            d_286_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_286_i2_)) and ((d_286_i2_) < ((d_278_v5_).length(0)))):
                (d_278_v5_)[(d_286_i2_)] = (d_286_i2_) - (d_272_v0_)
        d_287_v12_: _dafny.Map
        d_287_v12_ = _dafny.Map({d_272_v0_: d_272_v0_})
        if ((d_273_v1_) and (default__.fm0((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_287_v12_, d_272_v0_, d_271_globalState_)) if (d_273_v1_ if d_273_v1_ else True) else (d_272_v0_) == ((0) - (d_272_v0_))):
            d_288_v13_: D0
            d_288_v13_ = D0_DC0(d_276_v3_)
            d_289_v14_: _dafny.MultiSet
            d_289_v14_ = _dafny.MultiSet([d_276_v3_, d_276_v3_, (d_288_v13_).cf0])
            index42_ = default__.safeIndex(30, (d_278_v5_).length(0))
            (d_278_v5_)[index42_] = default__.safeDivisionInt(d_272_v0_, ((d_289_v14_)[d_276_v3_] if (d_276_v3_) in (d_289_v14_) else 996))
            index43_ = default__.safeIndex(30, (d_278_v5_).length(0))
            index44_ = default__.safeIndex(673, (d_274_v2_).length(0))
            index45_ = default__.safeIndex(673, (d_274_v2_).length(0))
            rhs47_ = (d_272_v0_) - (d_272_v0_)
            rhs48_ = (d_272_v0_) + (d_272_v0_)
            rhs49_ = d_273_v1_
            rhs50_ = d_274_v2_
            rhs51_ = d_273_v1_
            lhs24_ = d_278_v5_
            lhs25_ = default__.safeIndex(30, (d_278_v5_).length(0))
            lhs26_ = d_271_globalState_
            lhs27_ = d_274_v2_
            lhs28_ = default__.safeIndex(673, (d_274_v2_).length(0))
            lhs29_ = d_274_v2_
            lhs30_ = default__.safeIndex(673, (d_274_v2_).length(0))
            lhs24_[lhs25_] = rhs47_
            lhs26_.f1 = rhs48_
            lhs27_[lhs28_] = rhs49_
            d_274_v2_ = rhs50_
            lhs29_[lhs30_] = rhs51_
            d_290_v15_: _dafny.Array
            nw53_ = _dafny.Array(_dafny.Seq({}), 8)
            d_290_v15_ = nw53_
            index46_ = default__.safeIndex(659, (d_290_v15_).length(0))
            (d_290_v15_)[index46_] = _dafny.SeqWithoutIsStrInference([(d_278_v5_)[default__.safeIndex(30, (d_278_v5_).length(0))], default__.fm1(d_273_v1_, d_271_globalState_)])
            d_291_v16_: _dafny.Seq
            d_291_v16_ = _dafny.SeqWithoutIsStrInference([d_272_v0_])
            index47_ = default__.safeIndex(659, (d_290_v15_).length(0))
            (d_290_v15_)[index47_] = d_291_v16_
            d_292_v17_: D0
            d_292_v17_ = D0_DC1(d_273_v1_, d_274_v2_, (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))])
            pat_let_tv3_ = d_274_v2_
            pat_let_tv4_ = d_274_v2_
            pat_let_tv5_ = d_274_v2_
            d_293_v18_: int
            d_294_v19_: bool
            d_295_v20_: _dafny.Array
            out0_: int
            out1_: bool
            out2_: _dafny.Array
            def iife22_(_pat_let5_0):
                def iife23_(d_296_dt__update__tmp_h0_):
                    def iife24_(_pat_let6_0):
                        def iife25_(d_297_dt__update_hcf2_h0_):
                            def iife26_(_pat_let7_0):
                                def iife27_(d_298_dt__update_hcf3_h0_):
                                    return D0_DC1((d_296_dt__update__tmp_h0_).cf1, d_297_dt__update_hcf2_h0_, d_298_dt__update_hcf3_h0_)
                                return iife27_(_pat_let7_0)
                            return iife26_(not((pat_let_tv5_)[default__.safeIndex(673, (pat_let_tv4_).length(0))]))
                        return iife25_(_pat_let6_0)
                    return iife24_(pat_let_tv3_)
                return iife23_(_pat_let5_0)
            def iife21_(_pat_let4_0):
                def iife28_(d_299_dt__update__tmp_h1_):
                    def iife29_(_pat_let8_0):
                        def iife30_(d_300_dt__update_hcf3_h1_):
                            return D0_DC1((d_299_dt__update__tmp_h1_).cf1, (d_299_dt__update__tmp_h1_).cf2, d_300_dt__update_hcf3_h1_)
                        return iife30_(_pat_let8_0)
                    return iife29_(False)
                return iife28_(_pat_let4_0)
            out0_, out1_, out2_ = default__.m0(iife21_(iife22_(d_292_v17_)), (d_278_v5_)[default__.safeIndex(30, (d_278_v5_).length(0))], d_271_globalState_)
            d_293_v18_ = out0_
            d_294_v19_ = out1_
            d_295_v20_ = out2_
            d_292_v17_ = d_292_v17_
            d_301_v21_: int
            d_302_v22_: bool
            d_303_v23_: _dafny.Array
            out3_: int
            out4_: bool
            out5_: _dafny.Array
            out3_, out4_, out5_ = default__.m0(d_292_v17_, d_293_v18_, d_271_globalState_)
            d_301_v21_ = out3_
            d_302_v22_ = out4_
            d_303_v23_ = out5_
        elif True:
            d_280_v7_ = default__.fm2(_dafny.SeqWithoutIsStrInference([d_272_v0_, default__.fm1(default__.fm0((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_287_v12_, (0) - (d_272_v0_), d_271_globalState_), d_271_globalState_)]), default__.safeDivisionInt(default__.fm1(d_273_v1_, d_271_globalState_), d_272_v0_), d_272_v0_, not(d_273_v1_), d_271_globalState_)
            d_304_v24_: _dafny.MultiSet
            d_304_v24_ = _dafny.MultiSet([d_272_v0_, 556])
            d_273_v1_ = ((d_304_v24_).ispropersubset(d_304_v24_)) or (d_273_v1_)
            index48_ = default__.safeIndex(673, (d_274_v2_).length(0))
            (d_274_v2_)[index48_] = not ((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]) or ((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))])
            (d_271_globalState_).f1 = d_272_v0_
            d_305_v25_: _dafny.MultiSet
            d_305_v25_ = _dafny.MultiSet([(d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]])
            d_306_v26_: _dafny.Map
            d_306_v26_ = _dafny.Map({d_273_v1_: ((d_305_v25_)[d_273_v1_] if (d_273_v1_) in (d_305_v25_) else d_272_v0_)})
            d_307_v27_: _dafny.Map
            d_307_v27_ = _dafny.Map({d_280_v7_: len(d_306_v26_)})
            d_308_v28_: D0
            d_308_v28_ = D0_DC0(_dafny.Set({d_273_v1_}))
            d_309_v29_: _dafny.MultiSet
            d_309_v29_ = _dafny.MultiSet([default__.fm3(default__.fm0(d_273_v1_, d_287_v12_, d_272_v0_, d_271_globalState_), d_272_v0_, d_276_v3_, d_307_v27_, d_271_globalState_), d_308_v28_, D0_DC0(d_276_v3_), d_308_v28_, d_308_v28_])
            (d_271_globalState_).f1 = (d_309_v29_).cardinality
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_274_v2_).length(0)):
            d_310_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_310_i3_)) and ((d_310_i3_) < ((d_274_v2_).length(0)))):
                (d_274_v2_)[(d_310_i3_)] = False
        index49_ = default__.safeIndex(71, (d_278_v5_).length(0))
        (d_278_v5_)[index49_] = len(d_280_v7_)
        d_311_v30_: _dafny.Set
        d_311_v30_ = _dafny.Set({d_272_v0_})
        d_312_v31_: _dafny.Map
        d_312_v31_ = _dafny.Map({d_311_v30_: d_272_v0_})
        d_313_v32_: _dafny.Seq
        d_313_v32_ = _dafny.SeqWithoutIsStrInference([(0) - (d_272_v0_), 405, 226, len(d_312_v31_)])
        index50_ = default__.safeIndex(71, (d_278_v5_).length(0))
        (d_278_v5_)[index50_] = len(d_313_v32_)
        d_314_v33_: _dafny.Map
        d_314_v33_ = _dafny.Map({d_272_v0_: (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]})
        d_315_v34_: _dafny.Map
        d_315_v34_ = _dafny.Map({d_274_v2_: d_314_v33_})
        d_316_v35_: C4
        nw54_ = C4()
        nw54_.ctor__(d_273_v1_, d_315_v34_, 745, (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])
        d_316_v35_ = nw54_
        (d_271_globalState_).f1 = -937
        d_317_v37_: D7
        def iife31_():
            coll13_ = _dafny.Map()
            compr_13_: str
            for compr_13_ in (d_280_v7_).Elements:
                d_318_v36_: str = compr_13_
                if (d_318_v36_) in (d_280_v7_):
                    coll13_[d_318_v36_] = d_272_v0_
            return _dafny.Map(coll13_)
        d_317_v37_ = D7_DC16(iife31_()
)
        d_319_v38_: _dafny.Map
        d_319_v38_ = _dafny.Map({(d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]: d_317_v37_})
        d_319_v38_ = d_319_v38_
        d_320_v39_: _dafny.MultiSet
        d_320_v39_ = _dafny.MultiSet([(d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], d_272_v0_])
        hi3_ = (d_320_v39_).cardinality
        for d_321_i4_ in range((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], hi3_):
            index51_ = default__.safeIndex(71, (d_278_v5_).length(0))
            (d_278_v5_)[index51_] = (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]
            d_322_v40_: _dafny.Seq
            d_322_v40_ = _dafny.SeqWithoutIsStrInference([(d_316_v35_).f11, (d_316_v35_).f11])
            d_323_v41_: str
            d_323_v41_ = _dafny.CodePoint('h')
            d_324_v42_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_324_v42_ = nw55_
            d_325_v43_: _dafny.Map
            d_325_v43_ = _dafny.Map({True: (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]})
            d_326_v44_: C5
            nw56_ = C5()
            nw56_.ctor__(d_324_v42_, _dafny.SeqWithoutIsStrInference([len(d_325_v43_), (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]]), d_321_i4_, (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])
            d_326_v44_ = nw56_
            d_327_v45_: _dafny.MultiSet
            d_327_v45_ = _dafny.MultiSet([d_326_v44_])
            d_328_v46_: D5
            d_328_v46_ = D5_DC11(d_322_v40_)
            d_329_v47_: D8
            d_329_v47_ = D8_DC19(d_273_v1_, (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_279_v6_, d_272_v0_, (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])
            d_330_v48_: _dafny.Array
            nw57_ = _dafny.Array(None, 29)
            nw57_[int(0)] = (d_322_v40_) + (d_322_v40_)
            nw57_[int(1)] = d_322_v40_
            nw57_[int(2)] = d_322_v40_
            nw57_[int(3)] = d_322_v40_
            nw57_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]])
            nw57_[int(5)] = d_322_v40_
            nw57_[int(6)] = (d_322_v40_ if d_273_v1_ else (_dafny.SeqWithoutIsStrInference([(d_316_v35_).f11, (d_316_v35_).f11])).set(default__.safeIndex(d_272_v0_, len(_dafny.SeqWithoutIsStrInference([(d_316_v35_).f11, (d_316_v35_).f11]))), d_273_v1_))
            nw57_[int(7)] = d_322_v40_
            nw57_[int(8)] = d_322_v40_
            nw57_[int(9)] = (default__.fm27(-187, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_331_i5_ in range(default__.abs(-901))]), d_323_v41_, d_271_globalState_)).set(default__.safeIndex(((d_327_v45_)[d_326_v44_] if (d_326_v44_) in (d_327_v45_) else 538), len(default__.fm27(-187, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_332_i5_ in range(default__.abs(-901))]), d_323_v41_, d_271_globalState_))), (d_316_v35_).f11)
            nw57_[int(10)] = d_322_v40_
            nw57_[int(11)] = d_322_v40_
            nw57_[int(12)] = d_322_v40_
            nw57_[int(13)] = d_322_v40_
            nw57_[int(14)] = d_322_v40_
            nw57_[int(15)] = d_322_v40_
            nw57_[int(16)] = (d_328_v46_).cf20
            nw57_[int(17)] = d_322_v40_
            nw57_[int(18)] = (d_322_v40_) + (d_322_v40_)
            nw57_[int(19)] = (d_322_v40_) + (d_322_v40_)
            nw57_[int(20)] = d_322_v40_
            nw57_[int(21)] = (d_322_v40_) + (d_322_v40_)
            nw57_[int(22)] = default__.fm21(d_271_globalState_)
            nw57_[int(23)] = d_322_v40_
            nw57_[int(24)] = (d_322_v40_) + (_dafny.SeqWithoutIsStrInference([(d_329_v47_).cf34]))
            nw57_[int(25)] = _dafny.SeqWithoutIsStrInference([(d_316_v35_).f11])
            nw57_[int(26)] = d_322_v40_
            nw57_[int(27)] = (d_322_v40_ if (d_316_v35_).f11 else d_322_v40_)
            nw57_[int(28)] = (d_328_v46_).cf20
            d_330_v48_ = nw57_
            index52_ = default__.safeIndex(592, (d_330_v48_).length(0))
            (d_330_v48_)[index52_] = d_322_v40_
            index53_ = default__.safeIndex(592, (d_330_v48_).length(0))
            (d_330_v48_)[index53_] = d_322_v40_
            d_333_v49_: D0
            d_333_v49_ = D0_DC0(d_276_v3_)
            d_334_v50_: _dafny.Map
            d_334_v50_ = _dafny.Map({d_333_v49_: d_273_v1_})
            d_335_v51_: D1
            d_335_v51_ = D1_DC2(d_334_v50_)
            d_336_v52_: D2
            d_336_v52_ = D2_DC6(d_335_v51_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vku")), (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])
            d_337_v53_: _dafny.Map
            d_337_v53_ = _dafny.Map({d_273_v1_: (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]})
            d_338_v54_: C1
            nw58_ = C1()
            nw58_.ctor__(d_323_v41_, d_321_i4_, ((d_337_v53_)[(d_316_v35_).f11] if ((d_316_v35_).f11) in (d_337_v53_) else len(d_313_v32_)))
            d_338_v54_ = nw58_
            d_339_v55_: _dafny.Map
            d_339_v55_ = _dafny.Map({d_338_v54_: d_280_v7_})
            d_280_v7_ = ((d_336_v52_).cf11) + (((d_339_v55_)[d_338_v54_] if (d_338_v54_) in (d_339_v55_) else d_280_v7_))
            def lambda13_(d_340_v2_):
                def lambda14_(d_341_i6_):
                    return (d_340_v2_)[default__.safeIndex(673, (d_340_v2_).length(0))]

                return lambda14_

            init7_ = lambda13_(d_274_v2_)
            nw59_ = _dafny.Array(None, 22)
            for i0_7_ in range(nw59_.length(0)):
                nw59_[i0_7_] = init7_(i0_7_)
            d_274_v2_ = nw59_
        d_342_i7_: int
        d_342_i7_ = 0
        with _dafny.label("0"):
            while (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]:
                with _dafny.c_label("0"):
                    if (d_342_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_342_i7_ = (d_342_i7_) + (1)
                    (d_271_globalState_).f1 = (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]
                    d_343_v56_: _dafny.Array
                    def lambda15_(d_344_v35_):
                        def lambda16_(d_345_i8_):
                            return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([(d_344_v35_).f11, (d_344_v35_).f11]))

                        return lambda16_

                    init8_ = lambda15_(d_316_v35_)
                    nw60_ = _dafny.Array(None, 14)
                    for i0_8_ in range(nw60_.length(0)):
                        nw60_[i0_8_] = init8_(i0_8_)
                    d_343_v56_ = nw60_
                    d_346_v57_: _dafny.Seq
                    d_346_v57_ = _dafny.SeqWithoutIsStrInference([True, d_273_v1_])
                    index54_ = default__.safeIndex(535, (d_343_v56_).length(0))
                    (d_343_v56_)[index54_] = (d_346_v57_) + (d_346_v57_)
                    d_347_v58_: D1
                    d_347_v58_ = D1_DC4(d_273_v1_)
                    d_348_v59_: D2
                    d_348_v59_ = D2_DC7((d_316_v35_).f11, len(d_276_v3_), d_314_v33_, d_347_v58_)
                    index55_ = default__.safeIndex(535, (d_343_v56_).length(0))
                    (d_343_v56_)[index55_] = ((d_346_v57_ if not(default__.fm0((d_348_v59_).cf13, d_287_v12_, len(d_280_v7_), d_271_globalState_)) else d_346_v57_)) + ((_dafny.SeqWithoutIsStrInference([d_273_v1_, ((d_314_v33_)[d_272_v0_] if (d_272_v0_) in (d_314_v33_) else (d_316_v35_).f11)])).set(default__.safeIndex(d_272_v0_, len(_dafny.SeqWithoutIsStrInference([d_273_v1_, ((d_314_v33_)[d_272_v0_] if (d_272_v0_) in (d_314_v33_) else (d_316_v35_).f11)]))), True))
                    d_349_v60_: C0
                    nw61_ = C0()
                    nw61_.ctor__((d_311_v30_).issubset(_dafny.Set({(d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]})))
                    d_349_v60_ = nw61_
                    d_350_v61_: T0
                    nw62_ = C4()
                    nw62_.ctor__((d_316_v35_).f11, d_316_v35_.f12, (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], default__.fm1((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_271_globalState_))
                    d_350_v61_ = nw62_
                    d_351_v62_: _dafny.MultiSet
                    d_351_v62_ = _dafny.MultiSet([d_350_v61_, d_350_v61_, d_350_v61_, d_350_v61_])
                    d_352_v63_: _dafny.Map
                    d_352_v63_ = _dafny.Map({d_351_v62_: False})
                    d_353_v64_: T0
                    nw63_ = C4()
                    nw63_.ctor__(default__.fm0((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_287_v12_, len((d_346_v57_).set(default__.safeIndex(d_272_v0_, len(d_346_v57_)), (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))])), d_271_globalState_), d_315_v34_, d_272_v0_, len((d_352_v63_) | (d_352_v63_)))
                    d_353_v64_ = nw63_
                    d_354_v65_: str
                    d_354_v65_ = _dafny.CodePoint('k')
                    rhs52_ = (d_280_v7_) + (_dafny.SeqWithoutIsStrInference([d_354_v65_, default__.fm8(_dafny.SeqWithoutIsStrInference([d_354_v65_ for d_355_i9_ in range(default__.abs(828))]), _dafny.CodePoint('p'), False, (d_316_v35_).f11, d_271_globalState_), d_354_v65_, d_354_v65_]))
                    rhs53_ = d_350_v61_
                    d_280_v7_ = rhs52_
                    d_353_v64_ = rhs53_
                    pass
            pass
        index56_ = default__.safeIndex(673, (d_274_v2_).length(0))
        (d_274_v2_)[index56_] = (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]
        d_356_v66_: D0
        d_356_v66_ = D0_DC0(d_276_v3_)
        d_357_v67_: _dafny.Map
        d_357_v67_ = _dafny.Map({d_356_v66_: (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]})
        d_358_v68_: D1
        d_358_v68_ = D1_DC2(d_357_v67_)
        source5_ = d_358_v68_
        if source5_.is_DC3:
            d_359___mcc_h0_ = source5_.cf5
            d_360___mcc_h1_ = source5_.cf6
            d_361___mcc_h2_ = source5_.cf7
            d_362_cf7_ = d_361___mcc_h2_
            d_363_cf6_ = d_360___mcc_h1_
            d_364_cf5_ = d_359___mcc_h0_
            d_365_v69_: _dafny.Map
            d_365_v69_ = _dafny.Map({d_364_cf5_: default__.fm1(False, d_271_globalState_)})
            d_366_v70_: bool
            d_367_v71_: bool
            d_368_v72_: int
            out6_: bool
            out7_: bool
            out8_: int
            out6_, out7_, out8_ = (d_316_v35_).m1(not(not((d_362_cf7_) <= (((d_365_v69_)[False] if (False) in (d_365_v69_) else d_363_cf6_)))), len(d_280_v7_), 103, d_271_globalState_)
            d_366_v70_ = out6_
            d_367_v71_ = out7_
            d_368_v72_ = out8_
            d_369_v73_: _dafny.Seq
            d_369_v73_ = _dafny.SeqWithoutIsStrInference([default__.fm0(not(d_364_cf5_), d_287_v12_, d_363_cf6_, d_271_globalState_), d_273_v1_, d_366_v70_])
            d_370_v74_: _dafny.Seq
            d_370_v74_ = _dafny.SeqWithoutIsStrInference([d_311_v30_, d_311_v30_])
            d_371_v75_: bool
            d_372_v76_: bool
            d_373_v77_: int
            out9_: bool
            out10_: bool
            out11_: int
            out9_, out10_, out11_ = (d_316_v35_).m1((d_369_v73_)[default__.safeIndex(d_272_v0_, len(d_369_v73_))], (0) - ((0) - ((len(d_370_v74_)) - (d_363_cf6_))), len(_dafny.Map({d_362_cf7_: d_363_cf6_})), d_271_globalState_)
            d_371_v75_ = out9_
            d_372_v76_ = out10_
            d_373_v77_ = out11_
            (d_316_v35_).m5((0) - (((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))] if d_364_cf5_ else d_272_v0_)), d_271_globalState_)
            d_374_v78_: str
            d_374_v78_ = _dafny.CodePoint('q')
            d_375_v79_: _dafny.Set
            d_375_v79_ = _dafny.Set({_dafny.MultiSet([d_374_v78_])})
            d_376_v80_: _dafny.MultiSet
            d_376_v80_ = _dafny.MultiSet([d_374_v78_, d_374_v78_, d_374_v78_, d_374_v78_, d_374_v78_])
            d_375_v79_ = _dafny.Set({d_376_v80_, d_376_v80_})
        elif source5_.is_DC4:
            d_377___mcc_h3_ = source5_.cf8
            d_378_cf8_ = d_377___mcc_h3_
            (d_316_v35_).m5(d_272_v0_, d_271_globalState_)
            d_379_v81_: C2
            nw64_ = C2()
            nw64_.ctor__((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], d_272_v0_)
            d_379_v81_ = nw64_
            d_274_v2_ = d_274_v2_
            d_380_v82_: _dafny.MultiSet
            d_380_v82_ = _dafny.MultiSet([not(False)])
            if ((((d_380_v82_).set(d_273_v1_, default__.abs((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])) if d_378_cf8_ else (d_380_v82_).set((d_316_v35_).f11, default__.abs(d_272_v0_)))).cardinality) > ((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]):
                (d_316_v35_).m5(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdrbwar"))) + (d_280_v7_)), d_271_globalState_)
                (d_271_globalState_).f1 = (785 if not(d_378_cf8_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmolciku"))))
                d_381_v83_: _dafny.Seq
                d_381_v83_ = _dafny.SeqWithoutIsStrInference([d_274_v2_])
                d_274_v2_ = (d_381_v83_)[default__.safeIndex((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], len(d_381_v83_))]
                d_382_v84_: _dafny.Array
                nw65_ = _dafny.Array(None, 14)
                d_382_v84_ = nw65_
                d_383_v85_: _dafny.Seq
                d_383_v85_ = _dafny.SeqWithoutIsStrInference([d_382_v84_, d_382_v84_, d_382_v84_])
                d_382_v84_ = (d_383_v85_)[default__.safeIndex((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], len(d_383_v85_))]
                index57_ = default__.safeIndex(673, (d_274_v2_).length(0))
                (d_274_v2_)[index57_] = (d_316_v35_).f11
            elif True:
                d_384_v86_: C0
                nw66_ = C0()
                nw66_.ctor__((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))])
                d_384_v86_ = nw66_
                d_385_v87_: C0
                d_385_v87_ = d_384_v86_
                d_385_v87_ = d_385_v87_
                d_378_cf8_ = (d_316_v35_).f11
                d_386_v88_: _dafny.Seq
                d_386_v88_ = _dafny.SeqWithoutIsStrInference([(d_378_cf8_) or (not(d_384_v86_.f10)), d_378_cf8_])
                d_386_v88_ = (d_386_v88_) + (d_386_v88_)
                d_387_v89_: D1
                d_387_v89_ = D1_DC3((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], (0) - (d_272_v0_), (d_320_v39_).cardinality)
                index58_ = default__.safeIndex(673, (d_274_v2_).length(0))
                index59_ = default__.safeIndex(71, (d_278_v5_).length(0))
                rhs54_ = ((len(d_311_v30_)) <= (-192) if d_384_v86_.f10 else (d_316_v35_).f11)
                rhs55_ = (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]
                rhs56_ = default__.safeDivisionInt((d_272_v0_) - (len(d_386_v88_)), (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))])
                rhs57_ = ((len(d_314_v33_)) * (d_272_v0_)) == (d_272_v0_)
                rhs58_ = (d_387_v89_).cf5
                lhs31_ = d_274_v2_
                lhs32_ = default__.safeIndex(673, (d_274_v2_).length(0))
                lhs33_ = d_278_v5_
                lhs34_ = default__.safeIndex(71, (d_278_v5_).length(0))
                d_273_v1_ = rhs54_
                lhs31_[lhs32_] = rhs55_
                lhs33_[lhs34_] = rhs56_
                d_378_cf8_ = rhs57_
                d_378_cf8_ = rhs58_
                (d_271_globalState_).f1 = default__.safeModuloInt(d_272_v0_, ((d_287_v12_)[(d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]] if ((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]) in (d_287_v12_) else d_272_v0_))
        elif True:
            d_388___mcc_h4_ = source5_.cf4
            d_389_cf4_ = d_388___mcc_h4_
            d_390_v90_: _dafny.Seq
            d_390_v90_ = _dafny.SeqWithoutIsStrInference([d_274_v2_, d_274_v2_])
            d_391_v91_: _dafny.Array
            nw67_ = _dafny.Array(None, 12)
            nw67_[int(0)] = d_274_v2_
            nw67_[int(1)] = (d_390_v90_)[default__.safeIndex((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], len(d_390_v90_))]
            nw67_[int(2)] = d_274_v2_
            nw67_[int(3)] = d_274_v2_
            nw67_[int(4)] = (D0_DC1((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_274_v2_, (d_316_v35_).f11)).cf2
            nw67_[int(5)] = d_274_v2_
            nw67_[int(6)] = d_274_v2_
            nw67_[int(7)] = d_274_v2_
            nw67_[int(8)] = d_274_v2_
            nw67_[int(9)] = d_274_v2_
            nw67_[int(10)] = d_274_v2_
            nw67_[int(11)] = d_274_v2_
            d_391_v91_ = nw67_
            index60_ = default__.safeIndex(615, (d_391_v91_).length(0))
            (d_391_v91_)[index60_] = (d_274_v2_ if (d_316_v35_).f11 else d_274_v2_)
            index61_ = default__.safeIndex(615, (d_391_v91_).length(0))
            (d_391_v91_)[index61_] = d_274_v2_
            d_287_v12_ = (d_287_v12_).set((0) - ((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]), (len(d_280_v7_)) + ((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]))
            index62_ = default__.safeIndex(673, (d_274_v2_).length(0))
            (d_274_v2_)[index62_] = d_273_v1_
            d_272_v0_ = (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]
        if (d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))]:
            d_316_v35_ = d_316_v35_
            d_392_v92_: str
            d_392_v92_ = _dafny.CodePoint('p')
            d_392_v92_ = (d_280_v7_)[default__.safeIndex(default__.safeDivisionInt(d_272_v0_, len(d_276_v3_)), len(d_280_v7_))]
            d_287_v12_ = (d_287_v12_).set(default__.fm1((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_271_globalState_), default__.fm1((d_274_v2_)[default__.safeIndex(673, (d_274_v2_).length(0))], d_271_globalState_))
            d_272_v0_ = default__.fm1(d_273_v1_, d_271_globalState_)
            d_393_v94_: _dafny.Array
            nw68_ = _dafny.Array(None, 6)
            nw68_[int(0)] = d_311_v30_
            nw68_[int(1)] = d_311_v30_
            nw68_[int(2)] = d_311_v30_
            nw68_[int(3)] = _dafny.Set({d_272_v0_, d_272_v0_, d_272_v0_, (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]})
            def iife32_():
                coll14_ = _dafny.Set()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(877, 434):
                    d_394_v93_: int = compr_14_
                    if ((877) <= (d_394_v93_)) and ((d_394_v93_) < (434)):
                        coll14_ = coll14_.union(_dafny.Set([(d_394_v93_) * (d_272_v0_)]))
                return _dafny.Set(coll14_)
            nw68_[int(4)] = iife32_()
            
            nw68_[int(5)] = _dafny.Set({d_272_v0_, d_272_v0_, d_272_v0_, d_272_v0_, d_272_v0_})
            d_393_v94_ = nw68_
            d_395_v95_: D11
            d_395_v95_ = D11_DC28(d_280_v7_, d_393_v94_, (d_316_v35_).f11, d_273_v1_)
            d_396_v96_: D11
            d_396_v96_ = D11_DC29(d_395_v95_)
            d_396_v96_ = d_396_v96_
        elif True:
            d_397_v97_: _dafny.Set
            d_397_v97_ = _dafny.Set({d_274_v2_})
            d_398_v98_: _dafny.Map
            d_398_v98_ = _dafny.Map({d_273_v1_: (d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))]})
            index63_ = default__.safeIndex(673, (d_274_v2_).length(0))
            (d_274_v2_)[index63_] = not(not(default__.fm0((d_316_v35_).f11, (d_287_v12_).set(((d_320_v39_).set(d_272_v0_, default__.abs(len(d_397_v97_)))).cardinality, len(d_398_v98_)), d_272_v0_, d_271_globalState_)))
            d_399_v99_: C0
            nw69_ = C0()
            nw69_.ctor__((d_316_v35_).f11)
            d_399_v99_ = nw69_
            d_400_v100_: C0
            d_400_v100_ = d_399_v99_
            d_400_v100_ = d_400_v100_
            (d_399_v99_).f10 = default__.fm0(default__.fm0((d_316_v35_).f11, d_287_v12_, d_272_v0_, d_271_globalState_), d_287_v12_, len(default__.fm40((d_278_v5_)[default__.safeIndex(71, (d_278_v5_).length(0))], d_399_v99_.f10, d_271_globalState_)), d_271_globalState_)
            (d_271_globalState_).f1 = (0) - (d_272_v0_)
            (d_271_globalState_).f1 = -252
        _dafny.print(_dafny.string_of((d_271_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_272_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_273_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v2_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v3_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_v4_) == (_dafny.Map({_dafny.Set({False}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_279_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_280_v7_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_281_v8_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klhmhabdk"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_282_v9_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_287_v12_) == (_dafny.Map({-379: -379, -4: 567}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_311_v30_) == (_dafny.Set({-379}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_312_v31_) == (_dafny.Map({_dafny.Set({-379}): -379}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_313_v32_) == (_dafny.SeqWithoutIsStrInference([379, 405, 226, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_314_v33_) == (_dafny.Map({-379: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_315_v34_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_316_v35_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_316_v35_.f12)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_317_v37_).cf30) == (_dafny.Map({_dafny.CodePoint('i'): -379, _dafny.CodePoint('r'): -379, _dafny.CodePoint('e'): -379, _dafny.CodePoint('n'): -379, _dafny.CodePoint('x'): -379, _dafny.CodePoint('w'): -379, _dafny.CodePoint('f'): -379, _dafny.CodePoint('h'): -379, _dafny.CodePoint('c'): -379}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_319_v38_) == (_dafny.Map({False: D7_DC16(_dafny.Map({_dafny.CodePoint('i'): -379, _dafny.CodePoint('r'): -379, _dafny.CodePoint('e'): -379, _dafny.CodePoint('n'): -379, _dafny.CodePoint('x'): -379, _dafny.CodePoint('w'): -379, _dafny.CodePoint('f'): -379, _dafny.CodePoint('h'): -379, _dafny.CodePoint('c'): -379}))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_320_v39_) == (_dafny.MultiSet([4, -379]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_342_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_356_v66_).cf0) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_357_v67_) == (_dafny.Map({D0_DC0(_dafny.Set({False})): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_358_v68_).cf4) == (_dafny.Map({D0_DC0(_dafny.Set({False})): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
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
        return lambda: D1_DC3(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(D1.default()(), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)}, {self.cf11.VerbatimString(True)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC10(D4, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC12(D5, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(int(0), _dafny.MultiSet({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({self.cf31.VerbatimString(True)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(False, False, _dafny.Set({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf34', Any), ('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf39', Any), ('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(False, int(0), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC25()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC25(D10, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D11_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC28(D11, NamedTuple('DC28', [('cf54', Any), ('cf55', Any), ('cf56', Any), ('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({self.cf54.VerbatimString(True)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)

class D12_DC31(D12, NamedTuple('DC31', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC33(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D13_DC32)

class D13_DC33(D13, NamedTuple('DC33', [('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC32(D13, NamedTuple('DC32', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC32({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC32) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    def m1(self, p0, p1, p2, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self._f0: bool = False
        self._f2: int = int(0)
        self._f3: int = int(0)
        self._f4: int = int(0)
        self._f5: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5

    @property
    def f0(self):
        return self._f0
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
    def f5(self):
        return self._f5

class C0:
    def  __init__(self):
        self.f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f10):
        (self).f10 = f10


class C1(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f6: int = int(0)
        self._f13: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f13, f6, f7):
        (self)._f13 = f13
        (self)._f6 = f6
        (self).f7 = f7

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_401_v0_: _dafny.Map
        d_401_v0_ = _dafny.Map({p2: p1})
        d_402_v1_: _dafny.Seq
        d_402_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm"))
        d_403_v2_: _dafny.Seq
        d_403_v2_ = _dafny.SeqWithoutIsStrInference([default__.fm0(not(p0), ((d_401_v0_).set(p1, self.f7)).set(len(d_402_v1_), (self).f6), (self).f6, globalState), p0, p0, True, p0])
        d_404_v3_: _dafny.Array
        nw70_ = _dafny.Array(None, 9)
        nw70_[int(0)] = p0
        nw70_[int(1)] = p0
        nw70_[int(2)] = p0
        nw70_[int(3)] = True
        nw70_[int(4)] = p0
        nw70_[int(5)] = p0
        nw70_[int(6)] = p0
        nw70_[int(7)] = p0
        nw70_[int(8)] = p0
        d_404_v3_ = nw70_
        d_405_v4_: _dafny.Map
        d_405_v4_ = _dafny.Map({d_404_v3_: d_403_v2_})
        d_406_v5_: _dafny.Set
        d_406_v5_ = _dafny.Set({d_403_v2_, ((d_405_v4_)[d_404_v3_] if (d_404_v3_) in (d_405_v4_) else d_403_v2_), d_403_v2_, d_403_v2_, d_403_v2_})
        rhs59_ = d_406_v5_
        rhs60_ = p0
        rhs61_ = ((self).f6) + (p2)
        lhs35_ = self
        d_406_v5_ = rhs59_
        r1 = rhs60_
        lhs35_.f7 = rhs61_
        d_407_v6_: C0
        nw71_ = C0()
        nw71_.ctor__(p0)
        d_407_v6_ = nw71_
        d_408_v7_: _dafny.Map
        d_408_v7_ = _dafny.Map({d_402_v1_: d_407_v6_})
        d_408_v7_ = (d_408_v7_).set(d_402_v1_, d_407_v6_)
        d_409_v8_: _dafny.Array
        nw72_ = _dafny.Array(int(0), 11)
        d_409_v8_ = nw72_
        d_410_v9_: D2
        d_410_v9_ = D2_DC5(d_409_v8_)
        source6_ = (d_410_v9_ if d_407_v6_.f10 else d_410_v9_)
        if source6_.is_DC6:
            d_411___mcc_h0_ = source6_.cf10
            d_412___mcc_h1_ = source6_.cf11
            d_413___mcc_h2_ = source6_.cf12
            d_414_cf12_ = d_413___mcc_h2_
            d_415_cf11_ = d_412___mcc_h1_
            d_416_cf10_ = d_411___mcc_h0_
            d_415_cf11_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yo")))
            (globalState).f1 = self.f7
            d_417_v10_: D1
            d_417_v10_ = D1_DC4(d_407_v6_.f10)
            (d_407_v6_).f10 = (d_417_v10_).cf8
            d_418_v11_: _dafny.Seq
            d_418_v11_ = _dafny.SeqWithoutIsStrInference([d_407_v6_])
            d_419_v12_: _dafny.Map
            d_419_v12_ = _dafny.Map({(0) - (p2): d_407_v6_})
            d_420_v13_: _dafny.Array
            nw73_ = _dafny.Array(None, 11)
            nw73_[int(0)] = d_407_v6_
            nw73_[int(1)] = d_407_v6_
            nw73_[int(2)] = d_407_v6_
            nw73_[int(3)] = d_407_v6_
            nw73_[int(4)] = d_407_v6_
            nw73_[int(5)] = d_407_v6_
            nw73_[int(6)] = d_407_v6_
            nw73_[int(7)] = d_407_v6_
            nw73_[int(8)] = (d_418_v11_)[default__.safeIndex(582, len(d_418_v11_))]
            nw73_[int(9)] = ((d_419_v12_)[self.f7] if (self.f7) in (d_419_v12_) else d_407_v6_)
            nw73_[int(10)] = d_407_v6_
            d_420_v13_ = nw73_
            rhs62_ = d_407_v6_.f10
            rhs63_ = d_407_v6_.f10
            rhs64_ = d_407_v6_.f10
            rhs65_ = d_420_v13_
            rhs66_ = p0
            lhs36_ = d_407_v6_
            r1 = rhs62_
            lhs36_.f10 = rhs63_
            r0 = rhs64_
            d_420_v13_ = rhs65_
            r0 = rhs66_
        elif source6_.is_DC7:
            d_421___mcc_h3_ = source6_.cf13
            d_422___mcc_h4_ = source6_.cf14
            d_423___mcc_h5_ = source6_.cf15
            d_424___mcc_h6_ = source6_.cf16
            d_425_cf16_ = d_424___mcc_h6_
            d_426_cf15_ = d_423___mcc_h5_
            d_427_cf14_ = d_422___mcc_h4_
            d_428_cf13_ = d_421___mcc_h3_
            index64_ = default__.safeIndex(845, (d_409_v8_).length(0))
            (d_409_v8_)[index64_] = 341
            index65_ = default__.safeIndex(845, (d_409_v8_).length(0))
            (d_409_v8_)[index65_] = -810
            r2 = d_427_cf14_
            d_429_v14_: str
            d_429_v14_ = _dafny.CodePoint('t')
            d_430_v15_: _dafny.MultiSet
            d_430_v15_ = _dafny.MultiSet([(d_407_v6_.f10) in (d_403_v2_), d_407_v6_.f10])
            d_431_v16_: _dafny.Seq
            d_431_v16_ = _dafny.SeqWithoutIsStrInference([(len(d_403_v2_)) * (self.f7)])
            d_432_v17_: _dafny.Array
            nw74_ = _dafny.Array(int(0), 10)
            d_432_v17_ = nw74_
            d_433_v18_: _dafny.Map
            d_433_v18_ = _dafny.Map({d_407_v6_.f10: d_432_v17_})
            d_434_v19_: _dafny.Set
            d_434_v19_ = _dafny.Set({p0})
            d_435_v20_: _dafny.Map
            d_435_v20_ = _dafny.Map({D0_DC0(d_434_v19_): d_407_v6_.f10})
            d_436_v21_: D1
            d_436_v21_ = D1_DC2(d_435_v20_)
            d_437_v22_: _dafny.MultiSet
            d_437_v22_ = _dafny.MultiSet([(self).f6, 135, p2, (self).f6, (self).f6])
            d_438_v23_: D2
            d_438_v23_ = D2_DC6(d_436_v21_, d_402_v1_, ((d_437_v22_)[595] if (595) in (d_437_v22_) else self.f7))
            d_439_v24_: _dafny.Map
            d_439_v24_ = _dafny.Map({((d_433_v18_)[d_407_v6_.f10] if (d_407_v6_.f10) in (d_433_v18_) else d_432_v17_): d_438_v23_})
            index66_ = default__.safeIndex(845, (d_409_v8_).length(0))
            rhs67_ = default__.fm25(globalState)
            rhs68_ = (0) - ((0) - (len(d_431_v16_)))
            rhs69_ = _dafny.MultiSet([d_428_cf13_, (d_439_v24_) != (d_439_v24_)])
            rhs70_ = self.f7
            lhs37_ = globalState
            lhs38_ = d_409_v8_
            lhs39_ = default__.safeIndex(845, (d_409_v8_).length(0))
            d_429_v14_ = rhs67_
            lhs37_.f1 = rhs68_
            d_430_v15_ = rhs69_
            lhs38_[lhs39_] = rhs70_
            d_440_v25_: _dafny.Set
            d_440_v25_ = _dafny.Set({(self).f6, p1, len(_dafny.Set({(self).f6, (self).f6}))})
            d_441_v26_: _dafny.MultiSet
            d_441_v26_ = _dafny.MultiSet([d_440_v25_, d_440_v25_, d_440_v25_, d_440_v25_, d_440_v25_])
            r2 = ((d_441_v26_) - ((d_441_v26_).intersection(d_441_v26_))).cardinality
        elif source6_.is_DC5:
            d_442___mcc_h7_ = source6_.cf9
            d_443_cf9_ = d_442___mcc_h7_
            d_444_v27_: str
            d_444_v27_ = _dafny.CodePoint('d')
            d_444_v27_ = d_444_v27_
            r1 = d_407_v6_.f10
            index67_ = default__.safeIndex(963, (d_404_v3_).length(0))
            (d_404_v3_)[index67_] = True
            index68_ = default__.safeIndex(963, (d_404_v3_).length(0))
            (d_404_v3_)[index68_] = not(d_407_v6_.f10)
            r0 = (d_404_v3_)[default__.safeIndex(963, (d_404_v3_).length(0))]
        elif True:
            d_445___mcc_h8_ = source6_.cf17
            d_446_cf17_ = d_445___mcc_h8_
            index69_ = default__.safeIndex(364, (d_404_v3_).length(0))
            (d_404_v3_)[index69_] = p0
            d_447_v28_: _dafny.Seq
            d_447_v28_ = _dafny.SeqWithoutIsStrInference([d_402_v1_, d_402_v1_, d_402_v1_])
            d_448_v29_: D0
            d_448_v29_ = D0_DC0(_dafny.Set({p0, p0}))
            d_449_v30_: _dafny.Map
            d_449_v30_ = _dafny.Map({d_448_v29_: d_407_v6_.f10})
            d_450_v31_: D1
            d_450_v31_ = D1_DC2(d_449_v30_)
            d_451_v32_: D2
            d_451_v32_ = D2_DC6(d_450_v31_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukxgv")), (self).f6)
            index70_ = default__.safeIndex(364, (d_404_v3_).length(0))
            rhs71_ = (d_447_v28_)[default__.safeIndex(len((d_402_v1_) + (d_402_v1_)), len(d_447_v28_))]
            rhs72_ = (((d_451_v32_).cf11) + (d_402_v1_)) + ((d_402_v1_) + (d_402_v1_))
            rhs73_ = default__.fm0(p0, d_401_v0_, self.f7, globalState)
            lhs40_ = d_404_v3_
            lhs41_ = default__.safeIndex(364, (d_404_v3_).length(0))
            d_402_v1_ = rhs71_
            d_402_v1_ = rhs72_
            lhs40_[lhs41_] = rhs73_
            d_452_v33_: _dafny.Map
            d_452_v33_ = _dafny.Map({d_407_v6_.f10: p2})
            d_453_v34_: _dafny.MultiSet
            d_453_v34_ = _dafny.MultiSet([d_409_v8_])
            index71_ = default__.safeIndex(364, (d_404_v3_).length(0))
            rhs74_ = ((0) - (((d_452_v33_)[d_407_v6_.f10] if (d_407_v6_.f10) in (d_452_v33_) else (0) - (p2)))) - (((d_453_v34_).set(d_409_v8_, default__.abs((0) - (len(d_402_v1_))))).cardinality)
            rhs75_ = ((p2) < ((self).f6)) or (d_407_v6_.f10)
            lhs42_ = self
            lhs43_ = d_404_v3_
            lhs44_ = default__.safeIndex(364, (d_404_v3_).length(0))
            lhs42_.f7 = rhs74_
            lhs43_[lhs44_] = rhs75_
            d_454_v35_: _dafny.Set
            d_454_v35_ = _dafny.Set({p2})
            d_455_v37_: _dafny.Seq
            def iife33_():
                coll15_ = _dafny.Set()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(118, -912):
                    d_456_v36_: int = compr_15_
                    if ((118) <= (d_456_v36_)) and ((d_456_v36_) < (-912)):
                        coll15_ = coll15_.union(_dafny.Set([(d_456_v36_) + (p1)]))
                return _dafny.Set(coll15_)
            d_455_v37_ = _dafny.SeqWithoutIsStrInference([(d_454_v35_).intersection(d_454_v35_), d_454_v35_, iife33_()
            ])
            d_455_v37_ = default__.fm26((0) - ((d_451_v32_).cf12), p0, globalState)
            d_457_v39_: C0
            nw75_ = C0()
            def iife34_():
                coll16_ = _dafny.Set()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(-209, 576):
                    d_458_v38_: int = compr_16_
                    if ((-209) <= (d_458_v38_)) and ((d_458_v38_) < (576)):
                        coll16_ = coll16_.union(_dafny.Set([(d_458_v38_) + (len(d_403_v2_))]))
                return _dafny.Set(coll16_)
            nw75_.ctor__((iife34_()
            ).ispropersubset(d_454_v35_))
            d_457_v39_ = nw75_
        (self).f7 = len(d_402_v1_)
        if False:
            index72_ = default__.safeIndex(752, (d_409_v8_).length(0))
            (d_409_v8_)[index72_] = p2
            index73_ = default__.safeIndex(752, (d_409_v8_).length(0))
            (d_409_v8_)[index73_] = p2
            (globalState).f1 = self.f7
            d_459_v40_: _dafny.Array
            def lambda17_(d_460_v6_, d_461_v2_):
                def lambda18_(d_462_i0_):
                    return (_dafny.SeqWithoutIsStrInference([d_460_v6_.f10, d_460_v6_.f10])) + (d_461_v2_)

                return lambda18_

            init9_ = lambda17_(d_407_v6_, d_403_v2_)
            nw76_ = _dafny.Array(None, 28)
            for i0_9_ in range(nw76_.length(0)):
                nw76_[i0_9_] = init9_(i0_9_)
            d_459_v40_ = nw76_
            d_463_v41_: D0
            d_463_v41_ = D0_DC1(d_407_v6_.f10, d_404_v3_, d_407_v6_.f10)
            d_464_v42_: _dafny.Map
            d_464_v42_ = _dafny.Map({d_407_v6_.f10: not(d_407_v6_.f10)})
            d_465_v43_: D5
            d_465_v43_ = D5_DC11(d_403_v2_)
            d_466_v44_: _dafny.Seq
            d_466_v44_ = _dafny.SeqWithoutIsStrInference([d_403_v2_, d_403_v2_])
            d_467_v45_: _dafny.Set
            d_467_v45_ = _dafny.Set({self.f7})
            d_468_v46_: _dafny.Set
            d_468_v46_ = _dafny.Set({len(d_467_v45_)})
            pat_let_tv6_ = d_407_v6_
            nw77_ = _dafny.Array(None, 28)
            nw77_[int(0)] = d_403_v2_
            nw77_[int(1)] = ((d_403_v2_).set(default__.safeIndex(len((d_402_v1_).set(default__.safeIndex(self.f7, len(d_402_v1_)), (self).f13)), len(d_403_v2_)), d_407_v6_.f10)) + (d_403_v2_)
            nw77_[int(2)] = d_403_v2_
            nw77_[int(3)] = d_403_v2_
            nw77_[int(4)] = _dafny.SeqWithoutIsStrInference([p0, (d_463_v41_).cf3])
            nw77_[int(5)] = _dafny.SeqWithoutIsStrInference([((d_464_v42_)[d_407_v6_.f10] if (d_407_v6_.f10) in (d_464_v42_) else d_407_v6_.f10)])
            nw77_[int(6)] = d_403_v2_
            nw77_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_407_v6_.f10, d_407_v6_.f10])) + (((D5_DC11(d_403_v2_)).cf20).set(default__.safeIndex((self).f6, len((D5_DC11(d_403_v2_)).cf20)), d_407_v6_.f10))
            nw77_[int(8)] = _dafny.SeqWithoutIsStrInference([d_407_v6_.f10, d_407_v6_.f10])
            nw77_[int(9)] = (d_403_v2_) + ((D5_DC11(_dafny.SeqWithoutIsStrInference([False]))).cf20)
            nw77_[int(10)] = (d_403_v2_) + (d_403_v2_)
            nw77_[int(11)] = (d_465_v43_).cf20
            nw77_[int(12)] = d_403_v2_
            nw77_[int(13)] = (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, p0, p0]))
            nw77_[int(14)] = (d_466_v44_)[default__.safeIndex(p2, len(d_466_v44_))]
            nw77_[int(15)] = d_403_v2_
            nw77_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_403_v2_)[default__.safeIndex(self.f7, len(d_403_v2_))], d_407_v6_.f10])
            nw77_[int(17)] = d_403_v2_
            nw77_[int(18)] = (d_403_v2_) + (d_403_v2_)
            nw77_[int(19)] = d_403_v2_
            nw77_[int(20)] = d_403_v2_
            nw77_[int(21)] = (d_403_v2_) + (d_403_v2_)
            nw77_[int(22)] = d_403_v2_
            nw77_[int(23)] = d_403_v2_
            nw77_[int(24)] = _dafny.SeqWithoutIsStrInference([default__.fm0(d_407_v6_.f10, d_401_v0_, p1, globalState)])
            nw77_[int(25)] = d_403_v2_
            def iife35_(_pat_let9_0):
                def iife36_(d_469_dt__update__tmp_h0_):
                    def iife37_(_pat_let10_0):
                        def iife38_(d_470_dt__update_hcf1_h0_):
                            return D0_DC1(d_470_dt__update_hcf1_h0_, (d_469_dt__update__tmp_h0_).cf2, (d_469_dt__update__tmp_h0_).cf3)
                        return iife38_(_pat_let10_0)
                    return iife37_(pat_let_tv6_.f10)
                return iife36_(_pat_let9_0)
            nw77_[int(26)] = _dafny.SeqWithoutIsStrInference([(iife35_(d_463_v41_)).cf1])
            nw77_[int(27)] = default__.fm27(len(d_467_v45_), d_402_v1_, (self).f13, globalState)
            d_459_v40_ = nw77_
            r0 = not((d_407_v6_.f10 if d_407_v6_.f10 else (len(d_468_v46_)) < ((d_409_v8_)[default__.safeIndex(752, (d_409_v8_).length(0))])))
            index74_ = default__.safeIndex(752, (d_409_v8_).length(0))
            (d_409_v8_)[index74_] = self.f7
        elif True:
            d_471_v47_: _dafny.Map
            d_471_v47_ = _dafny.Map({d_409_v8_: (self).f13})
            d_471_v47_ = d_471_v47_
            d_472_v48_: _dafny.Map
            d_472_v48_ = _dafny.Map({p2: d_407_v6_.f10})
            d_473_v49_: _dafny.Seq
            d_473_v49_ = _dafny.SeqWithoutIsStrInference([d_472_v48_])
            d_474_v50_: D1
            d_474_v50_ = D1_DC3(False, len((d_473_v49_)[default__.safeIndex(len(d_402_v1_), len(d_473_v49_))]), p1)
            r0 = (d_407_v6_.f10 if (d_474_v50_).cf5 else True)
            index75_ = default__.safeIndex(342, (d_404_v3_).length(0))
            (d_404_v3_)[index75_] = (d_407_v6_.f10) or (d_407_v6_.f10)
            index76_ = default__.safeIndex(100, (d_409_v8_).length(0))
            (d_409_v8_)[index76_] = p1
            d_475_v51_: _dafny.Set
            d_475_v51_ = _dafny.Set({d_407_v6_.f10, d_407_v6_.f10})
            d_476_v52_: _dafny.Map
            d_476_v52_ = _dafny.Map({len(d_475_v51_): d_475_v51_})
            index77_ = default__.safeIndex(342, (d_404_v3_).length(0))
            index78_ = default__.safeIndex(100, (d_409_v8_).length(0))
            rhs76_ = d_407_v6_.f10
            rhs77_ = d_407_v6_.f10
            rhs78_ = (0) - (len((_dafny.Map({p2: d_475_v51_})) | ((d_476_v52_) | (d_476_v52_))))
            rhs79_ = p1
            lhs45_ = d_407_v6_
            lhs46_ = d_404_v3_
            lhs47_ = default__.safeIndex(342, (d_404_v3_).length(0))
            lhs48_ = globalState
            lhs49_ = d_409_v8_
            lhs50_ = default__.safeIndex(100, (d_409_v8_).length(0))
            lhs45_.f10 = rhs76_
            lhs46_[lhs47_] = rhs77_
            lhs48_.f1 = rhs78_
            lhs49_[lhs50_] = rhs79_
            source7_ = d_474_v50_
            if source7_.is_DC3:
                d_477___mcc_h9_ = source7_.cf5
                d_478___mcc_h10_ = source7_.cf6
                d_479___mcc_h11_ = source7_.cf7
                d_480_cf7_ = d_479___mcc_h11_
                d_481_cf6_ = d_478___mcc_h10_
                d_482_cf5_ = d_477___mcc_h9_
                d_483_v53_: str
                d_483_v53_ = _dafny.CodePoint('u')
                d_484_v54_: _dafny.Set
                d_484_v54_ = _dafny.Set({self.f7})
                d_483_v53_ = ((self).f13 if (p1) > (len(d_484_v54_)) else d_483_v53_)
                d_485_v55_: D5
                d_485_v55_ = D5_DC12(d_481_cf6_)
                d_486_v56_: _dafny.MultiSet
                d_486_v56_ = _dafny.MultiSet([(d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]])
                d_487_v57_: _dafny.Map
                d_487_v57_ = _dafny.Map({p2: d_486_v56_})
                d_485_v55_ = default__.fm28((default__.fm27(len(d_487_v57_), d_402_v1_, (self).f13, globalState) if False else d_403_v2_), p2, globalState)
                d_474_v50_ = d_474_v50_
                d_488_v59_: _dafny.Map
                d_488_v59_ = _dafny.Map({d_481_cf6_: d_483_v53_})
                d_489_v60_: _dafny.Set
                d_489_v60_ = _dafny.Set({d_488_v59_, _dafny.Map({532: (self).f13})})
                def iife39_():
                    coll17_ = _dafny.Map()
                    compr_17_: _dafny.Map
                    for compr_17_ in (d_489_v60_).Elements:
                        d_490_v58_: _dafny.Map = compr_17_
                        if (d_490_v58_) in (d_489_v60_):
                            coll17_[d_490_v58_] = d_481_cf6_
                    return _dafny.Map(coll17_)
                d_481_cf6_ = ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]) - ((0) - (len(iife39_()
                )))
            elif source7_.is_DC4:
                d_491___mcc_h12_ = source7_.cf8
                d_492_cf8_ = d_491___mcc_h12_
                d_493_v61_: _dafny.Seq
                d_493_v61_ = _dafny.SeqWithoutIsStrInference([(d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))], (self).f6])
                d_494_v62_: _dafny.Array
                nw78_ = _dafny.Array(None, 5)
                nw78_[int(0)] = d_493_v61_
                nw78_[int(1)] = d_493_v61_
                nw78_[int(2)] = d_493_v61_
                nw78_[int(3)] = d_493_v61_
                nw78_[int(4)] = d_493_v61_
                d_494_v62_ = nw78_
                index79_ = default__.safeIndex(328, (d_494_v62_).length(0))
                (d_494_v62_)[index79_] = d_493_v61_
                index80_ = default__.safeIndex(342, (d_404_v3_).length(0))
                index81_ = default__.safeIndex(328, (d_494_v62_).length(0))
                rhs80_ = ((_dafny.SeqWithoutIsStrInference([(self).f13 for d_495_i1_ in range(default__.abs(403))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "endmc")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmwqy")))
                rhs81_ = d_407_v6_.f10
                rhs82_ = d_493_v61_
                lhs51_ = d_404_v3_
                lhs52_ = default__.safeIndex(342, (d_404_v3_).length(0))
                lhs53_ = d_494_v62_
                lhs54_ = default__.safeIndex(328, (d_494_v62_).length(0))
                d_402_v1_ = rhs80_
                lhs51_[lhs52_] = rhs81_
                lhs53_[lhs54_] = rhs82_
                d_496_v63_: _dafny.Array
                nw79_ = _dafny.Array(int(0), 6)
                d_496_v63_ = nw79_
                d_497_v64_: _dafny.Seq
                d_497_v64_ = _dafny.SeqWithoutIsStrInference([d_496_v63_])
                d_498_v65_: _dafny.Map
                d_498_v65_ = _dafny.Map({d_402_v1_: False})
                d_496_v63_ = (d_497_v64_)[default__.safeIndex((len(d_498_v65_)) * (p1), len(d_497_v64_))]
                d_499_v66_: _dafny.Map
                d_499_v66_ = _dafny.Map({p0: d_492_cf8_})
                d_500_v67_: _dafny.Seq
                d_500_v67_ = _dafny.SeqWithoutIsStrInference([d_499_v66_])
                d_499_v66_ = ((d_500_v67_)[default__.safeIndex(p2, len(d_500_v67_))]).set(not(d_492_cf8_), (d_404_v3_)[default__.safeIndex(342, (d_404_v3_).length(0))])
                d_501_v68_: _dafny.MultiSet
                d_501_v68_ = _dafny.MultiSet([168, (self).f6, p1])
                d_502_v69_: _dafny.Map
                d_502_v69_ = _dafny.Map({d_492_cf8_: d_403_v2_})
                d_503_v70_: _dafny.MultiSet
                d_503_v70_ = _dafny.MultiSet([False, d_407_v6_.f10])
                index82_ = default__.safeIndex(342, (d_404_v3_).length(0))
                rhs83_ = p0
                rhs84_ = (((d_401_v0_)[len(d_502_v69_)] if (len(d_502_v69_)) in (d_401_v0_) else default__.fm1(d_407_v6_.f10, globalState))) <= ((0) - (self.f7))
                rhs85_ = _dafny.MultiSet([((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_407_v6_.f10, d_407_v6_.f10]))) | (d_503_v70_)).cardinality])
                rhs86_ = not(d_407_v6_.f10)
                lhs55_ = d_404_v3_
                lhs56_ = default__.safeIndex(342, (d_404_v3_).length(0))
                lhs57_ = d_407_v6_
                lhs55_[lhs56_] = rhs83_
                lhs57_.f10 = rhs84_
                d_501_v68_ = rhs85_
                r0 = rhs86_
            elif True:
                d_504___mcc_h13_ = source7_.cf4
                d_505_cf4_ = d_504___mcc_h13_
                d_506_v71_: C0
                nw80_ = C0()
                nw80_.ctor__(not(d_407_v6_.f10))
                d_506_v71_ = nw80_
                d_507_v72_: _dafny.Array
                def lambda19_(d_508_v71_, d_509_v6_):
                    def lambda20_(d_510_i2_):
                        return default__.safeModuloInt(d_510_i2_, len(_dafny.Map({d_508_v71_.f10: _dafny.SeqWithoutIsStrInference([d_509_v6_.f10])})))

                    return lambda20_

                init10_ = lambda19_(d_506_v71_, d_407_v6_)
                nw81_ = _dafny.Array(None, 11)
                for i0_10_ in range(nw81_.length(0)):
                    nw81_[i0_10_] = init10_(i0_10_)
                d_507_v72_ = nw81_
                d_507_v72_ = d_507_v72_
                d_511_v73_: _dafny.Array
                def lambda21_(d_512_v0_):
                    def lambda22_(d_513_i3_):
                        return d_512_v0_

                    return lambda22_

                init11_ = lambda21_(d_401_v0_)
                nw82_ = _dafny.Array(None, 10)
                for i0_11_ in range(nw82_.length(0)):
                    nw82_[i0_11_] = init11_(i0_11_)
                d_511_v73_ = nw82_
                index83_ = default__.safeIndex(939, (d_511_v73_).length(0))
                (d_511_v73_)[index83_] = _dafny.Map({p1: default__.fm1(default__.fm0(d_407_v6_.f10, d_401_v0_, p2, globalState), globalState)})
                d_514_v74_: _dafny.Array
                def lambda23_(d_515_p1_):
                    def lambda24_(d_516_i4_):
                        return _dafny.SeqWithoutIsStrInference([959, d_515_p1_])

                    return lambda24_

                init12_ = lambda23_(p1)
                nw83_ = _dafny.Array(None, 22)
                for i0_12_ in range(nw83_.length(0)):
                    nw83_[i0_12_] = init12_(i0_12_)
                d_514_v74_ = nw83_
                index84_ = default__.safeIndex(939, (d_511_v73_).length(0))
                rhs87_ = self.f7
                rhs88_ = default__.safeModuloInt((0) - (default__.safeDivisionInt((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))], -575)), ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]) + ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]))
                rhs89_ = (d_401_v0_) | (d_401_v0_)
                rhs90_ = default__.fm0(not (True) or (p0), _dafny.Map({-267: self.f7}), (d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))], globalState)
                rhs91_ = d_514_v74_
                lhs58_ = globalState
                lhs59_ = d_511_v73_
                lhs60_ = default__.safeIndex(939, (d_511_v73_).length(0))
                lhs61_ = d_407_v6_
                lhs58_.f1 = rhs87_
                r2 = rhs88_
                lhs59_[lhs60_] = rhs89_
                lhs61_.f10 = rhs90_
                d_514_v74_ = rhs91_
                (d_506_v71_).f10 = False
            if p0:
                index85_ = default__.safeIndex(342, (d_404_v3_).length(0))
                (d_404_v3_)[index85_] = ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]) == ((p2) + ((0) - ((self).f6)))
                (globalState).f1 = 846
                d_517_v75_: _dafny.Map
                d_517_v75_ = _dafny.Map({D0_DC0(d_475_v51_): not(d_407_v6_.f10)})
                d_518_v76_: D1
                d_518_v76_ = D1_DC2(d_517_v75_)
                d_519_v78_: D0
                d_519_v78_ = D0_DC0(_dafny.Set({(d_404_v3_)[default__.safeIndex(342, (d_404_v3_).length(0))]}))
                pat_let_tv7_ = d_475_v51_
                pat_let_tv8_ = d_475_v51_
                pat_let_tv9_ = d_517_v75_
                d_520_v79_: _dafny.Array
                nw84_ = _dafny.Array(None, 23)
                nw84_[int(0)] = d_518_v76_
                nw84_[int(1)] = d_518_v76_
                nw84_[int(2)] = d_518_v76_
                nw84_[int(3)] = d_518_v76_
                def iife40_():
                    def iife45_(_pat_let13_0):
                        def iife46_(d_524_dt__update__tmp_h1_):
                            def iife47_(_pat_let14_0):
                                def iife48_(d_525_dt__update_hcf0_h0_):
                                    return D0_DC0(d_525_dt__update_hcf0_h0_)
                                return iife48_(_pat_let14_0)
                            return iife47_(pat_let_tv8_)
                        return iife46_(_pat_let13_0)
                    coll18_ = _dafny.Map()
                    def iife41_(_pat_let11_0):
                        def iife42_(d_521_dt__update__tmp_h1_):
                            def iife43_(_pat_let12_0):
                                def iife44_(d_522_dt__update_hcf0_h0_):
                                    return D0_DC0(d_522_dt__update_hcf0_h0_)
                                return iife44_(_pat_let12_0)
                            return iife43_(pat_let_tv7_)
                        return iife42_(_pat_let11_0)
                    compr_18_: D0
                    for compr_18_ in (_dafny.Set({d_519_v78_, iife41_(D0_DC0(d_475_v51_))})).Elements:
                        d_523_v77_: D0 = compr_18_
                        if (d_523_v77_) in (_dafny.Set({d_519_v78_, iife45_(D0_DC0(d_475_v51_))})):
                            coll18_[d_523_v77_] = True
                    return _dafny.Map(coll18_)
                nw84_[int(4)] = D1_DC2(iife40_()
)
                nw84_[int(5)] = d_518_v76_
                nw84_[int(6)] = d_518_v76_
                nw84_[int(7)] = d_518_v76_
                nw84_[int(8)] = d_518_v76_
                nw84_[int(9)] = d_518_v76_
                def iife49_(_pat_let15_0):
                    def iife50_(d_526_dt__update__tmp_h2_):
                        def iife51_(_pat_let16_0):
                            def iife52_(d_527_dt__update_hcf4_h0_):
                                return D1_DC2(d_527_dt__update_hcf4_h0_)
                            return iife52_(_pat_let16_0)
                        return iife51_(pat_let_tv9_)
                    return iife50_(_pat_let15_0)
                nw84_[int(10)] = iife49_(D1_DC2(d_517_v75_))
                nw84_[int(11)] = d_518_v76_
                nw84_[int(12)] = D1_DC2(d_517_v75_)
                nw84_[int(13)] = d_518_v76_
                nw84_[int(14)] = d_518_v76_
                nw84_[int(15)] = D1_DC2(d_517_v75_)
                nw84_[int(16)] = D1_DC2(d_517_v75_)
                nw84_[int(17)] = d_518_v76_
                nw84_[int(18)] = d_518_v76_
                nw84_[int(19)] = D1_DC2(d_517_v75_)
                nw84_[int(20)] = d_518_v76_
                nw84_[int(21)] = d_518_v76_
                nw84_[int(22)] = D1_DC2(_dafny.Map({d_519_v78_: False}))
                d_520_v79_ = nw84_
                d_528_v80_: D6
                d_528_v80_ = D6_DC13(d_520_v79_)
                d_529_v81_: _dafny.Array
                nw85_ = _dafny.Array(None, 23)
                nw85_[int(0)] = d_520_v79_
                nw85_[int(1)] = d_520_v79_
                nw85_[int(2)] = d_520_v79_
                nw85_[int(3)] = d_520_v79_
                nw85_[int(4)] = d_520_v79_
                nw85_[int(5)] = (d_528_v80_).cf22
                nw85_[int(6)] = d_520_v79_
                nw85_[int(7)] = d_520_v79_
                nw85_[int(8)] = d_520_v79_
                nw85_[int(9)] = d_520_v79_
                nw85_[int(10)] = d_520_v79_
                nw85_[int(11)] = d_520_v79_
                nw85_[int(12)] = d_520_v79_
                nw85_[int(13)] = d_520_v79_
                nw85_[int(14)] = d_520_v79_
                nw85_[int(15)] = d_520_v79_
                nw85_[int(16)] = d_520_v79_
                nw85_[int(17)] = d_520_v79_
                nw85_[int(18)] = d_520_v79_
                nw85_[int(19)] = d_520_v79_
                nw85_[int(20)] = d_520_v79_
                nw85_[int(21)] = d_520_v79_
                nw85_[int(22)] = d_520_v79_
                d_529_v81_ = nw85_
                index86_ = default__.safeIndex(940, (d_529_v81_).length(0))
                (d_529_v81_)[index86_] = d_520_v79_
                d_530_v82_: _dafny.MultiSet
                d_530_v82_ = _dafny.MultiSet([True, d_407_v6_.f10, d_407_v6_.f10, not(p0), (d_404_v3_)[default__.safeIndex(342, (d_404_v3_).length(0))]])
                d_531_v83_: _dafny.Map
                d_531_v83_ = _dafny.Map({(d_407_v6_.f10 if default__.fm0(p0, d_401_v0_, self.f7, globalState) else (d_404_v3_)[default__.safeIndex(342, (d_404_v3_).length(0))]): d_520_v79_})
                index87_ = default__.safeIndex(940, (d_529_v81_).length(0))
                rhs92_ = default__.safeModuloInt((self.f7) + ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]), (d_530_v82_).cardinality)
                rhs93_ = ((d_531_v83_)[d_407_v6_.f10] if (d_407_v6_.f10) in (d_531_v83_) else d_520_v79_)
                lhs62_ = globalState
                lhs63_ = d_529_v81_
                lhs64_ = default__.safeIndex(940, (d_529_v81_).length(0))
                lhs62_.f1 = rhs92_
                lhs63_[lhs64_] = rhs93_
                d_532_v84_: C0
                nw86_ = C0()
                nw86_.ctor__(p0)
                d_532_v84_ = nw86_
                d_533_v85_: _dafny.Array
                nw87_ = _dafny.Array(D0.default()(), 9)
                d_533_v85_ = nw87_
                d_534_v86_: _dafny.Seq
                d_534_v86_ = _dafny.SeqWithoutIsStrInference([d_533_v85_, (d_533_v85_ if p0 else d_533_v85_)])
                d_533_v85_ = (d_534_v86_)[default__.safeIndex((self.f7) - (self.f7), len(d_534_v86_))]
            elif True:
                index88_ = default__.safeIndex(100, (d_409_v8_).length(0))
                (d_409_v8_)[index88_] = (0) - ((d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))])
                d_402_v1_ = ((d_402_v1_) + (d_402_v1_)) + ((d_402_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvmvk"))))
                (d_407_v6_).f10 = d_407_v6_.f10
                d_535_v88_: _dafny.Set
                d_535_v88_ = _dafny.Set({d_402_v1_})
                def iife53_():
                    coll19_ = _dafny.Map()
                    compr_19_: _dafny.Seq
                    for compr_19_ in (d_535_v88_).Elements:
                        d_536_v87_: _dafny.Seq = compr_19_
                        if (d_536_v87_) in (d_535_v88_):
                            coll19_[d_536_v87_] = (d_409_v8_)[default__.safeIndex(100, (d_409_v8_).length(0))]
                    return _dafny.Map(coll19_)
                d_475_v51_ = (default__.fm3(p0, p1, d_475_v51_, iife53_()
                , globalState)).cf0
                (d_407_v6_).f10 = not(not ((d_404_v3_)[default__.safeIndex(342, (d_404_v3_).length(0))]) or (d_407_v6_.f10))
        r1 = default__.fm0(p0, (d_401_v0_) | (d_401_v0_), (self).f6, globalState)
        r0 = True
        r1 = (_dafny.Map({d_404_v3_: self.f7})) != (_dafny.Map({d_404_v3_: p2}))
        r2 = default__.fm1(d_407_v6_.f10, globalState)
        return r0, r1, r2

    def m10(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: D1 = D1.default()()
        r3: bool = False
        d_537_v0_: bool
        d_537_v0_ = False
        d_538_v1_: _dafny.Map
        d_538_v1_ = _dafny.Map({d_537_v0_: 654})
        hi4_ = len((d_538_v1_) | (d_538_v1_))
        for d_539_i0_ in range((default__.fm1(d_537_v0_, globalState)) * (-878), hi4_):
            d_540_v2_: _dafny.Set
            d_540_v2_ = _dafny.Set({d_539_i0_})
            d_541_v3_: _dafny.Set
            d_541_v3_ = _dafny.Set({612, len(d_540_v2_), (self).f6})
            d_542_v4_: _dafny.Seq
            d_542_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twoo"))
            d_543_v5_: _dafny.Seq
            d_543_v5_ = _dafny.SeqWithoutIsStrInference([d_537_v0_, d_537_v0_])
            d_544_v6_: _dafny.Set
            d_544_v6_ = _dafny.Set({d_537_v0_})
            d_545_v7_: _dafny.Map
            d_545_v7_ = _dafny.Map({d_537_v0_: d_542_v4_})
            d_546_v8_: _dafny.Array
            nw88_ = _dafny.Array(None, 17)
            nw88_[int(0)] = len(d_541_v3_)
            nw88_[int(1)] = p0
            nw88_[int(2)] = (self).f6
            nw88_[int(3)] = p0
            nw88_[int(4)] = (self).f6
            nw88_[int(5)] = d_539_i0_
            nw88_[int(6)] = len(d_542_v4_)
            nw88_[int(7)] = (0) - ((_dafny.MultiSet(d_543_v5_)).cardinality)
            nw88_[int(8)] = len(d_544_v6_)
            nw88_[int(9)] = d_539_i0_
            nw88_[int(10)] = d_539_i0_
            nw88_[int(11)] = p0
            nw88_[int(12)] = 502
            nw88_[int(13)] = d_539_i0_
            nw88_[int(14)] = len(d_543_v5_)
            nw88_[int(15)] = d_539_i0_
            nw88_[int(16)] = len(((d_545_v7_)[d_537_v0_] if (d_537_v0_) in (d_545_v7_) else d_542_v4_))
            d_546_v8_ = nw88_
            d_547_v9_: _dafny.Seq
            d_547_v9_ = _dafny.SeqWithoutIsStrInference([d_539_i0_, (0) - (d_539_i0_)])
            d_548_v10_: _dafny.Map
            d_548_v10_ = _dafny.Map({d_547_v9_: d_539_i0_})
            d_549_v11_: _dafny.Map
            d_549_v11_ = _dafny.Map({d_546_v8_: len(d_548_v10_)})
            d_550_v12_: _dafny.Map
            d_550_v12_ = _dafny.Map({d_539_i0_: d_539_i0_})
            d_549_v11_ = (d_549_v11_).set((d_546_v8_ if default__.fm0(not(d_537_v0_), d_550_v12_, 432, globalState) else d_546_v8_), (self).f6)
            r3 = d_537_v0_
            d_544_v6_ = d_544_v6_
            d_551_v13_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Map({}), 3)
            d_551_v13_ = nw89_
            d_552_v14_: _dafny.Map
            d_552_v14_ = _dafny.Map({(self).f13: p0})
            index89_ = default__.safeIndex(34, (d_551_v13_).length(0))
            (d_551_v13_)[index89_] = d_552_v14_
            d_553_v15_: D7
            d_553_v15_ = D7_DC16(_dafny.Map({(self).f13: p0}))
            index90_ = default__.safeIndex(34, (d_551_v13_).length(0))
            (d_551_v13_)[index90_] = (d_552_v14_ if d_537_v0_ else (d_553_v15_).cf30)
        d_554_v16_: _dafny.Array
        def lambda25_(d_555_i1_):
            return default__.safeDivisionInt(d_555_i1_, (self).f6)

        init13_ = lambda25_
        nw90_ = _dafny.Array(None, 12)
        for i0_13_ in range(nw90_.length(0)):
            nw90_[i0_13_] = init13_(i0_13_)
        d_554_v16_ = nw90_
        d_556_v17_: _dafny.Map
        d_556_v17_ = _dafny.Map({p0: self.f7})
        index91_ = default__.safeIndex(139, (d_554_v16_).length(0))
        (d_554_v16_)[index91_] = ((d_556_v17_)[self.f7] if (self.f7) in (d_556_v17_) else (self).f6)
        d_557_v18_: _dafny.MultiSet
        d_557_v18_ = _dafny.MultiSet([d_537_v0_])
        index92_ = default__.safeIndex(139, (d_554_v16_).length(0))
        rhs94_ = p0
        rhs95_ = (d_557_v18_).issubset(d_557_v18_)
        lhs65_ = d_554_v16_
        lhs66_ = default__.safeIndex(139, (d_554_v16_).length(0))
        lhs65_[lhs66_] = rhs94_
        r3 = rhs95_
        d_558_v19_: _dafny.Seq
        d_558_v19_ = _dafny.SeqWithoutIsStrInference([((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]) - ((self).f6)])
        index93_ = default__.safeIndex(139, (d_554_v16_).length(0))
        index94_ = default__.safeIndex(139, (d_554_v16_).length(0))
        index95_ = default__.safeIndex(139, (d_554_v16_).length(0))
        rhs96_ = (self).f6
        rhs97_ = (((self).f6 if d_537_v0_ else (d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))])) - ((0) - (default__.safeModuloInt(self.f7, (d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))])))
        rhs98_ = (d_558_v19_)[default__.safeIndex(default__.safeModuloInt((self).f6, p0), len(d_558_v19_))]
        lhs67_ = d_554_v16_
        lhs68_ = default__.safeIndex(139, (d_554_v16_).length(0))
        lhs69_ = d_554_v16_
        lhs70_ = default__.safeIndex(139, (d_554_v16_).length(0))
        lhs71_ = d_554_v16_
        lhs72_ = default__.safeIndex(139, (d_554_v16_).length(0))
        lhs67_[lhs68_] = rhs96_
        lhs69_[lhs70_] = rhs97_
        lhs71_[lhs72_] = rhs98_
        hi5_ = p0
        for d_559_i2_ in range(default__.fm1(d_537_v0_, globalState), hi5_):
            d_560_v20_: _dafny.Set
            d_560_v20_ = _dafny.Set({(d_559_i2_) != ((0) - ((0) - (d_559_i2_))), (d_537_v0_) == (d_537_v0_), d_537_v0_, default__.fm0(d_537_v0_, d_556_v17_, (d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))], globalState), False})
            d_560_v20_ = d_560_v20_
            d_561_v21_: _dafny.Map
            d_561_v21_ = _dafny.Map({d_558_v19_: self.f7})
            index96_ = default__.safeIndex(139, (d_554_v16_).length(0))
            rhs99_ = (default__.fm29(globalState)) | ((d_561_v21_) | (d_561_v21_))
            rhs100_ = (0) - ((p0) - ((p0) + (410)))
            rhs101_ = ((d_558_v19_) + ((_dafny.SeqWithoutIsStrInference([(0) - (-628) for d_562_i3_ in range(default__.abs(-878))])) + (d_558_v19_))).set(default__.safeIndex((D1_DC3(d_537_v0_, len(d_538_v1_), (self).f6)).cf6, len((d_558_v19_) + ((_dafny.SeqWithoutIsStrInference([(0) - (-628) for d_563_i3_ in range(default__.abs(-878))])) + (d_558_v19_)))), (self).f6)
            lhs73_ = d_554_v16_
            lhs74_ = default__.safeIndex(139, (d_554_v16_).length(0))
            d_561_v21_ = rhs99_
            lhs73_[lhs74_] = rhs100_
            d_558_v19_ = rhs101_
            d_564_v22_: _dafny.Array
            nw91_ = _dafny.Array(False, 5)
            d_564_v22_ = nw91_
            d_564_v22_ = d_564_v22_
            d_565_v23_: _dafny.Array
            def lambda26_(d_566_v17_):
                def lambda27_(d_567_i4_):
                    return d_566_v17_

                return lambda27_

            init14_ = lambda26_(d_556_v17_)
            nw92_ = _dafny.Array(None, 2)
            for i0_14_ in range(nw92_.length(0)):
                nw92_[i0_14_] = init14_(i0_14_)
            d_565_v23_ = nw92_
            d_568_v24_: _dafny.Map
            d_568_v24_ = _dafny.Map({default__.fm1(d_537_v0_, globalState): d_565_v23_})
            d_568_v24_ = (d_568_v24_).set(default__.safeModuloInt((self).f6, (0) - (d_559_i2_)), d_565_v23_)
        hi6_ = self.f7
        for d_569_i5_ in range(default__.safeModuloInt(p0, len(d_556_v17_)), hi6_):
            index97_ = default__.safeIndex(139, (d_554_v16_).length(0))
            (d_554_v16_)[index97_] = default__.safeDivisionInt(-363, d_569_i5_)
            (globalState).f1 = p0
            index98_ = default__.safeIndex(139, (d_554_v16_).length(0))
            (d_554_v16_)[index98_] = self.f7
            d_570_v25_: _dafny.Array
            nw93_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_570_v25_ = nw93_
            index99_ = default__.safeIndex(627, (d_570_v25_).length(0))
            (d_570_v25_)[index99_] = (self).f13
            index100_ = default__.safeIndex(627, (d_570_v25_).length(0))
            (d_570_v25_)[index100_] = (self).f13
        d_571_v26_: _dafny.Set
        d_571_v26_ = _dafny.Set({self.f7, (d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]})
        if (_dafny.Set({-627})).isdisjoint(_dafny.Set({len(d_558_v19_), len(d_571_v26_), 208})):
            d_572_v27_: C0
            nw94_ = C0()
            nw94_.ctor__((d_537_v0_ if True else d_537_v0_))
            d_572_v27_ = nw94_
            d_537_v0_ = d_537_v0_
            d_573_v29_: _dafny.Seq
            d_573_v29_ = _dafny.SeqWithoutIsStrInference([d_537_v0_, d_537_v0_, d_572_v27_.f10])
            d_574_v30_: _dafny.Seq
            d_574_v30_ = _dafny.SeqWithoutIsStrInference([d_573_v29_, d_573_v29_])
            def iife54_():
                coll20_ = _dafny.Map()
                compr_20_: _dafny.Seq
                for compr_20_ in (d_574_v30_).Elements:
                    d_575_v28_: _dafny.Seq = compr_20_
                    if (d_575_v28_) in (d_574_v30_):
                        coll20_[d_575_v28_] = p0
                return _dafny.Map(coll20_)
            (self).f7 = ((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))] if not(d_572_v27_.f10) else (len(iife54_()
            ) if d_572_v27_.f10 else (self).f6))
            (d_572_v27_).f10 = d_537_v0_
            d_576_v31_: _dafny.Seq
            d_576_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cykvhpkly"))
            d_577_v32_: _dafny.MultiSet
            d_577_v32_ = _dafny.MultiSet([(d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))], (d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]])
            index101_ = default__.safeIndex(139, (d_554_v16_).length(0))
            rhs102_ = ((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))] if d_572_v27_.f10 else 200)
            rhs103_ = d_554_v16_
            rhs104_ = default__.safeModuloInt((len(d_576_v31_)) + (len(_dafny.Set({(d_577_v32_).cardinality, -162, -131, 355}))), ((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]) - (self.f7))
            rhs105_ = (self.f7) - (self.f7)
            rhs106_ = (d_573_v29_) + ((d_573_v29_) + (d_573_v29_))
            lhs75_ = globalState
            lhs76_ = d_554_v16_
            lhs77_ = default__.safeIndex(139, (d_554_v16_).length(0))
            lhs78_ = globalState
            lhs75_.f1 = rhs102_
            d_554_v16_ = rhs103_
            lhs76_[lhs77_] = rhs104_
            lhs78_.f1 = rhs105_
            d_573_v29_ = rhs106_
        elif True:
            if d_537_v0_:
                d_578_v33_: str
                d_578_v33_ = _dafny.CodePoint('i')
                d_578_v33_ = d_578_v33_
                d_579_v34_: C0
                nw95_ = C0()
                nw95_.ctor__(d_537_v0_)
                d_579_v34_ = nw95_
                d_580_v35_: _dafny.Set
                d_580_v35_ = _dafny.Set({d_554_v16_, d_554_v16_, d_554_v16_, d_554_v16_})
                rhs107_ = d_580_v35_
                rhs108_ = default__.fm1((_dafny.MultiSet(d_558_v19_)).ispropersubset(_dafny.MultiSet(d_558_v19_)), globalState)
                lhs79_ = globalState
                d_580_v35_ = rhs107_
                lhs79_.f1 = rhs108_
                (d_579_v34_).f10 = d_537_v0_
                d_581_v36_: _dafny.Map
                d_581_v36_ = _dafny.Map({(d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]: d_554_v16_})
                d_581_v36_ = (d_581_v36_).set((self).f6, d_554_v16_)
            elif True:
                d_537_v0_ = not((d_537_v0_) == ((False) or (d_537_v0_)))
                d_537_v0_ = (d_537_v0_ if d_537_v0_ else d_537_v0_)
                d_582_v37_: D2
                d_582_v37_ = D2_DC5(d_554_v16_)
                d_583_v38_: _dafny.Seq
                d_583_v38_ = _dafny.SeqWithoutIsStrInference([d_554_v16_, d_554_v16_, (d_582_v37_).cf9, d_554_v16_, d_554_v16_])
                d_554_v16_ = (d_554_v16_ if d_537_v0_ else (d_583_v38_)[default__.safeIndex(self.f7, len(d_583_v38_))])
                d_537_v0_ = not(d_537_v0_)
                d_584_v39_: _dafny.Array
                nw96_ = _dafny.Array(None, 9)
                nw96_[int(0)] = d_537_v0_
                nw96_[int(1)] = False
                nw96_[int(2)] = d_537_v0_
                nw96_[int(3)] = d_537_v0_
                nw96_[int(4)] = d_537_v0_
                nw96_[int(5)] = True
                nw96_[int(6)] = d_537_v0_
                nw96_[int(7)] = d_537_v0_
                nw96_[int(8)] = d_537_v0_
                d_584_v39_ = nw96_
                d_585_v40_: D0
                d_585_v40_ = D0_DC1(d_537_v0_, d_584_v39_, False)
                d_586_v41_: _dafny.Seq
                d_586_v41_ = _dafny.SeqWithoutIsStrInference([d_585_v40_, d_585_v40_, d_585_v40_])
                d_587_v42_: _dafny.Map
                d_587_v42_ = _dafny.Map({default__.fm1(True, globalState): not(not(d_537_v0_))})
                d_588_v43_: _dafny.Seq
                d_588_v43_ = _dafny.SeqWithoutIsStrInference([(d_587_v42_) | (_dafny.Map({(d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))]: True})), d_587_v42_, d_587_v42_])
                rhs109_ = (d_586_v41_) + (d_586_v41_)
                rhs110_ = d_588_v43_
                rhs111_ = ((len((_dafny.Map({self.f7: d_537_v0_})).set(self.f7, d_537_v0_))) * (p0)) * (364)
                rhs112_ = d_558_v19_
                lhs80_ = globalState
                d_586_v41_ = rhs109_
                d_588_v43_ = rhs110_
                lhs80_.f1 = rhs111_
                d_558_v19_ = rhs112_
            index102_ = default__.safeIndex(139, (d_554_v16_).length(0))
            (d_554_v16_)[index102_] = (self).f6
            (self).f7 = (default__.fm1(d_537_v0_, globalState) if d_537_v0_ else (0) - ((self).f6))
            d_589_v44_: _dafny.Array
            def lambda28_(d_590_v0_):
                def lambda29_(d_591_i6_):
                    return d_590_v0_

                return lambda29_

            init15_ = lambda28_(d_537_v0_)
            nw97_ = _dafny.Array(None, 20)
            for i0_15_ in range(nw97_.length(0)):
                nw97_[i0_15_] = init15_(i0_15_)
            d_589_v44_ = nw97_
            d_592_v45_: D0
            d_592_v45_ = D0_DC1(d_537_v0_, d_589_v44_, d_537_v0_)
            d_593_v46_: int
            d_594_v47_: bool
            d_595_v48_: _dafny.Array
            out12_: int
            out13_: bool
            out14_: _dafny.Array
            out12_, out13_, out14_ = default__.m0(d_592_v45_, (_dafny.MultiSet([d_554_v16_])).cardinality, globalState)
            d_593_v46_ = out12_
            d_594_v47_ = out13_
            d_595_v48_ = out14_
            d_596_v49_: _dafny.Map
            d_596_v49_ = _dafny.Map({d_556_v17_: self.f7})
            d_597_v51_: _dafny.Map
            d_597_v51_ = _dafny.Map({d_537_v0_: d_537_v0_})
            d_598_v52_: _dafny.Map
            d_598_v52_ = _dafny.Map({p0: d_597_v51_})
            def iife55_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in (d_598_v52_).keys.Elements:
                    d_599_v50_: int = compr_21_
                    if (d_599_v50_) in (d_598_v52_):
                        coll21_[(d_599_v50_) * (p0)] = self.f7
                return _dafny.Map(coll21_)
            (globalState).f1 = len((d_596_v49_) | ((d_596_v49_).set(iife55_()
            , (self).f6)))
        r0 = d_557_v18_
        r1 = p0
        d_600_v54_: _dafny.MultiSet
        d_600_v54_ = _dafny.MultiSet([d_571_v26_])
        d_601_v55_: _dafny.Seq
        d_601_v55_ = _dafny.SeqWithoutIsStrInference([d_600_v54_])
        d_602_v56_: D1
        def iife56_():
            coll22_ = _dafny.Map()
            compr_22_: _dafny.Set
            for compr_22_ in ((d_601_v55_)[default__.safeIndex((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))], len(d_601_v55_))]).Elements:
                d_603_v53_: _dafny.Set = compr_22_
                if (d_603_v53_) in ((d_601_v55_)[default__.safeIndex((d_554_v16_)[default__.safeIndex(139, (d_554_v16_).length(0))], len(d_601_v55_))]):
                    coll22_[d_603_v53_] = d_537_v0_
            return _dafny.Map(coll22_)
        d_602_v56_ = D1_DC4((self.f7) <= (len(iife56_()
)))
        r2 = d_602_v56_
        r3 = d_537_v0_
        return r0, r1, r2, r3

    @property
    def f13(self):
        return self._f13

class C2(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6, f7):
        (self)._f6 = f6
        (self).f7 = f7

    def fm23(self, p0, p1, p2, globalState):
        if True:
            return D1_DC4(True)
        elif True:
            return D1_DC4(True)

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_604_v0_: D1
        d_604_v0_ = D1_DC3(p0, p1, self.f7)
        if ((d_604_v0_ if p0 else d_604_v0_)).cf5:
            r1 = p0
            if p0:
                d_605_v1_: _dafny.Seq
                d_605_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ignf"))
                d_606_v2_: C0
                nw98_ = C0()
                nw98_.ctor__(p0)
                d_606_v2_ = nw98_
                d_607_v3_: _dafny.Set
                d_607_v3_ = _dafny.Set({d_606_v2_, d_606_v2_, d_606_v2_})
                d_608_v4_: _dafny.Map
                d_608_v4_ = _dafny.Map({len(d_607_v3_): d_606_v2_.f10})
                d_609_v5_: D1
                d_609_v5_ = D1_DC4(d_606_v2_.f10)
                d_610_v6_: D2
                d_610_v6_ = D2_DC7(p0, (0) - (self.f7), d_608_v4_, d_609_v5_)
                d_611_v7_: str
                d_611_v7_ = _dafny.CodePoint('x')
                d_612_v8_: _dafny.Map
                d_612_v8_ = _dafny.Map({d_611_v7_: True})
                d_613_v9_: _dafny.Set
                d_613_v9_ = _dafny.Set({d_606_v2_.f10, not((D2_DC7(p0, self.f7, d_608_v4_, d_609_v5_)).cf13)})
                d_614_v11_: _dafny.Map
                d_614_v11_ = _dafny.Map({d_605_v1_: d_611_v7_})
                pat_let_tv10_ = d_613_v9_
                d_615_v12_: _dafny.Map
                def iife58_():
                    coll23_ = _dafny.Map()
                    compr_23_: _dafny.Seq
                    for compr_23_ in (d_614_v11_).keys.Elements:
                        d_616_v10_: _dafny.Seq = compr_23_
                        if (d_616_v10_) in (d_614_v11_):
                            coll23_[d_616_v10_] = self.f7
                    return _dafny.Map(coll23_)
                def iife57_(_pat_let17_0):
                    def iife59_(d_617_dt__update__tmp_h0_):
                        def iife60_(_pat_let18_0):
                            def iife61_(d_618_dt__update_hcf0_h0_):
                                return D0_DC0(d_618_dt__update_hcf0_h0_)
                            return iife61_(_pat_let18_0)
                        return iife60_(pat_let_tv10_)
                    return iife59_(_pat_let17_0)
                d_615_v12_ = _dafny.Map({iife57_(default__.fm3((d_610_v6_).cf13, len(d_612_v8_), d_613_v9_, iife58_()
                , globalState)): d_606_v2_.f10})
                d_619_v13_: D1
                d_619_v13_ = D1_DC2((d_615_v12_) | (d_615_v12_))
                rhs113_ = default__.fm1(d_606_v2_.f10, globalState)
                rhs114_ = (self).f6
                rhs115_ = d_605_v1_
                rhs116_ = d_605_v1_
                rhs117_ = D1_DC2(d_615_v12_)
                lhs81_ = self
                lhs81_.f7 = rhs113_
                r2 = rhs114_
                d_605_v1_ = rhs115_
                d_605_v1_ = rhs116_
                d_619_v13_ = rhs117_
                d_620_v14_: _dafny.Map
                d_620_v14_ = _dafny.Map({398: d_606_v2_})
                d_621_v15_: _dafny.Set
                d_621_v15_ = _dafny.Set({d_620_v14_, _dafny.Map({p1: d_606_v2_})})
                d_622_v16_: C0
                nw99_ = C0()
                nw99_.ctor__((d_621_v15_).issubset(_dafny.Set({d_620_v14_})))
                d_622_v16_ = nw99_
                d_623_v17_: _dafny.Seq
                d_623_v17_ = _dafny.SeqWithoutIsStrInference([d_622_v16_.f10, d_606_v2_.f10])
                d_623_v17_ = (d_623_v17_) + (d_623_v17_)
                d_624_v18_: C0
                nw100_ = C0()
                nw100_.ctor__(p0)
                d_624_v18_ = nw100_
                d_625_v19_: _dafny.Array
                def lambda30_(d_626_i0_):
                    return (d_626_i0_) * (923)

                init16_ = lambda30_
                nw101_ = _dafny.Array(None, 10)
                for i0_16_ in range(nw101_.length(0)):
                    nw101_[i0_16_] = init16_(i0_16_)
                d_625_v19_ = nw101_
                index103_ = default__.safeIndex(28, (d_625_v19_).length(0))
                (d_625_v19_)[index103_] = p2
                d_627_v20_: _dafny.Array
                def lambda31_(d_628_v7_):
                    def lambda32_(d_629_i1_):
                        return d_628_v7_

                    return lambda32_

                init17_ = lambda31_(d_611_v7_)
                nw102_ = _dafny.Array(None, 20)
                for i0_17_ in range(nw102_.length(0)):
                    nw102_[i0_17_] = init17_(i0_17_)
                d_627_v20_ = nw102_
                index104_ = default__.safeIndex(681, (d_627_v20_).length(0))
                (d_627_v20_)[index104_] = d_611_v7_
                index105_ = default__.safeIndex(28, (d_625_v19_).length(0))
                index106_ = default__.safeIndex(681, (d_627_v20_).length(0))
                rhs118_ = default__.safeModuloInt((self).f6, (self.f7 if False else self.f7))
                rhs119_ = d_605_v1_
                rhs120_ = d_611_v7_
                rhs121_ = d_606_v2_.f10
                lhs82_ = d_625_v19_
                lhs83_ = default__.safeIndex(28, (d_625_v19_).length(0))
                lhs84_ = d_627_v20_
                lhs85_ = default__.safeIndex(681, (d_627_v20_).length(0))
                lhs86_ = d_606_v2_
                lhs82_[lhs83_] = rhs118_
                d_605_v1_ = rhs119_
                lhs84_[lhs85_] = rhs120_
                lhs86_.f10 = rhs121_
            elif True:
                d_630_v21_: _dafny.Array
                nw103_ = _dafny.Array(False, 10)
                d_630_v21_ = nw103_
                d_631_v22_: _dafny.Seq
                d_631_v22_ = _dafny.SeqWithoutIsStrInference([p1])
                d_632_v23_: _dafny.Map
                d_632_v23_ = _dafny.Map({d_631_v22_: p0})
                rhs122_ = d_630_v21_
                rhs123_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f7 for d_633_i2_ in range(default__.abs(655))]): p0})
                rhs124_ = d_630_v21_
                d_630_v21_ = rhs122_
                d_632_v23_ = rhs123_
                d_630_v21_ = rhs124_
                d_634_v24_: _dafny.Map
                d_634_v24_ = _dafny.Map({p0: p0})
                d_635_v25_: _dafny.Seq
                d_635_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yujxjkyll"))
                d_636_v26_: _dafny.Seq
                d_636_v26_ = _dafny.SeqWithoutIsStrInference([False])
                (self).m8(d_634_v24_, D2_DC6(default__.fm24(p0, D1_DC4(p0), self.f7, globalState), d_635_v25_, len((d_636_v26_).set(default__.safeIndex((self).f6, len(d_636_v26_)), not(p0)))), d_635_v25_, globalState)
                (globalState).f1 = -712
                r0 = False
                d_634_v24_ = (d_634_v24_).set(p0, p0)
            d_637_v27_: _dafny.Array
            nw104_ = _dafny.Array(False, 3)
            d_637_v27_ = nw104_
            d_637_v27_ = d_637_v27_
            (globalState).f1 = (((self).f6) + (p1)) * ((p2) - ((self).f6))
            d_638_v28_: _dafny.Array
            def lambda33_(d_639_i3_):
                return (d_639_i3_) * ((self).f6)

            init18_ = lambda33_
            nw105_ = _dafny.Array(None, 2)
            for i0_18_ in range(nw105_.length(0)):
                nw105_[i0_18_] = init18_(i0_18_)
            d_638_v28_ = nw105_
            d_640_v29_: _dafny.Set
            d_640_v29_ = _dafny.Set({p1, (self).f6, self.f7})
            rhs125_ = d_638_v28_
            rhs126_ = d_638_v28_
            rhs127_ = d_638_v28_
            rhs128_ = d_640_v29_
            d_638_v28_ = rhs125_
            d_638_v28_ = rhs126_
            d_638_v28_ = rhs127_
            d_640_v29_ = rhs128_
        elif True:
            r0 = p0
            d_641_v30_: _dafny.Map
            d_641_v30_ = _dafny.Map({p1: _dafny.Set({p0, p0, p0, p0})})
            d_642_v31_: _dafny.Set
            d_642_v31_ = _dafny.Set({True})
            d_641_v30_ = (d_641_v30_).set(138, (d_642_v31_) - (d_642_v31_))
            d_643_v32_: _dafny.Array
            def lambda34_(d_644_i4_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqtii"))

            init19_ = lambda34_
            nw106_ = _dafny.Array(None, 23)
            for i0_19_ in range(nw106_.length(0)):
                nw106_[i0_19_] = init19_(i0_19_)
            d_643_v32_ = nw106_
            d_643_v32_ = d_643_v32_
            d_645_v33_: D1
            d_645_v33_ = D1_DC4(not(not (p0) or (p0)))
            source8_ = d_645_v33_
            if source8_.is_DC3:
                d_646___mcc_h0_ = source8_.cf5
                d_647___mcc_h1_ = source8_.cf6
                d_648___mcc_h2_ = source8_.cf7
                d_649_cf7_ = d_648___mcc_h2_
                d_650_cf6_ = d_647___mcc_h1_
                d_651_cf5_ = d_646___mcc_h0_
                d_652_v34_: _dafny.Seq
                d_652_v34_ = _dafny.SeqWithoutIsStrInference([d_651_cf5_, p0])
                d_653_v35_: _dafny.MultiSet
                d_653_v35_ = _dafny.MultiSet([d_652_v34_, d_652_v34_, d_652_v34_, _dafny.SeqWithoutIsStrInference([d_651_cf5_]), _dafny.SeqWithoutIsStrInference([d_651_cf5_])])
                d_654_v36_: _dafny.Map
                d_654_v36_ = _dafny.Map({not(((0) - (p2)) > (297)): default__.safeDivisionInt(d_650_cf6_, ((d_653_v35_)[_dafny.SeqWithoutIsStrInference([p0])] if (_dafny.SeqWithoutIsStrInference([p0])) in (d_653_v35_) else (self).f6))})
                d_654_v36_ = (d_654_v36_).set(False, (0) - ((self.f7) - (self.f7)))
                d_655_v37_: C0
                nw107_ = C0()
                nw107_.ctor__(d_651_cf5_)
                d_655_v37_ = nw107_
                d_656_v38_: _dafny.Seq
                d_656_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbw"))
                rhs129_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tttqannns"))
                rhs130_ = d_656_v38_
                rhs131_ = d_652_v34_
                d_656_v38_ = rhs129_
                d_656_v38_ = rhs130_
                d_652_v34_ = rhs131_
                d_657_v39_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 7)
                d_657_v39_ = nw108_
                index107_ = default__.safeIndex(280, (d_657_v39_).length(0))
                (d_657_v39_)[index107_] = d_649_cf7_
                index108_ = default__.safeIndex(280, (d_657_v39_).length(0))
                (d_657_v39_)[index108_] = p1
            elif source8_.is_DC4:
                d_658___mcc_h3_ = source8_.cf8
                d_659_cf8_ = d_658___mcc_h3_
                d_660_v40_: _dafny.Seq
                d_660_v40_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(self.f7, p2), self.f7, (self).f6])
                d_660_v40_ = d_660_v40_
                d_661_v41_: _dafny.Map
                d_661_v41_ = _dafny.Map({default__.fm1(p0, globalState): (True) or (p0)})
                d_662_v42_: _dafny.Map
                d_662_v42_ = _dafny.Map({self.f7: p2})
                d_661_v41_ = (d_661_v41_).set((self).f6, default__.fm0(d_659_cf8_, d_662_v42_, (self).f6, globalState))
                d_663_v43_: _dafny.Array
                def lambda35_(d_664_cf8_):
                    def lambda36_(d_665_i5_):
                        return d_664_cf8_

                    return lambda36_

                init20_ = lambda35_(d_659_cf8_)
                nw109_ = _dafny.Array(None, 23)
                for i0_20_ in range(nw109_.length(0)):
                    nw109_[i0_20_] = init20_(i0_20_)
                d_663_v43_ = nw109_
                index109_ = default__.safeIndex(709, (d_663_v43_).length(0))
                (d_663_v43_)[index109_] = p0
                index110_ = default__.safeIndex(709, (d_663_v43_).length(0))
                (d_663_v43_)[index110_] = p0
                d_666_v44_: _dafny.Map
                d_666_v44_ = _dafny.Map({(d_663_v43_)[default__.safeIndex(709, (d_663_v43_).length(0))]: p0})
                d_667_v45_: D0
                d_667_v45_ = D0_DC0(d_642_v31_)
                d_668_v46_: _dafny.Map
                d_668_v46_ = _dafny.Map({d_667_v45_: (d_663_v43_)[default__.safeIndex(709, (d_663_v43_).length(0))]})
                d_669_v47_: D1
                d_669_v47_ = D1_DC2(d_668_v46_)
                d_670_v48_: _dafny.Seq
                d_670_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpeiqsqiw"))
                d_671_v49_: D2
                d_671_v49_ = D2_DC6(d_669_v47_, d_670_v48_, default__.fm1(default__.fm0(d_659_cf8_, d_662_v42_, self.f7, globalState), globalState))
                (self).m8(d_666_v44_, d_671_v49_, (d_670_v48_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sibwdulqv"))), globalState)
            elif True:
                d_672___mcc_h4_ = source8_.cf4
                d_673_cf4_ = d_672___mcc_h4_
                d_674_v50_: _dafny.Array
                nw110_ = _dafny.Array(D0.default()(), 19)
                d_674_v50_ = nw110_
                d_675_v51_: _dafny.Array
                nw111_ = _dafny.Array(None, 1)
                nw111_[int(0)] = p0
                d_675_v51_ = nw111_
                d_676_v52_: _dafny.Set
                d_676_v52_ = _dafny.Set({d_675_v51_})
                d_677_v53_: _dafny.Array
                nw112_ = _dafny.Array(None, 20)
                nw112_[int(0)] = d_676_v52_
                nw112_[int(1)] = d_676_v52_
                nw112_[int(2)] = d_676_v52_
                nw112_[int(3)] = d_676_v52_
                nw112_[int(4)] = d_676_v52_
                nw112_[int(5)] = d_676_v52_
                nw112_[int(6)] = _dafny.Set({d_675_v51_})
                nw112_[int(7)] = (d_676_v52_) - (d_676_v52_)
                nw112_[int(8)] = d_676_v52_
                nw112_[int(9)] = d_676_v52_
                nw112_[int(10)] = d_676_v52_
                nw112_[int(11)] = _dafny.Set({d_675_v51_})
                nw112_[int(12)] = d_676_v52_
                nw112_[int(13)] = d_676_v52_
                nw112_[int(14)] = d_676_v52_
                nw112_[int(15)] = d_676_v52_
                nw112_[int(16)] = d_676_v52_
                nw112_[int(17)] = _dafny.Set({d_675_v51_, d_675_v51_})
                nw112_[int(18)] = _dafny.Set({d_675_v51_, d_675_v51_})
                nw112_[int(19)] = (d_676_v52_) | (d_676_v52_)
                d_677_v53_ = nw112_
                index111_ = default__.safeIndex(763, (d_677_v53_).length(0))
                (d_677_v53_)[index111_] = d_676_v52_
                d_678_v54_: _dafny.Map
                d_678_v54_ = _dafny.Map({p0: p0})
                index112_ = default__.safeIndex(763, (d_677_v53_).length(0))
                rhs132_ = d_674_v50_
                rhs133_ = d_642_v31_
                rhs134_ = (((d_678_v54_).set(p0, p0)) | (d_678_v54_)) == ((d_678_v54_) | (d_678_v54_))
                rhs135_ = (d_676_v52_ if (True) or (p0) else _dafny.Set({d_675_v51_}))
                lhs87_ = d_677_v53_
                lhs88_ = default__.safeIndex(763, (d_677_v53_).length(0))
                d_674_v50_ = rhs132_
                d_642_v31_ = rhs133_
                r1 = rhs134_
                lhs87_[lhs88_] = rhs135_
                r0 = p0
                d_679_v55_: _dafny.Array
                def lambda37_(d_680_p2_):
                    def lambda38_(d_681_i6_):
                        return (d_681_i6_) + (d_680_p2_)

                    return lambda38_

                init21_ = lambda37_(p2)
                nw113_ = _dafny.Array(None, 8)
                for i0_21_ in range(nw113_.length(0)):
                    nw113_[i0_21_] = init21_(i0_21_)
                d_679_v55_ = nw113_
                d_682_v56_: D2
                d_682_v56_ = D2_DC5(d_679_v55_)
                d_683_v57_: D2
                d_683_v57_ = D2_DC8(D2_DC5((d_682_v56_).cf9))
                d_684_v58_: _dafny.Map
                d_684_v58_ = _dafny.Map({d_683_v57_: default__.fm1(p0, globalState)})
                d_684_v58_ = (d_684_v58_).set(d_683_v57_, default__.safeDivisionInt(self.f7, (self).f6))
                (globalState).f1 = self.f7
            d_685_v59_: C0
            nw114_ = C0()
            nw114_.ctor__((p1) == ((self).f6))
            d_685_v59_ = nw114_
        r1 = (default__.fm1(True, globalState)) < ((self).f6)
        d_686_v60_: _dafny.Map
        d_686_v60_ = _dafny.Map({p0: p0})
        d_687_v61_: D1
        d_687_v61_ = D1_DC4(p0)
        d_688_v62_: _dafny.Map
        d_688_v62_ = _dafny.Map({p2: (d_687_v61_).cf8})
        hi7_ = (D2_DC7(p0, (self).f6, d_688_v62_, d_687_v61_)).cf14
        for d_689_i7_ in range((0) - ((self.f7) + (len(d_686_v60_))), hi7_):
            d_690_v63_: _dafny.Seq
            d_690_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgk"))
            d_691_v64_: _dafny.Map
            d_691_v64_ = _dafny.Map({p2: d_690_v63_})
            d_692_v65_: _dafny.Seq
            d_692_v65_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_693_v66_: _dafny.Array
            nw115_ = _dafny.Array(None, 19)
            nw115_[int(0)] = (self).f6
            nw115_[int(1)] = ((d_604_v0_).cf6 if (d_604_v0_).cf5 else (self).f6)
            nw115_[int(2)] = (0) - (p2)
            nw115_[int(3)] = self.f7
            nw115_[int(4)] = p1
            nw115_[int(5)] = len((d_690_v63_) + (((d_691_v64_)[d_689_i7_] if (d_689_i7_) in (d_691_v64_) else d_690_v63_)))
            nw115_[int(6)] = ((0) - ((self).f6)) + (p2)
            nw115_[int(7)] = p2
            nw115_[int(8)] = p2
            nw115_[int(9)] = p1
            nw115_[int(10)] = d_689_i7_
            nw115_[int(11)] = (0) - (default__.fm1(p0, globalState))
            nw115_[int(12)] = p2
            nw115_[int(13)] = d_689_i7_
            nw115_[int(14)] = (0) - (d_689_i7_)
            nw115_[int(15)] = (d_692_v65_)[default__.safeIndex(p2, len(d_692_v65_))]
            nw115_[int(16)] = (d_604_v0_).cf6
            nw115_[int(17)] = (0) - (p2)
            nw115_[int(18)] = default__.safeDivisionInt(default__.fm1(p0, globalState), self.f7)
            d_693_v66_ = nw115_
            index113_ = default__.safeIndex(488, (d_693_v66_).length(0))
            (d_693_v66_)[index113_] = (self).f6
            index114_ = default__.safeIndex(488, (d_693_v66_).length(0))
            (d_693_v66_)[index114_] = 242
            d_694_v67_: _dafny.Map
            d_694_v67_ = _dafny.Map({p0: (len(d_690_v63_)) * (p1)})
            d_694_v67_ = (d_694_v67_).set(p0, p2)
            d_695_v68_: _dafny.Map
            d_695_v68_ = _dafny.Map({(d_692_v65_)[default__.safeIndex(len(d_692_v65_), len(d_692_v65_))]: p1})
            d_695_v68_ = (d_695_v68_).set((self).f6, ((d_694_v67_)[p0] if (p0) in (d_694_v67_) else (self).f6))
            if p0:
                d_696_v69_: _dafny.Array
                nw116_ = _dafny.Array(False, 27)
                d_696_v69_ = nw116_
                index115_ = default__.safeIndex(578, (d_696_v69_).length(0))
                (d_696_v69_)[index115_] = ((0) - (self.f7)) <= (len(_dafny.Map({default__.fm0(False, d_695_v68_, (self).f6, globalState): (d_693_v66_)[default__.safeIndex(488, (d_693_v66_).length(0))]})))
                index116_ = default__.safeIndex(578, (d_696_v69_).length(0))
                (d_696_v69_)[index116_] = not(p0)
                index117_ = default__.safeIndex(488, (d_693_v66_).length(0))
                (d_693_v66_)[index117_] = default__.fm1(p0, globalState)
                d_697_v70_: _dafny.MultiSet
                d_697_v70_ = _dafny.MultiSet([len(d_690_v63_)])
                index118_ = default__.safeIndex(578, (d_696_v69_).length(0))
                rhs136_ = default__.safeDivisionInt((0) - ((145) * (p2)), p2)
                rhs137_ = ((604) <= ((self).f6)) == ((d_697_v70_).isdisjoint(d_697_v70_))
                lhs89_ = globalState
                lhs90_ = d_696_v69_
                lhs91_ = default__.safeIndex(578, (d_696_v69_).length(0))
                lhs89_.f1 = rhs136_
                lhs90_[lhs91_] = rhs137_
                d_698_v71_: str
                d_698_v71_ = _dafny.CodePoint('i')
                d_699_v72_: _dafny.Seq
                d_699_v72_ = _dafny.SeqWithoutIsStrInference([d_690_v63_])
                d_700_v73_: _dafny.Seq
                d_700_v73_ = _dafny.SeqWithoutIsStrInference([(d_696_v69_)[default__.safeIndex(578, (d_696_v69_).length(0))]])
                d_701_v74_: _dafny.Set
                d_701_v74_ = _dafny.Set({d_690_v63_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_702_i8_ in range(default__.abs(387))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_703_i8_ in range(default__.abs(387))]))), d_698_v71_), (d_699_v72_)[default__.safeIndex(len(d_700_v73_), len(d_699_v72_))]})
                d_701_v74_ = d_701_v74_
                d_704_v75_: _dafny.Set
                d_704_v75_ = _dafny.Set({d_692_v65_})
                rhs138_ = (d_704_v75_).ispropersubset(d_704_v75_)
                rhs139_ = (self).f6
                lhs92_ = globalState
                r0 = rhs138_
                lhs92_.f1 = rhs139_
            elif True:
                d_705_v76_: T0
                nw117_ = C1()
                nw117_.ctor__(_dafny.CodePoint('x'), d_689_i7_, (len(d_686_v60_)) - (p1))
                d_705_v76_ = nw117_
                d_706_v77_: D8
                d_706_v77_ = D8_DC18(d_705_v76_)
                d_707_v78_: D8
                d_707_v78_ = D8_DC18((d_706_v77_).cf33)
                d_705_v76_ = (d_706_v77_).cf33
                d_708_v79_: C0
                nw118_ = C0()
                nw118_.ctor__((p0) or (p0))
                d_708_v79_ = nw118_
                d_693_v66_ = d_693_v66_
                d_709_v80_: _dafny.Set
                d_709_v80_ = _dafny.Set({p0})
                d_710_v81_: D5
                d_710_v81_ = D5_DC12(len(d_709_v80_))
                d_711_v82_: _dafny.Seq
                d_711_v82_ = _dafny.SeqWithoutIsStrInference([d_708_v79_.f10])
                d_710_v81_ = default__.fm28(d_711_v82_, (0) - (d_689_i7_), globalState)
                d_712_v83_: _dafny.Array
                nw119_ = _dafny.Array(None, 12)
                nw119_[int(0)] = p0
                nw119_[int(1)] = p0
                nw119_[int(2)] = d_708_v79_.f10
                nw119_[int(3)] = d_708_v79_.f10
                nw119_[int(4)] = p0
                nw119_[int(5)] = p0
                nw119_[int(6)] = False
                nw119_[int(7)] = p0
                nw119_[int(8)] = p0
                nw119_[int(9)] = p0
                nw119_[int(10)] = p0
                nw119_[int(11)] = d_708_v79_.f10
                d_712_v83_ = nw119_
                d_713_v84_: _dafny.Seq
                d_713_v84_ = _dafny.SeqWithoutIsStrInference([d_712_v83_])
                d_714_v85_: _dafny.Array
                nw120_ = _dafny.Array(None, 11)
                nw120_[int(0)] = d_712_v83_
                nw120_[int(1)] = d_712_v83_
                nw120_[int(2)] = d_712_v83_
                nw120_[int(3)] = d_712_v83_
                nw120_[int(4)] = d_712_v83_
                nw120_[int(5)] = d_712_v83_
                nw120_[int(6)] = d_712_v83_
                nw120_[int(7)] = (d_713_v84_)[default__.safeIndex(p2, len(d_713_v84_))]
                nw120_[int(8)] = d_712_v83_
                nw120_[int(9)] = d_712_v83_
                nw120_[int(10)] = d_712_v83_
                d_714_v85_ = nw120_
                d_715_v86_: _dafny.Array
                nw121_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_715_v86_ = nw121_
                d_716_v87_: str
                d_716_v87_ = _dafny.CodePoint('o')
                index119_ = default__.safeIndex(866, (d_715_v86_).length(0))
                (d_715_v86_)[index119_] = d_716_v87_
                index120_ = default__.safeIndex(866, (d_715_v86_).length(0))
                rhs140_ = d_714_v85_
                rhs141_ = ((d_686_v60_)[p0] if (p0) in (d_686_v60_) else p0)
                rhs142_ = d_716_v87_
                rhs143_ = d_705_v76_
                rhs144_ = (d_705_v76_).f6
                lhs93_ = d_708_v79_
                lhs94_ = d_715_v86_
                lhs95_ = default__.safeIndex(866, (d_715_v86_).length(0))
                lhs96_ = globalState
                d_714_v85_ = rhs140_
                lhs93_.f10 = rhs141_
                lhs94_[lhs95_] = rhs142_
                d_705_v76_ = rhs143_
                lhs96_.f1 = rhs144_
        hi8_ = p2
        for d_717_i9_ in range((self).f6, hi8_):
            r0 = not(p0)
            if not(p0):
                d_718_v88_: _dafny.Seq
                d_718_v88_ = _dafny.SeqWithoutIsStrInference([d_717_i9_])
                d_719_v89_: _dafny.Array
                def lambda39_(d_720_i10_):
                    return (d_720_i10_) * (397)

                init22_ = lambda39_
                nw122_ = _dafny.Array(None, 12)
                for i0_22_ in range(nw122_.length(0)):
                    nw122_[i0_22_] = init22_(i0_22_)
                d_719_v89_ = nw122_
                d_721_v90_: _dafny.Set
                d_721_v90_ = _dafny.Set({d_718_v88_})
                d_722_v91_: str
                d_722_v91_ = _dafny.CodePoint('i')
                d_723_v92_: _dafny.MultiSet
                d_723_v92_ = _dafny.MultiSet([d_722_v91_, d_722_v91_])
                d_724_v93_: _dafny.Seq
                d_724_v93_ = _dafny.SeqWithoutIsStrInference([p0])
                d_725_v94_: _dafny.Array
                nw123_ = _dafny.Array(None, 23)
                nw123_[int(0)] = p2
                nw123_[int(1)] = p1
                nw123_[int(2)] = p2
                nw123_[int(3)] = 687
                nw123_[int(4)] = (0) - (d_717_i9_)
                nw123_[int(5)] = p1
                nw123_[int(6)] = default__.fm1(not(p0), globalState)
                nw123_[int(7)] = (0) - ((0) - (len(d_718_v88_)))
                nw123_[int(8)] = len(_dafny.SeqWithoutIsStrInference([d_719_v89_, d_719_v89_]))
                nw123_[int(9)] = d_717_i9_
                nw123_[int(10)] = len(d_721_v90_)
                nw123_[int(11)] = p2
                nw123_[int(12)] = (d_723_v92_).cardinality
                nw123_[int(13)] = len(d_724_v93_)
                nw123_[int(14)] = (self).f6
                nw123_[int(15)] = len(d_718_v88_)
                nw123_[int(16)] = 503
                nw123_[int(17)] = len(_dafny.SeqWithoutIsStrInference([(self).f6]))
                nw123_[int(18)] = p2
                nw123_[int(19)] = 835
                nw123_[int(20)] = self.f7
                nw123_[int(21)] = (self).f6
                nw123_[int(22)] = 610
                d_725_v94_ = nw123_
                d_726_v95_: _dafny.Map
                d_726_v95_ = _dafny.Map({d_719_v89_: p0})
                d_727_v96_: D2
                d_727_v96_ = D2_DC5(d_719_v89_)
                r0 = ((d_726_v95_)[(d_727_v96_).cf9] if ((d_727_v96_).cf9) in (d_726_v95_) else p0)
                d_728_v97_: C0
                nw124_ = C0()
                nw124_.ctor__(p0)
                d_728_v97_ = nw124_
                d_729_v98_: _dafny.MultiSet
                d_729_v98_ = _dafny.MultiSet([(p1) * (d_717_i9_)])
                d_730_v99_: D6
                d_730_v99_ = D6_DC14(p1, _dafny.MultiSet([p0, ((d_686_v60_)[d_728_v97_.f10] if (d_728_v97_.f10) in (d_686_v60_) else d_728_v97_.f10)]), len(d_686_v60_))
                d_731_v100_: _dafny.Map
                d_731_v100_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_730_v99_).cf23, p1]): p0})
                rhs145_ = (False if p0 else p0)
                rhs146_ = ((d_729_v98_)[len((d_718_v88_) + (d_718_v88_))] if (len((d_718_v88_) + (d_718_v88_))) in (d_729_v98_) else len(d_731_v100_))
                rhs147_ = d_728_v97_
                r0 = rhs145_
                r2 = rhs146_
                d_728_v97_ = rhs147_
                d_732_v101_: _dafny.Map
                d_732_v101_ = _dafny.Map({d_728_v97_.f10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksw"))})
                d_733_v102_: _dafny.Seq
                d_733_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqxjodfld"))
                d_732_v101_ = (d_732_v101_).set(d_728_v97_.f10, d_733_v102_)
                d_734_v103_: _dafny.Array
                nw125_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_734_v103_ = nw125_
                index121_ = default__.safeIndex(523, (d_734_v103_).length(0))
                (d_734_v103_)[index121_] = (d_722_v91_ if p0 else _dafny.CodePoint('j'))
                index122_ = default__.safeIndex(523, (d_734_v103_).length(0))
                (d_734_v103_)[index122_] = d_722_v91_
                d_735_v104_: _dafny.Array
                def lambda40_(d_736_v97_):
                    def lambda41_(d_737_i11_):
                        return d_736_v97_.f10

                    return lambda41_

                init23_ = lambda40_(d_728_v97_)
                nw126_ = _dafny.Array(None, 3)
                for i0_23_ in range(nw126_.length(0)):
                    nw126_[i0_23_] = init23_(i0_23_)
                d_735_v104_ = nw126_
                d_738_v105_: _dafny.Map
                d_738_v105_ = _dafny.Map({d_735_v104_: (default__.fm1(d_728_v97_.f10, globalState)) > (d_717_i9_)})
                d_738_v105_ = (d_738_v105_).set(d_735_v104_, (d_724_v93_)[default__.safeIndex(p1, len(d_724_v93_))])
            elif True:
                r1 = p0
                d_739_v106_: _dafny.Set
                d_739_v106_ = _dafny.Set({p0, True})
                d_740_v107_: D0
                d_740_v107_ = D0_DC0(d_739_v106_)
                d_741_v108_: D1
                d_741_v108_ = D1_DC2(_dafny.Map({d_740_v107_: p0}))
                d_742_v109_: _dafny.Seq
                d_742_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbkeac"))
                d_743_v110_: D2
                d_743_v110_ = D2_DC6(d_741_v108_, d_742_v109_, p1)
                (self).m8((d_686_v60_) | (d_686_v60_), d_743_v110_, (d_742_v109_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_744_i12_ in range(default__.abs(674))])), globalState)
                d_745_v111_: _dafny.MultiSet
                d_745_v111_ = _dafny.MultiSet([p2, d_717_i9_])
                d_746_v112_: _dafny.Map
                d_746_v112_ = _dafny.Map({(d_745_v111_).cardinality: self.f7})
                r1 = default__.fm0(p0, d_746_v112_, d_717_i9_, globalState)
                d_747_v113_: D9
                d_747_v113_ = D9_DC22(_dafny.SeqWithoutIsStrInference([self.f7]))
                d_748_v114_: _dafny.Map
                d_748_v114_ = _dafny.Map({d_745_v111_: (p1) != (len((d_747_v113_).cf45))})
                d_748_v114_ = (d_748_v114_).set(d_745_v111_, p0)
                d_749_v115_: _dafny.Array
                def lambda42_(d_750_i13_):
                    return _dafny.MultiSet([_dafny.MultiSet([_dafny.CodePoint('f')])])

                init24_ = lambda42_
                nw127_ = _dafny.Array(None, 15)
                for i0_24_ in range(nw127_.length(0)):
                    nw127_[i0_24_] = init24_(i0_24_)
                d_749_v115_ = nw127_
                d_751_v116_: str
                d_751_v116_ = _dafny.CodePoint('b')
                d_752_v117_: _dafny.MultiSet
                d_752_v117_ = _dafny.MultiSet([d_751_v116_, _dafny.CodePoint('t'), d_751_v116_, d_751_v116_, d_751_v116_])
                d_753_v118_: _dafny.MultiSet
                d_753_v118_ = _dafny.MultiSet([d_752_v117_, _dafny.MultiSet([d_751_v116_]), d_752_v117_])
                index123_ = default__.safeIndex(63, (d_749_v115_).length(0))
                (d_749_v115_)[index123_] = (d_753_v118_) | (d_753_v118_)
                index124_ = default__.safeIndex(63, (d_749_v115_).length(0))
                (d_749_v115_)[index124_] = d_753_v118_
            (self).f7 = p1
            d_754_v119_: _dafny.Map
            d_754_v119_ = _dafny.Map({p0: p2})
            d_755_v120_: _dafny.Seq
            d_755_v120_ = _dafny.SeqWithoutIsStrInference([p0])
            d_756_v121_: _dafny.Seq
            d_756_v121_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kifqmnwow"))
            d_757_v122_: _dafny.Array
            nw128_ = _dafny.Array(None, 17)
            nw128_[int(0)] = p1
            nw128_[int(1)] = p2
            nw128_[int(2)] = default__.fm1(p0, globalState)
            nw128_[int(3)] = self.f7
            nw128_[int(4)] = (self).f6
            nw128_[int(5)] = (0) - ((self).f6)
            nw128_[int(6)] = default__.safeDivisionInt(p1, p2)
            nw128_[int(7)] = p2
            nw128_[int(8)] = ((0) - (p2) if p0 else d_717_i9_)
            nw128_[int(9)] = (p1) + (((d_754_v119_)[p0] if (p0) in (d_754_v119_) else len(default__.fm30((d_755_v120_)[default__.safeIndex(p1, len(d_755_v120_))], p0, p1, p1, globalState))))
            nw128_[int(10)] = (d_717_i9_) * (p1)
            nw128_[int(11)] = d_717_i9_
            nw128_[int(12)] = (d_717_i9_) + (p2)
            nw128_[int(13)] = -181
            nw128_[int(14)] = len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_758_i14_ in range(default__.abs(116))])) + (d_756_v121_))
            nw128_[int(15)] = 767
            nw128_[int(16)] = 277
            d_757_v122_ = nw128_
            index125_ = default__.safeIndex(704, (d_757_v122_).length(0))
            (d_757_v122_)[index125_] = self.f7
            d_759_v123_: _dafny.Seq
            d_759_v123_ = _dafny.SeqWithoutIsStrInference([d_717_i9_])
            d_760_v124_: _dafny.Map
            d_760_v124_ = _dafny.Map({(0) - (d_717_i9_): len(_dafny.Map({(d_759_v123_)[default__.safeIndex(len(_dafny.Map({self.f7: p0})), len(d_759_v123_))]: True}))})
            index126_ = default__.safeIndex(704, (d_757_v122_).length(0))
            (d_757_v122_)[index126_] = default__.safeDivisionInt((default__.fm1(p0, globalState)) * (len(d_686_v60_)), ((0) - (((d_754_v119_)[p0] if (p0) in (d_754_v119_) else ((d_760_v124_)[p2] if (p2) in (d_760_v124_) else default__.fm1(p0, globalState))))) * ((self).f6))
        d_761_v125_: _dafny.Map
        d_761_v125_ = _dafny.Map({p2: (self).f6})
        d_762_v126_: _dafny.Seq
        d_762_v126_ = _dafny.SeqWithoutIsStrInference([len(d_761_v125_)])
        hi9_ = p1
        for d_763_i15_ in range(((d_762_v126_)[default__.safeIndex(self.f7, len(d_762_v126_))]) * (p1), hi9_):
            if (p0) or ((p0 if (d_604_v0_).cf5 else False)):
                d_764_v127_: _dafny.Seq
                d_764_v127_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vg"))
                (self).m9(p0, (len(d_764_v127_) if not(p0) else d_763_i15_), p1, -200, globalState)
                d_765_v128_: _dafny.Array
                nw129_ = _dafny.Array(None, 6)
                nw129_[int(0)] = self.f7
                nw129_[int(1)] = d_763_i15_
                nw129_[int(2)] = d_763_i15_
                nw129_[int(3)] = ((0) - (-823) if p0 else p1)
                nw129_[int(4)] = p2
                nw129_[int(5)] = p1
                d_765_v128_ = nw129_
                d_765_v128_ = d_765_v128_
                d_766_v129_: C0
                nw130_ = C0()
                nw130_.ctor__((p0) == (p0))
                d_766_v129_ = nw130_
                (d_766_v129_).f10 = p0
                (self).f7 = (0) - ((d_763_i15_ if ((d_688_v62_)[d_763_i15_] if (d_763_i15_) in (d_688_v62_) else p0) else (self.f7) - (d_763_i15_)))
            elif True:
                d_767_v130_: _dafny.Array
                nw131_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_767_v130_ = nw131_
                d_768_v131_: _dafny.Seq
                d_768_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oiiryjckw"))
                rhs148_ = (d_768_v131_) <= (d_768_v131_)
                rhs149_ = d_767_v130_
                r1 = rhs148_
                d_767_v130_ = rhs149_
                d_769_v132_: str
                d_769_v132_ = _dafny.CodePoint('t')
                d_770_v133_: D7
                d_770_v133_ = D7_DC17(d_768_v131_, default__.fm31((self).f6, p1, self.f7, d_769_v132_, globalState))
                d_771_v134_: _dafny.Seq
                d_771_v134_ = _dafny.SeqWithoutIsStrInference([p0])
                d_772_v135_: _dafny.MultiSet
                d_772_v135_ = _dafny.MultiSet([len(d_771_v134_), 885, len(d_771_v134_)])
                d_773_v136_: _dafny.MultiSet
                d_773_v136_ = _dafny.MultiSet([(d_772_v135_).cardinality])
                d_774_v137_: _dafny.Map
                d_774_v137_ = _dafny.Map({p0: len(d_762_v126_)})
                d_775_v138_: T0
                nw132_ = C1()
                nw132_.ctor__(d_769_v132_, self.f7, ((d_773_v136_)[d_763_i15_] if (d_763_i15_) in (d_773_v136_) else len(d_774_v137_)))
                d_775_v138_ = nw132_
                d_776_v139_: _dafny.Map
                d_776_v139_ = _dafny.Map({d_770_v133_: d_775_v138_})
                d_777_v140_: _dafny.Set
                d_777_v140_ = _dafny.Set({p0, p0})
                d_778_v141_: D0
                d_778_v141_ = D0_DC0(d_777_v140_)
                d_779_v142_: _dafny.Map
                d_779_v142_ = _dafny.Map({d_778_v141_: p0})
                d_780_v143_: D1
                d_780_v143_ = D1_DC2(d_779_v142_)
                d_781_v144_: D2
                d_781_v144_ = D2_DC6(d_780_v143_, d_768_v131_, (d_775_v138_).f6)
                d_782_v145_: D2
                d_782_v145_ = D2_DC7(not(p0), self.f7, d_688_v62_, (self).fm23(260, p1, d_781_v144_, globalState))
                d_783_v146_: _dafny.Set
                d_783_v146_ = _dafny.Set({self.f7})
                rhs150_ = d_776_v139_
                rhs151_ = (len(d_771_v134_)) + (((self).f6) + (len(_dafny.Map({p1: True}))))
                rhs152_ = ((d_782_v145_).cf14) <= (default__.safeModuloInt(len(d_783_v146_), d_775_v138_.f7))
                rhs153_ = p0
                lhs97_ = self
                d_776_v139_ = rhs150_
                lhs97_.f7 = rhs151_
                r0 = rhs152_
                r0 = rhs153_
                r0 = (p0 if p0 else ((d_775_v138_).f6) > (len(d_768_v131_)))
                r1 = p0
                r0 = p0
            d_784_v147_: _dafny.Array
            def lambda43_(d_785_p2_):
                def lambda44_(d_786_i16_):
                    return (d_786_i16_) + (d_785_p2_)

                return lambda44_

            init25_ = lambda43_(p2)
            nw133_ = _dafny.Array(None, 7)
            for i0_25_ in range(nw133_.length(0)):
                nw133_[i0_25_] = init25_(i0_25_)
            d_784_v147_ = nw133_
            d_787_v148_: _dafny.Seq
            d_787_v148_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jaegurgi"))
            d_788_v149_: str
            d_788_v149_ = _dafny.CodePoint('s')
            d_789_v150_: _dafny.Map
            d_789_v150_ = _dafny.Map({p0: self.f7})
            rhs154_ = d_784_v147_
            rhs155_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amfxpo"))).set(default__.safeIndex(295, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amfxpo")))), d_788_v149_)
            rhs156_ = d_784_v147_
            rhs157_ = p0
            rhs158_ = ((d_762_v126_) + (d_762_v126_)).set(default__.safeIndex((0) - (p1), len((d_762_v126_) + (d_762_v126_))), default__.safeDivisionInt((self).f6, len(d_789_v150_)))
            d_784_v147_ = rhs154_
            d_787_v148_ = rhs155_
            d_784_v147_ = rhs156_
            r1 = rhs157_
            d_762_v126_ = rhs158_
            index127_ = default__.safeIndex(488, (d_784_v147_).length(0))
            (d_784_v147_)[index127_] = ((self).f6) * (d_763_i15_)
            index128_ = default__.safeIndex(488, (d_784_v147_).length(0))
            (d_784_v147_)[index128_] = default__.safeModuloInt(p1, (self).f6)
            d_790_v151_: C0
            nw134_ = C0()
            nw134_.ctor__(p0)
            d_790_v151_ = nw134_
        d_791_v152_: _dafny.Seq
        d_791_v152_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akifd"))
        d_792_v153_: _dafny.Seq
        d_792_v153_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(d_791_v152_)})])
        if default__.fm0(((self).f6) not in ((d_792_v153_)[default__.safeIndex((self).f6, len(d_792_v153_))]), (_dafny.Map({p2: self.f7})) | (_dafny.Map({len(d_761_v125_): p2})), p2, globalState):
            d_793_v154_: _dafny.Seq
            d_793_v154_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1 for d_794_i18_ in range(default__.abs(854))]), d_762_v126_, d_762_v126_, _dafny.SeqWithoutIsStrInference([self.f7 for d_795_i19_ in range(default__.abs(-931))]), d_762_v126_])
            d_796_v155_: _dafny.Array
            nw135_ = _dafny.Array(None, 3)
            nw135_[int(0)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_797_i17_ in range(default__.abs(259))])) + (d_791_v152_)
            nw135_[int(1)] = default__.fm2((d_793_v154_)[default__.safeIndex(p1, len(d_793_v154_))], self.f7, len(_dafny.SeqWithoutIsStrInference([184 for d_798_i20_ in range(default__.abs(873))])), p0, globalState)
            nw135_[int(2)] = d_791_v152_
            d_796_v155_ = nw135_
            index129_ = default__.safeIndex(403, (d_796_v155_).length(0))
            (d_796_v155_)[index129_] = (d_791_v152_) + (d_791_v152_)
            d_799_v156_: _dafny.Array
            nw136_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_799_v156_ = nw136_
            d_800_v157_: _dafny.Array
            nw137_ = _dafny.Array(False, 14)
            d_800_v157_ = nw137_
            index130_ = default__.safeIndex(578, (d_799_v156_).length(0))
            (d_799_v156_)[index130_] = d_800_v157_
            index131_ = default__.safeIndex(403, (d_796_v155_).length(0))
            index132_ = default__.safeIndex(578, (d_799_v156_).length(0))
            rhs159_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hv"))
            rhs160_ = d_800_v157_
            lhs98_ = d_796_v155_
            lhs99_ = default__.safeIndex(403, (d_796_v155_).length(0))
            lhs100_ = d_799_v156_
            lhs101_ = default__.safeIndex(578, (d_799_v156_).length(0))
            lhs98_[lhs99_] = rhs159_
            lhs100_[lhs101_] = rhs160_
            (globalState).f1 = (self).f6
            d_801_v158_: _dafny.Array
            def lambda45_(d_802_v126_):
                def lambda46_(d_803_i21_):
                    return D9_DC22(d_802_v126_)

                return lambda46_

            init26_ = lambda45_(d_762_v126_)
            nw138_ = _dafny.Array(None, 19)
            for i0_26_ in range(nw138_.length(0)):
                nw138_[i0_26_] = init26_(i0_26_)
            d_801_v158_ = nw138_
            index133_ = default__.safeIndex(324, (d_801_v158_).length(0))
            (d_801_v158_)[index133_] = D9_DC22(d_762_v126_)
            d_804_v159_: str
            d_804_v159_ = _dafny.CodePoint('u')
            d_805_v160_: _dafny.Map
            d_805_v160_ = _dafny.Map({p0: p1})
            d_806_v161_: D5
            d_806_v161_ = D5_DC12(len(_dafny.SeqWithoutIsStrInference([d_805_v160_])))
            d_807_v162_: _dafny.Seq
            d_807_v162_ = _dafny.SeqWithoutIsStrInference([d_806_v161_, d_806_v161_])
            d_808_v163_: _dafny.Seq
            d_808_v163_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, True, p0])
            d_809_v164_: D6
            d_809_v164_ = D6_DC15(p0, p2, 79, p2)
            index134_ = default__.safeIndex(324, (d_801_v158_).length(0))
            rhs161_ = default__.safeModuloInt(len(d_807_v162_), p1)
            rhs162_ = default__.fm32(p0, not(not((d_808_v163_) <= (d_808_v163_))), p0, p1, globalState)
            rhs163_ = d_804_v159_
            rhs164_ = (d_809_v164_).cf26
            lhs102_ = globalState
            lhs103_ = d_801_v158_
            lhs104_ = default__.safeIndex(324, (d_801_v158_).length(0))
            lhs102_.f1 = rhs161_
            lhs103_[lhs104_] = rhs162_
            d_804_v159_ = rhs163_
            r1 = rhs164_
            d_810_v165_: _dafny.Array
            nw139_ = _dafny.Array(int(0), 20)
            d_810_v165_ = nw139_
            d_810_v165_ = d_810_v165_
            index135_ = default__.safeIndex(535, (d_810_v165_).length(0))
            (d_810_v165_)[index135_] = ((0) - ((d_762_v126_)[default__.safeIndex(p2, len(d_762_v126_))])) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))) for d_811_i22_ in range(default__.abs(316))])))
            index136_ = default__.safeIndex(535, (d_810_v165_).length(0))
            (d_810_v165_)[index136_] = p1
        elif True:
            d_762_v126_ = d_762_v126_
            d_812_v166_: _dafny.MultiSet
            d_812_v166_ = _dafny.MultiSet([p0, p0, False, p0, (p0) or (p0)])
            d_812_v166_ = d_812_v166_
            d_813_v167_: D10
            d_813_v167_ = D10_DC24(_dafny.CodePoint('y'))
            d_814_v168_: C1
            nw140_ = C1()
            nw140_.ctor__((d_813_v167_).cf51, p2, (self).f6)
            d_814_v168_ = nw140_
            r1 = not ((d_791_v152_) == (d_791_v152_)) or (p0)
            d_815_v169_: C1
            nw141_ = C1()
            nw141_.ctor__((d_814_v168_).f13, default__.fm1(p0, globalState), (-480) + (p1))
            d_815_v169_ = nw141_
        r0 = p0
        d_816_v170_: _dafny.MultiSet
        d_816_v170_ = _dafny.MultiSet([p0])
        r1 = ((p0 if not(p0) else p0)) in ((default__.fm33(globalState) if p0 else d_816_v170_))
        d_817_v171_: _dafny.MultiSet
        d_817_v171_ = _dafny.MultiSet([(self).f6])
        r2 = default__.safeDivisionInt(default__.safeDivisionInt(len(_dafny.Map({self.f7: p0})), self.f7), (0) - (((d_817_v171_)[p2] if (p2) in (d_817_v171_) else (d_817_v171_).cardinality)))
        return r0, r1, r2

    def m8(self, p0, p1, p2, globalState):
        d_818_v0_: bool
        d_818_v0_ = True
        rhs165_ = d_818_v0_
        rhs166_ = self.f7
        lhs105_ = globalState
        d_818_v0_ = rhs165_
        lhs105_.f1 = rhs166_
        d_819_v1_: _dafny.Seq
        d_819_v1_ = _dafny.SeqWithoutIsStrInference([p2])
        d_820_v2_: _dafny.MultiSet
        d_820_v2_ = _dafny.MultiSet([len(d_819_v1_)])
        d_821_v3_: _dafny.MultiSet
        d_821_v3_ = _dafny.MultiSet([(d_820_v2_).cardinality, len(_dafny.Map({d_818_v0_: _dafny.SeqWithoutIsStrInference([(self).f6, self.f7])}))])
        d_822_v4_: _dafny.Array
        nw142_ = _dafny.Array(None, 6)
        nw142_[int(0)] = len(p2)
        nw142_[int(1)] = (self).f6
        nw142_[int(2)] = ((d_821_v3_)[683] if (683) in (d_821_v3_) else (self).f6)
        nw142_[int(3)] = self.f7
        nw142_[int(4)] = self.f7
        nw142_[int(5)] = (self).f6
        d_822_v4_ = nw142_
        d_823_v5_: _dafny.Set
        d_823_v5_ = _dafny.Set({d_822_v4_, d_822_v4_, d_822_v4_, d_822_v4_, d_822_v4_})
        d_824_v6_: _dafny.Set
        d_824_v6_ = _dafny.Set({(0) - (len(d_823_v5_)), (self).f6, self.f7})
        def iife62_():
            coll24_ = _dafny.Set()
            compr_24_: int
            for compr_24_ in ((d_824_v6_) - (d_824_v6_)).Elements:
                d_825_v7_: int = compr_24_
                if (d_825_v7_) in ((d_824_v6_) - (d_824_v6_)):
                    coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_825_v7_, 264)]))
            return _dafny.Set(coll24_)
        d_824_v6_ = iife62_()
        
        if not(d_818_v0_):
            d_818_v0_ = d_818_v0_
            d_826_v8_: _dafny.Map
            d_826_v8_ = _dafny.Map({(self).f6: (len(_dafny.Map({556: (self).f6}))) != ((self).f6)})
            d_827_v9_: _dafny.Map
            d_827_v9_ = _dafny.Map({self.f7: (self).f6})
            d_818_v0_ = ((d_826_v8_)[(self).f6] if ((self).f6) in (d_826_v8_) else default__.fm0(d_818_v0_, d_827_v9_, 472, globalState))
            d_828_v10_: str
            d_828_v10_ = _dafny.CodePoint('g')
            d_829_v11_: C1
            nw143_ = C1()
            nw143_.ctor__(d_828_v10_, default__.safeModuloInt(-990, len((_dafny.SeqWithoutIsStrInference([self.f7])).set(default__.safeIndex(len(p2), len(_dafny.SeqWithoutIsStrInference([self.f7]))), len(_dafny.Set({151, (self).f6}))))), (-72 if d_818_v0_ else len(p2)))
            d_829_v11_ = nw143_
            (self).f7 = self.f7
            d_830_v12_: _dafny.MultiSet
            d_830_v12_ = _dafny.MultiSet([d_818_v0_])
            d_831_v13_: C1
            nw144_ = C1()
            nw144_.ctor__((d_829_v11_).f13, self.f7, default__.safeDivisionInt((d_830_v12_).cardinality, self.f7))
            d_831_v13_ = nw144_
        elif True:
            d_832_v14_: _dafny.Seq
            d_832_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jx"))
            d_833_v15_: _dafny.Array
            nw145_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_833_v15_ = nw145_
            index137_ = default__.safeIndex(166, (d_833_v15_).length(0))
            (d_833_v15_)[index137_] = (d_820_v2_) - (d_820_v2_)
            d_834_v16_: _dafny.Map
            d_834_v16_ = _dafny.Map({self.f7: d_818_v0_})
            d_835_v19_: _dafny.Seq
            d_835_v19_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_836_v20_: _dafny.Array
            nw146_ = _dafny.Array(None, 5)
            nw146_[int(0)] = d_834_v16_
            nw146_[int(1)] = _dafny.Map({self.f7: d_818_v0_})
            def iife63_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(856, 645):
                    d_837_v17_: int = compr_25_
                    if ((856) <= (d_837_v17_)) and ((d_837_v17_) < (645)):
                        coll25_[(d_837_v17_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrpgkhh"))))] = d_818_v0_
                return _dafny.Map(coll25_)
            nw146_[int(2)] = iife63_()
            
            def iife64_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in (d_835_v19_).Elements:
                    d_838_v18_: int = compr_26_
                    if (d_838_v18_) in (d_835_v19_):
                        coll26_[default__.safeModuloInt(d_838_v18_, (self).f6)] = d_818_v0_
                return _dafny.Map(coll26_)
            nw146_[int(3)] = (d_834_v16_) | (iife64_()
            )
            nw146_[int(4)] = d_834_v16_
            d_836_v20_ = nw146_
            index138_ = default__.safeIndex(166, (d_833_v15_).length(0))
            rhs167_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_839_i0_ in range(default__.abs(223))])) + (p2)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_840_i1_ in range(default__.abs(983))]))
            rhs168_ = _dafny.MultiSet(d_835_v19_)
            rhs169_ = d_818_v0_
            rhs170_ = not(((self).f6) != ((self.f7) * ((self).f6)))
            rhs171_ = d_836_v20_
            lhs106_ = d_833_v15_
            lhs107_ = default__.safeIndex(166, (d_833_v15_).length(0))
            d_832_v14_ = rhs167_
            lhs106_[lhs107_] = rhs168_
            d_818_v0_ = rhs169_
            d_818_v0_ = rhs170_
            d_836_v20_ = rhs171_
            d_841_v22_: _dafny.Array
            def lambda47_(d_842_v19_):
                def lambda48_(d_843_i2_):
                    def iife65_():
                        coll27_ = _dafny.Set()
                        compr_27_: _dafny.Seq
                        for compr_27_ in (_dafny.SeqWithoutIsStrInference([d_842_v19_, d_842_v19_])).Elements:
                            d_844_v21_: _dafny.Seq = compr_27_
                            if (d_844_v21_) in (_dafny.SeqWithoutIsStrInference([d_842_v19_, d_842_v19_])):
                                coll27_ = coll27_.union(_dafny.Set([d_844_v21_]))
                        return _dafny.Set(coll27_)
                    return (_dafny.Set({(d_842_v19_).set(default__.safeIndex((self).f6, len(d_842_v19_)), (self).f6)})) | (iife65_()
                    )

                return lambda48_

            init27_ = lambda47_(d_835_v19_)
            nw147_ = _dafny.Array(None, 15)
            for i0_27_ in range(nw147_.length(0)):
                nw147_[i0_27_] = init27_(i0_27_)
            d_841_v22_ = nw147_
            d_845_v23_: _dafny.Set
            d_845_v23_ = _dafny.Set({d_835_v19_, d_835_v19_})
            index139_ = default__.safeIndex(93, (d_841_v22_).length(0))
            (d_841_v22_)[index139_] = d_845_v23_
            d_846_v24_: _dafny.Map
            d_846_v24_ = _dafny.Map({d_835_v19_: False})
            index140_ = default__.safeIndex(93, (d_841_v22_).length(0))
            def iife66_():
                coll28_ = _dafny.Set()
                compr_28_: _dafny.Seq
                for compr_28_ in ((d_846_v24_).set(d_835_v19_, not(False))).keys.Elements:
                    d_847_v25_: _dafny.Seq = compr_28_
                    if (d_847_v25_) in ((d_846_v24_).set(d_835_v19_, not(False))):
                        coll28_ = coll28_.union(_dafny.Set([d_847_v25_]))
                return _dafny.Set(coll28_)
            (d_841_v22_)[index140_] = (iife66_()
             if d_818_v0_ else d_845_v23_)
            if not(d_818_v0_):
                (globalState).f1 = self.f7
                d_848_v26_: C0
                nw148_ = C0()
                nw148_.ctor__(d_818_v0_)
                d_848_v26_ = nw148_
                index141_ = default__.safeIndex(666, (d_822_v4_).length(0))
                (d_822_v4_)[index141_] = (0) - (self.f7)
                d_849_v27_: _dafny.Seq
                d_849_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f6: (d_821_v3_).cardinality})])
                d_850_v29_: str
                d_850_v29_ = _dafny.CodePoint('f')
                index142_ = default__.safeIndex(666, (d_822_v4_).length(0))
                def iife67_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in (_dafny.SeqWithoutIsStrInference([(self).f6 for d_851_i3_ in range(default__.abs(284))])).Elements:
                        d_852_v28_: int = compr_29_
                        if (d_852_v28_) in (_dafny.SeqWithoutIsStrInference([(self).f6 for d_851_i3_ in range(default__.abs(284))])):
                            coll29_[(d_852_v28_) * (self.f7)] = self.f7
                    return _dafny.Map(coll29_)
                rhs172_ = self.f7
                rhs173_ = (0) - (len(((d_849_v27_)[default__.safeIndex((self).f6, len(d_849_v27_))]) | (iife67_()
                )))
                rhs174_ = len((d_832_v14_) + (((d_832_v14_).set(default__.safeIndex((self).f6, len(d_832_v14_)), d_850_v29_) if d_848_v26_.f10 else _dafny.SeqWithoutIsStrInference([d_850_v29_ for d_853_i4_ in range(default__.abs(950))]))))
                rhs175_ = p2
                lhs108_ = globalState
                lhs109_ = globalState
                lhs110_ = d_822_v4_
                lhs111_ = default__.safeIndex(666, (d_822_v4_).length(0))
                lhs108_.f1 = rhs172_
                lhs109_.f1 = rhs173_
                lhs110_[lhs111_] = rhs174_
                d_832_v14_ = rhs175_
                d_854_v30_: C0
                nw149_ = C0()
                nw149_.ctor__(d_818_v0_)
                d_854_v30_ = nw149_
                index143_ = default__.safeIndex(666, (d_822_v4_).length(0))
                (d_822_v4_)[index143_] = default__.safeModuloInt(222, (self.f7) - ((self).f6))
            elif True:
                d_855_v31_: _dafny.Map
                d_855_v31_ = _dafny.Map({self.f7: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_856_i5_ in range(default__.abs(-201))]))})
                rhs176_ = True
                rhs177_ = default__.fm0(d_818_v0_, d_855_v31_, default__.safeModuloInt(734, (self).f6), globalState)
                d_818_v0_ = rhs176_
                d_818_v0_ = rhs177_
                (globalState).f1 = default__.safeDivisionInt(self.f7, self.f7)
                (globalState).f1 = 767
                (self).m9(d_818_v0_, self.f7, self.f7, (self).f6, globalState)
                (globalState).f1 = len(p2)
            d_857_v32_: _dafny.Array
            def lambda49_(d_858_v0_):
                def lambda50_(d_859_i6_):
                    return d_858_v0_

                return lambda50_

            init28_ = lambda49_(d_818_v0_)
            nw150_ = _dafny.Array(None, 16)
            for i0_28_ in range(nw150_.length(0)):
                nw150_[i0_28_] = init28_(i0_28_)
            d_857_v32_ = nw150_
            index144_ = default__.safeIndex(387, (d_857_v32_).length(0))
            (d_857_v32_)[index144_] = not (d_818_v0_) or (d_818_v0_)
            d_860_v33_: D10
            d_860_v33_ = D10_DC26(d_818_v0_)
            pat_let_tv11_ = d_818_v0_
            d_861_v34_: _dafny.Map
            def iife68_(_pat_let19_0):
                def iife69_(d_862_dt__update__tmp_h0_):
                    def iife70_(_pat_let20_0):
                        def iife71_(d_863_dt__update_hcf52_h0_):
                            return D10_DC26(d_863_dt__update_hcf52_h0_)
                        return iife71_(_pat_let20_0)
                    return iife70_(pat_let_tv11_)
                return iife69_(_pat_let19_0)
            d_861_v34_ = _dafny.Map({iife68_(d_860_v33_): d_818_v0_})
            d_864_v35_: _dafny.Map
            d_864_v35_ = _dafny.Map({len(d_861_v34_): (self).f6})
            d_865_v36_: _dafny.Set
            d_865_v36_ = _dafny.Set({d_818_v0_, default__.fm0(d_818_v0_, d_864_v35_, self.f7, globalState)})
            index145_ = default__.safeIndex(387, (d_857_v32_).length(0))
            (d_857_v32_)[index145_] = (d_865_v36_).issubset(d_865_v36_)
            d_866_v37_: _dafny.Map
            d_866_v37_ = _dafny.Map({(d_857_v32_)[default__.safeIndex(387, (d_857_v32_).length(0))]: len(p2)})
            (globalState).f1 = (((d_866_v37_)[False] if (False) in (d_866_v37_) else self.f7)) - (self.f7)
        d_867_v39_: _dafny.Map
        d_867_v39_ = _dafny.Map({(self).f6: -142})
        d_868_v40_: _dafny.Set
        d_868_v40_ = _dafny.Set({d_818_v0_, d_818_v0_, default__.fm0(d_818_v0_, d_867_v39_, 685, globalState)})
        d_869_v41_: _dafny.Set
        d_869_v41_ = _dafny.Set({d_824_v6_})
        d_870_v42_: _dafny.Array
        nw151_ = _dafny.Array(None, 24)
        nw151_[int(0)] = d_818_v0_
        nw151_[int(1)] = (d_818_v0_) == (d_818_v0_)
        nw151_[int(2)] = d_818_v0_
        nw151_[int(3)] = d_818_v0_
        nw151_[int(4)] = (d_818_v0_ if not(d_818_v0_) else d_818_v0_)
        nw151_[int(5)] = not(((self).f6) in (_dafny.Set({self.f7, self.f7, (self).f6, self.f7})))
        nw151_[int(6)] = d_818_v0_
        nw151_[int(7)] = True
        nw151_[int(8)] = not(d_818_v0_)
        nw151_[int(9)] = d_818_v0_
        nw151_[int(10)] = (d_820_v2_).ispropersubset(_dafny.MultiSet([85]))
        def iife72_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in (_dafny.Map({len(d_824_v6_): d_820_v2_})).keys.Elements:
                d_871_v38_: int = compr_30_
                if (d_871_v38_) in (_dafny.Map({len(d_824_v6_): d_820_v2_})):
                    coll30_[default__.safeDivisionInt(d_871_v38_, len(d_824_v6_))] = d_818_v0_
            return _dafny.Map(coll30_)
        nw151_[int(11)] = (779) == (len(iife72_()
        ))
        nw151_[int(12)] = d_818_v0_
        nw151_[int(13)] = False
        nw151_[int(14)] = (p2) != (p2)
        nw151_[int(15)] = (d_818_v0_) not in (d_868_v40_)
        nw151_[int(16)] = (p2) < (p2)
        nw151_[int(17)] = d_818_v0_
        nw151_[int(18)] = d_818_v0_
        nw151_[int(19)] = d_818_v0_
        nw151_[int(20)] = not(not(not((d_869_v41_).isdisjoint(default__.fm34(globalState)))))
        nw151_[int(21)] = d_818_v0_
        nw151_[int(22)] = d_818_v0_
        nw151_[int(23)] = d_818_v0_
        d_870_v42_ = nw151_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_870_v42_).length(0)):
            d_872_i7_: int = guard_loop_2_
            if (True) and (((0) <= (d_872_i7_)) and ((d_872_i7_) < ((d_870_v42_).length(0)))):
                (d_870_v42_)[(d_872_i7_)] = ((self).f6) > (self.f7)
        d_818_v0_ = d_818_v0_
        d_873_v43_: str
        d_873_v43_ = _dafny.CodePoint('k')
        d_874_v44_: C1
        nw152_ = C1()
        nw152_.ctor__(d_873_v43_, self.f7, (0) - ((self).f6))
        d_874_v44_ = nw152_

    def m9(self, p0, p1, p2, p3, globalState):
        d_875_v0_: bool
        d_875_v0_ = True
        d_876_v1_: _dafny.Seq
        d_876_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdfh"))
        d_875_v0_ = (default__.fm0(p0, _dafny.Map({len(d_876_v1_): p2}), self.f7, globalState)) or (p0)
        d_877_v2_: str
        d_877_v2_ = _dafny.CodePoint('a')
        d_878_v3_: T0
        nw153_ = C1()
        nw153_.ctor__(d_877_v2_, self.f7, p2)
        d_878_v3_ = nw153_
        rhs178_ = (d_878_v3_).f6
        rhs179_ = d_878_v3_
        lhs112_ = self
        lhs112_.f7 = rhs178_
        d_878_v3_ = rhs179_
        (d_878_v3_).f7 = d_878_v3_.f7
        if not(not(p0)):
            d_879_v4_: _dafny.Seq
            d_879_v4_ = _dafny.SeqWithoutIsStrInference([self.f7])
            (d_878_v3_).f7 = default__.safeDivisionInt(default__.safeDivisionInt(len(d_879_v4_), p1), default__.fm1(d_875_v0_, globalState))
            (globalState).f1 = (0) - (p2)
            d_880_v5_: _dafny.Seq
            d_880_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjotuokyi")), _dafny.SeqWithoutIsStrInference([d_877_v2_ for d_881_i0_ in range(default__.abs(-105))])])
            (d_878_v3_).f7 = len(((d_880_v5_) + (d_880_v5_)) + (_dafny.SeqWithoutIsStrInference([d_876_v1_ for d_882_i1_ in range(default__.abs(135))])))
            d_883_v6_: D10
            d_883_v6_ = D10_DC24(d_877_v2_)
            d_884_v7_: _dafny.Array
            nw154_ = _dafny.Array(None, 7)
            nw154_[int(0)] = d_877_v2_
            nw154_[int(1)] = d_877_v2_
            nw154_[int(2)] = d_877_v2_
            nw154_[int(3)] = _dafny.CodePoint('t')
            nw154_[int(4)] = d_877_v2_
            nw154_[int(5)] = d_877_v2_
            nw154_[int(6)] = (d_883_v6_).cf51
            d_884_v7_ = nw154_
            d_884_v7_ = d_884_v7_
            d_885_v8_: _dafny.Seq
            d_885_v8_ = _dafny.SeqWithoutIsStrInference([d_875_v0_])
            d_886_v9_: D5
            d_886_v9_ = D5_DC11(d_885_v8_)
            (d_878_v3_).f7 = len(((d_886_v9_).cf20).set(default__.safeIndex(353, len((d_886_v9_).cf20)), p0))
        elif True:
            d_887_v10_: _dafny.Map
            d_887_v10_ = _dafny.Map({d_875_v0_: p2})
            d_888_v11_: _dafny.Array
            nw155_ = _dafny.Array(None, 16)
            nw155_[int(0)] = d_887_v10_
            nw155_[int(1)] = _dafny.Map({p0: (d_878_v3_).f6})
            nw155_[int(2)] = d_887_v10_
            nw155_[int(3)] = (_dafny.Map({False: p3})) | (d_887_v10_)
            nw155_[int(4)] = d_887_v10_
            nw155_[int(5)] = (d_887_v10_).set(not(p0), p1)
            nw155_[int(6)] = (d_887_v10_) | (default__.fm30(p0, False, d_878_v3_.f7, p2, globalState))
            nw155_[int(7)] = (default__.fm30(p0, p0, -807, p2, globalState)) | (d_887_v10_)
            nw155_[int(8)] = (d_887_v10_) | (d_887_v10_)
            nw155_[int(9)] = d_887_v10_
            nw155_[int(10)] = (d_887_v10_) | (d_887_v10_)
            nw155_[int(11)] = (d_887_v10_).set(d_875_v0_, 834)
            nw155_[int(12)] = (d_887_v10_) | (d_887_v10_)
            nw155_[int(13)] = d_887_v10_
            nw155_[int(14)] = d_887_v10_
            nw155_[int(15)] = ((d_887_v10_).set(p0, p3)).set(p0, len(default__.fm35((d_878_v3_).f6, d_878_v3_.f7, globalState)))
            d_888_v11_ = nw155_
            index146_ = default__.safeIndex(857, (d_888_v11_).length(0))
            (d_888_v11_)[index146_] = (d_887_v10_) | (_dafny.Map({d_875_v0_: d_878_v3_.f7}))
            index147_ = default__.safeIndex(857, (d_888_v11_).length(0))
            (d_888_v11_)[index147_] = default__.fm30(p0, p0, (p1) - (len(_dafny.SeqWithoutIsStrInference([d_878_v3_.f7]))), (0) - ((-943) * (d_878_v3_.f7)), globalState)
            d_889_v12_: _dafny.Seq
            d_889_v12_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_890_v13_: C0
            nw156_ = C0()
            nw156_.ctor__(p0)
            d_890_v13_ = nw156_
            d_891_v14_: _dafny.MultiSet
            d_891_v14_ = _dafny.MultiSet([d_890_v13_])
            if (d_878_v3_.f7) != (default__.safeDivisionInt(len(default__.fm2((d_889_v12_).set(default__.safeIndex((d_891_v14_).cardinality, len(d_889_v12_)), p3), 986, 475, p0, globalState)), p2)):
                d_892_v15_: _dafny.Set
                d_892_v15_ = _dafny.Set({p0})
                d_893_v16_: _dafny.Map
                d_893_v16_ = _dafny.Map({p0: (d_892_v15_).issubset(_dafny.Set({d_890_v13_.f10}))})
                d_893_v16_ = d_893_v16_
                d_894_v17_: D10
                d_894_v17_ = D10_DC24(d_877_v2_)
                d_895_v18_: _dafny.Array
                nw157_ = _dafny.Array(None, 22)
                nw157_[int(0)] = d_877_v2_
                nw157_[int(1)] = d_877_v2_
                nw157_[int(2)] = d_877_v2_
                nw157_[int(3)] = default__.fm25(globalState)
                nw157_[int(4)] = d_877_v2_
                nw157_[int(5)] = d_877_v2_
                nw157_[int(6)] = (d_894_v17_).cf51
                nw157_[int(7)] = d_877_v2_
                nw157_[int(8)] = d_877_v2_
                nw157_[int(9)] = d_877_v2_
                nw157_[int(10)] = _dafny.CodePoint('p')
                nw157_[int(11)] = d_877_v2_
                nw157_[int(12)] = d_877_v2_
                nw157_[int(13)] = default__.fm25(globalState)
                nw157_[int(14)] = d_877_v2_
                nw157_[int(15)] = d_877_v2_
                nw157_[int(16)] = d_877_v2_
                nw157_[int(17)] = d_877_v2_
                nw157_[int(18)] = d_877_v2_
                nw157_[int(19)] = d_877_v2_
                nw157_[int(20)] = d_877_v2_
                nw157_[int(21)] = d_877_v2_
                d_895_v18_ = nw157_
                d_896_v19_: _dafny.Seq
                d_896_v19_ = _dafny.SeqWithoutIsStrInference([d_890_v13_, d_890_v13_, d_890_v13_])
                d_897_v20_: C0
                d_897_v20_ = (d_896_v19_)[default__.safeIndex(d_878_v3_.f7, len(d_896_v19_))]
                d_898_v21_: _dafny.Map
                d_898_v21_ = _dafny.Map({d_895_v18_: d_897_v20_})
                (d_878_v3_).f7 = len(d_898_v21_)
                d_899_v22_: _dafny.Map
                d_899_v22_ = _dafny.Map({d_895_v18_: p2})
                d_899_v22_ = (d_899_v22_).set(d_895_v18_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hh"))))
                d_892_v15_ = (d_892_v15_).intersection(d_892_v15_)
                d_900_v23_: _dafny.Seq
                d_900_v23_ = _dafny.SeqWithoutIsStrInference([d_890_v13_.f10])
                d_901_v24_: D10
                d_901_v24_ = D10_DC26(d_890_v13_.f10)
                d_902_v25_: _dafny.Array
                nw158_ = _dafny.Array(None, 24)
                nw158_[int(0)] = d_900_v23_
                nw158_[int(1)] = d_900_v23_
                nw158_[int(2)] = d_900_v23_
                nw158_[int(3)] = d_900_v23_
                nw158_[int(4)] = d_900_v23_
                nw158_[int(5)] = d_900_v23_
                nw158_[int(6)] = d_900_v23_
                nw158_[int(7)] = default__.fm27(p2, d_876_v1_, _dafny.CodePoint('b'), globalState)
                nw158_[int(8)] = (d_900_v23_ if d_890_v13_.f10 else d_900_v23_)
                nw158_[int(9)] = (d_900_v23_) + (d_900_v23_)
                nw158_[int(10)] = d_900_v23_
                nw158_[int(11)] = d_900_v23_
                nw158_[int(12)] = d_900_v23_
                nw158_[int(13)] = (d_900_v23_ if d_890_v13_.f10 else d_900_v23_)
                nw158_[int(14)] = _dafny.SeqWithoutIsStrInference([d_890_v13_.f10, d_890_v13_.f10, d_890_v13_.f10, d_890_v13_.f10, (d_900_v23_)[default__.safeIndex(p2, len(d_900_v23_))]])
                nw158_[int(15)] = _dafny.SeqWithoutIsStrInference([(d_901_v24_).cf52, True])
                nw158_[int(16)] = (d_900_v23_) + (_dafny.SeqWithoutIsStrInference([d_890_v13_.f10]))
                nw158_[int(17)] = _dafny.SeqWithoutIsStrInference([d_875_v0_, d_890_v13_.f10])
                nw158_[int(18)] = d_900_v23_
                nw158_[int(19)] = (d_900_v23_) + (d_900_v23_)
                nw158_[int(20)] = _dafny.SeqWithoutIsStrInference([True])
                nw158_[int(21)] = d_900_v23_
                nw158_[int(22)] = _dafny.SeqWithoutIsStrInference([p0, d_875_v0_, d_875_v0_])
                nw158_[int(23)] = d_900_v23_
                d_902_v25_ = nw158_
                index148_ = default__.safeIndex(520, (d_902_v25_).length(0))
                (d_902_v25_)[index148_] = d_900_v23_
                d_903_v26_: _dafny.Map
                d_903_v26_ = _dafny.Map({d_890_v13_.f10: d_900_v23_})
                index149_ = default__.safeIndex(520, (d_902_v25_).length(0))
                (d_902_v25_)[index149_] = ((d_903_v26_)[(d_890_v13_.f10 if d_890_v13_.f10 else d_875_v0_)] if ((d_890_v13_.f10 if d_890_v13_.f10 else d_875_v0_)) in (d_903_v26_) else d_900_v23_)
            elif True:
                d_904_v27_: _dafny.Array
                nw159_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_904_v27_ = nw159_
                d_905_v28_: _dafny.Array
                nw160_ = _dafny.Array(False, 17)
                d_905_v28_ = nw160_
                index150_ = default__.safeIndex(18, (d_905_v28_).length(0))
                (d_905_v28_)[index150_] = d_890_v13_.f10
                index151_ = default__.safeIndex(18, (d_905_v28_).length(0))
                rhs180_ = d_904_v27_
                rhs181_ = self.f7
                rhs182_ = d_890_v13_
                rhs183_ = p0
                lhs113_ = globalState
                lhs114_ = d_905_v28_
                lhs115_ = default__.safeIndex(18, (d_905_v28_).length(0))
                d_904_v27_ = rhs180_
                lhs113_.f1 = rhs181_
                d_890_v13_ = rhs182_
                lhs114_[lhs115_] = rhs183_
                d_906_v29_: _dafny.MultiSet
                d_906_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D0_DC0(_dafny.Set({d_875_v0_})) for d_907_i2_ in range(default__.abs(440))])])
                d_908_v30_: _dafny.Map
                d_908_v30_ = _dafny.Map({(d_878_v3_).f6: d_890_v13_.f10})
                d_909_v31_: D1
                d_909_v31_ = D1_DC4(d_890_v13_.f10)
                d_910_v32_: D2
                d_910_v32_ = D2_DC7((d_905_v28_)[default__.safeIndex(18, (d_905_v28_).length(0))], 257, d_908_v30_, d_909_v31_)
                d_911_v33_: _dafny.Set
                d_911_v33_ = _dafny.Set({not((d_905_v28_)[default__.safeIndex(18, (d_905_v28_).length(0))]), (d_910_v32_).cf13})
                d_912_v34_: D0
                d_912_v34_ = D0_DC0(d_911_v33_)
                d_913_v35_: _dafny.Seq
                d_913_v35_ = _dafny.SeqWithoutIsStrInference([d_912_v34_])
                d_914_v36_: _dafny.Seq
                d_914_v36_ = _dafny.SeqWithoutIsStrInference([p0])
                index152_ = default__.safeIndex(18, (d_905_v28_).length(0))
                rhs184_ = (D10_DC26((d_905_v28_)[default__.safeIndex(18, (d_905_v28_).length(0))])).cf52
                rhs185_ = ((d_906_v29_)[d_913_v35_] if (d_913_v35_) in (d_906_v29_) else d_878_v3_.f7)
                rhs186_ = (d_914_v36_) < (d_914_v36_)
                lhs116_ = d_905_v28_
                lhs117_ = default__.safeIndex(18, (d_905_v28_).length(0))
                lhs118_ = d_878_v3_
                lhs116_[lhs117_] = rhs184_
                lhs118_.f7 = rhs185_
                d_875_v0_ = rhs186_
                (d_890_v13_).f10 = (d_905_v28_)[default__.safeIndex(18, (d_905_v28_).length(0))]
                (d_890_v13_).f10 = not(not(not((d_877_v2_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqk"))))))
                d_915_v37_: _dafny.MultiSet
                d_915_v37_ = _dafny.MultiSet([((d_878_v3_).f6) * (d_878_v3_.f7)])
                (d_878_v3_).f7 = ((d_915_v37_)[(d_878_v3_).f6] if ((d_878_v3_).f6) in (d_915_v37_) else (288) - (d_878_v3_.f7))
            d_916_v38_: _dafny.Array
            nw161_ = _dafny.Array(int(0), 21)
            d_916_v38_ = nw161_
            index153_ = default__.safeIndex(902, (d_916_v38_).length(0))
            (d_916_v38_)[index153_] = d_878_v3_.f7
            index154_ = default__.safeIndex(902, (d_916_v38_).length(0))
            rhs187_ = self.f7
            rhs188_ = default__.safeDivisionInt((self.f7) - (len(_dafny.SeqWithoutIsStrInference([d_877_v2_ for d_917_i3_ in range(default__.abs(-376))]))), p3)
            lhs119_ = d_878_v3_
            lhs120_ = d_916_v38_
            lhs121_ = default__.safeIndex(902, (d_916_v38_).length(0))
            lhs119_.f7 = rhs187_
            lhs120_[lhs121_] = rhs188_
            d_918_v39_: D2
            d_918_v39_ = D2_DC5(d_916_v38_)
            pat_let_tv12_ = d_916_v38_
            def iife73_(_pat_let21_0):
                def iife74_(d_919_dt__update__tmp_h0_):
                    def iife75_(_pat_let22_0):
                        def iife76_(d_920_dt__update_hcf9_h0_):
                            return D2_DC5(d_920_dt__update_hcf9_h0_)
                        return iife76_(_pat_let22_0)
                    return iife75_(pat_let_tv12_)
                return iife74_(_pat_let21_0)
            d_918_v39_ = iife73_(d_918_v39_)
            d_921_v40_: _dafny.Seq
            d_921_v40_ = _dafny.SeqWithoutIsStrInference([d_875_v0_])
            d_922_v41_: D9
            d_922_v41_ = D9_DC23((d_921_v40_)[default__.safeIndex((d_878_v3_).f6, len(d_921_v40_))], 20, d_875_v0_, d_875_v0_, d_878_v3_.f7)
            index155_ = default__.safeIndex(902, (d_916_v38_).length(0))
            (d_916_v38_)[index155_] = (0) - ((0) - ((d_922_v41_).cf47))
        d_923_v42_: _dafny.Array
        def lambda51_(d_924_v0_):
            def lambda52_(d_925_i4_):
                return not(d_924_v0_)

            return lambda52_

        init29_ = lambda51_(d_875_v0_)
        nw162_ = _dafny.Array(None, 28)
        for i0_29_ in range(nw162_.length(0)):
            nw162_[i0_29_] = init29_(i0_29_)
        d_923_v42_ = nw162_
        d_923_v42_ = d_923_v42_
        d_926_v43_: _dafny.Map
        d_926_v43_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: d_875_v0_})), p3])): d_875_v0_})
        d_927_v44_: _dafny.Map
        d_927_v44_ = _dafny.Map({((d_926_v43_)[p3] if (p3) in (d_926_v43_) else True): p2})
        d_928_i5_: int
        d_928_i5_ = 0
        with _dafny.label("1"):
            pat_let_tv13_ = d_875_v0_
            pat_let_tv14_ = d_875_v0_
            pat_let_tv15_ = d_875_v0_
            pat_let_tv16_ = p0
            def lambda53_(source9_):
                if source9_.is_DC12:
                    d_934___mcc_h0_ = source9_.cf21
                    d_935_cf21_ = d_934___mcc_h0_
                    if pat_let_tv13_:
                        return not(pat_let_tv14_)
                    elif True:
                        return pat_let_tv15_
                elif True:
                    d_936___mcc_h1_ = source9_.cf20
                    d_937_cf20_ = d_936___mcc_h1_
                    return pat_let_tv16_

            while lambda53_(D5_DC12(((d_927_v44_)[p0] if (p0) in (d_927_v44_) else p3))):
                with _dafny.c_label("1"):
                    if (d_928_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_928_i5_ = (d_928_i5_) + (1)
                    d_929_v45_: _dafny.Map
                    d_929_v45_ = _dafny.Map({(self).f6: 706})
                    d_930_v46_: _dafny.Set
                    d_930_v46_ = _dafny.Set({d_878_v3_.f7, d_878_v3_.f7})
                    d_931_v47_: _dafny.Set
                    d_931_v47_ = _dafny.Set({p0})
                    d_929_v45_ = (d_929_v45_).set(-226, default__.safeModuloInt(len(d_930_v46_), len(d_931_v47_)))
                    d_876_v1_ = d_876_v1_
                    d_932_v49_: _dafny.MultiSet
                    def iife77_():
                        coll31_ = _dafny.Map()
                        compr_31_: int
                        for compr_31_ in (d_926_v43_).keys.Elements:
                            d_933_v48_: int = compr_31_
                            if (d_933_v48_) in (d_926_v43_):
                                coll31_[default__.safeModuloInt(d_933_v48_, 983)] = p3
                        return _dafny.Map(coll31_)
                    d_932_v49_ = _dafny.MultiSet([iife77_()
                    , d_929_v45_, d_929_v45_])
                    d_932_v49_ = d_932_v49_
                    d_875_v0_ = d_875_v0_
                    pass
            pass


class C3(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f6: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f6, f7):
        (self)._f6 = f6
        (self).f7 = f7

    def fm15(self, p0, p1, globalState):
        return -928

    def fm16(self, globalState):
        return not(not(((self).f6) < (self.f7)))

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_938_v0_: _dafny.Set
        d_938_v0_ = _dafny.Set({p0})
        d_939_v1_: _dafny.Map
        d_939_v1_ = _dafny.Map({p0: p1})
        d_940_v2_: _dafny.Seq
        d_940_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqysp"))
        d_941_v3_: _dafny.Seq
        d_941_v3_ = _dafny.SeqWithoutIsStrInference([d_939_v1_, d_939_v1_, (d_939_v1_).set(p0, len(d_940_v2_)), d_939_v1_, d_939_v1_])
        d_942_v4_: _dafny.Seq
        d_942_v4_ = _dafny.SeqWithoutIsStrInference([p0])
        d_943_v5_: _dafny.Map
        d_943_v5_ = _dafny.Map({self.f7: (len(d_938_v0_)) + (len((d_941_v3_)[default__.safeIndex(len(d_942_v4_), len(d_941_v3_))]))})
        d_944_v6_: str
        d_944_v6_ = _dafny.CodePoint('a')
        d_945_v7_: _dafny.Seq
        d_945_v7_ = _dafny.SeqWithoutIsStrInference([d_940_v2_, d_940_v2_])
        d_946_v8_: _dafny.Seq
        d_946_v8_ = _dafny.SeqWithoutIsStrInference([p1, self.f7, self.f7])
        d_947_v9_: _dafny.Seq
        d_947_v9_ = _dafny.SeqWithoutIsStrInference([d_946_v8_])
        d_948_v10_: _dafny.Map
        d_948_v10_ = _dafny.Map({p2: p0})
        d_949_v11_: D1
        d_949_v11_ = D1_DC4(p0)
        d_950_v12_: D2
        d_950_v12_ = D2_DC7(p0, self.f7, d_948_v10_, d_949_v11_)
        pat_let_tv17_ = p0
        pat_let_tv18_ = p0
        pat_let_tv19_ = p0
        pat_let_tv20_ = p0
        pat_let_tv21_ = d_948_v10_
        def lambda54_(source10_):
            if source10_.is_DC6:
                d_951___mcc_h0_ = source10_.cf10
                d_952___mcc_h1_ = source10_.cf11
                d_953___mcc_h2_ = source10_.cf12
                d_954_cf12_ = d_953___mcc_h2_
                d_955_cf11_ = d_952___mcc_h1_
                d_956_cf10_ = d_951___mcc_h0_
                return pat_let_tv17_
            elif source10_.is_DC7:
                d_957___mcc_h3_ = source10_.cf13
                d_958___mcc_h4_ = source10_.cf14
                d_959___mcc_h5_ = source10_.cf15
                d_960___mcc_h6_ = source10_.cf16
                d_961_cf16_ = d_960___mcc_h6_
                d_962_cf15_ = d_959___mcc_h5_
                d_963_cf14_ = d_958___mcc_h4_
                d_964_cf13_ = d_957___mcc_h3_
                return True
            elif source10_.is_DC5:
                d_965___mcc_h7_ = source10_.cf9
                d_966_cf9_ = d_965___mcc_h7_
                return pat_let_tv18_
            elif True:
                d_967___mcc_h8_ = source10_.cf17
                d_968_cf17_ = d_967___mcc_h8_
                return not (pat_let_tv19_) or (pat_let_tv20_)

        def iife78_(_pat_let23_0):
            def iife79_(d_969_dt__update__tmp_h0_):
                def iife80_(_pat_let24_0):
                    def iife81_(d_970_dt__update_hcf15_h0_):
                        return D2_DC7((d_969_dt__update__tmp_h0_).cf13, (d_969_dt__update__tmp_h0_).cf14, d_970_dt__update_hcf15_h0_, (d_969_dt__update__tmp_h0_).cf16)
                    return iife81_(_pat_let24_0)
                return iife80_(pat_let_tv21_)
            return iife79_(_pat_let23_0)
        rhs189_ = default__.fm17(_dafny.MultiSet((d_945_v7_)[default__.safeIndex(len(d_946_v8_), len(d_945_v7_))]), not (p0) or (p0), self.f7, d_947_v9_, globalState)
        rhs190_ = lambda54_(iife78_(d_950_v12_))
        rhs191_ = d_944_v6_
        rhs192_ = default__.safeDivisionInt(306, (self.f7 if p0 else len(default__.fm18(default__.fm0(p0, d_943_v5_, (self).f6, globalState), p1, globalState))))
        rhs193_ = p0
        lhs122_ = self
        d_943_v5_ = rhs189_
        r1 = rhs190_
        d_944_v6_ = rhs191_
        lhs122_.f7 = rhs192_
        r1 = rhs193_
        d_971_i0_: int
        d_971_i0_ = 0
        with _dafny.label("2"):
            while True:
                with _dafny.c_label("2"):
                    if (d_971_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_971_i0_ = (d_971_i0_) + (1)
                    d_972_v13_: int
                    d_973_v14_: int
                    d_974_v15_: int
                    out15_: int
                    out16_: int
                    out17_: int
                    out15_, out16_, out17_ = (self).m7(p0, ((d_948_v10_)[p1] if (p1) in (d_948_v10_) else p0), p2, globalState)
                    d_972_v13_ = out15_
                    d_973_v14_ = out16_
                    d_974_v15_ = out17_
                    d_975_v16_: _dafny.Array
                    nw163_ = _dafny.Array(int(0), 11)
                    d_975_v16_ = nw163_
                    d_976_v17_: _dafny.Map
                    d_976_v17_ = _dafny.Map({p0: p0})
                    index156_ = default__.safeIndex(219, (d_975_v16_).length(0))
                    (d_975_v16_)[index156_] = len((d_976_v17_) | (d_976_v17_))
                    d_977_v18_: _dafny.MultiSet
                    d_977_v18_ = _dafny.MultiSet([p0])
                    index157_ = default__.safeIndex(219, (d_975_v16_).length(0))
                    (d_975_v16_)[index157_] = ((d_977_v18_)[p0] if (p0) in (d_977_v18_) else (self).f6)
                    d_978_v19_: _dafny.MultiSet
                    d_978_v19_ = _dafny.MultiSet([default__.fm1(p0, globalState), ((0) - ((0) - (p1))) - (-411), (self.f7) - (p1)])
                    d_979_v20_: _dafny.MultiSet
                    d_979_v20_ = d_977_v18_
                    rhs194_ = (p0) and (p0)
                    rhs195_ = d_978_v19_
                    rhs196_ = (default__.fm19(p0, globalState) if not(not(p0)) else (d_977_v18_).set(True, default__.abs(p2)))
                    rhs197_ = 217
                    lhs123_ = globalState
                    r0 = rhs194_
                    d_978_v19_ = rhs195_
                    d_979_v20_ = rhs196_
                    lhs123_.f1 = rhs197_
                    r1 = p0
                    pass
            pass
        d_944_v6_ = d_944_v6_
        d_980_i1_: int
        d_980_i1_ = 0
        with _dafny.label("3"):
            while False:
                with _dafny.c_label("3"):
                    if (d_980_i1_) >= (100):
                        raise _dafny.Break("3")
                    d_980_i1_ = (d_980_i1_) + (1)
                    d_981_v21_: _dafny.Array
                    nw164_ = _dafny.Array(False, 13)
                    d_981_v21_ = nw164_
                    d_982_v22_: D0
                    d_982_v22_ = D0_DC1(p0, d_981_v21_, p0)
                    d_983_v23_: _dafny.Array
                    nw165_ = _dafny.Array(None, 1)
                    nw165_[int(0)] = d_982_v22_
                    d_983_v23_ = nw165_
                    index158_ = default__.safeIndex(782, (d_983_v23_).length(0))
                    (d_983_v23_)[index158_] = d_982_v22_
                    index159_ = default__.safeIndex(782, (d_983_v23_).length(0))
                    (d_983_v23_)[index159_] = D0_DC1(p0, d_981_v21_, p0)
                    d_984_v24_: _dafny.Map
                    d_984_v24_ = _dafny.Map({not(p0): d_944_v6_})
                    d_985_v25_: _dafny.Map
                    d_985_v25_ = _dafny.Map({not((p0) not in (d_984_v24_)): p0})
                    d_985_v25_ = (d_985_v25_).set(p0, p0)
                    if ((0) - (p1)) > (p2):
                        d_986_v26_: _dafny.Array
                        def lambda55_(d_987_i2_):
                            return (d_987_i2_) + ((self).f6)

                        init30_ = lambda55_
                        nw166_ = _dafny.Array(None, 17)
                        for i0_30_ in range(nw166_.length(0)):
                            nw166_[i0_30_] = init30_(i0_30_)
                        d_986_v26_ = nw166_
                        index160_ = default__.safeIndex(372, (d_986_v26_).length(0))
                        (d_986_v26_)[index160_] = self.f7
                        d_988_v27_: _dafny.Set
                        d_988_v27_ = _dafny.Set({p1, 792, p1})
                        index161_ = default__.safeIndex(372, (d_986_v26_).length(0))
                        rhs198_ = p0
                        rhs199_ = p2
                        rhs200_ = (0) - (len(((d_946_v8_) + (d_946_v8_)) + (_dafny.SeqWithoutIsStrInference([len(default__.fm2(d_946_v8_, len(d_988_v27_), self.f7, p0, globalState))]))))
                        rhs201_ = d_940_v2_
                        rhs202_ = default__.safeModuloInt((0) - (self.f7), len((d_946_v8_) + (d_946_v8_)))
                        lhs124_ = self
                        lhs125_ = self
                        lhs126_ = d_986_v26_
                        lhs127_ = default__.safeIndex(372, (d_986_v26_).length(0))
                        r1 = rhs198_
                        lhs124_.f7 = rhs199_
                        lhs125_.f7 = rhs200_
                        d_940_v2_ = rhs201_
                        lhs126_[lhs127_] = rhs202_
                        r1 = not(p0)
                        d_940_v2_ = d_940_v2_
                        d_989_v28_: C0
                        nw167_ = C0()
                        nw167_.ctor__(True)
                        d_989_v28_ = nw167_
                        d_990_v29_: _dafny.MultiSet
                        d_990_v29_ = _dafny.MultiSet([_dafny.MultiSet([p0, p0, p0])])
                        d_991_v30_: _dafny.MultiSet
                        d_991_v30_ = _dafny.MultiSet([p0, p0])
                        d_992_v31_: _dafny.MultiSet
                        d_992_v31_ = d_991_v30_
                        d_990_v29_ = (d_990_v29_) | ((d_990_v29_) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_992_v31_]))))
                    elif True:
                        d_993_v32_: _dafny.Set
                        d_993_v32_ = _dafny.Set({p2})
                        d_994_v33_: _dafny.Map
                        d_994_v33_ = _dafny.Map({d_993_v32_: self.f7})
                        d_994_v33_ = (d_994_v33_).set(d_993_v32_, p2)
                        index162_ = default__.safeIndex(51, (d_981_v21_).length(0))
                        (d_981_v21_)[index162_] = (len(d_948_v10_)) != (p2)
                        d_995_v34_: _dafny.Array
                        def lambda56_(d_996_v4_):
                            def lambda57_(d_997_i3_):
                                return d_996_v4_

                            return lambda57_

                        init31_ = lambda56_(d_942_v4_)
                        nw168_ = _dafny.Array(None, 6)
                        for i0_31_ in range(nw168_.length(0)):
                            nw168_[i0_31_] = init31_(i0_31_)
                        d_995_v34_ = nw168_
                        index163_ = default__.safeIndex(51, (d_981_v21_).length(0))
                        rhs203_ = -542
                        rhs204_ = p0
                        rhs205_ = d_995_v34_
                        lhs128_ = globalState
                        lhs129_ = d_981_v21_
                        lhs130_ = default__.safeIndex(51, (d_981_v21_).length(0))
                        lhs128_.f1 = rhs203_
                        lhs129_[lhs130_] = rhs204_
                        d_995_v34_ = rhs205_
                        d_998_v35_: C0
                        nw169_ = C0()
                        nw169_.ctor__(not (p0) or ((d_981_v21_)[default__.safeIndex(51, (d_981_v21_).length(0))]))
                        d_998_v35_ = nw169_
                        index164_ = default__.safeIndex(51, (d_981_v21_).length(0))
                        rhs206_ = p0
                        rhs207_ = d_998_v35_.f10
                        rhs208_ = default__.fm0((self).fm16(globalState), d_943_v5_, (self).fm15(p1, 285, globalState), globalState)
                        rhs209_ = d_998_v35_.f10
                        rhs210_ = (self).fm15(p2, p1, globalState)
                        lhs131_ = d_981_v21_
                        lhs132_ = default__.safeIndex(51, (d_981_v21_).length(0))
                        lhs133_ = d_998_v35_
                        lhs134_ = d_998_v35_
                        lhs135_ = globalState
                        r1 = rhs206_
                        lhs131_[lhs132_] = rhs207_
                        lhs133_.f10 = rhs208_
                        lhs134_.f10 = rhs209_
                        lhs135_.f1 = rhs210_
                        d_999_v36_: D0
                        d_999_v36_ = D0_DC0(d_938_v0_)
                        d_1000_v37_: _dafny.Map
                        d_1000_v37_ = _dafny.Map({d_999_v36_: (d_981_v21_)[default__.safeIndex(51, (d_981_v21_).length(0))]})
                        d_1001_v38_: D1
                        d_1001_v38_ = D1_DC2(d_1000_v37_)
                        d_1002_v39_: _dafny.Set
                        d_1002_v39_ = _dafny.Set({d_1001_v38_, d_1001_v38_, D1_DC2(d_1000_v37_), d_1001_v38_})
                        d_1003_v40_: bool
                        out18_: bool
                        out18_ = (self).m6((_dafny.Map({len(d_939_v1_): d_998_v35_.f10})) | (d_948_v10_), p1, (d_1002_v39_) | (d_1002_v39_), globalState)
                        d_1003_v40_ = out18_
                    d_949_v11_ = d_949_v11_
                    pass
            pass
        if (len(d_940_v2_)) != (p1):
            d_1004_v41_: C0
            nw170_ = C0()
            nw170_.ctor__(p0)
            d_1004_v41_ = nw170_
            d_944_v6_ = default__.fm20(d_940_v2_, d_1004_v41_.f10, self.f7, globalState)
            d_1005_v42_: _dafny.MultiSet
            d_1005_v42_ = _dafny.MultiSet([d_948_v10_, d_948_v10_])
            r1 = not((d_1005_v42_) == (d_1005_v42_))
            d_1006_v43_: _dafny.Array
            nw171_ = _dafny.Array(_dafny.MultiSet({}), 24)
            d_1006_v43_ = nw171_
            d_1007_v44_: _dafny.MultiSet
            d_1007_v44_ = _dafny.MultiSet([p2, len(d_940_v2_), 556, p1, len(d_948_v10_)])
            index165_ = default__.safeIndex(395, (d_1006_v43_).length(0))
            (d_1006_v43_)[index165_] = (d_1007_v44_) - (_dafny.MultiSet([len(d_940_v2_), (self).fm15((self).f6, 156, globalState), (d_1007_v44_).cardinality]))
            index166_ = default__.safeIndex(395, (d_1006_v43_).length(0))
            (d_1006_v43_)[index166_] = _dafny.MultiSet(d_946_v8_)
            d_1008_v45_: _dafny.Array
            nw172_ = _dafny.Array(False, 27)
            d_1008_v45_ = nw172_
            d_1009_v46_: _dafny.MultiSet
            d_1009_v46_ = _dafny.MultiSet([d_1008_v45_])
            d_1010_v47_: _dafny.Map
            d_1010_v47_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mubemdm"))})
            d_1011_v48_: bool
            out19_: bool
            out19_ = (self).m6((d_948_v10_) | (_dafny.Map({((d_1009_v46_).set(d_1008_v45_, default__.abs(len(d_1010_v47_)))).cardinality: d_1004_v41_.f10})), p2, _dafny.Set({D1_DC2(_dafny.Map({D0_DC0(d_938_v0_): d_1004_v41_.f10}))}), globalState)
            d_1011_v48_ = out19_
        elif True:
            d_1012_v49_: _dafny.Array
            nw173_ = _dafny.Array(None, 21)
            nw173_[int(0)] = d_944_v6_
            nw173_[int(1)] = _dafny.CodePoint('l')
            nw173_[int(2)] = d_944_v6_
            nw173_[int(3)] = d_944_v6_
            nw173_[int(4)] = default__.fm20(d_940_v2_, p0, len(d_948_v10_), globalState)
            nw173_[int(5)] = _dafny.CodePoint('c')
            nw173_[int(6)] = d_944_v6_
            nw173_[int(7)] = (d_944_v6_ if p0 else d_944_v6_)
            nw173_[int(8)] = d_944_v6_
            nw173_[int(9)] = (d_940_v2_)[default__.safeIndex(p1, len(d_940_v2_))]
            nw173_[int(10)] = d_944_v6_
            nw173_[int(11)] = d_944_v6_
            nw173_[int(12)] = _dafny.CodePoint('o')
            nw173_[int(13)] = d_944_v6_
            nw173_[int(14)] = d_944_v6_
            nw173_[int(15)] = d_944_v6_
            nw173_[int(16)] = _dafny.CodePoint('u')
            nw173_[int(17)] = d_944_v6_
            nw173_[int(18)] = default__.fm20(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpvixhs")), p0, p2, globalState)
            nw173_[int(19)] = d_944_v6_
            nw173_[int(20)] = (d_944_v6_ if p0 else d_944_v6_)
            d_1012_v49_ = nw173_
            d_1012_v49_ = d_1012_v49_
            d_1013_v50_: C0
            nw174_ = C0()
            nw174_.ctor__(p0)
            d_1013_v50_ = nw174_
            d_1014_v51_: C0
            d_1014_v51_ = (d_1013_v50_ if p0 else d_1013_v50_)
            source11_ = d_1014_v51_
            d_1015___mcc_h9_ = source11_
            d_1016_cf19_ = d_1015___mcc_h9_
            d_1017_v52_: _dafny.Set
            d_1017_v52_ = _dafny.Set({p1})
            d_1018_v53_: _dafny.Array
            nw175_ = _dafny.Array(int(0), 21)
            d_1018_v53_ = nw175_
            d_1019_v54_: _dafny.Map
            d_1019_v54_ = _dafny.Map({d_1018_v53_: len(d_943_v5_)})
            d_1020_v55_: _dafny.Array
            nw176_ = _dafny.Array(None, 14)
            nw176_[int(0)] = len(d_1017_v52_)
            nw176_[int(1)] = self.f7
            nw176_[int(2)] = p2
            nw176_[int(3)] = (278) + (p2)
            nw176_[int(4)] = p1
            nw176_[int(5)] = ((self).f6) - (p1)
            nw176_[int(6)] = self.f7
            nw176_[int(7)] = (0) - (p1)
            nw176_[int(8)] = (d_946_v8_)[default__.safeIndex((0) - (p1), len(d_946_v8_))]
            nw176_[int(9)] = ((self).f6) + (p2)
            nw176_[int(10)] = (self).f6
            nw176_[int(11)] = len((d_1019_v54_) | (d_1019_v54_))
            nw176_[int(12)] = self.f7
            nw176_[int(13)] = ((self).f6) + (p2)
            d_1020_v55_ = nw176_
            index167_ = default__.safeIndex(144, (d_1018_v53_).length(0))
            (d_1018_v53_)[index167_] = (self).fm15(p2, self.f7, globalState)
            d_1021_v57_: _dafny.Array
            nw177_ = _dafny.Array(None, 11)
            nw177_[int(0)] = d_1013_v50_.f10
            nw177_[int(1)] = (self).fm16(globalState)
            nw177_[int(2)] = d_1013_v50_.f10
            nw177_[int(3)] = not(d_1016_cf19_.f10)
            nw177_[int(4)] = (d_1013_v50_.f10 if d_1013_v50_.f10 else p0)
            nw177_[int(5)] = not (d_1013_v50_.f10) or (d_1016_cf19_.f10)
            nw177_[int(6)] = (p0) == (p0)
            def iife82_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(283, -466):
                    d_1022_v56_: int = compr_32_
                    if ((283) <= (d_1022_v56_)) and ((d_1022_v56_) < (-466)):
                        coll32_[default__.safeModuloInt(d_1022_v56_, self.f7)] = len(d_938_v0_)
                return _dafny.Map(coll32_)
            nw177_[int(7)] = not (d_1013_v50_.f10) or (default__.fm0(d_1016_cf19_.f10, iife82_()
            , (self).f6, globalState))
            nw177_[int(8)] = d_1016_cf19_.f10
            nw177_[int(9)] = d_1016_cf19_.f10
            nw177_[int(10)] = not(d_1016_cf19_.f10)
            d_1021_v57_ = nw177_
            index168_ = default__.safeIndex(701, (d_1021_v57_).length(0))
            (d_1021_v57_)[index168_] = d_1016_cf19_.f10
            index169_ = default__.safeIndex(144, (d_1018_v53_).length(0))
            index170_ = default__.safeIndex(701, (d_1021_v57_).length(0))
            rhs211_ = p2
            rhs212_ = d_1013_v50_.f10
            rhs213_ = (p2) - ((p2) + (p1))
            rhs214_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1020_v55_, d_1020_v55_, d_1018_v53_, d_1020_v55_])), p2)
            lhs136_ = d_1018_v53_
            lhs137_ = default__.safeIndex(144, (d_1018_v53_).length(0))
            lhs138_ = d_1021_v57_
            lhs139_ = default__.safeIndex(701, (d_1021_v57_).length(0))
            lhs140_ = globalState
            lhs141_ = globalState
            lhs136_[lhs137_] = rhs211_
            lhs138_[lhs139_] = rhs212_
            lhs140_.f1 = rhs213_
            lhs141_.f1 = rhs214_
            index171_ = default__.safeIndex(701, (d_1021_v57_).length(0))
            (d_1021_v57_)[index171_] = (False) and (default__.fm0((d_1021_v57_)[default__.safeIndex(701, (d_1021_v57_).length(0))], d_943_v5_, self.f7, globalState))
            d_1013_v50_ = d_1016_cf19_
            d_1023_v58_: C0
            nw178_ = C0()
            nw178_.ctor__((d_940_v2_) != (d_940_v2_))
            d_1023_v58_ = nw178_
            d_1024_v59_: D0
            d_1024_v59_ = D0_DC0(d_938_v0_)
            d_1025_v60_: _dafny.Map
            d_1025_v60_ = _dafny.Map({d_1024_v59_: d_1013_v50_.f10})
            d_1026_v61_: _dafny.Map
            d_1026_v61_ = _dafny.Map({p0: d_1025_v60_})
            d_1027_v62_: D1
            d_1027_v62_ = D1_DC2(((d_1026_v61_)[d_1013_v50_.f10] if (d_1013_v50_.f10) in (d_1026_v61_) else _dafny.Map({D0_DC0(d_938_v0_): False})))
            d_1027_v62_ = d_1027_v62_
            d_1028_v63_: _dafny.Map
            d_1028_v63_ = _dafny.Map({p0: d_1013_v50_.f10})
            d_1029_v64_: _dafny.Seq
            d_1029_v64_ = _dafny.SeqWithoutIsStrInference([d_1028_v63_])
            d_1030_v65_: _dafny.MultiSet
            d_1030_v65_ = _dafny.MultiSet([self.f7, p1])
            d_1031_v66_: _dafny.Array
            def lambda58_(d_1032_v50_):
                def lambda59_(d_1033_i4_):
                    return d_1032_v50_.f10

                return lambda59_

            init32_ = lambda58_(d_1013_v50_)
            nw179_ = _dafny.Array(None, 20)
            for i0_32_ in range(nw179_.length(0)):
                nw179_[i0_32_] = init32_(i0_32_)
            d_1031_v66_ = nw179_
            d_1034_v67_: D0
            d_1034_v67_ = D0_DC1(d_1013_v50_.f10, d_1031_v66_, d_1013_v50_.f10)
            d_1035_v68_: _dafny.MultiSet
            d_1035_v68_ = _dafny.MultiSet([d_1013_v50_.f10, True, p0])
            d_1036_v69_: _dafny.Array
            nw180_ = _dafny.Array(None, 27)
            nw180_[int(0)] = (p0 if d_1013_v50_.f10 else True)
            nw180_[int(1)] = p0
            nw180_[int(2)] = not((p1) == (len(((d_1029_v64_)[default__.safeIndex((self).f6, len(d_1029_v64_))]).set(d_1013_v50_.f10, d_1013_v50_.f10))))
            nw180_[int(3)] = (d_950_v12_).cf13
            nw180_[int(4)] = d_1013_v50_.f10
            nw180_[int(5)] = not(p0)
            nw180_[int(6)] = d_1013_v50_.f10
            nw180_[int(7)] = d_1013_v50_.f10
            nw180_[int(8)] = False
            nw180_[int(9)] = d_1013_v50_.f10
            nw180_[int(10)] = (d_938_v0_).issubset(d_938_v0_)
            nw180_[int(11)] = p0
            nw180_[int(12)] = d_1013_v50_.f10
            nw180_[int(13)] = p0
            nw180_[int(14)] = d_1013_v50_.f10
            nw180_[int(15)] = (d_1030_v65_).ispropersubset(_dafny.MultiSet(d_946_v8_))
            nw180_[int(16)] = d_1013_v50_.f10
            nw180_[int(17)] = (p2) >= (p2)
            nw180_[int(18)] = p0
            nw180_[int(19)] = (False) or (True)
            nw180_[int(20)] = not((d_1013_v50_.f10 if d_1013_v50_.f10 else (d_1034_v67_).cf3))
            nw180_[int(21)] = (d_1035_v68_) == (_dafny.MultiSet([default__.fm0(d_1013_v50_.f10, _dafny.Map({p2: p2}), p1, globalState)]))
            nw180_[int(22)] = default__.fm0(d_1013_v50_.f10, d_943_v5_, p1, globalState)
            nw180_[int(23)] = p0
            nw180_[int(24)] = p0
            nw180_[int(25)] = (self).fm16(globalState)
            nw180_[int(26)] = d_1013_v50_.f10
            d_1036_v69_ = nw180_
            index172_ = default__.safeIndex(401, (d_1036_v69_).length(0))
            (d_1036_v69_)[index172_] = p0
            d_1037_v70_: _dafny.Set
            d_1037_v70_ = _dafny.Set({(self).f6})
            d_1038_v72_: _dafny.Array
            nw181_ = _dafny.Array(D2.default()(), 16)
            d_1038_v72_ = nw181_
            d_1039_v73_: _dafny.Set
            d_1039_v73_ = _dafny.Set({d_1038_v72_, d_1038_v72_})
            index173_ = default__.safeIndex(401, (d_1036_v69_).length(0))
            def iife83_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in _dafny.IntegerRange(288, 394):
                    d_1040_v71_: int = compr_33_
                    if ((288) <= (d_1040_v71_)) and ((d_1040_v71_) < (394)):
                        coll33_ = coll33_.union(_dafny.Set([default__.safeDivisionInt(d_1040_v71_, p1)]))
                return _dafny.Set(coll33_)
            rhs215_ = (p0 if (d_1037_v70_).isdisjoint(iife83_()
            ) else True)
            rhs216_ = not(((d_1039_v73_) | (d_1039_v73_)).issubset((d_1039_v73_).intersection(d_1039_v73_)))
            rhs217_ = (0) - ((p2) - (len(d_943_v5_)))
            lhs142_ = d_1036_v69_
            lhs143_ = default__.safeIndex(401, (d_1036_v69_).length(0))
            lhs144_ = self
            lhs142_[lhs143_] = rhs215_
            r1 = rhs216_
            lhs144_.f7 = rhs217_
            r0 = p0
        if (self).fm16(globalState):
            d_939_v1_ = (d_939_v1_).set(p0, (len(d_940_v2_)) * ((0) - (p2)))
            rhs218_ = (d_942_v4_ if True else d_942_v4_)
            rhs219_ = d_942_v4_
            d_942_v4_ = rhs218_
            d_942_v4_ = rhs219_
            r2 = p2
            d_1041_v74_: _dafny.Map
            d_1041_v74_ = _dafny.Map({d_944_v6_: p2})
            r1 = (d_942_v4_)[default__.safeIndex((((d_1041_v74_)[d_944_v6_] if (d_944_v6_) in (d_1041_v74_) else len(d_946_v8_))) + (p1), len(d_942_v4_))]
            d_1042_v75_: _dafny.Map
            d_1042_v75_ = _dafny.Map({p0: p0})
            d_1043_v76_: _dafny.Set
            d_1043_v76_ = _dafny.Set({d_942_v4_})
            d_1044_v77_: _dafny.Array
            nw182_ = _dafny.Array(None, 11)
            nw182_[int(0)] = default__.safeModuloInt((self).fm15(p1, p1, globalState), (0) - (len(d_1042_v75_)))
            nw182_[int(1)] = len(d_1043_v76_)
            nw182_[int(2)] = 886
            nw182_[int(3)] = p1
            nw182_[int(4)] = p2
            nw182_[int(5)] = p1
            nw182_[int(6)] = (p2) * ((d_946_v8_)[default__.safeIndex(p2, len(d_946_v8_))])
            nw182_[int(7)] = self.f7
            nw182_[int(8)] = p1
            nw182_[int(9)] = (self).f6
            nw182_[int(10)] = p2
            d_1044_v77_ = nw182_
            index174_ = default__.safeIndex(566, (d_1044_v77_).length(0))
            (d_1044_v77_)[index174_] = p2
            index175_ = default__.safeIndex(566, (d_1044_v77_).length(0))
            rhs220_ = False
            rhs221_ = (((d_939_v1_)[p0] if (p0) in (d_939_v1_) else self.f7)) * (default__.safeDivisionInt((self).f6, (self).f6))
            lhs145_ = d_1044_v77_
            lhs146_ = default__.safeIndex(566, (d_1044_v77_).length(0))
            r1 = rhs220_
            lhs145_[lhs146_] = rhs221_
        elif True:
            d_1045_v78_: C0
            nw183_ = C0()
            nw183_.ctor__(p0)
            d_1045_v78_ = nw183_
            d_1046_v79_: C0
            d_1046_v79_ = d_1045_v78_
            rhs222_ = d_1045_v78_
            rhs223_ = p0
            d_1046_v79_ = rhs222_
            r1 = rhs223_
            d_1047_v80_: _dafny.Array
            def lambda60_(d_1048_v4_):
                def lambda61_(d_1049_i5_):
                    return d_1048_v4_

                return lambda61_

            init33_ = lambda60_(d_942_v4_)
            nw184_ = _dafny.Array(None, 20)
            for i0_33_ in range(nw184_.length(0)):
                nw184_[i0_33_] = init33_(i0_33_)
            d_1047_v80_ = nw184_
            index176_ = default__.safeIndex(844, (d_1047_v80_).length(0))
            (d_1047_v80_)[index176_] = d_942_v4_
            index177_ = default__.safeIndex(844, (d_1047_v80_).length(0))
            (d_1047_v80_)[index177_] = (d_942_v4_) + (d_942_v4_)
            d_1050_v81_: _dafny.Map
            d_1050_v81_ = _dafny.Map({p1: default__.fm21(globalState)})
            d_948_v10_ = (d_948_v10_).set(self.f7, (((d_1050_v81_)[self.f7] if (self.f7) in (d_1050_v81_) else (d_1047_v80_)[default__.safeIndex(844, (d_1047_v80_).length(0))])) < ((d_1047_v80_)[default__.safeIndex(844, (d_1047_v80_).length(0))]))
            d_1051_v82_: _dafny.Array
            nw185_ = _dafny.Array(False, 11)
            d_1051_v82_ = nw185_
            index178_ = default__.safeIndex(317, (d_1051_v82_).length(0))
            (d_1051_v82_)[index178_] = p0
            index179_ = default__.safeIndex(317, (d_1051_v82_).length(0))
            (d_1051_v82_)[index179_] = ((self).f6) != ((0) - (default__.fm1(d_1045_v78_.f10, globalState)))
            d_1052_v83_: _dafny.Array
            nw186_ = _dafny.Array(_dafny.Map({}), 2)
            d_1052_v83_ = nw186_
            d_1053_v84_: _dafny.MultiSet
            d_1053_v84_ = _dafny.MultiSet([d_944_v6_, d_944_v6_, d_944_v6_, d_944_v6_, d_944_v6_])
            index180_ = default__.safeIndex(922, (d_1052_v83_).length(0))
            (d_1052_v83_)[index180_] = _dafny.Map({d_1053_v84_: (d_1051_v82_)[default__.safeIndex(317, (d_1051_v82_).length(0))]})
            d_1054_v85_: _dafny.Map
            d_1054_v85_ = _dafny.Map({d_1053_v84_: not(d_1045_v78_.f10)})
            index181_ = default__.safeIndex(922, (d_1052_v83_).length(0))
            rhs224_ = (0) - (len(_dafny.SeqWithoutIsStrInference([d_944_v6_ for d_1055_i6_ in range(default__.abs(280))])))
            rhs225_ = d_1045_v78_.f10
            rhs226_ = d_944_v6_
            rhs227_ = (_dafny.Map({d_1053_v84_: True})) | (d_1054_v85_)
            lhs147_ = globalState
            lhs148_ = d_1045_v78_
            lhs149_ = d_1052_v83_
            lhs150_ = default__.safeIndex(922, (d_1052_v83_).length(0))
            lhs147_.f1 = rhs224_
            lhs148_.f10 = rhs225_
            d_944_v6_ = rhs226_
            lhs149_[lhs150_] = rhs227_
        r0 = p0
        d_1056_v86_: _dafny.Set
        d_1056_v86_ = _dafny.Set({(self).f6, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oh"))).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oh")))), _dafny.CodePoint('f')))})
        r1 = (d_1056_v86_).isdisjoint(d_1056_v86_)
        r2 = ((self).f6) + (324)
        return r0, r1, r2

    def m6(self, p0, p1, p2, globalState):
        r0: bool = False
        d_1057_v0_: bool
        d_1057_v0_ = False
        d_1058_i0_: int
        d_1058_i0_ = 0
        with _dafny.label("4"):
            while not(d_1057_v0_):
                with _dafny.c_label("4"):
                    if (d_1058_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_1058_i0_ = (d_1058_i0_) + (1)
                    d_1059_v1_: _dafny.Seq
                    d_1059_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ea"))
                    d_1060_v2_: _dafny.Array
                    nw187_ = _dafny.Array(int(0), 13)
                    d_1060_v2_ = nw187_
                    d_1061_v3_: D2
                    d_1061_v3_ = D2_DC5(d_1060_v2_)
                    d_1062_v4_: _dafny.Array
                    nw188_ = _dafny.Array(None, 20)
                    nw188_[int(0)] = d_1060_v2_
                    nw188_[int(1)] = d_1060_v2_
                    nw188_[int(2)] = d_1060_v2_
                    nw188_[int(3)] = d_1060_v2_
                    nw188_[int(4)] = d_1060_v2_
                    nw188_[int(5)] = d_1060_v2_
                    nw188_[int(6)] = d_1060_v2_
                    nw188_[int(7)] = d_1060_v2_
                    nw188_[int(8)] = d_1060_v2_
                    nw188_[int(9)] = d_1060_v2_
                    nw188_[int(10)] = d_1060_v2_
                    nw188_[int(11)] = d_1060_v2_
                    nw188_[int(12)] = (d_1061_v3_).cf9
                    nw188_[int(13)] = d_1060_v2_
                    nw188_[int(14)] = d_1060_v2_
                    nw188_[int(15)] = d_1060_v2_
                    nw188_[int(16)] = (d_1061_v3_).cf9
                    nw188_[int(17)] = d_1060_v2_
                    nw188_[int(18)] = d_1060_v2_
                    nw188_[int(19)] = d_1060_v2_
                    d_1062_v4_ = nw188_
                    index182_ = default__.safeIndex(974, (d_1062_v4_).length(0))
                    (d_1062_v4_)[index182_] = (d_1060_v2_ if d_1057_v0_ else d_1060_v2_)
                    d_1063_v5_: str
                    d_1063_v5_ = _dafny.CodePoint('g')
                    d_1064_v6_: _dafny.Map
                    d_1064_v6_ = _dafny.Map({len(default__.fm18(d_1057_v0_, (self).f6, globalState)): d_1063_v5_})
                    d_1065_v7_: _dafny.Map
                    d_1065_v7_ = _dafny.Map({False: len(d_1064_v6_)})
                    d_1066_v8_: _dafny.Seq
                    d_1066_v8_ = _dafny.SeqWithoutIsStrInference([self.f7])
                    d_1067_v9_: _dafny.Seq
                    d_1067_v9_ = _dafny.SeqWithoutIsStrInference([p1, 938, len(d_1066_v8_)])
                    d_1068_v10_: _dafny.Seq
                    d_1068_v10_ = _dafny.SeqWithoutIsStrInference([d_1057_v0_])
                    index183_ = default__.safeIndex(974, (d_1062_v4_).length(0))
                    rhs228_ = (p1) + ((len(d_1065_v7_)) + ((self).f6))
                    rhs229_ = not(not(not(not((_dafny.SeqWithoutIsStrInference([default__.fm1(d_1057_v0_, globalState)])) < (_dafny.SeqWithoutIsStrInference([len(d_1067_v9_)]))))))
                    rhs230_ = (d_1059_v1_) + (d_1059_v1_)
                    rhs231_ = d_1060_v2_
                    rhs232_ = (d_1068_v10_)[default__.safeIndex(p1, len(d_1068_v10_))]
                    lhs151_ = globalState
                    lhs152_ = d_1062_v4_
                    lhs153_ = default__.safeIndex(974, (d_1062_v4_).length(0))
                    lhs151_.f1 = rhs228_
                    r0 = rhs229_
                    d_1059_v1_ = rhs230_
                    lhs152_[lhs153_] = rhs231_
                    r0 = rhs232_
                    d_1065_v7_ = (d_1065_v7_).set(d_1057_v0_, (self).f6)
                    d_1057_v0_ = d_1057_v0_
                    d_1069_v11_: C0
                    nw189_ = C0()
                    nw189_.ctor__(d_1057_v0_)
                    d_1069_v11_ = nw189_
                    d_1069_v11_ = d_1069_v11_
                    pass
            pass
        d_1070_v12_: _dafny.Set
        d_1070_v12_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tayy")))})
        d_1071_v13_: _dafny.Map
        d_1071_v13_ = _dafny.Map({d_1057_v0_: d_1057_v0_})
        d_1072_v14_: _dafny.Map
        d_1072_v14_ = _dafny.Map({((d_1071_v13_)[d_1057_v0_] if (d_1057_v0_) in (d_1071_v13_) else (self).fm16(globalState)): d_1070_v12_})
        d_1073_v16_: _dafny.Seq
        d_1073_v16_ = _dafny.SeqWithoutIsStrInference([not(not(d_1057_v0_)), (self).fm16(globalState)])
        d_1074_v17_: _dafny.Array
        nw190_ = _dafny.Array(False, 10)
        d_1074_v17_ = nw190_
        d_1075_v18_: D0
        d_1075_v18_ = D0_DC1(d_1057_v0_, d_1074_v17_, not(d_1057_v0_))
        d_1076_v19_: _dafny.Map
        d_1076_v19_ = _dafny.Map({p1: p0})
        d_1077_v20_: _dafny.Array
        def lambda62_(d_1078_i1_):
            return default__.safeModuloInt(d_1078_i1_, (self).f6)

        init34_ = lambda62_
        nw191_ = _dafny.Array(None, 6)
        for i0_34_ in range(nw191_.length(0)):
            nw191_[i0_34_] = init34_(i0_34_)
        d_1077_v20_ = nw191_
        d_1079_v21_: _dafny.MultiSet
        d_1079_v21_ = _dafny.MultiSet([d_1077_v20_, d_1077_v20_])
        d_1080_v22_: _dafny.Array
        nw192_ = _dafny.Array(None, 21)
        nw192_[int(0)] = d_1057_v0_
        nw192_[int(1)] = (self.f7) > ((self).f6)
        nw192_[int(2)] = d_1057_v0_
        nw192_[int(3)] = d_1057_v0_
        def iife84_():
            coll34_ = _dafny.Set()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(-142, 168):
                d_1081_v15_: int = compr_34_
                if ((-142) <= (d_1081_v15_)) and ((d_1081_v15_) < (168)):
                    coll34_ = coll34_.union(_dafny.Set([(d_1081_v15_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkk"))))]))
            return _dafny.Set(coll34_)
        nw192_[int(4)] = (d_1070_v12_).isdisjoint(((d_1072_v14_)[d_1057_v0_] if (d_1057_v0_) in (d_1072_v14_) else iife84_()
        ))
        nw192_[int(5)] = not(d_1057_v0_)
        nw192_[int(6)] = (d_1073_v16_)[default__.safeIndex(p1, len(d_1073_v16_))]
        nw192_[int(7)] = d_1057_v0_
        nw192_[int(8)] = d_1057_v0_
        nw192_[int(9)] = not (not((d_1075_v18_).cf3)) or (not(d_1057_v0_))
        nw192_[int(10)] = (d_1073_v16_) < (d_1073_v16_)
        nw192_[int(11)] = d_1057_v0_
        nw192_[int(12)] = d_1057_v0_
        nw192_[int(13)] = (p1) not in (((d_1076_v19_)[p1] if (p1) in (d_1076_v19_) else _dafny.Map({(self).f6: d_1057_v0_})))
        nw192_[int(14)] = d_1057_v0_
        nw192_[int(15)] = ((p0)[p1] if (p1) in (p0) else d_1057_v0_)
        nw192_[int(16)] = d_1057_v0_
        nw192_[int(17)] = d_1057_v0_
        nw192_[int(18)] = d_1057_v0_
        nw192_[int(19)] = (d_1079_v21_).issubset(d_1079_v21_)
        nw192_[int(20)] = d_1057_v0_
        d_1080_v22_ = nw192_
        index184_ = default__.safeIndex(143, (d_1080_v22_).length(0))
        (d_1080_v22_)[index184_] = False
        index185_ = default__.safeIndex(11, (d_1080_v22_).length(0))
        (d_1080_v22_)[index185_] = d_1057_v0_
        d_1082_v23_: _dafny.Seq
        d_1082_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
        d_1083_v24_: _dafny.MultiSet
        d_1083_v24_ = _dafny.MultiSet([d_1057_v0_, d_1057_v0_, not(not(d_1057_v0_))])
        d_1084_v25_: _dafny.Set
        d_1084_v25_ = _dafny.Set({d_1057_v0_, (d_1073_v16_)[default__.safeIndex(723, len(d_1073_v16_))]})
        d_1085_v26_: D0
        d_1085_v26_ = D0_DC0(d_1084_v25_)
        pat_let_tv22_ = d_1057_v0_
        pat_let_tv23_ = d_1071_v13_
        pat_let_tv24_ = d_1071_v13_
        pat_let_tv25_ = d_1057_v0_
        pat_let_tv26_ = d_1057_v0_
        index186_ = default__.safeIndex(143, (d_1080_v22_).length(0))
        index187_ = default__.safeIndex(11, (d_1080_v22_).length(0))
        def lambda63_(source12_):
            if source12_.is_DC1:
                d_1086___mcc_h0_ = source12_.cf1
                d_1087___mcc_h1_ = source12_.cf2
                d_1088___mcc_h2_ = source12_.cf3
                d_1089_cf3_ = d_1088___mcc_h2_
                d_1090_cf2_ = d_1087___mcc_h1_
                d_1091_cf1_ = d_1086___mcc_h0_
                if (pat_let_tv22_) in (pat_let_tv23_):
                    return (pat_let_tv24_)[pat_let_tv25_]
                elif True:
                    return pat_let_tv26_
            elif True:
                d_1092___mcc_h3_ = source12_.cf0
                d_1093_cf0_ = d_1092___mcc_h3_
                return False

        rhs233_ = not(((((_dafny.MultiSet(d_1073_v16_)).set(d_1057_v0_, default__.abs(len(d_1082_v23_)))).set(d_1057_v0_, default__.abs(p1))).set(d_1057_v0_, default__.abs(144))) == ((default__.fm22(p1, globalState)).intersection(d_1083_v24_)))
        rhs234_ = lambda63_(d_1085_v26_)
        rhs235_ = (self).f6
        lhs154_ = d_1080_v22_
        lhs155_ = default__.safeIndex(143, (d_1080_v22_).length(0))
        lhs156_ = d_1080_v22_
        lhs157_ = default__.safeIndex(11, (d_1080_v22_).length(0))
        lhs158_ = self
        lhs154_[lhs155_] = rhs233_
        lhs156_[lhs157_] = rhs234_
        lhs158_.f7 = rhs235_
        pat_let_tv27_ = d_1082_v23_
        def lambda64_(source13_):
            if source13_.is_DC1:
                d_1094___mcc_h4_ = source13_.cf1
                d_1095___mcc_h5_ = source13_.cf2
                d_1096___mcc_h6_ = source13_.cf3
                d_1097_cf3_ = d_1096___mcc_h6_
                d_1098_cf2_ = d_1095___mcc_h5_
                d_1099_cf1_ = d_1094___mcc_h4_
                return pat_let_tv27_
            elif True:
                d_1100___mcc_h7_ = source13_.cf0
                d_1101_cf0_ = d_1100___mcc_h7_
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1102_i2_ in range(default__.abs(979))])

        d_1082_v23_ = lambda64_(d_1085_v26_)
        hi10_ = p1
        for d_1103_i3_ in range(default__.safeDivisionInt((self).f6, p1), hi10_):
            d_1104_v27_: _dafny.Map
            d_1104_v27_ = _dafny.Map({(self).f6: False})
            d_1104_v27_ = (d_1104_v27_).set(len((d_1070_v12_ if True else d_1070_v12_)), (d_1080_v22_)[default__.safeIndex(143, (d_1080_v22_).length(0))])
            d_1057_v0_ = False
            d_1057_v0_ = (d_1080_v22_)[default__.safeIndex(143, (d_1080_v22_).length(0))]
            d_1105_v28_: str
            d_1105_v28_ = _dafny.CodePoint('g')
            d_1106_v29_: _dafny.Seq
            d_1106_v29_ = _dafny.SeqWithoutIsStrInference([d_1084_v25_])
            index188_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            index189_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            rhs236_ = not((d_1080_v22_)[default__.safeIndex(143, (d_1080_v22_).length(0))])
            rhs237_ = not((d_1105_v28_) not in (d_1082_v23_))
            rhs238_ = (self).f6
            rhs239_ = default__.safeModuloInt(default__.safeModuloInt((self).f6, self.f7), d_1103_i3_)
            rhs240_ = ((d_1106_v29_)[default__.safeIndex(d_1103_i3_, len(d_1106_v29_))]).ispropersubset((d_1084_v25_) - (d_1084_v25_))
            lhs159_ = d_1080_v22_
            lhs160_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            lhs161_ = globalState
            lhs162_ = globalState
            lhs163_ = d_1080_v22_
            lhs164_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            r0 = rhs236_
            lhs159_[lhs160_] = rhs237_
            lhs161_.f1 = rhs238_
            lhs162_.f1 = rhs239_
            lhs163_[lhs164_] = rhs240_
        hi11_ = p1
        for d_1107_i4_ in range(p1, hi11_):
            d_1108_v30_: str
            d_1108_v30_ = _dafny.CodePoint('c')
            d_1109_v31_: _dafny.Seq
            d_1109_v31_ = _dafny.SeqWithoutIsStrInference([d_1084_v25_, d_1084_v25_])
            d_1110_v32_: T0
            nw193_ = C1()
            nw193_.ctor__(d_1108_v30_, default__.safeModuloInt((self).f6, 95), default__.safeModuloInt(len((d_1109_v31_)[default__.safeIndex((self).f6, len(d_1109_v31_))]), d_1107_i4_))
            d_1110_v32_ = nw193_
            d_1111_v33_: _dafny.Map
            d_1111_v33_ = _dafny.Map({True: d_1110_v32_})
            d_1110_v32_ = ((d_1111_v33_)[(d_1110_v32_.f7) < (p1)] if ((d_1110_v32_.f7) < (p1)) in (d_1111_v33_) else d_1110_v32_)
            index190_ = default__.safeIndex(509, (d_1077_v20_).length(0))
            (d_1077_v20_)[index190_] = (0) - ((_dafny.MultiSet([d_1108_v30_])).cardinality)
            index191_ = default__.safeIndex(509, (d_1077_v20_).length(0))
            index192_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            rhs241_ = d_1057_v0_
            rhs242_ = d_1107_i4_
            rhs243_ = d_1077_v20_
            rhs244_ = self.f7
            rhs245_ = (d_1080_v22_)[default__.safeIndex(143, (d_1080_v22_).length(0))]
            lhs165_ = globalState
            lhs166_ = d_1077_v20_
            lhs167_ = default__.safeIndex(509, (d_1077_v20_).length(0))
            lhs168_ = d_1080_v22_
            lhs169_ = default__.safeIndex(143, (d_1080_v22_).length(0))
            d_1057_v0_ = rhs241_
            lhs165_.f1 = rhs242_
            d_1077_v20_ = rhs243_
            lhs166_[lhs167_] = rhs244_
            lhs168_[lhs169_] = rhs245_
            d_1112_v34_: _dafny.Array
            def lambda65_(d_1113_v23_):
                def lambda66_(d_1114_i5_):
                    return d_1113_v23_

                return lambda66_

            init35_ = lambda65_(d_1082_v23_)
            nw194_ = _dafny.Array(None, 21)
            for i0_35_ in range(nw194_.length(0)):
                nw194_[i0_35_] = init35_(i0_35_)
            d_1112_v34_ = nw194_
            index193_ = default__.safeIndex(561, (d_1112_v34_).length(0))
            (d_1112_v34_)[index193_] = d_1082_v23_
            index194_ = default__.safeIndex(561, (d_1112_v34_).length(0))
            (d_1112_v34_)[index194_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xr"))) + (d_1082_v23_)) + ((d_1082_v23_) + (d_1082_v23_))
            d_1115_v35_: C2
            nw195_ = C2()
            nw195_.ctor__(default__.safeModuloInt(256, (d_1110_v32_).f6), p1)
            d_1115_v35_ = nw195_
        index195_ = default__.safeIndex(821, (d_1077_v20_).length(0))
        (d_1077_v20_)[index195_] = (self).f6
        index196_ = default__.safeIndex(821, (d_1077_v20_).length(0))
        (d_1077_v20_)[index196_] = p1
        d_1116_v36_: _dafny.Map
        d_1116_v36_ = _dafny.Map({-348: (self).f6})
        d_1117_v37_: _dafny.Map
        d_1117_v37_ = _dafny.Map({d_1116_v36_: (d_1080_v22_)[default__.safeIndex(143, (d_1080_v22_).length(0))]})
        r0 = ((d_1117_v37_)[d_1116_v36_] if (d_1116_v36_) in (d_1117_v37_) else default__.fm0(d_1057_v0_, d_1116_v36_, p1, globalState))
        return r0

    def m7(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        d_1118_v0_: _dafny.Array
        nw196_ = _dafny.Array(False, 5)
        d_1118_v0_ = nw196_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1118_v0_).length(0)):
            d_1119_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1119_i0_)) and ((d_1119_i0_) < ((d_1118_v0_).length(0)))):
                (d_1118_v0_)[(d_1119_i0_)] = p1
        d_1120_i1_: int
        d_1120_i1_ = 0
        with _dafny.label("5"):
            while p1:
                with _dafny.c_label("5"):
                    if (d_1120_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1120_i1_ = (d_1120_i1_) + (1)
                    d_1121_v1_: bool
                    d_1121_v1_ = False
                    d_1122_v2_: _dafny.Seq
                    d_1122_v2_ = _dafny.SeqWithoutIsStrInference([d_1121_v1_])
                    rhs246_ = d_1118_v0_
                    rhs247_ = p0
                    rhs248_ = _dafny.SeqWithoutIsStrInference([not (p0) or (True), not(((self).f6) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgkqtd"))))), p0, p1])
                    rhs249_ = ((p2) + (p2)) == ((self).f6)
                    d_1118_v0_ = rhs246_
                    d_1121_v1_ = rhs247_
                    d_1122_v2_ = rhs248_
                    d_1121_v1_ = rhs249_
                    d_1123_v3_: _dafny.Array
                    def lambda67_(d_1124_i2_):
                        return default__.safeModuloInt(d_1124_i2_, (self).f6)

                    init36_ = lambda67_
                    nw197_ = _dafny.Array(None, 29)
                    for i0_36_ in range(nw197_.length(0)):
                        nw197_[i0_36_] = init36_(i0_36_)
                    d_1123_v3_ = nw197_
                    d_1125_v4_: _dafny.Seq
                    d_1125_v4_ = _dafny.SeqWithoutIsStrInference([d_1123_v3_])
                    d_1126_v5_: _dafny.Map
                    d_1126_v5_ = _dafny.Map({(d_1125_v4_)[default__.safeIndex(-925, len(d_1125_v4_))]: d_1121_v1_})
                    d_1126_v5_ = _dafny.Map({d_1123_v3_: d_1121_v1_})
                    d_1127_v6_: D0
                    d_1127_v6_ = D0_DC1(True, d_1118_v0_, d_1121_v1_)
                    d_1128_v7_: _dafny.Map
                    d_1128_v7_ = _dafny.Map({273: d_1121_v1_})
                    d_1129_v8_: int
                    d_1130_v9_: bool
                    d_1131_v10_: _dafny.Array
                    out20_: int
                    out21_: bool
                    out22_: _dafny.Array
                    out20_, out21_, out22_ = default__.m0(d_1127_v6_, default__.fm1(not(((d_1128_v7_)[(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blo")))))] if ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "blo")))))) in (d_1128_v7_) else False)), globalState), globalState)
                    d_1129_v8_ = out20_
                    d_1130_v9_ = out21_
                    d_1131_v10_ = out22_
                    d_1132_v11_: _dafny.MultiSet
                    d_1132_v11_ = _dafny.MultiSet([p0])
                    d_1133_v12_: _dafny.Seq
                    d_1133_v12_ = _dafny.SeqWithoutIsStrInference([(d_1132_v11_).cardinality, d_1129_v8_])
                    d_1134_v13_: _dafny.Map
                    d_1134_v13_ = _dafny.Map({(d_1133_v12_)[default__.safeIndex((self).f6, len(d_1133_v12_))]: 144})
                    d_1135_v14_: _dafny.Seq
                    d_1135_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aywp"))
                    if default__.fm0(d_1121_v1_, d_1134_v13_, len(d_1135_v14_), globalState):
                        d_1136_v15_: _dafny.Set
                        d_1136_v15_ = _dafny.Set({(self).f6, -188})
                        d_1137_v16_: C1
                        nw198_ = C1()
                        nw198_.ctor__(_dafny.CodePoint('s'), len(d_1136_v15_), len(default__.fm36(d_1129_v8_, True, (self).f6, d_1130_v9_, globalState)))
                        d_1137_v16_ = nw198_
                        (globalState).f1 = len((_dafny.SeqWithoutIsStrInference([d_1135_v14_])).set(default__.safeIndex(self.f7, len(_dafny.SeqWithoutIsStrInference([d_1135_v14_]))), d_1135_v14_))
                        d_1138_v17_: D9
                        d_1138_v17_ = D9_DC22((_dafny.SeqWithoutIsStrInference([d_1129_v8_ for d_1139_i3_ in range(default__.abs(-814))])) + (d_1133_v12_))
                        d_1138_v17_ = d_1138_v17_
                        d_1140_v18_: C1
                        nw199_ = C1()
                        nw199_.ctor__(default__.fm25(globalState), default__.fm1(d_1121_v1_, globalState), (self).f6)
                        d_1140_v18_ = nw199_
                        d_1135_v14_ = _dafny.SeqWithoutIsStrInference([(d_1137_v16_).f13 for d_1141_i4_ in range(default__.abs(440))])
                    elif True:
                        d_1142_v19_: str
                        d_1142_v19_ = _dafny.CodePoint('i')
                        d_1143_v20_: C1
                        nw200_ = C1()
                        nw200_.ctor__(d_1142_v19_, d_1129_v8_, ((self).f6) + (p2))
                        d_1143_v20_ = nw200_
                        (globalState).f1 = (self).f6
                        index197_ = default__.safeIndex(559, (d_1118_v0_).length(0))
                        (d_1118_v0_)[index197_] = (d_1142_v19_) in (d_1135_v14_)
                        index198_ = default__.safeIndex(795, (d_1123_v3_).length(0))
                        (d_1123_v3_)[index198_] = self.f7
                        index199_ = default__.safeIndex(559, (d_1118_v0_).length(0))
                        index200_ = default__.safeIndex(795, (d_1123_v3_).length(0))
                        rhs250_ = d_1118_v0_
                        rhs251_ = d_1118_v0_
                        rhs252_ = p0
                        rhs253_ = len(d_1135_v14_)
                        rhs254_ = -972
                        lhs170_ = d_1118_v0_
                        lhs171_ = default__.safeIndex(559, (d_1118_v0_).length(0))
                        lhs172_ = d_1123_v3_
                        lhs173_ = default__.safeIndex(795, (d_1123_v3_).length(0))
                        d_1118_v0_ = rhs250_
                        d_1118_v0_ = rhs251_
                        lhs170_[lhs171_] = rhs252_
                        r0 = rhs253_
                        lhs172_[lhs173_] = rhs254_
                        d_1144_v21_: _dafny.Map
                        d_1144_v21_ = _dafny.Map({((d_1122_v2_) + (d_1122_v2_)).set(default__.safeIndex((self).f6, len((d_1122_v2_) + (d_1122_v2_))), p0): (d_1121_v1_ if p1 else d_1121_v1_)})
                        d_1144_v21_ = (d_1144_v21_).set(default__.fm21(globalState), not((d_1118_v0_)[default__.safeIndex(559, (d_1118_v0_).length(0))]))
                        d_1145_v22_: _dafny.Set
                        d_1145_v22_ = _dafny.Set({p0})
                        d_1146_v23_: _dafny.Map
                        d_1146_v23_ = _dafny.Map({p1: d_1129_v8_})
                        d_1147_v25_: _dafny.MultiSet
                        d_1147_v25_ = _dafny.MultiSet([(self).f6])
                        d_1148_v28_: _dafny.Array
                        nw201_ = _dafny.Array(None, 17)
                        nw201_[int(0)] = (default__.fm1(p0, globalState)) - ((self).f6)
                        nw201_[int(1)] = ((d_1123_v3_)[default__.safeIndex(795, (d_1123_v3_).length(0))]) - (d_1129_v8_)
                        nw201_[int(2)] = default__.fm1(p1, globalState)
                        nw201_[int(3)] = d_1129_v8_
                        nw201_[int(4)] = (len(d_1145_v22_) if p0 else len(d_1146_v23_))
                        nw201_[int(5)] = default__.safeModuloInt(-436, p2)
                        nw201_[int(6)] = d_1129_v8_
                        nw201_[int(7)] = ((d_1123_v3_)[default__.safeIndex(795, (d_1123_v3_).length(0))]) + ((_dafny.MultiSet([d_1129_v8_])).cardinality)
                        nw201_[int(8)] = default__.fm1(d_1130_v9_, globalState)
                        def iife85_():
                            coll35_ = _dafny.Set()
                            compr_35_: int
                            for compr_35_ in _dafny.IntegerRange(219, 988):
                                d_1149_v24_: int = compr_35_
                                if ((219) <= (d_1149_v24_)) and ((d_1149_v24_) < (988)):
                                    coll35_ = coll35_.union(_dafny.Set([(d_1149_v24_) * (p2)]))
                            return _dafny.Set(coll35_)
                        def iife86_():
                            coll36_ = _dafny.Set()
                            compr_36_: int
                            for compr_36_ in (d_1147_v25_).Elements:
                                d_1150_v26_: int = compr_36_
                                if (d_1150_v26_) in (d_1147_v25_):
                                    coll36_ = coll36_.union(_dafny.Set([(d_1150_v26_) + (803)]))
                            return _dafny.Set(coll36_)
                        nw201_[int(9)] = len((iife85_()
                        ) | (iife86_()
                        ))
                        def iife87_():
                            coll37_ = _dafny.Map()
                            compr_37_: int
                            for compr_37_ in _dafny.IntegerRange(785, 192):
                                d_1151_v27_: int = compr_37_
                                if ((785) <= (d_1151_v27_)) and ((d_1151_v27_) < (192)):
                                    coll37_[default__.safeModuloInt(d_1151_v27_, d_1129_v8_)] = p1
                            return _dafny.Map(coll37_)
                        nw201_[int(10)] = ((d_1123_v3_)[default__.safeIndex(795, (d_1123_v3_).length(0))]) + (len(iife87_()
                        ))
                        nw201_[int(11)] = p2
                        nw201_[int(12)] = p2
                        nw201_[int(13)] = p2
                        nw201_[int(14)] = (self).fm15(len((_dafny.SeqWithoutIsStrInference([d_1133_v12_ for d_1152_i5_ in range(default__.abs(543))])).set(default__.safeIndex(77, len(_dafny.SeqWithoutIsStrInference([d_1133_v12_ for d_1153_i5_ in range(default__.abs(543))]))), d_1133_v12_)), self.f7, globalState)
                        nw201_[int(15)] = (270) * (p2)
                        nw201_[int(16)] = (self).f6
                        d_1148_v28_ = nw201_
                        d_1148_v28_ = d_1148_v28_
                    pass
            pass
        d_1154_v29_: _dafny.MultiSet
        d_1154_v29_ = _dafny.MultiSet([p2])
        d_1155_v30_: _dafny.Seq
        d_1155_v30_ = _dafny.SeqWithoutIsStrInference([default__.fm35(self.f7, (d_1154_v29_).cardinality, globalState)])
        d_1156_v31_: _dafny.Seq
        d_1156_v31_ = _dafny.SeqWithoutIsStrInference([len((d_1155_v30_)[default__.safeIndex(p2, len(d_1155_v30_))])])
        d_1157_v32_: _dafny.Seq
        d_1157_v32_ = _dafny.SeqWithoutIsStrInference([(d_1156_v31_)[default__.safeIndex(p2, len(d_1156_v31_))], p2])
        d_1158_v33_: _dafny.Seq
        d_1158_v33_ = _dafny.SeqWithoutIsStrInference([len(d_1157_v32_)])
        r1 = (0) - ((d_1156_v31_)[default__.safeIndex(p2, len(d_1156_v31_))])
        d_1159_v34_: str
        d_1159_v34_ = _dafny.CodePoint('e')
        d_1160_v35_: C1
        nw202_ = C1()
        nw202_.ctor__(d_1159_v34_, p2, len(_dafny.SeqWithoutIsStrInference([(D10_DC24(_dafny.CodePoint('e'))).cf51 for d_1161_i6_ in range(default__.abs(-485))])))
        d_1160_v35_ = nw202_
        rhs255_ = d_1159_v34_
        rhs256_ = d_1160_v35_
        d_1159_v34_ = rhs255_
        d_1160_v35_ = rhs256_
        d_1162_v36_: _dafny.Map
        d_1162_v36_ = _dafny.Map({(p2) + (p2): d_1160_v35_})
        d_1162_v36_ = (d_1162_v36_).set(p2, d_1160_v35_)
        d_1163_v37_: bool
        d_1163_v37_ = True
        d_1164_v38_: D10
        d_1164_v38_ = D10_DC26(d_1163_v37_)
        d_1163_v37_ = (d_1164_v38_).cf52
        d_1165_v39_: _dafny.Seq
        d_1165_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vdmx"))
        r0 = (len(d_1165_v39_)) * (p2)
        r1 = ((self).f6) + (default__.safeDivisionInt(p2, p2))
        r2 = p2
        return r0, r1, r2


class C4(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f6: int = int(0)
        self.f12: _dafny.Map = _dafny.Map({})
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f11, f12, f6, f7):
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f6 = f6
        (self).f7 = f7

    def fm11(self, p0, globalState):
        return D0_DC0((_dafny.Set({(self).f11})).intersection(_dafny.Set({(self).f11, False})))

    def fm12(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t'), _dafny.CodePoint('g')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1166_i0_ in range(default__.abs(-535))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1167_i1_ in range(default__.abs(511))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('u')])))

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        hi12_ = default__.safeDivisionInt(default__.fm1(p0, globalState), p1)
        for d_1168_i0_ in range((self).f6, hi12_):
            d_1169_v0_: _dafny.MultiSet
            d_1169_v0_ = _dafny.MultiSet([d_1168_i0_])
            d_1170_v1_: _dafny.Seq
            d_1170_v1_ = _dafny.SeqWithoutIsStrInference([77])
            d_1169_v0_ = _dafny.MultiSet(d_1170_v1_)
            d_1171_v2_: _dafny.Map
            d_1171_v2_ = _dafny.Map({d_1168_i0_: default__.safeDivisionInt(p2, 76)})
            (self).f7 = ((d_1171_v2_)[(0) - (default__.safeModuloInt((d_1170_v1_)[default__.safeIndex((0) - (d_1168_i0_), len(d_1170_v1_))], (self).f6))] if ((0) - (default__.safeModuloInt((d_1170_v1_)[default__.safeIndex((0) - (d_1168_i0_), len(d_1170_v1_))], (self).f6))) in (d_1171_v2_) else self.f7)
            d_1172_v3_: _dafny.Seq
            d_1172_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            d_1172_v3_ = d_1172_v3_
            (self).f7 = (p1) - (228)
        d_1173_v4_: _dafny.Seq
        d_1173_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwyinhv"))
        d_1174_v5_: _dafny.Array
        def lambda68_(d_1175_i1_):
            return (d_1175_i1_) * ((self).f6)

        init37_ = lambda68_
        nw203_ = _dafny.Array(None, 27)
        for i0_37_ in range(nw203_.length(0)):
            nw203_[i0_37_] = init37_(i0_37_)
        d_1174_v5_ = nw203_
        d_1176_v6_: _dafny.Seq
        d_1176_v6_ = _dafny.SeqWithoutIsStrInference([d_1174_v5_, d_1174_v5_])
        d_1177_v7_: _dafny.Array
        nw204_ = _dafny.Array(None, 9)
        nw204_[int(0)] = p2
        nw204_[int(1)] = p2
        nw204_[int(2)] = (0) - (p2)
        nw204_[int(3)] = len(d_1173_v4_)
        nw204_[int(4)] = p1
        nw204_[int(5)] = p1
        nw204_[int(6)] = self.f7
        nw204_[int(7)] = len(d_1176_v6_)
        nw204_[int(8)] = self.f7
        d_1177_v7_ = nw204_
        d_1178_v8_: D2
        d_1178_v8_ = D2_DC8(D2_DC5(d_1174_v5_))
        d_1178_v8_ = d_1178_v8_
        d_1179_v9_: C0
        nw205_ = C0()
        nw205_.ctor__(False)
        d_1179_v9_ = nw205_
        d_1180_v10_: _dafny.MultiSet
        d_1180_v10_ = _dafny.MultiSet([p0, d_1179_v9_.f10])
        d_1181_v11_: _dafny.Map
        d_1181_v11_ = _dafny.Map({_dafny.CodePoint('v'): d_1180_v10_})
        d_1182_v12_: str
        d_1182_v12_ = _dafny.CodePoint('n')
        d_1183_v13_: _dafny.MultiSet
        d_1183_v13_ = (((d_1181_v11_)[d_1182_v12_] if (d_1182_v12_) in (d_1181_v11_) else _dafny.MultiSet([d_1179_v9_.f10]))) | (d_1180_v10_)
        source14_ = d_1183_v13_
        d_1184___mcc_h0_ = source14_
        d_1185_cf18_ = d_1184___mcc_h0_
        d_1186_v14_: _dafny.Array
        nw206_ = _dafny.Array(False, 1)
        d_1186_v14_ = nw206_
        index201_ = default__.safeIndex(130, (d_1186_v14_).length(0))
        (d_1186_v14_)[index201_] = d_1179_v9_.f10
        index202_ = default__.safeIndex(130, (d_1186_v14_).length(0))
        (d_1186_v14_)[index202_] = True
        index203_ = default__.safeIndex(130, (d_1186_v14_).length(0))
        (d_1186_v14_)[index203_] = (self).f11
        d_1187_v15_: _dafny.MultiSet
        d_1187_v15_ = _dafny.MultiSet([default__.fm1(p0, globalState)])
        d_1188_v16_: _dafny.Map
        d_1188_v16_ = _dafny.Map({d_1187_v15_: p0})
        d_1189_v17_: _dafny.Map
        d_1189_v17_ = _dafny.Map({((d_1188_v16_)[(d_1187_v15_).set(520, default__.abs((self).f6))] if ((d_1187_v15_).set(520, default__.abs((self).f6))) in (d_1188_v16_) else d_1179_v9_.f10): 266})
        d_1189_v17_ = (d_1189_v17_) | ((d_1189_v17_) | (d_1189_v17_))
        d_1190_v18_: _dafny.Map
        d_1190_v18_ = _dafny.Map({self.f7: d_1179_v9_.f10})
        d_1191_v19_: _dafny.Map
        d_1191_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbfloex")): d_1182_v12_})
        d_1190_v18_ = (d_1190_v18_).set(len(d_1191_v19_), d_1179_v9_.f10)
        d_1192_v20_: _dafny.Map
        d_1192_v20_ = _dafny.Map({(0) - (p1): not((self).f11)})
        d_1193_v21_: D1
        d_1193_v21_ = D1_DC4(p0)
        d_1194_v22_: D2
        d_1194_v22_ = D2_DC7(p0, 79, d_1192_v20_, d_1193_v21_)
        source15_ = d_1194_v22_
        if source15_.is_DC6:
            d_1195___mcc_h1_ = source15_.cf10
            d_1196___mcc_h2_ = source15_.cf11
            d_1197___mcc_h3_ = source15_.cf12
            d_1198_cf12_ = d_1197___mcc_h3_
            d_1199_cf11_ = d_1196___mcc_h2_
            d_1200_cf10_ = d_1195___mcc_h1_
            d_1201_v23_: _dafny.Seq
            d_1201_v23_ = _dafny.SeqWithoutIsStrInference([d_1179_v9_.f10, d_1179_v9_.f10])
            r0 = not(((default__.fm13(globalState)).set(d_1179_v9_.f10, default__.abs(d_1198_cf12_))).issubset((default__.fm13(globalState)).set(p0, default__.abs((0) - (len(d_1201_v23_))))))
            r1 = (self).f11
            d_1202_v24_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.CodePoint('D'), 27)
            d_1202_v24_ = nw207_
            d_1202_v24_ = d_1202_v24_
            d_1203_v25_: C0
            nw208_ = C0()
            nw208_.ctor__(p0)
            d_1203_v25_ = nw208_
        elif source15_.is_DC7:
            d_1204___mcc_h4_ = source15_.cf13
            d_1205___mcc_h5_ = source15_.cf14
            d_1206___mcc_h6_ = source15_.cf15
            d_1207___mcc_h7_ = source15_.cf16
            d_1208_cf16_ = d_1207___mcc_h7_
            d_1209_cf15_ = d_1206___mcc_h6_
            d_1210_cf14_ = d_1205___mcc_h5_
            d_1211_cf13_ = d_1204___mcc_h4_
            d_1212_v26_: _dafny.Array
            def lambda69_(d_1213_i2_):
                return (False) and (False)

            init38_ = lambda69_
            nw209_ = _dafny.Array(None, 16)
            for i0_38_ in range(nw209_.length(0)):
                nw209_[i0_38_] = init38_(i0_38_)
            d_1212_v26_ = nw209_
            index204_ = default__.safeIndex(767, (d_1212_v26_).length(0))
            (d_1212_v26_)[index204_] = d_1179_v9_.f10
            index205_ = default__.safeIndex(767, (d_1212_v26_).length(0))
            (d_1212_v26_)[index205_] = d_1179_v9_.f10
            (d_1179_v9_).f10 = d_1211_cf13_
            d_1173_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvthhe"))) + (_dafny.SeqWithoutIsStrInference([d_1182_v12_ for d_1214_i3_ in range(default__.abs(504))]))
            d_1215_v27_: D0
            d_1215_v27_ = D0_DC0(_dafny.Set({(self).f11, (self).f11}))
            d_1216_v28_: _dafny.Map
            d_1216_v28_ = _dafny.Map({d_1215_v27_: d_1210_cf14_})
            d_1216_v28_ = (d_1216_v28_).set((self).fm11(d_1182_v12_, globalState), p1)
        elif source15_.is_DC5:
            d_1217___mcc_h8_ = source15_.cf9
            d_1218_cf9_ = d_1217___mcc_h8_
            (globalState).f1 = default__.fm1(default__.fm0(d_1179_v9_.f10, _dafny.Map({self.f7: p2}), (0) - (p1), globalState), globalState)
            if d_1179_v9_.f10:
                d_1219_v29_: _dafny.Map
                d_1219_v29_ = _dafny.Map({639: p2})
                d_1219_v29_ = (d_1219_v29_).set(((self).f6) - (750), len(d_1173_v4_))
                d_1220_v30_: D2
                d_1220_v30_ = D2_DC5(d_1174_v5_)
                d_1221_v31_: _dafny.Map
                d_1221_v31_ = _dafny.Map({d_1192_v20_: 243})
                rhs257_ = ((0) - (p2) if p0 else len(d_1221_v31_))
                rhs258_ = d_1220_v30_
                lhs174_ = globalState
                lhs174_.f1 = rhs257_
                d_1220_v30_ = rhs258_
                (d_1179_v9_).f10 = True
                d_1222_v32_: _dafny.Array
                nw210_ = _dafny.Array(False, 8)
                d_1222_v32_ = nw210_
                d_1223_v33_: D0
                d_1223_v33_ = D0_DC1(False, d_1222_v32_, (self).f11)
                d_1224_v34_: int
                d_1225_v35_: bool
                d_1226_v36_: _dafny.Array
                out23_: int
                out24_: bool
                out25_: _dafny.Array
                out23_, out24_, out25_ = default__.m0(d_1223_v33_, (self).f6, globalState)
                d_1224_v34_ = out23_
                d_1225_v35_ = out24_
                d_1226_v36_ = out25_
                d_1219_v29_ = (d_1219_v29_).set(self.f7, (self).f6)
            elif True:
                (globalState).f1 = (self).f6
                d_1227_v37_: _dafny.Array
                nw211_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_1227_v37_ = nw211_
                d_1227_v37_ = d_1227_v37_
                (d_1179_v9_).f10 = p0
                d_1192_v20_ = (d_1192_v20_).set(default__.safeDivisionInt(-709, self.f7), (self).f11)
                d_1228_v38_: _dafny.MultiSet
                d_1228_v38_ = _dafny.MultiSet([p2])
                d_1229_v39_: _dafny.Seq
                d_1229_v39_ = _dafny.SeqWithoutIsStrInference([d_1228_v38_])
                (globalState).f1 = len(d_1229_v39_)
            d_1174_v5_ = d_1218_cf9_
            d_1230_v40_: _dafny.Map
            d_1230_v40_ = _dafny.Map({d_1179_v9_.f10: (self).f11})
            index206_ = default__.safeIndex(844, (d_1218_cf9_).length(0))
            (d_1218_cf9_)[index206_] = len((d_1230_v40_) | (d_1230_v40_))
            index207_ = default__.safeIndex(844, (d_1218_cf9_).length(0))
            (d_1218_cf9_)[index207_] = ((self).f6 if False else p1)
        elif True:
            d_1231___mcc_h9_ = source15_.cf17
            d_1232_cf17_ = d_1231___mcc_h9_
            r1 = d_1179_v9_.f10
            index208_ = default__.safeIndex(132, (d_1174_v5_).length(0))
            (d_1174_v5_)[index208_] = p2
            d_1233_v41_: _dafny.Array
            def lambda70_(d_1234_v9_):
                def lambda71_(d_1235_i4_):
                    return d_1234_v9_.f10

                return lambda71_

            init39_ = lambda70_(d_1179_v9_)
            nw212_ = _dafny.Array(None, 7)
            for i0_39_ in range(nw212_.length(0)):
                nw212_[i0_39_] = init39_(i0_39_)
            d_1233_v41_ = nw212_
            index209_ = default__.safeIndex(158, (d_1233_v41_).length(0))
            (d_1233_v41_)[index209_] = not(p0)
            index210_ = default__.safeIndex(371, (d_1177_v7_).length(0))
            (d_1177_v7_)[index210_] = self.f7
            index211_ = default__.safeIndex(132, (d_1174_v5_).length(0))
            index212_ = default__.safeIndex(158, (d_1233_v41_).length(0))
            index213_ = default__.safeIndex(371, (d_1177_v7_).length(0))
            rhs259_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdoyxqucy"))) < (d_1173_v4_)
            rhs260_ = p2
            rhs261_ = p0
            rhs262_ = -836
            rhs263_ = (self).f11
            lhs175_ = d_1179_v9_
            lhs176_ = d_1174_v5_
            lhs177_ = default__.safeIndex(132, (d_1174_v5_).length(0))
            lhs178_ = d_1233_v41_
            lhs179_ = default__.safeIndex(158, (d_1233_v41_).length(0))
            lhs180_ = d_1177_v7_
            lhs181_ = default__.safeIndex(371, (d_1177_v7_).length(0))
            lhs175_.f10 = rhs259_
            lhs176_[lhs177_] = rhs260_
            lhs178_[lhs179_] = rhs261_
            lhs180_[lhs181_] = rhs262_
            r1 = rhs263_
            d_1236_v42_: _dafny.Array
            nw213_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_1236_v42_ = nw213_
            d_1236_v42_ = d_1236_v42_
            d_1237_v43_: _dafny.Seq
            d_1237_v43_ = _dafny.SeqWithoutIsStrInference([d_1233_v41_])
            d_1238_v44_: _dafny.MultiSet
            d_1238_v44_ = _dafny.MultiSet([(d_1237_v43_)[default__.safeIndex(self.f7, len(d_1237_v43_))]])
            d_1238_v44_ = (d_1238_v44_).set(d_1233_v41_, default__.abs(p1))
        d_1239_i5_: int
        d_1239_i5_ = 0
        with _dafny.label("6"):
            while (default__.safeDivisionInt(536, len(_dafny.Map({d_1179_v9_.f10: (self).f11})))) == (self.f7):
                with _dafny.c_label("6"):
                    if (d_1239_i5_) >= (100):
                        raise _dafny.Break("6")
                    d_1239_i5_ = (d_1239_i5_) + (1)
                    d_1182_v12_ = default__.fm14(194, len(d_1173_v4_), 767, globalState)
                    (d_1179_v9_).f10 = p0
                    d_1240_v45_: _dafny.MultiSet
                    d_1240_v45_ = _dafny.MultiSet([(d_1173_v4_) + (d_1173_v4_), (d_1173_v4_) + (d_1173_v4_), (d_1173_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uohlhbu")))])
                    (globalState).f1 = (d_1240_v45_).cardinality
                    d_1182_v12_ = d_1182_v12_
                    pass
            pass
        d_1241_v46_: _dafny.Seq
        d_1241_v46_ = _dafny.SeqWithoutIsStrInference([p1])
        r0 = (p2) in (d_1241_v46_)
        r1 = (d_1173_v4_) == (d_1173_v4_)
        r2 = self.f7
        return r0, r1, r2

    def m4(self, globalState):
        d_1242_v0_: _dafny.Map
        d_1242_v0_ = _dafny.Map({(self).f11: (False) == ((self).f11)})
        d_1243_v1_: _dafny.MultiSet
        d_1243_v1_ = _dafny.MultiSet([(self).f11])
        d_1242_v0_ = (d_1242_v0_).set((d_1243_v1_).isdisjoint(d_1243_v1_), True)
        d_1244_v2_: bool
        d_1244_v2_ = True
        d_1244_v2_ = (self).f11
        hi13_ = 110
        for d_1245_i0_ in range(self.f7, hi13_):
            if d_1244_v2_:
                d_1246_v3_: _dafny.MultiSet
                d_1246_v3_ = d_1243_v1_
                d_1247_v4_: _dafny.Map
                d_1247_v4_ = _dafny.Map({(self).f11: d_1246_v3_})
                d_1247_v4_ = (d_1247_v4_).set(not ((self).f11) or ((self).f11), d_1246_v3_)
                d_1248_v5_: _dafny.Set
                d_1248_v5_ = _dafny.Set({d_1244_v2_, (self).f11, d_1244_v2_})
                d_1249_v6_: D8
                d_1249_v6_ = D8_DC20((self).f11, (self).f11, d_1244_v2_, d_1244_v2_)
                d_1250_v7_: _dafny.Set
                d_1250_v7_ = _dafny.Set({default__.fm1((d_1249_v6_).cf39, globalState)})
                d_1251_v8_: T0
                nw214_ = C2()
                nw214_.ctor__(len(d_1248_v5_), len(d_1250_v7_))
                d_1251_v8_ = nw214_
                d_1252_v9_: _dafny.Seq
                d_1252_v9_ = _dafny.SeqWithoutIsStrInference([d_1244_v2_, (self).f11, False, d_1244_v2_, (self).f11])
                d_1253_v10_: D8
                d_1253_v10_ = D8_DC18(d_1251_v8_)
                d_1254_v11_: _dafny.Array
                nw215_ = _dafny.Array(None, 29)
                nw215_[int(0)] = d_1251_v8_
                nw215_[int(1)] = d_1251_v8_
                nw215_[int(2)] = d_1251_v8_
                nw215_[int(3)] = d_1251_v8_
                nw215_[int(4)] = d_1251_v8_
                nw215_[int(5)] = d_1251_v8_
                nw215_[int(6)] = d_1251_v8_
                nw215_[int(7)] = d_1251_v8_
                nw215_[int(8)] = d_1251_v8_
                nw215_[int(9)] = (d_1251_v8_ if (self).f11 else d_1251_v8_)
                nw215_[int(10)] = d_1251_v8_
                nw215_[int(11)] = d_1251_v8_
                nw215_[int(12)] = (d_1251_v8_ if (self).f11 else d_1251_v8_)
                nw215_[int(13)] = d_1251_v8_
                nw215_[int(14)] = d_1251_v8_
                nw215_[int(15)] = d_1251_v8_
                nw215_[int(16)] = d_1251_v8_
                nw215_[int(17)] = d_1251_v8_
                nw215_[int(18)] = d_1251_v8_
                nw215_[int(19)] = ((d_1253_v10_).cf33 if not((d_1252_v9_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1244_v2_, d_1244_v2_, d_1244_v2_])), len(d_1252_v9_))]) else d_1251_v8_)
                nw215_[int(20)] = d_1251_v8_
                nw215_[int(21)] = d_1251_v8_
                nw215_[int(22)] = d_1251_v8_
                nw215_[int(23)] = d_1251_v8_
                nw215_[int(24)] = d_1251_v8_
                nw215_[int(25)] = d_1251_v8_
                nw215_[int(26)] = d_1251_v8_
                nw215_[int(27)] = d_1251_v8_
                nw215_[int(28)] = d_1251_v8_
                d_1254_v11_ = nw215_
                index214_ = default__.safeIndex(661, (d_1254_v11_).length(0))
                (d_1254_v11_)[index214_] = d_1251_v8_
                index215_ = default__.safeIndex(661, (d_1254_v11_).length(0))
                (d_1254_v11_)[index215_] = (D8_DC18(d_1251_v8_)).cf33
                d_1255_v12_: _dafny.Seq
                d_1255_v12_ = _dafny.SeqWithoutIsStrInference([d_1251_v8_.f7, 598, (d_1245_i0_) + (d_1245_i0_)])
                (self).f7 = (0) - ((d_1255_v12_)[default__.safeIndex(default__.fm1(d_1244_v2_, globalState), len(d_1255_v12_))])
                d_1256_v13_: _dafny.Array
                def lambda72_(d_1257_i1_):
                    return (d_1257_i1_) - ((self).f6)

                init40_ = lambda72_
                nw216_ = _dafny.Array(None, 24)
                for i0_40_ in range(nw216_.length(0)):
                    nw216_[i0_40_] = init40_(i0_40_)
                d_1256_v13_ = nw216_
                index216_ = default__.safeIndex(819, (d_1256_v13_).length(0))
                (d_1256_v13_)[index216_] = (276 if d_1244_v2_ else d_1251_v8_.f7)
                index217_ = default__.safeIndex(819, (d_1256_v13_).length(0))
                (d_1256_v13_)[index217_] = (d_1251_v8_).f6
                d_1258_v14_: _dafny.Map
                d_1258_v14_ = _dafny.Map({(d_1254_v11_)[default__.safeIndex(661, (d_1254_v11_).length(0))]: d_1244_v2_})
                d_1258_v14_ = (d_1258_v14_).set((d_1254_v11_)[default__.safeIndex(661, (d_1254_v11_).length(0))], d_1244_v2_)
            elif True:
                d_1259_v15_: _dafny.Seq
                d_1259_v15_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_1260_v16_: _dafny.Map
                d_1260_v16_ = _dafny.Map({default__.fm1((self).f11, globalState): len(d_1259_v15_)})
                d_1261_v17_: _dafny.Map
                d_1261_v17_ = _dafny.Map({d_1260_v16_: (d_1245_i0_ if (self).f11 else d_1245_i0_)})
                d_1261_v17_ = (d_1261_v17_).set(d_1260_v16_, (self.f7) + (self.f7))
                d_1262_v18_: str
                d_1262_v18_ = _dafny.CodePoint('e')
                d_1263_v19_: _dafny.Array
                def lambda73_(d_1264_v2_):
                    def lambda74_(d_1265_i2_):
                        return d_1264_v2_

                    return lambda74_

                init41_ = lambda73_(d_1244_v2_)
                nw217_ = _dafny.Array(None, 29)
                for i0_41_ in range(nw217_.length(0)):
                    nw217_[i0_41_] = init41_(i0_41_)
                d_1263_v19_ = nw217_
                index218_ = default__.safeIndex(426, (d_1263_v19_).length(0))
                (d_1263_v19_)[index218_] = d_1244_v2_
                d_1266_v20_: _dafny.Array
                def lambda75_(d_1267_i3_):
                    return (d_1267_i3_) * (-171)

                init42_ = lambda75_
                nw218_ = _dafny.Array(None, 24)
                for i0_42_ in range(nw218_.length(0)):
                    nw218_[i0_42_] = init42_(i0_42_)
                d_1266_v20_ = nw218_
                index219_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                (d_1266_v20_)[index219_] = default__.safeDivisionInt((0) - ((self).f6), (self).f6)
                d_1268_v21_: _dafny.Seq
                d_1268_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxkhh"))
                d_1269_v22_: _dafny.Map
                d_1269_v22_ = _dafny.Map({d_1244_v2_: d_1268_v21_})
                index220_ = default__.safeIndex(426, (d_1263_v19_).length(0))
                index221_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                def iife88_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(63, 280):
                        d_1270_v23_: int = compr_38_
                        if ((63) <= (d_1270_v23_)) and ((d_1270_v23_) < (280)):
                            coll38_[(d_1270_v23_) - (self.f7)] = 689
                    return _dafny.Map(coll38_)
                rhs264_ = d_1262_v18_
                rhs265_ = (not(d_1244_v2_)) or ((d_1244_v2_) in (d_1243_v1_))
                rhs266_ = self.f7
                rhs267_ = (0) - (default__.safeModuloInt((0) - (-360), d_1245_i0_))
                rhs268_ = ((d_1269_v22_) | (d_1269_v22_) if default__.fm0((self).f11, iife88_()
                , self.f7, globalState) else _dafny.Map({not((self).f11): (self).fm12(False, (self).f11, globalState)}))
                lhs182_ = d_1263_v19_
                lhs183_ = default__.safeIndex(426, (d_1263_v19_).length(0))
                lhs184_ = d_1266_v20_
                lhs185_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                lhs186_ = globalState
                d_1262_v18_ = rhs264_
                lhs182_[lhs183_] = rhs265_
                lhs184_[lhs185_] = rhs266_
                lhs186_.f1 = rhs267_
                d_1269_v22_ = rhs268_
                d_1268_v21_ = (d_1268_v21_) + (d_1268_v21_)
                d_1271_v24_: _dafny.Map
                d_1271_v24_ = _dafny.Map({d_1266_v20_: d_1245_i0_})
                index222_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                index223_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                rhs269_ = (len(d_1271_v24_)) >= (len(default__.fm36(self.f7, (d_1263_v19_)[default__.safeIndex(426, (d_1263_v19_).length(0))], (self).f6, (self).f11, globalState)))
                rhs270_ = d_1266_v20_
                rhs271_ = (self).f6
                rhs272_ = (self.f7) + (self.f7)
                lhs187_ = d_1266_v20_
                lhs188_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                lhs189_ = d_1266_v20_
                lhs190_ = default__.safeIndex(190, (d_1266_v20_).length(0))
                d_1244_v2_ = rhs269_
                d_1266_v20_ = rhs270_
                lhs187_[lhs188_] = rhs271_
                lhs189_[lhs190_] = rhs272_
                d_1272_v25_: _dafny.Set
                d_1272_v25_ = _dafny.Set({d_1244_v2_})
                d_1244_v2_ = not((True) in (d_1272_v25_))
            d_1273_v26_: str
            d_1273_v26_ = _dafny.CodePoint('j')
            d_1274_v27_: _dafny.Map
            d_1274_v27_ = _dafny.Map({(_dafny.CodePoint('i') if d_1244_v2_ else d_1273_v26_): (self).f6})
            d_1275_v28_: _dafny.Map
            d_1275_v28_ = _dafny.Map({(self).f11: (self).f6})
            d_1276_v29_: _dafny.Seq
            d_1276_v29_ = _dafny.SeqWithoutIsStrInference([self.f7])
            d_1277_v30_: _dafny.MultiSet
            d_1277_v30_ = _dafny.MultiSet([self.f7, (self).f6, d_1245_i0_, (d_1243_v1_).cardinality])
            d_1278_v31_: _dafny.Set
            d_1278_v31_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([self.f7])), 774, (d_1277_v30_).cardinality})
            d_1279_v32_: _dafny.Array
            nw219_ = _dafny.Array(None, 23)
            nw219_[int(0)] = self.f7
            nw219_[int(1)] = (self).f6
            nw219_[int(2)] = ((d_1275_v28_)[d_1244_v2_] if (d_1244_v2_) in (d_1275_v28_) else self.f7)
            nw219_[int(3)] = (self).f6
            nw219_[int(4)] = (self).f6
            nw219_[int(5)] = d_1245_i0_
            nw219_[int(6)] = 802
            nw219_[int(7)] = len(d_1276_v29_)
            nw219_[int(8)] = d_1245_i0_
            nw219_[int(9)] = self.f7
            nw219_[int(10)] = d_1245_i0_
            nw219_[int(11)] = (self).f6
            nw219_[int(12)] = self.f7
            nw219_[int(13)] = (self).f6
            nw219_[int(14)] = (self).f6
            nw219_[int(15)] = self.f7
            nw219_[int(16)] = d_1245_i0_
            nw219_[int(17)] = self.f7
            nw219_[int(18)] = d_1245_i0_
            nw219_[int(19)] = d_1245_i0_
            nw219_[int(20)] = len(d_1278_v31_)
            nw219_[int(21)] = d_1245_i0_
            nw219_[int(22)] = (self).f6
            d_1279_v32_ = nw219_
            d_1280_v33_: _dafny.MultiSet
            d_1280_v33_ = _dafny.MultiSet([d_1279_v32_])
            d_1281_v34_: _dafny.Seq
            d_1281_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpevpx"))
            d_1282_v35_: _dafny.Array
            nw220_ = _dafny.Array(False, 9)
            d_1282_v35_ = nw220_
            d_1283_v36_: _dafny.Map
            d_1283_v36_ = _dafny.Map({d_1282_v35_: d_1245_i0_})
            d_1284_v37_: _dafny.Map
            d_1284_v37_ = _dafny.Map({default__.fm1(False, globalState): self.f7})
            d_1285_v38_: _dafny.Array
            def lambda76_(d_1286_i5_):
                return _dafny.CodePoint('f')

            init43_ = lambda76_
            nw221_ = _dafny.Array(None, 16)
            for i0_43_ in range(nw221_.length(0)):
                nw221_[i0_43_] = init43_(i0_43_)
            d_1285_v38_ = nw221_
            d_1287_v39_: _dafny.Map
            d_1287_v39_ = _dafny.Map({(self).f11: d_1285_v38_})
            rhs273_ = default__.fm37((len(_dafny.Set({d_1244_v2_}))) - (len(_dafny.SeqWithoutIsStrInference([self.f7 for d_1288_i4_ in range(default__.abs(714))]))), D1_DC3(not(True), d_1245_i0_, len(d_1281_v34_)), globalState)
            rhs274_ = (0) - (self.f7)
            rhs275_ = (d_1282_v35_) in (d_1283_v36_)
            rhs276_ = default__.fm0(not(not(d_1244_v2_)), (d_1284_v37_) | (d_1284_v37_), len(d_1287_v39_), globalState)
            rhs277_ = d_1280_v33_
            lhs191_ = self
            d_1274_v27_ = rhs273_
            lhs191_.f7 = rhs274_
            d_1244_v2_ = rhs275_
            d_1244_v2_ = rhs276_
            d_1280_v33_ = rhs277_
            d_1244_v2_ = not((self).f11)
            d_1289_v40_: _dafny.Array
            def lambda77_(d_1290_i6_):
                return D1_DC2(_dafny.Map({D0_DC0(_dafny.Set({(self).f11})): (self).f11}))

            init44_ = lambda77_
            nw222_ = _dafny.Array(None, 15)
            for i0_44_ in range(nw222_.length(0)):
                nw222_[i0_44_] = init44_(i0_44_)
            d_1289_v40_ = nw222_
            d_1291_v41_: D6
            d_1291_v41_ = D6_DC13(d_1289_v40_)
            source16_ = d_1291_v41_
            if source16_.is_DC14:
                d_1292___mcc_h0_ = source16_.cf23
                d_1293___mcc_h1_ = source16_.cf24
                d_1294___mcc_h2_ = source16_.cf25
                d_1295_cf25_ = d_1294___mcc_h2_
                d_1296_cf24_ = d_1293___mcc_h1_
                d_1297_cf23_ = d_1292___mcc_h0_
                d_1297_cf23_ = default__.safeModuloInt(((0) - ((self).f6)) + (len(d_1278_v31_)), (d_1297_cf23_) * (self.f7))
                d_1298_v42_: _dafny.Seq
                d_1298_v42_ = _dafny.SeqWithoutIsStrInference([d_1244_v2_])
                d_1299_v43_: _dafny.Seq
                d_1299_v43_ = _dafny.SeqWithoutIsStrInference([(d_1298_v42_)[default__.safeIndex(len(d_1298_v42_), len(d_1298_v42_))], True])
                index224_ = default__.safeIndex(453, (d_1279_v32_).length(0))
                (d_1279_v32_)[index224_] = ((d_1275_v28_)[d_1244_v2_] if (d_1244_v2_) in (d_1275_v28_) else len(d_1299_v43_))
                d_1300_v44_: D2
                d_1300_v44_ = D2_DC5(d_1279_v32_)
                d_1301_v45_: _dafny.Array
                nw223_ = _dafny.Array(None, 2)
                nw223_[int(0)] = d_1300_v44_
                nw223_[int(1)] = d_1300_v44_
                d_1301_v45_ = nw223_
                index225_ = default__.safeIndex(747, (d_1282_v35_).length(0))
                (d_1282_v35_)[index225_] = True
                index226_ = default__.safeIndex(900, (d_1282_v35_).length(0))
                (d_1282_v35_)[index226_] = d_1244_v2_
                index227_ = default__.safeIndex(453, (d_1279_v32_).length(0))
                index228_ = default__.safeIndex(747, (d_1282_v35_).length(0))
                index229_ = default__.safeIndex(900, (d_1282_v35_).length(0))
                rhs278_ = default__.fm1((self).f11, globalState)
                rhs279_ = d_1301_v45_
                rhs280_ = d_1273_v26_
                rhs281_ = (self).f11
                rhs282_ = (455) != ((len(_dafny.Set({d_1244_v2_, (self).f11, d_1244_v2_})) if (self).f11 else d_1245_i0_))
                lhs192_ = d_1279_v32_
                lhs193_ = default__.safeIndex(453, (d_1279_v32_).length(0))
                lhs194_ = d_1282_v35_
                lhs195_ = default__.safeIndex(747, (d_1282_v35_).length(0))
                lhs196_ = d_1282_v35_
                lhs197_ = default__.safeIndex(900, (d_1282_v35_).length(0))
                lhs192_[lhs193_] = rhs278_
                d_1301_v45_ = rhs279_
                d_1273_v26_ = rhs280_
                lhs194_[lhs195_] = rhs281_
                lhs196_[lhs197_] = rhs282_
                d_1244_v2_ = (d_1282_v35_)[default__.safeIndex(747, (d_1282_v35_).length(0))]
                d_1302_v46_: _dafny.Set
                d_1302_v46_ = _dafny.Set({d_1244_v2_, d_1244_v2_, False})
                d_1303_v47_: D7
                d_1303_v47_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dufjuu")), d_1302_v46_)
                d_1304_v48_: T0
                nw224_ = C2()
                nw224_.ctor__(d_1295_cf25_, d_1295_cf25_)
                d_1304_v48_ = nw224_
                d_1305_v49_: _dafny.Map
                d_1305_v49_ = _dafny.Map({d_1304_v48_: d_1302_v46_})
                pat_let_tv28_ = d_1305_v49_
                pat_let_tv29_ = d_1304_v48_
                pat_let_tv30_ = d_1304_v48_
                pat_let_tv31_ = d_1305_v49_
                pat_let_tv32_ = d_1302_v46_
                d_1306_v50_: _dafny.Map
                def iife89_(_pat_let25_0):
                    def iife90_(d_1307_dt__update__tmp_h0_):
                        def iife91_(_pat_let26_0):
                            def iife92_(d_1308_dt__update_hcf32_h0_):
                                return D7_DC17((d_1307_dt__update__tmp_h0_).cf31, d_1308_dt__update_hcf32_h0_)
                            return iife92_(_pat_let26_0)
                        return iife91_(((pat_let_tv28_)[pat_let_tv29_] if (pat_let_tv30_) in (pat_let_tv31_) else pat_let_tv32_))
                    return iife90_(_pat_let25_0)
                d_1306_v50_ = _dafny.Map({iife89_(d_1303_v47_): 374})
                d_1309_v51_: _dafny.Array
                nw225_ = _dafny.Array(int(0), 1)
                d_1309_v51_ = nw225_
                (self).f7 = ((d_1306_v50_)[D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txexlpnd")), d_1302_v46_)] if (D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txexlpnd")), d_1302_v46_)) in (d_1306_v50_) else len(_dafny.SeqWithoutIsStrInference([d_1309_v51_])))
            elif source16_.is_DC15:
                d_1310___mcc_h3_ = source16_.cf26
                d_1311___mcc_h4_ = source16_.cf27
                d_1312___mcc_h5_ = source16_.cf28
                d_1313___mcc_h6_ = source16_.cf29
                d_1314_cf29_ = d_1313___mcc_h6_
                d_1315_cf28_ = d_1312___mcc_h5_
                d_1316_cf27_ = d_1311___mcc_h4_
                d_1317_cf26_ = d_1310___mcc_h3_
                d_1317_cf26_ = d_1317_cf26_
                index230_ = default__.safeIndex(219, (d_1279_v32_).length(0))
                (d_1279_v32_)[index230_] = d_1314_cf29_
                index231_ = default__.safeIndex(219, (d_1279_v32_).length(0))
                (d_1279_v32_)[index231_] = d_1314_cf29_
                index232_ = default__.safeIndex(219, (d_1279_v32_).length(0))
                (d_1279_v32_)[index232_] = -383
                d_1317_cf26_ = (d_1316_cf27_) >= (-823)
            elif True:
                d_1318___mcc_h7_ = source16_.cf22
                d_1319_cf22_ = d_1318___mcc_h7_
                d_1284_v37_ = (d_1284_v37_).set((self).f6, self.f7)
                d_1320_v52_: _dafny.Map
                d_1320_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1244_v2_, d_1244_v2_]): d_1285_v38_})
                d_1321_v53_: _dafny.Seq
                d_1321_v53_ = _dafny.SeqWithoutIsStrInference([d_1244_v2_])
                d_1320_v52_ = (d_1320_v52_).set((d_1321_v53_) + ((d_1321_v53_).set(default__.safeIndex(self.f7, len(d_1321_v53_)), not((self).f11))), d_1285_v38_)
                d_1284_v37_ = (d_1284_v37_).set((0) - ((self).f6), d_1245_i0_)
                d_1275_v28_ = (d_1275_v28_).set(d_1244_v2_, (self).f6)
        d_1322_i7_: int
        d_1322_i7_ = 0
        with _dafny.label("7"):
            while (self).f11:
                with _dafny.c_label("7"):
                    if (d_1322_i7_) >= (100):
                        raise _dafny.Break("7")
                    d_1322_i7_ = (d_1322_i7_) + (1)
                    (self).f7 = self.f7
                    d_1323_v54_: _dafny.MultiSet
                    d_1323_v54_ = _dafny.MultiSet([(0) - (len(d_1242_v0_))])
                    d_1324_v55_: _dafny.Map
                    d_1324_v55_ = _dafny.Map({self.f7: (d_1323_v54_).cardinality})
                    d_1325_v56_: _dafny.Set
                    d_1325_v56_ = _dafny.Set({d_1324_v55_})
                    if (d_1325_v56_).issubset((d_1325_v56_) - (d_1325_v56_)):
                        d_1326_v57_: _dafny.Set
                        d_1326_v57_ = _dafny.Set({d_1244_v2_})
                        d_1327_v58_: D0
                        d_1327_v58_ = D0_DC0(d_1326_v57_)
                        d_1328_v59_: _dafny.Map
                        d_1328_v59_ = _dafny.Map({d_1327_v58_: (self).f11})
                        d_1328_v59_ = d_1328_v59_
                        d_1329_v60_: _dafny.Array
                        def lambda78_(d_1330_i8_):
                            return default__.safeModuloInt(d_1330_i8_, self.f7)

                        init45_ = lambda78_
                        nw226_ = _dafny.Array(None, 12)
                        for i0_45_ in range(nw226_.length(0)):
                            nw226_[i0_45_] = init45_(i0_45_)
                        d_1329_v60_ = nw226_
                        index233_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        (d_1329_v60_)[index233_] = (0) - (self.f7)
                        index234_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        (d_1329_v60_)[index234_] = (self).f6
                        d_1244_v2_ = (self).f11
                        d_1244_v2_ = (True) == (d_1244_v2_)
                        d_1331_v61_: _dafny.Array
                        nw227_ = _dafny.Array(False, 8)
                        d_1331_v61_ = nw227_
                        index235_ = default__.safeIndex(899, (d_1331_v61_).length(0))
                        (d_1331_v61_)[index235_] = not(d_1244_v2_)
                        d_1332_v62_: T0
                        nw228_ = C2()
                        nw228_.ctor__(self.f7, (self).f6)
                        d_1332_v62_ = nw228_
                        d_1333_v63_: _dafny.Map
                        d_1333_v63_ = _dafny.Map({d_1332_v62_: (self).f6})
                        index236_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        index237_ = default__.safeIndex(899, (d_1331_v61_).length(0))
                        index238_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        rhs283_ = (d_1243_v1_).ispropersubset(d_1243_v1_)
                        rhs284_ = ((d_1333_v63_)[d_1332_v62_] if (d_1332_v62_) in (d_1333_v63_) else (0) - ((0) - (((d_1332_v62_).f6) * (d_1332_v62_.f7))))
                        rhs285_ = (self).f11
                        rhs286_ = self.f7
                        lhs198_ = d_1329_v60_
                        lhs199_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        lhs200_ = d_1331_v61_
                        lhs201_ = default__.safeIndex(899, (d_1331_v61_).length(0))
                        lhs202_ = d_1329_v60_
                        lhs203_ = default__.safeIndex(134, (d_1329_v60_).length(0))
                        d_1244_v2_ = rhs283_
                        lhs198_[lhs199_] = rhs284_
                        lhs200_[lhs201_] = rhs285_
                        lhs202_[lhs203_] = rhs286_
                    elif True:
                        d_1244_v2_ = (self).f11
                        d_1334_v64_: _dafny.Array
                        nw229_ = _dafny.Array(_dafny.Set({}), 28)
                        d_1334_v64_ = nw229_
                        d_1335_v65_: _dafny.Array
                        def lambda79_(d_1336_v2_):
                            def lambda80_(d_1337_i9_):
                                return d_1336_v2_

                            return lambda80_

                        init46_ = lambda79_(d_1244_v2_)
                        nw230_ = _dafny.Array(None, 22)
                        for i0_46_ in range(nw230_.length(0)):
                            nw230_[i0_46_] = init46_(i0_46_)
                        d_1335_v65_ = nw230_
                        index239_ = default__.safeIndex(928, (d_1334_v64_).length(0))
                        (d_1334_v64_)[index239_] = _dafny.Set({d_1335_v65_})
                        d_1338_v66_: _dafny.Set
                        d_1338_v66_ = _dafny.Set({d_1335_v65_, d_1335_v65_})
                        d_1339_v67_: D11
                        d_1339_v67_ = D11_DC27(d_1338_v66_)
                        index240_ = default__.safeIndex(928, (d_1334_v64_).length(0))
                        (d_1334_v64_)[index240_] = (d_1339_v67_).cf53
                        d_1340_v68_: _dafny.Array
                        nw231_ = _dafny.Array(int(0), 1)
                        d_1340_v68_ = nw231_
                        d_1341_v69_: _dafny.Seq
                        d_1341_v69_ = _dafny.SeqWithoutIsStrInference([d_1242_v0_])
                        index241_ = default__.safeIndex(115, (d_1340_v68_).length(0))
                        (d_1340_v68_)[index241_] = len((d_1341_v69_).set(default__.safeIndex(self.f7, len(d_1341_v69_)), d_1242_v0_))
                        d_1342_v70_: _dafny.Seq
                        d_1342_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "um"))
                        index242_ = default__.safeIndex(115, (d_1340_v68_).length(0))
                        (d_1340_v68_)[index242_] = default__.fm1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ut"))) < (d_1342_v70_), globalState)
                        d_1343_v71_: _dafny.Seq
                        d_1343_v71_ = _dafny.SeqWithoutIsStrInference([self.f7])
                        d_1343_v71_ = d_1343_v71_
                        d_1244_v2_ = (self).f11
                    d_1344_v73_: _dafny.Seq
                    d_1344_v73_ = _dafny.SeqWithoutIsStrInference([(self).f6, -596])
                    d_1345_v74_: _dafny.Seq
                    d_1345_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utk"))
                    d_1346_v75_: C0
                    nw232_ = C0()
                    def iife93_():
                        coll39_ = _dafny.Map()
                        compr_39_: int
                        for compr_39_ in _dafny.IntegerRange(-881, -796):
                            d_1347_v72_: int = compr_39_
                            if ((-881) <= (d_1347_v72_)) and ((d_1347_v72_) < (-796)):
                                coll39_[(d_1347_v72_) * ((d_1243_v1_).cardinality)] = self.f7
                        return _dafny.Map(coll39_)
                    nw232_.ctor__(default__.fm0(default__.fm0(True, iife93_()
                    , (d_1344_v73_)[default__.safeIndex(len(d_1345_v74_), len(d_1344_v73_))], globalState), d_1324_v55_, (0) - ((self).f6), globalState))
                    d_1346_v75_ = nw232_
                    (self).f7 = (self).f6
                    pass
            pass
        d_1348_i10_: int
        d_1348_i10_ = 0
        with _dafny.label("8"):
            while False:
                with _dafny.c_label("8"):
                    if (d_1348_i10_) >= (100):
                        raise _dafny.Break("8")
                    d_1348_i10_ = (d_1348_i10_) + (1)
                    d_1349_v76_: _dafny.Array
                    nw233_ = _dafny.Array(False, 12)
                    d_1349_v76_ = nw233_
                    d_1350_v77_: _dafny.Seq
                    d_1350_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdmrd"))
                    d_1351_v78_: _dafny.Seq
                    d_1351_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpdetefa"))])
                    d_1352_v79_: str
                    d_1352_v79_ = _dafny.CodePoint('y')
                    d_1353_v80_: D12
                    d_1353_v80_ = D12_DC30(d_1351_v78_)
                    d_1354_v81_: _dafny.Array
                    nw234_ = _dafny.Array(D1.default()(), 24)
                    d_1354_v81_ = nw234_
                    d_1355_v82_: _dafny.Seq
                    d_1355_v82_ = _dafny.SeqWithoutIsStrInference([d_1354_v81_])
                    d_1356_v83_: _dafny.Seq
                    d_1356_v83_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7, (self).f6])
                    d_1357_v84_: _dafny.Seq
                    d_1357_v84_ = _dafny.SeqWithoutIsStrInference([d_1244_v2_, (self).f11])
                    d_1358_v85_: _dafny.Array
                    nw235_ = _dafny.Array(None, 23)
                    nw235_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1350_v77_])
                    nw235_[int(1)] = d_1351_v78_
                    nw235_[int(2)] = (d_1351_v78_).set(default__.safeIndex(len(_dafny.Map({(self).f6: (self).f11})), len(d_1351_v78_)), d_1350_v77_)
                    nw235_[int(3)] = d_1351_v78_
                    nw235_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_1350_v77_).set(default__.safeIndex((self).f6, len(d_1350_v77_)), d_1352_v79_), d_1350_v77_])
                    nw235_[int(5)] = d_1351_v78_
                    nw235_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_1350_v77_])) + ((d_1351_v78_).set(default__.safeIndex((self).f6, len(d_1351_v78_)), d_1350_v77_))
                    nw235_[int(7)] = d_1351_v78_
                    nw235_[int(8)] = d_1351_v78_
                    nw235_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1350_v77_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdn")), d_1350_v77_, d_1350_v77_, d_1350_v77_])
                    nw235_[int(10)] = d_1351_v78_
                    nw235_[int(11)] = ((d_1351_v78_) + (d_1351_v78_)).set(default__.safeIndex(self.f7, len((d_1351_v78_) + (d_1351_v78_))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smsluxg")))
                    nw235_[int(12)] = d_1351_v78_
                    nw235_[int(13)] = d_1351_v78_
                    nw235_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_1350_v77_ for d_1359_i11_ in range(default__.abs(196))])) + (d_1351_v78_)
                    nw235_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1350_v77_ for d_1360_i12_ in range(default__.abs(148))])
                    nw235_[int(16)] = d_1351_v78_
                    nw235_[int(17)] = ((d_1353_v80_).cf59).set(default__.safeIndex(len(d_1355_v82_), len((d_1353_v80_).cf59)), default__.fm2(d_1356_v83_, self.f7, self.f7, (self).f11, globalState))
                    nw235_[int(18)] = d_1351_v78_
                    nw235_[int(19)] = default__.fm38(globalState)
                    nw235_[int(20)] = d_1351_v78_
                    nw235_[int(21)] = (d_1351_v78_).set(default__.safeIndex(len(d_1357_v84_), len(d_1351_v78_)), d_1350_v77_)
                    nw235_[int(22)] = d_1351_v78_
                    d_1358_v85_ = nw235_
                    index243_ = default__.safeIndex(779, (d_1358_v85_).length(0))
                    (d_1358_v85_)[index243_] = d_1351_v78_
                    d_1361_v86_: _dafny.Seq
                    d_1361_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D10_DC26(d_1244_v2_)).cf52])])
                    index244_ = default__.safeIndex(779, (d_1358_v85_).length(0))
                    rhs287_ = d_1244_v2_
                    rhs288_ = ((d_1356_v83_)[default__.safeIndex((self).f6, len(d_1356_v83_))] if d_1244_v2_ else default__.fm1(True, globalState))
                    rhs289_ = (d_1349_v76_ if (self.f7) < (self.f7) else d_1349_v76_)
                    rhs290_ = not(((self.f7) > (self.f7)) and ((d_1361_v86_) <= (_dafny.SeqWithoutIsStrInference([d_1357_v84_]))))
                    rhs291_ = (d_1351_v78_) + ((d_1351_v78_).set(default__.safeIndex((self).f6, len(d_1351_v78_)), d_1350_v77_))
                    lhs204_ = globalState
                    lhs205_ = d_1358_v85_
                    lhs206_ = default__.safeIndex(779, (d_1358_v85_).length(0))
                    d_1244_v2_ = rhs287_
                    lhs204_.f1 = rhs288_
                    d_1349_v76_ = rhs289_
                    d_1244_v2_ = rhs290_
                    lhs205_[lhs206_] = rhs291_
                    d_1244_v2_ = (self.f7) <= ((self).f6)
                    index245_ = default__.safeIndex(187, (d_1349_v76_).length(0))
                    (d_1349_v76_)[index245_] = (self).f11
                    index246_ = default__.safeIndex(187, (d_1349_v76_).length(0))
                    (d_1349_v76_)[index246_] = ((self).f6) >= (default__.safeDivisionInt(self.f7, len(d_1356_v83_)))
                    (self).f7 = self.f7
                    pass
            pass
        d_1362_v87_: str
        d_1362_v87_ = _dafny.CodePoint('k')
        d_1363_v88_: D10
        d_1363_v88_ = D10_DC24(d_1362_v87_)
        d_1362_v87_ = (d_1363_v88_).cf51

    def m5(self, p0, globalState):
        d_1364_i0_: int
        d_1364_i0_ = 0
        with _dafny.label("9"):
            while ((self).f11) and ((self).f11):
                with _dafny.c_label("9"):
                    if (d_1364_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1364_i0_ = (d_1364_i0_) + (1)
                    d_1365_v0_: _dafny.Array
                    nw236_ = _dafny.Array(_dafny.Map({}), 15)
                    d_1365_v0_ = nw236_
                    d_1366_v1_: str
                    d_1366_v1_ = _dafny.CodePoint('j')
                    d_1367_v2_: _dafny.Seq
                    d_1367_v2_ = _dafny.SeqWithoutIsStrInference([d_1366_v1_])
                    d_1368_v3_: _dafny.Array
                    nw237_ = _dafny.Array(False, 4)
                    d_1368_v3_ = nw237_
                    d_1369_v4_: _dafny.Map
                    d_1369_v4_ = _dafny.Map({len(d_1367_v2_): d_1368_v3_})
                    index247_ = default__.safeIndex(459, (d_1365_v0_).length(0))
                    (d_1365_v0_)[index247_] = d_1369_v4_
                    index248_ = default__.safeIndex(459, (d_1365_v0_).length(0))
                    (d_1365_v0_)[index248_] = d_1369_v4_
                    d_1370_v5_: _dafny.Seq
                    d_1370_v5_ = _dafny.SeqWithoutIsStrInference([d_1367_v2_])
                    d_1371_v6_: _dafny.Map
                    d_1371_v6_ = _dafny.Map({(self).f11: (self).f11})
                    d_1372_v7_: _dafny.Seq
                    d_1372_v7_ = _dafny.SeqWithoutIsStrInference([self.f7, p0, len(d_1371_v6_)])
                    d_1373_v8_: _dafny.Seq
                    d_1373_v8_ = _dafny.SeqWithoutIsStrInference([len(d_1372_v7_)])
                    d_1374_v9_: _dafny.Map
                    d_1374_v9_ = _dafny.Map({d_1370_v5_: default__.safeModuloInt(len(d_1373_v8_), default__.fm1((self).f11, globalState))})
                    (self).f7 = ((d_1374_v9_)[d_1370_v5_] if (d_1370_v5_) in (d_1374_v9_) else default__.fm1((self).f11, globalState))
                    if not ((self.f7) != (self.f7)) or ((self).f11):
                        d_1375_v10_: bool
                        d_1375_v10_ = True
                        d_1376_v11_: _dafny.Map
                        d_1376_v11_ = _dafny.Map({len(d_1367_v2_): (self).f6})
                        d_1377_v12_: _dafny.Map
                        d_1377_v12_ = _dafny.Map({d_1366_v1_: ((d_1376_v11_)[self.f7] if (self.f7) in (d_1376_v11_) else self.f7)})
                        d_1378_v13_: D7
                        d_1378_v13_ = D7_DC16(d_1377_v12_)
                        pat_let_tv33_ = d_1377_v12_
                        pat_let_tv34_ = d_1377_v12_
                        def iife94_(_pat_let27_0):
                            def iife95_(d_1379_dt__update__tmp_h0_):
                                def iife96_(_pat_let28_0):
                                    def iife97_(d_1380_dt__update_hcf30_h0_):
                                        return D7_DC16(d_1380_dt__update_hcf30_h0_)
                                    return iife97_(_pat_let28_0)
                                return iife96_((pat_let_tv33_) | (pat_let_tv34_))
                            return iife95_(_pat_let27_0)
                        rhs292_ = d_1367_v2_
                        rhs293_ = d_1375_v10_
                        rhs294_ = default__.fm1((self).f11, globalState)
                        rhs295_ = iife94_(d_1378_v13_)
                        lhs207_ = globalState
                        d_1367_v2_ = rhs292_
                        d_1375_v10_ = rhs293_
                        lhs207_.f1 = rhs294_
                        d_1378_v13_ = rhs295_
                        (globalState).f1 = self.f7
                        (globalState).f1 = (-97 if True else self.f7)
                        d_1381_v14_: _dafny.Map
                        d_1381_v14_ = _dafny.Map({p0: (self).f11})
                        d_1381_v14_ = d_1381_v14_
                        index249_ = default__.safeIndex(464, (d_1368_v3_).length(0))
                        (d_1368_v3_)[index249_] = d_1375_v10_
                        d_1382_v15_: _dafny.Set
                        d_1382_v15_ = _dafny.Set({((d_1371_v6_)[True] if (True) in (d_1371_v6_) else d_1375_v10_), (self).f11, d_1375_v10_})
                        index250_ = default__.safeIndex(464, (d_1368_v3_).length(0))
                        rhs296_ = d_1378_v13_
                        rhs297_ = (False) in (d_1382_v15_)
                        rhs298_ = ((self).f6) - ((0) - (len(d_1367_v2_)))
                        lhs208_ = d_1368_v3_
                        lhs209_ = default__.safeIndex(464, (d_1368_v3_).length(0))
                        lhs210_ = globalState
                        d_1378_v13_ = rhs296_
                        lhs208_[lhs209_] = rhs297_
                        lhs210_.f1 = rhs298_
                    elif True:
                        index251_ = default__.safeIndex(928, (d_1368_v3_).length(0))
                        (d_1368_v3_)[index251_] = (self).f11
                        d_1383_v16_: D8
                        d_1383_v16_ = D8_DC20((self).f11, (self).f11, (self).f11, (self).f11)
                        d_1384_v17_: _dafny.Seq
                        def iife98_(_pat_let29_0):
                            def iife99_(d_1385_dt__update__tmp_h1_):
                                def iife100_(_pat_let30_0):
                                    def iife101_(d_1386_dt__update_hcf39_h0_):
                                        return D8_DC20(d_1386_dt__update_hcf39_h0_, (d_1385_dt__update__tmp_h1_).cf40, (d_1385_dt__update__tmp_h1_).cf41, (d_1385_dt__update__tmp_h1_).cf42)
                                    return iife101_(_pat_let30_0)
                                return iife100_((self).f11)
                            return iife99_(_pat_let29_0)
                        d_1384_v17_ = _dafny.SeqWithoutIsStrInference([(iife98_(d_1383_v16_)).cf39, (self).f11])
                        d_1387_v18_: _dafny.Map
                        d_1387_v18_ = _dafny.Map({(self).f11: d_1368_v3_})
                        d_1388_v19_: _dafny.Map
                        d_1388_v19_ = _dafny.Map({_dafny.Map({(self).f11: default__.fm35(len(d_1384_v17_), len(d_1387_v18_), globalState)}): (self).f11})
                        d_1389_v20_: _dafny.Map
                        d_1389_v20_ = _dafny.Map({(self).f11: d_1372_v7_})
                        index252_ = default__.safeIndex(928, (d_1368_v3_).length(0))
                        (d_1368_v3_)[index252_] = ((d_1388_v19_)[d_1389_v20_] if (d_1389_v20_) in (d_1388_v19_) else False)
                        d_1390_v21_: _dafny.Map
                        d_1390_v21_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctwrgg")): self.f7})
                        d_1391_v22_: _dafny.Map
                        d_1391_v22_ = _dafny.Map({len(d_1390_v21_): (d_1368_v3_)[default__.safeIndex(928, (d_1368_v3_).length(0))]})
                        d_1391_v22_ = (d_1391_v22_).set(p0, (d_1368_v3_)[default__.safeIndex(928, (d_1368_v3_).length(0))])
                        d_1392_v23_: _dafny.Array
                        nw238_ = _dafny.Array(D6.default()(), 19)
                        d_1392_v23_ = nw238_
                        index253_ = default__.safeIndex(590, (d_1392_v23_).length(0))
                        (d_1392_v23_)[index253_] = D6_DC15((self).f11, (self).f6, (self).f6, p0)
                        d_1393_v24_: _dafny.Array
                        nw239_ = _dafny.Array(int(0), 10)
                        d_1393_v24_ = nw239_
                        d_1394_v25_: D6
                        d_1394_v25_ = D6_DC15((d_1368_v3_)[default__.safeIndex(928, (d_1368_v3_).length(0))], default__.fm1(not((d_1368_v3_)[default__.safeIndex(928, (d_1368_v3_).length(0))]), globalState), self.f7, p0)
                        pat_let_tv35_ = p0
                        index254_ = default__.safeIndex(590, (d_1392_v23_).length(0))
                        def iife102_(_pat_let31_0):
                            def iife103_(d_1395_dt__update__tmp_h2_):
                                def iife104_(_pat_let32_0):
                                    def iife105_(d_1396_dt__update_hcf29_h0_):
                                        def iife106_(_pat_let33_0):
                                            def iife107_(d_1397_dt__update_hcf27_h0_):
                                                return D6_DC15((d_1395_dt__update__tmp_h2_).cf26, d_1397_dt__update_hcf27_h0_, (d_1395_dt__update__tmp_h2_).cf28, d_1396_dt__update_hcf29_h0_)
                                            return iife107_(_pat_let33_0)
                                        return iife106_(pat_let_tv35_)
                                    return iife105_(_pat_let32_0)
                                return iife104_((self).f6)
                            return iife103_(_pat_let31_0)
                        rhs299_ = iife102_(d_1394_v25_)
                        rhs300_ = d_1393_v24_
                        lhs211_ = d_1392_v23_
                        lhs212_ = default__.safeIndex(590, (d_1392_v23_).length(0))
                        lhs211_[lhs212_] = rhs299_
                        d_1393_v24_ = rhs300_
                        index255_ = default__.safeIndex(928, (d_1368_v3_).length(0))
                        (d_1368_v3_)[index255_] = (self.f7) == (p0)
                        d_1372_v7_ = (d_1373_v8_) + (_dafny.SeqWithoutIsStrInference([912 for d_1398_i1_ in range(default__.abs(84))]))
                    d_1399_v26_: _dafny.MultiSet
                    d_1399_v26_ = _dafny.MultiSet([False, (self).f11])
                    index256_ = default__.safeIndex(897, (d_1368_v3_).length(0))
                    (d_1368_v3_)[index256_] = (d_1399_v26_).issubset(d_1399_v26_)
                    index257_ = default__.safeIndex(897, (d_1368_v3_).length(0))
                    (d_1368_v3_)[index257_] = (self).f11
                    pass
            pass
        d_1400_v27_: _dafny.Seq
        d_1400_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnumj"))
        d_1401_v28_: str
        d_1401_v28_ = _dafny.CodePoint('r')
        d_1402_v29_: _dafny.Seq
        d_1402_v29_ = _dafny.SeqWithoutIsStrInference([self.f7, self.f7])
        d_1403_v30_: _dafny.MultiSet
        d_1403_v30_ = _dafny.MultiSet([self.f7, self.f7])
        d_1404_v31_: _dafny.Array
        nw240_ = _dafny.Array(None, 18)
        nw240_[int(0)] = d_1400_v27_
        nw240_[int(1)] = d_1400_v27_
        nw240_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeil")) if (self).f11 else d_1400_v27_)
        nw240_[int(3)] = d_1400_v27_
        nw240_[int(4)] = (d_1400_v27_).set(default__.safeIndex(p0, len(d_1400_v27_)), d_1401_v28_)
        nw240_[int(5)] = d_1400_v27_
        nw240_[int(6)] = d_1400_v27_
        nw240_[int(7)] = d_1400_v27_
        nw240_[int(8)] = d_1400_v27_
        nw240_[int(9)] = d_1400_v27_
        nw240_[int(10)] = d_1400_v27_
        nw240_[int(11)] = d_1400_v27_
        nw240_[int(12)] = (default__.fm2(d_1402_v29_, ((d_1403_v30_)[p0] if (p0) in (d_1403_v30_) else (self).f6), (self).f6, (self).f11, globalState)) + (d_1400_v27_)
        nw240_[int(13)] = (d_1400_v27_) + (d_1400_v27_)
        nw240_[int(14)] = d_1400_v27_
        nw240_[int(15)] = d_1400_v27_
        nw240_[int(16)] = (d_1400_v27_) + (d_1400_v27_)
        nw240_[int(17)] = d_1400_v27_
        d_1404_v31_ = nw240_
        d_1405_v32_: _dafny.Map
        d_1405_v32_ = _dafny.Map({self.f7: len(d_1400_v27_)})
        index258_ = default__.safeIndex(56, (d_1404_v31_).length(0))
        (d_1404_v31_)[index258_] = (self).fm12(default__.fm0((self).f11, d_1405_v32_, (0) - (p0), globalState), (self).f11, globalState)
        index259_ = default__.safeIndex(56, (d_1404_v31_).length(0))
        (d_1404_v31_)[index259_] = d_1400_v27_
        d_1406_v33_: _dafny.Set
        d_1406_v33_ = _dafny.Set({p0, (d_1402_v29_)[default__.safeIndex(-56, len(d_1402_v29_))]})
        d_1407_v34_: _dafny.Map
        d_1407_v34_ = _dafny.Map({(self).f11: D12_DC31(282)})
        d_1408_v35_: _dafny.MultiSet
        d_1408_v35_ = _dafny.MultiSet([d_1407_v34_, d_1407_v34_])
        d_1409_v36_: _dafny.Map
        d_1409_v36_ = _dafny.Map({len(d_1406_v33_): d_1408_v35_})
        d_1409_v36_ = (d_1409_v36_).set(((self).f6 if (self).f11 else (self).f6), d_1408_v35_)
        d_1410_v37_: bool
        d_1410_v37_ = True
        d_1411_v38_: _dafny.Seq
        d_1411_v38_ = _dafny.SeqWithoutIsStrInference([True])
        d_1412_v39_: _dafny.Map
        d_1412_v39_ = _dafny.Map({d_1410_v37_: len((d_1404_v31_)[default__.safeIndex(56, (d_1404_v31_).length(0))])})
        d_1410_v37_ = ((_dafny.Set({self.f7, (self).f6, len(d_1411_v38_), (self).f6, len(d_1412_v39_)})) - (d_1406_v33_)) != (_dafny.Set({p0, p0, (0) - (-177), -522}))
        d_1413_v40_: D10
        d_1413_v40_ = D10_DC24(d_1401_v28_)
        source17_ = d_1413_v40_
        if source17_.is_DC25:
            d_1414_v41_: _dafny.Set
            d_1414_v41_ = _dafny.Set({(d_1411_v38_) == (d_1411_v38_), d_1410_v37_})
            d_1414_v41_ = d_1414_v41_
            d_1415_v42_: C3
            nw241_ = C3()
            nw241_.ctor__(len(d_1400_v27_), -326)
            d_1415_v42_ = nw241_
            d_1416_v43_: _dafny.Array
            nw242_ = _dafny.Array(int(0), 20)
            d_1416_v43_ = nw242_
            index260_ = default__.safeIndex(968, (d_1416_v43_).length(0))
            (d_1416_v43_)[index260_] = (0) - (p0)
            index261_ = default__.safeIndex(968, (d_1416_v43_).length(0))
            (d_1416_v43_)[index261_] = len(d_1402_v29_)
            if (p0) >= ((p0) - (-984)):
                d_1417_v44_: _dafny.Array
                nw243_ = _dafny.Array(False, 27)
                d_1417_v44_ = nw243_
                index262_ = default__.safeIndex(450, (d_1417_v44_).length(0))
                (d_1417_v44_)[index262_] = default__.fm0(default__.fm0(True, d_1405_v32_, (self).f6, globalState), d_1405_v32_, len(d_1414_v41_), globalState)
                index263_ = default__.safeIndex(450, (d_1417_v44_).length(0))
                (d_1417_v44_)[index263_] = (self).f11
                d_1418_v45_: _dafny.Array
                nw244_ = _dafny.Array(_dafny.CodePoint('D'), 12)
                d_1418_v45_ = nw244_
                index264_ = default__.safeIndex(652, (d_1418_v45_).length(0))
                (d_1418_v45_)[index264_] = d_1401_v28_
                index265_ = default__.safeIndex(450, (d_1417_v44_).length(0))
                index266_ = default__.safeIndex(652, (d_1418_v45_).length(0))
                rhs301_ = (d_1417_v44_)[default__.safeIndex(450, (d_1417_v44_).length(0))]
                rhs302_ = p0
                rhs303_ = (d_1417_v44_)[default__.safeIndex(450, (d_1417_v44_).length(0))]
                rhs304_ = d_1401_v28_
                rhs305_ = (self).f11
                lhs213_ = d_1417_v44_
                lhs214_ = default__.safeIndex(450, (d_1417_v44_).length(0))
                lhs215_ = globalState
                lhs216_ = d_1418_v45_
                lhs217_ = default__.safeIndex(652, (d_1418_v45_).length(0))
                lhs213_[lhs214_] = rhs301_
                lhs215_.f1 = rhs302_
                d_1410_v37_ = rhs303_
                lhs216_[lhs217_] = rhs304_
                d_1410_v37_ = rhs305_
                d_1419_v46_: _dafny.Map
                d_1419_v46_ = _dafny.Map({p0: True})
                d_1420_v47_: D0
                d_1420_v47_ = D0_DC0(d_1414_v41_)
                d_1421_v48_: _dafny.Map
                d_1421_v48_ = _dafny.Map({d_1420_v47_: not((d_1417_v44_)[default__.safeIndex(450, (d_1417_v44_).length(0))])})
                d_1422_v49_: D1
                d_1422_v49_ = D1_DC2(d_1421_v48_)
                pat_let_tv36_ = d_1421_v48_
                d_1423_v50_: _dafny.Set
                def iife108_(_pat_let34_0):
                    def iife109_(d_1424_dt__update__tmp_h3_):
                        def iife110_(_pat_let35_0):
                            def iife111_(d_1425_dt__update_hcf4_h0_):
                                return D1_DC2(d_1425_dt__update_hcf4_h0_)
                            return iife111_(_pat_let35_0)
                        return iife110_(pat_let_tv36_)
                    return iife109_(_pat_let34_0)
                d_1423_v50_ = _dafny.Set({d_1422_v49_, iife108_(D1_DC2(d_1421_v48_)), d_1422_v49_, d_1422_v49_})
                d_1426_v51_: bool
                out26_: bool
                out26_ = (d_1415_v42_).m6((d_1419_v46_) | (d_1419_v46_), self.f7, d_1423_v50_, globalState)
                d_1426_v51_ = out26_
                d_1427_v52_: _dafny.Array
                def lambda81_(d_1428_v51_, d_1429_v44_):
                    def lambda82_(d_1430_i2_):
                        return D6_DC14(self.f7, _dafny.MultiSet([d_1428_v51_, (d_1429_v44_)[default__.safeIndex(450, (d_1429_v44_).length(0))]]), (self).f6)

                    return lambda82_

                init47_ = lambda81_(d_1426_v51_, d_1417_v44_)
                nw245_ = _dafny.Array(None, 17)
                for i0_47_ in range(nw245_.length(0)):
                    nw245_[i0_47_] = init47_(i0_47_)
                d_1427_v52_ = nw245_
                d_1431_v53_: D6
                d_1431_v53_ = D6_DC14(len(_dafny.SeqWithoutIsStrInference([not((d_1417_v44_)[default__.safeIndex(450, (d_1417_v44_).length(0))]), d_1426_v51_])), _dafny.MultiSet([d_1410_v37_]), self.f7)
                index267_ = default__.safeIndex(706, (d_1427_v52_).length(0))
                (d_1427_v52_)[index267_] = d_1431_v53_
                index268_ = default__.safeIndex(706, (d_1427_v52_).length(0))
                (d_1427_v52_)[index268_] = d_1431_v53_
                index269_ = default__.safeIndex(968, (d_1416_v43_).length(0))
                (d_1416_v43_)[index269_] = default__.fm1(((d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))]) > (135), globalState)
            elif True:
                d_1432_v54_: C0
                nw246_ = C0()
                nw246_.ctor__(((d_1404_v31_)[default__.safeIndex(56, (d_1404_v31_).length(0))]) == (d_1400_v27_))
                d_1432_v54_ = nw246_
                d_1433_v55_: _dafny.MultiSet
                d_1433_v55_ = _dafny.MultiSet([(self).f11, (self).f11, (self).f11])
                d_1434_v56_: D6
                d_1434_v56_ = D6_DC15((d_1410_v37_) == (d_1410_v37_), p0, (857) * ((d_1402_v29_)[default__.safeIndex(p0, len(d_1402_v29_))]), default__.safeDivisionInt((_dafny.MultiSet([(d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))]])).cardinality, (self).f6))
                pat_let_tv37_ = d_1401_v28_
                d_1435_v57_: _dafny.Seq
                def iife112_(_pat_let36_0):
                    def iife113_(d_1436_dt__update__tmp_h4_):
                        def iife114_(_pat_let37_0):
                            def iife115_(d_1437_dt__update_hcf51_h0_):
                                return D10_DC24(d_1437_dt__update_hcf51_h0_)
                            return iife115_(_pat_let37_0)
                        return iife114_(pat_let_tv37_)
                    return iife113_(_pat_let36_0)
                d_1435_v57_ = _dafny.SeqWithoutIsStrInference([(d_1413_v40_ if not(d_1432_v54_.f10) else iife112_(d_1413_v40_)), d_1413_v40_, d_1413_v40_, d_1413_v40_])
                d_1438_v58_: _dafny.MultiSet
                d_1438_v58_ = _dafny.MultiSet([d_1432_v54_.f10])
                rhs306_ = ((d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))]) - ((self.f7) + ((d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))]))
                rhs307_ = d_1438_v58_
                rhs308_ = d_1434_v56_
                rhs309_ = (_dafny.SeqWithoutIsStrInference([default__.fm39(len(_dafny.SeqWithoutIsStrInference([p0])), _dafny.CodePoint('h'), p0, (self).f6, globalState), d_1413_v40_])) + (((d_1435_v57_).set(default__.safeIndex((self).f6, len(d_1435_v57_)), d_1413_v40_)) + (d_1435_v57_))
                rhs310_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jl"))
                lhs218_ = globalState
                lhs218_.f1 = rhs306_
                d_1433_v55_ = rhs307_
                d_1434_v56_ = rhs308_
                d_1435_v57_ = rhs309_
                d_1400_v27_ = rhs310_
                rhs311_ = 753
                rhs312_ = p0
                rhs313_ = p0
                lhs219_ = globalState
                lhs220_ = globalState
                lhs221_ = globalState
                lhs219_.f1 = rhs311_
                lhs220_.f1 = rhs312_
                lhs221_.f1 = rhs313_
                (globalState).f1 = ((self.f7) + ((d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))])) - ((d_1416_v43_)[default__.safeIndex(968, (d_1416_v43_).length(0))])
                d_1405_v32_ = (d_1405_v32_).set((356) - (((d_1403_v30_)[p0] if (p0) in (d_1403_v30_) else (self).f6)), self.f7)
        elif source17_.is_DC26:
            d_1439___mcc_h0_ = source17_.cf52
            d_1440_cf52_ = d_1439___mcc_h0_
            d_1441_v59_: _dafny.Array
            nw247_ = _dafny.Array(False, 22)
            d_1441_v59_ = nw247_
            d_1442_v60_: _dafny.Map
            d_1442_v60_ = _dafny.Map({(self).f6: d_1440_cf52_})
            index270_ = default__.safeIndex(505, (d_1441_v59_).length(0))
            (d_1441_v59_)[index270_] = ((d_1442_v60_)[default__.fm1(d_1440_cf52_, globalState)] if (default__.fm1(d_1440_cf52_, globalState)) in (d_1442_v60_) else (self).f11)
            index271_ = default__.safeIndex(505, (d_1441_v59_).length(0))
            (d_1441_v59_)[index271_] = (self).f11
            d_1441_v59_ = d_1441_v59_
            d_1410_v37_ = (self).f11
            d_1443_v61_: _dafny.Array
            nw248_ = _dafny.Array(int(0), 20)
            d_1443_v61_ = nw248_
            index272_ = default__.safeIndex(860, (d_1443_v61_).length(0))
            (d_1443_v61_)[index272_] = ((self).f6 if d_1440_cf52_ else (0) - (self.f7))
            index273_ = default__.safeIndex(860, (d_1443_v61_).length(0))
            (d_1443_v61_)[index273_] = ((0) - (self.f7)) - ((self).f6)
        elif True:
            d_1444___mcc_h1_ = source17_.cf51
            d_1445_cf51_ = d_1444___mcc_h1_
            d_1410_v37_ = not(d_1410_v37_)
            d_1446_v62_: D5
            d_1446_v62_ = D5_DC11(d_1411_v38_)
            d_1447_v63_: _dafny.Map
            d_1447_v63_ = _dafny.Map({d_1446_v62_: (self).f6})
            rhs314_ = (-995) * (default__.safeModuloInt((self).f6, 477))
            rhs315_ = default__.fm0(d_1410_v37_, d_1405_v32_, (self.f7) - ((self).f6), globalState)
            rhs316_ = (d_1411_v38_)[default__.safeIndex(((d_1447_v63_)[d_1446_v62_] if (d_1446_v62_) in (d_1447_v63_) else (self).f6), len(d_1411_v38_))]
            lhs222_ = globalState
            lhs222_.f1 = rhs314_
            d_1410_v37_ = rhs315_
            d_1410_v37_ = rhs316_
            d_1448_v64_: _dafny.Set
            d_1448_v64_ = _dafny.Set({d_1410_v37_})
            d_1449_v65_: _dafny.MultiSet
            d_1449_v65_ = _dafny.MultiSet([default__.fm0((self).f11, _dafny.Map({len(d_1448_v64_): len(d_1411_v38_)}), len(d_1402_v29_), globalState)])
            if (d_1449_v65_).isdisjoint(d_1449_v65_):
                d_1450_v66_: C2
                nw249_ = C2()
                nw249_.ctor__((self).f6, 383)
                d_1450_v66_ = nw249_
                d_1451_v69_: D1
                d_1451_v69_ = D1_DC4((self).f11)
                d_1452_v70_: D2
                def iife116_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(377, -191):
                        d_1453_v68_: int = compr_40_
                        if ((377) <= (d_1453_v68_)) and ((d_1453_v68_) < (-191)):
                            coll40_[default__.safeDivisionInt(d_1453_v68_, self.f7)] = d_1410_v37_
                    return _dafny.Map(coll40_)
                d_1452_v70_ = D2_DC7((self).f11, (d_1402_v29_)[default__.safeIndex(self.f7, len(d_1402_v29_))], iife116_()
, d_1451_v69_)
                d_1454_v71_: D8
                def iife117_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(135, 264):
                        d_1455_v67_: int = compr_41_
                        if ((135) <= (d_1455_v67_)) and ((d_1455_v67_) < (264)):
                            coll41_[(d_1455_v67_) * (len(d_1400_v27_))] = 277
                    return _dafny.Map(coll41_)
                d_1454_v71_ = D8_DC20(default__.fm0(d_1410_v37_, iife117_()
, (self).f6, globalState), d_1410_v37_, (d_1452_v70_).cf13, False)
                d_1456_v72_: _dafny.Map
                d_1456_v72_ = _dafny.Map({len(d_1405_v32_): d_1454_v71_})
                d_1456_v72_ = (d_1456_v72_).set(p0, d_1454_v71_)
                d_1410_v37_ = (d_1411_v38_)[default__.safeIndex(((0) - (len(d_1411_v38_))) + (self.f7), len(d_1411_v38_))]
                d_1457_v73_: _dafny.Array
                nw250_ = _dafny.Array(_dafny.CodePoint('D'), 16)
                d_1457_v73_ = nw250_
                d_1457_v73_ = d_1457_v73_
                d_1458_v74_: _dafny.Map
                d_1458_v74_ = _dafny.Map({d_1410_v37_: d_1410_v37_})
                d_1459_v75_: D0
                d_1459_v75_ = D0_DC0(d_1448_v64_)
                d_1460_v76_: _dafny.Map
                d_1460_v76_ = _dafny.Map({d_1459_v75_: d_1410_v37_})
                d_1461_v77_: D1
                d_1461_v77_ = D1_DC2(d_1460_v76_)
                pat_let_tv38_ = d_1459_v75_
                pat_let_tv39_ = d_1460_v76_
                d_1462_v78_: D2
                def iife119_(_pat_let39_0):
                    def iife120_(d_1463_dt__update__tmp_h6_):
                        def iife121_(_pat_let40_0):
                            def iife122_(d_1464_dt__update_hcf4_h1_):
                                return D1_DC2(d_1464_dt__update_hcf4_h1_)
                            return iife122_(_pat_let40_0)
                        return iife121_(_dafny.Map({pat_let_tv38_: (self).f11}))
                    return iife120_(_pat_let39_0)
                def iife118_(_pat_let38_0):
                    def iife123_(d_1465_dt__update__tmp_h7_):
                        def iife124_(_pat_let41_0):
                            def iife125_(d_1466_dt__update_hcf4_h2_):
                                return D1_DC2(d_1466_dt__update_hcf4_h2_)
                            return iife125_(_pat_let41_0)
                        return iife124_(pat_let_tv39_)
                    return iife123_(_pat_let38_0)
                d_1462_v78_ = D2_DC6(iife118_(iife119_(d_1461_v77_)), (d_1404_v31_)[default__.safeIndex(56, (d_1404_v31_).length(0))], (self).f6)
                (d_1450_v66_).m8(d_1458_v74_, d_1462_v78_, _dafny.SeqWithoutIsStrInference([d_1445_cf51_ for d_1467_i3_ in range(default__.abs(119))]), globalState)
            elif True:
                d_1468_v79_: _dafny.Map
                d_1468_v79_ = _dafny.Map({d_1410_v37_: (d_1448_v64_) - (d_1448_v64_)})
                d_1469_v80_: _dafny.Array
                nw251_ = _dafny.Array(_dafny.Set({}), 1)
                d_1469_v80_ = nw251_
                d_1470_v81_: _dafny.Map
                d_1470_v81_ = _dafny.Map({268: (self).f11})
                rhs317_ = p0
                rhs318_ = d_1468_v79_
                rhs319_ = default__.safeDivisionInt(self.f7, ((self).f6) + (len(d_1470_v81_)))
                rhs320_ = (d_1469_v80_ if default__.fm0(d_1410_v37_, d_1405_v32_, (d_1449_v65_).cardinality, globalState) else d_1469_v80_)
                lhs223_ = globalState
                lhs224_ = globalState
                lhs223_.f1 = rhs317_
                d_1468_v79_ = rhs318_
                lhs224_.f1 = rhs319_
                d_1469_v80_ = rhs320_
                d_1471_v82_: _dafny.Array
                nw252_ = _dafny.Array(None, 20)
                nw252_[int(0)] = d_1410_v37_
                nw252_[int(1)] = (self).f11
                nw252_[int(2)] = d_1410_v37_
                nw252_[int(3)] = d_1410_v37_
                nw252_[int(4)] = not(d_1410_v37_)
                nw252_[int(5)] = (self).f11
                nw252_[int(6)] = False
                nw252_[int(7)] = not((self).f11)
                nw252_[int(8)] = (self).f11
                nw252_[int(9)] = (self).f11
                nw252_[int(10)] = (self).f11
                nw252_[int(11)] = d_1410_v37_
                nw252_[int(12)] = d_1410_v37_
                nw252_[int(13)] = d_1410_v37_
                nw252_[int(14)] = (self).f11
                nw252_[int(15)] = d_1410_v37_
                nw252_[int(16)] = d_1410_v37_
                nw252_[int(17)] = (self).f11
                nw252_[int(18)] = not((self).f11)
                nw252_[int(19)] = d_1410_v37_
                d_1471_v82_ = nw252_
                d_1472_v83_: _dafny.Array
                nw253_ = _dafny.Array(None, 6)
                nw253_[int(0)] = d_1471_v82_
                nw253_[int(1)] = d_1471_v82_
                nw253_[int(2)] = d_1471_v82_
                nw253_[int(3)] = d_1471_v82_
                nw253_[int(4)] = d_1471_v82_
                nw253_[int(5)] = d_1471_v82_
                d_1472_v83_ = nw253_
                index274_ = default__.safeIndex(732, (d_1472_v83_).length(0))
                (d_1472_v83_)[index274_] = d_1471_v82_
                index275_ = default__.safeIndex(732, (d_1472_v83_).length(0))
                (d_1472_v83_)[index275_] = d_1471_v82_
                d_1473_v84_: _dafny.Map
                d_1473_v84_ = _dafny.Map({self.f7: d_1471_v82_})
                (globalState).f1 = len((d_1473_v84_) | (d_1473_v84_))
                d_1410_v37_ = default__.fm0(((self).f6) <= ((self).f6), _dafny.Map({(self).f6: (self).f6}), (0) - (self.f7), globalState)
                d_1410_v37_ = (self).f11
            d_1474_v85_: C1
            nw254_ = C1()
            nw254_.ctor__((_dafny.CodePoint('b') if False else d_1401_v28_), (len(d_1400_v27_)) * (len(d_1412_v39_)), self.f7)
            d_1474_v85_ = nw254_
        d_1475_v86_: _dafny.Array
        nw255_ = _dafny.Array(None, 1)
        nw255_[int(0)] = (self).f11
        d_1475_v86_ = nw255_
        d_1475_v86_ = d_1475_v86_

    @property
    def f11(self):
        return self._f11

class C5(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f6: int = int(0)
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f7(self):
        return self._f7
    @f7.setter
    def f7(self, value):
        self._f7 = value
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f8, f9, f6, f7):
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f6 = f6
        (self).f7 = f7

    def fm4(self, globalState):
        return False

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_1476_v0_: str
        d_1476_v0_ = _dafny.CodePoint('p')
        d_1477_v1_: _dafny.Seq
        d_1477_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybpygmgw"))
        d_1476_v0_ = (d_1477_v1_)[default__.safeIndex((p2 if p0 else p2), len(d_1477_v1_))]
        d_1478_v2_: _dafny.Array
        nw256_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_1478_v2_ = nw256_
        d_1479_v3_: _dafny.Array
        nw257_ = _dafny.Array(None, 26)
        nw257_[int(0)] = p0
        nw257_[int(1)] = p0
        nw257_[int(2)] = p0
        nw257_[int(3)] = p0
        nw257_[int(4)] = p0
        nw257_[int(5)] = p0
        nw257_[int(6)] = p0
        nw257_[int(7)] = p0
        nw257_[int(8)] = p0
        nw257_[int(9)] = p0
        nw257_[int(10)] = p0
        nw257_[int(11)] = p0
        nw257_[int(12)] = p0
        nw257_[int(13)] = p0
        nw257_[int(14)] = p0
        nw257_[int(15)] = p0
        nw257_[int(16)] = p0
        nw257_[int(17)] = True
        nw257_[int(18)] = not(p0)
        nw257_[int(19)] = p0
        nw257_[int(20)] = p0
        nw257_[int(21)] = p0
        nw257_[int(22)] = p0
        nw257_[int(23)] = p0
        nw257_[int(24)] = p0
        nw257_[int(25)] = True
        d_1479_v3_ = nw257_
        index276_ = default__.safeIndex(377, (d_1478_v2_).length(0))
        (d_1478_v2_)[index276_] = d_1479_v3_
        index277_ = default__.safeIndex(377, (d_1478_v2_).length(0))
        (d_1478_v2_)[index277_] = d_1479_v3_
        d_1480_v4_: _dafny.Map
        d_1480_v4_ = _dafny.Map({not(p0): p0})
        d_1481_v5_: _dafny.Map
        d_1481_v5_ = _dafny.Map({default__.safeDivisionInt(len(d_1480_v4_), 387): (self).f6})
        r2 = len(d_1481_v5_)
        hi14_ = p2
        for d_1482_i0_ in range((0) - ((self).f6), hi14_):
            d_1483_v6_: _dafny.Array
            nw258_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_1483_v6_ = nw258_
            d_1484_v7_: _dafny.Set
            d_1484_v7_ = _dafny.Set({p0, p0})
            d_1485_v8_: D0
            d_1485_v8_ = D0_DC0(d_1484_v7_)
            d_1486_v9_: _dafny.Map
            d_1486_v9_ = _dafny.Map({d_1485_v8_: p0})
            pat_let_tv40_ = d_1484_v7_
            d_1487_v11_: _dafny.Set
            def iife126_(_pat_let42_0):
                def iife127_(d_1488_dt__update__tmp_h0_):
                    def iife128_(_pat_let43_0):
                        def iife129_(d_1489_dt__update_hcf0_h0_):
                            return D0_DC0(d_1489_dt__update_hcf0_h0_)
                        return iife129_(_pat_let43_0)
                    return iife128_(pat_let_tv40_)
                return iife127_(_pat_let42_0)
            d_1487_v11_ = _dafny.Set({d_1485_v8_, iife126_(d_1485_v8_), D0_DC0(d_1484_v7_)})
            d_1490_v13_: _dafny.Seq
            d_1490_v13_ = _dafny.SeqWithoutIsStrInference([d_1485_v8_])
            d_1491_v14_: D1
            d_1491_v14_ = D1_DC2(d_1486_v9_)
            d_1492_v15_: _dafny.Array
            nw259_ = _dafny.Array(None, 21)
            nw259_[int(0)] = d_1486_v9_
            nw259_[int(1)] = d_1486_v9_
            nw259_[int(2)] = d_1486_v9_
            nw259_[int(3)] = d_1486_v9_
            nw259_[int(4)] = _dafny.Map({d_1485_v8_: True})
            nw259_[int(5)] = d_1486_v9_
            nw259_[int(6)] = d_1486_v9_
            nw259_[int(7)] = d_1486_v9_
            nw259_[int(8)] = d_1486_v9_
            nw259_[int(9)] = d_1486_v9_
            def iife130_():
                coll42_ = _dafny.Map()
                compr_42_: D0
                for compr_42_ in (d_1487_v11_).Elements:
                    d_1493_v10_: D0 = compr_42_
                    if (d_1493_v10_) in (d_1487_v11_):
                        coll42_[d_1493_v10_] = p0
                return _dafny.Map(coll42_)
            nw259_[int(10)] = iife130_()
            
            nw259_[int(11)] = d_1486_v9_
            nw259_[int(12)] = d_1486_v9_
            def iife131_():
                coll43_ = _dafny.Map()
                compr_43_: D0
                for compr_43_ in (d_1490_v13_).Elements:
                    d_1494_v12_: D0 = compr_43_
                    if (d_1494_v12_) in (d_1490_v13_):
                        coll43_[d_1494_v12_] = p0
                return _dafny.Map(coll43_)
            nw259_[int(13)] = iife131_()
            
            nw259_[int(14)] = d_1486_v9_
            nw259_[int(15)] = (d_1486_v9_).set(d_1485_v8_, False)
            nw259_[int(16)] = d_1486_v9_
            nw259_[int(17)] = d_1486_v9_
            nw259_[int(18)] = d_1486_v9_
            nw259_[int(19)] = _dafny.Map({D0_DC0(d_1484_v7_): p0})
            nw259_[int(20)] = (d_1491_v14_).cf4
            d_1492_v15_ = nw259_
            index278_ = default__.safeIndex(429, (d_1483_v6_).length(0))
            (d_1483_v6_)[index278_] = (d_1492_v15_ if p0 else d_1492_v15_)
            d_1495_v16_: _dafny.Seq
            d_1495_v16_ = _dafny.SeqWithoutIsStrInference([d_1492_v15_])
            index279_ = default__.safeIndex(429, (d_1483_v6_).length(0))
            (d_1483_v6_)[index279_] = (d_1495_v16_)[default__.safeIndex((D1_DC3(p0, (0) - (p1), self.f7)).cf6, len(d_1495_v16_))]
            d_1496_v17_: _dafny.Map
            d_1496_v17_ = _dafny.Map({65: (self).f9})
            d_1497_v18_: _dafny.Seq
            d_1497_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f7: d_1482_i0_}))])])
            d_1498_v19_: _dafny.Map
            d_1498_v19_ = _dafny.Map({p0: default__.fm5(p0, p2, p2, p0, globalState)})
            d_1499_v20_: _dafny.Map
            d_1499_v20_ = _dafny.Map({len((self).f9): p0})
            d_1500_v21_: _dafny.Seq
            d_1500_v21_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1501_v22_: _dafny.Map
            d_1501_v22_ = _dafny.Map({d_1500_v21_: d_1482_i0_})
            d_1502_v23_: _dafny.Array
            nw260_ = _dafny.Array(None, 18)
            nw260_[int(0)] = (self).f9
            nw260_[int(1)] = ((self).f9) + ((self).f9)
            nw260_[int(2)] = ((self).f9) + ((self).f9)
            nw260_[int(3)] = ((d_1496_v17_)[d_1482_i0_] if (d_1482_i0_) in (d_1496_v17_) else (self).f9)
            nw260_[int(4)] = (self).f9
            nw260_[int(5)] = (self).f9
            nw260_[int(6)] = ((self).f9) + ((d_1497_v18_)[default__.safeIndex(p2, len(d_1497_v18_))])
            nw260_[int(7)] = (self).f9
            nw260_[int(8)] = (self).f9
            nw260_[int(9)] = ((d_1498_v19_)[p0] if (p0) in (d_1498_v19_) else (self).f9)
            nw260_[int(10)] = (self).f9
            nw260_[int(11)] = (self).f9
            nw260_[int(12)] = _dafny.SeqWithoutIsStrInference([p1, 821, (0) - (len(d_1499_v20_)), (self).f6, p1])
            nw260_[int(13)] = _dafny.SeqWithoutIsStrInference([(self).f6 for d_1503_i1_ in range(default__.abs(-17))])
            nw260_[int(14)] = ((self).f9) + ((self).f9)
            nw260_[int(15)] = ((d_1496_v17_)[((d_1501_v22_)[d_1500_v21_] if (d_1500_v21_) in (d_1501_v22_) else d_1482_i0_)] if (((d_1501_v22_)[d_1500_v21_] if (d_1500_v21_) in (d_1501_v22_) else d_1482_i0_)) in (d_1496_v17_) else (self).f9)
            nw260_[int(16)] = ((self).f9) + ((self).f9)
            nw260_[int(17)] = _dafny.SeqWithoutIsStrInference([(self).f6 for d_1504_i2_ in range(default__.abs(919))])
            d_1502_v23_ = nw260_
            index280_ = default__.safeIndex(973, (d_1502_v23_).length(0))
            (d_1502_v23_)[index280_] = _dafny.SeqWithoutIsStrInference([(self).f6])
            index281_ = default__.safeIndex(973, (d_1502_v23_).length(0))
            (d_1502_v23_)[index281_] = ((default__.fm5(p0, p1, p2, p0, globalState)) + ((d_1497_v18_)[default__.safeIndex((self).f6, len(d_1497_v18_))]) if p0 else (self).f9)
            d_1505_v24_: _dafny.MultiSet
            d_1505_v24_ = _dafny.MultiSet([p2, self.f7, (self).f6, (self).f6])
            d_1506_v25_: _dafny.Set
            d_1506_v25_ = _dafny.Set({len(d_1499_v20_), -479, p2})
            d_1507_v26_: _dafny.Array
            def lambda83_(d_1508_p0_):
                def lambda84_(d_1509_i3_):
                    return default__.safeDivisionInt(d_1509_i3_, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(d_1508_p0_)])).cardinality for d_1510_i4_ in range(default__.abs(924))])))

                return lambda84_

            init48_ = lambda83_(p0)
            nw261_ = _dafny.Array(None, 15)
            for i0_48_ in range(nw261_.length(0)):
                nw261_[i0_48_] = init48_(i0_48_)
            d_1507_v26_ = nw261_
            d_1511_v27_: _dafny.Map
            d_1511_v27_ = _dafny.Map({d_1476_v0_: d_1507_v26_})
            d_1512_v28_: _dafny.Array
            nw262_ = _dafny.Array(None, 23)
            nw262_[int(0)] = (self).f6
            nw262_[int(1)] = p2
            nw262_[int(2)] = len(d_1477_v1_)
            nw262_[int(3)] = p1
            nw262_[int(4)] = (len((d_1502_v23_)[default__.safeIndex(973, (d_1502_v23_).length(0))]) if not(not(p0)) else d_1482_i0_)
            nw262_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")))
            nw262_[int(6)] = (0) - (d_1482_i0_)
            nw262_[int(7)] = d_1482_i0_
            nw262_[int(8)] = d_1482_i0_
            nw262_[int(9)] = (self).f6
            nw262_[int(10)] = (self).f6
            nw262_[int(11)] = (d_1482_i0_) - ((self).f6)
            nw262_[int(12)] = ((d_1502_v23_)[default__.safeIndex(973, (d_1502_v23_).length(0))])[default__.safeIndex((self).f6, len((d_1502_v23_)[default__.safeIndex(973, (d_1502_v23_).length(0))]))]
            nw262_[int(13)] = self.f7
            nw262_[int(14)] = d_1482_i0_
            nw262_[int(15)] = len(d_1498_v19_)
            nw262_[int(16)] = (0) - (p2)
            nw262_[int(17)] = ((d_1505_v24_)[len(d_1506_v25_)] if (len(d_1506_v25_)) in (d_1505_v24_) else d_1482_i0_)
            nw262_[int(18)] = (self).f6
            nw262_[int(19)] = self.f7
            nw262_[int(20)] = len(d_1511_v27_)
            nw262_[int(21)] = default__.safeModuloInt(self.f7, self.f7)
            nw262_[int(22)] = p1
            d_1512_v28_ = nw262_
            d_1507_v26_ = (D2_DC5(d_1512_v28_)).cf9
            if p0:
                d_1513_v29_: _dafny.MultiSet
                d_1513_v29_ = _dafny.MultiSet([p0])
                d_1514_v30_: _dafny.Array
                nw263_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_1514_v30_ = nw263_
                rhs321_ = ((d_1480_v4_)[not (p0) or (p0)] if (not (p0) or (p0)) in (d_1480_v4_) else (-273) < (p1))
                rhs322_ = ((d_1513_v29_) | (_dafny.MultiSet([p0]))) | (d_1513_v29_)
                rhs323_ = d_1514_v30_
                rhs324_ = p0
                r0 = rhs321_
                d_1513_v29_ = rhs322_
                d_1514_v30_ = rhs323_
                r1 = rhs324_
                d_1515_v31_: C0
                nw264_ = C0()
                nw264_.ctor__(not((d_1506_v25_).isdisjoint(d_1506_v25_)))
                d_1515_v31_ = nw264_
                d_1477_v1_ = d_1477_v1_
                (globalState).f1 = ((d_1482_i0_) - (255)) + (self.f7)
                d_1476_v0_ = d_1476_v0_
            elif True:
                d_1516_v32_: _dafny.Map
                d_1516_v32_ = _dafny.Map({d_1477_v1_: p0})
                d_1516_v32_ = (d_1516_v32_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), p0)
                (self).f7 = default__.safeDivisionInt(self.f7, default__.safeDivisionInt(853, len(d_1477_v1_)))
                d_1517_v33_: _dafny.Seq
                d_1517_v33_ = _dafny.SeqWithoutIsStrInference([d_1479_v3_, (d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))], d_1479_v3_])
                d_1479_v3_ = (d_1517_v33_)[default__.safeIndex(d_1482_i0_, len(d_1517_v33_))]
                r1 = p0
                d_1518_v34_: _dafny.MultiSet
                d_1518_v34_ = _dafny.MultiSet([True, p0, p0, p0, p0])
                d_1519_v35_: _dafny.MultiSet
                d_1519_v35_ = d_1518_v34_
                d_1520_v36_: _dafny.Map
                d_1520_v36_ = _dafny.Map({p0: d_1518_v34_})
                d_1521_v37_: _dafny.Array
                nw265_ = _dafny.Array(None, 24)
                nw265_[int(0)] = (d_1518_v34_ if p0 else _dafny.MultiSet([p0, p0]))
                nw265_[int(1)] = d_1518_v34_
                nw265_[int(2)] = default__.fm6(globalState)
                nw265_[int(3)] = d_1518_v34_
                nw265_[int(4)] = (d_1519_v35_)
                nw265_[int(5)] = d_1518_v34_
                nw265_[int(6)] = _dafny.MultiSet([not(p0), p0, p0])
                nw265_[int(7)] = d_1518_v34_
                nw265_[int(8)] = _dafny.MultiSet(d_1500_v21_)
                nw265_[int(9)] = d_1518_v34_
                nw265_[int(10)] = d_1518_v34_
                nw265_[int(11)] = (_dafny.MultiSet([p0])).set(p0, default__.abs(p1))
                nw265_[int(12)] = d_1518_v34_
                nw265_[int(13)] = (d_1518_v34_) | (_dafny.MultiSet([p0]))
                nw265_[int(14)] = d_1518_v34_
                nw265_[int(15)] = d_1518_v34_
                nw265_[int(16)] = (d_1518_v34_).intersection(((d_1520_v36_)[p0] if (p0) in (d_1520_v36_) else default__.fm6(globalState)))
                nw265_[int(17)] = (_dafny.MultiSet([p0])) - (d_1518_v34_)
                nw265_[int(18)] = (d_1518_v34_).intersection(d_1518_v34_)
                nw265_[int(19)] = d_1518_v34_
                nw265_[int(20)] = d_1518_v34_
                nw265_[int(21)] = d_1518_v34_
                nw265_[int(22)] = _dafny.MultiSet([not(p0), p0, p0])
                nw265_[int(23)] = (d_1518_v34_).set(p0, default__.abs(p2))
                d_1521_v37_ = nw265_
                d_1521_v37_ = d_1521_v37_
        d_1522_v38_: C0
        nw266_ = C0()
        nw266_.ctor__(p0)
        d_1522_v38_ = nw266_
        arr0_ = (d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))]
        index282_ = default__.safeIndex(667, ((d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))]).length(0))
        arr0_[index282_] = d_1522_v38_.f10
        arr1_ = (d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))]
        index283_ = default__.safeIndex(667, ((d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))]).length(0))
        arr1_[index283_] = p0
        d_1523_v39_: _dafny.MultiSet
        d_1523_v39_ = _dafny.MultiSet([((d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))])[default__.safeIndex(667, ((d_1478_v2_)[default__.safeIndex(377, (d_1478_v2_).length(0))]).length(0))], p0])
        d_1524_v40_: _dafny.MultiSet
        d_1524_v40_ = d_1523_v39_
        pat_let_tv41_ = d_1522_v38_
        def lambda85_(source18_):
            d_1525___mcc_h0_ = source18_
            d_1526_cf18_ = d_1525___mcc_h0_
            return pat_let_tv41_.f10

        r0 = lambda85_(d_1524_v40_)
        r1 = (d_1477_v1_) < ((d_1477_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrjbfvxeg"))))
        r2 = p2
        return r0, r1, r2

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1527_v0_: _dafny.Array
        nw267_ = _dafny.Array(int(0), 17)
        d_1527_v0_ = nw267_
        d_1528_v1_: str
        d_1528_v1_ = _dafny.CodePoint('g')
        d_1529_v2_: _dafny.Set
        d_1529_v2_ = _dafny.Set({d_1528_v1_})
        index284_ = default__.safeIndex(491, (d_1527_v0_).length(0))
        (d_1527_v0_)[index284_] = len(d_1529_v2_)
        index285_ = default__.safeIndex(491, (d_1527_v0_).length(0))
        (d_1527_v0_)[index285_] = self.f7
        d_1530_v3_: _dafny.Set
        d_1530_v3_ = _dafny.Set({(self).f6, self.f7})
        d_1530_v3_ = (d_1530_v3_) | (d_1530_v3_)
        d_1531_v4_: _dafny.Seq
        d_1531_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wukhtqknx"))
        d_1532_v5_: _dafny.MultiSet
        d_1532_v5_ = _dafny.MultiSet([default__.fm2((self).f9, self.f7, self.f7, not(p1), globalState), d_1531_v4_, d_1531_v4_, d_1531_v4_])
        d_1532_v5_ = d_1532_v5_
        r0 = (p2) and (p1)
        d_1533_i0_: int
        d_1533_i0_ = 0
        with _dafny.label("10"):
            while not((d_1531_v4_) <= (_dafny.SeqWithoutIsStrInference([d_1528_v1_ for d_1587_i1_ in range(default__.abs(139))]))):
                with _dafny.c_label("10"):
                    if (d_1533_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1533_i0_ = (d_1533_i0_) + (1)
                    d_1534_v6_: _dafny.Set
                    d_1534_v6_ = _dafny.Set({p1, p1})
                    d_1535_v7_: D0
                    d_1535_v7_ = D0_DC0(d_1534_v6_)
                    d_1536_v8_: _dafny.Map
                    d_1536_v8_ = _dafny.Map({d_1535_v7_: p1})
                    d_1537_v9_: D1
                    d_1537_v9_ = D1_DC2(d_1536_v8_)
                    d_1538_v10_: D2
                    d_1538_v10_ = D2_DC6(d_1537_v9_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgkjag")), len(d_1534_v6_))
                    d_1539_v11_: _dafny.Seq
                    d_1539_v11_ = _dafny.SeqWithoutIsStrInference([p1, False])
                    d_1540_v12_: D1
                    d_1540_v12_ = D1_DC3(p2, len(d_1539_v11_), (self).f6)
                    pat_let_tv42_ = p1
                    pat_let_tv43_ = d_1536_v8_
                    pat_let_tv44_ = d_1537_v9_
                    pat_let_tv45_ = d_1531_v4_
                    d_1541_v13_: _dafny.Array
                    nw268_ = _dafny.Array(None, 25)
                    nw268_[int(0)] = d_1538_v10_
                    nw268_[int(1)] = d_1538_v10_
                    nw268_[int(2)] = d_1538_v10_
                    def iife132_(_pat_let44_0):
                        def iife133_(d_1543_dt__update__tmp_h0_):
                            def iife134_(_pat_let45_0):
                                def iife135_(d_1544_dt__update_hcf5_h0_):
                                    def iife136_(_pat_let46_0):
                                        def iife137_(d_1545_dt__update_hcf7_h0_):
                                            return D1_DC3(d_1544_dt__update_hcf5_h0_, (d_1543_dt__update__tmp_h0_).cf6, d_1545_dt__update_hcf7_h0_)
                                        return iife137_(_pat_let46_0)
                                    return iife136_(self.f7)
                                return iife135_(_pat_let45_0)
                            return iife134_(pat_let_tv42_)
                        return iife133_(_pat_let44_0)
                    nw268_[int(3)] = D2_DC6(d_1537_v9_, _dafny.SeqWithoutIsStrInference([d_1528_v1_ for d_1542_i2_ in range(default__.abs(759))]), (iife132_(d_1540_v12_)).cf7)
                    nw268_[int(4)] = D2_DC6(d_1537_v9_, d_1531_v4_, self.f7)
                    nw268_[int(5)] = d_1538_v10_
                    def iife138_(_pat_let47_0):
                        def iife139_(d_1546_dt__update__tmp_h1_):
                            def iife140_(_pat_let48_0):
                                def iife141_(d_1547_dt__update_hcf12_h0_):
                                    def iife142_(_pat_let49_0):
                                        def iife143_(d_1548_dt__update_hcf10_h0_):
                                            return D2_DC6(d_1548_dt__update_hcf10_h0_, (d_1546_dt__update__tmp_h1_).cf11, d_1547_dt__update_hcf12_h0_)
                                        return iife143_(_pat_let49_0)
                                    return iife142_(D1_DC2(pat_let_tv43_))
                                return iife141_(_pat_let48_0)
                            return iife140_(self.f7)
                        return iife139_(_pat_let47_0)
                    nw268_[int(6)] = iife138_(d_1538_v10_)
                    nw268_[int(7)] = d_1538_v10_
                    nw268_[int(8)] = d_1538_v10_
                    nw268_[int(9)] = d_1538_v10_
                    nw268_[int(10)] = d_1538_v10_
                    nw268_[int(11)] = (d_1538_v10_ if p1 else d_1538_v10_)
                    nw268_[int(12)] = d_1538_v10_
                    nw268_[int(13)] = d_1538_v10_
                    nw268_[int(14)] = d_1538_v10_
                    def iife144_(_pat_let50_0):
                        def iife145_(d_1549_dt__update__tmp_h2_):
                            def iife146_(_pat_let51_0):
                                def iife147_(d_1550_dt__update_hcf10_h1_):
                                    def iife148_(_pat_let52_0):
                                        def iife149_(d_1551_dt__update_hcf11_h0_):
                                            return D2_DC6(d_1550_dt__update_hcf10_h1_, d_1551_dt__update_hcf11_h0_, (d_1549_dt__update__tmp_h2_).cf12)
                                        return iife149_(_pat_let52_0)
                                    return iife148_(pat_let_tv45_)
                                return iife147_(_pat_let51_0)
                            return iife146_(pat_let_tv44_)
                        return iife145_(_pat_let50_0)
                    nw268_[int(15)] = iife144_(d_1538_v10_)
                    nw268_[int(16)] = d_1538_v10_
                    nw268_[int(17)] = default__.fm7(globalState)
                    nw268_[int(18)] = d_1538_v10_
                    nw268_[int(19)] = d_1538_v10_
                    nw268_[int(20)] = default__.fm7(globalState)
                    nw268_[int(21)] = d_1538_v10_
                    nw268_[int(22)] = d_1538_v10_
                    nw268_[int(23)] = d_1538_v10_
                    nw268_[int(24)] = d_1538_v10_
                    d_1541_v13_ = nw268_
                    d_1541_v13_ = d_1541_v13_
                    rhs325_ = not(False)
                    rhs326_ = p2
                    r0 = rhs325_
                    r0 = rhs326_
                    d_1552_v14_: C0
                    nw269_ = C0()
                    nw269_.ctor__(p1)
                    d_1552_v14_ = nw269_
                    d_1552_v14_ = d_1552_v14_
                    d_1553_v15_: D2
                    d_1553_v15_ = D2_DC5(d_1527_v0_)
                    source19_ = d_1553_v15_
                    if source19_.is_DC6:
                        d_1554___mcc_h0_ = source19_.cf10
                        d_1555___mcc_h1_ = source19_.cf11
                        d_1556___mcc_h2_ = source19_.cf12
                        d_1557_cf12_ = d_1556___mcc_h2_
                        d_1558_cf11_ = d_1555___mcc_h1_
                        d_1559_cf10_ = d_1554___mcc_h0_
                        (d_1552_v14_).f10 = (d_1557_cf12_) not in (((self).f9) + ((self).f9))
                        d_1531_v4_ = d_1531_v4_
                        index286_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        (d_1527_v0_)[index286_] = (len(_dafny.SeqWithoutIsStrInference([d_1552_v14_, d_1552_v14_]))) + ((d_1557_cf12_) - ((self).f6))
                        (globalState).f1 = (self).f6
                    elif source19_.is_DC7:
                        d_1560___mcc_h3_ = source19_.cf13
                        d_1561___mcc_h4_ = source19_.cf14
                        d_1562___mcc_h5_ = source19_.cf15
                        d_1563___mcc_h6_ = source19_.cf16
                        d_1564_cf16_ = d_1563___mcc_h6_
                        d_1565_cf15_ = d_1562___mcc_h5_
                        d_1566_cf14_ = d_1561___mcc_h4_
                        d_1567_cf13_ = d_1560___mcc_h3_
                        d_1527_v0_ = d_1527_v0_
                        d_1568_v16_: _dafny.Map
                        d_1568_v16_ = _dafny.Map({(self).f6: len(d_1531_v4_)})
                        d_1569_v17_: _dafny.Map
                        d_1569_v17_ = _dafny.Map({self.f7: (d_1568_v16_) | (d_1568_v16_)})
                        d_1570_v18_: _dafny.Map
                        d_1570_v18_ = _dafny.Map({d_1552_v14_.f10: not(False)})
                        index287_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        rhs327_ = len((d_1570_v18_) | (_dafny.Map({d_1552_v14_.f10: p2})))
                        rhs328_ = d_1567_cf13_
                        rhs329_ = d_1539_v11_
                        rhs330_ = -312
                        rhs331_ = d_1569_v17_
                        lhs225_ = self
                        lhs226_ = d_1552_v14_
                        lhs227_ = d_1527_v0_
                        lhs228_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        lhs225_.f7 = rhs327_
                        lhs226_.f10 = rhs328_
                        d_1539_v11_ = rhs329_
                        lhs227_[lhs228_] = rhs330_
                        d_1569_v17_ = rhs331_
                        rhs332_ = p2
                        rhs333_ = p2
                        lhs229_ = d_1552_v14_
                        lhs230_ = d_1552_v14_
                        lhs229_.f10 = rhs332_
                        lhs230_.f10 = rhs333_
                        d_1571_v20_: _dafny.Set
                        d_1571_v20_ = _dafny.Set({d_1531_v4_})
                        def iife150_():
                            coll44_ = _dafny.Set()
                            compr_44_: D2
                            for compr_44_ in (_dafny.Map({d_1538_v10_: d_1567_cf13_})).keys.Elements:
                                d_1572_v19_: D2 = compr_44_
                                if (d_1572_v19_) in (_dafny.Map({d_1538_v10_: d_1567_cf13_})):
                                    coll44_ = coll44_.union(_dafny.Set([d_1572_v19_]))
                            return _dafny.Set(coll44_)
                        d_1567_cf13_ = not((default__.fm2((self).f9, len(iife150_()
                        ), self.f7, d_1552_v14_.f10, globalState)) in (d_1571_v20_))
                    elif source19_.is_DC5:
                        d_1573___mcc_h7_ = source19_.cf9
                        d_1574_cf9_ = d_1573___mcc_h7_
                        d_1575_v21_: _dafny.Map
                        d_1575_v21_ = _dafny.Map({(d_1534_v6_).ispropersubset(d_1534_v6_): d_1527_v0_})
                        d_1575_v21_ = (d_1575_v21_).set(d_1552_v14_.f10, d_1574_cf9_)
                        d_1528_v1_ = default__.fm8((d_1531_v4_) + (_dafny.SeqWithoutIsStrInference([d_1528_v1_ for d_1576_i3_ in range(default__.abs(200))])), d_1528_v1_, d_1552_v14_.f10, (self).fm4(globalState), globalState)
                        d_1577_v22_: _dafny.Seq
                        d_1577_v22_ = _dafny.SeqWithoutIsStrInference([(d_1538_v10_).cf11])
                        d_1578_v23_: _dafny.Array
                        nw270_ = _dafny.Array(None, 7)
                        nw270_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1528_v1_ for d_1579_i4_ in range(default__.abs(81))])
                        nw270_[int(1)] = (d_1531_v4_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f9])), len(d_1531_v4_)), d_1528_v1_)
                        nw270_[int(2)] = d_1531_v4_
                        nw270_[int(3)] = d_1531_v4_
                        nw270_[int(4)] = d_1531_v4_
                        nw270_[int(5)] = (d_1577_v22_)[default__.safeIndex(self.f7, len(d_1577_v22_))]
                        nw270_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfyudb"))
                        d_1578_v23_ = nw270_
                        d_1580_v24_: _dafny.Map
                        d_1580_v24_ = _dafny.Map({(d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))]: d_1578_v23_})
                        d_1580_v24_ = (d_1580_v24_).set((self).f6, d_1578_v23_)
                        d_1581_v25_: _dafny.Seq
                        d_1581_v25_ = _dafny.SeqWithoutIsStrInference([(d_1552_v14_ if d_1552_v14_.f10 else d_1552_v14_), d_1552_v14_, d_1552_v14_])
                        d_1582_v26_: _dafny.MultiSet
                        d_1582_v26_ = _dafny.MultiSet([d_1552_v14_.f10, d_1552_v14_.f10])
                        index288_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        rhs334_ = (d_1581_v25_)[default__.safeIndex((d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))], len(d_1581_v25_))]
                        rhs335_ = (default__.safeModuloInt((d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))], (self).f6)) * ((d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))])
                        rhs336_ = _dafny.SeqWithoutIsStrInference([d_1531_v4_, _dafny.SeqWithoutIsStrInference([d_1528_v1_ for d_1583_i5_ in range(default__.abs(387))])])
                        rhs337_ = d_1528_v1_
                        rhs338_ = ((self).f6) == (((d_1582_v26_).cardinality) * ((self).f6))
                        lhs231_ = d_1527_v0_
                        lhs232_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        d_1552_v14_ = rhs334_
                        lhs231_[lhs232_] = rhs335_
                        d_1577_v22_ = rhs336_
                        d_1528_v1_ = rhs337_
                        r0 = rhs338_
                    elif True:
                        d_1584___mcc_h8_ = source19_.cf17
                        d_1585_cf17_ = d_1584___mcc_h8_
                        index289_ = default__.safeIndex(491, (d_1527_v0_).length(0))
                        (d_1527_v0_)[index289_] = (self).f6
                        (globalState).f1 = self.f7
                        d_1586_v27_: _dafny.Set
                        d_1586_v27_ = _dafny.Set({D2_DC5(d_1527_v0_), d_1553_v15_, d_1553_v15_, d_1553_v15_})
                        d_1586_v27_ = (d_1586_v27_ if d_1552_v14_.f10 else d_1586_v27_)
                        d_1539_v11_ = d_1539_v11_
                    pass
            pass
        d_1588_v28_: _dafny.Map
        d_1588_v28_ = _dafny.Map({len(d_1531_v4_): (d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))]})
        d_1589_v29_: _dafny.Set
        d_1589_v29_ = _dafny.Set({p1})
        d_1590_v30_: D0
        d_1590_v30_ = D0_DC0(d_1589_v29_)
        d_1591_v31_: _dafny.Map
        d_1591_v31_ = _dafny.Map({d_1590_v30_: p1})
        d_1592_v32_: D1
        d_1592_v32_ = D1_DC2(d_1591_v31_)
        d_1593_v33_: D2
        d_1593_v33_ = D2_DC6(d_1592_v32_, d_1531_v4_, 229)
        d_1594_v34_: _dafny.Map
        d_1594_v34_ = _dafny.Map({((d_1588_v28_)[(self).f6] if ((self).f6) in (d_1588_v28_) else len(default__.fm9(len(d_1531_v4_), globalState))): D2_DC8(d_1593_v33_)})
        d_1595_v35_: _dafny.MultiSet
        d_1595_v35_ = _dafny.MultiSet([p1])
        d_1596_v36_: _dafny.Seq
        d_1596_v36_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        d_1597_v37_: D2
        d_1597_v37_ = D2_DC8(D2_DC8(D2_DC8(d_1593_v33_)))
        d_1594_v34_ = (d_1594_v34_).set((0) - (default__.safeModuloInt(((d_1595_v35_)[(d_1596_v36_)[default__.safeIndex((d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))], len(d_1596_v36_))]] if ((d_1596_v36_)[default__.safeIndex((d_1527_v0_)[default__.safeIndex(491, (d_1527_v0_).length(0))], len(d_1596_v36_))]) in (d_1595_v35_) else self.f7), len(d_1531_v4_))), d_1597_v37_)
        r0 = p1
        d_1598_v38_: _dafny.Array
        def lambda86_(d_1599_v29_):
            def lambda87_(d_1600_i6_):
                return D0_DC0(d_1599_v29_)

            return lambda87_

        init49_ = lambda86_(d_1589_v29_)
        nw271_ = _dafny.Array(None, 20)
        for i0_49_ in range(nw271_.length(0)):
            nw271_[i0_49_] = init49_(i0_49_)
        d_1598_v38_ = nw271_
        r1 = d_1598_v38_
        return r0, r1

    def m3(self, p0, p1, globalState):
        d_1601_v0_: bool
        d_1601_v0_ = True
        d_1602_v1_: _dafny.Seq
        d_1602_v1_ = _dafny.SeqWithoutIsStrInference([d_1601_v0_])
        d_1603_v2_: _dafny.Map
        d_1603_v2_ = _dafny.Map({(d_1602_v1_).set(default__.safeIndex((self).f6, len(d_1602_v1_)), d_1601_v0_): (self).f6})
        d_1604_v3_: _dafny.MultiSet
        d_1604_v3_ = _dafny.MultiSet([(self).f6])
        (self).f7 = ((d_1603_v2_)[(d_1602_v1_) + (d_1602_v1_)] if ((d_1602_v1_) + (d_1602_v1_)) in (d_1603_v2_) else default__.safeDivisionInt(((d_1604_v3_).set(957, default__.abs((self).f6))).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adkmbscvj")))))
        d_1605_i0_: int
        d_1605_i0_ = 0
        with _dafny.label("11"):
            while d_1601_v0_:
                with _dafny.c_label("11"):
                    if (d_1605_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1605_i0_ = (d_1605_i0_) + (1)
                    d_1606_v4_: _dafny.Seq
                    d_1606_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvswxolf"))
                    d_1606_v4_ = default__.fm2((self).f9, (self.f7 if (d_1602_v1_)[default__.safeIndex(self.f7, len(d_1602_v1_))] else (self).f6), (self).f6, not(d_1601_v0_), globalState)
                    d_1607_v5_: C0
                    nw272_ = C0()
                    nw272_.ctor__(d_1601_v0_)
                    d_1607_v5_ = nw272_
                    d_1608_v6_: D0
                    d_1608_v6_ = D0_DC1(True, p0, (p1) == (p1))
                    source20_ = d_1608_v6_
                    if source20_.is_DC1:
                        d_1609___mcc_h0_ = source20_.cf1
                        d_1610___mcc_h1_ = source20_.cf2
                        d_1611___mcc_h2_ = source20_.cf3
                        d_1612_cf3_ = d_1611___mcc_h2_
                        d_1613_cf2_ = d_1610___mcc_h1_
                        d_1614_cf1_ = d_1609___mcc_h0_
                        d_1615_v7_: _dafny.Array
                        nw273_ = _dafny.Array(None, 1)
                        d_1615_v7_ = nw273_
                        index290_ = default__.safeIndex(928, (d_1615_v7_).length(0))
                        (d_1615_v7_)[index290_] = d_1607_v5_
                        d_1616_v8_: C0
                        d_1616_v8_ = d_1607_v5_
                        d_1617_v9_: _dafny.Seq
                        d_1617_v9_ = _dafny.SeqWithoutIsStrInference([(d_1616_v8_), d_1607_v5_, d_1607_v5_, d_1607_v5_])
                        index291_ = default__.safeIndex(928, (d_1615_v7_).length(0))
                        (d_1615_v7_)[index291_] = (d_1617_v9_)[default__.safeIndex(len(default__.fm10(d_1607_v5_.f10, d_1601_v0_, d_1602_v1_, True, globalState)), len(d_1617_v9_))]
                        index292_ = default__.safeIndex(159, (p0).length(0))
                        (p0)[index292_] = d_1601_v0_
                        index293_ = default__.safeIndex(159, (p0).length(0))
                        (p0)[index293_] = d_1612_cf3_
                        rhs339_ = d_1606_v4_
                        rhs340_ = d_1607_v5_.f10
                        rhs341_ = (d_1604_v3_) != ((d_1604_v3_) - (d_1604_v3_))
                        lhs233_ = d_1607_v5_
                        d_1606_v4_ = rhs339_
                        d_1614_cf1_ = rhs340_
                        lhs233_.f10 = rhs341_
                        d_1618_v10_: _dafny.Array
                        nw274_ = _dafny.Array(int(0), 22)
                        d_1618_v10_ = nw274_
                        index294_ = default__.safeIndex(393, ((self).f8).length(0))
                        ((self).f8)[index294_] = d_1618_v10_
                        index295_ = default__.safeIndex(393, ((self).f8).length(0))
                        ((self).f8)[index295_] = d_1618_v10_
                    elif True:
                        d_1619___mcc_h3_ = source20_.cf0
                        d_1620_cf0_ = d_1619___mcc_h3_
                        (globalState).f1 = p1
                        d_1606_v4_ = ((d_1606_v4_) + (d_1606_v4_)) + (d_1606_v4_)
                        d_1621_v11_: _dafny.Array
                        def lambda88_(d_1622_i1_):
                            return default__.safeDivisionInt(d_1622_i1_, -444)

                        init50_ = lambda88_
                        nw275_ = _dafny.Array(None, 9)
                        for i0_50_ in range(nw275_.length(0)):
                            nw275_[i0_50_] = init50_(i0_50_)
                        d_1621_v11_ = nw275_
                        d_1623_v12_: T0
                        nw276_ = C2()
                        nw276_.ctor__(len(d_1602_v1_), len(_dafny.SeqWithoutIsStrInference([(self).f6 for d_1624_i2_ in range(default__.abs(397))])))
                        d_1623_v12_ = nw276_
                        d_1625_v13_: _dafny.Map
                        d_1625_v13_ = _dafny.Map({d_1623_v12_: True})
                        d_1626_v14_: _dafny.Map
                        d_1626_v14_ = _dafny.Map({d_1601_v0_: len(d_1625_v13_)})
                        index296_ = default__.safeIndex(176, (d_1621_v11_).length(0))
                        (d_1621_v11_)[index296_] = len(d_1626_v14_)
                        index297_ = default__.safeIndex(213, (p0).length(0))
                        (p0)[index297_] = d_1601_v0_
                        index298_ = default__.safeIndex(176, (d_1621_v11_).length(0))
                        index299_ = default__.safeIndex(213, (p0).length(0))
                        rhs342_ = len(d_1606_v4_)
                        rhs343_ = d_1607_v5_.f10
                        lhs234_ = d_1621_v11_
                        lhs235_ = default__.safeIndex(176, (d_1621_v11_).length(0))
                        lhs236_ = p0
                        lhs237_ = default__.safeIndex(213, (p0).length(0))
                        lhs234_[lhs235_] = rhs342_
                        lhs236_[lhs237_] = rhs343_
                        d_1627_v15_: C0
                        nw277_ = C0()
                        nw277_.ctor__(not(d_1607_v5_.f10))
                        d_1627_v15_ = nw277_
                    d_1628_v16_: _dafny.Set
                    d_1628_v16_ = _dafny.Set({True, d_1607_v5_.f10, (-182) == (self.f7)})
                    d_1629_v17_: _dafny.Array
                    def lambda89_(d_1630_i3_):
                        return (d_1630_i3_) + ((self).f6)

                    init51_ = lambda89_
                    nw278_ = _dafny.Array(None, 7)
                    for i0_51_ in range(nw278_.length(0)):
                        nw278_[i0_51_] = init51_(i0_51_)
                    d_1629_v17_ = nw278_
                    index300_ = default__.safeIndex(118, (d_1629_v17_).length(0))
                    (d_1629_v17_)[index300_] = self.f7
                    index301_ = default__.safeIndex(118, (d_1629_v17_).length(0))
                    rhs344_ = d_1628_v16_
                    rhs345_ = p1
                    lhs238_ = d_1629_v17_
                    lhs239_ = default__.safeIndex(118, (d_1629_v17_).length(0))
                    d_1628_v16_ = rhs344_
                    lhs238_[lhs239_] = rhs345_
                    pass
            pass
        d_1631_v18_: C0
        nw279_ = C0()
        nw279_.ctor__(False)
        d_1631_v18_ = nw279_
        d_1632_v19_: _dafny.Map
        d_1632_v19_ = _dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1631_v18_, d_1631_v18_]))]))).cardinality]))})
        d_1633_v20_: _dafny.Seq
        d_1633_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmtta"))
        d_1634_v21_: _dafny.Map
        d_1634_v21_ = _dafny.Map({d_1631_v18_.f10: p1})
        d_1635_v22_: _dafny.MultiSet
        d_1635_v22_ = _dafny.MultiSet([d_1601_v0_])
        d_1636_v23_: D6
        d_1636_v23_ = D6_DC14(self.f7, d_1635_v22_, (_dafny.MultiSet([self.f7])).cardinality)
        d_1637_v24_: _dafny.Array
        nw280_ = _dafny.Array(None, 26)
        nw280_[int(0)] = (self).f9
        nw280_[int(1)] = _dafny.SeqWithoutIsStrInference([(self).f6, ((self).f9)[default__.safeIndex(len(d_1632_v19_), len((self).f9))], (self).f6])
        nw280_[int(2)] = _dafny.SeqWithoutIsStrInference([p1])
        nw280_[int(3)] = _dafny.SeqWithoutIsStrInference([(self).f6 for d_1638_i5_ in range(default__.abs(423))])
        nw280_[int(4)] = ((self).f9).set(default__.safeIndex(self.f7, len((self).f9)), len(d_1633_v20_))
        nw280_[int(5)] = ((self).f9) + (_dafny.SeqWithoutIsStrInference([self.f7, -262]))
        nw280_[int(6)] = (self).f9
        nw280_[int(7)] = (self).f9
        nw280_[int(8)] = (self).f9
        nw280_[int(9)] = (self).f9
        nw280_[int(10)] = default__.fm5(d_1631_v18_.f10, ((d_1634_v21_)[d_1601_v0_] if (d_1601_v0_) in (d_1634_v21_) else len(_dafny.Map({d_1601_v0_: p1}))), default__.fm1(d_1631_v18_.f10, globalState), d_1601_v0_, globalState)
        nw280_[int(11)] = ((self).f9).set(default__.safeIndex(self.f7, len((self).f9)), (d_1636_v23_).cf23)
        nw280_[int(12)] = (self).f9
        nw280_[int(13)] = (self).f9
        nw280_[int(14)] = default__.fm5(True, -647, p1, d_1631_v18_.f10, globalState)
        nw280_[int(15)] = (self).f9
        nw280_[int(16)] = (self).f9
        nw280_[int(17)] = (self).f9
        nw280_[int(18)] = _dafny.SeqWithoutIsStrInference([(0) - (default__.fm1(False, globalState)), self.f7, (self).f6, (self).f6])
        nw280_[int(19)] = (self).f9
        nw280_[int(20)] = (self).f9
        nw280_[int(21)] = default__.fm35(self.f7, self.f7, globalState)
        nw280_[int(22)] = (self).f9
        nw280_[int(23)] = (self).f9
        nw280_[int(24)] = (default__.fm5(d_1601_v0_, p1, self.f7, d_1601_v0_, globalState)) + ((self).f9)
        nw280_[int(25)] = (self).f9
        d_1637_v24_ = nw280_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1637_v24_).length(0)):
            d_1639_i4_: int = guard_loop_4_
            if (True) and (((0) <= (d_1639_i4_)) and ((d_1639_i4_) < ((d_1637_v24_).length(0)))):
                (d_1637_v24_)[(d_1639_i4_)] = (_dafny.SeqWithoutIsStrInference([len(d_1634_v21_) for d_1640_i6_ in range(default__.abs(835))])) + ((self).f9)
        d_1641_v25_: str
        d_1641_v25_ = _dafny.CodePoint('t')
        d_1642_v26_: C1
        nw281_ = C1()
        nw281_.ctor__(default__.fm8(d_1633_v20_, d_1641_v25_, d_1631_v18_.f10, not(d_1631_v18_.f10), globalState), len(d_1633_v20_), self.f7)
        d_1642_v26_ = nw281_
        d_1642_v26_ = d_1642_v26_
        d_1643_v27_: C3
        nw282_ = C3()
        nw282_.ctor__(default__.safeModuloInt(len(default__.fm2(_dafny.SeqWithoutIsStrInference([24]), self.f7, (0) - (self.f7), d_1601_v0_, globalState)), self.f7), 876)
        d_1643_v27_ = nw282_
        d_1644_v28_: C3
        nw283_ = C3()
        nw283_.ctor__(default__.safeDivisionInt((d_1643_v27_).fm15(self.f7, p1, globalState), (self).f6), ((self).f9)[default__.safeIndex(p1, len((self).f9))])
        d_1644_v28_ = nw283_

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
