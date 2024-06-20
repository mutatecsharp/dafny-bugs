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
        return -290

    @staticmethod
    def fm1(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])).Elements:
                d_2_v0_: str = compr_0_
                if (d_2_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')])):
                    coll0_[d_2_v0_] = -708
            return _dafny.Map(coll0_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ln"))), 589, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_0_i0_ in range(default__.abs(873))]))])) + (_dafny.SeqWithoutIsStrInference([-916 for d_1_i1_ in range(default__.abs(46))]))) + ((_dafny.SeqWithoutIsStrInference([len(iife0_()
) for d_3_i2_ in range(default__.abs(2))])) + ((D5_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True, True})) for d_4_i3_ in range(default__.abs(609))]))).cf25))

    @staticmethod
    def fm2(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: str
            for compr_1_ in (_dafny.MultiSet([_dafny.CodePoint('v')])).Elements:
                d_5_v0_: str = compr_1_
                if (d_5_v0_) in (_dafny.MultiSet([_dafny.CodePoint('v')])):
                    coll1_[d_5_v0_] = True
            return _dafny.Map(coll1_)
        return ((D7_DC20(_dafny.MultiSet([len(iife1_()
)]))).cf30) == ((_dafny.MultiSet([-675, (_dafny.MultiSet([429, 870])).cardinality, 495])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-209]))))

    @staticmethod
    def fm3(p0, globalState):
        return D0_DC2(len(_dafny.SeqWithoutIsStrInference([-501])), (_dafny.MultiSet([_dafny.CodePoint('t')])).issubset(_dafny.MultiSet([_dafny.CodePoint('t')])))

    @staticmethod
    def fm4(p0, p1, globalState):
        return D0_DC0((True if True else True))

    @staticmethod
    def fm13(p0, globalState):
        return (_dafny.Set({True})).intersection((_dafny.Set({True, False})).intersection(_dafny.Set({True, False, False, False})))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        if not((True) and (False)):
            return (D4_DC12(_dafny.SeqWithoutIsStrInference([False]))).cf18
        elif True:
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([not(not(not(not(not(not(not(False)))))))]))

    @staticmethod
    def fm16(p0, globalState):
        source0_ = D8_DC27(D8_DC26(not(False)))
        if source0_.is_DC24:
            d_6___mcc_h0_ = source0_.cf36
            d_7_cf36_ = d_6___mcc_h0_
            if d_7_cf36_:
                return D2_DC5(-268, d_7_cf36_, d_7_cf36_)
            elif True:
                return D2_DC5(417, True, not(d_7_cf36_))
        elif source0_.is_DC25:
            d_8___mcc_h1_ = source0_.cf37
            d_9___mcc_h2_ = source0_.cf38
            d_10___mcc_h3_ = source0_.cf39
            d_11___mcc_h4_ = source0_.cf40
            d_12_cf40_ = d_11___mcc_h4_
            d_13_cf39_ = d_10___mcc_h3_
            d_14_cf38_ = d_9___mcc_h2_
            d_15_cf37_ = d_8___mcc_h1_
            return D2_DC5(d_14_cf38_, d_15_cf37_, d_15_cf37_)
        elif source0_.is_DC26:
            d_16___mcc_h5_ = source0_.cf41
            d_17_cf41_ = d_16___mcc_h5_
            return D2_DC5(-893, d_17_cf41_, d_17_cf41_)
        elif source0_.is_DC23:
            d_18___mcc_h6_ = source0_.cf35
            d_19_cf35_ = d_18___mcc_h6_
            return D2_DC5((0) - (len(_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality)}))), True, True)
        elif True:
            d_20___mcc_h7_ = source0_.cf42
            d_21_cf42_ = d_20___mcc_h7_
            return D2_DC5(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})), (D21_DC58(False, len(_dafny.Map({not(not(False)): True})))).cf89, True)

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True]) if not(False) else _dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([True, not(True), False, not(True)])) + (_dafny.SeqWithoutIsStrInference([not(not(False))])))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(837, -677):
                d_22_v0_: int = compr_2_
                if ((837) <= (d_22_v0_)) and ((d_22_v0_) < (-677)):
                    coll2_ = coll2_.union(_dafny.Set([default__.safeDivisionInt(d_22_v0_, 509)]))
            return _dafny.Set(coll2_)
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-800, 437):
                d_23_v1_: int = compr_3_
                if ((-800) <= (d_23_v1_)) and ((d_23_v1_) < (437)):
                    coll3_ = coll3_.union(_dafny.Set([default__.safeDivisionInt(d_23_v1_, 642)]))
            return _dafny.Set(coll3_)
        return D0_DC0((iife2_()
).ispropersubset(iife3_()
))

    @staticmethod
    def fm20(p0, p1, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))) | (_dafny.MultiSet([False]))) - ((_dafny.MultiSet([not(not(True)), True])).intersection(_dafny.MultiSet([False])))

    @staticmethod
    def fm21(p0, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxqour"))

    @staticmethod
    def fm22(p0, globalState):
        return (_dafny.MultiSet([False, True])) | (_dafny.MultiSet([not(False), True]))

    @staticmethod
    def fm23(p0, p1, globalState):
        return D2_DC4(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D22_DC60(False, False, not(True))).cf93])]))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return D2_DC5((0) - ((0) - (default__.safeModuloInt((0) - (len(_dafny.Map({True: False}))), len(_dafny.Map({218: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gikgwt")))}))))), True, True)

    @staticmethod
    def fm25(p0, globalState):
        if (815) != (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_24_i0_ in range(default__.abs(16))]))):
            return D3_DC9(33, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkarebnxj"))))
        elif True:
            return D3_DC9((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwdkvrqt"))), 992])).cardinality, 731)

    @staticmethod
    def fm27(globalState):
        return ((_dafny.Set({False})) | (_dafny.Set({True, False}))).intersection((_dafny.Set({True})))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(618, -394):
                d_25_v0_: int = compr_4_
                if ((618) <= (d_25_v0_)) and ((d_25_v0_) < (-394)):
                    coll4_[default__.safeDivisionInt(d_25_v0_, len(_dafny.Map({D12_DC37(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), True): _dafny.Map({True: 930})})))] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eskdxn"))), 383, 492])
            return _dafny.Map(coll4_)
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([len(iife4_()
        ), -861])) + (_dafny.SeqWithoutIsStrInference([789, 210]))) + (_dafny.SeqWithoutIsStrInference([(0) - (-831)])))

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        source1_ = D23_DC64()
        if source1_.is_DC64:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txelk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj")))
        elif True:
            d_26___mcc_h0_ = source1_.cf98
            d_27_cf98_ = d_26___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnvlqggk"))

    @staticmethod
    def fm30(p0, p1, globalState):
        source2_ = _dafny.CodePoint('v')
        d_28___mcc_h0_ = source2_
        d_29_cf44_ = d_28___mcc_h0_
        return D6_DC19(D6_DC18((0) - (len(_dafny.SeqWithoutIsStrInference([867, 843])))))

    @staticmethod
    def fm31(p0, globalState):
        return D0_DC2(230, (-160) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lpmdnlj")))))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(False)]))) + ((_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([True, not(False), not(True)])))

    @staticmethod
    def fm35(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_30_i0_ in range(default__.abs(691))])

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([False, (D17_DC49(False)).cf73])})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([not(not(False))]), _dafny.SeqWithoutIsStrInference([True, True, True]), _dafny.SeqWithoutIsStrInference([True, False])}))) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True, False, False, True]), _dafny.SeqWithoutIsStrInference([not(True)]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(True), True, not(not(True))])}))

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return D4_DC12(_dafny.SeqWithoutIsStrInference([True, True]))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_31_i0_ in range(default__.abs(577))])

    @staticmethod
    def fm41(p0, p1, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm42(p0, p1, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])).Elements:
                d_32_v0_: str = compr_5_
                if (d_32_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')])):
                    coll5_[d_32_v0_] = 112
            return _dafny.Map(coll5_)
        return (_dafny.Map({_dafny.CodePoint('a'): _dafny.MultiSet((D5_DC15(_dafny.SeqWithoutIsStrInference([-461, len(_dafny.SeqWithoutIsStrInference([iife5_()
])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pu"))), 485]))).cf25)})) | (_dafny.Map({_dafny.CodePoint('y'): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([303]))}))

    @staticmethod
    def fm43(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, True])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([False, not(False)])))

    @staticmethod
    def fm44(p0, p1, globalState):
        if True:
            return D2_DC5((0) - (len(_dafny.Map({False: True}))), False, False)
        elif True:
            return D2_DC5(-672, False, False)

    @staticmethod
    def fm45(p0, p1, p2, globalState):
        return ((_dafny.Map({False: len(_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({True: 310})), 244])).cardinality: 464}))}) if True else _dafny.Map({False: (_dafny.MultiSet([False, False])).cardinality}))) | (_dafny.Map({True: -60}))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        if not (not(True)) or (False):
            return D6_DC18(779)
        elif True:
            return D6_DC18(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(False), True, True]), _dafny.SeqWithoutIsStrInference([True, False])])))

    @staticmethod
    def fm47(globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(30, -859):
                d_34_v0_: int = compr_6_
                if ((30) <= (d_34_v0_)) and ((d_34_v0_) < (-859)):
                    coll6_[(d_34_v0_) - ((0) - (len(_dafny.SeqWithoutIsStrInference([-906, 953]))))] = len(_dafny.Map({False: False}))
            return _dafny.Map(coll6_)
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryegaxv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpurpehd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrthtypa"))})).Elements:
                d_35_v1_: _dafny.Seq = compr_7_
                if (d_35_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryegaxv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpurpehd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrthtypa"))})):
                    coll7_[d_35_v1_] = _dafny.SeqWithoutIsStrInference([854, (_dafny.MultiSet([998])).cardinality])
            return _dafny.Map(coll7_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_33_i0_ in range(default__.abs(247))]): _dafny.SeqWithoutIsStrInference([562, 654, len(iife6_()
        )])})) | (iife7_()
        )

    @staticmethod
    def fm48(globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-656, 360):
                d_36_v0_: int = compr_8_
                if ((-656) <= (d_36_v0_)) and ((d_36_v0_) < (360)):
                    coll8_[(d_36_v0_) + (-774)] = _dafny.Map({False: len(_dafny.Set({_dafny.CodePoint('k')}))})
            return _dafny.Map(coll8_)
        return D11_DC30(iife8_()
)

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        if not((True if False else not(not(True)))):
            return _dafny.Map({139: not(not(False))})
        elif True:
            return (_dafny.Map({len(_dafny.Map({False: 377})): False})) | (_dafny.Map({409: False}))

    @staticmethod
    def fm50(p0, p1, globalState):
        return D4_DC14(True)

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(540, 589):
                    d_37_v0_: int = compr_11_
                    if ((540) <= (d_37_v0_)) and ((d_37_v0_) < (589)):
                        coll11_[(d_37_v0_) * (394)] = _dafny.CodePoint('a')
                return _dafny.Map(coll11_)
            coll9_ = _dafny.Set()
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(540, 589):
                    d_37_v0_: int = compr_10_
                    if ((540) <= (d_37_v0_)) and ((d_37_v0_) < (589)):
                        coll10_[(d_37_v0_) * (394)] = _dafny.CodePoint('a')
                return _dafny.Map(coll10_)
            compr_9_: int
            for compr_9_ in (iife10_()
            ).keys.Elements:
                d_38_v1_: int = compr_9_
                if (d_38_v1_) in (iife11_()
                ):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_38_v1_, 130)]))
            return _dafny.Set(coll9_)
        return ((_dafny.Map({False: iife9_()
        })) | (_dafny.Map({True: _dafny.Set({717})}))) | (_dafny.Map({True: _dafny.Set({len(_dafny.Set({True}))})}))

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([617])).Elements:
                d_39_v0_: int = compr_12_
                if (d_39_v0_) in (_dafny.SeqWithoutIsStrInference([617])):
                    coll12_[default__.safeDivisionInt(d_39_v0_, 96)] = False
            return _dafny.Map(coll12_)
        def iife13_():
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(600, 102):
                    d_40_v1_: int = compr_17_
                    if ((600) <= (d_40_v1_)) and ((d_40_v1_) < (102)):
                        coll17_[(d_40_v1_) * ((_dafny.MultiSet([True, True])).cardinality)] = -401
                return _dafny.Map(coll17_)
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in (_dafny.SeqWithoutIsStrInference([503])).Elements:
                    d_41_v2_: int = compr_18_
                    if (d_41_v2_) in (_dafny.SeqWithoutIsStrInference([503])):
                        coll18_[(d_41_v2_) + (-371)] = False
                return _dafny.Map(coll18_)
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(915, -501):
                    d_42_v3_: int = compr_19_
                    if ((915) <= (d_42_v3_)) and ((d_42_v3_) < (-501)):
                        coll19_[default__.safeModuloInt(d_42_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swbdrpg"))))] = False
                return _dafny.Map(coll19_)
            coll13_ = _dafny.Set()
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(600, 102):
                    d_40_v1_: int = compr_14_
                    if ((600) <= (d_40_v1_)) and ((d_40_v1_) < (102)):
                        coll14_[(d_40_v1_) * ((_dafny.MultiSet([True, True])).cardinality)] = -401
                return _dafny.Map(coll14_)
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in (_dafny.SeqWithoutIsStrInference([503])).Elements:
                    d_41_v2_: int = compr_15_
                    if (d_41_v2_) in (_dafny.SeqWithoutIsStrInference([503])):
                        coll15_[(d_41_v2_) + (-371)] = False
                return _dafny.Map(coll15_)
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(915, -501):
                    d_42_v3_: int = compr_16_
                    if ((915) <= (d_42_v3_)) and ((d_42_v3_) < (-501)):
                        coll16_[default__.safeModuloInt(d_42_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swbdrpg"))))] = False
                return _dafny.Map(coll16_)
            compr_13_: _dafny.Map
            for compr_13_ in (_dafny.Set({_dafny.Map({(0) - ((_dafny.MultiSet([len(iife14_()
            )])).cardinality): True}), _dafny.Map({443: True}), iife15_()
            , iife16_()
            })).Elements:
                d_43_v4_: _dafny.Map = compr_13_
                if (d_43_v4_) in (_dafny.Set({_dafny.Map({(0) - ((_dafny.MultiSet([len(iife17_()
                )])).cardinality): True}), _dafny.Map({443: True}), iife18_()
                , iife19_()
                })):
                    coll13_ = coll13_.union(_dafny.Set([d_43_v4_]))
            return _dafny.Set(coll13_)
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('m'), _dafny.CodePoint('y')]))})).Elements:
                d_44_v5_: int = compr_20_
                if (d_44_v5_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r'), _dafny.CodePoint('m'), _dafny.CodePoint('y')]))})):
                    coll20_[(d_44_v5_) * (512)] = not(not(False))
            return _dafny.Map(coll20_)
        return ((_dafny.Set({_dafny.Map({len(_dafny.Set({len(_dafny.Map({-114: 398}))})): True}), iife12_()
        , _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): not(True)})})) | (iife13_()
        )) - (_dafny.Set({iife20_()
        }))

    @staticmethod
    def fm53(p0, p1, globalState):
        return D2_DC4(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))

    @staticmethod
    def fm54(p0, globalState):
        if not(not((True if True else not(True)))):
            return D3_DC8((D3_DC8(_dafny.Map({not(True): False}))).cf11)
        elif True:
            return D3_DC8(_dafny.Map({not(not(False)): True}))

    @staticmethod
    def fm55(globalState):
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: D14
            for compr_21_ in (_dafny.Set({D14_DC41(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osq")), 762), D14_DC41(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnwyed")), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skj"))))})).Elements:
                d_45_v0_: D14 = compr_21_
                if (d_45_v0_) in (_dafny.Set({D14_DC41(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osq")), 762), D14_DC41(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnwyed")), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skj"))))})):
                    coll21_ = coll21_.union(_dafny.Set([d_45_v0_]))
            return _dafny.Set(coll21_)
        return D4_DC13(not(True), (-318) == (len(_dafny.Set({True, not(not(not(True))), False, not(not(False)), True}))), len(iife21_()
), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnji"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvs")))), not(not(True)))

    @staticmethod
    def fm56(p0, globalState):
        if False:
            return _dafny.MultiSet([364, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpnljp")))])
        elif True:
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([738]))

    @staticmethod
    def fm57(p0, p1, p2, globalState):
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(228, 940):
                d_46_v0_: int = compr_22_
                if ((228) <= (d_46_v0_)) and ((d_46_v0_) < (940)):
                    coll22_[(d_46_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "huket"))))] = 757
            return _dafny.Map(coll22_)
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(846, 847):
                d_47_v1_: int = compr_23_
                if ((846) <= (d_47_v1_)) and ((d_47_v1_) < (847)):
                    coll23_[default__.safeDivisionInt(d_47_v1_, len(_dafny.Map({False: 635})))] = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([False, True]))))
            return _dafny.Map(coll23_)
        return _dafny.MultiSet([iife22_()
        , (iife23_()
        ) | (_dafny.Map({-619: -169}))])

    @staticmethod
    def fm58(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_48_i0_ in range(default__.abs(650))])])

    @staticmethod
    def fm59(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({True: True})) | (_dafny.Map({False: False}))) | ((_dafny.Map({False: False})) | (_dafny.Map({False: (D2_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulwymqieg"))), True, True)).cf7})))

    @staticmethod
    def fm60(p0, globalState):
        source3_ = D12_DC35(False)
        if source3_.is_DC35:
            d_49___mcc_h0_ = source3_.cf54
            d_50_cf54_ = d_49___mcc_h0_
            return _dafny.Map({d_50_cf54_: _dafny.Map({d_50_cf54_: d_50_cf54_})})
        elif source3_.is_DC36:
            d_51___mcc_h1_ = source3_.cf55
            d_52___mcc_h2_ = source3_.cf56
            d_53___mcc_h3_ = source3_.cf57
            d_54___mcc_h4_ = source3_.cf58
            d_55_cf58_ = d_54___mcc_h4_
            d_56_cf57_ = d_53___mcc_h3_
            d_57_cf56_ = d_52___mcc_h2_
            d_58_cf55_ = d_51___mcc_h1_
            return (_dafny.Map({not(False): _dafny.Map({True: not(False)})})) | (_dafny.Map({True: _dafny.Map({True: True})}))
        elif source3_.is_DC37:
            d_59___mcc_h5_ = source3_.cf59
            d_60___mcc_h6_ = source3_.cf60
            d_61_cf60_ = d_60___mcc_h6_
            d_62_cf59_ = d_59___mcc_h5_
            return _dafny.Map({not(d_61_cf60_): _dafny.Map({d_61_cf60_: d_61_cf60_})})
        elif source3_.is_DC34:
            d_63___mcc_h7_ = source3_.cf53
            d_64_cf53_ = d_63___mcc_h7_
            return _dafny.Map({False: _dafny.Map({True: not(False)})})
        elif True:
            d_65___mcc_h8_ = source3_.cf61
            d_66_cf61_ = d_65___mcc_h8_
            return _dafny.Map({True: _dafny.Map({True: True})})

    @staticmethod
    def fm61(globalState):
        def iife24_():
            def iife26_():
                coll26_ = _dafny.Set()
                compr_26_: _dafny.Map
                for compr_26_ in (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({False: not(True)})})).Elements:
                    d_69_v1_: _dafny.Map = compr_26_
                    if (d_69_v1_) in (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({False: not(True)})})):
                        coll26_ = coll26_.union(_dafny.Set([d_69_v1_]))
                return _dafny.Set(coll26_)
            coll24_ = _dafny.Map()
            def iife25_():
                coll25_ = _dafny.Set()
                compr_25_: _dafny.Map
                for compr_25_ in (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({False: not(True)})})).Elements:
                    d_67_v1_: _dafny.Map = compr_25_
                    if (d_67_v1_) in (_dafny.Set({_dafny.Map({True: False}), _dafny.Map({False: not(True)})})):
                        coll25_ = coll25_.union(_dafny.Set([d_67_v1_]))
                return _dafny.Set(coll25_)
            compr_24_: _dafny.Map
            for compr_24_ in ((D27_DC73(iife25_()
)).cf113).Elements:
                d_68_v0_: _dafny.Map = compr_24_
                if (d_68_v0_) in ((D27_DC73(iife26_()
)).cf113):
                    coll24_[d_68_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ax"))
            return _dafny.Map(coll24_)
        return (_dafny.Map({_dafny.Map({True: True}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsgynkkcj"))})) | (iife24_()
        )

    @staticmethod
    def fm62(p0, globalState):
        def iife27_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in _dafny.IntegerRange(348, 296):
                d_70_v0_: int = compr_27_
                if ((348) <= (d_70_v0_)) and ((d_70_v0_) < (296)):
                    coll27_[default__.safeModuloInt(d_70_v0_, 490)] = _dafny.Map({False: 57})
            return _dafny.Map(coll27_)
        return iife27_()
        

    @staticmethod
    def fm63(p0, globalState):
        return D0_DC1()

    @staticmethod
    def fm64(p0, p1, p2, p3, globalState):
        return (_dafny.Set({_dafny.MultiSet([True])})) - (_dafny.Set({_dafny.MultiSet([True, False]), _dafny.MultiSet([True]), _dafny.MultiSet([False, True, True])}))

    @staticmethod
    def fm65(p0, p1, globalState):
        return ((_dafny.Map({_dafny.CodePoint('y'): _dafny.CodePoint('p')})) | (_dafny.Map({_dafny.CodePoint('q'): _dafny.CodePoint('h')}))) | (_dafny.Map({_dafny.CodePoint('b'): _dafny.CodePoint('u')}))

    @staticmethod
    def fm66(globalState):
        return _dafny.Map({54: default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([False, True, True])), len(_dafny.SeqWithoutIsStrInference([911, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality, (_dafny.MultiSet([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bybpefmp")))})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))), len(_dafny.Set({_dafny.CodePoint('x')})), (_dafny.MultiSet([63, 840])).cardinality])).cardinality])))})

    @staticmethod
    def fm67(p0, p1, p2, globalState):
        return D8_DC24(False)

    @staticmethod
    def fm68(p0, globalState):
        def iife28_():
            coll28_ = _dafny.Map()
            compr_28_: str
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('e')])).Elements:
                d_71_v0_: str = compr_28_
                if (d_71_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l'), _dafny.CodePoint('e')])):
                    coll28_[d_71_v0_] = False
            return _dafny.Map(coll28_)
        def iife29_():
            coll29_ = _dafny.Map()
            compr_29_: str
            for compr_29_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_72_i0_ in range(default__.abs(47))])).Elements:
                d_73_v1_: str = compr_29_
                if (d_73_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_72_i0_ in range(default__.abs(47))])):
                    coll29_[d_73_v1_] = False
            return _dafny.Map(coll29_)
        return ((_dafny.Map({_dafny.CodePoint('u'): True})) | (iife28_()
        )) | ((iife29_()
        ) | (_dafny.Map({_dafny.CodePoint('x'): False})))

    @staticmethod
    def fm69(p0, p1, p2, globalState):
        def iife30_():
            coll30_ = _dafny.Set()
            compr_30_: int
            for compr_30_ in _dafny.IntegerRange(-13, 636):
                d_74_v0_: int = compr_30_
                if ((-13) <= (d_74_v0_)) and ((d_74_v0_) < (636)):
                    coll30_ = coll30_.union(_dafny.Set([(d_74_v0_) - (722)]))
            return _dafny.Set(coll30_)
        if (iife30_()
        ).issubset(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-237 for d_75_i0_ in range(default__.abs(144))]))})):
            return D20_DC54((_dafny.CodePoint('f')), _dafny.Map({-555: False}), _dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_76_i1_ in range(default__.abs(578))])}), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjsq")))))
        elif True:
            return D20_DC54(_dafny.CodePoint('x'), _dafny.Map({-840: False}), _dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aquqvkpq"))}), 471)

    @staticmethod
    def fm70(p0, p1, p2, p3, globalState):
        return ((D6_DC17(_dafny.Set({199, len(_dafny.Set({False, True})), 937, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_77_i0_ in range(default__.abs(225))])), 754})) if not(True) else D6_DC17(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))])).cardinality})))).cf27

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: int = int(0)
        r3: bool = False
        d_78_v0_: _dafny.Array
        def lambda0_(d_79_i0_):
            return default__.safeModuloInt(d_79_i0_, len(_dafny.SeqWithoutIsStrInference([515, -90])))

        init0_ = lambda0_
        nw0_ = _dafny.Array(None, 1)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_78_v0_ = nw0_
        d_80_v1_: _dafny.MultiSet
        d_80_v1_ = _dafny.MultiSet([p0, p0, p2])
        d_81_v2_: int
        d_81_v2_ = 788
        index0_ = default__.safeIndex(288, (d_78_v0_).length(0))
        (d_78_v0_)[index0_] = ((d_80_v1_)[p0] if (p0) in (d_80_v1_) else d_81_v2_)
        d_82_v3_: C7
        nw1_ = C7()
        nw1_.ctor__()
        d_82_v3_ = nw1_
        d_83_v4_: _dafny.Seq
        d_83_v4_ = _dafny.SeqWithoutIsStrInference([d_82_v3_])
        d_84_v5_: _dafny.Map
        d_84_v5_ = _dafny.Map({d_83_v4_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cn"))})
        d_85_v6_: _dafny.Seq
        d_85_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvva"))
        index1_ = default__.safeIndex(288, (d_78_v0_).length(0))
        rhs0_ = 286
        rhs1_ = (d_84_v5_) | (d_84_v5_)
        rhs2_ = d_85_v6_
        lhs0_ = d_78_v0_
        lhs1_ = default__.safeIndex(288, (d_78_v0_).length(0))
        lhs0_[lhs1_] = rhs0_
        d_84_v5_ = rhs1_
        d_85_v6_ = rhs2_
        r2 = (d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]
        d_86_v7_: _dafny.MultiSet
        d_86_v7_ = (d_80_v1_ if p2 else d_80_v1_)
        d_87_v8_: _dafny.Set
        d_87_v8_ = _dafny.Set({d_81_v2_})
        d_88_v9_: D3
        d_88_v9_ = D3_DC9(d_81_v2_, d_81_v2_)
        d_89_v10_: _dafny.Map
        d_89_v10_ = _dafny.Map({p2: d_81_v2_})
        d_90_v11_: str
        d_90_v11_ = _dafny.CodePoint('s')
        rhs3_ = (len(d_87_v8_)) - (len((_dafny.Map({p0: (d_88_v9_).cf13})) | (d_89_v10_)))
        rhs4_ = False
        rhs5_ = d_86_v7_
        rhs6_ = ((d_85_v6_).set(default__.safeIndex(d_81_v2_, len(d_85_v6_)), d_90_v11_)) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtswhfxei"))) + (d_85_v6_))
        rhs7_ = (d_87_v8_) | (default__.fm70(p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qecc")), d_81_v2_, p2, globalState))
        lhs2_ = globalState
        r2 = rhs3_
        lhs2_.f4 = rhs4_
        d_86_v7_ = rhs5_
        r1 = rhs6_
        d_87_v8_ = rhs7_
        d_91_i1_: int
        d_91_i1_ = 0
        with _dafny.label("0"):
            while p2:
                with _dafny.c_label("0"):
                    if (d_91_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_91_i1_ = (d_91_i1_) + (1)
                    d_92_v12_: _dafny.Array
                    def lambda1_(d_93_v2_, d_94_v0_):
                        def lambda2_(d_95_i2_):
                            return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_93_v2_])).set(default__.safeIndex((d_94_v0_)[default__.safeIndex(288, (d_94_v0_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_93_v2_]))), (d_94_v0_)[default__.safeIndex(288, (d_94_v0_).length(0))]))).isdisjoint(_dafny.MultiSet([(0) - (d_93_v2_), (_dafny.MultiSet([d_93_v2_, (d_94_v0_)[default__.safeIndex(288, (d_94_v0_).length(0))]])).cardinality]))

                        return lambda2_

                    init1_ = lambda1_(d_81_v2_, d_78_v0_)
                    nw2_ = _dafny.Array(None, 10)
                    for i0_1_ in range(nw2_.length(0)):
                        nw2_[i0_1_] = init1_(i0_1_)
                    d_92_v12_ = nw2_
                    index2_ = default__.safeIndex(530, (d_92_v12_).length(0))
                    (d_92_v12_)[index2_] = p2
                    d_96_v13_: C4
                    nw3_ = C4()
                    nw3_.ctor__()
                    d_96_v13_ = nw3_
                    d_97_v14_: _dafny.Map
                    d_97_v14_ = _dafny.Map({d_96_v13_: d_87_v8_})
                    d_98_v15_: _dafny.Set
                    d_98_v15_ = _dafny.Set({d_87_v8_, ((d_97_v14_)[d_96_v13_] if (d_96_v13_) in (d_97_v14_) else d_87_v8_)})
                    d_99_v16_: _dafny.Seq
                    d_99_v16_ = _dafny.SeqWithoutIsStrInference([p2, p2, False])
                    d_100_v17_: _dafny.Map
                    d_100_v17_ = _dafny.Map({(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]: d_89_v10_})
                    d_101_v18_: _dafny.Map
                    d_101_v18_ = _dafny.Map({497: _dafny.Set({_dafny.Set({(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))], d_81_v2_, default__.fm0((d_99_v16_)[default__.safeIndex(d_81_v2_, len(d_99_v16_))], d_100_v17_, p2, globalState), (d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]})})})
                    index3_ = default__.safeIndex(530, (d_92_v12_).length(0))
                    (d_92_v12_)[index3_] = not(((d_98_v15_) - (_dafny.Set({d_87_v8_, d_87_v8_, _dafny.Set({d_81_v2_}), d_87_v8_, d_87_v8_}))) == (((d_101_v18_)[809] if (809) in (d_101_v18_) else d_98_v15_)))
                    d_102_v19_: _dafny.Map
                    d_102_v19_ = _dafny.Map({p0: (d_92_v12_)[default__.safeIndex(530, (d_92_v12_).length(0))]})
                    d_103_v20_: _dafny.Map
                    d_103_v20_ = _dafny.Map({p2: d_102_v19_})
                    d_104_v21_: _dafny.MultiSet
                    d_104_v21_ = _dafny.MultiSet([d_81_v2_, d_81_v2_, 213, (d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
                    d_105_v22_: _dafny.Seq
                    d_105_v22_ = _dafny.SeqWithoutIsStrInference([d_104_v21_])
                    d_106_v23_: _dafny.Set
                    d_106_v23_ = _dafny.Set({p2})
                    d_107_v24_: _dafny.Array
                    nw4_ = _dafny.Array(None, 28)
                    nw4_[int(0)] = d_104_v21_
                    nw4_[int(1)] = d_104_v21_
                    nw4_[int(2)] = d_104_v21_
                    nw4_[int(3)] = d_104_v21_
                    nw4_[int(4)] = d_104_v21_
                    nw4_[int(5)] = d_104_v21_
                    nw4_[int(6)] = d_104_v21_
                    nw4_[int(7)] = d_104_v21_
                    nw4_[int(8)] = d_104_v21_
                    nw4_[int(9)] = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
                    nw4_[int(10)] = d_104_v21_
                    nw4_[int(11)] = d_104_v21_
                    nw4_[int(12)] = d_104_v21_
                    nw4_[int(13)] = d_104_v21_
                    nw4_[int(14)] = d_104_v21_
                    nw4_[int(15)] = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
                    nw4_[int(16)] = d_104_v21_
                    nw4_[int(17)] = (d_105_v22_)[default__.safeIndex(523, len(d_105_v22_))]
                    nw4_[int(18)] = d_104_v21_
                    nw4_[int(19)] = _dafny.MultiSet([d_81_v2_, len(d_106_v23_)])
                    nw4_[int(20)] = d_104_v21_
                    nw4_[int(21)] = d_104_v21_
                    nw4_[int(22)] = d_104_v21_
                    nw4_[int(23)] = d_104_v21_
                    nw4_[int(24)] = d_104_v21_
                    nw4_[int(25)] = d_104_v21_
                    nw4_[int(26)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_81_v2_]))
                    nw4_[int(27)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([926]))
                    d_107_v24_ = nw4_
                    d_108_v25_: C2
                    nw5_ = C2()
                    nw5_.ctor__(_dafny.SeqWithoutIsStrInference([d_81_v2_]), p2, d_103_v20_, d_99_v16_, d_107_v24_)
                    d_108_v25_ = nw5_
                    d_109_v26_: _dafny.Map
                    d_109_v26_ = _dafny.Map({d_108_v25_: d_104_v21_})
                    d_110_v27_: _dafny.Seq
                    d_110_v27_ = _dafny.SeqWithoutIsStrInference([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
                    d_111_v28_: _dafny.Array
                    nw6_ = _dafny.Array(None, 19)
                    nw6_[int(0)] = d_104_v21_
                    nw6_[int(1)] = d_104_v21_
                    nw6_[int(2)] = d_104_v21_
                    nw6_[int(3)] = d_104_v21_
                    nw6_[int(4)] = d_104_v21_
                    nw6_[int(5)] = d_104_v21_
                    nw6_[int(6)] = d_104_v21_
                    nw6_[int(7)] = d_104_v21_
                    nw6_[int(8)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([157 for d_112_i3_ in range(default__.abs(524))]))
                    nw6_[int(9)] = _dafny.MultiSet([602])
                    nw6_[int(10)] = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
                    nw6_[int(11)] = d_104_v21_
                    nw6_[int(12)] = _dafny.MultiSet([d_81_v2_])
                    nw6_[int(13)] = d_104_v21_
                    nw6_[int(14)] = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))], d_81_v2_, 873, d_81_v2_])
                    nw6_[int(15)] = d_104_v21_
                    nw6_[int(16)] = d_104_v21_
                    nw6_[int(17)] = ((d_109_v26_)[d_108_v25_] if (d_108_v25_) in (d_109_v26_) else d_104_v21_)
                    nw6_[int(18)] = _dafny.MultiSet(d_110_v27_)
                    d_111_v28_ = nw6_
                    d_113_v29_: C1
                    nw7_ = C1()
                    nw7_.ctor__(p0, True, (d_103_v20_) | (d_103_v20_), (d_99_v16_) + (d_99_v16_), d_107_v24_)
                    d_113_v29_ = nw7_
                    d_113_v29_ = d_113_v29_
                    index4_ = default__.safeIndex(288, (d_78_v0_).length(0))
                    (d_78_v0_)[index4_] = d_81_v2_
                    d_114_v30_: _dafny.Array
                    def lambda3_(d_115_v6_):
                        def lambda4_(d_116_i4_):
                            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (d_115_v6_)

                        return lambda4_

                    init2_ = lambda3_(d_85_v6_)
                    nw8_ = _dafny.Array(None, 16)
                    for i0_2_ in range(nw8_.length(0)):
                        nw8_[i0_2_] = init2_(i0_2_)
                    d_114_v30_ = nw8_
                    d_117_v31_: T3
                    nw9_ = C2()
                    nw9_.ctor__(d_110_v27_, (d_87_v8_).isdisjoint(d_87_v8_), (d_103_v20_ if (d_113_v29_).f17 else d_103_v20_), d_99_v16_, d_111_v28_)
                    d_117_v31_ = nw9_
                    rhs8_ = d_114_v30_
                    rhs9_ = d_81_v2_
                    rhs10_ = d_117_v31_
                    rhs11_ = (d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]
                    rhs12_ = d_92_v12_
                    d_114_v30_ = rhs8_
                    d_81_v2_ = rhs9_
                    d_117_v31_ = rhs10_
                    d_81_v2_ = rhs11_
                    d_92_v12_ = rhs12_
                    pass
            pass
        d_118_v32_: _dafny.Seq
        d_118_v32_ = _dafny.SeqWithoutIsStrInference([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
        d_119_v33_: _dafny.Seq
        d_119_v33_ = _dafny.SeqWithoutIsStrInference([p2])
        d_120_v34_: _dafny.Map
        d_120_v34_ = _dafny.Map({p2: p0})
        d_121_v35_: _dafny.Map
        d_121_v35_ = _dafny.Map({(d_119_v33_)[default__.safeIndex(len(d_120_v34_), len(d_119_v33_))]: p0})
        d_122_v36_: _dafny.Map
        d_122_v36_ = _dafny.Map({p2: d_121_v35_})
        d_123_v37_: _dafny.MultiSet
        d_123_v37_ = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
        d_124_v38_: _dafny.Map
        d_124_v38_ = _dafny.Map({False: d_123_v37_})
        d_125_v39_: _dafny.Map
        d_125_v39_ = _dafny.Map({d_119_v33_: p2})
        d_126_v40_: _dafny.Array
        nw10_ = _dafny.Array(_dafny.MultiSet({}), 19)
        d_126_v40_ = nw10_
        d_127_v41_: C5
        nw11_ = C5()
        nw11_.ctor__(p0, d_126_v40_, d_122_v36_, d_119_v33_)
        d_127_v41_ = nw11_
        d_128_v42_: _dafny.Map
        d_128_v42_ = _dafny.Map({p2: d_127_v41_})
        d_129_v43_: _dafny.MultiSet
        d_129_v43_ = _dafny.MultiSet([d_128_v42_])
        d_130_v44_: D14
        d_130_v44_ = D14_DC41(d_85_v6_, d_81_v2_)
        d_131_v45_: _dafny.Seq
        d_131_v45_ = _dafny.SeqWithoutIsStrInference([d_130_v44_])
        d_132_v46_: D26
        d_132_v46_ = D26_DC70(d_131_v45_)
        d_133_v47_: _dafny.Array
        nw12_ = _dafny.Array(None, 25)
        nw12_[int(0)] = d_123_v37_
        nw12_[int(1)] = d_123_v37_
        nw12_[int(2)] = _dafny.MultiSet(d_118_v32_)
        nw12_[int(3)] = d_123_v37_
        nw12_[int(4)] = ((d_124_v38_)[((d_125_v39_)[d_119_v33_] if (d_119_v33_) in (d_125_v39_) else not(p0))] if (((d_125_v39_)[d_119_v33_] if (d_119_v33_) in (d_125_v39_) else not(p0))) in (d_124_v38_) else d_123_v37_)
        nw12_[int(5)] = ((d_124_v38_)[p2] if (p2) in (d_124_v38_) else d_123_v37_)
        nw12_[int(6)] = _dafny.MultiSet([((d_129_v43_)[d_128_v42_] if (d_128_v42_) in (d_129_v43_) else (d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]), len(d_85_v6_)])
        nw12_[int(7)] = d_123_v37_
        nw12_[int(8)] = d_123_v37_
        nw12_[int(9)] = d_123_v37_
        nw12_[int(10)] = d_123_v37_
        nw12_[int(11)] = ((d_124_v38_)[True] if (True) in (d_124_v38_) else _dafny.MultiSet([d_81_v2_, (d_123_v37_).cardinality, len(d_119_v33_)]))
        nw12_[int(12)] = _dafny.MultiSet([d_81_v2_])
        nw12_[int(13)] = d_123_v37_
        nw12_[int(14)] = (_dafny.MultiSet(d_118_v32_)).set(274, default__.abs(len((d_132_v46_).cf106)))
        nw12_[int(15)] = _dafny.MultiSet([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]])
        nw12_[int(16)] = d_123_v37_
        nw12_[int(17)] = d_123_v37_
        nw12_[int(18)] = d_123_v37_
        nw12_[int(19)] = d_123_v37_
        nw12_[int(20)] = default__.fm28(d_85_v6_, p0, d_81_v2_, globalState)
        nw12_[int(21)] = d_123_v37_
        nw12_[int(22)] = d_123_v37_
        nw12_[int(23)] = (default__.fm56(d_90_v11_, globalState)).set(d_81_v2_, default__.abs((d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))]))
        nw12_[int(24)] = _dafny.MultiSet([d_81_v2_])
        d_133_v47_ = nw12_
        d_134_v48_: C2
        nw13_ = C2()
        nw13_.ctor__((d_118_v32_) + (d_118_v32_), not(p0), d_122_v36_, d_119_v33_, d_133_v47_)
        d_134_v48_ = nw13_
        if not(p0):
            d_135_v49_: int
            d_136_v50_: int
            d_137_v51_: D0
            out0_: int
            out1_: int
            out2_: D0
            out0_, out1_, out2_ = (d_82_v3_).m2((d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))], globalState)
            d_135_v49_ = out0_
            d_136_v50_ = out1_
            d_137_v51_ = out2_
            d_138_v52_: _dafny.Array
            nw14_ = _dafny.Array(_dafny.MultiSet({}), 21)
            d_138_v52_ = nw14_
            index5_ = default__.safeIndex(880, (d_138_v52_).length(0))
            (d_138_v52_)[index5_] = d_80_v1_
            index6_ = default__.safeIndex(880, (d_138_v52_).length(0))
            (d_138_v52_)[index6_] = d_80_v1_
            d_139_v53_: _dafny.Map
            d_139_v53_ = _dafny.Map({d_85_v6_: d_78_v0_})
            d_140_v54_: _dafny.Map
            d_140_v54_ = _dafny.Map({123: d_78_v0_})
            d_141_v55_: _dafny.Seq
            d_141_v55_ = _dafny.SeqWithoutIsStrInference([((d_140_v54_)[d_136_v50_] if (d_136_v50_) in (d_140_v54_) else d_78_v0_), d_78_v0_])
            d_139_v53_ = (d_139_v53_).set((d_85_v6_) + (d_85_v6_), (d_141_v55_)[default__.safeIndex((d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))], len(d_141_v55_))])
            d_142_v56_: _dafny.Array
            nw15_ = _dafny.Array(None, 8)
            nw15_[int(0)] = (d_127_v41_).f19
            nw15_[int(1)] = p0
            nw15_[int(2)] = (d_127_v41_).f19
            nw15_[int(3)] = (d_127_v41_).f19
            nw15_[int(4)] = (d_127_v41_).f19
            nw15_[int(5)] = p2
            nw15_[int(6)] = (d_127_v41_).f19
            nw15_[int(7)] = p0
            d_142_v56_ = nw15_
            d_143_v57_: _dafny.Map
            d_143_v57_ = _dafny.Map({-756: p2})
            index7_ = default__.safeIndex(36, (d_142_v56_).length(0))
            (d_142_v56_)[index7_] = (d_135_v49_) not in (d_143_v57_)
            index8_ = default__.safeIndex(36, (d_142_v56_).length(0))
            (d_142_v56_)[index8_] = p0
            d_144_v58_: _dafny.Array
            nw16_ = _dafny.Array(None, 11)
            nw16_[int(0)] = _dafny.CodePoint('s')
            nw16_[int(1)] = d_90_v11_
            nw16_[int(2)] = d_90_v11_
            nw16_[int(3)] = d_90_v11_
            nw16_[int(4)] = d_90_v11_
            nw16_[int(5)] = d_90_v11_
            nw16_[int(6)] = d_90_v11_
            nw16_[int(7)] = d_90_v11_
            nw16_[int(8)] = d_90_v11_
            nw16_[int(9)] = d_90_v11_
            nw16_[int(10)] = d_90_v11_
            d_144_v58_ = nw16_
            d_145_v59_: _dafny.Seq
            d_145_v59_ = _dafny.SeqWithoutIsStrInference([d_144_v58_])
            d_143_v57_ = (d_143_v57_).set(len((d_145_v59_) + (d_145_v59_)), (d_127_v41_).f19)
        elif True:
            d_146_v60_: _dafny.Array
            def lambda5_(d_147_v37_):
                def lambda6_(d_148_i5_):
                    return (d_147_v37_).ispropersubset(d_147_v37_)

                return lambda6_

            init3_ = lambda5_(d_123_v37_)
            nw17_ = _dafny.Array(None, 24)
            for i0_3_ in range(nw17_.length(0)):
                nw17_[i0_3_] = init3_(i0_3_)
            d_146_v60_ = nw17_
            d_149_v61_: C4
            nw18_ = C4()
            nw18_.ctor__()
            d_149_v61_ = nw18_
            d_150_v62_: _dafny.Map
            d_150_v62_ = _dafny.Map({len(d_85_v6_): _dafny.Set({d_149_v61_, d_149_v61_})})
            d_151_v63_: _dafny.Set
            d_151_v63_ = _dafny.Set({d_149_v61_})
            d_152_v64_: _dafny.Map
            d_152_v64_ = _dafny.Map({_dafny.Set({d_149_v61_}): d_90_v11_})
            index9_ = default__.safeIndex(640, (d_146_v60_).length(0))
            (d_146_v60_)[index9_] = ((_dafny.Map({((d_150_v62_)[(d_123_v37_).cardinality] if ((d_123_v37_).cardinality) in (d_150_v62_) else d_151_v63_): d_90_v11_})).set(d_151_v63_, d_90_v11_)) == (d_152_v64_)
            index10_ = default__.safeIndex(640, (d_146_v60_).length(0))
            (d_146_v60_)[index10_] = p2
            d_153_v65_: _dafny.Array
            nw19_ = _dafny.Array(None, 15)
            nw19_[int(0)] = _dafny.CodePoint('p')
            nw19_[int(1)] = d_90_v11_
            nw19_[int(2)] = d_90_v11_
            nw19_[int(3)] = d_90_v11_
            nw19_[int(4)] = d_90_v11_
            nw19_[int(5)] = d_90_v11_
            nw19_[int(6)] = d_90_v11_
            nw19_[int(7)] = d_90_v11_
            nw19_[int(8)] = d_90_v11_
            nw19_[int(9)] = d_90_v11_
            nw19_[int(10)] = (d_85_v6_)[default__.safeIndex(d_81_v2_, len(d_85_v6_))]
            nw19_[int(11)] = d_90_v11_
            nw19_[int(12)] = d_90_v11_
            nw19_[int(13)] = d_90_v11_
            nw19_[int(14)] = d_90_v11_
            d_153_v65_ = nw19_
            d_154_v66_: _dafny.Map
            d_154_v66_ = _dafny.Map({d_146_v60_: (d_153_v65_ if (d_146_v60_)[default__.safeIndex(640, (d_146_v60_).length(0))] else d_153_v65_)})
            d_154_v66_ = (d_154_v66_).set(d_146_v60_, d_153_v65_)
            d_85_v6_ = d_85_v6_
            d_89_v10_ = (d_89_v10_) | (_dafny.Map({p0: d_81_v2_}))
            r3 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpbqponc"))) <= ((_dafny.SeqWithoutIsStrInference([d_90_v11_ for d_155_i6_ in range(default__.abs(359))])) + (_dafny.SeqWithoutIsStrInference([d_90_v11_ for d_156_i7_ in range(default__.abs(-418))])))
        r0 = (d_118_v32_) + (_dafny.SeqWithoutIsStrInference([(d_78_v0_)[default__.safeIndex(288, (d_78_v0_).length(0))] for d_157_i8_ in range(default__.abs(75))]))
        r1 = p2
        d_158_v67_: D8
        d_158_v67_ = D8_DC25((d_127_v41_).f19, len(d_119_v33_), d_81_v2_, d_89_v10_)
        r2 = default__.safeModuloInt((d_81_v2_) - (len(default__.fm35((d_80_v1_).cardinality, d_158_v67_, globalState))), (d_134_v48_).fm8((d_127_v41_).f19, p2, globalState))
        r3 = p0
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_159_v0_: _dafny.Array
        nw20_ = _dafny.Array(int(0), 6)
        d_159_v0_ = nw20_
        d_160_v1_: bool
        d_160_v1_ = False
        d_161_v2_: _dafny.MultiSet
        d_161_v2_ = _dafny.MultiSet([d_160_v1_])
        d_162_v3_: int
        d_162_v3_ = 846
        d_163_v4_: _dafny.Map
        d_163_v4_ = _dafny.Map({d_162_v3_: d_162_v3_})
        d_164_globalState_: GlobalState
        nw21_ = GlobalState()
        nw21_.ctor__(False, d_159_v0_, -689, 665, False, d_161_v2_, 409, -44, False, False, d_163_v4_, False)
        d_164_globalState_ = nw21_
        def iife31_():
            coll31_ = _dafny.Map()
            compr_31_: int
            for compr_31_ in _dafny.IntegerRange(481, -667):
                d_165_v5_: int = compr_31_
                if ((481) <= (d_165_v5_)) and ((d_165_v5_) < (-667)):
                    coll31_[(d_165_v5_) + (len(_dafny.SeqWithoutIsStrInference([d_160_v1_])))] = _dafny.Map({d_160_v1_: d_162_v3_})
            return _dafny.Map(coll31_)
        hi0_ = default__.fm0(d_160_v1_, iife31_()
        , d_160_v1_, d_164_globalState_)
        for d_166_i0_ in range(d_162_v3_, hi0_):
            d_167_v7_: _dafny.Array
            def lambda7_(d_168_v3_):
                def lambda8_(d_169_i1_):
                    def iife32_():
                        coll32_ = _dafny.Set()
                        compr_32_: int
                        for compr_32_ in _dafny.IntegerRange(-584, -15):
                            d_170_v6_: int = compr_32_
                            if ((-584) <= (d_170_v6_)) and ((d_170_v6_) < (-15)):
                                coll32_ = coll32_.union(_dafny.Set([(d_170_v6_) + ((0) - (d_168_v3_))]))
                        return _dafny.Set(coll32_)
                    return iife32_()
                    

                return lambda8_

            init4_ = lambda7_(d_162_v3_)
            nw22_ = _dafny.Array(None, 7)
            for i0_4_ in range(nw22_.length(0)):
                nw22_[i0_4_] = init4_(i0_4_)
            d_167_v7_ = nw22_
            d_167_v7_ = d_167_v7_
            d_160_v1_ = d_160_v1_
            d_171_v8_: str
            d_171_v8_ = _dafny.CodePoint('k')
            d_172_v9_: _dafny.Map
            d_172_v9_ = _dafny.Map({d_171_v8_: _dafny.CodePoint('o')})
            d_173_v10_: _dafny.Seq
            d_173_v10_ = _dafny.SeqWithoutIsStrInference([len(d_172_v9_)])
            d_174_v11_: D0
            d_174_v11_ = D0_DC2((d_161_v2_).cardinality, d_160_v1_)
            d_175_v12_: _dafny.Map
            d_175_v12_ = _dafny.Map({(d_173_v10_) == (default__.fm1((d_174_v11_).cf2, d_160_v1_, d_164_globalState_)): d_160_v1_})
            d_175_v12_ = (d_175_v12_).set(False, default__.fm2(d_160_v1_, d_164_globalState_))
            if d_160_v1_:
                d_160_v1_ = d_160_v1_
                d_176_v13_: _dafny.Seq
                d_176_v13_ = _dafny.SeqWithoutIsStrInference([d_160_v1_, d_160_v1_])
                d_177_v14_: _dafny.Map
                d_177_v14_ = _dafny.Map({d_173_v10_: d_160_v1_})
                d_178_v15_: _dafny.Set
                d_178_v15_ = _dafny.Set({len(d_177_v14_)})
                d_162_v3_ = (default__.safeDivisionInt(len(d_176_v13_), len(d_178_v15_))) * ((890) + (701))
                (d_164_globalState_).f4 = d_160_v1_
                d_162_v3_ = d_162_v3_
                d_179_v16_: D0
                d_179_v16_ = D0_DC1()
                d_179_v16_ = d_179_v16_
            elif True:
                d_180_v17_: _dafny.Seq
                d_181_v18_: bool
                d_182_v19_: int
                d_183_v20_: bool
                out3_: _dafny.Seq
                out4_: bool
                out5_: int
                out6_: bool
                out3_, out4_, out5_, out6_ = default__.m0((d_160_v1_ if d_160_v1_ else d_160_v1_), d_174_v11_, (d_166_i0_) <= (d_166_i0_), d_164_globalState_)
                d_180_v17_ = out3_
                d_181_v18_ = out4_
                d_182_v19_ = out5_
                d_183_v20_ = out6_
                d_184_v21_: _dafny.MultiSet
                d_184_v21_ = d_161_v2_
                d_181_v18_ = ((d_184_v21_)).issubset(d_161_v2_)
                d_185_v22_: _dafny.Seq
                d_185_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lomvp"))
                d_186_v23_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_186_v23_ = nw23_
                index11_ = default__.safeIndex(770, (d_186_v23_).length(0))
                (d_186_v23_)[index11_] = (d_161_v2_ if d_181_v18_ else d_161_v2_)
                index12_ = default__.safeIndex(770, (d_186_v23_).length(0))
                rhs13_ = not (d_183_v20_) or (not(default__.fm2(d_183_v20_, d_164_globalState_)))
                rhs14_ = D0_DC2(len(d_173_v10_), d_181_v18_)
                rhs15_ = d_185_v22_
                rhs16_ = (_dafny.MultiSet([d_183_v20_, False, d_160_v1_]) if d_160_v1_ else (_dafny.MultiSet([True, d_160_v1_]) if d_181_v18_ else d_161_v2_))
                lhs3_ = d_186_v23_
                lhs4_ = default__.safeIndex(770, (d_186_v23_).length(0))
                d_181_v18_ = rhs13_
                d_174_v11_ = rhs14_
                d_185_v22_ = rhs15_
                lhs3_[lhs4_] = rhs16_
                d_187_v24_: _dafny.Map
                d_187_v24_ = _dafny.Map({d_185_v22_: d_160_v1_})
                (d_164_globalState_).f4 = ((d_187_v24_)[d_185_v22_] if (d_185_v22_) in (d_187_v24_) else d_181_v18_)
                d_188_v25_: _dafny.Set
                d_188_v25_ = _dafny.Set({d_160_v1_})
                d_188_v25_ = d_188_v25_
        hi1_ = d_162_v3_
        for d_189_i2_ in range(d_162_v3_, hi1_):
            d_190_v26_: _dafny.Set
            d_190_v26_ = _dafny.Set({False, d_160_v1_, d_160_v1_, d_160_v1_, default__.fm2(d_160_v1_, d_164_globalState_)})
            d_191_v27_: _dafny.Map
            d_191_v27_ = _dafny.Map({-77: d_160_v1_})
            d_192_v28_: _dafny.Seq
            d_192_v28_ = _dafny.SeqWithoutIsStrInference([d_160_v1_, d_160_v1_, d_160_v1_, d_160_v1_])
            d_193_v29_: _dafny.Seq
            d_193_v29_ = _dafny.SeqWithoutIsStrInference([d_160_v1_, ((d_191_v27_)[(0) - (len(d_192_v28_))] if ((0) - (len(d_192_v28_))) in (d_191_v27_) else d_160_v1_), d_160_v1_, d_160_v1_, not(d_160_v1_)])
            d_194_v30_: _dafny.Map
            d_194_v30_ = _dafny.Map({(d_190_v26_).intersection(d_190_v26_): (d_192_v28_) == (_dafny.SeqWithoutIsStrInference([d_160_v1_]))})
            d_194_v30_ = (d_194_v30_).set(d_190_v26_, d_160_v1_)
            d_195_v31_: _dafny.Array
            def lambda9_(d_196_i2_):
                def lambda10_(d_197_i3_):
                    return _dafny.MultiSet([d_196_i2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tywvlxp")))])

                return lambda10_

            init5_ = lambda9_(d_189_i2_)
            nw24_ = _dafny.Array(None, 20)
            for i0_5_ in range(nw24_.length(0)):
                nw24_[i0_5_] = init5_(i0_5_)
            d_195_v31_ = nw24_
            d_198_v32_: _dafny.Map
            d_198_v32_ = _dafny.Map({d_190_v26_: 305})
            d_199_v33_: _dafny.MultiSet
            d_199_v33_ = _dafny.MultiSet([d_162_v3_, len(d_198_v32_)])
            d_200_v34_: _dafny.Map
            d_200_v34_ = _dafny.Map({d_160_v1_: default__.fm2(d_160_v1_, d_164_globalState_)})
            d_201_v35_: _dafny.Map
            d_201_v35_ = _dafny.Map({d_160_v1_: d_200_v34_})
            d_202_v36_: C3
            nw25_ = C3()
            nw25_.ctor__(default__.fm2(d_160_v1_, d_164_globalState_), d_195_v31_, (((d_199_v33_)[(0) - (d_189_i2_)] if ((0) - (d_189_i2_)) in (d_199_v33_) else d_189_i2_)) < (d_189_i2_), d_201_v35_, _dafny.SeqWithoutIsStrInference([True, d_160_v1_, d_160_v1_]))
            d_202_v36_ = nw25_
            d_160_v1_ = (False if d_160_v1_ else not(((d_200_v34_)[(d_202_v36_).f20] if ((d_202_v36_).f20) in (d_200_v34_) else (d_192_v28_)[default__.safeIndex(d_189_i2_, len(d_192_v28_))])))
            d_203_v37_: _dafny.Seq
            d_203_v37_ = _dafny.SeqWithoutIsStrInference([d_189_i2_, d_189_i2_])
            d_204_v38_: C2
            nw26_ = C2()
            nw26_.ctor__(d_203_v37_, True, d_201_v35_, d_192_v28_, d_195_v31_)
            d_204_v38_ = nw26_
            d_205_v39_: _dafny.Map
            d_205_v39_ = _dafny.Map({d_204_v38_: (d_202_v36_).f20})
            d_160_v1_ = ((_dafny.MultiSet([(0) - (d_189_i2_), 881])).ispropersubset(d_199_v33_) if (d_202_v36_).f20 else ((d_205_v39_)[d_204_v38_] if (d_204_v38_) in (d_205_v39_) else d_160_v1_))
        d_206_v40_: _dafny.Set
        d_206_v40_ = _dafny.Set({d_160_v1_, d_160_v1_})
        hi2_ = d_162_v3_
        for d_207_i4_ in range(len(d_206_v40_), hi2_):
            d_208_v41_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_208_v41_ = nw27_
            d_209_v42_: _dafny.Array
            nw28_ = _dafny.Array(None, 1)
            nw28_[int(0)] = default__.fm2(False, d_164_globalState_)
            d_209_v42_ = nw28_
            index13_ = default__.safeIndex(552, (d_208_v41_).length(0))
            (d_208_v41_)[index13_] = d_209_v42_
            index14_ = default__.safeIndex(552, (d_208_v41_).length(0))
            (d_208_v41_)[index14_] = d_209_v42_
            if d_160_v1_:
                d_210_v43_: _dafny.Seq
                d_210_v43_ = _dafny.SeqWithoutIsStrInference([d_160_v1_, d_160_v1_, d_160_v1_])
                (d_164_globalState_).f4 = (d_160_v1_ if (d_210_v43_)[default__.safeIndex(233, len(d_210_v43_))] else d_160_v1_)
                (d_164_globalState_).f4 = d_160_v1_
                d_211_v44_: _dafny.Seq
                d_211_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwe"))
                d_212_v45_: str
                d_212_v45_ = _dafny.CodePoint('n')
                d_213_v46_: _dafny.Seq
                d_213_v46_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, d_207_i4_])
                d_214_v47_: _dafny.Map
                d_214_v47_ = _dafny.Map({d_160_v1_: d_160_v1_})
                d_215_v48_: _dafny.Map
                d_215_v48_ = _dafny.Map({d_160_v1_: d_214_v47_})
                d_216_v49_: _dafny.Array
                nw29_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_216_v49_ = nw29_
                d_217_v50_: C6
                nw30_ = C6()
                nw30_.ctor__(d_213_v46_, d_160_v1_, d_215_v48_, d_210_v43_, d_216_v49_)
                d_217_v50_ = nw30_
                d_218_v51_: _dafny.Seq
                d_218_v51_ = _dafny.SeqWithoutIsStrInference([d_217_v50_])
                d_219_v52_: _dafny.Map
                d_219_v52_ = _dafny.Map({(len((d_211_v44_).set(default__.safeIndex(d_162_v3_, len(d_211_v44_)), d_212_v45_))) * (len(d_218_v51_)): d_212_v45_})
                d_219_v52_ = (d_219_v52_).set(d_162_v3_, d_212_v45_)
                d_220_v53_: _dafny.Map
                d_220_v53_ = _dafny.Map({d_207_i4_: d_211_v44_})
                d_221_v54_: bool
                out7_: bool
                out7_ = (d_217_v50_).m5(d_206_v40_, (d_211_v44_) + (((d_220_v53_)[761] if (761) in (d_220_v53_) else d_211_v44_)), (0) - (default__.safeDivisionInt(718, len(d_211_v44_))), d_164_globalState_)
                d_221_v54_ = out7_
                (d_164_globalState_).f4 = (d_207_i4_) <= (d_207_i4_)
            elif True:
                d_222_v55_: _dafny.Seq
                d_222_v55_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, d_207_i4_])
                d_223_v56_: _dafny.Map
                d_223_v56_ = _dafny.Map({d_160_v1_: len(d_222_v55_)})
                d_224_v57_: _dafny.Seq
                d_224_v57_ = _dafny.SeqWithoutIsStrInference([default__.fm0(False, _dafny.Map({d_162_v3_: d_223_v56_}), d_160_v1_, d_164_globalState_)])
                d_224_v57_ = (d_222_v55_) + (d_222_v55_)
                d_225_v58_: _dafny.Array
                def lambda11_(d_226_v1_):
                    def lambda12_(d_227_i5_):
                        return D4_DC12(_dafny.SeqWithoutIsStrInference([d_226_v1_]))

                    return lambda12_

                init6_ = lambda11_(d_160_v1_)
                nw31_ = _dafny.Array(None, 13)
                for i0_6_ in range(nw31_.length(0)):
                    nw31_[i0_6_] = init6_(i0_6_)
                d_225_v58_ = nw31_
                index15_ = default__.safeIndex(680, (d_159_v0_).length(0))
                (d_159_v0_)[index15_] = default__.safeDivisionInt((0) - ((0) - (d_207_i4_)), d_162_v3_)
                index16_ = default__.safeIndex(296, (d_159_v0_).length(0))
                (d_159_v0_)[index16_] = (d_162_v3_) + (d_162_v3_)
                d_228_v59_: _dafny.Seq
                d_228_v59_ = _dafny.SeqWithoutIsStrInference([d_225_v58_, d_225_v58_, d_225_v58_])
                d_229_v60_: _dafny.Map
                d_229_v60_ = _dafny.Map({d_207_i4_: d_160_v1_})
                d_230_v61_: _dafny.Map
                d_230_v61_ = _dafny.Map({d_162_v3_: d_223_v56_})
                index17_ = default__.safeIndex(680, (d_159_v0_).length(0))
                index18_ = default__.safeIndex(296, (d_159_v0_).length(0))
                rhs17_ = (d_228_v59_)[default__.safeIndex(len((d_229_v60_ if d_160_v1_ else d_229_v60_)), len(d_228_v59_))]
                rhs18_ = default__.fm0(True, d_230_v61_, (True if ((d_229_v60_)[d_207_i4_] if (d_207_i4_) in (d_229_v60_) else d_160_v1_) else d_160_v1_), d_164_globalState_)
                rhs19_ = (d_207_i4_) - (default__.safeModuloInt(d_162_v3_, (0) - (len(default__.fm68(d_160_v1_, d_164_globalState_)))))
                rhs20_ = (d_207_i4_) - (d_162_v3_)
                lhs5_ = d_159_v0_
                lhs6_ = default__.safeIndex(680, (d_159_v0_).length(0))
                lhs7_ = d_159_v0_
                lhs8_ = default__.safeIndex(296, (d_159_v0_).length(0))
                d_225_v58_ = rhs17_
                lhs5_[lhs6_] = rhs18_
                d_162_v3_ = rhs19_
                lhs7_[lhs8_] = rhs20_
                (d_164_globalState_).f4 = d_160_v1_
                d_231_v62_: _dafny.Seq
                d_231_v62_ = _dafny.SeqWithoutIsStrInference([d_209_v42_])
                d_232_v63_: D21
                d_232_v63_ = D21_DC55(d_231_v62_)
                d_233_v64_: D0
                d_233_v64_ = D0_DC2(len((d_232_v63_).cf84), True)
                d_234_v65_: _dafny.Map
                d_234_v65_ = _dafny.Map({d_160_v1_: d_160_v1_})
                d_235_v66_: _dafny.Seq
                d_236_v67_: bool
                d_237_v68_: int
                d_238_v69_: bool
                out8_: _dafny.Seq
                out9_: bool
                out10_: int
                out11_: bool
                out8_, out9_, out10_, out11_ = default__.m0(d_160_v1_, d_233_v64_, ((d_234_v65_)[True] if (True) in (d_234_v65_) else d_160_v1_), d_164_globalState_)
                d_235_v66_ = out8_
                d_236_v67_ = out9_
                d_237_v68_ = out10_
                d_238_v69_ = out11_
                d_239_v70_: _dafny.Map
                d_239_v70_ = _dafny.Map({False: d_234_v65_})
                d_240_v71_: _dafny.Seq
                d_240_v71_ = _dafny.SeqWithoutIsStrInference([d_236_v67_])
                d_241_v72_: _dafny.Array
                def lambda13_(d_242_v55_):
                    def lambda14_(d_243_i6_):
                        return _dafny.MultiSet(d_242_v55_)

                    return lambda14_

                init7_ = lambda13_(d_222_v55_)
                nw32_ = _dafny.Array(None, 9)
                for i0_7_ in range(nw32_.length(0)):
                    nw32_[i0_7_] = init7_(i0_7_)
                d_241_v72_ = nw32_
                d_244_v73_: T1
                nw33_ = C1()
                nw33_.ctor__(d_160_v1_, d_160_v1_, d_239_v70_, d_240_v71_, d_241_v72_)
                d_244_v73_ = nw33_
                d_245_v74_: C5
                nw34_ = C5()
                nw34_.ctor__(False, d_244_v73_.f12, (d_244_v73_).f13, (d_244_v73_).f14)
                d_245_v74_ = nw34_
                d_246_v75_: _dafny.Map
                d_246_v75_ = _dafny.Map({len((_dafny.Map({d_162_v3_: d_245_v74_})).set((d_159_v0_)[default__.safeIndex(680, (d_159_v0_).length(0))], d_245_v74_)): d_209_v42_})
                d_247_v76_: _dafny.Map
                d_247_v76_ = _dafny.Map({d_244_v73_: ((d_246_v75_)[(0) - (d_207_i4_)] if ((0) - (d_207_i4_)) in (d_246_v75_) else (d_208_v41_)[default__.safeIndex(552, (d_208_v41_).length(0))])})
                d_248_v77_: _dafny.Map
                d_248_v77_ = _dafny.Map({d_160_v1_: d_244_v73_})
                d_247_v76_ = (d_247_v76_).set(((d_248_v77_)[(d_245_v74_).f19] if ((d_245_v74_).f19) in (d_248_v77_) else d_244_v73_), (d_208_v41_)[default__.safeIndex(552, (d_208_v41_).length(0))])
            d_249_v78_: _dafny.Map
            d_249_v78_ = _dafny.Map({d_160_v1_: False})
            d_250_v79_: _dafny.Map
            d_250_v79_ = _dafny.Map({False: d_249_v78_})
            d_251_v80_: _dafny.Array
            def lambda15_(d_252_i4_):
                def lambda16_(d_253_i7_):
                    return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkxripbcw"))), d_252_i4_])

                return lambda16_

            init8_ = lambda15_(d_207_i4_)
            nw35_ = _dafny.Array(None, 23)
            for i0_8_ in range(nw35_.length(0)):
                nw35_[i0_8_] = init8_(i0_8_)
            d_251_v80_ = nw35_
            d_254_v81_: T1
            nw36_ = C1()
            nw36_.ctor__(not(False), d_160_v1_, d_250_v79_, _dafny.SeqWithoutIsStrInference([d_160_v1_]), d_251_v80_)
            d_254_v81_ = nw36_
            d_255_v82_: _dafny.Map
            d_255_v82_ = _dafny.Map({d_162_v3_: d_254_v81_})
            d_256_v83_: _dafny.Map
            d_256_v83_ = _dafny.Map({d_255_v82_: d_160_v1_})
            d_257_v84_: _dafny.Map
            d_257_v84_ = _dafny.Map({False: d_207_i4_})
            d_258_v85_: _dafny.Map
            d_258_v85_ = _dafny.Map({d_207_i4_: d_257_v84_})
            d_256_v83_ = (d_256_v83_).set((d_255_v82_).set(d_162_v3_, d_254_v81_), (default__.fm0(d_160_v1_, d_258_v85_, d_160_v1_, d_164_globalState_)) <= (len(d_257_v84_)))
            index19_ = default__.safeIndex(552, (d_208_v41_).length(0))
            (d_208_v41_)[index19_] = d_209_v42_
        (d_164_globalState_).f4 = (d_161_v2_).issubset(d_161_v2_)
        if d_160_v1_:
            index20_ = default__.safeIndex(947, (d_159_v0_).length(0))
            (d_159_v0_)[index20_] = d_162_v3_
            d_259_v86_: str
            d_259_v86_ = _dafny.CodePoint('i')
            d_260_v87_: _dafny.Map
            d_260_v87_ = _dafny.Map({d_160_v1_: d_259_v86_})
            index21_ = default__.safeIndex(14, (d_159_v0_).length(0))
            (d_159_v0_)[index21_] = len(d_260_v87_)
            d_261_v88_: _dafny.Seq
            d_261_v88_ = _dafny.SeqWithoutIsStrInference([d_160_v1_])
            d_262_v89_: _dafny.Map
            d_262_v89_ = _dafny.Map({d_162_v3_: _dafny.Map({default__.fm2(d_160_v1_, d_164_globalState_): len(d_261_v88_)})})
            d_263_v90_: _dafny.MultiSet
            d_263_v90_ = _dafny.MultiSet([D11_DC30(d_262_v89_)])
            d_264_v91_: D4
            d_264_v91_ = D4_DC13(d_160_v1_, d_160_v1_, d_162_v3_, d_162_v3_, (d_261_v88_)[default__.safeIndex((d_263_v90_).cardinality, len(d_261_v88_))])
            d_265_v92_: _dafny.Seq
            d_265_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_266_v93_: D4
            d_266_v93_ = D4_DC14(d_160_v1_)
            d_267_v94_: D12
            d_267_v94_ = D12_DC37(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "de")), d_160_v1_)
            d_268_v95_: _dafny.MultiSet
            d_268_v95_ = _dafny.MultiSet([default__.fm29((d_264_v91_).cf23, d_160_v1_, d_160_v1_, len(d_265_v92_), d_164_globalState_), ((d_265_v92_) + (default__.fm40(d_266_v93_, d_162_v3_, d_160_v1_, not(d_160_v1_), d_164_globalState_))).set(default__.safeIndex((0) - (d_162_v3_), len((d_265_v92_) + (default__.fm40(d_266_v93_, d_162_v3_, d_160_v1_, not(d_160_v1_), d_164_globalState_)))), default__.fm39(-338, d_160_v1_, d_162_v3_, (d_267_v94_).cf59, d_164_globalState_))])
            index22_ = default__.safeIndex(947, (d_159_v0_).length(0))
            index23_ = default__.safeIndex(14, (d_159_v0_).length(0))
            rhs21_ = d_162_v3_
            rhs22_ = ((d_268_v95_)[d_265_v92_] if (d_265_v92_) in (d_268_v95_) else d_162_v3_)
            rhs23_ = default__.safeModuloInt((d_161_v2_).cardinality, (d_162_v3_) + (len(d_261_v88_)))
            rhs24_ = default__.safeModuloInt(393, len((d_261_v88_) + (d_261_v88_)))
            lhs9_ = d_159_v0_
            lhs10_ = default__.safeIndex(947, (d_159_v0_).length(0))
            lhs11_ = d_159_v0_
            lhs12_ = default__.safeIndex(14, (d_159_v0_).length(0))
            lhs9_[lhs10_] = rhs21_
            lhs11_[lhs12_] = rhs22_
            d_162_v3_ = rhs23_
            d_162_v3_ = rhs24_
            d_269_v96_: _dafny.Map
            d_269_v96_ = _dafny.Map({not(d_160_v1_): d_162_v3_})
            d_270_v97_: _dafny.Map
            d_270_v97_ = _dafny.Map({d_160_v1_: d_160_v1_})
            d_271_v98_: _dafny.Map
            d_271_v98_ = _dafny.Map({d_160_v1_: d_270_v97_})
            d_272_v99_: _dafny.Array
            nw37_ = _dafny.Array(_dafny.MultiSet({}), 28)
            d_272_v99_ = nw37_
            d_273_v100_: _dafny.Map
            d_273_v100_ = _dafny.Map({348: d_272_v99_})
            d_274_v101_: C6
            nw38_ = C6()
            nw38_.ctor__((_dafny.SeqWithoutIsStrInference([d_162_v3_, (0) - (d_162_v3_)])).set(default__.safeIndex(default__.fm0(True, _dafny.Map({d_162_v3_: d_269_v96_}), d_160_v1_, d_164_globalState_), len(_dafny.SeqWithoutIsStrInference([d_162_v3_, (0) - (d_162_v3_)]))), len(d_265_v92_)), d_160_v1_, (d_271_v98_) | (_dafny.Map({d_160_v1_: d_270_v97_})), default__.fm15(d_160_v1_, d_162_v3_, len(d_265_v92_), -254, d_164_globalState_), ((d_273_v100_)[(d_159_v0_)[default__.safeIndex(947, (d_159_v0_).length(0))]] if ((d_159_v0_)[default__.safeIndex(947, (d_159_v0_).length(0))]) in (d_273_v100_) else d_272_v99_))
            d_274_v101_ = nw38_
            d_275_v102_: _dafny.Seq
            d_275_v102_ = _dafny.SeqWithoutIsStrInference([d_162_v3_])
            d_275_v102_ = d_275_v102_
            d_259_v86_ = d_259_v86_
            d_276_v103_: _dafny.MultiSet
            d_276_v103_ = _dafny.MultiSet([237, d_162_v3_])
            d_277_v104_: int
            out12_: int
            out12_ = (d_274_v101_).m4(d_162_v3_, (594) + (d_162_v3_), d_265_v92_, ((d_276_v103_).set((0) - ((d_159_v0_)[default__.safeIndex(947, (d_159_v0_).length(0))]), default__.abs(len(d_163_v4_)))).intersection(d_276_v103_), d_164_globalState_)
            d_277_v104_ = out12_
        elif True:
            d_278_v105_: _dafny.Seq
            d_278_v105_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_162_v3_: d_160_v1_}))])
            d_278_v105_ = d_278_v105_
            d_279_v106_: _dafny.Seq
            d_279_v106_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isu"))
            d_280_v107_: _dafny.MultiSet
            d_280_v107_ = _dafny.MultiSet([631])
            d_281_v108_: D7
            d_281_v108_ = D7_DC20(d_280_v107_)
            d_282_v109_: D11
            d_282_v109_ = D11_DC32(780, d_162_v3_)
            d_162_v3_ = default__.safeModuloInt(default__.safeDivisionInt(len(d_279_v106_), ((d_281_v108_).cf30).cardinality), (d_282_v109_).cf47)
            d_283_v110_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.CodePoint('D'), 19)
            d_283_v110_ = nw39_
            d_284_v111_: str
            d_284_v111_ = _dafny.CodePoint('g')
            index24_ = default__.safeIndex(673, (d_283_v110_).length(0))
            (d_283_v110_)[index24_] = d_284_v111_
            index25_ = default__.safeIndex(673, (d_283_v110_).length(0))
            (d_283_v110_)[index25_] = (d_284_v111_ if d_160_v1_ else d_284_v111_)
            d_285_v112_: C7
            nw40_ = C7()
            nw40_.ctor__()
            d_285_v112_ = nw40_
            d_286_v113_: _dafny.Array
            def lambda17_(d_287_v107_):
                def lambda18_(d_288_i8_):
                    return d_287_v107_

                return lambda18_

            init9_ = lambda17_(d_280_v107_)
            nw41_ = _dafny.Array(None, 7)
            for i0_9_ in range(nw41_.length(0)):
                nw41_[i0_9_] = init9_(i0_9_)
            d_286_v113_ = nw41_
            d_289_v114_: _dafny.Map
            d_289_v114_ = _dafny.Map({d_160_v1_: True})
            d_290_v115_: _dafny.Map
            d_290_v115_ = _dafny.Map({d_160_v1_: d_289_v114_})
            d_291_v116_: _dafny.Seq
            d_291_v116_ = _dafny.SeqWithoutIsStrInference([d_160_v1_])
            d_292_v117_: T1
            nw42_ = C5()
            nw42_.ctor__(d_160_v1_, d_286_v113_, d_290_v115_, d_291_v116_)
            d_292_v117_ = nw42_
            d_293_v118_: _dafny.Map
            d_293_v118_ = _dafny.Map({d_292_v117_: d_285_v112_})
            d_294_v119_: _dafny.Array
            nw43_ = _dafny.Array(None, 17)
            nw43_[int(0)] = d_285_v112_
            nw43_[int(1)] = d_285_v112_
            nw43_[int(2)] = d_285_v112_
            nw43_[int(3)] = d_285_v112_
            nw43_[int(4)] = d_285_v112_
            nw43_[int(5)] = d_285_v112_
            nw43_[int(6)] = d_285_v112_
            nw43_[int(7)] = d_285_v112_
            nw43_[int(8)] = ((d_293_v118_)[d_292_v117_] if (d_292_v117_) in (d_293_v118_) else d_285_v112_)
            nw43_[int(9)] = d_285_v112_
            nw43_[int(10)] = d_285_v112_
            nw43_[int(11)] = d_285_v112_
            nw43_[int(12)] = d_285_v112_
            nw43_[int(13)] = d_285_v112_
            nw43_[int(14)] = d_285_v112_
            nw43_[int(15)] = d_285_v112_
            nw43_[int(16)] = d_285_v112_
            d_294_v119_ = nw43_
            d_295_v120_: _dafny.Set
            d_295_v120_ = _dafny.Set({d_294_v119_, d_294_v119_, d_294_v119_, d_294_v119_, d_294_v119_})
            d_295_v120_ = _dafny.Set({d_294_v119_})
            d_296_v121_: _dafny.Array
            nw44_ = _dafny.Array(None, 11)
            d_296_v121_ = nw44_
            d_297_v122_: C3
            nw45_ = C3()
            nw45_.ctor__(d_160_v1_, d_292_v117_.f12, d_160_v1_, (d_292_v117_).f13, _dafny.SeqWithoutIsStrInference([d_160_v1_, d_160_v1_, d_160_v1_]))
            d_297_v122_ = nw45_
            index26_ = default__.safeIndex(19, (d_296_v121_).length(0))
            (d_296_v121_)[index26_] = d_297_v122_
            index27_ = default__.safeIndex(19, (d_296_v121_).length(0))
            (d_296_v121_)[index27_] = d_297_v122_
        d_298_i9_: int
        d_298_i9_ = 0
        with _dafny.label("1"):
            while d_160_v1_:
                with _dafny.c_label("1"):
                    if (d_298_i9_) >= (100):
                        raise _dafny.Break("1")
                    d_298_i9_ = (d_298_i9_) + (1)
                    d_299_v123_: str
                    d_299_v123_ = _dafny.CodePoint('n')
                    d_300_v124_: _dafny.Map
                    d_300_v124_ = _dafny.Map({d_299_v123_: d_162_v3_})
                    d_301_v125_: _dafny.Seq
                    d_301_v125_ = _dafny.SeqWithoutIsStrInference([42])
                    d_302_v126_: _dafny.Map
                    d_302_v126_ = _dafny.Map({d_300_v124_: d_301_v125_})
                    d_162_v3_ = len(d_302_v126_)
                    d_162_v3_ = (d_162_v3_) - (d_162_v3_)
                    d_303_v127_: _dafny.Array
                    def lambda19_(d_304_v125_):
                        def lambda20_(d_305_i10_):
                            return _dafny.MultiSet(d_304_v125_)

                        return lambda20_

                    init10_ = lambda19_(d_301_v125_)
                    nw46_ = _dafny.Array(None, 12)
                    for i0_10_ in range(nw46_.length(0)):
                        nw46_[i0_10_] = init10_(i0_10_)
                    d_303_v127_ = nw46_
                    d_306_v128_: C5
                    nw47_ = C5()
                    nw47_.ctor__(d_160_v1_, d_303_v127_, _dafny.Map({True: _dafny.Map({d_160_v1_: not(d_160_v1_)})}), _dafny.SeqWithoutIsStrInference([not(not(d_160_v1_)), d_160_v1_]))
                    d_306_v128_ = nw47_
                    d_307_v129_: _dafny.Array
                    nw48_ = _dafny.Array(None, 21)
                    nw48_[int(0)] = d_306_v128_
                    nw48_[int(1)] = d_306_v128_
                    nw48_[int(2)] = d_306_v128_
                    nw48_[int(3)] = d_306_v128_
                    nw48_[int(4)] = d_306_v128_
                    nw48_[int(5)] = d_306_v128_
                    nw48_[int(6)] = d_306_v128_
                    nw48_[int(7)] = d_306_v128_
                    nw48_[int(8)] = d_306_v128_
                    nw48_[int(9)] = d_306_v128_
                    nw48_[int(10)] = d_306_v128_
                    nw48_[int(11)] = d_306_v128_
                    nw48_[int(12)] = d_306_v128_
                    nw48_[int(13)] = d_306_v128_
                    nw48_[int(14)] = d_306_v128_
                    nw48_[int(15)] = d_306_v128_
                    nw48_[int(16)] = (d_306_v128_ if not((d_306_v128_).f19) else d_306_v128_)
                    nw48_[int(17)] = d_306_v128_
                    nw48_[int(18)] = d_306_v128_
                    nw48_[int(19)] = (D22_DC59(d_306_v128_)).cf91
                    nw48_[int(20)] = d_306_v128_
                    d_307_v129_ = nw48_
                    nw49_ = _dafny.Array(None, 6)
                    nw49_[int(0)] = d_306_v128_
                    nw49_[int(1)] = d_306_v128_
                    nw49_[int(2)] = d_306_v128_
                    nw49_[int(3)] = d_306_v128_
                    nw49_[int(4)] = d_306_v128_
                    nw49_[int(5)] = (d_306_v128_ if (d_306_v128_).f19 else d_306_v128_)
                    d_307_v129_ = nw49_
                    d_308_v130_: _dafny.Seq
                    d_308_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqmjqhts"))
                    d_309_v131_: _dafny.Seq
                    d_309_v131_ = _dafny.SeqWithoutIsStrInference([d_308_v130_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "th"))])
                    d_310_v132_: _dafny.Seq
                    d_310_v132_ = _dafny.SeqWithoutIsStrInference([not((d_306_v128_).f19), (_dafny.MultiSet(d_309_v131_)).issubset(_dafny.MultiSet([d_308_v130_]))])
                    d_311_v133_: _dafny.Map
                    d_311_v133_ = _dafny.Map({d_160_v1_: d_160_v1_})
                    d_312_v134_: _dafny.Map
                    d_312_v134_ = _dafny.Map({d_160_v1_: (d_311_v133_).set((d_306_v128_).f19, (d_306_v128_).f19)})
                    d_313_v135_: C2
                    nw50_ = C2()
                    nw50_.ctor__(_dafny.SeqWithoutIsStrInference([d_162_v3_ for d_314_i11_ in range(default__.abs(485))]), (d_306_v128_).f19, d_312_v134_, d_310_v132_, d_303_v127_)
                    d_313_v135_ = nw50_
                    d_315_v136_: _dafny.Map
                    d_315_v136_ = _dafny.Map({d_159_v0_: d_313_v135_})
                    d_316_v137_: D23
                    d_316_v137_ = D23_DC63(d_313_v135_)
                    d_317_v138_: _dafny.Array
                    nw51_ = _dafny.Array(None, 22)
                    nw51_[int(0)] = d_313_v135_
                    nw51_[int(1)] = ((d_315_v136_)[d_159_v0_] if (d_159_v0_) in (d_315_v136_) else d_313_v135_)
                    nw51_[int(2)] = d_313_v135_
                    nw51_[int(3)] = d_313_v135_
                    nw51_[int(4)] = (d_316_v137_).cf98
                    nw51_[int(5)] = d_313_v135_
                    nw51_[int(6)] = d_313_v135_
                    nw51_[int(7)] = d_313_v135_
                    nw51_[int(8)] = d_313_v135_
                    nw51_[int(9)] = d_313_v135_
                    nw51_[int(10)] = d_313_v135_
                    nw51_[int(11)] = d_313_v135_
                    nw51_[int(12)] = d_313_v135_
                    nw51_[int(13)] = d_313_v135_
                    nw51_[int(14)] = d_313_v135_
                    nw51_[int(15)] = d_313_v135_
                    nw51_[int(16)] = (d_316_v137_).cf98
                    nw51_[int(17)] = d_313_v135_
                    nw51_[int(18)] = d_313_v135_
                    nw51_[int(19)] = d_313_v135_
                    nw51_[int(20)] = d_313_v135_
                    nw51_[int(21)] = d_313_v135_
                    d_317_v138_ = nw51_
                    index28_ = default__.safeIndex(338, (d_317_v138_).length(0))
                    (d_317_v138_)[index28_] = d_313_v135_
                    index29_ = default__.safeIndex(338, (d_317_v138_).length(0))
                    rhs25_ = ((d_310_v132_) + (d_310_v132_)) + ((d_310_v132_) + (d_310_v132_))
                    rhs26_ = d_313_v135_
                    rhs27_ = 306
                    rhs28_ = not(True)
                    rhs29_ = d_310_v132_
                    lhs13_ = d_317_v138_
                    lhs14_ = default__.safeIndex(338, (d_317_v138_).length(0))
                    lhs15_ = d_164_globalState_
                    d_310_v132_ = rhs25_
                    lhs13_[lhs14_] = rhs26_
                    d_162_v3_ = rhs27_
                    lhs15_.f4 = rhs28_
                    d_310_v132_ = rhs29_
                    pass
            pass
        d_318_v139_: _dafny.Map
        d_318_v139_ = _dafny.Map({True: d_162_v3_})
        d_319_v140_: _dafny.Map
        d_319_v140_ = _dafny.Map({d_162_v3_: d_318_v139_})
        d_162_v3_ = (default__.fm0(d_160_v1_, d_319_v140_, d_160_v1_, d_164_globalState_) if d_160_v1_ else d_162_v3_)
        d_320_v141_: _dafny.Seq
        d_320_v141_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_321_v142_: _dafny.Seq
        d_321_v142_ = _dafny.SeqWithoutIsStrInference([d_159_v0_])
        d_322_v143_: _dafny.Seq
        d_322_v143_ = _dafny.SeqWithoutIsStrInference([d_321_v142_, d_321_v142_, d_321_v142_, d_321_v142_, d_321_v142_])
        hi3_ = len((d_322_v143_)[default__.safeIndex((0) - ((0) - (d_162_v3_)), len(d_322_v143_))])
        for d_323_i12_ in range(len(d_320_v141_), hi3_):
            d_324_v145_: _dafny.Array
            def lambda21_(d_325_v2_, d_326_v1_):
                def lambda22_(d_327_i13_):
                    def iife33_():
                        coll33_ = _dafny.Set()
                        compr_33_: _dafny.MultiSet
                        for compr_33_ in (_dafny.MultiSet([d_325_v2_, _dafny.MultiSet([d_326_v1_])])).Elements:
                            d_328_v144_: _dafny.MultiSet = compr_33_
                            if (d_328_v144_) in (_dafny.MultiSet([d_325_v2_, _dafny.MultiSet([d_326_v1_])])):
                                coll33_ = coll33_.union(_dafny.Set([d_328_v144_]))
                        return _dafny.Set(coll33_)
                    return (iife33_()
                    ) - (_dafny.Set({d_325_v2_}))

                return lambda22_

            init11_ = lambda21_(d_161_v2_, d_160_v1_)
            nw52_ = _dafny.Array(None, 3)
            for i0_11_ in range(nw52_.length(0)):
                nw52_[i0_11_] = init11_(i0_11_)
            d_324_v145_ = nw52_
            d_329_v146_: _dafny.Set
            d_329_v146_ = _dafny.Set({d_161_v2_})
            index30_ = default__.safeIndex(294, (d_324_v145_).length(0))
            (d_324_v145_)[index30_] = d_329_v146_
            index31_ = default__.safeIndex(294, (d_324_v145_).length(0))
            (d_324_v145_)[index31_] = d_329_v146_
            d_330_v147_: _dafny.Array
            nw53_ = _dafny.Array(None, 15)
            d_330_v147_ = nw53_
            d_331_v148_: C7
            nw54_ = C7()
            nw54_.ctor__()
            d_331_v148_ = nw54_
            index32_ = default__.safeIndex(213, (d_330_v147_).length(0))
            (d_330_v147_)[index32_] = d_331_v148_
            index33_ = default__.safeIndex(213, (d_330_v147_).length(0))
            (d_330_v147_)[index33_] = (d_331_v148_ if (d_160_v1_) == (not(d_160_v1_)) else d_331_v148_)
            if d_160_v1_:
                d_332_v149_: _dafny.MultiSet
                d_332_v149_ = _dafny.MultiSet([d_323_i12_])
                d_333_v150_: _dafny.Seq
                d_333_v150_ = _dafny.SeqWithoutIsStrInference([(0) - (d_323_i12_)])
                d_334_v151_: _dafny.Array
                nw55_ = _dafny.Array(None, 5)
                nw55_[int(0)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_162_v3_]))
                nw55_[int(1)] = d_332_v149_
                nw55_[int(2)] = d_332_v149_
                nw55_[int(3)] = d_332_v149_
                nw55_[int(4)] = _dafny.MultiSet(d_333_v150_)
                d_334_v151_ = nw55_
                d_335_v152_: _dafny.Map
                d_335_v152_ = _dafny.Map({d_160_v1_: d_160_v1_})
                d_336_v153_: _dafny.Map
                d_336_v153_ = _dafny.Map({d_160_v1_: d_335_v152_})
                d_337_v154_: _dafny.Seq
                d_337_v154_ = _dafny.SeqWithoutIsStrInference([d_160_v1_])
                d_338_v155_: C3
                nw56_ = C3()
                nw56_.ctor__(d_160_v1_, d_334_v151_, d_160_v1_, d_336_v153_, d_337_v154_)
                d_338_v155_ = nw56_
                d_339_v156_: D24
                d_339_v156_ = D24_DC65(d_338_v155_)
                d_338_v155_ = (d_339_v156_).cf99
                (d_164_globalState_).f4 = not(default__.fm2((d_338_v155_).f20, d_164_globalState_))
                d_340_v157_: _dafny.Seq
                d_341_v158_: bool
                d_342_v159_: int
                d_343_v160_: bool
                out13_: _dafny.Seq
                out14_: bool
                out15_: int
                out16_: bool
                out13_, out14_, out15_, out16_ = default__.m0((d_338_v155_).f20, D0_DC2(d_162_v3_, (d_338_v155_).f20), ((d_338_v155_).f20) or (d_160_v1_), d_164_globalState_)
                d_340_v157_ = out13_
                d_341_v158_ = out14_
                d_342_v159_ = out15_
                d_343_v160_ = out16_
                d_344_v161_: int
                d_345_v162_: int
                d_346_v163_: bool
                d_347_v164_: _dafny.Map
                out17_: int
                out18_: int
                out19_: bool
                out20_: _dafny.Map
                out17_, out18_, out19_, out20_ = (d_338_v155_).m9(True, not(False), d_320_v141_, d_164_globalState_)
                d_344_v161_ = out17_
                d_345_v162_ = out18_
                d_346_v163_ = out19_
                d_347_v164_ = out20_
                d_348_v165_: C2
                nw57_ = C2()
                nw57_.ctor__(d_340_v157_, (d_345_v162_) < (d_323_i12_), d_336_v153_, d_337_v154_, d_334_v151_)
                d_348_v165_ = nw57_
            elif True:
                d_162_v3_ = default__.fm0(not(d_160_v1_), d_319_v140_, (d_320_v141_) == (d_320_v141_), d_164_globalState_)
                d_349_v167_: _dafny.Array
                def lambda23_(d_350_i12_, d_351_v1_, d_352_v141_, d_353_v4_, d_354_v3_):
                    def lambda24_(d_355_i14_):
                        def iife34_():
                            coll34_ = _dafny.Map()
                            compr_34_: int
                            for compr_34_ in (d_353_v4_).keys.Elements:
                                d_357_v166_: int = compr_34_
                                if (d_357_v166_) in (d_353_v4_):
                                    coll34_[(d_357_v166_) * (len(_dafny.SeqWithoutIsStrInference([False])))] = d_351_v1_
                            return _dafny.Map(coll34_)
                        return (D20_DC54(_dafny.CodePoint('e'), _dafny.Map({d_350_i12_: d_351_v1_}), _dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_356_i15_ in range(default__.abs(895))])}), -472) if False else D20_DC54(_dafny.CodePoint('g'), iife34_()
, _dafny.Map({d_351_v1_: d_352_v141_}), d_354_v3_))

                    return lambda24_

                init12_ = lambda23_(d_323_i12_, d_160_v1_, d_320_v141_, d_163_v4_, d_162_v3_)
                nw58_ = _dafny.Array(None, 15)
                for i0_12_ in range(nw58_.length(0)):
                    nw58_[i0_12_] = init12_(i0_12_)
                d_349_v167_ = nw58_
                d_358_v168_: _dafny.MultiSet
                d_358_v168_ = _dafny.MultiSet([d_323_i12_, d_323_i12_])
                index34_ = default__.safeIndex(915, (d_349_v167_).length(0))
                (d_349_v167_)[index34_] = default__.fm69(141, ((d_358_v168_)[d_323_i12_] if (d_323_i12_) in (d_358_v168_) else d_162_v3_), d_162_v3_, d_164_globalState_)
                d_359_v169_: _dafny.Map
                d_359_v169_ = _dafny.Map({(0) - (d_323_i12_): True})
                d_360_v170_: _dafny.Map
                d_360_v170_ = _dafny.Map({d_160_v1_: d_320_v141_})
                d_361_v171_: _dafny.Set
                d_361_v171_ = _dafny.Set({d_162_v3_})
                d_362_v172_: _dafny.Set
                d_362_v172_ = _dafny.Set({d_361_v171_})
                d_363_v173_: D20
                d_363_v173_ = D20_DC54(_dafny.CodePoint('p'), d_359_v169_, d_360_v170_, len(d_362_v172_))
                index35_ = default__.safeIndex(915, (d_349_v167_).length(0))
                (d_349_v167_)[index35_] = d_363_v173_
                d_162_v3_ = d_323_i12_
                d_364_v174_: _dafny.Array
                def lambda25_(d_365_v141_):
                    def lambda26_(d_366_i16_):
                        return d_365_v141_

                    return lambda26_

                init13_ = lambda25_(d_320_v141_)
                nw59_ = _dafny.Array(None, 28)
                for i0_13_ in range(nw59_.length(0)):
                    nw59_[i0_13_] = init13_(i0_13_)
                d_364_v174_ = nw59_
                d_367_v175_: _dafny.Seq
                d_367_v175_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, d_162_v3_, -477, d_162_v3_, d_323_i12_])
                d_368_v176_: _dafny.Seq
                d_368_v176_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_323_i12_ for d_369_i17_ in range(default__.abs(898))]), d_367_v175_, d_367_v175_, d_367_v175_, d_367_v175_])
                d_370_v177_: _dafny.Map
                d_370_v177_ = _dafny.Map({772: d_358_v168_})
                d_371_v178_: _dafny.Array
                nw60_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_371_v178_ = nw60_
                d_372_v179_: T4
                nw61_ = C6()
                nw61_.ctor__((d_367_v175_) + ((d_368_v176_)[default__.safeIndex(d_162_v3_, len(d_368_v176_))]), (d_160_v1_) and (d_160_v1_), _dafny.Map({True: default__.fm59(((d_370_v177_)[d_323_i12_] if (d_323_i12_) in (d_370_v177_) else d_358_v168_), d_160_v1_, d_160_v1_, d_162_v3_, d_164_globalState_)}), _dafny.SeqWithoutIsStrInference([d_160_v1_]), d_371_v178_)
                d_372_v179_ = nw61_
                index36_ = default__.safeIndex(903, (d_159_v0_).length(0))
                (d_159_v0_)[index36_] = 985
                index37_ = default__.safeIndex(903, (d_159_v0_).length(0))
                rhs30_ = not(d_160_v1_)
                rhs31_ = d_364_v174_
                rhs32_ = d_372_v179_
                rhs33_ = d_323_i12_
                rhs34_ = d_320_v141_
                lhs16_ = d_159_v0_
                lhs17_ = default__.safeIndex(903, (d_159_v0_).length(0))
                d_160_v1_ = rhs30_
                d_364_v174_ = rhs31_
                d_372_v179_ = rhs32_
                lhs16_[lhs17_] = rhs33_
                d_320_v141_ = rhs34_
                d_373_v180_: _dafny.Map
                d_373_v180_ = _dafny.Map({d_160_v1_: d_160_v1_})
                d_374_v181_: _dafny.Map
                d_374_v181_ = _dafny.Map({d_160_v1_: d_373_v180_})
                d_375_v182_: _dafny.Seq
                d_375_v182_ = _dafny.SeqWithoutIsStrInference([d_160_v1_, d_160_v1_, d_160_v1_, d_160_v1_])
                d_376_v183_: C3
                nw62_ = C3()
                nw62_.ctor__(d_160_v1_, d_371_v178_, d_160_v1_, (d_374_v181_) | (default__.fm60(len(d_373_v180_), d_164_globalState_)), (d_375_v182_) + (_dafny.SeqWithoutIsStrInference([d_160_v1_])))
                d_376_v183_ = nw62_
                d_376_v183_ = d_376_v183_
            d_377_v184_: int
            d_378_v185_: int
            d_379_v186_: D0
            out21_: int
            out22_: int
            out23_: D0
            out21_, out22_, out23_ = (d_331_v148_).m2(993, d_164_globalState_)
            d_377_v184_ = out21_
            d_378_v185_ = out22_
            d_379_v186_ = out23_
        d_380_v187_: _dafny.Array
        nw63_ = _dafny.Array(_dafny.MultiSet({}), 27)
        d_380_v187_ = nw63_
        d_381_v188_: _dafny.Map
        d_381_v188_ = _dafny.Map({d_160_v1_: d_160_v1_})
        d_382_v189_: _dafny.Seq
        d_382_v189_ = _dafny.SeqWithoutIsStrInference([d_160_v1_])
        d_383_v190_: T2
        nw64_ = C3()
        nw64_.ctor__(d_160_v1_, d_380_v187_, d_160_v1_, _dafny.Map({d_160_v1_: d_381_v188_}), d_382_v189_)
        d_383_v190_ = nw64_
        d_384_v191_: D14
        d_384_v191_ = D14_DC40(d_383_v190_)
        pat_let_tv0_ = d_383_v190_
        def iife35_(_pat_let0_0):
            def iife36_(d_385_dt__update__tmp_h0_):
                def iife37_(_pat_let1_0):
                    def iife38_(d_386_dt__update_hcf63_h0_):
                        return D14_DC40(d_386_dt__update_hcf63_h0_)
                    return iife38_(_pat_let1_0)
                return iife37_(pat_let_tv0_)
            return iife36_(_pat_let0_0)
        source4_ = iife35_(d_384_v191_)
        if source4_.is_DC41:
            d_387___mcc_h0_ = source4_.cf64
            d_388___mcc_h1_ = source4_.cf65
            d_389_cf65_ = d_388___mcc_h1_
            d_390_cf64_ = d_387___mcc_h0_
            if (d_383_v190_).f15:
                (d_164_globalState_).f4 = (d_162_v3_) < (d_389_cf65_)
                d_320_v141_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_391_i18_ in range(default__.abs(243))])) + (d_320_v141_)
                d_392_v192_: _dafny.Array
                nw65_ = _dafny.Array(None, 8)
                nw65_[int(0)] = False
                nw65_[int(1)] = (d_383_v190_).f15
                nw65_[int(2)] = default__.fm2((d_383_v190_).f15, d_164_globalState_)
                nw65_[int(3)] = not(d_160_v1_)
                nw65_[int(4)] = d_160_v1_
                nw65_[int(5)] = False
                nw65_[int(6)] = False
                nw65_[int(7)] = (d_383_v190_).f15
                d_392_v192_ = nw65_
                d_393_v193_: _dafny.Map
                d_393_v193_ = _dafny.Map({d_392_v192_: d_160_v1_})
                d_394_v194_: _dafny.Array
                nw66_ = _dafny.Array(None, 17)
                nw66_[int(0)] = d_160_v1_
                nw66_[int(1)] = ((d_393_v193_)[d_392_v192_] if (d_392_v192_) in (d_393_v193_) else (d_383_v190_).f15)
                nw66_[int(2)] = (d_383_v190_).f15
                nw66_[int(3)] = (d_383_v190_).f15
                nw66_[int(4)] = (d_383_v190_).f15
                nw66_[int(5)] = (d_383_v190_).f15
                nw66_[int(6)] = not((d_320_v141_) != (d_390_cf64_))
                nw66_[int(7)] = d_160_v1_
                nw66_[int(8)] = d_160_v1_
                nw66_[int(9)] = default__.fm2((d_383_v190_).f15, d_164_globalState_)
                nw66_[int(10)] = ((d_383_v190_).f15) == ((d_383_v190_).f15)
                nw66_[int(11)] = (d_383_v190_).f15
                nw66_[int(12)] = (d_383_v190_).f15
                nw66_[int(13)] = (d_383_v190_).f15
                nw66_[int(14)] = (d_383_v190_).f15
                nw66_[int(15)] = (d_383_v190_).f15
                nw66_[int(16)] = (d_383_v190_).f15
                d_394_v194_ = nw66_
                d_395_v195_: _dafny.MultiSet
                d_395_v195_ = _dafny.MultiSet([len(d_320_v141_), d_162_v3_])
                d_396_v196_: _dafny.Map
                d_396_v196_ = _dafny.Map({d_395_v195_: d_162_v3_})
                d_397_v197_: _dafny.Set
                d_397_v197_ = _dafny.Set({d_389_cf65_})
                d_398_v198_: _dafny.Set
                d_398_v198_ = _dafny.Set({_dafny.CodePoint('g')})
                d_399_v199_: D5
                d_399_v199_ = D5_DC16((d_383_v190_).f15)
                nw67_ = _dafny.Array(None, 20)
                nw67_[int(0)] = not((len(d_396_v196_)) not in (d_397_v197_))
                nw67_[int(1)] = d_160_v1_
                nw67_[int(2)] = not((d_383_v190_).f15)
                nw67_[int(3)] = (d_162_v3_) < ((0) - (len(_dafny.Map({(d_161_v2_).cardinality: d_398_v198_}))))
                nw67_[int(4)] = not(d_160_v1_)
                nw67_[int(5)] = (d_383_v190_).f15
                nw67_[int(6)] = not((d_383_v190_).f15)
                nw67_[int(7)] = (d_383_v190_).f15
                nw67_[int(8)] = (d_383_v190_).f15
                nw67_[int(9)] = (d_383_v190_).f15
                nw67_[int(10)] = d_160_v1_
                nw67_[int(11)] = (d_383_v190_).f15
                nw67_[int(12)] = ((d_383_v190_).f15) == (not(d_160_v1_))
                nw67_[int(13)] = d_160_v1_
                nw67_[int(14)] = d_160_v1_
                nw67_[int(15)] = ((d_383_v190_).f15 if d_160_v1_ else (d_383_v190_).f15)
                nw67_[int(16)] = (d_383_v190_).f15
                nw67_[int(17)] = (False) and (d_160_v1_)
                nw67_[int(18)] = (d_383_v190_).f15
                nw67_[int(19)] = (d_399_v199_).cf26
                d_394_v194_ = nw67_
                d_389_cf65_ = (d_162_v3_ if (d_390_cf64_) != (d_320_v141_) else (d_383_v190_).fm8(default__.fm2(True, d_164_globalState_), d_160_v1_, d_164_globalState_))
                d_162_v3_ = d_389_cf65_
            elif True:
                d_400_v200_: _dafny.Array
                def lambda27_(d_401_v1_):
                    def lambda28_(d_402_i19_):
                        return d_401_v1_

                    return lambda28_

                init14_ = lambda27_(d_160_v1_)
                nw68_ = _dafny.Array(None, 2)
                for i0_14_ in range(nw68_.length(0)):
                    nw68_[i0_14_] = init14_(i0_14_)
                d_400_v200_ = nw68_
                index38_ = default__.safeIndex(430, (d_400_v200_).length(0))
                (d_400_v200_)[index38_] = default__.fm2(False, d_164_globalState_)
                d_403_v201_: _dafny.Array
                def lambda29_(d_404_i20_):
                    return _dafny.CodePoint('n')

                init15_ = lambda29_
                nw69_ = _dafny.Array(None, 15)
                for i0_15_ in range(nw69_.length(0)):
                    nw69_[i0_15_] = init15_(i0_15_)
                d_403_v201_ = nw69_
                d_405_v202_: _dafny.Seq
                d_405_v202_ = _dafny.SeqWithoutIsStrInference([d_403_v201_])
                d_406_v203_: _dafny.Seq
                d_406_v203_ = d_405_v202_
                index39_ = default__.safeIndex(430, (d_400_v200_).length(0))
                rhs35_ = (d_383_v190_).f15
                rhs36_ = (0) - (d_389_cf65_)
                rhs37_ = (d_383_v190_).f15
                rhs38_ = (d_383_v190_).f15
                rhs39_ = ((d_405_v202_) + ((d_405_v202_).set(default__.safeIndex(d_389_cf65_, len(d_405_v202_)), d_403_v201_))) < ((d_406_v203_))
                lhs18_ = d_164_globalState_
                lhs19_ = d_400_v200_
                lhs20_ = default__.safeIndex(430, (d_400_v200_).length(0))
                lhs21_ = d_164_globalState_
                d_160_v1_ = rhs35_
                d_389_cf65_ = rhs36_
                lhs18_.f4 = rhs37_
                lhs19_[lhs20_] = rhs38_
                lhs21_.f4 = rhs39_
                d_407_v204_: C4
                nw70_ = C4()
                nw70_.ctor__()
                d_407_v204_ = nw70_
                d_408_v205_: _dafny.Array
                nw71_ = _dafny.Array(None, 6)
                nw71_[int(0)] = d_407_v204_
                nw71_[int(1)] = d_407_v204_
                nw71_[int(2)] = d_407_v204_
                nw71_[int(3)] = d_407_v204_
                nw71_[int(4)] = d_407_v204_
                nw71_[int(5)] = d_407_v204_
                d_408_v205_ = nw71_
                d_409_v206_: _dafny.Map
                d_409_v206_ = _dafny.Map({default__.fm2((d_383_v190_).f15, d_164_globalState_): d_407_v204_})
                index40_ = default__.safeIndex(324, (d_408_v205_).length(0))
                (d_408_v205_)[index40_] = ((d_409_v206_)[(d_400_v200_)[default__.safeIndex(430, (d_400_v200_).length(0))]] if ((d_400_v200_)[default__.safeIndex(430, (d_400_v200_).length(0))]) in (d_409_v206_) else d_407_v204_)
                index41_ = default__.safeIndex(324, (d_408_v205_).length(0))
                (d_408_v205_)[index41_] = d_407_v204_
                d_162_v3_ = default__.safeDivisionInt((((d_161_v2_)[(d_383_v190_).f15] if ((d_383_v190_).f15) in (d_161_v2_) else (d_383_v190_).fm5(False, d_164_globalState_))) * (len(d_320_v141_)), d_162_v3_)
                d_389_cf65_ = d_389_cf65_
                d_162_v3_ = default__.safeDivisionInt(d_389_cf65_, d_162_v3_)
            d_410_v207_: T0
            nw72_ = C5()
            nw72_.ctor__((d_383_v190_).f15, d_383_v190_.f12, ((d_383_v190_).f13).set((d_383_v190_).f15, d_381_v188_), _dafny.SeqWithoutIsStrInference([d_160_v1_]))
            d_410_v207_ = nw72_
            d_411_v208_: _dafny.Map
            d_411_v208_ = _dafny.Map({d_410_v207_: d_389_cf65_})
            d_162_v3_ = default__.safeModuloInt(d_389_cf65_, default__.safeModuloInt(len((d_411_v208_).set(d_410_v207_, d_389_cf65_)), d_162_v3_))
            d_412_v209_: _dafny.Array
            nw73_ = _dafny.Array(None, 11)
            d_412_v209_ = nw73_
            index42_ = default__.safeIndex(978, (d_412_v209_).length(0))
            (d_412_v209_)[index42_] = d_383_v190_
            d_413_v210_: D21
            d_413_v210_ = D21_DC56(d_383_v190_)
            index43_ = default__.safeIndex(978, (d_412_v209_).length(0))
            rhs40_ = (d_413_v210_).cf85
            rhs41_ = d_159_v0_
            rhs42_ = (d_160_v1_) == (d_160_v1_)
            rhs43_ = len(d_320_v141_)
            rhs44_ = d_389_cf65_
            lhs22_ = d_412_v209_
            lhs23_ = default__.safeIndex(978, (d_412_v209_).length(0))
            lhs24_ = d_164_globalState_
            lhs22_[lhs23_] = rhs40_
            d_159_v0_ = rhs41_
            lhs24_.f4 = rhs42_
            d_389_cf65_ = rhs43_
            d_162_v3_ = rhs44_
            d_414_v211_: _dafny.Array
            nw74_ = _dafny.Array(False, 2)
            d_414_v211_ = nw74_
            d_415_v212_: _dafny.Set
            d_415_v212_ = _dafny.Set({d_414_v211_})
            d_416_v213_: D22
            d_416_v213_ = D22_DC62(d_414_v211_)
            d_417_v214_: _dafny.Seq
            d_417_v214_ = _dafny.SeqWithoutIsStrInference([d_414_v211_, d_414_v211_, (d_416_v213_).cf97, d_414_v211_, d_414_v211_])
            (d_164_globalState_).f4 = (d_414_v211_) in ((d_415_v212_).intersection(_dafny.Set({(d_417_v214_)[default__.safeIndex(d_162_v3_, len(d_417_v214_))]})))
        elif source4_.is_DC42:
            d_418_v215_: D0
            d_418_v215_ = D0_DC2(d_162_v3_, (d_383_v190_).f15)
            d_419_v216_: _dafny.Seq
            d_420_v217_: bool
            d_421_v218_: int
            d_422_v219_: bool
            out24_: _dafny.Seq
            out25_: bool
            out26_: int
            out27_: bool
            out24_, out25_, out26_, out27_ = default__.m0(d_160_v1_, d_418_v215_, (d_383_v190_).f15, d_164_globalState_)
            d_419_v216_ = out24_
            d_420_v217_ = out25_
            d_421_v218_ = out26_
            d_422_v219_ = out27_
            d_423_v220_: _dafny.Seq
            d_423_v220_ = _dafny.SeqWithoutIsStrInference([d_383_v190_])
            (d_164_globalState_).f4 = (d_421_v218_) <= (len(d_423_v220_))
            d_421_v218_ = (d_383_v190_).fm5((d_383_v190_).f15, d_164_globalState_)
            d_424_v221_: D8
            d_424_v221_ = D8_DC25(not((d_383_v190_).f15), len(d_419_v216_), d_421_v218_, d_318_v139_)
            d_425_v222_: D8
            d_425_v222_ = D8_DC27(d_424_v221_)
            source5_ = d_425_v222_
            if source5_.is_DC24:
                d_426___mcc_h4_ = source5_.cf36
                d_427_cf36_ = d_426___mcc_h4_
                index44_ = default__.safeIndex(226, (d_159_v0_).length(0))
                (d_159_v0_)[index44_] = (d_383_v190_).fm5((d_383_v190_).f15, d_164_globalState_)
                index45_ = default__.safeIndex(226, (d_159_v0_).length(0))
                (d_159_v0_)[index45_] = (0) - (d_162_v3_)
                d_428_v223_: _dafny.Array
                nw75_ = _dafny.Array(False, 25)
                d_428_v223_ = nw75_
                index46_ = default__.safeIndex(139, (d_428_v223_).length(0))
                (d_428_v223_)[index46_] = (d_421_v218_) == (d_421_v218_)
                d_429_v224_: _dafny.Map
                d_429_v224_ = _dafny.Map({d_162_v3_: (d_383_v190_).f15})
                index47_ = default__.safeIndex(139, (d_428_v223_).length(0))
                (d_428_v223_)[index47_] = (default__.fm2((d_383_v190_).f15, d_164_globalState_)) == ((len(d_429_v224_)) == ((d_159_v0_)[default__.safeIndex(226, (d_159_v0_).length(0))]))
                d_421_v218_ = len(d_320_v141_)
                d_430_v225_: str
                d_430_v225_ = _dafny.CodePoint('m')
                d_431_v226_: _dafny.MultiSet
                d_431_v226_ = _dafny.MultiSet([(d_159_v0_)[default__.safeIndex(226, (d_159_v0_).length(0))]])
                index48_ = default__.safeIndex(226, (d_159_v0_).length(0))
                (d_159_v0_)[index48_] = (len((_dafny.Map({d_430_v225_: (d_383_v190_).f15})) | (_dafny.Map({d_430_v225_: (d_383_v190_).f15})))) + ((d_162_v3_) + ((d_431_v226_).cardinality))
            elif source5_.is_DC25:
                d_432___mcc_h5_ = source5_.cf37
                d_433___mcc_h6_ = source5_.cf38
                d_434___mcc_h7_ = source5_.cf39
                d_435___mcc_h8_ = source5_.cf40
                d_436_cf40_ = d_435___mcc_h8_
                d_437_cf39_ = d_434___mcc_h7_
                d_438_cf38_ = d_433___mcc_h6_
                d_439_cf37_ = d_432___mcc_h5_
                (d_164_globalState_).f4 = not(False)
                d_440_v227_: _dafny.Array
                def lambda30_(d_441_v141_):
                    def lambda31_(d_442_i21_):
                        return d_441_v141_

                    return lambda31_

                init16_ = lambda30_(d_320_v141_)
                nw76_ = _dafny.Array(None, 27)
                for i0_16_ in range(nw76_.length(0)):
                    nw76_[i0_16_] = init16_(i0_16_)
                d_440_v227_ = nw76_
                d_440_v227_ = d_440_v227_
                d_443_v228_: C6
                nw77_ = C6()
                nw77_.ctor__(_dafny.SeqWithoutIsStrInference([(d_161_v2_).cardinality]), d_160_v1_, (d_383_v190_).f13, (d_383_v190_).f14, (d_380_v187_ if not(d_439_cf37_) else d_383_v190_.f12))
                d_443_v228_ = nw77_
                d_437_cf39_ = d_421_v218_
            elif source5_.is_DC26:
                d_444___mcc_h9_ = source5_.cf41
                d_445_cf41_ = d_444___mcc_h9_
                d_421_v218_ = d_162_v3_
                d_421_v218_ = d_421_v218_
                d_446_v229_: _dafny.Map
                d_446_v229_ = _dafny.Map({_dafny.CodePoint('h'): (d_383_v190_).f15})
                d_447_v230_: str
                d_447_v230_ = _dafny.CodePoint('a')
                d_446_v229_ = _dafny.Map({d_447_v230_: False})
                d_162_v3_ = (d_421_v218_) * (71)
            elif source5_.is_DC23:
                d_448___mcc_h10_ = source5_.cf35
                d_449_cf35_ = d_448___mcc_h10_
                d_450_v231_: C3
                nw78_ = C3()
                nw78_.ctor__(d_160_v1_, d_383_v190_.f12, False, (d_383_v190_).f13, d_382_v189_)
                d_450_v231_ = nw78_
                d_451_v232_: _dafny.Map
                d_451_v232_ = _dafny.Map({d_450_v231_: (d_383_v190_).f15})
                d_452_v233_: _dafny.Seq
                d_453_v234_: bool
                d_454_v235_: int
                d_455_v236_: bool
                out28_: _dafny.Seq
                out29_: bool
                out30_: int
                out31_: bool
                out28_, out29_, out30_, out31_ = default__.m0((d_450_v231_) in (d_451_v232_), d_418_v215_, (d_450_v231_).f20, d_164_globalState_)
                d_452_v233_ = out28_
                d_453_v234_ = out29_
                d_454_v235_ = out30_
                d_455_v236_ = out31_
                (d_164_globalState_).f4 = (d_383_v190_).f15
                d_456_v237_: int
                d_457_v238_: int
                d_458_v239_: bool
                d_459_v240_: _dafny.Map
                out32_: int
                out33_: int
                out34_: bool
                out35_: _dafny.Map
                out32_, out33_, out34_, out35_ = (d_450_v231_).m9((d_421_v218_) > ((_dafny.MultiSet([d_450_v231_, d_450_v231_, d_450_v231_])).cardinality), (d_383_v190_).f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgwsvo")), d_164_globalState_)
                d_456_v237_ = out32_
                d_457_v238_ = out33_
                d_458_v239_ = out34_
                d_459_v240_ = out35_
                d_460_v241_: _dafny.Seq
                d_461_v242_: bool
                d_462_v243_: int
                d_463_v244_: bool
                out36_: _dafny.Seq
                out37_: bool
                out38_: int
                out39_: bool
                out36_, out37_, out38_, out39_ = default__.m0(d_160_v1_, D0_DC2(d_457_v238_, d_455_v236_), ((d_383_v190_).f14)[default__.safeIndex(d_454_v235_, len((d_383_v190_).f14))], d_164_globalState_)
                d_460_v241_ = out36_
                d_461_v242_ = out37_
                d_462_v243_ = out38_
                d_463_v244_ = out39_
            elif True:
                d_464___mcc_h11_ = source5_.cf42
                d_465_cf42_ = d_464___mcc_h11_
                (d_164_globalState_).f4 = (d_383_v190_).f15
                d_382_v189_ = (d_383_v190_).f14
                index49_ = default__.safeIndex(553, (d_159_v0_).length(0))
                (d_159_v0_)[index49_] = default__.safeDivisionInt(d_421_v218_, len(d_206_v40_))
                d_466_v245_: _dafny.Array
                nw79_ = _dafny.Array(False, 6)
                d_466_v245_ = nw79_
                index50_ = default__.safeIndex(359, (d_466_v245_).length(0))
                (d_466_v245_)[index50_] = (len(d_163_v4_)) == (d_162_v3_)
                index51_ = default__.safeIndex(553, (d_159_v0_).length(0))
                index52_ = default__.safeIndex(359, (d_466_v245_).length(0))
                rhs45_ = (d_318_v139_ if (d_383_v190_).f15 else ((_dafny.Map({d_160_v1_: d_421_v218_})).set((d_383_v190_).f15, d_162_v3_)) | (_dafny.Map({(d_383_v190_).f15: 969})))
                rhs46_ = not(d_160_v1_)
                rhs47_ = d_421_v218_
                rhs48_ = not(d_160_v1_)
                lhs25_ = d_159_v0_
                lhs26_ = default__.safeIndex(553, (d_159_v0_).length(0))
                lhs27_ = d_466_v245_
                lhs28_ = default__.safeIndex(359, (d_466_v245_).length(0))
                d_318_v139_ = rhs45_
                d_420_v217_ = rhs46_
                lhs25_[lhs26_] = rhs47_
                lhs27_[lhs28_] = rhs48_
                d_160_v1_ = (d_383_v190_).f15
        elif source4_.is_DC40:
            d_467___mcc_h2_ = source4_.cf63
            d_468_cf63_ = d_467___mcc_h2_
            d_160_v1_ = default__.fm2(True, d_164_globalState_)
            d_469_v246_: _dafny.Array
            nw80_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_469_v246_ = nw80_
            d_470_v247_: _dafny.Seq
            d_470_v247_ = _dafny.SeqWithoutIsStrInference([d_469_v246_])
            d_471_v248_: _dafny.Map
            d_471_v248_ = _dafny.Map({(d_470_v247_)[default__.safeIndex(d_162_v3_, len(d_470_v247_))]: (0) - (((d_163_v4_)[d_162_v3_] if (d_162_v3_) in (d_163_v4_) else (d_468_cf63_).fm5((d_383_v190_).f15, d_164_globalState_)))})
            d_472_v249_: _dafny.Set
            d_472_v249_ = _dafny.Set({d_162_v3_, len((d_468_cf63_).f14)})
            d_471_v248_ = (d_471_v248_).set(d_469_v246_, len((d_472_v249_).intersection(d_472_v249_)))
            d_473_v250_: bool
            d_474_v251_: int
            d_475_v252_: _dafny.Set
            d_476_v253_: bool
            out40_: bool
            out41_: int
            out42_: _dafny.Set
            out43_: bool
            out40_, out41_, out42_, out43_ = (d_383_v190_).m3(d_164_globalState_)
            d_473_v250_ = out40_
            d_474_v251_ = out41_
            d_475_v252_ = out42_
            d_476_v253_ = out43_
            (d_164_globalState_).f4 = d_473_v250_
        elif True:
            d_477___mcc_h3_ = source4_.cf66
            d_478_cf66_ = d_477___mcc_h3_
            d_479_v254_: _dafny.Seq
            d_479_v254_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - ((d_383_v190_).fm5((d_383_v190_).f15, d_164_globalState_)))])
            d_479_v254_ = (d_479_v254_) + ((d_479_v254_) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_383_v190_).f15: d_160_v1_})), d_162_v3_, d_162_v3_])).set(default__.safeIndex(len(d_320_v141_), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(d_383_v190_).f15: d_160_v1_})), d_162_v3_, d_162_v3_]))), d_162_v3_)))
            d_480_v255_: D12
            d_480_v255_ = D12_DC37(d_320_v141_, (d_383_v190_).f15)
            d_481_v256_: _dafny.Map
            d_481_v256_ = _dafny.Map({d_162_v3_: d_320_v141_})
            d_160_v1_ = (((d_480_v255_).cf59) + (((d_481_v256_)[(d_479_v254_)[default__.safeIndex(d_162_v3_, len(d_479_v254_))]] if ((d_479_v254_)[default__.safeIndex(d_162_v3_, len(d_479_v254_))]) in (d_481_v256_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_482_i22_ in range(default__.abs(503))])))) == (d_320_v141_)
            d_483_v257_: C0
            nw81_ = C0()
            nw81_.ctor__()
            d_483_v257_ = nw81_
            d_484_v258_: D21
            d_484_v258_ = D21_DC56(d_383_v190_)
            source6_ = d_484_v258_
            if source6_.is_DC56:
                d_485___mcc_h12_ = source6_.cf85
                d_486_cf85_ = d_485___mcc_h12_
                d_487_v259_: bool
                d_488_v260_: int
                d_489_v261_: _dafny.Set
                d_490_v262_: bool
                out44_: bool
                out45_: int
                out46_: _dafny.Set
                out47_: bool
                out44_, out45_, out46_, out47_ = (d_383_v190_).m3(d_164_globalState_)
                d_487_v259_ = out44_
                d_488_v260_ = out45_
                d_489_v261_ = out46_
                d_490_v262_ = out47_
                d_491_v264_: _dafny.Map
                d_491_v264_ = _dafny.Map({d_486_cf85_: d_488_v260_})
                d_492_v265_: _dafny.Seq
                d_492_v265_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_479_v254_): len(d_491_v264_)}), d_163_v4_])
                d_493_v266_: D11
                d_493_v266_ = D11_DC32(d_488_v260_, d_162_v3_)
                d_494_v267_: _dafny.Array
                nw82_ = _dafny.Array(None, 7)
                def iife39_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in _dafny.IntegerRange(-646, 701):
                        d_495_v263_: int = compr_35_
                        if ((-646) <= (d_495_v263_)) and ((d_495_v263_) < (701)):
                            coll35_[(d_495_v263_) + (814)] = len(d_479_v254_)
                    return _dafny.Map(coll35_)
                nw82_[int(0)] = _dafny.SeqWithoutIsStrInference([iife39_()
 for d_496_i23_ in range(default__.abs(395))])
                nw82_[int(1)] = d_492_v265_
                nw82_[int(2)] = (d_492_v265_) + (d_492_v265_)
                nw82_[int(3)] = d_492_v265_
                nw82_[int(4)] = (d_492_v265_) + (_dafny.SeqWithoutIsStrInference([d_163_v4_]))
                nw82_[int(5)] = (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_162_v3_: d_162_v3_}) for d_497_i24_ in range(default__.abs(168))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_493_v266_).cf46: (0) - (d_162_v3_)}), (d_492_v265_)[default__.safeIndex((0) - (d_488_v260_), len(d_492_v265_))], d_163_v4_]))
                nw82_[int(6)] = d_492_v265_
                d_494_v267_ = nw82_
                index53_ = default__.safeIndex(119, (d_494_v267_).length(0))
                (d_494_v267_)[index53_] = _dafny.SeqWithoutIsStrInference([d_163_v4_ for d_498_i25_ in range(default__.abs(384))])
                index54_ = default__.safeIndex(119, (d_494_v267_).length(0))
                (d_494_v267_)[index54_] = d_492_v265_
                d_499_v268_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_499_v268_ = nw83_
                d_500_v269_: _dafny.Map
                d_500_v269_ = _dafny.Map({(d_488_v260_) * (d_488_v260_): d_499_v268_})
                d_501_v270_: C3
                nw84_ = C3()
                nw84_.ctor__((d_383_v190_).f15, d_383_v190_.f12, (d_383_v190_).f15, (d_383_v190_).f13, (d_383_v190_).f14)
                d_501_v270_ = nw84_
                d_502_v271_: D24
                d_502_v271_ = D24_DC65(d_501_v270_)
                d_503_v272_: _dafny.Map
                d_503_v272_ = _dafny.Map({d_502_v271_: d_499_v268_})
                d_500_v269_ = (d_500_v269_).set(len(d_320_v141_), ((d_503_v272_)[d_502_v271_] if (d_502_v271_) in (d_503_v272_) else d_499_v268_))
                d_483_v257_ = d_483_v257_
            elif source6_.is_DC57:
                d_504___mcc_h13_ = source6_.cf86
                d_505___mcc_h14_ = source6_.cf87
                d_506___mcc_h15_ = source6_.cf88
                d_507_cf88_ = d_506___mcc_h15_
                d_508_cf87_ = d_505___mcc_h14_
                d_509_cf86_ = d_504___mcc_h13_
                d_510_v273_: C7
                nw85_ = C7()
                nw85_.ctor__()
                d_510_v273_ = nw85_
                nw86_ = C7()
                nw86_.ctor__()
                d_510_v273_ = nw86_
                d_160_v1_ = not(not((((d_161_v2_)[True] if (True) in (d_161_v2_) else (0) - (d_162_v3_))) <= (default__.safeDivisionInt(d_162_v3_, -986))))
                d_511_v274_: _dafny.Set
                d_511_v274_ = _dafny.Set({d_162_v3_})
                d_512_v275_: _dafny.Map
                d_512_v275_ = _dafny.Map({d_483_v257_: d_511_v274_})
                d_513_v276_: C2
                nw87_ = C2()
                nw87_.ctor__(d_479_v254_, d_160_v1_, _dafny.Map({(d_383_v190_).f15: _dafny.Map({not(True): (d_383_v190_).f15})}), d_382_v189_, d_383_v190_.f12)
                d_513_v276_ = nw87_
                d_514_v277_: _dafny.Seq
                d_514_v277_ = _dafny.SeqWithoutIsStrInference([d_513_v276_])
                d_512_v275_ = (d_512_v275_).set(d_483_v257_, _dafny.Set({len(d_514_v277_), ((d_161_v2_)[(d_383_v190_).f15] if ((d_383_v190_).f15) in (d_161_v2_) else (0) - (len((D14_DC41(d_320_v141_, len(d_511_v274_))).cf64))), (0) - (d_162_v3_)}))
                d_515_v278_: _dafny.Array
                nw88_ = _dafny.Array(False, 18)
                d_515_v278_ = nw88_
                d_516_v279_: D2
                d_516_v279_ = D2_DC6(d_515_v278_, d_515_v278_)
                d_517_v280_: _dafny.Array
                nw89_ = _dafny.Array(None, 15)
                nw89_[int(0)] = d_515_v278_
                nw89_[int(1)] = d_515_v278_
                nw89_[int(2)] = d_515_v278_
                nw89_[int(3)] = d_515_v278_
                nw89_[int(4)] = d_515_v278_
                nw89_[int(5)] = d_515_v278_
                nw89_[int(6)] = (d_516_v279_).cf9
                nw89_[int(7)] = d_515_v278_
                nw89_[int(8)] = d_515_v278_
                nw89_[int(9)] = d_515_v278_
                nw89_[int(10)] = d_515_v278_
                nw89_[int(11)] = d_515_v278_
                nw89_[int(12)] = d_515_v278_
                nw89_[int(13)] = d_515_v278_
                nw89_[int(14)] = d_515_v278_
                d_517_v280_ = nw89_
                index55_ = default__.safeIndex(475, (d_517_v280_).length(0))
                (d_517_v280_)[index55_] = d_515_v278_
                index56_ = default__.safeIndex(475, (d_517_v280_).length(0))
                (d_517_v280_)[index56_] = d_515_v278_
            elif source6_.is_DC58:
                d_518___mcc_h16_ = source6_.cf89
                d_519___mcc_h17_ = source6_.cf90
                d_520_cf90_ = d_519___mcc_h17_
                d_521_cf89_ = d_518___mcc_h16_
                d_163_v4_ = _dafny.Map({d_520_cf90_: d_520_cf90_})
                d_522_v281_: _dafny.Array
                def lambda32_(d_523_v141_):
                    def lambda33_(d_524_i26_):
                        return d_523_v141_

                    return lambda33_

                init17_ = lambda32_(d_320_v141_)
                nw90_ = _dafny.Array(None, 7)
                for i0_17_ in range(nw90_.length(0)):
                    nw90_[i0_17_] = init17_(i0_17_)
                d_522_v281_ = nw90_
                index57_ = default__.safeIndex(651, (d_522_v281_).length(0))
                (d_522_v281_)[index57_] = d_320_v141_
                d_525_v282_: _dafny.MultiSet
                d_525_v282_ = _dafny.MultiSet([d_162_v3_])
                index58_ = default__.safeIndex(651, (d_522_v281_).length(0))
                rhs49_ = (_dafny.MultiSet([d_520_cf90_])).issubset((d_525_v282_).set(d_520_cf90_, default__.abs(d_520_cf90_)))
                rhs50_ = not(False)
                rhs51_ = (0) - ((d_520_cf90_) * (d_162_v3_))
                rhs52_ = ((d_320_v141_) + (default__.fm21(d_160_v1_, d_164_globalState_))) + (d_320_v141_)
                lhs29_ = d_164_globalState_
                lhs30_ = d_522_v281_
                lhs31_ = default__.safeIndex(651, (d_522_v281_).length(0))
                d_521_cf89_ = rhs49_
                lhs29_.f4 = rhs50_
                d_520_cf90_ = rhs51_
                lhs30_[lhs31_] = rhs52_
                d_526_v283_: _dafny.Map
                d_526_v283_ = _dafny.Map({d_520_cf90_: (d_383_v190_).f15})
                d_527_v284_: _dafny.MultiSet
                d_527_v284_ = d_161_v2_
                d_528_v285_: _dafny.Seq
                d_528_v285_ = _dafny.SeqWithoutIsStrInference([d_527_v284_])
                d_529_v286_: str
                d_529_v286_ = _dafny.CodePoint('n')
                d_530_v287_: _dafny.Array
                d_531_v288_: int
                out48_: _dafny.Array
                out49_: int
                out48_, out49_ = (d_483_v257_).m7(default__.safeModuloInt(d_162_v3_, len(d_526_v283_)), d_528_v285_, d_162_v3_, d_529_v286_, d_164_globalState_)
                d_530_v287_ = out48_
                d_531_v288_ = out49_
                d_532_v289_: _dafny.Seq
                d_532_v289_ = _dafny.SeqWithoutIsStrInference([(d_522_v281_)[default__.safeIndex(651, (d_522_v281_).length(0))], (d_522_v281_)[default__.safeIndex(651, (d_522_v281_).length(0))], d_320_v141_, (d_522_v281_)[default__.safeIndex(651, (d_522_v281_).length(0))]])
                (d_164_globalState_).f4 = (d_532_v289_) <= ((d_532_v289_).set(default__.safeIndex(722, len(d_532_v289_)), (d_522_v281_)[default__.safeIndex(651, (d_522_v281_).length(0))]))
            elif True:
                d_533___mcc_h18_ = source6_.cf84
                d_534_cf84_ = d_533___mcc_h18_
                (d_164_globalState_).f4 = default__.fm2((d_383_v190_).f15, d_164_globalState_)
                index59_ = default__.safeIndex(90, (d_159_v0_).length(0))
                (d_159_v0_)[index59_] = d_162_v3_
                index60_ = default__.safeIndex(90, (d_159_v0_).length(0))
                (d_159_v0_)[index60_] = len(d_320_v141_)
                d_535_v290_: C5
                nw91_ = C5()
                nw91_.ctor__((d_162_v3_) < (d_162_v3_), d_383_v190_.f12, (d_383_v190_).f13, d_382_v189_)
                d_535_v290_ = nw91_
                d_535_v290_ = d_535_v290_
                d_536_v291_: T4
                nw92_ = C2()
                nw92_.ctor__(d_479_v254_, (d_383_v190_).f15, _dafny.Map({(d_535_v290_).f19: d_381_v188_}), (d_382_v189_) + ((d_383_v190_).f14), d_383_v190_.f12)
                d_536_v291_ = nw92_
                d_536_v291_ = d_536_v291_
        d_537_i27_: int
        d_537_i27_ = 0
        with _dafny.label("2"):
            while (d_383_v190_).f15:
                with _dafny.c_label("2"):
                    if (d_537_i27_) >= (100):
                        raise _dafny.Break("2")
                    d_537_i27_ = (d_537_i27_) + (1)
                    d_538_v292_: _dafny.Array
                    def lambda34_(d_539_v3_):
                        def lambda35_(d_540_i28_):
                            return (_dafny.SeqWithoutIsStrInference([d_539_v3_])) + (_dafny.SeqWithoutIsStrInference([d_539_v3_]))

                        return lambda35_

                    init18_ = lambda34_(d_162_v3_)
                    nw93_ = _dafny.Array(None, 24)
                    for i0_18_ in range(nw93_.length(0)):
                        nw93_[i0_18_] = init18_(i0_18_)
                    d_538_v292_ = nw93_
                    d_541_v293_: _dafny.Seq
                    d_541_v293_ = _dafny.SeqWithoutIsStrInference([-548])
                    index61_ = default__.safeIndex(251, (d_538_v292_).length(0))
                    (d_538_v292_)[index61_] = d_541_v293_
                    index62_ = default__.safeIndex(251, (d_538_v292_).length(0))
                    (d_538_v292_)[index62_] = (_dafny.SeqWithoutIsStrInference([d_162_v3_ for d_542_i29_ in range(default__.abs(80))])) + (d_541_v293_)
                    d_543_v294_: C0
                    nw94_ = C0()
                    nw94_.ctor__()
                    d_543_v294_ = nw94_
                    d_543_v294_ = d_543_v294_
                    d_544_v295_: _dafny.Array
                    def lambda36_(d_545_v141_):
                        def lambda37_(d_546_i30_):
                            return d_545_v141_

                        return lambda37_

                    init19_ = lambda36_(d_320_v141_)
                    nw95_ = _dafny.Array(None, 17)
                    for i0_19_ in range(nw95_.length(0)):
                        nw95_[i0_19_] = init19_(i0_19_)
                    d_544_v295_ = nw95_
                    index63_ = default__.safeIndex(358, (d_544_v295_).length(0))
                    (d_544_v295_)[index63_] = d_320_v141_
                    d_547_v296_: str
                    d_547_v296_ = _dafny.CodePoint('u')
                    index64_ = default__.safeIndex(358, (d_544_v295_).length(0))
                    (d_544_v295_)[index64_] = (((d_320_v141_).set(default__.safeIndex(d_162_v3_, len(d_320_v141_)), d_547_v296_)) + (d_320_v141_)) + (d_320_v141_)
                    d_548_v297_: _dafny.Map
                    d_548_v297_ = _dafny.Map({len(d_320_v141_): (d_383_v190_).f15})
                    d_548_v297_ = (default__.fm49((d_383_v190_).f15, False, not((d_383_v190_).f15), _dafny.CodePoint('n'), d_164_globalState_)) | ((d_548_v297_) | (d_548_v297_))
                    pass
            pass
        if d_160_v1_:
            index65_ = default__.safeIndex(330, (d_159_v0_).length(0))
            (d_159_v0_)[index65_] = 708
            index66_ = default__.safeIndex(330, (d_159_v0_).length(0))
            (d_159_v0_)[index66_] = ((-942) + (d_162_v3_)) + ((d_383_v190_).fm5((d_383_v190_).f15, d_164_globalState_))
            d_549_v298_: _dafny.Array
            def lambda38_(d_550_i31_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saum"))

            init20_ = lambda38_
            nw96_ = _dafny.Array(None, 8)
            for i0_20_ in range(nw96_.length(0)):
                nw96_[i0_20_] = init20_(i0_20_)
            d_549_v298_ = nw96_
            d_551_v299_: _dafny.Seq
            d_551_v299_ = _dafny.SeqWithoutIsStrInference([d_549_v298_, d_549_v298_])
            d_552_v300_: str
            d_552_v300_ = _dafny.CodePoint('n')
            d_553_v301_: _dafny.Map
            d_553_v301_ = _dafny.Map({d_552_v300_: d_549_v298_})
            d_554_v302_: _dafny.Map
            d_554_v302_ = _dafny.Map({len(d_163_v4_): d_549_v298_})
            d_555_v303_: _dafny.Map
            d_555_v303_ = _dafny.Map({(d_159_v0_)[default__.safeIndex(330, (d_159_v0_).length(0))]: ((d_554_v302_)[(d_159_v0_)[default__.safeIndex(330, (d_159_v0_).length(0))]] if ((d_159_v0_)[default__.safeIndex(330, (d_159_v0_).length(0))]) in (d_554_v302_) else d_549_v298_)})
            d_556_v304_: _dafny.Array
            nw97_ = _dafny.Array(None, 22)
            nw97_[int(0)] = d_549_v298_
            nw97_[int(1)] = d_549_v298_
            nw97_[int(2)] = (d_551_v299_)[default__.safeIndex(len(d_320_v141_), len(d_551_v299_))]
            nw97_[int(3)] = d_549_v298_
            nw97_[int(4)] = d_549_v298_
            nw97_[int(5)] = d_549_v298_
            nw97_[int(6)] = d_549_v298_
            nw97_[int(7)] = d_549_v298_
            nw97_[int(8)] = d_549_v298_
            nw97_[int(9)] = d_549_v298_
            nw97_[int(10)] = d_549_v298_
            nw97_[int(11)] = d_549_v298_
            nw97_[int(12)] = d_549_v298_
            nw97_[int(13)] = d_549_v298_
            nw97_[int(14)] = d_549_v298_
            nw97_[int(15)] = ((d_553_v301_)[d_552_v300_] if (d_552_v300_) in (d_553_v301_) else ((d_554_v302_)[d_162_v3_] if (d_162_v3_) in (d_554_v302_) else d_549_v298_))
            nw97_[int(16)] = d_549_v298_
            nw97_[int(17)] = d_549_v298_
            nw97_[int(18)] = d_549_v298_
            nw97_[int(19)] = d_549_v298_
            nw97_[int(20)] = d_549_v298_
            nw97_[int(21)] = d_549_v298_
            d_556_v304_ = nw97_
            index67_ = default__.safeIndex(233, (d_556_v304_).length(0))
            (d_556_v304_)[index67_] = d_549_v298_
            index68_ = default__.safeIndex(233, (d_556_v304_).length(0))
            (d_556_v304_)[index68_] = d_549_v298_
            d_160_v1_ = (len(d_320_v141_)) > ((d_159_v0_)[default__.safeIndex(330, (d_159_v0_).length(0))])
            d_557_v305_: _dafny.Array
            nw98_ = _dafny.Array(False, 24)
            d_557_v305_ = nw98_
            index69_ = default__.safeIndex(378, (d_557_v305_).length(0))
            (d_557_v305_)[index69_] = default__.fm2(d_160_v1_, d_164_globalState_)
            index70_ = default__.safeIndex(378, (d_557_v305_).length(0))
            (d_557_v305_)[index70_] = (d_383_v190_).f15
            d_558_v306_: _dafny.Array
            def lambda39_(d_559_v300_):
                def lambda40_(d_560_i32_):
                    return d_559_v300_

                return lambda40_

            init21_ = lambda39_(d_552_v300_)
            nw99_ = _dafny.Array(None, 21)
            for i0_21_ in range(nw99_.length(0)):
                nw99_[i0_21_] = init21_(i0_21_)
            d_558_v306_ = nw99_
            d_561_v307_: _dafny.Seq
            d_561_v307_ = _dafny.SeqWithoutIsStrInference([d_558_v306_, d_558_v306_, d_558_v306_])
            d_562_v308_: D20
            d_562_v308_ = D20_DC52((d_561_v307_)[default__.safeIndex((0) - (d_162_v3_), len(d_561_v307_))])
            d_563_v309_: _dafny.Map
            d_563_v309_ = _dafny.Map({(d_159_v0_)[default__.safeIndex(330, (d_159_v0_).length(0))]: d_562_v308_})
            d_564_v310_: _dafny.Map
            d_564_v310_ = _dafny.Map({d_563_v309_: d_162_v3_})
            d_565_v312_: D7
            d_565_v312_ = D7_DC20(_dafny.MultiSet([d_162_v3_]))
            d_566_v313_: _dafny.Map
            d_566_v313_ = _dafny.Map({d_565_v312_: d_162_v3_})
            d_567_v314_: _dafny.Map
            def iife40_():
                coll36_ = _dafny.Set()
                compr_36_: int
                for compr_36_ in _dafny.IntegerRange(730, 922):
                    d_568_v311_: int = compr_36_
                    if ((730) <= (d_568_v311_)) and ((d_568_v311_) < (922)):
                        coll36_ = coll36_.union(_dafny.Set([(d_568_v311_) - (d_162_v3_)]))
                return _dafny.Set(coll36_)
            d_567_v314_ = _dafny.Map({len(_dafny.Set({_dafny.Set({150}), iife40_()
            })): d_566_v313_})
            index71_ = default__.safeIndex(330, (d_159_v0_).length(0))
            rhs53_ = len(((((d_567_v314_)[d_162_v3_] if (d_162_v3_) in (d_567_v314_) else _dafny.Map({d_565_v312_: (d_383_v190_).fm8((d_383_v190_).f15, (d_383_v190_).f15, d_164_globalState_)}))) | (d_566_v313_)) | (d_566_v313_))
            rhs54_ = d_564_v310_
            rhs55_ = (d_383_v190_).fm8((d_383_v190_).f15, not(((d_383_v190_).f15 if (d_557_v305_)[default__.safeIndex(378, (d_557_v305_).length(0))] else (d_383_v190_).f15)), d_164_globalState_)
            lhs32_ = d_159_v0_
            lhs33_ = default__.safeIndex(330, (d_159_v0_).length(0))
            lhs32_[lhs33_] = rhs53_
            d_564_v310_ = rhs54_
            d_162_v3_ = rhs55_
        elif True:
            d_569_v315_: bool
            d_570_v316_: int
            d_571_v317_: _dafny.Set
            d_572_v318_: bool
            out50_: bool
            out51_: int
            out52_: _dafny.Set
            out53_: bool
            out50_, out51_, out52_, out53_ = (d_383_v190_).m3(d_164_globalState_)
            d_569_v315_ = out50_
            d_570_v316_ = out51_
            d_571_v317_ = out52_
            d_572_v318_ = out53_
            d_573_v319_: _dafny.Seq
            d_573_v319_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, d_162_v3_])
            d_574_v320_: C2
            nw100_ = C2()
            nw100_.ctor__(d_573_v319_, (d_383_v190_).f15, (d_383_v190_).f13, ((d_383_v190_).f14).set(default__.safeIndex(len(d_320_v141_), len((d_383_v190_).f14)), not((d_383_v190_).f15)), d_383_v190_.f12)
            d_574_v320_ = nw100_
            d_575_v321_: bool
            out54_: bool
            out54_ = (d_574_v320_).m6(d_164_globalState_)
            d_575_v321_ = out54_
            d_576_v322_: bool
            d_577_v323_: int
            d_578_v324_: _dafny.Set
            d_579_v325_: bool
            out55_: bool
            out56_: int
            out57_: _dafny.Set
            out58_: bool
            out55_, out56_, out57_, out58_ = (d_383_v190_).m3(d_164_globalState_)
            d_576_v322_ = out55_
            d_577_v323_ = out56_
            d_578_v324_ = out57_
            d_579_v325_ = out58_
            d_580_v326_: str
            d_580_v326_ = _dafny.CodePoint('d')
            d_581_v327_: _dafny.Array
            nw101_ = _dafny.Array(None, 2)
            nw101_[int(0)] = (d_320_v141_)[default__.safeIndex(d_570_v316_, len(d_320_v141_))]
            nw101_[int(1)] = d_580_v326_
            d_581_v327_ = nw101_
            d_582_v328_: _dafny.Seq
            d_582_v328_ = _dafny.SeqWithoutIsStrInference([d_581_v327_, d_581_v327_])
            d_583_v329_: _dafny.Seq
            d_583_v329_ = d_582_v328_
            d_584_v330_: _dafny.Seq
            d_584_v330_ = _dafny.SeqWithoutIsStrInference([d_583_v329_, d_583_v329_, d_583_v329_, d_583_v329_, d_582_v328_])
            rhs56_ = d_320_v141_
            rhs57_ = d_584_v330_
            rhs58_ = d_320_v141_
            d_320_v141_ = rhs56_
            d_584_v330_ = rhs57_
            d_320_v141_ = rhs58_
        d_585_i33_: int
        d_585_i33_ = 0
        with _dafny.label("3"):
            while d_160_v1_:
                with _dafny.c_label("3"):
                    if (d_585_i33_) >= (100):
                        raise _dafny.Break("3")
                    d_585_i33_ = (d_585_i33_) + (1)
                    (d_164_globalState_).f4 = (d_382_v189_)[default__.safeIndex(d_162_v3_, len(d_382_v189_))]
                    if True:
                        d_586_v331_: _dafny.Map
                        d_586_v331_ = _dafny.Map({(d_383_v190_).f15: d_159_v0_})
                        rhs59_ = len(d_321_v142_)
                        rhs60_ = not (((d_383_v190_).f14)[default__.safeIndex(len(d_586_v331_), len((d_383_v190_).f14))]) or ((d_160_v1_) or ((d_383_v190_).f15))
                        rhs61_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) < (23)
                        rhs62_ = d_160_v1_
                        lhs34_ = d_164_globalState_
                        lhs35_ = d_164_globalState_
                        d_162_v3_ = rhs59_
                        lhs34_.f4 = rhs60_
                        lhs35_.f4 = rhs61_
                        d_160_v1_ = rhs62_
                        d_587_v332_: _dafny.Map
                        d_587_v332_ = _dafny.Map({(d_321_v142_)[default__.safeIndex(66, len(d_321_v142_))]: False})
                        d_588_v333_: _dafny.Set
                        d_588_v333_ = _dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference([d_162_v3_, d_162_v3_, d_162_v3_, d_162_v3_]))), d_162_v3_})
                        d_589_v334_: _dafny.Array
                        nw102_ = _dafny.Array(None, 21)
                        nw102_[int(0)] = d_160_v1_
                        nw102_[int(1)] = (d_383_v190_).f15
                        nw102_[int(2)] = (d_383_v190_).f15
                        nw102_[int(3)] = (d_159_v0_) in (d_587_v332_)
                        nw102_[int(4)] = (d_383_v190_).f15
                        nw102_[int(5)] = (d_383_v190_).f15
                        nw102_[int(6)] = (d_383_v190_).f15
                        nw102_[int(7)] = (d_383_v190_).f15
                        nw102_[int(8)] = d_160_v1_
                        nw102_[int(9)] = (d_383_v190_).f15
                        nw102_[int(10)] = True
                        nw102_[int(11)] = (d_383_v190_).f15
                        nw102_[int(12)] = (not((d_383_v190_).f15)) or ((d_383_v190_).f15)
                        nw102_[int(13)] = ((d_383_v190_).f15) or (not((d_383_v190_).f15))
                        nw102_[int(14)] = (d_588_v333_).ispropersubset(d_588_v333_)
                        nw102_[int(15)] = True
                        nw102_[int(16)] = d_160_v1_
                        nw102_[int(17)] = not(d_160_v1_)
                        nw102_[int(18)] = (d_320_v141_) < (d_320_v141_)
                        nw102_[int(19)] = d_160_v1_
                        nw102_[int(20)] = not((d_383_v190_).f15)
                        d_589_v334_ = nw102_
                        index72_ = default__.safeIndex(80, (d_589_v334_).length(0))
                        (d_589_v334_)[index72_] = d_160_v1_
                        index73_ = default__.safeIndex(80, (d_589_v334_).length(0))
                        (d_589_v334_)[index73_] = (d_383_v190_).f15
                        d_590_v335_: D12
                        d_590_v335_ = D12_DC37((d_320_v141_) + (d_320_v141_), d_160_v1_)
                        d_590_v335_ = d_590_v335_
                        d_591_v336_: _dafny.Map
                        d_591_v336_ = _dafny.Map({(d_383_v190_).f14: (d_163_v4_) | (_dafny.Map({d_162_v3_: d_162_v3_}))})
                        d_591_v336_ = d_591_v336_
                        d_592_v337_: str
                        d_592_v337_ = _dafny.CodePoint('o')
                        d_592_v337_ = (d_320_v141_)[default__.safeIndex(default__.safeDivisionInt(140, d_162_v3_), len(d_320_v141_))]
                    elif True:
                        d_593_v338_: _dafny.MultiSet
                        d_593_v338_ = _dafny.MultiSet([d_162_v3_])
                        d_594_v339_: _dafny.Map
                        d_594_v339_ = _dafny.Map({d_593_v338_: d_162_v3_})
                        d_594_v339_ = d_594_v339_
                        d_160_v1_ = ((d_383_v190_).f15) not in ((d_206_v40_) - (d_206_v40_))
                        d_595_v340_: C1
                        nw103_ = C1()
                        nw103_.ctor__((d_383_v190_).f15, d_160_v1_, default__.fm60(d_162_v3_, d_164_globalState_), (d_383_v190_).f14, d_380_v187_)
                        d_595_v340_ = nw103_
                        d_596_v341_: D11
                        d_596_v341_ = D11_DC33((d_383_v190_).f15, 766, (d_383_v190_).f15, d_595_v340_, (d_383_v190_).f15)
                        d_596_v341_ = d_596_v341_
                        d_597_v342_: bool
                        d_598_v343_: int
                        d_599_v344_: _dafny.Set
                        d_600_v345_: bool
                        out59_: bool
                        out60_: int
                        out61_: _dafny.Set
                        out62_: bool
                        out59_, out60_, out61_, out62_ = (d_383_v190_).m3(d_164_globalState_)
                        d_597_v342_ = out59_
                        d_598_v343_ = out60_
                        d_599_v344_ = out61_
                        d_600_v345_ = out62_
                        d_601_v346_: _dafny.Seq
                        d_601_v346_ = _dafny.SeqWithoutIsStrInference([d_598_v343_, d_162_v3_, d_162_v3_])
                        d_602_v347_: _dafny.Set
                        d_602_v347_ = _dafny.Set({len(d_381_v188_), d_598_v343_, d_598_v343_})
                        d_603_v349_: C6
                        nw104_ = C6()
                        def iife41_():
                            coll37_ = _dafny.Set()
                            compr_37_: int
                            for compr_37_ in (d_602_v347_).Elements:
                                d_604_v348_: int = compr_37_
                                if (d_604_v348_) in (d_602_v347_):
                                    coll37_ = coll37_.union(_dafny.Set([default__.safeDivisionInt(d_604_v348_, 22)]))
                            return _dafny.Set(coll37_)
                        nw104_.ctor__(d_601_v346_, (default__.fm70((d_383_v190_).f15, d_320_v141_, d_162_v3_, (d_383_v190_).f15, d_164_globalState_)) != (iife41_()
                        ), ((d_383_v190_).f13) | ((d_383_v190_).f13), _dafny.SeqWithoutIsStrInference([(d_383_v190_).f15, (d_383_v190_).f15]), d_380_v187_)
                        d_603_v349_ = nw104_
                    d_162_v3_ = d_162_v3_
                    d_605_v350_: _dafny.Array
                    nw105_ = _dafny.Array(False, 21)
                    d_605_v350_ = nw105_
                    d_606_v351_: _dafny.Seq
                    d_606_v351_ = _dafny.SeqWithoutIsStrInference([d_605_v350_])
                    d_607_v352_: _dafny.Map
                    d_607_v352_ = _dafny.Map({(d_606_v351_)[default__.safeIndex(966, len(d_606_v351_))]: (d_383_v190_).f15})
                    d_607_v352_ = (d_607_v352_).set(d_605_v350_, (d_383_v190_).f15)
                    pass
            pass
        d_608_v353_: str
        d_608_v353_ = _dafny.CodePoint('b')
        d_609_v354_: _dafny.Map
        d_609_v354_ = _dafny.Map({d_608_v353_: (d_383_v190_).f15})
        d_610_v355_: _dafny.MultiSet
        d_610_v355_ = _dafny.MultiSet([d_609_v354_])
        d_611_i34_: int
        d_611_i34_ = 0
        with _dafny.label("4"):
            while default__.fm2((d_610_v355_).isdisjoint(d_610_v355_), d_164_globalState_):
                with _dafny.c_label("4"):
                    if (d_611_i34_) >= (100):
                        raise _dafny.Break("4")
                    d_611_i34_ = (d_611_i34_) + (1)
                    d_612_v356_: D4
                    d_612_v356_ = D4_DC12(_dafny.SeqWithoutIsStrInference([True]))
                    source7_ = d_612_v356_
                    if source7_.is_DC13:
                        d_613___mcc_h19_ = source7_.cf19
                        d_614___mcc_h20_ = source7_.cf20
                        d_615___mcc_h21_ = source7_.cf21
                        d_616___mcc_h22_ = source7_.cf22
                        d_617___mcc_h23_ = source7_.cf23
                        d_618_cf23_ = d_617___mcc_h23_
                        d_619_cf22_ = d_616___mcc_h22_
                        d_620_cf21_ = d_615___mcc_h21_
                        d_621_cf20_ = d_614___mcc_h20_
                        d_622_cf19_ = d_613___mcc_h19_
                        d_623_v357_: D0
                        d_623_v357_ = D0_DC2(d_619_cf22_, (d_383_v190_).f15)
                        d_624_v358_: _dafny.Set
                        d_624_v358_ = _dafny.Set({d_320_v141_, d_320_v141_})
                        d_625_v359_: _dafny.Seq
                        d_626_v360_: bool
                        d_627_v361_: int
                        d_628_v362_: bool
                        out63_: _dafny.Seq
                        out64_: bool
                        out65_: int
                        out66_: bool
                        out63_, out64_, out65_, out66_ = default__.m0(d_621_cf20_, d_623_v357_, (d_624_v358_).isdisjoint(d_624_v358_), d_164_globalState_)
                        d_625_v359_ = out63_
                        d_626_v360_ = out64_
                        d_627_v361_ = out65_
                        d_628_v362_ = out66_
                        d_319_v140_ = (d_319_v140_).set(default__.fm0(not(d_622_cf19_), d_319_v140_, d_626_v360_, d_164_globalState_), d_318_v139_)
                        d_619_cf22_ = default__.safeModuloInt(d_620_cf21_, d_620_cf21_)
                        d_629_v363_: _dafny.Map
                        d_629_v363_ = _dafny.Map({d_383_v190_: d_320_v141_})
                        d_320_v141_ = ((d_629_v363_)[d_383_v190_] if (d_383_v190_) in (d_629_v363_) else d_320_v141_)
                    elif source7_.is_DC14:
                        d_630___mcc_h24_ = source7_.cf24
                        d_631_cf24_ = d_630___mcc_h24_
                        d_320_v141_ = (d_320_v141_) + (_dafny.SeqWithoutIsStrInference([d_608_v353_ for d_632_i35_ in range(default__.abs(594))]))
                        d_633_v364_: bool
                        d_634_v365_: int
                        d_635_v366_: _dafny.Set
                        d_636_v367_: bool
                        out67_: bool
                        out68_: int
                        out69_: _dafny.Set
                        out70_: bool
                        out67_, out68_, out69_, out70_ = (d_383_v190_).m3(d_164_globalState_)
                        d_633_v364_ = out67_
                        d_634_v365_ = out68_
                        d_635_v366_ = out69_
                        d_636_v367_ = out70_
                        d_637_v368_: C2
                        nw106_ = C2()
                        nw106_.ctor__((d_383_v190_).fm6(d_320_v141_, d_164_globalState_), d_636_v367_, (_dafny.Map({True: d_381_v188_})) | ((d_383_v190_).f13), (d_383_v190_).f14, d_383_v190_.f12)
                        d_637_v368_ = nw106_
                        d_637_v368_ = d_637_v368_
                        d_638_v369_: D0
                        d_638_v369_ = D0_DC2(d_162_v3_, (d_383_v190_).f15)
                        d_639_v370_: _dafny.Seq
                        d_640_v371_: bool
                        d_641_v372_: int
                        d_642_v373_: bool
                        out71_: _dafny.Seq
                        out72_: bool
                        out73_: int
                        out74_: bool
                        out71_, out72_, out73_, out74_ = default__.m0((d_383_v190_).f15, d_638_v369_, (d_383_v190_).f15, d_164_globalState_)
                        d_639_v370_ = out71_
                        d_640_v371_ = out72_
                        d_641_v372_ = out73_
                        d_642_v373_ = out74_
                    elif True:
                        d_643___mcc_h25_ = source7_.cf18
                        d_644_cf18_ = d_643___mcc_h25_
                        d_645_v374_: C3
                        nw107_ = C3()
                        nw107_.ctor__(d_160_v1_, d_383_v190_.f12, (d_383_v190_).f15, (d_383_v190_).f13, (d_383_v190_).f14)
                        d_645_v374_ = nw107_
                        d_646_v375_: _dafny.Seq
                        d_646_v375_ = _dafny.SeqWithoutIsStrInference([d_645_v374_])
                        d_647_v376_: _dafny.Map
                        d_647_v376_ = _dafny.Map({d_646_v375_: d_320_v141_})
                        d_160_v1_ = (((d_647_v376_)[d_646_v375_] if (d_646_v375_) in (d_647_v376_) else d_320_v141_)) == (d_320_v141_)
                        d_648_v377_: _dafny.Seq
                        d_648_v377_ = _dafny.SeqWithoutIsStrInference([d_162_v3_])
                        (d_164_globalState_).f4 = not((d_648_v377_) == (d_648_v377_))
                        (d_164_globalState_).f4 = (d_645_v374_).f20
                        d_649_v378_: C5
                        nw108_ = C5()
                        nw108_.ctor__((d_383_v190_).f15, d_380_v187_, (d_383_v190_).f13, (d_382_v189_) + (d_644_cf18_))
                        d_649_v378_ = nw108_
                    d_650_v379_: D5
                    d_650_v379_ = D5_DC16(d_160_v1_)
                    d_651_v380_: _dafny.Set
                    d_651_v380_ = _dafny.Set({d_650_v379_})
                    d_652_v381_: _dafny.Seq
                    d_652_v381_ = _dafny.SeqWithoutIsStrInference([d_650_v379_])
                    d_653_v383_: _dafny.Seq
                    def iife42_():
                        coll38_ = _dafny.Set()
                        compr_38_: D5
                        for compr_38_ in (d_652_v381_).Elements:
                            d_654_v382_: D5 = compr_38_
                            if (d_654_v382_) in (d_652_v381_):
                                coll38_ = coll38_.union(_dafny.Set([d_654_v382_]))
                        return _dafny.Set(coll38_)
                    d_653_v383_ = _dafny.SeqWithoutIsStrInference([d_651_v380_, (d_651_v380_).intersection(_dafny.Set({d_650_v379_})), _dafny.Set({d_650_v379_, d_650_v379_}), (d_651_v380_) - (iife42_()
                    )])
                    d_653_v383_ = d_653_v383_
                    if not(True):
                        d_655_v384_: _dafny.Map
                        d_655_v384_ = _dafny.Map({d_159_v0_: (d_383_v190_).f15})
                        d_656_v385_: C1
                        nw109_ = C1()
                        nw109_.ctor__((d_383_v190_).f15, True, (d_383_v190_).f13, ((d_383_v190_).f14 if d_160_v1_ else d_382_v189_), (d_383_v190_.f12 if ((d_655_v384_)[d_159_v0_] if (d_159_v0_) in (d_655_v384_) else (d_383_v190_).f15) else d_383_v190_.f12))
                        d_656_v385_ = nw109_
                        d_657_v386_: _dafny.Map
                        d_657_v386_ = _dafny.Map({d_159_v0_: d_162_v3_})
                        d_658_v387_: _dafny.Map
                        d_658_v387_ = _dafny.Map({(d_656_v385_).fm5(d_160_v1_, d_164_globalState_): (d_159_v0_) in (d_657_v386_)})
                        d_658_v387_ = (d_658_v387_).set(d_162_v3_, (d_656_v385_).f17)
                        d_162_v3_ = (d_162_v3_) + ((438) * (d_162_v3_))
                        d_659_v388_: C6
                        nw110_ = C6()
                        nw110_.ctor__(_dafny.SeqWithoutIsStrInference([(D20_DC53(d_318_v139_, d_162_v3_, d_162_v3_)).cf79 for d_660_i36_ in range(default__.abs(353))]), (d_383_v190_).f15, (d_383_v190_).f13, (d_383_v190_).f14, d_383_v190_.f12)
                        d_659_v388_ = nw110_
                        d_661_v389_: _dafny.Map
                        d_661_v389_ = _dafny.Map({d_659_v388_: (d_383_v190_).f15})
                        d_662_v390_: _dafny.Seq
                        d_662_v390_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, len(d_661_v389_)])
                        d_663_v391_: _dafny.Set
                        d_663_v391_ = _dafny.Set({d_162_v3_, default__.safeModuloInt(d_162_v3_, len(d_662_v390_))})
                        d_663_v391_ = d_663_v391_
                        d_162_v3_ = (0) - ((default__.safeDivisionInt((d_161_v2_).cardinality, d_162_v3_)) - (d_162_v3_))
                    elif True:
                        d_664_v392_: _dafny.Array
                        nw111_ = _dafny.Array(_dafny.Map({}), 22)
                        d_664_v392_ = nw111_
                        index74_ = default__.safeIndex(782, (d_664_v392_).length(0))
                        (d_664_v392_)[index74_] = ((d_381_v188_).set((d_383_v190_).f15, (d_383_v190_).f15)).set((d_383_v190_).f15, default__.fm2((d_383_v190_).f15, d_164_globalState_))
                        d_665_v393_: _dafny.Array
                        nw112_ = _dafny.Array(None, 14)
                        d_665_v393_ = nw112_
                        d_666_v394_: T0
                        nw113_ = C5()
                        nw113_.ctor__((d_383_v190_).f15, d_383_v190_.f12, (d_383_v190_).f13, (d_383_v190_).f14)
                        d_666_v394_ = nw113_
                        index75_ = default__.safeIndex(862, (d_665_v393_).length(0))
                        (d_665_v393_)[index75_] = d_666_v394_
                        index76_ = default__.safeIndex(782, (d_664_v392_).length(0))
                        index77_ = default__.safeIndex(862, (d_665_v393_).length(0))
                        rhs63_ = d_160_v1_
                        rhs64_ = d_381_v188_
                        rhs65_ = d_162_v3_
                        rhs66_ = d_666_v394_
                        lhs36_ = d_164_globalState_
                        lhs37_ = d_664_v392_
                        lhs38_ = default__.safeIndex(782, (d_664_v392_).length(0))
                        lhs39_ = d_665_v393_
                        lhs40_ = default__.safeIndex(862, (d_665_v393_).length(0))
                        lhs36_.f4 = rhs63_
                        lhs37_[lhs38_] = rhs64_
                        d_162_v3_ = rhs65_
                        lhs39_[lhs40_] = rhs66_
                        d_162_v3_ = d_162_v3_
                        d_667_v395_: _dafny.MultiSet
                        d_667_v395_ = _dafny.MultiSet([d_162_v3_])
                        d_668_v396_: D0
                        d_668_v396_ = D0_DC2(d_162_v3_, d_160_v1_)
                        d_669_v397_: _dafny.Seq
                        d_670_v398_: bool
                        d_671_v399_: int
                        d_672_v400_: bool
                        out75_: _dafny.Seq
                        out76_: bool
                        out77_: int
                        out78_: bool
                        out75_, out76_, out77_, out78_ = default__.m0((d_667_v395_) == (d_667_v395_), d_668_v396_, (d_383_v190_).f15, d_164_globalState_)
                        d_669_v397_ = out75_
                        d_670_v398_ = out76_
                        d_671_v399_ = out77_
                        d_672_v400_ = out78_
                        d_672_v400_ = False
                        index78_ = default__.safeIndex(84, (d_159_v0_).length(0))
                        (d_159_v0_)[index78_] = d_671_v399_
                        index79_ = default__.safeIndex(84, (d_159_v0_).length(0))
                        (d_159_v0_)[index79_] = (d_383_v190_).fm7(d_671_v399_, d_164_globalState_)
                    d_673_v401_: D8
                    d_673_v401_ = D8_DC26((d_383_v190_).f15)
                    d_674_v402_: _dafny.Array
                    nw114_ = _dafny.Array(None, 6)
                    nw114_[int(0)] = d_673_v401_
                    nw114_[int(1)] = d_673_v401_
                    nw114_[int(2)] = D8_DC26(d_160_v1_)
                    nw114_[int(3)] = d_673_v401_
                    nw114_[int(4)] = d_673_v401_
                    nw114_[int(5)] = d_673_v401_
                    d_674_v402_ = nw114_
                    index80_ = default__.safeIndex(116, (d_674_v402_).length(0))
                    (d_674_v402_)[index80_] = D8_DC26((d_383_v190_).f15)
                    pat_let_tv1_ = d_383_v190_
                    pat_let_tv2_ = d_160_v1_
                    index81_ = default__.safeIndex(116, (d_674_v402_).length(0))
                    def iife43_(_pat_let2_0):
                        def iife44_(d_675_dt__update__tmp_h2_):
                            def iife45_(_pat_let3_0):
                                def iife46_(d_676_dt__update_hcf41_h0_):
                                    return D8_DC26(d_676_dt__update_hcf41_h0_)
                                return iife46_(_pat_let3_0)
                            return iife45_((not((pat_let_tv1_).f15)) and (pat_let_tv2_))
                        return iife44_(_pat_let2_0)
                    (d_674_v402_)[index81_] = iife43_(d_673_v401_)
                    pass
            pass
        d_677_v403_: _dafny.Array
        nw115_ = _dafny.Array(False, 6)
        d_677_v403_ = nw115_
        d_678_v404_: _dafny.MultiSet
        d_678_v404_ = _dafny.MultiSet([d_677_v403_, d_677_v403_])
        index82_ = default__.safeIndex(462, (d_677_v403_).length(0))
        (d_677_v403_)[index82_] = (d_678_v404_).ispropersubset(_dafny.MultiSet([d_677_v403_]))
        index83_ = default__.safeIndex(462, (d_677_v403_).length(0))
        (d_677_v403_)[index83_] = (d_162_v3_) == (len(((d_320_v141_) + (d_320_v141_)).set(default__.safeIndex(d_162_v3_, len((d_320_v141_) + (d_320_v141_))), d_608_v353_)))
        d_679_v405_: T4
        nw116_ = C2()
        nw116_.ctor__(_dafny.SeqWithoutIsStrInference([len(d_320_v141_) for d_680_i37_ in range(default__.abs(342))]), ((d_383_v190_).f15) == ((d_383_v190_).f15), (d_383_v190_).f13, (d_383_v190_).f14, d_383_v190_.f12)
        d_679_v405_ = nw116_
        d_679_v405_ = d_679_v405_
        d_681_v406_: C3
        nw117_ = C3()
        nw117_.ctor__((d_383_v190_).f15, d_383_v190_.f12, False, _dafny.Map({False: d_381_v188_}), d_382_v189_)
        d_681_v406_ = nw117_
        d_682_v407_: _dafny.MultiSet
        d_682_v407_ = _dafny.MultiSet([d_681_v406_])
        d_683_v408_: _dafny.Seq
        d_683_v408_ = _dafny.SeqWithoutIsStrInference([d_162_v3_, ((d_682_v407_)[d_681_v406_] if (d_681_v406_) in (d_682_v407_) else d_162_v3_), d_162_v3_])
        d_684_v409_: _dafny.Map
        d_684_v409_ = _dafny.Map({d_162_v3_: d_683_v408_})
        d_685_v410_: D5
        d_685_v410_ = D5_DC15(d_683_v408_)
        d_686_v411_: _dafny.Map
        d_686_v411_ = _dafny.Map({d_160_v1_: (d_685_v410_).cf25})
        d_687_v412_: _dafny.Array
        nw118_ = _dafny.Array(None, 12)
        nw118_[int(0)] = (d_683_v408_) + (d_683_v408_)
        nw118_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_683_v408_)[default__.safeIndex(d_162_v3_, len(d_683_v408_))], d_162_v3_])
        nw118_[int(2)] = (d_683_v408_) + (d_683_v408_)
        nw118_[int(3)] = d_683_v408_
        nw118_[int(4)] = (d_683_v408_) + (_dafny.SeqWithoutIsStrInference([d_162_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkdkja")))]))
        nw118_[int(5)] = d_683_v408_
        nw118_[int(6)] = d_683_v408_
        nw118_[int(7)] = (d_683_v408_) + ((D5_DC15(((d_684_v409_)[d_162_v3_] if (d_162_v3_) in (d_684_v409_) else ((d_686_v411_)[d_160_v1_] if (d_160_v1_) in (d_686_v411_) else _dafny.SeqWithoutIsStrInference([d_162_v3_]))))).cf25)
        nw118_[int(8)] = (d_681_v406_).fm6(d_320_v141_, d_164_globalState_)
        nw118_[int(9)] = d_683_v408_
        nw118_[int(10)] = d_683_v408_
        nw118_[int(11)] = d_683_v408_
        d_687_v412_ = nw118_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_687_v412_).length(0)):
            d_688_i38_: int = guard_loop_0_
            if (True) and (((0) <= (d_688_i38_)) and ((d_688_i38_) < ((d_687_v412_).length(0)))):
                (d_687_v412_)[(d_688_i38_)] = ((d_683_v408_) + (d_683_v408_)) + (d_683_v408_)
        _dafny.print(_dafny.string_of((d_159_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v2_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_162_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v4_) == (_dafny.Map({846: 846}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_164_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_164_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_164_globalState_).f5) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_164_globalState_).f10) == (_dafny.Map({846: 846}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v40_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_298_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_318_v139_) == (_dafny.Map({True: 3}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_319_v140_) == (_dafny.Map({3: _dafny.Map({True: 3})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_320_v141_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_321_v142_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_322_v143_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_381_v188_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_382_v189_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_383_v190_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_383_v190_).f13) == (_dafny.Map({False: _dafny.Map({False: False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_383_v190_).f14) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_384_v191_).cf63).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_384_v191_).cf63).f13) == (_dafny.Map({False: _dafny.Map({False: False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_384_v191_).cf63).f14) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_537_i27_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_585_i33_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_608_v353_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_609_v354_) == (_dafny.Map({_dafny.CodePoint('b'): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_610_v355_) == (_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('b'): False})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_611_i34_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_677_v403_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_678_v404_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_681_v406_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_682_v407_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_683_v408_) == (_dafny.SeqWithoutIsStrInference([3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_684_v409_) == (_dafny.Map({3: _dafny.SeqWithoutIsStrInference([3, 1, 3])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_685_v410_).cf25) == (_dafny.SeqWithoutIsStrInference([3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_686_v411_) == (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([3, 1, 3])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[0]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[1]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[2]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[3]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[4]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[5]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[6]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[7]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[8]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[9]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[10]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_687_v412_)[11]) == (_dafny.SeqWithoutIsStrInference([3, 1, 3, 3, 1, 3, 3, 1, 3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
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

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
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
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4
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
        return lambda: D3_DC9(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC9(D3, NamedTuple('DC9', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(False, False, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC18(D6, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(_dafny.Map({}), False, False)
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

class D7_DC21(D7, NamedTuple('DC21', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D8_DC27)

class D8_DC24(D8, NamedTuple('DC24', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC27(D8, NamedTuple('DC27', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC27({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC27) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)

class D9_DC28(D9, NamedTuple('DC28', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC29(D10, NamedTuple('DC29', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D11_DC33)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC32(D11, NamedTuple('DC32', [('cf46', Any), ('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC33(D11, NamedTuple('DC33', [('cf48', Any), ('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC33({_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC33) and self.cf48 == __o.cf48 and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC35(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D12_DC37)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D12_DC38)

class D12_DC35(D12, NamedTuple('DC35', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC37(D12, NamedTuple('DC37', [('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC37({self.cf59.VerbatimString(True)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC37) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC38(D12, NamedTuple('DC38', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC38({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC38) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)

class D13_DC39(D13, NamedTuple('DC39', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC41(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D14_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D14_DC43)

class D14_DC41(D14, NamedTuple('DC41', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({self.cf64.VerbatimString(True)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC42(D14, NamedTuple('DC42', [])):
    def __dafnystr__(self) -> str:
        return f'D14.DC42'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC42)
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC43(D14, NamedTuple('DC43', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC43({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC43) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)

class D15_DC44(D15, NamedTuple('DC44', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)

class D16_DC45(D16, NamedTuple('DC45', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC47(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D17_DC49)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC47(D17, NamedTuple('DC47', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC49(D17, NamedTuple('DC49', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC49({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC49) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)

class D18_DC50(D18, NamedTuple('DC50', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)

class D19_DC51(D19, NamedTuple('DC51', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC53(_dafny.Map({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)

class D20_DC53(D20, NamedTuple('DC53', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC54(D20, NamedTuple('DC54', [('cf80', Any), ('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC56(None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D21_DC58)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC56(D21, NamedTuple('DC56', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC57(D21, NamedTuple('DC57', [('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC58(D21, NamedTuple('DC58', [('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC58({_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC58) and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC60(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D22_DC62)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)

class D22_DC60(D22, NamedTuple('DC60', [('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC62(D22, NamedTuple('DC62', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC62({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC62) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC64()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D23_DC63)

class D23_DC64(D23, NamedTuple('DC64', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC63(D23, NamedTuple('DC63', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC63({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC63) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC66(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D24_DC66)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D24_DC67)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D24_DC65)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D24_DC68)

class D24_DC66(D24, NamedTuple('DC66', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC66({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC66) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC67(D24, NamedTuple('DC67', [('cf101', Any), ('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC67({_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC67) and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC65(D24, NamedTuple('DC65', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC65({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC65) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC68(D24, NamedTuple('DC68', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC68({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC68) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D25_DC69)

class D25_DC69(D25, NamedTuple('DC69', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC69({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC69) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC71(int(0), False, _dafny.Array(None, 0), _dafny.Set({}), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D26_DC70)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D26_DC72)

class D26_DC71(D26, NamedTuple('DC71', [('cf107', Any), ('cf108', Any), ('cf109', Any), ('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)}, {_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf107 == __o.cf107 and self.cf108 == __o.cf108 and self.cf109 == __o.cf109 and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC70(D26, NamedTuple('DC70', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC70({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC70) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC72(D26, NamedTuple('DC72', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC72({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC72) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC74(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D27_DC74)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D27_DC75)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D27_DC73)

class D27_DC74(D27, NamedTuple('DC74', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC74({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC74) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC75(D27, NamedTuple('DC75', [('cf115', Any), ('cf116', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC75({self.cf115.VerbatimString(True)}, {_dafny.string_of(self.cf116)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC75) and self.cf115 == __o.cf115 and self.cf116 == __o.cf116
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC73(D27, NamedTuple('DC73', [('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC73({_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC73) and self.cf113 == __o.cf113
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
    def m3(self, globalState):
        pass


class T2:
    pass

class T3:
    pass
    def m4(self, p0, p1, p2, p3, globalState):
        pass


class T4:
    pass
    def m5(self, p0, p1, p2, globalState):
        pass

    def m6(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f4: bool = False
        self._f0: bool = False
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: int = int(0)
        self._f3: int = int(0)
        self._f5: _dafny.MultiSet = _dafny.MultiSet({})
        self._f6: int = int(0)
        self._f7: int = int(0)
        self._f8: bool = False
        self._f9: bool = False
        self._f10: _dafny.Map = _dafny.Map({})
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11

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
    def f5(self):
        return self._f5
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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm14(self, p0, globalState):
        return len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fihm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajtstts")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_689_i0_ in range(default__.abs(166))])))

    def m7(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_690_v0_: _dafny.Array
        def lambda41_(d_691_p2_, d_692_p0_):
            def lambda42_(d_693_i0_):
                return (d_691_p2_) <= (d_692_p0_)

            return lambda42_

        init22_ = lambda41_(p2, p0)
        nw119_ = _dafny.Array(None, 3)
        for i0_22_ in range(nw119_.length(0)):
            nw119_[i0_22_] = init22_(i0_22_)
        d_690_v0_ = nw119_
        d_694_v1_: bool
        d_694_v1_ = True
        index84_ = default__.safeIndex(306, (d_690_v0_).length(0))
        (d_690_v0_)[index84_] = d_694_v1_
        index85_ = default__.safeIndex(306, (d_690_v0_).length(0))
        (d_690_v0_)[index85_] = d_694_v1_
        d_695_v2_: _dafny.Map
        d_695_v2_ = _dafny.Map({d_694_v1_: False})
        d_696_v3_: _dafny.Seq
        d_696_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juqts"))
        d_697_v4_: _dafny.Map
        d_697_v4_ = _dafny.Map({d_696_v3_: p3})
        d_698_v5_: _dafny.Set
        d_698_v5_ = _dafny.Set({(0) - (p0)})
        d_695_v2_ = (d_695_v2_).set(d_694_v1_, (d_698_v5_).issubset(_dafny.Set({len(d_697_v4_), (0) - (len(d_698_v5_))})))
        d_699_v6_: D0
        d_699_v6_ = D0_DC2(((_dafny.MultiSet([p2, (self).fm14(p0, globalState)])).cardinality) + (len(d_695_v2_)), (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))])
        source8_ = d_699_v6_
        if source8_.is_DC1:
            (globalState).f4 = d_694_v1_
            r1 = default__.safeDivisionInt(p2, p2)
            d_694_v1_ = not(not (((0) - (p0)) >= (p0)) or (not (d_694_v1_) or (d_694_v1_)))
            d_696_v3_ = (d_696_v3_ if d_694_v1_ else (d_696_v3_).set(default__.safeIndex(p0, len(d_696_v3_)), p3))
        elif source8_.is_DC2:
            d_700___mcc_h0_ = source8_.cf1
            d_701___mcc_h1_ = source8_.cf2
            d_702_cf2_ = d_701___mcc_h1_
            d_703_cf1_ = d_700___mcc_h0_
            d_704_v7_: _dafny.Array
            nw120_ = _dafny.Array(_dafny.Seq({}), 4)
            d_704_v7_ = nw120_
            d_705_v8_: _dafny.Seq
            d_705_v8_ = _dafny.SeqWithoutIsStrInference([d_694_v1_, d_702_cf2_])
            d_706_v9_: _dafny.Seq
            d_706_v9_ = _dafny.SeqWithoutIsStrInference([d_705_v8_, d_705_v8_])
            index86_ = default__.safeIndex(306, (d_704_v7_).length(0))
            (d_704_v7_)[index86_] = d_706_v9_
            d_707_v10_: D2
            d_707_v10_ = D2_DC4(d_706_v9_)
            d_708_v11_: _dafny.MultiSet
            d_708_v11_ = _dafny.MultiSet([d_694_v1_])
            d_709_v12_: _dafny.Seq
            d_709_v12_ = _dafny.SeqWithoutIsStrInference([(d_708_v11_).cardinality])
            index87_ = default__.safeIndex(306, (d_704_v7_).length(0))
            (d_704_v7_)[index87_] = (((d_707_v10_).cf4).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([p0, 585])) + (d_709_v12_)), len((d_707_v10_).cf4)), default__.fm15((d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))], p2, p0, p2, globalState))).set(default__.safeIndex(len(d_696_v3_), len(((d_707_v10_).cf4).set(default__.safeIndex(len((_dafny.SeqWithoutIsStrInference([p0, 585])) + (d_709_v12_)), len((d_707_v10_).cf4)), default__.fm15((d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))], p2, p0, p2, globalState)))), ((d_705_v8_) + (_dafny.SeqWithoutIsStrInference([True]))).set(default__.safeIndex(p2, len((d_705_v8_) + (_dafny.SeqWithoutIsStrInference([True])))), d_694_v1_))
            (globalState).f4 = not(d_702_cf2_)
            d_710_v13_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_710_v13_ = nw121_
            index88_ = default__.safeIndex(968, (d_710_v13_).length(0))
            (d_710_v13_)[index88_] = d_690_v0_
            index89_ = default__.safeIndex(968, (d_710_v13_).length(0))
            (d_710_v13_)[index89_] = d_690_v0_
            d_711_v14_: _dafny.MultiSet
            d_711_v14_ = (d_708_v11_).set(False, default__.abs((d_708_v11_).cardinality))
            d_712_v15_: _dafny.Array
            nw122_ = _dafny.Array(None, 16)
            nw122_[int(0)] = d_711_v14_
            nw122_[int(1)] = d_711_v14_
            nw122_[int(2)] = d_708_v11_
            nw122_[int(3)] = d_711_v14_
            nw122_[int(4)] = d_711_v14_
            nw122_[int(5)] = _dafny.MultiSet([d_702_cf2_])
            nw122_[int(6)] = d_711_v14_
            nw122_[int(7)] = d_711_v14_
            nw122_[int(8)] = d_711_v14_
            nw122_[int(9)] = (d_708_v11_ if False else d_711_v14_)
            nw122_[int(10)] = d_711_v14_
            nw122_[int(11)] = d_711_v14_
            nw122_[int(12)] = d_711_v14_
            nw122_[int(13)] = d_711_v14_
            nw122_[int(14)] = d_711_v14_
            nw122_[int(15)] = d_711_v14_
            d_712_v15_ = nw122_
            d_712_v15_ = d_712_v15_
        elif True:
            d_713___mcc_h2_ = source8_.cf0
            d_714_cf0_ = d_713___mcc_h2_
            index90_ = default__.safeIndex(306, (d_690_v0_).length(0))
            (d_690_v0_)[index90_] = (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))]
            d_715_v16_: _dafny.Map
            d_715_v16_ = _dafny.Map({p2: d_696_v3_})
            d_716_v17_: _dafny.Map
            d_716_v17_ = _dafny.Map({len(d_715_v16_): (d_698_v5_) - (d_698_v5_)})
            d_717_v18_: _dafny.MultiSet
            d_717_v18_ = _dafny.MultiSet([d_694_v1_])
            def iife47_():
                coll39_ = _dafny.Set()
                compr_39_: int
                for compr_39_ in _dafny.IntegerRange(667, 946):
                    d_718_v19_: int = compr_39_
                    if ((667) <= (d_718_v19_)) and ((d_718_v19_) < (946)):
                        coll39_ = coll39_.union(_dafny.Set([default__.safeModuloInt(d_718_v19_, 199)]))
                return _dafny.Set(coll39_)
            d_716_v17_ = (d_716_v17_).set(((d_717_v18_) - (_dafny.MultiSet([not(True)]))).cardinality, (d_698_v5_).intersection(iife47_()
            ))
            r1 = default__.safeModuloInt(len(d_696_v3_), p2)
            r1 = (0) - (p0)
        r1 = p0
        d_719_i1_: int
        d_719_i1_ = 0
        with _dafny.label("5"):
            pat_let_tv3_ = d_694_v1_
            pat_let_tv4_ = d_690_v0_
            pat_let_tv5_ = d_690_v0_
            pat_let_tv6_ = d_694_v1_
            def lambda43_(source9_):
                if source9_.is_DC1:
                    return (D0_DC0(pat_let_tv3_)).cf0
                elif source9_.is_DC2:
                    d_733___mcc_h3_ = source9_.cf1
                    d_734___mcc_h4_ = source9_.cf2
                    d_735_cf2_ = d_734___mcc_h4_
                    d_736_cf1_ = d_733___mcc_h3_
                    return (pat_let_tv5_)[default__.safeIndex(306, (pat_let_tv4_).length(0))]
                elif True:
                    d_737___mcc_h5_ = source9_.cf0
                    d_738_cf0_ = d_737___mcc_h5_
                    return pat_let_tv6_

            while lambda43_(d_699_v6_):
                with _dafny.c_label("5"):
                    if (d_719_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_719_i1_ = (d_719_i1_) + (1)
                    (globalState).f4 = d_694_v1_
                    d_720_v20_: D0
                    d_720_v20_ = D0_DC0((d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))])
                    d_721_v21_: _dafny.Map
                    d_721_v21_ = _dafny.Map({(d_720_v20_).cf0: d_690_v0_})
                    d_721_v21_ = d_721_v21_
                    if (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))]:
                        d_722_v22_: _dafny.Seq
                        d_722_v22_ = _dafny.SeqWithoutIsStrInference([p2, 601])
                        r1 = ((d_722_v22_)[default__.safeIndex(p0, len(d_722_v22_))]) + (p0)
                        r1 = p0
                        d_723_v23_: _dafny.Seq
                        d_723_v23_ = _dafny.SeqWithoutIsStrInference([d_694_v1_])
                        d_724_v24_: _dafny.Seq
                        d_724_v24_ = _dafny.SeqWithoutIsStrInference([d_723_v23_])
                        d_725_v25_: D2
                        d_725_v25_ = D2_DC7(D2_DC4(d_724_v24_))
                        d_725_v25_ = d_725_v25_
                        d_726_v26_: _dafny.Seq
                        d_727_v27_: bool
                        d_728_v28_: int
                        d_729_v29_: bool
                        out79_: _dafny.Seq
                        out80_: bool
                        out81_: int
                        out82_: bool
                        out79_, out80_, out81_, out82_ = default__.m0(d_694_v1_, d_699_v6_, (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))], globalState)
                        d_726_v26_ = out79_
                        d_727_v27_ = out80_
                        d_728_v28_ = out81_
                        d_729_v29_ = out82_
                        (globalState).f4 = (929) >= (608)
                    elif True:
                        index91_ = default__.safeIndex(306, (d_690_v0_).length(0))
                        (d_690_v0_)[index91_] = d_694_v1_
                        d_730_v30_: _dafny.Seq
                        d_730_v30_ = _dafny.SeqWithoutIsStrInference([d_690_v0_])
                        r1 = (0) - (default__.safeDivisionInt(len(d_730_v30_), (p0) + (p0)))
                        d_731_v31_: _dafny.MultiSet
                        d_731_v31_ = _dafny.MultiSet([d_694_v1_])
                        d_732_v32_: _dafny.Seq
                        d_732_v32_ = _dafny.SeqWithoutIsStrInference([(d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))]])
                        r1 = (self).fm14((((d_731_v31_).set((d_732_v32_)[default__.safeIndex(p0, len(d_732_v32_))], default__.abs(p2))).cardinality) - (p2), globalState)
                        d_698_v5_ = (d_698_v5_) - (_dafny.Set({p0}))
                        (globalState).f4 = (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))]
                    (globalState).f4 = d_694_v1_
                    pass
            pass
        d_694_v1_ = d_694_v1_
        d_739_v33_: _dafny.Array
        nw123_ = _dafny.Array(None, 22)
        nw123_[int(0)] = d_690_v0_
        nw123_[int(1)] = d_690_v0_
        nw123_[int(2)] = d_690_v0_
        nw123_[int(3)] = d_690_v0_
        nw123_[int(4)] = d_690_v0_
        nw123_[int(5)] = d_690_v0_
        nw123_[int(6)] = d_690_v0_
        nw123_[int(7)] = d_690_v0_
        nw123_[int(8)] = d_690_v0_
        nw123_[int(9)] = d_690_v0_
        nw123_[int(10)] = d_690_v0_
        nw123_[int(11)] = d_690_v0_
        nw123_[int(12)] = d_690_v0_
        nw123_[int(13)] = d_690_v0_
        nw123_[int(14)] = d_690_v0_
        nw123_[int(15)] = d_690_v0_
        nw123_[int(16)] = d_690_v0_
        nw123_[int(17)] = d_690_v0_
        nw123_[int(18)] = d_690_v0_
        nw123_[int(19)] = d_690_v0_
        nw123_[int(20)] = d_690_v0_
        nw123_[int(21)] = d_690_v0_
        d_739_v33_ = nw123_
        d_740_v34_: _dafny.Seq
        d_740_v34_ = _dafny.SeqWithoutIsStrInference([(d_739_v33_ if (d_690_v0_)[default__.safeIndex(306, (d_690_v0_).length(0))] else d_739_v33_)])
        r0 = (d_740_v34_)[default__.safeIndex(p0, len(d_740_v34_))]
        r1 = (default__.safeModuloInt((0) - (p2), p0)) - (p2)
        return r0, r1


class C1(T1, T0):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        self.f18: bool = False
        self._f17: bool = False
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
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f17, f18, f13, f14, f12):
        (self)._f17 = f17
        (self).f18 = f18
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f12 = f12

    def fm6(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([-74])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqldcxev")))]))

    def fm5(self, p0, globalState):
        return len((_dafny.Map({_dafny.MultiSet([False, (self).f17]): -572})) | (_dafny.Map({_dafny.MultiSet([self.f18]): -820})))

    def fm17(self, p0, p1, globalState):
        return (self).f17

    def m3(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_741_v0_: C0
        nw124_ = C0()
        nw124_.ctor__()
        d_741_v0_ = nw124_
        d_742_v1_: C0
        nw125_ = C0()
        nw125_.ctor__()
        d_742_v1_ = nw125_
        d_743_v2_: _dafny.Array
        nw126_ = _dafny.Array(None, 11)
        nw126_[int(0)] = d_742_v1_
        nw126_[int(1)] = d_741_v0_
        nw126_[int(2)] = d_742_v1_
        nw126_[int(3)] = (d_741_v0_ if self.f18 else d_742_v1_)
        nw126_[int(4)] = d_741_v0_
        nw126_[int(5)] = d_742_v1_
        nw126_[int(6)] = d_741_v0_
        nw126_[int(7)] = d_741_v0_
        nw126_[int(8)] = d_741_v0_
        nw126_[int(9)] = d_742_v1_
        nw126_[int(10)] = d_742_v1_
        d_743_v2_ = nw126_
        index92_ = default__.safeIndex(451, (d_743_v2_).length(0))
        (d_743_v2_)[index92_] = d_741_v0_
        index93_ = default__.safeIndex(451, (d_743_v2_).length(0))
        (d_743_v2_)[index93_] = d_741_v0_
        d_744_v3_: int
        d_744_v3_ = 18
        d_745_v4_: _dafny.Seq
        d_745_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgpt"))
        r1 = (((0) - (d_744_v3_) if (self).f17 else d_744_v3_) if (default__.fm2(self.f18, globalState)) or (not((self).f17)) else len(d_745_v4_))
        d_746_v5_: _dafny.Seq
        d_746_v5_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
        d_747_v6_: D2
        d_747_v6_ = D2_DC4(d_746_v5_)
        d_748_i0_: int
        d_748_i0_ = 0
        with _dafny.label("6"):
            pat_let_tv7_ = d_744_v3_
            pat_let_tv8_ = d_744_v3_
            def lambda44_(source10_):
                if source10_.is_DC5:
                    d_755___mcc_h0_ = source10_.cf5
                    d_756___mcc_h1_ = source10_.cf6
                    d_757___mcc_h2_ = source10_.cf7
                    d_758_cf7_ = d_757___mcc_h2_
                    d_759_cf6_ = d_756___mcc_h1_
                    d_760_cf5_ = d_755___mcc_h0_
                    return (_dafny.Set({d_758_cf7_})).issubset(_dafny.Set({d_758_cf7_, d_758_cf7_}))
                elif source10_.is_DC6:
                    d_761___mcc_h3_ = source10_.cf8
                    d_762___mcc_h4_ = source10_.cf9
                    d_763_cf9_ = d_762___mcc_h4_
                    d_764_cf8_ = d_761___mcc_h3_
                    return (self).f17
                elif source10_.is_DC4:
                    d_765___mcc_h5_ = source10_.cf4
                    d_766_cf4_ = d_765___mcc_h5_
                    return self.f18
                elif True:
                    d_767___mcc_h6_ = source10_.cf10
                    d_768_cf10_ = d_767___mcc_h6_
                    return (pat_let_tv7_) == (pat_let_tv8_)

            while lambda44_(d_747_v6_):
                with _dafny.c_label("6"):
                    if (d_748_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_748_i0_ = (d_748_i0_) + (1)
                    d_749_v7_: _dafny.Array
                    nw127_ = _dafny.Array(int(0), 26)
                    d_749_v7_ = nw127_
                    d_750_v8_: _dafny.Map
                    d_750_v8_ = _dafny.Map({(0) - (d_744_v3_): d_749_v7_})
                    d_750_v8_ = (d_750_v8_).set(d_744_v3_, d_749_v7_)
                    d_751_v9_: _dafny.Map
                    d_751_v9_ = _dafny.Map({self.f18: d_744_v3_})
                    d_751_v9_ = (d_751_v9_).set(default__.fm2(True, globalState), d_744_v3_)
                    d_752_v10_: _dafny.MultiSet
                    d_752_v10_ = _dafny.MultiSet([True, self.f18])
                    d_753_v11_: _dafny.Set
                    d_753_v11_ = _dafny.Set({((d_741_v0_).fm14(d_744_v3_, globalState) if (self).f17 else (d_752_v10_).cardinality), (0) - ((d_744_v3_) - (d_744_v3_)), d_744_v3_})
                    d_753_v11_ = (d_753_v11_).intersection((_dafny.Set({d_744_v3_})) | (d_753_v11_))
                    d_754_v12_: _dafny.Seq
                    d_754_v12_ = _dafny.SeqWithoutIsStrInference([d_749_v7_, d_749_v7_, d_749_v7_])
                    d_749_v7_ = (d_754_v12_)[default__.safeIndex(d_744_v3_, len(d_754_v12_))]
                    pass
            pass
        d_769_v13_: _dafny.MultiSet
        d_769_v13_ = _dafny.MultiSet([True, (self).f17])
        d_770_v14_: _dafny.Seq
        d_770_v14_ = _dafny.SeqWithoutIsStrInference([((d_769_v13_)[self.f18] if (self.f18) in (d_769_v13_) else d_744_v3_), d_744_v3_])
        d_771_v15_: D0
        d_771_v15_ = D0_DC2(d_744_v3_, (self).f17)
        d_772_v16_: _dafny.Map
        d_772_v16_ = _dafny.Map({default__.fm2((self).f17, globalState): (d_771_v15_).cf1})
        r3 = not(not((self).fm17(((d_770_v14_).set(default__.safeIndex(d_744_v3_, len(d_770_v14_)), d_744_v3_)) + (_dafny.SeqWithoutIsStrInference([d_744_v3_])), len(_dafny.Map({((d_772_v16_)[False] if (False) in (d_772_v16_) else len(d_770_v14_)): self.f18})), globalState)))
        r0 = default__.fm2((self).f17, globalState)
        r1 = d_744_v3_
        d_773_v17_: _dafny.Set
        d_773_v17_ = _dafny.Set({(self).f14, ((self).f14) + (default__.fm18(self.f18, d_744_v3_, False, len(_dafny.SeqWithoutIsStrInference([d_744_v3_])), globalState)), (self).f14, (self).f14, (_dafny.SeqWithoutIsStrInference([(self).fm17(d_770_v14_, len(d_745_v4_), globalState)])) + ((self).f14)})
        r2 = d_773_v17_
        r3 = self.f18
        return r0, r1, r2, r3

    @property
    def f17(self):
        return self._f17

class C2(T3, T0, T4, T2, T1):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f16: _dafny.Seq = _dafny.Seq({})
        self._f15: bool = False
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f16(self):
        return self._f16
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f16, f15, f13, f14, f12):
        (self)._f16 = f16
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f12 = f12

    def fm9(self, globalState):
        source11_ = D8_DC26((self).f15)
        if source11_.is_DC24:
            d_774___mcc_h0_ = source11_.cf36
            d_775_cf36_ = d_774___mcc_h0_
            return _dafny.MultiSet((self).f14)
        elif source11_.is_DC25:
            d_776___mcc_h1_ = source11_.cf37
            d_777___mcc_h2_ = source11_.cf38
            d_778___mcc_h3_ = source11_.cf39
            d_779___mcc_h4_ = source11_.cf40
            d_780_cf40_ = d_779___mcc_h4_
            d_781_cf39_ = d_778___mcc_h3_
            d_782_cf38_ = d_777___mcc_h2_
            d_783_cf37_ = d_776___mcc_h1_
            return _dafny.MultiSet([False, False])
        elif source11_.is_DC26:
            d_784___mcc_h5_ = source11_.cf41
            d_785_cf41_ = d_784___mcc_h5_
            if (self).f15:
                return _dafny.MultiSet([d_785_cf41_])
            elif True:
                return _dafny.MultiSet([(D0_DC0(d_785_cf41_)).cf0])
        elif source11_.is_DC23:
            d_786___mcc_h6_ = source11_.cf35
            d_787_cf35_ = d_786___mcc_h6_
            return _dafny.MultiSet([(self).f15, (self).f15, False])
        elif True:
            d_788___mcc_h7_ = source11_.cf42
            d_789_cf42_ = d_788___mcc_h7_
            return (_dafny.MultiSet((self).f14))

    def fm7(self, p0, globalState):
        def iife48_():
            coll40_ = _dafny.Map()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(209, 497):
                d_790_v0_: int = compr_40_
                if ((209) <= (d_790_v0_)) and ((d_790_v0_) < (497)):
                    coll40_[default__.safeDivisionInt(d_790_v0_, 494)] = (self).f15
            return _dafny.Map(coll40_)
        return default__.safeDivisionInt((len(iife48_()
        )) + (len(_dafny.Map({-566: (self).f16}))), 239)

    def fm8(self, p0, p1, globalState):
        return 844

    def fm6(self, p0, globalState):
        def iife49_():
            def iife51_():
                coll43_ = _dafny.Set()
                compr_43_: _dafny.Seq
                for compr_43_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjnkal")): (self).f15})).keys.Elements:
                    d_794_v1_: _dafny.Seq = compr_43_
                    if (d_794_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjnkal")): (self).f15})):
                        coll43_ = coll43_.union(_dafny.Set([d_794_v1_]))
                return _dafny.Set(coll43_)
            coll41_ = _dafny.Map()
            def iife50_():
                coll42_ = _dafny.Set()
                compr_42_: _dafny.Seq
                for compr_42_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjnkal")): (self).f15})).keys.Elements:
                    d_792_v1_: _dafny.Seq = compr_42_
                    if (d_792_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjnkal")): (self).f15})):
                        coll42_ = coll42_.union(_dafny.Set([d_792_v1_]))
                return _dafny.Set(coll42_)
            compr_41_: _dafny.Seq
            for compr_41_ in (iife50_()
            ).Elements:
                d_793_v0_: _dafny.Seq = compr_41_
                if (d_793_v0_) in (iife51_()
                ):
                    coll41_[d_793_v0_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oy")))
            return _dafny.Map(coll41_)
        return ((self).f16).set(default__.safeIndex((0) - (((0) - (len((D6_DC17(_dafny.Set({553, 605}))).cf27)) if True else (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nan")) for d_791_i0_ in range(default__.abs(517))]))))), len((self).f16)), len(iife49_()
        ))

    def fm5(self, p0, globalState):
        return len(((self).f14) + ((self).f14))

    def fm10(self, p0, p1, p2, p3, globalState):
        return _dafny.Set({-129, (_dafny.MultiSet([len((self).f14), len((self).f14), 527, 772])).cardinality})

    def fm38(self, p0, p1, p2, p3, globalState):
        return (0) - ((0) - ((len(((D3_DC8(_dafny.Map({(self).f15: (self).f15}))).cf11) | (_dafny.Map({False: (self).f15})))) * (default__.safeDivisionInt(999, 912))))

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_795_v0_: _dafny.Map
        d_795_v0_ = _dafny.Map({(self).f15: (self).f15})
        d_796_v1_: D3
        d_796_v1_ = D3_DC8(d_795_v0_)
        d_797_v3_: _dafny.Map
        d_797_v3_ = _dafny.Map({p1: (self).f15})
        d_798_v4_: _dafny.Seq
        def iife52_():
            coll44_ = _dafny.Map()
            compr_44_: int
            for compr_44_ in (d_797_v3_).keys.Elements:
                d_799_v2_: int = compr_44_
                if (d_799_v2_) in (d_797_v3_):
                    coll44_[(d_799_v2_) + (p1)] = d_796_v1_
            return _dafny.Map(coll44_)
        d_798_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: d_796_v1_}), iife52_()
        ])
        d_800_v5_: _dafny.Map
        d_800_v5_ = _dafny.Map({724: D3_DC8(d_795_v0_)})
        (globalState).f4 = not(((d_798_v4_) < ((d_798_v4_).set(default__.safeIndex(p0, len(d_798_v4_)), d_800_v5_))) == ((self).f15))
        hi4_ = 414
        for d_801_i0_ in range(len((self).f14), hi4_):
            d_802_v6_: str
            d_802_v6_ = _dafny.CodePoint('m')
            d_803_v7_: _dafny.Map
            d_803_v7_ = _dafny.Map({(self).f15: d_802_v6_})
            d_804_v8_: _dafny.Array
            nw128_ = _dafny.Array(None, 13)
            nw128_[int(0)] = _dafny.CodePoint('j')
            nw128_[int(1)] = d_802_v6_
            nw128_[int(2)] = d_802_v6_
            nw128_[int(3)] = (p2)[default__.safeIndex(d_801_i0_, len(p2))]
            nw128_[int(4)] = d_802_v6_
            nw128_[int(5)] = d_802_v6_
            nw128_[int(6)] = d_802_v6_
            nw128_[int(7)] = d_802_v6_
            nw128_[int(8)] = d_802_v6_
            nw128_[int(9)] = ((d_803_v7_)[False] if (False) in (d_803_v7_) else d_802_v6_)
            nw128_[int(10)] = (d_802_v6_ if (self).f15 else d_802_v6_)
            nw128_[int(11)] = d_802_v6_
            nw128_[int(12)] = d_802_v6_
            d_804_v8_ = nw128_
            index94_ = default__.safeIndex(344, (d_804_v8_).length(0))
            (d_804_v8_)[index94_] = (d_802_v6_ if False else default__.fm39(p1, not((self).f15), p0, p2, globalState))
            index95_ = default__.safeIndex(344, (d_804_v8_).length(0))
            (d_804_v8_)[index95_] = d_802_v6_
            d_805_v9_: _dafny.MultiSet
            d_805_v9_ = _dafny.MultiSet([(len(p2) if (self).f15 else (0) - (p0))])
            d_805_v9_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))
            (globalState).f4 = not((self).f15)
            r0 = (0) - ((p0) + ((p0) - (p0)))
        d_806_v10_: _dafny.Array
        def lambda45_(d_807_i1_):
            return (self).f15

        init23_ = lambda45_
        nw129_ = _dafny.Array(None, 13)
        for i0_23_ in range(nw129_.length(0)):
            nw129_[i0_23_] = init23_(i0_23_)
        d_806_v10_ = nw129_
        index96_ = default__.safeIndex(859, (d_806_v10_).length(0))
        (d_806_v10_)[index96_] = (self).f15
        d_808_v11_: str
        d_808_v11_ = _dafny.CodePoint('h')
        index97_ = default__.safeIndex(859, (d_806_v10_).length(0))
        (d_806_v10_)[index97_] = (d_808_v11_) in (p2)
        (globalState).f4 = False
        d_809_v12_: D0
        d_809_v12_ = D0_DC2(p1, (d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))])
        d_810_i2_: int
        d_810_i2_ = 0
        with _dafny.label("7"):
            pat_let_tv9_ = p1
            def lambda46_(source12_):
                if source12_.is_DC1:
                    return (self).f15
                elif source12_.is_DC2:
                    d_824___mcc_h0_ = source12_.cf1
                    d_825___mcc_h1_ = source12_.cf2
                    d_826_cf2_ = d_825___mcc_h1_
                    d_827_cf1_ = d_824___mcc_h0_
                    return (_dafny.Set({(self).f16})).isdisjoint(_dafny.Set({(self).f16}))
                elif True:
                    d_828___mcc_h2_ = source12_.cf0
                    d_829_cf0_ = d_828___mcc_h2_
                    return not((_dafny.SeqWithoutIsStrInference([d_829_cf0_])) == ((self).f14))

            def iife53_(_pat_let4_0):
                def iife54_(d_830_dt__update__tmp_h0_):
                    def iife55_(_pat_let5_0):
                        def iife56_(d_831_dt__update_hcf1_h0_):
                            return D0_DC2(d_831_dt__update_hcf1_h0_, (d_830_dt__update__tmp_h0_).cf2)
                        return iife56_(_pat_let5_0)
                    return iife55_(pat_let_tv9_)
                return iife54_(_pat_let4_0)
            while lambda46_(iife53_(d_809_v12_)):
                with _dafny.c_label("7"):
                    if (d_810_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_810_i2_ = (d_810_i2_) + (1)
                    r0 = len(p2)
                    d_811_v13_: _dafny.Map
                    d_811_v13_ = _dafny.Map({(self).f15: p1})
                    d_812_v14_: _dafny.Seq
                    d_812_v14_ = _dafny.SeqWithoutIsStrInference([d_811_v13_, d_811_v13_, d_811_v13_])
                    d_813_v15_: D8
                    d_813_v15_ = D8_DC23(_dafny.Map({(self).f15: p0}))
                    d_814_v16_: _dafny.Map
                    d_814_v16_ = _dafny.Map({len(p2): p1})
                    d_815_v17_: _dafny.Array
                    nw130_ = _dafny.Array(None, 15)
                    nw130_[int(0)] = (d_812_v14_)[default__.safeIndex(p1, len(d_812_v14_))]
                    nw130_[int(1)] = (d_811_v13_) | (d_811_v13_)
                    nw130_[int(2)] = (d_811_v13_).set((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], p0)
                    nw130_[int(3)] = (_dafny.Map({(self).f15: 827}) if False else _dafny.Map({(self).f15: 235}))
                    nw130_[int(4)] = (d_813_v15_).cf35
                    nw130_[int(5)] = d_811_v13_
                    nw130_[int(6)] = (d_811_v13_) | (d_811_v13_)
                    nw130_[int(7)] = (d_811_v13_).set((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], len((self).f16))
                    nw130_[int(8)] = (d_811_v13_).set((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], p0)
                    nw130_[int(9)] = (d_811_v13_).set((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], (p3).cardinality)
                    nw130_[int(10)] = d_811_v13_
                    nw130_[int(11)] = (d_811_v13_).set((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], ((d_814_v16_)[p1] if (p1) in (d_814_v16_) else p1))
                    nw130_[int(12)] = (d_811_v13_).set((self).f15, p0)
                    nw130_[int(13)] = d_811_v13_
                    nw130_[int(14)] = (_dafny.Map({(d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))]: 97}) if (self).f15 else _dafny.Map({(self).f15: p1}))
                    d_815_v17_ = nw130_
                    index98_ = default__.safeIndex(699, (d_815_v17_).length(0))
                    (d_815_v17_)[index98_] = d_811_v13_
                    index99_ = default__.safeIndex(699, (d_815_v17_).length(0))
                    (d_815_v17_)[index99_] = (d_811_v13_) | (d_811_v13_)
                    d_816_v18_: T1
                    nw131_ = C1()
                    nw131_.ctor__(not((self).f15), (self).f15, (self).f13, (self).f14, self.f12)
                    d_816_v18_ = nw131_
                    d_817_v19_: _dafny.MultiSet
                    d_817_v19_ = _dafny.MultiSet([d_816_v18_])
                    d_818_v20_: _dafny.Seq
                    d_818_v20_ = _dafny.SeqWithoutIsStrInference([d_816_v18_, d_816_v18_])
                    d_819_v21_: _dafny.Map
                    d_819_v21_ = _dafny.Map({(d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))]: d_816_v18_})
                    d_820_v22_: _dafny.Map
                    d_820_v22_ = _dafny.Map({d_808_v11_: (d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))]})
                    d_821_v23_: _dafny.Array
                    nw132_ = _dafny.Array(None, 19)
                    nw132_[int(0)] = d_817_v19_
                    nw132_[int(1)] = d_817_v19_
                    nw132_[int(2)] = d_817_v19_
                    nw132_[int(3)] = d_817_v19_
                    nw132_[int(4)] = _dafny.MultiSet(d_818_v20_)
                    nw132_[int(5)] = d_817_v19_
                    nw132_[int(6)] = d_817_v19_
                    nw132_[int(7)] = d_817_v19_
                    nw132_[int(8)] = (_dafny.MultiSet([d_816_v18_, d_816_v18_, d_816_v18_, d_816_v18_, d_816_v18_])) - (d_817_v19_)
                    nw132_[int(9)] = d_817_v19_
                    nw132_[int(10)] = d_817_v19_
                    nw132_[int(11)] = _dafny.MultiSet([d_816_v18_, ((d_819_v21_)[(self).f15] if ((self).f15) in (d_819_v21_) else d_816_v18_), d_816_v18_])
                    nw132_[int(12)] = _dafny.MultiSet([d_816_v18_, d_816_v18_])
                    nw132_[int(13)] = _dafny.MultiSet([d_816_v18_, d_816_v18_, d_816_v18_])
                    nw132_[int(14)] = _dafny.MultiSet([d_816_v18_, ((d_819_v21_)[((d_820_v22_)[d_808_v11_] if (d_808_v11_) in (d_820_v22_) else (d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))])] if (((d_820_v22_)[d_808_v11_] if (d_808_v11_) in (d_820_v22_) else (d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))])) in (d_819_v21_) else d_816_v18_)])
                    nw132_[int(15)] = (d_817_v19_) - (d_817_v19_)
                    nw132_[int(16)] = d_817_v19_
                    nw132_[int(17)] = d_817_v19_
                    nw132_[int(18)] = d_817_v19_
                    d_821_v23_ = nw132_
                    d_822_v24_: _dafny.Map
                    d_822_v24_ = _dafny.Map({(d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))]: d_821_v23_})
                    d_823_v25_: _dafny.Seq
                    d_823_v25_ = _dafny.SeqWithoutIsStrInference([((d_822_v24_)[default__.fm2((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], globalState)] if (default__.fm2((d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], globalState)) in (d_822_v24_) else d_821_v23_), d_821_v23_, d_821_v23_])
                    rhs67_ = p1
                    rhs68_ = p0
                    rhs69_ = (d_823_v25_)[default__.safeIndex(p1, len(d_823_v25_))]
                    r0 = rhs67_
                    r0 = rhs68_
                    d_821_v23_ = rhs69_
                    index100_ = default__.safeIndex(859, (d_806_v10_).length(0))
                    (d_806_v10_)[index100_] = default__.fm2(((self).f14) == ((d_816_v18_).f14), globalState)
                    pass
            pass
        r0 = default__.safeModuloInt(len((p2).set(default__.safeIndex((0) - (p1), len(p2)), d_808_v11_)), p1)
        r0 = (self).fm38(p1, (0) - (((self).f16)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qifnlwcgl"))), len((self).f16))]), p0, (d_806_v10_)[default__.safeIndex(859, (d_806_v10_).length(0))], globalState)
        return r0

    def m3(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_832_v0_: D2
        d_832_v0_ = D2_DC4(_dafny.SeqWithoutIsStrInference([(self).f14 for d_833_i0_ in range(default__.abs(27))]))
        d_832_v0_ = d_832_v0_
        d_834_v1_: int
        d_834_v1_ = -924
        d_835_v2_: _dafny.Set
        d_835_v2_ = _dafny.Set({d_834_v1_})
        d_836_v3_: D8
        d_836_v3_ = D8_DC24((self).f15)
        d_837_v4_: D8
        d_837_v4_ = D8_DC27(d_836_v3_)
        d_838_v5_: str
        d_838_v5_ = _dafny.CodePoint('d')
        d_839_v6_: _dafny.Map
        d_839_v6_ = _dafny.Map({d_837_v4_: d_838_v5_})
        d_840_v7_: _dafny.Seq
        d_840_v7_ = _dafny.SeqWithoutIsStrInference([d_839_v6_])
        d_841_v8_: _dafny.Seq
        d_841_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjlnmny"))
        d_842_v9_: _dafny.Array
        nw133_ = _dafny.Array(None, 27)
        nw133_[int(0)] = (self).f15
        nw133_[int(1)] = (self).f15
        nw133_[int(2)] = (d_835_v2_).issubset(d_835_v2_)
        nw133_[int(3)] = not((self).f15)
        nw133_[int(4)] = ((self).f15) and (((self).f14)[default__.safeIndex(d_834_v1_, len((self).f14))])
        nw133_[int(5)] = (self).f15
        nw133_[int(6)] = (d_834_v1_) < (len(d_840_v7_))
        nw133_[int(7)] = (self).f15
        nw133_[int(8)] = default__.fm2((self).f15, globalState)
        nw133_[int(9)] = not ((self).f15) or ((self).f15)
        nw133_[int(10)] = (self).f15
        nw133_[int(11)] = (self).f15
        nw133_[int(12)] = (self).f15
        nw133_[int(13)] = (self).f15
        nw133_[int(14)] = (self).f15
        nw133_[int(15)] = (self).f15
        nw133_[int(16)] = (d_841_v8_) != (d_841_v8_)
        nw133_[int(17)] = False
        nw133_[int(18)] = (self).f15
        nw133_[int(19)] = False
        nw133_[int(20)] = (self).f15
        nw133_[int(21)] = False
        nw133_[int(22)] = (self).f15
        nw133_[int(23)] = (self).f15
        nw133_[int(24)] = (d_834_v1_) > (len(_dafny.Map({True: 722})))
        nw133_[int(25)] = (self).f15
        nw133_[int(26)] = False
        d_842_v9_ = nw133_
        d_843_v10_: D7
        d_843_v10_ = D7_DC22((self).f15)
        index101_ = default__.safeIndex(4, (d_842_v9_).length(0))
        (d_842_v9_)[index101_] = (d_843_v10_).cf34
        d_844_v11_: _dafny.MultiSet
        d_844_v11_ = _dafny.MultiSet([d_834_v1_, d_834_v1_])
        index102_ = default__.safeIndex(4, (d_842_v9_).length(0))
        (d_842_v9_)[index102_] = ((_dafny.MultiSet([d_834_v1_, d_834_v1_, d_834_v1_])) - ((d_844_v11_).set(d_834_v1_, default__.abs(len(d_841_v8_))))).ispropersubset(d_844_v11_)
        (self).m12(not((d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))]), globalState)
        d_845_v12_: _dafny.Map
        d_845_v12_ = _dafny.Map({(d_834_v1_ if (d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))] else 271): d_834_v1_})
        d_845_v12_ = (d_845_v12_).set(d_834_v1_, d_834_v1_)
        d_846_v13_: _dafny.Map
        d_846_v13_ = _dafny.Map({default__.fm2((d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))], globalState): (d_841_v8_).set(default__.safeIndex(len((self).f16), len(d_841_v8_)), d_838_v5_)})
        d_847_v14_: _dafny.Set
        d_847_v14_ = _dafny.Set({(d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))], (d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))], (d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))]})
        d_848_v15_: _dafny.Seq
        d_848_v15_ = _dafny.SeqWithoutIsStrInference([len(d_846_v13_), len((d_847_v14_).intersection(_dafny.Set({False, (d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))]}))), ((d_844_v11_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jefu")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jefu")))) in (d_844_v11_) else d_834_v1_), 520, (len((self).f14)) - (d_834_v1_)])
        rhs70_ = d_838_v5_
        rhs71_ = ((self).f16).set(default__.safeIndex(d_834_v1_, len((self).f16)), d_834_v1_)
        d_838_v5_ = rhs70_
        d_848_v15_ = rhs71_
        d_849_v16_: D4
        d_849_v16_ = D4_DC14((d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))])
        d_841_v8_ = (d_841_v8_ if not(False) else default__.fm40(d_849_v16_, d_834_v1_, True, (self).f15, globalState))
        r0 = (d_834_v1_) > (default__.safeModuloInt(d_834_v1_, len(d_845_v12_)))
        d_850_v17_: _dafny.MultiSet
        d_850_v17_ = _dafny.MultiSet([(self).f14])
        r1 = ((d_850_v17_).cardinality) * (d_834_v1_)
        r2 = default__.fm41((len(d_846_v13_)) - ((0) - (d_834_v1_)), d_838_v5_, globalState)
        r3 = (d_842_v9_)[default__.safeIndex(4, (d_842_v9_).length(0))]
        return r0, r1, r2, r3

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        d_851_v0_: str
        d_851_v0_ = _dafny.CodePoint('t')
        d_852_v1_: _dafny.MultiSet
        d_852_v1_ = _dafny.MultiSet([len((self).f14), len((self).f14), 660])
        d_853_v2_: _dafny.Map
        d_853_v2_ = _dafny.Map({d_851_v0_: d_852_v1_})
        d_854_v3_: _dafny.Map
        d_854_v3_ = _dafny.Map({p2: default__.fm2((self).f15, globalState)})
        d_855_v4_: D7
        d_855_v4_ = D7_DC21(d_854_v3_, (self).f15, (self).f15)
        d_853_v2_ = default__.fm42(d_855_v4_, p1, globalState)
        d_856_v5_: _dafny.MultiSet
        d_856_v5_ = _dafny.MultiSet([(self).f15])
        hi5_ = len((p1) + (p1))
        for d_857_i0_ in range(((d_856_v5_)[(self).f15] if ((self).f15) in (d_856_v5_) else p2), hi5_):
            d_858_v6_: _dafny.Array
            nw134_ = _dafny.Array(False, 21)
            d_858_v6_ = nw134_
            pat_let_tv10_ = d_858_v6_
            def iife57_(_pat_let6_0):
                def iife58_(d_859_dt__update__tmp_h0_):
                    def iife59_(_pat_let7_0):
                        def iife60_(d_860_dt__update_hcf8_h0_):
                            return D2_DC6(d_860_dt__update_hcf8_h0_, (d_859_dt__update__tmp_h0_).cf9)
                        return iife60_(_pat_let7_0)
                    return iife59_(pat_let_tv10_)
                return iife58_(_pat_let6_0)
            source13_ = iife57_(D2_DC6(d_858_v6_, d_858_v6_))
            if source13_.is_DC5:
                d_861___mcc_h0_ = source13_.cf5
                d_862___mcc_h1_ = source13_.cf6
                d_863___mcc_h2_ = source13_.cf7
                d_864_cf7_ = d_863___mcc_h2_
                d_865_cf6_ = d_862___mcc_h1_
                d_866_cf5_ = d_861___mcc_h0_
                d_867_v8_: _dafny.Array
                nw135_ = _dafny.Array(None, 19)
                nw135_[int(0)] = p2
                nw135_[int(1)] = d_866_cf5_
                nw135_[int(2)] = (d_866_cf5_ if (self).f15 else (self).fm5(d_865_cf6_, globalState))
                nw135_[int(3)] = -356
                nw135_[int(4)] = d_857_i0_
                nw135_[int(5)] = d_866_cf5_
                nw135_[int(6)] = (d_866_cf5_ if not(d_864_cf7_) else d_866_cf5_)
                nw135_[int(7)] = p2
                nw135_[int(8)] = (p2) + (d_866_cf5_)
                nw135_[int(9)] = p2
                nw135_[int(10)] = len(p1)
                nw135_[int(11)] = d_857_i0_
                nw135_[int(12)] = len(_dafny.SeqWithoutIsStrInference([d_851_v0_ for d_868_i1_ in range(default__.abs(754))]))
                nw135_[int(13)] = (209) - (d_857_i0_)
                nw135_[int(14)] = p2
                nw135_[int(15)] = d_866_cf5_
                def iife61_():
                    coll45_ = _dafny.Map()
                    compr_45_: int
                    for compr_45_ in _dafny.IntegerRange(624, 717):
                        d_869_v7_: int = compr_45_
                        if ((624) <= (d_869_v7_)) and ((d_869_v7_) < (717)):
                            coll45_[(d_869_v7_) - (p2)] = d_851_v0_
                    return _dafny.Map(coll45_)
                nw135_[int(16)] = default__.safeModuloInt(len(iife61_()
                ), p2)
                nw135_[int(17)] = default__.safeModuloInt(d_866_cf5_, len(p1))
                nw135_[int(18)] = d_857_i0_
                d_867_v8_ = nw135_
                d_867_v8_ = d_867_v8_
                d_866_cf5_ = default__.safeDivisionInt(p2, (d_866_cf5_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmitbisb")))))
                (globalState).f4 = d_864_cf7_
                d_866_cf5_ = (-39) + (p2)
            elif source13_.is_DC6:
                d_870___mcc_h3_ = source13_.cf8
                d_871___mcc_h4_ = source13_.cf9
                d_872_cf9_ = d_871___mcc_h4_
                d_873_cf8_ = d_870___mcc_h3_
                d_874_v9_: _dafny.Seq
                d_874_v9_ = _dafny.SeqWithoutIsStrInference([(60) * (d_857_i0_)])
                d_874_v9_ = (_dafny.SeqWithoutIsStrInference([p2, p2])) + (d_874_v9_)
                d_875_v10_: _dafny.MultiSet
                d_875_v10_ = d_856_v5_
                d_876_v11_: D7
                d_876_v11_ = D7_DC22((self).f15)
                d_877_v12_: _dafny.Seq
                d_877_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm43(d_875_v10_, d_876_v11_, globalState)])
                index103_ = default__.safeIndex(352, (d_858_v6_).length(0))
                (d_858_v6_)[index103_] = ((d_877_v12_).set(default__.safeIndex(len(p0), len(d_877_v12_)), (self).f14)) <= (d_877_v12_)
                index104_ = default__.safeIndex(352, (d_858_v6_).length(0))
                (d_858_v6_)[index104_] = default__.fm2((self).f15, globalState)
                d_878_v13_: D5
                d_878_v13_ = D5_DC16((d_858_v6_)[default__.safeIndex(352, (d_858_v6_).length(0))])
                d_879_v14_: _dafny.Map
                d_879_v14_ = _dafny.Map({d_878_v13_: self.f12})
                d_880_v15_: C1
                nw136_ = C1()
                nw136_.ctor__((d_858_v6_)[default__.safeIndex(352, (d_858_v6_).length(0))], (d_858_v6_)[default__.safeIndex(352, (d_858_v6_).length(0))], (self).f13, default__.fm43(d_875_v10_, d_876_v11_, globalState), ((d_879_v14_)[d_878_v13_] if (d_878_v13_) in (d_879_v14_) else self.f12))
                d_880_v15_ = nw136_
                d_881_v16_: _dafny.Array
                def lambda47_(d_882_i0_):
                    def lambda48_(d_883_i2_):
                        return default__.safeDivisionInt(d_883_i2_, d_882_i0_)

                    return lambda48_

                init24_ = lambda47_(d_857_i0_)
                nw137_ = _dafny.Array(None, 20)
                for i0_24_ in range(nw137_.length(0)):
                    nw137_[i0_24_] = init24_(i0_24_)
                d_881_v16_ = nw137_
                d_884_v17_: _dafny.Seq
                d_884_v17_ = _dafny.SeqWithoutIsStrInference([d_881_v16_, d_881_v16_])
                d_885_v18_: _dafny.Map
                d_885_v18_ = _dafny.Map({False: d_884_v17_})
                d_886_v19_: _dafny.Map
                d_886_v19_ = _dafny.Map({p2: d_881_v16_})
                index105_ = default__.safeIndex(352, (d_858_v6_).length(0))
                (d_858_v6_)[index105_] = (((((d_885_v18_)[d_880_v15_.f18] if (d_880_v15_.f18) in (d_885_v18_) else _dafny.SeqWithoutIsStrInference([d_881_v16_, ((d_886_v19_)[p2] if (p2) in (d_886_v19_) else d_881_v16_)]))).set(default__.safeIndex(((d_852_v1_)[p2] if (p2) in (d_852_v1_) else -225), len(((d_885_v18_)[d_880_v15_.f18] if (d_880_v15_.f18) in (d_885_v18_) else _dafny.SeqWithoutIsStrInference([d_881_v16_, ((d_886_v19_)[p2] if (p2) in (d_886_v19_) else d_881_v16_)])))), d_881_v16_)) + (d_884_v17_)) == (d_884_v17_)
            elif source13_.is_DC4:
                d_887___mcc_h5_ = source13_.cf4
                d_888_cf4_ = d_887___mcc_h5_
                (globalState).f4 = (d_854_v3_) != (d_854_v3_)
                d_889_v20_: _dafny.Seq
                d_889_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdnbe"))
                d_889_v20_ = ((d_889_v20_) + (_dafny.SeqWithoutIsStrInference([d_851_v0_ for d_890_i3_ in range(default__.abs(233))]))) + ((default__.fm40(D4_DC14((self).f15), d_857_i0_, (self).f15, (self).f15, globalState)) + (_dafny.SeqWithoutIsStrInference([d_851_v0_ for d_891_i4_ in range(default__.abs(-820))])))
                d_892_v21_: str
                d_892_v21_ = d_851_v0_
                d_892_v21_ = d_892_v21_
                d_893_v22_: int
                d_893_v22_ = 123
                d_893_v22_ = d_893_v22_
            elif True:
                d_894___mcc_h6_ = source13_.cf10
                d_895_cf10_ = d_894___mcc_h6_
                d_896_v23_: _dafny.Array
                nw138_ = _dafny.Array(int(0), 25)
                d_896_v23_ = nw138_
                index106_ = default__.safeIndex(973, (d_896_v23_).length(0))
                (d_896_v23_)[index106_] = (0) - (len((self).f16))
                index107_ = default__.safeIndex(973, (d_896_v23_).length(0))
                (d_896_v23_)[index107_] = (0) - ((d_857_i0_) - (p2))
                d_897_v24_: _dafny.Seq
                d_897_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ak"))
                d_897_v24_ = d_897_v24_
                d_858_v6_ = (d_858_v6_ if (self).f15 else d_858_v6_)
                index108_ = default__.safeIndex(973, (d_896_v23_).length(0))
                (d_896_v23_)[index108_] = (0) - (((d_896_v23_)[default__.safeIndex(973, (d_896_v23_).length(0))]) * (p2))
            r0 = (((self).f16) + ((self).f16)) <= ((self).f16)
            d_898_v25_: _dafny.Array
            nw139_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_898_v25_ = nw139_
            d_899_v26_: _dafny.Array
            def lambda49_(d_900_p2_):
                def lambda50_(d_901_i5_):
                    return (d_901_i5_) + (d_900_p2_)

                return lambda50_

            init25_ = lambda49_(p2)
            nw140_ = _dafny.Array(None, 3)
            for i0_25_ in range(nw140_.length(0)):
                nw140_[i0_25_] = init25_(i0_25_)
            d_899_v26_ = nw140_
            index109_ = default__.safeIndex(630, (d_898_v25_).length(0))
            (d_898_v25_)[index109_] = d_899_v26_
            index110_ = default__.safeIndex(630, (d_898_v25_).length(0))
            (d_898_v25_)[index110_] = d_899_v26_
            (globalState).f4 = (len(p0)) < (p2)
        d_902_v27_: D4
        d_902_v27_ = D4_DC12((self).f14)
        source14_ = default__.fm44(d_902_v27_, (self).f15, globalState)
        if source14_.is_DC5:
            d_903___mcc_h7_ = source14_.cf5
            d_904___mcc_h8_ = source14_.cf6
            d_905___mcc_h9_ = source14_.cf7
            d_906_cf7_ = d_905___mcc_h9_
            d_907_cf6_ = d_904___mcc_h8_
            d_908_cf5_ = d_903___mcc_h7_
            d_908_cf5_ = (d_908_cf5_) - (p2)
            d_909_v28_: _dafny.Array
            def lambda51_(d_910_p2_):
                def lambda52_(d_911_i6_):
                    return (d_911_i6_) + (d_910_p2_)

                return lambda52_

            init26_ = lambda51_(p2)
            nw141_ = _dafny.Array(None, 29)
            for i0_26_ in range(nw141_.length(0)):
                nw141_[i0_26_] = init26_(i0_26_)
            d_909_v28_ = nw141_
            index111_ = default__.safeIndex(816, (d_909_v28_).length(0))
            (d_909_v28_)[index111_] = (0) - (d_908_cf5_)
            d_912_v29_: _dafny.Map
            d_912_v29_ = _dafny.Map({d_851_v0_: p2})
            index112_ = default__.safeIndex(816, (d_909_v28_).length(0))
            rhs72_ = d_907_cf6_
            rhs73_ = 927
            rhs74_ = d_908_cf5_
            rhs75_ = (((d_912_v29_)[d_851_v0_] if (d_851_v0_) in (d_912_v29_) else p2)) + (p2)
            lhs41_ = d_909_v28_
            lhs42_ = default__.safeIndex(816, (d_909_v28_).length(0))
            r0 = rhs72_
            lhs41_[lhs42_] = rhs73_
            d_908_cf5_ = rhs74_
            d_908_cf5_ = rhs75_
            d_913_v30_: _dafny.Map
            d_913_v30_ = _dafny.Map({(d_909_v28_)[default__.safeIndex(816, (d_909_v28_).length(0))]: p2})
            d_914_v31_: _dafny.Map
            d_914_v31_ = _dafny.Map({d_913_v30_: d_851_v0_})
            d_915_v32_: _dafny.Map
            d_915_v32_ = _dafny.Map({d_854_v3_: len(d_914_v31_)})
            d_916_v33_: _dafny.Map
            d_916_v33_ = _dafny.Map({(self).f15: len(d_915_v32_)})
            d_916_v33_ = (d_916_v33_).set(not ((self).f15) or ((self).f15), ((0) - (p2) if ((d_854_v3_)[(0) - ((d_909_v28_)[default__.safeIndex(816, (d_909_v28_).length(0))])] if ((0) - ((d_909_v28_)[default__.safeIndex(816, (d_909_v28_).length(0))])) in (d_854_v3_) else d_907_cf6_) else p2))
            index113_ = default__.safeIndex(816, (d_909_v28_).length(0))
            (d_909_v28_)[index113_] = d_908_cf5_
        elif source14_.is_DC6:
            d_917___mcc_h10_ = source14_.cf8
            d_918___mcc_h11_ = source14_.cf9
            d_919_cf9_ = d_918___mcc_h11_
            d_920_cf8_ = d_917___mcc_h10_
            d_921_v34_: _dafny.Set
            d_921_v34_ = _dafny.Set({p2, p2})
            r0 = not((d_921_v34_).isdisjoint((d_921_v34_) | (_dafny.Set({p2}))))
            d_851_v0_ = d_851_v0_
            d_922_v35_: _dafny.Seq
            d_922_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvd"))
            d_922_v35_ = d_922_v35_
            rhs76_ = (p2) < (859)
            rhs77_ = (self).f15
            lhs43_ = globalState
            lhs44_ = globalState
            lhs43_.f4 = rhs76_
            lhs44_.f4 = rhs77_
        elif source14_.is_DC4:
            d_923___mcc_h12_ = source14_.cf4
            d_924_cf4_ = d_923___mcc_h12_
            d_925_v36_: D0
            d_925_v36_ = D0_DC2(len((self).f16), (self).f15)
            d_926_v37_: _dafny.Seq
            d_927_v38_: bool
            d_928_v39_: int
            d_929_v40_: bool
            out83_: _dafny.Seq
            out84_: bool
            out85_: int
            out86_: bool
            def iife62_(_pat_let8_0):
                def iife63_(d_930_dt__update__tmp_h1_):
                    def iife64_(_pat_let9_0):
                        def iife65_(d_931_dt__update_hcf2_h0_):
                            return D0_DC2((d_930_dt__update__tmp_h1_).cf1, d_931_dt__update_hcf2_h0_)
                        return iife65_(_pat_let9_0)
                    return iife64_((self).f15)
                return iife63_(_pat_let8_0)
            out83_, out84_, out85_, out86_ = default__.m0(((self).f15 if (self).f15 else (self).f15), iife62_(d_925_v36_), (self).f15, globalState)
            d_926_v37_ = out83_
            d_927_v38_ = out84_
            d_928_v39_ = out85_
            d_929_v40_ = out86_
            d_902_v27_ = d_902_v27_
            d_932_v41_: C0
            nw142_ = C0()
            nw142_.ctor__()
            d_932_v41_ = nw142_
            d_933_v42_: _dafny.Array
            def lambda53_(d_934_v39_):
                def lambda54_(d_935_i7_):
                    return (d_935_i7_) + (d_934_v39_)

                return lambda54_

            init27_ = lambda53_(d_928_v39_)
            nw143_ = _dafny.Array(None, 8)
            for i0_27_ in range(nw143_.length(0)):
                nw143_[i0_27_] = init27_(i0_27_)
            d_933_v42_ = nw143_
            index114_ = default__.safeIndex(818, (d_933_v42_).length(0))
            (d_933_v42_)[index114_] = p2
            index115_ = default__.safeIndex(818, (d_933_v42_).length(0))
            (d_933_v42_)[index115_] = ((d_856_v5_)[d_927_v38_] if (d_927_v38_) in (d_856_v5_) else len(_dafny.SeqWithoutIsStrInference([d_851_v0_ for d_936_i8_ in range(default__.abs(368))])))
        elif True:
            d_937___mcc_h13_ = source14_.cf10
            d_938_cf10_ = d_937___mcc_h13_
            d_939_v43_: _dafny.Map
            d_939_v43_ = _dafny.Map({(self).f15: (0) - (len(p1))})
            d_940_v44_: _dafny.Seq
            d_940_v44_ = _dafny.SeqWithoutIsStrInference([d_939_v43_, d_939_v43_, default__.fm45((self).f15, ((self).f16)[default__.safeIndex(p2, len((self).f16))], p2, globalState), (_dafny.Map({True: p2})) | (d_939_v43_)])
            d_940_v44_ = ((_dafny.SeqWithoutIsStrInference([d_939_v43_, d_939_v43_])) + (_dafny.SeqWithoutIsStrInference([d_939_v43_]))).set(default__.safeIndex(len((d_939_v43_) | (d_939_v43_)), len((_dafny.SeqWithoutIsStrInference([d_939_v43_, d_939_v43_])) + (_dafny.SeqWithoutIsStrInference([d_939_v43_])))), d_939_v43_)
            d_941_v45_: int
            d_941_v45_ = 138
            d_941_v45_ = (0) - (d_941_v45_)
            d_942_v46_: _dafny.Set
            d_942_v46_ = _dafny.Set({default__.safeDivisionInt(p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahgwcjry"))))})
            d_942_v46_ = (d_942_v46_).intersection(_dafny.Set({d_941_v45_, d_941_v45_}))
            d_851_v0_ = _dafny.CodePoint('n')
        if (p2) != (p2):
            d_943_v47_: int
            d_943_v47_ = 92
            d_943_v47_ = p2
            if (self).f15:
                d_944_v48_: _dafny.Array
                nw144_ = _dafny.Array(int(0), 4)
                d_944_v48_ = nw144_
                index116_ = default__.safeIndex(813, (d_944_v48_).length(0))
                (d_944_v48_)[index116_] = p2
                index117_ = default__.safeIndex(813, (d_944_v48_).length(0))
                (d_944_v48_)[index117_] = (p2) - (default__.safeModuloInt(p2, len(p1)))
                d_945_v49_: _dafny.Array
                nw145_ = _dafny.Array(False, 3)
                d_945_v49_ = nw145_
                index118_ = default__.safeIndex(379, (d_945_v49_).length(0))
                (d_945_v49_)[index118_] = (self).f15
                d_946_v50_: _dafny.Set
                d_946_v50_ = _dafny.Set({d_944_v48_})
                d_947_v51_: _dafny.Map
                d_947_v51_ = _dafny.Map({len(d_946_v50_): p2})
                d_948_v52_: _dafny.Map
                d_948_v52_ = _dafny.Map({default__.fm2((self).f15, globalState): len(d_947_v51_)})
                d_949_v53_: _dafny.Seq
                d_949_v53_ = _dafny.SeqWithoutIsStrInference([d_948_v52_, (d_948_v52_).set((D0_DC2(len(p1), (self).f15)).cf2, p2)])
                d_950_v54_: _dafny.Map
                d_950_v54_ = _dafny.Map({d_944_v48_: _dafny.MultiSet(d_949_v53_)})
                d_951_v55_: _dafny.MultiSet
                d_951_v55_ = _dafny.MultiSet([d_948_v52_, _dafny.Map({(self).f15: 401}), d_948_v52_])
                index119_ = default__.safeIndex(379, (d_945_v49_).length(0))
                (d_945_v49_)[index119_] = not((((d_950_v54_)[d_944_v48_] if (d_944_v48_) in (d_950_v54_) else d_951_v55_)).isdisjoint(d_951_v55_))
                index120_ = default__.safeIndex(813, (d_944_v48_).length(0))
                (d_944_v48_)[index120_] = d_943_v47_
                d_952_v56_: D0
                d_952_v56_ = D0_DC0((self).f15)
                d_952_v56_ = d_952_v56_
                d_953_v57_: _dafny.Map
                d_953_v57_ = _dafny.Map({(d_945_v49_)[default__.safeIndex(379, (d_945_v49_).length(0))]: True})
                d_954_v58_: C1
                nw146_ = C1()
                nw146_.ctor__(False, not((d_945_v49_)[default__.safeIndex(379, (d_945_v49_).length(0))]), _dafny.Map({False: d_953_v57_}), _dafny.SeqWithoutIsStrInference([(d_945_v49_)[default__.safeIndex(379, (d_945_v49_).length(0))]]), self.f12)
                d_954_v58_ = nw146_
                d_955_v59_: C1
                nw147_ = C1()
                nw147_.ctor__((_dafny.Map({d_954_v58_: (d_944_v48_)[default__.safeIndex(813, (d_944_v48_).length(0))]})) == (_dafny.Map({d_954_v58_: d_943_v47_})), d_954_v58_.f18, (self).f13, (self).f14, self.f12)
                d_955_v59_ = nw147_
                d_955_v59_ = d_954_v58_
            elif True:
                rhs78_ = d_851_v0_
                rhs79_ = (self).f15
                lhs45_ = globalState
                d_851_v0_ = rhs78_
                lhs45_.f4 = rhs79_
                d_956_v60_: _dafny.Array
                nw148_ = _dafny.Array(False, 26)
                d_956_v60_ = nw148_
                index121_ = default__.safeIndex(9, (d_956_v60_).length(0))
                (d_956_v60_)[index121_] = False
                d_957_v61_: _dafny.Array
                nw149_ = _dafny.Array(int(0), 13)
                d_957_v61_ = nw149_
                index122_ = default__.safeIndex(892, (d_957_v61_).length(0))
                (d_957_v61_)[index122_] = -21
                d_958_v62_: _dafny.Seq
                d_958_v62_ = _dafny.SeqWithoutIsStrInference([d_957_v61_, d_957_v61_])
                d_959_v63_: _dafny.Map
                d_959_v63_ = _dafny.Map({(self).f15: (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(len(_dafny.Set({(self).f15})), len(p1))] for d_960_i9_ in range(default__.abs(595))]))))})
                index123_ = default__.safeIndex(9, (d_956_v60_).length(0))
                index124_ = default__.safeIndex(892, (d_957_v61_).length(0))
                rhs80_ = (self).f15
                rhs81_ = len((self).f14)
                rhs82_ = default__.safeDivisionInt((0) - (len((d_958_v62_) + (_dafny.SeqWithoutIsStrInference([d_957_v61_, d_957_v61_, d_957_v61_])))), p2)
                rhs83_ = (p2 if not(True) else (len(p1)) + (d_943_v47_))
                rhs84_ = len((d_959_v63_) | (d_959_v63_))
                lhs46_ = d_956_v60_
                lhs47_ = default__.safeIndex(9, (d_956_v60_).length(0))
                lhs48_ = d_957_v61_
                lhs49_ = default__.safeIndex(892, (d_957_v61_).length(0))
                lhs46_[lhs47_] = rhs80_
                lhs48_[lhs49_] = rhs81_
                d_943_v47_ = rhs82_
                d_943_v47_ = rhs83_
                d_943_v47_ = rhs84_
                d_961_v64_: _dafny.Array
                nw150_ = _dafny.Array(_dafny.Seq({}), 5)
                d_961_v64_ = nw150_
                d_962_v65_: D3
                d_962_v65_ = D3_DC8(_dafny.Map({(self).f15: (d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))]}))
                d_963_v66_: _dafny.Map
                d_963_v66_ = _dafny.Map({(self).f15: not((self).f15)})
                d_964_v67_: _dafny.Seq
                d_964_v67_ = _dafny.SeqWithoutIsStrInference([D3_DC8(d_963_v66_)])
                index125_ = default__.safeIndex(588, (d_961_v64_).length(0))
                (d_961_v64_)[index125_] = (_dafny.SeqWithoutIsStrInference([d_962_v65_, d_962_v65_])) + (d_964_v67_)
                d_965_v68_: _dafny.Set
                d_965_v68_ = _dafny.Set({(d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))]})
                index126_ = default__.safeIndex(588, (d_961_v64_).length(0))
                (d_961_v64_)[index126_] = (_dafny.SeqWithoutIsStrInference([d_962_v65_, d_962_v65_])).set(default__.safeIndex(len(d_965_v68_), len(_dafny.SeqWithoutIsStrInference([d_962_v65_, d_962_v65_]))), d_962_v65_)
                d_966_v69_: D6
                d_966_v69_ = D6_DC18((d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))])
                d_967_v70_: D5
                d_967_v70_ = D5_DC16((self).f15)
                d_968_v71_: _dafny.Map
                d_968_v71_ = _dafny.Map({d_967_v70_: (_dafny.MultiSet([len((self).f16)])).cardinality})
                d_969_v72_: D2
                d_969_v72_ = D2_DC5(((d_968_v71_)[D5_DC16((d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))])] if (D5_DC16((d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))])) in (d_968_v71_) else (d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))]), (d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))], ((self).f14)[default__.safeIndex(d_943_v47_, len((self).f14))])
                d_970_v73_: _dafny.Map
                d_970_v73_ = _dafny.Map({d_966_v69_: d_969_v72_})
                d_971_v75_: _dafny.MultiSet
                d_971_v75_ = _dafny.MultiSet([D6_DC18(d_943_v47_), d_966_v69_, d_966_v69_])
                d_972_v76_: _dafny.Array
                nw151_ = _dafny.Array(None, 5)
                nw151_[int(0)] = (d_970_v73_).set(d_966_v69_, d_969_v72_)
                def iife66_():
                    coll46_ = _dafny.Map()
                    compr_46_: D6
                    for compr_46_ in (((d_971_v75_).set(D6_DC18((d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))]), default__.abs(p2))).set(d_966_v69_, default__.abs(p2))).Elements:
                        d_973_v74_: D6 = compr_46_
                        if (d_973_v74_) in (((d_971_v75_).set(D6_DC18((d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))]), default__.abs(p2))).set(d_966_v69_, default__.abs(p2))):
                            coll46_[d_973_v74_] = d_969_v72_
                    return _dafny.Map(coll46_)
                nw151_[int(1)] = iife66_()
                
                nw151_[int(2)] = _dafny.Map({d_966_v69_: d_969_v72_})
                nw151_[int(3)] = d_970_v73_
                nw151_[int(4)] = d_970_v73_
                d_972_v76_ = nw151_
                index127_ = default__.safeIndex(298, (d_972_v76_).length(0))
                (d_972_v76_)[index127_] = d_970_v73_
                d_974_v77_: _dafny.Map
                d_974_v77_ = _dafny.Map({d_943_v47_: p2})
                index128_ = default__.safeIndex(298, (d_972_v76_).length(0))
                index129_ = default__.safeIndex(892, (d_957_v61_).length(0))
                rhs85_ = d_970_v73_
                rhs86_ = (len(p1)) >= (default__.safeModuloInt(len(d_974_v77_), d_943_v47_))
                rhs87_ = p2
                rhs88_ = ((self).f14)[default__.safeIndex((d_957_v61_)[default__.safeIndex(892, (d_957_v61_).length(0))], len((self).f14))]
                rhs89_ = (d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))]
                lhs50_ = d_972_v76_
                lhs51_ = default__.safeIndex(298, (d_972_v76_).length(0))
                lhs52_ = d_957_v61_
                lhs53_ = default__.safeIndex(892, (d_957_v61_).length(0))
                lhs54_ = globalState
                lhs55_ = globalState
                lhs50_[lhs51_] = rhs85_
                r0 = rhs86_
                lhs52_[lhs53_] = rhs87_
                lhs54_.f4 = rhs88_
                lhs55_.f4 = rhs89_
                d_975_v78_: C1
                nw152_ = C1()
                nw152_.ctor__(((d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))]) and ((d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))]), not((d_956_v60_)[default__.safeIndex(9, (d_956_v60_).length(0))]), (self).f13, (self).f14, self.f12)
                d_975_v78_ = nw152_
                d_975_v78_ = d_975_v78_
            d_976_v79_: C0
            nw153_ = C0()
            nw153_.ctor__()
            d_976_v79_ = nw153_
            d_977_v80_: _dafny.Seq
            d_977_v80_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f12, self.f12])
            d_978_v81_: C1
            nw154_ = C1()
            nw154_.ctor__((self).f15, (self).f15, (self).f13, ((self).f14) + ((self).f14), (d_977_v80_)[default__.safeIndex(len(p1), len(d_977_v80_))])
            d_978_v81_ = nw154_
            (globalState).f4 = d_978_v81_.f18
        elif True:
            (globalState).f4 = (self).f15
            d_979_v82_: _dafny.Array
            def lambda55_(d_980_p2_):
                def lambda56_(d_981_i10_):
                    return default__.safeDivisionInt(d_981_i10_, d_980_p2_)

                return lambda56_

            init28_ = lambda55_(p2)
            nw155_ = _dafny.Array(None, 16)
            for i0_28_ in range(nw155_.length(0)):
                nw155_[i0_28_] = init28_(i0_28_)
            d_979_v82_ = nw155_
            d_982_v83_: _dafny.MultiSet
            d_982_v83_ = d_856_v5_
            d_983_v84_: _dafny.Map
            d_983_v84_ = _dafny.Map({d_979_v82_: d_982_v83_})
            index130_ = default__.safeIndex(811, (d_979_v82_).length(0))
            (d_979_v82_)[index130_] = len(d_983_v84_)
            d_984_v85_: D7
            d_984_v85_ = D7_DC22((self).f15)
            d_985_v86_: _dafny.Map
            def iife67_(_pat_let10_0):
                def iife68_(d_986_dt__update__tmp_h2_):
                    def iife69_(_pat_let11_0):
                        def iife70_(d_987_dt__update_hcf34_h0_):
                            return D7_DC22(d_987_dt__update_hcf34_h0_)
                        return iife70_(_pat_let11_0)
                    return iife69_((self).f15)
                return iife68_(_pat_let10_0)
            d_985_v86_ = _dafny.Map({iife67_(d_984_v85_): -494})
            index131_ = default__.safeIndex(811, (d_979_v82_).length(0))
            (d_979_v82_)[index131_] = ((d_985_v86_)[D7_DC22(((self).f14)[default__.safeIndex(p2, len((self).f14))])] if (D7_DC22(((self).f14)[default__.safeIndex(p2, len((self).f14))])) in (d_985_v86_) else p2)
            d_988_v87_: _dafny.Array
            nw156_ = _dafny.Array(False, 16)
            d_988_v87_ = nw156_
            d_989_v88_: _dafny.Seq
            d_989_v88_ = _dafny.SeqWithoutIsStrInference([d_988_v87_, d_988_v87_, d_988_v87_])
            d_989_v88_ = (d_989_v88_).set(default__.safeIndex((0) - (p2), len(d_989_v88_)), d_988_v87_)
            d_990_v89_: C1
            nw157_ = C1()
            nw157_.ctor__(((d_854_v3_)[p2] if (p2) in (d_854_v3_) else default__.fm2((self).f15, globalState)), (self).f15, (self).f13, (self).f14, self.f12)
            d_990_v89_ = nw157_
            index132_ = default__.safeIndex(811, (d_979_v82_).length(0))
            (d_979_v82_)[index132_] = (d_979_v82_)[default__.safeIndex(811, (d_979_v82_).length(0))]
        d_991_i11_: int
        d_991_i11_ = 0
        with _dafny.label("8"):
            while (self).f15:
                with _dafny.c_label("8"):
                    if (d_991_i11_) >= (100):
                        raise _dafny.Break("8")
                    d_991_i11_ = (d_991_i11_) + (1)
                    d_992_v90_: int
                    d_992_v90_ = 683
                    d_993_v91_: _dafny.Seq
                    d_993_v91_ = _dafny.SeqWithoutIsStrInference([(self).f15, not((self).f15), (self).f15, (len((_dafny.Map({(self).f15: (self).f15})).set((self).f15, (self).f15))) == (d_992_v90_), (self).f15])
                    d_994_v92_: _dafny.Array
                    nw158_ = _dafny.Array(int(0), 26)
                    d_994_v92_ = nw158_
                    rhs90_ = 109
                    rhs91_ = (self).f14
                    rhs92_ = p2
                    rhs93_ = d_994_v92_
                    d_992_v90_ = rhs90_
                    d_993_v91_ = rhs91_
                    d_992_v90_ = rhs92_
                    d_994_v92_ = rhs93_
                    d_995_v93_: D6
                    d_995_v93_ = D6_DC18(d_992_v90_)
                    d_996_v94_: _dafny.Set
                    d_996_v94_ = _dafny.Set({d_854_v3_})
                    d_995_v93_ = default__.fm46(p1, (self).f15, d_996_v94_, globalState)
                    d_992_v90_ = d_992_v90_
                    d_997_v95_: _dafny.Array
                    nw159_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                    d_997_v95_ = nw159_
                    d_998_v96_: _dafny.Map
                    d_998_v96_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "naxhcjn"))})
                    d_999_v97_: _dafny.Map
                    d_999_v97_ = _dafny.Map({(self).f15: (self).f15})
                    d_1000_v98_: _dafny.Map
                    d_1000_v98_ = _dafny.Map({default__.fm2((self).f15, globalState): p2})
                    index133_ = default__.safeIndex(757, (d_997_v95_).length(0))
                    (d_997_v95_)[index133_] = ((d_998_v96_)[default__.fm0((self).f15, _dafny.Map({len(d_999_v97_): d_1000_v98_}), (self).f15, globalState)] if (default__.fm0((self).f15, _dafny.Map({len(d_999_v97_): d_1000_v98_}), (self).f15, globalState)) in (d_998_v96_) else p1)
                    index134_ = default__.safeIndex(757, (d_997_v95_).length(0))
                    (d_997_v95_)[index134_] = p1
                    pass
            pass
        d_1001_v99_: C0
        nw160_ = C0()
        nw160_.ctor__()
        d_1001_v99_ = nw160_
        r0 = (self).f15
        return r0

    def m6(self, globalState):
        r0: bool = False
        d_1002_v0_: int
        d_1002_v0_ = 58
        d_1003_v1_: _dafny.Set
        d_1003_v1_ = _dafny.Set({d_1002_v0_})
        d_1004_v2_: _dafny.Map
        d_1004_v2_ = _dafny.Map({True: d_1002_v0_})
        d_1005_v3_: _dafny.Map
        d_1005_v3_ = _dafny.Map({len(d_1003_v1_): d_1004_v2_})
        d_1006_v4_: D11
        d_1006_v4_ = D11_DC30(d_1005_v3_)
        d_1007_v5_: _dafny.Set
        d_1007_v5_ = _dafny.Set({(self).f15, (self).f15, (self).f15})
        d_1008_v6_: _dafny.Array
        nw161_ = _dafny.Array(None, 26)
        nw161_[int(0)] = d_1002_v0_
        nw161_[int(1)] = d_1002_v0_
        nw161_[int(2)] = d_1002_v0_
        nw161_[int(3)] = d_1002_v0_
        nw161_[int(4)] = (0) - ((0) - (d_1002_v0_))
        nw161_[int(5)] = len((self).f14)
        nw161_[int(6)] = (0) - (d_1002_v0_)
        nw161_[int(7)] = d_1002_v0_
        nw161_[int(8)] = d_1002_v0_
        nw161_[int(9)] = d_1002_v0_
        nw161_[int(10)] = d_1002_v0_
        nw161_[int(11)] = d_1002_v0_
        nw161_[int(12)] = len((self).f16)
        nw161_[int(13)] = d_1002_v0_
        nw161_[int(14)] = d_1002_v0_
        nw161_[int(15)] = d_1002_v0_
        nw161_[int(16)] = default__.fm0(default__.fm2((self).f15, globalState), (d_1006_v4_).cf45, (self).f15, globalState)
        nw161_[int(17)] = (0) - (d_1002_v0_)
        nw161_[int(18)] = d_1002_v0_
        nw161_[int(19)] = 522
        nw161_[int(20)] = d_1002_v0_
        nw161_[int(21)] = len(d_1007_v5_)
        nw161_[int(22)] = d_1002_v0_
        nw161_[int(23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjmqcatm")))
        nw161_[int(24)] = -426
        nw161_[int(25)] = d_1002_v0_
        d_1008_v6_ = nw161_
        d_1009_v7_: _dafny.Seq
        d_1009_v7_ = _dafny.SeqWithoutIsStrInference([d_1008_v6_, d_1008_v6_])
        d_1010_v8_: _dafny.Array
        nw162_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_1010_v8_ = nw162_
        d_1011_v9_: _dafny.Array
        nw163_ = _dafny.Array(False, 8)
        d_1011_v9_ = nw163_
        index135_ = default__.safeIndex(820, (d_1010_v8_).length(0))
        (d_1010_v8_)[index135_] = d_1011_v9_
        d_1012_v10_: D0
        d_1012_v10_ = D0_DC1()
        pat_let_tv11_ = d_1007_v5_
        pat_let_tv12_ = d_1002_v0_
        pat_let_tv13_ = d_1002_v0_
        index136_ = default__.safeIndex(820, (d_1010_v8_).length(0))
        def lambda57_(source15_):
            if source15_.is_DC1:
                return _dafny.SeqWithoutIsStrInference([len(pat_let_tv11_) for d_1013_i0_ in range(default__.abs(767))])
            elif source15_.is_DC2:
                d_1014___mcc_h0_ = source15_.cf1
                d_1015___mcc_h1_ = source15_.cf2
                d_1016_cf2_ = d_1015___mcc_h1_
                d_1017_cf1_ = d_1014___mcc_h0_
                return (((self).f16).set(default__.safeIndex(pat_let_tv12_, len((self).f16)), d_1017_cf1_)) + (_dafny.SeqWithoutIsStrInference([d_1017_cf1_, pat_let_tv13_]))
            elif True:
                d_1018___mcc_h2_ = source15_.cf0
                d_1019_cf0_ = d_1018___mcc_h2_
                return (self).f16

        rhs94_ = d_1009_v7_
        rhs95_ = d_1011_v9_
        rhs96_ = len(lambda57_(d_1012_v10_))
        rhs97_ = default__.fm2((self).f15, globalState)
        lhs56_ = d_1010_v8_
        lhs57_ = default__.safeIndex(820, (d_1010_v8_).length(0))
        lhs58_ = globalState
        d_1009_v7_ = rhs94_
        lhs56_[lhs57_] = rhs95_
        d_1002_v0_ = rhs96_
        lhs58_.f4 = rhs97_
        rhs98_ = (self).f15
        rhs99_ = ((self).f16)[default__.safeIndex(len(d_1007_v5_), len((self).f16))]
        lhs59_ = globalState
        lhs59_.f4 = rhs98_
        d_1002_v0_ = rhs99_
        d_1020_i1_: int
        d_1020_i1_ = 0
        with _dafny.label("9"):
            while (self).f15:
                with _dafny.c_label("9"):
                    if (d_1020_i1_) >= (100):
                        raise _dafny.Break("9")
                    d_1020_i1_ = (d_1020_i1_) + (1)
                    d_1021_v11_: _dafny.MultiSet
                    d_1021_v11_ = _dafny.MultiSet([(self).f15, (self).f15])
                    d_1022_v12_: _dafny.Map
                    d_1022_v12_ = _dafny.Map({d_1002_v0_: (self).fm38(233, (d_1021_v11_).cardinality, d_1002_v0_, (self).f15, globalState)})
                    d_1022_v12_ = (d_1022_v12_).set(d_1002_v0_, d_1002_v0_)
                    d_1023_v13_: D2
                    d_1023_v13_ = D2_DC6((d_1010_v8_)[default__.safeIndex(820, (d_1010_v8_).length(0))], d_1011_v9_)
                    source16_ = d_1023_v13_
                    if source16_.is_DC5:
                        d_1024___mcc_h3_ = source16_.cf5
                        d_1025___mcc_h4_ = source16_.cf6
                        d_1026___mcc_h5_ = source16_.cf7
                        d_1027_cf7_ = d_1026___mcc_h5_
                        d_1028_cf6_ = d_1025___mcc_h4_
                        d_1029_cf5_ = d_1024___mcc_h3_
                        d_1030_v14_: C0
                        nw164_ = C0()
                        nw164_.ctor__()
                        d_1030_v14_ = nw164_
                        d_1031_v15_: _dafny.Array
                        nw165_ = _dafny.Array(_dafny.Map({}), 17)
                        d_1031_v15_ = nw165_
                        d_1032_v16_: _dafny.Map
                        d_1032_v16_ = _dafny.Map({(self).f15: (self).f15})
                        index137_ = default__.safeIndex(20, (d_1031_v15_).length(0))
                        (d_1031_v15_)[index137_] = d_1032_v16_
                        index138_ = default__.safeIndex(20, (d_1031_v15_).length(0))
                        (d_1031_v15_)[index138_] = (((self).f13)[d_1027_cf7_] if (d_1027_cf7_) in ((self).f13) else d_1032_v16_)
                        d_1033_v17_: _dafny.Array
                        nw166_ = _dafny.Array(_dafny.Seq({}), 16)
                        d_1033_v17_ = nw166_
                        d_1034_v18_: _dafny.MultiSet
                        d_1034_v18_ = _dafny.MultiSet([d_1002_v0_])
                        d_1035_v19_: _dafny.Seq
                        d_1035_v19_ = _dafny.SeqWithoutIsStrInference([D7_DC20(d_1034_v18_), D7_DC20(d_1034_v18_)])
                        d_1036_v20_: _dafny.Seq
                        d_1036_v20_ = _dafny.SeqWithoutIsStrInference([d_1035_v19_, d_1035_v19_, (d_1035_v19_).set(default__.safeIndex(d_1002_v0_, len(d_1035_v19_)), D7_DC20(d_1034_v18_)), d_1035_v19_, d_1035_v19_])
                        index139_ = default__.safeIndex(224, (d_1033_v17_).length(0))
                        (d_1033_v17_)[index139_] = (d_1036_v20_) + (d_1036_v20_)
                        index140_ = default__.safeIndex(224, (d_1033_v17_).length(0))
                        (d_1033_v17_)[index140_] = d_1036_v20_
                        d_1037_v21_: _dafny.Array
                        def lambda58_(d_1038_v11_):
                            def lambda59_(d_1039_i2_):
                                return d_1038_v11_

                            return lambda59_

                        init29_ = lambda58_(d_1021_v11_)
                        nw167_ = _dafny.Array(None, 11)
                        for i0_29_ in range(nw167_.length(0)):
                            nw167_[i0_29_] = init29_(i0_29_)
                        d_1037_v21_ = nw167_
                        index141_ = default__.safeIndex(734, (d_1037_v21_).length(0))
                        (d_1037_v21_)[index141_] = d_1021_v11_
                        index142_ = default__.safeIndex(734, (d_1037_v21_).length(0))
                        (d_1037_v21_)[index142_] = ((d_1021_v11_).set(True, default__.abs(-767))) | (d_1021_v11_)
                    elif source16_.is_DC6:
                        d_1040___mcc_h6_ = source16_.cf8
                        d_1041___mcc_h7_ = source16_.cf9
                        d_1042_cf9_ = d_1041___mcc_h7_
                        d_1043_cf8_ = d_1040___mcc_h6_
                        d_1044_v22_: _dafny.Map
                        d_1044_v22_ = _dafny.Map({True: (self).f15})
                        d_1045_v23_: _dafny.Map
                        d_1045_v23_ = _dafny.Map({d_1044_v22_: True})
                        d_1045_v23_ = (d_1045_v23_).set(_dafny.Map({(self).f15: (self).f15}), (self).f15)
                        d_1046_v24_: _dafny.Seq
                        d_1046_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrgw"))
                        d_1002_v0_ = len((d_1046_v24_) + ((d_1046_v24_) + (d_1046_v24_)))
                        d_1046_v24_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvucslniw"))) + (d_1046_v24_)) + ((d_1046_v24_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjffmj"))))
                        d_1047_v25_: D8
                        d_1047_v25_ = D8_DC26((self).f15)
                        d_1048_v26_: _dafny.Map
                        d_1048_v26_ = _dafny.Map({d_1043_cf8_: d_1047_v25_})
                        d_1049_v27_: D12
                        d_1049_v27_ = D12_DC34(d_1048_v26_)
                        d_1048_v26_ = (d_1049_v27_).cf53
                    elif source16_.is_DC4:
                        d_1050___mcc_h8_ = source16_.cf4
                        d_1051_cf4_ = d_1050___mcc_h8_
                        (globalState).f4 = (self).f15
                        (globalState).f4 = (self).f15
                        d_1052_v28_: _dafny.Array
                        def lambda60_(d_1053_i3_):
                            return _dafny.CodePoint('s')

                        init30_ = lambda60_
                        nw168_ = _dafny.Array(None, 28)
                        for i0_30_ in range(nw168_.length(0)):
                            nw168_[i0_30_] = init30_(i0_30_)
                        d_1052_v28_ = nw168_
                        index143_ = default__.safeIndex(476, (d_1052_v28_).length(0))
                        (d_1052_v28_)[index143_] = _dafny.CodePoint('y')
                        d_1054_v29_: str
                        d_1054_v29_ = _dafny.CodePoint('a')
                        index144_ = default__.safeIndex(476, (d_1052_v28_).length(0))
                        (d_1052_v28_)[index144_] = d_1054_v29_
                        d_1055_v30_: _dafny.Seq
                        d_1055_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqykhwa"))
                        d_1002_v0_ = len((d_1055_v30_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpcuoqb"))))
                    elif True:
                        d_1056___mcc_h9_ = source16_.cf10
                        d_1057_cf10_ = d_1056___mcc_h9_
                        index145_ = default__.safeIndex(370, (d_1008_v6_).length(0))
                        (d_1008_v6_)[index145_] = d_1002_v0_
                        index146_ = default__.safeIndex(370, (d_1008_v6_).length(0))
                        (d_1008_v6_)[index146_] = (0) - (d_1002_v0_)
                        (globalState).f4 = (self).f15
                        (globalState).f4 = (d_1002_v0_) != ((d_1008_v6_)[default__.safeIndex(370, (d_1008_v6_).length(0))])
                        d_1058_v31_: D2
                        d_1058_v31_ = D2_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vljqrh"))), (self).f15, (self).f15)
                        index147_ = default__.safeIndex(370, (d_1008_v6_).length(0))
                        (d_1008_v6_)[index147_] = (d_1002_v0_ if True else (d_1058_v31_).cf5)
                    d_1059_v32_: _dafny.Array
                    def lambda61_(d_1060_i4_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1061_i5_ in range(default__.abs(-301))])

                    init31_ = lambda61_
                    nw169_ = _dafny.Array(None, 11)
                    for i0_31_ in range(nw169_.length(0)):
                        nw169_[i0_31_] = init31_(i0_31_)
                    d_1059_v32_ = nw169_
                    d_1062_v33_: D12
                    d_1062_v33_ = D12_DC36(d_1059_v32_, d_1002_v0_, default__.fm47(globalState), _dafny.Set({d_1011_v9_, d_1011_v9_}))
                    index148_ = default__.safeIndex(644, (d_1011_v9_).length(0))
                    (d_1011_v9_)[index148_] = default__.fm2(not((self).f15), globalState)
                    d_1063_v34_: _dafny.Seq
                    d_1063_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npgudenrk"))
                    index149_ = default__.safeIndex(644, (d_1011_v9_).length(0))
                    rhs100_ = (self).f15
                    rhs101_ = d_1062_v33_
                    rhs102_ = (d_1063_v34_) < ((d_1063_v34_ if (self).f15 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eghjdwmk"))))
                    rhs103_ = ((self).f15) or ((self).f15)
                    lhs60_ = globalState
                    lhs61_ = globalState
                    lhs62_ = d_1011_v9_
                    lhs63_ = default__.safeIndex(644, (d_1011_v9_).length(0))
                    lhs60_.f4 = rhs100_
                    d_1062_v33_ = rhs101_
                    lhs61_.f4 = rhs102_
                    lhs62_[lhs63_] = rhs103_
                    d_1002_v0_ = d_1002_v0_
                    pass
            pass
        d_1002_v0_ = d_1002_v0_
        pat_let_tv14_ = d_1005_v3_
        pat_let_tv15_ = d_1005_v3_
        d_1064_v35_: _dafny.Array
        nw170_ = _dafny.Array(None, 15)
        nw170_[int(0)] = d_1006_v4_
        nw170_[int(1)] = d_1006_v4_
        nw170_[int(2)] = (d_1006_v4_ if True else d_1006_v4_)
        nw170_[int(3)] = d_1006_v4_
        nw170_[int(4)] = d_1006_v4_
        nw170_[int(5)] = d_1006_v4_
        nw170_[int(6)] = d_1006_v4_
        nw170_[int(7)] = d_1006_v4_
        def iife72_(_pat_let13_0):
            def iife73_(d_1065_dt__update__tmp_h0_):
                def iife74_(_pat_let14_0):
                    def iife75_(d_1066_dt__update_hcf45_h0_):
                        return D11_DC30(d_1066_dt__update_hcf45_h0_)
                    return iife75_(_pat_let14_0)
                return iife74_(pat_let_tv14_)
            return iife73_(_pat_let13_0)
        def iife71_(_pat_let12_0):
            def iife76_(d_1067_dt__update__tmp_h1_):
                def iife77_(_pat_let15_0):
                    def iife78_(d_1068_dt__update_hcf45_h1_):
                        return D11_DC30(d_1068_dt__update_hcf45_h1_)
                    return iife78_(_pat_let15_0)
                return iife77_(pat_let_tv15_)
            return iife76_(_pat_let12_0)
        nw170_[int(8)] = iife71_(iife72_(d_1006_v4_))
        nw170_[int(9)] = d_1006_v4_
        nw170_[int(10)] = d_1006_v4_
        nw170_[int(11)] = D11_DC30(_dafny.Map({d_1002_v0_: d_1004_v2_}))
        nw170_[int(12)] = default__.fm48(globalState)
        nw170_[int(13)] = D11_DC30(d_1005_v3_)
        nw170_[int(14)] = default__.fm48(globalState)
        d_1064_v35_ = nw170_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1064_v35_).length(0)):
            d_1069_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_1069_i6_)) and ((d_1069_i6_) < ((d_1064_v35_).length(0)))):
                (d_1064_v35_)[(d_1069_i6_)] = d_1006_v4_
        if (self).f15:
            d_1004_v2_ = (d_1004_v2_) | ((_dafny.Map({not((self).f15): d_1002_v0_})) | (d_1004_v2_))
            d_1070_v36_: C1
            nw171_ = C1()
            nw171_.ctor__((self).f15, ((self).f16) == (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1002_v0_: d_1002_v0_})) for d_1071_i7_ in range(default__.abs(202))])), (self).f13, ((self).f14).set(default__.safeIndex(d_1002_v0_, len((self).f14)), (self).f15), self.f12)
            d_1070_v36_ = nw171_
            index150_ = default__.safeIndex(820, (d_1010_v8_).length(0))
            (d_1010_v8_)[index150_] = d_1011_v9_
            d_1072_v37_: D11
            d_1072_v37_ = D11_DC31()
            d_1072_v37_ = D11_DC31()
            d_1002_v0_ = (0) - (-572)
        elif True:
            d_1073_v38_: D6
            d_1073_v38_ = D6_DC18(d_1002_v0_)
            d_1074_v39_: _dafny.Seq
            d_1074_v39_ = _dafny.SeqWithoutIsStrInference([(d_1002_v0_) + (len((self).f14))])
            rhs104_ = d_1073_v38_
            rhs105_ = (((self).f16).set(default__.safeIndex(-437, len((self).f16)), d_1002_v0_)) + ((self).f16)
            d_1073_v38_ = rhs104_
            d_1074_v39_ = rhs105_
            d_1002_v0_ = 709
            d_1075_v40_: _dafny.Seq
            d_1075_v40_ = _dafny.SeqWithoutIsStrInference([d_1011_v9_, d_1011_v9_])
            d_1076_v41_: _dafny.MultiSet
            d_1076_v41_ = _dafny.MultiSet([(d_1075_v40_)[default__.safeIndex(d_1002_v0_, len(d_1075_v40_))], d_1011_v9_])
            d_1077_v42_: str
            d_1077_v42_ = _dafny.CodePoint('b')
            d_1078_v43_: _dafny.Map
            d_1078_v43_ = _dafny.Map({(d_1076_v41_) | (d_1076_v41_): d_1077_v42_})
            d_1078_v43_ = (d_1078_v43_).set((d_1076_v41_) | (d_1076_v41_), d_1077_v42_)
            d_1079_v44_: _dafny.Array
            nw172_ = _dafny.Array(None, 15)
            nw172_[int(0)] = d_1008_v6_
            nw172_[int(1)] = ((d_1009_v7_)[default__.safeIndex(d_1002_v0_, len(d_1009_v7_))] if default__.fm2((self).f15, globalState) else d_1008_v6_)
            nw172_[int(2)] = d_1008_v6_
            nw172_[int(3)] = d_1008_v6_
            nw172_[int(4)] = d_1008_v6_
            nw172_[int(5)] = d_1008_v6_
            nw172_[int(6)] = d_1008_v6_
            nw172_[int(7)] = d_1008_v6_
            nw172_[int(8)] = d_1008_v6_
            nw172_[int(9)] = d_1008_v6_
            nw172_[int(10)] = d_1008_v6_
            nw172_[int(11)] = d_1008_v6_
            nw172_[int(12)] = d_1008_v6_
            nw172_[int(13)] = (d_1009_v7_)[default__.safeIndex(d_1002_v0_, len(d_1009_v7_))]
            nw172_[int(14)] = d_1008_v6_
            d_1079_v44_ = nw172_
            d_1079_v44_ = (d_1079_v44_ if (self).f15 else d_1079_v44_)
            d_1080_v45_: D5
            d_1080_v45_ = D5_DC16((self).f15)
            source17_ = d_1080_v45_
            if source17_.is_DC16:
                d_1081___mcc_h10_ = source17_.cf26
                d_1082_cf26_ = d_1081___mcc_h10_
                d_1002_v0_ = (0) - (d_1002_v0_)
                d_1083_v46_: D5
                d_1083_v46_ = D5_DC15((self).f16)
                index151_ = default__.safeIndex(837, (d_1011_v9_).length(0))
                (d_1011_v9_)[index151_] = d_1082_cf26_
                d_1084_v47_: _dafny.MultiSet
                d_1084_v47_ = _dafny.MultiSet([(self).f16, _dafny.SeqWithoutIsStrInference([d_1002_v0_]), ((self).f16).set(default__.safeIndex(d_1002_v0_, len((self).f16)), d_1002_v0_), (self).f16])
                d_1085_v48_: _dafny.Seq
                d_1085_v48_ = _dafny.SeqWithoutIsStrInference([d_1084_v47_, d_1084_v47_])
                index152_ = default__.safeIndex(837, (d_1011_v9_).length(0))
                rhs106_ = d_1083_v46_
                rhs107_ = ((d_1084_v47_).intersection(d_1084_v47_)).issubset((d_1084_v47_) | ((d_1085_v48_)[default__.safeIndex(149, len(d_1085_v48_))]))
                rhs108_ = (((self).f15) == (d_1082_cf26_)) and ((self).f15)
                rhs109_ = d_1073_v38_
                rhs110_ = (not((self).f15) if d_1082_cf26_ else not((self).f15))
                lhs64_ = globalState
                lhs65_ = globalState
                lhs66_ = d_1011_v9_
                lhs67_ = default__.safeIndex(837, (d_1011_v9_).length(0))
                d_1083_v46_ = rhs106_
                lhs64_.f4 = rhs107_
                lhs65_.f4 = rhs108_
                d_1073_v38_ = rhs109_
                lhs66_[lhs67_] = rhs110_
                (globalState).f4 = d_1082_cf26_
                d_1086_v49_: _dafny.Seq
                d_1086_v49_ = _dafny.SeqWithoutIsStrInference([(d_1011_v9_)[default__.safeIndex(837, (d_1011_v9_).length(0))], (d_1011_v9_)[default__.safeIndex(837, (d_1011_v9_).length(0))], (d_1011_v9_)[default__.safeIndex(837, (d_1011_v9_).length(0))]])
                d_1086_v49_ = (_dafny.SeqWithoutIsStrInference([(d_1011_v9_)[default__.safeIndex(837, (d_1011_v9_).length(0))], d_1082_cf26_, True, not((d_1011_v9_)[default__.safeIndex(837, (d_1011_v9_).length(0))]), default__.fm2((d_1086_v49_)[default__.safeIndex(-489, len(d_1086_v49_))], globalState)])) + (d_1086_v49_)
            elif True:
                d_1087___mcc_h11_ = source17_.cf25
                d_1088_cf25_ = d_1087___mcc_h11_
                d_1002_v0_ = len((self).f14)
                d_1089_v50_: _dafny.Map
                d_1089_v50_ = _dafny.Map({(self).f15: d_1011_v9_})
                d_1089_v50_ = (d_1089_v50_).set((d_1003_v1_).ispropersubset(d_1003_v1_), (d_1010_v8_)[default__.safeIndex(820, (d_1010_v8_).length(0))])
                d_1090_v51_: _dafny.Seq
                d_1090_v51_ = _dafny.SeqWithoutIsStrInference([d_1077_v42_])
                d_1091_v52_: D8
                d_1091_v52_ = D8_DC25((d_1007_v5_).issubset(d_1007_v5_), (len(d_1090_v51_) if (self).f15 else d_1002_v0_), d_1002_v0_, (d_1004_v2_) | (d_1004_v2_))
                d_1091_v52_ = d_1091_v52_
                d_1092_v53_: str
                d_1093_v54_: bool
                out87_: str
                out88_: bool
                out87_, out88_ = (self).m11(globalState)
                d_1092_v53_ = out87_
                d_1093_v54_ = out88_
        d_1094_v55_: T1
        nw173_ = C1()
        nw173_.ctor__((self).f15, True, (self).f13, (self).f14, self.f12)
        d_1094_v55_ = nw173_
        d_1095_v56_: D3
        d_1095_v56_ = D3_DC10(not((self).f15), d_1002_v0_, d_1094_v55_)
        r0 = not((d_1095_v56_).cf14)
        return r0

    def m11(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: bool = False
        d_1096_v0_: _dafny.Seq
        d_1096_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyjxdokyt"))
        d_1097_v1_: _dafny.MultiSet
        d_1097_v1_ = _dafny.MultiSet([d_1096_v0_, d_1096_v0_, d_1096_v0_, d_1096_v0_, d_1096_v0_])
        d_1098_v2_: int
        d_1098_v2_ = 132
        r1 = ((default__.fm2((self).f15, globalState)) not in ((self).f14)) and ((_dafny.MultiSet([d_1096_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpqincvt")), d_1096_v0_, d_1096_v0_, default__.fm40(D4_DC14((self).f15), d_1098_v2_, False, (self).f15, globalState)])).issubset(d_1097_v1_))
        d_1099_v3_: D0
        d_1099_v3_ = D0_DC2(len(d_1096_v0_), (self).f15)
        pat_let_tv16_ = d_1097_v1_
        def lambda62_(source19_):
            if source19_.is_DC1:
                return D8_DC23(_dafny.Map({(self).f15: (pat_let_tv16_).cardinality}))
            elif source19_.is_DC2:
                d_1100___mcc_h8_ = source19_.cf1
                d_1101___mcc_h9_ = source19_.cf2
                d_1102_cf2_ = d_1101___mcc_h9_
                d_1103_cf1_ = d_1100___mcc_h8_
                return D8_DC23(_dafny.Map({not(True): d_1103_cf1_}))
            elif True:
                d_1104___mcc_h10_ = source19_.cf0
                d_1105_cf0_ = d_1104___mcc_h10_
                return D8_DC23(_dafny.Map({d_1105_cf0_: 117}))

        source18_ = lambda62_(d_1099_v3_)
        if source18_.is_DC24:
            d_1106___mcc_h0_ = source18_.cf36
            d_1107_cf36_ = d_1106___mcc_h0_
            d_1108_v4_: _dafny.Set
            d_1108_v4_ = _dafny.Set({d_1107_cf36_, (self).f15, True})
            d_1109_v5_: _dafny.Map
            d_1109_v5_ = _dafny.Map({len((d_1108_v4_) | (d_1108_v4_)): d_1107_cf36_})
            d_1110_v6_: str
            d_1110_v6_ = _dafny.CodePoint('f')
            d_1111_v7_: _dafny.Map
            d_1111_v7_ = _dafny.Map({d_1110_v6_: ((d_1109_v5_)[d_1098_v2_] if (d_1098_v2_) in (d_1109_v5_) else d_1107_cf36_)})
            d_1109_v5_ = (d_1109_v5_).set(len(d_1111_v7_), (self).f15)
            d_1112_v8_: _dafny.MultiSet
            d_1112_v8_ = _dafny.MultiSet([d_1107_cf36_, False])
            (globalState).f4 = ((d_1112_v8_).cardinality) >= (((self).fm38(d_1098_v2_, d_1098_v2_, (0) - (d_1098_v2_), (self).f15, globalState)) - (d_1098_v2_))
            d_1113_v9_: _dafny.Array
            nw174_ = _dafny.Array(_dafny.Seq({}), 22)
            d_1113_v9_ = nw174_
            d_1114_v10_: D4
            d_1114_v10_ = D4_DC12(_dafny.SeqWithoutIsStrInference([(self).f15, d_1107_cf36_]))
            d_1115_v11_: _dafny.Map
            d_1115_v11_ = _dafny.Map({d_1107_cf36_: (self).f14})
            nw175_ = _dafny.Array(None, 24)
            nw175_[int(0)] = ((self).f14) + ((self).f14)
            nw175_[int(1)] = (self).f14
            nw175_[int(2)] = (self).f14
            nw175_[int(3)] = (self).f14
            nw175_[int(4)] = (d_1114_v10_).cf18
            nw175_[int(5)] = ((self).f14) + ((self).f14)
            nw175_[int(6)] = (self).f14
            nw175_[int(7)] = (self).f14
            nw175_[int(8)] = (self).f14
            nw175_[int(9)] = (self).f14
            nw175_[int(10)] = (self).f14
            nw175_[int(11)] = (self).f14
            nw175_[int(12)] = (self).f14
            nw175_[int(13)] = _dafny.SeqWithoutIsStrInference([default__.fm2((self).f15, globalState)])
            nw175_[int(14)] = (self).f14
            nw175_[int(15)] = _dafny.SeqWithoutIsStrInference([(self).f15, d_1107_cf36_])
            nw175_[int(16)] = _dafny.SeqWithoutIsStrInference([(self).f15, d_1107_cf36_])
            nw175_[int(17)] = ((self).f14) + ((self).f14)
            nw175_[int(18)] = ((d_1115_v11_)[d_1107_cf36_] if (d_1107_cf36_) in (d_1115_v11_) else (self).f14)
            nw175_[int(19)] = ((self).f14) + ((self).f14)
            nw175_[int(20)] = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15])
            nw175_[int(21)] = (self).f14
            nw175_[int(22)] = (self).f14
            nw175_[int(23)] = (self).f14
            d_1113_v9_ = nw175_
            (globalState).f4 = ((d_1110_v6_) not in (d_1096_v0_)) or (d_1107_cf36_)
        elif source18_.is_DC25:
            d_1116___mcc_h1_ = source18_.cf37
            d_1117___mcc_h2_ = source18_.cf38
            d_1118___mcc_h3_ = source18_.cf39
            d_1119___mcc_h4_ = source18_.cf40
            d_1120_cf40_ = d_1119___mcc_h4_
            d_1121_cf39_ = d_1118___mcc_h3_
            d_1122_cf38_ = d_1117___mcc_h2_
            d_1123_cf37_ = d_1116___mcc_h1_
            d_1124_v12_: _dafny.Array
            nw176_ = _dafny.Array(_dafny.CodePoint('D'), 23)
            d_1124_v12_ = nw176_
            d_1125_v13_: str
            d_1125_v13_ = _dafny.CodePoint('l')
            index153_ = default__.safeIndex(678, (d_1124_v12_).length(0))
            (d_1124_v12_)[index153_] = d_1125_v13_
            index154_ = default__.safeIndex(678, (d_1124_v12_).length(0))
            rhs111_ = (d_1098_v2_) - (d_1122_cf38_)
            rhs112_ = (default__.safeDivisionInt(421, d_1121_cf39_)) - (d_1098_v2_)
            rhs113_ = default__.fm2((self).f15, globalState)
            rhs114_ = d_1125_v13_
            rhs115_ = d_1125_v13_
            lhs68_ = d_1124_v12_
            lhs69_ = default__.safeIndex(678, (d_1124_v12_).length(0))
            d_1122_cf38_ = rhs111_
            d_1098_v2_ = rhs112_
            d_1123_cf37_ = rhs113_
            r0 = rhs114_
            lhs68_[lhs69_] = rhs115_
            d_1126_v14_: _dafny.Array
            nw177_ = _dafny.Array(int(0), 19)
            d_1126_v14_ = nw177_
            index155_ = default__.safeIndex(267, (d_1126_v14_).length(0))
            (d_1126_v14_)[index155_] = (d_1122_cf38_) + (d_1098_v2_)
            index156_ = default__.safeIndex(267, (d_1126_v14_).length(0))
            (d_1126_v14_)[index156_] = d_1122_cf38_
            d_1127_v15_: _dafny.Seq
            d_1128_v16_: bool
            d_1129_v17_: int
            d_1130_v18_: bool
            out89_: _dafny.Seq
            out90_: bool
            out91_: int
            out92_: bool
            out89_, out90_, out91_, out92_ = default__.m0((self).f15, (d_1099_v3_ if (self).f15 else d_1099_v3_), (d_1125_v13_) in (d_1096_v0_), globalState)
            d_1127_v15_ = out89_
            d_1128_v16_ = out90_
            d_1129_v17_ = out91_
            d_1130_v18_ = out92_
            if not((d_1128_v16_) and (not(d_1130_v18_))):
                d_1131_v19_: D4
                d_1131_v19_ = D4_DC14(True)
                d_1132_v20_: D4
                d_1132_v20_ = D4_DC12((self).f14)
                d_1133_v21_: _dafny.Array
                nw178_ = _dafny.Array(None, 16)
                nw178_[int(0)] = (self).f14
                nw178_[int(1)] = _dafny.SeqWithoutIsStrInference([(d_1131_v19_).cf24, d_1128_v16_, d_1130_v18_, not(not(d_1130_v18_)), d_1128_v16_])
                nw178_[int(2)] = (d_1132_v20_).cf18
                nw178_[int(3)] = (self).f14
                nw178_[int(4)] = (self).f14
                nw178_[int(5)] = (self).f14
                nw178_[int(6)] = (self).f14
                nw178_[int(7)] = (self).f14
                nw178_[int(8)] = (self).f14
                nw178_[int(9)] = (_dafny.SeqWithoutIsStrInference([default__.fm2(d_1130_v18_, globalState)])) + ((self).f14)
                nw178_[int(10)] = (self).f14
                nw178_[int(11)] = (_dafny.SeqWithoutIsStrInference([d_1123_cf37_])) + ((self).f14)
                nw178_[int(12)] = (self).f14
                nw178_[int(13)] = (self).f14
                nw178_[int(14)] = (self).f14
                nw178_[int(15)] = (self).f14
                d_1133_v21_ = nw178_
                d_1134_v22_: _dafny.Map
                d_1134_v22_ = _dafny.Map({d_1129_v17_: d_1123_cf37_})
                d_1135_v23_: _dafny.Map
                d_1135_v23_ = _dafny.Map({((d_1134_v22_)[d_1098_v2_] if (d_1098_v2_) in (d_1134_v22_) else (self).f15): (self).f15})
                index157_ = default__.safeIndex(916, (d_1133_v21_).length(0))
                (d_1133_v21_)[index157_] = _dafny.SeqWithoutIsStrInference([(self).f15, not(not(((d_1135_v23_)[(self).f15] if ((self).f15) in (d_1135_v23_) else d_1123_cf37_))), d_1123_cf37_])
                d_1136_v24_: _dafny.MultiSet
                d_1136_v24_ = _dafny.MultiSet([(self).f15])
                d_1137_v25_: _dafny.MultiSet
                d_1137_v25_ = d_1136_v24_
                index158_ = default__.safeIndex(916, (d_1133_v21_).length(0))
                (d_1133_v21_)[index158_] = (default__.fm43(d_1137_v25_, D7_DC22((self).f15), globalState)) + (((self).f14) + ((self).f14))
                d_1138_v26_: _dafny.Map
                d_1138_v26_ = _dafny.Map({575: d_1098_v2_})
                d_1139_v27_: _dafny.Map
                d_1139_v27_ = _dafny.Map({default__.fm49(d_1123_cf37_, d_1128_v16_, d_1128_v16_, default__.fm39(len(d_1138_v26_), (self).f15, d_1121_cf39_, d_1096_v0_, globalState), globalState): d_1128_v16_})
                d_1139_v27_ = d_1139_v27_
                d_1140_v28_: _dafny.Array
                nw179_ = _dafny.Array(False, 14)
                d_1140_v28_ = nw179_
                d_1140_v28_ = d_1140_v28_
                d_1140_v28_ = d_1140_v28_
                d_1141_v30_: D6
                def iife79_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(908, -420):
                        d_1142_v29_: int = compr_47_
                        if ((908) <= (d_1142_v29_)) and ((d_1142_v29_) < (-420)):
                            coll47_ = coll47_.union(_dafny.Set([(d_1142_v29_) * (-300)]))
                    return _dafny.Set(coll47_)
                d_1141_v30_ = D6_DC17(iife79_()
)
                d_1143_v31_: _dafny.Map
                d_1143_v31_ = _dafny.Map({d_1141_v30_: default__.safeModuloInt(d_1121_cf39_, d_1122_cf38_)})
                d_1143_v31_ = (d_1143_v31_).set(d_1141_v30_, 664)
            elif True:
                d_1129_v17_ = default__.safeModuloInt(d_1121_cf39_, d_1122_cf38_)
                d_1144_v32_: _dafny.MultiSet
                d_1144_v32_ = _dafny.MultiSet([(self).f15, d_1123_cf37_])
                d_1145_v33_: _dafny.Map
                d_1145_v33_ = _dafny.Map({_dafny.CodePoint('x'): d_1098_v2_})
                d_1146_v34_: _dafny.MultiSet
                d_1146_v34_ = _dafny.MultiSet([d_1122_cf38_])
                d_1147_v35_: _dafny.Array
                nw180_ = _dafny.Array(None, 20)
                nw180_[int(0)] = (d_1144_v32_).issubset(d_1144_v32_)
                nw180_[int(1)] = (_dafny.Map({default__.fm39(d_1121_cf39_, (self).f15, (0) - (d_1129_v17_), d_1096_v0_, globalState): d_1129_v17_})) != (d_1145_v33_)
                nw180_[int(2)] = (not(True)) == (d_1123_cf37_)
                nw180_[int(3)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1148_i0_ in range(default__.abs(231))])) <= (d_1096_v0_)
                nw180_[int(4)] = (self).f15
                nw180_[int(5)] = (d_1130_v18_ if d_1123_cf37_ else d_1130_v18_)
                nw180_[int(6)] = d_1130_v18_
                nw180_[int(7)] = d_1123_cf37_
                nw180_[int(8)] = (False if d_1123_cf37_ else (self).f15)
                nw180_[int(9)] = (self).f15
                nw180_[int(10)] = (self).f15
                nw180_[int(11)] = (self).f15
                nw180_[int(12)] = ((self).f14)[default__.safeIndex(d_1121_cf39_, len((self).f14))]
                nw180_[int(13)] = d_1128_v16_
                nw180_[int(14)] = (self).f15
                nw180_[int(15)] = False
                nw180_[int(16)] = d_1130_v18_
                nw180_[int(17)] = d_1123_cf37_
                nw180_[int(18)] = (((d_1146_v34_)[len(d_1127_v15_)] if (len(d_1127_v15_)) in (d_1146_v34_) else d_1121_cf39_)) <= (d_1122_cf38_)
                nw180_[int(19)] = d_1128_v16_
                d_1147_v35_ = nw180_
                index159_ = default__.safeIndex(596, (d_1147_v35_).length(0))
                (d_1147_v35_)[index159_] = (not((self).f15)) or (d_1123_cf37_)
                index160_ = default__.safeIndex(596, (d_1147_v35_).length(0))
                (d_1147_v35_)[index160_] = d_1128_v16_
                d_1149_v36_: _dafny.MultiSet
                d_1149_v36_ = _dafny.MultiSet([d_1147_v35_])
                index161_ = default__.safeIndex(596, (d_1147_v35_).length(0))
                (d_1147_v35_)[index161_] = (d_1149_v36_).issubset(_dafny.MultiSet([d_1147_v35_, d_1147_v35_]))
                d_1150_v37_: _dafny.Map
                d_1150_v37_ = _dafny.Map({(d_1120_cf40_).set((d_1147_v35_)[default__.safeIndex(596, (d_1147_v35_).length(0))], 293): d_1096_v0_})
                d_1151_v38_: _dafny.Map
                d_1151_v38_ = _dafny.Map({d_1128_v16_: (d_1121_cf39_) == (d_1129_v17_)})
                index162_ = default__.safeIndex(267, (d_1126_v14_).length(0))
                rhs116_ = ((d_1150_v37_) | (d_1150_v37_)) | (d_1150_v37_)
                rhs117_ = (d_1127_v15_)[default__.safeIndex((0) - (d_1098_v2_), len(d_1127_v15_))]
                rhs118_ = not(((d_1151_v38_)[default__.fm2((d_1147_v35_)[default__.safeIndex(596, (d_1147_v35_).length(0))], globalState)] if (default__.fm2((d_1147_v35_)[default__.safeIndex(596, (d_1147_v35_).length(0))], globalState)) in (d_1151_v38_) else False))
                rhs119_ = d_1147_v35_
                lhs70_ = d_1126_v14_
                lhs71_ = default__.safeIndex(267, (d_1126_v14_).length(0))
                lhs72_ = globalState
                d_1150_v37_ = rhs116_
                lhs70_[lhs71_] = rhs117_
                lhs72_.f4 = rhs118_
                d_1147_v35_ = rhs119_
                index163_ = default__.safeIndex(596, (d_1147_v35_).length(0))
                (d_1147_v35_)[index163_] = d_1123_cf37_
        elif source18_.is_DC26:
            d_1152___mcc_h5_ = source18_.cf41
            d_1153_cf41_ = d_1152___mcc_h5_
            d_1154_v39_: C0
            nw181_ = C0()
            nw181_.ctor__()
            d_1154_v39_ = nw181_
            d_1155_v40_: _dafny.Seq
            d_1155_v40_ = _dafny.SeqWithoutIsStrInference([(self).f15])
            d_1156_v41_: _dafny.Map
            d_1156_v41_ = _dafny.Map({(0) - (d_1098_v2_): d_1096_v0_})
            d_1157_v42_: _dafny.Map
            d_1157_v42_ = _dafny.Map({d_1153_cf41_: len(((d_1156_v41_)[d_1098_v2_] if (d_1098_v2_) in (d_1156_v41_) else d_1096_v0_))})
            d_1158_v43_: _dafny.Array
            nw182_ = _dafny.Array(None, 15)
            nw182_[int(0)] = (len((self).f16)) > (d_1098_v2_)
            nw182_[int(1)] = (self).f15
            nw182_[int(2)] = (d_1098_v2_) <= (((d_1157_v42_)[(self).f15] if ((self).f15) in (d_1157_v42_) else d_1098_v2_))
            nw182_[int(3)] = d_1153_cf41_
            nw182_[int(4)] = False
            nw182_[int(5)] = (self).f15
            nw182_[int(6)] = (self).f15
            nw182_[int(7)] = True
            nw182_[int(8)] = d_1153_cf41_
            nw182_[int(9)] = d_1153_cf41_
            nw182_[int(10)] = (d_1098_v2_) != ((0) - (d_1098_v2_))
            nw182_[int(11)] = d_1153_cf41_
            nw182_[int(12)] = (self).f15
            nw182_[int(13)] = (d_1096_v0_) <= (default__.fm40(default__.fm50((self).f15, d_1153_cf41_, globalState), d_1098_v2_, d_1153_cf41_, (self).f15, globalState))
            nw182_[int(14)] = d_1153_cf41_
            d_1158_v43_ = nw182_
            index164_ = default__.safeIndex(542, (d_1158_v43_).length(0))
            (d_1158_v43_)[index164_] = (self).f15
            d_1159_v44_: _dafny.Array
            def lambda63_(d_1160_v2_):
                def lambda64_(d_1161_i1_):
                    return (d_1161_i1_) - (d_1160_v2_)

                return lambda64_

            init32_ = lambda63_(d_1098_v2_)
            nw183_ = _dafny.Array(None, 14)
            for i0_32_ in range(nw183_.length(0)):
                nw183_[i0_32_] = init32_(i0_32_)
            d_1159_v44_ = nw183_
            index165_ = default__.safeIndex(542, (d_1158_v43_).length(0))
            rhs120_ = (d_1155_v40_) + ((self).f14)
            rhs121_ = (self).f15
            rhs122_ = d_1159_v44_
            rhs123_ = d_1098_v2_
            lhs73_ = d_1158_v43_
            lhs74_ = default__.safeIndex(542, (d_1158_v43_).length(0))
            d_1155_v40_ = rhs120_
            lhs73_[lhs74_] = rhs121_
            d_1159_v44_ = rhs122_
            d_1098_v2_ = rhs123_
            d_1162_v45_: _dafny.Set
            d_1162_v45_ = _dafny.Set({d_1098_v2_, d_1098_v2_})
            d_1163_v46_: _dafny.Map
            d_1163_v46_ = _dafny.Map({(self).f15: d_1162_v45_})
            d_1164_v47_: str
            d_1164_v47_ = _dafny.CodePoint('l')
            d_1163_v46_ = default__.fm51(d_1153_cf41_, _dafny.CodePoint('f'), (self).f15, d_1164_v47_, globalState)
            d_1165_v48_: _dafny.Map
            d_1165_v48_ = _dafny.Map({d_1098_v2_: len(d_1096_v0_)})
            d_1166_v49_: D11
            d_1166_v49_ = D11_DC32((0) - (((d_1165_v48_)[d_1098_v2_] if (d_1098_v2_) in (d_1165_v48_) else d_1098_v2_)), d_1098_v2_)
            d_1167_v50_: _dafny.Seq
            d_1167_v50_ = _dafny.SeqWithoutIsStrInference([d_1155_v40_])
            pat_let_tv17_ = d_1096_v0_
            d_1168_v51_: _dafny.Array
            nw184_ = _dafny.Array(None, 20)
            nw184_[int(0)] = d_1166_v49_
            nw184_[int(1)] = d_1166_v49_
            nw184_[int(2)] = d_1166_v49_
            nw184_[int(3)] = D11_DC32(len((d_1167_v50_)[default__.safeIndex(d_1098_v2_, len(d_1167_v50_))]), d_1098_v2_)
            nw184_[int(4)] = d_1166_v49_
            def iife80_(_pat_let16_0):
                def iife81_(d_1169_dt__update__tmp_h0_):
                    def iife82_(_pat_let17_0):
                        def iife83_(d_1170_dt__update_hcf46_h0_):
                            return D11_DC32(d_1170_dt__update_hcf46_h0_, (d_1169_dt__update__tmp_h0_).cf47)
                        return iife83_(_pat_let17_0)
                    return iife82_(len(pat_let_tv17_))
                return iife81_(_pat_let16_0)
            nw184_[int(5)] = iife80_(d_1166_v49_)
            nw184_[int(6)] = d_1166_v49_
            nw184_[int(7)] = d_1166_v49_
            nw184_[int(8)] = d_1166_v49_
            nw184_[int(9)] = d_1166_v49_
            nw184_[int(10)] = d_1166_v49_
            nw184_[int(11)] = d_1166_v49_
            nw184_[int(12)] = d_1166_v49_
            nw184_[int(13)] = d_1166_v49_
            nw184_[int(14)] = D11_DC32(d_1098_v2_, len(d_1096_v0_))
            nw184_[int(15)] = (d_1166_v49_ if (self).f15 else d_1166_v49_)
            nw184_[int(16)] = d_1166_v49_
            nw184_[int(17)] = d_1166_v49_
            nw184_[int(18)] = d_1166_v49_
            nw184_[int(19)] = d_1166_v49_
            d_1168_v51_ = nw184_
            d_1168_v51_ = d_1168_v51_
        elif source18_.is_DC23:
            d_1171___mcc_h6_ = source18_.cf35
            d_1172_cf35_ = d_1171___mcc_h6_
            d_1173_v52_: _dafny.Map
            d_1173_v52_ = _dafny.Map({d_1098_v2_: (self).f15})
            d_1174_v53_: _dafny.Map
            d_1174_v53_ = _dafny.Map({len((D12_DC37(d_1096_v0_, (self).f15)).cf59): ((d_1173_v52_)[653] if (653) in (d_1173_v52_) else (self).f15)})
            def iife84_():
                coll48_ = _dafny.Map()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(-743, 427):
                    d_1175_v54_: int = compr_48_
                    if ((-743) <= (d_1175_v54_)) and ((d_1175_v54_) < (427)):
                        coll48_[(d_1175_v54_) * (567)] = (self).f15
                return _dafny.Map(coll48_)
            d_1174_v53_ = iife84_()
            
            d_1176_v55_: _dafny.Array
            def lambda65_(d_1177_i2_):
                return ((self).f15 if (self).f15 else (self).f15)

            init33_ = lambda65_
            nw185_ = _dafny.Array(None, 16)
            for i0_33_ in range(nw185_.length(0)):
                nw185_[i0_33_] = init33_(i0_33_)
            d_1176_v55_ = nw185_
            d_1178_v56_: _dafny.MultiSet
            d_1178_v56_ = _dafny.MultiSet([(self).f15])
            index166_ = default__.safeIndex(940, (d_1176_v55_).length(0))
            (d_1176_v55_)[index166_] = (d_1178_v56_).isdisjoint(_dafny.MultiSet([(self).f15, (self).f15]))
            d_1179_v57_: _dafny.MultiSet
            d_1179_v57_ = d_1178_v56_
            d_1180_v58_: D7
            d_1180_v58_ = D7_DC22((self).f15)
            d_1181_v59_: _dafny.Seq
            d_1181_v59_ = _dafny.SeqWithoutIsStrInference([default__.fm43(d_1179_v57_, d_1180_v58_, globalState), ((self).f14 if (self).f15 else (self).f14), (self).f14, ((self).f14) + ((self).f14)])
            d_1182_v60_: _dafny.Set
            d_1182_v60_ = _dafny.Set({d_1098_v2_, ((self).f16)[default__.safeIndex(847, len((self).f16))], d_1098_v2_, (_dafny.MultiSet([d_1173_v52_])).cardinality, (0) - (d_1098_v2_)})
            index167_ = default__.safeIndex(940, (d_1176_v55_).length(0))
            rhs124_ = default__.fm2((self).f15, globalState)
            rhs125_ = d_1181_v59_
            rhs126_ = (d_1182_v60_) | (_dafny.Set({len(d_1096_v0_), d_1098_v2_}))
            lhs75_ = d_1176_v55_
            lhs76_ = default__.safeIndex(940, (d_1176_v55_).length(0))
            lhs75_[lhs76_] = rhs124_
            d_1181_v59_ = rhs125_
            d_1182_v60_ = rhs126_
            d_1183_v61_: _dafny.Seq
            d_1183_v61_ = _dafny.SeqWithoutIsStrInference([d_1176_v55_, d_1176_v55_, d_1176_v55_, d_1176_v55_])
            d_1098_v2_ = (0) - ((len(d_1183_v61_)) * (d_1098_v2_))
            d_1173_v52_ = d_1174_v53_
        elif True:
            d_1184___mcc_h7_ = source18_.cf42
            d_1185_cf42_ = d_1184___mcc_h7_
            d_1186_v62_: _dafny.MultiSet
            d_1186_v62_ = _dafny.MultiSet([(self).f15])
            d_1187_v63_: _dafny.Map
            d_1187_v63_ = _dafny.Map({(self).f15: len(d_1096_v0_)})
            d_1188_v64_: D8
            d_1188_v64_ = D8_DC25((self).f15, 698, (d_1186_v62_).cardinality, (d_1187_v63_).set((self).f15, 220))
            d_1188_v64_ = d_1188_v64_
            d_1189_v65_: _dafny.Map
            d_1189_v65_ = _dafny.Map({d_1098_v2_: (d_1098_v2_) + (d_1098_v2_)})
            d_1189_v65_ = (d_1189_v65_).set((self).fm5(not((self).f15), globalState), d_1098_v2_)
            d_1190_v66_: _dafny.Array
            nw186_ = _dafny.Array(_dafny.Map({}), 27)
            d_1190_v66_ = nw186_
            d_1191_v67_: _dafny.Array
            nw187_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_1191_v67_ = nw187_
            d_1192_v68_: _dafny.Array
            nw188_ = _dafny.Array(None, 9)
            nw188_[int(0)] = False
            nw188_[int(1)] = (self).f15
            nw188_[int(2)] = (self).f15
            nw188_[int(3)] = (self).f15
            nw188_[int(4)] = (self).f15
            nw188_[int(5)] = True
            nw188_[int(6)] = True
            nw188_[int(7)] = False
            nw188_[int(8)] = (self).f15
            d_1192_v68_ = nw188_
            index168_ = default__.safeIndex(961, (d_1191_v67_).length(0))
            (d_1191_v67_)[index168_] = d_1192_v68_
            d_1193_v69_: _dafny.Array
            nw189_ = _dafny.Array(int(0), 20)
            d_1193_v69_ = nw189_
            index169_ = default__.safeIndex(871, (d_1193_v69_).length(0))
            (d_1193_v69_)[index169_] = d_1098_v2_
            d_1194_v70_: _dafny.Map
            d_1194_v70_ = _dafny.Map({d_1098_v2_: (self).f15})
            d_1195_v71_: _dafny.Set
            d_1195_v71_ = _dafny.Set({d_1194_v70_})
            d_1196_v72_: D4
            d_1196_v72_ = D4_DC12((self).f14)
            d_1197_v73_: _dafny.Map
            d_1197_v73_ = _dafny.Map({d_1098_v2_: d_1196_v72_})
            d_1198_v74_: _dafny.Set
            d_1198_v74_ = _dafny.Set({d_1197_v73_, d_1197_v73_, d_1197_v73_, d_1197_v73_, _dafny.Map({d_1098_v2_: d_1196_v72_})})
            d_1199_v75_: _dafny.Map
            d_1199_v75_ = _dafny.Map({(self).f15: (self).f15})
            d_1200_v76_: T1
            nw190_ = C1()
            nw190_.ctor__((self).f15, (self).f15, ((self).f13).set((self).f15, d_1199_v75_), (self).f14, self.f12)
            d_1200_v76_ = nw190_
            d_1201_v77_: D3
            d_1201_v77_ = D3_DC10((self).f15, d_1098_v2_, d_1200_v76_)
            d_1202_v78_: _dafny.Set
            d_1202_v78_ = _dafny.Set({(self).f15, (self).f15, (d_1201_v77_).cf14})
            d_1203_v79_: _dafny.Map
            d_1203_v79_ = _dafny.Map({(self).f15: d_1202_v78_})
            d_1204_v80_: _dafny.Map
            d_1204_v80_ = _dafny.Map({d_1202_v78_: d_1203_v79_})
            d_1205_v81_: str
            d_1205_v81_ = _dafny.CodePoint('l')
            d_1206_v82_: D4
            d_1206_v82_ = D4_DC14((self).f15)
            d_1207_v83_: _dafny.Map
            d_1207_v83_ = _dafny.Map({d_1205_v81_: default__.fm40(d_1206_v82_, len((self).f16), (self).f15, (self).f15, globalState)})
            index170_ = default__.safeIndex(961, (d_1191_v67_).length(0))
            index171_ = default__.safeIndex(871, (d_1193_v69_).length(0))
            rhs127_ = (d_1190_v66_ if (_dafny.Set({d_1197_v73_})).isdisjoint(d_1198_v74_) else d_1190_v66_)
            rhs128_ = not((self).f15)
            rhs129_ = d_1192_v68_
            rhs130_ = len(((d_1204_v80_)[d_1202_v78_] if (d_1202_v78_) in (d_1204_v80_) else d_1203_v79_))
            rhs131_ = default__.fm52(d_1207_v83_, not((self).f15), default__.fm2((self).f15, globalState), globalState)
            lhs77_ = d_1191_v67_
            lhs78_ = default__.safeIndex(961, (d_1191_v67_).length(0))
            lhs79_ = d_1193_v69_
            lhs80_ = default__.safeIndex(871, (d_1193_v69_).length(0))
            d_1190_v66_ = rhs127_
            r1 = rhs128_
            lhs77_[lhs78_] = rhs129_
            lhs79_[lhs80_] = rhs130_
            d_1195_v71_ = rhs131_
            r1 = False
        d_1098_v2_ = d_1098_v2_
        r1 = not((len((self).f16)) > (len((self).f16)))
        d_1208_v84_: _dafny.Map
        d_1208_v84_ = _dafny.Map({(self).f15: d_1098_v2_})
        d_1209_v85_: D8
        d_1209_v85_ = D8_DC23(d_1208_v84_)
        d_1210_v86_: C1
        nw191_ = C1()
        nw191_.ctor__((self).f15, (self).f15, (self).f13, (self).f14, self.f12)
        d_1210_v86_ = nw191_
        d_1211_v87_: D11
        d_1211_v87_ = D11_DC33((self).f15, d_1098_v2_, False, d_1210_v86_, (self).f15)
        pat_let_tv18_ = d_1096_v0_
        pat_let_tv19_ = d_1098_v2_
        pat_let_tv20_ = d_1098_v2_
        pat_let_tv21_ = d_1096_v0_
        pat_let_tv22_ = d_1098_v2_
        def lambda66_(source20_):
            if source20_.is_DC24:
                d_1212___mcc_h11_ = source20_.cf36
                d_1213_cf36_ = d_1212___mcc_h11_
                return (len(_dafny.Map({pat_let_tv18_: (_dafny.MultiSet((self).f14)).cardinality}))) + (pat_let_tv19_)
            elif source20_.is_DC25:
                d_1214___mcc_h12_ = source20_.cf37
                d_1215___mcc_h13_ = source20_.cf38
                d_1216___mcc_h14_ = source20_.cf39
                d_1217___mcc_h15_ = source20_.cf40
                d_1218_cf40_ = d_1217___mcc_h15_
                d_1219_cf39_ = d_1216___mcc_h14_
                d_1220_cf38_ = d_1215___mcc_h13_
                d_1221_cf37_ = d_1214___mcc_h12_
                return default__.safeDivisionInt(pat_let_tv20_, d_1220_cf38_)
            elif source20_.is_DC26:
                d_1222___mcc_h16_ = source20_.cf41
                d_1223_cf41_ = d_1222___mcc_h16_
                return len((self).f14)
            elif source20_.is_DC23:
                d_1224___mcc_h17_ = source20_.cf35
                d_1225_cf35_ = d_1224___mcc_h17_
                return len(pat_let_tv21_)
            elif True:
                d_1226___mcc_h18_ = source20_.cf42
                d_1227_cf42_ = d_1226___mcc_h18_
                return pat_let_tv22_

        rhs132_ = lambda66_(d_1209_v85_)
        rhs133_ = (d_1098_v2_ if (self).f15 else (d_1211_v87_).cf49)
        d_1098_v2_ = rhs132_
        d_1098_v2_ = rhs133_
        if d_1210_v86_.f18:
            d_1098_v2_ = len((self).f14)
            d_1228_v88_: _dafny.Array
            nw192_ = _dafny.Array(int(0), 14)
            d_1228_v88_ = nw192_
            index172_ = default__.safeIndex(739, (d_1228_v88_).length(0))
            (d_1228_v88_)[index172_] = default__.safeDivisionInt(d_1098_v2_, d_1098_v2_)
            index173_ = default__.safeIndex(739, (d_1228_v88_).length(0))
            (d_1228_v88_)[index173_] = (d_1098_v2_ if default__.fm2((d_1210_v86_).f17, globalState) else d_1098_v2_)
            (globalState).f4 = d_1210_v86_.f18
            d_1229_v89_: _dafny.Map
            d_1229_v89_ = _dafny.Map({(d_1228_v88_)[default__.safeIndex(739, (d_1228_v88_).length(0))]: 49})
            d_1229_v89_ = (d_1229_v89_).set(-248, (d_1228_v88_)[default__.safeIndex(739, (d_1228_v88_).length(0))])
            (d_1210_v86_).f18 = False
        elif True:
            (d_1210_v86_).f18 = d_1210_v86_.f18
            d_1230_v90_: _dafny.Seq
            d_1230_v90_ = _dafny.SeqWithoutIsStrInference([((self).f14).set(default__.safeIndex(d_1098_v2_, len((self).f14)), True), _dafny.SeqWithoutIsStrInference([(self).f15])])
            d_1231_v91_: D2
            d_1231_v91_ = D2_DC4(d_1230_v90_)
            d_1231_v91_ = (d_1231_v91_ if d_1210_v86_.f18 else default__.fm53(769, (d_1210_v86_).f17, globalState))
            d_1232_v92_: str
            d_1232_v92_ = _dafny.CodePoint('y')
            d_1233_v93_: _dafny.Map
            d_1233_v93_ = _dafny.Map({d_1210_v86_.f18: d_1232_v92_})
            d_1233_v93_ = ((d_1233_v93_) | (d_1233_v93_)).set(d_1210_v86_.f18, d_1232_v92_)
            d_1234_v94_: _dafny.Set
            d_1234_v94_ = _dafny.Set({(0) - (d_1098_v2_), len((self).f16)})
            d_1235_v95_: _dafny.Seq
            d_1235_v95_ = _dafny.SeqWithoutIsStrInference([d_1096_v0_])
            d_1098_v2_ = (d_1098_v2_) - (len(((d_1230_v90_)[default__.safeIndex(len((self).fm10((d_1210_v86_).f17, d_1234_v94_, (d_1210_v86_).fm6((d_1235_v95_)[default__.safeIndex(d_1098_v2_, len(d_1235_v95_))], globalState), d_1098_v2_, globalState)), len(d_1230_v90_))]) + ((self).f14)))
            d_1236_v96_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_1236_v96_ = nw193_
            d_1237_v97_: _dafny.Array
            nw194_ = _dafny.Array(None, 11)
            nw194_[int(0)] = (d_1235_v95_)[default__.safeIndex(d_1098_v2_, len(d_1235_v95_))]
            nw194_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nddakcmt"))
            nw194_[int(2)] = d_1096_v0_
            nw194_[int(3)] = d_1096_v0_
            nw194_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkkowjg"))
            nw194_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
            nw194_[int(6)] = d_1096_v0_
            nw194_[int(7)] = d_1096_v0_
            nw194_[int(8)] = d_1096_v0_
            nw194_[int(9)] = d_1096_v0_
            nw194_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kw"))
            d_1237_v97_ = nw194_
            d_1238_v98_: _dafny.Array
            def lambda67_(d_1239_v86_):
                def lambda68_(d_1240_i3_):
                    return (d_1239_v86_).f17

                return lambda68_

            init34_ = lambda67_(d_1210_v86_)
            nw195_ = _dafny.Array(None, 17)
            for i0_34_ in range(nw195_.length(0)):
                nw195_[i0_34_] = init34_(i0_34_)
            d_1238_v98_ = nw195_
            d_1241_v99_: _dafny.Set
            d_1241_v99_ = _dafny.Set({d_1238_v98_, d_1238_v98_})
            d_1242_v100_: D12
            d_1242_v100_ = D12_DC36(d_1237_v97_, d_1098_v2_, default__.fm47(globalState), d_1241_v99_)
            d_1243_v101_: _dafny.Array
            d_1243_v101_ = d_1236_v96_
            rhs134_ = d_1210_v86_.f18
            rhs135_ = (d_1243_v101_)
            rhs136_ = d_1210_v86_.f18
            rhs137_ = (0) - ((self).fm8(not((self).f15), (self).f15, globalState))
            rhs138_ = D12_DC36(d_1237_v97_, d_1098_v2_, _dafny.Map({d_1096_v0_: (self).f16}), _dafny.Set({d_1238_v98_}))
            lhs81_ = globalState
            lhs82_ = d_1210_v86_
            lhs81_.f4 = rhs134_
            d_1236_v96_ = rhs135_
            lhs82_.f18 = rhs136_
            d_1098_v2_ = rhs137_
            d_1242_v100_ = rhs138_
        d_1244_v102_: str
        d_1244_v102_ = _dafny.CodePoint('x')
        r0 = d_1244_v102_
        r1 = (d_1210_v86_).f17
        return r0, r1

    def m12(self, p0, globalState):
        d_1245_v0_: _dafny.Map
        d_1245_v0_ = _dafny.Map({not(p0): p0})
        d_1246_v1_: D3
        d_1246_v1_ = D3_DC8(d_1245_v0_)
        d_1247_v2_: int
        d_1247_v2_ = -280
        d_1248_v3_: _dafny.Array
        nw196_ = _dafny.Array(None, 14)
        nw196_[int(0)] = default__.fm54((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1249_i0_ in range(default__.abs(-545))]))), globalState)
        nw196_[int(1)] = d_1246_v1_
        nw196_[int(2)] = d_1246_v1_
        nw196_[int(3)] = d_1246_v1_
        nw196_[int(4)] = d_1246_v1_
        nw196_[int(5)] = d_1246_v1_
        nw196_[int(6)] = d_1246_v1_
        nw196_[int(7)] = d_1246_v1_
        nw196_[int(8)] = d_1246_v1_
        nw196_[int(9)] = (D3_DC8(d_1245_v0_) if ((self).f14)[default__.safeIndex(d_1247_v2_, len((self).f14))] else d_1246_v1_)
        nw196_[int(10)] = d_1246_v1_
        nw196_[int(11)] = d_1246_v1_
        nw196_[int(12)] = default__.fm54(d_1247_v2_, globalState)
        nw196_[int(13)] = d_1246_v1_
        d_1248_v3_ = nw196_
        index174_ = default__.safeIndex(841, (d_1248_v3_).length(0))
        (d_1248_v3_)[index174_] = d_1246_v1_
        index175_ = default__.safeIndex(841, (d_1248_v3_).length(0))
        rhs139_ = default__.fm54(d_1247_v2_, globalState)
        rhs140_ = d_1247_v2_
        rhs141_ = p0
        lhs83_ = d_1248_v3_
        lhs84_ = default__.safeIndex(841, (d_1248_v3_).length(0))
        lhs85_ = globalState
        lhs83_[lhs84_] = rhs139_
        d_1247_v2_ = rhs140_
        lhs85_.f4 = rhs141_
        d_1250_v4_: _dafny.MultiSet
        d_1250_v4_ = _dafny.MultiSet([not(False), (self).f15, p0, (self).f15, True])
        d_1251_v5_: _dafny.MultiSet
        d_1251_v5_ = _dafny.MultiSet([(d_1250_v4_).isdisjoint(d_1250_v4_)])
        d_1247_v2_ = ((d_1251_v5_)[(d_1247_v2_) != (d_1247_v2_)] if ((d_1247_v2_) != (d_1247_v2_)) in (d_1251_v5_) else -22)
        d_1252_v6_: str
        d_1252_v6_ = _dafny.CodePoint('b')
        d_1252_v6_ = d_1252_v6_
        d_1253_v7_: _dafny.Set
        d_1253_v7_ = _dafny.Set({d_1247_v2_})
        d_1254_v8_: D6
        d_1254_v8_ = D6_DC19(D6_DC17(d_1253_v7_))
        d_1255_v9_: D6
        d_1255_v9_ = D6_DC19(D6_DC19(d_1254_v8_))
        d_1255_v9_ = d_1255_v9_
        d_1256_v10_: _dafny.Array
        def lambda69_(d_1257_v2_):
            def lambda70_(d_1258_i1_):
                return (d_1258_i1_) * (d_1257_v2_)

            return lambda70_

        init35_ = lambda69_(d_1247_v2_)
        nw197_ = _dafny.Array(None, 23)
        for i0_35_ in range(nw197_.length(0)):
            nw197_[i0_35_] = init35_(i0_35_)
        d_1256_v10_ = nw197_
        d_1259_v11_: _dafny.Seq
        d_1259_v11_ = _dafny.SeqWithoutIsStrInference([d_1256_v10_, d_1256_v10_, d_1256_v10_])
        d_1247_v2_ = len(d_1259_v11_)
        d_1260_v12_: D4
        d_1260_v12_ = D4_DC13(p0, p0, d_1247_v2_, d_1247_v2_, default__.fm2((self).f15, globalState))
        d_1261_v13_: D7
        d_1261_v13_ = D7_DC22(p0)
        d_1262_v14_: _dafny.Map
        d_1262_v14_ = _dafny.Map({d_1260_v12_: (d_1261_v13_).cf34})
        d_1263_v15_: _dafny.Seq
        d_1263_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        index176_ = default__.safeIndex(705, (d_1256_v10_).length(0))
        (d_1256_v10_)[index176_] = len(d_1263_v15_)
        d_1264_v16_: _dafny.Seq
        d_1264_v16_ = _dafny.SeqWithoutIsStrInference([d_1262_v14_, (d_1262_v14_).set(default__.fm55(globalState), (self).f15)])
        d_1265_v18_: _dafny.MultiSet
        d_1265_v18_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hxxlj")))])
        pat_let_tv23_ = d_1265_v18_
        d_1266_v19_: _dafny.Set
        def iife85_(_pat_let18_0):
            def iife86_(d_1267_dt__update__tmp_h0_):
                def iife87_(_pat_let19_0):
                    def iife88_(d_1268_dt__update_hcf22_h0_):
                        def iife89_(_pat_let20_0):
                            def iife90_(d_1269_dt__update_hcf19_h0_):
                                return D4_DC13(d_1269_dt__update_hcf19_h0_, (d_1267_dt__update__tmp_h0_).cf20, (d_1267_dt__update__tmp_h0_).cf21, d_1268_dt__update_hcf22_h0_, (d_1267_dt__update__tmp_h0_).cf23)
                            return iife90_(_pat_let20_0)
                        return iife89_((self).f15)
                    return iife88_(_pat_let19_0)
                return iife87_((pat_let_tv23_).cardinality)
            return iife86_(_pat_let18_0)
        d_1266_v19_ = _dafny.Set({default__.fm55(globalState), iife85_(D4_DC13(p0, (self).f15, d_1247_v2_, d_1247_v2_, p0))})
        d_1270_v20_: _dafny.Seq
        def iife91_():
            coll49_ = _dafny.Map()
            compr_49_: D4
            for compr_49_ in (d_1266_v19_).Elements:
                d_1271_v17_: D4 = compr_49_
                if (d_1271_v17_) in (d_1266_v19_):
                    coll49_[d_1271_v17_] = (self).f15
            return _dafny.Map(coll49_)
        d_1270_v20_ = _dafny.SeqWithoutIsStrInference([d_1262_v14_, d_1262_v14_, (d_1264_v16_)[default__.safeIndex((default__.fm56(d_1252_v6_, globalState)).cardinality, len(d_1264_v16_))], iife91_()
        , d_1262_v14_])
        d_1272_v21_: _dafny.Set
        d_1272_v21_ = _dafny.Set({(self).f15, p0, False})
        index177_ = default__.safeIndex(705, (d_1256_v10_).length(0))
        rhs142_ = (d_1270_v20_)[default__.safeIndex((d_1247_v2_) * (d_1247_v2_), len(d_1270_v20_))]
        rhs143_ = len(d_1272_v21_)
        rhs144_ = len(_dafny.Set({(d_1247_v2_) >= (d_1247_v2_), default__.fm2(p0, globalState), p0, not(p0)}))
        lhs86_ = d_1256_v10_
        lhs87_ = default__.safeIndex(705, (d_1256_v10_).length(0))
        d_1262_v14_ = rhs142_
        lhs86_[lhs87_] = rhs143_
        d_1247_v2_ = rhs144_


class C3(T0, T2, T1):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f15: bool = False
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        self._f20: bool = False
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
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f20, f12, f15, f13, f14):
        (self)._f20 = f20
        (self).f12 = f12
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f14 = f14

    def fm5(self, p0, globalState):
        return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpicrv")))

    def fm7(self, p0, globalState):
        return (0) - (default__.safeDivisionInt(default__.safeModuloInt(46, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')]))).cardinality), (len(_dafny.Map({(self).f15: (self).f15}))) - (len((self).f14))))

    def fm8(self, p0, p1, globalState):
        source21_ = D7_DC22((self).f15)
        if source21_.is_DC21:
            d_1273___mcc_h0_ = source21_.cf31
            d_1274___mcc_h1_ = source21_.cf32
            d_1275___mcc_h2_ = source21_.cf33
            d_1276_cf33_ = d_1275___mcc_h2_
            d_1277_cf32_ = d_1274___mcc_h1_
            d_1278_cf31_ = d_1273___mcc_h0_
            return len(_dafny.Map({_dafny.CodePoint('n'): (0) - (len((self).f14))}))
        elif source21_.is_DC22:
            d_1279___mcc_h3_ = source21_.cf34
            d_1280_cf34_ = d_1279___mcc_h3_
            return (len(_dafny.Map({not((self).f15): (self).f15}))) + ((0) - (len(_dafny.Set({_dafny.CodePoint('h'), _dafny.CodePoint('s')}))))
        elif True:
            d_1281___mcc_h4_ = source21_.cf30
            d_1282_cf30_ = d_1281___mcc_h4_
            return (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pp")))) * (315))

    def fm6(self, p0, globalState):
        def iife92_():
            coll50_ = _dafny.Set()
            compr_50_: int
            for compr_50_ in _dafny.IntegerRange(117, 419):
                d_1283_v0_: int = compr_50_
                if ((117) <= (d_1283_v0_)) and ((d_1283_v0_) < (419)):
                    coll50_ = coll50_.union(_dafny.Set([(d_1283_v0_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfbwyy"))), 226])))]))
            return _dafny.Set(coll50_)
        return (_dafny.SeqWithoutIsStrInference([547])) + ((_dafny.SeqWithoutIsStrInference([394, len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f20, (self).f15])): (self).f15})): _dafny.CodePoint('w')})), len(iife92_()
        )])) + (_dafny.SeqWithoutIsStrInference([-612, 737, -32])))

    def fm33(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('a')])) - (_dafny.MultiSet([_dafny.CodePoint('l'), _dafny.CodePoint('y')]))) - ((_dafny.MultiSet([_dafny.CodePoint('g'), _dafny.CodePoint('p'), _dafny.CodePoint('r'), _dafny.CodePoint('u'), _dafny.CodePoint('f')])) | (_dafny.MultiSet([(_dafny.CodePoint('w'))])))

    def fm34(self, p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcwxe"))))) + (default__.safeModuloInt(-276, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1284_i0_ in range(default__.abs(-153))]))))

    def m3(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_1285_v0_: _dafny.Map
        d_1285_v0_ = _dafny.Map({not((self).f20): self.f12})
        d_1286_v1_: C1
        nw198_ = C1()
        nw198_.ctor__((self).f20, (self).f20, (self).f13, (self).f14, ((d_1285_v0_)[(self).f20] if ((self).f20) in (d_1285_v0_) else self.f12))
        d_1286_v1_ = nw198_
        d_1287_v2_: _dafny.Map
        d_1287_v2_ = _dafny.Map({d_1286_v1_: (default__.fm2((d_1286_v1_).f17, globalState) if d_1286_v1_.f18 else (d_1286_v1_).f17)})
        d_1287_v2_ = (d_1287_v2_).set(d_1286_v1_, (d_1286_v1_).f17)
        d_1288_v3_: int
        d_1288_v3_ = 369
        hi6_ = (self).fm34((d_1286_v1_).f17, d_1288_v3_, (self).f20, (d_1286_v1_).f17, globalState)
        for d_1289_i0_ in range(d_1288_v3_, hi6_):
            d_1290_v4_: D4
            d_1290_v4_ = D4_DC12((self).f14)
            d_1291_v5_: _dafny.Seq
            d_1291_v5_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14, (self).f14])
            d_1292_v6_: _dafny.Array
            nw199_ = _dafny.Array(None, 20)
            nw199_[int(0)] = (self).f14
            nw199_[int(1)] = (d_1290_v4_).cf18
            nw199_[int(2)] = (self).f14
            nw199_[int(3)] = (((self).f14).set(default__.safeIndex(d_1289_i0_, len((self).f14)), not(d_1286_v1_.f18))) + (_dafny.SeqWithoutIsStrInference([False, (self).f20]))
            nw199_[int(4)] = (self).f14
            nw199_[int(5)] = (self).f14
            nw199_[int(6)] = (d_1291_v5_)[default__.safeIndex(-846, len(d_1291_v5_))]
            nw199_[int(7)] = (self).f14
            nw199_[int(8)] = (d_1291_v5_)[default__.safeIndex(833, len(d_1291_v5_))]
            nw199_[int(9)] = (self).f14
            nw199_[int(10)] = ((self).f14) + ((self).f14)
            nw199_[int(11)] = ((self).f14) + (_dafny.SeqWithoutIsStrInference([d_1286_v1_.f18, d_1286_v1_.f18, (d_1286_v1_).f17]))
            nw199_[int(12)] = (self).f14
            nw199_[int(13)] = (self).f14
            nw199_[int(14)] = ((self).f14).set(default__.safeIndex(d_1289_i0_, len((self).f14)), (d_1286_v1_).f17)
            nw199_[int(15)] = _dafny.SeqWithoutIsStrInference([(d_1286_v1_).f17])
            nw199_[int(16)] = ((self).f14) + ((self).f14)
            nw199_[int(17)] = (self).f14
            nw199_[int(18)] = (self).f14
            nw199_[int(19)] = _dafny.SeqWithoutIsStrInference([d_1286_v1_.f18])
            d_1292_v6_ = nw199_
            d_1293_v7_: _dafny.Seq
            d_1293_v7_ = _dafny.SeqWithoutIsStrInference([d_1289_i0_, d_1288_v3_])
            d_1294_v8_: T1
            nw200_ = C1()
            nw200_.ctor__((self).f15, default__.fm2(d_1286_v1_.f18, globalState), (self).f13, (self).f14, self.f12)
            d_1294_v8_ = nw200_
            d_1295_v9_: _dafny.Set
            d_1295_v9_ = _dafny.Set({D3_DC10((d_1286_v1_).fm17(d_1293_v7_, 515, globalState), d_1288_v3_, d_1294_v8_)})
            index178_ = default__.safeIndex(38, (d_1292_v6_).length(0))
            (d_1292_v6_)[index178_] = (_dafny.SeqWithoutIsStrInference([d_1286_v1_.f18])) + ((d_1291_v5_)[default__.safeIndex(len(d_1295_v9_), len(d_1291_v5_))])
            index179_ = default__.safeIndex(38, (d_1292_v6_).length(0))
            (d_1292_v6_)[index179_] = (d_1294_v8_).f14
            d_1296_v10_: _dafny.Map
            d_1296_v10_ = _dafny.Map({(d_1286_v1_).f17: 91})
            d_1297_v11_: D8
            d_1297_v11_ = D8_DC25((self).f15, d_1288_v3_, d_1288_v3_, d_1296_v10_)
            d_1298_v12_: _dafny.Seq
            d_1298_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jas"))
            d_1299_v13_: _dafny.Map
            d_1299_v13_ = _dafny.Map({(self).f15: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1300_i2_ in range(default__.abs(505))])})
            d_1301_v14_: str
            d_1301_v14_ = _dafny.CodePoint('j')
            d_1302_v15_: _dafny.Map
            d_1302_v15_ = _dafny.Map({d_1289_i0_: d_1298_v12_})
            d_1303_v16_: _dafny.Array
            nw201_ = _dafny.Array(None, 14)
            nw201_[int(0)] = default__.fm35(d_1288_v3_, d_1297_v11_, globalState)
            nw201_[int(1)] = d_1298_v12_
            nw201_[int(2)] = d_1298_v12_
            nw201_[int(3)] = d_1298_v12_
            nw201_[int(4)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1304_i1_ in range(default__.abs(809))])
            nw201_[int(5)] = d_1298_v12_
            nw201_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpq"))
            nw201_[int(7)] = d_1298_v12_
            nw201_[int(8)] = (d_1298_v12_) + (d_1298_v12_)
            nw201_[int(9)] = ((d_1299_v13_)[False] if (False) in (d_1299_v13_) else d_1298_v12_)
            nw201_[int(10)] = (d_1298_v12_).set(default__.safeIndex(d_1288_v3_, len(d_1298_v12_)), d_1301_v14_)
            nw201_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpvq"))) + (((d_1302_v15_)[(0) - (d_1288_v3_)] if ((0) - (d_1288_v3_)) in (d_1302_v15_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytiar"))))
            nw201_[int(12)] = d_1298_v12_
            nw201_[int(13)] = d_1298_v12_
            d_1303_v16_ = nw201_
            index180_ = default__.safeIndex(562, (d_1303_v16_).length(0))
            (d_1303_v16_)[index180_] = d_1298_v12_
            index181_ = default__.safeIndex(562, (d_1303_v16_).length(0))
            rhs145_ = (d_1298_v12_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iw"))) + (d_1298_v12_))
            rhs146_ = d_1286_v1_.f18
            lhs88_ = d_1303_v16_
            lhs89_ = default__.safeIndex(562, (d_1303_v16_).length(0))
            lhs90_ = globalState
            lhs88_[lhs89_] = rhs145_
            lhs90_.f4 = rhs146_
            r1 = d_1288_v3_
            d_1305_v17_: C1
            nw202_ = C1()
            nw202_.ctor__(d_1286_v1_.f18, d_1286_v1_.f18, (self).f13, (self).f14, self.f12)
            d_1305_v17_ = nw202_
        d_1306_v18_: _dafny.Seq
        d_1306_v18_ = _dafny.SeqWithoutIsStrInference([d_1288_v3_])
        d_1307_v19_: _dafny.Seq
        d_1307_v19_ = _dafny.SeqWithoutIsStrInference([d_1306_v18_])
        d_1308_v20_: _dafny.Set
        d_1308_v20_ = _dafny.Set({(d_1307_v19_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cekuw"))), len(d_1307_v19_))]})
        d_1309_v21_: _dafny.Map
        d_1309_v21_ = _dafny.Map({(self).f15: (d_1306_v18_).set(default__.safeIndex((0) - (d_1288_v3_), len(d_1306_v18_)), len(d_1308_v20_))})
        hi7_ = (len((self).f14)) + (d_1288_v3_)
        for d_1310_i3_ in range(len(((d_1309_v21_)[(d_1286_v1_).f17] if ((d_1286_v1_).f17) in (d_1309_v21_) else d_1306_v18_)), hi7_):
            d_1311_v22_: _dafny.Map
            d_1311_v22_ = _dafny.Map({d_1288_v3_: (d_1286_v1_).fm17((_dafny.SeqWithoutIsStrInference([d_1310_i3_])).set(default__.safeIndex(d_1310_i3_, len(_dafny.SeqWithoutIsStrInference([d_1310_i3_]))), (0) - (d_1288_v3_)), d_1310_i3_, globalState)})
            d_1312_v23_: _dafny.Seq
            d_1312_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_1311_v22_ = (d_1311_v22_).set(len(d_1311_v22_), (d_1312_v23_) < (d_1312_v23_))
            (globalState).f4 = (self).f20
            d_1313_v24_: _dafny.Array
            def lambda71_(d_1314_v18_, d_1315_v1_, d_1316_v21_):
                def lambda72_(d_1317_i4_):
                    return (d_1317_i4_) + ((d_1314_v18_)[default__.safeIndex((_dafny.MultiSet(((d_1316_v21_)[d_1315_v1_.f18] if (d_1315_v1_.f18) in (d_1316_v21_) else _dafny.SeqWithoutIsStrInference([len(d_1314_v18_) for d_1318_i5_ in range(default__.abs(375))])))).cardinality, len(d_1314_v18_))])

                return lambda72_

            init36_ = lambda71_(d_1306_v18_, d_1286_v1_, d_1309_v21_)
            nw203_ = _dafny.Array(None, 23)
            for i0_36_ in range(nw203_.length(0)):
                nw203_[i0_36_] = init36_(i0_36_)
            d_1313_v24_ = nw203_
            index182_ = default__.safeIndex(219, (d_1313_v24_).length(0))
            (d_1313_v24_)[index182_] = (d_1286_v1_).fm5(True, globalState)
            index183_ = default__.safeIndex(219, (d_1313_v24_).length(0))
            rhs147_ = (d_1288_v3_) + (d_1310_i3_)
            rhs148_ = (d_1288_v3_) == (d_1288_v3_)
            rhs149_ = d_1310_i3_
            lhs91_ = d_1313_v24_
            lhs92_ = default__.safeIndex(219, (d_1313_v24_).length(0))
            r1 = rhs147_
            r0 = rhs148_
            lhs91_[lhs92_] = rhs149_
            r0 = not ((d_1286_v1_).f17) or (not(not((d_1286_v1_).f17)))
        d_1319_v25_: _dafny.Seq
        d_1319_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
        d_1320_v26_: _dafny.Map
        d_1320_v26_ = _dafny.Map({d_1288_v3_: d_1319_v25_})
        d_1321_v27_: _dafny.Seq
        d_1321_v27_ = _dafny.SeqWithoutIsStrInference([((d_1320_v26_)[len(_dafny.Map({(self).f15: (self).f14}))] if (len(_dafny.Map({(self).f15: (self).f14}))) in (d_1320_v26_) else d_1319_v25_), d_1319_v25_])
        d_1322_v28_: _dafny.Set
        d_1322_v28_ = _dafny.Set({(0) - ((self).fm7((0) - (len(d_1306_v18_)), globalState)), (d_1288_v3_) + (d_1288_v3_), len(d_1321_v27_)})
        d_1323_v29_: _dafny.MultiSet
        d_1323_v29_ = _dafny.MultiSet([len(d_1308_v20_)])
        d_1324_v30_: str
        d_1324_v30_ = _dafny.CodePoint('b')
        rhs150_ = d_1288_v3_
        rhs151_ = _dafny.Set({(d_1323_v29_).cardinality, default__.safeDivisionInt(d_1288_v3_, d_1288_v3_), d_1288_v3_, d_1288_v3_})
        rhs152_ = (len(_dafny.SeqWithoutIsStrInference([d_1324_v30_, _dafny.CodePoint('c'), d_1324_v30_, d_1324_v30_]))) - (d_1288_v3_)
        rhs153_ = d_1319_v25_
        d_1288_v3_ = rhs150_
        d_1322_v28_ = rhs151_
        r1 = rhs152_
        d_1319_v25_ = rhs153_
        d_1325_v31_: _dafny.Map
        d_1325_v31_ = _dafny.Map({436: (d_1286_v1_).fm17(d_1306_v18_, len(d_1322_v28_), globalState)})
        d_1326_v32_: D7
        d_1326_v32_ = D7_DC21(d_1325_v31_, False, ((self).f14)[default__.safeIndex(d_1288_v3_, len((self).f14))])
        d_1327_v33_: _dafny.MultiSet
        d_1327_v33_ = _dafny.MultiSet([d_1325_v31_, (d_1325_v31_) | ((d_1326_v32_).cf31), (d_1325_v31_).set(d_1288_v3_, not(False)), d_1325_v31_])
        d_1327_v33_ = d_1327_v33_
        d_1328_v34_: _dafny.Set
        d_1328_v34_ = _dafny.Set({(d_1286_v1_).f17})
        rhs154_ = (d_1288_v3_) <= (d_1288_v3_)
        rhs155_ = len(_dafny.Map({(d_1328_v34_).intersection(_dafny.Set({False})): d_1288_v3_}))
        r0 = rhs154_
        r1 = rhs155_
        r0 = (d_1323_v29_).issubset(d_1323_v29_)
        d_1329_v35_: D4
        d_1329_v35_ = D4_DC13(not((self).f15), d_1286_v1_.f18, d_1288_v3_, d_1288_v3_, (self).f20)
        pat_let_tv24_ = d_1325_v31_
        pat_let_tv25_ = d_1286_v1_
        pat_let_tv26_ = d_1306_v18_
        pat_let_tv27_ = d_1288_v3_
        def lambda73_(source22_):
            if source22_.is_DC13:
                d_1330___mcc_h0_ = source22_.cf19
                d_1331___mcc_h1_ = source22_.cf20
                d_1332___mcc_h2_ = source22_.cf21
                d_1333___mcc_h3_ = source22_.cf22
                d_1334___mcc_h4_ = source22_.cf23
                d_1335_cf23_ = d_1334___mcc_h4_
                d_1336_cf22_ = d_1333___mcc_h3_
                d_1337_cf21_ = d_1332___mcc_h2_
                d_1338_cf20_ = d_1331___mcc_h1_
                d_1339_cf19_ = d_1330___mcc_h0_
                return len(pat_let_tv24_)
            elif source22_.is_DC14:
                d_1340___mcc_h5_ = source22_.cf24
                d_1341_cf24_ = d_1340___mcc_h5_
                return len(_dafny.Map({(self).f15: (pat_let_tv25_).f17}))
            elif True:
                d_1342___mcc_h6_ = source22_.cf18
                d_1343_cf18_ = d_1342___mcc_h6_
                return (len(pat_let_tv26_)) + (len(_dafny.SeqWithoutIsStrInference([pat_let_tv27_ for d_1344_i6_ in range(default__.abs(665))])))

        r1 = lambda73_(d_1329_v35_)
        d_1345_v37_: _dafny.Map
        d_1345_v37_ = _dafny.Map({d_1286_v1_.f18: d_1288_v3_})
        d_1346_v38_: _dafny.Set
        d_1346_v38_ = _dafny.Set({d_1319_v25_, d_1319_v25_, d_1319_v25_, _dafny.SeqWithoutIsStrInference([d_1324_v30_ for d_1347_i7_ in range(default__.abs(58))]), default__.fm35(d_1288_v3_, D8_DC25((self).f15, d_1288_v3_, d_1288_v3_, d_1345_v37_), globalState)})
        def iife93_():
            coll51_ = _dafny.Set()
            compr_51_: _dafny.Seq
            for compr_51_ in (d_1321_v27_).Elements:
                d_1348_v36_: _dafny.Seq = compr_51_
                if (d_1348_v36_) in (d_1321_v27_):
                    coll51_ = coll51_.union(_dafny.Set([d_1348_v36_]))
            return _dafny.Set(coll51_)
        r2 = default__.fm36(d_1288_v3_, d_1288_v3_, (iife93_()
         if (d_1286_v1_).f17 else d_1346_v38_), globalState)
        r3 = not ((self).f15) or (not((d_1286_v1_).f17))
        return r0, r1, r2, r3

    def m9(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        r3: _dafny.Map = _dafny.Map({})
        (self).m10(globalState)
        d_1349_v0_: int
        d_1349_v0_ = -546
        d_1350_v1_: _dafny.Set
        d_1350_v1_ = _dafny.Set({d_1349_v0_})
        d_1351_i0_: int
        d_1351_i0_ = 0
        with _dafny.label("10"):
            while (d_1350_v1_).issubset(d_1350_v1_):
                with _dafny.c_label("10"):
                    if (d_1351_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1351_i0_ = (d_1351_i0_) + (1)
                    d_1352_v2_: _dafny.Seq
                    d_1352_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vakf"))
                    d_1353_v3_: str
                    d_1353_v3_ = _dafny.CodePoint('a')
                    rhs156_ = d_1352_v2_
                    rhs157_ = (d_1349_v0_) * (d_1349_v0_)
                    rhs158_ = (self).f20
                    rhs159_ = d_1353_v3_
                    lhs93_ = globalState
                    d_1352_v2_ = rhs156_
                    r0 = rhs157_
                    lhs93_.f4 = rhs158_
                    d_1353_v3_ = rhs159_
                    d_1354_v4_: _dafny.Map
                    d_1354_v4_ = _dafny.Map({897: p1})
                    d_1355_v5_: _dafny.Array
                    nw204_ = _dafny.Array(None, 17)
                    nw204_[int(0)] = (self).f20
                    nw204_[int(1)] = p0
                    nw204_[int(2)] = p1
                    nw204_[int(3)] = not(p1)
                    nw204_[int(4)] = p0
                    nw204_[int(5)] = True
                    nw204_[int(6)] = (self).f15
                    nw204_[int(7)] = p1
                    nw204_[int(8)] = (self).f15
                    nw204_[int(9)] = not(p0)
                    nw204_[int(10)] = (self).f20
                    nw204_[int(11)] = (self).f15
                    nw204_[int(12)] = p0
                    nw204_[int(13)] = p1
                    nw204_[int(14)] = p0
                    nw204_[int(15)] = default__.fm2(p1, globalState)
                    nw204_[int(16)] = ((d_1354_v4_)[d_1349_v0_] if (d_1349_v0_) in (d_1354_v4_) else (self).f20)
                    d_1355_v5_ = nw204_
                    d_1356_v6_: D2
                    d_1356_v6_ = D2_DC6(d_1355_v5_, d_1355_v5_)
                    d_1357_v7_: _dafny.Map
                    d_1357_v7_ = _dafny.Map({p0: d_1355_v5_})
                    pat_let_tv28_ = d_1355_v5_
                    d_1358_v8_: _dafny.Array
                    nw205_ = _dafny.Array(None, 29)
                    nw205_[int(0)] = d_1355_v5_
                    nw205_[int(1)] = d_1355_v5_
                    nw205_[int(2)] = d_1355_v5_
                    nw205_[int(3)] = d_1355_v5_
                    nw205_[int(4)] = d_1355_v5_
                    nw205_[int(5)] = d_1355_v5_
                    nw205_[int(6)] = d_1355_v5_
                    nw205_[int(7)] = d_1355_v5_
                    nw205_[int(8)] = d_1355_v5_
                    nw205_[int(9)] = d_1355_v5_
                    nw205_[int(10)] = d_1355_v5_
                    nw205_[int(11)] = d_1355_v5_
                    nw205_[int(12)] = d_1355_v5_
                    nw205_[int(13)] = d_1355_v5_
                    nw205_[int(14)] = d_1355_v5_
                    nw205_[int(15)] = d_1355_v5_
                    nw205_[int(16)] = d_1355_v5_
                    nw205_[int(17)] = (d_1355_v5_ if (self).f20 else d_1355_v5_)
                    nw205_[int(18)] = (d_1356_v6_).cf9
                    nw205_[int(19)] = d_1355_v5_
                    nw205_[int(20)] = ((d_1357_v7_)[True] if (True) in (d_1357_v7_) else d_1355_v5_)
                    def iife94_(_pat_let21_0):
                        def iife95_(d_1359_dt__update__tmp_h0_):
                            def iife96_(_pat_let22_0):
                                def iife97_(d_1360_dt__update_hcf8_h0_):
                                    return D2_DC6(d_1360_dt__update_hcf8_h0_, (d_1359_dt__update__tmp_h0_).cf9)
                                return iife97_(_pat_let22_0)
                            return iife96_(pat_let_tv28_)
                        return iife95_(_pat_let21_0)
                    nw205_[int(21)] = (iife94_(D2_DC6(d_1355_v5_, d_1355_v5_))).cf9
                    nw205_[int(22)] = d_1355_v5_
                    nw205_[int(23)] = d_1355_v5_
                    nw205_[int(24)] = d_1355_v5_
                    nw205_[int(25)] = d_1355_v5_
                    nw205_[int(26)] = d_1355_v5_
                    nw205_[int(27)] = d_1355_v5_
                    nw205_[int(28)] = d_1355_v5_
                    d_1358_v8_ = nw205_
                    index184_ = default__.safeIndex(911, (d_1355_v5_).length(0))
                    (d_1355_v5_)[index184_] = (self).f15
                    index185_ = default__.safeIndex(911, (d_1355_v5_).length(0))
                    rhs160_ = d_1358_v8_
                    rhs161_ = True
                    rhs162_ = ((0) - (-910)) != (d_1349_v0_)
                    rhs163_ = d_1349_v0_
                    rhs164_ = d_1349_v0_
                    lhs94_ = d_1355_v5_
                    lhs95_ = default__.safeIndex(911, (d_1355_v5_).length(0))
                    lhs96_ = globalState
                    d_1358_v8_ = rhs160_
                    lhs94_[lhs95_] = rhs161_
                    lhs96_.f4 = rhs162_
                    d_1349_v0_ = rhs163_
                    r0 = rhs164_
                    r0 = (d_1349_v0_) - (d_1349_v0_)
                    d_1361_v9_: _dafny.MultiSet
                    d_1361_v9_ = _dafny.MultiSet([(d_1355_v5_)[default__.safeIndex(911, (d_1355_v5_).length(0))], (self).f15, p0, not(p0)])
                    r1 = (d_1349_v0_) - ((d_1361_v9_).cardinality)
                    pass
            pass
        d_1362_v10_: _dafny.MultiSet
        d_1362_v10_ = _dafny.MultiSet([(self).f20, p0, p1])
        hi8_ = (d_1362_v10_).cardinality
        for d_1363_i1_ in range(d_1349_v0_, hi8_):
            d_1364_v11_: _dafny.Map
            d_1364_v11_ = _dafny.Map({p2: (self).f20})
            if ((d_1364_v11_)[p2] if (p2) in (d_1364_v11_) else ((self).f15) and (p0)):
                d_1365_v12_: _dafny.Map
                d_1365_v12_ = _dafny.Map({((self).f15 if not((self).f20) else p0): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omhw"))})
                d_1366_v13_: str
                d_1366_v13_ = _dafny.CodePoint('m')
                rhs165_ = d_1365_v12_
                rhs166_ = (164) + ((d_1362_v10_).cardinality)
                rhs167_ = d_1366_v13_
                d_1365_v12_ = rhs165_
                d_1349_v0_ = rhs166_
                d_1366_v13_ = rhs167_
                d_1367_v14_: D4
                d_1367_v14_ = D4_DC12(_dafny.SeqWithoutIsStrInference([(self).f15]))
                d_1368_v15_: _dafny.Map
                d_1368_v15_ = _dafny.Map({(self).f14: d_1349_v0_})
                d_1369_v16_: _dafny.Map
                d_1369_v16_ = _dafny.Map({d_1363_i1_: len(d_1368_v15_)})
                d_1370_v17_: _dafny.Seq
                d_1370_v17_ = _dafny.SeqWithoutIsStrInference([d_1369_v16_])
                d_1371_v18_: _dafny.Map
                d_1371_v18_ = _dafny.Map({(0) - (d_1363_i1_): (0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymee"))).set(default__.safeIndex(len(d_1370_v17_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymee")))), d_1366_v13_)))})
                d_1372_v19_: C1
                nw206_ = C1()
                nw206_.ctor__((self).f20, default__.fm2(p0, globalState), ((self).f13 if ((self).f14)[default__.safeIndex((0) - (len(p2)), len((self).f14))] else (self).f13), ((d_1367_v14_).cf18).set(default__.safeIndex(len((d_1371_v18_).set(d_1363_i1_, len(d_1350_v1_))), len((d_1367_v14_).cf18)), p1), self.f12)
                d_1372_v19_ = nw206_
                (d_1372_v19_).f18 = p0
                d_1373_v20_: _dafny.Array
                def lambda74_(d_1374_v19_):
                    def lambda75_(d_1375_i2_):
                        return (d_1374_v19_).f17

                    return lambda75_

                init37_ = lambda74_(d_1372_v19_)
                nw207_ = _dafny.Array(None, 22)
                for i0_37_ in range(nw207_.length(0)):
                    nw207_[i0_37_] = init37_(i0_37_)
                d_1373_v20_ = nw207_
                d_1376_v21_: _dafny.Seq
                d_1376_v21_ = _dafny.SeqWithoutIsStrInference([d_1373_v20_])
                d_1377_v22_: _dafny.Map
                d_1377_v22_ = _dafny.Map({((d_1376_v21_) + (d_1376_v21_)).set(default__.safeIndex((self).fm7(708, globalState), len((d_1376_v21_) + (d_1376_v21_))), d_1373_v20_): d_1363_i1_})
                d_1377_v22_ = (d_1377_v22_).set(d_1376_v21_, (0) - ((d_1363_i1_) + (d_1349_v0_)))
                d_1378_v23_: _dafny.Map
                d_1378_v23_ = _dafny.Map({(self).f15: p0})
                d_1379_v24_: D3
                d_1379_v24_ = D3_DC8(d_1378_v23_)
                d_1380_v25_: _dafny.Map
                d_1380_v25_ = _dafny.Map({d_1349_v0_: True})
                d_1379_v24_ = D3_DC8((d_1378_v23_).set(((d_1380_v25_)[d_1349_v0_] if (d_1349_v0_) in (d_1380_v25_) else (self).f20), False))
            elif True:
                d_1381_v26_: D4
                d_1381_v26_ = D4_DC12((self).f14)
                d_1382_v27_: _dafny.MultiSet
                d_1382_v27_ = _dafny.MultiSet([d_1363_i1_, len(d_1350_v1_), d_1363_i1_])
                d_1383_v28_: str
                d_1383_v28_ = _dafny.CodePoint('v')
                d_1384_v29_: _dafny.Array
                nw208_ = _dafny.Array(None, 24)
                nw208_[int(0)] = d_1381_v26_
                nw208_[int(1)] = d_1381_v26_
                nw208_[int(2)] = d_1381_v26_
                nw208_[int(3)] = d_1381_v26_
                nw208_[int(4)] = d_1381_v26_
                nw208_[int(5)] = d_1381_v26_
                nw208_[int(6)] = d_1381_v26_
                nw208_[int(7)] = d_1381_v26_
                nw208_[int(8)] = d_1381_v26_
                nw208_[int(9)] = d_1381_v26_
                nw208_[int(10)] = d_1381_v26_
                nw208_[int(11)] = d_1381_v26_
                nw208_[int(12)] = d_1381_v26_
                nw208_[int(13)] = default__.fm37((self).f14, d_1382_v27_, 843, _dafny.CodePoint('p'), globalState)
                def iife98_(_pat_let23_0):
                    def iife99_(d_1385_dt__update__tmp_h1_):
                        def iife100_(_pat_let24_0):
                            def iife101_(d_1386_dt__update_hcf18_h0_):
                                return D4_DC12(d_1386_dt__update_hcf18_h0_)
                            return iife101_(_pat_let24_0)
                        return iife100_((self).f14)
                    return iife99_(_pat_let23_0)
                nw208_[int(14)] = iife98_(d_1381_v26_)
                nw208_[int(15)] = d_1381_v26_
                nw208_[int(16)] = d_1381_v26_
                nw208_[int(17)] = d_1381_v26_
                nw208_[int(18)] = d_1381_v26_
                nw208_[int(19)] = default__.fm37((self).f14, d_1382_v27_, d_1363_i1_, d_1383_v28_, globalState)
                def iife102_(_pat_let25_0):
                    def iife103_(d_1387_dt__update__tmp_h2_):
                        def iife104_(_pat_let26_0):
                            def iife105_(d_1388_dt__update_hcf18_h1_):
                                return D4_DC12(d_1388_dt__update_hcf18_h1_)
                            return iife105_(_pat_let26_0)
                        return iife104_((self).f14)
                    return iife103_(_pat_let25_0)
                nw208_[int(20)] = iife102_(d_1381_v26_)
                nw208_[int(21)] = d_1381_v26_
                nw208_[int(22)] = d_1381_v26_
                nw208_[int(23)] = d_1381_v26_
                d_1384_v29_ = nw208_
                d_1389_v30_: _dafny.Array
                nw209_ = _dafny.Array(None, 6)
                nw209_[int(0)] = (self).f15
                nw209_[int(1)] = p1
                nw209_[int(2)] = p1
                nw209_[int(3)] = p1
                nw209_[int(4)] = (self).f20
                nw209_[int(5)] = (self).f20
                d_1389_v30_ = nw209_
                d_1390_v31_: _dafny.MultiSet
                d_1390_v31_ = _dafny.MultiSet([d_1389_v30_, d_1389_v30_])
                rhs168_ = d_1363_i1_
                rhs169_ = d_1384_v29_
                rhs170_ = default__.fm2((d_1389_v30_) not in (d_1390_v31_), globalState)
                r0 = rhs168_
                d_1384_v29_ = rhs169_
                r2 = rhs170_
                d_1391_v32_: _dafny.Seq
                d_1391_v32_ = _dafny.SeqWithoutIsStrInference([(d_1349_v0_) > ((D0_DC2(d_1363_i1_, p0)).cf1), not (not(not((self).f15))) or (not(p1))])
                d_1391_v32_ = ((self).f14) + ((self).f14)
                d_1392_v33_: _dafny.Seq
                d_1392_v33_ = _dafny.SeqWithoutIsStrInference([p2, p2, (p2) + (p2), p2])
                d_1392_v33_ = (d_1392_v33_) + ((d_1392_v33_) + (_dafny.SeqWithoutIsStrInference([p2, p2])))
                index186_ = default__.safeIndex(443, (d_1389_v30_).length(0))
                (d_1389_v30_)[index186_] = p0
                d_1393_v34_: _dafny.Set
                d_1393_v34_ = _dafny.Set({(self).f15, (self).f20, (self).f20})
                index187_ = default__.safeIndex(443, (d_1389_v30_).length(0))
                (d_1389_v30_)[index187_] = (len(d_1393_v34_)) > (d_1363_i1_)
                (self).m10(globalState)
            d_1394_v35_: T4
            nw210_ = C2()
            nw210_.ctor__(_dafny.SeqWithoutIsStrInference([78]), (self).f20, (self).f13, (self).f14, self.f12)
            d_1394_v35_ = nw210_
            d_1395_v36_: _dafny.Seq
            d_1395_v36_ = _dafny.SeqWithoutIsStrInference([d_1394_v35_])
            d_1396_v37_: _dafny.Array
            nw211_ = _dafny.Array(None, 12)
            nw211_[int(0)] = (d_1394_v35_ if (self).f20 else d_1394_v35_)
            nw211_[int(1)] = d_1394_v35_
            nw211_[int(2)] = d_1394_v35_
            nw211_[int(3)] = d_1394_v35_
            nw211_[int(4)] = d_1394_v35_
            nw211_[int(5)] = d_1394_v35_
            nw211_[int(6)] = d_1394_v35_
            nw211_[int(7)] = d_1394_v35_
            nw211_[int(8)] = (d_1395_v36_)[default__.safeIndex(487, len(d_1395_v36_))]
            nw211_[int(9)] = d_1394_v35_
            nw211_[int(10)] = d_1394_v35_
            nw211_[int(11)] = d_1394_v35_
            d_1396_v37_ = nw211_
            index188_ = default__.safeIndex(556, (d_1396_v37_).length(0))
            (d_1396_v37_)[index188_] = d_1394_v35_
            d_1397_v38_: _dafny.Set
            d_1397_v38_ = _dafny.Set({(self).f15, False, p1, p0})
            d_1398_v39_: _dafny.Seq
            d_1398_v39_ = _dafny.SeqWithoutIsStrInference([d_1363_i1_, len(d_1397_v38_), d_1363_i1_, len((self).f14)])
            index189_ = default__.safeIndex(556, (d_1396_v37_).length(0))
            nw212_ = C2()
            nw212_.ctor__(d_1398_v39_, (self).f15, (self).f13, (self).f14, self.f12)
            (d_1396_v37_)[index189_] = nw212_
            d_1399_v40_: C0
            nw213_ = C0()
            nw213_.ctor__()
            d_1399_v40_ = nw213_
            d_1400_v41_: _dafny.Seq
            d_1400_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1401_i3_ in range(default__.abs(219))])])
            d_1400_v41_ = (d_1400_v41_) + (d_1400_v41_)
        r0 = default__.safeModuloInt(d_1349_v0_, (default__.fm57(p2, (self).f15, _dafny.Map({(d_1362_v10_).cardinality: (self).f15}), globalState)).cardinality)
        d_1402_v42_: str
        d_1402_v42_ = _dafny.CodePoint('r')
        d_1402_v42_ = d_1402_v42_
        d_1403_v43_: T1
        nw214_ = C1()
        nw214_.ctor__((self).f15, (self).f15, (self).f13, (self).f14, self.f12)
        d_1403_v43_ = nw214_
        d_1404_v44_: D3
        d_1404_v44_ = D3_DC10(p0, d_1349_v0_, d_1403_v43_)
        d_1404_v44_ = d_1404_v44_
        r0 = d_1349_v0_
        r1 = d_1349_v0_
        r2 = (self).f20
        d_1405_v46_: _dafny.Map
        d_1405_v46_ = _dafny.Map({d_1349_v0_: p0})
        def iife106_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(455, 775):
                d_1406_v45_: int = compr_52_
                if ((455) <= (d_1406_v45_)) and ((d_1406_v45_) < (775)):
                    coll52_[default__.safeModuloInt(d_1406_v45_, d_1349_v0_)] = (D7_DC21(_dafny.Map({(d_1362_v10_).cardinality: (self).f15}), (self).f15, p1)).cf32
            return _dafny.Map(coll52_)
        r3 = ((iife106_()
        ) | (d_1405_v46_)) | (d_1405_v46_)
        return r0, r1, r2, r3

    def m10(self, globalState):
        d_1407_v0_: _dafny.Array
        nw215_ = _dafny.Array(False, 23)
        d_1407_v0_ = nw215_
        d_1408_v1_: int
        d_1408_v1_ = 414
        d_1409_v2_: _dafny.Map
        d_1409_v2_ = _dafny.Map({d_1407_v0_: d_1408_v1_})
        d_1409_v2_ = d_1409_v2_
        d_1408_v1_ = d_1408_v1_
        d_1410_v3_: D8
        d_1410_v3_ = D8_DC24((self).f15)
        pat_let_tv29_ = d_1408_v1_
        pat_let_tv30_ = d_1408_v1_
        def lambda76_(source23_):
            if source23_.is_DC24:
                d_1411___mcc_h0_ = source23_.cf36
                d_1412_cf36_ = d_1411___mcc_h0_
                return pat_let_tv29_
            elif source23_.is_DC25:
                d_1413___mcc_h1_ = source23_.cf37
                d_1414___mcc_h2_ = source23_.cf38
                d_1415___mcc_h3_ = source23_.cf39
                d_1416___mcc_h4_ = source23_.cf40
                d_1417_cf40_ = d_1416___mcc_h4_
                d_1418_cf39_ = d_1415___mcc_h3_
                d_1419_cf38_ = d_1414___mcc_h2_
                d_1420_cf37_ = d_1413___mcc_h1_
                return len((_dafny.SeqWithoutIsStrInference([D3_DC9(906, len(_dafny.Set({d_1420_cf37_})))])) + (_dafny.SeqWithoutIsStrInference([D3_DC9(d_1419_cf38_, d_1418_cf39_)])))
            elif source23_.is_DC26:
                d_1421___mcc_h5_ = source23_.cf41
                d_1422_cf41_ = d_1421___mcc_h5_
                return (pat_let_tv30_) + (len(_dafny.Map({d_1422_cf41_: len(_dafny.Map({(self).f15: (self).f15}))})))
            elif source23_.is_DC23:
                d_1423___mcc_h6_ = source23_.cf35
                d_1424_cf35_ = d_1423___mcc_h6_
                return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tajrmdyiq")))
            elif True:
                d_1425___mcc_h7_ = source23_.cf42
                d_1426_cf42_ = d_1425___mcc_h7_
                return len((_dafny.Set({(self).f20, True, (self).f20, (self).f15, (self).f15})) - (_dafny.Set({(self).f20, (self).f20})))

        d_1408_v1_ = lambda76_(d_1410_v3_)
        hi9_ = d_1408_v1_
        for d_1427_i0_ in range(d_1408_v1_, hi9_):
            d_1428_v4_: str
            d_1428_v4_ = _dafny.CodePoint('k')
            d_1428_v4_ = d_1428_v4_
            d_1429_v5_: _dafny.Seq
            d_1429_v5_ = _dafny.SeqWithoutIsStrInference([d_1408_v1_])
            (globalState).f4 = ((d_1429_v5_) + ((d_1429_v5_).set(default__.safeIndex(-57, len(d_1429_v5_)), d_1408_v1_))) <= (d_1429_v5_)
            d_1409_v2_ = d_1409_v2_
            d_1430_v6_: _dafny.Set
            d_1430_v6_ = _dafny.Set({d_1427_i0_, d_1408_v1_})
            d_1431_v7_: _dafny.Seq
            d_1431_v7_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (d_1430_v6_).ispropersubset(d_1430_v6_), (self).f20, (default__.fm2((self).f15, globalState) if (self).f20 else (self).f15)])
            d_1431_v7_ = ((self).f14) + (d_1431_v7_)
        d_1432_v8_: _dafny.Map
        d_1432_v8_ = _dafny.Map({(0) - (d_1408_v1_): (self).f20})
        index190_ = default__.safeIndex(796, (d_1407_v0_).length(0))
        (d_1407_v0_)[index190_] = ((d_1432_v8_)[len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1433_i1_ in range(default__.abs(113))]))] if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1434_i1_ in range(default__.abs(113))]))) in (d_1432_v8_) else True)
        index191_ = default__.safeIndex(796, (d_1407_v0_).length(0))
        (d_1407_v0_)[index191_] = False
        (globalState).f4 = (default__.safeModuloInt(407, 562)) < (d_1408_v1_)

    @property
    def f20(self):
        return self._f20

class C4(T4):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm10(self, p0, p1, p2, p3, globalState):
        def iife107_():
            def iife109_():
                coll55_ = _dafny.Map()
                compr_55_: int
                for compr_55_ in _dafny.IntegerRange(914, 952):
                    d_1435_v0_: int = compr_55_
                    if ((914) <= (d_1435_v0_)) and ((d_1435_v0_) < (952)):
                        coll55_[default__.safeDivisionInt(d_1435_v0_, -486)] = True
                return _dafny.Map(coll55_)
            coll53_ = _dafny.Set()
            def iife108_():
                coll54_ = _dafny.Map()
                compr_54_: int
                for compr_54_ in _dafny.IntegerRange(914, 952):
                    d_1435_v0_: int = compr_54_
                    if ((914) <= (d_1435_v0_)) and ((d_1435_v0_) < (952)):
                        coll54_[default__.safeDivisionInt(d_1435_v0_, -486)] = True
                return _dafny.Map(coll54_)
            compr_53_: int
            for compr_53_ in ((_dafny.Map({495: False})) | (iife108_()
            )).keys.Elements:
                d_1436_v1_: int = compr_53_
                if (d_1436_v1_) in ((_dafny.Map({495: False})) | (iife109_()
                )):
                    coll53_ = coll53_.union(_dafny.Set([(d_1436_v1_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))]))
            return _dafny.Set(coll53_)
        return iife107_()
        

    def fm26(self, p0, p1, p2, globalState):
        return (D5_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False)]))]))).cf25

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        d_1437_v0_: str
        d_1437_v0_ = _dafny.CodePoint('s')
        d_1438_v1_: _dafny.Set
        d_1438_v1_ = _dafny.Set({d_1437_v0_})
        d_1439_v2_: _dafny.Map
        d_1439_v2_ = _dafny.Map({len(d_1438_v1_): (_dafny.MultiSet([d_1437_v0_])).cardinality})
        d_1440_v3_: bool
        d_1440_v3_ = False
        d_1441_v4_: D0
        d_1441_v4_ = D0_DC2(p2, not(d_1440_v3_))
        if (((d_1439_v2_)[148] if (148) in (d_1439_v2_) else (d_1441_v4_).cf1)) == (((d_1439_v2_)[p2] if (p2) in (d_1439_v2_) else p2)):
            (globalState).f4 = d_1440_v3_
            d_1442_v5_: _dafny.Seq
            d_1442_v5_ = _dafny.SeqWithoutIsStrInference([p1])
            r0 = (p2) < (default__.safeModuloInt(p2, len(d_1442_v5_)))
            d_1443_v6_: _dafny.Map
            d_1443_v6_ = _dafny.Map({d_1440_v3_: p2})
            d_1444_v7_: _dafny.Map
            d_1444_v7_ = _dafny.Map({p2: d_1443_v6_})
            d_1445_v8_: D3
            d_1445_v8_ = D3_DC9((default__.fm0(d_1440_v3_, d_1444_v7_, d_1440_v3_, globalState)) * ((0) - (p2)), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1446_i0_ in range(default__.abs(662))])))
            source24_ = d_1445_v8_
            if source24_.is_DC9:
                d_1447___mcc_h0_ = source24_.cf12
                d_1448___mcc_h1_ = source24_.cf13
                d_1449_cf13_ = d_1448___mcc_h1_
                d_1450_cf12_ = d_1447___mcc_h0_
                (globalState).f4 = default__.fm2(not(d_1440_v3_), globalState)
                d_1449_cf13_ = default__.fm0(True, d_1444_v7_, not(d_1440_v3_), globalState)
                d_1451_v9_: D6
                d_1451_v9_ = D6_DC18(p2)
                d_1451_v9_ = D6_DC18(default__.safeDivisionInt(d_1450_cf12_, d_1449_cf13_))
                (globalState).f4 = d_1440_v3_
            elif source24_.is_DC10:
                d_1452___mcc_h2_ = source24_.cf14
                d_1453___mcc_h3_ = source24_.cf15
                d_1454___mcc_h4_ = source24_.cf16
                d_1455_cf16_ = d_1454___mcc_h4_
                d_1456_cf15_ = d_1453___mcc_h3_
                d_1457_cf14_ = d_1452___mcc_h2_
                d_1457_cf14_ = ((d_1455_cf16_).f14)[default__.safeIndex((0) - (default__.safeDivisionInt(d_1456_cf15_, d_1456_cf15_)), len((d_1455_cf16_).f14))]
                d_1458_v10_: _dafny.Array
                def lambda77_(d_1459_v0_, d_1460_p1_):
                    def lambda78_(d_1461_i1_):
                        return (_dafny.SeqWithoutIsStrInference([d_1459_v0_ for d_1462_i2_ in range(default__.abs(220))])).set(default__.safeIndex(len(d_1460_p1_), len(_dafny.SeqWithoutIsStrInference([d_1459_v0_ for d_1463_i2_ in range(default__.abs(220))]))), _dafny.CodePoint('f'))

                    return lambda78_

                init38_ = lambda77_(d_1437_v0_, p1)
                nw216_ = _dafny.Array(None, 12)
                for i0_38_ in range(nw216_.length(0)):
                    nw216_[i0_38_] = init38_(i0_38_)
                d_1458_v10_ = nw216_
                index192_ = default__.safeIndex(237, (d_1458_v10_).length(0))
                (d_1458_v10_)[index192_] = p1
                index193_ = default__.safeIndex(237, (d_1458_v10_).length(0))
                (d_1458_v10_)[index193_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqnkr"))) + (p1)
                d_1464_v11_: C1
                nw217_ = C1()
                nw217_.ctor__(d_1457_cf14_, d_1440_v3_, (d_1455_cf16_).f13, ((d_1455_cf16_).f14 if d_1440_v3_ else (d_1455_cf16_).f14), d_1455_cf16_.f12)
                d_1464_v11_ = nw217_
                d_1464_v11_ = d_1464_v11_
                d_1465_v12_: bool
                d_1466_v13_: _dafny.Seq
                d_1467_v14_: _dafny.Seq
                out93_: bool
                out94_: _dafny.Seq
                out95_: _dafny.Seq
                out93_, out94_, out95_ = (self).m8(globalState)
                d_1465_v12_ = out93_
                d_1466_v13_ = out94_
                d_1467_v14_ = out95_
            elif source24_.is_DC8:
                d_1468___mcc_h5_ = source24_.cf11
                d_1469_cf11_ = d_1468___mcc_h5_
                d_1470_v15_: _dafny.Map
                d_1470_v15_ = _dafny.Map({p2: d_1440_v3_})
                d_1471_v16_: _dafny.MultiSet
                d_1471_v16_ = _dafny.MultiSet([p2, p2])
                d_1472_v17_: _dafny.Seq
                d_1472_v17_ = _dafny.SeqWithoutIsStrInference([d_1471_v16_, d_1471_v16_, d_1471_v16_, d_1471_v16_, (d_1471_v16_).set(955, default__.abs(p2))])
                d_1473_v18_: _dafny.Map
                d_1473_v18_ = _dafny.Map({d_1470_v15_: ((d_1472_v17_)[default__.safeIndex((0) - (p2), len(d_1472_v17_))]).cardinality})
                d_1474_v19_: _dafny.Array
                nw218_ = _dafny.Array(None, 19)
                nw218_[int(0)] = p2
                nw218_[int(1)] = p2
                nw218_[int(2)] = len((p0).intersection(default__.fm27(globalState)))
                nw218_[int(3)] = p2
                nw218_[int(4)] = p2
                nw218_[int(5)] = p2
                nw218_[int(6)] = p2
                nw218_[int(7)] = ((d_1443_v6_)[d_1440_v3_] if (d_1440_v3_) in (d_1443_v6_) else p2)
                nw218_[int(8)] = p2
                nw218_[int(9)] = (p2) - (p2)
                nw218_[int(10)] = (656) + (default__.fm0(d_1440_v3_, _dafny.Map({(0) - ((0) - (p2)): d_1443_v6_}), d_1440_v3_, globalState))
                nw218_[int(11)] = p2
                nw218_[int(12)] = p2
                nw218_[int(13)] = p2
                nw218_[int(14)] = p2
                nw218_[int(15)] = ((d_1473_v18_)[d_1470_v15_] if (d_1470_v15_) in (d_1473_v18_) else p2)
                nw218_[int(16)] = p2
                nw218_[int(17)] = len(d_1469_cf11_)
                nw218_[int(18)] = p2
                d_1474_v19_ = nw218_
                index194_ = default__.safeIndex(157, (d_1474_v19_).length(0))
                (d_1474_v19_)[index194_] = p2
                index195_ = default__.safeIndex(157, (d_1474_v19_).length(0))
                (d_1474_v19_)[index195_] = p2
                d_1475_v20_: _dafny.Seq
                d_1475_v20_ = _dafny.SeqWithoutIsStrInference([d_1440_v3_])
                index196_ = default__.safeIndex(157, (d_1474_v19_).length(0))
                rhs171_ = d_1440_v3_
                rhs172_ = (((d_1474_v19_)[default__.safeIndex(157, (d_1474_v19_).length(0))]) - (p2)) - ((len(d_1475_v20_)) * (p2))
                lhs97_ = d_1474_v19_
                lhs98_ = default__.safeIndex(157, (d_1474_v19_).length(0))
                d_1440_v3_ = rhs171_
                lhs97_[lhs98_] = rhs172_
                d_1476_v21_: _dafny.Map
                d_1476_v21_ = _dafny.Map({d_1440_v3_: d_1469_cf11_})
                d_1477_v22_: _dafny.Array
                nw219_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_1477_v22_ = nw219_
                d_1478_v23_: C1
                nw220_ = C1()
                nw220_.ctor__(d_1440_v3_, d_1440_v3_, d_1476_v21_, _dafny.SeqWithoutIsStrInference([d_1440_v3_, not(d_1440_v3_)]), d_1477_v22_)
                d_1478_v23_ = nw220_
                d_1479_v24_: C0
                nw221_ = C0()
                nw221_.ctor__()
                d_1479_v24_ = nw221_
                d_1480_v25_: _dafny.Map
                d_1480_v25_ = _dafny.Map({d_1479_v24_: default__.fm2(True, globalState)})
                d_1481_v26_: _dafny.Array
                nw222_ = _dafny.Array(None, 15)
                nw222_[int(0)] = d_1478_v23_.f18
                nw222_[int(1)] = d_1440_v3_
                nw222_[int(2)] = d_1478_v23_.f18
                nw222_[int(3)] = ((d_1480_v25_)[d_1479_v24_] if (d_1479_v24_) in (d_1480_v25_) else d_1478_v23_.f18)
                nw222_[int(4)] = not((d_1478_v23_).f17)
                nw222_[int(5)] = d_1478_v23_.f18
                nw222_[int(6)] = d_1440_v3_
                nw222_[int(7)] = d_1478_v23_.f18
                nw222_[int(8)] = d_1440_v3_
                nw222_[int(9)] = True
                nw222_[int(10)] = (d_1478_v23_).f17
                nw222_[int(11)] = d_1478_v23_.f18
                nw222_[int(12)] = (d_1478_v23_).f17
                nw222_[int(13)] = (d_1478_v23_).f17
                nw222_[int(14)] = (d_1478_v23_).f17
                d_1481_v26_ = nw222_
                d_1482_v27_: _dafny.Map
                d_1482_v27_ = _dafny.Map({d_1481_v26_: d_1478_v23_.f18})
                index197_ = default__.safeIndex(157, (d_1474_v19_).length(0))
                (d_1474_v19_)[index197_] = default__.fm0((d_1440_v3_ if (d_1475_v20_)[default__.safeIndex(len(d_1482_v27_), len(d_1475_v20_))] else (d_1478_v23_).f17), d_1444_v7_, True, globalState)
            elif True:
                d_1483___mcc_h6_ = source24_.cf17
                d_1484_cf17_ = d_1483___mcc_h6_
                d_1485_v28_: int
                d_1485_v28_ = -781
                d_1485_v28_ = (0) - (p2)
                d_1485_v28_ = len(default__.fm27(globalState))
                d_1486_v29_: _dafny.Set
                d_1486_v29_ = _dafny.Set({p1})
                (globalState).f4 = not(((d_1486_v29_) | (d_1486_v29_)).issubset(d_1486_v29_))
                d_1487_v30_: _dafny.Array
                nw223_ = _dafny.Array(int(0), 20)
                d_1487_v30_ = nw223_
                d_1488_v31_: _dafny.Map
                d_1488_v31_ = _dafny.Map({d_1487_v30_: d_1485_v28_})
                d_1489_v32_: _dafny.MultiSet
                d_1489_v32_ = _dafny.MultiSet([p2])
                d_1490_v33_: _dafny.Array
                nw224_ = _dafny.Array(D5.default()(), 23)
                d_1490_v33_ = nw224_
                d_1491_v34_: _dafny.Map
                d_1491_v34_ = _dafny.Map({len(p1): d_1488_v31_})
                rhs173_ = (_dafny.Map({d_1487_v30_: d_1485_v28_})) | (((d_1491_v34_)[p2] if (p2) in (d_1491_v34_) else d_1488_v31_))
                rhs174_ = (d_1489_v32_ if (-977) <= (d_1485_v28_) else d_1489_v32_)
                rhs175_ = (d_1490_v33_ if (p2) > (p2) else d_1490_v33_)
                d_1488_v31_ = rhs173_
                d_1489_v32_ = rhs174_
                d_1490_v33_ = rhs175_
            d_1492_v35_: _dafny.Array
            nw225_ = _dafny.Array(_dafny.MultiSet({}), 18)
            d_1492_v35_ = nw225_
            index198_ = default__.safeIndex(236, (d_1492_v35_).length(0))
            (d_1492_v35_)[index198_] = default__.fm28(p1, d_1440_v3_, p2, globalState)
            d_1493_v36_: _dafny.MultiSet
            d_1493_v36_ = _dafny.MultiSet([p2])
            index199_ = default__.safeIndex(236, (d_1492_v35_).length(0))
            (d_1492_v35_)[index199_] = d_1493_v36_
            (globalState).f4 = d_1440_v3_
        elif True:
            d_1494_v37_: _dafny.Map
            d_1494_v37_ = _dafny.Map({not(d_1440_v3_): d_1440_v3_})
            d_1495_v38_: _dafny.Array
            nw226_ = _dafny.Array(None, 15)
            nw226_[int(0)] = d_1440_v3_
            nw226_[int(1)] = d_1440_v3_
            nw226_[int(2)] = d_1440_v3_
            nw226_[int(3)] = d_1440_v3_
            nw226_[int(4)] = d_1440_v3_
            nw226_[int(5)] = not(not(d_1440_v3_))
            nw226_[int(6)] = (d_1440_v3_) and (d_1440_v3_)
            nw226_[int(7)] = (True) not in (d_1494_v37_)
            nw226_[int(8)] = d_1440_v3_
            nw226_[int(9)] = d_1440_v3_
            nw226_[int(10)] = (p1) <= (p1)
            nw226_[int(11)] = d_1440_v3_
            nw226_[int(12)] = d_1440_v3_
            nw226_[int(13)] = not (d_1440_v3_) or (d_1440_v3_)
            nw226_[int(14)] = d_1440_v3_
            d_1495_v38_ = nw226_
            index200_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            (d_1495_v38_)[index200_] = d_1440_v3_
            d_1496_v39_: _dafny.MultiSet
            d_1496_v39_ = _dafny.MultiSet([d_1441_v4_])
            index201_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            rhs176_ = not (((d_1494_v37_)[d_1440_v3_] if (d_1440_v3_) in (d_1494_v37_) else False)) or (True)
            rhs177_ = d_1440_v3_
            rhs178_ = d_1496_v39_
            lhs99_ = globalState
            lhs100_ = d_1495_v38_
            lhs101_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            lhs99_.f4 = rhs176_
            lhs100_[lhs101_] = rhs177_
            d_1496_v39_ = rhs178_
            index202_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            (d_1495_v38_)[index202_] = (d_1495_v38_)[default__.safeIndex(731, (d_1495_v38_).length(0))]
            index203_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            (d_1495_v38_)[index203_] = not (d_1440_v3_) or ((d_1495_v38_)[default__.safeIndex(731, (d_1495_v38_).length(0))])
            d_1497_v40_: _dafny.Array
            nw227_ = _dafny.Array(D2.default()(), 17)
            d_1497_v40_ = nw227_
            d_1498_v41_: _dafny.Map
            d_1498_v41_ = _dafny.Map({d_1497_v40_: p2})
            d_1499_v42_: _dafny.Seq
            d_1499_v42_ = _dafny.SeqWithoutIsStrInference([((d_1498_v41_)[d_1497_v40_] if (d_1497_v40_) in (d_1498_v41_) else p2)])
            d_1499_v42_ = (d_1499_v42_) + (_dafny.SeqWithoutIsStrInference([77]))
            index204_ = default__.safeIndex(731, (d_1495_v38_).length(0))
            (d_1495_v38_)[index204_] = (d_1495_v38_)[default__.safeIndex(731, (d_1495_v38_).length(0))]
        d_1500_v43_: int
        d_1500_v43_ = 584
        d_1501_v44_: _dafny.Seq
        d_1501_v44_ = _dafny.SeqWithoutIsStrInference([d_1440_v3_])
        d_1500_v43_ = ((d_1500_v43_) + (len(d_1501_v44_))) + (d_1500_v43_)
        d_1500_v43_ = (0) - ((p2) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1502_i3_ in range(default__.abs(607))])), p2)))
        d_1503_v45_: D0
        d_1503_v45_ = D0_DC1()
        d_1503_v45_ = d_1503_v45_
        hi10_ = (p2) * (d_1500_v43_)
        for d_1504_i4_ in range((p2) * (p2), hi10_):
            d_1500_v43_ = (default__.safeModuloInt(p2, d_1500_v43_)) + (d_1504_i4_)
            d_1500_v43_ = ((375) - ((0) - (-426))) + (-117)
            d_1505_v46_: _dafny.Array
            def lambda79_(d_1506_v3_):
                def lambda80_(d_1507_i5_):
                    return (_dafny.MultiSet([d_1506_v3_, d_1506_v3_])) == (_dafny.MultiSet([not(d_1506_v3_)]))

                return lambda80_

            init39_ = lambda79_(d_1440_v3_)
            nw228_ = _dafny.Array(None, 6)
            for i0_39_ in range(nw228_.length(0)):
                nw228_[i0_39_] = init39_(i0_39_)
            d_1505_v46_ = nw228_
            d_1505_v46_ = d_1505_v46_
            d_1440_v3_ = d_1440_v3_
        d_1508_v47_: _dafny.Array
        def lambda81_(d_1509_p2_):
            def lambda82_(d_1510_i6_):
                return default__.safeDivisionInt(d_1510_i6_, d_1509_p2_)

            return lambda82_

        init40_ = lambda81_(p2)
        nw229_ = _dafny.Array(None, 18)
        for i0_40_ in range(nw229_.length(0)):
            nw229_[i0_40_] = init40_(i0_40_)
        d_1508_v47_ = nw229_
        index205_ = default__.safeIndex(852, (d_1508_v47_).length(0))
        (d_1508_v47_)[index205_] = len(p0)
        d_1511_v48_: _dafny.Map
        d_1511_v48_ = _dafny.Map({d_1437_v0_: d_1440_v3_})
        index206_ = default__.safeIndex(852, (d_1508_v47_).length(0))
        (d_1508_v47_)[index206_] = (len((p1).set(default__.safeIndex(len(p0), len(p1)), d_1437_v0_))) * (len(d_1511_v48_))
        r0 = d_1440_v3_
        return r0

    def m6(self, globalState):
        r0: bool = False
        d_1512_v0_: bool
        d_1512_v0_ = False
        d_1513_v1_: _dafny.Seq
        d_1513_v1_ = _dafny.SeqWithoutIsStrInference([d_1512_v0_])
        d_1514_v2_: _dafny.Seq
        d_1514_v2_ = _dafny.SeqWithoutIsStrInference([d_1513_v1_])
        d_1515_v3_: D2
        d_1515_v3_ = D2_DC4(d_1514_v2_)
        d_1516_v4_: D2
        d_1516_v4_ = D2_DC7(d_1515_v3_)
        d_1517_v5_: D2
        d_1517_v5_ = D2_DC7(d_1516_v4_)
        d_1518_v6_: _dafny.Map
        d_1518_v6_ = _dafny.Map({d_1512_v0_: d_1517_v5_})
        d_1518_v6_ = (d_1518_v6_).set(d_1512_v0_, d_1517_v5_)
        if d_1512_v0_:
            d_1519_v7_: _dafny.Array
            nw230_ = _dafny.Array(_dafny.Map({}), 15)
            d_1519_v7_ = nw230_
            d_1520_v8_: _dafny.Array
            nw231_ = _dafny.Array(None, 6)
            nw231_[int(0)] = d_1512_v0_
            nw231_[int(1)] = d_1512_v0_
            nw231_[int(2)] = d_1512_v0_
            nw231_[int(3)] = (default__.fm2(d_1512_v0_, globalState)) or (d_1512_v0_)
            nw231_[int(4)] = d_1512_v0_
            nw231_[int(5)] = (d_1512_v0_) and (d_1512_v0_)
            d_1520_v8_ = nw231_
            d_1521_v9_: int
            d_1521_v9_ = 774
            d_1522_v10_: D2
            d_1522_v10_ = D2_DC6(d_1520_v8_, d_1520_v8_)
            d_1523_v11_: D3
            d_1523_v11_ = D3_DC9(d_1521_v9_, len(d_1513_v1_))
            rhs179_ = d_1519_v7_
            rhs180_ = (d_1522_v10_).cf8
            rhs181_ = (0) - ((d_1523_v11_).cf12)
            d_1519_v7_ = rhs179_
            d_1520_v8_ = rhs180_
            d_1521_v9_ = rhs181_
            d_1524_v12_: _dafny.Seq
            d_1524_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcj"))
            d_1525_v13_: _dafny.Map
            d_1525_v13_ = _dafny.Map({d_1512_v0_: d_1512_v0_})
            d_1526_v14_: _dafny.Map
            d_1526_v14_ = _dafny.Map({d_1512_v0_: d_1525_v13_})
            d_1527_v15_: _dafny.Array
            nw232_ = _dafny.Array(_dafny.MultiSet({}), 12)
            d_1527_v15_ = nw232_
            d_1528_v16_: T1
            nw233_ = C1()
            nw233_.ctor__(d_1512_v0_, d_1512_v0_, d_1526_v14_, d_1513_v1_, d_1527_v15_)
            d_1528_v16_ = nw233_
            d_1529_v17_: D3
            d_1529_v17_ = D3_DC10(True, 194, d_1528_v16_)
            d_1530_v18_: str
            d_1530_v18_ = _dafny.CodePoint('d')
            d_1531_v19_: _dafny.Array
            nw234_ = _dafny.Array(None, 9)
            nw234_[int(0)] = d_1524_v12_
            nw234_[int(1)] = d_1524_v12_
            nw234_[int(2)] = (default__.fm29((d_1529_v17_).cf14, d_1512_v0_, d_1512_v0_, d_1521_v9_, globalState)) + (d_1524_v12_)
            nw234_[int(3)] = (d_1524_v12_).set(default__.safeIndex(d_1521_v9_, len(d_1524_v12_)), d_1530_v18_)
            nw234_[int(4)] = d_1524_v12_
            nw234_[int(5)] = (d_1524_v12_) + (d_1524_v12_)
            nw234_[int(6)] = d_1524_v12_
            nw234_[int(7)] = (d_1524_v12_) + (d_1524_v12_)
            nw234_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pysu"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndnibmluf")))
            d_1531_v19_ = nw234_
            index207_ = default__.safeIndex(325, (d_1531_v19_).length(0))
            (d_1531_v19_)[index207_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxcb"))
            index208_ = default__.safeIndex(325, (d_1531_v19_).length(0))
            (d_1531_v19_)[index208_] = default__.fm29(d_1512_v0_, default__.fm2(True, globalState), d_1512_v0_, d_1521_v9_, globalState)
            d_1532_v20_: _dafny.MultiSet
            d_1532_v20_ = _dafny.MultiSet([d_1512_v0_, d_1512_v0_, True])
            d_1521_v9_ = default__.safeModuloInt((d_1521_v9_) - ((d_1532_v20_).cardinality), d_1521_v9_)
            if (d_1521_v9_) <= (d_1521_v9_):
                d_1533_v21_: C0
                nw235_ = C0()
                nw235_.ctor__()
                d_1533_v21_ = nw235_
                d_1520_v8_ = d_1520_v8_
                d_1534_v22_: _dafny.Array
                nw236_ = _dafny.Array(int(0), 29)
                d_1534_v22_ = nw236_
                index209_ = default__.safeIndex(599, (d_1534_v22_).length(0))
                (d_1534_v22_)[index209_] = d_1521_v9_
                d_1535_v23_: _dafny.Set
                d_1535_v23_ = _dafny.Set({d_1512_v0_})
                index210_ = default__.safeIndex(599, (d_1534_v22_).length(0))
                (d_1534_v22_)[index210_] = (d_1521_v9_) + (default__.safeModuloInt(len(d_1535_v23_), d_1521_v9_))
                d_1536_v24_: _dafny.Seq
                d_1536_v24_ = _dafny.SeqWithoutIsStrInference([d_1520_v8_, d_1520_v8_])
                index211_ = default__.safeIndex(599, (d_1534_v22_).length(0))
                (d_1534_v22_)[index211_] = len((d_1536_v24_) + ((_dafny.SeqWithoutIsStrInference([d_1520_v8_, d_1520_v8_])) + (d_1536_v24_)))
                d_1530_v18_ = d_1530_v18_
            elif True:
                d_1537_v25_: _dafny.Seq
                d_1537_v25_ = _dafny.SeqWithoutIsStrInference([(d_1521_v9_ if d_1512_v0_ else (d_1528_v16_).fm5(d_1512_v0_, globalState)), d_1521_v9_])
                d_1521_v9_ = (d_1537_v25_)[default__.safeIndex((d_1521_v9_) + (len((d_1531_v19_)[default__.safeIndex(325, (d_1531_v19_).length(0))])), len(d_1537_v25_))]
                d_1538_v26_: _dafny.Array
                nw237_ = _dafny.Array(int(0), 9)
                d_1538_v26_ = nw237_
                d_1539_v27_: D2
                d_1539_v27_ = D2_DC5(len(d_1524_v12_), d_1512_v0_, not(d_1512_v0_))
                index212_ = default__.safeIndex(544, (d_1538_v26_).length(0))
                def iife110_():
                    coll56_ = _dafny.Map()
                    compr_56_: int
                    for compr_56_ in (_dafny.SeqWithoutIsStrInference([d_1521_v9_])).Elements:
                        d_1540_v28_: int = compr_56_
                        if (d_1540_v28_) in (_dafny.SeqWithoutIsStrInference([d_1521_v9_])):
                            coll56_[(d_1540_v28_) * (d_1521_v9_)] = _dafny.MultiSet([283])
                    return _dafny.Map(coll56_)
                (d_1538_v26_)[index212_] = ((d_1539_v27_).cf5) * (len(iife110_()
                ))
                index213_ = default__.safeIndex(544, (d_1538_v26_).length(0))
                (d_1538_v26_)[index213_] = 973
                d_1541_v29_: C0
                nw238_ = C0()
                nw238_.ctor__()
                d_1541_v29_ = nw238_
                index214_ = default__.safeIndex(544, (d_1538_v26_).length(0))
                (d_1538_v26_)[index214_] = (d_1538_v26_)[default__.safeIndex(544, (d_1538_v26_).length(0))]
                d_1542_v30_: _dafny.Map
                d_1542_v30_ = _dafny.Map({d_1521_v9_: (d_1520_v8_ if d_1512_v0_ else d_1520_v8_)})
                d_1542_v30_ = ((d_1542_v30_) | (d_1542_v30_)) | ((_dafny.Map({len((d_1531_v19_)[default__.safeIndex(325, (d_1531_v19_).length(0))]): d_1520_v8_})).set(len((d_1537_v25_).set(default__.safeIndex(len(d_1537_v25_), len(d_1537_v25_)), d_1521_v9_)), d_1520_v8_))
            d_1543_v31_: _dafny.Set
            d_1543_v31_ = _dafny.Set({d_1521_v9_, len(_dafny.Map({not(d_1512_v0_): d_1525_v13_})), d_1521_v9_, d_1521_v9_})
            d_1544_v32_: _dafny.Map
            d_1544_v32_ = _dafny.Map({d_1543_v31_: len(d_1524_v12_)})
            d_1545_v33_: _dafny.Map
            d_1545_v33_ = _dafny.Map({d_1512_v0_: (d_1544_v32_).set(_dafny.Set({180}), (0) - (d_1521_v9_))})
            d_1545_v33_ = (d_1545_v33_).set(d_1512_v0_, d_1544_v32_)
        elif True:
            d_1546_v34_: _dafny.Seq
            d_1546_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbusqdbk"))
            d_1547_v35_: _dafny.Map
            d_1547_v35_ = _dafny.Map({d_1546_v34_: 993})
            d_1548_v36_: _dafny.Array
            nw239_ = _dafny.Array(False, 21)
            d_1548_v36_ = nw239_
            d_1549_v37_: D2
            d_1549_v37_ = D2_DC6(d_1548_v36_, d_1548_v36_)
            d_1550_v38_: _dafny.Map
            d_1550_v38_ = _dafny.Map({(d_1547_v35_) | (d_1547_v35_): d_1549_v37_})
            d_1550_v38_ = (d_1550_v38_).set(d_1547_v35_, D2_DC6(d_1548_v36_, d_1548_v36_))
            d_1551_v39_: int
            d_1551_v39_ = 153
            d_1552_v40_: _dafny.Map
            d_1552_v40_ = _dafny.Map({d_1551_v39_: d_1512_v0_})
            r0 = ((d_1552_v40_)[d_1551_v39_] if (d_1551_v39_) in (d_1552_v40_) else d_1512_v0_)
            if ((len(d_1546_v34_)) + (d_1551_v39_)) < (len(_dafny.SeqWithoutIsStrInference([(0) - (d_1551_v39_) for d_1553_i0_ in range(default__.abs(299))]))):
                d_1554_v41_: _dafny.Map
                d_1554_v41_ = _dafny.Map({d_1512_v0_: d_1551_v39_})
                d_1555_v42_: _dafny.Map
                d_1555_v42_ = _dafny.Map({d_1551_v39_: d_1554_v41_})
                d_1556_v43_: _dafny.Map
                d_1556_v43_ = _dafny.Map({d_1512_v0_: d_1512_v0_})
                d_1557_v44_: D3
                d_1557_v44_ = D3_DC8(d_1556_v43_)
                d_1558_v45_: _dafny.Map
                d_1558_v45_ = _dafny.Map({d_1512_v0_: (d_1557_v44_).cf11})
                d_1559_v46_: _dafny.MultiSet
                d_1559_v46_ = _dafny.MultiSet([len(d_1554_v41_), d_1551_v39_])
                d_1560_v47_: _dafny.Seq
                d_1560_v47_ = _dafny.SeqWithoutIsStrInference([d_1551_v39_, (0) - (d_1551_v39_)])
                d_1561_v48_: _dafny.Array
                nw240_ = _dafny.Array(None, 15)
                nw240_[int(0)] = default__.fm28(d_1546_v34_, d_1512_v0_, d_1551_v39_, globalState)
                nw240_[int(1)] = d_1559_v46_
                nw240_[int(2)] = d_1559_v46_
                nw240_[int(3)] = d_1559_v46_
                nw240_[int(4)] = d_1559_v46_
                nw240_[int(5)] = d_1559_v46_
                nw240_[int(6)] = _dafny.MultiSet(d_1560_v47_)
                nw240_[int(7)] = d_1559_v46_
                nw240_[int(8)] = d_1559_v46_
                nw240_[int(9)] = _dafny.MultiSet([-746, d_1551_v39_, d_1551_v39_])
                nw240_[int(10)] = d_1559_v46_
                nw240_[int(11)] = d_1559_v46_
                nw240_[int(12)] = d_1559_v46_
                nw240_[int(13)] = _dafny.MultiSet([d_1551_v39_])
                nw240_[int(14)] = d_1559_v46_
                d_1561_v48_ = nw240_
                d_1562_v49_: C1
                nw241_ = C1()
                nw241_.ctor__(d_1512_v0_, d_1512_v0_, d_1558_v45_, d_1513_v1_, d_1561_v48_)
                d_1562_v49_ = nw241_
                d_1563_v50_: _dafny.Map
                d_1563_v50_ = _dafny.Map({d_1562_v49_: ((d_1556_v43_)[(d_1562_v49_).f17] if ((d_1562_v49_).f17) in (d_1556_v43_) else d_1562_v49_.f18)})
                d_1564_v51_: _dafny.Set
                d_1564_v51_ = _dafny.Set({default__.fm0(d_1512_v0_, (d_1555_v42_).set(len(d_1563_v50_), d_1554_v41_), (d_1562_v49_).f17, globalState), (0) - ((0) - (d_1551_v39_)), d_1551_v39_, d_1551_v39_})
                def iife111_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in (d_1552_v40_).keys.Elements:
                        d_1565_v52_: int = compr_57_
                        if (d_1565_v52_) in (d_1552_v40_):
                            coll57_ = coll57_.union(_dafny.Set([default__.safeDivisionInt(d_1565_v52_, (_dafny.MultiSet([True, not(True)])).cardinality)]))
                    return _dafny.Set(coll57_)
                rhs182_ = ((iife111_()
                ).intersection(d_1564_v51_)) - (d_1564_v51_)
                rhs183_ = d_1548_v36_
                rhs184_ = d_1548_v36_
                rhs185_ = (d_1551_v39_) * ((0) - (default__.safeDivisionInt(-947, d_1551_v39_)))
                d_1564_v51_ = rhs182_
                d_1548_v36_ = rhs183_
                d_1548_v36_ = rhs184_
                d_1551_v39_ = rhs185_
                (globalState).f4 = (d_1562_v49_).f17
                d_1566_v53_: str
                d_1566_v53_ = _dafny.CodePoint('s')
                d_1567_v54_: _dafny.Array
                nw242_ = _dafny.Array(None, 24)
                nw242_[int(0)] = d_1546_v34_
                nw242_[int(1)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di"))) + ((d_1546_v34_).set(default__.safeIndex(-842, len(d_1546_v34_)), d_1566_v53_))).set(default__.safeIndex(d_1551_v39_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "di"))) + ((d_1546_v34_).set(default__.safeIndex(-842, len(d_1546_v34_)), d_1566_v53_)))), d_1566_v53_)
                nw242_[int(2)] = d_1546_v34_
                nw242_[int(3)] = default__.fm29((d_1562_v49_).fm17(d_1560_v47_, (0) - ((d_1562_v49_).fm5(d_1562_v49_.f18, globalState)), globalState), (d_1562_v49_).f17, d_1562_v49_.f18, 305, globalState)
                nw242_[int(4)] = d_1546_v34_
                nw242_[int(5)] = d_1546_v34_
                nw242_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qafisv"))
                nw242_[int(7)] = d_1546_v34_
                nw242_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxdoh"))
                nw242_[int(9)] = d_1546_v34_
                nw242_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnfm"))
                nw242_[int(11)] = (d_1546_v34_).set(default__.safeIndex(d_1551_v39_, len(d_1546_v34_)), d_1566_v53_)
                nw242_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwfvvvjpx"))
                nw242_[int(13)] = (d_1546_v34_) + (d_1546_v34_)
                nw242_[int(14)] = (d_1546_v34_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pawwyk")))
                nw242_[int(15)] = d_1546_v34_
                nw242_[int(16)] = default__.fm29(d_1562_v49_.f18, d_1512_v0_, d_1512_v0_, (D3_DC9(d_1551_v39_, d_1551_v39_)).cf12, globalState)
                nw242_[int(17)] = (d_1546_v34_) + (d_1546_v34_)
                nw242_[int(18)] = d_1546_v34_
                nw242_[int(19)] = (d_1546_v34_) + (_dafny.SeqWithoutIsStrInference([d_1566_v53_ for d_1568_i1_ in range(default__.abs(-652))]))
                nw242_[int(20)] = d_1546_v34_
                nw242_[int(21)] = d_1546_v34_
                nw242_[int(22)] = (d_1546_v34_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqvycdvnb")))
                nw242_[int(23)] = d_1546_v34_
                d_1567_v54_ = nw242_
                index215_ = default__.safeIndex(982, (d_1567_v54_).length(0))
                (d_1567_v54_)[index215_] = d_1546_v34_
                d_1569_v55_: _dafny.Seq
                d_1569_v55_ = _dafny.SeqWithoutIsStrInference([d_1546_v34_])
                index216_ = default__.safeIndex(982, (d_1567_v54_).length(0))
                (d_1567_v54_)[index216_] = ((d_1569_v55_)[default__.safeIndex(d_1551_v39_, len(d_1569_v55_))]) + (d_1546_v34_)
                d_1570_v56_: _dafny.Map
                d_1570_v56_ = _dafny.Map({_dafny.MultiSet((d_1567_v54_)[default__.safeIndex(982, (d_1567_v54_).length(0))]): (d_1551_v39_) * (d_1551_v39_)})
                d_1571_v57_: _dafny.MultiSet
                d_1571_v57_ = _dafny.MultiSet([d_1566_v53_])
                d_1572_v58_: _dafny.Map
                d_1572_v58_ = _dafny.Map({d_1551_v39_: d_1551_v39_})
                d_1570_v56_ = (d_1570_v56_).set((d_1571_v57_) - ((d_1571_v57_).set(d_1566_v53_, default__.abs(-271))), ((d_1572_v58_)[len(d_1546_v34_)] if (len(d_1546_v34_)) in (d_1572_v58_) else d_1551_v39_))
                d_1551_v39_ = d_1551_v39_
            elif True:
                (globalState).f4 = (d_1512_v0_) and (d_1512_v0_)
                (globalState).f4 = (d_1512_v0_ if not (d_1512_v0_) or (d_1512_v0_) else d_1512_v0_)
                d_1573_v59_: C0
                nw243_ = C0()
                nw243_.ctor__()
                d_1573_v59_ = nw243_
                (globalState).f4 = d_1512_v0_
                d_1574_v60_: _dafny.Seq
                d_1574_v60_ = _dafny.SeqWithoutIsStrInference([d_1546_v34_])
                d_1575_v61_: _dafny.Seq
                d_1575_v61_ = _dafny.SeqWithoutIsStrInference([d_1551_v39_, d_1551_v39_])
                d_1576_v62_: _dafny.Map
                d_1576_v62_ = _dafny.Map({True: not(d_1512_v0_)})
                d_1577_v63_: _dafny.Map
                d_1577_v63_ = _dafny.Map({d_1512_v0_: d_1576_v62_})
                d_1578_v64_: _dafny.Array
                nw244_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_1578_v64_ = nw244_
                d_1579_v65_: C1
                nw245_ = C1()
                nw245_.ctor__((len((d_1574_v60_)[default__.safeIndex(d_1551_v39_, len(d_1574_v60_))])) == (d_1551_v39_), (((d_1547_v35_)[d_1546_v34_] if (d_1546_v34_) in (d_1547_v35_) else d_1551_v39_)) in (d_1575_v61_), d_1577_v63_, (d_1514_v2_)[default__.safeIndex(d_1551_v39_, len(d_1514_v2_))], d_1578_v64_)
                d_1579_v65_ = nw245_
                d_1579_v65_ = d_1579_v65_
            (globalState).f4 = d_1512_v0_
            d_1580_v66_: _dafny.Array
            nw246_ = _dafny.Array(_dafny.Set({}), 22)
            d_1580_v66_ = nw246_
            d_1581_v67_: _dafny.Set
            d_1581_v67_ = _dafny.Set({d_1551_v39_, d_1551_v39_})
            index217_ = default__.safeIndex(26, (d_1580_v66_).length(0))
            (d_1580_v66_)[index217_] = (d_1581_v67_).intersection(d_1581_v67_)
            index218_ = default__.safeIndex(26, (d_1580_v66_).length(0))
            (d_1580_v66_)[index218_] = d_1581_v67_
        d_1582_v68_: str
        d_1582_v68_ = _dafny.CodePoint('b')
        d_1582_v68_ = d_1582_v68_
        d_1583_v69_: int
        d_1583_v69_ = 475
        hi11_ = d_1583_v69_
        for d_1584_i2_ in range(d_1583_v69_, hi11_):
            d_1585_v70_: _dafny.Seq
            d_1585_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqbu"))
            d_1585_v70_ = (d_1585_v70_) + (d_1585_v70_)
            d_1586_v71_: _dafny.Map
            d_1586_v71_ = _dafny.Map({d_1584_i2_: d_1583_v69_})
            d_1587_v72_: _dafny.Map
            d_1587_v72_ = _dafny.Map({d_1512_v0_: ((d_1586_v71_)[d_1583_v69_] if (d_1583_v69_) in (d_1586_v71_) else d_1584_i2_)})
            d_1587_v72_ = (d_1587_v72_).set(d_1512_v0_, d_1584_i2_)
            d_1588_v73_: _dafny.Array
            nw247_ = _dafny.Array(_dafny.Map({}), 8)
            d_1588_v73_ = nw247_
            d_1588_v73_ = d_1588_v73_
            d_1589_v74_: _dafny.MultiSet
            d_1589_v74_ = _dafny.MultiSet([d_1583_v69_])
            d_1590_v76_: _dafny.MultiSet
            d_1590_v76_ = _dafny.MultiSet([d_1582_v68_])
            def iife112_():
                coll58_ = _dafny.Map()
                compr_58_: str
                for compr_58_ in ((d_1590_v76_).set(d_1582_v68_, default__.abs(d_1584_i2_))).Elements:
                    d_1591_v75_: str = compr_58_
                    if (d_1591_v75_) in ((d_1590_v76_).set(d_1582_v68_, default__.abs(d_1584_i2_))):
                        coll58_[d_1591_v75_] = d_1512_v0_
                return _dafny.Map(coll58_)
            if (d_1589_v74_).ispropersubset(_dafny.MultiSet([d_1583_v69_, d_1583_v69_, len(d_1585_v70_), d_1583_v69_, len(iife112_()
            )])):
                d_1592_v77_: _dafny.Map
                d_1592_v77_ = _dafny.Map({_dafny.CodePoint('a'): (d_1583_v69_) - (d_1584_i2_)})
                d_1593_v78_: D6
                d_1593_v78_ = D6_DC19(default__.fm30(d_1583_v69_, d_1583_v69_, globalState))
                rhs186_ = d_1592_v77_
                rhs187_ = d_1593_v78_
                rhs188_ = (((d_1587_v72_)[d_1512_v0_] if (d_1512_v0_) in (d_1587_v72_) else d_1584_i2_)) <= (d_1584_i2_)
                rhs189_ = d_1583_v69_
                rhs190_ = (d_1583_v69_) * (d_1584_i2_)
                d_1592_v77_ = rhs186_
                d_1593_v78_ = rhs187_
                r0 = rhs188_
                d_1583_v69_ = rhs189_
                d_1583_v69_ = rhs190_
                (globalState).f4 = (d_1513_v1_)[default__.safeIndex(d_1583_v69_, len(d_1513_v1_))]
                (globalState).f4 = d_1512_v0_
                r0 = True
                d_1587_v72_ = (d_1587_v72_).set(not (d_1512_v0_) or (default__.fm2(d_1512_v0_, globalState)), d_1583_v69_)
            elif True:
                d_1583_v69_ = d_1583_v69_
                d_1594_v79_: _dafny.Set
                d_1594_v79_ = _dafny.Set({d_1512_v0_})
                d_1583_v69_ = (len((d_1594_v79_) | (d_1594_v79_))) + (d_1583_v69_)
                d_1595_v80_: _dafny.Seq
                d_1596_v81_: bool
                d_1597_v82_: int
                d_1598_v83_: bool
                out96_: _dafny.Seq
                out97_: bool
                out98_: int
                out99_: bool
                out96_, out97_, out98_, out99_ = default__.m0((_dafny.Set({False, default__.fm2(d_1512_v0_, globalState)})).ispropersubset(d_1594_v79_), default__.fm31(d_1584_i2_, globalState), True, globalState)
                d_1595_v80_ = out96_
                d_1596_v81_ = out97_
                d_1597_v82_ = out98_
                d_1598_v83_ = out99_
                d_1599_v84_: _dafny.Map
                d_1599_v84_ = _dafny.Map({len(d_1585_v70_): d_1587_v72_})
                d_1597_v82_ = (0) - (default__.fm0(d_1596_v81_, d_1599_v84_, d_1512_v0_, globalState))
                d_1600_v85_: D7
                d_1600_v85_ = D7_DC20(_dafny.MultiSet([d_1584_i2_, d_1597_v82_, d_1583_v69_, 747]))
                d_1601_v86_: _dafny.Map
                d_1601_v86_ = _dafny.Map({d_1512_v0_: d_1598_v83_})
                d_1602_v87_: _dafny.Map
                d_1602_v87_ = _dafny.Map({d_1598_v83_: d_1601_v86_})
                d_1603_v88_: _dafny.Array
                nw248_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_1603_v88_ = nw248_
                d_1604_v89_: C1
                nw249_ = C1()
                nw249_.ctor__(d_1598_v83_, ((d_1600_v85_).cf30).issubset(default__.fm28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkodyaqi")), False, d_1583_v69_, globalState)), d_1602_v87_, d_1513_v1_, d_1603_v88_)
                d_1604_v89_ = nw249_
        d_1605_v90_: _dafny.Array
        nw250_ = _dafny.Array(_dafny.MultiSet({}), 3)
        d_1605_v90_ = nw250_
        d_1606_v91_: _dafny.Seq
        d_1606_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckfjn"))
        d_1607_v92_: _dafny.Map
        d_1607_v92_ = _dafny.Map({d_1512_v0_: d_1606_v91_})
        index219_ = default__.safeIndex(546, (d_1605_v90_).length(0))
        (d_1605_v90_)[index219_] = default__.fm28(((d_1607_v92_)[True] if (True) in (d_1607_v92_) else d_1606_v91_), True, d_1583_v69_, globalState)
        d_1608_v93_: _dafny.Seq
        d_1608_v93_ = _dafny.SeqWithoutIsStrInference([d_1583_v69_])
        d_1609_v94_: _dafny.Seq
        d_1609_v94_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1608_v93_)[default__.safeIndex(d_1583_v69_, len(d_1608_v93_))]), d_1583_v69_, d_1583_v69_, d_1583_v69_])
        d_1610_v95_: _dafny.MultiSet
        d_1610_v95_ = _dafny.MultiSet([len(d_1609_v94_), d_1583_v69_])
        index220_ = default__.safeIndex(546, (d_1605_v90_).length(0))
        (d_1605_v90_)[index220_] = d_1610_v95_
        if d_1512_v0_:
            d_1611_v96_: _dafny.MultiSet
            d_1611_v96_ = _dafny.MultiSet([True, d_1512_v0_])
            d_1612_v97_: _dafny.MultiSet
            d_1612_v97_ = _dafny.MultiSet([d_1611_v96_, d_1611_v96_])
            d_1613_v98_: _dafny.Set
            d_1613_v98_ = _dafny.Set({False})
            d_1614_v99_: _dafny.Array
            nw251_ = _dafny.Array(None, 16)
            nw251_[int(0)] = d_1583_v69_
            nw251_[int(1)] = (d_1583_v69_ if d_1512_v0_ else d_1583_v69_)
            nw251_[int(2)] = d_1583_v69_
            nw251_[int(3)] = default__.safeModuloInt(472, d_1583_v69_)
            nw251_[int(4)] = d_1583_v69_
            nw251_[int(5)] = -742
            nw251_[int(6)] = default__.safeDivisionInt(d_1583_v69_, d_1583_v69_)
            nw251_[int(7)] = d_1583_v69_
            nw251_[int(8)] = default__.safeModuloInt(-773, ((d_1612_v97_)[d_1611_v96_] if (d_1611_v96_) in (d_1612_v97_) else len(d_1613_v98_)))
            nw251_[int(9)] = d_1583_v69_
            nw251_[int(10)] = (d_1611_v96_).cardinality
            nw251_[int(11)] = d_1583_v69_
            nw251_[int(12)] = d_1583_v69_
            nw251_[int(13)] = 674
            nw251_[int(14)] = d_1583_v69_
            nw251_[int(15)] = d_1583_v69_
            d_1614_v99_ = nw251_
            d_1614_v99_ = d_1614_v99_
            (globalState).f4 = default__.fm2((d_1583_v69_) == (d_1583_v69_), globalState)
            d_1583_v69_ = -413
            d_1615_v100_: _dafny.Map
            d_1615_v100_ = _dafny.Map({(d_1513_v1_)[default__.safeIndex(d_1583_v69_, len(d_1513_v1_))]: d_1512_v0_})
            d_1583_v69_ = len(d_1615_v100_)
            d_1616_v101_: D3
            d_1616_v101_ = D3_DC8((d_1615_v100_) | ((_dafny.Map({d_1512_v0_: d_1512_v0_})).set(d_1512_v0_, d_1512_v0_)))
            source25_ = d_1616_v101_
            if source25_.is_DC9:
                d_1617___mcc_h0_ = source25_.cf12
                d_1618___mcc_h1_ = source25_.cf13
                d_1619_cf13_ = d_1618___mcc_h1_
                d_1620_cf12_ = d_1617___mcc_h0_
                d_1621_v102_: _dafny.MultiSet
                d_1621_v102_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1582_v68_ for d_1622_i3_ in range(default__.abs(-38))])])
                index221_ = default__.safeIndex(382, (d_1614_v99_).length(0))
                (d_1614_v99_)[index221_] = 881
                d_1623_v103_: _dafny.Map
                d_1623_v103_ = _dafny.Map({d_1620_cf12_: True})
                d_1624_v105_: _dafny.Seq
                def iife113_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in (d_1623_v103_).keys.Elements:
                        d_1625_v104_: int = compr_59_
                        if (d_1625_v104_) in (d_1623_v103_):
                            coll59_[(d_1625_v104_) * (d_1620_cf12_)] = d_1512_v0_
                    return _dafny.Map(coll59_)
                d_1624_v105_ = _dafny.SeqWithoutIsStrInference([d_1623_v103_, (iife113_()
                 if not(d_1512_v0_) else _dafny.Map({d_1619_cf13_: d_1512_v0_}))])
                index222_ = default__.safeIndex(382, (d_1614_v99_).length(0))
                rhs191_ = d_1621_v102_
                rhs192_ = d_1583_v69_
                rhs193_ = (d_1513_v1_)[default__.safeIndex((d_1609_v94_)[default__.safeIndex(d_1583_v69_, len(d_1609_v94_))], len(d_1513_v1_))]
                rhs194_ = len(d_1624_v105_)
                lhs102_ = d_1614_v99_
                lhs103_ = default__.safeIndex(382, (d_1614_v99_).length(0))
                lhs104_ = globalState
                d_1621_v102_ = rhs191_
                lhs102_[lhs103_] = rhs192_
                lhs104_.f4 = rhs193_
                d_1583_v69_ = rhs194_
                d_1626_v106_: C1
                nw252_ = C1()
                nw252_.ctor__((d_1608_v93_) < (d_1609_v94_), True, _dafny.Map({d_1512_v0_: d_1615_v100_}), (d_1513_v1_) + (d_1513_v1_), d_1605_v90_)
                d_1626_v106_ = nw252_
                d_1627_v107_: _dafny.Array
                def lambda83_(d_1628_v99_):
                    def lambda84_(d_1629_i4_):
                        return default__.safeDivisionInt(d_1629_i4_, (d_1628_v99_)[default__.safeIndex(382, (d_1628_v99_).length(0))])

                    return lambda84_

                init41_ = lambda83_(d_1614_v99_)
                nw253_ = _dafny.Array(None, 14)
                for i0_41_ in range(nw253_.length(0)):
                    nw253_[i0_41_] = init41_(i0_41_)
                d_1627_v107_ = nw253_
                d_1630_v108_: _dafny.Set
                d_1630_v108_ = _dafny.Set({(0) - (d_1619_cf13_), d_1619_cf13_, len(d_1606_v91_), len(d_1606_v91_), d_1583_v69_})
                nw254_ = _dafny.Array(None, 7)
                nw254_[int(0)] = ((0) - (d_1583_v69_)) * (len(d_1609_v94_))
                nw254_[int(1)] = len(d_1630_v108_)
                nw254_[int(2)] = d_1583_v69_
                nw254_[int(3)] = d_1583_v69_
                nw254_[int(4)] = d_1619_cf13_
                nw254_[int(5)] = d_1619_cf13_
                nw254_[int(6)] = default__.safeModuloInt((d_1614_v99_)[default__.safeIndex(382, (d_1614_v99_).length(0))], (d_1614_v99_)[default__.safeIndex(382, (d_1614_v99_).length(0))])
                d_1627_v107_ = nw254_
                d_1583_v69_ = d_1620_cf12_
            elif source25_.is_DC10:
                d_1631___mcc_h2_ = source25_.cf14
                d_1632___mcc_h3_ = source25_.cf15
                d_1633___mcc_h4_ = source25_.cf16
                d_1634_cf16_ = d_1633___mcc_h4_
                d_1635_cf15_ = d_1632___mcc_h3_
                d_1636_cf14_ = d_1631___mcc_h2_
                d_1512_v0_ = d_1636_cf14_
                d_1637_v109_: _dafny.Array
                nw255_ = _dafny.Array(None, 23)
                nw255_[int(0)] = d_1614_v99_
                nw255_[int(1)] = d_1614_v99_
                nw255_[int(2)] = d_1614_v99_
                nw255_[int(3)] = d_1614_v99_
                nw255_[int(4)] = d_1614_v99_
                nw255_[int(5)] = d_1614_v99_
                nw255_[int(6)] = d_1614_v99_
                nw255_[int(7)] = d_1614_v99_
                nw255_[int(8)] = d_1614_v99_
                nw255_[int(9)] = d_1614_v99_
                nw255_[int(10)] = d_1614_v99_
                nw255_[int(11)] = d_1614_v99_
                nw255_[int(12)] = d_1614_v99_
                nw255_[int(13)] = d_1614_v99_
                nw255_[int(14)] = (d_1614_v99_ if default__.fm2(d_1512_v0_, globalState) else d_1614_v99_)
                nw255_[int(15)] = d_1614_v99_
                nw255_[int(16)] = d_1614_v99_
                nw255_[int(17)] = d_1614_v99_
                nw255_[int(18)] = d_1614_v99_
                nw255_[int(19)] = d_1614_v99_
                nw255_[int(20)] = d_1614_v99_
                nw255_[int(21)] = d_1614_v99_
                nw255_[int(22)] = d_1614_v99_
                d_1637_v109_ = nw255_
                d_1638_v110_: _dafny.Array
                nw256_ = _dafny.Array(False, 17)
                d_1638_v110_ = nw256_
                index223_ = default__.safeIndex(319, (d_1638_v110_).length(0))
                (d_1638_v110_)[index223_] = not((d_1583_v69_) < (d_1635_cf15_))
                index224_ = default__.safeIndex(572, (d_1614_v99_).length(0))
                (d_1614_v99_)[index224_] = default__.safeDivisionInt(d_1635_cf15_, d_1583_v69_)
                d_1639_v112_: _dafny.Map
                def iife114_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(409, 735):
                        d_1640_v111_: int = compr_60_
                        if ((409) <= (d_1640_v111_)) and ((d_1640_v111_) < (735)):
                            coll60_[default__.safeDivisionInt(d_1640_v111_, len(_dafny.Map({d_1636_cf14_: d_1583_v69_})))] = d_1512_v0_
                    return _dafny.Map(coll60_)
                d_1639_v112_ = _dafny.Map({d_1635_cf15_: len(iife114_()
                )})
                index225_ = default__.safeIndex(319, (d_1638_v110_).length(0))
                index226_ = default__.safeIndex(572, (d_1614_v99_).length(0))
                rhs195_ = d_1637_v109_
                rhs196_ = d_1635_cf15_
                rhs197_ = (d_1583_v69_) <= (d_1583_v69_)
                rhs198_ = (len(d_1639_v112_)) - ((0) - ((0) - ((d_1635_cf15_) + (d_1635_cf15_))))
                rhs199_ = ((704) - (d_1635_cf15_)) - (d_1635_cf15_)
                lhs105_ = d_1638_v110_
                lhs106_ = default__.safeIndex(319, (d_1638_v110_).length(0))
                lhs107_ = d_1614_v99_
                lhs108_ = default__.safeIndex(572, (d_1614_v99_).length(0))
                d_1637_v109_ = rhs195_
                d_1635_cf15_ = rhs196_
                lhs105_[lhs106_] = rhs197_
                lhs107_[lhs108_] = rhs198_
                d_1635_cf15_ = rhs199_
                d_1613_v98_ = d_1613_v98_
                index227_ = default__.safeIndex(572, (d_1614_v99_).length(0))
                (d_1614_v99_)[index227_] = d_1583_v69_
            elif source25_.is_DC8:
                d_1641___mcc_h5_ = source25_.cf11
                d_1642_cf11_ = d_1641___mcc_h5_
                d_1643_v113_: _dafny.Array
                def lambda85_(d_1644_v69_):
                    def lambda86_(d_1645_i5_):
                        return D0_DC2(d_1644_v69_, True)

                    return lambda86_

                init42_ = lambda85_(d_1583_v69_)
                nw257_ = _dafny.Array(None, 13)
                for i0_42_ in range(nw257_.length(0)):
                    nw257_[i0_42_] = init42_(i0_42_)
                d_1643_v113_ = nw257_
                d_1646_v114_: D0
                d_1646_v114_ = D0_DC2(d_1583_v69_, d_1512_v0_)
                index228_ = default__.safeIndex(509, (d_1643_v113_).length(0))
                (d_1643_v113_)[index228_] = d_1646_v114_
                pat_let_tv31_ = d_1512_v0_
                pat_let_tv32_ = globalState
                pat_let_tv33_ = d_1512_v0_
                index229_ = default__.safeIndex(509, (d_1643_v113_).length(0))
                def iife116_(_pat_let28_0):
                    def iife117_(d_1647_dt__update__tmp_h0_):
                        def iife118_(_pat_let29_0):
                            def iife119_(d_1648_dt__update_hcf1_h0_):
                                return D0_DC2(d_1648_dt__update_hcf1_h0_, (d_1647_dt__update__tmp_h0_).cf2)
                            return iife119_(_pat_let29_0)
                        return iife118_(len(_dafny.SeqWithoutIsStrInference([default__.fm2(pat_let_tv31_, pat_let_tv32_)])))
                    return iife117_(_pat_let28_0)
                def iife115_(_pat_let27_0):
                    def iife120_(d_1649_dt__update__tmp_h1_):
                        def iife121_(_pat_let30_0):
                            def iife122_(d_1650_dt__update_hcf2_h0_):
                                return D0_DC2((d_1649_dt__update__tmp_h1_).cf1, d_1650_dt__update_hcf2_h0_)
                            return iife122_(_pat_let30_0)
                        return iife121_((False) == (pat_let_tv33_))
                    return iife120_(_pat_let27_0)
                (d_1643_v113_)[index229_] = iife115_(iife116_(d_1646_v114_))
                d_1651_v115_: bool
                d_1652_v116_: _dafny.Seq
                d_1653_v117_: _dafny.Seq
                out100_: bool
                out101_: _dafny.Seq
                out102_: _dafny.Seq
                out100_, out101_, out102_ = (self).m8(globalState)
                d_1651_v115_ = out100_
                d_1652_v116_ = out101_
                d_1653_v117_ = out102_
                d_1613_v98_ = ((d_1613_v98_).intersection(d_1613_v98_)) | (d_1613_v98_)
                d_1654_v118_: _dafny.Map
                d_1654_v118_ = _dafny.Map({d_1651_v115_: d_1583_v69_})
                d_1655_v119_: _dafny.Seq
                d_1655_v119_ = _dafny.SeqWithoutIsStrInference([d_1654_v118_])
                d_1656_v120_: _dafny.Map
                d_1656_v120_ = _dafny.Map({d_1583_v69_: d_1512_v0_})
                d_1657_v121_: D8
                d_1657_v121_ = D8_DC23(d_1654_v118_)
                d_1658_v122_: _dafny.Array
                nw258_ = _dafny.Array(None, 24)
                nw258_[int(0)] = d_1654_v118_
                nw258_[int(1)] = d_1654_v118_
                nw258_[int(2)] = (d_1655_v119_)[default__.safeIndex(len(d_1656_v120_), len(d_1655_v119_))]
                nw258_[int(3)] = (d_1654_v118_) | (d_1654_v118_)
                nw258_[int(4)] = (d_1654_v118_).set(d_1651_v115_, len(d_1609_v94_))
                nw258_[int(5)] = d_1654_v118_
                nw258_[int(6)] = (d_1654_v118_).set(not(d_1512_v0_), 191)
                nw258_[int(7)] = d_1654_v118_
                nw258_[int(8)] = (d_1655_v119_)[default__.safeIndex(d_1583_v69_, len(d_1655_v119_))]
                nw258_[int(9)] = d_1654_v118_
                nw258_[int(10)] = d_1654_v118_
                nw258_[int(11)] = (d_1654_v118_) | (d_1654_v118_)
                nw258_[int(12)] = d_1654_v118_
                nw258_[int(13)] = d_1654_v118_
                nw258_[int(14)] = (d_1657_v121_).cf35
                nw258_[int(15)] = _dafny.Map({d_1651_v115_: -276})
                nw258_[int(16)] = d_1654_v118_
                nw258_[int(17)] = (_dafny.Map({d_1651_v115_: d_1583_v69_})).set(d_1512_v0_, 858)
                nw258_[int(18)] = d_1654_v118_
                nw258_[int(19)] = d_1654_v118_
                nw258_[int(20)] = d_1654_v118_
                nw258_[int(21)] = d_1654_v118_
                nw258_[int(22)] = (d_1654_v118_).set(d_1651_v115_, len(_dafny.SeqWithoutIsStrInference([len(d_1608_v93_) for d_1659_i6_ in range(default__.abs(-849))])))
                nw258_[int(23)] = (d_1654_v118_) | (d_1654_v118_)
                d_1658_v122_ = nw258_
                d_1658_v122_ = d_1658_v122_
            elif True:
                d_1660___mcc_h6_ = source25_.cf17
                d_1661_cf17_ = d_1660___mcc_h6_
                d_1583_v69_ = d_1583_v69_
                (globalState).f4 = d_1512_v0_
                d_1512_v0_ = (len(d_1606_v91_)) <= (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydrbvsno"))) + (d_1606_v91_)))
                d_1662_v123_: C0
                nw259_ = C0()
                nw259_.ctor__()
                d_1662_v123_ = nw259_
        elif True:
            d_1582_v68_ = d_1582_v68_
            d_1663_v124_: _dafny.Map
            d_1663_v124_ = _dafny.Map({d_1512_v0_: d_1512_v0_})
            d_1664_v125_: D3
            d_1664_v125_ = D3_DC11(D3_DC8(d_1663_v124_))
            source26_ = d_1664_v125_
            if source26_.is_DC9:
                d_1665___mcc_h7_ = source26_.cf12
                d_1666___mcc_h8_ = source26_.cf13
                d_1667_cf13_ = d_1666___mcc_h8_
                d_1668_cf12_ = d_1665___mcc_h7_
                d_1669_v126_: _dafny.Array
                nw260_ = _dafny.Array(None, 22)
                nw260_[int(0)] = d_1606_v91_
                nw260_[int(1)] = d_1606_v91_
                nw260_[int(2)] = d_1606_v91_
                nw260_[int(3)] = ((d_1607_v92_)[d_1512_v0_] if (d_1512_v0_) in (d_1607_v92_) else d_1606_v91_)
                nw260_[int(4)] = d_1606_v91_
                nw260_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xunrvxhf"))
                nw260_[int(6)] = d_1606_v91_
                nw260_[int(7)] = d_1606_v91_
                nw260_[int(8)] = d_1606_v91_
                nw260_[int(9)] = d_1606_v91_
                nw260_[int(10)] = d_1606_v91_
                nw260_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1582_v68_ for d_1670_i7_ in range(default__.abs(623))])
                nw260_[int(12)] = d_1606_v91_
                nw260_[int(13)] = (d_1606_v91_).set(default__.safeIndex(d_1668_cf12_, len(d_1606_v91_)), d_1582_v68_)
                nw260_[int(14)] = d_1606_v91_
                nw260_[int(15)] = d_1606_v91_
                nw260_[int(16)] = d_1606_v91_
                nw260_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                nw260_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_1582_v68_ for d_1671_i8_ in range(default__.abs(580))])).set(default__.safeIndex(len(_dafny.Map({d_1667_cf13_: d_1512_v0_})), len(_dafny.SeqWithoutIsStrInference([d_1582_v68_ for d_1672_i8_ in range(default__.abs(580))]))), d_1582_v68_)
                nw260_[int(19)] = d_1606_v91_
                nw260_[int(20)] = d_1606_v91_
                nw260_[int(21)] = d_1606_v91_
                d_1669_v126_ = nw260_
                d_1673_v127_: _dafny.Map
                d_1673_v127_ = _dafny.Map({(0) - (d_1667_cf13_): d_1669_v126_})
                d_1673_v127_ = (d_1673_v127_).set(d_1667_cf13_, d_1669_v126_)
                d_1674_v128_: _dafny.Map
                d_1674_v128_ = _dafny.Map({331: (d_1610_v95_).issubset(d_1610_v95_)})
                d_1674_v128_ = (d_1674_v128_).set(d_1668_cf12_, (399) > (d_1583_v69_))
                d_1583_v69_ = d_1583_v69_
                rhs200_ = d_1583_v69_
                rhs201_ = (d_1514_v2_)[default__.safeIndex(d_1583_v69_, len(d_1514_v2_))]
                d_1583_v69_ = rhs200_
                d_1513_v1_ = rhs201_
            elif source26_.is_DC10:
                d_1675___mcc_h9_ = source26_.cf14
                d_1676___mcc_h10_ = source26_.cf15
                d_1677___mcc_h11_ = source26_.cf16
                d_1678_cf16_ = d_1677___mcc_h11_
                d_1679_cf15_ = d_1676___mcc_h10_
                d_1680_cf14_ = d_1675___mcc_h9_
                d_1679_cf15_ = -590
                d_1681_v129_: _dafny.Array
                nw261_ = _dafny.Array(int(0), 17)
                d_1681_v129_ = nw261_
                d_1682_v130_: D8
                d_1682_v130_ = D8_DC24(d_1680_cf14_)
                rhs202_ = (d_1682_v130_).cf36
                rhs203_ = d_1681_v129_
                lhs109_ = globalState
                lhs109_.f4 = rhs202_
                d_1681_v129_ = rhs203_
                d_1583_v69_ = 0
                d_1606_v91_ = d_1606_v91_
            elif source26_.is_DC8:
                d_1683___mcc_h12_ = source26_.cf11
                d_1684_cf11_ = d_1683___mcc_h12_
                d_1583_v69_ = (len(d_1606_v91_) if (d_1583_v69_) < (d_1583_v69_) else 962)
                d_1685_v131_: _dafny.Array
                nw262_ = _dafny.Array(False, 25)
                d_1685_v131_ = nw262_
                d_1685_v131_ = d_1685_v131_
                d_1686_v132_: _dafny.Map
                d_1686_v132_ = _dafny.Map({d_1512_v0_: d_1684_cf11_})
                d_1687_v133_: C1
                nw263_ = C1()
                nw263_.ctor__(d_1512_v0_, d_1512_v0_, d_1686_v132_, d_1513_v1_, d_1605_v90_)
                d_1687_v133_ = nw263_
                d_1688_v134_: _dafny.Array
                def lambda87_(d_1689_i9_):
                    return (d_1689_i9_) + (len((D4_DC12(_dafny.SeqWithoutIsStrInference([False]))).cf18))

                init43_ = lambda87_
                nw264_ = _dafny.Array(None, 24)
                for i0_43_ in range(nw264_.length(0)):
                    nw264_[i0_43_] = init43_(i0_43_)
                d_1688_v134_ = nw264_
                index230_ = default__.safeIndex(526, (d_1688_v134_).length(0))
                (d_1688_v134_)[index230_] = (d_1609_v94_)[default__.safeIndex(d_1583_v69_, len(d_1609_v94_))]
                index231_ = default__.safeIndex(526, (d_1688_v134_).length(0))
                (d_1688_v134_)[index231_] = (d_1583_v69_) * (d_1583_v69_)
            elif True:
                d_1690___mcc_h13_ = source26_.cf17
                d_1691_cf17_ = d_1690___mcc_h13_
                d_1692_v135_: _dafny.Set
                d_1692_v135_ = _dafny.Set({d_1583_v69_, d_1583_v69_})
                d_1693_v136_: D6
                d_1693_v136_ = D6_DC17(d_1692_v135_)
                d_1694_v137_: D6
                d_1694_v137_ = D6_DC17((d_1693_v136_).cf27)
                d_1695_v138_: _dafny.Seq
                d_1695_v138_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1583_v69_ for d_1696_i10_ in range(default__.abs(640))])])
                d_1697_v139_: _dafny.Map
                d_1697_v139_ = _dafny.Map({d_1694_v137_: (d_1695_v138_) + (d_1695_v138_)})
                d_1697_v139_ = (d_1697_v139_).set(D6_DC17(d_1692_v135_), (_dafny.SeqWithoutIsStrInference([d_1608_v93_, _dafny.SeqWithoutIsStrInference([d_1583_v69_]), d_1608_v93_])) + (_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_1583_v69_, d_1583_v69_])).set(default__.safeIndex((D3_DC9(d_1583_v69_, d_1583_v69_)).cf12, len(_dafny.SeqWithoutIsStrInference([d_1583_v69_, d_1583_v69_]))), (0) - (d_1583_v69_)) for d_1698_i11_ in range(default__.abs(473))])))
                d_1608_v93_ = ((d_1695_v138_)[default__.safeIndex(d_1583_v69_, len(d_1695_v138_))]).set(default__.safeIndex(d_1583_v69_, len((d_1695_v138_)[default__.safeIndex(d_1583_v69_, len(d_1695_v138_))])), -147)
                d_1699_v140_: C0
                nw265_ = C0()
                nw265_.ctor__()
                d_1699_v140_ = nw265_
                d_1583_v69_ = 149
            d_1700_v141_: _dafny.Array
            nw266_ = _dafny.Array(False, 23)
            d_1700_v141_ = nw266_
            index232_ = default__.safeIndex(757, (d_1700_v141_).length(0))
            (d_1700_v141_)[index232_] = False
            index233_ = default__.safeIndex(757, (d_1700_v141_).length(0))
            (d_1700_v141_)[index233_] = d_1512_v0_
            if ((d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))]) and ((d_1583_v69_) != ((0) - (d_1583_v69_))):
                d_1606_v91_ = default__.fm29(d_1512_v0_, d_1512_v0_, not(d_1512_v0_), -320, globalState)
                d_1701_v142_: _dafny.Set
                d_1701_v142_ = _dafny.Set({d_1583_v69_})
                d_1702_v143_: _dafny.MultiSet
                d_1702_v143_ = _dafny.MultiSet([d_1512_v0_, (d_1701_v142_).ispropersubset(d_1701_v142_), d_1512_v0_, (d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))]])
                d_1702_v143_ = d_1702_v143_
                d_1703_v144_: _dafny.Array
                nw267_ = _dafny.Array(int(0), 24)
                d_1703_v144_ = nw267_
                d_1703_v144_ = d_1703_v144_
                d_1704_v145_: _dafny.Map
                d_1704_v145_ = _dafny.Map({d_1583_v69_: d_1583_v69_})
                d_1704_v145_ = d_1704_v145_
                d_1705_v146_: _dafny.MultiSet
                d_1705_v146_ = _dafny.MultiSet([d_1703_v144_, d_1703_v144_, d_1703_v144_, d_1703_v144_])
                d_1705_v146_ = d_1705_v146_
            elif True:
                d_1583_v69_ = d_1583_v69_
                d_1706_v147_: _dafny.MultiSet
                d_1706_v147_ = _dafny.MultiSet([(d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))]])
                d_1707_v148_: _dafny.Map
                d_1707_v148_ = _dafny.Map({d_1582_v68_: d_1512_v0_})
                index234_ = default__.safeIndex(546, (d_1605_v90_).length(0))
                (d_1605_v90_)[index234_] = _dafny.MultiSet([((d_1706_v147_)[(d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))]] if ((d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))]) in (d_1706_v147_) else len(_dafny.SeqWithoutIsStrInference([d_1583_v69_, len(d_1707_v148_)])))])
                index235_ = default__.safeIndex(757, (d_1700_v141_).length(0))
                (d_1700_v141_)[index235_] = d_1512_v0_
                d_1583_v69_ = d_1583_v69_
                (globalState).f4 = (d_1512_v0_) or ((d_1700_v141_)[default__.safeIndex(757, (d_1700_v141_).length(0))])
            d_1512_v0_ = d_1512_v0_
        r0 = d_1512_v0_
        return r0

    def m8(self, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Seq = _dafny.Seq({})
        d_1708_v0_: bool
        d_1708_v0_ = True
        d_1709_v1_: _dafny.Array
        nw268_ = _dafny.Array(None, 6)
        nw268_[int(0)] = not(not(d_1708_v0_))
        nw268_[int(1)] = False
        nw268_[int(2)] = d_1708_v0_
        nw268_[int(3)] = d_1708_v0_
        nw268_[int(4)] = d_1708_v0_
        nw268_[int(5)] = d_1708_v0_
        d_1709_v1_ = nw268_
        d_1710_v2_: _dafny.Array
        nw269_ = _dafny.Array(None, 15)
        nw269_[int(0)] = d_1709_v1_
        nw269_[int(1)] = d_1709_v1_
        nw269_[int(2)] = d_1709_v1_
        nw269_[int(3)] = d_1709_v1_
        nw269_[int(4)] = d_1709_v1_
        nw269_[int(5)] = (d_1709_v1_ if d_1708_v0_ else d_1709_v1_)
        nw269_[int(6)] = d_1709_v1_
        nw269_[int(7)] = d_1709_v1_
        nw269_[int(8)] = d_1709_v1_
        nw269_[int(9)] = d_1709_v1_
        nw269_[int(10)] = d_1709_v1_
        nw269_[int(11)] = d_1709_v1_
        nw269_[int(12)] = d_1709_v1_
        nw269_[int(13)] = d_1709_v1_
        nw269_[int(14)] = d_1709_v1_
        d_1710_v2_ = nw269_
        d_1710_v2_ = d_1710_v2_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1709_v1_).length(0)):
            d_1711_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1711_i0_)) and ((d_1711_i0_) < ((d_1709_v1_).length(0)))):
                (d_1709_v1_)[(d_1711_i0_)] = not(not(d_1708_v0_))
        d_1712_v3_: _dafny.Array
        def lambda88_(d_1713_i2_):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1714_i3_ in range(default__.abs(73))])

        init44_ = lambda88_
        nw270_ = _dafny.Array(None, 10)
        for i0_44_ in range(nw270_.length(0)):
            nw270_[i0_44_] = init44_(i0_44_)
        d_1712_v3_ = nw270_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1712_v3_).length(0)):
            d_1715_i1_: int = guard_loop_3_
            if (True) and (((0) <= (d_1715_i1_)) and ((d_1715_i1_) < ((d_1712_v3_).length(0)))):
                (d_1712_v3_)[(d_1715_i1_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vaderto"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1716_i4_ in range(default__.abs(411))]))
        d_1717_v4_: _dafny.Seq
        d_1717_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfbewqfba"))
        d_1718_v5_: int
        d_1718_v5_ = 571
        d_1719_v6_: _dafny.Array
        nw271_ = _dafny.Array(None, 8)
        nw271_[int(0)] = 867
        nw271_[int(1)] = d_1718_v5_
        nw271_[int(2)] = d_1718_v5_
        nw271_[int(3)] = d_1718_v5_
        nw271_[int(4)] = d_1718_v5_
        nw271_[int(5)] = d_1718_v5_
        nw271_[int(6)] = d_1718_v5_
        nw271_[int(7)] = d_1718_v5_
        d_1719_v6_ = nw271_
        d_1720_v7_: _dafny.Map
        d_1720_v7_ = _dafny.Map({len(d_1717_v4_): d_1719_v6_})
        d_1721_v8_: _dafny.Set
        d_1721_v8_ = _dafny.Set({len(d_1720_v7_)})
        d_1722_v10_: _dafny.Seq
        d_1722_v10_ = _dafny.SeqWithoutIsStrInference([d_1708_v0_, d_1708_v0_])
        d_1723_v11_: _dafny.Map
        d_1723_v11_ = _dafny.Map({d_1708_v0_: d_1718_v5_})
        d_1724_v12_: _dafny.Map
        d_1724_v12_ = _dafny.Map({len(d_1722_v10_): d_1723_v11_})
        d_1725_v13_: D4
        def iife123_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in _dafny.IntegerRange(453, 964):
                d_1726_v9_: int = compr_61_
                if ((453) <= (d_1726_v9_)) and ((d_1726_v9_) < (964)):
                    coll61_ = coll61_.union(_dafny.Set([default__.safeModuloInt(d_1726_v9_, d_1718_v5_)]))
            return _dafny.Set(coll61_)
        d_1725_v13_ = D4_DC13(default__.fm2(d_1708_v0_, globalState), (d_1721_v8_).issubset(iife123_()
), default__.fm0(d_1708_v0_, d_1724_v12_, d_1708_v0_, globalState), default__.safeModuloInt(d_1718_v5_, d_1718_v5_), True)
        source27_ = d_1725_v13_
        if source27_.is_DC13:
            d_1727___mcc_h0_ = source27_.cf19
            d_1728___mcc_h1_ = source27_.cf20
            d_1729___mcc_h2_ = source27_.cf21
            d_1730___mcc_h3_ = source27_.cf22
            d_1731___mcc_h4_ = source27_.cf23
            d_1732_cf23_ = d_1731___mcc_h4_
            d_1733_cf22_ = d_1730___mcc_h3_
            d_1734_cf21_ = d_1729___mcc_h2_
            d_1735_cf20_ = d_1728___mcc_h1_
            d_1736_cf19_ = d_1727___mcc_h0_
            index236_ = default__.safeIndex(635, (d_1719_v6_).length(0))
            (d_1719_v6_)[index236_] = 608
            d_1737_v14_: str
            d_1737_v14_ = _dafny.CodePoint('e')
            d_1738_v15_: _dafny.Map
            d_1738_v15_ = _dafny.Map({d_1734_cf21_: d_1737_v14_})
            index237_ = default__.safeIndex(635, (d_1719_v6_).length(0))
            (d_1719_v6_)[index237_] = (d_1718_v5_) + ((0) - (len(default__.fm32(d_1736_cf19_, d_1737_v14_, len(d_1738_v15_), 322, globalState))))
            d_1739_v16_: _dafny.Map
            d_1739_v16_ = _dafny.Map({d_1708_v0_: d_1736_cf19_})
            d_1740_v17_: _dafny.Map
            d_1740_v17_ = _dafny.Map({d_1708_v0_: d_1739_v16_})
            d_1741_v18_: _dafny.Array
            def lambda89_(d_1742_cf22_):
                def lambda90_(d_1743_i5_):
                    return _dafny.MultiSet([d_1742_cf22_])

                return lambda90_

            init45_ = lambda89_(d_1733_cf22_)
            nw272_ = _dafny.Array(None, 14)
            for i0_45_ in range(nw272_.length(0)):
                nw272_[i0_45_] = init45_(i0_45_)
            d_1741_v18_ = nw272_
            d_1744_v19_: T1
            nw273_ = C1()
            nw273_.ctor__(d_1732_cf23_, default__.fm2(d_1736_cf19_, globalState), d_1740_v17_, d_1722_v10_, d_1741_v18_)
            d_1744_v19_ = nw273_
            d_1745_v20_: _dafny.Array
            def lambda91_(d_1746_v14_):
                def lambda92_(d_1747_i6_):
                    return _dafny.Map({_dafny.CodePoint('s'): d_1746_v14_})

                return lambda92_

            init46_ = lambda91_(d_1737_v14_)
            nw274_ = _dafny.Array(None, 28)
            for i0_46_ in range(nw274_.length(0)):
                nw274_[i0_46_] = init46_(i0_46_)
            d_1745_v20_ = nw274_
            d_1748_v21_: _dafny.Seq
            d_1748_v21_ = _dafny.SeqWithoutIsStrInference([d_1744_v19_, d_1744_v19_, d_1744_v19_])
            index238_ = default__.safeIndex(635, (d_1719_v6_).length(0))
            rhs204_ = d_1734_cf21_
            rhs205_ = ((d_1744_v19_ if d_1735_cf20_ else d_1744_v19_) if (d_1736_cf19_ if d_1732_cf23_ else d_1708_v0_) else (d_1744_v19_ if d_1732_cf23_ else (d_1748_v21_)[default__.safeIndex(d_1734_cf21_, len(d_1748_v21_))]))
            rhs206_ = not(default__.fm2(d_1708_v0_, globalState))
            rhs207_ = d_1745_v20_
            lhs110_ = d_1719_v6_
            lhs111_ = default__.safeIndex(635, (d_1719_v6_).length(0))
            lhs110_[lhs111_] = rhs204_
            d_1744_v19_ = rhs205_
            d_1736_cf19_ = rhs206_
            d_1745_v20_ = rhs207_
            d_1749_v22_: _dafny.Map
            d_1749_v22_ = _dafny.Map({d_1709_v1_: d_1736_cf19_})
            d_1750_v23_: _dafny.Array
            nw275_ = _dafny.Array(int(0), 23)
            d_1750_v23_ = nw275_
            d_1751_v24_: _dafny.Array
            nw276_ = _dafny.Array(None, 18)
            nw276_[int(0)] = d_1750_v23_
            nw276_[int(1)] = d_1750_v23_
            nw276_[int(2)] = d_1750_v23_
            nw276_[int(3)] = d_1750_v23_
            nw276_[int(4)] = d_1750_v23_
            nw276_[int(5)] = d_1750_v23_
            nw276_[int(6)] = d_1750_v23_
            nw276_[int(7)] = d_1750_v23_
            nw276_[int(8)] = d_1750_v23_
            nw276_[int(9)] = d_1750_v23_
            nw276_[int(10)] = d_1750_v23_
            nw276_[int(11)] = d_1750_v23_
            nw276_[int(12)] = d_1750_v23_
            nw276_[int(13)] = d_1750_v23_
            nw276_[int(14)] = d_1750_v23_
            nw276_[int(15)] = d_1750_v23_
            nw276_[int(16)] = d_1750_v23_
            nw276_[int(17)] = d_1750_v23_
            d_1751_v24_ = nw276_
            d_1752_v25_: _dafny.Array
            d_1752_v25_ = d_1751_v24_
            rhs208_ = d_1749_v22_
            rhs209_ = (d_1717_v4_) + (d_1717_v4_)
            rhs210_ = (d_1752_v25_)
            d_1749_v22_ = rhs208_
            d_1717_v4_ = rhs209_
            d_1751_v24_ = rhs210_
            d_1753_v26_: _dafny.Array
            d_1753_v26_ = d_1744_v19_.f12
            d_1754_v27_: T2
            nw277_ = C3()
            nw277_.ctor__(True, (d_1753_v26_), (len((d_1744_v19_).f14)) in (d_1721_v8_), d_1740_v17_, (d_1744_v19_).f14)
            d_1754_v27_ = nw277_
            d_1755_v28_: D14
            d_1755_v28_ = D14_DC40(d_1754_v27_)
            d_1754_v27_ = (d_1755_v28_).cf63
        elif source27_.is_DC14:
            d_1756___mcc_h5_ = source27_.cf24
            d_1757_cf24_ = d_1756___mcc_h5_
            d_1758_v29_: str
            d_1758_v29_ = _dafny.CodePoint('c')
            d_1759_v30_: _dafny.Seq
            d_1759_v30_ = _dafny.SeqWithoutIsStrInference([d_1717_v4_, d_1717_v4_, d_1717_v4_])
            rhs211_ = False
            rhs212_ = (d_1717_v4_ if (d_1758_v29_) not in (d_1717_v4_) else ((d_1759_v30_)[default__.safeIndex(536, len(d_1759_v30_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmbesbw"))))
            d_1708_v0_ = rhs211_
            d_1717_v4_ = rhs212_
            d_1760_v31_: _dafny.Map
            d_1760_v31_ = _dafny.Map({d_1708_v0_: d_1757_cf24_})
            d_1761_v32_: _dafny.Map
            d_1761_v32_ = _dafny.Map({d_1757_cf24_: d_1760_v31_})
            d_1762_v33_: _dafny.Array
            nw278_ = _dafny.Array(_dafny.MultiSet({}), 11)
            d_1762_v33_ = nw278_
            d_1763_v34_: C1
            nw279_ = C1()
            nw279_.ctor__(d_1757_cf24_, d_1757_cf24_, (d_1761_v32_) | (d_1761_v32_), default__.fm32(d_1708_v0_, _dafny.CodePoint('o'), d_1718_v5_, d_1718_v5_, globalState), d_1762_v33_)
            d_1763_v34_ = nw279_
            r2 = _dafny.SeqWithoutIsStrInference([(d_1763_v34_).f17, d_1708_v0_])
            d_1764_v35_: C0
            nw280_ = C0()
            nw280_.ctor__()
            d_1764_v35_ = nw280_
        elif True:
            d_1765___mcc_h6_ = source27_.cf18
            d_1766_cf18_ = d_1765___mcc_h6_
            d_1708_v0_ = d_1708_v0_
            d_1767_v36_: _dafny.Map
            d_1767_v36_ = _dafny.Map({d_1718_v5_: d_1718_v5_})
            d_1718_v5_ = ((d_1767_v36_)[723] if (723) in (d_1767_v36_) else (0) - (d_1718_v5_))
            d_1718_v5_ = ((0) - (default__.fm0(d_1708_v0_, d_1724_v12_, d_1708_v0_, globalState))) - (d_1718_v5_)
            d_1768_v37_: _dafny.MultiSet
            d_1768_v37_ = _dafny.MultiSet([d_1708_v0_])
            d_1723_v11_ = (d_1723_v11_) | (default__.fm45(False, (d_1768_v37_).cardinality, (0) - ((0) - (d_1718_v5_)), globalState))
        d_1708_v0_ = d_1708_v0_
        d_1769_v38_: _dafny.Map
        d_1769_v38_ = _dafny.Map({d_1718_v5_: d_1708_v0_})
        d_1770_v39_: _dafny.Seq
        d_1770_v39_ = _dafny.SeqWithoutIsStrInference([403])
        d_1771_v40_: D7
        d_1771_v40_ = D7_DC21((d_1769_v38_) | (_dafny.Map({d_1718_v5_: d_1708_v0_})), default__.fm2(d_1708_v0_, globalState), not((d_1718_v5_) <= ((d_1770_v39_)[default__.safeIndex(d_1718_v5_, len(d_1770_v39_))])))
        source28_ = d_1771_v40_
        if source28_.is_DC21:
            d_1772___mcc_h7_ = source28_.cf31
            d_1773___mcc_h8_ = source28_.cf32
            d_1774___mcc_h9_ = source28_.cf33
            d_1775_cf33_ = d_1774___mcc_h9_
            d_1776_cf32_ = d_1773___mcc_h8_
            d_1777_cf31_ = d_1772___mcc_h7_
            if d_1708_v0_:
                d_1708_v0_ = default__.fm2(d_1775_cf33_, globalState)
                d_1778_v41_: _dafny.Map
                d_1778_v41_ = _dafny.Map({d_1718_v5_: d_1717_v4_})
                d_1779_v42_: _dafny.MultiSet
                d_1779_v42_ = _dafny.MultiSet([((d_1778_v41_)[21] if (21) in (d_1778_v41_) else d_1717_v4_)])
                d_1780_v43_: _dafny.MultiSet
                d_1780_v43_ = _dafny.MultiSet([d_1718_v5_, d_1718_v5_])
                d_1781_v44_: _dafny.Map
                d_1781_v44_ = _dafny.Map({True: d_1775_cf33_})
                d_1782_v45_: str
                d_1782_v45_ = _dafny.CodePoint('r')
                d_1783_v46_: _dafny.Array
                nw281_ = _dafny.Array(None, 29)
                nw281_[int(0)] = d_1780_v43_
                nw281_[int(1)] = d_1780_v43_
                nw281_[int(2)] = d_1780_v43_
                nw281_[int(3)] = d_1780_v43_
                nw281_[int(4)] = d_1780_v43_
                nw281_[int(5)] = d_1780_v43_
                nw281_[int(6)] = d_1780_v43_
                nw281_[int(7)] = d_1780_v43_
                nw281_[int(8)] = d_1780_v43_
                nw281_[int(9)] = d_1780_v43_
                nw281_[int(10)] = ((d_1780_v43_).set(len(d_1769_v38_), default__.abs(len(d_1781_v44_)))).set(d_1718_v5_, default__.abs(-474))
                nw281_[int(11)] = d_1780_v43_
                nw281_[int(12)] = d_1780_v43_
                nw281_[int(13)] = _dafny.MultiSet([(d_1780_v43_).cardinality])
                nw281_[int(14)] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1718_v5_]))])
                nw281_[int(15)] = _dafny.MultiSet([d_1718_v5_, d_1718_v5_])
                nw281_[int(16)] = d_1780_v43_
                nw281_[int(17)] = d_1780_v43_
                nw281_[int(18)] = d_1780_v43_
                nw281_[int(19)] = d_1780_v43_
                nw281_[int(20)] = d_1780_v43_
                nw281_[int(21)] = d_1780_v43_
                nw281_[int(22)] = d_1780_v43_
                nw281_[int(23)] = default__.fm56(d_1782_v45_, globalState)
                nw281_[int(24)] = d_1780_v43_
                nw281_[int(25)] = d_1780_v43_
                nw281_[int(26)] = _dafny.MultiSet([d_1718_v5_])
                nw281_[int(27)] = d_1780_v43_
                nw281_[int(28)] = d_1780_v43_
                d_1783_v46_ = nw281_
                d_1784_v47_: C1
                nw282_ = C1()
                nw282_.ctor__(d_1775_cf33_, d_1708_v0_, _dafny.Map({d_1708_v0_: _dafny.Map({d_1775_cf33_: d_1776_cf32_})}), _dafny.SeqWithoutIsStrInference([d_1775_cf33_]), d_1783_v46_)
                d_1784_v47_ = nw282_
                d_1785_v48_: _dafny.Map
                d_1785_v48_ = _dafny.Map({d_1784_v47_: d_1775_cf33_})
                (globalState).f4 = (((d_1779_v42_) | (d_1779_v42_)).cardinality) == (len(d_1785_v48_))
                index239_ = default__.safeIndex(552, (d_1709_v1_).length(0))
                (d_1709_v1_)[index239_] = d_1776_cf32_
                index240_ = default__.safeIndex(552, (d_1709_v1_).length(0))
                (d_1709_v1_)[index240_] = (d_1784_v47_).f17
                rhs213_ = ((d_1718_v5_) * (d_1718_v5_)) * (d_1718_v5_)
                rhs214_ = (d_1718_v5_) - (len((d_1722_v10_) + (d_1722_v10_)))
                d_1718_v5_ = rhs213_
                d_1718_v5_ = rhs214_
                index241_ = default__.safeIndex(443, (d_1712_v3_).length(0))
                (d_1712_v3_)[index241_] = d_1717_v4_
                index242_ = default__.safeIndex(443, (d_1712_v3_).length(0))
                (d_1712_v3_)[index242_] = d_1717_v4_
            elif True:
                d_1786_v49_: _dafny.Map
                d_1786_v49_ = _dafny.Map({len(default__.fm58(len(d_1722_v10_), globalState)): len(d_1770_v39_)})
                d_1786_v49_ = (d_1786_v49_).set(d_1718_v5_, d_1718_v5_)
                d_1787_v50_: _dafny.Map
                d_1787_v50_ = _dafny.Map({d_1717_v4_: default__.fm0(d_1776_cf32_, _dafny.Map({((d_1786_v49_)[d_1718_v5_] if (d_1718_v5_) in (d_1786_v49_) else d_1718_v5_): d_1723_v11_}), False, globalState)})
                d_1787_v50_ = (d_1787_v50_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1788_i7_ in range(default__.abs(773))]), (0) - (-880))
                d_1789_v51_: D2
                d_1789_v51_ = D2_DC5(len(d_1717_v4_), (d_1775_cf33_ if d_1775_cf33_ else d_1708_v0_), d_1775_cf33_)
                d_1789_v51_ = d_1789_v51_
                def lambda93_(d_1790_v10_, d_1791_v39_):
                    def lambda94_(d_1792_i8_):
                        return (d_1790_v10_)[default__.safeIndex((_dafny.MultiSet(d_1791_v39_)).cardinality, len(d_1790_v10_))]

                    return lambda94_

                init47_ = lambda93_(d_1722_v10_, d_1770_v39_)
                nw283_ = _dafny.Array(None, 3)
                for i0_47_ in range(nw283_.length(0)):
                    nw283_[i0_47_] = init47_(i0_47_)
                d_1709_v1_ = nw283_
                d_1718_v5_ = d_1718_v5_
            d_1793_v52_: _dafny.MultiSet
            d_1793_v52_ = _dafny.MultiSet([d_1718_v5_, d_1718_v5_])
            d_1794_v53_: _dafny.Map
            d_1794_v53_ = _dafny.Map({d_1718_v5_: d_1793_v52_})
            d_1723_v11_ = (d_1723_v11_).set((d_1718_v5_) <= (d_1718_v5_), default__.safeDivisionInt(d_1718_v5_, len(d_1794_v53_)))
            r0 = d_1775_cf33_
            if d_1775_cf33_:
                d_1718_v5_ = d_1718_v5_
                d_1795_v54_: _dafny.Map
                d_1795_v54_ = _dafny.Map({not(d_1708_v0_): d_1775_cf33_})
                d_1796_v55_: _dafny.Map
                d_1796_v55_ = _dafny.Map({d_1775_cf33_: d_1795_v54_})
                d_1797_v56_: _dafny.Array
                nw284_ = _dafny.Array(None, 4)
                nw284_[int(0)] = d_1793_v52_
                nw284_[int(1)] = d_1793_v52_
                nw284_[int(2)] = _dafny.MultiSet([d_1718_v5_, d_1718_v5_, (0) - (d_1718_v5_)])
                nw284_[int(3)] = d_1793_v52_
                d_1797_v56_ = nw284_
                d_1798_v57_: T1
                nw285_ = C1()
                nw285_.ctor__(d_1708_v0_, d_1708_v0_, d_1796_v55_, d_1722_v10_, d_1797_v56_)
                d_1798_v57_ = nw285_
                d_1799_v58_: _dafny.Map
                d_1799_v58_ = _dafny.Map({len(d_1723_v11_): d_1798_v57_})
                d_1800_v59_: _dafny.MultiSet
                d_1800_v59_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfqfoua"))])
                d_1801_v60_: str
                d_1801_v60_ = _dafny.CodePoint('w')
                (globalState).f4 = ((d_1770_v39_) != (_dafny.SeqWithoutIsStrInference([d_1718_v5_, len(d_1799_v58_)]))) or ((d_1718_v5_) <= (((d_1800_v59_)[(d_1717_v4_).set(default__.safeIndex(d_1718_v5_, len(d_1717_v4_)), d_1801_v60_)] if ((d_1717_v4_).set(default__.safeIndex(d_1718_v5_, len(d_1717_v4_)), d_1801_v60_)) in (d_1800_v59_) else d_1718_v5_)))
                d_1802_v61_: D4
                d_1802_v61_ = D4_DC14(d_1775_cf33_)
                d_1708_v0_ = (d_1802_v61_).cf24
                index243_ = default__.safeIndex(600, (d_1709_v1_).length(0))
                (d_1709_v1_)[index243_] = (d_1708_v0_) or (d_1776_cf32_)
                d_1803_v62_: _dafny.Set
                d_1803_v62_ = _dafny.Set({d_1801_v60_})
                index244_ = default__.safeIndex(600, (d_1709_v1_).length(0))
                (d_1709_v1_)[index244_] = default__.fm2((d_1803_v62_).issubset(d_1803_v62_), globalState)
                d_1804_v63_: D4
                d_1804_v63_ = D4_DC12(d_1722_v10_)
                r0 = (default__.fm44(d_1804_v63_, True, globalState)).cf6
            elif True:
                d_1805_v64_: _dafny.Map
                d_1805_v64_ = _dafny.Map({d_1776_cf32_: d_1776_cf32_})
                def iife124_():
                    coll62_ = _dafny.Set()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(835, 884):
                        d_1806_v65_: int = compr_62_
                        if ((835) <= (d_1806_v65_)) and ((d_1806_v65_) < (884)):
                            coll62_ = coll62_.union(_dafny.Set([default__.safeDivisionInt(d_1806_v65_, d_1718_v5_)]))
                    return _dafny.Set(coll62_)
                d_1805_v64_ = default__.fm59(d_1793_v52_, d_1776_cf32_, (d_1775_cf33_ if d_1776_cf32_ else d_1776_cf32_), (d_1718_v5_ if d_1776_cf32_ else len(iife124_()
                )), globalState)
                (globalState).f4 = (d_1805_v64_) == (_dafny.Map({d_1708_v0_: d_1708_v0_}))
                index245_ = default__.safeIndex(89, (d_1709_v1_).length(0))
                (d_1709_v1_)[index245_] = d_1775_cf33_
                d_1807_v66_: _dafny.Seq
                d_1807_v66_ = _dafny.SeqWithoutIsStrInference([d_1721_v8_, d_1721_v8_])
                d_1808_v67_: _dafny.Seq
                d_1808_v67_ = _dafny.SeqWithoutIsStrInference([d_1807_v66_, d_1807_v66_, d_1807_v66_, d_1807_v66_, d_1807_v66_])
                index246_ = default__.safeIndex(89, (d_1709_v1_).length(0))
                (d_1709_v1_)[index246_] = ((d_1808_v67_)[default__.safeIndex(916, len(d_1808_v67_))]) < (d_1807_v66_)
                d_1722_v10_ = _dafny.SeqWithoutIsStrInference([(d_1722_v10_)[default__.safeIndex(d_1718_v5_, len(d_1722_v10_))]])
                d_1809_v68_: _dafny.Set
                d_1809_v68_ = _dafny.Set({d_1775_cf33_, True, d_1776_cf32_})
                d_1810_v69_: D7
                d_1810_v69_ = D7_DC20(d_1793_v52_)
                d_1809_v68_ = _dafny.Set({(d_1709_v1_)[default__.safeIndex(89, (d_1709_v1_).length(0))], (d_1793_v52_).issubset((d_1810_v69_).cf30), (d_1722_v10_)[default__.safeIndex(len(d_1722_v10_), len(d_1722_v10_))]})
        elif source28_.is_DC22:
            d_1811___mcc_h10_ = source28_.cf34
            d_1812_cf34_ = d_1811___mcc_h10_
            d_1812_cf34_ = d_1812_cf34_
            d_1813_v70_: str
            d_1813_v70_ = _dafny.CodePoint('m')
            d_1814_v71_: _dafny.Map
            d_1814_v71_ = _dafny.Map({(d_1813_v70_ if not(d_1708_v0_) else d_1813_v70_): d_1709_v1_})
            d_1814_v71_ = (d_1814_v71_).set((d_1717_v4_)[default__.safeIndex(d_1718_v5_, len(d_1717_v4_))], d_1709_v1_)
            d_1815_v72_: _dafny.Map
            d_1815_v72_ = _dafny.Map({d_1812_cf34_: _dafny.Map({d_1812_cf34_: d_1708_v0_})})
            d_1816_v73_: _dafny.MultiSet
            d_1816_v73_ = _dafny.MultiSet([d_1812_cf34_, d_1812_cf34_])
            d_1817_v74_: _dafny.MultiSet
            d_1817_v74_ = _dafny.MultiSet([(d_1816_v73_).cardinality])
            d_1818_v75_: _dafny.Set
            d_1818_v75_ = _dafny.Set({d_1812_cf34_})
            d_1819_v76_: _dafny.Array
            nw286_ = _dafny.Array(None, 21)
            nw286_[int(0)] = d_1817_v74_
            nw286_[int(1)] = (d_1817_v74_).set(d_1718_v5_, default__.abs(269))
            nw286_[int(2)] = d_1817_v74_
            nw286_[int(3)] = (d_1817_v74_).set(len(d_1818_v75_), default__.abs(d_1718_v5_))
            nw286_[int(4)] = d_1817_v74_
            nw286_[int(5)] = d_1817_v74_
            nw286_[int(6)] = d_1817_v74_
            nw286_[int(7)] = d_1817_v74_
            nw286_[int(8)] = d_1817_v74_
            nw286_[int(9)] = d_1817_v74_
            nw286_[int(10)] = d_1817_v74_
            nw286_[int(11)] = d_1817_v74_
            nw286_[int(12)] = _dafny.MultiSet([d_1718_v5_, d_1718_v5_])
            nw286_[int(13)] = d_1817_v74_
            nw286_[int(14)] = d_1817_v74_
            nw286_[int(15)] = d_1817_v74_
            nw286_[int(16)] = d_1817_v74_
            nw286_[int(17)] = d_1817_v74_
            nw286_[int(18)] = d_1817_v74_
            nw286_[int(19)] = d_1817_v74_
            nw286_[int(20)] = d_1817_v74_
            d_1819_v76_ = nw286_
            d_1820_v77_: C1
            nw287_ = C1()
            nw287_.ctor__(d_1708_v0_, d_1708_v0_, d_1815_v72_, _dafny.SeqWithoutIsStrInference([d_1812_cf34_]), d_1819_v76_)
            d_1820_v77_ = nw287_
            pat_let_tv34_ = d_1718_v5_
            def iife125_(_pat_let31_0):
                def iife126_(d_1821_dt__update__tmp_h0_):
                    def iife127_(_pat_let32_0):
                        def iife128_(d_1822_dt__update_hcf49_h0_):
                            return D11_DC33((d_1821_dt__update__tmp_h0_).cf48, d_1822_dt__update_hcf49_h0_, (d_1821_dt__update__tmp_h0_).cf50, (d_1821_dt__update__tmp_h0_).cf51, (d_1821_dt__update__tmp_h0_).cf52)
                        return iife128_(_pat_let32_0)
                    return iife127_(pat_let_tv34_)
                return iife126_(_pat_let31_0)
            source29_ = iife125_(D11_DC33(d_1812_cf34_, d_1718_v5_, not(True), d_1820_v77_, (d_1820_v77_).f17))
            if source29_.is_DC31:
                (d_1820_v77_).f18 = d_1708_v0_
                index247_ = default__.safeIndex(586, (d_1719_v6_).length(0))
                (d_1719_v6_)[index247_] = d_1718_v5_
                index248_ = default__.safeIndex(586, (d_1719_v6_).length(0))
                (d_1719_v6_)[index248_] = (0) - (d_1718_v5_)
                d_1718_v5_ = (d_1719_v6_)[default__.safeIndex(586, (d_1719_v6_).length(0))]
                index249_ = default__.safeIndex(17, (d_1709_v1_).length(0))
                (d_1709_v1_)[index249_] = (True if d_1708_v0_ else (d_1820_v77_).f17)
                index250_ = default__.safeIndex(17, (d_1709_v1_).length(0))
                (d_1709_v1_)[index250_] = not(False)
            elif source29_.is_DC32:
                d_1823___mcc_h12_ = source29_.cf46
                d_1824___mcc_h13_ = source29_.cf47
                d_1825_cf47_ = d_1824___mcc_h13_
                d_1826_cf46_ = d_1823___mcc_h12_
                d_1718_v5_ = (0) - (d_1718_v5_)
                d_1719_v6_ = d_1719_v6_
                index251_ = default__.safeIndex(392, (d_1710_v2_).length(0))
                (d_1710_v2_)[index251_] = d_1709_v1_
                index252_ = default__.safeIndex(392, (d_1710_v2_).length(0))
                (d_1710_v2_)[index252_] = d_1709_v1_
                (globalState).f4 = (d_1820_v77_).f17
            elif source29_.is_DC33:
                d_1827___mcc_h14_ = source29_.cf48
                d_1828___mcc_h15_ = source29_.cf49
                d_1829___mcc_h16_ = source29_.cf50
                d_1830___mcc_h17_ = source29_.cf51
                d_1831___mcc_h18_ = source29_.cf52
                d_1832_cf52_ = d_1831___mcc_h18_
                d_1833_cf51_ = d_1830___mcc_h17_
                d_1834_cf50_ = d_1829___mcc_h16_
                d_1835_cf49_ = d_1828___mcc_h15_
                d_1836_cf48_ = d_1827___mcc_h14_
                d_1837_v78_: T2
                nw288_ = C3()
                nw288_.ctor__(d_1836_cf48_, d_1819_v76_, (d_1820_v77_).f17, default__.fm60(d_1835_cf49_, globalState), d_1722_v10_)
                d_1837_v78_ = nw288_
                d_1838_v79_: D14
                d_1838_v79_ = D14_DC40(d_1837_v78_)
                d_1839_v80_: D14
                d_1839_v80_ = D14_DC43((D14_DC42() if (d_1833_cf51_).f17 else d_1838_v79_))
                rhs215_ = d_1820_v77_.f18
                rhs216_ = d_1718_v5_
                rhs217_ = d_1839_v80_
                rhs218_ = d_1718_v5_
                lhs112_ = d_1820_v77_
                lhs112_.f18 = rhs215_
                d_1718_v5_ = rhs216_
                d_1839_v80_ = rhs217_
                d_1718_v5_ = rhs218_
                def iife129_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(-568, -612):
                        d_1840_v81_: int = compr_63_
                        if ((-568) <= (d_1840_v81_)) and ((d_1840_v81_) < (-612)):
                            coll63_[default__.safeModuloInt(d_1840_v81_, d_1835_cf49_)] = len(_dafny.SeqWithoutIsStrInference([d_1833_cf51_.f18]))
                    return _dafny.Map(coll63_)
                d_1718_v5_ = default__.safeDivisionInt(d_1835_cf49_, default__.safeModuloInt(428, len(iife129_()
                )))
                index253_ = default__.safeIndex(916, (d_1709_v1_).length(0))
                (d_1709_v1_)[index253_] = not(d_1820_v77_.f18)
                d_1841_v82_: _dafny.Map
                d_1841_v82_ = _dafny.Map({d_1836_cf48_: (d_1833_cf51_).f17})
                d_1842_v83_: _dafny.MultiSet
                d_1842_v83_ = d_1816_v73_
                d_1843_v84_: C2
                nw289_ = C2()
                nw289_.ctor__(_dafny.SeqWithoutIsStrInference([d_1835_cf49_]), d_1708_v0_, ((d_1837_v78_).f13).set(d_1833_cf51_.f18, d_1841_v82_), default__.fm43(_dafny.MultiSet((d_1837_v78_).f14), D7_DC22((d_1833_cf51_).f17), globalState), d_1837_v78_.f12)
                d_1843_v84_ = nw289_
                d_1844_v85_: _dafny.Map
                d_1844_v85_ = _dafny.Map({d_1832_cf52_: d_1843_v84_})
                index254_ = default__.safeIndex(916, (d_1709_v1_).length(0))
                rhs219_ = (d_1833_cf51_).f17
                rhs220_ = (d_1835_cf49_) >= (default__.safeDivisionInt(len(d_1844_v85_), d_1835_cf49_))
                rhs221_ = d_1717_v4_
                rhs222_ = d_1709_v1_
                rhs223_ = d_1709_v1_
                lhs113_ = d_1709_v1_
                lhs114_ = default__.safeIndex(916, (d_1709_v1_).length(0))
                lhs115_ = globalState
                lhs113_[lhs114_] = rhs219_
                lhs115_.f4 = rhs220_
                d_1717_v4_ = rhs221_
                d_1709_v1_ = rhs222_
                d_1709_v1_ = rhs223_
                d_1718_v5_ = d_1835_cf49_
            elif True:
                d_1845___mcc_h19_ = source29_.cf45
                d_1846_cf45_ = d_1845___mcc_h19_
                d_1847_v86_: _dafny.Map
                d_1847_v86_ = _dafny.Map({(_dafny.Map({d_1708_v0_: d_1820_v77_.f18}) if not(d_1820_v77_.f18) else _dafny.Map({(d_1820_v77_).f17: not((d_1820_v77_).f17)})): (d_1717_v4_ if d_1820_v77_.f18 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfjvynxmm")))})
                index255_ = default__.safeIndex(288, (d_1719_v6_).length(0))
                (d_1719_v6_)[index255_] = d_1718_v5_
                index256_ = default__.safeIndex(288, (d_1719_v6_).length(0))
                rhs224_ = default__.fm61(globalState)
                rhs225_ = ((613) * (d_1718_v5_)) - (d_1718_v5_)
                lhs116_ = d_1719_v6_
                lhs117_ = default__.safeIndex(288, (d_1719_v6_).length(0))
                d_1847_v86_ = rhs224_
                lhs116_[lhs117_] = rhs225_
                d_1848_v87_: _dafny.Map
                d_1848_v87_ = _dafny.Map({d_1709_v1_: (d_1820_v77_).f17})
                d_1708_v0_ = (d_1848_v87_) != ((d_1848_v87_).set(d_1709_v1_, False))
                d_1849_v88_: bool
                d_1850_v89_: int
                d_1851_v90_: _dafny.Set
                d_1852_v91_: bool
                out103_: bool
                out104_: int
                out105_: _dafny.Set
                out106_: bool
                out103_, out104_, out105_, out106_ = (d_1820_v77_).m3(globalState)
                d_1849_v88_ = out103_
                d_1850_v89_ = out104_
                d_1851_v90_ = out105_
                d_1852_v91_ = out106_
                index257_ = default__.safeIndex(288, (d_1719_v6_).length(0))
                (d_1719_v6_)[index257_] = (d_1719_v6_)[default__.safeIndex(288, (d_1719_v6_).length(0))]
            d_1817_v74_ = d_1817_v74_
        elif True:
            d_1853___mcc_h11_ = source28_.cf30
            d_1854_cf30_ = d_1853___mcc_h11_
            d_1855_v92_: _dafny.Array
            def lambda95_(d_1856_i9_):
                return _dafny.CodePoint('g')

            init48_ = lambda95_
            nw290_ = _dafny.Array(None, 16)
            for i0_48_ in range(nw290_.length(0)):
                nw290_[i0_48_] = init48_(i0_48_)
            d_1855_v92_ = nw290_
            d_1857_v93_: _dafny.Set
            d_1857_v93_ = _dafny.Set({d_1855_v92_})
            rhs226_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ahyfwsh"))
            rhs227_ = default__.fm0((d_1708_v0_ if d_1708_v0_ else False), (d_1724_v12_).set(len(d_1857_v93_), d_1723_v11_), d_1708_v0_, globalState)
            d_1717_v4_ = rhs226_
            d_1718_v5_ = rhs227_
            if not(d_1708_v0_):
                d_1718_v5_ = d_1718_v5_
                index258_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                (d_1719_v6_)[index258_] = 958
                index259_ = default__.safeIndex(169, (d_1719_v6_).length(0))
                (d_1719_v6_)[index259_] = d_1718_v5_
                d_1858_v94_: D8
                d_1858_v94_ = D8_DC26(d_1708_v0_)
                d_1859_v95_: _dafny.Map
                d_1859_v95_ = _dafny.Map({d_1709_v1_: d_1858_v94_})
                d_1860_v96_: D12
                d_1860_v96_ = D12_DC34(d_1859_v95_)
                d_1861_v97_: _dafny.MultiSet
                d_1861_v97_ = _dafny.MultiSet([D12_DC38(d_1860_v96_)])
                index260_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                index261_ = default__.safeIndex(169, (d_1719_v6_).length(0))
                rhs228_ = d_1718_v5_
                rhs229_ = ((d_1861_v97_) - (d_1861_v97_)).issubset(d_1861_v97_)
                rhs230_ = len(default__.fm27(globalState))
                rhs231_ = ((d_1718_v5_) - ((0) - (d_1718_v5_))) * (d_1718_v5_)
                rhs232_ = d_1708_v0_
                lhs118_ = d_1719_v6_
                lhs119_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                lhs120_ = globalState
                lhs121_ = d_1719_v6_
                lhs122_ = default__.safeIndex(169, (d_1719_v6_).length(0))
                lhs123_ = globalState
                lhs118_[lhs119_] = rhs228_
                lhs120_.f4 = rhs229_
                d_1718_v5_ = rhs230_
                lhs121_[lhs122_] = rhs231_
                lhs123_.f4 = rhs232_
                index262_ = default__.safeIndex(554, (d_1709_v1_).length(0))
                (d_1709_v1_)[index262_] = True
                d_1862_v98_: _dafny.Map
                d_1862_v98_ = _dafny.Map({(d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))]: d_1718_v5_})
                d_1863_v99_: str
                d_1863_v99_ = _dafny.CodePoint('j')
                index263_ = default__.safeIndex(554, (d_1709_v1_).length(0))
                index264_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                rhs233_ = d_1708_v0_
                rhs234_ = len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1864_i10_ in range(default__.abs(658))])).set(default__.safeIndex(((d_1862_v98_)[(d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))]] if ((d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))]) in (d_1862_v98_) else (d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))]), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1865_i10_ in range(default__.abs(658))]))), d_1863_v99_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adwid"))) + ((d_1717_v4_).set(default__.safeIndex((d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))], len(d_1717_v4_)), d_1863_v99_))))
                lhs124_ = d_1709_v1_
                lhs125_ = default__.safeIndex(554, (d_1709_v1_).length(0))
                lhs126_ = d_1719_v6_
                lhs127_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                lhs124_[lhs125_] = rhs233_
                lhs126_[lhs127_] = rhs234_
                d_1866_v100_: _dafny.Array
                nw291_ = _dafny.Array(_dafny.Map({}), 5)
                d_1866_v100_ = nw291_
                d_1867_v101_: _dafny.Map
                d_1867_v101_ = _dafny.Map({default__.safeModuloInt(default__.fm0(False, d_1724_v12_, d_1708_v0_, globalState), d_1718_v5_): d_1866_v100_})
                d_1867_v101_ = (d_1867_v101_).set((d_1719_v6_)[default__.safeIndex(360, (d_1719_v6_).length(0))], d_1866_v100_)
                d_1868_v102_: _dafny.Array
                nw292_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1868_v102_ = nw292_
                d_1869_v103_: _dafny.Array
                d_1869_v103_ = d_1868_v102_
                d_1870_v104_: _dafny.Array
                nw293_ = _dafny.Array(int(0), 23)
                d_1870_v104_ = nw293_
                d_1871_v105_: _dafny.Map
                d_1871_v105_ = _dafny.Map({d_1717_v4_: len(d_1717_v4_)})
                index265_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                index266_ = default__.safeIndex(554, (d_1709_v1_).length(0))
                index267_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                rhs235_ = d_1869_v103_
                rhs236_ = (((d_1871_v105_)[(d_1717_v4_).set(default__.safeIndex(d_1718_v5_, len(d_1717_v4_)), d_1863_v99_)] if ((d_1717_v4_).set(default__.safeIndex(d_1718_v5_, len(d_1717_v4_)), d_1863_v99_)) in (d_1871_v105_) else default__.fm0((d_1709_v1_)[default__.safeIndex(554, (d_1709_v1_).length(0))], d_1724_v12_, (d_1722_v10_)[default__.safeIndex(default__.fm0(d_1708_v0_, d_1724_v12_, False, globalState), len(d_1722_v10_))], globalState))) + ((0) - (default__.safeModuloInt(d_1718_v5_, len(d_1723_v11_))))
                rhs237_ = (d_1709_v1_)[default__.safeIndex(554, (d_1709_v1_).length(0))]
                rhs238_ = d_1870_v104_
                rhs239_ = (d_1770_v39_)[default__.safeIndex(d_1718_v5_, len(d_1770_v39_))]
                lhs128_ = d_1719_v6_
                lhs129_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                lhs130_ = d_1709_v1_
                lhs131_ = default__.safeIndex(554, (d_1709_v1_).length(0))
                lhs132_ = d_1719_v6_
                lhs133_ = default__.safeIndex(360, (d_1719_v6_).length(0))
                d_1869_v103_ = rhs235_
                lhs128_[lhs129_] = rhs236_
                lhs130_[lhs131_] = rhs237_
                d_1870_v104_ = rhs238_
                lhs132_[lhs133_] = rhs239_
            elif True:
                d_1718_v5_ = d_1718_v5_
                (globalState).f4 = False
                d_1872_v106_: _dafny.Map
                d_1872_v106_ = _dafny.Map({d_1718_v5_: d_1770_v39_})
                d_1873_v107_: _dafny.Map
                d_1873_v107_ = _dafny.Map({d_1708_v0_: d_1872_v106_})
                index268_ = default__.safeIndex(272, (d_1719_v6_).length(0))
                (d_1719_v6_)[index268_] = len(((d_1873_v107_)[d_1708_v0_] if (d_1708_v0_) in (d_1873_v107_) else d_1872_v106_))
                index269_ = default__.safeIndex(272, (d_1719_v6_).length(0))
                rhs240_ = d_1708_v0_
                rhs241_ = default__.fm0((d_1717_v4_) < (d_1717_v4_), d_1724_v12_, d_1708_v0_, globalState)
                rhs242_ = (d_1718_v5_) < (((0) - (d_1718_v5_) if d_1708_v0_ else d_1718_v5_))
                lhs134_ = d_1719_v6_
                lhs135_ = default__.safeIndex(272, (d_1719_v6_).length(0))
                lhs136_ = globalState
                r0 = rhs240_
                lhs134_[lhs135_] = rhs241_
                lhs136_.f4 = rhs242_
                d_1718_v5_ = (d_1719_v6_)[default__.safeIndex(272, (d_1719_v6_).length(0))]
                d_1718_v5_ = (d_1719_v6_)[default__.safeIndex(272, (d_1719_v6_).length(0))]
            index270_ = default__.safeIndex(943, (d_1719_v6_).length(0))
            (d_1719_v6_)[index270_] = (0) - (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))), d_1718_v5_))
            index271_ = default__.safeIndex(943, (d_1719_v6_).length(0))
            (d_1719_v6_)[index271_] = (len(d_1717_v4_) if d_1708_v0_ else d_1718_v5_)
            if True:
                d_1874_v108_: _dafny.Seq
                d_1874_v108_ = _dafny.SeqWithoutIsStrInference([d_1709_v1_, d_1709_v1_])
                (globalState).f4 = ((d_1874_v108_) + (d_1874_v108_)) != (d_1874_v108_)
                d_1875_v109_: _dafny.Map
                d_1875_v109_ = _dafny.Map({len(d_1722_v10_): d_1770_v39_})
                d_1876_v110_: _dafny.Array
                nw294_ = _dafny.Array(None, 6)
                nw294_[int(0)] = (d_1770_v39_).set(default__.safeIndex(d_1718_v5_, len(d_1770_v39_)), 341)
                nw294_[int(1)] = d_1770_v39_
                nw294_[int(2)] = _dafny.SeqWithoutIsStrInference([68])
                nw294_[int(3)] = (_dafny.SeqWithoutIsStrInference([(d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))], d_1718_v5_])) + (((d_1875_v109_)[(d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))]] if ((d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))]) in (d_1875_v109_) else d_1770_v39_))
                nw294_[int(4)] = d_1770_v39_
                nw294_[int(5)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1877_i12_ in range(default__.abs(-106))])) for d_1878_i11_ in range(default__.abs(126))])
                d_1876_v110_ = nw294_
                index272_ = default__.safeIndex(157, (d_1876_v110_).length(0))
                (d_1876_v110_)[index272_] = default__.fm1(d_1708_v0_, d_1708_v0_, globalState)
                d_1879_v111_: D0
                d_1879_v111_ = D0_DC1()
                d_1880_v113_: _dafny.Map
                d_1880_v113_ = _dafny.Map({(d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))]: d_1709_v1_})
                d_1881_v114_: _dafny.MultiSet
                d_1881_v114_ = _dafny.MultiSet([d_1880_v113_, _dafny.Map({len(d_1717_v4_): d_1709_v1_})])
                index273_ = default__.safeIndex(157, (d_1876_v110_).length(0))
                index274_ = default__.safeIndex(943, (d_1719_v6_).length(0))
                index275_ = default__.safeIndex(943, (d_1719_v6_).length(0))
                def iife130_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(940, 896):
                        d_1883_v112_: int = compr_64_
                        if ((940) <= (d_1883_v112_)) and ((d_1883_v112_) < (896)):
                            coll64_[default__.safeDivisionInt(d_1883_v112_, d_1718_v5_)] = (D8_DC25(d_1708_v0_, (d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))], len(_dafny.Map({D7_DC20(d_1854_cf30_): d_1718_v5_})), _dafny.Map({d_1708_v0_: (d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))]}))).cf40
                    return _dafny.Map(coll64_)
                rhs243_ = (d_1770_v39_) + (_dafny.SeqWithoutIsStrInference([635 for d_1882_i13_ in range(default__.abs(594))]))
                rhs244_ = ((-415) * (d_1718_v5_)) - ((d_1719_v6_)[default__.safeIndex(943, (d_1719_v6_).length(0))])
                rhs245_ = default__.safeModuloInt(default__.fm0(not(d_1708_v0_), iife130_()
                , d_1708_v0_, globalState), ((d_1881_v114_)[d_1880_v113_] if (d_1880_v113_) in (d_1881_v114_) else d_1718_v5_))
                rhs246_ = d_1879_v111_
                lhs137_ = d_1876_v110_
                lhs138_ = default__.safeIndex(157, (d_1876_v110_).length(0))
                lhs139_ = d_1719_v6_
                lhs140_ = default__.safeIndex(943, (d_1719_v6_).length(0))
                lhs141_ = d_1719_v6_
                lhs142_ = default__.safeIndex(943, (d_1719_v6_).length(0))
                lhs137_[lhs138_] = rhs243_
                lhs139_[lhs140_] = rhs244_
                lhs141_[lhs142_] = rhs245_
                d_1879_v111_ = rhs246_
                def lambda96_(d_1884_v0_):
                    def lambda97_(d_1885_i14_):
                        return d_1884_v0_

                    return lambda97_

                init49_ = lambda96_(d_1708_v0_)
                nw295_ = _dafny.Array(None, 5)
                for i0_49_ in range(nw295_.length(0)):
                    nw295_[i0_49_] = init49_(i0_49_)
                d_1709_v1_ = nw295_
                d_1886_v115_: _dafny.Set
                d_1886_v115_ = _dafny.Set({d_1717_v4_})
                (globalState).f4 = (d_1886_v115_).issubset(d_1886_v115_)
                d_1887_v116_: _dafny.MultiSet
                d_1887_v116_ = _dafny.MultiSet([d_1708_v0_, d_1708_v0_, d_1708_v0_])
                d_1888_v117_: _dafny.Map
                d_1888_v117_ = _dafny.Map({default__.fm2(d_1708_v0_, globalState): True})
                d_1889_v118_: _dafny.Map
                d_1889_v118_ = _dafny.Map({default__.fm2(d_1708_v0_, globalState): d_1888_v117_})
                d_1890_v119_: _dafny.Array
                nw296_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_1890_v119_ = nw296_
                d_1891_v120_: C1
                nw297_ = C1()
                nw297_.ctor__(d_1708_v0_, (d_1887_v116_).ispropersubset(d_1887_v116_), d_1889_v118_, d_1722_v10_, d_1890_v119_)
                d_1891_v120_ = nw297_
            elif True:
                d_1770_v39_ = d_1770_v39_
                d_1892_v121_: _dafny.Map
                d_1892_v121_ = _dafny.Map({d_1708_v0_: d_1708_v0_})
                d_1893_v122_: _dafny.Map
                d_1893_v122_ = _dafny.Map({not(d_1708_v0_): d_1892_v121_})
                d_1894_v123_: _dafny.Array
                def lambda98_(d_1895_cf30_):
                    def lambda99_(d_1896_i15_):
                        return d_1895_cf30_

                    return lambda99_

                init50_ = lambda98_(d_1854_cf30_)
                nw298_ = _dafny.Array(None, 29)
                for i0_50_ in range(nw298_.length(0)):
                    nw298_[i0_50_] = init50_(i0_50_)
                d_1894_v123_ = nw298_
                d_1897_v124_: T1
                nw299_ = C1()
                nw299_.ctor__(not(d_1708_v0_), d_1708_v0_, d_1893_v122_, d_1722_v10_, d_1894_v123_)
                d_1897_v124_ = nw299_
                d_1897_v124_ = d_1897_v124_
                d_1708_v0_ = d_1708_v0_
                d_1898_v125_: str
                d_1898_v125_ = _dafny.CodePoint('o')
                index276_ = default__.safeIndex(139, (d_1855_v92_).length(0))
                (d_1855_v92_)[index276_] = d_1898_v125_
                index277_ = default__.safeIndex(139, (d_1855_v92_).length(0))
                (d_1855_v92_)[index277_] = d_1898_v125_
                d_1717_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nugcjll"))
        r0 = d_1708_v0_
        r1 = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1899_i17_ in range(default__.abs(600))]) for d_1900_i16_ in range(default__.abs(254))])
        r2 = (d_1722_v10_) + (d_1722_v10_)
        return r0, r1, r2


class C5(T0, T1, T4):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f19, f12, f13, f14):
        (self)._f19 = f19
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14

    def fm5(self, p0, globalState):
        return default__.safeModuloInt(944, (0) - (len(_dafny.Map({_dafny.MultiSet([D2_DC4(_dafny.SeqWithoutIsStrInference([(self).f14])), D2_DC4(_dafny.SeqWithoutIsStrInference([(self).f14]))]): (self).f14}))))

    def fm6(self, p0, globalState):
        def iife131_():
            coll65_ = _dafny.Map()
            compr_65_: int
            for compr_65_ in (_dafny.SeqWithoutIsStrInference([-887 for d_1901_i0_ in range(default__.abs(-832))])).Elements:
                d_1902_v0_: int = compr_65_
                if (d_1902_v0_) in (_dafny.SeqWithoutIsStrInference([-887 for d_1901_i0_ in range(default__.abs(-832))])):
                    coll65_[(d_1902_v0_) - (435)] = 150
            return _dafny.Map(coll65_)
        return ((_dafny.SeqWithoutIsStrInference([408])) + (_dafny.SeqWithoutIsStrInference([436, len(_dafny.SeqWithoutIsStrInference([len((self).f14), len(_dafny.Map({-25: 366}))])), (0) - (len(iife131_()
        ))]))) + (_dafny.SeqWithoutIsStrInference([(0) - (-155) for d_1903_i1_ in range(default__.abs(699))]))

    def fm10(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Set({293})).intersection(_dafny.Set({-796}))) - ((_dafny.Set({len(_dafny.Set({-378, 984, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnblljb")))}))})) - (_dafny.Set({-459})))

    def m3(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_1904_v0_: int
        d_1904_v0_ = 208
        d_1905_v1_: _dafny.MultiSet
        d_1905_v1_ = _dafny.MultiSet([d_1904_v0_])
        d_1906_v2_: _dafny.MultiSet
        d_1906_v2_ = default__.fm20((self).f19, d_1904_v0_, globalState)
        def lambda100_(source30_):
            d_1907___mcc_h0_ = source30_
            d_1908_cf3_ = d_1907___mcc_h0_
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xh")))])

        d_1905_v1_ = lambda100_(d_1906_v2_)
        d_1909_v3_: D0
        d_1909_v3_ = D0_DC1()
        source31_ = d_1909_v3_
        if source31_.is_DC1:
            d_1910_v4_: _dafny.Array
            def lambda101_(d_1911_i0_):
                return (self).f19

            init51_ = lambda101_
            nw300_ = _dafny.Array(None, 21)
            for i0_51_ in range(nw300_.length(0)):
                nw300_[i0_51_] = init51_(i0_51_)
            d_1910_v4_ = nw300_
            index278_ = default__.safeIndex(297, (d_1910_v4_).length(0))
            (d_1910_v4_)[index278_] = (self).f19
            d_1912_v5_: _dafny.Map
            d_1912_v5_ = _dafny.Map({(self).f19: (self).f19})
            d_1913_v6_: _dafny.MultiSet
            d_1913_v6_ = _dafny.MultiSet([d_1912_v5_, d_1912_v5_])
            index279_ = default__.safeIndex(297, (d_1910_v4_).length(0))
            (d_1910_v4_)[index279_] = (d_1913_v6_).issubset(((d_1913_v6_).set(d_1912_v5_, default__.abs(493)) if default__.fm2((self).f19, globalState) else d_1913_v6_))
            d_1914_v7_: _dafny.Seq
            d_1914_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sacmgtow"))
            d_1914_v7_ = (default__.fm21((self).f19, globalState)).set(default__.safeIndex(d_1904_v0_, len(default__.fm21((self).f19, globalState))), _dafny.CodePoint('m'))
            d_1915_v8_: _dafny.Array
            nw301_ = _dafny.Array(int(0), 22)
            d_1915_v8_ = nw301_
            index280_ = default__.safeIndex(412, (d_1915_v8_).length(0))
            (d_1915_v8_)[index280_] = (d_1904_v0_) - (d_1904_v0_)
            index281_ = default__.safeIndex(412, (d_1915_v8_).length(0))
            (d_1915_v8_)[index281_] = d_1904_v0_
            d_1916_v9_: D0
            d_1916_v9_ = D0_DC2(363, (self).f19)
            d_1916_v9_ = d_1916_v9_
        elif source31_.is_DC2:
            d_1917___mcc_h1_ = source31_.cf1
            d_1918___mcc_h2_ = source31_.cf2
            d_1919_cf2_ = d_1918___mcc_h2_
            d_1920_cf1_ = d_1917___mcc_h1_
            d_1921_v10_: _dafny.Map
            d_1921_v10_ = _dafny.Map({False: d_1919_cf2_})
            d_1922_v11_: _dafny.Map
            d_1922_v11_ = _dafny.Map({len(d_1921_v10_): d_1919_cf2_})
            d_1923_v12_: _dafny.MultiSet
            d_1923_v12_ = _dafny.MultiSet([d_1919_cf2_])
            d_1924_v13_: _dafny.Seq
            d_1924_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsmors"))
            rhs247_ = d_1922_v11_
            rhs248_ = not(((self).f14)[default__.safeIndex(d_1904_v0_, len((self).f14))])
            rhs249_ = (default__.safeModuloInt(d_1904_v0_, (0) - (d_1904_v0_))) <= (((d_1923_v12_)[True] if (True) in (d_1923_v12_) else len(_dafny.SeqWithoutIsStrInference([len(d_1924_v13_)]))))
            d_1922_v11_ = rhs247_
            d_1919_cf2_ = rhs248_
            r3 = rhs249_
            r0 = d_1919_cf2_
            d_1925_v14_: C0
            nw302_ = C0()
            nw302_.ctor__()
            d_1925_v14_ = nw302_
            d_1926_v15_: _dafny.Map
            d_1926_v15_ = _dafny.Map({_dafny.MultiSet([(self).f19, d_1919_cf2_]): d_1925_v14_})
            d_1926_v15_ = (d_1926_v15_).set((d_1923_v12_) | (d_1923_v12_), d_1925_v14_)
            d_1927_v16_: D0
            d_1927_v16_ = D0_DC0(d_1919_cf2_)
            d_1928_v17_: D2
            d_1928_v17_ = D2_DC5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbajbykx"))), (self).f19, (d_1927_v16_).cf0)
            d_1929_v18_: D2
            d_1929_v18_ = D2_DC7(d_1928_v17_)
            d_1929_v18_ = d_1929_v18_
        elif True:
            d_1930___mcc_h3_ = source31_.cf0
            d_1931_cf0_ = d_1930___mcc_h3_
            d_1932_v19_: _dafny.Array
            def lambda102_(d_1933_v0_):
                def lambda103_(d_1934_i1_):
                    return (d_1934_i1_) - (d_1933_v0_)

                return lambda103_

            init52_ = lambda102_(d_1904_v0_)
            nw303_ = _dafny.Array(None, 11)
            for i0_52_ in range(nw303_.length(0)):
                nw303_[i0_52_] = init52_(i0_52_)
            d_1932_v19_ = nw303_
            d_1932_v19_ = d_1932_v19_
            rhs250_ = (d_1904_v0_) * (d_1904_v0_)
            rhs251_ = d_1904_v0_
            d_1904_v0_ = rhs250_
            r1 = rhs251_
            d_1935_v20_: C0
            nw304_ = C0()
            nw304_.ctor__()
            d_1935_v20_ = nw304_
            d_1936_v21_: str
            d_1936_v21_ = _dafny.CodePoint('i')
            d_1937_v22_: _dafny.Seq
            d_1937_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            d_1931_cf0_ = (d_1936_v21_) not in (d_1937_v22_)
        d_1938_v23_: _dafny.Seq
        d_1938_v23_ = _dafny.SeqWithoutIsStrInference([(self).f14, _dafny.SeqWithoutIsStrInference([(self).f19])])
        d_1939_v24_: D2
        d_1939_v24_ = D2_DC4(d_1938_v23_)
        d_1906_v2_ = default__.fm22(d_1939_v24_, globalState)
        if (self).f19:
            d_1904_v0_ = d_1904_v0_
            d_1940_v25_: _dafny.Array
            nw305_ = _dafny.Array(int(0), 16)
            d_1940_v25_ = nw305_
            d_1941_v26_: _dafny.MultiSet
            d_1941_v26_ = _dafny.MultiSet([d_1940_v25_, d_1940_v25_])
            d_1942_v27_: _dafny.Set
            d_1942_v27_ = _dafny.Set({d_1904_v0_})
            d_1943_v28_: _dafny.Seq
            d_1943_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpo"))
            d_1944_v29_: _dafny.Array
            nw306_ = _dafny.Array(None, 22)
            nw306_[int(0)] = ((self).f19 if (self).f19 else (self).f19)
            nw306_[int(1)] = (self).f19
            nw306_[int(2)] = (self).f19
            nw306_[int(3)] = (self).f19
            nw306_[int(4)] = ((self).f19) == (default__.fm2((self).f19, globalState))
            nw306_[int(5)] = (self).f19
            nw306_[int(6)] = (self).f19
            nw306_[int(7)] = (self).f19
            nw306_[int(8)] = ((self).f19) in (_dafny.SeqWithoutIsStrInference([(self).f19, (self).f19, (self).f19, (self).f19, (self).f19]))
            nw306_[int(9)] = not(False)
            nw306_[int(10)] = (d_1941_v26_).issubset((d_1941_v26_).set(d_1940_v25_, default__.abs(d_1904_v0_)))
            nw306_[int(11)] = (-711) != (d_1904_v0_)
            nw306_[int(12)] = (self).f19
            nw306_[int(13)] = (self).f19
            nw306_[int(14)] = (d_1904_v0_) in ((self).fm10(False, d_1942_v27_, _dafny.SeqWithoutIsStrInference([-833 for d_1945_i2_ in range(default__.abs(348))]), d_1904_v0_, globalState))
            nw306_[int(15)] = (self).f19
            nw306_[int(16)] = default__.fm2((self).f19, globalState)
            nw306_[int(17)] = (d_1943_v28_) == (d_1943_v28_)
            nw306_[int(18)] = False
            nw306_[int(19)] = (self).f19
            nw306_[int(20)] = (self).f19
            nw306_[int(21)] = (True) == ((self).f19)
            d_1944_v29_ = nw306_
            index282_ = default__.safeIndex(88, (d_1944_v29_).length(0))
            (d_1944_v29_)[index282_] = (self).f19
            index283_ = default__.safeIndex(88, (d_1944_v29_).length(0))
            (d_1944_v29_)[index283_] = True
            d_1946_v30_: _dafny.Seq
            d_1946_v30_ = _dafny.SeqWithoutIsStrInference([d_1904_v0_])
            d_1904_v0_ = ((0) - (d_1904_v0_)) + ((d_1946_v30_)[default__.safeIndex(-412, len(d_1946_v30_))])
            if (self).f19:
                index284_ = default__.safeIndex(596, (d_1940_v25_).length(0))
                (d_1940_v25_)[index284_] = len((d_1942_v27_).intersection(d_1942_v27_))
                d_1947_v31_: _dafny.Array
                nw307_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1947_v31_ = nw307_
                d_1948_v32_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.CodePoint('D'), 14)
                d_1948_v32_ = nw308_
                index285_ = default__.safeIndex(840, (d_1947_v31_).length(0))
                (d_1947_v31_)[index285_] = d_1948_v32_
                d_1949_v33_: _dafny.Map
                d_1949_v33_ = _dafny.Map({615: (self).f19})
                index286_ = default__.safeIndex(596, (d_1940_v25_).length(0))
                index287_ = default__.safeIndex(840, (d_1947_v31_).length(0))
                rhs252_ = d_1904_v0_
                rhs253_ = default__.safeDivisionInt(d_1904_v0_, d_1904_v0_)
                rhs254_ = default__.safeDivisionInt(501, 576)
                rhs255_ = d_1948_v32_
                rhs256_ = (d_1946_v30_)[default__.safeIndex(len(d_1949_v33_), len(d_1946_v30_))]
                lhs143_ = d_1940_v25_
                lhs144_ = default__.safeIndex(596, (d_1940_v25_).length(0))
                lhs145_ = d_1947_v31_
                lhs146_ = default__.safeIndex(840, (d_1947_v31_).length(0))
                lhs143_[lhs144_] = rhs252_
                d_1904_v0_ = rhs253_
                r1 = rhs254_
                lhs145_[lhs146_] = rhs255_
                r1 = rhs256_
                d_1950_v34_: _dafny.Array
                nw309_ = _dafny.Array(int(0), 23)
                d_1950_v34_ = nw309_
                d_1950_v34_ = d_1950_v34_
                d_1946_v30_ = (d_1946_v30_) + (d_1946_v30_)
                d_1951_v35_: D0
                d_1951_v35_ = D0_DC0((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))])
                index288_ = default__.safeIndex(88, (d_1944_v29_).length(0))
                index289_ = default__.safeIndex(88, (d_1944_v29_).length(0))
                index290_ = default__.safeIndex(596, (d_1940_v25_).length(0))
                rhs257_ = False
                rhs258_ = (d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))]
                rhs259_ = (len(d_1943_v28_) if default__.fm2((d_1951_v35_).cf0, globalState) else (d_1940_v25_)[default__.safeIndex(596, (d_1940_v25_).length(0))])
                lhs147_ = d_1944_v29_
                lhs148_ = default__.safeIndex(88, (d_1944_v29_).length(0))
                lhs149_ = d_1944_v29_
                lhs150_ = default__.safeIndex(88, (d_1944_v29_).length(0))
                lhs151_ = d_1940_v25_
                lhs152_ = default__.safeIndex(596, (d_1940_v25_).length(0))
                lhs147_[lhs148_] = rhs257_
                lhs149_[lhs150_] = rhs258_
                lhs151_[lhs152_] = rhs259_
                d_1904_v0_ = (self).fm5((self).f19, globalState)
            elif True:
                (globalState).f4 = (self).f19
                (globalState).f4 = (self).f19
                rhs260_ = (0) - (d_1904_v0_)
                rhs261_ = d_1944_v29_
                r1 = rhs260_
                d_1944_v29_ = rhs261_
                r3 = not((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))])
                d_1952_v36_: C0
                nw310_ = C0()
                nw310_.ctor__()
                d_1952_v36_ = nw310_
            d_1953_v37_: str
            d_1953_v37_ = _dafny.CodePoint('r')
            source32_ = (d_1939_v24_ if ((self).f14)[default__.safeIndex(d_1904_v0_, len((self).f14))] else default__.fm23(d_1904_v0_, d_1953_v37_, globalState))
            if source32_.is_DC5:
                d_1954___mcc_h4_ = source32_.cf5
                d_1955___mcc_h5_ = source32_.cf6
                d_1956___mcc_h6_ = source32_.cf7
                d_1957_cf7_ = d_1956___mcc_h6_
                d_1958_cf6_ = d_1955___mcc_h5_
                d_1959_cf5_ = d_1954___mcc_h4_
                d_1959_cf5_ = 174
                (globalState).f4 = d_1957_cf7_
                d_1960_v38_: D0
                d_1960_v38_ = D0_DC2(d_1959_cf5_, True)
                d_1961_v39_: _dafny.Seq
                d_1962_v40_: bool
                d_1963_v41_: int
                d_1964_v42_: bool
                out107_: _dafny.Seq
                out108_: bool
                out109_: int
                out110_: bool
                out107_, out108_, out109_, out110_ = default__.m0(False, d_1960_v38_, (self).f19, globalState)
                d_1961_v39_ = out107_
                d_1962_v40_ = out108_
                d_1963_v41_ = out109_
                d_1964_v42_ = out110_
                d_1965_v43_: _dafny.Set
                d_1965_v43_ = _dafny.Set({d_1944_v29_, d_1944_v29_, d_1944_v29_, d_1944_v29_})
                d_1965_v43_ = ((d_1965_v43_) | (d_1965_v43_)) - (d_1965_v43_)
            elif source32_.is_DC6:
                d_1966___mcc_h7_ = source32_.cf8
                d_1967___mcc_h8_ = source32_.cf9
                d_1968_cf9_ = d_1967___mcc_h8_
                d_1969_cf8_ = d_1966___mcc_h7_
                r1 = (d_1904_v0_) * (default__.safeModuloInt(d_1904_v0_, len((self).f14)))
                d_1970_v44_: D2
                d_1970_v44_ = D2_DC5(d_1904_v0_, (self).f19, (self).f19)
                d_1971_v45_: D2
                d_1971_v45_ = D2_DC7(d_1970_v44_)
                d_1971_v45_ = d_1971_v45_
                (globalState).f4 = (d_1904_v0_) == ((d_1904_v0_) * ((d_1946_v30_)[default__.safeIndex(d_1904_v0_, len(d_1946_v30_))]))
                r3 = (self).f19
            elif source32_.is_DC4:
                d_1972___mcc_h9_ = source32_.cf4
                d_1973_cf4_ = d_1972___mcc_h9_
                d_1974_v46_: _dafny.Map
                d_1974_v46_ = _dafny.Map({d_1953_v37_: (self).f19})
                d_1974_v46_ = (d_1974_v46_).set(d_1953_v37_, (d_1946_v30_) <= (d_1946_v30_))
                index291_ = default__.safeIndex(673, (d_1940_v25_).length(0))
                (d_1940_v25_)[index291_] = d_1904_v0_
                index292_ = default__.safeIndex(673, (d_1940_v25_).length(0))
                (d_1940_v25_)[index292_] = (0) - ((default__.safeDivisionInt(d_1904_v0_, d_1904_v0_)) + (d_1904_v0_))
                d_1975_v47_: _dafny.Map
                d_1975_v47_ = _dafny.Map({(self).f19: d_1946_v30_})
                d_1976_v48_: _dafny.Map
                d_1976_v48_ = _dafny.Map({(self).f19: not((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))])})
                d_1977_v49_: _dafny.Map
                d_1977_v49_ = _dafny.Map({d_1975_v47_: ((d_1976_v48_)[(d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))]] if ((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))]) in (d_1976_v48_) else (d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))])})
                d_1977_v49_ = (d_1977_v49_).set(d_1975_v47_, (d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))])
                d_1978_v50_: _dafny.Array
                def lambda104_(d_1979_v37_):
                    def lambda105_(d_1980_i3_):
                        return d_1979_v37_

                    return lambda105_

                init53_ = lambda104_(d_1953_v37_)
                nw311_ = _dafny.Array(None, 9)
                for i0_53_ in range(nw311_.length(0)):
                    nw311_[i0_53_] = init53_(i0_53_)
                d_1978_v50_ = nw311_
                index293_ = default__.safeIndex(346, (d_1978_v50_).length(0))
                (d_1978_v50_)[index293_] = d_1953_v37_
                index294_ = default__.safeIndex(346, (d_1978_v50_).length(0))
                (d_1978_v50_)[index294_] = (d_1953_v37_ if True else d_1953_v37_)
            elif True:
                d_1981___mcc_h10_ = source32_.cf10
                d_1982_cf10_ = d_1981___mcc_h10_
                d_1983_v51_: _dafny.Map
                d_1983_v51_ = _dafny.Map({d_1904_v0_: d_1904_v0_})
                r1 = len((_dafny.Map({len((self).f14): d_1904_v0_})) | (d_1983_v51_))
                r0 = default__.fm2(not ((self).f19) or ((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))]), globalState)
                d_1904_v0_ = (d_1904_v0_) * (d_1904_v0_)
                d_1984_v52_: _dafny.Set
                d_1984_v52_ = _dafny.Set({default__.fm24((d_1944_v29_)[default__.safeIndex(88, (d_1944_v29_).length(0))], _dafny.MultiSet(d_1946_v30_), 749, globalState)})
                (globalState).f4 = (d_1984_v52_).ispropersubset(d_1984_v52_)
        elif True:
            d_1985_v53_: _dafny.Array
            nw312_ = _dafny.Array(False, 27)
            d_1985_v53_ = nw312_
            index295_ = default__.safeIndex(34, (d_1985_v53_).length(0))
            (d_1985_v53_)[index295_] = (self).f19
            index296_ = default__.safeIndex(34, (d_1985_v53_).length(0))
            (d_1985_v53_)[index296_] = ((self).f14)[default__.safeIndex(d_1904_v0_, len((self).f14))]
            d_1986_v54_: _dafny.Map
            d_1986_v54_ = _dafny.Map({d_1904_v0_: d_1985_v53_})
            d_1987_v55_: _dafny.Array
            nw313_ = _dafny.Array(None, 29)
            nw313_[int(0)] = d_1985_v53_
            nw313_[int(1)] = d_1985_v53_
            nw313_[int(2)] = d_1985_v53_
            nw313_[int(3)] = d_1985_v53_
            nw313_[int(4)] = (d_1985_v53_ if (d_1985_v53_)[default__.safeIndex(34, (d_1985_v53_).length(0))] else d_1985_v53_)
            nw313_[int(5)] = ((d_1986_v54_)[d_1904_v0_] if (d_1904_v0_) in (d_1986_v54_) else d_1985_v53_)
            nw313_[int(6)] = d_1985_v53_
            nw313_[int(7)] = d_1985_v53_
            nw313_[int(8)] = d_1985_v53_
            nw313_[int(9)] = d_1985_v53_
            nw313_[int(10)] = d_1985_v53_
            nw313_[int(11)] = d_1985_v53_
            nw313_[int(12)] = d_1985_v53_
            nw313_[int(13)] = d_1985_v53_
            nw313_[int(14)] = d_1985_v53_
            nw313_[int(15)] = d_1985_v53_
            nw313_[int(16)] = d_1985_v53_
            nw313_[int(17)] = d_1985_v53_
            nw313_[int(18)] = d_1985_v53_
            nw313_[int(19)] = d_1985_v53_
            nw313_[int(20)] = d_1985_v53_
            nw313_[int(21)] = d_1985_v53_
            nw313_[int(22)] = d_1985_v53_
            nw313_[int(23)] = d_1985_v53_
            nw313_[int(24)] = d_1985_v53_
            nw313_[int(25)] = d_1985_v53_
            nw313_[int(26)] = d_1985_v53_
            nw313_[int(27)] = d_1985_v53_
            nw313_[int(28)] = d_1985_v53_
            d_1987_v55_ = nw313_
            d_1987_v55_ = d_1987_v55_
            d_1988_v56_: C1
            nw314_ = C1()
            nw314_.ctor__((d_1985_v53_)[default__.safeIndex(34, (d_1985_v53_).length(0))], ((d_1985_v53_)[default__.safeIndex(34, (d_1985_v53_).length(0))]) or ((d_1985_v53_)[default__.safeIndex(34, (d_1985_v53_).length(0))]), (self).f13, (self).f14, self.f12)
            d_1988_v56_ = nw314_
            d_1989_v57_: _dafny.Array
            nw315_ = _dafny.Array(int(0), 9)
            d_1989_v57_ = nw315_
            index297_ = default__.safeIndex(993, (d_1989_v57_).length(0))
            (d_1989_v57_)[index297_] = (d_1904_v0_) * (d_1904_v0_)
            index298_ = default__.safeIndex(993, (d_1989_v57_).length(0))
            (d_1989_v57_)[index298_] = ((d_1988_v56_).fm5(not(False), globalState)) * (d_1904_v0_)
            d_1990_v58_: _dafny.Map
            d_1990_v58_ = _dafny.Map({(self).f14: (d_1988_v56_).f17})
            r3 = (d_1990_v58_) == ((d_1990_v58_).set((self).f14, (self).f19))
        d_1991_i4_: int
        d_1991_i4_ = 0
        with _dafny.label("11"):
            while (self).f19:
                with _dafny.c_label("11"):
                    if (d_1991_i4_) >= (100):
                        raise _dafny.Break("11")
                    d_1991_i4_ = (d_1991_i4_) + (1)
                    d_1992_v59_: _dafny.Array
                    nw316_ = _dafny.Array(None, 1)
                    nw316_[int(0)] = not(not (True) or (True))
                    d_1992_v59_ = nw316_
                    index299_ = default__.safeIndex(941, (d_1992_v59_).length(0))
                    (d_1992_v59_)[index299_] = (self).f19
                    index300_ = default__.safeIndex(941, (d_1992_v59_).length(0))
                    (d_1992_v59_)[index300_] = (self).f19
                    d_1904_v0_ = (43) - (d_1904_v0_)
                    d_1993_v60_: str
                    d_1993_v60_ = _dafny.CodePoint('s')
                    d_1994_v61_: _dafny.Set
                    d_1994_v61_ = _dafny.Set({d_1993_v60_})
                    index301_ = default__.safeIndex(941, (d_1992_v59_).length(0))
                    (d_1992_v59_)[index301_] = (d_1993_v60_) in ((d_1994_v61_ if False else d_1994_v61_))
                    d_1995_v62_: _dafny.Map
                    d_1995_v62_ = _dafny.Map({d_1904_v0_: (self).f13})
                    d_1996_v63_: _dafny.Map
                    d_1996_v63_ = _dafny.Map({(d_1992_v59_)[default__.safeIndex(941, (d_1992_v59_).length(0))]: (self).f19})
                    d_1997_v64_: D3
                    d_1997_v64_ = D3_DC8(d_1996_v63_)
                    d_1998_v65_: C1
                    nw317_ = C1()
                    nw317_.ctor__((d_1992_v59_)[default__.safeIndex(941, (d_1992_v59_).length(0))], not(not ((d_1992_v59_)[default__.safeIndex(941, (d_1992_v59_).length(0))]) or ((self).f19)), (((d_1995_v62_)[d_1904_v0_] if (d_1904_v0_) in (d_1995_v62_) else (self).f13)) | (((self).f13).set((self).f19, (d_1997_v64_).cf11)), (self).f14, self.f12)
                    d_1998_v65_ = nw317_
                    pass
            pass
        d_1999_v66_: _dafny.Seq
        d_1999_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyqhtx"))
        d_2000_v67_: _dafny.Map
        d_2000_v67_ = _dafny.Map({d_1999_v66_: not(((self).f14)[default__.safeIndex(d_1904_v0_, len((self).f14))])})
        r3 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqw"))) in (d_2000_v67_)
        r0 = (D0_DC0((self).f19)).cf0
        r1 = d_1904_v0_
        d_2001_v68_: _dafny.Set
        d_2001_v68_ = _dafny.Set({(self).f14})
        r2 = d_2001_v68_
        d_2002_v69_: str
        d_2002_v69_ = _dafny.CodePoint('r')
        d_2003_v70_: _dafny.Map
        d_2003_v70_ = _dafny.Map({(self).f19: d_2002_v69_})
        r3 = (default__.fm2((self).f19, globalState)) or (((self).f19) not in (d_2003_v70_))
        return r0, r1, r2, r3

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        d_2004_v0_: D3
        d_2004_v0_ = D3_DC9(956, p2)
        source33_ = d_2004_v0_
        if source33_.is_DC9:
            d_2005___mcc_h0_ = source33_.cf12
            d_2006___mcc_h1_ = source33_.cf13
            d_2007_cf13_ = d_2006___mcc_h1_
            d_2008_cf12_ = d_2005___mcc_h0_
            d_2009_v1_: D4
            d_2009_v1_ = D4_DC12((self).f14)
            d_2010_v2_: _dafny.MultiSet
            d_2010_v2_ = _dafny.MultiSet([(d_2009_v1_).cf18])
            d_2011_v3_: _dafny.Seq
            d_2011_v3_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2012_v4_: _dafny.Map
            d_2012_v4_ = _dafny.Map({d_2010_v2_: ((d_2011_v3_) + (d_2011_v3_)).set(default__.safeIndex(d_2008_cf12_, len((d_2011_v3_) + (d_2011_v3_))), p1)})
            d_2012_v4_ = (d_2012_v4_).set(d_2010_v2_, d_2011_v3_)
            d_2013_v5_: _dafny.Array
            nw318_ = _dafny.Array(False, 11)
            d_2013_v5_ = nw318_
            d_2014_v6_: T1
            nw319_ = C1()
            nw319_.ctor__((D2_DC5(d_2008_cf12_, (self).f19, (self).f19)).cf6, (self).f19, (self).f13, (self).f14, self.f12)
            d_2014_v6_ = nw319_
            index302_ = default__.safeIndex(140, (d_2013_v5_).length(0))
            def iife132_(_pat_let33_0):
                def iife133_(d_2015_dt__update__tmp_h0_):
                    def iife134_(_pat_let34_0):
                        def iife135_(d_2016_dt__update_hcf14_h0_):
                            return D3_DC10(d_2016_dt__update_hcf14_h0_, (d_2015_dt__update__tmp_h0_).cf15, (d_2015_dt__update__tmp_h0_).cf16)
                        return iife135_(_pat_let34_0)
                    return iife134_((self).f19)
                return iife133_(_pat_let33_0)
            (d_2013_v5_)[index302_] = (iife132_(D3_DC10((self).f19, d_2008_cf12_, d_2014_v6_))).cf14
            d_2017_v7_: _dafny.Map
            d_2017_v7_ = _dafny.Map({(self).f19: (self).f19})
            index303_ = default__.safeIndex(140, (d_2013_v5_).length(0))
            (d_2013_v5_)[index303_] = ((d_2017_v7_) | (d_2017_v7_)) != (d_2017_v7_)
            d_2018_v8_: _dafny.Seq
            d_2018_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbxdea"))
            d_2018_v8_ = d_2018_v8_
            (globalState).f4 = (p2) <= (default__.safeModuloInt(649, d_2008_cf12_))
        elif source33_.is_DC10:
            d_2019___mcc_h2_ = source33_.cf14
            d_2020___mcc_h3_ = source33_.cf15
            d_2021___mcc_h4_ = source33_.cf16
            d_2022_cf16_ = d_2021___mcc_h4_
            d_2023_cf15_ = d_2020___mcc_h3_
            d_2024_cf14_ = d_2019___mcc_h2_
            d_2025_v9_: _dafny.Seq
            d_2025_v9_ = _dafny.SeqWithoutIsStrInference([p2])
            r0 = default__.fm2((_dafny.SeqWithoutIsStrInference([p2])) == (d_2025_v9_), globalState)
            d_2026_v10_: _dafny.Map
            d_2026_v10_ = _dafny.Map({((d_2022_cf16_).f14)[default__.safeIndex(d_2023_cf15_, len((d_2022_cf16_).f14))]: 351})
            d_2026_v10_ = (d_2026_v10_).set(not((d_2023_cf15_) == (d_2023_cf15_)), d_2023_cf15_)
            d_2027_v11_: _dafny.Map
            d_2027_v11_ = _dafny.Map({(self).f19: (self).f19})
            d_2028_v12_: C1
            nw320_ = C1()
            nw320_.ctor__(d_2024_cf14_, not (d_2024_cf14_) or (d_2024_cf14_), ((d_2022_cf16_).f13) | (((d_2022_cf16_).f13).set(d_2024_cf14_, d_2027_v11_)), (self).f14, (d_2022_cf16_.f12 if d_2024_cf14_ else self.f12))
            d_2028_v12_ = nw320_
            source34_ = default__.fm25(d_2025_v9_, globalState)
            if source34_.is_DC9:
                d_2029___mcc_h7_ = source34_.cf12
                d_2030___mcc_h8_ = source34_.cf13
                d_2031_cf13_ = d_2030___mcc_h8_
                d_2032_cf12_ = d_2029___mcc_h7_
                d_2033_v13_: _dafny.Array
                def lambda106_(d_2034_p2_):
                    def lambda107_(d_2035_i0_):
                        return default__.safeDivisionInt(d_2035_i0_, d_2034_p2_)

                    return lambda107_

                init54_ = lambda106_(p2)
                nw321_ = _dafny.Array(None, 8)
                for i0_54_ in range(nw321_.length(0)):
                    nw321_[i0_54_] = init54_(i0_54_)
                d_2033_v13_ = nw321_
                d_2033_v13_ = (d_2033_v13_ if d_2024_cf14_ else d_2033_v13_)
                (d_2028_v12_).f18 = not (not (d_2028_v12_.f18) or (d_2028_v12_.f18)) or ((self).f19)
                index304_ = default__.safeIndex(3, (d_2033_v13_).length(0))
                (d_2033_v13_)[index304_] = d_2032_cf12_
                d_2036_v14_: D5
                d_2036_v14_ = D5_DC15(d_2025_v9_)
                index305_ = default__.safeIndex(3, (d_2033_v13_).length(0))
                rhs262_ = not((_dafny.SeqWithoutIsStrInference([p2 for d_2037_i1_ in range(default__.abs(789))])) < (((d_2036_v14_).cf25) + (d_2025_v9_)))
                rhs263_ = (d_2025_v9_)[default__.safeIndex(d_2032_cf12_, len(d_2025_v9_))]
                rhs264_ = default__.safeModuloInt(d_2031_cf13_, d_2031_cf13_)
                rhs265_ = d_2024_cf14_
                lhs153_ = globalState
                lhs154_ = d_2033_v13_
                lhs155_ = default__.safeIndex(3, (d_2033_v13_).length(0))
                lhs156_ = d_2028_v12_
                lhs153_.f4 = rhs262_
                d_2023_cf15_ = rhs263_
                lhs154_[lhs155_] = rhs264_
                lhs156_.f18 = rhs265_
                d_2038_v15_: C0
                nw322_ = C0()
                nw322_.ctor__()
                d_2038_v15_ = nw322_
            elif source34_.is_DC10:
                d_2039___mcc_h9_ = source34_.cf14
                d_2040___mcc_h10_ = source34_.cf15
                d_2041___mcc_h11_ = source34_.cf16
                d_2042_cf16_ = d_2041___mcc_h11_
                d_2043_cf15_ = d_2040___mcc_h10_
                d_2044_cf14_ = d_2039___mcc_h9_
                d_2045_v16_: _dafny.Array
                def lambda108_(d_2046_v12_):
                    def lambda109_(d_2047_i2_):
                        return not(d_2046_v12_.f18)

                    return lambda109_

                init55_ = lambda108_(d_2028_v12_)
                nw323_ = _dafny.Array(None, 2)
                for i0_55_ in range(nw323_.length(0)):
                    nw323_[i0_55_] = init55_(i0_55_)
                d_2045_v16_ = nw323_
                d_2048_v17_: _dafny.Set
                d_2048_v17_ = _dafny.Set({d_2045_v16_, d_2045_v16_, d_2045_v16_, d_2045_v16_, d_2045_v16_})
                d_2048_v17_ = d_2048_v17_
                d_2049_v18_: _dafny.Map
                d_2049_v18_ = _dafny.Map({(d_2028_v12_).f17: p1})
                d_2049_v18_ = (d_2049_v18_).set(d_2044_cf14_, (p1) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2050_i3_ in range(default__.abs(951))])))
                d_2051_v19_: _dafny.Array
                nw324_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_2051_v19_ = nw324_
                index306_ = default__.safeIndex(262, (d_2051_v19_).length(0))
                (d_2051_v19_)[index306_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "su"))
                index307_ = default__.safeIndex(262, (d_2051_v19_).length(0))
                (d_2051_v19_)[index307_] = (p1) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yi"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rffepbepm"))))
                (globalState).f4 = default__.fm2((d_2028_v12_).f17, globalState)
            elif source34_.is_DC8:
                d_2052___mcc_h12_ = source34_.cf11
                d_2053_cf11_ = d_2052___mcc_h12_
                d_2054_v20_: _dafny.Array
                nw325_ = _dafny.Array(False, 15)
                d_2054_v20_ = nw325_
                index308_ = default__.safeIndex(286, (d_2054_v20_).length(0))
                (d_2054_v20_)[index308_] = (self).f19
                index309_ = default__.safeIndex(286, (d_2054_v20_).length(0))
                rhs266_ = (d_2028_v12_).f17
                rhs267_ = (d_2028_v12_).fm17((_dafny.SeqWithoutIsStrInference([d_2023_cf15_ for d_2055_i4_ in range(default__.abs(22))])) + (d_2025_v9_), len((p1 if d_2024_cf14_ else p1)), globalState)
                lhs157_ = d_2054_v20_
                lhs158_ = default__.safeIndex(286, (d_2054_v20_).length(0))
                lhs157_[lhs158_] = rhs266_
                d_2024_cf14_ = rhs267_
                d_2056_v21_: C0
                nw326_ = C0()
                nw326_.ctor__()
                d_2056_v21_ = nw326_
                index310_ = default__.safeIndex(286, (d_2054_v20_).length(0))
                (d_2054_v20_)[index310_] = not (((d_2053_cf11_)[not((d_2054_v20_)[default__.safeIndex(286, (d_2054_v20_).length(0))])] if (not((d_2054_v20_)[default__.safeIndex(286, (d_2054_v20_).length(0))])) in (d_2053_cf11_) else d_2028_v12_.f18)) or (((d_2054_v20_)[default__.safeIndex(286, (d_2054_v20_).length(0))] if (d_2054_v20_)[default__.safeIndex(286, (d_2054_v20_).length(0))] else d_2024_cf14_))
                d_2023_cf15_ = p2
            elif True:
                d_2057___mcc_h13_ = source34_.cf17
                d_2058_cf17_ = d_2057___mcc_h13_
                r0 = d_2028_v12_.f18
                d_2059_v22_: _dafny.MultiSet
                d_2059_v22_ = _dafny.MultiSet([p2, p2])
                d_2060_v23_: _dafny.Map
                d_2060_v23_ = _dafny.Map({p1: d_2023_cf15_})
                d_2061_v24_: D6
                d_2061_v24_ = D6_DC17(_dafny.Set({p2, 300, p2}))
                d_2062_v25_: _dafny.Set
                d_2062_v25_ = _dafny.Set({p2})
                pat_let_tv35_ = d_2062_v25_
                d_2063_v26_: _dafny.Array
                nw327_ = _dafny.Array(None, 24)
                nw327_[int(0)] = p2
                nw327_[int(1)] = default__.safeDivisionInt(d_2023_cf15_, d_2023_cf15_)
                nw327_[int(2)] = d_2023_cf15_
                nw327_[int(3)] = d_2023_cf15_
                nw327_[int(4)] = (self).fm5((d_2028_v12_).f17, globalState)
                nw327_[int(5)] = (0) - (p2)
                nw327_[int(6)] = -268
                nw327_[int(7)] = d_2023_cf15_
                nw327_[int(8)] = ((d_2026_v10_)[(d_2028_v12_).f17] if ((d_2028_v12_).f17) in (d_2026_v10_) else ((d_2059_v22_)[d_2023_cf15_] if (d_2023_cf15_) in (d_2059_v22_) else p2))
                nw327_[int(9)] = len(d_2060_v23_)
                nw327_[int(10)] = d_2023_cf15_
                nw327_[int(11)] = d_2023_cf15_
                nw327_[int(12)] = p2
                nw327_[int(13)] = len((d_2022_cf16_).fm6(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2064_i5_ in range(default__.abs(755))]), globalState))
                nw327_[int(14)] = (0) - (d_2023_cf15_)
                nw327_[int(15)] = d_2023_cf15_
                nw327_[int(16)] = (d_2023_cf15_) * (d_2023_cf15_)
                nw327_[int(17)] = (d_2023_cf15_) - (d_2023_cf15_)
                nw327_[int(18)] = d_2023_cf15_
                nw327_[int(19)] = p2
                nw327_[int(20)] = p2
                nw327_[int(21)] = (0) - (-289)
                nw327_[int(22)] = d_2023_cf15_
                def iife136_(_pat_let35_0):
                    def iife137_(d_2065_dt__update__tmp_h1_):
                        def iife138_(_pat_let36_0):
                            def iife139_(d_2066_dt__update_hcf27_h0_):
                                return D6_DC17(d_2066_dt__update_hcf27_h0_)
                            return iife139_(_pat_let36_0)
                        return iife138_(pat_let_tv35_)
                    return iife137_(_pat_let35_0)
                nw327_[int(23)] = (len((iife136_(d_2061_v24_)).cf27)) + (842)
                d_2063_v26_ = nw327_
                d_2067_v27_: D3
                d_2067_v27_ = D3_DC10((d_2028_v12_).f17, d_2023_cf15_, d_2022_cf16_)
                index311_ = default__.safeIndex(210, (d_2063_v26_).length(0))
                (d_2063_v26_)[index311_] = (d_2067_v27_).cf15
                index312_ = default__.safeIndex(210, (d_2063_v26_).length(0))
                (d_2063_v26_)[index312_] = (0) - ((0) - ((d_2023_cf15_) - (d_2023_cf15_)))
                d_2068_v28_: _dafny.MultiSet
                d_2068_v28_ = _dafny.MultiSet([d_2024_cf14_])
                d_2069_v29_: _dafny.Seq
                d_2069_v29_ = _dafny.SeqWithoutIsStrInference([d_2068_v28_, d_2068_v28_, d_2068_v28_])
                d_2070_v30_: _dafny.Map
                d_2070_v30_ = _dafny.Map({not((d_2028_v12_).f17): d_2068_v28_})
                d_2071_v31_: D4
                d_2071_v31_ = D4_DC13(d_2028_v12_.f18, (d_2028_v12_).f17, p2, (d_2028_v12_).fm5((d_2028_v12_).f17, globalState), d_2028_v12_.f18)
                d_2072_v32_: _dafny.Array
                nw328_ = _dafny.Array(None, 27)
                nw328_[int(0)] = d_2068_v28_
                nw328_[int(1)] = _dafny.MultiSet([d_2028_v12_.f18, (d_2028_v12_).f17])
                nw328_[int(2)] = d_2068_v28_
                nw328_[int(3)] = d_2068_v28_
                nw328_[int(4)] = d_2068_v28_
                nw328_[int(5)] = d_2068_v28_
                nw328_[int(6)] = d_2068_v28_
                nw328_[int(7)] = (d_2069_v29_)[default__.safeIndex(d_2023_cf15_, len(d_2069_v29_))]
                nw328_[int(8)] = (_dafny.MultiSet([d_2028_v12_.f18])) - (_dafny.MultiSet((d_2022_cf16_).f14))
                nw328_[int(9)] = d_2068_v28_
                nw328_[int(10)] = (d_2068_v28_).intersection(d_2068_v28_)
                nw328_[int(11)] = d_2068_v28_
                nw328_[int(12)] = _dafny.MultiSet([d_2024_cf14_, (d_2028_v12_).f17])
                nw328_[int(13)] = (d_2068_v28_) - (_dafny.MultiSet((d_2022_cf16_).f14))
                nw328_[int(14)] = d_2068_v28_
                nw328_[int(15)] = (_dafny.MultiSet((d_2022_cf16_).f14)).set(d_2024_cf14_, default__.abs(d_2023_cf15_))
                nw328_[int(16)] = d_2068_v28_
                nw328_[int(17)] = ((d_2070_v30_)[d_2028_v12_.f18] if (d_2028_v12_.f18) in (d_2070_v30_) else _dafny.MultiSet([d_2028_v12_.f18]))
                nw328_[int(18)] = d_2068_v28_
                nw328_[int(19)] = d_2068_v28_
                nw328_[int(20)] = (d_2068_v28_) | (_dafny.MultiSet([True, (d_2028_v12_).fm17(d_2025_v9_, d_2023_cf15_, globalState)]))
                nw328_[int(21)] = (d_2068_v28_ if (self).f19 else d_2068_v28_)
                nw328_[int(22)] = _dafny.MultiSet([d_2028_v12_.f18, False])
                nw328_[int(23)] = d_2068_v28_
                nw328_[int(24)] = d_2068_v28_
                nw328_[int(25)] = _dafny.MultiSet([(d_2071_v31_).cf20])
                nw328_[int(26)] = _dafny.MultiSet(((self).f14) + ((d_2022_cf16_).f14))
                d_2072_v32_ = nw328_
                index313_ = default__.safeIndex(313, (d_2072_v32_).length(0))
                (d_2072_v32_)[index313_] = (d_2068_v28_).set((d_2028_v12_).f17, default__.abs(p2))
                index314_ = default__.safeIndex(210, (d_2063_v26_).length(0))
                index315_ = default__.safeIndex(313, (d_2072_v32_).length(0))
                rhs268_ = (d_2063_v26_)[default__.safeIndex(210, (d_2063_v26_).length(0))]
                rhs269_ = True
                rhs270_ = (0) - ((268) + (p2))
                rhs271_ = (d_2068_v28_) - (d_2068_v28_)
                rhs272_ = (d_2028_v12_).f17
                lhs159_ = d_2063_v26_
                lhs160_ = default__.safeIndex(210, (d_2063_v26_).length(0))
                lhs161_ = globalState
                lhs162_ = d_2072_v32_
                lhs163_ = default__.safeIndex(313, (d_2072_v32_).length(0))
                lhs164_ = globalState
                lhs159_[lhs160_] = rhs268_
                lhs161_.f4 = rhs269_
                d_2023_cf15_ = rhs270_
                lhs162_[lhs163_] = rhs271_
                lhs164_.f4 = rhs272_
                d_2073_v33_: str
                d_2073_v33_ = _dafny.CodePoint('u')
                d_2073_v33_ = d_2073_v33_
        elif source33_.is_DC8:
            d_2074___mcc_h5_ = source33_.cf11
            d_2075_cf11_ = d_2074___mcc_h5_
            if (self).f19:
                d_2076_v34_: _dafny.Array
                def lambda110_(d_2077_i6_):
                    return (self).f19

                init56_ = lambda110_
                nw329_ = _dafny.Array(None, 11)
                for i0_56_ in range(nw329_.length(0)):
                    nw329_[i0_56_] = init56_(i0_56_)
                d_2076_v34_ = nw329_
                index316_ = default__.safeIndex(888, (d_2076_v34_).length(0))
                (d_2076_v34_)[index316_] = (self).f19
                d_2078_v35_: D4
                d_2078_v35_ = D4_DC13(True, (self).f19, p2, p2, (self).f19)
                d_2079_v36_: _dafny.Map
                d_2079_v36_ = _dafny.Map({p2: (self).f19})
                index317_ = default__.safeIndex(888, (d_2076_v34_).length(0))
                rhs273_ = (self).f19
                rhs274_ = ((self).f19 if (self).f19 else (self).f19)
                rhs275_ = ((d_2079_v36_)[p2] if (p2) in (d_2079_v36_) else default__.fm2((self).f19, globalState))
                rhs276_ = D4_DC13((self).f19, (self).f19, (0) - (p2), 894, not((self).f19))
                lhs165_ = d_2076_v34_
                lhs166_ = default__.safeIndex(888, (d_2076_v34_).length(0))
                lhs167_ = globalState
                r0 = rhs273_
                lhs165_[lhs166_] = rhs274_
                lhs167_.f4 = rhs275_
                d_2078_v35_ = rhs276_
                d_2080_v37_: _dafny.Map
                d_2080_v37_ = _dafny.Map({p2: len(_dafny.Set({(d_2076_v34_)[default__.safeIndex(888, (d_2076_v34_).length(0))]}))})
                (globalState).f4 = not((((d_2080_v37_)[p2] if (p2) in (d_2080_v37_) else p2)) <= (p2))
                r0 = (d_2076_v34_)[default__.safeIndex(888, (d_2076_v34_).length(0))]
                d_2081_v38_: _dafny.Array
                nw330_ = _dafny.Array(int(0), 4)
                d_2081_v38_ = nw330_
                d_2082_v39_: _dafny.Map
                d_2082_v39_ = _dafny.Map({d_2081_v38_: p2})
                d_2082_v39_ = (d_2082_v39_).set(d_2081_v38_, (p2) * (p2))
                d_2083_v40_: _dafny.Array
                nw331_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_2083_v40_ = nw331_
                index318_ = default__.safeIndex(67, (d_2083_v40_).length(0))
                (d_2083_v40_)[index318_] = d_2076_v34_
                index319_ = default__.safeIndex(67, (d_2083_v40_).length(0))
                (d_2083_v40_)[index319_] = d_2076_v34_
            elif True:
                (globalState).f4 = (((self).f14)[default__.safeIndex(p2, len((self).f14))]) and ((self).f19)
                d_2084_v41_: _dafny.Map
                d_2084_v41_ = _dafny.Map({(D5_DC16((self).f19)).cf26: len(p1)})
                d_2085_v42_: _dafny.MultiSet
                d_2085_v42_ = _dafny.MultiSet([len((d_2084_v41_).set((self).f19, p2)), p2, p2])
                d_2085_v42_ = d_2085_v42_
                d_2086_v43_: C1
                nw332_ = C1()
                nw332_.ctor__((self).f19, (self).f19, _dafny.Map({(self).f19: _dafny.Map({(self).f19: (self).f19})}), ((self).f14).set(default__.safeIndex(p2, len((self).f14)), not(default__.fm2(True, globalState))), self.f12)
                d_2086_v43_ = nw332_
                d_2087_v44_: int
                d_2087_v44_ = 199
                d_2087_v44_ = p2
                d_2088_v45_: T4
                nw333_ = C4()
                nw333_.ctor__()
                d_2088_v45_ = nw333_
                d_2089_v46_: _dafny.Seq
                d_2089_v46_ = _dafny.SeqWithoutIsStrInference([d_2088_v45_, d_2088_v45_, d_2088_v45_])
                d_2090_v47_: _dafny.Map
                d_2090_v47_ = _dafny.Map({(d_2086_v43_).f17: ((d_2089_v46_)[default__.safeIndex(p2, len(d_2089_v46_))] if not((d_2086_v43_).f17) else d_2088_v45_)})
                d_2088_v45_ = ((d_2090_v47_)[(self).f19] if ((self).f19) in (d_2090_v47_) else d_2088_v45_)
            d_2091_v48_: _dafny.MultiSet
            d_2091_v48_ = _dafny.MultiSet([(0) - (p2)])
            d_2092_v49_: _dafny.Seq
            d_2092_v49_ = _dafny.SeqWithoutIsStrInference([d_2091_v48_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (len((((self).f14).set(default__.safeIndex(p2, len((self).f14)), (self).f19)).set(default__.safeIndex(960, len(((self).f14).set(default__.safeIndex(p2, len((self).f14)), (self).f19))), (self).f19))))) for d_2093_i7_ in range(default__.abs(865))]))])
            d_2092_v49_ = d_2092_v49_
            d_2094_v50_: D5
            d_2094_v50_ = D5_DC16((self).f19)
            d_2095_v51_: _dafny.Map
            d_2095_v51_ = _dafny.Map({d_2075_cf11_: d_2094_v50_})
            d_2095_v51_ = (d_2095_v51_).set((d_2075_cf11_) | (d_2075_cf11_), d_2094_v50_)
            d_2096_v52_: _dafny.Array
            nw334_ = _dafny.Array(False, 24)
            d_2096_v52_ = nw334_
            index320_ = default__.safeIndex(56, (d_2096_v52_).length(0))
            (d_2096_v52_)[index320_] = True
            index321_ = default__.safeIndex(56, (d_2096_v52_).length(0))
            (d_2096_v52_)[index321_] = (self).f19
        elif True:
            d_2097___mcc_h6_ = source33_.cf17
            d_2098_cf17_ = d_2097___mcc_h6_
            d_2099_v53_: _dafny.Array
            nw335_ = _dafny.Array(int(0), 29)
            d_2099_v53_ = nw335_
            index322_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            (d_2099_v53_)[index322_] = (0) - (len(p1))
            d_2100_v54_: _dafny.Map
            d_2100_v54_ = _dafny.Map({not((self).f19): p2})
            d_2101_v55_: _dafny.Map
            d_2101_v55_ = _dafny.Map({p2: d_2100_v54_})
            index323_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            (d_2099_v53_)[index323_] = default__.fm0((self).f19, (default__.fm62(355, globalState) if (self).f19 else d_2101_v55_), (self).f19, globalState)
            d_2102_v56_: _dafny.Seq
            d_2102_v56_ = _dafny.SeqWithoutIsStrInference([-570])
            d_2103_v57_: _dafny.Set
            d_2103_v57_ = _dafny.Set({(default__.fm1((self).f19, (self).f19, globalState)) + (d_2102_v56_), (_dafny.SeqWithoutIsStrInference([-499, (d_2099_v53_)[default__.safeIndex(477, (d_2099_v53_).length(0))], (d_2099_v53_)[default__.safeIndex(477, (d_2099_v53_).length(0))], p2, p2])) + (_dafny.SeqWithoutIsStrInference([(d_2099_v53_)[default__.safeIndex(477, (d_2099_v53_).length(0))], (d_2099_v53_)[default__.safeIndex(477, (d_2099_v53_).length(0))], (d_2099_v53_)[default__.safeIndex(477, (d_2099_v53_).length(0))]])), d_2102_v56_})
            d_2103_v57_ = d_2103_v57_
            d_2104_v58_: D0
            d_2104_v58_ = D0_DC2(455, (self).f19)
            d_2105_v59_: _dafny.Seq
            d_2106_v60_: bool
            d_2107_v61_: int
            d_2108_v62_: bool
            out111_: _dafny.Seq
            out112_: bool
            out113_: int
            out114_: bool
            out111_, out112_, out113_, out114_ = default__.m0(not((self).f19), (d_2104_v58_ if (self).f19 else d_2104_v58_), True, globalState)
            d_2105_v59_ = out111_
            d_2106_v60_ = out112_
            d_2107_v61_ = out113_
            d_2108_v62_ = out114_
            index324_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            index325_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            rhs277_ = d_2107_v61_
            rhs278_ = (self).f19
            rhs279_ = d_2107_v61_
            lhs168_ = d_2099_v53_
            lhs169_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            lhs170_ = d_2099_v53_
            lhs171_ = default__.safeIndex(477, (d_2099_v53_).length(0))
            lhs168_[lhs169_] = rhs277_
            d_2108_v62_ = rhs278_
            lhs170_[lhs171_] = rhs279_
        d_2109_i8_: int
        d_2109_i8_ = 0
        with _dafny.label("12"):
            while (p2) >= (p2):
                with _dafny.c_label("12"):
                    if (d_2109_i8_) >= (100):
                        raise _dafny.Break("12")
                    d_2109_i8_ = (d_2109_i8_) + (1)
                    r0 = not((p2) == (p2))
                    d_2110_v63_: _dafny.Array
                    nw336_ = _dafny.Array(D12.default()(), 2)
                    d_2110_v63_ = nw336_
                    d_2111_v64_: _dafny.Set
                    d_2111_v64_ = _dafny.Set({d_2110_v63_, d_2110_v63_, d_2110_v63_})
                    d_2111_v64_ = (d_2111_v64_) | (d_2111_v64_)
                    d_2112_v65_: _dafny.Map
                    d_2112_v65_ = _dafny.Map({p2: (self).f19})
                    d_2113_v66_: D7
                    d_2113_v66_ = D7_DC21(d_2112_v65_, not ((self).f19) or ((self).f19), not(((self).f19) == ((self).f19)))
                    d_2113_v66_ = d_2113_v66_
                    d_2114_v67_: int
                    d_2114_v67_ = 443
                    d_2114_v67_ = 126
                    pass
            pass
        d_2115_v68_: str
        d_2115_v68_ = _dafny.CodePoint('u')
        d_2116_v69_: _dafny.Map
        d_2116_v69_ = _dafny.Map({d_2115_v68_: (self).f19})
        d_2117_v70_: _dafny.Map
        d_2117_v70_ = _dafny.Map({(self).f19: len(d_2116_v69_)})
        d_2118_v71_: _dafny.Map
        d_2118_v71_ = _dafny.Map({p2: d_2117_v70_})
        d_2119_v72_: _dafny.MultiSet
        d_2119_v72_ = _dafny.MultiSet([p2])
        if ((d_2119_v72_).set(p2, default__.abs(p2))).issubset((_dafny.MultiSet([default__.fm0(not((self).f19), d_2118_v71_, (self).f19, globalState), p2, 899, len(p1)])) | (d_2119_v72_)):
            r0 = (self).f19
            d_2120_v73_: _dafny.Array
            def lambda111_(d_2121_i9_):
                return False

            init57_ = lambda111_
            nw337_ = _dafny.Array(None, 24)
            for i0_57_ in range(nw337_.length(0)):
                nw337_[i0_57_] = init57_(i0_57_)
            d_2120_v73_ = nw337_
            index326_ = default__.safeIndex(161, (d_2120_v73_).length(0))
            (d_2120_v73_)[index326_] = (self).f19
            index327_ = default__.safeIndex(161, (d_2120_v73_).length(0))
            (d_2120_v73_)[index327_] = ((0) - (p2)) == (p2)
            if (d_2120_v73_)[default__.safeIndex(161, (d_2120_v73_).length(0))]:
                d_2122_v74_: D0
                d_2122_v74_ = D0_DC2(p2, default__.fm2(False, globalState))
                d_2123_v75_: _dafny.Map
                d_2123_v75_ = _dafny.Map({(self).f19: (self).f19})
                d_2124_v76_: C1
                nw338_ = C1()
                nw338_.ctor__((d_2120_v73_)[default__.safeIndex(161, (d_2120_v73_).length(0))], (d_2120_v73_)[default__.safeIndex(161, (d_2120_v73_).length(0))], ((self).f13).set((self).f19, d_2123_v75_), (self).f14, self.f12)
                d_2124_v76_ = nw338_
                d_2125_v77_: _dafny.Seq
                d_2125_v77_ = _dafny.SeqWithoutIsStrInference([d_2124_v76_])
                d_2126_v78_: _dafny.Map
                d_2126_v78_ = _dafny.Map({p2: p2})
                d_2127_v79_: _dafny.Set
                d_2127_v79_ = _dafny.Set({(d_2125_v77_)[default__.safeIndex(len(d_2126_v78_), len(d_2125_v77_))]})
                d_2128_v80_: _dafny.Map
                d_2128_v80_ = _dafny.Map({706: len(d_2127_v79_)})
                d_2129_v81_: _dafny.Seq
                d_2130_v82_: bool
                d_2131_v83_: int
                d_2132_v84_: bool
                out115_: _dafny.Seq
                out116_: bool
                out117_: int
                out118_: bool
                out115_, out116_, out117_, out118_ = default__.m0((self).f19, d_2122_v74_, (d_2126_v78_) != (d_2128_v80_), globalState)
                d_2129_v81_ = out115_
                d_2130_v82_ = out116_
                d_2131_v83_ = out117_
                d_2132_v84_ = out118_
                d_2133_v85_: _dafny.Seq
                d_2133_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaj"))
                d_2134_v86_: D12
                d_2134_v86_ = D12_DC37(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruokmp")), d_2132_v84_)
                def iife140_(_pat_let37_0):
                    def iife141_(d_2135_dt__update__tmp_h2_):
                        def iife142_(_pat_let38_0):
                            def iife143_(d_2136_dt__update_hcf60_h0_):
                                return D12_DC37((d_2135_dt__update__tmp_h2_).cf59, d_2136_dt__update_hcf60_h0_)
                            return iife143_(_pat_let38_0)
                        return iife142_(True)
                    return iife141_(_pat_let37_0)
                d_2133_v85_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unebn"))) + ((iife140_(d_2134_v86_)).cf59)
                d_2137_v87_: _dafny.Set
                d_2137_v87_ = _dafny.Set({d_2120_v73_})
                r0 = (d_2137_v87_).isdisjoint((d_2137_v87_) | (d_2137_v87_))
                d_2129_v81_ = d_2129_v81_
                d_2126_v78_ = (d_2126_v78_).set(len(d_2117_v70_), (0) - (p2))
            elif True:
                index328_ = default__.safeIndex(161, (d_2120_v73_).length(0))
                (d_2120_v73_)[index328_] = default__.fm2(default__.fm2(not((d_2120_v73_)[default__.safeIndex(161, (d_2120_v73_).length(0))]), globalState), globalState)
                d_2138_v88_: int
                d_2138_v88_ = 350
                d_2138_v88_ = (self).fm5(not((self).f19), globalState)
                d_2139_v89_: C3
                nw339_ = C3()
                nw339_.ctor__(True, self.f12, (d_2120_v73_)[default__.safeIndex(161, (d_2120_v73_).length(0))], (self).f13, (self).f14)
                d_2139_v89_ = nw339_
                (globalState).f4 = not(not(((self).f14)[default__.safeIndex(22, len((self).f14))]))
                d_2140_v90_: _dafny.Map
                d_2140_v90_ = _dafny.Map({p2: p2})
                d_2141_v91_: _dafny.Seq
                d_2141_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2138_v88_: len(p1)}), d_2140_v90_, d_2140_v90_])
                d_2142_v92_: _dafny.Array
                nw340_ = _dafny.Array(None, 4)
                d_2142_v92_ = nw340_
                d_2143_v93_: C1
                nw341_ = C1()
                nw341_.ctor__((d_2139_v89_).f20, True, (self).f13, (self).f14, self.f12)
                d_2143_v93_ = nw341_
                index329_ = default__.safeIndex(987, (d_2142_v92_).length(0))
                (d_2142_v92_)[index329_] = d_2143_v93_
                d_2144_v94_: _dafny.MultiSet
                d_2144_v94_ = _dafny.MultiSet([(d_2119_v72_) | (_dafny.MultiSet([(0) - ((0) - (d_2138_v88_))])), d_2119_v72_])
                d_2145_v95_: _dafny.Seq
                d_2145_v95_ = _dafny.SeqWithoutIsStrInference([(d_2143_v93_).fm5(False, globalState), d_2138_v88_])
                index330_ = default__.safeIndex(987, (d_2142_v92_).length(0))
                rhs280_ = ((self).f14)[default__.safeIndex(260, len((self).f14))]
                rhs281_ = _dafny.SeqWithoutIsStrInference([d_2140_v90_, _dafny.Map({(0) - (p2): d_2138_v88_}), (d_2140_v90_) | (d_2140_v90_), d_2140_v90_, d_2140_v90_])
                rhs282_ = d_2143_v93_
                rhs283_ = (_dafny.MultiSet([_dafny.MultiSet(d_2145_v95_), d_2119_v72_]) if ((self).f19) or (False) else d_2144_v94_)
                rhs284_ = d_2120_v73_
                lhs172_ = globalState
                lhs173_ = d_2142_v92_
                lhs174_ = default__.safeIndex(987, (d_2142_v92_).length(0))
                lhs172_.f4 = rhs280_
                d_2141_v91_ = rhs281_
                lhs173_[lhs174_] = rhs282_
                d_2144_v94_ = rhs283_
                d_2120_v73_ = rhs284_
            d_2146_v96_: int
            d_2146_v96_ = 938
            d_2146_v96_ = (len(p1)) + (p2)
            d_2147_v97_: _dafny.Seq
            d_2147_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            d_2147_v97_ = (d_2147_v97_) + ((p1) + (d_2147_v97_))
        elif True:
            r0 = (p2) != (p2)
            d_2148_v98_: int
            d_2148_v98_ = 336
            d_2149_v99_: _dafny.Map
            d_2149_v99_ = _dafny.Map({p2: p2})
            d_2150_v100_: _dafny.Seq
            d_2150_v100_ = _dafny.SeqWithoutIsStrInference([((d_2149_v99_)[p2] if (p2) in (d_2149_v99_) else p2), p2])
            d_2148_v98_ = (d_2150_v100_)[default__.safeIndex((len(p1)) + (len((self).f14)), len(d_2150_v100_))]
            d_2151_v101_: _dafny.Map
            d_2151_v101_ = _dafny.Map({((self).f14) + ((self).f14): p1})
            d_2151_v101_ = (d_2151_v101_).set(((self).f14 if not((self).f19) else (self).f14), _dafny.SeqWithoutIsStrInference([d_2115_v68_ for d_2152_i10_ in range(default__.abs(139))]))
            d_2148_v98_ = p2
            d_2153_v102_: _dafny.Array
            def lambda112_(d_2154_p1_):
                def lambda113_(d_2155_i11_):
                    return d_2154_p1_

                return lambda113_

            init58_ = lambda112_(p1)
            nw342_ = _dafny.Array(None, 8)
            for i0_58_ in range(nw342_.length(0)):
                nw342_[i0_58_] = init58_(i0_58_)
            d_2153_v102_ = nw342_
            d_2156_v104_: _dafny.Seq
            d_2156_v104_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpwcbgmj"))])
            d_2157_v105_: _dafny.Array
            def lambda114_(d_2158_i12_):
                return (self).f19

            init59_ = lambda114_
            nw343_ = _dafny.Array(None, 17)
            for i0_59_ in range(nw343_.length(0)):
                nw343_[i0_59_] = init59_(i0_59_)
            d_2157_v105_ = nw343_
            d_2159_v106_: _dafny.Set
            d_2159_v106_ = _dafny.Set({d_2157_v105_})
            d_2160_v107_: D12
            def iife144_():
                coll66_ = _dafny.Map()
                compr_66_: _dafny.Seq
                for compr_66_ in (d_2156_v104_).Elements:
                    d_2161_v103_: _dafny.Seq = compr_66_
                    if (d_2161_v103_) in (d_2156_v104_):
                        coll66_[d_2161_v103_] = d_2150_v100_
                return _dafny.Map(coll66_)
            d_2160_v107_ = D12_DC36(d_2153_v102_, default__.fm0((self).f19, (_dafny.Map({-214: d_2117_v70_})).set(p2, d_2117_v70_), (self).f19, globalState), iife144_()
, d_2159_v106_)
            d_2162_v108_: _dafny.Map
            d_2162_v108_ = _dafny.Map({d_2148_v98_: d_2160_v107_})
            d_2163_v109_: _dafny.MultiSet
            d_2163_v109_ = _dafny.MultiSet([(self).f19])
            d_2164_v110_: _dafny.Map
            d_2164_v110_ = _dafny.Map({(self).f19: False})
            d_2165_v111_: _dafny.Array
            nw344_ = _dafny.Array(None, 14)
            nw344_[int(0)] = d_2148_v98_
            nw344_[int(1)] = len(_dafny.Map({(self).f19: p2}))
            nw344_[int(2)] = -405
            nw344_[int(3)] = d_2148_v98_
            nw344_[int(4)] = ((d_2117_v70_)[(self).f19] if ((self).f19) in (d_2117_v70_) else d_2148_v98_)
            nw344_[int(5)] = len((self).f14)
            nw344_[int(6)] = ((d_2149_v99_)[d_2148_v98_] if (d_2148_v98_) in (d_2149_v99_) else d_2148_v98_)
            nw344_[int(7)] = p2
            nw344_[int(8)] = ((d_2149_v99_)[d_2148_v98_] if (d_2148_v98_) in (d_2149_v99_) else len(d_2162_v108_))
            nw344_[int(9)] = (d_2148_v98_) - (((d_2163_v109_)[False] if (False) in (d_2163_v109_) else len(p1)))
            nw344_[int(10)] = len(d_2164_v110_)
            nw344_[int(11)] = p2
            nw344_[int(12)] = (len((self).f14)) + (len(d_2150_v100_))
            nw344_[int(13)] = p2
            d_2165_v111_ = nw344_
            d_2166_v112_: _dafny.Map
            d_2166_v112_ = _dafny.Map({d_2148_v98_: d_2157_v105_})
            rhs285_ = d_2148_v98_
            rhs286_ = d_2148_v98_
            rhs287_ = (((d_2163_v109_)[(self).f19] if ((self).f19) in (d_2163_v109_) else len(d_2166_v112_))) - (len(d_2149_v99_))
            rhs288_ = d_2165_v111_
            rhs289_ = d_2115_v68_
            d_2148_v98_ = rhs285_
            d_2148_v98_ = rhs286_
            d_2148_v98_ = rhs287_
            d_2165_v111_ = rhs288_
            d_2115_v68_ = rhs289_
        d_2167_i13_: int
        d_2167_i13_ = 0
        with _dafny.label("13"):
            while True:
                with _dafny.c_label("13"):
                    if (d_2167_i13_) >= (100):
                        raise _dafny.Break("13")
                    d_2167_i13_ = (d_2167_i13_) + (1)
                    d_2168_v113_: int
                    d_2168_v113_ = 997
                    d_2168_v113_ = d_2168_v113_
                    d_2169_v114_: D4
                    d_2169_v114_ = D4_DC14((self).f19)
                    d_2170_v115_: _dafny.Map
                    d_2170_v115_ = _dafny.Map({(self).f19: default__.fm40(d_2169_v114_, 373, (self).f19, (self).f19, globalState)})
                    d_2171_v116_: _dafny.Seq
                    d_2171_v116_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14, (self).f14, (self).f14, (self).f14])
                    d_2172_v117_: _dafny.Set
                    d_2172_v117_ = _dafny.Set({(d_2171_v116_)[default__.safeIndex(len(p0), len(d_2171_v116_))]})
                    d_2170_v115_ = (d_2170_v115_).set((p2) < (len(d_2172_v117_)), p1)
                    d_2173_v118_: _dafny.Seq
                    d_2173_v118_ = _dafny.SeqWithoutIsStrInference([d_2115_v68_])
                    d_2174_v119_: _dafny.Map
                    d_2174_v119_ = _dafny.Map({(self).f19: ((d_2116_v69_)[d_2115_v68_] if (d_2115_v68_) in (d_2116_v69_) else (self).f19)})
                    d_2173_v118_ = (default__.fm29((self).f19, (self).f19, (self).f19, d_2168_v113_, globalState)) + ((default__.fm40(d_2169_v114_, p2, ((d_2174_v119_)[(self).f19] if ((self).f19) in (d_2174_v119_) else (self).f19), (self).f19, globalState) if (self).f19 else _dafny.SeqWithoutIsStrInference([d_2115_v68_, d_2115_v68_, d_2115_v68_])))
                    if (not ((self).f19) or (False) if (self).f19 else True):
                        d_2168_v113_ = d_2168_v113_
                        d_2175_v120_: _dafny.Seq
                        d_2175_v120_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jchc")))), p2, len(d_2116_v69_)])
                        (globalState).f4 = ((self).f19) == ((d_2175_v120_) == (_dafny.SeqWithoutIsStrInference([p2])))
                        (globalState).f4 = (self).f19
                        d_2176_v121_: C1
                        nw345_ = C1()
                        nw345_.ctor__((self).f19, (self).f19, (self).f13, ((self).f14) + (_dafny.SeqWithoutIsStrInference([(self).f19])), self.f12)
                        d_2176_v121_ = nw345_
                        d_2177_v122_: _dafny.Array
                        nw346_ = _dafny.Array(int(0), 16)
                        d_2177_v122_ = nw346_
                        index331_ = default__.safeIndex(34, (d_2177_v122_).length(0))
                        (d_2177_v122_)[index331_] = p2
                        index332_ = default__.safeIndex(34, (d_2177_v122_).length(0))
                        (d_2177_v122_)[index332_] = ((d_2119_v72_).intersection(d_2119_v72_)).cardinality
                    elif True:
                        d_2178_v123_: _dafny.Array
                        nw347_ = _dafny.Array(_dafny.Map({}), 2)
                        d_2178_v123_ = nw347_
                        d_2179_v124_: _dafny.Array
                        nw348_ = _dafny.Array(None, 6)
                        nw348_[int(0)] = d_2178_v123_
                        nw348_[int(1)] = d_2178_v123_
                        nw348_[int(2)] = d_2178_v123_
                        nw348_[int(3)] = d_2178_v123_
                        nw348_[int(4)] = d_2178_v123_
                        nw348_[int(5)] = d_2178_v123_
                        d_2179_v124_ = nw348_
                        index333_ = default__.safeIndex(226, (d_2179_v124_).length(0))
                        (d_2179_v124_)[index333_] = d_2178_v123_
                        index334_ = default__.safeIndex(226, (d_2179_v124_).length(0))
                        (d_2179_v124_)[index334_] = d_2178_v123_
                        d_2180_v125_: C1
                        nw349_ = C1()
                        nw349_.ctor__((self).f19, (self).f19, (self).f13, (self).f14, self.f12)
                        d_2180_v125_ = nw349_
                        d_2181_v126_: _dafny.Array
                        nw350_ = _dafny.Array(False, 15)
                        d_2181_v126_ = nw350_
                        index335_ = default__.safeIndex(520, (d_2181_v126_).length(0))
                        (d_2181_v126_)[index335_] = (d_2115_v68_) not in (p1)
                        d_2182_v127_: _dafny.Array
                        nw351_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                        d_2182_v127_ = nw351_
                        index336_ = default__.safeIndex(834, (d_2182_v127_).length(0))
                        (d_2182_v127_)[index336_] = d_2173_v118_
                        index337_ = default__.safeIndex(520, (d_2181_v126_).length(0))
                        index338_ = default__.safeIndex(834, (d_2182_v127_).length(0))
                        rhs290_ = default__.fm56(_dafny.CodePoint('a'), globalState)
                        rhs291_ = default__.safeDivisionInt(d_2168_v113_, d_2168_v113_)
                        rhs292_ = ((d_2180_v125_).f17 if (d_2180_v125_).f17 else not((d_2180_v125_).f17))
                        rhs293_ = (self).f19
                        rhs294_ = p1
                        lhs175_ = d_2181_v126_
                        lhs176_ = default__.safeIndex(520, (d_2181_v126_).length(0))
                        lhs177_ = d_2182_v127_
                        lhs178_ = default__.safeIndex(834, (d_2182_v127_).length(0))
                        d_2119_v72_ = rhs290_
                        d_2168_v113_ = rhs291_
                        r0 = rhs292_
                        lhs175_[lhs176_] = rhs293_
                        lhs177_[lhs178_] = rhs294_
                        d_2183_v128_: _dafny.Array
                        nw352_ = _dafny.Array(int(0), 18)
                        d_2183_v128_ = nw352_
                        index339_ = default__.safeIndex(877, (d_2183_v128_).length(0))
                        (d_2183_v128_)[index339_] = default__.safeModuloInt(d_2168_v113_, d_2168_v113_)
                        d_2184_v129_: _dafny.Set
                        d_2184_v129_ = _dafny.Set({d_2181_v126_})
                        index340_ = default__.safeIndex(877, (d_2183_v128_).length(0))
                        rhs295_ = d_2168_v113_
                        rhs296_ = (len((d_2184_v129_) | (_dafny.Set({d_2181_v126_})))) <= ((d_2168_v113_ if not((self).f19) else 981))
                        lhs179_ = d_2183_v128_
                        lhs180_ = default__.safeIndex(877, (d_2183_v128_).length(0))
                        lhs181_ = globalState
                        lhs179_[lhs180_] = rhs295_
                        lhs181_.f4 = rhs296_
                        index341_ = default__.safeIndex(520, (d_2181_v126_).length(0))
                        (d_2181_v126_)[index341_] = default__.fm2(d_2180_v125_.f18, globalState)
                    pass
            pass
        (globalState).f4 = True
        d_2185_v130_: _dafny.Map
        d_2185_v130_ = _dafny.Map({p2: (self).f19})
        d_2186_v131_: D7
        d_2186_v131_ = D7_DC21(d_2185_v130_, (p2) == (p2), (not((self).f19)) and ((self).f19))
        d_2186_v131_ = d_2186_v131_
        r0 = (self).f19
        return r0

    def m6(self, globalState):
        r0: bool = False
        if not ((self).f19) or ((self).f19):
            d_2187_v0_: D8
            d_2187_v0_ = D8_DC26((self).f19)
            def iife146_(_pat_let40_0):
                def iife147_(d_2188_dt__update__tmp_h0_):
                    def iife148_(_pat_let41_0):
                        def iife149_(d_2189_dt__update_hcf41_h0_):
                            return D8_DC26(d_2189_dt__update_hcf41_h0_)
                        return iife149_(_pat_let41_0)
                    return iife148_((self).f19)
                return iife147_(_pat_let40_0)
            def iife145_(_pat_let39_0):
                def iife150_(d_2190_dt__update__tmp_h1_):
                    def iife151_(_pat_let42_0):
                        def iife152_(d_2191_dt__update_hcf41_h1_):
                            return D8_DC26(d_2191_dt__update_hcf41_h1_)
                        return iife152_(_pat_let42_0)
                    return iife151_((self).f19)
                return iife150_(_pat_let39_0)
            source35_ = iife145_(iife146_(d_2187_v0_))
            if source35_.is_DC24:
                d_2192___mcc_h0_ = source35_.cf36
                d_2193_cf36_ = d_2192___mcc_h0_
                d_2193_cf36_ = not(d_2193_cf36_)
                d_2194_v1_: _dafny.Seq
                d_2194_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkgacn"))
                d_2194_v1_ = d_2194_v1_
                (globalState).f4 = d_2193_cf36_
                d_2195_v2_: D0
                d_2195_v2_ = D0_DC0(not((self).f19))
                r0 = (d_2195_v2_).cf0
            elif source35_.is_DC25:
                d_2196___mcc_h1_ = source35_.cf37
                d_2197___mcc_h2_ = source35_.cf38
                d_2198___mcc_h3_ = source35_.cf39
                d_2199___mcc_h4_ = source35_.cf40
                d_2200_cf40_ = d_2199___mcc_h4_
                d_2201_cf39_ = d_2198___mcc_h3_
                d_2202_cf38_ = d_2197___mcc_h2_
                d_2203_cf37_ = d_2196___mcc_h1_
                d_2202_cf38_ = d_2202_cf38_
                (globalState).f4 = default__.fm2(default__.fm2((self).f19, globalState), globalState)
                d_2204_v3_: _dafny.Array
                def lambda115_(d_2205_i0_):
                    return True

                init60_ = lambda115_
                nw353_ = _dafny.Array(None, 11)
                for i0_60_ in range(nw353_.length(0)):
                    nw353_[i0_60_] = init60_(i0_60_)
                d_2204_v3_ = nw353_
                d_2206_v4_: _dafny.Set
                d_2206_v4_ = _dafny.Set({d_2204_v3_, d_2204_v3_})
                d_2201_cf39_ = len((d_2206_v4_) - ((d_2206_v4_) | (d_2206_v4_)))
                (globalState).f4 = not(((self).f19) and ((self).f19))
            elif source35_.is_DC26:
                d_2207___mcc_h5_ = source35_.cf41
                d_2208_cf41_ = d_2207___mcc_h5_
                d_2209_v5_: int
                d_2209_v5_ = 824
                d_2209_v5_ = d_2209_v5_
                d_2210_v6_: str
                d_2210_v6_ = _dafny.CodePoint('i')
                d_2210_v6_ = d_2210_v6_
                d_2211_v7_: _dafny.Set
                d_2211_v7_ = _dafny.Set({d_2208_cf41_})
                d_2212_v8_: _dafny.Set
                d_2212_v8_ = d_2211_v7_
                rhs297_ = (d_2211_v7_).isdisjoint((d_2212_v8_))
                rhs298_ = default__.safeDivisionInt(d_2209_v5_, default__.safeModuloInt(d_2209_v5_, 41))
                lhs182_ = globalState
                lhs182_.f4 = rhs297_
                d_2209_v5_ = rhs298_
                d_2213_v9_: T2
                nw354_ = C3()
                nw354_.ctor__(d_2208_cf41_, self.f12, d_2208_cf41_, (self).f13, (self).f14)
                d_2213_v9_ = nw354_
                d_2214_v10_: D14
                d_2214_v10_ = D14_DC40(d_2213_v9_)
                d_2215_v11_: _dafny.Map
                d_2215_v11_ = _dafny.Map({True: d_2213_v9_})
                pat_let_tv36_ = d_2215_v11_
                pat_let_tv37_ = d_2215_v11_
                pat_let_tv38_ = d_2213_v9_
                def iife153_(_pat_let43_0):
                    def iife154_(d_2216_dt__update__tmp_h2_):
                        def iife155_(_pat_let44_0):
                            def iife156_(d_2217_dt__update_hcf63_h0_):
                                return D14_DC40(d_2217_dt__update_hcf63_h0_)
                            return iife156_(_pat_let44_0)
                        return iife155_(((pat_let_tv36_)[(self).f19] if ((self).f19) in (pat_let_tv37_) else pat_let_tv38_))
                    return iife154_(_pat_let43_0)
                d_2214_v10_ = iife153_(d_2214_v10_)
            elif source35_.is_DC23:
                d_2218___mcc_h6_ = source35_.cf35
                d_2219_cf35_ = d_2218___mcc_h6_
                d_2220_v12_: _dafny.Array
                def lambda116_(d_2221_i1_):
                    return (self).f19

                init61_ = lambda116_
                nw355_ = _dafny.Array(None, 18)
                for i0_61_ in range(nw355_.length(0)):
                    nw355_[i0_61_] = init61_(i0_61_)
                d_2220_v12_ = nw355_
                index342_ = default__.safeIndex(941, (d_2220_v12_).length(0))
                (d_2220_v12_)[index342_] = ((self).f19 if (self).f19 else (self).f19)
                index343_ = default__.safeIndex(941, (d_2220_v12_).length(0))
                (d_2220_v12_)[index343_] = (self).f19
                d_2222_v13_: C0
                nw356_ = C0()
                nw356_.ctor__()
                d_2222_v13_ = nw356_
                d_2223_v14_: int
                d_2223_v14_ = -413
                d_2224_v15_: _dafny.MultiSet
                d_2224_v15_ = _dafny.MultiSet([(0) - (d_2223_v14_)])
                d_2223_v14_ = (d_2223_v14_) * ((d_2224_v15_).cardinality)
                d_2225_v16_: _dafny.Seq
                d_2225_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhlgooi"))
                d_2226_v17_: _dafny.Seq
                d_2226_v17_ = _dafny.SeqWithoutIsStrInference([(d_2225_v16_) <= (d_2225_v16_)])
                d_2226_v17_ = (self).f14
            elif True:
                d_2227___mcc_h7_ = source35_.cf42
                d_2228_cf42_ = d_2227___mcc_h7_
                d_2229_v18_: str
                d_2229_v18_ = _dafny.CodePoint('d')
                d_2229_v18_ = d_2229_v18_
                (globalState).f4 = (self).f19
                d_2230_v19_: int
                d_2230_v19_ = -786
                d_2230_v19_ = 756
                d_2231_v20_: _dafny.Array
                nw357_ = _dafny.Array(_dafny.Set({}), 17)
                d_2231_v20_ = nw357_
                d_2232_v21_: _dafny.MultiSet
                d_2232_v21_ = _dafny.MultiSet([d_2230_v19_])
                d_2233_v22_: _dafny.Map
                d_2233_v22_ = _dafny.Map({True: (0) - (d_2230_v19_)})
                d_2234_v23_: _dafny.Set
                d_2234_v23_ = _dafny.Set({d_2230_v19_})
                d_2235_v24_: _dafny.Seq
                d_2235_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gb"))
                d_2236_v25_: _dafny.MultiSet
                d_2236_v25_ = _dafny.MultiSet([(self).f19, (self).f19])
                d_2237_v26_: _dafny.Seq
                d_2237_v26_ = _dafny.SeqWithoutIsStrInference([d_2230_v19_])
                d_2238_v27_: _dafny.Map
                d_2238_v27_ = _dafny.Map({(self).f19: d_2229_v18_})
                d_2239_v28_: _dafny.MultiSet
                d_2239_v28_ = _dafny.MultiSet([d_2238_v27_, d_2238_v27_, d_2238_v27_, d_2238_v27_, _dafny.Map({(self).f19: d_2229_v18_})])
                d_2240_v29_: _dafny.Array
                nw358_ = _dafny.Array(None, 27)
                nw358_[int(0)] = len(_dafny.Set({d_2230_v19_}))
                nw358_[int(1)] = d_2230_v19_
                nw358_[int(2)] = d_2230_v19_
                nw358_[int(3)] = d_2230_v19_
                nw358_[int(4)] = d_2230_v19_
                nw358_[int(5)] = (d_2232_v21_).cardinality
                nw358_[int(6)] = 618
                nw358_[int(7)] = ((d_2233_v22_)[not((self).f19)] if (not((self).f19)) in (d_2233_v22_) else d_2230_v19_)
                nw358_[int(8)] = 303
                nw358_[int(9)] = d_2230_v19_
                nw358_[int(10)] = len(d_2234_v23_)
                nw358_[int(11)] = 851
                nw358_[int(12)] = 582
                nw358_[int(13)] = d_2230_v19_
                nw358_[int(14)] = d_2230_v19_
                nw358_[int(15)] = d_2230_v19_
                nw358_[int(16)] = d_2230_v19_
                nw358_[int(17)] = len(d_2235_v24_)
                nw358_[int(18)] = ((d_2236_v25_)[not((self).f19)] if (not((self).f19)) in (d_2236_v25_) else len(d_2237_v26_))
                nw358_[int(19)] = ((d_2236_v25_).set((self).f19, default__.abs((0) - (((d_2233_v22_)[False] if (False) in (d_2233_v22_) else (d_2239_v28_).cardinality))))).cardinality
                nw358_[int(20)] = len(d_2237_v26_)
                nw358_[int(21)] = d_2230_v19_
                nw358_[int(22)] = d_2230_v19_
                nw358_[int(23)] = 819
                nw358_[int(24)] = len(d_2235_v24_)
                nw358_[int(25)] = d_2230_v19_
                nw358_[int(26)] = d_2230_v19_
                d_2240_v29_ = nw358_
                d_2241_v30_: _dafny.Set
                d_2241_v30_ = _dafny.Set({d_2240_v29_})
                index344_ = default__.safeIndex(643, (d_2231_v20_).length(0))
                (d_2231_v20_)[index344_] = (d_2241_v30_ if (self).f19 else d_2241_v30_)
                index345_ = default__.safeIndex(643, (d_2231_v20_).length(0))
                (d_2231_v20_)[index345_] = (d_2241_v30_) | (d_2241_v30_)
            d_2242_v31_: C4
            nw359_ = C4()
            nw359_.ctor__()
            d_2242_v31_ = nw359_
            d_2243_v32_: int
            d_2243_v32_ = 522
            d_2244_v33_: _dafny.Map
            d_2244_v33_ = _dafny.Map({d_2243_v32_: d_2243_v32_})
            d_2244_v33_ = (d_2244_v33_).set((0) - (d_2243_v32_), d_2243_v32_)
            d_2245_v34_: str
            d_2245_v34_ = _dafny.CodePoint('t')
            d_2246_v35_: _dafny.Seq
            d_2246_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uonlgggk"))
            d_2247_v36_: _dafny.Map
            d_2247_v36_ = _dafny.Map({d_2245_v34_: d_2246_v35_})
            d_2248_v37_: _dafny.Set
            d_2248_v37_ = _dafny.Set({d_2242_v31_})
            d_2249_v38_: _dafny.Map
            d_2249_v38_ = _dafny.Map({(self).f19: True})
            d_2250_v39_: _dafny.Seq
            d_2250_v39_ = _dafny.SeqWithoutIsStrInference([d_2243_v32_])
            d_2251_v40_: _dafny.Map
            d_2251_v40_ = _dafny.Map({(self).f19: (d_2250_v39_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2245_v34_])), len(d_2250_v39_))]})
            d_2252_v41_: _dafny.Map
            d_2252_v41_ = _dafny.Map({len((self).f14): d_2251_v40_})
            d_2253_v42_: _dafny.Map
            d_2253_v42_ = _dafny.Map({d_2243_v32_: (self).f19})
            d_2254_v43_: D7
            d_2254_v43_ = D7_DC21(d_2253_v42_, not((self).f19), (self).f19)
            d_2255_v44_: _dafny.Set
            d_2255_v44_ = _dafny.Set({d_2243_v32_, (0) - (d_2243_v32_)})
            d_2256_v45_: _dafny.Array
            nw360_ = _dafny.Array(None, 21)
            nw360_[int(0)] = d_2243_v32_
            nw360_[int(1)] = len(d_2248_v37_)
            nw360_[int(2)] = d_2243_v32_
            nw360_[int(3)] = d_2243_v32_
            nw360_[int(4)] = default__.fm0(((d_2249_v38_)[(self).f19] if ((self).f19) in (d_2249_v38_) else (self).f19), d_2252_v41_, (self).f19, globalState)
            nw360_[int(5)] = d_2243_v32_
            nw360_[int(6)] = d_2243_v32_
            nw360_[int(7)] = d_2243_v32_
            nw360_[int(8)] = d_2243_v32_
            nw360_[int(9)] = d_2243_v32_
            nw360_[int(10)] = default__.fm0((self).f19, d_2252_v41_, (self).f19, globalState)
            nw360_[int(11)] = ((d_2244_v33_)[len(_dafny.Map({(d_2254_v43_).cf32: d_2245_v34_}))] if (len(_dafny.Map({(d_2254_v43_).cf32: d_2245_v34_}))) in (d_2244_v33_) else d_2243_v32_)
            nw360_[int(12)] = (0) - (d_2243_v32_)
            nw360_[int(13)] = len(d_2246_v35_)
            nw360_[int(14)] = d_2243_v32_
            nw360_[int(15)] = len(d_2255_v44_)
            nw360_[int(16)] = d_2243_v32_
            nw360_[int(17)] = len((self).f14)
            nw360_[int(18)] = d_2243_v32_
            nw360_[int(19)] = d_2243_v32_
            nw360_[int(20)] = len((d_2250_v39_).set(default__.safeIndex((0) - (len(d_2246_v35_)), len(d_2250_v39_)), len(d_2250_v39_)))
            d_2256_v45_ = nw360_
            d_2257_v46_: _dafny.Set
            d_2257_v46_ = _dafny.Set({d_2256_v45_})
            d_2258_v47_: _dafny.Map
            d_2258_v47_ = _dafny.Map({((d_2247_v36_)[d_2245_v34_] if (d_2245_v34_) in (d_2247_v36_) else d_2246_v35_): d_2257_v46_})
            d_2258_v47_ = (d_2258_v47_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hya")), (d_2257_v46_) | (d_2257_v46_))
            d_2243_v32_ = (default__.fm0((self).f19, d_2252_v41_, True, globalState)) - (len(d_2250_v39_))
        elif True:
            d_2259_v48_: _dafny.Array
            def lambda117_(d_2260_i2_):
                return D8_DC24((self).f19)

            init62_ = lambda117_
            nw361_ = _dafny.Array(None, 4)
            for i0_62_ in range(nw361_.length(0)):
                nw361_[i0_62_] = init62_(i0_62_)
            d_2259_v48_ = nw361_
            d_2261_v49_: D8
            d_2261_v49_ = D8_DC24((self).f19)
            index346_ = default__.safeIndex(659, (d_2259_v48_).length(0))
            (d_2259_v48_)[index346_] = d_2261_v49_
            d_2262_v50_: _dafny.Array
            nw362_ = _dafny.Array(int(0), 11)
            d_2262_v50_ = nw362_
            index347_ = default__.safeIndex(659, (d_2259_v48_).length(0))
            rhs299_ = D8_DC24((self).f19)
            rhs300_ = d_2262_v50_
            lhs183_ = d_2259_v48_
            lhs184_ = default__.safeIndex(659, (d_2259_v48_).length(0))
            lhs183_[lhs184_] = rhs299_
            d_2262_v50_ = rhs300_
            d_2263_v51_: _dafny.Array
            nw363_ = _dafny.Array(_dafny.Map({}), 5)
            d_2263_v51_ = nw363_
            d_2264_v52_: int
            d_2264_v52_ = 815
            d_2265_v53_: _dafny.Map
            d_2265_v53_ = _dafny.Map({(self).f19: d_2264_v52_})
            d_2266_v54_: _dafny.Set
            d_2266_v54_ = _dafny.Set({((d_2265_v53_)[(self).f19] if ((self).f19) in (d_2265_v53_) else d_2264_v52_)})
            d_2267_v55_: _dafny.Map
            d_2267_v55_ = _dafny.Map({(self).f19: D6_DC17(d_2266_v54_)})
            index348_ = default__.safeIndex(934, (d_2263_v51_).length(0))
            (d_2263_v51_)[index348_] = d_2267_v55_
            index349_ = default__.safeIndex(934, (d_2263_v51_).length(0))
            (d_2263_v51_)[index349_] = (d_2267_v55_) | (d_2267_v55_)
            d_2268_v56_: C4
            nw364_ = C4()
            nw364_.ctor__()
            d_2268_v56_ = nw364_
            d_2269_v57_: _dafny.Array
            def lambda118_(d_2270_v52_):
                def lambda119_(d_2271_i3_):
                    return (D4_DC13((self).f19, (self).f19, (0) - (d_2270_v52_), d_2270_v52_, (self).f19)).cf19

                return lambda119_

            init63_ = lambda118_(d_2264_v52_)
            nw365_ = _dafny.Array(None, 17)
            for i0_63_ in range(nw365_.length(0)):
                nw365_[i0_63_] = init63_(i0_63_)
            d_2269_v57_ = nw365_
            d_2272_v58_: _dafny.Set
            d_2272_v58_ = _dafny.Set({(self).f19})
            index350_ = default__.safeIndex(164, (d_2269_v57_).length(0))
            (d_2269_v57_)[index350_] = (d_2272_v58_).issubset(d_2272_v58_)
            index351_ = default__.safeIndex(164, (d_2269_v57_).length(0))
            (d_2269_v57_)[index351_] = (self).f19
            d_2273_v59_: D8
            d_2273_v59_ = D8_DC25((self).f19, d_2264_v52_, d_2264_v52_, d_2265_v53_)
            index352_ = default__.safeIndex(187, (d_2262_v50_).length(0))
            (d_2262_v50_)[index352_] = (d_2273_v59_).cf38
            index353_ = default__.safeIndex(187, (d_2262_v50_).length(0))
            (d_2262_v50_)[index353_] = (d_2264_v52_ if ((self).f19) and (True) else d_2264_v52_)
        if (self).f19:
            (globalState).f4 = (self).f19
            d_2274_v60_: int
            d_2274_v60_ = 855
            d_2275_v61_: _dafny.Set
            d_2275_v61_ = _dafny.Set({_dafny.CodePoint('v')})
            d_2276_v62_: _dafny.Set
            d_2276_v62_ = _dafny.Set({(351) * (d_2274_v60_), (d_2274_v60_) * (len(d_2275_v61_)), d_2274_v60_})
            d_2277_v63_: _dafny.Seq
            d_2277_v63_ = _dafny.SeqWithoutIsStrInference([d_2274_v60_, d_2274_v60_])
            d_2276_v62_ = (self).fm10((self).f19, (_dafny.Set({265, d_2274_v60_})).intersection(d_2276_v62_), (d_2277_v63_) + (d_2277_v63_), d_2274_v60_, globalState)
            d_2278_v64_: _dafny.Map
            d_2278_v64_ = _dafny.Map({(d_2274_v60_) * (len(_dafny.Map({(self).f19: 489}))): ((self).f14) + ((self).f14)})
            d_2278_v64_ = (d_2278_v64_).set(d_2274_v60_, ((self).f14) + ((self).f14))
            d_2279_v65_: str
            d_2279_v65_ = _dafny.CodePoint('p')
            d_2279_v65_ = d_2279_v65_
            (globalState).f4 = (self).f19
        elif True:
            d_2280_v66_: _dafny.MultiSet
            d_2280_v66_ = _dafny.MultiSet([(self).f19, (self).f19])
            d_2281_v67_: D0
            d_2281_v67_ = D0_DC1()
            d_2282_v68_: int
            d_2282_v68_ = 421
            source36_ = (d_2281_v67_ if (d_2280_v66_).ispropersubset(d_2280_v66_) else default__.fm63(d_2282_v68_, globalState))
            if source36_.is_DC1:
                d_2283_v69_: _dafny.Seq
                d_2283_v69_ = _dafny.SeqWithoutIsStrInference([not((self).f19), not((self).f19)])
                d_2284_v70_: _dafny.MultiSet
                d_2284_v70_ = _dafny.MultiSet([d_2282_v68_])
                d_2285_v71_: _dafny.Seq
                d_2285_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwkupqk"))
                d_2286_v72_: _dafny.Seq
                d_2286_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqexl")), d_2285_v71_])
                d_2287_v73_: _dafny.Seq
                d_2287_v73_ = _dafny.SeqWithoutIsStrInference([(d_2286_v72_)[default__.safeIndex(d_2282_v68_, len(d_2286_v72_))], d_2285_v71_, (d_2286_v72_)[default__.safeIndex((0) - (d_2282_v68_), len(d_2286_v72_))], d_2285_v71_])
                rhs301_ = (self).f19
                rhs302_ = ((self).f14).set(default__.safeIndex((d_2284_v70_).cardinality, len((self).f14)), (self).f19)
                rhs303_ = (d_2283_v69_)[default__.safeIndex(d_2282_v68_, len(d_2283_v69_))]
                rhs304_ = len((d_2286_v72_)[default__.safeIndex(d_2282_v68_, len(d_2286_v72_))])
                rhs305_ = ((self).f19 if (d_2282_v68_) > (len(_dafny.SeqWithoutIsStrInference([(self).f19]))) else (self).f19)
                lhs185_ = globalState
                lhs186_ = globalState
                lhs187_ = globalState
                lhs185_.f4 = rhs301_
                d_2283_v69_ = rhs302_
                lhs186_.f4 = rhs303_
                d_2282_v68_ = rhs304_
                lhs187_.f4 = rhs305_
                d_2288_v74_: _dafny.Set
                d_2288_v74_ = _dafny.Set({d_2282_v68_})
                d_2289_v75_: _dafny.Map
                d_2289_v75_ = _dafny.Map({d_2288_v74_: d_2282_v68_})
                d_2290_v76_: _dafny.Seq
                d_2290_v76_ = _dafny.SeqWithoutIsStrInference([d_2284_v70_, d_2284_v70_])
                d_2289_v75_ = (d_2289_v75_).set(_dafny.Set({(0) - (len(d_2290_v76_)), d_2282_v68_}), d_2282_v68_)
                d_2291_v77_: C3
                nw366_ = C3()
                nw366_.ctor__((self).f19, self.f12, (self).f19, (self).f13, (d_2283_v69_) + (d_2283_v69_))
                d_2291_v77_ = nw366_
                d_2292_v78_: C0
                nw367_ = C0()
                nw367_.ctor__()
                d_2292_v78_ = nw367_
            elif source36_.is_DC2:
                d_2293___mcc_h8_ = source36_.cf1
                d_2294___mcc_h9_ = source36_.cf2
                d_2295_cf2_ = d_2294___mcc_h9_
                d_2296_cf1_ = d_2293___mcc_h8_
                d_2297_v79_: _dafny.Array
                def lambda120_(d_2298_v68_):
                    def lambda121_(d_2299_i4_):
                        return (_dafny.Map({d_2298_v68_: _dafny.SeqWithoutIsStrInference([709])})) | (_dafny.Map({d_2298_v68_: _dafny.SeqWithoutIsStrInference([144])}))

                    return lambda121_

                init64_ = lambda120_(d_2282_v68_)
                nw368_ = _dafny.Array(None, 10)
                for i0_64_ in range(nw368_.length(0)):
                    nw368_[i0_64_] = init64_(i0_64_)
                d_2297_v79_ = nw368_
                d_2300_v80_: _dafny.Map
                d_2300_v80_ = _dafny.Map({(self).f19: d_2296_cf1_})
                d_2301_v81_: _dafny.Map
                d_2301_v81_ = _dafny.Map({((d_2300_v80_)[(self).f19] if ((self).f19) in (d_2300_v80_) else d_2282_v68_): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2302_i6_ in range(default__.abs(-747))])) for d_2303_i5_ in range(default__.abs(481))])})
                index354_ = default__.safeIndex(263, (d_2297_v79_).length(0))
                (d_2297_v79_)[index354_] = d_2301_v81_
                index355_ = default__.safeIndex(263, (d_2297_v79_).length(0))
                (d_2297_v79_)[index355_] = (d_2301_v81_) | (d_2301_v81_)
                d_2295_cf2_ = d_2295_cf2_
                d_2304_v82_: _dafny.Seq
                d_2304_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veo"))
                d_2305_v83_: _dafny.Seq
                d_2305_v83_ = _dafny.SeqWithoutIsStrInference([d_2296_cf1_, d_2282_v68_])
                d_2306_v84_: _dafny.Set
                d_2306_v84_ = _dafny.Set({d_2295_cf2_})
                d_2307_v85_: _dafny.Set
                d_2307_v85_ = _dafny.Set({(d_2305_v83_)[default__.safeIndex(len(d_2306_v84_), len(d_2305_v83_))]})
                d_2308_v86_: _dafny.Map
                d_2308_v86_ = _dafny.Map({d_2304_v82_: (0) - (len(d_2307_v85_))})
                d_2309_v87_: D14
                d_2309_v87_ = D14_DC41(d_2304_v82_, d_2296_cf1_)
                d_2308_v86_ = (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxkuwhg")): d_2282_v68_})).set((d_2309_v87_).cf64, (d_2282_v68_ if d_2295_cf2_ else len(d_2304_v82_)))
                d_2310_v88_: C1
                nw369_ = C1()
                nw369_.ctor__((d_2282_v68_) <= (len(d_2304_v82_)), d_2295_cf2_, (self).f13, (self).f14, self.f12)
                d_2310_v88_ = nw369_
            elif True:
                d_2311___mcc_h10_ = source36_.cf0
                d_2312_cf0_ = d_2311___mcc_h10_
                d_2313_v89_: _dafny.Map
                d_2313_v89_ = _dafny.Map({(self).f19: d_2282_v68_})
                d_2314_v90_: _dafny.Map
                d_2314_v90_ = _dafny.Map({_dafny.Map({len(d_2313_v89_): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onysgsey"))}): (self).f19})
                d_2315_v91_: _dafny.Seq
                d_2315_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghunfsbs"))
                (globalState).f4 = ((d_2314_v90_)[_dafny.Map({d_2282_v68_: d_2315_v91_})] if (_dafny.Map({d_2282_v68_: d_2315_v91_})) in (d_2314_v90_) else (default__.fm2((self).f19, globalState) if (self).f19 else default__.fm2((self).f19, globalState)))
                d_2316_v92_: _dafny.Array
                def lambda122_(d_2317_cf0_):
                    def lambda123_(d_2318_i7_):
                        return d_2317_cf0_

                    return lambda123_

                init65_ = lambda122_(d_2312_cf0_)
                nw370_ = _dafny.Array(None, 26)
                for i0_65_ in range(nw370_.length(0)):
                    nw370_[i0_65_] = init65_(i0_65_)
                d_2316_v92_ = nw370_
                d_2319_v93_: _dafny.Set
                d_2319_v93_ = _dafny.Set({d_2316_v92_})
                d_2320_v94_: _dafny.Map
                d_2320_v94_ = _dafny.Map({(self).f14: _dafny.Set({d_2316_v92_})})
                d_2321_v95_: _dafny.Seq
                d_2321_v95_ = _dafny.SeqWithoutIsStrInference([620, d_2282_v68_, d_2282_v68_, d_2282_v68_, d_2282_v68_])
                d_2322_v96_: T0
                nw371_ = C3()
                nw371_.ctor__(d_2312_cf0_, self.f12, (self).f19, (self).f13, _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19]))
                d_2322_v96_ = nw371_
                d_2323_v97_: _dafny.Map
                d_2323_v97_ = _dafny.Map({d_2312_cf0_: d_2322_v96_})
                d_2324_v98_: _dafny.Array
                nw372_ = _dafny.Array(None, 16)
                nw372_[int(0)] = (d_2319_v93_).ispropersubset(((d_2320_v94_)[(self).f14] if ((self).f14) in (d_2320_v94_) else _dafny.Set({d_2316_v92_})))
                nw372_[int(1)] = ((self).f19 if (self).f19 else d_2312_cf0_)
                nw372_[int(2)] = (self).f19
                nw372_[int(3)] = (d_2321_v95_) <= (d_2321_v95_)
                nw372_[int(4)] = (self).f19
                nw372_[int(5)] = d_2312_cf0_
                nw372_[int(6)] = d_2312_cf0_
                nw372_[int(7)] = (d_2282_v68_) <= (d_2282_v68_)
                nw372_[int(8)] = ((self).f19) in (d_2323_v97_)
                nw372_[int(9)] = d_2312_cf0_
                nw372_[int(10)] = True
                nw372_[int(11)] = (self).f19
                nw372_[int(12)] = d_2312_cf0_
                nw372_[int(13)] = (self).f19
                nw372_[int(14)] = (self).f19
                nw372_[int(15)] = (self).f19
                d_2324_v98_ = nw372_
                index356_ = default__.safeIndex(67, (d_2324_v98_).length(0))
                (d_2324_v98_)[index356_] = ((self).f14)[default__.safeIndex(d_2282_v68_, len((self).f14))]
                index357_ = default__.safeIndex(67, (d_2324_v98_).length(0))
                (d_2324_v98_)[index357_] = (self).f19
                d_2325_v99_: _dafny.Set
                d_2325_v99_ = _dafny.Set({d_2280_v66_, d_2280_v66_, d_2280_v66_})
                d_2326_v100_: D7
                d_2326_v100_ = D7_DC22(not((d_2324_v98_)[default__.safeIndex(67, (d_2324_v98_).length(0))]))
                d_2327_v101_: _dafny.Map
                d_2327_v101_ = _dafny.Map({402: d_2282_v68_})
                d_2325_v99_ = (default__.fm64(d_2312_cf0_, d_2326_v100_, ((d_2327_v101_)[d_2282_v68_] if (d_2282_v68_) in (d_2327_v101_) else d_2282_v68_), d_2282_v68_, globalState)) | ((d_2325_v99_) | (d_2325_v99_))
                d_2328_v102_: _dafny.Array
                nw373_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2328_v102_ = nw373_
                d_2329_v103_: _dafny.Map
                d_2329_v103_ = _dafny.Map({d_2282_v68_: (d_2326_v100_).cf34})
                d_2330_v104_: D11
                d_2330_v104_ = D11_DC31()
                d_2331_v105_: C1
                nw374_ = C1()
                nw374_.ctor__((self).f19, (d_2324_v98_)[default__.safeIndex(67, (d_2324_v98_).length(0))], (self).f13, (self).f14, d_2322_v96_.f12)
                d_2331_v105_ = nw374_
                d_2332_v106_: _dafny.Map
                d_2332_v106_ = _dafny.Map({d_2331_v105_: (d_2331_v105_).f17})
                d_2333_v107_: C2
                nw375_ = C2()
                nw375_.ctor__(_dafny.SeqWithoutIsStrInference([d_2282_v68_, d_2282_v68_]), ((d_2332_v106_)[d_2331_v105_] if (d_2331_v105_) in (d_2332_v106_) else d_2331_v105_.f18), (self).f13, ((self).f14).set(default__.safeIndex(len(d_2313_v89_), len((self).f14)), (d_2331_v105_).f17), d_2322_v96_.f12)
                d_2333_v107_ = nw375_
                d_2334_v108_: _dafny.Map
                d_2334_v108_ = _dafny.Map({d_2330_v104_: d_2333_v107_})
                d_2335_v109_: _dafny.MultiSet
                d_2335_v109_ = _dafny.MultiSet([d_2334_v108_, (d_2334_v108_).set(D11_DC31(), d_2333_v107_), d_2334_v108_, d_2334_v108_, d_2334_v108_])
                d_2336_v110_: _dafny.Array
                nw376_ = _dafny.Array(None, 21)
                nw376_[int(0)] = d_2282_v68_
                nw376_[int(1)] = d_2282_v68_
                nw376_[int(2)] = 247
                nw376_[int(3)] = len(d_2329_v103_)
                nw376_[int(4)] = d_2282_v68_
                nw376_[int(5)] = d_2282_v68_
                nw376_[int(6)] = d_2282_v68_
                nw376_[int(7)] = d_2282_v68_
                nw376_[int(8)] = d_2282_v68_
                nw376_[int(9)] = d_2282_v68_
                nw376_[int(10)] = d_2282_v68_
                nw376_[int(11)] = (0) - (d_2282_v68_)
                nw376_[int(12)] = (d_2335_v109_).cardinality
                nw376_[int(13)] = len((self).f14)
                nw376_[int(14)] = d_2282_v68_
                nw376_[int(15)] = d_2282_v68_
                nw376_[int(16)] = d_2282_v68_
                nw376_[int(17)] = d_2282_v68_
                nw376_[int(18)] = d_2282_v68_
                nw376_[int(19)] = d_2282_v68_
                nw376_[int(20)] = 238
                d_2336_v110_ = nw376_
                index358_ = default__.safeIndex(454, (d_2328_v102_).length(0))
                (d_2328_v102_)[index358_] = d_2336_v110_
                d_2337_v111_: str
                d_2337_v111_ = _dafny.CodePoint('d')
                index359_ = default__.safeIndex(67, (d_2324_v98_).length(0))
                index360_ = default__.safeIndex(454, (d_2328_v102_).length(0))
                rhs306_ = (d_2315_v91_).set(default__.safeIndex(default__.safeModuloInt(619, d_2282_v68_), len(d_2315_v91_)), d_2337_v111_)
                rhs307_ = ((self).f14)[default__.safeIndex((d_2333_v107_).fm38((0) - (d_2282_v68_), d_2282_v68_, d_2282_v68_, (self).f19, globalState), len((self).f14))]
                rhs308_ = ((d_2331_v105_.f18) or (d_2331_v105_.f18) if (d_2331_v105_).f17 else ((d_2331_v105_).f17) and (False))
                rhs309_ = (d_2282_v68_) * (682)
                rhs310_ = d_2336_v110_
                lhs188_ = d_2324_v98_
                lhs189_ = default__.safeIndex(67, (d_2324_v98_).length(0))
                lhs190_ = d_2328_v102_
                lhs191_ = default__.safeIndex(454, (d_2328_v102_).length(0))
                d_2315_v91_ = rhs306_
                lhs188_[lhs189_] = rhs307_
                r0 = rhs308_
                d_2282_v68_ = rhs309_
                lhs190_[lhs191_] = rhs310_
            d_2338_v112_: _dafny.Seq
            d_2338_v112_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ip"))
            d_2339_v113_: _dafny.MultiSet
            d_2339_v113_ = _dafny.MultiSet([d_2338_v112_])
            (globalState).f4 = (not((self).f19) if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2338_v112_, d_2338_v112_]))).ispropersubset(d_2339_v113_) else (self).f19)
            d_2340_v114_: C3
            nw377_ = C3()
            nw377_.ctor__(True, self.f12, (self).f19, (self).f13, (self).f14)
            d_2340_v114_ = nw377_
            d_2341_v115_: _dafny.Array
            def lambda124_(d_2342_v68_):
                def lambda125_(d_2343_i8_):
                    return default__.safeDivisionInt(d_2343_i8_, d_2342_v68_)

                return lambda125_

            init66_ = lambda124_(d_2282_v68_)
            nw378_ = _dafny.Array(None, 21)
            for i0_66_ in range(nw378_.length(0)):
                nw378_[i0_66_] = init66_(i0_66_)
            d_2341_v115_ = nw378_
            d_2344_v116_: _dafny.MultiSet
            d_2344_v116_ = _dafny.MultiSet([d_2282_v68_, d_2282_v68_])
            d_2345_v117_: _dafny.Seq
            d_2345_v117_ = _dafny.SeqWithoutIsStrInference([(d_2344_v116_).cardinality])
            index361_ = default__.safeIndex(25, (d_2341_v115_).length(0))
            (d_2341_v115_)[index361_] = ((d_2344_v116_)[d_2282_v68_] if (d_2282_v68_) in (d_2344_v116_) else len(d_2345_v117_))
            d_2346_v118_: _dafny.Array
            nw379_ = _dafny.Array(False, 13)
            d_2346_v118_ = nw379_
            d_2347_v119_: _dafny.Set
            d_2347_v119_ = _dafny.Set({(d_2340_v114_).f20, False, (self).f19, (d_2340_v114_).f20, (d_2340_v114_).f20})
            index362_ = default__.safeIndex(743, (d_2346_v118_).length(0))
            (d_2346_v118_)[index362_] = (d_2347_v119_).issubset(d_2347_v119_)
            d_2348_v120_: _dafny.Map
            d_2348_v120_ = _dafny.Map({len(d_2347_v119_): d_2347_v119_})
            index363_ = default__.safeIndex(754, (d_2341_v115_).length(0))
            (d_2341_v115_)[index363_] = default__.safeDivisionInt(d_2282_v68_, len(d_2348_v120_))
            d_2349_v121_: _dafny.Map
            d_2349_v121_ = _dafny.Map({(0) - (d_2282_v68_): (self).f19})
            d_2350_v122_: _dafny.Array
            def lambda126_(d_2351_v112_):
                def lambda127_(d_2352_i9_):
                    return d_2351_v112_

                return lambda127_

            init67_ = lambda126_(d_2338_v112_)
            nw380_ = _dafny.Array(None, 2)
            for i0_67_ in range(nw380_.length(0)):
                nw380_[i0_67_] = init67_(i0_67_)
            d_2350_v122_ = nw380_
            d_2353_v123_: _dafny.Map
            d_2353_v123_ = _dafny.Map({d_2338_v112_: default__.fm1((d_2340_v114_).f20, (d_2340_v114_).f20, globalState)})
            d_2354_v124_: D2
            d_2354_v124_ = D2_DC6(d_2346_v118_, d_2346_v118_)
            d_2355_v125_: _dafny.Set
            d_2355_v125_ = _dafny.Set({(d_2354_v124_).cf9, d_2346_v118_})
            d_2356_v126_: D12
            d_2356_v126_ = D12_DC36(d_2350_v122_, len(d_2345_v117_), d_2353_v123_, d_2355_v125_)
            index364_ = default__.safeIndex(25, (d_2341_v115_).length(0))
            index365_ = default__.safeIndex(743, (d_2346_v118_).length(0))
            index366_ = default__.safeIndex(754, (d_2341_v115_).length(0))
            rhs311_ = d_2282_v68_
            rhs312_ = (d_2282_v68_ if True else len(d_2349_v121_))
            rhs313_ = default__.fm2(not(True), globalState)
            rhs314_ = (d_2356_v126_).cf56
            lhs192_ = d_2341_v115_
            lhs193_ = default__.safeIndex(25, (d_2341_v115_).length(0))
            lhs194_ = d_2346_v118_
            lhs195_ = default__.safeIndex(743, (d_2346_v118_).length(0))
            lhs196_ = d_2341_v115_
            lhs197_ = default__.safeIndex(754, (d_2341_v115_).length(0))
            d_2282_v68_ = rhs311_
            lhs192_[lhs193_] = rhs312_
            lhs194_[lhs195_] = rhs313_
            lhs196_[lhs197_] = rhs314_
            d_2282_v68_ = (d_2282_v68_ if (self).f19 else (d_2341_v115_)[default__.safeIndex(25, (d_2341_v115_).length(0))])
        r0 = not (not((self).f19)) or ((self).f19)
        d_2357_v127_: int
        d_2357_v127_ = 372
        d_2357_v127_ = d_2357_v127_
        d_2358_v128_: _dafny.Seq
        d_2358_v128_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwb"))
        d_2359_v129_: _dafny.Map
        d_2359_v129_ = _dafny.Map({default__.fm65(-587, (self).f19, globalState): d_2358_v128_})
        d_2360_v130_: str
        d_2360_v130_ = _dafny.CodePoint('v')
        d_2358_v128_ = ((d_2359_v129_)[(_dafny.Map({d_2360_v130_: d_2360_v130_})).set(d_2360_v130_, d_2360_v130_)] if ((_dafny.Map({d_2360_v130_: d_2360_v130_})).set(d_2360_v130_, d_2360_v130_)) in (d_2359_v129_) else (d_2358_v128_) + (d_2358_v128_))
        d_2357_v127_ = d_2357_v127_
        d_2361_v131_: _dafny.Seq
        d_2361_v131_ = _dafny.SeqWithoutIsStrInference([d_2357_v127_])
        d_2362_v132_: T1
        nw381_ = C1()
        nw381_.ctor__((self).f19, (self).f19, (self).f13, _dafny.SeqWithoutIsStrInference([(self).f19]), self.f12)
        d_2362_v132_ = nw381_
        d_2363_v133_: D3
        d_2363_v133_ = D3_DC10(False, len(d_2361_v131_), d_2362_v132_)
        d_2364_v134_: _dafny.MultiSet
        d_2364_v134_ = _dafny.MultiSet([d_2363_v133_])
        r0 = not ((d_2363_v133_) not in (d_2364_v134_)) or ((self).f19)
        return r0

    @property
    def f19(self):
        return self._f19

class C6(T3, T4, T2, T1, T0):
    def  __init__(self):
        self._f12: _dafny.Array = _dafny.Array(None, 0)
        self._f16: _dafny.Seq = _dafny.Seq({})
        self._f15: bool = False
        self._f13: _dafny.Map = _dafny.Map({})
        self._f14: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f16(self):
        return self._f16
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    def ctor__(self, f16, f15, f13, f14, f12):
        (self)._f16 = f16
        (self)._f15 = f15
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f12 = f12

    def fm9(self, globalState):
        if (self).f15:
            return _dafny.MultiSet([(self).f15, (self).f15])
        elif True:
            return _dafny.MultiSet([(self).f15, (self).f15, (self).f15, (self).f15])

    def fm7(self, p0, globalState):
        return default__.safeDivisionInt((0) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qi"))))), default__.safeModuloInt(len(_dafny.Map({not((self).f15): _dafny.MultiSet([(self).f15, (self).f15])})), -561))

    def fm8(self, p0, p1, globalState):
        return ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2365_i0_ in range(default__.abs(354))]))) + (len((self).f14))) + (((D0_DC2(-182, False)).cf1) * (-297))

    def fm6(self, p0, globalState):
        source37_ = D0_DC2((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2366_i0_ in range(default__.abs(409))])))), (self).f15)
        if source37_.is_DC1:
            return (self).f16
        elif source37_.is_DC2:
            d_2367___mcc_h0_ = source37_.cf1
            d_2368___mcc_h1_ = source37_.cf2
            d_2369_cf2_ = d_2368___mcc_h1_
            d_2370_cf1_ = d_2367___mcc_h0_
            return (self).f16
        elif True:
            d_2371___mcc_h2_ = source37_.cf0
            d_2372_cf0_ = d_2371___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet((self).f14)).cardinality)])) + (((self).f16).set(default__.safeIndex(len(_dafny.Map({d_2372_cf0_: len(_dafny.Map({(self).f15: (self).f15}))})), len((self).f16)), 17))

    def fm5(self, p0, globalState):
        return 932

    def fm10(self, p0, p1, p2, p3, globalState):
        return (_dafny.Set({(0) - (len(_dafny.Map({D0_DC2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), (self).f15): (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2373_i0_ in range(default__.abs(287))])))})))})) | ((_dafny.Set({832})) | (_dafny.Set({254, -852})))

    def fm11(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gitfodix"))

    def fm12(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydouas")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlaj"))])

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_2374_v0_: _dafny.Set
        d_2374_v0_ = _dafny.Set({(self).f15})
        d_2375_v1_: _dafny.Map
        d_2375_v1_ = _dafny.Map({d_2374_v0_: ((_dafny.MultiSet([(self).f15])).cardinality) >= (p0)})
        d_2375_v1_ = (d_2375_v1_).set(d_2374_v0_, (self).f15)
        d_2376_v2_: D0
        d_2376_v2_ = D0_DC2(p1, (self).f15)
        d_2377_v3_: _dafny.Map
        d_2377_v3_ = _dafny.Map({p1: (self).f15})
        d_2378_v4_: _dafny.Map
        d_2378_v4_ = _dafny.Map({(421) <= (len(_dafny.Set({(self).f15, (d_2376_v2_).cf2}))): ((d_2377_v3_)[(0) - ((D0_DC2(100, (self).f15)).cf1)] if ((0) - ((D0_DC2(100, (self).f15)).cf1)) in (d_2377_v3_) else not((self).f15))})
        d_2378_v4_ = (d_2378_v4_).set(((self).f15) or ((self).f15), False)
        if (default__.fm13(p1, globalState)) == (d_2374_v0_):
            d_2379_v5_: _dafny.Map
            d_2379_v5_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2380_i0_ in range(default__.abs(349))]): p2})
            d_2379_v5_ = ((d_2379_v5_ if False else d_2379_v5_)) | (_dafny.Map({p2: p2}))
            d_2381_v6_: C0
            nw382_ = C0()
            nw382_.ctor__()
            d_2381_v6_ = nw382_
            d_2382_v7_: _dafny.Array
            nw383_ = _dafny.Array(False, 18)
            d_2382_v7_ = nw383_
            d_2383_v8_: _dafny.MultiSet
            d_2383_v8_ = _dafny.MultiSet([p2])
            d_2384_v9_: _dafny.Map
            d_2384_v9_ = _dafny.Map({not((self).f15): d_2383_v8_})
            index367_ = default__.safeIndex(890, (d_2382_v7_).length(0))
            (d_2382_v7_)[index367_] = not((p2) not in (((d_2384_v9_)[False] if (False) in (d_2384_v9_) else d_2383_v8_)))
            index368_ = default__.safeIndex(890, (d_2382_v7_).length(0))
            (d_2382_v7_)[index368_] = (self).f15
            r0 = (0) - (default__.safeModuloInt(476, (p1) * (p1)))
            d_2385_v10_: D2
            d_2385_v10_ = D2_DC6(d_2382_v7_, d_2382_v7_)
            d_2386_v11_: _dafny.Map
            d_2386_v11_ = _dafny.Map({d_2385_v10_: len((d_2378_v4_) | (d_2378_v4_))})
            pat_let_tv39_ = d_2382_v7_
            def iife157_(_pat_let45_0):
                def iife158_(d_2387_dt__update__tmp_h0_):
                    def iife159_(_pat_let46_0):
                        def iife160_(d_2388_dt__update_hcf8_h0_):
                            return D2_DC6(d_2388_dt__update_hcf8_h0_, (d_2387_dt__update__tmp_h0_).cf9)
                        return iife160_(_pat_let46_0)
                    return iife159_(pat_let_tv39_)
                return iife158_(_pat_let45_0)
            d_2386_v11_ = (d_2386_v11_).set(iife157_(d_2385_v10_), default__.safeModuloInt(p0, p0))
        elif True:
            (globalState).f4 = (self).f15
            d_2389_v12_: C0
            nw384_ = C0()
            nw384_.ctor__()
            d_2389_v12_ = nw384_
            r0 = p1
            d_2390_v13_: _dafny.Array
            nw385_ = _dafny.Array(int(0), 3)
            d_2390_v13_ = nw385_
            index369_ = default__.safeIndex(933, (d_2390_v13_).length(0))
            (d_2390_v13_)[index369_] = p0
            index370_ = default__.safeIndex(933, (d_2390_v13_).length(0))
            (d_2390_v13_)[index370_] = (0) - (((p0) - (496)) + (p0))
            index371_ = default__.safeIndex(933, (d_2390_v13_).length(0))
            (d_2390_v13_)[index371_] = 660
        hi12_ = len(d_2374_v0_)
        for d_2391_i1_ in range(p1, hi12_):
            d_2378_v4_ = (d_2378_v4_).set((self).f15, not((self).f15))
            d_2392_v14_: C0
            nw386_ = C0()
            nw386_.ctor__()
            d_2392_v14_ = nw386_
            d_2393_v15_: _dafny.Map
            d_2393_v15_ = _dafny.Map({(d_2391_i1_) > (d_2391_i1_): 659})
            r0 = len(d_2393_v15_)
            d_2394_v16_: _dafny.Array
            nw387_ = _dafny.Array(int(0), 28)
            d_2394_v16_ = nw387_
            index372_ = default__.safeIndex(543, (d_2394_v16_).length(0))
            (d_2394_v16_)[index372_] = default__.safeModuloInt(p0, d_2391_i1_)
            index373_ = default__.safeIndex(543, (d_2394_v16_).length(0))
            (d_2394_v16_)[index373_] = (d_2391_i1_) * (p0)
        if (self).f15:
            (globalState).f4 = (self).f15
            r0 = default__.safeDivisionInt((p1) + (p1), 738)
            d_2395_v17_: _dafny.Map
            d_2395_v17_ = _dafny.Map({p1: p2})
            d_2395_v17_ = (d_2395_v17_).set(p1, (p2) + (p2))
            d_2396_v18_: str
            d_2396_v18_ = _dafny.CodePoint('q')
            d_2397_v19_: _dafny.Seq
            d_2397_v19_ = _dafny.SeqWithoutIsStrInference([(p2) + ((p2).set(default__.safeIndex(p1, len(p2)), d_2396_v18_)), p2, (p2) + (_dafny.SeqWithoutIsStrInference([d_2396_v18_ for d_2398_i2_ in range(default__.abs(371))])), (self).fm11(globalState)])
            d_2399_v20_: _dafny.Map
            d_2399_v20_ = _dafny.Map({len(d_2377_v3_): d_2397_v19_})
            d_2397_v19_ = ((d_2399_v20_)[(0) - ((p0 if not((self).f15) else p1))] if ((0) - ((p0 if not((self).f15) else p1))) in (d_2399_v20_) else d_2397_v19_)
            r0 = p0
        elif True:
            source38_ = default__.fm16(default__.fm2(False, globalState), globalState)
            if source38_.is_DC5:
                d_2400___mcc_h0_ = source38_.cf5
                d_2401___mcc_h1_ = source38_.cf6
                d_2402___mcc_h2_ = source38_.cf7
                d_2403_cf7_ = d_2402___mcc_h2_
                d_2404_cf6_ = d_2401___mcc_h1_
                d_2405_cf5_ = d_2400___mcc_h0_
                d_2406_v21_: _dafny.Array
                nw388_ = _dafny.Array(_dafny.Map({}), 3)
                d_2406_v21_ = nw388_
                index374_ = default__.safeIndex(513, (d_2406_v21_).length(0))
                (d_2406_v21_)[index374_] = d_2378_v4_
                d_2407_v22_: _dafny.Set
                d_2407_v22_ = _dafny.Set({(self).f14, (self).f14})
                d_2408_v23_: _dafny.Map
                d_2408_v23_ = _dafny.Map({len(d_2407_v22_): p1})
                index375_ = default__.safeIndex(513, (d_2406_v21_).length(0))
                rhs315_ = 343
                rhs316_ = ((d_2408_v23_)[d_2405_cf5_] if (d_2405_cf5_) in (d_2408_v23_) else p0)
                rhs317_ = d_2378_v4_
                rhs318_ = p1
                rhs319_ = (self).f15
                lhs198_ = d_2406_v21_
                lhs199_ = default__.safeIndex(513, (d_2406_v21_).length(0))
                lhs200_ = globalState
                r0 = rhs315_
                d_2405_cf5_ = rhs316_
                lhs198_[lhs199_] = rhs317_
                d_2405_cf5_ = rhs318_
                lhs200_.f4 = rhs319_
                r0 = (0) - (p1)
                d_2409_v24_: _dafny.Array
                nw389_ = _dafny.Array(int(0), 8)
                d_2409_v24_ = nw389_
                d_2409_v24_ = d_2409_v24_
                d_2405_cf5_ = d_2405_cf5_
            elif source38_.is_DC6:
                d_2410___mcc_h3_ = source38_.cf8
                d_2411___mcc_h4_ = source38_.cf9
                d_2412_cf9_ = d_2411___mcc_h4_
                d_2413_cf8_ = d_2410___mcc_h3_
                d_2414_v25_: _dafny.Seq
                d_2414_v25_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                d_2414_v25_ = d_2414_v25_
                (globalState).f4 = not((p1) > (default__.safeModuloInt((0) - (p1), p1)))
                d_2415_v26_: _dafny.Seq
                d_2415_v26_ = _dafny.SeqWithoutIsStrInference([d_2414_v25_, (self).f14, d_2414_v25_])
                d_2416_v27_: D2
                d_2416_v27_ = D2_DC4(d_2415_v26_)
                d_2417_v28_: D2
                d_2417_v28_ = D2_DC7(d_2416_v27_)
                d_2418_v29_: _dafny.Map
                d_2418_v29_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15, (self).f15, (self).f15]): d_2417_v28_})
                d_2418_v29_ = (d_2418_v29_).set((self).f14, d_2417_v28_)
                d_2419_v30_: _dafny.Set
                d_2419_v30_ = _dafny.Set({p0})
                d_2420_v31_: _dafny.Set
                d_2420_v31_ = _dafny.Set({len(d_2419_v30_), p0})
                d_2421_v32_: _dafny.Map
                d_2421_v32_ = _dafny.Map({default__.safeDivisionInt(89, len(d_2420_v31_)): d_2413_cf8_})
                d_2412_cf9_ = ((d_2421_v32_)[(0) - ((p1) * (290))] if ((0) - ((p1) * (290))) in (d_2421_v32_) else d_2412_cf9_)
            elif source38_.is_DC4:
                d_2422___mcc_h5_ = source38_.cf4
                d_2423_cf4_ = d_2422___mcc_h5_
                d_2424_v33_: _dafny.Array
                def lambda128_(d_2425_v4_):
                    def lambda129_(d_2426_i3_):
                        return (d_2426_i3_) * (len(d_2425_v4_))

                    return lambda129_

                init68_ = lambda128_(d_2378_v4_)
                nw390_ = _dafny.Array(None, 7)
                for i0_68_ in range(nw390_.length(0)):
                    nw390_[i0_68_] = init68_(i0_68_)
                d_2424_v33_ = nw390_
                d_2427_v34_: D0
                d_2427_v34_ = D0_DC1()
                d_2428_v35_: _dafny.Seq
                d_2428_v35_ = _dafny.SeqWithoutIsStrInference([p2])
                d_2429_v36_: D2
                d_2429_v36_ = D2_DC5(p0, (self).f15, not((self).f15))
                d_2430_v37_: D0
                d_2430_v37_ = D0_DC0((self).f15)
                d_2431_v38_: str
                d_2431_v38_ = _dafny.CodePoint('i')
                rhs320_ = d_2424_v33_
                rhs321_ = (d_2427_v34_ if (d_2429_v36_).cf7 else d_2427_v34_)
                rhs322_ = (self).fm12(d_2430_v37_, d_2431_v38_, globalState)
                rhs323_ = p0
                rhs324_ = len(_dafny.SeqWithoutIsStrInference([(self).f15, (self).f15]))
                d_2424_v33_ = rhs320_
                d_2427_v34_ = rhs321_
                d_2428_v35_ = rhs322_
                r0 = rhs323_
                r0 = rhs324_
                d_2378_v4_ = (d_2378_v4_).set((self).f15, ((self).f14)[default__.safeIndex((0) - (p0), len((self).f14))])
                (globalState).f4 = default__.fm2((p0) < (p1), globalState)
                index376_ = default__.safeIndex(447, (d_2424_v33_).length(0))
                (d_2424_v33_)[index376_] = len(d_2374_v0_)
                index377_ = default__.safeIndex(447, (d_2424_v33_).length(0))
                (d_2424_v33_)[index377_] = p0
            elif True:
                d_2432___mcc_h6_ = source38_.cf10
                d_2433_cf10_ = d_2432___mcc_h6_
                d_2434_v39_: _dafny.Map
                d_2434_v39_ = _dafny.Map({p0: p1})
                d_2435_v40_: _dafny.MultiSet
                d_2435_v40_ = _dafny.MultiSet([d_2434_v39_])
                (globalState).f4 = (d_2435_v40_) != ((d_2435_v40_) - (d_2435_v40_))
                d_2436_v41_: D2
                d_2436_v41_ = D2_DC5(p0, False, (self).f15)
                d_2437_v42_: T1
                nw391_ = C1()
                nw391_.ctor__(True, (d_2436_v41_).cf6, (self).f13, (self).f14, self.f12)
                d_2437_v42_ = nw391_
                d_2438_v43_: _dafny.MultiSet
                d_2438_v43_ = _dafny.MultiSet([d_2437_v42_])
                (globalState).f4 = ((d_2438_v43_) | (d_2438_v43_)) == (d_2438_v43_)
                (globalState).f4 = (not((self).f15)) not in (d_2374_v0_)
                d_2439_v44_: _dafny.Array
                def lambda130_(d_2440_i4_):
                    return D0_DC0((self).f15)

                init69_ = lambda130_
                nw392_ = _dafny.Array(None, 27)
                for i0_69_ in range(nw392_.length(0)):
                    nw392_[i0_69_] = init69_(i0_69_)
                d_2439_v44_ = nw392_
                index378_ = default__.safeIndex(351, (d_2439_v44_).length(0))
                (d_2439_v44_)[index378_] = default__.fm19(_dafny.Set({(self).f15, (self).f15}), (self).f15, _dafny.CodePoint('l'), p1, globalState)
                d_2441_v45_: D0
                d_2441_v45_ = D0_DC0((self).f15)
                index379_ = default__.safeIndex(351, (d_2439_v44_).length(0))
                (d_2439_v44_)[index379_] = d_2441_v45_
            d_2442_v46_: _dafny.Seq
            d_2442_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xl"))
            d_2442_v46_ = d_2442_v46_
            d_2443_v47_: C0
            nw393_ = C0()
            nw393_.ctor__()
            d_2443_v47_ = nw393_
            d_2444_v48_: _dafny.Array
            nw394_ = _dafny.Array(False, 23)
            d_2444_v48_ = nw394_
            d_2445_v49_: T4
            nw395_ = C4()
            nw395_.ctor__()
            d_2445_v49_ = nw395_
            d_2446_v50_: _dafny.Map
            d_2446_v50_ = _dafny.Map({d_2445_v49_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hku"))})
            index380_ = default__.safeIndex(205, (d_2444_v48_).length(0))
            (d_2444_v48_)[index380_] = not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2447_i5_ in range(default__.abs(51))])) < (((d_2446_v50_)[d_2445_v49_] if (d_2445_v49_) in (d_2446_v50_) else (self).fm11(globalState))))
            d_2448_v51_: C1
            nw396_ = C1()
            nw396_.ctor__(not(not (True) or ((self).f15)), (self).f15, (self).f13, _dafny.SeqWithoutIsStrInference([(self).f15, False]), self.f12)
            d_2448_v51_ = nw396_
            d_2449_v52_: _dafny.MultiSet
            d_2449_v52_ = _dafny.MultiSet([(self).f15])
            d_2450_v53_: _dafny.Map
            d_2450_v53_ = _dafny.Map({d_2449_v52_: (self).f15})
            d_2451_v54_: C0
            d_2451_v54_ = d_2443_v47_
            index381_ = default__.safeIndex(205, (d_2444_v48_).length(0))
            rhs325_ = (d_2448_v51_).f17
            rhs326_ = default__.safeDivisionInt(p0, default__.safeModuloInt(p0, len(d_2450_v53_)))
            rhs327_ = (d_2451_v54_)
            rhs328_ = False
            rhs329_ = d_2448_v51_
            lhs201_ = globalState
            lhs202_ = d_2444_v48_
            lhs203_ = default__.safeIndex(205, (d_2444_v48_).length(0))
            lhs201_.f4 = rhs325_
            r0 = rhs326_
            d_2443_v47_ = rhs327_
            lhs202_[lhs203_] = rhs328_
            d_2448_v51_ = rhs329_
            d_2452_v55_: C3
            nw397_ = C3()
            nw397_.ctor__((d_2448_v51_).f17, self.f12, (self).f15, (self).f13, ((self).f14 if (d_2448_v51_).f17 else (self).f14))
            d_2452_v55_ = nw397_
            d_2442_v46_ = p2
        d_2453_v56_: _dafny.Array
        nw398_ = _dafny.Array(None, 1)
        nw398_[int(0)] = p2
        d_2453_v56_ = nw398_
        index382_ = default__.safeIndex(761, (d_2453_v56_).length(0))
        (d_2453_v56_)[index382_] = p2
        index383_ = default__.safeIndex(761, (d_2453_v56_).length(0))
        (d_2453_v56_)[index383_] = p2
        r0 = default__.safeDivisionInt(p0, len((d_2453_v56_)[default__.safeIndex(761, (d_2453_v56_).length(0))]))
        return r0

    def m3(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_2454_v0_: _dafny.Seq
        d_2454_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxffsdmo"))
        d_2455_v1_: int
        d_2455_v1_ = 30
        hi13_ = (self).fm5((self).f15, globalState)
        for d_2456_i0_ in range(default__.safeModuloInt(len(d_2454_v0_), (self).fm7(d_2455_v1_, globalState)), hi13_):
            d_2457_v2_: _dafny.MultiSet
            d_2457_v2_ = _dafny.MultiSet([d_2455_v1_, d_2455_v1_])
            d_2458_v3_: _dafny.MultiSet
            d_2458_v3_ = _dafny.MultiSet([((d_2457_v2_)[d_2455_v1_] if (d_2455_v1_) in (d_2457_v2_) else d_2455_v1_)])
            d_2459_v4_: _dafny.Set
            d_2459_v4_ = _dafny.Set({228, (d_2457_v2_).cardinality, d_2456_i0_, d_2456_i0_, d_2456_i0_})
            def iife161_():
                coll67_ = _dafny.Set()
                compr_67_: int
                for compr_67_ in _dafny.IntegerRange(211, -38):
                    d_2460_v5_: int = compr_67_
                    if ((211) <= (d_2460_v5_)) and ((d_2460_v5_) < (-38)):
                        coll67_ = coll67_.union(_dafny.Set([(d_2460_v5_) + (d_2455_v1_)]))
                return _dafny.Set(coll67_)
            d_2459_v4_ = ((d_2459_v4_) | (iife161_()
            )).intersection(d_2459_v4_)
            d_2461_v6_: _dafny.Array
            nw399_ = _dafny.Array(False, 28)
            d_2461_v6_ = nw399_
            d_2462_v7_: D2
            d_2462_v7_ = D2_DC6(d_2461_v6_, d_2461_v6_)
            pat_let_tv40_ = d_2461_v6_
            def iife162_(_pat_let47_0):
                def iife163_(d_2463_dt__update__tmp_h0_):
                    def iife164_(_pat_let48_0):
                        def iife165_(d_2464_dt__update_hcf8_h0_):
                            return D2_DC6(d_2464_dt__update_hcf8_h0_, (d_2463_dt__update__tmp_h0_).cf9)
                        return iife165_(_pat_let48_0)
                    return iife164_(pat_let_tv40_)
                return iife163_(_pat_let47_0)
            source39_ = iife162_(d_2462_v7_)
            if source39_.is_DC5:
                d_2465___mcc_h0_ = source39_.cf5
                d_2466___mcc_h1_ = source39_.cf6
                d_2467___mcc_h2_ = source39_.cf7
                d_2468_cf7_ = d_2467___mcc_h2_
                d_2469_cf6_ = d_2466___mcc_h1_
                d_2470_cf5_ = d_2465___mcc_h0_
                d_2470_cf5_ = default__.safeDivisionInt(168, d_2470_cf5_)
                d_2469_cf6_ = d_2468_cf7_
                r1 = ((0) - (d_2455_v1_)) * ((self).fm5((self).f15, globalState))
                d_2454_v0_ = d_2454_v0_
            elif source39_.is_DC6:
                d_2471___mcc_h3_ = source39_.cf8
                d_2472___mcc_h4_ = source39_.cf9
                d_2473_cf9_ = d_2472___mcc_h4_
                d_2474_cf8_ = d_2471___mcc_h3_
                r1 = d_2456_i0_
                d_2475_v8_: _dafny.Map
                d_2475_v8_ = _dafny.Map({(self).f15: 297})
                d_2476_v9_: _dafny.Map
                d_2476_v9_ = _dafny.Map({d_2455_v1_: d_2475_v8_})
                d_2477_v10_: _dafny.Map
                d_2477_v10_ = _dafny.Map({d_2456_i0_: (self).f15})
                d_2478_v11_: _dafny.Set
                d_2478_v11_ = _dafny.Set({(self).f15, False})
                r1 = default__.fm0((self).f15, d_2476_v9_, ((d_2477_v10_)[(0) - (len(d_2478_v11_))] if ((0) - (len(d_2478_v11_))) in (d_2477_v10_) else (self).f15), globalState)
                d_2454_v0_ = d_2454_v0_
                index384_ = default__.safeIndex(184, (d_2474_cf8_).length(0))
                (d_2474_cf8_)[index384_] = (self).f15
                index385_ = default__.safeIndex(184, (d_2474_cf8_).length(0))
                (d_2474_cf8_)[index385_] = (self).f15
            elif source39_.is_DC4:
                d_2479___mcc_h5_ = source39_.cf4
                d_2480_cf4_ = d_2479___mcc_h5_
                d_2481_v12_: _dafny.Array
                def lambda131_(d_2482_v1_):
                    def lambda132_(d_2483_i1_):
                        return default__.safeDivisionInt(d_2483_i1_, (0) - (d_2482_v1_))

                    return lambda132_

                init70_ = lambda131_(d_2455_v1_)
                nw400_ = _dafny.Array(None, 9)
                for i0_70_ in range(nw400_.length(0)):
                    nw400_[i0_70_] = init70_(i0_70_)
                d_2481_v12_ = nw400_
                d_2484_v13_: D12
                d_2484_v13_ = D12_DC35((self).f15)
                d_2485_v14_: _dafny.Array
                nw401_ = _dafny.Array(None, 11)
                nw401_[int(0)] = d_2461_v6_
                nw401_[int(1)] = d_2461_v6_
                nw401_[int(2)] = d_2461_v6_
                nw401_[int(3)] = d_2461_v6_
                nw401_[int(4)] = d_2461_v6_
                nw401_[int(5)] = d_2461_v6_
                nw401_[int(6)] = d_2461_v6_
                nw401_[int(7)] = d_2461_v6_
                nw401_[int(8)] = d_2461_v6_
                nw401_[int(9)] = d_2461_v6_
                nw401_[int(10)] = d_2461_v6_
                d_2485_v14_ = nw401_
                index386_ = default__.safeIndex(708, (d_2485_v14_).length(0))
                (d_2485_v14_)[index386_] = d_2461_v6_
                d_2486_v15_: _dafny.Array
                def lambda133_(d_2487_i2_):
                    return (_dafny.MultiSet([_dafny.CodePoint('j')])) | (_dafny.MultiSet([_dafny.CodePoint('h')]))

                init71_ = lambda133_
                nw402_ = _dafny.Array(None, 10)
                for i0_71_ in range(nw402_.length(0)):
                    nw402_[i0_71_] = init71_(i0_71_)
                d_2486_v15_ = nw402_
                d_2488_v16_: str
                d_2488_v16_ = _dafny.CodePoint('j')
                d_2489_v17_: _dafny.MultiSet
                d_2489_v17_ = _dafny.MultiSet([d_2488_v16_, d_2488_v16_])
                d_2490_v18_: _dafny.Seq
                d_2490_v18_ = _dafny.SeqWithoutIsStrInference([d_2489_v17_, d_2489_v17_])
                index387_ = default__.safeIndex(163, (d_2486_v15_).length(0))
                (d_2486_v15_)[index387_] = (d_2490_v18_)[default__.safeIndex(d_2455_v1_, len(d_2490_v18_))]
                index388_ = default__.safeIndex(708, (d_2485_v14_).length(0))
                index389_ = default__.safeIndex(163, (d_2486_v15_).length(0))
                rhs330_ = d_2481_v12_
                rhs331_ = d_2484_v13_
                rhs332_ = d_2461_v6_
                rhs333_ = not((self).f15)
                rhs334_ = _dafny.MultiSet([d_2488_v16_, d_2488_v16_, d_2488_v16_, d_2488_v16_])
                lhs204_ = d_2485_v14_
                lhs205_ = default__.safeIndex(708, (d_2485_v14_).length(0))
                lhs206_ = globalState
                lhs207_ = d_2486_v15_
                lhs208_ = default__.safeIndex(163, (d_2486_v15_).length(0))
                d_2481_v12_ = rhs330_
                d_2484_v13_ = rhs331_
                lhs204_[lhs205_] = rhs332_
                lhs206_.f4 = rhs333_
                lhs207_[lhs208_] = rhs334_
                r0 = (not ((self).f15) or ((self).f15)) or ((self).f15)
                d_2491_v19_: D0
                d_2491_v19_ = D0_DC2(d_2455_v1_, (self).f15)
                d_2492_v20_: _dafny.MultiSet
                d_2492_v20_ = _dafny.MultiSet([(self).f15])
                d_2493_v21_: _dafny.Map
                d_2493_v21_ = _dafny.Map({d_2492_v20_: d_2461_v6_})
                d_2494_v22_: _dafny.Seq
                d_2495_v23_: bool
                d_2496_v24_: int
                d_2497_v25_: bool
                out119_: _dafny.Seq
                out120_: bool
                out121_: int
                out122_: bool
                out119_, out120_, out121_, out122_ = default__.m0((self).f15, d_2491_v19_, (d_2456_i0_) == (len(d_2493_v21_)), globalState)
                d_2494_v22_ = out119_
                d_2495_v23_ = out120_
                d_2496_v24_ = out121_
                d_2497_v25_ = out122_
                d_2498_v26_: C1
                nw403_ = C1()
                nw403_.ctor__((d_2488_v16_) not in (d_2454_v0_), (d_2497_v25_ if (self).f15 else d_2495_v23_), ((self).f13) | ((self).f13), (self).f14, self.f12)
                d_2498_v26_ = nw403_
                d_2498_v26_ = d_2498_v26_
            elif True:
                d_2499___mcc_h6_ = source39_.cf10
                d_2500_cf10_ = d_2499___mcc_h6_
                d_2501_v27_: _dafny.Map
                d_2501_v27_ = _dafny.Map({(self).f15: True})
                d_2502_v28_: D8
                d_2502_v28_ = D8_DC26((self).f15)
                d_2503_v29_: _dafny.Map
                d_2503_v29_ = _dafny.Map({d_2501_v27_: (d_2502_v28_).cf41})
                d_2503_v29_ = (d_2503_v29_).set(d_2501_v27_, ((self).f15 if (self).f15 else (self).f15))
                d_2504_v30_: _dafny.MultiSet
                d_2504_v30_ = _dafny.MultiSet([True, (self).f15, (self).f15, (self).f15, not((self).f15)])
                d_2504_v30_ = d_2504_v30_
                d_2505_v31_: _dafny.Set
                d_2505_v31_ = _dafny.Set({(self).f15})
                d_2506_v32_: _dafny.Set
                d_2506_v32_ = d_2505_v31_
                r3 = (((d_2506_v32_) if (self).f15 else _dafny.Set({(self).f15, False}))).issubset(d_2505_v31_)
                index390_ = default__.safeIndex(798, (d_2461_v6_).length(0))
                (d_2461_v6_)[index390_] = (self).f15
                index391_ = default__.safeIndex(798, (d_2461_v6_).length(0))
                (d_2461_v6_)[index391_] = (self).f15
            d_2507_v33_: C3
            nw404_ = C3()
            nw404_.ctor__(default__.fm2((self).f15, globalState), self.f12, (self).f15, (self).f13, ((self).f14) + ((self).f14))
            d_2507_v33_ = nw404_
            r0 = (d_2507_v33_).f20
        (globalState).f4 = (default__.safeModuloInt((self).fm5((self).f15, globalState), 90)) != (default__.safeDivisionInt(d_2455_v1_, d_2455_v1_))
        d_2508_v34_: _dafny.Array
        def lambda134_(d_2509_v1_):
            def lambda135_(d_2510_i3_):
                return _dafny.Set({d_2509_v1_, d_2509_v1_, -781, d_2509_v1_})

            return lambda135_

        init72_ = lambda134_(d_2455_v1_)
        nw405_ = _dafny.Array(None, 20)
        for i0_72_ in range(nw405_.length(0)):
            nw405_[i0_72_] = init72_(i0_72_)
        d_2508_v34_ = nw405_
        d_2511_v35_: _dafny.Set
        d_2511_v35_ = _dafny.Set({d_2455_v1_})
        index392_ = default__.safeIndex(562, (d_2508_v34_).length(0))
        (d_2508_v34_)[index392_] = d_2511_v35_
        index393_ = default__.safeIndex(562, (d_2508_v34_).length(0))
        (d_2508_v34_)[index393_] = d_2511_v35_
        d_2512_v36_: _dafny.Set
        d_2512_v36_ = _dafny.Set({d_2454_v0_})
        d_2513_v37_: _dafny.MultiSet
        d_2513_v37_ = _dafny.MultiSet([len(d_2512_v36_)])
        d_2514_v38_: str
        d_2514_v38_ = _dafny.CodePoint('a')
        d_2515_v39_: _dafny.Map
        d_2515_v39_ = _dafny.Map({d_2514_v38_: (self).f15})
        d_2516_i4_: int
        d_2516_i4_ = 0
        with _dafny.label("14"):
            while ((0) - (d_2455_v1_)) <= (default__.safeDivisionInt((0) - (((d_2513_v37_)[d_2455_v1_] if (d_2455_v1_) in (d_2513_v37_) else len(d_2515_v39_))), len(d_2454_v0_))):
                with _dafny.c_label("14"):
                    if (d_2516_i4_) >= (100):
                        raise _dafny.Break("14")
                    d_2516_i4_ = (d_2516_i4_) + (1)
                    r0 = (d_2455_v1_) not in ((self).f16)
                    d_2517_v40_: D2
                    d_2517_v40_ = D2_DC5(d_2455_v1_, (self).f15, (self).f15)
                    rhs335_ = d_2517_v40_
                    rhs336_ = d_2514_v38_
                    rhs337_ = default__.fm2((self).f15, globalState)
                    lhs209_ = globalState
                    d_2517_v40_ = rhs335_
                    d_2514_v38_ = rhs336_
                    lhs209_.f4 = rhs337_
                    d_2518_v41_: _dafny.Map
                    d_2518_v41_ = _dafny.Map({814: (self).f15})
                    d_2519_v42_: _dafny.Map
                    d_2519_v42_ = _dafny.Map({((self).f15 if ((d_2518_v41_)[len(_dafny.SeqWithoutIsStrInference([d_2514_v38_]))] if (len(_dafny.SeqWithoutIsStrInference([d_2514_v38_]))) in (d_2518_v41_) else (self).f15) else (self).f15): (self).f15})
                    d_2520_v43_: _dafny.Set
                    d_2520_v43_ = _dafny.Set({(self).f15})
                    d_2521_v44_: _dafny.Set
                    d_2521_v44_ = d_2520_v43_
                    rhs338_ = (self).f15
                    rhs339_ = (self).f15
                    rhs340_ = (d_2455_v1_) - (len((d_2521_v44_)))
                    rhs341_ = ((d_2519_v42_) | (d_2519_v42_)).set((d_2455_v1_) != (548), (self).f15)
                    lhs210_ = globalState
                    lhs211_ = globalState
                    lhs210_.f4 = rhs338_
                    lhs211_.f4 = rhs339_
                    d_2455_v1_ = rhs340_
                    d_2519_v42_ = rhs341_
                    d_2522_v45_: D17
                    d_2522_v45_ = D17_DC46((self).f13)
                    d_2523_v46_: T1
                    nw406_ = C5()
                    nw406_.ctor__((self).f15, self.f12, ((self).f13) | ((d_2522_v45_).cf69), (self).f14)
                    d_2523_v46_ = nw406_
                    d_2524_v47_: _dafny.MultiSet
                    d_2524_v47_ = _dafny.MultiSet([not(((self).f16) < (_dafny.SeqWithoutIsStrInference([(D11_DC32(d_2455_v1_, len(d_2454_v0_))).cf46, d_2455_v1_]))), False])
                    d_2525_v48_: _dafny.Map
                    d_2525_v48_ = _dafny.Map({d_2455_v1_: d_2524_v47_})
                    rhs342_ = d_2455_v1_
                    rhs343_ = d_2523_v46_
                    rhs344_ = ((d_2524_v47_ if (self).f15 else ((d_2525_v48_)[(self).fm8((self).f15, (self).f15, globalState)] if ((self).fm8((self).f15, (self).f15, globalState)) in (d_2525_v48_) else d_2524_v47_))).set((self).f15, default__.abs((0) - (len(d_2454_v0_))))
                    rhs345_ = (self).f15
                    r1 = rhs342_
                    d_2523_v46_ = rhs343_
                    d_2524_v47_ = rhs344_
                    r0 = rhs345_
                    pass
            pass
        d_2526_i5_: int
        d_2526_i5_ = 0
        with _dafny.label("15"):
            while True:
                with _dafny.c_label("15"):
                    if (d_2526_i5_) >= (100):
                        raise _dafny.Break("15")
                    d_2526_i5_ = (d_2526_i5_) + (1)
                    rhs346_ = (749) != ((d_2455_v1_) * (d_2455_v1_))
                    rhs347_ = not((self).f15)
                    lhs212_ = globalState
                    lhs213_ = globalState
                    lhs212_.f4 = rhs346_
                    lhs213_.f4 = rhs347_
                    d_2527_v49_: _dafny.Array
                    def lambda136_(d_2528_v1_):
                        def lambda137_(d_2529_i6_):
                            return default__.safeModuloInt(d_2529_i6_, d_2528_v1_)

                        return lambda137_

                    init73_ = lambda136_(d_2455_v1_)
                    nw407_ = _dafny.Array(None, 17)
                    for i0_73_ in range(nw407_.length(0)):
                        nw407_[i0_73_] = init73_(i0_73_)
                    d_2527_v49_ = nw407_
                    index394_ = default__.safeIndex(341, (d_2527_v49_).length(0))
                    (d_2527_v49_)[index394_] = d_2455_v1_
                    index395_ = default__.safeIndex(341, (d_2527_v49_).length(0))
                    (d_2527_v49_)[index395_] = d_2455_v1_
                    d_2530_v50_: _dafny.Array
                    nw408_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                    d_2530_v50_ = nw408_
                    index396_ = default__.safeIndex(253, (d_2530_v50_).length(0))
                    (d_2530_v50_)[index396_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kga"))
                    index397_ = default__.safeIndex(253, (d_2530_v50_).length(0))
                    (d_2530_v50_)[index397_] = d_2454_v0_
                    d_2455_v1_ = (0) - ((d_2527_v49_)[default__.safeIndex(341, (d_2527_v49_).length(0))])
                    pass
            pass
        d_2455_v1_ = default__.safeModuloInt(len(((self).f16) + ((self).f16)), ((D11_DC32(d_2455_v1_, d_2455_v1_)).cf46) - (len(d_2454_v0_)))
        r0 = (self).f15
        d_2531_v51_: _dafny.Map
        d_2531_v51_ = _dafny.Map({(self).f15: d_2455_v1_})
        d_2532_v52_: D8
        d_2532_v52_ = D8_DC25((self).f15, d_2455_v1_, d_2455_v1_, (d_2531_v51_).set((self).f15, d_2455_v1_))
        pat_let_tv41_ = d_2454_v0_
        pat_let_tv42_ = d_2531_v51_
        pat_let_tv43_ = d_2455_v1_
        def iife166_(_pat_let49_0):
            def iife167_(d_2533_dt__update__tmp_h1_):
                def iife168_(_pat_let50_0):
                    def iife169_(d_2534_dt__update_hcf39_h0_):
                        def iife170_(_pat_let51_0):
                            def iife171_(d_2535_dt__update_hcf40_h0_):
                                def iife172_(_pat_let52_0):
                                    def iife173_(d_2536_dt__update_hcf38_h0_):
                                        return D8_DC25((d_2533_dt__update__tmp_h1_).cf37, d_2536_dt__update_hcf38_h0_, d_2534_dt__update_hcf39_h0_, d_2535_dt__update_hcf40_h0_)
                                    return iife173_(_pat_let52_0)
                                return iife172_(pat_let_tv43_)
                            return iife171_(_pat_let51_0)
                        return iife170_(pat_let_tv42_)
                    return iife169_(_pat_let50_0)
                return iife168_(len(pat_let_tv41_))
            return iife167_(_pat_let49_0)
        r1 = (iife166_(d_2532_v52_)).cf39
        r2 = _dafny.Set({(self).f14, (self).f14, default__.fm32((self).f15, d_2514_v38_, (default__.fm20((self).f15, d_2455_v1_, globalState)).cardinality, (0) - (-843), globalState), (self).f14, (self).f14})
        r3 = ((d_2454_v0_) == (d_2454_v0_) if (d_2455_v1_) < (-187) else (self).f15)
        return r0, r1, r2, r3

    def m5(self, p0, p1, p2, globalState):
        r0: bool = False
        d_2537_v0_: _dafny.Map
        d_2537_v0_ = _dafny.Map({(self).f15: (self).f15})
        d_2537_v0_ = (d_2537_v0_).set((self).f15, (self).f15)
        d_2538_v1_: _dafny.Seq
        d_2538_v1_ = _dafny.SeqWithoutIsStrInference([((self).f16).set(default__.safeIndex(p2, len((self).f16)), p2)])
        d_2539_v2_: _dafny.Map
        d_2539_v2_ = _dafny.Map({len(d_2538_v1_): (self).f15})
        d_2539_v2_ = (d_2539_v2_).set(p2, (self).f15)
        d_2540_v3_: _dafny.Seq
        d_2540_v3_ = _dafny.SeqWithoutIsStrInference([True])
        d_2541_v4_: str
        d_2541_v4_ = _dafny.CodePoint('n')
        d_2542_v5_: _dafny.Map
        d_2542_v5_ = _dafny.Map({((self).f15) == ((self).f15): d_2541_v4_})
        d_2543_v6_: int
        d_2543_v6_ = 875
        d_2544_v7_: _dafny.MultiSet
        d_2544_v7_ = _dafny.MultiSet([(self).f15, (self).f15])
        rhs348_ = ((self).f14) + ((self).f14)
        rhs349_ = d_2542_v5_
        rhs350_ = p2
        rhs351_ = ((d_2544_v7_).intersection(d_2544_v7_)).ispropersubset(_dafny.MultiSet((self).f14))
        d_2540_v3_ = rhs348_
        d_2542_v5_ = rhs349_
        d_2543_v6_ = rhs350_
        r0 = rhs351_
        d_2545_v8_: C5
        nw409_ = C5()
        nw409_.ctor__((self).f15, self.f12, (_dafny.Map({(self).f15: d_2537_v0_})) | ((self).f13), d_2540_v3_)
        d_2545_v8_ = nw409_
        d_2546_v9_: T0
        nw410_ = C5()
        nw410_.ctor__((self).f15, self.f12, _dafny.Map({not(False): d_2537_v0_}), (self).f14)
        d_2546_v9_ = nw410_
        d_2547_v10_: _dafny.Map
        d_2547_v10_ = _dafny.Map({d_2546_v9_: p2})
        r0 = (d_2546_v9_) in (d_2547_v10_)
        d_2548_v11_: str
        d_2548_v11_ = d_2541_v4_
        d_2549_v12_: _dafny.Map
        d_2549_v12_ = _dafny.Map({631: d_2543_v6_})
        d_2550_v13_: _dafny.Array
        nw411_ = _dafny.Array(None, 28)
        nw411_[int(0)] = _dafny.CodePoint('h')
        nw411_[int(1)] = d_2541_v4_
        nw411_[int(2)] = _dafny.CodePoint('m')
        nw411_[int(3)] = d_2541_v4_
        nw411_[int(4)] = d_2541_v4_
        nw411_[int(5)] = (d_2548_v11_)
        nw411_[int(6)] = d_2541_v4_
        nw411_[int(7)] = d_2541_v4_
        nw411_[int(8)] = _dafny.CodePoint('n')
        nw411_[int(9)] = d_2541_v4_
        nw411_[int(10)] = d_2541_v4_
        nw411_[int(11)] = d_2541_v4_
        nw411_[int(12)] = d_2541_v4_
        nw411_[int(13)] = d_2541_v4_
        nw411_[int(14)] = d_2541_v4_
        nw411_[int(15)] = (_dafny.CodePoint('i') if not((self).f15) else d_2541_v4_)
        nw411_[int(16)] = d_2541_v4_
        nw411_[int(17)] = d_2541_v4_
        nw411_[int(18)] = default__.fm39(((d_2549_v12_)[d_2543_v6_] if (d_2543_v6_) in (d_2549_v12_) else d_2543_v6_), (self).f15, p2, p1, globalState)
        nw411_[int(19)] = d_2541_v4_
        nw411_[int(20)] = d_2541_v4_
        nw411_[int(21)] = d_2541_v4_
        nw411_[int(22)] = d_2541_v4_
        nw411_[int(23)] = (d_2548_v11_)
        nw411_[int(24)] = d_2541_v4_
        nw411_[int(25)] = d_2541_v4_
        nw411_[int(26)] = d_2541_v4_
        nw411_[int(27)] = d_2541_v4_
        d_2550_v13_ = nw411_
        index398_ = default__.safeIndex(438, (d_2550_v13_).length(0))
        (d_2550_v13_)[index398_] = _dafny.CodePoint('q')
        index399_ = default__.safeIndex(438, (d_2550_v13_).length(0))
        (d_2550_v13_)[index399_] = d_2541_v4_
        r0 = not ((p2) in (_dafny.MultiSet([len(p1)]))) or ((self).f15)
        return r0

    def m6(self, globalState):
        r0: bool = False
        d_2551_v0_: int
        d_2551_v0_ = 614
        d_2552_v1_: _dafny.Array
        nw412_ = _dafny.Array(None, 23)
        nw412_[int(0)] = (self).f15
        nw412_[int(1)] = False
        nw412_[int(2)] = False
        nw412_[int(3)] = (self).f15
        nw412_[int(4)] = (self).f15
        nw412_[int(5)] = False
        nw412_[int(6)] = (self).f15
        nw412_[int(7)] = (self).f15
        nw412_[int(8)] = not((self).f15)
        nw412_[int(9)] = (self).f15
        nw412_[int(10)] = (self).f15
        nw412_[int(11)] = False
        nw412_[int(12)] = (self).f15
        nw412_[int(13)] = False
        nw412_[int(14)] = (self).f15
        nw412_[int(15)] = (self).f15
        nw412_[int(16)] = (self).f15
        nw412_[int(17)] = (self).f15
        nw412_[int(18)] = (self).f15
        nw412_[int(19)] = (self).f15
        nw412_[int(20)] = (self).f15
        nw412_[int(21)] = (self).f15
        nw412_[int(22)] = (self).f15
        d_2552_v1_ = nw412_
        d_2553_v2_: _dafny.Map
        d_2553_v2_ = _dafny.Map({431: d_2552_v1_})
        d_2554_v3_: _dafny.Map
        d_2554_v3_ = _dafny.Map({((self).f14)[default__.safeIndex(len((d_2553_v2_).set(d_2551_v0_, d_2552_v1_)), len((self).f14))]: (self).f15})
        hi14_ = len(d_2554_v3_)
        for d_2555_i0_ in range(d_2551_v0_, hi14_):
            d_2556_v4_: C1
            nw413_ = C1()
            nw413_.ctor__((D5_DC16((self).f15)).cf26, (self).f15, ((self).f13) | ((self).f13), (_dafny.SeqWithoutIsStrInference([default__.fm2(not((self).f15), globalState)])).set(default__.safeIndex((0) - (d_2555_i0_), len(_dafny.SeqWithoutIsStrInference([default__.fm2(not((self).f15), globalState)]))), (self).f15), self.f12)
            d_2556_v4_ = nw413_
            d_2557_v5_: D2
            d_2557_v5_ = D2_DC6(d_2552_v1_, (D2_DC6(d_2552_v1_, d_2552_v1_)).cf9)
            d_2558_v6_: _dafny.Array
            nw414_ = _dafny.Array(None, 22)
            nw414_[int(0)] = d_2552_v1_
            nw414_[int(1)] = (d_2557_v5_).cf9
            nw414_[int(2)] = d_2552_v1_
            nw414_[int(3)] = d_2552_v1_
            nw414_[int(4)] = d_2552_v1_
            nw414_[int(5)] = d_2552_v1_
            nw414_[int(6)] = d_2552_v1_
            nw414_[int(7)] = d_2552_v1_
            nw414_[int(8)] = d_2552_v1_
            nw414_[int(9)] = d_2552_v1_
            nw414_[int(10)] = d_2552_v1_
            nw414_[int(11)] = d_2552_v1_
            nw414_[int(12)] = d_2552_v1_
            nw414_[int(13)] = d_2552_v1_
            nw414_[int(14)] = d_2552_v1_
            nw414_[int(15)] = d_2552_v1_
            nw414_[int(16)] = d_2552_v1_
            nw414_[int(17)] = d_2552_v1_
            nw414_[int(18)] = d_2552_v1_
            nw414_[int(19)] = d_2552_v1_
            nw414_[int(20)] = (d_2557_v5_).cf8
            nw414_[int(21)] = d_2552_v1_
            d_2558_v6_ = nw414_
            d_2559_v7_: _dafny.Map
            d_2559_v7_ = _dafny.Map({(d_2556_v4_).f17: d_2558_v6_})
            d_2559_v7_ = (d_2559_v7_).set((self).f15, d_2558_v6_)
            if not (not(False)) or ((self).f15):
                d_2560_v8_: _dafny.Array
                nw415_ = _dafny.Array(_dafny.Map({}), 12)
                d_2560_v8_ = nw415_
                d_2560_v8_ = d_2560_v8_
                d_2561_v9_: _dafny.Seq
                d_2561_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfwv"))
                d_2562_v10_: _dafny.Map
                d_2562_v10_ = _dafny.Map({d_2552_v1_: len((self).fm6(d_2561_v9_, globalState))})
                d_2562_v10_ = (d_2562_v10_).set(d_2552_v1_, d_2555_i0_)
                d_2563_v11_: bool
                d_2564_v12_: int
                d_2565_v13_: _dafny.Set
                d_2566_v14_: bool
                out123_: bool
                out124_: int
                out125_: _dafny.Set
                out126_: bool
                out123_, out124_, out125_, out126_ = (d_2556_v4_).m3(globalState)
                d_2563_v11_ = out123_
                d_2564_v12_ = out124_
                d_2565_v13_ = out125_
                d_2566_v14_ = out126_
                d_2567_v15_: D0
                d_2567_v15_ = D0_DC2(d_2551_v0_, not(d_2566_v14_))
                d_2568_v16_: str
                d_2568_v16_ = _dafny.CodePoint('s')
                d_2569_v17_: _dafny.MultiSet
                d_2569_v17_ = _dafny.MultiSet([d_2568_v16_, d_2568_v16_, d_2568_v16_, d_2568_v16_])
                d_2570_v18_: _dafny.Seq
                d_2571_v19_: bool
                d_2572_v20_: int
                d_2573_v21_: bool
                out127_: _dafny.Seq
                out128_: bool
                out129_: int
                out130_: bool
                out127_, out128_, out129_, out130_ = default__.m0((self).f15, d_2567_v15_, (d_2569_v17_).issubset(d_2569_v17_), globalState)
                d_2570_v18_ = out127_
                d_2571_v19_ = out128_
                d_2572_v20_ = out129_
                d_2573_v21_ = out130_
                d_2551_v0_ = d_2551_v0_
            elif True:
                d_2574_v22_: _dafny.MultiSet
                d_2574_v22_ = _dafny.MultiSet([(d_2556_v4_).f17])
                d_2574_v22_ = d_2574_v22_
                d_2575_v23_: _dafny.Set
                d_2575_v23_ = _dafny.Set({(self).f14})
                d_2576_v24_: _dafny.Seq
                d_2576_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ai"))
                d_2577_v25_: C4
                nw416_ = C4()
                nw416_.ctor__()
                d_2577_v25_ = nw416_
                d_2578_v26_: _dafny.Array
                nw417_ = _dafny.Array(None, 28)
                nw417_[int(0)] = d_2577_v25_
                nw417_[int(1)] = d_2577_v25_
                nw417_[int(2)] = d_2577_v25_
                nw417_[int(3)] = d_2577_v25_
                nw417_[int(4)] = d_2577_v25_
                nw417_[int(5)] = d_2577_v25_
                nw417_[int(6)] = d_2577_v25_
                nw417_[int(7)] = d_2577_v25_
                nw417_[int(8)] = d_2577_v25_
                nw417_[int(9)] = d_2577_v25_
                nw417_[int(10)] = d_2577_v25_
                nw417_[int(11)] = d_2577_v25_
                nw417_[int(12)] = d_2577_v25_
                nw417_[int(13)] = d_2577_v25_
                nw417_[int(14)] = d_2577_v25_
                nw417_[int(15)] = d_2577_v25_
                nw417_[int(16)] = d_2577_v25_
                nw417_[int(17)] = d_2577_v25_
                nw417_[int(18)] = d_2577_v25_
                nw417_[int(19)] = d_2577_v25_
                nw417_[int(20)] = d_2577_v25_
                nw417_[int(21)] = d_2577_v25_
                nw417_[int(22)] = d_2577_v25_
                nw417_[int(23)] = d_2577_v25_
                nw417_[int(24)] = d_2577_v25_
                nw417_[int(25)] = d_2577_v25_
                nw417_[int(26)] = d_2577_v25_
                nw417_[int(27)] = d_2577_v25_
                d_2578_v26_ = nw417_
                d_2579_v27_: _dafny.MultiSet
                d_2579_v27_ = _dafny.MultiSet([d_2578_v26_])
                rhs352_ = not(((len(d_2576_v24_)) * (d_2555_i0_)) <= ((d_2555_i0_) + ((d_2579_v27_).cardinality)))
                rhs353_ = (0) - (-885)
                rhs354_ = (self).f15
                rhs355_ = (d_2575_v23_) | (d_2575_v23_)
                lhs214_ = d_2556_v4_
                lhs215_ = globalState
                lhs214_.f18 = rhs352_
                d_2551_v0_ = rhs353_
                lhs215_.f4 = rhs354_
                d_2575_v23_ = rhs355_
                d_2580_v28_: _dafny.MultiSet
                d_2580_v28_ = _dafny.MultiSet([780])
                d_2581_v29_: _dafny.Seq
                d_2581_v29_ = _dafny.SeqWithoutIsStrInference([(d_2556_v4_.f18) and (True), not(False), ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2582_i1_ in range(default__.abs(258))])))) >= (((d_2580_v28_)[852] if (852) in (d_2580_v28_) else d_2551_v0_))])
                d_2581_v29_ = (d_2581_v29_) + (_dafny.SeqWithoutIsStrInference([(d_2556_v4_).f17, d_2556_v4_.f18, False]))
                d_2583_v30_: _dafny.Array
                def lambda138_(d_2584_v0_):
                    def lambda139_(d_2585_i2_):
                        return (d_2585_i2_) + (d_2584_v0_)

                    return lambda139_

                init74_ = lambda138_(d_2551_v0_)
                nw418_ = _dafny.Array(None, 4)
                for i0_74_ in range(nw418_.length(0)):
                    nw418_[i0_74_] = init74_(i0_74_)
                d_2583_v30_ = nw418_
                index400_ = default__.safeIndex(236, (d_2583_v30_).length(0))
                (d_2583_v30_)[index400_] = (d_2551_v0_) * (d_2555_i0_)
                index401_ = default__.safeIndex(236, (d_2583_v30_).length(0))
                (d_2583_v30_)[index401_] = d_2555_i0_
                d_2586_v31_: D11
                d_2586_v31_ = D11_DC31()
                d_2587_v32_: _dafny.Array
                nw419_ = _dafny.Array(None, 17)
                nw419_[int(0)] = d_2586_v31_
                nw419_[int(1)] = d_2586_v31_
                nw419_[int(2)] = d_2586_v31_
                nw419_[int(3)] = d_2586_v31_
                nw419_[int(4)] = d_2586_v31_
                nw419_[int(5)] = d_2586_v31_
                nw419_[int(6)] = D11_DC31()
                nw419_[int(7)] = d_2586_v31_
                nw419_[int(8)] = D11_DC31()
                nw419_[int(9)] = d_2586_v31_
                nw419_[int(10)] = d_2586_v31_
                nw419_[int(11)] = d_2586_v31_
                nw419_[int(12)] = d_2586_v31_
                nw419_[int(13)] = D11_DC31()
                nw419_[int(14)] = d_2586_v31_
                nw419_[int(15)] = D11_DC31()
                nw419_[int(16)] = d_2586_v31_
                d_2587_v32_ = nw419_
                index402_ = default__.safeIndex(259, (d_2587_v32_).length(0))
                (d_2587_v32_)[index402_] = D11_DC31()
                d_2588_v33_: D12
                d_2588_v33_ = D12_DC37(d_2576_v24_, d_2556_v4_.f18)
                d_2589_v34_: D4
                d_2589_v34_ = D4_DC14(d_2556_v4_.f18)
                d_2590_v35_: _dafny.Map
                d_2590_v35_ = _dafny.Map({True: default__.fm40(d_2589_v34_, (d_2583_v30_)[default__.safeIndex(236, (d_2583_v30_).length(0))], d_2556_v4_.f18, (self).f15, globalState)})
                d_2591_v36_: _dafny.Array
                nw420_ = _dafny.Array(None, 22)
                nw420_[int(0)] = (d_2576_v24_) + (d_2576_v24_)
                nw420_[int(1)] = d_2576_v24_
                nw420_[int(2)] = d_2576_v24_
                nw420_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snwcxjku"))
                nw420_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fhlw"))
                nw420_[int(5)] = (d_2588_v33_).cf59
                nw420_[int(6)] = d_2576_v24_
                nw420_[int(7)] = d_2576_v24_
                nw420_[int(8)] = (d_2576_v24_) + (d_2576_v24_)
                nw420_[int(9)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2592_i3_ in range(default__.abs(-67))])) + (d_2576_v24_)
                nw420_[int(10)] = d_2576_v24_
                nw420_[int(11)] = d_2576_v24_
                nw420_[int(12)] = ((d_2590_v35_)[d_2556_v4_.f18] if (d_2556_v4_.f18) in (d_2590_v35_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bn")))
                nw420_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2593_i4_ in range(default__.abs(438))])
                nw420_[int(14)] = (d_2576_v24_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xykacskdp")))
                nw420_[int(15)] = d_2576_v24_
                nw420_[int(16)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2594_i5_ in range(default__.abs(71))])) + (d_2576_v24_)
                nw420_[int(17)] = d_2576_v24_
                nw420_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftpjtp"))
                nw420_[int(19)] = d_2576_v24_
                nw420_[int(20)] = (d_2576_v24_) + (d_2576_v24_)
                nw420_[int(21)] = d_2576_v24_
                d_2591_v36_ = nw420_
                index403_ = default__.safeIndex(792, (d_2591_v36_).length(0))
                (d_2591_v36_)[index403_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htrpkjrkt"))) + (d_2576_v24_)
                d_2595_v37_: _dafny.Set
                d_2595_v37_ = _dafny.Set({(self).f15})
                index404_ = default__.safeIndex(259, (d_2587_v32_).length(0))
                index405_ = default__.safeIndex(792, (d_2591_v36_).length(0))
                rhs356_ = ((_dafny.Set({not((self).f15)})) - (d_2595_v37_)).ispropersubset(_dafny.Set({(self).f15}))
                rhs357_ = d_2576_v24_
                rhs358_ = d_2586_v31_
                rhs359_ = d_2576_v24_
                rhs360_ = d_2556_v4_.f18
                lhs216_ = d_2556_v4_
                lhs217_ = d_2587_v32_
                lhs218_ = default__.safeIndex(259, (d_2587_v32_).length(0))
                lhs219_ = d_2591_v36_
                lhs220_ = default__.safeIndex(792, (d_2591_v36_).length(0))
                lhs221_ = d_2556_v4_
                lhs216_.f18 = rhs356_
                d_2576_v24_ = rhs357_
                lhs217_[lhs218_] = rhs358_
                lhs219_[lhs220_] = rhs359_
                lhs221_.f18 = rhs360_
            d_2596_v38_: _dafny.Array
            def lambda140_(d_2597_v0_):
                def lambda141_(d_2598_i6_):
                    return (d_2598_i6_) + (d_2597_v0_)

                return lambda141_

            init75_ = lambda140_(d_2551_v0_)
            nw421_ = _dafny.Array(None, 20)
            for i0_75_ in range(nw421_.length(0)):
                nw421_[i0_75_] = init75_(i0_75_)
            d_2596_v38_ = nw421_
            index406_ = default__.safeIndex(597, (d_2596_v38_).length(0))
            (d_2596_v38_)[index406_] = d_2555_i0_
            index407_ = default__.safeIndex(597, (d_2596_v38_).length(0))
            (d_2596_v38_)[index407_] = d_2551_v0_
        d_2599_v39_: str
        d_2599_v39_ = _dafny.CodePoint('o')
        d_2600_v40_: _dafny.Map
        d_2600_v40_ = _dafny.Map({True: d_2599_v39_})
        index408_ = default__.safeIndex(209, (d_2552_v1_).length(0))
        (d_2552_v1_)[index408_] = not((_dafny.Map({(self).f15: d_2599_v39_})) != (d_2600_v40_))
        d_2601_v41_: _dafny.Seq
        d_2601_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jimfygn"))
        d_2602_v42_: _dafny.Seq
        d_2602_v42_ = _dafny.SeqWithoutIsStrInference([d_2601_v41_, d_2601_v41_, d_2601_v41_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twglwr"))])
        d_2603_v43_: _dafny.Map
        d_2603_v43_ = _dafny.Map({(self).f15: d_2602_v42_})
        index409_ = default__.safeIndex(209, (d_2552_v1_).length(0))
        (d_2552_v1_)[index409_] = (len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2604_i7_ in range(default__.abs(931))]), d_2601_v41_])) + (((d_2603_v43_)[(self).f15] if ((self).f15) in (d_2603_v43_) else _dafny.SeqWithoutIsStrInference([d_2601_v41_ for d_2605_i8_ in range(default__.abs(86))]))))) <= (d_2551_v0_)
        d_2606_v44_: _dafny.Set
        d_2606_v44_ = _dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('e'), d_2599_v39_, d_2599_v39_, _dafny.CodePoint('f')})
        d_2607_v45_: _dafny.Map
        d_2607_v45_ = _dafny.Map({-38: d_2606_v44_})
        d_2607_v45_ = d_2607_v45_
        d_2608_v46_: _dafny.Map
        d_2608_v46_ = _dafny.Map({(self).f15: d_2551_v0_})
        d_2609_v47_: D8
        d_2609_v47_ = D8_DC23(d_2608_v46_)
        d_2610_v48_: _dafny.Map
        d_2610_v48_ = _dafny.Map({d_2609_v47_: False})
        d_2610_v48_ = (d_2610_v48_).set(d_2609_v47_, (self).f15)
        if (self).f15:
            if (self).f15:
                d_2599_v39_ = d_2599_v39_
                d_2611_v49_: _dafny.Seq
                d_2611_v49_ = _dafny.SeqWithoutIsStrInference([not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]), (d_2554_v3_) != (d_2554_v3_), (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]])
                d_2612_v50_: _dafny.Array
                def lambda142_(d_2613_v1_):
                    def lambda143_(d_2614_i9_):
                        return _dafny.Map({(d_2613_v1_)[default__.safeIndex(209, (d_2613_v1_).length(0))]: (self).f15})

                    return lambda143_

                init76_ = lambda142_(d_2552_v1_)
                nw422_ = _dafny.Array(None, 7)
                for i0_76_ in range(nw422_.length(0)):
                    nw422_[i0_76_] = init76_(i0_76_)
                d_2612_v50_ = nw422_
                d_2615_v51_: _dafny.Array
                d_2615_v51_ = d_2612_v50_
                d_2616_v52_: _dafny.Seq
                d_2616_v52_ = _dafny.SeqWithoutIsStrInference([d_2612_v50_, d_2612_v50_])
                d_2617_v53_: D12
                d_2617_v53_ = D12_DC37(_dafny.SeqWithoutIsStrInference([d_2599_v39_ for d_2618_i10_ in range(default__.abs(548))]), (self).f15)
                d_2619_v54_: _dafny.Array
                nw423_ = _dafny.Array(None, 26)
                nw423_[int(0)] = d_2612_v50_
                nw423_[int(1)] = (d_2612_v50_ if (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))] else d_2612_v50_)
                nw423_[int(2)] = d_2612_v50_
                nw423_[int(3)] = d_2612_v50_
                nw423_[int(4)] = d_2612_v50_
                nw423_[int(5)] = d_2612_v50_
                nw423_[int(6)] = d_2612_v50_
                nw423_[int(7)] = d_2612_v50_
                nw423_[int(8)] = d_2612_v50_
                nw423_[int(9)] = d_2612_v50_
                nw423_[int(10)] = d_2612_v50_
                nw423_[int(11)] = d_2612_v50_
                nw423_[int(12)] = d_2612_v50_
                nw423_[int(13)] = d_2612_v50_
                nw423_[int(14)] = d_2612_v50_
                nw423_[int(15)] = d_2612_v50_
                nw423_[int(16)] = d_2612_v50_
                nw423_[int(17)] = d_2612_v50_
                nw423_[int(18)] = d_2612_v50_
                nw423_[int(19)] = (d_2612_v50_)
                nw423_[int(20)] = d_2612_v50_
                nw423_[int(21)] = d_2612_v50_
                nw423_[int(22)] = d_2612_v50_
                nw423_[int(23)] = (d_2616_v52_)[default__.safeIndex(len((d_2617_v53_).cf59), len(d_2616_v52_))]
                nw423_[int(24)] = d_2612_v50_
                nw423_[int(25)] = d_2612_v50_
                d_2619_v54_ = nw423_
                index410_ = default__.safeIndex(935, (d_2619_v54_).length(0))
                (d_2619_v54_)[index410_] = d_2612_v50_
                d_2620_v55_: _dafny.MultiSet
                d_2620_v55_ = _dafny.MultiSet([(self).f15])
                d_2621_v56_: _dafny.MultiSet
                d_2621_v56_ = _dafny.MultiSet([(d_2620_v55_).cardinality])
                d_2622_v57_: _dafny.Map
                d_2622_v57_ = _dafny.Map({d_2621_v56_: d_2599_v39_})
                index411_ = default__.safeIndex(935, (d_2619_v54_).length(0))
                rhs361_ = (((self).f14) + (d_2611_v49_) if not((d_2551_v0_) != (len(d_2601_v41_))) else (self).f14)
                rhs362_ = d_2612_v50_
                rhs363_ = (len(d_2601_v41_)) - (len(_dafny.SeqWithoutIsStrInference([(D4_DC13((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (self).f15, d_2551_v0_, d_2551_v0_, False)).cf21 for d_2623_i11_ in range(default__.abs(419))])))
                rhs364_ = d_2622_v57_
                lhs222_ = d_2619_v54_
                lhs223_ = default__.safeIndex(935, (d_2619_v54_).length(0))
                d_2611_v49_ = rhs361_
                lhs222_[lhs223_] = rhs362_
                d_2551_v0_ = rhs363_
                d_2622_v57_ = rhs364_
                d_2624_v58_: _dafny.MultiSet
                d_2624_v58_ = _dafny.MultiSet([d_2601_v41_, d_2601_v41_])
                d_2551_v0_ = (((d_2624_v58_).intersection(d_2624_v58_) if (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))] else (_dafny.MultiSet(d_2602_v42_)) - (_dafny.MultiSet([d_2601_v41_])))).cardinality
                d_2625_v59_: C4
                nw424_ = C4()
                nw424_.ctor__()
                d_2625_v59_ = nw424_
                d_2551_v0_ = (self).fm7((d_2551_v0_) + (d_2551_v0_), globalState)
            elif True:
                d_2626_v60_: _dafny.MultiSet
                d_2626_v60_ = _dafny.MultiSet([d_2552_v1_, (d_2552_v1_ if (self).f15 else d_2552_v1_)])
                d_2627_v61_: _dafny.Seq
                d_2627_v61_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2552_v1_, d_2552_v1_, d_2552_v1_]), d_2626_v60_, (d_2626_v60_) | (d_2626_v60_), d_2626_v60_])
                d_2626_v60_ = (d_2627_v61_)[default__.safeIndex((d_2551_v0_) + (d_2551_v0_), len(d_2627_v61_))]
                d_2628_v62_: C4
                nw425_ = C4()
                nw425_.ctor__()
                d_2628_v62_ = nw425_
                r0 = (self).f15
                (globalState).f4 = not (not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])) or ((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])
                d_2629_v63_: _dafny.Array
                nw426_ = _dafny.Array(int(0), 23)
                d_2629_v63_ = nw426_
                index412_ = default__.safeIndex(591, (d_2629_v63_).length(0))
                (d_2629_v63_)[index412_] = default__.safeDivisionInt(-180, d_2551_v0_)
                index413_ = default__.safeIndex(591, (d_2629_v63_).length(0))
                (d_2629_v63_)[index413_] = d_2551_v0_
            if (self).f15:
                d_2630_v64_: D2
                d_2630_v64_ = D2_DC5(d_2551_v0_, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])
                pat_let_tv44_ = d_2608_v46_
                pat_let_tv45_ = d_2552_v1_
                pat_let_tv46_ = d_2552_v1_
                pat_let_tv47_ = d_2552_v1_
                pat_let_tv48_ = d_2552_v1_
                pat_let_tv49_ = d_2608_v46_
                d_2631_v65_: _dafny.Array
                nw427_ = _dafny.Array(None, 2)
                def iife174_(_pat_let53_0):
                    def iife175_(d_2632_dt__update__tmp_h1_):
                        def iife176_(_pat_let54_0):
                            def iife177_(d_2633_dt__update_hcf5_h0_):
                                def iife178_(_pat_let55_0):
                                    def iife179_(d_2634_dt__update_hcf6_h0_):
                                        return D2_DC5(d_2633_dt__update_hcf5_h0_, d_2634_dt__update_hcf6_h0_, (d_2632_dt__update__tmp_h1_).cf7)
                                    return iife179_(_pat_let55_0)
                                return iife178_((self).f15)
                            return iife177_(_pat_let54_0)
                        return iife176_(((pat_let_tv44_)[(pat_let_tv46_)[default__.safeIndex(209, (pat_let_tv45_).length(0))]] if ((pat_let_tv48_)[default__.safeIndex(209, (pat_let_tv47_).length(0))]) in (pat_let_tv49_) else -254))
                    return iife175_(_pat_let53_0)
                nw427_[int(0)] = iife174_(D2_DC5(910, not((self).f15), default__.fm2((self).f15, globalState)))
                nw427_[int(1)] = d_2630_v64_
                d_2631_v65_ = nw427_
                index414_ = default__.safeIndex(661, (d_2631_v65_).length(0))
                (d_2631_v65_)[index414_] = d_2630_v64_
                index415_ = default__.safeIndex(661, (d_2631_v65_).length(0))
                (d_2631_v65_)[index415_] = d_2630_v64_
                d_2635_v66_: _dafny.Array
                def lambda144_(d_2636_v41_):
                    def lambda145_(d_2637_i12_):
                        return d_2636_v41_

                    return lambda145_

                init77_ = lambda144_(d_2601_v41_)
                nw428_ = _dafny.Array(None, 18)
                for i0_77_ in range(nw428_.length(0)):
                    nw428_[i0_77_] = init77_(i0_77_)
                d_2635_v66_ = nw428_
                index416_ = default__.safeIndex(236, (d_2635_v66_).length(0))
                (d_2635_v66_)[index416_] = (d_2601_v41_).set(default__.safeIndex(d_2551_v0_, len(d_2601_v41_)), d_2599_v39_)
                index417_ = default__.safeIndex(236, (d_2635_v66_).length(0))
                (d_2635_v66_)[index417_] = ((d_2601_v41_) + (d_2601_v41_)) + ((d_2601_v41_) + (d_2601_v41_))
                d_2551_v0_ = (0) - (d_2551_v0_)
                d_2601_v41_ = (d_2635_v66_)[default__.safeIndex(236, (d_2635_v66_).length(0))]
                index418_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                rhs365_ = (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]
                rhs366_ = not(True)
                rhs367_ = not (default__.fm2((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], globalState)) or ((self).f15)
                lhs224_ = globalState
                lhs225_ = globalState
                lhs226_ = d_2552_v1_
                lhs227_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                lhs224_.f4 = rhs365_
                lhs225_.f4 = rhs366_
                lhs226_[lhs227_] = rhs367_
            elif True:
                rhs368_ = (self).f15
                rhs369_ = d_2551_v0_
                rhs370_ = (d_2551_v0_) + (d_2551_v0_)
                rhs371_ = (d_2602_v42_).set(default__.safeIndex((self).fm8(not((self).f15), (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], globalState), len(d_2602_v42_)), d_2601_v41_)
                lhs228_ = globalState
                lhs228_.f4 = rhs368_
                d_2551_v0_ = rhs369_
                d_2551_v0_ = rhs370_
                d_2602_v42_ = rhs371_
                d_2638_v67_: _dafny.Map
                d_2638_v67_ = _dafny.Map({d_2551_v0_: d_2608_v46_})
                d_2639_v68_: D11
                d_2639_v68_ = D11_DC30(d_2638_v67_)
                d_2639_v68_ = d_2639_v68_
                d_2640_v69_: _dafny.Array
                nw429_ = _dafny.Array(int(0), 29)
                d_2640_v69_ = nw429_
                d_2641_v70_: _dafny.Seq
                d_2641_v70_ = _dafny.SeqWithoutIsStrInference([d_2640_v69_, d_2640_v69_, d_2640_v69_, d_2640_v69_, d_2640_v69_])
                d_2640_v69_ = (d_2641_v70_)[default__.safeIndex(d_2551_v0_, len(d_2641_v70_))]
                d_2642_v71_: _dafny.MultiSet
                d_2642_v71_ = _dafny.MultiSet([d_2551_v0_])
                index419_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                (d_2552_v1_)[index419_] = (d_2642_v71_).issubset(d_2642_v71_)
                d_2643_v72_: _dafny.Map
                d_2643_v72_ = _dafny.Map({d_2551_v0_: (d_2551_v0_) - (d_2551_v0_)})
                d_2644_v73_: _dafny.Seq
                d_2644_v73_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14, (self).f14])
                d_2645_v74_: D2
                d_2645_v74_ = D2_DC4(d_2644_v73_)
                d_2646_v75_: _dafny.MultiSet
                d_2646_v75_ = _dafny.MultiSet([(self).f15, (self).f15])
                index420_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                index421_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                rhs372_ = (d_2643_v72_).set((0) - (d_2551_v0_), (_dafny.MultiSet(((self).f16).set(default__.safeIndex(len(d_2608_v46_), len((self).f16)), d_2551_v0_))).cardinality)
                rhs373_ = D2_DC4((_dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])) + (d_2644_v73_))
                rhs374_ = (d_2646_v75_) != ((d_2646_v75_) | (d_2646_v75_))
                rhs375_ = (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]
                rhs376_ = ((d_2551_v0_ if (self).f15 else 454) if (self).f15 else (self).fm8(not(not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])), default__.fm2(True, globalState), globalState))
                lhs229_ = d_2552_v1_
                lhs230_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                lhs231_ = d_2552_v1_
                lhs232_ = default__.safeIndex(209, (d_2552_v1_).length(0))
                d_2643_v72_ = rhs372_
                d_2645_v74_ = rhs373_
                lhs229_[lhs230_] = rhs374_
                lhs231_[lhs232_] = rhs375_
                d_2551_v0_ = rhs376_
            d_2647_v76_: _dafny.Array
            nw430_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_2647_v76_ = nw430_
            d_2648_v77_: _dafny.Array
            d_2648_v77_ = d_2647_v76_
            d_2649_v78_: _dafny.Array
            nw431_ = _dafny.Array(None, 29)
            nw431_[int(0)] = d_2647_v76_
            nw431_[int(1)] = (d_2648_v77_ if (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))] else d_2648_v77_)
            nw431_[int(2)] = d_2648_v77_
            nw431_[int(3)] = d_2648_v77_
            nw431_[int(4)] = d_2648_v77_
            nw431_[int(5)] = d_2648_v77_
            nw431_[int(6)] = d_2648_v77_
            nw431_[int(7)] = d_2648_v77_
            nw431_[int(8)] = d_2647_v76_
            nw431_[int(9)] = d_2647_v76_
            nw431_[int(10)] = d_2648_v77_
            nw431_[int(11)] = d_2648_v77_
            nw431_[int(12)] = d_2648_v77_
            nw431_[int(13)] = d_2648_v77_
            nw431_[int(14)] = d_2648_v77_
            nw431_[int(15)] = d_2648_v77_
            nw431_[int(16)] = d_2647_v76_
            nw431_[int(17)] = d_2647_v76_
            nw431_[int(18)] = d_2648_v77_
            nw431_[int(19)] = d_2648_v77_
            nw431_[int(20)] = d_2648_v77_
            nw431_[int(21)] = d_2647_v76_
            nw431_[int(22)] = d_2647_v76_
            nw431_[int(23)] = d_2648_v77_
            nw431_[int(24)] = (d_2647_v76_ if (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))] else d_2647_v76_)
            nw431_[int(25)] = d_2648_v77_
            nw431_[int(26)] = d_2648_v77_
            nw431_[int(27)] = d_2647_v76_
            nw431_[int(28)] = d_2648_v77_
            d_2649_v78_ = nw431_
            index422_ = default__.safeIndex(899, (d_2649_v78_).length(0))
            (d_2649_v78_)[index422_] = d_2648_v77_
            index423_ = default__.safeIndex(899, (d_2649_v78_).length(0))
            (d_2649_v78_)[index423_] = d_2647_v76_
            if (self).f15:
                d_2608_v46_ = (d_2608_v46_).set((d_2551_v0_) <= ((0) - (((d_2608_v46_)[(self).f15] if ((self).f15) in (d_2608_v46_) else 372))), default__.safeModuloInt(len((d_2601_v41_).set(default__.safeIndex(d_2551_v0_, len(d_2601_v41_)), _dafny.CodePoint('n'))), d_2551_v0_))
                d_2650_v79_: _dafny.Set
                d_2650_v79_ = _dafny.Set({(d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (self).f15})
                d_2651_v80_: T1
                nw432_ = C1()
                nw432_.ctor__((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (self).f15, (self).f13, (self).f14, self.f12)
                d_2651_v80_ = nw432_
                d_2652_v81_: D3
                d_2652_v81_ = D3_DC10((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], len(d_2650_v79_), d_2651_v80_)
                d_2653_v82_: D3
                d_2653_v82_ = D3_DC11(d_2652_v81_)
                pat_let_tv50_ = d_2652_v81_
                d_2654_v83_: _dafny.Array
                nw433_ = _dafny.Array(None, 19)
                nw433_[int(0)] = d_2653_v82_
                nw433_[int(1)] = D3_DC11(d_2652_v81_)
                nw433_[int(2)] = d_2653_v82_
                nw433_[int(3)] = d_2653_v82_
                nw433_[int(4)] = d_2653_v82_
                nw433_[int(5)] = d_2653_v82_
                nw433_[int(6)] = d_2653_v82_
                nw433_[int(7)] = d_2653_v82_
                nw433_[int(8)] = d_2653_v82_
                nw433_[int(9)] = d_2653_v82_
                nw433_[int(10)] = d_2653_v82_
                nw433_[int(11)] = D3_DC11(d_2652_v81_)
                nw433_[int(12)] = d_2653_v82_
                def iife180_(_pat_let56_0):
                    def iife181_(d_2655_dt__update__tmp_h11_):
                        def iife182_(_pat_let57_0):
                            def iife183_(d_2656_dt__update_hcf17_h0_):
                                return D3_DC11(d_2656_dt__update_hcf17_h0_)
                            return iife183_(_pat_let57_0)
                        return iife182_(pat_let_tv50_)
                    return iife181_(_pat_let56_0)
                nw433_[int(13)] = iife180_(d_2653_v82_)
                nw433_[int(14)] = d_2653_v82_
                nw433_[int(15)] = d_2653_v82_
                nw433_[int(16)] = d_2653_v82_
                nw433_[int(17)] = d_2653_v82_
                nw433_[int(18)] = d_2653_v82_
                d_2654_v83_ = nw433_
                d_2657_v84_: _dafny.Array
                d_2657_v84_ = d_2654_v83_
                d_2658_v85_: _dafny.MultiSet
                d_2658_v85_ = _dafny.MultiSet([(d_2657_v84_)])
                d_2659_v86_: _dafny.MultiSet
                d_2659_v86_ = _dafny.MultiSet([((d_2658_v85_ if (self).f15 else d_2658_v85_)).cardinality])
                d_2659_v86_ = (d_2659_v86_) - (d_2659_v86_)
                d_2551_v0_ = default__.safeDivisionInt(d_2551_v0_, d_2551_v0_)
                d_2660_v87_: T4
                nw434_ = C4()
                nw434_.ctor__()
                d_2660_v87_ = nw434_
                d_2661_v88_: _dafny.Seq
                d_2661_v88_ = _dafny.SeqWithoutIsStrInference([d_2552_v1_])
                rhs377_ = default__.safeDivisionInt((0) - ((d_2551_v0_) - (d_2551_v0_)), (734) + ((0) - (len(d_2661_v88_))))
                rhs378_ = d_2660_v87_
                rhs379_ = (self).f15
                d_2551_v0_ = rhs377_
                d_2660_v87_ = rhs378_
                r0 = rhs379_
                d_2601_v41_ = d_2601_v41_
            elif True:
                d_2551_v0_ = (self).fm5((self).f15, globalState)
                d_2662_v89_: T0
                nw435_ = C2()
                nw435_.ctor__((self).f16, (self).f15, ((self).f13).set(False, d_2554_v3_), (self).f14, self.f12)
                d_2662_v89_ = nw435_
                d_2663_v90_: _dafny.Map
                d_2663_v90_ = _dafny.Map({d_2551_v0_: _dafny.Map({d_2662_v89_: _dafny.Map({(self).f15: (self).f15})})})
                d_2664_v91_: _dafny.Map
                d_2664_v91_ = _dafny.Map({d_2662_v89_: d_2554_v3_})
                d_2665_v92_: _dafny.Map
                d_2665_v92_ = _dafny.Map({((d_2663_v90_)[-934] if (-934) in (d_2663_v90_) else d_2664_v91_): (self).fm5((self).f15, globalState)})
                d_2551_v0_ = default__.safeModuloInt((len(d_2601_v41_)) * (((d_2665_v92_)[_dafny.Map({d_2662_v89_: d_2554_v3_})] if (_dafny.Map({d_2662_v89_: d_2554_v3_})) in (d_2665_v92_) else d_2551_v0_)), d_2551_v0_)
                d_2666_v93_: _dafny.Set
                d_2666_v93_ = _dafny.Set({d_2551_v0_, 450})
                d_2667_v94_: _dafny.Map
                d_2667_v94_ = _dafny.Map({d_2666_v93_: (self).f15})
                d_2667_v94_ = (d_2667_v94_).set(d_2666_v93_, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])
                (globalState).f4 = (self).f15
                d_2668_v95_: C2
                nw436_ = C2()
                nw436_.ctor__((_dafny.SeqWithoutIsStrInference([d_2551_v0_])).set(default__.safeIndex(d_2551_v0_, len(_dafny.SeqWithoutIsStrInference([d_2551_v0_]))), d_2551_v0_), not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]), (self).f13, (self).f14, d_2662_v89_.f12)
                d_2668_v95_ = nw436_
            d_2601_v41_ = (d_2601_v41_) + ((_dafny.SeqWithoutIsStrInference([d_2599_v39_ for d_2669_i13_ in range(default__.abs(879))])) + (d_2601_v41_))
        elif True:
            (globalState).f4 = (self).f15
            d_2670_v96_: _dafny.Map
            d_2670_v96_ = _dafny.Map({d_2601_v41_: (self).f15})
            d_2671_v97_: D12
            d_2671_v97_ = D12_DC37(d_2601_v41_, (self).f15)
            index424_ = default__.safeIndex(209, (d_2552_v1_).length(0))
            (d_2552_v1_)[index424_] = ((d_2670_v96_)[d_2601_v41_] if (d_2601_v41_) in (d_2670_v96_) else (d_2671_v97_).cf60)
            if not((self).f15):
                d_2672_v98_: _dafny.Array
                nw437_ = _dafny.Array(int(0), 2)
                d_2672_v98_ = nw437_
                d_2673_v99_: C1
                nw438_ = C1()
                nw438_.ctor__((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], (self).f13, (self).f14, self.f12)
                d_2673_v99_ = nw438_
                d_2674_v100_: D11
                d_2674_v100_ = D11_DC33((self).f15, d_2551_v0_, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], d_2673_v99_, (d_2673_v99_).f17)
                index425_ = default__.safeIndex(879, (d_2672_v98_).length(0))
                (d_2672_v98_)[index425_] = (d_2551_v0_) * ((0) - ((d_2674_v100_).cf49))
                d_2675_v101_: D17
                d_2675_v101_ = D17_DC46((self).f13)
                d_2676_v102_: C2
                nw439_ = C2()
                nw439_.ctor__((self).f16, d_2673_v99_.f18, (d_2675_v101_).cf69, (self).f14, self.f12)
                d_2676_v102_ = nw439_
                d_2677_v103_: _dafny.Set
                d_2677_v103_ = _dafny.Set({d_2551_v0_})
                index426_ = default__.safeIndex(879, (d_2672_v98_).length(0))
                rhs380_ = len((_dafny.Set({d_2551_v0_})).intersection(d_2677_v103_))
                rhs381_ = (d_2551_v0_ if d_2673_v99_.f18 else (620) * (d_2551_v0_))
                rhs382_ = d_2676_v102_
                rhs383_ = not((58) != ((0) - ((0) - (d_2551_v0_))))
                lhs233_ = d_2672_v98_
                lhs234_ = default__.safeIndex(879, (d_2672_v98_).length(0))
                lhs235_ = d_2673_v99_
                d_2551_v0_ = rhs380_
                lhs233_[lhs234_] = rhs381_
                d_2676_v102_ = rhs382_
                lhs235_.f18 = rhs383_
                d_2678_v104_: D0
                d_2678_v104_ = D0_DC0(not(False))
                d_2679_v105_: D4
                d_2679_v105_ = D4_DC13((d_2678_v104_).cf0, not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]), 519, (self).fm5(d_2673_v99_.f18, globalState), (d_2673_v99_).f17)
                d_2680_v106_: _dafny.Map
                d_2680_v106_ = _dafny.Map({(d_2679_v105_).cf21: d_2551_v0_})
                d_2681_v107_: _dafny.MultiSet
                d_2681_v107_ = _dafny.MultiSet([d_2551_v0_, len(d_2601_v41_)])
                d_2680_v106_ = (d_2680_v106_).set(((d_2681_v107_).intersection(d_2681_v107_)).cardinality, d_2551_v0_)
                d_2682_v108_: _dafny.Array
                nw440_ = _dafny.Array(_dafny.Map({}), 19)
                d_2682_v108_ = nw440_
                d_2682_v108_ = d_2682_v108_
                r0 = (d_2673_v99_).f17
                d_2683_v109_: _dafny.Array
                nw441_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_2683_v109_ = nw441_
                d_2684_v110_: _dafny.MultiSet
                d_2684_v110_ = _dafny.MultiSet([(self).f15, False, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], d_2673_v99_.f18, (d_2673_v99_).f17])
                d_2685_v111_: _dafny.Array
                nw442_ = _dafny.Array(None, 2)
                nw442_[int(0)] = _dafny.MultiSet([d_2673_v99_.f18])
                nw442_[int(1)] = d_2684_v110_
                d_2685_v111_ = nw442_
                index427_ = default__.safeIndex(323, (d_2683_v109_).length(0))
                (d_2683_v109_)[index427_] = d_2685_v111_
                index428_ = default__.safeIndex(323, (d_2683_v109_).length(0))
                (d_2683_v109_)[index428_] = d_2685_v111_
            elif True:
                d_2686_v112_: _dafny.Array
                nw443_ = _dafny.Array(D14.default()(), 2)
                d_2686_v112_ = nw443_
                d_2687_v113_: T2
                nw444_ = C3()
                nw444_.ctor__((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], self.f12, not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]), (self).f13, (self).f14)
                d_2687_v113_ = nw444_
                d_2688_v114_: D14
                d_2688_v114_ = D14_DC40(d_2687_v113_)
                index429_ = default__.safeIndex(463, (d_2686_v112_).length(0))
                (d_2686_v112_)[index429_] = d_2688_v114_
                index430_ = default__.safeIndex(463, (d_2686_v112_).length(0))
                rhs384_ = d_2688_v114_
                rhs385_ = d_2601_v41_
                rhs386_ = d_2599_v39_
                rhs387_ = (0) - (d_2551_v0_)
                lhs236_ = d_2686_v112_
                lhs237_ = default__.safeIndex(463, (d_2686_v112_).length(0))
                lhs236_[lhs237_] = rhs384_
                d_2601_v41_ = rhs385_
                d_2599_v39_ = rhs386_
                d_2551_v0_ = rhs387_
                d_2551_v0_ = ((self).f16)[default__.safeIndex(len((d_2601_v41_) + (d_2601_v41_)), len((self).f16))]
                d_2599_v39_ = d_2599_v39_
                (globalState).f4 = (d_2551_v0_) > (d_2551_v0_)
                d_2689_v115_: _dafny.Array
                def lambda146_(d_2690_v0_):
                    def lambda147_(d_2691_i14_):
                        return (d_2691_i14_) - (d_2690_v0_)

                    return lambda147_

                init78_ = lambda146_(d_2551_v0_)
                nw445_ = _dafny.Array(None, 19)
                for i0_78_ in range(nw445_.length(0)):
                    nw445_[i0_78_] = init78_(i0_78_)
                d_2689_v115_ = nw445_
                index431_ = default__.safeIndex(336, (d_2689_v115_).length(0))
                (d_2689_v115_)[index431_] = default__.safeDivisionInt(d_2551_v0_, d_2551_v0_)
                d_2692_v116_: _dafny.MultiSet
                d_2692_v116_ = _dafny.MultiSet([(d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]])
                d_2693_v117_: _dafny.MultiSet
                d_2693_v117_ = _dafny.MultiSet([d_2551_v0_])
                index432_ = default__.safeIndex(336, (d_2689_v115_).length(0))
                rhs388_ = (default__.fm20((self).f15, d_2551_v0_, globalState)).issubset(d_2692_v116_)
                rhs389_ = d_2551_v0_
                rhs390_ = len((d_2687_v113_).f14)
                rhs391_ = default__.fm59(d_2693_v117_, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], default__.fm2((d_2687_v113_).f15, globalState), ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2551_v0_, d_2551_v0_, d_2551_v0_, 841]))) | (d_2693_v117_)).cardinality, globalState)
                lhs238_ = globalState
                lhs239_ = d_2689_v115_
                lhs240_ = default__.safeIndex(336, (d_2689_v115_).length(0))
                lhs238_.f4 = rhs388_
                lhs239_[lhs240_] = rhs389_
                d_2551_v0_ = rhs390_
                d_2554_v3_ = rhs391_
            d_2694_v118_: C3
            nw446_ = C3()
            nw446_.ctor__((self).f15, self.f12, not((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]), _dafny.Map({(self).f15: d_2554_v3_}), default__.fm15((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], d_2551_v0_, d_2551_v0_, 466, globalState))
            d_2694_v118_ = nw446_
            d_2695_v119_: _dafny.Map
            d_2695_v119_ = _dafny.Map({d_2694_v118_: d_2551_v0_})
            d_2696_v120_: C2
            nw447_ = C2()
            nw447_.ctor__((self).f16, (d_2695_v119_) != (d_2695_v119_), _dafny.Map({(d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]: d_2554_v3_}), (self).f14, self.f12)
            d_2696_v120_ = nw447_
            d_2696_v120_ = d_2696_v120_
            index433_ = default__.safeIndex(209, (d_2552_v1_).length(0))
            (d_2552_v1_)[index433_] = (self).f15
        hi15_ = len((d_2554_v3_) | (d_2554_v3_))
        for d_2697_i15_ in range(default__.safeModuloInt((self).fm5((self).f15, globalState), d_2551_v0_), hi15_):
            d_2698_v121_: _dafny.Array
            def lambda148_(d_2699_i15_):
                def lambda149_(d_2700_i16_):
                    return (_dafny.MultiSet([(self).f16, (self).f16, (self).f16])) | (_dafny.MultiSet([(self).f16, (self).f16, _dafny.SeqWithoutIsStrInference([d_2699_i15_ for d_2701_i17_ in range(default__.abs(402))])]))

                return lambda149_

            init79_ = lambda148_(d_2697_i15_)
            nw448_ = _dafny.Array(None, 13)
            for i0_79_ in range(nw448_.length(0)):
                nw448_[i0_79_] = init79_(i0_79_)
            d_2698_v121_ = nw448_
            index434_ = default__.safeIndex(801, (d_2698_v121_).length(0))
            (d_2698_v121_)[index434_] = _dafny.MultiSet([((self).f16).set(default__.safeIndex((self).fm7((0) - (((self).f16)[default__.safeIndex(d_2551_v0_, len((self).f16))]), globalState), len((self).f16)), d_2551_v0_), (self).f16])
            d_2702_v122_: _dafny.Seq
            d_2702_v122_ = _dafny.SeqWithoutIsStrInference([(self).f16, (self).f16])
            d_2703_v123_: _dafny.MultiSet
            d_2703_v123_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2697_i15_ for d_2704_i18_ in range(default__.abs(-238))]), ((d_2702_v122_)[default__.safeIndex((0) - (d_2551_v0_), len(d_2702_v122_))]).set(default__.safeIndex(d_2697_i15_, len((d_2702_v122_)[default__.safeIndex((0) - (d_2551_v0_), len(d_2702_v122_))])), d_2697_i15_)])
            d_2705_v124_: _dafny.Map
            d_2705_v124_ = _dafny.Map({(d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]: d_2703_v123_})
            d_2706_v125_: _dafny.Map
            d_2706_v125_ = _dafny.Map({(d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]: ((d_2705_v124_)[True] if (True) in (d_2705_v124_) else d_2703_v123_)})
            index435_ = default__.safeIndex(801, (d_2698_v121_).length(0))
            (d_2698_v121_)[index435_] = (d_2703_v123_) - (((d_2705_v124_)[False] if (False) in (d_2705_v124_) else d_2703_v123_))
            d_2707_v126_: _dafny.Array
            def lambda150_(d_2708_v39_):
                def lambda151_(d_2709_i19_):
                    return _dafny.SeqWithoutIsStrInference([d_2708_v39_ for d_2710_i20_ in range(default__.abs(375))])

                return lambda151_

            init80_ = lambda150_(d_2599_v39_)
            nw449_ = _dafny.Array(None, 14)
            for i0_80_ in range(nw449_.length(0)):
                nw449_[i0_80_] = init80_(i0_80_)
            d_2707_v126_ = nw449_
            index436_ = default__.safeIndex(248, (d_2707_v126_).length(0))
            def iife184_(_pat_let58_0):
                def iife185_(d_2711_dt__update__tmp_h12_):
                    def iife186_(_pat_let59_0):
                        def iife187_(d_2712_dt__update_hcf24_h0_):
                            return D4_DC14(d_2712_dt__update_hcf24_h0_)
                        return iife187_(_pat_let59_0)
                    return iife186_((self).f15)
                return iife185_(_pat_let58_0)
            (d_2707_v126_)[index436_] = default__.fm40(iife184_(D4_DC14((self).f15)), len(d_2610_v48_), (self).f15, (d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))], globalState)
            index437_ = default__.safeIndex(248, (d_2707_v126_).length(0))
            (d_2707_v126_)[index437_] = (_dafny.SeqWithoutIsStrInference([d_2599_v39_ for d_2713_i21_ in range(default__.abs(49))])) + (d_2601_v41_)
            d_2714_v127_: _dafny.Array
            nw450_ = _dafny.Array(None, 5)
            nw450_[int(0)] = d_2697_i15_
            nw450_[int(1)] = d_2697_i15_
            nw450_[int(2)] = d_2697_i15_
            nw450_[int(3)] = d_2697_i15_
            nw450_[int(4)] = d_2551_v0_
            d_2714_v127_ = nw450_
            d_2715_v128_: _dafny.Seq
            d_2715_v128_ = _dafny.SeqWithoutIsStrInference([d_2714_v127_, d_2714_v127_])
            d_2715_v128_ = _dafny.SeqWithoutIsStrInference([d_2714_v127_])
            (globalState).f4 = (self).f15
        r0 = not (((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))]) == ((d_2552_v1_)[default__.safeIndex(209, (d_2552_v1_).length(0))])) or ((d_2551_v0_) < (d_2551_v0_))
        return r0


class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def m1(self, p0, p1, globalState):
        r0: D0 = D0.default()()
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        d_2716_v0_: bool
        d_2716_v0_ = False
        if d_2716_v0_:
            d_2717_v1_: _dafny.Seq
            d_2717_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hlnnql"))
            d_2718_v2_: int
            d_2718_v2_ = -982
            d_2719_v4_: D0
            def iife188_():
                coll68_ = _dafny.Map()
                compr_68_: int
                for compr_68_ in _dafny.IntegerRange(603, 114):
                    d_2720_v3_: int = compr_68_
                    if ((603) <= (d_2720_v3_)) and ((d_2720_v3_) < (114)):
                        coll68_[(d_2720_v3_) + ((_dafny.MultiSet([d_2716_v0_, d_2716_v0_])).cardinality)] = d_2716_v0_
                return _dafny.Map(coll68_)
            d_2719_v4_ = D0_DC2(default__.safeModuloInt(len(d_2717_v1_), d_2718_v2_), (D0_DC2(len(iife188_()
), d_2716_v0_)).cf2)
            source40_ = d_2719_v4_
            if source40_.is_DC1:
                d_2721_v5_: _dafny.Map
                d_2721_v5_ = _dafny.Map({d_2716_v0_: (p1)[default__.safeIndex(d_2718_v2_, len(p1))]})
                d_2722_v6_: _dafny.MultiSet
                d_2722_v6_ = _dafny.MultiSet([len(d_2721_v5_)])
                d_2722_v6_ = d_2722_v6_
                d_2723_v7_: _dafny.Seq
                d_2724_v8_: bool
                d_2725_v9_: int
                d_2726_v10_: bool
                out131_: _dafny.Seq
                out132_: bool
                out133_: int
                out134_: bool
                out131_, out132_, out133_, out134_ = default__.m0((d_2719_v4_).cf2, d_2719_v4_, default__.fm2(d_2716_v0_, globalState), globalState)
                d_2723_v7_ = out131_
                d_2724_v8_ = out132_
                d_2725_v9_ = out133_
                d_2726_v10_ = out134_
                (globalState).f4 = default__.fm2(not(d_2716_v0_), globalState)
                d_2725_v9_ = -206
            elif source40_.is_DC2:
                d_2727___mcc_h0_ = source40_.cf1
                d_2728___mcc_h1_ = source40_.cf2
                d_2729_cf2_ = d_2728___mcc_h1_
                d_2730_cf1_ = d_2727___mcc_h0_
                d_2731_v11_: _dafny.Array
                nw451_ = _dafny.Array(int(0), 8)
                d_2731_v11_ = nw451_
                index438_ = default__.safeIndex(677, (d_2731_v11_).length(0))
                (d_2731_v11_)[index438_] = d_2730_cf1_
                d_2732_v12_: _dafny.MultiSet
                d_2732_v12_ = _dafny.MultiSet([d_2719_v4_])
                index439_ = default__.safeIndex(677, (d_2731_v11_).length(0))
                (d_2731_v11_)[index439_] = ((d_2732_v12_)[d_2719_v4_] if (d_2719_v4_) in (d_2732_v12_) else 863)
                d_2733_v13_: _dafny.Array
                nw452_ = _dafny.Array(None, 1)
                nw452_[int(0)] = D0_DC2((0) - ((0) - (d_2730_cf1_)), d_2716_v0_)
                d_2733_v13_ = nw452_
                index440_ = default__.safeIndex(496, (d_2733_v13_).length(0))
                (d_2733_v13_)[index440_] = D0_DC2(-257, d_2716_v0_)
                d_2734_v14_: _dafny.Map
                d_2734_v14_ = _dafny.Map({p0: d_2718_v2_})
                d_2735_v15_: _dafny.Map
                d_2735_v15_ = _dafny.Map({not(d_2729_cf2_): ((d_2734_v14_)[p0] if (p0) in (d_2734_v14_) else (d_2731_v11_)[default__.safeIndex(677, (d_2731_v11_).length(0))])})
                d_2736_v16_: _dafny.Seq
                d_2736_v16_ = _dafny.SeqWithoutIsStrInference([d_2735_v15_, d_2735_v15_, d_2735_v15_, d_2735_v15_])
                index441_ = default__.safeIndex(496, (d_2733_v13_).length(0))
                (d_2733_v13_)[index441_] = default__.fm3(d_2736_v16_, globalState)
                d_2717_v1_ = _dafny.SeqWithoutIsStrInference([p0 for d_2737_i0_ in range(default__.abs(504))])
                d_2730_cf1_ = (d_2731_v11_)[default__.safeIndex(677, (d_2731_v11_).length(0))]
            elif True:
                d_2738___mcc_h2_ = source40_.cf0
                d_2739_cf0_ = d_2738___mcc_h2_
                d_2718_v2_ = d_2718_v2_
                d_2718_v2_ = d_2718_v2_
                d_2740_v17_: _dafny.Set
                d_2740_v17_ = _dafny.Set({d_2716_v0_, d_2716_v0_, d_2716_v0_})
                d_2741_v18_: _dafny.Array
                nw453_ = _dafny.Array(int(0), 20)
                d_2741_v18_ = nw453_
                d_2742_v19_: _dafny.Seq
                d_2742_v19_ = _dafny.SeqWithoutIsStrInference([True, d_2716_v0_])
                d_2743_v20_: D0
                d_2743_v20_ = D0_DC0(d_2739_cf0_)
                d_2744_v21_: _dafny.Seq
                d_2744_v21_ = _dafny.SeqWithoutIsStrInference([d_2743_v20_])
                d_2745_v22_: _dafny.Array
                nw454_ = _dafny.Array(None, 21)
                nw454_[int(0)] = len(_dafny.Map({d_2740_v17_: len(_dafny.SeqWithoutIsStrInference([d_2716_v0_, False, d_2716_v0_, d_2716_v0_, d_2739_cf0_]))}))
                nw454_[int(1)] = d_2718_v2_
                nw454_[int(2)] = ((_dafny.MultiSet([d_2716_v0_])).cardinality) + (d_2718_v2_)
                nw454_[int(3)] = d_2718_v2_
                nw454_[int(4)] = d_2718_v2_
                nw454_[int(5)] = d_2718_v2_
                nw454_[int(6)] = len((d_2717_v1_) + ((_dafny.SeqWithoutIsStrInference([p0 for d_2746_i1_ in range(default__.abs(516))])).set(default__.safeIndex(d_2718_v2_, len(_dafny.SeqWithoutIsStrInference([p0 for d_2747_i1_ in range(default__.abs(516))]))), p0)))
                nw454_[int(7)] = ((0) - (d_2718_v2_)) * (d_2718_v2_)
                nw454_[int(8)] = d_2718_v2_
                nw454_[int(9)] = d_2718_v2_
                nw454_[int(10)] = d_2718_v2_
                nw454_[int(11)] = default__.safeModuloInt(680, len(_dafny.Map({d_2716_v0_: d_2741_v18_})))
                nw454_[int(12)] = d_2718_v2_
                nw454_[int(13)] = 555
                nw454_[int(14)] = (_dafny.MultiSet(d_2742_v19_)).cardinality
                nw454_[int(15)] = -863
                nw454_[int(16)] = len(_dafny.SeqWithoutIsStrInference([True, True]))
                nw454_[int(17)] = ((0) - (d_2718_v2_)) * (len(d_2717_v1_))
                nw454_[int(18)] = len((d_2744_v21_).set(default__.safeIndex(d_2718_v2_, len(d_2744_v21_)), default__.fm4(p0, -652, globalState)))
                nw454_[int(19)] = d_2718_v2_
                nw454_[int(20)] = d_2718_v2_
                d_2745_v22_ = nw454_
                index442_ = default__.safeIndex(593, (d_2745_v22_).length(0))
                (d_2745_v22_)[index442_] = d_2718_v2_
                index443_ = default__.safeIndex(593, (d_2745_v22_).length(0))
                (d_2745_v22_)[index443_] = (d_2718_v2_ if d_2716_v0_ else d_2718_v2_)
                d_2748_v23_: _dafny.Map
                d_2748_v23_ = _dafny.Map({d_2739_cf0_: d_2739_cf0_})
                d_2749_v24_: _dafny.Map
                d_2749_v24_ = _dafny.Map({True: d_2748_v23_})
                d_2750_v25_: _dafny.MultiSet
                d_2750_v25_ = _dafny.MultiSet([(d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))]])
                d_2751_v26_: _dafny.Map
                d_2751_v26_ = _dafny.Map({d_2739_cf0_: d_2750_v25_})
                d_2752_v27_: _dafny.Array
                nw455_ = _dafny.Array(None, 26)
                nw455_[int(0)] = d_2750_v25_
                nw455_[int(1)] = d_2750_v25_
                nw455_[int(2)] = d_2750_v25_
                nw455_[int(3)] = d_2750_v25_
                nw455_[int(4)] = d_2750_v25_
                nw455_[int(5)] = d_2750_v25_
                nw455_[int(6)] = d_2750_v25_
                nw455_[int(7)] = d_2750_v25_
                nw455_[int(8)] = _dafny.MultiSet([len(d_2748_v23_), (d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))]])
                nw455_[int(9)] = d_2750_v25_
                nw455_[int(10)] = d_2750_v25_
                nw455_[int(11)] = ((d_2751_v26_)[not(d_2739_cf0_)] if (not(d_2739_cf0_)) in (d_2751_v26_) else default__.fm28(d_2717_v1_, d_2716_v0_, d_2718_v2_, globalState))
                nw455_[int(12)] = d_2750_v25_
                nw455_[int(13)] = d_2750_v25_
                nw455_[int(14)] = d_2750_v25_
                nw455_[int(15)] = d_2750_v25_
                nw455_[int(16)] = _dafny.MultiSet(p1)
                nw455_[int(17)] = (d_2750_v25_).set((d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))], default__.abs((d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))]))
                nw455_[int(18)] = d_2750_v25_
                nw455_[int(19)] = d_2750_v25_
                nw455_[int(20)] = d_2750_v25_
                nw455_[int(21)] = d_2750_v25_
                nw455_[int(22)] = _dafny.MultiSet([(d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))], d_2718_v2_, (d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))], (d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))]])
                nw455_[int(23)] = d_2750_v25_
                nw455_[int(24)] = _dafny.MultiSet([(d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))]])
                nw455_[int(25)] = d_2750_v25_
                d_2752_v27_ = nw455_
                d_2753_v28_: C6
                nw456_ = C6()
                nw456_.ctor__(p1, not(d_2739_cf0_), (d_2749_v24_) | (d_2749_v24_), _dafny.SeqWithoutIsStrInference([d_2716_v0_, (d_2742_v19_)[default__.safeIndex((d_2745_v22_)[default__.safeIndex(593, (d_2745_v22_).length(0))], len(d_2742_v19_))], True, d_2716_v0_]), d_2752_v27_)
                d_2753_v28_ = nw456_
            d_2754_v29_: C0
            nw457_ = C0()
            nw457_.ctor__()
            d_2754_v29_ = nw457_
            d_2716_v0_ = (default__.safeDivisionInt(d_2718_v2_, d_2718_v2_)) not in (p1)
            d_2716_v0_ = d_2716_v0_
            d_2755_v30_: D2
            d_2755_v30_ = D2_DC5(len(_dafny.Set({d_2716_v0_})), d_2716_v0_, False)
            d_2756_v31_: _dafny.Set
            d_2756_v31_ = _dafny.Set({D2_DC5(d_2718_v2_, d_2716_v0_, d_2716_v0_), d_2755_v30_, d_2755_v30_})
            d_2757_v32_: _dafny.Map
            d_2757_v32_ = _dafny.Map({d_2718_v2_: d_2716_v0_})
            rhs392_ = d_2718_v2_
            rhs393_ = default__.safeModuloInt(116, len((d_2756_v31_) | (d_2756_v31_)))
            rhs394_ = d_2716_v0_
            rhs395_ = (False) == (((d_2757_v32_)[d_2718_v2_] if (d_2718_v2_) in (d_2757_v32_) else d_2716_v0_))
            rhs396_ = ((d_2718_v2_) * (d_2718_v2_)) == (d_2718_v2_)
            lhs241_ = globalState
            lhs242_ = globalState
            lhs243_ = globalState
            d_2718_v2_ = rhs392_
            d_2718_v2_ = rhs393_
            lhs241_.f4 = rhs394_
            lhs242_.f4 = rhs395_
            lhs243_.f4 = rhs396_
        elif True:
            d_2758_v33_: int
            d_2758_v33_ = -101
            d_2759_v34_: _dafny.MultiSet
            d_2759_v34_ = _dafny.MultiSet([d_2758_v33_])
            d_2760_v35_: _dafny.Seq
            d_2760_v35_ = _dafny.SeqWithoutIsStrInference([d_2716_v0_, d_2716_v0_])
            d_2761_v36_: _dafny.Map
            d_2761_v36_ = _dafny.Map({d_2716_v0_: d_2716_v0_})
            d_2762_v37_: _dafny.Map
            d_2762_v37_ = _dafny.Map({False: d_2761_v36_})
            d_2763_v38_: _dafny.Array
            nw458_ = _dafny.Array(_dafny.MultiSet({}), 8)
            d_2763_v38_ = nw458_
            d_2764_v39_: T4
            nw459_ = C2()
            nw459_.ctor__(_dafny.SeqWithoutIsStrInference([d_2758_v33_]), d_2716_v0_, d_2762_v37_, d_2760_v35_, d_2763_v38_)
            d_2764_v39_ = nw459_
            d_2765_v40_: _dafny.Map
            d_2765_v40_ = _dafny.Map({d_2764_v39_: d_2716_v0_})
            d_2766_v41_: _dafny.Seq
            d_2766_v41_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2716_v0_]), d_2760_v35_])
            d_2767_v42_: _dafny.Array
            nw460_ = _dafny.Array(None, 6)
            nw460_[int(0)] = len(d_2760_v35_)
            nw460_[int(1)] = 120
            nw460_[int(2)] = d_2758_v33_
            nw460_[int(3)] = len(default__.fm45(d_2716_v0_, d_2758_v33_, d_2758_v33_, globalState))
            nw460_[int(4)] = d_2758_v33_
            nw460_[int(5)] = 540
            d_2767_v42_ = nw460_
            d_2768_v43_: _dafny.MultiSet
            d_2768_v43_ = _dafny.MultiSet([d_2767_v42_, d_2767_v42_])
            d_2769_v44_: T2
            nw461_ = C3()
            nw461_.ctor__(True, d_2763_v38_, d_2716_v0_, d_2762_v37_, d_2760_v35_)
            d_2769_v44_ = nw461_
            d_2770_v45_: _dafny.Map
            d_2770_v45_ = _dafny.Map({(0) - (d_2758_v33_): d_2769_v44_})
            d_2771_v46_: _dafny.Set
            d_2771_v46_ = _dafny.Set({d_2758_v33_})
            d_2772_v47_: C1
            nw462_ = C1()
            nw462_.ctor__((d_2769_v44_).f15, (d_2769_v44_).f15, d_2762_v37_, (d_2769_v44_).f14, d_2763_v38_)
            d_2772_v47_ = nw462_
            d_2773_v48_: _dafny.Seq
            d_2773_v48_ = _dafny.SeqWithoutIsStrInference([d_2772_v47_])
            d_2774_v49_: D11
            d_2774_v49_ = D11_DC33(d_2716_v0_, (0) - (len(d_2771_v46_)), d_2716_v0_, (d_2773_v48_)[default__.safeIndex(d_2758_v33_, len(d_2773_v48_))], (d_2772_v47_).f17)
            d_2775_v50_: _dafny.Array
            nw463_ = _dafny.Array(None, 25)
            nw463_[int(0)] = ((d_2759_v34_).cardinality) * (d_2758_v33_)
            nw463_[int(1)] = (default__.fm20(d_2716_v0_, d_2758_v33_, globalState)).cardinality
            nw463_[int(2)] = len(d_2760_v35_)
            nw463_[int(3)] = d_2758_v33_
            nw463_[int(4)] = 964
            nw463_[int(5)] = d_2758_v33_
            nw463_[int(6)] = (0) - ((0) - (d_2758_v33_))
            nw463_[int(7)] = (d_2758_v33_) * (len(d_2765_v40_))
            nw463_[int(8)] = 732
            nw463_[int(9)] = d_2758_v33_
            nw463_[int(10)] = d_2758_v33_
            nw463_[int(11)] = (d_2758_v33_) + (d_2758_v33_)
            nw463_[int(12)] = d_2758_v33_
            nw463_[int(13)] = d_2758_v33_
            nw463_[int(14)] = d_2758_v33_
            nw463_[int(15)] = len((d_2766_v41_)[default__.safeIndex((d_2768_v43_).cardinality, len(d_2766_v41_))])
            nw463_[int(16)] = d_2758_v33_
            nw463_[int(17)] = (p1)[default__.safeIndex(d_2758_v33_, len(p1))]
            nw463_[int(18)] = (d_2758_v33_ if d_2716_v0_ else (0) - (d_2758_v33_))
            nw463_[int(19)] = d_2758_v33_
            nw463_[int(20)] = len((d_2770_v45_ if d_2716_v0_ else _dafny.Map({d_2758_v33_: d_2769_v44_})))
            nw463_[int(21)] = d_2758_v33_
            nw463_[int(22)] = (d_2769_v44_).fm7((0) - (d_2758_v33_), globalState)
            nw463_[int(23)] = default__.safeModuloInt(d_2758_v33_, d_2758_v33_)
            nw463_[int(24)] = (d_2774_v49_).cf49
            d_2775_v50_ = nw463_
            index444_ = default__.safeIndex(479, (d_2767_v42_).length(0))
            (d_2767_v42_)[index444_] = d_2758_v33_
            index445_ = default__.safeIndex(479, (d_2767_v42_).length(0))
            (d_2767_v42_)[index445_] = d_2758_v33_
            d_2776_v51_: C0
            nw464_ = C0()
            nw464_.ctor__()
            d_2776_v51_ = nw464_
            d_2777_v52_: C0
            d_2777_v52_ = d_2776_v51_
            d_2778_v53_: _dafny.Map
            d_2778_v53_ = _dafny.Map({False: d_2777_v52_})
            d_2779_v54_: _dafny.Map
            d_2779_v54_ = _dafny.Map({d_2778_v53_: p0})
            d_2779_v54_ = d_2779_v54_
            if not((p1) <= (p1)):
                d_2780_v55_: _dafny.Map
                d_2780_v55_ = _dafny.Map({(d_2767_v42_)[default__.safeIndex(479, (d_2767_v42_).length(0))]: d_2758_v33_})
                d_2780_v55_ = _dafny.Map({(d_2767_v42_)[default__.safeIndex(479, (d_2767_v42_).length(0))]: (p1)[default__.safeIndex((d_2767_v42_)[default__.safeIndex(479, (d_2767_v42_).length(0))], len(p1))]})
                (d_2772_v47_).f18 = d_2772_v47_.f18
                d_2781_v56_: C5
                nw465_ = C5()
                nw465_.ctor__(not((d_2769_v44_).f15), d_2769_v44_.f12, d_2762_v37_, (d_2769_v44_).f14)
                d_2781_v56_ = nw465_
                d_2782_v57_: _dafny.Map
                d_2782_v57_ = _dafny.Map({(d_2772_v47_).f17: d_2781_v56_})
                d_2783_v58_: _dafny.Map
                d_2783_v58_ = _dafny.Map({d_2775_v50_: d_2775_v50_})
                d_2782_v57_ = (d_2782_v57_).set(((d_2783_v58_).set(d_2767_v42_, d_2775_v50_)) != (d_2783_v58_), d_2781_v56_)
                d_2784_v59_: _dafny.Array
                def lambda152_(d_2785_v46_):
                    def lambda153_(d_2786_i2_):
                        return d_2785_v46_

                    return lambda153_

                init81_ = lambda152_(d_2771_v46_)
                nw466_ = _dafny.Array(None, 2)
                for i0_81_ in range(nw466_.length(0)):
                    nw466_[i0_81_] = init81_(i0_81_)
                d_2784_v59_ = nw466_
                index446_ = default__.safeIndex(405, (d_2784_v59_).length(0))
                (d_2784_v59_)[index446_] = d_2771_v46_
                index447_ = default__.safeIndex(405, (d_2784_v59_).length(0))
                (d_2784_v59_)[index447_] = d_2771_v46_
                d_2787_v60_: _dafny.Array
                def lambda154_(d_2788_v44_):
                    def lambda155_(d_2789_i3_):
                        return (d_2788_v44_).f15

                    return lambda155_

                init82_ = lambda154_(d_2769_v44_)
                nw467_ = _dafny.Array(None, 13)
                for i0_82_ in range(nw467_.length(0)):
                    nw467_[i0_82_] = init82_(i0_82_)
                d_2787_v60_ = nw467_
                d_2790_v61_: _dafny.Seq
                d_2790_v61_ = _dafny.SeqWithoutIsStrInference([d_2787_v60_])
                d_2780_v55_ = (d_2780_v55_).set(len(d_2790_v61_), default__.safeModuloInt(len(d_2771_v46_), (d_2767_v42_)[default__.safeIndex(479, (d_2767_v42_).length(0))]))
            elif True:
                d_2761_v36_ = (d_2761_v36_).set(False, (d_2772_v47_).f17)
                d_2791_v62_: _dafny.Seq
                d_2791_v62_ = _dafny.SeqWithoutIsStrInference([d_2775_v50_])
                d_2791_v62_ = d_2791_v62_
                d_2758_v33_ = (d_2758_v33_) * (d_2758_v33_)
                d_2792_v63_: _dafny.Array
                def lambda156_(d_2793_v47_):
                    def lambda157_(d_2794_i4_):
                        return (d_2793_v47_).f17

                    return lambda157_

                init83_ = lambda156_(d_2772_v47_)
                nw468_ = _dafny.Array(None, 15)
                for i0_83_ in range(nw468_.length(0)):
                    nw468_[i0_83_] = init83_(i0_83_)
                d_2792_v63_ = nw468_
                index448_ = default__.safeIndex(810, (d_2792_v63_).length(0))
                (d_2792_v63_)[index448_] = (d_2769_v44_).f15
                index449_ = default__.safeIndex(810, (d_2792_v63_).length(0))
                (d_2792_v63_)[index449_] = False
                d_2795_v64_: C5
                nw469_ = C5()
                nw469_.ctor__(False, d_2763_v38_, (d_2769_v44_).f13, d_2760_v35_)
                d_2795_v64_ = nw469_
                d_2796_v65_: _dafny.Set
                d_2796_v65_ = _dafny.Set({d_2795_v64_})
                d_2796_v65_ = d_2796_v65_
            index450_ = default__.safeIndex(479, (d_2767_v42_).length(0))
            (d_2767_v42_)[index450_] = (d_2758_v33_) - (-417)
            index451_ = default__.safeIndex(479, (d_2767_v42_).length(0))
            (d_2767_v42_)[index451_] = 320
        r2 = True
        d_2797_v66_: _dafny.Array
        def lambda158_(d_2798_v0_):
            def lambda159_(d_2799_i5_):
                return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djqnvnmv"))), (0) - (-161), 775, len(_dafny.SeqWithoutIsStrInference([True, d_2798_v0_]))])

            return lambda159_

        init84_ = lambda158_(d_2716_v0_)
        nw470_ = _dafny.Array(None, 4)
        for i0_84_ in range(nw470_.length(0)):
            nw470_[i0_84_] = init84_(i0_84_)
        d_2797_v66_ = nw470_
        d_2800_v67_: _dafny.Map
        d_2800_v67_ = _dafny.Map({d_2716_v0_: True})
        d_2801_v68_: _dafny.Map
        d_2801_v68_ = _dafny.Map({d_2716_v0_: d_2800_v67_})
        d_2802_v69_: _dafny.Seq
        d_2802_v69_ = _dafny.SeqWithoutIsStrInference([d_2716_v0_, default__.fm2(d_2716_v0_, globalState), d_2716_v0_, d_2716_v0_, d_2716_v0_])
        d_2803_v70_: C3
        nw471_ = C3()
        nw471_.ctor__(d_2716_v0_, d_2797_v66_, d_2716_v0_, d_2801_v68_, d_2802_v69_)
        d_2803_v70_ = nw471_
        if False:
            d_2804_v71_: int
            d_2804_v71_ = -129
            d_2805_v72_: _dafny.Set
            d_2805_v72_ = _dafny.Set({(d_2803_v70_).f20, (d_2803_v70_).f20, (d_2802_v69_)[default__.safeIndex(d_2804_v71_, len(d_2802_v69_))], (d_2803_v70_).f20})
            d_2806_v73_: _dafny.Map
            d_2806_v73_ = _dafny.Map({len(d_2805_v72_): (d_2803_v70_).f20})
            d_2806_v73_ = (d_2806_v73_).set((0) - (d_2804_v71_), (d_2803_v70_).f20)
            r2 = not((d_2803_v70_).f20)
            (globalState).f4 = d_2716_v0_
            if (d_2803_v70_).f20:
                d_2716_v0_ = not((d_2802_v69_)[default__.safeIndex(d_2804_v71_, len(d_2802_v69_))])
                d_2804_v71_ = 313
                d_2807_v74_: D0
                d_2807_v74_ = D0_DC0((d_2803_v70_).f20)
                d_2808_v75_: _dafny.Seq
                d_2808_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yv"))
                r2 = ((d_2807_v74_).cf0) or ((len(d_2808_v75_)) != (d_2804_v71_))
                d_2809_v76_: _dafny.Map
                d_2809_v76_ = _dafny.Map({d_2803_v70_: d_2806_v73_})
                d_2804_v71_ = (default__.safeModuloInt(d_2804_v71_, (_dafny.MultiSet([(d_2803_v70_).f20, ((d_2806_v73_)[d_2804_v71_] if (d_2804_v71_) in (d_2806_v73_) else (d_2803_v70_).f20)])).cardinality) if d_2716_v0_ else default__.safeDivisionInt(len(((d_2809_v76_)[d_2803_v70_] if (d_2803_v70_) in (d_2809_v76_) else d_2806_v73_)), d_2804_v71_))
                d_2810_v77_: str
                d_2810_v77_ = _dafny.CodePoint('n')
                d_2811_v78_: _dafny.Map
                d_2811_v78_ = _dafny.Map({not((d_2803_v70_).f20): d_2810_v77_})
                d_2812_v79_: _dafny.MultiSet
                d_2812_v79_ = _dafny.MultiSet([len(d_2811_v78_)])
                d_2813_v80_: _dafny.Seq
                d_2813_v80_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(p1)]), d_2812_v79_, _dafny.MultiSet([(0) - (d_2804_v71_), -650]), default__.fm28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), (d_2803_v70_).f20, d_2804_v71_, globalState)])
                rhs397_ = ((d_2812_v79_)[d_2804_v71_] if (d_2804_v71_) in (d_2812_v79_) else default__.safeModuloInt((0) - (len(d_2808_v75_)), d_2804_v71_))
                rhs398_ = d_2804_v71_
                rhs399_ = _dafny.CodePoint('s')
                rhs400_ = d_2804_v71_
                rhs401_ = (d_2813_v80_)[default__.safeIndex(default__.safeModuloInt(d_2804_v71_, d_2804_v71_), len(d_2813_v80_))]
                d_2804_v71_ = rhs397_
                d_2804_v71_ = rhs398_
                d_2810_v77_ = rhs399_
                d_2804_v71_ = rhs400_
                d_2812_v79_ = rhs401_
            elif True:
                d_2814_v81_: _dafny.Seq
                d_2814_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhmrbo"))
                d_2815_v82_: _dafny.Set
                d_2815_v82_ = _dafny.Set({d_2814_v81_})
                d_2816_v83_: T1
                nw472_ = C1()
                nw472_.ctor__((d_2803_v70_).f20, (d_2803_v70_).f20, d_2801_v68_, d_2802_v69_, d_2797_v66_)
                d_2816_v83_ = nw472_
                d_2817_v84_: D3
                d_2817_v84_ = D3_DC10((d_2803_v70_).f20, len(d_2806_v73_), d_2816_v83_)
                d_2818_v85_: _dafny.Map
                d_2818_v85_ = _dafny.Map({_dafny.CodePoint('a'): d_2804_v71_})
                d_2819_v86_: _dafny.Map
                d_2819_v86_ = _dafny.Map({(d_2803_v70_).f20: d_2804_v71_})
                d_2820_v87_: _dafny.MultiSet
                d_2820_v87_ = _dafny.MultiSet([(d_2803_v70_).f20])
                d_2821_v88_: _dafny.Set
                d_2821_v88_ = _dafny.Set({858, ((d_2820_v87_)[d_2716_v0_] if (d_2716_v0_) in (d_2820_v87_) else d_2804_v71_)})
                d_2822_v89_: _dafny.Map
                d_2822_v89_ = _dafny.Map({len(d_2806_v73_): d_2804_v71_})
                d_2823_v90_: _dafny.Array
                nw473_ = _dafny.Array(None, 22)
                nw473_[int(0)] = len(d_2814_v81_)
                nw473_[int(1)] = d_2804_v71_
                nw473_[int(2)] = d_2804_v71_
                nw473_[int(3)] = d_2804_v71_
                nw473_[int(4)] = len((d_2802_v69_) + (default__.fm18(d_2716_v0_, d_2804_v71_, False, d_2804_v71_, globalState)))
                nw473_[int(5)] = len(d_2815_v82_)
                nw473_[int(6)] = (d_2817_v84_).cf15
                nw473_[int(7)] = d_2804_v71_
                nw473_[int(8)] = d_2804_v71_
                nw473_[int(9)] = d_2804_v71_
                nw473_[int(10)] = 318
                nw473_[int(11)] = 263
                nw473_[int(12)] = (d_2804_v71_ if d_2716_v0_ else d_2804_v71_)
                nw473_[int(13)] = ((d_2818_v85_)[p0] if (p0) in (d_2818_v85_) else d_2804_v71_)
                nw473_[int(14)] = d_2804_v71_
                nw473_[int(15)] = len(_dafny.Map({d_2804_v71_: d_2804_v71_}))
                nw473_[int(16)] = d_2804_v71_
                nw473_[int(17)] = (0) - (d_2804_v71_)
                nw473_[int(18)] = len((d_2814_v81_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sa"))))
                nw473_[int(19)] = ((d_2819_v86_)[(d_2803_v70_).f20] if ((d_2803_v70_).f20) in (d_2819_v86_) else len(d_2821_v88_))
                nw473_[int(20)] = len(d_2822_v89_)
                nw473_[int(21)] = (d_2804_v71_) + (d_2804_v71_)
                d_2823_v90_ = nw473_
                d_2824_v91_: C6
                nw474_ = C6()
                nw474_.ctor__(p1, d_2716_v0_, d_2801_v68_, _dafny.SeqWithoutIsStrInference([(d_2803_v70_).f20]), d_2797_v66_)
                d_2824_v91_ = nw474_
                d_2825_v92_: _dafny.Set
                d_2825_v92_ = _dafny.Set({d_2824_v91_, d_2824_v91_, d_2824_v91_})
                index452_ = default__.safeIndex(267, (d_2823_v90_).length(0))
                (d_2823_v90_)[index452_] = default__.safeModuloInt(len(d_2825_v92_), d_2804_v71_)
                d_2826_v93_: T0
                nw475_ = C5()
                nw475_.ctor__(((d_2806_v73_)[(0) - (d_2804_v71_)] if ((0) - (d_2804_v71_)) in (d_2806_v73_) else (d_2803_v70_).f20), d_2797_v66_, d_2801_v68_, (d_2816_v83_).f14)
                d_2826_v93_ = nw475_
                index453_ = default__.safeIndex(267, (d_2823_v90_).length(0))
                rhs402_ = (d_2804_v71_) == (d_2804_v71_)
                rhs403_ = (0) - (d_2804_v71_)
                rhs404_ = d_2826_v93_
                rhs405_ = default__.fm66(globalState)
                rhs406_ = False
                lhs244_ = globalState
                lhs245_ = d_2823_v90_
                lhs246_ = default__.safeIndex(267, (d_2823_v90_).length(0))
                lhs244_.f4 = rhs402_
                lhs245_[lhs246_] = rhs403_
                d_2826_v93_ = rhs404_
                d_2822_v89_ = rhs405_
                r2 = rhs406_
                index454_ = default__.safeIndex(267, (d_2823_v90_).length(0))
                (d_2823_v90_)[index454_] = d_2804_v71_
                d_2827_v94_: D3
                d_2827_v94_ = D3_DC11(D3_DC9((d_2823_v90_)[default__.safeIndex(267, (d_2823_v90_).length(0))], 18))
                d_2828_v95_: _dafny.Map
                d_2828_v95_ = _dafny.Map({default__.fm13(d_2804_v71_, globalState): d_2827_v94_})
                d_2829_v96_: D8
                d_2829_v96_ = D8_DC25(d_2716_v0_, d_2804_v71_, d_2804_v71_, _dafny.Map({not(d_2716_v0_): (d_2823_v90_)[default__.safeIndex(267, (d_2823_v90_).length(0))]}))
                d_2830_v97_: _dafny.Map
                d_2830_v97_ = _dafny.Map({len(d_2814_v81_): (d_2829_v96_).cf40})
                d_2831_v98_: D3
                d_2831_v98_ = D3_DC9(default__.fm0((d_2803_v70_).f20, d_2830_v97_, (d_2803_v70_).f20, globalState), len(d_2814_v81_))
                d_2828_v95_ = (d_2828_v95_).set(d_2805_v72_, D3_DC11(d_2831_v98_))
                d_2819_v86_ = d_2819_v86_
                index455_ = default__.safeIndex(267, (d_2823_v90_).length(0))
                (d_2823_v90_)[index455_] = (d_2804_v71_) - ((d_2823_v90_)[default__.safeIndex(267, (d_2823_v90_).length(0))])
            d_2832_v100_: _dafny.MultiSet
            def iife189_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(430, -595):
                    d_2833_v99_: int = compr_69_
                    if ((430) <= (d_2833_v99_)) and ((d_2833_v99_) < (-595)):
                        coll69_[default__.safeDivisionInt(d_2833_v99_, d_2804_v71_)] = d_2804_v71_
                return _dafny.Map(coll69_)
            d_2832_v100_ = _dafny.MultiSet([iife189_()
            ])
            d_2834_v102_: _dafny.Seq
            d_2834_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcsn"))
            d_2835_v103_: _dafny.Map
            def iife190_():
                coll70_ = _dafny.Map()
                compr_70_: int
                for compr_70_ in _dafny.IntegerRange(430, 674):
                    d_2836_v101_: int = compr_70_
                    if ((430) <= (d_2836_v101_)) and ((d_2836_v101_) < (674)):
                        coll70_[default__.safeDivisionInt(d_2836_v101_, (0) - (d_2804_v71_))] = 887
                return _dafny.Map(coll70_)
            d_2835_v103_ = _dafny.Map({(d_2804_v71_ if d_2716_v0_ else d_2804_v71_): (d_2832_v100_).set(iife190_()
            , default__.abs(len(d_2834_v102_)))})
            d_2837_v104_: _dafny.Map
            d_2837_v104_ = _dafny.Map({d_2804_v71_: d_2804_v71_})
            d_2835_v103_ = (d_2835_v103_).set((0) - (d_2804_v71_), _dafny.MultiSet([d_2837_v104_]))
        elif True:
            d_2838_v105_: int
            d_2838_v105_ = 983
            d_2839_v106_: _dafny.Map
            d_2839_v106_ = _dafny.Map({d_2716_v0_: d_2838_v105_})
            d_2840_v107_: D8
            d_2840_v107_ = D8_DC24(True)
            d_2841_v108_: _dafny.Array
            nw476_ = _dafny.Array(None, 11)
            nw476_[int(0)] = D8_DC24((d_2803_v70_).f20)
            nw476_[int(1)] = default__.fm67(d_2838_v105_, (0) - ((0) - (len(d_2839_v106_))), d_2838_v105_, globalState)
            nw476_[int(2)] = d_2840_v107_
            nw476_[int(3)] = d_2840_v107_
            nw476_[int(4)] = d_2840_v107_
            nw476_[int(5)] = d_2840_v107_
            nw476_[int(6)] = d_2840_v107_
            nw476_[int(7)] = d_2840_v107_
            nw476_[int(8)] = d_2840_v107_
            nw476_[int(9)] = d_2840_v107_
            nw476_[int(10)] = d_2840_v107_
            d_2841_v108_ = nw476_
            index456_ = default__.safeIndex(187, (d_2841_v108_).length(0))
            (d_2841_v108_)[index456_] = d_2840_v107_
            index457_ = default__.safeIndex(187, (d_2841_v108_).length(0))
            (d_2841_v108_)[index457_] = default__.fm67(d_2838_v105_, (0) - (d_2838_v105_), d_2838_v105_, globalState)
            d_2842_v109_: _dafny.Map
            d_2842_v109_ = _dafny.Map({(0) - (d_2838_v105_): _dafny.Map({d_2716_v0_: d_2838_v105_})})
            d_2838_v105_ = default__.safeModuloInt(default__.fm0((d_2803_v70_).f20, d_2842_v109_, (d_2803_v70_).f20, globalState), d_2838_v105_)
            if default__.fm2(d_2716_v0_, globalState):
                d_2843_v110_: _dafny.Set
                d_2843_v110_ = _dafny.Set({True})
                d_2844_v111_: _dafny.MultiSet
                d_2844_v111_ = _dafny.MultiSet([(p1)[default__.safeIndex(d_2838_v105_, len(p1))]])
                rhs407_ = (_dafny.MultiSet(p1)).isdisjoint(d_2844_v111_)
                rhs408_ = d_2843_v110_
                lhs247_ = globalState
                lhs247_.f4 = rhs407_
                d_2843_v110_ = rhs408_
                d_2838_v105_ = d_2838_v105_
                d_2845_v112_: _dafny.Array
                nw477_ = _dafny.Array(None, 21)
                nw477_[int(0)] = (d_2803_v70_).f20
                nw477_[int(1)] = (d_2803_v70_).f20
                nw477_[int(2)] = (d_2803_v70_).f20
                nw477_[int(3)] = (d_2803_v70_).f20
                nw477_[int(4)] = (d_2803_v70_).f20
                nw477_[int(5)] = d_2716_v0_
                nw477_[int(6)] = (d_2803_v70_).f20
                nw477_[int(7)] = (d_2803_v70_).f20
                nw477_[int(8)] = (d_2803_v70_).f20
                nw477_[int(9)] = (d_2803_v70_).f20
                nw477_[int(10)] = (d_2803_v70_).f20
                nw477_[int(11)] = (d_2803_v70_).f20
                nw477_[int(12)] = (d_2803_v70_).f20
                nw477_[int(13)] = (d_2803_v70_).f20
                nw477_[int(14)] = (d_2803_v70_).f20
                nw477_[int(15)] = (d_2803_v70_).f20
                nw477_[int(16)] = (d_2803_v70_).f20
                nw477_[int(17)] = (d_2803_v70_).f20
                nw477_[int(18)] = (d_2803_v70_).f20
                nw477_[int(19)] = (d_2803_v70_).f20
                nw477_[int(20)] = (d_2803_v70_).f20
                d_2845_v112_ = nw477_
                d_2846_v113_: _dafny.Array
                nw478_ = _dafny.Array(int(0), 21)
                d_2846_v113_ = nw478_
                d_2847_v114_: _dafny.Map
                d_2847_v114_ = _dafny.Map({d_2845_v112_: d_2846_v113_})
                d_2848_v115_: C5
                nw479_ = C5()
                nw479_.ctor__((d_2803_v70_).f20, d_2797_v66_, d_2801_v68_, d_2802_v69_)
                d_2848_v115_ = nw479_
                d_2849_v116_: _dafny.Map
                d_2849_v116_ = _dafny.Map({((d_2847_v114_)[d_2845_v112_] if (d_2845_v112_) in (d_2847_v114_) else d_2846_v113_): d_2848_v115_})
                rhs409_ = d_2849_v116_
                rhs410_ = default__.fm2((d_2803_v70_).f20, globalState)
                d_2849_v116_ = rhs409_
                r2 = rhs410_
                index458_ = default__.safeIndex(274, (d_2845_v112_).length(0))
                (d_2845_v112_)[index458_] = (d_2803_v70_).f20
                d_2850_v117_: _dafny.Array
                nw480_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_2850_v117_ = nw480_
                index459_ = default__.safeIndex(919, (d_2850_v117_).length(0))
                (d_2850_v117_)[index459_] = p0
                d_2851_v118_: _dafny.MultiSet
                d_2851_v118_ = _dafny.MultiSet([(d_2803_v70_).f20])
                d_2852_v119_: _dafny.Seq
                d_2852_v119_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnbue"))
                d_2853_v120_: D14
                d_2853_v120_ = D14_DC41(d_2852_v119_, d_2838_v105_)
                index460_ = default__.safeIndex(274, (d_2845_v112_).length(0))
                index461_ = default__.safeIndex(919, (d_2850_v117_).length(0))
                rhs411_ = (d_2838_v105_) < ((((d_2851_v118_)[(d_2803_v70_).f20] if ((d_2803_v70_).f20) in (d_2851_v118_) else d_2838_v105_)) * ((d_2853_v120_).cf65))
                rhs412_ = p0
                lhs248_ = d_2845_v112_
                lhs249_ = default__.safeIndex(274, (d_2845_v112_).length(0))
                lhs250_ = d_2850_v117_
                lhs251_ = default__.safeIndex(919, (d_2850_v117_).length(0))
                lhs248_[lhs249_] = rhs411_
                lhs250_[lhs251_] = rhs412_
                d_2854_v121_: C0
                nw481_ = C0()
                nw481_.ctor__()
                d_2854_v121_ = nw481_
                d_2855_v122_: C0
                d_2855_v122_ = d_2854_v121_
                d_2856_v123_: _dafny.Map
                d_2856_v123_ = _dafny.Map({d_2855_v122_: (p1).set(default__.safeIndex(d_2838_v105_, len(p1)), d_2838_v105_)})
                d_2857_v124_: _dafny.Seq
                d_2857_v124_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_2858_v125_: D5
                d_2858_v125_ = D5_DC15(p1)
                d_2859_v126_: _dafny.Map
                d_2859_v126_ = _dafny.Map({default__.fm2(d_2716_v0_, globalState): p1})
                d_2860_v127_: _dafny.Array
                nw482_ = _dafny.Array(None, 26)
                nw482_[int(0)] = _dafny.SeqWithoutIsStrInference([d_2838_v105_])
                nw482_[int(1)] = p1
                nw482_[int(2)] = (p1).set(default__.safeIndex((_dafny.MultiSet([(d_2803_v70_).f20])).cardinality, len(p1)), d_2838_v105_)
                nw482_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_2838_v105_, d_2838_v105_]) if (d_2848_v115_).f19 else p1)
                nw482_[int(4)] = p1
                nw482_[int(5)] = _dafny.SeqWithoutIsStrInference([d_2838_v105_])
                nw482_[int(6)] = p1
                nw482_[int(7)] = ((p1).set(default__.safeIndex(d_2838_v105_, len(p1)), d_2838_v105_)) + (default__.fm1((d_2803_v70_).f20, (d_2803_v70_).f20, globalState))
                nw482_[int(8)] = (p1).set(default__.safeIndex(d_2838_v105_, len(p1)), d_2838_v105_)
                nw482_[int(9)] = ((d_2856_v123_)[d_2855_v122_] if (d_2855_v122_) in (d_2856_v123_) else (d_2857_v124_)[default__.safeIndex(d_2838_v105_, len(d_2857_v124_))])
                nw482_[int(10)] = p1
                nw482_[int(11)] = (p1) + (p1)
                nw482_[int(12)] = _dafny.SeqWithoutIsStrInference([875 for d_2861_i6_ in range(default__.abs(672))])
                nw482_[int(13)] = p1
                nw482_[int(14)] = _dafny.SeqWithoutIsStrInference([d_2838_v105_, d_2838_v105_, ((d_2851_v118_)[(d_2848_v115_).f19] if ((d_2848_v115_).f19) in (d_2851_v118_) else d_2838_v105_), d_2838_v105_, d_2838_v105_])
                nw482_[int(15)] = p1
                nw482_[int(16)] = p1
                nw482_[int(17)] = _dafny.SeqWithoutIsStrInference([d_2838_v105_ for d_2862_i7_ in range(default__.abs(924))])
                nw482_[int(18)] = default__.fm1((d_2845_v112_)[default__.safeIndex(274, (d_2845_v112_).length(0))], not((d_2803_v70_).f20), globalState)
                nw482_[int(19)] = p1
                nw482_[int(20)] = p1
                nw482_[int(21)] = p1
                nw482_[int(22)] = (d_2858_v125_).cf25
                nw482_[int(23)] = p1
                nw482_[int(24)] = (d_2848_v115_).fm6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "morqadjvx")), globalState)
                nw482_[int(25)] = ((d_2859_v126_)[default__.fm2((d_2803_v70_).f20, globalState)] if (default__.fm2((d_2803_v70_).f20, globalState)) in (d_2859_v126_) else _dafny.SeqWithoutIsStrInference([len(d_2839_v106_)]))
                d_2860_v127_ = nw482_
                index462_ = default__.safeIndex(948, (d_2860_v127_).length(0))
                (d_2860_v127_)[index462_] = (p1) + (p1)
                d_2863_v128_: _dafny.Map
                d_2863_v128_ = _dafny.Map({d_2838_v105_: len(d_2843_v110_)})
                index463_ = default__.safeIndex(948, (d_2860_v127_).length(0))
                (d_2860_v127_)[index463_] = (_dafny.SeqWithoutIsStrInference([d_2838_v105_, len((d_2852_v119_).set(default__.safeIndex(d_2838_v105_, len(d_2852_v119_)), p0)), len(d_2863_v128_)])) + ((p1) + (p1))
            elif True:
                d_2864_v129_: D0
                d_2864_v129_ = D0_DC2(len(p1), (d_2803_v70_).f20)
                d_2865_v130_: _dafny.Array
                nw483_ = _dafny.Array(_dafny.Map({}), 25)
                d_2865_v130_ = nw483_
                d_2866_v131_: _dafny.Array
                d_2866_v131_ = d_2865_v130_
                d_2867_v132_: _dafny.Seq
                d_2868_v133_: bool
                d_2869_v134_: int
                d_2870_v135_: bool
                out135_: _dafny.Seq
                out136_: bool
                out137_: int
                out138_: bool
                out135_, out136_, out137_, out138_ = default__.m0((d_2803_v70_).f20, d_2864_v129_, (d_2838_v105_) <= (len(_dafny.Map({d_2866_v131_: (d_2803_v70_).f20}))), globalState)
                d_2867_v132_ = out135_
                d_2868_v133_ = out136_
                d_2869_v134_ = out137_
                d_2870_v135_ = out138_
                d_2871_v136_: _dafny.Seq
                d_2871_v136_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nn"))
                d_2872_v137_: _dafny.Map
                d_2872_v137_ = _dafny.Map({(d_2871_v136_) != (d_2871_v136_): _dafny.SeqWithoutIsStrInference([not(not((d_2803_v70_).f20))])})
                d_2872_v137_ = (d_2872_v137_).set((d_2803_v70_).f20, d_2802_v69_)
                d_2873_v138_: _dafny.Array
                nw484_ = _dafny.Array(None, 14)
                nw484_[int(0)] = p0
                nw484_[int(1)] = p0
                nw484_[int(2)] = p0
                nw484_[int(3)] = p0
                nw484_[int(4)] = p0
                nw484_[int(5)] = p0
                nw484_[int(6)] = _dafny.CodePoint('f')
                nw484_[int(7)] = p0
                nw484_[int(8)] = p0
                nw484_[int(9)] = p0
                nw484_[int(10)] = p0
                nw484_[int(11)] = p0
                nw484_[int(12)] = _dafny.CodePoint('f')
                nw484_[int(13)] = _dafny.CodePoint('v')
                d_2873_v138_ = nw484_
                index464_ = default__.safeIndex(356, (d_2873_v138_).length(0))
                (d_2873_v138_)[index464_] = p0
                index465_ = default__.safeIndex(356, (d_2873_v138_).length(0))
                (d_2873_v138_)[index465_] = p0
                r2 = d_2716_v0_
                d_2874_v139_: T1
                nw485_ = C1()
                nw485_.ctor__((d_2803_v70_).f20, (d_2803_v70_).f20, d_2801_v68_, d_2802_v69_, d_2797_v66_)
                d_2874_v139_ = nw485_
                nw486_ = C5()
                nw486_.ctor__((d_2869_v134_) >= (d_2869_v134_), d_2874_v139_.f12, (d_2874_v139_).f13, (d_2802_v69_) + (default__.fm15((d_2803_v70_).f20, (p1)[default__.safeIndex(d_2869_v134_, len(p1))], d_2838_v105_, d_2838_v105_, globalState)))
                d_2874_v139_ = nw486_
            d_2875_v140_: C6
            nw487_ = C6()
            nw487_.ctor__((p1) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvbnjdfav"))) for d_2876_i8_ in range(default__.abs(523))])), False, (d_2801_v68_) | (d_2801_v68_), d_2802_v69_, d_2797_v66_)
            d_2875_v140_ = nw487_
            if not(d_2716_v0_):
                d_2877_v141_: D3
                d_2877_v141_ = D3_DC9(d_2838_v105_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2878_i9_ in range(default__.abs(708))])))
                d_2879_v142_: C0
                nw488_ = C0()
                nw488_.ctor__()
                d_2879_v142_ = nw488_
                d_2880_v143_: _dafny.Map
                d_2880_v143_ = _dafny.Map({d_2879_v142_: True})
                d_2881_v144_: D0
                d_2881_v144_ = D0_DC2((d_2877_v141_).cf12, ((d_2880_v143_)[d_2879_v142_] if (d_2879_v142_) in (d_2880_v143_) else d_2716_v0_))
                d_2882_v145_: _dafny.Seq
                d_2883_v146_: bool
                d_2884_v147_: int
                d_2885_v148_: bool
                out139_: _dafny.Seq
                out140_: bool
                out141_: int
                out142_: bool
                out139_, out140_, out141_, out142_ = default__.m0(True, d_2881_v144_, d_2716_v0_, globalState)
                d_2882_v145_ = out139_
                d_2883_v146_ = out140_
                d_2884_v147_ = out141_
                d_2885_v148_ = out142_
                d_2886_v149_: _dafny.Seq
                d_2886_v149_ = _dafny.SeqWithoutIsStrInference([d_2839_v106_])
                d_2887_v150_: _dafny.Seq
                d_2888_v151_: bool
                d_2889_v152_: int
                d_2890_v153_: bool
                out143_: _dafny.Seq
                out144_: bool
                out145_: int
                out146_: bool
                out143_, out144_, out145_, out146_ = default__.m0(d_2885_v148_, default__.fm3(d_2886_v149_, globalState), d_2885_v148_, globalState)
                d_2887_v150_ = out143_
                d_2888_v151_ = out144_
                d_2889_v152_ = out145_
                d_2890_v153_ = out146_
                d_2891_v154_: _dafny.Array
                d_2891_v154_ = d_2797_v66_
                d_2892_v155_: C2
                nw489_ = C2()
                nw489_.ctor__(p1, False, _dafny.Map({d_2883_v146_: d_2800_v67_}), d_2802_v69_, (d_2797_v66_ if (d_2803_v70_).f20 else (d_2891_v154_)))
                d_2892_v155_ = nw489_
                d_2892_v155_ = d_2892_v155_
                d_2885_v148_ = d_2885_v148_
                d_2893_v156_: _dafny.Array
                def lambda160_(d_2894_v148_):
                    def lambda161_(d_2895_i10_):
                        return d_2894_v148_

                    return lambda161_

                init85_ = lambda160_(d_2885_v148_)
                nw490_ = _dafny.Array(None, 5)
                for i0_85_ in range(nw490_.length(0)):
                    nw490_[i0_85_] = init85_(i0_85_)
                d_2893_v156_ = nw490_
                d_2896_v157_: _dafny.MultiSet
                d_2896_v157_ = _dafny.MultiSet([d_2890_v153_, d_2888_v151_])
                index466_ = default__.safeIndex(949, (d_2893_v156_).length(0))
                (d_2893_v156_)[index466_] = (d_2896_v157_).issubset(d_2896_v157_)
                index467_ = default__.safeIndex(949, (d_2893_v156_).length(0))
                (d_2893_v156_)[index467_] = d_2716_v0_
            elif True:
                d_2838_v105_ = d_2838_v105_
                d_2838_v105_ = (0) - (d_2838_v105_)
                d_2838_v105_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ak")))
                d_2802_v69_ = d_2802_v69_
                d_2897_v158_: _dafny.Array
                nw491_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2897_v158_ = nw491_
                d_2898_v159_: _dafny.Array
                nw492_ = _dafny.Array(False, 29)
                d_2898_v159_ = nw492_
                index468_ = default__.safeIndex(160, (d_2897_v158_).length(0))
                (d_2897_v158_)[index468_] = d_2898_v159_
                d_2899_v160_: _dafny.Map
                d_2899_v160_ = _dafny.Map({p0: (d_2803_v70_).f20})
                d_2900_v161_: _dafny.Map
                d_2900_v161_ = _dafny.Map({d_2838_v105_: d_2716_v0_})
                d_2901_v162_: D7
                d_2901_v162_ = D7_DC21(d_2900_v161_, d_2716_v0_, ((d_2800_v67_)[d_2716_v0_] if (d_2716_v0_) in (d_2800_v67_) else True))
                index469_ = default__.safeIndex(160, (d_2897_v158_).length(0))
                nw493_ = _dafny.Array(None, 7)
                nw493_[int(0)] = ((d_2803_v70_).f20 if (d_2803_v70_).f20 else (d_2803_v70_).f20)
                nw493_[int(1)] = (d_2838_v105_) == (d_2838_v105_)
                nw493_[int(2)] = ((d_2899_v160_)[_dafny.CodePoint('o')] if (_dafny.CodePoint('o')) in (d_2899_v160_) else (d_2803_v70_).f20)
                nw493_[int(3)] = (d_2803_v70_).f20
                nw493_[int(4)] = (d_2901_v162_).cf32
                nw493_[int(5)] = not(((d_2900_v161_)[(d_2875_v140_).fm7(d_2838_v105_, globalState)] if ((d_2875_v140_).fm7(d_2838_v105_, globalState)) in (d_2900_v161_) else (d_2803_v70_).f20))
                nw493_[int(6)] = (d_2803_v70_).f20
                (d_2897_v158_)[index469_] = nw493_
        d_2902_v163_: int
        d_2902_v163_ = 519
        d_2903_v164_: _dafny.MultiSet
        d_2903_v164_ = _dafny.MultiSet([d_2902_v163_])
        hi16_ = ((d_2903_v164_)[d_2902_v163_] if (d_2902_v163_) in (d_2903_v164_) else d_2902_v163_)
        for d_2904_i11_ in range(d_2902_v163_, hi16_):
            d_2905_v165_: C1
            nw494_ = C1()
            nw494_.ctor__((d_2803_v70_).f20, True, d_2801_v68_, _dafny.SeqWithoutIsStrInference([(d_2803_v70_).f20, (d_2803_v70_).f20]), d_2797_v66_)
            d_2905_v165_ = nw494_
            d_2906_v166_: _dafny.Array
            nw495_ = _dafny.Array(False, 11)
            d_2906_v166_ = nw495_
            d_2907_v167_: _dafny.Map
            d_2907_v167_ = _dafny.Map({d_2902_v163_: True})
            d_2908_v168_: _dafny.Map
            d_2908_v168_ = _dafny.Map({d_2716_v0_: len(d_2907_v167_)})
            d_2909_v169_: _dafny.Set
            d_2909_v169_ = _dafny.Set({d_2716_v0_})
            d_2910_v170_: _dafny.Map
            d_2910_v170_ = _dafny.Map({(len(d_2908_v168_) if d_2716_v0_ else len(d_2909_v169_)): d_2906_v166_})
            d_2906_v166_ = ((d_2910_v170_)[(d_2902_v163_) * (d_2904_i11_)] if ((d_2902_v163_) * (d_2904_i11_)) in (d_2910_v170_) else d_2906_v166_)
            d_2800_v67_ = (d_2800_v67_).set((_dafny.SeqWithoutIsStrInference([d_2905_v165_.f18])) <= (d_2802_v69_), not(d_2716_v0_))
            d_2911_v171_: _dafny.Map
            d_2911_v171_ = _dafny.Map({(d_2905_v165_).f17: d_2906_v166_})
            d_2911_v171_ = (d_2911_v171_).set(not(not(d_2716_v0_)), d_2906_v166_)
        if d_2716_v0_:
            d_2912_v172_: _dafny.Seq
            d_2912_v172_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sniuh"))
            d_2913_v173_: int
            d_2914_v174_: int
            d_2915_v175_: bool
            d_2916_v176_: _dafny.Map
            out147_: int
            out148_: int
            out149_: bool
            out150_: _dafny.Map
            out147_, out148_, out149_, out150_ = (d_2803_v70_).m9((True) or (d_2716_v0_), False, (d_2912_v172_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmgiy"))), globalState)
            d_2913_v173_ = out147_
            d_2914_v174_ = out148_
            d_2915_v175_ = out149_
            d_2916_v176_ = out150_
            d_2917_v177_: _dafny.Map
            d_2917_v177_ = _dafny.Map({d_2914_v174_: 474})
            d_2917_v177_ = (d_2917_v177_).set(len(_dafny.Set({default__.fm2(d_2716_v0_, globalState)})), d_2902_v163_)
            d_2918_v178_: C4
            nw496_ = C4()
            nw496_.ctor__()
            d_2918_v178_ = nw496_
            (globalState).f4 = (len(d_2802_v69_)) >= (d_2914_v174_)
            d_2919_v179_: D11
            d_2919_v179_ = D11_DC32(default__.safeModuloInt(d_2913_v173_, d_2913_v173_), default__.safeModuloInt(d_2902_v163_, 922))
            d_2919_v179_ = D11_DC32(d_2913_v173_, d_2914_v174_)
        elif True:
            d_2902_v163_ = d_2902_v163_
            d_2920_v180_: _dafny.Map
            d_2920_v180_ = _dafny.Map({d_2902_v163_: d_2902_v163_})
            d_2921_v181_: _dafny.Set
            d_2921_v181_ = _dafny.Set({d_2902_v163_, ((d_2903_v164_)[len(d_2802_v69_)] if (len(d_2802_v69_)) in (d_2903_v164_) else d_2902_v163_), d_2902_v163_})
            d_2922_v182_: _dafny.Seq
            d_2922_v182_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jttuiim"))
            d_2923_v183_: _dafny.Map
            d_2923_v183_ = _dafny.Map({d_2716_v0_: d_2922_v182_})
            d_2920_v180_ = (d_2920_v180_).set((d_2902_v163_ if d_2716_v0_ else d_2902_v163_), (len(d_2921_v181_)) * ((0) - (len(((d_2923_v183_)[(d_2803_v70_).f20] if ((d_2803_v70_).f20) in (d_2923_v183_) else d_2922_v182_)))))
            d_2924_v184_: C2
            nw497_ = C2()
            nw497_.ctor__(_dafny.SeqWithoutIsStrInference([-849]), (d_2803_v70_).f20, d_2801_v68_, d_2802_v69_, d_2797_v66_)
            d_2924_v184_ = nw497_
            d_2925_v185_: _dafny.Set
            d_2925_v185_ = _dafny.Set({d_2924_v184_})
            d_2926_v186_: _dafny.Seq
            d_2926_v186_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-741, d_2902_v163_]), p1, _dafny.SeqWithoutIsStrInference([len(d_2922_v182_)]), p1, (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_2803_v70_).f20]))).cardinality for d_2927_i12_ in range(default__.abs(296))])).set(default__.safeIndex(d_2902_v163_, len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_2803_v70_).f20]))).cardinality for d_2928_i12_ in range(default__.abs(296))]))), len(p1))])
            rhs413_ = default__.safeDivisionInt(d_2902_v163_, (0) - (default__.safeDivisionInt(d_2902_v163_, len(d_2926_v186_))))
            rhs414_ = d_2902_v163_
            rhs415_ = not(not((d_2803_v70_).f20))
            rhs416_ = (d_2925_v185_).intersection(d_2925_v185_)
            d_2902_v163_ = rhs413_
            d_2902_v163_ = rhs414_
            d_2716_v0_ = rhs415_
            d_2925_v185_ = rhs416_
            d_2929_v187_: int
            out151_: int
            out151_ = (d_2924_v184_).m4(d_2902_v163_, (d_2902_v163_) * ((p1)[default__.safeIndex(d_2902_v163_, len(p1))]), d_2922_v182_, default__.fm56(p0, globalState), globalState)
            d_2929_v187_ = out151_
            d_2902_v163_ = (0) - ((0) - (d_2929_v187_))
        d_2930_v188_: D0
        d_2930_v188_ = D0_DC1()
        r0 = d_2930_v188_
        r1 = (d_2800_v67_) | (d_2800_v67_)
        r2 = d_2716_v0_
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: D0 = D0.default()()
        d_2931_v0_: _dafny.Array
        def lambda162_(d_2932_i1_):
            return False

        init86_ = lambda162_
        nw498_ = _dafny.Array(None, 12)
        for i0_86_ in range(nw498_.length(0)):
            nw498_[i0_86_] = init86_(i0_86_)
        d_2931_v0_ = nw498_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2931_v0_).length(0)):
            d_2933_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_2933_i0_)) and ((d_2933_i0_) < ((d_2931_v0_).length(0)))):
                (d_2931_v0_)[(d_2933_i0_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbrqflr"))) != ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2934_i2_ in range(default__.abs(630))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))))
        d_2935_v1_: bool
        d_2935_v1_ = False
        d_2936_v2_: _dafny.Set
        d_2936_v2_ = _dafny.Set({(d_2935_v1_ if d_2935_v1_ else d_2935_v1_), True, d_2935_v1_})
        d_2936_v2_ = (d_2936_v2_).intersection(d_2936_v2_)
        d_2937_v3_: _dafny.Array
        nw499_ = _dafny.Array(int(0), 19)
        d_2937_v3_ = nw499_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2937_v3_).length(0)):
            d_2938_i3_: int = guard_loop_5_
            if (True) and (((0) <= (d_2938_i3_)) and ((d_2938_i3_) < ((d_2937_v3_).length(0)))):
                (d_2937_v3_)[(d_2938_i3_)] = (d_2938_i3_) + (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufueuh")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbaoaupfj")))))
        d_2939_v4_: _dafny.Seq
        d_2939_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nubb"))
        hi17_ = -637
        for d_2940_i4_ in range(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktxfhm")) if d_2935_v1_ else d_2939_v4_)), hi17_):
            rhs417_ = p0
            rhs418_ = default__.fm2(d_2935_v1_, globalState)
            rhs419_ = p0
            rhs420_ = d_2935_v1_
            lhs252_ = globalState
            r1 = rhs417_
            d_2935_v1_ = rhs418_
            r0 = rhs419_
            lhs252_.f4 = rhs420_
            index470_ = default__.safeIndex(178, (d_2937_v3_).length(0))
            (d_2937_v3_)[index470_] = d_2940_i4_
            index471_ = default__.safeIndex(178, (d_2937_v3_).length(0))
            (d_2937_v3_)[index471_] = p0
            r0 = d_2940_i4_
            d_2941_v5_: _dafny.Array
            nw500_ = _dafny.Array(int(0), 23)
            d_2941_v5_ = nw500_
            d_2941_v5_ = d_2941_v5_
        d_2942_i5_: int
        d_2942_i5_ = 0
        with _dafny.label("16"):
            while (d_2935_v1_) or (d_2935_v1_):
                with _dafny.c_label("16"):
                    if (d_2942_i5_) >= (100):
                        raise _dafny.Break("16")
                    d_2942_i5_ = (d_2942_i5_) + (1)
                    r1 = -328
                    if d_2935_v1_:
                        d_2943_v6_: _dafny.Map
                        d_2943_v6_ = _dafny.Map({default__.fm2(d_2935_v1_, globalState): p0})
                        d_2943_v6_ = (d_2943_v6_).set(d_2935_v1_, p0)
                        r1 = p0
                        d_2944_v7_: _dafny.Map
                        d_2944_v7_ = _dafny.Map({d_2935_v1_: (d_2935_v1_ if d_2935_v1_ else d_2935_v1_)})
                        d_2945_v8_: _dafny.Seq
                        d_2945_v8_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_2944_v7_ = (d_2944_v7_).set((p0) in (d_2945_v8_), d_2935_v1_)
                        d_2946_v9_: _dafny.MultiSet
                        d_2946_v9_ = _dafny.MultiSet([d_2935_v1_, not(d_2935_v1_)])
                        d_2935_v1_ = (d_2946_v9_).ispropersubset(_dafny.MultiSet([d_2935_v1_]))
                        d_2947_v10_: _dafny.Array
                        nw501_ = _dafny.Array(_dafny.MultiSet({}), 21)
                        d_2947_v10_ = nw501_
                        d_2948_v11_: _dafny.MultiSet
                        d_2948_v11_ = _dafny.MultiSet([d_2935_v1_, d_2935_v1_, not(d_2935_v1_)])
                        index472_ = default__.safeIndex(359, (d_2947_v10_).length(0))
                        (d_2947_v10_)[index472_] = d_2948_v11_
                        d_2949_v12_: _dafny.Seq
                        d_2949_v12_ = _dafny.SeqWithoutIsStrInference([d_2937_v3_])
                        index473_ = default__.safeIndex(359, (d_2947_v10_).length(0))
                        rhs421_ = not((d_2949_v12_) < (d_2949_v12_))
                        rhs422_ = d_2948_v11_
                        rhs423_ = d_2939_v4_
                        lhs253_ = globalState
                        lhs254_ = d_2947_v10_
                        lhs255_ = default__.safeIndex(359, (d_2947_v10_).length(0))
                        lhs253_.f4 = rhs421_
                        lhs254_[lhs255_] = rhs422_
                        d_2939_v4_ = rhs423_
                    elif True:
                        d_2950_v13_: _dafny.Map
                        d_2950_v13_ = _dafny.Map({d_2935_v1_: p0})
                        d_2951_v14_: _dafny.Map
                        d_2951_v14_ = _dafny.Map({d_2935_v1_: d_2950_v13_})
                        d_2952_v15_: _dafny.Map
                        d_2952_v15_ = _dafny.Map({p0: ((d_2951_v14_)[d_2935_v1_] if (d_2935_v1_) in (d_2951_v14_) else d_2950_v13_)})
                        d_2953_v16_: _dafny.Seq
                        d_2953_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_2935_v1_, d_2952_v15_, d_2935_v1_, globalState)])
                        index474_ = default__.safeIndex(176, (d_2937_v3_).length(0))
                        (d_2937_v3_)[index474_] = ((d_2950_v13_)[False] if (False) in (d_2950_v13_) else len(d_2953_v16_))
                        d_2954_v17_: _dafny.Seq
                        d_2954_v17_ = _dafny.SeqWithoutIsStrInference([True])
                        index475_ = default__.safeIndex(176, (d_2937_v3_).length(0))
                        (d_2937_v3_)[index475_] = (p0) + (len((d_2954_v17_) + (d_2954_v17_)))
                        (globalState).f4 = default__.fm2(not(not(not(d_2935_v1_))), globalState)
                        d_2955_v18_: _dafny.Array
                        def lambda163_(d_2956_i6_):
                            return _dafny.CodePoint('g')

                        init87_ = lambda163_
                        nw502_ = _dafny.Array(None, 21)
                        for i0_87_ in range(nw502_.length(0)):
                            nw502_[i0_87_] = init87_(i0_87_)
                        d_2955_v18_ = nw502_
                        d_2957_v19_: D20
                        d_2957_v19_ = D20_DC52(d_2955_v18_)
                        d_2958_v20_: _dafny.Set
                        d_2958_v20_ = _dafny.Set({d_2955_v18_, d_2955_v18_, (d_2957_v19_).cf76, d_2955_v18_, d_2955_v18_})
                        r1 = len((d_2958_v20_).intersection(d_2958_v20_))
                        (globalState).f4 = ((0) - (p0)) != (default__.safeDivisionInt((0) - (p0), p0))
                        d_2959_v21_: _dafny.Set
                        d_2959_v21_ = _dafny.Set({p0})
                        rhs424_ = not(d_2935_v1_)
                        rhs425_ = d_2935_v1_
                        rhs426_ = (d_2959_v21_).ispropersubset(d_2959_v21_)
                        rhs427_ = ((d_2950_v13_)[d_2935_v1_] if (d_2935_v1_) in (d_2950_v13_) else p0)
                        lhs256_ = globalState
                        lhs257_ = globalState
                        lhs258_ = globalState
                        lhs256_.f4 = rhs424_
                        lhs257_.f4 = rhs425_
                        lhs258_.f4 = rhs426_
                        r1 = rhs427_
                    d_2960_v22_: _dafny.MultiSet
                    d_2960_v22_ = _dafny.MultiSet([((0) - (p0) if d_2935_v1_ else -108)])
                    rhs428_ = (d_2960_v22_).intersection(d_2960_v22_)
                    rhs429_ = (0) - ((0) - (p0))
                    rhs430_ = d_2937_v3_
                    d_2960_v22_ = rhs428_
                    r1 = rhs429_
                    d_2937_v3_ = rhs430_
                    d_2961_v23_: _dafny.Map
                    d_2961_v23_ = _dafny.Map({d_2935_v1_: d_2935_v1_})
                    r1 = len((d_2961_v23_ if d_2935_v1_ else _dafny.Map({d_2935_v1_: d_2935_v1_})))
                    pass
            pass
        d_2962_v24_: _dafny.Seq
        d_2962_v24_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2963_v25_: _dafny.Map
        d_2963_v25_ = _dafny.Map({False: _dafny.Map({True: d_2935_v1_})})
        d_2964_v26_: _dafny.Seq
        d_2964_v26_ = _dafny.SeqWithoutIsStrInference([True, d_2935_v1_])
        d_2965_v27_: _dafny.MultiSet
        d_2965_v27_ = _dafny.MultiSet([(0) - (p0)])
        d_2966_v28_: _dafny.Array
        def lambda164_(d_2967_v26_):
            def lambda165_(d_2968_i8_):
                return _dafny.MultiSet([len((D4_DC12(d_2967_v26_)).cf18)])

            return lambda165_

        init88_ = lambda164_(d_2964_v26_)
        nw503_ = _dafny.Array(None, 20)
        for i0_88_ in range(nw503_.length(0)):
            nw503_[i0_88_] = init88_(i0_88_)
        d_2966_v28_ = nw503_
        d_2969_v29_: T4
        nw504_ = C5()
        nw504_.ctor__(d_2935_v1_, d_2966_v28_, d_2963_v25_, _dafny.SeqWithoutIsStrInference([d_2935_v1_, d_2935_v1_]))
        d_2969_v29_ = nw504_
        d_2970_v30_: _dafny.Map
        d_2970_v30_ = _dafny.Map({d_2969_v29_: d_2965_v27_})
        d_2971_v31_: _dafny.Array
        nw505_ = _dafny.Array(None, 7)
        nw505_[int(0)] = d_2965_v27_
        nw505_[int(1)] = d_2965_v27_
        nw505_[int(2)] = d_2965_v27_
        nw505_[int(3)] = d_2965_v27_
        nw505_[int(4)] = _dafny.MultiSet([p0, (0) - (-973), p0, len(_dafny.Map({len(d_2964_v26_): d_2935_v1_})), len(_dafny.SeqWithoutIsStrInference([d_2935_v1_, d_2935_v1_, d_2935_v1_]))])
        nw505_[int(5)] = ((d_2970_v30_)[d_2969_v29_] if (d_2969_v29_) in (d_2970_v30_) else d_2965_v27_)
        nw505_[int(6)] = d_2965_v27_
        d_2971_v31_ = nw505_
        d_2972_v32_: C2
        nw506_ = C2()
        nw506_.ctor__((d_2962_v24_) + (_dafny.SeqWithoutIsStrInference([p0 for d_2973_i7_ in range(default__.abs(-379))])), d_2935_v1_, d_2963_v25_, (d_2964_v26_).set(default__.safeIndex(p0, len(d_2964_v26_)), False), d_2971_v31_)
        d_2972_v32_ = nw506_
        d_2974_v33_: _dafny.Map
        d_2974_v33_ = _dafny.Map({p0: 739})
        r0 = (206 if d_2935_v1_ else ((d_2974_v33_)[p0] if (p0) in (d_2974_v33_) else p0))
        d_2975_v34_: D8
        d_2975_v34_ = D8_DC26(d_2935_v1_)
        pat_let_tv51_ = p0
        pat_let_tv52_ = p0
        pat_let_tv53_ = p0
        pat_let_tv54_ = p0
        pat_let_tv55_ = p0
        pat_let_tv56_ = globalState
        def lambda166_(source41_):
            if source41_.is_DC24:
                d_2976___mcc_h0_ = source41_.cf36
                d_2977_cf36_ = d_2976___mcc_h0_
                return pat_let_tv51_
            elif source41_.is_DC25:
                d_2978___mcc_h1_ = source41_.cf37
                d_2979___mcc_h2_ = source41_.cf38
                d_2980___mcc_h3_ = source41_.cf39
                d_2981___mcc_h4_ = source41_.cf40
                d_2982_cf40_ = d_2981___mcc_h4_
                d_2983_cf39_ = d_2980___mcc_h3_
                d_2984_cf38_ = d_2979___mcc_h2_
                d_2985_cf37_ = d_2978___mcc_h1_
                return pat_let_tv52_
            elif source41_.is_DC26:
                d_2986___mcc_h5_ = source41_.cf41
                d_2987_cf41_ = d_2986___mcc_h5_
                return (0) - (pat_let_tv53_)
            elif source41_.is_DC23:
                d_2988___mcc_h6_ = source41_.cf35
                d_2989_cf35_ = d_2988___mcc_h6_
                return (0) - (pat_let_tv54_)
            elif True:
                d_2990___mcc_h7_ = source41_.cf42
                d_2991_cf42_ = d_2990___mcc_h7_
                return (0) - (pat_let_tv55_)

        def iife191_(_pat_let60_0):
            def iife192_(d_2992_dt__update__tmp_h0_):
                def iife193_(_pat_let61_0):
                    def iife194_(d_2993_dt__update_hcf41_h0_):
                        return D8_DC26(d_2993_dt__update_hcf41_h0_)
                    return iife194_(_pat_let61_0)
                return iife193_(default__.fm2(True, pat_let_tv56_))
            return iife192_(_pat_let60_0)
        r1 = lambda166_(iife191_(d_2975_v34_))
        d_2994_v35_: D0
        d_2994_v35_ = D0_DC1()
        r2 = d_2994_v35_
        return r0, r1, r2

