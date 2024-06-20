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
    def fm3(p0, p1, p2, p3, globalState):
        return D1_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")), -80)

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return _dafny.MultiSet([(_dafny.MultiSet([(_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qt"))))])).cardinality, len(_dafny.Set({True}))])).isdisjoint(_dafny.MultiSet([-958])), not (not(False)) or (False), not(False)])

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('j')

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        source0_ = D1_DC4(_dafny.Set({934}))
        if source0_.is_DC5:
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_0_i0_ in range(default__.abs(207))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1_i1_ in range(default__.abs(626))]))
        elif source0_.is_DC6:
            d_2___mcc_h0_ = source0_.cf5
            d_3___mcc_h1_ = source0_.cf6
            d_4_cf6_ = d_3___mcc_h1_
            d_5_cf5_ = d_2___mcc_h0_
            return d_5_cf5_
        elif True:
            d_6___mcc_h2_ = source0_.cf4
            d_7_cf4_ = d_6___mcc_h2_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_8_i2_ in range(default__.abs(997))])

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.Map({not (True) or (False): ((_dafny.MultiSet([len(_dafny.Map({False: -626}))])).cardinality) >= (984)})

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "je"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfsuhwou")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")) if True else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_9_i0_ in range(default__.abs(168))])))

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([780, 411, 574, len(_dafny.Set({True})), 362])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({True}))), 986, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, False]))]): False}))])), (_dafny.SeqWithoutIsStrInference([187])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, True})) for d_10_i0_ in range(default__.abs(161))])), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_11_i1_ in range(default__.abs(338))])), 623, len(_dafny.Map({True: len(_dafny.Set({len(_dafny.Map({True: True})), len(_dafny.SeqWithoutIsStrInference([(0) - (-883)]))}))}))]), (_dafny.SeqWithoutIsStrInference([940])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('i'): False})) for d_12_i2_ in range(default__.abs(221))]))})

    @staticmethod
    def fm15(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([710])): -724}), _dafny.Map({-115: 120})])

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        return (_dafny.MultiSet([False])).intersection((_dafny.MultiSet([not(True), True])).intersection(_dafny.MultiSet([False, False, False])))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: D0
            for compr_0_ in (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D0_DC3(False)])) + (_dafny.SeqWithoutIsStrInference([D0_DC3(False) for d_13_i0_ in range(default__.abs(-390))])))).Elements:
                d_14_v0_: D0 = compr_0_
                if (d_14_v0_) in (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([D0_DC3(False)])) + (_dafny.SeqWithoutIsStrInference([D0_DC3(False) for d_13_i0_ in range(default__.abs(-390))])))):
                    coll0_[d_14_v0_] = not((len(_dafny.SeqWithoutIsStrInference([True]))) == (713))
            return _dafny.Map(coll0_)
        return iife0_()
        

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm19(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jx")))})), -653, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pbdbqo"))), -821, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")))})]))).cardinality])).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), len((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([387]))})))]))) | (_dafny.MultiSet([(0) - (len(_dafny.Map({D10_DC30(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrnolacmv")))): 291})))]))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return D0_DC2()

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return (_dafny.Set({True, not(False), False, False})).intersection((_dafny.Set({True})).intersection(_dafny.Set({True})))

    @staticmethod
    def fm22(p0, p1, globalState):
        source1_ = _dafny.MultiSet([False])
        d_15___mcc_h0_ = source1_
        d_16_cf64_ = d_15___mcc_h0_
        return D3_DC8(_dafny.Map({not(False): False}))

    @staticmethod
    def fm23(p0, globalState):
        return (_dafny.Map({D3_DC8(_dafny.Map({False: not(True)})): False})) | ((_dafny.Map({D3_DC8(_dafny.Map({False: True})): False}) if False else _dafny.Map({D3_DC8(_dafny.Map({False: False})): True})))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return D7_DC21((_dafny.Map({D3_DC8(_dafny.Map({True: True})): False})) | (_dafny.Map({D3_DC8(_dafny.Map({True: True})): True})))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return D0_DC1((True) or (True), (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))})).ispropersubset(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([792]))).cardinality})))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        source2_ = D4_DC14(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_17_i0_ in range(default__.abs(315))]))])), not(not(True)), (_dafny.MultiSet([len(_dafny.Set({len(_dafny.Set({807}))})), len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: 865}))])).cardinality, not(True), len(_dafny.Map({not(not(True)): True})))
        if source2_.is_DC12:
            d_18___mcc_h0_ = source2_.cf14
            d_19___mcc_h1_ = source2_.cf15
            d_20___mcc_h2_ = source2_.cf16
            d_21_cf16_ = d_20___mcc_h2_
            d_22_cf15_ = d_19___mcc_h1_
            d_23_cf14_ = d_18___mcc_h0_
            return _dafny.Map({_dafny.Map({d_21_cf16_: d_22_cf15_}): d_22_cf15_})
        elif source2_.is_DC13:
            d_24___mcc_h3_ = source2_.cf17
            d_25___mcc_h4_ = source2_.cf18
            d_26___mcc_h5_ = source2_.cf19
            d_27_cf19_ = d_26___mcc_h5_
            d_28_cf18_ = d_25___mcc_h4_
            d_29_cf17_ = d_24___mcc_h3_
            return _dafny.Map({_dafny.Map({not(d_29_cf17_): len(_dafny.Set({d_27_cf19_}))}): d_28_cf18_})
        elif source2_.is_DC14:
            d_30___mcc_h6_ = source2_.cf20
            d_31___mcc_h7_ = source2_.cf21
            d_32___mcc_h8_ = source2_.cf22
            d_33___mcc_h9_ = source2_.cf23
            d_34___mcc_h10_ = source2_.cf24
            d_35_cf24_ = d_34___mcc_h10_
            d_36_cf23_ = d_33___mcc_h9_
            d_37_cf22_ = d_32___mcc_h8_
            d_38_cf21_ = d_31___mcc_h7_
            d_39_cf20_ = d_30___mcc_h6_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: _dafny.Map
                for compr_1_ in (_dafny.Map({_dafny.Map({True: (0) - (d_37_cf22_)}): not(d_38_cf21_)})).keys.Elements:
                    d_40_v0_: _dafny.Map = compr_1_
                    if (d_40_v0_) in (_dafny.Map({_dafny.Map({True: (0) - (d_37_cf22_)}): not(d_38_cf21_)})):
                        coll1_[d_40_v0_] = d_39_cf20_
                return _dafny.Map(coll1_)
            return (_dafny.Map({_dafny.Map({d_36_cf23_: len(_dafny.Set({d_35_cf24_}))}): d_39_cf20_})) | (iife1_()
            )
        elif source2_.is_DC11:
            d_41___mcc_h11_ = source2_.cf13
            d_42_cf13_ = d_41___mcc_h11_
            return (_dafny.Map({_dafny.Map({False: 438}): len(_dafny.Map({90: True}))})) | (_dafny.Map({_dafny.Map({False: 345}): 705}))
        elif True:
            d_43___mcc_h12_ = source2_.cf25
            d_44_cf25_ = d_43___mcc_h12_
            if True:
                def iife2_():
                    coll2_ = _dafny.Map()
                    compr_2_: _dafny.Map
                    for compr_2_ in (_dafny.MultiSet([_dafny.Map({True: 322})])).Elements:
                        d_45_v1_: _dafny.Map = compr_2_
                        if (d_45_v1_) in (_dafny.MultiSet([_dafny.Map({True: 322})])):
                            coll2_[d_45_v1_] = 288
                    return _dafny.Map(coll2_)
                return iife2_()
                
            elif True:
                return _dafny.Map({_dafny.Map({True: (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjrwfbgy"))): -569})))}): (0) - (len(_dafny.Map({486: -839})))})

    @staticmethod
    def fm27(p0, globalState):
        return D4_DC13(not (True) or (True), 915, (False if True else False))

    @staticmethod
    def fm28(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet([475])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([619, 473, len(_dafny.Map({True: False}))])))).cardinality, 547])

    @staticmethod
    def fm29(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('f'): False}) for d_46_i0_ in range(default__.abs(611))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.CodePoint('i'): not(True)})]))

    @staticmethod
    def fm30(globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qqwr"))])

    @staticmethod
    def fm31(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-711, 941):
                d_47_v0_: int = compr_3_
                if ((-711) <= (d_47_v0_)) and ((d_47_v0_) < (941)):
                    coll3_[default__.safeModuloInt(d_47_v0_, 433)] = True
            return _dafny.Map(coll3_)
        return (_dafny.Map({832: True})) | ((_dafny.Map({828: not(False)})) | (iife3_()
        ))

    @staticmethod
    def fm32(p0, p1, globalState):
        return (_dafny.Map({True: _dafny.Map({True: not(False)})})) | ((_dafny.Map({True: _dafny.Map({True: True})})) | (_dafny.Map({not(False): _dafny.Map({False: True})})))

    @staticmethod
    def fm33(globalState):
        return _dafny.SeqWithoutIsStrInference([(True) or (False), True])

    @staticmethod
    def fm34(p0, globalState):
        def iife4_():
            def iife8_():
                def iife10_():
                    coll10_ = _dafny.Map()
                    compr_10_: int
                    for compr_10_ in (_dafny.Map({690: False})).keys.Elements:
                        d_48_v1_: int = compr_10_
                        if (d_48_v1_) in (_dafny.Map({690: False})):
                            coll10_[default__.safeDivisionInt(d_48_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajo"))))] = 229
                    return _dafny.Map(coll10_)
                coll8_ = _dafny.Map()
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in (_dafny.Map({690: False})).keys.Elements:
                        d_48_v1_: int = compr_9_
                        if (d_48_v1_) in (_dafny.Map({690: False})):
                            coll9_[default__.safeDivisionInt(d_48_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajo"))))] = 229
                    return _dafny.Map(coll9_)
                compr_8_: int
                for compr_8_ in (iife9_()
                ).keys.Elements:
                    d_49_v0_: int = compr_8_
                    if (d_49_v0_) in (iife10_()
                    ):
                        coll8_[(d_49_v0_) + (len(_dafny.Map({True: D3_DC8(_dafny.Map({False: False}))})))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvslhv"))
                return _dafny.Map(coll8_)
            coll4_ = _dafny.Set()
            def iife5_():
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in (_dafny.Map({690: False})).keys.Elements:
                        d_48_v1_: int = compr_7_
                        if (d_48_v1_) in (_dafny.Map({690: False})):
                            coll7_[default__.safeDivisionInt(d_48_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajo"))))] = 229
                    return _dafny.Map(coll7_)
                coll5_ = _dafny.Map()
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: int
                    for compr_6_ in (_dafny.Map({690: False})).keys.Elements:
                        d_48_v1_: int = compr_6_
                        if (d_48_v1_) in (_dafny.Map({690: False})):
                            coll6_[default__.safeDivisionInt(d_48_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajo"))))] = 229
                    return _dafny.Map(coll6_)
                compr_5_: int
                for compr_5_ in (iife6_()
                ).keys.Elements:
                    d_49_v0_: int = compr_5_
                    if (d_49_v0_) in (iife7_()
                    ):
                        coll5_[(d_49_v0_) + (len(_dafny.Map({True: D3_DC8(_dafny.Map({False: False}))})))] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvslhv"))
                return _dafny.Map(coll5_)
            compr_4_: int
            for compr_4_ in (iife5_()
            ).keys.Elements:
                d_50_v2_: int = compr_4_
                if (d_50_v2_) in (iife8_()
                ):
                    coll4_ = coll4_.union(_dafny.Set([(d_50_v2_) - ((0) - (-61))]))
            return _dafny.Set(coll4_)
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(567, 438):
                d_51_v3_: int = compr_11_
                if ((567) <= (d_51_v3_)) and ((d_51_v3_) < (438)):
                    coll11_ = coll11_.union(_dafny.Set([(d_51_v3_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_52_i0_ in range(default__.abs(-116))])))]))
            return _dafny.Set(coll11_)
        return ((iife4_()
        ) | (iife11_()
        )) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjjpql")))])).cardinality, 404, len(_dafny.Map({508: _dafny.CodePoint('y')}))])).cardinality, 944}))

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")) for d_53_i0_ in range(default__.abs(570))])).Elements:
                d_54_v0_: _dafny.Seq = compr_12_
                if (d_54_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")) for d_53_i0_ in range(default__.abs(570))])):
                    coll12_ = coll12_.union(_dafny.Set([d_54_v0_]))
            return _dafny.Set(coll12_)
        def iife13_():
            coll13_ = _dafny.Set()
            compr_13_: _dafny.Seq
            for compr_13_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iofop")): _dafny.SeqWithoutIsStrInference([True, False])})).keys.Elements:
                d_55_v1_: _dafny.Seq = compr_13_
                if (d_55_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iofop")): _dafny.SeqWithoutIsStrInference([True, False])})):
                    coll13_ = coll13_.union(_dafny.Set([d_55_v1_]))
            return _dafny.Set(coll13_)
        def iife14_():
            coll14_ = _dafny.Set()
            compr_14_: _dafny.Seq
            for compr_14_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thbw"))])).Elements:
                d_56_v2_: _dafny.Seq = compr_14_
                if (d_56_v2_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thbw"))])):
                    coll14_ = coll14_.union(_dafny.Set([d_56_v2_]))
            return _dafny.Set(coll14_)
        return ((iife12_()
        ) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kejowhwr"))}))) | ((iife13_()
        ).intersection(iife14_()
        ))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return D3_DC9((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([949]))).intersection(_dafny.MultiSet([-759, 320])), D0_DC2(), 829)

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        def iife15_():
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: str
                for compr_17_ in (_dafny.MultiSet([_dafny.CodePoint('h')])).Elements:
                    d_57_v0_: str = compr_17_
                    if (d_57_v0_) in (_dafny.MultiSet([_dafny.CodePoint('h')])):
                        coll17_[d_57_v0_] = False
                return _dafny.Map(coll17_)
            coll15_ = _dafny.Set()
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: str
                for compr_16_ in (_dafny.MultiSet([_dafny.CodePoint('h')])).Elements:
                    d_57_v0_: str = compr_16_
                    if (d_57_v0_) in (_dafny.MultiSet([_dafny.CodePoint('h')])):
                        coll16_[d_57_v0_] = False
                return _dafny.Map(coll16_)
            compr_15_: str
            for compr_15_ in (iife16_()
            ).keys.Elements:
                d_58_v1_: str = compr_15_
                if (d_58_v1_) in (iife17_()
                ):
                    coll15_ = coll15_.union(_dafny.Set([d_58_v1_]))
            return _dafny.Set(coll15_)
        return (D3_DC9(_dafny.MultiSet([198, len(iife15_()
)]), D0_DC2(), 690)).cf11

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return (D12_DC35(True, 723, 135)).cf57

    @staticmethod
    def fm39(globalState):
        if not((-773) == (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([11])), len(_dafny.Set({True}))])))):
            return (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnr")))})) | (_dafny.Map({False: 643}))
        elif True:
            return (_dafny.Map({False: 824})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmlngaghv")))}))

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: int = int(0)
        d_59_v0_: bool
        d_59_v0_ = True
        if (d_59_v0_) or ((d_59_v0_) or (d_59_v0_)):
            d_60_v1_: int
            d_60_v1_ = 125
            d_61_v2_: _dafny.Seq
            d_61_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "or"))
            d_62_v3_: _dafny.MultiSet
            d_62_v3_ = _dafny.MultiSet([p0])
            d_63_v4_: D9
            d_63_v4_ = D9_DC27(d_62_v3_)
            d_64_v5_: _dafny.Seq
            d_64_v5_ = _dafny.SeqWithoutIsStrInference([d_63_v4_])
            d_65_v6_: _dafny.MultiSet
            d_65_v6_ = _dafny.MultiSet([d_59_v0_])
            d_66_v7_: _dafny.Map
            d_66_v7_ = _dafny.Map({True: d_65_v6_})
            d_67_v8_: _dafny.Map
            d_67_v8_ = _dafny.Map({len(d_66_v7_): d_60_v1_})
            nw0_ = _dafny.Array(None, 14)
            nw0_[int(0)] = d_60_v1_
            nw0_[int(1)] = len((d_61_v2_) + (d_61_v2_))
            nw0_[int(2)] = d_60_v1_
            nw0_[int(3)] = d_60_v1_
            nw0_[int(4)] = -822
            nw0_[int(5)] = (0) - (len(d_64_v5_))
            nw0_[int(6)] = d_60_v1_
            nw0_[int(7)] = d_60_v1_
            nw0_[int(8)] = d_60_v1_
            nw0_[int(9)] = d_60_v1_
            nw0_[int(10)] = -244
            nw0_[int(11)] = d_60_v1_
            nw0_[int(12)] = d_60_v1_
            nw0_[int(13)] = len(d_67_v8_)
            r0 = nw0_
            d_68_v9_: _dafny.Array
            nw1_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_68_v9_ = nw1_
            d_69_v10_: _dafny.Array
            def lambda0_(d_70_v1_):
                def lambda1_(d_71_i0_):
                    return (d_71_i0_) * (d_70_v1_)

                return lambda1_

            init0_ = lambda0_(d_60_v1_)
            nw2_ = _dafny.Array(None, 8)
            for i0_0_ in range(nw2_.length(0)):
                nw2_[i0_0_] = init0_(i0_0_)
            d_69_v10_ = nw2_
            index0_ = default__.safeIndex(51, (d_68_v9_).length(0))
            (d_68_v9_)[index0_] = d_69_v10_
            index1_ = default__.safeIndex(51, (d_68_v9_).length(0))
            nw3_ = _dafny.Array(int(0), 17)
            (d_68_v9_)[index1_] = nw3_
            r3 = default__.safeDivisionInt(d_60_v1_, (d_60_v1_ if d_59_v0_ else d_60_v1_))
            d_72_v11_: _dafny.Seq
            d_72_v11_ = _dafny.SeqWithoutIsStrInference([d_60_v1_])
            d_73_v12_: _dafny.Map
            d_73_v12_ = _dafny.Map({d_59_v0_: (d_72_v11_)[default__.safeIndex(617, len(d_72_v11_))]})
            d_74_v13_: _dafny.MultiSet
            d_74_v13_ = _dafny.MultiSet([-256, 656])
            rhs0_ = (d_74_v13_).ispropersubset(d_74_v13_)
            rhs1_ = (d_73_v12_) | (d_73_v12_)
            rhs2_ = d_59_v0_
            rhs3_ = _dafny.SeqWithoutIsStrInference([p0 for d_75_i1_ in range(default__.abs(-981))])
            rhs4_ = d_60_v1_
            r1 = rhs0_
            d_73_v12_ = rhs1_
            r1 = rhs2_
            d_61_v2_ = rhs3_
            d_60_v1_ = rhs4_
            r1 = d_59_v0_
        elif True:
            d_76_v14_: int
            d_76_v14_ = 1
            d_77_v15_: _dafny.Seq
            d_77_v15_ = _dafny.SeqWithoutIsStrInference([d_59_v0_])
            rhs5_ = (0) - (d_76_v14_)
            rhs6_ = 844
            rhs7_ = (0) - ((len((d_77_v15_) + (d_77_v15_))) + (d_76_v14_))
            lhs0_ = globalState
            lhs1_ = globalState
            lhs2_ = globalState
            lhs0_.f2 = rhs5_
            lhs1_.f2 = rhs6_
            lhs2_.f2 = rhs7_
            d_76_v14_ = 148
            d_78_v16_: _dafny.Seq
            d_78_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tuyexyypg"))
            d_79_v17_: C2
            nw4_ = C2()
            nw4_.ctor__(len(d_78_v16_), d_76_v14_)
            d_79_v17_ = nw4_
            (globalState).f2 = d_76_v14_
            (globalState).f2 = (default__.safeDivisionInt(default__.fm37(d_59_v0_, 805, not(d_59_v0_), globalState), d_76_v14_)) - (491)
        d_80_v18_: int
        d_80_v18_ = -649
        hi0_ = d_80_v18_
        for d_81_i2_ in range(d_80_v18_, hi0_):
            d_82_v19_: T1
            nw5_ = C3()
            nw5_.ctor__(False, (d_81_i2_) * (d_80_v18_), default__.safeModuloInt(d_80_v18_, d_81_i2_))
            d_82_v19_ = nw5_
            nw6_ = C1()
            nw6_.ctor__(d_81_i2_, (d_81_i2_) * (-405))
            d_82_v19_ = nw6_
            d_83_v20_: C2
            nw7_ = C2()
            nw7_.ctor__(d_80_v18_, d_82_v19_.f6)
            d_83_v20_ = nw7_
            d_84_v21_: _dafny.Seq
            d_84_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krfj"))
            d_84_v21_ = ((default__.fm7(p0, True, d_59_v0_, d_80_v18_, globalState)) + ((d_84_v21_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwr"))).set(default__.safeIndex(d_81_i2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwr")))), p0)))).set(default__.safeIndex(446, len((default__.fm7(p0, True, d_59_v0_, d_80_v18_, globalState)) + ((d_84_v21_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwr"))).set(default__.safeIndex(d_81_i2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwr")))), p0))))), (d_84_v21_)[default__.safeIndex(d_81_i2_, len(d_84_v21_))])
            d_85_v22_: _dafny.Array
            def lambda2_(d_86_v0_):
                def lambda3_(d_87_i3_):
                    return d_86_v0_

                return lambda3_

            init1_ = lambda2_(d_59_v0_)
            nw8_ = _dafny.Array(None, 27)
            for i0_1_ in range(nw8_.length(0)):
                nw8_[i0_1_] = init1_(i0_1_)
            d_85_v22_ = nw8_
            d_88_v23_: D7
            d_88_v23_ = D7_DC22(d_85_v22_, d_84_v21_)
            d_59_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuhemdl"))) < ((d_88_v23_).cf39)
        d_89_v24_: _dafny.Array
        nw9_ = _dafny.Array(int(0), 27)
        d_89_v24_ = nw9_
        d_90_v25_: D5
        d_90_v25_ = D5_DC16(d_89_v24_)
        d_91_v26_: _dafny.Map
        d_91_v26_ = _dafny.Map({d_90_v25_: d_80_v18_})
        d_92_v27_: _dafny.MultiSet
        d_92_v27_ = _dafny.MultiSet([d_91_v26_, _dafny.Map({d_90_v25_: d_80_v18_}), _dafny.Map({D5_DC16(d_89_v24_): d_80_v18_})])
        d_92_v27_ = (d_92_v27_).intersection(d_92_v27_)
        d_93_v28_: _dafny.Seq
        d_93_v28_ = _dafny.SeqWithoutIsStrInference([not(d_59_v0_)])
        d_94_v29_: _dafny.Set
        d_94_v29_ = _dafny.Set({d_93_v28_})
        d_95_v30_: _dafny.MultiSet
        d_95_v30_ = _dafny.MultiSet([d_59_v0_])
        hi1_ = (d_95_v30_).cardinality
        for d_96_i4_ in range(len((d_94_v29_) | (d_94_v29_)), hi1_):
            d_97_v31_: D4
            d_97_v31_ = D4_DC14(d_80_v18_, d_59_v0_, d_96_i4_, d_59_v0_, d_96_i4_)
            def iife18_(_pat_let0_0):
                def iife19_(d_98_dt__update__tmp_h0_):
                    def iife20_(_pat_let1_0):
                        def iife21_(d_99_dt__update_hcf20_h0_):
                            def iife22_(_pat_let2_0):
                                def iife23_(d_100_dt__update_hcf24_h0_):
                                    return D4_DC14(d_99_dt__update_hcf20_h0_, (d_98_dt__update__tmp_h0_).cf21, (d_98_dt__update__tmp_h0_).cf22, (d_98_dt__update__tmp_h0_).cf23, d_100_dt__update_hcf24_h0_)
                                return iife23_(_pat_let2_0)
                            return iife22_(d_96_i4_)
                        return iife21_(_pat_let1_0)
                    return iife20_(d_96_i4_)
                return iife19_(_pat_let0_0)
            d_95_v30_ = _dafny.MultiSet([not((iife18_(d_97_v31_)).cf21)])
            (globalState).f2 = d_80_v18_
            d_101_v32_: _dafny.Map
            d_101_v32_ = _dafny.Map({d_96_i4_: d_59_v0_})
            d_102_v33_: _dafny.Seq
            d_102_v33_ = _dafny.SeqWithoutIsStrInference([d_96_i4_, d_80_v18_])
            r1 = default__.fm38(d_101_v32_, (d_96_i4_) in (_dafny.SeqWithoutIsStrInference([len(d_102_v33_)])), True, (d_80_v18_) > (d_96_i4_), globalState)
            d_103_v34_: C1
            nw10_ = C1()
            nw10_.ctor__(-663, default__.fm37(d_59_v0_, d_80_v18_, d_59_v0_, globalState))
            d_103_v34_ = nw10_
        d_104_v35_: _dafny.Seq
        d_104_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyt"))
        d_104_v35_ = (d_104_v35_) + (d_104_v35_)
        d_105_v36_: D0
        d_105_v36_ = D0_DC0(d_59_v0_)
        d_106_v37_: D10
        d_106_v37_ = D10_DC29(_dafny.Map({d_80_v18_: len(d_104_v35_)}))
        rhs8_ = d_105_v36_
        rhs9_ = d_106_v37_
        rhs10_ = d_59_v0_
        rhs11_ = d_59_v0_
        d_105_v36_ = rhs8_
        d_106_v37_ = rhs9_
        r1 = rhs10_
        d_59_v0_ = rhs11_
        d_107_v38_: _dafny.Map
        d_107_v38_ = _dafny.Map({(d_104_v35_) + (d_104_v35_): d_89_v24_})
        r0 = ((d_107_v38_)[(d_104_v35_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcykgfojv"))).set(default__.safeIndex(-354, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcykgfojv")))), _dafny.CodePoint('r')))] if ((d_104_v35_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcykgfojv"))).set(default__.safeIndex(-354, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcykgfojv")))), _dafny.CodePoint('r')))) in (d_107_v38_) else d_89_v24_)
        r1 = d_59_v0_
        d_108_v39_: _dafny.MultiSet
        d_108_v39_ = _dafny.MultiSet([d_104_v35_, d_104_v35_])
        r2 = d_108_v39_
        d_109_v40_: _dafny.MultiSet
        d_109_v40_ = _dafny.MultiSet([d_80_v18_])
        d_110_v41_: _dafny.Seq
        d_110_v41_ = _dafny.SeqWithoutIsStrInference([747, (0) - (((d_109_v40_)[d_80_v18_] if (d_80_v18_) in (d_109_v40_) else d_80_v18_))])
        r3 = len(d_110_v41_)
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_111_v0_: int
        d_111_v0_ = -723
        d_112_v1_: str
        d_112_v1_ = _dafny.CodePoint('i')
        d_113_v2_: bool
        d_113_v2_ = True
        d_114_v3_: _dafny.Map
        d_114_v3_ = _dafny.Map({d_112_v1_: d_113_v2_})
        d_115_v4_: _dafny.MultiSet
        d_115_v4_ = _dafny.MultiSet([d_111_v0_, d_111_v0_, d_111_v0_])
        d_116_v5_: _dafny.Map
        d_116_v5_ = _dafny.Map({d_112_v1_: d_115_v4_})
        d_117_globalState_: GlobalState
        nw11_ = GlobalState()
        nw11_.ctor__(275, 628, -930, 254, (_dafny.MultiSet([d_111_v0_, (0) - (len(d_114_v3_)), d_111_v0_, d_111_v0_, d_111_v0_])) | (((d_116_v5_)[d_112_v1_] if (d_112_v1_) in (d_116_v5_) else d_115_v4_)))
        d_117_globalState_ = nw11_
        if d_113_v2_:
            d_118_v6_: _dafny.Array
            d_119_v7_: bool
            d_120_v8_: _dafny.MultiSet
            d_121_v9_: int
            out0_: _dafny.Array
            out1_: bool
            out2_: _dafny.MultiSet
            out3_: int
            out0_, out1_, out2_, out3_ = default__.m0(d_112_v1_, d_117_globalState_)
            d_118_v6_ = out0_
            d_119_v7_ = out1_
            d_120_v8_ = out2_
            d_121_v9_ = out3_
            d_122_v10_: C3
            nw12_ = C3()
            nw12_.ctor__(d_119_v7_, d_111_v0_, d_121_v9_)
            d_122_v10_ = nw12_
            d_123_v11_: C1
            nw13_ = C1()
            nw13_.ctor__(default__.safeModuloInt(d_121_v9_, d_121_v9_), d_111_v0_)
            d_123_v11_ = nw13_
            nw14_ = C1()
            nw14_.ctor__(default__.safeDivisionInt(d_121_v9_, (0) - (d_111_v0_)), d_111_v0_)
            d_123_v11_ = nw14_
            d_124_v12_: _dafny.Seq
            d_124_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_125_v13_: _dafny.Map
            d_125_v13_ = _dafny.Map({(d_124_v12_) < (d_124_v12_): d_111_v0_})
            d_126_v14_: _dafny.MultiSet
            d_126_v14_ = _dafny.MultiSet([d_123_v11_])
            d_125_v13_ = (d_125_v13_).set((not((d_122_v10_).f7)) or (d_113_v2_), ((d_126_v14_)[d_123_v11_] if (d_123_v11_) in (d_126_v14_) else d_111_v0_))
            d_119_v7_ = not(d_119_v7_)
        elif True:
            d_127_v15_: _dafny.Seq
            d_127_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw"))
            d_128_v16_: _dafny.Map
            d_128_v16_ = _dafny.Map({933: d_111_v0_})
            d_129_v17_: _dafny.Seq
            d_129_v17_ = _dafny.SeqWithoutIsStrInference([d_113_v2_])
            d_130_v18_: _dafny.MultiSet
            d_130_v18_ = _dafny.MultiSet([d_113_v2_, d_113_v2_])
            d_131_v19_: _dafny.MultiSet
            d_131_v19_ = d_130_v18_
            d_132_v20_: _dafny.Array
            nw15_ = _dafny.Array(None, 21)
            nw15_[int(0)] = len(d_127_v15_)
            nw15_[int(1)] = len(d_128_v16_)
            nw15_[int(2)] = default__.safeModuloInt(d_111_v0_, d_111_v0_)
            nw15_[int(3)] = default__.safeDivisionInt(d_111_v0_, len(_dafny.Set({d_111_v0_})))
            nw15_[int(4)] = len((d_129_v17_) + (d_129_v17_))
            nw15_[int(5)] = d_111_v0_
            nw15_[int(6)] = -903
            nw15_[int(7)] = (d_111_v0_) + ((d_115_v4_).cardinality)
            nw15_[int(8)] = d_111_v0_
            nw15_[int(9)] = d_111_v0_
            nw15_[int(10)] = d_111_v0_
            nw15_[int(11)] = len(d_127_v15_)
            nw15_[int(12)] = d_111_v0_
            nw15_[int(13)] = len(default__.fm35(True, len(_dafny.SeqWithoutIsStrInference([d_111_v0_ for d_133_i0_ in range(default__.abs(813))])), d_111_v0_, d_111_v0_, d_117_globalState_))
            nw15_[int(14)] = ((d_131_v19_)).cardinality
            nw15_[int(15)] = 281
            nw15_[int(16)] = d_111_v0_
            nw15_[int(17)] = d_111_v0_
            nw15_[int(18)] = -225
            nw15_[int(19)] = len(d_129_v17_)
            nw15_[int(20)] = d_111_v0_
            d_132_v20_ = nw15_
            index2_ = default__.safeIndex(190, (d_132_v20_).length(0))
            (d_132_v20_)[index2_] = d_111_v0_
            d_134_v21_: D0
            d_134_v21_ = D0_DC2()
            d_135_v22_: D3
            d_135_v22_ = D3_DC9(d_115_v4_, d_134_v21_, d_111_v0_)
            d_136_v23_: _dafny.Array
            nw16_ = _dafny.Array(None, 1)
            nw16_[int(0)] = d_135_v22_
            d_136_v23_ = nw16_
            pat_let_tv0_ = d_134_v21_
            index3_ = default__.safeIndex(694, (d_136_v23_).length(0))
            def iife24_(_pat_let3_0):
                def iife25_(d_137_dt__update__tmp_h0_):
                    def iife26_(_pat_let4_0):
                        def iife27_(d_138_dt__update_hcf10_h0_):
                            return D3_DC9((d_137_dt__update__tmp_h0_).cf9, d_138_dt__update_hcf10_h0_, (d_137_dt__update__tmp_h0_).cf11)
                        return iife27_(_pat_let4_0)
                    return iife26_(pat_let_tv0_)
                return iife25_(_pat_let3_0)
            (d_136_v23_)[index3_] = iife24_(default__.fm36(D5_DC19(d_111_v0_, d_112_v1_, d_111_v0_), d_112_v1_, d_113_v2_, d_117_globalState_))
            index4_ = default__.safeIndex(190, (d_132_v20_).length(0))
            index5_ = default__.safeIndex(694, (d_136_v23_).length(0))
            rhs12_ = not(d_113_v2_)
            rhs13_ = ((524) - (len((d_127_v15_).set(default__.safeIndex(d_111_v0_, len(d_127_v15_)), d_112_v1_)))) - ((0) - ((0) - (len(d_129_v17_))))
            rhs14_ = d_135_v22_
            lhs3_ = d_132_v20_
            lhs4_ = default__.safeIndex(190, (d_132_v20_).length(0))
            lhs5_ = d_136_v23_
            lhs6_ = default__.safeIndex(694, (d_136_v23_).length(0))
            d_113_v2_ = rhs12_
            lhs3_[lhs4_] = rhs13_
            lhs5_[lhs6_] = rhs14_
            d_139_v24_: D7
            d_139_v24_ = D7_DC23((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], d_127_v15_, len(d_128_v16_))
            d_140_v25_: _dafny.MultiSet
            d_140_v25_ = _dafny.MultiSet([d_127_v15_, (d_139_v24_).cf41, d_127_v15_])
            if (d_140_v25_) == (d_140_v25_):
                d_141_v26_: _dafny.Array
                d_142_v27_: bool
                d_143_v28_: _dafny.MultiSet
                d_144_v29_: int
                out4_: _dafny.Array
                out5_: bool
                out6_: _dafny.MultiSet
                out7_: int
                out4_, out5_, out6_, out7_ = default__.m0(d_112_v1_, d_117_globalState_)
                d_141_v26_ = out4_
                d_142_v27_ = out5_
                d_143_v28_ = out6_
                d_144_v29_ = out7_
                d_145_v30_: C3
                nw17_ = C3()
                nw17_.ctor__(d_142_v27_, d_111_v0_, d_111_v0_)
                d_145_v30_ = nw17_
                d_146_v31_: _dafny.Map
                d_146_v31_ = _dafny.Map({((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]) + (608): d_145_v30_})
                d_147_v32_: _dafny.Array
                nw18_ = _dafny.Array(None, 3)
                nw18_[int(0)] = (d_145_v30_).f7
                nw18_[int(1)] = ((d_114_v3_)[d_112_v1_] if (d_112_v1_) in (d_114_v3_) else (d_145_v30_).f7)
                nw18_[int(2)] = d_142_v27_
                d_147_v32_ = nw18_
                index6_ = default__.safeIndex(376, (d_147_v32_).length(0))
                (d_147_v32_)[index6_] = not(d_113_v2_)
                d_148_v33_: _dafny.Map
                d_148_v33_ = _dafny.Map({d_112_v1_: d_127_v15_})
                d_149_v34_: _dafny.Map
                d_149_v34_ = _dafny.Map({len(((d_148_v33_)[d_112_v1_] if (d_112_v1_) in (d_148_v33_) else d_127_v15_)): (d_145_v30_).f7})
                d_150_v35_: _dafny.MultiSet
                d_150_v35_ = _dafny.MultiSet([d_149_v34_, d_149_v34_])
                d_151_v36_: _dafny.Set
                d_151_v36_ = _dafny.Set({d_113_v2_})
                d_152_v37_: _dafny.Seq
                d_152_v37_ = _dafny.SeqWithoutIsStrInference([(d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], len(_dafny.Set({True, not(d_113_v2_)}))])
                index7_ = default__.safeIndex(376, (d_147_v32_).length(0))
                rhs15_ = d_146_v31_
                rhs16_ = (d_145_v30_).fm0((d_150_v35_) - (d_150_v35_), d_113_v2_, d_127_v15_, d_117_globalState_)
                rhs17_ = ((d_145_v30_).fm2(209, d_151_v36_, d_117_globalState_)) >= (len(d_152_v37_))
                rhs18_ = d_144_v29_
                lhs7_ = d_147_v32_
                lhs8_ = default__.safeIndex(376, (d_147_v32_).length(0))
                lhs9_ = d_117_globalState_
                d_146_v31_ = rhs15_
                lhs7_[lhs8_] = rhs16_
                d_142_v27_ = rhs17_
                lhs9_.f2 = rhs18_
                d_153_v38_: C3
                nw19_ = C3()
                nw19_.ctor__(d_113_v2_, (d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], d_111_v0_)
                d_153_v38_ = nw19_
                d_154_v39_: C1
                nw20_ = C1()
                nw20_.ctor__((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], (0) - ((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]))
                d_154_v39_ = nw20_
                d_155_v40_: _dafny.Seq
                d_155_v40_ = _dafny.SeqWithoutIsStrInference([d_147_v32_, d_147_v32_, d_147_v32_, d_147_v32_])
                d_156_v41_: D7
                d_156_v41_ = D7_DC22(d_147_v32_, d_127_v15_)
                d_157_v42_: _dafny.Array
                nw21_ = _dafny.Array(None, 13)
                nw21_[int(0)] = (d_155_v40_)[default__.safeIndex((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], len(d_155_v40_))]
                nw21_[int(1)] = d_147_v32_
                nw21_[int(2)] = d_147_v32_
                nw21_[int(3)] = d_147_v32_
                nw21_[int(4)] = (d_155_v40_)[default__.safeIndex((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], len(d_155_v40_))]
                nw21_[int(5)] = d_147_v32_
                nw21_[int(6)] = (d_156_v41_).cf38
                nw21_[int(7)] = d_147_v32_
                nw21_[int(8)] = d_147_v32_
                nw21_[int(9)] = (d_156_v41_).cf38
                nw21_[int(10)] = d_147_v32_
                nw21_[int(11)] = d_147_v32_
                nw21_[int(12)] = d_147_v32_
                d_157_v42_ = nw21_
                d_158_v43_: _dafny.Seq
                d_158_v43_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_112_v1_ for d_159_i1_ in range(default__.abs(167))])])
                index8_ = default__.safeIndex(190, (d_132_v20_).length(0))
                rhs19_ = (0) - (default__.safeDivisionInt(d_144_v29_, len(_dafny.SeqWithoutIsStrInference([d_113_v2_, (d_153_v38_).f7, d_113_v2_, (d_153_v38_).f7, True]))))
                rhs20_ = (d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]
                rhs21_ = d_157_v42_
                rhs22_ = len((d_158_v43_)[default__.safeIndex((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))], len(d_158_v43_))])
                lhs10_ = d_132_v20_
                lhs11_ = default__.safeIndex(190, (d_132_v20_).length(0))
                lhs12_ = d_117_globalState_
                lhs10_[lhs11_] = rhs19_
                lhs12_.f2 = rhs20_
                d_157_v42_ = rhs21_
                d_111_v0_ = rhs22_
            elif True:
                (d_117_globalState_).f2 = d_111_v0_
                d_160_v44_: C3
                nw22_ = C3()
                nw22_.ctor__(d_113_v2_, ((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]) + (11), (d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))])
                d_160_v44_ = nw22_
                d_161_v45_: C3
                d_161_v45_ = d_160_v44_
                d_160_v44_ = (d_161_v45_)
                d_162_v46_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_162_v46_ = nw23_
                index9_ = default__.safeIndex(558, (d_162_v46_).length(0))
                (d_162_v46_)[index9_] = d_130_v18_
                index10_ = default__.safeIndex(558, (d_162_v46_).length(0))
                rhs23_ = (0) - ((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))])
                rhs24_ = (d_130_v18_) | (d_130_v18_)
                lhs13_ = d_162_v46_
                lhs14_ = default__.safeIndex(558, (d_162_v46_).length(0))
                d_111_v0_ = rhs23_
                lhs13_[lhs14_] = rhs24_
                d_163_v47_: _dafny.Array
                def lambda4_(d_164_v44_):
                    def lambda5_(d_165_i2_):
                        return (d_164_v44_).f7

                    return lambda5_

                init2_ = lambda4_(d_160_v44_)
                nw24_ = _dafny.Array(None, 10)
                for i0_2_ in range(nw24_.length(0)):
                    nw24_[i0_2_] = init2_(i0_2_)
                d_163_v47_ = nw24_
                d_166_v48_: _dafny.Map
                d_166_v48_ = _dafny.Map({d_112_v1_: d_163_v47_})
                d_167_v49_: _dafny.Map
                d_167_v49_ = _dafny.Map({((d_166_v48_)[d_112_v1_] if (d_112_v1_) in (d_166_v48_) else d_163_v47_): (d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]})
                d_167_v49_ = (d_167_v49_) | ((d_167_v49_) | (_dafny.Map({d_163_v47_: len(d_127_v15_)})))
                index11_ = default__.safeIndex(190, (d_132_v20_).length(0))
                (d_132_v20_)[index11_] = (d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]
            d_168_v50_: C0
            nw25_ = C0()
            nw25_.ctor__(-984)
            d_168_v50_ = nw25_
            d_169_v51_: _dafny.MultiSet
            d_169_v51_ = _dafny.MultiSet([d_168_v50_])
            d_170_v52_: T1
            nw26_ = C1()
            nw26_.ctor__((((d_169_v51_)[d_168_v50_] if (d_168_v50_) in (d_169_v51_) else len(d_127_v15_))) * ((d_168_v50_).f8), ((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]) + (d_111_v0_))
            d_170_v52_ = nw26_
            d_170_v52_ = d_170_v52_
            d_171_v53_: _dafny.Seq
            d_171_v53_ = _dafny.SeqWithoutIsStrInference([d_111_v0_, ((d_128_v16_)[(d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]] if ((d_132_v20_)[default__.safeIndex(190, (d_132_v20_).length(0))]) in (d_128_v16_) else (d_170_v52_).f5), (d_170_v52_).f5])
            d_172_v54_: C3
            nw27_ = C3()
            nw27_.ctor__(d_113_v2_, len(d_171_v53_), (d_168_v50_).f8)
            d_172_v54_ = nw27_
            d_127_v15_ = d_127_v15_
        d_173_i3_: int
        d_173_i3_ = 0
        with _dafny.label("0"):
            while d_113_v2_:
                with _dafny.c_label("0"):
                    if (d_173_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_173_i3_ = (d_173_i3_) + (1)
                    rhs25_ = (d_113_v2_ if d_113_v2_ else not((d_115_v4_).ispropersubset(d_115_v4_)))
                    rhs26_ = d_113_v2_
                    rhs27_ = d_113_v2_
                    d_113_v2_ = rhs25_
                    d_113_v2_ = rhs26_
                    d_113_v2_ = rhs27_
                    d_174_v55_: _dafny.MultiSet
                    d_174_v55_ = _dafny.MultiSet([d_113_v2_, d_113_v2_])
                    d_175_v56_: _dafny.Map
                    d_175_v56_ = _dafny.Map({d_113_v2_: d_111_v0_})
                    (d_117_globalState_).f2 = ((d_174_v55_)[((d_114_v3_)[d_112_v1_] if (d_112_v1_) in (d_114_v3_) else d_113_v2_)] if (((d_114_v3_)[d_112_v1_] if (d_112_v1_) in (d_114_v3_) else d_113_v2_)) in (d_174_v55_) else (((d_175_v56_)[d_113_v2_] if (d_113_v2_) in (d_175_v56_) else d_111_v0_)) + (d_111_v0_))
                    d_113_v2_ = not(d_113_v2_)
                    d_176_v57_: _dafny.Seq
                    d_176_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqy"))
                    d_177_v58_: _dafny.Map
                    d_177_v58_ = _dafny.Map({d_113_v2_: d_112_v1_})
                    d_178_v59_: _dafny.Seq
                    d_178_v59_ = _dafny.SeqWithoutIsStrInference([d_113_v2_])
                    d_176_v57_ = default__.fm7(((d_177_v58_)[d_113_v2_] if (d_113_v2_) in (d_177_v58_) else d_112_v1_), d_113_v2_, (len(d_178_v59_)) == ((0) - (d_111_v0_)), default__.safeModuloInt((0) - (d_111_v0_), d_111_v0_), d_117_globalState_)
                    pass
            pass
        d_179_v60_: _dafny.Seq
        d_179_v60_ = _dafny.SeqWithoutIsStrInference([False])
        d_180_v61_: _dafny.Set
        d_180_v61_ = _dafny.Set({not((d_179_v60_)[default__.safeIndex(d_111_v0_, len(d_179_v60_))])})
        d_181_v62_: _dafny.Seq
        d_181_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uapcj"))
        d_182_v63_: _dafny.Array
        nw28_ = _dafny.Array(None, 9)
        nw28_[int(0)] = (d_111_v0_ if d_113_v2_ else (0) - (d_111_v0_))
        nw28_[int(1)] = (0) - (((0) - (d_111_v0_)) * (d_111_v0_))
        nw28_[int(2)] = len(_dafny.SeqWithoutIsStrInference([d_112_v1_ for d_183_i5_ in range(default__.abs(804))]))
        nw28_[int(3)] = d_111_v0_
        nw28_[int(4)] = len(d_180_v61_)
        nw28_[int(5)] = d_111_v0_
        nw28_[int(6)] = len((d_181_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qmhnk"))))
        nw28_[int(7)] = len((d_180_v61_).intersection(d_180_v61_))
        nw28_[int(8)] = d_111_v0_
        d_182_v63_ = nw28_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_182_v63_).length(0)):
            d_184_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_184_i4_)) and ((d_184_i4_) < ((d_182_v63_).length(0)))):
                (d_182_v63_)[(d_184_i4_)] = (d_184_i4_) + (len(d_179_v60_))
        d_185_v64_: D1
        d_185_v64_ = D1_DC6(d_181_v62_, (0) - (len((d_180_v61_ if d_113_v2_ else d_180_v61_))))
        source3_ = d_185_v64_
        if source3_.is_DC5:
            rhs28_ = d_112_v1_
            rhs29_ = d_113_v2_
            d_112_v1_ = rhs28_
            d_113_v2_ = rhs29_
            d_111_v0_ = (0) - (default__.safeModuloInt(-266, (d_111_v0_) + (d_111_v0_)))
            d_186_v65_: _dafny.Array
            d_187_v66_: bool
            d_188_v67_: _dafny.MultiSet
            d_189_v68_: int
            out8_: _dafny.Array
            out9_: bool
            out10_: _dafny.MultiSet
            out11_: int
            out8_, out9_, out10_, out11_ = default__.m0(d_112_v1_, d_117_globalState_)
            d_186_v65_ = out8_
            d_187_v66_ = out9_
            d_188_v67_ = out10_
            d_189_v68_ = out11_
            d_187_v66_ = True
        elif source3_.is_DC6:
            d_190___mcc_h0_ = source3_.cf5
            d_191___mcc_h1_ = source3_.cf6
            d_192_cf6_ = d_191___mcc_h1_
            d_193_cf5_ = d_190___mcc_h0_
            d_113_v2_ = ((d_180_v61_).issubset(d_180_v61_) if d_113_v2_ else False)
            d_194_v69_: C1
            nw29_ = C1()
            nw29_.ctor__(len((_dafny.Set({default__.fm37(d_113_v2_, d_111_v0_, d_113_v2_, d_117_globalState_)})) - (_dafny.Set({d_111_v0_, d_111_v0_, d_111_v0_, 678, d_192_cf6_}))), d_192_cf6_)
            d_194_v69_ = nw29_
            d_194_v69_ = d_194_v69_
            d_195_v70_: D9
            d_195_v70_ = D9_DC28(d_113_v2_, d_192_cf6_, d_111_v0_)
            d_113_v2_ = not((d_195_v70_).cf46)
            d_196_v71_: _dafny.Array
            nw30_ = _dafny.Array(False, 9)
            d_196_v71_ = nw30_
            index12_ = default__.safeIndex(220, (d_196_v71_).length(0))
            (d_196_v71_)[index12_] = d_113_v2_
            index13_ = default__.safeIndex(220, (d_196_v71_).length(0))
            (d_196_v71_)[index13_] = d_113_v2_
        elif True:
            d_197___mcc_h2_ = source3_.cf4
            d_198_cf4_ = d_197___mcc_h2_
            if ((d_111_v0_) <= (d_111_v0_) if d_113_v2_ else d_113_v2_):
                d_199_v72_: _dafny.Seq
                d_199_v72_ = _dafny.SeqWithoutIsStrInference([len(default__.fm7(d_112_v1_, d_113_v2_, d_113_v2_, 373, d_117_globalState_))])
                d_199_v72_ = d_199_v72_
                d_200_v73_: _dafny.Map
                d_200_v73_ = _dafny.Map({d_181_v62_: d_113_v2_})
                d_200_v73_ = (d_200_v73_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xajwiulrn")), d_113_v2_)
                d_113_v2_ = False
                (d_117_globalState_).f2 = d_111_v0_
                d_201_v74_: _dafny.Array
                nw31_ = _dafny.Array(None, 26)
                d_201_v74_ = nw31_
                d_202_v75_: _dafny.Map
                d_202_v75_ = _dafny.Map({d_111_v0_: True})
                d_203_v76_: _dafny.MultiSet
                d_203_v76_ = _dafny.MultiSet([((d_202_v75_)[len(d_198_cf4_)] if (len(d_198_cf4_)) in (d_202_v75_) else True)])
                d_204_v77_: _dafny.Map
                d_204_v77_ = _dafny.Map({d_113_v2_: d_113_v2_})
                d_205_v78_: _dafny.Map
                d_205_v78_ = _dafny.Map({False: ((d_203_v76_)[d_113_v2_] if (d_113_v2_) in (d_203_v76_) else len(d_204_v77_))})
                d_206_v79_: _dafny.Map
                d_206_v79_ = (d_205_v78_) | (d_205_v78_)
                rhs30_ = d_201_v74_
                rhs31_ = (0) - (d_111_v0_)
                rhs32_ = d_205_v78_
                rhs33_ = (default__.safeDivisionInt(d_111_v0_, d_111_v0_)) == (d_111_v0_)
                lhs15_ = d_117_globalState_
                d_201_v74_ = rhs30_
                lhs15_.f2 = rhs31_
                d_206_v79_ = rhs32_
                d_113_v2_ = rhs33_
            elif True:
                d_113_v2_ = d_113_v2_
                d_207_v80_: _dafny.Seq
                d_207_v80_ = _dafny.SeqWithoutIsStrInference([d_111_v0_])
                d_208_v81_: _dafny.Map
                d_208_v81_ = _dafny.Map({d_207_v80_: d_111_v0_})
                d_209_v82_: _dafny.Map
                d_209_v82_ = _dafny.Map({d_181_v62_: d_111_v0_})
                d_111_v0_ = len((d_208_v81_).set((d_207_v80_) + (_dafny.SeqWithoutIsStrInference([d_111_v0_])), ((d_209_v82_)[(d_181_v62_).set(default__.safeIndex(d_111_v0_, len(d_181_v62_)), d_112_v1_)] if ((d_181_v62_).set(default__.safeIndex(d_111_v0_, len(d_181_v62_)), d_112_v1_)) in (d_209_v82_) else d_111_v0_)))
                d_179_v60_ = d_179_v60_
                d_210_v83_: C1
                nw32_ = C1()
                nw32_.ctor__(896, 501)
                d_210_v83_ = nw32_
                d_211_v84_: _dafny.Map
                d_211_v84_ = _dafny.Map({d_111_v0_: d_210_v83_})
                d_212_v85_: _dafny.Array
                nw33_ = _dafny.Array(None, 6)
                nw33_[int(0)] = d_210_v83_
                nw33_[int(1)] = ((d_211_v84_)[d_111_v0_] if (d_111_v0_) in (d_211_v84_) else d_210_v83_)
                nw33_[int(2)] = d_210_v83_
                nw33_[int(3)] = d_210_v83_
                nw33_[int(4)] = d_210_v83_
                nw33_[int(5)] = d_210_v83_
                d_212_v85_ = nw33_
                index14_ = default__.safeIndex(628, (d_212_v85_).length(0))
                (d_212_v85_)[index14_] = d_210_v83_
                index15_ = default__.safeIndex(628, (d_212_v85_).length(0))
                (d_212_v85_)[index15_] = d_210_v83_
                d_213_v86_: _dafny.Map
                d_213_v86_ = _dafny.Map({602: d_113_v2_})
                d_214_v87_: _dafny.MultiSet
                d_214_v87_ = _dafny.MultiSet([d_213_v86_])
                (d_117_globalState_).f2 = (default__.safeModuloInt(d_111_v0_, d_111_v0_)) + (default__.fm37(d_113_v2_, d_111_v0_, (d_210_v83_).fm0(d_214_v87_, d_113_v2_, d_181_v62_, d_117_globalState_), d_117_globalState_))
            d_215_v88_: _dafny.Array
            d_216_v89_: bool
            d_217_v90_: _dafny.MultiSet
            d_218_v91_: int
            out12_: _dafny.Array
            out13_: bool
            out14_: _dafny.MultiSet
            out15_: int
            out12_, out13_, out14_, out15_ = default__.m0(d_112_v1_, d_117_globalState_)
            d_215_v88_ = out12_
            d_216_v89_ = out13_
            d_217_v90_ = out14_
            d_218_v91_ = out15_
            d_219_v92_: _dafny.Array
            def lambda6_(d_220_v89_):
                def lambda7_(d_221_i6_):
                    return d_220_v89_

                return lambda7_

            init3_ = lambda6_(d_216_v89_)
            nw34_ = _dafny.Array(None, 29)
            for i0_3_ in range(nw34_.length(0)):
                nw34_[i0_3_] = init3_(i0_3_)
            d_219_v92_ = nw34_
            d_222_v93_: D7
            d_222_v93_ = D7_DC22(d_219_v92_, _dafny.SeqWithoutIsStrInference([d_112_v1_ for d_223_i7_ in range(default__.abs(22))]))
            d_224_v94_: _dafny.Seq
            d_224_v94_ = _dafny.SeqWithoutIsStrInference([d_219_v92_, (d_222_v93_).cf38, d_219_v92_])
            (d_117_globalState_).f2 = len((d_224_v94_ if d_113_v2_ else d_224_v94_))
            d_225_v95_: C1
            nw35_ = C1()
            nw35_.ctor__(220, d_218_v91_)
            d_225_v95_ = nw35_
            d_225_v95_ = d_225_v95_
        d_111_v0_ = ((0) - (d_111_v0_)) + (d_111_v0_)
        d_226_v96_: _dafny.Map
        d_226_v96_ = _dafny.Map({not(d_113_v2_): d_113_v2_})
        d_227_v97_: _dafny.Seq
        d_227_v97_ = _dafny.SeqWithoutIsStrInference([(d_115_v4_).cardinality])
        d_228_v98_: _dafny.Map
        d_228_v98_ = _dafny.Map({(d_227_v97_)[default__.safeIndex(d_111_v0_, len(d_227_v97_))]: d_113_v2_})
        d_229_v99_: _dafny.MultiSet
        d_229_v99_ = _dafny.MultiSet([d_113_v2_])
        d_230_i8_: int
        d_230_i8_ = 0
        with _dafny.label("1"):
            pat_let_tv1_ = d_113_v2_
            pat_let_tv2_ = d_113_v2_
            pat_let_tv3_ = d_113_v2_
            pat_let_tv4_ = d_113_v2_
            def lambda8_(source4_):
                if source4_.is_DC5:
                    return pat_let_tv1_
                elif source4_.is_DC6:
                    d_236___mcc_h3_ = source4_.cf5
                    d_237___mcc_h4_ = source4_.cf6
                    d_238_cf6_ = d_237___mcc_h4_
                    d_239_cf5_ = d_236___mcc_h3_
                    return pat_let_tv2_
                elif True:
                    d_240___mcc_h5_ = source4_.cf4
                    d_241_cf4_ = d_240___mcc_h5_
                    return not (pat_let_tv3_) or (pat_let_tv4_)

            while lambda8_(default__.fm3(default__.fm37(False, len(d_226_v96_), default__.fm38(d_228_v98_, d_113_v2_, d_113_v2_, d_113_v2_, d_117_globalState_), d_117_globalState_), 377, _dafny.SeqWithoutIsStrInference([d_229_v99_]), (d_227_v97_).set(default__.safeIndex(d_111_v0_, len(d_227_v97_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "daff")))), d_117_globalState_)):
                with _dafny.c_label("1"):
                    if (d_230_i8_) >= (100):
                        raise _dafny.Break("1")
                    d_230_i8_ = (d_230_i8_) + (1)
                    d_231_v100_: _dafny.Array
                    d_232_v101_: bool
                    d_233_v102_: _dafny.MultiSet
                    d_234_v103_: int
                    out16_: _dafny.Array
                    out17_: bool
                    out18_: _dafny.MultiSet
                    out19_: int
                    out16_, out17_, out18_, out19_ = default__.m0(d_112_v1_, d_117_globalState_)
                    d_231_v100_ = out16_
                    d_232_v101_ = out17_
                    d_233_v102_ = out18_
                    d_234_v103_ = out19_
                    (d_117_globalState_).f2 = (d_234_v103_ if (d_234_v103_) < (d_234_v103_) else (d_229_v99_).cardinality)
                    d_235_v104_: C1
                    nw36_ = C1()
                    nw36_.ctor__(623, (d_111_v0_) * (d_111_v0_))
                    d_235_v104_ = nw36_
                    d_182_v63_ = d_231_v100_
                    pass
            pass
        d_242_v105_: _dafny.Seq
        d_242_v105_ = _dafny.SeqWithoutIsStrInference([d_115_v4_])
        hi2_ = len(d_181_v62_)
        for d_243_i9_ in range(len((d_242_v105_) + (d_242_v105_)), hi2_):
            d_113_v2_ = d_113_v2_
            d_244_v107_: _dafny.Set
            d_244_v107_ = _dafny.Set({default__.fm37(d_113_v2_, len(d_181_v62_), d_113_v2_, d_117_globalState_)})
            def iife28_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in (d_244_v107_).Elements:
                    d_245_v106_: int = compr_18_
                    if (d_245_v106_) in (d_244_v107_):
                        coll18_[default__.safeDivisionInt(d_245_v106_, d_111_v0_)] = d_113_v2_
                return _dafny.Map(coll18_)
            d_228_v98_ = ((iife28_()
            ) | (d_228_v98_) if d_113_v2_ else _dafny.Map({d_111_v0_: d_113_v2_}))
            d_246_v108_: _dafny.Map
            d_246_v108_ = _dafny.Map({d_179_v60_: d_111_v0_})
            d_246_v108_ = (d_246_v108_).set(d_179_v60_, d_111_v0_)
            d_113_v2_ = d_113_v2_
        d_247_v109_: _dafny.Array
        nw37_ = _dafny.Array(_dafny.Map({}), 15)
        d_247_v109_ = nw37_
        d_248_v110_: _dafny.Map
        d_248_v110_ = _dafny.Map({(0) - (d_111_v0_): default__.fm37(d_113_v2_, len(d_226_v96_), d_113_v2_, d_117_globalState_)})
        index16_ = default__.safeIndex(135, (d_247_v109_).length(0))
        (d_247_v109_)[index16_] = (d_248_v110_).set(d_111_v0_, d_111_v0_)
        index17_ = default__.safeIndex(135, (d_247_v109_).length(0))
        (d_247_v109_)[index17_] = ((d_248_v110_) | (_dafny.Map({536: d_111_v0_})) if not (d_113_v2_) or (d_113_v2_) else d_248_v110_)
        index18_ = default__.safeIndex(303, (d_182_v63_).length(0))
        (d_182_v63_)[index18_] = (-644) + (d_111_v0_)
        index19_ = default__.safeIndex(303, (d_182_v63_).length(0))
        rhs34_ = not(True)
        rhs35_ = ((0) - (default__.safeModuloInt(d_111_v0_, d_111_v0_))) * ((d_111_v0_) + (len(d_181_v62_)))
        rhs36_ = (0) - ((d_111_v0_) - (d_111_v0_))
        rhs37_ = d_111_v0_
        lhs16_ = d_117_globalState_
        lhs17_ = d_182_v63_
        lhs18_ = default__.safeIndex(303, (d_182_v63_).length(0))
        lhs19_ = d_117_globalState_
        d_113_v2_ = rhs34_
        lhs16_.f2 = rhs35_
        lhs17_[lhs18_] = rhs36_
        lhs19_.f2 = rhs37_
        d_249_v111_: _dafny.Map
        d_249_v111_ = _dafny.Map({d_113_v2_: (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]})
        d_249_v111_ = (d_249_v111_).set(not(d_113_v2_), d_111_v0_)
        d_250_v112_: D3
        d_250_v112_ = D3_DC8(d_226_v96_)
        d_251_v113_: _dafny.Seq
        d_251_v113_ = _dafny.SeqWithoutIsStrInference([d_250_v112_, d_250_v112_])
        d_252_v114_: D16
        d_252_v114_ = D16_DC40(d_251_v113_)
        d_251_v113_ = (d_251_v113_) + ((d_251_v113_) + ((d_252_v114_).cf66))
        d_253_v117_: _dafny.Map
        d_253_v117_ = _dafny.Map({(d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]: d_228_v98_})
        d_254_v119_: _dafny.Set
        d_254_v119_ = _dafny.Set({d_111_v0_})
        d_255_v121_: _dafny.Map
        d_255_v121_ = _dafny.Map({d_228_v98_: d_228_v98_})
        d_256_v122_: _dafny.Array
        nw38_ = _dafny.Array(None, 23)
        nw38_[int(0)] = d_228_v98_
        nw38_[int(1)] = d_228_v98_
        nw38_[int(2)] = d_228_v98_
        nw38_[int(3)] = d_228_v98_
        nw38_[int(4)] = _dafny.Map({284: False})
        nw38_[int(5)] = (d_228_v98_) | (d_228_v98_)
        nw38_[int(6)] = d_228_v98_
        nw38_[int(7)] = d_228_v98_
        def iife29_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in (d_115_v4_).Elements:
                d_257_v115_: int = compr_19_
                if (d_257_v115_) in (d_115_v4_):
                    coll19_[default__.safeDivisionInt(d_257_v115_, -416)] = d_113_v2_
            return _dafny.Map(coll19_)
        nw38_[int(8)] = (iife29_()
        ) | (d_228_v98_)
        nw38_[int(9)] = (d_228_v98_) | (d_228_v98_)
        nw38_[int(10)] = (d_228_v98_) | (d_228_v98_)
        nw38_[int(11)] = d_228_v98_
        nw38_[int(12)] = ((d_228_v98_).set((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_113_v2_)) | (d_228_v98_)
        nw38_[int(13)] = _dafny.Map({197: d_113_v2_})
        def iife30_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in (d_115_v4_).Elements:
                d_258_v116_: int = compr_20_
                if (d_258_v116_) in (d_115_v4_):
                    coll20_[default__.safeDivisionInt(d_258_v116_, (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])] = d_113_v2_
            return _dafny.Map(coll20_)
        nw38_[int(14)] = (d_228_v98_ if d_113_v2_ else iife30_()
        )
        nw38_[int(15)] = ((d_253_v117_)[len(d_181_v62_)] if (len(d_181_v62_)) in (d_253_v117_) else d_228_v98_)
        def iife31_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in (d_254_v119_).Elements:
                d_259_v118_: int = compr_21_
                if (d_259_v118_) in (d_254_v119_):
                    coll21_[default__.safeModuloInt(d_259_v118_, d_111_v0_)] = d_113_v2_
            return _dafny.Map(coll21_)
        nw38_[int(16)] = iife31_()
        
        def iife32_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(831, 188):
                d_260_v120_: int = compr_22_
                if ((831) <= (d_260_v120_)) and ((d_260_v120_) < (188)):
                    coll22_[(d_260_v120_) + ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])] = d_113_v2_
            return _dafny.Map(coll22_)
        nw38_[int(17)] = iife32_()
        
        nw38_[int(18)] = d_228_v98_
        nw38_[int(19)] = d_228_v98_
        nw38_[int(20)] = (_dafny.Map({(d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]: d_113_v2_})) | (d_228_v98_)
        nw38_[int(21)] = ((d_255_v121_)[d_228_v98_] if (d_228_v98_) in (d_255_v121_) else d_228_v98_)
        nw38_[int(22)] = _dafny.Map({(d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]: not(d_113_v2_)})
        d_256_v122_ = nw38_
        d_261_v123_: C2
        nw39_ = C2()
        nw39_.ctor__(d_111_v0_, default__.fm37(d_113_v2_, 799, default__.fm38(_dafny.Map({d_111_v0_: d_113_v2_}), d_113_v2_, d_113_v2_, d_113_v2_, d_117_globalState_), d_117_globalState_))
        d_261_v123_ = nw39_
        d_262_v124_: _dafny.MultiSet
        d_262_v124_ = _dafny.MultiSet([d_261_v123_])
        rhs38_ = d_256_v122_
        rhs39_ = (d_262_v124_).issubset((_dafny.MultiSet([d_261_v123_]) if d_113_v2_ else d_262_v124_))
        d_256_v122_ = rhs38_
        d_113_v2_ = rhs39_
        d_263_v125_: int
        d_264_v126_: bool
        out20_: int
        out21_: bool
        out20_, out21_ = (d_261_v123_).m1(d_117_globalState_)
        d_263_v125_ = out20_
        d_264_v126_ = out21_
        hi3_ = d_263_v125_
        for d_265_i10_ in range(d_263_v125_, hi3_):
            if d_113_v2_:
                d_226_v96_ = (d_226_v96_).set(not(not(d_264_v126_)), (106) > ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]))
                def iife33_():
                    coll23_ = _dafny.Set()
                    compr_23_: int
                    for compr_23_ in _dafny.IntegerRange(816, 77):
                        d_266_v128_: int = compr_23_
                        if ((816) <= (d_266_v128_)) and ((d_266_v128_) < (77)):
                            coll23_ = coll23_.union(_dafny.Set([default__.safeDivisionInt(d_266_v128_, d_265_i10_)]))
                    return _dafny.Set(coll23_)
                def iife34_():
                    coll24_ = _dafny.Set()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(276, 504):
                        d_267_v127_: int = compr_24_
                        if ((276) <= (d_267_v127_)) and ((d_267_v127_) < (504)):
                            coll24_ = coll24_.union(_dafny.Set([default__.safeModuloInt(d_267_v127_, (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])]))
                    return _dafny.Set(coll24_)
                d_113_v2_ = ((iife33_()
                ) | (d_254_v119_)).issubset((iife34_()
                ) | (_dafny.Set({(d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], (_dafny.MultiSet([d_111_v0_, d_263_v125_])).cardinality})))
                d_113_v2_ = ((d_115_v4_).cardinality) > ((d_261_v123_).fm2((d_261_v123_).fm2((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_180_v61_, d_117_globalState_), _dafny.Set({True}), d_117_globalState_))
                d_268_v129_: _dafny.Array
                nw40_ = _dafny.Array(_dafny.Seq({}), 27)
                d_268_v129_ = nw40_
                index20_ = default__.safeIndex(361, (d_268_v129_).length(0))
                (d_268_v129_)[index20_] = (d_179_v60_) + (d_179_v60_)
                index21_ = default__.safeIndex(361, (d_268_v129_).length(0))
                (d_268_v129_)[index21_] = _dafny.SeqWithoutIsStrInference([d_264_v126_, not (False) or (d_113_v2_)])
                d_269_v130_: _dafny.Array
                def lambda9_(d_270_v2_):
                    def lambda10_(d_271_i11_):
                        return (not(False)) or (not(d_270_v2_))

                    return lambda10_

                init4_ = lambda9_(d_113_v2_)
                nw41_ = _dafny.Array(None, 16)
                for i0_4_ in range(nw41_.length(0)):
                    nw41_[i0_4_] = init4_(i0_4_)
                d_269_v130_ = nw41_
                index22_ = default__.safeIndex(891, (d_269_v130_).length(0))
                (d_269_v130_)[index22_] = d_264_v126_
                d_272_v131_: D16
                d_272_v131_ = D16_DC42((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], len(d_181_v62_), d_112_v1_)
                d_273_v132_: _dafny.Map
                d_273_v132_ = _dafny.Map({(d_268_v129_)[default__.safeIndex(361, (d_268_v129_).length(0))]: len(d_181_v62_)})
                index23_ = default__.safeIndex(891, (d_269_v130_).length(0))
                rhs40_ = (d_272_v131_).cf70
                rhs41_ = d_113_v2_
                rhs42_ = not((len(d_273_v132_)) < (d_263_v125_))
                lhs20_ = d_269_v130_
                lhs21_ = default__.safeIndex(891, (d_269_v130_).length(0))
                d_112_v1_ = rhs40_
                d_113_v2_ = rhs41_
                lhs20_[lhs21_] = rhs42_
            elif True:
                d_274_v133_: _dafny.Map
                d_274_v133_ = _dafny.Map({(d_227_v97_) + ((d_227_v97_).set(default__.safeIndex(len(d_226_v96_), len(d_227_v97_)), d_263_v125_)): d_263_v125_})
                d_274_v133_ = (d_274_v133_).set(d_227_v97_, len(d_227_v97_))
                index24_ = default__.safeIndex(303, (d_182_v63_).length(0))
                (d_182_v63_)[index24_] = d_265_i10_
                (d_261_v123_).m2(d_263_v125_, d_117_globalState_)
                d_275_v134_: D7
                d_275_v134_ = D7_DC23(default__.fm37(d_113_v2_, d_111_v0_, d_264_v126_, d_117_globalState_), _dafny.SeqWithoutIsStrInference([d_112_v1_ for d_276_i12_ in range(default__.abs(681))]), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nltodsjko"))), (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]))
                d_275_v134_ = d_275_v134_
                d_182_v63_ = d_182_v63_
            if d_113_v2_:
                d_277_v135_: _dafny.Map
                d_277_v135_ = _dafny.Map({_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmmvfaq"))]): d_113_v2_})
                d_278_v136_: _dafny.MultiSet
                d_278_v136_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xsltdk")), d_181_v62_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etaeaj"))])
                d_279_v137_: _dafny.Map
                d_279_v137_ = _dafny.Map({d_264_v126_: d_181_v62_})
                d_277_v135_ = (d_277_v135_).set((d_278_v136_).set(((d_279_v137_)[d_113_v2_] if (d_113_v2_) in (d_279_v137_) else d_181_v62_), default__.abs(d_265_i10_)), (d_181_v62_) <= (d_181_v62_))
                d_280_v138_: C0
                nw42_ = C0()
                nw42_.ctor__(d_265_i10_)
                d_280_v138_ = nw42_
                d_281_v139_: C1
                nw43_ = C1()
                nw43_.ctor__(len(d_181_v62_), (d_265_i10_) - (d_263_v125_))
                d_281_v139_ = nw43_
                d_282_v140_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.Seq({}), 23)
                d_282_v140_ = nw44_
                index25_ = default__.safeIndex(398, (d_282_v140_).length(0))
                (d_282_v140_)[index25_] = _dafny.SeqWithoutIsStrInference([d_265_i10_])
                d_283_v141_: _dafny.Map
                d_283_v141_ = _dafny.Map({d_281_v139_: d_111_v0_})
                index26_ = default__.safeIndex(398, (d_282_v140_).length(0))
                rhs43_ = d_281_v139_
                rhs44_ = d_227_v97_
                rhs45_ = (d_113_v2_) == (d_113_v2_)
                rhs46_ = ((d_227_v97_)[default__.safeIndex(len(d_283_v141_), len(d_227_v97_))] if d_113_v2_ else d_111_v0_)
                lhs22_ = d_282_v140_
                lhs23_ = default__.safeIndex(398, (d_282_v140_).length(0))
                lhs24_ = d_117_globalState_
                d_281_v139_ = rhs43_
                lhs22_[lhs23_] = rhs44_
                d_113_v2_ = rhs45_
                lhs24_.f2 = rhs46_
                d_113_v2_ = d_264_v126_
                d_112_v1_ = d_112_v1_
            elif True:
                (d_261_v123_).m4((d_265_i10_) >= ((_dafny.MultiSet([d_264_v126_])).cardinality), d_117_globalState_)
                d_284_v142_: _dafny.Array
                nw45_ = _dafny.Array(None, 2)
                nw45_[int(0)] = d_181_v62_
                nw45_[int(1)] = (d_181_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmsgbdkrd")))
                d_284_v142_ = nw45_
                index27_ = default__.safeIndex(903, (d_284_v142_).length(0))
                (d_284_v142_)[index27_] = default__.fm11(74, len(d_254_v119_), d_111_v0_, d_111_v0_, d_117_globalState_)
                d_285_v143_: _dafny.Seq
                d_285_v143_ = _dafny.SeqWithoutIsStrInference([d_181_v62_])
                index28_ = default__.safeIndex(903, (d_284_v142_).length(0))
                (d_284_v142_)[index28_] = ((d_181_v62_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laqun")))) + ((d_181_v62_) + ((d_285_v143_)[default__.safeIndex(d_265_i10_, len(d_285_v143_))]))
                d_286_v144_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Seq({}), 28)
                d_286_v144_ = nw46_
                d_286_v144_ = d_286_v144_
                d_287_v145_: T0
                nw47_ = C1()
                nw47_.ctor__(len((_dafny.Map({d_264_v126_: d_263_v125_})) | (d_249_v111_)), 296)
                d_287_v145_ = nw47_
                rhs47_ = d_264_v126_
                rhs48_ = default__.fm28(d_263_v125_, d_117_globalState_)
                rhs49_ = d_264_v126_
                rhs50_ = d_287_v145_
                d_113_v2_ = rhs47_
                d_227_v97_ = rhs48_
                d_264_v126_ = rhs49_
                d_287_v145_ = rhs50_
                d_288_v146_: _dafny.Map
                d_288_v146_ = (d_249_v111_).set(d_113_v2_, d_263_v125_)
                d_289_v147_: _dafny.Array
                nw48_ = _dafny.Array(None, 26)
                nw48_[int(0)] = d_288_v146_
                nw48_[int(1)] = d_288_v146_
                nw48_[int(2)] = d_288_v146_
                nw48_[int(3)] = d_288_v146_
                nw48_[int(4)] = d_288_v146_
                nw48_[int(5)] = d_249_v111_
                nw48_[int(6)] = d_249_v111_
                nw48_[int(7)] = d_288_v146_
                nw48_[int(8)] = d_249_v111_
                nw48_[int(9)] = d_249_v111_
                nw48_[int(10)] = d_288_v146_
                nw48_[int(11)] = _dafny.Map({d_264_v126_: -772})
                nw48_[int(12)] = d_288_v146_
                nw48_[int(13)] = d_288_v146_
                nw48_[int(14)] = d_288_v146_
                nw48_[int(15)] = (d_288_v146_ if d_264_v126_ else d_288_v146_)
                nw48_[int(16)] = d_288_v146_
                nw48_[int(17)] = d_249_v111_
                nw48_[int(18)] = d_288_v146_
                nw48_[int(19)] = d_288_v146_
                nw48_[int(20)] = d_288_v146_
                nw48_[int(21)] = d_288_v146_
                nw48_[int(22)] = d_288_v146_
                nw48_[int(23)] = d_288_v146_
                nw48_[int(24)] = d_288_v146_
                nw48_[int(25)] = d_288_v146_
                d_289_v147_ = nw48_
                index29_ = default__.safeIndex(118, (d_289_v147_).length(0))
                (d_289_v147_)[index29_] = d_249_v111_
                index30_ = default__.safeIndex(118, (d_289_v147_).length(0))
                (d_289_v147_)[index30_] = d_288_v146_
            d_290_v148_: _dafny.Map
            d_290_v148_ = _dafny.Map({d_264_v126_: _dafny.Set({d_113_v2_, d_264_v126_, d_264_v126_})})
            d_291_v149_: C1
            nw49_ = C1()
            nw49_.ctor__(len(((d_290_v148_)[d_113_v2_] if (d_113_v2_) in (d_290_v148_) else _dafny.Set({d_113_v2_}))), d_111_v0_)
            d_291_v149_ = nw49_
            d_292_v150_: _dafny.Map
            d_292_v150_ = _dafny.Map({(0) - ((0) - (len(_dafny.Map({640: d_113_v2_})))): d_291_v149_})
            d_293_v151_: D11
            d_293_v151_ = D11_DC31(((d_292_v150_)[(0) - (d_111_v0_)] if ((0) - (d_111_v0_)) in (d_292_v150_) else d_291_v149_))
            pat_let_tv5_ = d_291_v149_
            def iife35_(_pat_let5_0):
                def iife36_(d_294_dt__update__tmp_h2_):
                    def iife37_(_pat_let6_0):
                        def iife38_(d_295_dt__update_hcf51_h0_):
                            return D11_DC31(d_295_dt__update_hcf51_h0_)
                        return iife38_(_pat_let6_0)
                    return iife37_(pat_let_tv5_)
                return iife36_(_pat_let5_0)
            source5_ = iife35_(d_293_v151_)
            if source5_.is_DC32:
                d_296___mcc_h6_ = source5_.cf52
                d_297___mcc_h7_ = source5_.cf53
                d_298___mcc_h8_ = source5_.cf54
                d_299_cf54_ = d_298___mcc_h8_
                d_300_cf53_ = d_297___mcc_h7_
                d_301_cf52_ = d_296___mcc_h6_
                d_302_v152_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_302_v152_ = nw50_
                d_303_v153_: D1
                d_303_v153_ = D1_DC5()
                rhs51_ = (d_261_v123_).fm9(d_303_v153_, d_113_v2_, d_117_globalState_)
                rhs52_ = d_113_v2_
                rhs53_ = d_111_v0_
                rhs54_ = ((d_254_v119_) | (d_254_v119_)) - (d_254_v119_)
                rhs55_ = d_302_v152_
                lhs25_ = d_117_globalState_
                d_299_cf54_ = rhs51_
                d_264_v126_ = rhs52_
                lhs25_.f2 = rhs53_
                d_254_v119_ = rhs54_
                d_302_v152_ = rhs55_
                d_300_cf53_ = (d_263_v125_) <= (d_265_i10_)
                d_304_v154_: C2
                nw51_ = C2()
                nw51_.ctor__(320, (0) - (d_111_v0_))
                d_304_v154_ = nw51_
                d_305_v155_: _dafny.Array
                nw52_ = _dafny.Array(False, 10)
                d_305_v155_ = nw52_
                d_306_v156_: D7
                d_306_v156_ = D7_DC22(d_305_v155_, d_181_v62_)
                d_305_v155_ = (d_306_v156_).cf38
            elif True:
                d_307___mcc_h9_ = source5_.cf51
                d_308_cf51_ = d_307___mcc_h9_
                d_113_v2_ = (d_265_i10_) != (d_263_v125_)
                d_113_v2_ = (d_229_v99_) == (default__.fm5(d_227_v97_, not(d_264_v126_), (d_308_cf51_).fm13(d_265_i10_, _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbeh"))]), d_117_globalState_), d_117_globalState_))
                d_309_v157_: T1
                nw53_ = C3()
                nw53_.ctor__(d_113_v2_, (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_265_i10_)
                d_309_v157_ = nw53_
                d_310_v158_: _dafny.Map
                d_310_v158_ = _dafny.Map({d_181_v62_: d_309_v157_})
                d_311_v159_: T0
                nw54_ = C2()
                nw54_.ctor__(d_263_v125_, ((d_249_v111_)[d_113_v2_] if (d_113_v2_) in (d_249_v111_) else len((d_310_v158_).set(d_181_v62_, d_309_v157_))))
                d_311_v159_ = nw54_
                d_312_v160_: T0
                d_312_v160_ = d_311_v159_
                rhs56_ = False
                rhs57_ = d_264_v126_
                rhs58_ = d_312_v160_
                d_264_v126_ = rhs56_
                d_113_v2_ = rhs57_
                d_312_v160_ = rhs58_
                d_313_v161_: _dafny.MultiSet
                d_313_v161_ = _dafny.MultiSet([d_181_v62_, d_181_v62_])
                d_314_v162_: C0
                nw55_ = C0()
                nw55_.ctor__((d_291_v149_).fm13((d_308_cf51_).fm13(d_309_v157_.f6, d_313_v161_, d_117_globalState_), d_313_v161_, d_117_globalState_))
                d_314_v162_ = nw55_
                d_315_v163_: _dafny.Seq
                d_315_v163_ = _dafny.SeqWithoutIsStrInference([d_314_v162_, d_314_v162_, d_314_v162_])
                d_316_v164_: _dafny.Map
                d_316_v164_ = _dafny.Map({(d_315_v163_) + (d_315_v163_): d_111_v0_})
                index31_ = default__.safeIndex(303, (d_182_v63_).length(0))
                rhs59_ = (0) - (((d_314_v162_).f8) * ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]))
                rhs60_ = (d_316_v164_).set(_dafny.SeqWithoutIsStrInference([d_314_v162_]), (0) - ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]))
                lhs26_ = d_182_v63_
                lhs27_ = default__.safeIndex(303, (d_182_v63_).length(0))
                lhs26_[lhs27_] = rhs59_
                d_316_v164_ = rhs60_
            d_317_v165_: T1
            nw56_ = C3()
            nw56_.ctor__((d_180_v61_).isdisjoint(_dafny.Set({d_113_v2_})), (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_111_v0_)
            d_317_v165_ = nw56_
            d_317_v165_ = d_317_v165_
        d_111_v0_ = default__.fm37(default__.fm38(d_228_v98_, d_113_v2_, d_113_v2_, d_264_v126_, d_117_globalState_), (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_113_v2_, d_117_globalState_)
        if (d_254_v119_) == (d_254_v119_):
            d_318_v166_: _dafny.Array
            nw57_ = _dafny.Array(None, 13)
            nw57_[int(0)] = d_182_v63_
            nw57_[int(1)] = d_182_v63_
            nw57_[int(2)] = d_182_v63_
            nw57_[int(3)] = d_182_v63_
            nw57_[int(4)] = d_182_v63_
            nw57_[int(5)] = d_182_v63_
            nw57_[int(6)] = d_182_v63_
            nw57_[int(7)] = d_182_v63_
            nw57_[int(8)] = d_182_v63_
            nw57_[int(9)] = d_182_v63_
            nw57_[int(10)] = d_182_v63_
            nw57_[int(11)] = d_182_v63_
            nw57_[int(12)] = d_182_v63_
            d_318_v166_ = nw57_
            index32_ = default__.safeIndex(574, (d_318_v166_).length(0))
            (d_318_v166_)[index32_] = d_182_v63_
            index33_ = default__.safeIndex(574, (d_318_v166_).length(0))
            (d_318_v166_)[index33_] = d_182_v63_
            d_319_v167_: _dafny.Map
            d_319_v167_ = _dafny.Map({len(d_180_v61_): d_182_v63_})
            d_320_v168_: _dafny.Map
            d_320_v168_ = _dafny.Map({(d_319_v167_) | (d_319_v167_): d_264_v126_})
            if ((d_320_v168_)[d_319_v167_] if (d_319_v167_) in (d_320_v168_) else d_264_v126_):
                d_321_v169_: D4
                d_321_v169_ = D4_DC14((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], d_113_v2_, len(d_181_v62_), not(not(d_264_v126_)), d_111_v0_)
                d_322_v170_: _dafny.Map
                d_322_v170_ = _dafny.Map({d_321_v169_: 8})
                d_323_v171_: _dafny.Map
                d_323_v171_ = _dafny.Map({d_113_v2_: d_322_v170_})
                d_324_v172_: T0
                nw58_ = C1()
                nw58_.ctor__(d_263_v125_, len(d_323_v171_))
                d_324_v172_ = nw58_
                d_325_v173_: _dafny.Map
                d_325_v173_ = _dafny.Map({d_263_v125_: d_226_v96_})
                d_326_v174_: _dafny.Array
                def lambda11_(d_327_v2_):
                    def lambda12_(d_328_i13_):
                        return d_327_v2_

                    return lambda12_

                init5_ = lambda11_(d_113_v2_)
                nw59_ = _dafny.Array(None, 28)
                for i0_5_ in range(nw59_.length(0)):
                    nw59_[i0_5_] = init5_(i0_5_)
                d_326_v174_ = nw59_
                index34_ = default__.safeIndex(623, (d_326_v174_).length(0))
                (d_326_v174_)[index34_] = (d_111_v0_) == (d_263_v125_)
                d_329_v175_: T1
                nw60_ = C1()
                nw60_.ctor__(d_111_v0_, d_111_v0_)
                d_329_v175_ = nw60_
                d_330_v176_: D9
                d_330_v176_ = D9_DC28(d_264_v126_, len(_dafny.SeqWithoutIsStrInference([d_329_v175_, d_329_v175_, d_329_v175_, d_329_v175_])), d_111_v0_)
                d_331_v177_: _dafny.Seq
                d_331_v177_ = _dafny.SeqWithoutIsStrInference([d_330_v176_])
                index35_ = default__.safeIndex(623, (d_326_v174_).length(0))
                index36_ = default__.safeIndex(303, (d_182_v63_).length(0))
                rhs61_ = d_324_v172_
                rhs62_ = d_325_v173_
                rhs63_ = not(d_113_v2_)
                rhs64_ = ((len(d_331_v177_)) - (len(d_179_v60_)) if d_113_v2_ else d_329_v175_.f6)
                lhs28_ = d_326_v174_
                lhs29_ = default__.safeIndex(623, (d_326_v174_).length(0))
                lhs30_ = d_182_v63_
                lhs31_ = default__.safeIndex(303, (d_182_v63_).length(0))
                d_324_v172_ = rhs61_
                d_325_v173_ = rhs62_
                lhs28_[lhs29_] = rhs63_
                lhs30_[lhs31_] = rhs64_
                index37_ = default__.safeIndex(623, (d_326_v174_).length(0))
                (d_326_v174_)[index37_] = d_264_v126_
                d_332_v178_: int
                d_333_v179_: bool
                out22_: int
                out23_: bool
                out22_, out23_ = (d_261_v123_).m1(d_117_globalState_)
                d_332_v178_ = out22_
                d_333_v179_ = out23_
                index38_ = default__.safeIndex(623, (d_326_v174_).length(0))
                rhs65_ = ((d_262_v124_) | (d_262_v124_)).set(d_261_v123_, default__.abs(328))
                rhs66_ = (d_326_v174_)[default__.safeIndex(623, (d_326_v174_).length(0))]
                rhs67_ = d_261_v123_
                rhs68_ = d_113_v2_
                rhs69_ = (d_179_v60_)[default__.safeIndex(len((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_329_v175_).f5])): d_112_v1_})).set((0) - (d_329_v175_.f6), d_112_v1_)), len(d_179_v60_))]
                lhs32_ = d_326_v174_
                lhs33_ = default__.safeIndex(623, (d_326_v174_).length(0))
                d_262_v124_ = rhs65_
                lhs32_[lhs33_] = rhs66_
                d_261_v123_ = rhs67_
                d_113_v2_ = rhs68_
                d_333_v179_ = rhs69_
                d_113_v2_ = (len(d_181_v62_)) != (998)
            elif True:
                (d_261_v123_).m4(not(True), d_117_globalState_)
                d_334_v180_: _dafny.Array
                nw61_ = _dafny.Array(False, 29)
                d_334_v180_ = nw61_
                index39_ = default__.safeIndex(549, (d_334_v180_).length(0))
                (d_334_v180_)[index39_] = not(not (d_113_v2_) or (True))
                index40_ = default__.safeIndex(549, (d_334_v180_).length(0))
                (d_334_v180_)[index40_] = d_113_v2_
                (d_117_globalState_).f2 = d_111_v0_
                d_335_v181_: int
                d_336_v182_: bool
                out24_: int
                out25_: bool
                out24_, out25_ = (d_261_v123_).m1(d_117_globalState_)
                d_335_v181_ = out24_
                d_336_v182_ = out25_
                d_115_v4_ = ((d_115_v4_).set(len(d_181_v62_), default__.abs(d_335_v181_))) - (d_115_v4_)
            d_337_v183_: _dafny.Array
            nw62_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_337_v183_ = nw62_
            d_338_v184_: _dafny.Array
            nw63_ = _dafny.Array(False, 16)
            d_338_v184_ = nw63_
            index41_ = default__.safeIndex(6, (d_337_v183_).length(0))
            (d_337_v183_)[index41_] = d_338_v184_
            index42_ = default__.safeIndex(6, (d_337_v183_).length(0))
            index43_ = default__.safeIndex(303, (d_182_v63_).length(0))
            rhs70_ = d_263_v125_
            rhs71_ = d_338_v184_
            rhs72_ = d_263_v125_
            rhs73_ = d_181_v62_
            rhs74_ = (d_181_v62_ if d_113_v2_ else d_181_v62_)
            lhs34_ = d_117_globalState_
            lhs35_ = d_337_v183_
            lhs36_ = default__.safeIndex(6, (d_337_v183_).length(0))
            lhs37_ = d_182_v63_
            lhs38_ = default__.safeIndex(303, (d_182_v63_).length(0))
            lhs34_.f2 = rhs70_
            lhs35_[lhs36_] = rhs71_
            lhs37_[lhs38_] = rhs72_
            d_181_v62_ = rhs73_
            d_181_v62_ = rhs74_
            d_339_v185_: _dafny.Array
            nw64_ = _dafny.Array(None, 2)
            d_339_v185_ = nw64_
            d_340_v186_: C1
            nw65_ = C1()
            nw65_.ctor__((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))], len(d_248_v110_))
            d_340_v186_ = nw65_
            index44_ = default__.safeIndex(47, (d_339_v185_).length(0))
            (d_339_v185_)[index44_] = d_340_v186_
            index45_ = default__.safeIndex(47, (d_339_v185_).length(0))
            (d_339_v185_)[index45_] = d_340_v186_
            rhs75_ = (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]
            rhs76_ = d_112_v1_
            lhs39_ = d_117_globalState_
            lhs39_.f2 = rhs75_
            d_112_v1_ = rhs76_
        elif True:
            d_181_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcqcxq"))
            (d_261_v123_).m2(default__.safeDivisionInt(len(d_181_v62_), (d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))]), d_117_globalState_)
            d_264_v126_ = not ((_dafny.CodePoint('a')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "miaw")))) or (((d_228_v98_)[461] if (461) in (d_228_v98_) else d_113_v2_))
            if (d_264_v126_) or (not(d_113_v2_)):
                d_341_v187_: C0
                nw66_ = C0()
                nw66_.ctor__((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])
                d_341_v187_ = nw66_
                d_342_v188_: _dafny.Seq
                d_342_v188_ = _dafny.SeqWithoutIsStrInference([d_249_v111_, d_249_v111_])
                d_343_v189_: _dafny.Set
                d_343_v189_ = _dafny.Set({d_112_v1_})
                d_344_v190_: _dafny.Array
                nw67_ = _dafny.Array(None, 11)
                nw67_[int(0)] = d_249_v111_
                nw67_[int(1)] = (d_249_v111_) | (d_249_v111_)
                nw67_[int(2)] = (d_249_v111_) | (d_249_v111_)
                nw67_[int(3)] = d_249_v111_
                nw67_[int(4)] = (d_342_v188_)[default__.safeIndex((d_341_v187_).f8, len(d_342_v188_))]
                nw67_[int(5)] = (_dafny.Map({d_264_v126_: d_263_v125_})) | (_dafny.Map({d_264_v126_: len(d_343_v189_)}))
                nw67_[int(6)] = default__.fm39(d_117_globalState_)
                nw67_[int(7)] = d_249_v111_
                nw67_[int(8)] = d_249_v111_
                nw67_[int(9)] = d_249_v111_
                nw67_[int(10)] = d_249_v111_
                d_344_v190_ = nw67_
                index46_ = default__.safeIndex(123, (d_344_v190_).length(0))
                (d_344_v190_)[index46_] = (_dafny.Map({not(d_264_v126_): d_263_v125_})) | (_dafny.Map({d_264_v126_: 193}))
                index47_ = default__.safeIndex(123, (d_344_v190_).length(0))
                rhs77_ = default__.safeModuloInt(d_263_v125_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awmjrtbl"))))
                rhs78_ = d_227_v97_
                rhs79_ = (d_342_v188_)[default__.safeIndex(d_263_v125_, len(d_342_v188_))]
                rhs80_ = (d_261_v123_).fm2((0) - ((0) - ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])), d_180_v61_, d_117_globalState_)
                rhs81_ = d_264_v126_
                lhs40_ = d_117_globalState_
                lhs41_ = d_344_v190_
                lhs42_ = default__.safeIndex(123, (d_344_v190_).length(0))
                lhs43_ = d_117_globalState_
                lhs40_.f2 = rhs77_
                d_227_v97_ = rhs78_
                lhs41_[lhs42_] = rhs79_
                lhs43_.f2 = rhs80_
                d_113_v2_ = rhs81_
                def iife39_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(243, 405):
                        d_345_v191_: int = compr_25_
                        if ((243) <= (d_345_v191_)) and ((d_345_v191_) < (405)):
                            coll25_[default__.safeModuloInt(d_345_v191_, ((d_115_v4_).set(len(d_180_v61_), default__.abs(len(d_226_v96_)))).cardinality)] = (d_341_v187_).f8
                    return _dafny.Map(coll25_)
                rhs82_ = (d_111_v0_) + ((d_182_v63_)[default__.safeIndex(303, (d_182_v63_).length(0))])
                rhs83_ = d_112_v1_
                rhs84_ = iife39_()

                lhs44_ = d_117_globalState_
                lhs44_.f2 = rhs82_
                d_112_v1_ = rhs83_
                d_248_v110_ = rhs84_
                d_181_v62_ = (d_181_v62_).set(default__.safeIndex(d_263_v125_, len(d_181_v62_)), d_112_v1_)
                d_346_v192_: _dafny.Set
                d_346_v192_ = (d_180_v61_) - (d_180_v61_)
                d_346_v192_ = d_180_v61_
            elif True:
                d_347_v193_: int
                d_348_v194_: bool
                out26_: int
                out27_: bool
                out26_, out27_ = (d_261_v123_).m1(d_117_globalState_)
                d_347_v193_ = out26_
                d_348_v194_ = out27_
                d_349_v195_: C2
                nw68_ = C2()
                nw68_.ctor__((d_227_v97_)[default__.safeIndex(d_111_v0_, len(d_227_v97_))], 136)
                d_349_v195_ = nw68_
                d_350_v196_: _dafny.Seq
                d_350_v196_ = _dafny.SeqWithoutIsStrInference([d_181_v62_])
                d_351_v197_: _dafny.Seq
                d_351_v197_ = _dafny.SeqWithoutIsStrInference([d_254_v119_])
                index48_ = default__.safeIndex(303, (d_182_v63_).length(0))
                index49_ = default__.safeIndex(303, (d_182_v63_).length(0))
                rhs85_ = ((d_228_v98_)[len((d_350_v196_)[default__.safeIndex(len((d_351_v197_)[default__.safeIndex(d_347_v193_, len(d_351_v197_))]), len(d_350_v196_))])] if (len((d_350_v196_)[default__.safeIndex(len((d_351_v197_)[default__.safeIndex(d_347_v193_, len(d_351_v197_))]), len(d_350_v196_))])) in (d_228_v98_) else d_348_v194_)
                rhs86_ = (d_111_v0_) + ((d_347_v193_) * (955))
                rhs87_ = d_263_v125_
                lhs45_ = d_182_v63_
                lhs46_ = default__.safeIndex(303, (d_182_v63_).length(0))
                lhs47_ = d_182_v63_
                lhs48_ = default__.safeIndex(303, (d_182_v63_).length(0))
                d_348_v194_ = rhs85_
                lhs45_[lhs46_] = rhs86_
                lhs47_[lhs48_] = rhs87_
                d_263_v125_ = d_263_v125_
                (d_117_globalState_).f2 = default__.safeModuloInt(d_111_v0_, d_263_v125_)
            d_113_v2_ = not (not(d_113_v2_)) or (d_264_v126_)
        _dafny.print(_dafny.string_of(d_111_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_112_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_113_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_v3_) == (_dafny.Map({_dafny.CodePoint('i'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v4_) == (_dafny.MultiSet([-723, -723, -723]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v5_) == (_dafny.Map({_dafny.CodePoint('i'): _dafny.MultiSet([-723, -723, -723])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_117_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_117_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_117_globalState_).f4) == (_dafny.MultiSet([-723, -723, -723, -723, -723, -723, -723, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_173_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v60_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v61_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_181_v62_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v63_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_185_v64_).cf5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v64_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v96_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v97_) == (_dafny.SeqWithoutIsStrInference([3]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v98_) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v99_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_230_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v105_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([-723, -723, -723])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_247_v109_)[0]) == (_dafny.Map({0: 690, 536: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v110_) == (_dafny.Map({0: 690}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_249_v111_) == (_dafny.Map({False: 0, True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_250_v112_).cf8) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_251_v113_) == (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True}))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v114_).cf66) == (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({False: True})), D3_DC8(_dafny.Map({False: True}))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v117_) == (_dafny.Map({0: _dafny.Map({690: True, 3: True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v119_) == (_dafny.Set({0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_255_v121_) == (_dafny.Map({_dafny.Map({690: True, 3: True}): _dafny.Map({690: True, 3: True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[0]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[1]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[2]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[3]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[4]) == (_dafny.Map({284: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[5]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[6]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[7]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[8]) == (_dafny.Map({2: False, 690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[9]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[10]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[11]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[12]) == (_dafny.Map({690: True, 3: True, 0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[13]) == (_dafny.Map({197: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[14]) == (_dafny.Map({-723: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[15]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[16]) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[17]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[18]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[19]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[20]) == (_dafny.Map({0: False, 690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[21]) == (_dafny.Map({690: True, 3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_256_v122_)[22]) == (_dafny.Map({0: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v124_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_263_v125_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_264_v126_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False)
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2)
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
        return lambda: D1_DC5()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({self.cf5.VerbatimString(True)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

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
        return lambda: D3_DC9(_dafny.MultiSet({}), D0.default()(), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC9(D3, NamedTuple('DC9', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC12(D4, NamedTuple('DC12', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf13 == __o.cf13
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
        return lambda: D5_DC17(int(0), _dafny.Array(None, 0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)

class D5_DC17(D5, NamedTuple('DC17', [('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)

class D6_DC20(D6, NamedTuple('DC20', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC22(_dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D7_DC24)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D7_DC25)

class D7_DC22(D7, NamedTuple('DC22', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf38)}, {self.cf39.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf40)}, {self.cf41.VerbatimString(True)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC24(D7, NamedTuple('DC24', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC24'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC24)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC25(D7, NamedTuple('DC25', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC25({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC25) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)

class D8_DC26(D8, NamedTuple('DC26', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC28(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D9_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D9_DC27)

class D9_DC28(D9, NamedTuple('DC28', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC28({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC28) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC27(D9, NamedTuple('DC27', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC27({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC27) and self.cf45 == __o.cf45
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

class D10_DC30(D10, NamedTuple('DC30', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC30({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC30) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC32(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D11_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)

class D11_DC32(D11, NamedTuple('DC32', [('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC32({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC32) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC31(D11, NamedTuple('DC31', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC34(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D12_DC36)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC34(D12, NamedTuple('DC34', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC36(D12, NamedTuple('DC36', [('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC36({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {self.cf62.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC36) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)

class D13_DC37(D13, NamedTuple('DC37', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D14_DC38)

class D14_DC38(D14, NamedTuple('DC38', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC38({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC38) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)

class D15_DC39(D15, NamedTuple('DC39', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC41(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D16_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D16_DC42)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D16_DC40)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D16_DC43)

class D16_DC41(D16, NamedTuple('DC41', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value
    def m1(self, globalState):
        pass

    def m2(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self._f0: int = int(0)
        self._f1: int = int(0)
        self._f3: int = int(0)
        self._f4: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4

class C0:
    def  __init__(self):
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f8):
        (self)._f8 = f8

    def fm4(self, p0, p1, p2, globalState):
        return (self).f8

    @property
    def f8(self):
        return self._f8

class C1(T0, T1):
    def  __init__(self):
        self._f6: int = int(0)
        self._f5: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f5, f6):
        (self)._f5 = f5
        (self).f6 = f6

    def fm0(self, p0, p1, p2, globalState):
        return (self.f6) > (188)

    def fm1(self, globalState):
        def iife40_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(110, 681):
                d_352_v0_: int = compr_26_
                if ((110) <= (d_352_v0_)) and ((d_352_v0_) < (681)):
                    coll26_ = coll26_.union(_dafny.Set([(d_352_v0_) * (self.f6)]))
            return _dafny.Set(coll26_)
        def iife41_():
            coll27_ = _dafny.Set()
            compr_27_: int
            for compr_27_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f6 for d_353_i0_ in range(default__.abs(467))])), self.f6])).Elements:
                d_354_v1_: int = compr_27_
                if (d_354_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f6 for d_353_i0_ in range(default__.abs(467))])), self.f6])):
                    coll27_ = coll27_.union(_dafny.Set([default__.safeDivisionInt(d_354_v1_, 499)]))
            return _dafny.Set(coll27_)
        return ((iife40_()
        ) - (iife41_()
        )).ispropersubset(_dafny.Set({231}))

    def fm2(self, p0, p1, globalState):
        return self.f6

    def fm12(self, p0, p1, p2, p3, globalState):
        return len((_dafny.Set({(self).f5})).intersection(_dafny.Set({(self).f5})))

    def fm13(self, p0, p1, globalState):
        return (self).f5

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_355_v0_: _dafny.Array
        def lambda13_(d_356_i0_):
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjwhec"))

        init6_ = lambda13_
        nw69_ = _dafny.Array(None, 1)
        for i0_6_ in range(nw69_.length(0)):
            nw69_[i0_6_] = init6_(i0_6_)
        d_355_v0_ = nw69_
        d_357_v1_: _dafny.Seq
        d_357_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oll"))
        index50_ = default__.safeIndex(241, (d_355_v0_).length(0))
        (d_355_v0_)[index50_] = d_357_v1_
        d_358_v2_: str
        d_358_v2_ = _dafny.CodePoint('r')
        index51_ = default__.safeIndex(241, (d_355_v0_).length(0))
        rhs88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alijtn"))
        rhs89_ = d_358_v2_
        lhs49_ = d_355_v0_
        lhs50_ = default__.safeIndex(241, (d_355_v0_).length(0))
        lhs49_[lhs50_] = rhs88_
        d_358_v2_ = rhs89_
        d_359_v3_: _dafny.Array
        nw70_ = _dafny.Array(False, 29)
        d_359_v3_ = nw70_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_359_v3_).length(0)):
            d_360_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_360_i1_)) and ((d_360_i1_) < ((d_359_v3_).length(0)))):
                (d_359_v3_)[(d_360_i1_)] = not (not((_dafny.SeqWithoutIsStrInference([d_358_v2_ for d_361_i2_ in range(default__.abs(-399))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djet"))))) or (False)
        d_362_v4_: _dafny.Array
        nw71_ = _dafny.Array(int(0), 11)
        d_362_v4_ = nw71_
        d_363_v5_: _dafny.Seq
        d_363_v5_ = _dafny.SeqWithoutIsStrInference([d_362_v4_])
        d_364_v6_: _dafny.Set
        d_364_v6_ = _dafny.Set({(d_363_v5_)[default__.safeIndex((self).f5, len(d_363_v5_))]})
        d_364_v6_ = (_dafny.Set({d_362_v4_, d_362_v4_})).intersection(d_364_v6_)
        index52_ = default__.safeIndex(779, (d_362_v4_).length(0))
        (d_362_v4_)[index52_] = (85) * ((self).f5)
        index53_ = default__.safeIndex(779, (d_362_v4_).length(0))
        (d_362_v4_)[index53_] = (self).f5
        d_365_v7_: _dafny.Seq
        d_365_v7_ = _dafny.SeqWithoutIsStrInference([(self).f5])
        d_366_v8_: _dafny.Set
        d_366_v8_ = _dafny.Set({d_365_v7_})
        d_367_v9_: bool
        d_367_v9_ = True
        d_368_i3_: int
        d_368_i3_ = 0
        with _dafny.label("2"):
            while ((d_366_v8_).intersection(d_366_v8_)).ispropersubset(default__.fm14(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwogcvy"))), d_367_v9_, globalState)):
                with _dafny.c_label("2"):
                    if (d_368_i3_) >= (100):
                        raise _dafny.Break("2")
                    d_368_i3_ = (d_368_i3_) + (1)
                    (self).f6 = 131
                    r1 = not(not (d_367_v9_) or (d_367_v9_))
                    d_369_v10_: _dafny.Map
                    d_369_v10_ = _dafny.Map({(self).f5: self.f6})
                    d_370_v11_: _dafny.Map
                    d_370_v11_ = _dafny.Map({(D1_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyctsial")), (self).f5)).cf6: _dafny.SeqWithoutIsStrInference([d_369_v10_, d_369_v10_])})
                    d_370_v11_ = (d_370_v11_).set((self).f5, default__.fm15(d_358_v2_, globalState))
                    index54_ = default__.safeIndex(615, (d_359_v3_).length(0))
                    (d_359_v3_)[index54_] = d_367_v9_
                    index55_ = default__.safeIndex(615, (d_359_v3_).length(0))
                    (d_359_v3_)[index55_] = False
                    pass
            pass
        if (d_365_v7_) == (_dafny.SeqWithoutIsStrInference([self.f6])):
            index56_ = default__.safeIndex(716, (d_359_v3_).length(0))
            (d_359_v3_)[index56_] = not(d_367_v9_)
            index57_ = default__.safeIndex(716, (d_359_v3_).length(0))
            (d_359_v3_)[index57_] = d_367_v9_
            d_371_v12_: D0
            d_371_v12_ = D0_DC0(not(False))
            source6_ = d_371_v12_
            if source6_.is_DC1:
                d_372___mcc_h0_ = source6_.cf1
                d_373___mcc_h1_ = source6_.cf2
                d_374_cf2_ = d_373___mcc_h1_
                d_375_cf1_ = d_372___mcc_h0_
                (self).f6 = (self).f5
                index58_ = default__.safeIndex(716, (d_359_v3_).length(0))
                (d_359_v3_)[index58_] = d_367_v9_
                d_376_v13_: C0
                nw72_ = C0()
                nw72_.ctor__(85)
                d_376_v13_ = nw72_
                d_377_v14_: C0
                nw73_ = C0()
                nw73_.ctor__((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])
                d_377_v14_ = nw73_
            elif source6_.is_DC2:
                d_357_v1_ = (d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]
                d_378_v15_: _dafny.Seq
                d_378_v15_ = _dafny.SeqWithoutIsStrInference([d_367_v9_])
                d_379_v16_: _dafny.Map
                d_379_v16_ = _dafny.Map({-348: not((d_378_v15_)[default__.safeIndex((0) - ((self).f5), len(d_378_v15_))])})
                d_380_v18_: C0
                nw74_ = C0()
                nw74_.ctor__(self.f6)
                d_380_v18_ = nw74_
                d_381_v19_: _dafny.Map
                d_381_v19_ = _dafny.Map({d_380_v18_: (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]})
                d_382_v20_: _dafny.Set
                d_382_v20_ = _dafny.Set({len(d_381_v19_)})
                d_383_v21_: _dafny.MultiSet
                d_383_v21_ = _dafny.MultiSet([d_367_v9_, d_367_v9_, (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]])
                def iife42_():
                    coll28_ = _dafny.Map()
                    compr_28_: int
                    for compr_28_ in (d_382_v20_).Elements:
                        d_384_v17_: int = compr_28_
                        if (d_384_v17_) in (d_382_v20_):
                            coll28_[default__.safeModuloInt(d_384_v17_, (d_380_v18_).f8)] = False
                    return _dafny.Map(coll28_)
                d_367_v9_ = ((_dafny.MultiSet([False])) | (d_383_v21_)).ispropersubset(default__.fm16(d_379_v16_, (self).f5, iife42_()
                , globalState))
                d_385_v22_: _dafny.Array
                def lambda14_(d_386_v3_):
                    def lambda15_(d_387_i4_):
                        return (d_386_v3_)[default__.safeIndex(716, (d_386_v3_).length(0))]

                    return lambda15_

                init7_ = lambda14_(d_359_v3_)
                nw75_ = _dafny.Array(None, 11)
                for i0_7_ in range(nw75_.length(0)):
                    nw75_[i0_7_] = init7_(i0_7_)
                d_385_v22_ = nw75_
                d_388_v23_: _dafny.Map
                d_388_v23_ = _dafny.Map({d_380_v18_: d_385_v22_})
                (globalState).f2 = default__.safeDivisionInt(self.f6, (self).fm12(d_367_v9_, len(d_388_v23_), (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], len(default__.fm17((d_380_v18_).f8, D0_DC0(d_367_v9_), self.f6, True, globalState)), globalState))
                d_389_v24_: _dafny.Array
                nw76_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_389_v24_ = nw76_
                d_390_v25_: D0
                d_390_v25_ = D0_DC1(False, (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))])
                pat_let_tv6_ = d_367_v9_
                index59_ = default__.safeIndex(716, (d_359_v3_).length(0))
                index60_ = default__.safeIndex(779, (d_362_v4_).length(0))
                def iife43_(_pat_let7_0):
                    def iife44_(d_391_dt__update__tmp_h0_):
                        def iife45_(_pat_let8_0):
                            def iife46_(d_392_dt__update_hcf2_h0_):
                                return D0_DC1((d_391_dt__update__tmp_h0_).cf1, d_392_dt__update_hcf2_h0_)
                            return iife46_(_pat_let8_0)
                        return iife45_(pat_let_tv6_)
                    return iife44_(_pat_let7_0)
                rhs90_ = (default__.safeModuloInt((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], (self).f5)) >= (self.f6)
                rhs91_ = (d_389_v24_ if (iife43_(d_390_v25_)).cf2 else d_389_v24_)
                rhs92_ = default__.safeDivisionInt((0) - ((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]), self.f6)
                rhs93_ = d_385_v22_
                rhs94_ = d_365_v7_
                lhs51_ = d_359_v3_
                lhs52_ = default__.safeIndex(716, (d_359_v3_).length(0))
                lhs53_ = d_362_v4_
                lhs54_ = default__.safeIndex(779, (d_362_v4_).length(0))
                lhs51_[lhs52_] = rhs90_
                d_389_v24_ = rhs91_
                lhs53_[lhs54_] = rhs92_
                d_385_v22_ = rhs93_
                d_365_v7_ = rhs94_
            elif source6_.is_DC3:
                d_393___mcc_h2_ = source6_.cf3
                d_394_cf3_ = d_393___mcc_h2_
                d_394_cf3_ = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
                d_357_v1_ = (d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]
                def lambda16_(d_395_i5_):
                    return default__.safeModuloInt(d_395_i5_, (self).f5)

                init8_ = lambda16_
                nw77_ = _dafny.Array(None, 27)
                for i0_8_ in range(nw77_.length(0)):
                    nw77_[i0_8_] = init8_(i0_8_)
                d_362_v4_ = nw77_
                d_396_v26_: _dafny.Array
                def lambda17_(d_397_v9_):
                    def lambda18_(d_398_i6_):
                        return d_397_v9_

                    return lambda18_

                init9_ = lambda17_(d_367_v9_)
                nw78_ = _dafny.Array(None, 18)
                for i0_9_ in range(nw78_.length(0)):
                    nw78_[i0_9_] = init9_(i0_9_)
                d_396_v26_ = nw78_
                d_399_v27_: _dafny.Map
                d_399_v27_ = _dafny.Map({d_394_cf3_: (self).f5})
                d_400_v28_: _dafny.Map
                d_400_v28_ = _dafny.Map({d_396_v26_: _dafny.Map({d_394_cf3_: (self).f5})})
                d_401_v29_: _dafny.Array
                nw79_ = _dafny.Array(None, 3)
                nw79_[int(0)] = (_dafny.Map({d_396_v26_: _dafny.Map({True: self.f6})})).set(d_396_v26_, d_399_v27_)
                nw79_[int(1)] = d_400_v28_
                nw79_[int(2)] = (d_400_v28_) | (d_400_v28_)
                d_401_v29_ = nw79_
                index61_ = default__.safeIndex(61, (d_401_v29_).length(0))
                (d_401_v29_)[index61_] = d_400_v28_
                d_402_v30_: _dafny.Map
                d_402_v30_ = _dafny.Map({d_367_v9_: d_367_v9_})
                d_403_v31_: D3
                d_403_v31_ = D3_DC8(d_402_v30_)
                index62_ = default__.safeIndex(61, (d_401_v29_).length(0))
                index63_ = default__.safeIndex(716, (d_359_v3_).length(0))
                rhs95_ = len((d_403_v31_).cf8)
                rhs96_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(len((d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]), self.f6), (_dafny.MultiSet([(d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))], d_357_v1_])).cardinality))
                rhs97_ = (_dafny.Map({d_396_v26_: d_399_v27_})) | (_dafny.Map({d_396_v26_: d_399_v27_}))
                rhs98_ = ((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]) >= (self.f6)
                lhs55_ = self
                lhs56_ = globalState
                lhs57_ = d_401_v29_
                lhs58_ = default__.safeIndex(61, (d_401_v29_).length(0))
                lhs59_ = d_359_v3_
                lhs60_ = default__.safeIndex(716, (d_359_v3_).length(0))
                lhs55_.f6 = rhs95_
                lhs56_.f2 = rhs96_
                lhs57_[lhs58_] = rhs97_
                lhs59_[lhs60_] = rhs98_
            elif True:
                d_404___mcc_h3_ = source6_.cf0
                d_405_cf0_ = d_404___mcc_h3_
                index64_ = default__.safeIndex(779, (d_362_v4_).length(0))
                (d_362_v4_)[index64_] = ((0) - (self.f6)) - (self.f6)
                d_406_v32_: _dafny.MultiSet
                d_406_v32_ = _dafny.MultiSet([d_365_v7_])
                d_406_v32_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]]), _dafny.SeqWithoutIsStrInference([(self).fm12(d_367_v9_, (self).f5, d_367_v9_, (self).fm12((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], d_367_v9_, self.f6, globalState), globalState)]), (D4_DC11(d_365_v7_)).cf13, d_365_v7_])
                d_407_v33_: C0
                nw80_ = C0()
                nw80_.ctor__((self).f5)
                d_407_v33_ = nw80_
                d_408_v34_: _dafny.Seq
                d_408_v34_ = _dafny.SeqWithoutIsStrInference([d_405_cf0_])
                d_409_v35_: _dafny.MultiSet
                d_409_v35_ = _dafny.MultiSet([len(d_408_v34_), 979, 490])
                d_409_v35_ = _dafny.MultiSet((d_365_v7_) + ((d_365_v7_) + (d_365_v7_)))
            d_410_v36_: _dafny.Map
            d_410_v36_ = _dafny.Map({134: d_367_v9_})
            d_411_v37_: _dafny.MultiSet
            d_411_v37_ = _dafny.MultiSet([d_410_v36_])
            d_412_v38_: _dafny.Map
            d_412_v38_ = _dafny.Map({False: (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]})
            d_413_v39_: _dafny.Set
            d_413_v39_ = _dafny.Set({(d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]})
            d_414_v40_: _dafny.Array
            nw81_ = _dafny.Array(None, 21)
            nw81_[int(0)] = d_367_v9_
            nw81_[int(1)] = d_367_v9_
            nw81_[int(2)] = not((self).fm0(d_411_v37_, True, d_357_v1_, globalState))
            nw81_[int(3)] = False
            nw81_[int(4)] = not(True)
            nw81_[int(5)] = (self).fm1(globalState)
            nw81_[int(6)] = (d_367_v9_) not in (d_412_v38_)
            nw81_[int(7)] = d_367_v9_
            nw81_[int(8)] = d_367_v9_
            nw81_[int(9)] = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
            nw81_[int(10)] = d_367_v9_
            nw81_[int(11)] = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
            nw81_[int(12)] = d_367_v9_
            nw81_[int(13)] = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
            nw81_[int(14)] = d_367_v9_
            nw81_[int(15)] = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
            nw81_[int(16)] = (d_413_v39_).issubset(d_413_v39_)
            nw81_[int(17)] = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
            nw81_[int(18)] = False
            nw81_[int(19)] = d_367_v9_
            nw81_[int(20)] = d_367_v9_
            d_414_v40_ = nw81_
            d_414_v40_ = d_414_v40_
            d_415_v41_: _dafny.Array
            nw82_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_415_v41_ = nw82_
            index65_ = default__.safeIndex(609, (d_415_v41_).length(0))
            (d_415_v41_)[index65_] = d_414_v40_
            index66_ = default__.safeIndex(609, (d_415_v41_).length(0))
            (d_415_v41_)[index66_] = d_414_v40_
            d_416_v42_: _dafny.Map
            d_416_v42_ = _dafny.Map({d_362_v4_: (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]})
            d_417_v43_: _dafny.MultiSet
            d_417_v43_ = _dafny.MultiSet([self.f6])
            d_418_v44_: _dafny.Map
            d_418_v44_ = _dafny.Map({d_416_v42_: d_417_v43_})
            d_419_v45_: D3
            d_419_v45_ = D3_DC9(((d_418_v44_)[(d_416_v42_).set(d_362_v4_, -664)] if ((d_416_v42_).set(d_362_v4_, -664)) in (d_418_v44_) else d_417_v43_), D0_DC2(), len(_dafny.SeqWithoutIsStrInference([d_358_v2_ for d_420_i7_ in range(default__.abs(29))])))
            source7_ = d_419_v45_
            if source7_.is_DC9:
                d_421___mcc_h4_ = source7_.cf9
                d_422___mcc_h5_ = source7_.cf10
                d_423___mcc_h6_ = source7_.cf11
                d_424_cf11_ = d_423___mcc_h6_
                d_425_cf10_ = d_422___mcc_h5_
                d_426_cf9_ = d_421___mcc_h4_
                d_362_v4_ = d_362_v4_
                d_427_v46_: _dafny.Map
                d_427_v46_ = _dafny.Map({d_367_v9_: d_365_v7_})
                d_428_v47_: C0
                nw83_ = C0()
                nw83_.ctor__(len(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_358_v2_ for d_429_i8_ in range(default__.abs(933))]): len(((d_427_v46_)[(d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]] if ((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]) in (d_427_v46_) else _dafny.SeqWithoutIsStrInference([(0) - (self.f6) for d_430_i9_ in range(default__.abs(10))])))})))
                d_428_v47_ = nw83_
                index67_ = default__.safeIndex(241, (d_355_v0_).length(0))
                (d_355_v0_)[index67_] = d_357_v1_
                (globalState).f2 = (d_428_v47_).f8
            elif source7_.is_DC8:
                d_431___mcc_h7_ = source7_.cf8
                d_432_cf8_ = d_431___mcc_h7_
                index68_ = default__.safeIndex(779, (d_362_v4_).length(0))
                (d_362_v4_)[index68_] = (0) - ((self).fm12(True, (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], self.f6, globalState))
                d_367_v9_ = (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]
                d_433_v48_: C0
                nw84_ = C0()
                nw84_.ctor__((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])
                d_433_v48_ = nw84_
                d_434_v49_: _dafny.Map
                d_434_v49_ = _dafny.Map({not (d_367_v9_) or ((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))]): default__.safeDivisionInt((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], (d_365_v7_)[default__.safeIndex((self).f5, len(d_365_v7_))])})
                rhs99_ = (0) - (self.f6)
                rhs100_ = d_434_v49_
                rhs101_ = d_358_v2_
                lhs61_ = globalState
                lhs61_.f2 = rhs99_
                d_434_v49_ = rhs100_
                d_358_v2_ = rhs101_
            elif True:
                d_435___mcc_h8_ = source7_.cf12
                d_436_cf12_ = d_435___mcc_h8_
                d_437_v50_: _dafny.Seq
                d_437_v50_ = _dafny.SeqWithoutIsStrInference([d_367_v9_])
                d_437_v50_ = d_437_v50_
                rhs102_ = d_362_v4_
                rhs103_ = ((0) - ((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])) not in (d_417_v43_)
                d_362_v4_ = rhs102_
                d_367_v9_ = rhs103_
                d_438_v51_: _dafny.Array
                nw85_ = _dafny.Array(None, 23)
                nw85_[int(0)] = d_358_v2_
                nw85_[int(1)] = d_358_v2_
                nw85_[int(2)] = d_358_v2_
                nw85_[int(3)] = d_358_v2_
                nw85_[int(4)] = d_358_v2_
                nw85_[int(5)] = default__.fm18((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], self.f6, (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], globalState)
                nw85_[int(6)] = d_358_v2_
                nw85_[int(7)] = d_358_v2_
                nw85_[int(8)] = d_358_v2_
                nw85_[int(9)] = d_358_v2_
                nw85_[int(10)] = (d_358_v2_ if (d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))] else default__.fm18((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))], self.f6, -788, d_367_v9_, globalState))
                nw85_[int(11)] = d_358_v2_
                nw85_[int(12)] = d_358_v2_
                nw85_[int(13)] = d_358_v2_
                nw85_[int(14)] = d_358_v2_
                nw85_[int(15)] = d_358_v2_
                nw85_[int(16)] = d_358_v2_
                nw85_[int(17)] = d_358_v2_
                nw85_[int(18)] = d_358_v2_
                nw85_[int(19)] = d_358_v2_
                nw85_[int(20)] = d_358_v2_
                nw85_[int(21)] = d_358_v2_
                nw85_[int(22)] = d_358_v2_
                d_438_v51_ = nw85_
                d_438_v51_ = d_438_v51_
                d_439_v52_: D0
                d_439_v52_ = D0_DC3((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))])
                d_440_v53_: _dafny.Array
                nw86_ = _dafny.Array(None, 10)
                nw86_[int(0)] = d_439_v52_
                nw86_[int(1)] = d_439_v52_
                nw86_[int(2)] = d_439_v52_
                nw86_[int(3)] = D0_DC3(d_367_v9_)
                nw86_[int(4)] = d_439_v52_
                nw86_[int(5)] = d_439_v52_
                nw86_[int(6)] = D0_DC3((d_359_v3_)[default__.safeIndex(716, (d_359_v3_).length(0))])
                nw86_[int(7)] = d_439_v52_
                nw86_[int(8)] = d_439_v52_
                nw86_[int(9)] = d_439_v52_
                d_440_v53_ = nw86_
                index69_ = default__.safeIndex(424, (d_440_v53_).length(0))
                (d_440_v53_)[index69_] = d_439_v52_
                index70_ = default__.safeIndex(424, (d_440_v53_).length(0))
                (d_440_v53_)[index70_] = d_439_v52_
        elif True:
            index71_ = default__.safeIndex(241, (d_355_v0_).length(0))
            (d_355_v0_)[index71_] = (d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]
            (globalState).f2 = default__.safeModuloInt((self).fm13(self.f6, _dafny.MultiSet([d_357_v1_, _dafny.SeqWithoutIsStrInference([d_358_v2_ for d_441_i10_ in range(default__.abs(-505))]), d_357_v1_, (d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]]), globalState), 393)
            if d_367_v9_:
                d_442_v54_: _dafny.Seq
                d_442_v54_ = _dafny.SeqWithoutIsStrInference([(d_357_v1_ if d_367_v9_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ltkjnhdi")))])
                d_443_v55_: _dafny.Map
                d_443_v55_ = _dafny.Map({d_367_v9_: d_367_v9_})
                index72_ = default__.safeIndex(779, (d_362_v4_).length(0))
                (d_362_v4_)[index72_] = len((d_442_v54_).set(default__.safeIndex(default__.safeModuloInt(len(d_443_v55_), 835), len(d_442_v54_)), (_dafny.SeqWithoutIsStrInference([d_358_v2_ for d_444_i11_ in range(default__.abs(145))])) + ((d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))])))
                index73_ = default__.safeIndex(3, (d_359_v3_).length(0))
                (d_359_v3_)[index73_] = ((self).f5) != ((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])
                index74_ = default__.safeIndex(3, (d_359_v3_).length(0))
                (d_359_v3_)[index74_] = (self).fm1(globalState)
                d_445_v56_: C0
                nw87_ = C0()
                nw87_.ctor__((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])
                d_445_v56_ = nw87_
                d_446_v57_: _dafny.MultiSet
                d_446_v57_ = _dafny.MultiSet([(d_445_v56_).f8, (self).f5, self.f6, (0) - (len(d_357_v1_)), (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]])
                d_447_v58_: _dafny.Map
                d_447_v58_ = _dafny.Map({d_357_v1_: (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]})
                d_448_v59_: _dafny.Map
                d_448_v59_ = _dafny.Map({self.f6: (d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))]})
                d_449_v60_: _dafny.Map
                d_449_v60_ = _dafny.Map({(d_359_v3_)[default__.safeIndex(3, (d_359_v3_).length(0))]: _dafny.MultiSet(d_365_v7_)})
                d_450_v61_: _dafny.Array
                nw88_ = _dafny.Array(None, 29)
                nw88_[int(0)] = d_446_v57_
                nw88_[int(1)] = _dafny.MultiSet([(self).f5])
                nw88_[int(2)] = d_446_v57_
                nw88_[int(3)] = d_446_v57_
                nw88_[int(4)] = d_446_v57_
                nw88_[int(5)] = (_dafny.MultiSet(d_365_v7_)) - (d_446_v57_)
                nw88_[int(6)] = d_446_v57_
                nw88_[int(7)] = (_dafny.MultiSet(d_365_v7_)) | (d_446_v57_)
                nw88_[int(8)] = d_446_v57_
                nw88_[int(9)] = (default__.fm19((d_359_v3_)[default__.safeIndex(3, (d_359_v3_).length(0))], (d_445_v56_).f8, (d_359_v3_)[default__.safeIndex(3, (d_359_v3_).length(0))], globalState)).set(self.f6, default__.abs(((d_447_v58_)[(d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]] if ((d_355_v0_)[default__.safeIndex(241, (d_355_v0_).length(0))]) in (d_447_v58_) else (self).f5)))
                nw88_[int(10)] = d_446_v57_
                nw88_[int(11)] = (_dafny.MultiSet([(self).f5])).set(((d_448_v59_)[(self).f5] if ((self).f5) in (d_448_v59_) else self.f6), default__.abs(690))
                nw88_[int(12)] = _dafny.MultiSet([(0) - (len(d_357_v1_)), (d_446_v57_).cardinality])
                nw88_[int(13)] = d_446_v57_
                nw88_[int(14)] = d_446_v57_
                nw88_[int(15)] = d_446_v57_
                nw88_[int(16)] = (d_446_v57_) | (((d_449_v60_)[False] if (False) in (d_449_v60_) else d_446_v57_))
                nw88_[int(17)] = d_446_v57_
                nw88_[int(18)] = d_446_v57_
                nw88_[int(19)] = d_446_v57_
                nw88_[int(20)] = d_446_v57_
                nw88_[int(21)] = d_446_v57_
                nw88_[int(22)] = (d_446_v57_) | (d_446_v57_)
                nw88_[int(23)] = d_446_v57_
                nw88_[int(24)] = _dafny.MultiSet([(self).f5])
                nw88_[int(25)] = d_446_v57_
                nw88_[int(26)] = d_446_v57_
                nw88_[int(27)] = (d_446_v57_ if (d_359_v3_)[default__.safeIndex(3, (d_359_v3_).length(0))] else _dafny.MultiSet(d_365_v7_))
                nw88_[int(28)] = _dafny.MultiSet((d_365_v7_) + (_dafny.SeqWithoutIsStrInference([(d_445_v56_).f8])))
                d_450_v61_ = nw88_
                index75_ = default__.safeIndex(293, (d_450_v61_).length(0))
                (d_450_v61_)[index75_] = (d_446_v57_).intersection(d_446_v57_)
                index76_ = default__.safeIndex(293, (d_450_v61_).length(0))
                index77_ = default__.safeIndex(3, (d_359_v3_).length(0))
                rhs104_ = d_446_v57_
                rhs105_ = not (not(d_367_v9_)) or (not (d_367_v9_) or (not(not(d_367_v9_))))
                rhs106_ = _dafny.SeqWithoutIsStrInference([d_357_v1_])
                lhs62_ = d_450_v61_
                lhs63_ = default__.safeIndex(293, (d_450_v61_).length(0))
                lhs64_ = d_359_v3_
                lhs65_ = default__.safeIndex(3, (d_359_v3_).length(0))
                lhs62_[lhs63_] = rhs104_
                lhs64_[lhs65_] = rhs105_
                d_442_v54_ = rhs106_
                d_451_v62_: _dafny.Array
                d_452_v63_: bool
                d_453_v64_: _dafny.MultiSet
                d_454_v65_: int
                out28_: _dafny.Array
                out29_: bool
                out30_: _dafny.MultiSet
                out31_: int
                out28_, out29_, out30_, out31_ = default__.m0(d_358_v2_, globalState)
                d_451_v62_ = out28_
                d_452_v63_ = out29_
                d_453_v64_ = out30_
                d_454_v65_ = out31_
            elif True:
                d_455_v66_: C0
                nw89_ = C0()
                nw89_.ctor__((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))])
                d_455_v66_ = nw89_
                d_456_v67_: D4
                d_456_v67_ = D4_DC14((d_362_v4_)[default__.safeIndex(779, (d_362_v4_).length(0))], False, (self).f5, d_367_v9_, (d_455_v66_).f8)
                (globalState).f2 = (d_456_v67_).cf22
                d_457_v68_: _dafny.Array
                d_458_v69_: bool
                d_459_v70_: _dafny.MultiSet
                d_460_v71_: int
                out32_: _dafny.Array
                out33_: bool
                out34_: _dafny.MultiSet
                out35_: int
                out32_, out33_, out34_, out35_ = default__.m0(_dafny.CodePoint('p'), globalState)
                d_457_v68_ = out32_
                d_458_v69_ = out33_
                d_459_v70_ = out34_
                d_460_v71_ = out35_
                r1 = d_458_v69_
                r1 = d_367_v9_
            r1 = d_367_v9_
            d_367_v9_ = not(d_367_v9_)
        r0 = default__.safeModuloInt((self).f5, (self).f5)
        r1 = d_367_v9_
        return r0, r1

    def m2(self, p0, globalState):
        d_461_v0_: C0
        nw90_ = C0()
        nw90_.ctor__(self.f6)
        d_461_v0_ = nw90_
        d_462_v1_: _dafny.Array
        def lambda19_(d_463_v0_):
            def lambda20_(d_464_i1_):
                return default__.safeModuloInt(d_464_i1_, (d_463_v0_).f8)

            return lambda20_

        init10_ = lambda19_(d_461_v0_)
        nw91_ = _dafny.Array(None, 27)
        for i0_10_ in range(nw91_.length(0)):
            nw91_[i0_10_] = init10_(i0_10_)
        d_462_v1_ = nw91_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_462_v1_).length(0)):
            d_465_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_465_i0_)) and ((d_465_i0_) < ((d_462_v1_).length(0)))):
                (d_462_v1_)[(d_465_i0_)] = (d_465_i0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_466_i2_ in range(default__.abs(-155))])))
        d_467_v2_: bool
        d_467_v2_ = True
        (self).f6 = ((d_461_v0_).f8 if d_467_v2_ else self.f6)
        d_468_v3_: _dafny.Array
        def lambda21_(d_469_v0_):
            def lambda22_(d_470_i4_):
                return ((d_469_v0_).f8) == (self.f6)

            return lambda22_

        init11_ = lambda21_(d_461_v0_)
        nw92_ = _dafny.Array(None, 19)
        for i0_11_ in range(nw92_.length(0)):
            nw92_[i0_11_] = init11_(i0_11_)
        d_468_v3_ = nw92_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_468_v3_).length(0)):
            d_471_i3_: int = guard_loop_3_
            if (True) and (((0) <= (d_471_i3_)) and ((d_471_i3_) < ((d_468_v3_).length(0)))):
                (d_468_v3_)[(d_471_i3_)] = d_467_v2_
        d_472_v4_: _dafny.Array
        def lambda23_(d_473_v2_, d_474_v0_):
            def lambda24_(d_475_i6_):
                return (_dafny.Map({d_473_v2_: (d_474_v0_).f8})) | (_dafny.Map({not(d_473_v2_): (d_474_v0_).f8}))

            return lambda24_

        init12_ = lambda23_(d_467_v2_, d_461_v0_)
        nw93_ = _dafny.Array(None, 8)
        for i0_12_ in range(nw93_.length(0)):
            nw93_[i0_12_] = init12_(i0_12_)
        d_472_v4_ = nw93_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_472_v4_).length(0)):
            d_476_i5_: int = guard_loop_4_
            if (True) and (((0) <= (d_476_i5_)) and ((d_476_i5_) < ((d_472_v4_).length(0)))):
                (d_472_v4_)[(d_476_i5_)] = _dafny.Map({not(False): len((_dafny.Map({_dafny.CodePoint('k'): len(_dafny.SeqWithoutIsStrInference([D1_DC4(_dafny.Set({(0) - ((self).f5)})), D1_DC4(_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bof")))})), len(_dafny.Map({d_467_v2_: d_467_v2_})), (0) - ((d_461_v0_).f8), (d_461_v0_).f8, self.f6}))]))}) if not(d_467_v2_) else _dafny.Map({_dafny.CodePoint('i'): (0) - ((d_461_v0_).f8)})))})
        d_477_v5_: str
        d_477_v5_ = _dafny.CodePoint('m')
        d_478_v6_: _dafny.Set
        d_478_v6_ = _dafny.Set({d_467_v2_, d_467_v2_, True, d_467_v2_, d_467_v2_})
        rhs107_ = d_468_v3_
        rhs108_ = (d_477_v5_ if True else d_477_v5_)
        rhs109_ = (0) - (((self).fm2((d_461_v0_).f8, d_478_v6_, globalState) if d_467_v2_ else (d_461_v0_).f8))
        rhs110_ = (d_461_v0_).f8
        rhs111_ = (d_461_v0_).f8
        lhs66_ = globalState
        lhs67_ = self
        lhs68_ = self
        d_468_v3_ = rhs107_
        d_477_v5_ = rhs108_
        lhs66_.f2 = rhs109_
        lhs67_.f6 = rhs110_
        lhs68_.f6 = rhs111_


class C2(T0, T1):
    def  __init__(self):
        self._f6: int = int(0)
        self._f5: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f5, f6):
        (self)._f5 = f5
        (self).f6 = f6

    def fm0(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([-363, len(_dafny.Map({False: len(_dafny.Set({True, True}))}))])])).isdisjoint((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([self.f6])])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False})), -401]) for d_479_i0_ in range(default__.abs(567))]))))

    def fm1(self, globalState):
        return ((self).f5) in ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_480_i0_ in range(default__.abs(874))])) + (_dafny.SeqWithoutIsStrInference([(self).f5 for d_481_i1_ in range(default__.abs(22))])))

    def fm2(self, p0, p1, globalState):
        return -302

    def fm8(self, p0, p1, p2, p3, globalState):
        return ((0) - ((self).f5)) > ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))))

    def fm9(self, p0, p1, globalState):
        return True

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_482_v0_: bool
        d_482_v0_ = True
        r1 = d_482_v0_
        d_483_v1_: str
        d_483_v1_ = _dafny.CodePoint('q')
        d_484_v2_: _dafny.Array
        d_485_v3_: bool
        d_486_v4_: _dafny.MultiSet
        d_487_v5_: int
        out36_: _dafny.Array
        out37_: bool
        out38_: _dafny.MultiSet
        out39_: int
        out36_, out37_, out38_, out39_ = default__.m0(d_483_v1_, globalState)
        d_484_v2_ = out36_
        d_485_v3_ = out37_
        d_486_v4_ = out38_
        d_487_v5_ = out39_
        index78_ = default__.safeIndex(689, (d_484_v2_).length(0))
        (d_484_v2_)[index78_] = d_487_v5_
        d_488_v6_: D1
        d_488_v6_ = D1_DC5()
        d_489_v7_: _dafny.MultiSet
        d_489_v7_ = _dafny.MultiSet([d_485_v3_])
        index79_ = default__.safeIndex(689, (d_484_v2_).length(0))
        (d_484_v2_)[index79_] = ((d_489_v7_).cardinality if (self).fm9(d_488_v6_, d_482_v0_, globalState) else (self).f5)
        hi4_ = default__.safeModuloInt(247, d_487_v5_)
        for d_490_i0_ in range((282) + ((0) - (d_487_v5_)), hi4_):
            d_491_v8_: _dafny.Seq
            d_491_v8_ = _dafny.SeqWithoutIsStrInference([(self.f6) != (d_490_i0_)])
            if (d_491_v8_)[default__.safeIndex((self.f6) - (d_490_i0_), len(d_491_v8_))]:
                d_492_v9_: _dafny.MultiSet
                d_492_v9_ = _dafny.MultiSet([d_490_i0_])
                d_493_v10_: C0
                nw94_ = C0()
                nw94_.ctor__((d_492_v9_).cardinality)
                d_493_v10_ = nw94_
                d_494_v11_: _dafny.Seq
                d_494_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ij"))
                d_495_v12_: _dafny.Map
                d_495_v12_ = _dafny.Map({d_482_v0_: d_494_v11_})
                d_494_v11_ = ((((d_495_v12_)[d_485_v3_] if (d_485_v3_) in (d_495_v12_) else d_494_v11_)).set(default__.safeIndex((self).f5, len(((d_495_v12_)[d_485_v3_] if (d_485_v3_) in (d_495_v12_) else d_494_v11_))), d_483_v1_)) + (d_494_v11_)
                d_496_v13_: _dafny.Map
                d_496_v13_ = _dafny.Map({(0) - (d_487_v5_): d_482_v0_})
                d_497_v14_: _dafny.MultiSet
                d_497_v14_ = _dafny.MultiSet([d_496_v13_, d_496_v13_, d_496_v13_])
                d_498_v15_: _dafny.Array
                nw95_ = _dafny.Array(None, 8)
                nw95_[int(0)] = not (((d_496_v13_)[len((d_491_v8_).set(default__.safeIndex((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))], len(d_491_v8_)), d_485_v3_))] if (len((d_491_v8_).set(default__.safeIndex((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))], len(d_491_v8_)), d_485_v3_))) in (d_496_v13_) else True)) or (d_485_v3_)
                nw95_[int(1)] = d_485_v3_
                nw95_[int(2)] = (False if d_485_v3_ else d_482_v0_)
                nw95_[int(3)] = d_485_v3_
                nw95_[int(4)] = (self).fm8(d_494_v11_, d_485_v3_, (self).fm0(d_497_v14_, d_485_v3_, d_494_v11_, globalState), d_482_v0_, globalState)
                nw95_[int(5)] = d_485_v3_
                nw95_[int(6)] = False
                nw95_[int(7)] = d_485_v3_
                d_498_v15_ = nw95_
                index80_ = default__.safeIndex(868, (d_498_v15_).length(0))
                (d_498_v15_)[index80_] = d_482_v0_
                index81_ = default__.safeIndex(868, (d_498_v15_).length(0))
                (d_498_v15_)[index81_] = ((d_492_v9_) | (d_492_v9_)) != (d_492_v9_)
                d_499_v16_: _dafny.Set
                d_499_v16_ = _dafny.Set({d_485_v3_, d_482_v0_, d_482_v0_})
                d_500_v17_: C0
                nw96_ = C0()
                nw96_.ctor__((len(d_499_v16_) if d_482_v0_ else d_490_i0_))
                d_500_v17_ = nw96_
                d_485_v3_ = (d_498_v15_)[default__.safeIndex(868, (d_498_v15_).length(0))]
            elif True:
                d_501_v18_: _dafny.Array
                d_502_v19_: bool
                d_503_v20_: _dafny.MultiSet
                d_504_v21_: int
                out40_: _dafny.Array
                out41_: bool
                out42_: _dafny.MultiSet
                out43_: int
                out40_, out41_, out42_, out43_ = default__.m0(d_483_v1_, globalState)
                d_501_v18_ = out40_
                d_502_v19_ = out41_
                d_503_v20_ = out42_
                d_504_v21_ = out43_
                r1 = d_482_v0_
                d_505_v22_: C0
                nw97_ = C0()
                nw97_.ctor__((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))])
                d_505_v22_ = nw97_
                d_506_v23_: _dafny.Seq
                d_506_v23_ = _dafny.SeqWithoutIsStrInference([(d_505_v22_).f8, d_490_i0_])
                d_507_v24_: _dafny.Seq
                d_507_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gygdbkdr"))
                d_508_v25_: _dafny.Set
                d_508_v25_ = _dafny.Set({d_490_i0_})
                d_509_v26_: _dafny.Array
                nw98_ = _dafny.Array(None, 12)
                nw98_[int(0)] = True
                nw98_[int(1)] = ((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))]) <= ((d_506_v23_)[default__.safeIndex(d_490_i0_, len(d_506_v23_))])
                nw98_[int(2)] = (False) and (d_485_v3_)
                nw98_[int(3)] = (self).fm8(d_507_v24_, (d_491_v8_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([163 for d_510_i1_ in range(default__.abs(-830))])), len(d_491_v8_))], d_502_v19_, d_485_v3_, globalState)
                nw98_[int(4)] = (_dafny.Set({(d_489_v7_).cardinality, len(d_491_v8_)})).isdisjoint(d_508_v25_)
                nw98_[int(5)] = d_482_v0_
                nw98_[int(6)] = (self).fm9(d_488_v6_, d_502_v19_, globalState)
                nw98_[int(7)] = (not(d_485_v3_)) == (d_485_v3_)
                nw98_[int(8)] = not(d_482_v0_)
                nw98_[int(9)] = d_482_v0_
                nw98_[int(10)] = d_502_v19_
                nw98_[int(11)] = not(d_502_v19_)
                d_509_v26_ = nw98_
                nw99_ = _dafny.Array(False, 12)
                d_509_v26_ = nw99_
                d_485_v3_ = not(True)
            d_511_v27_: _dafny.MultiSet
            d_511_v27_ = _dafny.MultiSet([d_490_i0_, len(d_491_v8_)])
            d_512_v28_: _dafny.Set
            d_512_v28_ = _dafny.Set({d_511_v27_})
            d_513_v29_: _dafny.Map
            d_513_v29_ = _dafny.Map({d_487_v5_: d_512_v28_})
            d_514_v30_: _dafny.Seq
            d_514_v30_ = _dafny.SeqWithoutIsStrInference([d_511_v27_])
            def iife47_():
                coll29_ = _dafny.Set()
                compr_29_: _dafny.MultiSet
                for compr_29_ in (d_514_v30_).Elements:
                    d_515_v31_: _dafny.MultiSet = compr_29_
                    if (d_515_v31_) in (d_514_v30_):
                        coll29_ = coll29_.union(_dafny.Set([d_515_v31_]))
                return _dafny.Set(coll29_)
            d_513_v29_ = (d_513_v29_).set((self).f5, iife47_()
            )
            d_516_v32_: _dafny.Map
            d_516_v32_ = _dafny.Map({516: d_482_v0_})
            d_517_v33_: _dafny.Seq
            d_517_v33_ = _dafny.SeqWithoutIsStrInference([d_516_v32_])
            d_518_v34_: _dafny.Seq
            d_518_v34_ = _dafny.SeqWithoutIsStrInference([d_517_v33_])
            d_519_v35_: _dafny.Seq
            d_519_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cldrgrfef"))
            d_520_v36_: _dafny.Array
            nw100_ = _dafny.Array(None, 12)
            nw100_[int(0)] = d_485_v3_
            nw100_[int(1)] = (d_482_v0_ if d_485_v3_ else (d_491_v8_)[default__.safeIndex((d_489_v7_).cardinality, len(d_491_v8_))])
            nw100_[int(2)] = d_485_v3_
            nw100_[int(3)] = (self).fm0(_dafny.MultiSet((d_518_v34_)[default__.safeIndex((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))], len(d_518_v34_))]), False, d_519_v35_, globalState)
            nw100_[int(4)] = (False if d_482_v0_ else d_482_v0_)
            nw100_[int(5)] = (d_482_v0_) and (d_485_v3_)
            nw100_[int(6)] = d_482_v0_
            nw100_[int(7)] = d_482_v0_
            nw100_[int(8)] = (d_487_v5_) > ((0) - ((0) - ((self).f5)))
            nw100_[int(9)] = d_485_v3_
            nw100_[int(10)] = d_485_v3_
            nw100_[int(11)] = (d_482_v0_ if d_482_v0_ else d_485_v3_)
            d_520_v36_ = nw100_
            index82_ = default__.safeIndex(486, (d_520_v36_).length(0))
            (d_520_v36_)[index82_] = d_485_v3_
            index83_ = default__.safeIndex(486, (d_520_v36_).length(0))
            (d_520_v36_)[index83_] = ((d_489_v7_).intersection(d_489_v7_)).issubset((_dafny.MultiSet(d_491_v8_)) | (d_489_v7_))
            d_521_v37_: _dafny.Array
            def lambda25_(d_522_v36_, d_523_v0_):
                def lambda26_(d_524_i2_):
                    return _dafny.Map({(d_522_v36_)[default__.safeIndex(486, (d_522_v36_).length(0))]: d_523_v0_})

                return lambda26_

            init13_ = lambda25_(d_520_v36_, d_482_v0_)
            nw101_ = _dafny.Array(None, 29)
            for i0_13_ in range(nw101_.length(0)):
                nw101_[i0_13_] = init13_(i0_13_)
            d_521_v37_ = nw101_
            d_525_v38_: _dafny.Map
            d_525_v38_ = _dafny.Map({(0) - ((self).f5): (d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))]})
            d_526_v39_: _dafny.Map
            d_526_v39_ = _dafny.Map({d_482_v0_: True})
            index84_ = default__.safeIndex(439, (d_521_v37_).length(0))
            (d_521_v37_)[index84_] = (default__.fm10((d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))], d_525_v38_, not(d_485_v3_), (d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))], globalState)) | (d_526_v39_)
            index85_ = default__.safeIndex(439, (d_521_v37_).length(0))
            (d_521_v37_)[index85_] = d_526_v39_
        index86_ = default__.safeIndex(689, (d_484_v2_).length(0))
        (d_484_v2_)[index86_] = self.f6
        d_527_v40_: _dafny.Seq
        d_527_v40_ = _dafny.SeqWithoutIsStrInference([777, (d_484_v2_)[default__.safeIndex(689, (d_484_v2_).length(0))]])
        d_528_v41_: _dafny.Map
        d_528_v41_ = _dafny.Map({d_482_v0_: d_485_v3_})
        d_529_v42_: _dafny.Set
        d_529_v42_ = _dafny.Set({d_482_v0_})
        index87_ = default__.safeIndex(689, (d_484_v2_).length(0))
        (d_484_v2_)[index87_] = (d_487_v5_) - (((d_527_v40_)[default__.safeIndex(self.f6, len(d_527_v40_))]) + ((self).fm2(len(d_528_v41_), d_529_v42_, globalState)))
        r0 = (d_527_v40_)[default__.safeIndex(d_487_v5_, len(d_527_v40_))]
        r1 = d_482_v0_
        return r0, r1

    def m2(self, p0, globalState):
        d_530_v0_: bool
        d_530_v0_ = False
        if d_530_v0_:
            d_531_v1_: _dafny.Seq
            d_531_v1_ = _dafny.SeqWithoutIsStrInference([p0, self.f6])
            d_532_v2_: str
            d_532_v2_ = _dafny.CodePoint('q')
            d_533_v3_: _dafny.Map
            d_533_v3_ = _dafny.Map({d_532_v2_: (0) - ((self).f5)})
            d_534_v4_: _dafny.Seq
            d_534_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwfr"))
            rhs112_ = d_531_v1_
            rhs113_ = (default__.fm11((0) - ((0) - (p0)), (self).f5, self.f6, self.f6, globalState)) == (default__.fm11((self).f5, len((d_533_v3_).set((d_534_v4_)[default__.safeIndex((0) - (p0), len(d_534_v4_))], 996)), p0, len(d_534_v4_), globalState))
            d_531_v1_ = rhs112_
            d_530_v0_ = rhs113_
            d_535_v5_: D1
            d_535_v5_ = D1_DC6(d_534_v4_, self.f6)
            d_536_v6_: D0
            d_536_v6_ = D0_DC0(False)
            d_537_v7_: _dafny.MultiSet
            d_537_v7_ = _dafny.MultiSet([(d_536_v6_).cf0, d_530_v0_])
            d_538_v8_: _dafny.MultiSet
            d_538_v8_ = _dafny.MultiSet([d_537_v7_])
            d_530_v0_ = (self).fm8((d_534_v4_) + ((d_535_v5_).cf5), not (d_530_v0_) or (not(False)), d_530_v0_, (d_538_v8_).isdisjoint(d_538_v8_), globalState)
            d_539_v9_: _dafny.Array
            nw102_ = _dafny.Array(_dafny.MultiSet({}), 3)
            d_539_v9_ = nw102_
            d_540_v10_: _dafny.Set
            d_540_v10_ = _dafny.Set({d_530_v0_, d_530_v0_, False})
            d_541_v11_: _dafny.MultiSet
            d_541_v11_ = _dafny.MultiSet([d_532_v2_, d_532_v2_])
            d_542_v12_: _dafny.MultiSet
            d_542_v12_ = _dafny.MultiSet([(self).f5, (self).fm2((self).f5, d_540_v10_, globalState), ((d_541_v11_)[d_532_v2_] if (d_532_v2_) in (d_541_v11_) else ((d_533_v3_)[d_532_v2_] if (d_532_v2_) in (d_533_v3_) else 660)), p0, (self).f5])
            index88_ = default__.safeIndex(379, (d_539_v9_).length(0))
            (d_539_v9_)[index88_] = d_542_v12_
            index89_ = default__.safeIndex(379, (d_539_v9_).length(0))
            (d_539_v9_)[index89_] = _dafny.MultiSet([p0])
            d_543_v13_: _dafny.Set
            d_543_v13_ = _dafny.Set({p0})
            d_544_v14_: D1
            d_544_v14_ = D1_DC4(d_543_v13_)
            source8_ = d_544_v14_
            if source8_.is_DC5:
                d_545_v15_: _dafny.Array
                nw103_ = _dafny.Array(int(0), 29)
                d_545_v15_ = nw103_
                index90_ = default__.safeIndex(30, (d_545_v15_).length(0))
                (d_545_v15_)[index90_] = (self).fm2(self.f6, d_540_v10_, globalState)
                index91_ = default__.safeIndex(30, (d_545_v15_).length(0))
                rhs114_ = d_530_v0_
                rhs115_ = (d_541_v11_).isdisjoint(d_541_v11_)
                rhs116_ = (d_545_v15_ if d_530_v0_ else d_545_v15_)
                rhs117_ = (0) - ((self).fm2(self.f6, (_dafny.Set({d_530_v0_})) - (d_540_v10_), globalState))
                lhs69_ = d_545_v15_
                lhs70_ = default__.safeIndex(30, (d_545_v15_).length(0))
                d_530_v0_ = rhs114_
                d_530_v0_ = rhs115_
                d_545_v15_ = rhs116_
                lhs69_[lhs70_] = rhs117_
                (self).f6 = (self).f5
                d_546_v16_: _dafny.Array
                nw104_ = _dafny.Array(False, 3)
                d_546_v16_ = nw104_
                index92_ = default__.safeIndex(324, (d_546_v16_).length(0))
                (d_546_v16_)[index92_] = not(d_530_v0_)
                index93_ = default__.safeIndex(324, (d_546_v16_).length(0))
                (d_546_v16_)[index93_] = d_530_v0_
                index94_ = default__.safeIndex(30, (d_545_v15_).length(0))
                (d_545_v15_)[index94_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npeqafbc"))) + (d_534_v4_))
            elif source8_.is_DC6:
                d_547___mcc_h0_ = source8_.cf5
                d_548___mcc_h1_ = source8_.cf6
                d_549_cf6_ = d_548___mcc_h1_
                d_550_cf5_ = d_547___mcc_h0_
                d_551_v17_: _dafny.Array
                nw105_ = _dafny.Array(None, 3)
                nw105_[int(0)] = d_549_cf6_
                nw105_[int(1)] = self.f6
                nw105_[int(2)] = self.f6
                d_551_v17_ = nw105_
                index95_ = default__.safeIndex(739, (d_551_v17_).length(0))
                (d_551_v17_)[index95_] = 728
                index96_ = default__.safeIndex(739, (d_551_v17_).length(0))
                (d_551_v17_)[index96_] = (self).f5
                d_552_v18_: C0
                nw106_ = C0()
                nw106_.ctor__(p0)
                d_552_v18_ = nw106_
                d_553_v19_: _dafny.MultiSet
                d_553_v19_ = _dafny.MultiSet([(d_552_v18_ if d_530_v0_ else d_552_v18_), d_552_v18_])
                d_553_v19_ = (d_553_v19_) | (d_553_v19_)
                (globalState).f2 = (self).f5
            elif True:
                d_554___mcc_h2_ = source8_.cf4
                d_555_cf4_ = d_554___mcc_h2_
                d_556_v20_: D0
                d_556_v20_ = D0_DC2()
                d_557_v21_: _dafny.Array
                nw107_ = _dafny.Array(None, 15)
                nw107_[int(0)] = (d_556_v20_ if d_530_v0_ else d_556_v20_)
                nw107_[int(1)] = D0_DC2()
                nw107_[int(2)] = d_556_v20_
                nw107_[int(3)] = d_556_v20_
                nw107_[int(4)] = d_556_v20_
                nw107_[int(5)] = d_556_v20_
                nw107_[int(6)] = d_556_v20_
                nw107_[int(7)] = d_556_v20_
                nw107_[int(8)] = d_556_v20_
                nw107_[int(9)] = D0_DC2()
                nw107_[int(10)] = d_556_v20_
                nw107_[int(11)] = d_556_v20_
                nw107_[int(12)] = D0_DC2()
                nw107_[int(13)] = D0_DC2()
                nw107_[int(14)] = (d_556_v20_ if d_530_v0_ else d_556_v20_)
                d_557_v21_ = nw107_
                index97_ = default__.safeIndex(433, (d_557_v21_).length(0))
                (d_557_v21_)[index97_] = D0_DC2()
                index98_ = default__.safeIndex(433, (d_557_v21_).length(0))
                (d_557_v21_)[index98_] = d_556_v20_
                d_558_v22_: _dafny.Map
                d_558_v22_ = _dafny.Map({d_534_v4_: (((d_539_v9_)[default__.safeIndex(379, (d_539_v9_).length(0))]).cardinality) * (p0)})
                d_558_v22_ = (d_558_v22_).set((d_534_v4_ if d_530_v0_ else d_534_v4_), (len(d_543_v13_)) * ((self).f5))
                d_559_v23_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 25)
                d_559_v23_ = nw108_
                d_560_v24_: T1
                nw109_ = C1()
                nw109_.ctor__((0) - (p0), p0)
                d_560_v24_ = nw109_
                d_561_v25_: _dafny.Set
                d_561_v25_ = _dafny.Set({d_560_v24_, d_560_v24_, d_560_v24_, d_560_v24_, d_560_v24_})
                d_562_v26_: _dafny.Array
                nw110_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_562_v26_ = nw110_
                index99_ = default__.safeIndex(249, (d_562_v26_).length(0))
                (d_562_v26_)[index99_] = d_559_v23_
                d_563_v27_: _dafny.Map
                d_563_v27_ = _dafny.Map({(self).f5: d_530_v0_})
                d_564_v28_: D5
                d_564_v28_ = D5_DC17(self.f6, d_559_v23_, d_563_v27_)
                index100_ = default__.safeIndex(249, (d_562_v26_).length(0))
                rhs118_ = d_559_v23_
                rhs119_ = d_561_v25_
                rhs120_ = (d_564_v28_).cf28
                rhs121_ = ((p0) < ((self).fm2((self).f5, d_540_v10_, globalState))) and (((self).f5) == (p0))
                rhs122_ = ((d_537_v7_) - (_dafny.MultiSet([not(True)]))).ispropersubset(d_537_v7_)
                lhs71_ = d_562_v26_
                lhs72_ = default__.safeIndex(249, (d_562_v26_).length(0))
                d_559_v23_ = rhs118_
                d_561_v25_ = rhs119_
                lhs71_[lhs72_] = rhs120_
                d_530_v0_ = rhs121_
                d_530_v0_ = rhs122_
                (self).f6 = (d_560_v24_).f5
            d_530_v0_ = (d_534_v4_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gyqisghb"))) + (d_534_v4_))
        elif True:
            d_565_v30_: _dafny.Map
            d_565_v30_ = _dafny.Map({self.f6: d_530_v0_})
            d_566_v31_: C1
            nw111_ = C1()
            nw111_.ctor__((self).f5, (self).f5)
            d_566_v31_ = nw111_
            d_567_v32_: _dafny.Map
            d_567_v32_ = _dafny.Map({d_530_v0_: d_566_v31_})
            d_568_v33_: _dafny.Seq
            d_568_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynpxt"))
            d_569_v34_: _dafny.Map
            d_569_v34_ = _dafny.Map({d_568_v33_: d_567_v32_})
            d_570_v35_: _dafny.Array
            nw112_ = _dafny.Array(None, 28)
            nw112_[int(0)] = p0
            nw112_[int(1)] = p0
            nw112_[int(2)] = self.f6
            nw112_[int(3)] = default__.safeDivisionInt((self).f5, self.f6)
            nw112_[int(4)] = (0) - ((self).f5)
            nw112_[int(5)] = (self).f5
            def iife48_():
                coll30_ = _dafny.Set()
                compr_30_: int
                for compr_30_ in _dafny.IntegerRange(590, 716):
                    d_571_v29_: int = compr_30_
                    if ((590) <= (d_571_v29_)) and ((d_571_v29_) < (716)):
                        coll30_ = coll30_.union(_dafny.Set([(d_571_v29_) * (self.f6)]))
                return _dafny.Set(coll30_)
            nw112_[int(6)] = len(iife48_()
            )
            nw112_[int(7)] = (len(d_565_v30_)) - (775)
            nw112_[int(8)] = self.f6
            nw112_[int(9)] = len((d_567_v32_) | (((d_569_v34_)[d_568_v33_] if (d_568_v33_) in (d_569_v34_) else d_567_v32_)))
            nw112_[int(10)] = p0
            nw112_[int(11)] = self.f6
            nw112_[int(12)] = default__.safeModuloInt((self).f5, (self).f5)
            nw112_[int(13)] = ((self).f5) + (57)
            nw112_[int(14)] = p0
            nw112_[int(15)] = (self).f5
            nw112_[int(16)] = p0
            nw112_[int(17)] = (self.f6 if d_530_v0_ else (self).f5)
            nw112_[int(18)] = self.f6
            nw112_[int(19)] = (self).f5
            nw112_[int(20)] = p0
            nw112_[int(21)] = (self).f5
            nw112_[int(22)] = (0) - ((self.f6 if True else p0))
            nw112_[int(23)] = (self.f6) + (p0)
            nw112_[int(24)] = self.f6
            nw112_[int(25)] = self.f6
            nw112_[int(26)] = (p0) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmkmdiv"))))
            nw112_[int(27)] = self.f6
            d_570_v35_ = nw112_
            d_570_v35_ = (D5_DC17(self.f6, d_570_v35_, d_565_v30_)).cf28
            d_572_v36_: D5
            d_572_v36_ = D5_DC16(d_570_v35_)
            d_573_v37_: _dafny.Seq
            d_573_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_572_v36_, d_572_v36_, d_572_v36_, d_572_v36_, d_572_v36_})])
            rhs123_ = d_530_v0_
            rhs124_ = d_530_v0_
            rhs125_ = d_573_v37_
            rhs126_ = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tomjydsx"))) + ((d_568_v33_) + (d_568_v33_)))
            rhs127_ = not (True) or (d_530_v0_)
            lhs73_ = globalState
            d_530_v0_ = rhs123_
            d_530_v0_ = rhs124_
            d_573_v37_ = rhs125_
            lhs73_.f2 = rhs126_
            d_530_v0_ = rhs127_
            d_574_v38_: _dafny.Map
            d_574_v38_ = _dafny.Map({d_530_v0_: self.f6})
            d_575_v39_: _dafny.Map
            d_575_v39_ = _dafny.Map({(self).f5: d_574_v38_})
            d_575_v39_ = (d_575_v39_).set(p0, (_dafny.Map({d_530_v0_: self.f6}) if not(not(d_530_v0_)) else _dafny.Map({d_530_v0_: self.f6})))
            d_576_v40_: _dafny.Seq
            d_576_v40_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_577_i0_ in range(default__.abs(47))]))])
            d_578_v41_: _dafny.Array
            nw113_ = _dafny.Array(None, 3)
            nw113_[int(0)] = d_530_v0_
            nw113_[int(1)] = (d_576_v40_) != ((d_576_v40_).set(default__.safeIndex(self.f6, len(d_576_v40_)), len(d_576_v40_)))
            nw113_[int(2)] = d_530_v0_
            d_578_v41_ = nw113_
            d_578_v41_ = d_578_v41_
            index101_ = default__.safeIndex(267, (d_570_v35_).length(0))
            (d_570_v35_)[index101_] = self.f6
            index102_ = default__.safeIndex(267, (d_570_v35_).length(0))
            (d_570_v35_)[index102_] = p0
        d_579_v42_: _dafny.Array
        nw114_ = _dafny.Array(_dafny.CodePoint('D'), 14)
        d_579_v42_ = nw114_
        d_580_v43_: str
        d_580_v43_ = _dafny.CodePoint('a')
        index103_ = default__.safeIndex(508, (d_579_v42_).length(0))
        (d_579_v42_)[index103_] = d_580_v43_
        d_581_v44_: _dafny.MultiSet
        d_581_v44_ = _dafny.MultiSet([d_530_v0_, d_530_v0_])
        d_582_v45_: _dafny.MultiSet
        d_582_v45_ = _dafny.MultiSet([d_581_v44_])
        index104_ = default__.safeIndex(508, (d_579_v42_).length(0))
        (d_579_v42_)[index104_] = (d_580_v43_ if (d_582_v45_).isdisjoint(d_582_v45_) else d_580_v43_)
        d_583_v46_: _dafny.Seq
        d_583_v46_ = _dafny.SeqWithoutIsStrInference([d_530_v0_, d_530_v0_])
        if (d_583_v46_)[default__.safeIndex(-612, len(d_583_v46_))]:
            d_584_v47_: _dafny.Array
            nw115_ = _dafny.Array(int(0), 12)
            d_584_v47_ = nw115_
            d_585_v48_: _dafny.Map
            d_585_v48_ = _dafny.Map({d_581_v44_: d_530_v0_})
            d_586_v49_: _dafny.Seq
            d_586_v49_ = _dafny.SeqWithoutIsStrInference([p0, len(d_585_v48_)])
            d_587_v50_: _dafny.Seq
            d_587_v50_ = _dafny.SeqWithoutIsStrInference([(d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))], d_580_v43_])
            d_588_v51_: _dafny.Seq
            d_588_v51_ = _dafny.SeqWithoutIsStrInference([(d_586_v49_)[default__.safeIndex(len(d_587_v50_), len(d_586_v49_))]])
            d_589_v52_: _dafny.Map
            d_589_v52_ = _dafny.Map({len(d_586_v49_): d_530_v0_})
            d_590_v53_: D5
            d_590_v53_ = D5_DC17((self).f5, d_584_v47_, (d_589_v52_) | (d_589_v52_))
            d_590_v53_ = d_590_v53_
            (globalState).f2 = (self).f5
            if (d_530_v0_) == ((self.f6) <= (p0)):
                d_591_v54_: _dafny.Set
                d_591_v54_ = _dafny.Set({d_581_v44_})
                d_592_v55_: _dafny.MultiSet
                d_592_v55_ = _dafny.MultiSet([d_591_v54_])
                d_593_v57_: _dafny.Map
                d_593_v57_ = _dafny.Map({d_580_v43_: d_530_v0_})
                def iife49_():
                    coll31_ = _dafny.Map()
                    compr_31_: str
                    for compr_31_ in (d_593_v57_).keys.Elements:
                        d_594_v56_: str = compr_31_
                        if (d_594_v56_) in (d_593_v57_):
                            coll31_[d_594_v56_] = d_530_v0_
                    return _dafny.Map(coll31_)
                (self).f6 = ((p0) + ((self).f5)) * (((d_592_v55_)[d_591_v54_] if (d_591_v54_) in (d_592_v55_) else len(iife49_()
                )))
                d_595_v58_: _dafny.Map
                d_595_v58_ = _dafny.Map({d_584_v47_: not (d_530_v0_) or (d_530_v0_)})
                d_596_v59_: _dafny.Seq
                d_596_v59_ = _dafny.SeqWithoutIsStrInference([d_595_v58_])
                d_595_v58_ = (d_595_v58_) | ((d_596_v59_)[default__.safeIndex(-172, len(d_596_v59_))])
                index105_ = default__.safeIndex(508, (d_579_v42_).length(0))
                (d_579_v42_)[index105_] = (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]
                d_587_v50_ = d_587_v50_
                d_597_v60_: _dafny.Map
                d_597_v60_ = _dafny.Map({self.f6: (self).f5})
                d_597_v60_ = (d_597_v60_).set(p0, len((d_587_v50_) + (d_587_v50_)))
            elif True:
                (globalState).f2 = (self).f5
                d_598_v61_: _dafny.MultiSet
                d_598_v61_ = _dafny.MultiSet([len(d_587_v50_), self.f6])
                d_599_v62_: D0
                d_599_v62_ = D0_DC2()
                d_600_v63_: D3
                d_600_v63_ = D3_DC9(_dafny.MultiSet([p0]), d_599_v62_, (self).f5)
                d_601_v64_: _dafny.Set
                d_601_v64_ = _dafny.Set({self.f6})
                d_602_v65_: _dafny.MultiSet
                d_602_v65_ = _dafny.MultiSet([(d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))], (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]])
                d_603_v66_: _dafny.Array
                nw116_ = _dafny.Array(None, 20)
                nw116_[int(0)] = d_530_v0_
                nw116_[int(1)] = (d_586_v49_) <= (d_588_v51_)
                nw116_[int(2)] = False
                nw116_[int(3)] = not(False)
                nw116_[int(4)] = (_dafny.MultiSet([(d_600_v63_).cf11])).issubset(d_598_v61_)
                nw116_[int(5)] = (d_601_v64_) != (_dafny.Set({(self).f5}))
                nw116_[int(6)] = d_530_v0_
                nw116_[int(7)] = d_530_v0_
                nw116_[int(8)] = (d_581_v44_).isdisjoint(d_581_v44_)
                nw116_[int(9)] = (d_587_v50_) <= (d_587_v50_)
                nw116_[int(10)] = (d_598_v61_).isdisjoint(_dafny.MultiSet([(self).f5, p0]))
                nw116_[int(11)] = d_530_v0_
                nw116_[int(12)] = (self.f6) > (888)
                nw116_[int(13)] = d_530_v0_
                nw116_[int(14)] = d_530_v0_
                nw116_[int(15)] = not(d_530_v0_)
                nw116_[int(16)] = ((_dafny.MultiSet([_dafny.CodePoint('b'), d_580_v43_])).set(d_580_v43_, default__.abs(p0))).issubset(d_602_v65_)
                nw116_[int(17)] = d_530_v0_
                nw116_[int(18)] = d_530_v0_
                nw116_[int(19)] = d_530_v0_
                d_603_v66_ = nw116_
                index106_ = default__.safeIndex(284, (d_603_v66_).length(0))
                (d_603_v66_)[index106_] = d_530_v0_
                index107_ = default__.safeIndex(284, (d_603_v66_).length(0))
                (d_603_v66_)[index107_] = ((self).f5) >= ((self).f5)
                index108_ = default__.safeIndex(284, (d_603_v66_).length(0))
                (d_603_v66_)[index108_] = d_530_v0_
                d_604_v67_: _dafny.MultiSet
                d_604_v67_ = _dafny.MultiSet([d_584_v47_])
                d_605_v68_: _dafny.Map
                d_605_v68_ = _dafny.Map({d_604_v67_: d_586_v49_})
                d_605_v68_ = (d_605_v68_).set(d_604_v67_, ((_dafny.SeqWithoutIsStrInference([(self).f5])).set(default__.safeIndex(len(d_588_v51_), len(_dafny.SeqWithoutIsStrInference([(self).f5]))), p0)) + (d_586_v49_))
                d_606_v69_: _dafny.Set
                d_606_v69_ = _dafny.Set({d_603_v66_, d_603_v66_})
                d_606_v69_ = d_606_v69_
            d_607_v70_: _dafny.Array
            d_608_v71_: bool
            d_609_v72_: _dafny.MultiSet
            d_610_v73_: int
            out44_: _dafny.Array
            out45_: bool
            out46_: _dafny.MultiSet
            out47_: int
            out44_, out45_, out46_, out47_ = default__.m0(_dafny.CodePoint('l'), globalState)
            d_607_v70_ = out44_
            d_608_v71_ = out45_
            d_609_v72_ = out46_
            d_610_v73_ = out47_
            d_611_v74_: _dafny.Set
            d_611_v74_ = _dafny.Set({(d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]})
            d_611_v74_ = d_611_v74_
        elif True:
            d_612_v75_: _dafny.Map
            d_612_v75_ = _dafny.Map({d_530_v0_: d_530_v0_})
            d_613_v76_: _dafny.Seq
            d_613_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycjiany"))
            d_530_v0_ = ((d_612_v75_)[(len(d_613_v76_)) > (self.f6)] if ((len(d_613_v76_)) > (self.f6)) in (d_612_v75_) else ((0) - (self.f6)) > (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wml")))))
            d_614_v77_: _dafny.Array
            nw117_ = _dafny.Array(_dafny.Seq({}), 8)
            d_614_v77_ = nw117_
            d_615_v78_: _dafny.Set
            d_615_v78_ = _dafny.Set({(self).f5, p0, (self).f5, p0})
            d_616_v79_: _dafny.Map
            d_616_v79_ = _dafny.Map({len(d_615_v78_): 952})
            d_617_v80_: _dafny.Seq
            d_617_v80_ = _dafny.SeqWithoutIsStrInference([p0])
            d_618_v81_: C0
            nw118_ = C0()
            nw118_.ctor__(len(_dafny.Map({d_530_v0_: ((d_616_v79_)[p0] if (p0) in (d_616_v79_) else len(d_617_v80_))})))
            d_618_v81_ = nw118_
            d_619_v82_: _dafny.Seq
            d_619_v82_ = _dafny.SeqWithoutIsStrInference([d_618_v81_, d_618_v81_])
            index109_ = default__.safeIndex(65, (d_614_v77_).length(0))
            (d_614_v77_)[index109_] = d_619_v82_
            index110_ = default__.safeIndex(65, (d_614_v77_).length(0))
            (d_614_v77_)[index110_] = d_619_v82_
            d_620_v83_: C1
            nw119_ = C1()
            nw119_.ctor__(579, 464)
            d_620_v83_ = nw119_
            d_530_v0_ = not(((d_583_v46_)[default__.safeIndex((0) - (len(d_615_v78_)), len(d_583_v46_))] if d_530_v0_ else d_530_v0_))
            d_621_v84_: _dafny.Array
            def lambda27_(d_622_i1_):
                return default__.safeModuloInt(d_622_i1_, (self).f5)

            init14_ = lambda27_
            nw120_ = _dafny.Array(None, 4)
            for i0_14_ in range(nw120_.length(0)):
                nw120_[i0_14_] = init14_(i0_14_)
            d_621_v84_ = nw120_
            index111_ = default__.safeIndex(436, (d_621_v84_).length(0))
            (d_621_v84_)[index111_] = (0) - ((self.f6) * (self.f6))
            index112_ = default__.safeIndex(436, (d_621_v84_).length(0))
            (d_621_v84_)[index112_] = (0) - ((p0) - (self.f6))
        d_623_i2_: int
        d_623_i2_ = 0
        with _dafny.label("3"):
            while not (not(d_530_v0_)) or (d_530_v0_):
                with _dafny.c_label("3"):
                    if (d_623_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_623_i2_ = (d_623_i2_) + (1)
                    (self).m4((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f5 for d_624_i3_ in range(default__.abs(89))]))).ispropersubset(_dafny.MultiSet([(self).f5])), globalState)
                    (self).f6 = 841
                    d_625_v86_: _dafny.Seq
                    def iife50_():
                        coll32_ = _dafny.Set()
                        compr_32_: int
                        for compr_32_ in _dafny.IntegerRange(379, -220):
                            d_626_v85_: int = compr_32_
                            if ((379) <= (d_626_v85_)) and ((d_626_v85_) < (-220)):
                                coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_626_v85_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxvn"))))]))
                        return _dafny.Set(coll32_)
                    d_625_v86_ = _dafny.SeqWithoutIsStrInference([iife50_()
                    ])
                    source9_ = default__.fm20((d_625_v86_)[default__.safeIndex(self.f6, len(d_625_v86_))], not((self.f6) == ((self).f5)), False, globalState)
                    if source9_.is_DC1:
                        d_627___mcc_h3_ = source9_.cf1
                        d_628___mcc_h4_ = source9_.cf2
                        d_629_cf2_ = d_628___mcc_h4_
                        d_630_cf1_ = d_627___mcc_h3_
                        d_631_v87_: _dafny.Seq
                        d_631_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekgniqmts"))
                        index113_ = default__.safeIndex(508, (d_579_v42_).length(0))
                        (d_579_v42_)[index113_] = (d_631_v87_)[default__.safeIndex((self.f6) + (862), len(d_631_v87_))]
                        d_629_cf2_ = d_630_cf1_
                        d_632_v88_: _dafny.Array
                        def lambda28_(d_633_p0_):
                            def lambda29_(d_634_i4_):
                                return (_dafny.MultiSet([729])) | (_dafny.MultiSet([d_633_p0_, (self).f5, 180]))

                            return lambda29_

                        init15_ = lambda28_(p0)
                        nw121_ = _dafny.Array(None, 7)
                        for i0_15_ in range(nw121_.length(0)):
                            nw121_[i0_15_] = init15_(i0_15_)
                        d_632_v88_ = nw121_
                        d_632_v88_ = d_632_v88_
                        d_635_v89_: _dafny.Map
                        d_635_v89_ = _dafny.Map({d_630_cf1_: p0})
                        d_636_v90_: _dafny.Set
                        d_636_v90_ = _dafny.Set({d_630_cf1_})
                        d_637_v91_: _dafny.Seq
                        d_637_v91_ = _dafny.SeqWithoutIsStrInference([d_636_v90_, d_636_v90_, d_636_v90_, d_636_v90_])
                        d_638_v92_: _dafny.Map
                        d_638_v92_ = _dafny.Map({473: self.f6})
                        d_639_v93_: _dafny.MultiSet
                        d_639_v93_ = _dafny.MultiSet([self.f6])
                        d_640_v94_: _dafny.Map
                        d_640_v94_ = _dafny.Map({d_639_v93_: d_530_v0_})
                        d_641_v95_: _dafny.Set
                        d_641_v95_ = _dafny.Set({(d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]})
                        d_642_v96_: _dafny.Seq
                        d_642_v96_ = _dafny.SeqWithoutIsStrInference([self.f6, len(d_631_v87_)])
                        d_643_v98_: _dafny.MultiSet
                        def iife51_():
                            coll33_ = _dafny.Set()
                            compr_33_: str
                            for compr_33_ in (d_641_v95_).Elements:
                                d_644_v97_: str = compr_33_
                                if (d_644_v97_) in (d_641_v95_):
                                    coll33_ = coll33_.union(_dafny.Set([d_644_v97_]))
                            return _dafny.Set(coll33_)
                        d_643_v98_ = _dafny.MultiSet([d_641_v95_, iife51_()
                        ])
                        d_645_v99_: _dafny.Array
                        nw122_ = _dafny.Array(None, 28)
                        nw122_[int(0)] = (0) - ((self).fm2(len(_dafny.Set({d_630_cf1_, d_630_cf1_})), _dafny.Set({d_530_v0_}), globalState))
                        nw122_[int(1)] = (self).f5
                        nw122_[int(2)] = ((self).f5 if d_629_cf2_ else (self).f5)
                        nw122_[int(3)] = len(d_635_v89_)
                        nw122_[int(4)] = self.f6
                        nw122_[int(5)] = self.f6
                        nw122_[int(6)] = (self).f5
                        nw122_[int(7)] = p0
                        nw122_[int(8)] = p0
                        nw122_[int(9)] = (self).f5
                        nw122_[int(10)] = (self).f5
                        nw122_[int(11)] = -309
                        nw122_[int(12)] = self.f6
                        nw122_[int(13)] = (self).fm2(self.f6, d_636_v90_, globalState)
                        nw122_[int(14)] = len(d_631_v87_)
                        nw122_[int(15)] = (len(d_637_v91_) if not(d_629_cf2_) else self.f6)
                        nw122_[int(16)] = (0) - (p0)
                        nw122_[int(17)] = self.f6
                        nw122_[int(18)] = (len(d_638_v92_)) * (len(d_640_v94_))
                        nw122_[int(19)] = len((d_583_v46_) + (d_583_v46_))
                        nw122_[int(20)] = (self).f5
                        nw122_[int(21)] = (p0) * (self.f6)
                        nw122_[int(22)] = (0) - ((self).fm2(len(d_641_v95_), d_636_v90_, globalState))
                        nw122_[int(23)] = (0) - ((self).f5)
                        nw122_[int(24)] = (d_642_v96_)[default__.safeIndex((self).f5, len(d_642_v96_))]
                        nw122_[int(25)] = (self).f5
                        nw122_[int(26)] = self.f6
                        nw122_[int(27)] = default__.safeDivisionInt((self).f5, (0) - (((d_643_v98_)[d_641_v95_] if (d_641_v95_) in (d_643_v98_) else (self).f5)))
                        d_645_v99_ = nw122_
                        d_646_v100_: D4
                        d_646_v100_ = D4_DC13(not(d_530_v0_), len(d_638_v92_), d_629_cf2_)
                        index114_ = default__.safeIndex(508, (d_579_v42_).length(0))
                        rhs128_ = d_645_v99_
                        rhs129_ = (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]
                        rhs130_ = d_645_v99_
                        rhs131_ = len((d_642_v96_) + (d_642_v96_))
                        rhs132_ = (d_646_v100_).cf17
                        lhs74_ = d_579_v42_
                        lhs75_ = default__.safeIndex(508, (d_579_v42_).length(0))
                        lhs76_ = globalState
                        d_645_v99_ = rhs128_
                        lhs74_[lhs75_] = rhs129_
                        d_645_v99_ = rhs130_
                        lhs76_.f2 = rhs131_
                        d_530_v0_ = rhs132_
                    elif source9_.is_DC2:
                        d_647_v101_: _dafny.Set
                        d_647_v101_ = _dafny.Set({d_530_v0_})
                        d_530_v0_ = ((d_647_v101_).issubset(d_647_v101_)) == (d_530_v0_)
                        d_648_v102_: _dafny.Array
                        nw123_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_648_v102_ = nw123_
                        d_649_v103_: _dafny.Array
                        nw124_ = _dafny.Array(False, 20)
                        d_649_v103_ = nw124_
                        d_650_v104_: _dafny.Seq
                        d_650_v104_ = _dafny.SeqWithoutIsStrInference([d_649_v103_])
                        index115_ = default__.safeIndex(133, (d_648_v102_).length(0))
                        (d_648_v102_)[index115_] = (d_650_v104_)[default__.safeIndex(390, len(d_650_v104_))]
                        index116_ = default__.safeIndex(133, (d_648_v102_).length(0))
                        (d_648_v102_)[index116_] = d_649_v103_
                        d_651_v105_: _dafny.MultiSet
                        d_651_v105_ = _dafny.MultiSet([d_647_v101_])
                        d_652_v106_: C1
                        nw125_ = C1()
                        nw125_.ctor__(len(_dafny.SeqWithoutIsStrInference([d_580_v43_ for d_653_i5_ in range(default__.abs(107))])), ((d_651_v105_).set(default__.fm21(p0, p0, d_530_v0_, globalState), default__.abs(self.f6))).cardinality)
                        d_652_v106_ = nw125_
                        (self).f6 = (0) - (p0)
                    elif source9_.is_DC3:
                        d_654___mcc_h5_ = source9_.cf3
                        d_655_cf3_ = d_654___mcc_h5_
                        d_656_v107_: _dafny.Seq
                        d_656_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsid"))
                        d_657_v108_: _dafny.Seq
                        d_657_v108_ = _dafny.SeqWithoutIsStrInference([d_656_v107_])
                        d_658_v109_: _dafny.Map
                        d_658_v109_ = _dafny.Map({self.f6: d_657_v108_})
                        d_658_v109_ = (d_658_v109_).set(p0, d_657_v108_)
                        d_656_v107_ = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_659_i6_ in range(default__.abs(-803))]) if d_655_cf3_ else (((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj")))).set(default__.safeIndex(p0, len((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj"))))), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))])).set(default__.safeIndex(p0, len(((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj")))).set(default__.safeIndex(p0, len((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj"))))), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]))), d_580_v43_))).set(default__.safeIndex((0) - (p0), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_660_i6_ in range(default__.abs(-803))]) if d_655_cf3_ else (((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj")))).set(default__.safeIndex(p0, len((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj"))))), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))])).set(default__.safeIndex(p0, len(((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj")))).set(default__.safeIndex(p0, len((d_656_v107_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geuucj"))))), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]))), d_580_v43_)))), d_580_v43_)
                        d_661_v110_: _dafny.Seq
                        d_661_v110_ = _dafny.SeqWithoutIsStrInference([p0, self.f6])
                        d_662_v111_: C0
                        nw126_ = C0()
                        nw126_.ctor__(len(_dafny.Set({d_661_v110_})))
                        d_662_v111_ = nw126_
                        d_663_v112_: _dafny.Array
                        nw127_ = _dafny.Array(int(0), 8)
                        d_663_v112_ = nw127_
                        d_664_v113_: _dafny.Map
                        d_664_v113_ = _dafny.Map({self.f6: (d_583_v46_)[default__.safeIndex(self.f6, len(d_583_v46_))]})
                        d_665_v114_: D5
                        d_665_v114_ = D5_DC17(len(_dafny.SeqWithoutIsStrInference([self.f6, len(_dafny.SeqWithoutIsStrInference([d_655_cf3_])), 288])), d_663_v112_, d_664_v113_)
                        d_666_v115_: _dafny.MultiSet
                        d_666_v115_ = _dafny.MultiSet([D5_DC17(len(d_583_v46_), d_663_v112_, d_664_v113_), d_665_v114_])
                        d_667_v116_: _dafny.Array
                        nw128_ = _dafny.Array(False, 28)
                        d_667_v116_ = nw128_
                        index117_ = default__.safeIndex(571, (d_667_v116_).length(0))
                        (d_667_v116_)[index117_] = d_655_cf3_
                        index118_ = default__.safeIndex(80, (d_667_v116_).length(0))
                        (d_667_v116_)[index118_] = d_655_cf3_
                        index119_ = default__.safeIndex(571, (d_667_v116_).length(0))
                        index120_ = default__.safeIndex(80, (d_667_v116_).length(0))
                        rhs133_ = d_666_v115_
                        rhs134_ = d_530_v0_
                        rhs135_ = not(d_530_v0_)
                        lhs77_ = d_667_v116_
                        lhs78_ = default__.safeIndex(571, (d_667_v116_).length(0))
                        lhs79_ = d_667_v116_
                        lhs80_ = default__.safeIndex(80, (d_667_v116_).length(0))
                        d_666_v115_ = rhs133_
                        lhs77_[lhs78_] = rhs134_
                        lhs79_[lhs80_] = rhs135_
                    elif True:
                        d_668___mcc_h6_ = source9_.cf0
                        d_669_cf0_ = d_668___mcc_h6_
                        d_670_v117_: _dafny.Set
                        d_670_v117_ = _dafny.Set({p0})
                        d_670_v117_ = ((d_670_v117_).intersection(d_670_v117_)).intersection(d_670_v117_)
                        (globalState).f2 = p0
                        d_671_v118_: C1
                        nw129_ = C1()
                        nw129_.ctor__((self).f5, (self).f5)
                        d_671_v118_ = nw129_
                        d_672_v119_: T0
                        nw130_ = C1()
                        nw130_.ctor__(p0, (0) - (self.f6))
                        d_672_v119_ = nw130_
                        d_673_v120_: T0
                        d_673_v120_ = d_672_v119_
                        d_672_v119_ = (d_673_v120_)
                    if d_530_v0_:
                        d_530_v0_ = d_530_v0_
                        d_674_v121_: _dafny.Seq
                        d_674_v121_ = _dafny.SeqWithoutIsStrInference([self.f6])
                        d_675_v123_: _dafny.Map
                        def iife52_():
                            coll34_ = _dafny.Set()
                            compr_34_: int
                            for compr_34_ in (d_674_v121_).Elements:
                                d_676_v122_: int = compr_34_
                                if (d_676_v122_) in (d_674_v121_):
                                    coll34_ = coll34_.union(_dafny.Set([(d_676_v122_) + (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({561}))) for d_677_i7_ in range(default__.abs(562))])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vudi")))})))]))
                            return _dafny.Set(coll34_)
                        d_675_v123_ = _dafny.Map({default__.fm22(iife52_()
                        , d_530_v0_, globalState): d_530_v0_})
                        d_678_v124_: D7
                        d_678_v124_ = D7_DC21(d_675_v123_)
                        d_679_v125_: _dafny.Set
                        d_679_v125_ = _dafny.Set({d_530_v0_})
                        d_680_v126_: _dafny.Map
                        d_680_v126_ = _dafny.Map({d_530_v0_: (self).fm1(globalState)})
                        d_681_v127_: D3
                        d_681_v127_ = D3_DC8(d_680_v126_)
                        d_682_v129_: _dafny.Seq
                        d_682_v129_ = _dafny.SeqWithoutIsStrInference([d_681_v127_])
                        d_683_v131_: _dafny.Seq
                        d_683_v131_ = _dafny.SeqWithoutIsStrInference([d_675_v123_])
                        d_684_v132_: _dafny.Array
                        nw131_ = _dafny.Array(None, 28)
                        nw131_[int(0)] = (d_675_v123_) | (d_675_v123_)
                        nw131_[int(1)] = (d_678_v124_).cf37
                        nw131_[int(2)] = d_675_v123_
                        nw131_[int(3)] = d_675_v123_
                        nw131_[int(4)] = d_675_v123_
                        nw131_[int(5)] = (d_675_v123_) | ((default__.fm23((self).fm2(p0, d_679_v125_, globalState), globalState)).set(d_681_v127_, True))
                        nw131_[int(6)] = d_675_v123_
                        nw131_[int(7)] = (_dafny.Map({d_681_v127_: d_530_v0_})) | (d_675_v123_)
                        nw131_[int(8)] = _dafny.Map({d_681_v127_: d_530_v0_})
                        nw131_[int(9)] = (d_675_v123_) | (d_675_v123_)
                        def iife53_():
                            coll35_ = _dafny.Map()
                            compr_35_: D3
                            for compr_35_ in (d_682_v129_).Elements:
                                d_685_v128_: D3 = compr_35_
                                if (d_685_v128_) in (d_682_v129_):
                                    coll35_[d_685_v128_] = d_530_v0_
                            return _dafny.Map(coll35_)
                        nw131_[int(10)] = ((iife53_()
                        ).set(d_681_v127_, (self).fm1(globalState))) | (d_675_v123_)
                        def iife54_():
                            coll36_ = _dafny.Map()
                            compr_36_: D3
                            for compr_36_ in (d_675_v123_).keys.Elements:
                                d_686_v130_: D3 = compr_36_
                                if (d_686_v130_) in (d_675_v123_):
                                    coll36_[d_686_v130_] = d_530_v0_
                            return _dafny.Map(coll36_)
                        nw131_[int(11)] = iife54_()
                        
                        nw131_[int(12)] = d_675_v123_
                        nw131_[int(13)] = d_675_v123_
                        nw131_[int(14)] = d_675_v123_
                        nw131_[int(15)] = (d_675_v123_ if d_530_v0_ else d_675_v123_)
                        nw131_[int(16)] = d_675_v123_
                        nw131_[int(17)] = d_675_v123_
                        nw131_[int(18)] = d_675_v123_
                        nw131_[int(19)] = (d_683_v131_)[default__.safeIndex(p0, len(d_683_v131_))]
                        nw131_[int(20)] = (d_675_v123_) | (d_675_v123_)
                        nw131_[int(21)] = d_675_v123_
                        nw131_[int(22)] = (d_675_v123_) | (_dafny.Map({d_681_v127_: d_530_v0_}))
                        nw131_[int(23)] = (d_675_v123_) | (_dafny.Map({d_681_v127_: (d_583_v46_)[default__.safeIndex(self.f6, len(d_583_v46_))]}))
                        nw131_[int(24)] = _dafny.Map({d_681_v127_: not(d_530_v0_)})
                        nw131_[int(25)] = d_675_v123_
                        nw131_[int(26)] = d_675_v123_
                        nw131_[int(27)] = d_675_v123_
                        d_684_v132_ = nw131_
                        index121_ = default__.safeIndex(526, (d_684_v132_).length(0))
                        (d_684_v132_)[index121_] = (default__.fm23(self.f6, globalState)) | (d_675_v123_)
                        d_687_v133_: _dafny.Set
                        d_687_v133_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f5, (0) - (p0), self.f6])})
                        index122_ = default__.safeIndex(526, (d_684_v132_).length(0))
                        rhs136_ = ((d_687_v133_) | (d_687_v133_)).ispropersubset((d_687_v133_) | (d_687_v133_))
                        rhs137_ = d_530_v0_
                        rhs138_ = ((0) - (p0)) >= (p0)
                        rhs139_ = (len(d_583_v46_)) * (((self).f5) * ((self).f5))
                        rhs140_ = (d_675_v123_) | ((default__.fm23((0) - (len(d_583_v46_)), globalState)) | ((d_675_v123_).set(d_681_v127_, d_530_v0_)))
                        lhs81_ = globalState
                        lhs82_ = d_684_v132_
                        lhs83_ = default__.safeIndex(526, (d_684_v132_).length(0))
                        d_530_v0_ = rhs136_
                        d_530_v0_ = rhs137_
                        d_530_v0_ = rhs138_
                        lhs81_.f2 = rhs139_
                        lhs82_[lhs83_] = rhs140_
                        d_688_v134_: _dafny.Array
                        nw132_ = _dafny.Array(None, 10)
                        nw132_[int(0)] = d_530_v0_
                        nw132_[int(1)] = d_530_v0_
                        nw132_[int(2)] = d_530_v0_
                        nw132_[int(3)] = d_530_v0_
                        nw132_[int(4)] = d_530_v0_
                        nw132_[int(5)] = (p0) != (self.f6)
                        nw132_[int(6)] = d_530_v0_
                        nw132_[int(7)] = d_530_v0_
                        nw132_[int(8)] = d_530_v0_
                        nw132_[int(9)] = d_530_v0_
                        d_688_v134_ = nw132_
                        d_688_v134_ = d_688_v134_
                        d_530_v0_ = not(d_530_v0_)
                        d_689_v135_: _dafny.Array
                        nw133_ = _dafny.Array(_dafny.Seq({}), 22)
                        d_689_v135_ = nw133_
                        index123_ = default__.safeIndex(431, (d_689_v135_).length(0))
                        (d_689_v135_)[index123_] = (d_674_v121_ if d_530_v0_ else d_674_v121_)
                        index124_ = default__.safeIndex(431, (d_689_v135_).length(0))
                        (d_689_v135_)[index124_] = d_674_v121_
                    elif True:
                        d_690_v136_: _dafny.Array
                        d_691_v137_: bool
                        d_692_v138_: _dafny.MultiSet
                        d_693_v139_: int
                        out48_: _dafny.Array
                        out49_: bool
                        out50_: _dafny.MultiSet
                        out51_: int
                        out48_, out49_, out50_, out51_ = default__.m0(default__.fm18(d_530_v0_, self.f6, self.f6, True, globalState), globalState)
                        d_690_v136_ = out48_
                        d_691_v137_ = out49_
                        d_692_v138_ = out50_
                        d_693_v139_ = out51_
                        d_694_v140_: _dafny.Set
                        d_694_v140_ = _dafny.Set({not(d_530_v0_)})
                        d_693_v139_ = (self).fm2((self).fm2(self.f6, d_694_v140_, globalState), d_694_v140_, globalState)
                        d_695_v141_: _dafny.Map
                        d_695_v141_ = _dafny.Map({d_530_v0_: d_690_v136_})
                        d_696_v142_: _dafny.Set
                        d_696_v142_ = _dafny.Set({616, d_693_v139_})
                        d_697_v143_: C1
                        nw134_ = C1()
                        nw134_.ctor__(default__.safeDivisionInt(self.f6, len(d_695_v141_)), len(d_696_v142_))
                        d_697_v143_ = nw134_
                        rhs141_ = d_580_v43_
                        rhs142_ = d_697_v143_
                        rhs143_ = d_691_v137_
                        rhs144_ = default__.safeDivisionInt((0) - (d_693_v139_), d_693_v139_)
                        lhs84_ = globalState
                        d_580_v43_ = rhs141_
                        d_697_v143_ = rhs142_
                        d_530_v0_ = rhs143_
                        lhs84_.f2 = rhs144_
                        d_698_v144_: _dafny.Seq
                        d_698_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnmpaxw"))
                        d_699_v145_: C1
                        nw135_ = C1()
                        nw135_.ctor__((0) - (d_693_v139_), (0) - (((self).f5) - (len(d_698_v144_))))
                        d_699_v145_ = nw135_
                        d_700_v146_: D0
                        d_700_v146_ = D0_DC3(d_530_v0_)
                        rhs145_ = self.f6
                        rhs146_ = d_700_v146_
                        rhs147_ = d_691_v137_
                        rhs148_ = (d_697_v143_).fm2((self.f6) - (922), d_694_v140_, globalState)
                        rhs149_ = (0) - ((len((d_698_v144_) + ((d_698_v144_).set(default__.safeIndex((_dafny.MultiSet([d_691_v137_])).cardinality, len(d_698_v144_)), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))])))) * (p0))
                        lhs85_ = globalState
                        lhs86_ = globalState
                        lhs85_.f2 = rhs145_
                        d_700_v146_ = rhs146_
                        d_691_v137_ = rhs147_
                        d_693_v139_ = rhs148_
                        lhs86_.f2 = rhs149_
                    pass
            pass
        d_701_v147_: _dafny.Array
        def lambda30_(d_702_v0_):
            def lambda31_(d_703_i8_):
                return (d_703_i8_) * (len(_dafny.Set({d_702_v0_, d_702_v0_, not(d_702_v0_)})))

            return lambda31_

        init16_ = lambda30_(d_530_v0_)
        nw136_ = _dafny.Array(None, 2)
        for i0_16_ in range(nw136_.length(0)):
            nw136_[i0_16_] = init16_(i0_16_)
        d_701_v147_ = nw136_
        d_704_v148_: _dafny.Map
        d_704_v148_ = _dafny.Map({d_701_v147_: self.f6})
        d_705_v149_: _dafny.Map
        d_705_v149_ = _dafny.Map({(self).f5: d_704_v148_})
        d_705_v149_ = (d_705_v149_).set(((d_581_v44_)[d_530_v0_] if (d_530_v0_) in (d_581_v44_) else (self).f5), d_704_v148_)
        if True:
            d_580_v43_ = d_580_v43_
            d_706_v150_: _dafny.Set
            d_706_v150_ = _dafny.Set({d_701_v147_, d_701_v147_})
            d_706_v150_ = d_706_v150_
            d_707_v151_: _dafny.Seq
            d_707_v151_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            index125_ = default__.safeIndex(729, (d_701_v147_).length(0))
            (d_701_v147_)[index125_] = len(d_707_v151_)
            index126_ = default__.safeIndex(729, (d_701_v147_).length(0))
            (d_701_v147_)[index126_] = self.f6
            d_708_v152_: D0
            d_708_v152_ = D0_DC2()
            source10_ = d_708_v152_
            if source10_.is_DC1:
                d_709___mcc_h7_ = source10_.cf1
                d_710___mcc_h8_ = source10_.cf2
                d_711_cf2_ = d_710___mcc_h8_
                d_712_cf1_ = d_709___mcc_h7_
                d_713_v153_: _dafny.Array
                nw137_ = _dafny.Array(False, 3)
                d_713_v153_ = nw137_
                d_714_v154_: _dafny.Map
                d_714_v154_ = _dafny.Map({d_713_v153_: (self).f5})
                d_714_v154_ = (d_714_v154_).set(d_713_v153_, self.f6)
                (self).f6 = default__.safeDivisionInt(234, (self).f5)
                index127_ = default__.safeIndex(729, (d_701_v147_).length(0))
                (d_701_v147_)[index127_] = 236
                index128_ = default__.safeIndex(971, (d_713_v153_).length(0))
                (d_713_v153_)[index128_] = d_530_v0_
                index129_ = default__.safeIndex(971, (d_713_v153_).length(0))
                (d_713_v153_)[index129_] = d_712_cf1_
            elif source10_.is_DC2:
                d_530_v0_ = (d_530_v0_) or (d_530_v0_)
                d_715_v155_: _dafny.Map
                d_715_v155_ = _dafny.Map({True: d_530_v0_})
                index130_ = default__.safeIndex(729, (d_701_v147_).length(0))
                (d_701_v147_)[index130_] = (len((d_715_v155_) | (default__.fm10((self).f5, _dafny.Map({self.f6: p0}), d_530_v0_, p0, globalState)))) + ((p0) - (p0))
                d_716_v156_: _dafny.Array
                nw138_ = _dafny.Array(D7.default()(), 6)
                d_716_v156_ = nw138_
                d_717_v157_: D3
                d_717_v157_ = D3_DC8(d_715_v155_)
                d_718_v158_: _dafny.Map
                d_718_v158_ = _dafny.Map({d_717_v157_: d_530_v0_})
                d_719_v159_: _dafny.Map
                d_719_v159_ = _dafny.Map({D7_DC21(d_718_v158_): d_716_v156_})
                d_720_v160_: _dafny.Seq
                d_720_v160_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_530_v0_: d_530_v0_})])
                index131_ = default__.safeIndex(729, (d_701_v147_).length(0))
                rhs150_ = default__.safeModuloInt(default__.safeModuloInt((d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))], (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]), (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))])
                rhs151_ = ((d_719_v159_)[default__.fm24(d_720_v160_, -597, d_530_v0_, globalState)] if (default__.fm24(d_720_v160_, -597, d_530_v0_, globalState)) in (d_719_v159_) else d_716_v156_)
                lhs87_ = d_701_v147_
                lhs88_ = default__.safeIndex(729, (d_701_v147_).length(0))
                lhs87_[lhs88_] = rhs150_
                d_716_v156_ = rhs151_
                d_721_v161_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_721_v161_ = nw139_
                d_722_v162_: _dafny.Array
                nw140_ = _dafny.Array(None, 17)
                nw140_[int(0)] = d_530_v0_
                nw140_[int(1)] = d_530_v0_
                nw140_[int(2)] = d_530_v0_
                nw140_[int(3)] = d_530_v0_
                nw140_[int(4)] = d_530_v0_
                nw140_[int(5)] = False
                nw140_[int(6)] = not(False)
                nw140_[int(7)] = d_530_v0_
                nw140_[int(8)] = d_530_v0_
                nw140_[int(9)] = True
                nw140_[int(10)] = d_530_v0_
                nw140_[int(11)] = d_530_v0_
                nw140_[int(12)] = not(d_530_v0_)
                nw140_[int(13)] = d_530_v0_
                nw140_[int(14)] = False
                nw140_[int(15)] = False
                nw140_[int(16)] = False
                d_722_v162_ = nw140_
                index132_ = default__.safeIndex(967, (d_721_v161_).length(0))
                (d_721_v161_)[index132_] = d_722_v162_
                index133_ = default__.safeIndex(967, (d_721_v161_).length(0))
                rhs152_ = ((d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]) < ((0) - ((d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]))
                rhs153_ = d_722_v162_
                rhs154_ = _dafny.SeqWithoutIsStrInference([d_580_v43_ for d_723_i9_ in range(default__.abs(362))])
                lhs89_ = d_721_v161_
                lhs90_ = default__.safeIndex(967, (d_721_v161_).length(0))
                d_530_v0_ = rhs152_
                lhs89_[lhs90_] = rhs153_
                d_707_v151_ = rhs154_
            elif source10_.is_DC3:
                d_724___mcc_h9_ = source10_.cf3
                d_725_cf3_ = d_724___mcc_h9_
                d_726_v163_: _dafny.Set
                d_726_v163_ = _dafny.Set({_dafny.CodePoint('h'), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))], (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))], (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))], (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]})
                d_727_v164_: _dafny.MultiSet
                d_727_v164_ = _dafny.MultiSet([len((d_707_v151_).set(default__.safeIndex(210, len(d_707_v151_)), (d_579_v42_)[default__.safeIndex(508, (d_579_v42_).length(0))]))])
                rhs155_ = d_726_v163_
                rhs156_ = default__.safeDivisionInt(((d_727_v164_).cardinality) - (self.f6), (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))])
                lhs91_ = globalState
                d_726_v163_ = rhs155_
                lhs91_.f2 = rhs156_
                d_728_v165_: _dafny.Array
                def lambda32_(d_729_i10_):
                    return (d_729_i10_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsc"))))

                init17_ = lambda32_
                nw141_ = _dafny.Array(None, 28)
                for i0_17_ in range(nw141_.length(0)):
                    nw141_[i0_17_] = init17_(i0_17_)
                d_728_v165_ = nw141_
                d_730_v166_: _dafny.Seq
                d_730_v166_ = _dafny.SeqWithoutIsStrInference([d_728_v165_, d_728_v165_, d_728_v165_])
                d_728_v165_ = (d_730_v166_)[default__.safeIndex(self.f6, len(d_730_v166_))]
                d_731_v167_: D0
                d_731_v167_ = D0_DC1(d_725_cf3_, d_530_v0_)
                d_732_v168_: _dafny.Seq
                d_732_v168_ = _dafny.SeqWithoutIsStrInference([d_727_v164_])
                d_733_v169_: _dafny.Set
                d_733_v169_ = _dafny.Set({not(d_530_v0_)})
                d_734_v170_: _dafny.Map
                d_734_v170_ = _dafny.Map({d_530_v0_: (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]})
                d_735_v171_: _dafny.Map
                d_735_v171_ = _dafny.Map({(self).fm2(p0, d_733_v169_, globalState): (d_734_v170_).set(d_725_cf3_, p0)})
                d_736_v172_: _dafny.Array
                nw142_ = _dafny.Array(None, 8)
                nw142_[int(0)] = D0_DC1(False, d_530_v0_)
                nw142_[int(1)] = D0_DC1(d_530_v0_, d_530_v0_)
                nw142_[int(2)] = D0_DC1((self).fm8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yotfvv")), d_725_cf3_, d_530_v0_, d_530_v0_, globalState), d_530_v0_)
                nw142_[int(3)] = d_731_v167_
                nw142_[int(4)] = D0_DC1(d_725_cf3_, d_725_cf3_)
                nw142_[int(5)] = default__.fm25(d_725_cf3_, (self).f5, d_732_v168_, d_735_v171_, globalState)
                nw142_[int(6)] = d_731_v167_
                nw142_[int(7)] = d_731_v167_
                d_736_v172_ = nw142_
                index134_ = default__.safeIndex(51, (d_736_v172_).length(0))
                (d_736_v172_)[index134_] = default__.fm25(d_725_cf3_, self.f6, d_732_v168_, d_735_v171_, globalState)
                index135_ = default__.safeIndex(51, (d_736_v172_).length(0))
                (d_736_v172_)[index135_] = default__.fm25((d_725_cf3_ if d_530_v0_ else True), self.f6, (d_732_v168_) + (d_732_v168_), d_735_v171_, globalState)
                index136_ = default__.safeIndex(729, (d_701_v147_).length(0))
                (d_701_v147_)[index136_] = (self).f5
            elif True:
                d_737___mcc_h10_ = source10_.cf0
                d_738_cf0_ = d_737___mcc_h10_
                d_739_v173_: D0
                d_739_v173_ = D0_DC3(d_530_v0_)
                d_740_v174_: _dafny.MultiSet
                d_740_v174_ = _dafny.MultiSet([d_580_v43_, _dafny.CodePoint('s'), d_580_v43_])
                d_741_v175_: _dafny.Map
                d_741_v175_ = _dafny.Map({(d_739_v173_).cf3: (d_740_v174_).cardinality})
                d_742_v176_: _dafny.Seq
                d_742_v176_ = _dafny.SeqWithoutIsStrInference([d_741_v175_])
                d_743_v177_: _dafny.Map
                d_743_v177_ = _dafny.Map({len(d_742_v176_): d_530_v0_})
                d_744_v178_: _dafny.Map
                d_744_v178_ = _dafny.Map({len(d_743_v177_): d_738_cf0_})
                d_745_v179_: C0
                nw143_ = C0()
                nw143_.ctor__(len(d_743_v177_))
                d_745_v179_ = nw143_
                d_746_v182_: D1
                d_746_v182_ = D1_DC5()
                d_747_v183_: _dafny.Map
                d_747_v183_ = _dafny.Map({d_741_v175_: (self).fm9(d_746_v182_, d_530_v0_, globalState)})
                d_748_v184_: _dafny.Map
                d_748_v184_ = _dafny.Map({len(d_583_v46_): d_741_v175_})
                d_749_v185_: _dafny.Map
                d_749_v185_ = _dafny.Map({_dafny.Map({d_738_cf0_: 352}): p0})
                d_750_v186_: _dafny.Map
                d_750_v186_ = d_749_v185_
                d_751_v190_: _dafny.Map
                d_751_v190_ = _dafny.Map({d_741_v175_: d_707_v151_})
                d_752_v191_: _dafny.Array
                nw144_ = _dafny.Array(None, 14)
                def iife55_():
                    def iife57_():
                        coll39_ = _dafny.Map()
                        compr_39_: _dafny.Map
                        for compr_39_ in (d_747_v183_).keys.Elements:
                            d_753_v181_: _dafny.Map = compr_39_
                            if (d_753_v181_) in (d_747_v183_):
                                coll39_[d_753_v181_] = (d_581_v44_).cardinality
                        return _dafny.Map(coll39_)
                    coll37_ = _dafny.Map()
                    def iife56_():
                        coll38_ = _dafny.Map()
                        compr_38_: _dafny.Map
                        for compr_38_ in (d_747_v183_).keys.Elements:
                            d_753_v181_: _dafny.Map = compr_38_
                            if (d_753_v181_) in (d_747_v183_):
                                coll38_[d_753_v181_] = (d_581_v44_).cardinality
                        return _dafny.Map(coll38_)
                    compr_37_: _dafny.Map
                    for compr_37_ in (iife56_()
                    ).keys.Elements:
                        d_754_v180_: _dafny.Map = compr_37_
                        if (d_754_v180_) in (iife57_()
                        ):
                            coll37_[d_754_v180_] = (d_745_v179_).f8
                    return _dafny.Map(coll37_)
                nw144_[int(0)] = iife55_()
                
                nw144_[int(1)] = _dafny.Map({((d_748_v184_)[(d_745_v179_).f8] if ((d_745_v179_).f8) in (d_748_v184_) else d_741_v175_): len(_dafny.Map({self.f6: d_530_v0_}))})
                nw144_[int(2)] = d_749_v185_
                nw144_[int(3)] = (d_750_v186_)
                nw144_[int(4)] = _dafny.Map({d_741_v175_: p0})
                def iife58_():
                    coll40_ = _dafny.Map()
                    compr_40_: _dafny.Map
                    for compr_40_ in (d_747_v183_).keys.Elements:
                        d_755_v187_: _dafny.Map = compr_40_
                        if (d_755_v187_) in (d_747_v183_):
                            def iife59_():
                                coll41_ = _dafny.Map()
                                compr_41_: D3
                                for compr_41_ in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({d_530_v0_: d_738_cf0_}))])).Elements:
                                    d_756_v188_: D3 = compr_41_
                                    if (d_756_v188_) in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({d_530_v0_: d_738_cf0_}))])):
                                        coll41_[d_756_v188_] = (d_745_v179_).f8
                                return _dafny.Map(coll41_)
                            coll40_[d_755_v187_] = len(_dafny.SeqWithoutIsStrInference([iife59_()
                            ]))
                    return _dafny.Map(coll40_)
                nw144_[int(5)] = (iife58_()
                ) | (d_749_v185_)
                nw144_[int(6)] = d_749_v185_
                nw144_[int(7)] = d_749_v185_
                nw144_[int(8)] = d_749_v185_
                nw144_[int(9)] = (d_749_v185_) | (_dafny.Map({d_741_v175_: (d_745_v179_).f8}))
                def iife60_():
                    coll42_ = _dafny.Map()
                    compr_42_: _dafny.Map
                    for compr_42_ in (d_751_v190_).keys.Elements:
                        d_757_v189_: _dafny.Map = compr_42_
                        if (d_757_v189_) in (d_751_v190_):
                            coll42_[d_757_v189_] = (self).f5
                    return _dafny.Map(coll42_)
                nw144_[int(10)] = iife60_()
                
                nw144_[int(11)] = (d_749_v185_).set((d_741_v175_).set(d_530_v0_, p0), (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))])
                nw144_[int(12)] = d_749_v185_
                nw144_[int(13)] = d_749_v185_
                d_752_v191_ = nw144_
                index137_ = default__.safeIndex(945, (d_752_v191_).length(0))
                (d_752_v191_)[index137_] = d_749_v185_
                d_758_v192_: D4
                d_758_v192_ = D4_DC15(default__.fm27(437, globalState))
                index138_ = default__.safeIndex(945, (d_752_v191_).length(0))
                rhs157_ = d_745_v179_
                rhs158_ = default__.fm26(d_530_v0_, (d_707_v151_)[default__.safeIndex((d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))], len(d_707_v151_))], d_758_v192_, globalState)
                lhs92_ = d_752_v191_
                lhs93_ = default__.safeIndex(945, (d_752_v191_).length(0))
                d_745_v179_ = rhs157_
                lhs92_[lhs93_] = rhs158_
                d_759_v193_: C0
                nw145_ = C0()
                nw145_.ctor__(p0)
                d_759_v193_ = nw145_
                d_530_v0_ = d_530_v0_
                d_760_v194_: _dafny.Set
                d_760_v194_ = _dafny.Set({d_738_cf0_})
                d_761_v195_: _dafny.Seq
                d_761_v195_ = _dafny.SeqWithoutIsStrInference([p0])
                d_762_v196_: _dafny.Set
                d_762_v196_ = _dafny.Set({len(d_761_v195_), (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]})
                d_763_v197_: _dafny.Seq
                d_763_v197_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_762_v196_])): (d_745_v179_).f8})), -505, (0) - ((d_759_v193_).f8), (self).f5])
                d_764_v198_: _dafny.Seq
                d_764_v198_ = _dafny.SeqWithoutIsStrInference([d_743_v177_, d_743_v177_, d_744_v178_, (d_744_v178_).set(len(d_761_v195_), d_738_cf0_)])
                d_765_v199_: D5
                d_765_v199_ = D5_DC18(d_738_cf0_, p0, False)
                pat_let_tv7_ = d_738_cf0_
                d_766_v200_: _dafny.Map
                def iife61_(_pat_let9_0):
                    def iife62_(d_767_dt__update__tmp_h0_):
                        def iife63_(_pat_let10_0):
                            def iife64_(d_768_dt__update_hcf30_h0_):
                                return D5_DC18(d_768_dt__update_hcf30_h0_, (d_767_dt__update__tmp_h0_).cf31, (d_767_dt__update__tmp_h0_).cf32)
                            return iife64_(_pat_let10_0)
                        return iife63_(pat_let_tv7_)
                    return iife62_(_pat_let9_0)
                d_766_v200_ = _dafny.Map({iife61_(d_765_v199_): (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]})
                d_769_v201_: _dafny.Array
                nw146_ = _dafny.Array(False, 27)
                d_769_v201_ = nw146_
                d_770_v202_: _dafny.Array
                nw147_ = _dafny.Array(None, 23)
                nw147_[int(0)] = (self).f5
                nw147_[int(1)] = self.f6
                nw147_[int(2)] = (self).fm2(p0, d_760_v194_, globalState)
                nw147_[int(3)] = (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]
                nw147_[int(4)] = 195
                nw147_[int(5)] = (self).f5
                nw147_[int(6)] = self.f6
                nw147_[int(7)] = (self).f5
                nw147_[int(8)] = len(d_764_v198_)
                nw147_[int(9)] = len(d_707_v151_)
                nw147_[int(10)] = (d_759_v193_).f8
                nw147_[int(11)] = (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]
                nw147_[int(12)] = self.f6
                nw147_[int(13)] = self.f6
                nw147_[int(14)] = len((d_766_v200_).set(d_765_v199_, self.f6))
                nw147_[int(15)] = (self).f5
                nw147_[int(16)] = (self).f5
                nw147_[int(17)] = len(_dafny.Map({d_580_v43_: d_769_v201_}))
                nw147_[int(18)] = len(d_744_v178_)
                nw147_[int(19)] = (d_701_v147_)[default__.safeIndex(729, (d_701_v147_).length(0))]
                nw147_[int(20)] = len(d_763_v197_)
                nw147_[int(21)] = (0) - ((d_759_v193_).f8)
                nw147_[int(22)] = 79
                d_770_v202_ = nw147_
                d_771_v203_: D5
                d_771_v203_ = D5_DC16(d_770_v202_)
                pat_let_tv8_ = d_770_v202_
                def iife65_(_pat_let11_0):
                    def iife66_(d_772_dt__update__tmp_h1_):
                        def iife67_(_pat_let12_0):
                            def iife68_(d_773_dt__update_hcf26_h0_):
                                return D5_DC16(d_773_dt__update_hcf26_h0_)
                            return iife68_(_pat_let12_0)
                        return iife67_(pat_let_tv8_)
                    return iife66_(_pat_let11_0)
                d_771_v203_ = iife65_(d_771_v203_)
            d_774_v204_: _dafny.Array
            nw148_ = _dafny.Array(_dafny.Map({}), 18)
            d_774_v204_ = nw148_
            d_774_v204_ = d_774_v204_
        elif True:
            d_530_v0_ = ((d_581_v44_) | (d_581_v44_)).issubset(d_581_v44_)
            d_775_v205_: _dafny.Map
            d_775_v205_ = _dafny.Map({p0: self.f6})
            d_775_v205_ = d_775_v205_
            if True:
                (globalState).f2 = self.f6
                d_776_v206_: C0
                nw149_ = C0()
                nw149_.ctor__(self.f6)
                d_776_v206_ = nw149_
                d_777_v207_: _dafny.Set
                d_777_v207_ = _dafny.Set({d_776_v206_})
                d_530_v0_ = not((self.f6) >= ((len(d_583_v46_)) * (len(d_777_v207_))))
                d_778_v208_: _dafny.Set
                d_778_v208_ = _dafny.Set({d_530_v0_})
                d_779_v209_: _dafny.Array
                nw150_ = _dafny.Array(None, 25)
                nw150_[int(0)] = d_583_v46_
                nw150_[int(1)] = d_583_v46_
                nw150_[int(2)] = d_583_v46_
                nw150_[int(3)] = ((d_583_v46_).set(default__.safeIndex(-17, len(d_583_v46_)), d_530_v0_) if d_530_v0_ else d_583_v46_)
                nw150_[int(4)] = (d_583_v46_) + (_dafny.SeqWithoutIsStrInference([d_530_v0_, d_530_v0_]))
                nw150_[int(5)] = d_583_v46_
                nw150_[int(6)] = d_583_v46_
                nw150_[int(7)] = d_583_v46_
                nw150_[int(8)] = d_583_v46_
                nw150_[int(9)] = d_583_v46_
                nw150_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_530_v0_])) + (d_583_v46_)
                nw150_[int(11)] = d_583_v46_
                nw150_[int(12)] = d_583_v46_
                nw150_[int(13)] = (d_583_v46_).set(default__.safeIndex(p0, len(d_583_v46_)), d_530_v0_)
                nw150_[int(14)] = d_583_v46_
                nw150_[int(15)] = (d_583_v46_) + (d_583_v46_)
                nw150_[int(16)] = (d_583_v46_) + (d_583_v46_)
                nw150_[int(17)] = (d_583_v46_).set(default__.safeIndex(len(d_778_v208_), len(d_583_v46_)), False)
                nw150_[int(18)] = (d_583_v46_).set(default__.safeIndex(p0, len(d_583_v46_)), d_530_v0_)
                nw150_[int(19)] = d_583_v46_
                nw150_[int(20)] = _dafny.SeqWithoutIsStrInference([d_530_v0_])
                nw150_[int(21)] = _dafny.SeqWithoutIsStrInference([d_530_v0_, d_530_v0_, d_530_v0_, d_530_v0_, True])
                nw150_[int(22)] = d_583_v46_
                nw150_[int(23)] = _dafny.SeqWithoutIsStrInference([False, d_530_v0_, d_530_v0_])
                nw150_[int(24)] = (_dafny.SeqWithoutIsStrInference([(d_583_v46_)[default__.safeIndex((d_776_v206_).f8, len(d_583_v46_))]])) + (_dafny.SeqWithoutIsStrInference([d_530_v0_]))
                d_779_v209_ = nw150_
                index139_ = default__.safeIndex(914, (d_779_v209_).length(0))
                (d_779_v209_)[index139_] = d_583_v46_
                index140_ = default__.safeIndex(914, (d_779_v209_).length(0))
                (d_779_v209_)[index140_] = d_583_v46_
                d_780_v210_: _dafny.Array
                nw151_ = _dafny.Array(False, 24)
                d_780_v210_ = nw151_
                index141_ = default__.safeIndex(970, (d_780_v210_).length(0))
                (d_780_v210_)[index141_] = (d_530_v0_) not in (d_583_v46_)
                index142_ = default__.safeIndex(970, (d_780_v210_).length(0))
                rhs159_ = (not(d_530_v0_)) in ((d_779_v209_)[default__.safeIndex(914, (d_779_v209_).length(0))])
                rhs160_ = (d_776_v206_).f8
                rhs161_ = d_530_v0_
                rhs162_ = default__.safeModuloInt(((d_776_v206_).f8) + ((d_776_v206_).fm4(d_580_v43_, d_581_v44_, d_530_v0_, globalState)), (d_776_v206_).f8)
                lhs94_ = globalState
                lhs95_ = d_780_v210_
                lhs96_ = default__.safeIndex(970, (d_780_v210_).length(0))
                lhs97_ = globalState
                d_530_v0_ = rhs159_
                lhs94_.f2 = rhs160_
                lhs95_[lhs96_] = rhs161_
                lhs97_.f2 = rhs162_
                d_781_v211_: _dafny.Seq
                d_781_v211_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nep"))
                (self).f6 = (0) - ((len(d_781_v211_)) * (449))
            elif True:
                d_701_v147_ = d_701_v147_
                d_782_v212_: D7
                d_782_v212_ = D7_DC24()
                d_782_v212_ = d_782_v212_
                d_530_v0_ = not(d_530_v0_)
                d_783_v213_: C1
                nw152_ = C1()
                nw152_.ctor__(p0, (self).f5)
                d_783_v213_ = nw152_
                d_784_v214_: _dafny.Map
                d_784_v214_ = _dafny.Map({d_783_v213_: d_530_v0_})
                index143_ = default__.safeIndex(425, (d_701_v147_).length(0))
                (d_701_v147_)[index143_] = len(d_784_v214_)
                d_785_v215_: _dafny.Array
                nw153_ = _dafny.Array(False, 23)
                d_785_v215_ = nw153_
                index144_ = default__.safeIndex(818, (d_785_v215_).length(0))
                (d_785_v215_)[index144_] = d_530_v0_
                index145_ = default__.safeIndex(425, (d_701_v147_).length(0))
                index146_ = default__.safeIndex(818, (d_785_v215_).length(0))
                rhs163_ = (self).f5
                rhs164_ = d_530_v0_
                lhs98_ = d_701_v147_
                lhs99_ = default__.safeIndex(425, (d_701_v147_).length(0))
                lhs100_ = d_785_v215_
                lhs101_ = default__.safeIndex(818, (d_785_v215_).length(0))
                lhs98_[lhs99_] = rhs163_
                lhs100_[lhs101_] = rhs164_
                d_786_v216_: _dafny.Map
                d_786_v216_ = _dafny.Map({d_785_v215_: d_581_v44_})
                d_786_v216_ = (d_786_v216_).set(d_785_v215_, d_581_v44_)
            d_787_v217_: _dafny.Array
            nw154_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
            d_787_v217_ = nw154_
            d_788_v218_: _dafny.Seq
            d_788_v218_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chcl"))
            index147_ = default__.safeIndex(702, (d_787_v217_).length(0))
            (d_787_v217_)[index147_] = d_788_v218_
            index148_ = default__.safeIndex(702, (d_787_v217_).length(0))
            (d_787_v217_)[index148_] = d_788_v218_
            d_789_v219_: _dafny.Map
            d_789_v219_ = _dafny.Map({d_530_v0_: self.f6})
            d_790_v220_: C0
            nw155_ = C0()
            nw155_.ctor__(((d_789_v219_)[d_530_v0_] if (d_530_v0_) in (d_789_v219_) else len(d_788_v218_)))
            d_790_v220_ = nw155_

    def m4(self, p0, globalState):
        if (p0 if p0 else p0):
            d_791_v0_: _dafny.Set
            d_791_v0_ = _dafny.Set({self.f6})
            d_792_v1_: D1
            d_792_v1_ = D1_DC4(_dafny.Set({self.f6}))
            d_793_v2_: str
            d_793_v2_ = _dafny.CodePoint('a')
            d_794_v3_: _dafny.Map
            d_794_v3_ = _dafny.Map({p0: self.f6})
            d_795_v4_: _dafny.Map
            d_795_v4_ = _dafny.Map({d_793_v2_: d_794_v3_})
            d_796_v5_: _dafny.Array
            nw156_ = _dafny.Array(None, 21)
            nw156_[int(0)] = p0
            nw156_[int(1)] = ((self).f5) in (d_791_v0_)
            nw156_[int(2)] = p0
            nw156_[int(3)] = p0
            nw156_[int(4)] = False
            nw156_[int(5)] = (not(p0)) and (p0)
            nw156_[int(6)] = (d_791_v0_).isdisjoint((d_792_v1_).cf4)
            nw156_[int(7)] = not (True) or (p0)
            nw156_[int(8)] = p0
            nw156_[int(9)] = not((len(_dafny.Set({False, False}))) >= (self.f6))
            nw156_[int(10)] = not((d_793_v2_) in (d_795_v4_))
            nw156_[int(11)] = True
            nw156_[int(12)] = p0
            nw156_[int(13)] = p0
            nw156_[int(14)] = True
            nw156_[int(15)] = p0
            nw156_[int(16)] = p0
            nw156_[int(17)] = p0
            nw156_[int(18)] = p0
            nw156_[int(19)] = p0
            nw156_[int(20)] = p0
            d_796_v5_ = nw156_
            index149_ = default__.safeIndex(24, (d_796_v5_).length(0))
            (d_796_v5_)[index149_] = ((0) - ((self).f5)) <= ((self).f5)
            index150_ = default__.safeIndex(24, (d_796_v5_).length(0))
            (d_796_v5_)[index150_] = (self).fm9(D1_DC5(), ((self).fm1(globalState)) == (p0), globalState)
            index151_ = default__.safeIndex(24, (d_796_v5_).length(0))
            (d_796_v5_)[index151_] = True
            d_797_v6_: _dafny.Set
            d_797_v6_ = _dafny.Set({p0, (self.f6) == (len(d_791_v0_)), p0})
            d_797_v6_ = ((_dafny.Set({(d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))], (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]})) - (d_797_v6_)) | (d_797_v6_)
            d_798_v7_: _dafny.Seq
            d_798_v7_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f5]))])
            d_799_v8_: _dafny.Map
            d_799_v8_ = _dafny.Map({self.f6: len(d_798_v7_)})
            d_800_v9_: _dafny.Seq
            d_800_v9_ = _dafny.SeqWithoutIsStrInference([p0, (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))], (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]])
            d_801_v10_: _dafny.MultiSet
            d_801_v10_ = _dafny.MultiSet([p0, not((d_800_v9_)[default__.safeIndex(self.f6, len(d_800_v9_))]), (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]])
            d_802_v11_: _dafny.Array
            nw157_ = _dafny.Array(None, 18)
            nw157_[int(0)] = self.f6
            nw157_[int(1)] = len(d_799_v8_)
            nw157_[int(2)] = (self).f5
            nw157_[int(3)] = (self).f5
            nw157_[int(4)] = (self).f5
            nw157_[int(5)] = (self).f5
            nw157_[int(6)] = self.f6
            nw157_[int(7)] = (_dafny.MultiSet(d_798_v7_)).cardinality
            nw157_[int(8)] = self.f6
            nw157_[int(9)] = 199
            nw157_[int(10)] = (self).f5
            nw157_[int(11)] = (self).f5
            nw157_[int(12)] = self.f6
            nw157_[int(13)] = (self).f5
            nw157_[int(14)] = self.f6
            nw157_[int(15)] = (self).f5
            nw157_[int(16)] = (d_801_v10_).cardinality
            nw157_[int(17)] = (self).f5
            d_802_v11_ = nw157_
            d_803_v12_: _dafny.Map
            d_803_v12_ = _dafny.Map({(self).f5: not((d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))])})
            d_804_v13_: D5
            d_804_v13_ = D5_DC17(len(default__.fm11((self).f5, (self).f5, (self).f5, (0) - (len(d_798_v7_)), globalState)), d_802_v11_, d_803_v12_)
            d_805_v16_: _dafny.Map
            d_805_v16_ = _dafny.Map({d_793_v2_: (self).f5})
            pat_let_tv9_ = d_802_v11_
            d_806_v18_: _dafny.Array
            nw158_ = _dafny.Array(None, 28)
            nw158_[int(0)] = d_804_v13_
            def iife69_():
                coll43_ = _dafny.Map()
                compr_43_: int
                for compr_43_ in _dafny.IntegerRange(171, 954):
                    d_807_v14_: int = compr_43_
                    if ((171) <= (d_807_v14_)) and ((d_807_v14_) < (954)):
                        coll43_[(d_807_v14_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljpgjdur"))))] = (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]
                return _dafny.Map(coll43_)
            nw158_[int(1)] = D5_DC17((self).f5, d_802_v11_, iife69_()
)
            def iife70_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in (d_791_v0_).Elements:
                    d_808_v15_: int = compr_44_
                    if (d_808_v15_) in (d_791_v0_):
                        coll44_[(d_808_v15_) + (len(d_803_v12_))] = (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]
                return _dafny.Map(coll44_)
            nw158_[int(2)] = D5_DC17((d_801_v10_).cardinality, d_802_v11_, iife70_()
)
            nw158_[int(3)] = (d_804_v13_ if (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))] else d_804_v13_)
            nw158_[int(4)] = d_804_v13_
            nw158_[int(5)] = d_804_v13_
            nw158_[int(6)] = d_804_v13_
            nw158_[int(7)] = d_804_v13_
            nw158_[int(8)] = d_804_v13_
            nw158_[int(9)] = d_804_v13_
            def iife71_(_pat_let13_0):
                def iife72_(d_809_dt__update__tmp_h0_):
                    def iife73_(_pat_let14_0):
                        def iife74_(d_810_dt__update_hcf28_h0_):
                            return D5_DC17((d_809_dt__update__tmp_h0_).cf27, d_810_dt__update_hcf28_h0_, (d_809_dt__update__tmp_h0_).cf29)
                        return iife74_(_pat_let14_0)
                    return iife73_(pat_let_tv9_)
                return iife72_(_pat_let13_0)
            nw158_[int(10)] = iife71_(d_804_v13_)
            nw158_[int(11)] = d_804_v13_
            nw158_[int(12)] = d_804_v13_
            nw158_[int(13)] = D5_DC17((0) - (self.f6), d_802_v11_, d_803_v12_)
            def iife75_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in (d_803_v12_).keys.Elements:
                    d_811_v17_: int = compr_45_
                    if (d_811_v17_) in (d_803_v12_):
                        coll45_[default__.safeModuloInt(d_811_v17_, (self).f5)] = (d_796_v5_)[default__.safeIndex(24, (d_796_v5_).length(0))]
                return _dafny.Map(coll45_)
            nw158_[int(14)] = D5_DC17(((d_805_v16_)[d_793_v2_] if (d_793_v2_) in (d_805_v16_) else (self).f5), d_802_v11_, iife75_()
)
            nw158_[int(15)] = d_804_v13_
            nw158_[int(16)] = D5_DC17(639, d_802_v11_, d_803_v12_)
            nw158_[int(17)] = d_804_v13_
            nw158_[int(18)] = d_804_v13_
            nw158_[int(19)] = d_804_v13_
            nw158_[int(20)] = d_804_v13_
            nw158_[int(21)] = d_804_v13_
            nw158_[int(22)] = d_804_v13_
            nw158_[int(23)] = d_804_v13_
            nw158_[int(24)] = d_804_v13_
            nw158_[int(25)] = d_804_v13_
            nw158_[int(26)] = d_804_v13_
            nw158_[int(27)] = D5_DC17(self.f6, d_802_v11_, d_803_v12_)
            d_806_v18_ = nw158_
            nw159_ = _dafny.Array(None, 6)
            nw159_[int(0)] = d_804_v13_
            nw159_[int(1)] = d_804_v13_
            nw159_[int(2)] = d_804_v13_
            nw159_[int(3)] = d_804_v13_
            nw159_[int(4)] = d_804_v13_
            nw159_[int(5)] = d_804_v13_
            d_806_v18_ = nw159_
            d_801_v10_ = d_801_v10_
        elif True:
            d_812_v19_: _dafny.Map
            d_812_v19_ = _dafny.Map({self.f6: self.f6})
            d_813_v20_: _dafny.Map
            d_813_v20_ = _dafny.Map({p0: p0})
            (globalState).f2 = default__.safeModuloInt(((d_812_v19_)[len(default__.fm10(-670, d_812_v19_, p0, self.f6, globalState))] if (len(default__.fm10(-670, d_812_v19_, p0, self.f6, globalState))) in (d_812_v19_) else self.f6), len(d_813_v20_))
            d_814_v21_: str
            d_814_v21_ = _dafny.CodePoint('h')
            d_814_v21_ = d_814_v21_
            (globalState).f2 = self.f6
            d_815_v22_: bool
            d_815_v22_ = False
            d_816_v23_: _dafny.Seq
            d_816_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            d_817_v24_: _dafny.Seq
            d_817_v24_ = _dafny.SeqWithoutIsStrInference([self.f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ml"))), (self).f5, len(d_816_v23_)])
            d_818_v25_: _dafny.Map
            d_818_v25_ = _dafny.Map({(self).f5: p0})
            d_819_v26_: _dafny.Seq
            d_819_v26_ = _dafny.SeqWithoutIsStrInference([not(not(d_815_v22_)), ((d_818_v25_)[(self).f5] if ((self).f5) in (d_818_v25_) else p0)])
            d_815_v22_ = (((self).f5) - (len(d_817_v24_))) == (len(d_819_v26_))
            d_820_v27_: _dafny.Set
            d_820_v27_ = _dafny.Set({880, len(default__.fm11((self).f5, 574, self.f6, self.f6, globalState)), self.f6})
            d_820_v27_ = _dafny.Set({(self).f5, (0) - (self.f6), default__.safeModuloInt((self).f5, (0) - (self.f6))})
        (self).f6 = (self).f5
        d_821_v28_: bool
        d_821_v28_ = False
        d_821_v28_ = p0
        d_822_v30_: _dafny.Seq
        d_822_v30_ = _dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({True: True}))])
        d_823_v31_: D7
        def iife76_():
            coll46_ = _dafny.Map()
            compr_46_: D3
            for compr_46_ in (d_822_v30_).Elements:
                d_824_v29_: D3 = compr_46_
                if (d_824_v29_) in (d_822_v30_):
                    coll46_[d_824_v29_] = False
            return _dafny.Map(coll46_)
        d_823_v31_ = D7_DC21(iife76_()
)
        pat_let_tv10_ = d_821_v28_
        pat_let_tv11_ = d_821_v28_
        pat_let_tv12_ = p0
        pat_let_tv13_ = p0
        def lambda33_(source11_):
            if source11_.is_DC22:
                d_825___mcc_h1_ = source11_.cf38
                d_826___mcc_h2_ = source11_.cf39
                d_827_cf39_ = d_826___mcc_h2_
                d_828_cf38_ = d_825___mcc_h1_
                return pat_let_tv10_
            elif source11_.is_DC23:
                d_829___mcc_h3_ = source11_.cf40
                d_830___mcc_h4_ = source11_.cf41
                d_831___mcc_h5_ = source11_.cf42
                d_832_cf42_ = d_831___mcc_h5_
                d_833_cf41_ = d_830___mcc_h4_
                d_834_cf40_ = d_829___mcc_h3_
                return not(pat_let_tv11_)
            elif source11_.is_DC24:
                return pat_let_tv12_
            elif source11_.is_DC21:
                d_835___mcc_h6_ = source11_.cf37
                d_836_cf37_ = d_835___mcc_h6_
                return True
            elif True:
                d_837___mcc_h7_ = source11_.cf43
                d_838_cf43_ = d_837___mcc_h7_
                return (_dafny.SeqWithoutIsStrInference([-767, len(_dafny.Set({pat_let_tv13_}))])) < (_dafny.SeqWithoutIsStrInference([self.f6]))

        if lambda33_(d_823_v31_):
            if p0:
                d_821_v28_ = d_821_v28_
                d_839_v32_: str
                d_839_v32_ = _dafny.CodePoint('b')
                d_840_v33_: _dafny.MultiSet
                d_840_v33_ = _dafny.MultiSet([d_839_v32_, d_839_v32_, d_839_v32_])
                d_841_v34_: D9
                d_841_v34_ = D9_DC27(d_840_v33_)
                d_840_v33_ = (d_840_v33_) | ((d_841_v34_).cf45)
                d_842_v35_: C0
                nw160_ = C0()
                nw160_.ctor__(default__.safeModuloInt((self).f5, 541))
                d_842_v35_ = nw160_
                d_821_v28_ = p0
                d_821_v28_ = ((len(_dafny.Map({True: self.f6}))) * ((self).f5)) > ((d_842_v35_).f8)
            elif True:
                d_843_v36_: _dafny.Seq
                d_843_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_843_v36_ = d_843_v36_
                d_844_v37_: _dafny.Array
                def lambda34_(d_845_v36_):
                    def lambda35_(d_846_i0_):
                        return (d_846_i0_) - (len(d_845_v36_))

                    return lambda35_

                init18_ = lambda34_(d_843_v36_)
                nw161_ = _dafny.Array(None, 6)
                for i0_18_ in range(nw161_.length(0)):
                    nw161_[i0_18_] = init18_(i0_18_)
                d_844_v37_ = nw161_
                d_844_v37_ = d_844_v37_
                d_847_v38_: _dafny.Set
                d_847_v38_ = _dafny.Set({False, d_821_v28_})
                d_848_v39_: C1
                nw162_ = C1()
                nw162_.ctor__(((self).fm2(self.f6, d_847_v38_, globalState)) - ((0) - ((0) - (self.f6))), (self.f6 if d_821_v28_ else 941))
                d_848_v39_ = nw162_
                d_849_v40_: C1
                nw163_ = C1()
                nw163_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srfo"))), (0) - ((d_848_v39_).fm13((0) - (self.f6), _dafny.MultiSet([d_843_v36_, d_843_v36_]), globalState)))
                d_849_v40_ = nw163_
                d_850_v41_: D0
                d_850_v41_ = D0_DC1(False, d_821_v28_)
                d_851_v42_: D9
                d_851_v42_ = D9_DC28(d_821_v28_, self.f6, (d_849_v40_).fm13(971, _dafny.MultiSet([d_843_v36_, d_843_v36_]), globalState))
                d_852_v43_: _dafny.Seq
                d_852_v43_ = _dafny.SeqWithoutIsStrInference([p0])
                d_853_v44_: _dafny.Array
                nw164_ = _dafny.Array(None, 26)
                nw164_[int(0)] = d_821_v28_
                nw164_[int(1)] = d_821_v28_
                nw164_[int(2)] = d_821_v28_
                nw164_[int(3)] = d_821_v28_
                nw164_[int(4)] = p0
                nw164_[int(5)] = (d_850_v41_).cf2
                nw164_[int(6)] = d_821_v28_
                nw164_[int(7)] = p0
                nw164_[int(8)] = False
                nw164_[int(9)] = d_821_v28_
                nw164_[int(10)] = p0
                nw164_[int(11)] = p0
                nw164_[int(12)] = p0
                nw164_[int(13)] = d_821_v28_
                nw164_[int(14)] = p0
                nw164_[int(15)] = p0
                nw164_[int(16)] = (d_851_v42_).cf46
                nw164_[int(17)] = d_821_v28_
                nw164_[int(18)] = True
                nw164_[int(19)] = p0
                nw164_[int(20)] = p0
                nw164_[int(21)] = (d_852_v43_)[default__.safeIndex((self).f5, len(d_852_v43_))]
                nw164_[int(22)] = p0
                nw164_[int(23)] = p0
                nw164_[int(24)] = p0
                nw164_[int(25)] = d_821_v28_
                d_853_v44_ = nw164_
                d_854_v45_: _dafny.Seq
                d_854_v45_ = _dafny.SeqWithoutIsStrInference([d_853_v44_, d_853_v44_])
                d_855_v46_: _dafny.Map
                d_855_v46_ = _dafny.Map({p0: (self).f5})
                d_856_v48_: str
                d_856_v48_ = _dafny.CodePoint('b')
                d_857_v49_: _dafny.Map
                d_857_v49_ = _dafny.Map({d_855_v46_: d_856_v48_})
                d_858_v50_: _dafny.Map
                def iife77_():
                    coll47_ = _dafny.Set()
                    compr_47_: int
                    for compr_47_ in _dafny.IntegerRange(340, 549):
                        d_859_v47_: int = compr_47_
                        if ((340) <= (d_859_v47_)) and ((d_859_v47_) < (549)):
                            coll47_ = coll47_.union(_dafny.Set([default__.safeModuloInt(d_859_v47_, len(d_852_v43_))]))
                    return _dafny.Set(coll47_)
                d_858_v50_ = _dafny.Map({d_854_v45_: (_dafny.Map({d_855_v46_: (d_843_v36_)[default__.safeIndex(len(iife77_()
                ), len(d_843_v36_))]})) == (d_857_v49_)})
                d_858_v50_ = (d_858_v50_).set(d_854_v45_, d_821_v28_)
            d_821_v28_ = not (p0) or ((_dafny.MultiSet([d_821_v28_])).issubset(_dafny.MultiSet([d_821_v28_])))
            d_860_v51_: _dafny.Map
            d_860_v51_ = _dafny.Map({p0: (self).f5})
            d_861_v52_: _dafny.Map
            d_861_v52_ = (d_860_v51_) | (d_860_v51_)
            source12_ = d_861_v52_
            d_862___mcc_h0_ = source12_
            d_863_cf7_ = d_862___mcc_h0_
            d_864_v53_: _dafny.Seq
            d_864_v53_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f5), (self).f5])
            d_865_v54_: _dafny.Map
            d_865_v54_ = _dafny.Map({self.f6: len(d_864_v53_)})
            d_866_v55_: str
            d_866_v55_ = _dafny.CodePoint('a')
            d_867_v56_: _dafny.Map
            d_867_v56_ = _dafny.Map({d_865_v54_: d_866_v55_})
            d_867_v56_ = d_867_v56_
            d_868_v58_: D10
            d_868_v58_ = D10_DC29(d_865_v54_)
            d_869_v59_: _dafny.Seq
            d_869_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcfhvxpm"))
            d_870_v60_: _dafny.Map
            d_870_v60_ = _dafny.Map({p0: d_869_v59_})
            d_871_v61_: _dafny.Array
            nw165_ = _dafny.Array(None, 26)
            nw165_[int(0)] = d_865_v54_
            nw165_[int(1)] = _dafny.Map({(self).f5: self.f6})
            nw165_[int(2)] = _dafny.Map({len(d_863_cf7_): (self).f5})
            def iife78_():
                coll48_ = _dafny.Map()
                compr_48_: int
                for compr_48_ in _dafny.IntegerRange(353, 769):
                    d_872_v57_: int = compr_48_
                    if ((353) <= (d_872_v57_)) and ((d_872_v57_) < (769)):
                        coll48_[(d_872_v57_) + (self.f6)] = self.f6
                return _dafny.Map(coll48_)
            nw165_[int(3)] = iife78_()
            
            nw165_[int(4)] = (d_868_v58_).cf49
            nw165_[int(5)] = (d_865_v54_) | (d_865_v54_)
            nw165_[int(6)] = d_865_v54_
            nw165_[int(7)] = (d_865_v54_) | (d_865_v54_)
            nw165_[int(8)] = d_865_v54_
            nw165_[int(9)] = (d_865_v54_) | (d_865_v54_)
            nw165_[int(10)] = d_865_v54_
            nw165_[int(11)] = d_865_v54_
            nw165_[int(12)] = (d_865_v54_).set(self.f6, len(d_870_v60_))
            nw165_[int(13)] = d_865_v54_
            nw165_[int(14)] = d_865_v54_
            nw165_[int(15)] = d_865_v54_
            nw165_[int(16)] = d_865_v54_
            nw165_[int(17)] = d_865_v54_
            nw165_[int(18)] = (d_865_v54_) | (d_865_v54_)
            nw165_[int(19)] = d_865_v54_
            nw165_[int(20)] = d_865_v54_
            nw165_[int(21)] = d_865_v54_
            nw165_[int(22)] = d_865_v54_
            nw165_[int(23)] = d_865_v54_
            nw165_[int(24)] = d_865_v54_
            nw165_[int(25)] = d_865_v54_
            d_871_v61_ = nw165_
            d_873_v63_: _dafny.Set
            d_873_v63_ = _dafny.Set({self.f6})
            index152_ = default__.safeIndex(750, (d_871_v61_).length(0))
            def iife79_():
                coll49_ = _dafny.Map()
                compr_49_: int
                for compr_49_ in (d_873_v63_).Elements:
                    d_874_v62_: int = compr_49_
                    if (d_874_v62_) in (d_873_v63_):
                        coll49_[default__.safeDivisionInt(d_874_v62_, self.f6)] = (self).f5
                return _dafny.Map(coll49_)
            (d_871_v61_)[index152_] = iife79_()
            
            index153_ = default__.safeIndex(750, (d_871_v61_).length(0))
            def iife80_():
                coll50_ = _dafny.Map()
                compr_50_: int
                for compr_50_ in (d_865_v54_).keys.Elements:
                    d_875_v64_: int = compr_50_
                    if (d_875_v64_) in (d_865_v54_):
                        coll50_[default__.safeModuloInt(d_875_v64_, (self).f5)] = (self).f5
                return _dafny.Map(coll50_)
            (d_871_v61_)[index153_] = (d_865_v54_) | (iife80_()
            )
            d_869_v59_ = d_869_v59_
            (globalState).f2 = (self).f5
            d_876_v65_: _dafny.MultiSet
            d_876_v65_ = _dafny.MultiSet([(self).f5])
            if (d_876_v65_).isdisjoint(d_876_v65_):
                d_876_v65_ = d_876_v65_
                d_877_v66_: _dafny.Array
                nw166_ = _dafny.Array(int(0), 29)
                d_877_v66_ = nw166_
                d_878_v67_: _dafny.Set
                d_878_v67_ = _dafny.Set({p0, p0, d_821_v28_, not(p0)})
                index154_ = default__.safeIndex(761, (d_877_v66_).length(0))
                (d_877_v66_)[index154_] = (self).fm2(self.f6, d_878_v67_, globalState)
                d_879_v68_: _dafny.Seq
                d_879_v68_ = _dafny.SeqWithoutIsStrInference([p0])
                index155_ = default__.safeIndex(761, (d_877_v66_).length(0))
                (d_877_v66_)[index155_] = (len(_dafny.Map({_dafny.MultiSet(d_879_v68_): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_880_i1_ in range(default__.abs(-838))])).set(default__.safeIndex(620, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_881_i1_ in range(default__.abs(-838))]))), _dafny.CodePoint('l'))}))) * ((self).f5)
                d_882_v69_: str
                d_882_v69_ = _dafny.CodePoint('g')
                d_883_v70_: _dafny.Map
                d_883_v70_ = _dafny.Map({self.f6: d_882_v69_})
                d_882_v69_ = ((d_883_v70_)[((d_877_v66_)[default__.safeIndex(761, (d_877_v66_).length(0))] if p0 else self.f6)] if (((d_877_v66_)[default__.safeIndex(761, (d_877_v66_).length(0))] if p0 else self.f6)) in (d_883_v70_) else _dafny.CodePoint('e'))
                d_884_v71_: _dafny.Array
                def lambda36_(d_885_p0_):
                    def lambda37_(d_886_i2_):
                        return d_885_p0_

                    return lambda37_

                init19_ = lambda36_(p0)
                nw167_ = _dafny.Array(None, 7)
                for i0_19_ in range(nw167_.length(0)):
                    nw167_[i0_19_] = init19_(i0_19_)
                d_884_v71_ = nw167_
                index156_ = default__.safeIndex(169, (d_884_v71_).length(0))
                (d_884_v71_)[index156_] = (self).fm1(globalState)
                d_887_v72_: D1
                d_887_v72_ = D1_DC5()
                index157_ = default__.safeIndex(169, (d_884_v71_).length(0))
                rhs165_ = p0
                rhs166_ = (self).f5
                rhs167_ = ((self).fm9(d_887_v72_, True, globalState)) and (d_821_v28_)
                lhs102_ = d_884_v71_
                lhs103_ = default__.safeIndex(169, (d_884_v71_).length(0))
                lhs104_ = globalState
                lhs102_[lhs103_] = rhs165_
                lhs104_.f2 = rhs166_
                d_821_v28_ = rhs167_
                d_888_v73_: _dafny.Array
                d_889_v74_: bool
                d_890_v75_: _dafny.MultiSet
                d_891_v76_: int
                out52_: _dafny.Array
                out53_: bool
                out54_: _dafny.MultiSet
                out55_: int
                out52_, out53_, out54_, out55_ = default__.m0(_dafny.CodePoint('m'), globalState)
                d_888_v73_ = out52_
                d_889_v74_ = out53_
                d_890_v75_ = out54_
                d_891_v76_ = out55_
            elif True:
                d_892_v77_: D1
                d_892_v77_ = D1_DC5()
                d_893_v78_: _dafny.Map
                d_893_v78_ = _dafny.Map({d_892_v77_: ((self).f5) - ((self).f5)})
                d_893_v78_ = (d_893_v78_).set(d_892_v77_, (0) - ((0) - (self.f6)))
                d_894_v79_: C0
                nw168_ = C0()
                nw168_.ctor__(780)
                d_894_v79_ = nw168_
                d_895_v80_: _dafny.Seq
                d_895_v80_ = _dafny.SeqWithoutIsStrInference([d_821_v28_, p0])
                d_896_v81_: _dafny.Seq
                d_896_v81_ = _dafny.SeqWithoutIsStrInference([len(d_895_v80_), self.f6])
                d_897_v82_: C0
                nw169_ = C0()
                nw169_.ctor__((self).fm2((self).f5, default__.fm21(-113, len(d_896_v81_), False, globalState), globalState))
                d_897_v82_ = nw169_
                d_898_v83_: _dafny.Seq
                d_898_v83_ = _dafny.SeqWithoutIsStrInference([d_894_v79_])
                d_897_v82_ = (d_898_v83_)[default__.safeIndex((d_894_v79_).f8, len(d_898_v83_))]
                d_899_v84_: _dafny.Array
                nw170_ = _dafny.Array(_dafny.Seq({}), 29)
                d_899_v84_ = nw170_
                index158_ = default__.safeIndex(154, (d_899_v84_).length(0))
                (d_899_v84_)[index158_] = d_896_v81_
                d_900_v85_: _dafny.Set
                d_900_v85_ = _dafny.Set({d_876_v65_, d_876_v65_})
                d_901_v86_: C1
                nw171_ = C1()
                nw171_.ctor__(len(d_900_v85_), (d_897_v82_).f8)
                d_901_v86_ = nw171_
                d_902_v87_: _dafny.Seq
                d_902_v87_ = _dafny.SeqWithoutIsStrInference([d_901_v86_])
                d_903_v88_: D11
                d_903_v88_ = D11_DC31(d_901_v86_)
                d_904_v89_: _dafny.Array
                nw172_ = _dafny.Array(None, 23)
                nw172_[int(0)] = d_901_v86_
                nw172_[int(1)] = d_901_v86_
                nw172_[int(2)] = d_901_v86_
                nw172_[int(3)] = d_901_v86_
                nw172_[int(4)] = d_901_v86_
                nw172_[int(5)] = d_901_v86_
                nw172_[int(6)] = d_901_v86_
                nw172_[int(7)] = (d_902_v87_)[default__.safeIndex((d_897_v82_).f8, len(d_902_v87_))]
                nw172_[int(8)] = d_901_v86_
                nw172_[int(9)] = d_901_v86_
                nw172_[int(10)] = d_901_v86_
                nw172_[int(11)] = (d_903_v88_).cf51
                nw172_[int(12)] = d_901_v86_
                nw172_[int(13)] = d_901_v86_
                nw172_[int(14)] = d_901_v86_
                nw172_[int(15)] = d_901_v86_
                nw172_[int(16)] = d_901_v86_
                nw172_[int(17)] = d_901_v86_
                nw172_[int(18)] = d_901_v86_
                nw172_[int(19)] = d_901_v86_
                nw172_[int(20)] = d_901_v86_
                nw172_[int(21)] = d_901_v86_
                nw172_[int(22)] = d_901_v86_
                d_904_v89_ = nw172_
                d_905_v90_: _dafny.Set
                d_905_v90_ = _dafny.Set({d_904_v89_, d_904_v89_, d_904_v89_, d_904_v89_, d_904_v89_})
                index159_ = default__.safeIndex(154, (d_899_v84_).length(0))
                (d_899_v84_)[index159_] = _dafny.SeqWithoutIsStrInference([len((_dafny.Set({d_904_v89_})) - (d_905_v90_)), (d_897_v82_).f8])
                d_906_v91_: _dafny.Array
                def lambda38_(d_907_i3_):
                    return default__.safeModuloInt(d_907_i3_, self.f6)

                init20_ = lambda38_
                nw173_ = _dafny.Array(None, 28)
                for i0_20_ in range(nw173_.length(0)):
                    nw173_[i0_20_] = init20_(i0_20_)
                d_906_v91_ = nw173_
                d_908_v92_: D5
                d_908_v92_ = D5_DC16(d_906_v91_)
                d_909_v93_: _dafny.Array
                nw174_ = _dafny.Array(None, 24)
                nw174_[int(0)] = d_906_v91_
                nw174_[int(1)] = d_906_v91_
                nw174_[int(2)] = d_906_v91_
                nw174_[int(3)] = d_906_v91_
                nw174_[int(4)] = d_906_v91_
                nw174_[int(5)] = d_906_v91_
                nw174_[int(6)] = d_906_v91_
                nw174_[int(7)] = d_906_v91_
                nw174_[int(8)] = d_906_v91_
                nw174_[int(9)] = d_906_v91_
                nw174_[int(10)] = d_906_v91_
                nw174_[int(11)] = d_906_v91_
                nw174_[int(12)] = d_906_v91_
                nw174_[int(13)] = d_906_v91_
                nw174_[int(14)] = d_906_v91_
                nw174_[int(15)] = (d_908_v92_).cf26
                nw174_[int(16)] = d_906_v91_
                nw174_[int(17)] = d_906_v91_
                nw174_[int(18)] = d_906_v91_
                nw174_[int(19)] = d_906_v91_
                nw174_[int(20)] = (d_908_v92_).cf26
                nw174_[int(21)] = d_906_v91_
                nw174_[int(22)] = d_906_v91_
                nw174_[int(23)] = d_906_v91_
                d_909_v93_ = nw174_
                index160_ = default__.safeIndex(644, (d_909_v93_).length(0))
                (d_909_v93_)[index160_] = d_906_v91_
                index161_ = default__.safeIndex(644, (d_909_v93_).length(0))
                rhs168_ = (0) - ((d_894_v79_).f8)
                rhs169_ = d_906_v91_
                lhs105_ = self
                lhs106_ = d_909_v93_
                lhs107_ = default__.safeIndex(644, (d_909_v93_).length(0))
                lhs105_.f6 = rhs168_
                lhs106_[lhs107_] = rhs169_
            d_910_v94_: str
            d_910_v94_ = _dafny.CodePoint('u')
            d_910_v94_ = d_910_v94_
        elif True:
            (globalState).f2 = (self).f5
            d_911_v95_: _dafny.Array
            nw175_ = _dafny.Array(None, 1)
            nw175_[int(0)] = p0
            d_911_v95_ = nw175_
            index162_ = default__.safeIndex(913, (d_911_v95_).length(0))
            (d_911_v95_)[index162_] = (self.f6) <= (self.f6)
            index163_ = default__.safeIndex(913, (d_911_v95_).length(0))
            (d_911_v95_)[index163_] = p0
            d_821_v28_ = p0
            d_912_v96_: _dafny.Seq
            d_912_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hjndu"))
            d_912_v96_ = d_912_v96_
            if (d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))]:
                d_821_v28_ = ((d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))]) == (p0)
                d_913_v97_: _dafny.MultiSet
                d_913_v97_ = _dafny.MultiSet([(d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))]])
                d_914_v98_: _dafny.Set
                d_914_v98_ = _dafny.Set({(0) - (len(d_912_v96_))})
                d_915_v99_: _dafny.MultiSet
                d_915_v99_ = _dafny.MultiSet([(d_913_v97_).issubset(d_913_v97_), (d_914_v98_).isdisjoint(d_914_v98_)])
                d_916_v100_: _dafny.Seq
                d_916_v100_ = _dafny.SeqWithoutIsStrInference([(d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))]])
                d_917_v101_: _dafny.Seq
                d_917_v101_ = _dafny.SeqWithoutIsStrInference([d_914_v98_])
                d_913_v97_ = ((d_915_v99_) - (_dafny.MultiSet((d_916_v100_).set(default__.safeIndex((0) - (self.f6), len(d_916_v100_)), p0)))).set((d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))], default__.abs(((self).f5 if False else len((d_917_v101_)[default__.safeIndex((self).f5, len(d_917_v101_))]))))
                index164_ = default__.safeIndex(913, (d_911_v95_).length(0))
                (d_911_v95_)[index164_] = not (d_821_v28_) or ((d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))])
                d_918_v102_: C1
                nw176_ = C1()
                nw176_.ctor__((self.f6) + ((self).f5), self.f6)
                d_918_v102_ = nw176_
                index165_ = default__.safeIndex(913, (d_911_v95_).length(0))
                (d_911_v95_)[index165_] = d_821_v28_
            elif True:
                d_919_v103_: C1
                nw177_ = C1()
                nw177_.ctor__((self).f5, (self).f5)
                d_919_v103_ = nw177_
                d_920_v104_: _dafny.Map
                d_920_v104_ = _dafny.Map({d_919_v103_: p0})
                d_921_v105_: _dafny.Map
                d_921_v105_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f5 for d_922_i4_ in range(default__.abs(808))]): d_920_v104_})
                d_923_v106_: _dafny.Array
                def lambda39_(d_924_v95_):
                    def lambda40_(d_925_i5_):
                        return (_dafny.SeqWithoutIsStrInference([(d_924_v95_)[default__.safeIndex(913, (d_924_v95_).length(0))], (d_924_v95_)[default__.safeIndex(913, (d_924_v95_).length(0))]])) + (_dafny.SeqWithoutIsStrInference([(d_924_v95_)[default__.safeIndex(913, (d_924_v95_).length(0))]]))

                    return lambda40_

                init21_ = lambda39_(d_911_v95_)
                nw178_ = _dafny.Array(None, 3)
                for i0_21_ in range(nw178_.length(0)):
                    nw178_[i0_21_] = init21_(i0_21_)
                d_923_v106_ = nw178_
                d_926_v107_: _dafny.Seq
                d_926_v107_ = _dafny.SeqWithoutIsStrInference([self.f6])
                rhs170_ = self.f6
                rhs171_ = _dafny.Map({d_926_v107_: d_920_v104_})
                rhs172_ = d_923_v106_
                lhs108_ = globalState
                lhs108_.f2 = rhs170_
                d_921_v105_ = rhs171_
                d_923_v106_ = rhs172_
                d_912_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kobrta"))
                d_927_v108_: _dafny.Seq
                d_927_v108_ = _dafny.SeqWithoutIsStrInference([(d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))], False])
                d_927_v108_ = ((d_927_v108_) + (d_927_v108_) if (p0) == ((d_911_v95_)[default__.safeIndex(913, (d_911_v95_).length(0))]) else d_927_v108_)
                d_928_v109_: D5
                d_928_v109_ = D5_DC18(d_821_v28_, (self).f5, not(d_821_v28_))
                rhs173_ = (len(d_926_v107_)) - (self.f6)
                rhs174_ = d_928_v109_
                lhs109_ = self
                lhs109_.f6 = rhs173_
                d_928_v109_ = rhs174_
                index166_ = default__.safeIndex(913, (d_911_v95_).length(0))
                (d_911_v95_)[index166_] = p0
        d_929_v110_: _dafny.Seq
        d_929_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gahubgf"))
        d_930_v111_: _dafny.Map
        d_930_v111_ = _dafny.Map({len(d_929_v110_): True})
        d_931_v112_: D0
        d_931_v112_ = D0_DC1(d_821_v28_, d_821_v28_)
        d_932_v113_: _dafny.MultiSet
        d_932_v113_ = _dafny.MultiSet([False])
        d_933_v114_: _dafny.Array
        nw179_ = _dafny.Array(int(0), 26)
        d_933_v114_ = nw179_
        d_934_v115_: _dafny.Map
        d_934_v115_ = _dafny.Map({d_933_v114_: d_821_v28_})
        d_935_v116_: str
        d_935_v116_ = _dafny.CodePoint('a')
        d_936_v117_: _dafny.Set
        d_936_v117_ = _dafny.Set({self.f6, self.f6})
        d_937_v118_: _dafny.Seq
        d_937_v118_ = _dafny.SeqWithoutIsStrInference([(self).f5, self.f6, (self).f5])
        d_938_v119_: _dafny.Map
        d_938_v119_ = _dafny.Map({p0: d_930_v111_})
        d_939_v120_: _dafny.MultiSet
        d_939_v120_ = _dafny.MultiSet([_dafny.Map({(_dafny.MultiSet([len(d_937_v118_), self.f6])).cardinality: p0}), ((d_938_v119_)[p0] if (p0) in (d_938_v119_) else d_930_v111_)])
        d_940_v121_: _dafny.Array
        nw180_ = _dafny.Array(None, 26)
        nw180_[int(0)] = d_821_v28_
        nw180_[int(1)] = (self).fm0(_dafny.MultiSet([d_930_v111_, d_930_v111_]), p0, d_929_v110_, globalState)
        nw180_[int(2)] = not(p0)
        nw180_[int(3)] = (d_931_v112_).cf1
        nw180_[int(4)] = d_821_v28_
        nw180_[int(5)] = not(not(d_821_v28_))
        nw180_[int(6)] = (_dafny.MultiSet([p0])) == (d_932_v113_)
        nw180_[int(7)] = ((d_934_v115_)[d_933_v114_] if (d_933_v114_) in (d_934_v115_) else d_821_v28_)
        nw180_[int(8)] = p0
        nw180_[int(9)] = True
        nw180_[int(10)] = not(p0)
        nw180_[int(11)] = ((default__.fm11(self.f6, (self).f5, (d_932_v113_).cardinality, self.f6, globalState)).set(default__.safeIndex((self).f5, len(default__.fm11(self.f6, (self).f5, (d_932_v113_).cardinality, self.f6, globalState))), d_935_v116_)) != (d_929_v110_)
        nw180_[int(12)] = (d_936_v117_).issubset(d_936_v117_)
        nw180_[int(13)] = p0
        nw180_[int(14)] = p0
        nw180_[int(15)] = p0
        nw180_[int(16)] = (self).fm0(d_939_v120_, p0, d_929_v110_, globalState)
        nw180_[int(17)] = d_821_v28_
        nw180_[int(18)] = d_821_v28_
        nw180_[int(19)] = d_821_v28_
        nw180_[int(20)] = (d_932_v113_).isdisjoint(d_932_v113_)
        nw180_[int(21)] = p0
        nw180_[int(22)] = p0
        nw180_[int(23)] = (p0) or (False)
        nw180_[int(24)] = d_821_v28_
        nw180_[int(25)] = d_821_v28_
        d_940_v121_ = nw180_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_940_v121_).length(0)):
            d_941_i6_: int = guard_loop_5_
            if (True) and (((0) <= (d_941_i6_)) and ((d_941_i6_) < ((d_940_v121_).length(0)))):
                (d_940_v121_)[(d_941_i6_)] = ((self).f5) == (self.f6)
        d_942_v122_: _dafny.Map
        d_942_v122_ = _dafny.Map({self.f6: d_929_v110_})
        d_943_v123_: _dafny.Set
        d_943_v123_ = _dafny.Set({not(d_821_v28_)})
        d_942_v122_ = (d_942_v122_).set((0) - ((self).fm2(self.f6, d_943_v123_, globalState)), d_929_v110_)


class C3(T1, T0):
    def  __init__(self):
        self._f6: int = int(0)
        self._f5: int = int(0)
        self._f7: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f7, f5, f6):
        (self)._f7 = f7
        (self)._f5 = f5
        (self).f6 = f6

    def fm2(self, p0, p1, globalState):
        def iife81_():
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in (_dafny.Map({(self).f5: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbwcwhig"))})).keys.Elements:
                d_944_v0_: int = compr_51_
                if (d_944_v0_) in (_dafny.Map({(self).f5: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbwcwhig"))})):
                    coll51_[(d_944_v0_) - ((self).f5)] = _dafny.CodePoint('a')
            return _dafny.Map(coll51_)
        def iife82_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(649, 978):
                d_945_v1_: int = compr_52_
                if ((649) <= (d_945_v1_)) and ((d_945_v1_) < (978)):
                    coll52_[(d_945_v1_) - ((self).f5)] = _dafny.CodePoint('h')
            return _dafny.Map(coll52_)
        return len(((iife81_()
        ) | (iife82_()
        )) | ((_dafny.Map({(_dafny.MultiSet([self.f6])).cardinality: _dafny.CodePoint('x')})) | (_dafny.Map({(_dafny.MultiSet([(self).f7])).cardinality: _dafny.CodePoint('u')}))))

    def fm0(self, p0, p1, p2, globalState):
        if False:
            return (self).f7
        elif True:
            return (self).f7

    def fm1(self, globalState):
        return (self).f7

    def m1(self, globalState):
        r0: int = int(0)
        r1: bool = False
        d_946_i0_: int
        d_946_i0_ = 0
        with _dafny.label("4"):
            while (D0_DC1((self).f7, (self).f7)).cf2:
                with _dafny.c_label("4"):
                    if (d_946_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_946_i0_ = (d_946_i0_) + (1)
                    d_947_v0_: _dafny.Map
                    d_947_v0_ = _dafny.Map({(self).f7: False})
                    r1 = ((self).f7 if (self).f7 else ((d_947_v0_)[(self).f7] if ((self).f7) in (d_947_v0_) else (self).f7))
                    d_948_v1_: D0
                    d_948_v1_ = D0_DC3((self).f7)
                    r1 = (d_948_v1_).cf3
                    d_949_v2_: _dafny.Set
                    d_949_v2_ = _dafny.Set({self.f6})
                    d_950_v3_: _dafny.Array
                    nw181_ = _dafny.Array(False, 9)
                    d_950_v3_ = nw181_
                    d_951_v4_: _dafny.Set
                    d_951_v4_ = _dafny.Set({(self).f7, (self).f7, (self).f7})
                    index167_ = default__.safeIndex(870, (d_950_v3_).length(0))
                    (d_950_v3_)[index167_] = (d_951_v4_).issubset(d_951_v4_)
                    d_952_v5_: D1
                    d_952_v5_ = D1_DC4(_dafny.Set({self.f6, len(_dafny.Map({(self).f7: 836})), self.f6, (self).f5, 337}))
                    d_953_v6_: _dafny.Set
                    d_953_v6_ = _dafny.Set({_dafny.Map({853: (self).f7})})
                    index168_ = default__.safeIndex(870, (d_950_v3_).length(0))
                    rhs175_ = (d_952_v5_).cf4
                    rhs176_ = (self).f7
                    rhs177_ = _dafny.Map({(d_953_v6_).isdisjoint(d_953_v6_): (self).f7})
                    lhs110_ = d_950_v3_
                    lhs111_ = default__.safeIndex(870, (d_950_v3_).length(0))
                    d_949_v2_ = rhs175_
                    lhs110_[lhs111_] = rhs176_
                    d_947_v0_ = rhs177_
                    d_954_v7_: _dafny.Seq
                    d_954_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eijplewq"))
                    d_955_v8_: D1
                    d_955_v8_ = D1_DC6(d_954_v7_, (self).f5)
                    d_956_v9_: _dafny.Seq
                    d_956_v9_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (0) - (self.f6), (self).f5, (self).fm2((self).f5, d_951_v4_, globalState)])
                    d_957_v10_: _dafny.Seq
                    d_957_v10_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f6), (d_955_v8_).cf6, (d_956_v9_)[default__.safeIndex(168, len(d_956_v9_))], 240, self.f6])
                    r0 = (d_957_v10_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnkhhatg")) for d_958_i1_ in range(default__.abs(877))]))), len(d_957_v10_))]
                    pass
            pass
        r1 = (self).f7
        d_959_v11_: _dafny.Array
        nw182_ = _dafny.Array(False, 4)
        d_959_v11_ = nw182_
        d_960_v12_: _dafny.Map
        d_960_v12_ = _dafny.Map({(self).f5: D0_DC3((self).f7)})
        d_961_v13_: _dafny.Map
        d_961_v13_ = _dafny.Map({self.f6: d_960_v12_})
        d_962_v14_: str
        d_962_v14_ = _dafny.CodePoint('q')
        d_963_v15_: _dafny.Map
        d_963_v15_ = _dafny.Map({d_962_v14_: 916})
        d_964_v16_: _dafny.Seq
        d_964_v16_ = _dafny.SeqWithoutIsStrInference([d_963_v15_])
        d_965_v17_: _dafny.MultiSet
        d_965_v17_ = _dafny.MultiSet([d_960_v12_, ((d_961_v13_)[(0) - (len(d_964_v16_))] if ((0) - (len(d_964_v16_))) in (d_961_v13_) else d_960_v12_)])
        index169_ = default__.safeIndex(709, (d_959_v11_).length(0))
        (d_959_v11_)[index169_] = not((d_965_v17_).ispropersubset(_dafny.MultiSet([d_960_v12_])))
        index170_ = default__.safeIndex(709, (d_959_v11_).length(0))
        (d_959_v11_)[index170_] = (self).f7
        d_966_v18_: _dafny.MultiSet
        d_966_v18_ = _dafny.MultiSet([(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))], (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]])
        d_967_i2_: int
        d_967_i2_ = 0
        with _dafny.label("5"):
            while ((d_966_v18_).cardinality) != ((0) - (self.f6)):
                with _dafny.c_label("5"):
                    if (d_967_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_967_i2_ = (d_967_i2_) + (1)
                    d_968_v19_: _dafny.Array
                    nw183_ = _dafny.Array(_dafny.Array(None, 0), 11)
                    d_968_v19_ = nw183_
                    d_969_v20_: _dafny.Array
                    nw184_ = _dafny.Array(int(0), 26)
                    d_969_v20_ = nw184_
                    index171_ = default__.safeIndex(201, (d_968_v19_).length(0))
                    (d_968_v19_)[index171_] = d_969_v20_
                    d_970_v21_: _dafny.Map
                    d_970_v21_ = _dafny.Map({(0) - ((self).f5): d_969_v20_})
                    index172_ = default__.safeIndex(201, (d_968_v19_).length(0))
                    (d_968_v19_)[index172_] = ((d_970_v21_)[self.f6] if (self.f6) in (d_970_v21_) else d_969_v20_)
                    d_971_v22_: _dafny.Seq
                    d_971_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgosnrps"))
                    d_972_v23_: _dafny.Map
                    d_972_v23_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_962_v14_ for d_973_i3_ in range(default__.abs(127))])) + (d_971_v22_): (d_968_v19_)[default__.safeIndex(201, (d_968_v19_).length(0))]})
                    d_974_v24_: D1
                    d_974_v24_ = D1_DC6(d_971_v22_, self.f6)
                    index173_ = default__.safeIndex(201, (d_968_v19_).length(0))
                    (d_968_v19_)[index173_] = ((d_972_v23_)[((d_974_v24_).cf5) + (d_971_v22_)] if (((d_974_v24_).cf5) + (d_971_v22_)) in (d_972_v23_) else d_969_v20_)
                    d_962_v14_ = d_962_v14_
                    if (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]:
                        d_975_v25_: _dafny.Seq
                        d_975_v25_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                        d_976_v26_: _dafny.Seq
                        d_976_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f5, self.f6])])
                        d_977_v27_: _dafny.Seq
                        d_977_v27_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                        r1 = (not((d_975_v25_) < ((d_976_v26_)[default__.safeIndex(len(_dafny.Set({len(d_977_v27_)})), len(d_976_v26_))])) if ((self).f7 if (self).f7 else (self).f7) else False)
                        d_978_v28_: _dafny.Set
                        d_978_v28_ = _dafny.Set({d_959_v11_, d_959_v11_, d_959_v11_})
                        d_979_v29_: _dafny.Array
                        nw185_ = _dafny.Array(None, 15)
                        nw185_[int(0)] = d_978_v28_
                        nw185_[int(1)] = _dafny.Set({d_959_v11_})
                        nw185_[int(2)] = _dafny.Set({d_959_v11_})
                        nw185_[int(3)] = d_978_v28_
                        nw185_[int(4)] = (d_978_v28_) | (d_978_v28_)
                        nw185_[int(5)] = d_978_v28_
                        nw185_[int(6)] = _dafny.Set({d_959_v11_, d_959_v11_, d_959_v11_, d_959_v11_})
                        nw185_[int(7)] = (d_978_v28_) - (d_978_v28_)
                        nw185_[int(8)] = d_978_v28_
                        nw185_[int(9)] = _dafny.Set({d_959_v11_})
                        nw185_[int(10)] = d_978_v28_
                        nw185_[int(11)] = _dafny.Set({d_959_v11_})
                        nw185_[int(12)] = d_978_v28_
                        nw185_[int(13)] = d_978_v28_
                        nw185_[int(14)] = d_978_v28_
                        d_979_v29_ = nw185_
                        index174_ = default__.safeIndex(155, (d_979_v29_).length(0))
                        (d_979_v29_)[index174_] = d_978_v28_
                        index175_ = default__.safeIndex(155, (d_979_v29_).length(0))
                        (d_979_v29_)[index175_] = d_978_v28_
                        d_980_v30_: _dafny.Set
                        d_980_v30_ = _dafny.Set({(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]})
                        d_980_v30_ = ((d_980_v30_ if (self).f7 else d_980_v30_)).intersection(d_980_v30_)
                        d_981_v31_: _dafny.Map
                        d_981_v31_ = _dafny.Map({d_959_v11_: (self).f5})
                        d_982_v32_: _dafny.Map
                        d_982_v32_ = _dafny.Map({(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]: (d_981_v31_) == (d_981_v31_)})
                        d_982_v32_ = (d_982_v32_) | (d_982_v32_)
                        d_983_v33_: _dafny.Seq
                        d_983_v33_ = _dafny.SeqWithoutIsStrInference([d_971_v22_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dajoybt")), (d_971_v22_).set(default__.safeIndex(self.f6, len(d_971_v22_)), d_962_v14_), d_971_v22_])
                        d_971_v22_ = (((d_983_v33_)[default__.safeIndex(self.f6, len(d_983_v33_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iugifadp")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "viaetyyar")))
                    elif True:
                        d_984_v34_: _dafny.Seq
                        d_984_v34_ = _dafny.SeqWithoutIsStrInference([self.f6, self.f6])
                        d_985_v35_: _dafny.Map
                        d_985_v35_ = _dafny.Map({((self).f5) != ((self).f5): (d_984_v34_)[default__.safeIndex((self).f5, len(d_984_v34_))]})
                        d_986_v36_: _dafny.Map
                        d_986_v36_ = d_985_v35_
                        d_985_v35_ = (d_985_v35_) | ((d_986_v36_))
                        d_987_v37_: _dafny.Seq
                        d_987_v37_ = _dafny.SeqWithoutIsStrInference([d_966_v18_, d_966_v18_])
                        (globalState).f2 = (default__.fm3((self).f5, 434, d_987_v37_, d_984_v34_, globalState)).cf6
                        r1 = (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]
                        d_988_v38_: C0
                        nw186_ = C0()
                        nw186_.ctor__(len(d_971_v22_))
                        d_988_v38_ = nw186_
                        r0 = (d_988_v38_).f8
                    pass
            pass
        index176_ = default__.safeIndex(709, (d_959_v11_).length(0))
        (d_959_v11_)[index176_] = (self).f7
        if (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]:
            d_989_v39_: C0
            nw187_ = C0()
            nw187_.ctor__((self).f5)
            d_989_v39_ = nw187_
            d_989_v39_ = d_989_v39_
            nw188_ = C0()
            nw188_.ctor__((d_989_v39_).f8)
            d_989_v39_ = nw188_
            d_990_v40_: _dafny.Seq
            d_990_v40_ = _dafny.SeqWithoutIsStrInference([(self).f5])
            d_991_v41_: _dafny.Map
            d_991_v41_ = _dafny.Map({d_989_v39_: (self).f7})
            d_992_v42_: _dafny.Map
            d_992_v42_ = _dafny.Map({(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]: d_962_v14_})
            d_993_v43_: _dafny.Array
            nw189_ = _dafny.Array(None, 11)
            nw189_[int(0)] = ((d_989_v39_).f8) - (len(d_990_v40_))
            nw189_[int(1)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sksxj")))
            nw189_[int(2)] = ((d_989_v39_).f8) + ((0) - ((d_989_v39_).f8))
            nw189_[int(3)] = (self).f5
            nw189_[int(4)] = (d_990_v40_)[default__.safeIndex(len(d_991_v41_), len(d_990_v40_))]
            nw189_[int(5)] = self.f6
            nw189_[int(6)] = self.f6
            nw189_[int(7)] = (0) - ((d_989_v39_).fm4(d_962_v14_, default__.fm5(d_990_v40_, (self).f7, self.f6, globalState), (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))], globalState))
            nw189_[int(8)] = len(d_992_v42_)
            nw189_[int(9)] = ((self).f5) - (self.f6)
            nw189_[int(10)] = (d_989_v39_).f8
            d_993_v43_ = nw189_
            d_994_v44_: _dafny.Map
            d_994_v44_ = _dafny.Map({d_993_v43_: self.f6})
            d_995_v45_: _dafny.Array
            nw190_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
            d_995_v45_ = nw190_
            d_996_v46_: _dafny.Seq
            d_996_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfn"))
            d_997_v47_: _dafny.Map
            d_997_v47_ = _dafny.Map({(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]: d_993_v43_})
            d_998_v48_: _dafny.MultiSet
            d_998_v48_ = _dafny.MultiSet([d_959_v11_])
            rhs178_ = ((d_997_v47_)[(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]] if ((d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]) in (d_997_v47_) else d_993_v43_)
            rhs179_ = d_994_v44_
            rhs180_ = d_995_v45_
            rhs181_ = d_996_v46_
            rhs182_ = ((d_998_v48_)[d_959_v11_] if (d_959_v11_) in (d_998_v48_) else (d_989_v39_).fm4(d_962_v14_, default__.fm5(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f7: (d_989_v39_).f8})) for d_999_i4_ in range(default__.abs(530))]), (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))], (self).f5, globalState), (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))], globalState))
            d_993_v43_ = rhs178_
            d_994_v44_ = rhs179_
            d_995_v45_ = rhs180_
            d_996_v46_ = rhs181_
            r0 = rhs182_
            d_1000_v49_: _dafny.Map
            d_1000_v49_ = _dafny.Map({(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]: self.f6})
            (globalState).f2 = default__.safeDivisionInt((len(d_1000_v49_)) - ((d_989_v39_).f8), (d_989_v39_).f8)
            index177_ = default__.safeIndex(709, (d_959_v11_).length(0))
            (d_959_v11_)[index177_] = (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]
        elif True:
            (globalState).f2 = (self).f5
            index178_ = default__.safeIndex(709, (d_959_v11_).length(0))
            (d_959_v11_)[index178_] = True
            r1 = ((d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]) and (False)
            d_1001_v50_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Map({}), 18)
            d_1001_v50_ = nw191_
            d_1002_v51_: _dafny.Seq
            d_1002_v51_ = _dafny.SeqWithoutIsStrInference([d_1001_v50_, d_1001_v50_, d_1001_v50_])
            d_1001_v50_ = (d_1002_v51_)[default__.safeIndex(self.f6, len(d_1002_v51_))]
            d_1003_v52_: _dafny.Array
            def lambda41_(d_1004_v11_):
                def lambda42_(d_1005_i5_):
                    return (_dafny.SeqWithoutIsStrInference([(self).f7])) + (_dafny.SeqWithoutIsStrInference([(D0_DC1(not((d_1004_v11_)[default__.safeIndex(709, (d_1004_v11_).length(0))]), (d_1004_v11_)[default__.safeIndex(709, (d_1004_v11_).length(0))])).cf1]))

                return lambda42_

            init22_ = lambda41_(d_959_v11_)
            nw192_ = _dafny.Array(None, 7)
            for i0_22_ in range(nw192_.length(0)):
                nw192_[i0_22_] = init22_(i0_22_)
            d_1003_v52_ = nw192_
            d_1006_v53_: _dafny.Seq
            d_1006_v53_ = _dafny.SeqWithoutIsStrInference([(d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))], (d_959_v11_)[default__.safeIndex(709, (d_959_v11_).length(0))]])
            index179_ = default__.safeIndex(989, (d_1003_v52_).length(0))
            (d_1003_v52_)[index179_] = d_1006_v53_
            index180_ = default__.safeIndex(989, (d_1003_v52_).length(0))
            (d_1003_v52_)[index180_] = d_1006_v53_
        d_1007_v54_: _dafny.Array
        nw193_ = _dafny.Array(int(0), 26)
        d_1007_v54_ = nw193_
        d_1008_v55_: _dafny.Map
        d_1008_v55_ = _dafny.Map({(self).f5: d_1007_v54_})
        r0 = ((self).f5) * (default__.safeDivisionInt(self.f6, len(d_1008_v55_)))
        r1 = (self).f7
        return r0, r1

    def m2(self, p0, globalState):
        if ((True) == ((self).f7) if (self).f7 else (self).f7):
            d_1009_v0_: _dafny.Array
            nw194_ = _dafny.Array(_dafny.Seq({}), 17)
            d_1009_v0_ = nw194_
            d_1010_v1_: _dafny.Array
            def lambda43_(d_1011_i0_):
                return default__.safeModuloInt(d_1011_i0_, self.f6)

            init23_ = lambda43_
            nw195_ = _dafny.Array(None, 29)
            for i0_23_ in range(nw195_.length(0)):
                nw195_[i0_23_] = init23_(i0_23_)
            d_1010_v1_ = nw195_
            d_1012_v2_: _dafny.Seq
            d_1012_v2_ = _dafny.SeqWithoutIsStrInference([d_1010_v1_, d_1010_v1_])
            d_1013_v3_: _dafny.Seq
            d_1013_v3_ = _dafny.SeqWithoutIsStrInference([d_1010_v1_, (d_1012_v2_)[default__.safeIndex(p0, len(d_1012_v2_))]])
            index181_ = default__.safeIndex(197, (d_1009_v0_).length(0))
            (d_1009_v0_)[index181_] = d_1012_v2_
            index182_ = default__.safeIndex(197, (d_1009_v0_).length(0))
            (d_1009_v0_)[index182_] = d_1013_v3_
            d_1014_v4_: D0
            d_1014_v4_ = D0_DC1((self).f7, (self).f7)
            d_1015_v5_: _dafny.MultiSet
            d_1015_v5_ = _dafny.MultiSet([d_1014_v4_])
            d_1016_v6_: _dafny.Seq
            d_1016_v6_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1010_v1_: self.f6}))])
            d_1017_v7_: _dafny.Set
            d_1017_v7_ = _dafny.Set({d_1015_v5_, (d_1015_v5_).set(d_1014_v4_, default__.abs(len(d_1016_v6_)))})
            if (d_1017_v7_).issubset((d_1017_v7_ if False else d_1017_v7_)):
                d_1018_v8_: str
                d_1018_v8_ = _dafny.CodePoint('f')
                d_1018_v8_ = _dafny.CodePoint('t')
                d_1019_v9_: _dafny.Seq
                d_1019_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxjhrkcig"))
                d_1020_v10_: _dafny.Map
                d_1020_v10_ = _dafny.Map({(self).f7: p0})
                d_1021_v11_: _dafny.Set
                d_1021_v11_ = _dafny.Set({(self).f7})
                d_1022_v12_: D1
                d_1022_v12_ = D1_DC6(d_1019_v9_, (D1_DC6(default__.fm7(d_1018_v8_, (self).f7, (self).f7, (self).f5, globalState), len(d_1021_v11_))).cf6)
                d_1023_v13_: _dafny.Seq
                d_1023_v13_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_1024_v14_: _dafny.Array
                nw196_ = _dafny.Array(None, 21)
                nw196_[int(0)] = ((_dafny.SeqWithoutIsStrInference([d_1018_v8_ for d_1025_i1_ in range(default__.abs(-602))])) + (d_1019_v9_)).set(default__.safeIndex(806, len((_dafny.SeqWithoutIsStrInference([d_1018_v8_ for d_1026_i1_ in range(default__.abs(-602))])) + (d_1019_v9_))), default__.fm6(len(d_1020_v10_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjqkqeiu"))), (self).f7, (self).f7, globalState))
                nw196_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1018_v8_ for d_1027_i2_ in range(default__.abs(224))]) if not((self).f7) else d_1019_v9_)
                nw196_[int(2)] = d_1019_v9_
                nw196_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")))), d_1018_v8_)
                nw196_[int(4)] = (default__.fm7(d_1018_v8_, (self).f7, (self).f7, self.f6, globalState)).set(default__.safeIndex(p0, len(default__.fm7(d_1018_v8_, (self).f7, (self).f7, self.f6, globalState))), d_1018_v8_)
                nw196_[int(5)] = (d_1022_v12_).cf5
                nw196_[int(6)] = _dafny.SeqWithoutIsStrInference([d_1018_v8_ for d_1028_i3_ in range(default__.abs(430))])
                nw196_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1018_v8_ for d_1029_i4_ in range(default__.abs(-983))])
                nw196_[int(8)] = d_1019_v9_
                nw196_[int(9)] = d_1019_v9_
                nw196_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcns"))
                nw196_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmyra"))
                nw196_[int(12)] = d_1019_v9_
                nw196_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyonfg"))
                nw196_[int(14)] = d_1019_v9_
                nw196_[int(15)] = (d_1019_v9_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npwl")))
                nw196_[int(16)] = d_1019_v9_
                nw196_[int(17)] = d_1019_v9_
                nw196_[int(18)] = default__.fm7(_dafny.CodePoint('i'), (d_1023_v13_)[default__.safeIndex(len(d_1019_v9_), len(d_1023_v13_))], (self).f7, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "biesqdee"))), globalState)
                nw196_[int(19)] = d_1019_v9_
                nw196_[int(20)] = d_1019_v9_
                d_1024_v14_ = nw196_
                d_1030_v15_: _dafny.MultiSet
                d_1030_v15_ = _dafny.MultiSet([self.f6, self.f6])
                index183_ = default__.safeIndex(928, (d_1010_v1_).length(0))
                (d_1010_v1_)[index183_] = ((_dafny.MultiSet([p0])) | (d_1030_v15_)).cardinality
                index184_ = default__.safeIndex(928, (d_1010_v1_).length(0))
                rhs183_ = d_1024_v14_
                rhs184_ = (self).fm2(len(d_1016_v6_), d_1021_v11_, globalState)
                lhs112_ = d_1010_v1_
                lhs113_ = default__.safeIndex(928, (d_1010_v1_).length(0))
                d_1024_v14_ = rhs183_
                lhs112_[lhs113_] = rhs184_
                d_1031_v16_: _dafny.Map
                d_1031_v16_ = _dafny.Map({d_1018_v8_: d_1018_v8_})
                d_1032_v17_: _dafny.Array
                nw197_ = _dafny.Array(False, 7)
                d_1032_v17_ = nw197_
                d_1033_v18_: _dafny.MultiSet
                d_1033_v18_ = _dafny.MultiSet([d_1032_v17_, d_1032_v17_, d_1032_v17_])
                d_1034_v19_: T0
                nw198_ = C1()
                nw198_.ctor__(len(d_1023_v13_), (d_1033_v18_).cardinality)
                d_1034_v19_ = nw198_
                d_1035_v20_: _dafny.Map
                d_1035_v20_ = _dafny.Map({d_1034_v19_: _dafny.CodePoint('n')})
                d_1018_v8_ = ((d_1031_v16_)[((d_1035_v20_)[d_1034_v19_] if (d_1034_v19_) in (d_1035_v20_) else d_1018_v8_)] if (((d_1035_v20_)[d_1034_v19_] if (d_1034_v19_) in (d_1035_v20_) else d_1018_v8_)) in (d_1031_v16_) else (d_1019_v9_)[default__.safeIndex((self).f5, len(d_1019_v9_))])
                d_1036_v21_: _dafny.Array
                nw199_ = _dafny.Array(_dafny.Seq({}), 29)
                d_1036_v21_ = nw199_
                d_1037_v22_: D4
                d_1037_v22_ = D4_DC11(d_1016_v6_)
                d_1038_v23_: _dafny.Map
                d_1038_v23_ = _dafny.Map({d_1036_v21_: d_1037_v22_})
                def iife83_(_pat_let15_0):
                    def iife84_(d_1039_dt__update__tmp_h0_):
                        def iife85_(_pat_let16_0):
                            def iife86_(d_1041_dt__update_hcf13_h0_):
                                return D4_DC11(d_1041_dt__update_hcf13_h0_)
                            return iife86_(_pat_let16_0)
                        return iife85_(_dafny.SeqWithoutIsStrInference([456 for d_1040_i5_ in range(default__.abs(-764))]))
                    return iife84_(_pat_let15_0)
                d_1038_v23_ = (d_1038_v23_).set(d_1036_v21_, iife83_(d_1037_v22_))
                d_1023_v13_ = d_1023_v13_
            elif True:
                d_1042_v24_: T1
                nw200_ = C1()
                nw200_.ctor__((self).f5, (self).f5)
                d_1042_v24_ = nw200_
                d_1043_v25_: _dafny.Array
                nw201_ = _dafny.Array(None, 8)
                nw201_[int(0)] = d_1042_v24_
                nw201_[int(1)] = d_1042_v24_
                nw201_[int(2)] = d_1042_v24_
                nw201_[int(3)] = d_1042_v24_
                nw201_[int(4)] = d_1042_v24_
                nw201_[int(5)] = d_1042_v24_
                nw201_[int(6)] = d_1042_v24_
                nw201_[int(7)] = d_1042_v24_
                d_1043_v25_ = nw201_
                index185_ = default__.safeIndex(971, (d_1043_v25_).length(0))
                (d_1043_v25_)[index185_] = d_1042_v24_
                index186_ = default__.safeIndex(971, (d_1043_v25_).length(0))
                (d_1043_v25_)[index186_] = d_1042_v24_
                d_1044_v26_: C0
                nw202_ = C0()
                nw202_.ctor__(self.f6)
                d_1044_v26_ = nw202_
                d_1045_v27_: int
                d_1046_v28_: bool
                out56_: int
                out57_: bool
                out56_, out57_ = (d_1042_v24_).m1(globalState)
                d_1045_v27_ = out56_
                d_1046_v28_ = out57_
                d_1047_v29_: _dafny.Map
                d_1047_v29_ = _dafny.Map({d_1012_v2_: (d_1044_v26_).f8})
                (d_1042_v24_).f6 = ((d_1047_v29_)[_dafny.SeqWithoutIsStrInference([d_1010_v1_, d_1010_v1_])] if (_dafny.SeqWithoutIsStrInference([d_1010_v1_, d_1010_v1_])) in (d_1047_v29_) else (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1048_i6_ in range(default__.abs(351))]))) - (40))
                d_1049_v30_: str
                d_1049_v30_ = _dafny.CodePoint('s')
                d_1050_v31_: _dafny.Seq
                d_1050_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yern"))
                d_1046_v28_ = (d_1049_v30_) in (d_1050_v31_)
            def iife87_(_pat_let17_0):
                def iife88_(d_1051_dt__update__tmp_h1_):
                    def iife89_(_pat_let18_0):
                        def iife90_(d_1052_dt__update_hcf1_h0_):
                            return D0_DC1(d_1052_dt__update_hcf1_h0_, (d_1051_dt__update__tmp_h1_).cf2)
                        return iife90_(_pat_let18_0)
                    return iife89_((self).f7)
                return iife88_(_pat_let17_0)
            d_1014_v4_ = (d_1014_v4_ if (self).f7 else iife87_(d_1014_v4_))
            d_1053_v32_: _dafny.Map
            d_1053_v32_ = _dafny.Map({(self).f7: (self).f7})
            d_1054_v33_: D3
            d_1054_v33_ = D3_DC8(d_1053_v32_)
            pat_let_tv14_ = d_1053_v32_
            def iife91_(_pat_let19_0):
                def iife92_(d_1055_dt__update__tmp_h2_):
                    def iife93_(_pat_let20_0):
                        def iife94_(d_1056_dt__update_hcf8_h0_):
                            return D3_DC8(d_1056_dt__update_hcf8_h0_)
                        return iife94_(_pat_let20_0)
                    return iife93_(pat_let_tv14_)
                return iife92_(_pat_let19_0)
            source13_ = iife91_(d_1054_v33_)
            if source13_.is_DC9:
                d_1057___mcc_h0_ = source13_.cf9
                d_1058___mcc_h1_ = source13_.cf10
                d_1059___mcc_h2_ = source13_.cf11
                d_1060_cf11_ = d_1059___mcc_h2_
                d_1061_cf10_ = d_1058___mcc_h1_
                d_1062_cf9_ = d_1057___mcc_h0_
                d_1063_v34_: _dafny.Array
                nw203_ = _dafny.Array(None, 12)
                d_1063_v34_ = nw203_
                d_1064_v35_: C2
                nw204_ = C2()
                nw204_.ctor__((self).f5, self.f6)
                d_1064_v35_ = nw204_
                index187_ = default__.safeIndex(597, (d_1063_v34_).length(0))
                (d_1063_v34_)[index187_] = d_1064_v35_
                index188_ = default__.safeIndex(597, (d_1063_v34_).length(0))
                (d_1063_v34_)[index188_] = d_1064_v35_
                d_1065_v36_: _dafny.Map
                d_1065_v36_ = _dafny.Map({(self).f7: p0})
                d_1066_v37_: _dafny.Map
                d_1066_v37_ = d_1065_v36_
                d_1066_v37_ = _dafny.Map({(self).f7: self.f6})
                d_1067_v38_: C2
                nw205_ = C2()
                nw205_.ctor__(((self).f5) + (self.f6), self.f6)
                d_1067_v38_ = nw205_
                d_1068_v39_: _dafny.Seq
                d_1068_v39_ = _dafny.SeqWithoutIsStrInference([(D9_DC28((self).f7, (self).f5, (self).f5)).cf46])
                d_1068_v39_ = _dafny.SeqWithoutIsStrInference([(self).f7, (self).f7, (self).f7])
            elif source13_.is_DC8:
                d_1069___mcc_h3_ = source13_.cf8
                d_1070_cf8_ = d_1069___mcc_h3_
                d_1071_v40_: bool
                d_1071_v40_ = False
                d_1072_v41_: _dafny.Set
                d_1072_v41_ = _dafny.Set({d_1053_v32_, d_1053_v32_, d_1053_v32_, d_1070_cf8_})
                d_1071_v40_ = not(((d_1072_v41_) | (d_1072_v41_)).isdisjoint(d_1072_v41_))
                d_1073_v42_: _dafny.Seq
                d_1073_v42_ = _dafny.SeqWithoutIsStrInference([True])
                index189_ = default__.safeIndex(456, (d_1010_v1_).length(0))
                (d_1010_v1_)[index189_] = default__.safeDivisionInt((self).f5, len(d_1073_v42_))
                d_1074_v43_: _dafny.Map
                d_1074_v43_ = _dafny.Map({_dafny.CodePoint('m'): d_1010_v1_})
                d_1075_v44_: _dafny.Seq
                d_1075_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxsi"))
                d_1076_v45_: _dafny.MultiSet
                d_1076_v45_ = _dafny.MultiSet([d_1075_v44_])
                d_1077_v46_: str
                d_1077_v46_ = _dafny.CodePoint('d')
                index190_ = default__.safeIndex(456, (d_1010_v1_).length(0))
                rhs185_ = self.f6
                rhs186_ = -212
                rhs187_ = ((len(d_1016_v6_)) + (self.f6)) >= ((((d_1076_v45_)[d_1075_v44_] if (d_1075_v44_) in (d_1076_v45_) else 291)) - (p0))
                rhs188_ = _dafny.Map({d_1077_v46_: d_1010_v1_})
                lhs114_ = d_1010_v1_
                lhs115_ = default__.safeIndex(456, (d_1010_v1_).length(0))
                lhs116_ = self
                lhs114_[lhs115_] = rhs185_
                lhs116_.f6 = rhs186_
                d_1071_v40_ = rhs187_
                d_1074_v43_ = rhs188_
                d_1078_v47_: _dafny.Array
                def lambda44_(d_1079_v40_):
                    def lambda45_(d_1080_i7_):
                        return D0_DC0(d_1079_v40_)

                    return lambda45_

                init24_ = lambda44_(d_1071_v40_)
                nw206_ = _dafny.Array(None, 10)
                for i0_24_ in range(nw206_.length(0)):
                    nw206_[i0_24_] = init24_(i0_24_)
                d_1078_v47_ = nw206_
                d_1081_v48_: _dafny.Seq
                d_1081_v48_ = _dafny.SeqWithoutIsStrInference([d_1078_v47_])
                d_1082_v49_: D1
                d_1082_v49_ = D1_DC5()
                d_1083_v50_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Map({}), 28)
                d_1083_v50_ = nw207_
                rhs189_ = (_dafny.SeqWithoutIsStrInference([d_1078_v47_])) + (d_1081_v48_)
                rhs190_ = d_1082_v49_
                rhs191_ = d_1083_v50_
                d_1081_v48_ = rhs189_
                d_1082_v49_ = rhs190_
                d_1083_v50_ = rhs191_
                (globalState).f2 = (self).f5
            elif True:
                d_1084___mcc_h4_ = source13_.cf12
                d_1085_cf12_ = d_1084___mcc_h4_
                d_1086_v51_: _dafny.Seq
                d_1086_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_1087_v52_: _dafny.Map
                d_1087_v52_ = _dafny.Map({d_1016_v6_: p0})
                d_1088_v53_: D1
                d_1088_v53_ = D1_DC6(d_1086_v51_, self.f6)
                d_1089_v54_: T1
                nw208_ = C2()
                nw208_.ctor__(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1090_i8_ in range(default__.abs(826))])), (default__.fm5(_dafny.SeqWithoutIsStrInference([p0, self.f6, p0]), (self).f7, len((d_1088_v53_).cf5), globalState)).cardinality)
                d_1089_v54_ = nw208_
                d_1091_v55_: _dafny.Array
                nw209_ = _dafny.Array(None, 14)
                nw209_[int(0)] = (self).f7
                nw209_[int(1)] = ((self).f7) or ((self).f7)
                nw209_[int(2)] = (self).f7
                nw209_[int(3)] = (self).f7
                nw209_[int(4)] = (self).f7
                nw209_[int(5)] = (self).f7
                nw209_[int(6)] = (d_1086_v51_) <= (d_1086_v51_)
                nw209_[int(7)] = (self).f7
                nw209_[int(8)] = (self).f7
                nw209_[int(9)] = not(((self).f7) or ((self).f7))
                nw209_[int(10)] = (p0) < (self.f6)
                nw209_[int(11)] = (default__.fm28(p0, globalState)) in (d_1087_v52_)
                nw209_[int(12)] = (_dafny.Set({d_1089_v54_, d_1089_v54_, d_1089_v54_, d_1089_v54_, d_1089_v54_})).ispropersubset(_dafny.Set({d_1089_v54_}))
                nw209_[int(13)] = not(((self).f7) or (False))
                d_1091_v55_ = nw209_
                d_1092_v56_: _dafny.Set
                d_1092_v56_ = _dafny.Set({len(d_1086_v51_)})
                index191_ = default__.safeIndex(150, (d_1091_v55_).length(0))
                (d_1091_v55_)[index191_] = ((self).f5) in (d_1092_v56_)
                index192_ = default__.safeIndex(150, (d_1091_v55_).length(0))
                (d_1091_v55_)[index192_] = not((self).f7)
                d_1093_v57_: _dafny.Seq
                d_1093_v57_ = _dafny.SeqWithoutIsStrInference([(d_1091_v55_)[default__.safeIndex(150, (d_1091_v55_).length(0))]])
                d_1094_v58_: _dafny.MultiSet
                d_1094_v58_ = _dafny.MultiSet([_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(d_1091_v55_)[default__.safeIndex(150, (d_1091_v55_).length(0))]])) + (d_1093_v57_))])
                (globalState).f2 = (d_1094_v58_).cardinality
                d_1095_v59_: _dafny.Map
                d_1095_v59_ = _dafny.Map({(d_1091_v55_)[default__.safeIndex(150, (d_1091_v55_).length(0))]: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffv")))})
                index193_ = default__.safeIndex(150, (d_1091_v55_).length(0))
                (d_1091_v55_)[index193_] = (-77) <= (default__.safeModuloInt((self).fm2(len(d_1095_v59_), _dafny.Set({(self).f7, (d_1091_v55_)[default__.safeIndex(150, (d_1091_v55_).length(0))], (self).f7}), globalState), p0))
                index194_ = default__.safeIndex(150, (d_1091_v55_).length(0))
                (d_1091_v55_)[index194_] = (self).f7
            d_1096_v60_: C2
            nw210_ = C2()
            nw210_.ctor__((d_1016_v6_)[default__.safeIndex(self.f6, len(d_1016_v6_))], (self).f5)
            d_1096_v60_ = nw210_
        elif True:
            d_1097_v61_: str
            d_1097_v61_ = _dafny.CodePoint('a')
            d_1097_v61_ = d_1097_v61_
            d_1098_v62_: _dafny.Array
            def lambda46_(d_1099_i9_):
                return default__.safeModuloInt(d_1099_i9_, (self).f5)

            init25_ = lambda46_
            nw211_ = _dafny.Array(None, 10)
            for i0_25_ in range(nw211_.length(0)):
                nw211_[i0_25_] = init25_(i0_25_)
            d_1098_v62_ = nw211_
            d_1100_v63_: _dafny.Map
            d_1100_v63_ = _dafny.Map({d_1097_v61_: (self).f7})
            d_1101_v64_: _dafny.MultiSet
            d_1101_v64_ = _dafny.MultiSet([d_1100_v63_, d_1100_v63_])
            d_1102_v65_: D12
            d_1102_v65_ = D12_DC33(d_1101_v64_)
            d_1103_v66_: _dafny.Seq
            d_1103_v66_ = _dafny.SeqWithoutIsStrInference([d_1100_v63_])
            d_1104_v67_: _dafny.Array
            nw212_ = _dafny.Array(None, 11)
            nw212_[int(0)] = d_1101_v64_
            nw212_[int(1)] = d_1101_v64_
            nw212_[int(2)] = d_1101_v64_
            nw212_[int(3)] = _dafny.MultiSet(default__.fm29(self.f6, globalState))
            nw212_[int(4)] = d_1101_v64_
            nw212_[int(5)] = (d_1102_v65_).cf55
            nw212_[int(6)] = (d_1101_v64_) - (_dafny.MultiSet(d_1103_v66_))
            nw212_[int(7)] = _dafny.MultiSet([d_1100_v63_])
            nw212_[int(8)] = (d_1101_v64_).intersection(d_1101_v64_)
            nw212_[int(9)] = d_1101_v64_
            nw212_[int(10)] = d_1101_v64_
            d_1104_v67_ = nw212_
            index195_ = default__.safeIndex(713, (d_1104_v67_).length(0))
            (d_1104_v67_)[index195_] = d_1101_v64_
            d_1105_v68_: _dafny.Seq
            d_1105_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsvv"))
            index196_ = default__.safeIndex(713, (d_1104_v67_).length(0))
            rhs192_ = d_1098_v62_
            rhs193_ = (d_1105_v68_)[default__.safeIndex(len(d_1105_v68_), len(d_1105_v68_))]
            rhs194_ = p0
            rhs195_ = _dafny.MultiSet([d_1100_v63_])
            lhs117_ = self
            lhs118_ = d_1104_v67_
            lhs119_ = default__.safeIndex(713, (d_1104_v67_).length(0))
            d_1098_v62_ = rhs192_
            d_1097_v61_ = rhs193_
            lhs117_.f6 = rhs194_
            lhs118_[lhs119_] = rhs195_
            d_1105_v68_ = d_1105_v68_
            d_1106_v69_: bool
            d_1106_v69_ = True
            d_1107_v70_: _dafny.Set
            d_1107_v70_ = _dafny.Set({self.f6})
            rhs196_ = d_1097_v61_
            rhs197_ = (_dafny.Set({self.f6})).ispropersubset((d_1107_v70_) - (d_1107_v70_))
            rhs198_ = (d_1097_v61_ if ((self).f5) >= (len(d_1105_v68_)) else d_1097_v61_)
            d_1097_v61_ = rhs196_
            d_1106_v69_ = rhs197_
            d_1097_v61_ = rhs198_
            d_1097_v61_ = (d_1097_v61_ if (self).f7 else d_1097_v61_)
        d_1108_i10_: int
        d_1108_i10_ = 0
        with _dafny.label("6"):
            while not ((self).f7) or (((self).f7) and ((self).f7)):
                with _dafny.c_label("6"):
                    if (d_1108_i10_) >= (100):
                        raise _dafny.Break("6")
                    d_1108_i10_ = (d_1108_i10_) + (1)
                    d_1109_v71_: _dafny.Array
                    nw213_ = _dafny.Array(_dafny.Set({}), 16)
                    d_1109_v71_ = nw213_
                    d_1110_v72_: _dafny.Set
                    d_1110_v72_ = _dafny.Set({(self).f7})
                    index197_ = default__.safeIndex(775, (d_1109_v71_).length(0))
                    (d_1109_v71_)[index197_] = d_1110_v72_
                    index198_ = default__.safeIndex(775, (d_1109_v71_).length(0))
                    (d_1109_v71_)[index198_] = d_1110_v72_
                    d_1111_v73_: _dafny.Map
                    d_1111_v73_ = _dafny.Map({(self).f7: p0})
                    d_1112_v74_: _dafny.Seq
                    d_1112_v74_ = _dafny.SeqWithoutIsStrInference([-750, (((d_1111_v73_)[(self).f7] if ((self).f7) in (d_1111_v73_) else -885)) * (537)])
                    (globalState).f2 = (d_1112_v74_)[default__.safeIndex((self).f5, len(d_1112_v74_))]
                    d_1113_v75_: C1
                    nw214_ = C1()
                    nw214_.ctor__((0) - ((p0 if (D12_DC35((self).f7, self.f6, 388)).cf57 else self.f6)), p0)
                    d_1113_v75_ = nw214_
                    d_1114_v76_: T1
                    nw215_ = C2()
                    nw215_.ctor__(p0, self.f6)
                    d_1114_v76_ = nw215_
                    d_1114_v76_ = d_1114_v76_
                    pass
            pass
        if not(not((self).f7)):
            d_1115_v77_: _dafny.Set
            d_1115_v77_ = _dafny.Set({not((self).f7)})
            (globalState).f2 = ((self).f5) - ((len(d_1115_v77_)) - ((self).f5))
            d_1116_v78_: _dafny.Array
            nw216_ = _dafny.Array(int(0), 13)
            d_1116_v78_ = nw216_
            index199_ = default__.safeIndex(446, (d_1116_v78_).length(0))
            (d_1116_v78_)[index199_] = self.f6
            d_1117_v79_: _dafny.MultiSet
            d_1117_v79_ = _dafny.MultiSet([(self).f7, (self).f7, (self).f7, (self).f7, (self).f7])
            index200_ = default__.safeIndex(446, (d_1116_v78_).length(0))
            (d_1116_v78_)[index200_] = ((d_1117_v79_)[(self).f7] if ((self).f7) in (d_1117_v79_) else p0)
            d_1118_v80_: _dafny.Array
            nw217_ = _dafny.Array(False, 9)
            d_1118_v80_ = nw217_
            d_1119_v81_: _dafny.Seq
            d_1119_v81_ = _dafny.SeqWithoutIsStrInference([(self).f5])
            index201_ = default__.safeIndex(910, (d_1118_v80_).length(0))
            (d_1118_v80_)[index201_] = (d_1119_v81_) == (d_1119_v81_)
            d_1120_v82_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
            d_1120_v82_ = nw218_
            d_1121_v83_: _dafny.Seq
            d_1121_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thwvihdbx"))
            d_1122_v84_: str
            d_1122_v84_ = _dafny.CodePoint('m')
            index202_ = default__.safeIndex(305, (d_1120_v82_).length(0))
            (d_1120_v82_)[index202_] = ((d_1121_v83_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no")))).set(default__.safeIndex((self).f5, len((d_1121_v83_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "no"))))), d_1122_v84_)
            index203_ = default__.safeIndex(910, (d_1118_v80_).length(0))
            index204_ = default__.safeIndex(305, (d_1120_v82_).length(0))
            rhs199_ = (default__.safeModuloInt((self).f5, p0)) + ((self).fm2(p0, _dafny.Set({(self).f7, (self).f7, (self).f7}), globalState))
            rhs200_ = (self).f7
            rhs201_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsn"))).set(default__.safeIndex(self.f6, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsn")))), (d_1121_v83_)[default__.safeIndex(p0, len(d_1121_v83_))])) + ((d_1121_v83_) + (d_1121_v83_))
            lhs120_ = globalState
            lhs121_ = d_1118_v80_
            lhs122_ = default__.safeIndex(910, (d_1118_v80_).length(0))
            lhs123_ = d_1120_v82_
            lhs124_ = default__.safeIndex(305, (d_1120_v82_).length(0))
            lhs120_.f2 = rhs199_
            lhs121_[lhs122_] = rhs200_
            lhs123_[lhs124_] = rhs201_
            index205_ = default__.safeIndex(910, (d_1118_v80_).length(0))
            (d_1118_v80_)[index205_] = False
            d_1123_v85_: _dafny.Map
            d_1123_v85_ = _dafny.Map({(d_1116_v78_)[default__.safeIndex(446, (d_1116_v78_).length(0))]: 525})
            d_1124_v86_: _dafny.Map
            d_1124_v86_ = _dafny.Map({default__.safeModuloInt(self.f6, (0) - ((d_1116_v78_)[default__.safeIndex(446, (d_1116_v78_).length(0))])): ((self).fm2(301, _dafny.Set({(self).f7, (d_1118_v80_)[default__.safeIndex(910, (d_1118_v80_).length(0))]}), globalState)) == (len(d_1123_v85_))})
            d_1124_v86_ = (d_1124_v86_).set((self).f5, (_dafny.SeqWithoutIsStrInference([(0) - ((d_1116_v78_)[default__.safeIndex(446, (d_1116_v78_).length(0))])])) == (d_1119_v81_))
        elif True:
            d_1125_v87_: _dafny.Seq
            d_1125_v87_ = _dafny.SeqWithoutIsStrInference([(self).f7, not((self).f7), (self).f7, (self).f7])
            d_1126_v88_: _dafny.Set
            d_1126_v88_ = _dafny.Set({False})
            d_1127_v89_: _dafny.Array
            nw219_ = _dafny.Array(None, 13)
            nw219_[int(0)] = (self).f7
            nw219_[int(1)] = (self).f7
            nw219_[int(2)] = (self).f7
            nw219_[int(3)] = (d_1125_v87_) <= (d_1125_v87_)
            nw219_[int(4)] = (d_1126_v88_).isdisjoint(d_1126_v88_)
            nw219_[int(5)] = (self).f7
            nw219_[int(6)] = (self).f7
            nw219_[int(7)] = (d_1126_v88_) != (_dafny.Set({(self).f7, not((self).f7), (self).f7, (self).f7, (self).f7}))
            nw219_[int(8)] = (self).f7
            nw219_[int(9)] = (self).f7
            nw219_[int(10)] = (not(False) if (self).f7 else (self).f7)
            nw219_[int(11)] = (self).f7
            nw219_[int(12)] = (self).f7
            d_1127_v89_ = nw219_
            d_1127_v89_ = d_1127_v89_
            if (self).f7:
                d_1128_v90_: bool
                d_1128_v90_ = False
                d_1128_v90_ = ((self).f5) >= ((self).fm2(self.f6, d_1126_v88_, globalState))
                index206_ = default__.safeIndex(712, (d_1127_v89_).length(0))
                (d_1127_v89_)[index206_] = (self).f7
                index207_ = default__.safeIndex(712, (d_1127_v89_).length(0))
                (d_1127_v89_)[index207_] = (self).fm1(globalState)
                d_1129_v91_: _dafny.Array
                nw220_ = _dafny.Array(None, 4)
                nw220_[int(0)] = ((self).f5) + (-587)
                nw220_[int(1)] = (self).f5
                nw220_[int(2)] = p0
                nw220_[int(3)] = default__.safeDivisionInt(p0, 35)
                d_1129_v91_ = nw220_
                index208_ = default__.safeIndex(692, (d_1129_v91_).length(0))
                (d_1129_v91_)[index208_] = 428
                index209_ = default__.safeIndex(692, (d_1129_v91_).length(0))
                (d_1129_v91_)[index209_] = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiamx"))) if d_1128_v90_ else self.f6)) + ((p0) * (140))
                d_1130_v92_: D4
                d_1130_v92_ = D4_DC14(p0, d_1128_v90_, (d_1129_v91_)[default__.safeIndex(692, (d_1129_v91_).length(0))], True, self.f6)
                d_1131_v93_: _dafny.Map
                d_1131_v93_ = _dafny.Map({(d_1130_v92_).cf22: d_1129_v91_})
                d_1131_v93_ = (d_1131_v93_).set((self).f5, d_1129_v91_)
                d_1132_v94_: _dafny.Seq
                d_1132_v94_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1133_v96_: D1
                def iife95_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in (d_1132_v94_).Elements:
                        d_1134_v95_: int = compr_53_
                        if (d_1134_v95_) in (d_1132_v94_):
                            coll53_ = coll53_.union(_dafny.Set([default__.safeDivisionInt(d_1134_v95_, len(_dafny.SeqWithoutIsStrInference([not(False), True])))]))
                    return _dafny.Set(coll53_)
                d_1133_v96_ = D1_DC4(iife95_()
)
                d_1135_v97_: _dafny.Set
                d_1135_v97_ = _dafny.Set({(self).f5, p0, (d_1129_v91_)[default__.safeIndex(692, (d_1129_v91_).length(0))], p0})
                pat_let_tv15_ = d_1135_v97_
                def iife96_(_pat_let21_0):
                    def iife97_(d_1136_dt__update__tmp_h4_):
                        def iife98_(_pat_let22_0):
                            def iife99_(d_1137_dt__update_hcf4_h0_):
                                return D1_DC4(d_1137_dt__update_hcf4_h0_)
                            return iife99_(_pat_let22_0)
                        return iife98_(pat_let_tv15_)
                    return iife97_(_pat_let21_0)
                d_1133_v96_ = iife96_(d_1133_v96_)
            elif True:
                d_1138_v98_: bool
                d_1138_v98_ = False
                d_1139_v99_: _dafny.Seq
                d_1139_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ympy"))
                d_1140_v100_: T0
                nw221_ = C2()
                nw221_.ctor__(p0, self.f6)
                d_1140_v100_ = nw221_
                d_1141_v101_: _dafny.MultiSet
                d_1141_v101_ = _dafny.MultiSet([d_1140_v100_])
                d_1142_v103_: _dafny.Seq
                d_1142_v103_ = _dafny.SeqWithoutIsStrInference([self.f6])
                d_1143_v104_: _dafny.Map
                d_1143_v104_ = _dafny.Map({p0: (self).f5})
                d_1144_v105_: str
                d_1144_v105_ = _dafny.CodePoint('n')
                d_1145_v106_: D5
                d_1145_v106_ = D5_DC19(len(d_1126_v88_), d_1144_v105_, len(_dafny.Map({(self).f7: True})))
                def iife100_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in (d_1142_v103_).Elements:
                        d_1146_v102_: int = compr_54_
                        if (d_1146_v102_) in (d_1142_v103_):
                            coll54_[default__.safeModuloInt(d_1146_v102_, p0)] = 825
                    return _dafny.Map(coll54_)
                rhs202_ = (d_1140_v100_) not in (d_1141_v101_)
                rhs203_ = ((self).fm2((self).f5, _dafny.Set({(self).f7}), globalState)) in ((iife100_()
                ) | (d_1143_v104_))
                rhs204_ = d_1138_v98_
                rhs205_ = ((d_1139_v99_).set(default__.safeIndex(p0, len(d_1139_v99_)), (d_1145_v106_).cf34) if True else _dafny.SeqWithoutIsStrInference([d_1144_v105_ for d_1147_i11_ in range(default__.abs(272))]))
                d_1138_v98_ = rhs202_
                d_1138_v98_ = rhs203_
                d_1138_v98_ = rhs204_
                d_1139_v99_ = rhs205_
                d_1138_v98_ = (self).f7
                d_1148_v107_: C2
                nw222_ = C2()
                nw222_.ctor__(-813, (self).f5)
                d_1148_v107_ = nw222_
                d_1149_v108_: _dafny.Array
                def lambda47_(d_1150_i12_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfkrct"))

                init26_ = lambda47_
                nw223_ = _dafny.Array(None, 29)
                for i0_26_ in range(nw223_.length(0)):
                    nw223_[i0_26_] = init26_(i0_26_)
                d_1149_v108_ = nw223_
                d_1149_v108_ = (d_1149_v108_ if (d_1138_v98_) or (False) else d_1149_v108_)
                d_1138_v98_ = d_1138_v98_
            index210_ = default__.safeIndex(491, (d_1127_v89_).length(0))
            (d_1127_v89_)[index210_] = (True if (self).f7 else not((self).f7))
            index211_ = default__.safeIndex(491, (d_1127_v89_).length(0))
            (d_1127_v89_)[index211_] = (self).f7
            d_1151_v109_: _dafny.Map
            d_1151_v109_ = _dafny.Map({(d_1127_v89_)[default__.safeIndex(491, (d_1127_v89_).length(0))]: (d_1127_v89_)[default__.safeIndex(491, (d_1127_v89_).length(0))]})
            d_1151_v109_ = (d_1151_v109_).set((d_1127_v89_)[default__.safeIndex(491, (d_1127_v89_).length(0))], (self).f7)
            index212_ = default__.safeIndex(491, (d_1127_v89_).length(0))
            (d_1127_v89_)[index212_] = (d_1127_v89_)[default__.safeIndex(491, (d_1127_v89_).length(0))]
        d_1152_v110_: _dafny.MultiSet
        d_1152_v110_ = _dafny.MultiSet([(self).f5])
        (self).f6 = default__.safeModuloInt((self).f5, (d_1152_v110_).cardinality)
        d_1153_v111_: _dafny.Seq
        d_1153_v111_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfpqnji"))
        d_1154_v112_: str
        d_1154_v112_ = _dafny.CodePoint('s')
        d_1155_v113_: _dafny.Map
        d_1155_v113_ = _dafny.Map({(self).f5: (self).f7})
        d_1156_v114_: _dafny.Seq
        d_1156_v114_ = _dafny.SeqWithoutIsStrInference([d_1155_v113_])
        hi5_ = len(default__.fm7(d_1154_v112_, (self).fm0(_dafny.MultiSet(d_1156_v114_), (self).f7, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcfqhi")), globalState), (self).f7, (self).f5, globalState))
        for d_1157_i13_ in range((p0) - (len(d_1153_v111_)), hi5_):
            d_1158_v115_: D4
            d_1158_v115_ = D4_DC12((self).f7, d_1157_i13_, (self).f7)
            if ((d_1158_v115_).cf15) not in ((d_1152_v110_).intersection(_dafny.MultiSet([self.f6]))):
                d_1159_v116_: _dafny.Seq
                d_1159_v116_ = _dafny.SeqWithoutIsStrInference([(self).f5, (self).f5])
                d_1160_v117_: _dafny.Map
                d_1160_v117_ = _dafny.Map({d_1153_v111_: d_1159_v116_})
                d_1160_v117_ = d_1160_v117_
                (globalState).f2 = d_1157_i13_
                d_1161_v118_: _dafny.Seq
                d_1161_v118_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_1162_v119_: _dafny.Seq
                d_1162_v119_ = _dafny.SeqWithoutIsStrInference([(d_1161_v118_).set(default__.safeIndex(self.f6, len(d_1161_v118_)), (self).f7)])
                d_1163_v120_: D12
                d_1163_v120_ = D12_DC34(not((self).f7))
                d_1164_v121_: _dafny.Map
                def iife101_(_pat_let23_0):
                    def iife102_(d_1165_dt__update__tmp_h5_):
                        def iife103_(_pat_let24_0):
                            def iife104_(d_1166_dt__update_hcf56_h0_):
                                return D12_DC34(d_1166_dt__update_hcf56_h0_)
                            return iife104_(_pat_let24_0)
                        return iife103_(not((self).f7))
                    return iife102_(_pat_let23_0)
                d_1164_v121_ = _dafny.Map({self.f6: iife101_(d_1163_v120_)})
                d_1167_v122_: _dafny.Map
                d_1167_v122_ = _dafny.Map({(d_1162_v119_)[default__.safeIndex(len(d_1164_v121_), len(d_1162_v119_))]: default__.safeDivisionInt(d_1157_i13_, d_1157_i13_)})
                d_1167_v122_ = (d_1167_v122_).set((_dafny.SeqWithoutIsStrInference([True])).set(default__.safeIndex(len(d_1153_v111_), len(_dafny.SeqWithoutIsStrInference([True]))), True), default__.safeModuloInt(len(d_1155_v113_), self.f6))
                d_1168_v123_: bool
                d_1168_v123_ = True
                d_1169_v124_: _dafny.MultiSet
                d_1169_v124_ = _dafny.MultiSet([d_1154_v112_, d_1154_v112_])
                d_1168_v123_ = (d_1169_v124_).issubset(d_1169_v124_)
                d_1168_v123_ = d_1168_v123_
            elif True:
                d_1170_v125_: C1
                nw224_ = C1()
                nw224_.ctor__(d_1157_i13_, self.f6)
                d_1170_v125_ = nw224_
                d_1171_v126_: bool
                d_1171_v126_ = True
                d_1172_v127_: _dafny.Seq
                d_1172_v127_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_1171_v126_ = not((d_1172_v127_)[default__.safeIndex(d_1157_i13_, len(d_1172_v127_))])
                d_1173_v128_: _dafny.Array
                def lambda48_(d_1174_v126_):
                    def lambda49_(d_1175_i14_):
                        return d_1174_v126_

                    return lambda49_

                init27_ = lambda48_(d_1171_v126_)
                nw225_ = _dafny.Array(None, 4)
                for i0_27_ in range(nw225_.length(0)):
                    nw225_[i0_27_] = init27_(i0_27_)
                d_1173_v128_ = nw225_
                d_1176_v129_: D7
                d_1176_v129_ = D7_DC22(d_1173_v128_, d_1153_v111_)
                d_1177_v130_: D7
                d_1177_v130_ = D7_DC25(d_1176_v129_)
                d_1178_v131_: _dafny.Set
                d_1178_v131_ = _dafny.Set({d_1177_v130_})
                d_1179_v132_: _dafny.Seq
                d_1179_v132_ = _dafny.SeqWithoutIsStrInference([d_1178_v131_, d_1178_v131_])
                d_1180_v133_: C2
                nw226_ = C2()
                nw226_.ctor__(len(d_1153_v111_), len((d_1178_v131_) | ((d_1179_v132_)[default__.safeIndex(857, len(d_1179_v132_))])))
                d_1180_v133_ = nw226_
                d_1181_v134_: _dafny.Array
                nw227_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                d_1181_v134_ = nw227_
                d_1182_v135_: _dafny.Map
                d_1182_v135_ = _dafny.Map({d_1181_v134_: (self).f5})
                d_1182_v135_ = (d_1182_v135_).set(d_1181_v134_, d_1157_i13_)
                (globalState).f2 = (self).f5
            d_1183_v136_: D12
            d_1183_v136_ = D12_DC35((self).f7, 90, (222) * (self.f6))
            d_1183_v136_ = d_1183_v136_
            d_1184_v137_: _dafny.Set
            d_1184_v137_ = _dafny.Set({True, True})
            d_1185_v138_: _dafny.MultiSet
            d_1185_v138_ = _dafny.MultiSet([d_1184_v137_, d_1184_v137_, d_1184_v137_])
            d_1186_v139_: _dafny.Seq
            d_1186_v139_ = _dafny.SeqWithoutIsStrInference([d_1184_v137_, d_1184_v137_])
            if ((d_1184_v137_) - (d_1184_v137_)) in ((d_1185_v138_).intersection(_dafny.MultiSet(d_1186_v139_))):
                d_1184_v137_ = d_1184_v137_
                (globalState).f2 = default__.safeDivisionInt((p0) * ((0) - (len(default__.fm28(self.f6, globalState)))), (self).f5)
                d_1187_v140_: _dafny.Seq
                d_1187_v140_ = _dafny.SeqWithoutIsStrInference([False, (d_1183_v136_).cf57])
                d_1188_v141_: _dafny.Seq
                d_1188_v141_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f7, (self).f7, False])), d_1157_i13_, (self).fm2((self).f5, default__.fm21((_dafny.MultiSet(d_1187_v140_)).cardinality, (0) - ((self).f5), False, globalState), globalState), self.f6])
                (self).f6 = (default__.fm5(d_1188_v141_, ((d_1184_v137_)).issubset(d_1184_v137_), (d_1157_i13_) * (len(_dafny.SeqWithoutIsStrInference([d_1154_v112_ for d_1189_i15_ in range(default__.abs(103))]))), globalState)).cardinality
                d_1190_v142_: _dafny.Array
                nw228_ = _dafny.Array(None, 5)
                nw228_[int(0)] = not((self).f7)
                nw228_[int(1)] = not(True)
                nw228_[int(2)] = (self).f7
                nw228_[int(3)] = (d_1157_i13_) != (len(d_1153_v111_))
                nw228_[int(4)] = (self).f7
                d_1190_v142_ = nw228_
                index213_ = default__.safeIndex(936, (d_1190_v142_).length(0))
                (d_1190_v142_)[index213_] = (self).f7
                index214_ = default__.safeIndex(936, (d_1190_v142_).length(0))
                rhs206_ = d_1153_v111_
                rhs207_ = True
                lhs125_ = d_1190_v142_
                lhs126_ = default__.safeIndex(936, (d_1190_v142_).length(0))
                d_1153_v111_ = rhs206_
                lhs125_[lhs126_] = rhs207_
                d_1187_v140_ = _dafny.SeqWithoutIsStrInference([False])
            elif True:
                (globalState).f2 = d_1157_i13_
                d_1191_v143_: D4
                d_1191_v143_ = D4_DC13((self).f7, d_1157_i13_, (self).f7)
                pat_let_tv16_ = p0
                d_1192_v144_: _dafny.Array
                nw229_ = _dafny.Array(None, 24)
                nw229_[int(0)] = d_1191_v143_
                def iife105_(_pat_let25_0):
                    def iife106_(d_1193_dt__update__tmp_h6_):
                        def iife107_(_pat_let26_0):
                            def iife108_(d_1194_dt__update_hcf17_h0_):
                                return D4_DC13(d_1194_dt__update_hcf17_h0_, (d_1193_dt__update__tmp_h6_).cf18, (d_1193_dt__update__tmp_h6_).cf19)
                            return iife108_(_pat_let26_0)
                        return iife107_(False)
                    return iife106_(_pat_let25_0)
                nw229_[int(1)] = (iife105_(d_1191_v143_) if (self).f7 else default__.fm27(self.f6, globalState))
                nw229_[int(2)] = default__.fm27(self.f6, globalState)
                def iife109_(_pat_let27_0):
                    def iife110_(d_1195_dt__update__tmp_h7_):
                        def iife111_(_pat_let28_0):
                            def iife112_(d_1196_dt__update_hcf18_h0_):
                                return D4_DC13((d_1195_dt__update__tmp_h7_).cf17, d_1196_dt__update_hcf18_h0_, (d_1195_dt__update__tmp_h7_).cf19)
                            return iife112_(_pat_let28_0)
                        return iife111_(pat_let_tv16_)
                    return iife110_(_pat_let27_0)
                nw229_[int(3)] = iife109_(d_1191_v143_)
                nw229_[int(4)] = d_1191_v143_
                nw229_[int(5)] = d_1191_v143_
                nw229_[int(6)] = d_1191_v143_
                nw229_[int(7)] = default__.fm27(self.f6, globalState)
                nw229_[int(8)] = d_1191_v143_
                nw229_[int(9)] = D4_DC13((self).f7, self.f6, (self).f7)
                nw229_[int(10)] = d_1191_v143_
                nw229_[int(11)] = d_1191_v143_
                nw229_[int(12)] = D4_DC13((self).f7, p0, (self).f7)
                nw229_[int(13)] = d_1191_v143_
                nw229_[int(14)] = d_1191_v143_
                nw229_[int(15)] = default__.fm27(self.f6, globalState)
                nw229_[int(16)] = d_1191_v143_
                nw229_[int(17)] = (D4_DC13((self).f7, self.f6, not((self).f7)) if (self).f7 else D4_DC13((self).f7, p0, (self).f7))
                nw229_[int(18)] = D4_DC13((self).fm1(globalState), self.f6, (self).f7)
                nw229_[int(19)] = d_1191_v143_
                nw229_[int(20)] = d_1191_v143_
                nw229_[int(21)] = d_1191_v143_
                nw229_[int(22)] = d_1191_v143_
                nw229_[int(23)] = d_1191_v143_
                d_1192_v144_ = nw229_
                index215_ = default__.safeIndex(245, (d_1192_v144_).length(0))
                (d_1192_v144_)[index215_] = (d_1191_v143_ if (self).f7 else d_1191_v143_)
                index216_ = default__.safeIndex(245, (d_1192_v144_).length(0))
                (d_1192_v144_)[index216_] = d_1191_v143_
                d_1197_v145_: _dafny.Array
                nw230_ = _dafny.Array(None, 4)
                nw230_[int(0)] = d_1157_i13_
                nw230_[int(1)] = p0
                nw230_[int(2)] = self.f6
                nw230_[int(3)] = p0
                d_1197_v145_ = nw230_
                d_1198_v146_: _dafny.Map
                d_1198_v146_ = _dafny.Map({d_1184_v137_: d_1157_i13_})
                d_1199_v147_: _dafny.Map
                d_1199_v147_ = _dafny.Map({(self).f7: 256})
                d_1200_v148_: T1
                nw231_ = C2()
                nw231_.ctor__(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1201_i16_ in range(default__.abs(-827))])), (self).f5)
                d_1200_v148_ = nw231_
                d_1202_v149_: _dafny.Set
                d_1202_v149_ = _dafny.Set({d_1200_v148_, d_1200_v148_})
                nw232_ = _dafny.Array(None, 15)
                nw232_[int(0)] = p0
                nw232_[int(1)] = ((self).f5) - (d_1157_i13_)
                nw232_[int(2)] = d_1157_i13_
                nw232_[int(3)] = default__.safeDivisionInt(d_1157_i13_, -158)
                nw232_[int(4)] = ((d_1198_v146_)[d_1184_v137_] if (d_1184_v137_) in (d_1198_v146_) else (self).f5)
                nw232_[int(5)] = len((d_1199_v147_) | (_dafny.Map({(self).f7: 627})))
                nw232_[int(6)] = (self).f5
                nw232_[int(7)] = p0
                nw232_[int(8)] = len(d_1202_v149_)
                nw232_[int(9)] = d_1200_v148_.f6
                nw232_[int(10)] = d_1200_v148_.f6
                nw232_[int(11)] = ((d_1199_v147_)[(self).f7] if ((self).f7) in (d_1199_v147_) else (d_1158_v115_).cf15)
                nw232_[int(12)] = (self.f6 if (self).f7 else d_1157_i13_)
                nw232_[int(13)] = d_1200_v148_.f6
                nw232_[int(14)] = (0) - ((d_1200_v148_).fm2(p0, d_1184_v137_, globalState))
                d_1197_v145_ = nw232_
                index217_ = default__.safeIndex(601, (d_1197_v145_).length(0))
                (d_1197_v145_)[index217_] = self.f6
                d_1203_v150_: _dafny.MultiSet
                d_1203_v150_ = _dafny.MultiSet([(self).f7])
                index218_ = default__.safeIndex(601, (d_1197_v145_).length(0))
                (d_1197_v145_)[index218_] = (self).fm2(d_1200_v148_.f6, default__.fm21((0) - (len(d_1153_v111_)), (d_1203_v150_).cardinality, True, globalState), globalState)
                d_1204_v151_: C2
                nw233_ = C2()
                nw233_.ctor__((self).f5, d_1157_i13_)
                d_1204_v151_ = nw233_
            d_1205_v152_: _dafny.Seq
            d_1205_v152_ = _dafny.SeqWithoutIsStrInference([d_1157_i13_])
            (self).f6 = len((d_1205_v152_).set(default__.safeIndex(p0, len(d_1205_v152_)), d_1157_i13_))
        if (self.f6) <= (918):
            d_1206_v153_: bool
            d_1206_v153_ = False
            d_1206_v153_ = (self).f7
            d_1207_v154_: C0
            nw234_ = C0()
            nw234_.ctor__(p0)
            d_1207_v154_ = nw234_
            d_1208_v155_: _dafny.Seq
            d_1208_v155_ = _dafny.SeqWithoutIsStrInference([d_1207_v154_])
            d_1209_v156_: _dafny.Map
            d_1209_v156_ = _dafny.Map({len(d_1153_v111_): d_1208_v155_})
            d_1210_v157_: _dafny.Set
            d_1210_v157_ = _dafny.Set({d_1206_v153_})
            (self).f6 = (self).fm2(len(d_1209_v156_), d_1210_v157_, globalState)
            d_1211_v158_: D0
            d_1211_v158_ = D0_DC2()
            d_1212_v159_: D3
            d_1212_v159_ = D3_DC9(d_1152_v110_, d_1211_v158_, self.f6)
            d_1213_v160_: D3
            d_1213_v160_ = D3_DC10(d_1212_v159_)
            d_1213_v160_ = D3_DC10(d_1212_v159_)
            d_1214_v161_: _dafny.Array
            nw235_ = _dafny.Array(False, 6)
            d_1214_v161_ = nw235_
            d_1215_v162_: _dafny.Set
            d_1215_v162_ = _dafny.Set({d_1214_v161_})
            d_1216_v163_: _dafny.Map
            d_1216_v163_ = _dafny.Map({d_1154_v112_: d_1215_v162_})
            d_1217_v164_: _dafny.Map
            d_1217_v164_ = _dafny.Map({d_1206_v153_: d_1214_v161_})
            d_1218_v165_: _dafny.Array
            nw236_ = _dafny.Array(None, 15)
            nw236_[int(0)] = d_1215_v162_
            nw236_[int(1)] = d_1215_v162_
            nw236_[int(2)] = _dafny.Set({d_1214_v161_, d_1214_v161_, d_1214_v161_, d_1214_v161_, d_1214_v161_})
            nw236_[int(3)] = d_1215_v162_
            nw236_[int(4)] = d_1215_v162_
            nw236_[int(5)] = (d_1215_v162_) | (((d_1216_v163_)[d_1154_v112_] if (d_1154_v112_) in (d_1216_v163_) else d_1215_v162_))
            nw236_[int(6)] = d_1215_v162_
            nw236_[int(7)] = (_dafny.Set({((d_1217_v164_)[not(d_1206_v153_)] if (not(d_1206_v153_)) in (d_1217_v164_) else d_1214_v161_)})) | (d_1215_v162_)
            nw236_[int(8)] = d_1215_v162_
            nw236_[int(9)] = d_1215_v162_
            nw236_[int(10)] = _dafny.Set({d_1214_v161_, d_1214_v161_, d_1214_v161_})
            nw236_[int(11)] = d_1215_v162_
            nw236_[int(12)] = d_1215_v162_
            nw236_[int(13)] = _dafny.Set({d_1214_v161_})
            nw236_[int(14)] = (d_1215_v162_) | (d_1215_v162_)
            d_1218_v165_ = nw236_
            d_1219_v166_: _dafny.Map
            d_1219_v166_ = _dafny.Map({len(d_1153_v111_): d_1215_v162_})
            d_1220_v167_: _dafny.Seq
            d_1220_v167_ = _dafny.SeqWithoutIsStrInference([((d_1219_v166_)[len(_dafny.SeqWithoutIsStrInference([d_1154_v112_ for d_1221_i17_ in range(default__.abs(-798))]))] if (len(_dafny.SeqWithoutIsStrInference([d_1154_v112_ for d_1222_i17_ in range(default__.abs(-798))]))) in (d_1219_v166_) else d_1215_v162_)])
            index219_ = default__.safeIndex(659, (d_1218_v165_).length(0))
            (d_1218_v165_)[index219_] = (d_1220_v167_)[default__.safeIndex((self).f5, len(d_1220_v167_))]
            index220_ = default__.safeIndex(659, (d_1218_v165_).length(0))
            (d_1218_v165_)[index220_] = d_1215_v162_
            if not(d_1206_v153_):
                d_1223_v168_: _dafny.Set
                d_1223_v168_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkkcfpn")))})
                d_1224_v169_: _dafny.Set
                d_1224_v169_ = _dafny.Set({-116, len(d_1223_v168_), (d_1207_v154_).f8, (p0) - ((self).f5)})
                d_1223_v168_ = ((d_1223_v168_ if (self).f7 else d_1224_v169_)).intersection(d_1223_v168_)
                d_1225_v170_: T0
                nw237_ = C1()
                nw237_.ctor__(self.f6, (615 if not(d_1206_v153_) else (d_1207_v154_).f8))
                d_1225_v170_ = nw237_
                d_1226_v171_: _dafny.MultiSet
                d_1226_v171_ = _dafny.MultiSet([d_1206_v153_, (self).f7, d_1206_v153_, d_1206_v153_, (self).f7])
                d_1227_v172_: _dafny.Array
                nw238_ = _dafny.Array(int(0), 28)
                d_1227_v172_ = nw238_
                d_1228_v173_: _dafny.Map
                d_1228_v173_ = _dafny.Map({d_1227_v172_: d_1206_v153_})
                d_1229_v174_: C1
                nw239_ = C1()
                nw239_.ctor__(default__.safeDivisionInt((d_1207_v154_).f8, ((d_1226_v171_)[(self).f7] if ((self).f7) in (d_1226_v171_) else (d_1207_v154_).f8)), (len(d_1228_v173_)) - (self.f6))
                d_1229_v174_ = nw239_
                d_1230_v175_: _dafny.Array
                nw240_ = _dafny.Array(D4.default()(), 2)
                d_1230_v175_ = nw240_
                d_1231_v176_: _dafny.Seq
                d_1231_v176_ = _dafny.SeqWithoutIsStrInference([len(d_1153_v111_), 345])
                d_1232_v177_: D4
                d_1232_v177_ = D4_DC11(d_1231_v176_)
                index221_ = default__.safeIndex(653, (d_1230_v175_).length(0))
                (d_1230_v175_)[index221_] = d_1232_v177_
                d_1233_v178_: _dafny.Map
                d_1233_v178_ = _dafny.Map({d_1227_v172_: d_1229_v174_})
                d_1234_v179_: _dafny.Map
                d_1234_v179_ = _dafny.Map({d_1207_v154_: (d_1207_v154_).f8})
                d_1235_v180_: _dafny.Map
                d_1235_v180_ = _dafny.Map({d_1234_v179_: (self).f7})
                index222_ = default__.safeIndex(653, (d_1230_v175_).length(0))
                rhs208_ = d_1225_v170_
                rhs209_ = d_1229_v174_
                rhs210_ = d_1232_v177_
                rhs211_ = ((d_1235_v180_)[(d_1234_v179_) | (d_1234_v179_)] if ((d_1234_v179_) | (d_1234_v179_)) in (d_1235_v180_) else d_1206_v153_)
                rhs212_ = d_1233_v178_
                lhs127_ = d_1230_v175_
                lhs128_ = default__.safeIndex(653, (d_1230_v175_).length(0))
                d_1225_v170_ = rhs208_
                d_1229_v174_ = rhs209_
                lhs127_[lhs128_] = rhs210_
                d_1206_v153_ = rhs211_
                d_1233_v178_ = rhs212_
                d_1236_v181_: _dafny.Seq
                d_1236_v181_ = _dafny.SeqWithoutIsStrInference([d_1227_v172_, d_1227_v172_, d_1227_v172_, d_1227_v172_, d_1227_v172_])
                d_1206_v153_ = ((d_1236_v181_) != (d_1236_v181_) if d_1206_v153_ else (d_1210_v157_).ispropersubset(default__.fm21((d_1207_v154_).f8, (d_1207_v154_).f8, False, globalState)))
                d_1237_v182_: _dafny.MultiSet
                d_1237_v182_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nduue")), d_1153_v111_])
                d_1238_v183_: C1
                nw241_ = C1()
                nw241_.ctor__(self.f6, (d_1229_v174_).fm13(len(d_1153_v111_), d_1237_v182_, globalState))
                d_1238_v183_ = nw241_
                d_1239_v184_: _dafny.Array
                nw242_ = _dafny.Array(_dafny.Set({}), 7)
                d_1239_v184_ = nw242_
                d_1239_v184_ = d_1239_v184_
            elif True:
                d_1240_v185_: _dafny.Array
                nw243_ = _dafny.Array(None, 13)
                d_1240_v185_ = nw243_
                d_1240_v185_ = d_1240_v185_
                d_1241_v186_: _dafny.MultiSet
                d_1241_v186_ = _dafny.MultiSet([d_1206_v153_])
                d_1242_v187_: C2
                nw244_ = C2()
                nw244_.ctor__(p0, ((d_1241_v186_) - (_dafny.MultiSet([True]))).cardinality)
                d_1242_v187_ = nw244_
                d_1242_v187_ = d_1242_v187_
                d_1243_v188_: _dafny.Seq
                d_1243_v188_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_1206_v153_ = not(not((d_1243_v188_)[default__.safeIndex((d_1207_v154_).f8, len(d_1243_v188_))]))
                index223_ = default__.safeIndex(406, (d_1214_v161_).length(0))
                (d_1214_v161_)[index223_] = (self).f7
                d_1244_v189_: _dafny.Seq
                d_1244_v189_ = _dafny.SeqWithoutIsStrInference([(d_1207_v154_).f8, (d_1207_v154_).f8])
                index224_ = default__.safeIndex(406, (d_1214_v161_).length(0))
                (d_1214_v161_)[index224_] = not (((d_1207_v154_).f8) >= (len(d_1244_v189_))) or (d_1206_v153_)
                index225_ = default__.safeIndex(406, (d_1214_v161_).length(0))
                (d_1214_v161_)[index225_] = (self).f7
        elif True:
            if not((default__.fm30(globalState)).issubset((_dafny.MultiSet([d_1153_v111_, d_1153_v111_])).intersection(_dafny.MultiSet([d_1153_v111_])))):
                d_1245_v190_: C0
                nw245_ = C0()
                nw245_.ctor__(389)
                d_1245_v190_ = nw245_
                d_1246_v191_: _dafny.Seq
                d_1246_v191_ = _dafny.SeqWithoutIsStrInference([(d_1245_v190_).f8])
                d_1247_v192_: _dafny.MultiSet
                d_1247_v192_ = _dafny.MultiSet([(self).f7])
                d_1248_v193_: T0
                nw246_ = C1()
                nw246_.ctor__((d_1246_v191_)[default__.safeIndex((d_1245_v190_).fm4(d_1154_v112_, d_1247_v192_, (self).f7, globalState), len(d_1246_v191_))], (self).f5)
                d_1248_v193_ = nw246_
                d_1249_v194_: _dafny.Map
                d_1249_v194_ = _dafny.Map({d_1248_v193_: (self).f7})
                d_1250_v195_: D12
                d_1250_v195_ = D12_DC35((self).f7, len((d_1249_v194_) | (d_1249_v194_)), len(d_1153_v111_))
                d_1250_v195_ = d_1250_v195_
                (self).f6 = default__.safeModuloInt(p0, (self).f5)
                d_1251_v196_: bool
                d_1251_v196_ = False
                d_1251_v196_ = (self).f7
                d_1252_v197_: _dafny.Array
                d_1253_v198_: bool
                d_1254_v199_: _dafny.MultiSet
                d_1255_v200_: int
                out58_: _dafny.Array
                out59_: bool
                out60_: _dafny.MultiSet
                out61_: int
                out58_, out59_, out60_, out61_ = default__.m0(_dafny.CodePoint('r'), globalState)
                d_1252_v197_ = out58_
                d_1253_v198_ = out59_
                d_1254_v199_ = out60_
                d_1255_v200_ = out61_
            elif True:
                (globalState).f2 = 34
                d_1256_v201_: _dafny.Array
                nw247_ = _dafny.Array(False, 1)
                d_1256_v201_ = nw247_
                index226_ = default__.safeIndex(354, (d_1256_v201_).length(0))
                (d_1256_v201_)[index226_] = ((self).f7 if (self).f7 else True)
                d_1257_v202_: _dafny.Seq
                d_1257_v202_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1154_v112_ for d_1258_i18_ in range(default__.abs(951))]), _dafny.SeqWithoutIsStrInference([d_1154_v112_ for d_1259_i19_ in range(default__.abs(773))])])
                d_1260_v203_: _dafny.Map
                d_1260_v203_ = _dafny.Map({p0: len((d_1257_v202_)[default__.safeIndex((self).f5, len(d_1257_v202_))])})
                d_1261_v204_: _dafny.Seq
                d_1261_v204_ = _dafny.SeqWithoutIsStrInference([len(d_1260_v203_), p0])
                index227_ = default__.safeIndex(354, (d_1256_v201_).length(0))
                (d_1256_v201_)[index227_] = ((d_1261_v204_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f7])), len(d_1261_v204_))]) < ((self).f5)
                d_1262_v205_: C0
                nw248_ = C0()
                nw248_.ctor__(self.f6)
                d_1262_v205_ = nw248_
                d_1263_v206_: _dafny.Array
                nw249_ = _dafny.Array(None, 5)
                nw249_[int(0)] = d_1262_v205_
                nw249_[int(1)] = d_1262_v205_
                nw249_[int(2)] = d_1262_v205_
                nw249_[int(3)] = d_1262_v205_
                nw249_[int(4)] = d_1262_v205_
                d_1263_v206_ = nw249_
                index228_ = default__.safeIndex(966, (d_1263_v206_).length(0))
                (d_1263_v206_)[index228_] = (d_1262_v205_ if (d_1256_v201_)[default__.safeIndex(354, (d_1256_v201_).length(0))] else d_1262_v205_)
                index229_ = default__.safeIndex(966, (d_1263_v206_).length(0))
                (d_1263_v206_)[index229_] = d_1262_v205_
                d_1260_v203_ = (d_1260_v203_).set((self).f5, default__.safeModuloInt(p0, self.f6))
                d_1264_v207_: T1
                nw250_ = C1()
                nw250_.ctor__((self.f6 if (self).f7 else (self).f5), ((d_1262_v205_).f8) - ((d_1262_v205_).f8))
                d_1264_v207_ = nw250_
                d_1265_v208_: _dafny.Seq
                d_1265_v208_ = _dafny.SeqWithoutIsStrInference([(d_1256_v201_)[default__.safeIndex(354, (d_1256_v201_).length(0))], (self).f7, (self).f7])
                d_1266_v209_: C2
                nw251_ = C2()
                nw251_.ctor__(len(d_1153_v111_), len(d_1265_v208_))
                d_1266_v209_ = nw251_
                d_1267_v210_: _dafny.Seq
                d_1267_v210_ = _dafny.SeqWithoutIsStrInference([d_1266_v209_])
                index230_ = default__.safeIndex(354, (d_1256_v201_).length(0))
                rhs213_ = d_1264_v207_
                rhs214_ = (d_1266_v209_) not in (d_1267_v210_)
                rhs215_ = d_1154_v112_
                rhs216_ = d_1261_v204_
                lhs129_ = d_1256_v201_
                lhs130_ = default__.safeIndex(354, (d_1256_v201_).length(0))
                d_1264_v207_ = rhs213_
                lhs129_[lhs130_] = rhs214_
                d_1154_v112_ = rhs215_
                d_1261_v204_ = rhs216_
            d_1268_v211_: _dafny.MultiSet
            d_1268_v211_ = _dafny.MultiSet([(self).f7, (self).f7])
            d_1269_v212_: C1
            nw252_ = C1()
            nw252_.ctor__(((_dafny.MultiSet([(self).f7])) | (d_1268_v211_)).cardinality, ((0) - ((self).f5) if (self).f7 else p0))
            d_1269_v212_ = nw252_
            d_1270_v213_: T1
            nw253_ = C2()
            nw253_.ctor__((self).f5, self.f6)
            d_1270_v213_ = nw253_
            d_1270_v213_ = d_1270_v213_
            d_1271_v214_: bool
            d_1271_v214_ = True
            d_1271_v214_ = (self).fm1(globalState)
            d_1270_v213_ = d_1270_v213_

    def m3(self, globalState):
        r0: T0 = None
        d_1272_v0_: _dafny.Map
        d_1272_v0_ = _dafny.Map({(self).f7: (self).f7})
        d_1272_v0_ = (d_1272_v0_).set((self).f7, (self).f7)
        d_1273_v1_: _dafny.Map
        d_1273_v1_ = _dafny.Map({self.f6: 683})
        (self).f6 = (-594) * (len(default__.fm10((self).f5, d_1273_v1_, (self).f7, self.f6, globalState)))
        d_1274_v2_: _dafny.MultiSet
        d_1274_v2_ = _dafny.MultiSet([(self).f7])
        d_1275_v3_: _dafny.Seq
        d_1275_v3_ = _dafny.SeqWithoutIsStrInference([(self).f7])
        d_1276_v4_: _dafny.Set
        d_1276_v4_ = _dafny.Set({(self).f7, not((self).f7), (self).f7})
        d_1277_v5_: D9
        d_1277_v5_ = D9_DC28((d_1274_v2_).issubset(d_1274_v2_), (self).f5, (self).fm2(len(d_1275_v3_), d_1276_v4_, globalState))
        d_1277_v5_ = D9_DC28((self).f7, self.f6, self.f6)
        d_1278_v6_: _dafny.Set
        d_1278_v6_ = _dafny.Set({len(d_1272_v0_)})
        d_1279_v7_: _dafny.Seq
        d_1279_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exw"))
        d_1280_v8_: _dafny.MultiSet
        d_1280_v8_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1281_i0_ in range(default__.abs(-957))]), d_1279_v7_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1282_i1_ in range(default__.abs(547))]), d_1279_v7_])
        d_1283_v9_: _dafny.Map
        d_1283_v9_ = _dafny.Map({(self).f5: (self).f7})
        d_1284_v10_: _dafny.Seq
        d_1284_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(0) - (self.f6)}), (d_1278_v6_) | (_dafny.Set({(d_1280_v8_).cardinality})), (_dafny.Set({self.f6, (self).f5, len(d_1283_v9_)})).intersection(d_1278_v6_)])
        d_1285_v11_: bool
        d_1285_v11_ = True
        d_1286_v12_: _dafny.Set
        d_1286_v12_ = _dafny.Set({default__.fm18(not((self).f7), self.f6, (self).f5, d_1285_v11_, globalState)})
        d_1287_v13_: str
        d_1287_v13_ = _dafny.CodePoint('e')
        d_1288_v14_: T0
        nw254_ = C2()
        nw254_.ctor__(self.f6, len(_dafny.SeqWithoutIsStrInference([43])))
        d_1288_v14_ = nw254_
        d_1289_v15_: _dafny.Seq
        d_1289_v15_ = _dafny.SeqWithoutIsStrInference([d_1288_v14_])
        d_1290_v16_: _dafny.Map
        d_1290_v16_ = _dafny.Map({d_1287_v13_: (_dafny.MultiSet(d_1289_v15_)).cardinality})
        rhs217_ = (self).f5
        rhs218_ = d_1284_v10_
        rhs219_ = ((self).f5) < (((d_1290_v16_)[d_1287_v13_] if (d_1287_v13_) in (d_1290_v16_) else 149))
        rhs220_ = _dafny.Set({d_1287_v13_})
        rhs221_ = ((self).f7) or (((d_1272_v0_)[(self).f7] if ((self).f7) in (d_1272_v0_) else (self).f7))
        lhs131_ = globalState
        lhs131_.f2 = rhs217_
        d_1284_v10_ = rhs218_
        d_1285_v11_ = rhs219_
        d_1286_v12_ = rhs220_
        d_1285_v11_ = rhs221_
        d_1291_v17_: _dafny.Array
        nw255_ = _dafny.Array(int(0), 4)
        d_1291_v17_ = nw255_
        d_1292_v18_: D5
        d_1292_v18_ = D5_DC17(self.f6, d_1291_v17_, (default__.fm31(self.f6, False, True, globalState)).set(self.f6, (self).f7))
        source14_ = d_1292_v18_
        if source14_.is_DC17:
            d_1293___mcc_h0_ = source14_.cf27
            d_1294___mcc_h1_ = source14_.cf28
            d_1295___mcc_h2_ = source14_.cf29
            d_1296_cf29_ = d_1295___mcc_h2_
            d_1297_cf28_ = d_1294___mcc_h1_
            d_1298_cf27_ = d_1293___mcc_h0_
            (globalState).f2 = len(_dafny.Set({_dafny.Map({d_1285_v11_: d_1285_v11_}), d_1272_v0_}))
            d_1299_v19_: _dafny.Seq
            d_1299_v19_ = _dafny.SeqWithoutIsStrInference([(len(d_1279_v7_)) + (d_1298_cf27_)])
            rhs222_ = d_1285_v11_
            rhs223_ = (d_1299_v19_)[default__.safeIndex((0) - (d_1298_cf27_), len(d_1299_v19_))]
            rhs224_ = 928
            rhs225_ = (402) != (len(d_1275_v3_))
            rhs226_ = _dafny.Map({(d_1287_v13_ if d_1285_v11_ else d_1287_v13_): ((self).f5) - (self.f6)})
            lhs132_ = self
            lhs133_ = globalState
            d_1285_v11_ = rhs222_
            lhs132_.f6 = rhs223_
            lhs133_.f2 = rhs224_
            d_1285_v11_ = rhs225_
            d_1290_v16_ = rhs226_
            d_1274_v2_ = (d_1274_v2_).intersection(d_1274_v2_)
            (globalState).f2 = (self).f5
        elif source14_.is_DC18:
            d_1300___mcc_h3_ = source14_.cf30
            d_1301___mcc_h4_ = source14_.cf31
            d_1302___mcc_h5_ = source14_.cf32
            d_1303_cf32_ = d_1302___mcc_h5_
            d_1304_cf31_ = d_1301___mcc_h4_
            d_1305_cf30_ = d_1300___mcc_h3_
            (globalState).f2 = (self).f5
            if (self.f6) == (default__.safeModuloInt(135, self.f6)):
                (self).f6 = (self).f5
                d_1306_v20_: C0
                nw256_ = C0()
                nw256_.ctor__((self.f6 if (self).f7 else self.f6))
                d_1306_v20_ = nw256_
                d_1307_v21_: C2
                nw257_ = C2()
                nw257_.ctor__(868, len(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_1287_v13_ for d_1308_i2_ in range(default__.abs(107))])})))
                d_1307_v21_ = nw257_
                d_1309_v22_: _dafny.Seq
                d_1309_v22_ = _dafny.SeqWithoutIsStrInference([d_1307_v21_, d_1307_v21_, d_1307_v21_])
                d_1307_v21_ = (d_1309_v22_)[default__.safeIndex((0) - (default__.safeModuloInt(495, (d_1306_v20_).f8)), len(d_1309_v22_))]
                d_1310_v23_: _dafny.Array
                def lambda50_(d_1311_cf30_):
                    def lambda51_(d_1312_i3_):
                        return d_1311_cf30_

                    return lambda51_

                init28_ = lambda50_(d_1305_cf30_)
                nw258_ = _dafny.Array(None, 22)
                for i0_28_ in range(nw258_.length(0)):
                    nw258_[i0_28_] = init28_(i0_28_)
                d_1310_v23_ = nw258_
                index231_ = default__.safeIndex(623, (d_1310_v23_).length(0))
                (d_1310_v23_)[index231_] = d_1303_cf32_
                index232_ = default__.safeIndex(623, (d_1310_v23_).length(0))
                (d_1310_v23_)[index232_] = (d_1278_v6_).issubset(d_1278_v6_)
                d_1313_v24_: _dafny.Map
                d_1313_v24_ = _dafny.Map({not((d_1310_v23_)[default__.safeIndex(623, (d_1310_v23_).length(0))]): default__.safeDivisionInt(d_1304_cf31_, (self).fm2((d_1306_v20_).f8, d_1276_v4_, globalState))})
                d_1314_v25_: _dafny.Array
                nw259_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 3)
                d_1314_v25_ = nw259_
                index233_ = default__.safeIndex(168, (d_1314_v25_).length(0))
                (d_1314_v25_)[index233_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucngrye"))
                d_1315_v26_: _dafny.Array
                nw260_ = _dafny.Array(None, 20)
                d_1315_v26_ = nw260_
                d_1316_v27_: _dafny.Map
                d_1316_v27_ = _dafny.Map({self.f6: d_1315_v26_})
                d_1317_v28_: _dafny.Array
                nw261_ = _dafny.Array(None, 4)
                nw261_[int(0)] = ((d_1316_v27_)[d_1304_cf31_] if (d_1304_cf31_) in (d_1316_v27_) else d_1315_v26_)
                nw261_[int(1)] = d_1315_v26_
                nw261_[int(2)] = d_1315_v26_
                nw261_[int(3)] = d_1315_v26_
                d_1317_v28_ = nw261_
                index234_ = default__.safeIndex(143, (d_1317_v28_).length(0))
                (d_1317_v28_)[index234_] = d_1315_v26_
                d_1318_v29_: D1
                d_1318_v29_ = D1_DC5()
                d_1319_v30_: D5
                d_1319_v30_ = D5_DC19((d_1306_v20_).f8, d_1287_v13_, (self).f5)
                d_1320_v31_: _dafny.MultiSet
                d_1320_v31_ = _dafny.MultiSet([self.f6, 148, d_1304_cf31_])
                index235_ = default__.safeIndex(623, (d_1310_v23_).length(0))
                index236_ = default__.safeIndex(168, (d_1314_v25_).length(0))
                index237_ = default__.safeIndex(143, (d_1317_v28_).length(0))
                rhs227_ = (d_1307_v21_).fm9(d_1318_v29_, ((self).f5) >= (len(d_1278_v6_)), globalState)
                rhs228_ = (d_1319_v30_).cf34
                rhs229_ = _dafny.Map({d_1305_cf30_: (603) + (self.f6)})
                rhs230_ = default__.fm11(d_1304_cf31_, ((d_1320_v31_)[(self).f5] if ((self).f5) in (d_1320_v31_) else (d_1306_v20_).f8), default__.safeModuloInt((0) - (len(d_1272_v0_)), (d_1306_v20_).fm4(d_1287_v13_, d_1274_v2_, (self).f7, globalState)), default__.safeModuloInt(d_1304_cf31_, self.f6), globalState)
                rhs231_ = d_1315_v26_
                lhs134_ = d_1310_v23_
                lhs135_ = default__.safeIndex(623, (d_1310_v23_).length(0))
                lhs136_ = d_1314_v25_
                lhs137_ = default__.safeIndex(168, (d_1314_v25_).length(0))
                lhs138_ = d_1317_v28_
                lhs139_ = default__.safeIndex(143, (d_1317_v28_).length(0))
                lhs134_[lhs135_] = rhs227_
                d_1287_v13_ = rhs228_
                d_1313_v24_ = rhs229_
                lhs136_[lhs137_] = rhs230_
                lhs138_[lhs139_] = rhs231_
            elif True:
                d_1287_v13_ = d_1287_v13_
                d_1321_v32_: _dafny.Array
                nw262_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 24)
                d_1321_v32_ = nw262_
                index238_ = default__.safeIndex(296, (d_1321_v32_).length(0))
                (d_1321_v32_)[index238_] = (d_1279_v7_) + (d_1279_v7_)
                index239_ = default__.safeIndex(296, (d_1321_v32_).length(0))
                (d_1321_v32_)[index239_] = d_1279_v7_
                d_1322_v33_: _dafny.Array
                def lambda52_(d_1323_v7_, d_1324_v6_, d_1325_cf32_, d_1326_v2_, d_1327_v32_):
                    def lambda53_(d_1328_i4_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D7_DC23(((d_1326_v2_)[d_1325_cf32_] if (d_1325_cf32_) in (d_1326_v2_) else (self).f5), (d_1327_v32_)[default__.safeIndex(296, (d_1327_v32_).length(0))], (self).f5)]))).issubset(_dafny.MultiSet([D7_DC23(self.f6, d_1323_v7_, len(d_1324_v6_))]))

                    return lambda53_

                init29_ = lambda52_(d_1279_v7_, d_1278_v6_, d_1303_cf32_, d_1274_v2_, d_1321_v32_)
                nw263_ = _dafny.Array(None, 6)
                for i0_29_ in range(nw263_.length(0)):
                    nw263_[i0_29_] = init29_(i0_29_)
                d_1322_v33_ = nw263_
                d_1322_v33_ = d_1322_v33_
                d_1329_v34_: _dafny.Seq
                d_1329_v34_ = _dafny.SeqWithoutIsStrInference([(D1_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htofcuax")), self.f6)).cf5, (d_1321_v32_)[default__.safeIndex(296, (d_1321_v32_).length(0))], (d_1321_v32_)[default__.safeIndex(296, (d_1321_v32_).length(0))]])
                d_1330_v35_: C2
                nw264_ = C2()
                nw264_.ctor__((self).f5, len(d_1329_v34_))
                d_1330_v35_ = nw264_
                d_1331_v36_: C2
                nw265_ = C2()
                nw265_.ctor__(-392, len(d_1273_v1_))
                d_1331_v36_ = nw265_
            d_1332_v37_: _dafny.Set
            d_1332_v37_ = _dafny.Set({d_1274_v2_})
            if (d_1332_v37_).ispropersubset(d_1332_v37_):
                d_1303_cf32_ = d_1285_v11_
                d_1333_v38_: D4
                d_1333_v38_ = D4_DC13(d_1303_cf32_, (self).f5, d_1303_cf32_)
                rhs232_ = True
                rhs233_ = d_1305_cf30_
                rhs234_ = not(((self).f5) < ((d_1333_v38_).cf18))
                d_1285_v11_ = rhs232_
                d_1285_v11_ = rhs233_
                d_1305_cf30_ = rhs234_
                d_1334_v39_: D12
                d_1334_v39_ = D12_DC36(d_1285_v11_, d_1305_cf30_, d_1279_v7_)
                d_1335_v40_: T1
                nw266_ = C2()
                nw266_.ctor__(len(_dafny.Map({(self).f7: d_1304_cf31_})), len(d_1283_v9_))
                d_1335_v40_ = nw266_
                d_1283_v9_ = (d_1283_v9_).set(default__.safeDivisionInt(self.f6, len(_dafny.Map({d_1334_v39_: d_1335_v40_}))), (self).f7)
                d_1336_v41_: _dafny.Array
                nw267_ = _dafny.Array(False, 27)
                d_1336_v41_ = nw267_
                d_1337_v42_: _dafny.Map
                d_1337_v42_ = _dafny.Map({d_1335_v40_: d_1336_v41_})
                d_1337_v42_ = d_1337_v42_
                (globalState).f2 = (self).f5
            elif True:
                d_1338_v43_: _dafny.MultiSet
                d_1338_v43_ = _dafny.MultiSet([(self).f5])
                d_1339_v44_: _dafny.Array
                nw268_ = _dafny.Array(None, 9)
                nw268_[int(0)] = d_1338_v43_
                nw268_[int(1)] = d_1338_v43_
                nw268_[int(2)] = d_1338_v43_
                nw268_[int(3)] = (d_1338_v43_).set((d_1338_v43_).cardinality, default__.abs((self).f5))
                nw268_[int(4)] = (d_1338_v43_).intersection(d_1338_v43_)
                nw268_[int(5)] = _dafny.MultiSet([-390, self.f6, (self).f5])
                nw268_[int(6)] = d_1338_v43_
                nw268_[int(7)] = (d_1338_v43_) - (d_1338_v43_)
                nw268_[int(8)] = (d_1338_v43_).set((self).f5, default__.abs((self).f5))
                d_1339_v44_ = nw268_
                d_1339_v44_ = d_1339_v44_
                d_1340_v45_: C0
                nw269_ = C0()
                nw269_.ctor__((D4_DC13(d_1303_cf32_, (self).f5, d_1303_cf32_)).cf18)
                d_1340_v45_ = nw269_
                d_1341_v46_: _dafny.Seq
                d_1341_v46_ = _dafny.SeqWithoutIsStrInference([self.f6, (d_1340_v45_).f8, d_1304_cf31_, (self).f5, (d_1340_v45_).f8])
                d_1341_v46_ = d_1341_v46_
                d_1342_v47_: C1
                nw270_ = C1()
                nw270_.ctor__((d_1340_v45_).fm4(d_1287_v13_, d_1274_v2_, d_1303_cf32_, globalState), 578)
                d_1342_v47_ = nw270_
                d_1342_v47_ = d_1342_v47_
                d_1343_v48_: _dafny.Array
                nw271_ = _dafny.Array(None, 18)
                d_1343_v48_ = nw271_
                d_1344_v49_: C2
                nw272_ = C2()
                nw272_.ctor__(326, len(d_1275_v3_))
                d_1344_v49_ = nw272_
                index240_ = default__.safeIndex(260, (d_1343_v48_).length(0))
                (d_1343_v48_)[index240_] = d_1344_v49_
                index241_ = default__.safeIndex(260, (d_1343_v48_).length(0))
                rhs235_ = d_1344_v49_
                rhs236_ = self.f6
                rhs237_ = d_1342_v47_
                lhs140_ = d_1343_v48_
                lhs141_ = default__.safeIndex(260, (d_1343_v48_).length(0))
                lhs142_ = globalState
                lhs140_[lhs141_] = rhs235_
                lhs142_.f2 = rhs236_
                d_1342_v47_ = rhs237_
            d_1345_v50_: D12
            d_1345_v50_ = D12_DC36(d_1285_v11_, d_1303_cf32_, d_1279_v7_)
            d_1346_v51_: _dafny.Seq
            d_1346_v51_ = _dafny.SeqWithoutIsStrInference([d_1279_v7_, d_1279_v7_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ey"))])
            d_1347_v52_: _dafny.Array
            nw273_ = _dafny.Array(None, 18)
            nw273_[int(0)] = d_1279_v7_
            nw273_[int(1)] = (default__.fm11((self).f5, self.f6, d_1304_cf31_, self.f6, globalState)) + (d_1279_v7_)
            nw273_[int(2)] = (d_1345_v50_).cf62
            nw273_[int(3)] = (d_1279_v7_).set(default__.safeIndex(self.f6, len(d_1279_v7_)), d_1287_v13_)
            nw273_[int(4)] = d_1279_v7_
            nw273_[int(5)] = d_1279_v7_
            nw273_[int(6)] = d_1279_v7_
            nw273_[int(7)] = d_1279_v7_
            nw273_[int(8)] = (d_1279_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaddku")))
            nw273_[int(9)] = d_1279_v7_
            nw273_[int(10)] = d_1279_v7_
            nw273_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpastef"))
            nw273_[int(12)] = d_1279_v7_
            nw273_[int(13)] = d_1279_v7_
            nw273_[int(14)] = d_1279_v7_
            nw273_[int(15)] = (d_1346_v51_)[default__.safeIndex(d_1304_cf31_, len(d_1346_v51_))]
            nw273_[int(16)] = _dafny.SeqWithoutIsStrInference([d_1287_v13_ for d_1348_i5_ in range(default__.abs(5))])
            nw273_[int(17)] = d_1279_v7_
            d_1347_v52_ = nw273_
            nw274_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_1347_v52_ = nw274_
        elif source14_.is_DC19:
            d_1349___mcc_h6_ = source14_.cf33
            d_1350___mcc_h7_ = source14_.cf34
            d_1351___mcc_h8_ = source14_.cf35
            d_1352_cf35_ = d_1351___mcc_h8_
            d_1353_cf34_ = d_1350___mcc_h7_
            d_1354_cf33_ = d_1349___mcc_h6_
            d_1354_cf33_ = d_1354_cf33_
            d_1355_v53_: _dafny.Map
            d_1355_v53_ = _dafny.Map({d_1285_v11_: (d_1272_v0_) | (d_1272_v0_)})
            d_1355_v53_ = default__.fm32((self).f5, not (not(True)) or ((self).f7), globalState)
            d_1356_v54_: C0
            nw275_ = C0()
            nw275_.ctor__(d_1354_cf33_)
            d_1356_v54_ = nw275_
            d_1356_v54_ = d_1356_v54_
            d_1357_v55_: _dafny.Array
            def lambda54_(d_1358_v0_):
                def lambda55_(d_1359_i6_):
                    return _dafny.SeqWithoutIsStrInference([(D0_DC3((D4_DC14(len(d_1358_v0_), (self).f7, 675, (self).f7, self.f6)).cf21)).cf3])

                return lambda55_

            init30_ = lambda54_(d_1272_v0_)
            nw276_ = _dafny.Array(None, 2)
            for i0_30_ in range(nw276_.length(0)):
                nw276_[i0_30_] = init30_(i0_30_)
            d_1357_v55_ = nw276_
            index242_ = default__.safeIndex(223, (d_1357_v55_).length(0))
            (d_1357_v55_)[index242_] = _dafny.SeqWithoutIsStrInference([d_1285_v11_])
            d_1360_v56_: _dafny.Seq
            d_1360_v56_ = _dafny.SeqWithoutIsStrInference([d_1275_v3_, _dafny.SeqWithoutIsStrInference([True, (self).f7]), d_1275_v3_, d_1275_v3_, d_1275_v3_])
            d_1361_v57_: _dafny.Seq
            d_1361_v57_ = _dafny.SeqWithoutIsStrInference([(d_1360_v56_)[default__.safeIndex(211, len(d_1360_v56_))], d_1275_v3_])
            d_1362_v58_: _dafny.Array
            nw277_ = _dafny.Array(False, 3)
            d_1362_v58_ = nw277_
            d_1363_v59_: _dafny.MultiSet
            d_1363_v59_ = _dafny.MultiSet([d_1362_v58_])
            index243_ = default__.safeIndex(223, (d_1357_v55_).length(0))
            (d_1357_v55_)[index243_] = (((d_1360_v56_)[default__.safeIndex((d_1363_v59_).cardinality, len(d_1360_v56_))]) + (d_1275_v3_)) + (default__.fm33(globalState))
        elif True:
            d_1364___mcc_h9_ = source14_.cf26
            d_1365_cf26_ = d_1364___mcc_h9_
            d_1287_v13_ = d_1287_v13_
            d_1366_v60_: _dafny.Array
            nw278_ = _dafny.Array(_dafny.Seq({}), 18)
            d_1366_v60_ = nw278_
            index244_ = default__.safeIndex(59, (d_1366_v60_).length(0))
            (d_1366_v60_)[index244_] = _dafny.SeqWithoutIsStrInference([(self).f7])
            index245_ = default__.safeIndex(59, (d_1366_v60_).length(0))
            (d_1366_v60_)[index245_] = (d_1275_v3_) + (default__.fm33(globalState))
            d_1367_v61_: _dafny.Map
            d_1367_v61_ = _dafny.Map({self.f6: d_1279_v7_})
            d_1285_v11_ = (d_1367_v61_) != (d_1367_v61_)
            d_1368_v62_: _dafny.Seq
            d_1368_v62_ = _dafny.SeqWithoutIsStrInference([(self).f5, -235])
            d_1369_v63_: D4
            d_1369_v63_ = D4_DC11(d_1368_v62_)
            pat_let_tv17_ = d_1368_v62_
            d_1370_v64_: _dafny.Array
            nw279_ = _dafny.Array(None, 1)
            def iife113_(_pat_let29_0):
                def iife114_(d_1371_dt__update__tmp_h0_):
                    def iife115_(_pat_let30_0):
                        def iife116_(d_1372_dt__update_hcf13_h0_):
                            return D4_DC11(d_1372_dt__update_hcf13_h0_)
                        return iife116_(_pat_let30_0)
                    return iife115_(pat_let_tv17_)
                return iife114_(_pat_let29_0)
            nw279_[int(0)] = iife113_(d_1369_v63_)
            d_1370_v64_ = nw279_
            index246_ = default__.safeIndex(342, (d_1370_v64_).length(0))
            (d_1370_v64_)[index246_] = d_1369_v63_
            index247_ = default__.safeIndex(342, (d_1370_v64_).length(0))
            (d_1370_v64_)[index247_] = default__.fm27(self.f6, globalState)
        if not((d_1276_v4_).isdisjoint(_dafny.Set({d_1285_v11_, d_1285_v11_, (self).f7, not((self).f7)}))):
            d_1373_v65_: _dafny.MultiSet
            d_1373_v65_ = _dafny.MultiSet([d_1283_v9_])
            d_1374_v66_: _dafny.Seq
            d_1374_v66_ = _dafny.SeqWithoutIsStrInference([d_1373_v65_])
            if (self).fm0((d_1374_v66_)[default__.safeIndex(self.f6, len(d_1374_v66_))], (d_1274_v2_).ispropersubset(d_1274_v2_), (d_1279_v7_) + (d_1279_v7_), globalState):
                d_1375_v67_: _dafny.Map
                d_1375_v67_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1376_i7_ in range(default__.abs(58))]): default__.safeModuloInt((self).f5, (self).f5)})
                d_1377_v68_: _dafny.Array
                nw280_ = _dafny.Array(False, 11)
                d_1377_v68_ = nw280_
                d_1378_v69_: _dafny.MultiSet
                d_1378_v69_ = _dafny.MultiSet([d_1377_v68_])
                d_1379_v70_: _dafny.Map
                d_1379_v70_ = _dafny.Map({d_1288_v14_: d_1375_v67_})
                rhs238_ = ((d_1379_v70_)[d_1288_v14_] if (d_1288_v14_) in (d_1379_v70_) else (d_1375_v67_) | (d_1375_v67_))
                rhs239_ = d_1378_v69_
                d_1375_v67_ = rhs238_
                d_1378_v69_ = rhs239_
                index248_ = default__.safeIndex(950, (d_1377_v68_).length(0))
                (d_1377_v68_)[index248_] = d_1285_v11_
                index249_ = default__.safeIndex(44, (d_1377_v68_).length(0))
                (d_1377_v68_)[index249_] = (self.f6) <= (self.f6)
                d_1380_v71_: D0
                d_1380_v71_ = D0_DC3((self).f7)
                d_1381_v73_: _dafny.Map
                d_1381_v73_ = _dafny.Map({d_1285_v11_: self.f6})
                d_1382_v74_: _dafny.Seq
                d_1382_v74_ = _dafny.SeqWithoutIsStrInference([(687) - ((self).f5), default__.safeDivisionInt(self.f6, self.f6), (self.f6) + (self.f6), (self).f5, (self).fm2(self.f6, d_1276_v4_, globalState)])
                index250_ = default__.safeIndex(950, (d_1377_v68_).length(0))
                index251_ = default__.safeIndex(44, (d_1377_v68_).length(0))
                def iife117_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in _dafny.IntegerRange(-393, 743):
                        d_1383_v72_: int = compr_55_
                        if ((-393) <= (d_1383_v72_)) and ((d_1383_v72_) < (743)):
                            coll55_[default__.safeModuloInt(d_1383_v72_, 513)] = self.f6
                    return _dafny.Map(coll55_)
                rhs240_ = (d_1380_v71_).cf3
                rhs241_ = iife117_()

                rhs242_ = ((self.f6) - (((d_1381_v73_)[d_1285_v11_] if (d_1285_v11_) in (d_1381_v73_) else (self).f5))) * ((0) - (self.f6))
                rhs243_ = (d_1382_v74_)[default__.safeIndex((self).f5, len(d_1382_v74_))]
                rhs244_ = (d_1275_v3_)[default__.safeIndex(len((d_1279_v7_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esd"))).set(default__.safeIndex((self).fm2(self.f6, d_1276_v4_, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esd")))), d_1287_v13_))), len(d_1275_v3_))]
                lhs143_ = d_1377_v68_
                lhs144_ = default__.safeIndex(950, (d_1377_v68_).length(0))
                lhs145_ = globalState
                lhs146_ = globalState
                lhs147_ = d_1377_v68_
                lhs148_ = default__.safeIndex(44, (d_1377_v68_).length(0))
                lhs143_[lhs144_] = rhs240_
                d_1273_v1_ = rhs241_
                lhs145_.f2 = rhs242_
                lhs146_.f2 = rhs243_
                lhs147_[lhs148_] = rhs244_
                (globalState).f2 = (0) - (default__.safeModuloInt(((self).f5 if (d_1377_v68_)[default__.safeIndex(950, (d_1377_v68_).length(0))] else (0) - (self.f6)), (self).fm2((self).f5, d_1276_v4_, globalState)))
                d_1283_v9_ = (d_1283_v9_).set((self).f5, (self).fm0(d_1373_v65_, (self).f7, d_1279_v7_, globalState))
                d_1279_v7_ = default__.fm7(d_1287_v13_, (d_1377_v68_)[default__.safeIndex(950, (d_1377_v68_).length(0))], d_1285_v11_, (0) - (len(d_1382_v74_)), globalState)
            elif True:
                d_1384_v75_: _dafny.Array
                nw281_ = _dafny.Array(None, 13)
                d_1384_v75_ = nw281_
                d_1385_v76_: _dafny.Array
                nw282_ = _dafny.Array(None, 2)
                nw282_[int(0)] = d_1384_v75_
                nw282_[int(1)] = d_1384_v75_
                d_1385_v76_ = nw282_
                d_1386_v77_: _dafny.Array
                nw283_ = _dafny.Array(None, 1)
                nw283_[int(0)] = d_1385_v76_
                d_1386_v77_ = nw283_
                index252_ = default__.safeIndex(607, (d_1386_v77_).length(0))
                (d_1386_v77_)[index252_] = d_1385_v76_
                index253_ = default__.safeIndex(607, (d_1386_v77_).length(0))
                nw284_ = _dafny.Array(_dafny.Array(None, 0), 16)
                (d_1386_v77_)[index253_] = nw284_
                (globalState).f2 = -132
                index254_ = default__.safeIndex(266, (d_1291_v17_).length(0))
                (d_1291_v17_)[index254_] = (self.f6) - (self.f6)
                index255_ = default__.safeIndex(266, (d_1291_v17_).length(0))
                (d_1291_v17_)[index255_] = (self).f5
                d_1387_v78_: C0
                nw285_ = C0()
                nw285_.ctor__(((self).f5) + (len(d_1279_v7_)))
                d_1387_v78_ = nw285_
                d_1287_v13_ = d_1287_v13_
            if d_1285_v11_:
                d_1388_v79_: D10
                d_1388_v79_ = D10_DC30((self).f5)
                (globalState).f2 = (d_1388_v79_).cf50
                d_1285_v11_ = (D12_DC34((self).f7)).cf56
                d_1389_v80_: _dafny.Array
                def lambda56_(d_1390_v11_, d_1391_v7_):
                    def lambda57_(d_1392_i8_):
                        return _dafny.Map({_dafny.Map({d_1390_v11_: D7_DC23(self.f6, d_1391_v7_, (self).f5)}): _dafny.CodePoint('c')})

                    return lambda57_

                init31_ = lambda56_(d_1285_v11_, d_1279_v7_)
                nw286_ = _dafny.Array(None, 14)
                for i0_31_ in range(nw286_.length(0)):
                    nw286_[i0_31_] = init31_(i0_31_)
                d_1389_v80_ = nw286_
                d_1393_v81_: D12
                d_1393_v81_ = D12_DC36(((d_1272_v0_)[d_1285_v11_] if (d_1285_v11_) in (d_1272_v0_) else (self).f7), d_1285_v11_, d_1279_v7_)
                d_1394_v82_: D7
                d_1394_v82_ = D7_DC23((self).f5, d_1279_v7_, 845)
                d_1395_v83_: _dafny.Map
                d_1395_v83_ = _dafny.Map({_dafny.Map({(d_1393_v81_).cf60: d_1394_v82_}): d_1287_v13_})
                index256_ = default__.safeIndex(931, (d_1389_v80_).length(0))
                (d_1389_v80_)[index256_] = (d_1395_v83_) | (d_1395_v83_)
                d_1396_v84_: _dafny.Seq
                d_1396_v84_ = _dafny.SeqWithoutIsStrInference([(d_1275_v3_) + (d_1275_v3_), ((_dafny.SeqWithoutIsStrInference([(self).f7])).set(default__.safeIndex((self).f5, len(_dafny.SeqWithoutIsStrInference([(self).f7]))), not((self).f7))) + (d_1275_v3_), (d_1275_v3_) + ((d_1275_v3_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsnlo"))), len(d_1275_v3_)), False))])
                d_1397_v85_: _dafny.Array
                def lambda58_(d_1398_v13_):
                    def lambda59_(d_1399_i9_):
                        return (d_1398_v13_ if (self).f7 else _dafny.CodePoint('u'))

                    return lambda59_

                init32_ = lambda58_(d_1287_v13_)
                nw287_ = _dafny.Array(None, 15)
                for i0_32_ in range(nw287_.length(0)):
                    nw287_[i0_32_] = init32_(i0_32_)
                d_1397_v85_ = nw287_
                index257_ = default__.safeIndex(992, (d_1397_v85_).length(0))
                (d_1397_v85_)[index257_] = d_1287_v13_
                index258_ = default__.safeIndex(304, (d_1291_v17_).length(0))
                (d_1291_v17_)[index258_] = (self).f5
                index259_ = default__.safeIndex(931, (d_1389_v80_).length(0))
                index260_ = default__.safeIndex(992, (d_1397_v85_).length(0))
                index261_ = default__.safeIndex(304, (d_1291_v17_).length(0))
                rhs245_ = d_1287_v13_
                rhs246_ = d_1395_v83_
                rhs247_ = (d_1396_v84_).set(default__.safeIndex((d_1274_v2_).cardinality, len(d_1396_v84_)), d_1275_v3_)
                rhs248_ = default__.fm18((d_1285_v11_) or (d_1285_v11_), (self.f6) + ((0) - (len(d_1278_v6_))), (self).f5, d_1285_v11_, globalState)
                rhs249_ = default__.safeModuloInt((self.f6) + (580), default__.safeModuloInt((d_1274_v2_).cardinality, self.f6))
                lhs149_ = d_1389_v80_
                lhs150_ = default__.safeIndex(931, (d_1389_v80_).length(0))
                lhs151_ = d_1397_v85_
                lhs152_ = default__.safeIndex(992, (d_1397_v85_).length(0))
                lhs153_ = d_1291_v17_
                lhs154_ = default__.safeIndex(304, (d_1291_v17_).length(0))
                d_1287_v13_ = rhs245_
                lhs149_[lhs150_] = rhs246_
                d_1396_v84_ = rhs247_
                lhs151_[lhs152_] = rhs248_
                lhs153_[lhs154_] = rhs249_
                d_1400_v86_: _dafny.Seq
                d_1400_v86_ = _dafny.SeqWithoutIsStrInference([d_1274_v2_])
                d_1401_v87_: D0
                d_1401_v87_ = D0_DC1(not(True), (self).f7)
                d_1402_v88_: C0
                nw288_ = C0()
                nw288_.ctor__((d_1291_v17_)[default__.safeIndex(304, (d_1291_v17_).length(0))])
                d_1402_v88_ = nw288_
                d_1403_v89_: _dafny.Set
                d_1403_v89_ = _dafny.Set({d_1402_v88_})
                d_1404_v90_: _dafny.Seq
                d_1404_v90_ = _dafny.SeqWithoutIsStrInference([(d_1402_v88_).fm4(_dafny.CodePoint('m'), d_1274_v2_, (self).f7, globalState), (d_1291_v17_)[default__.safeIndex(304, (d_1291_v17_).length(0))]])
                d_1405_v91_: _dafny.Array
                nw289_ = _dafny.Array(None, 23)
                nw289_[int(0)] = (self).f7
                nw289_[int(1)] = not(not (not((self).f7)) or ((self).f7))
                nw289_[int(2)] = d_1285_v11_
                nw289_[int(3)] = not(d_1285_v11_)
                nw289_[int(4)] = d_1285_v11_
                nw289_[int(5)] = d_1285_v11_
                nw289_[int(6)] = (self).f7
                nw289_[int(7)] = d_1285_v11_
                nw289_[int(8)] = d_1285_v11_
                nw289_[int(9)] = d_1285_v11_
                nw289_[int(10)] = ((d_1400_v86_)[default__.safeIndex(776, len(d_1400_v86_))]).ispropersubset(d_1274_v2_)
                nw289_[int(11)] = (self).f7
                nw289_[int(12)] = (d_1275_v3_)[default__.safeIndex((self).f5, len(d_1275_v3_))]
                nw289_[int(13)] = (self).f7
                nw289_[int(14)] = ((self).f7) or (((d_1283_v9_)[(self).f5] if ((self).f5) in (d_1283_v9_) else d_1285_v11_))
                nw289_[int(15)] = not (False) or (d_1285_v11_)
                nw289_[int(16)] = not(True)
                nw289_[int(17)] = (d_1401_v87_).cf1
                nw289_[int(18)] = (d_1403_v89_).ispropersubset(d_1403_v89_)
                nw289_[int(19)] = ((d_1284_v10_)[default__.safeIndex(759, len(d_1284_v10_))]).issubset(default__.fm34(d_1404_v90_, globalState))
                nw289_[int(20)] = not(not((self).fm1(globalState)))
                nw289_[int(21)] = False
                nw289_[int(22)] = (self).f7
                d_1405_v91_ = nw289_
                d_1405_v91_ = d_1405_v91_
                index262_ = default__.safeIndex(304, (d_1291_v17_).length(0))
                (d_1291_v17_)[index262_] = default__.safeModuloInt(29, (d_1402_v88_).f8)
            elif True:
                (globalState).f2 = ((self).fm2((self).f5, d_1276_v4_, globalState)) * (self.f6)
                d_1406_v92_: _dafny.Array
                nw290_ = _dafny.Array(None, 7)
                nw290_[int(0)] = d_1291_v17_
                nw290_[int(1)] = d_1291_v17_
                nw290_[int(2)] = d_1291_v17_
                nw290_[int(3)] = d_1291_v17_
                nw290_[int(4)] = d_1291_v17_
                nw290_[int(5)] = d_1291_v17_
                nw290_[int(6)] = d_1291_v17_
                d_1406_v92_ = nw290_
                index263_ = default__.safeIndex(885, (d_1406_v92_).length(0))
                (d_1406_v92_)[index263_] = d_1291_v17_
                d_1407_v93_: D5
                d_1407_v93_ = D5_DC16(d_1291_v17_)
                index264_ = default__.safeIndex(885, (d_1406_v92_).length(0))
                (d_1406_v92_)[index264_] = (d_1407_v93_).cf26
                d_1285_v11_ = ((self).f7) and (((self).f5) != ((self).f5))
                d_1408_v94_: _dafny.MultiSet
                d_1408_v94_ = _dafny.MultiSet([207, (self).f5])
                d_1409_v95_: D0
                d_1409_v95_ = D0_DC2()
                d_1410_v96_: D3
                d_1410_v96_ = D3_DC9(d_1408_v94_, d_1409_v95_, self.f6)
                d_1411_v97_: _dafny.Map
                d_1411_v97_ = _dafny.Map({d_1285_v11_: d_1410_v96_})
                d_1412_v98_: D3
                d_1412_v98_ = D3_DC10(((d_1411_v97_)[d_1285_v11_] if (d_1285_v11_) in (d_1411_v97_) else d_1410_v96_))
                d_1413_v99_: _dafny.Map
                d_1413_v99_ = _dafny.Map({(self).f5: d_1410_v96_})
                d_1414_v100_: _dafny.Seq
                d_1414_v100_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                d_1415_v101_: _dafny.Set
                d_1415_v101_ = _dafny.Set({d_1412_v98_, D3_DC10(((d_1413_v99_)[len(d_1414_v100_)] if (len(d_1414_v100_)) in (d_1413_v99_) else d_1410_v96_))})
                (self).f6 = len(d_1415_v101_)
                d_1287_v13_ = d_1287_v13_
            d_1416_v102_: D5
            d_1416_v102_ = D5_DC19(self.f6, d_1287_v13_, self.f6)
            d_1416_v102_ = d_1416_v102_
            d_1417_v103_: _dafny.Map
            d_1417_v103_ = _dafny.Map({(self).f7: self.f6})
            d_1417_v103_ = (d_1417_v103_).set(not(not (d_1285_v11_) or (d_1285_v11_)), (self).f5)
            d_1418_v104_: _dafny.Array
            def lambda60_(d_1419_v11_):
                def lambda61_(d_1420_i10_):
                    return ((self).f7 if d_1419_v11_ else d_1419_v11_)

                return lambda61_

            init33_ = lambda60_(d_1285_v11_)
            nw291_ = _dafny.Array(None, 18)
            for i0_33_ in range(nw291_.length(0)):
                nw291_[i0_33_] = init33_(i0_33_)
            d_1418_v104_ = nw291_
            index265_ = default__.safeIndex(164, (d_1418_v104_).length(0))
            (d_1418_v104_)[index265_] = (self).f7
            d_1421_v105_: _dafny.Seq
            d_1421_v105_ = _dafny.SeqWithoutIsStrInference([self.f6])
            d_1422_v106_: C1
            nw292_ = C1()
            nw292_.ctor__(len(d_1279_v7_), ((self).f5 if (self).f7 else len(d_1421_v105_)))
            d_1422_v106_ = nw292_
            index266_ = default__.safeIndex(164, (d_1418_v104_).length(0))
            rhs250_ = d_1285_v11_
            rhs251_ = d_1422_v106_
            rhs252_ = d_1285_v11_
            lhs155_ = d_1418_v104_
            lhs156_ = default__.safeIndex(164, (d_1418_v104_).length(0))
            lhs155_[lhs156_] = rhs250_
            d_1422_v106_ = rhs251_
            d_1285_v11_ = rhs252_
        elif True:
            d_1423_v107_: _dafny.Array
            def lambda62_(d_1424_v11_):
                def lambda63_(d_1425_i11_):
                    return d_1424_v11_

                return lambda63_

            init34_ = lambda62_(d_1285_v11_)
            nw293_ = _dafny.Array(None, 5)
            for i0_34_ in range(nw293_.length(0)):
                nw293_[i0_34_] = init34_(i0_34_)
            d_1423_v107_ = nw293_
            d_1426_v108_: _dafny.Array
            nw294_ = _dafny.Array(None, 7)
            nw294_[int(0)] = (d_1423_v107_ if (self).f7 else d_1423_v107_)
            nw294_[int(1)] = d_1423_v107_
            nw294_[int(2)] = d_1423_v107_
            nw294_[int(3)] = d_1423_v107_
            nw294_[int(4)] = d_1423_v107_
            nw294_[int(5)] = d_1423_v107_
            nw294_[int(6)] = d_1423_v107_
            d_1426_v108_ = nw294_
            index267_ = default__.safeIndex(275, (d_1426_v108_).length(0))
            (d_1426_v108_)[index267_] = d_1423_v107_
            index268_ = default__.safeIndex(275, (d_1426_v108_).length(0))
            nw295_ = _dafny.Array(False, 11)
            (d_1426_v108_)[index268_] = nw295_
            d_1427_v109_: _dafny.Array
            d_1428_v110_: bool
            d_1429_v111_: _dafny.MultiSet
            d_1430_v112_: int
            out62_: _dafny.Array
            out63_: bool
            out64_: _dafny.MultiSet
            out65_: int
            out62_, out63_, out64_, out65_ = default__.m0(d_1287_v13_, globalState)
            d_1427_v109_ = out62_
            d_1428_v110_ = out63_
            d_1429_v111_ = out64_
            d_1430_v112_ = out65_
            d_1431_v113_: _dafny.Array
            d_1432_v114_: bool
            d_1433_v115_: _dafny.MultiSet
            d_1434_v116_: int
            out66_: _dafny.Array
            out67_: bool
            out68_: _dafny.MultiSet
            out69_: int
            out66_, out67_, out68_, out69_ = default__.m0(_dafny.CodePoint('y'), globalState)
            d_1431_v113_ = out66_
            d_1432_v114_ = out67_
            d_1433_v115_ = out68_
            d_1434_v116_ = out69_
            (globalState).f2 = (self).fm2(d_1430_v112_, d_1276_v4_, globalState)
            d_1435_v117_: C0
            nw296_ = C0()
            nw296_.ctor__((self).f5)
            d_1435_v117_ = nw296_
            d_1436_v118_: _dafny.MultiSet
            d_1436_v118_ = _dafny.MultiSet([d_1435_v117_])
            d_1437_v119_: _dafny.Map
            d_1437_v119_ = _dafny.Map({(d_1287_v13_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwwebb"))): d_1436_v118_})
            d_1438_v120_: _dafny.MultiSet
            d_1438_v120_ = _dafny.MultiSet([326, (self).f5])
            d_1439_v121_: _dafny.Seq
            d_1439_v121_ = _dafny.SeqWithoutIsStrInference([d_1438_v120_])
            d_1430_v112_ = (((d_1437_v119_)[(d_1438_v120_).issubset((d_1439_v121_)[default__.safeIndex((self).f5, len(d_1439_v121_))])] if ((d_1438_v120_).issubset((d_1439_v121_)[default__.safeIndex((self).f5, len(d_1439_v121_))])) in (d_1437_v119_) else d_1436_v118_)).cardinality
        nw297_ = C2()
        nw297_.ctor__((self.f6) + (159), (0) - ((self).f5))
        r0 = nw297_
        return r0

    @property
    def f7(self):
        return self._f7
