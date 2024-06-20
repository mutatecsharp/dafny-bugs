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
    def fm0(globalState):
        return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmfo"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kambpnln")) if not(True) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crnml")))))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([497])).Elements:
                d_0_v0_: int = compr_0_
                if (d_0_v0_) in (_dafny.SeqWithoutIsStrInference([497])):
                    coll0_[(d_0_v0_) - (len(_dafny.SeqWithoutIsStrInference([False])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kee")))
            return _dafny.Map(coll0_)
        return ((_dafny.Map({682: 629})) | (iife0_()
        )) | (_dafny.Map({len(_dafny.Map({True: False})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oehpvis")))}))

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')])

    @staticmethod
    def fm3(p0, globalState):
        if not((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "po")))) != (384)):
            return (_dafny.MultiSet([not(False), (D4_DC13(not(True), _dafny.SeqWithoutIsStrInference([41 for d_1_i0_ in range(default__.abs(-529))]))).cf21])).ispropersubset(_dafny.MultiSet([True, True]))
        elif True:
            return (-543) != (len(_dafny.Map({False: not(True)})))

    @staticmethod
    def fm4(p0, p1, globalState):
        return (_dafny.Map({not(not(True)): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgvltt"))) for d_2_i0_ in range(default__.abs(610))]))})) | ((_dafny.Map({False: 320})) | (_dafny.Map({False: 549})))

    @staticmethod
    def fm5(p0, globalState):
        return (_dafny.MultiSet([not(False), True])).intersection(_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([False, False, True, not((D13_DC34(False, True, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_3_i0_ in range(default__.abs(923))]))]))])).cardinality, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))).cf54), False]) if not(False) else _dafny.SeqWithoutIsStrInference([False, not(False), False]))))

    @staticmethod
    def fm6(p0, p1, globalState):
        return ((_dafny.Set({False, True})) - (_dafny.Set({not((D25_DC63(551, not(True), True, False)).cf99)}))) - ((_dafny.Set({True, False}) if not(not(True)) else _dafny.Set({True})))

    @staticmethod
    def fm7(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([True])

    @staticmethod
    def fm8(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: D9
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([D9_DC27(len(_dafny.Set({_dafny.Map({747: False})})))])).Elements:
                d_6_v0_: D9 = compr_1_
                if (d_6_v0_) in (_dafny.SeqWithoutIsStrInference([D9_DC27(len(_dafny.Set({_dafny.Map({747: False})})))])):
                    coll1_ = coll1_.union(_dafny.Set([d_6_v0_]))
            return _dafny.Set(coll1_)
        return ((_dafny.MultiSet([95, len(_dafny.SeqWithoutIsStrInference([D9_DC26((_dafny.MultiSet([-684, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tx"))) for d_4_i1_ in range(default__.abs(624))]))])).cardinality) for d_5_i0_ in range(default__.abs(83))])), len(_dafny.SeqWithoutIsStrInference([True]))])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csvnyvpqu")))])))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([897, (0) - (len(_dafny.SeqWithoutIsStrInference([True])))]))).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(iife1_()
        )]))])))

    @staticmethod
    def fm9(globalState):
        return ((_dafny.Map({not(False): not(True)})) | (_dafny.Map({False: False}))) | (_dafny.Map({False: not(False)}))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        if not(False):
            return D2_DC6(_dafny.CodePoint('o'), True, True, 502)
        elif True:
            return D2_DC6(_dafny.CodePoint('y'), False, True, 787)

    @staticmethod
    def fm17(p0, globalState):
        return D3_DC11(D3_DC9(168, False))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([389, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkid")))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([609, 108, 299, 356, 454])), len(_dafny.Map({True: len(_dafny.Map({False: 286}))})), len(_dafny.Map({True: 867})), 361, 155]))

    @staticmethod
    def fm22(p0, p1, globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return D3_DC11((D3_DC10(868, len(_dafny.SeqWithoutIsStrInference([True, False]))) if True else D3_DC9(100, True)))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return _dafny.CodePoint('w')

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        if (-383) == (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qldwqoe"))), 769]))):
            if not(True):
                return _dafny.SeqWithoutIsStrInference([-827])
            elif True:
                return _dafny.SeqWithoutIsStrInference([-628])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tcvlq"))): 320}): False}))]))).cardinality for d_7_i0_ in range(default__.abs(-762))])) + (_dafny.SeqWithoutIsStrInference([(0) - (-517) for d_8_i1_ in range(default__.abs(99))]))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return _dafny.Map({D3_DC9(-426, False): -56})

    @staticmethod
    def fm33(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(683, 668):
                d_9_v0_: int = compr_2_
                if ((683) <= (d_9_v0_)) and ((d_9_v0_) < (668)):
                    coll2_ = coll2_.union(_dafny.Set([(d_9_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_10_i0_ in range(default__.abs(202))])))]))
            return _dafny.Set(coll2_)
        return (_dafny.Set({590, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncy")))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "md")))})) | (iife2_()
        )

    @staticmethod
    def fm34(p0, p1, globalState):
        return (D26_DC65(_dafny.MultiSet([_dafny.MultiSet([False, True])]))).cf102

    @staticmethod
    def fm35(p0, globalState):
        return D0_DC1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhndkcyca"))))

    @staticmethod
    def fm36(p0, globalState):
        return D6_DC18((0) - (len(_dafny.Set({-147}))), len((_dafny.SeqWithoutIsStrInference([-719])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True]))]))), True, _dafny.CodePoint('t'), 428)

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return (_dafny.Map({(D9_DC26(38)).cf44: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_11_i0_ in range(default__.abs(-408))])})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_12_i1_ in range(default__.abs(200))])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_13_i2_ in range(default__.abs(754))])})) | (_dafny.Map({591: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apixoryg"))})))

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm39(globalState):
        return D3_DC9(default__.safeModuloInt((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).cardinality, len(_dafny.SeqWithoutIsStrInference([477, len(_dafny.SeqWithoutIsStrInference([543, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymxyj")))]))]))), not(False))

    @staticmethod
    def fm40(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_14_i0_ in range(default__.abs(-591))])).Elements:
                d_15_v0_: str = compr_3_
                if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_14_i0_ in range(default__.abs(-591))])):
                    coll3_[d_15_v0_] = len(_dafny.SeqWithoutIsStrInference([True]))
            return _dafny.Map(coll3_)
        return D3_DC11((D3_DC10(len(iife3_()
), 294) if False else D3_DC10(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({608, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_16_i1_ in range(default__.abs(410))]))}))])), 952)))

    @staticmethod
    def fm41(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: D6
            for compr_4_ in (_dafny.MultiSet([D6_DC20(D6_DC18(707, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njoljke"))), False, _dafny.CodePoint('i'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcpaxxcwa")))))])).Elements:
                d_17_v0_: D6 = compr_4_
                if (d_17_v0_) in (_dafny.MultiSet([D6_DC20(D6_DC18(707, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njoljke"))), False, _dafny.CodePoint('i'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcpaxxcwa")))))])):
                    coll4_[d_17_v0_] = 715
            return _dafny.Map(coll4_)
        return (_dafny.MultiSet([iife4_()
        ])) - (_dafny.MultiSet([_dafny.Map({D6_DC20(D6_DC16(_dafny.Map({929: True}))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfk")))}), _dafny.Map({D6_DC20(D6_DC20(D6_DC20(D6_DC20(D6_DC18(899, 505, True, _dafny.CodePoint('e'), (_dafny.MultiSet([(0) - (-690), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_18_i0_ in range(default__.abs(851))]))])).cardinality))))): 121})]))

    @staticmethod
    def fm42(globalState):
        return _dafny.Map({_dafny.CodePoint('q'): default__.safeModuloInt(32, len(_dafny.Map({346: True})))})

    @staticmethod
    def fm43(globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm44(globalState):
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: _dafny.MultiSet
            for compr_5_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True), False]), _dafny.MultiSet([True]), _dafny.MultiSet([True])])).Elements:
                d_19_v0_: _dafny.MultiSet = compr_5_
                if (d_19_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([not(True), False]), _dafny.MultiSet([True]), _dafny.MultiSet([True])])):
                    coll5_ = coll5_.union(_dafny.Set([d_19_v0_]))
            return _dafny.Set(coll5_)
        return ((_dafny.Set({_dafny.MultiSet([True, not(True), True, False, True])})).intersection(_dafny.Set({_dafny.MultiSet([False]), _dafny.MultiSet([True])}))).intersection(iife5_()
        )

    @staticmethod
    def fm45(p0, p1, globalState):
        source0_ = D18_DC44(_dafny.Map({D13_DC34(not(True), True, 606, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nagwkirpq"))): False}))
        if source0_.is_DC45:
            d_20___mcc_h0_ = source0_.cf75
            d_21___mcc_h1_ = source0_.cf76
            d_22___mcc_h2_ = source0_.cf77
            d_23_cf77_ = d_22___mcc_h2_
            d_24_cf76_ = d_21___mcc_h1_
            d_25_cf75_ = d_20___mcc_h0_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')])]))) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_26_i0_ in range(default__.abs(478))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_27_i1_ in range(default__.abs(549))])]))
        elif source0_.is_DC46:
            return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')])])
        elif source0_.is_DC44:
            d_28___mcc_h3_ = source0_.cf74
            d_29_cf74_ = d_28___mcc_h3_
            return _dafny.MultiSet([(D14_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqrhbqnd")))).cf64, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])])
        elif True:
            d_30___mcc_h4_ = source0_.cf78
            d_31_cf78_ = d_30___mcc_h4_
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')])]))

    @staticmethod
    def fm46(globalState):
        return _dafny.MultiSet([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([685])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([_dafny.CodePoint('j')])).cardinality for d_32_i0_ in range(default__.abs(830))]))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bs"))), len(_dafny.Map({D3_DC11(D3_DC10(323, 384)): True}))]))) - (_dafny.MultiSet([(D19_DC49(325, 780)).cf81, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkkho")))])), _dafny.MultiSet([219, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, True, False]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, not(True), True, True]), _dafny.SeqWithoutIsStrInference([not(not(True))])]))}))])])

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D8_DC24(_dafny.CodePoint('f'), _dafny.Map({True: False})) for d_33_i0_ in range(default__.abs(113))])

    @staticmethod
    def fm48(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([D6_DC18(693, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmkabar"))), False, _dafny.CodePoint('d'), len(_dafny.SeqWithoutIsStrInference([True]))), D6_DC18(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mt")))])), 972, not(True), _dafny.CodePoint('e'), 894)])

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        if False:
            return D6_DC17(True)
        elif True:
            return D6_DC17(True)

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return D0_DC0(len(_dafny.Map({True: 978})))

    @staticmethod
    def fm51(p0, p1, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_34_i0_ in range(default__.abs(-656))]): 805})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxllmt")): -687}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovmiu")): (0) - (len(_dafny.SeqWithoutIsStrInference([True])))}))

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, False, not(False), False, False]), (_dafny.MultiSet([not(False)])).intersection(_dafny.MultiSet([True]))])

    @staticmethod
    def fm53(p0, globalState):
        source1_ = D4_DC13(not(not(True)), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (_dafny.MultiSet([127, 985])).cardinality}))]))
        if source1_.is_DC13:
            d_35___mcc_h0_ = source1_.cf21
            d_36___mcc_h1_ = source1_.cf22
            d_37_cf22_ = d_36___mcc_h1_
            d_38_cf21_ = d_35___mcc_h0_
            def iife6_():
                def iife8_():
                    coll8_ = _dafny.Set()
                    compr_8_: str
                    for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])).Elements:
                        d_41_v1_: str = compr_8_
                        if (d_41_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])):
                            coll8_ = coll8_.union(_dafny.Set([d_41_v1_]))
                    return _dafny.Set(coll8_)
                coll6_ = _dafny.Map()
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: str
                    for compr_7_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])).Elements:
                        d_39_v1_: str = compr_7_
                        if (d_39_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m')])):
                            coll7_ = coll7_.union(_dafny.Set([d_39_v1_]))
                    return _dafny.Set(coll7_)
                compr_6_: D13
                for compr_6_ in (_dafny.MultiSet([D13_DC35(195, (_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({d_38_cf21_: d_38_cf21_})): len(d_37_cf22_)}))])).cardinality, d_38_cf21_, -88, D11_DC30(_dafny.Map({len(iife7_()
): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgeiumb"))}))), D13_DC35(-979, -35, d_38_cf21_, -701, D11_DC30(_dafny.Map({-944: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohcyaelk"))}))), D13_DC35((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_38_cf21_]))).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({d_38_cf21_})): 880})])), d_38_cf21_, len(_dafny.Set({d_38_cf21_})), D11_DC30(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, not(d_38_cf21_)])).cardinality])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddei"))})))])).Elements:
                    d_40_v0_: D13 = compr_6_
                    if (d_40_v0_) in (_dafny.MultiSet([D13_DC35(195, (_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({d_38_cf21_: d_38_cf21_})): len(d_37_cf22_)}))])).cardinality, d_38_cf21_, -88, D11_DC30(_dafny.Map({len(iife8_()
): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cgeiumb"))}))), D13_DC35(-979, -35, d_38_cf21_, -701, D11_DC30(_dafny.Map({-944: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohcyaelk"))}))), D13_DC35((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_38_cf21_]))).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Set({d_38_cf21_})): 880})])), d_38_cf21_, len(_dafny.Set({d_38_cf21_})), D11_DC30(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, not(d_38_cf21_)])).cardinality])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddei"))})))])):
                        coll6_[d_40_v0_] = _dafny.CodePoint('h')
                return _dafny.Map(coll6_)
            return (iife6_()
            ) | (_dafny.Map({D13_DC35(len(_dafny.Map({_dafny.CodePoint('n'): -211})), (_dafny.MultiSet(d_37_cf22_)).cardinality, d_38_cf21_, 833, D11_DC30(_dafny.Map({596: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlorlrnk"))}))): _dafny.CodePoint('t')}))
        elif True:
            d_42___mcc_h2_ = source1_.cf20
            d_43_cf20_ = d_42___mcc_h2_
            def iife9_():
                def iife11_():
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwdhkf"))) for d_44_i0_ in range(default__.abs(413))])).Elements:
                        d_45_v3_: int = compr_11_
                        if (d_45_v3_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwdhkf"))) for d_44_i0_ in range(default__.abs(413))])):
                            coll11_[default__.safeDivisionInt(d_45_v3_, -591)] = True
                    return _dafny.Map(coll11_)
                coll9_ = _dafny.Map()
                def iife10_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwdhkf"))) for d_44_i0_ in range(default__.abs(413))])).Elements:
                        d_45_v3_: int = compr_10_
                        if (d_45_v3_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwdhkf"))) for d_44_i0_ in range(default__.abs(413))])):
                            coll10_[default__.safeDivisionInt(d_45_v3_, -591)] = True
                    return _dafny.Map(coll10_)
                compr_9_: D13
                for compr_9_ in (_dafny.SeqWithoutIsStrInference([D13_DC35(len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference([876, 41])), False, (0) - (len(iife10_()
)), D11_DC30(_dafny.Map({654: _dafny.SeqWithoutIsStrInference([(D8_DC24(_dafny.CodePoint('m'), _dafny.Map({not(False): True}))).cf41 for d_46_i1_ in range(default__.abs(273))])})))])).Elements:
                    d_47_v2_: D13 = compr_9_
                    if (d_47_v2_) in (_dafny.SeqWithoutIsStrInference([D13_DC35(len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference([876, 41])), False, (0) - (len(iife11_()
)), D11_DC30(_dafny.Map({654: _dafny.SeqWithoutIsStrInference([(D8_DC24(_dafny.CodePoint('m'), _dafny.Map({not(False): True}))).cf41 for d_46_i1_ in range(default__.abs(273))])})))])):
                        coll9_[d_47_v2_] = _dafny.CodePoint('e')
                return _dafny.Map(coll9_)
            return (iife9_()
            ) | (_dafny.Map({D13_DC35(694, -984, True, 87, D11_DC30(_dafny.Map({559: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))}))): _dafny.CodePoint('t')}))

    @staticmethod
    def fm54(p0, p1, globalState):
        source2_ = D3_DC10(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))), len(_dafny.Map({_dafny.CodePoint('w'): False})))
        if source2_.is_DC9:
            d_48___mcc_h0_ = source2_.cf15
            d_49___mcc_h1_ = source2_.cf16
            d_50_cf16_ = d_49___mcc_h1_
            d_51_cf15_ = d_48___mcc_h0_
            return D6_DC16(_dafny.Map({-500: d_50_cf16_}))
        elif source2_.is_DC10:
            d_52___mcc_h2_ = source2_.cf17
            d_53___mcc_h3_ = source2_.cf18
            d_54_cf18_ = d_53___mcc_h3_
            d_55_cf17_ = d_52___mcc_h2_
            if not(False):
                return D6_DC16(_dafny.Map({len(_dafny.Map({True: True})): False}))
            elif True:
                return D6_DC16(_dafny.Map({469: True}))
        elif source2_.is_DC8:
            d_56___mcc_h4_ = source2_.cf14
            d_57_cf14_ = d_56___mcc_h4_
            def iife12_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(961, 472):
                    d_58_v0_: int = compr_12_
                    if ((961) <= (d_58_v0_)) and ((d_58_v0_) < (472)):
                        coll12_[(d_58_v0_) - (48)] = True
                return _dafny.Map(coll12_)
            return D6_DC16(iife12_()
)
        elif True:
            d_59___mcc_h5_ = source2_.cf19
            d_60_cf19_ = d_59___mcc_h5_
            def iife13_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in (_dafny.MultiSet([637])).Elements:
                    d_61_v1_: int = compr_13_
                    if (d_61_v1_) in (_dafny.MultiSet([637])):
                        coll13_[default__.safeModuloInt(d_61_v1_, -791)] = (D13_DC34(True, False, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(not(False)))]))).cardinality, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnysmsacj")))).cf53
                return _dafny.Map(coll13_)
            return D6_DC16(iife13_()
)

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return D2_DC5(_dafny.Map({True: len(_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_62_i0_ in range(default__.abs(41))])}))}))

    @staticmethod
    def fm56(p0, globalState):
        return D11_DC31(((_dafny.MultiSet([279, len(_dafny.Map({179: 887}))])) | (_dafny.MultiSet([len(_dafny.Map({860: _dafny.Map({True: len(_dafny.Map({3: (_dafny.MultiSet([702, len(_dafny.Map({True: 902}))])).cardinality}))})}))]))).cardinality)

    @staticmethod
    def fm57(p0, globalState):
        def iife14_():
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(-74, 163):
                    d_63_v0_: int = compr_16_
                    if ((-74) <= (d_63_v0_)) and ((d_63_v0_) < (163)):
                        coll16_[default__.safeModuloInt(d_63_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_64_i0_ in range(default__.abs(-969))])))] = 474
                return _dafny.Map(coll16_)
            coll14_ = _dafny.Set()
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-74, 163):
                    d_63_v0_: int = compr_15_
                    if ((-74) <= (d_63_v0_)) and ((d_63_v0_) < (163)):
                        coll15_[default__.safeModuloInt(d_63_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_64_i0_ in range(default__.abs(-969))])))] = 474
                return _dafny.Map(coll15_)
            compr_14_: _dafny.Map
            for compr_14_ in (_dafny.Map({_dafny.Map({len(_dafny.Map({not(True): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({D8_DC23(len(iife15_()
), (_dafny.MultiSet([len(_dafny.Map({True: 438}))])).cardinality, False): True})), 416])): len(_dafny.Map({True: not(False)}))}))})): False}): not(True)})).keys.Elements:
                d_65_v1_: _dafny.Map = compr_14_
                if (d_65_v1_) in (_dafny.Map({_dafny.Map({len(_dafny.Map({not(True): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({D8_DC23(len(iife16_()
), (_dafny.MultiSet([len(_dafny.Map({True: 438}))])).cardinality, False): True})), 416])): len(_dafny.Map({True: not(False)}))}))})): False}): not(True)})):
                    coll14_ = coll14_.union(_dafny.Set([d_65_v1_]))
            return _dafny.Set(coll14_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(-841, 918):
                d_66_v2_: int = compr_17_
                if ((-841) <= (d_66_v2_)) and ((d_66_v2_) < (918)):
                    coll17_[(d_66_v2_) + (7)] = _dafny.CodePoint('p')
            return _dafny.Map(coll17_)
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(405, 398):
                d_67_v3_: int = compr_18_
                if ((405) <= (d_67_v3_)) and ((d_67_v3_) < (398)):
                    coll18_[default__.safeDivisionInt(d_67_v3_, (0) - (len(_dafny.SeqWithoutIsStrInference([True]))))] = True
            return _dafny.Map(coll18_)
        return ((iife14_()
        ) | (_dafny.Set({_dafny.Map({len((_dafny.SeqWithoutIsStrInference([True]))): not(True)}), _dafny.Map({len(_dafny.Map({True: len(_dafny.Map({len(iife17_()
        ): _dafny.MultiSet([True])}))})): True})}))).intersection(_dafny.Set({iife18_()
        , _dafny.Map({621: True})}))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        return (D16_DC41(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_68_i0_ in range(default__.abs(-156))]))).cf69

    @staticmethod
    def fm59(p0, globalState):
        return D8_DC23((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjlhd")))) - (len(_dafny.SeqWithoutIsStrInference([not(not(True))]))), 360, not(not((_dafny.MultiSet([False, not(not(False))])).issubset(_dafny.MultiSet([False])))))

    @staticmethod
    def fm60(globalState):
        if not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).isdisjoint(_dafny.MultiSet([False]))):
            return _dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('w')})
        elif True:
            return _dafny.Set({_dafny.CodePoint('p')})

    @staticmethod
    def fm61(p0, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shgscmrg"))), 934])])) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, True])), 816]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, False, True, True]))})) for d_69_i0_ in range(default__.abs(274))]), _dafny.SeqWithoutIsStrInference([650])]))) - (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([970, 482]), _dafny.SeqWithoutIsStrInference([550])]))

    @staticmethod
    def fm62(p0, p1, globalState):
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgdwgnu")): True})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mobipcov")): False}))

    @staticmethod
    def fm63(p0, p1, p2, p3, globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_70_i0_ in range(default__.abs(996))])): (_dafny.MultiSet([False])).cardinality})).keys.Elements:
                d_71_v0_: int = compr_19_
                if (d_71_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_70_i0_ in range(default__.abs(996))])): (_dafny.MultiSet([False])).cardinality})):
                    coll19_[default__.safeModuloInt(d_71_v0_, 482)] = True
            return _dafny.Map(coll19_)
        return ((_dafny.Set({_dafny.Map({False: True})})) - (_dafny.Set({_dafny.Map({True: True})}))) - (_dafny.Set({_dafny.Map({False: False}), _dafny.Map({not(True): not((D3_DC9(len(iife19_()
), True)).cf16)})}))

    @staticmethod
    def fm64(p0, p1, p2, globalState):
        source3_ = D4_DC13(False, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpwp"))), -500, (0) - (len((D3_DC8(_dafny.Set({False}))).cf14))]))
        if source3_.is_DC13:
            d_72___mcc_h0_ = source3_.cf21
            d_73___mcc_h1_ = source3_.cf22
            d_74_cf22_ = d_73___mcc_h1_
            d_75_cf21_ = d_72___mcc_h0_
            if d_75_cf21_:
                def iife20_():
                    def iife22_():
                        coll22_ = _dafny.Map()
                        compr_22_: D13
                        for compr_22_ in (_dafny.SeqWithoutIsStrInference([D13_DC34(d_75_cf21_, False, 725, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_76_i1_ in range(default__.abs(961))])) for d_77_i0_ in range(default__.abs(216))])).Elements:
                            d_78_v1_: D13 = compr_22_
                            if (d_78_v1_) in (_dafny.SeqWithoutIsStrInference([D13_DC34(d_75_cf21_, False, 725, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_76_i1_ in range(default__.abs(961))])) for d_77_i0_ in range(default__.abs(216))])):
                                coll22_[d_78_v1_] = -316
                        return _dafny.Map(coll22_)
                    coll20_ = _dafny.Map()
                    def iife21_():
                        coll21_ = _dafny.Map()
                        compr_21_: D13
                        for compr_21_ in (_dafny.SeqWithoutIsStrInference([D13_DC34(d_75_cf21_, False, 725, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_76_i1_ in range(default__.abs(961))])) for d_77_i0_ in range(default__.abs(216))])).Elements:
                            d_78_v1_: D13 = compr_21_
                            if (d_78_v1_) in (_dafny.SeqWithoutIsStrInference([D13_DC34(d_75_cf21_, False, 725, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_76_i1_ in range(default__.abs(961))])) for d_77_i0_ in range(default__.abs(216))])):
                                coll21_[d_78_v1_] = -316
                        return _dafny.Map(coll21_)
                    compr_20_: D13
                    for compr_20_ in (iife21_()
                    ).keys.Elements:
                        d_79_v0_: D13 = compr_20_
                        if (d_79_v0_) in (iife22_()
                        ):
                            coll20_[d_79_v0_] = d_75_cf21_
                    return _dafny.Map(coll20_)
                return iife20_()
                
            elif True:
                return _dafny.Map({D13_DC34(d_75_cf21_, d_75_cf21_, (0) - (-435), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucbh"))): d_75_cf21_})
        elif True:
            d_80___mcc_h2_ = source3_.cf20
            d_81_cf20_ = d_80___mcc_h2_
            if False:
                return _dafny.Map({D13_DC34(False, not(False), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))): 221})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxpbdad"))): False})
            elif True:
                return (D18_DC44(_dafny.Map({D13_DC34(False, not(True), len(_dafny.Map({True: True})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lu"))): False}))).cf74

    @staticmethod
    def fm65(p0, p1, p2, p3, globalState):
        return D4_DC13(True, (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhn"))), -381, len(_dafny.SeqWithoutIsStrInference([False])), 858])) + (_dafny.SeqWithoutIsStrInference([324 for d_82_i0_ in range(default__.abs(-57))])))

    @staticmethod
    def fm66(p0, p1, p2, globalState):
        source4_ = D8_DC23(len((D8_DC24(_dafny.CodePoint('n'), _dafny.Map({False: True}))).cf42), 462, True)
        if source4_.is_DC23:
            d_83___mcc_h0_ = source4_.cf38
            d_84___mcc_h1_ = source4_.cf39
            d_85___mcc_h2_ = source4_.cf40
            d_86_cf40_ = d_85___mcc_h2_
            d_87_cf39_ = d_84___mcc_h1_
            d_88_cf38_ = d_83___mcc_h0_
            return D8_DC24(_dafny.CodePoint('m'), _dafny.Map({True: d_86_cf40_}))
        elif source4_.is_DC24:
            d_89___mcc_h3_ = source4_.cf41
            d_90___mcc_h4_ = source4_.cf42
            d_91_cf42_ = d_90___mcc_h4_
            d_92_cf41_ = d_89___mcc_h3_
            return D8_DC24(d_92_cf41_, d_91_cf42_)
        elif True:
            d_93___mcc_h5_ = source4_.cf37
            d_94_cf37_ = d_93___mcc_h5_
            return D8_DC24(_dafny.CodePoint('p'), _dafny.Map({not(True): True}))

    @staticmethod
    def fm67(p0, globalState):
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(-31, 806):
                d_95_v0_: int = compr_23_
                if ((-31) <= (d_95_v0_)) and ((d_95_v0_) < (806)):
                    coll23_[default__.safeDivisionInt(d_95_v0_, -805)] = D8_DC24(_dafny.CodePoint('h'), _dafny.Map({False: False}))
            return _dafny.Map(coll23_)
        return _dafny.Map({(D27_DC68(iife23_()
)).cf109: (530) > ((0) - (len(_dafny.Map({855: False}))))})

    @staticmethod
    def m0(p0, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: _dafny.Seq = _dafny.Seq({})
        d_96_v0_: _dafny.Array
        nw0_ = _dafny.Array(False, 27)
        d_96_v0_ = nw0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_96_v0_).length(0)):
            d_97_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_97_i0_)) and ((d_97_i0_) < ((d_96_v0_).length(0)))):
                (d_96_v0_)[(d_97_i0_)] = (D1_DC4(True, True)).cf6
        d_98_v1_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 28)
        d_98_v1_ = nw1_
        d_99_v2_: int
        d_99_v2_ = -554
        d_100_v3_: _dafny.MultiSet
        d_100_v3_ = _dafny.MultiSet([(0) - (d_99_v2_)])
        d_101_v4_: _dafny.Map
        d_101_v4_ = _dafny.Map({(d_100_v3_).cardinality: d_98_v1_})
        d_102_v5_: _dafny.Array
        nw2_ = _dafny.Array(None, 4)
        nw2_[int(0)] = d_98_v1_
        nw2_[int(1)] = (d_98_v1_ if default__.fm3(d_99_v2_, globalState) else d_98_v1_)
        nw2_[int(2)] = d_98_v1_
        nw2_[int(3)] = ((d_101_v4_)[d_99_v2_] if (d_99_v2_) in (d_101_v4_) else d_98_v1_)
        d_102_v5_ = nw2_
        d_102_v5_ = d_102_v5_
        d_103_v6_: _dafny.Array
        nw3_ = _dafny.Array(None, 2)
        nw3_[int(0)] = d_96_v0_
        nw3_[int(1)] = d_96_v0_
        d_103_v6_ = nw3_
        d_104_v7_: bool
        d_104_v7_ = True
        index0_ = default__.safeIndex(151, (d_96_v0_).length(0))
        (d_96_v0_)[index0_] = d_104_v7_
        d_105_v8_: str
        d_105_v8_ = _dafny.CodePoint('r')
        d_106_v9_: D2
        d_106_v9_ = D2_DC6(d_105_v8_, d_104_v7_, d_104_v7_, d_99_v2_)
        d_107_v10_: _dafny.Set
        d_107_v10_ = _dafny.Set({d_99_v2_})
        d_108_v11_: _dafny.Seq
        d_108_v11_ = _dafny.SeqWithoutIsStrInference([d_104_v7_, d_104_v7_])
        pat_let_tv0_ = d_104_v7_
        pat_let_tv1_ = d_105_v8_
        pat_let_tv2_ = d_104_v7_
        pat_let_tv3_ = d_105_v8_
        pat_let_tv4_ = d_104_v7_
        pat_let_tv5_ = d_99_v2_
        pat_let_tv6_ = d_105_v8_
        d_109_v12_: _dafny.Array
        nw4_ = _dafny.Array(None, 26)
        nw4_[int(0)] = d_106_v9_
        nw4_[int(1)] = d_106_v9_
        def iife24_(_pat_let0_0):
            def iife25_(d_110_dt__update__tmp_h0_):
                def iife26_(_pat_let1_0):
                    def iife27_(d_111_dt__update_hcf10_h0_):
                        def iife28_(_pat_let2_0):
                            def iife29_(d_112_dt__update_hcf9_h0_):
                                def iife30_(_pat_let3_0):
                                    def iife31_(d_113_dt__update_hcf11_h0_):
                                        return D2_DC6(d_112_dt__update_hcf9_h0_, d_111_dt__update_hcf10_h0_, d_113_dt__update_hcf11_h0_, (d_110_dt__update__tmp_h0_).cf12)
                                    return iife31_(_pat_let3_0)
                                return iife30_(pat_let_tv2_)
                            return iife29_(_pat_let2_0)
                        return iife28_(pat_let_tv1_)
                    return iife27_(_pat_let1_0)
                return iife26_(pat_let_tv0_)
            return iife25_(_pat_let0_0)
        nw4_[int(2)] = iife24_(D2_DC6(d_105_v8_, True, d_104_v7_, len(d_107_v10_)))
        nw4_[int(3)] = d_106_v9_
        nw4_[int(4)] = d_106_v9_
        nw4_[int(5)] = D2_DC6(d_105_v8_, True, False, -709)
        nw4_[int(6)] = d_106_v9_
        nw4_[int(7)] = D2_DC6(d_105_v8_, d_104_v7_, d_104_v7_, len(d_108_v11_))
        nw4_[int(8)] = d_106_v9_
        def iife32_(_pat_let4_0):
            def iife33_(d_114_dt__update__tmp_h1_):
                def iife34_(_pat_let5_0):
                    def iife35_(d_115_dt__update_hcf9_h1_):
                        def iife36_(_pat_let6_0):
                            def iife37_(d_116_dt__update_hcf11_h1_):
                                return D2_DC6(d_115_dt__update_hcf9_h1_, (d_114_dt__update__tmp_h1_).cf10, d_116_dt__update_hcf11_h1_, (d_114_dt__update__tmp_h1_).cf12)
                            return iife37_(_pat_let6_0)
                        return iife36_(pat_let_tv4_)
                    return iife35_(_pat_let5_0)
                return iife34_(pat_let_tv3_)
            return iife33_(_pat_let4_0)
        nw4_[int(9)] = iife32_(d_106_v9_)
        nw4_[int(10)] = d_106_v9_
        nw4_[int(11)] = d_106_v9_
        nw4_[int(12)] = d_106_v9_
        nw4_[int(13)] = (d_106_v9_ if False else d_106_v9_)
        nw4_[int(14)] = d_106_v9_
        nw4_[int(15)] = d_106_v9_
        nw4_[int(16)] = d_106_v9_
        nw4_[int(17)] = d_106_v9_
        nw4_[int(18)] = d_106_v9_
        nw4_[int(19)] = d_106_v9_
        nw4_[int(20)] = D2_DC6(d_105_v8_, d_104_v7_, d_104_v7_, d_99_v2_)
        nw4_[int(21)] = D2_DC6(d_105_v8_, d_104_v7_, not(d_104_v7_), d_99_v2_)
        nw4_[int(22)] = d_106_v9_
        nw4_[int(23)] = d_106_v9_
        nw4_[int(24)] = d_106_v9_
        def iife38_(_pat_let7_0):
            def iife39_(d_117_dt__update__tmp_h2_):
                def iife40_(_pat_let8_0):
                    def iife41_(d_118_dt__update_hcf12_h0_):
                        def iife42_(_pat_let9_0):
                            def iife43_(d_119_dt__update_hcf9_h2_):
                                return D2_DC6(d_119_dt__update_hcf9_h2_, (d_117_dt__update__tmp_h2_).cf10, (d_117_dt__update__tmp_h2_).cf11, d_118_dt__update_hcf12_h0_)
                            return iife43_(_pat_let9_0)
                        return iife42_(pat_let_tv6_)
                    return iife41_(_pat_let8_0)
                return iife40_(pat_let_tv5_)
            return iife39_(_pat_let7_0)
        nw4_[int(25)] = iife38_(default__.fm10(d_104_v7_, d_104_v7_, d_104_v7_, globalState))
        d_109_v12_ = nw4_
        index1_ = default__.safeIndex(246, (d_109_v12_).length(0))
        (d_109_v12_)[index1_] = d_106_v9_
        d_120_v13_: _dafny.Map
        d_120_v13_ = _dafny.Map({d_105_v8_: False})
        d_121_v15_: _dafny.Set
        d_121_v15_ = _dafny.Set({_dafny.CodePoint('l')})
        d_122_v16_: _dafny.Map
        d_122_v16_ = _dafny.Map({d_104_v7_: False})
        d_123_v17_: _dafny.Seq
        d_123_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmwdor"))
        d_124_v18_: _dafny.Seq
        d_124_v18_ = _dafny.SeqWithoutIsStrInference([len(d_123_v17_), d_99_v2_, len(d_123_v17_)])
        d_125_v19_: D4
        d_125_v19_ = D4_DC12(d_100_v3_)
        pat_let_tv7_ = d_104_v7_
        index2_ = default__.safeIndex(151, (d_96_v0_).length(0))
        index3_ = default__.safeIndex(246, (d_109_v12_).length(0))
        def iife44_():
            coll24_ = _dafny.Set()
            compr_24_: str
            for compr_24_ in (d_120_v13_).keys.Elements:
                d_126_v14_: str = compr_24_
                if (d_126_v14_) in (d_120_v13_):
                    coll24_ = coll24_.union(_dafny.Set([d_126_v14_]))
            return _dafny.Set(coll24_)
        def lambda0_(source5_):
            if source5_.is_DC13:
                d_127___mcc_h0_ = source5_.cf21
                d_128___mcc_h1_ = source5_.cf22
                d_129_cf22_ = d_128___mcc_h1_
                d_130_cf21_ = d_127___mcc_h0_
                return d_130_cf21_
            elif True:
                d_131___mcc_h2_ = source5_.cf20
                d_132_cf20_ = d_131___mcc_h2_
                return pat_let_tv7_

        rhs0_ = d_103_v6_
        rhs1_ = ((_dafny.SeqWithoutIsStrInference([d_99_v2_, d_99_v2_, len(d_122_v16_), d_99_v2_, len(d_122_v16_)])) == (d_124_v18_) if (d_121_v15_).ispropersubset(iife44_()
        ) else False)
        rhs2_ = lambda0_(d_125_v19_)
        rhs3_ = d_99_v2_
        rhs4_ = D2_DC6((_dafny.CodePoint('x') if d_104_v7_ else d_105_v8_), d_104_v7_, d_104_v7_, d_99_v2_)
        lhs0_ = d_96_v0_
        lhs1_ = default__.safeIndex(151, (d_96_v0_).length(0))
        lhs2_ = d_109_v12_
        lhs3_ = default__.safeIndex(246, (d_109_v12_).length(0))
        d_103_v6_ = rhs0_
        lhs0_[lhs1_] = rhs1_
        r0 = rhs2_
        d_99_v2_ = rhs3_
        lhs2_[lhs3_] = rhs4_
        d_133_v20_: _dafny.Map
        d_133_v20_ = _dafny.Map({d_99_v2_: d_99_v2_})
        def iife45_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in (d_124_v18_).Elements:
                d_134_v21_: int = compr_25_
                if (d_134_v21_) in (d_124_v18_):
                    coll25_[(d_134_v21_) - (d_99_v2_)] = d_99_v2_
            return _dafny.Map(coll25_)
        d_99_v2_ = len(((d_133_v20_) | (_dafny.Map({364: d_99_v2_}))) | (iife45_()
        ))
        d_99_v2_ = -925
        d_135_v22_: _dafny.MultiSet
        d_135_v22_ = _dafny.MultiSet([d_104_v7_])
        d_136_v23_: _dafny.Map
        d_136_v23_ = _dafny.Map({len(d_108_v11_): d_135_v22_})
        d_137_v24_: D1
        d_137_v24_ = D1_DC3(d_104_v7_, default__.safeModuloInt(d_99_v2_, len(d_136_v23_)), d_99_v2_)
        source6_ = d_137_v24_
        if source6_.is_DC3:
            d_138___mcc_h3_ = source6_.cf3
            d_139___mcc_h4_ = source6_.cf4
            d_140___mcc_h5_ = source6_.cf5
            d_141_cf5_ = d_140___mcc_h5_
            d_142_cf4_ = d_139___mcc_h4_
            d_143_cf3_ = d_138___mcc_h3_
            source7_ = D3_DC9(d_141_cf5_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))])
            if source7_.is_DC9:
                d_144___mcc_h9_ = source7_.cf15
                d_145___mcc_h10_ = source7_.cf16
                d_146_cf16_ = d_145___mcc_h10_
                d_147_cf15_ = d_144___mcc_h9_
                d_148_v25_: C0
                nw5_ = C0()
                nw5_.ctor__(d_98_v1_)
                d_148_v25_ = nw5_
                d_141_cf5_ = d_99_v2_
                d_149_v26_: D4
                d_149_v26_ = D4_DC13(d_104_v7_, d_124_v18_)
                d_150_v27_: D4
                d_150_v27_ = D4_DC13(False, (d_149_v26_).cf22)
                d_151_v28_: _dafny.Array
                nw6_ = _dafny.Array(None, 14)
                nw6_[int(0)] = d_150_v27_
                nw6_[int(1)] = D4_DC13(d_146_cf16_, d_124_v18_)
                nw6_[int(2)] = d_150_v27_
                nw6_[int(3)] = default__.fm65(False, d_143_cf3_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], False, globalState)
                nw6_[int(4)] = d_150_v27_
                nw6_[int(5)] = d_150_v27_
                nw6_[int(6)] = d_150_v27_
                nw6_[int(7)] = d_150_v27_
                nw6_[int(8)] = d_150_v27_
                nw6_[int(9)] = d_150_v27_
                nw6_[int(10)] = D4_DC13(d_146_cf16_, d_124_v18_)
                nw6_[int(11)] = D4_DC13((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_124_v18_)
                nw6_[int(12)] = d_149_v26_
                nw6_[int(13)] = d_150_v27_
                d_151_v28_ = nw6_
                index4_ = default__.safeIndex(945, (d_151_v28_).length(0))
                (d_151_v28_)[index4_] = d_149_v26_
                d_152_v29_: _dafny.Map
                d_152_v29_ = _dafny.Map({d_141_cf5_: d_143_cf3_})
                index5_ = default__.safeIndex(945, (d_151_v28_).length(0))
                (d_151_v28_)[index5_] = D4_DC13((len(d_152_v29_)) < (d_141_cf5_), (d_124_v18_) + (_dafny.SeqWithoutIsStrInference([d_147_cf15_])))
                d_153_v30_: _dafny.Map
                d_153_v30_ = _dafny.Map({d_123_v17_: d_104_v7_})
                d_154_v31_: T0
                nw7_ = C5()
                nw7_.ctor__((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], (d_153_v30_) | (d_153_v30_), default__.fm2(d_104_v7_, globalState))
                d_154_v31_ = nw7_
                d_154_v31_ = d_154_v31_
            elif source7_.is_DC10:
                d_155___mcc_h11_ = source7_.cf17
                d_156___mcc_h12_ = source7_.cf18
                d_157_cf18_ = d_156___mcc_h12_
                d_158_cf17_ = d_155___mcc_h11_
                d_159_v32_: _dafny.Seq
                d_159_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]]), d_108_v11_, d_108_v11_])
                d_160_v33_: D16
                d_160_v33_ = D16_DC41(d_159_v32_)
                d_161_v34_: _dafny.Map
                d_161_v34_ = _dafny.Map({default__.fm2(d_143_cf3_, globalState): not(d_104_v7_)})
                d_162_v36_: _dafny.Map
                d_162_v36_ = _dafny.Map({d_123_v17_: d_141_cf5_})
                d_163_v37_: _dafny.Map
                d_163_v37_ = _dafny.Map({(d_124_v18_)[default__.safeIndex(len(_dafny.Map({d_143_cf3_: len(d_123_v17_)})), len(d_124_v18_))]: d_104_v7_})
                d_164_v38_: T1
                nw8_ = C8()
                def iife46_():
                    coll26_ = _dafny.Map()
                    compr_26_: _dafny.Seq
                    for compr_26_ in (d_162_v36_).keys.Elements:
                        d_165_v35_: _dafny.Seq = compr_26_
                        if (d_165_v35_) in (d_162_v36_):
                            coll26_[d_165_v35_] = (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]
                    return _dafny.Map(coll26_)
                nw8_.ctor__((d_159_v32_) + ((d_160_v33_).cf69), (not(d_143_cf3_) if (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))] else (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]), (d_161_v34_) | (iife46_()
                ), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtqsgwsq")), (len(d_163_v37_)) - (d_158_cf17_), d_96_v0_)
                d_164_v38_ = nw8_
                rhs5_ = (d_164_v38_ if (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))] else d_164_v38_)
                rhs6_ = d_143_cf3_
                d_164_v38_ = rhs5_
                d_104_v7_ = rhs6_
                d_99_v2_ = default__.safeModuloInt(d_157_cf18_, d_99_v2_)
                d_166_v39_: C5
                nw9_ = C5()
                nw9_.ctor__(d_104_v7_, d_161_v34_, ((d_164_v38_).f2) + ((d_164_v38_).f2))
                d_166_v39_ = nw9_
                d_166_v39_ = d_166_v39_
                d_143_cf3_ = (d_104_v7_ if (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))] else not (False) or (d_104_v7_))
            elif source7_.is_DC8:
                d_167___mcc_h13_ = source7_.cf14
                d_168_cf14_ = d_167___mcc_h13_
                d_169_v40_: _dafny.Array
                nw10_ = _dafny.Array(D13.default()(), 15)
                d_169_v40_ = nw10_
                d_170_v41_: D22
                d_170_v41_ = D22_DC54(d_169_v40_)
                d_171_v42_: D22
                d_171_v42_ = D22_DC56(d_170_v41_)
                d_172_v43_: _dafny.Map
                d_172_v43_ = _dafny.Map({(d_99_v2_) - (d_141_cf5_): d_171_v42_})
                d_173_v44_: _dafny.Map
                d_173_v44_ = _dafny.Map({d_99_v2_: d_172_v43_})
                rhs7_ = (d_99_v2_) - (len(d_123_v17_))
                rhs8_ = (((d_173_v44_)[(0) - (d_99_v2_)] if ((0) - (d_99_v2_)) in (d_173_v44_) else d_172_v43_)) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_99_v2_ for d_174_i1_ in range(default__.abs(333))])): d_171_v42_}))
                d_99_v2_ = rhs7_
                d_172_v43_ = rhs8_
                index6_ = default__.safeIndex(590, (d_98_v1_).length(0))
                (d_98_v1_)[index6_] = d_142_cf4_
                index7_ = default__.safeIndex(590, (d_98_v1_).length(0))
                (d_98_v1_)[index7_] = d_99_v2_
                d_175_v45_: _dafny.Map
                d_175_v45_ = _dafny.Map({d_104_v7_: d_141_cf5_})
                d_176_v46_: _dafny.Array
                def lambda1_(d_177_i2_):
                    return default__.safeModuloInt(d_177_i2_, -154)

                init0_ = lambda1_
                nw11_ = _dafny.Array(None, 12)
                for i0_0_ in range(nw11_.length(0)):
                    nw11_[i0_0_] = init0_(i0_0_)
                d_176_v46_ = nw11_
                index8_ = default__.safeIndex(590, (d_98_v1_).length(0))
                (d_98_v1_)[index8_] = (d_142_cf4_ if (d_104_v7_ if d_104_v7_ else (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]) else len(((d_175_v45_).set((d_108_v11_)[default__.safeIndex((0) - (d_142_cf4_), len(d_108_v11_))], (0) - ((0) - (len(_dafny.Set({d_176_v46_, d_176_v46_})))))).set((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_142_cf4_)))
                d_178_v47_: _dafny.Map
                d_178_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_105_v8_ for d_179_i3_ in range(default__.abs(997))]): d_143_cf3_})
                d_180_v48_: C3
                nw12_ = C3()
                nw12_.ctor__(821, d_96_v0_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_178_v47_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsqriyfwh")))
                d_180_v48_ = nw12_
                d_181_v49_: _dafny.MultiSet
                d_181_v49_ = _dafny.MultiSet([d_180_v48_])
                d_182_v50_: _dafny.Map
                d_182_v50_ = _dafny.Map({d_142_cf4_: _dafny.SeqWithoutIsStrInference([d_105_v8_ for d_183_i4_ in range(default__.abs(833))])})
                d_184_v51_: _dafny.Seq
                d_184_v51_ = _dafny.SeqWithoutIsStrInference([d_180_v48_, d_180_v48_, d_180_v48_])
                index9_ = default__.safeIndex(151, (d_96_v0_).length(0))
                rhs9_ = (d_168_cf14_).issubset(d_168_cf14_)
                rhs10_ = d_143_cf3_
                rhs11_ = ((d_182_v50_)[(0) - ((0) - (d_142_cf4_))] if ((0) - ((0) - (d_142_cf4_))) in (d_182_v50_) else d_123_v17_)
                rhs12_ = _dafny.MultiSet((d_184_v51_) + (_dafny.SeqWithoutIsStrInference([d_180_v48_])))
                lhs4_ = d_96_v0_
                lhs5_ = default__.safeIndex(151, (d_96_v0_).length(0))
                r0 = rhs9_
                lhs4_[lhs5_] = rhs10_
                d_123_v17_ = rhs11_
                d_181_v49_ = rhs12_
            elif True:
                d_185___mcc_h14_ = source7_.cf19
                d_186_cf19_ = d_185___mcc_h14_
                d_99_v2_ = (0) - ((d_99_v2_) + (default__.safeModuloInt(d_142_cf4_, d_141_cf5_)))
                d_187_v52_: D3
                d_187_v52_ = D3_DC9(d_142_cf4_, d_143_cf3_)
                d_141_cf5_ = (d_187_v52_).cf15
                d_143_cf3_ = d_143_cf3_
                d_141_cf5_ = d_141_cf5_
            if default__.fm3((171) * (len(_dafny.SeqWithoutIsStrInference([d_105_v8_ for d_188_i5_ in range(default__.abs(748))]))), globalState):
                d_104_v7_ = (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]
                index10_ = default__.safeIndex(151, (d_96_v0_).length(0))
                (d_96_v0_)[index10_] = d_143_cf3_
                d_122_v16_ = (d_122_v16_).set(d_104_v7_, False)
                d_189_v53_: _dafny.Array
                def lambda2_(d_190_v15_, d_191_cf5_, d_192_v7_, d_193_v2_, d_194_cf4_, d_195_v17_, d_196_v8_, d_197_v16_, d_198_v0_):
                    def lambda3_(d_199_i6_):
                        return (_dafny.Map({_dafny.Map({(0) - ((D13_DC35(len(d_190_v15_), d_191_cf5_, d_192_v7_, d_193_v2_, D11_DC30(_dafny.Map({d_194_cf4_: d_195_v17_})))).cf58): D8_DC24(d_196_v8_, d_197_v16_)}): (d_198_v0_)[default__.safeIndex(151, (d_198_v0_).length(0))]})) | (_dafny.Map({_dafny.Map({d_191_cf5_: D8_DC24(d_196_v8_, d_197_v16_)}): d_192_v7_}))

                    return lambda3_

                init1_ = lambda2_(d_121_v15_, d_141_cf5_, d_104_v7_, d_99_v2_, d_142_cf4_, d_123_v17_, d_105_v8_, d_122_v16_, d_96_v0_)
                nw13_ = _dafny.Array(None, 29)
                for i0_1_ in range(nw13_.length(0)):
                    nw13_[i0_1_] = init1_(i0_1_)
                d_189_v53_ = nw13_
                d_200_v54_: _dafny.Map
                d_200_v54_ = _dafny.Map({476: default__.fm66(d_105_v8_, d_142_cf4_, len(_dafny.Set({d_105_v8_, d_105_v8_})), globalState)})
                d_201_v55_: _dafny.Map
                d_201_v55_ = _dafny.Map({d_200_v54_: (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]})
                index11_ = default__.safeIndex(515, (d_189_v53_).length(0))
                (d_189_v53_)[index11_] = d_201_v55_
                index12_ = default__.safeIndex(515, (d_189_v53_).length(0))
                (d_189_v53_)[index12_] = (default__.fm67(False, globalState)) | (d_201_v55_)
                r0 = (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]
            elif True:
                d_202_v56_: _dafny.Array
                nw14_ = _dafny.Array(_dafny.Seq({}), 2)
                d_202_v56_ = nw14_
                d_203_v57_: _dafny.Array
                nw15_ = _dafny.Array(None, 6)
                nw15_[int(0)] = _dafny.CodePoint('y')
                nw15_[int(1)] = d_105_v8_
                nw15_[int(2)] = d_105_v8_
                nw15_[int(3)] = d_105_v8_
                nw15_[int(4)] = d_105_v8_
                nw15_[int(5)] = d_105_v8_
                d_203_v57_ = nw15_
                d_204_v58_: _dafny.Seq
                d_204_v58_ = _dafny.SeqWithoutIsStrInference([d_203_v57_])
                d_205_v59_: _dafny.Seq
                d_205_v59_ = d_204_v58_
                index13_ = default__.safeIndex(725, (d_202_v56_).length(0))
                (d_202_v56_)[index13_] = (d_205_v59_)
                d_206_v60_: _dafny.Set
                d_206_v60_ = _dafny.Set({d_104_v7_})
                index14_ = default__.safeIndex(725, (d_202_v56_).length(0))
                rhs13_ = ((len(d_124_v18_)) - (len(d_206_v60_))) * ((157 if default__.fm3(d_99_v2_, globalState) else ((d_135_v22_)[True] if (True) in (d_135_v22_) else d_142_cf4_)))
                rhs14_ = (0) - (d_141_cf5_)
                rhs15_ = d_141_cf5_
                rhs16_ = (d_204_v58_) + (d_204_v58_)
                lhs6_ = d_202_v56_
                lhs7_ = default__.safeIndex(725, (d_202_v56_).length(0))
                d_141_cf5_ = rhs13_
                d_141_cf5_ = rhs14_
                d_99_v2_ = rhs15_
                lhs6_[lhs7_] = rhs16_
                d_123_v17_ = d_123_v17_
                d_207_v61_: _dafny.Map
                d_207_v61_ = _dafny.Map({d_123_v17_: (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))]})
                d_208_v62_: C1
                nw16_ = C1()
                nw16_.ctor__()
                d_208_v62_ = nw16_
                d_209_v63_: _dafny.Seq
                d_209_v63_ = _dafny.SeqWithoutIsStrInference([d_208_v62_, d_208_v62_, d_208_v62_])
                d_210_v64_: C5
                nw17_ = C5()
                nw17_.ctor__(not(d_143_cf3_), (d_207_v61_) | (d_207_v61_), ((d_123_v17_) + ((d_123_v17_).set(default__.safeIndex(len(d_209_v63_), len(d_123_v17_)), d_105_v8_))).set(default__.safeIndex((0) - (d_99_v2_), len((d_123_v17_) + ((d_123_v17_).set(default__.safeIndex(len(d_209_v63_), len(d_123_v17_)), d_105_v8_)))), _dafny.CodePoint('s')))
                d_210_v64_ = nw17_
                d_211_v65_: _dafny.Seq
                d_211_v65_ = _dafny.SeqWithoutIsStrInference([d_210_v64_, d_210_v64_, d_210_v64_])
                d_210_v64_ = (d_210_v64_ if (d_210_v64_).fm11((d_210_v64_).f15, globalState) else (d_211_v65_)[default__.safeIndex(len(default__.fm2(d_143_cf3_, globalState)), len(d_211_v65_))])
                d_143_cf3_ = not((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))])
                d_212_v66_: _dafny.MultiSet
                d_212_v66_ = _dafny.MultiSet([d_105_v8_, d_105_v8_])
                d_213_v67_: T2
                nw18_ = C3()
                nw18_.ctor__((0) - (d_99_v2_), d_96_v0_, (d_142_cf4_) not in (d_124_v18_), default__.fm62(d_99_v2_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], globalState), (d_123_v17_).set(default__.safeIndex(d_99_v2_, len(d_123_v17_)), d_105_v8_))
                d_213_v67_ = nw18_
                rhs17_ = _dafny.MultiSet([d_105_v8_])
                rhs18_ = d_104_v7_
                rhs19_ = (((d_213_v67_).f4 if d_104_v7_ else d_99_v2_)) - (900)
                rhs20_ = d_213_v67_
                rhs21_ = d_99_v2_
                d_212_v66_ = rhs17_
                d_143_cf3_ = rhs18_
                d_99_v2_ = rhs19_
                d_213_v67_ = rhs20_
                d_99_v2_ = rhs21_
            d_214_v68_: D4
            d_214_v68_ = D4_DC13(d_143_cf3_, d_124_v18_)
            source8_ = d_214_v68_
            if source8_.is_DC13:
                d_215___mcc_h15_ = source8_.cf21
                d_216___mcc_h16_ = source8_.cf22
                d_217_cf22_ = d_216___mcc_h16_
                d_218_cf21_ = d_215___mcc_h15_
                d_124_v18_ = d_217_cf22_
                d_99_v2_ = d_142_cf4_
                d_104_v7_ = default__.fm3((d_124_v18_)[default__.safeIndex(d_142_cf4_, len(d_124_v18_))], globalState)
                d_219_v69_: _dafny.Map
                d_219_v69_ = _dafny.Map({len(d_124_v18_): d_217_cf22_})
                d_220_v70_: _dafny.Map
                d_220_v70_ = _dafny.Map({d_218_cf21_: len(d_219_v69_)})
                index15_ = default__.safeIndex(101, (d_98_v1_).length(0))
                (d_98_v1_)[index15_] = ((d_220_v70_)[d_143_cf3_] if (d_143_cf3_) in (d_220_v70_) else d_142_cf4_)
                index16_ = default__.safeIndex(101, (d_98_v1_).length(0))
                (d_98_v1_)[index16_] = len(d_217_cf22_)
            elif True:
                d_221___mcc_h17_ = source8_.cf20
                d_222_cf20_ = d_221___mcc_h17_
                d_99_v2_ = (d_142_cf4_) + (d_99_v2_)
                d_223_v71_: _dafny.Map
                d_223_v71_ = _dafny.Map({d_143_cf3_: d_141_cf5_})
                d_224_v72_: _dafny.Map
                d_224_v72_ = _dafny.Map({d_123_v17_: d_223_v71_})
                d_142_cf4_ = len(d_224_v72_)
                d_225_v73_: _dafny.Map
                d_225_v73_ = _dafny.Map({(0) - (d_141_cf5_): _dafny.CodePoint('v')})
                d_225_v73_ = (d_225_v73_).set(d_99_v2_, d_105_v8_)
                d_226_v74_: _dafny.Map
                d_226_v74_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfgd")): not(d_104_v7_)})
                d_227_v75_: C11
                nw19_ = C11()
                nw19_.ctor__(d_143_cf3_, d_142_cf4_, d_96_v0_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_226_v74_, d_123_v17_)
                d_227_v75_ = nw19_
                d_228_v76_: _dafny.MultiSet
                d_228_v76_ = _dafny.MultiSet([d_227_v75_, d_227_v75_, d_227_v75_, d_227_v75_])
                d_229_v77_: _dafny.Map
                d_229_v77_ = _dafny.Map({len(d_123_v17_): d_228_v76_})
                d_223_v71_ = (d_223_v71_).set((d_228_v76_).issubset(((d_229_v77_)[d_141_cf5_] if (d_141_cf5_) in (d_229_v77_) else d_228_v76_)), d_142_cf4_)
            d_230_v78_: _dafny.Map
            d_230_v78_ = _dafny.Map({_dafny.Map({len(d_107_v10_): len(_dafny.SeqWithoutIsStrInference([False, d_143_cf3_]))}): d_99_v2_})
            d_231_v79_: _dafny.Map
            d_231_v79_ = _dafny.Map({d_123_v17_: d_104_v7_})
            d_232_v80_: T2
            nw20_ = C10()
            nw20_.ctor__((d_100_v3_).cardinality, (0) - (len(d_230_v78_)), d_96_v0_, d_143_cf3_, d_231_v79_, d_123_v17_)
            d_232_v80_ = nw20_
            d_233_v81_: _dafny.Set
            d_233_v81_ = _dafny.Set({d_232_v80_})
            d_234_v82_: _dafny.Map
            d_234_v82_ = _dafny.Map({len(d_233_v81_): d_143_cf3_})
            d_235_v83_: _dafny.Map
            d_235_v83_ = _dafny.Map({d_234_v82_: d_98_v1_})
            d_236_v84_: _dafny.Array
            nw21_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_236_v84_ = nw21_
            index17_ = default__.safeIndex(722, (d_236_v84_).length(0))
            (d_236_v84_)[index17_] = d_105_v8_
            d_237_v86_: _dafny.Seq
            def iife47_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in (d_124_v18_).Elements:
                    d_238_v85_: int = compr_27_
                    if (d_238_v85_) in (d_124_v18_):
                        coll27_ = coll27_.union(_dafny.Set([(d_238_v85_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onh"))))]))
                return _dafny.Set(coll27_)
            d_237_v86_ = _dafny.SeqWithoutIsStrInference([iife47_()
            , (d_107_v10_) | (_dafny.Set({d_141_cf5_, len(d_123_v17_), d_142_cf4_, 969})), d_107_v10_])
            index18_ = default__.safeIndex(722, (d_236_v84_).length(0))
            rhs22_ = (d_237_v86_)[default__.safeIndex(d_99_v2_, len(d_237_v86_))]
            rhs23_ = d_235_v83_
            rhs24_ = d_99_v2_
            rhs25_ = default__.fm38(not(d_143_cf3_), d_232_v80_.f3, d_141_cf5_, globalState)
            rhs26_ = d_232_v80_
            lhs8_ = d_236_v84_
            lhs9_ = default__.safeIndex(722, (d_236_v84_).length(0))
            d_107_v10_ = rhs22_
            d_235_v83_ = rhs23_
            d_142_cf4_ = rhs24_
            lhs8_[lhs9_] = rhs25_
            d_232_v80_ = rhs26_
        elif source6_.is_DC4:
            d_239___mcc_h6_ = source6_.cf6
            d_240___mcc_h7_ = source6_.cf7
            d_241_cf7_ = d_240___mcc_h7_
            d_242_cf6_ = d_239___mcc_h6_
            d_243_v87_: D23
            d_243_v87_ = D23_DC58()
            d_244_v88_: _dafny.MultiSet
            d_244_v88_ = _dafny.MultiSet([d_243_v87_, D23_DC58()])
            d_244_v88_ = d_244_v88_
            d_245_v89_: _dafny.Map
            d_245_v89_ = _dafny.Map({d_123_v17_: d_242_cf6_})
            d_246_v90_: C11
            nw22_ = C11()
            nw22_.ctor__(d_241_cf7_, d_99_v2_, d_96_v0_, default__.fm3(d_99_v2_, globalState), (d_245_v89_) | (d_245_v89_), d_123_v17_)
            d_246_v90_ = nw22_
            d_247_v91_: _dafny.MultiSet
            d_247_v91_ = _dafny.MultiSet([d_133_v20_, d_133_v20_])
            def iife48_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(-393, 105):
                    d_248_v92_: int = compr_28_
                    if ((-393) <= (d_248_v92_)) and ((d_248_v92_) < (105)):
                        coll28_[(d_248_v92_) - (d_99_v2_)] = d_99_v2_
                return _dafny.Map(coll28_)
            d_247_v91_ = (d_247_v91_).intersection((_dafny.MultiSet([iife48_()
            ])).intersection(d_247_v91_))
            d_99_v2_ = d_99_v2_
        elif True:
            d_249___mcc_h8_ = source6_.cf2
            d_250_cf2_ = d_249___mcc_h8_
            d_99_v2_ = d_99_v2_
            d_251_v93_: _dafny.Map
            d_251_v93_ = _dafny.Map({d_99_v2_: d_123_v17_})
            d_252_v94_: D11
            d_252_v94_ = D11_DC30(d_251_v93_)
            d_253_v95_: D13
            d_253_v95_ = D13_DC35(d_99_v2_, d_99_v2_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], len(d_108_v11_), d_252_v94_)
            d_254_v96_: _dafny.Map
            d_254_v96_ = _dafny.Map({(d_250_cf2_).set(default__.safeIndex(d_99_v2_, len(d_250_cf2_)), d_105_v8_): True})
            d_255_v97_: C11
            nw23_ = C11()
            nw23_.ctor__((d_100_v3_).ispropersubset(_dafny.MultiSet([default__.fm0(globalState), (d_253_v95_).cf60])), d_99_v2_, d_96_v0_, False, (d_254_v96_) | (d_254_v96_), d_250_cf2_)
            d_255_v97_ = nw23_
            d_256_v98_: D6
            d_256_v98_ = D6_DC18(d_99_v2_, d_99_v2_, (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_105_v8_, 227)
            pat_let_tv8_ = d_96_v0_
            pat_let_tv9_ = d_96_v0_
            d_257_v99_: _dafny.Map
            def iife49_(_pat_let10_0):
                def iife50_(d_258_dt__update__tmp_h3_):
                    def iife51_(_pat_let11_0):
                        def iife52_(d_259_dt__update_hcf21_h0_):
                            return D4_DC13(d_259_dt__update_hcf21_h0_, (d_258_dt__update__tmp_h3_).cf22)
                        return iife52_(_pat_let11_0)
                    return iife51_((pat_let_tv9_)[default__.safeIndex(151, (pat_let_tv8_).length(0))])
                return iife50_(_pat_let10_0)
            d_257_v99_ = _dafny.Map({iife49_(D4_DC13(not((d_255_v97_).f6), d_124_v18_)): d_250_cf2_})
            d_260_v100_: D4
            d_260_v100_ = D4_DC13((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], d_124_v18_)
            d_261_v101_: _dafny.Array
            nw24_ = _dafny.Array(None, 16)
            nw24_[int(0)] = (d_123_v17_) + ((d_255_v97_).fm14(d_104_v7_, (d_255_v97_).f6, d_105_v8_, d_99_v2_, globalState))
            nw24_[int(1)] = ((d_123_v17_).set(default__.safeIndex(d_99_v2_, len(d_123_v17_)), d_105_v8_)) + (d_123_v17_)
            nw24_[int(2)] = d_123_v17_
            nw24_[int(3)] = ((d_255_v97_).fm14((d_255_v97_).f6, (d_106_v9_).cf10, d_105_v8_, d_99_v2_, globalState)) + (_dafny.SeqWithoutIsStrInference([(d_123_v17_)[default__.safeIndex(d_99_v2_, len(d_123_v17_))], d_105_v8_, d_105_v8_]))
            nw24_[int(4)] = (d_123_v17_) + (d_250_cf2_)
            nw24_[int(5)] = (d_250_cf2_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c')]))
            nw24_[int(6)] = (d_250_cf2_).set(default__.safeIndex(d_99_v2_, len(d_250_cf2_)), (d_256_v98_).cf32)
            nw24_[int(7)] = d_123_v17_
            nw24_[int(8)] = (d_123_v17_) + ((_dafny.SeqWithoutIsStrInference([default__.fm38(True, d_104_v7_, d_99_v2_, globalState)])).set(default__.safeIndex(len(d_123_v17_), len(_dafny.SeqWithoutIsStrInference([default__.fm38(True, d_104_v7_, d_99_v2_, globalState)]))), _dafny.CodePoint('y')))
            nw24_[int(9)] = ((d_257_v99_)[d_260_v100_] if (d_260_v100_) in (d_257_v99_) else d_123_v17_)
            nw24_[int(10)] = _dafny.SeqWithoutIsStrInference([d_105_v8_ for d_262_i7_ in range(default__.abs(448))])
            nw24_[int(11)] = _dafny.SeqWithoutIsStrInference([d_105_v8_])
            nw24_[int(12)] = d_123_v17_
            nw24_[int(13)] = d_250_cf2_
            nw24_[int(14)] = d_123_v17_
            nw24_[int(15)] = d_123_v17_
            d_261_v101_ = nw24_
            index19_ = default__.safeIndex(38, (d_261_v101_).length(0))
            (d_261_v101_)[index19_] = d_123_v17_
            d_263_v103_: _dafny.MultiSet
            d_263_v103_ = _dafny.MultiSet([d_105_v8_])
            index20_ = default__.safeIndex(151, (d_96_v0_).length(0))
            index21_ = default__.safeIndex(38, (d_261_v101_).length(0))
            def iife53_():
                coll29_ = _dafny.Set()
                compr_29_: int
                for compr_29_ in (_dafny.MultiSet([594])).Elements:
                    d_264_v102_: int = compr_29_
                    if (d_264_v102_) in (_dafny.MultiSet([594])):
                        coll29_ = coll29_.union(_dafny.Set([default__.safeDivisionInt(d_264_v102_, -109)]))
                return _dafny.Set(coll29_)
            rhs27_ = (iife53_()
            ).isdisjoint(d_107_v10_)
            rhs28_ = ((0) - ((d_99_v2_) + (d_99_v2_))) - (len(_dafny.SeqWithoutIsStrInference([(d_255_v97_).f6, (d_255_v97_).f6])))
            rhs29_ = (0) - (((d_263_v103_).cardinality if (d_105_v8_) not in (d_250_cf2_) else d_99_v2_))
            rhs30_ = (d_123_v17_) + ((d_250_cf2_) + ((d_255_v97_).fm14((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], (d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))], _dafny.CodePoint('b'), d_99_v2_, globalState)))
            rhs31_ = d_99_v2_
            lhs10_ = d_96_v0_
            lhs11_ = default__.safeIndex(151, (d_96_v0_).length(0))
            lhs12_ = d_261_v101_
            lhs13_ = default__.safeIndex(38, (d_261_v101_).length(0))
            lhs10_[lhs11_] = rhs27_
            d_99_v2_ = rhs28_
            d_99_v2_ = rhs29_
            lhs12_[lhs13_] = rhs30_
            d_99_v2_ = rhs31_
            if (d_255_v97_).f6:
                d_265_v104_: _dafny.Map
                d_265_v104_ = _dafny.Map({(default__.fm22(True, d_99_v2_, globalState) if d_104_v7_ else _dafny.CodePoint('m')): default__.safeModuloInt(d_99_v2_, d_99_v2_)})
                d_265_v104_ = (d_265_v104_).set(d_105_v8_, d_99_v2_)
                rhs32_ = d_99_v2_
                rhs33_ = d_261_v101_
                d_99_v2_ = rhs32_
                d_261_v101_ = rhs33_
                index22_ = default__.safeIndex(313, (d_98_v1_).length(0))
                (d_98_v1_)[index22_] = d_99_v2_
                d_266_v105_: C7
                nw25_ = C7()
                nw25_.ctor__(d_99_v2_, d_96_v0_, False, d_254_v96_, _dafny.SeqWithoutIsStrInference([d_105_v8_ for d_267_i8_ in range(default__.abs(607))]))
                d_266_v105_ = nw25_
                d_268_v106_: _dafny.Map
                d_268_v106_ = _dafny.Map({d_99_v2_: d_104_v7_})
                d_269_v107_: _dafny.Set
                d_269_v107_ = _dafny.Set({d_266_v105_, d_266_v105_, (D25_DC62(d_266_v105_, len(d_268_v106_), (d_255_v97_).fm11((d_255_v97_).f6, globalState))).cf94})
                d_270_v108_: _dafny.Seq
                d_270_v108_ = _dafny.SeqWithoutIsStrInference([d_269_v107_])
                index23_ = default__.safeIndex(313, (d_98_v1_).length(0))
                (d_98_v1_)[index23_] = len(d_270_v108_)
                d_271_v109_: C8
                nw26_ = C8()
                nw26_.ctor__(_dafny.SeqWithoutIsStrInference([d_108_v11_, d_108_v11_, d_108_v11_, d_108_v11_]), (d_255_v97_).f6, d_254_v96_, d_123_v17_, (d_98_v1_)[default__.safeIndex(313, (d_98_v1_).length(0))], d_96_v0_)
                d_271_v109_ = nw26_
                d_272_v110_: _dafny.Seq
                d_272_v110_ = _dafny.SeqWithoutIsStrInference([d_271_v109_])
                d_273_v111_: D19
                d_273_v111_ = D19_DC48(d_272_v110_)
                d_274_v112_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_274_v112_ = nw27_
                index24_ = default__.safeIndex(181, (d_274_v112_).length(0))
                (d_274_v112_)[index24_] = d_105_v8_
                index25_ = default__.safeIndex(181, (d_274_v112_).length(0))
                rhs34_ = d_273_v111_
                rhs35_ = _dafny.SeqWithoutIsStrInference([len(d_268_v106_) for d_275_i9_ in range(default__.abs(449))])
                rhs36_ = d_105_v8_
                rhs37_ = (d_255_v97_).f6
                lhs14_ = d_274_v112_
                lhs15_ = default__.safeIndex(181, (d_274_v112_).length(0))
                d_273_v111_ = rhs34_
                d_124_v18_ = rhs35_
                lhs14_[lhs15_] = rhs36_
                r1 = rhs37_
                d_276_v113_: _dafny.Array
                nw28_ = _dafny.Array(D2.default()(), 16)
                d_276_v113_ = nw28_
                d_277_v114_: _dafny.Map
                out0_: _dafny.Map
                out0_ = (d_255_v97_).m2(d_276_v113_, globalState)
                d_277_v114_ = out0_
            elif True:
                d_99_v2_ = ((len((d_261_v101_)[default__.safeIndex(38, (d_261_v101_).length(0))])) - (d_99_v2_)) + (len((d_108_v11_) + (d_108_v11_)))
                d_278_v115_: _dafny.Map
                d_278_v115_ = _dafny.Map({(_dafny.MultiSet([True])).cardinality: not((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))])})
                r1 = not (((d_278_v115_)[d_99_v2_] if (d_99_v2_) in (d_278_v115_) else d_104_v7_)) or ((d_96_v0_)[default__.safeIndex(151, (d_96_v0_).length(0))])
                r1 = default__.fm3(d_99_v2_, globalState)
                d_133_v20_ = (d_133_v20_).set(d_99_v2_, ((0) - ((0) - (d_99_v2_))) - (d_99_v2_))
                d_279_v116_: C0
                nw29_ = C0()
                nw29_.ctor__(d_98_v1_)
                d_279_v116_ = nw29_
        r0 = True
        d_280_v117_: _dafny.MultiSet
        d_280_v117_ = _dafny.MultiSet([d_124_v18_, d_124_v18_, _dafny.SeqWithoutIsStrInference([d_99_v2_ for d_281_i10_ in range(default__.abs(561))])])
        r1 = ((d_280_v117_).cardinality) == (d_99_v2_)
        r2 = default__.fm7(d_104_v7_, d_104_v7_, globalState)
        d_282_v118_: _dafny.Seq
        d_282_v118_ = _dafny.SeqWithoutIsStrInference([(d_133_v20_).set(d_99_v2_, (0) - (d_99_v2_)), d_133_v20_, d_133_v20_, d_133_v20_, _dafny.Map({d_99_v2_: d_99_v2_})])
        r3 = d_282_v118_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_283_globalState_: GlobalState
        nw30_ = GlobalState()
        nw30_.ctor__(True)
        d_283_globalState_ = nw30_
        d_284_v0_: bool
        d_284_v0_ = True
        d_285_v1_: _dafny.Array
        def lambda4_(d_286_i0_):
            return not(not((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejsfewk"))) for d_287_i1_ in range(default__.abs(409))]), _dafny.SeqWithoutIsStrInference([902, 812, 389, 649])])) <= (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([485, 593])]))))

        init2_ = lambda4_
        nw31_ = _dafny.Array(None, 15)
        for i0_2_ in range(nw31_.length(0)):
            nw31_[i0_2_] = init2_(i0_2_)
        d_285_v1_ = nw31_
        d_288_v2_: int
        d_288_v2_ = 59
        rhs38_ = (default__.fm0(d_283_globalState_)) > (d_288_v2_)
        rhs39_ = (not(not(False))) and ((d_288_v2_) >= (d_288_v2_))
        rhs40_ = d_285_v1_
        rhs41_ = (d_284_v0_ if d_284_v0_ else d_284_v0_)
        rhs42_ = (d_288_v2_) - (d_288_v2_)
        d_284_v0_ = rhs38_
        d_284_v0_ = rhs39_
        d_285_v1_ = rhs40_
        d_284_v0_ = rhs41_
        d_288_v2_ = rhs42_
        d_289_v3_: bool
        d_290_v4_: bool
        d_291_v5_: _dafny.Seq
        d_292_v6_: _dafny.Seq
        out1_: bool
        out2_: bool
        out3_: _dafny.Seq
        out4_: _dafny.Seq
        out1_, out2_, out3_, out4_ = default__.m0(_dafny.Map({d_285_v1_: d_284_v0_}), d_283_globalState_)
        d_289_v3_ = out1_
        d_290_v4_ = out2_
        d_291_v5_ = out3_
        d_292_v6_ = out4_
        d_293_v7_: _dafny.Array
        def lambda5_(d_294_v2_):
            def lambda6_(d_295_i2_):
                return (_dafny.Map({d_294_v2_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_296_i3_ in range(default__.abs(734))]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_294_v2_ for d_297_i4_ in range(default__.abs(503))])): d_294_v2_}))

            return lambda6_

        init3_ = lambda5_(d_288_v2_)
        nw32_ = _dafny.Array(None, 2)
        for i0_3_ in range(nw32_.length(0)):
            nw32_[i0_3_] = init3_(i0_3_)
        d_293_v7_ = nw32_
        d_298_v8_: _dafny.Seq
        d_298_v8_ = _dafny.SeqWithoutIsStrInference([d_291_v5_])
        d_299_v9_: _dafny.MultiSet
        d_299_v9_ = _dafny.MultiSet([d_284_v0_, d_289_v3_])
        d_300_v10_: _dafny.Seq
        d_300_v10_ = _dafny.SeqWithoutIsStrInference([d_299_v9_])
        d_301_v11_: str
        d_301_v11_ = _dafny.CodePoint('p')
        index26_ = default__.safeIndex(1, (d_293_v7_).length(0))
        (d_293_v7_)[index26_] = default__.fm1((d_298_v8_)[default__.safeIndex(((d_300_v10_)[default__.safeIndex(d_288_v2_, len(d_300_v10_))]).cardinality, len(d_298_v8_))], d_288_v2_, d_288_v2_, d_301_v11_, d_283_globalState_)
        d_302_v12_: _dafny.Map
        d_302_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_301_v11_, d_301_v11_]): d_284_v0_})
        d_303_v13_: _dafny.Map
        d_303_v13_ = _dafny.Map({len(d_302_v12_): d_288_v2_})
        d_304_v14_: D0
        d_304_v14_ = D0_DC0(d_288_v2_)
        index27_ = default__.safeIndex(1, (d_293_v7_).length(0))
        (d_293_v7_)[index27_] = (d_303_v13_) | ((d_303_v13_) | (_dafny.Map({default__.fm0(d_283_globalState_): (d_304_v14_).cf0})))
        d_290_v4_ = not((d_288_v2_) < (d_288_v2_))
        d_305_v15_: _dafny.Array
        def lambda7_(d_306_v2_):
            def lambda8_(d_307_i5_):
                return (d_307_i5_) - (d_306_v2_)

            return lambda8_

        init4_ = lambda7_(d_288_v2_)
        nw33_ = _dafny.Array(None, 5)
        for i0_4_ in range(nw33_.length(0)):
            nw33_[i0_4_] = init4_(i0_4_)
        d_305_v15_ = nw33_
        index28_ = default__.safeIndex(353, (d_305_v15_).length(0))
        (d_305_v15_)[index28_] = (default__.fm0(d_283_globalState_)) * (len(_dafny.SeqWithoutIsStrInference([d_290_v4_])))
        pat_let_tv10_ = d_289_v3_
        pat_let_tv11_ = d_299_v9_
        pat_let_tv12_ = d_299_v9_
        pat_let_tv13_ = d_289_v3_
        pat_let_tv14_ = d_288_v2_
        pat_let_tv15_ = d_288_v2_
        index29_ = default__.safeIndex(353, (d_305_v15_).length(0))
        def lambda9_(source9_):
            if source9_.is_DC1:
                d_308___mcc_h0_ = source9_.cf1
                d_309_cf1_ = d_308___mcc_h0_
                if (pat_let_tv10_) in (pat_let_tv11_):
                    return (pat_let_tv12_)[pat_let_tv13_]
                elif True:
                    return (0) - (pat_let_tv14_)
            elif True:
                d_310___mcc_h1_ = source9_.cf0
                d_311_cf0_ = d_310___mcc_h1_
                return pat_let_tv15_

        rhs43_ = d_305_v15_
        rhs44_ = lambda9_(d_304_v14_)
        rhs45_ = default__.safeModuloInt(d_288_v2_, d_288_v2_)
        rhs46_ = d_304_v14_
        rhs47_ = d_288_v2_
        lhs16_ = d_305_v15_
        lhs17_ = default__.safeIndex(353, (d_305_v15_).length(0))
        d_305_v15_ = rhs43_
        d_288_v2_ = rhs44_
        d_288_v2_ = rhs45_
        d_304_v14_ = rhs46_
        lhs16_[lhs17_] = rhs47_
        d_312_i6_: int
        d_312_i6_ = 0
        with _dafny.label("0"):
            while ((d_302_v12_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjwie"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjwie"))) in (d_302_v12_) else d_290_v4_):
                with _dafny.c_label("0"):
                    if (d_312_i6_) >= (100):
                        raise _dafny.Break("0")
                    d_312_i6_ = (d_312_i6_) + (1)
                    d_288_v2_ = (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]
                    d_313_v16_: _dafny.Map
                    d_313_v16_ = _dafny.Map({d_285_v1_: d_290_v4_})
                    d_314_v17_: bool
                    d_315_v18_: bool
                    d_316_v19_: _dafny.Seq
                    d_317_v20_: _dafny.Seq
                    out5_: bool
                    out6_: bool
                    out7_: _dafny.Seq
                    out8_: _dafny.Seq
                    out5_, out6_, out7_, out8_ = default__.m0(d_313_v16_, d_283_globalState_)
                    d_314_v17_ = out5_
                    d_315_v18_ = out6_
                    d_316_v19_ = out7_
                    d_317_v20_ = out8_
                    d_318_v21_: _dafny.Set
                    d_318_v21_ = _dafny.Set({d_288_v2_})
                    index30_ = default__.safeIndex(353, (d_305_v15_).length(0))
                    (d_305_v15_)[index30_] = (len((d_316_v19_ if d_315_v18_ else d_291_v5_))) + ((len(d_318_v21_)) - (d_288_v2_))
                    d_319_v22_: _dafny.Seq
                    d_319_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtuvm"))
                    d_319_v22_ = (d_319_v22_) + (d_319_v22_)
                    pass
            pass
        d_288_v2_ = default__.safeDivisionInt((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_288_v2_)
        d_320_v23_: _dafny.Seq
        d_320_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyfj"))
        d_321_v24_: _dafny.Seq
        d_321_v24_ = _dafny.SeqWithoutIsStrInference([len(d_320_v23_)])
        d_322_v25_: _dafny.Map
        d_322_v25_ = _dafny.Map({d_289_v3_: len(d_321_v24_)})
        d_323_v26_: _dafny.Seq
        d_323_v26_ = _dafny.SeqWithoutIsStrInference([d_322_v25_])
        d_324_v27_: _dafny.Map
        d_324_v27_ = _dafny.Map({(150) >= (d_288_v2_): default__.safeDivisionInt(len(_dafny.Map({(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]: d_290_v4_})), len((d_323_v26_)[default__.safeIndex(d_288_v2_, len(d_323_v26_))]))})
        d_322_v25_ = d_324_v27_
        d_325_v28_: _dafny.Map
        d_325_v28_ = _dafny.Map({(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]: d_320_v23_})
        d_320_v23_ = (d_320_v23_) + ((((d_325_v28_)[d_288_v2_] if (d_288_v2_) in (d_325_v28_) else default__.fm2(not(d_290_v4_), d_283_globalState_))).set(default__.safeIndex((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], len(((d_325_v28_)[d_288_v2_] if (d_288_v2_) in (d_325_v28_) else default__.fm2(not(d_290_v4_), d_283_globalState_)))), _dafny.CodePoint('q')))
        d_326_i7_: int
        d_326_i7_ = 0
        with _dafny.label("1"):
            while d_290_v4_:
                with _dafny.c_label("1"):
                    if (d_326_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_326_i7_ = (d_326_i7_) + (1)
                    d_327_v29_: _dafny.Array
                    nw34_ = _dafny.Array(_dafny.Seq({}), 13)
                    d_327_v29_ = nw34_
                    rhs48_ = (d_291_v5_).set(default__.safeIndex((((d_324_v27_)[False] if (False) in (d_324_v27_) else (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])) + (796), len(d_291_v5_)), d_290_v4_)
                    rhs49_ = d_327_v29_
                    d_291_v5_ = rhs48_
                    d_327_v29_ = rhs49_
                    d_328_v30_: D0
                    d_328_v30_ = D0_DC1(d_288_v2_)
                    source10_ = d_328_v30_
                    if source10_.is_DC1:
                        d_329___mcc_h2_ = source10_.cf1
                        d_330_cf1_ = d_329___mcc_h2_
                        d_290_v4_ = not(d_289_v3_)
                        d_330_cf1_ = d_288_v2_
                        d_331_v31_: D1
                        d_331_v31_ = D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrs")))
                        d_332_v32_: _dafny.Array
                        nw35_ = _dafny.Array(None, 23)
                        nw35_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkmnvysfk"))
                        nw35_[int(1)] = (d_320_v23_) + (d_320_v23_)
                        nw35_[int(2)] = (d_320_v23_) + (d_320_v23_)
                        nw35_[int(3)] = d_320_v23_
                        nw35_[int(4)] = d_320_v23_
                        nw35_[int(5)] = ((d_331_v31_).cf2).set(default__.safeIndex((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], len((d_331_v31_).cf2)), d_301_v11_)
                        nw35_[int(6)] = ((d_331_v31_).cf2) + (d_320_v23_)
                        nw35_[int(7)] = _dafny.SeqWithoutIsStrInference([d_301_v11_ for d_333_i8_ in range(default__.abs(385))])
                        nw35_[int(8)] = d_320_v23_
                        nw35_[int(9)] = (d_320_v23_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "up")))
                        nw35_[int(10)] = (d_320_v23_) + (d_320_v23_)
                        nw35_[int(11)] = d_320_v23_
                        nw35_[int(12)] = d_320_v23_
                        nw35_[int(13)] = d_320_v23_
                        nw35_[int(14)] = d_320_v23_
                        nw35_[int(15)] = d_320_v23_
                        nw35_[int(16)] = d_320_v23_
                        nw35_[int(17)] = d_320_v23_
                        nw35_[int(18)] = d_320_v23_
                        nw35_[int(19)] = (_dafny.SeqWithoutIsStrInference([d_301_v11_ for d_334_i9_ in range(default__.abs(14))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpfqtw")))
                        nw35_[int(20)] = d_320_v23_
                        nw35_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqpxpwr"))
                        nw35_[int(22)] = d_320_v23_
                        d_332_v32_ = nw35_
                        index31_ = default__.safeIndex(373, (d_332_v32_).length(0))
                        (d_332_v32_)[index31_] = _dafny.SeqWithoutIsStrInference([d_301_v11_ for d_335_i10_ in range(default__.abs(-950))])
                        index32_ = default__.safeIndex(373, (d_332_v32_).length(0))
                        (d_332_v32_)[index32_] = d_320_v23_
                        d_336_v33_: _dafny.Seq
                        d_336_v33_ = _dafny.SeqWithoutIsStrInference([d_321_v24_, (d_321_v24_ if d_289_v3_ else d_321_v24_), d_321_v24_])
                        d_337_v34_: _dafny.Map
                        d_337_v34_ = _dafny.Map({d_284_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdb"))})
                        d_336_v33_ = ((d_336_v33_) + (d_336_v33_)).set(default__.safeIndex(len(d_337_v34_), len((d_336_v33_) + (d_336_v33_))), d_321_v24_)
                    elif True:
                        d_338___mcc_h3_ = source10_.cf0
                        d_339_cf0_ = d_338___mcc_h3_
                        d_340_v35_: _dafny.Map
                        d_340_v35_ = _dafny.Map({d_284_v0_: d_289_v3_})
                        d_341_v36_: _dafny.Map
                        d_341_v36_ = _dafny.Map({(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]: d_285_v1_})
                        d_342_v37_: _dafny.Seq
                        d_342_v37_ = _dafny.SeqWithoutIsStrInference([d_320_v23_])
                        index33_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        rhs50_ = len(d_341_v36_)
                        rhs51_ = _dafny.Map({(d_320_v23_) not in (d_342_v37_): default__.fm3(d_339_cf0_, d_283_globalState_)})
                        rhs52_ = (default__.safeModuloInt(215, (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])) + (d_288_v2_)
                        lhs18_ = d_305_v15_
                        lhs19_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        lhs18_[lhs19_] = rhs50_
                        d_340_v35_ = rhs51_
                        d_339_cf0_ = rhs52_
                        d_343_v38_: _dafny.Map
                        d_343_v38_ = _dafny.Map({d_321_v24_: default__.fm0(d_283_globalState_)})
                        d_344_v39_: D2
                        d_344_v39_ = D2_DC5(d_324_v27_)
                        d_345_v40_: _dafny.Set
                        d_345_v40_ = _dafny.Set({False})
                        d_346_v41_: _dafny.Map
                        d_346_v41_ = _dafny.Map({len(d_345_v40_): True})
                        d_347_v42_: _dafny.Array
                        nw36_ = _dafny.Array(None, 24)
                        nw36_[int(0)] = d_322_v25_
                        nw36_[int(1)] = (_dafny.Map({d_289_v3_: d_288_v2_})) | ((((d_344_v39_).cf8).set(d_284_v0_, 500)).set(d_284_v0_, len(d_345_v40_)))
                        nw36_[int(2)] = _dafny.Map({d_290_v4_: (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]})
                        nw36_[int(3)] = d_322_v25_
                        nw36_[int(4)] = (d_324_v27_) | ((d_323_v26_)[default__.safeIndex(133, len(d_323_v26_))])
                        nw36_[int(5)] = (d_324_v27_).set(not(d_290_v4_), (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])
                        nw36_[int(6)] = default__.fm4(d_288_v2_, ((d_346_v41_)[(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]] if ((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]) in (d_346_v41_) else default__.fm3((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_283_globalState_)), d_283_globalState_)
                        nw36_[int(7)] = d_324_v27_
                        nw36_[int(8)] = d_322_v25_
                        nw36_[int(9)] = d_322_v25_
                        nw36_[int(10)] = _dafny.Map({True: len(d_291_v5_)})
                        nw36_[int(11)] = d_322_v25_
                        nw36_[int(12)] = d_322_v25_
                        nw36_[int(13)] = d_324_v27_
                        nw36_[int(14)] = (_dafny.Map({(d_291_v5_)[default__.safeIndex((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], len(d_291_v5_))]: len((d_293_v7_)[default__.safeIndex(1, (d_293_v7_).length(0))])})) | (d_322_v25_)
                        nw36_[int(15)] = d_324_v27_
                        nw36_[int(16)] = default__.fm4(d_288_v2_, d_284_v0_, d_283_globalState_)
                        nw36_[int(17)] = d_322_v25_
                        nw36_[int(18)] = d_322_v25_
                        nw36_[int(19)] = d_324_v27_
                        nw36_[int(20)] = d_322_v25_
                        nw36_[int(21)] = (_dafny.Map({d_290_v4_: (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]})) | (d_324_v27_)
                        nw36_[int(22)] = d_324_v27_
                        nw36_[int(23)] = d_322_v25_
                        d_347_v42_ = nw36_
                        index34_ = default__.safeIndex(103, (d_347_v42_).length(0))
                        (d_347_v42_)[index34_] = (d_324_v27_) | (d_324_v27_)
                        d_348_v43_: _dafny.Seq
                        d_348_v43_ = _dafny.SeqWithoutIsStrInference([d_343_v38_])
                        d_349_v44_: _dafny.Map
                        d_349_v44_ = _dafny.Map({_dafny.Map({(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]: d_284_v0_}): 169})
                        index35_ = default__.safeIndex(103, (d_347_v42_).length(0))
                        index36_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        index37_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        rhs53_ = ((d_343_v38_) | ((d_348_v43_)[default__.safeIndex(-648, len(d_348_v43_))])) | ((d_343_v38_) | (d_343_v38_))
                        rhs54_ = d_340_v35_
                        rhs55_ = ((d_324_v27_).set(d_290_v4_, d_288_v2_)) | ((d_324_v27_).set(not(d_290_v4_), d_339_cf0_))
                        rhs56_ = len(d_349_v44_)
                        rhs57_ = default__.safeModuloInt(d_339_cf0_, (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])
                        lhs20_ = d_347_v42_
                        lhs21_ = default__.safeIndex(103, (d_347_v42_).length(0))
                        lhs22_ = d_305_v15_
                        lhs23_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        lhs24_ = d_305_v15_
                        lhs25_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        d_343_v38_ = rhs53_
                        d_340_v35_ = rhs54_
                        lhs20_[lhs21_] = rhs55_
                        lhs22_[lhs23_] = rhs56_
                        lhs24_[lhs25_] = rhs57_
                        index38_ = default__.safeIndex(353, (d_305_v15_).length(0))
                        (d_305_v15_)[index38_] = d_288_v2_
                        d_350_v45_: _dafny.Array
                        nw37_ = _dafny.Array(_dafny.Seq({}), 17)
                        d_350_v45_ = nw37_
                        d_351_v46_: _dafny.Seq
                        d_351_v46_ = _dafny.SeqWithoutIsStrInference([d_285_v1_, d_285_v1_])
                        index39_ = default__.safeIndex(521, (d_350_v45_).length(0))
                        (d_350_v45_)[index39_] = (d_351_v46_ if d_290_v4_ else _dafny.SeqWithoutIsStrInference([d_285_v1_]))
                        index40_ = default__.safeIndex(521, (d_350_v45_).length(0))
                        (d_350_v45_)[index40_] = (d_351_v46_) + (d_351_v46_)
                    d_285_v1_ = d_285_v1_
                    d_290_v4_ = d_289_v3_
                    pass
            pass
        d_352_v47_: D1
        d_352_v47_ = D1_DC4(d_284_v0_, (len(d_320_v23_)) in (d_321_v24_))
        d_352_v47_ = D1_DC4((d_291_v5_) < (d_291_v5_), d_289_v3_)
        d_353_v48_: _dafny.Array
        def lambda10_(d_354_v3_, d_355_v15_):
            def lambda11_(d_356_i11_):
                return D1_DC3(d_354_v3_, (d_355_v15_)[default__.safeIndex(353, (d_355_v15_).length(0))], (d_355_v15_)[default__.safeIndex(353, (d_355_v15_).length(0))])

            return lambda11_

        init5_ = lambda10_(d_289_v3_, d_305_v15_)
        nw38_ = _dafny.Array(None, 24)
        for i0_5_ in range(nw38_.length(0)):
            nw38_[i0_5_] = init5_(i0_5_)
        d_353_v48_ = nw38_
        d_353_v48_ = d_353_v48_
        d_357_i12_: int
        d_357_i12_ = 0
        with _dafny.label("2"):
            while d_284_v0_:
                with _dafny.c_label("2"):
                    if (d_357_i12_) >= (100):
                        raise _dafny.Break("2")
                    d_357_i12_ = (d_357_i12_) + (1)
                    d_358_v49_: _dafny.MultiSet
                    d_358_v49_ = _dafny.MultiSet([d_288_v2_])
                    d_359_v50_: _dafny.Map
                    d_359_v50_ = _dafny.Map({d_358_v49_: d_320_v23_})
                    d_359_v50_ = (_dafny.Map({d_358_v49_: d_320_v23_})) | (_dafny.Map({d_358_v49_: d_320_v23_}))
                    d_288_v2_ = (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]
                    d_288_v2_ = (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]
                    d_360_v51_: _dafny.Map
                    d_360_v51_ = _dafny.Map({d_285_v1_: d_284_v0_})
                    d_361_v52_: bool
                    d_362_v53_: bool
                    d_363_v54_: _dafny.Seq
                    d_364_v55_: _dafny.Seq
                    out9_: bool
                    out10_: bool
                    out11_: _dafny.Seq
                    out12_: _dafny.Seq
                    out9_, out10_, out11_, out12_ = default__.m0(d_360_v51_, d_283_globalState_)
                    d_361_v52_ = out9_
                    d_362_v53_ = out10_
                    d_363_v54_ = out11_
                    d_364_v55_ = out12_
                    pass
            pass
        hi0_ = (0) - (default__.safeModuloInt((0) - (d_288_v2_), d_288_v2_))
        for d_365_i13_ in range(-33, hi0_):
            index41_ = default__.safeIndex(625, (d_285_v1_).length(0))
            (d_285_v1_)[index41_] = d_290_v4_
            d_366_v56_: D3
            d_366_v56_ = D3_DC8(_dafny.Set({d_290_v4_, d_289_v3_, d_290_v4_}))
            index42_ = default__.safeIndex(625, (d_285_v1_).length(0))
            (d_285_v1_)[index42_] = default__.fm3(len((d_366_v56_).cf14), d_283_globalState_)
            index43_ = default__.safeIndex(353, (d_305_v15_).length(0))
            (d_305_v15_)[index43_] = ((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]) + (d_365_i13_)
            d_288_v2_ = d_288_v2_
            d_367_v57_: _dafny.Map
            d_367_v57_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_288_v2_ for d_368_i14_ in range(default__.abs(262))]): d_320_v23_})
            d_288_v2_ = len((((d_367_v57_)[_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]])] if (_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]])) in (d_367_v57_) else _dafny.SeqWithoutIsStrInference([d_301_v11_ for d_369_i15_ in range(default__.abs(45))]))).set(default__.safeIndex(d_365_i13_, len(((d_367_v57_)[_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]])] if (_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]])) in (d_367_v57_) else _dafny.SeqWithoutIsStrInference([d_301_v11_ for d_370_i15_ in range(default__.abs(45))])))), d_301_v11_))
        index44_ = default__.safeIndex(709, (d_285_v1_).length(0))
        (d_285_v1_)[index44_] = d_284_v0_
        index45_ = default__.safeIndex(709, (d_285_v1_).length(0))
        (d_285_v1_)[index45_] = d_290_v4_
        hi1_ = default__.safeModuloInt(d_288_v2_, (0) - (d_288_v2_))
        for d_371_i16_ in range((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], hi1_):
            index46_ = default__.safeIndex(709, (d_285_v1_).length(0))
            (d_285_v1_)[index46_] = (d_289_v3_) or (d_289_v3_)
            d_372_v58_: _dafny.Map
            d_372_v58_ = _dafny.Map({d_289_v3_: _dafny.Map({d_284_v0_: d_371_i16_})})
            source11_ = D2_DC5((d_322_v25_) | (((d_372_v58_)[d_284_v0_] if (d_284_v0_) in (d_372_v58_) else d_322_v25_)))
            if source11_.is_DC6:
                d_373___mcc_h4_ = source11_.cf9
                d_374___mcc_h5_ = source11_.cf10
                d_375___mcc_h6_ = source11_.cf11
                d_376___mcc_h7_ = source11_.cf12
                d_377_cf12_ = d_376___mcc_h7_
                d_378_cf11_ = d_375___mcc_h6_
                d_379_cf10_ = d_374___mcc_h5_
                d_380_cf9_ = d_373___mcc_h4_
                index47_ = default__.safeIndex(353, (d_305_v15_).length(0))
                rhs58_ = (default__.fm5(_dafny.Map({d_320_v23_: d_371_i16_}), d_283_globalState_)).cardinality
                rhs59_ = (D3_DC9(d_377_cf12_, (d_291_v5_)[default__.safeIndex(((d_303_v13_)[d_377_cf12_] if (d_377_cf12_) in (d_303_v13_) else d_371_i16_), len(d_291_v5_))])).cf16
                lhs26_ = d_305_v15_
                lhs27_ = default__.safeIndex(353, (d_305_v15_).length(0))
                lhs26_[lhs27_] = rhs58_
                d_379_cf10_ = rhs59_
                index48_ = default__.safeIndex(709, (d_285_v1_).length(0))
                (d_285_v1_)[index48_] = default__.fm3(default__.safeDivisionInt(d_371_i16_, d_377_cf12_), d_283_globalState_)
                d_381_v59_: _dafny.Set
                d_381_v59_ = _dafny.Set({(0) - (d_377_cf12_), default__.safeModuloInt(d_288_v2_, len(default__.fm6(d_378_cf11_, d_301_v11_, d_283_globalState_)))})
                d_381_v59_ = d_381_v59_
                d_321_v24_ = d_321_v24_
            elif source11_.is_DC5:
                d_382___mcc_h8_ = source11_.cf8
                d_383_cf8_ = d_382___mcc_h8_
                index49_ = default__.safeIndex(353, (d_305_v15_).length(0))
                (d_305_v15_)[index49_] = 147
                d_384_v60_: _dafny.Set
                d_384_v60_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([(d_285_v1_)[default__.safeIndex(709, (d_285_v1_).length(0))]])), d_371_i16_})
                d_303_v13_ = (d_303_v13_).set(((d_322_v25_)[True] if (True) in (d_322_v25_) else len(d_384_v60_)), default__.safeDivisionInt(d_371_i16_, ((d_383_cf8_)[d_290_v4_] if (d_290_v4_) in (d_383_cf8_) else (0) - (d_288_v2_))))
                d_385_v61_: _dafny.Array
                def lambda12_(d_386_v5_):
                    def lambda13_(d_387_i17_):
                        return d_386_v5_

                    return lambda13_

                init6_ = lambda12_(d_291_v5_)
                nw39_ = _dafny.Array(None, 8)
                for i0_6_ in range(nw39_.length(0)):
                    nw39_[i0_6_] = init6_(i0_6_)
                d_385_v61_ = nw39_
                d_388_v62_: _dafny.Map
                d_388_v62_ = _dafny.Map({(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]: d_291_v5_})
                index50_ = default__.safeIndex(609, (d_385_v61_).length(0))
                (d_385_v61_)[index50_] = ((d_388_v62_)[d_288_v2_] if (d_288_v2_) in (d_388_v62_) else d_291_v5_)
                index51_ = default__.safeIndex(609, (d_385_v61_).length(0))
                (d_385_v61_)[index51_] = (default__.fm7(d_290_v4_, True, d_283_globalState_)) + (d_291_v5_)
                d_389_v63_: _dafny.Map
                d_389_v63_ = _dafny.Map({d_321_v24_: True})
                index52_ = default__.safeIndex(353, (d_305_v15_).length(0))
                index53_ = default__.safeIndex(709, (d_285_v1_).length(0))
                rhs60_ = (len(d_320_v23_) if d_284_v0_ else len(d_389_v63_))
                rhs61_ = ((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]) < ((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])
                lhs28_ = d_305_v15_
                lhs29_ = default__.safeIndex(353, (d_305_v15_).length(0))
                lhs30_ = d_285_v1_
                lhs31_ = default__.safeIndex(709, (d_285_v1_).length(0))
                lhs28_[lhs29_] = rhs60_
                lhs30_[lhs31_] = rhs61_
            elif True:
                d_390___mcc_h9_ = source11_.cf13
                d_391_cf13_ = d_390___mcc_h9_
                d_392_v64_: _dafny.Map
                d_392_v64_ = _dafny.Map({d_289_v3_: d_290_v4_})
                d_393_v65_: _dafny.Map
                d_393_v65_ = _dafny.Map({d_392_v64_: d_371_i16_})
                d_288_v2_ = default__.safeModuloInt(((d_393_v65_)[d_392_v64_] if (d_392_v64_) in (d_393_v65_) else d_371_i16_), default__.safeModuloInt(d_288_v2_, len(_dafny.SeqWithoutIsStrInference([d_301_v11_ for d_394_i18_ in range(default__.abs(746))]))))
                d_395_v66_: _dafny.Array
                def lambda14_(d_396_v23_):
                    def lambda15_(d_397_i19_):
                        return (D1_DC2(d_396_v23_)).cf2

                    return lambda15_

                init7_ = lambda14_(d_320_v23_)
                nw40_ = _dafny.Array(None, 26)
                for i0_7_ in range(nw40_.length(0)):
                    nw40_[i0_7_] = init7_(i0_7_)
                d_395_v66_ = nw40_
                nw41_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_395_v66_ = nw41_
                index54_ = default__.safeIndex(353, (d_305_v15_).length(0))
                (d_305_v15_)[index54_] = (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]
                d_398_v67_: _dafny.Seq
                d_398_v67_ = _dafny.SeqWithoutIsStrInference([d_305_v15_, d_305_v15_, d_305_v15_, (d_305_v15_ if d_284_v0_ else d_305_v15_)])
                d_398_v67_ = d_398_v67_
            d_399_v68_: _dafny.MultiSet
            d_399_v68_ = _dafny.MultiSet([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_288_v2_])
            d_400_v69_: _dafny.Map
            d_400_v69_ = _dafny.Map({d_301_v11_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmmd")))})
            d_401_v70_: D4
            d_401_v70_ = D4_DC12(d_399_v68_)
            d_402_v71_: _dafny.Map
            d_402_v71_ = _dafny.Map({d_301_v11_: (d_399_v68_).set(d_371_i16_, default__.abs(d_371_i16_))})
            d_403_v72_: _dafny.Map
            d_403_v72_ = _dafny.Map({d_289_v3_: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_371_i16_, (d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]]))})
            d_404_v73_: _dafny.Array
            nw42_ = _dafny.Array(None, 29)
            nw42_[int(0)] = d_399_v68_
            nw42_[int(1)] = d_399_v68_
            nw42_[int(2)] = _dafny.MultiSet([d_288_v2_])
            nw42_[int(3)] = (d_399_v68_).intersection(d_399_v68_)
            nw42_[int(4)] = (d_399_v68_) - (_dafny.MultiSet(d_321_v24_))
            nw42_[int(5)] = _dafny.MultiSet([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]])
            nw42_[int(6)] = _dafny.MultiSet((d_321_v24_).set(default__.safeIndex((d_299_v9_).cardinality, len(d_321_v24_)), len(d_321_v24_)))
            nw42_[int(7)] = _dafny.MultiSet([((d_400_v69_)[d_301_v11_] if (d_301_v11_) in (d_400_v69_) else len(d_320_v23_))])
            nw42_[int(8)] = d_399_v68_
            nw42_[int(9)] = _dafny.MultiSet(d_321_v24_)
            nw42_[int(10)] = d_399_v68_
            nw42_[int(11)] = (d_401_v70_).cf20
            nw42_[int(12)] = d_399_v68_
            nw42_[int(13)] = d_399_v68_
            nw42_[int(14)] = (d_399_v68_).intersection((_dafny.MultiSet(d_321_v24_)).set(d_371_i16_, default__.abs(d_371_i16_)))
            nw42_[int(15)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(((d_293_v7_)[default__.safeIndex(1, (d_293_v7_).length(0))])[len(_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_288_v2_]))] if (len(_dafny.SeqWithoutIsStrInference([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_288_v2_]))) in ((d_293_v7_)[default__.safeIndex(1, (d_293_v7_).length(0))]) else len(d_291_v5_))]))) - (_dafny.MultiSet([(d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_371_i16_]))
            nw42_[int(16)] = default__.fm8((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], d_320_v23_, d_283_globalState_)
            nw42_[int(17)] = (d_399_v68_).intersection(d_399_v68_)
            nw42_[int(18)] = d_399_v68_
            nw42_[int(19)] = _dafny.MultiSet(d_321_v24_)
            nw42_[int(20)] = d_399_v68_
            nw42_[int(21)] = _dafny.MultiSet(d_321_v24_)
            nw42_[int(22)] = (d_399_v68_) - (d_399_v68_)
            nw42_[int(23)] = ((d_402_v71_)[d_301_v11_] if (d_301_v11_) in (d_402_v71_) else (d_399_v68_).set((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))], default__.abs((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))])))
            nw42_[int(24)] = ((d_403_v72_)[True] if (True) in (d_403_v72_) else d_399_v68_)
            nw42_[int(25)] = d_399_v68_
            nw42_[int(26)] = _dafny.MultiSet((d_321_v24_) + (d_321_v24_))
            nw42_[int(27)] = d_399_v68_
            nw42_[int(28)] = d_399_v68_
            d_404_v73_ = nw42_
            d_404_v73_ = d_404_v73_
            d_405_v74_: _dafny.Map
            d_405_v74_ = _dafny.Map({d_284_v0_: ((d_305_v15_)[default__.safeIndex(353, (d_305_v15_).length(0))]) == (default__.fm0(d_283_globalState_))})
            rhs62_ = (d_285_v1_)[default__.safeIndex(709, (d_285_v1_).length(0))]
            rhs63_ = default__.fm9(d_283_globalState_)
            d_284_v0_ = rhs62_
            d_405_v74_ = rhs63_
        _dafny.print(_dafny.string_of((d_283_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_284_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_285_v1_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_288_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_289_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_290_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v5_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v6_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({-554: -554, 0: 0}), _dafny.Map({-554: -554}), _dafny.Map({-554: -554}), _dafny.Map({-554: -554}), _dafny.Map({0: 0})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v7_)[0]) == (_dafny.Map({0: 734, 503: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v7_)[1]) == (_dafny.Map({1: 0, 9: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_298_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v9_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v10_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_301_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_302_v12_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p'), _dafny.CodePoint('p')]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_303_v13_) == (_dafny.Map({1: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v14_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v15_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v15_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v15_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v15_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_305_v15_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_312_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_320_v23_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_321_v24_) == (_dafny.SeqWithoutIsStrInference([4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_322_v25_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_323_v26_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: 1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_324_v27_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_325_v28_) == (_dafny.Map({0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyfj"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_326_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v47_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v47_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[0]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[0]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[0]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[1]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[1]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[1]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[2]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[2]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[2]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[3]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[3]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[3]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[4]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[4]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[4]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[5]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[5]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[5]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[6]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[6]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[6]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[7]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[7]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[7]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[8]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[8]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[8]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[9]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[9]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[9]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[10]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[10]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[10]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[11]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[11]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[11]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[12]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[12]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[12]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[13]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[13]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[13]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[14]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[14]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[14]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[15]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[15]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[15]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[16]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[16]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[16]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[17]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[17]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[17]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[18]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[18]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[18]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[19]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[19]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[19]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[20]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[20]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[20]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[21]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[21]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[21]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[22]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[22]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[22]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[23]).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[23]).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_353_v48_)[23]).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_357_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
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

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({self.cf2.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(_dafny.CodePoint('D'), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0), False)
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

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC17(D6, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC20(D6, NamedTuple('DC20', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC21(D7, NamedTuple('DC21', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(int(0), int(0), False)
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

class D8_DC23(D8, NamedTuple('DC23', [('cf38', Any), ('cf39', Any), ('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC26(int(0))
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

class D9_DC26(D9, NamedTuple('DC26', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC29(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)

class D10_DC29(D10, NamedTuple('DC29', [('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC32(D12, NamedTuple('DC32', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC34(False, False, int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)

class D13_DC34(D13, NamedTuple('DC34', [('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {self.cf56.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC38(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)

class D14_DC38(D14, NamedTuple('DC38', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({self.cf64.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC39(D14, NamedTuple('DC39', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC37(D14, NamedTuple('DC37', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D15_DC40)

class D15_DC40(D15, NamedTuple('DC40', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC40({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC40) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC42(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)

class D16_DC42(D16, NamedTuple('DC42', [('cf70', Any), ('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC41(D16, NamedTuple('DC41', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)

class D17_DC43(D17, NamedTuple('DC43', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC45(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D18_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D18_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)

class D18_DC45(D18, NamedTuple('DC45', [('cf75', Any), ('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC45({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC45) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC46(D18, NamedTuple('DC46', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC46'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC46)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC47(D18, NamedTuple('DC47', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC49(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC49(D19, NamedTuple('DC49', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC52()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)

class D20_DC52(D20, NamedTuple('DC52', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC51(D20, NamedTuple('DC51', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)

class D21_DC53(D21, NamedTuple('DC53', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC55(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)

class D22_DC55(D22, NamedTuple('DC55', [('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({self.cf86.VerbatimString(True)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC54(D22, NamedTuple('DC54', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC56(D22, NamedTuple('DC56', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC58()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)

class D23_DC58(D23, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC59(D23, NamedTuple('DC59', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC60(D24, NamedTuple('DC60', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC62(None, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D25_DC61)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D25_DC64)

class D25_DC62(D25, NamedTuple('DC62', [('cf94', Any), ('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC63(D25, NamedTuple('DC63', [('cf97', Any), ('cf98', Any), ('cf99', Any), ('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC61(D25, NamedTuple('DC61', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC61({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC61) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC64(D25, NamedTuple('DC64', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC64({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC64) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC66(_dafny.Map({}), False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D26_DC66)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D26_DC65)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D26_DC67)

class D26_DC66(D26, NamedTuple('DC66', [('cf103', Any), ('cf104', Any), ('cf105', Any), ('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC66({_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC66) and self.cf103 == __o.cf103 and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC65(D26, NamedTuple('DC65', [('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC65({_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC65) and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC67(D26, NamedTuple('DC67', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC67({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC67) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC69(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D27_DC69)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D27_DC68)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D27_DC70)

class D27_DC69(D27, NamedTuple('DC69', [('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC69({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC69) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC68(D27, NamedTuple('DC68', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC68({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC68) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC70(D27, NamedTuple('DC70', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC70({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC70) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    def m1(self, p0, p1, globalState):
        pass


class T2:
    pass

class GlobalState:
    def  __init__(self):
        self._f0: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0):
        (self)._f0 = f0

    @property
    def f0(self):
        return self._f0

class C0:
    def  __init__(self):
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f14):
        (self)._f14 = f14

    def fm29(self, p0, p1, p2, globalState):
        source12_ = D1_DC3(True, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False]), _dafny.MultiSet([True])])): _dafny.SeqWithoutIsStrInference([False])})), 337)
        if source12_.is_DC3:
            d_406___mcc_h0_ = source12_.cf3
            d_407___mcc_h1_ = source12_.cf4
            d_408___mcc_h2_ = source12_.cf5
            d_409_cf5_ = d_408___mcc_h2_
            d_410_cf4_ = d_407___mcc_h1_
            d_411_cf3_ = d_406___mcc_h0_
            return _dafny.CodePoint('u')
        elif source12_.is_DC4:
            d_412___mcc_h3_ = source12_.cf6
            d_413___mcc_h4_ = source12_.cf7
            d_414_cf7_ = d_413___mcc_h4_
            d_415_cf6_ = d_412___mcc_h3_
            return _dafny.CodePoint('f')
        elif True:
            d_416___mcc_h5_ = source12_.cf2
            d_417_cf2_ = d_416___mcc_h5_
            def iife54_():
                coll30_ = _dafny.Set()
                compr_30_: int
                for compr_30_ in (_dafny.Map({201: False})).keys.Elements:
                    d_418_v0_: int = compr_30_
                    if (d_418_v0_) in (_dafny.Map({201: False})):
                        coll30_ = coll30_.union(_dafny.Set([(d_418_v0_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_419_i1_ in range(default__.abs(34))])): True}))})) for d_420_i0_ in range(default__.abs(723))])))]))
                return _dafny.Set(coll30_)
            return (D6_DC18(len(iife54_()
), 226, False, _dafny.CodePoint('w'), 571)).cf32

    def fm30(self, globalState):
        return not (True) or (((_dafny.MultiSet([238, len(_dafny.Map({673: len(_dafny.Map({False: True}))}))])).cardinality) != ((0) - (len(_dafny.Map({-31: -428})))))

    @property
    def f14(self):
        return self._f14

class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm28(self, p0, p1, globalState):
        return ((0) - (len(_dafny.Map({_dafny.CodePoint('g'): False})))) + (len(_dafny.Map({(0) - (-498): not(not(False))})))

    def m12(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: D5 = D5.default()()
        d_421_v0_: _dafny.Array
        def lambda16_(d_422_p0_):
            def lambda17_(d_423_i0_):
                return d_422_p0_

            return lambda17_

        init8_ = lambda16_(p0)
        nw43_ = _dafny.Array(None, 3)
        for i0_8_ in range(nw43_.length(0)):
            nw43_[i0_8_] = init8_(i0_8_)
        d_421_v0_ = nw43_
        d_424_v1_: _dafny.Map
        d_424_v1_ = _dafny.Map({d_421_v0_: p0})
        d_425_v2_: bool
        d_426_v3_: bool
        d_427_v4_: _dafny.Seq
        d_428_v5_: _dafny.Seq
        out13_: bool
        out14_: bool
        out15_: _dafny.Seq
        out16_: _dafny.Seq
        out13_, out14_, out15_, out16_ = default__.m0(d_424_v1_, globalState)
        d_425_v2_ = out13_
        d_426_v3_ = out14_
        d_427_v4_ = out15_
        d_428_v5_ = out16_
        d_429_i1_: int
        d_429_i1_ = 0
        with _dafny.label("3"):
            while p0:
                with _dafny.c_label("3"):
                    if (d_429_i1_) >= (100):
                        raise _dafny.Break("3")
                    d_429_i1_ = (d_429_i1_) + (1)
                    d_430_v6_: int
                    d_430_v6_ = 534
                    r0 = d_430_v6_
                    d_431_v7_: _dafny.Seq
                    d_431_v7_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odqlsiwfo")))])
                    d_432_v8_: _dafny.MultiSet
                    d_432_v8_ = _dafny.MultiSet([p0])
                    d_433_v9_: _dafny.Seq
                    d_433_v9_ = _dafny.SeqWithoutIsStrInference([d_432_v8_])
                    d_434_v10_: str
                    d_434_v10_ = _dafny.CodePoint('g')
                    d_435_v11_: _dafny.MultiSet
                    d_435_v11_ = _dafny.MultiSet([d_434_v10_])
                    d_436_v13_: _dafny.Array
                    nw44_ = _dafny.Array(None, 21)
                    nw44_[int(0)] = len(d_431_v7_)
                    nw44_[int(1)] = (0) - (d_430_v6_)
                    nw44_[int(2)] = d_430_v6_
                    nw44_[int(3)] = d_430_v6_
                    nw44_[int(4)] = ((d_433_v9_)[default__.safeIndex(d_430_v6_, len(d_433_v9_))]).cardinality
                    nw44_[int(5)] = d_430_v6_
                    nw44_[int(6)] = d_430_v6_
                    nw44_[int(7)] = 693
                    nw44_[int(8)] = (0) - ((d_431_v7_)[default__.safeIndex(d_430_v6_, len(d_431_v7_))])
                    nw44_[int(9)] = (d_432_v8_).cardinality
                    nw44_[int(10)] = d_430_v6_
                    def iife55_():
                        coll31_ = _dafny.Set()
                        compr_31_: int
                        for compr_31_ in _dafny.IntegerRange(666, 789):
                            d_437_v12_: int = compr_31_
                            if ((666) <= (d_437_v12_)) and ((d_437_v12_) < (789)):
                                coll31_ = coll31_.union(_dafny.Set([default__.safeModuloInt(d_437_v12_, len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spwpy"))})))]))
                        return _dafny.Set(coll31_)
                    nw44_[int(11)] = ((d_435_v11_)[d_434_v10_] if (d_434_v10_) in (d_435_v11_) else len(iife55_()
                    ))
                    nw44_[int(12)] = d_430_v6_
                    nw44_[int(13)] = 347
                    nw44_[int(14)] = (0) - (d_430_v6_)
                    nw44_[int(15)] = d_430_v6_
                    nw44_[int(16)] = 367
                    nw44_[int(17)] = d_430_v6_
                    nw44_[int(18)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ledgplr")))
                    nw44_[int(19)] = d_430_v6_
                    nw44_[int(20)] = (d_432_v8_).cardinality
                    d_436_v13_ = nw44_
                    d_438_v14_: C0
                    nw45_ = C0()
                    nw45_.ctor__(d_436_v13_)
                    d_438_v14_ = nw45_
                    d_439_v15_: _dafny.MultiSet
                    d_439_v15_ = _dafny.MultiSet([d_430_v6_, d_430_v6_])
                    d_440_v16_: _dafny.Set
                    d_440_v16_ = _dafny.Set({d_439_v15_})
                    if (938) <= ((len(d_440_v16_)) - (default__.fm0(globalState))):
                        d_441_v17_: C0
                        nw46_ = C0()
                        nw46_.ctor__((d_438_v14_).f14)
                        d_441_v17_ = nw46_
                        d_442_v18_: _dafny.Seq
                        d_442_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hv"))
                        index55_ = default__.safeIndex(848, (d_421_v0_).length(0))
                        (d_421_v0_)[index55_] = (d_434_v10_) in (d_442_v18_)
                        index56_ = default__.safeIndex(848, (d_421_v0_).length(0))
                        (d_421_v0_)[index56_] = d_426_v3_
                        index57_ = default__.safeIndex(848, (d_421_v0_).length(0))
                        (d_421_v0_)[index57_] = True
                        d_434_v10_ = d_434_v10_
                        d_443_v19_: C0
                        nw47_ = C0()
                        nw47_.ctor__((d_441_v17_).f14)
                        d_443_v19_ = nw47_
                    elif True:
                        d_430_v6_ = (d_430_v6_) * (default__.safeModuloInt(d_430_v6_, d_430_v6_))
                        d_431_v7_ = d_431_v7_
                        d_444_v20_: C0
                        nw48_ = C0()
                        nw48_.ctor__((d_438_v14_).f14)
                        d_444_v20_ = nw48_
                        d_445_v21_: _dafny.Seq
                        d_445_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxrlpmtg"))
                        d_446_v22_: _dafny.Map
                        d_446_v22_ = _dafny.Map({p0: d_445_v21_})
                        d_445_v21_ = (((d_446_v22_)[d_426_v3_] if (d_426_v3_) in (d_446_v22_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_447_i2_ in range(default__.abs(516))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_448_i3_ in range(default__.abs(655))]))
                        d_449_v23_: _dafny.Map
                        d_449_v23_ = _dafny.Map({d_426_v3_: d_431_v7_})
                        d_450_v24_: _dafny.Array
                        nw49_ = _dafny.Array(None, 13)
                        nw49_[int(0)] = d_431_v7_
                        nw49_[int(1)] = (d_431_v7_) + (default__.fm31(d_430_v6_, d_425_v2_, 917, d_434_v10_, globalState))
                        nw49_[int(2)] = _dafny.SeqWithoutIsStrInference([760])
                        nw49_[int(3)] = ((d_449_v23_)[d_425_v2_] if (d_425_v2_) in (d_449_v23_) else d_431_v7_)
                        nw49_[int(4)] = _dafny.SeqWithoutIsStrInference([d_430_v6_])
                        nw49_[int(5)] = d_431_v7_
                        nw49_[int(6)] = d_431_v7_
                        nw49_[int(7)] = ((d_449_v23_)[d_425_v2_] if (d_425_v2_) in (d_449_v23_) else (d_431_v7_).set(default__.safeIndex(d_430_v6_, len(d_431_v7_)), (0) - (d_430_v6_)))
                        nw49_[int(8)] = d_431_v7_
                        nw49_[int(9)] = d_431_v7_
                        nw49_[int(10)] = _dafny.SeqWithoutIsStrInference([778, d_430_v6_])
                        nw49_[int(11)] = d_431_v7_
                        nw49_[int(12)] = _dafny.SeqWithoutIsStrInference([d_430_v6_ for d_451_i4_ in range(default__.abs(-269))])
                        d_450_v24_ = nw49_
                        index58_ = default__.safeIndex(46, (d_450_v24_).length(0))
                        (d_450_v24_)[index58_] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: d_426_v3_})) for d_452_i5_ in range(default__.abs(338))])
                        index59_ = default__.safeIndex(46, (d_450_v24_).length(0))
                        (d_450_v24_)[index59_] = d_431_v7_
                    if d_426_v3_:
                        d_426_v3_ = not(d_426_v3_)
                        d_436_v13_ = (d_438_v14_).f14
                        d_421_v0_ = d_421_v0_
                        r0 = (0) - (default__.safeModuloInt(d_430_v6_, (d_430_v6_ if d_426_v3_ else default__.fm0(globalState))))
                        d_453_v25_: _dafny.Seq
                        d_453_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqxkm"))
                        d_439_v15_ = (default__.fm8(904, d_453_v25_, globalState)).set(len(d_453_v25_), default__.abs(d_430_v6_))
                    elif True:
                        d_454_v26_: _dafny.Map
                        d_454_v26_ = _dafny.Map({p0: p0})
                        d_454_v26_ = default__.fm9(globalState)
                        d_430_v6_ = (d_430_v6_) * (d_430_v6_)
                        d_455_v27_: _dafny.Seq
                        d_455_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                        d_456_v28_: _dafny.Map
                        d_456_v28_ = _dafny.Map({d_421_v0_: len((default__.fm2(False, globalState)) + (d_455_v27_))})
                        d_457_v29_: D1
                        d_457_v29_ = D1_DC2(d_455_v27_)
                        rhs64_ = d_456_v28_
                        rhs65_ = (((d_433_v9_)[default__.safeIndex(d_430_v6_, len(d_433_v9_))]).cardinality) > ((0) - (d_430_v6_))
                        rhs66_ = ((d_455_v27_) + ((d_457_v29_).cf2)) + (d_455_v27_)
                        d_456_v28_ = rhs64_
                        r2 = rhs65_
                        d_455_v27_ = rhs66_
                        d_458_v30_: _dafny.Seq
                        d_458_v30_ = _dafny.SeqWithoutIsStrInference([d_455_v27_, d_455_v27_])
                        d_459_v31_: _dafny.Map
                        d_459_v31_ = _dafny.Map({d_430_v6_: default__.fm2(d_426_v3_, globalState)})
                        d_460_v32_: _dafny.Map
                        d_460_v32_ = _dafny.Map({d_438_v14_: _dafny.SeqWithoutIsStrInference([d_434_v10_ for d_461_i7_ in range(default__.abs(180))])})
                        d_462_v33_: _dafny.Array
                        nw50_ = _dafny.Array(None, 24)
                        nw50_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsprls"))
                        nw50_[int(1)] = (d_455_v27_) + ((d_458_v30_)[default__.safeIndex((0) - (d_430_v6_), len(d_458_v30_))])
                        nw50_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fviebtams"))) + (d_455_v27_)
                        nw50_[int(3)] = (d_455_v27_) + (_dafny.SeqWithoutIsStrInference([d_434_v10_ for d_463_i6_ in range(default__.abs(-198))]))
                        nw50_[int(4)] = ((d_459_v31_)[(0) - (d_430_v6_)] if ((0) - (d_430_v6_)) in (d_459_v31_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqw")))
                        nw50_[int(5)] = d_455_v27_
                        nw50_[int(6)] = d_455_v27_
                        nw50_[int(7)] = ((d_460_v32_)[d_438_v14_] if (d_438_v14_) in (d_460_v32_) else d_455_v27_)
                        nw50_[int(8)] = d_455_v27_
                        nw50_[int(9)] = d_455_v27_
                        nw50_[int(10)] = d_455_v27_
                        nw50_[int(11)] = d_455_v27_
                        nw50_[int(12)] = d_455_v27_
                        nw50_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_464_i8_ in range(default__.abs(263))])
                        nw50_[int(14)] = d_455_v27_
                        nw50_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfx"))) + (d_455_v27_)
                        nw50_[int(16)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_465_i9_ in range(default__.abs(137))])
                        nw50_[int(17)] = (d_455_v27_) + (d_455_v27_)
                        nw50_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obxuhqq"))
                        nw50_[int(19)] = d_455_v27_
                        nw50_[int(20)] = d_455_v27_
                        nw50_[int(21)] = (d_455_v27_) + ((((d_459_v31_)[d_430_v6_] if (d_430_v6_) in (d_459_v31_) else _dafny.SeqWithoutIsStrInference([d_434_v10_ for d_466_i10_ in range(default__.abs(995))]))).set(default__.safeIndex(d_430_v6_, len(((d_459_v31_)[d_430_v6_] if (d_430_v6_) in (d_459_v31_) else _dafny.SeqWithoutIsStrInference([d_434_v10_ for d_467_i10_ in range(default__.abs(995))])))), d_434_v10_))
                        nw50_[int(22)] = (default__.fm2(d_426_v3_, globalState)) + (d_455_v27_)
                        nw50_[int(23)] = d_455_v27_
                        d_462_v33_ = nw50_
                        index60_ = default__.safeIndex(904, (d_462_v33_).length(0))
                        (d_462_v33_)[index60_] = d_455_v27_
                        d_468_v34_: _dafny.Set
                        d_468_v34_ = _dafny.Set({_dafny.MultiSet([True]), d_432_v8_})
                        index61_ = default__.safeIndex(904, (d_462_v33_).length(0))
                        (d_462_v33_)[index61_] = (d_455_v27_).set(default__.safeIndex(len(d_468_v34_), len(d_455_v27_)), _dafny.CodePoint('s'))
                        d_469_v35_: _dafny.Map
                        d_469_v35_ = _dafny.Map({(d_462_v33_)[default__.safeIndex(904, (d_462_v33_).length(0))]: d_430_v6_})
                        d_470_v36_: _dafny.Map
                        d_470_v36_ = _dafny.Map({(d_425_v2_ if d_426_v3_ else True): d_469_v35_})
                        d_470_v36_ = (d_470_v36_).set((568) != (len(d_455_v27_)), d_469_v35_)
                    pass
            pass
        d_471_v37_: _dafny.Seq
        d_471_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipd"))
        d_472_v38_: int
        d_472_v38_ = 611
        d_473_v39_: str
        d_473_v39_ = _dafny.CodePoint('i')
        d_471_v37_ = (((default__.fm2(d_425_v2_, globalState)) + (d_471_v37_)) + ((d_471_v37_) + (d_471_v37_))).set(default__.safeIndex(d_472_v38_, len(((default__.fm2(d_425_v2_, globalState)) + (d_471_v37_)) + ((d_471_v37_) + (d_471_v37_)))), (d_473_v39_ if p0 else d_473_v39_))
        d_474_i11_: int
        d_474_i11_ = 0
        with _dafny.label("4"):
            while p0:
                with _dafny.c_label("4"):
                    if (d_474_i11_) >= (100):
                        raise _dafny.Break("4")
                    d_474_i11_ = (d_474_i11_) + (1)
                    d_475_v40_: _dafny.Seq
                    d_475_v40_ = _dafny.SeqWithoutIsStrInference([d_472_v38_])
                    rhs67_ = d_475_v40_
                    rhs68_ = ((d_472_v38_) - (d_472_v38_) if False else d_472_v38_)
                    d_475_v40_ = rhs67_
                    d_472_v38_ = rhs68_
                    d_476_v41_: _dafny.Map
                    d_476_v41_ = _dafny.Map({True: d_471_v37_})
                    d_476_v41_ = (d_476_v41_).set((d_427_v4_)[default__.safeIndex(d_472_v38_, len(d_427_v4_))], d_471_v37_)
                    d_477_v42_: _dafny.MultiSet
                    d_477_v42_ = _dafny.MultiSet([d_421_v0_, d_421_v0_])
                    rhs69_ = d_475_v40_
                    rhs70_ = (((_dafny.MultiSet([d_421_v0_, d_421_v0_, d_421_v0_, d_421_v0_])) - (d_477_v42_)) - ((d_477_v42_) | (d_477_v42_))).cardinality
                    d_475_v40_ = rhs69_
                    d_472_v38_ = rhs70_
                    d_478_v43_: _dafny.Array
                    def lambda18_(d_479_v38_):
                        def lambda19_(d_480_i12_):
                            return (d_480_i12_) + (d_479_v38_)

                        return lambda19_

                    init9_ = lambda18_(d_472_v38_)
                    nw51_ = _dafny.Array(None, 8)
                    for i0_9_ in range(nw51_.length(0)):
                        nw51_[i0_9_] = init9_(i0_9_)
                    d_478_v43_ = nw51_
                    d_481_v44_: C0
                    nw52_ = C0()
                    nw52_.ctor__(d_478_v43_)
                    d_481_v44_ = nw52_
                    pass
            pass
        d_482_v45_: _dafny.Map
        d_482_v45_ = _dafny.Map({d_472_v38_: 368})
        hi2_ = len(d_482_v45_)
        for d_483_i13_ in range((0) - ((d_472_v38_ if False else d_472_v38_)), hi2_):
            d_484_v46_: _dafny.Set
            d_484_v46_ = _dafny.Set({d_421_v0_})
            d_485_v47_: _dafny.Array
            nw53_ = _dafny.Array(None, 7)
            nw53_[int(0)] = (0) - (len(d_484_v46_))
            nw53_[int(1)] = 326
            nw53_[int(2)] = default__.fm0(globalState)
            nw53_[int(3)] = d_472_v38_
            nw53_[int(4)] = d_472_v38_
            nw53_[int(5)] = d_472_v38_
            nw53_[int(6)] = -701
            d_485_v47_ = nw53_
            d_486_v48_: C0
            nw54_ = C0()
            nw54_.ctor__(d_485_v47_)
            d_486_v48_ = nw54_
            d_487_v49_: _dafny.Array
            nw55_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_487_v49_ = nw55_
            d_487_v49_ = d_487_v49_
            d_488_v50_: _dafny.MultiSet
            d_488_v50_ = _dafny.MultiSet([d_473_v39_])
            d_489_v51_: _dafny.Map
            d_489_v51_ = _dafny.Map({d_426_v3_: (d_488_v50_).cardinality})
            d_489_v51_ = (d_489_v51_).set(not(d_425_v2_), d_472_v38_)
            if d_425_v2_:
                d_490_v52_: _dafny.Map
                d_490_v52_ = _dafny.Map({d_483_i13_: d_473_v39_})
                d_491_v53_: D3
                d_491_v53_ = D3_DC9((0) - (len(d_490_v52_)), d_426_v3_)
                pat_let_tv16_ = d_471_v37_
                pat_let_tv17_ = globalState
                def iife56_(_pat_let12_0):
                    def iife57_(d_492_dt__update__tmp_h0_):
                        def iife58_(_pat_let13_0):
                            def iife59_(d_493_dt__update_hcf16_h0_):
                                return D3_DC9((d_492_dt__update__tmp_h0_).cf15, d_493_dt__update_hcf16_h0_)
                            return iife59_(_pat_let13_0)
                        return iife58_(default__.fm3(len(pat_let_tv16_), pat_let_tv17_))
                    return iife57_(_pat_let12_0)
                d_425_v2_ = (d_491_v53_) in ((default__.fm32(d_483_i13_, 420, d_472_v38_, d_471_v37_, globalState)).set(iife56_(d_491_v53_), d_483_i13_))
                d_494_v54_: _dafny.Map
                d_494_v54_ = _dafny.Map({d_483_i13_: p0})
                d_495_v55_: D1
                d_495_v55_ = D1_DC4(((d_494_v54_)[d_472_v38_] if (d_472_v38_) in (d_494_v54_) else default__.fm3(d_472_v38_, globalState)), d_426_v3_)
                d_495_v55_ = d_495_v55_
                d_472_v38_ = d_472_v38_
                r2 = d_425_v2_
                d_496_v56_: _dafny.Array
                nw56_ = _dafny.Array(None, 20)
                nw56_[int(0)] = d_421_v0_
                nw56_[int(1)] = d_421_v0_
                nw56_[int(2)] = d_421_v0_
                nw56_[int(3)] = d_421_v0_
                nw56_[int(4)] = d_421_v0_
                nw56_[int(5)] = d_421_v0_
                nw56_[int(6)] = d_421_v0_
                nw56_[int(7)] = d_421_v0_
                nw56_[int(8)] = d_421_v0_
                nw56_[int(9)] = d_421_v0_
                nw56_[int(10)] = d_421_v0_
                nw56_[int(11)] = d_421_v0_
                nw56_[int(12)] = d_421_v0_
                nw56_[int(13)] = d_421_v0_
                nw56_[int(14)] = d_421_v0_
                nw56_[int(15)] = d_421_v0_
                nw56_[int(16)] = d_421_v0_
                nw56_[int(17)] = d_421_v0_
                nw56_[int(18)] = d_421_v0_
                nw56_[int(19)] = d_421_v0_
                d_496_v56_ = nw56_
                d_496_v56_ = d_496_v56_
            elif True:
                def lambda20_(d_497_v39_):
                    def lambda21_(d_498_i14_):
                        return d_497_v39_

                    return lambda21_

                init10_ = lambda20_(d_473_v39_)
                nw57_ = _dafny.Array(None, 11)
                for i0_10_ in range(nw57_.length(0)):
                    nw57_[i0_10_] = init10_(i0_10_)
                d_487_v49_ = nw57_
                d_499_v57_: _dafny.Seq
                d_499_v57_ = _dafny.SeqWithoutIsStrInference([len(d_471_v37_)])
                index62_ = default__.safeIndex(374, ((d_486_v48_).f14).length(0))
                ((d_486_v48_).f14)[index62_] = (self).fm28(p0, len(d_499_v57_), globalState)
                d_500_v58_: D6
                d_500_v58_ = D6_DC17(p0)
                index63_ = default__.safeIndex(374, ((d_486_v48_).f14).length(0))
                rhs71_ = d_482_v45_
                rhs72_ = (-354) - (len((d_471_v37_) + (d_471_v37_)))
                rhs73_ = d_426_v3_
                rhs74_ = d_500_v58_
                lhs32_ = (d_486_v48_).f14
                lhs33_ = default__.safeIndex(374, ((d_486_v48_).f14).length(0))
                d_482_v45_ = rhs71_
                lhs32_[lhs33_] = rhs72_
                d_426_v3_ = rhs73_
                d_500_v58_ = rhs74_
                d_425_v2_ = d_426_v3_
                d_501_v59_: _dafny.Set
                def iife60_(_pat_let14_0):
                    def iife61_(d_502_dt__update__tmp_h1_):
                        def iife62_(_pat_let15_0):
                            def iife63_(d_503_dt__update_hcf28_h0_):
                                return D6_DC17(d_503_dt__update_hcf28_h0_)
                            return iife63_(_pat_let15_0)
                        return iife62_(False)
                    return iife61_(_pat_let14_0)
                d_501_v59_ = _dafny.Set({iife60_(d_500_v58_), d_500_v58_, d_500_v58_, d_500_v58_})
                d_501_v59_ = ((d_501_v59_) - (d_501_v59_)) - (d_501_v59_)
                r0 = d_472_v38_
        d_504_v60_: _dafny.MultiSet
        d_504_v60_ = _dafny.MultiSet([not(d_426_v3_), d_426_v3_, p0])
        hi3_ = ((d_504_v60_)[d_425_v2_] if (d_425_v2_) in (d_504_v60_) else d_472_v38_)
        for d_505_i15_ in range(d_472_v38_, hi3_):
            r0 = d_505_i15_
            d_506_v61_: D0
            d_506_v61_ = D0_DC1(334)
            source13_ = d_506_v61_
            if source13_.is_DC1:
                d_507___mcc_h0_ = source13_.cf1
                d_508_cf1_ = d_507___mcc_h0_
                r0 = default__.safeModuloInt(d_472_v38_, (d_505_i15_) + (d_508_cf1_))
                d_509_v62_: _dafny.Array
                nw58_ = _dafny.Array(int(0), 13)
                d_509_v62_ = nw58_
                d_510_v63_: _dafny.Map
                d_510_v63_ = _dafny.Map({(d_509_v62_ if d_425_v2_ else d_509_v62_): d_425_v2_})
                d_510_v63_ = (d_510_v63_).set(d_509_v62_, False)
                d_473_v39_ = d_473_v39_
                d_511_v64_: _dafny.Set
                d_511_v64_ = _dafny.Set({d_472_v38_, d_472_v38_, d_505_i15_, d_505_i15_})
                d_512_v65_: _dafny.Seq
                d_512_v65_ = _dafny.SeqWithoutIsStrInference([d_472_v38_])
                d_513_v66_: D3
                d_513_v66_ = D3_DC9(d_472_v38_, d_426_v3_)
                d_514_v67_: _dafny.Map
                d_514_v67_ = _dafny.Map({len(d_512_v65_): d_513_v66_})
                d_515_v69_: _dafny.Seq
                def iife64_():
                    coll32_ = _dafny.Set()
                    compr_32_: int
                    for compr_32_ in (d_514_v67_).keys.Elements:
                        d_516_v68_: int = compr_32_
                        if (d_516_v68_) in (d_514_v67_):
                            coll32_ = coll32_.union(_dafny.Set([(d_516_v68_) - (96)]))
                    return _dafny.Set(coll32_)
                d_515_v69_ = _dafny.SeqWithoutIsStrInference([d_511_v64_, iife64_()
                , d_511_v64_])
                d_517_v70_: _dafny.Set
                d_517_v70_ = _dafny.Set({d_511_v64_, d_511_v64_, default__.fm33(len(_dafny.SeqWithoutIsStrInference([d_426_v3_, not(d_425_v2_)])), globalState)})
                d_518_v71_: _dafny.Set
                d_518_v71_ = _dafny.Set({len((d_515_v69_)[default__.safeIndex(d_472_v38_, len(d_515_v69_))]), len(d_517_v70_)})
                r2 = (d_472_v38_) in (d_518_v71_)
            elif True:
                d_519___mcc_h1_ = source13_.cf0
                d_520_cf0_ = d_519___mcc_h1_
                d_521_v72_: _dafny.Map
                d_521_v72_ = _dafny.Map({d_425_v2_: p0})
                d_521_v72_ = (d_521_v72_).set(p0, not(not((_dafny.SeqWithoutIsStrInference([d_426_v3_, d_426_v3_])) <= (d_427_v4_))))
                d_522_v73_: _dafny.Seq
                d_522_v73_ = _dafny.SeqWithoutIsStrInference([d_471_v37_, d_471_v37_, d_471_v37_, d_471_v37_, d_471_v37_])
                d_523_v74_: _dafny.Map
                d_523_v74_ = _dafny.Map({(d_427_v4_)[default__.safeIndex(d_520_cf0_, len(d_427_v4_))]: d_421_v0_})
                d_524_v75_: _dafny.Array
                nw59_ = _dafny.Array(None, 24)
                nw59_[int(0)] = d_520_cf0_
                nw59_[int(1)] = d_472_v38_
                nw59_[int(2)] = default__.safeDivisionInt(d_505_i15_, len(d_522_v73_))
                nw59_[int(3)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))) + (d_520_cf0_)
                nw59_[int(4)] = (d_505_i15_) * (d_505_i15_)
                nw59_[int(5)] = len((_dafny.SeqWithoutIsStrInference([d_473_v39_ for d_525_i16_ in range(default__.abs(988))])) + (d_471_v37_))
                nw59_[int(6)] = len(default__.fm4(d_472_v38_, d_426_v3_, globalState))
                nw59_[int(7)] = len(d_471_v37_)
                nw59_[int(8)] = d_472_v38_
                nw59_[int(9)] = (d_520_cf0_) - (len(d_427_v4_))
                nw59_[int(10)] = d_505_i15_
                nw59_[int(11)] = d_505_i15_
                nw59_[int(12)] = (d_472_v38_ if d_425_v2_ else d_520_cf0_)
                nw59_[int(13)] = d_472_v38_
                nw59_[int(14)] = (0) - (d_505_i15_)
                nw59_[int(15)] = d_472_v38_
                nw59_[int(16)] = default__.safeModuloInt(len(d_521_v72_), d_505_i15_)
                nw59_[int(17)] = len(d_523_v74_)
                nw59_[int(18)] = d_505_i15_
                nw59_[int(19)] = (_dafny.MultiSet([(0) - (len(d_471_v37_)), d_472_v38_])).cardinality
                nw59_[int(20)] = (0) - (d_505_i15_)
                nw59_[int(21)] = d_472_v38_
                nw59_[int(22)] = d_472_v38_
                nw59_[int(23)] = d_520_cf0_
                d_524_v75_ = nw59_
                index64_ = default__.safeIndex(699, (d_524_v75_).length(0))
                (d_524_v75_)[index64_] = d_505_i15_
                index65_ = default__.safeIndex(699, (d_524_v75_).length(0))
                (d_524_v75_)[index65_] = len((((d_471_v37_) + (d_471_v37_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp")))).set(default__.safeIndex(default__.fm0(globalState), len(((d_471_v37_) + (d_471_v37_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp"))))), d_473_v39_))
                d_526_v76_: _dafny.MultiSet
                d_526_v76_ = _dafny.MultiSet([(self).fm28(d_425_v2_, -524, globalState)])
                d_527_v77_: _dafny.Seq
                d_527_v77_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_473_v39_: ((d_526_v76_)[d_472_v38_] if (d_472_v38_) in (d_526_v76_) else d_520_cf0_)}))])
                d_528_v78_: _dafny.MultiSet
                d_528_v78_ = _dafny.MultiSet([d_520_cf0_, len(d_527_v77_), d_505_i15_])
                r2 = ((d_528_v78_) | (d_526_v76_)).ispropersubset(_dafny.MultiSet([(d_524_v75_)[default__.safeIndex(699, (d_524_v75_).length(0))]]))
                d_529_v79_: _dafny.Set
                d_529_v79_ = _dafny.Set({d_425_v2_})
                rhs75_ = d_505_i15_
                rhs76_ = default__.fm3((self).fm28(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihpnbc"))), globalState), globalState)
                rhs77_ = ((d_529_v79_ if d_426_v3_ else _dafny.Set({d_426_v3_, not(False)}))) != ((_dafny.Set({default__.fm3((d_524_v75_)[default__.safeIndex(699, (d_524_v75_).length(0))], globalState)})) - (d_529_v79_))
                r0 = rhs75_
                d_425_v2_ = rhs76_
                d_425_v2_ = rhs77_
            d_530_v80_: _dafny.Array
            nw60_ = _dafny.Array(int(0), 5)
            d_530_v80_ = nw60_
            index66_ = default__.safeIndex(915, (d_530_v80_).length(0))
            (d_530_v80_)[index66_] = (len(d_471_v37_)) - (d_472_v38_)
            index67_ = default__.safeIndex(915, (d_530_v80_).length(0))
            (d_530_v80_)[index67_] = default__.safeDivisionInt(((d_504_v60_)[d_426_v3_] if (d_426_v3_) in (d_504_v60_) else d_472_v38_), d_472_v38_)
            if False:
                index68_ = default__.safeIndex(277, (d_421_v0_).length(0))
                (d_421_v0_)[index68_] = d_426_v3_
                index69_ = default__.safeIndex(277, (d_421_v0_).length(0))
                (d_421_v0_)[index69_] = p0
                d_531_v81_: _dafny.MultiSet
                d_531_v81_ = _dafny.MultiSet([d_505_i15_])
                r1 = default__.fm3((d_531_v81_).cardinality, globalState)
                d_532_v82_: _dafny.Array
                nw61_ = _dafny.Array(_dafny.Set({}), 28)
                d_532_v82_ = nw61_
                d_533_v83_: _dafny.Array
                nw62_ = _dafny.Array(None, 3)
                nw62_[int(0)] = d_473_v39_
                nw62_[int(1)] = d_473_v39_
                nw62_[int(2)] = d_473_v39_
                d_533_v83_ = nw62_
                d_534_v84_: _dafny.Set
                d_534_v84_ = _dafny.Set({d_533_v83_, d_533_v83_, d_533_v83_})
                index70_ = default__.safeIndex(374, (d_532_v82_).length(0))
                (d_532_v82_)[index70_] = (d_534_v84_) | (d_534_v84_)
                index71_ = default__.safeIndex(374, (d_532_v82_).length(0))
                (d_532_v82_)[index71_] = d_534_v84_
                d_535_v85_: _dafny.Set
                d_535_v85_ = _dafny.Set({not(not(not((d_421_v0_)[default__.safeIndex(277, (d_421_v0_).length(0))]))), not(d_426_v3_), d_425_v2_})
                d_535_v85_ = ((D3_DC8(d_535_v85_)).cf14).intersection(default__.fm6(d_426_v3_, d_473_v39_, globalState))
                d_536_v86_: _dafny.Seq
                d_536_v86_ = _dafny.SeqWithoutIsStrInference([(0) - (d_505_i15_), (d_530_v80_)[default__.safeIndex(915, (d_530_v80_).length(0))], len(d_427_v4_)])
                index72_ = default__.safeIndex(915, (d_530_v80_).length(0))
                (d_530_v80_)[index72_] = (len(d_536_v86_)) - (715)
            elif True:
                d_530_v80_ = d_530_v80_
                d_537_v87_: _dafny.MultiSet
                d_537_v87_ = _dafny.MultiSet([d_530_v80_, d_530_v80_, d_530_v80_, d_530_v80_, d_530_v80_])
                d_537_v87_ = d_537_v87_
                d_471_v37_ = _dafny.SeqWithoutIsStrInference([d_473_v39_ for d_538_i17_ in range(default__.abs(316))])
                d_539_v88_: _dafny.Seq
                d_539_v88_ = _dafny.SeqWithoutIsStrInference([d_471_v37_])
                d_540_v89_: _dafny.Map
                d_540_v89_ = _dafny.Map({d_472_v38_: (d_539_v88_) + (d_539_v88_)})
                d_541_v90_: _dafny.Map
                d_541_v90_ = _dafny.Map({D3_DC8(_dafny.Set({p0})): d_539_v88_})
                d_542_v91_: D3
                d_542_v91_ = D3_DC8(_dafny.Set({d_426_v3_}))
                d_540_v89_ = (d_540_v89_).set(len((_dafny.SeqWithoutIsStrInference([d_426_v3_])) + (d_427_v4_)), (((d_541_v90_)[d_542_v91_] if (d_542_v91_) in (d_541_v90_) else d_539_v88_) if False else d_539_v88_))
                r2 = d_425_v2_
        r0 = default__.safeModuloInt(default__.safeModuloInt(d_472_v38_, d_472_v38_), default__.safeModuloInt(d_472_v38_, d_472_v38_))
        d_543_v92_: _dafny.Set
        d_543_v92_ = _dafny.Set({d_426_v3_})
        r1 = ((default__.fm6(d_425_v2_, _dafny.CodePoint('y'), globalState) if d_425_v2_ else d_543_v92_)).ispropersubset(_dafny.Set({True, d_425_v2_}))
        r2 = True
        d_544_v93_: D5
        d_544_v93_ = D5_DC15(d_472_v38_, p0, d_421_v0_)
        r3 = d_544_v93_
        return r0, r1, r2, r3


class C2(T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    def ctor__(self, f13, f3, f1, f2):
        (self)._f13 = f13
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm12(self, p0, p1, p2, p3, globalState):
        return (0) - ((self).f13)

    def fm11(self, p0, globalState):
        return self.f3

    def fm27(self, p0, p1, p2, globalState):
        return (default__.safeModuloInt((self).f13, (self).f13)) - (default__.safeDivisionInt((self).f13, (0) - ((self).f13)))

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        d_545_v0_: _dafny.Array
        nw63_ = _dafny.Array(False, 2)
        d_545_v0_ = nw63_
        index73_ = default__.safeIndex(303, (d_545_v0_).length(0))
        (d_545_v0_)[index73_] = default__.fm3((self).f13, globalState)
        index74_ = default__.safeIndex(303, (d_545_v0_).length(0))
        (d_545_v0_)[index74_] = not (self.f3) or (self.f3)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_545_v0_).length(0)):
            d_546_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_546_i0_)) and ((d_546_i0_) < ((d_545_v0_).length(0)))):
                (d_545_v0_)[(d_546_i0_)] = False
        r1 = (self).f13
        d_547_v1_: D6
        d_547_v1_ = D6_DC17(self.f3)
        pat_let_tv18_ = p1
        pat_let_tv19_ = p1
        pat_let_tv20_ = p0
        def lambda22_(source14_):
            if source14_.is_DC17:
                d_548___mcc_h0_ = source14_.cf28
                d_549_cf28_ = d_548___mcc_h0_
                return pat_let_tv18_
            elif source14_.is_DC18:
                d_550___mcc_h1_ = source14_.cf29
                d_551___mcc_h2_ = source14_.cf30
                d_552___mcc_h3_ = source14_.cf31
                d_553___mcc_h4_ = source14_.cf32
                d_554___mcc_h5_ = source14_.cf33
                d_555_cf33_ = d_554___mcc_h5_
                d_556_cf32_ = d_553___mcc_h4_
                d_557_cf31_ = d_552___mcc_h3_
                d_558_cf30_ = d_551___mcc_h2_
                d_559_cf29_ = d_550___mcc_h1_
                return d_558_cf30_
            elif source14_.is_DC19:
                d_560___mcc_h6_ = source14_.cf34
                d_561_cf34_ = d_560___mcc_h6_
                return pat_let_tv19_
            elif source14_.is_DC16:
                d_562___mcc_h7_ = source14_.cf27
                d_563_cf27_ = d_562___mcc_h7_
                def iife65_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(499, 733):
                        d_564_v2_: int = compr_33_
                        if ((499) <= (d_564_v2_)) and ((d_564_v2_) < (733)):
                            coll33_[(d_564_v2_) - ((self).f13)] = self.f3
                    return _dafny.Map(coll33_)
                return default__.safeDivisionInt(len(iife65_()
                ), pat_let_tv20_)
            elif True:
                d_565___mcc_h8_ = source14_.cf35
                d_566_cf35_ = d_565___mcc_h8_
                return (0) - (len(((self).f2) + ((self).f2)))

        r2 = lambda22_(d_547_v1_)
        d_567_v3_: _dafny.Seq
        d_567_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axj"))
        d_567_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_568_i1_ in range(default__.abs(649))])
        if (192) == (p1):
            d_569_v4_: str
            d_569_v4_ = _dafny.CodePoint('m')
            d_570_v5_: _dafny.MultiSet
            d_570_v5_ = _dafny.MultiSet([_dafny.Map({p0: p1})])
            d_571_v6_: _dafny.Map
            d_571_v6_ = _dafny.Map({(d_570_v5_).cardinality: p0})
            d_572_v7_: D6
            d_572_v7_ = D6_DC18(len(d_571_v6_), len(d_567_v3_), self.f3, d_569_v4_, len(d_571_v6_))
            d_569_v4_ = (d_572_v7_).cf32
            d_573_v8_: D3
            d_573_v8_ = D3_DC10(151, p0)
            d_574_v9_: D3
            d_574_v9_ = D3_DC11(d_573_v8_)
            d_575_v10_: _dafny.Seq
            d_575_v10_ = _dafny.SeqWithoutIsStrInference([d_573_v8_, d_573_v8_])
            d_576_v11_: _dafny.Seq
            d_576_v11_ = _dafny.SeqWithoutIsStrInference([(self).f2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyygrb"))])
            d_577_v12_: D3
            d_577_v12_ = D3_DC11((d_575_v10_)[default__.safeIndex((self).fm27(p0, (d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))], _dafny.MultiSet(d_576_v11_), globalState), len(d_575_v10_))])
            d_578_v13_: _dafny.Set
            d_578_v13_ = _dafny.Set({d_577_v12_, D3_DC11(d_573_v8_), d_577_v12_, d_577_v12_})
            d_578_v13_ = ((d_578_v13_ if True else d_578_v13_) if (self).fm11((d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))], globalState) else (d_578_v13_ if (d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))] else d_578_v13_))
            (self).f3 = self.f3
            d_567_v3_ = d_567_v3_
            r1 = p1
        elif True:
            index75_ = default__.safeIndex(303, (d_545_v0_).length(0))
            (d_545_v0_)[index75_] = ((self).f2) == (d_567_v3_)
            r1 = default__.safeModuloInt(p0, p1)
            (self).f3 = (p1) != (p1)
            if self.f3:
                d_579_v14_: D1
                d_579_v14_ = D1_DC2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scdvhk")))
                d_580_v15_: _dafny.Map
                def iife66_(_pat_let16_0):
                    def iife67_(d_581_dt__update__tmp_h0_):
                        def iife68_(_pat_let17_0):
                            def iife69_(d_582_dt__update_hcf2_h0_):
                                return D1_DC2(d_582_dt__update_hcf2_h0_)
                            return iife69_(_pat_let17_0)
                        return iife68_((self).f2)
                    return iife67_(_pat_let16_0)
                d_580_v15_ = _dafny.Map({p1: iife66_(d_579_v14_)})
                (self).f3 = not ((d_580_v15_) == (d_580_v15_)) or ((d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))])
                d_583_v16_: _dafny.Seq
                d_583_v16_ = _dafny.SeqWithoutIsStrInference([330, p0])
                d_584_v17_: D4
                d_584_v17_ = D4_DC13(self.f3, (d_583_v16_ if not((d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))]) else d_583_v16_))
                pat_let_tv21_ = d_545_v0_
                pat_let_tv22_ = d_545_v0_
                pat_let_tv23_ = d_545_v0_
                pat_let_tv24_ = d_545_v0_
                def iife71_(_pat_let19_0):
                    def iife72_(d_585_dt__update__tmp_h1_):
                        def iife73_(_pat_let20_0):
                            def iife74_(d_586_dt__update_hcf21_h0_):
                                return D4_DC13(d_586_dt__update_hcf21_h0_, (d_585_dt__update__tmp_h1_).cf22)
                            return iife74_(_pat_let20_0)
                        return iife73_((pat_let_tv22_)[default__.safeIndex(303, (pat_let_tv21_).length(0))])
                    return iife72_(_pat_let19_0)
                def iife70_(_pat_let18_0):
                    def iife75_(d_587_dt__update__tmp_h2_):
                        def iife76_(_pat_let21_0):
                            def iife77_(d_588_dt__update_hcf21_h1_):
                                return D4_DC13(d_588_dt__update_hcf21_h1_, (d_587_dt__update__tmp_h2_).cf22)
                            return iife77_(_pat_let21_0)
                        return iife76_((pat_let_tv24_)[default__.safeIndex(303, (pat_let_tv23_).length(0))])
                    return iife75_(_pat_let18_0)
                d_584_v17_ = iife70_(iife71_(d_584_v17_))
                r2 = p0
                d_589_v18_: _dafny.Array
                def lambda23_(d_590_p0_):
                    def lambda24_(d_591_i2_):
                        return (d_591_i2_) * (d_590_p0_)

                    return lambda24_

                init11_ = lambda23_(p0)
                nw64_ = _dafny.Array(None, 25)
                for i0_11_ in range(nw64_.length(0)):
                    nw64_[i0_11_] = init11_(i0_11_)
                d_589_v18_ = nw64_
                d_592_v19_: _dafny.Map
                d_592_v19_ = _dafny.Map({(d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))]: d_589_v18_})
                d_592_v19_ = d_592_v19_
                r2 = (p1) - ((self).f13)
            elif True:
                d_593_v20_: D2
                d_593_v20_ = D2_DC5(_dafny.Map({False: (self).f13}))
                d_594_v21_: D2
                d_594_v21_ = D2_DC7(d_593_v20_)
                d_595_v22_: D2
                d_595_v22_ = D2_DC7(d_594_v21_)
                d_595_v22_ = d_595_v22_
                d_596_v23_: D5
                d_596_v23_ = D5_DC15(p0, self.f3, d_545_v0_)
                pat_let_tv25_ = d_545_v0_
                def iife79_(_pat_let23_0):
                    def iife80_(d_597_dt__update__tmp_h3_):
                        def iife81_(_pat_let24_0):
                            def iife82_(d_598_dt__update_hcf26_h0_):
                                return D5_DC15((d_597_dt__update__tmp_h3_).cf24, (d_597_dt__update__tmp_h3_).cf25, d_598_dt__update_hcf26_h0_)
                            return iife82_(_pat_let24_0)
                        return iife81_(pat_let_tv25_)
                    return iife80_(_pat_let23_0)
                def iife78_(_pat_let22_0):
                    def iife83_(d_599_dt__update__tmp_h4_):
                        def iife84_(_pat_let25_0):
                            def iife85_(d_600_dt__update_hcf25_h0_):
                                return D5_DC15((d_599_dt__update__tmp_h4_).cf24, d_600_dt__update_hcf25_h0_, (d_599_dt__update__tmp_h4_).cf26)
                            return iife85_(_pat_let25_0)
                        return iife84_(self.f3)
                    return iife83_(_pat_let22_0)
                d_596_v23_ = iife78_(iife79_(d_596_v23_))
                d_601_v24_: C1
                nw65_ = C1()
                nw65_.ctor__()
                d_601_v24_ = nw65_
                d_602_v25_: D0
                d_602_v25_ = D0_DC1(981)
                d_603_v26_: _dafny.Map
                d_603_v26_ = _dafny.Map({(self).f2: (d_602_v25_).cf1})
                d_604_v27_: _dafny.MultiSet
                d_604_v27_ = _dafny.MultiSet([default__.fm5(d_603_v26_, globalState)])
                d_605_v28_: D0
                d_605_v28_ = D0_DC0(p1)
                d_606_v29_: _dafny.MultiSet
                d_606_v29_ = _dafny.MultiSet([not(self.f3)])
                d_607_v30_: _dafny.Seq
                d_607_v30_ = _dafny.SeqWithoutIsStrInference([True, (d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))]])
                d_608_v31_: D8
                d_608_v31_ = D8_DC22(d_606_v29_)
                d_604_v27_ = (default__.fm34((d_605_v28_).cf0, (d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))], globalState) if True else _dafny.MultiSet([d_606_v29_, _dafny.MultiSet(d_607_v30_), (d_608_v31_).cf37]))
                d_609_v32_: _dafny.Set
                d_609_v32_ = _dafny.Set({(self).f13, p1, p1})
                d_610_v33_: _dafny.Array
                nw66_ = _dafny.Array(None, 8)
                nw66_[int(0)] = default__.safeModuloInt(p0, len(d_609_v32_))
                nw66_[int(1)] = (self).f13
                nw66_[int(2)] = p1
                nw66_[int(3)] = (0) - (p0)
                nw66_[int(4)] = p1
                nw66_[int(5)] = p0
                nw66_[int(6)] = p0
                nw66_[int(7)] = (609) + (p1)
                d_610_v33_ = nw66_
                index76_ = default__.safeIndex(340, (d_610_v33_).length(0))
                (d_610_v33_)[index76_] = 382
                index77_ = default__.safeIndex(340, (d_610_v33_).length(0))
                (d_610_v33_)[index77_] = (p0) - (p1)
            d_611_v34_: _dafny.Seq
            d_611_v34_ = _dafny.SeqWithoutIsStrInference([d_567_v3_, d_567_v3_, (self).f2])
            d_612_v35_: _dafny.Seq
            d_612_v35_ = _dafny.SeqWithoutIsStrInference([p0, (self).f13])
            d_613_v36_: str
            d_613_v36_ = _dafny.CodePoint('u')
            d_614_v37_: _dafny.Set
            d_614_v37_ = _dafny.Set({(self).f13, p1, len(d_611_v34_), len(_dafny.Map({(self).fm12(p0, (self).f13, (d_612_v35_)[default__.safeIndex(p1, len(d_612_v35_))], False, globalState): default__.fm31(len((self).f2), self.f3, p1, d_613_v36_, globalState)})), p1})
            d_615_v38_: _dafny.Map
            d_615_v38_ = _dafny.Map({p1: d_613_v36_})
            r1 = ((self).f13 if (d_614_v37_).ispropersubset(d_614_v37_) else ((self).f13 if (d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))] else len(d_615_v38_)))
        d_616_v39_: _dafny.Map
        d_616_v39_ = _dafny.Map({True: not((d_545_v0_)[default__.safeIndex(303, (d_545_v0_).length(0))])})
        r0 = (d_616_v39_) | (d_616_v39_)
        r1 = (0) - ((0) - (default__.fm0(globalState)))
        r2 = p1
        return r0, r1, r2

    @property
    def f13(self):
        return self._f13

class C3(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5, f3, f1, f2):
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return self.f3

    def fm14(self, p0, p1, p2, p3, globalState):
        return (self).f2

    def fm12(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm11(self, p0, globalState):
        return self.f3

    def fm25(self, p0, globalState):
        return (self).f2

    def fm26(self, p0, p1, p2, globalState):
        return ((self).f4) < (230)

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        d_617_v0_: _dafny.Set
        d_617_v0_ = _dafny.Set({p1})
        index78_ = default__.safeIndex(620, ((self).f5).length(0))
        ((self).f5)[index78_] = (d_617_v0_).issubset(_dafny.Set({p1}))
        index79_ = default__.safeIndex(620, ((self).f5).length(0))
        ((self).f5)[index79_] = (p0) == ((self).fm12(p0, p1, (self).f4, self.f3, globalState))
        r2 = p0
        hi4_ = 99
        for d_618_i0_ in range(p0, hi4_):
            d_619_v1_: _dafny.Array
            nw67_ = _dafny.Array(int(0), 15)
            d_619_v1_ = nw67_
            index80_ = default__.safeIndex(149, (d_619_v1_).length(0))
            (d_619_v1_)[index80_] = d_618_i0_
            index81_ = default__.safeIndex(149, (d_619_v1_).length(0))
            (d_619_v1_)[index81_] = (self).f4
            d_619_v1_ = d_619_v1_
            d_620_v2_: _dafny.MultiSet
            d_620_v2_ = _dafny.MultiSet([self.f3])
            d_621_v3_: _dafny.Map
            d_621_v3_ = _dafny.Map({(d_620_v2_).issubset(d_620_v2_): (self.f3 if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])})
            d_622_v4_: _dafny.MultiSet
            d_622_v4_ = _dafny.MultiSet([(0) - (p0)])
            d_623_v5_: _dafny.Seq
            d_623_v5_ = _dafny.SeqWithoutIsStrInference([d_618_i0_, p0])
            d_621_v3_ = (d_621_v3_).set((d_622_v4_).ispropersubset((_dafny.MultiSet([(0) - (p0), (0) - (d_618_i0_), len(d_623_v5_), p0, p1])).set((self).fm12(p1, p1, len((self).f2), self.f3, globalState), default__.abs((_dafny.MultiSet([self.f3])).cardinality))), self.f3)
            d_624_v6_: str
            d_624_v6_ = _dafny.CodePoint('k')
            index82_ = default__.safeIndex(620, ((self).f5).length(0))
            ((self).f5)[index82_] = (d_624_v6_) not in ((self).f2)
        d_625_i1_: int
        d_625_i1_ = 0
        with _dafny.label("5"):
            while not(not(self.f3)):
                with _dafny.c_label("5"):
                    if (d_625_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_625_i1_ = (d_625_i1_) + (1)
                    d_626_v7_: _dafny.Map
                    d_626_v7_ = _dafny.Map({(p1) + (360): (self).f4})
                    d_627_v8_: _dafny.Seq
                    d_627_v8_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                    d_628_v9_: _dafny.Map
                    d_628_v9_ = _dafny.Map({self.f3: len(_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], self.f3]))})
                    d_629_v10_: _dafny.Set
                    d_629_v10_ = _dafny.Set({D2_DC5(d_628_v9_)})
                    d_626_v7_ = (d_626_v7_).set((d_627_v8_)[default__.safeIndex((self).f4, len(d_627_v8_))], len(d_629_v10_))
                    index83_ = default__.safeIndex(620, ((self).f5).length(0))
                    ((self).f5)[index83_] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                    d_630_v11_: str
                    d_630_v11_ = _dafny.CodePoint('n')
                    d_630_v11_ = d_630_v11_
                    d_631_v12_: _dafny.Array
                    nw68_ = _dafny.Array(None, 6)
                    nw68_[int(0)] = False
                    nw68_[int(1)] = self.f3
                    nw68_[int(2)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                    nw68_[int(3)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                    nw68_[int(4)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                    nw68_[int(5)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                    d_631_v12_ = nw68_
                    d_632_v13_: _dafny.Set
                    d_632_v13_ = _dafny.Set({(d_631_v12_ if self.f3 else d_631_v12_), d_631_v12_, d_631_v12_, d_631_v12_})
                    d_632_v13_ = _dafny.Set({d_631_v12_})
                    pass
            pass
        d_633_i2_: int
        d_633_i2_ = 0
        with _dafny.label("6"):
            while ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]:
                with _dafny.c_label("6"):
                    if (d_633_i2_) >= (100):
                        raise _dafny.Break("6")
                    d_633_i2_ = (d_633_i2_) + (1)
                    d_634_v14_: D0
                    d_634_v14_ = D0_DC0((self).f4)
                    d_635_v15_: str
                    d_635_v15_ = _dafny.CodePoint('e')
                    d_636_v16_: D6
                    d_636_v16_ = D6_DC18(len((self).f2), (d_634_v14_).cf0, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], d_635_v15_, (self).f4)
                    d_637_v17_: _dafny.MultiSet
                    d_637_v17_ = _dafny.MultiSet([340, p0, (self).f4, (self).f4])
                    if ((_dafny.MultiSet([(self).f4])).issubset(d_637_v17_) if (d_636_v16_).cf31 else self.f3):
                        d_638_v18_: _dafny.Array
                        nw69_ = _dafny.Array(int(0), 2)
                        d_638_v18_ = nw69_
                        d_639_v19_: _dafny.MultiSet
                        d_639_v19_ = _dafny.MultiSet([self.f3])
                        index84_ = default__.safeIndex(672, (d_638_v18_).length(0))
                        (d_638_v18_)[index84_] = ((d_639_v19_)[False] if (False) in (d_639_v19_) else len((self).f2))
                        d_640_v20_: _dafny.Seq
                        d_640_v20_ = _dafny.SeqWithoutIsStrInference([(self).fm11(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], globalState), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]])
                        index85_ = default__.safeIndex(672, (d_638_v18_).length(0))
                        (d_638_v18_)[index85_] = (0) - (((p0 if (d_640_v20_)[default__.safeIndex(p0, len(d_640_v20_))] else p1)) + (default__.safeModuloInt(p1, default__.fm0(globalState))))
                        index86_ = default__.safeIndex(620, ((self).f5).length(0))
                        ((self).f5)[index86_] = (self).fm13(globalState)
                        d_641_v21_: _dafny.Map
                        d_641_v21_ = _dafny.Map({d_638_v18_: not(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])})
                        d_641_v21_ = (d_641_v21_).set(d_638_v18_, self.f3)
                        d_642_v22_: _dafny.Map
                        d_642_v22_ = _dafny.Map({self.f3: p0})
                        d_642_v22_ = (d_642_v22_).set(self.f3, (p0) - ((self).f4))
                        d_641_v21_ = (_dafny.Map({d_638_v18_: True})) | ((_dafny.Map({d_638_v18_: False})) | (d_641_v21_))
                    elif True:
                        r1 = default__.fm0(globalState)
                        d_643_v24_: _dafny.Map
                        def iife86_():
                            coll34_ = _dafny.Map()
                            compr_34_: int
                            for compr_34_ in (d_637_v17_).Elements:
                                d_644_v23_: int = compr_34_
                                if (d_644_v23_) in (d_637_v17_):
                                    coll34_[(d_644_v23_) * ((self).f4)] = self.f3
                            return _dafny.Map(coll34_)
                        d_643_v24_ = _dafny.Map({iife86_()
                        : p0})
                        rhs78_ = ((len((self).f2)) + (p0)) * ((self).f4)
                        rhs79_ = (d_643_v24_) | (d_643_v24_)
                        rhs80_ = p1
                        r2 = rhs78_
                        d_643_v24_ = rhs79_
                        r1 = rhs80_
                        d_645_v25_: _dafny.Array
                        nw70_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                        d_645_v25_ = nw70_
                        d_646_v26_: _dafny.Array
                        nw71_ = _dafny.Array(_dafny.Array(None, 0), 12)
                        d_646_v26_ = nw71_
                        d_647_v27_: _dafny.Array
                        nw72_ = _dafny.Array(int(0), 28)
                        d_647_v27_ = nw72_
                        index87_ = default__.safeIndex(266, (d_647_v27_).length(0))
                        (d_647_v27_)[index87_] = len((d_617_v0_) | (d_617_v0_))
                        d_648_v28_: _dafny.Seq
                        d_648_v28_ = _dafny.SeqWithoutIsStrInference([d_645_v25_])
                        d_649_v29_: _dafny.Seq
                        d_649_v29_ = _dafny.SeqWithoutIsStrInference([(self).fm14(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], default__.fm3(p1, globalState), d_635_v15_, p0, globalState)])
                        d_650_v30_: _dafny.Map
                        d_650_v30_ = _dafny.Map({((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]: self.f3})
                        index88_ = default__.safeIndex(620, ((self).f5).length(0))
                        index89_ = default__.safeIndex(266, (d_647_v27_).length(0))
                        rhs81_ = self.f3
                        rhs82_ = (d_648_v28_)[default__.safeIndex(759, len(d_648_v28_))]
                        rhs83_ = d_646_v26_
                        rhs84_ = ((self).f4) - (len(d_649_v29_))
                        rhs85_ = ((_dafny.Map({self.f3: self.f3})) | (d_650_v30_)) | (_dafny.Map({self.f3: self.f3}))
                        lhs34_ = (self).f5
                        lhs35_ = default__.safeIndex(620, ((self).f5).length(0))
                        lhs36_ = d_647_v27_
                        lhs37_ = default__.safeIndex(266, (d_647_v27_).length(0))
                        lhs34_[lhs35_] = rhs81_
                        d_645_v25_ = rhs82_
                        d_646_v26_ = rhs83_
                        lhs36_[lhs37_] = rhs84_
                        r0 = rhs85_
                        d_651_v31_: _dafny.MultiSet
                        d_651_v31_ = _dafny.MultiSet([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]])
                        index90_ = default__.safeIndex(620, ((self).f5).length(0))
                        ((self).f5)[index90_] = (d_651_v31_) != (d_651_v31_)
                        d_652_v32_: _dafny.Seq
                        d_652_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfebqujas"))
                        index91_ = default__.safeIndex(620, ((self).f5).length(0))
                        rhs86_ = d_652_v32_
                        rhs87_ = self.f3
                        rhs88_ = p0
                        lhs38_ = (self).f5
                        lhs39_ = default__.safeIndex(620, ((self).f5).length(0))
                        d_652_v32_ = rhs86_
                        lhs38_[lhs39_] = rhs87_
                        r1 = rhs88_
                    d_653_v33_: _dafny.Array
                    def lambda25_(d_654_p1_):
                        def lambda26_(d_655_i3_):
                            return (d_655_i3_) * (d_654_p1_)

                        return lambda26_

                    init12_ = lambda25_(p1)
                    nw73_ = _dafny.Array(None, 17)
                    for i0_12_ in range(nw73_.length(0)):
                        nw73_[i0_12_] = init12_(i0_12_)
                    d_653_v33_ = nw73_
                    d_656_v34_: C0
                    nw74_ = C0()
                    nw74_.ctor__(d_653_v33_)
                    d_656_v34_ = nw74_
                    d_657_v35_: _dafny.Seq
                    d_657_v35_ = _dafny.SeqWithoutIsStrInference([((d_637_v17_)[p1] if (p1) in (d_637_v17_) else p1)])
                    if not(not((p1) != (((d_657_v35_)[default__.safeIndex((d_637_v17_).cardinality, len(d_657_v35_))]) - ((self).f4)))):
                        d_658_v36_: _dafny.Seq
                        d_658_v36_ = _dafny.SeqWithoutIsStrInference([(self).f1, (self).f1, (self).f1, (self).f1])
                        d_659_v37_: _dafny.Map
                        d_659_v37_ = _dafny.Map({self.f3: 332})
                        d_660_v38_: C2
                        nw75_ = C2()
                        nw75_.ctor__(((self).f4) + ((0) - ((self).f4)), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (d_658_v36_)[default__.safeIndex(p0, len(d_658_v36_))], (_dafny.SeqWithoutIsStrInference([d_635_v15_ for d_661_i4_ in range(default__.abs(-691))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsrcir"))).set(default__.safeIndex(len(d_659_v37_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsrcir")))), d_635_v15_)))
                        d_660_v38_ = nw75_
                        index92_ = default__.safeIndex(620, ((self).f5).length(0))
                        ((self).f5)[index92_] = ((((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else (self).fm26((self).f4, self.f3, False, globalState)) if self.f3 else self.f3)
                        d_662_v39_: _dafny.Seq
                        d_662_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynnlqps"))
                        rhs89_ = d_662_v39_
                        rhs90_ = (0) - ((self).f4)
                        rhs91_ = d_635_v15_
                        d_662_v39_ = rhs89_
                        r1 = rhs90_
                        d_635_v15_ = rhs91_
                        (self).f3 = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                        d_663_v40_: _dafny.Map
                        d_663_v40_ = _dafny.Map({(d_660_v38_).f13: default__.safeModuloInt((self).f4, p0)})
                        d_663_v40_ = (d_663_v40_).set(((self).f4) + (p0), (len(d_657_v35_)) - ((d_660_v38_).f13))
                    elif True:
                        r2 = p1
                        r1 = default__.safeModuloInt((self).f4, p1)
                        index93_ = default__.safeIndex(307, (d_653_v33_).length(0))
                        (d_653_v33_)[index93_] = (self).f4
                        index94_ = default__.safeIndex(307, (d_653_v33_).length(0))
                        (d_653_v33_)[index94_] = 234
                        r2 = (self).f4
                        index95_ = default__.safeIndex(307, (d_653_v33_).length(0))
                        (d_653_v33_)[index95_] = (d_653_v33_)[default__.safeIndex(307, (d_653_v33_).length(0))]
                    d_664_v41_: _dafny.Set
                    d_664_v41_ = _dafny.Set({(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]) and (((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])})
                    d_664_v41_ = (d_664_v41_) | ((d_664_v41_) | (d_664_v41_))
                    pass
            pass
        if self.f3:
            d_665_v42_: _dafny.Array
            def lambda27_(d_666_i5_):
                return ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]

            init13_ = lambda27_
            nw76_ = _dafny.Array(None, 3)
            for i0_13_ in range(nw76_.length(0)):
                nw76_[i0_13_] = init13_(i0_13_)
            d_665_v42_ = nw76_
            d_667_v43_: _dafny.Map
            d_667_v43_ = _dafny.Map({d_665_v42_: (self).f4})
            d_667_v43_ = (d_667_v43_).set(d_665_v42_, p0)
            d_668_v44_: _dafny.MultiSet
            d_668_v44_ = _dafny.MultiSet([self.f3, True])
            d_669_v45_: _dafny.Map
            d_669_v45_ = _dafny.Map({_dafny.Map({-745: self.f3}): self.f3})
            d_670_v46_: _dafny.Seq
            d_670_v46_ = _dafny.SeqWithoutIsStrInference([d_669_v45_])
            d_671_v47_: _dafny.Seq
            d_671_v47_ = _dafny.SeqWithoutIsStrInference([len((self).f2), p1, len(_dafny.Map({(self).f4: p0})), p1, len((d_670_v46_)[default__.safeIndex(p1, len(d_670_v46_))])])
            d_672_v48_: _dafny.Map
            d_672_v48_ = _dafny.Map({self.f3: (self).f4})
            d_673_v49_: _dafny.Seq
            d_673_v49_ = _dafny.SeqWithoutIsStrInference([len((self).f2), len(d_671_v47_), len(d_672_v48_)])
            d_674_v50_: _dafny.Set
            d_674_v50_ = _dafny.Set({((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]})
            d_675_v51_: D1
            d_675_v51_ = D1_DC3(self.f3, 823, p0)
            d_676_v52_: _dafny.MultiSet
            d_676_v52_ = _dafny.MultiSet([(d_675_v51_).cf4, (self).f4, 933, -450, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "il")))])
            d_677_v53_: _dafny.Map
            d_677_v53_ = _dafny.Map({d_668_v44_: (self).f4})
            d_678_v54_: _dafny.Seq
            d_678_v54_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
            d_679_v55_: _dafny.Array
            nw77_ = _dafny.Array(None, 24)
            nw77_[int(0)] = default__.safeDivisionInt(p0, -741)
            nw77_[int(1)] = (self).f4
            nw77_[int(2)] = (self).f4
            nw77_[int(3)] = (self).f4
            nw77_[int(4)] = (self).f4
            nw77_[int(5)] = p1
            nw77_[int(6)] = len((self).f2)
            nw77_[int(7)] = (-569) + (p1)
            nw77_[int(8)] = default__.safeDivisionInt(p0, p0)
            nw77_[int(9)] = (self).f4
            nw77_[int(10)] = p1
            nw77_[int(11)] = (d_668_v44_).cardinality
            nw77_[int(12)] = (0) - ((0) - ((self).f4))
            nw77_[int(13)] = (len(d_671_v47_)) + (p0)
            nw77_[int(14)] = p0
            nw77_[int(15)] = default__.safeDivisionInt((0) - (len(d_674_v50_)), (d_676_v52_).cardinality)
            nw77_[int(16)] = default__.safeDivisionInt((self).f4, ((d_677_v53_)[d_668_v44_] if (d_668_v44_) in (d_677_v53_) else p0))
            nw77_[int(17)] = (self).f4
            nw77_[int(18)] = (0) - ((p0) + (p1))
            nw77_[int(19)] = p0
            nw77_[int(20)] = default__.safeModuloInt(-374, len(d_674_v50_))
            nw77_[int(21)] = p0
            nw77_[int(22)] = p0
            nw77_[int(23)] = len(d_678_v54_)
            d_679_v55_ = nw77_
            d_679_v55_ = d_679_v55_
            index96_ = default__.safeIndex(556, (d_679_v55_).length(0))
            (d_679_v55_)[index96_] = (self).f4
            index97_ = default__.safeIndex(556, (d_679_v55_).length(0))
            (d_679_v55_)[index97_] = (default__.fm0(globalState) if not(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]) else p0)
            if self.f3:
                d_678_v54_ = d_678_v54_
                (self).f3 = (default__.safeDivisionInt(p1, (self).fm12(p0, 202, (d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], globalState))) > (default__.safeModuloInt(p0, -580))
                r1 = (0) - (default__.fm0(globalState))
                d_680_v56_: str
                d_680_v56_ = _dafny.CodePoint('l')
                d_680_v56_ = d_680_v56_
                d_681_v57_: _dafny.Array
                nw78_ = _dafny.Array(_dafny.Set({}), 27)
                d_681_v57_ = nw78_
                d_682_v58_: _dafny.Map
                d_682_v58_ = _dafny.Map({self.f3: d_681_v57_})
                d_682_v58_ = (d_682_v58_).set(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], d_681_v57_)
            elif True:
                d_683_v59_: C1
                nw79_ = C1()
                nw79_.ctor__()
                d_683_v59_ = nw79_
                index98_ = default__.safeIndex(620, ((self).f5).length(0))
                ((self).f5)[index98_] = False
                d_684_v60_: _dafny.Map
                d_684_v60_ = _dafny.Map({(0) - (p0): d_665_v42_})
                (self).f3 = (d_684_v60_) == (d_684_v60_)
                d_685_v61_: D1
                d_685_v61_ = D1_DC2(default__.fm2(self.f3, globalState))
                d_686_v62_: _dafny.Map
                d_686_v62_ = _dafny.Map({38: ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]})
                d_687_v63_: _dafny.Map
                d_687_v63_ = _dafny.Map({d_686_v62_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_688_i6_ in range(default__.abs(131))])})
                d_689_v64_: _dafny.Map
                d_689_v64_ = _dafny.Map({len(d_687_v63_): (self).f4})
                d_690_v65_: _dafny.Map
                d_690_v65_ = _dafny.Map({self.f3: ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]})
                d_691_v66_: str
                d_691_v66_ = _dafny.CodePoint('b')
                d_692_v67_: _dafny.Seq
                d_692_v67_ = _dafny.SeqWithoutIsStrInference([d_671_v47_])
                d_693_v68_: _dafny.Map
                d_693_v68_ = _dafny.Map({((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]: d_671_v47_})
                d_694_v69_: _dafny.Array
                nw80_ = _dafny.Array(None, 23)
                nw80_[int(0)] = d_673_v49_
                nw80_[int(1)] = _dafny.SeqWithoutIsStrInference([(0) - ((d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))])])
                nw80_[int(2)] = d_673_v49_
                nw80_[int(3)] = (_dafny.SeqWithoutIsStrInference([len((d_685_v61_).cf2)])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([len((d_685_v61_).cf2)]))), p1)
                nw80_[int(4)] = d_671_v47_
                nw80_[int(5)] = d_673_v49_
                nw80_[int(6)] = (d_673_v49_) + (d_671_v47_)
                nw80_[int(7)] = (d_673_v49_) + (d_673_v49_)
                nw80_[int(8)] = default__.fm31(((d_689_v64_)[p0] if (p0) in (d_689_v64_) else len(d_690_v65_)), not(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]), (d_676_v52_).cardinality, d_691_v66_, globalState)
                nw80_[int(9)] = d_673_v49_
                nw80_[int(10)] = default__.fm31(p1, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (0) - (p1), d_691_v66_, globalState)
                nw80_[int(11)] = (_dafny.SeqWithoutIsStrInference([(d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))] for d_695_i7_ in range(default__.abs(14))]) if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else d_671_v47_)
                nw80_[int(12)] = (d_673_v49_ if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else d_673_v49_)
                nw80_[int(13)] = _dafny.SeqWithoutIsStrInference([len(d_674_v50_) for d_696_i8_ in range(default__.abs(24))])
                nw80_[int(14)] = ((d_692_v67_)[default__.safeIndex((self).f4, len(d_692_v67_))]) + (d_671_v47_)
                nw80_[int(15)] = d_673_v49_
                nw80_[int(16)] = ((d_693_v68_)[((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]] if (((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]) in (d_693_v68_) else d_673_v49_)
                nw80_[int(17)] = _dafny.SeqWithoutIsStrInference([(0) - (p1)])
                nw80_[int(18)] = (_dafny.SeqWithoutIsStrInference([p1])) + (d_673_v49_)
                nw80_[int(19)] = d_673_v49_
                nw80_[int(20)] = d_671_v47_
                nw80_[int(21)] = d_673_v49_
                nw80_[int(22)] = d_671_v47_
                d_694_v69_ = nw80_
                index99_ = default__.safeIndex(740, (d_694_v69_).length(0))
                (d_694_v69_)[index99_] = (_dafny.SeqWithoutIsStrInference([p1, ((d_672_v48_)[(d_678_v54_)[default__.safeIndex((self).f4, len(d_678_v54_))]] if ((d_678_v54_)[default__.safeIndex((self).f4, len(d_678_v54_))]) in (d_672_v48_) else p0), (self).f4, default__.fm0(globalState)])).set(default__.safeIndex((d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))], len(_dafny.SeqWithoutIsStrInference([p1, ((d_672_v48_)[(d_678_v54_)[default__.safeIndex((self).f4, len(d_678_v54_))]] if ((d_678_v54_)[default__.safeIndex((self).f4, len(d_678_v54_))]) in (d_672_v48_) else p0), (self).f4, default__.fm0(globalState)]))), p1)
                index100_ = default__.safeIndex(740, (d_694_v69_).length(0))
                index101_ = default__.safeIndex(556, (d_679_v55_).length(0))
                index102_ = default__.safeIndex(556, (d_679_v55_).length(0))
                index103_ = default__.safeIndex(620, ((self).f5).length(0))
                rhs92_ = (d_692_v67_)[default__.safeIndex(p1, len(d_692_v67_))]
                rhs93_ = d_691_v66_
                rhs94_ = (len((d_686_v62_) | (d_686_v62_))) + ((len(d_678_v54_) if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else (d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))]))
                rhs95_ = (556) * (default__.safeDivisionInt(740, (d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))]))
                rhs96_ = self.f3
                lhs40_ = d_694_v69_
                lhs41_ = default__.safeIndex(740, (d_694_v69_).length(0))
                lhs42_ = d_679_v55_
                lhs43_ = default__.safeIndex(556, (d_679_v55_).length(0))
                lhs44_ = d_679_v55_
                lhs45_ = default__.safeIndex(556, (d_679_v55_).length(0))
                lhs46_ = (self).f5
                lhs47_ = default__.safeIndex(620, ((self).f5).length(0))
                lhs40_[lhs41_] = rhs92_
                d_691_v66_ = rhs93_
                lhs42_[lhs43_] = rhs94_
                lhs44_[lhs45_] = rhs95_
                lhs46_[lhs47_] = rhs96_
                d_697_v70_: _dafny.Array
                nw81_ = _dafny.Array(D4.default()(), 18)
                d_697_v70_ = nw81_
                d_698_v71_: _dafny.Map
                d_698_v71_ = _dafny.Map({self.f3: d_697_v70_})
                d_699_v72_: _dafny.Seq
                d_699_v72_ = _dafny.SeqWithoutIsStrInference([d_697_v70_, d_697_v70_, d_697_v70_])
                d_700_v73_: _dafny.Map
                d_700_v73_ = _dafny.Map({p1: d_697_v70_})
                d_701_v74_: D9
                d_701_v74_ = D9_DC25(d_697_v70_)
                d_702_v75_: _dafny.Array
                nw82_ = _dafny.Array(None, 23)
                nw82_[int(0)] = d_697_v70_
                nw82_[int(1)] = d_697_v70_
                nw82_[int(2)] = d_697_v70_
                nw82_[int(3)] = d_697_v70_
                nw82_[int(4)] = d_697_v70_
                nw82_[int(5)] = d_697_v70_
                nw82_[int(6)] = d_697_v70_
                nw82_[int(7)] = ((d_698_v71_)[self.f3] if (self.f3) in (d_698_v71_) else d_697_v70_)
                nw82_[int(8)] = d_697_v70_
                nw82_[int(9)] = d_697_v70_
                nw82_[int(10)] = (d_699_v72_)[default__.safeIndex((self).f4, len(d_699_v72_))]
                nw82_[int(11)] = d_697_v70_
                nw82_[int(12)] = d_697_v70_
                nw82_[int(13)] = d_697_v70_
                nw82_[int(14)] = d_697_v70_
                nw82_[int(15)] = ((d_700_v73_)[(d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))]] if ((d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))]) in (d_700_v73_) else (d_701_v74_).cf43)
                nw82_[int(16)] = d_697_v70_
                nw82_[int(17)] = (D9_DC25(d_697_v70_)).cf43
                nw82_[int(18)] = d_697_v70_
                nw82_[int(19)] = d_697_v70_
                nw82_[int(20)] = d_697_v70_
                nw82_[int(21)] = d_697_v70_
                nw82_[int(22)] = d_697_v70_
                d_702_v75_ = nw82_
                index104_ = default__.safeIndex(666, (d_702_v75_).length(0))
                (d_702_v75_)[index104_] = d_697_v70_
                index105_ = default__.safeIndex(666, (d_702_v75_).length(0))
                rhs97_ = d_617_v0_
                rhs98_ = d_697_v70_
                rhs99_ = self.f3
                rhs100_ = d_665_v42_
                lhs48_ = d_702_v75_
                lhs49_ = default__.safeIndex(666, (d_702_v75_).length(0))
                lhs50_ = self
                d_617_v0_ = rhs97_
                lhs48_[lhs49_] = rhs98_
                lhs50_.f3 = rhs99_
                d_665_v42_ = rhs100_
            index106_ = default__.safeIndex(556, (d_679_v55_).length(0))
            (d_679_v55_)[index106_] = (d_679_v55_)[default__.safeIndex(556, (d_679_v55_).length(0))]
        elif True:
            d_703_v76_: _dafny.Array
            nw83_ = _dafny.Array(int(0), 29)
            d_703_v76_ = nw83_
            index107_ = default__.safeIndex(404, (d_703_v76_).length(0))
            (d_703_v76_)[index107_] = p1
            d_704_v77_: _dafny.MultiSet
            d_704_v77_ = _dafny.MultiSet([(self).fm13(globalState)])
            d_705_v78_: _dafny.Set
            d_705_v78_ = _dafny.Set({d_704_v77_})
            index108_ = default__.safeIndex(404, (d_703_v76_).length(0))
            (d_703_v76_)[index108_] = default__.safeModuloInt(len((d_705_v78_ if True else d_705_v78_)), p0)
            if ((self.f3) == (False)) == ((p0) == ((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))])):
                d_706_v79_: _dafny.Seq
                d_706_v79_ = _dafny.SeqWithoutIsStrInference([(self).f4, p1, p0])
                d_706_v79_ = d_706_v79_
                d_707_v80_: _dafny.MultiSet
                d_707_v80_ = _dafny.MultiSet([p1, 617])
                index109_ = default__.safeIndex(404, (d_703_v76_).length(0))
                index110_ = default__.safeIndex(404, (d_703_v76_).length(0))
                rhs101_ = (d_707_v80_) | (d_707_v80_)
                rhs102_ = ((0) - ((0) - (p0))) + ((self).f4)
                rhs103_ = p0
                lhs51_ = d_703_v76_
                lhs52_ = default__.safeIndex(404, (d_703_v76_).length(0))
                lhs53_ = d_703_v76_
                lhs54_ = default__.safeIndex(404, (d_703_v76_).length(0))
                d_707_v80_ = rhs101_
                lhs51_[lhs52_] = rhs102_
                lhs53_[lhs54_] = rhs103_
                d_708_v81_: _dafny.Array
                nw84_ = _dafny.Array(None, 10)
                nw84_[int(0)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                nw84_[int(1)] = (p1) < (p1)
                nw84_[int(2)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                nw84_[int(3)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                nw84_[int(4)] = (self.f3) == (self.f3)
                nw84_[int(5)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                nw84_[int(6)] = True
                nw84_[int(7)] = (default__.fm3(len((self).f2), globalState)) or (((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])
                nw84_[int(8)] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                nw84_[int(9)] = ((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))]) < (-703)
                d_708_v81_ = nw84_
                d_709_v82_: _dafny.Set
                d_709_v82_ = _dafny.Set({self.f3, not(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]), self.f3, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], self.f3})
                d_710_v83_: _dafny.Seq
                d_710_v83_ = _dafny.SeqWithoutIsStrInference([(((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] if self.f3 else self.f3), (self.f3) not in (d_709_v82_)])
                d_711_v84_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.Map({}), 20)
                d_711_v84_ = nw85_
                d_712_v85_: _dafny.Map
                d_712_v85_ = _dafny.Map({p1: d_706_v79_})
                index111_ = default__.safeIndex(773, (d_711_v84_).length(0))
                (d_711_v84_)[index111_] = d_712_v85_
                d_713_v86_: str
                d_713_v86_ = _dafny.CodePoint('c')
                index112_ = default__.safeIndex(773, (d_711_v84_).length(0))
                rhs104_ = (len(d_706_v79_)) * (p1)
                rhs105_ = d_708_v81_
                rhs106_ = ((d_710_v83_) + (((d_710_v83_) + (_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], False, self.f3, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]]))).set(default__.safeIndex(len(default__.fm31((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], d_713_v86_, globalState)), len((d_710_v83_) + (_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], False, self.f3, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]])))), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]))).set(default__.safeIndex((self).f4, len((d_710_v83_) + (((d_710_v83_) + (_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], False, self.f3, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]]))).set(default__.safeIndex(len(default__.fm31((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], d_713_v86_, globalState)), len((d_710_v83_) + (_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], False, self.f3, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]])))), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])))), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])
                rhs107_ = d_712_v85_
                lhs55_ = d_711_v84_
                lhs56_ = default__.safeIndex(773, (d_711_v84_).length(0))
                r1 = rhs104_
                d_708_v81_ = rhs105_
                d_710_v83_ = rhs106_
                lhs55_[lhs56_] = rhs107_
                d_714_v87_: D5
                d_714_v87_ = D5_DC15(len((self).f2), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], d_708_v81_)
                d_715_v88_: _dafny.Map
                d_715_v88_ = _dafny.Map({p0: (d_714_v87_).cf26})
                d_708_v81_ = ((d_715_v88_)[p0] if (p0) in (d_715_v88_) else d_708_v81_)
                def lambda28_(d_716_p0_):
                    def lambda29_(d_717_i9_):
                        return (d_717_i9_) + (d_716_p0_)

                    return lambda29_

                init14_ = lambda28_(p0)
                nw86_ = _dafny.Array(None, 14)
                for i0_14_ in range(nw86_.length(0)):
                    nw86_[i0_14_] = init14_(i0_14_)
                d_703_v76_ = nw86_
            elif True:
                index113_ = default__.safeIndex(620, ((self).f5).length(0))
                ((self).f5)[index113_] = ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]
                d_718_v89_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                d_718_v89_ = nw87_
                d_719_v90_: str
                d_719_v90_ = _dafny.CodePoint('i')
                index114_ = default__.safeIndex(837, (d_718_v89_).length(0))
                (d_718_v89_)[index114_] = d_719_v90_
                index115_ = default__.safeIndex(837, (d_718_v89_).length(0))
                index116_ = default__.safeIndex(404, (d_703_v76_).length(0))
                rhs108_ = (d_719_v90_ if not((737) >= ((self).f4)) else (d_719_v90_ if ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))] else d_719_v90_))
                rhs109_ = p0
                lhs57_ = d_718_v89_
                lhs58_ = default__.safeIndex(837, (d_718_v89_).length(0))
                lhs59_ = d_703_v76_
                lhs60_ = default__.safeIndex(404, (d_703_v76_).length(0))
                lhs57_[lhs58_] = rhs108_
                lhs59_[lhs60_] = rhs109_
                d_720_v91_: _dafny.Seq
                d_720_v91_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
                d_721_v92_: _dafny.Seq
                d_721_v92_ = _dafny.SeqWithoutIsStrInference([(d_720_v91_)[default__.safeIndex((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], len(d_720_v91_))], (self.f3 if self.f3 else ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))])])
                d_721_v92_ = default__.fm7(not (((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]) or (True), ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], globalState)
                d_722_v93_: C1
                nw88_ = C1()
                nw88_.ctor__()
                d_722_v93_ = nw88_
                d_723_v94_: _dafny.Seq
                d_723_v94_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmwbhci"))
                d_724_v95_: _dafny.MultiSet
                d_724_v95_ = _dafny.MultiSet([p1, (self).f4])
                d_725_v96_: _dafny.Seq
                d_725_v96_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puaoww"))])
                d_726_v97_: _dafny.Seq
                d_726_v97_ = _dafny.SeqWithoutIsStrInference([(d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))]])
                d_727_v98_: D2
                d_727_v98_ = D2_DC6((d_723_v94_)[default__.safeIndex((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], len(d_723_v94_))], self.f3, self.f3, len((d_725_v96_)[default__.safeIndex((0) - (len(d_726_v97_)), len(d_725_v96_))]))
                d_728_v99_: _dafny.Map
                d_728_v99_ = _dafny.Map({((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]: p0})
                d_729_v100_: _dafny.Seq
                d_729_v100_ = _dafny.SeqWithoutIsStrInference([572, (d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], len((default__.fm31((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))], ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (d_724_v95_).cardinality, (d_727_v98_).cf9, globalState)) + (d_726_v97_)), (((d_728_v99_)[True] if (True) in (d_728_v99_) else (d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))])) + ((d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))]), 363])
                rhs110_ = d_723_v94_
                rhs111_ = (d_723_v94_) + ((self).f2)
                rhs112_ = (d_729_v100_) + (d_729_v100_)
                rhs113_ = _dafny.Map({self.f3: (p0) - (len(d_729_v100_))})
                rhs114_ = (d_718_v89_)[default__.safeIndex(837, (d_718_v89_).length(0))]
                d_723_v94_ = rhs110_
                d_723_v94_ = rhs111_
                d_729_v100_ = rhs112_
                d_728_v99_ = rhs113_
                d_719_v90_ = rhs114_
            r1 = p1
            r1 = (d_703_v76_)[default__.safeIndex(404, (d_703_v76_).length(0))]
            d_730_v101_: str
            d_730_v101_ = _dafny.CodePoint('j')
            d_731_v102_: _dafny.MultiSet
            d_731_v102_ = _dafny.MultiSet([_dafny.CodePoint('d'), d_730_v101_, d_730_v101_])
            d_731_v102_ = (d_731_v102_) - (d_731_v102_)
        d_732_v103_: _dafny.MultiSet
        d_732_v103_ = _dafny.MultiSet([((self).f5)[default__.safeIndex(620, ((self).f5).length(0))]])
        d_733_v104_: D8
        d_733_v104_ = D8_DC23(((d_732_v103_)[True] if (True) in (d_732_v103_) else p0), (self).f4, self.f3)
        d_734_v105_: _dafny.Map
        d_734_v105_ = _dafny.Map({(d_733_v104_).cf40: self.f3})
        r0 = ((d_734_v105_) | (d_734_v105_)).set(self.f3, ((self).f2) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_735_i10_ in range(default__.abs(416))])))
        r1 = (self).f4
        r2 = p0
        return r0, r1, r2

    def m10(self, p0, globalState):
        r0: D5 = D5.default()()
        r1: bool = False
        if self.f3:
            d_736_v0_: str
            d_736_v0_ = _dafny.CodePoint('o')
            d_736_v0_ = d_736_v0_
            d_737_v1_: _dafny.Array
            nw89_ = _dafny.Array(int(0), 26)
            d_737_v1_ = nw89_
            d_737_v1_ = d_737_v1_
            d_738_v2_: _dafny.Seq
            d_738_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_738_v2_ = (self).f2
            d_739_v3_: int
            d_739_v3_ = -585
            rhs115_ = self.f3
            rhs116_ = d_739_v3_
            lhs61_ = self
            lhs61_.f3 = rhs115_
            d_739_v3_ = rhs116_
            d_740_v4_: C1
            nw90_ = C1()
            nw90_.ctor__()
            d_740_v4_ = nw90_
        elif True:
            d_741_v5_: _dafny.Map
            d_741_v5_ = _dafny.Map({True: p0})
            d_741_v5_ = (d_741_v5_).set(not((p0) or (self.f3)), not(((self).f4) >= ((self).f4)))
            d_742_v6_: _dafny.Set
            d_742_v6_ = _dafny.Set({False, p0, p0})
            index117_ = default__.safeIndex(786, ((self).f5).length(0))
            ((self).f5)[index117_] = (d_742_v6_).issubset(d_742_v6_)
            d_743_v7_: _dafny.Map
            d_743_v7_ = _dafny.Map({default__.safeModuloInt((self).f4, (self).f4): True})
            index118_ = default__.safeIndex(786, ((self).f5).length(0))
            rhs117_ = (((self).f2) == ((self).f2) if self.f3 else self.f3)
            rhs118_ = d_743_v7_
            rhs119_ = True
            lhs62_ = (self).f5
            lhs63_ = default__.safeIndex(786, ((self).f5).length(0))
            lhs62_[lhs63_] = rhs117_
            d_743_v7_ = rhs118_
            r1 = rhs119_
            d_744_v8_: _dafny.Seq
            d_744_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
            d_744_v8_ = d_744_v8_
            d_745_v9_: C2
            nw91_ = C2()
            nw91_.ctor__(default__.safeDivisionInt((self).f4, (0) - ((self).f4)), True, (self).f1, d_744_v8_)
            d_745_v9_ = nw91_
            d_746_v10_: D0
            d_746_v10_ = D0_DC1((self).f4)
            source15_ = d_746_v10_
            if source15_.is_DC1:
                d_747___mcc_h0_ = source15_.cf1
                d_748_cf1_ = d_747___mcc_h0_
                d_749_v11_: _dafny.Array
                nw92_ = _dafny.Array(None, 3)
                nw92_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhq"))
                nw92_[int(1)] = (self).f2
                nw92_[int(2)] = (self).f2
                d_749_v11_ = nw92_
                d_749_v11_ = d_749_v11_
                d_744_v8_ = (self).f2
                d_750_v12_: _dafny.Array
                def lambda30_(d_751_i0_):
                    return (d_751_i0_) - (792)

                init15_ = lambda30_
                nw93_ = _dafny.Array(None, 28)
                for i0_15_ in range(nw93_.length(0)):
                    nw93_[i0_15_] = init15_(i0_15_)
                d_750_v12_ = nw93_
                d_752_v13_: C0
                nw94_ = C0()
                nw94_.ctor__(d_750_v12_)
                d_752_v13_ = nw94_
                index119_ = default__.safeIndex(207, (d_750_v12_).length(0))
                (d_750_v12_)[index119_] = default__.fm0(globalState)
                index120_ = default__.safeIndex(207, (d_750_v12_).length(0))
                (d_750_v12_)[index120_] = default__.safeDivisionInt(default__.fm0(globalState), len(d_743_v7_))
            elif True:
                d_753___mcc_h1_ = source15_.cf0
                d_754_cf0_ = d_753___mcc_h1_
                index121_ = default__.safeIndex(786, ((self).f5).length(0))
                ((self).f5)[index121_] = self.f3
                d_755_v14_: _dafny.Map
                d_755_v14_ = _dafny.Map({(d_745_v9_).f13: _dafny.CodePoint('l')})
                d_756_v15_: D8
                d_756_v15_ = D8_DC24(_dafny.CodePoint('q'), d_741_v5_)
                d_757_v16_: _dafny.Seq
                d_757_v16_ = _dafny.SeqWithoutIsStrInference([d_756_v15_])
                d_758_v17_: _dafny.Map
                d_758_v17_ = _dafny.Map({d_755_v14_: d_757_v16_})
                d_759_v18_: _dafny.Array
                nw95_ = _dafny.Array(None, 27)
                nw95_[int(0)] = p0
                nw95_[int(1)] = False
                nw95_[int(2)] = p0
                nw95_[int(3)] = True
                nw95_[int(4)] = p0
                nw95_[int(5)] = True
                nw95_[int(6)] = self.f3
                nw95_[int(7)] = ((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]
                nw95_[int(8)] = ((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]
                nw95_[int(9)] = self.f3
                nw95_[int(10)] = self.f3
                nw95_[int(11)] = True
                nw95_[int(12)] = self.f3
                nw95_[int(13)] = False
                nw95_[int(14)] = not(self.f3)
                nw95_[int(15)] = default__.fm3(d_754_cf0_, globalState)
                nw95_[int(16)] = p0
                nw95_[int(17)] = True
                nw95_[int(18)] = False
                nw95_[int(19)] = True
                nw95_[int(20)] = self.f3
                nw95_[int(21)] = self.f3
                nw95_[int(22)] = not(False)
                nw95_[int(23)] = ((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]
                nw95_[int(24)] = ((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]
                nw95_[int(25)] = self.f3
                nw95_[int(26)] = ((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]
                d_759_v18_ = nw95_
                d_760_v19_: D5
                d_760_v19_ = D5_DC15(573, p0, d_759_v18_)
                d_761_v20_: _dafny.Seq
                d_761_v20_ = _dafny.SeqWithoutIsStrInference([p0, False])
                d_762_v21_: _dafny.Map
                d_762_v21_ = _dafny.Map({False: len(d_761_v20_)})
                d_763_v22_: _dafny.Seq
                d_763_v22_ = _dafny.SeqWithoutIsStrInference([d_762_v21_, d_762_v21_, _dafny.Map({False: (self).f4})])
                d_764_v24_: _dafny.Set
                d_764_v24_ = _dafny.Set({d_762_v21_, d_762_v21_, d_762_v21_})
                d_765_v25_: str
                d_765_v25_ = _dafny.CodePoint('e')
                d_766_v26_: _dafny.Seq
                d_766_v26_ = _dafny.SeqWithoutIsStrInference([default__.fm2(((self).f5)[default__.safeIndex(786, ((self).f5).length(0))], globalState), d_744_v8_])
                d_767_v27_: _dafny.Set
                d_767_v27_ = _dafny.Set({(self).fm25(not((D2_DC6(d_765_v25_, p0, (self).fm13(globalState), (self).f4)).cf10), globalState), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_768_i1_ in range(default__.abs(560))]), (d_766_v26_)[default__.safeIndex((self).f4, len(d_766_v26_))]})
                d_769_v28_: _dafny.Array
                nw96_ = _dafny.Array(None, 23)
                nw96_[int(0)] = self.f3
                nw96_[int(1)] = self.f3
                nw96_[int(2)] = ((d_741_v5_)[((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]] if (((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]) in (d_741_v5_) else p0)
                nw96_[int(3)] = ((d_743_v7_)[(d_745_v9_).f13] if ((d_745_v9_).f13) in (d_743_v7_) else p0)
                nw96_[int(4)] = (not(False)) == (p0)
                nw96_[int(5)] = self.f3
                nw96_[int(6)] = (d_760_v19_).cf25
                nw96_[int(7)] = p0
                nw96_[int(8)] = p0
                nw96_[int(9)] = (((self).f5)[default__.safeIndex(786, ((self).f5).length(0))]) or (p0)
                nw96_[int(10)] = self.f3
                nw96_[int(11)] = (d_744_v8_) == (d_744_v8_)
                nw96_[int(12)] = self.f3
                nw96_[int(13)] = p0
                nw96_[int(14)] = not(self.f3)
                nw96_[int(15)] = ((self).f2) != (d_744_v8_)
                nw96_[int(16)] = p0
                nw96_[int(17)] = not((self).fm11(p0, globalState))
                def iife87_():
                    coll35_ = _dafny.Set()
                    compr_35_: _dafny.Map
                    for compr_35_ in (d_763_v22_).Elements:
                        d_770_v23_: _dafny.Map = compr_35_
                        if (d_770_v23_) in (d_763_v22_):
                            coll35_ = coll35_.union(_dafny.Set([d_770_v23_]))
                    return _dafny.Set(coll35_)
                nw96_[int(18)] = (d_764_v24_).issubset(iife87_()
                )
                nw96_[int(19)] = not (False) or (((self).f5)[default__.safeIndex(786, ((self).f5).length(0))])
                nw96_[int(20)] = (d_767_v27_).ispropersubset(d_767_v27_)
                nw96_[int(21)] = True
                nw96_[int(22)] = ((d_745_v9_).f13) >= ((0) - (len(d_742_v6_)))
                d_769_v28_ = nw96_
                index122_ = default__.safeIndex(786, ((self).f5).length(0))
                rhs120_ = d_758_v17_
                rhs121_ = ((self).f4) not in (d_743_v7_)
                rhs122_ = d_769_v28_
                lhs64_ = (self).f5
                lhs65_ = default__.safeIndex(786, ((self).f5).length(0))
                d_758_v17_ = rhs120_
                lhs64_[lhs65_] = rhs121_
                d_769_v28_ = rhs122_
                index123_ = default__.safeIndex(786, ((self).f5).length(0))
                rhs123_ = ((self).f4) > ((d_745_v9_).f13)
                rhs124_ = _dafny.Map({(d_745_v9_).f13: d_765_v25_})
                lhs66_ = (self).f5
                lhs67_ = default__.safeIndex(786, ((self).f5).length(0))
                lhs66_[lhs67_] = rhs123_
                d_755_v14_ = rhs124_
                d_771_v29_: _dafny.MultiSet
                d_771_v29_ = _dafny.MultiSet([(d_745_v9_).f13, (0) - ((d_745_v9_).f13)])
                d_772_v30_: _dafny.Map
                d_772_v30_ = _dafny.Map({((d_771_v29_)[433] if (433) in (d_771_v29_) else (0) - (d_754_cf0_)): _dafny.SeqWithoutIsStrInference([d_765_v25_ for d_773_i2_ in range(default__.abs(295))])})
                d_754_cf0_ = (((d_745_v9_).f13) + (len(d_772_v30_))) - ((self).f4)
        d_774_v31_: _dafny.Seq
        d_774_v31_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f4)])
        if (len(d_774_v31_)) <= (((self).f4) * ((self).f4)):
            if self.f3:
                d_775_v32_: D5
                d_775_v32_ = D5_DC15((self).f4, p0, (self).f5)
                d_776_v33_: _dafny.Map
                d_776_v33_ = _dafny.Map({(d_774_v31_)[default__.safeIndex((self).f4, len(d_774_v31_))]: (d_775_v32_).cf25})
                d_776_v33_ = (d_776_v33_).set((self).f4, self.f3)
                d_777_v34_: _dafny.Seq
                d_777_v34_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f4: p0})])
                d_778_v35_: str
                d_778_v35_ = _dafny.CodePoint('d')
                d_779_v36_: C2
                nw97_ = C2()
                nw97_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_776_v33_ for d_780_i3_ in range(default__.abs(-83))])) + (d_777_v34_)), p0, ((self).f1).set((self).fm14(p0, True, _dafny.CodePoint('y'), (self).f4, globalState), self.f3), ((self).f2).set(default__.safeIndex((0) - ((self).f4), len((self).f2)), d_778_v35_))
                d_779_v36_ = nw97_
                d_781_v37_: D6
                d_781_v37_ = D6_DC18((self).f4, (self).f4, p0, d_778_v35_, (0) - ((self).f4))
                pat_let_tv26_ = d_778_v35_
                def iife88_(_pat_let26_0):
                    def iife89_(d_782_dt__update__tmp_h0_):
                        def iife90_(_pat_let27_0):
                            def iife91_(d_783_dt__update_hcf32_h0_):
                                def iife92_(_pat_let28_0):
                                    def iife93_(d_784_dt__update_hcf30_h0_):
                                        def iife94_(_pat_let29_0):
                                            def iife95_(d_785_dt__update_hcf31_h0_):
                                                return D6_DC18((d_782_dt__update__tmp_h0_).cf29, d_784_dt__update_hcf30_h0_, d_785_dt__update_hcf31_h0_, d_783_dt__update_hcf32_h0_, (d_782_dt__update__tmp_h0_).cf33)
                                            return iife95_(_pat_let29_0)
                                        return iife94_(self.f3)
                                    return iife93_(_pat_let28_0)
                                return iife92_((self).f4)
                            return iife91_(_pat_let27_0)
                        return iife90_(pat_let_tv26_)
                    return iife89_(_pat_let26_0)
                d_778_v35_ = (iife88_(d_781_v37_)).cf32
                d_774_v31_ = d_774_v31_
                d_786_v38_: _dafny.Seq
                d_786_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqd"))
                d_786_v38_ = d_786_v38_
            elif True:
                d_787_v39_: _dafny.Array
                nw98_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_787_v39_ = nw98_
                index124_ = default__.safeIndex(302, (d_787_v39_).length(0))
                (d_787_v39_)[index124_] = (self).f5
                index125_ = default__.safeIndex(302, (d_787_v39_).length(0))
                (d_787_v39_)[index125_] = (self).f5
                d_788_v40_: int
                d_788_v40_ = 729
                d_789_v41_: _dafny.Array
                nw99_ = _dafny.Array(_dafny.Seq({}), 24)
                d_789_v41_ = nw99_
                d_790_v42_: _dafny.Array
                def lambda31_(d_791_p0_):
                    def lambda32_(d_792_i4_):
                        return default__.safeDivisionInt(d_792_i4_, (_dafny.MultiSet([self.f3, d_791_p0_, self.f3])).cardinality)

                    return lambda32_

                init16_ = lambda31_(p0)
                nw100_ = _dafny.Array(None, 18)
                for i0_16_ in range(nw100_.length(0)):
                    nw100_[i0_16_] = init16_(i0_16_)
                d_790_v42_ = nw100_
                index126_ = default__.safeIndex(318, (d_790_v42_).length(0))
                (d_790_v42_)[index126_] = default__.fm0(globalState)
                index127_ = default__.safeIndex(318, (d_790_v42_).length(0))
                rhs125_ = ((self).fm12(d_788_v40_, 760, d_788_v40_, self.f3, globalState)) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_793_i5_ in range(default__.abs(931))])))
                rhs126_ = (d_789_v41_ if self.f3 else (d_789_v41_ if p0 else d_789_v41_))
                rhs127_ = d_788_v40_
                lhs68_ = d_790_v42_
                lhs69_ = default__.safeIndex(318, (d_790_v42_).length(0))
                d_788_v40_ = rhs125_
                d_789_v41_ = rhs126_
                lhs68_[lhs69_] = rhs127_
                d_794_v43_: C0
                nw101_ = C0()
                nw101_.ctor__(d_790_v42_)
                d_794_v43_ = nw101_
                d_795_v44_: str
                d_795_v44_ = _dafny.CodePoint('r')
                d_796_v45_: _dafny.Map
                d_796_v45_ = _dafny.Map({self.f3: (self).f2})
                d_797_v46_: D3
                d_797_v46_ = D3_DC10(d_788_v40_, (d_790_v42_)[default__.safeIndex(318, (d_790_v42_).length(0))])
                d_798_v47_: D3
                d_798_v47_ = D3_DC11(d_797_v46_)
                d_799_v48_: _dafny.Map
                d_799_v48_ = _dafny.Map({d_798_v47_: len(_dafny.SeqWithoutIsStrInference([(self).f4 for d_800_i6_ in range(default__.abs(30))]))})
                d_795_v44_ = (d_794_v43_).fm29(d_796_v45_, d_799_v48_, len(_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (d_787_v39_)[default__.safeIndex(302, (d_787_v39_).length(0))], (self).f5, (d_787_v39_)[default__.safeIndex(302, (d_787_v39_).length(0))]])), globalState)
                d_801_v49_: _dafny.Seq
                d_801_v49_ = _dafny.SeqWithoutIsStrInference([self.f3])
                index128_ = default__.safeIndex(318, (d_790_v42_).length(0))
                (d_790_v42_)[index128_] = (0) - (((len(d_801_v49_) if p0 else d_788_v40_) if ((self).f2) <= ((self).f2) else (self).f4))
            r1 = True
            d_802_v50_: C1
            nw102_ = C1()
            nw102_.ctor__()
            d_802_v50_ = nw102_
            if self.f3:
                d_803_v51_: _dafny.Seq
                d_803_v51_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
                d_804_v52_: _dafny.Set
                d_804_v52_ = _dafny.Set({d_803_v51_})
                d_805_v53_: _dafny.Array
                nw103_ = _dafny.Array(None, 7)
                nw103_[int(0)] = (self).f4
                nw103_[int(1)] = 449
                nw103_[int(2)] = (0) - ((self).f4)
                nw103_[int(3)] = len(d_804_v52_)
                nw103_[int(4)] = (d_774_v31_)[default__.safeIndex((self).f4, len(d_774_v31_))]
                nw103_[int(5)] = (self).f4
                nw103_[int(6)] = (self).f4
                d_805_v53_ = nw103_
                d_806_v54_: C0
                nw104_ = C0()
                nw104_.ctor__((d_805_v53_ if p0 else d_805_v53_))
                d_806_v54_ = nw104_
                d_807_v55_: _dafny.Map
                d_807_v55_ = _dafny.Map({(self).f4: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaukntjia"))})
                d_807_v55_ = (d_807_v55_).set((0) - ((self).f4), (self).f2)
                r1 = not(p0)
                d_808_v56_: int
                d_808_v56_ = -984
                d_809_v57_: _dafny.MultiSet
                d_809_v57_ = _dafny.MultiSet([self.f3, p0])
                d_810_v58_: D3
                d_810_v58_ = D3_DC9(d_808_v56_, self.f3)
                d_808_v56_ = (((d_809_v57_).cardinality) * ((self).f4)) - ((len(d_803_v51_) if not((d_810_v58_).cf16) else d_808_v56_))
                d_808_v56_ = (self).f4
            elif True:
                d_811_v59_: int
                d_811_v59_ = 283
                d_811_v59_ = (0) - ((((self).f4) - ((self).f4)) + (d_811_v59_))
                d_812_v60_: _dafny.Array
                nw105_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_812_v60_ = nw105_
                d_813_v61_: _dafny.MultiSet
                d_813_v61_ = _dafny.MultiSet([d_811_v59_, (self).f4])
                index129_ = default__.safeIndex(136, (d_812_v60_).length(0))
                (d_812_v60_)[index129_] = d_813_v61_
                d_814_v62_: _dafny.MultiSet
                d_814_v62_ = _dafny.MultiSet([p0])
                d_815_v63_: str
                d_815_v63_ = _dafny.CodePoint('r')
                d_816_v64_: D2
                d_816_v64_ = D2_DC6(d_815_v63_, p0, p0, d_811_v59_)
                d_817_v65_: _dafny.Seq
                d_817_v65_ = _dafny.SeqWithoutIsStrInference([d_774_v31_, (d_774_v31_) + (default__.fm31(d_811_v59_, p0, len(_dafny.Set({p0})), (d_816_v64_).cf9, globalState))])
                d_818_v66_: D0
                d_818_v66_ = D0_DC1((self).f4)
                d_819_v67_: _dafny.Seq
                d_819_v67_ = _dafny.SeqWithoutIsStrInference([self.f3])
                index130_ = default__.safeIndex(136, (d_812_v60_).length(0))
                rhs128_ = (d_814_v62_).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, not(False)])))
                rhs129_ = (d_817_v65_)[default__.safeIndex((d_811_v59_) - (((d_814_v62_)[p0] if (p0) in (d_814_v62_) else (self).f4)), len(d_817_v65_))]
                rhs130_ = (((self).f4) < ((d_818_v66_).cf1)) in (_dafny.MultiSet(d_819_v67_))
                rhs131_ = d_813_v61_
                lhs70_ = self
                lhs71_ = self
                lhs72_ = d_812_v60_
                lhs73_ = default__.safeIndex(136, (d_812_v60_).length(0))
                lhs70_.f3 = rhs128_
                d_774_v31_ = rhs129_
                lhs71_.f3 = rhs130_
                lhs72_[lhs73_] = rhs131_
                d_820_v68_: _dafny.Array
                def lambda33_(d_821_p0_):
                    def lambda34_(d_822_i7_):
                        return d_821_p0_

                    return lambda34_

                init17_ = lambda33_(p0)
                nw106_ = _dafny.Array(None, 4)
                for i0_17_ in range(nw106_.length(0)):
                    nw106_[i0_17_] = init17_(i0_17_)
                d_820_v68_ = nw106_
                d_823_v69_: _dafny.Map
                d_823_v69_ = _dafny.Map({p0: (self).f5})
                d_820_v68_ = ((d_823_v69_)[(d_811_v59_) == ((self).f4)] if ((d_811_v59_) == ((self).f4)) in (d_823_v69_) else (self).f5)
                d_824_v70_: _dafny.Map
                d_824_v70_ = _dafny.Map({(self).f2: -549})
                d_824_v70_ = (d_824_v70_).set(_dafny.SeqWithoutIsStrInference([d_815_v63_ for d_825_i8_ in range(default__.abs(1))]), (d_774_v31_)[default__.safeIndex((self).f4, len(d_774_v31_))])
                (self).f3 = (((self).f4) * (d_811_v59_)) > (d_811_v59_)
            if (self.f3 if not(self.f3) else self.f3):
                d_826_v71_: _dafny.Array
                nw107_ = _dafny.Array(None, 1)
                nw107_[int(0)] = (self).f4
                d_826_v71_ = nw107_
                d_827_v72_: C0
                nw108_ = C0()
                nw108_.ctor__(d_826_v71_)
                d_827_v72_ = nw108_
                d_828_v73_: int
                d_828_v73_ = 949
                d_829_v74_: _dafny.Map
                d_829_v74_ = _dafny.Map({p0: (self).f4})
                d_830_v75_: _dafny.Map
                d_830_v75_ = _dafny.Map({(self).f4: not(self.f3)})
                d_828_v73_ = (0) - ((len(d_830_v75_) if not(((0) - (d_828_v73_)) >= ((0) - (((d_829_v74_)[p0] if (p0) in (d_829_v74_) else d_828_v73_)))) else 694))
                d_831_v76_: str
                d_831_v76_ = _dafny.CodePoint('f')
                d_831_v76_ = _dafny.CodePoint('k')
                (self).f3 = (d_828_v73_) == (default__.safeDivisionInt(d_828_v73_, d_828_v73_))
                d_828_v73_ = d_828_v73_
            elif True:
                d_832_v77_: _dafny.Array
                nw109_ = _dafny.Array(int(0), 15)
                d_832_v77_ = nw109_
                d_833_v78_: D5
                d_833_v78_ = D5_DC14(d_832_v77_)
                pat_let_tv27_ = d_832_v77_
                def iife96_(_pat_let30_0):
                    def iife97_(d_834_dt__update__tmp_h1_):
                        def iife98_(_pat_let31_0):
                            def iife99_(d_835_dt__update_hcf23_h0_):
                                return D5_DC14(d_835_dt__update_hcf23_h0_)
                            return iife99_(_pat_let31_0)
                        return iife98_(pat_let_tv27_)
                    return iife97_(_pat_let30_0)
                r0 = iife96_(d_833_v78_)
                r1 = (p0 if not(p0) else self.f3)
                d_836_v79_: C0
                nw110_ = C0()
                nw110_.ctor__(d_832_v77_)
                d_836_v79_ = nw110_
                d_837_v80_: _dafny.Seq
                d_837_v80_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpkogcfjm")))])
                d_837_v80_ = _dafny.SeqWithoutIsStrInference([p0, not(p0), not(not(False))])
                (self).f3 = self.f3
        elif True:
            d_838_v81_: _dafny.MultiSet
            d_838_v81_ = _dafny.MultiSet([False])
            d_839_v82_: str
            d_839_v82_ = _dafny.CodePoint('a')
            d_840_v83_: _dafny.Map
            d_840_v83_ = _dafny.Map({d_774_v31_: d_839_v82_})
            d_841_v84_: _dafny.Seq
            d_841_v84_ = _dafny.SeqWithoutIsStrInference([self.f3])
            d_842_v85_: D8
            d_842_v85_ = D8_DC23((self).f4, default__.fm0(globalState), self.f3)
            d_843_v86_: _dafny.Array
            nw111_ = _dafny.Array(None, 25)
            nw111_[int(0)] = not(p0)
            nw111_[int(1)] = self.f3
            nw111_[int(2)] = True
            nw111_[int(3)] = self.f3
            nw111_[int(4)] = not(p0)
            nw111_[int(5)] = not((d_838_v81_).isdisjoint(d_838_v81_))
            nw111_[int(6)] = self.f3
            nw111_[int(7)] = not(self.f3)
            nw111_[int(8)] = p0
            nw111_[int(9)] = p0
            nw111_[int(10)] = self.f3
            nw111_[int(11)] = self.f3
            nw111_[int(12)] = (((d_840_v83_)[default__.fm31(975, p0, (self).f4, d_839_v82_, globalState)] if (default__.fm31(975, p0, (self).f4, d_839_v82_, globalState)) in (d_840_v83_) else d_839_v82_)) not in ((self).f2)
            nw111_[int(13)] = (d_841_v84_) < (d_841_v84_)
            nw111_[int(14)] = (d_842_v85_).cf40
            nw111_[int(15)] = (_dafny.SeqWithoutIsStrInference([(self).f4])) < (d_774_v31_)
            nw111_[int(16)] = p0
            nw111_[int(17)] = p0
            nw111_[int(18)] = p0
            nw111_[int(19)] = True
            nw111_[int(20)] = not(p0)
            nw111_[int(21)] = self.f3
            nw111_[int(22)] = not(self.f3)
            nw111_[int(23)] = self.f3
            nw111_[int(24)] = self.f3
            d_843_v86_ = nw111_
            d_843_v86_ = d_843_v86_
            r1 = ((self).f4) != ((self).f4)
            d_844_v87_: int
            d_844_v87_ = 810
            d_845_v88_: D5
            d_845_v88_ = D5_DC15(d_844_v87_, self.f3, (self).f5)
            rhs132_ = default__.fm0(globalState)
            rhs133_ = (d_845_v88_).cf26
            rhs134_ = d_844_v87_
            d_844_v87_ = rhs132_
            d_843_v86_ = rhs133_
            d_844_v87_ = rhs134_
            d_846_v89_: _dafny.Array
            nw112_ = _dafny.Array(int(0), 7)
            d_846_v89_ = nw112_
            d_847_v90_: D0
            d_847_v90_ = D0_DC0(len((self).f2))
            index131_ = default__.safeIndex(139, (d_846_v89_).length(0))
            (d_846_v89_)[index131_] = (d_847_v90_).cf0
            index132_ = default__.safeIndex(139, (d_846_v89_).length(0))
            (d_846_v89_)[index132_] = d_844_v87_
            if self.f3:
                d_848_v92_: _dafny.Array
                def lambda35_(d_849_p0_, d_850_v82_, d_851_v87_):
                    def lambda36_(d_852_i9_):
                        def iife100_():
                            coll36_ = _dafny.Set()
                            compr_36_: _dafny.Map
                            for compr_36_ in (_dafny.Map({_dafny.Map({d_849_p0_: d_850_v82_}): d_851_v87_})).keys.Elements:
                                d_853_v91_: _dafny.Map = compr_36_
                                if (d_853_v91_) in (_dafny.Map({_dafny.Map({d_849_p0_: d_850_v82_}): d_851_v87_})):
                                    coll36_ = coll36_.union(_dafny.Set([d_853_v91_]))
                            return _dafny.Set(coll36_)
                        return iife100_()
                        

                    return lambda36_

                init18_ = lambda35_(p0, d_839_v82_, d_844_v87_)
                nw113_ = _dafny.Array(None, 26)
                for i0_18_ in range(nw113_.length(0)):
                    nw113_[i0_18_] = init18_(i0_18_)
                d_848_v92_ = nw113_
                d_854_v93_: _dafny.Map
                d_854_v93_ = _dafny.Map({self.f3: d_839_v82_})
                d_855_v94_: _dafny.Set
                d_855_v94_ = _dafny.Set({d_854_v93_, d_854_v93_})
                index133_ = default__.safeIndex(537, (d_848_v92_).length(0))
                (d_848_v92_)[index133_] = d_855_v94_
                index134_ = default__.safeIndex(537, (d_848_v92_).length(0))
                (d_848_v92_)[index134_] = d_855_v94_
                d_856_v95_: _dafny.Seq
                d_856_v95_ = _dafny.SeqWithoutIsStrInference([d_846_v89_])
                d_856_v95_ = _dafny.SeqWithoutIsStrInference([d_846_v89_])
                r1 = ((_dafny.SeqWithoutIsStrInference([(self).f4, (d_846_v89_)[default__.safeIndex(139, (d_846_v89_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([323 for d_857_i10_ in range(default__.abs(-759))]))) == (d_774_v31_)
                index135_ = default__.safeIndex(254, ((self).f5).length(0))
                ((self).f5)[index135_] = False
                d_858_v96_: _dafny.MultiSet
                d_858_v96_ = _dafny.MultiSet([d_841_v84_, d_841_v84_])
                index136_ = default__.safeIndex(254, ((self).f5).length(0))
                ((self).f5)[index136_] = ((not(self.f3)) or (self.f3) if p0 else (d_858_v96_).ispropersubset(d_858_v96_))
                d_859_v97_: _dafny.Array
                def lambda37_(d_860_i11_):
                    return ((self).f2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icurqkgo")))

                init19_ = lambda37_
                nw114_ = _dafny.Array(None, 9)
                for i0_19_ in range(nw114_.length(0)):
                    nw114_[i0_19_] = init19_(i0_19_)
                d_859_v97_ = nw114_
                index137_ = default__.safeIndex(592, (d_859_v97_).length(0))
                (d_859_v97_)[index137_] = (self).f2
                index138_ = default__.safeIndex(592, (d_859_v97_).length(0))
                (d_859_v97_)[index138_] = (((self).f2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvoau")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_861_i12_ in range(default__.abs(148))]))
            elif True:
                index139_ = default__.safeIndex(310, ((self).f5).length(0))
                ((self).f5)[index139_] = True
                index140_ = default__.safeIndex(310, ((self).f5).length(0))
                ((self).f5)[index140_] = p0
                d_862_v98_: _dafny.Array
                nw115_ = _dafny.Array(_dafny.CodePoint('D'), 5)
                d_862_v98_ = nw115_
                d_863_v99_: _dafny.Map
                d_863_v99_ = _dafny.Map({not(p0): default__.fm0(globalState)})
                d_864_v100_: _dafny.Map
                d_864_v100_ = _dafny.Map({len(d_863_v99_): d_862_v98_})
                d_865_v101_: _dafny.Seq
                d_865_v101_ = _dafny.SeqWithoutIsStrInference([d_774_v31_])
                d_866_v102_: _dafny.Array
                nw116_ = _dafny.Array(None, 2)
                nw116_[int(0)] = d_862_v98_
                nw116_[int(1)] = ((d_864_v100_)[len(d_865_v101_)] if (len(d_865_v101_)) in (d_864_v100_) else d_862_v98_)
                d_866_v102_ = nw116_
                index141_ = default__.safeIndex(369, (d_866_v102_).length(0))
                (d_866_v102_)[index141_] = d_862_v98_
                index142_ = default__.safeIndex(369, (d_866_v102_).length(0))
                (d_866_v102_)[index142_] = d_862_v98_
                d_867_v103_: _dafny.Map
                d_867_v103_ = _dafny.Map({default__.safeDivisionInt((self).f4, (d_846_v89_)[default__.safeIndex(139, (d_846_v89_).length(0))]): (self).f2})
                d_867_v103_ = d_867_v103_
                d_868_v104_: _dafny.Map
                d_868_v104_ = _dafny.Map({not(default__.fm3((d_846_v89_)[default__.safeIndex(139, (d_846_v89_).length(0))], globalState)): ((self).f5)[default__.safeIndex(310, ((self).f5).length(0))]})
                d_868_v104_ = (d_868_v104_).set(p0, ((self).f5)[default__.safeIndex(310, ((self).f5).length(0))])
                index143_ = default__.safeIndex(139, (d_846_v89_).length(0))
                (d_846_v89_)[index143_] = (self).f4
        if False:
            d_869_v105_: str
            d_869_v105_ = _dafny.CodePoint('g')
            d_870_v106_: _dafny.Array
            nw117_ = _dafny.Array(int(0), 11)
            d_870_v106_ = nw117_
            d_871_v107_: _dafny.Map
            d_871_v107_ = _dafny.Map({d_869_v105_: d_870_v106_})
            d_872_v108_: _dafny.Map
            d_872_v108_ = _dafny.Map({p0: ((d_871_v107_)[d_869_v105_] if (d_869_v105_) in (d_871_v107_) else d_870_v106_)})
            d_873_v109_: C0
            nw118_ = C0()
            nw118_.ctor__(((d_872_v108_)[self.f3] if (self.f3) in (d_872_v108_) else d_870_v106_))
            d_873_v109_ = nw118_
            d_874_v110_: _dafny.Map
            d_874_v110_ = _dafny.Map({len(d_774_v31_): (self).f4})
            d_875_v111_: _dafny.Map
            d_875_v111_ = _dafny.Map({p0: len((self).f2)})
            d_876_v112_: _dafny.Map
            d_876_v112_ = _dafny.Map({len(d_875_v111_): p0})
            d_877_v113_: D6
            d_877_v113_ = D6_DC16(d_876_v112_)
            d_878_v114_: _dafny.Map
            d_878_v114_ = _dafny.Map({d_877_v113_: _dafny.Map({124: (self).f4})})
            d_879_v115_: _dafny.Set
            d_879_v115_ = _dafny.Set({d_874_v110_, (d_874_v110_) | (d_874_v110_), ((d_878_v114_)[d_877_v113_] if (d_877_v113_) in (d_878_v114_) else d_874_v110_)})
            d_879_v115_ = d_879_v115_
            d_880_v116_: D3
            d_880_v116_ = D3_DC9((self).f4, p0)
            def iife101_(_pat_let32_0):
                def iife102_(d_881_dt__update__tmp_h2_):
                    def iife103_(_pat_let33_0):
                        def iife104_(d_882_dt__update_hcf16_h0_):
                            return D3_DC9((d_881_dt__update__tmp_h2_).cf15, d_882_dt__update_hcf16_h0_)
                        return iife104_(_pat_let33_0)
                    return iife103_(not(self.f3))
                return iife102_(_pat_let32_0)
            source16_ = iife101_(d_880_v116_)
            if source16_.is_DC9:
                d_883___mcc_h2_ = source16_.cf15
                d_884___mcc_h3_ = source16_.cf16
                d_885_cf16_ = d_884___mcc_h3_
                d_886_cf15_ = d_883___mcc_h2_
                index144_ = default__.safeIndex(210, ((self).f5).length(0))
                ((self).f5)[index144_] = (d_886_cf15_) >= (573)
                d_887_v117_: _dafny.MultiSet
                d_887_v117_ = _dafny.MultiSet([d_885_cf16_, p0, d_885_cf16_])
                index145_ = default__.safeIndex(210, ((self).f5).length(0))
                ((self).f5)[index145_] = (d_887_v117_) != (d_887_v117_)
                d_888_v118_: _dafny.Set
                d_888_v118_ = _dafny.Set({d_869_v105_})
                d_889_v119_: D6
                d_889_v119_ = D6_DC17(self.f3)
                d_890_v120_: _dafny.Map
                d_890_v120_ = _dafny.Map({(d_888_v118_).issubset(d_888_v118_): d_889_v119_})
                d_890_v120_ = _dafny.Map({(d_885_cf16_ if p0 else (self).fm11(((self).f5)[default__.safeIndex(210, ((self).f5).length(0))], globalState)): d_889_v119_})
                index146_ = default__.safeIndex(210, ((self).f5).length(0))
                rhs135_ = self.f3
                rhs136_ = ((0) - (default__.fm0(globalState)) if p0 else d_886_cf15_)
                lhs74_ = (self).f5
                lhs75_ = default__.safeIndex(210, ((self).f5).length(0))
                lhs74_[lhs75_] = rhs135_
                d_886_cf15_ = rhs136_
                d_891_v121_: _dafny.Array
                def lambda38_(d_892_v105_):
                    def lambda39_(d_893_i13_):
                        return d_892_v105_

                    return lambda39_

                init20_ = lambda38_(d_869_v105_)
                nw119_ = _dafny.Array(None, 1)
                for i0_20_ in range(nw119_.length(0)):
                    nw119_[i0_20_] = init20_(i0_20_)
                d_891_v121_ = nw119_
                index147_ = default__.safeIndex(708, (d_891_v121_).length(0))
                (d_891_v121_)[index147_] = d_869_v105_
                d_894_v122_: _dafny.Map
                d_894_v122_ = _dafny.Map({p0: ((self).f5)[default__.safeIndex(210, ((self).f5).length(0))]})
                d_895_v123_: D8
                d_895_v123_ = D8_DC24(d_869_v105_, d_894_v122_)
                pat_let_tv28_ = d_894_v122_
                index148_ = default__.safeIndex(708, (d_891_v121_).length(0))
                def iife105_(_pat_let34_0):
                    def iife106_(d_896_dt__update__tmp_h3_):
                        def iife107_(_pat_let35_0):
                            def iife108_(d_897_dt__update_hcf42_h0_):
                                return D8_DC24((d_896_dt__update__tmp_h3_).cf41, d_897_dt__update_hcf42_h0_)
                            return iife108_(_pat_let35_0)
                        return iife107_(pat_let_tv28_)
                    return iife106_(_pat_let34_0)
                (d_891_v121_)[index148_] = (iife105_(d_895_v123_)).cf41
            elif source16_.is_DC10:
                d_898___mcc_h4_ = source16_.cf17
                d_899___mcc_h5_ = source16_.cf18
                d_900_cf18_ = d_899___mcc_h5_
                d_901_cf17_ = d_898___mcc_h4_
                (self).f3 = default__.fm3((0) - ((0) - (d_901_cf17_)), globalState)
                index149_ = default__.safeIndex(519, ((self).f5).length(0))
                ((self).f5)[index149_] = not(p0)
                index150_ = default__.safeIndex(519, ((self).f5).length(0))
                ((self).f5)[index150_] = p0
                index151_ = default__.safeIndex(519, ((self).f5).length(0))
                ((self).f5)[index151_] = not(self.f3)
                d_902_v124_: _dafny.Seq
                d_902_v124_ = _dafny.SeqWithoutIsStrInference([self.f3])
                d_903_v125_: _dafny.Set
                d_903_v125_ = _dafny.Set({((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]})
                d_904_v126_: D2
                d_904_v126_ = D2_DC6(d_869_v105_, self.f3, p0, (self).f4)
                d_905_v127_: D2
                d_905_v127_ = D2_DC6((d_904_v126_).cf9, ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))], p0, 379)
                d_906_v128_: _dafny.Array
                nw120_ = _dafny.Array(None, 24)
                nw120_[int(0)] = (len(d_774_v31_)) < (d_901_cf17_)
                nw120_[int(1)] = ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]
                nw120_[int(2)] = self.f3
                nw120_[int(3)] = ((d_876_v112_)[(0) - ((0) - (d_901_cf17_))] if ((0) - ((0) - (d_901_cf17_))) in (d_876_v112_) else p0)
                nw120_[int(4)] = self.f3
                nw120_[int(5)] = (d_902_v124_)[default__.safeIndex((self).f4, len(d_902_v124_))]
                nw120_[int(6)] = not(not((d_903_v125_) != (d_903_v125_)))
                nw120_[int(7)] = ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]
                nw120_[int(8)] = ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]
                nw120_[int(9)] = ((d_876_v112_)[(self).f4] if ((self).f4) in (d_876_v112_) else p0)
                nw120_[int(10)] = False
                nw120_[int(11)] = self.f3
                nw120_[int(12)] = (d_905_v127_).cf10
                nw120_[int(13)] = True
                nw120_[int(14)] = p0
                nw120_[int(15)] = p0
                nw120_[int(16)] = ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]
                nw120_[int(17)] = ((self).f2) != ((self).f2)
                nw120_[int(18)] = p0
                nw120_[int(19)] = p0
                nw120_[int(20)] = ((self).f5)[default__.safeIndex(519, ((self).f5).length(0))]
                nw120_[int(21)] = (self.f3 if (d_902_v124_)[default__.safeIndex((self).f4, len(d_902_v124_))] else False)
                nw120_[int(22)] = p0
                nw120_[int(23)] = ((self).f4) == (321)
                d_906_v128_ = nw120_
                d_906_v128_ = d_906_v128_
            elif source16_.is_DC8:
                d_907___mcc_h6_ = source16_.cf14
                d_908_cf14_ = d_907___mcc_h6_
                d_909_v129_: _dafny.Map
                d_909_v129_ = _dafny.Map({(self).f5: p0})
                d_910_v130_: bool
                d_911_v131_: bool
                d_912_v132_: _dafny.Seq
                d_913_v133_: _dafny.Seq
                out17_: bool
                out18_: bool
                out19_: _dafny.Seq
                out20_: _dafny.Seq
                out17_, out18_, out19_, out20_ = default__.m0(d_909_v129_, globalState)
                d_910_v130_ = out17_
                d_911_v131_ = out18_
                d_912_v132_ = out19_
                d_913_v133_ = out20_
                d_914_v134_: C1
                nw121_ = C1()
                nw121_.ctor__()
                d_914_v134_ = nw121_
                d_915_v135_: C0
                nw122_ = C0()
                nw122_.ctor__(d_870_v106_)
                d_915_v135_ = nw122_
                index152_ = default__.safeIndex(19, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index152_] = (default__.fm0(globalState)) + ((self).f4)
                index153_ = default__.safeIndex(19, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index153_] = (((self).f4) * ((self).f4)) + ((self).f4)
            elif True:
                d_916___mcc_h7_ = source16_.cf19
                d_917_cf19_ = d_916___mcc_h7_
                d_918_v136_: _dafny.Seq
                d_918_v136_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                (self).f3 = (D6_DC18(334, (0) - (len(d_918_v136_)), p0, d_869_v105_, (self).f4)).cf31
                d_919_v137_: _dafny.Set
                d_919_v137_ = _dafny.Set({self.f3})
                d_920_v138_: _dafny.Seq
                d_920_v138_ = _dafny.SeqWithoutIsStrInference([d_919_v137_, (d_919_v137_) - (d_919_v137_)])
                d_921_v139_: _dafny.Map
                d_921_v139_ = _dafny.Map({p0: d_919_v137_})
                d_920_v138_ = _dafny.SeqWithoutIsStrInference([(((d_921_v139_)[self.f3] if (self.f3) in (d_921_v139_) else d_919_v137_)).intersection(d_919_v137_)])
                (self).f3 = p0
                d_922_v140_: int
                d_922_v140_ = -13
                d_923_v141_: D2
                d_923_v141_ = D2_DC5(d_875_v111_)
                d_924_v142_: D2
                d_924_v142_ = D2_DC7(d_923_v141_)
                d_922_v140_ = ((d_922_v140_) - ((self).f4)) + ((0) - (len(_dafny.SeqWithoutIsStrInference([d_924_v142_, d_924_v142_, d_924_v142_, d_924_v142_, d_924_v142_]))))
            d_925_v143_: _dafny.Set
            d_925_v143_ = _dafny.Set({p0})
            d_926_v144_: D3
            d_926_v144_ = D3_DC8(d_925_v143_)
            source17_ = d_926_v144_
            if source17_.is_DC9:
                d_927___mcc_h8_ = source17_.cf15
                d_928___mcc_h9_ = source17_.cf16
                d_929_cf16_ = d_928___mcc_h9_
                d_930_cf15_ = d_927___mcc_h8_
                d_930_cf15_ = d_930_cf15_
                d_931_v145_: D2
                d_931_v145_ = D2_DC5(d_875_v111_)
                d_932_v146_: _dafny.Map
                d_932_v146_ = _dafny.Map({(d_873_v109_).f14: (d_931_v145_ if p0 else D2_DC5(d_875_v111_))})
                d_932_v146_ = (d_932_v146_).set((d_873_v109_).f14, d_931_v145_)
                d_930_cf15_ = d_930_cf15_
                d_869_v105_ = d_869_v105_
            elif source17_.is_DC10:
                d_933___mcc_h10_ = source17_.cf17
                d_934___mcc_h11_ = source17_.cf18
                d_935_cf18_ = d_934___mcc_h11_
                d_936_cf17_ = d_933___mcc_h10_
                r1 = default__.fm3((self).f4, globalState)
                d_936_cf17_ = (0) - (d_936_cf17_)
                (self).f3 = p0
                d_937_v147_: D10
                d_937_v147_ = D10_DC28(d_874_v110_)
                d_935_cf18_ = (0) - (((d_935_cf18_) + ((self).f4)) - ((0) - (len((d_874_v110_ if self.f3 else (d_937_v147_).cf46)))))
            elif source17_.is_DC8:
                d_938___mcc_h12_ = source17_.cf14
                d_939_cf14_ = d_938___mcc_h12_
                d_940_v148_: _dafny.Set
                d_940_v148_ = _dafny.Set({(self).f4, (self).f4, len(_dafny.SeqWithoutIsStrInference([len(d_874_v110_) for d_941_i14_ in range(default__.abs(-725))]))})
                d_942_v149_: D6
                d_942_v149_ = D6_DC19((len(d_940_v148_) if self.f3 else (self).f4))
                d_942_v149_ = D6_DC19((self).f4)
                d_943_v150_: _dafny.Array
                nw123_ = _dafny.Array(D4.default()(), 8)
                d_943_v150_ = nw123_
                d_944_v151_: D9
                d_944_v151_ = D9_DC25(d_943_v150_)
                d_945_v152_: int
                d_945_v152_ = 740
                rhs137_ = True
                rhs138_ = not(False)
                rhs139_ = (self).fm13(globalState)
                rhs140_ = D9_DC25(d_943_v150_)
                rhs141_ = (0) - ((self).f4)
                lhs76_ = self
                r1 = rhs137_
                r1 = rhs138_
                lhs76_.f3 = rhs139_
                d_944_v151_ = rhs140_
                d_945_v152_ = rhs141_
                index154_ = default__.safeIndex(604, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index154_] = (0) - (d_945_v152_)
                index155_ = default__.safeIndex(604, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index155_] = (((0) - ((self).f4)) + (len(d_774_v31_))) + (default__.safeModuloInt(len(_dafny.Map({self.f3: self.f3})), d_945_v152_))
                d_946_v153_: _dafny.Map
                d_946_v153_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f")) if self.f3 else (_dafny.SeqWithoutIsStrInference([d_869_v105_ for d_947_i15_ in range(default__.abs(27))])).set(default__.safeIndex(default__.fm0(globalState), len(_dafny.SeqWithoutIsStrInference([d_869_v105_ for d_948_i15_ in range(default__.abs(27))]))), d_869_v105_)): (self).f2})
                d_946_v153_ = (d_946_v153_).set((self).f2, (self).f2)
            elif True:
                d_949___mcc_h13_ = source17_.cf19
                d_950_cf19_ = d_949___mcc_h13_
                d_951_v154_: _dafny.Map
                d_951_v154_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cckx")): (self).f4})
                index156_ = default__.safeIndex(913, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index156_] = (((d_951_v154_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rn"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rn"))) in (d_951_v154_) else (self).f4)) - ((self).f4)
                index157_ = default__.safeIndex(913, ((d_873_v109_).f14).length(0))
                ((d_873_v109_).f14)[index157_] = ((self).f4) + ((self).f4)
                r1 = (self).fm11(p0, globalState)
                d_952_v155_: _dafny.Seq
                d_952_v155_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndc"))
                d_952_v155_ = ((d_952_v155_) + ((self).f2)) + (((self).f2) + (default__.fm2(p0, globalState)))
                d_953_v156_: bool
                d_954_v157_: bool
                d_955_v158_: _dafny.Seq
                d_956_v159_: _dafny.Seq
                out21_: bool
                out22_: bool
                out23_: _dafny.Seq
                out24_: _dafny.Seq
                out21_, out22_, out23_, out24_ = default__.m0(_dafny.Map({(self).f5: p0}), globalState)
                d_953_v156_ = out21_
                d_954_v157_ = out22_
                d_955_v158_ = out23_
                d_956_v159_ = out24_
            d_957_v160_: int
            d_957_v160_ = 811
            d_957_v160_ = (self).fm12((self).f4, -323, d_957_v160_, self.f3, globalState)
        elif True:
            d_958_v161_: int
            d_958_v161_ = 356
            d_959_v162_: _dafny.Map
            d_959_v162_ = _dafny.Map({(self).f5: (self).f4})
            d_960_v163_: _dafny.Map
            d_960_v163_ = _dafny.Map({d_958_v161_: d_959_v162_})
            rhs142_ = default__.fm0(globalState)
            rhs143_ = (d_958_v161_) - (len(((d_960_v163_)[(self).f4] if ((self).f4) in (d_960_v163_) else d_959_v162_)))
            rhs144_ = ((self).f4) + (d_958_v161_)
            d_958_v161_ = rhs142_
            d_958_v161_ = rhs143_
            d_958_v161_ = rhs144_
            d_961_v164_: _dafny.Seq
            d_961_v164_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "soixpidow"))
            rhs145_ = d_961_v164_
            rhs146_ = (self).f4
            d_961_v164_ = rhs145_
            d_958_v161_ = rhs146_
            (self).f3 = self.f3
            d_962_v165_: _dafny.Seq
            d_962_v165_ = _dafny.SeqWithoutIsStrInference([p0])
            (self).f3 = (d_962_v165_) <= ((d_962_v165_) + (_dafny.SeqWithoutIsStrInference([p0, not(p0)])))
            d_963_v166_: _dafny.Array
            def lambda40_(d_964_i16_):
                return (d_964_i16_) + (-944)

            init21_ = lambda40_
            nw124_ = _dafny.Array(None, 8)
            for i0_21_ in range(nw124_.length(0)):
                nw124_[i0_21_] = init21_(i0_21_)
            d_963_v166_ = nw124_
            index158_ = default__.safeIndex(763, (d_963_v166_).length(0))
            (d_963_v166_)[index158_] = (453) + ((self).f4)
            index159_ = default__.safeIndex(763, (d_963_v166_).length(0))
            (d_963_v166_)[index159_] = (self).f4
        d_965_v167_: _dafny.Map
        d_965_v167_ = _dafny.Map({self.f3: False})
        d_966_v168_: _dafny.Map
        d_966_v168_ = _dafny.Map({(self).f4: d_965_v167_})
        d_967_v169_: _dafny.Array
        nw125_ = _dafny.Array(int(0), 10)
        d_967_v169_ = nw125_
        d_968_v170_: _dafny.Map
        d_968_v170_ = _dafny.Map({((d_966_v168_)[len((self).f2)] if (len((self).f2)) in (d_966_v168_) else _dafny.Map({not(self.f3): True})): d_967_v169_})
        d_969_v171_: _dafny.Seq
        d_969_v171_ = _dafny.SeqWithoutIsStrInference([d_965_v167_])
        d_968_v170_ = (d_968_v170_).set((d_965_v167_) | ((d_969_v171_)[default__.safeIndex((self).f4, len(d_969_v171_))]), d_967_v169_)
        d_970_v172_: _dafny.Map
        d_970_v172_ = _dafny.Map({(self).f4: p0})
        d_971_v173_: _dafny.Seq
        d_971_v173_ = _dafny.SeqWithoutIsStrInference([D6_DC17(self.f3)])
        d_972_v174_: _dafny.Map
        d_972_v174_ = _dafny.Map({len(d_971_v173_): 868})
        (self).f3 = ((d_970_v172_)[len(d_972_v174_)] if (len(d_972_v174_)) in (d_970_v172_) else self.f3)
        d_973_v175_: _dafny.Set
        d_973_v175_ = _dafny.Set({(self).f5, (self).f5})
        d_974_v176_: _dafny.Map
        d_974_v176_ = _dafny.Map({(d_973_v175_) - (d_973_v175_): default__.safeModuloInt(len(d_965_v167_), default__.fm0(globalState))})
        d_974_v176_ = (d_974_v176_).set(d_973_v175_, (len(d_972_v174_)) * ((self).f4))
        d_975_v177_: D5
        d_975_v177_ = D5_DC14(d_967_v169_)
        r0 = d_975_v177_
        d_976_v178_: _dafny.Set
        d_976_v178_ = _dafny.Set({(self).fm12((d_774_v31_)[default__.safeIndex((self).f4, len(d_774_v31_))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhfa"))), len(_dafny.SeqWithoutIsStrInference([(self).f4 for d_977_i17_ in range(default__.abs(867))])), self.f3, globalState), (self).f4, (self).f4, (0) - ((self).f4)})
        d_978_v179_: _dafny.Map
        d_978_v179_ = _dafny.Map({d_976_v178_: d_973_v175_})
        d_979_v180_: _dafny.Map
        d_979_v180_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): ((d_978_v179_)[d_976_v178_] if (d_976_v178_) in (d_978_v179_) else _dafny.Set({(self).f5}))})
        r1 = ((((d_979_v180_)[(self).f4] if ((self).f4) in (d_979_v180_) else d_973_v175_)).intersection(d_973_v175_)) == (d_973_v175_)
        return r0, r1

    def m11(self, p0, p1, globalState):
        r0: int = int(0)
        r0 = default__.fm0(globalState)
        d_980_i0_: int
        d_980_i0_ = 0
        with _dafny.label("7"):
            while self.f3:
                with _dafny.c_label("7"):
                    if (d_980_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_980_i0_ = (d_980_i0_) + (1)
                    d_981_v0_: _dafny.Seq
                    d_981_v0_ = _dafny.SeqWithoutIsStrInference([p1, (self).f4])
                    d_982_v1_: _dafny.Map
                    d_982_v1_ = _dafny.Map({_dafny.Set({self.f3, self.f3}): (self).f2})
                    d_983_v3_: _dafny.Map
                    d_983_v3_ = _dafny.Map({(self).f2: (0) - (len(d_981_v0_))})
                    d_984_v4_: _dafny.Seq
                    d_984_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm3(544, globalState)])
                    d_985_v5_: _dafny.Array
                    nw126_ = _dafny.Array(None, 19)
                    nw126_[int(0)] = p1
                    nw126_[int(1)] = (d_981_v0_)[default__.safeIndex((0) - (p1), len(d_981_v0_))]
                    nw126_[int(2)] = (self).f4
                    nw126_[int(3)] = (0) - (p1)
                    nw126_[int(4)] = ((0) - (p1) if self.f3 else len(d_981_v0_))
                    nw126_[int(5)] = len(d_981_v0_)
                    nw126_[int(6)] = p1
                    nw126_[int(7)] = (784) + (p1)
                    nw126_[int(8)] = (self).f4
                    nw126_[int(9)] = p1
                    nw126_[int(10)] = (self).f4
                    def iife109_():
                        coll37_ = _dafny.Set()
                        compr_37_: _dafny.Set
                        for compr_37_ in (d_982_v1_).keys.Elements:
                            d_986_v2_: _dafny.Set = compr_37_
                            if (d_986_v2_) in (d_982_v1_):
                                coll37_ = coll37_.union(_dafny.Set([d_986_v2_]))
                        return _dafny.Set(coll37_)
                    nw126_[int(11)] = len(iife109_()
                    )
                    nw126_[int(12)] = ((d_983_v3_)[(self).f2] if ((self).f2) in (d_983_v3_) else default__.fm0(globalState))
                    nw126_[int(13)] = p1
                    nw126_[int(14)] = (p1) - (p1)
                    nw126_[int(15)] = len(d_984_v4_)
                    nw126_[int(16)] = (self).f4
                    nw126_[int(17)] = (self).f4
                    nw126_[int(18)] = default__.safeDivisionInt(p1, (d_981_v0_)[default__.safeIndex(p1, len(d_981_v0_))])
                    d_985_v5_ = nw126_
                    index160_ = default__.safeIndex(510, (d_985_v5_).length(0))
                    (d_985_v5_)[index160_] = (self).f4
                    index161_ = default__.safeIndex(510, (d_985_v5_).length(0))
                    (d_985_v5_)[index161_] = (((0) - ((self).f4)) * (p1)) * (193)
                    d_987_v6_: _dafny.MultiSet
                    d_987_v6_ = _dafny.MultiSet([(d_985_v5_)[default__.safeIndex(510, (d_985_v5_).length(0))], (self).f4])
                    (self).f3 = ((default__.fm35(self.f3, globalState)).cf1) in (d_987_v6_)
                    (self).f3 = self.f3
                    r0 = (d_985_v5_)[default__.safeIndex(510, (d_985_v5_).length(0))]
                    pass
            pass
        d_988_v7_: str
        d_988_v7_ = _dafny.CodePoint('i')
        d_988_v7_ = (d_988_v7_ if not(self.f3) else d_988_v7_)
        d_989_v8_: _dafny.MultiSet
        d_989_v8_ = _dafny.MultiSet([p1, (self).f4, p1, (self).f4])
        d_990_v9_: D3
        d_990_v9_ = D3_DC9((d_989_v8_).cardinality, True)
        d_991_v10_: D6
        d_991_v10_ = D6_DC19((self).f4)
        d_992_v11_: _dafny.Map
        d_992_v11_ = _dafny.Map({d_990_v9_: (d_991_v10_).cf34})
        source18_ = default__.fm36((d_992_v11_) | (d_992_v11_), globalState)
        if source18_.is_DC17:
            d_993___mcc_h0_ = source18_.cf28
            d_994_cf28_ = d_993___mcc_h0_
            d_995_v12_: _dafny.Array
            nw127_ = _dafny.Array(None, 2)
            nw127_[int(0)] = _dafny.MultiSet([p1])
            nw127_[int(1)] = d_989_v8_
            d_995_v12_ = nw127_
            d_996_v13_: _dafny.Map
            d_996_v13_ = _dafny.Map({d_995_v12_: 267})
            d_996_v13_ = (d_996_v13_).set(d_995_v12_, (self).f4)
            d_997_v14_: _dafny.Array
            nw128_ = _dafny.Array(int(0), 10)
            d_997_v14_ = nw128_
            d_998_v15_: _dafny.Map
            d_998_v15_ = _dafny.Map({self.f3: d_997_v14_})
            d_999_v16_: _dafny.Seq
            d_999_v16_ = _dafny.SeqWithoutIsStrInference([d_997_v14_, d_997_v14_, d_997_v14_, d_997_v14_, d_997_v14_])
            d_1000_v17_: _dafny.Array
            nw129_ = _dafny.Array(None, 19)
            nw129_[int(0)] = d_997_v14_
            nw129_[int(1)] = d_997_v14_
            nw129_[int(2)] = ((d_998_v15_)[(d_990_v9_).cf16] if ((d_990_v9_).cf16) in (d_998_v15_) else d_997_v14_)
            nw129_[int(3)] = d_997_v14_
            nw129_[int(4)] = d_997_v14_
            nw129_[int(5)] = (d_999_v16_)[default__.safeIndex((self).f4, len(d_999_v16_))]
            nw129_[int(6)] = d_997_v14_
            nw129_[int(7)] = d_997_v14_
            nw129_[int(8)] = d_997_v14_
            nw129_[int(9)] = d_997_v14_
            nw129_[int(10)] = d_997_v14_
            nw129_[int(11)] = d_997_v14_
            nw129_[int(12)] = d_997_v14_
            nw129_[int(13)] = d_997_v14_
            nw129_[int(14)] = d_997_v14_
            nw129_[int(15)] = d_997_v14_
            nw129_[int(16)] = d_997_v14_
            nw129_[int(17)] = ((d_998_v15_)[not(False)] if (not(False)) in (d_998_v15_) else d_997_v14_)
            nw129_[int(18)] = d_997_v14_
            d_1000_v17_ = nw129_
            index162_ = default__.safeIndex(393, (d_1000_v17_).length(0))
            (d_1000_v17_)[index162_] = d_997_v14_
            index163_ = default__.safeIndex(393, (d_1000_v17_).length(0))
            rhs147_ = d_997_v14_
            rhs148_ = (self).f4
            rhs149_ = d_994_cf28_
            lhs77_ = d_1000_v17_
            lhs78_ = default__.safeIndex(393, (d_1000_v17_).length(0))
            lhs77_[lhs78_] = rhs147_
            r0 = rhs148_
            d_994_cf28_ = rhs149_
            r0 = 911
            d_1001_v18_: _dafny.Seq
            d_1001_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orqcpbsg"))
            d_1001_v18_ = (self).f2
        elif source18_.is_DC18:
            d_1002___mcc_h1_ = source18_.cf29
            d_1003___mcc_h2_ = source18_.cf30
            d_1004___mcc_h3_ = source18_.cf31
            d_1005___mcc_h4_ = source18_.cf32
            d_1006___mcc_h5_ = source18_.cf33
            d_1007_cf33_ = d_1006___mcc_h5_
            d_1008_cf32_ = d_1005___mcc_h4_
            d_1009_cf31_ = d_1004___mcc_h3_
            d_1010_cf30_ = d_1003___mcc_h2_
            d_1011_cf29_ = d_1002___mcc_h1_
            d_1012_v19_: _dafny.Seq
            d_1012_v19_ = _dafny.SeqWithoutIsStrInference([(self).f2, _dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1013_i1_ in range(default__.abs(313))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saoqwp")), (self).f2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ao"))])
            d_1009_cf31_ = (d_988_v7_) in ((d_1012_v19_)[default__.safeIndex((self).f4, len(d_1012_v19_))])
            d_1014_v20_: _dafny.MultiSet
            d_1014_v20_ = _dafny.MultiSet([d_1009_cf31_, d_1009_cf31_])
            d_988_v7_ = (d_988_v7_ if (d_1014_v20_).isdisjoint(_dafny.MultiSet([self.f3])) else d_988_v7_)
            d_1015_v21_: _dafny.Seq
            d_1015_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dm"))
            d_1015_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iljigkhs"))
            d_1016_v22_: _dafny.Seq
            d_1016_v22_ = _dafny.SeqWithoutIsStrInference([not(not(self.f3)), d_1009_cf31_, d_1009_cf31_])
            (self).f3 = (d_1015_v21_) == ((d_1015_v21_) + (((self).f2).set(default__.safeIndex(len(d_1016_v22_), len((self).f2)), d_988_v7_)))
        elif source18_.is_DC19:
            d_1017___mcc_h6_ = source18_.cf34
            d_1018_cf34_ = d_1017___mcc_h6_
            d_1019_v23_: _dafny.Seq
            d_1019_v23_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
            if (d_1019_v23_)[default__.safeIndex(-324, len(d_1019_v23_))]:
                (self).f3 = self.f3
                (self).f3 = (self.f3) == (self.f3)
                d_1020_v24_: _dafny.Map
                d_1020_v24_ = _dafny.Map({((self).f2).set(default__.safeIndex(d_1018_cf34_, len((self).f2)), ((self).f2)[default__.safeIndex((self).f4, len((self).f2))]): (self).f4})
                d_1020_v24_ = (d_1020_v24_).set((self).f2, (self).f4)
                d_1021_v25_: _dafny.Map
                d_1021_v25_ = _dafny.Map({self.f3: (d_1019_v23_)[default__.safeIndex(p1, len(d_1019_v23_))]})
                d_1021_v25_ = (d_1021_v25_).set(self.f3, False)
                d_1022_v26_: _dafny.Map
                d_1022_v26_ = _dafny.Map({d_1021_v25_: self.f3})
                d_1022_v26_ = (d_1022_v26_).set(d_1021_v25_, self.f3)
            elif True:
                d_1023_v27_: _dafny.Seq
                d_1023_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brwwhg"))
                d_1023_v27_ = (default__.fm2(self.f3, globalState)).set(default__.safeIndex(d_1018_cf34_, len(default__.fm2(self.f3, globalState))), d_988_v7_)
                d_1024_v28_: _dafny.Array
                nw130_ = _dafny.Array(int(0), 20)
                d_1024_v28_ = nw130_
                d_1025_v29_: C0
                nw131_ = C0()
                nw131_.ctor__(d_1024_v28_)
                d_1025_v29_ = nw131_
                index164_ = default__.safeIndex(576, ((d_1025_v29_).f14).length(0))
                ((d_1025_v29_).f14)[index164_] = (self).f4
                d_1026_v30_: D1
                d_1026_v30_ = D1_DC2(d_1023_v27_)
                index165_ = default__.safeIndex(576, ((d_1025_v29_).f14).length(0))
                rhs150_ = default__.safeDivisionInt((self).f4, p1)
                rhs151_ = self.f3
                rhs152_ = (d_1023_v27_) < ((d_1026_v30_).cf2)
                rhs153_ = (self).f4
                lhs79_ = self
                lhs80_ = self
                lhs81_ = (d_1025_v29_).f14
                lhs82_ = default__.safeIndex(576, ((d_1025_v29_).f14).length(0))
                r0 = rhs150_
                lhs79_.f3 = rhs151_
                lhs80_.f3 = rhs152_
                lhs81_[lhs82_] = rhs153_
                d_1027_v31_: _dafny.Seq
                d_1027_v31_ = _dafny.SeqWithoutIsStrInference([p1, d_1018_cf34_, (self).f4])
                d_1028_v32_: D4
                d_1028_v32_ = D4_DC13(self.f3, (d_1027_v31_) + (d_1027_v31_))
                d_1028_v32_ = d_1028_v32_
                d_1029_v33_: _dafny.Map
                d_1029_v33_ = _dafny.Map({self.f3: (self).f2})
                d_1030_v34_: _dafny.Map
                d_1030_v34_ = _dafny.Map({(d_989_v8_).cardinality: d_1023_v27_})
                d_1031_v36_: _dafny.Set
                d_1031_v36_ = _dafny.Set({True, not(self.f3), self.f3})
                d_1032_v37_: _dafny.Map
                d_1032_v37_ = _dafny.Map({d_1018_cf34_: d_1030_v34_})
                d_1033_v40_: D11
                d_1033_v40_ = D11_DC30(d_1030_v34_)
                d_1034_v41_: _dafny.Map
                d_1034_v41_ = _dafny.Map({self.f3: d_1030_v34_})
                d_1035_v43_: _dafny.Set
                d_1035_v43_ = _dafny.Set({p1})
                d_1036_v44_: _dafny.Array
                nw132_ = _dafny.Array(None, 25)
                nw132_[int(0)] = (_dafny.Map({len((self).f2): ((d_1029_v33_)[self.f3] if (self.f3) in (d_1029_v33_) else (self).f2)})).set((self).f4, (self).f2)
                def iife110_():
                    coll38_ = _dafny.Map()
                    compr_38_: int
                    for compr_38_ in (d_1027_v31_).Elements:
                        d_1037_v35_: int = compr_38_
                        if (d_1037_v35_) in (d_1027_v31_):
                            coll38_[(d_1037_v35_) + ((d_991_v10_).cf34)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivjrwehm"))
                    return _dafny.Map(coll38_)
                nw132_[int(1)] = (d_1030_v34_) | (iife110_()
                )
                nw132_[int(2)] = d_1030_v34_
                nw132_[int(3)] = d_1030_v34_
                nw132_[int(4)] = default__.fm37(((d_1025_v29_).f14)[default__.safeIndex(576, ((d_1025_v29_).f14).length(0))], p1, ((d_1025_v29_).f14)[default__.safeIndex(576, ((d_1025_v29_).f14).length(0))], d_1031_v36_, globalState)
                nw132_[int(5)] = (d_1030_v34_).set((_dafny.MultiSet([d_1018_cf34_])).cardinality, d_1023_v27_)
                nw132_[int(6)] = d_1030_v34_
                nw132_[int(7)] = d_1030_v34_
                nw132_[int(8)] = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqqsxqjwf"))})
                nw132_[int(9)] = d_1030_v34_
                nw132_[int(10)] = (((d_1032_v37_)[(self).f4] if ((self).f4) in (d_1032_v37_) else d_1030_v34_)) | (d_1030_v34_)
                nw132_[int(11)] = (d_1030_v34_) | (d_1030_v34_)
                nw132_[int(12)] = d_1030_v34_
                nw132_[int(13)] = d_1030_v34_
                nw132_[int(14)] = (d_1030_v34_) | (d_1030_v34_)
                nw132_[int(15)] = d_1030_v34_
                def iife111_():
                    def iife112_():
                        coll40_ = _dafny.Set()
                        compr_40_: int
                        for compr_40_ in _dafny.IntegerRange(83, 585):
                            d_1039_v39_: int = compr_40_
                            if ((83) <= (d_1039_v39_)) and ((d_1039_v39_) < (585)):
                                coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_1039_v39_, ((d_1025_v29_).f14)[default__.safeIndex(576, ((d_1025_v29_).f14).length(0))])]))
                        return _dafny.Set(coll40_)
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(232, 565):
                        d_1038_v38_: int = compr_39_
                        if ((232) <= (d_1038_v38_)) and ((d_1038_v38_) < (565)):
                            coll39_[(d_1038_v38_) - (len(iife112_()
                            ))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhbt"))
                    return _dafny.Map(coll39_)
                nw132_[int(16)] = iife111_()
                
                nw132_[int(17)] = (d_1030_v34_).set(p1, (self).f2)
                nw132_[int(18)] = (d_1030_v34_).set((self).fm12(d_1018_cf34_, d_1018_cf34_, len(_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1040_i2_ in range(default__.abs(114))])), self.f3, globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekvspg")))
                nw132_[int(19)] = ((d_1033_v40_).cf49) | (((d_1034_v41_)[self.f3] if (self.f3) in (d_1034_v41_) else d_1030_v34_))
                nw132_[int(20)] = d_1030_v34_
                nw132_[int(21)] = d_1030_v34_
                nw132_[int(22)] = d_1030_v34_
                def iife113_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in (d_1035_v43_).Elements:
                        d_1041_v42_: int = compr_41_
                        if (d_1041_v42_) in (d_1035_v43_):
                            coll41_[default__.safeModuloInt(d_1041_v42_, p1)] = (self).f2
                    return _dafny.Map(coll41_)
                nw132_[int(23)] = (iife113_()
                ) | (d_1030_v34_)
                nw132_[int(24)] = _dafny.Map({p1: (self).f2})
                d_1036_v44_ = nw132_
                index166_ = default__.safeIndex(895, (d_1036_v44_).length(0))
                (d_1036_v44_)[index166_] = (d_1030_v34_).set(p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onhe")))
                d_1042_v45_: _dafny.Array
                nw133_ = _dafny.Array(D10.default()(), 25)
                d_1042_v45_ = nw133_
                d_1043_v46_: _dafny.MultiSet
                d_1043_v46_ = _dafny.MultiSet([False])
                index167_ = default__.safeIndex(895, (d_1036_v44_).length(0))
                rhs154_ = -907
                rhs155_ = d_1030_v34_
                rhs156_ = (self).f2
                rhs157_ = d_1042_v45_
                rhs158_ = (d_1043_v46_).intersection((d_1043_v46_).intersection(_dafny.MultiSet([self.f3, self.f3])))
                lhs83_ = d_1036_v44_
                lhs84_ = default__.safeIndex(895, (d_1036_v44_).length(0))
                r0 = rhs154_
                lhs83_[lhs84_] = rhs155_
                d_1023_v27_ = rhs156_
                d_1042_v45_ = rhs157_
                d_1043_v46_ = rhs158_
            d_1044_v47_: _dafny.Set
            d_1044_v47_ = _dafny.Set({d_988_v7_, d_988_v7_})
            index168_ = default__.safeIndex(217, ((self).f5).length(0))
            ((self).f5)[index168_] = (d_1044_v47_).issubset(d_1044_v47_)
            index169_ = default__.safeIndex(217, ((self).f5).length(0))
            ((self).f5)[index169_] = (self.f3) == (self.f3)
            d_1045_v49_: _dafny.Map
            d_1045_v49_ = _dafny.Map({d_1018_cf34_: ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]})
            d_1046_v50_: _dafny.Map
            def iife114_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in _dafny.IntegerRange(117, -55):
                    d_1047_v48_: int = compr_42_
                    if ((117) <= (d_1047_v48_)) and ((d_1047_v48_) < (-55)):
                        coll42_[(d_1047_v48_) - (p1)] = True
                return _dafny.Map(coll42_)
            d_1046_v50_ = _dafny.Map({len((iife114_()
            ) | (d_1045_v49_)): p1})
            d_1048_v51_: D8
            d_1048_v51_ = D8_DC23(d_1018_cf34_, p1, ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))])
            d_1046_v50_ = (d_1046_v50_).set(d_1018_cf34_, ((self).f4 if ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))] else (d_1048_v51_).cf38))
            d_1049_v52_: _dafny.Map
            d_1049_v52_ = _dafny.Map({541: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sxbfhfglq"))})
            d_1050_v53_: D1
            d_1050_v53_ = D1_DC3(((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], len(((d_1049_v52_)[p1] if (p1) in (d_1049_v52_) else (self).f2)), len(((self).f2) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1051_i3_ in range(default__.abs(-886))]))))
            source19_ = d_1050_v53_
            if source19_.is_DC3:
                d_1052___mcc_h9_ = source19_.cf3
                d_1053___mcc_h10_ = source19_.cf4
                d_1054___mcc_h11_ = source19_.cf5
                d_1055_cf5_ = d_1054___mcc_h11_
                d_1056_cf4_ = d_1053___mcc_h10_
                d_1057_cf3_ = d_1052___mcc_h9_
                d_1058_v54_: _dafny.Map
                d_1058_v54_ = _dafny.Map({self.f3: d_1055_cf5_})
                d_1059_v55_: _dafny.Set
                d_1059_v55_ = _dafny.Set({((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]})
                index170_ = default__.safeIndex(217, ((self).f5).length(0))
                ((self).f5)[index170_] = ((0) - ((0) - (p1))) >= (((d_1058_v54_)[(self).fm26((self).f4, self.f3, ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], globalState)] if ((self).fm26((self).f4, self.f3, ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], globalState)) in (d_1058_v54_) else len(d_1059_v55_)))
                d_1060_v56_: _dafny.Seq
                d_1060_v56_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_1046_v50_ = (d_1046_v50_).set((d_1018_cf34_) - (len(_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]]))), default__.safeModuloInt((d_1060_v56_)[default__.safeIndex((self).f4, len(d_1060_v56_))], (0) - ((self).f4)))
                d_1061_v57_: _dafny.Set
                d_1061_v57_ = _dafny.Set({d_1018_cf34_, d_1018_cf34_})
                r0 = len(((_dafny.Set({d_1056_cf4_})) | (d_1061_v57_)).intersection(d_1061_v57_))
                d_1062_v58_: _dafny.MultiSet
                d_1062_v58_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1063_i4_ in range(default__.abs(216))])])
                index171_ = default__.safeIndex(217, ((self).f5).length(0))
                ((self).f5)[index171_] = ((self).f2) not in (d_1062_v58_)
            elif source19_.is_DC4:
                d_1064___mcc_h12_ = source19_.cf6
                d_1065___mcc_h13_ = source19_.cf7
                d_1066_cf7_ = d_1065___mcc_h13_
                d_1067_cf6_ = d_1064___mcc_h12_
                d_1068_v59_: _dafny.Seq
                d_1068_v59_ = _dafny.SeqWithoutIsStrInference([len((self).f2)])
                d_1067_cf6_ = not(((d_1068_v59_) + (d_1068_v59_)) < (d_1068_v59_))
                r0 = default__.safeModuloInt(p1, (d_989_v8_).cardinality)
                (self).f3 = not(self.f3)
                d_1069_v60_: C1
                nw134_ = C1()
                nw134_.ctor__()
                d_1069_v60_ = nw134_
            elif True:
                d_1070___mcc_h14_ = source19_.cf2
                d_1071_cf2_ = d_1070___mcc_h14_
                d_1072_v61_: _dafny.Map
                d_1072_v61_ = _dafny.Map({((self).f4) * (d_1018_cf34_): D11_DC30(_dafny.Map({len(_dafny.Map({self.f3: ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]})): d_1071_cf2_}))})
                d_1073_v62_: D11
                d_1073_v62_ = D11_DC30(d_1049_v52_)
                d_1072_v61_ = (d_1072_v61_).set(p1, d_1073_v62_)
                d_1074_v63_: _dafny.Array
                nw135_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_1074_v63_ = nw135_
                d_1075_v64_: _dafny.Array
                def lambda41_(d_1076_v7_):
                    def lambda42_(d_1077_i5_):
                        return d_1076_v7_

                    return lambda42_

                init22_ = lambda41_(d_988_v7_)
                nw136_ = _dafny.Array(None, 19)
                for i0_22_ in range(nw136_.length(0)):
                    nw136_[i0_22_] = init22_(i0_22_)
                d_1075_v64_ = nw136_
                index172_ = default__.safeIndex(448, (d_1074_v63_).length(0))
                (d_1074_v63_)[index172_] = d_1075_v64_
                index173_ = default__.safeIndex(448, (d_1074_v63_).length(0))
                (d_1074_v63_)[index173_] = d_1075_v64_
                d_1078_v65_: _dafny.Seq
                d_1078_v65_ = _dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])
                d_1079_v66_: _dafny.Array
                nw137_ = _dafny.Array(_dafny.MultiSet({}), 13)
                d_1079_v66_ = nw137_
                d_1080_v67_: _dafny.Map
                d_1080_v67_ = _dafny.Map({_dafny.Map({((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]: d_1078_v65_}): d_1079_v66_})
                d_1081_v68_: _dafny.Map
                d_1081_v68_ = _dafny.Map({False: d_1078_v65_})
                d_1080_v67_ = (d_1080_v67_).set(((d_1081_v68_).set(((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], d_1078_v65_)) | ((d_1081_v68_).set(((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], d_1078_v65_)), d_1079_v66_)
                d_1078_v65_ = _dafny.SeqWithoutIsStrInference([(self).f4, default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))], self.f3, ((self).f5)[default__.safeIndex(217, ((self).f5).length(0))]])), (self).f4)])
        elif source18_.is_DC16:
            d_1082___mcc_h7_ = source18_.cf27
            d_1083_cf27_ = d_1082___mcc_h7_
            (self).f3 = not(not(not(self.f3)))
            d_1084_v69_: _dafny.Array
            nw138_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_1084_v69_ = nw138_
            d_1085_v70_: _dafny.Set
            d_1085_v70_ = _dafny.Set({self.f3})
            d_1086_v71_: _dafny.Map
            d_1086_v71_ = _dafny.Map({d_988_v7_: d_1084_v69_})
            rhs159_ = (self.f3) not in ((d_1085_v70_).intersection(d_1085_v70_))
            rhs160_ = ((d_1086_v71_)[d_988_v7_] if (d_988_v7_) in (d_1086_v71_) else d_1084_v69_)
            rhs161_ = d_988_v7_
            rhs162_ = (default__.safeModuloInt(p1, (d_989_v8_).cardinality)) + ((0) - (len((self).f2)))
            lhs85_ = self
            lhs85_.f3 = rhs159_
            d_1084_v69_ = rhs160_
            d_988_v7_ = rhs161_
            r0 = rhs162_
            d_1087_v72_: C2
            nw139_ = C2()
            nw139_.ctor__(p1, self.f3, _dafny.Map({(self).f2: self.f3}), (self).f2)
            d_1087_v72_ = nw139_
            d_1088_v73_: D6
            d_1088_v73_ = D6_DC17(self.f3)
            index174_ = default__.safeIndex(282, ((self).f5).length(0))
            ((self).f5)[index174_] = (d_1088_v73_).cf28
            index175_ = default__.safeIndex(282, ((self).f5).length(0))
            ((self).f5)[index175_] = self.f3
        elif True:
            d_1089___mcc_h8_ = source18_.cf35
            d_1090_cf35_ = d_1089___mcc_h8_
            d_1091_v74_: _dafny.MultiSet
            d_1091_v74_ = _dafny.MultiSet([(self).f2])
            if (d_1091_v74_).isdisjoint(_dafny.MultiSet([(self).f2])):
                d_1092_v75_: _dafny.Seq
                d_1092_v75_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(p1, p1)])
                d_1092_v75_ = (d_1092_v75_) + (d_1092_v75_)
                d_1093_v76_: _dafny.Map
                d_1093_v76_ = _dafny.Map({p1: self.f3})
                d_1093_v76_ = (d_1093_v76_).set((self).f4, self.f3)
                d_1094_v77_: _dafny.Array
                def lambda43_(d_1095_p1_):
                    def lambda44_(d_1096_i6_):
                        return default__.safeDivisionInt(d_1096_i6_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f3])): d_1095_p1_})))

                    return lambda44_

                init23_ = lambda43_(p1)
                nw140_ = _dafny.Array(None, 5)
                for i0_23_ in range(nw140_.length(0)):
                    nw140_[i0_23_] = init23_(i0_23_)
                d_1094_v77_ = nw140_
                d_1097_v78_: C0
                nw141_ = C0()
                nw141_.ctor__(d_1094_v77_)
                d_1097_v78_ = nw141_
                d_1098_v79_: _dafny.Seq
                d_1098_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xui"))
                rhs163_ = d_1097_v78_
                rhs164_ = d_1098_v79_
                rhs165_ = (0) - ((0) - ((self).f4))
                d_1097_v78_ = rhs163_
                d_1098_v79_ = rhs164_
                r0 = rhs165_
                d_1099_v80_: _dafny.Seq
                d_1099_v80_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f2: self.f3})])
                d_1100_v81_: C2
                nw142_ = C2()
                nw142_.ctor__(p1, self.f3, (d_1099_v80_)[default__.safeIndex((self).f4, len(d_1099_v80_))], (self).f2)
                d_1100_v81_ = nw142_
                d_1101_v82_: C1
                nw143_ = C1()
                nw143_.ctor__()
                d_1101_v82_ = nw143_
            elif True:
                d_1102_v83_: D4
                d_1102_v83_ = D4_DC13(self.f3, _dafny.SeqWithoutIsStrInference([(self).f4 for d_1103_i7_ in range(default__.abs(738))]))
                d_1104_v84_: _dafny.Seq
                d_1104_v84_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_1105_v85_: _dafny.Set
                d_1105_v85_ = _dafny.Set({(d_1102_v83_ if not(self.f3) else d_1102_v83_), D4_DC13(self.f3, d_1104_v84_)})
                d_1106_v86_: _dafny.Map
                d_1106_v86_ = _dafny.Map({(D6_DC18(p1, (self).f4, self.f3, d_988_v7_, (0) - (p1))).cf31: d_1105_v85_})
                d_1105_v85_ = ((d_1106_v86_)[self.f3] if (self.f3) in (d_1106_v86_) else (d_1105_v85_).intersection(d_1105_v85_))
                rhs166_ = (p1) >= (p1)
                rhs167_ = self.f3
                rhs168_ = (self).fm12((0) - (default__.safeDivisionInt((self).f4, (0) - (len((self).fm14(self.f3, self.f3, d_988_v7_, len((self).f2), globalState))))), (self).f4, (_dafny.MultiSet([not(self.f3)])).cardinality, self.f3, globalState)
                lhs86_ = self
                lhs87_ = self
                lhs86_.f3 = rhs166_
                lhs87_.f3 = rhs167_
                r0 = rhs168_
                d_1107_v87_: _dafny.Array
                nw144_ = _dafny.Array(int(0), 4)
                d_1107_v87_ = nw144_
                d_1108_v88_: C0
                nw145_ = C0()
                nw145_.ctor__(d_1107_v87_)
                d_1108_v88_ = nw145_
                (self).f3 = self.f3
                d_1109_v89_: _dafny.Array
                def lambda45_(d_1110_v7_):
                    def lambda46_(d_1111_i8_):
                        return d_1110_v7_

                    return lambda46_

                init24_ = lambda45_(d_988_v7_)
                nw146_ = _dafny.Array(None, 24)
                for i0_24_ in range(nw146_.length(0)):
                    nw146_[i0_24_] = init24_(i0_24_)
                d_1109_v89_ = nw146_
                d_1112_v90_: _dafny.MultiSet
                d_1112_v90_ = _dafny.MultiSet([d_1109_v89_])
                (self).f3 = ((d_1112_v90_).set(d_1109_v89_, default__.abs((self).f4))).isdisjoint(d_1112_v90_)
            r0 = (0) - (len((((self).f2) + (_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1113_i9_ in range(default__.abs(258))]))).set(default__.safeIndex((self).f4, len(((self).f2) + (_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1114_i9_ in range(default__.abs(258))])))), d_988_v7_)))
            r0 = (self).f4
            if (self.f3) == (self.f3):
                d_1115_v91_: _dafny.Seq
                d_1115_v91_ = _dafny.SeqWithoutIsStrInference([(self).f4])
                d_1116_v92_: _dafny.Array
                nw147_ = _dafny.Array(None, 24)
                nw147_[int(0)] = d_988_v7_
                nw147_[int(1)] = d_988_v7_
                nw147_[int(2)] = d_988_v7_
                nw147_[int(3)] = d_988_v7_
                nw147_[int(4)] = d_988_v7_
                nw147_[int(5)] = d_988_v7_
                nw147_[int(6)] = d_988_v7_
                nw147_[int(7)] = d_988_v7_
                nw147_[int(8)] = d_988_v7_
                nw147_[int(9)] = d_988_v7_
                nw147_[int(10)] = d_988_v7_
                nw147_[int(11)] = d_988_v7_
                nw147_[int(12)] = d_988_v7_
                nw147_[int(13)] = d_988_v7_
                nw147_[int(14)] = d_988_v7_
                nw147_[int(15)] = d_988_v7_
                nw147_[int(16)] = d_988_v7_
                nw147_[int(17)] = default__.fm38(self.f3, False, (self).f4, globalState)
                nw147_[int(18)] = d_988_v7_
                nw147_[int(19)] = d_988_v7_
                nw147_[int(20)] = d_988_v7_
                nw147_[int(21)] = d_988_v7_
                nw147_[int(22)] = _dafny.CodePoint('o')
                nw147_[int(23)] = d_988_v7_
                d_1116_v92_ = nw147_
                d_1117_v93_: _dafny.Seq
                d_1117_v93_ = _dafny.SeqWithoutIsStrInference([d_1116_v92_, d_1116_v92_])
                d_1118_v94_: D4
                d_1118_v94_ = D4_DC13(False, (d_1115_v91_).set(default__.safeIndex((self).f4, len(d_1115_v91_)), len(d_1117_v93_)))
                d_1119_v95_: _dafny.Array
                nw148_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_1119_v95_ = nw148_
                d_1120_v96_: _dafny.Array
                nw149_ = _dafny.Array(int(0), 16)
                d_1120_v96_ = nw149_
                index176_ = default__.safeIndex(752, (d_1119_v95_).length(0))
                (d_1119_v95_)[index176_] = d_1120_v96_
                d_1121_v97_: _dafny.Map
                d_1121_v97_ = _dafny.Map({self.f3: len((self).f2)})
                index177_ = default__.safeIndex(752, (d_1119_v95_).length(0))
                rhs169_ = D4_DC13(False, d_1115_v91_)
                rhs170_ = self.f3
                rhs171_ = (p1 if self.f3 else ((d_1121_v97_)[not(self.f3)] if (not(self.f3)) in (d_1121_v97_) else (self).f4))
                rhs172_ = (d_988_v7_) in (_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1122_i10_ in range(default__.abs(116))]))
                rhs173_ = d_1120_v96_
                lhs88_ = self
                lhs89_ = self
                lhs90_ = d_1119_v95_
                lhs91_ = default__.safeIndex(752, (d_1119_v95_).length(0))
                d_1118_v94_ = rhs169_
                lhs88_.f3 = rhs170_
                r0 = rhs171_
                lhs89_.f3 = rhs172_
                lhs90_[lhs91_] = rhs173_
                d_1123_v98_: _dafny.Array
                nw150_ = _dafny.Array(False, 27)
                d_1123_v98_ = nw150_
                index178_ = default__.safeIndex(568, ((self).f5).length(0))
                ((self).f5)[index178_] = (d_989_v8_) != ((d_989_v8_).set(p1, default__.abs((self).f4)))
                index179_ = default__.safeIndex(568, ((self).f5).length(0))
                rhs174_ = (self).f4
                rhs175_ = default__.fm0(globalState)
                rhs176_ = d_1123_v98_
                rhs177_ = ((self).f4) != ((self).f4)
                rhs178_ = self.f3
                lhs92_ = self
                lhs93_ = (self).f5
                lhs94_ = default__.safeIndex(568, ((self).f5).length(0))
                r0 = rhs174_
                r0 = rhs175_
                d_1123_v98_ = rhs176_
                lhs92_.f3 = rhs177_
                lhs93_[lhs94_] = rhs178_
                (self).f3 = ((self).f5)[default__.safeIndex(568, ((self).f5).length(0))]
                r0 = ((d_989_v8_)[(self).f4] if ((self).f4) in (d_989_v8_) else p1)
                index180_ = default__.safeIndex(752, (d_1119_v95_).length(0))
                (d_1119_v95_)[index180_] = (d_1119_v95_)[default__.safeIndex(752, (d_1119_v95_).length(0))]
            elif True:
                d_1124_v99_: _dafny.Map
                d_1124_v99_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1125_i11_ in range(default__.abs(248))]): p1})
                d_1124_v99_ = (d_1124_v99_).set((self).f2, -930)
                (self).f3 = self.f3
                d_1126_v100_: _dafny.Map
                d_1126_v100_ = _dafny.Map({(self).f4: not(self.f3)})
                d_1127_v101_: _dafny.Seq
                d_1127_v101_ = _dafny.SeqWithoutIsStrInference([self.f3])
                d_1128_v102_: _dafny.Seq
                d_1128_v102_ = _dafny.SeqWithoutIsStrInference([len(d_1127_v101_)])
                d_1129_v103_: _dafny.Array
                nw151_ = _dafny.Array(None, 15)
                nw151_[int(0)] = len(d_1126_v100_)
                nw151_[int(1)] = p1
                nw151_[int(2)] = p1
                nw151_[int(3)] = default__.safeDivisionInt(len(d_1128_v102_), 991)
                nw151_[int(4)] = p1
                nw151_[int(5)] = p1
                nw151_[int(6)] = (0) - (default__.safeDivisionInt(p1, (0) - ((self).f4)))
                nw151_[int(7)] = (self).f4
                nw151_[int(8)] = ((self).f4) + (-855)
                nw151_[int(9)] = default__.safeModuloInt((self).f4, p1)
                nw151_[int(10)] = (default__.fm39(globalState)).cf15
                nw151_[int(11)] = ((0) - (p1)) + (p1)
                nw151_[int(12)] = (self).f4
                nw151_[int(13)] = p1
                nw151_[int(14)] = (self).f4
                d_1129_v103_ = nw151_
                index181_ = default__.safeIndex(636, (d_1129_v103_).length(0))
                (d_1129_v103_)[index181_] = ((self).f4) * (len(default__.fm2(self.f3, globalState)))
                index182_ = default__.safeIndex(636, (d_1129_v103_).length(0))
                (d_1129_v103_)[index182_] = p1
                d_1130_v104_: _dafny.Map
                d_1130_v104_ = _dafny.Map({self.f3: self.f3})
                d_1131_v105_: _dafny.Map
                def iife115_(_pat_let36_0):
                    def iife116_(d_1132_dt__update__tmp_h0_):
                        def iife117_(_pat_let37_0):
                            def iife118_(d_1133_dt__update_hcf25_h0_):
                                return D5_DC15((d_1132_dt__update__tmp_h0_).cf24, d_1133_dt__update_hcf25_h0_, (d_1132_dt__update__tmp_h0_).cf26)
                            return iife118_(_pat_let37_0)
                        return iife117_(self.f3)
                    return iife116_(_pat_let36_0)
                d_1131_v105_ = _dafny.Map({(iife115_(D5_DC15((self).f4, True, (self).f5))).cf25: d_1130_v104_})
                d_1134_v106_: _dafny.Set
                d_1134_v106_ = _dafny.Set({self.f3})
                d_1135_v107_: D6
                d_1135_v107_ = D6_DC18((self).f4, len(d_1134_v106_), self.f3, d_988_v7_, p1)
                d_1136_v108_: _dafny.MultiSet
                d_1136_v108_ = _dafny.MultiSet([True, (d_1135_v107_).cf31])
                d_1137_v109_: _dafny.Seq
                d_1137_v109_ = d_1127_v101_
                d_1138_v110_: _dafny.Map
                d_1138_v110_ = _dafny.Map({self.f3: d_1131_v105_})
                d_1139_v111_: _dafny.Seq
                d_1139_v111_ = _dafny.SeqWithoutIsStrInference([((d_1138_v110_)[self.f3] if (self.f3) in (d_1138_v110_) else _dafny.Map({self.f3: d_1130_v104_}))])
                rhs179_ = self.f3
                rhs180_ = ((_dafny.MultiSet((d_1137_v109_))) - (d_1136_v108_)).ispropersubset(d_1136_v108_)
                rhs181_ = ((d_1139_v111_)[default__.safeIndex(((d_989_v8_)[(self).f4] if ((self).f4) in (d_989_v8_) else 662), len(d_1139_v111_))] if self.f3 else d_1131_v105_)
                lhs95_ = self
                lhs96_ = self
                lhs95_.f3 = rhs179_
                lhs96_.f3 = rhs180_
                d_1131_v105_ = rhs181_
                d_1140_v112_: _dafny.Map
                d_1140_v112_ = _dafny.Map({(d_1128_v102_) + (d_1128_v102_): (self.f3) == (not(self.f3))})
                d_1140_v112_ = (d_1140_v112_).set((_dafny.SeqWithoutIsStrInference([(d_1129_v103_)[default__.safeIndex(636, (d_1129_v103_).length(0))]])) + (d_1128_v102_), self.f3)
        (self).f3 = self.f3
        d_1141_v113_: _dafny.Map
        d_1141_v113_ = _dafny.Map({self.f3: p1})
        d_1142_v114_: _dafny.Map
        d_1142_v114_ = _dafny.Map({self.f3: ((d_1141_v113_)[self.f3] if (self.f3) in (d_1141_v113_) else p1)})
        pat_let_tv29_ = p1
        def lambda47_(source21_):
            if source21_.is_DC6:
                d_1143___mcc_h18_ = source21_.cf9
                d_1144___mcc_h19_ = source21_.cf10
                d_1145___mcc_h20_ = source21_.cf11
                d_1146___mcc_h21_ = source21_.cf12
                d_1147_cf12_ = d_1146___mcc_h21_
                d_1148_cf11_ = d_1145___mcc_h20_
                d_1149_cf10_ = d_1144___mcc_h19_
                d_1150_cf9_ = d_1143___mcc_h18_
                return D9_DC27(d_1147_cf12_)
            elif source21_.is_DC5:
                d_1151___mcc_h22_ = source21_.cf8
                d_1152_cf8_ = d_1151___mcc_h22_
                return D9_DC27(len(_dafny.Set({self.f3, self.f3})))
            elif True:
                d_1153___mcc_h23_ = source21_.cf13
                d_1154_cf13_ = d_1153___mcc_h23_
                return D9_DC27(pat_let_tv29_)

        source20_ = lambda47_(D2_DC5(d_1141_v113_))
        if source20_.is_DC26:
            d_1155___mcc_h15_ = source20_.cf44
            d_1156_cf44_ = d_1155___mcc_h15_
            d_1157_v115_: _dafny.Array
            nw152_ = _dafny.Array(_dafny.Seq({}), 14)
            d_1157_v115_ = nw152_
            d_1158_v116_: _dafny.Seq
            d_1158_v116_ = _dafny.SeqWithoutIsStrInference([d_1157_v115_, d_1157_v115_, d_1157_v115_, d_1157_v115_])
            d_1157_v115_ = (d_1158_v116_)[default__.safeIndex(p1, len(d_1158_v116_))]
            if (d_1156_cf44_) != ((0) - ((p1) - ((0) - (default__.fm0(globalState))))):
                d_1159_v117_: D3
                d_1159_v117_ = D3_DC9(p1, (d_990_v9_).cf16)
                d_1160_v118_: D3
                d_1160_v118_ = D3_DC11(d_1159_v117_)
                pat_let_tv30_ = d_1159_v117_
                def iife119_(_pat_let38_0):
                    def iife120_(d_1161_dt__update__tmp_h1_):
                        def iife121_(_pat_let39_0):
                            def iife122_(d_1162_dt__update_hcf19_h0_):
                                return D3_DC11(d_1162_dt__update_hcf19_h0_)
                            return iife122_(_pat_let39_0)
                        return iife121_(pat_let_tv30_)
                    return iife120_(_pat_let38_0)
                d_1160_v118_ = iife119_(default__.fm40((d_989_v8_).cardinality, self.f3, 281, p1, globalState))
                d_1163_v119_: _dafny.Seq
                d_1163_v119_ = _dafny.SeqWithoutIsStrInference([(0) - (len((self).f2))])
                d_1164_v120_: C2
                nw153_ = C2()
                nw153_.ctor__(p1, not((401) != ((d_1163_v119_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1165_i12_ in range(default__.abs(396))])), len(d_1163_v119_))])), (self).f1, (self).f2)
                d_1164_v120_ = nw153_
                d_1166_v121_: _dafny.Set
                d_1166_v121_ = _dafny.Set({d_1156_cf44_, (d_1156_cf44_) - ((0) - ((0) - ((d_1164_v120_).f13))), d_1156_cf44_})
                d_1166_v121_ = d_1166_v121_
                r0 = default__.safeModuloInt((0) - (d_1156_cf44_), (self).f4)
                d_1167_v122_: _dafny.Map
                d_1167_v122_ = _dafny.Map({p1: self.f3})
                d_1168_v123_: D6
                d_1168_v123_ = D6_DC16(d_1167_v122_)
                d_1169_v124_: _dafny.Map
                d_1169_v124_ = _dafny.Map({D6_DC20(d_1168_v123_): (0) - ((self).f4)})
                d_1170_v125_: _dafny.MultiSet
                d_1170_v125_ = _dafny.MultiSet([d_1169_v124_])
                d_1171_v126_: D6
                d_1171_v126_ = D6_DC20(d_1168_v123_)
                d_1172_v127_: _dafny.Seq
                d_1172_v127_ = _dafny.SeqWithoutIsStrInference([d_1169_v124_, _dafny.Map({d_1171_v126_: (self).f4}), d_1169_v124_, d_1169_v124_, d_1169_v124_])
                d_1173_v128_: _dafny.Map
                d_1173_v128_ = _dafny.Map({self.f3: d_1172_v127_})
                d_1174_v129_: _dafny.Array
                nw154_ = _dafny.Array(None, 9)
                nw154_[int(0)] = d_1170_v125_
                nw154_[int(1)] = d_1170_v125_
                nw154_[int(2)] = d_1170_v125_
                nw154_[int(3)] = _dafny.MultiSet(((d_1173_v128_)[self.f3] if (self.f3) in (d_1173_v128_) else d_1172_v127_))
                nw154_[int(4)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1169_v124_, d_1169_v124_]))
                nw154_[int(5)] = default__.fm41((self).f4, d_1156_cf44_, globalState)
                nw154_[int(6)] = (d_1170_v125_).set(d_1169_v124_, default__.abs((d_1164_v120_).f13))
                nw154_[int(7)] = (_dafny.MultiSet(d_1172_v127_)).intersection(d_1170_v125_)
                nw154_[int(8)] = d_1170_v125_
                d_1174_v129_ = nw154_
                index183_ = default__.safeIndex(958, (d_1174_v129_).length(0))
                (d_1174_v129_)[index183_] = d_1170_v125_
                index184_ = default__.safeIndex(958, (d_1174_v129_).length(0))
                (d_1174_v129_)[index184_] = d_1170_v125_
            elif True:
                (self).f3 = ((self).f4) > (76)
                d_988_v7_ = _dafny.CodePoint('y')
                (self).f3 = False
                d_1175_v130_: _dafny.Map
                d_1175_v130_ = _dafny.Map({(self).f5: self.f3})
                d_1176_v131_: bool
                d_1177_v132_: bool
                d_1178_v133_: _dafny.Seq
                d_1179_v134_: _dafny.Seq
                out25_: bool
                out26_: bool
                out27_: _dafny.Seq
                out28_: _dafny.Seq
                out25_, out26_, out27_, out28_ = default__.m0((_dafny.Map({(self).f5: self.f3})) | ((d_1175_v130_).set((self).f5, self.f3)), globalState)
                d_1176_v131_ = out25_
                d_1177_v132_ = out26_
                d_1178_v133_ = out27_
                d_1179_v134_ = out28_
                r0 = default__.safeModuloInt((d_991_v10_).cf34, (0) - (len((self).f2)))
            d_1156_cf44_ = default__.fm0(globalState)
            d_1156_cf44_ = (self).f4
        elif source20_.is_DC27:
            d_1180___mcc_h16_ = source20_.cf45
            d_1181_cf45_ = d_1180___mcc_h16_
            d_1182_v135_: _dafny.Seq
            d_1182_v135_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_1183_v136_: _dafny.Set
            d_1183_v136_ = _dafny.Set({d_1182_v135_})
            d_1183_v136_ = ((d_1183_v136_ if self.f3 else d_1183_v136_)) | (d_1183_v136_)
            d_1184_v137_: _dafny.Seq
            d_1184_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gv"))
            d_1184_v137_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pceontkto"))
            d_1185_v138_: D11
            d_1185_v138_ = D11_DC31((self).f4)
            pat_let_tv31_ = p1
            def iife123_(_pat_let40_0):
                def iife124_(d_1186_dt__update__tmp_h2_):
                    def iife125_(_pat_let41_0):
                        def iife126_(d_1187_dt__update_hcf50_h0_):
                            return D11_DC31(d_1187_dt__update_hcf50_h0_)
                        return iife126_(_pat_let41_0)
                    return iife125_(pat_let_tv31_)
                return iife124_(_pat_let40_0)
            source22_ = iife123_(d_1185_v138_)
            if source22_.is_DC31:
                d_1188___mcc_h24_ = source22_.cf50
                d_1189_cf50_ = d_1188___mcc_h24_
                (self).f3 = self.f3
                d_1190_v139_: _dafny.Map
                d_1190_v139_ = _dafny.Map({default__.fm38(False, self.f3, ((d_989_v8_)[len(d_1142_v114_)] if (len(d_1142_v114_)) in (d_989_v8_) else (self).f4), globalState): (self).f4})
                d_1191_v140_: _dafny.Array
                nw155_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1191_v140_ = nw155_
                d_1192_v141_: _dafny.Array
                nw156_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_1192_v141_ = nw156_
                index185_ = default__.safeIndex(343, (d_1191_v140_).length(0))
                (d_1191_v140_)[index185_] = d_1192_v141_
                d_1193_v142_: _dafny.Map
                d_1193_v142_ = _dafny.Map({not (self.f3) or (True): d_1192_v141_})
                index186_ = default__.safeIndex(343, (d_1191_v140_).length(0))
                rhs182_ = (default__.fm42(globalState)) | (d_1190_v139_)
                rhs183_ = ((d_1193_v142_)[self.f3] if (self.f3) in (d_1193_v142_) else d_1192_v141_)
                lhs97_ = d_1191_v140_
                lhs98_ = default__.safeIndex(343, (d_1191_v140_).length(0))
                d_1190_v139_ = rhs182_
                lhs97_[lhs98_] = rhs183_
                d_1194_v143_: C2
                nw157_ = C2()
                nw157_.ctor__(689, ((0) - (d_1189_cf50_)) != (d_1189_cf50_), (self).f1, d_1184_v137_)
                d_1194_v143_ = nw157_
                d_1195_v144_: _dafny.Map
                d_1195_v144_ = _dafny.Map({d_988_v7_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlbjot"))})
                d_1184_v137_ = ((d_1195_v144_)[d_988_v7_] if (d_988_v7_) in (d_1195_v144_) else (self).f2)
            elif True:
                d_1196___mcc_h25_ = source22_.cf49
                d_1197_cf49_ = d_1196___mcc_h25_
                d_1198_v145_: _dafny.Array
                nw158_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_1198_v145_ = nw158_
                index187_ = default__.safeIndex(292, (d_1198_v145_).length(0))
                (d_1198_v145_)[index187_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                index188_ = default__.safeIndex(292, (d_1198_v145_).length(0))
                (d_1198_v145_)[index188_] = _dafny.SeqWithoutIsStrInference([d_988_v7_ for d_1199_i13_ in range(default__.abs(-345))])
                d_1200_v146_: _dafny.Array
                def lambda48_(d_1201_i14_):
                    return default__.safeModuloInt(d_1201_i14_, (self).f4)

                init25_ = lambda48_
                nw159_ = _dafny.Array(None, 20)
                for i0_25_ in range(nw159_.length(0)):
                    nw159_[i0_25_] = init25_(i0_25_)
                d_1200_v146_ = nw159_
                index189_ = default__.safeIndex(936, (d_1200_v146_).length(0))
                (d_1200_v146_)[index189_] = p1
                index190_ = default__.safeIndex(936, (d_1200_v146_).length(0))
                (d_1200_v146_)[index190_] = d_1181_cf45_
                d_1181_cf45_ = p1
                d_1202_v147_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.Seq({}), 24)
                d_1202_v147_ = nw160_
                d_1203_v148_: _dafny.Seq
                d_1203_v148_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3, self.f3, self.f3, self.f3])
                index191_ = default__.safeIndex(245, (d_1202_v147_).length(0))
                (d_1202_v147_)[index191_] = d_1203_v148_
                index192_ = default__.safeIndex(245, (d_1202_v147_).length(0))
                (d_1202_v147_)[index192_] = default__.fm43(globalState)
            d_1181_cf45_ = p1
        elif True:
            d_1204___mcc_h17_ = source20_.cf43
            d_1205_cf43_ = d_1204___mcc_h17_
            d_1206_v149_: _dafny.Set
            d_1206_v149_ = _dafny.Set({self.f3})
            if ((d_1206_v149_).intersection(d_1206_v149_)).isdisjoint((_dafny.Set({not(False), self.f3}) if self.f3 else d_1206_v149_)):
                d_1207_v150_: _dafny.MultiSet
                d_1207_v150_ = _dafny.MultiSet([not(self.f3), self.f3])
                (self).f3 = (_dafny.Set({d_1207_v150_})).ispropersubset(default__.fm44(globalState))
                d_1208_v151_: _dafny.Map
                d_1208_v151_ = _dafny.Map({((self).f2)[default__.safeIndex(default__.fm0(globalState), len((self).f2))]: (_dafny.CodePoint('x') if self.f3 else d_988_v7_)})
                d_1208_v151_ = (d_1208_v151_).set(d_988_v7_, (d_988_v7_ if self.f3 else d_988_v7_))
                d_1209_v152_: _dafny.Seq
                d_1209_v152_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1210_v153_: _dafny.Seq
                d_1210_v153_ = _dafny.SeqWithoutIsStrInference([d_1209_v152_])
                d_1211_v154_: D9
                d_1211_v154_ = D9_DC26((self).f4)
                d_1212_v155_: _dafny.Array
                nw161_ = _dafny.Array(None, 22)
                nw161_[int(0)] = 155
                nw161_[int(1)] = p1
                nw161_[int(2)] = p1
                nw161_[int(3)] = len(_dafny.Map({80: self.f3}))
                nw161_[int(4)] = 681
                nw161_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbqdrprp")))
                nw161_[int(6)] = p1
                nw161_[int(7)] = p1
                nw161_[int(8)] = (self).f4
                nw161_[int(9)] = (self).f4
                nw161_[int(10)] = (self).f4
                nw161_[int(11)] = p1
                nw161_[int(12)] = 161
                nw161_[int(13)] = (self).f4
                nw161_[int(14)] = len(d_1210_v153_)
                nw161_[int(15)] = (self).f4
                nw161_[int(16)] = p1
                nw161_[int(17)] = (self).f4
                nw161_[int(18)] = p1
                nw161_[int(19)] = (d_1211_v154_).cf44
                nw161_[int(20)] = p1
                nw161_[int(21)] = len(_dafny.SeqWithoutIsStrInference([self.f3]))
                d_1212_v155_ = nw161_
                d_1213_v156_: _dafny.Seq
                d_1213_v156_ = _dafny.SeqWithoutIsStrInference([d_1212_v155_])
                rhs184_ = ((self).f4) - ((self).f4)
                rhs185_ = (d_1213_v156_) != (d_1213_v156_)
                rhs186_ = (p1) <= (default__.safeDivisionInt((0) - (p1), p1))
                lhs99_ = self
                lhs100_ = self
                r0 = rhs184_
                lhs99_.f3 = rhs185_
                lhs100_.f3 = rhs186_
                d_1214_v157_: _dafny.Set
                d_1214_v157_ = _dafny.Set({p1})
                (self).f3 = (d_1214_v157_).isdisjoint(d_1214_v157_)
                d_1215_v158_: _dafny.Seq
                d_1215_v158_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
                rhs187_ = ((self).f4) < (default__.safeDivisionInt(p1, len(d_1215_v158_)))
                rhs188_ = (_dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])) + (d_1209_v152_)
                lhs101_ = self
                lhs101_.f3 = rhs187_
                d_1209_v152_ = rhs188_
            elif True:
                d_1216_v159_: C2
                nw162_ = C2()
                nw162_.ctor__((0) - ((self).f4), self.f3, (self).f1, (self).f2)
                d_1216_v159_ = nw162_
                d_1217_v160_: _dafny.MultiSet
                d_1217_v160_ = _dafny.MultiSet([(self).f2])
                d_1217_v160_ = default__.fm45((self).f4, self.f3, globalState)
                rhs189_ = default__.safeDivisionInt((self).f4, default__.safeModuloInt(len(d_1206_v149_), p1))
                rhs190_ = p1
                rhs191_ = (D10_DC29(self.f3, self.f3)).cf47
                rhs192_ = (default__.safeDivisionInt(p1, len(d_1206_v149_))) < ((self).f4)
                lhs102_ = self
                lhs103_ = self
                r0 = rhs189_
                r0 = rhs190_
                lhs102_.f3 = rhs191_
                lhs103_.f3 = rhs192_
                d_1218_v161_: D8
                d_1218_v161_ = D8_DC23(p1, (self).f4, self.f3)
                d_1219_v162_: _dafny.Map
                d_1219_v162_ = _dafny.Map({p1: self.f3})
                d_1220_v163_: _dafny.MultiSet
                d_1220_v163_ = _dafny.MultiSet([self.f3, self.f3, self.f3, self.f3])
                d_1221_v164_: _dafny.Set
                d_1221_v164_ = _dafny.Set({(self).f4, 951, (self).f4})
                d_1222_v165_: _dafny.Map
                d_1222_v165_ = _dafny.Map({d_1221_v164_: p1})
                d_1223_v166_: _dafny.Seq
                d_1223_v166_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - ((self).f4))])
                d_1224_v167_: C1
                nw163_ = C1()
                nw163_.ctor__()
                d_1224_v167_ = nw163_
                d_1225_v168_: _dafny.MultiSet
                d_1225_v168_ = _dafny.MultiSet([d_1224_v167_, d_1224_v167_])
                pat_let_tv32_ = p1
                pat_let_tv33_ = d_1222_v165_
                pat_let_tv34_ = p1
                pat_let_tv35_ = d_1216_v159_
                pat_let_tv36_ = d_1216_v159_
                d_1226_v169_: _dafny.Array
                nw164_ = _dafny.Array(None, 23)
                nw164_[int(0)] = d_1218_v161_
                nw164_[int(1)] = D8_DC23(p1, len((d_1219_v162_).set(p1, self.f3)), self.f3)
                nw164_[int(2)] = D8_DC23((d_1220_v163_).cardinality, p1, default__.fm3(p1, globalState))
                nw164_[int(3)] = D8_DC23((d_1216_v159_).f13, p1, self.f3)
                nw164_[int(4)] = d_1218_v161_
                nw164_[int(5)] = d_1218_v161_
                def iife127_(_pat_let42_0):
                    def iife128_(d_1227_dt__update__tmp_h3_):
                        def iife129_(_pat_let43_0):
                            def iife130_(d_1228_dt__update_hcf38_h0_):
                                def iife131_(_pat_let44_0):
                                    def iife132_(d_1229_dt__update_hcf40_h0_):
                                        return D8_DC23(d_1228_dt__update_hcf38_h0_, (d_1227_dt__update__tmp_h3_).cf39, d_1229_dt__update_hcf40_h0_)
                                    return iife132_(_pat_let44_0)
                                return iife131_(self.f3)
                            return iife130_(_pat_let43_0)
                        return iife129_(pat_let_tv32_)
                    return iife128_(_pat_let42_0)
                nw164_[int(6)] = iife127_(d_1218_v161_)
                def iife133_(_pat_let45_0):
                    def iife134_(d_1230_dt__update__tmp_h4_):
                        def iife135_(_pat_let46_0):
                            def iife136_(d_1231_dt__update_hcf38_h1_):
                                def iife137_(_pat_let47_0):
                                    def iife138_(d_1232_dt__update_hcf40_h1_):
                                        return D8_DC23(d_1231_dt__update_hcf38_h1_, (d_1230_dt__update__tmp_h4_).cf39, d_1232_dt__update_hcf40_h1_)
                                    return iife138_(_pat_let47_0)
                                return iife137_(self.f3)
                            return iife136_(_pat_let46_0)
                        return iife135_(len(pat_let_tv33_))
                    return iife134_(_pat_let45_0)
                nw164_[int(7)] = iife133_(d_1218_v161_)
                nw164_[int(8)] = d_1218_v161_
                nw164_[int(9)] = d_1218_v161_
                nw164_[int(10)] = d_1218_v161_
                nw164_[int(11)] = d_1218_v161_
                nw164_[int(12)] = D8_DC23(714, p1, self.f3)
                nw164_[int(13)] = d_1218_v161_
                def iife139_(_pat_let48_0):
                    def iife140_(d_1233_dt__update__tmp_h5_):
                        def iife141_(_pat_let49_0):
                            def iife142_(d_1234_dt__update_hcf40_h2_):
                                def iife143_(_pat_let50_0):
                                    def iife144_(d_1235_dt__update_hcf39_h0_):
                                        return D8_DC23((d_1233_dt__update__tmp_h5_).cf38, d_1235_dt__update_hcf39_h0_, d_1234_dt__update_hcf40_h2_)
                                    return iife144_(_pat_let50_0)
                                return iife143_(pat_let_tv34_)
                            return iife142_(_pat_let49_0)
                        return iife141_(self.f3)
                    return iife140_(_pat_let48_0)
                nw164_[int(14)] = iife139_(D8_DC23((_dafny.MultiSet([d_989_v8_])).cardinality, (self).f4, self.f3))
                nw164_[int(15)] = D8_DC23(len(d_1223_v166_), (self).f4, self.f3)
                def iife145_(_pat_let51_0):
                    def iife146_(d_1236_dt__update__tmp_h6_):
                        def iife147_(_pat_let52_0):
                            def iife148_(d_1237_dt__update_hcf40_h3_):
                                return D8_DC23((d_1236_dt__update__tmp_h6_).cf38, (d_1236_dt__update__tmp_h6_).cf39, d_1237_dt__update_hcf40_h3_)
                            return iife148_(_pat_let52_0)
                        return iife147_(False)
                    return iife146_(_pat_let51_0)
                nw164_[int(16)] = iife145_(d_1218_v161_)
                nw164_[int(17)] = d_1218_v161_
                nw164_[int(18)] = d_1218_v161_
                nw164_[int(19)] = d_1218_v161_
                def iife149_(_pat_let53_0):
                    def iife150_(d_1238_dt__update__tmp_h7_):
                        def iife151_(_pat_let54_0):
                            def iife152_(d_1239_dt__update_hcf38_h2_):
                                def iife153_(_pat_let55_0):
                                    def iife154_(d_1240_dt__update_hcf40_h4_):
                                        return D8_DC23(d_1239_dt__update_hcf38_h2_, (d_1238_dt__update__tmp_h7_).cf39, d_1240_dt__update_hcf40_h4_)
                                    return iife154_(_pat_let55_0)
                                return iife153_(self.f3)
                            return iife152_(_pat_let54_0)
                        return iife151_((pat_let_tv35_).f13)
                    return iife150_(_pat_let53_0)
                nw164_[int(20)] = (d_1218_v161_ if self.f3 else iife149_(d_1218_v161_))
                def iife155_(_pat_let56_0):
                    def iife156_(d_1241_dt__update__tmp_h8_):
                        def iife157_(_pat_let57_0):
                            def iife158_(d_1242_dt__update_hcf38_h3_):
                                return D8_DC23(d_1242_dt__update_hcf38_h3_, (d_1241_dt__update__tmp_h8_).cf39, (d_1241_dt__update__tmp_h8_).cf40)
                            return iife158_(_pat_let57_0)
                        return iife157_((pat_let_tv36_).f13)
                    return iife156_(_pat_let56_0)
                nw164_[int(21)] = (iife155_(d_1218_v161_) if not(self.f3) else d_1218_v161_)
                nw164_[int(22)] = D8_DC23(((d_1225_v168_)[d_1224_v167_] if (d_1224_v167_) in (d_1225_v168_) else (d_1216_v159_).f13), (d_1216_v159_).f13, self.f3)
                d_1226_v169_ = nw164_
                index193_ = default__.safeIndex(266, (d_1226_v169_).length(0))
                (d_1226_v169_)[index193_] = d_1218_v161_
                index194_ = default__.safeIndex(266, (d_1226_v169_).length(0))
                (d_1226_v169_)[index194_] = D8_DC23((d_1216_v159_).f13, (p1 if self.f3 else (0) - ((d_1216_v159_).f13)), self.f3)
                d_1243_v170_: _dafny.Map
                d_1243_v170_ = _dafny.Map({(self).f2: (self).f4})
                d_1244_v171_: _dafny.MultiSet
                d_1244_v171_ = _dafny.MultiSet([d_989_v8_, _dafny.MultiSet([len(d_1243_v170_), (self).f4])])
                d_1244_v171_ = default__.fm46(globalState)
            (self).f3 = self.f3
            d_1245_v172_: _dafny.Seq
            d_1245_v172_ = _dafny.SeqWithoutIsStrInference([not(self.f3)])
            d_1246_v173_: _dafny.Seq
            d_1246_v173_ = _dafny.SeqWithoutIsStrInference([d_1245_v172_])
            r0 = (D9_DC27(len(d_1246_v173_))).cf45
            d_1247_v174_: _dafny.Array
            nw165_ = _dafny.Array(_dafny.CodePoint('D'), 21)
            d_1247_v174_ = nw165_
            index195_ = default__.safeIndex(330, (d_1247_v174_).length(0))
            (d_1247_v174_)[index195_] = _dafny.CodePoint('l')
            d_1248_v175_: _dafny.MultiSet
            d_1248_v175_ = _dafny.MultiSet([d_1245_v172_])
            d_1249_v176_: D11
            d_1249_v176_ = D11_DC31((0) - ((d_1248_v175_).cardinality))
            d_1250_v177_: _dafny.Seq
            d_1250_v177_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f4)])
            d_1251_v178_: _dafny.Set
            d_1251_v178_ = _dafny.Set({p1, (0) - ((0) - (p1))})
            d_1252_v179_: _dafny.Map
            d_1252_v179_ = _dafny.Map({(self).f4: _dafny.CodePoint('h')})
            d_1253_v181_: _dafny.Array
            nw166_ = _dafny.Array(None, 12)
            nw166_[int(0)] = p1
            nw166_[int(1)] = (d_1250_v177_)[default__.safeIndex((0) - ((self).f4), len(d_1250_v177_))]
            nw166_[int(2)] = p1
            nw166_[int(3)] = (self).f4
            nw166_[int(4)] = p1
            nw166_[int(5)] = len(d_1251_v178_)
            nw166_[int(6)] = len(d_1252_v179_)
            nw166_[int(7)] = len(d_1252_v179_)
            nw166_[int(8)] = p1
            nw166_[int(9)] = len(default__.fm2(self.f3, globalState))
            def iife159_():
                coll43_ = _dafny.Set()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(344, -998):
                    d_1254_v180_: int = compr_43_
                    if ((344) <= (d_1254_v180_)) and ((d_1254_v180_) < (-998)):
                        coll43_ = coll43_.union(_dafny.Set([(d_1254_v180_) - (p1)]))
                return _dafny.Set(coll43_)
            nw166_[int(10)] = len(iife159_()
            )
            nw166_[int(11)] = (self).f4
            d_1253_v181_ = nw166_
            d_1255_v182_: _dafny.Map
            d_1255_v182_ = _dafny.Map({d_1249_v176_: d_1253_v181_})
            d_1256_v183_: _dafny.Seq
            d_1256_v183_ = _dafny.SeqWithoutIsStrInference([d_1250_v177_])
            d_1257_v184_: D6
            d_1257_v184_ = D6_DC18((self).f4, (self).f4, default__.fm3(len(d_1256_v183_), globalState), d_988_v7_, (self).f4)
            index196_ = default__.safeIndex(330, (d_1247_v174_).length(0))
            rhs193_ = (d_1255_v182_) != ((_dafny.Map({d_1249_v176_: d_1253_v181_})) | (d_1255_v182_))
            rhs194_ = (d_1257_v184_).cf32
            rhs195_ = self.f3
            lhs104_ = self
            lhs105_ = d_1247_v174_
            lhs106_ = default__.safeIndex(330, (d_1247_v174_).length(0))
            lhs107_ = self
            lhs104_.f3 = rhs193_
            lhs105_[lhs106_] = rhs194_
            lhs107_.f3 = rhs195_
        r0 = 349
        return r0


class C4(T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    def ctor__(self, f3, f1, f2):
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm12(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt(642, len(((self).f2 if not(self.f3) else ((self).f2).set(default__.safeIndex(356, len((self).f2)), _dafny.CodePoint('f')))))

    def fm11(self, p0, globalState):
        return not(not (not(self.f3)) or (self.f3))

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        hi5_ = p0
        for d_1258_i0_ in range((D2_DC6(default__.fm24(self.f3, self.f3, p0, globalState), self.f3, True, 755)).cf12, hi5_):
            r2 = len((self).f2)
            d_1259_v0_: _dafny.Map
            d_1259_v0_ = _dafny.Map({p0: self.f3})
            d_1260_v1_: _dafny.MultiSet
            d_1260_v1_ = _dafny.MultiSet([self.f3, self.f3])
            d_1261_v2_: _dafny.Seq
            d_1261_v2_ = _dafny.SeqWithoutIsStrInference([self.f3])
            d_1262_v3_: _dafny.Array
            nw167_ = _dafny.Array(None, 29)
            nw167_[int(0)] = (self).fm11(self.f3, globalState)
            nw167_[int(1)] = self.f3
            nw167_[int(2)] = (d_1261_v2_)[default__.safeIndex(p0, len(d_1261_v2_))]
            nw167_[int(3)] = self.f3
            nw167_[int(4)] = self.f3
            nw167_[int(5)] = self.f3
            nw167_[int(6)] = self.f3
            nw167_[int(7)] = self.f3
            nw167_[int(8)] = True
            nw167_[int(9)] = self.f3
            nw167_[int(10)] = self.f3
            nw167_[int(11)] = self.f3
            nw167_[int(12)] = self.f3
            nw167_[int(13)] = self.f3
            nw167_[int(14)] = self.f3
            nw167_[int(15)] = self.f3
            nw167_[int(16)] = self.f3
            nw167_[int(17)] = self.f3
            nw167_[int(18)] = False
            nw167_[int(19)] = self.f3
            nw167_[int(20)] = self.f3
            nw167_[int(21)] = self.f3
            nw167_[int(22)] = self.f3
            nw167_[int(23)] = not(self.f3)
            nw167_[int(24)] = self.f3
            nw167_[int(25)] = (d_1261_v2_)[default__.safeIndex(p1, len(d_1261_v2_))]
            nw167_[int(26)] = self.f3
            nw167_[int(27)] = self.f3
            nw167_[int(28)] = self.f3
            d_1262_v3_ = nw167_
            d_1263_v4_: T2
            nw168_ = C3()
            nw168_.ctor__((d_1260_v1_).cardinality, d_1262_v3_, self.f3, (self).f1, (self).f2)
            d_1263_v4_ = nw168_
            d_1264_v5_: _dafny.Map
            d_1264_v5_ = _dafny.Map({self.f3: (_dafny.MultiSet([p1, default__.fm0(globalState)])).cardinality})
            d_1265_v6_: _dafny.Map
            d_1265_v6_ = _dafny.Map({d_1263_v4_: len(_dafny.Map({len(d_1264_v5_): d_1263_v4_.f3}))})
            d_1266_v7_: _dafny.Set
            d_1266_v7_ = _dafny.Set({self.f3, self.f3, d_1263_v4_.f3})
            d_1267_v8_: _dafny.Array
            nw169_ = _dafny.Array(None, 28)
            nw169_[int(0)] = p0
            nw169_[int(1)] = (len(d_1259_v0_)) + (p0)
            nw169_[int(2)] = len((self).f2)
            nw169_[int(3)] = p1
            nw169_[int(4)] = d_1258_i0_
            nw169_[int(5)] = 911
            nw169_[int(6)] = p0
            nw169_[int(7)] = p1
            nw169_[int(8)] = default__.safeDivisionInt(p0, len((self).f2))
            nw169_[int(9)] = d_1258_i0_
            nw169_[int(10)] = (0) - (p0)
            nw169_[int(11)] = p0
            nw169_[int(12)] = d_1258_i0_
            nw169_[int(13)] = len((self).f2)
            nw169_[int(14)] = (d_1260_v1_).cardinality
            nw169_[int(15)] = default__.safeModuloInt(len((self).f2), len(_dafny.Map({self.f3: p0})))
            nw169_[int(16)] = (p1) - (250)
            nw169_[int(17)] = (self).fm12(len(d_1265_v6_), d_1258_i0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), d_1263_v4_.f3, globalState)
            nw169_[int(18)] = default__.safeDivisionInt((d_1263_v4_).f4, (d_1263_v4_).f4)
            nw169_[int(19)] = p0
            nw169_[int(20)] = p1
            nw169_[int(21)] = (d_1263_v4_).f4
            nw169_[int(22)] = p0
            nw169_[int(23)] = (d_1263_v4_).f4
            nw169_[int(24)] = ((d_1263_v4_).fm12(d_1258_i0_, (d_1263_v4_).f4, d_1258_i0_, d_1263_v4_.f3, globalState)) + (p1)
            nw169_[int(25)] = (d_1263_v4_).f4
            nw169_[int(26)] = len((d_1266_v7_) | (d_1266_v7_))
            nw169_[int(27)] = ((d_1260_v1_)[True] if (True) in (d_1260_v1_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1268_i1_ in range(default__.abs(819))])))
            d_1267_v8_ = nw169_
            index197_ = default__.safeIndex(870, (d_1267_v8_).length(0))
            (d_1267_v8_)[index197_] = p0
            index198_ = default__.safeIndex(870, (d_1267_v8_).length(0))
            (d_1267_v8_)[index198_] = ((d_1263_v4_).f4) - (p0)
            d_1269_v9_: C3
            nw170_ = C3()
            nw170_.ctor__(p1, (d_1263_v4_).f5, (p0) == ((d_1263_v4_).f4), (self).f1, (self).f2)
            d_1269_v9_ = nw170_
            d_1270_v10_: _dafny.MultiSet
            d_1270_v10_ = _dafny.MultiSet([d_1267_v8_])
            (d_1263_v4_).f3 = (d_1270_v10_).issubset(d_1270_v10_)
        d_1271_v11_: _dafny.Array
        nw171_ = _dafny.Array(int(0), 23)
        d_1271_v11_ = nw171_
        d_1272_v12_: _dafny.Seq
        d_1272_v12_ = _dafny.SeqWithoutIsStrInference([d_1271_v11_])
        d_1272_v12_ = (d_1272_v12_) + (_dafny.SeqWithoutIsStrInference([d_1271_v11_]))
        (self).f3 = not(False)
        d_1273_i2_: int
        d_1273_i2_ = 0
        with _dafny.label("8"):
            while not(self.f3):
                with _dafny.c_label("8"):
                    if (d_1273_i2_) >= (100):
                        raise _dafny.Break("8")
                    d_1273_i2_ = (d_1273_i2_) + (1)
                    d_1274_v13_: _dafny.Array
                    nw172_ = _dafny.Array(_dafny.Map({}), 19)
                    d_1274_v13_ = nw172_
                    d_1275_v14_: _dafny.Array
                    nw173_ = _dafny.Array(None, 15)
                    nw173_[int(0)] = self.f3
                    nw173_[int(1)] = True
                    nw173_[int(2)] = self.f3
                    nw173_[int(3)] = self.f3
                    nw173_[int(4)] = self.f3
                    nw173_[int(5)] = self.f3
                    nw173_[int(6)] = self.f3
                    nw173_[int(7)] = self.f3
                    nw173_[int(8)] = self.f3
                    nw173_[int(9)] = self.f3
                    nw173_[int(10)] = False
                    nw173_[int(11)] = self.f3
                    nw173_[int(12)] = self.f3
                    nw173_[int(13)] = self.f3
                    nw173_[int(14)] = self.f3
                    d_1275_v14_ = nw173_
                    d_1276_v15_: str
                    d_1276_v15_ = _dafny.CodePoint('o')
                    index199_ = default__.safeIndex(195, (d_1274_v13_).length(0))
                    (d_1274_v13_)[index199_] = (_dafny.Map({d_1275_v14_: d_1276_v15_})).set(d_1275_v14_, _dafny.CodePoint('a'))
                    d_1277_v16_: _dafny.Map
                    d_1277_v16_ = _dafny.Map({d_1275_v14_: d_1276_v15_})
                    index200_ = default__.safeIndex(195, (d_1274_v13_).length(0))
                    (d_1274_v13_)[index200_] = d_1277_v16_
                    (self).f3 = self.f3
                    d_1278_v17_: _dafny.Seq
                    d_1278_v17_ = _dafny.SeqWithoutIsStrInference([self.f3, not(self.f3), self.f3])
                    d_1279_v18_: _dafny.Seq
                    d_1279_v18_ = d_1278_v17_
                    source23_ = d_1279_v18_
                    d_1280___mcc_h0_ = source23_
                    d_1281_cf51_ = d_1280___mcc_h0_
                    d_1282_v19_: _dafny.Map
                    d_1282_v19_ = _dafny.Map({p1: self.f3})
                    d_1283_v20_: _dafny.MultiSet
                    d_1283_v20_ = _dafny.MultiSet([p1])
                    d_1284_v21_: _dafny.Map
                    d_1284_v21_ = _dafny.Map({d_1283_v20_: p1})
                    d_1282_v19_ = (d_1282_v19_).set(default__.safeModuloInt(p0, len(d_1284_v21_)), self.f3)
                    d_1281_cf51_ = (_dafny.SeqWithoutIsStrInference([self.f3, self.f3])) + ((_dafny.SeqWithoutIsStrInference([self.f3, self.f3])) + (d_1278_v17_))
                    d_1285_v22_: _dafny.MultiSet
                    d_1285_v22_ = _dafny.MultiSet([self.f3, False])
                    d_1286_v23_: _dafny.Map
                    d_1286_v23_ = _dafny.Map({d_1271_v11_: d_1271_v11_})
                    index201_ = default__.safeIndex(658, (d_1271_v11_).length(0))
                    (d_1271_v11_)[index201_] = p0
                    d_1287_v24_: _dafny.Seq
                    d_1287_v24_ = _dafny.SeqWithoutIsStrInference([d_1286_v23_, d_1286_v23_])
                    index202_ = default__.safeIndex(658, (d_1271_v11_).length(0))
                    rhs196_ = d_1271_v11_
                    rhs197_ = d_1285_v22_
                    rhs198_ = (d_1287_v24_)[default__.safeIndex(p1, len(d_1287_v24_))]
                    rhs199_ = (p0) + (p1)
                    rhs200_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxvovuu")))
                    lhs108_ = d_1271_v11_
                    lhs109_ = default__.safeIndex(658, (d_1271_v11_).length(0))
                    d_1271_v11_ = rhs196_
                    d_1285_v22_ = rhs197_
                    d_1286_v23_ = rhs198_
                    r1 = rhs199_
                    lhs108_[lhs109_] = rhs200_
                    d_1283_v20_ = (((_dafny.MultiSet([len((self).f2)])).set(len((self).f2), default__.abs(p0))) - (d_1283_v20_)).intersection(d_1283_v20_)
                    index203_ = default__.safeIndex(223, (d_1275_v14_).length(0))
                    (d_1275_v14_)[index203_] = self.f3
                    d_1288_v25_: _dafny.MultiSet
                    d_1288_v25_ = _dafny.MultiSet([self.f3])
                    d_1289_v26_: _dafny.Seq
                    d_1289_v26_ = _dafny.SeqWithoutIsStrInference([(d_1288_v25_).cardinality])
                    d_1290_v27_: _dafny.Seq
                    d_1290_v27_ = _dafny.SeqWithoutIsStrInference([d_1289_v26_, d_1289_v26_, d_1289_v26_])
                    d_1291_v28_: _dafny.Map
                    d_1291_v28_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_1292_i4_ in range(default__.abs(748))]) for d_1293_i3_ in range(default__.abs(394))])) + (d_1290_v27_): not(self.f3)})
                    index204_ = default__.safeIndex(223, (d_1275_v14_).length(0))
                    (d_1275_v14_)[index204_] = ((d_1291_v28_)[(d_1290_v27_) + (d_1290_v27_)] if ((d_1290_v27_) + (d_1290_v27_)) in (d_1291_v28_) else self.f3)
                    pass
            pass
        d_1294_v29_: str
        d_1294_v29_ = _dafny.CodePoint('w')
        d_1295_v30_: D2
        d_1295_v30_ = D2_DC6(d_1294_v29_, self.f3, self.f3, p1)
        source24_ = (d_1295_v30_ if self.f3 else d_1295_v30_)
        if source24_.is_DC6:
            d_1296___mcc_h1_ = source24_.cf9
            d_1297___mcc_h2_ = source24_.cf10
            d_1298___mcc_h3_ = source24_.cf11
            d_1299___mcc_h4_ = source24_.cf12
            d_1300_cf12_ = d_1299___mcc_h4_
            d_1301_cf11_ = d_1298___mcc_h3_
            d_1302_cf10_ = d_1297___mcc_h2_
            d_1303_cf9_ = d_1296___mcc_h1_
            d_1304_v31_: _dafny.Seq
            d_1304_v31_ = _dafny.SeqWithoutIsStrInference([(self).f2, (self).f2, (self).f2])
            d_1304_v31_ = d_1304_v31_
            d_1305_v32_: _dafny.MultiSet
            d_1305_v32_ = _dafny.MultiSet([d_1302_cf10_, d_1302_cf10_])
            r2 = ((d_1305_v32_)[d_1301_cf11_] if (d_1301_cf11_) in (d_1305_v32_) else 250)
            d_1306_v33_: _dafny.Set
            d_1306_v33_ = _dafny.Set({self.f3, d_1301_cf11_})
            d_1307_v34_: _dafny.Map
            d_1307_v34_ = _dafny.Map({d_1301_cf11_: p0})
            d_1300_cf12_ = (len(d_1306_v33_)) - (((d_1307_v34_)[d_1301_cf11_] if (d_1301_cf11_) in (d_1307_v34_) else p1))
            d_1308_v35_: D5
            d_1308_v35_ = D5_DC14((d_1272_v12_)[default__.safeIndex(176, len(d_1272_v12_))])
            source25_ = d_1308_v35_
            if source25_.is_DC15:
                d_1309___mcc_h7_ = source25_.cf24
                d_1310___mcc_h8_ = source25_.cf25
                d_1311___mcc_h9_ = source25_.cf26
                d_1312_cf26_ = d_1311___mcc_h9_
                d_1313_cf25_ = d_1310___mcc_h8_
                d_1314_cf24_ = d_1309___mcc_h7_
                d_1315_v36_: _dafny.Array
                nw174_ = _dafny.Array(D10.default()(), 6)
                d_1315_v36_ = nw174_
                d_1316_v37_: _dafny.Seq
                d_1316_v37_ = _dafny.SeqWithoutIsStrInference([d_1315_v36_])
                d_1317_v38_: _dafny.MultiSet
                d_1317_v38_ = _dafny.MultiSet([(d_1316_v37_)[default__.safeIndex(len(d_1307_v34_), len(d_1316_v37_))], d_1315_v36_, d_1315_v36_, d_1315_v36_])
                d_1317_v38_ = d_1317_v38_
                d_1301_cf11_ = (d_1302_cf10_) == (self.f3)
                r2 = ((0) - ((0) - (default__.safeModuloInt(d_1300_cf12_, 292))) if d_1301_cf11_ else default__.safeModuloInt(((d_1305_v32_)[d_1302_cf10_] if (d_1302_cf10_) in (d_1305_v32_) else d_1314_cf24_), p0))
                index205_ = default__.safeIndex(471, (d_1271_v11_).length(0))
                (d_1271_v11_)[index205_] = p1
                d_1318_v39_: _dafny.Seq
                d_1318_v39_ = _dafny.SeqWithoutIsStrInference([d_1301_cf11_])
                d_1319_v40_: _dafny.Map
                d_1319_v40_ = _dafny.Map({(d_1312_cf26_ if d_1302_cf10_ else d_1312_cf26_): (self.f3) in (d_1318_v39_)})
                d_1320_v41_: _dafny.Map
                d_1320_v41_ = _dafny.Map({d_1301_cf11_: self.f3})
                d_1321_v42_: D8
                d_1321_v42_ = D8_DC24(d_1294_v29_, d_1320_v41_)
                d_1322_v43_: _dafny.Seq
                d_1322_v43_ = _dafny.SeqWithoutIsStrInference([d_1321_v42_])
                d_1323_v44_: _dafny.Seq
                d_1323_v44_ = _dafny.SeqWithoutIsStrInference([p0])
                index206_ = default__.safeIndex(471, (d_1271_v11_).length(0))
                rhs201_ = d_1300_cf12_
                rhs202_ = d_1300_cf12_
                rhs203_ = d_1294_v29_
                rhs204_ = d_1319_v40_
                rhs205_ = (len((d_1322_v43_) + (default__.fm47(len(d_1323_v44_), 480, len(d_1318_v39_), True, globalState)))) != (d_1314_cf24_)
                lhs110_ = d_1271_v11_
                lhs111_ = default__.safeIndex(471, (d_1271_v11_).length(0))
                lhs110_[lhs111_] = rhs201_
                d_1300_cf12_ = rhs202_
                d_1303_cf9_ = rhs203_
                d_1319_v40_ = rhs204_
                d_1301_cf11_ = rhs205_
            elif True:
                d_1324___mcc_h10_ = source25_.cf23
                d_1325_cf23_ = d_1324___mcc_h10_
                d_1326_v45_: _dafny.Array
                nw175_ = _dafny.Array(None, 6)
                nw175_[int(0)] = self.f3
                nw175_[int(1)] = self.f3
                nw175_[int(2)] = d_1301_cf11_
                nw175_[int(3)] = self.f3
                nw175_[int(4)] = self.f3
                nw175_[int(5)] = d_1302_cf10_
                d_1326_v45_ = nw175_
                d_1327_v46_: C3
                nw176_ = C3()
                nw176_.ctor__((0) - (default__.safeModuloInt(d_1300_cf12_, p1)), d_1326_v45_, self.f3, (self).f1, (self).f2)
                d_1327_v46_ = nw176_
                d_1328_v47_: _dafny.Seq
                d_1328_v47_ = _dafny.SeqWithoutIsStrInference([len((self).f2), d_1300_cf12_])
                d_1307_v34_ = (d_1307_v34_).set((d_1328_v47_) != (_dafny.SeqWithoutIsStrInference([d_1300_cf12_])), d_1300_cf12_)
                d_1302_cf10_ = self.f3
                d_1302_cf10_ = not (self.f3) or (d_1301_cf11_)
        elif source24_.is_DC5:
            d_1329___mcc_h5_ = source24_.cf8
            d_1330_cf8_ = d_1329___mcc_h5_
            d_1331_v48_: _dafny.Array
            def lambda49_(d_1332_i5_):
                return self.f3

            init26_ = lambda49_
            nw177_ = _dafny.Array(None, 18)
            for i0_26_ in range(nw177_.length(0)):
                nw177_[i0_26_] = init26_(i0_26_)
            d_1331_v48_ = nw177_
            d_1333_v49_: _dafny.Map
            d_1333_v49_ = _dafny.Map({self.f3: d_1331_v48_})
            r2 = len(d_1333_v49_)
            d_1334_v50_: _dafny.Set
            d_1334_v50_ = _dafny.Set({self.f3})
            r2 = (len((d_1334_v50_) - (d_1334_v50_))) + (((0) - (-932)) - (default__.fm0(globalState)))
            d_1271_v11_ = d_1271_v11_
            d_1335_v51_: _dafny.Seq
            d_1335_v51_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1336_v52_: _dafny.Seq
            d_1336_v52_ = _dafny.SeqWithoutIsStrInference([p1, len(d_1335_v51_)])
            r1 = (-71) + ((len(d_1336_v52_)) * (len((self).f2)))
        elif True:
            d_1337___mcc_h6_ = source24_.cf13
            d_1338_cf13_ = d_1337___mcc_h6_
            r2 = p1
            (self).f3 = not (True) or (self.f3)
            if (p1) < ((0) - ((p1) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppfqtk")))))):
                d_1339_v53_: _dafny.Seq
                d_1339_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjsmxnhxl"))
                d_1339_v53_ = (self).f2
                index207_ = default__.safeIndex(741, (d_1271_v11_).length(0))
                (d_1271_v11_)[index207_] = (0) - (p1)
                index208_ = default__.safeIndex(741, (d_1271_v11_).length(0))
                (d_1271_v11_)[index208_] = p0
                (self).m9((self.f3) == (self.f3), self.f3, self.f3, globalState)
                d_1340_v54_: _dafny.Seq
                d_1340_v54_ = _dafny.SeqWithoutIsStrInference([len(d_1339_v53_), p0])
                d_1341_v55_: _dafny.Set
                d_1341_v55_ = _dafny.Set({161})
                d_1342_v56_: _dafny.Map
                d_1342_v56_ = _dafny.Map({(_dafny.Set({_dafny.Set({p0, (d_1340_v54_)[default__.safeIndex(p1, len(d_1340_v54_))]}), d_1341_v55_, d_1341_v55_})) - (_dafny.Set({d_1341_v55_, d_1341_v55_})): (0) - (default__.fm0(globalState))})
                d_1343_v58_: _dafny.Set
                def iife160_():
                    coll44_ = _dafny.Set()
                    compr_44_: int
                    for compr_44_ in (d_1340_v54_).Elements:
                        d_1344_v57_: int = compr_44_
                        if (d_1344_v57_) in (d_1340_v54_):
                            coll44_ = coll44_.union(_dafny.Set([(d_1344_v57_) * (815)]))
                    return _dafny.Set(coll44_)
                d_1343_v58_ = _dafny.Set({iife160_()
                , d_1341_v55_})
                d_1342_v56_ = (d_1342_v56_).set(d_1343_v58_, p0)
                index209_ = default__.safeIndex(741, (d_1271_v11_).length(0))
                (d_1271_v11_)[index209_] = p1
            elif True:
                d_1345_v59_: _dafny.Map
                d_1345_v59_ = _dafny.Map({p0: False})
                (self).m9(False, self.f3, (p0) not in (d_1345_v59_), globalState)
                d_1346_v60_: _dafny.Seq
                d_1346_v60_ = _dafny.SeqWithoutIsStrInference([p1, p0, (0) - (p1), (p0) - (p0)])
                d_1346_v60_ = (_dafny.SeqWithoutIsStrInference([706])) + (_dafny.SeqWithoutIsStrInference([p1 for d_1347_i6_ in range(default__.abs(-234))]))
                d_1348_v61_: _dafny.Array
                nw178_ = _dafny.Array(False, 1)
                d_1348_v61_ = nw178_
                index210_ = default__.safeIndex(386, (d_1348_v61_).length(0))
                (d_1348_v61_)[index210_] = self.f3
                d_1349_v62_: _dafny.MultiSet
                d_1349_v62_ = _dafny.MultiSet([self.f3, self.f3, self.f3])
                index211_ = default__.safeIndex(386, (d_1348_v61_).length(0))
                rhs206_ = (p0) * (p1)
                rhs207_ = ((0) - ((d_1349_v62_).cardinality)) <= (default__.safeModuloInt(p1, p0))
                rhs208_ = not(self.f3)
                lhs112_ = self
                lhs113_ = d_1348_v61_
                lhs114_ = default__.safeIndex(386, (d_1348_v61_).length(0))
                r2 = rhs206_
                lhs112_.f3 = rhs207_
                lhs113_[lhs114_] = rhs208_
                d_1350_v63_: _dafny.MultiSet
                d_1350_v63_ = _dafny.MultiSet([706, p0, len(_dafny.Set({p0}))])
                d_1351_v64_: D11
                d_1351_v64_ = D11_DC31((d_1350_v63_).cardinality)
                d_1352_v65_: _dafny.Map
                d_1352_v65_ = _dafny.Map({self.f3: False})
                d_1353_v66_: _dafny.Map
                d_1353_v66_ = _dafny.Map({d_1351_v64_: D2_DC6(d_1294_v29_, False, self.f3, len(d_1352_v65_))})
                r2 = len(d_1353_v66_)
                d_1354_v67_: _dafny.Set
                d_1354_v67_ = _dafny.Set({self.f3})
                index212_ = default__.safeIndex(280, (d_1271_v11_).length(0))
                (d_1271_v11_)[index212_] = (0) - ((len(d_1354_v67_)) - (len((self).f2)))
                index213_ = default__.safeIndex(280, (d_1271_v11_).length(0))
                (d_1271_v11_)[index213_] = 688
            d_1294_v29_ = ((self).f2)[default__.safeIndex(p0, len((self).f2))]
        d_1355_v68_: _dafny.Map
        d_1355_v68_ = _dafny.Map({p0: self.f3})
        d_1356_v69_: D6
        d_1356_v69_ = D6_DC16(d_1355_v68_)
        pat_let_tv37_ = p0
        pat_let_tv38_ = p0
        pat_let_tv39_ = d_1355_v68_
        def lambda50_(source26_):
            if source26_.is_DC17:
                d_1357___mcc_h11_ = source26_.cf28
                d_1358_cf28_ = d_1357___mcc_h11_
                return 706
            elif source26_.is_DC18:
                d_1359___mcc_h12_ = source26_.cf29
                d_1360___mcc_h13_ = source26_.cf30
                d_1361___mcc_h14_ = source26_.cf31
                d_1362___mcc_h15_ = source26_.cf32
                d_1363___mcc_h16_ = source26_.cf33
                d_1364_cf33_ = d_1363___mcc_h16_
                d_1365_cf32_ = d_1362___mcc_h15_
                d_1366_cf31_ = d_1361___mcc_h14_
                d_1367_cf30_ = d_1360___mcc_h13_
                d_1368_cf29_ = d_1359___mcc_h12_
                return d_1367_cf30_
            elif source26_.is_DC19:
                d_1369___mcc_h17_ = source26_.cf34
                d_1370_cf34_ = d_1369___mcc_h17_
                return pat_let_tv37_
            elif source26_.is_DC16:
                d_1371___mcc_h18_ = source26_.cf27
                d_1372_cf27_ = d_1371___mcc_h18_
                return -982
            elif True:
                d_1373___mcc_h19_ = source26_.cf35
                d_1374_cf35_ = d_1373___mcc_h19_
                return pat_let_tv38_

        def iife161_(_pat_let58_0):
            def iife162_(d_1375_dt__update__tmp_h0_):
                def iife163_(_pat_let59_0):
                    def iife164_(d_1376_dt__update_hcf27_h0_):
                        return D6_DC16(d_1376_dt__update_hcf27_h0_)
                    return iife164_(_pat_let59_0)
                return iife163_(pat_let_tv39_)
            return iife162_(_pat_let58_0)
        r2 = lambda50_(iife161_(d_1356_v69_))
        d_1377_v70_: _dafny.Map
        d_1377_v70_ = _dafny.Map({self.f3: self.f3})
        r0 = (d_1377_v70_) | ((d_1377_v70_) | (d_1377_v70_))
        d_1378_v71_: _dafny.MultiSet
        d_1378_v71_ = _dafny.MultiSet([False])
        r1 = default__.safeDivisionInt((p1) + ((d_1378_v71_).cardinality), p0)
        r2 = default__.safeModuloInt((p0) * (p1), default__.fm0(globalState))
        return r0, r1, r2

    def m9(self, p0, p1, p2, globalState):
        d_1379_v0_: str
        d_1379_v0_ = _dafny.CodePoint('t')
        d_1380_v1_: _dafny.MultiSet
        d_1380_v1_ = _dafny.MultiSet([p2, self.f3, False])
        d_1379_v0_ = (d_1379_v0_ if (d_1380_v1_).isdisjoint(d_1380_v1_) else d_1379_v0_)
        d_1381_v2_: int
        d_1381_v2_ = 941
        d_1382_v3_: _dafny.Seq
        d_1382_v3_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1383_v4_: _dafny.Seq
        d_1383_v4_ = _dafny.SeqWithoutIsStrInference([d_1381_v2_])
        d_1384_v5_: T1
        nw179_ = C2()
        nw179_.ctor__(len((self).f2), p1, (self).f1, (self).f2)
        d_1384_v5_ = nw179_
        d_1385_v6_: _dafny.Seq
        d_1385_v6_ = _dafny.SeqWithoutIsStrInference([d_1384_v5_, d_1384_v5_])
        d_1386_v7_: _dafny.Array
        nw180_ = _dafny.Array(None, 19)
        nw180_[int(0)] = p2
        nw180_[int(1)] = p2
        nw180_[int(2)] = (_dafny.MultiSet([d_1381_v2_])).ispropersubset(_dafny.MultiSet([d_1381_v2_]))
        nw180_[int(3)] = True
        nw180_[int(4)] = (_dafny.SeqWithoutIsStrInference([True])) == (d_1382_v3_)
        nw180_[int(5)] = (self.f3 if not(p1) else False)
        nw180_[int(6)] = p1
        nw180_[int(7)] = not(p2)
        nw180_[int(8)] = (d_1382_v3_)[default__.safeIndex(len(_dafny.Set({default__.fm0(globalState)})), len(d_1382_v3_))]
        nw180_[int(9)] = p1
        nw180_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_1379_v0_ for d_1387_i0_ in range(default__.abs(-431))])) <= ((self).f2)
        nw180_[int(11)] = (d_1381_v2_) == ((d_1383_v4_)[default__.safeIndex(d_1381_v2_, len(d_1383_v4_))])
        nw180_[int(12)] = p2
        nw180_[int(13)] = (d_1384_v5_) in (d_1385_v6_)
        nw180_[int(14)] = ((d_1382_v3_)[default__.safeIndex(d_1381_v2_, len(d_1382_v3_))] if p2 else p0)
        nw180_[int(15)] = p0
        nw180_[int(16)] = self.f3
        nw180_[int(17)] = ((self).f2) <= ((self).f2)
        nw180_[int(18)] = d_1384_v5_.f3
        d_1386_v7_ = nw180_
        d_1386_v7_ = d_1386_v7_
        d_1388_v8_: _dafny.Array
        def lambda51_(d_1389_i1_):
            return D9_DC26(892)

        init27_ = lambda51_
        nw181_ = _dafny.Array(None, 11)
        for i0_27_ in range(nw181_.length(0)):
            nw181_[i0_27_] = init27_(i0_27_)
        d_1388_v8_ = nw181_
        d_1390_v9_: _dafny.Map
        d_1390_v9_ = _dafny.Map({d_1388_v8_: d_1386_v7_})
        d_1391_v10_: _dafny.Seq
        d_1391_v10_ = _dafny.SeqWithoutIsStrInference([d_1386_v7_])
        d_1392_v11_: _dafny.Map
        d_1392_v11_ = _dafny.Map({d_1381_v2_: d_1386_v7_})
        d_1393_v12_: _dafny.Map
        d_1393_v12_ = _dafny.Map({p2: d_1381_v2_})
        d_1394_v13_: _dafny.Seq
        d_1394_v13_ = _dafny.SeqWithoutIsStrInference([d_1390_v9_])
        d_1395_v14_: _dafny.Array
        nw182_ = _dafny.Array(None, 25)
        nw182_[int(0)] = _dafny.Map({d_1388_v8_: d_1386_v7_})
        nw182_[int(1)] = d_1390_v9_
        nw182_[int(2)] = (d_1390_v9_).set(d_1388_v8_, d_1386_v7_)
        nw182_[int(3)] = d_1390_v9_
        nw182_[int(4)] = _dafny.Map({d_1388_v8_: d_1386_v7_})
        nw182_[int(5)] = (d_1390_v9_) | (_dafny.Map({d_1388_v8_: d_1386_v7_}))
        nw182_[int(6)] = d_1390_v9_
        nw182_[int(7)] = d_1390_v9_
        nw182_[int(8)] = d_1390_v9_
        nw182_[int(9)] = d_1390_v9_
        nw182_[int(10)] = d_1390_v9_
        nw182_[int(11)] = d_1390_v9_
        nw182_[int(12)] = d_1390_v9_
        nw182_[int(13)] = (d_1390_v9_) | (d_1390_v9_)
        nw182_[int(14)] = _dafny.Map({d_1388_v8_: (d_1391_v10_)[default__.safeIndex((d_1380_v1_).cardinality, len(d_1391_v10_))]})
        nw182_[int(15)] = _dafny.Map({d_1388_v8_: d_1386_v7_})
        nw182_[int(16)] = d_1390_v9_
        nw182_[int(17)] = d_1390_v9_
        nw182_[int(18)] = d_1390_v9_
        nw182_[int(19)] = (d_1390_v9_) | (_dafny.Map({d_1388_v8_: d_1386_v7_}))
        nw182_[int(20)] = d_1390_v9_
        nw182_[int(21)] = _dafny.Map({d_1388_v8_: ((d_1392_v11_)[(0) - (len(d_1393_v12_))] if ((0) - (len(d_1393_v12_))) in (d_1392_v11_) else d_1386_v7_)})
        nw182_[int(22)] = ((d_1394_v13_)[default__.safeIndex(d_1381_v2_, len(d_1394_v13_))]) | ((_dafny.Map({d_1388_v8_: d_1386_v7_})).set(d_1388_v8_, d_1386_v7_))
        nw182_[int(23)] = (d_1390_v9_) | (d_1390_v9_)
        nw182_[int(24)] = d_1390_v9_
        d_1395_v14_ = nw182_
        index214_ = default__.safeIndex(890, (d_1395_v14_).length(0))
        (d_1395_v14_)[index214_] = d_1390_v9_
        d_1396_v15_: _dafny.MultiSet
        d_1396_v15_ = _dafny.MultiSet([d_1381_v2_, d_1381_v2_])
        d_1397_v16_: _dafny.Set
        d_1397_v16_ = _dafny.Set({p0})
        d_1398_v17_: D3
        d_1398_v17_ = D3_DC8(d_1397_v16_)
        pat_let_tv40_ = d_1384_v5_
        pat_let_tv41_ = d_1384_v5_
        pat_let_tv42_ = d_1384_v5_
        pat_let_tv43_ = d_1384_v5_
        index215_ = default__.safeIndex(890, (d_1395_v14_).length(0))
        def lambda52_(source27_):
            if source27_.is_DC9:
                d_1399___mcc_h0_ = source27_.cf15
                d_1400___mcc_h1_ = source27_.cf16
                d_1401_cf16_ = d_1400___mcc_h1_
                d_1402_cf15_ = d_1399___mcc_h0_
                return not(pat_let_tv40_.f3)
            elif source27_.is_DC10:
                d_1403___mcc_h2_ = source27_.cf17
                d_1404___mcc_h3_ = source27_.cf18
                d_1405_cf18_ = d_1404___mcc_h3_
                d_1406_cf17_ = d_1403___mcc_h2_
                return pat_let_tv41_.f3
            elif source27_.is_DC8:
                d_1407___mcc_h4_ = source27_.cf14
                d_1408_cf14_ = d_1407___mcc_h4_
                return pat_let_tv42_.f3
            elif True:
                d_1409___mcc_h5_ = source27_.cf19
                d_1410_cf19_ = d_1409___mcc_h5_
                return pat_let_tv43_.f3

        rhs209_ = (d_1383_v4_) < ((d_1383_v4_).set(default__.safeIndex(d_1381_v2_, len(d_1383_v4_)), d_1381_v2_))
        rhs210_ = (0) - (((d_1396_v15_)[(d_1381_v2_) - (d_1381_v2_)] if ((d_1381_v2_) - (d_1381_v2_)) in (d_1396_v15_) else (d_1381_v2_) + (d_1381_v2_)))
        rhs211_ = p0
        rhs212_ = not(lambda52_(d_1398_v17_))
        rhs213_ = d_1390_v9_
        lhs115_ = d_1384_v5_
        lhs116_ = self
        lhs117_ = self
        lhs118_ = d_1395_v14_
        lhs119_ = default__.safeIndex(890, (d_1395_v14_).length(0))
        lhs115_.f3 = rhs209_
        d_1381_v2_ = rhs210_
        lhs116_.f3 = rhs211_
        lhs117_.f3 = rhs212_
        lhs118_[lhs119_] = rhs213_
        d_1411_i2_: int
        d_1411_i2_ = 0
        with _dafny.label("9"):
            while ((self).f2) == ((self).f2):
                with _dafny.c_label("9"):
                    if (d_1411_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_1411_i2_ = (d_1411_i2_) + (1)
                    d_1412_v18_: D6
                    d_1412_v18_ = D6_DC17(d_1384_v5_.f3)
                    d_1413_v19_: D6
                    d_1413_v19_ = D6_DC20(d_1412_v18_)
                    d_1414_v20_: D6
                    d_1414_v20_ = D6_DC20(d_1412_v18_)
                    d_1415_v21_: D6
                    d_1415_v21_ = D6_DC20(d_1413_v19_)
                    source28_ = d_1415_v21_
                    if source28_.is_DC17:
                        d_1416___mcc_h6_ = source28_.cf28
                        d_1417_cf28_ = d_1416___mcc_h6_
                        d_1418_v22_: _dafny.Map
                        d_1418_v22_ = _dafny.Map({d_1381_v2_: d_1384_v5_.f3})
                        d_1419_v23_: _dafny.Map
                        d_1419_v23_ = _dafny.Map({d_1381_v2_: d_1381_v2_})
                        d_1420_v24_: _dafny.Map
                        d_1420_v24_ = _dafny.Map({((d_1419_v23_)[d_1381_v2_] if (d_1381_v2_) in (d_1419_v23_) else ((d_1393_v12_)[p1] if (p1) in (d_1393_v12_) else len((d_1384_v5_).f2))): d_1393_v12_})
                        (d_1384_v5_).f3 = (len(d_1418_v22_)) in (d_1420_v24_)
                        d_1421_v25_: C2
                        nw183_ = C2()
                        nw183_.ctor__((d_1381_v2_) - (-24), d_1384_v5_.f3, (d_1384_v5_).f1, (d_1384_v5_).f2)
                        d_1421_v25_ = nw183_
                        d_1422_v26_: _dafny.Set
                        d_1422_v26_ = _dafny.Set({d_1381_v2_, d_1381_v2_})
                        d_1423_v27_: D6
                        d_1423_v27_ = D6_DC18(len(d_1422_v26_), (d_1421_v25_).f13, True, d_1379_v0_, (d_1421_v25_).f13)
                        d_1424_v28_: _dafny.Seq
                        d_1424_v28_ = _dafny.SeqWithoutIsStrInference([d_1423_v27_])
                        d_1425_v29_: _dafny.Seq
                        d_1425_v29_ = _dafny.SeqWithoutIsStrInference([(d_1424_v28_) + (d_1424_v28_), d_1424_v28_, _dafny.SeqWithoutIsStrInference([d_1423_v27_, D6_DC18(282, (d_1421_v25_).f13, d_1417_cf28_, d_1379_v0_, (d_1421_v25_).f13)]), (default__.fm48(d_1384_v5_.f3, d_1381_v2_, globalState)) + (_dafny.SeqWithoutIsStrInference([d_1423_v27_ for d_1426_i3_ in range(default__.abs(815))]))])
                        d_1425_v29_ = (d_1425_v29_) + (_dafny.SeqWithoutIsStrInference([d_1424_v28_]))
                        d_1427_v30_: _dafny.Map
                        d_1427_v30_ = _dafny.Map({True: (_dafny.SeqWithoutIsStrInference([932, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1428_i5_ in range(default__.abs(72))]))])).set(default__.safeIndex(d_1381_v2_, len(_dafny.SeqWithoutIsStrInference([932, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1429_i5_ in range(default__.abs(72))]))]))), (d_1421_v25_).f13)})
                        d_1430_v31_: _dafny.Map
                        d_1430_v31_ = _dafny.Map({p1: d_1384_v5_.f3})
                        d_1431_v32_: D9
                        d_1431_v32_ = D9_DC27(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ina"))))
                        d_1432_v33_: _dafny.Map
                        d_1432_v33_ = _dafny.Map({d_1431_v32_: not(d_1384_v5_.f3)})
                        d_1433_v34_: _dafny.Array
                        nw184_ = _dafny.Array(None, 25)
                        nw184_[int(0)] = d_1383_v4_
                        nw184_[int(1)] = d_1383_v4_
                        nw184_[int(2)] = d_1383_v4_
                        nw184_[int(3)] = (d_1383_v4_) + (_dafny.SeqWithoutIsStrInference([(0) - ((d_1383_v4_)[default__.safeIndex((d_1421_v25_).f13, len(d_1383_v4_))]) for d_1434_i4_ in range(default__.abs(391))]))
                        nw184_[int(4)] = (d_1383_v4_) + (((d_1427_v30_)[d_1384_v5_.f3] if (d_1384_v5_.f3) in (d_1427_v30_) else d_1383_v4_))
                        nw184_[int(5)] = d_1383_v4_
                        nw184_[int(6)] = d_1383_v4_
                        nw184_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1381_v2_, d_1381_v2_, ((d_1380_v1_)[d_1384_v5_.f3] if (d_1384_v5_.f3) in (d_1380_v1_) else len(d_1430_v31_)), (d_1421_v25_).f13, (d_1421_v25_).f13])
                        nw184_[int(8)] = d_1383_v4_
                        nw184_[int(9)] = d_1383_v4_
                        nw184_[int(10)] = _dafny.SeqWithoutIsStrInference([((d_1419_v23_)[len((d_1384_v5_).f2)] if (len((d_1384_v5_).f2)) in (d_1419_v23_) else (d_1421_v25_).f13) for d_1435_i6_ in range(default__.abs(423))])
                        nw184_[int(11)] = d_1383_v4_
                        nw184_[int(12)] = _dafny.SeqWithoutIsStrInference([(d_1421_v25_).f13 for d_1436_i7_ in range(default__.abs(672))])
                        nw184_[int(13)] = d_1383_v4_
                        nw184_[int(14)] = (_dafny.SeqWithoutIsStrInference([(d_1384_v5_).fm12(d_1381_v2_, (d_1421_v25_).f13, len((self).f2), d_1384_v5_.f3, globalState)]) if d_1417_cf28_ else _dafny.SeqWithoutIsStrInference([(d_1421_v25_).f13]))
                        nw184_[int(15)] = d_1383_v4_
                        nw184_[int(16)] = d_1383_v4_
                        nw184_[int(17)] = d_1383_v4_
                        nw184_[int(18)] = d_1383_v4_
                        nw184_[int(19)] = d_1383_v4_
                        nw184_[int(20)] = d_1383_v4_
                        nw184_[int(21)] = d_1383_v4_
                        nw184_[int(22)] = d_1383_v4_
                        nw184_[int(23)] = (d_1383_v4_) + ((_dafny.SeqWithoutIsStrInference([len(d_1397_v16_) for d_1437_i8_ in range(default__.abs(819))])).set(default__.safeIndex(-923, len(_dafny.SeqWithoutIsStrInference([len(d_1397_v16_) for d_1438_i8_ in range(default__.abs(819))]))), len(d_1432_v33_)))
                        nw184_[int(24)] = (d_1383_v4_) + (d_1383_v4_)
                        d_1433_v34_ = nw184_
                        index216_ = default__.safeIndex(467, (d_1433_v34_).length(0))
                        (d_1433_v34_)[index216_] = d_1383_v4_
                        index217_ = default__.safeIndex(467, (d_1433_v34_).length(0))
                        (d_1433_v34_)[index217_] = default__.fm31((254) + (default__.fm0(globalState)), p1, default__.safeModuloInt((d_1421_v25_).f13, (d_1421_v25_).f13), d_1379_v0_, globalState)
                    elif source28_.is_DC18:
                        d_1439___mcc_h7_ = source28_.cf29
                        d_1440___mcc_h8_ = source28_.cf30
                        d_1441___mcc_h9_ = source28_.cf31
                        d_1442___mcc_h10_ = source28_.cf32
                        d_1443___mcc_h11_ = source28_.cf33
                        d_1444_cf33_ = d_1443___mcc_h11_
                        d_1445_cf32_ = d_1442___mcc_h10_
                        d_1446_cf31_ = d_1441___mcc_h9_
                        d_1447_cf30_ = d_1440___mcc_h8_
                        d_1448_cf29_ = d_1439___mcc_h7_
                        (d_1384_v5_).f3 = d_1384_v5_.f3
                        d_1381_v2_ = (d_1444_cf33_ if p1 else d_1448_cf29_)
                        d_1449_v35_: _dafny.Map
                        d_1449_v35_ = _dafny.Map({_dafny.MultiSet((d_1383_v4_).set(default__.safeIndex(d_1381_v2_, len(d_1383_v4_)), d_1448_cf29_)): d_1384_v5_.f3})
                        d_1450_v36_: _dafny.Map
                        d_1450_v36_ = _dafny.Map({not(d_1384_v5_.f3): d_1449_v35_})
                        d_1450_v36_ = (d_1450_v36_).set(d_1384_v5_.f3, (d_1449_v35_).set(_dafny.MultiSet([160, d_1381_v2_]), not(p1)))
                        d_1451_v37_: _dafny.Seq
                        d_1451_v37_ = _dafny.SeqWithoutIsStrInference([d_1396_v15_, d_1396_v15_, d_1396_v15_, d_1396_v15_])
                        d_1396_v15_ = (d_1451_v37_)[default__.safeIndex(default__.safeModuloInt(d_1444_cf33_, (0) - (d_1444_cf33_)), len(d_1451_v37_))]
                    elif source28_.is_DC19:
                        d_1452___mcc_h12_ = source28_.cf34
                        d_1453_cf34_ = d_1452___mcc_h12_
                        d_1453_cf34_ = d_1453_cf34_
                        d_1381_v2_ = default__.fm0(globalState)
                        d_1454_v38_: D9
                        d_1454_v38_ = D9_DC26(d_1381_v2_)
                        (d_1384_v5_).f3 = (((d_1454_v38_).cf44) != (d_1453_cf34_)) or ((d_1382_v3_)[default__.safeIndex(d_1453_cf34_, len(d_1382_v3_))])
                        d_1455_v39_: _dafny.Map
                        d_1455_v39_ = _dafny.Map({d_1381_v2_: True})
                        d_1455_v39_ = _dafny.Map({(d_1380_v1_).cardinality: p2})
                    elif source28_.is_DC16:
                        d_1456___mcc_h13_ = source28_.cf27
                        d_1457_cf27_ = d_1456___mcc_h13_
                        (d_1384_v5_).f3 = ((d_1396_v15_).set(len(d_1393_v12_), default__.abs(d_1381_v2_))).issubset((d_1396_v15_) - (d_1396_v15_))
                        d_1381_v2_ = (0) - (default__.safeModuloInt((d_1381_v2_) - (d_1381_v2_), len(d_1457_cf27_)))
                        d_1458_v40_: _dafny.Array
                        def lambda53_(d_1459_i9_):
                            return (d_1459_i9_) * (len((self).f2))

                        init28_ = lambda53_
                        nw185_ = _dafny.Array(None, 28)
                        for i0_28_ in range(nw185_.length(0)):
                            nw185_[i0_28_] = init28_(i0_28_)
                        d_1458_v40_ = nw185_
                        index218_ = default__.safeIndex(730, (d_1458_v40_).length(0))
                        (d_1458_v40_)[index218_] = default__.safeDivisionInt(d_1381_v2_, d_1381_v2_)
                        index219_ = default__.safeIndex(730, (d_1458_v40_).length(0))
                        (d_1458_v40_)[index219_] = (d_1381_v2_) - (d_1381_v2_)
                        d_1460_v41_: _dafny.Array
                        nw186_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                        d_1460_v41_ = nw186_
                        index220_ = default__.safeIndex(921, (d_1460_v41_).length(0))
                        (d_1460_v41_)[index220_] = ((self).f2) + (default__.fm2(p2, globalState))
                        index221_ = default__.safeIndex(921, (d_1460_v41_).length(0))
                        (d_1460_v41_)[index221_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyiwhk"))).set(default__.safeIndex(len(d_1457_cf27_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyiwhk")))), d_1379_v0_)).set(default__.safeIndex(d_1381_v2_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyiwhk"))).set(default__.safeIndex(len(d_1457_cf27_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyiwhk")))), d_1379_v0_))), (d_1379_v0_ if not(d_1384_v5_.f3) else d_1379_v0_))
                    elif True:
                        d_1461___mcc_h14_ = source28_.cf35
                        d_1462_cf35_ = d_1461___mcc_h14_
                        d_1463_v42_: D1
                        d_1463_v42_ = D1_DC3(p1, d_1381_v2_, d_1381_v2_)
                        (d_1384_v5_).f3 = (454) < ((0) - (((d_1383_v4_)[default__.safeIndex(497, len(d_1383_v4_))] if not(d_1384_v5_.f3) else (d_1463_v42_).cf4)))
                        d_1464_v44_: D6
                        d_1464_v44_ = D6_DC17(not(p0))
                        d_1465_v45_: _dafny.Map
                        d_1465_v45_ = _dafny.Map({d_1464_v44_: d_1381_v2_})
                        d_1466_v46_: D3
                        d_1466_v46_ = D3_DC10(d_1381_v2_, (0) - (d_1381_v2_))
                        d_1467_v47_: _dafny.Array
                        nw187_ = _dafny.Array(None, 28)
                        nw187_[int(0)] = d_1381_v2_
                        nw187_[int(1)] = d_1381_v2_
                        nw187_[int(2)] = (d_1380_v1_).cardinality
                        nw187_[int(3)] = 319
                        nw187_[int(4)] = d_1381_v2_
                        nw187_[int(5)] = d_1381_v2_
                        nw187_[int(6)] = d_1381_v2_
                        def iife165_():
                            coll45_ = _dafny.Map()
                            compr_45_: D6
                            for compr_45_ in (d_1465_v45_).keys.Elements:
                                d_1468_v43_: D6 = compr_45_
                                if (d_1468_v43_) in (d_1465_v45_):
                                    coll45_[d_1468_v43_] = d_1379_v0_
                            return _dafny.Map(coll45_)
                        nw187_[int(7)] = len(iife165_()
                        )
                        nw187_[int(8)] = d_1381_v2_
                        nw187_[int(9)] = len((d_1384_v5_).f2)
                        nw187_[int(10)] = d_1381_v2_
                        nw187_[int(11)] = d_1381_v2_
                        nw187_[int(12)] = d_1381_v2_
                        nw187_[int(13)] = (0) - (((d_1393_v12_)[d_1384_v5_.f3] if (d_1384_v5_.f3) in (d_1393_v12_) else d_1381_v2_))
                        nw187_[int(14)] = d_1381_v2_
                        nw187_[int(15)] = d_1381_v2_
                        nw187_[int(16)] = d_1381_v2_
                        nw187_[int(17)] = d_1381_v2_
                        nw187_[int(18)] = (d_1384_v5_).fm12(d_1381_v2_, len((d_1384_v5_).f2), d_1381_v2_, d_1384_v5_.f3, globalState)
                        nw187_[int(19)] = d_1381_v2_
                        nw187_[int(20)] = d_1381_v2_
                        nw187_[int(21)] = d_1381_v2_
                        nw187_[int(22)] = (d_1380_v1_).cardinality
                        nw187_[int(23)] = d_1381_v2_
                        nw187_[int(24)] = d_1381_v2_
                        nw187_[int(25)] = d_1381_v2_
                        nw187_[int(26)] = (d_1466_v46_).cf17
                        nw187_[int(27)] = d_1381_v2_
                        d_1467_v47_ = nw187_
                        d_1469_v48_: C0
                        nw188_ = C0()
                        nw188_.ctor__(d_1467_v47_)
                        d_1469_v48_ = nw188_
                        d_1386_v7_ = d_1386_v7_
                        d_1381_v2_ = (self).fm12(-956, (d_1381_v2_) - (d_1381_v2_), d_1381_v2_, self.f3, globalState)
                    d_1470_v49_: D2
                    d_1470_v49_ = D2_DC6(_dafny.CodePoint('v'), p0, d_1384_v5_.f3, d_1381_v2_)
                    source29_ = d_1470_v49_
                    if source29_.is_DC6:
                        d_1471___mcc_h15_ = source29_.cf9
                        d_1472___mcc_h16_ = source29_.cf10
                        d_1473___mcc_h17_ = source29_.cf11
                        d_1474___mcc_h18_ = source29_.cf12
                        d_1475_cf12_ = d_1474___mcc_h18_
                        d_1476_cf11_ = d_1473___mcc_h17_
                        d_1477_cf10_ = d_1472___mcc_h16_
                        d_1478_cf9_ = d_1471___mcc_h15_
                        d_1479_v50_: _dafny.Map
                        d_1479_v50_ = _dafny.Map({d_1477_cf10_: d_1478_cf9_})
                        d_1480_v51_: _dafny.Map
                        d_1480_v51_ = _dafny.Map({(len(d_1479_v50_)) - (120): d_1383_v4_})
                        d_1383_v4_ = ((d_1480_v51_)[len(d_1391_v10_)] if (len(d_1391_v10_)) in (d_1480_v51_) else d_1383_v4_)
                        rhs214_ = (0) - (default__.safeDivisionInt(179, (d_1381_v2_ if p0 else 249)))
                        rhs215_ = (((d_1396_v15_)[len(d_1382_v3_)] if (len(d_1382_v3_)) in (d_1396_v15_) else default__.fm0(globalState))) + ((-47 if d_1384_v5_.f3 else len(_dafny.SeqWithoutIsStrInference([len((d_1384_v5_).f2) for d_1481_i10_ in range(default__.abs(8))]))))
                        d_1475_cf12_ = rhs214_
                        d_1381_v2_ = rhs215_
                        d_1381_v2_ = -84
                        d_1482_v52_: _dafny.Set
                        d_1482_v52_ = _dafny.Set({d_1381_v2_})
                        rhs216_ = d_1477_cf10_
                        rhs217_ = True
                        rhs218_ = d_1383_v4_
                        rhs219_ = p2
                        rhs220_ = (p2 if (d_1381_v2_) in (d_1482_v52_) else (p0 if False else not(d_1384_v5_.f3)))
                        lhs120_ = d_1384_v5_
                        lhs121_ = self
                        lhs122_ = d_1384_v5_
                        lhs123_ = d_1384_v5_
                        lhs120_.f3 = rhs216_
                        lhs121_.f3 = rhs217_
                        d_1383_v4_ = rhs218_
                        lhs122_.f3 = rhs219_
                        lhs123_.f3 = rhs220_
                    elif source29_.is_DC5:
                        d_1483___mcc_h19_ = source29_.cf8
                        d_1484_cf8_ = d_1483___mcc_h19_
                        (d_1384_v5_).f3 = (d_1379_v0_) not in (((self).f2) + ((d_1384_v5_).f2))
                        d_1485_v53_: D11
                        d_1485_v53_ = D11_DC31(d_1381_v2_)
                        d_1486_v54_: _dafny.Map
                        d_1486_v54_ = _dafny.Map({default__.fm49(d_1485_v53_, (d_1380_v1_).cardinality, d_1381_v2_, globalState): d_1381_v2_})
                        d_1487_v55_: _dafny.Map
                        d_1487_v55_ = _dafny.Map({d_1486_v54_: p0})
                        d_1487_v55_ = (d_1487_v55_).set((d_1486_v54_) | (d_1486_v54_), (False) or (self.f3))
                        d_1484_cf8_ = (d_1484_cf8_).set(p0, d_1381_v2_)
                        rhs221_ = p0
                        rhs222_ = not(not(((d_1384_v5_).fm11(not((self).fm11(default__.fm3(d_1381_v2_, globalState), globalState)), globalState)) and (d_1384_v5_.f3)))
                        lhs124_ = d_1384_v5_
                        lhs125_ = d_1384_v5_
                        lhs124_.f3 = rhs221_
                        lhs125_.f3 = rhs222_
                    elif True:
                        d_1488___mcc_h20_ = source29_.cf13
                        d_1489_cf13_ = d_1488___mcc_h20_
                        d_1490_v56_: C2
                        nw189_ = C2()
                        nw189_.ctor__(d_1381_v2_, False, (d_1384_v5_).f1, (d_1384_v5_).f2)
                        d_1490_v56_ = nw189_
                        d_1491_v57_: C1
                        nw190_ = C1()
                        nw190_.ctor__()
                        d_1491_v57_ = nw190_
                        d_1492_v58_: _dafny.Array
                        nw191_ = _dafny.Array(_dafny.Seq({}), 15)
                        d_1492_v58_ = nw191_
                        rhs223_ = (d_1490_v56_).f13
                        rhs224_ = p2
                        rhs225_ = default__.safeModuloInt((0) - (d_1381_v2_), d_1381_v2_)
                        rhs226_ = (d_1492_v58_ if d_1384_v5_.f3 else d_1492_v58_)
                        rhs227_ = p0
                        lhs126_ = d_1384_v5_
                        lhs127_ = d_1384_v5_
                        d_1381_v2_ = rhs223_
                        lhs126_.f3 = rhs224_
                        d_1381_v2_ = rhs225_
                        d_1492_v58_ = rhs226_
                        lhs127_.f3 = rhs227_
                        (d_1384_v5_).f3 = p1
                    d_1381_v2_ = d_1381_v2_
                    if False:
                        d_1493_v59_: _dafny.Array
                        def lambda54_(d_1494_v2_):
                            def lambda55_(d_1495_i11_):
                                return (d_1495_i11_) * (d_1494_v2_)

                            return lambda55_

                        init29_ = lambda54_(d_1381_v2_)
                        nw192_ = _dafny.Array(None, 22)
                        for i0_29_ in range(nw192_.length(0)):
                            nw192_[i0_29_] = init29_(i0_29_)
                        d_1493_v59_ = nw192_
                        index222_ = default__.safeIndex(845, (d_1493_v59_).length(0))
                        (d_1493_v59_)[index222_] = d_1381_v2_
                        index223_ = default__.safeIndex(845, (d_1493_v59_).length(0))
                        (d_1493_v59_)[index223_] = d_1381_v2_
                        (self).f3 = ((d_1384_v5_).f2) not in (_dafny.SeqWithoutIsStrInference([(d_1384_v5_).f2 for d_1496_i12_ in range(default__.abs(891))]))
                        rhs228_ = (d_1381_v2_) - ((d_1493_v59_)[default__.safeIndex(845, (d_1493_v59_).length(0))])
                        rhs229_ = p1
                        rhs230_ = d_1379_v0_
                        lhs128_ = d_1384_v5_
                        d_1381_v2_ = rhs228_
                        lhs128_.f3 = rhs229_
                        d_1379_v0_ = rhs230_
                        d_1497_v60_: _dafny.Seq
                        d_1497_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjg"))
                        d_1497_v60_ = (d_1384_v5_).f2
                        d_1498_v61_: _dafny.Map
                        d_1498_v61_ = _dafny.Map({d_1381_v2_: (d_1382_v3_)[default__.safeIndex((0) - ((d_1493_v59_)[default__.safeIndex(845, (d_1493_v59_).length(0))]), len(d_1382_v3_))]})
                        d_1498_v61_ = d_1498_v61_
                    elif True:
                        d_1499_v62_: _dafny.Array
                        def lambda56_(d_1500_v0_, d_1501_v5_, d_1502_v2_):
                            def lambda57_(d_1503_i13_):
                                return (d_1503_i13_) * ((D2_DC6(d_1500_v0_, d_1501_v5_.f3, d_1501_v5_.f3, d_1502_v2_)).cf12)

                            return lambda57_

                        init30_ = lambda56_(d_1379_v0_, d_1384_v5_, d_1381_v2_)
                        nw193_ = _dafny.Array(None, 27)
                        for i0_30_ in range(nw193_.length(0)):
                            nw193_[i0_30_] = init30_(i0_30_)
                        d_1499_v62_ = nw193_
                        d_1504_v63_: _dafny.Set
                        d_1504_v63_ = _dafny.Set({-299})
                        index224_ = default__.safeIndex(377, (d_1499_v62_).length(0))
                        (d_1499_v62_)[index224_] = len(_dafny.Map({len(d_1504_v63_): p1}))
                        index225_ = default__.safeIndex(377, (d_1499_v62_).length(0))
                        (d_1499_v62_)[index225_] = d_1381_v2_
                        index226_ = default__.safeIndex(377, (d_1499_v62_).length(0))
                        (d_1499_v62_)[index226_] = (771) * (((d_1499_v62_)[default__.safeIndex(377, (d_1499_v62_).length(0))] if d_1384_v5_.f3 else d_1381_v2_))
                        d_1505_v64_: C1
                        nw194_ = C1()
                        nw194_.ctor__()
                        d_1505_v64_ = nw194_
                        d_1415_v21_ = D6_DC20(d_1413_v19_)
                        d_1506_v65_: C2
                        nw195_ = C2()
                        nw195_.ctor__(d_1381_v2_, p0, (d_1384_v5_).f1, ((d_1384_v5_).f2).set(default__.safeIndex(d_1381_v2_, len((d_1384_v5_).f2)), d_1379_v0_))
                        d_1506_v65_ = nw195_
                        d_1507_v66_: _dafny.Array
                        nw196_ = _dafny.Array(None, 16)
                        nw196_[int(0)] = (d_1506_v65_ if d_1384_v5_.f3 else d_1506_v65_)
                        nw196_[int(1)] = d_1506_v65_
                        nw196_[int(2)] = d_1506_v65_
                        nw196_[int(3)] = (d_1506_v65_ if p0 else d_1506_v65_)
                        nw196_[int(4)] = d_1506_v65_
                        nw196_[int(5)] = d_1506_v65_
                        nw196_[int(6)] = d_1506_v65_
                        nw196_[int(7)] = d_1506_v65_
                        nw196_[int(8)] = d_1506_v65_
                        nw196_[int(9)] = d_1506_v65_
                        nw196_[int(10)] = d_1506_v65_
                        nw196_[int(11)] = (D13_DC33(d_1506_v65_)).cf52
                        nw196_[int(12)] = (d_1506_v65_ if p0 else d_1506_v65_)
                        nw196_[int(13)] = d_1506_v65_
                        nw196_[int(14)] = d_1506_v65_
                        nw196_[int(15)] = d_1506_v65_
                        d_1507_v66_ = nw196_
                        index227_ = default__.safeIndex(235, (d_1507_v66_).length(0))
                        (d_1507_v66_)[index227_] = d_1506_v65_
                        index228_ = default__.safeIndex(235, (d_1507_v66_).length(0))
                        rhs231_ = d_1384_v5_.f3
                        rhs232_ = d_1506_v65_
                        lhs129_ = d_1384_v5_
                        lhs130_ = d_1507_v66_
                        lhs131_ = default__.safeIndex(235, (d_1507_v66_).length(0))
                        lhs129_.f3 = rhs231_
                        lhs130_[lhs131_] = rhs232_
                    pass
            pass
        d_1508_i14_: int
        d_1508_i14_ = 0
        with _dafny.label("10"):
            while d_1384_v5_.f3:
                with _dafny.c_label("10"):
                    if (d_1508_i14_) >= (100):
                        raise _dafny.Break("10")
                    d_1508_i14_ = (d_1508_i14_) + (1)
                    d_1509_v67_: _dafny.Seq
                    d_1509_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aw"))
                    d_1509_v67_ = ((d_1384_v5_).f2) + ((d_1509_v67_) + ((d_1384_v5_).f2))
                    d_1381_v2_ = default__.fm0(globalState)
                    d_1510_v68_: D5
                    d_1510_v68_ = D5_DC15(d_1381_v2_, d_1384_v5_.f3, d_1386_v7_)
                    (d_1384_v5_).f3 = ((d_1510_v68_).cf24) < (d_1381_v2_)
                    (self).f3 = p0
                    pass
            pass
        hi6_ = (0) - (d_1381_v2_)
        for d_1511_i15_ in range((d_1381_v2_) * (d_1381_v2_), hi6_):
            d_1512_v69_: _dafny.Array
            nw197_ = _dafny.Array(int(0), 4)
            d_1512_v69_ = nw197_
            index229_ = default__.safeIndex(90, (d_1512_v69_).length(0))
            (d_1512_v69_)[index229_] = d_1511_i15_
            d_1513_v70_: D1
            d_1513_v70_ = D1_DC3(d_1384_v5_.f3, d_1511_i15_, (self).fm12(d_1511_i15_, d_1381_v2_, d_1381_v2_, p1, globalState))
            index230_ = default__.safeIndex(90, (d_1512_v69_).length(0))
            (d_1512_v69_)[index230_] = ((len((d_1384_v5_).f2)) + ((d_1513_v70_).cf5)) + (d_1381_v2_)
            (self).f3 = (d_1384_v5_.f3) or (d_1384_v5_.f3)
            d_1514_v71_: _dafny.Map
            d_1514_v71_ = _dafny.Map({_dafny.CodePoint('m'): d_1386_v7_})
            rhs233_ = (d_1511_i15_) <= (d_1511_i15_)
            rhs234_ = ((d_1514_v71_)[d_1379_v0_] if (d_1379_v0_) in (d_1514_v71_) else d_1386_v7_)
            rhs235_ = p2
            rhs236_ = (not (False) or (p0)) == (not(d_1384_v5_.f3))
            lhs132_ = self
            lhs133_ = d_1384_v5_
            lhs134_ = self
            lhs132_.f3 = rhs233_
            d_1386_v7_ = rhs234_
            lhs133_.f3 = rhs235_
            lhs134_.f3 = rhs236_
            d_1381_v2_ = (0) - (((d_1512_v69_)[default__.safeIndex(90, (d_1512_v69_).length(0))] if (d_1381_v2_) >= (-781) else default__.safeModuloInt((d_1512_v69_)[default__.safeIndex(90, (d_1512_v69_).length(0))], d_1511_i15_)))


class C5(T0):
    def  __init__(self):
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    def ctor__(self, f15, f1, f2):
        (self)._f15 = f15
        (self)._f1 = f1
        (self)._f2 = f2

    def fm11(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({-315: -858}))), len((self).f2), 204]) if (self).f15 else _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwmv"))): len((self).f2)}))) for d_1515_i0_ in range(default__.abs(361))]))) == (_dafny.SeqWithoutIsStrInference([-970]))

    @property
    def f15(self):
        return self._f15

class C6(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self.f12: int = int(0)
        self._f11: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f11, f12, f4, f5, f3, f1, f2):
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return ((_dafny.MultiSet([self.f3])) | (_dafny.MultiSet([self.f3, self.f3]))).issubset(_dafny.MultiSet([self.f3, self.f3, self.f3, self.f3]))

    def fm14(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iweafjioe"))) + ((self).f2)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kruhckdc")))

    def fm12(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm11(self, p0, globalState):
        if self.f3:
            return self.f3
        elif True:
            return (True) or (self.f3)

    def fm21(self, p0, p1, globalState):
        return default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1516_i0_ in range(default__.abs(215))])), self.f12)

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        if self.f3:
            d_1517_v0_: str
            d_1517_v0_ = _dafny.CodePoint('h')
            d_1518_v2_: D3
            def iife166_():
                coll46_ = _dafny.Set()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(493, 7):
                    d_1520_v1_: int = compr_46_
                    if ((493) <= (d_1520_v1_)) and ((d_1520_v1_) < (7)):
                        coll46_ = coll46_.union(_dafny.Set([default__.safeModuloInt(d_1520_v1_, p1)]))
                return _dafny.Set(coll46_)
            d_1518_v2_ = D3_DC10(len(_dafny.SeqWithoutIsStrInference([(self).f11 for d_1519_i0_ in range(default__.abs(687))])), len(iife166_()
))
            d_1521_v3_: D3
            d_1521_v3_ = D3_DC11(d_1518_v2_)
            d_1522_v4_: _dafny.MultiSet
            d_1522_v4_ = _dafny.MultiSet([d_1521_v3_])
            d_1523_v5_: _dafny.Array
            def lambda58_(d_1524_p0_):
                def lambda59_(d_1525_i1_):
                    return default__.safeModuloInt(d_1525_i1_, (0) - (d_1524_p0_))

                return lambda59_

            init31_ = lambda58_(p0)
            nw198_ = _dafny.Array(None, 5)
            for i0_31_ in range(nw198_.length(0)):
                nw198_[i0_31_] = init31_(i0_31_)
            d_1523_v5_ = nw198_
            d_1526_v6_: _dafny.MultiSet
            d_1526_v6_ = _dafny.MultiSet([self.f3, self.f3])
            d_1527_v7_: _dafny.Seq
            d_1527_v7_ = _dafny.SeqWithoutIsStrInference([self.f3])
            rhs237_ = default__.fm22(self.f3, len((self).f2), globalState)
            rhs238_ = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([default__.fm23((d_1526_v6_).cardinality, (self).f11, self.f12, d_1517_v0_, globalState), d_1521_v3_, d_1521_v3_, d_1521_v3_, d_1521_v3_])).set(default__.safeIndex((0) - (default__.safeModuloInt(self.f12, (self).fm12((self).f4, self.f12, len(_dafny.Map({self.f3: (self).f11})), self.f3, globalState))), len(_dafny.SeqWithoutIsStrInference([default__.fm23((d_1526_v6_).cardinality, (self).f11, self.f12, d_1517_v0_, globalState), d_1521_v3_, d_1521_v3_, d_1521_v3_, d_1521_v3_]))), d_1521_v3_))
            rhs239_ = (len(d_1527_v7_)) * (default__.safeModuloInt((self).f4, (self).f11))
            rhs240_ = d_1523_v5_
            rhs241_ = self.f3
            lhs135_ = self
            d_1517_v0_ = rhs237_
            d_1522_v4_ = rhs238_
            r1 = rhs239_
            d_1523_v5_ = rhs240_
            lhs135_.f3 = rhs241_
            (self).f3 = self.f3
            (self).f3 = not(self.f3)
            d_1528_v8_: _dafny.Map
            d_1528_v8_ = _dafny.Map({p0: self.f3})
            d_1529_v9_: _dafny.MultiSet
            d_1529_v9_ = _dafny.MultiSet([(0) - (len(default__.fm2(self.f3, globalState)))])
            d_1530_v10_: _dafny.Map
            d_1530_v10_ = _dafny.Map({self.f12: d_1529_v9_})
            d_1531_v11_: _dafny.Set
            d_1531_v11_ = _dafny.Set({self.f3})
            d_1532_v12_: _dafny.Seq
            d_1532_v12_ = _dafny.SeqWithoutIsStrInference([922])
            d_1528_v8_ = (_dafny.Map({len((d_1530_v10_).set(-641, _dafny.MultiSet([len(d_1531_v11_), 495, self.f12, self.f12, (d_1532_v12_)[default__.safeIndex((self).f11, len(d_1532_v12_))]]))): False})) | (d_1528_v8_)
            d_1533_v13_: D5
            d_1533_v13_ = D5_DC14(d_1523_v5_)
            d_1533_v13_ = (d_1533_v13_ if self.f3 else d_1533_v13_)
        elif True:
            d_1534_v14_: D3
            d_1534_v14_ = D3_DC9(len((self).f2), self.f3)
            index231_ = default__.safeIndex(703, ((self).f5).length(0))
            ((self).f5)[index231_] = (d_1534_v14_).cf16
            d_1535_v15_: _dafny.MultiSet
            d_1535_v15_ = _dafny.MultiSet([(self).f4])
            d_1536_v16_: _dafny.Map
            d_1536_v16_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f4: (self).f11})) for d_1537_i2_ in range(default__.abs(754))])) == (_dafny.SeqWithoutIsStrInference([(d_1535_v15_).cardinality, (self).f11, p0])): self.f3})
            index232_ = default__.safeIndex(703, ((self).f5).length(0))
            ((self).f5)[index232_] = ((d_1536_v16_)[self.f3] if (self.f3) in (d_1536_v16_) else self.f3)
            d_1538_v17_: _dafny.Map
            d_1538_v17_ = _dafny.Map({(default__.fm8(p0, (self).f2, globalState)).cardinality: d_1536_v16_})
            d_1539_v18_: _dafny.Map
            d_1539_v18_ = _dafny.Map({self.f3: 708})
            d_1540_v19_: D1
            d_1540_v19_ = D1_DC3(((self).f5)[default__.safeIndex(703, ((self).f5).length(0))], p0, (self).f11)
            d_1541_v20_: _dafny.Seq
            d_1541_v20_ = _dafny.SeqWithoutIsStrInference([(d_1540_v19_).cf3])
            d_1542_v21_: _dafny.Array
            nw199_ = _dafny.Array(None, 9)
            nw199_[int(0)] = (self).f11
            nw199_[int(1)] = default__.fm0(globalState)
            nw199_[int(2)] = self.f12
            nw199_[int(3)] = len(d_1539_v18_)
            nw199_[int(4)] = len(d_1541_v20_)
            nw199_[int(5)] = len(default__.fm9(globalState))
            nw199_[int(6)] = (self).f11
            nw199_[int(7)] = -422
            nw199_[int(8)] = self.f12
            d_1542_v21_ = nw199_
            d_1543_v22_: _dafny.Set
            d_1543_v22_ = _dafny.Set({d_1542_v21_, d_1542_v21_})
            d_1544_v23_: _dafny.Map
            d_1544_v23_ = _dafny.Map({d_1538_v17_: d_1543_v22_})
            def iife167_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in _dafny.IntegerRange(984, 273):
                    d_1545_v24_: int = compr_47_
                    if ((984) <= (d_1545_v24_)) and ((d_1545_v24_) < (273)):
                        coll47_[(d_1545_v24_) * (len(_dafny.SeqWithoutIsStrInference([p0])))] = d_1536_v16_
                return _dafny.Map(coll47_)
            d_1544_v23_ = (d_1544_v23_).set((d_1538_v17_) | (iife167_()
            ), d_1543_v22_)
            if self.f3:
                d_1546_v25_: _dafny.Array
                def lambda60_(d_1547_p0_):
                    def lambda61_(d_1548_i3_):
                        return _dafny.SeqWithoutIsStrInference([((self).f2)[default__.safeIndex(d_1547_p0_, len((self).f2))] for d_1549_i4_ in range(default__.abs(304))])

                    return lambda61_

                init32_ = lambda60_(p0)
                nw200_ = _dafny.Array(None, 13)
                for i0_32_ in range(nw200_.length(0)):
                    nw200_[i0_32_] = init32_(i0_32_)
                d_1546_v25_ = nw200_
                index233_ = default__.safeIndex(782, (d_1546_v25_).length(0))
                (d_1546_v25_)[index233_] = ((self).f2) + ((self).f2)
                index234_ = default__.safeIndex(782, (d_1546_v25_).length(0))
                (d_1546_v25_)[index234_] = (self).f2
                r1 = (default__.safeDivisionInt((self).fm12(p0, len(_dafny.Map({p0: p1})), len((self).f2), ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))], globalState), (self).f4)) * ((self).f4)
                index235_ = default__.safeIndex(703, ((self).f5).length(0))
                ((self).f5)[index235_] = self.f3
                (self).f3 = (self).fm11(self.f3, globalState)
                d_1550_v26_: _dafny.Set
                d_1550_v26_ = _dafny.Set({d_1536_v16_})
                d_1551_v27_: _dafny.Map
                d_1551_v27_ = _dafny.Map({d_1536_v16_: p0})
                index236_ = default__.safeIndex(703, ((self).f5).length(0))
                def iife168_():
                    coll48_ = _dafny.Set()
                    compr_48_: _dafny.Map
                    for compr_48_ in (d_1551_v27_).keys.Elements:
                        d_1552_v28_: _dafny.Map = compr_48_
                        if (d_1552_v28_) in (d_1551_v27_):
                            coll48_ = coll48_.union(_dafny.Set([d_1552_v28_]))
                    return _dafny.Set(coll48_)
                ((self).f5)[index236_] = (d_1550_v26_).isdisjoint(iife168_()
                )
            elif True:
                index237_ = default__.safeIndex(703, ((self).f5).length(0))
                ((self).f5)[index237_] = self.f3
                d_1553_v29_: _dafny.Array
                nw201_ = _dafny.Array(None, 22)
                nw201_[int(0)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(1)] = self.f3
                nw201_[int(2)] = self.f3
                nw201_[int(3)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(4)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(5)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(6)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(7)] = self.f3
                nw201_[int(8)] = self.f3
                nw201_[int(9)] = self.f3
                nw201_[int(10)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(11)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(12)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(13)] = self.f3
                nw201_[int(14)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(15)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(16)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                nw201_[int(17)] = self.f3
                nw201_[int(18)] = default__.fm3(p0, globalState)
                nw201_[int(19)] = self.f3
                nw201_[int(20)] = not(not(self.f3))
                nw201_[int(21)] = ((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]
                d_1553_v29_ = nw201_
                d_1554_v30_: _dafny.Map
                d_1554_v30_ = _dafny.Map({d_1553_v29_: self.f3})
                d_1555_v31_: bool
                d_1556_v32_: bool
                d_1557_v33_: _dafny.Seq
                d_1558_v34_: _dafny.Seq
                out29_: bool
                out30_: bool
                out31_: _dafny.Seq
                out32_: _dafny.Seq
                out29_, out30_, out31_, out32_ = default__.m0(d_1554_v30_, globalState)
                d_1555_v31_ = out29_
                d_1556_v32_ = out30_
                d_1557_v33_ = out31_
                d_1558_v34_ = out32_
                d_1559_v35_: D3
                d_1559_v35_ = D3_DC10(p1, len((self).f2))
                d_1559_v35_ = d_1559_v35_
                d_1560_v36_: D6
                d_1560_v36_ = D6_DC17(((self).f5)[default__.safeIndex(703, ((self).f5).length(0))])
                d_1561_v37_: _dafny.MultiSet
                d_1561_v37_ = _dafny.MultiSet([not(((self).f5)[default__.safeIndex(703, ((self).f5).length(0))]), d_1556_v32_, (d_1560_v36_).cf28])
                d_1561_v37_ = d_1561_v37_
                d_1562_v38_: D5
                d_1562_v38_ = D5_DC14(d_1542_v21_)
                d_1563_v39_: C0
                nw202_ = C0()
                nw202_.ctor__((d_1562_v38_).cf23)
                d_1563_v39_ = nw202_
            d_1564_v40_: C1
            nw203_ = C1()
            nw203_.ctor__()
            d_1564_v40_ = nw203_
            index238_ = default__.safeIndex(703, ((self).f5).length(0))
            ((self).f5)[index238_] = (self).fm13(globalState)
        d_1565_v41_: str
        d_1565_v41_ = _dafny.CodePoint('s')
        d_1566_v42_: _dafny.Map
        d_1566_v42_ = _dafny.Map({d_1565_v41_: (-615 if True else p1)})
        d_1566_v42_ = (d_1566_v42_).set(((self).f2)[default__.safeIndex((self).f4, len((self).f2))], (self).f4)
        d_1567_v43_: C2
        nw204_ = C2()
        nw204_.ctor__((self).f4, self.f3, (self).f1, (self).f2)
        d_1567_v43_ = nw204_
        d_1568_v44_: D13
        d_1568_v44_ = D13_DC33(d_1567_v43_)
        source30_ = d_1568_v44_
        if source30_.is_DC34:
            d_1569___mcc_h0_ = source30_.cf53
            d_1570___mcc_h1_ = source30_.cf54
            d_1571___mcc_h2_ = source30_.cf55
            d_1572___mcc_h3_ = source30_.cf56
            d_1573_cf56_ = d_1572___mcc_h3_
            d_1574_cf55_ = d_1571___mcc_h2_
            d_1575_cf54_ = d_1570___mcc_h1_
            d_1576_cf53_ = d_1569___mcc_h0_
            index239_ = default__.safeIndex(504, ((self).f5).length(0))
            ((self).f5)[index239_] = d_1575_cf54_
            index240_ = default__.safeIndex(504, ((self).f5).length(0))
            ((self).f5)[index240_] = (d_1565_v41_) not in ((self).f2)
            d_1577_v45_: _dafny.Set
            d_1577_v45_ = _dafny.Set({d_1565_v41_})
            d_1578_v46_: _dafny.Seq
            d_1578_v46_ = _dafny.SeqWithoutIsStrInference([d_1577_v45_])
            d_1579_v47_: _dafny.Map
            d_1579_v47_ = _dafny.Map({d_1574_cf55_: len((d_1578_v46_)[default__.safeIndex(p1, len(d_1578_v46_))])})
            d_1579_v47_ = (d_1579_v47_).set(default__.fm0(globalState), (self.f12) - (d_1574_cf55_))
            d_1565_v41_ = d_1565_v41_
            (self).f3 = self.f3
        elif source30_.is_DC35:
            d_1580___mcc_h4_ = source30_.cf57
            d_1581___mcc_h5_ = source30_.cf58
            d_1582___mcc_h6_ = source30_.cf59
            d_1583___mcc_h7_ = source30_.cf60
            d_1584___mcc_h8_ = source30_.cf61
            d_1585_cf61_ = d_1584___mcc_h8_
            d_1586_cf60_ = d_1583___mcc_h7_
            d_1587_cf59_ = d_1582___mcc_h6_
            d_1588_cf58_ = d_1581___mcc_h5_
            d_1589_cf57_ = d_1580___mcc_h4_
            (self).f3 = (((self).f2) + ((self).f2)) <= ((self).f2)
            d_1590_v48_: C4
            nw205_ = C4()
            nw205_.ctor__((d_1565_v41_) in ((self).f2), (self).f1, (self).f2)
            d_1590_v48_ = nw205_
            (self).f3 = d_1587_cf59_
            d_1591_v49_: C4
            nw206_ = C4()
            nw206_.ctor__(not(self.f3), (self).f1, (self).f2)
            d_1591_v49_ = nw206_
        elif source30_.is_DC33:
            d_1592___mcc_h9_ = source30_.cf52
            d_1593_cf52_ = d_1592___mcc_h9_
            d_1594_v50_: _dafny.Seq
            d_1594_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agpskpj"))
            d_1594_v50_ = d_1594_v50_
            d_1595_v51_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
            d_1595_v51_ = nw207_
            d_1595_v51_ = d_1595_v51_
            (self).f3 = (d_1567_v43_).fm11(self.f3, globalState)
            d_1596_v52_: D6
            d_1596_v52_ = D6_DC18(self.f12, self.f12, self.f3, d_1565_v41_, self.f12)
            d_1597_v53_: D6
            d_1597_v53_ = D6_DC20(d_1596_v52_)
            d_1597_v53_ = d_1597_v53_
        elif True:
            d_1598___mcc_h10_ = source30_.cf62
            d_1599_cf62_ = d_1598___mcc_h10_
            (self).f12 = -574
            (self).f3 = self.f3
            (self).f12 = (self).f11
            d_1600_v54_: _dafny.Map
            d_1600_v54_ = _dafny.Map({True: p0})
            d_1601_v55_: _dafny.MultiSet
            d_1601_v55_ = _dafny.MultiSet([(self).f2])
            (self).f3 = (p1) == ((len(d_1600_v54_) if self.f3 else (0) - ((d_1567_v43_).fm27(p1, self.f3, d_1601_v55_, globalState))))
        r1 = default__.safeDivisionInt(p0, self.f12)
        r1 = 575
        (self).f3 = (self.f3) or (self.f3)
        d_1602_v56_: _dafny.Map
        d_1602_v56_ = _dafny.Map({self.f3: self.f3})
        r0 = d_1602_v56_
        r1 = (0) - ((p1) + ((d_1567_v43_).f13))
        r2 = (default__.safeModuloInt(p0, (self).f11)) * ((d_1567_v43_).f13)
        return r0, r1, r2

    def m7(self, p0, p1, p2, p3, globalState):
        (self).f12 = default__.safeModuloInt(p2, default__.safeDivisionInt((self).f11, 612))
        d_1603_v0_: D6
        d_1603_v0_ = D6_DC19((self).f11)
        d_1604_v1_: D13
        d_1604_v1_ = D13_DC34(p0, self.f3, ((d_1603_v0_).cf34 if self.f3 else self.f12), default__.fm2((self).fm13(globalState), globalState))
        source31_ = d_1604_v1_
        if source31_.is_DC34:
            d_1605___mcc_h0_ = source31_.cf53
            d_1606___mcc_h1_ = source31_.cf54
            d_1607___mcc_h2_ = source31_.cf55
            d_1608___mcc_h3_ = source31_.cf56
            d_1609_cf56_ = d_1608___mcc_h3_
            d_1610_cf55_ = d_1607___mcc_h2_
            d_1611_cf54_ = d_1606___mcc_h1_
            d_1612_cf53_ = d_1605___mcc_h0_
            d_1613_v2_: str
            d_1613_v2_ = _dafny.CodePoint('j')
            d_1609_cf56_ = (((self).f2) + ((d_1609_cf56_) + ((self).f2))).set(default__.safeIndex((self).f4, len(((self).f2) + ((d_1609_cf56_) + ((self).f2)))), d_1613_v2_)
            if d_1612_cf53_:
                d_1614_v3_: D2
                d_1614_v3_ = D2_DC6(d_1613_v2_, p0, d_1612_cf53_, d_1610_cf55_)
                d_1615_v4_: D5
                d_1615_v4_ = D5_DC15((-841) * ((self).fm12(self.f12, (self).f11, self.f12, (d_1614_v3_).cf10, globalState)), d_1612_cf53_, (self).f5)
                d_1616_v5_: _dafny.Seq
                d_1616_v5_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5, (self).f5])
                def iife169_(_pat_let60_0):
                    def iife170_(d_1617_dt__update__tmp_h0_):
                        def iife171_(_pat_let61_0):
                            def iife172_(d_1618_dt__update_hcf24_h0_):
                                return D5_DC15(d_1618_dt__update_hcf24_h0_, (d_1617_dt__update__tmp_h0_).cf25, (d_1617_dt__update__tmp_h0_).cf26)
                            return iife172_(_pat_let61_0)
                        return iife171_((self).f4)
                    return iife170_(_pat_let60_0)
                d_1615_v4_ = iife169_(D5_DC15(default__.fm0(globalState), d_1612_cf53_, (d_1616_v5_)[default__.safeIndex(self.f12, len(d_1616_v5_))]))
                (self).f3 = d_1612_cf53_
                d_1619_v6_: _dafny.Map
                d_1619_v6_ = _dafny.Map({(self).f5: self.f3})
                d_1620_v7_: bool
                d_1621_v8_: bool
                d_1622_v9_: _dafny.Seq
                d_1623_v10_: _dafny.Seq
                out33_: bool
                out34_: bool
                out35_: _dafny.Seq
                out36_: _dafny.Seq
                out33_, out34_, out35_, out36_ = default__.m0((d_1619_v6_) | (_dafny.Map({(self).f5: d_1612_cf53_})), globalState)
                d_1620_v7_ = out33_
                d_1621_v8_ = out34_
                d_1622_v9_ = out35_
                d_1623_v10_ = out36_
                d_1624_v11_: C1
                nw208_ = C1()
                nw208_.ctor__()
                d_1624_v11_ = nw208_
                d_1625_v12_: _dafny.Array
                def lambda62_(d_1626_cf56_):
                    def lambda63_(d_1627_i0_):
                        return (D1_DC2(d_1626_cf56_)).cf2

                    return lambda63_

                init33_ = lambda62_(d_1609_cf56_)
                nw209_ = _dafny.Array(None, 14)
                for i0_33_ in range(nw209_.length(0)):
                    nw209_[i0_33_] = init33_(i0_33_)
                d_1625_v12_ = nw209_
                nw210_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_1625_v12_ = nw210_
            elif True:
                d_1628_v13_: C4
                nw211_ = C4()
                nw211_.ctor__(d_1612_cf53_, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wttr")))
                d_1628_v13_ = nw211_
                d_1629_v14_: _dafny.Seq
                d_1629_v14_ = _dafny.SeqWithoutIsStrInference([d_1610_cf55_, d_1610_cf55_])
                (self).f12 = (d_1629_v14_)[default__.safeIndex(default__.fm0(globalState), len(d_1629_v14_))]
                d_1630_v15_: _dafny.Array
                nw212_ = _dafny.Array(int(0), 5)
                d_1630_v15_ = nw212_
                d_1631_v16_: _dafny.MultiSet
                d_1631_v16_ = _dafny.MultiSet([d_1630_v15_])
                d_1632_v17_: _dafny.Map
                d_1632_v17_ = _dafny.Map({(self).f2: ((d_1631_v16_)[d_1630_v15_] if (d_1630_v15_) in (d_1631_v16_) else len((d_1629_v14_).set(default__.safeIndex(self.f12, len(d_1629_v14_)), self.f12)))})
                d_1633_v18_: _dafny.Map
                d_1633_v18_ = _dafny.Map({(self).f4: self.f3})
                d_1610_cf55_ = (default__.safeModuloInt(d_1610_cf55_, ((d_1632_v17_)[d_1609_cf56_] if (d_1609_cf56_) in (d_1632_v17_) else d_1610_cf55_))) + (len(d_1633_v18_))
                d_1634_v19_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Map({}), 9)
                d_1634_v19_ = nw213_
                d_1635_v20_: _dafny.MultiSet
                d_1635_v20_ = _dafny.MultiSet([p0])
                d_1636_v21_: _dafny.Map
                d_1636_v21_ = _dafny.Map({d_1635_v20_: p0})
                index241_ = default__.safeIndex(147, (d_1634_v19_).length(0))
                (d_1634_v19_)[index241_] = (d_1636_v21_).set(d_1635_v20_, p0)
                d_1637_v23_: _dafny.Seq
                d_1637_v23_ = _dafny.SeqWithoutIsStrInference([d_1635_v20_])
                index242_ = default__.safeIndex(147, (d_1634_v19_).length(0))
                def iife173_():
                    coll49_ = _dafny.Map()
                    compr_49_: _dafny.MultiSet
                    for compr_49_ in ((d_1637_v23_) + (d_1637_v23_)).Elements:
                        d_1638_v22_: _dafny.MultiSet = compr_49_
                        if (d_1638_v22_) in ((d_1637_v23_) + (d_1637_v23_)):
                            coll49_[d_1638_v22_] = (d_1609_cf56_) < ((self).f2)
                    return _dafny.Map(coll49_)
                rhs242_ = (0) - (d_1610_cf55_)
                rhs243_ = p2
                rhs244_ = p2
                rhs245_ = iife173_()

                lhs136_ = self
                lhs137_ = d_1634_v19_
                lhs138_ = default__.safeIndex(147, (d_1634_v19_).length(0))
                d_1610_cf55_ = rhs242_
                d_1610_cf55_ = rhs243_
                lhs136_.f12 = rhs244_
                lhs137_[lhs138_] = rhs245_
                d_1639_v24_: C1
                nw214_ = C1()
                nw214_.ctor__()
                d_1639_v24_ = nw214_
            d_1604_v1_ = d_1604_v1_
            d_1640_v25_: _dafny.Map
            d_1640_v25_ = _dafny.Map({(self).f2: self.f3})
            def iife174_():
                coll50_ = _dafny.Map()
                compr_50_: _dafny.Seq
                for compr_50_ in (_dafny.Map({(self).fm14(p0, p0, _dafny.CodePoint('w'), (self).f4, globalState): self.f12})).keys.Elements:
                    d_1641_v26_: _dafny.Seq = compr_50_
                    if (d_1641_v26_) in (_dafny.Map({(self).fm14(p0, p0, _dafny.CodePoint('w'), (self).f4, globalState): self.f12})):
                        coll50_[d_1641_v26_] = (D3_DC9(d_1610_cf55_, p0)).cf16
                return _dafny.Map(coll50_)
            d_1640_v25_ = iife174_()
            
        elif source31_.is_DC35:
            d_1642___mcc_h4_ = source31_.cf57
            d_1643___mcc_h5_ = source31_.cf58
            d_1644___mcc_h6_ = source31_.cf59
            d_1645___mcc_h7_ = source31_.cf60
            d_1646___mcc_h8_ = source31_.cf61
            d_1647_cf61_ = d_1646___mcc_h8_
            d_1648_cf60_ = d_1645___mcc_h7_
            d_1649_cf59_ = d_1644___mcc_h6_
            d_1650_cf58_ = d_1643___mcc_h5_
            d_1651_cf57_ = d_1642___mcc_h4_
            index243_ = default__.safeIndex(830, ((self).f5).length(0))
            ((self).f5)[index243_] = d_1649_cf59_
            d_1652_v27_: _dafny.MultiSet
            d_1652_v27_ = _dafny.MultiSet([self.f3])
            d_1653_v28_: _dafny.Map
            d_1653_v28_ = _dafny.Map({(0) - (p2): ((p1)[p2] if (p2) in (p1) else default__.fm0(globalState))})
            d_1654_v29_: _dafny.Map
            d_1654_v29_ = _dafny.Map({len(d_1653_v28_): d_1652_v27_})
            d_1655_v30_: _dafny.Seq
            d_1655_v30_ = _dafny.SeqWithoutIsStrInference([d_1652_v27_, d_1652_v27_])
            d_1656_v31_: _dafny.Array
            nw215_ = _dafny.Array(None, 13)
            nw215_[int(0)] = d_1652_v27_
            nw215_[int(1)] = d_1652_v27_
            nw215_[int(2)] = (d_1652_v27_).intersection(d_1652_v27_)
            nw215_[int(3)] = ((d_1654_v29_)[(self).f4] if ((self).f4) in (d_1654_v29_) else d_1652_v27_)
            nw215_[int(4)] = d_1652_v27_
            nw215_[int(5)] = d_1652_v27_
            nw215_[int(6)] = d_1652_v27_
            nw215_[int(7)] = d_1652_v27_
            nw215_[int(8)] = d_1652_v27_
            nw215_[int(9)] = d_1652_v27_
            nw215_[int(10)] = _dafny.MultiSet([d_1649_cf59_])
            nw215_[int(11)] = d_1652_v27_
            nw215_[int(12)] = (d_1652_v27_) | ((d_1655_v30_)[default__.safeIndex(398, len(d_1655_v30_))])
            d_1656_v31_ = nw215_
            index244_ = default__.safeIndex(376, (d_1656_v31_).length(0))
            (d_1656_v31_)[index244_] = d_1652_v27_
            d_1657_v32_: _dafny.Array
            def lambda64_(d_1658_p2_):
                def lambda65_(d_1659_i1_):
                    return default__.safeDivisionInt(d_1659_i1_, d_1658_p2_)

                return lambda65_

            init34_ = lambda64_(p2)
            nw216_ = _dafny.Array(None, 17)
            for i0_34_ in range(nw216_.length(0)):
                nw216_[i0_34_] = init34_(i0_34_)
            d_1657_v32_ = nw216_
            index245_ = default__.safeIndex(380, (d_1657_v32_).length(0))
            (d_1657_v32_)[index245_] = default__.fm0(globalState)
            d_1660_v33_: _dafny.Map
            d_1660_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ktq")): d_1648_cf60_})
            d_1661_v34_: str
            d_1661_v34_ = _dafny.CodePoint('f')
            index246_ = default__.safeIndex(830, ((self).f5).length(0))
            index247_ = default__.safeIndex(376, (d_1656_v31_).length(0))
            index248_ = default__.safeIndex(380, (d_1657_v32_).length(0))
            rhs246_ = (self).fm13(globalState)
            rhs247_ = not(d_1649_cf59_)
            rhs248_ = (d_1652_v27_).intersection(default__.fm5((d_1660_v33_).set((self).fm14(d_1649_cf59_, self.f3, d_1661_v34_, self.f12, globalState), (self).f11), globalState))
            rhs249_ = p2
            lhs139_ = (self).f5
            lhs140_ = default__.safeIndex(830, ((self).f5).length(0))
            lhs141_ = self
            lhs142_ = d_1656_v31_
            lhs143_ = default__.safeIndex(376, (d_1656_v31_).length(0))
            lhs144_ = d_1657_v32_
            lhs145_ = default__.safeIndex(380, (d_1657_v32_).length(0))
            lhs139_[lhs140_] = rhs246_
            lhs141_.f3 = rhs247_
            lhs142_[lhs143_] = rhs248_
            lhs144_[lhs145_] = rhs249_
            d_1662_v35_: _dafny.Array
            nw217_ = _dafny.Array(False, 21)
            d_1662_v35_ = nw217_
            d_1663_v37_: C3
            nw218_ = C3()
            def iife175_():
                coll51_ = _dafny.Map()
                compr_51_: _dafny.Seq
                for compr_51_ in (d_1660_v33_).keys.Elements:
                    d_1664_v36_: _dafny.Seq = compr_51_
                    if (d_1664_v36_) in (d_1660_v33_):
                        coll51_[d_1664_v36_] = p0
                return _dafny.Map(coll51_)
            nw218_.ctor__((self).f11, d_1662_v35_, not(((self).f5)[default__.safeIndex(830, ((self).f5).length(0))]), ((self).f1) | (iife175_()
            ), ((self).f2) + ((self).f2))
            d_1663_v37_ = nw218_
            d_1665_v38_: _dafny.Map
            d_1665_v38_ = _dafny.Map({d_1648_cf60_: p0})
            d_1666_v39_: _dafny.Seq
            d_1666_v39_ = _dafny.SeqWithoutIsStrInference([d_1665_v38_])
            d_1667_v40_: _dafny.Seq
            d_1667_v40_ = _dafny.SeqWithoutIsStrInference([d_1648_cf60_, len(d_1666_v39_), p2, d_1648_cf60_, (self).f4])
            pat_let_tv44_ = d_1667_v40_
            pat_let_tv45_ = d_1667_v40_
            def iife176_(_pat_let62_0):
                def iife177_(d_1668_dt__update__tmp_h1_):
                    def iife178_(_pat_let63_0):
                        def iife179_(d_1669_dt__update_hcf24_h1_):
                            return D5_DC15(d_1669_dt__update_hcf24_h1_, (d_1668_dt__update__tmp_h1_).cf25, (d_1668_dt__update__tmp_h1_).cf26)
                        return iife179_(_pat_let63_0)
                    return iife178_((pat_let_tv44_)[default__.safeIndex(511, len(pat_let_tv45_))])
                return iife177_(_pat_let62_0)
            source32_ = iife176_(D5_DC15(d_1651_cf57_, (d_1663_v37_).fm11(not((d_1663_v37_).fm13(globalState)), globalState), d_1662_v35_))
            if source32_.is_DC15:
                d_1670___mcc_h11_ = source32_.cf24
                d_1671___mcc_h12_ = source32_.cf25
                d_1672___mcc_h13_ = source32_.cf26
                d_1673_cf26_ = d_1672___mcc_h13_
                d_1674_cf25_ = d_1671___mcc_h12_
                d_1675_cf24_ = d_1670___mcc_h11_
                d_1676_v42_: _dafny.Map
                def iife180_():
                    coll52_ = _dafny.Set()
                    compr_52_: _dafny.MultiSet
                    for compr_52_ in (_dafny.Set({p1, p1, p1, p1, p1})).Elements:
                        d_1677_v41_: _dafny.MultiSet = compr_52_
                        if (d_1677_v41_) in (_dafny.Set({p1, p1, p1, p1, p1})):
                            coll52_ = coll52_.union(_dafny.Set([d_1677_v41_]))
                    return _dafny.Set(coll52_)
                d_1676_v42_ = _dafny.Map({default__.safeModuloInt((self).f11, d_1651_cf57_): iife180_()
                })
                def iife181_():
                    coll53_ = _dafny.Set()
                    compr_53_: _dafny.MultiSet
                    for compr_53_ in (_dafny.Set({p1, p1})).Elements:
                        d_1678_v43_: _dafny.MultiSet = compr_53_
                        if (d_1678_v43_) in (_dafny.Set({p1, p1})):
                            coll53_ = coll53_.union(_dafny.Set([d_1678_v43_]))
                    return _dafny.Set(coll53_)
                d_1676_v42_ = (d_1676_v42_).set((self).f11, (iife181_()
                ) | (_dafny.Set({_dafny.MultiSet(d_1667_v40_)})))
                index249_ = default__.safeIndex(830, ((self).f5).length(0))
                ((self).f5)[index249_] = self.f3
                d_1653_v28_ = d_1653_v28_
                d_1679_v44_: C0
                nw219_ = C0()
                nw219_.ctor__(d_1657_v32_)
                d_1679_v44_ = nw219_
            elif True:
                d_1680___mcc_h14_ = source32_.cf23
                d_1681_cf23_ = d_1680___mcc_h14_
                d_1682_v45_: C0
                nw220_ = C0()
                nw220_.ctor__(d_1657_v32_)
                d_1682_v45_ = nw220_
                index250_ = default__.safeIndex(830, ((self).f5).length(0))
                ((self).f5)[index250_] = p0
                index251_ = default__.safeIndex(830, ((self).f5).length(0))
                ((self).f5)[index251_] = p0
                (self).f3 = d_1649_cf59_
            index252_ = default__.safeIndex(380, (d_1657_v32_).length(0))
            (d_1657_v32_)[index252_] = p2
        elif source31_.is_DC33:
            d_1683___mcc_h9_ = source31_.cf52
            d_1684_cf52_ = d_1683___mcc_h9_
            rhs250_ = (self).f4
            rhs251_ = not(True)
            rhs252_ = (d_1684_cf52_).f13
            lhs146_ = self
            lhs147_ = self
            lhs148_ = self
            lhs146_.f12 = rhs250_
            lhs147_.f3 = rhs251_
            lhs148_.f12 = rhs252_
            (self).f12 = (default__.safeModuloInt((self).f4, (d_1684_cf52_).f13)) * (p2)
            index253_ = default__.safeIndex(837, ((self).f5).length(0))
            ((self).f5)[index253_] = p0
            d_1685_v46_: str
            d_1685_v46_ = _dafny.CodePoint('s')
            d_1686_v47_: _dafny.Map
            d_1686_v47_ = _dafny.Map({self.f12: self.f3})
            d_1687_v48_: _dafny.Seq
            d_1687_v48_ = _dafny.SeqWithoutIsStrInference([(d_1684_cf52_).f13])
            d_1688_v49_: _dafny.Seq
            d_1688_v49_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_1687_v48_), p1, _dafny.MultiSet([len((self).f2), (0) - ((self).f4)]), _dafny.MultiSet([(d_1684_cf52_).f13])])
            d_1689_v50_: _dafny.Map
            d_1689_v50_ = _dafny.Map({(d_1688_v49_)[default__.safeIndex(len((self).f2), len(d_1688_v49_))]: p0})
            d_1690_v51_: D4
            d_1690_v51_ = D4_DC13(self.f3, d_1687_v48_)
            d_1691_v52_: _dafny.Map
            d_1691_v52_ = _dafny.Map({False: self.f3})
            index254_ = default__.safeIndex(837, ((self).f5).length(0))
            rhs253_ = not((238) != ((len(((self).f2).set(default__.safeIndex(self.f12, len((self).f2)), d_1685_v46_))) * (len((d_1686_v47_).set(804, p0)))))
            rhs254_ = False
            rhs255_ = (0) - (((self).f4) - (self.f12))
            rhs256_ = self.f3
            rhs257_ = not(((d_1689_v50_)[_dafny.MultiSet((((d_1690_v51_).cf22).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f4])), len((d_1690_v51_).cf22)), (d_1684_cf52_).f13)) + (d_1687_v48_))] if (_dafny.MultiSet((((d_1690_v51_).cf22).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f4])), len((d_1690_v51_).cf22)), (d_1684_cf52_).f13)) + (d_1687_v48_))) in (d_1689_v50_) else ((d_1691_v52_).set(p0, p0)) != (d_1691_v52_)))
            lhs149_ = self
            lhs150_ = (self).f5
            lhs151_ = default__.safeIndex(837, ((self).f5).length(0))
            lhs152_ = self
            lhs153_ = self
            lhs154_ = self
            lhs149_.f3 = rhs253_
            lhs150_[lhs151_] = rhs254_
            lhs152_.f12 = rhs255_
            lhs153_.f3 = rhs256_
            lhs154_.f3 = rhs257_
            d_1692_v53_: _dafny.Array
            nw221_ = _dafny.Array(None, 16)
            nw221_[int(0)] = False
            nw221_[int(1)] = ((self).f5)[default__.safeIndex(837, ((self).f5).length(0))]
            nw221_[int(2)] = p0
            nw221_[int(3)] = self.f3
            nw221_[int(4)] = p0
            nw221_[int(5)] = p0
            nw221_[int(6)] = ((self).f5)[default__.safeIndex(837, ((self).f5).length(0))]
            nw221_[int(7)] = self.f3
            nw221_[int(8)] = p0
            nw221_[int(9)] = ((self).f5)[default__.safeIndex(837, ((self).f5).length(0))]
            nw221_[int(10)] = p0
            nw221_[int(11)] = ((self).f5)[default__.safeIndex(837, ((self).f5).length(0))]
            nw221_[int(12)] = p0
            nw221_[int(13)] = True
            nw221_[int(14)] = ((self).f5)[default__.safeIndex(837, ((self).f5).length(0))]
            nw221_[int(15)] = not(p0)
            d_1692_v53_ = nw221_
            d_1693_v54_: _dafny.Seq
            d_1693_v54_ = _dafny.SeqWithoutIsStrInference([d_1692_v53_])
            index255_ = default__.safeIndex(837, ((self).f5).length(0))
            ((self).f5)[index255_] = ((d_1693_v54_) + (d_1693_v54_)) == (((d_1693_v54_) + (d_1693_v54_)).set(default__.safeIndex(self.f12, len((d_1693_v54_) + (d_1693_v54_))), d_1692_v53_))
        elif True:
            d_1694___mcc_h10_ = source31_.cf62
            d_1695_cf62_ = d_1694___mcc_h10_
            if self.f3:
                (self).f12 = (self).fm21(self.f3, self.f12, globalState)
                (self).f12 = default__.fm0(globalState)
                index256_ = default__.safeIndex(822, ((self).f5).length(0))
                ((self).f5)[index256_] = self.f3
                d_1696_v55_: _dafny.Map
                d_1696_v55_ = _dafny.Map({(self).f11: self.f3})
                d_1697_v56_: _dafny.Set
                d_1697_v56_ = _dafny.Set({p0})
                index257_ = default__.safeIndex(822, ((self).f5).length(0))
                rhs258_ = (0) - (len((d_1696_v55_) | (d_1696_v55_)))
                rhs259_ = (self).fm21(self.f3, p2, globalState)
                rhs260_ = self.f3
                rhs261_ = (d_1697_v56_) == (d_1697_v56_)
                lhs155_ = self
                lhs156_ = self
                lhs157_ = (self).f5
                lhs158_ = default__.safeIndex(822, ((self).f5).length(0))
                lhs159_ = self
                lhs155_.f12 = rhs258_
                lhs156_.f12 = rhs259_
                lhs157_[lhs158_] = rhs260_
                lhs159_.f3 = rhs261_
                index258_ = default__.safeIndex(822, ((self).f5).length(0))
                ((self).f5)[index258_] = ((self).f5)[default__.safeIndex(822, ((self).f5).length(0))]
                (self).f12 = len(d_1696_v55_)
            elif True:
                d_1698_v57_: _dafny.Map
                d_1698_v57_ = _dafny.Map({(self).f5: self.f3})
                d_1699_v58_: bool
                d_1700_v59_: bool
                d_1701_v60_: _dafny.Seq
                d_1702_v61_: _dafny.Seq
                out37_: bool
                out38_: bool
                out39_: _dafny.Seq
                out40_: _dafny.Seq
                out37_, out38_, out39_, out40_ = default__.m0(d_1698_v57_, globalState)
                d_1699_v58_ = out37_
                d_1700_v59_ = out38_
                d_1701_v60_ = out39_
                d_1702_v61_ = out40_
                (self).f3 = d_1699_v58_
                d_1703_v62_: C4
                nw222_ = C4()
                nw222_.ctor__(d_1700_v59_, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dph")))
                d_1703_v62_ = nw222_
                d_1704_v63_: _dafny.Set
                d_1704_v63_ = _dafny.Set({p2})
                d_1705_v64_: _dafny.Seq
                d_1705_v64_ = _dafny.SeqWithoutIsStrInference([(self).f4, self.f12, (self).f11, (0) - ((self).f4), (len(d_1704_v63_)) - (len(_dafny.Map({self.f12: p0})))])
                (self).f12 = len(d_1705_v64_)
                d_1706_v65_: _dafny.Array
                def lambda66_(d_1707_v58_):
                    def lambda67_(d_1708_i2_):
                        return d_1707_v58_

                    return lambda67_

                init35_ = lambda66_(d_1699_v58_)
                nw223_ = _dafny.Array(None, 16)
                for i0_35_ in range(nw223_.length(0)):
                    nw223_[i0_35_] = init35_(i0_35_)
                d_1706_v65_ = nw223_
                d_1709_v66_: _dafny.Array
                nw224_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                d_1709_v66_ = nw224_
                d_1710_v67_: D14
                d_1710_v67_ = D14_DC37(d_1709_v66_)
                d_1711_v68_: _dafny.Array
                nw225_ = _dafny.Array(None, 26)
                nw225_[int(0)] = d_1709_v66_
                nw225_[int(1)] = d_1709_v66_
                nw225_[int(2)] = d_1709_v66_
                nw225_[int(3)] = d_1709_v66_
                nw225_[int(4)] = d_1709_v66_
                nw225_[int(5)] = d_1709_v66_
                nw225_[int(6)] = d_1709_v66_
                nw225_[int(7)] = d_1709_v66_
                nw225_[int(8)] = d_1709_v66_
                nw225_[int(9)] = d_1709_v66_
                nw225_[int(10)] = d_1709_v66_
                nw225_[int(11)] = d_1709_v66_
                nw225_[int(12)] = (d_1710_v67_).cf63
                nw225_[int(13)] = d_1709_v66_
                nw225_[int(14)] = d_1709_v66_
                nw225_[int(15)] = d_1709_v66_
                nw225_[int(16)] = d_1709_v66_
                nw225_[int(17)] = d_1709_v66_
                nw225_[int(18)] = d_1709_v66_
                nw225_[int(19)] = d_1709_v66_
                nw225_[int(20)] = d_1709_v66_
                nw225_[int(21)] = d_1709_v66_
                nw225_[int(22)] = d_1709_v66_
                nw225_[int(23)] = d_1709_v66_
                nw225_[int(24)] = d_1709_v66_
                nw225_[int(25)] = d_1709_v66_
                d_1711_v68_ = nw225_
                d_1712_v69_: _dafny.Array
                d_1712_v69_ = d_1711_v68_
                d_1713_v70_: str
                d_1713_v70_ = _dafny.CodePoint('j')
                rhs262_ = (_dafny.CodePoint('h')) not in (_dafny.SeqWithoutIsStrInference([d_1713_v70_, d_1713_v70_]))
                rhs263_ = (self).f5
                rhs264_ = d_1712_v69_
                d_1699_v58_ = rhs262_
                d_1706_v65_ = rhs263_
                d_1712_v69_ = rhs264_
            d_1714_v71_: _dafny.Set
            d_1714_v71_ = _dafny.Set({self.f3})
            d_1715_v72_: _dafny.Map
            d_1715_v72_ = _dafny.Map({(self).f4: d_1714_v71_})
            d_1716_v73_: _dafny.Array
            nw226_ = _dafny.Array(int(0), 23)
            d_1716_v73_ = nw226_
            d_1717_v74_: _dafny.Seq
            d_1717_v74_ = _dafny.SeqWithoutIsStrInference([d_1716_v73_])
            d_1718_v75_: _dafny.Array
            nw227_ = _dafny.Array(None, 26)
            nw227_[int(0)] = d_1714_v71_
            nw227_[int(1)] = (d_1714_v71_) - (d_1714_v71_)
            nw227_[int(2)] = (d_1714_v71_).intersection(_dafny.Set({self.f3}))
            nw227_[int(3)] = d_1714_v71_
            nw227_[int(4)] = d_1714_v71_
            nw227_[int(5)] = d_1714_v71_
            nw227_[int(6)] = d_1714_v71_
            nw227_[int(7)] = (d_1714_v71_).intersection(d_1714_v71_)
            nw227_[int(8)] = (d_1714_v71_) | (d_1714_v71_)
            nw227_[int(9)] = d_1714_v71_
            nw227_[int(10)] = d_1714_v71_
            nw227_[int(11)] = d_1714_v71_
            nw227_[int(12)] = _dafny.Set({self.f3, p0})
            nw227_[int(13)] = (d_1714_v71_) - (d_1714_v71_)
            nw227_[int(14)] = d_1714_v71_
            nw227_[int(15)] = (d_1714_v71_) | (d_1714_v71_)
            nw227_[int(16)] = d_1714_v71_
            nw227_[int(17)] = d_1714_v71_
            nw227_[int(18)] = d_1714_v71_
            nw227_[int(19)] = _dafny.Set({self.f3, self.f3})
            nw227_[int(20)] = (d_1714_v71_) | (d_1714_v71_)
            nw227_[int(21)] = d_1714_v71_
            nw227_[int(22)] = ((d_1715_v72_)[len(d_1717_v74_)] if (len(d_1717_v74_)) in (d_1715_v72_) else _dafny.Set({p0}))
            nw227_[int(23)] = d_1714_v71_
            nw227_[int(24)] = d_1714_v71_
            nw227_[int(25)] = d_1714_v71_
            d_1718_v75_ = nw227_
            index259_ = default__.safeIndex(793, (d_1718_v75_).length(0))
            (d_1718_v75_)[index259_] = _dafny.Set({p0})
            index260_ = default__.safeIndex(793, (d_1718_v75_).length(0))
            (d_1718_v75_)[index260_] = (d_1714_v71_).intersection((d_1714_v71_) - (_dafny.Set({p0})))
            d_1714_v71_ = d_1714_v71_
            d_1719_v76_: D3
            d_1719_v76_ = D3_DC8((d_1718_v75_)[default__.safeIndex(793, (d_1718_v75_).length(0))])
            d_1720_v77_: D3
            d_1720_v77_ = D3_DC11(d_1719_v76_)
            d_1721_v78_: _dafny.Seq
            d_1721_v78_ = _dafny.SeqWithoutIsStrInference([self.f3, not(p0), not(p0)])
            d_1722_v79_: _dafny.Map
            d_1722_v79_ = _dafny.Map({self.f3: d_1721_v78_})
            d_1720_v77_ = default__.fm40(len((self).f2), self.f3, len(((d_1722_v79_)[self.f3] if (self.f3) in (d_1722_v79_) else d_1721_v78_)), (0) - ((self.f12) + (self.f12)), globalState)
        if self.f3:
            d_1723_v80_: str
            d_1723_v80_ = _dafny.CodePoint('d')
            d_1724_v81_: _dafny.Set
            d_1724_v81_ = _dafny.Set({p0, self.f3, self.f3, p0, False})
            d_1725_v82_: D3
            d_1725_v82_ = D3_DC8(d_1724_v81_)
            d_1726_v83_: _dafny.Set
            d_1726_v83_ = _dafny.Set({D3_DC8(d_1724_v81_), d_1725_v82_})
            d_1727_v84_: _dafny.Seq
            d_1727_v84_ = _dafny.SeqWithoutIsStrInference([self.f3, p0])
            d_1723_v80_ = default__.fm38(not(not((d_1726_v83_).isdisjoint(d_1726_v83_))), (d_1727_v84_) < (d_1727_v84_), 769, globalState)
            d_1728_v85_: _dafny.Array
            nw228_ = _dafny.Array(int(0), 4)
            d_1728_v85_ = nw228_
            index261_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            (d_1728_v85_)[index261_] = (self).f4
            index262_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            (d_1728_v85_)[index262_] = self.f12
            index263_ = default__.safeIndex(705, ((self).f5).length(0))
            ((self).f5)[index263_] = p0
            index264_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            index265_ = default__.safeIndex(705, ((self).f5).length(0))
            index266_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            rhs265_ = p2
            rhs266_ = False
            rhs267_ = (self.f12) * (default__.safeModuloInt(807, default__.fm0(globalState)))
            lhs160_ = d_1728_v85_
            lhs161_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            lhs162_ = (self).f5
            lhs163_ = default__.safeIndex(705, ((self).f5).length(0))
            lhs164_ = d_1728_v85_
            lhs165_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            lhs160_[lhs161_] = rhs265_
            lhs162_[lhs163_] = rhs266_
            lhs164_[lhs165_] = rhs267_
            (self).f3 = ((self).f5)[default__.safeIndex(705, ((self).f5).length(0))]
            index267_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            rhs268_ = 240
            rhs269_ = d_1728_v85_
            lhs166_ = d_1728_v85_
            lhs167_ = default__.safeIndex(951, (d_1728_v85_).length(0))
            lhs166_[lhs167_] = rhs268_
            d_1728_v85_ = rhs269_
        elif True:
            d_1729_v86_: _dafny.Set
            d_1729_v86_ = _dafny.Set({p0})
            d_1730_v87_: D3
            d_1730_v87_ = D3_DC8(d_1729_v86_)
            d_1731_v88_: D3
            d_1731_v88_ = D3_DC11(D3_DC11(d_1730_v87_))
            source33_ = d_1731_v88_
            if source33_.is_DC9:
                d_1732___mcc_h15_ = source33_.cf15
                d_1733___mcc_h16_ = source33_.cf16
                d_1734_cf16_ = d_1733___mcc_h16_
                d_1735_cf15_ = d_1732___mcc_h15_
                d_1736_v89_: _dafny.Map
                d_1736_v89_ = _dafny.Map({self.f12: d_1735_cf15_})
                d_1737_v90_: _dafny.Seq
                d_1737_v90_ = _dafny.SeqWithoutIsStrInference([len((d_1736_v89_).set(self.f12, self.f12))])
                rhs270_ = ((len((self).f2)) * ((d_1737_v90_)[default__.safeIndex(self.f12, len(d_1737_v90_))])) == (default__.safeDivisionInt(p2, len((self).f2)))
                rhs271_ = default__.fm0(globalState)
                rhs272_ = default__.safeDivisionInt(len((d_1729_v86_) | (d_1729_v86_)), (self).f4)
                d_1734_cf16_ = rhs270_
                d_1735_cf15_ = rhs271_
                d_1735_cf15_ = rhs272_
                (self).f3 = False
                d_1738_v91_: _dafny.Array
                nw229_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_1738_v91_ = nw229_
                d_1739_v92_: D14
                d_1739_v92_ = D14_DC38((self).f2)
                index268_ = default__.safeIndex(377, (d_1738_v91_).length(0))
                (d_1738_v91_)[index268_] = (d_1739_v92_).cf64
                d_1740_v93_: _dafny.Seq
                d_1740_v93_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
                d_1741_v94_: str
                d_1741_v94_ = _dafny.CodePoint('n')
                index269_ = default__.safeIndex(377, (d_1738_v91_).length(0))
                (d_1738_v91_)[index269_] = ((self).f2) + (((self).f2).set(default__.safeIndex(len(d_1740_v93_), len((self).f2)), d_1741_v94_))
                index270_ = default__.safeIndex(766, ((self).f5).length(0))
                ((self).f5)[index270_] = self.f3
                d_1742_v95_: _dafny.Array
                def lambda68_(d_1743_i3_):
                    return default__.safeModuloInt(d_1743_i3_, (self).f11)

                init36_ = lambda68_
                nw230_ = _dafny.Array(None, 18)
                for i0_36_ in range(nw230_.length(0)):
                    nw230_[i0_36_] = init36_(i0_36_)
                d_1742_v95_ = nw230_
                index271_ = default__.safeIndex(889, (d_1742_v95_).length(0))
                (d_1742_v95_)[index271_] = p2
                index272_ = default__.safeIndex(766, ((self).f5).length(0))
                index273_ = default__.safeIndex(889, (d_1742_v95_).length(0))
                index274_ = default__.safeIndex(377, (d_1738_v91_).length(0))
                rhs273_ = False
                rhs274_ = not(True)
                rhs275_ = d_1735_cf15_
                rhs276_ = (self).f2
                lhs168_ = (self).f5
                lhs169_ = default__.safeIndex(766, ((self).f5).length(0))
                lhs170_ = d_1742_v95_
                lhs171_ = default__.safeIndex(889, (d_1742_v95_).length(0))
                lhs172_ = d_1738_v91_
                lhs173_ = default__.safeIndex(377, (d_1738_v91_).length(0))
                lhs168_[lhs169_] = rhs273_
                d_1734_cf16_ = rhs274_
                lhs170_[lhs171_] = rhs275_
                lhs172_[lhs173_] = rhs276_
            elif source33_.is_DC10:
                d_1744___mcc_h17_ = source33_.cf17
                d_1745___mcc_h18_ = source33_.cf18
                d_1746_cf18_ = d_1745___mcc_h18_
                d_1747_cf17_ = d_1744___mcc_h17_
                (self).f3 = self.f3
                d_1748_v96_: _dafny.Array
                nw231_ = _dafny.Array(int(0), 18)
                d_1748_v96_ = nw231_
                d_1749_v97_: _dafny.Set
                d_1749_v97_ = _dafny.Set({self.f12})
                rhs277_ = d_1748_v96_
                rhs278_ = d_1747_cf17_
                rhs279_ = d_1749_v97_
                d_1748_v96_ = rhs277_
                d_1747_cf17_ = rhs278_
                d_1749_v97_ = rhs279_
                index275_ = default__.safeIndex(682, ((self).f5).length(0))
                ((self).f5)[index275_] = p0
                d_1750_v98_: _dafny.MultiSet
                d_1750_v98_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1751_i4_ in range(default__.abs(34))])])
                d_1752_v99_: _dafny.Seq
                d_1752_v99_ = _dafny.SeqWithoutIsStrInference([(self).f2])
                index276_ = default__.safeIndex(682, ((self).f5).length(0))
                rhs280_ = (self).fm12(default__.fm0(globalState), d_1746_cf18_, d_1746_cf18_, p0, globalState)
                rhs281_ = (d_1746_cf18_ if self.f3 else (self).fm21(self.f3, (self).f4, globalState))
                rhs282_ = p0
                rhs283_ = ((self).f11) < (((d_1750_v98_) | (_dafny.MultiSet(d_1752_v99_))).cardinality)
                lhs174_ = (self).f5
                lhs175_ = default__.safeIndex(682, ((self).f5).length(0))
                lhs176_ = self
                d_1747_cf17_ = rhs280_
                d_1746_cf18_ = rhs281_
                lhs174_[lhs175_] = rhs282_
                lhs176_.f3 = rhs283_
                d_1753_v100_: _dafny.Array
                nw232_ = _dafny.Array(False, 24)
                d_1753_v100_ = nw232_
                d_1754_v101_: _dafny.Set
                d_1754_v101_ = _dafny.Set({d_1753_v100_})
                d_1755_v102_: _dafny.Map
                d_1755_v102_ = _dafny.Map({d_1747_cf17_: ((self).f5)[default__.safeIndex(682, ((self).f5).length(0))]})
                d_1756_v103_: _dafny.Set
                d_1756_v103_ = _dafny.Set({d_1755_v102_})
                d_1757_v104_: _dafny.Map
                d_1757_v104_ = _dafny.Map({_dafny.Map({d_1746_cf18_: p0}): d_1747_cf17_})
                d_1758_v106_: _dafny.Seq
                def iife182_():
                    coll54_ = _dafny.Set()
                    compr_54_: _dafny.Map
                    for compr_54_ in (d_1757_v104_).keys.Elements:
                        d_1759_v105_: _dafny.Map = compr_54_
                        if (d_1759_v105_) in (d_1757_v104_):
                            coll54_ = coll54_.union(_dafny.Set([d_1759_v105_]))
                    return _dafny.Set(coll54_)
                d_1758_v106_ = _dafny.SeqWithoutIsStrInference([d_1756_v103_, iife182_()
                , d_1756_v103_])
                d_1760_v107_: _dafny.MultiSet
                d_1760_v107_ = _dafny.MultiSet([self.f3])
                d_1761_v108_: _dafny.Map
                d_1761_v108_ = _dafny.Map({d_1755_v102_: d_1760_v107_})
                d_1762_v110_: _dafny.Array
                nw233_ = _dafny.Array(None, 10)
                nw233_[int(0)] = (d_1756_v103_).intersection(d_1756_v103_)
                nw233_[int(1)] = _dafny.Set({d_1755_v102_})
                nw233_[int(2)] = (d_1758_v106_)[default__.safeIndex(d_1746_cf18_, len(d_1758_v106_))]
                nw233_[int(3)] = d_1756_v103_
                nw233_[int(4)] = d_1756_v103_
                nw233_[int(5)] = d_1756_v103_
                nw233_[int(6)] = d_1756_v103_
                def iife183_():
                    coll55_ = _dafny.Set()
                    compr_55_: _dafny.Map
                    for compr_55_ in (d_1761_v108_).keys.Elements:
                        d_1763_v109_: _dafny.Map = compr_55_
                        if (d_1763_v109_) in (d_1761_v108_):
                            coll55_ = coll55_.union(_dafny.Set([d_1763_v109_]))
                    return _dafny.Set(coll55_)
                nw233_[int(7)] = iife183_()
                
                nw233_[int(8)] = d_1756_v103_
                nw233_[int(9)] = d_1756_v103_
                d_1762_v110_ = nw233_
                index277_ = default__.safeIndex(100, (d_1762_v110_).length(0))
                (d_1762_v110_)[index277_] = _dafny.Set({_dafny.Map({d_1746_cf18_: p0}), d_1755_v102_, d_1755_v102_, d_1755_v102_})
                index278_ = default__.safeIndex(100, (d_1762_v110_).length(0))
                rhs284_ = d_1754_v101_
                rhs285_ = (d_1756_v103_) | (_dafny.Set({d_1755_v102_, _dafny.Map({d_1746_cf18_: True}), d_1755_v102_, d_1755_v102_}))
                lhs177_ = d_1762_v110_
                lhs178_ = default__.safeIndex(100, (d_1762_v110_).length(0))
                d_1754_v101_ = rhs284_
                lhs177_[lhs178_] = rhs285_
            elif source33_.is_DC8:
                d_1764___mcc_h19_ = source33_.cf14
                d_1765_cf14_ = d_1764___mcc_h19_
                (self).f3 = False
                d_1766_v111_: _dafny.Map
                d_1766_v111_ = _dafny.Map({p2: (self).f11})
                d_1767_v112_: _dafny.MultiSet
                d_1767_v112_ = _dafny.MultiSet([(self).f4])
                index279_ = default__.safeIndex(869, ((self).f5).length(0))
                ((self).f5)[index279_] = self.f3
                d_1768_v113_: D6
                d_1768_v113_ = D6_DC17(p0)
                d_1769_v114_: _dafny.Set
                d_1769_v114_ = _dafny.Set({258, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1770_i5_ in range(default__.abs(-838))]))})
                index280_ = default__.safeIndex(869, ((self).f5).length(0))
                rhs286_ = ((d_1766_v111_) | (d_1766_v111_)) | (_dafny.Map({p2: 749}))
                rhs287_ = (d_1768_v113_).cf28
                rhs288_ = ((0) - ((self).f11)) == (default__.fm0(globalState))
                rhs289_ = (p1 if (d_1769_v114_).ispropersubset(d_1769_v114_) else (d_1767_v112_) | (p1))
                rhs290_ = self.f3
                lhs179_ = self
                lhs180_ = self
                lhs181_ = (self).f5
                lhs182_ = default__.safeIndex(869, ((self).f5).length(0))
                d_1766_v111_ = rhs286_
                lhs179_.f3 = rhs287_
                lhs180_.f3 = rhs288_
                d_1767_v112_ = rhs289_
                lhs181_[lhs182_] = rhs290_
                (self).f12 = (self).f4
                d_1771_v115_: _dafny.Seq
                d_1771_v115_ = _dafny.SeqWithoutIsStrInference([(self).f11])
                d_1772_v116_: str
                d_1772_v116_ = _dafny.CodePoint('l')
                d_1773_v117_: _dafny.Seq
                d_1773_v117_ = _dafny.SeqWithoutIsStrInference([len(d_1771_v115_), (self).f11, (self).f11, len((self).fm14(self.f3, self.f3, d_1772_v116_, (self).f4, globalState))])
                d_1774_v118_: _dafny.Map
                d_1774_v118_ = _dafny.Map({(self).f2: d_1773_v117_})
                (self).f12 = len(((d_1774_v118_)[(self).f2] if ((self).f2) in (d_1774_v118_) else d_1771_v115_))
            elif True:
                d_1775___mcc_h20_ = source33_.cf19
                d_1776_cf19_ = d_1775___mcc_h20_
                d_1777_v119_: _dafny.Array
                nw234_ = _dafny.Array(None, 5)
                nw234_[int(0)] = (self).f11
                nw234_[int(1)] = self.f12
                nw234_[int(2)] = (self).f4
                nw234_[int(3)] = self.f12
                nw234_[int(4)] = 317
                d_1777_v119_ = nw234_
                d_1778_v120_: C0
                nw235_ = C0()
                nw235_.ctor__(d_1777_v119_)
                d_1778_v120_ = nw235_
                d_1778_v120_ = d_1778_v120_
                index281_ = default__.safeIndex(730, ((self).f5).length(0))
                ((self).f5)[index281_] = (p1).ispropersubset(p1)
                index282_ = default__.safeIndex(730, ((self).f5).length(0))
                ((self).f5)[index282_] = not((p2) != ((self).f11))
                (self).f12 = (self).f4
                d_1779_v121_: str
                d_1779_v121_ = _dafny.CodePoint('o')
                d_1779_v121_ = _dafny.CodePoint('g')
            d_1780_v122_: _dafny.Array
            def lambda69_(d_1781_i6_):
                return (d_1781_i6_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([133]))]))))

            init37_ = lambda69_
            nw236_ = _dafny.Array(None, 17)
            for i0_37_ in range(nw236_.length(0)):
                nw236_[i0_37_] = init37_(i0_37_)
            d_1780_v122_ = nw236_
            d_1780_v122_ = d_1780_v122_
            d_1782_v123_: _dafny.MultiSet
            d_1782_v123_ = _dafny.MultiSet([False])
            (self).f3 = (d_1782_v123_).ispropersubset(d_1782_v123_)
            d_1783_v124_: C2
            nw237_ = C2()
            nw237_.ctor__((self).f4, not(True), (self).f1, (self).f2)
            d_1783_v124_ = nw237_
            d_1784_v125_: _dafny.Seq
            d_1784_v125_ = _dafny.SeqWithoutIsStrInference([d_1783_v124_, d_1783_v124_, d_1783_v124_])
            d_1785_v126_: _dafny.Map
            d_1785_v126_ = _dafny.Map({d_1784_v125_: p1})
            rhs291_ = d_1785_v126_
            rhs292_ = self.f3
            lhs183_ = self
            d_1785_v126_ = rhs291_
            lhs183_.f3 = rhs292_
            d_1786_v127_: _dafny.Array
            nw238_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_1786_v127_ = nw238_
            d_1787_v128_: D0
            d_1787_v128_ = D0_DC0((d_1783_v124_).f13)
            d_1788_v129_: _dafny.Array
            nw239_ = _dafny.Array(None, 5)
            nw239_[int(0)] = d_1787_v128_
            nw239_[int(1)] = default__.fm50(self.f12, p2, self.f3, globalState)
            nw239_[int(2)] = D0_DC0((d_1783_v124_).f13)
            nw239_[int(3)] = d_1787_v128_
            nw239_[int(4)] = d_1787_v128_
            d_1788_v129_ = nw239_
            index283_ = default__.safeIndex(319, (d_1786_v127_).length(0))
            (d_1786_v127_)[index283_] = d_1788_v129_
            index284_ = default__.safeIndex(319, (d_1786_v127_).length(0))
            (d_1786_v127_)[index284_] = d_1788_v129_
        d_1789_i7_: int
        d_1789_i7_ = 0
        with _dafny.label("11"):
            while self.f3:
                with _dafny.c_label("11"):
                    if (d_1789_i7_) >= (100):
                        raise _dafny.Break("11")
                    d_1789_i7_ = (d_1789_i7_) + (1)
                    if self.f3:
                        d_1790_v130_: bool
                        out41_: bool
                        out41_ = (self).m8(p0, globalState)
                        d_1790_v130_ = out41_
                        (self).f12 = (self).f4
                        (self).f12 = (self.f12) - ((len((self).f2) if p0 else p2))
                        d_1791_v131_: _dafny.Map
                        d_1791_v131_ = _dafny.Map({True: (self).f4})
                        d_1791_v131_ = (d_1791_v131_).set((505) > (p2), (self.f12) - ((self).f11))
                        d_1792_v132_: _dafny.Array
                        def lambda70_(d_1793_i8_):
                            return default__.safeDivisionInt(d_1793_i8_, self.f12)

                        init38_ = lambda70_
                        nw240_ = _dafny.Array(None, 7)
                        for i0_38_ in range(nw240_.length(0)):
                            nw240_[i0_38_] = init38_(i0_38_)
                        d_1792_v132_ = nw240_
                        d_1792_v132_ = d_1792_v132_
                    elif True:
                        d_1794_v133_: _dafny.Seq
                        d_1794_v133_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
                        rhs293_ = default__.fm7(True, self.f3, globalState)
                        rhs294_ = False
                        lhs184_ = self
                        d_1794_v133_ = rhs293_
                        lhs184_.f3 = rhs294_
                        (self).f12 = (0) - (p2)
                        d_1795_v134_: _dafny.Map
                        d_1795_v134_ = _dafny.Map({(self).f11: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1796_i9_ in range(default__.abs(631))])})
                        d_1797_v135_: C3
                        nw241_ = C3()
                        nw241_.ctor__(p2, (self).f5, self.f3, (self).f1, ((d_1795_v134_)[self.f12] if (self.f12) in (d_1795_v134_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kvxnj"))))
                        d_1797_v135_ = nw241_
                        d_1798_v136_: _dafny.Array
                        nw242_ = _dafny.Array(int(0), 14)
                        d_1798_v136_ = nw242_
                        index285_ = default__.safeIndex(238, (d_1798_v136_).length(0))
                        (d_1798_v136_)[index285_] = p2
                        index286_ = default__.safeIndex(238, (d_1798_v136_).length(0))
                        (d_1798_v136_)[index286_] = default__.fm0(globalState)
                        d_1799_v137_: _dafny.Set
                        d_1799_v137_ = _dafny.Set({p2})
                        (self).f3 = (self.f12) not in (d_1799_v137_)
                    d_1800_v138_: _dafny.Map
                    d_1800_v138_ = _dafny.Map({p0: (self).f4})
                    rhs295_ = (d_1800_v138_) | (d_1800_v138_)
                    rhs296_ = 76
                    lhs185_ = self
                    d_1800_v138_ = rhs295_
                    lhs185_.f12 = rhs296_
                    d_1801_v139_: D1
                    d_1801_v139_ = D1_DC2((self).f2)
                    source34_ = d_1801_v139_
                    if source34_.is_DC3:
                        d_1802___mcc_h21_ = source34_.cf3
                        d_1803___mcc_h22_ = source34_.cf4
                        d_1804___mcc_h23_ = source34_.cf5
                        d_1805_cf5_ = d_1804___mcc_h23_
                        d_1806_cf4_ = d_1803___mcc_h22_
                        d_1807_cf3_ = d_1802___mcc_h21_
                        d_1808_v140_: _dafny.Set
                        d_1808_v140_ = _dafny.Set({default__.safeDivisionInt((self).f4, (_dafny.MultiSet([(self).f2])).cardinality)})
                        d_1808_v140_ = _dafny.Set({(self).f4})
                        d_1809_v141_: _dafny.Seq
                        d_1809_v141_ = _dafny.SeqWithoutIsStrInference([d_1805_cf5_, (self).f11, (self).f4])
                        d_1809_v141_ = d_1809_v141_
                        d_1810_v142_: _dafny.Array
                        nw243_ = _dafny.Array(int(0), 27)
                        d_1810_v142_ = nw243_
                        index287_ = default__.safeIndex(622, (d_1810_v142_).length(0))
                        (d_1810_v142_)[index287_] = (self).f4
                        index288_ = default__.safeIndex(622, (d_1810_v142_).length(0))
                        (d_1810_v142_)[index288_] = (0) - (d_1806_cf4_)
                        d_1810_v142_ = d_1810_v142_
                    elif source34_.is_DC4:
                        d_1811___mcc_h24_ = source34_.cf6
                        d_1812___mcc_h25_ = source34_.cf7
                        d_1813_cf7_ = d_1812___mcc_h25_
                        d_1814_cf6_ = d_1811___mcc_h24_
                        d_1814_cf6_ = self.f3
                        d_1815_v143_: _dafny.Map
                        d_1815_v143_ = _dafny.Map({(self).f5: (self).fm21(p0, (self).f11, globalState)})
                        d_1815_v143_ = (d_1815_v143_).set((self).f5, ((0) - (len(_dafny.Set({d_1814_cf6_}))) if d_1814_cf6_ else p2))
                        (self).f12 = ((self).f11) - (p2)
                        d_1816_v144_: str
                        d_1816_v144_ = _dafny.CodePoint('f')
                        d_1817_v145_: _dafny.Array
                        nw244_ = _dafny.Array(None, 15)
                        nw244_[int(0)] = (self).f2
                        nw244_[int(1)] = ((self).f2 if d_1813_cf7_ else (self).f2)
                        nw244_[int(2)] = _dafny.SeqWithoutIsStrInference([d_1816_v144_, d_1816_v144_])
                        nw244_[int(3)] = (self).f2
                        nw244_[int(4)] = ((self).f2).set(default__.safeIndex((self).f11, len((self).f2)), default__.fm24(p0, p0, self.f12, globalState))
                        nw244_[int(5)] = (self).f2
                        nw244_[int(6)] = ((self).f2 if default__.fm3((self).f11, globalState) else (self).f2)
                        nw244_[int(7)] = ((self).f2) + ((self).f2)
                        nw244_[int(8)] = (self).f2
                        nw244_[int(9)] = (self).f2
                        nw244_[int(10)] = (self).f2
                        nw244_[int(11)] = ((self).f2) + ((self).f2)
                        nw244_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_1816_v144_])) + ((self).f2)
                        nw244_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1816_v144_ for d_1818_i10_ in range(default__.abs(579))])
                        nw244_[int(14)] = (((self).f2) + (_dafny.SeqWithoutIsStrInference([d_1816_v144_]))).set(default__.safeIndex(p2, len(((self).f2) + (_dafny.SeqWithoutIsStrInference([d_1816_v144_])))), d_1816_v144_)
                        d_1817_v145_ = nw244_
                        index289_ = default__.safeIndex(877, (d_1817_v145_).length(0))
                        (d_1817_v145_)[index289_] = (self).f2
                        index290_ = default__.safeIndex(877, (d_1817_v145_).length(0))
                        (d_1817_v145_)[index290_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1819_i11_ in range(default__.abs(501))])
                    elif True:
                        d_1820___mcc_h26_ = source34_.cf2
                        d_1821_cf2_ = d_1820___mcc_h26_
                        d_1822_v146_: _dafny.Set
                        d_1822_v146_ = _dafny.Set({(self).f2})
                        (self).f12 = (self).fm21(((self).f2) in (d_1822_v146_), self.f12, globalState)
                        d_1823_v148_: _dafny.Seq
                        d_1823_v148_ = _dafny.SeqWithoutIsStrInference([d_1821_cf2_])
                        d_1824_v149_: str
                        d_1824_v149_ = _dafny.CodePoint('p')
                        d_1825_v150_: C3
                        nw245_ = C3()
                        def iife184_():
                            coll56_ = _dafny.Map()
                            compr_56_: _dafny.Seq
                            for compr_56_ in (d_1823_v148_).Elements:
                                d_1826_v147_: _dafny.Seq = compr_56_
                                if (d_1826_v147_) in (d_1823_v148_):
                                    coll56_[d_1826_v147_] = False
                            return _dafny.Map(coll56_)
                        nw245_.ctor__(self.f12, (self).f5, self.f3, iife184_()
                        , (d_1821_cf2_).set(default__.safeIndex((self).f11, len(d_1821_cf2_)), d_1824_v149_))
                        d_1825_v150_ = nw245_
                        d_1827_v151_: _dafny.Map
                        d_1827_v151_ = _dafny.Map({self.f3: p0})
                        d_1828_v152_: _dafny.Map
                        d_1828_v152_ = _dafny.Map({d_1824_v149_: d_1827_v151_})
                        d_1829_v153_: _dafny.Map
                        d_1829_v153_ = _dafny.Map({self.f12: d_1827_v151_})
                        d_1830_v154_: _dafny.Set
                        d_1830_v154_ = _dafny.Set({p1})
                        d_1831_v155_: _dafny.Seq
                        d_1831_v155_ = _dafny.SeqWithoutIsStrInference([d_1827_v151_])
                        d_1832_v156_: _dafny.Array
                        nw246_ = _dafny.Array(None, 22)
                        nw246_[int(0)] = _dafny.Map({(D6_DC18((self).f11, p2, (d_1825_v150_).fm26(-455, p0, p0, globalState), d_1824_v149_, self.f12)).cf31: self.f3})
                        nw246_[int(1)] = (d_1827_v151_) | (d_1827_v151_)
                        nw246_[int(2)] = (_dafny.Map({p0: True})) | (d_1827_v151_)
                        nw246_[int(3)] = d_1827_v151_
                        nw246_[int(4)] = default__.fm9(globalState)
                        nw246_[int(5)] = _dafny.Map({p0: self.f3})
                        nw246_[int(6)] = d_1827_v151_
                        nw246_[int(7)] = (_dafny.Map({(self).fm11(p0, globalState): True})) | (d_1827_v151_)
                        nw246_[int(8)] = (d_1827_v151_ if not(self.f3) else ((d_1828_v152_)[d_1824_v149_] if (d_1824_v149_) in (d_1828_v152_) else d_1827_v151_))
                        nw246_[int(9)] = d_1827_v151_
                        nw246_[int(10)] = ((d_1829_v153_)[len((self).fm14(self.f3, False, d_1824_v149_, len(d_1830_v154_), globalState))] if (len((self).fm14(self.f3, False, d_1824_v149_, len(d_1830_v154_), globalState))) in (d_1829_v153_) else (d_1831_v155_)[default__.safeIndex(p2, len(d_1831_v155_))])
                        nw246_[int(11)] = d_1827_v151_
                        nw246_[int(12)] = (d_1827_v151_) | (d_1827_v151_)
                        nw246_[int(13)] = (d_1827_v151_) | (d_1827_v151_)
                        nw246_[int(14)] = (d_1831_v155_)[default__.safeIndex((0) - ((self).f11), len(d_1831_v155_))]
                        nw246_[int(15)] = d_1827_v151_
                        nw246_[int(16)] = d_1827_v151_
                        nw246_[int(17)] = (d_1831_v155_)[default__.safeIndex((self).f11, len(d_1831_v155_))]
                        nw246_[int(18)] = d_1827_v151_
                        nw246_[int(19)] = d_1827_v151_
                        nw246_[int(20)] = (d_1827_v151_) | (d_1827_v151_)
                        nw246_[int(21)] = _dafny.Map({self.f3: p0})
                        d_1832_v156_ = nw246_
                        d_1832_v156_ = d_1832_v156_
                        d_1833_v157_: _dafny.Set
                        d_1833_v157_ = _dafny.Set({self.f12})
                        d_1834_v158_: _dafny.Set
                        d_1834_v158_ = d_1833_v157_
                        d_1835_v159_: _dafny.Map
                        d_1835_v159_ = _dafny.Map({(self).f11: (d_1833_v157_).isdisjoint((d_1834_v158_))})
                        d_1836_v160_: _dafny.Map
                        d_1836_v160_ = _dafny.Map({(self).f5: d_1821_cf2_})
                        d_1837_v161_: _dafny.Seq
                        d_1837_v161_ = _dafny.SeqWithoutIsStrInference([self.f12])
                        d_1838_v162_: _dafny.Map
                        d_1838_v162_ = _dafny.Map({d_1824_v149_: (self).f11})
                        d_1835_v159_ = (d_1835_v159_).set(len(d_1836_v160_), (len(d_1837_v161_)) >= ((0) - (((d_1838_v162_)[d_1824_v149_] if (d_1824_v149_) in (d_1838_v162_) else (0) - ((self).f4)))))
                    (self).f3 = p0
                    pass
            pass
        d_1839_v163_: _dafny.Seq
        d_1839_v163_ = _dafny.SeqWithoutIsStrInference([(self).f11])
        d_1840_v164_: _dafny.Map
        d_1840_v164_ = _dafny.Map({self.f12: (self).fm21(p0, (self).f11, globalState)})
        d_1841_v165_: _dafny.Seq
        d_1841_v165_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p0, p0})) for d_1842_i13_ in range(default__.abs(884))])])
        d_1843_v166_: _dafny.Map
        d_1843_v166_ = _dafny.Map({p2: (d_1841_v165_)[default__.safeIndex((self).f4, len(d_1841_v165_))]})
        d_1844_v167_: _dafny.Array
        nw247_ = _dafny.Array(None, 26)
        nw247_[int(0)] = _dafny.MultiSet([self.f12, (self).f4])
        nw247_[int(1)] = p1
        nw247_[int(2)] = p1
        nw247_[int(3)] = (p1) - (default__.fm8(p2, (self).f2, globalState))
        nw247_[int(4)] = _dafny.MultiSet(d_1839_v163_)
        nw247_[int(5)] = p1
        nw247_[int(6)] = p1
        nw247_[int(7)] = p1
        nw247_[int(8)] = p1
        nw247_[int(9)] = _dafny.MultiSet([(self).f11, (self).fm12(self.f12, p2, (self).f11, not(True), globalState), (self).f11, (self).f4, (self).f11])
        nw247_[int(10)] = (p1) - (_dafny.MultiSet([len(_dafny.Map({(self).f5: (0) - ((self).f4)})), (self).fm12(self.f12, p2, (self).f11, self.f3, globalState), 588]))
        nw247_[int(11)] = (p1) | (p1)
        nw247_[int(12)] = (p1).set(self.f12, default__.abs((self).f11))
        nw247_[int(13)] = (_dafny.MultiSet(d_1839_v163_)).intersection(p1)
        nw247_[int(14)] = p1
        nw247_[int(15)] = p1
        nw247_[int(16)] = (_dafny.MultiSet([len(d_1840_v164_)])) - (p1)
        nw247_[int(17)] = (p1).intersection(p1)
        nw247_[int(18)] = p1
        nw247_[int(19)] = ((_dafny.MultiSet(((d_1843_v166_)[self.f12] if (self.f12) in (d_1843_v166_) else d_1839_v163_))).set((self).f4, default__.abs(333))) - (p1)
        nw247_[int(20)] = (p1) | ((p1).set((0) - ((self).f11), default__.abs(self.f12)))
        nw247_[int(21)] = (p1) - (p1)
        nw247_[int(22)] = p1
        nw247_[int(23)] = p1
        nw247_[int(24)] = p1
        nw247_[int(25)] = p1
        d_1844_v167_ = nw247_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1844_v167_).length(0)):
            d_1845_i12_: int = guard_loop_2_
            if (True) and (((0) <= (d_1845_i12_)) and ((d_1845_i12_) < ((d_1844_v167_).length(0)))):
                (d_1844_v167_)[(d_1845_i12_)] = _dafny.MultiSet([default__.safeDivisionInt(p2, (self).f4), (self).f4, (self).f11, (self).f11])
        (self).f3 = ((self).f11) < (len((self).f2))

    def m8(self, p0, globalState):
        r0: bool = False
        d_1846_v0_: D0
        d_1846_v0_ = D0_DC0((self).f4)
        d_1847_v1_: _dafny.Map
        d_1847_v1_ = _dafny.Map({(d_1846_v0_).cf0: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xx")))})
        d_1848_v2_: _dafny.MultiSet
        d_1848_v2_ = _dafny.MultiSet([d_1847_v1_, d_1847_v1_, d_1847_v1_])
        def iife185_():
            coll57_ = _dafny.Map()
            compr_57_: int
            for compr_57_ in _dafny.IntegerRange(478, 344):
                d_1849_v3_: int = compr_57_
                if ((478) <= (d_1849_v3_)) and ((d_1849_v3_) < (344)):
                    coll57_[default__.safeDivisionInt(d_1849_v3_, 116)] = (self).f4
            return _dafny.Map(coll57_)
        d_1848_v2_ = ((d_1848_v2_) - (d_1848_v2_)) | ((d_1848_v2_).set(iife185_()
        , default__.abs(601)))
        d_1850_v4_: _dafny.Seq
        d_1850_v4_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
        d_1851_v5_: _dafny.Array
        def lambda71_(d_1852_i1_):
            return (d_1852_i1_) - ((_dafny.MultiSet([(self).f11])).cardinality)

        init39_ = lambda71_
        nw248_ = _dafny.Array(None, 23)
        for i0_39_ in range(nw248_.length(0)):
            nw248_[i0_39_] = init39_(i0_39_)
        d_1851_v5_ = nw248_
        d_1853_v6_: _dafny.Set
        d_1853_v6_ = _dafny.Set({d_1851_v5_})
        d_1854_v7_: _dafny.Array
        nw249_ = _dafny.Array(None, 22)
        nw249_[int(0)] = (self).fm12(self.f12, -859, (self).f4, p0, globalState)
        nw249_[int(1)] = 163
        nw249_[int(2)] = self.f12
        nw249_[int(3)] = 727
        nw249_[int(4)] = default__.fm0(globalState)
        nw249_[int(5)] = self.f12
        nw249_[int(6)] = 674
        nw249_[int(7)] = (self).f4
        nw249_[int(8)] = (self).f4
        nw249_[int(9)] = (self).f4
        nw249_[int(10)] = self.f12
        nw249_[int(11)] = (0) - (len(d_1850_v4_))
        nw249_[int(12)] = self.f12
        nw249_[int(13)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1855_i0_ in range(default__.abs(129))]))
        nw249_[int(14)] = self.f12
        nw249_[int(15)] = (default__.fm8(len(d_1853_v6_), (self).f2, globalState)).cardinality
        nw249_[int(16)] = -538
        nw249_[int(17)] = (self).f4
        nw249_[int(18)] = len((self).f2)
        nw249_[int(19)] = self.f12
        nw249_[int(20)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f12]))).cardinality
        nw249_[int(21)] = (self).f11
        d_1854_v7_ = nw249_
        d_1856_v8_: C0
        nw250_ = C0()
        nw250_.ctor__(d_1854_v7_)
        d_1856_v8_ = nw250_
        d_1857_v9_: _dafny.MultiSet
        d_1857_v9_ = _dafny.MultiSet([(self).f11, self.f12])
        d_1858_v10_: _dafny.Seq
        d_1858_v10_ = _dafny.SeqWithoutIsStrInference([self.f3])
        d_1859_v11_: D14
        d_1859_v11_ = D14_DC39(326, (((d_1857_v9_)[(self).f11] if ((self).f11) in (d_1857_v9_) else len(d_1858_v10_))) + ((self).f11), (self).f4)
        d_1859_v11_ = d_1859_v11_
        d_1860_v12_: _dafny.Seq
        d_1860_v12_ = _dafny.SeqWithoutIsStrInference([len(d_1858_v10_)])
        if (d_1860_v12_) != (d_1860_v12_):
            d_1861_v13_: str
            d_1861_v13_ = _dafny.CodePoint('c')
            d_1862_v14_: C4
            nw251_ = C4()
            nw251_.ctor__(self.f3, _dafny.Map({(self).fm14(p0, self.f3, d_1861_v13_, len((self).f2), globalState): self.f3}), (self).f2)
            d_1862_v14_ = nw251_
            d_1863_v15_: _dafny.Map
            d_1863_v15_ = _dafny.Map({p0: (self).f5})
            d_1863_v15_ = (d_1863_v15_).set(True, (self).f5)
            d_1864_v16_: _dafny.Seq
            d_1864_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugutx"))
            d_1864_v16_ = (self).fm14((p0) and (self.f3), p0, d_1861_v13_, (self).f4, globalState)
            d_1865_v17_: C0
            nw252_ = C0()
            nw252_.ctor__((d_1856_v8_).f14)
            d_1865_v17_ = nw252_
            if p0:
                d_1866_v18_: _dafny.Map
                d_1866_v18_ = _dafny.Map({p0: False})
                d_1867_v20_: _dafny.Seq
                d_1867_v20_ = _dafny.SeqWithoutIsStrInference([d_1864_v16_, default__.fm2(p0, globalState), (self).f2, d_1864_v16_, d_1864_v16_])
                d_1868_v21_: C4
                nw253_ = C4()
                def iife186_():
                    coll58_ = _dafny.Map()
                    compr_58_: _dafny.Seq
                    for compr_58_ in (d_1867_v20_).Elements:
                        d_1869_v19_: _dafny.Seq = compr_58_
                        if (d_1869_v19_) in (d_1867_v20_):
                            coll58_[d_1869_v19_] = True
                    return _dafny.Map(coll58_)
                nw253_.ctor__((len(_dafny.SeqWithoutIsStrInference([len(d_1866_v18_), len(d_1860_v12_)]))) not in (d_1860_v12_), ((self).f1) | (iife186_()
                ), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykpsh")))
                d_1868_v21_ = nw253_
                d_1870_v22_: _dafny.Map
                d_1870_v22_ = _dafny.Map({(self).f5: ((self).f2) + ((self).f2)})
                d_1864_v16_ = ((d_1870_v22_)[(self).f5] if ((self).f5) in (d_1870_v22_) else d_1864_v16_)
                d_1871_v23_: _dafny.MultiSet
                d_1871_v23_ = _dafny.MultiSet([d_1864_v16_])
                d_1872_v24_: _dafny.Map
                d_1872_v24_ = _dafny.Map({not(p0): (self).f4})
                d_1873_v25_: _dafny.Seq
                d_1873_v25_ = _dafny.SeqWithoutIsStrInference([d_1872_v24_, d_1872_v24_])
                d_1874_v26_: _dafny.Map
                d_1874_v26_ = _dafny.Map({(d_1871_v23_).intersection(d_1871_v23_): d_1873_v25_})
                d_1874_v26_ = (d_1874_v26_).set((d_1871_v23_).set(d_1864_v16_, default__.abs(self.f12)), d_1873_v25_)
                (self).f12 = ((d_1868_v21_).fm12((default__.fm10(self.f3, self.f3, self.f3, globalState)).cf12, self.f12, (d_1857_v9_).cardinality, self.f3, globalState)) - (len(d_1864_v16_))
                (self).f12 = default__.safeModuloInt(len((self).f2), (D3_DC9(643, default__.fm3((self).fm21(p0, (self).f11, globalState), globalState))).cf15)
            elif True:
                r0 = (not(not(self.f3))) and (not (p0) or (p0))
                d_1875_v27_: _dafny.Set
                d_1875_v27_ = _dafny.Set({(_dafny.MultiSet([self.f3])).cardinality})
                d_1876_v28_: _dafny.Map
                d_1876_v28_ = _dafny.Map({d_1875_v27_: (912 if p0 else (self).f11)})
                d_1877_v29_: _dafny.Map
                d_1877_v29_ = _dafny.Map({d_1865_v17_: d_1875_v27_})
                d_1878_v30_: _dafny.Map
                d_1878_v30_ = _dafny.Map({((d_1847_v1_)[(self).f11] if ((self).f11) in (d_1847_v1_) else self.f12): not(self.f3)})
                d_1879_v31_: _dafny.MultiSet
                d_1879_v31_ = _dafny.MultiSet([d_1878_v30_])
                d_1880_v32_: D3
                d_1880_v32_ = D3_DC9((self).f11, p0)
                d_1876_v28_ = (d_1876_v28_).set((d_1875_v27_) | (((d_1877_v29_)[d_1865_v17_] if (d_1865_v17_) in (d_1877_v29_) else _dafny.Set({834, (d_1879_v31_).cardinality}))), (d_1880_v32_).cf15)
                d_1881_v33_: D6
                d_1881_v33_ = D6_DC16(d_1878_v30_)
                d_1882_v34_: D6
                d_1882_v34_ = D6_DC20(d_1881_v33_)
                d_1883_v35_: D6
                d_1883_v35_ = D6_DC20((d_1882_v34_).cf35)
                d_1884_v36_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1884_v36_ = nw254_
                d_1885_v37_: _dafny.Array
                nw255_ = _dafny.Array(None, 18)
                nw255_[int(0)] = d_1851_v5_
                nw255_[int(1)] = (d_1856_v8_).f14
                nw255_[int(2)] = (d_1865_v17_).f14
                nw255_[int(3)] = (d_1856_v8_).f14
                nw255_[int(4)] = (d_1856_v8_).f14
                nw255_[int(5)] = (d_1856_v8_).f14
                nw255_[int(6)] = (d_1865_v17_).f14
                nw255_[int(7)] = (d_1865_v17_).f14
                nw255_[int(8)] = (d_1856_v8_).f14
                nw255_[int(9)] = (d_1856_v8_).f14
                nw255_[int(10)] = (d_1856_v8_).f14
                nw255_[int(11)] = d_1851_v5_
                nw255_[int(12)] = d_1851_v5_
                nw255_[int(13)] = d_1851_v5_
                nw255_[int(14)] = d_1851_v5_
                nw255_[int(15)] = (d_1865_v17_).f14
                nw255_[int(16)] = d_1854_v7_
                nw255_[int(17)] = (d_1865_v17_).f14
                d_1885_v37_ = nw255_
                index291_ = default__.safeIndex(777, (d_1884_v36_).length(0))
                (d_1884_v36_)[index291_] = d_1885_v37_
                index292_ = default__.safeIndex(777, (d_1884_v36_).length(0))
                rhs297_ = d_1883_v35_
                rhs298_ = d_1885_v37_
                rhs299_ = (self).f2
                lhs186_ = d_1884_v36_
                lhs187_ = default__.safeIndex(777, (d_1884_v36_).length(0))
                d_1882_v34_ = rhs297_
                lhs186_[lhs187_] = rhs298_
                d_1864_v16_ = rhs299_
                d_1886_v38_: T1
                nw256_ = C4()
                nw256_.ctor__(p0, (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")): not(self.f3)})) | ((self).f1), d_1864_v16_)
                d_1886_v38_ = nw256_
                d_1887_v39_: _dafny.Map
                d_1887_v39_ = _dafny.Map({d_1864_v16_: (self).f11})
                rhs300_ = ((self.f12) <= (self.f12)) or (p0)
                rhs301_ = d_1886_v38_
                rhs302_ = ((self.f12) >= ((self).f11) if (d_1886_v38_.f3) and (not(d_1886_v38_.f3)) else (default__.fm5(default__.fm51((self).f11, self.f12, globalState), globalState)).ispropersubset(default__.fm5(d_1887_v39_, globalState)))
                lhs188_ = self
                lhs189_ = self
                lhs188_.f3 = rhs300_
                d_1886_v38_ = rhs301_
                lhs189_.f3 = rhs302_
                d_1854_v7_ = ((d_1856_v8_).f14 if p0 else (d_1865_v17_).f14)
        elif True:
            d_1888_v40_: T0
            nw257_ = C5()
            nw257_.ctor__(p0, (self).f1, (self).f2)
            d_1888_v40_ = nw257_
            d_1889_v41_: _dafny.Seq
            d_1889_v41_ = _dafny.SeqWithoutIsStrInference([d_1888_v40_, d_1888_v40_, d_1888_v40_])
            d_1890_v42_: _dafny.Array
            def lambda72_(d_1891_i2_):
                return not(self.f3)

            init40_ = lambda72_
            nw258_ = _dafny.Array(None, 19)
            for i0_40_ in range(nw258_.length(0)):
                nw258_[i0_40_] = init40_(i0_40_)
            d_1890_v42_ = nw258_
            rhs303_ = d_1889_v41_
            rhs304_ = (self).f5
            d_1889_v41_ = rhs303_
            d_1890_v42_ = rhs304_
            if not (p0) or (p0):
                d_1892_v43_: _dafny.MultiSet
                d_1892_v43_ = _dafny.MultiSet([not(self.f3)])
                d_1893_v44_: C2
                nw259_ = C2()
                nw259_.ctor__(((d_1892_v43_)[self.f3] if (self.f3) in (d_1892_v43_) else self.f12), p0, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wppe")): self.f3}), default__.fm2(self.f3, globalState))
                d_1893_v44_ = nw259_
                d_1892_v43_ = d_1892_v43_
                d_1894_v45_: _dafny.Set
                d_1894_v45_ = _dafny.Set({self.f12})
                d_1894_v45_ = (_dafny.Set({(self).f4, default__.fm0(globalState), (d_1893_v44_).f13, 318, (self).f4}) if self.f3 else d_1894_v45_)
                r0 = self.f3
                d_1895_v46_: _dafny.Seq
                d_1895_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                d_1895_v46_ = ((d_1888_v40_).f2) + (((self).f2) + ((d_1888_v40_).f2))
            elif True:
                d_1896_v47_: _dafny.Array
                def lambda73_(d_1897_p0_):
                    def lambda74_(d_1898_i3_):
                        return (D1_DC3(d_1897_p0_, self.f12, 757) if d_1897_p0_ else D1_DC3(self.f3, (self).f11, (self).f4))

                    return lambda74_

                init41_ = lambda73_(p0)
                nw260_ = _dafny.Array(None, 4)
                for i0_41_ in range(nw260_.length(0)):
                    nw260_[i0_41_] = init41_(i0_41_)
                d_1896_v47_ = nw260_
                d_1899_v48_: D1
                d_1899_v48_ = D1_DC3(self.f3, (self).f11, (self).f4)
                index293_ = default__.safeIndex(878, (d_1896_v47_).length(0))
                (d_1896_v47_)[index293_] = d_1899_v48_
                d_1900_v49_: C3
                nw261_ = C3()
                nw261_.ctor__((self).f11, (self).f5, self.f3, (d_1888_v40_).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctsrc")))
                d_1900_v49_ = nw261_
                d_1901_v50_: _dafny.MultiSet
                d_1901_v50_ = _dafny.MultiSet([d_1900_v49_, d_1900_v49_])
                d_1902_v51_: _dafny.Map
                d_1902_v51_ = _dafny.Map({p0: (d_1901_v50_).cardinality})
                d_1903_v52_: _dafny.Set
                d_1903_v52_ = _dafny.Set({(self).f4, (0) - ((self).f11), len(d_1902_v51_), (self).f11, (self).f11})
                index294_ = default__.safeIndex(878, (d_1896_v47_).length(0))
                rhs305_ = not((self).fm13(globalState))
                rhs306_ = not(p0)
                rhs307_ = ((self).f11) + (len(d_1903_v52_))
                rhs308_ = d_1899_v48_
                rhs309_ = ((d_1903_v52_) | (_dafny.Set({(self).f11, self.f12}))).ispropersubset(_dafny.Set({(self).f4, (self).f4, len(d_1903_v52_)}))
                lhs190_ = self
                lhs191_ = self
                lhs192_ = d_1896_v47_
                lhs193_ = default__.safeIndex(878, (d_1896_v47_).length(0))
                lhs194_ = self
                r0 = rhs305_
                lhs190_.f3 = rhs306_
                lhs191_.f12 = rhs307_
                lhs192_[lhs193_] = rhs308_
                lhs194_.f3 = rhs309_
                d_1904_v53_: C2
                nw262_ = C2()
                nw262_.ctor__((self).f4, p0, (d_1888_v40_).f1, (d_1888_v40_).f2)
                d_1904_v53_ = nw262_
                d_1905_v54_: _dafny.Map
                d_1905_v54_ = _dafny.Map({d_1904_v53_: d_1890_v42_})
                d_1906_v55_: _dafny.Map
                d_1906_v55_ = _dafny.Map({((self).f11) + (len(d_1905_v54_)): (d_1888_v40_).f2})
                d_1906_v55_ = d_1906_v55_
                (self).f3 = self.f3
                r0 = self.f3
                (self).f12 = ((d_1904_v53_).f13) + ((d_1904_v53_).f13)
            d_1907_v56_: _dafny.Set
            d_1907_v56_ = _dafny.Set({909})
            d_1908_v57_: _dafny.Set
            d_1908_v57_ = _dafny.Set({d_1907_v56_, d_1907_v56_})
            d_1908_v57_ = d_1908_v57_
            d_1909_v58_: _dafny.Seq
            d_1909_v58_ = _dafny.SeqWithoutIsStrInference([d_1854_v7_])
            d_1851_v5_ = (d_1909_v58_)[default__.safeIndex(self.f12, len(d_1909_v58_))]
            d_1910_v59_: C5
            nw263_ = C5()
            nw263_.ctor__(self.f3, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncifpw")))
            d_1910_v59_ = nw263_
        d_1911_v60_: _dafny.Array
        nw264_ = _dafny.Array(None, 2)
        nw264_[int(0)] = p0
        nw264_[int(1)] = p0
        d_1911_v60_ = nw264_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1911_v60_).length(0)):
            d_1912_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_1912_i4_)) and ((d_1912_i4_) < ((d_1911_v60_).length(0)))):
                (d_1911_v60_)[(d_1912_i4_)] = not((self.f3) or ((D10_DC29(p0, p0)).cf48))
        d_1913_v61_: _dafny.Set
        d_1913_v61_ = _dafny.Set({self.f3})
        d_1914_v62_: _dafny.Map
        d_1914_v62_ = _dafny.Map({(self).f4: d_1913_v61_})
        d_1915_v63_: D3
        d_1915_v63_ = D3_DC8(((d_1914_v62_)[(self).f4] if ((self).f4) in (d_1914_v62_) else d_1913_v61_))
        pat_let_tv46_ = p0
        pat_let_tv47_ = p0
        pat_let_tv48_ = p0
        def lambda75_(source35_):
            if source35_.is_DC9:
                d_1916___mcc_h0_ = source35_.cf15
                d_1917___mcc_h1_ = source35_.cf16
                d_1918_cf16_ = d_1917___mcc_h1_
                d_1919_cf15_ = d_1916___mcc_h0_
                return ((self).f11) <= ((self).f4)
            elif source35_.is_DC10:
                d_1920___mcc_h2_ = source35_.cf17
                d_1921___mcc_h3_ = source35_.cf18
                d_1922_cf18_ = d_1921___mcc_h3_
                d_1923_cf17_ = d_1920___mcc_h2_
                return (d_1923_cf17_) > (len(_dafny.Map({pat_let_tv46_: (self).f11})))
            elif source35_.is_DC8:
                d_1924___mcc_h4_ = source35_.cf14
                d_1925_cf14_ = d_1924___mcc_h4_
                return pat_let_tv47_
            elif True:
                d_1926___mcc_h5_ = source35_.cf19
                d_1927_cf19_ = d_1926___mcc_h5_
                return pat_let_tv48_

        rhs310_ = lambda75_(d_1915_v63_)
        rhs311_ = default__.safeModuloInt((self).f11, len((self).f2))
        lhs195_ = self
        r0 = rhs310_
        lhs195_.f12 = rhs311_
        r0 = p0
        return r0

    @property
    def f11(self):
        return self._f11

class C7(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f4, f5, f3, f1, f2):
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return ((_dafny.MultiSet([len((self).f2), (self).f4, (self).f4])) | (_dafny.MultiSet([(self).f4]))).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f4, (0) - ((0) - (len(_dafny.Map({_dafny.Set({(self).f4, (self).f4}): self.f3}))))])))

    def fm14(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvsoenjvi"))) + ((self).f2)

    def fm12(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm11(self, p0, globalState):
        def iife187_():
            coll59_ = _dafny.Set()
            compr_59_: D1
            for compr_59_ in (_dafny.SeqWithoutIsStrInference([D1_DC4(self.f3, self.f3) for d_1928_i0_ in range(default__.abs(732))])).Elements:
                d_1929_v0_: D1 = compr_59_
                if (d_1929_v0_) in (_dafny.SeqWithoutIsStrInference([D1_DC4(self.f3, self.f3) for d_1928_i0_ in range(default__.abs(732))])):
                    coll59_ = coll59_.union(_dafny.Set([d_1929_v0_]))
            return _dafny.Set(coll59_)
        return (iife187_()
        ).issubset(_dafny.Set({D1_DC4(self.f3, self.f3), D1_DC4(True, True)}))

    def fm18(self, p0, p1, globalState):
        if self.f3:
            if self.f3:
                return D4_DC12(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f4, (self).f4])))
            elif True:
                return D4_DC12(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f4) for d_1930_i0_ in range(default__.abs(-10))]))]))
        elif True:
            return D4_DC12(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(_dafny.MultiSet([True, self.f3])).cardinality: False})])), (self).f4]))

    def fm19(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        r2 = (0) - (((0) - (default__.safeDivisionInt(p0, p1))) + (default__.fm0(globalState)))
        d_1931_v0_: str
        d_1931_v0_ = _dafny.CodePoint('g')
        d_1932_v1_: _dafny.Set
        d_1932_v1_ = _dafny.Set({p1})
        d_1933_v2_: _dafny.Seq
        d_1933_v2_ = _dafny.SeqWithoutIsStrInference([len(d_1932_v1_)])
        d_1934_v3_: _dafny.Seq
        d_1934_v3_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
        d_1935_v4_: D3
        d_1935_v4_ = D3_DC9(333, self.f3)
        d_1936_v5_: _dafny.Map
        d_1936_v5_ = _dafny.Map({d_1931_v0_: not(self.f3)})
        d_1937_v6_: _dafny.MultiSet
        d_1937_v6_ = _dafny.MultiSet([self.f3, self.f3, self.f3])
        d_1938_v7_: _dafny.Array
        nw265_ = _dafny.Array(None, 26)
        nw265_[int(0)] = (self).f4
        nw265_[int(1)] = p1
        nw265_[int(2)] = (0) - (len((self).f2))
        nw265_[int(3)] = (D2_DC6(d_1931_v0_, self.f3, False, p0)).cf12
        nw265_[int(4)] = ((0) - (p0)) + (p1)
        nw265_[int(5)] = default__.safeDivisionInt(-932, len(d_1933_v2_))
        nw265_[int(6)] = ((self).f4) * (p1)
        nw265_[int(7)] = len((self).f2)
        nw265_[int(8)] = (self).f4
        nw265_[int(9)] = 923
        nw265_[int(10)] = default__.safeDivisionInt(p1, len(d_1933_v2_))
        nw265_[int(11)] = (0) - (len((d_1932_v1_) - (d_1932_v1_)))
        nw265_[int(12)] = (self).f4
        nw265_[int(13)] = p0
        nw265_[int(14)] = (p0 if (d_1934_v3_)[default__.safeIndex(len(d_1934_v3_), len(d_1934_v3_))] else (self).f4)
        nw265_[int(15)] = len(d_1933_v2_)
        nw265_[int(16)] = (d_1933_v2_)[default__.safeIndex((self).fm19((self).f4, (0) - (p0), D3_DC11(d_1935_v4_), d_1933_v2_, globalState), len(d_1933_v2_))]
        nw265_[int(17)] = p1
        nw265_[int(18)] = (p1 if self.f3 else (0) - ((self).f4))
        nw265_[int(19)] = len(d_1936_v5_)
        nw265_[int(20)] = (self).f4
        nw265_[int(21)] = default__.safeDivisionInt(88, (0) - ((d_1937_v6_).cardinality))
        nw265_[int(22)] = (self).f4
        nw265_[int(23)] = p1
        nw265_[int(24)] = (0) - (default__.fm0(globalState))
        nw265_[int(25)] = (p1) - (p0)
        d_1938_v7_ = nw265_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1938_v7_).length(0)):
            d_1939_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1939_i0_)) and ((d_1939_i0_) < ((d_1938_v7_).length(0)))):
                (d_1938_v7_)[(d_1939_i0_)] = (d_1939_i0_) + (((_dafny.MultiSet([(d_1933_v2_)[default__.safeIndex(len((self).f2), len(d_1933_v2_))]])) | (_dafny.MultiSet([p1, p1, len((self).f2), p1, p0]))).cardinality)
        hi7_ = p1
        for d_1940_i1_ in range(default__.safeModuloInt(p1, p0), hi7_):
            (self).f3 = ((self).f4) > (default__.safeDivisionInt(227, d_1940_i1_))
            d_1941_v8_: _dafny.Seq
            d_1941_v8_ = _dafny.SeqWithoutIsStrInference([default__.fm7(self.f3, self.f3, globalState), d_1934_v3_])
            d_1934_v3_ = ((d_1941_v8_)[default__.safeIndex(len(((self).f2).set(default__.safeIndex((self).f4, len((self).f2)), d_1931_v0_)), len(d_1941_v8_))]) + (d_1934_v3_)
            d_1942_v9_: _dafny.Array
            nw266_ = _dafny.Array(False, 17)
            d_1942_v9_ = nw266_
            rhs312_ = _dafny.SeqWithoutIsStrInference([(True if True else self.f3)])
            rhs313_ = (d_1942_v9_ if self.f3 else (self).f5)
            rhs314_ = len((self).f2)
            rhs315_ = self.f3
            lhs196_ = self
            d_1934_v3_ = rhs312_
            d_1942_v9_ = rhs313_
            r2 = rhs314_
            lhs196_.f3 = rhs315_
            if (self.f3) or (self.f3):
                (self).f3 = self.f3
                index295_ = default__.safeIndex(25, (d_1938_v7_).length(0))
                (d_1938_v7_)[index295_] = (p0 if self.f3 else (self).f4)
                index296_ = default__.safeIndex(25, (d_1938_v7_).length(0))
                rhs316_ = (26 if ((self).f4) > (len(d_1933_v2_)) else p1)
                rhs317_ = (p1) + (len((_dafny.Set({self.f3})) - (_dafny.Set({self.f3}))))
                lhs197_ = d_1938_v7_
                lhs198_ = default__.safeIndex(25, (d_1938_v7_).length(0))
                lhs197_[lhs198_] = rhs316_
                r1 = rhs317_
                d_1943_v10_: _dafny.Seq
                d_1943_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phdcd"))
                d_1943_v10_ = d_1943_v10_
                d_1944_v11_: _dafny.Map
                d_1944_v11_ = _dafny.Map({self.f3: d_1933_v2_})
                r2 = ((0) - (default__.safeDivisionInt(len(d_1944_v11_), p1))) * ((p1) * ((d_1938_v7_)[default__.safeIndex(25, (d_1938_v7_).length(0))]))
                d_1945_v12_: _dafny.Seq
                d_1945_v12_ = _dafny.SeqWithoutIsStrInference([d_1935_v4_])
                d_1946_v13_: D3
                d_1946_v13_ = D3_DC11((d_1945_v12_)[default__.safeIndex(p1, len(d_1945_v12_))])
                (self).f3 = (self.f3) and ((default__.fm0(globalState)) < ((self).fm19(p0, (d_1938_v7_)[default__.safeIndex(25, (d_1938_v7_).length(0))], d_1946_v13_, _dafny.SeqWithoutIsStrInference([(d_1938_v7_)[default__.safeIndex(25, (d_1938_v7_).length(0))] for d_1947_i2_ in range(default__.abs(-165))]), globalState)))
            elif True:
                index297_ = default__.safeIndex(21, (d_1942_v9_).length(0))
                (d_1942_v9_)[index297_] = self.f3
                d_1948_v14_: _dafny.Seq
                d_1948_v14_ = _dafny.SeqWithoutIsStrInference([d_1933_v2_, d_1933_v2_, ((_dafny.SeqWithoutIsStrInference([len((self).f2)])) + (_dafny.SeqWithoutIsStrInference([d_1940_i1_]))).set(default__.safeIndex((0) - (d_1940_i1_), len((_dafny.SeqWithoutIsStrInference([len((self).f2)])) + (_dafny.SeqWithoutIsStrInference([d_1940_i1_])))), (self).fm19(d_1940_i1_, d_1940_i1_, D3_DC11(D3_DC11(d_1935_v4_)), d_1933_v2_, globalState)), _dafny.SeqWithoutIsStrInference([p0, p0, p1]), d_1933_v2_])
                index298_ = default__.safeIndex(21, (d_1942_v9_).length(0))
                rhs318_ = self.f3
                rhs319_ = d_1938_v7_
                rhs320_ = self.f3
                rhs321_ = (self).fm13(globalState)
                rhs322_ = len((d_1948_v14_)[default__.safeIndex(d_1940_i1_, len(d_1948_v14_))])
                lhs199_ = d_1942_v9_
                lhs200_ = default__.safeIndex(21, (d_1942_v9_).length(0))
                lhs201_ = self
                lhs202_ = self
                lhs199_[lhs200_] = rhs318_
                d_1938_v7_ = rhs319_
                lhs201_.f3 = rhs320_
                lhs202_.f3 = rhs321_
                r2 = rhs322_
                r1 = p1
                index299_ = default__.safeIndex(104, (d_1938_v7_).length(0))
                (d_1938_v7_)[index299_] = d_1940_i1_
                index300_ = default__.safeIndex(104, (d_1938_v7_).length(0))
                index301_ = default__.safeIndex(21, (d_1942_v9_).length(0))
                rhs323_ = ((self).f4) - (p0)
                rhs324_ = default__.safeModuloInt((d_1940_i1_) * (p1), p0)
                rhs325_ = (d_1934_v3_) < (d_1934_v3_)
                lhs203_ = d_1938_v7_
                lhs204_ = default__.safeIndex(104, (d_1938_v7_).length(0))
                lhs205_ = d_1942_v9_
                lhs206_ = default__.safeIndex(21, (d_1942_v9_).length(0))
                r1 = rhs323_
                lhs203_[lhs204_] = rhs324_
                lhs205_[lhs206_] = rhs325_
                d_1949_v15_: _dafny.Array
                nw267_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                d_1949_v15_ = nw267_
                d_1949_v15_ = (d_1949_v15_ if (d_1942_v9_)[default__.safeIndex(21, (d_1942_v9_).length(0))] else d_1949_v15_)
                d_1950_v16_: _dafny.Map
                d_1950_v16_ = _dafny.Map({p1: (d_1942_v9_)[default__.safeIndex(21, (d_1942_v9_).length(0))]})
                d_1951_v17_: _dafny.Map
                d_1951_v17_ = _dafny.Map({(((d_1950_v16_)[(d_1933_v2_)[default__.safeIndex(p1, len(d_1933_v2_))]] if ((d_1933_v2_)[default__.safeIndex(p1, len(d_1933_v2_))]) in (d_1950_v16_) else (d_1942_v9_)[default__.safeIndex(21, (d_1942_v9_).length(0))])) or (self.f3): (self).f2})
                r2 = len(d_1951_v17_)
        d_1952_v18_: D5
        d_1952_v18_ = D5_DC15((self).f4, not(self.f3), (self).f5)
        d_1953_v19_: _dafny.Map
        d_1953_v19_ = _dafny.Map({p1: (self).f5})
        d_1954_v20_: D3
        d_1954_v20_ = D3_DC9(len((self).f2), (self).fm13(globalState))
        d_1955_v21_: _dafny.Map
        d_1955_v21_ = _dafny.Map({(_dafny.MultiSet([len(d_1953_v19_), (0) - (p1), (d_1954_v20_).cf15, (d_1937_v6_).cardinality])).cardinality: (0) - (p0)})
        d_1956_v22_: _dafny.Map
        d_1956_v22_ = _dafny.Map({(d_1952_v18_).cf25: d_1955_v21_})
        d_1956_v22_ = (d_1956_v22_).set(self.f3, d_1955_v21_)
        hi8_ = 0
        for d_1957_i3_ in range(374, hi8_):
            d_1958_v23_: D3
            d_1958_v23_ = D3_DC10((0) - (p0), (0) - (p1))
            pat_let_tv49_ = p1
            def iife188_(_pat_let64_0):
                def iife189_(d_1959_dt__update__tmp_h0_):
                    def iife190_(_pat_let65_0):
                        def iife191_(d_1960_dt__update_hcf18_h0_):
                            return D3_DC10((d_1959_dt__update__tmp_h0_).cf17, d_1960_dt__update_hcf18_h0_)
                        return iife191_(_pat_let65_0)
                    return iife190_(pat_let_tv49_)
                return iife189_(_pat_let64_0)
            source36_ = iife188_(d_1958_v23_)
            if source36_.is_DC9:
                d_1961___mcc_h0_ = source36_.cf15
                d_1962___mcc_h1_ = source36_.cf16
                d_1963_cf16_ = d_1962___mcc_h1_
                d_1964_cf15_ = d_1961___mcc_h0_
                d_1934_v3_ = d_1934_v3_
                d_1965_v24_: _dafny.Seq
                d_1965_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcmpvem"))
                d_1965_v24_ = (self).f2
                r2 = (p0) - ((self).f4)
                r2 = (self).f4
            elif source36_.is_DC10:
                d_1966___mcc_h2_ = source36_.cf17
                d_1967___mcc_h3_ = source36_.cf18
                d_1968_cf18_ = d_1967___mcc_h3_
                d_1969_cf17_ = d_1966___mcc_h2_
                d_1938_v7_ = d_1938_v7_
                d_1970_v25_: _dafny.Map
                d_1970_v25_ = _dafny.Map({d_1938_v7_: self.f3})
                d_1970_v25_ = (d_1970_v25_).set(d_1938_v7_, self.f3)
                (self).f3 = self.f3
                d_1953_v19_ = (d_1953_v19_).set(d_1968_cf18_, (self).f5)
            elif source36_.is_DC8:
                d_1971___mcc_h4_ = source36_.cf14
                d_1972_cf14_ = d_1971___mcc_h4_
                d_1931_v0_ = d_1931_v0_
                d_1938_v7_ = d_1938_v7_
                d_1973_v26_: _dafny.Map
                d_1973_v26_ = _dafny.Map({d_1931_v0_: d_1957_i3_})
                rhs326_ = (0) - (((d_1957_i3_) - ((self).f4)) + (default__.safeDivisionInt(d_1957_i3_, p0)))
                rhs327_ = ((d_1937_v6_).isdisjoint((d_1937_v6_).set(False, default__.abs(len(_dafny.Set({d_1957_i3_, ((d_1973_v26_)[_dafny.CodePoint('i')] if (_dafny.CodePoint('i')) in (d_1973_v26_) else d_1957_i3_), d_1957_i3_})))))) and (self.f3)
                lhs207_ = self
                r2 = rhs326_
                lhs207_.f3 = rhs327_
                d_1974_v27_: _dafny.Map
                d_1974_v27_ = _dafny.Map({(self).f5: not(True)})
                d_1975_v28_: bool
                d_1976_v29_: bool
                d_1977_v30_: _dafny.Seq
                d_1978_v31_: _dafny.Seq
                out42_: bool
                out43_: bool
                out44_: _dafny.Seq
                out45_: _dafny.Seq
                out42_, out43_, out44_, out45_ = default__.m0((d_1974_v27_) | (_dafny.Map({(self).f5: (d_1934_v3_)[default__.safeIndex((d_1937_v6_).cardinality, len(d_1934_v3_))]})), globalState)
                d_1975_v28_ = out42_
                d_1976_v29_ = out43_
                d_1977_v30_ = out44_
                d_1978_v31_ = out45_
            elif True:
                d_1979___mcc_h5_ = source36_.cf19
                d_1980_cf19_ = d_1979___mcc_h5_
                d_1981_v32_: _dafny.Array
                nw268_ = _dafny.Array(None, 12)
                nw268_[int(0)] = (d_1933_v2_ if self.f3 else d_1933_v2_)
                nw268_[int(1)] = _dafny.SeqWithoutIsStrInference([p0, (0) - (p1), (0) - (len(d_1955_v21_))])
                nw268_[int(2)] = d_1933_v2_
                nw268_[int(3)] = d_1933_v2_
                nw268_[int(4)] = d_1933_v2_
                nw268_[int(5)] = default__.fm20(self.f3, self.f3, self.f3, d_1934_v3_, globalState)
                nw268_[int(6)] = d_1933_v2_
                nw268_[int(7)] = d_1933_v2_
                nw268_[int(8)] = (d_1933_v2_) + (d_1933_v2_)
                nw268_[int(9)] = (d_1933_v2_) + (d_1933_v2_)
                nw268_[int(10)] = d_1933_v2_
                nw268_[int(11)] = d_1933_v2_
                d_1981_v32_ = nw268_
                nw269_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1981_v32_ = nw269_
                d_1982_v33_: _dafny.Array
                def lambda76_(d_1983_p0_):
                    def lambda77_(d_1984_i4_):
                        return (_dafny.Map({self.f3: len(_dafny.Map({self.f3: self.f3}))})) | (_dafny.Map({self.f3: d_1983_p0_}))

                    return lambda77_

                init42_ = lambda76_(p0)
                nw270_ = _dafny.Array(None, 26)
                for i0_42_ in range(nw270_.length(0)):
                    nw270_[i0_42_] = init42_(i0_42_)
                d_1982_v33_ = nw270_
                d_1985_v34_: _dafny.Map
                d_1985_v34_ = _dafny.Map({self.f3: p1})
                index302_ = default__.safeIndex(295, (d_1982_v33_).length(0))
                (d_1982_v33_)[index302_] = d_1985_v34_
                d_1986_v35_: _dafny.MultiSet
                d_1986_v35_ = _dafny.MultiSet([d_1931_v0_])
                index303_ = default__.safeIndex(295, (d_1982_v33_).length(0))
                (d_1982_v33_)[index303_] = default__.fm4(((d_1986_v35_)[d_1931_v0_] if (d_1931_v0_) in (d_1986_v35_) else (self).f4), (self.f3) not in (d_1937_v6_), globalState)
                (self).f3 = self.f3
                index304_ = default__.safeIndex(168, ((self).f5).length(0))
                ((self).f5)[index304_] = self.f3
                d_1987_v36_: _dafny.Seq
                d_1987_v36_ = _dafny.SeqWithoutIsStrInference([d_1938_v7_])
                index305_ = default__.safeIndex(168, ((self).f5).length(0))
                ((self).f5)[index305_] = (d_1987_v36_) < (d_1987_v36_)
            if ((self).f4) in ((d_1933_v2_ if self.f3 else d_1933_v2_)):
                (self).f3 = not(True)
                index306_ = default__.safeIndex(321, ((self).f5).length(0))
                ((self).f5)[index306_] = self.f3
                index307_ = default__.safeIndex(321, ((self).f5).length(0))
                ((self).f5)[index307_] = (default__.safeModuloInt(p1, (self).f4)) == ((d_1957_i3_ if True else d_1957_i3_))
                d_1988_v37_: _dafny.Seq
                d_1988_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cst"))
                d_1988_v37_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djhll"))) + (d_1988_v37_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djhll"))) + (d_1988_v37_))), d_1931_v0_)
                d_1989_v38_: _dafny.Map
                d_1989_v38_ = _dafny.Map({d_1957_i3_: (d_1955_v21_).set(p0, p0)})
                d_1990_v39_: C5
                nw271_ = C5()
                nw271_.ctor__(not(((self).f5)[default__.safeIndex(321, ((self).f5).length(0))]), (self).f1, ((d_1988_v37_) + (d_1988_v37_)).set(default__.safeIndex(len(d_1989_v38_), len((d_1988_v37_) + (d_1988_v37_))), d_1931_v0_))
                d_1990_v39_ = nw271_
                d_1991_v40_: _dafny.Set
                d_1991_v40_ = _dafny.Set({default__.fm22(((self).f5)[default__.safeIndex(321, ((self).f5).length(0))], 725, globalState)})
                d_1992_v41_: _dafny.Array
                nw272_ = _dafny.Array(None, 10)
                nw272_[int(0)] = not(default__.fm3(p0, globalState))
                nw272_[int(1)] = ((self).f5)[default__.safeIndex(321, ((self).f5).length(0))]
                nw272_[int(2)] = self.f3
                nw272_[int(3)] = self.f3
                nw272_[int(4)] = (d_1990_v39_).fm11((d_1990_v39_).f15, globalState)
                nw272_[int(5)] = (d_1931_v0_) not in (d_1991_v40_)
                nw272_[int(6)] = (self.f3) == (((self).f5)[default__.safeIndex(321, ((self).f5).length(0))])
                nw272_[int(7)] = (d_1990_v39_).f15
                nw272_[int(8)] = not(self.f3)
                nw272_[int(9)] = ((self).f5)[default__.safeIndex(321, ((self).f5).length(0))]
                d_1992_v41_ = nw272_
                d_1993_v42_: D6
                d_1993_v42_ = D6_DC17(self.f3)
                d_1994_v43_: _dafny.Array
                nw273_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_1994_v43_ = nw273_
                d_1995_v44_: _dafny.Map
                d_1995_v44_ = _dafny.Map({d_1994_v43_: True})
                nw274_ = _dafny.Array(None, 13)
                nw274_[int(0)] = (len(_dafny.SeqWithoutIsStrInference([d_1931_v0_ for d_1996_i5_ in range(default__.abs(131))]))) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abxt"))) for d_1997_i6_ in range(default__.abs(-276))]))
                nw274_[int(1)] = (d_1990_v39_).f15
                nw274_[int(2)] = ((self).f5)[default__.safeIndex(321, ((self).f5).length(0))]
                nw274_[int(3)] = self.f3
                nw274_[int(4)] = (False) or (not(self.f3))
                nw274_[int(5)] = (d_1993_v42_).cf28
                nw274_[int(6)] = not((d_1990_v39_).f15)
                nw274_[int(7)] = True
                nw274_[int(8)] = (self.f3 if self.f3 else self.f3)
                nw274_[int(9)] = (d_1990_v39_).f15
                nw274_[int(10)] = (d_1990_v39_).f15
                nw274_[int(11)] = (d_1990_v39_).f15
                nw274_[int(12)] = ((d_1995_v44_)[d_1994_v43_] if (d_1994_v43_) in (d_1995_v44_) else (d_1990_v39_).fm11(((self).f5)[default__.safeIndex(321, ((self).f5).length(0))], globalState))
                d_1992_v41_ = nw274_
            elif True:
                (self).f3 = self.f3
                d_1931_v0_ = d_1931_v0_
                d_1938_v7_ = d_1938_v7_
                d_1998_v45_: _dafny.MultiSet
                d_1998_v45_ = _dafny.MultiSet([p0])
                r2 = default__.safeDivisionInt((d_1998_v45_).cardinality, ((0) - (len(_dafny.Map({d_1955_v21_: self.f3})))) + (len(_dafny.SeqWithoutIsStrInference([d_1931_v0_ for d_1999_i7_ in range(default__.abs(78))]))))
                d_2000_v46_: _dafny.Seq
                d_2000_v46_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([self.f3])])
                d_2001_v47_: _dafny.Map
                d_2001_v47_ = _dafny.Map({(self).f4: (self).f2})
                d_2002_v48_: _dafny.Array
                nw275_ = _dafny.Array(None, 3)
                nw275_[int(0)] = d_2000_v46_
                nw275_[int(1)] = (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f3])) for d_2003_i8_ in range(default__.abs(627))])).set(default__.safeIndex(len(d_2001_v47_), len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f3])) for d_2004_i8_ in range(default__.abs(627))]))), d_1937_v6_)
                nw275_[int(2)] = d_2000_v46_
                d_2002_v48_ = nw275_
                index308_ = default__.safeIndex(582, (d_2002_v48_).length(0))
                (d_2002_v48_)[index308_] = d_2000_v46_
                index309_ = default__.safeIndex(582, (d_2002_v48_).length(0))
                rhs328_ = (default__.fm52((self).fm13(globalState), 775, self.f3, globalState)) + ((d_2000_v46_) + (d_2000_v46_))
                rhs329_ = d_1998_v45_
                lhs208_ = d_2002_v48_
                lhs209_ = default__.safeIndex(582, (d_2002_v48_).length(0))
                lhs208_[lhs209_] = rhs328_
                d_1998_v45_ = rhs329_
            d_2005_v49_: _dafny.Set
            d_2005_v49_ = _dafny.Set({self.f3})
            d_2006_v50_: _dafny.Map
            d_2006_v50_ = _dafny.Map({(self).f4: d_2005_v49_})
            d_2007_v51_: C6
            nw276_ = C6()
            nw276_.ctor__((821) + (p0), len(d_2006_v50_), default__.safeDivisionInt(p0, p0), (self).f5, self.f3, (self).f1, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "decv"))) + ((default__.fm2((self).fm11(self.f3, globalState), globalState)).set(default__.safeIndex(p0, len(default__.fm2((self).fm11(self.f3, globalState), globalState))), d_1931_v0_)))
            d_2007_v51_ = nw276_
            d_2008_v52_: _dafny.Map
            d_2008_v52_ = _dafny.Map({self.f3: self.f3})
            (self).f3 = (d_1957_i3_) < ((d_2007_v51_.f12) * (len((d_2008_v52_).set(self.f3, self.f3))))
        r2 = (self).fm12((self).f4, (self).f4, 479, self.f3, globalState)
        r0 = default__.fm9(globalState)
        r1 = (p1) - ((self).f4)
        r2 = p0
        return r0, r1, r2


class C8(T1, T2, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f10: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f10, f3, f1, f2, f4, f5):
        (self)._f10 = f10
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f4 = f4
        (self)._f5 = f5

    def fm12(self, p0, p1, p2, p3, globalState):
        if (D3_DC9((self).f4, self.f3)).cf16:
            return (self).f4
        elif True:
            return (self).f4

    def fm11(self, p0, globalState):
        source37_ = (D0_DC1((self).f4) if self.f3 else D0_DC1(len((self).f10)))
        if source37_.is_DC1:
            d_2009___mcc_h0_ = source37_.cf1
            d_2010_cf1_ = d_2009___mcc_h0_
            return self.f3
        elif True:
            d_2011___mcc_h1_ = source37_.cf0
            d_2012_cf0_ = d_2011___mcc_h1_
            return not((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f3]))).issubset(_dafny.MultiSet([self.f3])))

    def fm13(self, globalState):
        def iife192_():
            coll60_ = _dafny.Map()
            compr_60_: int
            for compr_60_ in (_dafny.SeqWithoutIsStrInference([(0) - ((self).f4) for d_2013_i0_ in range(default__.abs(583))])).Elements:
                d_2014_v0_: int = compr_60_
                if (d_2014_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - ((self).f4) for d_2013_i0_ in range(default__.abs(583))])):
                    coll60_[(d_2014_v0_) + ((self).f4)] = self.f3
            return _dafny.Map(coll60_)
        return ((self).f4) not in (((D6_DC16(_dafny.Map({len(_dafny.Set({False})): self.f3}))).cf27) | (iife192_()
        ))

    def fm14(self, p0, p1, p2, p3, globalState):
        return ((self).f2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))

    def fm16(self, p0, p1, p2, globalState):
        return ((self).f2).set(default__.safeIndex(((self).f4 if self.f3 else len(_dafny.Set({_dafny.Map({_dafny.CodePoint('b'): self.f3})}))), len((self).f2)), _dafny.CodePoint('h'))

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        def lambda78_(source38_):
            if source38_.is_DC9:
                d_2015___mcc_h0_ = source38_.cf15
                d_2016___mcc_h1_ = source38_.cf16
                d_2017_cf16_ = d_2016___mcc_h1_
                d_2018_cf15_ = d_2015___mcc_h0_
                return d_2017_cf16_
            elif source38_.is_DC10:
                d_2019___mcc_h2_ = source38_.cf17
                d_2020___mcc_h3_ = source38_.cf18
                d_2021_cf18_ = d_2020___mcc_h3_
                d_2022_cf17_ = d_2019___mcc_h2_
                return self.f3
            elif source38_.is_DC8:
                d_2023___mcc_h4_ = source38_.cf14
                d_2024_cf14_ = d_2023___mcc_h4_
                return (True) == (False)
            elif True:
                d_2025___mcc_h5_ = source38_.cf19
                d_2026_cf19_ = d_2025___mcc_h5_
                return self.f3

        if not(lambda78_(default__.fm17(p1, globalState))):
            d_2027_v0_: _dafny.Seq
            d_2027_v0_ = _dafny.SeqWithoutIsStrInference([p0, p1])
            r2 = (d_2027_v0_)[default__.safeIndex((self).f4, len(d_2027_v0_))]
            (self).f3 = self.f3
            d_2028_v1_: D5
            d_2028_v1_ = D5_DC15(p1, self.f3, (self).f5)
            d_2028_v1_ = D5_DC15(p0, not((p0) >= (p0)), (self).f5)
            d_2029_v2_: _dafny.Seq
            d_2029_v2_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
            d_2030_v3_: str
            d_2030_v3_ = _dafny.CodePoint('y')
            d_2031_v4_: _dafny.Map
            d_2031_v4_ = _dafny.Map({d_2030_v3_: p1})
            d_2032_v5_: _dafny.Map
            d_2032_v5_ = _dafny.Map({self.f3: p0})
            d_2033_v6_: _dafny.MultiSet
            d_2033_v6_ = _dafny.MultiSet([p1, 549])
            d_2034_v7_: _dafny.Array
            nw277_ = _dafny.Array(None, 21)
            nw277_[int(0)] = (d_2029_v2_)[default__.safeIndex(len((self).f2), len(d_2029_v2_))]
            nw277_[int(1)] = self.f3
            nw277_[int(2)] = (not(self.f3) if self.f3 else self.f3)
            nw277_[int(3)] = self.f3
            nw277_[int(4)] = not(self.f3)
            nw277_[int(5)] = self.f3
            nw277_[int(6)] = ((0) - (((d_2031_v4_)[d_2030_v3_] if (d_2030_v3_) in (d_2031_v4_) else p0))) >= (len(d_2029_v2_))
            nw277_[int(7)] = self.f3
            nw277_[int(8)] = self.f3
            nw277_[int(9)] = not (self.f3) or (self.f3)
            nw277_[int(10)] = self.f3
            nw277_[int(11)] = (p0) not in (d_2027_v0_)
            nw277_[int(12)] = False
            nw277_[int(13)] = (((self).f2).set(default__.safeIndex(p0, len((self).f2)), d_2030_v3_)) <= ((self).f2)
            nw277_[int(14)] = not (self.f3) or (self.f3)
            nw277_[int(15)] = (self.f3 if self.f3 else self.f3)
            nw277_[int(16)] = not (self.f3) or (self.f3)
            nw277_[int(17)] = (len(d_2032_v5_)) >= ((self).fm12(871, p0, p0, self.f3, globalState))
            nw277_[int(18)] = False
            nw277_[int(19)] = True
            nw277_[int(20)] = (659) not in (d_2033_v6_)
            d_2034_v7_ = nw277_
            d_2034_v7_ = ((d_2034_v7_ if self.f3 else (self).f5) if self.f3 else (self).f5)
            if ((self).f4) > (p1):
                d_2035_v8_: _dafny.Seq
                d_2035_v8_ = _dafny.SeqWithoutIsStrInference([((self).f2)[default__.safeIndex(p1, len((self).f2))]])
                d_2035_v8_ = (d_2035_v8_) + ((self).f2)
                d_2036_v9_: _dafny.Array
                nw278_ = _dafny.Array(_dafny.Set({}), 9)
                d_2036_v9_ = nw278_
                d_2037_v10_: _dafny.Set
                d_2037_v10_ = _dafny.Set({p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsar")))})
                index310_ = default__.safeIndex(839, (d_2036_v9_).length(0))
                (d_2036_v9_)[index310_] = d_2037_v10_
                index311_ = default__.safeIndex(839, (d_2036_v9_).length(0))
                (d_2036_v9_)[index311_] = _dafny.Set({p0, (self).fm12((self).fm12(len(d_2027_v0_), (self).f4, (self).f4, self.f3, globalState), p1, (self).f4, self.f3, globalState)})
                index312_ = default__.safeIndex(839, (d_2036_v9_).length(0))
                (d_2036_v9_)[index312_] = (d_2036_v9_)[default__.safeIndex(839, (d_2036_v9_).length(0))]
                d_2038_v11_: _dafny.Array
                nw279_ = _dafny.Array(None, 8)
                nw279_[int(0)] = _dafny.CodePoint('m')
                nw279_[int(1)] = d_2030_v3_
                nw279_[int(2)] = d_2030_v3_
                nw279_[int(3)] = d_2030_v3_
                nw279_[int(4)] = d_2030_v3_
                nw279_[int(5)] = _dafny.CodePoint('e')
                nw279_[int(6)] = d_2030_v3_
                nw279_[int(7)] = d_2030_v3_
                d_2038_v11_ = nw279_
                d_2039_v12_: _dafny.Seq
                d_2039_v12_ = _dafny.SeqWithoutIsStrInference([d_2038_v11_, d_2038_v11_])
                d_2040_v13_: _dafny.Array
                nw280_ = _dafny.Array(None, 3)
                nw280_[int(0)] = (d_2039_v12_)[default__.safeIndex(p0, len(d_2039_v12_))]
                nw280_[int(1)] = d_2038_v11_
                nw280_[int(2)] = d_2038_v11_
                d_2040_v13_ = nw280_
                d_2041_v14_: _dafny.Set
                d_2041_v14_ = _dafny.Set({not((self).fm13(globalState))})
                d_2042_v15_: _dafny.Array
                d_2042_v15_ = d_2040_v13_
                rhs330_ = len(d_2041_v14_)
                rhs331_ = not(self.f3)
                rhs332_ = (d_2042_v15_)
                rhs333_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2043_i0_ in range(default__.abs(380))])) + (((self).f2).set(default__.safeIndex((self).f4, len((self).f2)), d_2030_v3_))) + (((self).f2) + (d_2035_v8_))
                lhs210_ = self
                r1 = rhs330_
                lhs210_.f3 = rhs331_
                d_2040_v13_ = rhs332_
                d_2035_v8_ = rhs333_
                d_2044_v16_: _dafny.Map
                d_2044_v16_ = _dafny.Map({(self).f5: (self).fm14(self.f3, default__.fm3(len(d_2027_v0_), globalState), d_2030_v3_, p1, globalState)})
                rhs334_ = (self).f4
                rhs335_ = (d_2035_v8_ if not (self.f3) or (self.f3) else (((d_2044_v16_)[d_2034_v7_] if (d_2034_v7_) in (d_2044_v16_) else d_2035_v8_)) + ((self).f2))
                rhs336_ = p1
                rhs337_ = self.f3
                lhs211_ = self
                r2 = rhs334_
                d_2035_v8_ = rhs335_
                r2 = rhs336_
                lhs211_.f3 = rhs337_
            elif True:
                d_2045_v17_: _dafny.Map
                d_2045_v17_ = _dafny.Map({d_2034_v7_: p0})
                d_2045_v17_ = (d_2045_v17_).set(d_2034_v7_, (d_2027_v0_)[default__.safeIndex((self).f4, len(d_2027_v0_))])
                r1 = (self).f4
                (self).f3 = self.f3
                d_2046_v18_: _dafny.Seq
                d_2046_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xildsfr"))
                d_2046_v18_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jogglsdwh"))) + ((self).f2)
                d_2047_v19_: C2
                nw281_ = C2()
                nw281_.ctor__((p0 if self.f3 else p1), self.f3, (self).f1, _dafny.SeqWithoutIsStrInference([d_2030_v3_ for d_2048_i1_ in range(default__.abs(701))]))
                d_2047_v19_ = nw281_
        elif True:
            index313_ = default__.safeIndex(407, ((self).f5).length(0))
            ((self).f5)[index313_] = not((p1) < (p0))
            d_2049_v20_: _dafny.Set
            d_2049_v20_ = _dafny.Set({len((self).f2), p1, p1})
            d_2050_v21_: _dafny.MultiSet
            d_2050_v21_ = _dafny.MultiSet([d_2049_v20_])
            index314_ = default__.safeIndex(75, ((self).f5).length(0))
            ((self).f5)[index314_] = (_dafny.MultiSet([d_2049_v20_])).ispropersubset(d_2050_v21_)
            d_2051_v22_: _dafny.Set
            d_2051_v22_ = _dafny.Set({not(self.f3)})
            index315_ = default__.safeIndex(407, ((self).f5).length(0))
            index316_ = default__.safeIndex(75, ((self).f5).length(0))
            rhs338_ = (p1) + (62)
            rhs339_ = p1
            rhs340_ = not(self.f3)
            rhs341_ = (((self).f2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggplncq")))) < (default__.fm2(True, globalState))
            rhs342_ = default__.fm3(len(d_2051_v22_), globalState)
            lhs212_ = (self).f5
            lhs213_ = default__.safeIndex(407, ((self).f5).length(0))
            lhs214_ = (self).f5
            lhs215_ = default__.safeIndex(75, ((self).f5).length(0))
            lhs216_ = self
            r2 = rhs338_
            r2 = rhs339_
            lhs212_[lhs213_] = rhs340_
            lhs214_[lhs215_] = rhs341_
            lhs216_.f3 = rhs342_
            d_2052_v23_: _dafny.Seq
            d_2052_v23_ = _dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(407, ((self).f5).length(0))]])
            d_2053_v24_: _dafny.Array
            def lambda79_(d_2054_i2_):
                return False

            init43_ = lambda79_
            nw282_ = _dafny.Array(None, 25)
            for i0_43_ in range(nw282_.length(0)):
                nw282_[i0_43_] = init43_(i0_43_)
            d_2053_v24_ = nw282_
            d_2055_v25_: _dafny.MultiSet
            d_2055_v25_ = _dafny.MultiSet([p0, len((self).f2), p1, p0, p0])
            d_2056_v26_: _dafny.Map
            d_2056_v26_ = _dafny.Map({d_2053_v24_: d_2055_v25_})
            d_2057_v27_: _dafny.Array
            nw283_ = _dafny.Array(None, 18)
            nw283_[int(0)] = (d_2052_v23_) + (d_2052_v23_)
            nw283_[int(1)] = ((d_2052_v23_).set(default__.safeIndex((0) - ((((d_2056_v26_)[d_2053_v24_] if (d_2053_v24_) in (d_2056_v26_) else d_2055_v25_)).cardinality), len(d_2052_v23_)), ((self).f5)[default__.safeIndex(407, ((self).f5).length(0))])) + (d_2052_v23_)
            nw283_[int(2)] = d_2052_v23_
            nw283_[int(3)] = d_2052_v23_
            nw283_[int(4)] = default__.fm7(self.f3, (d_2052_v23_)[default__.safeIndex(p0, len(d_2052_v23_))], globalState)
            nw283_[int(5)] = _dafny.SeqWithoutIsStrInference([self.f3, ((self).f5)[default__.safeIndex(407, ((self).f5).length(0))]])
            nw283_[int(6)] = default__.fm7(((self).f5)[default__.safeIndex(407, ((self).f5).length(0))], self.f3, globalState)
            nw283_[int(7)] = ((self).f10)[default__.safeIndex(396, len((self).f10))]
            nw283_[int(8)] = (d_2052_v23_ if ((self).f5)[default__.safeIndex(407, ((self).f5).length(0))] else d_2052_v23_)
            nw283_[int(9)] = (d_2052_v23_) + (d_2052_v23_)
            nw283_[int(10)] = (d_2052_v23_).set(default__.safeIndex((self).f4, len(d_2052_v23_)), ((self).f5)[default__.safeIndex(407, ((self).f5).length(0))])
            nw283_[int(11)] = ((d_2052_v23_) + (d_2052_v23_)).set(default__.safeIndex(-153, len((d_2052_v23_) + (d_2052_v23_))), self.f3)
            nw283_[int(12)] = d_2052_v23_
            nw283_[int(13)] = d_2052_v23_
            nw283_[int(14)] = (d_2052_v23_) + (d_2052_v23_)
            nw283_[int(15)] = _dafny.SeqWithoutIsStrInference([self.f3, False])
            nw283_[int(16)] = d_2052_v23_
            nw283_[int(17)] = d_2052_v23_
            d_2057_v27_ = nw283_
            d_2057_v27_ = d_2057_v27_
            d_2058_v28_: _dafny.Array
            nw284_ = _dafny.Array(None, 3)
            nw284_[int(0)] = d_2053_v24_
            nw284_[int(1)] = d_2053_v24_
            nw284_[int(2)] = d_2053_v24_
            d_2058_v28_ = nw284_
            index317_ = default__.safeIndex(185, (d_2058_v28_).length(0))
            (d_2058_v28_)[index317_] = d_2053_v24_
            index318_ = default__.safeIndex(185, (d_2058_v28_).length(0))
            (d_2058_v28_)[index318_] = d_2053_v24_
            d_2059_v29_: _dafny.Array
            nw285_ = _dafny.Array(int(0), 29)
            d_2059_v29_ = nw285_
            d_2060_v30_: C0
            nw286_ = C0()
            nw286_.ctor__(d_2059_v29_)
            d_2060_v30_ = nw286_
            r1 = len((self).f2)
        index319_ = default__.safeIndex(108, ((self).f5).length(0))
        ((self).f5)[index319_] = self.f3
        d_2061_v31_: D11
        d_2061_v31_ = D11_DC31((self).f4)
        pat_let_tv50_ = p1
        pat_let_tv51_ = p1
        index320_ = default__.safeIndex(108, ((self).f5).length(0))
        def lambda80_(source39_):
            if source39_.is_DC31:
                d_2062___mcc_h6_ = source39_.cf50
                d_2063_cf50_ = d_2062___mcc_h6_
                return not (self.f3) or (self.f3)
            elif True:
                d_2064___mcc_h7_ = source39_.cf49
                d_2065_cf49_ = d_2064___mcc_h7_
                return (pat_let_tv50_) > (pat_let_tv51_)

        rhs343_ = p1
        rhs344_ = self.f3
        rhs345_ = lambda80_(d_2061_v31_)
        rhs346_ = p1
        rhs347_ = False
        lhs217_ = self
        lhs218_ = (self).f5
        lhs219_ = default__.safeIndex(108, ((self).f5).length(0))
        lhs220_ = self
        r2 = rhs343_
        lhs217_.f3 = rhs344_
        lhs218_[lhs219_] = rhs345_
        r1 = rhs346_
        lhs220_.f3 = rhs347_
        d_2066_v32_: _dafny.Array
        nw287_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_2066_v32_ = nw287_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2066_v32_).length(0)):
            d_2067_i3_: int = guard_loop_5_
            if (True) and (((0) <= (d_2067_i3_)) and ((d_2067_i3_) < ((d_2066_v32_).length(0)))):
                (d_2066_v32_)[(d_2067_i3_)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihcix"))
        r1 = (0) - (len((self).f2))
        d_2068_v33_: _dafny.Array
        nw288_ = _dafny.Array(int(0), 23)
        d_2068_v33_ = nw288_
        index321_ = default__.safeIndex(461, (d_2068_v33_).length(0))
        (d_2068_v33_)[index321_] = p0
        index322_ = default__.safeIndex(461, (d_2068_v33_).length(0))
        (d_2068_v33_)[index322_] = (0) - ((p0) - (default__.safeDivisionInt((self).f4, p1)))
        index323_ = default__.safeIndex(108, ((self).f5).length(0))
        ((self).f5)[index323_] = self.f3
        d_2069_v34_: _dafny.Map
        d_2069_v34_ = _dafny.Map({((self).f5)[default__.safeIndex(108, ((self).f5).length(0))]: self.f3})
        r0 = (d_2069_v34_) | (d_2069_v34_)
        d_2070_v35_: _dafny.Set
        d_2070_v35_ = _dafny.Set({((self).f5)[default__.safeIndex(108, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(108, ((self).f5).length(0))]})
        r1 = len(((d_2070_v35_).intersection(d_2070_v35_)).intersection((d_2070_v35_).intersection(d_2070_v35_)))
        r2 = p1
        return r0, r1, r2

    def m6(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_2071_v0_: str
        d_2071_v0_ = _dafny.CodePoint('m')
        d_2072_v1_: _dafny.Map
        d_2072_v1_ = _dafny.Map({p2: (self).f4})
        d_2073_v2_: D6
        d_2073_v2_ = D6_DC18(p2, p2, self.f3, d_2071_v0_, ((d_2072_v1_)[p2] if (p2) in (d_2072_v1_) else p2))
        d_2074_v3_: _dafny.Map
        d_2074_v3_ = _dafny.Map({d_2073_v2_: (D6_DC18(489, (self).f4, self.f3, d_2071_v0_, (self).f4)).cf32})
        d_2074_v3_ = (d_2074_v3_).set(d_2073_v2_, d_2071_v0_)
        d_2075_v4_: _dafny.Seq
        d_2075_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtawg"))
        d_2075_v4_ = (_dafny.SeqWithoutIsStrInference([(d_2075_v4_)[default__.safeIndex(p2, len(d_2075_v4_))] for d_2076_i0_ in range(default__.abs(804))]) if self.f3 else (self).f2)
        d_2077_v5_: _dafny.Array
        nw289_ = _dafny.Array(None, 1)
        nw289_[int(0)] = p2
        d_2077_v5_ = nw289_
        d_2078_v6_: D5
        d_2078_v6_ = D5_DC14(d_2077_v5_)
        d_2079_v7_: _dafny.Map
        d_2079_v7_ = _dafny.Map({d_2078_v6_: False})
        if ((d_2079_v7_)[d_2078_v6_] if (d_2078_v6_) in (d_2079_v7_) else self.f3):
            d_2080_v9_: _dafny.Array
            def lambda81_(d_2081_i1_):
                def iife193_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in _dafny.IntegerRange(4, 213):
                        d_2082_v8_: int = compr_61_
                        if ((4) <= (d_2082_v8_)) and ((d_2082_v8_) < (213)):
                            coll61_ = coll61_.union(_dafny.Set([(d_2082_v8_) + (502)]))
                    return _dafny.Set(coll61_)
                return _dafny.Map({self.f3: iife193_()
                })

            init44_ = lambda81_
            nw290_ = _dafny.Array(None, 13)
            for i0_44_ in range(nw290_.length(0)):
                nw290_[i0_44_] = init44_(i0_44_)
            d_2080_v9_ = nw290_
            d_2083_v10_: _dafny.MultiSet
            d_2083_v10_ = _dafny.MultiSet([p2])
            d_2084_v11_: _dafny.Seq
            d_2084_v11_ = _dafny.SeqWithoutIsStrInference([d_2083_v10_])
            d_2085_v12_: _dafny.Set
            d_2085_v12_ = _dafny.Set({(self).f4, p2, (self).f4, (self).f4})
            d_2086_v13_: _dafny.Map
            d_2086_v13_ = _dafny.Map({default__.fm3(len(d_2084_v11_), globalState): d_2085_v12_})
            index324_ = default__.safeIndex(552, (d_2080_v9_).length(0))
            (d_2080_v9_)[index324_] = d_2086_v13_
            d_2087_v14_: _dafny.Set
            d_2087_v14_ = _dafny.Set({self.f3, self.f3})
            index325_ = default__.safeIndex(552, (d_2080_v9_).length(0))
            rhs348_ = d_2086_v13_
            rhs349_ = ((d_2087_v14_) | (d_2087_v14_)).issubset(d_2087_v14_)
            lhs221_ = d_2080_v9_
            lhs222_ = default__.safeIndex(552, (d_2080_v9_).length(0))
            lhs223_ = self
            lhs221_[lhs222_] = rhs348_
            lhs223_.f3 = rhs349_
            index326_ = default__.safeIndex(631, ((self).f5).length(0))
            ((self).f5)[index326_] = self.f3
            index327_ = default__.safeIndex(631, ((self).f5).length(0))
            ((self).f5)[index327_] = (self).fm11((len((_dafny.Map({p2: self.f3})).set((self).f4, self.f3))) < (-305), globalState)
            d_2088_v15_: D8
            d_2088_v15_ = D8_DC23(len(d_2072_v1_), p2, ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))])
            index328_ = default__.safeIndex(631, ((self).f5).length(0))
            ((self).f5)[index328_] = (self).fm11((d_2088_v15_).cf40, globalState)
            if ((_dafny.Set({not(False), self.f3, self.f3})).intersection(d_2087_v14_)).issubset(d_2087_v14_):
                index329_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                (d_2077_v5_)[index329_] = p2
                index330_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                (d_2077_v5_)[index330_] = (self).f4
                d_2089_v16_: _dafny.Map
                d_2089_v16_ = _dafny.Map({((self).f5)[default__.safeIndex(631, ((self).f5).length(0))]: False})
                d_2090_v17_: _dafny.Seq
                d_2090_v17_ = _dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(631, ((self).f5).length(0))], True, ((d_2089_v16_)[self.f3] if (self.f3) in (d_2089_v16_) else self.f3), self.f3])
                d_2091_v18_: _dafny.MultiSet
                d_2091_v18_ = _dafny.MultiSet([d_2090_v17_, _dafny.SeqWithoutIsStrInference([self.f3, ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))], self.f3, ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))]])])
                index331_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                (d_2077_v5_)[index331_] = (d_2091_v18_).cardinality
                d_2092_v19_: _dafny.Array
                nw291_ = _dafny.Array(False, 12)
                d_2092_v19_ = nw291_
                d_2092_v19_ = d_2092_v19_
                d_2093_v20_: C4
                nw292_ = C4()
                nw292_.ctor__((d_2083_v10_).isdisjoint(d_2083_v10_), _dafny.Map({(self).f2: True}), (d_2075_v4_) + (d_2075_v4_))
                d_2093_v20_ = nw292_
                d_2094_v21_: _dafny.Map
                d_2094_v21_ = _dafny.Map({((self).f5)[default__.safeIndex(631, ((self).f5).length(0))]: d_2072_v1_})
                index332_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                index333_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                rhs350_ = default__.safeModuloInt(len(d_2094_v21_), -125)
                rhs351_ = d_2071_v0_
                rhs352_ = (self).f4
                lhs224_ = d_2077_v5_
                lhs225_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                lhs226_ = d_2077_v5_
                lhs227_ = default__.safeIndex(973, (d_2077_v5_).length(0))
                lhs224_[lhs225_] = rhs350_
                d_2071_v0_ = rhs351_
                lhs226_[lhs227_] = rhs352_
            elif True:
                d_2095_v22_: int
                d_2095_v22_ = 816
                d_2095_v22_ = (0) - (((0) - ((self).f4) if ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))] else ((0) - (p2)) + (d_2095_v22_)))
                d_2096_v23_: _dafny.Map
                d_2096_v23_ = _dafny.Map({d_2077_v5_: (self).f2})
                d_2097_v24_: _dafny.Map
                d_2097_v24_ = _dafny.Map({self.f3: d_2096_v23_})
                d_2097_v24_ = (d_2097_v24_).set(self.f3, _dafny.Map({d_2077_v5_: (self).f2}))
                d_2098_v25_: _dafny.MultiSet
                d_2098_v25_ = _dafny.MultiSet([self.f3, ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))]])
                d_2099_v26_: C2
                nw293_ = C2()
                nw293_.ctor__((self).f4, self.f3, (self).f1, (self).f2)
                d_2099_v26_ = nw293_
                d_2100_v27_: _dafny.Map
                d_2100_v27_ = _dafny.Map({d_2099_v26_: self.f3})
                d_2101_v28_: D13
                d_2101_v28_ = D13_DC33(d_2099_v26_)
                d_2102_v29_: _dafny.Seq
                d_2102_v29_ = _dafny.SeqWithoutIsStrInference([not(self.f3), self.f3, ((d_2100_v27_)[(D13_DC33((d_2101_v28_).cf52)).cf52] if ((D13_DC33((d_2101_v28_).cf52)).cf52) in (d_2100_v27_) else (self).fm13(globalState))])
                index334_ = default__.safeIndex(631, ((self).f5).length(0))
                rhs353_ = 239
                rhs354_ = (_dafny.MultiSet(d_2102_v29_)) | ((d_2098_v25_) | (d_2098_v25_))
                rhs355_ = ((d_2075_v4_) + (d_2075_v4_)) + ((_dafny.SeqWithoutIsStrInference([d_2071_v0_ for d_2103_i2_ in range(default__.abs(-897))])) + (_dafny.SeqWithoutIsStrInference([d_2071_v0_ for d_2104_i3_ in range(default__.abs(325))])))
                rhs356_ = p2
                rhs357_ = ((self).f5)[default__.safeIndex(631, ((self).f5).length(0))]
                lhs228_ = (self).f5
                lhs229_ = default__.safeIndex(631, ((self).f5).length(0))
                d_2095_v22_ = rhs353_
                d_2098_v25_ = rhs354_
                d_2075_v4_ = rhs355_
                d_2095_v22_ = rhs356_
                lhs228_[lhs229_] = rhs357_
                index335_ = default__.safeIndex(462, (d_2077_v5_).length(0))
                (d_2077_v5_)[index335_] = ((d_2099_v26_).f13) * (p2)
                index336_ = default__.safeIndex(462, (d_2077_v5_).length(0))
                (d_2077_v5_)[index336_] = (d_2099_v26_).f13
                index337_ = default__.safeIndex(462, (d_2077_v5_).length(0))
                (d_2077_v5_)[index337_] = p2
            d_2105_v30_: int
            d_2105_v30_ = 868
            d_2105_v30_ = d_2105_v30_
        elif True:
            d_2106_v31_: _dafny.Seq
            d_2106_v31_ = _dafny.SeqWithoutIsStrInference([self.f3, (d_2073_v2_).cf31])
            (self).f3 = not((d_2106_v31_)[default__.safeIndex((self).f4, len(d_2106_v31_))])
            d_2107_v32_: _dafny.Seq
            d_2107_v32_ = _dafny.SeqWithoutIsStrInference([len(d_2072_v1_), (self).f4])
            d_2108_v33_: _dafny.Map
            d_2108_v33_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wftn"))})
            d_2109_v34_: D11
            d_2109_v34_ = D11_DC30(d_2108_v33_)
            d_2110_v35_: D13
            d_2110_v35_ = D13_DC35(18, (d_2107_v32_)[default__.safeIndex(p2, len(d_2107_v32_))], self.f3, len((d_2075_v4_).set(default__.safeIndex(p2, len(d_2075_v4_)), d_2071_v0_)), d_2109_v34_)
            d_2111_v36_: _dafny.Map
            d_2111_v36_ = _dafny.Map({d_2110_v35_: d_2071_v0_})
            rhs358_ = d_2071_v0_
            rhs359_ = (default__.fm53((self).f4, globalState)) | (d_2111_v36_)
            rhs360_ = default__.fm3(p2, globalState)
            lhs230_ = self
            d_2071_v0_ = rhs358_
            d_2111_v36_ = rhs359_
            lhs230_.f3 = rhs360_
            (self).f3 = self.f3
            d_2112_v37_: _dafny.Map
            d_2112_v37_ = _dafny.Map({p2: self.f3})
            d_2113_v38_: _dafny.MultiSet
            d_2113_v38_ = _dafny.MultiSet([d_2075_v4_, (self).f2])
            if ((d_2112_v37_)[((d_2113_v38_).cardinality) - (674)] if (((d_2113_v38_).cardinality) - (674)) in (d_2112_v37_) else self.f3):
                index338_ = default__.safeIndex(617, ((self).f5).length(0))
                ((self).f5)[index338_] = True
                index339_ = default__.safeIndex(617, ((self).f5).length(0))
                def iife194_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(387, 160):
                        d_2114_v39_: int = compr_62_
                        if ((387) <= (d_2114_v39_)) and ((d_2114_v39_) < (160)):
                            coll62_[default__.safeModuloInt(d_2114_v39_, p2)] = (self).f4
                    return _dafny.Map(coll62_)
                ((self).f5)[index339_] = ((iife194_()
                ) | (d_2072_v1_)) == (d_2072_v1_)
                d_2115_v40_: _dafny.Map
                d_2115_v40_ = _dafny.Map({self.f3: p2})
                d_2116_v41_: _dafny.Set
                d_2116_v41_ = _dafny.Set({default__.fm54(d_2115_v40_, self.f3, globalState)})
                d_2117_v42_: _dafny.Map
                d_2117_v42_ = _dafny.Map({False: d_2116_v41_})
                d_2118_v43_: _dafny.Array
                def lambda82_(d_2119_i4_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fswednf"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hs")))

                init45_ = lambda82_
                nw294_ = _dafny.Array(None, 26)
                for i0_45_ in range(nw294_.length(0)):
                    nw294_[i0_45_] = init45_(i0_45_)
                d_2118_v43_ = nw294_
                index340_ = default__.safeIndex(617, ((self).f5).length(0))
                index341_ = default__.safeIndex(617, ((self).f5).length(0))
                rhs361_ = (d_2117_v42_).set((True) or (not(False)), (d_2116_v41_) - (d_2116_v41_))
                rhs362_ = d_2112_v37_
                rhs363_ = ((self).f5)[default__.safeIndex(617, ((self).f5).length(0))]
                rhs364_ = ((self).f5)[default__.safeIndex(617, ((self).f5).length(0))]
                rhs365_ = d_2118_v43_
                lhs231_ = (self).f5
                lhs232_ = default__.safeIndex(617, ((self).f5).length(0))
                lhs233_ = (self).f5
                lhs234_ = default__.safeIndex(617, ((self).f5).length(0))
                d_2117_v42_ = rhs361_
                d_2112_v37_ = rhs362_
                lhs231_[lhs232_] = rhs363_
                lhs233_[lhs234_] = rhs364_
                d_2118_v43_ = rhs365_
                d_2120_v44_: _dafny.Array
                def lambda83_(d_2121_i5_):
                    return ((self).f5)[default__.safeIndex(617, ((self).f5).length(0))]

                init46_ = lambda83_
                nw295_ = _dafny.Array(None, 3)
                for i0_46_ in range(nw295_.length(0)):
                    nw295_[i0_46_] = init46_(i0_46_)
                d_2120_v44_ = nw295_
                d_2122_v45_: _dafny.Map
                d_2122_v45_ = _dafny.Map({d_2120_v44_: False})
                d_2123_v46_: bool
                d_2124_v47_: bool
                d_2125_v48_: _dafny.Seq
                d_2126_v49_: _dafny.Seq
                out46_: bool
                out47_: bool
                out48_: _dafny.Seq
                out49_: _dafny.Seq
                out46_, out47_, out48_, out49_ = default__.m0(d_2122_v45_, globalState)
                d_2123_v46_ = out46_
                d_2124_v47_ = out47_
                d_2125_v48_ = out48_
                d_2126_v49_ = out49_
                d_2127_v50_: D2
                d_2127_v50_ = D2_DC6(d_2071_v0_, d_2124_v47_, self.f3, (self).f4)
                d_2128_v51_: _dafny.Map
                d_2128_v51_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p2, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))]): d_2127_v50_})
                d_2128_v51_ = (d_2128_v51_).set(d_2107_v32_, d_2127_v50_)
                d_2129_v52_: bool
                d_2130_v53_: bool
                d_2131_v54_: _dafny.Seq
                d_2132_v55_: _dafny.Seq
                out50_: bool
                out51_: bool
                out52_: _dafny.Seq
                out53_: _dafny.Seq
                out50_, out51_, out52_, out53_ = default__.m0(d_2122_v45_, globalState)
                d_2129_v52_ = out50_
                d_2130_v53_ = out51_
                d_2131_v54_ = out52_
                d_2132_v55_ = out53_
            elif True:
                d_2133_v56_: _dafny.Array
                def lambda84_(d_2134_i6_):
                    return not (self.f3) or (self.f3)

                init47_ = lambda84_
                nw296_ = _dafny.Array(None, 26)
                for i0_47_ in range(nw296_.length(0)):
                    nw296_[i0_47_] = init47_(i0_47_)
                d_2133_v56_ = nw296_
                d_2135_v57_: _dafny.Seq
                d_2135_v57_ = _dafny.SeqWithoutIsStrInference([d_2133_v56_, d_2133_v56_, (self).f5, (self).f5])
                d_2136_v58_: _dafny.Map
                d_2136_v58_ = _dafny.Map({self.f3: (self).f4})
                rhs366_ = self.f3
                rhs367_ = (d_2135_v57_)[default__.safeIndex(((self).f4) + (len(d_2136_v58_)), len(d_2135_v57_))]
                lhs235_ = self
                lhs235_.f3 = rhs366_
                d_2133_v56_ = rhs367_
                d_2137_v59_: int
                d_2137_v59_ = 411
                d_2137_v59_ = (0) - (p2)
                (self).f3 = self.f3
                d_2136_v58_ = (d_2136_v58_).set(False, (0) - ((self).f4))
                d_2138_v60_: _dafny.MultiSet
                d_2138_v60_ = _dafny.MultiSet([(self).f4, len(d_2107_v32_)])
                d_2139_v61_: D2
                d_2139_v61_ = D2_DC7(D2_DC7(default__.fm10(self.f3, self.f3, self.f3, globalState)))
                d_2140_v62_: _dafny.Map
                d_2140_v62_ = _dafny.Map({d_2138_v60_: d_2139_v61_})
                d_2140_v62_ = (d_2140_v62_).set((_dafny.MultiSet([d_2137_v59_, p2])) - (d_2138_v60_), d_2139_v61_)
            d_2141_v63_: int
            d_2141_v63_ = -289
            d_2142_v64_: _dafny.Set
            d_2142_v64_ = _dafny.Set({(self).f4})
            d_2143_v65_: _dafny.Map
            d_2143_v65_ = _dafny.Map({len((self).f2): d_2142_v64_})
            d_2144_v66_: _dafny.MultiSet
            d_2144_v66_ = _dafny.MultiSet([self.f3])
            d_2141_v63_ = len(((d_2143_v65_)[((d_2144_v66_ if self.f3 else d_2144_v66_)).cardinality] if (((d_2144_v66_ if self.f3 else d_2144_v66_)).cardinality) in (d_2143_v65_) else d_2142_v64_))
        if self.f3:
            d_2145_v67_: _dafny.Set
            d_2145_v67_ = _dafny.Set({(self).f4})
            (self).f3 = (d_2145_v67_).ispropersubset(d_2145_v67_)
            d_2075_v4_ = (self).fm14(not(self.f3), ((0) - (len((self).f2))) < (132), d_2071_v0_, (self).f4, globalState)
            (self).f3 = self.f3
            d_2071_v0_ = d_2071_v0_
            d_2146_v68_: _dafny.Map
            d_2146_v68_ = _dafny.Map({(self).f4: self.f3})
            d_2147_v69_: C3
            nw297_ = C3()
            nw297_.ctor__(((d_2072_v1_)[len((d_2075_v4_).set(default__.safeIndex(len(d_2075_v4_), len(d_2075_v4_)), d_2071_v0_))] if (len((d_2075_v4_).set(default__.safeIndex(len(d_2075_v4_), len(d_2075_v4_)), d_2071_v0_))) in (d_2072_v1_) else (self).f4), (self).f5, ((0) - ((self).f4)) != (len(d_2146_v68_)), (self).f1, d_2075_v4_)
            d_2147_v69_ = nw297_
        elif True:
            d_2148_v70_: _dafny.Set
            d_2148_v70_ = _dafny.Set({self.f3})
            d_2149_v71_: D14
            d_2149_v71_ = D14_DC39(p2, (0) - (len(((_dafny.SeqWithoutIsStrInference([d_2071_v0_ for d_2150_i7_ in range(default__.abs(181))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([d_2071_v0_ for d_2151_i7_ in range(default__.abs(181))]))), d_2071_v0_)) + (d_2075_v4_))), len(d_2148_v70_))
            source40_ = d_2149_v71_
            if source40_.is_DC38:
                d_2152___mcc_h0_ = source40_.cf64
                d_2153_cf64_ = d_2152___mcc_h0_
                d_2154_v72_: int
                d_2154_v72_ = -855
                d_2154_v72_ = d_2154_v72_
                d_2071_v0_ = d_2071_v0_
                d_2072_v1_ = (d_2072_v1_).set(d_2154_v72_, p2)
                (self).f3 = (d_2148_v70_).ispropersubset(d_2148_v70_)
            elif source40_.is_DC39:
                d_2155___mcc_h1_ = source40_.cf65
                d_2156___mcc_h2_ = source40_.cf66
                d_2157___mcc_h3_ = source40_.cf67
                d_2158_cf67_ = d_2157___mcc_h3_
                d_2159_cf66_ = d_2156___mcc_h2_
                d_2160_cf65_ = d_2155___mcc_h1_
                d_2161_v73_: _dafny.Map
                d_2161_v73_ = _dafny.Map({(self).f5: self.f3})
                d_2162_v74_: bool
                d_2163_v75_: bool
                d_2164_v76_: _dafny.Seq
                d_2165_v77_: _dafny.Seq
                out54_: bool
                out55_: bool
                out56_: _dafny.Seq
                out57_: _dafny.Seq
                out54_, out55_, out56_, out57_ = default__.m0(d_2161_v73_, globalState)
                d_2162_v74_ = out54_
                d_2163_v75_ = out55_
                d_2164_v76_ = out56_
                d_2165_v77_ = out57_
                d_2162_v74_ = not(d_2163_v75_)
                (self).f3 = d_2162_v74_
                index342_ = default__.safeIndex(759, ((self).f5).length(0))
                ((self).f5)[index342_] = (d_2163_v75_ if not(d_2162_v74_) else self.f3)
                index343_ = default__.safeIndex(759, ((self).f5).length(0))
                ((self).f5)[index343_] = (True) and (d_2162_v74_)
            elif True:
                d_2166___mcc_h4_ = source40_.cf63
                d_2167_cf63_ = d_2166___mcc_h4_
                (self).f3 = self.f3
                d_2077_v5_ = d_2077_v5_
                d_2168_v78_: C1
                nw298_ = C1()
                nw298_.ctor__()
                d_2168_v78_ = nw298_
                d_2169_v79_: _dafny.Array
                nw299_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2169_v79_ = nw299_
                d_2170_v80_: _dafny.Map
                d_2170_v80_ = _dafny.Map({d_2169_v79_: (len(_dafny.SeqWithoutIsStrInference([(self).f4]))) > ((self).f4)})
                d_2171_v81_: _dafny.MultiSet
                d_2171_v81_ = _dafny.MultiSet([self.f3, self.f3, self.f3, self.f3, self.f3])
                d_2170_v80_ = (d_2170_v80_).set(d_2169_v79_, (self.f3) not in (d_2171_v81_))
            d_2172_v82_: D2
            d_2172_v82_ = D2_DC6(d_2071_v0_, self.f3, self.f3, len(_dafny.SeqWithoutIsStrInference([self.f3])))
            d_2172_v82_ = d_2172_v82_
            d_2173_v83_: _dafny.Array
            nw300_ = _dafny.Array(None, 26)
            d_2173_v83_ = nw300_
            d_2173_v83_ = d_2173_v83_
            d_2174_v84_: _dafny.Map
            d_2174_v84_ = _dafny.Map({default__.safeModuloInt((self).f4, len(d_2075_v4_)): d_2077_v5_})
            r0 = ((d_2174_v84_)[p2] if (p2) in (d_2174_v84_) else d_2077_v5_)
            index344_ = default__.safeIndex(969, ((self).f5).length(0))
            ((self).f5)[index344_] = ((self).f2) < (d_2075_v4_)
            index345_ = default__.safeIndex(969, ((self).f5).length(0))
            ((self).f5)[index345_] = (default__.safeModuloInt(p2, p2)) <= (p2)
        d_2175_v85_: _dafny.Map
        d_2175_v85_ = _dafny.Map({d_2077_v5_: self.f3})
        index346_ = default__.safeIndex(241, ((self).f5).length(0))
        ((self).f5)[index346_] = ((d_2175_v85_)[d_2077_v5_] if (d_2077_v5_) in (d_2175_v85_) else self.f3)
        d_2176_v86_: D6
        d_2176_v86_ = D6_DC17(self.f3)
        index347_ = default__.safeIndex(241, ((self).f5).length(0))
        ((self).f5)[index347_] = (d_2176_v86_).cf28
        d_2177_v87_: _dafny.Seq
        d_2177_v87_ = _dafny.SeqWithoutIsStrInference([len(d_2075_v4_)])
        d_2178_v88_: _dafny.MultiSet
        d_2178_v88_ = _dafny.MultiSet([not((len(d_2177_v87_)) < ((self).f4)), self.f3])
        d_2179_v89_: _dafny.Array
        nw301_ = _dafny.Array(_dafny.Array(None, 0), 25)
        d_2179_v89_ = nw301_
        d_2180_v90_: _dafny.Array
        def lambda85_(d_2181_v0_):
            def lambda86_(d_2182_i8_):
                return d_2181_v0_

            return lambda86_

        init48_ = lambda85_(d_2071_v0_)
        nw302_ = _dafny.Array(None, 2)
        for i0_48_ in range(nw302_.length(0)):
            nw302_[i0_48_] = init48_(i0_48_)
        d_2180_v90_ = nw302_
        index348_ = default__.safeIndex(950, (d_2179_v89_).length(0))
        (d_2179_v89_)[index348_] = (d_2180_v90_ if ((self).f5)[default__.safeIndex(241, ((self).f5).length(0))] else d_2180_v90_)
        index349_ = default__.safeIndex(241, ((self).f5).length(0))
        index350_ = default__.safeIndex(950, (d_2179_v89_).length(0))
        rhs368_ = (self).fm13(globalState)
        rhs369_ = d_2178_v88_
        rhs370_ = d_2180_v90_
        rhs371_ = self.f3
        lhs236_ = (self).f5
        lhs237_ = default__.safeIndex(241, ((self).f5).length(0))
        lhs238_ = d_2179_v89_
        lhs239_ = default__.safeIndex(950, (d_2179_v89_).length(0))
        lhs240_ = self
        lhs236_[lhs237_] = rhs368_
        d_2178_v88_ = rhs369_
        lhs238_[lhs239_] = rhs370_
        lhs240_.f3 = rhs371_
        r0 = d_2077_v5_
        return r0

    @property
    def f10(self):
        return self._f10

class C9(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Array = _dafny.Array(None, 0)
        self._f9: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f8, f9, f4, f5, f3, f1, f2):
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return not(((_dafny.Set({_dafny.Set({True, self.f3})})) | (_dafny.Set({_dafny.Set({self.f3, False, self.f3, self.f3, self.f3}), _dafny.Set({True})}))).isdisjoint(_dafny.Set({_dafny.Set({self.f3}), _dafny.Set({False})})))

    def fm14(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2183_i0_ in range(default__.abs(-895))])

    def fm12(self, p0, p1, p2, p3, globalState):
        return default__.safeModuloInt((0) - (((self).f4) + ((self).f4)), (self).f4)

    def fm11(self, p0, globalState):
        return self.f3

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        d_2184_v0_: _dafny.Set
        d_2184_v0_ = _dafny.Set({len((self).f2)})
        d_2185_i0_: int
        d_2185_i0_ = 0
        with _dafny.label("12"):
            while (d_2184_v0_).isdisjoint((d_2184_v0_) | (d_2184_v0_)):
                with _dafny.c_label("12"):
                    if (d_2185_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_2185_i0_ = (d_2185_i0_) + (1)
                    d_2186_v1_: _dafny.Map
                    d_2186_v1_ = _dafny.Map({self.f3: True})
                    (self).f3 = ((p1) - (322)) == ((len((self).f2)) - (len(d_2186_v1_)))
                    index351_ = default__.safeIndex(587, ((self).f5).length(0))
                    ((self).f5)[index351_] = self.f3
                    index352_ = default__.safeIndex(587, ((self).f5).length(0))
                    ((self).f5)[index352_] = self.f3
                    r2 = ((self).f4 if self.f3 else (p0) * (p1))
                    d_2187_v2_: _dafny.Array
                    nw303_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                    d_2187_v2_ = nw303_
                    d_2188_v3_: _dafny.Map
                    d_2188_v3_ = _dafny.Map({d_2187_v2_: (self).fm11(True, globalState)})
                    d_2188_v3_ = (d_2188_v3_).set(d_2187_v2_, (p0) in (d_2184_v0_))
                    pass
            pass
        if (self).fm11((p0) > ((self).f4), globalState):
            d_2189_v4_: _dafny.Array
            nw304_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_2189_v4_ = nw304_
            rhs372_ = p0
            rhs373_ = d_2189_v4_
            r2 = rhs372_
            d_2189_v4_ = rhs373_
            r2 = p0
            d_2190_v5_: _dafny.Seq
            d_2190_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqtdfit"))
            d_2190_v5_ = ((default__.fm2(self.f3, globalState)) + ((self).f2)) + (d_2190_v5_)
            d_2191_v6_: _dafny.MultiSet
            d_2191_v6_ = _dafny.MultiSet([838, (self).f4, len((self).f2)])
            d_2192_v7_: D4
            d_2192_v7_ = D4_DC12(d_2191_v6_)
            source41_ = d_2192_v7_
            if source41_.is_DC13:
                d_2193___mcc_h0_ = source41_.cf21
                d_2194___mcc_h1_ = source41_.cf22
                d_2195_cf22_ = d_2194___mcc_h1_
                d_2196_cf21_ = d_2193___mcc_h0_
                d_2197_v8_: _dafny.Map
                d_2197_v8_ = _dafny.Map({-162: p0})
                d_2198_v9_: _dafny.Map
                d_2198_v9_ = _dafny.Map({(self).f8: d_2190_v5_})
                d_2199_v10_: _dafny.Seq
                d_2199_v10_ = _dafny.SeqWithoutIsStrInference([d_2196_cf21_])
                d_2200_v12_: _dafny.Map
                d_2200_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krvvypne")): p0})
                d_2201_v13_: _dafny.Array
                nw305_ = _dafny.Array(None, 23)
                nw305_[int(0)] = p0
                nw305_[int(1)] = p0
                nw305_[int(2)] = (self).f4
                nw305_[int(3)] = p1
                nw305_[int(4)] = p0
                nw305_[int(5)] = (self).f4
                nw305_[int(6)] = p0
                nw305_[int(7)] = p1
                nw305_[int(8)] = p0
                nw305_[int(9)] = p0
                nw305_[int(10)] = (self).f4
                nw305_[int(11)] = (_dafny.MultiSet([p0, p1, p0, p1, (self).f4])).cardinality
                nw305_[int(12)] = (0) - (p0)
                nw305_[int(13)] = ((d_2197_v8_)[(self).f4] if ((self).f4) in (d_2197_v8_) else p1)
                nw305_[int(14)] = p1
                nw305_[int(15)] = p1
                nw305_[int(16)] = len(d_2198_v9_)
                nw305_[int(17)] = (self).f4
                nw305_[int(18)] = p0
                nw305_[int(19)] = len(d_2199_v10_)
                nw305_[int(20)] = len(d_2197_v8_)
                nw305_[int(21)] = p0
                def iife195_():
                    coll63_ = _dafny.Map()
                    compr_63_: _dafny.Seq
                    for compr_63_ in (d_2200_v12_).keys.Elements:
                        d_2202_v11_: _dafny.Seq = compr_63_
                        if (d_2202_v11_) in (d_2200_v12_):
                            coll63_[d_2202_v11_] = d_2196_cf21_
                    return _dafny.Map(coll63_)
                nw305_[int(22)] = (0) - ((0) - (len(iife195_()
                )))
                d_2201_v13_ = nw305_
                d_2203_v14_: D5
                d_2203_v14_ = D5_DC14(d_2201_v13_)
                pat_let_tv52_ = d_2201_v13_
                d_2204_v15_: _dafny.Seq
                def iife196_(_pat_let66_0):
                    def iife197_(d_2205_dt__update__tmp_h0_):
                        def iife198_(_pat_let67_0):
                            def iife199_(d_2206_dt__update_hcf23_h0_):
                                return D5_DC14(d_2206_dt__update_hcf23_h0_)
                            return iife199_(_pat_let67_0)
                        return iife198_(pat_let_tv52_)
                    return iife197_(_pat_let66_0)
                d_2204_v15_ = _dafny.SeqWithoutIsStrInference([d_2201_v13_, d_2201_v13_, d_2201_v13_, (iife196_(d_2203_v14_)).cf23])
                d_2207_v16_: _dafny.Map
                d_2207_v16_ = _dafny.Map({False: len(d_2190_v5_)})
                d_2208_v17_: D2
                d_2208_v17_ = D2_DC5(d_2207_v16_)
                d_2209_v18_: _dafny.Set
                d_2209_v18_ = _dafny.Set({d_2208_v17_})
                d_2210_v19_: D1
                d_2210_v19_ = D1_DC3(d_2196_cf21_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvsmecttx"))), len(d_2209_v18_))
                d_2211_v20_: _dafny.Map
                d_2211_v20_ = _dafny.Map({(0) - (default__.safeModuloInt((self).f4, len(d_2190_v5_))): (d_2210_v19_).cf3})
                index353_ = default__.safeIndex(284, (d_2201_v13_).length(0))
                (d_2201_v13_)[index353_] = len(d_2195_cf22_)
                index354_ = default__.safeIndex(284, (d_2201_v13_).length(0))
                def iife200_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(608, 28):
                        d_2212_v21_: int = compr_64_
                        if ((608) <= (d_2212_v21_)) and ((d_2212_v21_) < (28)):
                            coll64_[(d_2212_v21_) + (((d_2207_v16_)[self.f3] if (self.f3) in (d_2207_v16_) else (self).f4))] = self.f3
                    return _dafny.Map(coll64_)
                rhs374_ = (d_2204_v15_).set(default__.safeIndex((0) - ((len((iife200_()
                ).set((self).f4, default__.fm3(len(d_2184_v0_), globalState)))) * (p0)), len(d_2204_v15_)), d_2201_v13_)
                rhs375_ = d_2211_v20_
                rhs376_ = default__.fm0(globalState)
                lhs241_ = d_2201_v13_
                lhs242_ = default__.safeIndex(284, (d_2201_v13_).length(0))
                d_2204_v15_ = rhs374_
                d_2211_v20_ = rhs375_
                lhs241_[lhs242_] = rhs376_
                d_2213_v22_: _dafny.MultiSet
                d_2213_v22_ = _dafny.MultiSet([d_2196_cf21_, not(d_2196_cf21_), d_2196_cf21_])
                d_2214_v23_: _dafny.Map
                d_2214_v23_ = _dafny.Map({not(not (d_2196_cf21_) or (self.f3)): d_2213_v22_})
                d_2214_v23_ = d_2214_v23_
                d_2215_v24_: bool
                d_2216_v25_: bool
                out58_: bool
                out59_: bool
                out58_, out59_ = (self).m5(globalState)
                d_2215_v24_ = out58_
                d_2216_v25_ = out59_
                d_2217_v26_: C0
                nw306_ = C0()
                nw306_.ctor__(d_2201_v13_)
                d_2217_v26_ = nw306_
            elif True:
                d_2218___mcc_h2_ = source41_.cf20
                d_2219_cf20_ = d_2218___mcc_h2_
                d_2220_v27_: _dafny.Map
                d_2220_v27_ = _dafny.Map({(self).f4: d_2190_v5_})
                d_2221_v28_: _dafny.Set
                d_2221_v28_ = _dafny.Set({self.f3, self.f3})
                d_2222_v29_: _dafny.Seq
                d_2222_v29_ = _dafny.SeqWithoutIsStrInference([(self).fm13(globalState), self.f3])
                d_2223_v30_: T0
                nw307_ = C5()
                nw307_.ctor__(True, (self).f1, (self).f2)
                d_2223_v30_ = nw307_
                d_2224_v31_: _dafny.MultiSet
                d_2224_v31_ = _dafny.MultiSet([self.f3])
                d_2225_v32_: _dafny.Map
                d_2225_v32_ = _dafny.Map({d_2223_v30_: (d_2224_v31_).cardinality})
                d_2226_v33_: _dafny.Array
                nw308_ = _dafny.Array(None, 15)
                nw308_[int(0)] = p0
                nw308_[int(1)] = p1
                nw308_[int(2)] = p0
                nw308_[int(3)] = p1
                nw308_[int(4)] = len(d_2220_v27_)
                nw308_[int(5)] = 451
                nw308_[int(6)] = len(d_2221_v28_)
                nw308_[int(7)] = p0
                nw308_[int(8)] = (0) - (p1)
                nw308_[int(9)] = (self).f4
                nw308_[int(10)] = (d_2219_cf20_).cardinality
                nw308_[int(11)] = len(_dafny.SeqWithoutIsStrInference([(self).f4 for d_2227_i1_ in range(default__.abs(79))]))
                nw308_[int(12)] = len(_dafny.SeqWithoutIsStrInference([(self).f4 for d_2228_i2_ in range(default__.abs(308))]))
                nw308_[int(13)] = len(d_2222_v29_)
                nw308_[int(14)] = len((d_2225_v32_).set(d_2223_v30_, p1))
                d_2226_v33_ = nw308_
                d_2229_v34_: C0
                nw309_ = C0()
                nw309_.ctor__(d_2226_v33_)
                d_2229_v34_ = nw309_
                (self).f3 = not ((self.f3 if self.f3 else True)) or (((self).f4) > (default__.fm0(globalState)))
                index355_ = default__.safeIndex(388, (d_2226_v33_).length(0))
                (d_2226_v33_)[index355_] = (self).f4
                index356_ = default__.safeIndex(388, (d_2226_v33_).length(0))
                (d_2226_v33_)[index356_] = 942
                d_2230_v35_: _dafny.Map
                d_2230_v35_ = _dafny.Map({self.f3: d_2189_v4_})
                d_2230_v35_ = (d_2230_v35_).set((len(default__.fm2(self.f3, globalState))) >= (len(_dafny.SeqWithoutIsStrInference([p1, p0]))), d_2189_v4_)
            d_2231_v36_: _dafny.Array
            nw310_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_2231_v36_ = nw310_
            d_2232_v37_: _dafny.Seq
            d_2232_v37_ = _dafny.SeqWithoutIsStrInference([d_2231_v36_, d_2231_v36_])
            d_2233_v38_: str
            d_2233_v38_ = _dafny.CodePoint('x')
            d_2234_v39_: _dafny.Map
            d_2234_v39_ = _dafny.Map({d_2233_v38_: self.f3})
            d_2235_v41_: _dafny.Array
            nw311_ = _dafny.Array(None, 5)
            nw311_[int(0)] = d_2231_v36_
            nw311_[int(1)] = d_2231_v36_
            nw311_[int(2)] = d_2231_v36_
            def iife201_():
                coll65_ = _dafny.Set()
                compr_65_: str
                for compr_65_ in ((d_2234_v39_).set(_dafny.CodePoint('v'), self.f3)).keys.Elements:
                    d_2236_v40_: str = compr_65_
                    if (d_2236_v40_) in ((d_2234_v39_).set(_dafny.CodePoint('v'), self.f3)):
                        coll65_ = coll65_.union(_dafny.Set([d_2236_v40_]))
                return _dafny.Set(coll65_)
            nw311_[int(3)] = (d_2232_v37_)[default__.safeIndex(len(iife201_()
            ), len(d_2232_v37_))]
            nw311_[int(4)] = d_2231_v36_
            d_2235_v41_ = nw311_
            index357_ = default__.safeIndex(238, (d_2235_v41_).length(0))
            (d_2235_v41_)[index357_] = d_2231_v36_
            index358_ = default__.safeIndex(196, ((self).f5).length(0))
            ((self).f5)[index358_] = self.f3
            index359_ = default__.safeIndex(695, ((self).f8).length(0))
            ((self).f8)[index359_] = self.f3
            d_2237_v42_: _dafny.Seq
            d_2237_v42_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
            index360_ = default__.safeIndex(238, (d_2235_v41_).length(0))
            index361_ = default__.safeIndex(196, ((self).f5).length(0))
            index362_ = default__.safeIndex(695, ((self).f8).length(0))
            rhs377_ = d_2231_v36_
            rhs378_ = self.f3
            rhs379_ = self.f3
            rhs380_ = (d_2237_v42_)[default__.safeIndex(((d_2191_v6_).set(245, default__.abs(p0))).cardinality, len(d_2237_v42_))]
            rhs381_ = (p0) - ((p1) + ((self).f4))
            lhs243_ = d_2235_v41_
            lhs244_ = default__.safeIndex(238, (d_2235_v41_).length(0))
            lhs245_ = (self).f5
            lhs246_ = default__.safeIndex(196, ((self).f5).length(0))
            lhs247_ = (self).f8
            lhs248_ = default__.safeIndex(695, ((self).f8).length(0))
            lhs249_ = self
            lhs243_[lhs244_] = rhs377_
            lhs245_[lhs246_] = rhs378_
            lhs247_[lhs248_] = rhs379_
            lhs249_.f3 = rhs380_
            r1 = rhs381_
        elif True:
            r1 = ((0) - ((self).f4)) - ((self).f4)
            d_2238_v43_: _dafny.Array
            nw312_ = _dafny.Array(int(0), 22)
            d_2238_v43_ = nw312_
            index363_ = default__.safeIndex(648, (d_2238_v43_).length(0))
            (d_2238_v43_)[index363_] = (self).f4
            d_2239_v44_: _dafny.Set
            d_2239_v44_ = _dafny.Set({True})
            index364_ = default__.safeIndex(648, (d_2238_v43_).length(0))
            (d_2238_v43_)[index364_] = (default__.safeDivisionInt(p1, p1)) * (len(d_2239_v44_))
            d_2240_v45_: str
            d_2240_v45_ = _dafny.CodePoint('m')
            d_2241_v46_: _dafny.Map
            d_2241_v46_ = _dafny.Map({self.f3: self.f3})
            d_2242_v47_: D8
            d_2242_v47_ = D8_DC24(d_2240_v45_, (d_2241_v46_) | (_dafny.Map({self.f3: self.f3})))
            source42_ = d_2242_v47_
            if source42_.is_DC23:
                d_2243___mcc_h3_ = source42_.cf38
                d_2244___mcc_h4_ = source42_.cf39
                d_2245___mcc_h5_ = source42_.cf40
                d_2246_cf40_ = d_2245___mcc_h5_
                d_2247_cf39_ = d_2244___mcc_h4_
                d_2248_cf38_ = d_2243___mcc_h3_
                d_2249_v48_: _dafny.Map
                d_2249_v48_ = _dafny.Map({-766: self.f3})
                d_2246_cf40_ = default__.fm3(default__.safeDivisionInt(len(d_2249_v48_), p0), globalState)
                d_2250_v49_: bool
                d_2251_v50_: bool
                out60_: bool
                out61_: bool
                out60_, out61_ = (self).m5(globalState)
                d_2250_v49_ = out60_
                d_2251_v50_ = out61_
                d_2251_v50_ = not(d_2250_v49_)
                d_2252_v51_: _dafny.Map
                d_2252_v51_ = _dafny.Map({d_2238_v43_: ((self).f4 if d_2250_v49_ else (self).f4)})
                d_2253_v52_: _dafny.Seq
                d_2253_v52_ = _dafny.SeqWithoutIsStrInference([d_2238_v43_])
                d_2252_v51_ = (d_2252_v51_).set((d_2253_v52_)[default__.safeIndex((self).f4, len(d_2253_v52_))], (0) - (p0))
            elif source42_.is_DC24:
                d_2254___mcc_h6_ = source42_.cf41
                d_2255___mcc_h7_ = source42_.cf42
                d_2256_cf42_ = d_2255___mcc_h7_
                d_2257_cf41_ = d_2254___mcc_h6_
                d_2258_v53_: _dafny.Seq
                d_2258_v53_ = _dafny.SeqWithoutIsStrInference([p0, (self).f4, p1])
                d_2258_v53_ = (d_2258_v53_) + (d_2258_v53_)
                index365_ = default__.safeIndex(246, ((self).f5).length(0))
                ((self).f5)[index365_] = False
                index366_ = default__.safeIndex(246, ((self).f5).length(0))
                ((self).f5)[index366_] = default__.fm3((self).f4, globalState)
                d_2259_v54_: _dafny.Map
                d_2259_v54_ = _dafny.Map({((self).f5)[default__.safeIndex(246, ((self).f5).length(0))]: (self).f2})
                index367_ = default__.safeIndex(246, ((self).f5).length(0))
                index368_ = default__.safeIndex(648, (d_2238_v43_).length(0))
                index369_ = default__.safeIndex(246, ((self).f5).length(0))
                rhs382_ = not (self.f3) or ((d_2239_v44_).issubset(d_2239_v44_))
                rhs383_ = default__.fm22(((self).f5)[default__.safeIndex(246, ((self).f5).length(0))], (len(((d_2259_v54_)[True] if (True) in (d_2259_v54_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))))) + ((self).f4), globalState)
                rhs384_ = default__.fm3(len(_dafny.Map({d_2257_cf41_: (self).f4})), globalState)
                rhs385_ = (d_2238_v43_)[default__.safeIndex(648, (d_2238_v43_).length(0))]
                rhs386_ = (False if ((self).f5)[default__.safeIndex(246, ((self).f5).length(0))] else ((self).f5)[default__.safeIndex(246, ((self).f5).length(0))])
                lhs250_ = self
                lhs251_ = (self).f5
                lhs252_ = default__.safeIndex(246, ((self).f5).length(0))
                lhs253_ = d_2238_v43_
                lhs254_ = default__.safeIndex(648, (d_2238_v43_).length(0))
                lhs255_ = (self).f5
                lhs256_ = default__.safeIndex(246, ((self).f5).length(0))
                lhs250_.f3 = rhs382_
                d_2257_cf41_ = rhs383_
                lhs251_[lhs252_] = rhs384_
                lhs253_[lhs254_] = rhs385_
                lhs255_[lhs256_] = rhs386_
                d_2260_v55_: D9
                d_2260_v55_ = D9_DC27(-385)
                d_2261_v56_: _dafny.Seq
                d_2261_v56_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3, default__.fm3((d_2238_v43_)[default__.safeIndex(648, (d_2238_v43_).length(0))], globalState)])
                d_2262_v57_: _dafny.Seq
                d_2262_v57_ = _dafny.SeqWithoutIsStrInference([d_2261_v56_, _dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(246, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(246, ((self).f5).length(0))]]), d_2261_v56_, d_2261_v56_, d_2261_v56_])
                d_2263_v58_: T2
                nw313_ = C8()
                nw313_.ctor__(d_2262_v57_, self.f3, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imqca")), (self).f4, (self).f8)
                d_2263_v58_ = nw313_
                d_2264_v59_: _dafny.Map
                d_2264_v59_ = _dafny.Map({d_2260_v55_: d_2263_v58_})
                d_2264_v59_ = (d_2264_v59_).set(d_2260_v55_, d_2263_v58_)
            elif True:
                d_2265___mcc_h8_ = source42_.cf37
                d_2266_cf37_ = d_2265___mcc_h8_
                d_2267_v60_: _dafny.Seq
                d_2267_v60_ = _dafny.SeqWithoutIsStrInference([len(d_2241_v46_), p0])
                d_2241_v46_ = (d_2241_v46_).set(self.f3, (d_2184_v0_).issubset(_dafny.Set({p0, (d_2266_cf37_).cardinality, p1, len(d_2267_v60_), p0})))
                d_2238_v43_ = d_2238_v43_
                r2 = (default__.safeDivisionInt(p1, (self).f4)) - (default__.safeModuloInt((d_2238_v43_)[default__.safeIndex(648, (d_2238_v43_).length(0))], p1))
                index370_ = default__.safeIndex(648, (d_2238_v43_).length(0))
                (d_2238_v43_)[index370_] = default__.fm0(globalState)
            r2 = (self).f4
            index371_ = default__.safeIndex(648, (d_2238_v43_).length(0))
            (d_2238_v43_)[index371_] = (d_2238_v43_)[default__.safeIndex(648, (d_2238_v43_).length(0))]
        d_2268_i3_: int
        d_2268_i3_ = 0
        with _dafny.label("13"):
            while (p0) > (p1):
                with _dafny.c_label("13"):
                    if (d_2268_i3_) >= (100):
                        raise _dafny.Break("13")
                    d_2268_i3_ = (d_2268_i3_) + (1)
                    (self).f3 = self.f3
                    d_2269_v61_: _dafny.Array
                    nw314_ = _dafny.Array(int(0), 4)
                    d_2269_v61_ = nw314_
                    d_2270_v62_: _dafny.Seq
                    d_2270_v62_ = _dafny.SeqWithoutIsStrInference([self.f3])
                    rhs387_ = not(True)
                    rhs388_ = (self).f4
                    rhs389_ = ((d_2270_v62_).set(default__.safeIndex((self).f4, len(d_2270_v62_)), True)) <= ((d_2270_v62_) + (default__.fm7(self.f3, self.f3, globalState)))
                    rhs390_ = d_2269_v61_
                    lhs257_ = self
                    lhs258_ = self
                    lhs257_.f3 = rhs387_
                    r2 = rhs388_
                    lhs258_.f3 = rhs389_
                    d_2269_v61_ = rhs390_
                    d_2271_v63_: _dafny.Seq
                    d_2271_v63_ = _dafny.SeqWithoutIsStrInference([(self).f2, (self).f2])
                    d_2272_v64_: _dafny.Map
                    d_2272_v64_ = _dafny.Map({p0: _dafny.Map({(d_2271_v63_)[default__.safeIndex(p1, len(d_2271_v63_))]: self.f3})})
                    d_2273_v65_: C8
                    nw315_ = C8()
                    nw315_.ctor__(_dafny.SeqWithoutIsStrInference([d_2270_v62_, default__.fm7(self.f3, (default__.fm10(True, self.f3, self.f3, globalState)).cf11, globalState)]), self.f3, ((d_2272_v64_)[p0] if (p0) in (d_2272_v64_) else (self).f1), ((self).f2 if self.f3 else (self).f2), 451, (self).f8)
                    d_2273_v65_ = nw315_
                    d_2273_v65_ = d_2273_v65_
                    d_2274_v66_: _dafny.Map
                    d_2274_v66_ = _dafny.Map({self.f3: (self).f4})
                    d_2275_v67_: D2
                    d_2275_v67_ = D2_DC5(d_2274_v66_)
                    d_2276_v68_: _dafny.Map
                    d_2276_v68_ = _dafny.Map({d_2274_v66_: d_2275_v67_})
                    d_2276_v68_ = (d_2276_v68_).set(d_2274_v66_, default__.fm55(self.f3, (self).f2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfxldc")), globalState))
                    pass
            pass
        if self.f3:
            d_2277_v69_: _dafny.MultiSet
            d_2277_v69_ = _dafny.MultiSet([p1, default__.fm0(globalState)])
            d_2277_v69_ = d_2277_v69_
            index372_ = default__.safeIndex(15, ((self).f8).length(0))
            ((self).f8)[index372_] = self.f3
            index373_ = default__.safeIndex(15, ((self).f8).length(0))
            ((self).f8)[index373_] = self.f3
            d_2278_v70_: _dafny.Seq
            d_2278_v70_ = _dafny.SeqWithoutIsStrInference([p1])
            if not((d_2278_v70_) <= (d_2278_v70_)):
                d_2279_v71_: C5
                nw316_ = C5()
                nw316_.ctor__(((self).f8)[default__.safeIndex(15, ((self).f8).length(0))], (self).f1, ((self).f2) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlopb"))))
                d_2279_v71_ = nw316_
                d_2280_v72_: str
                d_2280_v72_ = _dafny.CodePoint('e')
                index374_ = default__.safeIndex(15, ((self).f8).length(0))
                rhs391_ = (d_2279_v71_).f15
                rhs392_ = d_2280_v72_
                rhs393_ = (d_2279_v71_).fm11(False, globalState)
                lhs259_ = (self).f8
                lhs260_ = default__.safeIndex(15, ((self).f8).length(0))
                lhs261_ = self
                lhs259_[lhs260_] = rhs391_
                d_2280_v72_ = rhs392_
                lhs261_.f3 = rhs393_
                r1 = (0) - (((0) - ((0) - (default__.safeModuloInt(p0, p0))) if not(not(self.f3)) else default__.safeDivisionInt((self).f4, p1)))
                d_2281_v73_: _dafny.Array
                def lambda87_(d_2282_v71_):
                    def lambda88_(d_2283_i4_):
                        return not((_dafny.Set({((self).f8)[default__.safeIndex(15, ((self).f8).length(0))], (d_2282_v71_).f15})).ispropersubset(_dafny.Set({((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]})))

                    return lambda88_

                init49_ = lambda87_(d_2279_v71_)
                nw317_ = _dafny.Array(None, 4)
                for i0_49_ in range(nw317_.length(0)):
                    nw317_[i0_49_] = init49_(i0_49_)
                d_2281_v73_ = nw317_
                d_2284_v74_: _dafny.Array
                def lambda89_(d_2285_p1_):
                    def lambda90_(d_2286_i5_):
                        return (d_2286_i5_) * ((D8_DC23(392, d_2285_p1_, self.f3)).cf38)

                    return lambda90_

                init50_ = lambda89_(p1)
                nw318_ = _dafny.Array(None, 29)
                for i0_50_ in range(nw318_.length(0)):
                    nw318_[i0_50_] = init50_(i0_50_)
                d_2284_v74_ = nw318_
                d_2287_v75_: D5
                d_2287_v75_ = D5_DC14(d_2284_v74_)
                d_2288_v76_: _dafny.Map
                d_2288_v76_ = _dafny.Map({(d_2284_v74_ if ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))] else (d_2287_v75_).cf23): default__.safeDivisionInt((self).f4, p0)})
                d_2289_v77_: _dafny.Map
                d_2289_v77_ = _dafny.Map({((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]: len(d_2184_v0_)})
                d_2290_v78_: _dafny.Map
                d_2290_v78_ = _dafny.Map({p1: self.f3})
                d_2291_v79_: _dafny.Map
                d_2291_v79_ = _dafny.Map({len(d_2184_v0_): ((d_2290_v78_)[p0] if (p0) in (d_2290_v78_) else (d_2279_v71_).f15)})
                index375_ = default__.safeIndex(15, ((self).f8).length(0))
                index376_ = default__.safeIndex(15, ((self).f8).length(0))
                rhs394_ = True
                rhs395_ = (d_2279_v71_).f15
                rhs396_ = (self).f5
                rhs397_ = (d_2288_v76_).set(d_2284_v74_, default__.safeModuloInt(len(d_2289_v77_), len((d_2291_v79_).set(p1, ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]))))
                rhs398_ = (0) - (-889)
                lhs262_ = (self).f8
                lhs263_ = default__.safeIndex(15, ((self).f8).length(0))
                lhs264_ = (self).f8
                lhs265_ = default__.safeIndex(15, ((self).f8).length(0))
                lhs262_[lhs263_] = rhs394_
                lhs264_[lhs265_] = rhs395_
                d_2281_v73_ = rhs396_
                d_2288_v76_ = rhs397_
                r2 = rhs398_
                d_2292_v80_: _dafny.Seq
                d_2292_v80_ = _dafny.SeqWithoutIsStrInference([((self).f8)[default__.safeIndex(15, ((self).f8).length(0))], ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]])
                d_2293_v81_: _dafny.Seq
                d_2293_v81_ = _dafny.SeqWithoutIsStrInference([d_2292_v80_])
                d_2294_v82_: D16
                d_2294_v82_ = D16_DC41(d_2293_v81_)
                d_2295_v83_: D5
                d_2295_v83_ = D5_DC15(p1, False, (self).f5)
                d_2296_v84_: _dafny.Map
                d_2296_v84_ = _dafny.Map({(d_2279_v71_).f15: (d_2295_v83_).cf26})
                d_2297_v85_: C8
                nw319_ = C8()
                nw319_.ctor__((d_2294_v82_).cf69, not (self.f3) or (self.f3), ((self).f1) | ((self).f1), ((self).f2).set(default__.safeIndex(290, len((self).f2)), d_2280_v72_), p0, ((d_2296_v84_)[self.f3] if (self.f3) in (d_2296_v84_) else d_2281_v73_))
                d_2297_v85_ = nw319_
            elif True:
                d_2298_v86_: _dafny.Seq
                d_2298_v86_ = _dafny.SeqWithoutIsStrInference([default__.fm3(853, globalState), ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]])
                d_2299_v87_: str
                d_2299_v87_ = _dafny.CodePoint('o')
                d_2300_v88_: _dafny.Map
                d_2300_v88_ = _dafny.Map({(self).f5: ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]})
                d_2301_v89_: _dafny.Map
                d_2301_v89_ = _dafny.Map({d_2299_v87_: len(d_2300_v88_)})
                d_2302_v90_: _dafny.MultiSet
                d_2302_v90_ = _dafny.MultiSet([((self).f8)[default__.safeIndex(15, ((self).f8).length(0))], self.f3])
                d_2303_v91_: D8
                d_2303_v91_ = D8_DC22(d_2302_v90_)
                d_2304_v92_: _dafny.Map
                d_2304_v92_ = _dafny.Map({p0: self.f3})
                d_2305_v93_: _dafny.Map
                d_2305_v93_ = _dafny.Map({d_2299_v87_: self.f3})
                d_2306_v94_: _dafny.Array
                nw320_ = _dafny.Array(None, 28)
                nw320_[int(0)] = (0) - (len((self).f2))
                nw320_[int(1)] = p1
                nw320_[int(2)] = p1
                nw320_[int(3)] = p0
                nw320_[int(4)] = len(d_2304_v92_)
                nw320_[int(5)] = (self).f4
                nw320_[int(6)] = p0
                nw320_[int(7)] = (self).f4
                nw320_[int(8)] = (self).f4
                nw320_[int(9)] = p1
                nw320_[int(10)] = p0
                nw320_[int(11)] = p1
                nw320_[int(12)] = 477
                nw320_[int(13)] = p0
                nw320_[int(14)] = (self).f4
                nw320_[int(15)] = (0) - (len(d_2278_v70_))
                nw320_[int(16)] = (self).f4
                nw320_[int(17)] = p0
                nw320_[int(18)] = (self).f4
                nw320_[int(19)] = -705
                nw320_[int(20)] = len(d_2305_v93_)
                nw320_[int(21)] = p1
                nw320_[int(22)] = p0
                nw320_[int(23)] = (0) - (p0)
                nw320_[int(24)] = p1
                nw320_[int(25)] = (self).f4
                nw320_[int(26)] = len(d_2304_v92_)
                nw320_[int(27)] = ((d_2277_v69_)[len((self).f2)] if (len((self).f2)) in (d_2277_v69_) else (0) - (p1))
                d_2306_v94_ = nw320_
                d_2307_v95_: _dafny.Map
                d_2307_v95_ = _dafny.Map({False: d_2306_v94_})
                d_2308_v96_: _dafny.Array
                nw321_ = _dafny.Array(None, 26)
                nw321_[int(0)] = (self).fm13(globalState)
                nw321_[int(1)] = not(not(self.f3))
                nw321_[int(2)] = (p1) == (p1)
                nw321_[int(3)] = ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]
                nw321_[int(4)] = self.f3
                nw321_[int(5)] = (self.f3 if (d_2298_v86_)[default__.safeIndex(p0, len(d_2298_v86_))] else ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))])
                nw321_[int(6)] = self.f3
                nw321_[int(7)] = not (self.f3) or (((self).f8)[default__.safeIndex(15, ((self).f8).length(0))])
                nw321_[int(8)] = ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]
                nw321_[int(9)] = not((d_2278_v70_) <= (_dafny.SeqWithoutIsStrInference([((d_2301_v89_)[d_2299_v87_] if (d_2299_v87_) in (d_2301_v89_) else (self).f4), p0, 109, (((d_2303_v91_).cf37).set(self.f3, default__.abs(p0))).cardinality])))
                nw321_[int(10)] = self.f3
                nw321_[int(11)] = not(not (((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]) or (self.f3))
                nw321_[int(12)] = (((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]) or (False)
                nw321_[int(13)] = ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]
                nw321_[int(14)] = (p0) < (-29)
                nw321_[int(15)] = (909) == (len(d_2307_v95_))
                nw321_[int(16)] = (((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]) == (self.f3)
                nw321_[int(17)] = self.f3
                nw321_[int(18)] = self.f3
                nw321_[int(19)] = self.f3
                nw321_[int(20)] = self.f3
                nw321_[int(21)] = not(((self).f8)[default__.safeIndex(15, ((self).f8).length(0))])
                nw321_[int(22)] = self.f3
                nw321_[int(23)] = default__.fm3(len(((self).f2).set(default__.safeIndex((0) - (-803), len((self).f2)), d_2299_v87_)), globalState)
                nw321_[int(24)] = not(self.f3)
                nw321_[int(25)] = self.f3
                d_2308_v96_ = nw321_
                d_2308_v96_ = d_2308_v96_
                d_2309_v97_: _dafny.Map
                d_2309_v97_ = _dafny.Map({(self).f4: D8_DC23(p1, p0, not(((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]))})
                d_2310_v98_: D8
                d_2310_v98_ = D8_DC23((self).f4, p1, ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))])
                d_2309_v97_ = (d_2309_v97_).set((self).f4, d_2310_v98_)
                d_2311_v99_: _dafny.Seq
                d_2311_v99_ = _dafny.SeqWithoutIsStrInference([((self).f2).set(default__.safeIndex(337, len((self).f2)), d_2299_v87_)])
                d_2312_v100_: _dafny.Map
                d_2312_v100_ = _dafny.Map({d_2311_v99_: p1})
                d_2312_v100_ = (d_2312_v100_).set(d_2311_v99_, 524)
                d_2313_v101_: _dafny.Map
                d_2313_v101_ = _dafny.Map({((self).f2) + ((self).f2): p1})
                d_2313_v101_ = (d_2313_v101_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2314_i6_ in range(default__.abs(969))]), (0) - (p1))
                (self).f3 = ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]
            index377_ = default__.safeIndex(15, ((self).f8).length(0))
            ((self).f8)[index377_] = ((self).f8)[default__.safeIndex(15, ((self).f8).length(0))]
            d_2315_v102_: _dafny.Seq
            d_2315_v102_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5])
            d_2315_v102_ = (d_2315_v102_) + ((d_2315_v102_) + (d_2315_v102_))
        elif True:
            d_2316_v103_: _dafny.Array
            nw322_ = _dafny.Array(int(0), 23)
            d_2316_v103_ = nw322_
            d_2317_v104_: C0
            nw323_ = C0()
            nw323_.ctor__(d_2316_v103_)
            d_2317_v104_ = nw323_
            r2 = (p0) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aegsjgofj"))))
            (self).f3 = (self).fm13(globalState)
            d_2318_v105_: _dafny.Map
            d_2318_v105_ = _dafny.Map({self.f3: (self).f5})
            d_2319_v106_: D5
            d_2319_v106_ = D5_DC15(p0, self.f3, (self).f8)
            d_2318_v105_ = (d_2318_v105_).set(self.f3, (d_2319_v106_).cf26)
            r2 = default__.fm0(globalState)
        r1 = p0
        d_2320_v107_: _dafny.Array
        nw324_ = _dafny.Array(_dafny.CodePoint('D'), 28)
        d_2320_v107_ = nw324_
        d_2321_v108_: _dafny.MultiSet
        d_2321_v108_ = _dafny.MultiSet([d_2320_v107_])
        d_2322_v109_: _dafny.Map
        d_2322_v109_ = _dafny.Map({d_2321_v108_: not(default__.fm3(p0, globalState))})
        d_2323_v110_: _dafny.Seq
        d_2323_v110_ = _dafny.SeqWithoutIsStrInference([False])
        d_2324_v111_: C6
        nw325_ = C6()
        nw325_.ctor__(p1, default__.safeModuloInt(p1, 843), p1, (self).f5, ((d_2322_v109_)[d_2321_v108_] if (d_2321_v108_) in (d_2322_v109_) else (d_2323_v110_)[default__.safeIndex((self).f4, len(d_2323_v110_))]), (_dafny.Map({(self).f2: self.f3})) | ((self).f1), (self).f2)
        d_2324_v111_ = nw325_
        d_2325_v112_: _dafny.Map
        d_2325_v112_ = _dafny.Map({True: self.f3})
        r0 = ((d_2325_v112_) | (d_2325_v112_)) | (default__.fm9(globalState))
        r1 = (p1) * ((p1) - ((self).f4))
        r2 = (self).f4
        return r0, r1, r2

    def m5(self, globalState):
        r0: bool = False
        r1: bool = False
        if False:
            r0 = self.f3
            (self).f3 = self.f3
            d_2326_v0_: _dafny.Seq
            d_2326_v0_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3, self.f3, self.f3, self.f3])
            d_2327_v1_: _dafny.Seq
            d_2327_v1_ = _dafny.SeqWithoutIsStrInference([d_2326_v0_, d_2326_v0_])
            d_2328_v2_: D1
            d_2328_v2_ = D1_DC2((self).f2)
            d_2329_v3_: _dafny.Set
            d_2329_v3_ = _dafny.Set({(self).f2, (d_2328_v2_).cf2, (self).f2})
            d_2330_v4_: D13
            d_2330_v4_ = D13_DC34(not(self.f3), self.f3, (self).f4, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2331_i1_ in range(default__.abs(786))]))
            d_2332_v5_: C8
            nw326_ = C8()
            nw326_.ctor__((d_2327_v1_) + (d_2327_v1_), (d_2329_v3_).isdisjoint(d_2329_v3_), (self).f1, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2333_i0_ in range(default__.abs(696))])) + ((d_2330_v4_).cf56), (0) - ((0) - ((self).f4)), (self).f5)
            d_2332_v5_ = nw326_
            d_2334_v6_: D11
            d_2334_v6_ = D11_DC31((self).f4)
            d_2335_v7_: int
            d_2335_v7_ = 479
            d_2336_v8_: _dafny.Array
            nw327_ = _dafny.Array(int(0), 15)
            d_2336_v8_ = nw327_
            d_2337_v9_: _dafny.MultiSet
            d_2337_v9_ = _dafny.MultiSet([d_2336_v8_])
            d_2338_v10_: _dafny.MultiSet
            d_2338_v10_ = _dafny.MultiSet([self.f3])
            d_2339_v11_: str
            d_2339_v11_ = _dafny.CodePoint('k')
            d_2340_v12_: _dafny.Map
            d_2340_v12_ = _dafny.Map({self.f3: self.f3})
            d_2341_v13_: _dafny.Array
            nw328_ = _dafny.Array(None, 12)
            nw328_[int(0)] = -526
            nw328_[int(1)] = (d_2332_v5_).fm12((self).f4, d_2335_v7_, (self).f4, False, globalState)
            nw328_[int(2)] = d_2335_v7_
            nw328_[int(3)] = (self).f4
            nw328_[int(4)] = (0) - ((self).f4)
            nw328_[int(5)] = d_2335_v7_
            nw328_[int(6)] = d_2335_v7_
            nw328_[int(7)] = (0) - (((d_2337_v9_)[d_2336_v8_] if (d_2336_v8_) in (d_2337_v9_) else ((d_2338_v10_)[self.f3] if (self.f3) in (d_2338_v10_) else d_2335_v7_)))
            nw328_[int(8)] = d_2335_v7_
            nw328_[int(9)] = (self).f4
            nw328_[int(10)] = (len((D8_DC24(d_2339_v11_, d_2340_v12_)).cf42)) + (d_2335_v7_)
            nw328_[int(11)] = len(d_2340_v12_)
            d_2341_v13_ = nw328_
            index378_ = default__.safeIndex(71, (d_2336_v8_).length(0))
            (d_2336_v8_)[index378_] = len((_dafny.SeqWithoutIsStrInference([D1_DC4(self.f3, self.f3) for d_2342_i2_ in range(default__.abs(468))])) + (_dafny.SeqWithoutIsStrInference([D1_DC4(self.f3, self.f3)])))
            index379_ = default__.safeIndex(71, (d_2336_v8_).length(0))
            rhs399_ = (default__.fm56(d_2335_v7_, globalState) if (d_2338_v10_).isdisjoint(d_2338_v10_) else d_2334_v6_)
            rhs400_ = default__.safeModuloInt((self).f4, (self).f4)
            rhs401_ = d_2335_v7_
            rhs402_ = d_2335_v7_
            lhs266_ = d_2336_v8_
            lhs267_ = default__.safeIndex(71, (d_2336_v8_).length(0))
            d_2334_v6_ = rhs399_
            d_2335_v7_ = rhs400_
            d_2335_v7_ = rhs401_
            lhs266_[lhs267_] = rhs402_
            d_2343_v14_: C8
            nw329_ = C8()
            nw329_.ctor__((_dafny.SeqWithoutIsStrInference([default__.fm7(self.f3, self.f3, globalState), d_2326_v0_, d_2326_v0_, d_2326_v0_])) + ((_dafny.SeqWithoutIsStrInference([d_2326_v0_])).set(default__.safeIndex(d_2335_v7_, len(_dafny.SeqWithoutIsStrInference([d_2326_v0_]))), d_2326_v0_)), (d_2326_v0_)[default__.safeIndex((d_2336_v8_)[default__.safeIndex(71, (d_2336_v8_).length(0))], len(d_2326_v0_))], (self).f1, (self).f2, (d_2336_v8_)[default__.safeIndex(71, (d_2336_v8_).length(0))], (self).f5)
            d_2343_v14_ = nw329_
        elif True:
            d_2344_v15_: _dafny.Array
            nw330_ = _dafny.Array(int(0), 7)
            d_2344_v15_ = nw330_
            index380_ = default__.safeIndex(303, (d_2344_v15_).length(0))
            (d_2344_v15_)[index380_] = (self).f4
            index381_ = default__.safeIndex(303, (d_2344_v15_).length(0))
            (d_2344_v15_)[index381_] = (self).f4
            d_2345_v16_: _dafny.Map
            d_2345_v16_ = _dafny.Map({(self).f4: self.f3})
            d_2346_v17_: C7
            nw331_ = C7()
            nw331_.ctor__((self).f4, (self).f5, ((d_2345_v16_)[(d_2344_v15_)[default__.safeIndex(303, (d_2344_v15_).length(0))]] if ((d_2344_v15_)[default__.safeIndex(303, (d_2344_v15_).length(0))]) in (d_2345_v16_) else self.f3), ((self).f1) | ((self).f1), (self).f2)
            d_2346_v17_ = nw331_
            d_2347_v18_: _dafny.Set
            d_2347_v18_ = _dafny.Set({self.f3})
            d_2348_v19_: _dafny.Map
            d_2348_v19_ = _dafny.Map({len(d_2347_v18_): (self).f5})
            d_2349_v20_: _dafny.Map
            d_2349_v20_ = _dafny.Map({d_2348_v19_: self.f3})
            d_2349_v20_ = (d_2349_v20_).set((d_2348_v19_) | (d_2348_v19_), (self.f3) or (self.f3))
            d_2350_v21_: str
            d_2350_v21_ = _dafny.CodePoint('l')
            index382_ = default__.safeIndex(303, (d_2344_v15_).length(0))
            (d_2344_v15_)[index382_] = (0) - (len(default__.fm57(d_2350_v21_, globalState)))
            r0 = (self.f3 if self.f3 else self.f3)
        d_2351_v22_: _dafny.Array
        def lambda91_(d_2352_i4_):
            return (d_2352_i4_) + ((self).f4)

        init51_ = lambda91_
        nw332_ = _dafny.Array(None, 4)
        for i0_51_ in range(nw332_.length(0)):
            nw332_[i0_51_] = init51_(i0_51_)
        d_2351_v22_ = nw332_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2351_v22_).length(0)):
            d_2353_i3_: int = guard_loop_6_
            if (True) and (((0) <= (d_2353_i3_)) and ((d_2353_i3_) < ((d_2351_v22_).length(0)))):
                (d_2351_v22_)[(d_2353_i3_)] = default__.safeDivisionInt(d_2353_i3_, ((self).f4) * ((self).f4))
        d_2354_v23_: C6
        nw333_ = C6()
        nw333_.ctor__((self).f4, (self).f4, default__.fm0(globalState), (self).f5, True, ((self).f1) | ((self).f1), (self).f2)
        d_2354_v23_ = nw333_
        d_2355_v24_: _dafny.Array
        def lambda92_(d_2356_i5_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htylxj"))) + ((self).f2)

        init52_ = lambda92_
        nw334_ = _dafny.Array(None, 5)
        for i0_52_ in range(nw334_.length(0)):
            nw334_[i0_52_] = init52_(i0_52_)
        d_2355_v24_ = nw334_
        rhs403_ = (d_2354_v23_).fm13(globalState)
        rhs404_ = d_2354_v23_
        rhs405_ = self.f3
        rhs406_ = d_2355_v24_
        lhs268_ = self
        r0 = rhs403_
        d_2354_v23_ = rhs404_
        lhs268_.f3 = rhs405_
        d_2355_v24_ = rhs406_
        d_2357_v25_: C5
        nw335_ = C5()
        nw335_.ctor__(self.f3, ((self).f1) | ((self).f1), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2358_i6_ in range(default__.abs(448))]))
        d_2357_v25_ = nw335_
        d_2359_v26_: str
        d_2359_v26_ = _dafny.CodePoint('p')
        d_2359_v26_ = d_2359_v26_
        d_2360_v27_: _dafny.Seq
        d_2360_v27_ = _dafny.SeqWithoutIsStrInference([(d_2357_v25_).f15])
        d_2361_v28_: _dafny.Map
        d_2361_v28_ = _dafny.Map({not(self.f3): (d_2360_v27_)[default__.safeIndex(len((self).f2), len(d_2360_v27_))]})
        d_2362_v29_: _dafny.Seq
        d_2362_v29_ = _dafny.SeqWithoutIsStrInference([838, (self).fm12(len(d_2361_v28_), d_2354_v23_.f12, d_2354_v23_.f12, (d_2357_v25_).f15, globalState), d_2354_v23_.f12, d_2354_v23_.f12, 867])
        (d_2354_v23_).f12 = (0) - (((d_2362_v29_)[default__.safeIndex(852, len(d_2362_v29_))]) + (default__.safeModuloInt((self).f4, (self).f4)))
        d_2363_v30_: D4
        d_2363_v30_ = D4_DC13((d_2357_v25_).f15, d_2362_v29_)
        pat_let_tv53_ = d_2357_v25_
        def iife202_(_pat_let68_0):
            def iife203_(d_2364_dt__update__tmp_h0_):
                def iife204_(_pat_let69_0):
                    def iife205_(d_2365_dt__update_hcf21_h0_):
                        return D4_DC13(d_2365_dt__update_hcf21_h0_, (d_2364_dt__update__tmp_h0_).cf22)
                    return iife205_(_pat_let69_0)
                return iife204_((pat_let_tv53_).f15)
            return iife203_(_pat_let68_0)
        r0 = (iife202_(d_2363_v30_)).cf21
        r1 = True
        return r0, r1

    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C10(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f7, f4, f5, f3, f1, f2):
        (self)._f7 = f7
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return self.f3

    def fm14(self, p0, p1, p2, p3, globalState):
        return ((self).f2) + ((self).f2)

    def fm12(self, p0, p1, p2, p3, globalState):
        return (self).f4

    def fm11(self, p0, globalState):
        return (-530) > ((self).f4)

    def fm15(self, p0, p1, globalState):
        return ((self).f7) - ((self).f7)

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        d_2366_v0_: _dafny.Seq
        d_2366_v0_ = _dafny.SeqWithoutIsStrInference([(self).fm11(self.f3, globalState)])
        d_2367_v1_: D3
        d_2367_v1_ = D3_DC10(default__.fm0(globalState), p0)
        d_2368_v2_: _dafny.MultiSet
        d_2368_v2_ = _dafny.MultiSet([self.f3])
        rhs407_ = (((_dafny.SeqWithoutIsStrInference([self.f3, self.f3, False])) + (_dafny.SeqWithoutIsStrInference([self.f3, True, self.f3])) if self.f3 else _dafny.SeqWithoutIsStrInference([self.f3]))).set(default__.safeIndex(420, len(((_dafny.SeqWithoutIsStrInference([self.f3, self.f3, False])) + (_dafny.SeqWithoutIsStrInference([self.f3, True, self.f3])) if self.f3 else _dafny.SeqWithoutIsStrInference([self.f3])))), (d_2368_v2_).issubset(d_2368_v2_))
        rhs408_ = d_2367_v1_
        rhs409_ = default__.safeDivisionInt(len((self).f2), (self).f7)
        d_2366_v0_ = rhs407_
        d_2367_v1_ = rhs408_
        r1 = rhs409_
        d_2369_v3_: _dafny.Set
        d_2369_v3_ = _dafny.Set({self.f3})
        d_2370_v4_: _dafny.Seq
        d_2370_v4_ = _dafny.SeqWithoutIsStrInference([(self).f4])
        d_2371_v5_: _dafny.MultiSet
        d_2371_v5_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([552, len(d_2369_v3_)]), d_2370_v4_])
        if (self.f3) == ((d_2371_v5_) == (_dafny.MultiSet([d_2370_v4_, d_2370_v4_, d_2370_v4_]))):
            d_2372_v6_: _dafny.Map
            d_2372_v6_ = _dafny.Map({(self).fm11(self.f3, globalState): d_2366_v0_})
            d_2372_v6_ = d_2372_v6_
            d_2373_v7_: D3
            d_2373_v7_ = D3_DC9(p1, self.f3)
            d_2374_v8_: _dafny.Map
            d_2374_v8_ = _dafny.Map({(self).f4: d_2373_v7_})
            d_2375_v9_: _dafny.Map
            d_2375_v9_ = _dafny.Map({p0: (self).f4})
            r2 = (self).fm12(default__.safeDivisionInt((self).f7, p1), (self).fm15(_dafny.Map({d_2374_v8_: p0}), d_2375_v9_, globalState), p1, (self.f3 if self.f3 else self.f3), globalState)
            d_2376_v10_: _dafny.MultiSet
            d_2376_v10_ = _dafny.MultiSet([(self).f5])
            d_2377_v11_: _dafny.Array
            nw336_ = _dafny.Array(None, 4)
            nw336_[int(0)] = (0) - (len((self).f2))
            nw336_[int(1)] = (self).f4
            nw336_[int(2)] = (d_2376_v10_).cardinality
            nw336_[int(3)] = p0
            d_2377_v11_ = nw336_
            d_2377_v11_ = d_2377_v11_
            d_2378_v12_: _dafny.Array
            nw337_ = _dafny.Array(_dafny.Seq({}), 11)
            d_2378_v12_ = nw337_
            d_2379_v13_: _dafny.Map
            d_2379_v13_ = _dafny.Map({(self).f7: False})
            d_2380_v14_: D0
            d_2380_v14_ = D0_DC0(len((d_2379_v13_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))), self.f3)))
            d_2381_v15_: _dafny.Seq
            d_2381_v15_ = _dafny.SeqWithoutIsStrInference([d_2380_v14_])
            index383_ = default__.safeIndex(145, (d_2378_v12_).length(0))
            (d_2378_v12_)[index383_] = d_2381_v15_
            index384_ = default__.safeIndex(145, (d_2378_v12_).length(0))
            (d_2378_v12_)[index384_] = _dafny.SeqWithoutIsStrInference([D0_DC0(len(((self).f2).set(default__.safeIndex((d_2368_v2_).cardinality, len((self).f2)), (D2_DC6(_dafny.CodePoint('r'), self.f3, True, p1)).cf9))) for d_2382_i0_ in range(default__.abs(702))])
            d_2383_v16_: _dafny.Seq
            d_2383_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hccwe"))
            d_2384_v17_: str
            d_2384_v17_ = _dafny.CodePoint('j')
            d_2385_v18_: _dafny.Seq
            d_2385_v18_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqaqunld"))).set(default__.safeIndex(846, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqaqunld")))), d_2384_v17_), d_2383_v16_, d_2383_v16_])
            d_2386_v19_: _dafny.Map
            d_2386_v19_ = _dafny.Map({self.f3: (self).f7})
            d_2387_v20_: _dafny.Map
            d_2387_v20_ = _dafny.Map({D3_DC10(p0, ((d_2386_v19_)[self.f3] if (self.f3) in (d_2386_v19_) else p0)): p0})
            d_2388_v21_: _dafny.Map
            d_2388_v21_ = _dafny.Map({d_2384_v17_: -156})
            d_2389_v22_: _dafny.Map
            d_2389_v22_ = _dafny.Map({d_2388_v21_: (self).f4})
            rhs410_ = ((d_2385_v18_)[default__.safeIndex(((d_2387_v20_)[d_2367_v1_] if (d_2367_v1_) in (d_2387_v20_) else 919), len(d_2385_v18_))]).set(default__.safeIndex((self).f7, len((d_2385_v18_)[default__.safeIndex(((d_2387_v20_)[d_2367_v1_] if (d_2367_v1_) in (d_2387_v20_) else 919), len(d_2385_v18_))])), d_2384_v17_)
            rhs411_ = (((d_2389_v22_)[_dafny.Map({d_2384_v17_: (self).f7})] if (_dafny.Map({d_2384_v17_: (self).f7})) in (d_2389_v22_) else (self).f7)) * ((self).f7)
            d_2383_v16_ = rhs410_
            r1 = rhs411_
        elif True:
            d_2390_v23_: _dafny.Array
            nw338_ = _dafny.Array(int(0), 20)
            d_2390_v23_ = nw338_
            d_2391_v24_: _dafny.Map
            d_2391_v24_ = _dafny.Map({-264: d_2390_v23_})
            d_2392_v25_: _dafny.MultiSet
            d_2392_v25_ = _dafny.MultiSet([len(d_2391_v24_)])
            d_2393_v26_: str
            d_2393_v26_ = _dafny.CodePoint('f')
            d_2394_v27_: _dafny.Set
            d_2394_v27_ = _dafny.Set({d_2393_v26_, d_2393_v26_, d_2393_v26_})
            r2 = ((self).f7) * ((self).fm12((d_2392_v25_).cardinality, len(d_2394_v27_), p0, self.f3, globalState))
            (self).f3 = self.f3
            d_2395_v28_: C7
            nw339_ = C7()
            nw339_.ctor__(p1, (self).f5, self.f3, (self).f1, (self).f2)
            d_2395_v28_ = nw339_
            d_2396_v29_: _dafny.Map
            d_2396_v29_ = _dafny.Map({(self).f5: self.f3})
            d_2397_v30_: bool
            d_2398_v31_: bool
            d_2399_v32_: _dafny.Seq
            d_2400_v33_: _dafny.Seq
            out62_: bool
            out63_: bool
            out64_: _dafny.Seq
            out65_: _dafny.Seq
            out62_, out63_, out64_, out65_ = default__.m0(d_2396_v29_, globalState)
            d_2397_v30_ = out62_
            d_2398_v31_ = out63_
            d_2399_v32_ = out64_
            d_2400_v33_ = out65_
            d_2401_v34_: _dafny.Map
            d_2401_v34_ = _dafny.Map({d_2390_v23_: (self).f5})
            d_2402_v35_: D0
            d_2402_v35_ = D0_DC1(696)
            d_2403_v37_: _dafny.Seq
            d_2403_v37_ = _dafny.SeqWithoutIsStrInference([(self).f2])
            d_2404_v38_: C9
            nw340_ = C9()
            def iife206_():
                coll66_ = _dafny.Map()
                compr_66_: _dafny.Seq
                for compr_66_ in (d_2403_v37_).Elements:
                    d_2405_v36_: _dafny.Seq = compr_66_
                    if (d_2405_v36_) in (d_2403_v37_):
                        coll66_[d_2405_v36_] = d_2398_v31_
                return _dafny.Map(coll66_)
            nw340_.ctor__((self).f5, d_2401_v34_, (d_2402_v35_).cf1, (self).f5, d_2398_v31_, (iife206_()
            ).set((self).f2, self.f3), (self).f2)
            d_2404_v38_ = nw340_
        d_2406_v39_: _dafny.Seq
        d_2406_v39_ = _dafny.SeqWithoutIsStrInference([d_2366_v0_])
        d_2407_v40_: _dafny.Seq
        d_2407_v40_ = d_2366_v0_
        d_2408_v41_: _dafny.Array
        nw341_ = _dafny.Array(None, 9)
        nw341_[int(0)] = _dafny.SeqWithoutIsStrInference([self.f3])
        nw341_[int(1)] = d_2366_v0_
        nw341_[int(2)] = _dafny.SeqWithoutIsStrInference([self.f3, self.f3])
        nw341_[int(3)] = d_2366_v0_
        nw341_[int(4)] = _dafny.SeqWithoutIsStrInference([self.f3])
        nw341_[int(5)] = d_2366_v0_
        nw341_[int(6)] = (d_2406_v39_)[default__.safeIndex((self).f7, len(d_2406_v39_))]
        nw341_[int(7)] = (d_2366_v0_) + (d_2366_v0_)
        nw341_[int(8)] = (d_2407_v40_)
        d_2408_v41_ = nw341_
        index385_ = default__.safeIndex(203, (d_2408_v41_).length(0))
        (d_2408_v41_)[index385_] = d_2366_v0_
        index386_ = default__.safeIndex(203, (d_2408_v41_).length(0))
        (d_2408_v41_)[index386_] = (_dafny.SeqWithoutIsStrInference([self.f3])) + (d_2366_v0_)
        d_2409_v42_: _dafny.Map
        d_2409_v42_ = _dafny.Map({(self).f4: (self).f1})
        d_2410_v43_: _dafny.Set
        d_2410_v43_ = _dafny.Set({(self).f7})
        d_2411_v45_: _dafny.Seq
        d_2411_v45_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpqhqgii"))])
        d_2412_v46_: C6
        nw342_ = C6()
        def iife207_():
            coll67_ = _dafny.Map()
            compr_67_: _dafny.Seq
            for compr_67_ in (d_2411_v45_).Elements:
                d_2413_v44_: _dafny.Seq = compr_67_
                if (d_2413_v44_) in (d_2411_v45_):
                    coll67_[d_2413_v44_] = self.f3
            return _dafny.Map(coll67_)
        nw342_.ctor__(p1, default__.safeDivisionInt(-575, (self).f7), p0, (self).f5, self.f3, ((d_2409_v42_)[len(d_2410_v43_)] if (len(d_2410_v43_)) in (d_2409_v42_) else iife207_()
        ), (self).f2)
        d_2412_v46_ = nw342_
        if False:
            d_2414_v47_: D6
            d_2414_v47_ = D6_DC16(_dafny.Map({d_2412_v46_.f12: self.f3}))
            d_2415_v48_: _dafny.Map
            d_2415_v48_ = _dafny.Map({157: self.f3})
            d_2416_v50_: _dafny.Map
            d_2416_v50_ = _dafny.Map({self.f3: (d_2368_v2_).cardinality})
            pat_let_tv54_ = d_2415_v48_
            pat_let_tv55_ = d_2415_v48_
            d_2417_v51_: _dafny.Array
            nw343_ = _dafny.Array(None, 19)
            nw343_[int(0)] = d_2414_v47_
            nw343_[int(1)] = d_2414_v47_
            nw343_[int(2)] = d_2414_v47_
            nw343_[int(3)] = d_2414_v47_
            def iife208_(_pat_let70_0):
                def iife209_(d_2418_dt__update__tmp_h0_):
                    def iife210_(_pat_let71_0):
                        def iife211_(d_2419_dt__update_hcf27_h0_):
                            return D6_DC16(d_2419_dt__update_hcf27_h0_)
                        return iife211_(_pat_let71_0)
                    return iife210_(pat_let_tv54_)
                return iife209_(_pat_let70_0)
            nw343_[int(4)] = iife208_(D6_DC16((d_2415_v48_).set(440, self.f3)))
            nw343_[int(5)] = d_2414_v47_
            def iife212_():
                coll68_ = _dafny.Map()
                compr_68_: int
                for compr_68_ in _dafny.IntegerRange(828, 777):
                    d_2420_v49_: int = compr_68_
                    if ((828) <= (d_2420_v49_)) and ((d_2420_v49_) < (777)):
                        coll68_[default__.safeDivisionInt(d_2420_v49_, d_2412_v46_.f12)] = self.f3
                return _dafny.Map(coll68_)
            nw343_[int(6)] = D6_DC16(iife212_()
)
            def iife213_(_pat_let72_0):
                def iife214_(d_2421_dt__update__tmp_h1_):
                    def iife215_(_pat_let73_0):
                        def iife216_(d_2422_dt__update_hcf27_h1_):
                            return D6_DC16(d_2422_dt__update_hcf27_h1_)
                        return iife216_(_pat_let73_0)
                    return iife215_(pat_let_tv55_)
                return iife214_(_pat_let72_0)
            nw343_[int(7)] = iife213_(D6_DC16(d_2415_v48_))
            nw343_[int(8)] = d_2414_v47_
            nw343_[int(9)] = d_2414_v47_
            nw343_[int(10)] = D6_DC16(d_2415_v48_)
            nw343_[int(11)] = default__.fm54(d_2416_v50_, True, globalState)
            nw343_[int(12)] = D6_DC16(_dafny.Map({(d_2412_v46_).f11: self.f3}))
            nw343_[int(13)] = d_2414_v47_
            nw343_[int(14)] = d_2414_v47_
            nw343_[int(15)] = d_2414_v47_
            nw343_[int(16)] = d_2414_v47_
            nw343_[int(17)] = D6_DC16(d_2415_v48_)
            nw343_[int(18)] = d_2414_v47_
            d_2417_v51_ = nw343_
            d_2423_v52_: _dafny.Seq
            d_2423_v52_ = _dafny.SeqWithoutIsStrInference([d_2417_v51_])
            d_2424_v53_: _dafny.Map
            d_2424_v53_ = _dafny.Map({self.f3: d_2370_v4_})
            d_2425_v55_: _dafny.Array
            nw344_ = _dafny.Array(None, 19)
            nw344_[int(0)] = (self).f4
            nw344_[int(1)] = p1
            nw344_[int(2)] = (p1) * (73)
            nw344_[int(3)] = len((d_2423_v52_) + (d_2423_v52_))
            nw344_[int(4)] = len(((d_2424_v53_)[self.f3] if (self.f3) in (d_2424_v53_) else d_2370_v4_))
            nw344_[int(5)] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_2366_v0_ for d_2426_i1_ in range(default__.abs(209))])))
            nw344_[int(6)] = (D9_DC26(len((_dafny.Map({self.f3: (d_2412_v46_).f11})).set(self.f3, p0)))).cf44
            nw344_[int(7)] = (d_2412_v46_).fm21(self.f3, (d_2370_v4_)[default__.safeIndex((self).f7, len(d_2370_v4_))], globalState)
            nw344_[int(8)] = (self).f7
            nw344_[int(9)] = (d_2412_v46_).f11
            nw344_[int(10)] = default__.safeDivisionInt((d_2412_v46_).f11, (0) - ((d_2412_v46_).f11))
            nw344_[int(11)] = (d_2412_v46_).f11
            nw344_[int(12)] = default__.safeModuloInt((d_2412_v46_).f11, (d_2412_v46_).f11)
            nw344_[int(13)] = (0) - ((0) - (len((d_2416_v50_).set(self.f3, 807))))
            nw344_[int(14)] = (0) - (d_2412_v46_.f12)
            nw344_[int(15)] = d_2412_v46_.f12
            nw344_[int(16)] = -920
            nw344_[int(17)] = default__.safeModuloInt((self).f7, len(d_2366_v0_))
            def iife217_():
                coll69_ = _dafny.Map()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(-611, 128):
                    d_2427_v54_: int = compr_69_
                    if ((-611) <= (d_2427_v54_)) and ((d_2427_v54_) < (128)):
                        coll69_[(d_2427_v54_) + ((self).f7)] = d_2368_v2_
                return _dafny.Map(coll69_)
            nw344_[int(18)] = len(iife217_()
            )
            d_2425_v55_ = nw344_
            index387_ = default__.safeIndex(614, (d_2425_v55_).length(0))
            (d_2425_v55_)[index387_] = p0
            index388_ = default__.safeIndex(614, (d_2425_v55_).length(0))
            (d_2425_v55_)[index388_] = (self).fm12(d_2412_v46_.f12, (d_2412_v46_).f11, default__.fm0(globalState), not (self.f3) or (self.f3), globalState)
            if not(False):
                r2 = default__.safeDivisionInt((d_2412_v46_).f11, d_2412_v46_.f12)
                (self).f3 = (False) or (not(self.f3))
                d_2428_v56_: C8
                nw345_ = C8()
                nw345_.ctor__(d_2406_v39_, self.f3, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")), len(d_2416_v50_), (self).f5)
                d_2428_v56_ = nw345_
                d_2429_v57_: _dafny.Seq
                d_2429_v57_ = _dafny.SeqWithoutIsStrInference([d_2428_v56_, d_2428_v56_, d_2428_v56_])
                d_2429_v57_ = d_2429_v57_
                d_2430_v58_: _dafny.Map
                d_2430_v58_ = _dafny.Map({(0) - ((0) - (d_2412_v46_.f12)): (d_2412_v46_).f11})
                d_2430_v58_ = d_2430_v58_
                index389_ = default__.safeIndex(203, (d_2408_v41_).length(0))
                rhs412_ = (d_2425_v55_)[default__.safeIndex(614, (d_2425_v55_).length(0))]
                rhs413_ = ((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))]) + ((default__.fm7(self.f3, not(self.f3), globalState)).set(default__.safeIndex(p1, len(default__.fm7(self.f3, not(self.f3), globalState))), not(self.f3)))
                rhs414_ = d_2366_v0_
                lhs269_ = d_2408_v41_
                lhs270_ = default__.safeIndex(203, (d_2408_v41_).length(0))
                r2 = rhs412_
                lhs269_[lhs270_] = rhs413_
                d_2366_v0_ = rhs414_
            elif True:
                d_2431_v59_: str
                d_2431_v59_ = _dafny.CodePoint('q')
                d_2432_v60_: C3
                nw346_ = C3()
                nw346_.ctor__(len((self).f2), (self).f5, self.f3, ((self).f1).set((self).f2, True), (d_2412_v46_).fm14(self.f3, self.f3, d_2431_v59_, (self).f7, globalState))
                d_2432_v60_ = nw346_
                (d_2412_v46_).f12 = (d_2370_v4_)[default__.safeIndex(p0, len(d_2370_v4_))]
                d_2433_v61_: _dafny.Seq
                d_2433_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mah"))
                rhs415_ = (len((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))])) - ((d_2412_v46_).f11)
                rhs416_ = (self).f2
                rhs417_ = p1
                r1 = rhs415_
                d_2433_v61_ = rhs416_
                r1 = rhs417_
                d_2434_v62_: _dafny.Map
                d_2434_v62_ = _dafny.Map({self.f3: self.f3})
                d_2435_v63_: D1
                d_2435_v63_ = D1_DC4(((d_2434_v62_)[self.f3] if (self.f3) in (d_2434_v62_) else self.f3), not(self.f3))
                d_2436_v64_: _dafny.Map
                d_2436_v64_ = _dafny.Map({self.f3: d_2435_v63_})
                d_2436_v64_ = ((d_2436_v64_) | (d_2436_v64_)) | (d_2436_v64_)
                (self).f3 = ((0) - ((len(_dafny.SeqWithoutIsStrInference([d_2431_v59_ for d_2437_i2_ in range(default__.abs(391))])) if self.f3 else (d_2425_v55_)[default__.safeIndex(614, (d_2425_v55_).length(0))]))) > ((self).f7)
            d_2438_v65_: _dafny.Map
            d_2438_v65_ = _dafny.Map({self.f3: default__.fm7(self.f3, self.f3, globalState)})
            (d_2412_v46_).f12 = len((d_2438_v65_).set(True, (d_2406_v39_)[default__.safeIndex(-614, len(d_2406_v39_))]))
            (self).f3 = self.f3
            d_2408_v41_ = d_2408_v41_
        elif True:
            d_2366_v0_ = (d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))]
            (self).f3 = self.f3
            if self.f3:
                d_2439_v66_: _dafny.Seq
                d_2439_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwuuwcna"))
                d_2439_v66_ = d_2439_v66_
                d_2440_v67_: str
                d_2440_v67_ = _dafny.CodePoint('a')
                (d_2412_v46_).f12 = len((d_2412_v46_).fm14((d_2412_v46_.f12) < (len(default__.fm20(self.f3, not(self.f3), self.f3, (d_2366_v0_).set(default__.safeIndex((d_2412_v46_).f11, len(d_2366_v0_)), True), globalState))), self.f3, d_2440_v67_, (self).f4, globalState))
                d_2439_v66_ = d_2439_v66_
                d_2440_v67_ = d_2440_v67_
                d_2441_v68_: _dafny.Set
                d_2441_v68_ = d_2410_v43_
                d_2442_v69_: _dafny.Map
                d_2442_v69_ = _dafny.Map({self.f3: d_2441_v68_})
                d_2443_v70_: _dafny.MultiSet
                d_2443_v70_ = _dafny.MultiSet([len(d_2410_v43_)])
                d_2444_v71_: C2
                nw347_ = C2()
                nw347_.ctor__((0) - ((d_2412_v46_).f11), self.f3, _dafny.Map({(self).f2: self.f3}), d_2439_v66_)
                d_2444_v71_ = nw347_
                d_2445_v72_: D13
                d_2445_v72_ = D13_DC33(d_2444_v71_)
                d_2446_v73_: _dafny.Seq
                d_2446_v73_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
                d_2447_v74_: _dafny.Array
                def lambda93_(d_2448_v46_):
                    def lambda94_(d_2449_i3_):
                        return default__.safeDivisionInt(d_2449_i3_, (d_2448_v46_).f11)

                    return lambda94_

                init53_ = lambda93_(d_2412_v46_)
                nw348_ = _dafny.Array(None, 29)
                for i0_53_ in range(nw348_.length(0)):
                    nw348_[i0_53_] = init53_(i0_53_)
                d_2447_v74_ = nw348_
                d_2450_v75_: _dafny.Map
                d_2450_v75_ = _dafny.Map({d_2447_v74_: (self).f5})
                d_2451_v76_: C9
                nw349_ = C9()
                nw349_.ctor__((d_2446_v73_)[default__.safeIndex(d_2412_v46_.f12, len(d_2446_v73_))], d_2450_v75_, (d_2370_v4_)[default__.safeIndex(373, len(d_2370_v4_))], (self).f5, self.f3, (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlgym")))
                d_2451_v76_ = nw349_
                d_2452_v77_: _dafny.Map
                d_2452_v77_ = _dafny.Map({d_2445_v72_: d_2451_v76_})
                d_2453_v78_: _dafny.Array
                nw350_ = _dafny.Array(None, 17)
                nw350_[int(0)] = (self).f4
                nw350_[int(1)] = len((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))])
                nw350_[int(2)] = (0) - (d_2412_v46_.f12)
                nw350_[int(3)] = p0
                nw350_[int(4)] = len(d_2439_v66_)
                nw350_[int(5)] = ((d_2443_v70_)[(d_2412_v46_).f11] if ((d_2412_v46_).f11) in (d_2443_v70_) else 832)
                nw350_[int(6)] = 922
                nw350_[int(7)] = (d_2412_v46_).f11
                nw350_[int(8)] = 844
                nw350_[int(9)] = -112
                nw350_[int(10)] = p0
                nw350_[int(11)] = d_2412_v46_.f12
                nw350_[int(12)] = len(d_2439_v66_)
                nw350_[int(13)] = len(d_2452_v77_)
                nw350_[int(14)] = d_2412_v46_.f12
                nw350_[int(15)] = (d_2412_v46_).f11
                nw350_[int(16)] = len((self).f2)
                d_2453_v78_ = nw350_
                d_2454_v79_: _dafny.Map
                d_2454_v79_ = _dafny.Map({d_2370_v4_: d_2447_v74_})
                d_2455_v80_: _dafny.Array
                nw351_ = _dafny.Array(None, 27)
                nw351_[int(0)] = self.f3
                nw351_[int(1)] = (d_2410_v43_).isdisjoint(d_2410_v43_)
                nw351_[int(2)] = self.f3
                nw351_[int(3)] = False
                nw351_[int(4)] = not((d_2368_v2_).issubset(d_2368_v2_))
                nw351_[int(5)] = self.f3
                nw351_[int(6)] = False
                nw351_[int(7)] = self.f3
                nw351_[int(8)] = self.f3
                nw351_[int(9)] = (len(default__.fm6(self.f3, d_2440_v67_, globalState))) <= (p1)
                nw351_[int(10)] = self.f3
                nw351_[int(11)] = self.f3
                nw351_[int(12)] = self.f3
                nw351_[int(13)] = (d_2370_v4_) < (d_2370_v4_)
                nw351_[int(14)] = self.f3
                nw351_[int(15)] = False
                nw351_[int(16)] = (p1) <= (len(d_2442_v69_))
                nw351_[int(17)] = self.f3
                nw351_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_2412_v46_.f12])) in (d_2454_v79_)
                nw351_[int(19)] = (d_2443_v70_).issubset(d_2443_v70_)
                nw351_[int(20)] = self.f3
                nw351_[int(21)] = self.f3
                nw351_[int(22)] = (d_2444_v71_).fm11(self.f3, globalState)
                nw351_[int(23)] = ((self).f2) != ((self).f2)
                nw351_[int(24)] = self.f3
                nw351_[int(25)] = (d_2412_v46_.f12) <= (d_2412_v46_.f12)
                nw351_[int(26)] = self.f3
                d_2455_v80_ = nw351_
                d_2455_v80_ = (d_2451_v76_).f8
            elif True:
                d_2456_v81_: C1
                nw352_ = C1()
                nw352_.ctor__()
                d_2456_v81_ = nw352_
                d_2456_v81_ = d_2456_v81_
                index390_ = default__.safeIndex(726, ((self).f5).length(0))
                ((self).f5)[index390_] = (d_2366_v0_)[default__.safeIndex((d_2412_v46_).f11, len(d_2366_v0_))]
                index391_ = default__.safeIndex(726, ((self).f5).length(0))
                ((self).f5)[index391_] = (d_2412_v46_).fm11(self.f3, globalState)
                r2 = d_2412_v46_.f12
                d_2457_v82_: _dafny.Map
                d_2457_v82_ = _dafny.Map({(self).f4: (self).f7})
                d_2458_v83_: _dafny.Map
                d_2458_v83_ = _dafny.Map({((self).f5)[default__.safeIndex(726, ((self).f5).length(0))]: ((self).f5)[default__.safeIndex(726, ((self).f5).length(0))]})
                d_2459_v84_: _dafny.Map
                d_2459_v84_ = _dafny.Map({False: len((self).f2)})
                d_2460_v85_: _dafny.Map
                d_2460_v85_ = _dafny.Map({(d_2412_v46_).f11: self.f3})
                d_2461_v86_: _dafny.MultiSet
                d_2461_v86_ = _dafny.MultiSet([(d_2412_v46_).f11, len(d_2460_v85_), (d_2412_v46_).f11, (d_2412_v46_).f11, d_2412_v46_.f12])
                d_2462_v87_: _dafny.Map
                d_2462_v87_ = _dafny.Map({((self).f5)[default__.safeIndex(726, ((self).f5).length(0))]: d_2461_v86_})
                d_2463_v88_: _dafny.Array
                nw353_ = _dafny.Array(None, 17)
                nw353_[int(0)] = len((default__.fm33((_dafny.MultiSet([((self).f5)[default__.safeIndex(726, ((self).f5).length(0))]])).cardinality, globalState)))
                nw353_[int(1)] = (d_2412_v46_).f11
                nw353_[int(2)] = d_2412_v46_.f12
                nw353_[int(3)] = (self).f7
                nw353_[int(4)] = (0) - (((d_2457_v82_)[(self).f4] if ((self).f4) in (d_2457_v82_) else len(d_2369_v3_)))
                nw353_[int(5)] = p1
                nw353_[int(6)] = p0
                nw353_[int(7)] = len((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))])
                nw353_[int(8)] = 484
                nw353_[int(9)] = 34
                nw353_[int(10)] = ((d_2368_v2_)[False] if (False) in (d_2368_v2_) else len(d_2458_v83_))
                nw353_[int(11)] = d_2412_v46_.f12
                nw353_[int(12)] = len(d_2459_v84_)
                nw353_[int(13)] = d_2412_v46_.f12
                nw353_[int(14)] = 291
                nw353_[int(15)] = d_2412_v46_.f12
                nw353_[int(16)] = len(d_2462_v87_)
                d_2463_v88_ = nw353_
                d_2464_v89_: str
                d_2464_v89_ = _dafny.CodePoint('g')
                d_2406_v39_ = default__.fm58(default__.fm31(default__.fm0(globalState), ((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))])[default__.safeIndex((self).f7, len((d_2408_v41_)[default__.safeIndex(203, (d_2408_v41_).length(0))]))], len(_dafny.Map({d_2463_v88_: len(d_2457_v82_)})), d_2464_v89_, globalState), d_2464_v89_, d_2459_v84_, globalState)
                index392_ = default__.safeIndex(726, ((self).f5).length(0))
                ((self).f5)[index392_] = ((d_2412_v46_).f11) < ((0) - (p0))
            d_2465_v90_: str
            d_2465_v90_ = _dafny.CodePoint('w')
            d_2466_v91_: D0
            d_2466_v91_ = D0_DC0(len(((self).f2).set(default__.safeIndex((self).fm12((d_2412_v46_).f11, d_2412_v46_.f12, p0, self.f3, globalState), len((self).f2)), d_2465_v90_)))
            d_2466_v91_ = d_2466_v91_
            d_2467_v92_: _dafny.Map
            d_2467_v92_ = _dafny.Map({(self).f4: self.f3})
            (self).f3 = (self.f3) or ((d_2467_v92_) == (d_2467_v92_))
        if True:
            r1 = d_2412_v46_.f12
            d_2468_v93_: C1
            nw354_ = C1()
            nw354_.ctor__()
            d_2468_v93_ = nw354_
            d_2469_v94_: _dafny.Map
            d_2469_v94_ = _dafny.Map({d_2468_v93_: False})
            d_2470_v95_: _dafny.Seq
            d_2470_v95_ = _dafny.SeqWithoutIsStrInference([d_2469_v94_])
            d_2471_v96_: T1
            nw355_ = C2()
            nw355_.ctor__(default__.safeModuloInt(-309, p0), True, (_dafny.Map({(self).f2: self.f3})) | ((self).f1), (self).f2)
            d_2471_v96_ = nw355_
            rhs418_ = d_2471_v96_.f3
            rhs419_ = (d_2412_v46_.f12) <= ((self).f4)
            rhs420_ = d_2470_v95_
            rhs421_ = d_2471_v96_
            lhs271_ = self
            lhs272_ = self
            lhs271_.f3 = rhs418_
            lhs272_.f3 = rhs419_
            d_2470_v95_ = rhs420_
            d_2471_v96_ = rhs421_
            d_2472_v97_: _dafny.Array
            def lambda95_(d_2473_p1_):
                def lambda96_(d_2474_i4_):
                    return (d_2474_i4_) * (d_2473_p1_)

                return lambda96_

            init54_ = lambda95_(p1)
            nw356_ = _dafny.Array(None, 8)
            for i0_54_ in range(nw356_.length(0)):
                nw356_[i0_54_] = init54_(i0_54_)
            d_2472_v97_ = nw356_
            d_2475_v98_: _dafny.Seq
            d_2475_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbi"))
            index393_ = default__.safeIndex(161, (d_2472_v97_).length(0))
            (d_2472_v97_)[index393_] = -246
            index394_ = default__.safeIndex(161, (d_2472_v97_).length(0))
            rhs422_ = d_2472_v97_
            rhs423_ = (d_2475_v98_) + ((d_2475_v98_ if d_2471_v96_.f3 else (self).f2))
            rhs424_ = d_2412_v46_.f12
            rhs425_ = self.f3
            lhs273_ = d_2472_v97_
            lhs274_ = default__.safeIndex(161, (d_2472_v97_).length(0))
            lhs275_ = self
            d_2472_v97_ = rhs422_
            d_2475_v98_ = rhs423_
            lhs273_[lhs274_] = rhs424_
            lhs275_.f3 = rhs425_
            r1 = (self).fm12(default__.safeModuloInt(375, (self).f4), default__.safeDivisionInt((self).f7, d_2412_v46_.f12), (448) - ((d_2472_v97_)[default__.safeIndex(161, (d_2472_v97_).length(0))]), d_2471_v96_.f3, globalState)
            if not(not(d_2471_v96_.f3)):
                (d_2412_v46_).f12 = p0
                d_2476_v99_: str
                d_2476_v99_ = _dafny.CodePoint('i')
                d_2477_v100_: C3
                nw357_ = C3()
                nw357_.ctor__(p1, (self).f5, d_2471_v96_.f3, (((d_2471_v96_).f1).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2478_i5_ in range(default__.abs(912))]), d_2471_v96_.f3)) | ((d_2471_v96_).f1), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2479_i6_ in range(default__.abs(686))])).set(default__.safeIndex((d_2412_v46_).f11, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2480_i6_ in range(default__.abs(686))]))), d_2476_v99_))
                d_2477_v100_ = nw357_
                index395_ = default__.safeIndex(161, (d_2472_v97_).length(0))
                (d_2472_v97_)[index395_] = len((d_2475_v98_ if self.f3 else (d_2471_v96_).f2))
                d_2481_v101_: _dafny.Map
                d_2481_v101_ = _dafny.Map({-199: p0})
                d_2482_v102_: _dafny.Map
                d_2482_v102_ = _dafny.Map({(d_2412_v46_).f11: d_2481_v101_})
                d_2483_v103_: _dafny.Map
                d_2483_v103_ = _dafny.Map({(self).f7: (self).f5})
                d_2482_v102_ = (d_2482_v102_).set(len(d_2483_v103_), d_2481_v101_)
                index396_ = default__.safeIndex(812, ((self).f5).length(0))
                ((self).f5)[index396_] = default__.fm3((d_2412_v46_).f11, globalState)
                index397_ = default__.safeIndex(812, ((self).f5).length(0))
                ((self).f5)[index397_] = (self.f3 if d_2471_v96_.f3 else True)
            elif True:
                (d_2471_v96_).f3 = not(self.f3)
                d_2484_v104_: C2
                nw358_ = C2()
                nw358_.ctor__((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2485_i7_ in range(default__.abs(303))]))), d_2471_v96_.f3, (self).f1, d_2475_v98_)
                d_2484_v104_ = nw358_
                d_2486_v105_: D13
                d_2486_v105_ = D13_DC36(D13_DC33(d_2484_v104_))
                d_2486_v105_ = d_2486_v105_
                d_2487_v106_: str
                d_2487_v106_ = _dafny.CodePoint('i')
                d_2487_v106_ = d_2487_v106_
                index398_ = default__.safeIndex(161, (d_2472_v97_).length(0))
                (d_2472_v97_)[index398_] = (p0) * (p1)
                d_2488_v107_: _dafny.Array
                nw359_ = _dafny.Array(False, 24)
                d_2488_v107_ = nw359_
                d_2488_v107_ = (self).f5
        elif True:
            (self).f3 = self.f3
            d_2489_v108_: str
            d_2489_v108_ = _dafny.CodePoint('p')
            d_2490_v109_: _dafny.Map
            d_2490_v109_ = _dafny.Map({self.f3: len(default__.fm31(d_2412_v46_.f12, self.f3, p1, d_2489_v108_, globalState))})
            d_2491_v110_: _dafny.Map
            d_2491_v110_ = _dafny.Map({False: (self).f2})
            d_2492_v111_: _dafny.Array
            nw360_ = _dafny.Array(None, 26)
            nw360_[int(0)] = d_2490_v109_
            nw360_[int(1)] = (d_2490_v109_).set(self.f3, p1)
            nw360_[int(2)] = (d_2490_v109_).set(self.f3, (self).fm12((self).f7, len(((d_2491_v110_)[self.f3] if (self.f3) in (d_2491_v110_) else (self).f2)), p0, self.f3, globalState))
            nw360_[int(3)] = d_2490_v109_
            nw360_[int(4)] = (d_2490_v109_).set(self.f3, (self).f7)
            nw360_[int(5)] = d_2490_v109_
            nw360_[int(6)] = (d_2490_v109_) | (d_2490_v109_)
            nw360_[int(7)] = (d_2490_v109_) | (d_2490_v109_)
            nw360_[int(8)] = d_2490_v109_
            nw360_[int(9)] = d_2490_v109_
            nw360_[int(10)] = d_2490_v109_
            nw360_[int(11)] = d_2490_v109_
            nw360_[int(12)] = d_2490_v109_
            nw360_[int(13)] = d_2490_v109_
            nw360_[int(14)] = _dafny.Map({(d_2366_v0_)[default__.safeIndex((self).f4, len(d_2366_v0_))]: (self).f4})
            nw360_[int(15)] = (d_2490_v109_) | (d_2490_v109_)
            nw360_[int(16)] = d_2490_v109_
            nw360_[int(17)] = d_2490_v109_
            nw360_[int(18)] = _dafny.Map({False: default__.fm0(globalState)})
            nw360_[int(19)] = d_2490_v109_
            nw360_[int(20)] = d_2490_v109_
            nw360_[int(21)] = (_dafny.Map({False: d_2412_v46_.f12})) | (d_2490_v109_)
            nw360_[int(22)] = d_2490_v109_
            nw360_[int(23)] = d_2490_v109_
            nw360_[int(24)] = (d_2490_v109_ if (d_2412_v46_).fm13(globalState) else _dafny.Map({False: (d_2412_v46_).f11}))
            nw360_[int(25)] = (_dafny.Map({self.f3: (0) - (len((self).f2))})) | (d_2490_v109_)
            d_2492_v111_ = nw360_
            index399_ = default__.safeIndex(869, (d_2492_v111_).length(0))
            (d_2492_v111_)[index399_] = (d_2490_v109_) | (d_2490_v109_)
            index400_ = default__.safeIndex(869, (d_2492_v111_).length(0))
            (d_2492_v111_)[index400_] = d_2490_v109_
            r1 = -593
            d_2493_v112_: C5
            nw361_ = C5()
            nw361_.ctor__(False, (self).f1, (self).f2)
            d_2493_v112_ = nw361_
            d_2494_v113_: _dafny.Map
            d_2494_v113_ = _dafny.Map({len((self).f2): d_2493_v112_})
            d_2495_v114_: _dafny.Set
            d_2495_v114_ = _dafny.Set({d_2494_v113_})
            d_2496_v115_: _dafny.MultiSet
            d_2496_v115_ = _dafny.MultiSet([len(_dafny.Set({(self).f4}))])
            d_2497_v116_: _dafny.Map
            d_2497_v116_ = _dafny.Map({(d_2495_v114_).intersection(d_2495_v114_): (d_2496_v115_).cardinality})
            r1 = ((d_2497_v116_)[d_2495_v114_] if (d_2495_v114_) in (d_2497_v116_) else d_2412_v46_.f12)
            d_2498_v117_: _dafny.Map
            d_2498_v117_ = _dafny.Map({495: default__.fm59(346, globalState)})
            d_2499_v118_: _dafny.Array
            def lambda97_(d_2500_v46_):
                def lambda98_(d_2501_i8_):
                    return default__.safeModuloInt(d_2501_i8_, d_2500_v46_.f12)

                return lambda98_

            init55_ = lambda97_(d_2412_v46_)
            nw362_ = _dafny.Array(None, 24)
            for i0_55_ in range(nw362_.length(0)):
                nw362_[i0_55_] = init55_(i0_55_)
            d_2499_v118_ = nw362_
            d_2502_v119_: _dafny.Map
            d_2502_v119_ = _dafny.Map({d_2499_v118_: (self).f5})
            d_2503_v120_: C9
            nw363_ = C9()
            nw363_.ctor__(((self).f5 if (d_2493_v112_).f15 else (self).f5), (d_2502_v119_) | (d_2502_v119_), default__.safeModuloInt(len(d_2410_v43_), p1), (self).f5, (d_2366_v0_)[default__.safeIndex((d_2412_v46_).f11, len(d_2366_v0_))], (self).f1, ((self).f2) + (_dafny.SeqWithoutIsStrInference([d_2489_v108_ for d_2504_i9_ in range(default__.abs(276))])))
            d_2503_v120_ = nw363_
            rhs426_ = d_2498_v117_
            rhs427_ = (d_2412_v46_).f11
            rhs428_ = (self.f3) == (self.f3)
            rhs429_ = self.f3
            rhs430_ = d_2503_v120_
            lhs276_ = self
            lhs277_ = self
            d_2498_v117_ = rhs426_
            r2 = rhs427_
            lhs276_.f3 = rhs428_
            lhs277_.f3 = rhs429_
            d_2503_v120_ = rhs430_
        d_2505_v121_: _dafny.Map
        d_2505_v121_ = _dafny.Map({self.f3: self.f3})
        r0 = (((d_2505_v121_).set(self.f3, self.f3)) | (d_2505_v121_)).set(self.f3, not (False) or (self.f3))
        d_2506_v122_: _dafny.MultiSet
        d_2506_v122_ = _dafny.MultiSet([p0])
        r1 = ((p1 if (self).fm11(self.f3, globalState) else (d_2506_v122_).cardinality) if not (self.f3) or (self.f3) else default__.safeDivisionInt((self).f4, d_2412_v46_.f12))
        d_2507_v123_: D16
        d_2507_v123_ = D16_DC42(self.f3, self.f3, self.f3)
        pat_let_tv56_ = p1
        def lambda99_(source43_):
            if source43_.is_DC42:
                d_2508___mcc_h0_ = source43_.cf70
                d_2509___mcc_h1_ = source43_.cf71
                d_2510___mcc_h2_ = source43_.cf72
                d_2511_cf72_ = d_2510___mcc_h2_
                d_2512_cf71_ = d_2509___mcc_h1_
                d_2513_cf70_ = d_2508___mcc_h0_
                return len((self).f2)
            elif True:
                d_2514___mcc_h3_ = source43_.cf69
                d_2515_cf69_ = d_2514___mcc_h3_
                return pat_let_tv56_

        r2 = lambda99_(d_2507_v123_)
        return r0, r1, r2

    def m3(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r3: _dafny.Array = _dafny.Array(None, 0)
        d_2516_v0_: _dafny.MultiSet
        d_2516_v0_ = _dafny.MultiSet([(self).f7, len((self).f2)])
        d_2517_v2_: _dafny.Set
        d_2517_v2_ = _dafny.Set({(0) - (len((_dafny.Map({self.f3: self.f3})).set(self.f3, False)))})
        d_2518_i0_: int
        d_2518_i0_ = 0
        with _dafny.label("14"):
            def iife218_():
                coll70_ = _dafny.Set()
                compr_70_: int
                for compr_70_ in (d_2516_v0_).Elements:
                    d_2524_v1_: int = compr_70_
                    if (d_2524_v1_) in (d_2516_v0_):
                        coll70_ = coll70_.union(_dafny.Set([(d_2524_v1_) - (811)]))
                return _dafny.Set(coll70_)
            while (iife218_()
            ).ispropersubset(d_2517_v2_):
                with _dafny.c_label("14"):
                    if (d_2518_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2518_i0_ = (d_2518_i0_) + (1)
                    d_2519_v3_: C5
                    nw364_ = C5()
                    nw364_.ctor__(self.f3, (self).f1, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2520_i1_ in range(default__.abs(-749))]))
                    d_2519_v3_ = nw364_
                    d_2521_v4_: _dafny.Seq
                    d_2521_v4_ = _dafny.SeqWithoutIsStrInference([self.f3, self.f3, self.f3])
                    d_2522_v5_: _dafny.Seq
                    d_2522_v5_ = _dafny.SeqWithoutIsStrInference([d_2521_v4_])
                    d_2521_v4_ = (d_2522_v5_)[default__.safeIndex(default__.safeDivisionInt(751, (self).f4), len(d_2522_v5_))]
                    r1 = True
                    d_2523_v6_: D3
                    d_2523_v6_ = D3_DC9(546, not(not(self.f3)))
                    (self).f3 = (d_2523_v6_).cf16
                    pass
            pass
        d_2525_v7_: _dafny.Map
        d_2525_v7_ = _dafny.Map({self.f3: False})
        d_2526_v8_: int
        d_2526_v8_ = 154
        d_2527_v9_: _dafny.MultiSet
        d_2527_v9_ = _dafny.MultiSet([self.f3, False])
        rhs431_ = d_2525_v7_
        rhs432_ = (self).f4
        rhs433_ = default__.safeModuloInt(((d_2516_v0_).cardinality) + ((self).f4), (self).f7)
        rhs434_ = ((d_2527_v9_) | (d_2527_v9_)).issubset(_dafny.MultiSet([self.f3, self.f3]))
        d_2525_v7_ = rhs431_
        d_2526_v8_ = rhs432_
        d_2526_v8_ = rhs433_
        r1 = rhs434_
        d_2528_v10_: D8
        d_2528_v10_ = D8_DC23((self).f7, (self).f4, False)
        d_2529_v11_: _dafny.Seq
        d_2529_v11_ = _dafny.SeqWithoutIsStrInference([d_2528_v10_])
        if ((_dafny.SeqWithoutIsStrInference([d_2528_v10_, d_2528_v10_])) + (d_2529_v11_)) != (d_2529_v11_):
            d_2530_v12_: _dafny.Map
            d_2530_v12_ = _dafny.Map({self.f3: d_2526_v8_})
            d_2531_v13_: _dafny.Array
            def lambda100_(d_2532_i2_):
                return default__.safeDivisionInt(d_2532_i2_, (self).f4)

            init56_ = lambda100_
            nw365_ = _dafny.Array(None, 29)
            for i0_56_ in range(nw365_.length(0)):
                nw365_[i0_56_] = init56_(i0_56_)
            d_2531_v13_ = nw365_
            d_2533_v14_: D5
            d_2533_v14_ = D5_DC14(d_2531_v13_)
            d_2534_v15_: _dafny.Seq
            d_2534_v15_ = _dafny.SeqWithoutIsStrInference([d_2533_v14_, d_2533_v14_])
            if (((d_2530_v12_)[self.f3] if (self.f3) in (d_2530_v12_) else (self).f7)) != (len(d_2534_v15_)):
                d_2535_v16_: _dafny.Map
                d_2535_v16_ = _dafny.Map({self.f3: default__.fm9(globalState)})
                rhs435_ = d_2535_v16_
                rhs436_ = d_2534_v15_
                d_2535_v16_ = rhs435_
                d_2534_v15_ = rhs436_
                r1 = self.f3
                d_2536_v17_: _dafny.Map
                d_2536_v17_ = _dafny.Map({(self).f5: self.f3})
                d_2537_v18_: bool
                d_2538_v19_: bool
                d_2539_v20_: _dafny.Seq
                d_2540_v21_: _dafny.Seq
                out66_: bool
                out67_: bool
                out68_: _dafny.Seq
                out69_: _dafny.Seq
                out66_, out67_, out68_, out69_ = default__.m0(d_2536_v17_, globalState)
                d_2537_v18_ = out66_
                d_2538_v19_ = out67_
                d_2539_v20_ = out68_
                d_2540_v21_ = out69_
                d_2541_v22_: _dafny.Map
                d_2541_v22_ = _dafny.Map({not(d_2538_v19_): d_2531_v13_})
                d_2542_v23_: C0
                nw366_ = C0()
                nw366_.ctor__(((d_2541_v22_)[self.f3] if (self.f3) in (d_2541_v22_) else d_2531_v13_))
                d_2542_v23_ = nw366_
                r2 = ((self).f2) + ((self).f2)
            elif True:
                index401_ = default__.safeIndex(407, (d_2531_v13_).length(0))
                (d_2531_v13_)[index401_] = (self).f4
                index402_ = default__.safeIndex(407, (d_2531_v13_).length(0))
                rhs437_ = self.f3
                rhs438_ = (d_2516_v0_ if self.f3 else d_2516_v0_)
                rhs439_ = (self).f4
                rhs440_ = (d_2526_v8_) * (len((self).f2))
                rhs441_ = self.f3
                lhs278_ = self
                lhs279_ = d_2531_v13_
                lhs280_ = default__.safeIndex(407, (d_2531_v13_).length(0))
                lhs278_.f3 = rhs437_
                d_2516_v0_ = rhs438_
                d_2526_v8_ = rhs439_
                lhs279_[lhs280_] = rhs440_
                r1 = rhs441_
                d_2543_v25_: _dafny.Seq
                d_2543_v25_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_2544_v26_: _dafny.Map
                def iife219_():
                    coll71_ = _dafny.Map()
                    compr_71_: _dafny.Seq
                    for compr_71_ in (d_2543_v25_).Elements:
                        d_2545_v24_: _dafny.Seq = compr_71_
                        if (d_2545_v24_) in (d_2543_v25_):
                            coll71_[d_2545_v24_] = d_2526_v8_
                    return _dafny.Map(coll71_)
                d_2544_v26_ = _dafny.Map({(0) - ((self).f4): iife219_()
                })
                d_2546_v27_: _dafny.Map
                d_2546_v27_ = _dafny.Map({(self).f7: d_2530_v12_})
                d_2547_v28_: _dafny.Map
                d_2547_v28_ = _dafny.Map({d_2546_v27_: (self).f4})
                d_2548_v29_: _dafny.Seq
                d_2548_v29_ = p0
                d_2549_v30_: _dafny.Map
                d_2549_v30_ = _dafny.Map({d_2548_v29_: 103})
                d_2544_v26_ = (d_2544_v26_).set(((self).f7) - (((d_2547_v28_)[d_2546_v27_] if (d_2546_v27_) in (d_2547_v28_) else (self).f7)), (d_2549_v30_).set(d_2548_v29_, (d_2531_v13_)[default__.safeIndex(407, (d_2531_v13_).length(0))]))
                d_2550_v31_: _dafny.Set
                d_2550_v31_ = d_2517_v2_
                d_2551_v32_: _dafny.Set
                d_2551_v32_ = _dafny.Set({d_2550_v31_})
                d_2552_v33_: _dafny.Map
                d_2552_v33_ = _dafny.Map({(self).f5: len(d_2551_v32_)})
                d_2553_v34_: D5
                d_2553_v34_ = D5_DC15((d_2531_v13_)[default__.safeIndex(407, (d_2531_v13_).length(0))], self.f3, (self).f5)
                d_2554_v37_: _dafny.Map
                d_2554_v37_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2555_i3_ in range(default__.abs(-224))]): d_2526_v8_})
                d_2556_v38_: T2
                nw367_ = C7()
                def iife220_(_pat_let74_0):
                    def iife221_(d_2557_dt__update__tmp_h0_):
                        def iife222_(_pat_let75_0):
                            def iife223_(d_2558_dt__update_hcf25_h0_):
                                return D5_DC15((d_2557_dt__update__tmp_h0_).cf24, d_2558_dt__update_hcf25_h0_, (d_2557_dt__update__tmp_h0_).cf26)
                            return iife223_(_pat_let75_0)
                        return iife222_(self.f3)
                    return iife221_(_pat_let74_0)
                def iife224_():
                    def iife226_():
                        coll74_ = _dafny.Map()
                        compr_74_: _dafny.Seq
                        for compr_74_ in (d_2554_v37_).keys.Elements:
                            d_2559_v36_: _dafny.Seq = compr_74_
                            if (d_2559_v36_) in (d_2554_v37_):
                                coll74_[d_2559_v36_] = (self).f7
                        return _dafny.Map(coll74_)
                    coll72_ = _dafny.Map()
                    def iife225_():
                        coll73_ = _dafny.Map()
                        compr_73_: _dafny.Seq
                        for compr_73_ in (d_2554_v37_).keys.Elements:
                            d_2559_v36_: _dafny.Seq = compr_73_
                            if (d_2559_v36_) in (d_2554_v37_):
                                coll73_[d_2559_v36_] = (self).f7
                        return _dafny.Map(coll73_)
                    compr_72_: _dafny.Seq
                    for compr_72_ in (iife225_()
                    ).keys.Elements:
                        d_2560_v35_: _dafny.Seq = compr_72_
                        if (d_2560_v35_) in (iife226_()
                        ):
                            coll72_[d_2560_v35_] = not(self.f3)
                    return _dafny.Map(coll72_)
                nw367_.ctor__((self).f7, (iife220_(d_2553_v34_)).cf26, self.f3, iife224_()
                , (self).f2)
                d_2556_v38_ = nw367_
                d_2561_v39_: _dafny.Map
                d_2561_v39_ = _dafny.Map({p0: d_2556_v38_})
                rhs442_ = _dafny.Set({(len(d_2552_v33_)) + (d_2526_v8_)})
                rhs443_ = self.f3
                rhs444_ = (not((p0)[default__.safeIndex((d_2556_v38_).f4, len(p0))]) if (len(d_2561_v39_)) <= ((0) - ((d_2516_v0_).cardinality)) else self.f3)
                lhs281_ = self
                lhs282_ = self
                d_2517_v2_ = rhs442_
                lhs281_.f3 = rhs443_
                lhs282_.f3 = rhs444_
                d_2562_v40_: _dafny.Set
                d_2562_v40_ = _dafny.Set({not(d_2556_v38_.f3)})
                d_2562_v40_ = d_2562_v40_
                d_2563_v41_: _dafny.Array
                nw368_ = _dafny.Array(int(0), 22)
                d_2563_v41_ = nw368_
                d_2564_v42_: C0
                nw369_ = C0()
                nw369_.ctor__(d_2563_v41_)
                d_2564_v42_ = nw369_
            d_2526_v8_ = (d_2526_v8_) + (len((self).f2))
            r2 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv"))
            d_2526_v8_ = len((_dafny.Map({(self).f5: self.f3})).set((self).f5, self.f3))
            d_2565_v43_: _dafny.Seq
            d_2565_v43_ = _dafny.SeqWithoutIsStrInference([d_2531_v13_, d_2531_v13_, d_2531_v13_, d_2531_v13_])
            d_2566_v44_: _dafny.Seq
            d_2566_v44_ = _dafny.SeqWithoutIsStrInference([(self).f7])
            d_2526_v8_ = (len((_dafny.SeqWithoutIsStrInference([d_2531_v13_, d_2531_v13_, (d_2565_v43_)[default__.safeIndex((self).f4, len(d_2565_v43_))]])).set(default__.safeIndex(len(_dafny.Map({d_2531_v13_: (d_2566_v44_).set(default__.safeIndex(d_2526_v8_, len(d_2566_v44_)), (self).f4)})), len(_dafny.SeqWithoutIsStrInference([d_2531_v13_, d_2531_v13_, (d_2565_v43_)[default__.safeIndex((self).f4, len(d_2565_v43_))]]))), d_2531_v13_))) * (default__.safeDivisionInt((d_2566_v44_)[default__.safeIndex((self).f4, len(d_2566_v44_))], d_2526_v8_))
        elif True:
            d_2567_v45_: _dafny.Map
            d_2567_v45_ = _dafny.Map({(self).f2: d_2526_v8_})
            if (default__.fm5(d_2567_v45_, globalState)) != (d_2527_v9_):
                d_2568_v46_: C2
                nw370_ = C2()
                nw370_.ctor__((self).f7, self.f3, (self).f1, (self).f2)
                d_2568_v46_ = nw370_
                d_2569_v47_: D13
                d_2569_v47_ = D13_DC33(d_2568_v46_)
                d_2570_v48_: _dafny.Map
                d_2570_v48_ = _dafny.Map({d_2569_v47_: (d_2568_v46_).f13})
                d_2571_v49_: str
                d_2571_v49_ = _dafny.CodePoint('r')
                d_2572_v50_: D13
                d_2572_v50_ = D13_DC34(self.f3, self.f3, (d_2568_v46_).f13, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmtl")))
                d_2573_v51_: _dafny.Array
                def lambda101_(d_2574_v46_):
                    def lambda102_(d_2575_i4_):
                        return (d_2575_i4_) * ((d_2574_v46_).f13)

                    return lambda102_

                init57_ = lambda101_(d_2568_v46_)
                nw371_ = _dafny.Array(None, 12)
                for i0_57_ in range(nw371_.length(0)):
                    nw371_[i0_57_] = init57_(i0_57_)
                d_2573_v51_ = nw371_
                d_2576_v52_: C0
                nw372_ = C0()
                nw372_.ctor__(d_2573_v51_)
                d_2576_v52_ = nw372_
                d_2577_v53_: _dafny.Map
                d_2577_v53_ = _dafny.Map({d_2576_v52_: self.f3})
                d_2578_v54_: _dafny.Seq
                d_2578_v54_ = _dafny.SeqWithoutIsStrInference([(d_2568_v46_).f13, (self).f7, (self).f4])
                d_2579_v55_: _dafny.Set
                d_2579_v55_ = _dafny.Set({(self).f2})
                d_2580_v56_: _dafny.Array
                nw373_ = _dafny.Array(None, 15)
                nw373_[int(0)] = (0) - ((0) - ((self).f4))
                nw373_[int(1)] = len(d_2570_v48_)
                nw373_[int(2)] = len((self).fm14(self.f3, self.f3, d_2571_v49_, d_2526_v8_, globalState))
                nw373_[int(3)] = (self).f7
                nw373_[int(4)] = len((d_2572_v50_).cf56)
                nw373_[int(5)] = (self).f4
                nw373_[int(6)] = (len(d_2577_v53_)) + (d_2526_v8_)
                nw373_[int(7)] = (d_2578_v54_)[default__.safeIndex((self).f7, len(d_2578_v54_))]
                nw373_[int(8)] = ((d_2568_v46_).f13) + ((self).f4)
                nw373_[int(9)] = (len(d_2579_v55_)) - ((d_2568_v46_).f13)
                nw373_[int(10)] = d_2526_v8_
                nw373_[int(11)] = (d_2526_v8_) - (len((self).f2))
                nw373_[int(12)] = (self).f4
                nw373_[int(13)] = (d_2568_v46_).f13
                nw373_[int(14)] = -352
                d_2580_v56_ = nw373_
                index403_ = default__.safeIndex(833, ((d_2576_v52_).f14).length(0))
                ((d_2576_v52_).f14)[index403_] = len((self).f2)
                d_2581_v58_: _dafny.Seq
                def iife227_():
                    coll75_ = _dafny.Set()
                    compr_75_: int
                    for compr_75_ in (_dafny.SeqWithoutIsStrInference([d_2526_v8_ for d_2582_i5_ in range(default__.abs(16))])).Elements:
                        d_2583_v57_: int = compr_75_
                        if (d_2583_v57_) in (_dafny.SeqWithoutIsStrInference([d_2526_v8_ for d_2582_i5_ in range(default__.abs(16))])):
                            coll75_ = coll75_.union(_dafny.Set([default__.safeModuloInt(d_2583_v57_, -509)]))
                    return _dafny.Set(coll75_)
                d_2581_v58_ = _dafny.SeqWithoutIsStrInference([iife227_()
                , (d_2517_v2_) | (d_2517_v2_), default__.fm33(len(d_2578_v54_), globalState)])
                index404_ = default__.safeIndex(833, ((d_2576_v52_).f14).length(0))
                ((d_2576_v52_).f14)[index404_] = len(d_2581_v58_)
                d_2584_v59_: _dafny.Seq
                d_2584_v59_ = p0
                d_2585_v60_: _dafny.Seq
                d_2585_v60_ = _dafny.SeqWithoutIsStrInference([(d_2584_v59_)])
                d_2586_v61_: T1
                nw374_ = C8()
                nw374_.ctor__(d_2585_v60_, (p0)[default__.safeIndex(((d_2576_v52_).f14)[default__.safeIndex(833, ((d_2576_v52_).f14).length(0))], len(p0))], (self).f1, (self).f2, ((self).f7) * (len(d_2517_v2_)), (self).f5)
                d_2586_v61_ = nw374_
                rhs445_ = (d_2576_v52_).f14
                rhs446_ = d_2586_v61_
                rhs447_ = ((d_2516_v0_).cardinality if (len((d_2586_v61_).f2)) < (871) else len((d_2586_v61_).f2))
                rhs448_ = d_2586_v61_.f3
                d_2580_v56_ = rhs445_
                d_2586_v61_ = rhs446_
                d_2526_v8_ = rhs447_
                r1 = rhs448_
                d_2587_v62_: _dafny.Array
                nw375_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_2587_v62_ = nw375_
                index405_ = default__.safeIndex(330, (d_2587_v62_).length(0))
                (d_2587_v62_)[index405_] = _dafny.SeqWithoutIsStrInference([d_2571_v49_ for d_2588_i6_ in range(default__.abs(-3))])
                index406_ = default__.safeIndex(330, (d_2587_v62_).length(0))
                (d_2587_v62_)[index406_] = (d_2586_v61_).f2
                d_2589_v63_: _dafny.Map
                d_2589_v63_ = _dafny.Map({d_2578_v54_: default__.fm3(d_2526_v8_, globalState)})
                d_2590_v65_: _dafny.Set
                d_2590_v65_ = _dafny.Set({d_2578_v54_, _dafny.SeqWithoutIsStrInference([(self).f4])})
                d_2591_v67_: _dafny.Seq
                d_2591_v67_ = _dafny.SeqWithoutIsStrInference([d_2589_v63_, d_2589_v63_, d_2589_v63_])
                d_2592_v71_: _dafny.Array
                nw376_ = _dafny.Array(None, 22)
                nw376_[int(0)] = d_2589_v63_
                nw376_[int(1)] = _dafny.Map({d_2578_v54_: d_2586_v61_.f3})
                def iife228_():
                    coll76_ = _dafny.Map()
                    compr_76_: _dafny.Seq
                    for compr_76_ in (d_2590_v65_).Elements:
                        d_2593_v64_: _dafny.Seq = compr_76_
                        if (d_2593_v64_) in (d_2590_v65_):
                            coll76_[d_2593_v64_] = self.f3
                    return _dafny.Map(coll76_)
                nw376_[int(2)] = iife228_()
                
                nw376_[int(3)] = d_2589_v63_
                nw376_[int(4)] = (d_2589_v63_) | (_dafny.Map({d_2578_v54_: self.f3}))
                nw376_[int(5)] = (d_2589_v63_) | (d_2589_v63_)
                nw376_[int(6)] = d_2589_v63_
                nw376_[int(7)] = (d_2589_v63_) | (d_2589_v63_)
                nw376_[int(8)] = d_2589_v63_
                nw376_[int(9)] = (_dafny.Map({d_2578_v54_: not(d_2586_v61_.f3)})) | ((d_2589_v63_).set(default__.fm31(d_2526_v8_, self.f3, default__.fm0(globalState), d_2571_v49_, globalState), False))
                def iife229_():
                    coll77_ = _dafny.Map()
                    compr_77_: _dafny.Seq
                    for compr_77_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([364])})).Elements:
                        d_2594_v66_: _dafny.Seq = compr_77_
                        if (d_2594_v66_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([364])})):
                            coll77_[d_2594_v66_] = d_2586_v61_.f3
                    return _dafny.Map(coll77_)
                nw376_[int(10)] = (d_2589_v63_) | (iife229_()
                )
                nw376_[int(11)] = d_2589_v63_
                nw376_[int(12)] = d_2589_v63_
                nw376_[int(13)] = d_2589_v63_
                nw376_[int(14)] = (d_2591_v67_)[default__.safeIndex((d_2568_v46_).f13, len(d_2591_v67_))]
                def iife230_():
                    def iife232_():
                        coll80_ = _dafny.Map()
                        compr_80_: _dafny.Seq
                        for compr_80_ in (d_2589_v63_).keys.Elements:
                            d_2595_v69_: _dafny.Seq = compr_80_
                            if (d_2595_v69_) in (d_2589_v63_):
                                coll80_[d_2595_v69_] = d_2526_v8_
                        return _dafny.Map(coll80_)
                    coll78_ = _dafny.Map()
                    def iife231_():
                        coll79_ = _dafny.Map()
                        compr_79_: _dafny.Seq
                        for compr_79_ in (d_2589_v63_).keys.Elements:
                            d_2595_v69_: _dafny.Seq = compr_79_
                            if (d_2595_v69_) in (d_2589_v63_):
                                coll79_[d_2595_v69_] = d_2526_v8_
                        return _dafny.Map(coll79_)
                    compr_78_: _dafny.Seq
                    for compr_78_ in (iife231_()
                    ).keys.Elements:
                        d_2596_v68_: _dafny.Seq = compr_78_
                        if (d_2596_v68_) in (iife232_()
                        ):
                            coll78_[d_2596_v68_] = self.f3
                    return _dafny.Map(coll78_)
                nw376_[int(15)] = (d_2589_v63_) | (iife230_()
                )
                def iife233_():
                    coll81_ = _dafny.Map()
                    compr_81_: _dafny.Seq
                    for compr_81_ in (_dafny.SeqWithoutIsStrInference([d_2578_v54_ for d_2597_i7_ in range(default__.abs(705))])).Elements:
                        d_2598_v70_: _dafny.Seq = compr_81_
                        if (d_2598_v70_) in (_dafny.SeqWithoutIsStrInference([d_2578_v54_ for d_2597_i7_ in range(default__.abs(705))])):
                            coll81_[d_2598_v70_] = not(self.f3)
                    return _dafny.Map(coll81_)
                nw376_[int(16)] = iife233_()
                
                nw376_[int(17)] = (d_2589_v63_) | ((d_2589_v63_).set(d_2578_v54_, d_2586_v61_.f3))
                nw376_[int(18)] = (d_2589_v63_) | (d_2589_v63_)
                nw376_[int(19)] = (d_2589_v63_) | ((d_2591_v67_)[default__.safeIndex(((d_2576_v52_).f14)[default__.safeIndex(833, ((d_2576_v52_).f14).length(0))], len(d_2591_v67_))])
                nw376_[int(20)] = d_2589_v63_
                nw376_[int(21)] = d_2589_v63_
                d_2592_v71_ = nw376_
                d_2599_v72_: _dafny.Seq
                d_2599_v72_ = _dafny.SeqWithoutIsStrInference([d_2578_v54_, d_2578_v54_])
                index407_ = default__.safeIndex(983, (d_2592_v71_).length(0))
                (d_2592_v71_)[index407_] = ((d_2589_v63_).set((d_2599_v72_)[default__.safeIndex((self).f4, len(d_2599_v72_))], self.f3)) | (d_2589_v63_)
                d_2600_v74_: _dafny.MultiSet
                d_2600_v74_ = _dafny.MultiSet([d_2578_v54_, d_2578_v54_])
                index408_ = default__.safeIndex(983, (d_2592_v71_).length(0))
                def iife234_():
                    coll82_ = _dafny.Map()
                    compr_82_: _dafny.Seq
                    for compr_82_ in ((d_2600_v74_) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2578_v54_, d_2578_v54_])))).Elements:
                        d_2601_v73_: _dafny.Seq = compr_82_
                        if (d_2601_v73_) in ((d_2600_v74_) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2578_v54_, d_2578_v54_])))):
                            coll82_[d_2601_v73_] = (d_2572_v50_).cf53
                    return _dafny.Map(coll82_)
                (d_2592_v71_)[index408_] = iife234_()
                
                d_2602_v75_: _dafny.Array
                nw377_ = _dafny.Array(None, 8)
                d_2602_v75_ = nw377_
                d_2603_v76_: C5
                nw378_ = C5()
                nw378_.ctor__(self.f3, (d_2586_v61_).f1, (self).f2)
                d_2603_v76_ = nw378_
                index409_ = default__.safeIndex(721, (d_2602_v75_).length(0))
                (d_2602_v75_)[index409_] = d_2603_v76_
                index410_ = default__.safeIndex(721, (d_2602_v75_).length(0))
                nw379_ = C5()
                nw379_.ctor__(d_2586_v61_.f3, _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcotffvl")): not(d_2586_v61_.f3)}), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpehooaj"))) + ((self).fm14((d_2603_v76_).f15, True, d_2571_v49_, ((d_2576_v52_).f14)[default__.safeIndex(833, ((d_2576_v52_).f14).length(0))], globalState)))
                (d_2602_v75_)[index410_] = nw379_
            elif True:
                d_2526_v8_ = default__.safeModuloInt((self).f4, len((_dafny.SeqWithoutIsStrInference([(self).f5])) + (_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5]))))
                d_2604_v77_: D13
                d_2604_v77_ = D13_DC34(self.f3, ((d_2525_v7_)[self.f3] if (self.f3) in (d_2525_v7_) else False), (self).f7, (self).f2)
                d_2605_v78_: C2
                nw380_ = C2()
                nw380_.ctor__(d_2526_v8_, self.f3, (self).f1, (d_2604_v77_).cf56)
                d_2605_v78_ = nw380_
                d_2606_v79_: _dafny.Array
                nw381_ = _dafny.Array(int(0), 26)
                d_2606_v79_ = nw381_
                index411_ = default__.safeIndex(549, (d_2606_v79_).length(0))
                (d_2606_v79_)[index411_] = d_2526_v8_
                index412_ = default__.safeIndex(549, (d_2606_v79_).length(0))
                rhs449_ = d_2605_v78_
                rhs450_ = (0) - ((0) - ((self).f7))
                lhs283_ = d_2606_v79_
                lhs284_ = default__.safeIndex(549, (d_2606_v79_).length(0))
                d_2605_v78_ = rhs449_
                lhs283_[lhs284_] = rhs450_
                d_2607_v80_: _dafny.Array
                nw382_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_2607_v80_ = nw382_
                d_2607_v80_ = d_2607_v80_
                d_2608_v81_: _dafny.Array
                nw383_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_2608_v81_ = nw383_
                d_2609_v82_: str
                d_2609_v82_ = _dafny.CodePoint('j')
                index413_ = default__.safeIndex(526, (d_2608_v81_).length(0))
                (d_2608_v81_)[index413_] = d_2609_v82_
                index414_ = default__.safeIndex(526, (d_2608_v81_).length(0))
                rhs451_ = d_2609_v82_
                rhs452_ = d_2608_v81_
                rhs453_ = (p0)[default__.safeIndex(len((self).f2), len(p0))]
                rhs454_ = (((self).f2) + ((self).f2)) + (((self).f2) + ((self).f2))
                lhs285_ = d_2608_v81_
                lhs286_ = default__.safeIndex(526, (d_2608_v81_).length(0))
                lhs285_[lhs286_] = rhs451_
                d_2608_v81_ = rhs452_
                r1 = rhs453_
                r2 = rhs454_
                d_2610_v83_: D14
                d_2610_v83_ = D14_DC37(d_2608_v81_)
                index415_ = default__.safeIndex(549, (d_2606_v79_).length(0))
                rhs455_ = d_2610_v83_
                rhs456_ = (d_2605_v78_).f13
                lhs287_ = d_2606_v79_
                lhs288_ = default__.safeIndex(549, (d_2606_v79_).length(0))
                d_2610_v83_ = rhs455_
                lhs287_[lhs288_] = rhs456_
            (self).f3 = (((self).f4) - (d_2526_v8_)) < ((d_2527_v9_).cardinality)
            if not(((0) - ((self).f4)) >= ((len((self).f2)) * (-120))):
                d_2611_v84_: _dafny.Seq
                d_2611_v84_ = _dafny.SeqWithoutIsStrInference([(d_2526_v8_) < ((self).f4)])
                d_2611_v84_ = d_2611_v84_
                (self).f3 = ((d_2527_v9_).cardinality) > (d_2526_v8_)
                d_2612_v85_: str
                d_2612_v85_ = _dafny.CodePoint('x')
                d_2612_v85_ = d_2612_v85_
                d_2613_v86_: C3
                nw384_ = C3()
                nw384_.ctor__((self).f4, (self).f5, (320) != (d_2526_v8_), (self).f1, (self).f2)
                d_2613_v86_ = nw384_
                d_2614_v87_: _dafny.Array
                nw385_ = _dafny.Array(False, 6)
                d_2614_v87_ = nw385_
                d_2614_v87_ = (self).f5
            elif True:
                d_2526_v8_ = (self).f7
                d_2615_v88_: _dafny.Map
                d_2615_v88_ = _dafny.Map({self.f3: _dafny.CodePoint('o')})
                d_2616_v89_: str
                d_2616_v89_ = _dafny.CodePoint('l')
                d_2615_v88_ = (d_2615_v88_).set(((self).f7) == ((self).f4), d_2616_v89_)
                d_2617_v90_: _dafny.Map
                d_2617_v90_ = _dafny.Map({((self).f2 if self.f3 else (self).f2): (default__.fm60(globalState))})
                d_2617_v90_ = (d_2617_v90_) | (d_2617_v90_)
                d_2618_v91_: _dafny.Seq
                d_2618_v91_ = _dafny.SeqWithoutIsStrInference([default__.fm9(globalState), default__.fm9(globalState)])
                d_2619_v92_: _dafny.Map
                d_2619_v92_ = _dafny.Map({d_2618_v91_: self.f3})
                d_2619_v92_ = (d_2619_v92_).set(d_2618_v91_, self.f3)
                d_2526_v8_ = (self).f4
            r2 = (((self).f2) + ((self).f2)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghwey")))
            d_2620_v93_: _dafny.Array
            nw386_ = _dafny.Array(None, 1)
            d_2620_v93_ = nw386_
            d_2620_v93_ = d_2620_v93_
        d_2621_i8_: int
        d_2621_i8_ = 0
        with _dafny.label("15"):
            while self.f3:
                with _dafny.c_label("15"):
                    if (d_2621_i8_) >= (100):
                        raise _dafny.Break("15")
                    d_2621_i8_ = (d_2621_i8_) + (1)
                    d_2622_v94_: _dafny.Array
                    nw387_ = _dafny.Array(None, 23)
                    d_2622_v94_ = nw387_
                    d_2623_v95_: C3
                    nw388_ = C3()
                    nw388_.ctor__((self).f7, (self).f5, self.f3, (self).f1, (self).f2)
                    d_2623_v95_ = nw388_
                    index416_ = default__.safeIndex(746, (d_2622_v94_).length(0))
                    (d_2622_v94_)[index416_] = (d_2623_v95_ if self.f3 else d_2623_v95_)
                    index417_ = default__.safeIndex(746, (d_2622_v94_).length(0))
                    (d_2622_v94_)[index417_] = d_2623_v95_
                    r1 = self.f3
                    d_2624_v96_: str
                    d_2624_v96_ = _dafny.CodePoint('q')
                    d_2625_v97_: _dafny.Map
                    d_2625_v97_ = _dafny.Map({self.f3: d_2624_v96_})
                    d_2626_v98_: D2
                    d_2626_v98_ = D2_DC6(d_2624_v96_, self.f3, True, 516)
                    d_2625_v97_ = (d_2625_v97_).set(not((self.f3) == (True)), (d_2626_v98_).cf9)
                    d_2627_v99_: _dafny.Set
                    d_2627_v99_ = _dafny.Set({(self).f7})
                    d_2628_v100_: _dafny.Set
                    d_2628_v100_ = _dafny.Set({d_2624_v96_})
                    d_2629_v101_: _dafny.Set
                    d_2629_v101_ = d_2628_v100_
                    rhs457_ = d_2627_v99_
                    rhs458_ = d_2629_v101_
                    d_2627_v99_ = rhs457_
                    d_2629_v101_ = rhs458_
                    pass
            pass
        d_2526_v8_ = (((self).f7) * ((self).f7)) - (((self).f7 if self.f3 else (self).f4))
        d_2630_v102_: D13
        d_2630_v102_ = D13_DC34(self.f3, self.f3, (self).f7, (self).f2)
        d_2631_v103_: _dafny.Map
        d_2631_v103_ = _dafny.Map({d_2630_v102_: self.f3})
        d_2632_v105_: _dafny.Map
        d_2632_v105_ = _dafny.Map({d_2526_v8_: (self).f4})
        d_2633_v106_: _dafny.Map
        d_2633_v106_ = _dafny.Map({d_2632_v105_: d_2631_v103_})
        d_2634_v107_: D18
        d_2634_v107_ = D18_DC44(d_2631_v103_)
        d_2635_v108_: _dafny.Array
        nw389_ = _dafny.Array(None, 24)
        nw389_[int(0)] = d_2631_v103_
        nw389_[int(1)] = d_2631_v103_
        def iife235_():
            coll83_ = _dafny.Map()
            compr_83_: D13
            for compr_83_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2630_v102_]))).Elements:
                d_2636_v104_: D13 = compr_83_
                if (d_2636_v104_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2630_v102_]))):
                    coll83_[d_2636_v104_] = not(self.f3)
            return _dafny.Map(coll83_)
        nw389_[int(2)] = iife235_()
        
        nw389_[int(3)] = (d_2631_v103_) | (((d_2633_v106_)[d_2632_v105_] if (d_2632_v105_) in (d_2633_v106_) else d_2631_v103_))
        nw389_[int(4)] = _dafny.Map({d_2630_v102_: False})
        nw389_[int(5)] = d_2631_v103_
        nw389_[int(6)] = (d_2631_v103_).set(d_2630_v102_, self.f3)
        nw389_[int(7)] = d_2631_v103_
        nw389_[int(8)] = (d_2631_v103_) | (d_2631_v103_)
        nw389_[int(9)] = d_2631_v103_
        nw389_[int(10)] = (d_2631_v103_).set(d_2630_v102_, self.f3)
        nw389_[int(11)] = d_2631_v103_
        nw389_[int(12)] = (d_2631_v103_) | (_dafny.Map({d_2630_v102_: self.f3}))
        nw389_[int(13)] = d_2631_v103_
        nw389_[int(14)] = ((d_2634_v107_).cf74 if self.f3 else d_2631_v103_)
        nw389_[int(15)] = d_2631_v103_
        nw389_[int(16)] = d_2631_v103_
        nw389_[int(17)] = d_2631_v103_
        nw389_[int(18)] = _dafny.Map({d_2630_v102_: not(self.f3)})
        nw389_[int(19)] = (d_2631_v103_) | (d_2631_v103_)
        nw389_[int(20)] = d_2631_v103_
        nw389_[int(21)] = d_2631_v103_
        nw389_[int(22)] = (d_2631_v103_) | ((_dafny.Map({d_2630_v102_: self.f3})).set(d_2630_v102_, self.f3))
        nw389_[int(23)] = d_2631_v103_
        d_2635_v108_ = nw389_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_2635_v108_).length(0)):
            d_2637_i9_: int = guard_loop_7_
            if (True) and (((0) <= (d_2637_i9_)) and ((d_2637_i9_) < ((d_2635_v108_).length(0)))):
                (d_2635_v108_)[(d_2637_i9_)] = d_2631_v103_
        d_2638_v109_: str
        d_2638_v109_ = _dafny.CodePoint('v')
        r0 = default__.fm31((self).f7, self.f3, default__.safeModuloInt(794, 979), d_2638_v109_, globalState)
        r1 = not (default__.fm3(d_2526_v8_, globalState)) or (True)
        r2 = ((self).f2) + ((self).f2)
        d_2639_v110_: _dafny.Array
        nw390_ = _dafny.Array(int(0), 22)
        d_2639_v110_ = nw390_
        r3 = d_2639_v110_
        return r0, r1, r2, r3

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        d_2640_v0_: str
        d_2640_v0_ = _dafny.CodePoint('k')
        d_2640_v0_ = d_2640_v0_
        d_2641_v1_: _dafny.Map
        d_2641_v1_ = _dafny.Map({self.f3: p0})
        hi9_ = len(d_2641_v1_)
        for d_2642_i0_ in range((self).f7, hi9_):
            d_2643_v2_: _dafny.Map
            d_2643_v2_ = _dafny.Map({default__.fm6(self.f3, _dafny.CodePoint('j'), globalState): self.f3})
            index418_ = default__.safeIndex(337, ((self).f5).length(0))
            ((self).f5)[index418_] = self.f3
            index419_ = default__.safeIndex(337, ((self).f5).length(0))
            rhs459_ = default__.safeModuloInt(default__.safeDivisionInt(p0, (self).f7), (self).f7)
            rhs460_ = d_2643_v2_
            rhs461_ = self.f3
            lhs289_ = (self).f5
            lhs290_ = default__.safeIndex(337, ((self).f5).length(0))
            r0 = rhs459_
            d_2643_v2_ = rhs460_
            lhs289_[lhs290_] = rhs461_
            d_2640_v0_ = d_2640_v0_
            d_2644_v4_: _dafny.Seq
            def iife236_():
                coll84_ = _dafny.Set()
                compr_84_: int
                for compr_84_ in _dafny.IntegerRange(998, 916):
                    d_2645_v3_: int = compr_84_
                    if ((998) <= (d_2645_v3_)) and ((d_2645_v3_) < (916)):
                        coll84_ = coll84_.union(_dafny.Set([default__.safeModuloInt(d_2645_v3_, p0)]))
                return _dafny.Set(coll84_)
            d_2644_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm0(globalState), len(iife236_()
            )])
            d_2646_v5_: _dafny.Seq
            d_2646_v5_ = _dafny.SeqWithoutIsStrInference([len(d_2644_v4_), p0])
            d_2647_v6_: _dafny.Set
            d_2647_v6_ = _dafny.Set({len(d_2646_v5_), d_2642_i0_})
            d_2648_v7_: _dafny.Seq
            d_2648_v7_ = _dafny.SeqWithoutIsStrInference([self.f3])
            d_2649_v8_: _dafny.Array
            nw391_ = _dafny.Array(False, 3)
            d_2649_v8_ = nw391_
            d_2650_v9_: C8
            nw392_ = C8()
            nw392_.ctor__(_dafny.SeqWithoutIsStrInference([d_2648_v7_, d_2648_v7_]), ((self).f5)[default__.safeIndex(337, ((self).f5).length(0))], (_dafny.Map({(self).f2: ((self).f5)[default__.safeIndex(337, ((self).f5).length(0))]})).set(_dafny.SeqWithoutIsStrInference([d_2640_v0_ for d_2651_i1_ in range(default__.abs(947))]), True), _dafny.SeqWithoutIsStrInference([d_2640_v0_ for d_2652_i2_ in range(default__.abs(691))]), (self).f4, d_2649_v8_)
            d_2650_v9_ = nw392_
            d_2653_v10_: _dafny.Seq
            d_2653_v10_ = _dafny.SeqWithoutIsStrInference([d_2650_v9_])
            d_2654_v11_: D19
            d_2654_v11_ = D19_DC48(d_2653_v10_)
            d_2641_v1_ = (d_2641_v1_).set((d_2642_i0_) <= ((self).f4), (len(_dafny.Set({_dafny.Set({(self).f4, d_2642_i0_}), d_2647_v6_})) if ((self).f5)[default__.safeIndex(337, ((self).f5).length(0))] else len((d_2654_v11_).cf79)))
            d_2655_v12_: _dafny.Map
            d_2655_v12_ = _dafny.Map({(self).f7: (self).f4})
            d_2656_v13_: _dafny.MultiSet
            d_2656_v13_ = _dafny.MultiSet([d_2646_v5_, d_2646_v5_])
            d_2657_v14_: _dafny.Array
            nw393_ = _dafny.Array(None, 11)
            nw393_[int(0)] = (self).f4
            nw393_[int(1)] = (self).f4
            nw393_[int(2)] = p0
            nw393_[int(3)] = d_2642_i0_
            nw393_[int(4)] = ((d_2655_v12_)[(self).f4] if ((self).f4) in (d_2655_v12_) else ((d_2656_v13_).set(_dafny.SeqWithoutIsStrInference([d_2642_i0_ for d_2658_i3_ in range(default__.abs(-111))]), default__.abs(len(d_2647_v6_)))).cardinality)
            nw393_[int(5)] = d_2642_i0_
            nw393_[int(6)] = (0) - ((0) - (p0))
            nw393_[int(7)] = d_2642_i0_
            nw393_[int(8)] = d_2642_i0_
            nw393_[int(9)] = len(d_2641_v1_)
            nw393_[int(10)] = len((self).f2)
            d_2657_v14_ = nw393_
            d_2659_v15_: C0
            nw394_ = C0()
            nw394_.ctor__(d_2657_v14_)
            d_2659_v15_ = nw394_
        index420_ = default__.safeIndex(879, ((self).f5).length(0))
        ((self).f5)[index420_] = self.f3
        d_2660_v16_: _dafny.Seq
        d_2660_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfp"))
        d_2661_v17_: _dafny.Array
        nw395_ = _dafny.Array(int(0), 24)
        d_2661_v17_ = nw395_
        d_2662_v18_: _dafny.MultiSet
        d_2662_v18_ = _dafny.MultiSet([d_2661_v17_, d_2661_v17_, d_2661_v17_])
        index421_ = default__.safeIndex(879, ((self).f5).length(0))
        rhs462_ = default__.fm22(True, (d_2662_v18_).cardinality, globalState)
        rhs463_ = (p0) >= ((self).f4)
        rhs464_ = (((self).f2) + (_dafny.SeqWithoutIsStrInference([d_2640_v0_ for d_2663_i4_ in range(default__.abs(-451))]))) + ((self).f2)
        lhs291_ = (self).f5
        lhs292_ = default__.safeIndex(879, ((self).f5).length(0))
        d_2640_v0_ = rhs462_
        lhs291_[lhs292_] = rhs463_
        d_2660_v16_ = rhs464_
        d_2664_v19_: _dafny.Array
        nw396_ = _dafny.Array(D3.default()(), 12)
        d_2664_v19_ = nw396_
        _ingredientsd_0 = []
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_2664_v19_).length(0)):
            d_2665_i5_: int = guard_loop_8_
            if (True) and (((0) <= (d_2665_i5_)) and ((d_2665_i5_) < ((d_2664_v19_).length(0)))):
                _ingredientsd_0.append((d_2664_v19_, int(d_2665_i5_), D3_DC9(p0, (D3_DC8(_dafny.Set({False}))) in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Set({not(((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]), ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], True})), D3_DC8(_dafny.Set({((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]})), D3_DC8(_dafny.Set({self.f3}))])))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_2666_v20_: _dafny.Seq
        d_2666_v20_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        hi10_ = (0) - ((d_2666_v20_)[default__.safeIndex(len(default__.fm7((D1_DC4(self.f3, self.f3)).cf6, self.f3, globalState)), len(d_2666_v20_))])
        for d_2667_i6_ in range(default__.safeModuloInt((self).f4, (self).f4), hi10_):
            r0 = d_2667_i6_
            d_2668_v21_: _dafny.Map
            d_2668_v21_ = _dafny.Map({self.f3: False})
            d_2668_v21_ = (d_2668_v21_).set(((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], self.f3)
            d_2669_v22_: _dafny.Array
            nw397_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_2669_v22_ = nw397_
            rhs465_ = d_2669_v22_
            rhs466_ = p0
            d_2669_v22_ = rhs465_
            r0 = rhs466_
            d_2670_v23_: D10
            d_2670_v23_ = D10_DC29(((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], self.f3)
            d_2671_v24_: _dafny.Map
            d_2671_v24_ = _dafny.Map({d_2670_v23_: ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]})
            d_2672_v25_: _dafny.Array
            nw398_ = _dafny.Array(False, 11)
            d_2672_v25_ = nw398_
            d_2673_v26_: _dafny.Map
            d_2673_v26_ = _dafny.Map({((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]: d_2672_v25_})
            d_2671_v24_ = (d_2671_v24_).set(d_2670_v23_, (d_2673_v26_) == (d_2673_v26_))
        if (642) >= ((self).f4):
            d_2674_v27_: _dafny.Map
            d_2674_v27_ = _dafny.Map({len(d_2660_v16_): 535})
            d_2674_v27_ = (d_2674_v27_ if ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))] else d_2674_v27_)
            r0 = ((self).fm12(p0, (self).f4, (self).f4, ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], globalState)) + (((self).f7) * ((self).f7))
            (self).f3 = self.f3
            if not(((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]):
                d_2675_v28_: C1
                nw399_ = C1()
                nw399_.ctor__()
                d_2675_v28_ = nw399_
                d_2676_v29_: _dafny.Seq
                d_2676_v29_ = _dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], self.f3])
                d_2677_v30_: _dafny.Seq
                d_2677_v30_ = _dafny.SeqWithoutIsStrInference([d_2676_v29_, d_2676_v29_])
                d_2678_v31_: _dafny.Array
                nw400_ = _dafny.Array(None, 9)
                nw400_[int(0)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(1)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(2)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(3)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(4)] = self.f3
                nw400_[int(5)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(6)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(7)] = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                nw400_[int(8)] = True
                d_2678_v31_ = nw400_
                d_2679_v32_: C8
                nw401_ = C8()
                nw401_.ctor__(d_2677_v30_, self.f3, (self).f1, (self).f2, (d_2675_v28_).fm28(((self).f5)[default__.safeIndex(879, ((self).f5).length(0))], (self).f4, globalState), d_2678_v31_)
                d_2679_v32_ = nw401_
                d_2680_v33_: _dafny.Seq
                d_2680_v33_ = _dafny.SeqWithoutIsStrInference([d_2679_v32_])
                d_2681_v34_: _dafny.Seq
                d_2681_v34_ = _dafny.SeqWithoutIsStrInference([(d_2680_v33_)[default__.safeIndex((self).f7, len(d_2680_v33_))]])
                d_2682_v35_: D19
                d_2682_v35_ = D19_DC48(d_2681_v34_)
                d_2683_v36_: _dafny.Seq
                d_2683_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2679_v32_]), d_2680_v33_, d_2680_v33_, d_2680_v33_, d_2681_v34_])
                d_2684_v37_: _dafny.Array
                nw402_ = _dafny.Array(None, 21)
                nw402_[int(0)] = d_2682_v35_
                nw402_[int(1)] = d_2682_v35_
                nw402_[int(2)] = d_2682_v35_
                nw402_[int(3)] = d_2682_v35_
                nw402_[int(4)] = d_2682_v35_
                nw402_[int(5)] = d_2682_v35_
                nw402_[int(6)] = D19_DC48(d_2680_v33_)
                nw402_[int(7)] = D19_DC48(d_2681_v34_)
                nw402_[int(8)] = D19_DC48(d_2680_v33_)
                nw402_[int(9)] = D19_DC48(d_2681_v34_)
                nw402_[int(10)] = d_2682_v35_
                nw402_[int(11)] = d_2682_v35_
                nw402_[int(12)] = d_2682_v35_
                nw402_[int(13)] = D19_DC48((d_2683_v36_)[default__.safeIndex((self).f7, len(d_2683_v36_))])
                nw402_[int(14)] = d_2682_v35_
                nw402_[int(15)] = d_2682_v35_
                nw402_[int(16)] = d_2682_v35_
                nw402_[int(17)] = d_2682_v35_
                nw402_[int(18)] = d_2682_v35_
                nw402_[int(19)] = d_2682_v35_
                nw402_[int(20)] = D19_DC48(d_2680_v33_)
                d_2684_v37_ = nw402_
                d_2685_v38_: _dafny.Seq
                d_2685_v38_ = _dafny.SeqWithoutIsStrInference([d_2684_v37_])
                d_2686_v39_: _dafny.Array
                nw403_ = _dafny.Array(None, 22)
                nw403_[int(0)] = d_2684_v37_
                nw403_[int(1)] = (d_2685_v38_)[default__.safeIndex((self).f7, len(d_2685_v38_))]
                nw403_[int(2)] = d_2684_v37_
                nw403_[int(3)] = d_2684_v37_
                nw403_[int(4)] = d_2684_v37_
                nw403_[int(5)] = d_2684_v37_
                nw403_[int(6)] = d_2684_v37_
                nw403_[int(7)] = (d_2685_v38_)[default__.safeIndex(p0, len(d_2685_v38_))]
                nw403_[int(8)] = d_2684_v37_
                nw403_[int(9)] = d_2684_v37_
                nw403_[int(10)] = d_2684_v37_
                nw403_[int(11)] = d_2684_v37_
                nw403_[int(12)] = d_2684_v37_
                nw403_[int(13)] = d_2684_v37_
                nw403_[int(14)] = d_2684_v37_
                nw403_[int(15)] = d_2684_v37_
                nw403_[int(16)] = d_2684_v37_
                nw403_[int(17)] = d_2684_v37_
                nw403_[int(18)] = (d_2685_v38_)[default__.safeIndex((self).f4, len(d_2685_v38_))]
                nw403_[int(19)] = d_2684_v37_
                nw403_[int(20)] = d_2684_v37_
                nw403_[int(21)] = d_2684_v37_
                d_2686_v39_ = nw403_
                index422_ = default__.safeIndex(86, (d_2686_v39_).length(0))
                (d_2686_v39_)[index422_] = d_2684_v37_
                index423_ = default__.safeIndex(86, (d_2686_v39_).length(0))
                (d_2686_v39_)[index423_] = d_2684_v37_
                d_2676_v29_ = (d_2676_v29_)
                (self).f3 = self.f3
                d_2687_v40_: C0
                nw404_ = C0()
                nw404_.ctor__(d_2661_v17_)
                d_2687_v40_ = nw404_
            elif True:
                d_2688_v41_: _dafny.Seq
                d_2688_v41_ = _dafny.SeqWithoutIsStrInference([self.f3])
                d_2689_v42_: _dafny.Seq
                d_2690_v43_: bool
                d_2691_v44_: _dafny.Seq
                d_2692_v45_: _dafny.Array
                out70_: _dafny.Seq
                out71_: bool
                out72_: _dafny.Seq
                out73_: _dafny.Array
                out70_, out71_, out72_, out73_ = (self).m3(d_2688_v41_, globalState)
                d_2689_v42_ = out70_
                d_2690_v43_ = out71_
                d_2691_v44_ = out72_
                d_2692_v45_ = out73_
                d_2693_v46_: _dafny.Array
                nw405_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_2693_v46_ = nw405_
                index424_ = default__.safeIndex(332, (d_2693_v46_).length(0))
                (d_2693_v46_)[index424_] = d_2640_v0_
                d_2694_v47_: _dafny.Map
                d_2694_v47_ = _dafny.Map({d_2690_v43_: d_2640_v0_})
                index425_ = default__.safeIndex(332, (d_2693_v46_).length(0))
                (d_2693_v46_)[index425_] = ((d_2694_v47_)[self.f3] if (self.f3) in (d_2694_v47_) else _dafny.CodePoint('v'))
                (self).f3 = (self.f3) == (d_2690_v43_)
                rhs467_ = _dafny.SeqWithoutIsStrInference([(self).f7, (self).f7, p0])
                rhs468_ = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
                d_2689_v42_ = rhs467_
                d_2690_v43_ = rhs468_
                index426_ = default__.safeIndex(693, (d_2692_v45_).length(0))
                (d_2692_v45_)[index426_] = (self).f7
                index427_ = default__.safeIndex(693, (d_2692_v45_).length(0))
                (d_2692_v45_)[index427_] = p0
            index428_ = default__.safeIndex(278, (d_2661_v17_).length(0))
            (d_2661_v17_)[index428_] = (self).f7
            index429_ = default__.safeIndex(278, (d_2661_v17_).length(0))
            (d_2661_v17_)[index429_] = p0
        elif True:
            (self).f3 = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
            r0 = (((self).f7) * (default__.fm0(globalState))) + ((self).f4)
            r0 = 226
            r1 = (self).fm11(self.f3, globalState)
            index430_ = default__.safeIndex(879, ((self).f5).length(0))
            ((self).f5)[index430_] = self.f3
        r0 = p0
        r1 = ((self).f5)[default__.safeIndex(879, ((self).f5).length(0))]
        return r0, r1

    @property
    def f7(self):
        return self._f7

class C11(T2, T1, T0):
    def  __init__(self):
        self._f3: bool = False
        self._f1: _dafny.Map = _dafny.Map({})
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: _dafny.Array = _dafny.Array(None, 0)
        self._f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f3(self):
        return self._f3
    @f3.setter
    def f3(self, value):
        self._f3 = value
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f6, f4, f5, f3, f1, f2):
        (self)._f6 = f6
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f3 = f3
        (self)._f1 = f1
        (self)._f2 = f2

    def fm13(self, globalState):
        return True

    def fm14(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2695_i0_ in range(default__.abs(732))])) + ((self).f2)

    def fm12(self, p0, p1, p2, p3, globalState):
        return len(_dafny.SeqWithoutIsStrInference([((self).f4) + ((0) - ((self).f4)), 623, (self).f4]))

    def fm11(self, p0, globalState):
        return (self).f6

    def m1(self, p0, p1, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: int = int(0)
        r2: int = int(0)
        d_2696_v0_: _dafny.Seq
        d_2696_v0_ = _dafny.SeqWithoutIsStrInference([self.f3])
        d_2697_v1_: _dafny.Map
        d_2697_v1_ = _dafny.Map({True: p0})
        r1 = (default__.safeModuloInt(len(d_2696_v0_), len(default__.fm7((self).f6, not(self.f3), globalState)))) + (len(d_2697_v1_))
        d_2698_v2_: _dafny.Seq
        d_2698_v2_ = _dafny.SeqWithoutIsStrInference([d_2696_v0_])
        d_2699_v4_: _dafny.Set
        d_2699_v4_ = _dafny.Set({(self).f2})
        d_2700_v5_: D20
        def iife237_():
            coll85_ = _dafny.Map()
            compr_85_: _dafny.Seq
            for compr_85_ in (d_2699_v4_).Elements:
                d_2701_v3_: _dafny.Seq = compr_85_
                if (d_2701_v3_) in (d_2699_v4_):
                    coll85_[d_2701_v3_] = self.f3
            return _dafny.Map(coll85_)
        d_2700_v5_ = D20_DC51(iife237_()
)
        d_2702_v6_: C8
        nw406_ = C8()
        nw406_.ctor__(d_2698_v2_, self.f3, ((self).f1 if (self).f6 else ((d_2700_v5_).cf83).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uokvrp")), (self).f6)), (self).f2, ((self).f4) * (p0), (self).f5)
        d_2702_v6_ = nw406_
        hi11_ = (self).f4
        for d_2703_i0_ in range(p1, hi11_):
            d_2704_v7_: str
            d_2704_v7_ = _dafny.CodePoint('b')
            d_2705_v8_: _dafny.Map
            d_2705_v8_ = _dafny.Map({(self).f6: self.f3})
            d_2706_v9_: D8
            d_2706_v9_ = D8_DC24(d_2704_v7_, d_2705_v8_)
            d_2707_v10_: _dafny.Map
            d_2707_v10_ = _dafny.Map({d_2706_v9_: d_2696_v0_})
            d_2708_v11_: _dafny.Seq
            d_2708_v11_ = _dafny.SeqWithoutIsStrInference([d_2703_i0_, (self).f4, len(_dafny.SeqWithoutIsStrInference([d_2704_v7_ for d_2709_i1_ in range(default__.abs(896))]))])
            if (((d_2707_v10_)[d_2706_v9_] if (d_2706_v9_) in (d_2707_v10_) else default__.fm7((self).f6, (d_2702_v6_).fm11(self.f3, globalState), globalState))) <= ((d_2696_v0_).set(default__.safeIndex(len(((self).f2).set(default__.safeIndex(len(d_2708_v11_), len((self).f2)), d_2704_v7_)), len(d_2696_v0_)), self.f3)):
                d_2710_v12_: _dafny.Array
                nw407_ = _dafny.Array(None, 1)
                nw407_[int(0)] = p0
                d_2710_v12_ = nw407_
                index431_ = default__.safeIndex(318, (d_2710_v12_).length(0))
                (d_2710_v12_)[index431_] = p1
                d_2711_v13_: _dafny.MultiSet
                d_2711_v13_ = _dafny.MultiSet([self.f3])
                index432_ = default__.safeIndex(318, (d_2710_v12_).length(0))
                (d_2710_v12_)[index432_] = default__.safeModuloInt((0) - (len(d_2705_v8_)), (((d_2711_v13_)[self.f3] if (self.f3) in (d_2711_v13_) else default__.fm0(globalState)) if (self).f6 else d_2703_i0_))
                d_2712_v14_: D14
                d_2712_v14_ = D14_DC39(len((self).f2), d_2703_i0_, len((self).f2))
                d_2713_v15_: _dafny.Map
                d_2713_v15_ = _dafny.Map({len((default__.fm31(p1, (self).f6, -281, d_2704_v7_, globalState)) + (_dafny.SeqWithoutIsStrInference([(d_2712_v14_).cf67]))): d_2710_v12_})
                d_2714_v16_: _dafny.Array
                nw408_ = _dafny.Array(_dafny.Map({}), 3)
                d_2714_v16_ = nw408_
                d_2715_v17_: _dafny.Map
                d_2715_v17_ = _dafny.Map({p1: default__.fm0(globalState)})
                index433_ = default__.safeIndex(235, (d_2714_v16_).length(0))
                (d_2714_v16_)[index433_] = d_2715_v17_
                d_2716_v18_: _dafny.Array
                nw409_ = _dafny.Array(_dafny.Seq({}), 1)
                d_2716_v18_ = nw409_
                index434_ = default__.safeIndex(235, (d_2714_v16_).length(0))
                def iife238_():
                    coll86_ = _dafny.Map()
                    compr_86_: int
                    for compr_86_ in (_dafny.Map({(self).f4: (self).f4})).keys.Elements:
                        d_2717_v19_: int = compr_86_
                        if (d_2717_v19_) in (_dafny.Map({(self).f4: (self).f4})):
                            coll86_[(d_2717_v19_) * ((self).f4)] = d_2703_i0_
                    return _dafny.Map(coll86_)
                rhs469_ = d_2713_v15_
                rhs470_ = iife238_()

                rhs471_ = d_2712_v14_
                rhs472_ = d_2716_v18_
                rhs473_ = p0
                lhs293_ = d_2714_v16_
                lhs294_ = default__.safeIndex(235, (d_2714_v16_).length(0))
                d_2713_v15_ = rhs469_
                lhs293_[lhs294_] = rhs470_
                d_2712_v14_ = rhs471_
                d_2716_v18_ = rhs472_
                r1 = rhs473_
                r1 = d_2703_i0_
                d_2718_v20_: _dafny.Map
                d_2718_v20_ = _dafny.Map({d_2704_v7_: p0})
                d_2718_v20_ = _dafny.Map({d_2704_v7_: (self).f4})
                index435_ = default__.safeIndex(318, (d_2710_v12_).length(0))
                (d_2710_v12_)[index435_] = (d_2710_v12_)[default__.safeIndex(318, (d_2710_v12_).length(0))]
            elif True:
                d_2719_v21_: _dafny.Array
                def lambda103_(d_2720_p1_):
                    def lambda104_(d_2721_i2_):
                        return (d_2721_i2_) + (d_2720_p1_)

                    return lambda104_

                init58_ = lambda103_(p1)
                nw410_ = _dafny.Array(None, 11)
                for i0_58_ in range(nw410_.length(0)):
                    nw410_[i0_58_] = init58_(i0_58_)
                d_2719_v21_ = nw410_
                d_2722_v22_: C0
                nw411_ = C0()
                nw411_.ctor__(d_2719_v21_)
                d_2722_v22_ = nw411_
                d_2723_v23_: _dafny.Seq
                d_2723_v23_ = _dafny.SeqWithoutIsStrInference([d_2722_v22_, d_2722_v22_, d_2722_v22_, d_2722_v22_])
                (self).f3 = ((d_2723_v23_) <= (d_2723_v23_)) and ((self).f6)
                d_2724_v24_: _dafny.MultiSet
                d_2724_v24_ = _dafny.MultiSet([p1])
                d_2725_v25_: C6
                nw412_ = C6()
                nw412_.ctor__((p0) + ((d_2724_v24_).cardinality), p0, d_2703_i0_, (self).f5, (True) == ((self).f6), ((self).f1) | ((self).f1), ((self).f2 if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laxa"))))
                d_2725_v25_ = nw412_
                r2 = len((self).f2)
                d_2708_v11_ = d_2708_v11_
                d_2726_v26_: C2
                nw413_ = C2()
                nw413_.ctor__(p0, self.f3, (self).f1, (self).f2)
                d_2726_v26_ = nw413_
                d_2727_v27_: _dafny.Seq
                d_2727_v27_ = _dafny.SeqWithoutIsStrInference([d_2726_v26_])
                (self).f3 = default__.fm3(len(d_2727_v27_), globalState)
            d_2728_v28_: _dafny.Array
            def lambda105_(d_2729_i3_):
                return default__.safeDivisionInt(d_2729_i3_, 34)

            init59_ = lambda105_
            nw414_ = _dafny.Array(None, 2)
            for i0_59_ in range(nw414_.length(0)):
                nw414_[i0_59_] = init59_(i0_59_)
            d_2728_v28_ = nw414_
            d_2730_v29_: C0
            nw415_ = C0()
            nw415_.ctor__(d_2728_v28_)
            d_2730_v29_ = nw415_
            d_2731_v30_: _dafny.MultiSet
            d_2731_v30_ = _dafny.MultiSet([d_2708_v11_, d_2708_v11_, d_2708_v11_, d_2708_v11_, _dafny.SeqWithoutIsStrInference([p1])])
            rhs474_ = d_2697_v1_
            rhs475_ = ((d_2731_v30_).intersection(d_2731_v30_)).isdisjoint(default__.fm61(not(self.f3), globalState))
            rhs476_ = d_2730_v29_
            rhs477_ = (d_2730_v29_).f14
            lhs295_ = self
            d_2697_v1_ = rhs474_
            lhs295_.f3 = rhs475_
            d_2730_v29_ = rhs476_
            d_2728_v28_ = rhs477_
            d_2732_v31_: C1
            nw416_ = C1()
            nw416_.ctor__()
            d_2732_v31_ = nw416_
            d_2733_v32_: _dafny.Map
            d_2733_v32_ = _dafny.Map({d_2732_v31_: p0})
            d_2734_v33_: _dafny.Map
            d_2734_v33_ = _dafny.Map({(d_2730_v29_).f14: d_2733_v32_})
            d_2735_v34_: _dafny.MultiSet
            d_2735_v34_ = _dafny.MultiSet([(self).f6])
            d_2736_v35_: _dafny.Set
            d_2736_v35_ = _dafny.Set({(self).f6, (self).f6})
            rhs478_ = (_dafny.Set({p0, (self).f4, p1})).isdisjoint(_dafny.Set({((d_2735_v34_)[self.f3] if (self.f3) in (d_2735_v34_) else len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lofhhys"))).set(default__.safeIndex(len(d_2736_v35_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lofhhys")))), _dafny.CodePoint('x'))))}))
            rhs479_ = (d_2734_v33_) | ((d_2734_v33_) | (_dafny.Map({(d_2730_v29_).f14: d_2733_v32_})))
            rhs480_ = len((d_2708_v11_) + (d_2708_v11_))
            rhs481_ = self.f3
            rhs482_ = (self).fm12(default__.safeModuloInt(p0, -521), len(d_2736_v35_), default__.safeModuloInt(p1, (0) - (d_2703_i0_)), (_dafny.CodePoint('l')) not in ((self).f2), globalState)
            lhs296_ = self
            lhs297_ = self
            lhs296_.f3 = rhs478_
            d_2734_v33_ = rhs479_
            r1 = rhs480_
            lhs297_.f3 = rhs481_
            r2 = rhs482_
            d_2737_v36_: _dafny.Set
            d_2737_v36_ = _dafny.Set({d_2703_i0_, d_2703_i0_})
            d_2738_v37_: _dafny.Array
            nw417_ = _dafny.Array(False, 6)
            d_2738_v37_ = nw417_
            index436_ = default__.safeIndex(720, (d_2728_v28_).length(0))
            (d_2728_v28_)[index436_] = (self).f4
            d_2739_v38_: _dafny.Seq
            d_2739_v38_ = _dafny.SeqWithoutIsStrInference([d_2737_v36_])
            d_2740_v39_: _dafny.Seq
            d_2740_v39_ = _dafny.SeqWithoutIsStrInference([(self).f5])
            index437_ = default__.safeIndex(720, (d_2728_v28_).length(0))
            rhs483_ = (d_2739_v38_)[default__.safeIndex(p1, len(d_2739_v38_))]
            rhs484_ = (d_2740_v39_)[default__.safeIndex(d_2703_i0_, len(d_2740_v39_))]
            rhs485_ = ((self).f4) - (p1)
            rhs486_ = 328
            lhs298_ = d_2728_v28_
            lhs299_ = default__.safeIndex(720, (d_2728_v28_).length(0))
            d_2737_v36_ = rhs483_
            d_2738_v37_ = rhs484_
            lhs298_[lhs299_] = rhs485_
            r1 = rhs486_
        if self.f3:
            (self).f3 = (self).f6
            d_2741_v40_: T2
            nw418_ = C3()
            nw418_.ctor__(333, (self).f5, (d_2696_v0_) < (d_2696_v0_), (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyvxifc")))
            d_2741_v40_ = nw418_
            d_2741_v40_ = d_2741_v40_
            d_2742_v41_: _dafny.Array
            def lambda106_(d_2743_v40_):
                def lambda107_(d_2744_i4_):
                    return (_dafny.Set({d_2743_v40_.f3}) if d_2743_v40_.f3 else _dafny.Set({d_2743_v40_.f3, True}))

                return lambda107_

            init60_ = lambda106_(d_2741_v40_)
            nw419_ = _dafny.Array(None, 2)
            for i0_60_ in range(nw419_.length(0)):
                nw419_[i0_60_] = init60_(i0_60_)
            d_2742_v41_ = nw419_
            index438_ = default__.safeIndex(87, (d_2742_v41_).length(0))
            (d_2742_v41_)[index438_] = _dafny.Set({(self).f6, (self).f6, (d_2696_v0_)[default__.safeIndex(p1, len(d_2696_v0_))]})
            d_2745_v42_: str
            d_2745_v42_ = _dafny.CodePoint('l')
            d_2746_v43_: _dafny.Set
            d_2746_v43_ = _dafny.Set({(self).f6, self.f3})
            index439_ = default__.safeIndex(87, (d_2742_v41_).length(0))
            (d_2742_v41_)[index439_] = ((default__.fm6(d_2741_v40_.f3, d_2745_v42_, globalState)) - (_dafny.Set({self.f3}))) - (d_2746_v43_)
            index440_ = default__.safeIndex(585, ((self).f5).length(0))
            ((self).f5)[index440_] = self.f3
            index441_ = default__.safeIndex(585, ((self).f5).length(0))
            ((self).f5)[index441_] = (self).f6
            d_2747_v44_: _dafny.Seq
            d_2747_v44_ = _dafny.SeqWithoutIsStrInference([(d_2741_v40_).f4])
            d_2748_v45_: T1
            nw420_ = C2()
            nw420_.ctor__(p0, True, default__.fm62(len(d_2747_v44_), (self).f6, globalState), (d_2741_v40_).f2)
            d_2748_v45_ = nw420_
            d_2749_v46_: T1
            d_2749_v46_ = d_2748_v45_
            d_2750_v47_: _dafny.Seq
            d_2750_v47_ = _dafny.SeqWithoutIsStrInference([d_2748_v45_, d_2748_v45_, d_2748_v45_])
            d_2751_v48_: _dafny.Map
            d_2751_v48_ = _dafny.Map({(0) - (p0): d_2748_v45_})
            d_2752_v49_: _dafny.Array
            nw421_ = _dafny.Array(None, 18)
            nw421_[int(0)] = (d_2749_v46_)
            nw421_[int(1)] = d_2748_v45_
            nw421_[int(2)] = (d_2748_v45_)
            nw421_[int(3)] = d_2748_v45_
            nw421_[int(4)] = d_2748_v45_
            nw421_[int(5)] = d_2748_v45_
            nw421_[int(6)] = (d_2750_v47_)[default__.safeIndex(p0, len(d_2750_v47_))]
            nw421_[int(7)] = (d_2748_v45_)
            nw421_[int(8)] = d_2748_v45_
            nw421_[int(9)] = d_2748_v45_
            nw421_[int(10)] = d_2748_v45_
            nw421_[int(11)] = d_2748_v45_
            nw421_[int(12)] = d_2748_v45_
            nw421_[int(13)] = d_2748_v45_
            nw421_[int(14)] = d_2748_v45_
            nw421_[int(15)] = ((d_2751_v48_)[p0] if (p0) in (d_2751_v48_) else d_2748_v45_)
            nw421_[int(16)] = d_2748_v45_
            nw421_[int(17)] = d_2748_v45_
            d_2752_v49_ = nw421_
            index442_ = default__.safeIndex(874, (d_2752_v49_).length(0))
            (d_2752_v49_)[index442_] = d_2748_v45_
            index443_ = default__.safeIndex(874, (d_2752_v49_).length(0))
            (d_2752_v49_)[index443_] = d_2748_v45_
        elif True:
            d_2753_v50_: _dafny.Map
            d_2753_v50_ = _dafny.Map({(self).f4: self.f3})
            d_2754_v51_: str
            d_2754_v51_ = _dafny.CodePoint('i')
            d_2755_v52_: T0
            nw422_ = C5()
            nw422_.ctor__((len(d_2753_v50_)) <= (default__.fm0(globalState)), (self).f1, ((self).f2).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f2, (self).f2])), len((self).f2)), d_2754_v51_))
            d_2755_v52_ = nw422_
            d_2755_v52_ = d_2755_v52_
            d_2756_v53_: _dafny.Map
            d_2756_v53_ = _dafny.Map({d_2754_v51_: p0})
            d_2757_v54_: _dafny.MultiSet
            d_2757_v54_ = _dafny.MultiSet([self.f3, (self).f6, True, self.f3, (self).f6])
            d_2758_v55_: _dafny.Map
            d_2758_v55_ = _dafny.Map({True: (self).f6})
            d_2759_v56_: _dafny.MultiSet
            d_2759_v56_ = _dafny.MultiSet([(d_2755_v52_).f2, (d_2755_v52_).f2, (d_2755_v52_).f2])
            d_2760_v57_: _dafny.Seq
            d_2760_v57_ = _dafny.SeqWithoutIsStrInference([(self).f4, p1, -492, (self).f4, p0])
            d_2761_v58_: D8
            d_2761_v58_ = D8_DC24(d_2754_v51_, d_2758_v55_)
            d_2762_v59_: _dafny.Map
            d_2762_v59_ = _dafny.Map({(self).f6: (d_2761_v58_).cf41})
            d_2763_v60_: C4
            nw423_ = C4()
            nw423_.ctor__(self.f3, (d_2755_v52_).f1, (d_2755_v52_).f2)
            d_2763_v60_ = nw423_
            d_2764_v61_: _dafny.MultiSet
            d_2764_v61_ = _dafny.MultiSet([p0, (self).f4])
            d_2765_v62_: _dafny.Array
            nw424_ = _dafny.Array(None, 24)
            nw424_[int(0)] = (p1) * (p0)
            nw424_[int(1)] = (p1) * (len(d_2756_v53_))
            nw424_[int(2)] = (d_2757_v54_).cardinality
            nw424_[int(3)] = p0
            nw424_[int(4)] = len((d_2758_v55_) | (d_2758_v55_))
            nw424_[int(5)] = (242) - (p1)
            nw424_[int(6)] = default__.safeDivisionInt((d_2759_v56_).cardinality, (self).f4)
            nw424_[int(7)] = (d_2760_v57_)[default__.safeIndex(800, len(d_2760_v57_))]
            nw424_[int(8)] = (self).f4
            nw424_[int(9)] = (self).f4
            nw424_[int(10)] = -279
            nw424_[int(11)] = p1
            nw424_[int(12)] = 507
            nw424_[int(13)] = p0
            nw424_[int(14)] = len(d_2762_v59_)
            nw424_[int(15)] = (0) - (default__.safeDivisionInt(len(d_2697_v1_), p0))
            nw424_[int(16)] = (p0 if self.f3 else (self).f4)
            nw424_[int(17)] = p1
            nw424_[int(18)] = (172 if self.f3 else (_dafny.MultiSet([d_2763_v60_])).cardinality)
            nw424_[int(19)] = p0
            nw424_[int(20)] = ((d_2764_v61_).cardinality) + ((self).f4)
            nw424_[int(21)] = p0
            nw424_[int(22)] = p0
            nw424_[int(23)] = p1
            d_2765_v62_ = nw424_
            index444_ = default__.safeIndex(633, (d_2765_v62_).length(0))
            (d_2765_v62_)[index444_] = p1
            index445_ = default__.safeIndex(633, (d_2765_v62_).length(0))
            (d_2765_v62_)[index445_] = p0
            r1 = p0
            d_2766_v63_: _dafny.Map
            d_2766_v63_ = _dafny.Map({(self).f4: ((d_2757_v54_).intersection(d_2757_v54_)).cardinality})
            d_2766_v63_ = (d_2766_v63_) | (d_2766_v63_)
            if self.f3:
                (self).f3 = (d_2702_v6_).fm11((self).f6, globalState)
                d_2765_v62_ = d_2765_v62_
                index446_ = default__.safeIndex(620, ((self).f5).length(0))
                ((self).f5)[index446_] = (self).f6
                index447_ = default__.safeIndex(620, ((self).f5).length(0))
                ((self).f5)[index447_] = self.f3
                d_2767_v64_: _dafny.Array
                nw425_ = _dafny.Array(False, 16)
                d_2767_v64_ = nw425_
                d_2768_v65_: C7
                nw426_ = C7()
                nw426_.ctor__((self).f4, d_2767_v64_, ((self).f5)[default__.safeIndex(620, ((self).f5).length(0))], (d_2755_v52_).f1, (d_2755_v52_).f2)
                d_2768_v65_ = nw426_
                d_2697_v1_ = (d_2697_v1_).set(((d_2760_v57_)[default__.safeIndex(p1, len(d_2760_v57_))]) >= ((self).f4), (self).f4)
            elif True:
                (self).f3 = ((0) - ((self).f4)) <= ((self).f4)
                r1 = 452
                r2 = p1
                (self).f3 = ((d_2753_v50_)[p1] if (p1) in (d_2753_v50_) else not (self.f3) or (True))
                d_2769_v66_: _dafny.Seq
                d_2769_v66_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f2: self.f3})])
                d_2770_v67_: _dafny.Map
                d_2770_v67_ = _dafny.Map({self.f3: (self).f1})
                d_2771_v68_: T1
                nw427_ = C8()
                nw427_.ctor__((d_2702_v6_).f10, self.f3, ((d_2769_v66_)[default__.safeIndex((self).f4, len(d_2769_v66_))]) | ((((d_2770_v67_)[self.f3] if (self.f3) in (d_2770_v67_) else (self).f1)).set((d_2755_v52_).f2, self.f3)), (self).f2, len((_dafny.SeqWithoutIsStrInference([d_2754_v51_ for d_2772_i5_ in range(default__.abs(655))])) + ((d_2755_v52_).f2)), (self).f5)
                d_2771_v68_ = nw427_
                d_2773_v69_: T1
                d_2773_v69_ = d_2771_v68_
                d_2771_v68_ = (d_2773_v69_)
        if self.f3:
            d_2774_v70_: D11
            d_2774_v70_ = D11_DC31((self).f4)
            rhs487_ = (default__.fm49(d_2774_v70_, len((self).f2), p0, globalState)).cf28
            rhs488_ = ((self).f4) - (((d_2702_v6_).fm12((0) - (p1), p1, 124, True, globalState)) - (len((self).f2)))
            rhs489_ = default__.fm0(globalState)
            lhs300_ = self
            lhs300_.f3 = rhs487_
            r2 = rhs488_
            r2 = rhs489_
            d_2775_v71_: _dafny.MultiSet
            d_2775_v71_ = _dafny.MultiSet([self.f3])
            d_2776_v72_: _dafny.Map
            d_2776_v72_ = _dafny.Map({(self).f2: (0) - (((d_2775_v71_).set((self).f6, default__.abs(p0))).cardinality)})
            if ((_dafny.MultiSet([self.f3, True, (self).f6])) - (d_2775_v71_)).ispropersubset((_dafny.MultiSet(d_2696_v0_)) | (default__.fm5(d_2776_v72_, globalState))):
                d_2777_v73_: str
                d_2777_v73_ = _dafny.CodePoint('j')
                rhs490_ = d_2777_v73_
                rhs491_ = ((p1) * (p1)) > (((d_2775_v71_)[(d_2696_v0_)[default__.safeIndex(len((self).f2), len(d_2696_v0_))]] if ((d_2696_v0_)[default__.safeIndex(len((self).f2), len(d_2696_v0_))]) in (d_2775_v71_) else p1))
                lhs301_ = self
                d_2777_v73_ = rhs490_
                lhs301_.f3 = rhs491_
                (self).f3 = self.f3
                d_2778_v74_: _dafny.Array
                def lambda108_(d_2779_i6_):
                    return default__.safeModuloInt(d_2779_i6_, (self).f4)

                init61_ = lambda108_
                nw428_ = _dafny.Array(None, 15)
                for i0_61_ in range(nw428_.length(0)):
                    nw428_[i0_61_] = init61_(i0_61_)
                d_2778_v74_ = nw428_
                d_2780_v75_: D8
                d_2780_v75_ = D8_DC22(d_2775_v71_)
                d_2781_v76_: _dafny.Seq
                d_2781_v76_ = _dafny.SeqWithoutIsStrInference([d_2780_v75_])
                d_2782_v77_: _dafny.Set
                d_2782_v77_ = _dafny.Set({d_2781_v76_})
                d_2783_v78_: _dafny.Seq
                d_2783_v78_ = _dafny.SeqWithoutIsStrInference([len(d_2782_v77_), p1])
                d_2784_v79_: _dafny.MultiSet
                d_2784_v79_ = _dafny.MultiSet([d_2778_v74_])
                nw429_ = _dafny.Array(None, 17)
                nw429_[int(0)] = p0
                nw429_[int(1)] = default__.safeModuloInt(len(d_2783_v78_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gc"))))
                nw429_[int(2)] = p1
                nw429_[int(3)] = ((d_2784_v79_)[d_2778_v74_] if (d_2778_v74_) in (d_2784_v79_) else (self).f4)
                nw429_[int(4)] = default__.safeDivisionInt((self).f4, len((self).f2))
                nw429_[int(5)] = p1
                nw429_[int(6)] = p1
                nw429_[int(7)] = (0) - ((self).f4)
                nw429_[int(8)] = 28
                nw429_[int(9)] = p1
                nw429_[int(10)] = p1
                nw429_[int(11)] = (self).f4
                nw429_[int(12)] = len(((self).f2) + ((self).f2))
                nw429_[int(13)] = (self).f4
                nw429_[int(14)] = (self).f4
                nw429_[int(15)] = p0
                nw429_[int(16)] = len(default__.fm20(True, (self).f6, (self).fm11(True, globalState), d_2696_v0_, globalState))
                d_2778_v74_ = nw429_
                d_2785_v80_: C2
                nw430_ = C2()
                nw430_.ctor__((self).f4, (self).f6, (self).f1, (self).f2)
                d_2785_v80_ = nw430_
                d_2786_v82_: T1
                nw431_ = C10()
                def iife239_():
                    coll87_ = _dafny.Map()
                    compr_87_: _dafny.Seq
                    for compr_87_ in ((self).f1).keys.Elements:
                        d_2787_v81_: _dafny.Seq = compr_87_
                        if (d_2787_v81_) in ((self).f1):
                            coll87_[d_2787_v81_] = (d_2696_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtww"))), len(d_2696_v0_))]
                    return _dafny.Map(coll87_)
                nw431_.ctor__(p0, (self).f4, (self).f5, (self).f6, iife239_()
                , (self).f2)
                d_2786_v82_ = nw431_
                d_2788_v83_: T1
                d_2788_v83_ = d_2786_v82_
                d_2788_v83_ = d_2788_v83_
            elif True:
                d_2789_v84_: C6
                nw432_ = C6()
                nw432_.ctor__((self).f4, 781, len(_dafny.SeqWithoutIsStrInference([self.f3, self.f3])), (self).f5, not((True if (self).f6 else not(False))), (self).f1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "es")))
                d_2789_v84_ = nw432_
                d_2790_v85_: _dafny.Map
                d_2790_v85_ = _dafny.Map({(self).f6: (self).f6})
                d_2791_v86_: _dafny.Set
                d_2791_v86_ = _dafny.Set({_dafny.Map({(d_2696_v0_)[default__.safeIndex((self).f4, len(d_2696_v0_))]: self.f3}), d_2790_v85_})
                d_2792_v87_: _dafny.MultiSet
                d_2792_v87_ = _dafny.MultiSet([p0])
                d_2793_v88_: _dafny.Map
                d_2793_v88_ = _dafny.Map({(d_2792_v87_).cardinality: (self).f6})
                def iife240_():
                    coll88_ = _dafny.Set()
                    compr_88_: _dafny.Map
                    for compr_88_ in (default__.fm63(((d_2793_v88_)[(d_2789_v84_).f11] if ((d_2789_v84_).f11) in (d_2793_v88_) else self.f3), self.f3, self.f3, (self).f6, globalState)).Elements:
                        d_2794_v89_: _dafny.Map = compr_88_
                        if (d_2794_v89_) in (default__.fm63(((d_2793_v88_)[(d_2789_v84_).f11] if ((d_2789_v84_).f11) in (d_2793_v88_) else self.f3), self.f3, self.f3, (self).f6, globalState)):
                            coll88_ = coll88_.union(_dafny.Set([d_2794_v89_]))
                    return _dafny.Set(coll88_)
                (self).f3 = (iife240_()
                ).issubset(d_2791_v86_)
                d_2795_v90_: _dafny.Array
                def lambda109_(d_2796_v0_):
                    def lambda110_(d_2797_i7_):
                        return default__.safeDivisionInt(d_2797_i7_, len(d_2796_v0_))

                    return lambda110_

                init62_ = lambda109_(d_2696_v0_)
                nw433_ = _dafny.Array(None, 21)
                for i0_62_ in range(nw433_.length(0)):
                    nw433_[i0_62_] = init62_(i0_62_)
                d_2795_v90_ = nw433_
                d_2798_v91_: _dafny.Seq
                d_2798_v91_ = _dafny.SeqWithoutIsStrInference([d_2795_v90_, d_2795_v90_, d_2795_v90_, d_2795_v90_])
                d_2799_v92_: C0
                nw434_ = C0()
                nw434_.ctor__((d_2798_v91_)[default__.safeIndex((self).f4, len(d_2798_v91_))])
                d_2799_v92_ = nw434_
                d_2800_v93_: str
                d_2800_v93_ = _dafny.CodePoint('x')
                d_2801_v94_: _dafny.Map
                d_2801_v94_ = _dafny.Map({d_2800_v93_: self.f3})
                d_2802_v96_: _dafny.Set
                d_2802_v96_ = _dafny.Set({_dafny.CodePoint('t'), d_2800_v93_, d_2800_v93_})
                index448_ = default__.safeIndex(938, ((self).f5).length(0))
                def iife241_():
                    coll89_ = _dafny.Set()
                    compr_89_: str
                    for compr_89_ in (d_2801_v94_).keys.Elements:
                        d_2803_v95_: str = compr_89_
                        if (d_2803_v95_) in (d_2801_v94_):
                            coll89_ = coll89_.union(_dafny.Set([d_2803_v95_]))
                    return _dafny.Set(coll89_)
                ((self).f5)[index448_] = not((d_2802_v96_).issubset(iife241_()
                ))
                index449_ = default__.safeIndex(938, ((self).f5).length(0))
                ((self).f5)[index449_] = (p0) <= ((self).f4)
                d_2775_v71_ = d_2775_v71_
            d_2804_v97_: _dafny.Map
            d_2804_v97_ = _dafny.Map({(0) - (default__.safeModuloInt(-3, (self).f4)): (self).f6})
            d_2805_v98_: D6
            d_2805_v98_ = D6_DC16(d_2804_v97_)
            d_2804_v97_ = (d_2805_v98_).cf27
            (self).f3 = (self).f6
            r1 = (self).fm12(p0, -259, (self).f4, not(((d_2696_v0_)[default__.safeIndex(len((self).f2), len(d_2696_v0_))] if True else self.f3)), globalState)
        elif True:
            if self.f3:
                d_2806_v99_: _dafny.Array
                def lambda111_(d_2807_p0_):
                    def lambda112_(d_2808_i8_):
                        return (d_2808_i8_) * ((0) - (d_2807_p0_))

                    return lambda112_

                init63_ = lambda111_(p0)
                nw435_ = _dafny.Array(None, 11)
                for i0_63_ in range(nw435_.length(0)):
                    nw435_[i0_63_] = init63_(i0_63_)
                d_2806_v99_ = nw435_
                d_2806_v99_ = d_2806_v99_
                d_2809_v100_: _dafny.Map
                d_2809_v100_ = _dafny.Map({(self).f5: not((self).f6)})
                d_2810_v101_: bool
                d_2811_v102_: bool
                d_2812_v103_: _dafny.Seq
                d_2813_v104_: _dafny.Seq
                out74_: bool
                out75_: bool
                out76_: _dafny.Seq
                out77_: _dafny.Seq
                out74_, out75_, out76_, out77_ = default__.m0((_dafny.Map({(self).f5: (self).f6})) | (d_2809_v100_), globalState)
                d_2810_v101_ = out74_
                d_2811_v102_ = out75_
                d_2812_v103_ = out76_
                d_2813_v104_ = out77_
                d_2814_v105_: _dafny.Seq
                d_2814_v105_ = _dafny.SeqWithoutIsStrInference([(p0) + (p0)])
                d_2815_v106_: _dafny.MultiSet
                d_2815_v106_ = _dafny.MultiSet([d_2810_v101_])
                d_2816_v107_: _dafny.MultiSet
                d_2816_v107_ = _dafny.MultiSet([(self).f2])
                d_2817_v108_: _dafny.Map
                d_2817_v108_ = _dafny.Map({(d_2816_v107_).cardinality: d_2810_v101_})
                d_2818_v109_: _dafny.MultiSet
                d_2818_v109_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))), len((d_2814_v105_).set(default__.safeIndex(p1, len(d_2814_v105_)), p1)), 754, ((d_2815_v106_)[(self).f6] if ((self).f6) in (d_2815_v106_) else p0), len(d_2817_v108_)])
                d_2819_v110_: _dafny.Set
                d_2819_v110_ = _dafny.Set({d_2811_v102_})
                d_2820_v111_: _dafny.Set
                d_2820_v111_ = _dafny.Set({d_2819_v110_})
                d_2821_v112_: _dafny.Map
                d_2821_v112_ = _dafny.Map({d_2820_v111_: len(d_2814_v105_)})
                d_2814_v105_ = _dafny.SeqWithoutIsStrInference([((d_2818_v109_)[len(d_2814_v105_)] if (len(d_2814_v105_)) in (d_2818_v109_) else p1), ((d_2821_v112_)[d_2820_v111_] if (d_2820_v111_) in (d_2821_v112_) else len((self).f2))])
                index450_ = default__.safeIndex(169, (d_2806_v99_).length(0))
                (d_2806_v99_)[index450_] = (self).f4
                index451_ = default__.safeIndex(169, (d_2806_v99_).length(0))
                rhs492_ = ((self).f6) or (d_2811_v102_)
                rhs493_ = ((d_2818_v109_)[p0] if (p0) in (d_2818_v109_) else p0)
                rhs494_ = len((self).f2)
                lhs302_ = d_2806_v99_
                lhs303_ = default__.safeIndex(169, (d_2806_v99_).length(0))
                d_2811_v102_ = rhs492_
                r2 = rhs493_
                lhs302_[lhs303_] = rhs494_
                d_2822_v113_: T0
                nw436_ = C5()
                nw436_.ctor__(d_2810_v101_, _dafny.Map({(self).f2: (d_2812_v103_)[default__.safeIndex(len(d_2696_v0_), len(d_2812_v103_))]}), (self).f2)
                d_2822_v113_ = nw436_
                d_2823_v114_: _dafny.Map
                d_2823_v114_ = _dafny.Map({d_2697_v1_: d_2810_v101_})
                d_2824_v115_: str
                d_2824_v115_ = _dafny.CodePoint('p')
                d_2825_v116_: _dafny.Map
                d_2825_v116_ = _dafny.Map({not(d_2810_v101_): default__.fm6(self.f3, d_2824_v115_, globalState)})
                d_2826_v117_: _dafny.Array
                nw437_ = _dafny.Array(None, 20)
                nw437_[int(0)] = _dafny.Set({not(((d_2823_v114_)[d_2697_v1_] if (d_2697_v1_) in (d_2823_v114_) else not((self).f6))), d_2810_v101_, (self).f6})
                nw437_[int(1)] = (d_2819_v110_).intersection(_dafny.Set({False, self.f3}))
                nw437_[int(2)] = d_2819_v110_
                nw437_[int(3)] = (d_2819_v110_).intersection(_dafny.Set({True, d_2811_v102_, True, not((self).fm11(d_2811_v102_, globalState))}))
                nw437_[int(4)] = (((d_2825_v116_)[False] if (False) in (d_2825_v116_) else d_2819_v110_)) - (d_2819_v110_)
                nw437_[int(5)] = default__.fm6(d_2810_v101_, ((self).f2)[default__.safeIndex(894, len((self).f2))], globalState)
                nw437_[int(6)] = (d_2819_v110_) - (d_2819_v110_)
                nw437_[int(7)] = d_2819_v110_
                nw437_[int(8)] = d_2819_v110_
                nw437_[int(9)] = _dafny.Set({d_2810_v101_})
                nw437_[int(10)] = (d_2819_v110_) - (d_2819_v110_)
                nw437_[int(11)] = d_2819_v110_
                nw437_[int(12)] = d_2819_v110_
                nw437_[int(13)] = d_2819_v110_
                nw437_[int(14)] = (d_2819_v110_).intersection(d_2819_v110_)
                nw437_[int(15)] = (d_2819_v110_) | (d_2819_v110_)
                nw437_[int(16)] = (default__.fm6(d_2811_v102_, d_2824_v115_, globalState)).intersection(d_2819_v110_)
                nw437_[int(17)] = d_2819_v110_
                nw437_[int(18)] = d_2819_v110_
                nw437_[int(19)] = d_2819_v110_
                d_2826_v117_ = nw437_
                rhs495_ = d_2822_v113_
                rhs496_ = d_2826_v117_
                d_2822_v113_ = rhs495_
                d_2826_v117_ = rhs496_
            elif True:
                d_2827_v118_: _dafny.Array
                def lambda113_(d_2828_i9_):
                    return _dafny.CodePoint('x')

                init64_ = lambda113_
                nw438_ = _dafny.Array(None, 20)
                for i0_64_ in range(nw438_.length(0)):
                    nw438_[i0_64_] = init64_(i0_64_)
                d_2827_v118_ = nw438_
                d_2829_v119_: _dafny.Map
                d_2829_v119_ = _dafny.Map({self.f3: (self).f6})
                d_2830_v120_: _dafny.Map
                d_2830_v120_ = _dafny.Map({d_2827_v118_: p0})
                index452_ = default__.safeIndex(889, (d_2827_v118_).length(0))
                (d_2827_v118_)[index452_] = default__.fm38(((d_2829_v119_)[(self).f6] if ((self).f6) in (d_2829_v119_) else self.f3), (self).f6, len((d_2830_v120_).set(d_2827_v118_, 603)), globalState)
                index453_ = default__.safeIndex(889, (d_2827_v118_).length(0))
                (d_2827_v118_)[index453_] = ((self).f2)[default__.safeIndex((self).f4, len((self).f2))]
                d_2831_v121_: _dafny.Seq
                d_2831_v121_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqys"))
                d_2832_v122_: _dafny.Array
                nw439_ = _dafny.Array(None, 1)
                nw439_[int(0)] = (self).f5
                d_2832_v122_ = nw439_
                index454_ = default__.safeIndex(591, (d_2832_v122_).length(0))
                (d_2832_v122_)[index454_] = (self).f5
                d_2833_v123_: _dafny.Map
                d_2833_v123_ = _dafny.Map({p1: self.f3})
                d_2834_v124_: _dafny.Set
                d_2834_v124_ = _dafny.Set({((d_2833_v123_)[(self).f4] if ((self).f4) in (d_2833_v123_) else self.f3)})
                index455_ = default__.safeIndex(591, (d_2832_v122_).length(0))
                rhs497_ = (self).f2
                rhs498_ = (self).f5
                rhs499_ = (d_2834_v124_).isdisjoint((d_2834_v124_).intersection(d_2834_v124_))
                lhs304_ = d_2832_v122_
                lhs305_ = default__.safeIndex(591, (d_2832_v122_).length(0))
                lhs306_ = self
                d_2831_v121_ = rhs497_
                lhs304_[lhs305_] = rhs498_
                lhs306_.f3 = rhs499_
                (self).f3 = self.f3
                d_2835_v125_: _dafny.Array
                nw440_ = _dafny.Array(int(0), 9)
                d_2835_v125_ = nw440_
                d_2836_v126_: _dafny.Seq
                d_2836_v126_ = _dafny.SeqWithoutIsStrInference([d_2835_v125_, d_2835_v125_, d_2835_v125_])
                d_2835_v125_ = (d_2836_v126_)[default__.safeIndex(p1, len(d_2836_v126_))]
                d_2837_v127_: T0
                nw441_ = C5()
                nw441_.ctor__((self).f6, ((self).f1).set(d_2831_v121_, self.f3), d_2831_v121_)
                d_2837_v127_ = nw441_
                d_2838_v128_: _dafny.Set
                d_2838_v128_ = _dafny.Set({(self).f4})
                d_2839_v129_: _dafny.Seq
                d_2839_v129_ = _dafny.SeqWithoutIsStrInference([d_2838_v128_, d_2838_v128_])
                d_2840_v130_: _dafny.Set
                d_2840_v130_ = _dafny.Set({p1, len((d_2839_v129_)[default__.safeIndex(p1, len(d_2839_v129_))])})
                d_2841_v131_: _dafny.Map
                d_2841_v131_ = _dafny.Map({d_2837_v127_: d_2838_v128_})
                d_2842_v132_: _dafny.Map
                d_2842_v132_ = _dafny.Map({(0) - ((self).f4): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))})
                d_2843_v133_: D11
                d_2843_v133_ = D11_DC30(d_2842_v132_)
                d_2844_v134_: D13
                d_2844_v134_ = D13_DC36(D13_DC35((self).f4, (self).f4, (self).f6, 547, d_2843_v133_))
                d_2845_v135_: D13
                d_2845_v135_ = D13_DC36(d_2844_v134_)
                d_2846_v136_: _dafny.Map
                d_2846_v136_ = _dafny.Map({((d_2841_v131_)[d_2837_v127_] if (d_2837_v127_) in (d_2841_v131_) else d_2840_v130_): d_2845_v135_})
                d_2846_v136_ = (d_2846_v136_).set(_dafny.Set({len((self).f2)}), d_2845_v135_)
            d_2847_v137_: str
            d_2847_v137_ = _dafny.CodePoint('l')
            d_2848_v138_: C2
            nw442_ = C2()
            nw442_.ctor__((p0) - (default__.fm0(globalState)), not ((self).f6) or (self.f3), ((self).f1).set((self).f2, self.f3), ((self).f2) + ((self).f2))
            d_2848_v138_ = nw442_
            rhs500_ = (0) - (p1)
            rhs501_ = d_2847_v137_
            rhs502_ = d_2848_v138_
            rhs503_ = not ((((self).f1)[(self).f2] if ((self).f2) in ((self).f1) else (self).f6)) or (((d_2848_v138_).f13) > (p0))
            lhs307_ = self
            r1 = rhs500_
            d_2847_v137_ = rhs501_
            d_2848_v138_ = rhs502_
            lhs307_.f3 = rhs503_
            index456_ = default__.safeIndex(946, ((self).f5).length(0))
            ((self).f5)[index456_] = (self).f6
            index457_ = default__.safeIndex(946, ((self).f5).length(0))
            rhs504_ = (self).f6
            rhs505_ = (0) - (p1)
            rhs506_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwfasmko"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gak"))))
            rhs507_ = p0
            lhs308_ = (self).f5
            lhs309_ = default__.safeIndex(946, ((self).f5).length(0))
            lhs308_[lhs309_] = rhs504_
            r1 = rhs505_
            r1 = rhs506_
            r2 = rhs507_
            d_2849_v139_: D1
            d_2849_v139_ = D1_DC3((d_2702_v6_).fm13(globalState), len((default__.fm7(True, ((self).f5)[default__.safeIndex(946, ((self).f5).length(0))], globalState)) + (d_2696_v0_)), (0) - ((p0) + ((d_2848_v138_).f13)))
            source44_ = d_2849_v139_
            if source44_.is_DC3:
                d_2850___mcc_h0_ = source44_.cf3
                d_2851___mcc_h1_ = source44_.cf4
                d_2852___mcc_h2_ = source44_.cf5
                d_2853_cf5_ = d_2852___mcc_h2_
                d_2854_cf4_ = d_2851___mcc_h1_
                d_2855_cf3_ = d_2850___mcc_h0_
                d_2856_v140_: _dafny.Array
                def lambda114_(d_2857_cf3_):
                    def lambda115_(d_2858_i10_):
                        return ((D8_DC22(_dafny.MultiSet([d_2857_cf3_]))).cf37).ispropersubset(_dafny.MultiSet([d_2857_cf3_]))

                    return lambda115_

                init65_ = lambda114_(d_2855_cf3_)
                nw443_ = _dafny.Array(None, 16)
                for i0_65_ in range(nw443_.length(0)):
                    nw443_[i0_65_] = init65_(i0_65_)
                d_2856_v140_ = nw443_
                nw444_ = _dafny.Array(False, 29)
                d_2856_v140_ = nw444_
                index458_ = default__.safeIndex(946, ((self).f5).length(0))
                ((self).f5)[index458_] = self.f3
                d_2859_v141_: _dafny.Array
                nw445_ = _dafny.Array(None, 8)
                nw445_[int(0)] = (d_2848_v138_).f13
                nw445_[int(1)] = len((_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(946, ((self).f5).length(0))]])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([((self).f5)[default__.safeIndex(946, ((self).f5).length(0))]]))), d_2855_cf3_))
                nw445_[int(2)] = d_2854_cf4_
                nw445_[int(3)] = p0
                nw445_[int(4)] = default__.safeModuloInt(p0, 643)
                nw445_[int(5)] = len((self).f2)
                nw445_[int(6)] = len(default__.fm6(True, _dafny.CodePoint('e'), globalState))
                nw445_[int(7)] = p0
                d_2859_v141_ = nw445_
                d_2860_v142_: _dafny.Array
                nw446_ = _dafny.Array(None, 14)
                nw446_[int(0)] = d_2847_v137_
                nw446_[int(1)] = d_2847_v137_
                nw446_[int(2)] = d_2847_v137_
                nw446_[int(3)] = d_2847_v137_
                nw446_[int(4)] = _dafny.CodePoint('o')
                nw446_[int(5)] = default__.fm24(((self).f5)[default__.safeIndex(946, ((self).f5).length(0))], self.f3, len((self).f2), globalState)
                nw446_[int(6)] = d_2847_v137_
                nw446_[int(7)] = d_2847_v137_
                nw446_[int(8)] = d_2847_v137_
                nw446_[int(9)] = d_2847_v137_
                nw446_[int(10)] = d_2847_v137_
                nw446_[int(11)] = d_2847_v137_
                nw446_[int(12)] = d_2847_v137_
                nw446_[int(13)] = d_2847_v137_
                d_2860_v142_ = nw446_
                d_2861_v143_: D10
                d_2861_v143_ = D10_DC29(default__.fm3(-574, globalState), default__.fm3(d_2853_cf5_, globalState))
                d_2862_v144_: _dafny.Map
                d_2862_v144_ = _dafny.Map({d_2861_v143_: d_2860_v142_})
                d_2863_v145_: _dafny.Array
                nw447_ = _dafny.Array(None, 22)
                nw447_[int(0)] = d_2860_v142_
                nw447_[int(1)] = d_2860_v142_
                nw447_[int(2)] = d_2860_v142_
                nw447_[int(3)] = d_2860_v142_
                nw447_[int(4)] = d_2860_v142_
                nw447_[int(5)] = d_2860_v142_
                nw447_[int(6)] = d_2860_v142_
                nw447_[int(7)] = d_2860_v142_
                nw447_[int(8)] = d_2860_v142_
                nw447_[int(9)] = d_2860_v142_
                nw447_[int(10)] = d_2860_v142_
                nw447_[int(11)] = d_2860_v142_
                nw447_[int(12)] = d_2860_v142_
                nw447_[int(13)] = d_2860_v142_
                nw447_[int(14)] = d_2860_v142_
                nw447_[int(15)] = d_2860_v142_
                nw447_[int(16)] = d_2860_v142_
                nw447_[int(17)] = d_2860_v142_
                nw447_[int(18)] = d_2860_v142_
                nw447_[int(19)] = ((d_2862_v144_)[d_2861_v143_] if (d_2861_v143_) in (d_2862_v144_) else d_2860_v142_)
                nw447_[int(20)] = d_2860_v142_
                nw447_[int(21)] = d_2860_v142_
                d_2863_v145_ = nw447_
                d_2864_v146_: _dafny.Array
                d_2864_v146_ = d_2863_v145_
                rhs508_ = (d_2859_v141_ if self.f3 else d_2859_v141_)
                rhs509_ = d_2864_v146_
                rhs510_ = ((d_2848_v138_).f13) + ((d_2853_cf5_) * (p1))
                d_2859_v141_ = rhs508_
                d_2864_v146_ = rhs509_
                d_2853_cf5_ = rhs510_
                d_2865_v147_: _dafny.Array
                nw448_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_2865_v147_ = nw448_
                rhs511_ = default__.safeDivisionInt((d_2702_v6_).fm12(p1, p0, d_2853_cf5_, (self).f6, globalState), 381)
                rhs512_ = d_2865_v147_
                rhs513_ = d_2865_v147_
                rhs514_ = (True) not in (d_2697_v1_)
                lhs310_ = self
                d_2853_cf5_ = rhs511_
                d_2865_v147_ = rhs512_
                d_2865_v147_ = rhs513_
                lhs310_.f3 = rhs514_
            elif source44_.is_DC4:
                d_2866___mcc_h3_ = source44_.cf6
                d_2867___mcc_h4_ = source44_.cf7
                d_2868_cf7_ = d_2867___mcc_h4_
                d_2869_cf6_ = d_2866___mcc_h3_
                d_2870_v148_: _dafny.Set
                d_2870_v148_ = _dafny.Set({((self).f5)[default__.safeIndex(946, ((self).f5).length(0))], d_2868_cf7_})
                d_2871_v149_: _dafny.MultiSet
                d_2871_v149_ = _dafny.MultiSet([d_2870_v148_, d_2870_v148_])
                d_2872_v150_: _dafny.Map
                d_2872_v150_ = _dafny.Map({((d_2871_v149_) - (d_2871_v149_)).cardinality: (d_2868_cf7_) == (((self).f5)[default__.safeIndex(946, ((self).f5).length(0))])})
                d_2873_v151_: _dafny.Set
                d_2873_v151_ = _dafny.Set({835, (d_2848_v138_).f13})
                d_2872_v150_ = (d_2872_v150_).set((0) - (default__.safeModuloInt(len((self).f2), (d_2848_v138_).f13)), (d_2873_v151_).ispropersubset(d_2873_v151_))
                d_2874_v152_: _dafny.Array
                def lambda116_(d_2875_v137_):
                    def lambda117_(d_2876_i11_):
                        return default__.safeDivisionInt(d_2876_i11_, len(_dafny.SeqWithoutIsStrInference([d_2875_v137_ for d_2877_i12_ in range(default__.abs(-637))])))

                    return lambda117_

                init66_ = lambda116_(d_2847_v137_)
                nw449_ = _dafny.Array(None, 9)
                for i0_66_ in range(nw449_.length(0)):
                    nw449_[i0_66_] = init66_(i0_66_)
                d_2874_v152_ = nw449_
                index459_ = default__.safeIndex(452, (d_2874_v152_).length(0))
                (d_2874_v152_)[index459_] = p0
                d_2878_v153_: _dafny.Array
                nw450_ = _dafny.Array(False, 18)
                d_2878_v153_ = nw450_
                d_2879_v154_: C3
                nw451_ = C3()
                nw451_.ctor__(p1, d_2878_v153_, ((self).f5)[default__.safeIndex(946, ((self).f5).length(0))], (self).f1, (self).f2)
                d_2879_v154_ = nw451_
                d_2880_v155_: _dafny.Seq
                d_2880_v155_ = _dafny.SeqWithoutIsStrInference([d_2879_v154_])
                d_2881_v156_: _dafny.Seq
                d_2881_v156_ = _dafny.SeqWithoutIsStrInference([(d_2880_v155_)[default__.safeIndex((self).f4, len(d_2880_v155_))]])
                d_2882_v157_: _dafny.Map
                d_2882_v157_ = _dafny.Map({d_2881_v156_: p0})
                index460_ = default__.safeIndex(452, (d_2874_v152_).length(0))
                (d_2874_v152_)[index460_] = len(d_2882_v157_)
                d_2868_cf7_ = d_2868_cf7_
                r1 = (((d_2874_v152_)[default__.safeIndex(452, (d_2874_v152_).length(0))]) * ((self).f4)) + (((d_2848_v138_).f13) - ((0) - ((self).f4)))
            elif True:
                d_2883___mcc_h5_ = source44_.cf2
                d_2884_cf2_ = d_2883___mcc_h5_
                d_2885_v158_: _dafny.Array
                nw452_ = _dafny.Array(False, 19)
                d_2885_v158_ = nw452_
                d_2886_v159_: _dafny.Array
                def lambda118_(d_2887_cf2_):
                    def lambda119_(d_2888_i13_):
                        return d_2887_cf2_

                    return lambda119_

                init67_ = lambda118_(d_2884_cf2_)
                nw453_ = _dafny.Array(None, 8)
                for i0_67_ in range(nw453_.length(0)):
                    nw453_[i0_67_] = init67_(i0_67_)
                d_2886_v159_ = nw453_
                d_2889_v160_: _dafny.Seq
                d_2889_v160_ = _dafny.SeqWithoutIsStrInference([d_2885_v158_, d_2885_v158_, d_2885_v158_])
                d_2890_v161_: _dafny.Map
                d_2890_v161_ = _dafny.Map({(d_2848_v138_).f13: self.f3})
                index461_ = default__.safeIndex(946, ((self).f5).length(0))
                rhs515_ = (d_2889_v160_)[default__.safeIndex(p1, len(d_2889_v160_))]
                rhs516_ = d_2886_v159_
                rhs517_ = ((d_2848_v138_).f13) not in ((d_2890_v161_) | (d_2890_v161_))
                lhs311_ = (self).f5
                lhs312_ = default__.safeIndex(946, ((self).f5).length(0))
                d_2885_v158_ = rhs515_
                d_2886_v159_ = rhs516_
                lhs311_[lhs312_] = rhs517_
                d_2891_v162_: C1
                nw454_ = C1()
                nw454_.ctor__()
                d_2891_v162_ = nw454_
                d_2892_v163_: _dafny.Array
                nw455_ = _dafny.Array(D13.default()(), 25)
                d_2892_v163_ = nw455_
                d_2893_v164_: _dafny.Array
                nw456_ = _dafny.Array(_dafny.Array(None, 0), 27)
                d_2893_v164_ = nw456_
                d_2894_v165_: D22
                d_2894_v165_ = D22_DC54(d_2892_v163_)
                index462_ = default__.safeIndex(946, ((self).f5).length(0))
                rhs518_ = ((self).f4) + (p0)
                rhs519_ = ((self).f6) == (((d_2848_v138_).f13) < (p1))
                rhs520_ = (d_2894_v165_).cf85
                rhs521_ = d_2893_v164_
                lhs313_ = (self).f5
                lhs314_ = default__.safeIndex(946, ((self).f5).length(0))
                r1 = rhs518_
                lhs313_[lhs314_] = rhs519_
                d_2892_v163_ = rhs520_
                d_2893_v164_ = rhs521_
                d_2895_v166_: _dafny.Array
                def lambda120_(d_2896_v138_):
                    def lambda121_(d_2897_i14_):
                        return default__.safeDivisionInt(d_2897_i14_, (d_2896_v138_).f13)

                    return lambda121_

                init68_ = lambda120_(d_2848_v138_)
                nw457_ = _dafny.Array(None, 5)
                for i0_68_ in range(nw457_.length(0)):
                    nw457_[i0_68_] = init68_(i0_68_)
                d_2895_v166_ = nw457_
                index463_ = default__.safeIndex(907, (d_2895_v166_).length(0))
                (d_2895_v166_)[index463_] = 965
                d_2898_v167_: D11
                d_2898_v167_ = D11_DC30(_dafny.Map({731: d_2884_cf2_}))
                d_2899_v168_: D13
                d_2899_v168_ = D13_DC35(len(d_2890_v161_), (0) - ((d_2848_v138_).f13), False, (d_2848_v138_).f13, d_2898_v167_)
                d_2900_v169_: D18
                d_2900_v169_ = D18_DC44(default__.fm64((self).f6, (self).f6, (d_2899_v168_).cf59, globalState))
                index464_ = default__.safeIndex(907, (d_2895_v166_).length(0))
                rhs522_ = default__.safeDivisionInt((d_2848_v138_).f13, len(_dafny.Map({False: (self).f6})))
                rhs523_ = p1
                rhs524_ = d_2900_v169_
                lhs315_ = d_2895_v166_
                lhs316_ = default__.safeIndex(907, (d_2895_v166_).length(0))
                lhs315_[lhs316_] = rhs522_
                r2 = rhs523_
                d_2900_v169_ = rhs524_
            (self).f3 = ((self).f5)[default__.safeIndex(946, ((self).f5).length(0))]
        hi12_ = (self).f4
        for d_2901_i15_ in range(len(_dafny.SeqWithoutIsStrInference([(self).f2])), hi12_):
            r2 = (((self).f4) * ((self).f4) if (len(d_2697_v1_)) != (d_2901_i15_) else d_2901_i15_)
            d_2902_v170_: _dafny.Array
            def lambda122_(d_2903_i16_):
                return (d_2903_i16_) + (len((self).f2))

            init69_ = lambda122_
            nw458_ = _dafny.Array(None, 13)
            for i0_69_ in range(nw458_.length(0)):
                nw458_[i0_69_] = init69_(i0_69_)
            d_2902_v170_ = nw458_
            d_2904_v171_: _dafny.Map
            d_2904_v171_ = _dafny.Map({752: d_2902_v170_})
            d_2905_v172_: _dafny.Seq
            d_2905_v172_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_2906_v173_: _dafny.Map
            d_2906_v173_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([-329]): d_2902_v170_})
            d_2907_v174_: _dafny.Map
            d_2907_v174_ = _dafny.Map({((d_2904_v171_)[(d_2905_v172_)[default__.safeIndex(d_2901_i15_, len(d_2905_v172_))]] if ((d_2905_v172_)[default__.safeIndex(d_2901_i15_, len(d_2905_v172_))]) in (d_2904_v171_) else ((d_2906_v173_)[default__.fm20(False, (self).f6, (self).f6, d_2696_v0_, globalState)] if (default__.fm20(False, (self).f6, (self).f6, d_2696_v0_, globalState)) in (d_2906_v173_) else d_2902_v170_)): (self).f5})
            d_2908_v175_: _dafny.Set
            d_2908_v175_ = _dafny.Set({p0, d_2901_i15_, 220, (self).f4})
            d_2909_v176_: C9
            nw459_ = C9()
            nw459_.ctor__((self).f5, d_2907_v174_, 723, (self).f5, (self).f6, ((self).f1).set((self).f2, (d_2696_v0_)[default__.safeIndex(len(d_2908_v175_), len(d_2696_v0_))]), ((self).f2) + ((d_2702_v6_).fm16(d_2901_i15_, (self).f4, p1, globalState)))
            d_2909_v176_ = nw459_
            d_2910_v177_: C5
            nw460_ = C5()
            nw460_.ctor__((self).f6, ((self).f1) | (default__.fm62(p0, not(self.f3), globalState)), (self).f2)
            d_2910_v177_ = nw460_
            d_2910_v177_ = d_2910_v177_
            index465_ = default__.safeIndex(84, (d_2902_v170_).length(0))
            (d_2902_v170_)[index465_] = p1
            index466_ = default__.safeIndex(84, (d_2902_v170_).length(0))
            (d_2902_v170_)[index466_] = (d_2909_v176_).fm12(d_2901_i15_, p0, default__.safeModuloInt(71, p1), (default__.fm33(p0, globalState)).ispropersubset(d_2908_v175_), globalState)
        d_2911_v178_: str
        d_2911_v178_ = _dafny.CodePoint('w')
        d_2912_v179_: D8
        d_2912_v179_ = D8_DC24(d_2911_v178_, _dafny.Map({(self).fm13(globalState): (self).f6}))
        r0 = (d_2912_v179_).cf42
        r1 = default__.safeModuloInt(p0, len(_dafny.SeqWithoutIsStrInference([d_2911_v178_ for d_2913_i17_ in range(default__.abs(271))])))
        r2 = (self).f4
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2914_v0_: _dafny.Array
        def lambda123_(d_2915_i0_):
            return (d_2915_i0_) - (len(_dafny.Set({(self).f6})))

        init70_ = lambda123_
        nw461_ = _dafny.Array(None, 17)
        for i0_70_ in range(nw461_.length(0)):
            nw461_[i0_70_] = init70_(i0_70_)
        d_2914_v0_ = nw461_
        index467_ = default__.safeIndex(118, (d_2914_v0_).length(0))
        (d_2914_v0_)[index467_] = ((self).f4) + ((self).f4)
        index468_ = default__.safeIndex(118, (d_2914_v0_).length(0))
        (d_2914_v0_)[index468_] = ((self).f4) + ((self).f4)
        d_2916_v1_: D16
        d_2916_v1_ = D16_DC42((self).f6, False, (self).f6)
        if (d_2916_v1_).cf70:
            d_2917_v2_: _dafny.MultiSet
            d_2917_v2_ = _dafny.MultiSet([self.f3])
            d_2917_v2_ = (d_2917_v2_).set((self).f6, default__.abs((self).f4))
            d_2918_v3_: _dafny.Seq
            d_2918_v3_ = _dafny.SeqWithoutIsStrInference([(self).f4, (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]])
            d_2919_v4_: _dafny.Map
            d_2919_v4_ = _dafny.Map({(self).f5: d_2918_v3_})
            d_2920_v5_: _dafny.Map
            d_2920_v5_ = _dafny.Map({(d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]: d_2919_v4_})
            index469_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            (d_2914_v0_)[index469_] = len((((d_2920_v5_)[(d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]] if ((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]) in (d_2920_v5_) else d_2919_v4_)) | (d_2919_v4_))
            d_2921_v6_: D22
            d_2921_v6_ = D22_DC55((self).f2, (self).f4, d_2914_v0_)
            (self).f3 = (((self).f1)[(d_2921_v6_).cf86] if ((d_2921_v6_).cf86) in ((self).f1) else self.f3)
            (self).f3 = ((self).f6) and ((self).f6)
            d_2922_v7_: _dafny.Map
            d_2922_v7_ = _dafny.Map({((self).f4) * (len((self).f2)): (self).f2})
            d_2922_v7_ = (d_2922_v7_).set(-210, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feihncdx")))
        elif True:
            d_2923_v8_: str
            d_2923_v8_ = _dafny.CodePoint('r')
            d_2924_v9_: _dafny.Map
            d_2924_v9_ = _dafny.Map({(self).f2: ((self).f2 if True else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfkrh"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfkrh")))), d_2923_v8_))})
            d_2925_v10_: _dafny.Map
            d_2925_v10_ = _dafny.Map({(self).f4: (self).f2})
            d_2924_v9_ = (d_2924_v9_).set(((self).f2) + (((d_2925_v10_)[(self).f4] if ((self).f4) in (d_2925_v10_) else (self).f2)), (self).f2)
            d_2926_v11_: _dafny.Map
            d_2926_v11_ = _dafny.Map({(self).f6: (self).f6})
            d_2927_v12_: C4
            nw462_ = C4()
            nw462_.ctor__(((d_2926_v11_)[(self).f6] if ((self).f6) in (d_2926_v11_) else self.f3), default__.fm62((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))], (self).f6, globalState), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uckdk"))) + (_dafny.SeqWithoutIsStrInference([d_2923_v8_ for d_2928_i1_ in range(default__.abs(1))])))
            d_2927_v12_ = nw462_
            d_2929_v13_: _dafny.Seq
            d_2929_v13_ = _dafny.SeqWithoutIsStrInference([(self).f4, (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]])
            d_2930_v14_: _dafny.Seq
            d_2930_v14_ = _dafny.SeqWithoutIsStrInference([d_2929_v13_])
            d_2931_v15_: _dafny.Seq
            d_2931_v15_ = _dafny.SeqWithoutIsStrInference([((self).fm13(globalState) if True else (self).f6)])
            index470_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            rhs525_ = (d_2929_v13_)[default__.safeIndex(len(d_2926_v11_), len(d_2929_v13_))]
            rhs526_ = (d_2931_v15_)[default__.safeIndex((self).f4, len(d_2931_v15_))]
            rhs527_ = self.f3
            rhs528_ = _dafny.SeqWithoutIsStrInference([d_2929_v13_ for d_2932_i2_ in range(default__.abs(-793))])
            lhs317_ = d_2914_v0_
            lhs318_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            lhs319_ = self
            lhs320_ = self
            lhs317_[lhs318_] = rhs525_
            lhs319_.f3 = rhs526_
            lhs320_.f3 = rhs527_
            d_2930_v14_ = rhs528_
            d_2933_v16_: _dafny.Set
            d_2933_v16_ = default__.fm33(len((_dafny.SeqWithoutIsStrInference([(d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]])).set(default__.safeIndex((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]]))), (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))])), globalState)
            d_2934_v17_: _dafny.Map
            d_2934_v17_ = _dafny.Map({len((d_2930_v14_)[default__.safeIndex((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))], len(d_2930_v14_))]): (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]})
            d_2935_v18_: _dafny.Set
            d_2935_v18_ = _dafny.Set({((d_2934_v17_)[-718] if (-718) in (d_2934_v17_) else len((self).f2))})
            d_2933_v16_ = d_2935_v18_
            index471_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            (d_2914_v0_)[index471_] = (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]
        d_2936_v19_: _dafny.Map
        d_2936_v19_ = _dafny.Map({d_2914_v0_: (self).f5})
        d_2937_v20_: C9
        nw463_ = C9()
        nw463_.ctor__((self).f5, d_2936_v19_, 762, (self).f5, self.f3, (self).f1, (self).f2)
        d_2937_v20_ = nw463_
        d_2938_v21_: _dafny.Map
        d_2938_v21_ = _dafny.Map({not((self).f6): d_2937_v20_})
        d_2939_v22_: _dafny.Seq
        d_2939_v22_ = _dafny.SeqWithoutIsStrInference([(self).f2])
        d_2940_v23_: D23
        d_2940_v23_ = D23_DC57(d_2937_v20_)
        d_2938_v21_ = (d_2938_v21_).set(((d_2939_v22_)[default__.safeIndex(-462, len(d_2939_v22_))]) <= ((self).f2), (d_2940_v23_).cf90)
        hi13_ = (self).f4
        for d_2941_i3_ in range((self).f4, hi13_):
            index472_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            (d_2914_v0_)[index472_] = ((self).f4) + (default__.safeModuloInt(683, d_2941_i3_))
            d_2942_v24_: T1
            nw464_ = C2()
            nw464_.ctor__((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))], not(default__.fm3((self).f4, globalState)), (self).f1, (self).f2)
            d_2942_v24_ = nw464_
            d_2943_v25_: T1
            d_2943_v25_ = d_2942_v24_
            index473_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            index474_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            index475_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            rhs529_ = (0) - (d_2941_i3_)
            rhs530_ = ((0) - ((self).f4)) * (default__.safeModuloInt((d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))], (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]))
            rhs531_ = d_2914_v0_
            rhs532_ = (d_2914_v0_)[default__.safeIndex(118, (d_2914_v0_).length(0))]
            rhs533_ = (d_2943_v25_)
            lhs321_ = d_2914_v0_
            lhs322_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            lhs323_ = d_2914_v0_
            lhs324_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            lhs325_ = d_2914_v0_
            lhs326_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            lhs321_[lhs322_] = rhs529_
            lhs323_[lhs324_] = rhs530_
            d_2914_v0_ = rhs531_
            lhs325_[lhs326_] = rhs532_
            d_2942_v24_ = rhs533_
            d_2944_v27_: C3
            nw465_ = C3()
            def iife242_():
                coll90_ = _dafny.Map()
                compr_90_: _dafny.Seq
                for compr_90_ in (_dafny.MultiSet([(self).f2, (self).f2, (self).f2, (d_2942_v24_).f2])).Elements:
                    d_2945_v26_: _dafny.Seq = compr_90_
                    if (d_2945_v26_) in (_dafny.MultiSet([(self).f2, (self).f2, (self).f2, (d_2942_v24_).f2])):
                        coll90_[d_2945_v26_] = (self).f6
                return _dafny.Map(coll90_)
            nw465_.ctor__((self).f4, (self).f5, False, iife242_()
            , (self).f2)
            d_2944_v27_ = nw465_
            d_2946_v28_: _dafny.Seq
            d_2946_v28_ = _dafny.SeqWithoutIsStrInference([d_2944_v27_])
            d_2947_v29_: str
            d_2947_v29_ = _dafny.CodePoint('l')
            index476_ = default__.safeIndex(118, (d_2914_v0_).length(0))
            (d_2914_v0_)[index476_] = (len(d_2946_v28_)) + (len((_dafny.Map({d_2947_v29_: not(True)})) | (_dafny.Map({_dafny.CodePoint('p'): (self).f6}))))
            d_2948_v30_: _dafny.Array
            nw466_ = _dafny.Array(_dafny.CodePoint('D'), 2)
            d_2948_v30_ = nw466_
            index477_ = default__.safeIndex(88, (d_2948_v30_).length(0))
            (d_2948_v30_)[index477_] = d_2947_v29_
            index478_ = default__.safeIndex(88, (d_2948_v30_).length(0))
            (d_2948_v30_)[index478_] = (d_2947_v29_ if (self).f6 else d_2947_v29_)
        d_2949_v31_: _dafny.Array
        nw467_ = _dafny.Array(None, 6)
        nw467_[int(0)] = (self).f2
        nw467_[int(1)] = default__.fm2((self).f6, globalState)
        nw467_[int(2)] = ((self).f2) + ((self).f2)
        nw467_[int(3)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2950_i4_ in range(default__.abs(227))])) + ((self).f2)
        nw467_[int(4)] = (self).f2
        nw467_[int(5)] = ((self).f2) + ((self).f2)
        d_2949_v31_ = nw467_
        index479_ = default__.safeIndex(225, (d_2949_v31_).length(0))
        (d_2949_v31_)[index479_] = (self).f2
        index480_ = default__.safeIndex(225, (d_2949_v31_).length(0))
        (d_2949_v31_)[index480_] = (self).f2
        (self).f3 = ((self).f4) != (default__.safeModuloInt((self).f4, (self).f4))
        d_2951_v32_: D10
        d_2951_v32_ = D10_DC29((self).f6, not(self.f3))
        pat_let_tv57_ = d_2949_v31_
        pat_let_tv58_ = d_2949_v31_
        pat_let_tv59_ = d_2949_v31_
        pat_let_tv60_ = d_2949_v31_
        pat_let_tv61_ = d_2914_v0_
        pat_let_tv62_ = d_2914_v0_
        def lambda124_(source45_):
            if source45_.is_DC29:
                d_2952___mcc_h0_ = source45_.cf47
                d_2953___mcc_h1_ = source45_.cf48
                d_2954_cf48_ = d_2953___mcc_h1_
                d_2955_cf47_ = d_2952___mcc_h0_
                def iife243_():
                    coll91_ = _dafny.Map()
                    compr_91_: int
                    for compr_91_ in (_dafny.SeqWithoutIsStrInference([len((pat_let_tv58_)[default__.safeIndex(225, (pat_let_tv57_).length(0))])])).Elements:
                        d_2956_v33_: int = compr_91_
                        if (d_2956_v33_) in (_dafny.SeqWithoutIsStrInference([len((pat_let_tv60_)[default__.safeIndex(225, (pat_let_tv59_).length(0))])])):
                            coll91_[(d_2956_v33_) - (len(_dafny.SeqWithoutIsStrInference([d_2954_cf48_])))] = len(_dafny.Map({d_2955_cf47_: True}))
                    return _dafny.Map(coll91_)
                return iife243_()
                
            elif True:
                d_2957___mcc_h2_ = source45_.cf46
                d_2958_cf46_ = d_2957___mcc_h2_
                return (_dafny.Map({len(_dafny.Map({240: self.f3})): (pat_let_tv62_)[default__.safeIndex(118, (pat_let_tv61_).length(0))]})) | (d_2958_cf46_)

        r0 = lambda124_(d_2951_v32_)
        return r0

    @property
    def f6(self):
        return self._f6
