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
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfbosyfj")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_0_i0_ in range(default__.abs(990))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i1_ in range(default__.abs(890))]))])

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        def iife0_():
            def iife8_():
                def iife12_():
                    def iife14_():
                        coll14_ = _dafny.Map()
                        compr_14_: int
                        for compr_14_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_14_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll14_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll14_)
                    coll12_ = _dafny.Set()
                    def iife13_():
                        coll13_ = _dafny.Map()
                        compr_13_: int
                        for compr_13_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_13_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll13_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll13_)
                    compr_12_: int
                    for compr_12_ in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife13_()
                    ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})).keys.Elements:
                        d_11_v2_: int = compr_12_
                        if (d_11_v2_) in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife14_()
                        ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})):
                            coll12_ = coll12_.union(_dafny.Set([(d_11_v2_) * (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_5_i2_ in range(default__.abs(945))]))) for d_6_i1_ in range(default__.abs(751))])))]))
                    return _dafny.Set(coll12_)
                coll8_ = _dafny.Map()
                def iife9_():
                    def iife11_():
                        coll11_ = _dafny.Map()
                        compr_11_: int
                        for compr_11_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_11_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll11_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll11_)
                    coll9_ = _dafny.Set()
                    def iife10_():
                        coll10_ = _dafny.Map()
                        compr_10_: int
                        for compr_10_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_10_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll10_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll10_)
                    compr_9_: int
                    for compr_9_ in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife10_()
                    ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})).keys.Elements:
                        d_10_v2_: int = compr_9_
                        if (d_10_v2_) in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife11_()
                        ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})):
                            coll9_ = coll9_.union(_dafny.Set([(d_10_v2_) * (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_5_i2_ in range(default__.abs(945))]))) for d_6_i1_ in range(default__.abs(751))])))]))
                    return _dafny.Set(coll9_)
                compr_8_: _dafny.Set
                for compr_8_ in (_dafny.Set({iife9_()
                })).Elements:
                    d_7_v0_: _dafny.Set = compr_8_
                    if (d_7_v0_) in (_dafny.Set({iife12_()
                    })):
                        coll8_[d_7_v0_] = 90
                return _dafny.Map(coll8_)
            coll0_ = _dafny.Set()
            def iife1_():
                def iife5_():
                    def iife7_():
                        coll7_ = _dafny.Map()
                        compr_7_: int
                        for compr_7_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_7_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll7_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll7_)
                    coll5_ = _dafny.Set()
                    def iife6_():
                        coll6_ = _dafny.Map()
                        compr_6_: int
                        for compr_6_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_6_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll6_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll6_)
                    compr_5_: int
                    for compr_5_ in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife6_()
                    ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})).keys.Elements:
                        d_8_v2_: int = compr_5_
                        if (d_8_v2_) in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife7_()
                        ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})):
                            coll5_ = coll5_.union(_dafny.Set([(d_8_v2_) * (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_5_i2_ in range(default__.abs(945))]))) for d_6_i1_ in range(default__.abs(751))])))]))
                    return _dafny.Set(coll5_)
                coll1_ = _dafny.Map()
                def iife2_():
                    def iife4_():
                        coll4_ = _dafny.Map()
                        compr_4_: int
                        for compr_4_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_4_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll4_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll4_)
                    coll2_ = _dafny.Set()
                    def iife3_():
                        coll3_ = _dafny.Map()
                        compr_3_: int
                        for compr_3_ in _dafny.IntegerRange(280, 120):
                            d_3_v1_: int = compr_3_
                            if ((280) <= (d_3_v1_)) and ((d_3_v1_) < (120)):
                                coll3_[(d_3_v1_) * (-118)] = len(_dafny.Set({_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cth"))): len(_dafny.Map({340: 90}))}), _dafny.Map({len(_dafny.Set({True, True})): 622})}))
                        return _dafny.Map(coll3_)
                    compr_2_: int
                    for compr_2_ in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife3_()
                    ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})).keys.Elements:
                        d_4_v2_: int = compr_2_
                        if (d_4_v2_) in (_dafny.Map({(_dafny.MultiSet([_dafny.MultiSet([len(iife4_()
                        ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})), len(_dafny.Set({-625}))]))])])).cardinality: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-234, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpmel")))]))])).cardinality})):
                            coll2_ = coll2_.union(_dafny.Set([(d_4_v2_) * (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])) for d_5_i2_ in range(default__.abs(945))]))) for d_6_i1_ in range(default__.abs(751))])))]))
                    return _dafny.Set(coll2_)
                compr_1_: _dafny.Set
                for compr_1_ in (_dafny.Set({iife2_()
                })).Elements:
                    d_7_v0_: _dafny.Set = compr_1_
                    if (d_7_v0_) in (_dafny.Set({iife5_()
                    })):
                        coll1_[d_7_v0_] = 90
                return _dafny.Map(coll1_)
            compr_0_: int
            for compr_0_ in (_dafny.Map({len(iife1_()
            ): not(True)})).keys.Elements:
                d_9_v3_: int = compr_0_
                if (d_9_v3_) in (_dafny.Map({len(iife8_()
                ): not(True)})):
                    coll0_ = coll0_.union(_dafny.Set([(d_9_v3_) + (len(_dafny.Set({-920})))]))
            return _dafny.Set(coll0_)
        def iife15_():
            coll15_ = _dafny.Set()
            compr_15_: int
            for compr_15_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_12_i3_ in range(default__.abs(739))])): len(_dafny.Set({False}))})).keys.Elements:
                d_13_v4_: int = compr_15_
                if (d_13_v4_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_12_i3_ in range(default__.abs(739))])): len(_dafny.Set({False}))})):
                    coll15_ = coll15_.union(_dafny.Set([(d_13_v4_) * (438)]))
            return _dafny.Set(coll15_)
        return _dafny.Map({((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2_i0_ in range(default__.abs(424))])), len(_dafny.Map({len(_dafny.Map({(0) - (-981): not(False)})): D1_DC3(_dafny.SeqWithoutIsStrInference([not(True)]))}))])).cardinality) * (len(iife0_()
        )): (len(_dafny.Map({len(iife15_()
        ): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jiwdrv")))}))) - (630)})

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        source0_ = _dafny.Set({False})
        d_14___mcc_h0_ = source0_
        d_15_cf23_ = d_14___mcc_h0_
        return (_dafny.MultiSet([False])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (-621)]))) - ((_dafny.MultiSet([545])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([352, len(_dafny.SeqWithoutIsStrInference([True]))]))))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        source1_ = D14_DC32(_dafny.MultiSet([not(True)]), True, False, not(True), 542)
        if source1_.is_DC32:
            d_16___mcc_h0_ = source1_.cf56
            d_17___mcc_h1_ = source1_.cf57
            d_18___mcc_h2_ = source1_.cf58
            d_19___mcc_h3_ = source1_.cf59
            d_20___mcc_h4_ = source1_.cf60
            d_21_cf60_ = d_20___mcc_h4_
            d_22_cf59_ = d_19___mcc_h3_
            d_23_cf58_ = d_18___mcc_h2_
            d_24_cf57_ = d_17___mcc_h1_
            d_25_cf56_ = d_16___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([d_22_cf59_])) + (_dafny.SeqWithoutIsStrInference([d_23_cf58_]))
        elif True:
            d_26___mcc_h5_ = source1_.cf55
            d_27_cf55_ = d_26___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([not(True), True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        return _dafny.CodePoint('d')

    @staticmethod
    def fm15(globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False})) for d_28_i0_ in range(default__.abs(835))])) + (_dafny.SeqWithoutIsStrInference([215]))) + (_dafny.SeqWithoutIsStrInference([600 for d_29_i1_ in range(default__.abs(11))]))

    @staticmethod
    def fm18(p0, globalState):
        return _dafny.Set({not (False) or (not(False))})

    @staticmethod
    def fm19(p0, p1, globalState):
        return D0_DC0((406) == (-130))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return D2_DC9(-36, 441, (False) and (False))

    @staticmethod
    def fm21(p0, globalState):
        return (_dafny.MultiSet([False, False, not(True)])).intersection(_dafny.MultiSet([False]))

    @staticmethod
    def fm22(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvtodqg"))

    @staticmethod
    def fm23(p0, globalState):
        source2_ = D1_DC5(796)
        if source2_.is_DC4:
            d_30___mcc_h0_ = source2_.cf7
            d_31___mcc_h1_ = source2_.cf8
            d_32___mcc_h2_ = source2_.cf9
            d_33___mcc_h3_ = source2_.cf10
            d_34___mcc_h4_ = source2_.cf11
            d_35_cf11_ = d_34___mcc_h4_
            d_36_cf10_ = d_33___mcc_h3_
            d_37_cf9_ = d_32___mcc_h2_
            d_38_cf8_ = d_31___mcc_h1_
            d_39_cf7_ = d_30___mcc_h0_
            if not(True):
                return _dafny.SeqWithoutIsStrInference([d_38_cf8_])
            elif True:
                def iife16_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(-531, 718):
                        d_40_v0_: int = compr_16_
                        if ((-531) <= (d_40_v0_)) and ((d_40_v0_) < (718)):
                            coll16_[(d_40_v0_) * (d_37_cf9_)] = True
                    return _dafny.Map(coll16_)
                return _dafny.SeqWithoutIsStrInference([len(iife16_()
                )])
        elif source2_.is_DC5:
            d_41___mcc_h5_ = source2_.cf12
            d_42_cf12_ = d_41___mcc_h5_
            return (_dafny.SeqWithoutIsStrInference([d_42_cf12_, 584, (0) - (d_42_cf12_), d_42_cf12_])) + (_dafny.SeqWithoutIsStrInference([d_42_cf12_]))
        elif True:
            d_43___mcc_h6_ = source2_.cf6
            d_44_cf6_ = d_43___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([107])

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return _dafny.CodePoint('v')

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([244, 580])) <= (_dafny.SeqWithoutIsStrInference([-662, len(_dafny.SeqWithoutIsStrInference([not(False), True, True, True]))])), True])

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in (_dafny.Map({334: len(_dafny.Map({True: len(_dafny.Set({False, True}))}))})).keys.Elements:
                d_45_v0_: int = compr_17_
                if (d_45_v0_) in (_dafny.Map({334: len(_dafny.Map({True: len(_dafny.Set({False, True}))}))})):
                    coll17_[default__.safeModuloInt(d_45_v0_, len(_dafny.Set({True})))] = 969
            return _dafny.Map(coll17_)
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(-42, 599):
                d_46_v1_: int = compr_18_
                if ((-42) <= (d_46_v1_)) and ((d_46_v1_) < (599)):
                    coll18_[default__.safeModuloInt(d_46_v1_, 0)] = 335
            return _dafny.Map(coll18_)
        return (_dafny.MultiSet([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o')])): len(_dafny.Map({327: True}))}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsiwka"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjgqe")))})])).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife17_()
        , _dafny.Map({243: (_dafny.MultiSet([(D22_DC54(_dafny.CodePoint('u'), D9_DC22(False), True, True, True)).cf98, _dafny.CodePoint('x')])).cardinality}), _dafny.Map({len(_dafny.Map({True: True})): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(iife18_()
        )]))]))}), _dafny.Map({48: 789})]))) | (_dafny.MultiSet([_dafny.Map({162: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwfgndon")))})])))

    @staticmethod
    def fm27(p0, p1, globalState):
        source3_ = D22_DC54(_dafny.CodePoint('r'), D9_DC22(True), False, True, not(False))
        if source3_.is_DC54:
            d_47___mcc_h0_ = source3_.cf98
            d_48___mcc_h1_ = source3_.cf99
            d_49___mcc_h2_ = source3_.cf100
            d_50___mcc_h3_ = source3_.cf101
            d_51___mcc_h4_ = source3_.cf102
            d_52_cf102_ = d_51___mcc_h4_
            d_53_cf101_ = d_50___mcc_h3_
            d_54_cf100_ = d_49___mcc_h2_
            d_55_cf99_ = d_48___mcc_h1_
            d_56_cf98_ = d_47___mcc_h0_
            return _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqm")): d_54_cf100_})
        elif True:
            d_57___mcc_h5_ = source3_.cf97
            d_58_cf97_ = d_57___mcc_h5_
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smnurn")): not(False)})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_59_i0_ in range(default__.abs(629))]): not(True)}))

    @staticmethod
    def fm28(globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: D0
            for compr_19_ in (_dafny.SeqWithoutIsStrInference([D0_DC1(True, True, len(_dafny.Map({False: (_dafny.MultiSet([True, False])).cardinality})), not(True)), D0_DC1(not(False), not(False), 297, True), D0_DC1(False, not(False), 106, True)])).Elements:
                d_60_v0_: D0 = compr_19_
                if (d_60_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC1(True, True, len(_dafny.Map({False: (_dafny.MultiSet([True, False])).cardinality})), not(True)), D0_DC1(not(False), not(False), 297, True), D0_DC1(False, not(False), 106, True)])):
                    coll19_ = coll19_.union(_dafny.Set([d_60_v0_]))
            return _dafny.Set(coll19_)
        return (iife19_()
        ) - (_dafny.Set({D0_DC1(False, not(False), len(_dafny.Map({False: False})), False)}))

    @staticmethod
    def fm29(p0, p1, globalState):
        return ((_dafny.Map({len(_dafny.Map({False: False})): len(_dafny.SeqWithoutIsStrInference([False]))})) | (_dafny.Map({667: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jp"))))}))) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e')])): len((D10_DC24(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwtjp")), _dafny.SeqWithoutIsStrInference([_dafny.Map({False: 227})]), True)).cf44)}))

    @staticmethod
    def fm30(p0, p1, globalState):
        return ((_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojauypo")))})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))}))) | (_dafny.Map({False: -979}))

    @staticmethod
    def fm31(p0, globalState):
        return (D21_DC49(_dafny.Map({True: _dafny.Map({len(_dafny.Set({809, 745, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mf")))})): 326})}))).cf86

    @staticmethod
    def fm32(p0, p1, globalState):
        if (not(False)) and (not(False)):
            return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([True, True])}))
        elif True:
            return _dafny.Map({False: _dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm33(p0, p1, globalState):
        return D2_DC8()

    @staticmethod
    def fm34(globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in (_dafny.Set({283})).Elements:
                d_64_v0_: int = compr_20_
                if (d_64_v0_) in (_dafny.Set({283})):
                    coll20_[(d_64_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_65_i3_ in range(default__.abs(980))]))))] = (0) - ((0) - (-254))
            return _dafny.Map(coll20_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovj"))), 57]): _dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({(_dafny.MultiSet([False])).cardinality: 362}))])})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True]))]): _dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwg"))): 725}))])}))) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxbg"))) for d_61_i0_ in range(default__.abs(412))]): _dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlhmrsu"))): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False]))]))})) for d_62_i1_ in range(default__.abs(364))])})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Set({len(_dafny.Map({len(_dafny.Map({len(_dafny.Map({len(_dafny.Map({False: len(_dafny.Map({len(_dafny.Map({False: 992})): -660}))})): True})): False})): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_63_i2_ in range(default__.abs(-63))]))}))}): len(_dafny.SeqWithoutIsStrInference([True]))}))]): _dafny.SeqWithoutIsStrInference([D2_DC6(iife20_()
)])})))

    @staticmethod
    def fm35(p0, globalState):
        if True:
            return _dafny.SeqWithoutIsStrInference([D5_DC14((0) - (len(_dafny.Map({False: _dafny.Set({False, True, True})}))), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})) for d_66_i1_ in range(default__.abs(161))]))})), True) for d_67_i0_ in range(default__.abs(68))])
        elif True:
            return _dafny.SeqWithoutIsStrInference([D5_DC14(219, len(_dafny.Set({274})), False)])

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(691, 981):
                d_69_v0_: int = compr_21_
                if ((691) <= (d_69_v0_)) and ((d_69_v0_) < (981)):
                    coll21_[(d_69_v0_) - (202)] = -438
            return _dafny.Map(coll21_)
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(-193, 618):
                d_70_v1_: int = compr_22_
                if ((-193) <= (d_70_v1_)) and ((d_70_v1_) < (618)):
                    coll22_[default__.safeDivisionInt(d_70_v1_, 560)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ao")))
            return _dafny.Map(coll22_)
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({-68: -480}))]), _dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({(0) - (-248): 436})) for d_68_i0_ in range(default__.abs(-255))]), (_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({282: 895})), D2_DC6(iife21_()
), D2_DC6(iife22_()
)])) + (_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({(_dafny.MultiSet([-649])).cardinality: 384}))]))])

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: D6
            for compr_23_ in (_dafny.Map({D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrjkdbhoe"))): False})).keys.Elements:
                d_71_v0_: D6 = compr_23_
                if (d_71_v0_) in (_dafny.Map({D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrjkdbhoe"))): False})):
                    coll23_[d_71_v0_] = _dafny.CodePoint('n')
            return _dafny.Map(coll23_)
        return (_dafny.Map({D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqdccsvha"))): _dafny.CodePoint('h')})) | ((iife23_()
        ) | (_dafny.Map({D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxdmgbo"))): _dafny.CodePoint('d')})))

    @staticmethod
    def fm38(p0, p1, p2, globalState):
        return D6_DC16((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_72_i0_ in range(default__.abs(423))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_73_i1_ in range(default__.abs(135))])))

    @staticmethod
    def fm39(globalState):
        return 479

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        source4_ = D14_DC31(_dafny.SeqWithoutIsStrInference([_dafny.Set({not(True)}), _dafny.Set({False}), _dafny.Set({not(True)}), _dafny.Set({(D10_DC24(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hy")), _dafny.SeqWithoutIsStrInference([_dafny.Map({True: 855}) for d_74_i0_ in range(default__.abs(846))]), True)).cf46})]))
        if source4_.is_DC32:
            d_75___mcc_h0_ = source4_.cf56
            d_76___mcc_h1_ = source4_.cf57
            d_77___mcc_h2_ = source4_.cf58
            d_78___mcc_h3_ = source4_.cf59
            d_79___mcc_h4_ = source4_.cf60
            d_80_cf60_ = d_79___mcc_h4_
            d_81_cf59_ = d_78___mcc_h3_
            d_82_cf58_ = d_77___mcc_h2_
            d_83_cf57_ = d_76___mcc_h1_
            d_84_cf56_ = d_75___mcc_h0_
            if d_82_cf58_:
                return _dafny.CodePoint('v')
            elif True:
                return _dafny.CodePoint('i')
        elif True:
            d_85___mcc_h5_ = source4_.cf55
            d_86_cf55_ = d_85___mcc_h5_
            if True:
                return _dafny.CodePoint('r')
            elif True:
                return _dafny.CodePoint('q')

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({False, True}), (_dafny.Set({True})) - (_dafny.Set({not(not(True)), False, True, False})), (_dafny.Set({True, False, False})).intersection(_dafny.Set({(D2_DC7(True, True, -255, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))).cardinality)).cf15, True, True, not(False), False})), (_dafny.Set({True, False, True, False}) if True else _dafny.Set({not(not(True))}))])

    @staticmethod
    def fm42(p0, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dagoucki")))])) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-490]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.Map({False: _dafny.CodePoint('q')})})), 760]))))

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_87_i0_ in range(default__.abs(920))])), 385))])

    @staticmethod
    def fm45(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(-250) - (len(_dafny.SeqWithoutIsStrInference([not(True)])))])

    @staticmethod
    def fm46(p0, p1, p2, p3, globalState):
        def iife24_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in (_dafny.SeqWithoutIsStrInference([((D7_DC19(len(_dafny.Set({_dafny.CodePoint('v')})), _dafny.MultiSet([True]), True, _dafny.Map({True: 794}))).cf37).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_88_i0_ in range(default__.abs(680))]))])).Elements:
                d_89_v0_: int = compr_24_
                if (d_89_v0_) in (_dafny.SeqWithoutIsStrInference([((D7_DC19(len(_dafny.Set({_dafny.CodePoint('v')})), _dafny.MultiSet([True]), True, _dafny.Map({True: 794}))).cf37).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_88_i0_ in range(default__.abs(680))]))])):
                    coll24_[(d_89_v0_) - (423)] = 909
            return _dafny.Map(coll24_)
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(442, 367):
                d_90_v1_: int = compr_25_
                if ((442) <= (d_90_v1_)) and ((d_90_v1_) < (367)):
                    coll25_[(d_90_v1_) + (-352)] = _dafny.Map({492: False})
            return _dafny.Map(coll25_)
        def iife26_():
            coll26_ = _dafny.Map()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(812, 557):
                d_92_v2_: int = compr_26_
                if ((812) <= (d_92_v2_)) and ((d_92_v2_) < (557)):
                    coll26_[(d_92_v2_) * (894)] = True
            return _dafny.Map(coll26_)
        return ((_dafny.Map({528: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): False})})) | (_dafny.Map({117: _dafny.Map({len(_dafny.Map({(_dafny.MultiSet([False])).cardinality: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "chdvwyj"))), len(iife24_()
        ), len(_dafny.Map({False: 612}))]))).cardinality})): not(not(True))})}))) | ((iife25_()
        ) | (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([465 for d_91_i1_ in range(default__.abs(447))])): False})): iife26_()
        })))

    @staticmethod
    def fm47(p0, p1, p2, p3, globalState):
        source5_ = D16_DC37(_dafny.Map({610: True}))
        if source5_.is_DC38:
            d_93___mcc_h0_ = source5_.cf69
            d_94___mcc_h1_ = source5_.cf70
            d_95___mcc_h2_ = source5_.cf71
            d_96_cf71_ = d_95___mcc_h2_
            d_97_cf70_ = d_94___mcc_h1_
            d_98_cf69_ = d_93___mcc_h0_
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_96_cf71_]))])): d_96_cf71_})
        elif source5_.is_DC37:
            d_99___mcc_h3_ = source5_.cf68
            d_100_cf68_ = d_99___mcc_h3_
            def iife27_():
                coll27_ = _dafny.Map()
                compr_27_: int
                for compr_27_ in (d_100_cf68_).keys.Elements:
                    d_101_v0_: int = compr_27_
                    if (d_101_v0_) in (d_100_cf68_):
                        coll27_[default__.safeModuloInt(d_101_v0_, -772)] = True
                return _dafny.Map(coll27_)
            return (_dafny.Map({len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtfmeu")): False})): True})) | (iife27_()
            )
        elif True:
            d_102___mcc_h4_ = source5_.cf72
            d_103_cf72_ = d_102___mcc_h4_
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqifsibhn"))): True})

    @staticmethod
    def fm48(p0, globalState):
        return (_dafny.MultiSet([_dafny.CodePoint('o')])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(D3_DC11(_dafny.CodePoint('t'))).cf22 for d_104_i0_ in range(default__.abs(460))])))

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        def iife28_():
            coll28_ = _dafny.Set()
            compr_28_: str
            for compr_28_ in (_dafny.Map({_dafny.CodePoint('w'): 536})).keys.Elements:
                d_105_v0_: str = compr_28_
                if (d_105_v0_) in (_dafny.Map({_dafny.CodePoint('w'): 536})):
                    coll28_ = coll28_.union(_dafny.Set([d_105_v0_]))
            return _dafny.Set(coll28_)
        return ((_dafny.MultiSet([_dafny.MultiSet([len(iife28_()
        ), len(_dafny.Set({923}))])]) if True else _dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([575, 305, 46])), _dafny.MultiSet([-264, 693]), _dafny.MultiSet([632]), _dafny.MultiSet([170])]))) - (_dafny.MultiSet([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({178: not(False)}))])), _dafny.MultiSet([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([False]))}))])]))

    @staticmethod
    def fm50(globalState):
        def iife29_():
            coll29_ = _dafny.Set()
            compr_29_: int
            for compr_29_ in (_dafny.SeqWithoutIsStrInference([339, 982])).Elements:
                d_106_v0_: int = compr_29_
                if (d_106_v0_) in (_dafny.SeqWithoutIsStrInference([339, 982])):
                    def iife30_():
                        coll30_ = _dafny.Map()
                        compr_30_: int
                        for compr_30_ in _dafny.IntegerRange(-368, 100):
                            d_107_v1_: int = compr_30_
                            if ((-368) <= (d_107_v1_)) and ((d_107_v1_) < (100)):
                                coll30_[default__.safeDivisionInt(d_107_v1_, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_108_i0_ in range(default__.abs(-175))])): len(_dafny.Set({len(_dafny.Map({932: _dafny.MultiSet([733])}))}))})))] = 988
                        return _dafny.Map(coll30_)
                    coll29_ = coll29_.union(_dafny.Set([(d_106_v0_) - (len(_dafny.SeqWithoutIsStrInference([len(iife30_()
), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnb")))])))]))
            return _dafny.Set(coll29_)
        return (_dafny.MultiSet([_dafny.Set({-384})])) - ((_dafny.MultiSet([iife29_()
        ])).intersection(_dafny.MultiSet([_dafny.Set({len(_dafny.Map({True: 246}))})])))

    @staticmethod
    def fm51(p0, p1, p2, p3, globalState):
        return _dafny.Set({(0) - ((443 if False else 851)), 571, 829})

    @staticmethod
    def fm52(p0, p1, globalState):
        return (_dafny.Set({not(False)})) | (_dafny.Set({False}))

    @staticmethod
    def fm53(p0, p1, p2, p3, globalState):
        def iife31_():
            def iife33_():
                coll33_ = _dafny.Set()
                compr_33_: int
                for compr_33_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])).Elements:
                    d_111_v1_: int = compr_33_
                    if (d_111_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])):
                        coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_111_v1_, len(_dafny.Set({not(False)})))]))
                return _dafny.Set(coll33_)
            coll31_ = _dafny.Map()
            def iife32_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])).Elements:
                    d_109_v1_: int = compr_32_
                    if (d_109_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])):
                        coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_109_v1_, len(_dafny.Set({not(False)})))]))
                return _dafny.Set(coll32_)
            compr_31_: _dafny.Set
            for compr_31_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([-916])).cardinality, 671}), _dafny.Set({len(_dafny.Map({False: True}))}), iife32_()
            , _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True]))})])).Elements:
                d_110_v0_: _dafny.Set = compr_31_
                if (d_110_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([-916])).cardinality, 671}), _dafny.Set({len(_dafny.Map({False: True}))}), iife33_()
                , _dafny.Set({len(_dafny.SeqWithoutIsStrInference([True, True]))})])):
                    coll31_[d_110_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_112_i0_ in range(default__.abs(64))])))
            return _dafny.Map(coll31_)
        def iife34_():
            coll34_ = _dafny.Map()
            compr_34_: int
            for compr_34_ in _dafny.IntegerRange(874, 495):
                d_113_v2_: int = compr_34_
                if ((874) <= (d_113_v2_)) and ((d_113_v2_) < (495)):
                    coll34_[(d_113_v2_) * (-116)] = 823
            return _dafny.Map(coll34_)
        source6_ = D10_DC23(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(iife31_()
), 660]): _dafny.SeqWithoutIsStrInference([D2_DC6(iife34_()
), D2_DC6(_dafny.Map({461: 387}))])}))
        if source6_.is_DC24:
            d_114___mcc_h0_ = source6_.cf44
            d_115___mcc_h1_ = source6_.cf45
            d_116___mcc_h2_ = source6_.cf46
            d_117_cf46_ = d_116___mcc_h2_
            d_118_cf45_ = d_115___mcc_h1_
            d_119_cf44_ = d_114___mcc_h0_
            return D6_DC17(-489, len(_dafny.Map({d_117_cf46_: False})), 973, 287, 272)
        elif source6_.is_DC25:
            return D6_DC17(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tu"))), len(_dafny.Map({False: False})), -503, 26, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "grxj")))))
        elif True:
            d_120___mcc_h3_ = source6_.cf43
            d_121_cf43_ = d_120___mcc_h3_
            if False:
                def iife35_():
                    def iife37_():
                        coll37_ = _dafny.Map()
                        compr_37_: int
                        for compr_37_ in _dafny.IntegerRange(656, 744):
                            d_122_v4_: int = compr_37_
                            if ((656) <= (d_122_v4_)) and ((d_122_v4_) < (744)):
                                coll37_[(d_122_v4_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([True]))))] = False
                        return _dafny.Map(coll37_)
                    coll35_ = _dafny.Map()
                    def iife36_():
                        coll36_ = _dafny.Map()
                        compr_36_: int
                        for compr_36_ in _dafny.IntegerRange(656, 744):
                            d_122_v4_: int = compr_36_
                            if ((656) <= (d_122_v4_)) and ((d_122_v4_) < (744)):
                                coll36_[(d_122_v4_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([True]))))] = False
                        return _dafny.Map(coll36_)
                    compr_35_: _dafny.Map
                    for compr_35_ in (_dafny.SeqWithoutIsStrInference([iife36_()
                    ])).Elements:
                        d_123_v3_: _dafny.Map = compr_35_
                        if (d_123_v3_) in (_dafny.SeqWithoutIsStrInference([iife37_()
                        ])):
                            coll35_[d_123_v3_] = _dafny.Map({len(_dafny.Map({-843: not(False)})): 538})
                    return _dafny.Map(coll35_)
                return D6_DC17(len(_dafny.SeqWithoutIsStrInference([True])), -178, -617, 884, len(iife35_()
))
            elif True:
                return D6_DC17(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_124_i1_ in range(default__.abs(264))])), (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vghjbcr"))))), -987, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfqfeded"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        def iife38_():
            coll38_ = _dafny.Map()
            compr_38_: int
            for compr_38_ in _dafny.IntegerRange(-354, 574):
                d_125_v0_: int = compr_38_
                if ((-354) <= (d_125_v0_)) and ((d_125_v0_) < (574)):
                    coll38_[default__.safeModuloInt(d_125_v0_, -173)] = False
            return _dafny.Map(coll38_)
        def iife39_():
            coll39_ = _dafny.Map()
            compr_39_: int
            for compr_39_ in (_dafny.MultiSet([746, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwtt"))))])).Elements:
                d_126_v1_: int = compr_39_
                if (d_126_v1_) in (_dafny.MultiSet([746, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwtt"))))])):
                    coll39_[(d_126_v1_) - (len(_dafny.SeqWithoutIsStrInference([730])))] = _dafny.SeqWithoutIsStrInference([True])
            return _dafny.Map(coll39_)
        return D16_DC38(((0) - (len(iife38_()
))) - (443), (len(_dafny.Map({_dafny.Map({432: (0) - (len(_dafny.Set({614, (0) - (len(iife39_()
)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))})))}): _dafny.SeqWithoutIsStrInference([True, True, True])})) if True else (0) - (-784)), False)

    @staticmethod
    def fm55(globalState):
        def iife40_():
            coll40_ = _dafny.Set()
            compr_40_: int
            for compr_40_ in _dafny.IntegerRange(284, -98):
                d_127_v0_: int = compr_40_
                if ((284) <= (d_127_v0_)) and ((d_127_v0_) < (-98)):
                    coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_127_v0_, 210)]))
            return _dafny.Set(coll40_)
        def iife41_():
            coll41_ = _dafny.Set()
            compr_41_: int
            for compr_41_ in _dafny.IntegerRange(314, 577):
                d_128_v1_: int = compr_41_
                if ((314) <= (d_128_v1_)) and ((d_128_v1_) < (577)):
                    coll41_ = coll41_.union(_dafny.Set([default__.safeDivisionInt(d_128_v1_, -287)]))
            return _dafny.Set(coll41_)
        return _dafny.SeqWithoutIsStrInference([(iife40_()
        ) | (iife41_()
        )])

    @staticmethod
    def fm56(globalState):
        return D3_DC11((_dafny.CodePoint('j') if not(True) else _dafny.CodePoint('d')))

    @staticmethod
    def m0(p0, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: _dafny.Seq = _dafny.Seq({})
        r2: bool = False
        r3: _dafny.Map = _dafny.Map({})
        d_129_v0_: bool
        d_129_v0_ = True
        d_130_v1_: _dafny.Seq
        d_130_v1_ = _dafny.SeqWithoutIsStrInference([d_129_v0_])
        d_131_v2_: int
        d_131_v2_ = 762
        if (d_130_v1_)[default__.safeIndex(d_131_v2_, len(d_130_v1_))]:
            d_132_v3_: _dafny.MultiSet
            d_132_v3_ = _dafny.MultiSet([default__.fm39(globalState)])
            d_133_v4_: _dafny.Map
            d_133_v4_ = _dafny.Map({d_132_v3_: d_129_v0_})
            d_134_v5_: D0
            d_134_v5_ = D0_DC0(d_129_v0_)
            d_135_v6_: _dafny.Seq
            d_135_v6_ = _dafny.SeqWithoutIsStrInference([d_134_v5_, d_134_v5_])
            d_136_v7_: D0
            d_136_v7_ = D0_DC2((d_135_v6_)[default__.safeIndex(d_131_v2_, len(d_135_v6_))])
            d_133_v4_ = (d_133_v4_).set(default__.fm10(d_131_v2_, d_131_v2_, d_136_v7_, globalState), d_129_v0_)
            (globalState).f15 = not(d_129_v0_)
            d_137_v8_: _dafny.Array
            nw0_ = _dafny.Array(int(0), 29)
            d_137_v8_ = nw0_
            index0_ = default__.safeIndex(61, (d_137_v8_).length(0))
            (d_137_v8_)[index0_] = default__.safeDivisionInt(d_131_v2_, d_131_v2_)
            index1_ = default__.safeIndex(61, (d_137_v8_).length(0))
            (d_137_v8_)[index1_] = default__.fm39(globalState)
            (globalState).f22 = d_129_v0_
            (globalState).f12 = len(default__.fm22((d_137_v8_)[default__.safeIndex(61, (d_137_v8_).length(0))], d_129_v0_, globalState))
        elif True:
            d_138_v9_: _dafny.Map
            d_138_v9_ = _dafny.Map({d_131_v2_: True})
            d_139_v10_: _dafny.Map
            d_139_v10_ = _dafny.Map({d_129_v0_: not(d_129_v0_)})
            d_140_v11_: _dafny.Map
            d_140_v11_ = _dafny.Map({d_129_v0_: d_131_v2_})
            d_141_v12_: _dafny.Seq
            d_141_v12_ = _dafny.SeqWithoutIsStrInference([((d_140_v11_)[True] if (True) in (d_140_v11_) else d_131_v2_)])
            d_142_v13_: _dafny.Seq
            d_142_v13_ = _dafny.SeqWithoutIsStrInference([d_141_v12_, d_141_v12_, _dafny.SeqWithoutIsStrInference([len(d_130_v1_), d_131_v2_]), d_141_v12_, d_141_v12_])
            d_143_v14_: C2
            nw1_ = C2()
            nw1_.ctor__(d_142_v13_, d_129_v0_, d_129_v0_)
            d_143_v14_ = nw1_
            d_144_v15_: _dafny.Map
            d_144_v15_ = _dafny.Map({(d_129_v0_) == (((d_138_v9_)[len(d_139_v10_)] if (len(d_139_v10_)) in (d_138_v9_) else d_129_v0_)): d_143_v14_})
            d_145_v16_: _dafny.MultiSet
            d_145_v16_ = _dafny.MultiSet([d_129_v0_, d_129_v0_])
            d_146_v17_: C6
            nw2_ = C6()
            nw2_.ctor__((d_145_v16_).cardinality)
            d_146_v17_ = nw2_
            d_147_v18_: _dafny.Map
            d_147_v18_ = _dafny.Map({d_129_v0_: d_146_v17_})
            d_148_v19_: D0
            d_148_v19_ = D0_DC1(True, d_129_v0_, (d_141_v12_)[default__.safeIndex(len(d_147_v18_), len(d_141_v12_))], not(d_129_v0_))
            d_144_v15_ = (d_144_v15_).set((d_148_v19_).cf1, d_143_v14_)
            (globalState).f2 = d_129_v0_
            d_149_v20_: _dafny.Array
            def lambda0_(d_150_i0_):
                return True

            init0_ = lambda0_
            nw3_ = _dafny.Array(None, 28)
            for i0_0_ in range(nw3_.length(0)):
                nw3_[i0_0_] = init0_(i0_0_)
            d_149_v20_ = nw3_
            index2_ = default__.safeIndex(414, (d_149_v20_).length(0))
            (d_149_v20_)[index2_] = d_129_v0_
            index3_ = default__.safeIndex(414, (d_149_v20_).length(0))
            (d_149_v20_)[index3_] = False
            d_151_v21_: _dafny.Array
            nw4_ = _dafny.Array(None, 15)
            nw4_[int(0)] = d_145_v16_
            nw4_[int(1)] = d_145_v16_
            nw4_[int(2)] = (d_145_v16_) | (d_145_v16_)
            nw4_[int(3)] = _dafny.MultiSet(d_130_v1_)
            nw4_[int(4)] = (_dafny.MultiSet([False])) | ((d_145_v16_).set(d_129_v0_, default__.abs((0) - (d_131_v2_))))
            nw4_[int(5)] = d_145_v16_
            nw4_[int(6)] = _dafny.MultiSet([(d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]])
            nw4_[int(7)] = default__.fm9(True, d_146_v17_.f40, d_129_v0_, globalState)
            nw4_[int(8)] = default__.fm21(d_146_v17_.f40, globalState)
            nw4_[int(9)] = d_145_v16_
            nw4_[int(10)] = d_145_v16_
            nw4_[int(11)] = (d_145_v16_) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]])))
            nw4_[int(12)] = _dafny.MultiSet([d_129_v0_, (d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]])
            nw4_[int(13)] = (d_145_v16_) - (default__.fm9(d_129_v0_, d_131_v2_, (d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))], globalState))
            nw4_[int(14)] = d_145_v16_
            d_151_v21_ = nw4_
            index4_ = default__.safeIndex(416, (d_151_v21_).length(0))
            (d_151_v21_)[index4_] = d_145_v16_
            index5_ = default__.safeIndex(416, (d_151_v21_).length(0))
            (d_151_v21_)[index5_] = d_145_v16_
            d_152_v22_: _dafny.Seq
            d_152_v22_ = _dafny.SeqWithoutIsStrInference([p0])
            if (len((d_152_v22_) + (_dafny.SeqWithoutIsStrInference([p0])))) == (365):
                (globalState).f16 = (200) + (d_131_v2_)
                d_153_v23_: D10
                d_153_v23_ = D10_DC25()
                d_153_v23_ = D10_DC25()
                d_154_v24_: _dafny.Array
                def lambda1_(d_155_v17_, d_156_v2_, d_157_v20_, d_158_v0_):
                    def lambda2_(d_159_i1_):
                        return default__.safeDivisionInt(d_159_i1_, len(_dafny.Set({d_155_v17_.f40, (0) - (d_156_v2_), len(_dafny.Map({(d_157_v20_)[default__.safeIndex(414, (d_157_v20_).length(0))]: d_158_v0_}))})))

                    return lambda2_

                init1_ = lambda1_(d_146_v17_, d_131_v2_, d_149_v20_, d_129_v0_)
                nw5_ = _dafny.Array(None, 21)
                for i0_1_ in range(nw5_.length(0)):
                    nw5_[i0_1_] = init1_(i0_1_)
                d_154_v24_ = nw5_
                index6_ = default__.safeIndex(201, (d_154_v24_).length(0))
                (d_154_v24_)[index6_] = d_146_v17_.f40
                index7_ = default__.safeIndex(338, (d_154_v24_).length(0))
                (d_154_v24_)[index7_] = default__.fm39(globalState)
                d_160_v25_: D9
                d_160_v25_ = D9_DC22((d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))])
                d_161_v26_: _dafny.Map
                d_161_v26_ = _dafny.Map({d_130_v1_: d_149_v20_})
                d_162_v27_: D9
                d_162_v27_ = D9_DC21(d_161_v26_)
                d_163_v28_: _dafny.Seq
                d_163_v28_ = _dafny.SeqWithoutIsStrInference([d_162_v27_])
                index8_ = default__.safeIndex(201, (d_154_v24_).length(0))
                index9_ = default__.safeIndex(338, (d_154_v24_).length(0))
                rhs0_ = d_146_v17_.f40
                rhs1_ = len((d_141_v12_) + (d_141_v12_))
                rhs2_ = (d_160_v25_).cf42
                rhs3_ = len((d_163_v28_ if (d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))] else d_163_v28_))
                lhs0_ = d_146_v17_
                lhs1_ = d_154_v24_
                lhs2_ = default__.safeIndex(201, (d_154_v24_).length(0))
                lhs3_ = d_154_v24_
                lhs4_ = default__.safeIndex(338, (d_154_v24_).length(0))
                lhs0_.f40 = rhs0_
                lhs1_[lhs2_] = rhs1_
                d_129_v0_ = rhs2_
                lhs3_[lhs4_] = rhs3_
                (globalState).f2 = ((d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]) not in (d_130_v1_)
                d_138_v9_ = (d_138_v9_) | (_dafny.Map({d_146_v17_.f40: (d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]}))
            elif True:
                d_164_v29_: _dafny.Set
                d_164_v29_ = _dafny.Set({len(d_140_v11_)})
                d_165_v30_: D9
                d_165_v30_ = D9_DC22(((d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))] if (d_143_v14_).fm16(len(d_152_v22_), _dafny.SeqWithoutIsStrInference([D2_DC9(d_131_v2_, len(d_164_v29_), d_129_v0_)]), d_146_v17_.f40, d_146_v17_.f40, globalState) else d_129_v0_))
                d_165_v30_ = d_165_v30_
                r1 = d_130_v1_
                d_166_v31_: D2
                d_166_v31_ = D2_DC9(d_146_v17_.f40, d_131_v2_, ((d_138_v9_)[d_146_v17_.f40] if (d_146_v17_.f40) in (d_138_v9_) else (d_149_v20_)[default__.safeIndex(414, (d_149_v20_).length(0))]))
                d_167_v32_: _dafny.Seq
                d_167_v32_ = _dafny.SeqWithoutIsStrInference([d_166_v31_])
                d_168_v33_: _dafny.Map
                d_168_v33_ = _dafny.Map({d_131_v2_: d_146_v17_.f40})
                (globalState).f12 = (622 if (d_129_v0_ if d_129_v0_ else (d_143_v14_).fm16(d_146_v17_.f40, d_167_v32_, len(d_168_v33_), d_146_v17_.f40, globalState)) else default__.safeModuloInt((d_143_v14_).fm1(globalState), (_dafny.MultiSet([(d_146_v17_).fm2(d_129_v0_, d_131_v2_, globalState)])).cardinality))
                (globalState).f16 = (d_146_v17_).fm1(globalState)
                (globalState).f22 = (p0) in (d_152_v22_)
        d_169_v34_: _dafny.Array
        def lambda3_(d_170_v1_, d_171_v2_, d_172_v0_):
            def lambda4_(d_173_i2_):
                return (_dafny.Set({True, (d_170_v1_)[default__.safeIndex(d_171_v2_, len(d_170_v1_))]})).intersection(_dafny.Set({True, d_172_v0_}))

            return lambda4_

        init2_ = lambda3_(d_130_v1_, d_131_v2_, d_129_v0_)
        nw6_ = _dafny.Array(None, 8)
        for i0_2_ in range(nw6_.length(0)):
            nw6_[i0_2_] = init2_(i0_2_)
        d_169_v34_ = nw6_
        d_174_v35_: _dafny.Set
        d_174_v35_ = _dafny.Set({d_129_v0_})
        index10_ = default__.safeIndex(683, (d_169_v34_).length(0))
        (d_169_v34_)[index10_] = (d_174_v35_) | (d_174_v35_)
        d_175_v36_: _dafny.Seq
        d_175_v36_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_129_v0_: d_131_v2_})])
        d_176_v37_: D21
        d_176_v37_ = D21_DC52(len(_dafny.SeqWithoutIsStrInference([d_129_v0_])), d_175_v36_)
        d_177_v38_: _dafny.Map
        d_177_v38_ = _dafny.Map({d_129_v0_: d_176_v37_})
        d_178_v39_: _dafny.MultiSet
        d_178_v39_ = _dafny.MultiSet([d_177_v38_, d_177_v38_, d_177_v38_, (d_177_v38_).set(d_129_v0_, D21_DC52(d_131_v2_, d_175_v36_))])
        d_179_v40_: _dafny.Map
        d_179_v40_ = _dafny.Map({d_129_v0_: d_174_v35_})
        d_180_v41_: _dafny.Seq
        d_180_v41_ = _dafny.SeqWithoutIsStrInference([d_177_v38_])
        index11_ = default__.safeIndex(683, (d_169_v34_).length(0))
        rhs4_ = default__.fm18(((d_179_v40_)[d_129_v0_] if (d_129_v0_) in (d_179_v40_) else d_174_v35_), globalState)
        rhs5_ = (d_178_v39_).intersection(_dafny.MultiSet(d_180_v41_))
        lhs5_ = d_169_v34_
        lhs6_ = default__.safeIndex(683, (d_169_v34_).length(0))
        lhs5_[lhs6_] = rhs4_
        d_178_v39_ = rhs5_
        d_181_v42_: _dafny.Array
        nw7_ = _dafny.Array(_dafny.CodePoint('D'), 12)
        d_181_v42_ = nw7_
        d_182_v43_: _dafny.MultiSet
        d_182_v43_ = _dafny.MultiSet([d_131_v2_])
        d_183_v44_: _dafny.MultiSet
        d_183_v44_ = _dafny.MultiSet([d_182_v43_])
        rhs6_ = d_129_v0_
        rhs7_ = (0) - ((812) - ((((d_183_v44_)[d_182_v43_] if (d_182_v43_) in (d_183_v44_) else d_131_v2_)) + (len(default__.fm18(d_174_v35_, globalState)))))
        rhs8_ = (d_131_v2_) + (d_131_v2_)
        rhs9_ = d_181_v42_
        rhs10_ = p0
        lhs7_ = globalState
        lhs8_ = globalState
        lhs7_.f2 = rhs6_
        lhs8_.f12 = rhs7_
        d_131_v2_ = rhs8_
        d_181_v42_ = rhs9_
        r0 = rhs10_
        (globalState).f16 = (0) - (d_131_v2_)
        index12_ = default__.safeIndex(683, (d_169_v34_).length(0))
        (d_169_v34_)[index12_] = default__.fm52(not (d_129_v0_) or (True), (0) - ((len((d_169_v34_)[default__.safeIndex(683, (d_169_v34_).length(0))])) - (d_131_v2_)), globalState)
        hi0_ = d_131_v2_
        for d_184_i3_ in range(d_131_v2_, hi0_):
            (globalState).f2 = d_129_v0_
            d_131_v2_ = default__.safeModuloInt(d_131_v2_, (0) - (default__.safeDivisionInt((_dafny.MultiSet([d_184_i3_])).cardinality, d_131_v2_)))
            d_185_v45_: C8
            nw8_ = C8()
            nw8_.ctor__(d_182_v43_)
            d_185_v45_ = nw8_
            d_186_v46_: D25
            d_186_v46_ = D25_DC62(d_185_v45_)
            source7_ = D25_DC64(d_186_v46_)
            if source7_.is_DC63:
                d_187___mcc_h0_ = source7_.cf116
                d_188___mcc_h1_ = source7_.cf117
                d_189___mcc_h2_ = source7_.cf118
                d_190_cf118_ = d_189___mcc_h2_
                d_191_cf117_ = d_188___mcc_h1_
                d_192_cf116_ = d_187___mcc_h0_
                d_193_v47_: D15
                d_193_v47_ = D15_DC35(d_129_v0_)
                d_194_v48_: C5
                nw9_ = C5()
                nw9_.ctor__((d_193_v47_).cf65, d_191_cf117_)
                d_194_v48_ = nw9_
                d_195_v49_: _dafny.Seq
                d_195_v49_ = _dafny.SeqWithoutIsStrInference([d_194_v48_])
                d_196_v50_: _dafny.Map
                d_196_v50_ = _dafny.Map({d_195_v49_: d_129_v0_})
                d_196_v50_ = (d_196_v50_).set((d_195_v49_).set(default__.safeIndex(d_131_v2_, len(d_195_v49_)), d_194_v48_), d_129_v0_)
                (globalState).f24 = (d_129_v0_) and (d_129_v0_)
                d_197_v51_: _dafny.MultiSet
                d_197_v51_ = _dafny.MultiSet([True])
                d_198_v52_: _dafny.Map
                d_198_v52_ = _dafny.Map({(d_190_cf118_) + (d_184_i3_): d_197_v51_})
                d_198_v52_ = (d_198_v52_).set(d_190_cf118_, _dafny.MultiSet([d_191_cf117_, d_129_v0_]))
                d_199_v53_: _dafny.Map
                d_199_v53_ = _dafny.Map({d_182_v43_: d_129_v0_})
                (globalState).f15 = (d_199_v53_) == (d_199_v53_)
            elif source7_.is_DC62:
                d_200___mcc_h3_ = source7_.cf115
                d_201_cf115_ = d_200___mcc_h3_
                d_202_v54_: _dafny.Array
                nw10_ = _dafny.Array(int(0), 10)
                d_202_v54_ = nw10_
                d_203_v55_: D3
                d_203_v55_ = D3_DC10(d_202_v54_)
                d_202_v54_ = (d_203_v55_).cf21
                (globalState).f6 = (default__.fm22(d_184_i3_, d_129_v0_, globalState)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_204_i4_ in range(default__.abs(219))]))
                d_205_v56_: _dafny.Array
                def lambda5_(d_206_v0_):
                    def lambda6_(d_207_i5_):
                        return d_206_v0_

                    return lambda6_

                init3_ = lambda5_(d_129_v0_)
                nw11_ = _dafny.Array(None, 7)
                for i0_3_ in range(nw11_.length(0)):
                    nw11_[i0_3_] = init3_(i0_3_)
                d_205_v56_ = nw11_
                index13_ = default__.safeIndex(919, (d_205_v56_).length(0))
                (d_205_v56_)[index13_] = d_129_v0_
                index14_ = default__.safeIndex(919, (d_205_v56_).length(0))
                rhs11_ = d_129_v0_
                rhs12_ = d_129_v0_
                rhs13_ = d_129_v0_
                rhs14_ = d_129_v0_
                lhs9_ = d_205_v56_
                lhs10_ = default__.safeIndex(919, (d_205_v56_).length(0))
                lhs11_ = globalState
                lhs12_ = globalState
                lhs13_ = globalState
                lhs9_[lhs10_] = rhs11_
                lhs11_.f2 = rhs12_
                lhs12_.f2 = rhs13_
                lhs13_.f15 = rhs14_
                d_208_v57_: _dafny.MultiSet
                d_208_v57_ = _dafny.MultiSet([d_202_v54_, d_202_v54_])
                d_209_v58_: _dafny.Map
                d_209_v58_ = _dafny.Map({True: d_208_v57_})
                d_210_v59_: D26
                d_210_v59_ = D26_DC65(d_208_v57_)
                d_211_v60_: _dafny.Map
                d_211_v60_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmiq"))): ((d_208_v57_).set(d_202_v54_, default__.abs((((d_209_v58_)[(d_205_v56_)[default__.safeIndex(919, (d_205_v56_).length(0))]] if ((d_205_v56_)[default__.safeIndex(919, (d_205_v56_).length(0))]) in (d_209_v58_) else (d_210_v59_).cf120)).cardinality))) - (_dafny.MultiSet([d_202_v54_, d_202_v54_, d_202_v54_]))})
                d_212_v61_: _dafny.Map
                d_212_v61_ = _dafny.Map({d_203_v55_: d_131_v2_})
                d_213_v62_: C4
                nw12_ = C4()
                nw12_.ctor__(d_212_v61_, (d_205_v56_)[default__.safeIndex(919, (d_205_v56_).length(0))], d_129_v0_)
                d_213_v62_ = nw12_
                d_214_v63_: D16
                d_214_v63_ = D16_DC38((0) - (d_131_v2_), d_184_i3_, d_129_v0_)
                d_211_v60_ = (d_211_v60_).set(len((_dafny.Map({d_213_v62_: d_129_v0_})).set(d_213_v62_, (d_213_v62_).fm2(True, (d_214_v63_).cf69, globalState))), d_208_v57_)
            elif True:
                d_215___mcc_h4_ = source7_.cf119
                d_216_cf119_ = d_215___mcc_h4_
                d_217_v64_: bool
                d_218_v65_: bool
                out0_: bool
                out1_: bool
                out0_, out1_ = (d_185_v45_).m4(globalState)
                d_217_v64_ = out0_
                d_218_v65_ = out1_
                (globalState).f12 = len(default__.fm22(430, d_218_v65_, globalState))
                (globalState).f16 = d_184_i3_
                d_217_v64_ = d_217_v64_
            d_219_v66_: _dafny.Array
            def lambda7_(d_220_v2_):
                def lambda8_(d_221_i6_):
                    return default__.safeDivisionInt(d_221_i6_, d_220_v2_)

                return lambda8_

            init4_ = lambda7_(d_131_v2_)
            nw13_ = _dafny.Array(None, 12)
            for i0_4_ in range(nw13_.length(0)):
                nw13_[i0_4_] = init4_(i0_4_)
            d_219_v66_ = nw13_
            d_222_v67_: _dafny.Map
            d_222_v67_ = _dafny.Map({d_129_v0_: 668})
            index15_ = default__.safeIndex(78, (d_219_v66_).length(0))
            (d_219_v66_)[index15_] = (d_131_v2_) + (((d_222_v67_)[d_129_v0_] if (d_129_v0_) in (d_222_v67_) else d_131_v2_))
            index16_ = default__.safeIndex(78, (d_219_v66_).length(0))
            (d_219_v66_)[index16_] = ((0) - (default__.fm39(globalState))) + (default__.safeModuloInt(default__.fm39(globalState), d_131_v2_))
        r0 = p0
        r1 = default__.fm25(d_129_v0_, not (d_129_v0_) or (d_129_v0_), p0, globalState)
        r2 = (d_131_v2_) > (d_131_v2_)
        d_223_v68_: _dafny.MultiSet
        d_223_v68_ = _dafny.MultiSet([True, d_129_v0_])
        r3 = _dafny.Map({(d_223_v68_) | (d_223_v68_): not(((d_130_v1_).set(default__.safeIndex(d_131_v2_, len(d_130_v1_)), d_129_v0_)) < (d_130_v1_))})
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_224_v0_: int
        d_224_v0_ = 979
        d_225_v1_: _dafny.Seq
        d_225_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovdhmu"))
        d_226_v2_: _dafny.Set
        d_226_v2_ = _dafny.Set({d_224_v0_, (0) - (len(d_225_v1_))})
        d_227_v3_: _dafny.MultiSet
        d_227_v3_ = _dafny.MultiSet([d_225_v1_])
        d_228_v4_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.Array(None, 0), 8)
        d_228_v4_ = nw14_
        d_229_v5_: _dafny.Array
        def lambda9_(d_230_v2_):
            def lambda10_(d_231_i1_):
                return (d_231_i1_) - (len(d_230_v2_))

            return lambda10_

        init5_ = lambda9_(d_226_v2_)
        nw15_ = _dafny.Array(None, 17)
        for i0_5_ in range(nw15_.length(0)):
            nw15_[i0_5_] = init5_(i0_5_)
        d_229_v5_ = nw15_
        d_232_v6_: _dafny.Map
        d_232_v6_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_233_i0_ in range(default__.abs(434))]): d_229_v5_})
        d_234_v8_: str
        d_234_v8_ = _dafny.CodePoint('v')
        d_235_v9_: _dafny.Set
        d_235_v9_ = _dafny.Set({d_234_v8_})
        d_236_v10_: _dafny.Map
        def iife42_():
            coll42_ = _dafny.Map()
            compr_42_: str
            for compr_42_ in (d_235_v9_).Elements:
                d_237_v7_: str = compr_42_
                if (d_237_v7_) in (d_235_v9_):
                    coll42_[d_237_v7_] = -441
            return _dafny.Map(coll42_)
        d_236_v10_ = _dafny.Map({d_224_v0_: iife42_()
        })
        d_238_v11_: _dafny.Map
        d_238_v11_ = _dafny.Map({d_234_v8_: d_224_v0_})
        d_239_v12_: _dafny.Map
        d_239_v12_ = _dafny.Map({((d_236_v10_)[d_224_v0_] if (d_224_v0_) in (d_236_v10_) else d_238_v11_): d_224_v0_})
        d_240_globalState_: GlobalState
        nw16_ = GlobalState()
        nw16_.ctor__(d_226_v2_, d_227_v3_, True, True, _dafny.CodePoint('l'), d_228_v4_, d_225_v1_, d_232_v6_, False, True, 139, True, 66, -610, d_229_v5_, False, 796, True, True, _dafny.CodePoint('w'), d_225_v1_, 942, True, 735, False, d_239_v12_)
        d_240_globalState_ = nw16_
        d_241_v14_: bool
        d_241_v14_ = False
        d_242_v15_: _dafny.MultiSet
        d_242_v15_ = _dafny.MultiSet([_dafny.Map({d_241_v14_: d_224_v0_})])
        d_243_v16_: _dafny.Map
        d_243_v16_ = _dafny.Map({False: d_224_v0_})
        d_244_v17_: _dafny.Map
        d_244_v17_ = _dafny.Map({d_243_v16_: d_241_v14_})
        def iife43_():
            coll43_ = _dafny.Map()
            compr_43_: _dafny.Map
            for compr_43_ in (d_242_v15_).Elements:
                d_245_v13_: _dafny.Map = compr_43_
                if (d_245_v13_) in (d_242_v15_):
                    coll43_[d_245_v13_] = (D0_DC1(d_241_v14_, d_241_v14_, d_224_v0_, d_241_v14_)).cf2
            return _dafny.Map(coll43_)
        (d_240_globalState_).f12 = len((iife43_()
        ) | (d_244_v17_))
        if (d_227_v3_).ispropersubset(default__.fm0(False, d_240_globalState_)):
            d_246_v18_: str
            d_247_v19_: _dafny.Seq
            d_248_v20_: bool
            d_249_v21_: _dafny.Map
            out2_: str
            out3_: _dafny.Seq
            out4_: bool
            out5_: _dafny.Map
            out2_, out3_, out4_, out5_ = default__.m0(d_234_v8_, d_240_globalState_)
            d_246_v18_ = out2_
            d_247_v19_ = out3_
            d_248_v20_ = out4_
            d_249_v21_ = out5_
            d_250_v22_: D0
            d_250_v22_ = D0_DC0(d_248_v20_)
            d_250_v22_ = d_250_v22_
            (d_240_globalState_).f12 = d_224_v0_
            d_251_v23_: _dafny.MultiSet
            d_251_v23_ = _dafny.MultiSet([882])
            d_252_v24_: _dafny.Array
            nw17_ = _dafny.Array(None, 2)
            nw17_[int(0)] = _dafny.MultiSet([d_224_v0_])
            nw17_[int(1)] = d_251_v23_
            d_252_v24_ = nw17_
            d_253_v25_: D1
            d_253_v25_ = D1_DC3(d_247_v19_)
            rhs15_ = d_241_v14_
            rhs16_ = (d_253_v25_).cf6
            rhs17_ = d_252_v24_
            rhs18_ = d_246_v18_
            lhs14_ = d_240_globalState_
            lhs14_.f2 = rhs15_
            d_247_v19_ = rhs16_
            d_252_v24_ = rhs17_
            d_246_v18_ = rhs18_
            d_254_v26_: str
            d_255_v27_: _dafny.Seq
            d_256_v28_: bool
            d_257_v29_: _dafny.Map
            out6_: str
            out7_: _dafny.Seq
            out8_: bool
            out9_: _dafny.Map
            out6_, out7_, out8_, out9_ = default__.m0(d_234_v8_, d_240_globalState_)
            d_254_v26_ = out6_
            d_255_v27_ = out7_
            d_256_v28_ = out8_
            d_257_v29_ = out9_
        elif True:
            (d_240_globalState_).f15 = (d_226_v2_).isdisjoint(d_226_v2_)
            (d_240_globalState_).f24 = (d_241_v14_) or (d_241_v14_)
            d_258_v30_: C3
            nw18_ = C3()
            nw18_.ctor__(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_224_v0_])]), True, (d_241_v14_) and (d_241_v14_))
            d_258_v30_ = nw18_
            d_259_v31_: _dafny.Map
            d_259_v31_ = _dafny.Map({d_241_v14_: True})
            d_260_v32_: C1
            nw19_ = C1()
            nw19_.ctor__(d_224_v0_, len(d_259_v31_), d_241_v14_, False)
            d_260_v32_ = nw19_
            d_261_v33_: _dafny.Map
            d_261_v33_ = _dafny.Map({d_260_v32_: d_241_v14_})
            d_262_v34_: _dafny.Seq
            d_262_v34_ = _dafny.SeqWithoutIsStrInference([d_261_v33_])
            d_263_v35_: _dafny.MultiSet
            d_263_v35_ = _dafny.MultiSet([len(d_225_v1_)])
            d_264_v36_: T3
            nw20_ = C8()
            nw20_.ctor__(d_263_v35_)
            d_264_v36_ = nw20_
            d_265_v37_: _dafny.MultiSet
            d_265_v37_ = _dafny.MultiSet([d_264_v36_, d_264_v36_, d_264_v36_, d_264_v36_])
            d_266_v38_: _dafny.Map
            d_266_v38_ = _dafny.Map({(d_224_v0_) + (len((d_262_v34_)[default__.safeIndex((d_260_v32_).f36, len(d_262_v34_))])): (d_265_v37_).cardinality})
            d_266_v38_ = (d_266_v38_).set(d_224_v0_, (d_224_v0_) + (d_224_v0_))
            d_267_v39_: D3
            d_267_v39_ = D3_DC11(_dafny.CodePoint('b'))
            d_268_v40_: _dafny.Map
            d_268_v40_ = _dafny.Map({d_267_v39_: d_241_v14_})
            d_268_v40_ = (d_268_v40_).set(default__.fm56(d_240_globalState_), d_241_v14_)
        d_269_v41_: _dafny.Array
        def lambda11_(d_270_v14_):
            def lambda12_(d_271_i2_):
                return d_270_v14_

            return lambda12_

        init6_ = lambda11_(d_241_v14_)
        nw21_ = _dafny.Array(None, 4)
        for i0_6_ in range(nw21_.length(0)):
            nw21_[i0_6_] = init6_(i0_6_)
        d_269_v41_ = nw21_
        index17_ = default__.safeIndex(434, (d_269_v41_).length(0))
        (d_269_v41_)[index17_] = d_241_v14_
        index18_ = default__.safeIndex(434, (d_269_v41_).length(0))
        (d_269_v41_)[index18_] = d_241_v14_
        d_272_v42_: _dafny.Seq
        d_272_v42_ = _dafny.SeqWithoutIsStrInference([d_224_v0_, d_224_v0_, d_224_v0_, d_224_v0_])
        d_273_v43_: _dafny.Map
        d_273_v43_ = _dafny.Map({d_269_v41_: d_241_v14_})
        d_274_v44_: _dafny.Map
        d_274_v44_ = _dafny.Map({(d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]: (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]})
        d_275_v45_: _dafny.Seq
        d_275_v45_ = _dafny.SeqWithoutIsStrInference([d_241_v14_])
        d_276_v46_: _dafny.MultiSet
        d_276_v46_ = _dafny.MultiSet([len(d_275_v45_), d_224_v0_, d_224_v0_])
        d_277_v47_: T2
        nw22_ = C9()
        nw22_.ctor__(d_241_v14_, len(d_274_v44_), not((d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]), not(d_241_v14_), d_276_v46_)
        d_277_v47_ = nw22_
        d_278_v48_: _dafny.Seq
        d_278_v48_ = _dafny.SeqWithoutIsStrInference([d_277_v47_])
        d_279_v49_: _dafny.MultiSet
        d_279_v49_ = _dafny.MultiSet([len(d_273_v43_), d_224_v0_, len(d_278_v48_)])
        d_280_v50_: C8
        nw23_ = C8()
        nw23_.ctor__(_dafny.MultiSet([len((d_272_v42_).set(default__.safeIndex((d_279_v49_).cardinality, len(d_272_v42_)), -790))]))
        d_280_v50_ = nw23_
        d_281_v51_: _dafny.Array
        nw24_ = _dafny.Array(None, 2)
        nw24_[int(0)] = d_280_v50_
        nw24_[int(1)] = d_280_v50_
        d_281_v51_ = nw24_
        index19_ = default__.safeIndex(758, (d_281_v51_).length(0))
        (d_281_v51_)[index19_] = d_280_v50_
        index20_ = default__.safeIndex(758, (d_281_v51_).length(0))
        (d_281_v51_)[index20_] = d_280_v50_
        d_282_v52_: _dafny.Map
        d_282_v52_ = _dafny.Map({d_224_v0_: d_224_v0_})
        d_283_v55_: _dafny.Array
        nw25_ = _dafny.Array(None, 23)
        nw25_[int(0)] = d_282_v52_
        nw25_[int(1)] = d_282_v52_
        nw25_[int(2)] = d_282_v52_
        nw25_[int(3)] = d_282_v52_
        nw25_[int(4)] = _dafny.Map({len(_dafny.Map({d_241_v14_: d_241_v14_})): d_224_v0_})
        nw25_[int(5)] = d_282_v52_
        nw25_[int(6)] = d_282_v52_
        def iife44_():
            coll44_ = _dafny.Map()
            compr_44_: int
            for compr_44_ in (d_272_v42_).Elements:
                d_284_v53_: int = compr_44_
                if (d_284_v53_) in (d_272_v42_):
                    coll44_[(d_284_v53_) + (((d_282_v52_)[165] if (165) in (d_282_v52_) else d_224_v0_))] = d_224_v0_
            return _dafny.Map(coll44_)
        nw25_[int(7)] = iife44_()
        
        nw25_[int(8)] = (d_282_v52_).set(len(d_225_v1_), len(_dafny.SeqWithoutIsStrInference([(d_277_v47_).f27, d_241_v14_])))
        nw25_[int(9)] = d_282_v52_
        nw25_[int(10)] = (d_282_v52_).set(d_224_v0_, len(d_225_v1_))
        def iife45_():
            coll45_ = _dafny.Map()
            compr_45_: int
            for compr_45_ in (d_279_v49_).Elements:
                d_285_v54_: int = compr_45_
                if (d_285_v54_) in (d_279_v49_):
                    coll45_[(d_285_v54_) * (len(d_243_v16_))] = d_224_v0_
            return _dafny.Map(coll45_)
        nw25_[int(11)] = iife45_()
        
        nw25_[int(12)] = d_282_v52_
        nw25_[int(13)] = d_282_v52_
        nw25_[int(14)] = d_282_v52_
        nw25_[int(15)] = d_282_v52_
        nw25_[int(16)] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etqlwqsyc"))): len(d_275_v45_)})
        nw25_[int(17)] = d_282_v52_
        nw25_[int(18)] = d_282_v52_
        nw25_[int(19)] = d_282_v52_
        nw25_[int(20)] = d_282_v52_
        nw25_[int(21)] = d_282_v52_
        nw25_[int(22)] = d_282_v52_
        d_283_v55_ = nw25_
        d_286_v56_: D6
        d_286_v56_ = D6_DC16(d_225_v1_)
        d_287_v57_: _dafny.Map
        d_287_v57_ = _dafny.Map({d_283_v55_: len(((d_286_v56_).cf29 if (d_277_v47_).f27 else d_225_v1_))})
        d_224_v0_ = ((d_287_v57_)[d_283_v55_] if (d_283_v55_) in (d_287_v57_) else d_224_v0_)
        d_288_v58_: T0
        nw26_ = C6()
        nw26_.ctor__(d_224_v0_)
        d_288_v58_ = nw26_
        d_289_v59_: str
        d_290_v60_: int
        out10_: str
        out11_: int
        out10_, out11_ = (d_277_v47_).m2((d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))], d_288_v58_, (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))], d_240_globalState_)
        d_289_v59_ = out10_
        d_290_v60_ = out11_
        d_291_v61_: str
        d_292_v62_: _dafny.Seq
        d_293_v63_: bool
        d_294_v64_: _dafny.Map
        out12_: str
        out13_: _dafny.Seq
        out14_: bool
        out15_: _dafny.Map
        out12_, out13_, out14_, out15_ = default__.m0(d_289_v59_, d_240_globalState_)
        d_291_v61_ = out12_
        d_292_v62_ = out13_
        d_293_v63_ = out14_
        d_294_v64_ = out15_
        d_295_v65_: _dafny.MultiSet
        d_295_v65_ = _dafny.MultiSet([(d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]])
        if ((_dafny.MultiSet([(d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]])) - (d_295_v65_)).isdisjoint((d_295_v65_).set(d_293_v63_, default__.abs(d_290_v60_))):
            d_296_v66_: str
            d_297_v67_: int
            out16_: str
            out17_: int
            out16_, out17_ = (d_277_v47_).m2((d_277_v47_).f27, d_288_v58_, (d_277_v47_).f27, d_240_globalState_)
            d_296_v66_ = out16_
            d_297_v67_ = out17_
            index21_ = default__.safeIndex(49, (d_229_v5_).length(0))
            (d_229_v5_)[index21_] = d_224_v0_
            d_298_v68_: _dafny.MultiSet
            d_298_v68_ = _dafny.MultiSet([d_272_v42_])
            d_299_v69_: C3
            nw27_ = C3()
            nw27_.ctor__((_dafny.MultiSet([d_272_v42_, _dafny.SeqWithoutIsStrInference([d_297_v67_])]) if True else d_298_v68_), (d_288_v58_).fm2((d_277_v47_).f27, d_224_v0_, d_240_globalState_), (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))])
            d_299_v69_ = nw27_
            index22_ = default__.safeIndex(49, (d_229_v5_).length(0))
            rhs19_ = (d_297_v67_) * ((d_224_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhcidbx")))))
            rhs20_ = d_299_v69_
            lhs15_ = d_229_v5_
            lhs16_ = default__.safeIndex(49, (d_229_v5_).length(0))
            lhs15_[lhs16_] = rhs19_
            d_299_v69_ = rhs20_
            d_274_v44_ = (d_274_v44_).set(((d_272_v42_)[default__.safeIndex(d_224_v0_, len(d_272_v42_))]) == ((0) - ((d_229_v5_)[default__.safeIndex(49, (d_229_v5_).length(0))])), (d_277_v47_).f26)
            d_300_v70_: _dafny.Seq
            d_300_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oujjykne")), d_225_v1_])
            (d_299_v69_).m1((d_277_v47_).f26, (d_300_v70_).set(default__.safeIndex((0) - (d_224_v0_), len(d_300_v70_)), d_225_v1_), default__.fm25((d_277_v47_).f26, (d_277_v47_).f26, d_296_v66_, d_240_globalState_), d_240_globalState_)
            d_301_v71_: _dafny.Map
            d_301_v71_ = _dafny.Map({370: d_269_v41_})
            d_301_v71_ = (d_301_v71_).set((d_229_v5_)[default__.safeIndex(49, (d_229_v5_).length(0))], d_269_v41_)
        elif True:
            (d_240_globalState_).f2 = not ((d_277_v47_).f26) or ((d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))])
            d_302_v72_: _dafny.Map
            d_302_v72_ = _dafny.Map({d_290_v60_: d_241_v14_})
            d_303_v73_: _dafny.Map
            d_303_v73_ = _dafny.Map({(d_277_v47_).f26: d_302_v72_})
            d_304_v75_: _dafny.Set
            d_304_v75_ = _dafny.Set({(d_277_v47_).f26, (d_277_v47_).f26})
            d_305_v76_: _dafny.Map
            d_305_v76_ = _dafny.Map({d_304_v75_: (d_277_v47_).f26})
            d_306_v78_: _dafny.Array
            nw28_ = _dafny.Array(None, 12)
            nw28_[int(0)] = d_302_v72_
            nw28_[int(1)] = d_302_v72_
            nw28_[int(2)] = _dafny.Map({d_224_v0_: (d_277_v47_).f27})
            nw28_[int(3)] = (_dafny.Map({len(d_225_v1_): (d_277_v47_).f26})) | (_dafny.Map({len(d_272_v42_): (d_277_v47_).fm2((d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))], d_290_v60_, d_240_globalState_)}))
            nw28_[int(4)] = ((d_303_v73_)[d_293_v63_] if (d_293_v63_) in (d_303_v73_) else d_302_v72_)
            def iife46_():
                coll46_ = _dafny.Map()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(100, 8):
                    d_307_v74_: int = compr_46_
                    if ((100) <= (d_307_v74_)) and ((d_307_v74_) < (8)):
                        coll46_[(d_307_v74_) + (d_290_v60_)] = (d_277_v47_).f27
                return _dafny.Map(coll46_)
            nw28_[int(5)] = iife46_()
            
            nw28_[int(6)] = d_302_v72_
            nw28_[int(7)] = d_302_v72_
            nw28_[int(8)] = d_302_v72_
            nw28_[int(9)] = d_302_v72_
            nw28_[int(10)] = (_dafny.Map({d_224_v0_: (d_277_v47_).f26})).set(len(d_305_v76_), (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))])
            def iife47_():
                coll47_ = _dafny.Map()
                compr_47_: int
                for compr_47_ in (d_279_v49_).Elements:
                    d_308_v77_: int = compr_47_
                    if (d_308_v77_) in (d_279_v49_):
                        coll47_[default__.safeModuloInt(d_308_v77_, d_290_v60_)] = (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]
                return _dafny.Map(coll47_)
            nw28_[int(11)] = iife47_()
            
            d_306_v78_ = nw28_
            index23_ = default__.safeIndex(490, (d_306_v78_).length(0))
            (d_306_v78_)[index23_] = d_302_v72_
            index24_ = default__.safeIndex(490, (d_306_v78_).length(0))
            (d_306_v78_)[index24_] = (d_302_v72_) | (d_302_v72_)
            (d_277_v47_).m1((d_277_v47_).f27, _dafny.SeqWithoutIsStrInference([d_225_v1_, d_225_v1_]), d_292_v62_, d_240_globalState_)
            index25_ = default__.safeIndex(742, (d_229_v5_).length(0))
            (d_229_v5_)[index25_] = (d_224_v0_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rodtn"))))
            index26_ = default__.safeIndex(742, (d_229_v5_).length(0))
            (d_229_v5_)[index26_] = (d_224_v0_) + (d_224_v0_)
            d_309_v79_: C5
            nw29_ = C5()
            nw29_.ctor__((d_277_v47_).f27, (d_288_v58_).fm2(True, 144, d_240_globalState_))
            d_309_v79_ = nw29_
            d_310_v80_: D25
            d_310_v80_ = D25_DC62((d_281_v51_)[default__.safeIndex(758, (d_281_v51_).length(0))])
            d_311_v81_: _dafny.Array
            nw30_ = _dafny.Array(int(0), 18)
            d_311_v81_ = nw30_
            d_312_v82_: D3
            d_312_v82_ = D3_DC10(d_311_v81_)
            d_313_v83_: _dafny.Map
            d_313_v83_ = _dafny.Map({d_312_v82_: d_224_v0_})
            d_314_v84_: T1
            nw31_ = C4()
            nw31_.ctor__(d_313_v83_, (d_277_v47_).f27, (d_277_v47_).f27)
            d_314_v84_ = nw31_
            d_315_v85_: D24
            d_315_v85_ = D24_DC60(d_314_v84_, len(d_225_v1_), len(_dafny.Set({len(d_292_v62_)})), (d_314_v84_).f27)
            index27_ = default__.safeIndex(742, (d_229_v5_).length(0))
            index28_ = default__.safeIndex(758, (d_281_v51_).length(0))
            rhs21_ = (-755) - (8)
            rhs22_ = d_309_v79_
            rhs23_ = (d_310_v80_).cf115
            rhs24_ = d_292_v62_
            rhs25_ = default__.safeDivisionInt(d_224_v0_, (d_315_v85_).cf112)
            lhs17_ = d_229_v5_
            lhs18_ = default__.safeIndex(742, (d_229_v5_).length(0))
            lhs19_ = d_281_v51_
            lhs20_ = default__.safeIndex(758, (d_281_v51_).length(0))
            lhs17_[lhs18_] = rhs21_
            d_309_v79_ = rhs22_
            lhs19_[lhs20_] = rhs23_
            d_292_v62_ = rhs24_
            d_224_v0_ = rhs25_
        (d_240_globalState_).f2 = (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]
        d_316_v86_: _dafny.Array
        nw32_ = _dafny.Array(None, 27)
        d_316_v86_ = nw32_
        d_317_v87_: C5
        nw33_ = C5()
        nw33_.ctor__(d_293_v63_, d_241_v14_)
        d_317_v87_ = nw33_
        index29_ = default__.safeIndex(461, (d_316_v86_).length(0))
        (d_316_v86_)[index29_] = d_317_v87_
        d_318_v88_: _dafny.Seq
        d_318_v88_ = _dafny.SeqWithoutIsStrInference([d_288_v58_, d_288_v58_])
        index30_ = default__.safeIndex(461, (d_316_v86_).length(0))
        rhs26_ = d_289_v59_
        rhs27_ = d_224_v0_
        rhs28_ = d_317_v87_
        rhs29_ = not(d_241_v14_)
        rhs30_ = (_dafny.SeqWithoutIsStrInference([d_288_v58_, d_288_v58_, d_288_v58_, d_288_v58_, d_288_v58_])) <= ((d_318_v88_) + (d_318_v88_))
        lhs21_ = d_240_globalState_
        lhs22_ = d_240_globalState_
        lhs23_ = d_316_v86_
        lhs24_ = default__.safeIndex(461, (d_316_v86_).length(0))
        lhs25_ = d_240_globalState_
        lhs26_ = d_240_globalState_
        lhs21_.f19 = rhs26_
        lhs22_.f12 = rhs27_
        lhs23_[lhs24_] = rhs28_
        lhs25_.f15 = rhs29_
        lhs26_.f22 = rhs30_
        d_319_v89_: _dafny.Seq
        d_319_v89_ = _dafny.SeqWithoutIsStrInference([d_229_v5_, d_229_v5_, d_229_v5_])
        d_229_v5_ = (d_319_v89_)[default__.safeIndex(len(d_243_v16_), len(d_319_v89_))]
        source8_ = D7_DC18(d_277_v47_)
        if source8_.is_DC19:
            d_320___mcc_h0_ = source8_.cf36
            d_321___mcc_h1_ = source8_.cf37
            d_322___mcc_h2_ = source8_.cf38
            d_323___mcc_h3_ = source8_.cf39
            d_324_cf39_ = d_323___mcc_h3_
            d_325_cf38_ = d_322___mcc_h2_
            d_326_cf37_ = d_321___mcc_h1_
            d_327_cf36_ = d_320___mcc_h0_
            index31_ = default__.safeIndex(434, (d_269_v41_).length(0))
            (d_269_v41_)[index31_] = (d_288_v58_).fm2((d_277_v47_).f26, d_327_cf36_, d_240_globalState_)
            (d_317_v87_).m7((d_276_v46_).issubset(_dafny.MultiSet(d_272_v42_)), (d_277_v47_).f26, (0) - (default__.safeDivisionInt(d_327_cf36_, d_290_v60_)), d_276_v46_, d_240_globalState_)
            d_229_v5_ = d_229_v5_
            d_328_v90_: _dafny.Map
            d_328_v90_ = _dafny.Map({d_229_v5_: d_229_v5_})
            d_329_v91_: _dafny.Map
            d_329_v91_ = _dafny.Map({d_328_v90_: d_290_v60_})
            d_329_v91_ = (d_329_v91_).set(d_328_v90_, d_224_v0_)
        elif True:
            d_330___mcc_h4_ = source8_.cf35
            d_331_cf35_ = d_330___mcc_h4_
            d_332_v92_: _dafny.Map
            d_332_v92_ = _dafny.Map({d_225_v1_: len(d_225_v1_)})
            d_333_v93_: _dafny.Map
            d_333_v93_ = _dafny.Map({len(d_282_v52_): (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]})
            (d_240_globalState_).f16 = ((d_332_v92_)[_dafny.SeqWithoutIsStrInference([d_234_v8_ for d_334_i3_ in range(default__.abs(-115))])] if (_dafny.SeqWithoutIsStrInference([d_234_v8_ for d_335_i3_ in range(default__.abs(-115))])) in (d_332_v92_) else len(d_333_v93_))
            if (_dafny.MultiSet((d_225_v1_) + (d_225_v1_))) != (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_291_v61_]))):
                d_336_v94_: D0
                d_336_v94_ = D0_DC0((d_331_cf35_).f26)
                (d_240_globalState_).f16 = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(753, (d_277_v47_).fm4(d_336_v94_, d_224_v0_, d_224_v0_, d_240_globalState_)), 123))
                (d_240_globalState_).f24 = (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]
                index32_ = default__.safeIndex(434, (d_269_v41_).length(0))
                (d_269_v41_)[index32_] = ((d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))] if (d_331_cf35_).f26 else (_dafny.CodePoint('s')) not in (d_225_v1_))
                index33_ = default__.safeIndex(434, (d_269_v41_).length(0))
                (d_269_v41_)[index33_] = (d_331_cf35_).f27
                pat_let_tv0_ = d_277_v47_
                d_337_v95_: bool
                out18_: bool
                def iife48_(_pat_let0_0):
                    def iife49_(d_338_dt__update__tmp_h0_):
                        def iife50_(_pat_let1_0):
                            def iife51_(d_339_dt__update_hcf0_h0_):
                                return D0_DC0(d_339_dt__update_hcf0_h0_)
                            return iife51_(_pat_let1_0)
                        return iife50_((pat_let_tv0_).f26)
                    return iife49_(_pat_let0_0)
                out18_ = (d_317_v87_).m8(iife48_(d_336_v94_), (d_331_cf35_).f27, d_295_v65_, (d_224_v0_) * (d_290_v60_), d_240_globalState_)
                d_337_v95_ = out18_
            elif True:
                d_340_v96_: _dafny.Map
                d_340_v96_ = _dafny.Map({(d_331_cf35_).f27: d_275_v45_})
                d_340_v96_ = (d_340_v96_).set((d_277_v47_).f27, d_275_v45_)
                d_341_v97_: _dafny.Seq
                d_341_v97_ = _dafny.SeqWithoutIsStrInference([d_225_v1_, d_225_v1_, d_225_v1_, d_225_v1_])
                d_342_v98_: C9
                nw34_ = C9()
                nw34_.ctor__(False, d_224_v0_, (d_277_v47_).f27, ((d_341_v97_)[default__.safeIndex(d_224_v0_, len(d_341_v97_))]) != (d_225_v1_), d_279_v49_)
                d_342_v98_ = nw34_
                d_343_v99_: T1
                nw35_ = C5()
                nw35_.ctor__(not((d_331_cf35_).f26), (d_277_v47_).f27)
                d_343_v99_ = nw35_
                d_343_v99_ = d_343_v99_
                (d_240_globalState_).f22 = (d_288_v58_).fm2(not (d_241_v14_) or (not((d_331_cf35_).f26)), d_290_v60_, d_240_globalState_)
                d_344_v100_: D0
                d_344_v100_ = D0_DC0(False)
                (d_240_globalState_).f16 = (d_331_cf35_).fm4(d_344_v100_, default__.safeModuloInt(len(d_275_v45_), (d_272_v42_)[default__.safeIndex((d_342_v98_).f31, len(d_272_v42_))]), (d_224_v0_ if (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))] else d_224_v0_), d_240_globalState_)
            d_345_v101_: C6
            nw36_ = C6()
            nw36_.ctor__((d_290_v60_) * (d_290_v60_))
            d_345_v101_ = nw36_
            d_346_v102_: D0
            d_346_v102_ = D0_DC0(False)
            d_347_v103_: D5
            d_347_v103_ = D5_DC14(d_224_v0_, d_345_v101_.f40, (d_346_v102_).cf0)
            d_348_v104_: D5
            d_348_v104_ = D5_DC15(d_347_v103_)
            pat_let_tv1_ = d_347_v103_
            def iife52_(_pat_let2_0):
                def iife53_(d_349_dt__update__tmp_h1_):
                    def iife54_(_pat_let3_0):
                        def iife55_(d_350_dt__update_hcf28_h0_):
                            return D5_DC15(d_350_dt__update_hcf28_h0_)
                        return iife55_(_pat_let3_0)
                    return iife54_(pat_let_tv1_)
                return iife53_(_pat_let2_0)
            d_348_v104_ = ((iife52_(d_348_v104_) if d_293_v63_ else d_348_v104_) if ((d_331_cf35_).f27 if (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))] else True) else d_348_v104_)
        d_351_i4_: int
        d_351_i4_ = 0
        with _dafny.label("0"):
            while (d_295_v65_).issubset((d_295_v65_) - (d_295_v65_)):
                with _dafny.c_label("0"):
                    if (d_351_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_351_i4_ = (d_351_i4_) + (1)
                    d_352_v105_: _dafny.Map
                    d_352_v105_ = _dafny.Map({563: (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]})
                    d_353_v106_: D16
                    d_353_v106_ = D16_DC37(d_352_v105_)
                    d_354_v107_: _dafny.MultiSet
                    d_354_v107_ = _dafny.MultiSet([(d_353_v106_).cf68, d_352_v105_])
                    d_355_v108_: _dafny.MultiSet
                    d_355_v108_ = d_354_v107_
                    source9_ = d_355_v108_
                    d_356___mcc_h5_ = source9_
                    d_357_cf82_ = d_356___mcc_h5_
                    (d_240_globalState_).f2 = (d_277_v47_).f27
                    d_243_v16_ = d_243_v16_
                    d_290_v60_ = default__.safeDivisionInt(d_224_v0_, (d_288_v58_).fm1(d_240_globalState_))
                    d_358_v109_: _dafny.Map
                    d_358_v109_ = _dafny.Map({d_229_v5_: (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]})
                    rhs31_ = (d_277_v47_).f27
                    rhs32_ = len(((d_358_v109_) | (_dafny.Map({d_229_v5_: (d_277_v47_).f27}))) | (d_358_v109_))
                    lhs27_ = d_240_globalState_
                    d_293_v63_ = rhs31_
                    lhs27_.f12 = rhs32_
                    if (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]:
                        (d_240_globalState_).f22 = True
                        d_359_v110_: _dafny.Array
                        def lambda13_(d_360_v1_):
                            def lambda14_(d_361_i5_):
                                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glxoh"))) + (d_360_v1_)

                            return lambda14_

                        init7_ = lambda13_(d_225_v1_)
                        nw37_ = _dafny.Array(None, 1)
                        for i0_7_ in range(nw37_.length(0)):
                            nw37_[i0_7_] = init7_(i0_7_)
                        d_359_v110_ = nw37_
                        d_362_v111_: _dafny.Map
                        d_362_v111_ = _dafny.Map({d_290_v60_: d_225_v1_})
                        index34_ = default__.safeIndex(626, (d_359_v110_).length(0))
                        (d_359_v110_)[index34_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_363_i6_ in range(default__.abs(-70))])) + (((d_362_v111_)[d_224_v0_] if (d_224_v0_) in (d_362_v111_) else d_225_v1_))
                        index35_ = default__.safeIndex(626, (d_359_v110_).length(0))
                        (d_359_v110_)[index35_] = (d_225_v1_) + (d_225_v1_)
                        d_364_v112_: D3
                        d_364_v112_ = D3_DC10(d_229_v5_)
                        d_365_v113_: _dafny.Map
                        d_365_v113_ = _dafny.Map({d_364_v112_: d_290_v60_})
                        d_366_v114_: C4
                        nw38_ = C4()
                        nw38_.ctor__(d_365_v113_, (d_277_v47_).f27, (d_277_v47_).f26)
                        d_366_v114_ = nw38_
                        d_367_v115_: _dafny.Map
                        d_367_v115_ = _dafny.Map({d_366_v114_: d_290_v60_})
                        d_367_v115_ = d_367_v115_
                        (d_240_globalState_).f6 = d_225_v1_
                        (d_240_globalState_).f6 = (d_359_v110_)[default__.safeIndex(626, (d_359_v110_).length(0))]
                    elif True:
                        (d_240_globalState_).f24 = (d_277_v47_).f27
                        d_368_v116_: C1
                        nw39_ = C1()
                        nw39_.ctor__(len(d_272_v42_), d_290_v60_, (d_224_v0_) in (d_226_v2_), (d_277_v47_).f26)
                        d_368_v116_ = nw39_
                        d_369_v117_: _dafny.Array
                        def lambda15_(d_370_v61_):
                            def lambda16_(d_371_i7_):
                                return d_370_v61_

                            return lambda16_

                        init8_ = lambda15_(d_291_v61_)
                        nw40_ = _dafny.Array(None, 8)
                        for i0_8_ in range(nw40_.length(0)):
                            nw40_[i0_8_] = init8_(i0_8_)
                        d_369_v117_ = nw40_
                        index36_ = default__.safeIndex(346, (d_369_v117_).length(0))
                        (d_369_v117_)[index36_] = d_289_v59_
                        index37_ = default__.safeIndex(346, (d_369_v117_).length(0))
                        (d_369_v117_)[index37_] = d_291_v61_
                        index38_ = default__.safeIndex(434, (d_269_v41_).length(0))
                        (d_269_v41_)[index38_] = d_241_v14_
                        d_372_v118_: _dafny.Map
                        d_372_v118_ = _dafny.Map({(d_272_v42_)[default__.safeIndex((0) - (d_368_v116_.f37), len(d_272_v42_))]: d_234_v8_})
                        d_372_v118_ = (d_372_v118_).set(d_290_v60_, (d_369_v117_)[default__.safeIndex(346, (d_369_v117_).length(0))])
                    d_373_v119_: _dafny.Map
                    d_373_v119_ = _dafny.Map({d_272_v42_: d_224_v0_})
                    d_374_v120_: _dafny.Set
                    d_374_v120_ = _dafny.Set({(d_316_v86_)[default__.safeIndex(461, (d_316_v86_).length(0))]})
                    d_375_v121_: _dafny.Seq
                    d_375_v121_ = _dafny.SeqWithoutIsStrInference([d_374_v120_, d_374_v120_, d_374_v120_, d_374_v120_])
                    d_373_v119_ = (d_373_v119_).set((d_272_v42_) + (d_272_v42_), len((d_375_v121_)[default__.safeIndex(d_224_v0_, len(d_375_v121_))]))
                    d_376_v122_: _dafny.Array
                    nw41_ = _dafny.Array(None, 8)
                    d_376_v122_ = nw41_
                    d_376_v122_ = d_376_v122_
                    pass
            pass
        hi1_ = d_290_v60_
        for d_377_i8_ in range(len(_dafny.Set({False, (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))], (d_277_v47_).f26, d_241_v14_})), hi1_):
            d_269_v41_ = d_269_v41_
            d_378_v123_: _dafny.Set
            d_378_v123_ = _dafny.Set({(d_277_v47_).f26, (d_277_v47_).f26, (d_269_v41_)[default__.safeIndex(434, (d_269_v41_).length(0))]})
            d_379_v124_: D1
            d_379_v124_ = D1_DC5((0) - (len(d_378_v123_)))
            rhs33_ = (d_377_i8_) - ((d_379_v124_).cf12)
            rhs34_ = d_224_v0_
            rhs35_ = 855
            rhs36_ = (d_226_v2_).isdisjoint(_dafny.Set({d_377_i8_}))
            lhs28_ = d_240_globalState_
            lhs29_ = d_240_globalState_
            lhs30_ = d_240_globalState_
            lhs31_ = d_240_globalState_
            lhs28_.f16 = rhs33_
            lhs29_.f12 = rhs34_
            lhs30_.f12 = rhs35_
            lhs31_.f24 = rhs36_
            index39_ = default__.safeIndex(434, (d_269_v41_).length(0))
            rhs37_ = -455
            rhs38_ = (d_277_v47_).fm2(False, len(_dafny.Map({d_290_v60_: (0) - (-932)})), d_240_globalState_)
            lhs32_ = d_240_globalState_
            lhs33_ = d_269_v41_
            lhs34_ = default__.safeIndex(434, (d_269_v41_).length(0))
            lhs32_.f12 = rhs37_
            lhs33_[lhs34_] = rhs38_
            d_380_v125_: _dafny.MultiSet
            d_380_v125_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_224_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrpw"))), d_377_i8_, (0) - (d_377_i8_), 123])])
            d_381_v126_: C3
            nw42_ = C3()
            nw42_.ctor__(d_380_v125_, False, not((d_277_v47_).f27))
            d_381_v126_ = nw42_
            d_382_v127_: D23
            d_382_v127_ = D23_DC55(d_381_v126_)
            d_381_v126_ = (d_382_v127_).cf103
        d_383_v128_: bool
        d_384_v129_: bool
        out19_: bool
        out20_: bool
        out19_, out20_ = (d_280_v50_).m4(d_240_globalState_)
        d_383_v128_ = out19_
        d_384_v129_ = out20_
        index40_ = default__.safeIndex(434, (d_269_v41_).length(0))
        (d_269_v41_)[index40_] = not(not(d_241_v14_))
        _dafny.print(_dafny.string_of(d_224_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_225_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_v2_) == (_dafny.Set({979, -6}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_227_v3_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovdhmu"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_229_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_232_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_234_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_235_v9_) == (_dafny.Set({_dafny.CodePoint('v')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v10_) == (_dafny.Map({979: _dafny.Map({_dafny.CodePoint('v'): -441})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_238_v11_) == (_dafny.Map({_dafny.CodePoint('v'): 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_239_v12_) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('v'): -441}): 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f0) == (_dafny.Set({979, -6}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f1) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ovdhmu"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_240_globalState_.f6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_240_globalState_).f7)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f14)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_240_globalState_).f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_240_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_globalState_).f25) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('v'): -441}): 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_241_v14_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v15_) == (_dafny.MultiSet([_dafny.Map({False: 979})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v16_) == (_dafny.Map({False: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v17_) == (_dafny.Map({_dafny.Map({False: 979}): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v41_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v41_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v41_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v41_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v42_) == (_dafny.SeqWithoutIsStrInference([979, 979, 979, 979]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v43_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v44_) == (_dafny.Map({False: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v45_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v46_) == (_dafny.MultiSet([1, 979, 979]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_v47_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_v47_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_278_v48_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v49_) == (_dafny.MultiSet([1, 1, 979]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_281_v51_)[0]).f29) == (_dafny.MultiSet([4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_281_v51_)[1]).f29) == (_dafny.MultiSet([4]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_282_v52_) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[0]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[1]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[2]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[3]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[4]) == (_dafny.Map({1: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[5]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[6]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[7]) == (_dafny.Map({1958: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[8]) == (_dafny.Map({979: 979, 6: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[9]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[10]) == (_dafny.Map({979: 6}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[11]) == (_dafny.Map({1: 979, 979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[12]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[13]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[14]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[15]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[16]) == (_dafny.Map({9: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[17]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[18]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[19]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[20]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[21]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_283_v55_)[22]) == (_dafny.Map({979: 979}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_286_v56_).cf29).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_287_v57_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_289_v59_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_290_v60_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_291_v61_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_292_v62_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_293_v63_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v64_) == (_dafny.Map({_dafny.MultiSet([True, True, True, True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_295_v65_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v86_)[2]).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_316_v86_)[2]).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_318_v88_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_319_v89_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_351_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_383_v128_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_384_v129_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.Array(None, 0), int(0), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {self.cf10.VerbatimString(True)}, {self.cf11.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC12(D4, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(int(0), int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({self.cf29.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(int(0), _dafny.MultiSet({}), False, _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC20(D8, NamedTuple('DC20', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC22(D9, NamedTuple('DC22', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC24(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC24(D10, NamedTuple('DC24', [('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({self.cf44.VerbatimString(True)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC25(D10, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)

class D11_DC26(D11, NamedTuple('DC26', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC28(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC28(D12, NamedTuple('DC28', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC27(D12, NamedTuple('DC27', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC30(False, int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC30(D13, NamedTuple('DC30', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC32(_dafny.MultiSet({}), False, False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC32(D14, NamedTuple('DC32', [('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC31(D14, NamedTuple('DC31', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC34(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D15_DC35)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC34(D15, NamedTuple('DC34', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC35(D15, NamedTuple('DC35', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC35({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC35) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC38(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D16_DC38)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D16_DC37)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D16_DC39)

class D16_DC38(D16, NamedTuple('DC38', [('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC38({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC38) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC39(D16, NamedTuple('DC39', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC39({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC39) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC41(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D17_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D17_DC43)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC41(D17, NamedTuple('DC41', [('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC42(D17, NamedTuple('DC42', [('cf77', Any), ('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC42({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC42) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC43(D17, NamedTuple('DC43', [('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC43({_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC43) and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)

class D18_DC44(D18, NamedTuple('DC44', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)

class D19_DC45(D19, NamedTuple('DC45', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC47()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D20_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D20_DC46)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)

class D20_DC47(D20, NamedTuple('DC47', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC47'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC47)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC46(D20, NamedTuple('DC46', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC46({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC46) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC50(False, int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D21_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D21_DC49)

class D21_DC50(D21, NamedTuple('DC50', [('cf87', Any), ('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC50({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC50) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf91', Any), ('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC52(D21, NamedTuple('DC52', [('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC49(D21, NamedTuple('DC49', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC49({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC49) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC54(_dafny.CodePoint('D'), D9.default()(), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D22_DC54)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D22_DC53)

class D22_DC54(D22, NamedTuple('DC54', [('cf98', Any), ('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC54({_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC54) and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC53(D22, NamedTuple('DC53', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC53({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC53) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC56(int(0), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D23_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D23_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D23_DC55)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D23_DC58)

class D23_DC56(D23, NamedTuple('DC56', [('cf104', Any), ('cf105', Any), ('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC56({_dafny.string_of(self.cf104)}, {_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC56) and self.cf104 == __o.cf104 and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC57(D23, NamedTuple('DC57', [])):
    def __dafnystr__(self) -> str:
        return f'D23.DC57'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC57)
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC55(D23, NamedTuple('DC55', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC55({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC55) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC58(D23, NamedTuple('DC58', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC58({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC58) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC60(None, int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D24_DC59)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D24_DC61)

class D24_DC60(D24, NamedTuple('DC60', [('cf110', Any), ('cf111', Any), ('cf112', Any), ('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111 and self.cf112 == __o.cf112 and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC59(D24, NamedTuple('DC59', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC59({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC59) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC61(D24, NamedTuple('DC61', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC61({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC61) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC63(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D25_DC64)

class D25_DC63(D25, NamedTuple('DC63', [('cf116', Any), ('cf117', Any), ('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({self.cf116.VerbatimString(True)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC62(D25, NamedTuple('DC62', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC64(D25, NamedTuple('DC64', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC64({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC64) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC66(int(0), False, None)
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

class D26_DC66(D26, NamedTuple('DC66', [('cf121', Any), ('cf122', Any), ('cf123', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC66({_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC66) and self.cf121 == __o.cf121 and self.cf122 == __o.cf122 and self.cf123 == __o.cf123
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC65(D26, NamedTuple('DC65', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC65({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC65) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC67(D26, NamedTuple('DC67', [('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC67({_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC67) and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T2:
    pass
    def m2(self, p0, p1, p2, globalState):
        pass


class T3:
    pass
    def m3(self, p0, p1, p2, p3, globalState):
        pass

    def m4(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self.f5: _dafny.Array = _dafny.Array(None, 0)
        self.f6: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f12: int = int(0)
        self.f15: bool = False
        self.f16: int = int(0)
        self.f19: str = _dafny.CodePoint('D')
        self.f22: bool = False
        self.f24: bool = False
        self._f0: _dafny.Set = _dafny.Set({})
        self._f1: _dafny.MultiSet = _dafny.MultiSet({})
        self._f3: bool = False
        self._f4: str = _dafny.CodePoint('D')
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: bool = False
        self._f9: bool = False
        self._f10: int = int(0)
        self._f11: bool = False
        self._f13: int = int(0)
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f17: bool = False
        self._f18: bool = False
        self._f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f21: int = int(0)
        self._f23: int = int(0)
        self._f25: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25

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
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25

class C0:
    def  __init__(self):
        self._f32: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f32):
        (self)._f32 = f32

    def fm6(self, p0, p1, globalState):
        return (self).f32

    def fm7(self, p0, p1, p2, globalState):
        return True

    @property
    def f32(self):
        return self._f32

class C1(T2, T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        self.f37: int = int(0)
        self._f36: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f36, f37, f26, f27):
        (self)._f36 = f36
        (self).f37 = f37
        (self)._f26 = f26
        (self)._f27 = f27

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(self).f26])) + (_dafny.SeqWithoutIsStrInference([(self).f26]))) + ((_dafny.SeqWithoutIsStrInference([(self).f27])) + (_dafny.SeqWithoutIsStrInference([(self).f27, (self).f26])))

    def fm4(self, p0, p1, p2, globalState):
        return (self).f36

    def fm1(self, globalState):
        return self.f37

    def fm2(self, p0, p1, globalState):
        return (self).f26

    def m2(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        d_385_i0_: int
        d_385_i0_ = 0
        with _dafny.label("1"):
            while True:
                with _dafny.c_label("1"):
                    if (d_385_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_385_i0_ = (d_385_i0_) + (1)
                    (globalState).f12 = self.f37
                    (globalState).f24 = not (p2) or ((self).fm2((self).f26, self.f37, globalState))
                    d_386_v0_: _dafny.Array
                    def lambda17_(d_387_i1_):
                        return _dafny.SeqWithoutIsStrInference([(self).f36, self.f37])

                    init9_ = lambda17_
                    nw43_ = _dafny.Array(None, 10)
                    for i0_9_ in range(nw43_.length(0)):
                        nw43_[i0_9_] = init9_(i0_9_)
                    d_386_v0_ = nw43_
                    d_388_v1_: _dafny.Seq
                    d_388_v1_ = _dafny.SeqWithoutIsStrInference([self.f37, (self).f36])
                    index41_ = default__.safeIndex(661, (d_386_v0_).length(0))
                    (d_386_v0_)[index41_] = d_388_v1_
                    index42_ = default__.safeIndex(661, (d_386_v0_).length(0))
                    (d_386_v0_)[index42_] = d_388_v1_
                    (globalState).f2 = p2
                    pass
            pass
        d_389_v2_: _dafny.Map
        d_389_v2_ = _dafny.Map({not((self).f27): not (p2) or (p2)})
        d_390_v3_: _dafny.Array
        nw44_ = _dafny.Array(None, 5)
        d_390_v3_ = nw44_
        d_391_v4_: C0
        nw45_ = C0()
        nw45_.ctor__((self).f27)
        d_391_v4_ = nw45_
        index43_ = default__.safeIndex(845, (d_390_v3_).length(0))
        (d_390_v3_)[index43_] = d_391_v4_
        d_392_v5_: _dafny.Set
        d_392_v5_ = _dafny.Set({self.f37, 975})
        d_393_v6_: _dafny.Array
        nw46_ = _dafny.Array(False, 9)
        d_393_v6_ = nw46_
        d_394_v7_: _dafny.MultiSet
        d_394_v7_ = _dafny.MultiSet([(self).f36])
        d_395_v8_: _dafny.Map
        d_395_v8_ = _dafny.Map({d_393_v6_: ((d_394_v7_)[(self).f36] if ((self).f36) in (d_394_v7_) else self.f37)})
        index44_ = default__.safeIndex(845, (d_390_v3_).length(0))
        rhs39_ = ((-338) - (len(d_392_v5_))) + ((self).fm4(default__.fm19((self).f36, len(d_395_v8_), globalState), -144, (self).f36, globalState))
        rhs40_ = d_389_v2_
        rhs41_ = d_391_v4_
        lhs35_ = self
        lhs36_ = d_390_v3_
        lhs37_ = default__.safeIndex(845, (d_390_v3_).length(0))
        lhs35_.f37 = rhs39_
        d_389_v2_ = rhs40_
        lhs36_[lhs37_] = rhs41_
        d_396_v9_: D2
        d_396_v9_ = D2_DC7(p2, (self).f27, -868, (self).f36)
        d_397_v10_: _dafny.Map
        d_397_v10_ = _dafny.Map({self.f37: (d_396_v9_).cf17})
        d_398_v11_: D2
        d_398_v11_ = D2_DC6(d_397_v10_)
        hi2_ = len(_dafny.Set({d_398_v11_}))
        for d_399_i2_ in range(self.f37, hi2_):
            d_392_v5_ = d_392_v5_
            d_400_v12_: _dafny.Array
            nw47_ = _dafny.Array(int(0), 6)
            d_400_v12_ = nw47_
            index45_ = default__.safeIndex(100, (d_400_v12_).length(0))
            (d_400_v12_)[index45_] = default__.safeModuloInt(796, (self).f36)
            index46_ = default__.safeIndex(100, (d_400_v12_).length(0))
            (d_400_v12_)[index46_] = default__.safeModuloInt(d_399_i2_, d_399_i2_)
            d_401_v13_: str
            d_401_v13_ = _dafny.CodePoint('a')
            d_402_v14_: _dafny.Array
            nw48_ = _dafny.Array(None, 3)
            nw48_[int(0)] = d_401_v13_
            nw48_[int(1)] = d_401_v13_
            nw48_[int(2)] = d_401_v13_
            d_402_v14_ = nw48_
            d_403_v15_: _dafny.Map
            d_403_v15_ = _dafny.Map({(d_400_v12_)[default__.safeIndex(100, (d_400_v12_).length(0))]: d_402_v14_})
            d_403_v15_ = (d_403_v15_).set(d_399_i2_, (d_402_v14_ if p0 else d_402_v14_))
            index47_ = default__.safeIndex(955, (d_393_v6_).length(0))
            (d_393_v6_)[index47_] = (self).f26
            d_404_v16_: _dafny.Seq
            d_404_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tniumeoec"))
            d_405_v17_: _dafny.Map
            d_405_v17_ = _dafny.Map({(0) - (len(d_404_v16_)): d_394_v7_})
            index48_ = default__.safeIndex(955, (d_393_v6_).length(0))
            (d_393_v6_)[index48_] = ((d_400_v12_)[default__.safeIndex(100, (d_400_v12_).length(0))]) > ((0) - (((((d_405_v17_)[(self).f36] if ((self).f36) in (d_405_v17_) else d_394_v7_)).intersection(d_394_v7_)).cardinality))
        (self).f37 = self.f37
        d_406_v18_: _dafny.Seq
        d_406_v18_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_407_v19_: _dafny.Seq
        d_407_v19_ = _dafny.SeqWithoutIsStrInference([len(d_406_v18_)])
        d_408_v20_: _dafny.Seq
        d_408_v20_ = d_407_v19_
        d_409_v21_: _dafny.Map
        d_409_v21_ = _dafny.Map({(self).f26: ((d_408_v20_)) + (_dafny.SeqWithoutIsStrInference([len(d_406_v18_)]))})
        d_410_v22_: _dafny.Array
        def lambda18_(d_411_v4_, d_412_v18_, d_413_p0_):
            def lambda19_(d_414_i3_):
                return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([D2_DC9((self).f36, self.f37, (d_411_v4_).f32)]), _dafny.MultiSet([D2_DC9((self).f36, (self).f36, True), D2_DC9((self).f36, (self).f36, (self).f27), D2_DC9(self.f37, self.f37, (d_411_v4_).f32)])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([D2_DC9((self).f36, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.MultiSet(d_412_v18_): (d_411_v4_).f32}))])), (d_411_v4_).f32), D2_DC9(len(d_412_v18_), self.f37, d_413_p0_)])]))

            return lambda19_

        init10_ = lambda18_(d_391_v4_, d_406_v18_, p0)
        nw49_ = _dafny.Array(None, 1)
        for i0_10_ in range(nw49_.length(0)):
            nw49_[i0_10_] = init10_(i0_10_)
        d_410_v22_ = nw49_
        d_415_v23_: D2
        d_415_v23_ = D2_DC9((self).f36, self.f37, (self).f27)
        d_416_v24_: _dafny.MultiSet
        d_416_v24_ = _dafny.MultiSet([d_415_v23_])
        d_417_v25_: _dafny.MultiSet
        d_417_v25_ = _dafny.MultiSet([p2])
        d_418_v26_: _dafny.Seq
        d_418_v26_ = _dafny.SeqWithoutIsStrInference([d_416_v24_, _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_415_v23_])).set(default__.safeIndex((0) - ((d_417_v25_).cardinality), len(_dafny.SeqWithoutIsStrInference([d_415_v23_]))), default__.fm20((self).f36, p2, (self).f36, (d_396_v9_).cf17, globalState)))])
        index49_ = default__.safeIndex(193, (d_410_v22_).length(0))
        (d_410_v22_)[index49_] = (d_418_v26_) + (_dafny.SeqWithoutIsStrInference([d_416_v24_ for d_419_i4_ in range(default__.abs(62))]))
        d_420_v27_: _dafny.Map
        d_420_v27_ = _dafny.Map({self.f37: d_416_v24_})
        d_421_v28_: D0
        d_421_v28_ = D0_DC0((self).f26)
        index50_ = default__.safeIndex(193, (d_410_v22_).length(0))
        rhs42_ = d_409_v21_
        rhs43_ = ((d_418_v26_).set(default__.safeIndex(-716, len(d_418_v26_)), ((d_420_v27_)[self.f37] if (self.f37) in (d_420_v27_) else d_416_v24_))) + ((((d_418_v26_).set(default__.safeIndex((self).fm4(d_421_v28_, self.f37, (self).f36, globalState), len(d_418_v26_)), d_416_v24_)).set(default__.safeIndex((self).f36, len((d_418_v26_).set(default__.safeIndex((self).fm4(d_421_v28_, self.f37, (self).f36, globalState), len(d_418_v26_)), d_416_v24_))), d_416_v24_)) + (_dafny.SeqWithoutIsStrInference([d_416_v24_, d_416_v24_])))
        lhs38_ = d_410_v22_
        lhs39_ = default__.safeIndex(193, (d_410_v22_).length(0))
        d_409_v21_ = rhs42_
        lhs38_[lhs39_] = rhs43_
        d_422_v29_: _dafny.Array
        def lambda20_(d_423_v25_):
            def lambda21_(d_424_i5_):
                return d_423_v25_

            return lambda21_

        init11_ = lambda20_(d_417_v25_)
        nw50_ = _dafny.Array(None, 29)
        for i0_11_ in range(nw50_.length(0)):
            nw50_[i0_11_] = init11_(i0_11_)
        d_422_v29_ = nw50_
        d_422_v29_ = d_422_v29_
        d_425_v30_: _dafny.Map
        d_425_v30_ = _dafny.Map({p2: _dafny.CodePoint('c')})
        d_426_v31_: str
        d_426_v31_ = _dafny.CodePoint('m')
        r0 = ((d_425_v30_)[(d_391_v4_).f32] if ((d_391_v4_).f32) in (d_425_v30_) else d_426_v31_)
        d_427_v32_: _dafny.Seq
        d_427_v32_ = _dafny.SeqWithoutIsStrInference([d_426_v31_, d_426_v31_])
        d_428_v33_: _dafny.Map
        d_428_v33_ = _dafny.Map({len(d_427_v32_): (d_397_v10_).set(self.f37, (self).f36)})
        r1 = len(d_428_v33_)
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        d_429_v0_: _dafny.Array
        nw51_ = _dafny.Array(False, 3)
        d_429_v0_ = nw51_
        index51_ = default__.safeIndex(254, (d_429_v0_).length(0))
        (d_429_v0_)[index51_] = (self).f26
        d_430_v1_: _dafny.Array
        nw52_ = _dafny.Array(int(0), 11)
        d_430_v1_ = nw52_
        d_431_v2_: _dafny.Set
        d_431_v2_ = _dafny.Set({(self).f26})
        d_432_v3_: _dafny.Seq
        d_432_v3_ = _dafny.SeqWithoutIsStrInference([d_431_v2_])
        d_433_v4_: _dafny.Map
        d_433_v4_ = _dafny.Map({len((d_432_v3_)[default__.safeIndex(850, len(d_432_v3_))]): (self).f26})
        d_434_v5_: _dafny.Map
        d_434_v5_ = _dafny.Map({(self).f26: (0) - (len(d_433_v4_))})
        d_435_v6_: _dafny.MultiSet
        d_435_v6_ = _dafny.MultiSet([((d_434_v5_)[p0] if (p0) in (d_434_v5_) else self.f37), self.f37])
        d_436_v7_: _dafny.Seq
        d_436_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))
        d_437_v8_: str
        d_437_v8_ = _dafny.CodePoint('t')
        index52_ = default__.safeIndex(592, (d_430_v1_).length(0))
        (d_430_v1_)[index52_] = ((d_435_v6_)[len((d_436_v7_).set(default__.safeIndex(561, len(d_436_v7_)), d_437_v8_))] if (len((d_436_v7_).set(default__.safeIndex(561, len(d_436_v7_)), d_437_v8_))) in (d_435_v6_) else (self).fm1(globalState))
        d_438_v9_: _dafny.Array
        def lambda22_(d_439_v8_):
            def lambda23_(d_440_i0_):
                return d_439_v8_

            return lambda23_

        init12_ = lambda22_(d_437_v8_)
        nw53_ = _dafny.Array(None, 24)
        for i0_12_ in range(nw53_.length(0)):
            nw53_[i0_12_] = init12_(i0_12_)
        d_438_v9_ = nw53_
        index53_ = default__.safeIndex(254, (d_429_v0_).length(0))
        index54_ = default__.safeIndex(592, (d_430_v1_).length(0))
        rhs44_ = p0
        rhs45_ = self.f37
        rhs46_ = d_438_v9_
        lhs40_ = d_429_v0_
        lhs41_ = default__.safeIndex(254, (d_429_v0_).length(0))
        lhs42_ = d_430_v1_
        lhs43_ = default__.safeIndex(592, (d_430_v1_).length(0))
        lhs40_[lhs41_] = rhs44_
        lhs42_[lhs43_] = rhs45_
        d_438_v9_ = rhs46_
        d_441_v10_: _dafny.MultiSet
        d_441_v10_ = _dafny.MultiSet([p0, (self).f27])
        d_442_v11_: _dafny.Seq
        d_442_v11_ = _dafny.SeqWithoutIsStrInference([self.f37, len(d_436_v7_)])
        index55_ = default__.safeIndex(592, (d_430_v1_).length(0))
        index56_ = default__.safeIndex(592, (d_430_v1_).length(0))
        index57_ = default__.safeIndex(592, (d_430_v1_).length(0))
        index58_ = default__.safeIndex(254, (d_429_v0_).length(0))
        rhs47_ = self.f37
        rhs48_ = (((d_441_v10_) | (d_441_v10_)) | (d_441_v10_)).cardinality
        rhs49_ = -718
        rhs50_ = ((d_442_v11_) != (d_442_v11_) if (p0 if p0 else not(p0)) else (((d_434_v5_)[(self).f26] if ((self).f26) in (d_434_v5_) else len(d_434_v5_))) < (349))
        rhs51_ = (self).f26
        lhs44_ = d_430_v1_
        lhs45_ = default__.safeIndex(592, (d_430_v1_).length(0))
        lhs46_ = d_430_v1_
        lhs47_ = default__.safeIndex(592, (d_430_v1_).length(0))
        lhs48_ = d_430_v1_
        lhs49_ = default__.safeIndex(592, (d_430_v1_).length(0))
        lhs50_ = globalState
        lhs51_ = d_429_v0_
        lhs52_ = default__.safeIndex(254, (d_429_v0_).length(0))
        lhs44_[lhs45_] = rhs47_
        lhs46_[lhs47_] = rhs48_
        lhs48_[lhs49_] = rhs49_
        lhs50_.f22 = rhs50_
        lhs51_[lhs52_] = rhs51_
        hi3_ = (d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))]
        for d_443_i1_ in range(self.f37, hi3_):
            d_444_v12_: C0
            nw54_ = C0()
            nw54_.ctor__((d_441_v10_).issubset(default__.fm21(self.f37, globalState)))
            d_444_v12_ = nw54_
            d_445_v13_: _dafny.Set
            d_445_v13_ = _dafny.Set({self.f37, (d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))], (self).f36, (d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))]})
            d_446_v14_: _dafny.Map
            d_446_v14_ = _dafny.Map({d_445_v13_: d_429_v0_})
            d_446_v14_ = d_446_v14_
            index59_ = default__.safeIndex(592, (d_430_v1_).length(0))
            (d_430_v1_)[index59_] = d_443_i1_
            (globalState).f15 = (p2)[default__.safeIndex(default__.safeDivisionInt(self.f37, len(d_431_v2_)), len(p2))]
        d_447_v15_: _dafny.Map
        d_447_v15_ = _dafny.Map({d_431_v2_: not((self).f26)})
        d_447_v15_ = (_dafny.Map({d_431_v2_: (d_429_v0_)[default__.safeIndex(254, (d_429_v0_).length(0))]}) if p0 else d_447_v15_)
        hi4_ = (self.f37) * ((0) - ((d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))]))
        for d_448_i2_ in range((d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))], hi4_):
            d_449_v16_: C0
            nw55_ = C0()
            nw55_.ctor__((self).f27)
            d_449_v16_ = nw55_
            (globalState).f12 = ((d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))]) + (d_448_i2_)
            d_450_v17_: D1
            d_450_v17_ = D1_DC4(d_438_v9_, (d_430_v1_)[default__.safeIndex(592, (d_430_v1_).length(0))], (_dafny.MultiSet([d_437_v8_, d_437_v8_])).cardinality, default__.fm22(len(_dafny.SeqWithoutIsStrInference([d_437_v8_ for d_451_i3_ in range(default__.abs(344))])), (self).f27, globalState), _dafny.SeqWithoutIsStrInference([d_437_v8_ for d_452_i4_ in range(default__.abs(209))]))
            d_450_v17_ = (d_450_v17_ if (self).f26 else D1_DC4(d_438_v9_, len((d_450_v17_).cf11), self.f37, d_436_v7_, d_436_v7_))
            (globalState).f15 = False
        index60_ = default__.safeIndex(592, (d_430_v1_).length(0))
        (d_430_v1_)[index60_] = self.f37

    @property
    def f36(self):
        return self._f36

class C2(T2, T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        self._f35: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f35, f26, f27):
        (self)._f35 = f35
        (self)._f26 = f26
        (self)._f27 = f27

    def fm3(self, p0, p1, p2, globalState):
        def iife56_():
            coll48_ = _dafny.Set()
            compr_48_: int
            for compr_48_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({589: True}))) for d_453_i0_ in range(default__.abs(549))])).Elements:
                d_454_v0_: int = compr_48_
                if (d_454_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({589: True}))) for d_453_i0_ in range(default__.abs(549))])):
                    coll48_ = coll48_.union(_dafny.Set([default__.safeDivisionInt(d_454_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False}))])))]))
            return _dafny.Set(coll48_)
        if (_dafny.Set({len(_dafny.Map({_dafny.Set({len(_dafny.Map({(_dafny.MultiSet([len(iife56_()
        )])).cardinality: (self).f27})), len(_dafny.Set({True}))}): len(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f27, (self).f26]))}))})) == (_dafny.Set({(D0_DC1(False, (self).f26, 832, (self).f26)).cf3})):
            return _dafny.SeqWithoutIsStrInference([(self).f26, (self).f27])
        elif True:
            return _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])

    def fm4(self, p0, p1, p2, globalState):
        def lambda24_(source10_):
            if source10_.is_DC7:
                d_455___mcc_h0_ = source10_.cf14
                d_456___mcc_h1_ = source10_.cf15
                d_457___mcc_h2_ = source10_.cf16
                d_458___mcc_h3_ = source10_.cf17
                d_459_cf17_ = d_458___mcc_h3_
                d_460_cf16_ = d_457___mcc_h2_
                d_461_cf15_ = d_456___mcc_h1_
                d_462_cf14_ = d_455___mcc_h0_
                return _dafny.Map({(self).f27: _dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('g')])})
            elif source10_.is_DC8:
                return _dafny.Map({(self).f26: _dafny.MultiSet([_dafny.CodePoint('g')])})
            elif source10_.is_DC9:
                d_463___mcc_h4_ = source10_.cf18
                d_464___mcc_h5_ = source10_.cf19
                d_465___mcc_h6_ = source10_.cf20
                d_466_cf20_ = d_465___mcc_h6_
                d_467_cf19_ = d_464___mcc_h5_
                d_468_cf18_ = d_463___mcc_h4_
                return (_dafny.Map({False: _dafny.MultiSet([_dafny.CodePoint('e'), _dafny.CodePoint('b')])})) | (_dafny.Map({False: _dafny.MultiSet([_dafny.CodePoint('s')])}))
            elif True:
                d_469___mcc_h7_ = source10_.cf13
                d_470_cf13_ = d_469___mcc_h7_
                return _dafny.Map({(self).f26: _dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('b'), _dafny.CodePoint('r')])})

        return len(lambda24_(D2_DC8()))

    def fm1(self, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(D7_DC19(len(_dafny.SeqWithoutIsStrInference([(self).f26])), _dafny.MultiSet([(self).f27]), (self).f27, _dafny.Map({(self).f27: len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f26}) for d_471_i0_ in range(default__.abs(-748))]))}))).cf36, -153])])).cardinality

    def fm2(self, p0, p1, globalState):
        return False

    def fm16(self, p0, p1, p2, p3, globalState):
        source11_ = D6_DC17(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uot"))), 884, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f26]))).cardinality, -382, len(_dafny.Set({(self).f26, (self).f27, True})))
        if source11_.is_DC17:
            d_472___mcc_h0_ = source11_.cf30
            d_473___mcc_h1_ = source11_.cf31
            d_474___mcc_h2_ = source11_.cf32
            d_475___mcc_h3_ = source11_.cf33
            d_476___mcc_h4_ = source11_.cf34
            d_477_cf34_ = d_476___mcc_h4_
            d_478_cf33_ = d_475___mcc_h3_
            d_479_cf32_ = d_474___mcc_h2_
            d_480_cf31_ = d_473___mcc_h1_
            d_481_cf30_ = d_472___mcc_h0_
            return (self).f26
        elif True:
            d_482___mcc_h5_ = source11_.cf29
            d_483_cf29_ = d_482___mcc_h5_
            return not((self).f26)

    def fm17(self, p0, p1, p2, p3, globalState):
        return (self).f26

    def m2(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r1 = (self).fm1(globalState)
        d_484_v0_: _dafny.Set
        d_484_v0_ = _dafny.Set({p0})
        d_485_v1_: int
        d_485_v1_ = -337
        d_486_v2_: _dafny.Seq
        d_486_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajxqysp"))
        rhs52_ = (0) - (((d_485_v1_) * (len(d_486_v2_))) - (d_485_v1_))
        rhs53_ = (default__.fm18(d_484_v0_, globalState)) - (d_484_v0_)
        rhs54_ = d_485_v1_
        lhs53_ = globalState
        lhs53_.f16 = rhs52_
        d_484_v0_ = rhs53_
        r1 = rhs54_
        d_487_v3_: _dafny.Seq
        d_487_v3_ = _dafny.SeqWithoutIsStrInference([p0])
        d_488_v4_: _dafny.Seq
        d_488_v4_ = _dafny.SeqWithoutIsStrInference([(d_487_v3_).set(default__.safeIndex(len(d_487_v3_), len(d_487_v3_)), p2), (d_487_v3_) + (d_487_v3_)])
        d_488_v4_ = ((d_488_v4_) + (d_488_v4_)) + ((_dafny.SeqWithoutIsStrInference([d_487_v3_ for d_489_i0_ in range(default__.abs(567))])).set(default__.safeIndex(d_485_v1_, len(_dafny.SeqWithoutIsStrInference([d_487_v3_ for d_490_i0_ in range(default__.abs(567))]))), d_487_v3_))
        d_491_v5_: _dafny.Array
        nw56_ = _dafny.Array(False, 18)
        d_491_v5_ = nw56_
        index61_ = default__.safeIndex(800, (d_491_v5_).length(0))
        (d_491_v5_)[index61_] = True
        d_492_v6_: _dafny.Map
        d_492_v6_ = _dafny.Map({_dafny.MultiSet([d_485_v1_, len(d_486_v2_), d_485_v1_, d_485_v1_]): (self).f26})
        d_493_v7_: _dafny.MultiSet
        d_493_v7_ = _dafny.MultiSet([d_485_v1_, len(d_486_v2_)])
        d_494_v8_: _dafny.Map
        d_494_v8_ = _dafny.Map({False: len(d_486_v2_)})
        d_495_v9_: _dafny.Map
        d_495_v9_ = _dafny.Map({len(d_494_v8_): not(False)})
        index62_ = default__.safeIndex(800, (d_491_v5_).length(0))
        (d_491_v5_)[index62_] = ((d_492_v6_)[(d_493_v7_) | (_dafny.MultiSet([len(d_495_v9_)]))] if ((d_493_v7_) | (_dafny.MultiSet([len(d_495_v9_)]))) in (d_492_v6_) else (self).f26)
        d_496_v10_: str
        d_496_v10_ = _dafny.CodePoint('p')
        d_497_v11_: _dafny.Map
        d_497_v11_ = _dafny.Map({(d_486_v2_)[default__.safeIndex(d_485_v1_, len(d_486_v2_))]: d_496_v10_})
        d_498_v12_: _dafny.Seq
        d_498_v12_ = _dafny.SeqWithoutIsStrInference([d_497_v11_])
        d_498_v12_ = d_498_v12_
        d_499_v13_: _dafny.Seq
        d_499_v13_ = _dafny.SeqWithoutIsStrInference([-288, (d_493_v7_).cardinality, 180, d_485_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flhymnjo")))])
        d_500_v15_: _dafny.Array
        nw57_ = _dafny.Array(None, 18)
        nw57_[int(0)] = d_499_v13_
        nw57_[int(1)] = _dafny.SeqWithoutIsStrInference([d_485_v1_ for d_501_i1_ in range(default__.abs(220))])
        nw57_[int(2)] = (d_499_v13_) + (d_499_v13_)
        nw57_[int(3)] = d_499_v13_
        nw57_[int(4)] = d_499_v13_
        nw57_[int(5)] = ((d_499_v13_ if not((d_491_v5_)[default__.safeIndex(800, (d_491_v5_).length(0))]) else d_499_v13_)).set(default__.safeIndex(d_485_v1_, len((d_499_v13_ if not((d_491_v5_)[default__.safeIndex(800, (d_491_v5_).length(0))]) else d_499_v13_))), d_485_v1_)
        nw57_[int(6)] = (d_499_v13_) + (d_499_v13_)
        nw57_[int(7)] = _dafny.SeqWithoutIsStrInference([(d_493_v7_).cardinality for d_502_i2_ in range(default__.abs(671))])
        nw57_[int(8)] = d_499_v13_
        def iife57_():
            coll49_ = _dafny.Set()
            compr_49_: str
            for compr_49_ in (d_497_v11_).keys.Elements:
                d_503_v14_: str = compr_49_
                if (d_503_v14_) in (d_497_v11_):
                    coll49_ = coll49_.union(_dafny.Set([d_503_v14_]))
            return _dafny.Set(coll49_)
        nw57_[int(9)] = (_dafny.SeqWithoutIsStrInference([len(d_486_v2_)])).set(default__.safeIndex(len(iife57_()
        ), len(_dafny.SeqWithoutIsStrInference([len(d_486_v2_)]))), d_485_v1_)
        nw57_[int(10)] = d_499_v13_
        nw57_[int(11)] = d_499_v13_
        nw57_[int(12)] = d_499_v13_
        nw57_[int(13)] = _dafny.SeqWithoutIsStrInference([d_485_v1_ for d_504_i3_ in range(default__.abs(444))])
        nw57_[int(14)] = d_499_v13_
        nw57_[int(15)] = (d_499_v13_).set(default__.safeIndex(d_485_v1_, len(d_499_v13_)), d_485_v1_)
        nw57_[int(16)] = d_499_v13_
        nw57_[int(17)] = d_499_v13_
        d_500_v15_ = nw57_
        d_505_v16_: _dafny.Set
        d_505_v16_ = _dafny.Set({(d_493_v7_).cardinality})
        d_500_v15_ = (d_500_v15_ if ((0) - (d_485_v1_)) in (d_505_v16_) else d_500_v15_)
        r0 = d_496_v10_
        r1 = ((0) - (d_485_v1_)) * ((0) - (d_485_v1_))
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        d_506_v0_: _dafny.Array
        nw58_ = _dafny.Array(int(0), 1)
        d_506_v0_ = nw58_
        d_507_v1_: _dafny.MultiSet
        d_507_v1_ = _dafny.MultiSet([d_506_v0_])
        d_508_v2_: int
        d_508_v2_ = 240
        d_509_v3_: _dafny.Map
        d_509_v3_ = _dafny.Map({d_508_v2_: d_506_v0_})
        d_510_v4_: D0
        d_510_v4_ = D0_DC0(p0)
        d_511_v5_: str
        d_511_v5_ = _dafny.CodePoint('a')
        d_512_v6_: _dafny.MultiSet
        d_512_v6_ = _dafny.MultiSet([d_511_v5_, d_511_v5_, d_511_v5_, d_511_v5_])
        d_513_v7_: _dafny.Seq
        d_513_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cudjrcpt"))
        d_514_v8_: _dafny.Map
        d_514_v8_ = _dafny.Map({len(d_513_v7_): _dafny.MultiSet([d_506_v0_, d_506_v0_, d_506_v0_, d_506_v0_, d_506_v0_])})
        d_515_v9_: _dafny.Array
        nw59_ = _dafny.Array(None, 28)
        nw59_[int(0)] = d_507_v1_
        nw59_[int(1)] = _dafny.MultiSet([((d_509_v3_)[(self).fm4(d_510_v4_, d_508_v2_, (d_512_v6_).cardinality, globalState)] if ((self).fm4(d_510_v4_, d_508_v2_, (d_512_v6_).cardinality, globalState)) in (d_509_v3_) else d_506_v0_)])
        nw59_[int(2)] = d_507_v1_
        nw59_[int(3)] = d_507_v1_
        nw59_[int(4)] = d_507_v1_
        nw59_[int(5)] = d_507_v1_
        nw59_[int(6)] = d_507_v1_
        nw59_[int(7)] = _dafny.MultiSet([d_506_v0_, d_506_v0_])
        nw59_[int(8)] = d_507_v1_
        nw59_[int(9)] = d_507_v1_
        nw59_[int(10)] = d_507_v1_
        nw59_[int(11)] = ((d_507_v1_).set(d_506_v0_, default__.abs(892))) - (d_507_v1_)
        nw59_[int(12)] = (d_507_v1_).intersection(d_507_v1_)
        nw59_[int(13)] = d_507_v1_
        nw59_[int(14)] = d_507_v1_
        nw59_[int(15)] = d_507_v1_
        nw59_[int(16)] = d_507_v1_
        nw59_[int(17)] = (d_507_v1_) | (d_507_v1_)
        nw59_[int(18)] = d_507_v1_
        nw59_[int(19)] = d_507_v1_
        nw59_[int(20)] = (d_507_v1_) | (_dafny.MultiSet([d_506_v0_, d_506_v0_, d_506_v0_, d_506_v0_, d_506_v0_]))
        nw59_[int(21)] = d_507_v1_
        nw59_[int(22)] = (d_507_v1_) - (d_507_v1_)
        nw59_[int(23)] = d_507_v1_
        nw59_[int(24)] = d_507_v1_
        nw59_[int(25)] = d_507_v1_
        nw59_[int(26)] = ((d_514_v8_)[(0) - ((0) - (d_508_v2_))] if ((0) - ((0) - (d_508_v2_))) in (d_514_v8_) else _dafny.MultiSet([d_506_v0_, d_506_v0_, d_506_v0_]))
        nw59_[int(27)] = d_507_v1_
        d_515_v9_ = nw59_
        index63_ = default__.safeIndex(532, (d_515_v9_).length(0))
        (d_515_v9_)[index63_] = d_507_v1_
        index64_ = default__.safeIndex(532, (d_515_v9_).length(0))
        (d_515_v9_)[index64_] = d_507_v1_
        d_516_i0_: int
        d_516_i0_ = 0
        with _dafny.label("2"):
            while not((d_512_v6_).ispropersubset((_dafny.MultiSet([d_511_v5_])) | (_dafny.MultiSet([d_511_v5_])))):
                with _dafny.c_label("2"):
                    if (d_516_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_516_i0_ = (d_516_i0_) + (1)
                    d_517_v10_: _dafny.Map
                    d_517_v10_ = _dafny.Map({(self).f26: d_508_v2_})
                    d_517_v10_ = (d_517_v10_).set((self).f27, d_508_v2_)
                    d_518_v11_: C0
                    nw60_ = C0()
                    nw60_.ctor__(p0)
                    d_518_v11_ = nw60_
                    index65_ = default__.safeIndex(433, (d_506_v0_).length(0))
                    (d_506_v0_)[index65_] = (0) - (default__.safeDivisionInt((0) - ((self).fm1(globalState)), d_508_v2_))
                    index66_ = default__.safeIndex(433, (d_506_v0_).length(0))
                    (d_506_v0_)[index66_] = d_508_v2_
                    d_519_v12_: _dafny.Array
                    def lambda25_(d_520_v0_):
                        def lambda26_(d_521_i1_):
                            return default__.safeModuloInt(d_521_i1_, (d_520_v0_)[default__.safeIndex(433, (d_520_v0_).length(0))])

                        return lambda26_

                    init13_ = lambda25_(d_506_v0_)
                    nw61_ = _dafny.Array(None, 23)
                    for i0_13_ in range(nw61_.length(0)):
                        nw61_[i0_13_] = init13_(i0_13_)
                    d_519_v12_ = nw61_
                    d_522_v13_: D3
                    d_522_v13_ = D3_DC10(d_519_v12_)
                    d_522_v13_ = D3_DC10(d_519_v12_)
                    pass
            pass
        d_523_v14_: _dafny.Set
        d_523_v14_ = _dafny.Set({d_508_v2_})
        d_524_v15_: _dafny.MultiSet
        d_524_v15_ = _dafny.MultiSet([d_523_v14_])
        d_525_v16_: _dafny.MultiSet
        d_525_v16_ = _dafny.MultiSet([(self).fm17(len(d_513_v7_), (0) - (d_508_v2_), d_508_v2_, d_524_v15_, globalState), (self).f27, (self).f26, not(not(False))])
        (globalState).f12 = ((len(d_513_v7_)) * (((d_525_v16_)[not(p0)] if (not(p0)) in (d_525_v16_) else d_508_v2_))) - ((self).fm1(globalState))
        d_526_v17_: _dafny.Map
        d_526_v17_ = _dafny.Map({p0: (self).f26})
        (self).m12(d_508_v2_, len(d_526_v17_), globalState)
        d_527_v18_: T2
        nw62_ = C1()
        nw62_.ctor__(d_508_v2_, len(d_513_v7_), p0, (self).f26)
        d_527_v18_ = nw62_
        d_528_v19_: D7
        d_528_v19_ = D7_DC18(d_527_v18_)
        source12_ = d_528_v19_
        if source12_.is_DC19:
            d_529___mcc_h0_ = source12_.cf36
            d_530___mcc_h1_ = source12_.cf37
            d_531___mcc_h2_ = source12_.cf38
            d_532___mcc_h3_ = source12_.cf39
            d_533_cf39_ = d_532___mcc_h3_
            d_534_cf38_ = d_531___mcc_h2_
            d_535_cf37_ = d_530___mcc_h1_
            d_536_cf36_ = d_529___mcc_h0_
            (globalState).f15 = (d_527_v18_).f26
            (globalState).f12 = (d_508_v2_) + ((d_508_v2_) - (648))
            d_537_v20_: _dafny.Map
            d_537_v20_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f26])): d_536_cf36_})
            d_538_v21_: _dafny.Array
            nw63_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_538_v21_ = nw63_
            d_539_v22_: _dafny.Map
            d_539_v22_ = _dafny.Map({d_538_v21_: (d_527_v18_).f27})
            d_540_v23_: _dafny.Map
            d_540_v23_ = _dafny.Map({((d_539_v22_)[d_538_v21_] if (d_538_v21_) in (d_539_v22_) else (d_527_v18_).f27): d_526_v17_})
            d_536_cf36_ = (d_527_v18_).fm4(d_510_v4_, (len(d_537_v20_)) * (d_508_v2_), len(d_540_v23_), globalState)
            d_511_v5_ = d_511_v5_
        elif True:
            d_541___mcc_h4_ = source12_.cf35
            d_542_cf35_ = d_541___mcc_h4_
            d_543_v24_: _dafny.Array
            def lambda27_(d_544_v2_, d_545_v7_, d_546_v18_):
                def lambda28_(d_547_i2_):
                    return (D5_DC14(d_544_v2_, len(d_545_v7_), (d_546_v18_).f27)).cf27

                return lambda28_

            init14_ = lambda27_(d_508_v2_, d_513_v7_, d_527_v18_)
            nw64_ = _dafny.Array(None, 18)
            for i0_14_ in range(nw64_.length(0)):
                nw64_[i0_14_] = init14_(i0_14_)
            d_543_v24_ = nw64_
            rhs55_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htwwcultw"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlkne")))
            rhs56_ = d_543_v24_
            lhs54_ = globalState
            lhs54_.f6 = rhs55_
            d_543_v24_ = rhs56_
            (globalState).f12 = (d_508_v2_) * (970)
            d_548_v25_: _dafny.Array
            def lambda29_(d_549_v14_):
                def lambda30_(d_550_i3_):
                    return d_549_v14_

                return lambda30_

            init15_ = lambda29_(d_523_v14_)
            nw65_ = _dafny.Array(None, 19)
            for i0_15_ in range(nw65_.length(0)):
                nw65_[i0_15_] = init15_(i0_15_)
            d_548_v25_ = nw65_
            index67_ = default__.safeIndex(648, (d_548_v25_).length(0))
            def iife58_():
                coll50_ = _dafny.Set()
                compr_50_: int
                for compr_50_ in (_dafny.MultiSet(default__.fm23(d_511_v5_, globalState))).Elements:
                    d_551_v26_: int = compr_50_
                    if (d_551_v26_) in (_dafny.MultiSet(default__.fm23(d_511_v5_, globalState))):
                        coll50_ = coll50_.union(_dafny.Set([default__.safeDivisionInt(d_551_v26_, len(_dafny.Set({False})))]))
                return _dafny.Set(coll50_)
            (d_548_v25_)[index67_] = iife58_()
            
            index68_ = default__.safeIndex(648, (d_548_v25_).length(0))
            (d_548_v25_)[index68_] = d_523_v14_
            (globalState).f16 = (0) - (d_508_v2_)
        d_511_v5_ = (d_511_v5_ if (d_527_v18_).fm2(p0, d_508_v2_, globalState) else (d_511_v5_ if p0 else d_511_v5_))

    def m11(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_552_v0_: D0
        d_552_v0_ = D0_DC0((self).f26)
        d_553_v1_: _dafny.Seq
        d_553_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "euqbmr"))
        d_554_v2_: _dafny.Map
        d_554_v2_ = _dafny.Map({d_553_v1_: (self).f27})
        d_555_v3_: int
        d_555_v3_ = 671
        d_556_v4_: _dafny.Array
        nw66_ = _dafny.Array(int(0), 14)
        d_556_v4_ = nw66_
        d_557_v5_: _dafny.Map
        def iife59_(_pat_let4_0):
            def iife60_(d_558_dt__update__tmp_h0_):
                def iife61_(_pat_let5_0):
                    def iife62_(d_559_dt__update_hcf0_h0_):
                        return D0_DC0(d_559_dt__update_hcf0_h0_)
                    return iife62_(_pat_let5_0)
                return iife61_((self).f26)
            return iife60_(_pat_let4_0)
        d_557_v5_ = _dafny.Map({(self).fm4(iife59_(d_552_v0_), len(d_554_v2_), d_555_v3_, globalState): d_556_v4_})
        d_557_v5_ = (d_557_v5_).set(d_555_v3_, d_556_v4_)
        d_560_v6_: _dafny.Map
        d_560_v6_ = _dafny.Map({d_553_v1_: d_555_v3_})
        hi5_ = d_555_v3_
        for d_561_i0_ in range(len(d_560_v6_), hi5_):
            d_562_v7_: _dafny.Map
            d_562_v7_ = _dafny.Map({default__.fm24(d_553_v1_, (self).f26, d_555_v3_, globalState): (self).f27})
            d_562_v7_ = (d_562_v7_) | (d_562_v7_)
            d_563_v8_: str
            d_563_v8_ = _dafny.CodePoint('n')
            d_564_v9_: str
            d_565_v10_: _dafny.Seq
            d_566_v11_: bool
            d_567_v12_: _dafny.Map
            out21_: str
            out22_: _dafny.Seq
            out23_: bool
            out24_: _dafny.Map
            out21_, out22_, out23_, out24_ = default__.m0(d_563_v8_, globalState)
            d_564_v9_ = out21_
            d_565_v10_ = out22_
            d_566_v11_ = out23_
            d_567_v12_ = out24_
            index69_ = default__.safeIndex(234, (d_556_v4_).length(0))
            (d_556_v4_)[index69_] = d_555_v3_
            d_568_v13_: D6
            d_568_v13_ = D6_DC16((d_553_v1_) + (d_553_v1_))
            d_569_v14_: D2
            d_569_v14_ = D2_DC9(d_561_i0_, d_561_i0_, (self).f27)
            d_570_v15_: _dafny.MultiSet
            d_570_v15_ = _dafny.MultiSet([_dafny.Set({d_555_v3_, 201})])
            index70_ = default__.safeIndex(234, (d_556_v4_).length(0))
            rhs57_ = d_561_i0_
            rhs58_ = d_561_i0_
            rhs59_ = (d_568_v13_ if ((d_569_v14_).cf19) != (d_555_v3_) else D6_DC16(d_553_v1_))
            rhs60_ = (self).fm17(d_561_i0_, (0) - (d_561_i0_), d_555_v3_, d_570_v15_, globalState)
            lhs55_ = globalState
            lhs56_ = d_556_v4_
            lhs57_ = default__.safeIndex(234, (d_556_v4_).length(0))
            lhs58_ = globalState
            lhs55_.f12 = rhs57_
            lhs56_[lhs57_] = rhs58_
            d_568_v13_ = rhs59_
            lhs58_.f2 = rhs60_
            d_571_v16_: _dafny.Array
            def lambda31_(d_572_i1_):
                return (self).f27

            init16_ = lambda31_
            nw67_ = _dafny.Array(None, 13)
            for i0_16_ in range(nw67_.length(0)):
                nw67_[i0_16_] = init16_(i0_16_)
            d_571_v16_ = nw67_
            d_573_v17_: _dafny.Map
            d_573_v17_ = _dafny.Map({not((self).f27): (self).f27})
            d_571_v16_ = (d_571_v16_ if ((d_573_v17_)[d_566_v11_] if (d_566_v11_) in (d_573_v17_) else True) else p1)
        d_574_i2_: int
        d_574_i2_ = 0
        with _dafny.label("3"):
            while (False) == ((self).f26):
                with _dafny.c_label("3"):
                    if (d_574_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_574_i2_ = (d_574_i2_) + (1)
                    d_575_v18_: str
                    d_575_v18_ = _dafny.CodePoint('g')
                    d_576_v19_: _dafny.Set
                    d_576_v19_ = _dafny.Set({(self).f27})
                    d_577_v20_: _dafny.Map
                    d_577_v20_ = _dafny.Map({len((_dafny.SeqWithoutIsStrInference([d_575_v18_, d_575_v18_])) + (d_553_v1_)): (d_576_v19_).intersection(d_576_v19_)})
                    d_577_v20_ = d_577_v20_
                    (globalState).f22 = (self).f26
                    (globalState).f12 = d_555_v3_
                    (globalState).f16 = 308
                    pass
            pass
        d_578_v21_: _dafny.MultiSet
        d_578_v21_ = _dafny.MultiSet([d_555_v3_])
        d_579_v22_: _dafny.Set
        d_579_v22_ = _dafny.Set({(self).f27, False})
        index71_ = default__.safeIndex(156, (p1).length(0))
        (p1)[index71_] = (d_578_v21_).ispropersubset(_dafny.MultiSet([len(d_579_v22_), d_555_v3_]))
        d_580_v23_: _dafny.Map
        d_580_v23_ = _dafny.Map({(self).f27: False})
        d_581_v24_: _dafny.MultiSet
        d_581_v24_ = _dafny.MultiSet([((d_580_v23_)[(self).f27] if ((self).f27) in (d_580_v23_) else False)])
        d_582_v25_: _dafny.Array
        def lambda32_(d_583_i3_):
            return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([(self).f27]))})

        init17_ = lambda32_
        nw68_ = _dafny.Array(None, 6)
        for i0_17_ in range(nw68_.length(0)):
            nw68_[i0_17_] = init17_(i0_17_)
        d_582_v25_ = nw68_
        d_584_v26_: _dafny.Map
        d_584_v26_ = _dafny.Map({d_555_v3_: d_581_v24_})
        d_585_v27_: _dafny.Map
        d_585_v27_ = _dafny.Map({d_579_v22_: d_582_v25_})
        index72_ = default__.safeIndex(156, (p1).length(0))
        rhs61_ = (self).f26
        rhs62_ = ((d_584_v26_)[d_555_v3_] if (d_555_v3_) in (d_584_v26_) else _dafny.MultiSet([True, (self).f27]))
        rhs63_ = ((self).f27) and ((self).f26)
        rhs64_ = (d_582_v25_ if True else ((d_585_v27_)[_dafny.Set({(self).f27})] if (_dafny.Set({(self).f27})) in (d_585_v27_) else d_582_v25_))
        rhs65_ = (0) - (d_555_v3_)
        lhs59_ = p1
        lhs60_ = default__.safeIndex(156, (p1).length(0))
        lhs61_ = globalState
        lhs59_[lhs60_] = rhs61_
        d_581_v24_ = rhs62_
        lhs61_.f15 = rhs63_
        d_582_v25_ = rhs64_
        d_555_v3_ = rhs65_
        d_586_v28_: _dafny.Set
        d_586_v28_ = _dafny.Set({d_555_v3_})
        d_587_i4_: int
        d_587_i4_ = 0
        with _dafny.label("4"):
            while ((d_555_v3_) + (d_555_v3_)) not in (d_586_v28_):
                with _dafny.c_label("4"):
                    if (d_587_i4_) >= (100):
                        raise _dafny.Break("4")
                    d_587_i4_ = (d_587_i4_) + (1)
                    index73_ = default__.safeIndex(960, (d_556_v4_).length(0))
                    (d_556_v4_)[index73_] = d_555_v3_
                    d_588_v29_: _dafny.Set
                    d_588_v29_ = d_579_v22_
                    d_589_v30_: str
                    d_589_v30_ = _dafny.CodePoint('c')
                    d_590_v31_: _dafny.Seq
                    d_590_v31_ = _dafny.SeqWithoutIsStrInference([(self).f26, False, True])
                    d_591_v32_: _dafny.Map
                    d_591_v32_ = _dafny.Map({(self).f26: d_590_v31_})
                    index74_ = default__.safeIndex(960, (d_556_v4_).length(0))
                    rhs66_ = ((self).fm4(d_552_v0_, len(_dafny.Map({(p1)[default__.safeIndex(156, (p1).length(0))]: d_589_v30_})), len(d_591_v32_), globalState)) - (d_555_v3_)
                    rhs67_ = d_555_v3_
                    rhs68_ = d_588_v29_
                    lhs62_ = globalState
                    lhs63_ = d_556_v4_
                    lhs64_ = default__.safeIndex(960, (d_556_v4_).length(0))
                    lhs62_.f16 = rhs66_
                    lhs63_[lhs64_] = rhs67_
                    d_588_v29_ = rhs68_
                    d_592_v33_: C1
                    nw69_ = C1()
                    nw69_.ctor__(d_555_v3_, d_555_v3_, (self).f27, (p1)[default__.safeIndex(156, (p1).length(0))])
                    d_592_v33_ = nw69_
                    d_593_v34_: _dafny.Seq
                    d_593_v34_ = _dafny.SeqWithoutIsStrInference([d_592_v33_, d_592_v33_, d_592_v33_])
                    rhs69_ = (default__.fm22(d_555_v3_, (self).f26, globalState)) <= ((d_553_v1_) + (d_553_v1_))
                    rhs70_ = (d_593_v34_)[default__.safeIndex(d_555_v3_, len(d_593_v34_))]
                    rhs71_ = (d_589_v30_) in (d_553_v1_)
                    rhs72_ = d_589_v30_
                    lhs65_ = globalState
                    lhs66_ = globalState
                    lhs65_.f22 = rhs69_
                    d_592_v33_ = rhs70_
                    lhs66_.f22 = rhs71_
                    d_589_v30_ = rhs72_
                    d_594_v35_: C0
                    nw70_ = C0()
                    nw70_.ctor__((p1)[default__.safeIndex(156, (p1).length(0))])
                    d_594_v35_ = nw70_
                    d_595_v36_: C1
                    nw71_ = C1()
                    nw71_.ctor__((d_592_v33_).f36, d_555_v3_, (d_594_v35_).f32, (p1)[default__.safeIndex(156, (p1).length(0))])
                    d_595_v36_ = nw71_
                    pass
            pass
        hi6_ = (self).fm1(globalState)
        for d_596_i5_ in range(d_555_v3_, hi6_):
            (globalState).f19 = (d_553_v1_)[default__.safeIndex((d_555_v3_ if (self).f27 else d_596_i5_), len(d_553_v1_))]
            d_597_v37_: str
            d_597_v37_ = _dafny.CodePoint('n')
            if (-287) <= (len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbthv"))).set(default__.safeIndex(d_596_i5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbthv")))), d_597_v37_)) + (d_553_v1_))):
                d_598_v38_: _dafny.Map
                d_598_v38_ = _dafny.Map({d_596_i5_: (self).f27})
                d_599_v39_: D2
                d_599_v39_ = D2_DC9(len(d_598_v38_), d_555_v3_, False)
                d_600_v40_: _dafny.Seq
                d_600_v40_ = _dafny.SeqWithoutIsStrInference([d_599_v39_, d_599_v39_])
                d_601_v41_: C1
                nw72_ = C1()
                nw72_.ctor__(133, d_555_v3_, (self).fm16(d_596_i5_, d_600_v40_, len(_dafny.Set({d_596_i5_})), len(d_553_v1_), globalState), (self).f26)
                d_601_v41_ = nw72_
                d_602_v42_: _dafny.Map
                d_602_v42_ = _dafny.Map({(d_601_v41_).f36: _dafny.SeqWithoutIsStrInference([True, (p1)[default__.safeIndex(156, (p1).length(0))]])})
                d_603_v43_: _dafny.Seq
                d_603_v43_ = _dafny.SeqWithoutIsStrInference([True])
                d_604_v44_: _dafny.Map
                d_604_v44_ = _dafny.Map({((d_602_v42_)[(d_601_v41_).f36] if ((d_601_v41_).f36) in (d_602_v42_) else d_603_v43_): p1})
                rhs73_ = ((D9_DC21(d_604_v44_)).cf41) | (d_604_v44_)
                rhs74_ = ((d_596_i5_ if (self).f26 else d_555_v3_)) * (default__.safeDivisionInt((d_601_v41_).f36, d_596_i5_))
                lhs67_ = globalState
                d_604_v44_ = rhs73_
                lhs67_.f12 = rhs74_
                d_605_v45_: C1
                nw73_ = C1()
                nw73_.ctor__(d_601_v41_.f37, (d_601_v41_.f37) - (d_596_i5_), not((self).f27), (False) and ((self).f26))
                d_605_v45_ = nw73_
                d_606_v46_: _dafny.Map
                d_606_v46_ = _dafny.Map({(self).f26: (d_601_v41_).f36})
                d_580_v23_ = (d_580_v23_).set((len(d_606_v46_)) != (d_601_v41_.f37), (p1)[default__.safeIndex(156, (p1).length(0))])
                d_607_v47_: _dafny.Array
                nw74_ = _dafny.Array(D0.default()(), 21)
                d_607_v47_ = nw74_
                d_608_v48_: _dafny.Map
                d_608_v48_ = _dafny.Map({d_597_v37_: d_607_v47_})
                d_608_v48_ = (d_608_v48_).set(d_597_v37_, d_607_v47_)
            elif True:
                d_609_v49_: _dafny.MultiSet
                d_609_v49_ = _dafny.MultiSet([d_586_v28_])
                d_610_v50_: _dafny.Seq
                d_610_v50_ = _dafny.SeqWithoutIsStrInference([(self).fm17(d_555_v3_, d_596_i5_, d_555_v3_, d_609_v49_, globalState)])
                (globalState).f16 = (self).fm4(d_552_v0_, len(((self).fm3(d_553_v1_, d_610_v50_, len(_dafny.Map({d_555_v3_: (p1)[default__.safeIndex(156, (p1).length(0))]})), globalState)) + (d_610_v50_)), 864, globalState)
                (globalState).f16 = (0) - ((self).fm4(d_552_v0_, d_555_v3_, 400, globalState))
                (globalState).f12 = default__.safeModuloInt((0) - ((d_555_v3_ if (p1)[default__.safeIndex(156, (p1).length(0))] else d_555_v3_)), len((d_586_v28_).intersection(d_586_v28_)))
                (globalState).f12 = d_596_i5_
                (globalState).f24 = (p1)[default__.safeIndex(156, (p1).length(0))]
            d_611_v51_: _dafny.Seq
            d_611_v51_ = _dafny.SeqWithoutIsStrInference([not(False), (p1)[default__.safeIndex(156, (p1).length(0))], True])
            d_612_v52_: _dafny.Map
            d_612_v52_ = _dafny.Map({(d_611_v51_) + (_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(156, (p1).length(0))], (p1)[default__.safeIndex(156, (p1).length(0))]])): (self).f26})
            d_612_v52_ = d_612_v52_
            d_611_v51_ = (_dafny.SeqWithoutIsStrInference([(self).f27])) + ((d_611_v51_) + (_dafny.SeqWithoutIsStrInference([not((self).f26)])))
        r0 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmiwf"))
        return r0

    def m12(self, p0, p1, globalState):
        d_613_v0_: _dafny.MultiSet
        d_613_v0_ = _dafny.MultiSet([p1])
        d_614_v1_: C1
        nw75_ = C1()
        nw75_.ctor__(p1, p1, ((self).f27 if (self).f27 else (self).f27), (_dafny.MultiSet([p1, p1])).isdisjoint(d_613_v0_))
        d_614_v1_ = nw75_
        d_615_v2_: _dafny.Array
        def lambda33_(d_616_i0_):
            return (self).f26

        init18_ = lambda33_
        nw76_ = _dafny.Array(None, 17)
        for i0_18_ in range(nw76_.length(0)):
            nw76_[i0_18_] = init18_(i0_18_)
        d_615_v2_ = nw76_
        d_617_v3_: _dafny.Seq
        d_617_v3_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f27])
        d_618_v4_: _dafny.Map
        d_618_v4_ = _dafny.Map({(self).f26: d_617_v3_})
        d_619_v5_: D5
        d_619_v5_ = D5_DC13(d_618_v4_)
        rhs75_ = not ((_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27, (self).f27])) != (_dafny.SeqWithoutIsStrInference([(self).f27, (self).f26]))) or ((self).f26)
        rhs76_ = d_615_v2_
        rhs77_ = (self).f27
        rhs78_ = d_619_v5_
        lhs68_ = globalState
        lhs69_ = globalState
        lhs68_.f24 = rhs75_
        d_615_v2_ = rhs76_
        lhs69_.f22 = rhs77_
        d_619_v5_ = rhs78_
        d_620_v6_: _dafny.Seq
        d_620_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adlk"))
        hi7_ = default__.safeDivisionInt(d_614_v1_.f37, (0) - (len(d_617_v3_)))
        for d_621_i1_ in range(len(d_620_v6_), hi7_):
            (globalState).f12 = d_621_i1_
            d_622_v7_: _dafny.Array
            def lambda34_(d_623_v1_):
                def lambda35_(d_624_i2_):
                    return default__.safeDivisionInt(d_624_i2_, (d_623_v1_).f36)

                return lambda35_

            init19_ = lambda34_(d_614_v1_)
            nw77_ = _dafny.Array(None, 12)
            for i0_19_ in range(nw77_.length(0)):
                nw77_[i0_19_] = init19_(i0_19_)
            d_622_v7_ = nw77_
            index75_ = default__.safeIndex(313, (d_622_v7_).length(0))
            (d_622_v7_)[index75_] = d_614_v1_.f37
            d_625_v8_: str
            d_625_v8_ = _dafny.CodePoint('s')
            d_626_v9_: _dafny.Seq
            d_626_v9_ = _dafny.SeqWithoutIsStrInference([d_615_v2_, d_615_v2_])
            index76_ = default__.safeIndex(313, (d_622_v7_).length(0))
            rhs79_ = (d_620_v6_).set(default__.safeIndex(len(d_617_v3_), len(d_620_v6_)), d_625_v8_)
            rhs80_ = (0) - ((0) - (len(d_626_v9_)))
            rhs81_ = (self).f26
            rhs82_ = (self).f27
            rhs83_ = not((d_614_v1_).fm2((self).f26, (d_614_v1_.f37 if True else d_614_v1_.f37), globalState))
            lhs70_ = globalState
            lhs71_ = d_622_v7_
            lhs72_ = default__.safeIndex(313, (d_622_v7_).length(0))
            lhs73_ = globalState
            lhs74_ = globalState
            lhs75_ = globalState
            lhs70_.f6 = rhs79_
            lhs71_[lhs72_] = rhs80_
            lhs73_.f22 = rhs81_
            lhs74_.f24 = rhs82_
            lhs75_.f2 = rhs83_
            d_627_v10_: _dafny.Set
            d_627_v10_ = _dafny.Set({not((self).f27), (self).f27, not(False)})
            d_627_v10_ = ((d_627_v10_).intersection(d_627_v10_)).intersection((d_627_v10_) | (d_627_v10_))
            d_628_v11_: _dafny.Array
            nw78_ = _dafny.Array(None, 27)
            nw78_[int(0)] = d_615_v2_
            nw78_[int(1)] = d_615_v2_
            nw78_[int(2)] = d_615_v2_
            nw78_[int(3)] = d_615_v2_
            nw78_[int(4)] = d_615_v2_
            nw78_[int(5)] = d_615_v2_
            nw78_[int(6)] = d_615_v2_
            nw78_[int(7)] = d_615_v2_
            nw78_[int(8)] = d_615_v2_
            nw78_[int(9)] = d_615_v2_
            nw78_[int(10)] = d_615_v2_
            nw78_[int(11)] = d_615_v2_
            nw78_[int(12)] = d_615_v2_
            nw78_[int(13)] = d_615_v2_
            nw78_[int(14)] = d_615_v2_
            nw78_[int(15)] = d_615_v2_
            nw78_[int(16)] = d_615_v2_
            nw78_[int(17)] = d_615_v2_
            nw78_[int(18)] = d_615_v2_
            nw78_[int(19)] = d_615_v2_
            nw78_[int(20)] = d_615_v2_
            nw78_[int(21)] = d_615_v2_
            nw78_[int(22)] = d_615_v2_
            nw78_[int(23)] = d_615_v2_
            nw78_[int(24)] = d_615_v2_
            nw78_[int(25)] = d_615_v2_
            nw78_[int(26)] = d_615_v2_
            d_628_v11_ = nw78_
            d_629_v12_: _dafny.Seq
            d_629_v12_ = _dafny.SeqWithoutIsStrInference([d_628_v11_, d_628_v11_, d_628_v11_])
            rhs84_ = ((self).f27) == ((self).f26)
            rhs85_ = (d_629_v12_)[default__.safeIndex((d_614_v1_).f36, len(d_629_v12_))]
            lhs76_ = globalState
            lhs77_ = globalState
            lhs76_.f22 = rhs84_
            lhs77_.f5 = rhs85_
        index77_ = default__.safeIndex(575, (d_615_v2_).length(0))
        (d_615_v2_)[index77_] = (self).f27
        index78_ = default__.safeIndex(575, (d_615_v2_).length(0))
        (d_615_v2_)[index78_] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_630_i3_ in range(default__.abs(815))])) + (d_620_v6_)) < (d_620_v6_)
        (globalState).f16 = (0) - (p0)
        (d_614_v1_).f37 = len((d_620_v6_) + (d_620_v6_))

    @property
    def f35(self):
        return self._f35

class C3(T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        self._f34: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f34, f26, f27):
        (self)._f34 = f34
        (self)._f26 = f26
        (self)._f27 = f27

    def fm1(self, globalState):
        if (self).f26:
            return (len((D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoqe")))).cf29)) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmfpgf"))))
        elif True:
            return len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f26, (self).f26]))

    def fm2(self, p0, p1, globalState):
        return False

    def fm12(self, p0, globalState):
        return ((84 if (self).f27 else 719)) + (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(self).f26])), 676))

    def fm13(self, p0, p1, p2, globalState):
        return (len(_dafny.SeqWithoutIsStrInference([len((D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))).cf29) for d_631_i0_ in range(default__.abs(109))]))) - ((len(_dafny.Map({(self).f27: _dafny.Map({956: (self).f26})}))) * ((0) - (len(_dafny.Map({False: _dafny.CodePoint('k')})))))

    def m1(self, p0, p1, p2, globalState):
        d_632_v0_: _dafny.Array
        nw79_ = _dafny.Array(int(0), 21)
        d_632_v0_ = nw79_
        d_633_v1_: int
        d_633_v1_ = 281
        d_634_v2_: _dafny.Map
        d_634_v2_ = _dafny.Map({d_633_v1_: d_632_v0_})
        d_635_v3_: _dafny.Array
        nw80_ = _dafny.Array(None, 20)
        nw80_[int(0)] = d_632_v0_
        nw80_[int(1)] = d_632_v0_
        nw80_[int(2)] = d_632_v0_
        nw80_[int(3)] = d_632_v0_
        nw80_[int(4)] = d_632_v0_
        nw80_[int(5)] = d_632_v0_
        nw80_[int(6)] = d_632_v0_
        nw80_[int(7)] = d_632_v0_
        nw80_[int(8)] = d_632_v0_
        nw80_[int(9)] = d_632_v0_
        nw80_[int(10)] = d_632_v0_
        nw80_[int(11)] = d_632_v0_
        nw80_[int(12)] = d_632_v0_
        nw80_[int(13)] = ((d_634_v2_)[d_633_v1_] if (d_633_v1_) in (d_634_v2_) else d_632_v0_)
        nw80_[int(14)] = d_632_v0_
        nw80_[int(15)] = d_632_v0_
        nw80_[int(16)] = d_632_v0_
        nw80_[int(17)] = d_632_v0_
        nw80_[int(18)] = d_632_v0_
        nw80_[int(19)] = d_632_v0_
        d_635_v3_ = nw80_
        index79_ = default__.safeIndex(142, (d_635_v3_).length(0))
        (d_635_v3_)[index79_] = d_632_v0_
        d_636_v4_: _dafny.Seq
        d_636_v4_ = _dafny.SeqWithoutIsStrInference([d_632_v0_, d_632_v0_])
        index80_ = default__.safeIndex(142, (d_635_v3_).length(0))
        (d_635_v3_)[index80_] = (d_636_v4_)[default__.safeIndex(918, len(d_636_v4_))]
        d_637_v5_: _dafny.Map
        d_637_v5_ = _dafny.Map({(self).f27: (self).f26})
        d_638_v6_: _dafny.Seq
        d_638_v6_ = _dafny.SeqWithoutIsStrInference([d_637_v5_, _dafny.Map({p0: (self).f27}), d_637_v5_, d_637_v5_, _dafny.Map({(self).f27: (self).f26})])
        d_639_v7_: _dafny.Seq
        d_639_v7_ = _dafny.SeqWithoutIsStrInference([d_633_v1_, d_633_v1_, d_633_v1_])
        d_640_v8_: _dafny.Map
        d_640_v8_ = _dafny.Map({d_638_v6_: d_639_v7_})
        d_641_v9_: _dafny.Seq
        d_641_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgla"))
        d_642_v10_: _dafny.Map
        d_642_v10_ = _dafny.Map({d_641_v9_: d_638_v6_})
        d_640_v8_ = (d_640_v8_).set((((d_642_v10_)[d_641_v9_] if (d_641_v9_) in (d_642_v10_) else d_638_v6_)) + (_dafny.SeqWithoutIsStrInference([d_637_v5_])), (d_639_v7_) + (d_639_v7_))
        d_643_v11_: _dafny.Seq
        d_643_v11_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f26, p0, (d_633_v1_) == (len(d_641_v9_)), p0])
        d_643_v11_ = (d_643_v11_) + (d_643_v11_)
        d_644_v12_: _dafny.Seq
        d_644_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([_dafny.CodePoint('y')])])
        d_633_v1_ = (0) - ((self).fm13((d_644_v12_) + (d_644_v12_), (d_641_v9_)[default__.safeIndex(d_633_v1_, len(d_641_v9_))], (d_633_v1_) + (d_633_v1_), globalState))
        d_641_v9_ = (d_641_v9_) + (d_641_v9_)
        d_645_v13_: _dafny.Set
        d_645_v13_ = _dafny.Set({(self).f27})
        d_646_v14_: _dafny.Map
        d_646_v14_ = _dafny.Map({(0) - (len(d_645_v13_)): d_633_v1_})
        hi8_ = d_633_v1_
        for d_647_i0_ in range((0) - (((d_646_v14_)[d_633_v1_] if (d_633_v1_) in (d_646_v14_) else d_633_v1_)), hi8_):
            (globalState).f15 = (d_647_i0_) < (default__.safeDivisionInt(d_647_i0_, len(d_643_v11_)))
            d_648_v15_: _dafny.Map
            d_648_v15_ = _dafny.Map({d_647_i0_: (self).f27})
            d_649_v16_: _dafny.MultiSet
            d_649_v16_ = _dafny.MultiSet([d_633_v1_])
            d_650_v17_: _dafny.Map
            d_650_v17_ = _dafny.Map({(len(d_648_v15_)) + ((d_649_v16_).cardinality): d_641_v9_})
            d_651_v18_: str
            d_651_v18_ = _dafny.CodePoint('e')
            d_650_v17_ = (d_650_v17_).set(len(_dafny.SeqWithoutIsStrInference([default__.fm14((self).f26, d_647_i0_, d_633_v1_, globalState), d_651_v18_, d_651_v18_, d_651_v18_])), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_652_i1_ in range(default__.abs(934))])) + (d_641_v9_))
            d_633_v1_ = len(((_dafny.SeqWithoutIsStrInference([d_647_i0_]) if p0 else _dafny.SeqWithoutIsStrInference([d_647_i0_]))) + ((default__.fm15(globalState)) + (d_639_v7_)))
            d_653_v19_: T2
            nw81_ = C1()
            nw81_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylv"))), len(d_641_v9_), p0, (self).f27)
            d_653_v19_ = nw81_
            d_654_v20_: _dafny.Map
            d_654_v20_ = _dafny.Map({(D7_DC18(d_653_v19_)).cf35: d_647_i0_})
            d_655_v21_: _dafny.Array
            nw82_ = _dafny.Array(None, 12)
            nw82_[int(0)] = d_654_v20_
            nw82_[int(1)] = (d_654_v20_) | (_dafny.Map({d_653_v19_: len(d_636_v4_)}))
            nw82_[int(2)] = d_654_v20_
            nw82_[int(3)] = (d_654_v20_) | (_dafny.Map({d_653_v19_: d_633_v1_}))
            nw82_[int(4)] = d_654_v20_
            nw82_[int(5)] = d_654_v20_
            nw82_[int(6)] = (_dafny.Map({d_653_v19_: (d_649_v16_).cardinality})) | (d_654_v20_)
            nw82_[int(7)] = d_654_v20_
            nw82_[int(8)] = d_654_v20_
            nw82_[int(9)] = d_654_v20_
            nw82_[int(10)] = d_654_v20_
            nw82_[int(11)] = d_654_v20_
            d_655_v21_ = nw82_
            index81_ = default__.safeIndex(318, (d_655_v21_).length(0))
            (d_655_v21_)[index81_] = _dafny.Map({d_653_v19_: 196})
            index82_ = default__.safeIndex(318, (d_655_v21_).length(0))
            (d_655_v21_)[index82_] = (_dafny.Map({d_653_v19_: d_647_i0_})) | ((d_654_v20_) | (d_654_v20_))

    @property
    def f34(self):
        return self._f34

class C4(T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        self._f33: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f33, f26, f27):
        (self)._f33 = f33
        (self)._f26 = f26
        (self)._f27 = f27

    def fm1(self, globalState):
        source13_ = _dafny.Set({(self).f27})
        d_656___mcc_h0_ = source13_
        d_657_cf23_ = d_656___mcc_h0_
        return len(_dafny.SeqWithoutIsStrInference([not((self).f26), (self).f26, (self).f26]))

    def fm2(self, p0, p1, globalState):
        return (self).f27

    def m1(self, p0, p1, p2, globalState):
        d_658_v0_: int
        d_658_v0_ = 32
        (globalState).f12 = d_658_v0_
        hi9_ = d_658_v0_
        for d_659_i0_ in range(d_658_v0_, hi9_):
            if True:
                d_660_v1_: _dafny.MultiSet
                d_660_v1_ = _dafny.MultiSet([(0) - (d_659_i0_)])
                d_661_v2_: _dafny.MultiSet
                d_661_v2_ = _dafny.MultiSet([(_dafny.MultiSet([len(p2)])).ispropersubset(d_660_v1_), not ((self).f26) or (p0)])
                (globalState).f12 = ((d_661_v2_)[(self).f26] if ((self).f26) in (d_661_v2_) else (d_659_i0_) + (d_658_v0_))
                d_662_v3_: _dafny.Array
                def lambda36_(d_663_v0_):
                    def lambda37_(d_664_i1_):
                        return D2_DC7((self).f26, (self).f27, d_663_v0_, d_663_v0_)

                    return lambda37_

                init20_ = lambda36_(d_658_v0_)
                nw83_ = _dafny.Array(None, 13)
                for i0_20_ in range(nw83_.length(0)):
                    nw83_[i0_20_] = init20_(i0_20_)
                d_662_v3_ = nw83_
                d_662_v3_ = d_662_v3_
                d_665_v4_: _dafny.Map
                d_665_v4_ = _dafny.Map({d_658_v0_: 136})
                (globalState).f12 = (d_659_i0_) + (((d_665_v4_)[d_659_i0_] if (d_659_i0_) in (d_665_v4_) else (d_661_v2_).cardinality))
                d_666_v5_: C0
                nw84_ = C0()
                nw84_.ctor__((self).f27)
                d_666_v5_ = nw84_
                (globalState).f15 = not(not((d_658_v0_) >= (d_659_i0_)))
            elif True:
                d_667_v6_: _dafny.Map
                d_667_v6_ = _dafny.Map({(self).f27: p2})
                d_668_v7_: D5
                d_668_v7_ = D5_DC13(d_667_v6_)
                d_667_v6_ = ((d_668_v7_).cf24) | (d_667_v6_)
                d_669_v8_: D1
                d_669_v8_ = D1_DC3(p2)
                d_670_v9_: C0
                nw85_ = C0()
                nw85_.ctor__(p0)
                d_670_v9_ = nw85_
                d_671_v10_: _dafny.Seq
                d_671_v10_ = _dafny.SeqWithoutIsStrInference([d_670_v9_])
                d_672_v11_: _dafny.Map
                d_672_v11_ = _dafny.Map({d_669_v8_: (d_670_v9_) in (d_671_v10_)})
                d_673_v12_: _dafny.Map
                d_673_v12_ = _dafny.Map({(self).f26: (self).f27})
                d_674_v13_: _dafny.Seq
                d_674_v13_ = _dafny.SeqWithoutIsStrInference([d_672_v11_])
                d_675_v15_: _dafny.MultiSet
                d_675_v15_ = _dafny.MultiSet([d_669_v8_, D1_DC3(p2)])
                def iife63_():
                    coll51_ = _dafny.Map()
                    compr_51_: D1
                    for compr_51_ in (d_675_v15_).Elements:
                        d_676_v14_: D1 = compr_51_
                        if (d_676_v14_) in (d_675_v15_):
                            coll51_[d_676_v14_] = (d_670_v9_).f32
                    return _dafny.Map(coll51_)
                d_672_v11_ = (_dafny.Map({d_669_v8_: ((d_673_v12_)[(d_670_v9_).f32] if ((d_670_v9_).f32) in (d_673_v12_) else (self).f27)})) | (((d_674_v13_)[default__.safeIndex(d_659_i0_, len(d_674_v13_))]) | (iife63_()
                ))
                rhs86_ = d_658_v0_
                rhs87_ = d_658_v0_
                lhs78_ = globalState
                lhs79_ = globalState
                lhs78_.f12 = rhs86_
                lhs79_.f12 = rhs87_
                d_677_v16_: D5
                d_677_v16_ = D5_DC14(len(p2), d_658_v0_, True)
                d_678_v17_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_678_v17_ = nw86_
                d_679_v18_: _dafny.Seq
                d_679_v18_ = _dafny.SeqWithoutIsStrInference([d_678_v17_, d_678_v17_])
                (globalState).f12 = (((d_677_v16_).cf26) + (len(d_679_v18_))) - (d_658_v0_)
                d_680_v19_: _dafny.Seq
                d_680_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldus"))
                (globalState).f15 = (d_659_i0_) < ((len(d_680_v19_)) * (d_659_i0_))
            d_681_v20_: _dafny.Array
            def lambda38_(d_682_v0_):
                def lambda39_(d_683_i2_):
                    return (d_683_i2_) + ((D5_DC14(d_682_v0_, d_682_v0_, True)).cf26)

                return lambda39_

            init21_ = lambda38_(d_658_v0_)
            nw87_ = _dafny.Array(None, 9)
            for i0_21_ in range(nw87_.length(0)):
                nw87_[i0_21_] = init21_(i0_21_)
            d_681_v20_ = nw87_
            d_684_v21_: str
            d_684_v21_ = _dafny.CodePoint('o')
            d_685_v22_: _dafny.Set
            d_685_v22_ = _dafny.Set({d_684_v21_})
            d_686_v23_: _dafny.Map
            d_686_v23_ = _dafny.Map({d_685_v22_: d_658_v0_})
            d_687_v24_: _dafny.Seq
            d_687_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdyf"))
            d_688_v25_: _dafny.Seq
            d_688_v25_ = _dafny.SeqWithoutIsStrInference([d_658_v0_, d_659_i0_])
            d_689_v26_: _dafny.MultiSet
            d_689_v26_ = _dafny.MultiSet([d_688_v25_, d_688_v25_])
            d_690_v27_: T1
            nw88_ = C3()
            nw88_.ctor__(d_689_v26_, p0, p0)
            d_690_v27_ = nw88_
            d_691_v28_: _dafny.Array
            nw89_ = _dafny.Array(None, 14)
            nw89_[int(0)] = (d_659_i0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yghsv"))))
            nw89_[int(1)] = len((_dafny.Map({d_681_v20_: len(d_686_v23_)})) | (_dafny.Map({d_681_v20_: d_659_i0_})))
            nw89_[int(2)] = d_658_v0_
            nw89_[int(3)] = d_658_v0_
            nw89_[int(4)] = (0) - (d_658_v0_)
            nw89_[int(5)] = (0) - (d_659_i0_)
            nw89_[int(6)] = d_659_i0_
            nw89_[int(7)] = len((d_687_v24_) + (d_687_v24_))
            nw89_[int(8)] = len(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_658_v0_ for d_692_i3_ in range(default__.abs(546))]): d_690_v27_}))
            nw89_[int(9)] = d_659_i0_
            nw89_[int(10)] = d_658_v0_
            nw89_[int(11)] = d_659_i0_
            nw89_[int(12)] = d_659_i0_
            nw89_[int(13)] = d_658_v0_
            d_691_v28_ = nw89_
            d_693_v29_: D3
            d_693_v29_ = D3_DC10(d_691_v28_)
            d_681_v20_ = (d_693_v29_).cf21
            d_694_v30_: D7
            d_694_v30_ = D7_DC19(d_659_i0_, _dafny.MultiSet([not(p0), (self).f26]), False, _dafny.Map({False: d_658_v0_}))
            d_695_v31_: _dafny.MultiSet
            d_695_v31_ = _dafny.MultiSet([(d_690_v27_).f27])
            pat_let_tv2_ = d_684_v21_
            pat_let_tv3_ = d_695_v31_
            d_696_v32_: _dafny.Array
            nw90_ = _dafny.Array(None, 12)
            nw90_[int(0)] = (d_694_v30_ if (self).f27 else d_694_v30_)
            def iife64_(_pat_let6_0):
                def iife65_(d_697_dt__update__tmp_h0_):
                    def iife66_(_pat_let7_0):
                        def iife67_(d_699_dt__update_hcf36_h0_):
                            def iife68_(_pat_let8_0):
                                def iife69_(d_700_dt__update_hcf38_h0_):
                                    def iife70_(_pat_let9_0):
                                        def iife71_(d_701_dt__update_hcf37_h0_):
                                            return D7_DC19(d_699_dt__update_hcf36_h0_, d_701_dt__update_hcf37_h0_, d_700_dt__update_hcf38_h0_, (d_697_dt__update__tmp_h0_).cf39)
                                        return iife71_(_pat_let9_0)
                                    return iife70_(pat_let_tv3_)
                                return iife69_(_pat_let8_0)
                            return iife68_((self).f26)
                        return iife67_(_pat_let7_0)
                    return iife66_(len(_dafny.SeqWithoutIsStrInference([pat_let_tv2_ for d_698_i4_ in range(default__.abs(230))])))
                return iife65_(_pat_let6_0)
            nw90_[int(1)] = iife64_(d_694_v30_)
            nw90_[int(2)] = d_694_v30_
            nw90_[int(3)] = d_694_v30_
            nw90_[int(4)] = d_694_v30_
            nw90_[int(5)] = d_694_v30_
            nw90_[int(6)] = d_694_v30_
            nw90_[int(7)] = d_694_v30_
            nw90_[int(8)] = d_694_v30_
            nw90_[int(9)] = d_694_v30_
            nw90_[int(10)] = (d_694_v30_ if p0 else d_694_v30_)
            nw90_[int(11)] = d_694_v30_
            d_696_v32_ = nw90_
            d_702_v33_: _dafny.Map
            d_702_v33_ = _dafny.Map({(d_690_v27_).f26: 637})
            index83_ = default__.safeIndex(896, (d_696_v32_).length(0))
            (d_696_v32_)[index83_] = D7_DC19(d_658_v0_, d_695_v31_, p0, d_702_v33_)
            index84_ = default__.safeIndex(896, (d_696_v32_).length(0))
            (d_696_v32_)[index84_] = d_694_v30_
            d_703_v34_: _dafny.Array
            nw91_ = _dafny.Array(None, 4)
            nw91_[int(0)] = (d_690_v27_).f26
            nw91_[int(1)] = (d_690_v27_).f27
            nw91_[int(2)] = False
            nw91_[int(3)] = (d_690_v27_).f26
            d_703_v34_ = nw91_
            index85_ = default__.safeIndex(510, (d_703_v34_).length(0))
            (d_703_v34_)[index85_] = ((d_690_v27_).fm2((d_690_v27_).f26, d_659_i0_, globalState) if True else (d_690_v27_).f27)
            d_704_v35_: _dafny.Array
            nw92_ = _dafny.Array(D1.default()(), 29)
            d_704_v35_ = nw92_
            d_705_v36_: _dafny.Map
            d_705_v36_ = _dafny.Map({True: d_704_v35_})
            d_706_v37_: _dafny.Map
            d_706_v37_ = _dafny.Map({d_659_i0_: d_659_i0_})
            d_707_v38_: _dafny.Map
            d_707_v38_ = _dafny.Map({d_705_v36_: len(d_706_v37_)})
            index86_ = default__.safeIndex(510, (d_703_v34_).length(0))
            rhs88_ = not ((d_658_v0_) != (d_659_i0_)) or ((self).f26)
            rhs89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxf"))
            rhs90_ = default__.safeModuloInt(d_658_v0_, d_658_v0_)
            rhs91_ = ((d_658_v0_) != (((d_707_v38_)[d_705_v36_] if (d_705_v36_) in (d_707_v38_) else (d_688_v25_)[default__.safeIndex(d_659_i0_, len(d_688_v25_))]))) or ((d_659_i0_) <= (d_658_v0_))
            lhs80_ = d_703_v34_
            lhs81_ = default__.safeIndex(510, (d_703_v34_).length(0))
            lhs82_ = globalState
            lhs83_ = globalState
            lhs84_ = globalState
            lhs80_[lhs81_] = rhs88_
            lhs82_.f6 = rhs89_
            lhs83_.f12 = rhs90_
            lhs84_.f22 = rhs91_
        d_708_v39_: str
        d_708_v39_ = _dafny.CodePoint('q')
        d_709_v40_: _dafny.Set
        d_709_v40_ = _dafny.Set({(self).f26})
        d_710_v41_: _dafny.Array
        nw93_ = _dafny.Array(None, 21)
        nw93_[int(0)] = d_708_v39_
        nw93_[int(1)] = _dafny.CodePoint('f')
        nw93_[int(2)] = _dafny.CodePoint('l')
        nw93_[int(3)] = default__.fm14(p0, d_658_v0_, len(d_709_v40_), globalState)
        nw93_[int(4)] = d_708_v39_
        nw93_[int(5)] = d_708_v39_
        nw93_[int(6)] = _dafny.CodePoint('v')
        nw93_[int(7)] = d_708_v39_
        nw93_[int(8)] = d_708_v39_
        nw93_[int(9)] = d_708_v39_
        nw93_[int(10)] = d_708_v39_
        nw93_[int(11)] = d_708_v39_
        nw93_[int(12)] = d_708_v39_
        nw93_[int(13)] = d_708_v39_
        nw93_[int(14)] = _dafny.CodePoint('u')
        nw93_[int(15)] = d_708_v39_
        nw93_[int(16)] = _dafny.CodePoint('o')
        nw93_[int(17)] = d_708_v39_
        nw93_[int(18)] = d_708_v39_
        nw93_[int(19)] = _dafny.CodePoint('m')
        nw93_[int(20)] = d_708_v39_
        d_710_v41_ = nw93_
        d_711_v42_: _dafny.Map
        d_711_v42_ = _dafny.Map({d_658_v0_: False})
        d_712_v43_: _dafny.Seq
        d_712_v43_ = _dafny.SeqWithoutIsStrInference([d_711_v42_, d_711_v42_])
        d_713_v44_: D2
        d_713_v44_ = D2_DC7(True, False, d_658_v0_, (_dafny.MultiSet(d_712_v43_)).cardinality)
        d_714_v45_: _dafny.Map
        d_714_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_713_v44_, d_713_v44_]): d_658_v0_})
        d_715_v46_: _dafny.Seq
        d_715_v46_ = _dafny.SeqWithoutIsStrInference([d_713_v44_, d_713_v44_])
        d_716_v47_: _dafny.Seq
        d_716_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baa"))
        d_717_v48_: D1
        d_717_v48_ = D1_DC4(d_710_v41_, d_658_v0_, ((d_714_v45_)[d_715_v46_] if (d_715_v46_) in (d_714_v45_) else d_658_v0_), d_716_v47_, (d_716_v47_).set(default__.safeIndex((0) - (len(_dafny.Map({d_658_v0_: d_658_v0_}))), len(d_716_v47_)), d_708_v39_))
        source14_ = d_717_v48_
        if source14_.is_DC4:
            d_718___mcc_h0_ = source14_.cf7
            d_719___mcc_h1_ = source14_.cf8
            d_720___mcc_h2_ = source14_.cf9
            d_721___mcc_h3_ = source14_.cf10
            d_722___mcc_h4_ = source14_.cf11
            d_723_cf11_ = d_722___mcc_h4_
            d_724_cf10_ = d_721___mcc_h3_
            d_725_cf9_ = d_720___mcc_h2_
            d_726_cf8_ = d_719___mcc_h1_
            d_727_cf7_ = d_718___mcc_h0_
            d_728_v49_: _dafny.Seq
            d_728_v49_ = _dafny.SeqWithoutIsStrInference([d_658_v0_, d_726_cf8_, d_725_cf9_])
            d_729_v50_: _dafny.Seq
            d_729_v50_ = _dafny.SeqWithoutIsStrInference([d_728_v49_, d_728_v49_, (d_728_v49_).set(default__.safeIndex(len(d_728_v49_), len(d_728_v49_)), d_726_cf8_), _dafny.SeqWithoutIsStrInference([d_658_v0_]), d_728_v49_])
            d_730_v51_: C2
            nw94_ = C2()
            nw94_.ctor__(d_729_v50_, False, (self).f26)
            d_730_v51_ = nw94_
            d_731_v52_: _dafny.Seq
            d_731_v52_ = _dafny.SeqWithoutIsStrInference([d_730_v51_])
            (globalState).f12 = (default__.safeDivisionInt(len((p2).set(default__.safeIndex((0) - ((self).fm1(globalState)), len(p2)), not((self).f26))), d_725_cf9_) if (self).f26 else len((d_731_v52_) + (d_731_v52_)))
            d_732_v53_: _dafny.Map
            d_732_v53_ = _dafny.Map({len(p2): d_726_cf8_})
            d_733_v54_: D0
            d_733_v54_ = D0_DC0((self).f26)
            d_732_v53_ = (d_732_v53_).set(d_726_cf8_, (d_730_v51_).fm4(d_733_v54_, d_658_v0_, d_725_cf9_, globalState))
            (globalState).f2 = not((self).f26)
            if not((self).f27):
                d_734_v55_: _dafny.Array
                nw95_ = _dafny.Array(_dafny.Seq({}), 2)
                d_734_v55_ = nw95_
                index87_ = default__.safeIndex(59, (d_734_v55_).length(0))
                (d_734_v55_)[index87_] = (default__.fm15(globalState)) + (_dafny.SeqWithoutIsStrInference([d_726_cf8_ for d_735_i5_ in range(default__.abs(191))]))
                index88_ = default__.safeIndex(59, (d_734_v55_).length(0))
                rhs92_ = (0) - ((d_730_v51_).fm4(d_733_v54_, d_725_cf9_, d_658_v0_, globalState))
                rhs93_ = (p2)[default__.safeIndex(d_658_v0_, len(p2))]
                rhs94_ = d_724_cf10_
                rhs95_ = default__.fm23(d_708_v39_, globalState)
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = d_734_v55_
                lhs88_ = default__.safeIndex(59, (d_734_v55_).length(0))
                d_725_cf9_ = rhs92_
                lhs85_.f15 = rhs93_
                lhs86_.f6 = rhs94_
                lhs87_[lhs88_] = rhs95_
                (globalState).f19 = d_708_v39_
                (globalState).f6 = ((d_716_v47_).set(default__.safeIndex((0) - (d_658_v0_), len(d_716_v47_)), d_708_v39_)) + (_dafny.SeqWithoutIsStrInference([d_708_v39_ for d_736_i6_ in range(default__.abs(472))]))
                d_737_v56_: _dafny.Set
                d_737_v56_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "riqaxgwa")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jqtgis"))})
                (globalState).f24 = (d_730_v51_).fm2((d_737_v56_).ispropersubset(d_737_v56_), d_725_cf9_, globalState)
                (globalState).f15 = (self).f26
            elif True:
                d_738_v57_: _dafny.Array
                nw96_ = _dafny.Array(_dafny.Map({}), 6)
                d_738_v57_ = nw96_
                d_739_v58_: _dafny.Array
                nw97_ = _dafny.Array(D3.default()(), 11)
                d_739_v58_ = nw97_
                d_740_v59_: _dafny.Map
                d_740_v59_ = _dafny.Map({not((self).f26): d_727_cf7_})
                d_741_v60_: _dafny.Seq
                d_741_v60_ = _dafny.SeqWithoutIsStrInference([d_727_cf7_])
                d_742_v61_: _dafny.Array
                nw98_ = _dafny.Array(None, 15)
                nw98_[int(0)] = d_710_v41_
                nw98_[int(1)] = d_727_cf7_
                nw98_[int(2)] = d_727_cf7_
                nw98_[int(3)] = d_727_cf7_
                nw98_[int(4)] = d_727_cf7_
                nw98_[int(5)] = d_710_v41_
                nw98_[int(6)] = d_710_v41_
                nw98_[int(7)] = d_727_cf7_
                nw98_[int(8)] = d_710_v41_
                nw98_[int(9)] = d_727_cf7_
                nw98_[int(10)] = (d_727_cf7_ if True else d_710_v41_)
                nw98_[int(11)] = d_727_cf7_
                nw98_[int(12)] = ((d_740_v59_)[True] if (True) in (d_740_v59_) else d_727_cf7_)
                nw98_[int(13)] = d_727_cf7_
                nw98_[int(14)] = (d_741_v60_)[default__.safeIndex(-535, len(d_741_v60_))]
                d_742_v61_ = nw98_
                index89_ = default__.safeIndex(966, (d_742_v61_).length(0))
                (d_742_v61_)[index89_] = d_710_v41_
                index90_ = default__.safeIndex(966, (d_742_v61_).length(0))
                rhs96_ = (d_738_v57_ if p0 else d_738_v57_)
                rhs97_ = d_739_v58_
                rhs98_ = d_710_v41_
                lhs89_ = d_742_v61_
                lhs90_ = default__.safeIndex(966, (d_742_v61_).length(0))
                d_738_v57_ = rhs96_
                d_739_v58_ = rhs97_
                lhs89_[lhs90_] = rhs98_
                d_725_cf9_ = (d_725_cf9_ if (self).f27 else (len(d_724_cf10_)) * (d_725_cf9_))
                d_743_v62_: _dafny.Set
                d_743_v62_ = _dafny.Set({-955})
                d_744_v63_: D2
                d_744_v63_ = D2_DC9(len(d_743_v62_), d_658_v0_, False)
                d_745_v64_: _dafny.MultiSet
                d_745_v64_ = _dafny.MultiSet([not((self).f27), (self).f26, (self).f26])
                d_746_v65_: _dafny.Seq
                d_746_v65_ = _dafny.SeqWithoutIsStrInference([p2])
                d_747_v66_: _dafny.Array
                nw99_ = _dafny.Array(None, 12)
                nw99_[int(0)] = d_726_cf8_
                nw99_[int(1)] = (d_744_v63_).cf18
                nw99_[int(2)] = (d_745_v64_).cardinality
                nw99_[int(3)] = (0) - (d_725_cf9_)
                nw99_[int(4)] = (_dafny.MultiSet([(self).f27, p0, p0, False])).cardinality
                nw99_[int(5)] = d_658_v0_
                nw99_[int(6)] = d_726_cf8_
                nw99_[int(7)] = d_726_cf8_
                nw99_[int(8)] = len(d_723_cf11_)
                nw99_[int(9)] = d_658_v0_
                nw99_[int(10)] = len((p2) + ((d_746_v65_)[default__.safeIndex((d_730_v51_).fm1(globalState), len(d_746_v65_))]))
                nw99_[int(11)] = d_658_v0_
                d_747_v66_ = nw99_
                index91_ = default__.safeIndex(232, (d_747_v66_).length(0))
                (d_747_v66_)[index91_] = (d_658_v0_) * (d_725_cf9_)
                index92_ = default__.safeIndex(232, (d_747_v66_).length(0))
                rhs99_ = d_725_cf9_
                rhs100_ = d_747_v66_
                rhs101_ = not (not((p0 if (self).f26 else ((d_711_v42_)[d_726_cf8_] if (d_726_cf8_) in (d_711_v42_) else p0)))) or (((0) - (d_725_cf9_)) < (d_726_cf8_))
                rhs102_ = p0
                rhs103_ = default__.safeModuloInt(d_725_cf9_, d_658_v0_)
                lhs91_ = globalState
                lhs92_ = globalState
                lhs93_ = globalState
                lhs94_ = d_747_v66_
                lhs95_ = default__.safeIndex(232, (d_747_v66_).length(0))
                lhs91_.f16 = rhs99_
                d_747_v66_ = rhs100_
                lhs92_.f15 = rhs101_
                lhs93_.f2 = rhs102_
                lhs94_[lhs95_] = rhs103_
                d_748_v67_: _dafny.Seq
                d_748_v67_ = _dafny.SeqWithoutIsStrInference([p0])
                d_748_v67_ = d_748_v67_
                d_749_v68_: D7
                d_749_v68_ = D7_DC19(d_658_v0_, (d_745_v64_).set((self).f27, default__.abs(d_725_cf9_)), (self).f26, _dafny.Map({p0: d_725_cf9_}))
                d_750_v69_: _dafny.MultiSet
                d_750_v69_ = _dafny.MultiSet([d_726_cf8_, (0) - (d_725_cf9_)])
                d_751_v70_: _dafny.Map
                d_751_v70_ = _dafny.Map({(self).f26: (d_730_v51_).fm2(p0, d_658_v0_, globalState)})
                d_752_v71_: _dafny.Array
                nw100_ = _dafny.Array(None, 26)
                nw100_[int(0)] = False
                nw100_[int(1)] = False
                nw100_[int(2)] = (len(default__.fm22(((d_732_v53_)[d_726_cf8_] if (d_726_cf8_) in (d_732_v53_) else d_725_cf9_), (self).f27, globalState))) > (d_726_cf8_)
                nw100_[int(3)] = (self).f27
                nw100_[int(4)] = (self).f27
                nw100_[int(5)] = (d_749_v68_).cf38
                nw100_[int(6)] = (self).f27
                nw100_[int(7)] = p0
                nw100_[int(8)] = (self).f27
                nw100_[int(9)] = (d_716_v47_) <= (d_716_v47_)
                nw100_[int(10)] = (_dafny.MultiSet(d_728_v49_)).isdisjoint(d_750_v69_)
                nw100_[int(11)] = (self).f27
                nw100_[int(12)] = (self).f27
                nw100_[int(13)] = True
                nw100_[int(14)] = (d_724_cf10_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")))
                nw100_[int(15)] = ((self).f26) and (p0)
                nw100_[int(16)] = (self).f26
                nw100_[int(17)] = (self).f27
                nw100_[int(18)] = (self).f26
                nw100_[int(19)] = ((d_751_v70_)[(self).f26] if ((self).f26) in (d_751_v70_) else (self).f27)
                nw100_[int(20)] = True
                nw100_[int(21)] = False
                nw100_[int(22)] = (self).f27
                nw100_[int(23)] = (self).f27
                nw100_[int(24)] = not((_dafny.Map({(self).f26: (self).f27})) == (d_751_v70_))
                nw100_[int(25)] = not(p0)
                d_752_v71_ = nw100_
                index93_ = default__.safeIndex(0, (d_752_v71_).length(0))
                (d_752_v71_)[index93_] = False
                d_753_v72_: _dafny.Map
                d_753_v72_ = _dafny.Map({(self).f27: (d_747_v66_)[default__.safeIndex(232, (d_747_v66_).length(0))]})
                index94_ = default__.safeIndex(0, (d_752_v71_).length(0))
                rhs104_ = p0
                rhs105_ = (self).fm1(globalState)
                rhs106_ = d_708_v39_
                rhs107_ = ((p0 if (d_730_v51_).fm2(not((self).f27), (d_747_v66_)[default__.safeIndex(232, (d_747_v66_).length(0))], globalState) else (self).f27)) or ((len(d_753_v72_)) != (d_658_v0_))
                lhs96_ = globalState
                lhs97_ = globalState
                lhs98_ = globalState
                lhs99_ = d_752_v71_
                lhs100_ = default__.safeIndex(0, (d_752_v71_).length(0))
                lhs96_.f22 = rhs104_
                lhs97_.f16 = rhs105_
                lhs98_.f19 = rhs106_
                lhs99_[lhs100_] = rhs107_
        elif source14_.is_DC5:
            d_754___mcc_h5_ = source14_.cf12
            d_755_cf12_ = d_754___mcc_h5_
            if (self).f26:
                d_756_v73_: _dafny.Map
                d_756_v73_ = _dafny.Map({d_658_v0_: d_708_v39_})
                d_757_v74_: _dafny.Seq
                d_757_v74_ = _dafny.SeqWithoutIsStrInference([d_756_v73_])
                (globalState).f15 = (len((d_757_v74_) + (d_757_v74_))) >= (d_658_v0_)
                d_758_v75_: _dafny.MultiSet
                d_758_v75_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_755_cf12_])])
                d_759_v76_: C3
                nw101_ = C3()
                nw101_.ctor__(d_758_v75_, (p2)[default__.safeIndex(d_755_cf12_, len(p2))], (self).f27)
                d_759_v76_ = nw101_
                d_760_v77_: _dafny.Seq
                d_760_v77_ = _dafny.SeqWithoutIsStrInference([d_755_cf12_, d_658_v0_])
                d_761_v78_: _dafny.Map
                d_761_v78_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_658_v0_, d_658_v0_, (0) - (d_658_v0_)])) for d_762_i7_ in range(default__.abs(-971))])): (d_760_v77_)[default__.safeIndex((d_759_v76_).fm12(len(d_716_v47_), globalState), len(d_760_v77_))]})
                d_763_v79_: _dafny.Map
                d_763_v79_ = _dafny.Map({d_759_v76_: len(d_761_v78_)})
                d_764_v80_: _dafny.Map
                d_764_v80_ = _dafny.Map({d_763_v79_: d_708_v39_})
                d_765_v81_: _dafny.Map
                d_765_v81_ = _dafny.Map({d_658_v0_: d_761_v78_})
                d_766_v82_: _dafny.Array
                nw102_ = _dafny.Array(None, 19)
                nw102_[int(0)] = (p2) + ((p2).set(default__.safeIndex(d_755_cf12_, len(p2)), (self).f26))
                nw102_[int(1)] = p2
                nw102_[int(2)] = _dafny.SeqWithoutIsStrInference([not(p0)])
                nw102_[int(3)] = p2
                nw102_[int(4)] = default__.fm25(True, (self).f26, d_708_v39_, globalState)
                nw102_[int(5)] = p2
                nw102_[int(6)] = p2
                nw102_[int(7)] = (p2) + (p2)
                nw102_[int(8)] = default__.fm25((self).f27, (self).f26, ((d_764_v80_)[d_763_v79_] if (d_763_v79_) in (d_764_v80_) else d_708_v39_), globalState)
                nw102_[int(9)] = p2
                nw102_[int(10)] = p2
                nw102_[int(11)] = p2
                nw102_[int(12)] = ((p2 if (self).f26 else _dafny.SeqWithoutIsStrInference([False]))).set(default__.safeIndex(len(_dafny.Map({(d_759_v76_).fm12(d_755_cf12_, globalState): ((d_765_v81_)[(_dafny.MultiSet(default__.fm15(globalState))).cardinality] if ((_dafny.MultiSet(default__.fm15(globalState))).cardinality) in (d_765_v81_) else _dafny.Map({d_755_cf12_: d_658_v0_}))})), len((p2 if (self).f26 else _dafny.SeqWithoutIsStrInference([False])))), (self).f27)
                nw102_[int(13)] = p2
                nw102_[int(14)] = p2
                nw102_[int(15)] = (p2).set(default__.safeIndex(d_755_cf12_, len(p2)), (self).f27)
                nw102_[int(16)] = p2
                nw102_[int(17)] = p2
                nw102_[int(18)] = default__.fm25((self).f26, (self).f27, d_708_v39_, globalState)
                d_766_v82_ = nw102_
                index95_ = default__.safeIndex(171, (d_766_v82_).length(0))
                (d_766_v82_)[index95_] = p2
                index96_ = default__.safeIndex(171, (d_766_v82_).length(0))
                (d_766_v82_)[index96_] = p2
                d_767_v83_: _dafny.Array
                nw103_ = _dafny.Array(int(0), 20)
                d_767_v83_ = nw103_
                d_768_v84_: D3
                d_768_v84_ = D3_DC10(d_767_v83_)
                d_769_v85_: _dafny.Seq
                d_769_v85_ = _dafny.SeqWithoutIsStrInference([(d_768_v84_).cf21])
                d_770_v86_: D2
                d_770_v86_ = D2_DC9(121, d_658_v0_, False)
                d_771_v87_: _dafny.Array
                nw104_ = _dafny.Array(None, 27)
                nw104_[int(0)] = d_767_v83_
                nw104_[int(1)] = (d_769_v85_)[default__.safeIndex(d_658_v0_, len(d_769_v85_))]
                nw104_[int(2)] = d_767_v83_
                nw104_[int(3)] = d_767_v83_
                nw104_[int(4)] = d_767_v83_
                nw104_[int(5)] = d_767_v83_
                nw104_[int(6)] = d_767_v83_
                nw104_[int(7)] = d_767_v83_
                nw104_[int(8)] = d_767_v83_
                nw104_[int(9)] = d_767_v83_
                nw104_[int(10)] = d_767_v83_
                nw104_[int(11)] = (d_767_v83_ if (d_770_v86_).cf20 else d_767_v83_)
                nw104_[int(12)] = d_767_v83_
                nw104_[int(13)] = d_767_v83_
                nw104_[int(14)] = d_767_v83_
                nw104_[int(15)] = d_767_v83_
                nw104_[int(16)] = (d_767_v83_ if (self).f26 else d_767_v83_)
                nw104_[int(17)] = d_767_v83_
                nw104_[int(18)] = d_767_v83_
                nw104_[int(19)] = d_767_v83_
                nw104_[int(20)] = d_767_v83_
                nw104_[int(21)] = d_767_v83_
                nw104_[int(22)] = d_767_v83_
                nw104_[int(23)] = (d_768_v84_).cf21
                nw104_[int(24)] = d_767_v83_
                nw104_[int(25)] = d_767_v83_
                nw104_[int(26)] = d_767_v83_
                d_771_v87_ = nw104_
                d_771_v87_ = d_771_v87_
                d_772_v88_: _dafny.MultiSet
                d_772_v88_ = _dafny.MultiSet([(self).f27, (self).f27])
                d_773_v89_: _dafny.Array
                nw105_ = _dafny.Array(None, 25)
                nw105_[int(0)] = not (not(not(p0))) or (not(p0))
                nw105_[int(1)] = (self).f27
                nw105_[int(2)] = (self).f26
                nw105_[int(3)] = p0
                nw105_[int(4)] = not(p0)
                nw105_[int(5)] = p0
                nw105_[int(6)] = (self).f27
                nw105_[int(7)] = (self).f26
                nw105_[int(8)] = p0
                nw105_[int(9)] = False
                nw105_[int(10)] = (self).f27
                nw105_[int(11)] = (self).f27
                nw105_[int(12)] = (self).f27
                nw105_[int(13)] = not((d_772_v88_).isdisjoint(_dafny.MultiSet([(self).fm2(p0, d_658_v0_, globalState)])))
                nw105_[int(14)] = p0
                nw105_[int(15)] = (self).f26
                nw105_[int(16)] = p0
                nw105_[int(17)] = p0
                nw105_[int(18)] = not((self).f26)
                nw105_[int(19)] = (self).f27
                nw105_[int(20)] = (self).f27
                nw105_[int(21)] = (self).f27
                nw105_[int(22)] = (p0) and (False)
                nw105_[int(23)] = (self).f27
                nw105_[int(24)] = (len((d_766_v82_)[default__.safeIndex(171, (d_766_v82_).length(0))])) < (d_658_v0_)
                d_773_v89_ = nw105_
                index97_ = default__.safeIndex(422, (d_773_v89_).length(0))
                (d_773_v89_)[index97_] = (self).fm2(p0, (0) - (d_755_cf12_), globalState)
                index98_ = default__.safeIndex(422, (d_773_v89_).length(0))
                (d_773_v89_)[index98_] = p0
                d_774_v90_: _dafny.Seq
                d_774_v90_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).fm2(p0, d_658_v0_, globalState), (d_773_v89_)[default__.safeIndex(422, (d_773_v89_).length(0))]]), p2, (d_766_v82_)[default__.safeIndex(171, (d_766_v82_).length(0))]])
                d_775_v91_: _dafny.Array
                nw106_ = _dafny.Array(_dafny.Map({}), 9)
                d_775_v91_ = nw106_
                index99_ = default__.safeIndex(651, (d_775_v91_).length(0))
                (d_775_v91_)[index99_] = d_756_v73_
                index100_ = default__.safeIndex(422, (d_773_v89_).length(0))
                index101_ = default__.safeIndex(651, (d_775_v91_).length(0))
                rhs108_ = d_774_v90_
                rhs109_ = True
                rhs110_ = ((d_761_v78_)[d_755_cf12_] if (d_755_cf12_) in (d_761_v78_) else 953)
                rhs111_ = d_756_v73_
                lhs101_ = d_773_v89_
                lhs102_ = default__.safeIndex(422, (d_773_v89_).length(0))
                lhs103_ = globalState
                lhs104_ = d_775_v91_
                lhs105_ = default__.safeIndex(651, (d_775_v91_).length(0))
                d_774_v90_ = rhs108_
                lhs101_[lhs102_] = rhs109_
                lhs103_.f12 = rhs110_
                lhs104_[lhs105_] = rhs111_
            elif True:
                d_776_v92_: _dafny.Map
                d_776_v92_ = _dafny.Map({True: d_658_v0_})
                d_777_v93_: _dafny.Seq
                d_777_v93_ = _dafny.SeqWithoutIsStrInference([d_776_v92_])
                d_778_v94_: _dafny.Array
                nw107_ = _dafny.Array(None, 26)
                nw107_[int(0)] = (self).f27
                nw107_[int(1)] = p0
                nw107_[int(2)] = ((d_716_v47_).set(default__.safeIndex(d_755_cf12_, len(d_716_v47_)), _dafny.CodePoint('e'))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urnk")))
                nw107_[int(3)] = not (p0) or ((self).fm2((self).f27, d_658_v0_, globalState))
                nw107_[int(4)] = (self).f27
                nw107_[int(5)] = (self).fm2(p0, d_755_cf12_, globalState)
                nw107_[int(6)] = (self).f27
                nw107_[int(7)] = not((self).f27)
                nw107_[int(8)] = not((D5_DC14(d_658_v0_, d_755_cf12_, (self).f26)).cf27)
                nw107_[int(9)] = p0
                nw107_[int(10)] = ((0) - (d_755_cf12_)) >= (len(p2))
                nw107_[int(11)] = p0
                nw107_[int(12)] = (self).f26
                nw107_[int(13)] = (self).f27
                nw107_[int(14)] = p0
                nw107_[int(15)] = p0
                nw107_[int(16)] = (self).f26
                nw107_[int(17)] = (self).f27
                nw107_[int(18)] = (p2)[default__.safeIndex(d_658_v0_, len(p2))]
                nw107_[int(19)] = (self).f26
                nw107_[int(20)] = p0
                nw107_[int(21)] = (self).f26
                nw107_[int(22)] = True
                nw107_[int(23)] = p0
                nw107_[int(24)] = (self).f27
                nw107_[int(25)] = (d_777_v93_) != (d_777_v93_)
                d_778_v94_ = nw107_
                index102_ = default__.safeIndex(26, (d_778_v94_).length(0))
                (d_778_v94_)[index102_] = (p2)[default__.safeIndex(d_658_v0_, len(p2))]
                index103_ = default__.safeIndex(26, (d_778_v94_).length(0))
                (d_778_v94_)[index103_] = (self).f27
                d_779_v95_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 8)
                d_779_v95_ = nw108_
                d_780_v96_: _dafny.Map
                d_780_v96_ = _dafny.Map({d_779_v95_: (d_778_v94_)[default__.safeIndex(26, (d_778_v94_).length(0))]})
                (globalState).f22 = (len((d_780_v96_) | (_dafny.Map({d_779_v95_: p0})))) < (d_755_cf12_)
                d_781_v97_: C1
                nw109_ = C1()
                nw109_.ctor__(default__.safeModuloInt(len((d_716_v47_).set(default__.safeIndex(d_658_v0_, len(d_716_v47_)), d_708_v39_)), d_658_v0_), (self).fm1(globalState), (d_778_v94_)[default__.safeIndex(26, (d_778_v94_).length(0))], (self).f27)
                d_781_v97_ = nw109_
                d_782_v98_: D3
                d_782_v98_ = D3_DC11(d_708_v39_)
                d_783_v99_: _dafny.Seq
                d_783_v99_ = _dafny.SeqWithoutIsStrInference([len(d_709_v40_), (d_781_v97_).f36])
                d_784_v100_: _dafny.Map
                d_784_v100_ = _dafny.Map({d_778_v94_: (default__.fm23((d_782_v98_).cf22, globalState)) <= (d_783_v99_)})
                d_784_v100_ = (d_784_v100_).set(d_778_v94_, not(False))
                d_785_v101_: T2
                nw110_ = C1()
                nw110_.ctor__(len(_dafny.SeqWithoutIsStrInference([False, (self).f26])), d_781_v97_.f37, (d_778_v94_)[default__.safeIndex(26, (d_778_v94_).length(0))], p0)
                d_785_v101_ = nw110_
                d_786_v102_: D7
                d_786_v102_ = D7_DC18(d_785_v101_)
                index104_ = default__.safeIndex(63, (d_779_v95_).length(0))
                (d_779_v95_)[index104_] = d_658_v0_
                index105_ = default__.safeIndex(63, (d_779_v95_).length(0))
                index106_ = default__.safeIndex(26, (d_778_v94_).length(0))
                rhs112_ = d_786_v102_
                rhs113_ = d_781_v97_.f37
                rhs114_ = (self).f26
                rhs115_ = (_dafny.SeqWithoutIsStrInference([d_658_v0_])) < (d_783_v99_)
                rhs116_ = d_785_v101_
                lhs106_ = d_779_v95_
                lhs107_ = default__.safeIndex(63, (d_779_v95_).length(0))
                lhs108_ = d_778_v94_
                lhs109_ = default__.safeIndex(26, (d_778_v94_).length(0))
                lhs110_ = globalState
                d_786_v102_ = rhs112_
                lhs106_[lhs107_] = rhs113_
                lhs108_[lhs109_] = rhs114_
                lhs110_.f2 = rhs115_
                d_785_v101_ = rhs116_
            if not(not (True) or (p0)):
                (globalState).f24 = ((self).f27) not in (p2)
                (globalState).f22 = (self).f26
                rhs117_ = (0) - ((0) - (d_658_v0_))
                rhs118_ = _dafny.CodePoint('y')
                rhs119_ = 687
                lhs111_ = globalState
                lhs112_ = globalState
                lhs111_.f16 = rhs117_
                d_708_v39_ = rhs118_
                lhs112_.f16 = rhs119_
                (globalState).f15 = ((self).f26) or (p0)
                d_787_v103_: _dafny.Array
                def lambda40_(d_788_cf12_, d_789_v0_):
                    def lambda41_(d_790_i8_):
                        return _dafny.MultiSet([_dafny.Map({d_788_cf12_: d_789_v0_})])

                    return lambda41_

                init22_ = lambda40_(d_755_cf12_, d_658_v0_)
                nw111_ = _dafny.Array(None, 29)
                for i0_22_ in range(nw111_.length(0)):
                    nw111_[i0_22_] = init22_(i0_22_)
                d_787_v103_ = nw111_
                index107_ = default__.safeIndex(369, (d_787_v103_).length(0))
                (d_787_v103_)[index107_] = default__.fm26((self).fm2((self).f26, d_755_cf12_, globalState), p0, p0, globalState)
                d_791_v104_: _dafny.Map
                d_791_v104_ = _dafny.Map({d_658_v0_: d_755_cf12_})
                d_792_v105_: _dafny.MultiSet
                d_792_v105_ = _dafny.MultiSet([_dafny.Map({d_755_cf12_: d_658_v0_}), (d_791_v104_) | (_dafny.Map({-261: d_755_cf12_})), d_791_v104_])
                index108_ = default__.safeIndex(369, (d_787_v103_).length(0))
                (d_787_v103_)[index108_] = d_792_v105_
            elif True:
                d_793_v106_: _dafny.MultiSet
                d_793_v106_ = _dafny.MultiSet([not((self).fm2(p0, (self).fm1(globalState), globalState))])
                rhs120_ = not (not(not((self).f27))) or ((d_658_v0_) == (d_755_cf12_))
                rhs121_ = d_793_v106_
                lhs113_ = globalState
                lhs113_.f2 = rhs120_
                d_793_v106_ = rhs121_
                d_794_v107_: _dafny.Array
                def lambda42_(d_795_p2_):
                    def lambda43_(d_796_i9_):
                        return d_795_p2_

                    return lambda43_

                init23_ = lambda42_(p2)
                nw112_ = _dafny.Array(None, 4)
                for i0_23_ in range(nw112_.length(0)):
                    nw112_[i0_23_] = init23_(i0_23_)
                d_794_v107_ = nw112_
                index109_ = default__.safeIndex(74, (d_794_v107_).length(0))
                (d_794_v107_)[index109_] = p2
                index110_ = default__.safeIndex(74, (d_794_v107_).length(0))
                rhs122_ = d_755_cf12_
                rhs123_ = (d_755_cf12_) != (d_658_v0_)
                rhs124_ = ((p2 if True else p2)) + (p2)
                lhs114_ = globalState
                lhs115_ = d_794_v107_
                lhs116_ = default__.safeIndex(74, (d_794_v107_).length(0))
                d_658_v0_ = rhs122_
                lhs114_.f24 = rhs123_
                lhs115_[lhs116_] = rhs124_
                d_797_v108_: D2
                d_797_v108_ = D2_DC8()
                d_798_v109_: _dafny.Set
                d_798_v109_ = _dafny.Set({d_797_v108_})
                d_798_v109_ = d_798_v109_
                (globalState).f15 = (self).f26
                d_799_v110_: _dafny.Array
                def lambda44_(d_800_p0_):
                    def lambda45_(d_801_i10_):
                        return d_800_p0_

                    return lambda45_

                init24_ = lambda44_(p0)
                nw113_ = _dafny.Array(None, 17)
                for i0_24_ in range(nw113_.length(0)):
                    nw113_[i0_24_] = init24_(i0_24_)
                d_799_v110_ = nw113_
                d_802_v111_: _dafny.Map
                d_802_v111_ = _dafny.Map({(d_794_v107_)[default__.safeIndex(74, (d_794_v107_).length(0))]: d_799_v110_})
                d_803_v112_: D9
                d_803_v112_ = D9_DC21(d_802_v111_)
                d_804_v113_: _dafny.Seq
                d_804_v113_ = _dafny.SeqWithoutIsStrInference([d_803_v112_])
                (globalState).f24 = (d_803_v112_) in (d_804_v113_)
            d_805_v115_: _dafny.Array
            def lambda46_(d_806_cf12_, d_807_v0_):
                def lambda47_(d_808_i11_):
                    def iife72_():
                        coll52_ = _dafny.Set()
                        compr_52_: int
                        for compr_52_ in (_dafny.Set({d_807_v0_})).Elements:
                            d_809_v114_: int = compr_52_
                            if (d_809_v114_) in (_dafny.Set({d_807_v0_})):
                                coll52_ = coll52_.union(_dafny.Set([(d_809_v114_) + ((0) - (-311))]))
                        return _dafny.Set(coll52_)
                    return (iife72_()
                    ).intersection(_dafny.Set({len(_dafny.Map({d_807_v0_: _dafny.Map({d_807_v0_: d_806_cf12_})}))}))

                return lambda47_

            init25_ = lambda46_(d_755_cf12_, d_658_v0_)
            nw114_ = _dafny.Array(None, 18)
            for i0_25_ in range(nw114_.length(0)):
                nw114_[i0_25_] = init25_(i0_25_)
            d_805_v115_ = nw114_
            d_810_v116_: D1
            d_810_v116_ = D1_DC3((D1_DC3(_dafny.SeqWithoutIsStrInference([p0]))).cf6)
            d_811_v117_: _dafny.MultiSet
            d_811_v117_ = _dafny.MultiSet([len(default__.fm18(d_709_v40_, globalState))])
            d_812_v118_: _dafny.Seq
            d_812_v118_ = _dafny.SeqWithoutIsStrInference([336])
            pat_let_tv4_ = p0
            pat_let_tv5_ = d_811_v117_
            def iife73_(_pat_let10_0):
                def iife74_(d_813_dt__update__tmp_h1_):
                    def iife75_(_pat_let11_0):
                        def iife76_(d_814_dt__update_hcf15_h0_):
                            def iife77_(_pat_let12_0):
                                def iife78_(d_815_dt__update_hcf16_h0_):
                                    return D2_DC7((d_813_dt__update__tmp_h1_).cf14, d_814_dt__update_hcf15_h0_, d_815_dt__update_hcf16_h0_, (d_813_dt__update__tmp_h1_).cf17)
                                return iife78_(_pat_let12_0)
                            return iife77_((pat_let_tv5_).cardinality)
                        return iife76_(_pat_let11_0)
                    return iife75_(pat_let_tv4_)
                return iife74_(_pat_let10_0)
            rhs125_ = ((self).f27) and ((self).f26)
            rhs126_ = d_805_v115_
            rhs127_ = (0) - (default__.safeDivisionInt((iife73_(d_713_v44_)).cf17, (d_812_v118_)[default__.safeIndex(-276, len(d_812_v118_))]))
            rhs128_ = d_810_v116_
            lhs117_ = globalState
            lhs118_ = globalState
            lhs117_.f15 = rhs125_
            d_805_v115_ = rhs126_
            lhs118_.f12 = rhs127_
            d_810_v116_ = rhs128_
            d_816_v119_: _dafny.Map
            d_816_v119_ = _dafny.Map({(self).f26: (self).f26})
            d_816_v119_ = (d_816_v119_).set(p0, p0)
        elif True:
            d_817___mcc_h6_ = source14_.cf6
            d_818_cf6_ = d_817___mcc_h6_
            (globalState).f19 = d_708_v39_
            d_819_v120_: _dafny.Map
            d_819_v120_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")): d_658_v0_})
            (globalState).f12 = len(d_819_v120_)
            d_820_v121_: str
            d_821_v122_: _dafny.Seq
            d_822_v123_: bool
            d_823_v124_: _dafny.Map
            out25_: str
            out26_: _dafny.Seq
            out27_: bool
            out28_: _dafny.Map
            out25_, out26_, out27_, out28_ = default__.m0(d_708_v39_, globalState)
            d_820_v121_ = out25_
            d_821_v122_ = out26_
            d_822_v123_ = out27_
            d_823_v124_ = out28_
            if ((False if p0 else (self).fm2((self).f26, d_658_v0_, globalState)) if not((self).f26) else p0):
                d_824_v125_: _dafny.Map
                d_824_v125_ = _dafny.Map({p0: (self).fm1(globalState)})
                (globalState).f15 = (954) >= ((d_658_v0_) * (((d_824_v125_)[d_822_v123_] if (d_822_v123_) in (d_824_v125_) else 869)))
                d_825_v126_: _dafny.Map
                d_825_v126_ = _dafny.Map({d_658_v0_: d_658_v0_})
                d_825_v126_ = (d_825_v126_).set(d_658_v0_, len(d_821_v122_))
                d_822_v123_ = d_822_v123_
                d_826_v127_: _dafny.Array
                def lambda48_(d_827_i12_):
                    return (self).f26

                init26_ = lambda48_
                nw115_ = _dafny.Array(None, 6)
                for i0_26_ in range(nw115_.length(0)):
                    nw115_[i0_26_] = init26_(i0_26_)
                d_826_v127_ = nw115_
                d_828_v128_: _dafny.Map
                d_828_v128_ = _dafny.Map({243: d_826_v127_})
                d_829_v129_: _dafny.Array
                nw116_ = _dafny.Array(None, 23)
                nw116_[int(0)] = d_826_v127_
                nw116_[int(1)] = ((d_828_v128_)[d_658_v0_] if (d_658_v0_) in (d_828_v128_) else d_826_v127_)
                nw116_[int(2)] = d_826_v127_
                nw116_[int(3)] = d_826_v127_
                nw116_[int(4)] = d_826_v127_
                nw116_[int(5)] = d_826_v127_
                nw116_[int(6)] = d_826_v127_
                nw116_[int(7)] = d_826_v127_
                nw116_[int(8)] = d_826_v127_
                nw116_[int(9)] = d_826_v127_
                nw116_[int(10)] = d_826_v127_
                nw116_[int(11)] = d_826_v127_
                nw116_[int(12)] = d_826_v127_
                nw116_[int(13)] = d_826_v127_
                nw116_[int(14)] = d_826_v127_
                nw116_[int(15)] = d_826_v127_
                nw116_[int(16)] = d_826_v127_
                nw116_[int(17)] = d_826_v127_
                nw116_[int(18)] = d_826_v127_
                nw116_[int(19)] = d_826_v127_
                nw116_[int(20)] = d_826_v127_
                nw116_[int(21)] = d_826_v127_
                nw116_[int(22)] = d_826_v127_
                d_829_v129_ = nw116_
                (globalState).f5 = d_829_v129_
                index111_ = default__.safeIndex(492, (d_826_v127_).length(0))
                (d_826_v127_)[index111_] = p0
                index112_ = default__.safeIndex(492, (d_826_v127_).length(0))
                rhs129_ = d_820_v121_
                rhs130_ = not(p0)
                rhs131_ = (d_658_v0_) * ((0) - (d_658_v0_))
                rhs132_ = (d_821_v122_) + ((p2) + (d_821_v122_))
                rhs133_ = (self).f26
                lhs119_ = globalState
                lhs120_ = d_826_v127_
                lhs121_ = default__.safeIndex(492, (d_826_v127_).length(0))
                lhs122_ = globalState
                lhs123_ = globalState
                lhs119_.f19 = rhs129_
                lhs120_[lhs121_] = rhs130_
                lhs122_.f12 = rhs131_
                d_821_v122_ = rhs132_
                lhs123_.f15 = rhs133_
            elif True:
                d_830_v130_: _dafny.Set
                d_830_v130_ = _dafny.Set({d_658_v0_, d_658_v0_})
                d_831_v131_: _dafny.Map
                d_831_v131_ = _dafny.Map({d_658_v0_: len(d_830_v130_)})
                d_832_v132_: _dafny.Array
                nw117_ = _dafny.Array(None, 13)
                nw117_[int(0)] = d_658_v0_
                nw117_[int(1)] = d_658_v0_
                nw117_[int(2)] = d_658_v0_
                nw117_[int(3)] = d_658_v0_
                nw117_[int(4)] = d_658_v0_
                nw117_[int(5)] = 250
                nw117_[int(6)] = d_658_v0_
                nw117_[int(7)] = len(d_831_v131_)
                nw117_[int(8)] = 492
                nw117_[int(9)] = (0) - (d_658_v0_)
                nw117_[int(10)] = d_658_v0_
                nw117_[int(11)] = d_658_v0_
                nw117_[int(12)] = d_658_v0_
                d_832_v132_ = nw117_
                index113_ = default__.safeIndex(305, (d_832_v132_).length(0))
                (d_832_v132_)[index113_] = d_658_v0_
                index114_ = default__.safeIndex(305, (d_832_v132_).length(0))
                (d_832_v132_)[index114_] = (0) - ((d_658_v0_) + (d_658_v0_))
                (globalState).f24 = False
                d_833_v133_: D7
                d_833_v133_ = D7_DC19(d_658_v0_, default__.fm21((d_832_v132_)[default__.safeIndex(305, (d_832_v132_).length(0))], globalState), True, _dafny.Map({(self).f27: d_658_v0_}))
                (globalState).f22 = not(((len(d_716_v47_)) < (len(d_716_v47_)) if (self).f26 else (d_833_v133_).cf38))
                (globalState).f2 = (False) or ((self).f27)
                d_834_v134_: _dafny.Map
                d_834_v134_ = _dafny.Map({d_716_v47_: (self).f26})
                d_835_v135_: _dafny.Map
                d_835_v135_ = _dafny.Map({(self).f27: len(d_709_v40_)})
                d_834_v134_ = (default__.fm27(len(d_835_v135_), d_658_v0_, globalState)) | ((d_834_v134_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bd")), d_822_v123_))
        d_836_v136_: _dafny.Map
        d_836_v136_ = _dafny.Map({d_658_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ga"))})
        d_837_v137_: D0
        d_837_v137_ = D0_DC1((self).f26, (self).f26, len(d_836_v136_), p0)
        d_838_i13_: int
        d_838_i13_ = 0
        with _dafny.label("5"):
            while (default__.fm28(globalState)) == (_dafny.Set({d_837_v137_})):
                with _dafny.c_label("5"):
                    if (d_838_i13_) >= (100):
                        raise _dafny.Break("5")
                    d_838_i13_ = (d_838_i13_) + (1)
                    (globalState).f12 = d_658_v0_
                    (globalState).f12 = default__.safeDivisionInt(default__.safeModuloInt(d_658_v0_, d_658_v0_), len((p2 if (self).f27 else p2)))
                    d_839_v138_: _dafny.Array
                    nw118_ = _dafny.Array(int(0), 18)
                    d_839_v138_ = nw118_
                    d_840_v139_: _dafny.Seq
                    d_840_v139_ = _dafny.SeqWithoutIsStrInference([d_839_v138_])
                    rhs134_ = default__.safeDivisionInt(d_658_v0_, -344)
                    rhs135_ = ((self).f27) and ((d_840_v139_) == (d_840_v139_))
                    lhs124_ = globalState
                    lhs125_ = globalState
                    lhs124_.f12 = rhs134_
                    lhs125_.f15 = rhs135_
                    d_841_v140_: D5
                    d_841_v140_ = D5_DC14(803, d_658_v0_, (self).f27)
                    d_842_v141_: _dafny.Seq
                    d_842_v141_ = _dafny.SeqWithoutIsStrInference([(d_841_v140_).cf26, (self).fm1(globalState)])
                    d_658_v0_ = (d_842_v141_)[default__.safeIndex(291, len(d_842_v141_))]
                    pass
            pass
        d_843_v142_: _dafny.Array
        nw119_ = _dafny.Array(False, 14)
        d_843_v142_ = nw119_
        index115_ = default__.safeIndex(866, (d_843_v142_).length(0))
        (d_843_v142_)[index115_] = (self).f27
        index116_ = default__.safeIndex(866, (d_843_v142_).length(0))
        (d_843_v142_)[index116_] = (d_658_v0_) >= (d_658_v0_)
        index117_ = default__.safeIndex(866, (d_843_v142_).length(0))
        rhs136_ = d_716_v47_
        rhs137_ = (d_658_v0_) - (default__.safeDivisionInt(-402, len(d_709_v40_)))
        rhs138_ = p0
        rhs139_ = not((self).f26)
        lhs126_ = globalState
        lhs127_ = globalState
        lhs128_ = d_843_v142_
        lhs129_ = default__.safeIndex(866, (d_843_v142_).length(0))
        lhs130_ = globalState
        lhs126_.f6 = rhs136_
        lhs127_.f16 = rhs137_
        lhs128_[lhs129_] = rhs138_
        lhs130_.f24 = rhs139_

    def m9(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        r3: bool = False
        d_844_v0_: _dafny.Array
        nw120_ = _dafny.Array(_dafny.Map({}), 28)
        d_844_v0_ = nw120_
        d_845_v1_: int
        d_845_v1_ = -964
        d_846_v2_: _dafny.Map
        d_846_v2_ = _dafny.Map({d_845_v1_: d_845_v1_})
        d_847_v3_: _dafny.Map
        d_847_v3_ = _dafny.Map({d_845_v1_: d_846_v2_})
        index118_ = default__.safeIndex(584, (d_844_v0_).length(0))
        (d_844_v0_)[index118_] = d_847_v3_
        d_848_v4_: _dafny.Seq
        d_848_v4_ = _dafny.SeqWithoutIsStrInference([d_845_v1_, d_845_v1_])
        d_849_v5_: _dafny.MultiSet
        d_849_v5_ = _dafny.MultiSet([(self).fm2(p0, len(d_848_v4_), globalState), (self).f26])
        d_850_v6_: _dafny.Seq
        d_850_v6_ = _dafny.SeqWithoutIsStrInference([(d_849_v5_).issubset(d_849_v5_)])
        d_851_v7_: _dafny.Array
        nw121_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_851_v7_ = nw121_
        index119_ = default__.safeIndex(644, (d_851_v7_).length(0))
        (d_851_v7_)[index119_] = p1
        index120_ = default__.safeIndex(584, (d_844_v0_).length(0))
        index121_ = default__.safeIndex(644, (d_851_v7_).length(0))
        rhs140_ = d_847_v3_
        rhs141_ = d_850_v6_
        rhs142_ = p1
        lhs131_ = d_844_v0_
        lhs132_ = default__.safeIndex(584, (d_844_v0_).length(0))
        lhs133_ = d_851_v7_
        lhs134_ = default__.safeIndex(644, (d_851_v7_).length(0))
        lhs131_[lhs132_] = rhs140_
        d_850_v6_ = rhs141_
        lhs133_[lhs134_] = rhs142_
        d_845_v1_ = d_845_v1_
        (globalState).f2 = p0
        r0 = (self).f26
        d_852_v8_: _dafny.Array
        nw122_ = _dafny.Array(_dafny.Array(None, 0), 28)
        d_852_v8_ = nw122_
        d_853_v9_: _dafny.Array
        nw123_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
        d_853_v9_ = nw123_
        d_854_v10_: _dafny.Map
        d_854_v10_ = _dafny.Map({(0) - (d_845_v1_): d_853_v9_})
        index122_ = default__.safeIndex(34, (d_852_v8_).length(0))
        (d_852_v8_)[index122_] = ((d_854_v10_)[d_845_v1_] if (d_845_v1_) in (d_854_v10_) else d_853_v9_)
        d_855_v11_: _dafny.Array
        nw124_ = _dafny.Array(int(0), 23)
        d_855_v11_ = nw124_
        index123_ = default__.safeIndex(427, (p1).length(0))
        (p1)[index123_] = (self).f27
        index124_ = default__.safeIndex(34, (d_852_v8_).length(0))
        index125_ = default__.safeIndex(427, (p1).length(0))
        index126_ = default__.safeIndex(644, (d_851_v7_).length(0))
        rhs143_ = d_853_v9_
        rhs144_ = d_855_v11_
        rhs145_ = p0
        rhs146_ = p1
        lhs135_ = d_852_v8_
        lhs136_ = default__.safeIndex(34, (d_852_v8_).length(0))
        lhs137_ = p1
        lhs138_ = default__.safeIndex(427, (p1).length(0))
        lhs139_ = d_851_v7_
        lhs140_ = default__.safeIndex(644, (d_851_v7_).length(0))
        lhs135_[lhs136_] = rhs143_
        d_855_v11_ = rhs144_
        lhs137_[lhs138_] = rhs145_
        lhs139_[lhs140_] = rhs146_
        index127_ = default__.safeIndex(427, (p1).length(0))
        (p1)[index127_] = not((self).f27)
        d_856_v12_: str
        d_856_v12_ = _dafny.CodePoint('e')
        d_857_v13_: _dafny.Map
        d_857_v13_ = _dafny.Map({d_845_v1_: d_856_v12_})
        d_858_v14_: _dafny.Set
        d_858_v14_ = _dafny.Set({d_857_v13_})
        def iife79_():
            coll53_ = _dafny.Set()
            compr_53_: _dafny.Map
            for compr_53_ in (d_858_v14_).Elements:
                d_859_v15_: _dafny.Map = compr_53_
                if (d_859_v15_) in (d_858_v14_):
                    coll53_ = coll53_.union(_dafny.Set([d_859_v15_]))
            return _dafny.Set(coll53_)
        r0 = not(((d_858_v14_) - (iife79_()
        )).issubset(d_858_v14_))
        r1 = not((_dafny.SeqWithoutIsStrInference([(0) - (d_845_v1_) for d_860_i0_ in range(default__.abs(530))])) != (d_848_v4_))
        d_861_v16_: D9
        d_861_v16_ = D9_DC22((self).f27)
        r2 = ((0) - (d_845_v1_) if (d_861_v16_).cf42 else len(default__.fm15(globalState)))
        d_862_v17_: D2
        d_862_v17_ = D2_DC8()
        pat_let_tv6_ = p1
        pat_let_tv7_ = p1
        pat_let_tv8_ = p1
        pat_let_tv9_ = p1
        pat_let_tv10_ = d_850_v6_
        pat_let_tv11_ = d_850_v6_
        pat_let_tv12_ = p2
        pat_let_tv13_ = p2
        pat_let_tv14_ = d_850_v6_
        pat_let_tv15_ = d_850_v6_
        pat_let_tv16_ = d_845_v1_
        pat_let_tv17_ = d_845_v1_
        pat_let_tv18_ = d_845_v1_
        def lambda49_(source15_):
            if source15_.is_DC7:
                d_863___mcc_h0_ = source15_.cf14
                d_864___mcc_h1_ = source15_.cf15
                d_865___mcc_h2_ = source15_.cf16
                d_866___mcc_h3_ = source15_.cf17
                d_867_cf17_ = d_866___mcc_h3_
                d_868_cf16_ = d_865___mcc_h2_
                d_869_cf15_ = d_864___mcc_h1_
                d_870_cf14_ = d_863___mcc_h0_
                return (pat_let_tv7_)[default__.safeIndex(427, (pat_let_tv6_).length(0))]
            elif source15_.is_DC8:
                return (pat_let_tv9_)[default__.safeIndex(427, (pat_let_tv8_).length(0))]
            elif source15_.is_DC9:
                d_871___mcc_h4_ = source15_.cf18
                d_872___mcc_h5_ = source15_.cf19
                d_873___mcc_h6_ = source15_.cf20
                d_874_cf20_ = d_873___mcc_h6_
                d_875_cf19_ = d_872___mcc_h5_
                d_876_cf18_ = d_871___mcc_h4_
                if ((pat_let_tv10_)[default__.safeIndex(d_876_cf18_, len(pat_let_tv11_))]) in (pat_let_tv12_):
                    return (pat_let_tv13_)[(pat_let_tv14_)[default__.safeIndex(d_876_cf18_, len(pat_let_tv15_))]]
                elif True:
                    return (D5_DC14(752, 463, d_874_cf20_)).cf27
            elif True:
                d_877___mcc_h7_ = source15_.cf13
                d_878_cf13_ = d_877___mcc_h7_
                def iife80_():
                    coll54_ = _dafny.Set()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(-111, -734):
                        d_879_v18_: int = compr_54_
                        if ((-111) <= (d_879_v18_)) and ((d_879_v18_) < (-734)):
                            coll54_ = coll54_.union(_dafny.Set([(d_879_v18_) + (pat_let_tv18_)]))
                    return _dafny.Set(coll54_)
                return (_dafny.Set({pat_let_tv16_, pat_let_tv17_})) == (iife80_()
                )

        r3 = lambda49_(d_862_v17_)
        return r0, r1, r2, r3

    def m10(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: str = _dafny.CodePoint('D')
        d_880_v0_: _dafny.Map
        d_880_v0_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): not(p0)})
        d_881_v1_: _dafny.Seq
        d_881_v1_ = _dafny.SeqWithoutIsStrInference([p0, (self).fm2((self).f27, 771, globalState), p0])
        hi10_ = len((d_881_v1_) + (d_881_v1_))
        for d_882_i0_ in range(len(d_880_v0_), hi10_):
            d_883_v2_: _dafny.Array
            def lambda50_(d_884_v1_):
                def lambda51_(d_885_i1_):
                    return _dafny.SeqWithoutIsStrInference([d_884_v1_, d_884_v1_, d_884_v1_])

                return lambda51_

            init27_ = lambda50_(d_881_v1_)
            nw125_ = _dafny.Array(None, 19)
            for i0_27_ in range(nw125_.length(0)):
                nw125_[i0_27_] = init27_(i0_27_)
            d_883_v2_ = nw125_
            index128_ = default__.safeIndex(649, (d_883_v2_).length(0))
            (d_883_v2_)[index128_] = (_dafny.SeqWithoutIsStrInference([d_881_v1_])) + (_dafny.SeqWithoutIsStrInference([d_881_v1_]))
            d_886_v3_: _dafny.Seq
            d_886_v3_ = _dafny.SeqWithoutIsStrInference([d_881_v1_, d_881_v1_])
            index129_ = default__.safeIndex(649, (d_883_v2_).length(0))
            (d_883_v2_)[index129_] = (d_886_v3_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0]), d_881_v1_]))
            d_887_v4_: _dafny.Seq
            d_887_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usmgaxt"))
            d_888_v5_: _dafny.MultiSet
            d_888_v5_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p1 for d_889_i2_ in range(default__.abs(514))])), 123, p1, d_882_i0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_890_i3_ in range(default__.abs(886))]))])
            d_891_v6_: _dafny.Set
            d_891_v6_ = _dafny.Set({p1, (d_888_v5_).cardinality})
            d_892_v7_: _dafny.Map
            d_892_v7_ = _dafny.Map({d_891_v6_: (self).f27})
            d_893_v8_: _dafny.Seq
            d_893_v8_ = _dafny.SeqWithoutIsStrInference([-252, len(d_891_v6_)])
            d_894_v9_: _dafny.MultiSet
            d_894_v9_ = _dafny.MultiSet([(self).f26])
            d_895_v10_: _dafny.Array
            nw126_ = _dafny.Array(None, 20)
            nw126_[int(0)] = len(default__.fm29(False, d_882_i0_, globalState))
            nw126_[int(1)] = default__.safeDivisionInt(len(default__.fm22(len(d_887_v4_), (d_881_v1_)[default__.safeIndex(p1, len(d_881_v1_))], globalState)), d_882_i0_)
            nw126_[int(2)] = len(d_891_v6_)
            nw126_[int(3)] = len((d_891_v6_) - (d_891_v6_))
            nw126_[int(4)] = d_882_i0_
            nw126_[int(5)] = (p1) - (p1)
            nw126_[int(6)] = (0) - (len((d_892_v7_) | (d_892_v7_)))
            nw126_[int(7)] = default__.safeDivisionInt(p1, d_882_i0_)
            nw126_[int(8)] = (d_882_i0_) + (p1)
            nw126_[int(9)] = default__.safeDivisionInt(len(default__.fm30(d_882_i0_, p0, globalState)), len(d_881_v1_))
            nw126_[int(10)] = d_882_i0_
            nw126_[int(11)] = (self).fm1(globalState)
            nw126_[int(12)] = d_882_i0_
            nw126_[int(13)] = d_882_i0_
            nw126_[int(14)] = (0) - ((p1 if (self).f27 else d_882_i0_))
            nw126_[int(15)] = p1
            nw126_[int(16)] = default__.safeModuloInt(d_882_i0_, (d_893_v8_)[default__.safeIndex(d_882_i0_, len(d_893_v8_))])
            nw126_[int(17)] = p1
            nw126_[int(18)] = len(d_887_v4_)
            nw126_[int(19)] = (d_894_v9_).cardinality
            d_895_v10_ = nw126_
            index130_ = default__.safeIndex(625, (d_895_v10_).length(0))
            (d_895_v10_)[index130_] = p1
            index131_ = default__.safeIndex(625, (d_895_v10_).length(0))
            (d_895_v10_)[index131_] = (p1) - (d_882_i0_)
            d_896_v11_: C1
            nw127_ = C1()
            nw127_.ctor__((self).fm1(globalState), len(d_881_v1_), (self).f26, p0)
            d_896_v11_ = nw127_
            (d_896_v11_).f37 = ((d_894_v9_) - (d_894_v9_)).cardinality
        d_897_v12_: _dafny.MultiSet
        d_897_v12_ = _dafny.MultiSet([(self).f26, (self).f27, (self).fm2((self).f27, p1, globalState)])
        (globalState).f22 = (d_881_v1_)[default__.safeIndex((d_897_v12_).cardinality, len(d_881_v1_))]
        d_898_v13_: _dafny.Array
        def lambda52_(d_899_p1_, d_900_p0_):
            def lambda53_(d_901_i4_):
                return _dafny.SeqWithoutIsStrInference([(0) - ((D5_DC14(d_899_p1_, len(_dafny.SeqWithoutIsStrInference([-342, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_902_i5_ in range(default__.abs(671))]))])), d_900_p0_)).cf26)])

            return lambda53_

        init28_ = lambda52_(p1, p0)
        nw128_ = _dafny.Array(None, 13)
        for i0_28_ in range(nw128_.length(0)):
            nw128_[i0_28_] = init28_(i0_28_)
        d_898_v13_ = nw128_
        d_898_v13_ = d_898_v13_
        d_903_v14_: _dafny.Seq
        d_903_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgkdg"))
        d_904_v16_: _dafny.Map
        d_904_v16_ = _dafny.Map({602: _dafny.Set({(self).f26})})
        def iife81_():
            coll55_ = _dafny.Map()
            compr_55_: int
            for compr_55_ in _dafny.IntegerRange(-648, 519):
                d_905_v15_: int = compr_55_
                if ((-648) <= (d_905_v15_)) and ((d_905_v15_) < (519)):
                    coll55_[(d_905_v15_) * (p1)] = _dafny.Set({(self).f26})
            return _dafny.Map(coll55_)
        if (len(d_903_v14_)) not in ((iife81_()
         if p0 else d_904_v16_)):
            d_906_v17_: _dafny.Array
            nw129_ = _dafny.Array(False, 11)
            d_906_v17_ = nw129_
            index132_ = default__.safeIndex(104, (d_906_v17_).length(0))
            (d_906_v17_)[index132_] = (self).f26
            index133_ = default__.safeIndex(104, (d_906_v17_).length(0))
            (d_906_v17_)[index133_] = not(True)
            d_907_v18_: _dafny.Seq
            d_907_v18_ = _dafny.SeqWithoutIsStrInference([d_881_v1_])
            d_907_v18_ = d_907_v18_
            d_908_v19_: C0
            nw130_ = C0()
            nw130_.ctor__((self).f27)
            d_908_v19_ = nw130_
            d_909_v20_: _dafny.Map
            d_909_v20_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ettqukfnq")): d_908_v19_})
            d_910_v21_: _dafny.Array
            nw131_ = _dafny.Array(None, 16)
            nw131_[int(0)] = d_908_v19_
            nw131_[int(1)] = d_908_v19_
            nw131_[int(2)] = d_908_v19_
            nw131_[int(3)] = d_908_v19_
            nw131_[int(4)] = d_908_v19_
            nw131_[int(5)] = d_908_v19_
            nw131_[int(6)] = d_908_v19_
            nw131_[int(7)] = d_908_v19_
            nw131_[int(8)] = d_908_v19_
            nw131_[int(9)] = (d_908_v19_ if (d_906_v17_)[default__.safeIndex(104, (d_906_v17_).length(0))] else d_908_v19_)
            nw131_[int(10)] = d_908_v19_
            nw131_[int(11)] = d_908_v19_
            nw131_[int(12)] = ((d_909_v20_)[d_903_v14_] if (d_903_v14_) in (d_909_v20_) else d_908_v19_)
            nw131_[int(13)] = d_908_v19_
            nw131_[int(14)] = d_908_v19_
            nw131_[int(15)] = d_908_v19_
            d_910_v21_ = nw131_
            index134_ = default__.safeIndex(78, (d_910_v21_).length(0))
            (d_910_v21_)[index134_] = d_908_v19_
            index135_ = default__.safeIndex(78, (d_910_v21_).length(0))
            (d_910_v21_)[index135_] = d_908_v19_
            d_911_v22_: _dafny.Array
            nw132_ = _dafny.Array(None, 4)
            d_911_v22_ = nw132_
            d_912_v23_: _dafny.Seq
            d_912_v23_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_913_v24_: _dafny.MultiSet
            d_913_v24_ = _dafny.MultiSet([d_912_v23_, d_912_v23_, d_912_v23_])
            d_914_v25_: C3
            nw133_ = C3()
            nw133_.ctor__(d_913_v24_, (d_906_v17_)[default__.safeIndex(104, (d_906_v17_).length(0))], (self).f26)
            d_914_v25_ = nw133_
            index136_ = default__.safeIndex(905, (d_911_v22_).length(0))
            (d_911_v22_)[index136_] = d_914_v25_
            index137_ = default__.safeIndex(905, (d_911_v22_).length(0))
            (d_911_v22_)[index137_] = d_914_v25_
            d_915_v26_: _dafny.Map
            d_915_v26_ = _dafny.Map({p0: p0})
            d_915_v26_ = (d_915_v26_).set((_dafny.SeqWithoutIsStrInference([d_881_v1_])) != (d_907_v18_), (self).f26)
        elif True:
            d_916_v27_: _dafny.Map
            d_916_v27_ = _dafny.Map({not((self).f27): p1})
            d_917_v28_: _dafny.Map
            d_917_v28_ = _dafny.Map({len(d_916_v27_): p1})
            d_918_v29_: _dafny.Seq
            d_918_v29_ = _dafny.SeqWithoutIsStrInference([d_898_v13_, d_898_v13_])
            rhs147_ = p1
            rhs148_ = len((d_881_v1_) + ((d_881_v1_) + (d_881_v1_)))
            rhs149_ = (self).f26
            rhs150_ = (p1 if False else ((d_917_v28_)[p1] if (p1) in (d_917_v28_) else p1))
            rhs151_ = (d_918_v29_)[default__.safeIndex((self).fm1(globalState), len(d_918_v29_))]
            lhs141_ = globalState
            r2 = rhs147_
            r2 = rhs148_
            lhs141_.f22 = rhs149_
            r2 = rhs150_
            d_898_v13_ = rhs151_
            d_919_v30_: _dafny.Array
            nw134_ = _dafny.Array(False, 5)
            d_919_v30_ = nw134_
            d_920_v31_: str
            d_920_v31_ = _dafny.CodePoint('s')
            d_921_v32_: _dafny.Map
            d_921_v32_ = _dafny.Map({d_919_v30_: d_920_v31_})
            d_922_v33_: str
            d_923_v34_: _dafny.Seq
            d_924_v35_: bool
            d_925_v36_: _dafny.Map
            out29_: str
            out30_: _dafny.Seq
            out31_: bool
            out32_: _dafny.Map
            out29_, out30_, out31_, out32_ = default__.m0(((d_921_v32_)[d_919_v30_] if (d_919_v30_) in (d_921_v32_) else d_920_v31_), globalState)
            d_922_v33_ = out29_
            d_923_v34_ = out30_
            d_924_v35_ = out31_
            d_925_v36_ = out32_
            d_926_v37_: _dafny.Map
            d_926_v37_ = _dafny.Map({p1: d_923_v34_})
            d_927_v38_: C1
            nw135_ = C1()
            nw135_.ctor__(p1, len(((d_926_v37_)[p1] if (p1) in (d_926_v37_) else d_923_v34_)), (self).f26, (self).f26)
            d_927_v38_ = nw135_
            d_928_v39_: str
            d_929_v40_: _dafny.Seq
            d_930_v41_: bool
            d_931_v42_: _dafny.Map
            out33_: str
            out34_: _dafny.Seq
            out35_: bool
            out36_: _dafny.Map
            out33_, out34_, out35_, out36_ = default__.m0(_dafny.CodePoint('l'), globalState)
            d_928_v39_ = out33_
            d_929_v40_ = out34_
            d_930_v41_ = out35_
            d_931_v42_ = out36_
            if (D5_DC14(d_927_v38_.f37, (d_897_v12_).cardinality, d_930_v41_)).cf27:
                d_932_v43_: _dafny.Seq
                d_932_v43_ = _dafny.SeqWithoutIsStrInference([(0) - (d_927_v38_.f37), d_927_v38_.f37])
                d_933_v44_: _dafny.Seq
                d_933_v44_ = _dafny.SeqWithoutIsStrInference([d_932_v43_, d_932_v43_])
                d_934_v45_: D0
                d_934_v45_ = D0_DC0((self).f26)
                d_935_v46_: _dafny.Set
                d_935_v46_ = _dafny.Set({(d_927_v38_).fm4(d_934_v45_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uc"))), len(d_929_v40_), globalState)})
                d_936_v47_: T2
                nw136_ = C2()
                nw136_.ctor__(d_933_v44_, (p0 if d_924_v35_ else d_930_v41_), (d_935_v46_).ispropersubset(d_935_v46_))
                d_936_v47_ = nw136_
                nw137_ = C2()
                nw137_.ctor__(_dafny.SeqWithoutIsStrInference([d_932_v43_ for d_937_i6_ in range(default__.abs(-120))]), (d_903_v14_) <= (d_903_v14_), (d_936_v47_).f26)
                d_936_v47_ = nw137_
                d_919_v30_ = d_919_v30_
                (globalState).f2 = (d_936_v47_).f26
                d_938_v48_: _dafny.Set
                d_938_v48_ = _dafny.Set({(self).fm2(not((d_936_v47_).f27), d_927_v38_.f37, globalState)})
                d_939_v49_: _dafny.Map
                d_939_v49_ = _dafny.Map({d_922_v33_: d_927_v38_.f37})
                d_940_v50_: _dafny.Map
                d_940_v50_ = _dafny.Map({(d_927_v38_).f36: d_927_v38_})
                d_941_v51_: _dafny.Map
                d_941_v51_ = _dafny.Map({(d_927_v38_).f36: d_903_v14_})
                d_942_v52_: _dafny.Array
                nw138_ = _dafny.Array(None, 17)
                nw138_[int(0)] = len(d_938_v48_)
                nw138_[int(1)] = len(d_939_v49_)
                nw138_[int(2)] = (d_927_v38_).f36
                nw138_[int(3)] = len(d_940_v50_)
                nw138_[int(4)] = len(((d_941_v51_)[(d_927_v38_).f36] if ((d_927_v38_).f36) in (d_941_v51_) else d_903_v14_))
                nw138_[int(5)] = d_927_v38_.f37
                nw138_[int(6)] = (d_927_v38_).f36
                nw138_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umscmdh")))
                nw138_[int(8)] = (d_927_v38_).f36
                nw138_[int(9)] = (d_927_v38_).f36
                nw138_[int(10)] = d_927_v38_.f37
                nw138_[int(11)] = len(default__.fm31(p1, globalState))
                nw138_[int(12)] = (0) - ((d_927_v38_).f36)
                nw138_[int(13)] = (d_927_v38_).f36
                nw138_[int(14)] = d_927_v38_.f37
                nw138_[int(15)] = d_927_v38_.f37
                nw138_[int(16)] = (d_927_v38_).f36
                d_942_v52_ = nw138_
                d_943_v53_: _dafny.Map
                d_943_v53_ = _dafny.Map({d_942_v52_: d_928_v39_})
                d_943_v53_ = (d_943_v53_).set(d_942_v52_, d_922_v33_)
                r0 = default__.safeDivisionInt(((d_927_v38_).f36 if d_930_v41_ else len((d_936_v47_).fm3(d_903_v14_, d_881_v1_, (d_927_v38_).f36, globalState))), (d_927_v38_).f36)
            elif True:
                d_944_v54_: _dafny.Array
                nw139_ = _dafny.Array(None, 7)
                nw139_[int(0)] = default__.safeDivisionInt((d_927_v38_).f36, d_927_v38_.f37)
                nw139_[int(1)] = (d_927_v38_).f36
                nw139_[int(2)] = d_927_v38_.f37
                nw139_[int(3)] = p1
                nw139_[int(4)] = d_927_v38_.f37
                nw139_[int(5)] = (d_927_v38_).f36
                nw139_[int(6)] = (d_927_v38_.f37) + ((d_927_v38_).f36)
                d_944_v54_ = nw139_
                d_945_v55_: _dafny.Array
                nw140_ = _dafny.Array(None, 4)
                nw140_[int(0)] = d_944_v54_
                nw140_[int(1)] = d_944_v54_
                nw140_[int(2)] = d_944_v54_
                nw140_[int(3)] = d_944_v54_
                d_945_v55_ = nw140_
                rhs152_ = d_924_v35_
                rhs153_ = d_944_v54_
                rhs154_ = 89
                rhs155_ = d_945_v55_
                lhs142_ = globalState
                lhs143_ = d_927_v38_
                lhs142_.f2 = rhs152_
                d_944_v54_ = rhs153_
                lhs143_.f37 = rhs154_
                d_945_v55_ = rhs155_
                d_946_v56_: D0
                d_946_v56_ = D0_DC0((self).f26)
                (globalState).f12 = (d_927_v38_).fm4(d_946_v56_, (0) - (((d_927_v38_).f36) * (d_927_v38_.f37)), default__.safeDivisionInt(d_927_v38_.f37, (d_927_v38_).f36), globalState)
                (globalState).f22 = not(True)
                d_947_v57_: _dafny.Array
                nw141_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                d_947_v57_ = nw141_
                d_948_v58_: _dafny.Array
                nw142_ = _dafny.Array(None, 5)
                nw142_[int(0)] = d_927_v38_
                nw142_[int(1)] = d_927_v38_
                nw142_[int(2)] = d_927_v38_
                nw142_[int(3)] = d_927_v38_
                nw142_[int(4)] = d_927_v38_
                d_948_v58_ = nw142_
                index138_ = default__.safeIndex(832, (d_948_v58_).length(0))
                (d_948_v58_)[index138_] = d_927_v38_
                index139_ = default__.safeIndex(832, (d_948_v58_).length(0))
                rhs156_ = d_947_v57_
                rhs157_ = (self).f27
                rhs158_ = (826) * ((d_927_v38_).f36)
                rhs159_ = (d_927_v38_).f36
                rhs160_ = d_927_v38_
                lhs144_ = globalState
                lhs145_ = globalState
                lhs146_ = globalState
                lhs147_ = d_948_v58_
                lhs148_ = default__.safeIndex(832, (d_948_v58_).length(0))
                d_947_v57_ = rhs156_
                lhs144_.f22 = rhs157_
                lhs145_.f16 = rhs158_
                lhs146_.f16 = rhs159_
                lhs147_[lhs148_] = rhs160_
                (globalState).f22 = (self).f26
        d_949_v59_: _dafny.Array
        def lambda54_(d_950_i7_):
            return (self).f27

        init29_ = lambda54_
        nw143_ = _dafny.Array(None, 7)
        for i0_29_ in range(nw143_.length(0)):
            nw143_[i0_29_] = init29_(i0_29_)
        d_949_v59_ = nw143_
        d_951_v60_: _dafny.Map
        d_951_v60_ = _dafny.Map({(self).f26: (self).f27})
        d_952_v61_: bool
        d_953_v62_: bool
        d_954_v63_: int
        d_955_v64_: bool
        out37_: bool
        out38_: bool
        out39_: int
        out40_: bool
        out37_, out38_, out39_, out40_ = (self).m9((self).f27, d_949_v59_, (d_951_v60_ if (self).f26 else d_951_v60_), globalState)
        d_952_v61_ = out37_
        d_953_v62_ = out38_
        d_954_v63_ = out39_
        d_955_v64_ = out40_
        d_956_v65_: _dafny.Seq
        d_956_v65_ = _dafny.SeqWithoutIsStrInference([d_881_v1_])
        d_956_v65_ = ((_dafny.SeqWithoutIsStrInference([d_881_v1_])) + (d_956_v65_)) + (d_956_v65_)
        r0 = p1
        r1 = d_954_v63_
        r2 = (0) - (p1)
        r3 = default__.fm24(d_903_v14_, (self).f26, p1, globalState)
        return r0, r1, r2, r3

    @property
    def f33(self):
        return self._f33

class C5(T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f26, f27):
        (self)._f26 = f26
        (self)._f27 = f27

    def fm1(self, globalState):
        return -356

    def fm2(self, p0, p1, globalState):
        source16_ = (D0_DC2(D0_DC1((self).f27, (self).f27, 756, not(True))) if (self).f27 else D0_DC2(D0_DC0((self).f26)))
        if source16_.is_DC1:
            d_957___mcc_h0_ = source16_.cf1
            d_958___mcc_h1_ = source16_.cf2
            d_959___mcc_h2_ = source16_.cf3
            d_960___mcc_h3_ = source16_.cf4
            d_961_cf4_ = d_960___mcc_h3_
            d_962_cf3_ = d_959___mcc_h2_
            d_963_cf2_ = d_958___mcc_h1_
            d_964_cf1_ = d_957___mcc_h0_
            return (self).f27
        elif source16_.is_DC0:
            d_965___mcc_h4_ = source16_.cf0
            d_966_cf0_ = d_965___mcc_h4_
            return (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrscocei")))) == (468)
        elif True:
            d_967___mcc_h5_ = source16_.cf5
            d_968_cf5_ = d_967___mcc_h5_
            return (self).f26

    def fm5(self, p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('g')

    def m1(self, p0, p1, p2, globalState):
        d_969_v0_: C0
        nw144_ = C0()
        nw144_.ctor__((self).f26)
        d_969_v0_ = nw144_
        d_970_v2_: _dafny.Map
        d_970_v2_ = _dafny.Map({p2: p0})
        d_971_v3_: _dafny.Map
        def iife82_():
            coll56_ = _dafny.Map()
            compr_56_: _dafny.Seq
            for compr_56_ in (d_970_v2_).keys.Elements:
                d_972_v1_: _dafny.Seq = compr_56_
                if (d_972_v1_) in (d_970_v2_):
                    coll56_[d_972_v1_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmtrj")))
            return _dafny.Map(coll56_)
        d_971_v3_ = _dafny.Map({iife82_()
        : p0})
        d_973_v4_: int
        d_973_v4_ = -494
        d_974_v5_: _dafny.Map
        d_974_v5_ = _dafny.Map({p2: d_973_v4_})
        d_975_v7_: D2
        def iife83_():
            coll57_ = _dafny.Map()
            compr_57_: int
            for compr_57_ in _dafny.IntegerRange(828, 982):
                d_976_v6_: int = compr_57_
                if ((828) <= (d_976_v6_)) and ((d_976_v6_) < (982)):
                    coll57_[(d_976_v6_) - (d_973_v4_)] = d_973_v4_
            return _dafny.Map(coll57_)
        d_975_v7_ = D2_DC6(iife83_()
)
        d_971_v3_ = (d_971_v3_).set(d_974_v5_, not((len((d_975_v7_).cf13)) > (d_973_v4_)))
        (globalState).f6 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbjowp"))
        d_977_v8_: C0
        nw145_ = C0()
        nw145_.ctor__(not (p0) or ((self).f27))
        d_977_v8_ = nw145_
        d_978_v9_: D2
        d_978_v9_ = D2_DC9(530, len(p2), (self).f26)
        source17_ = d_978_v9_
        if source17_.is_DC7:
            d_979___mcc_h0_ = source17_.cf14
            d_980___mcc_h1_ = source17_.cf15
            d_981___mcc_h2_ = source17_.cf16
            d_982___mcc_h3_ = source17_.cf17
            d_983_cf17_ = d_982___mcc_h3_
            d_984_cf16_ = d_981___mcc_h2_
            d_985_cf15_ = d_980___mcc_h1_
            d_986_cf14_ = d_979___mcc_h0_
            d_987_v10_: _dafny.Array
            nw146_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_987_v10_ = nw146_
            d_988_v11_: _dafny.Seq
            d_988_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            index140_ = default__.safeIndex(397, (d_987_v10_).length(0))
            (d_987_v10_)[index140_] = d_988_v11_
            index141_ = default__.safeIndex(397, (d_987_v10_).length(0))
            (d_987_v10_)[index141_] = (d_988_v11_) + (d_988_v11_)
            (globalState).f22 = (self).fm2(d_986_cf14_, (0) - (d_973_v4_), globalState)
            d_989_v12_: _dafny.Array
            nw147_ = _dafny.Array(False, 20)
            d_989_v12_ = nw147_
            index142_ = default__.safeIndex(461, (d_989_v12_).length(0))
            (d_989_v12_)[index142_] = ((d_987_v10_)[default__.safeIndex(397, (d_987_v10_).length(0))]) <= ((d_987_v10_)[default__.safeIndex(397, (d_987_v10_).length(0))])
            index143_ = default__.safeIndex(461, (d_989_v12_).length(0))
            (d_989_v12_)[index143_] = False
            (globalState).f2 = (self).f27
        elif source17_.is_DC8:
            (globalState).f15 = (d_969_v0_).f32
            d_990_v13_: D0
            d_990_v13_ = D0_DC1((d_969_v0_).f32, (d_977_v8_).f32, d_973_v4_, (d_977_v8_).f32)
            d_991_v14_: _dafny.Map
            d_991_v14_ = _dafny.Map({d_973_v4_: (d_990_v13_).cf3})
            d_992_v15_: _dafny.MultiSet
            d_992_v15_ = _dafny.MultiSet([default__.fm8(40, d_973_v4_, (d_969_v0_).f32, d_973_v4_, globalState), d_991_v14_, d_991_v14_, d_991_v14_, d_991_v14_])
            d_993_v16_: str
            d_993_v16_ = _dafny.CodePoint('o')
            d_994_v17_: _dafny.Map
            d_994_v17_ = _dafny.Map({p0: d_993_v16_})
            d_995_v18_: _dafny.Map
            d_995_v18_ = _dafny.Map({len((d_994_v17_) | (_dafny.Map({p0: d_993_v16_}))): _dafny.MultiSet([d_991_v14_])})
            d_992_v15_ = ((d_995_v18_)[d_973_v4_] if (d_973_v4_) in (d_995_v18_) else (d_992_v15_) | (d_992_v15_))
            d_996_v19_: _dafny.Seq
            d_996_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sksvlyg"))
            (globalState).f6 = ((d_996_v19_).set(default__.safeIndex(d_973_v4_, len(d_996_v19_)), d_993_v16_)) + (d_996_v19_)
            d_997_v20_: _dafny.MultiSet
            d_997_v20_ = _dafny.MultiSet([(d_969_v0_).f32, (d_969_v0_).f32])
            d_997_v20_ = default__.fm9(not((d_977_v8_).f32), d_973_v4_, (d_969_v0_).f32, globalState)
        elif source17_.is_DC9:
            d_998___mcc_h4_ = source17_.cf18
            d_999___mcc_h5_ = source17_.cf19
            d_1000___mcc_h6_ = source17_.cf20
            d_1001_cf20_ = d_1000___mcc_h6_
            d_1002_cf19_ = d_999___mcc_h5_
            d_1003_cf18_ = d_998___mcc_h4_
            d_1004_v21_: _dafny.Seq
            d_1004_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfsobu"))
            d_1005_v22_: _dafny.MultiSet
            d_1005_v22_ = _dafny.MultiSet([True])
            d_1006_v23_: _dafny.Map
            d_1006_v23_ = _dafny.Map({(self).f26: d_1002_cf19_})
            d_1007_v24_: _dafny.MultiSet
            d_1007_v24_ = _dafny.MultiSet([(d_1005_v22_).cardinality, len(d_1006_v23_)])
            d_1008_v25_: str
            d_1008_v25_ = _dafny.CodePoint('l')
            d_1009_v26_: _dafny.Map
            d_1009_v26_ = _dafny.Map({(d_1004_v21_).set(default__.safeIndex(((d_1007_v24_)[(self).fm1(globalState)] if ((self).fm1(globalState)) in (d_1007_v24_) else d_1002_cf19_), len(d_1004_v21_)), d_1008_v25_): d_1008_v25_})
            d_1009_v26_ = (d_1009_v26_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ja")), d_1008_v25_)
            (globalState).f16 = d_973_v4_
            d_1003_cf18_ = (self).fm1(globalState)
            d_1010_v27_: _dafny.Map
            d_1010_v27_ = _dafny.Map({d_1003_cf18_: (d_969_v0_).f32})
            d_1011_v28_: _dafny.Map
            d_1011_v28_ = _dafny.Map({(d_1010_v27_).set(d_1003_cf18_, (self).f26): d_973_v4_})
            d_1003_cf18_ = len(d_1011_v28_)
        elif True:
            d_1012___mcc_h7_ = source17_.cf13
            d_1013_cf13_ = d_1012___mcc_h7_
            (globalState).f16 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_1014_i0_ in range(default__.abs(535))]))
            d_1015_v29_: _dafny.Seq
            d_1015_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cc"))
            d_1016_v30_: _dafny.Map
            d_1016_v30_ = _dafny.Map({(d_977_v8_).f32: d_973_v4_})
            d_1017_v31_: _dafny.Seq
            d_1017_v31_ = _dafny.SeqWithoutIsStrInference([d_973_v4_, d_973_v4_, d_973_v4_])
            d_1018_v32_: _dafny.Map
            d_1018_v32_ = _dafny.Map({d_973_v4_: (d_977_v8_).f32})
            d_1019_v33_: D1
            d_1019_v33_ = D1_DC5(d_973_v4_)
            d_1020_v34_: _dafny.Array
            nw148_ = _dafny.Array(None, 21)
            nw148_[int(0)] = d_973_v4_
            nw148_[int(1)] = 41
            nw148_[int(2)] = d_973_v4_
            nw148_[int(3)] = d_973_v4_
            nw148_[int(4)] = d_973_v4_
            nw148_[int(5)] = d_973_v4_
            nw148_[int(6)] = d_973_v4_
            nw148_[int(7)] = 887
            nw148_[int(8)] = d_973_v4_
            nw148_[int(9)] = d_973_v4_
            nw148_[int(10)] = len(d_1015_v29_)
            nw148_[int(11)] = d_973_v4_
            nw148_[int(12)] = d_973_v4_
            nw148_[int(13)] = d_973_v4_
            nw148_[int(14)] = ((d_1016_v30_)[(d_969_v0_).f32] if ((d_969_v0_).f32) in (d_1016_v30_) else d_973_v4_)
            nw148_[int(15)] = d_973_v4_
            nw148_[int(16)] = (d_1017_v31_)[default__.safeIndex(d_973_v4_, len(d_1017_v31_))]
            nw148_[int(17)] = len(d_1015_v29_)
            nw148_[int(18)] = d_973_v4_
            nw148_[int(19)] = len(d_1018_v32_)
            nw148_[int(20)] = (d_1019_v33_).cf12
            d_1020_v34_ = nw148_
            d_1021_v35_: D3
            d_1021_v35_ = D3_DC10(d_1020_v34_)
            d_1022_v36_: _dafny.Array
            nw149_ = _dafny.Array(None, 25)
            nw149_[int(0)] = d_1020_v34_
            nw149_[int(1)] = d_1020_v34_
            nw149_[int(2)] = d_1020_v34_
            nw149_[int(3)] = d_1020_v34_
            nw149_[int(4)] = (d_1021_v35_).cf21
            nw149_[int(5)] = d_1020_v34_
            nw149_[int(6)] = d_1020_v34_
            nw149_[int(7)] = d_1020_v34_
            nw149_[int(8)] = d_1020_v34_
            nw149_[int(9)] = d_1020_v34_
            nw149_[int(10)] = d_1020_v34_
            nw149_[int(11)] = d_1020_v34_
            nw149_[int(12)] = d_1020_v34_
            nw149_[int(13)] = d_1020_v34_
            nw149_[int(14)] = d_1020_v34_
            nw149_[int(15)] = d_1020_v34_
            nw149_[int(16)] = d_1020_v34_
            nw149_[int(17)] = d_1020_v34_
            nw149_[int(18)] = d_1020_v34_
            nw149_[int(19)] = d_1020_v34_
            nw149_[int(20)] = d_1020_v34_
            nw149_[int(21)] = d_1020_v34_
            nw149_[int(22)] = d_1020_v34_
            nw149_[int(23)] = d_1020_v34_
            nw149_[int(24)] = d_1020_v34_
            d_1022_v36_ = nw149_
            d_1022_v36_ = d_1022_v36_
            d_1023_v37_: _dafny.Map
            d_1023_v37_ = _dafny.Map({d_1015_v29_: d_973_v4_})
            d_1024_v38_: _dafny.MultiSet
            d_1024_v38_ = _dafny.MultiSet([p0, (d_969_v0_).fm6(_dafny.CodePoint('j'), (d_977_v8_).f32, globalState), (d_969_v0_).f32, (p2)[default__.safeIndex(len((d_1023_v37_).set(d_1015_v29_, d_973_v4_)), len(p2))], (d_977_v8_).f32])
            d_1025_v39_: _dafny.MultiSet
            d_1025_v39_ = _dafny.MultiSet([((d_1024_v38_).cardinality) + (d_973_v4_), d_973_v4_])
            d_1026_v40_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_1026_v40_ = nw150_
            d_1027_v41_: D1
            d_1027_v41_ = D1_DC4(d_1026_v40_, d_973_v4_, d_973_v4_, d_1015_v29_, d_1015_v29_)
            d_1028_v42_: _dafny.Map
            d_1028_v42_ = _dafny.Map({p0: (d_969_v0_).f32})
            pat_let_tv19_ = d_1028_v42_
            d_1029_v43_: _dafny.Map
            def iife84_(_pat_let13_0):
                def iife85_(d_1030_dt__update__tmp_h0_):
                    def iife86_(_pat_let14_0):
                        def iife87_(d_1031_dt__update_hcf8_h0_):
                            return D1_DC4((d_1030_dt__update__tmp_h0_).cf7, d_1031_dt__update_hcf8_h0_, (d_1030_dt__update__tmp_h0_).cf9, (d_1030_dt__update__tmp_h0_).cf10, (d_1030_dt__update__tmp_h0_).cf11)
                        return iife87_(_pat_let14_0)
                    return iife86_(len(pat_let_tv19_))
                return iife85_(_pat_let13_0)
            d_1029_v43_ = _dafny.Map({(d_969_v0_).f32: (iife84_(d_1027_v41_)).cf10})
            d_1032_v44_: D0
            d_1032_v44_ = D0_DC1((d_969_v0_).f32, (self).f26, d_973_v4_, not((self).f26))
            d_1025_v39_ = (default__.fm10(d_973_v4_, len(d_1029_v43_), D0_DC2(d_1032_v44_), globalState)) | (d_1025_v39_)
            if (self).f26:
                (globalState).f16 = (((_dafny.MultiSet([(self).f26])) | (d_1024_v38_)) - (d_1024_v38_)).cardinality
                (globalState).f16 = d_973_v4_
                d_1033_v45_: _dafny.Array
                nw151_ = _dafny.Array(None, 3)
                nw151_[int(0)] = p0
                nw151_[int(1)] = (d_969_v0_).f32
                nw151_[int(2)] = (d_969_v0_).f32
                d_1033_v45_ = nw151_
                index144_ = default__.safeIndex(752, (d_1033_v45_).length(0))
                (d_1033_v45_)[index144_] = (self).f26
                d_1034_v46_: D0
                d_1034_v46_ = D0_DC0((d_969_v0_).f32)
                index145_ = default__.safeIndex(752, (d_1033_v45_).length(0))
                rhs161_ = not(p0)
                rhs162_ = ((d_969_v0_).f32 if (d_1034_v46_).cf0 else not((self).f27))
                lhs149_ = d_1033_v45_
                lhs150_ = default__.safeIndex(752, (d_1033_v45_).length(0))
                lhs151_ = globalState
                lhs149_[lhs150_] = rhs161_
                lhs151_.f24 = rhs162_
                d_1020_v34_ = d_1020_v34_
                (globalState).f12 = default__.safeModuloInt((0) - (d_973_v4_), d_973_v4_)
            elif True:
                d_1035_v47_: _dafny.Array
                nw152_ = _dafny.Array(False, 14)
                d_1035_v47_ = nw152_
                index146_ = default__.safeIndex(358, (d_1035_v47_).length(0))
                (d_1035_v47_)[index146_] = (d_1025_v39_).ispropersubset(d_1025_v39_)
                d_1036_v48_: str
                d_1036_v48_ = _dafny.CodePoint('d')
                index147_ = default__.safeIndex(358, (d_1035_v47_).length(0))
                (d_1035_v47_)[index147_] = (d_977_v8_).fm6(d_1036_v48_, (d_977_v8_).f32, globalState)
                d_1037_v49_: D2
                d_1037_v49_ = D2_DC8()
                rhs163_ = d_1037_v49_
                rhs164_ = True
                lhs152_ = globalState
                d_1037_v49_ = rhs163_
                lhs152_.f24 = rhs164_
                d_1038_v50_: D0
                d_1038_v50_ = D0_DC0(True)
                d_1039_v51_: bool
                out41_: bool
                out41_ = (self).m8(d_1038_v50_, (self).f27, (_dafny.MultiSet([(d_969_v0_).f32])).set((d_969_v0_).f32, default__.abs(d_973_v4_)), d_973_v4_, globalState)
                d_1039_v51_ = out41_
                index148_ = default__.safeIndex(616, (d_1020_v34_).length(0))
                (d_1020_v34_)[index148_] = d_973_v4_
                index149_ = default__.safeIndex(616, (d_1020_v34_).length(0))
                (d_1020_v34_)[index149_] = ((d_973_v4_ if (d_1035_v47_)[default__.safeIndex(358, (d_1035_v47_).length(0))] else d_973_v4_)) - (len(d_1015_v29_))
                d_1036_v48_ = d_1036_v48_
        d_1040_v52_: str
        d_1040_v52_ = _dafny.CodePoint('e')
        d_1041_v53_: D1
        d_1041_v53_ = D1_DC3(default__.fm11((d_977_v8_).f32, (d_969_v0_).f32, d_1040_v52_, globalState))
        pat_let_tv20_ = p2
        pat_let_tv21_ = d_973_v4_
        pat_let_tv22_ = p2
        pat_let_tv23_ = d_969_v0_
        def iife88_(_pat_let15_0):
            def iife89_(d_1042_dt__update__tmp_h1_):
                def iife90_(_pat_let16_0):
                    def iife91_(d_1043_dt__update_hcf6_h0_):
                        return D1_DC3(d_1043_dt__update_hcf6_h0_)
                    return iife91_(_pat_let16_0)
                return iife90_((pat_let_tv20_).set(default__.safeIndex(pat_let_tv21_, len(pat_let_tv22_)), (pat_let_tv23_).f32))
            return iife89_(_pat_let15_0)
        source18_ = iife88_(d_1041_v53_)
        if source18_.is_DC4:
            d_1044___mcc_h8_ = source18_.cf7
            d_1045___mcc_h9_ = source18_.cf8
            d_1046___mcc_h10_ = source18_.cf9
            d_1047___mcc_h11_ = source18_.cf10
            d_1048___mcc_h12_ = source18_.cf11
            d_1049_cf11_ = d_1048___mcc_h12_
            d_1050_cf10_ = d_1047___mcc_h11_
            d_1051_cf9_ = d_1046___mcc_h10_
            d_1052_cf8_ = d_1045___mcc_h9_
            d_1053_cf7_ = d_1044___mcc_h8_
            d_1054_v54_: _dafny.Array
            nw153_ = _dafny.Array(_dafny.Array(None, 0), 4)
            d_1054_v54_ = nw153_
            d_1055_v55_: _dafny.Map
            d_1055_v55_ = _dafny.Map({d_1054_v54_: (self).fm1(globalState)})
            d_1056_v56_: _dafny.Map
            d_1056_v56_ = _dafny.Map({False: _dafny.Set({True, (d_969_v0_).f32, (self).f27})})
            d_1055_v55_ = (d_1055_v55_).set(d_1054_v54_, (len(d_1056_v56_)) + (d_1052_cf8_))
            d_1057_v57_: _dafny.Set
            d_1057_v57_ = _dafny.Set({not((d_977_v8_).f32), (self).f27})
            d_1058_v58_: _dafny.Set
            d_1058_v58_ = d_1057_v57_
            (globalState).f2 = (d_1057_v57_).ispropersubset((d_1058_v58_))
            d_973_v4_ = d_1051_cf9_
            d_1059_v59_: D2
            d_1059_v59_ = D2_DC8()
            source19_ = d_1059_v59_
            if source19_.is_DC7:
                d_1060___mcc_h15_ = source19_.cf14
                d_1061___mcc_h16_ = source19_.cf15
                d_1062___mcc_h17_ = source19_.cf16
                d_1063___mcc_h18_ = source19_.cf17
                d_1064_cf17_ = d_1063___mcc_h18_
                d_1065_cf16_ = d_1062___mcc_h17_
                d_1066_cf15_ = d_1061___mcc_h16_
                d_1067_cf14_ = d_1060___mcc_h15_
                d_1068_v60_: _dafny.Array
                def lambda55_(d_1069_p2_):
                    def lambda56_(d_1070_i1_):
                        return default__.safeModuloInt(d_1070_i1_, len(d_1069_p2_))

                    return lambda56_

                init30_ = lambda55_(p2)
                nw154_ = _dafny.Array(None, 18)
                for i0_30_ in range(nw154_.length(0)):
                    nw154_[i0_30_] = init30_(i0_30_)
                d_1068_v60_ = nw154_
                d_1071_v61_: _dafny.Map
                d_1071_v61_ = _dafny.Map({d_1068_v60_: (self).fm1(globalState)})
                d_1071_v61_ = (d_1071_v61_).set(d_1068_v60_, (d_1064_cf17_) - (len(d_1050_cf10_)))
                (globalState).f6 = (_dafny.SeqWithoutIsStrInference([d_1040_v52_ for d_1072_i2_ in range(default__.abs(45))])) + (d_1049_cf11_)
                d_977_v8_ = d_977_v8_
                (globalState).f6 = _dafny.SeqWithoutIsStrInference([d_1040_v52_ for d_1073_i3_ in range(default__.abs(830))])
            elif source19_.is_DC8:
                d_1074_v62_: _dafny.Seq
                d_1074_v62_ = _dafny.SeqWithoutIsStrInference([((_dafny.MultiSet([d_1040_v52_])).set(d_1040_v52_, default__.abs(204))).cardinality])
                d_1075_v63_: _dafny.Array
                nw155_ = _dafny.Array(False, 23)
                d_1075_v63_ = nw155_
                d_1076_v64_: _dafny.Map
                d_1076_v64_ = _dafny.Map({d_1075_v63_: (d_1074_v62_).set(default__.safeIndex(d_973_v4_, len(d_1074_v62_)), d_1051_cf9_)})
                d_1077_v65_: _dafny.Map
                d_1077_v65_ = _dafny.Map({d_1074_v62_: ((d_1076_v64_)[d_1075_v63_] if (d_1075_v63_) in (d_1076_v64_) else d_1074_v62_)})
                (globalState).f15 = not((d_1077_v65_) != (d_1077_v65_))
                d_1078_v66_: _dafny.Map
                d_1078_v66_ = _dafny.Map({d_1051_cf9_: d_1074_v62_})
                d_1078_v66_ = (d_1078_v66_).set(d_973_v4_, ((_dafny.SeqWithoutIsStrInference([d_1051_cf9_ for d_1079_i4_ in range(default__.abs(814))])) + (d_1074_v62_)).set(default__.safeIndex(d_973_v4_, len((_dafny.SeqWithoutIsStrInference([d_1051_cf9_ for d_1080_i4_ in range(default__.abs(814))])) + (d_1074_v62_))), d_973_v4_))
                d_1081_v67_: _dafny.Map
                d_1081_v67_ = _dafny.Map({d_973_v4_: d_1052_cf8_})
                d_1081_v67_ = (d_1081_v67_).set(-944, d_973_v4_)
                (globalState).f2 = (d_969_v0_).f32
            elif source19_.is_DC9:
                d_1082___mcc_h19_ = source19_.cf18
                d_1083___mcc_h20_ = source19_.cf19
                d_1084___mcc_h21_ = source19_.cf20
                d_1085_cf20_ = d_1084___mcc_h21_
                d_1086_cf19_ = d_1083___mcc_h20_
                d_1087_cf18_ = d_1082___mcc_h19_
                (globalState).f2 = not (d_1085_cf20_) or ((d_1087_cf18_) < (d_1052_cf8_))
                index150_ = default__.safeIndex(114, (d_1053_cf7_).length(0))
                (d_1053_cf7_)[index150_] = (d_1040_v52_ if (d_977_v8_).f32 else d_1040_v52_)
                d_1088_v68_: _dafny.Map
                d_1088_v68_ = _dafny.Map({d_1085_cf20_: (d_969_v0_).f32})
                d_1089_v69_: _dafny.MultiSet
                d_1089_v69_ = _dafny.MultiSet([d_973_v4_, (0) - ((d_1052_cf8_) - (d_973_v4_)), (0) - ((len(d_1088_v68_) if False else d_1087_cf18_)), (d_1051_cf9_) * (194)])
                index151_ = default__.safeIndex(114, (d_1053_cf7_).length(0))
                rhs165_ = d_1040_v52_
                rhs166_ = ((d_1089_v69_).intersection(d_1089_v69_)) | ((d_1089_v69_).intersection(_dafny.MultiSet([len(_dafny.Map({d_973_v4_: p0}))])))
                rhs167_ = d_1087_cf18_
                lhs153_ = d_1053_cf7_
                lhs154_ = default__.safeIndex(114, (d_1053_cf7_).length(0))
                lhs155_ = globalState
                lhs153_[lhs154_] = rhs165_
                d_1089_v69_ = rhs166_
                lhs155_.f12 = rhs167_
                d_1090_v70_: _dafny.Array
                def lambda57_(d_1091_v4_):
                    def lambda58_(d_1092_i5_):
                        return default__.safeDivisionInt(d_1092_i5_, d_1091_v4_)

                    return lambda58_

                init31_ = lambda57_(d_973_v4_)
                nw156_ = _dafny.Array(None, 1)
                for i0_31_ in range(nw156_.length(0)):
                    nw156_[i0_31_] = init31_(i0_31_)
                d_1090_v70_ = nw156_
                d_1093_v71_: D3
                d_1093_v71_ = D3_DC10(d_1090_v70_)
                d_1094_v72_: _dafny.Map
                d_1094_v72_ = _dafny.Map({d_1093_v71_: (self).fm1(globalState)})
                d_1095_v73_: T1
                nw157_ = C4()
                nw157_.ctor__(d_1094_v72_, (d_977_v8_).f32, (self).f27)
                d_1095_v73_ = nw157_
                d_1096_v74_: _dafny.Map
                d_1096_v74_ = _dafny.Map({d_1052_cf8_: d_1095_v73_})
                d_1097_v75_: _dafny.Array
                nw158_ = _dafny.Array(None, 15)
                nw158_[int(0)] = d_1087_cf18_
                nw158_[int(1)] = (self).fm1(globalState)
                nw158_[int(2)] = len(d_1057_v57_)
                nw158_[int(3)] = d_1086_cf19_
                nw158_[int(4)] = 507
                nw158_[int(5)] = len(d_1050_cf10_)
                nw158_[int(6)] = d_1086_cf19_
                nw158_[int(7)] = (0) - (d_1052_cf8_)
                nw158_[int(8)] = d_1051_cf9_
                nw158_[int(9)] = (D0_DC1((d_969_v0_).f32, (d_977_v8_).f32, len(d_1050_cf10_), (d_977_v8_).f32)).cf3
                nw158_[int(10)] = d_1086_cf19_
                nw158_[int(11)] = len(d_1096_v74_)
                nw158_[int(12)] = d_1086_cf19_
                nw158_[int(13)] = (0) - (d_1052_cf8_)
                nw158_[int(14)] = d_1086_cf19_
                d_1097_v75_ = nw158_
                d_1098_v76_: _dafny.MultiSet
                d_1098_v76_ = _dafny.MultiSet([d_1097_v75_, d_1097_v75_, d_1090_v70_])
                (globalState).f16 = (d_1098_v76_).cardinality
                d_1099_v77_: C1
                nw159_ = C1()
                nw159_.ctor__((d_1052_cf8_) + (d_1051_cf9_), (d_1051_cf9_) - (d_1051_cf9_), (d_1052_cf8_) >= (d_973_v4_), (not((d_977_v8_).f32)) or (d_1085_cf20_))
                d_1099_v77_ = nw159_
            elif True:
                d_1100___mcc_h22_ = source19_.cf13
                d_1101_cf13_ = d_1100___mcc_h22_
                d_1102_v78_: C1
                nw160_ = C1()
                nw160_.ctor__(len(p2), d_1051_cf9_, p0, (self).f27)
                d_1102_v78_ = nw160_
                d_1103_v79_: _dafny.Set
                d_1103_v79_ = _dafny.Set({d_1102_v78_, d_1102_v78_, d_1102_v78_})
                d_1103_v79_ = (d_1103_v79_) - (d_1103_v79_)
                (globalState).f22 = (True) and ((d_977_v8_).f32)
                (globalState).f12 = (d_1102_v78_).f36
                (globalState).f6 = (d_1049_cf11_) + ((d_1050_cf10_ if (d_969_v0_).f32 else d_1049_cf11_))
        elif source18_.is_DC5:
            d_1104___mcc_h13_ = source18_.cf12
            d_1105_cf12_ = d_1104___mcc_h13_
            d_1106_v80_: _dafny.Map
            d_1106_v80_ = _dafny.Map({(d_977_v8_).f32: _dafny.SeqWithoutIsStrInference([(d_969_v0_).f32, (d_969_v0_).f32])})
            d_1107_v81_: D5
            d_1107_v81_ = D5_DC13((default__.fm32((self).fm1(globalState), d_1105_cf12_, globalState) if True else (d_1106_v80_).set(p0, p2)))
            source20_ = d_1107_v81_
            if source20_.is_DC14:
                d_1108___mcc_h23_ = source20_.cf25
                d_1109___mcc_h24_ = source20_.cf26
                d_1110___mcc_h25_ = source20_.cf27
                d_1111_cf27_ = d_1110___mcc_h25_
                d_1112_cf26_ = d_1109___mcc_h24_
                d_1113_cf25_ = d_1108___mcc_h23_
                (globalState).f16 = (0) - (d_1112_cf26_)
                rhs168_ = d_1112_cf26_
                rhs169_ = 294
                rhs170_ = (d_969_v0_).f32
                lhs156_ = globalState
                lhs157_ = globalState
                lhs158_ = globalState
                lhs156_.f16 = rhs168_
                lhs157_.f12 = rhs169_
                lhs158_.f24 = rhs170_
                d_1114_v82_: _dafny.Array
                def lambda59_(d_1115_i6_):
                    return D2_DC8()

                init32_ = lambda59_
                nw161_ = _dafny.Array(None, 1)
                for i0_32_ in range(nw161_.length(0)):
                    nw161_[i0_32_] = init32_(i0_32_)
                d_1114_v82_ = nw161_
                index152_ = default__.safeIndex(247, (d_1114_v82_).length(0))
                (d_1114_v82_)[index152_] = D2_DC8()
                d_1116_v83_: _dafny.Seq
                d_1116_v83_ = _dafny.SeqWithoutIsStrInference([d_1040_v52_])
                d_1117_v84_: _dafny.Set
                d_1117_v84_ = _dafny.Set({p0, False})
                d_1118_v85_: _dafny.Map
                d_1118_v85_ = _dafny.Map({d_1116_v83_: len(d_1117_v84_)})
                index153_ = default__.safeIndex(247, (d_1114_v82_).length(0))
                (d_1114_v82_)[index153_] = default__.fm33((0) - ((743 if (self).f26 else d_1113_cf25_)), len((d_1118_v85_) | ((d_1118_v85_).set(d_1116_v83_, d_1105_cf12_))), globalState)
                (globalState).f16 = (0) - (d_1113_cf25_)
            elif source20_.is_DC13:
                d_1119___mcc_h26_ = source20_.cf24
                d_1120_cf24_ = d_1119___mcc_h26_
                d_1121_v86_: _dafny.Array
                nw162_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1121_v86_ = nw162_
                d_1122_v87_: _dafny.Map
                d_1122_v87_ = _dafny.Map({-46: d_1105_cf12_})
                d_1123_v88_: _dafny.Seq
                d_1123_v88_ = _dafny.SeqWithoutIsStrInference([d_1122_v87_])
                index154_ = default__.safeIndex(522, (d_1121_v86_).length(0))
                (d_1121_v86_)[index154_] = d_1123_v88_
                index155_ = default__.safeIndex(522, (d_1121_v86_).length(0))
                (d_1121_v86_)[index155_] = d_1123_v88_
                d_1124_v89_: _dafny.Array
                nw163_ = _dafny.Array(int(0), 12)
                d_1124_v89_ = nw163_
                d_1124_v89_ = d_1124_v89_
                d_1105_cf12_ = d_973_v4_
                d_1125_v90_: _dafny.Seq
                d_1125_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "priuoxrge"))
                d_1126_v91_: _dafny.MultiSet
                d_1126_v91_ = _dafny.MultiSet([(p1)[default__.safeIndex(d_973_v4_, len(p1))], d_1125_v90_])
                d_1126_v91_ = d_1126_v91_
            elif True:
                d_1127___mcc_h27_ = source20_.cf28
                d_1128_cf28_ = d_1127___mcc_h27_
                d_1129_v92_: _dafny.MultiSet
                d_1129_v92_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_977_v8_).f32])])
                d_1130_v93_: _dafny.Array
                def lambda60_(d_1131_v0_):
                    def lambda61_(d_1132_i7_):
                        return (d_1131_v0_).f32

                    return lambda61_

                init33_ = lambda60_(d_969_v0_)
                nw164_ = _dafny.Array(None, 13)
                for i0_33_ in range(nw164_.length(0)):
                    nw164_[i0_33_] = init33_(i0_33_)
                d_1130_v93_ = nw164_
                d_1133_v94_: _dafny.MultiSet
                d_1133_v94_ = _dafny.MultiSet([d_1130_v93_, d_1130_v93_, d_1130_v93_])
                d_1134_v95_: _dafny.Map
                d_1134_v95_ = _dafny.Map({(self).f27: d_973_v4_})
                d_1135_v96_: _dafny.Seq
                d_1135_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dq"))
                d_1136_v97_: _dafny.Seq
                d_1136_v97_ = _dafny.SeqWithoutIsStrInference([len(d_1135_v96_)])
                d_1137_v98_: _dafny.Set
                d_1137_v98_ = _dafny.Set({(d_969_v0_).f32, (self).f26, False, (609) == (((d_1129_v92_).set(p2, default__.abs(((d_1133_v94_)[d_1130_v93_] if (d_1130_v93_) in (d_1133_v94_) else d_973_v4_)))).cardinality), (d_969_v0_).fm7(d_1134_v95_, d_1136_v97_, (self).fm2((self).f26, (0) - (len(d_1135_v96_)), globalState), globalState)})
                rhs171_ = d_1137_v98_
                rhs172_ = d_978_v9_
                rhs173_ = d_1136_v97_
                d_1137_v98_ = rhs171_
                d_978_v9_ = rhs172_
                d_1136_v97_ = rhs173_
                d_1138_v99_: _dafny.MultiSet
                d_1138_v99_ = _dafny.MultiSet([d_1136_v97_, d_1136_v97_, d_1136_v97_])
                d_1139_v100_: C3
                nw165_ = C3()
                nw165_.ctor__((d_1138_v99_).set(d_1136_v97_, default__.abs(d_973_v4_)), (d_973_v4_) <= (d_1105_cf12_), ((d_977_v8_).f32) == (not((d_977_v8_).f32)))
                d_1139_v100_ = nw165_
                d_1136_v97_ = ((d_1136_v97_) + (_dafny.SeqWithoutIsStrInference([d_1105_cf12_, d_1105_cf12_]))) + (d_1136_v97_)
                (globalState).f22 = (d_977_v8_).f32
            source21_ = D1_DC3(_dafny.SeqWithoutIsStrInference([True, (self).f26, (d_969_v0_).f32]))
            if source21_.is_DC4:
                d_1140___mcc_h28_ = source21_.cf7
                d_1141___mcc_h29_ = source21_.cf8
                d_1142___mcc_h30_ = source21_.cf9
                d_1143___mcc_h31_ = source21_.cf10
                d_1144___mcc_h32_ = source21_.cf11
                d_1145_cf11_ = d_1144___mcc_h32_
                d_1146_cf10_ = d_1143___mcc_h31_
                d_1147_cf9_ = d_1142___mcc_h30_
                d_1148_cf8_ = d_1141___mcc_h29_
                d_1149_cf7_ = d_1140___mcc_h28_
                d_1150_v101_: _dafny.Array
                nw166_ = _dafny.Array(None, 11)
                nw166_[int(0)] = d_1148_cf8_
                nw166_[int(1)] = -155
                nw166_[int(2)] = d_1147_cf9_
                nw166_[int(3)] = d_1148_cf8_
                nw166_[int(4)] = 272
                nw166_[int(5)] = 842
                nw166_[int(6)] = d_973_v4_
                nw166_[int(7)] = d_973_v4_
                nw166_[int(8)] = d_1147_cf9_
                nw166_[int(9)] = len(d_1145_cf11_)
                nw166_[int(10)] = d_973_v4_
                d_1150_v101_ = nw166_
                d_1151_v102_: D3
                d_1151_v102_ = D3_DC10(d_1150_v101_)
                d_1152_v103_: _dafny.Map
                d_1152_v103_ = _dafny.Map({d_1151_v102_: d_973_v4_})
                d_1153_v104_: C4
                nw167_ = C4()
                nw167_.ctor__(d_1152_v103_, (d_969_v0_).f32, (d_977_v8_).f32)
                d_1153_v104_ = nw167_
                d_1154_v105_: _dafny.Seq
                d_1154_v105_ = _dafny.SeqWithoutIsStrInference([d_1153_v104_])
                d_1155_v106_: _dafny.Set
                d_1155_v106_ = _dafny.Set({(d_969_v0_).f32})
                d_1156_v107_: _dafny.Set
                d_1156_v107_ = _dafny.Set({len(d_1155_v106_)})
                d_1157_v108_: _dafny.MultiSet
                d_1157_v108_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_973_v4_ for d_1158_i8_ in range(default__.abs(231))])])
                d_1159_v109_: C3
                nw168_ = C3()
                nw168_.ctor__(d_1157_v108_, (d_977_v8_).f32, not(p0))
                d_1159_v109_ = nw168_
                d_1160_v110_: _dafny.Map
                d_1160_v110_ = _dafny.Map({(_dafny.Set({len(d_1154_v105_), d_1148_cf8_})).issubset(d_1156_v107_): d_1159_v109_})
                d_1160_v110_ = d_1160_v110_
                (globalState).f22 = not((len((d_1145_cf11_) + (d_1145_cf11_))) == (((0) - (len(_dafny.SeqWithoutIsStrInference([d_1040_v52_ for d_1161_i9_ in range(default__.abs(77))]))) if (d_977_v8_).f32 else d_973_v4_)))
                d_1162_v111_: _dafny.Array
                def lambda62_(d_1163_p2_, d_1164_v0_):
                    def lambda63_(d_1165_i10_):
                        return (d_1163_p2_) + (_dafny.SeqWithoutIsStrInference([(d_1164_v0_).f32, (self).f27]))

                    return lambda63_

                init34_ = lambda62_(p2, d_969_v0_)
                nw169_ = _dafny.Array(None, 17)
                for i0_34_ in range(nw169_.length(0)):
                    nw169_[i0_34_] = init34_(i0_34_)
                d_1162_v111_ = nw169_
                index156_ = default__.safeIndex(267, (d_1162_v111_).length(0))
                (d_1162_v111_)[index156_] = (p2) + (p2)
                index157_ = default__.safeIndex(267, (d_1162_v111_).length(0))
                (d_1162_v111_)[index157_] = p2
                d_1166_v112_: _dafny.Map
                d_1166_v112_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1040_v52_ for d_1167_i11_ in range(default__.abs(120))]): d_1148_cf8_})
                d_1166_v112_ = (d_1166_v112_).set((d_1145_cf11_) + (d_1146_cf10_), d_973_v4_)
            elif source21_.is_DC5:
                d_1168___mcc_h33_ = source21_.cf12
                d_1169_cf12_ = d_1168___mcc_h33_
                d_1170_v113_: _dafny.Seq
                d_1170_v113_ = _dafny.SeqWithoutIsStrInference([d_973_v4_, d_1169_cf12_])
                d_1171_v114_: _dafny.Seq
                d_1171_v114_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhtyej"))
                d_1172_v115_: _dafny.MultiSet
                d_1172_v115_ = _dafny.MultiSet([d_1171_v114_, d_1171_v114_])
                d_1173_v116_: _dafny.Seq
                d_1173_v116_ = _dafny.SeqWithoutIsStrInference([d_975_v7_, d_975_v7_, d_975_v7_])
                d_1174_v117_: _dafny.Map
                d_1174_v117_ = _dafny.Map({(d_1170_v113_).set(default__.safeIndex((d_1172_v115_).cardinality, len(d_1170_v113_)), d_1169_cf12_): d_1173_v116_})
                d_1175_v118_: _dafny.Map
                d_1175_v118_ = _dafny.Map({(d_977_v8_).f32: _dafny.Map({d_1170_v113_: d_1173_v116_})})
                d_1176_v119_: D10
                d_1176_v119_ = D10_DC23(_dafny.Map({d_1170_v113_: d_1173_v116_}))
                d_1177_v120_: _dafny.Array
                nw170_ = _dafny.Array(None, 23)
                nw170_[int(0)] = d_1174_v117_
                nw170_[int(1)] = ((d_1175_v118_)[True] if (True) in (d_1175_v118_) else d_1174_v117_)
                nw170_[int(2)] = d_1174_v117_
                nw170_[int(3)] = d_1174_v117_
                nw170_[int(4)] = d_1174_v117_
                nw170_[int(5)] = d_1174_v117_
                nw170_[int(6)] = (d_1174_v117_).set(default__.fm15(globalState), _dafny.SeqWithoutIsStrInference([d_975_v7_]))
                nw170_[int(7)] = (d_1174_v117_) | (d_1174_v117_)
                nw170_[int(8)] = (default__.fm34(globalState)).set(_dafny.SeqWithoutIsStrInference([d_973_v4_]), d_1173_v116_)
                nw170_[int(9)] = (d_1174_v117_) | (d_1174_v117_)
                nw170_[int(10)] = d_1174_v117_
                nw170_[int(11)] = _dafny.Map({d_1170_v113_: d_1173_v116_})
                nw170_[int(12)] = d_1174_v117_
                nw170_[int(13)] = (d_1174_v117_).set(d_1170_v113_, _dafny.SeqWithoutIsStrInference([d_975_v7_ for d_1178_i12_ in range(default__.abs(721))]))
                nw170_[int(14)] = d_1174_v117_
                nw170_[int(15)] = (d_1176_v119_).cf43
                nw170_[int(16)] = d_1174_v117_
                nw170_[int(17)] = d_1174_v117_
                nw170_[int(18)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1169_cf12_ for d_1179_i13_ in range(default__.abs(892))]): d_1173_v116_})
                nw170_[int(19)] = (d_1174_v117_) | (d_1174_v117_)
                nw170_[int(20)] = (d_1174_v117_).set(_dafny.SeqWithoutIsStrInference([d_1169_cf12_]), d_1173_v116_)
                nw170_[int(21)] = d_1174_v117_
                nw170_[int(22)] = (d_1174_v117_) | (d_1174_v117_)
                d_1177_v120_ = nw170_
                index158_ = default__.safeIndex(891, (d_1177_v120_).length(0))
                (d_1177_v120_)[index158_] = _dafny.Map({d_1170_v113_: d_1173_v116_})
                d_1180_v122_: _dafny.Map
                d_1180_v122_ = _dafny.Map({d_1170_v113_: d_973_v4_})
                index159_ = default__.safeIndex(891, (d_1177_v120_).length(0))
                def iife92_():
                    coll58_ = _dafny.Map()
                    compr_58_: _dafny.Seq
                    for compr_58_ in (d_1180_v122_).keys.Elements:
                        d_1181_v121_: _dafny.Seq = compr_58_
                        if (d_1181_v121_) in (d_1180_v122_):
                            coll58_[d_1181_v121_] = d_1173_v116_
                    return _dafny.Map(coll58_)
                (d_1177_v120_)[index159_] = ((d_1174_v117_) | ((iife92_()
                ).set(_dafny.SeqWithoutIsStrInference([d_1169_cf12_]), d_1173_v116_))) | (d_1174_v117_)
                rhs174_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ertaynatq")))) - (893)
                rhs175_ = d_1169_cf12_
                lhs159_ = globalState
                lhs160_ = globalState
                lhs159_.f16 = rhs174_
                lhs160_.f16 = rhs175_
                (globalState).f15 = (d_969_v0_).f32
                d_1182_v123_: _dafny.Array
                nw171_ = _dafny.Array(None, 12)
                d_1182_v123_ = nw171_
                d_1183_v124_: _dafny.MultiSet
                d_1183_v124_ = _dafny.MultiSet([d_1182_v123_, d_1182_v123_])
                rhs176_ = d_1183_v124_
                rhs177_ = (d_978_v9_).cf19
                lhs161_ = globalState
                d_1183_v124_ = rhs176_
                lhs161_.f12 = rhs177_
            elif True:
                d_1184___mcc_h34_ = source21_.cf6
                d_1185_cf6_ = d_1184___mcc_h34_
                d_1186_v125_: _dafny.Seq
                d_1186_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsdmh"))
                d_1187_v126_: _dafny.Seq
                d_1187_v126_ = _dafny.SeqWithoutIsStrInference([(self).fm1(globalState)])
                d_1188_v127_: _dafny.Map
                d_1188_v127_ = _dafny.Map({(d_969_v0_).f32: d_973_v4_})
                d_1189_v128_: _dafny.Map
                d_1189_v128_ = _dafny.Map({p0: False})
                d_1190_v129_: _dafny.Array
                nw172_ = _dafny.Array(None, 15)
                nw172_[int(0)] = (d_1187_v126_)[default__.safeIndex(d_1105_cf12_, len(d_1187_v126_))]
                nw172_[int(1)] = (0) - (d_1105_cf12_)
                nw172_[int(2)] = (0) - (d_973_v4_)
                nw172_[int(3)] = (self).fm1(globalState)
                nw172_[int(4)] = d_1105_cf12_
                nw172_[int(5)] = len(d_1188_v127_)
                nw172_[int(6)] = 698
                nw172_[int(7)] = d_973_v4_
                nw172_[int(8)] = len(d_1189_v128_)
                nw172_[int(9)] = d_973_v4_
                nw172_[int(10)] = d_1105_cf12_
                nw172_[int(11)] = (0) - (d_973_v4_)
                nw172_[int(12)] = d_1105_cf12_
                nw172_[int(13)] = d_973_v4_
                nw172_[int(14)] = d_1105_cf12_
                d_1190_v129_ = nw172_
                d_1191_v130_: _dafny.Map
                d_1191_v130_ = _dafny.Map({d_1190_v129_: d_1186_v125_})
                d_1192_v131_: C1
                nw173_ = C1()
                nw173_.ctor__(len(d_1186_v125_), (0) - (len(d_1187_v126_)), (d_1190_v129_) in (d_1191_v130_), (d_977_v8_).f32)
                d_1192_v131_ = nw173_
                d_1193_v132_: D0
                d_1193_v132_ = D0_DC0((self).f26)
                (globalState).f22 = not ((d_969_v0_).f32) or ((self).fm2((self).f26, (d_1192_v131_).fm4(d_1193_v132_, (d_1192_v131_).f36, d_973_v4_, globalState), globalState))
                d_1194_v133_: _dafny.Map
                d_1194_v133_ = _dafny.Map({(d_1186_v125_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxenwl"))): d_1040_v52_})
                d_1194_v133_ = (d_1194_v133_).set(d_1186_v125_, d_1040_v52_)
                d_1188_v127_ = (d_1188_v127_).set((d_977_v8_).fm6(d_1040_v52_, (d_969_v0_).f32, globalState), 487)
            if (d_977_v8_).f32:
                d_1195_v134_: _dafny.Seq
                d_1195_v134_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ck"))
                d_1196_v135_: _dafny.MultiSet
                d_1196_v135_ = _dafny.MultiSet([len(d_1195_v134_)])
                d_1197_v136_: _dafny.Map
                d_1197_v136_ = _dafny.Map({d_1196_v135_: (d_969_v0_).f32})
                d_1198_v137_: _dafny.Set
                d_1198_v137_ = _dafny.Set({(d_977_v8_).f32, (d_969_v0_).f32})
                d_1199_v138_: _dafny.Map
                d_1199_v138_ = _dafny.Map({d_973_v4_: d_1198_v137_})
                d_1200_v140_: _dafny.Array
                def lambda64_(d_1201_i15_):
                    return (self).f26

                init35_ = lambda64_
                nw174_ = _dafny.Array(None, 15)
                for i0_35_ in range(nw174_.length(0)):
                    nw174_[i0_35_] = init35_(i0_35_)
                d_1200_v140_ = nw174_
                d_1202_v141_: _dafny.MultiSet
                d_1202_v141_ = _dafny.MultiSet([d_1200_v140_])
                d_1203_v142_: _dafny.Map
                d_1203_v142_ = _dafny.Map({d_973_v4_: ((d_1202_v141_)[d_1200_v140_] if (d_1200_v140_) in (d_1202_v141_) else d_1105_cf12_)})
                d_1204_v143_: _dafny.Map
                d_1204_v143_ = _dafny.Map({d_1203_v142_: 53})
                d_1205_v144_: _dafny.MultiSet
                d_1205_v144_ = _dafny.MultiSet([(d_969_v0_).f32])
                d_1206_v145_: _dafny.Array
                nw175_ = _dafny.Array(None, 24)
                nw175_[int(0)] = d_1105_cf12_
                nw175_[int(1)] = len(d_1197_v136_)
                nw175_[int(2)] = (d_973_v4_) * (d_973_v4_)
                nw175_[int(3)] = len(((d_1199_v138_)[255] if (255) in (d_1199_v138_) else _dafny.Set({(d_977_v8_).f32})))
                def iife93_():
                    coll59_ = _dafny.Set()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(158, 746):
                        d_1207_v139_: int = compr_59_
                        if ((158) <= (d_1207_v139_)) and ((d_1207_v139_) < (746)):
                            coll59_ = coll59_.union(_dafny.Set([default__.safeModuloInt(d_1207_v139_, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cuhflyu")))) for d_1208_i14_ in range(default__.abs(858))])))]))
                    return _dafny.Set(coll59_)
                nw175_[int(4)] = len(iife93_()
                )
                nw175_[int(5)] = (0) - (d_973_v4_)
                nw175_[int(6)] = d_1105_cf12_
                nw175_[int(7)] = d_973_v4_
                nw175_[int(8)] = d_1105_cf12_
                nw175_[int(9)] = d_1105_cf12_
                nw175_[int(10)] = d_973_v4_
                nw175_[int(11)] = (d_973_v4_) - (d_973_v4_)
                nw175_[int(12)] = d_973_v4_
                nw175_[int(13)] = default__.safeDivisionInt(d_1105_cf12_, d_1105_cf12_)
                nw175_[int(14)] = d_1105_cf12_
                nw175_[int(15)] = d_1105_cf12_
                nw175_[int(16)] = (d_1105_cf12_) - (len(d_1204_v143_))
                nw175_[int(17)] = d_1105_cf12_
                nw175_[int(18)] = (d_1205_v144_).cardinality
                nw175_[int(19)] = d_973_v4_
                nw175_[int(20)] = len(d_1203_v142_)
                nw175_[int(21)] = d_973_v4_
                nw175_[int(22)] = ((self).fm1(globalState)) * (d_973_v4_)
                nw175_[int(23)] = ((0) - (d_973_v4_)) * (d_1105_cf12_)
                d_1206_v145_ = nw175_
                index160_ = default__.safeIndex(854, (d_1206_v145_).length(0))
                (d_1206_v145_)[index160_] = d_1105_cf12_
                index161_ = default__.safeIndex(854, (d_1206_v145_).length(0))
                (d_1206_v145_)[index161_] = (len(d_1195_v134_) if not(((self).f27) == ((self).f27)) else d_973_v4_)
                (globalState).f16 = d_973_v4_
                d_1209_v146_: _dafny.Map
                d_1209_v146_ = _dafny.Map({(d_969_v0_).f32: d_1105_cf12_})
                d_1210_v147_: D5
                d_1210_v147_ = D5_DC14((d_1206_v145_)[default__.safeIndex(854, (d_1206_v145_).length(0))], d_1105_cf12_, not((d_969_v0_).f32))
                d_1211_v148_: _dafny.Seq
                d_1211_v148_ = _dafny.SeqWithoutIsStrInference([d_1210_v147_, d_1210_v147_])
                d_1212_v149_: _dafny.Map
                d_1212_v149_ = _dafny.Map({(d_969_v0_).f32: (d_977_v8_).f32})
                d_1213_v150_: _dafny.Seq
                d_1213_v150_ = _dafny.SeqWithoutIsStrInference([((default__.fm35(((d_1209_v146_)[(d_969_v0_).f32] if ((d_969_v0_).f32) in (d_1209_v146_) else (d_1206_v145_)[default__.safeIndex(854, (d_1206_v145_).length(0))]), globalState)).set(default__.safeIndex(d_973_v4_, len(default__.fm35(((d_1209_v146_)[(d_969_v0_).f32] if ((d_969_v0_).f32) in (d_1209_v146_) else (d_1206_v145_)[default__.safeIndex(854, (d_1206_v145_).length(0))]), globalState))), d_1210_v147_)) == (d_1211_v148_), (d_1212_v149_) != (_dafny.Map({(self).f27: (d_969_v0_).f32})), (d_977_v8_).f32, (d_977_v8_).f32])
                d_1213_v150_ = (p2) + ((_dafny.SeqWithoutIsStrInference([(d_969_v0_).f32, (d_977_v8_).f32])) + ((d_1041_v53_).cf6))
                d_1214_v151_: _dafny.MultiSet
                d_1214_v151_ = _dafny.MultiSet([d_1195_v134_])
                index162_ = default__.safeIndex(854, (d_1206_v145_).length(0))
                (d_1206_v145_)[index162_] = ((d_1214_v151_)[d_1195_v134_] if (d_1195_v134_) in (d_1214_v151_) else (d_1206_v145_)[default__.safeIndex(854, (d_1206_v145_).length(0))])
                d_1215_v152_: C1
                nw176_ = C1()
                nw176_.ctor__((0) - ((d_1206_v145_)[default__.safeIndex(854, (d_1206_v145_).length(0))]), d_1105_cf12_, (d_969_v0_).f32, (self).f27)
                d_1215_v152_ = nw176_
            elif True:
                d_1105_cf12_ = (0) - (default__.safeDivisionInt(d_1105_cf12_, d_1105_cf12_))
                d_1216_v153_: _dafny.MultiSet
                d_1216_v153_ = _dafny.MultiSet([878, d_973_v4_])
                d_1217_v154_: bool
                out42_: bool
                out42_ = (self).m8(D0_DC0((d_969_v0_).f32), (d_977_v8_).f32, _dafny.MultiSet([(d_969_v0_).f32, (d_969_v0_).f32, p0]), default__.safeModuloInt(d_1105_cf12_, ((d_1216_v153_).set(d_1105_cf12_, default__.abs(d_973_v4_))).cardinality), globalState)
                d_1217_v154_ = out42_
                d_1218_v155_: _dafny.Array
                def lambda65_(d_1219_i16_):
                    return (d_1219_i16_) + (121)

                init36_ = lambda65_
                nw177_ = _dafny.Array(None, 29)
                for i0_36_ in range(nw177_.length(0)):
                    nw177_[i0_36_] = init36_(i0_36_)
                d_1218_v155_ = nw177_
                d_1220_v156_: D3
                d_1220_v156_ = D3_DC10(d_1218_v155_)
                d_1221_v157_: _dafny.Map
                d_1221_v157_ = _dafny.Map({d_1220_v156_: (0) - ((0) - (d_973_v4_))})
                d_1222_v158_: C4
                nw178_ = C4()
                nw178_.ctor__((d_1221_v157_ if (d_969_v0_).f32 else d_1221_v157_), (d_969_v0_).f32, (self).f27)
                d_1222_v158_ = nw178_
                (globalState).f2 = (d_977_v8_).f32
                d_1223_v159_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.Map({}), 21)
                d_1223_v159_ = nw179_
                d_1224_v160_: _dafny.Seq
                d_1224_v160_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnfdufwb"))
                d_1225_v161_: _dafny.Map
                d_1225_v161_ = _dafny.Map({d_1105_cf12_: d_1224_v160_})
                d_1226_v162_: _dafny.Seq
                d_1226_v162_ = _dafny.SeqWithoutIsStrInference([d_1225_v161_, d_1225_v161_, d_1225_v161_, d_1225_v161_])
                index163_ = default__.safeIndex(780, (d_1223_v159_).length(0))
                (d_1223_v159_)[index163_] = ((d_1226_v162_)[default__.safeIndex(d_1105_cf12_, len(d_1226_v162_))]) | (d_1225_v161_)
                d_1227_v163_: D5
                d_1227_v163_ = D5_DC13((d_1106_v80_).set((self).f26, _dafny.SeqWithoutIsStrInference([(d_969_v0_).f32])))
                d_1228_v164_: D5
                d_1228_v164_ = D5_DC15(d_1227_v163_)
                d_1229_v165_: D5
                d_1229_v165_ = D5_DC15(D5_DC15(d_1228_v164_))
                d_1230_v166_: D5
                d_1230_v166_ = D5_DC15(d_1228_v164_)
                d_1231_v167_: _dafny.Array
                nw180_ = _dafny.Array(None, 6)
                nw180_[int(0)] = D5_DC15(d_1227_v163_)
                nw180_[int(1)] = d_1230_v166_
                nw180_[int(2)] = d_1230_v166_
                nw180_[int(3)] = d_1230_v166_
                nw180_[int(4)] = d_1230_v166_
                nw180_[int(5)] = d_1230_v166_
                d_1231_v167_ = nw180_
                index164_ = default__.safeIndex(780, (d_1223_v159_).length(0))
                def iife94_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in ((d_975_v7_).cf13).keys.Elements:
                        d_1232_v168_: int = compr_60_
                        if (d_1232_v168_) in ((d_975_v7_).cf13):
                            coll60_[(d_1232_v168_) + (725)] = d_1224_v160_
                    return _dafny.Map(coll60_)
                rhs178_ = (iife94_()
                ) | ((d_1225_v161_) | (_dafny.Map({(0) - (d_973_v4_): d_1224_v160_})))
                rhs179_ = p0
                rhs180_ = d_1231_v167_
                rhs181_ = not((d_969_v0_).f32)
                lhs162_ = d_1223_v159_
                lhs163_ = default__.safeIndex(780, (d_1223_v159_).length(0))
                lhs164_ = globalState
                lhs162_[lhs163_] = rhs178_
                d_1217_v154_ = rhs179_
                d_1231_v167_ = rhs180_
                lhs164_.f2 = rhs181_
            if (self).f26:
                d_1233_v169_: _dafny.Map
                d_1233_v169_ = _dafny.Map({((0) - (d_1105_cf12_)) + (d_1105_cf12_): (d_977_v8_).f32})
                d_1234_v170_: _dafny.Map
                d_1234_v170_ = _dafny.Map({(d_977_v8_).f32: d_973_v4_})
                d_1233_v169_ = (d_1233_v169_).set(len(d_1234_v170_), (d_969_v0_).f32)
                d_1235_v171_: _dafny.Array
                nw181_ = _dafny.Array(D0.default()(), 16)
                d_1235_v171_ = nw181_
                d_1236_v172_: D0
                d_1236_v172_ = D0_DC0((self).f27)
                index165_ = default__.safeIndex(95, (d_1235_v171_).length(0))
                (d_1235_v171_)[index165_] = d_1236_v172_
                d_1237_v173_: _dafny.Map
                d_1237_v173_ = _dafny.Map({(self).f27: (d_977_v8_).f32})
                index166_ = default__.safeIndex(95, (d_1235_v171_).length(0))
                rhs182_ = D0_DC0(((d_1237_v173_)[(d_977_v8_).f32] if ((d_977_v8_).f32) in (d_1237_v173_) else (d_969_v0_).f32))
                rhs183_ = d_973_v4_
                lhs165_ = d_1235_v171_
                lhs166_ = default__.safeIndex(95, (d_1235_v171_).length(0))
                lhs167_ = globalState
                lhs165_[lhs166_] = rhs182_
                lhs167_.f16 = rhs183_
                d_1238_v174_: _dafny.Seq
                d_1238_v174_ = _dafny.SeqWithoutIsStrInference([d_1105_cf12_, d_973_v4_])
                d_1239_v175_: _dafny.Seq
                d_1239_v175_ = d_1238_v174_
                d_1240_v176_: _dafny.Seq
                d_1240_v176_ = _dafny.SeqWithoutIsStrInference([d_1239_v175_, d_1239_v175_])
                d_1241_v177_: _dafny.Seq
                d_1241_v177_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iywdtfdjg"))
                d_1242_v178_: D6
                d_1242_v178_ = D6_DC17(d_1105_cf12_, d_1105_cf12_, d_973_v4_, len((d_1241_v177_).set(default__.safeIndex(-75, len(d_1241_v177_)), d_1040_v52_)), d_973_v4_)
                d_1243_v179_: _dafny.Map
                d_1243_v179_ = _dafny.Map({d_1240_v176_: d_1242_v178_})
                d_1243_v179_ = d_1243_v179_
                d_1244_v180_: D0
                d_1244_v180_ = D0_DC1((d_969_v0_).f32, (d_977_v8_).f32, len(d_1241_v177_), p0)
                d_1245_v181_: C1
                nw182_ = C1()
                nw182_.ctor__(d_973_v4_, d_973_v4_, (self).f26, (self).f27)
                d_1245_v181_ = nw182_
                d_1246_v182_: _dafny.MultiSet
                d_1246_v182_ = _dafny.MultiSet([d_1245_v181_])
                d_1247_v183_: _dafny.Set
                d_1247_v183_ = _dafny.Set({(self).f26, (d_977_v8_).f32})
                d_1248_v184_: _dafny.Map
                d_1248_v184_ = _dafny.Map({True: d_1040_v52_})
                d_1249_v185_: D5
                d_1249_v185_ = D5_DC14(d_1245_v181_.f37, len(d_1241_v177_), (d_977_v8_).f32)
                d_1250_v186_: _dafny.Array
                nw183_ = _dafny.Array(None, 20)
                nw183_[int(0)] = not(not((-994) != (d_973_v4_)))
                nw183_[int(1)] = (d_977_v8_).f32
                nw183_[int(2)] = (self).fm2(True, (self).fm1(globalState), globalState)
                nw183_[int(3)] = (d_977_v8_).f32
                nw183_[int(4)] = (d_977_v8_).f32
                nw183_[int(5)] = not ((d_969_v0_).fm7(d_1234_v170_, d_1238_v174_, (self).f27, globalState)) or (True)
                nw183_[int(6)] = not ((d_969_v0_).f32) or ((d_1244_v180_).cf1)
                nw183_[int(7)] = (self).f26
                nw183_[int(8)] = (self).f27
                nw183_[int(9)] = (d_977_v8_).f32
                nw183_[int(10)] = ((d_969_v0_).fm7(d_1234_v170_, d_1238_v174_, not(p0), globalState)) and ((self).f26)
                nw183_[int(11)] = (((d_1246_v182_)[d_1245_v181_] if (d_1245_v181_) in (d_1246_v182_) else len(d_1238_v174_))) < (len(d_1247_v183_))
                nw183_[int(12)] = (d_969_v0_).fm6(d_1040_v52_, not((d_969_v0_).fm7(d_1234_v170_, _dafny.SeqWithoutIsStrInference([len(d_1248_v184_)]), (d_969_v0_).f32, globalState)), globalState)
                nw183_[int(13)] = True
                nw183_[int(14)] = not (not((self).f27)) or ((d_969_v0_).f32)
                nw183_[int(15)] = (d_1249_v185_).cf27
                nw183_[int(16)] = ((d_977_v8_).f32 if p0 else p0)
                nw183_[int(17)] = (not(not(p0))) == ((d_969_v0_).f32)
                nw183_[int(18)] = (d_977_v8_).f32
                nw183_[int(19)] = (p2) == (_dafny.SeqWithoutIsStrInference([p0, p0, (d_977_v8_).f32]))
                d_1250_v186_ = nw183_
                index167_ = default__.safeIndex(947, (d_1250_v186_).length(0))
                (d_1250_v186_)[index167_] = (d_977_v8_).f32
                index168_ = default__.safeIndex(947, (d_1250_v186_).length(0))
                (d_1250_v186_)[index168_] = not((d_969_v0_).f32)
                d_1251_v187_: _dafny.Array
                nw184_ = _dafny.Array(D1.default()(), 15)
                d_1251_v187_ = nw184_
                d_1252_v188_: _dafny.Seq
                d_1252_v188_ = _dafny.SeqWithoutIsStrInference([d_1251_v187_])
                (globalState).f22 = (d_1252_v188_) != (d_1252_v188_)
            elif True:
                (globalState).f12 = d_1105_cf12_
                d_1253_v189_: _dafny.Seq
                d_1253_v189_ = _dafny.SeqWithoutIsStrInference([911, len(p2)])
                d_1254_v190_: _dafny.Seq
                d_1254_v190_ = _dafny.SeqWithoutIsStrInference([d_1253_v189_])
                d_1255_v191_: T2
                nw185_ = C2()
                nw185_.ctor__(d_1254_v190_, (d_969_v0_).f32, (p2)[default__.safeIndex(d_973_v4_, len(p2))])
                d_1255_v191_ = nw185_
                d_1256_v192_: D7
                d_1256_v192_ = D7_DC18(d_1255_v191_)
                d_1257_v193_: _dafny.MultiSet
                d_1257_v193_ = _dafny.MultiSet([d_1256_v192_])
                d_1258_v194_: _dafny.Seq
                d_1258_v194_ = _dafny.SeqWithoutIsStrInference([d_1257_v193_])
                d_1259_v195_: _dafny.Seq
                d_1259_v195_ = _dafny.SeqWithoutIsStrInference([d_1256_v192_, d_1256_v192_])
                (globalState).f24 = not(((_dafny.MultiSet(d_1259_v195_)).set(d_1256_v192_, default__.abs(d_1105_cf12_))).issubset((d_1258_v194_)[default__.safeIndex(38, len(d_1258_v194_))]))
                d_1260_v196_: _dafny.Array
                nw186_ = _dafny.Array(False, 5)
                d_1260_v196_ = nw186_
                rhs184_ = d_1260_v196_
                rhs185_ = not ((self).f27) or ((d_977_v8_).f32)
                rhs186_ = (self).fm1(globalState)
                lhs168_ = globalState
                lhs169_ = globalState
                d_1260_v196_ = rhs184_
                lhs168_.f2 = rhs185_
                lhs169_.f16 = rhs186_
                d_1261_v197_: _dafny.Array
                nw187_ = _dafny.Array(int(0), 2)
                d_1261_v197_ = nw187_
                index169_ = default__.safeIndex(581, (d_1261_v197_).length(0))
                (d_1261_v197_)[index169_] = d_1105_cf12_
                index170_ = default__.safeIndex(581, (d_1261_v197_).length(0))
                (d_1261_v197_)[index170_] = d_1105_cf12_
                d_1262_v198_: _dafny.Set
                d_1262_v198_ = _dafny.Set({d_1261_v197_})
                d_1262_v198_ = (_dafny.Set({d_1261_v197_, d_1261_v197_, d_1261_v197_, d_1261_v197_})).intersection(_dafny.Set({d_1261_v197_, d_1261_v197_, d_1261_v197_}))
        elif True:
            d_1263___mcc_h14_ = source18_.cf6
            d_1264_cf6_ = d_1263___mcc_h14_
            d_1265_v199_: _dafny.Map
            d_1265_v199_ = _dafny.Map({(d_977_v8_).f32: p0})
            (globalState).f15 = ((d_1265_v199_)[(self).f27] if ((self).f27) in (d_1265_v199_) else (p2)[default__.safeIndex(450, len(p2))])
            d_1266_v200_: _dafny.Map
            d_1266_v200_ = _dafny.Map({(d_977_v8_).f32: d_973_v4_})
            d_1267_v201_: _dafny.Seq
            d_1267_v201_ = _dafny.SeqWithoutIsStrInference([d_973_v4_])
            d_1268_v202_: _dafny.Map
            d_1268_v202_ = _dafny.Map({(d_969_v0_).fm7(d_1266_v200_, d_1267_v201_, (p2)[default__.safeIndex(d_973_v4_, len(p2))], globalState): _dafny.Map({d_973_v4_: -45})})
            d_1269_v203_: _dafny.Map
            d_1269_v203_ = _dafny.Map({d_973_v4_: -349})
            d_1270_v204_: _dafny.Map
            d_1270_v204_ = _dafny.Map({((d_1269_v203_)[d_973_v4_] if (d_973_v4_) in (d_1269_v203_) else 912): 2})
            d_1271_v205_: _dafny.Seq
            d_1271_v205_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwqodhw"))
            d_1272_v206_: _dafny.Map
            d_1272_v206_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rphevkuq"))): (d_969_v0_).fm6(d_1040_v52_, (d_969_v0_).f32, globalState)})
            d_1273_v207_: _dafny.MultiSet
            d_1273_v207_ = _dafny.MultiSet([d_1267_v201_, d_1267_v201_])
            d_1274_v208_: C3
            nw188_ = C3()
            nw188_.ctor__(d_1273_v207_, (self).f27, (d_969_v0_).f32)
            d_1274_v208_ = nw188_
            d_1275_v209_: _dafny.Array
            nw189_ = _dafny.Array(None, 22)
            nw189_[int(0)] = d_973_v4_
            nw189_[int(1)] = d_973_v4_
            nw189_[int(2)] = len(d_1265_v199_)
            nw189_[int(3)] = d_973_v4_
            nw189_[int(4)] = d_973_v4_
            nw189_[int(5)] = len(d_1272_v206_)
            nw189_[int(6)] = len(d_1271_v205_)
            nw189_[int(7)] = len(_dafny.SeqWithoutIsStrInference([(d_969_v0_).f32]))
            nw189_[int(8)] = d_973_v4_
            nw189_[int(9)] = d_973_v4_
            nw189_[int(10)] = d_973_v4_
            nw189_[int(11)] = d_973_v4_
            nw189_[int(12)] = d_973_v4_
            nw189_[int(13)] = 138
            nw189_[int(14)] = d_973_v4_
            nw189_[int(15)] = d_973_v4_
            nw189_[int(16)] = d_973_v4_
            nw189_[int(17)] = 61
            nw189_[int(18)] = 540
            nw189_[int(19)] = d_973_v4_
            nw189_[int(20)] = d_973_v4_
            nw189_[int(21)] = len(_dafny.SeqWithoutIsStrInference([d_1274_v208_]))
            d_1275_v209_ = nw189_
            d_1276_v210_: D3
            d_1276_v210_ = D3_DC10(d_1275_v209_)
            d_1277_v211_: C4
            nw190_ = C4()
            nw190_.ctor__(_dafny.Map({d_1276_v210_: len(d_1271_v205_)}), (self).f27, (d_977_v8_).f32)
            d_1277_v211_ = nw190_
            d_1278_v212_: _dafny.Map
            d_1278_v212_ = _dafny.Map({d_1277_v211_: (self).f27})
            d_1279_v213_: _dafny.Map
            d_1279_v213_ = _dafny.Map({d_1267_v201_: d_973_v4_})
            d_1280_v214_: _dafny.Array
            nw191_ = _dafny.Array(None, 29)
            nw191_[int(0)] = d_973_v4_
            nw191_[int(1)] = -872
            nw191_[int(2)] = d_973_v4_
            nw191_[int(3)] = len((d_1268_v202_) | (_dafny.Map({(d_969_v0_).f32: d_1269_v203_})))
            nw191_[int(4)] = d_973_v4_
            nw191_[int(5)] = -557
            nw191_[int(6)] = default__.safeDivisionInt(-248, d_973_v4_)
            nw191_[int(7)] = d_973_v4_
            nw191_[int(8)] = d_973_v4_
            nw191_[int(9)] = (0) - (len(((d_1271_v205_).set(default__.safeIndex(len(d_1265_v199_), len(d_1271_v205_)), d_1040_v52_) if (self).f26 else d_1271_v205_)))
            nw191_[int(10)] = ((0) - (d_973_v4_)) - (d_973_v4_)
            nw191_[int(11)] = (354) - ((d_1267_v201_)[default__.safeIndex(d_973_v4_, len(d_1267_v201_))])
            nw191_[int(12)] = (0) - (default__.safeDivisionInt(d_973_v4_, len(d_1272_v206_)))
            nw191_[int(13)] = d_973_v4_
            nw191_[int(14)] = len((d_1271_v205_).set(default__.safeIndex(d_973_v4_, len(d_1271_v205_)), d_1040_v52_))
            nw191_[int(15)] = 850
            nw191_[int(16)] = 942
            nw191_[int(17)] = d_973_v4_
            nw191_[int(18)] = d_973_v4_
            nw191_[int(19)] = (0) - (len((d_1278_v212_) | (d_1278_v212_)))
            nw191_[int(20)] = d_973_v4_
            nw191_[int(21)] = len(d_1279_v213_)
            nw191_[int(22)] = (196 if (d_977_v8_).f32 else 664)
            nw191_[int(23)] = default__.safeModuloInt(d_973_v4_, d_973_v4_)
            nw191_[int(24)] = d_973_v4_
            nw191_[int(25)] = d_973_v4_
            nw191_[int(26)] = d_973_v4_
            nw191_[int(27)] = d_973_v4_
            nw191_[int(28)] = 480
            d_1280_v214_ = nw191_
            d_1280_v214_ = d_1275_v209_
            d_1278_v212_ = d_1278_v212_
            d_1281_v215_: _dafny.Array
            def lambda66_(d_1282_i17_):
                return _dafny.SeqWithoutIsStrInference([(self).f26, (self).f27])

            init37_ = lambda66_
            nw192_ = _dafny.Array(None, 8)
            for i0_37_ in range(nw192_.length(0)):
                nw192_[i0_37_] = init37_(i0_37_)
            d_1281_v215_ = nw192_
            index171_ = default__.safeIndex(769, (d_1281_v215_).length(0))
            (d_1281_v215_)[index171_] = d_1264_cf6_
            index172_ = default__.safeIndex(769, (d_1281_v215_).length(0))
            (d_1281_v215_)[index172_] = (p2) + (d_1264_cf6_)

    def m7(self, p0, p1, p2, p3, globalState):
        d_1283_i0_: int
        d_1283_i0_ = 0
        with _dafny.label("6"):
            while (self).f27:
                with _dafny.c_label("6"):
                    if (d_1283_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1283_i0_ = (d_1283_i0_) + (1)
                    d_1284_v0_: C1
                    nw193_ = C1()
                    nw193_.ctor__(p2, p2, (self).f26, p1)
                    d_1284_v0_ = nw193_
                    d_1285_v1_: _dafny.Array
                    nw194_ = _dafny.Array(None, 3)
                    nw194_[int(0)] = d_1284_v0_
                    nw194_[int(1)] = d_1284_v0_
                    nw194_[int(2)] = d_1284_v0_
                    d_1285_v1_ = nw194_
                    d_1286_v2_: _dafny.Seq
                    d_1286_v2_ = _dafny.SeqWithoutIsStrInference([d_1285_v1_])
                    (globalState).f2 = (d_1286_v2_) <= (d_1286_v2_)
                    d_1287_v3_: _dafny.Seq
                    d_1287_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jn"))
                    d_1288_v4_: _dafny.Set
                    d_1288_v4_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1289_i1_ in range(default__.abs(825))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yknlqwi"))), d_1287_v3_})
                    (d_1284_v0_).f37 = len(d_1288_v4_)
                    d_1290_v5_: _dafny.Seq
                    d_1290_v5_ = _dafny.SeqWithoutIsStrInference([d_1287_v3_])
                    d_1290_v5_ = (((d_1290_v5_) + (_dafny.SeqWithoutIsStrInference([d_1287_v3_ for d_1291_i2_ in range(default__.abs(330))]))) + (d_1290_v5_)).set(default__.safeIndex(d_1284_v0_.f37, len(((d_1290_v5_) + (_dafny.SeqWithoutIsStrInference([d_1287_v3_ for d_1292_i2_ in range(default__.abs(330))]))) + (d_1290_v5_))), d_1287_v3_)
                    d_1293_v6_: _dafny.Array
                    nw195_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_1293_v6_ = nw195_
                    d_1294_v7_: _dafny.Array
                    nw196_ = _dafny.Array(_dafny.Map({}), 26)
                    d_1294_v7_ = nw196_
                    index173_ = default__.safeIndex(961, (d_1293_v6_).length(0))
                    (d_1293_v6_)[index173_] = d_1294_v7_
                    index174_ = default__.safeIndex(961, (d_1293_v6_).length(0))
                    (d_1293_v6_)[index174_] = d_1294_v7_
                    pass
            pass
        d_1295_v8_: _dafny.Array
        def lambda67_(d_1296_i3_):
            return (self).f26

        init38_ = lambda67_
        nw197_ = _dafny.Array(None, 6)
        for i0_38_ in range(nw197_.length(0)):
            nw197_[i0_38_] = init38_(i0_38_)
        d_1295_v8_ = nw197_
        index175_ = default__.safeIndex(477, (d_1295_v8_).length(0))
        (d_1295_v8_)[index175_] = (self).f26
        d_1297_v9_: _dafny.Array
        nw198_ = _dafny.Array(int(0), 4)
        d_1297_v9_ = nw198_
        index176_ = default__.safeIndex(234, (d_1295_v8_).length(0))
        (d_1295_v8_)[index176_] = (self).f26
        d_1298_v10_: _dafny.Seq
        d_1298_v10_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1299_v11_: _dafny.MultiSet
        d_1299_v11_ = _dafny.MultiSet([d_1298_v10_, (_dafny.SeqWithoutIsStrInference([p2 for d_1300_i4_ in range(default__.abs(-396))])).set(default__.safeIndex(p2, len(_dafny.SeqWithoutIsStrInference([p2 for d_1301_i4_ in range(default__.abs(-396))]))), p2)])
        index177_ = default__.safeIndex(477, (d_1295_v8_).length(0))
        index178_ = default__.safeIndex(234, (d_1295_v8_).length(0))
        rhs187_ = False
        rhs188_ = (self).f27
        rhs189_ = d_1297_v9_
        rhs190_ = p1
        rhs191_ = ((d_1299_v11_) | (d_1299_v11_)).issubset(d_1299_v11_)
        lhs170_ = globalState
        lhs171_ = d_1295_v8_
        lhs172_ = default__.safeIndex(477, (d_1295_v8_).length(0))
        lhs173_ = d_1295_v8_
        lhs174_ = default__.safeIndex(234, (d_1295_v8_).length(0))
        lhs175_ = globalState
        lhs170_.f15 = rhs187_
        lhs171_[lhs172_] = rhs188_
        d_1297_v9_ = rhs189_
        lhs173_[lhs174_] = rhs190_
        lhs175_.f2 = rhs191_
        hi11_ = p2
        for d_1302_i5_ in range(default__.safeDivisionInt((0) - ((self).fm1(globalState)), p2), hi11_):
            (globalState).f12 = p2
            d_1303_v12_: _dafny.Map
            d_1303_v12_ = _dafny.Map({(self).fm2((d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))], (self).fm1(globalState), globalState): p2})
            (globalState).f16 = ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1304_i6_ in range(default__.abs(463))]))) + (len(d_1303_v12_)) if (d_1302_i5_) <= (p2) else 887)
            d_1305_v13_: _dafny.Seq
            d_1305_v13_ = _dafny.SeqWithoutIsStrInference([not(p1), (self).f26])
            (globalState).f15 = ((d_1305_v13_)[default__.safeIndex(p2, len(d_1305_v13_))] if not(not(True)) else p0)
            if p1:
                d_1306_v14_: _dafny.Map
                d_1306_v14_ = _dafny.Map({p0: p0})
                d_1307_v15_: _dafny.Map
                d_1307_v15_ = _dafny.Map({p2: len(d_1306_v14_)})
                d_1308_v16_: D2
                d_1308_v16_ = D2_DC6(d_1307_v15_)
                d_1309_v17_: _dafny.Seq
                d_1309_v17_ = _dafny.SeqWithoutIsStrInference([d_1308_v16_, D2_DC6(d_1307_v15_)])
                d_1310_v18_: _dafny.MultiSet
                d_1310_v18_ = _dafny.MultiSet([d_1309_v17_, d_1309_v17_, d_1309_v17_, d_1309_v17_, d_1309_v17_])
                index179_ = default__.safeIndex(477, (d_1295_v8_).length(0))
                (d_1295_v8_)[index179_] = (d_1310_v18_).issubset((_dafny.MultiSet([(d_1309_v17_).set(default__.safeIndex(700, len(d_1309_v17_)), d_1308_v16_)]) if (self).f26 else default__.fm36(p2, (d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))], not((self).fm2(p1, p2, globalState)), (self).f26, globalState)))
                d_1311_v19_: D3
                d_1311_v19_ = D3_DC10(d_1297_v9_)
                d_1312_v20_: _dafny.Map
                d_1312_v20_ = _dafny.Map({p2: d_1295_v8_})
                d_1313_v21_: _dafny.Map
                d_1313_v21_ = _dafny.Map({d_1311_v19_: len(d_1312_v20_)})
                d_1314_v22_: C4
                nw199_ = C4()
                nw199_.ctor__(d_1313_v21_, p1, p1)
                d_1314_v22_ = nw199_
                d_1315_v23_: _dafny.Seq
                d_1315_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gegvu"))
                (globalState).f6 = (d_1315_v23_) + (d_1315_v23_)
                d_1316_v24_: _dafny.Seq
                d_1316_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p2 for d_1317_i7_ in range(default__.abs(117))]), _dafny.SeqWithoutIsStrInference([p2]), _dafny.SeqWithoutIsStrInference([(0) - (len(d_1307_v15_)), d_1302_i5_]), d_1298_v10_, _dafny.SeqWithoutIsStrInference([d_1302_i5_, d_1302_i5_])])
                d_1318_v25_: C2
                nw200_ = C2()
                nw200_.ctor__(d_1316_v24_, (d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))], (self).f27)
                d_1318_v25_ = nw200_
                d_1319_v26_: _dafny.Map
                d_1319_v26_ = _dafny.Map({d_1298_v10_: d_1318_v25_})
                d_1320_v27_: _dafny.Array
                nw201_ = _dafny.Array(None, 3)
                nw201_[int(0)] = (d_1319_v26_) | (d_1319_v26_)
                nw201_[int(1)] = d_1319_v26_
                nw201_[int(2)] = d_1319_v26_
                d_1320_v27_ = nw201_
                index180_ = default__.safeIndex(450, (d_1320_v27_).length(0))
                (d_1320_v27_)[index180_] = d_1319_v26_
                index181_ = default__.safeIndex(450, (d_1320_v27_).length(0))
                (d_1320_v27_)[index181_] = d_1319_v26_
                d_1321_v28_: D0
                d_1321_v28_ = D0_DC0((self).f26)
                index182_ = default__.safeIndex(821, (d_1297_v9_).length(0))
                (d_1297_v9_)[index182_] = ((d_1318_v25_).fm4(d_1321_v28_, d_1302_i5_, p2, globalState)) - (d_1302_i5_)
                index183_ = default__.safeIndex(821, (d_1297_v9_).length(0))
                rhs192_ = (0) - ((606 if (d_1305_v13_)[default__.safeIndex(d_1302_i5_, len(d_1305_v13_))] else (d_1302_i5_) - (p2)))
                rhs193_ = d_1302_i5_
                lhs176_ = globalState
                lhs177_ = d_1297_v9_
                lhs178_ = default__.safeIndex(821, (d_1297_v9_).length(0))
                lhs176_.f12 = rhs192_
                lhs177_[lhs178_] = rhs193_
            elif True:
                d_1322_v29_: _dafny.Array
                nw202_ = _dafny.Array(None, 6)
                nw202_[int(0)] = d_1297_v9_
                nw202_[int(1)] = d_1297_v9_
                nw202_[int(2)] = d_1297_v9_
                nw202_[int(3)] = d_1297_v9_
                nw202_[int(4)] = d_1297_v9_
                nw202_[int(5)] = d_1297_v9_
                d_1322_v29_ = nw202_
                index184_ = default__.safeIndex(699, (d_1322_v29_).length(0))
                (d_1322_v29_)[index184_] = d_1297_v9_
                index185_ = default__.safeIndex(699, (d_1322_v29_).length(0))
                (d_1322_v29_)[index185_] = d_1297_v9_
                d_1323_v30_: _dafny.Seq
                d_1323_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctfdse"))
                d_1324_v31_: _dafny.Seq
                d_1324_v31_ = _dafny.SeqWithoutIsStrInference([(d_1322_v29_)[default__.safeIndex(699, (d_1322_v29_).length(0))], d_1297_v9_, d_1297_v9_])
                rhs194_ = (d_1322_v29_)[default__.safeIndex(699, (d_1322_v29_).length(0))]
                rhs195_ = d_1323_v30_
                rhs196_ = 812
                rhs197_ = (d_1324_v31_)[default__.safeIndex(p2, len(d_1324_v31_))]
                lhs179_ = globalState
                lhs180_ = globalState
                d_1297_v9_ = rhs194_
                lhs179_.f6 = rhs195_
                lhs180_.f16 = rhs196_
                d_1297_v9_ = rhs197_
                (globalState).f16 = (self).fm1(globalState)
                d_1325_v32_: _dafny.Array
                nw203_ = _dafny.Array(_dafny.Seq({}), 4)
                d_1325_v32_ = nw203_
                d_1326_v33_: _dafny.Array
                nw204_ = _dafny.Array(D1.default()(), 4)
                d_1326_v33_ = nw204_
                d_1327_v34_: _dafny.Seq
                d_1327_v34_ = _dafny.SeqWithoutIsStrInference([d_1326_v33_, d_1326_v33_])
                index186_ = default__.safeIndex(349, (d_1325_v32_).length(0))
                (d_1325_v32_)[index186_] = d_1327_v34_
                d_1328_v35_: _dafny.MultiSet
                d_1328_v35_ = _dafny.MultiSet([p3, p3])
                index187_ = default__.safeIndex(349, (d_1325_v32_).length(0))
                rhs198_ = (d_1327_v34_) + (d_1327_v34_)
                rhs199_ = d_1302_i5_
                rhs200_ = (d_1328_v35_).issubset(d_1328_v35_)
                lhs181_ = d_1325_v32_
                lhs182_ = default__.safeIndex(349, (d_1325_v32_).length(0))
                lhs183_ = globalState
                lhs184_ = globalState
                lhs181_[lhs182_] = rhs198_
                lhs183_.f16 = rhs199_
                lhs184_.f15 = rhs200_
                d_1329_v36_: C1
                nw205_ = C1()
                nw205_.ctor__(p2, d_1302_i5_, ((0) - (-125)) == ((p3).cardinality), (d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))])
                d_1329_v36_ = nw205_
        index188_ = default__.safeIndex(477, (d_1295_v8_).length(0))
        (d_1295_v8_)[index188_] = (self).f26
        d_1330_v37_: _dafny.Map
        d_1330_v37_ = _dafny.Map({946: True})
        (globalState).f24 = (len(d_1330_v37_)) == (p2)
        if (self).f26:
            d_1331_v38_: _dafny.Map
            d_1331_v38_ = _dafny.Map({(p3).cardinality: p2})
            d_1332_v39_: _dafny.MultiSet
            d_1332_v39_ = _dafny.MultiSet([default__.safeDivisionInt(p2, len(d_1331_v38_)), p2])
            d_1332_v39_ = (d_1332_v39_) - ((d_1332_v39_) - (d_1332_v39_))
            (globalState).f16 = len(d_1298_v10_)
            d_1333_v40_: _dafny.Seq
            d_1333_v40_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_1298_v10_ = _dafny.SeqWithoutIsStrInference([len(d_1333_v40_), p2, (599) - (p2)])
            (globalState).f22 = p1
            d_1334_v41_: _dafny.Array
            nw206_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_1334_v41_ = nw206_
            d_1334_v41_ = d_1334_v41_
        elif True:
            index189_ = default__.safeIndex(478, (d_1297_v9_).length(0))
            (d_1297_v9_)[index189_] = p2
            index190_ = default__.safeIndex(478, (d_1297_v9_).length(0))
            (d_1297_v9_)[index190_] = p2
            d_1335_v42_: D0
            d_1335_v42_ = D0_DC0(p0)
            d_1336_v43_: _dafny.Seq
            d_1336_v43_ = _dafny.SeqWithoutIsStrInference([True])
            d_1337_v44_: _dafny.MultiSet
            d_1337_v44_ = _dafny.MultiSet([p1])
            d_1338_v45_: bool
            out43_: bool
            out43_ = (self).m8(d_1335_v42_, (d_1336_v43_)[default__.safeIndex((self).fm1(globalState), len(d_1336_v43_))], d_1337_v44_, (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], globalState)
            d_1338_v45_ = out43_
            d_1339_v46_: _dafny.MultiSet
            d_1339_v46_ = _dafny.MultiSet([d_1295_v8_, d_1295_v8_])
            d_1340_v47_: _dafny.Seq
            d_1340_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhfpskf"))
            d_1341_v48_: _dafny.Map
            d_1341_v48_ = _dafny.Map({((d_1339_v46_).intersection(d_1339_v46_)).cardinality: (d_1340_v47_ if (self).f27 else d_1340_v47_)})
            d_1341_v48_ = (d_1341_v48_).set((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))) + (default__.fm22(p2, (self).fm2(p0, (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], globalState), globalState)))
            d_1342_v49_: _dafny.Map
            d_1342_v49_ = _dafny.Map({p1: (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]})
            d_1343_v50_: D7
            d_1343_v50_ = D7_DC19((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], d_1337_v44_, (self).f27, d_1342_v49_)
            if (d_1343_v50_).cf38:
                index191_ = default__.safeIndex(478, (d_1297_v9_).length(0))
                (d_1297_v9_)[index191_] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                (globalState).f12 = p2
                d_1344_v51_: _dafny.Map
                d_1344_v51_ = _dafny.Map({(d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]: _dafny.SeqWithoutIsStrInference([(d_1343_v50_).cf38])})
                d_1345_v52_: str
                d_1345_v52_ = _dafny.CodePoint('h')
                d_1344_v51_ = (d_1344_v51_).set((self).fm1(globalState), default__.fm11((d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))], (d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))], d_1345_v52_, globalState))
                d_1346_v53_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_1346_v53_ = nw207_
                d_1347_v54_: _dafny.Array
                nw208_ = _dafny.Array(None, 25)
                nw208_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbnprm"))
                nw208_[int(1)] = d_1340_v47_
                nw208_[int(2)] = d_1340_v47_
                nw208_[int(3)] = d_1340_v47_
                nw208_[int(4)] = d_1340_v47_
                nw208_[int(5)] = d_1340_v47_
                nw208_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                nw208_[int(7)] = d_1340_v47_
                nw208_[int(8)] = d_1340_v47_
                nw208_[int(9)] = d_1340_v47_
                nw208_[int(10)] = d_1340_v47_
                nw208_[int(11)] = default__.fm22(p2, d_1338_v45_, globalState)
                nw208_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hluwry"))
                nw208_[int(13)] = d_1340_v47_
                nw208_[int(14)] = d_1340_v47_
                nw208_[int(15)] = d_1340_v47_
                nw208_[int(16)] = d_1340_v47_
                nw208_[int(17)] = d_1340_v47_
                nw208_[int(18)] = d_1340_v47_
                nw208_[int(19)] = d_1340_v47_
                nw208_[int(20)] = d_1340_v47_
                nw208_[int(21)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkxwgoox"))).set(default__.safeIndex((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkxwgoox")))), d_1345_v52_)
                nw208_[int(22)] = d_1340_v47_
                nw208_[int(23)] = _dafny.SeqWithoutIsStrInference([d_1345_v52_ for d_1348_i8_ in range(default__.abs(773))])
                nw208_[int(24)] = d_1340_v47_
                d_1347_v54_ = nw208_
                d_1349_v55_: _dafny.Seq
                d_1349_v55_ = _dafny.SeqWithoutIsStrInference([d_1347_v54_, d_1347_v54_])
                index192_ = default__.safeIndex(192, (d_1346_v53_).length(0))
                (d_1346_v53_)[index192_] = (d_1347_v54_ if (self).fm2(True, (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], globalState) else (d_1349_v55_)[default__.safeIndex(p2, len(d_1349_v55_))])
                index193_ = default__.safeIndex(192, (d_1346_v53_).length(0))
                (d_1346_v53_)[index193_] = d_1347_v54_
                d_1350_v56_: _dafny.Map
                d_1350_v56_ = _dafny.Map({(self).f26: _dafny.SeqWithoutIsStrInference([(self).f27, d_1338_v45_])})
                d_1351_v57_: D5
                d_1351_v57_ = D5_DC13(d_1350_v56_)
                d_1351_v57_ = d_1351_v57_
            elif True:
                d_1352_v58_: str
                d_1352_v58_ = _dafny.CodePoint('a')
                d_1353_v59_: _dafny.Array
                nw209_ = _dafny.Array(None, 25)
                nw209_[int(0)] = p2
                nw209_[int(1)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(2)] = p2
                nw209_[int(3)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(4)] = p2
                nw209_[int(5)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(6)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(7)] = (0) - ((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))])
                nw209_[int(8)] = len(d_1298_v10_)
                nw209_[int(9)] = p2
                nw209_[int(10)] = p2
                nw209_[int(11)] = p2
                nw209_[int(12)] = 606
                nw209_[int(13)] = p2
                nw209_[int(14)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(15)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(16)] = p2
                nw209_[int(17)] = len(_dafny.SeqWithoutIsStrInference([d_1352_v58_]))
                nw209_[int(18)] = p2
                nw209_[int(19)] = p2
                nw209_[int(20)] = (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]
                nw209_[int(21)] = p2
                nw209_[int(22)] = len(d_1340_v47_)
                nw209_[int(23)] = p2
                nw209_[int(24)] = p2
                d_1353_v59_ = nw209_
                d_1354_v60_: _dafny.Map
                d_1354_v60_ = _dafny.Map({d_1353_v59_: len(d_1340_v47_)})
                d_1355_v61_: _dafny.Map
                d_1355_v61_ = _dafny.Map({d_1337_v44_: _dafny.Set({(self).f27})})
                d_1356_v62_: D5
                d_1356_v62_ = D5_DC14(len(((d_1355_v61_)[d_1337_v44_] if (d_1337_v44_) in (d_1355_v61_) else _dafny.Set({(self).f26, (self).f26}))), p2, (d_1336_v43_)[default__.safeIndex((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], len(d_1336_v43_))])
                index194_ = default__.safeIndex(478, (d_1297_v9_).length(0))
                rhs201_ = (len((d_1354_v60_) | (_dafny.Map({d_1353_v59_: (d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]}))) if (self).f27 else p2)
                rhs202_ = (d_1356_v62_).cf26
                rhs203_ = (d_1338_v45_ if not((p2) <= (p2)) else (d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))])
                lhs185_ = globalState
                lhs186_ = d_1297_v9_
                lhs187_ = default__.safeIndex(478, (d_1297_v9_).length(0))
                lhs188_ = globalState
                lhs185_.f12 = rhs201_
                lhs186_[lhs187_] = rhs202_
                lhs188_.f24 = rhs203_
                d_1357_v63_: C0
                nw210_ = C0()
                nw210_.ctor__((self).f26)
                d_1357_v63_ = nw210_
                d_1358_v64_: _dafny.Set
                d_1358_v64_ = _dafny.Set({(self).f26})
                d_1359_v65_: _dafny.Map
                d_1359_v65_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([D3_DC11(d_1352_v58_) for d_1360_i9_ in range(default__.abs(262))]): d_1357_v63_})
                d_1361_v66_: D3
                d_1361_v66_ = D3_DC11(d_1352_v58_)
                rhs204_ = ((d_1357_v63_).f32) in (((d_1358_v64_) if (self).f27 else d_1358_v64_))
                rhs205_ = ((d_1359_v65_)[_dafny.SeqWithoutIsStrInference([d_1361_v66_, D3_DC11(d_1352_v58_), d_1361_v66_, d_1361_v66_])] if (_dafny.SeqWithoutIsStrInference([d_1361_v66_, D3_DC11(d_1352_v58_), d_1361_v66_, d_1361_v66_])) in (d_1359_v65_) else d_1357_v63_)
                lhs189_ = globalState
                lhs189_.f2 = rhs204_
                d_1357_v63_ = rhs205_
                (globalState).f6 = d_1340_v47_
                d_1362_v67_: _dafny.Array
                nw211_ = _dafny.Array(None, 18)
                nw211_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxwll"))) + (d_1340_v47_)
                nw211_[int(1)] = d_1340_v47_
                nw211_[int(2)] = d_1340_v47_
                nw211_[int(3)] = (d_1340_v47_).set(default__.safeIndex((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], len(d_1340_v47_)), d_1352_v58_)
                nw211_[int(4)] = d_1340_v47_
                nw211_[int(5)] = d_1340_v47_
                nw211_[int(6)] = (d_1340_v47_) + (d_1340_v47_)
                nw211_[int(7)] = d_1340_v47_
                nw211_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lawmo"))
                nw211_[int(9)] = d_1340_v47_
                nw211_[int(10)] = default__.fm22((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], p0, globalState)
                nw211_[int(11)] = _dafny.SeqWithoutIsStrInference([d_1352_v58_ for d_1363_i10_ in range(default__.abs(716))])
                nw211_[int(12)] = d_1340_v47_
                nw211_[int(13)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1364_i11_ in range(default__.abs(715))])
                nw211_[int(14)] = d_1340_v47_
                nw211_[int(15)] = d_1340_v47_
                nw211_[int(16)] = d_1340_v47_
                nw211_[int(17)] = d_1340_v47_
                d_1362_v67_ = nw211_
                def lambda68_(d_1365_v58_):
                    def lambda69_(d_1366_i12_):
                        return _dafny.SeqWithoutIsStrInference([d_1365_v58_ for d_1367_i13_ in range(default__.abs(734))])

                    return lambda69_

                init39_ = lambda68_(d_1352_v58_)
                nw212_ = _dafny.Array(None, 3)
                for i0_39_ in range(nw212_.length(0)):
                    nw212_[i0_39_] = init39_(i0_39_)
                d_1362_v67_ = nw212_
                d_1368_v68_: _dafny.Array
                d_1368_v68_ = d_1295_v8_
                d_1369_v69_: _dafny.Seq
                d_1369_v69_ = _dafny.SeqWithoutIsStrInference([d_1295_v8_, (d_1368_v68_), d_1295_v8_, d_1295_v8_])
                d_1370_v70_: _dafny.Map
                d_1370_v70_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hnkb")): _dafny.SeqWithoutIsStrInference([d_1295_v8_])})
                d_1371_v71_: _dafny.Array
                nw213_ = _dafny.Array(None, 5)
                nw213_[int(0)] = d_1369_v69_
                nw213_[int(1)] = d_1369_v69_
                nw213_[int(2)] = (d_1369_v69_) + (d_1369_v69_)
                nw213_[int(3)] = (d_1369_v69_) + (d_1369_v69_)
                nw213_[int(4)] = (d_1369_v69_) + ((((d_1370_v70_)[default__.fm22((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], (d_1357_v63_).fm6(d_1352_v58_, (d_1357_v63_).f32, globalState), globalState)] if (default__.fm22((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], (d_1357_v63_).fm6(d_1352_v58_, (d_1357_v63_).f32, globalState), globalState)) in (d_1370_v70_) else d_1369_v69_)).set(default__.safeIndex(p2, len(((d_1370_v70_)[default__.fm22((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], (d_1357_v63_).fm6(d_1352_v58_, (d_1357_v63_).f32, globalState), globalState)] if (default__.fm22((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))], (d_1357_v63_).fm6(d_1352_v58_, (d_1357_v63_).f32, globalState), globalState)) in (d_1370_v70_) else d_1369_v69_))), d_1295_v8_))
                d_1371_v71_ = nw213_
                index195_ = default__.safeIndex(573, (d_1371_v71_).length(0))
                (d_1371_v71_)[index195_] = d_1369_v69_
                index196_ = default__.safeIndex(573, (d_1371_v71_).length(0))
                (d_1371_v71_)[index196_] = (d_1369_v69_) + ((d_1369_v69_) + (d_1369_v69_))
            if p1:
                d_1372_v72_: _dafny.Map
                d_1372_v72_ = _dafny.Map({d_1340_v47_: d_1295_v8_})
                d_1372_v72_ = (d_1372_v72_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1373_i14_ in range(default__.abs(162))]), d_1295_v8_)
                d_1374_v73_: C3
                nw214_ = C3()
                nw214_.ctor__(d_1299_v11_, (p1) == (d_1338_v45_), not (d_1338_v45_) or (d_1338_v45_))
                d_1374_v73_ = nw214_
                d_1375_v74_: bool
                out44_: bool
                out44_ = (self).m8(d_1335_v42_, (self).f26, _dafny.MultiSet([(d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))]]), default__.safeDivisionInt(553, len(d_1340_v47_)), globalState)
                d_1375_v74_ = out44_
                (globalState).f12 = default__.safeDivisionInt((0) - ((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]), p2)
                d_1376_v75_: str
                d_1376_v75_ = _dafny.CodePoint('h')
                (globalState).f19 = d_1376_v75_
            elif True:
                d_1377_v76_: _dafny.Map
                d_1377_v76_ = _dafny.Map({d_1330_v37_: p1})
                d_1378_v77_: C1
                nw215_ = C1()
                nw215_.ctor__(len(d_1377_v76_), p2, p0, False)
                d_1378_v77_ = nw215_
                d_1379_v78_: _dafny.Map
                d_1379_v78_ = _dafny.Map({d_1378_v77_: p3})
                d_1380_v79_: _dafny.Map
                d_1380_v79_ = _dafny.Map({p2: ((d_1379_v78_)[d_1378_v77_] if (d_1378_v77_) in (d_1379_v78_) else _dafny.MultiSet([(d_1378_v77_).f36]))})
                d_1381_v80_: C0
                nw216_ = C0()
                nw216_.ctor__(True)
                d_1381_v80_ = nw216_
                d_1382_v81_: _dafny.Seq
                d_1382_v81_ = _dafny.SeqWithoutIsStrInference([d_1381_v80_, d_1381_v80_])
                d_1383_v82_: _dafny.Map
                d_1383_v82_ = _dafny.Map({d_1380_v79_: (d_1382_v81_) + (d_1382_v81_)})
                d_1383_v82_ = (d_1383_v82_).set(d_1380_v79_, d_1382_v81_)
                d_1298_v10_ = (d_1298_v10_) + (d_1298_v10_)
                d_1384_v83_: _dafny.Array
                nw217_ = _dafny.Array(int(0), 29)
                d_1384_v83_ = nw217_
                d_1385_v84_: D6
                d_1385_v84_ = D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drgpsi")))
                d_1386_v85_: str
                d_1386_v85_ = _dafny.CodePoint('c')
                d_1387_v86_: _dafny.Map
                d_1387_v86_ = _dafny.Map({d_1385_v84_: d_1386_v85_})
                d_1388_v87_: _dafny.Set
                d_1388_v87_ = _dafny.Set({d_1387_v86_})
                d_1389_v89_: D3
                d_1389_v89_ = D3_DC10(d_1384_v83_)
                d_1390_v90_: C4
                nw218_ = C4()
                nw218_.ctor__(_dafny.Map({d_1389_v89_: (0) - ((d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))])}), p0, True)
                d_1390_v90_ = nw218_
                d_1391_v91_: _dafny.Seq
                d_1391_v91_ = _dafny.SeqWithoutIsStrInference([d_1390_v90_, d_1390_v90_, d_1390_v90_])
                d_1392_v92_: _dafny.Map
                d_1392_v92_ = _dafny.Map({len(d_1391_v91_): d_1378_v77_.f37})
                d_1393_v93_: D2
                def iife95_():
                    coll61_ = _dafny.Map()
                    compr_61_: int
                    for compr_61_ in (d_1392_v92_).keys.Elements:
                        d_1394_v88_: int = compr_61_
                        if (d_1394_v88_) in (d_1392_v92_):
                            coll61_[(d_1394_v88_) * (-800)] = p2
                    return _dafny.Map(coll61_)
                d_1393_v93_ = D2_DC6(iife95_()
)
                d_1395_v94_: _dafny.Map
                d_1395_v94_ = _dafny.Map({d_1298_v10_: (_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({d_1378_v77_.f37: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]])]))})) for d_1396_i15_ in range(default__.abs(719))])).set(default__.safeIndex(len(d_1340_v47_), len(_dafny.SeqWithoutIsStrInference([D2_DC6(_dafny.Map({d_1378_v77_.f37: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1297_v9_)[default__.safeIndex(478, (d_1297_v9_).length(0))]])]))})) for d_1397_i15_ in range(default__.abs(719))]))), d_1393_v93_)})
                d_1398_v95_: D10
                d_1398_v95_ = D10_DC23(d_1395_v94_)
                rhs206_ = d_1378_v77_.f37
                rhs207_ = d_1384_v83_
                rhs208_ = d_1295_v8_
                rhs209_ = _dafny.Set({(default__.fm37(d_1338_v45_, (d_1378_v77_).fm1(globalState), p2, d_1398_v95_, globalState)) | (d_1387_v86_)})
                rhs210_ = 569
                lhs190_ = globalState
                lhs191_ = globalState
                lhs190_.f16 = rhs206_
                d_1384_v83_ = rhs207_
                d_1295_v8_ = rhs208_
                d_1388_v87_ = rhs209_
                lhs191_.f16 = rhs210_
                index197_ = default__.safeIndex(477, (d_1295_v8_).length(0))
                (d_1295_v8_)[index197_] = ((d_1390_v90_).fm1(globalState)) <= ((len(_dafny.SeqWithoutIsStrInference([(d_1295_v8_)[default__.safeIndex(477, (d_1295_v8_).length(0))]]))) - ((d_1378_v77_).f36))
                d_1399_v96_: _dafny.Set
                d_1399_v96_ = _dafny.Set({len(d_1340_v47_)})
                index198_ = default__.safeIndex(477, (d_1295_v8_).length(0))
                (d_1295_v8_)[index198_] = ((len(d_1399_v96_)) - (p2)) < (p2)

    def m8(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_1400_v0_: _dafny.Array
        nw219_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
        d_1400_v0_ = nw219_
        d_1401_v1_: str
        d_1401_v1_ = _dafny.CodePoint('h')
        index199_ = default__.safeIndex(371, (d_1400_v0_).length(0))
        (d_1400_v0_)[index199_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsnm"))).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsnm")))), d_1401_v1_)
        d_1402_v2_: _dafny.Seq
        d_1402_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxuu"))
        index200_ = default__.safeIndex(371, (d_1400_v0_).length(0))
        (d_1400_v0_)[index200_] = d_1402_v2_
        d_1403_v3_: _dafny.Seq
        d_1403_v3_ = _dafny.SeqWithoutIsStrInference([(self).fm2((self).f26, p3, globalState), p1])
        (globalState).f12 = len((_dafny.SeqWithoutIsStrInference([(self).f26])) + (d_1403_v3_))
        (globalState).f16 = (0) - ((0) - (p3))
        d_1404_v4_: _dafny.Map
        d_1404_v4_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p3, p3])): 427})
        d_1405_v5_: D2
        d_1405_v5_ = D2_DC6(d_1404_v4_)
        d_1406_v7_: _dafny.Set
        d_1406_v7_ = _dafny.Set({p3})
        pat_let_tv24_ = d_1406_v7_
        pat_let_tv25_ = d_1406_v7_
        pat_let_tv26_ = p3
        d_1407_v8_: _dafny.Array
        nw220_ = _dafny.Array(None, 18)
        nw220_[int(0)] = d_1405_v5_
        nw220_[int(1)] = d_1405_v5_
        nw220_[int(2)] = d_1405_v5_
        nw220_[int(3)] = d_1405_v5_
        def iife96_(_pat_let17_0):
            def iife97_(d_1408_dt__update__tmp_h0_):
                def iife99_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in (pat_let_tv24_).Elements:
                        d_1409_v6_: int = compr_62_
                        if (d_1409_v6_) in (pat_let_tv25_):
                            coll62_[default__.safeDivisionInt(d_1409_v6_, 270)] = pat_let_tv26_
                    return _dafny.Map(coll62_)
                def iife98_(_pat_let18_0):
                    def iife100_(d_1410_dt__update_hcf13_h0_):
                        return D2_DC6(d_1410_dt__update_hcf13_h0_)
                    return iife100_(_pat_let18_0)
                return iife98_(iife99_()
                )
            return iife97_(_pat_let17_0)
        nw220_[int(4)] = iife96_(d_1405_v5_)
        nw220_[int(5)] = D2_DC6(_dafny.Map({p3: p3}))
        nw220_[int(6)] = D2_DC6(d_1404_v4_)
        nw220_[int(7)] = d_1405_v5_
        nw220_[int(8)] = d_1405_v5_
        nw220_[int(9)] = d_1405_v5_
        nw220_[int(10)] = d_1405_v5_
        nw220_[int(11)] = D2_DC6(d_1404_v4_)
        nw220_[int(12)] = d_1405_v5_
        nw220_[int(13)] = d_1405_v5_
        nw220_[int(14)] = d_1405_v5_
        nw220_[int(15)] = D2_DC6(d_1404_v4_)
        nw220_[int(16)] = d_1405_v5_
        nw220_[int(17)] = D2_DC6(_dafny.Map({p3: p3}))
        d_1407_v8_ = nw220_
        d_1411_v9_: _dafny.Map
        d_1411_v9_ = _dafny.Map({((self).fm5(d_1403_v3_, d_1403_v3_, p0, not(False), globalState)) not in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1412_i0_ in range(default__.abs(999))])): d_1407_v8_})
        d_1411_v9_ = (d_1411_v9_).set((self).f26, d_1407_v8_)
        hi12_ = (0) - (p3)
        for d_1413_i1_ in range(p3, hi12_):
            d_1414_v10_: _dafny.Array
            nw221_ = _dafny.Array(False, 28)
            d_1414_v10_ = nw221_
            d_1415_v11_: _dafny.Map
            d_1415_v11_ = _dafny.Map({p1: d_1406_v7_})
            index201_ = default__.safeIndex(653, (d_1414_v10_).length(0))
            (d_1414_v10_)[index201_] = not((_dafny.Set({p3, d_1413_i1_})).issubset(((d_1415_v11_)[(self).f27] if ((self).f27) in (d_1415_v11_) else d_1406_v7_)))
            d_1416_v12_: _dafny.Map
            d_1416_v12_ = _dafny.Map({p3: p1})
            d_1417_v13_: _dafny.MultiSet
            d_1417_v13_ = _dafny.MultiSet([d_1414_v10_])
            index202_ = default__.safeIndex(653, (d_1414_v10_).length(0))
            rhs211_ = (default__.safeModuloInt(len(d_1416_v12_), d_1413_i1_)) != (297)
            rhs212_ = ((d_1417_v13_) - (_dafny.MultiSet([d_1414_v10_, d_1414_v10_]))).ispropersubset(d_1417_v13_)
            lhs192_ = globalState
            lhs193_ = d_1414_v10_
            lhs194_ = default__.safeIndex(653, (d_1414_v10_).length(0))
            lhs192_.f22 = rhs211_
            lhs193_[lhs194_] = rhs212_
            (globalState).f16 = p3
            d_1418_v14_: _dafny.MultiSet
            d_1418_v14_ = _dafny.MultiSet([(default__.fm25((d_1414_v10_)[default__.safeIndex(653, (d_1414_v10_).length(0))], (d_1414_v10_)[default__.safeIndex(653, (d_1414_v10_).length(0))], _dafny.CodePoint('q'), globalState)) + (d_1403_v3_)])
            d_1418_v14_ = (d_1418_v14_) | (_dafny.MultiSet([d_1403_v3_, d_1403_v3_, _dafny.SeqWithoutIsStrInference([(d_1414_v10_)[default__.safeIndex(653, (d_1414_v10_).length(0))], (self).fm2(p1, p3, globalState)]), _dafny.SeqWithoutIsStrInference([(self).f26, (d_1414_v10_)[default__.safeIndex(653, (d_1414_v10_).length(0))], (self).f26])]))
            (globalState).f16 = d_1413_i1_
        d_1419_v15_: _dafny.Seq
        d_1419_v15_ = _dafny.SeqWithoutIsStrInference([(self).fm1(globalState)])
        d_1420_i2_: int
        d_1420_i2_ = 0
        with _dafny.label("7"):
            while ((self).f26 if not((d_1403_v3_)[default__.safeIndex(len((d_1419_v15_).set(default__.safeIndex(p3, len(d_1419_v15_)), 60)), len(d_1403_v3_))]) else False):
                with _dafny.c_label("7"):
                    if (d_1420_i2_) >= (100):
                        raise _dafny.Break("7")
                    d_1420_i2_ = (d_1420_i2_) + (1)
                    d_1421_v16_: D0
                    d_1421_v16_ = D0_DC1((self).f27, (self).f26, len(d_1403_v3_), p1)
                    d_1422_v17_: D0
                    d_1422_v17_ = D0_DC2(d_1421_v16_)
                    d_1423_v18_: D0
                    d_1423_v18_ = D0_DC2(d_1422_v17_)
                    source22_ = D0_DC2(d_1422_v17_)
                    if source22_.is_DC1:
                        d_1424___mcc_h0_ = source22_.cf1
                        d_1425___mcc_h1_ = source22_.cf2
                        d_1426___mcc_h2_ = source22_.cf3
                        d_1427___mcc_h3_ = source22_.cf4
                        d_1428_cf4_ = d_1427___mcc_h3_
                        d_1429_cf3_ = d_1426___mcc_h2_
                        d_1430_cf2_ = d_1425___mcc_h1_
                        d_1431_cf1_ = d_1424___mcc_h0_
                        d_1432_v19_: _dafny.Array
                        nw222_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                        d_1432_v19_ = nw222_
                        index203_ = default__.safeIndex(419, (d_1432_v19_).length(0))
                        (d_1432_v19_)[index203_] = d_1401_v1_
                        index204_ = default__.safeIndex(419, (d_1432_v19_).length(0))
                        (d_1432_v19_)[index204_] = d_1401_v1_
                        d_1433_v20_: _dafny.Map
                        d_1433_v20_ = _dafny.Map({d_1429_cf3_: not (d_1430_cf2_) or (True)})
                        d_1433_v20_ = (d_1433_v20_).set(d_1429_cf3_, (d_1429_cf3_) <= (p3))
                        (globalState).f24 = (d_1428_cf4_ if (self).fm2(p1, p3, globalState) else True)
                        d_1434_v21_: _dafny.Array
                        def lambda70_(d_1435_cf2_):
                            def lambda71_(d_1436_i3_):
                                return d_1435_cf2_

                            return lambda71_

                        init40_ = lambda70_(d_1430_cf2_)
                        nw223_ = _dafny.Array(None, 28)
                        for i0_40_ in range(nw223_.length(0)):
                            nw223_[i0_40_] = init40_(i0_40_)
                        d_1434_v21_ = nw223_
                        index205_ = default__.safeIndex(404, (d_1434_v21_).length(0))
                        (d_1434_v21_)[index205_] = d_1431_cf1_
                        index206_ = default__.safeIndex(404, (d_1434_v21_).length(0))
                        (d_1434_v21_)[index206_] = not(d_1431_cf1_)
                    elif source22_.is_DC0:
                        d_1437___mcc_h4_ = source22_.cf0
                        d_1438_cf0_ = d_1437___mcc_h4_
                        d_1439_v22_: _dafny.Array
                        def lambda72_(d_1440_p3_):
                            def lambda73_(d_1441_i4_):
                                return default__.safeModuloInt(d_1441_i4_, d_1440_p3_)

                            return lambda73_

                        init41_ = lambda72_(p3)
                        nw224_ = _dafny.Array(None, 6)
                        for i0_41_ in range(nw224_.length(0)):
                            nw224_[i0_41_] = init41_(i0_41_)
                        d_1439_v22_ = nw224_
                        d_1442_v23_: _dafny.Map
                        d_1442_v23_ = _dafny.Map({(self).f26: d_1439_v22_})
                        d_1443_v24_: _dafny.Array
                        nw225_ = _dafny.Array(None, 4)
                        nw225_[int(0)] = d_1439_v22_
                        nw225_[int(1)] = d_1439_v22_
                        nw225_[int(2)] = ((d_1442_v23_)[not(True)] if (not(True)) in (d_1442_v23_) else d_1439_v22_)
                        nw225_[int(3)] = d_1439_v22_
                        d_1443_v24_ = nw225_
                        index207_ = default__.safeIndex(175, (d_1443_v24_).length(0))
                        (d_1443_v24_)[index207_] = d_1439_v22_
                        index208_ = default__.safeIndex(175, (d_1443_v24_).length(0))
                        (d_1443_v24_)[index208_] = d_1439_v22_
                        d_1444_v25_: _dafny.MultiSet
                        d_1444_v25_ = _dafny.MultiSet([p3, p3, p3, p3, p3])
                        d_1445_v26_: _dafny.Array
                        nw226_ = _dafny.Array(None, 11)
                        nw226_[int(0)] = (728) > (p3)
                        nw226_[int(1)] = d_1438_cf0_
                        nw226_[int(2)] = d_1438_cf0_
                        nw226_[int(3)] = d_1438_cf0_
                        nw226_[int(4)] = not(((d_1400_v0_)[default__.safeIndex(371, (d_1400_v0_).length(0))]) != (d_1402_v2_))
                        nw226_[int(5)] = not((d_1444_v25_).issubset(d_1444_v25_))
                        nw226_[int(6)] = (self).fm2(False, p3, globalState)
                        nw226_[int(7)] = (self).f26
                        nw226_[int(8)] = (self).f27
                        nw226_[int(9)] = p1
                        nw226_[int(10)] = True
                        d_1445_v26_ = nw226_
                        index209_ = default__.safeIndex(717, (d_1445_v26_).length(0))
                        (d_1445_v26_)[index209_] = p1
                        d_1446_v27_: D6
                        d_1446_v27_ = D6_DC16(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqncb")))
                        pat_let_tv27_ = d_1402_v2_
                        pat_let_tv28_ = d_1402_v2_
                        d_1447_v28_: _dafny.Array
                        nw227_ = _dafny.Array(None, 28)
                        nw227_[int(0)] = D6_DC16(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1448_i5_ in range(default__.abs(794))]))
                        nw227_[int(1)] = d_1446_v27_
                        nw227_[int(2)] = d_1446_v27_
                        nw227_[int(3)] = default__.fm38(True, p3, (d_1400_v0_)[default__.safeIndex(371, (d_1400_v0_).length(0))], globalState)
                        nw227_[int(4)] = d_1446_v27_
                        nw227_[int(5)] = d_1446_v27_
                        nw227_[int(6)] = d_1446_v27_
                        nw227_[int(7)] = d_1446_v27_
                        nw227_[int(8)] = d_1446_v27_
                        nw227_[int(9)] = d_1446_v27_
                        def iife101_(_pat_let19_0):
                            def iife102_(d_1449_dt__update__tmp_h1_):
                                def iife103_(_pat_let20_0):
                                    def iife104_(d_1450_dt__update_hcf29_h0_):
                                        return D6_DC16(d_1450_dt__update_hcf29_h0_)
                                    return iife104_(_pat_let20_0)
                                return iife103_(pat_let_tv27_)
                            return iife102_(_pat_let19_0)
                        nw227_[int(10)] = iife101_(D6_DC16(d_1402_v2_))
                        nw227_[int(11)] = d_1446_v27_
                        nw227_[int(12)] = D6_DC16(d_1402_v2_)
                        def iife105_(_pat_let21_0):
                            def iife106_(d_1451_dt__update__tmp_h2_):
                                def iife107_(_pat_let22_0):
                                    def iife108_(d_1452_dt__update_hcf29_h1_):
                                        return D6_DC16(d_1452_dt__update_hcf29_h1_)
                                    return iife108_(_pat_let22_0)
                                return iife107_(pat_let_tv28_)
                            return iife106_(_pat_let21_0)
                        nw227_[int(13)] = iife105_(d_1446_v27_)
                        nw227_[int(14)] = d_1446_v27_
                        nw227_[int(15)] = d_1446_v27_
                        nw227_[int(16)] = d_1446_v27_
                        nw227_[int(17)] = d_1446_v27_
                        nw227_[int(18)] = D6_DC16(d_1402_v2_)
                        nw227_[int(19)] = d_1446_v27_
                        nw227_[int(20)] = d_1446_v27_
                        nw227_[int(21)] = d_1446_v27_
                        nw227_[int(22)] = d_1446_v27_
                        nw227_[int(23)] = d_1446_v27_
                        nw227_[int(24)] = d_1446_v27_
                        nw227_[int(25)] = D6_DC16((d_1400_v0_)[default__.safeIndex(371, (d_1400_v0_).length(0))])
                        nw227_[int(26)] = d_1446_v27_
                        nw227_[int(27)] = d_1446_v27_
                        d_1447_v28_ = nw227_
                        index210_ = default__.safeIndex(485, (d_1447_v28_).length(0))
                        (d_1447_v28_)[index210_] = d_1446_v27_
                        index211_ = default__.safeIndex(717, (d_1445_v26_).length(0))
                        index212_ = default__.safeIndex(485, (d_1447_v28_).length(0))
                        rhs213_ = not((self).f27)
                        rhs214_ = d_1446_v27_
                        lhs195_ = d_1445_v26_
                        lhs196_ = default__.safeIndex(717, (d_1445_v26_).length(0))
                        lhs197_ = d_1447_v28_
                        lhs198_ = default__.safeIndex(485, (d_1447_v28_).length(0))
                        lhs195_[lhs196_] = rhs213_
                        lhs197_[lhs198_] = rhs214_
                        d_1453_v29_: _dafny.Map
                        d_1453_v29_ = _dafny.Map({D1_DC5(p3): d_1401_v1_})
                        d_1454_v30_: _dafny.Seq
                        d_1454_v30_ = _dafny.SeqWithoutIsStrInference([d_1453_v29_])
                        d_1455_v31_: _dafny.Map
                        d_1455_v31_ = _dafny.Map({(d_1454_v30_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({D1_DC5(p3): d_1401_v1_}) for d_1456_i6_ in range(default__.abs(68))])): d_1403_v3_})
                        def iife109_():
                            coll63_ = _dafny.Map()
                            compr_63_: D1
                            for compr_63_ in (_dafny.Map({D1_DC5(p3): p3})).keys.Elements:
                                d_1457_v32_: D1 = compr_63_
                                if (d_1457_v32_) in (_dafny.Map({D1_DC5(p3): p3})):
                                    coll63_[d_1457_v32_] = d_1401_v1_
                            return _dafny.Map(coll63_)
                        d_1455_v31_ = (d_1455_v31_).set((_dafny.SeqWithoutIsStrInference([iife109_()
                        ])) + (d_1454_v30_), d_1403_v3_)
                        pat_let_tv29_ = d_1400_v0_
                        pat_let_tv30_ = d_1400_v0_
                        def iife110_(_pat_let23_0):
                            def iife111_(d_1458_dt__update__tmp_h3_):
                                def iife112_(_pat_let24_0):
                                    def iife113_(d_1459_dt__update_hcf29_h2_):
                                        return D6_DC16(d_1459_dt__update_hcf29_h2_)
                                    return iife113_(_pat_let24_0)
                                return iife112_((pat_let_tv30_)[default__.safeIndex(371, (pat_let_tv29_).length(0))])
                            return iife111_(_pat_let23_0)
                        (globalState).f6 = (iife110_(d_1446_v27_)).cf29
                    elif True:
                        d_1460___mcc_h5_ = source22_.cf5
                        d_1461_cf5_ = d_1460___mcc_h5_
                        d_1462_v33_: _dafny.Array
                        def lambda74_(d_1463_i7_):
                            return (self).f26

                        init42_ = lambda74_
                        nw228_ = _dafny.Array(None, 6)
                        for i0_42_ in range(nw228_.length(0)):
                            nw228_[i0_42_] = init42_(i0_42_)
                        d_1462_v33_ = nw228_
                        d_1464_v34_: _dafny.Map
                        d_1464_v34_ = _dafny.Map({d_1462_v33_: p3})
                        d_1464_v34_ = (d_1464_v34_).set(d_1462_v33_, p3)
                        d_1465_v35_: _dafny.Map
                        d_1465_v35_ = _dafny.Map({(self).f27: len((d_1400_v0_)[default__.safeIndex(371, (d_1400_v0_).length(0))])})
                        d_1466_v36_: _dafny.MultiSet
                        d_1466_v36_ = _dafny.MultiSet([d_1465_v35_, d_1465_v35_])
                        rhs215_ = (d_1400_v0_)[default__.safeIndex(371, (d_1400_v0_).length(0))]
                        rhs216_ = p1
                        rhs217_ = ((d_1466_v36_).cardinality) > (p3)
                        rhs218_ = (self).f27
                        rhs219_ = d_1401_v1_
                        lhs199_ = globalState
                        lhs200_ = globalState
                        lhs201_ = globalState
                        lhs202_ = globalState
                        lhs203_ = globalState
                        lhs199_.f6 = rhs215_
                        lhs200_.f24 = rhs216_
                        lhs201_.f15 = rhs217_
                        lhs202_.f24 = rhs218_
                        lhs203_.f19 = rhs219_
                        (globalState).f2 = p1
                        (globalState).f16 = p3
                    d_1467_v37_: str
                    d_1468_v38_: _dafny.Seq
                    d_1469_v39_: bool
                    d_1470_v40_: _dafny.Map
                    out45_: str
                    out46_: _dafny.Seq
                    out47_: bool
                    out48_: _dafny.Map
                    out45_, out46_, out47_, out48_ = default__.m0(d_1401_v1_, globalState)
                    d_1467_v37_ = out45_
                    d_1468_v38_ = out46_
                    d_1469_v39_ = out47_
                    d_1470_v40_ = out48_
                    d_1471_v41_: _dafny.Array
                    nw229_ = _dafny.Array(_dafny.Array(None, 0), 10)
                    d_1471_v41_ = nw229_
                    d_1472_v42_: _dafny.Array
                    nw230_ = _dafny.Array(int(0), 5)
                    d_1472_v42_ = nw230_
                    index213_ = default__.safeIndex(894, (d_1471_v41_).length(0))
                    (d_1471_v41_)[index213_] = d_1472_v42_
                    d_1473_v43_: T2
                    nw231_ = C1()
                    nw231_.ctor__(p3, p3, True, False)
                    d_1473_v43_ = nw231_
                    d_1474_v44_: _dafny.Seq
                    d_1474_v44_ = _dafny.SeqWithoutIsStrInference([d_1473_v43_])
                    index214_ = default__.safeIndex(894, (d_1471_v41_).length(0))
                    rhs220_ = (d_1472_v42_ if ((self).f27) or (p1) else d_1472_v42_)
                    rhs221_ = d_1474_v44_
                    rhs222_ = (0) - ((p3) + (p3))
                    rhs223_ = (d_1473_v43_).f26
                    lhs204_ = d_1471_v41_
                    lhs205_ = default__.safeIndex(894, (d_1471_v41_).length(0))
                    lhs206_ = globalState
                    lhs207_ = globalState
                    lhs204_[lhs205_] = rhs220_
                    d_1474_v44_ = rhs221_
                    lhs206_.f12 = rhs222_
                    lhs207_.f24 = rhs223_
                    d_1475_v45_: D3
                    d_1475_v45_ = D3_DC10((d_1471_v41_)[default__.safeIndex(894, (d_1471_v41_).length(0))])
                    d_1476_v46_: _dafny.Map
                    d_1476_v46_ = _dafny.Map({(d_1473_v43_).f26: (self).f26})
                    d_1477_v47_: _dafny.Map
                    d_1477_v47_ = _dafny.Map({d_1475_v45_: len(d_1476_v46_)})
                    d_1478_v48_: T1
                    nw232_ = C4()
                    nw232_.ctor__(d_1477_v47_, (d_1473_v43_).f26, (self).f26)
                    d_1478_v48_ = nw232_
                    d_1479_v49_: _dafny.Map
                    d_1479_v49_ = _dafny.Map({d_1478_v48_: not(False)})
                    d_1479_v49_ = (d_1479_v49_).set(d_1478_v48_, True)
                    pass
            pass
        r0 = not ((d_1403_v3_)[default__.safeIndex(p3, len(d_1403_v3_))]) or ((self).f27)
        return r0


class C6(T0):
    def  __init__(self):
        self.f40: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f40):
        (self).f40 = f40

    def fm1(self, globalState):
        return ((self.f40) - (len(_dafny.SeqWithoutIsStrInference([self.f40 for d_1480_i0_ in range(default__.abs(113))])))) - (self.f40)

    def fm2(self, p0, p1, globalState):
        def iife114_():
            coll64_ = _dafny.Set()
            compr_64_: int
            for compr_64_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f40, len(_dafny.Map({False: not(False)})), (0) - (self.f40)]))).Elements:
                d_1481_v0_: int = compr_64_
                if (d_1481_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f40, len(_dafny.Map({False: not(False)})), (0) - (self.f40)]))):
                    coll64_ = coll64_.union(_dafny.Set([default__.safeModuloInt(d_1481_v0_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True), False, False, not(False), True]))).cardinality)]))
            return _dafny.Set(coll64_)
        def iife115_():
            coll65_ = _dafny.Set()
            compr_65_: int
            for compr_65_ in _dafny.IntegerRange(178, 967):
                d_1483_v1_: int = compr_65_
                if ((178) <= (d_1483_v1_)) and ((d_1483_v1_) < (967)):
                    coll65_ = coll65_.union(_dafny.Set([(d_1483_v1_) * (self.f40)]))
            return _dafny.Set(coll65_)
        return ((iife114_()
        ) - (_dafny.Set({(0) - (self.f40), (_dafny.MultiSet([(_dafny.MultiSet([self.f40])).cardinality])).cardinality, (_dafny.MultiSet([True, not(True)])).cardinality, self.f40}))) == ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1482_i0_ in range(default__.abs(400))])), (0) - (self.f40), self.f40})) - (_dafny.Set({self.f40, len(_dafny.SeqWithoutIsStrInference([iife115_()
        ]))})))

    def fm44(self, globalState):
        return (not(not(False))) and (not((not(False)) == (True)))

    def m15(self, p0, p1, p2, p3, globalState):
        r0: D1 = D1.default()()
        r1: _dafny.Map = _dafny.Map({})
        r2: int = int(0)
        (globalState).f12 = (self).fm1(globalState)
        d_1484_v0_: _dafny.MultiSet
        d_1484_v0_ = _dafny.MultiSet([p2])
        d_1485_v1_: _dafny.Map
        d_1485_v1_ = _dafny.Map({((d_1484_v0_).cardinality) + (866): p2})
        d_1486_v2_: D16
        d_1486_v2_ = D16_DC37(d_1485_v1_)
        d_1485_v1_ = (d_1485_v1_) | ((d_1486_v2_).cf68)
        (globalState).f15 = (p3) and (p2)
        d_1487_v3_: _dafny.Seq
        d_1487_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymvn"))
        d_1488_v4_: _dafny.Map
        d_1488_v4_ = _dafny.Map({d_1487_v3_: p2})
        (globalState).f12 = (0) - (((d_1484_v0_)[(_dafny.Map({d_1487_v3_: p2})) != (d_1488_v4_)] if ((_dafny.Map({d_1487_v3_: p2})) != (d_1488_v4_)) in (d_1484_v0_) else p0))
        d_1489_i0_: int
        d_1489_i0_ = 0
        with _dafny.label("8"):
            while p2:
                with _dafny.c_label("8"):
                    if (d_1489_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1489_i0_ = (d_1489_i0_) + (1)
                    d_1490_v5_: D15
                    d_1490_v5_ = D15_DC35(p2)
                    d_1490_v5_ = d_1490_v5_
                    d_1491_v6_: _dafny.Array
                    def lambda75_(d_1492_i1_):
                        return default__.safeDivisionInt(d_1492_i1_, 669)

                    init43_ = lambda75_
                    nw233_ = _dafny.Array(None, 25)
                    for i0_43_ in range(nw233_.length(0)):
                        nw233_[i0_43_] = init43_(i0_43_)
                    d_1491_v6_ = nw233_
                    index215_ = default__.safeIndex(182, (d_1491_v6_).length(0))
                    (d_1491_v6_)[index215_] = self.f40
                    index216_ = default__.safeIndex(182, (d_1491_v6_).length(0))
                    (d_1491_v6_)[index216_] = 134
                    d_1493_v7_: T2
                    nw234_ = C1()
                    nw234_.ctor__(p0, self.f40, p2, p2)
                    d_1493_v7_ = nw234_
                    d_1494_v8_: D7
                    d_1494_v8_ = D7_DC18(d_1493_v7_)
                    pat_let_tv31_ = d_1493_v7_
                    pat_let_tv32_ = d_1493_v7_
                    d_1495_v9_: _dafny.Array
                    nw235_ = _dafny.Array(None, 18)
                    nw235_[int(0)] = d_1494_v8_
                    def iife116_(_pat_let25_0):
                        def iife117_(d_1496_dt__update__tmp_h0_):
                            def iife118_(_pat_let26_0):
                                def iife119_(d_1497_dt__update_hcf35_h0_):
                                    return D7_DC18(d_1497_dt__update_hcf35_h0_)
                                return iife119_(_pat_let26_0)
                            return iife118_(pat_let_tv31_)
                        return iife117_(_pat_let25_0)
                    nw235_[int(1)] = iife116_(d_1494_v8_)
                    nw235_[int(2)] = d_1494_v8_
                    nw235_[int(3)] = d_1494_v8_
                    nw235_[int(4)] = d_1494_v8_
                    nw235_[int(5)] = d_1494_v8_
                    nw235_[int(6)] = d_1494_v8_
                    def iife120_(_pat_let27_0):
                        def iife121_(d_1498_dt__update__tmp_h1_):
                            def iife122_(_pat_let28_0):
                                def iife123_(d_1499_dt__update_hcf35_h1_):
                                    return D7_DC18(d_1499_dt__update_hcf35_h1_)
                                return iife123_(_pat_let28_0)
                            return iife122_(pat_let_tv32_)
                        return iife121_(_pat_let27_0)
                    nw235_[int(7)] = iife120_(D7_DC18(d_1493_v7_))
                    nw235_[int(8)] = d_1494_v8_
                    nw235_[int(9)] = d_1494_v8_
                    nw235_[int(10)] = d_1494_v8_
                    nw235_[int(11)] = d_1494_v8_
                    nw235_[int(12)] = d_1494_v8_
                    nw235_[int(13)] = d_1494_v8_
                    nw235_[int(14)] = d_1494_v8_
                    nw235_[int(15)] = D7_DC18(d_1493_v7_)
                    nw235_[int(16)] = d_1494_v8_
                    nw235_[int(17)] = d_1494_v8_
                    d_1495_v9_ = nw235_
                    index217_ = default__.safeIndex(743, (d_1495_v9_).length(0))
                    (d_1495_v9_)[index217_] = d_1494_v8_
                    index218_ = default__.safeIndex(743, (d_1495_v9_).length(0))
                    (d_1495_v9_)[index218_] = d_1494_v8_
                    d_1500_v10_: _dafny.Array
                    nw236_ = _dafny.Array(False, 7)
                    d_1500_v10_ = nw236_
                    index219_ = default__.safeIndex(263, (d_1500_v10_).length(0))
                    (d_1500_v10_)[index219_] = (d_1493_v7_).f26
                    index220_ = default__.safeIndex(263, (d_1500_v10_).length(0))
                    (d_1500_v10_)[index220_] = p2
                    pass
            pass
        d_1501_i2_: int
        d_1501_i2_ = 0
        with _dafny.label("9"):
            while p2:
                with _dafny.c_label("9"):
                    if (d_1501_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_1501_i2_ = (d_1501_i2_) + (1)
                    d_1502_v11_: _dafny.Seq
                    d_1502_v11_ = _dafny.SeqWithoutIsStrInference([p3])
                    d_1503_v12_: _dafny.Set
                    d_1503_v12_ = _dafny.Set({self.f40, self.f40, p0, p0, p0})
                    d_1504_v13_: _dafny.Array
                    nw237_ = _dafny.Array(None, 8)
                    nw237_[int(0)] = p3
                    nw237_[int(1)] = (d_1502_v11_)[default__.safeIndex(len(d_1503_v12_), len(d_1502_v11_))]
                    nw237_[int(2)] = True
                    nw237_[int(3)] = (self).fm2(p2, p0, globalState)
                    nw237_[int(4)] = True
                    nw237_[int(5)] = p3
                    nw237_[int(6)] = p2
                    nw237_[int(7)] = p2
                    d_1504_v13_ = nw237_
                    index221_ = default__.safeIndex(147, (d_1504_v13_).length(0))
                    (d_1504_v13_)[index221_] = (self).fm44(globalState)
                    index222_ = default__.safeIndex(147, (d_1504_v13_).length(0))
                    (d_1504_v13_)[index222_] = (self.f40) > (((self).fm1(globalState)) + ((0) - (p0)))
                    r2 = p0
                    rhs224_ = (self).fm1(globalState)
                    rhs225_ = (self).fm1(globalState)
                    rhs226_ = len(_dafny.SeqWithoutIsStrInference([-197, (0) - ((self).fm1(globalState)), len(d_1503_v12_), ((0) - (self.f40)) + (self.f40), 452]))
                    rhs227_ = d_1487_v3_
                    lhs208_ = globalState
                    lhs209_ = globalState
                    lhs210_ = globalState
                    lhs211_ = globalState
                    lhs208_.f16 = rhs224_
                    lhs209_.f12 = rhs225_
                    lhs210_.f12 = rhs226_
                    lhs211_.f6 = rhs227_
                    d_1505_v14_: _dafny.Seq
                    d_1505_v14_ = _dafny.SeqWithoutIsStrInference([len(d_1487_v3_)])
                    d_1506_v15_: _dafny.MultiSet
                    d_1506_v15_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0, p0]), d_1505_v14_, d_1505_v14_])
                    d_1507_v16_: C3
                    nw238_ = C3()
                    nw238_.ctor__(d_1506_v15_, (p3 if (d_1504_v13_)[default__.safeIndex(147, (d_1504_v13_).length(0))] else True), (d_1504_v13_)[default__.safeIndex(147, (d_1504_v13_).length(0))])
                    d_1507_v16_ = nw238_
                    pass
            pass
        r0 = D1_DC5(default__.safeModuloInt(p0, p0))
        d_1508_v17_: _dafny.Seq
        d_1508_v17_ = _dafny.SeqWithoutIsStrInference([p3])
        d_1509_v18_: _dafny.Map
        d_1509_v18_ = _dafny.Map({d_1484_v0_: (d_1508_v17_)[default__.safeIndex(self.f40, len(d_1508_v17_))]})
        r1 = d_1509_v18_
        r2 = (self.f40) * (self.f40)
        return r0, r1, r2

    def m16(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        d_1510_v0_: _dafny.Seq
        d_1510_v0_ = _dafny.SeqWithoutIsStrInference([self.f40])
        d_1511_v1_: _dafny.Seq
        d_1511_v1_ = _dafny.SeqWithoutIsStrInference([d_1510_v0_])
        d_1512_v2_: bool
        d_1512_v2_ = False
        rhs228_ = p0
        rhs229_ = (_dafny.SeqWithoutIsStrInference([default__.fm45(d_1512_v2_, p0, -164, d_1512_v2_, globalState), d_1510_v0_, d_1510_v0_])) + (d_1511_v1_)
        lhs212_ = globalState
        lhs212_.f16 = rhs228_
        d_1511_v1_ = rhs229_
        rhs230_ = p0
        rhs231_ = self.f40
        lhs213_ = globalState
        lhs214_ = globalState
        lhs213_.f16 = rhs230_
        lhs214_.f12 = rhs231_
        d_1513_v3_: _dafny.Array
        nw239_ = _dafny.Array(int(0), 24)
        d_1513_v3_ = nw239_
        index223_ = default__.safeIndex(946, (d_1513_v3_).length(0))
        (d_1513_v3_)[index223_] = p0
        d_1514_v4_: _dafny.Map
        d_1514_v4_ = _dafny.Map({self.f40: default__.fm47(d_1512_v2_, d_1512_v2_, d_1512_v2_, not(d_1512_v2_), globalState)})
        d_1515_v5_: _dafny.Seq
        d_1515_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mswelsh"))
        d_1516_v6_: _dafny.Map
        d_1516_v6_ = _dafny.Map({len(d_1515_v5_): d_1512_v2_})
        d_1517_v7_: _dafny.MultiSet
        d_1517_v7_ = _dafny.MultiSet([(0) - (p0)])
        d_1518_v8_: _dafny.Map
        d_1518_v8_ = _dafny.Map({d_1517_v7_: self.f40})
        d_1519_v9_: _dafny.Map
        d_1519_v9_ = _dafny.Map({_dafny.Map({d_1512_v2_: self.f40}): len(d_1518_v8_)})
        d_1520_v10_: _dafny.Array
        nw240_ = _dafny.Array(None, 4)
        nw240_[int(0)] = default__.fm46(d_1512_v2_, p0, p0, d_1512_v2_, globalState)
        nw240_[int(1)] = d_1514_v4_
        nw240_[int(2)] = (d_1514_v4_) | (_dafny.Map({p0: d_1516_v6_}))
        nw240_[int(3)] = default__.fm46(d_1512_v2_, ((d_1519_v9_)[_dafny.Map({d_1512_v2_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxbdjm")))})] if (_dafny.Map({d_1512_v2_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxbdjm")))})) in (d_1519_v9_) else self.f40), p0, d_1512_v2_, globalState)
        d_1520_v10_ = nw240_
        index224_ = default__.safeIndex(385, (d_1520_v10_).length(0))
        (d_1520_v10_)[index224_] = d_1514_v4_
        d_1521_v11_: _dafny.Array
        nw241_ = _dafny.Array(False, 27)
        d_1521_v11_ = nw241_
        d_1522_v12_: _dafny.Seq
        d_1522_v12_ = _dafny.SeqWithoutIsStrInference([d_1521_v11_, d_1521_v11_])
        d_1523_v13_: _dafny.Seq
        d_1523_v13_ = _dafny.SeqWithoutIsStrInference([d_1522_v12_])
        d_1524_v14_: _dafny.Array
        nw242_ = _dafny.Array(None, 25)
        nw242_[int(0)] = (d_1522_v12_) + (d_1522_v12_)
        nw242_[int(1)] = d_1522_v12_
        nw242_[int(2)] = d_1522_v12_
        nw242_[int(3)] = d_1522_v12_
        nw242_[int(4)] = d_1522_v12_
        nw242_[int(5)] = d_1522_v12_
        nw242_[int(6)] = (d_1522_v12_) + ((d_1523_v13_)[default__.safeIndex(490, len(d_1523_v13_))])
        nw242_[int(7)] = d_1522_v12_
        nw242_[int(8)] = _dafny.SeqWithoutIsStrInference([d_1521_v11_, d_1521_v11_])
        nw242_[int(9)] = ((_dafny.SeqWithoutIsStrInference([d_1521_v11_])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajrpp")))])), len(_dafny.SeqWithoutIsStrInference([d_1521_v11_]))), d_1521_v11_)) + (d_1522_v12_)
        nw242_[int(10)] = d_1522_v12_
        nw242_[int(11)] = d_1522_v12_
        nw242_[int(12)] = d_1522_v12_
        nw242_[int(13)] = d_1522_v12_
        nw242_[int(14)] = d_1522_v12_
        nw242_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1521_v11_, d_1521_v11_, d_1521_v11_])
        nw242_[int(16)] = d_1522_v12_
        nw242_[int(17)] = (d_1522_v12_).set(default__.safeIndex(p0, len(d_1522_v12_)), d_1521_v11_)
        nw242_[int(18)] = (d_1522_v12_) + ((d_1522_v12_).set(default__.safeIndex(self.f40, len(d_1522_v12_)), d_1521_v11_))
        nw242_[int(19)] = d_1522_v12_
        nw242_[int(20)] = d_1522_v12_
        nw242_[int(21)] = d_1522_v12_
        nw242_[int(22)] = (d_1522_v12_) + (d_1522_v12_)
        nw242_[int(23)] = d_1522_v12_
        nw242_[int(24)] = d_1522_v12_
        d_1524_v14_ = nw242_
        index225_ = default__.safeIndex(119, (d_1524_v14_).length(0))
        (d_1524_v14_)[index225_] = d_1522_v12_
        d_1525_v15_: str
        d_1525_v15_ = _dafny.CodePoint('k')
        d_1526_v16_: _dafny.Map
        d_1526_v16_ = _dafny.Map({_dafny.MultiSet([p0]): d_1525_v15_})
        d_1527_v17_: _dafny.Map
        d_1527_v17_ = _dafny.Map({((d_1526_v16_)[d_1517_v7_] if (d_1517_v7_) in (d_1526_v16_) else (d_1515_v5_)[default__.safeIndex((0) - (p0), len(d_1515_v5_))]): d_1512_v2_})
        index226_ = default__.safeIndex(946, (d_1513_v3_).length(0))
        index227_ = default__.safeIndex(385, (d_1520_v10_).length(0))
        index228_ = default__.safeIndex(119, (d_1524_v14_).length(0))
        rhs232_ = p0
        rhs233_ = (d_1512_v2_ if d_1512_v2_ else (d_1512_v2_ if not(False) else d_1512_v2_))
        rhs234_ = d_1514_v4_
        rhs235_ = (d_1522_v12_).set(default__.safeIndex(p0, len(d_1522_v12_)), d_1521_v11_)
        rhs236_ = d_1527_v17_
        lhs215_ = d_1513_v3_
        lhs216_ = default__.safeIndex(946, (d_1513_v3_).length(0))
        lhs217_ = globalState
        lhs218_ = d_1520_v10_
        lhs219_ = default__.safeIndex(385, (d_1520_v10_).length(0))
        lhs220_ = d_1524_v14_
        lhs221_ = default__.safeIndex(119, (d_1524_v14_).length(0))
        lhs215_[lhs216_] = rhs232_
        lhs217_.f24 = rhs233_
        lhs218_[lhs219_] = rhs234_
        lhs220_[lhs221_] = rhs235_
        d_1527_v17_ = rhs236_
        (globalState).f24 = not(not(d_1512_v2_))
        d_1528_v18_: D6
        d_1528_v18_ = D6_DC16(_dafny.SeqWithoutIsStrInference([d_1525_v15_ for d_1529_i0_ in range(default__.abs(977))]))
        pat_let_tv33_ = d_1512_v2_
        pat_let_tv34_ = d_1512_v2_
        def lambda76_(source23_):
            if source23_.is_DC17:
                d_1530___mcc_h0_ = source23_.cf30
                d_1531___mcc_h1_ = source23_.cf31
                d_1532___mcc_h2_ = source23_.cf32
                d_1533___mcc_h3_ = source23_.cf33
                d_1534___mcc_h4_ = source23_.cf34
                d_1535_cf34_ = d_1534___mcc_h4_
                d_1536_cf33_ = d_1533___mcc_h3_
                d_1537_cf32_ = d_1532___mcc_h2_
                d_1538_cf31_ = d_1531___mcc_h1_
                d_1539_cf30_ = d_1530___mcc_h0_
                return pat_let_tv33_
            elif True:
                d_1540___mcc_h5_ = source23_.cf29
                d_1541_cf29_ = d_1540___mcc_h5_
                return pat_let_tv34_

        if lambda76_((d_1528_v18_ if d_1512_v2_ else d_1528_v18_)):
            d_1542_v19_: _dafny.Seq
            d_1542_v19_ = _dafny.SeqWithoutIsStrInference([d_1512_v2_])
            index229_ = default__.safeIndex(946, (d_1513_v3_).length(0))
            (d_1513_v3_)[index229_] = len((d_1542_v19_) + (((d_1542_v19_).set(default__.safeIndex(self.f40, len(d_1542_v19_)), d_1512_v2_)) + (d_1542_v19_)))
            (globalState).f24 = d_1512_v2_
            d_1513_v3_ = d_1513_v3_
            d_1513_v3_ = d_1513_v3_
            d_1543_v20_: D1
            d_1543_v20_ = D1_DC5((0) - (self.f40))
            d_1544_v21_: _dafny.Set
            d_1544_v21_ = _dafny.Set({d_1543_v20_})
            d_1545_v22_: _dafny.Map
            d_1545_v22_ = _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yeue")), d_1515_v5_]))})
            rhs237_ = (_dafny.Set({d_1543_v20_, d_1543_v20_, d_1543_v20_})) - (d_1544_v21_)
            rhs238_ = d_1545_v22_
            rhs239_ = d_1521_v11_
            d_1544_v21_ = rhs237_
            d_1545_v22_ = rhs238_
            d_1521_v11_ = rhs239_
        elif True:
            d_1546_v23_: _dafny.Seq
            d_1546_v23_ = _dafny.SeqWithoutIsStrInference([d_1515_v5_])
            d_1547_v24_: _dafny.Array
            nw243_ = _dafny.Array(None, 2)
            nw243_[int(0)] = d_1525_v15_
            nw243_[int(1)] = d_1525_v15_
            d_1547_v24_ = nw243_
            d_1548_v25_: D1
            d_1548_v25_ = D1_DC4(d_1547_v24_, p0, p0, d_1515_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjntvhbdw")))
            d_1549_v26_: _dafny.Array
            nw244_ = _dafny.Array(None, 20)
            nw244_[int(0)] = d_1515_v5_
            nw244_[int(1)] = d_1515_v5_
            nw244_[int(2)] = d_1515_v5_
            nw244_[int(3)] = d_1515_v5_
            nw244_[int(4)] = (d_1546_v23_)[default__.safeIndex(self.f40, len(d_1546_v23_))]
            nw244_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
            nw244_[int(6)] = d_1515_v5_
            nw244_[int(7)] = d_1515_v5_
            nw244_[int(8)] = d_1515_v5_
            nw244_[int(9)] = d_1515_v5_
            nw244_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agk"))
            nw244_[int(11)] = (d_1548_v25_).cf10
            nw244_[int(12)] = (d_1515_v5_) + (d_1515_v5_)
            nw244_[int(13)] = _dafny.SeqWithoutIsStrInference([d_1525_v15_ for d_1550_i1_ in range(default__.abs(566))])
            nw244_[int(14)] = d_1515_v5_
            nw244_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwu"))
            nw244_[int(16)] = (d_1515_v5_) + (d_1515_v5_)
            nw244_[int(17)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdid"))
            nw244_[int(18)] = (d_1515_v5_) + (d_1515_v5_)
            nw244_[int(19)] = d_1515_v5_
            d_1549_v26_ = nw244_
            index230_ = default__.safeIndex(141, (d_1549_v26_).length(0))
            (d_1549_v26_)[index230_] = d_1515_v5_
            d_1551_v27_: _dafny.Set
            d_1551_v27_ = _dafny.Set({d_1512_v2_})
            d_1552_v28_: _dafny.MultiSet
            d_1552_v28_ = _dafny.MultiSet([True])
            index231_ = default__.safeIndex(141, (d_1549_v26_).length(0))
            (d_1549_v26_)[index231_] = (d_1515_v5_) + (((d_1515_v5_ if True else (d_1515_v5_).set(default__.safeIndex(len(d_1551_v27_), len(d_1515_v5_)), d_1525_v15_))).set(default__.safeIndex(((d_1552_v28_)[False] if (False) in (d_1552_v28_) else len(_dafny.Map({(d_1517_v7_).cardinality: d_1512_v2_}))), len((d_1515_v5_ if True else (d_1515_v5_).set(default__.safeIndex(len(d_1551_v27_), len(d_1515_v5_)), d_1525_v15_)))), d_1525_v15_))
            d_1553_v29_: _dafny.Map
            d_1553_v29_ = _dafny.Map({d_1512_v2_: d_1512_v2_})
            (globalState).f22 = not((d_1512_v2_) or ((self.f40) <= (len(d_1553_v29_))))
            d_1554_v30_: D5
            d_1554_v30_ = D5_DC14(p0, self.f40, d_1512_v2_)
            d_1555_v31_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.Seq({}), 17)
            d_1555_v31_ = nw245_
            d_1556_v32_: _dafny.Map
            d_1556_v32_ = _dafny.Map({d_1555_v31_: (d_1513_v3_ if d_1512_v2_ else d_1513_v3_)})
            d_1557_v33_: _dafny.Map
            d_1557_v33_ = _dafny.Map({d_1512_v2_: (0) - (self.f40)})
            d_1558_v34_: D7
            d_1558_v34_ = D7_DC19((0) - ((d_1513_v3_)[default__.safeIndex(946, (d_1513_v3_).length(0))]), d_1552_v28_, d_1512_v2_, _dafny.Map({d_1512_v2_: (d_1513_v3_)[default__.safeIndex(946, (d_1513_v3_).length(0))]}))
            d_1559_v35_: D7
            d_1559_v35_ = D7_DC19(self.f40, d_1552_v28_, d_1512_v2_, (d_1558_v34_).cf39)
            pat_let_tv35_ = globalState
            def iife124_(_pat_let29_0):
                def iife125_(d_1560_dt__update__tmp_h0_):
                    def iife126_(_pat_let30_0):
                        def iife127_(d_1561_dt__update_hcf26_h0_):
                            def iife128_(_pat_let31_0):
                                def iife129_(d_1562_dt__update_hcf25_h0_):
                                    return D5_DC14(d_1562_dt__update_hcf25_h0_, d_1561_dt__update_hcf26_h0_, (d_1560_dt__update__tmp_h0_).cf27)
                                return iife129_(_pat_let31_0)
                            return iife128_(176)
                        return iife127_(_pat_let30_0)
                    return iife126_((self).fm1(pat_let_tv35_))
                return iife125_(_pat_let29_0)
            rhs240_ = ((d_1556_v32_)[d_1555_v31_] if (d_1555_v31_) in (d_1556_v32_) else d_1513_v3_)
            rhs241_ = (((d_1557_v33_)[d_1512_v2_] if (d_1512_v2_) in (d_1557_v33_) else len(d_1557_v33_))) <= ((d_1559_v35_).cf36)
            rhs242_ = iife124_(d_1554_v30_)
            lhs222_ = globalState
            d_1513_v3_ = rhs240_
            lhs222_.f24 = rhs241_
            d_1554_v30_ = rhs242_
            d_1522_v12_ = (d_1522_v12_) + ((d_1524_v14_)[default__.safeIndex(119, (d_1524_v14_).length(0))])
            d_1563_v36_: _dafny.Array
            nw246_ = _dafny.Array(None, 2)
            nw246_[int(0)] = d_1513_v3_
            nw246_[int(1)] = d_1513_v3_
            d_1563_v36_ = nw246_
            index232_ = default__.safeIndex(542, (d_1563_v36_).length(0))
            (d_1563_v36_)[index232_] = d_1513_v3_
            index233_ = default__.safeIndex(542, (d_1563_v36_).length(0))
            (d_1563_v36_)[index233_] = (d_1513_v3_ if True else d_1513_v3_)
        d_1564_v37_: _dafny.Set
        d_1564_v37_ = _dafny.Set({(d_1513_v3_)[default__.safeIndex(946, (d_1513_v3_).length(0))], len(d_1515_v5_), p0, (d_1513_v3_)[default__.safeIndex(946, (d_1513_v3_).length(0))]})
        d_1565_v38_: _dafny.Map
        d_1565_v38_ = _dafny.Map({d_1512_v2_: d_1564_v37_})
        d_1565_v38_ = (d_1565_v38_).set((d_1515_v5_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krmbocy"))), d_1564_v37_)
        d_1566_v39_: _dafny.Seq
        d_1566_v39_ = _dafny.SeqWithoutIsStrInference([d_1512_v2_, d_1512_v2_, d_1512_v2_, d_1512_v2_, d_1512_v2_])
        r0 = (d_1566_v39_) + (_dafny.SeqWithoutIsStrInference([d_1512_v2_, d_1512_v2_, False, d_1512_v2_]))
        return r0


class C7(T0):
    def  __init__(self):
        self.f38: bool = False
        self.f39: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self, f38, f39):
        (self).f38 = f38
        (self).f39 = f39

    def fm1(self, globalState):
        def lambda77_(source24_):
            d_1567___mcc_h0_ = source24_
            d_1568_cf40_ = d_1567___mcc_h0_
            return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f38, False, self.f38, self.f38, True])): 428})) | (_dafny.Map({818: len(_dafny.Map({False: (D2_DC9(795, -873, self.f38)).cf20}))}))

        return len(lambda77_(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.Map({True: -746})): self.f38}))])))

    def fm2(self, p0, p1, globalState):
        def iife130_():
            coll66_ = _dafny.Set()
            compr_66_: _dafny.MultiSet
            for compr_66_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (-201), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imnioove"))), 112]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([373])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([self.f38, self.f38])).cardinality])), _dafny.MultiSet([987])])).Elements:
                d_1569_v0_: _dafny.MultiSet = compr_66_
                if (d_1569_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(0) - (-201), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imnioove"))), 112]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([373])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([self.f38, self.f38])).cardinality])), _dafny.MultiSet([987])])):
                    coll66_ = coll66_.union(_dafny.Set([d_1569_v0_]))
            return _dafny.Set(coll66_)
        return (iife130_()
        ).ispropersubset((_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([699, -34, -43, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqgjwd"))), len(_dafny.Map({not(self.f38): True}))])), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhuvf")))])})) - (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihdf"))), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro")) for d_1570_i0_ in range(default__.abs(-567))])), 871]))])})))

    def m13(self, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        d_1571_v0_: int
        d_1571_v0_ = 616
        d_1572_v1_: D12
        d_1572_v1_ = D12_DC28(d_1571_v0_)
        d_1573_v2_: _dafny.Map
        d_1573_v2_ = _dafny.Map({(self).fm2(not(self.f38), d_1571_v0_, globalState): d_1571_v0_})
        pat_let_tv38_ = d_1571_v0_
        pat_let_tv39_ = d_1571_v0_
        def iife140_(_pat_let37_0):
            def iife141_(d_1601_dt__update__tmp_h0_):
                def iife142_(_pat_let38_0):
                    def iife143_(d_1602_dt__update_hcf49_h0_):
                        return D12_DC28(d_1602_dt__update_hcf49_h0_)
                    return iife143_(_pat_let38_0)
                return iife142_(pat_let_tv38_)
            return iife141_(_pat_let37_0)
        def iife139_(_pat_let36_0):
            def iife144_(d_1603_dt__update__tmp_h1_):
                def iife145_(_pat_let39_0):
                    def iife146_(d_1604_dt__update_hcf49_h1_):
                        return D12_DC28(d_1604_dt__update_hcf49_h1_)
                    return iife146_(_pat_let39_0)
                return iife145_(pat_let_tv39_)
            return iife144_(_pat_let36_0)
        hi13_ = ((d_1573_v2_)[self.f38] if (self.f38) in (d_1573_v2_) else d_1571_v0_)
        for d_1574_i0_ in range(len(_dafny.SeqWithoutIsStrInference([d_1572_v1_, d_1572_v1_, d_1572_v1_, iife139_(iife140_(d_1572_v1_)), d_1572_v1_])), hi13_):
            d_1575_v3_: _dafny.Map
            d_1575_v3_ = _dafny.Map({d_1571_v0_: d_1574_i0_})
            d_1576_v4_: D0
            d_1576_v4_ = D0_DC1(self.f38, self.f38, len(d_1575_v3_), self.f38)
            d_1576_v4_ = d_1576_v4_
            d_1577_v5_: D5
            d_1577_v5_ = D5_DC15(D5_DC14(d_1574_i0_, (self).fm1(globalState), (self).fm2(self.f38, d_1571_v0_, globalState)))
            d_1578_v6_: D5
            d_1578_v6_ = D5_DC15(d_1577_v5_)
            d_1579_v7_: _dafny.Array
            nw247_ = _dafny.Array(None, 8)
            nw247_[int(0)] = d_1578_v6_
            nw247_[int(1)] = D5_DC15(d_1577_v5_)
            nw247_[int(2)] = d_1578_v6_
            nw247_[int(3)] = d_1578_v6_
            nw247_[int(4)] = d_1578_v6_
            nw247_[int(5)] = d_1578_v6_
            nw247_[int(6)] = d_1578_v6_
            nw247_[int(7)] = d_1578_v6_
            d_1579_v7_ = nw247_
            index234_ = default__.safeIndex(53, (d_1579_v7_).length(0))
            (d_1579_v7_)[index234_] = (d_1578_v6_ if not(not(self.f38)) else d_1578_v6_)
            pat_let_tv36_ = d_1577_v5_
            pat_let_tv37_ = d_1577_v5_
            index235_ = default__.safeIndex(53, (d_1579_v7_).length(0))
            def iife132_(_pat_let33_0):
                def iife133_(d_1580_dt__update__tmp_h2_):
                    def iife134_(_pat_let34_0):
                        def iife135_(d_1581_dt__update_hcf28_h0_):
                            return D5_DC15(d_1581_dt__update_hcf28_h0_)
                        return iife135_(_pat_let34_0)
                    return iife134_(pat_let_tv36_)
                return iife133_(_pat_let33_0)
            def iife131_(_pat_let32_0):
                def iife136_(d_1582_dt__update__tmp_h3_):
                    def iife137_(_pat_let35_0):
                        def iife138_(d_1583_dt__update_hcf28_h1_):
                            return D5_DC15(d_1583_dt__update_hcf28_h1_)
                        return iife138_(_pat_let35_0)
                    return iife137_(pat_let_tv37_)
                return iife136_(_pat_let32_0)
            (d_1579_v7_)[index235_] = iife131_(iife132_(D5_DC15(D5_DC14((0) - (d_1571_v0_), d_1574_i0_, self.f38))))
            d_1584_v8_: _dafny.Seq
            d_1584_v8_ = _dafny.SeqWithoutIsStrInference([d_1571_v0_])
            d_1585_v9_: _dafny.Seq
            d_1585_v9_ = _dafny.SeqWithoutIsStrInference([(d_1584_v8_).set(default__.safeIndex(d_1574_i0_, len(d_1584_v8_)), d_1571_v0_)])
            d_1586_v10_: C2
            nw248_ = C2()
            nw248_.ctor__((d_1585_v9_).set(default__.safeIndex(d_1571_v0_, len(d_1585_v9_)), _dafny.SeqWithoutIsStrInference([d_1571_v0_, 250, d_1571_v0_, d_1574_i0_])), not(self.f38), self.f38)
            d_1586_v10_ = nw248_
            d_1587_v11_: _dafny.MultiSet
            d_1587_v11_ = _dafny.MultiSet([-248])
            d_1588_v12_: _dafny.Seq
            d_1588_v12_ = _dafny.SeqWithoutIsStrInference([self.f38, self.f38])
            d_1589_v13_: _dafny.Array
            nw249_ = _dafny.Array(None, 19)
            nw249_[int(0)] = d_1574_i0_
            nw249_[int(1)] = (d_1587_v11_).cardinality
            nw249_[int(2)] = d_1574_i0_
            nw249_[int(3)] = 812
            nw249_[int(4)] = 661
            nw249_[int(5)] = d_1574_i0_
            nw249_[int(6)] = d_1574_i0_
            nw249_[int(7)] = d_1574_i0_
            nw249_[int(8)] = len(d_1588_v12_)
            nw249_[int(9)] = d_1574_i0_
            nw249_[int(10)] = d_1574_i0_
            nw249_[int(11)] = d_1571_v0_
            nw249_[int(12)] = d_1571_v0_
            nw249_[int(13)] = (d_1584_v8_)[default__.safeIndex(d_1574_i0_, len(d_1584_v8_))]
            nw249_[int(14)] = d_1571_v0_
            nw249_[int(15)] = d_1574_i0_
            nw249_[int(16)] = d_1574_i0_
            nw249_[int(17)] = d_1574_i0_
            nw249_[int(18)] = d_1574_i0_
            d_1589_v13_ = nw249_
            d_1590_v14_: _dafny.Map
            d_1590_v14_ = _dafny.Map({d_1589_v13_: False})
            d_1591_v15_: D5
            d_1591_v15_ = D5_DC14(d_1574_i0_, (0) - (len(d_1588_v12_)), self.f38)
            d_1592_v16_: D2
            d_1592_v16_ = D2_DC7(self.f38, self.f38, len(d_1590_v14_), (d_1591_v15_).cf25)
            d_1593_v17_: _dafny.Array
            nw250_ = _dafny.Array(False, 7)
            d_1593_v17_ = nw250_
            index236_ = default__.safeIndex(687, (d_1593_v17_).length(0))
            (d_1593_v17_)[index236_] = self.f38
            d_1594_v18_: _dafny.Seq
            d_1594_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhjfmmfl"))
            d_1595_v19_: D3
            d_1595_v19_ = D3_DC11((d_1594_v18_)[default__.safeIndex(d_1574_i0_, len(d_1594_v18_))])
            d_1596_v20_: _dafny.Map
            d_1596_v20_ = _dafny.Map({False: d_1595_v19_})
            d_1597_v21_: _dafny.Map
            d_1597_v21_ = _dafny.Map({d_1584_v8_: _dafny.SeqWithoutIsStrInference([D2_DC6(d_1575_v3_) for d_1598_i1_ in range(default__.abs(726))])})
            d_1599_v22_: D10
            d_1599_v22_ = D10_DC23(d_1597_v21_)
            d_1600_v23_: _dafny.Map
            d_1600_v23_ = _dafny.Map({d_1599_v22_: self.f38})
            index237_ = default__.safeIndex(687, (d_1593_v17_).length(0))
            rhs243_ = (d_1586_v10_ if self.f38 else d_1586_v10_)
            rhs244_ = d_1592_v16_
            rhs245_ = len(d_1596_v20_)
            rhs246_ = d_1574_i0_
            rhs247_ = ((d_1600_v23_)[d_1599_v22_] if (d_1599_v22_) in (d_1600_v23_) else self.f38)
            lhs223_ = globalState
            lhs224_ = globalState
            lhs225_ = d_1593_v17_
            lhs226_ = default__.safeIndex(687, (d_1593_v17_).length(0))
            d_1586_v10_ = rhs243_
            d_1592_v16_ = rhs244_
            lhs223_.f16 = rhs245_
            lhs224_.f12 = rhs246_
            lhs225_[lhs226_] = rhs247_
            index238_ = default__.safeIndex(687, (d_1593_v17_).length(0))
            (d_1593_v17_)[index238_] = (d_1593_v17_)[default__.safeIndex(687, (d_1593_v17_).length(0))]
        d_1605_v24_: str
        d_1605_v24_ = _dafny.CodePoint('l')
        d_1606_v25_: str
        d_1607_v26_: _dafny.Seq
        d_1608_v27_: bool
        d_1609_v28_: _dafny.Map
        out49_: str
        out50_: _dafny.Seq
        out51_: bool
        out52_: _dafny.Map
        out49_, out50_, out51_, out52_ = default__.m0(d_1605_v24_, globalState)
        d_1606_v25_ = out49_
        d_1607_v26_ = out50_
        d_1608_v27_ = out51_
        d_1609_v28_ = out52_
        d_1610_v29_: _dafny.Map
        d_1610_v29_ = _dafny.Map({d_1606_v25_: 369})
        rhs248_ = not((default__.safeDivisionInt(d_1571_v0_, (self).fm1(globalState))) >= (d_1571_v0_))
        rhs249_ = (default__.safeModuloInt(809, 846)) * (((d_1610_v29_)[default__.fm40(d_1571_v0_, (self).fm2(d_1608_v27_, d_1571_v0_, globalState), d_1608_v27_, globalState)] if (default__.fm40(d_1571_v0_, (self).fm2(d_1608_v27_, d_1571_v0_, globalState), d_1608_v27_, globalState)) in (d_1610_v29_) else d_1571_v0_))
        lhs227_ = globalState
        lhs227_.f24 = rhs248_
        r1 = rhs249_
        d_1611_v30_: _dafny.MultiSet
        d_1611_v30_ = _dafny.MultiSet([d_1571_v0_, d_1571_v0_])
        d_1612_v31_: _dafny.Map
        d_1612_v31_ = _dafny.Map({d_1611_v30_: self.f38})
        d_1613_v32_: _dafny.Map
        d_1613_v32_ = _dafny.Map({d_1571_v0_: len(d_1612_v31_)})
        d_1613_v32_ = (d_1613_v32_).set((0) - (d_1571_v0_), d_1571_v0_)
        (self).f38 = d_1608_v27_
        d_1614_v33_: D10
        d_1614_v33_ = D10_DC25()
        pat_let_tv40_ = d_1611_v30_
        pat_let_tv41_ = d_1611_v30_
        pat_let_tv42_ = d_1571_v0_
        pat_let_tv43_ = d_1611_v30_
        def lambda78_(source25_):
            if source25_.is_DC24:
                d_1615___mcc_h0_ = source25_.cf44
                d_1616___mcc_h1_ = source25_.cf45
                d_1617___mcc_h2_ = source25_.cf46
                d_1618_cf46_ = d_1617___mcc_h2_
                d_1619_cf45_ = d_1616___mcc_h1_
                d_1620_cf44_ = d_1615___mcc_h0_
                return pat_let_tv40_
            elif source25_.is_DC25:
                return (pat_let_tv41_).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iowvgu"))), default__.abs(pat_let_tv42_))
            elif True:
                d_1621___mcc_h3_ = source25_.cf43
                d_1622_cf43_ = d_1621___mcc_h3_
                return pat_let_tv43_

        d_1611_v30_ = lambda78_(d_1614_v33_)
        r0 = d_1608_v27_
        r1 = (d_1571_v0_) * (d_1571_v0_)
        r2 = d_1608_v27_
        return r0, r1, r2

    def m14(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        r2: _dafny.Map = _dafny.Map({})
        r3: bool = False
        d_1623_v0_: int
        d_1623_v0_ = -497
        d_1624_v1_: _dafny.Seq
        d_1624_v1_ = _dafny.SeqWithoutIsStrInference([d_1623_v0_])
        d_1625_v2_: C2
        nw251_ = C2()
        nw251_.ctor__(_dafny.SeqWithoutIsStrInference([d_1624_v1_]), not(p0), self.f38)
        d_1625_v2_ = nw251_
        d_1626_v3_: _dafny.Seq
        d_1626_v3_ = _dafny.SeqWithoutIsStrInference([(d_1625_v2_).fm2(p1, d_1623_v0_, globalState), p0])
        d_1627_v4_: D1
        d_1627_v4_ = D1_DC3((_dafny.SeqWithoutIsStrInference([p1])) + (d_1626_v3_))
        source26_ = d_1627_v4_
        if source26_.is_DC4:
            d_1628___mcc_h0_ = source26_.cf7
            d_1629___mcc_h1_ = source26_.cf8
            d_1630___mcc_h2_ = source26_.cf9
            d_1631___mcc_h3_ = source26_.cf10
            d_1632___mcc_h4_ = source26_.cf11
            d_1633_cf11_ = d_1632___mcc_h4_
            d_1634_cf10_ = d_1631___mcc_h3_
            d_1635_cf9_ = d_1630___mcc_h2_
            d_1636_cf8_ = d_1629___mcc_h1_
            d_1637_cf7_ = d_1628___mcc_h0_
            d_1638_v5_: _dafny.Map
            d_1638_v5_ = _dafny.Map({p1: (self).fm2(p1, d_1636_cf8_, globalState)})
            d_1639_v6_: _dafny.Set
            d_1639_v6_ = _dafny.Set({self.f38, ((d_1638_v5_)[self.f38] if (self.f38) in (d_1638_v5_) else p1), not(self.f38), self.f38, p0})
            d_1640_v7_: _dafny.Map
            d_1640_v7_ = _dafny.Map({(d_1639_v6_ if p1 else d_1639_v6_): p1})
            d_1641_v8_: _dafny.Set
            d_1641_v8_ = d_1639_v6_
            rhs250_ = ((d_1639_v6_) - (d_1639_v6_)).issubset((d_1639_v6_).intersection(_dafny.Set({True})))
            rhs251_ = d_1626_v3_
            rhs252_ = d_1626_v3_
            rhs253_ = (0) - (default__.safeModuloInt(len(d_1639_v6_), d_1636_cf8_))
            rhs254_ = ((d_1640_v7_)[d_1641_v8_] if (d_1641_v8_) in (d_1640_v7_) else self.f38)
            lhs228_ = globalState
            lhs229_ = globalState
            lhs228_.f22 = rhs250_
            d_1626_v3_ = rhs251_
            d_1626_v3_ = rhs252_
            d_1635_cf9_ = rhs253_
            lhs229_.f22 = rhs254_
            d_1642_v9_: _dafny.Seq
            d_1642_v9_ = _dafny.SeqWithoutIsStrInference([d_1639_v6_])
            d_1643_v10_: D14
            d_1643_v10_ = D14_DC31(_dafny.SeqWithoutIsStrInference([d_1639_v6_]))
            d_1644_v11_: _dafny.Map
            d_1644_v11_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([d_1639_v6_])})
            d_1645_v12_: _dafny.Array
            nw252_ = _dafny.Array(None, 18)
            nw252_[int(0)] = d_1642_v9_
            nw252_[int(1)] = (d_1642_v9_) + ((d_1642_v9_).set(default__.safeIndex(767, len(d_1642_v9_)), _dafny.Set({p0, True})))
            nw252_[int(2)] = d_1642_v9_
            nw252_[int(3)] = ((d_1642_v9_).set(default__.safeIndex(d_1636_cf8_, len(d_1642_v9_)), d_1639_v6_)) + (d_1642_v9_)
            nw252_[int(4)] = d_1642_v9_
            nw252_[int(5)] = (d_1642_v9_) + (d_1642_v9_)
            nw252_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_1639_v6_, d_1639_v6_, d_1639_v6_])) + (d_1642_v9_)
            nw252_[int(7)] = _dafny.SeqWithoutIsStrInference([d_1639_v6_])
            nw252_[int(8)] = (d_1642_v9_) + (_dafny.SeqWithoutIsStrInference([d_1639_v6_ for d_1646_i0_ in range(default__.abs(799))]))
            nw252_[int(9)] = d_1642_v9_
            nw252_[int(10)] = d_1642_v9_
            nw252_[int(11)] = (d_1642_v9_).set(default__.safeIndex(d_1623_v0_, len(d_1642_v9_)), d_1639_v6_)
            nw252_[int(12)] = ((d_1643_v10_).cf55) + (d_1642_v9_)
            nw252_[int(13)] = ((d_1644_v11_)[p0] if (p0) in (d_1644_v11_) else d_1642_v9_)
            nw252_[int(14)] = (d_1643_v10_).cf55
            nw252_[int(15)] = _dafny.SeqWithoutIsStrInference([d_1639_v6_ for d_1647_i1_ in range(default__.abs(785))])
            nw252_[int(16)] = (d_1642_v9_ if self.f38 else _dafny.SeqWithoutIsStrInference([d_1639_v6_ for d_1648_i2_ in range(default__.abs(104))]))
            nw252_[int(17)] = (d_1642_v9_) + (d_1642_v9_)
            d_1645_v12_ = nw252_
            index239_ = default__.safeIndex(645, (d_1645_v12_).length(0))
            (d_1645_v12_)[index239_] = d_1642_v9_
            d_1649_v13_: D3
            d_1649_v13_ = D3_DC11(_dafny.CodePoint('a'))
            d_1650_v14_: _dafny.Seq
            d_1650_v14_ = _dafny.SeqWithoutIsStrInference([default__.fm41(not(self.f38), d_1649_v13_, d_1635_cf9_, p0, globalState)])
            d_1651_v15_: C5
            nw253_ = C5()
            nw253_.ctor__((d_1626_v3_)[default__.safeIndex(d_1635_cf9_, len(d_1626_v3_))], True)
            d_1651_v15_ = nw253_
            d_1652_v16_: _dafny.Map
            d_1652_v16_ = _dafny.Map({d_1651_v15_: p1})
            d_1653_v17_: D15
            d_1653_v17_ = D15_DC33(d_1652_v16_)
            pat_let_tv44_ = d_1652_v16_
            index240_ = default__.safeIndex(645, (d_1645_v12_).length(0))
            def iife147_(_pat_let40_0):
                def iife148_(d_1654_dt__update__tmp_h0_):
                    def iife149_(_pat_let41_0):
                        def iife150_(d_1655_dt__update_hcf61_h0_):
                            return D15_DC33(d_1655_dt__update_hcf61_h0_)
                        return iife150_(_pat_let41_0)
                    return iife149_(pat_let_tv44_)
                return iife148_(_pat_let40_0)
            rhs255_ = d_1636_cf8_
            rhs256_ = (d_1642_v9_) + ((d_1650_v14_)[default__.safeIndex(len((d_1634_cf10_).set(default__.safeIndex((d_1625_v2_).fm4(D0_DC0(p1), d_1623_v0_, len((iife147_(d_1653_v17_)).cf61), globalState), len(d_1634_cf10_)), p2)), len(d_1650_v14_))])
            lhs230_ = globalState
            lhs231_ = d_1645_v12_
            lhs232_ = default__.safeIndex(645, (d_1645_v12_).length(0))
            lhs230_.f16 = rhs255_
            lhs231_[lhs232_] = rhs256_
            d_1656_v18_: _dafny.Map
            d_1656_v18_ = _dafny.Map({d_1623_v0_: d_1635_cf9_})
            if ((0) - ((d_1636_cf8_ if p0 else d_1623_v0_))) == (len((d_1656_v18_).set(d_1623_v0_, d_1636_cf8_))):
                d_1656_v18_ = (d_1656_v18_).set(567, d_1636_cf8_)
                d_1657_v19_: _dafny.Array
                def lambda79_(d_1658_cf9_):
                    def lambda80_(d_1659_i3_):
                        return default__.safeDivisionInt(d_1659_i3_, (0) - (d_1658_cf9_))

                    return lambda80_

                init44_ = lambda79_(d_1635_cf9_)
                nw254_ = _dafny.Array(None, 21)
                for i0_44_ in range(nw254_.length(0)):
                    nw254_[i0_44_] = init44_(i0_44_)
                d_1657_v19_ = nw254_
                index241_ = default__.safeIndex(658, (d_1657_v19_).length(0))
                (d_1657_v19_)[index241_] = (d_1651_v15_).fm1(globalState)
                index242_ = default__.safeIndex(658, (d_1657_v19_).length(0))
                (d_1657_v19_)[index242_] = (d_1624_v1_)[default__.safeIndex((d_1651_v15_).fm1(globalState), len(d_1624_v1_))]
                d_1660_v20_: C0
                nw255_ = C0()
                nw255_.ctor__((d_1636_cf8_) <= ((d_1657_v19_)[default__.safeIndex(658, (d_1657_v19_).length(0))]))
                d_1660_v20_ = nw255_
                rhs257_ = d_1660_v20_
                rhs258_ = d_1657_v19_
                rhs259_ = ((d_1625_v2_).f35)[default__.safeIndex(d_1623_v0_, len((d_1625_v2_).f35))]
                d_1660_v20_ = rhs257_
                d_1657_v19_ = rhs258_
                d_1624_v1_ = rhs259_
                d_1661_v21_: D15
                d_1661_v21_ = D15_DC34(d_1636_cf8_, p0, len(d_1633_cf11_))
                d_1661_v21_ = D15_DC34((d_1657_v19_)[default__.safeIndex(658, (d_1657_v19_).length(0))], False, default__.safeDivisionInt(d_1623_v0_, d_1635_cf9_))
                d_1662_v22_: C5
                nw256_ = C5()
                nw256_.ctor__(True, (True) or (True))
                d_1662_v22_ = nw256_
            elif True:
                d_1663_v23_: _dafny.Array
                nw257_ = _dafny.Array(False, 1)
                d_1663_v23_ = nw257_
                d_1663_v23_ = d_1663_v23_
                d_1664_v24_: _dafny.Array
                nw258_ = _dafny.Array(None, 9)
                nw258_[int(0)] = default__.safeDivisionInt(d_1635_cf9_, d_1636_cf8_)
                nw258_[int(1)] = (d_1635_cf9_ if p1 else d_1635_cf9_)
                nw258_[int(2)] = d_1635_cf9_
                nw258_[int(3)] = (0) - (default__.safeModuloInt(len(d_1626_v3_), d_1635_cf9_))
                nw258_[int(4)] = 836
                nw258_[int(5)] = d_1636_cf8_
                nw258_[int(6)] = d_1623_v0_
                nw258_[int(7)] = ((0) - (d_1636_cf8_)) * (len(d_1634_cf10_))
                nw258_[int(8)] = d_1635_cf9_
                d_1664_v24_ = nw258_
                index243_ = default__.safeIndex(563, (d_1664_v24_).length(0))
                (d_1664_v24_)[index243_] = len((d_1626_v3_).set(default__.safeIndex(d_1623_v0_, len(d_1626_v3_)), self.f38))
                index244_ = default__.safeIndex(563, (d_1664_v24_).length(0))
                (d_1664_v24_)[index244_] = d_1623_v0_
                d_1665_v25_: _dafny.Map
                d_1665_v25_ = _dafny.Map({d_1663_v23_: d_1635_cf9_})
                d_1666_v26_: _dafny.Map
                d_1666_v26_ = _dafny.Map({p0: len(d_1665_v25_)})
                (globalState).f12 = len((d_1666_v26_) | ((d_1666_v26_ if p1 else _dafny.Map({p0: (d_1664_v24_)[default__.safeIndex(563, (d_1664_v24_).length(0))]}))))
                d_1667_v27_: _dafny.MultiSet
                d_1667_v27_ = _dafny.MultiSet([p1, not(p1), self.f38])
                d_1668_v28_: _dafny.Set
                d_1668_v28_ = _dafny.Set({(d_1664_v24_)[default__.safeIndex(563, (d_1664_v24_).length(0))]})
                d_1669_v29_: D2
                d_1669_v29_ = D2_DC9((d_1623_v0_) - ((0) - (d_1623_v0_)), (0) - (d_1623_v0_), (d_1668_v28_).issubset(_dafny.Set({(d_1667_v27_).cardinality})))
                d_1669_v29_ = d_1669_v29_
                (globalState).f16 = d_1635_cf9_
            (globalState).f22 = p0
        elif source26_.is_DC5:
            d_1670___mcc_h5_ = source26_.cf12
            d_1671_cf12_ = d_1670___mcc_h5_
            r3 = p1
            d_1672_v30_: _dafny.Set
            d_1672_v30_ = _dafny.Set({d_1671_cf12_, d_1623_v0_})
            d_1673_v31_: _dafny.Array
            nw259_ = _dafny.Array(None, 2)
            nw259_[int(0)] = p2
            nw259_[int(1)] = p2
            d_1673_v31_ = nw259_
            d_1674_v32_: _dafny.Seq
            d_1674_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryndpso"))
            d_1675_v33_: D1
            d_1675_v33_ = D1_DC4(d_1673_v31_, d_1671_cf12_, d_1623_v0_, d_1674_v32_, d_1674_v32_)
            d_1676_v34_: _dafny.Map
            d_1676_v34_ = _dafny.Map({p0: p0})
            d_1677_v35_: _dafny.Map
            d_1677_v35_ = _dafny.Map({d_1671_cf12_: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vi"))))})
            d_1678_v36_: _dafny.Set
            d_1678_v36_ = _dafny.Set({self.f38, p0, True, p1})
            d_1679_v37_: _dafny.MultiSet
            d_1679_v37_ = _dafny.MultiSet([d_1623_v0_, d_1671_cf12_, d_1671_cf12_, d_1671_cf12_])
            d_1680_v38_: _dafny.Array
            nw260_ = _dafny.Array(None, 24)
            nw260_[int(0)] = (0) - (d_1623_v0_)
            nw260_[int(1)] = d_1671_cf12_
            nw260_[int(2)] = (d_1675_v33_).cf8
            nw260_[int(3)] = d_1671_cf12_
            nw260_[int(4)] = (0) - (d_1623_v0_)
            nw260_[int(5)] = d_1623_v0_
            nw260_[int(6)] = len((d_1676_v34_) | (d_1676_v34_))
            nw260_[int(7)] = d_1671_cf12_
            nw260_[int(8)] = d_1671_cf12_
            nw260_[int(9)] = d_1671_cf12_
            nw260_[int(10)] = (d_1625_v2_).fm1(globalState)
            nw260_[int(11)] = default__.safeModuloInt(d_1671_cf12_, d_1671_cf12_)
            nw260_[int(12)] = d_1623_v0_
            nw260_[int(13)] = (default__.fm42(d_1623_v0_, globalState)).cardinality
            nw260_[int(14)] = d_1671_cf12_
            nw260_[int(15)] = d_1623_v0_
            nw260_[int(16)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))
            nw260_[int(17)] = d_1671_cf12_
            nw260_[int(18)] = default__.safeModuloInt(len(d_1677_v35_), len(d_1678_v36_))
            nw260_[int(19)] = (_dafny.MultiSet([d_1671_cf12_, d_1623_v0_])).cardinality
            nw260_[int(20)] = d_1623_v0_
            nw260_[int(21)] = (d_1623_v0_) * ((0) - ((d_1679_v37_).cardinality))
            nw260_[int(22)] = d_1623_v0_
            nw260_[int(23)] = d_1623_v0_
            d_1680_v38_ = nw260_
            index245_ = default__.safeIndex(662, (d_1680_v38_).length(0))
            (d_1680_v38_)[index245_] = len((d_1624_v1_) + (d_1624_v1_))
            index246_ = default__.safeIndex(662, (d_1680_v38_).length(0))
            rhs260_ = (d_1672_v30_) - (d_1672_v30_)
            rhs261_ = default__.fm43(d_1623_v0_, 909, len((d_1626_v3_) + (d_1626_v3_)), _dafny.Map({p1: p2}), globalState)
            rhs262_ = (p0) and (not(p0))
            rhs263_ = d_1623_v0_
            rhs264_ = (d_1671_cf12_) + (d_1671_cf12_)
            lhs233_ = globalState
            lhs234_ = globalState
            lhs235_ = d_1680_v38_
            lhs236_ = default__.safeIndex(662, (d_1680_v38_).length(0))
            d_1672_v30_ = rhs260_
            d_1624_v1_ = rhs261_
            lhs233_.f15 = rhs262_
            lhs234_.f16 = rhs263_
            lhs235_[lhs236_] = rhs264_
            d_1681_v39_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.MultiSet({}), 8)
            d_1681_v39_ = nw261_
            d_1681_v39_ = d_1681_v39_
            d_1682_v40_: D0
            d_1682_v40_ = D0_DC0(self.f38)
            pat_let_tv45_ = p0
            def iife151_(_pat_let42_0):
                def iife152_(d_1683_dt__update__tmp_h1_):
                    def iife153_(_pat_let43_0):
                        def iife154_(d_1684_dt__update_hcf0_h0_):
                            return D0_DC0(d_1684_dt__update_hcf0_h0_)
                        return iife154_(_pat_let43_0)
                    return iife153_(pat_let_tv45_)
                return iife152_(_pat_let42_0)
            d_1671_cf12_ = (d_1625_v2_).fm4(iife151_(d_1682_v40_), default__.safeModuloInt(d_1623_v0_, d_1671_cf12_), (d_1671_cf12_) - (d_1671_cf12_), globalState)
        elif True:
            d_1685___mcc_h6_ = source26_.cf6
            d_1686_cf6_ = d_1685___mcc_h6_
            d_1687_v41_: T0
            nw262_ = C6()
            nw262_.ctor__(d_1623_v0_)
            d_1687_v41_ = nw262_
            d_1688_v42_: str
            d_1689_v43_: int
            out53_: str
            out54_: int
            out53_, out54_ = (d_1625_v2_).m2((d_1623_v0_) != (d_1623_v0_), d_1687_v41_, self.f38, globalState)
            d_1688_v42_ = out53_
            d_1689_v43_ = out54_
            if p1:
                (globalState).f15 = p1
                (globalState).f22 = (d_1687_v41_).fm2(p1, d_1623_v0_, globalState)
                d_1690_v44_: _dafny.Map
                d_1690_v44_ = _dafny.Map({p0: p1})
                d_1691_v45_: D13
                d_1691_v45_ = D13_DC29(d_1690_v44_)
                d_1692_v46_: _dafny.Set
                d_1692_v46_ = _dafny.Set({d_1691_v45_, d_1691_v45_})
                (globalState).f15 = (d_1692_v46_).isdisjoint(d_1692_v46_)
                d_1693_v47_: _dafny.MultiSet
                d_1693_v47_ = _dafny.MultiSet([p0])
                d_1694_v48_: _dafny.Seq
                d_1694_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_1690_v44_): False})])
                (globalState).f24 = (((0) - ((d_1693_v47_).cardinality)) == (len(d_1694_v48_))) == ((len(d_1686_cf6_)) == (505))
                d_1623_v0_ = d_1689_v43_
            elif True:
                d_1695_v49_: _dafny.Set
                d_1695_v49_ = _dafny.Set({d_1689_v43_})
                d_1696_v50_: _dafny.Map
                d_1696_v50_ = _dafny.Map({d_1623_v0_: p1})
                d_1697_v51_: _dafny.Seq
                d_1697_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bam"))
                d_1698_v52_: D5
                d_1698_v52_ = D5_DC13(_dafny.Map({p0: d_1626_v3_}))
                d_1699_v53_: D5
                d_1699_v53_ = D5_DC15(d_1698_v52_)
                d_1700_v54_: _dafny.Map
                d_1700_v54_ = _dafny.Map({d_1699_v53_: d_1689_v43_})
                d_1701_v55_: _dafny.Array
                nw263_ = _dafny.Array(None, 25)
                nw263_[int(0)] = d_1689_v43_
                nw263_[int(1)] = len(d_1695_v49_)
                nw263_[int(2)] = d_1623_v0_
                nw263_[int(3)] = d_1623_v0_
                nw263_[int(4)] = -377
                nw263_[int(5)] = d_1623_v0_
                nw263_[int(6)] = d_1689_v43_
                nw263_[int(7)] = d_1623_v0_
                nw263_[int(8)] = len(_dafny.Map({_dafny.MultiSet([((d_1696_v50_)[d_1689_v43_] if (d_1689_v43_) in (d_1696_v50_) else self.f38)]): d_1689_v43_}))
                nw263_[int(9)] = (d_1687_v41_).fm1(globalState)
                nw263_[int(10)] = (0) - (d_1623_v0_)
                nw263_[int(11)] = d_1689_v43_
                nw263_[int(12)] = d_1623_v0_
                nw263_[int(13)] = d_1623_v0_
                nw263_[int(14)] = (0) - (d_1689_v43_)
                nw263_[int(15)] = d_1689_v43_
                nw263_[int(16)] = len(d_1697_v51_)
                nw263_[int(17)] = (d_1687_v41_).fm1(globalState)
                nw263_[int(18)] = len(d_1700_v54_)
                nw263_[int(19)] = d_1689_v43_
                nw263_[int(20)] = (0) - (d_1623_v0_)
                nw263_[int(21)] = len(d_1697_v51_)
                nw263_[int(22)] = d_1623_v0_
                nw263_[int(23)] = d_1689_v43_
                nw263_[int(24)] = d_1689_v43_
                d_1701_v55_ = nw263_
                d_1702_v56_: _dafny.MultiSet
                d_1702_v56_ = _dafny.MultiSet([self.f38, False, False])
                d_1703_v57_: _dafny.Map
                d_1703_v57_ = _dafny.Map({d_1701_v55_: d_1702_v56_})
                d_1703_v57_ = (d_1703_v57_) | (d_1703_v57_)
                (globalState).f22 = p0
                (globalState).f6 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qx"))
                index247_ = default__.safeIndex(868, (d_1701_v55_).length(0))
                (d_1701_v55_)[index247_] = d_1623_v0_
                index248_ = default__.safeIndex(868, (d_1701_v55_).length(0))
                (d_1701_v55_)[index248_] = (D1_DC5((0) - (d_1689_v43_))).cf12
                (globalState).f15 = (d_1697_v51_) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1704_i4_ in range(default__.abs(657))]))
            d_1705_v58_: T2
            nw264_ = C1()
            nw264_.ctor__(d_1623_v0_, d_1623_v0_, p1, p0)
            d_1705_v58_ = nw264_
            d_1706_v59_: _dafny.MultiSet
            d_1706_v59_ = _dafny.MultiSet([d_1705_v58_, d_1705_v58_, d_1705_v58_, d_1705_v58_])
            d_1707_v60_: D15
            d_1707_v60_ = D15_DC34(-234, p1, d_1623_v0_)
            d_1708_v61_: _dafny.Map
            d_1708_v61_ = _dafny.Map({p0: len(_dafny.Map({944: (d_1707_v60_).cf63}))})
            d_1709_v62_: _dafny.Set
            d_1709_v62_ = _dafny.Set({True, p1, not((d_1705_v58_).f27), (d_1705_v58_).f26, (d_1705_v58_).f26})
            d_1710_v63_: _dafny.MultiSet
            d_1710_v63_ = _dafny.MultiSet([p2])
            d_1711_v64_: _dafny.Array
            nw265_ = _dafny.Array(None, 8)
            nw265_[int(0)] = (len(d_1624_v1_) if p0 else d_1689_v43_)
            nw265_[int(1)] = d_1689_v43_
            nw265_[int(2)] = (d_1706_v59_).cardinality
            nw265_[int(3)] = d_1689_v43_
            nw265_[int(4)] = len(d_1686_cf6_)
            nw265_[int(5)] = 405
            nw265_[int(6)] = default__.safeModuloInt(((d_1708_v61_)[(d_1705_v58_).f26] if ((d_1705_v58_).f26) in (d_1708_v61_) else len(d_1709_v62_)), d_1623_v0_)
            nw265_[int(7)] = ((default__.fm48(len(_dafny.SeqWithoutIsStrInference([d_1623_v0_])), globalState)) | (d_1710_v63_)).cardinality
            d_1711_v64_ = nw265_
            index249_ = default__.safeIndex(347, (d_1711_v64_).length(0))
            (d_1711_v64_)[index249_] = d_1689_v43_
            d_1712_v65_: D0
            d_1712_v65_ = D0_DC0((d_1686_cf6_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p2 for d_1713_i5_ in range(default__.abs(362))])), len(d_1686_cf6_))])
            index250_ = default__.safeIndex(347, (d_1711_v64_).length(0))
            rhs265_ = (0) - ((d_1625_v2_).fm4(d_1712_v65_, d_1689_v43_, d_1689_v43_, globalState))
            rhs266_ = len(d_1686_cf6_)
            lhs237_ = d_1711_v64_
            lhs238_ = default__.safeIndex(347, (d_1711_v64_).length(0))
            lhs237_[lhs238_] = rhs265_
            d_1689_v43_ = rhs266_
            d_1714_v66_: _dafny.MultiSet
            d_1714_v66_ = _dafny.MultiSet([-658])
            d_1715_v67_: _dafny.MultiSet
            d_1715_v67_ = _dafny.MultiSet([(d_1714_v66_) - (d_1714_v66_)])
            d_1715_v67_ = default__.fm49((d_1714_v66_).set(d_1689_v43_, default__.abs(-948)), len((d_1686_cf6_).set(default__.safeIndex(d_1623_v0_, len(d_1686_cf6_)), p1)), (d_1705_v58_).f27, globalState)
        hi14_ = d_1623_v0_
        for d_1716_i6_ in range(len(_dafny.Map({206: d_1623_v0_})), hi14_):
            d_1717_v68_: _dafny.MultiSet
            d_1717_v68_ = _dafny.MultiSet([d_1623_v0_])
            d_1718_v69_: _dafny.Set
            d_1718_v69_ = _dafny.Set({p1, self.f38})
            d_1719_v70_: _dafny.Map
            d_1719_v70_ = _dafny.Map({(0) - (d_1623_v0_): p1})
            d_1720_v71_: _dafny.MultiSet
            d_1720_v71_ = _dafny.MultiSet([_dafny.Map({(d_1717_v68_).cardinality: p1}), _dafny.Map({len(d_1718_v69_): p0}), d_1719_v70_, d_1719_v70_])
            d_1721_v72_: _dafny.Map
            d_1721_v72_ = _dafny.Map({not(self.f38): d_1720_v71_})
            r3 = not(((d_1720_v71_).intersection(((d_1721_v72_)[p0] if (p0) in (d_1721_v72_) else d_1720_v71_))).ispropersubset(_dafny.MultiSet([_dafny.Map({d_1623_v0_: p1})])))
            d_1722_v73_: _dafny.MultiSet
            d_1722_v73_ = _dafny.MultiSet([True])
            d_1723_v74_: _dafny.Map
            d_1723_v74_ = _dafny.Map({self.f38: p0})
            d_1724_v75_: _dafny.Array
            nw266_ = _dafny.Array(None, 21)
            nw266_[int(0)] = p0
            nw266_[int(1)] = p0
            nw266_[int(2)] = p1
            nw266_[int(3)] = True
            nw266_[int(4)] = p0
            nw266_[int(5)] = p0
            nw266_[int(6)] = (d_1626_v3_)[default__.safeIndex(((d_1722_v73_)[self.f38] if (self.f38) in (d_1722_v73_) else d_1623_v0_), len(d_1626_v3_))]
            nw266_[int(7)] = self.f38
            nw266_[int(8)] = self.f38
            nw266_[int(9)] = p0
            nw266_[int(10)] = not(True)
            nw266_[int(11)] = p1
            nw266_[int(12)] = self.f38
            nw266_[int(13)] = self.f38
            nw266_[int(14)] = (d_1625_v2_).fm2(((d_1723_v74_)[self.f38] if (self.f38) in (d_1723_v74_) else self.f38), len(_dafny.SeqWithoutIsStrInference([p2 for d_1725_i7_ in range(default__.abs(-511))])), globalState)
            nw266_[int(15)] = p1
            nw266_[int(16)] = p1
            nw266_[int(17)] = p1
            nw266_[int(18)] = p0
            nw266_[int(19)] = True
            nw266_[int(20)] = not((D2_DC7(p1, True, d_1623_v0_, d_1623_v0_)).cf14)
            d_1724_v75_ = nw266_
            d_1726_v76_: D14
            d_1726_v76_ = D14_DC32(d_1722_v73_, p1, p0, p1, len(d_1718_v69_))
            d_1727_v77_: _dafny.Map
            d_1727_v77_ = _dafny.Map({d_1726_v76_: self.f38})
            index251_ = default__.safeIndex(329, (d_1724_v75_).length(0))
            (d_1724_v75_)[index251_] = (d_1625_v2_).fm17(d_1623_v0_, d_1716_i6_, len(d_1727_v77_), default__.fm50(globalState), globalState)
            index252_ = default__.safeIndex(329, (d_1724_v75_).length(0))
            rhs267_ = p1
            rhs268_ = d_1716_i6_
            lhs239_ = d_1724_v75_
            lhs240_ = default__.safeIndex(329, (d_1724_v75_).length(0))
            lhs241_ = globalState
            lhs239_[lhs240_] = rhs267_
            lhs241_.f16 = rhs268_
            d_1728_v78_: T0
            nw267_ = C6()
            nw267_.ctor__((d_1717_v68_).cardinality)
            d_1728_v78_ = nw267_
            d_1729_v79_: str
            d_1730_v80_: int
            out55_: str
            out56_: int
            out55_, out56_ = (d_1625_v2_).m2(p0, d_1728_v78_, not((self.f38 if p1 else p0)), globalState)
            d_1729_v79_ = out55_
            d_1730_v80_ = out56_
            d_1731_v81_: _dafny.Map
            d_1731_v81_ = _dafny.Map({p0: d_1623_v0_})
            d_1732_v82_: D7
            d_1732_v82_ = D7_DC19(d_1623_v0_, _dafny.MultiSet([p1, p0]), (d_1724_v75_)[default__.safeIndex(329, (d_1724_v75_).length(0))], d_1731_v81_)
            d_1733_v83_: _dafny.Map
            d_1733_v83_ = _dafny.Map({_dafny.Map({p1: self.f38}): (d_1732_v82_).cf36})
            index253_ = default__.safeIndex(329, (d_1724_v75_).length(0))
            (d_1724_v75_)[index253_] = ((d_1719_v70_)[default__.safeModuloInt(641, (0) - (len(d_1733_v83_)))] if (default__.safeModuloInt(641, (0) - (len(d_1733_v83_)))) in (d_1719_v70_) else p0)
        (globalState).f16 = d_1623_v0_
        d_1734_v84_: _dafny.Array
        nw268_ = _dafny.Array(False, 6)
        d_1734_v84_ = nw268_
        d_1734_v84_ = d_1734_v84_
        d_1735_v85_: _dafny.Seq
        d_1735_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcjrgaf"))
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tixwcfv"))) <= (d_1735_v85_):
            d_1736_v86_: _dafny.Array
            nw269_ = _dafny.Array(int(0), 27)
            d_1736_v86_ = nw269_
            d_1737_v87_: D3
            d_1737_v87_ = D3_DC10(d_1736_v86_)
            d_1738_v88_: _dafny.Map
            d_1738_v88_ = _dafny.Map({d_1737_v87_: len(_dafny.SeqWithoutIsStrInference([d_1623_v0_]))})
            d_1739_v89_: C4
            nw270_ = C4()
            nw270_.ctor__(d_1738_v88_, p1, self.f38)
            d_1739_v89_ = nw270_
            d_1740_v90_: _dafny.Seq
            d_1740_v90_ = _dafny.SeqWithoutIsStrInference([d_1735_v85_, d_1735_v85_, d_1735_v85_, d_1735_v85_])
            (d_1625_v2_).m1((True if not(not(False)) else self.f38), d_1740_v90_, (d_1626_v3_) + (d_1626_v3_), globalState)
            d_1741_v91_: _dafny.Map
            d_1741_v91_ = _dafny.Map({_dafny.CodePoint('w'): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1742_i8_ in range(default__.abs(971))])})
            rhs269_ = d_1741_v91_
            rhs270_ = True
            rhs271_ = d_1623_v0_
            rhs272_ = (d_1735_v85_) + (d_1735_v85_)
            lhs242_ = self
            lhs243_ = globalState
            lhs244_ = globalState
            d_1741_v91_ = rhs269_
            lhs242_.f38 = rhs270_
            lhs243_.f12 = rhs271_
            lhs244_.f6 = rhs272_
            r0 = False
            d_1743_v92_: _dafny.Array
            nw271_ = _dafny.Array(None, 9)
            nw271_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnu"))
            nw271_[int(1)] = (d_1735_v85_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxsmk")))
            nw271_[int(2)] = (d_1735_v85_) + (d_1735_v85_)
            nw271_[int(3)] = d_1735_v85_
            nw271_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            nw271_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fws"))
            nw271_[int(6)] = d_1735_v85_
            nw271_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cobamhx"))
            nw271_[int(8)] = d_1735_v85_
            d_1743_v92_ = nw271_
            d_1743_v92_ = d_1743_v92_
        elif True:
            d_1744_v93_: _dafny.Map
            d_1744_v93_ = _dafny.Map({d_1623_v0_: 223})
            d_1745_v94_: _dafny.Map
            d_1745_v94_ = _dafny.Map({d_1734_v84_: (_dafny.MultiSet([False, self.f38, self.f38, p1, p1])).cardinality})
            d_1746_v95_: _dafny.Set
            d_1746_v95_ = _dafny.Set({self.f38, not(True), p0, True, False})
            d_1747_v96_: _dafny.Array
            nw272_ = _dafny.Array(None, 18)
            nw272_[int(0)] = d_1623_v0_
            nw272_[int(1)] = (len(d_1744_v93_)) - (647)
            nw272_[int(2)] = (0) - ((0) - (len(d_1735_v85_)))
            nw272_[int(3)] = d_1623_v0_
            nw272_[int(4)] = d_1623_v0_
            nw272_[int(5)] = default__.safeDivisionInt(d_1623_v0_, d_1623_v0_)
            nw272_[int(6)] = 834
            nw272_[int(7)] = d_1623_v0_
            nw272_[int(8)] = (d_1623_v0_) + (d_1623_v0_)
            nw272_[int(9)] = ((d_1745_v94_)[d_1734_v84_] if (d_1734_v84_) in (d_1745_v94_) else d_1623_v0_)
            nw272_[int(10)] = d_1623_v0_
            nw272_[int(11)] = ((D16_DC38(d_1623_v0_, d_1623_v0_, self.f38)).cf70) - ((0) - (d_1623_v0_))
            nw272_[int(12)] = 174
            nw272_[int(13)] = d_1623_v0_
            nw272_[int(14)] = len(d_1746_v95_)
            nw272_[int(15)] = d_1623_v0_
            nw272_[int(16)] = d_1623_v0_
            nw272_[int(17)] = (d_1623_v0_ if not(self.f38) else d_1623_v0_)
            d_1747_v96_ = nw272_
            index254_ = default__.safeIndex(907, (d_1747_v96_).length(0))
            (d_1747_v96_)[index254_] = default__.safeDivisionInt((0) - (d_1623_v0_), d_1623_v0_)
            d_1748_v97_: D0
            d_1748_v97_ = D0_DC0(p0)
            d_1749_v98_: D0
            d_1749_v98_ = D0_DC2(d_1748_v97_)
            d_1750_v99_: _dafny.Seq
            d_1750_v99_ = _dafny.SeqWithoutIsStrInference([d_1749_v98_])
            index255_ = default__.safeIndex(907, (d_1747_v96_).length(0))
            rhs273_ = d_1623_v0_
            rhs274_ = default__.safeModuloInt(default__.safeDivisionInt((0) - (d_1623_v0_), -14), d_1623_v0_)
            rhs275_ = ((0) - (len((d_1735_v85_) + (d_1735_v85_)))) * ((d_1625_v2_).fm1(globalState))
            rhs276_ = _dafny.SeqWithoutIsStrInference([d_1749_v98_ for d_1751_i9_ in range(default__.abs(63))])
            rhs277_ = d_1623_v0_
            lhs245_ = globalState
            lhs246_ = d_1747_v96_
            lhs247_ = default__.safeIndex(907, (d_1747_v96_).length(0))
            lhs248_ = globalState
            d_1623_v0_ = rhs273_
            lhs245_.f12 = rhs274_
            lhs246_[lhs247_] = rhs275_
            d_1750_v99_ = rhs276_
            lhs248_.f12 = rhs277_
            (globalState).f22 = p1
            d_1752_v100_: _dafny.Array
            nw273_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1752_v100_ = nw273_
            index256_ = default__.safeIndex(144, (d_1752_v100_).length(0))
            (d_1752_v100_)[index256_] = d_1747_v96_
            index257_ = default__.safeIndex(144, (d_1752_v100_).length(0))
            (d_1752_v100_)[index257_] = d_1747_v96_
            d_1746_v95_ = (d_1746_v95_) | (_dafny.Set({p1}))
            d_1753_v101_: _dafny.Map
            d_1753_v101_ = _dafny.Map({False: (d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))]})
            if ((d_1753_v101_) == (d_1753_v101_) if True else self.f38):
                (globalState).f24 = ((self.f38) and (False)) or (p1)
                (globalState).f15 = self.f38
                d_1754_v102_: _dafny.Map
                d_1754_v102_ = _dafny.Map({p1: p0})
                d_1755_v103_: _dafny.Map
                d_1755_v103_ = _dafny.Map({(d_1752_v100_)[default__.safeIndex(144, (d_1752_v100_).length(0))]: p0})
                d_1754_v102_ = (d_1754_v102_).set(((d_1755_v103_)[d_1747_v96_] if (d_1747_v96_) in (d_1755_v103_) else p0), self.f38)
                d_1734_v84_ = d_1734_v84_
                index258_ = default__.safeIndex(907, (d_1747_v96_).length(0))
                (d_1747_v96_)[index258_] = ((d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))] if True else (d_1623_v0_) - (d_1623_v0_))
            elif True:
                d_1756_v104_: _dafny.Set
                d_1756_v104_ = _dafny.Set({((d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))]) * ((d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))]), d_1623_v0_, 177, 986, (d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))]})
                d_1757_v105_: _dafny.Seq
                d_1757_v105_ = _dafny.SeqWithoutIsStrInference([d_1735_v85_, d_1735_v85_, d_1735_v85_])
                d_1758_v106_: _dafny.Map
                d_1758_v106_ = _dafny.Map({_dafny.MultiSet((d_1757_v105_) + (d_1757_v105_)): _dafny.Set({(d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))], d_1623_v0_})})
                d_1759_v107_: _dafny.MultiSet
                d_1759_v107_ = _dafny.MultiSet([d_1735_v85_])
                d_1756_v104_ = ((d_1758_v106_)[(d_1759_v107_).set(d_1735_v85_, default__.abs(d_1623_v0_))] if ((d_1759_v107_).set(d_1735_v85_, default__.abs(d_1623_v0_))) in (d_1758_v106_) else d_1756_v104_)
                d_1760_v108_: D17
                d_1760_v108_ = D17_DC40((d_1625_v2_).f35)
                d_1761_v109_: _dafny.Map
                d_1761_v109_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([(d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))], d_1623_v0_])})
                d_1762_v110_: _dafny.MultiSet
                d_1762_v110_ = _dafny.MultiSet([d_1624_v1_])
                d_1763_v111_: _dafny.Seq
                d_1763_v111_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(((d_1760_v108_).cf73).set(default__.safeIndex((d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))], len((d_1760_v108_).cf73)), _dafny.SeqWithoutIsStrInference([len(((d_1761_v109_)[p2] if (p2) in (d_1761_v109_) else d_1624_v1_))]))), d_1762_v110_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((d_1625_v2_).f35)[default__.safeIndex(d_1623_v0_, len((d_1625_v2_).f35))] for d_1764_i10_ in range(default__.abs(777))])), d_1762_v110_, d_1762_v110_])
                d_1765_v112_: C3
                nw274_ = C3()
                nw274_.ctor__((d_1763_v111_)[default__.safeIndex(d_1623_v0_, len(d_1763_v111_))], self.f38, p1)
                d_1765_v112_ = nw274_
                d_1766_v113_: _dafny.Seq
                d_1766_v113_ = _dafny.SeqWithoutIsStrInference([d_1744_v93_, d_1744_v93_])
                d_1766_v113_ = d_1766_v113_
                (globalState).f16 = default__.safeDivisionInt(d_1623_v0_, (d_1747_v96_)[default__.safeIndex(907, (d_1747_v96_).length(0))])
                d_1767_v114_: C2
                nw275_ = C2()
                nw275_.ctor__((d_1625_v2_).f35, False, False)
                d_1767_v114_ = nw275_
        r0 = p1
        d_1768_v115_: _dafny.Map
        d_1768_v115_ = _dafny.Map({self.f38: d_1623_v0_})
        d_1769_v116_: _dafny.MultiSet
        d_1769_v116_ = _dafny.MultiSet([d_1768_v115_, d_1768_v115_])
        d_1770_v117_: D1
        d_1770_v117_ = D1_DC5(226)
        r1 = ((d_1769_v116_).set(d_1768_v115_, default__.abs((d_1770_v117_).cf12))).intersection((_dafny.MultiSet([d_1768_v115_])) - (_dafny.MultiSet([d_1768_v115_, d_1768_v115_])))
        d_1771_v118_: _dafny.Map
        d_1771_v118_ = _dafny.Map({self.f38: p2})
        d_1772_v119_: _dafny.Map
        d_1772_v119_ = _dafny.Map({(d_1626_v3_) + (d_1626_v3_): len((_dafny.Map({p1: p2}) if p1 else d_1771_v118_))})
        r2 = d_1772_v119_
        r3 = p1
        return r0, r1, r2, r3


class C8(T3):
    def  __init__(self):
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f29):
        (self)._f29 = f29

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: T0 = None
        (globalState).f16 = p0
        hi15_ = p0
        for d_1773_i0_ in range(p0, hi15_):
            d_1774_v0_: str
            d_1774_v0_ = _dafny.CodePoint('y')
            if (default__.safeDivisionInt(p3, -589)) == ((0) - (len((p2).set(default__.safeIndex(p1, len(p2)), d_1774_v0_)))):
                d_1775_v1_: bool
                d_1775_v1_ = False
                d_1776_v3_: _dafny.Map
                d_1776_v3_ = _dafny.Map({p0: d_1773_i0_})
                d_1777_v4_: _dafny.Set
                def iife155_():
                    coll67_ = _dafny.Map()
                    compr_67_: int
                    for compr_67_ in (d_1776_v3_).keys.Elements:
                        d_1778_v2_: int = compr_67_
                        if (d_1778_v2_) in (d_1776_v3_):
                            coll67_[default__.safeDivisionInt(d_1778_v2_, d_1773_i0_)] = d_1773_i0_
                    return _dafny.Map(coll67_)
                d_1777_v4_ = _dafny.Set({-207, len(_dafny.Map({p0: len(iife155_()
                )})), p3})
                d_1779_v5_: _dafny.Map
                d_1779_v5_ = _dafny.Map({d_1775_v1_: d_1777_v4_})
                d_1780_v6_: _dafny.Map
                d_1780_v6_ = _dafny.Map({(d_1773_i0_) * (p0): (d_1779_v5_) | (d_1779_v5_)})
                d_1780_v6_ = (d_1780_v6_).set(len(p2), (d_1779_v5_) | (d_1779_v5_))
                d_1776_v3_ = (d_1776_v3_).set(default__.fm39(globalState), p1)
                d_1781_v7_: _dafny.Array
                def lambda81_(d_1782_p1_):
                    def lambda82_(d_1783_i1_):
                        return (d_1783_i1_) + (d_1782_p1_)

                    return lambda82_

                init45_ = lambda81_(p1)
                nw276_ = _dafny.Array(None, 11)
                for i0_45_ in range(nw276_.length(0)):
                    nw276_[i0_45_] = init45_(i0_45_)
                d_1781_v7_ = nw276_
                index259_ = default__.safeIndex(905, (d_1781_v7_).length(0))
                (d_1781_v7_)[index259_] = (0) - (p3)
                d_1784_v8_: _dafny.Array
                def lambda83_(d_1785_v1_):
                    def lambda84_(d_1786_i2_):
                        return d_1785_v1_

                    return lambda84_

                init46_ = lambda83_(d_1775_v1_)
                nw277_ = _dafny.Array(None, 7)
                for i0_46_ in range(nw277_.length(0)):
                    nw277_[i0_46_] = init46_(i0_46_)
                d_1784_v8_ = nw277_
                index260_ = default__.safeIndex(931, (d_1784_v8_).length(0))
                (d_1784_v8_)[index260_] = not((d_1775_v1_) and (d_1775_v1_))
                d_1787_v9_: _dafny.Map
                d_1787_v9_ = _dafny.Map({False: True})
                d_1788_v10_: D13
                d_1788_v10_ = D13_DC29(d_1787_v9_)
                d_1789_v11_: _dafny.Seq
                d_1789_v11_ = _dafny.SeqWithoutIsStrInference([d_1773_i0_, len(p2), d_1773_i0_, p3, p0])
                index261_ = default__.safeIndex(905, (d_1781_v7_).length(0))
                index262_ = default__.safeIndex(931, (d_1784_v8_).length(0))
                rhs278_ = (109) * ((len((d_1788_v10_).cf50)) * (d_1773_i0_))
                rhs279_ = d_1775_v1_
                rhs280_ = (p0) * (p0)
                rhs281_ = ((p3) in (d_1789_v11_) if not(d_1775_v1_) else (d_1775_v1_) and (d_1775_v1_))
                rhs282_ = ((d_1773_i0_) + (p1)) + ((578 if d_1775_v1_ else d_1773_i0_))
                lhs249_ = d_1781_v7_
                lhs250_ = default__.safeIndex(905, (d_1781_v7_).length(0))
                lhs251_ = d_1784_v8_
                lhs252_ = default__.safeIndex(931, (d_1784_v8_).length(0))
                lhs253_ = globalState
                lhs254_ = globalState
                lhs255_ = globalState
                lhs249_[lhs250_] = rhs278_
                lhs251_[lhs252_] = rhs279_
                lhs253_.f16 = rhs280_
                lhs254_.f15 = rhs281_
                lhs255_.f12 = rhs282_
                (globalState).f22 = (d_1784_v8_)[default__.safeIndex(931, (d_1784_v8_).length(0))]
                d_1790_v12_: C0
                nw278_ = C0()
                nw278_.ctor__(not((d_1784_v8_)[default__.safeIndex(931, (d_1784_v8_).length(0))]))
                d_1790_v12_ = nw278_
            elif True:
                d_1791_v13_: bool
                d_1791_v13_ = True
                (globalState).f22 = not((d_1791_v13_) and (d_1791_v13_))
                (globalState).f16 = p1
                d_1792_v14_: _dafny.Set
                d_1792_v14_ = _dafny.Set({d_1774_v0_})
                d_1793_v15_: C1
                nw279_ = C1()
                nw279_.ctor__(p3, d_1773_i0_, d_1791_v13_, (len(p2)) == (len(d_1792_v14_)))
                d_1793_v15_ = nw279_
                d_1794_v16_: _dafny.Array
                nw280_ = _dafny.Array(False, 21)
                d_1794_v16_ = nw280_
                index263_ = default__.safeIndex(829, (d_1794_v16_).length(0))
                (d_1794_v16_)[index263_] = True
                index264_ = default__.safeIndex(829, (d_1794_v16_).length(0))
                (d_1794_v16_)[index264_] = not(d_1791_v13_)
                (globalState).f16 = d_1793_v15_.f37
            (globalState).f12 = d_1773_i0_
            d_1795_v17_: _dafny.Seq
            d_1795_v17_ = _dafny.SeqWithoutIsStrInference([d_1773_i0_, default__.fm39(globalState)])
            (globalState).f16 = len(d_1795_v17_)
            (globalState).f19 = d_1774_v0_
        d_1796_v18_: _dafny.Array
        def lambda85_(d_1797_p1_):
            def lambda86_(d_1798_i3_):
                return (d_1798_i3_) * (d_1797_p1_)

            return lambda86_

        init47_ = lambda85_(p1)
        nw281_ = _dafny.Array(None, 27)
        for i0_47_ in range(nw281_.length(0)):
            nw281_[i0_47_] = init47_(i0_47_)
        d_1796_v18_ = nw281_
        d_1799_v19_: _dafny.Seq
        d_1799_v19_ = _dafny.SeqWithoutIsStrInference([d_1796_v18_, d_1796_v18_, d_1796_v18_])
        d_1800_v20_: D3
        d_1800_v20_ = D3_DC10((d_1799_v19_)[default__.safeIndex(p0, len(d_1799_v19_))])
        pat_let_tv46_ = d_1796_v18_
        d_1801_v21_: _dafny.Map
        def iife156_(_pat_let44_0):
            def iife157_(d_1802_dt__update__tmp_h0_):
                def iife158_(_pat_let45_0):
                    def iife159_(d_1803_dt__update_hcf21_h0_):
                        return D3_DC10(d_1803_dt__update_hcf21_h0_)
                    return iife159_(_pat_let45_0)
                return iife158_(pat_let_tv46_)
            return iife157_(_pat_let44_0)
        d_1801_v21_ = _dafny.Map({iife156_(d_1800_v20_): p1})
        d_1804_v22_: _dafny.Set
        d_1804_v22_ = _dafny.Set({d_1796_v18_})
        d_1805_v23_: bool
        d_1805_v23_ = True
        d_1806_v24_: C4
        nw282_ = C4()
        nw282_.ctor__(d_1801_v21_, (d_1804_v22_).isdisjoint(_dafny.Set({d_1796_v18_, d_1796_v18_, d_1796_v18_, d_1796_v18_})), d_1805_v23_)
        d_1806_v24_ = nw282_
        d_1807_v25_: str
        d_1807_v25_ = _dafny.CodePoint('f')
        d_1808_v26_: str
        d_1809_v27_: _dafny.Seq
        d_1810_v28_: bool
        d_1811_v29_: _dafny.Map
        out57_: str
        out58_: _dafny.Seq
        out59_: bool
        out60_: _dafny.Map
        out57_, out58_, out59_, out60_ = default__.m0(d_1807_v25_, globalState)
        d_1808_v26_ = out57_
        d_1809_v27_ = out58_
        d_1810_v28_ = out59_
        d_1811_v29_ = out60_
        d_1812_v30_: _dafny.Map
        d_1812_v30_ = _dafny.Map({d_1805_v23_: p1})
        hi16_ = len(d_1812_v30_)
        for d_1813_i4_ in range((len(p2)) - (p1), hi16_):
            (globalState).f16 = d_1813_i4_
            (globalState).f16 = (d_1806_v24_).fm1(globalState)
            d_1809_v27_ = ((_dafny.SeqWithoutIsStrInference([d_1805_v23_])) + (d_1809_v27_)) + (d_1809_v27_)
            d_1814_v31_: _dafny.Seq
            d_1814_v31_ = _dafny.SeqWithoutIsStrInference([p3])
            d_1815_v32_: _dafny.MultiSet
            d_1815_v32_ = _dafny.MultiSet([d_1814_v31_])
            d_1816_v33_: T1
            nw283_ = C3()
            nw283_.ctor__(d_1815_v32_, d_1810_v28_, d_1805_v23_)
            d_1816_v33_ = nw283_
            d_1817_v34_: _dafny.Map
            d_1817_v34_ = _dafny.Map({d_1816_v33_: (self).f29})
            d_1818_v35_: _dafny.Map
            d_1818_v35_ = _dafny.Map({d_1817_v34_: (p1) < (76)})
            d_1818_v35_ = (d_1818_v35_).set((d_1817_v34_).set(d_1816_v33_, (self).f29), (d_1809_v27_)[default__.safeIndex((d_1816_v33_).fm1(globalState), len(d_1809_v27_))])
        d_1819_v36_: _dafny.Set
        d_1819_v36_ = _dafny.Set({d_1810_v28_})
        rhs283_ = ((d_1819_v36_).intersection(d_1819_v36_)).ispropersubset(d_1819_v36_)
        rhs284_ = p0
        lhs256_ = globalState
        r0 = rhs283_
        lhs256_.f16 = rhs284_
        r0 = not (d_1805_v23_) or (d_1805_v23_)
        r1 = not((d_1805_v23_) or (d_1805_v23_))
        r2 = (d_1799_v19_)[default__.safeIndex(p0, len(d_1799_v19_))]
        d_1820_v38_: T0
        nw284_ = C7()
        def iife160_():
            coll68_ = _dafny.Map()
            compr_68_: int
            for compr_68_ in _dafny.IntegerRange(-329, 305):
                d_1821_v37_: int = compr_68_
                if ((-329) <= (d_1821_v37_)) and ((d_1821_v37_) < (305)):
                    coll68_[(d_1821_v37_) + ((_dafny.MultiSet(d_1809_v27_)).cardinality)] = d_1812_v30_
            return _dafny.Map(coll68_)
        nw284_.ctor__(d_1805_v23_, iife160_()
        )
        d_1820_v38_ = nw284_
        r3 = d_1820_v38_
        return r0, r1, r2, r3

    def m4(self, globalState):
        r0: bool = False
        r1: bool = False
        d_1822_v0_: int
        d_1822_v0_ = 290
        d_1823_v1_: _dafny.Seq
        d_1823_v1_ = _dafny.SeqWithoutIsStrInference([d_1822_v0_])
        d_1824_v2_: C6
        nw285_ = C6()
        nw285_.ctor__(d_1822_v0_)
        d_1824_v2_ = nw285_
        d_1825_v3_: _dafny.Map
        d_1825_v3_ = _dafny.Map({d_1824_v2_: d_1823_v1_})
        d_1826_v4_: _dafny.MultiSet
        d_1826_v4_ = _dafny.MultiSet([d_1823_v1_, ((d_1825_v3_)[d_1824_v2_] if (d_1824_v2_) in (d_1825_v3_) else d_1823_v1_)])
        d_1827_v5_: bool
        d_1827_v5_ = False
        d_1828_v6_: C3
        nw286_ = C3()
        nw286_.ctor__(d_1826_v4_, False, d_1827_v5_)
        d_1828_v6_ = nw286_
        d_1829_v7_: _dafny.MultiSet
        d_1829_v7_ = _dafny.MultiSet([d_1828_v6_, d_1828_v6_])
        d_1830_v8_: _dafny.Seq
        d_1830_v8_ = _dafny.SeqWithoutIsStrInference([d_1827_v5_, False])
        d_1831_v9_: D6
        d_1831_v9_ = D6_DC17(((d_1829_v7_)[d_1828_v6_] if (d_1828_v6_) in (d_1829_v7_) else d_1824_v2_.f40), d_1824_v2_.f40, 51, (_dafny.MultiSet((d_1830_v8_).set(default__.safeIndex(d_1822_v0_, len(d_1830_v8_)), d_1827_v5_))).cardinality, d_1822_v0_)
        d_1832_v10_: str
        d_1832_v10_ = _dafny.CodePoint('v')
        d_1833_v11_: _dafny.MultiSet
        d_1833_v11_ = _dafny.MultiSet([d_1832_v10_])
        d_1834_v12_: _dafny.Seq
        d_1834_v12_ = _dafny.SeqWithoutIsStrInference([d_1833_v11_])
        pat_let_tv47_ = d_1828_v6_
        pat_let_tv48_ = d_1834_v12_
        pat_let_tv49_ = d_1832_v10_
        pat_let_tv50_ = globalState
        pat_let_tv51_ = d_1822_v0_
        pat_let_tv52_ = d_1824_v2_
        pat_let_tv53_ = d_1824_v2_
        pat_let_tv54_ = d_1822_v0_
        d_1835_v13_: _dafny.Array
        nw287_ = _dafny.Array(None, 23)
        nw287_[int(0)] = d_1831_v9_
        nw287_[int(1)] = d_1831_v9_
        nw287_[int(2)] = d_1831_v9_
        nw287_[int(3)] = d_1831_v9_
        nw287_[int(4)] = d_1831_v9_
        nw287_[int(5)] = d_1831_v9_
        nw287_[int(6)] = D6_DC17(d_1822_v0_, d_1822_v0_, (0) - (d_1824_v2_.f40), d_1824_v2_.f40, d_1824_v2_.f40)
        nw287_[int(7)] = d_1831_v9_
        nw287_[int(8)] = d_1831_v9_
        nw287_[int(9)] = D6_DC17(195, d_1822_v0_, d_1824_v2_.f40, d_1822_v0_, d_1822_v0_)
        nw287_[int(10)] = d_1831_v9_
        nw287_[int(11)] = d_1831_v9_
        def iife161_(_pat_let46_0):
            def iife162_(d_1836_dt__update__tmp_h0_):
                def iife163_(_pat_let47_0):
                    def iife164_(d_1837_dt__update_hcf30_h0_):
                        def iife165_(_pat_let48_0):
                            def iife166_(d_1838_dt__update_hcf31_h0_):
                                return D6_DC17(d_1837_dt__update_hcf30_h0_, d_1838_dt__update_hcf31_h0_, (d_1836_dt__update__tmp_h0_).cf32, (d_1836_dt__update__tmp_h0_).cf33, (d_1836_dt__update__tmp_h0_).cf34)
                            return iife166_(_pat_let48_0)
                        return iife165_(87)
                    return iife164_(_pat_let47_0)
                return iife163_((pat_let_tv47_).fm13(pat_let_tv48_, pat_let_tv49_, ((self).f29).cardinality, pat_let_tv50_))
            return iife162_(_pat_let46_0)
        nw287_[int(12)] = iife161_(d_1831_v9_)
        nw287_[int(13)] = d_1831_v9_
        nw287_[int(14)] = d_1831_v9_
        nw287_[int(15)] = d_1831_v9_
        def iife167_(_pat_let49_0):
            def iife168_(d_1839_dt__update__tmp_h1_):
                def iife169_(_pat_let50_0):
                    def iife170_(d_1840_dt__update_hcf34_h0_):
                        def iife171_(_pat_let51_0):
                            def iife172_(d_1841_dt__update_hcf33_h0_):
                                return D6_DC17((d_1839_dt__update__tmp_h1_).cf30, (d_1839_dt__update__tmp_h1_).cf31, (d_1839_dt__update__tmp_h1_).cf32, d_1841_dt__update_hcf33_h0_, d_1840_dt__update_hcf34_h0_)
                            return iife172_(_pat_let51_0)
                        return iife171_((0) - (pat_let_tv52_.f40))
                    return iife170_(_pat_let50_0)
                return iife169_(pat_let_tv51_)
            return iife168_(_pat_let49_0)
        nw287_[int(16)] = iife167_(D6_DC17(d_1822_v0_, len(default__.fm51(d_1822_v0_, d_1822_v0_, d_1827_v5_, 597, globalState)), d_1824_v2_.f40, d_1822_v0_, d_1824_v2_.f40))
        nw287_[int(17)] = d_1831_v9_
        nw287_[int(18)] = D6_DC17(d_1824_v2_.f40, (d_1828_v6_).fm13(d_1834_v12_, d_1832_v10_, d_1824_v2_.f40, globalState), d_1822_v0_, ((self).f29).cardinality, d_1822_v0_)
        nw287_[int(19)] = d_1831_v9_
        def iife173_(_pat_let52_0):
            def iife174_(d_1842_dt__update__tmp_h2_):
                def iife175_(_pat_let53_0):
                    def iife176_(d_1843_dt__update_hcf33_h1_):
                        def iife177_(_pat_let54_0):
                            def iife178_(d_1844_dt__update_hcf30_h1_):
                                return D6_DC17(d_1844_dt__update_hcf30_h1_, (d_1842_dt__update__tmp_h2_).cf31, (d_1842_dt__update__tmp_h2_).cf32, d_1843_dt__update_hcf33_h1_, (d_1842_dt__update__tmp_h2_).cf34)
                            return iife178_(_pat_let54_0)
                        return iife177_(pat_let_tv54_)
                    return iife176_(_pat_let53_0)
                return iife175_(pat_let_tv53_.f40)
            return iife174_(_pat_let52_0)
        nw287_[int(20)] = iife173_(d_1831_v9_)
        nw287_[int(21)] = d_1831_v9_
        nw287_[int(22)] = d_1831_v9_
        d_1835_v13_ = nw287_
        index265_ = default__.safeIndex(913, (d_1835_v13_).length(0))
        (d_1835_v13_)[index265_] = D6_DC17(d_1822_v0_, d_1822_v0_, d_1822_v0_, 66, d_1822_v0_)
        index266_ = default__.safeIndex(913, (d_1835_v13_).length(0))
        (d_1835_v13_)[index266_] = D6_DC17((len(_dafny.SeqWithoutIsStrInference([d_1832_v10_, d_1832_v10_]))) + (d_1824_v2_.f40), d_1824_v2_.f40, d_1824_v2_.f40, d_1822_v0_, (0) - (d_1824_v2_.f40))
        (globalState).f12 = d_1822_v0_
        d_1845_v14_: _dafny.Map
        d_1845_v14_ = _dafny.Map({not(d_1827_v5_): d_1822_v0_})
        d_1846_v15_: _dafny.Map
        d_1846_v15_ = _dafny.Map({not(d_1827_v5_): d_1845_v14_})
        d_1846_v15_ = (d_1846_v15_).set(d_1827_v5_, (_dafny.Map({d_1827_v5_: d_1822_v0_})) | (_dafny.Map({d_1827_v5_: d_1822_v0_})))
        if not(not((d_1830_v8_) <= (_dafny.SeqWithoutIsStrInference([d_1827_v5_])))):
            d_1847_v16_: _dafny.Seq
            d_1847_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyvdj"))
            d_1848_v17_: _dafny.MultiSet
            d_1848_v17_ = _dafny.MultiSet([d_1827_v5_])
            d_1849_v18_: _dafny.Array
            nw288_ = _dafny.Array(None, 19)
            nw288_[int(0)] = d_1824_v2_.f40
            nw288_[int(1)] = d_1822_v0_
            nw288_[int(2)] = 158
            nw288_[int(3)] = d_1824_v2_.f40
            nw288_[int(4)] = (d_1824_v2_).fm1(globalState)
            nw288_[int(5)] = default__.fm39(globalState)
            nw288_[int(6)] = len(d_1847_v16_)
            nw288_[int(7)] = d_1824_v2_.f40
            nw288_[int(8)] = default__.safeModuloInt(d_1824_v2_.f40, (d_1848_v17_).cardinality)
            nw288_[int(9)] = (d_1824_v2_.f40) + (d_1824_v2_.f40)
            nw288_[int(10)] = ((d_1824_v2_).fm1(globalState)) * (d_1822_v0_)
            nw288_[int(11)] = d_1822_v0_
            nw288_[int(12)] = d_1824_v2_.f40
            nw288_[int(13)] = (d_1824_v2_.f40) + (len(d_1845_v14_))
            nw288_[int(14)] = default__.safeDivisionInt(d_1822_v0_, d_1822_v0_)
            nw288_[int(15)] = -336
            nw288_[int(16)] = ((d_1848_v17_) | (_dafny.MultiSet([d_1827_v5_]))).cardinality
            nw288_[int(17)] = 416
            nw288_[int(18)] = default__.safeDivisionInt((d_1823_v1_)[default__.safeIndex(((self).f29).cardinality, len(d_1823_v1_))], d_1822_v0_)
            d_1849_v18_ = nw288_
            d_1849_v18_ = d_1849_v18_
            d_1850_v19_: _dafny.Map
            d_1850_v19_ = _dafny.Map({(d_1828_v6_).fm13(d_1834_v12_, d_1832_v10_, d_1824_v2_.f40, globalState): d_1847_v16_})
            d_1851_v20_: _dafny.Set
            d_1851_v20_ = _dafny.Set({d_1822_v0_, d_1824_v2_.f40})
            d_1852_v21_: _dafny.Map
            d_1852_v21_ = _dafny.Map({((d_1850_v19_)[len(d_1851_v20_)] if (len(d_1851_v20_)) in (d_1850_v19_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xicujy"))): ((_dafny.MultiSet(d_1847_v16_)).cardinality if d_1827_v5_ else (0) - ((0) - (d_1824_v2_.f40)))})
            d_1852_v21_ = (d_1852_v21_).set(d_1847_v16_, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akwbgonr"))) + (d_1847_v16_)))
            d_1853_v22_: D3
            d_1853_v22_ = D3_DC10(d_1849_v18_)
            d_1854_v23_: _dafny.Map
            d_1854_v23_ = _dafny.Map({d_1827_v5_: (d_1853_v22_).cf21})
            d_1855_v24_: _dafny.Map
            d_1855_v24_ = _dafny.Map({((d_1854_v23_)[d_1827_v5_] if (d_1827_v5_) in (d_1854_v23_) else d_1849_v18_): (default__.fm40(373, d_1827_v5_, True, globalState)) not in ((d_1847_v16_).set(default__.safeIndex(d_1824_v2_.f40, len(d_1847_v16_)), (d_1847_v16_)[default__.safeIndex(d_1824_v2_.f40, len(d_1847_v16_))]))})
            d_1855_v24_ = (d_1855_v24_).set(d_1849_v18_, (d_1824_v2_).fm44(globalState))
            d_1856_v25_: D14
            d_1856_v25_ = D14_DC32(d_1848_v17_, (d_1830_v8_)[default__.safeIndex(d_1824_v2_.f40, len(d_1830_v8_))], d_1827_v5_, d_1827_v5_, d_1824_v2_.f40)
            d_1857_v26_: D3
            d_1857_v26_ = D3_DC11(d_1832_v10_)
            d_1858_v27_: _dafny.Array
            nw289_ = _dafny.Array(None, 27)
            nw289_[int(0)] = d_1832_v10_
            nw289_[int(1)] = d_1832_v10_
            nw289_[int(2)] = d_1832_v10_
            nw289_[int(3)] = d_1832_v10_
            nw289_[int(4)] = d_1832_v10_
            nw289_[int(5)] = d_1832_v10_
            nw289_[int(6)] = d_1832_v10_
            nw289_[int(7)] = d_1832_v10_
            nw289_[int(8)] = (d_1857_v26_).cf22
            nw289_[int(9)] = d_1832_v10_
            nw289_[int(10)] = d_1832_v10_
            nw289_[int(11)] = d_1832_v10_
            nw289_[int(12)] = d_1832_v10_
            nw289_[int(13)] = d_1832_v10_
            nw289_[int(14)] = d_1832_v10_
            nw289_[int(15)] = d_1832_v10_
            nw289_[int(16)] = _dafny.CodePoint('v')
            nw289_[int(17)] = d_1832_v10_
            nw289_[int(18)] = d_1832_v10_
            nw289_[int(19)] = d_1832_v10_
            nw289_[int(20)] = d_1832_v10_
            nw289_[int(21)] = d_1832_v10_
            nw289_[int(22)] = d_1832_v10_
            nw289_[int(23)] = d_1832_v10_
            nw289_[int(24)] = d_1832_v10_
            nw289_[int(25)] = d_1832_v10_
            nw289_[int(26)] = _dafny.CodePoint('c')
            d_1858_v27_ = nw289_
            d_1859_v28_: D1
            d_1859_v28_ = D1_DC4(d_1858_v27_, d_1824_v2_.f40, d_1824_v2_.f40, d_1847_v16_, d_1847_v16_)
            d_1860_v29_: _dafny.Map
            d_1860_v29_ = _dafny.Map({d_1822_v0_: len(d_1851_v20_)})
            d_1861_v30_: _dafny.Seq
            d_1861_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({((d_1860_v29_)[d_1824_v2_.f40] if (d_1824_v2_.f40) in (d_1860_v29_) else len(d_1860_v29_)), d_1824_v2_.f40})])
            d_1862_v31_: _dafny.Array
            nw290_ = _dafny.Array(None, 22)
            nw290_[int(0)] = d_1827_v5_
            nw290_[int(1)] = d_1827_v5_
            nw290_[int(2)] = d_1827_v5_
            nw290_[int(3)] = d_1827_v5_
            nw290_[int(4)] = False
            nw290_[int(5)] = not(d_1827_v5_)
            nw290_[int(6)] = (d_1827_v5_ if (d_1856_v25_).cf58 else False)
            nw290_[int(7)] = False
            nw290_[int(8)] = ((d_1859_v28_).cf10) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyjjbux")))
            nw290_[int(9)] = d_1827_v5_
            nw290_[int(10)] = d_1827_v5_
            nw290_[int(11)] = d_1827_v5_
            nw290_[int(12)] = d_1827_v5_
            nw290_[int(13)] = (d_1827_v5_) or (d_1827_v5_)
            nw290_[int(14)] = d_1827_v5_
            nw290_[int(15)] = (d_1851_v20_).issubset((d_1861_v30_)[default__.safeIndex(d_1824_v2_.f40, len(d_1861_v30_))])
            nw290_[int(16)] = (not(True)) and (d_1827_v5_)
            nw290_[int(17)] = d_1827_v5_
            nw290_[int(18)] = (d_1827_v5_) and (d_1827_v5_)
            nw290_[int(19)] = False
            nw290_[int(20)] = False
            nw290_[int(21)] = not(d_1827_v5_)
            d_1862_v31_ = nw290_
            index267_ = default__.safeIndex(536, (d_1862_v31_).length(0))
            (d_1862_v31_)[index267_] = d_1827_v5_
            index268_ = default__.safeIndex(536, (d_1862_v31_).length(0))
            (d_1862_v31_)[index268_] = not((d_1824_v2_.f40) != (((0) - (d_1822_v0_) if d_1827_v5_ else d_1822_v0_)))
            index269_ = default__.safeIndex(663, (d_1849_v18_).length(0))
            (d_1849_v18_)[index269_] = d_1824_v2_.f40
            index270_ = default__.safeIndex(663, (d_1849_v18_).length(0))
            (d_1849_v18_)[index270_] = d_1824_v2_.f40
        elif True:
            r1 = (((d_1829_v7_)[d_1828_v6_] if (d_1828_v6_) in (d_1829_v7_) else d_1822_v0_)) < (d_1822_v0_)
            d_1863_v32_: _dafny.Seq
            d_1863_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vaiuecm"))
            (globalState).f6 = d_1863_v32_
            d_1864_v33_: _dafny.Array
            nw291_ = _dafny.Array(int(0), 18)
            d_1864_v33_ = nw291_
            index271_ = default__.safeIndex(192, (d_1864_v33_).length(0))
            (d_1864_v33_)[index271_] = -728
            d_1865_v34_: _dafny.Map
            d_1865_v34_ = _dafny.Map({480: _dafny.Map({d_1827_v5_: d_1822_v0_})})
            d_1866_v35_: T0
            nw292_ = C7()
            nw292_.ctor__(d_1827_v5_, d_1865_v34_)
            d_1866_v35_ = nw292_
            d_1867_v36_: _dafny.Set
            d_1867_v36_ = _dafny.Set({d_1866_v35_})
            index272_ = default__.safeIndex(192, (d_1864_v33_).length(0))
            (d_1864_v33_)[index272_] = len(((d_1867_v36_) | (d_1867_v36_) if not((d_1827_v5_) and (d_1827_v5_)) else (d_1867_v36_) - (d_1867_v36_)))
            d_1868_v37_: D3
            d_1868_v37_ = D3_DC10(d_1864_v33_)
            d_1869_v38_: _dafny.Map
            d_1869_v38_ = _dafny.Map({d_1868_v37_: d_1824_v2_.f40})
            d_1870_v39_: _dafny.MultiSet
            d_1870_v39_ = _dafny.MultiSet([d_1827_v5_, d_1827_v5_])
            d_1871_v40_: D14
            d_1871_v40_ = D14_DC32(d_1870_v39_, d_1827_v5_, not(d_1827_v5_), d_1827_v5_, d_1824_v2_.f40)
            d_1872_v41_: C1
            nw293_ = C1()
            nw293_.ctor__(((self).f29).cardinality, len(default__.fm52(d_1827_v5_, (d_1864_v33_)[default__.safeIndex(192, (d_1864_v33_).length(0))], globalState)), (d_1871_v40_).cf57, d_1827_v5_)
            d_1872_v41_ = nw293_
            d_1873_v42_: _dafny.Seq
            d_1873_v42_ = _dafny.SeqWithoutIsStrInference([d_1872_v41_])
            d_1874_v43_: C4
            nw294_ = C4()
            nw294_.ctor__(d_1869_v38_, d_1827_v5_, ((d_1873_v42_).set(default__.safeIndex((d_1872_v41_).f36, len(d_1873_v42_)), d_1872_v41_)) != (_dafny.SeqWithoutIsStrInference([d_1872_v41_, d_1872_v41_])))
            d_1874_v43_ = nw294_
            d_1875_v44_: _dafny.Seq
            d_1875_v44_ = _dafny.SeqWithoutIsStrInference([d_1863_v32_])
            d_1876_v45_: _dafny.Array
            nw295_ = _dafny.Array(None, 25)
            nw295_[int(0)] = d_1863_v32_
            nw295_[int(1)] = d_1863_v32_
            nw295_[int(2)] = d_1863_v32_
            nw295_[int(3)] = (d_1875_v44_)[default__.safeIndex((d_1864_v33_)[default__.safeIndex(192, (d_1864_v33_).length(0))], len(d_1875_v44_))]
            nw295_[int(4)] = d_1863_v32_
            nw295_[int(5)] = d_1863_v32_
            nw295_[int(6)] = d_1863_v32_
            nw295_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hiqyfxr"))
            nw295_[int(8)] = d_1863_v32_
            nw295_[int(9)] = ((d_1863_v32_).set(default__.safeIndex(-400, len(d_1863_v32_)), d_1832_v10_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfc")))
            nw295_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsbqonwkb"))
            nw295_[int(11)] = d_1863_v32_
            nw295_[int(12)] = d_1863_v32_
            nw295_[int(13)] = d_1863_v32_
            nw295_[int(14)] = d_1863_v32_
            nw295_[int(15)] = (d_1863_v32_).set(default__.safeIndex(d_1824_v2_.f40, len(d_1863_v32_)), d_1832_v10_)
            nw295_[int(16)] = d_1863_v32_
            nw295_[int(17)] = d_1863_v32_
            nw295_[int(18)] = d_1863_v32_
            nw295_[int(19)] = d_1863_v32_
            nw295_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kasxjomfh"))
            nw295_[int(21)] = d_1863_v32_
            nw295_[int(22)] = (d_1875_v44_)[default__.safeIndex(d_1824_v2_.f40, len(d_1875_v44_))]
            nw295_[int(23)] = d_1863_v32_
            nw295_[int(24)] = d_1863_v32_
            d_1876_v45_ = nw295_
            index273_ = default__.safeIndex(935, (d_1876_v45_).length(0))
            (d_1876_v45_)[index273_] = (d_1863_v32_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1877_i0_ in range(default__.abs(283))]))
            index274_ = default__.safeIndex(935, (d_1876_v45_).length(0))
            (d_1876_v45_)[index274_] = d_1863_v32_
        d_1878_v46_: _dafny.Seq
        d_1878_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tgaaxgqsq"))
        d_1879_v47_: D6
        d_1879_v47_ = D6_DC16((d_1878_v46_) + (d_1878_v46_))
        pat_let_tv55_ = d_1878_v46_
        def iife179_(_pat_let55_0):
            def iife180_(d_1880_dt__update__tmp_h3_):
                def iife181_(_pat_let56_0):
                    def iife182_(d_1881_dt__update_hcf29_h0_):
                        return D6_DC16(d_1881_dt__update_hcf29_h0_)
                    return iife182_(_pat_let56_0)
                return iife181_(pat_let_tv55_)
            return iife180_(_pat_let55_0)
        d_1879_v47_ = (d_1879_v47_ if d_1827_v5_ else iife179_(d_1879_v47_))
        hi17_ = (d_1824_v2_.f40) - (d_1822_v0_)
        for d_1882_i1_ in range((-597) * (d_1822_v0_), hi17_):
            d_1832_v10_ = d_1832_v10_
            d_1883_v48_: C0
            nw296_ = C0()
            nw296_.ctor__(True)
            d_1883_v48_ = nw296_
            (globalState).f22 = not((d_1827_v5_ if d_1827_v5_ else (d_1883_v48_).f32))
            d_1884_v49_: D0
            d_1884_v49_ = D0_DC1(True, (d_1883_v48_).f32, d_1824_v2_.f40, (d_1883_v48_).f32)
            d_1885_v50_: _dafny.Set
            d_1885_v50_ = _dafny.Set({d_1884_v49_})
            d_1885_v50_ = (d_1885_v50_) - ((d_1885_v50_) | (_dafny.Set({d_1884_v49_})))
        r0 = d_1827_v5_
        r1 = (d_1830_v8_)[default__.safeIndex(d_1824_v2_.f40, len(d_1830_v8_))]
        return r0, r1


class C9(T2, T3, T1, T0):
    def  __init__(self):
        self._f26: bool = False
        self._f27: bool = False
        self._f29: _dafny.MultiSet = _dafny.MultiSet({})
        self.f30: bool = False
        self._f31: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f30, f31, f26, f27, f29):
        (self).f30 = f30
        (self)._f31 = f31
        (self)._f26 = f26
        (self)._f27 = f27
        (self)._f29 = f29

    def fm3(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])) + (_dafny.SeqWithoutIsStrInference([(self).f27]))) + ((_dafny.SeqWithoutIsStrInference([(self).f26])) + (_dafny.SeqWithoutIsStrInference([(self).f27])))

    def fm4(self, p0, p1, p2, globalState):
        return ((self).f31) - ((self).f31)

    def fm1(self, globalState):
        return (self).f31

    def fm2(self, p0, p1, globalState):
        return ((self).f27) == ((self).f27)

    def m2(self, p0, p1, p2, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        d_1886_v0_: _dafny.Array
        nw297_ = _dafny.Array(_dafny.Map({}), 3)
        d_1886_v0_ = nw297_
        d_1887_v1_: _dafny.Map
        d_1887_v1_ = _dafny.Map({(self).f27: self.f30})
        index275_ = default__.safeIndex(333, (d_1886_v0_).length(0))
        (d_1886_v0_)[index275_] = d_1887_v1_
        index276_ = default__.safeIndex(333, (d_1886_v0_).length(0))
        rhs285_ = ((0) - ((self).f31) if self.f30 else default__.safeDivisionInt((self).f31, (self).f31))
        rhs286_ = (d_1887_v1_) | ((d_1887_v1_) | (d_1887_v1_))
        lhs257_ = globalState
        lhs258_ = d_1886_v0_
        lhs259_ = default__.safeIndex(333, (d_1886_v0_).length(0))
        lhs257_.f12 = rhs285_
        lhs258_[lhs259_] = rhs286_
        d_1888_v2_: _dafny.Seq
        d_1888_v2_ = _dafny.SeqWithoutIsStrInference([(((d_1886_v0_)[default__.safeIndex(333, (d_1886_v0_).length(0))])[(self).f27] if ((self).f27) in ((d_1886_v0_)[default__.safeIndex(333, (d_1886_v0_).length(0))]) else p0), not(self.f30), (self).fm2(True, len(_dafny.Map({(self).f27: (self).f31})), globalState)])
        d_1889_v3_: str
        d_1889_v3_ = _dafny.CodePoint('v')
        d_1890_v4_: _dafny.Map
        d_1890_v4_ = _dafny.Map({len(d_1888_v2_): d_1889_v3_})
        d_1891_v5_: _dafny.Map
        d_1891_v5_ = _dafny.Map({(self).f31: self.f30})
        d_1892_v6_: _dafny.Set
        d_1892_v6_ = _dafny.Set({((d_1890_v4_)[len(d_1891_v5_)] if (len(d_1891_v5_)) in (d_1890_v4_) else d_1889_v3_), d_1889_v3_})
        d_1893_i0_: int
        d_1893_i0_ = 0
        with _dafny.label("10"):
            while (d_1892_v6_).isdisjoint(d_1892_v6_):
                with _dafny.c_label("10"):
                    if (d_1893_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1893_i0_ = (d_1893_i0_) + (1)
                    d_1894_v7_: D0
                    d_1894_v7_ = D0_DC0(p2)
                    d_1895_v8_: _dafny.Map
                    d_1895_v8_ = _dafny.Map({(self).f31: (self).f31})
                    pat_let_tv56_ = globalState
                    pat_let_tv57_ = d_1895_v8_
                    def iife183_(_pat_let57_0):
                        def iife184_(d_1896_dt__update__tmp_h0_):
                            def iife185_(_pat_let58_0):
                                def iife186_(d_1897_dt__update_hcf0_h0_):
                                    return D0_DC0(d_1897_dt__update_hcf0_h0_)
                                return iife186_(_pat_let58_0)
                            return iife185_((_dafny.Map({(self).fm1(pat_let_tv56_): (self).f31})) != (pat_let_tv57_))
                        return iife184_(_pat_let57_0)
                    source27_ = iife183_(d_1894_v7_)
                    if source27_.is_DC1:
                        d_1898___mcc_h0_ = source27_.cf1
                        d_1899___mcc_h1_ = source27_.cf2
                        d_1900___mcc_h2_ = source27_.cf3
                        d_1901___mcc_h3_ = source27_.cf4
                        d_1902_cf4_ = d_1901___mcc_h3_
                        d_1903_cf3_ = d_1900___mcc_h2_
                        d_1904_cf2_ = d_1899___mcc_h1_
                        d_1905_cf1_ = d_1898___mcc_h0_
                        (globalState).f22 = not((d_1894_v7_).cf0)
                        d_1890_v4_ = d_1890_v4_
                        d_1906_v9_: _dafny.Array
                        def lambda87_(d_1907_p2_):
                            def lambda88_(d_1908_i1_):
                                return d_1907_p2_

                            return lambda88_

                        init48_ = lambda87_(p2)
                        nw298_ = _dafny.Array(None, 14)
                        for i0_48_ in range(nw298_.length(0)):
                            nw298_[i0_48_] = init48_(i0_48_)
                        d_1906_v9_ = nw298_
                        index277_ = default__.safeIndex(924, (d_1906_v9_).length(0))
                        (d_1906_v9_)[index277_] = ((self).f26 if True else self.f30)
                        index278_ = default__.safeIndex(924, (d_1906_v9_).length(0))
                        (d_1906_v9_)[index278_] = d_1902_cf4_
                        (globalState).f12 = (((self).f29)[d_1903_cf3_] if (d_1903_cf3_) in ((self).f29) else (self).f31)
                    elif source27_.is_DC0:
                        d_1909___mcc_h4_ = source27_.cf0
                        d_1910_cf0_ = d_1909___mcc_h4_
                        d_1911_v10_: _dafny.Array
                        def lambda89_(d_1912_p0_):
                            def lambda90_(d_1913_i2_):
                                return d_1912_p0_

                            return lambda90_

                        init49_ = lambda89_(p0)
                        nw299_ = _dafny.Array(None, 25)
                        for i0_49_ in range(nw299_.length(0)):
                            nw299_[i0_49_] = init49_(i0_49_)
                        d_1911_v10_ = nw299_
                        d_1914_v11_: _dafny.Seq
                        d_1914_v11_ = _dafny.SeqWithoutIsStrInference([d_1911_v10_, d_1911_v10_, d_1911_v10_, d_1911_v10_, d_1911_v10_])
                        (globalState).f12 = ((692 if p0 else (self).f31)) - (len((d_1914_v11_) + (d_1914_v11_)))
                        d_1915_v12_: _dafny.Seq
                        d_1915_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bex"))
                        d_1916_v13_: _dafny.Array
                        nw300_ = _dafny.Array(None, 24)
                        nw300_[int(0)] = (self).f31
                        nw300_[int(1)] = -773
                        nw300_[int(2)] = len(d_1915_v12_)
                        nw300_[int(3)] = (self).f31
                        nw300_[int(4)] = (self).f31
                        nw300_[int(5)] = (self).f31
                        nw300_[int(6)] = (self).f31
                        nw300_[int(7)] = (self).f31
                        nw300_[int(8)] = (self).f31
                        nw300_[int(9)] = (self).f31
                        nw300_[int(10)] = (self).f31
                        nw300_[int(11)] = (self).f31
                        nw300_[int(12)] = (self).f31
                        nw300_[int(13)] = (self).f31
                        nw300_[int(14)] = (self).f31
                        nw300_[int(15)] = (self).f31
                        nw300_[int(16)] = (self).f31
                        nw300_[int(17)] = len(d_1888_v2_)
                        nw300_[int(18)] = (self).f31
                        nw300_[int(19)] = (self).f31
                        nw300_[int(20)] = len(d_1915_v12_)
                        nw300_[int(21)] = (self).f31
                        nw300_[int(22)] = (self).f31
                        nw300_[int(23)] = len(_dafny.SeqWithoutIsStrInference([d_1888_v2_, d_1888_v2_]))
                        d_1916_v13_ = nw300_
                        d_1917_v14_: D3
                        d_1917_v14_ = D3_DC10(d_1916_v13_)
                        d_1918_v15_: _dafny.Map
                        d_1918_v15_ = _dafny.Map({d_1917_v14_: (self).f31})
                        d_1919_v16_: C4
                        nw301_ = C4()
                        nw301_.ctor__(d_1918_v15_, d_1910_cf0_, not((self).f26))
                        d_1919_v16_ = nw301_
                        (globalState).f2 = True
                        index279_ = default__.safeIndex(971, (d_1916_v13_).length(0))
                        (d_1916_v13_)[index279_] = (self).f31
                        index280_ = default__.safeIndex(971, (d_1916_v13_).length(0))
                        (d_1916_v13_)[index280_] = (self).f31
                    elif True:
                        d_1920___mcc_h5_ = source27_.cf5
                        d_1921_cf5_ = d_1920___mcc_h5_
                        d_1922_v17_: _dafny.Array
                        nw302_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                        d_1922_v17_ = nw302_
                        d_1923_v18_: _dafny.Seq
                        d_1923_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbooou"))
                        index281_ = default__.safeIndex(893, (d_1922_v17_).length(0))
                        (d_1922_v17_)[index281_] = d_1923_v18_
                        d_1924_v19_: _dafny.Set
                        d_1924_v19_ = _dafny.Set({self.f30})
                        index282_ = default__.safeIndex(893, (d_1922_v17_).length(0))
                        (d_1922_v17_)[index282_] = (d_1923_v18_ if (d_1924_v19_).ispropersubset(d_1924_v19_) else (d_1923_v18_) + (d_1923_v18_))
                        d_1925_v20_: _dafny.Set
                        d_1925_v20_ = _dafny.Set({(self).f31, (self).f31})
                        d_1926_v21_: _dafny.Array
                        nw303_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                        d_1926_v21_ = nw303_
                        d_1927_v22_: D1
                        d_1927_v22_ = D1_DC4(d_1926_v21_, 731, 279, _dafny.SeqWithoutIsStrInference([d_1889_v3_ for d_1928_i3_ in range(default__.abs(60))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
                        (globalState).f12 = default__.safeModuloInt((0) - (default__.safeModuloInt(len(d_1923_v18_), len(d_1925_v20_))), (d_1927_v22_).cf8)
                        d_1929_v23_: _dafny.Map
                        d_1929_v23_ = _dafny.Map({p0: (self).f31})
                        (globalState).f2 = (False if (self).f27 else (p1).fm2((self).fm2(self.f30, len(d_1929_v23_), globalState), (self).f31, globalState))
                        (globalState).f2 = (self).fm2(((self).f31) >= ((self).f31), ((self).f31) + (111), globalState)
                    d_1930_v24_: _dafny.Seq
                    d_1930_v24_ = _dafny.SeqWithoutIsStrInference([d_1889_v3_, d_1889_v3_])
                    d_1931_v25_: _dafny.Seq
                    d_1931_v25_ = _dafny.SeqWithoutIsStrInference([d_1930_v24_, _dafny.SeqWithoutIsStrInference([d_1889_v3_ for d_1932_i4_ in range(default__.abs(165))])])
                    d_1931_v25_ = d_1931_v25_
                    if self.f30:
                        d_1933_v26_: _dafny.Array
                        nw304_ = _dafny.Array(False, 4)
                        d_1933_v26_ = nw304_
                        index283_ = default__.safeIndex(357, (d_1933_v26_).length(0))
                        (d_1933_v26_)[index283_] = not((p1).fm2((self).f26, (self).f31, globalState))
                        index284_ = default__.safeIndex(357, (d_1933_v26_).length(0))
                        (d_1933_v26_)[index284_] = (self).fm2((self).f26, (self).f31, globalState)
                        d_1934_v27_: _dafny.Seq
                        d_1934_v27_ = _dafny.SeqWithoutIsStrInference([d_1933_v26_])
                        d_1935_v28_: _dafny.Array
                        def lambda91_(d_1936_i5_):
                            return (d_1936_i5_) + ((self).f31)

                        init50_ = lambda91_
                        nw305_ = _dafny.Array(None, 17)
                        for i0_50_ in range(nw305_.length(0)):
                            nw305_[i0_50_] = init50_(i0_50_)
                        d_1935_v28_ = nw305_
                        rhs287_ = d_1934_v27_
                        rhs288_ = (p1).fm2((d_1933_v26_)[default__.safeIndex(357, (d_1933_v26_).length(0))], (self).f31, globalState)
                        rhs289_ = d_1935_v28_
                        rhs290_ = default__.safeModuloInt((self).f31, (self).f31)
                        lhs260_ = globalState
                        lhs261_ = globalState
                        d_1934_v27_ = rhs287_
                        lhs260_.f24 = rhs288_
                        d_1935_v28_ = rhs289_
                        lhs261_.f16 = rhs290_
                        d_1937_v29_: _dafny.Seq
                        d_1937_v29_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f31]))])
                        index285_ = default__.safeIndex(292, (d_1935_v28_).length(0))
                        (d_1935_v28_)[index285_] = ((self).f31 if p0 else len(d_1937_v29_))
                        d_1938_v30_: _dafny.Seq
                        d_1938_v30_ = _dafny.SeqWithoutIsStrInference([d_1934_v27_, d_1934_v27_])
                        index286_ = default__.safeIndex(292, (d_1935_v28_).length(0))
                        rhs291_ = (0) - ((self).f31)
                        rhs292_ = (p1).fm2((d_1933_v26_) in ((d_1938_v30_)[default__.safeIndex(174, len(d_1938_v30_))]), 22, globalState)
                        rhs293_ = (self).f31
                        rhs294_ = self.f30
                        lhs262_ = globalState
                        lhs263_ = self
                        lhs264_ = d_1935_v28_
                        lhs265_ = default__.safeIndex(292, (d_1935_v28_).length(0))
                        lhs266_ = globalState
                        lhs262_.f12 = rhs291_
                        lhs263_.f30 = rhs292_
                        lhs264_[lhs265_] = rhs293_
                        lhs266_.f15 = rhs294_
                        d_1933_v26_ = d_1933_v26_
                        d_1939_v31_: _dafny.Map
                        d_1939_v31_ = _dafny.Map({D3_DC10(d_1935_v28_): (d_1935_v28_)[default__.safeIndex(292, (d_1935_v28_).length(0))]})
                        d_1940_v32_: C4
                        nw306_ = C4()
                        nw306_.ctor__(d_1939_v31_, p2, not(p0))
                        d_1940_v32_ = nw306_
                    elif True:
                        d_1941_v33_: _dafny.Map
                        d_1941_v33_ = _dafny.Map({p2: (self).f31})
                        d_1942_v34_: int
                        d_1943_v35_: int
                        d_1944_v36_: bool
                        out61_: int
                        out62_: int
                        out63_: bool
                        out61_, out62_, out63_ = (self).m6(d_1941_v33_, (0) - ((self).f31), d_1889_v3_, (self).f31, globalState)
                        d_1942_v34_ = out61_
                        d_1943_v35_ = out62_
                        d_1944_v36_ = out63_
                        d_1889_v3_ = _dafny.CodePoint('o')
                        d_1945_v37_: _dafny.Array
                        def lambda92_(d_1946_i6_):
                            return (self).f26

                        init51_ = lambda92_
                        nw307_ = _dafny.Array(None, 3)
                        for i0_51_ in range(nw307_.length(0)):
                            nw307_[i0_51_] = init51_(i0_51_)
                        d_1945_v37_ = nw307_
                        d_1945_v37_ = d_1945_v37_
                        d_1947_v38_: _dafny.Map
                        d_1947_v38_ = _dafny.Map({p1: p0})
                        d_1948_v39_: D12
                        d_1948_v39_ = D12_DC27(p1)
                        (globalState).f24 = ((d_1947_v38_)[(d_1948_v39_).cf48] if ((d_1948_v39_).cf48) in (d_1947_v38_) else p0)
                        (globalState).f24 = (self).f26
                    d_1949_v40_: _dafny.Set
                    d_1949_v40_ = _dafny.Set({d_1887_v1_, ((d_1886_v0_)[default__.safeIndex(333, (d_1886_v0_).length(0))]).set(p0, (self).f26), d_1887_v1_})
                    d_1949_v40_ = (d_1949_v40_) | (d_1949_v40_)
                    pass
            pass
        d_1950_v41_: D12
        d_1950_v41_ = D12_DC28(229)
        source28_ = d_1950_v41_
        if source28_.is_DC28:
            d_1951___mcc_h6_ = source28_.cf49
            d_1952_cf49_ = d_1951___mcc_h6_
            (globalState).f22 = ((self.f30 if self.f30 else p2) if self.f30 else self.f30)
            (globalState).f2 = (d_1888_v2_)[default__.safeIndex(237, len(d_1888_v2_))]
            d_1953_v42_: C5
            nw308_ = C5()
            nw308_.ctor__(False, False)
            d_1953_v42_ = nw308_
            (globalState).f12 = (self).f31
        elif True:
            d_1954___mcc_h7_ = source28_.cf48
            d_1955_cf48_ = d_1954___mcc_h7_
            d_1956_v43_: _dafny.Array
            def lambda93_(d_1957_i7_):
                return (d_1957_i7_) + ((self).f31)

            init52_ = lambda93_
            nw309_ = _dafny.Array(None, 26)
            for i0_52_ in range(nw309_.length(0)):
                nw309_[i0_52_] = init52_(i0_52_)
            d_1956_v43_ = nw309_
            d_1958_v44_: D3
            d_1958_v44_ = D3_DC10(d_1956_v43_)
            d_1959_v45_: _dafny.Map
            d_1959_v45_ = _dafny.Map({d_1958_v44_: len(_dafny.SeqWithoutIsStrInference([self.f30]))})
            d_1960_v46_: C4
            nw310_ = C4()
            nw310_.ctor__((_dafny.Map({d_1958_v44_: (self).f31})) | (d_1959_v45_), p2, (p0 if p2 else (self).f27))
            d_1960_v46_ = nw310_
            if (self).f26:
                d_1961_v47_: _dafny.Set
                d_1961_v47_ = _dafny.Set({(d_1960_v46_).fm1(globalState), (self).f31})
                d_1961_v47_ = d_1961_v47_
                d_1962_v48_: _dafny.Seq
                d_1962_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsjg"))
                (globalState).f6 = d_1962_v48_
                d_1963_v49_: _dafny.Set
                d_1963_v49_ = _dafny.Set({d_1962_v48_, d_1962_v48_, d_1962_v48_})
                d_1964_v50_: _dafny.MultiSet
                d_1964_v50_ = _dafny.MultiSet([d_1889_v3_])
                d_1965_v51_: D0
                d_1965_v51_ = D0_DC0(p2)
                d_1966_v52_: _dafny.Set
                d_1966_v52_ = _dafny.Set({(d_1888_v2_)[default__.safeIndex((self).f31, len(d_1888_v2_))]})
                d_1967_v53_: T2
                nw311_ = C1()
                nw311_.ctor__(len(d_1962_v48_), 800, (self).f27, p0)
                d_1967_v53_ = nw311_
                d_1968_v54_: _dafny.Seq
                d_1968_v54_ = _dafny.SeqWithoutIsStrInference([d_1967_v53_, d_1967_v53_])
                d_1969_v55_: _dafny.Map
                d_1969_v55_ = _dafny.Map({d_1968_v54_: (self).f31})
                d_1970_v56_: _dafny.Array
                nw312_ = _dafny.Array(None, 25)
                nw312_[int(0)] = (d_1963_v49_).ispropersubset(d_1963_v49_)
                nw312_[int(1)] = (self).f27
                nw312_[int(2)] = (self).f26
                nw312_[int(3)] = (d_1964_v50_).issubset(d_1964_v50_)
                nw312_[int(4)] = (d_1965_v51_).cf0
                nw312_[int(5)] = ((self).f27) or (self.f30)
                nw312_[int(6)] = ((self).f31) <= ((self).f31)
                nw312_[int(7)] = p0
                nw312_[int(8)] = not((self).f26)
                nw312_[int(9)] = not((_dafny.Set({self.f30, p2})).ispropersubset(d_1966_v52_))
                nw312_[int(10)] = (d_1955_cf48_).fm2((self).f27, (self).f31, globalState)
                nw312_[int(11)] = not(not((self.f30) or ((self).f27)))
                nw312_[int(12)] = p0
                nw312_[int(13)] = p0
                nw312_[int(14)] = self.f30
                nw312_[int(15)] = (self).fm2((self).f27, len(d_1969_v55_), globalState)
                nw312_[int(16)] = (d_1967_v53_).f26
                nw312_[int(17)] = not(((self).f27) == ((self).f27))
                nw312_[int(18)] = not((_dafny.Set({379})).isdisjoint(d_1961_v47_))
                nw312_[int(19)] = p2
                nw312_[int(20)] = not ((self).fm2((d_1967_v53_).f26, (self).f31, globalState)) or ((self).f27)
                nw312_[int(21)] = (self).f27
                nw312_[int(22)] = self.f30
                nw312_[int(23)] = (d_1967_v53_).f26
                nw312_[int(24)] = (d_1888_v2_)[default__.safeIndex((self).f31, len(d_1888_v2_))]
                d_1970_v56_ = nw312_
                index287_ = default__.safeIndex(56, (d_1970_v56_).length(0))
                (d_1970_v56_)[index287_] = not((d_1960_v46_).fm2((d_1967_v53_).f27, (self).f31, globalState))
                index288_ = default__.safeIndex(56, (d_1970_v56_).length(0))
                (d_1970_v56_)[index288_] = (d_1967_v53_).f26
                (globalState).f16 = len((d_1962_v48_) + ((d_1962_v48_) + (_dafny.SeqWithoutIsStrInference([d_1889_v3_ for d_1971_i8_ in range(default__.abs(-868))]))))
                d_1972_v57_: _dafny.Seq
                d_1972_v57_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                (globalState).f2 = (d_1888_v2_)[default__.safeIndex((len(d_1972_v57_)) + ((self).f31), len(d_1888_v2_))]
            elif True:
                d_1973_v58_: _dafny.Map
                d_1973_v58_ = _dafny.Map({(self).f31: (self).f31})
                d_1974_v59_: _dafny.Seq
                d_1974_v59_ = _dafny.SeqWithoutIsStrInference([len(d_1973_v58_), (self).f31])
                d_1975_v60_: T3
                nw313_ = C8()
                nw313_.ctor__(_dafny.MultiSet(d_1974_v59_))
                d_1975_v60_ = nw313_
                d_1976_v61_: _dafny.Seq
                d_1976_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtfvyyvhs"))
                d_1977_v62_: _dafny.Map
                d_1977_v62_ = _dafny.Map({d_1976_v61_: d_1888_v2_})
                d_1978_v63_: _dafny.Map
                d_1978_v63_ = _dafny.Map({(d_1975_v60_ if (self).f27 else d_1975_v60_): d_1977_v62_})
                d_1978_v63_ = (d_1978_v63_).set(d_1975_v60_, d_1977_v62_)
                index289_ = default__.safeIndex(435, (d_1956_v43_).length(0))
                (d_1956_v43_)[index289_] = (self).f31
                index290_ = default__.safeIndex(435, (d_1956_v43_).length(0))
                rhs295_ = (self).f31
                rhs296_ = True
                rhs297_ = (771) - ((self).f31)
                rhs298_ = (self).f31
                rhs299_ = (self).f27
                lhs267_ = d_1956_v43_
                lhs268_ = default__.safeIndex(435, (d_1956_v43_).length(0))
                lhs269_ = globalState
                lhs270_ = globalState
                lhs271_ = globalState
                lhs272_ = globalState
                lhs267_[lhs268_] = rhs295_
                lhs269_.f22 = rhs296_
                lhs270_.f12 = rhs297_
                lhs271_.f12 = rhs298_
                lhs272_.f22 = rhs299_
                d_1979_v64_: _dafny.MultiSet
                d_1979_v64_ = _dafny.MultiSet([(d_1956_v43_)[default__.safeIndex(435, (d_1956_v43_).length(0))]])
                d_1979_v64_ = d_1979_v64_
                d_1980_v65_: _dafny.MultiSet
                d_1980_v65_ = _dafny.MultiSet([d_1891_v5_])
                d_1981_v66_: _dafny.MultiSet
                d_1981_v66_ = d_1980_v65_
                index291_ = default__.safeIndex(435, (d_1956_v43_).length(0))
                rhs300_ = (d_1960_v46_).fm1(globalState)
                rhs301_ = (((d_1981_v66_)).cardinality) * (((self).f31 if self.f30 else len(d_1974_v59_)))
                rhs302_ = p0
                rhs303_ = ((d_1975_v60_).f29).intersection(((d_1975_v60_).f29) - (d_1979_v64_))
                lhs273_ = globalState
                lhs274_ = d_1956_v43_
                lhs275_ = default__.safeIndex(435, (d_1956_v43_).length(0))
                lhs276_ = globalState
                lhs273_.f12 = rhs300_
                lhs274_[lhs275_] = rhs301_
                lhs276_.f15 = rhs302_
                d_1979_v64_ = rhs303_
                d_1982_v67_: _dafny.MultiSet
                d_1982_v67_ = _dafny.MultiSet([d_1889_v3_, d_1889_v3_, _dafny.CodePoint('u'), d_1889_v3_])
                d_1983_v68_: _dafny.Array
                nw314_ = _dafny.Array(None, 7)
                nw314_[int(0)] = (_dafny.MultiSet([d_1889_v3_])).issubset(d_1982_v67_)
                nw314_[int(1)] = p0
                nw314_[int(2)] = not (True) or (p2)
                nw314_[int(3)] = not(not((self).f26))
                nw314_[int(4)] = p2
                nw314_[int(5)] = self.f30
                nw314_[int(6)] = (p0) or ((self).f27)
                d_1983_v68_ = nw314_
                index292_ = default__.safeIndex(355, (d_1983_v68_).length(0))
                (d_1983_v68_)[index292_] = (self).f26
                index293_ = default__.safeIndex(355, (d_1983_v68_).length(0))
                (d_1983_v68_)[index293_] = (self).f27
            d_1984_v69_: D2
            d_1984_v69_ = D2_DC7(p2, False, (self).f31, (((self).f29).cardinality) - ((self).f31))
            rhs304_ = (p2 if not(p2) else p0)
            rhs305_ = d_1984_v69_
            lhs277_ = globalState
            lhs277_.f15 = rhs304_
            d_1984_v69_ = rhs305_
            d_1985_v70_: _dafny.Array
            def lambda94_(d_1986_i9_):
                return (self).f26

            init53_ = lambda94_
            nw315_ = _dafny.Array(None, 23)
            for i0_53_ in range(nw315_.length(0)):
                nw315_[i0_53_] = init53_(i0_53_)
            d_1985_v70_ = nw315_
            index294_ = default__.safeIndex(795, (d_1985_v70_).length(0))
            (d_1985_v70_)[index294_] = (d_1955_cf48_).fm2((self).f26, (self).f31, globalState)
            index295_ = default__.safeIndex(795, (d_1985_v70_).length(0))
            (d_1985_v70_)[index295_] = p0
        d_1987_v71_: _dafny.Array
        nw316_ = _dafny.Array(False, 17)
        d_1987_v71_ = nw316_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_1987_v71_).length(0)):
            d_1988_i10_: int = guard_loop_0_
            if (True) and (((0) <= (d_1988_i10_)) and ((d_1988_i10_) < ((d_1987_v71_).length(0)))):
                (d_1987_v71_)[(d_1988_i10_)] = (D16_DC38((self).f31, 928, self.f30)).cf71
        hi18_ = (self).f31
        for d_1989_i11_ in range((self).f31, hi18_):
            d_1990_v72_: _dafny.Seq
            d_1990_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eobm"))
            (globalState).f6 = ((((d_1990_v72_).set(default__.safeIndex((self).f31, len(d_1990_v72_)), _dafny.CodePoint('e'))) + (d_1990_v72_)).set(default__.safeIndex(d_1989_i11_, len(((d_1990_v72_).set(default__.safeIndex((self).f31, len(d_1990_v72_)), _dafny.CodePoint('e'))) + (d_1990_v72_))), d_1889_v3_)) + (d_1990_v72_)
            (globalState).f6 = d_1990_v72_
            d_1991_v73_: _dafny.Map
            d_1991_v73_ = _dafny.Map({(self).f31: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxokab"))})
            d_1992_v74_: _dafny.Seq
            d_1992_v74_ = _dafny.SeqWithoutIsStrInference([d_1989_i11_, (self).f31])
            d_1993_v75_: _dafny.Seq
            d_1993_v75_ = _dafny.SeqWithoutIsStrInference([d_1992_v74_])
            d_1994_v76_: T2
            nw317_ = C2()
            nw317_.ctor__(d_1993_v75_, self.f30, p2)
            d_1994_v76_ = nw317_
            d_1995_v77_: _dafny.Array
            nw318_ = _dafny.Array(None, 3)
            nw318_[int(0)] = d_1994_v76_
            nw318_[int(1)] = d_1994_v76_
            nw318_[int(2)] = d_1994_v76_
            d_1995_v77_ = nw318_
            d_1996_v78_: _dafny.Seq
            d_1996_v78_ = _dafny.SeqWithoutIsStrInference([d_1995_v77_, d_1995_v77_])
            d_1997_v79_: _dafny.Set
            d_1997_v79_ = _dafny.Set({(self).f31, d_1989_i11_, len(d_1990_v72_), len(d_1991_v73_), len((d_1996_v78_) + (_dafny.SeqWithoutIsStrInference([d_1995_v77_, d_1995_v77_])))})
            def iife187_():
                coll69_ = _dafny.Set()
                compr_69_: int
                for compr_69_ in _dafny.IntegerRange(809, 689):
                    d_1998_v80_: int = compr_69_
                    if ((809) <= (d_1998_v80_)) and ((d_1998_v80_) < (689)):
                        coll69_ = coll69_.union(_dafny.Set([(d_1998_v80_) + ((self).f31)]))
                return _dafny.Set(coll69_)
            d_1997_v79_ = (iife187_()
            ).intersection(d_1997_v79_)
            (globalState).f12 = d_1989_i11_
        (globalState).f12 = (self).f31
        r0 = d_1889_v3_
        r1 = (self).f31
        return r0, r1

    def m1(self, p0, p1, p2, globalState):
        d_1999_v0_: D0
        d_1999_v0_ = D0_DC0((self).f27)
        d_2000_v1_: _dafny.Seq
        d_2000_v1_ = _dafny.SeqWithoutIsStrInference([(self).fm4(d_1999_v0_, (self).f31, (self).f31, globalState)])
        d_2001_v2_: _dafny.MultiSet
        d_2001_v2_ = _dafny.MultiSet([d_2000_v1_, d_2000_v1_])
        d_2002_v3_: C3
        nw319_ = C3()
        nw319_.ctor__(d_2001_v2_, self.f30, self.f30)
        d_2002_v3_ = nw319_
        d_2002_v3_ = d_2002_v3_
        d_2003_v4_: str
        d_2003_v4_ = _dafny.CodePoint('s')
        d_2004_v5_: _dafny.Map
        d_2004_v5_ = _dafny.Map({p0: d_2000_v1_})
        d_2005_v6_: _dafny.Map
        d_2005_v6_ = _dafny.Map({d_2003_v4_: len(((d_2004_v5_)[(self).f27] if ((self).f27) in (d_2004_v5_) else d_2000_v1_))})
        d_2006_v7_: _dafny.Array
        def lambda95_(d_2007_v4_):
            def lambda96_(d_2008_i1_):
                return d_2007_v4_

            return lambda96_

        init54_ = lambda95_(d_2003_v4_)
        nw320_ = _dafny.Array(None, 15)
        for i0_54_ in range(nw320_.length(0)):
            nw320_[i0_54_] = init54_(i0_54_)
        d_2006_v7_ = nw320_
        d_2009_v8_: _dafny.Seq
        d_2009_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbutbaoar"))
        d_2010_v9_: D1
        d_2010_v9_ = D1_DC4(d_2006_v7_, (self).f31, (self).f31, d_2009_v8_, d_2009_v8_)
        hi19_ = (d_2010_v9_).cf8
        for d_2011_i0_ in range(((d_2002_v3_).fm12(len(d_2005_v6_), globalState)) + ((self).f31), hi19_):
            d_2009_v8_ = (d_2009_v8_) + (d_2009_v8_)
            d_2012_v10_: _dafny.Array
            def lambda97_(d_2013_p0_):
                def lambda98_(d_2014_i2_):
                    return (d_2014_i2_) * (len(_dafny.Set({(self).f27, (self).f27, d_2013_p0_})))

                return lambda98_

            init55_ = lambda97_(p0)
            nw321_ = _dafny.Array(None, 3)
            for i0_55_ in range(nw321_.length(0)):
                nw321_[i0_55_] = init55_(i0_55_)
            d_2012_v10_ = nw321_
            index296_ = default__.safeIndex(745, (d_2012_v10_).length(0))
            (d_2012_v10_)[index296_] = ((self).f31) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sw"))))
            index297_ = default__.safeIndex(745, (d_2012_v10_).length(0))
            (d_2012_v10_)[index297_] = (self).f31
            d_2015_v11_: _dafny.Map
            d_2015_v11_ = _dafny.Map({self.f30: (self).fm3(d_2009_v8_, p2, (self).f31, globalState)})
            d_2016_v12_: D5
            d_2016_v12_ = D5_DC13((d_2015_v11_) | (d_2015_v11_))
            source29_ = d_2016_v12_
            if source29_.is_DC14:
                d_2017___mcc_h0_ = source29_.cf25
                d_2018___mcc_h1_ = source29_.cf26
                d_2019___mcc_h2_ = source29_.cf27
                d_2020_cf27_ = d_2019___mcc_h2_
                d_2021_cf26_ = d_2018___mcc_h1_
                d_2022_cf25_ = d_2017___mcc_h0_
                d_2023_v13_: _dafny.Map
                d_2023_v13_ = _dafny.Map({len(p2): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhvqkm"))})
                d_2024_v14_: _dafny.Map
                d_2024_v14_ = _dafny.Map({(self).f27: d_2022_cf25_})
                d_2025_v15_: _dafny.Seq
                d_2025_v15_ = _dafny.SeqWithoutIsStrInference([d_2024_v14_, d_2024_v14_])
                d_2026_v16_: D10
                d_2026_v16_ = D10_DC24(((d_2023_v13_)[d_2011_i0_] if (d_2011_i0_) in (d_2023_v13_) else d_2009_v8_), d_2025_v15_, (self).f26)
                d_2026_v16_ = d_2026_v16_
                (globalState).f12 = default__.safeModuloInt((self).f31, d_2021_cf26_)
                d_2027_v17_: _dafny.Seq
                d_2027_v17_ = _dafny.SeqWithoutIsStrInference([d_2000_v1_, d_2000_v1_])
                d_2028_v18_: C2
                nw322_ = C2()
                nw322_.ctor__(d_2027_v17_, True, self.f30)
                d_2028_v18_ = nw322_
                d_2029_v19_: _dafny.Map
                d_2029_v19_ = _dafny.Map({(d_2002_v3_).fm1(globalState): p0})
                d_2030_v20_: _dafny.MultiSet
                d_2030_v20_ = _dafny.MultiSet([d_2029_v19_])
                d_2031_v21_: _dafny.MultiSet
                d_2031_v21_ = d_2030_v20_
                d_2031_v21_ = d_2030_v20_
            elif source29_.is_DC13:
                d_2032___mcc_h3_ = source29_.cf24
                d_2033_cf24_ = d_2032___mcc_h3_
                (globalState).f15 = (self).f27
                d_2034_v22_: _dafny.Map
                d_2034_v22_ = _dafny.Map({d_2000_v1_: (0) - ((d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))])})
                d_2034_v22_ = (d_2034_v22_).set(d_2000_v1_, (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))])
                d_2003_v4_ = d_2003_v4_
                d_2035_v23_: D6
                d_2035_v23_ = D6_DC16(d_2009_v8_)
                d_2036_v24_: _dafny.Map
                d_2036_v24_ = _dafny.Map({(self).f26: (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))]})
                d_2037_v25_: _dafny.Map
                d_2037_v25_ = _dafny.Map({((_dafny.SeqWithoutIsStrInference([d_2011_i0_])) + (d_2000_v1_)).set(default__.safeIndex((d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))], len((_dafny.SeqWithoutIsStrInference([d_2011_i0_])) + (d_2000_v1_))), len((d_2035_v23_).cf29)): (d_2036_v24_) != (d_2036_v24_)})
                index298_ = default__.safeIndex(745, (d_2012_v10_).length(0))
                (d_2012_v10_)[index298_] = len(d_2037_v25_)
            elif True:
                d_2038___mcc_h4_ = source29_.cf28
                d_2039_cf28_ = d_2038___mcc_h4_
                d_2040_v26_: _dafny.Map
                d_2040_v26_ = _dafny.Map({d_2011_i0_: (self).fm2((p2)[default__.safeIndex((self).f31, len(p2))], d_2011_i0_, globalState)})
                d_2041_v27_: D16
                d_2041_v27_ = D16_DC37((d_2040_v26_).set((self).f31, p0))
                d_2042_v28_: _dafny.Map
                d_2042_v28_ = _dafny.Map({(0) - ((self).f31): (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))]})
                d_2043_v29_: _dafny.Map
                d_2043_v29_ = _dafny.Map({(d_2041_v27_).cf68: d_2042_v28_})
                d_2044_v31_: _dafny.Map
                d_2044_v31_ = _dafny.Map({_dafny.Map({103: (self).f26}): (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))]})
                d_2045_v32_: _dafny.Map
                d_2045_v32_ = d_2044_v31_
                def iife188_():
                    coll70_ = _dafny.Map()
                    compr_70_: _dafny.Map
                    for compr_70_ in ((d_2045_v32_)).keys.Elements:
                        d_2046_v30_: _dafny.Map = compr_70_
                        if (d_2046_v30_) in ((d_2045_v32_)):
                            coll70_[d_2046_v30_] = _dafny.Map({len(_dafny.Map({False: (self).f26})): (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))]})
                    return _dafny.Map(coll70_)
                d_2043_v29_ = (d_2043_v29_) | (iife188_()
                )
                d_2047_v33_: _dafny.Map
                d_2047_v33_ = _dafny.Map({len(d_2009_v8_): p2})
                d_2048_v34_: _dafny.Map
                d_2048_v34_ = _dafny.Map({p0: len(d_2009_v8_)})
                d_2049_v35_: D15
                d_2049_v35_ = D15_DC34((d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))], (self).f26, (self).f31)
                d_2050_v36_: _dafny.MultiSet
                d_2050_v36_ = _dafny.MultiSet([d_2049_v35_])
                d_2051_v37_: _dafny.Map
                d_2051_v37_ = _dafny.Map({((d_2047_v33_)[len(d_2048_v34_)] if (len(d_2048_v34_)) in (d_2047_v33_) else p2): (d_2050_v36_) | (d_2050_v36_)})
                d_2051_v37_ = _dafny.Map({(p2).set(default__.safeIndex((d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))], len(p2)), (self).f27): d_2050_v36_})
                d_2052_v38_: C0
                nw323_ = C0()
                nw323_.ctor__(True)
                d_2052_v38_ = nw323_
                d_2053_v39_: _dafny.Map
                d_2053_v39_ = _dafny.Map({(d_2052_v38_).f32: (self).f29})
                d_2054_v40_: _dafny.MultiSet
                d_2054_v40_ = _dafny.MultiSet([p0])
                rhs306_ = (d_2003_v4_ if self.f30 else d_2003_v4_)
                rhs307_ = d_2052_v38_
                rhs308_ = (len((d_2053_v39_) | (_dafny.Map({(self).fm2(p0, (d_2012_v10_)[default__.safeIndex(745, (d_2012_v10_).length(0))], globalState): (self).f29})))) * (len(_dafny.SeqWithoutIsStrInference([(self).f26])))
                rhs309_ = d_2009_v8_
                rhs310_ = not((d_2054_v40_).issubset(d_2054_v40_))
                lhs278_ = globalState
                lhs279_ = globalState
                lhs280_ = globalState
                lhs281_ = globalState
                lhs278_.f19 = rhs306_
                d_2052_v38_ = rhs307_
                lhs279_.f16 = rhs308_
                lhs280_.f6 = rhs309_
                lhs281_.f15 = rhs310_
                d_2055_v41_: _dafny.Array
                nw324_ = _dafny.Array(_dafny.Map({}), 2)
                d_2055_v41_ = nw324_
                d_2056_v42_: _dafny.Map
                d_2056_v42_ = _dafny.Map({(self).f27: (self).f26})
                d_2057_v43_: D13
                d_2057_v43_ = D13_DC29(d_2056_v42_)
                index299_ = default__.safeIndex(893, (d_2055_v41_).length(0))
                (d_2055_v41_)[index299_] = ((d_2057_v43_).cf50) | (d_2056_v42_)
                index300_ = default__.safeIndex(893, (d_2055_v41_).length(0))
                (d_2055_v41_)[index300_] = d_2056_v42_
            d_2058_v44_: _dafny.Map
            d_2058_v44_ = _dafny.Map({d_2012_v10_: (p2)[default__.safeIndex(len(p2), len(p2))]})
            d_2058_v44_ = (d_2058_v44_ if self.f30 else d_2058_v44_)
        d_2059_v45_: _dafny.MultiSet
        d_2059_v45_ = _dafny.MultiSet([(self).f26])
        d_2060_v46_: _dafny.Set
        d_2060_v46_ = _dafny.Set({False, (self).f27})
        d_2061_v47_: _dafny.Map
        d_2061_v47_ = _dafny.Map({d_2059_v45_: len(d_2060_v46_)})
        d_2062_v48_: D0
        d_2062_v48_ = D0_DC1(p0, False, len(d_2009_v8_), (p2)[default__.safeIndex(len(d_2061_v47_), len(p2))])
        d_2063_v49_: D0
        d_2063_v49_ = D0_DC2(d_2062_v48_)
        def lambda99_(source30_):
            if source30_.is_DC1:
                d_2064___mcc_h5_ = source30_.cf1
                d_2065___mcc_h6_ = source30_.cf2
                d_2066___mcc_h7_ = source30_.cf3
                d_2067___mcc_h8_ = source30_.cf4
                d_2068_cf4_ = d_2067___mcc_h8_
                d_2069_cf3_ = d_2066___mcc_h7_
                d_2070_cf2_ = d_2065___mcc_h6_
                d_2071_cf1_ = d_2064___mcc_h5_
                return d_2069_cf3_
            elif source30_.is_DC0:
                d_2072___mcc_h9_ = source30_.cf0
                d_2073_cf0_ = d_2072___mcc_h9_
                return 691
            elif True:
                d_2074___mcc_h10_ = source30_.cf5
                d_2075_cf5_ = d_2074___mcc_h10_
                return (self).f31

        (globalState).f12 = (0) - (lambda99_(d_2063_v49_))
        pat_let_tv58_ = d_2059_v45_
        d_2076_v50_: _dafny.Map
        def iife189_(_pat_let59_0):
            def iife190_(d_2078_dt__update__tmp_h0_):
                def iife191_(_pat_let60_0):
                    def iife192_(d_2079_dt__update_hcf37_h0_):
                        def iife193_(_pat_let61_0):
                            def iife194_(d_2080_dt__update_hcf36_h0_):
                                return D7_DC19(d_2080_dt__update_hcf36_h0_, d_2079_dt__update_hcf37_h0_, (d_2078_dt__update__tmp_h0_).cf38, (d_2078_dt__update__tmp_h0_).cf39)
                            return iife194_(_pat_let61_0)
                        return iife193_((self).f31)
                    return iife192_(_pat_let60_0)
                return iife191_(pat_let_tv58_)
            return iife190_(_pat_let59_0)
        d_2076_v50_ = _dafny.Map({(iife189_(D7_DC19((self).f31, d_2059_v45_, (d_2002_v3_).fm2(False, (self).f31, globalState), _dafny.Map({(self).f26: len(_dafny.SeqWithoutIsStrInference([d_2003_v4_ for d_2077_i4_ in range(default__.abs(359))]))})))).cf38: (self).f27})
        d_2081_i3_: int
        d_2081_i3_ = 0
        with _dafny.label("11"):
            while ((d_2076_v50_)[p0] if (p0) in (d_2076_v50_) else (self).f27):
                with _dafny.c_label("11"):
                    if (d_2081_i3_) >= (100):
                        raise _dafny.Break("11")
                    d_2081_i3_ = (d_2081_i3_) + (1)
                    d_2082_v51_: _dafny.Seq
                    d_2082_v51_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                    d_2083_v52_: _dafny.Map
                    d_2083_v52_ = _dafny.Map({(self).f26: len(d_2009_v8_)})
                    (globalState).f2 = (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(self).f31 for d_2084_i5_ in range(default__.abs(612))])), (0) - ((self).f31))) > ((0) - ((len(_dafny.Set({(self).f29, (d_2082_v51_)[default__.safeIndex((self).f31, len(d_2082_v51_))]})) if (self).f27 else ((d_2083_v52_)[(self).f26] if ((self).f26) in (d_2083_v52_) else (self).f31))))
                    d_2085_v53_: T1
                    nw325_ = C3()
                    nw325_.ctor__(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(self).f31]), (d_2000_v1_).set(default__.safeIndex((self).f31, len(d_2000_v1_)), -473)]), False, p0)
                    d_2085_v53_ = nw325_
                    d_2086_v54_: D20
                    d_2086_v54_ = D20_DC46(d_2085_v53_)
                    d_2087_v55_: _dafny.Array
                    nw326_ = _dafny.Array(None, 28)
                    nw326_[int(0)] = d_2085_v53_
                    nw326_[int(1)] = d_2085_v53_
                    nw326_[int(2)] = d_2085_v53_
                    nw326_[int(3)] = d_2085_v53_
                    nw326_[int(4)] = d_2085_v53_
                    nw326_[int(5)] = d_2085_v53_
                    nw326_[int(6)] = d_2085_v53_
                    nw326_[int(7)] = d_2085_v53_
                    nw326_[int(8)] = d_2085_v53_
                    nw326_[int(9)] = d_2085_v53_
                    nw326_[int(10)] = d_2085_v53_
                    nw326_[int(11)] = d_2085_v53_
                    nw326_[int(12)] = d_2085_v53_
                    nw326_[int(13)] = d_2085_v53_
                    nw326_[int(14)] = d_2085_v53_
                    nw326_[int(15)] = (d_2086_v54_).cf84
                    nw326_[int(16)] = d_2085_v53_
                    nw326_[int(17)] = d_2085_v53_
                    nw326_[int(18)] = d_2085_v53_
                    nw326_[int(19)] = d_2085_v53_
                    nw326_[int(20)] = d_2085_v53_
                    nw326_[int(21)] = d_2085_v53_
                    nw326_[int(22)] = d_2085_v53_
                    nw326_[int(23)] = d_2085_v53_
                    nw326_[int(24)] = d_2085_v53_
                    nw326_[int(25)] = d_2085_v53_
                    nw326_[int(26)] = d_2085_v53_
                    nw326_[int(27)] = d_2085_v53_
                    d_2087_v55_ = nw326_
                    index301_ = default__.safeIndex(309, (d_2087_v55_).length(0))
                    (d_2087_v55_)[index301_] = d_2085_v53_
                    index302_ = default__.safeIndex(309, (d_2087_v55_).length(0))
                    (d_2087_v55_)[index302_] = d_2085_v53_
                    d_2088_v56_: _dafny.Map
                    d_2088_v56_ = _dafny.Map({(self).f31: (d_2085_v53_).f27})
                    d_2089_v57_: _dafny.Array
                    nw327_ = _dafny.Array(None, 3)
                    nw327_[int(0)] = ((d_2088_v56_)[(self).f31] if ((self).f31) in (d_2088_v56_) else (d_2085_v53_).f27)
                    nw327_[int(1)] = (self).f26
                    nw327_[int(2)] = (self).f27
                    d_2089_v57_ = nw327_
                    d_2090_v58_: D9
                    d_2090_v58_ = D9_DC21(_dafny.Map({p2: d_2089_v57_}))
                    d_2091_v59_: _dafny.Array
                    def lambda100_(d_2092_v52_):
                        def lambda101_(d_2093_i6_):
                            return d_2092_v52_

                        return lambda101_

                    init56_ = lambda100_(d_2083_v52_)
                    nw328_ = _dafny.Array(None, 20)
                    for i0_56_ in range(nw328_.length(0)):
                        nw328_[i0_56_] = init56_(i0_56_)
                    d_2091_v59_ = nw328_
                    d_2094_v60_: _dafny.Map
                    d_2094_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(p2)[default__.safeIndex((self).f31, len(p2))]]): d_2089_v57_})
                    rhs311_ = d_2002_v3_
                    rhs312_ = D9_DC21(d_2094_v60_)
                    rhs313_ = d_2091_v59_
                    rhs314_ = 110
                    lhs282_ = globalState
                    d_2002_v3_ = rhs311_
                    d_2090_v58_ = rhs312_
                    d_2091_v59_ = rhs313_
                    lhs282_.f12 = rhs314_
                    d_2095_v61_: D1
                    d_2095_v61_ = D1_DC5((self).f31)
                    d_2095_v61_ = d_2095_v61_
                    pass
            pass
        d_2096_v62_: _dafny.Array
        def lambda102_(d_2097_i7_):
            return default__.safeDivisionInt(d_2097_i7_, (self).f31)

        init57_ = lambda102_
        nw329_ = _dafny.Array(None, 15)
        for i0_57_ in range(nw329_.length(0)):
            nw329_[i0_57_] = init57_(i0_57_)
        d_2096_v62_ = nw329_
        index303_ = default__.safeIndex(335, (d_2096_v62_).length(0))
        (d_2096_v62_)[index303_] = (self).f31
        d_2098_v63_: _dafny.Array
        nw330_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_2098_v63_ = nw330_
        index304_ = default__.safeIndex(174, (d_2098_v63_).length(0))
        (d_2098_v63_)[index304_] = d_2096_v62_
        d_2099_v64_: _dafny.Seq
        d_2099_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f27})])
        index305_ = default__.safeIndex(335, (d_2096_v62_).length(0))
        index306_ = default__.safeIndex(174, (d_2098_v63_).length(0))
        rhs315_ = (len((d_2099_v64_)[default__.safeIndex((self).f31, len(d_2099_v64_))])) - ((self).f31)
        rhs316_ = (self).f26
        rhs317_ = d_2096_v62_
        rhs318_ = d_2096_v62_
        lhs283_ = d_2096_v62_
        lhs284_ = default__.safeIndex(335, (d_2096_v62_).length(0))
        lhs285_ = globalState
        lhs286_ = d_2098_v63_
        lhs287_ = default__.safeIndex(174, (d_2098_v63_).length(0))
        lhs283_[lhs284_] = rhs315_
        lhs285_.f15 = rhs316_
        lhs286_[lhs287_] = rhs317_
        d_2096_v62_ = rhs318_
        (globalState).f15 = (self).f27

    def m3(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: T0 = None
        d_2100_v0_: _dafny.Seq
        d_2100_v0_ = _dafny.SeqWithoutIsStrInference([p3])
        (globalState).f16 = (d_2100_v0_)[default__.safeIndex((self).f31, len(d_2100_v0_))]
        hi20_ = -364
        for d_2101_i0_ in range(p3, hi20_):
            d_2102_v1_: _dafny.Array
            nw331_ = _dafny.Array(_dafny.Map({}), 29)
            d_2102_v1_ = nw331_
            d_2102_v1_ = d_2102_v1_
            d_2103_v2_: _dafny.Map
            d_2103_v2_ = _dafny.Map({(0) - ((self).f31): (self).f27})
            d_2103_v2_ = (d_2103_v2_).set(p0, self.f30)
            d_2104_v3_: _dafny.Array
            def lambda103_(d_2105_i1_):
                return False

            init58_ = lambda103_
            nw332_ = _dafny.Array(None, 9)
            for i0_58_ in range(nw332_.length(0)):
                nw332_[i0_58_] = init58_(i0_58_)
            d_2104_v3_ = nw332_
            index307_ = default__.safeIndex(913, (d_2104_v3_).length(0))
            (d_2104_v3_)[index307_] = (self).f27
            index308_ = default__.safeIndex(913, (d_2104_v3_).length(0))
            rhs319_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2106_i2_ in range(default__.abs(-29))])) + (p2)
            rhs320_ = ((p2) + (default__.fm22((0) - ((self).f31), self.f30, globalState))) == (p2)
            rhs321_ = (self).f26
            lhs288_ = globalState
            lhs289_ = globalState
            lhs290_ = d_2104_v3_
            lhs291_ = default__.safeIndex(913, (d_2104_v3_).length(0))
            lhs288_.f6 = rhs319_
            lhs289_.f2 = rhs320_
            lhs290_[lhs291_] = rhs321_
            (globalState).f16 = 763
        d_2107_i3_: int
        d_2107_i3_ = 0
        with _dafny.label("12"):
            while not((p3) <= (-399)):
                with _dafny.c_label("12"):
                    if (d_2107_i3_) >= (100):
                        raise _dafny.Break("12")
                    d_2107_i3_ = (d_2107_i3_) + (1)
                    d_2108_v4_: _dafny.Map
                    d_2108_v4_ = _dafny.Map({p1: self.f30})
                    d_2109_v5_: _dafny.Map
                    d_2109_v5_ = _dafny.Map({self.f30: p1})
                    d_2110_v6_: _dafny.Map
                    d_2110_v6_ = _dafny.Map({len(d_2108_v4_): d_2109_v5_})
                    d_2111_v7_: _dafny.MultiSet
                    d_2111_v7_ = _dafny.MultiSet([(self).f26])
                    d_2112_v8_: str
                    d_2112_v8_ = _dafny.CodePoint('o')
                    d_2113_v9_: int
                    d_2114_v10_: int
                    d_2115_v11_: bool
                    out64_: int
                    out65_: int
                    out66_: bool
                    out64_, out65_, out66_ = (self).m6(((d_2110_v6_)[p1] if (p1) in (d_2110_v6_) else default__.fm30((d_2111_v7_).cardinality, True, globalState)), (default__.fm39(globalState)) * (len(_dafny.Map({(self).f26: p0}))), d_2112_v8_, p1, globalState)
                    d_2113_v9_ = out64_
                    d_2114_v10_ = out65_
                    d_2115_v11_ = out66_
                    d_2116_v12_: _dafny.Array
                    def lambda104_(d_2117_p0_):
                        def lambda105_(d_2118_i4_):
                            return (d_2118_i4_) * ((_dafny.MultiSet([(self).f31, d_2117_p0_])).cardinality)

                        return lambda105_

                    init59_ = lambda104_(p0)
                    nw333_ = _dafny.Array(None, 5)
                    for i0_59_ in range(nw333_.length(0)):
                        nw333_[i0_59_] = init59_(i0_59_)
                    d_2116_v12_ = nw333_
                    d_2119_v13_: T3
                    nw334_ = C8()
                    nw334_.ctor__((self).f29)
                    d_2119_v13_ = nw334_
                    d_2120_v14_: _dafny.MultiSet
                    d_2120_v14_ = _dafny.MultiSet([d_2119_v13_])
                    d_2121_v15_: _dafny.Array
                    nw335_ = _dafny.Array(None, 10)
                    nw335_[int(0)] = self.f30
                    nw335_[int(1)] = (self).f26
                    nw335_[int(2)] = (self).f26
                    nw335_[int(3)] = True
                    nw335_[int(4)] = (self).fm2((self).f26, (d_2120_v14_).cardinality, globalState)
                    nw335_[int(5)] = d_2115_v11_
                    nw335_[int(6)] = d_2115_v11_
                    nw335_[int(7)] = self.f30
                    nw335_[int(8)] = self.f30
                    nw335_[int(9)] = (self).f26
                    d_2121_v15_ = nw335_
                    d_2122_v16_: _dafny.Seq
                    d_2122_v16_ = _dafny.SeqWithoutIsStrInference([d_2121_v15_, d_2121_v15_])
                    index309_ = default__.safeIndex(919, (d_2116_v12_).length(0))
                    (d_2116_v12_)[index309_] = ((d_2111_v7_)[d_2115_v11_] if (d_2115_v11_) in (d_2111_v7_) else len(d_2122_v16_))
                    d_2123_v17_: _dafny.Map
                    d_2123_v17_ = _dafny.Map({len(d_2109_v5_): len(p2)})
                    d_2124_v18_: _dafny.Map
                    d_2124_v18_ = _dafny.Map({d_2115_v11_: d_2123_v17_})
                    d_2125_v19_: D21
                    d_2125_v19_ = D21_DC49(d_2124_v18_)
                    index310_ = default__.safeIndex(919, (d_2116_v12_).length(0))
                    (d_2116_v12_)[index310_] = (default__.safeModuloInt(d_2113_v9_, p0)) - (default__.safeDivisionInt(d_2114_v10_, len((d_2125_v19_).cf86)))
                    index311_ = default__.safeIndex(586, (d_2121_v15_).length(0))
                    (d_2121_v15_)[index311_] = d_2115_v11_
                    d_2126_v20_: _dafny.Map
                    d_2126_v20_ = _dafny.Map({(self).f26: d_2112_v8_})
                    index312_ = default__.safeIndex(586, (d_2121_v15_).length(0))
                    (d_2121_v15_)[index312_] = (len(d_2126_v20_)) > (p1)
                    index313_ = default__.safeIndex(586, (d_2121_v15_).length(0))
                    (d_2121_v15_)[index313_] = ((self).f27) not in (_dafny.MultiSet([(self).f27, (self).fm2((self).f27, (d_2116_v12_)[default__.safeIndex(919, (d_2116_v12_).length(0))], globalState), False]))
                    pass
            pass
        if (self).f27:
            d_2127_v21_: _dafny.Array
            nw336_ = _dafny.Array(None, 1)
            nw336_[int(0)] = (self).f26
            d_2127_v21_ = nw336_
            d_2128_v22_: D13
            d_2128_v22_ = D13_DC30(self.f30, (self).f31, self.f30, d_2127_v21_)
            d_2129_v23_: _dafny.Seq
            d_2129_v23_ = _dafny.SeqWithoutIsStrInference([self.f30, self.f30])
            d_2130_v24_: T0
            nw337_ = C6()
            nw337_.ctor__(len(_dafny.Map({True: len(d_2129_v23_)})))
            d_2130_v24_ = nw337_
            d_2131_v25_: _dafny.Map
            d_2131_v25_ = _dafny.Map({(self).f27: d_2130_v24_})
            d_2132_v26_: _dafny.Set
            d_2132_v26_ = _dafny.Set({((d_2131_v25_)[(self).f26] if ((self).f26) in (d_2131_v25_) else d_2130_v24_)})
            d_2133_v27_: D22
            d_2133_v27_ = D22_DC53(d_2132_v26_)
            d_2134_v28_: _dafny.MultiSet
            d_2134_v28_ = _dafny.MultiSet([(self).f27])
            d_2135_v29_: _dafny.Array
            nw338_ = _dafny.Array(None, 18)
            nw338_[int(0)] = (p3) >= (p0)
            nw338_[int(1)] = (self).f27
            nw338_[int(2)] = not ((self).f27) or ((self).f27)
            nw338_[int(3)] = not((d_2128_v22_).cf51)
            nw338_[int(4)] = self.f30
            nw338_[int(5)] = (self).f27
            nw338_[int(6)] = ((d_2133_v27_).cf97) != (d_2132_v26_)
            nw338_[int(7)] = (self).f27
            nw338_[int(8)] = not ((self).f27) or (self.f30)
            nw338_[int(9)] = self.f30
            nw338_[int(10)] = (self).f26
            nw338_[int(11)] = (self).f26
            nw338_[int(12)] = ((self).f26) and ((self).f27)
            nw338_[int(13)] = not((d_2134_v28_).isdisjoint(_dafny.MultiSet(d_2129_v23_)))
            nw338_[int(14)] = self.f30
            nw338_[int(15)] = (((d_2134_v28_)[self.f30] if (self.f30) in (d_2134_v28_) else p3)) < (-253)
            nw338_[int(16)] = (self).fm2((self).f26, len(p2), globalState)
            nw338_[int(17)] = (self).f27
            d_2135_v29_ = nw338_
            d_2135_v29_ = d_2127_v21_
            d_2136_v30_: _dafny.Array
            def lambda106_(d_2137_i5_):
                return (d_2137_i5_) - ((self).f31)

            init60_ = lambda106_
            nw339_ = _dafny.Array(None, 6)
            for i0_60_ in range(nw339_.length(0)):
                nw339_[i0_60_] = init60_(i0_60_)
            d_2136_v30_ = nw339_
            d_2138_v31_: _dafny.Map
            d_2138_v31_ = _dafny.Map({p0: self.f30})
            index314_ = default__.safeIndex(247, (d_2136_v30_).length(0))
            (d_2136_v30_)[index314_] = (952) + (len(d_2138_v31_))
            index315_ = default__.safeIndex(247, (d_2136_v30_).length(0))
            (d_2136_v30_)[index315_] = p0
            d_2139_v32_: _dafny.Array
            nw340_ = _dafny.Array(_dafny.Map({}), 8)
            d_2139_v32_ = nw340_
            d_2140_v33_: _dafny.Map
            d_2140_v33_ = _dafny.Map({p1: (d_2136_v30_)[default__.safeIndex(247, (d_2136_v30_).length(0))]})
            index316_ = default__.safeIndex(462, (d_2139_v32_).length(0))
            (d_2139_v32_)[index316_] = (d_2140_v33_) | (d_2140_v33_)
            index317_ = default__.safeIndex(441, (d_2127_v21_).length(0))
            (d_2127_v21_)[index317_] = not(self.f30)
            d_2141_v34_: D0
            d_2141_v34_ = D0_DC1((d_2129_v23_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f26])), len(d_2129_v23_))], True, p1, (self).f26)
            d_2142_v35_: D0
            d_2142_v35_ = D0_DC2(d_2141_v34_)
            index318_ = default__.safeIndex(462, (d_2139_v32_).length(0))
            index319_ = default__.safeIndex(441, (d_2127_v21_).length(0))
            rhs322_ = d_2140_v33_
            rhs323_ = (self).f27
            rhs324_ = ((d_2136_v30_)[default__.safeIndex(247, (d_2136_v30_).length(0))]) <= (((d_2134_v28_)[(self).f26] if ((self).f26) in (d_2134_v28_) else len(d_2140_v33_)))
            rhs325_ = (self).f31
            rhs326_ = not((_dafny.MultiSet([p3, p3])).issubset(default__.fm10((self).f31, p3, D0_DC2(d_2142_v35_), globalState)))
            lhs292_ = d_2139_v32_
            lhs293_ = default__.safeIndex(462, (d_2139_v32_).length(0))
            lhs294_ = d_2127_v21_
            lhs295_ = default__.safeIndex(441, (d_2127_v21_).length(0))
            lhs296_ = globalState
            lhs297_ = globalState
            lhs292_[lhs293_] = rhs322_
            r0 = rhs323_
            lhs294_[lhs295_] = rhs324_
            lhs296_.f12 = rhs325_
            lhs297_.f22 = rhs326_
            d_2143_v36_: D3
            d_2143_v36_ = D3_DC10(d_2136_v30_)
            d_2144_v37_: _dafny.Map
            d_2144_v37_ = _dafny.Map({d_2143_v36_: p0})
            d_2145_v38_: C4
            nw341_ = C4()
            nw341_.ctor__(d_2144_v37_, (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))], (self).f27)
            d_2145_v38_ = nw341_
            d_2146_v39_: _dafny.MultiSet
            d_2146_v39_ = _dafny.MultiSet([d_2136_v30_])
            rhs327_ = (d_2146_v39_).issubset(d_2146_v39_)
            rhs328_ = (d_2100_v0_)[default__.safeIndex((d_2130_v24_).fm1(globalState), len(d_2100_v0_))]
            rhs329_ = (d_2145_v38_ if (self).f27 else d_2145_v38_)
            rhs330_ = p1
            lhs298_ = self
            lhs299_ = globalState
            lhs300_ = globalState
            lhs298_.f30 = rhs327_
            lhs299_.f12 = rhs328_
            d_2145_v38_ = rhs329_
            lhs300_.f12 = rhs330_
            d_2147_v40_: _dafny.Set
            d_2147_v40_ = _dafny.Set({(d_2130_v24_).fm1(globalState)})
            d_2148_v41_: _dafny.Set
            d_2148_v41_ = _dafny.Set({not((self).f26), (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))]})
            d_2149_v42_: _dafny.Map
            d_2149_v42_ = _dafny.Map({p3: d_2148_v41_})
            d_2150_v43_: D21
            d_2150_v43_ = D21_DC50((self).f26, len(d_2147_v40_), (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))], len(d_2149_v42_))
            source31_ = d_2150_v43_
            if source31_.is_DC50:
                d_2151___mcc_h0_ = source31_.cf87
                d_2152___mcc_h1_ = source31_.cf88
                d_2153___mcc_h2_ = source31_.cf89
                d_2154___mcc_h3_ = source31_.cf90
                d_2155_cf90_ = d_2154___mcc_h3_
                d_2156_cf89_ = d_2153___mcc_h2_
                d_2157_cf88_ = d_2152___mcc_h1_
                d_2158_cf87_ = d_2151___mcc_h0_
                d_2159_v44_: _dafny.Array
                nw342_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_2159_v44_ = nw342_
                index320_ = default__.safeIndex(562, (d_2159_v44_).length(0))
                (d_2159_v44_)[index320_] = d_2127_v21_
                index321_ = default__.safeIndex(441, (d_2127_v21_).length(0))
                index322_ = default__.safeIndex(562, (d_2159_v44_).length(0))
                index323_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                rhs331_ = d_2156_cf89_
                rhs332_ = d_2135_v29_
                rhs333_ = d_2157_cf88_
                rhs334_ = (p2) + (p2)
                rhs335_ = p2
                lhs301_ = d_2127_v21_
                lhs302_ = default__.safeIndex(441, (d_2127_v21_).length(0))
                lhs303_ = d_2159_v44_
                lhs304_ = default__.safeIndex(562, (d_2159_v44_).length(0))
                lhs305_ = d_2136_v30_
                lhs306_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                lhs307_ = globalState
                lhs308_ = globalState
                lhs301_[lhs302_] = rhs331_
                lhs303_[lhs304_] = rhs332_
                lhs305_[lhs306_] = rhs333_
                lhs307_.f6 = rhs334_
                lhs308_.f6 = rhs335_
                d_2160_v45_: _dafny.Array
                nw343_ = _dafny.Array(None, 1)
                d_2160_v45_ = nw343_
                index324_ = default__.safeIndex(519, (d_2160_v45_).length(0))
                (d_2160_v45_)[index324_] = d_2145_v38_
                index325_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                index326_ = default__.safeIndex(519, (d_2160_v45_).length(0))
                rhs336_ = not((p2) < (p2))
                rhs337_ = (self).f31
                rhs338_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icsw"))
                rhs339_ = d_2145_v38_
                lhs309_ = d_2136_v30_
                lhs310_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                lhs311_ = globalState
                lhs312_ = d_2160_v45_
                lhs313_ = default__.safeIndex(519, (d_2160_v45_).length(0))
                r0 = rhs336_
                lhs309_[lhs310_] = rhs337_
                lhs311_.f6 = rhs338_
                lhs312_[lhs313_] = rhs339_
                d_2161_v46_: _dafny.Array
                def lambda107_(d_2162_p3_):
                    def lambda108_(d_2163_i6_):
                        return _dafny.SeqWithoutIsStrInference([d_2162_p3_])

                    return lambda108_

                init61_ = lambda107_(p3)
                nw344_ = _dafny.Array(None, 15)
                for i0_61_ in range(nw344_.length(0)):
                    nw344_[i0_61_] = init61_(i0_61_)
                d_2161_v46_ = nw344_
                index327_ = default__.safeIndex(401, (d_2161_v46_).length(0))
                (d_2161_v46_)[index327_] = _dafny.SeqWithoutIsStrInference([(0) - (len(p2)) for d_2164_i7_ in range(default__.abs(827))])
                d_2165_v47_: D12
                d_2165_v47_ = D12_DC28(p3)
                index328_ = default__.safeIndex(401, (d_2161_v46_).length(0))
                (d_2161_v46_)[index328_] = (d_2100_v0_) + ((_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex((d_2100_v0_)[default__.safeIndex((d_2165_v47_).cf49, len(d_2100_v0_))], len(_dafny.SeqWithoutIsStrInference([p0]))), p3))
                d_2166_v48_: _dafny.Map
                d_2166_v48_ = _dafny.Map({(self).f26: len(_dafny.SeqWithoutIsStrInference([True]))})
                d_2167_v50_: D17
                def iife195_():
                    coll71_ = _dafny.Set()
                    compr_71_: int
                    for compr_71_ in (d_2147_v40_).Elements:
                        d_2168_v49_: int = compr_71_
                        if (d_2168_v49_) in (d_2147_v40_):
                            coll71_ = coll71_.union(_dafny.Set([(d_2168_v49_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxuvuil"))))]))
                    return _dafny.Set(coll71_)
                d_2167_v50_ = D17_DC41(not((iife195_()
).ispropersubset(_dafny.Set({len(d_2166_v48_), p3}))), (self).f26, p0)
                d_2167_v50_ = d_2167_v50_
            elif source31_.is_DC51:
                d_2169___mcc_h4_ = source31_.cf91
                d_2170___mcc_h5_ = source31_.cf92
                d_2171___mcc_h6_ = source31_.cf93
                d_2172___mcc_h7_ = source31_.cf94
                d_2173_cf94_ = d_2172___mcc_h7_
                d_2174_cf93_ = d_2171___mcc_h6_
                d_2175_cf92_ = d_2170___mcc_h5_
                d_2176_cf91_ = d_2169___mcc_h4_
                d_2177_v51_: _dafny.Map
                d_2177_v51_ = _dafny.Map({(self).f27: -449})
                d_2178_v52_: C7
                nw345_ = C7()
                nw345_.ctor__((d_2175_cf92_).f27, _dafny.Map({d_2176_cf91_: d_2177_v51_}))
                d_2178_v52_ = nw345_
                d_2179_v53_: _dafny.Map
                d_2179_v53_ = _dafny.Map({d_2178_v52_: p3})
                index329_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                (d_2136_v30_)[index329_] = (D17_DC42(len(d_2179_v53_), p1, p3)).cf79
                (globalState).f6 = ((p2 if not(self.f30) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hsr"))) if (p2) <= (p2) else p2)
                index330_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                (d_2136_v30_)[index330_] = (p1) + (d_2173_cf94_)
                d_2180_v54_: C1
                nw346_ = C1()
                nw346_.ctor__(default__.safeModuloInt(p3, (self).f31), d_2174_cf93_, (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))], (p0) == (d_2174_cf93_))
                d_2180_v54_ = nw346_
                nw347_ = C1()
                nw347_.ctor__((d_2180_v54_.f37) * (p3), p0, not((d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))]), (d_2175_cf92_).f26)
                d_2180_v54_ = nw347_
            elif source31_.is_DC52:
                d_2181___mcc_h8_ = source31_.cf95
                d_2182___mcc_h9_ = source31_.cf96
                d_2183_cf96_ = d_2182___mcc_h9_
                d_2184_cf95_ = d_2181___mcc_h8_
                (globalState).f12 = p1
                d_2100_v0_ = (d_2100_v0_) + (((d_2100_v0_) + ((d_2100_v0_).set(default__.safeIndex((0) - (p0), len(d_2100_v0_)), p0))).set(default__.safeIndex(p1, len((d_2100_v0_) + ((d_2100_v0_).set(default__.safeIndex((0) - (p0), len(d_2100_v0_)), p0)))), (self).f31))
                d_2185_v55_: _dafny.MultiSet
                d_2185_v55_ = _dafny.MultiSet([d_2100_v0_])
                d_2186_v56_: C3
                nw348_ = C3()
                nw348_.ctor__(d_2185_v55_, True, (self).f27)
                d_2186_v56_ = nw348_
                d_2187_v57_: _dafny.Array
                nw349_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_2187_v57_ = nw349_
                d_2188_v58_: _dafny.Array
                def lambda109_(d_2189_v28_):
                    def lambda110_(d_2190_i8_):
                        return d_2189_v28_

                    return lambda110_

                init62_ = lambda109_(d_2134_v28_)
                nw350_ = _dafny.Array(None, 9)
                for i0_62_ in range(nw350_.length(0)):
                    nw350_[i0_62_] = init62_(i0_62_)
                d_2188_v58_ = nw350_
                index331_ = default__.safeIndex(36, (d_2187_v57_).length(0))
                (d_2187_v57_)[index331_] = d_2188_v58_
                index332_ = default__.safeIndex(36, (d_2187_v57_).length(0))
                rhs340_ = (((d_2100_v0_)[default__.safeIndex(p1, len(d_2100_v0_))]) + (p0)) - ((0) - ((d_2100_v0_)[default__.safeIndex(-797, len(d_2100_v0_))]))
                rhs341_ = d_2188_v58_
                rhs342_ = d_2136_v30_
                rhs343_ = len((_dafny.Set({(d_2134_v28_).cardinality})).intersection(_dafny.Set({p1, (self).f31, p3})))
                lhs314_ = globalState
                lhs315_ = d_2187_v57_
                lhs316_ = default__.safeIndex(36, (d_2187_v57_).length(0))
                lhs317_ = globalState
                lhs314_.f16 = rhs340_
                lhs315_[lhs316_] = rhs341_
                d_2136_v30_ = rhs342_
                lhs317_.f16 = rhs343_
            elif True:
                d_2191___mcc_h10_ = source31_.cf86
                d_2192_cf86_ = d_2191___mcc_h10_
                d_2193_v59_: C4
                nw351_ = C4()
                nw351_.ctor__(((d_2145_v38_).f33) | ((d_2145_v38_).f33), (p2) != (p2), self.f30)
                d_2193_v59_ = nw351_
                d_2194_v60_: _dafny.Map
                d_2194_v60_ = _dafny.Map({(d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))]: (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))]})
                d_2195_v61_: C4
                nw352_ = C4()
                nw352_.ctor__(((d_2145_v38_).f33 if self.f30 else (d_2193_v59_).f33), (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))], ((d_2194_v60_)[not((d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))])] if (not((d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))])) in (d_2194_v60_) else (self).f26))
                d_2195_v61_ = nw352_
                d_2196_v62_: T1
                nw353_ = C4()
                nw353_.ctor__((d_2193_v59_).f33, (self).f27, (self).f26)
                d_2196_v62_ = nw353_
                d_2197_v63_: D21
                d_2197_v63_ = D21_DC51((self).f31, d_2196_v62_, 875, (self).f31)
                index333_ = default__.safeIndex(247, (d_2136_v30_).length(0))
                (d_2136_v30_)[index333_] = (d_2197_v63_).cf91
                d_2198_v64_: str
                d_2198_v64_ = _dafny.CodePoint('r')
                d_2199_v65_: _dafny.Set
                d_2199_v65_ = _dafny.Set({d_2198_v64_, d_2198_v64_})
                d_2200_v66_: C6
                nw354_ = C6()
                nw354_.ctor__(len(d_2199_v65_))
                d_2200_v66_ = nw354_
                d_2201_v67_: _dafny.MultiSet
                d_2201_v67_ = _dafny.MultiSet([d_2200_v66_])
                d_2202_v68_: _dafny.Map
                d_2202_v68_ = _dafny.Map({d_2201_v67_: p3})
                d_2203_v69_: D16
                d_2203_v69_ = D16_DC37(d_2138_v31_)
                d_2204_v71_: _dafny.Map
                d_2204_v71_ = _dafny.Map({not(False): d_2198_v64_})
                d_2205_v72_: D14
                d_2205_v72_ = D14_DC32(_dafny.MultiSet([(d_2196_v62_).f27]), (self).f26, (self).f27, (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))], len((d_2204_v71_).set((d_2196_v62_).f27, d_2198_v64_)))
                d_2206_v73_: _dafny.Map
                d_2206_v73_ = _dafny.Map({d_2205_v72_: (d_2196_v62_).fm2(self.f30, p1, globalState)})
                d_2207_v74_: _dafny.Map
                d_2207_v74_ = _dafny.Map({(self).f31: d_2206_v73_})
                d_2208_v75_: _dafny.Array
                nw355_ = _dafny.Array(None, 6)
                nw355_[int(0)] = (d_2138_v31_).set(len(d_2202_v68_), (self).f26)
                nw355_[int(1)] = d_2138_v31_
                nw355_[int(2)] = d_2138_v31_
                nw355_[int(3)] = (d_2138_v31_) | ((d_2203_v69_).cf68)
                nw355_[int(4)] = _dafny.Map({(d_2193_v59_).fm1(globalState): False})
                def iife196_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in (d_2207_v74_).keys.Elements:
                        d_2209_v70_: int = compr_72_
                        if (d_2209_v70_) in (d_2207_v74_):
                            coll72_[(d_2209_v70_) + (p1)] = (d_2196_v62_).f26
                    return _dafny.Map(coll72_)
                nw355_[int(5)] = iife196_()
                
                d_2208_v75_ = nw355_
                index334_ = default__.safeIndex(341, (d_2208_v75_).length(0))
                (d_2208_v75_)[index334_] = d_2138_v31_
                index335_ = default__.safeIndex(341, (d_2208_v75_).length(0))
                (d_2208_v75_)[index335_] = (d_2138_v31_ if (d_2127_v21_)[default__.safeIndex(441, (d_2127_v21_).length(0))] else (d_2138_v31_) | (d_2138_v31_))
        elif True:
            d_2210_v76_: _dafny.Set
            d_2210_v76_ = _dafny.Set({self.f30})
            d_2211_v77_: D6
            d_2211_v77_ = D6_DC17(len(d_2210_v76_), p0, p3, p1, (0) - (p0))
            d_2212_v78_: _dafny.Map
            d_2212_v78_ = _dafny.Map({False: (self).f26})
            d_2213_v79_: D16
            d_2213_v79_ = D16_DC38(len(d_2212_v78_), p1, self.f30)
            d_2214_v80_: _dafny.Map
            d_2214_v80_ = _dafny.Map({self.f30: len(_dafny.Set({d_2213_v79_}))})
            d_2215_v81_: D0
            d_2215_v81_ = D0_DC0(self.f30)
            d_2216_v82_: _dafny.Map
            d_2216_v82_ = _dafny.Map({(self).fm4(d_2215_v81_, p0, (0) - (p3), globalState): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})
            d_2217_v83_: _dafny.MultiSet
            d_2217_v83_ = _dafny.MultiSet([(self).f27])
            d_2218_v84_: _dafny.Set
            def iife197_(_pat_let62_0):
                def iife198_(d_2219_dt__update__tmp_h1_):
                    def iife199_(_pat_let63_0):
                        def iife200_(d_2220_dt__update_hcf36_h1_):
                            return D7_DC19(d_2220_dt__update_hcf36_h1_, (d_2219_dt__update__tmp_h1_).cf37, (d_2219_dt__update__tmp_h1_).cf38, (d_2219_dt__update__tmp_h1_).cf39)
                        return iife200_(_pat_let63_0)
                    return iife199_((self).f31)
                return iife198_(_pat_let62_0)
            def iife201_(_pat_let64_0):
                def iife202_(d_2221_dt__update__tmp_h0_):
                    def iife203_(_pat_let65_0):
                        def iife204_(d_2222_dt__update_hcf36_h0_):
                            return D7_DC19(d_2222_dt__update_hcf36_h0_, (d_2221_dt__update__tmp_h0_).cf37, (d_2221_dt__update__tmp_h0_).cf38, (d_2221_dt__update__tmp_h0_).cf39)
                        return iife204_(_pat_let65_0)
                    return iife203_((self).f31)
                return iife202_(_pat_let64_0)
            d_2218_v84_ = _dafny.Set({(self).f31, ((d_2217_v83_)[(self).fm2((iife197_(D7_DC19(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smafhihgt"))), d_2217_v83_, (self).f26, d_2214_v80_))).cf38, p0, globalState)] if ((self).fm2((iife201_(D7_DC19(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smafhihgt"))), d_2217_v83_, (self).f26, d_2214_v80_))).cf38, p0, globalState)) in (d_2217_v83_) else p0), p0})
            d_2223_v85_: _dafny.Map
            d_2223_v85_ = _dafny.Map({(self).f31: (self).f31})
            d_2211_v77_ = default__.fm53(((_dafny.MultiSet([p3, p3, len(d_2214_v80_)])).cardinality if (self).f27 else len(((d_2216_v82_)[p0] if (p0) in (d_2216_v82_) else p2))), p1, d_2218_v84_, d_2223_v85_, globalState)
            (globalState).f24 = ((p2) + (p2)) < ((p2) + (p2))
            (globalState).f15 = (self).f27
            (globalState).f16 = (p3) + (p1)
            d_2224_v86_: _dafny.Seq
            d_2224_v86_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1 for d_2225_i9_ in range(default__.abs(704))])])
            d_2226_v87_: C2
            nw356_ = C2()
            nw356_.ctor__((d_2224_v86_).set(default__.safeIndex((0) - ((self).f31), len(d_2224_v86_)), _dafny.SeqWithoutIsStrInference([p0])), (self).f27, (self).f27)
            d_2226_v87_ = nw356_
        d_2227_v88_: _dafny.Array
        nw357_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
        d_2227_v88_ = nw357_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_2227_v88_).length(0)):
            d_2228_i10_: int = guard_loop_1_
            if (True) and (((0) <= (d_2228_i10_)) and ((d_2228_i10_) < ((d_2227_v88_).length(0)))):
                (d_2227_v88_)[(d_2228_i10_)] = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('u') if self.f30 else _dafny.CodePoint('e')) for d_2229_i11_ in range(default__.abs(680))])
        d_2230_v89_: D0
        d_2230_v89_ = D0_DC0(self.f30)
        d_2231_v90_: _dafny.Seq
        d_2231_v90_ = _dafny.SeqWithoutIsStrInference([not(not(False))])
        d_2232_v91_: _dafny.Map
        d_2232_v91_ = _dafny.Map({self.f30: (self).f31})
        d_2233_v92_: _dafny.Array
        nw358_ = _dafny.Array(None, 26)
        nw358_[int(0)] = True
        nw358_[int(1)] = (self).fm2((self).f26, (self).f31, globalState)
        nw358_[int(2)] = self.f30
        nw358_[int(3)] = (-543) > (p1)
        nw358_[int(4)] = not ((self).f27) or (self.f30)
        nw358_[int(5)] = (d_2230_v89_).cf0
        nw358_[int(6)] = (d_2231_v90_)[default__.safeIndex(len(d_2232_v91_), len(d_2231_v90_))]
        nw358_[int(7)] = (d_2100_v0_) <= (d_2100_v0_)
        nw358_[int(8)] = ((self).fm1(globalState)) <= ((0) - (p1))
        nw358_[int(9)] = (self).f27
        nw358_[int(10)] = (self).f27
        nw358_[int(11)] = (self).f26
        nw358_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2234_i13_ in range(default__.abs(14))]))
        nw358_[int(13)] = self.f30
        nw358_[int(14)] = True
        nw358_[int(15)] = (self).f27
        nw358_[int(16)] = (self).f27
        nw358_[int(17)] = (self).f27
        nw358_[int(18)] = (self).f26
        nw358_[int(19)] = self.f30
        nw358_[int(20)] = (self).f26
        nw358_[int(21)] = not ((self).f27) or (self.f30)
        nw358_[int(22)] = self.f30
        nw358_[int(23)] = self.f30
        nw358_[int(24)] = self.f30
        nw358_[int(25)] = False
        d_2233_v92_ = nw358_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_2233_v92_).length(0)):
            d_2235_i12_: int = guard_loop_2_
            if (True) and (((0) <= (d_2235_i12_)) and ((d_2235_i12_) < ((d_2233_v92_).length(0)))):
                (d_2233_v92_)[(d_2235_i12_)] = not(((self).f26 if (self).f27 else False))
        r0 = ((self).fm2((self).f27, p0, globalState)) == ((_dafny.MultiSet(d_2100_v0_)) == ((self).f29))
        d_2236_v93_: C0
        nw359_ = C0()
        nw359_.ctor__((self).f27)
        d_2236_v93_ = nw359_
        d_2237_v94_: _dafny.Map
        d_2237_v94_ = _dafny.Map({self.f30: d_2236_v93_})
        r1 = (len(d_2237_v94_)) >= ((self).f31)
        d_2238_v95_: _dafny.Array
        nw360_ = _dafny.Array(None, 5)
        nw360_[int(0)] = p3
        nw360_[int(1)] = (self).f31
        nw360_[int(2)] = (self).f31
        nw360_[int(3)] = ((self).f31 if self.f30 else (self).f31)
        nw360_[int(4)] = (self).f31
        d_2238_v95_ = nw360_
        r2 = d_2238_v95_
        d_2239_v96_: T0
        nw361_ = C6()
        nw361_.ctor__(p0)
        d_2239_v96_ = nw361_
        r3 = d_2239_v96_
        return r0, r1, r2, r3

    def m4(self, globalState):
        r0: bool = False
        r1: bool = False
        d_2240_v0_: D6
        d_2240_v0_ = D6_DC16(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2241_i1_ in range(default__.abs(-172))]))
        d_2242_i0_: int
        d_2242_i0_ = 0
        with _dafny.label("13"):
            def lambda111_(source32_):
                if source32_.is_DC17:
                    d_2254___mcc_h0_ = source32_.cf30
                    d_2255___mcc_h1_ = source32_.cf31
                    d_2256___mcc_h2_ = source32_.cf32
                    d_2257___mcc_h3_ = source32_.cf33
                    d_2258___mcc_h4_ = source32_.cf34
                    d_2259_cf34_ = d_2258___mcc_h4_
                    d_2260_cf33_ = d_2257___mcc_h3_
                    d_2261_cf32_ = d_2256___mcc_h2_
                    d_2262_cf31_ = d_2255___mcc_h1_
                    d_2263_cf30_ = d_2254___mcc_h0_
                    return True
                elif True:
                    d_2264___mcc_h5_ = source32_.cf29
                    d_2265_cf29_ = d_2264___mcc_h5_
                    return not(True)

            while lambda111_(d_2240_v0_):
                with _dafny.c_label("13"):
                    if (d_2242_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_2242_i0_ = (d_2242_i0_) + (1)
                    d_2243_v1_: _dafny.Map
                    d_2243_v1_ = _dafny.Map({(self).f26: (self).f27})
                    d_2244_v2_: _dafny.Seq
                    d_2244_v2_ = _dafny.SeqWithoutIsStrInference([(d_2243_v1_) | (d_2243_v1_), d_2243_v1_])
                    d_2244_v2_ = d_2244_v2_
                    d_2245_v3_: _dafny.Set
                    d_2245_v3_ = _dafny.Set({(self).f31})
                    d_2246_v4_: _dafny.Seq
                    d_2246_v4_ = _dafny.SeqWithoutIsStrInference([-469, len(_dafny.Map({len(d_2245_v3_): (self).f26})), (self).f31])
                    d_2247_v5_: _dafny.Seq
                    d_2247_v5_ = _dafny.SeqWithoutIsStrInference([d_2246_v4_, d_2246_v4_, d_2246_v4_])
                    d_2248_v6_: C2
                    nw362_ = C2()
                    nw362_.ctor__(d_2247_v5_, self.f30, (self).f26)
                    d_2248_v6_ = nw362_
                    (globalState).f2 = (_dafny.MultiSet((d_2246_v4_) + (d_2246_v4_))).issubset(((self).f29) | ((self).f29))
                    d_2249_v7_: _dafny.Seq
                    d_2249_v7_ = _dafny.SeqWithoutIsStrInference([self.f30])
                    d_2250_v8_: _dafny.Map
                    d_2250_v8_ = _dafny.Map({(self).f31: (d_2248_v6_).fm1(globalState)})
                    d_2251_v9_: _dafny.Map
                    d_2251_v9_ = _dafny.Map({(self).f27: len(d_2250_v8_)})
                    d_2252_v10_: _dafny.Map
                    d_2252_v10_ = _dafny.Map({len(d_2249_v7_): d_2251_v9_})
                    d_2253_v11_: C7
                    nw363_ = C7()
                    nw363_.ctor__(self.f30, d_2252_v10_)
                    d_2253_v11_ = nw363_
                    pass
            pass
        d_2266_v12_: C6
        nw364_ = C6()
        nw364_.ctor__((self).f31)
        d_2266_v12_ = nw364_
        d_2267_v13_: _dafny.Set
        d_2267_v13_ = _dafny.Set({self.f30})
        d_2267_v13_ = d_2267_v13_
        if (self).f27:
            d_2268_v14_: D16
            d_2268_v14_ = D16_DC39(default__.fm54(True, (self).f31, d_2266_v12_.f40, globalState))
            source33_ = d_2268_v14_
            if source33_.is_DC38:
                d_2269___mcc_h6_ = source33_.cf69
                d_2270___mcc_h7_ = source33_.cf70
                d_2271___mcc_h8_ = source33_.cf71
                d_2272_cf71_ = d_2271___mcc_h8_
                d_2273_cf70_ = d_2270___mcc_h7_
                d_2274_cf69_ = d_2269___mcc_h6_
                d_2275_v15_: _dafny.Seq
                d_2275_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq"))
                d_2276_v16_: D3
                d_2276_v16_ = D3_DC11(default__.fm24(d_2275_v15_, (self).f26, d_2273_cf70_, globalState))
                d_2277_v17_: _dafny.Array
                def lambda112_(d_2278_i2_):
                    return not(False)

                init63_ = lambda112_
                nw365_ = _dafny.Array(None, 27)
                for i0_63_ in range(nw365_.length(0)):
                    nw365_[i0_63_] = init63_(i0_63_)
                d_2277_v17_ = nw365_
                index336_ = default__.safeIndex(77, (d_2277_v17_).length(0))
                (d_2277_v17_)[index336_] = self.f30
                d_2279_v18_: str
                d_2279_v18_ = _dafny.CodePoint('f')
                d_2280_v19_: _dafny.Map
                d_2280_v19_ = _dafny.Map({d_2274_cf69_: True})
                d_2281_v20_: _dafny.Array
                nw366_ = _dafny.Array(None, 19)
                nw366_[int(0)] = d_2266_v12_.f40
                nw366_[int(1)] = (0) - (d_2266_v12_.f40)
                nw366_[int(2)] = d_2266_v12_.f40
                nw366_[int(3)] = len(d_2280_v19_)
                nw366_[int(4)] = (self).f31
                nw366_[int(5)] = d_2266_v12_.f40
                nw366_[int(6)] = -305
                nw366_[int(7)] = 635
                nw366_[int(8)] = d_2274_cf69_
                nw366_[int(9)] = d_2273_cf70_
                nw366_[int(10)] = d_2266_v12_.f40
                nw366_[int(11)] = d_2273_cf70_
                nw366_[int(12)] = d_2273_cf70_
                nw366_[int(13)] = d_2266_v12_.f40
                nw366_[int(14)] = d_2274_cf69_
                nw366_[int(15)] = d_2266_v12_.f40
                nw366_[int(16)] = (self).f31
                nw366_[int(17)] = len(d_2275_v15_)
                nw366_[int(18)] = (self).f31
                d_2281_v20_ = nw366_
                d_2282_v21_: _dafny.MultiSet
                d_2282_v21_ = _dafny.MultiSet([d_2281_v20_, d_2281_v20_])
                d_2283_v22_: _dafny.Set
                d_2283_v22_ = _dafny.Set({d_2274_cf69_, ((d_2282_v21_)[d_2281_v20_] if (d_2281_v20_) in (d_2282_v21_) else d_2266_v12_.f40), d_2266_v12_.f40})
                index337_ = default__.safeIndex(77, (d_2277_v17_).length(0))
                rhs344_ = D3_DC11(d_2279_v18_)
                rhs345_ = (d_2267_v13_).ispropersubset((d_2267_v13_) - (_dafny.Set({self.f30, self.f30, (self).f26})))
                rhs346_ = d_2266_v12_.f40
                rhs347_ = len(d_2283_v22_)
                lhs318_ = d_2277_v17_
                lhs319_ = default__.safeIndex(77, (d_2277_v17_).length(0))
                lhs320_ = globalState
                d_2276_v16_ = rhs344_
                lhs318_[lhs319_] = rhs345_
                d_2274_cf69_ = rhs346_
                lhs320_.f16 = rhs347_
                (d_2266_v12_).f40 = d_2266_v12_.f40
                (globalState).f16 = d_2274_cf69_
                d_2284_v23_: _dafny.Seq
                d_2284_v23_ = _dafny.SeqWithoutIsStrInference([(self).f31])
                (globalState).f15 = (d_2274_cf69_) in (_dafny.SeqWithoutIsStrInference([len(d_2284_v23_)]))
            elif source33_.is_DC37:
                d_2285___mcc_h9_ = source33_.cf68
                d_2286_cf68_ = d_2285___mcc_h9_
                d_2287_v24_: _dafny.Array
                def lambda113_(d_2288_i3_):
                    return (self).f26

                init64_ = lambda113_
                nw367_ = _dafny.Array(None, 18)
                for i0_64_ in range(nw367_.length(0)):
                    nw367_[i0_64_] = init64_(i0_64_)
                d_2287_v24_ = nw367_
                d_2289_v25_: _dafny.Set
                d_2289_v25_ = _dafny.Set({(self).f31})
                d_2290_v26_: D12
                d_2290_v26_ = D12_DC28(d_2266_v12_.f40)
                d_2291_v27_: _dafny.Array
                nw368_ = _dafny.Array(None, 2)
                nw368_[int(0)] = len(d_2289_v25_)
                nw368_[int(1)] = (d_2290_v26_).cf49
                d_2291_v27_ = nw368_
                index338_ = default__.safeIndex(1, (d_2287_v24_).length(0))
                (d_2287_v24_)[index338_] = not((len(_dafny.SeqWithoutIsStrInference([d_2291_v27_, d_2291_v27_]))) <= ((self).f31))
                index339_ = default__.safeIndex(1, (d_2287_v24_).length(0))
                (d_2287_v24_)[index339_] = (self).f27
                (globalState).f2 = not((d_2266_v12_).fm44(globalState))
                index340_ = default__.safeIndex(117, (d_2291_v27_).length(0))
                (d_2291_v27_)[index340_] = default__.safeModuloInt((self).f31, -769)
                index341_ = default__.safeIndex(117, (d_2291_v27_).length(0))
                (d_2291_v27_)[index341_] = d_2266_v12_.f40
                (globalState).f2 = ((d_2287_v24_)[default__.safeIndex(1, (d_2287_v24_).length(0))]) == (not(self.f30))
            elif True:
                d_2292___mcc_h10_ = source33_.cf72
                d_2293_cf72_ = d_2292___mcc_h10_
                d_2294_v28_: _dafny.Seq
                d_2294_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "veta"))
                d_2295_v29_: _dafny.Map
                d_2295_v29_ = _dafny.Map({(self).f26: len(d_2294_v28_)})
                d_2296_v30_: str
                d_2296_v30_ = _dafny.CodePoint('t')
                d_2297_v31_: _dafny.MultiSet
                d_2297_v31_ = _dafny.MultiSet([_dafny.CodePoint('q'), d_2296_v30_])
                d_2295_v29_ = (d_2295_v29_).set((d_2297_v31_).issubset(_dafny.MultiSet([d_2296_v30_, d_2296_v30_, d_2296_v30_])), default__.safeModuloInt(d_2266_v12_.f40, (0) - (d_2266_v12_.f40)))
                r1 = not((d_2266_v12_).fm44(globalState))
                d_2298_v32_: _dafny.Array
                def lambda114_(d_2299_v12_):
                    def lambda115_(d_2300_i4_):
                        return (d_2300_i4_) + (d_2299_v12_.f40)

                    return lambda115_

                init65_ = lambda114_(d_2266_v12_)
                nw369_ = _dafny.Array(None, 21)
                for i0_65_ in range(nw369_.length(0)):
                    nw369_[i0_65_] = init65_(i0_65_)
                d_2298_v32_ = nw369_
                d_2301_v33_: _dafny.Seq
                d_2301_v33_ = _dafny.SeqWithoutIsStrInference([True, (self).f26, self.f30, (self).f26])
                d_2302_v34_: _dafny.Seq
                d_2302_v34_ = _dafny.SeqWithoutIsStrInference([len(d_2301_v33_), d_2266_v12_.f40])
                d_2303_v35_: _dafny.Set
                d_2303_v35_ = _dafny.Set({d_2266_v12_.f40, (self).fm1(globalState)})
                nw370_ = _dafny.Array(None, 4)
                nw370_[int(0)] = (0) - ((d_2302_v34_)[default__.safeIndex(default__.fm39(globalState), len(d_2302_v34_))])
                nw370_[int(1)] = ((0) - (len(d_2303_v35_))) - ((self).f31)
                nw370_[int(2)] = len(d_2294_v28_)
                nw370_[int(3)] = (self).f31
                d_2298_v32_ = nw370_
                d_2304_v36_: _dafny.Array
                nw371_ = _dafny.Array(False, 10)
                d_2304_v36_ = nw371_
                index342_ = default__.safeIndex(902, (d_2304_v36_).length(0))
                (d_2304_v36_)[index342_] = self.f30
                index343_ = default__.safeIndex(902, (d_2304_v36_).length(0))
                (d_2304_v36_)[index343_] = (not((self).f27) if (self).f26 else (self).f27)
            d_2305_v37_: _dafny.Seq
            d_2305_v37_ = _dafny.SeqWithoutIsStrInference([d_2266_v12_.f40, default__.safeModuloInt(d_2266_v12_.f40, ((self).f29).cardinality)])
            d_2306_v38_: _dafny.Seq
            d_2306_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivqetbs"))
            (globalState).f12 = (d_2305_v37_)[default__.safeIndex(len(d_2306_v38_), len(d_2305_v37_))]
            d_2307_v39_: _dafny.Array
            def lambda116_(d_2308_v38_):
                def lambda117_(d_2309_i5_):
                    return (len(d_2308_v38_)) <= ((self).f31)

                return lambda117_

            init66_ = lambda116_(d_2306_v38_)
            nw372_ = _dafny.Array(None, 2)
            for i0_66_ in range(nw372_.length(0)):
                nw372_[i0_66_] = init66_(i0_66_)
            d_2307_v39_ = nw372_
            index344_ = default__.safeIndex(191, (d_2307_v39_).length(0))
            (d_2307_v39_)[index344_] = (d_2266_v12_.f40) <= (788)
            index345_ = default__.safeIndex(191, (d_2307_v39_).length(0))
            (d_2307_v39_)[index345_] = (self.f30 if ((self).f31) > (((self).f29).cardinality) else (self).f26)
            d_2310_v40_: _dafny.Map
            d_2310_v40_ = _dafny.Map({d_2266_v12_.f40: not(self.f30)})
            d_2310_v40_ = d_2310_v40_
            (globalState).f15 = (d_2266_v12_).fm44(globalState)
        elif True:
            d_2311_v41_: C5
            nw373_ = C5()
            nw373_.ctor__((self).f27, not((self).f26))
            d_2311_v41_ = nw373_
            d_2311_v41_ = d_2311_v41_
            (globalState).f16 = default__.safeDivisionInt(d_2266_v12_.f40, (d_2266_v12_.f40) + (844))
            d_2312_v42_: D7
            d_2312_v42_ = D7_DC19(d_2266_v12_.f40, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([self.f30, (self).f26])), (D0_DC1((self).f27, (self).f26, 344, self.f30)).cf4, default__.fm30(d_2266_v12_.f40, (self).f27, globalState))
            d_2313_v43_: _dafny.Map
            d_2313_v43_ = _dafny.Map({(d_2266_v12_).fm44(globalState): (self).f31})
            pat_let_tv59_ = d_2313_v43_
            d_2314_v44_: _dafny.Map
            def iife205_(_pat_let66_0):
                def iife206_(d_2315_dt__update__tmp_h0_):
                    def iife207_(_pat_let67_0):
                        def iife208_(d_2316_dt__update_hcf39_h0_):
                            return D7_DC19((d_2315_dt__update__tmp_h0_).cf36, (d_2315_dt__update__tmp_h0_).cf37, (d_2315_dt__update__tmp_h0_).cf38, d_2316_dt__update_hcf39_h0_)
                        return iife208_(_pat_let67_0)
                    return iife207_(pat_let_tv59_)
                return iife206_(_pat_let66_0)
            d_2314_v44_ = _dafny.Map({(iife205_(d_2312_v42_)).cf39: ((self).f27) or (self.f30)})
            d_2314_v44_ = (d_2314_v44_).set((d_2313_v43_).set(self.f30, 990), not(True))
            if (self).f26:
                d_2317_v45_: _dafny.Array
                nw374_ = _dafny.Array(int(0), 25)
                d_2317_v45_ = nw374_
                index346_ = default__.safeIndex(166, (d_2317_v45_).length(0))
                (d_2317_v45_)[index346_] = (self).f31
                index347_ = default__.safeIndex(166, (d_2317_v45_).length(0))
                (d_2317_v45_)[index347_] = 537
                d_2318_v46_: _dafny.Seq
                d_2318_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohnwk"))
                d_2319_v47_: _dafny.Map
                d_2319_v47_ = _dafny.Map({-558: d_2318_v46_})
                d_2320_v48_: C1
                nw375_ = C1()
                nw375_.ctor__(len(((d_2319_v47_)[d_2266_v12_.f40] if (d_2266_v12_.f40) in (d_2319_v47_) else d_2318_v46_)), d_2266_v12_.f40, (self).f27, (self).f27)
                d_2320_v48_ = nw375_
                d_2321_v49_: _dafny.Map
                d_2321_v49_ = _dafny.Map({(d_2320_v48_).f36: d_2320_v48_})
                d_2322_v50_: _dafny.Seq
                d_2322_v50_ = _dafny.SeqWithoutIsStrInference([d_2320_v48_, d_2320_v48_, ((d_2321_v49_)[(self).f31] if ((self).f31) in (d_2321_v49_) else d_2320_v48_), d_2320_v48_])
                d_2323_v51_: _dafny.Map
                d_2323_v51_ = _dafny.Map({d_2322_v50_: d_2320_v48_.f37})
                d_2323_v51_ = d_2323_v51_
                d_2324_v52_: _dafny.Seq
                d_2324_v52_ = _dafny.SeqWithoutIsStrInference([len(d_2318_v46_)])
                d_2325_v53_: _dafny.Seq
                d_2325_v53_ = _dafny.SeqWithoutIsStrInference([len(d_2324_v52_)])
                d_2326_v54_: D0
                d_2326_v54_ = D0_DC0(not(True))
                d_2327_v55_: _dafny.MultiSet
                d_2327_v55_ = _dafny.MultiSet([(d_2324_v52_).set(default__.safeIndex((self).fm4(d_2326_v54_, 713, 453, globalState), len(d_2324_v52_)), d_2266_v12_.f40), _dafny.SeqWithoutIsStrInference([-911]), d_2324_v52_])
                d_2328_v56_: C3
                nw376_ = C3()
                nw376_.ctor__(d_2327_v55_, (self).f27, (self).f26)
                d_2328_v56_ = nw376_
                d_2329_v57_: _dafny.Map
                d_2329_v57_ = _dafny.Map({(self).f26: d_2328_v56_})
                d_2330_v58_: str
                d_2330_v58_ = _dafny.CodePoint('s')
                d_2331_v59_: D23
                d_2331_v59_ = D23_DC55(d_2328_v56_)
                d_2328_v56_ = ((d_2329_v57_)[(d_2330_v58_) not in ((d_2318_v46_).set(default__.safeIndex(d_2266_v12_.f40, len(d_2318_v46_)), d_2330_v58_))] if ((d_2330_v58_) not in ((d_2318_v46_).set(default__.safeIndex(d_2266_v12_.f40, len(d_2318_v46_)), d_2330_v58_))) in (d_2329_v57_) else ((d_2329_v57_)[(self).f26] if ((self).f26) in (d_2329_v57_) else (d_2331_v59_).cf103))
                (globalState).f2 = ((self).f27) and (self.f30)
                (globalState).f12 = default__.safeModuloInt(default__.safeDivisionInt((d_2317_v45_)[default__.safeIndex(166, (d_2317_v45_).length(0))], d_2266_v12_.f40), d_2266_v12_.f40)
            elif True:
                d_2332_v60_: _dafny.MultiSet
                d_2332_v60_ = _dafny.MultiSet([d_2266_v12_.f40])
                d_2332_v60_ = d_2332_v60_
                d_2333_v61_: _dafny.MultiSet
                d_2333_v61_ = _dafny.MultiSet([D1_DC3(_dafny.SeqWithoutIsStrInference([(self).f26]))])
                d_2334_v62_: _dafny.Seq
                d_2334_v62_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f27])
                d_2335_v63_: D1
                d_2335_v63_ = D1_DC3(d_2334_v62_)
                pat_let_tv60_ = d_2334_v62_
                d_2336_v64_: _dafny.Seq
                def iife209_(_pat_let68_0):
                    def iife210_(d_2337_dt__update__tmp_h1_):
                        def iife211_(_pat_let69_0):
                            def iife212_(d_2338_dt__update_hcf6_h0_):
                                return D1_DC3(d_2338_dt__update_hcf6_h0_)
                            return iife212_(_pat_let69_0)
                        return iife211_(pat_let_tv60_)
                    return iife210_(_pat_let68_0)
                d_2336_v64_ = _dafny.SeqWithoutIsStrInference([iife209_(d_2335_v63_), d_2335_v63_])
                d_2339_v65_: _dafny.Set
                d_2339_v65_ = _dafny.Set({len(d_2334_v62_)})
                d_2340_v66_: _dafny.Map
                d_2340_v66_ = _dafny.Map({d_2339_v65_: 979})
                d_2341_v67_: D24
                d_2341_v67_ = D24_DC59(d_2333_v61_)
                pat_let_tv61_ = d_2334_v62_
                pat_let_tv62_ = d_2334_v62_
                d_2342_v68_: _dafny.Array
                nw377_ = _dafny.Array(None, 27)
                nw377_[int(0)] = d_2333_v61_
                nw377_[int(1)] = (d_2333_v61_) | (d_2333_v61_)
                nw377_[int(2)] = d_2333_v61_
                nw377_[int(3)] = (d_2333_v61_).set(d_2335_v63_, default__.abs(len(d_2334_v62_)))
                nw377_[int(4)] = ((d_2333_v61_).set(d_2335_v63_, default__.abs(d_2266_v12_.f40))).intersection(d_2333_v61_)
                nw377_[int(5)] = d_2333_v61_
                nw377_[int(6)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2335_v63_, d_2335_v63_, d_2335_v63_]))) | (d_2333_v61_)
                nw377_[int(7)] = d_2333_v61_
                def iife213_(_pat_let70_0):
                    def iife214_(d_2343_dt__update__tmp_h2_):
                        def iife215_(_pat_let71_0):
                            def iife216_(d_2344_dt__update_hcf6_h1_):
                                return D1_DC3(d_2344_dt__update_hcf6_h1_)
                            return iife216_(_pat_let71_0)
                        return iife215_(pat_let_tv61_)
                    return iife214_(_pat_let70_0)
                nw377_[int(8)] = (_dafny.MultiSet([iife213_(d_2335_v63_)])) | (d_2333_v61_)
                nw377_[int(9)] = _dafny.MultiSet(d_2336_v64_)
                nw377_[int(10)] = (d_2333_v61_) | (d_2333_v61_)
                def iife217_(_pat_let72_0):
                    def iife218_(d_2345_dt__update__tmp_h3_):
                        def iife219_(_pat_let73_0):
                            def iife220_(d_2346_dt__update_hcf6_h2_):
                                return D1_DC3(d_2346_dt__update_hcf6_h2_)
                            return iife220_(_pat_let73_0)
                        return iife219_(pat_let_tv62_)
                    return iife218_(_pat_let72_0)
                nw377_[int(11)] = _dafny.MultiSet([d_2335_v63_, iife217_(d_2335_v63_)])
                nw377_[int(12)] = (_dafny.MultiSet(d_2336_v64_)) | (d_2333_v61_)
                nw377_[int(13)] = d_2333_v61_
                nw377_[int(14)] = (d_2333_v61_).set(d_2335_v63_, default__.abs(len(d_2340_v66_)))
                nw377_[int(15)] = (d_2333_v61_) - ((d_2341_v67_).cf109)
                nw377_[int(16)] = d_2333_v61_
                nw377_[int(17)] = d_2333_v61_
                nw377_[int(18)] = d_2333_v61_
                nw377_[int(19)] = (_dafny.MultiSet([d_2335_v63_])).set(d_2335_v63_, default__.abs((self).f31))
                nw377_[int(20)] = _dafny.MultiSet([d_2335_v63_, d_2335_v63_, d_2335_v63_, d_2335_v63_])
                nw377_[int(21)] = (d_2333_v61_).intersection(d_2333_v61_)
                nw377_[int(22)] = d_2333_v61_
                nw377_[int(23)] = _dafny.MultiSet([d_2335_v63_, D1_DC3((d_2334_v62_).set(default__.safeIndex((self).f31, len(d_2334_v62_)), (self).f26))])
                nw377_[int(24)] = d_2333_v61_
                nw377_[int(25)] = d_2333_v61_
                nw377_[int(26)] = d_2333_v61_
                d_2342_v68_ = nw377_
                index348_ = default__.safeIndex(641, (d_2342_v68_).length(0))
                (d_2342_v68_)[index348_] = _dafny.MultiSet(d_2336_v64_)
                d_2347_v69_: _dafny.Map
                d_2347_v69_ = _dafny.Map({len(_dafny.Map({d_2266_v12_.f40: self.f30})): (self).f26})
                index349_ = default__.safeIndex(641, (d_2342_v68_).length(0))
                rhs348_ = (d_2333_v61_).set(D1_DC3(d_2334_v62_), default__.abs((d_2266_v12_.f40 if (d_2311_v41_).fm2((self).f27, len(d_2347_v69_), globalState) else 120)))
                rhs349_ = (0) - (default__.safeModuloInt(d_2266_v12_.f40, d_2266_v12_.f40))
                rhs350_ = d_2266_v12_.f40
                rhs351_ = (d_2266_v12_.f40) + (d_2266_v12_.f40)
                rhs352_ = (default__.fm39(globalState) if (self).f27 else d_2266_v12_.f40)
                lhs321_ = d_2342_v68_
                lhs322_ = default__.safeIndex(641, (d_2342_v68_).length(0))
                lhs323_ = globalState
                lhs324_ = globalState
                lhs325_ = globalState
                lhs326_ = d_2266_v12_
                lhs321_[lhs322_] = rhs348_
                lhs323_.f12 = rhs349_
                lhs324_.f16 = rhs350_
                lhs325_.f16 = rhs351_
                lhs326_.f40 = rhs352_
                d_2348_v70_: _dafny.Array
                nw378_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_2348_v70_ = nw378_
                (globalState).f5 = (d_2348_v70_ if (self).f27 else d_2348_v70_)
                d_2334_v62_ = (d_2334_v62_ if (self).f27 else _dafny.SeqWithoutIsStrInference([self.f30]))
                d_2349_v71_: _dafny.MultiSet
                d_2349_v71_ = _dafny.MultiSet([(self).f26])
                d_2350_v72_: _dafny.Seq
                d_2350_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f27, True]), d_2349_v71_, _dafny.MultiSet([False])])
                d_2351_v73_: _dafny.Map
                d_2351_v73_ = _dafny.Map({(self).f26: d_2349_v71_})
                d_2352_v74_: D14
                d_2352_v74_ = D14_DC32(d_2349_v71_, (self).f27, (self).f26, self.f30, 407)
                d_2353_v75_: _dafny.Array
                nw379_ = _dafny.Array(None, 15)
                nw379_[int(0)] = d_2349_v71_
                nw379_[int(1)] = d_2349_v71_
                nw379_[int(2)] = d_2349_v71_
                nw379_[int(3)] = default__.fm21((self).f31, globalState)
                nw379_[int(4)] = (d_2350_v72_)[default__.safeIndex(d_2266_v12_.f40, len(d_2350_v72_))]
                nw379_[int(5)] = (d_2349_v71_) - (d_2349_v71_)
                nw379_[int(6)] = ((d_2351_v73_)[self.f30] if (self.f30) in (d_2351_v73_) else d_2349_v71_)
                nw379_[int(7)] = (d_2352_v74_).cf56
                nw379_[int(8)] = d_2349_v71_
                nw379_[int(9)] = d_2349_v71_
                nw379_[int(10)] = _dafny.MultiSet((d_2334_v62_ if (self).f27 else d_2334_v62_))
                nw379_[int(11)] = d_2349_v71_
                nw379_[int(12)] = (d_2350_v72_)[default__.safeIndex(d_2266_v12_.f40, len(d_2350_v72_))]
                nw379_[int(13)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).fm2((self).f27, 49, globalState)]))
                nw379_[int(14)] = d_2349_v71_
                d_2353_v75_ = nw379_
                d_2354_v76_: _dafny.Array
                def lambda118_(d_2355_v12_):
                    def lambda119_(d_2356_i6_):
                        return (d_2356_i6_) * ((0) - (d_2355_v12_.f40))

                    return lambda119_

                init67_ = lambda118_(d_2266_v12_)
                nw380_ = _dafny.Array(None, 19)
                for i0_67_ in range(nw380_.length(0)):
                    nw380_[i0_67_] = init67_(i0_67_)
                d_2354_v76_ = nw380_
                d_2357_v77_: _dafny.Seq
                d_2357_v77_ = _dafny.SeqWithoutIsStrInference([d_2354_v76_])
                index350_ = default__.safeIndex(434, (d_2353_v75_).length(0))
                (d_2353_v75_)[index350_] = default__.fm21(len(d_2357_v77_), globalState)
                d_2358_v78_: _dafny.Seq
                d_2358_v78_ = _dafny.SeqWithoutIsStrInference([(d_2266_v12_.f40) + (d_2266_v12_.f40)])
                d_2359_v79_: _dafny.Map
                d_2359_v79_ = _dafny.Map({d_2266_v12_.f40: d_2266_v12_.f40})
                index351_ = default__.safeIndex(434, (d_2353_v75_).length(0))
                rhs353_ = (d_2266_v12_.f40 if self.f30 else default__.safeModuloInt((self).f31, d_2266_v12_.f40))
                rhs354_ = (d_2358_v78_)[default__.safeIndex((self).fm1(globalState), len(d_2358_v78_))]
                rhs355_ = d_2349_v71_
                rhs356_ = (self).fm2((d_2266_v12_).fm2((self).f27, d_2266_v12_.f40, globalState), default__.safeDivisionInt(d_2266_v12_.f40, len(d_2359_v79_)), globalState)
                lhs327_ = globalState
                lhs328_ = d_2266_v12_
                lhs329_ = d_2353_v75_
                lhs330_ = default__.safeIndex(434, (d_2353_v75_).length(0))
                lhs331_ = globalState
                lhs327_.f16 = rhs353_
                lhs328_.f40 = rhs354_
                lhs329_[lhs330_] = rhs355_
                lhs331_.f22 = rhs356_
            d_2360_v80_: _dafny.Seq
            d_2360_v80_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f31, (self).f31])])
            d_2361_v81_: _dafny.Seq
            d_2361_v81_ = _dafny.SeqWithoutIsStrInference([d_2360_v80_])
            d_2362_v82_: T2
            nw381_ = C2()
            nw381_.ctor__(((d_2361_v81_)[default__.safeIndex(d_2266_v12_.f40, len(d_2361_v81_))] if self.f30 else d_2360_v80_), not (True) or (not((self).f26)), (self).f27)
            d_2362_v82_ = nw381_
            d_2362_v82_ = d_2362_v82_
        d_2363_v83_: _dafny.Seq
        d_2363_v83_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuuanc"))
        d_2364_v84_: _dafny.Seq
        d_2364_v84_ = _dafny.SeqWithoutIsStrInference([(884) + (len(d_2363_v83_)), d_2266_v12_.f40])
        d_2365_v85_: str
        d_2365_v85_ = _dafny.CodePoint('i')
        d_2366_v86_: _dafny.Map
        d_2366_v86_ = _dafny.Map({self.f30: d_2365_v85_})
        d_2364_v84_ = (d_2364_v84_) + ((d_2364_v84_) + (default__.fm43((self).f31, d_2266_v12_.f40, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjmphn"))), d_2366_v86_, globalState)))
        d_2367_v87_: _dafny.Array
        nw382_ = _dafny.Array(int(0), 14)
        d_2367_v87_ = nw382_
        index352_ = default__.safeIndex(537, (d_2367_v87_).length(0))
        (d_2367_v87_)[index352_] = (self).f31
        d_2368_v88_: _dafny.Array
        def lambda120_(d_2369_i7_):
            return (self).f27

        init68_ = lambda120_
        nw383_ = _dafny.Array(None, 12)
        for i0_68_ in range(nw383_.length(0)):
            nw383_[i0_68_] = init68_(i0_68_)
        d_2368_v88_ = nw383_
        d_2370_v89_: _dafny.Set
        d_2370_v89_ = _dafny.Set({d_2368_v88_})
        d_2371_v90_: _dafny.Seq
        d_2371_v90_ = _dafny.SeqWithoutIsStrInference([d_2370_v89_, d_2370_v89_, d_2370_v89_])
        d_2372_v91_: _dafny.Map
        d_2372_v91_ = _dafny.Map({d_2266_v12_.f40: d_2370_v89_})
        index353_ = default__.safeIndex(537, (d_2367_v87_).length(0))
        (d_2367_v87_)[index353_] = len(((d_2371_v90_)[default__.safeIndex((0) - (d_2266_v12_.f40), len(d_2371_v90_))]) | (((d_2372_v91_)[d_2266_v12_.f40] if (d_2266_v12_.f40) in (d_2372_v91_) else d_2370_v89_)))
        d_2373_v92_: _dafny.Seq
        d_2373_v92_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_2374_v93_: D16
        d_2374_v93_ = D16_DC38(718, len(d_2373_v92_), (self).f26)
        pat_let_tv63_ = d_2367_v87_
        pat_let_tv64_ = d_2367_v87_
        pat_let_tv65_ = d_2373_v92_
        pat_let_tv66_ = d_2267_v13_
        def lambda121_(source34_):
            if source34_.is_DC38:
                d_2375___mcc_h11_ = source34_.cf69
                d_2376___mcc_h12_ = source34_.cf70
                d_2377___mcc_h13_ = source34_.cf71
                d_2378_cf71_ = d_2377___mcc_h13_
                d_2379_cf70_ = d_2376___mcc_h12_
                d_2380_cf69_ = d_2375___mcc_h11_
                return (self).f26
            elif source34_.is_DC37:
                d_2381___mcc_h14_ = source34_.cf68
                d_2382_cf68_ = d_2381___mcc_h14_
                return not(((pat_let_tv64_)[default__.safeIndex(537, (pat_let_tv63_).length(0))]) < (len(pat_let_tv65_)))
            elif True:
                d_2383___mcc_h15_ = source34_.cf72
                d_2384_cf72_ = d_2383___mcc_h15_
                return not((pat_let_tv66_) != ((_dafny.Set({self.f30, True}))))

        r0 = lambda121_(d_2374_v93_)
        r1 = (self).f27
        return r0, r1

    def m5(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: int = int(0)
        r2: bool = False
        d_2385_v0_: _dafny.Seq
        d_2385_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flsqn"))
        d_2386_v1_: _dafny.Seq
        d_2386_v1_ = _dafny.SeqWithoutIsStrInference([len(d_2385_v0_), p1])
        d_2387_v2_: _dafny.Array
        nw384_ = _dafny.Array(None, 14)
        nw384_[int(0)] = p1
        nw384_[int(1)] = len(_dafny.SeqWithoutIsStrInference([self.f30]))
        nw384_[int(2)] = (self).f31
        nw384_[int(3)] = (self).f31
        nw384_[int(4)] = (self).f31
        nw384_[int(5)] = len(d_2386_v1_)
        nw384_[int(6)] = len(_dafny.Map({(self).f31: p1}))
        nw384_[int(7)] = (self).f31
        nw384_[int(8)] = p1
        nw384_[int(9)] = (self).f31
        nw384_[int(10)] = p0
        nw384_[int(11)] = p1
        nw384_[int(12)] = p1
        nw384_[int(13)] = p0
        d_2387_v2_ = nw384_
        d_2388_v3_: _dafny.Map
        d_2388_v3_ = _dafny.Map({(D3_DC10(d_2387_v2_)).cf21: p1})
        d_2388_v3_ = (d_2388_v3_).set(d_2387_v2_, (0) - (p0))
        r1 = default__.fm39(globalState)
        d_2389_v4_: _dafny.MultiSet
        d_2389_v4_ = _dafny.MultiSet([p2])
        d_2390_v5_: _dafny.Set
        d_2390_v5_ = _dafny.Set({len(d_2386_v1_), p1, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujlp"))), (0) - ((d_2389_v4_).cardinality)})
        d_2391_v6_: _dafny.Map
        d_2391_v6_ = _dafny.Map({p2: d_2390_v5_})
        d_2392_v7_: _dafny.Seq
        d_2392_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1, -51, len(_dafny.Map({(self).f31: (self).f27}))}), _dafny.Set({p0}), ((d_2391_v6_)[(self).f27] if ((self).f27) in (d_2391_v6_) else d_2390_v5_)])
        d_2393_v8_: _dafny.Map
        d_2393_v8_ = _dafny.Map({d_2392_v7_: (self).f26})
        d_2394_i0_: int
        d_2394_i0_ = 0
        with _dafny.label("14"):
            while not(((d_2393_v8_)[d_2392_v7_] if (d_2392_v7_) in (d_2393_v8_) else self.f30)):
                with _dafny.c_label("14"):
                    if (d_2394_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2394_i0_ = (d_2394_i0_) + (1)
                    d_2395_v9_: _dafny.Seq
                    d_2395_v9_ = _dafny.SeqWithoutIsStrInference([self.f30, True])
                    d_2396_v10_: _dafny.Map
                    d_2396_v10_ = _dafny.Map({(d_2395_v9_).set(default__.safeIndex(p1, len(d_2395_v9_)), True): p0})
                    d_2397_v11_: C5
                    nw385_ = C5()
                    nw385_.ctor__((self).f26, self.f30)
                    d_2397_v11_ = nw385_
                    d_2398_v12_: _dafny.Map
                    d_2398_v12_ = _dafny.Map({(self).f31: d_2397_v11_})
                    d_2399_v13_: _dafny.Set
                    d_2399_v13_ = _dafny.Set({self.f30})
                    d_2400_v14_: _dafny.Set
                    d_2400_v14_ = _dafny.Set({d_2399_v13_})
                    d_2401_v15_: _dafny.Map
                    d_2401_v15_ = _dafny.Map({len(d_2400_v14_): (self).f26})
                    d_2402_v16_: D16
                    d_2402_v16_ = D16_DC38(len(d_2398_v12_), len(d_2401_v15_), (self).f27)
                    d_2403_v17_: _dafny.Seq
                    d_2403_v17_ = _dafny.SeqWithoutIsStrInference([d_2402_v16_, d_2402_v16_])
                    r1 = (self).fm4(D0_DC0((self).f26), ((d_2396_v10_)[_dafny.SeqWithoutIsStrInference([(self).f27, self.f30])] if (_dafny.SeqWithoutIsStrInference([(self).f27, self.f30])) in (d_2396_v10_) else p1), default__.safeModuloInt(len(d_2395_v9_), len(d_2403_v17_)), globalState)
                    (globalState).f16 = 154
                    d_2404_v18_: _dafny.Map
                    d_2404_v18_ = _dafny.Map({((self).f26 if not((self).f27) else p2): not ((self).f26) or (self.f30)})
                    d_2404_v18_ = (d_2404_v18_).set(p2, self.f30)
                    d_2405_v19_: _dafny.MultiSet
                    d_2405_v19_ = _dafny.MultiSet([d_2386_v1_])
                    d_2406_v20_: C3
                    nw386_ = C3()
                    nw386_.ctor__(d_2405_v19_, True, (p2) in (_dafny.MultiSet([self.f30])))
                    d_2406_v20_ = nw386_
                    pass
            pass
        r1 = (self).f31
        (globalState).f15 = False
        if (self).fm2(p2, p1, globalState):
            d_2407_v21_: _dafny.Array
            nw387_ = _dafny.Array(_dafny.CodePoint('D'), 6)
            d_2407_v21_ = nw387_
            d_2408_v22_: _dafny.Array
            def lambda122_(d_2409_i1_):
                return (D20_DC47() if self.f30 else D20_DC47())

            init69_ = lambda122_
            nw388_ = _dafny.Array(None, 16)
            for i0_69_ in range(nw388_.length(0)):
                nw388_[i0_69_] = init69_(i0_69_)
            d_2408_v22_ = nw388_
            d_2410_v23_: D20
            d_2410_v23_ = D20_DC47()
            index354_ = default__.safeIndex(120, (d_2408_v22_).length(0))
            (d_2408_v22_)[index354_] = d_2410_v23_
            index355_ = default__.safeIndex(120, (d_2408_v22_).length(0))
            rhs357_ = (d_2385_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2411_i2_ in range(default__.abs(-395))]))
            rhs358_ = d_2407_v21_
            rhs359_ = d_2387_v2_
            rhs360_ = d_2410_v23_
            lhs332_ = globalState
            lhs333_ = d_2408_v22_
            lhs334_ = default__.safeIndex(120, (d_2408_v22_).length(0))
            lhs332_.f6 = rhs357_
            d_2407_v21_ = rhs358_
            d_2387_v2_ = rhs359_
            lhs333_[lhs334_] = rhs360_
            (globalState).f16 = (((d_2389_v4_) | (d_2389_v4_)) - (d_2389_v4_)).cardinality
            r1 = (0) - (p0)
            d_2412_v24_: _dafny.Map
            d_2412_v24_ = _dafny.Map({len(d_2390_v5_): (self).f26})
            d_2413_v25_: D0
            d_2413_v25_ = D0_DC0((self).f27)
            d_2414_v26_: _dafny.Seq
            d_2414_v26_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
            if (len((d_2412_v24_ if p2 else _dafny.Map({p1: (self).f26})))) > ((self).fm4(d_2413_v25_, len(d_2414_v26_), p0, globalState)):
                d_2415_v27_: _dafny.Map
                d_2415_v27_ = _dafny.Map({p0: p0})
                (globalState).f12 = len(((d_2415_v27_).set(p0, p0)) | (d_2415_v27_))
                d_2416_v28_: _dafny.Array
                nw389_ = _dafny.Array(_dafny.MultiSet({}), 15)
                d_2416_v28_ = nw389_
                index356_ = default__.safeIndex(654, (d_2416_v28_).length(0))
                (d_2416_v28_)[index356_] = (self).f29
                d_2417_v29_: str
                d_2417_v29_ = _dafny.CodePoint('p')
                d_2418_v30_: D6
                d_2418_v30_ = D6_DC16(d_2385_v0_)
                d_2419_v31_: _dafny.Array
                nw390_ = _dafny.Array(None, 16)
                nw390_[int(0)] = d_2385_v0_
                nw390_[int(1)] = d_2385_v0_
                nw390_[int(2)] = d_2385_v0_
                nw390_[int(3)] = d_2385_v0_
                nw390_[int(4)] = d_2385_v0_
                nw390_[int(5)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2420_i3_ in range(default__.abs(881))])) + (d_2385_v0_)
                nw390_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gixiulo"))
                nw390_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_2421_i4_ in range(default__.abs(521))])
                nw390_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxxceu"))
                nw390_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qkrwt")) if self.f30 else d_2385_v0_)
                nw390_[int(10)] = (d_2385_v0_).set(default__.safeIndex(len(d_2385_v0_), len(d_2385_v0_)), d_2417_v29_)
                nw390_[int(11)] = d_2385_v0_
                nw390_[int(12)] = d_2385_v0_
                nw390_[int(13)] = (d_2385_v0_) + (d_2385_v0_)
                nw390_[int(14)] = d_2385_v0_
                nw390_[int(15)] = (d_2418_v30_).cf29
                d_2419_v31_ = nw390_
                index357_ = default__.safeIndex(172, (d_2419_v31_).length(0))
                (d_2419_v31_)[index357_] = (d_2385_v0_) + (default__.fm22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aj"))), self.f30, globalState))
                d_2422_v32_: _dafny.Seq
                d_2422_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_2414_v26_), ((d_2389_v4_).set(True, default__.abs((d_2386_v1_)[default__.safeIndex(757, len(d_2386_v1_))]))).set((self).f27, default__.abs(p1))])
                d_2423_v33_: _dafny.Seq
                d_2423_v33_ = _dafny.SeqWithoutIsStrInference([default__.fm22(p0, (self).f26, globalState)])
                index358_ = default__.safeIndex(654, (d_2416_v28_).length(0))
                index359_ = default__.safeIndex(172, (d_2419_v31_).length(0))
                rhs361_ = (d_2422_v32_)[default__.safeIndex(((d_2389_v4_)[(self).f26] if ((self).f26) in (d_2389_v4_) else p1), len(d_2422_v32_))]
                rhs362_ = (((self).f29) - (default__.fm42(((self).f29).cardinality, globalState))) - ((self).f29)
                rhs363_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyblpod"))) + (((d_2423_v33_)[default__.safeIndex(p1, len(d_2423_v33_))]) + (d_2385_v0_))
                lhs335_ = d_2416_v28_
                lhs336_ = default__.safeIndex(654, (d_2416_v28_).length(0))
                lhs337_ = d_2419_v31_
                lhs338_ = default__.safeIndex(172, (d_2419_v31_).length(0))
                d_2389_v4_ = rhs361_
                lhs335_[lhs336_] = rhs362_
                lhs337_[lhs338_] = rhs363_
                (globalState).f16 = p1
                d_2424_v34_: C0
                nw391_ = C0()
                nw391_.ctor__(self.f30)
                d_2424_v34_ = nw391_
                d_2424_v34_ = d_2424_v34_
                d_2425_v35_: _dafny.Array
                def lambda123_(d_2426_i5_):
                    return (self).f27

                init70_ = lambda123_
                nw392_ = _dafny.Array(None, 5)
                for i0_70_ in range(nw392_.length(0)):
                    nw392_[i0_70_] = init70_(i0_70_)
                d_2425_v35_ = nw392_
                d_2427_v36_: D5
                d_2427_v36_ = D5_DC14(p1, p0, (d_2424_v34_).f32)
                d_2428_v37_: C1
                nw393_ = C1()
                nw393_.ctor__(p0, p0, (d_2427_v36_).cf27, (d_2415_v27_) == (d_2415_v27_))
                d_2428_v37_ = nw393_
                d_2429_v38_: _dafny.Map
                d_2429_v38_ = _dafny.Map({(d_2424_v34_).f32: -689})
                rhs364_ = d_2425_v35_
                rhs365_ = d_2428_v37_
                rhs366_ = _dafny.SeqWithoutIsStrInference([(d_2424_v34_).fm7(d_2429_v38_, d_2386_v1_, p2, globalState)])
                rhs367_ = not(self.f30)
                rhs368_ = not (self.f30) or ((False) or (p2))
                lhs339_ = globalState
                lhs340_ = globalState
                d_2425_v35_ = rhs364_
                d_2428_v37_ = rhs365_
                d_2414_v26_ = rhs366_
                lhs339_.f15 = rhs367_
                lhs340_.f15 = rhs368_
            elif True:
                d_2430_v39_: _dafny.Map
                d_2430_v39_ = _dafny.Map({(self).f31: p1})
                d_2431_v40_: _dafny.Map
                d_2431_v40_ = _dafny.Map({_dafny.Map({True: (self).f26}): d_2430_v39_})
                d_2431_v40_ = d_2431_v40_
                d_2432_v41_: str
                d_2432_v41_ = _dafny.CodePoint('t')
                d_2433_v42_: str
                d_2434_v43_: _dafny.Seq
                d_2435_v44_: bool
                d_2436_v45_: _dafny.Map
                out67_: str
                out68_: _dafny.Seq
                out69_: bool
                out70_: _dafny.Map
                out67_, out68_, out69_, out70_ = default__.m0(d_2432_v41_, globalState)
                d_2433_v42_ = out67_
                d_2434_v43_ = out68_
                d_2435_v44_ = out69_
                d_2436_v45_ = out70_
                d_2437_v46_: C0
                nw394_ = C0()
                nw394_.ctor__((p0) < ((self).f31))
                d_2437_v46_ = nw394_
                d_2438_v47_: C2
                nw395_ = C2()
                nw395_.ctor__(_dafny.SeqWithoutIsStrInference([d_2386_v1_ for d_2439_i6_ in range(default__.abs(437))]), (self).f26, (d_2437_v46_).f32)
                d_2438_v47_ = nw395_
                d_2440_v48_: C5
                nw396_ = C5()
                nw396_.ctor__((self).f26, (self).f27)
                d_2440_v48_ = nw396_
                d_2441_v49_: _dafny.Map
                d_2441_v49_ = _dafny.Map({(len(d_2385_v0_)) + (p0): d_2440_v48_})
                d_2442_v50_: _dafny.Map
                d_2442_v50_ = _dafny.Map({(d_2437_v46_).f32: (self).f27})
                d_2441_v49_ = (d_2441_v49_).set(len((d_2442_v50_) | (d_2442_v50_)), d_2440_v48_)
            d_2443_v51_: _dafny.Set
            d_2443_v51_ = _dafny.Set({(self).f26, (self).f27})
            (globalState).f16 = len(d_2443_v51_)
        elif True:
            d_2444_v52_: C5
            nw397_ = C5()
            nw397_.ctor__((126) <= (p0), (True if p2 else (self).f27))
            d_2444_v52_ = nw397_
            d_2445_v53_: _dafny.Array
            def lambda124_(d_2446_v0_):
                def lambda125_(d_2447_i7_):
                    return D6_DC16(d_2446_v0_)

                return lambda125_

            init71_ = lambda124_(d_2385_v0_)
            nw398_ = _dafny.Array(None, 13)
            for i0_71_ in range(nw398_.length(0)):
                nw398_[i0_71_] = init71_(i0_71_)
            d_2445_v53_ = nw398_
            d_2448_v54_: _dafny.Array
            def lambda126_(d_2449_i8_):
                return _dafny.CodePoint('r')

            init72_ = lambda126_
            nw399_ = _dafny.Array(None, 13)
            for i0_72_ in range(nw399_.length(0)):
                nw399_[i0_72_] = init72_(i0_72_)
            d_2448_v54_ = nw399_
            d_2450_v55_: D1
            d_2450_v55_ = D1_DC4(d_2448_v54_, p1, (d_2389_v4_).cardinality, d_2385_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_2451_i9_ in range(default__.abs(-896))]))
            index360_ = default__.safeIndex(67, (d_2445_v53_).length(0))
            (d_2445_v53_)[index360_] = D6_DC16((d_2450_v55_).cf11)
            d_2452_v56_: _dafny.Array
            def lambda127_(d_2453_v0_):
                def lambda128_(d_2454_i10_):
                    return d_2453_v0_

                return lambda128_

            init73_ = lambda127_(d_2385_v0_)
            nw400_ = _dafny.Array(None, 17)
            for i0_73_ in range(nw400_.length(0)):
                nw400_[i0_73_] = init73_(i0_73_)
            d_2452_v56_ = nw400_
            index361_ = default__.safeIndex(547, (d_2452_v56_).length(0))
            (d_2452_v56_)[index361_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2455_i11_ in range(default__.abs(704))])
            d_2456_v57_: D21
            d_2456_v57_ = D21_DC50(self.f30, (0) - ((self).f31), self.f30, (self).f31)
            index362_ = default__.safeIndex(67, (d_2445_v53_).length(0))
            index363_ = default__.safeIndex(547, (d_2452_v56_).length(0))
            rhs369_ = d_2444_v52_
            rhs370_ = (p1) == ((p1) + ((0) - ((0) - (-898))))
            rhs371_ = self.f30
            rhs372_ = default__.fm38((d_2456_v57_).cf89, (0) - ((self).fm1(globalState)), d_2385_v0_, globalState)
            rhs373_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('f') if False else _dafny.CodePoint('e')) for d_2457_i12_ in range(default__.abs(632))])
            lhs341_ = globalState
            lhs342_ = globalState
            lhs343_ = d_2445_v53_
            lhs344_ = default__.safeIndex(67, (d_2445_v53_).length(0))
            lhs345_ = d_2452_v56_
            lhs346_ = default__.safeIndex(547, (d_2452_v56_).length(0))
            d_2444_v52_ = rhs369_
            lhs341_.f2 = rhs370_
            lhs342_.f22 = rhs371_
            lhs343_[lhs344_] = rhs372_
            lhs345_[lhs346_] = rhs373_
            d_2458_v58_: _dafny.Set
            d_2458_v58_ = _dafny.Set({(self).f26})
            d_2459_v59_: _dafny.Seq
            d_2459_v59_ = _dafny.SeqWithoutIsStrInference([d_2458_v58_, d_2458_v58_])
            d_2460_v60_: _dafny.Map
            d_2460_v60_ = _dafny.Map({d_2459_v59_: d_2385_v0_})
            (globalState).f6 = ((d_2460_v60_)[d_2459_v59_] if (d_2459_v59_) in (d_2460_v60_) else d_2385_v0_)
            r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofrmiv"))) + ((d_2452_v56_)[default__.safeIndex(547, (d_2452_v56_).length(0))])
            (globalState).f16 = (self).f31
            d_2461_v61_: _dafny.Set
            d_2461_v61_ = _dafny.Set({default__.fm22(p0, self.f30, globalState), (d_2452_v56_)[default__.safeIndex(547, (d_2452_v56_).length(0))], ((d_2452_v56_)[default__.safeIndex(547, (d_2452_v56_).length(0))] if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")))})
            d_2462_v62_: _dafny.Map
            d_2462_v62_ = _dafny.Map({(d_2452_v56_)[default__.safeIndex(547, (d_2452_v56_).length(0))]: p0})
            def iife221_():
                coll73_ = _dafny.Set()
                compr_73_: _dafny.Seq
                for compr_73_ in (d_2462_v62_).keys.Elements:
                    d_2463_v63_: _dafny.Seq = compr_73_
                    if (d_2463_v63_) in (d_2462_v62_):
                        coll73_ = coll73_.union(_dafny.Set([d_2463_v63_]))
                return _dafny.Set(coll73_)
            rhs374_ = (_dafny.Set({(d_2452_v56_)[default__.safeIndex(547, (d_2452_v56_).length(0))]})) | (iife221_()
            )
            rhs375_ = (self).f26
            lhs347_ = globalState
            d_2461_v61_ = rhs374_
            lhs347_.f2 = rhs375_
        r0 = default__.fm22(p1, not((p2) == (self.f30)), globalState)
        r1 = (0) - ((self).fm4(D0_DC0((self).f27), (p0) + (p1), p1, globalState))
        r2 = p2
        return r0, r1, r2

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_2464_v0_: _dafny.Seq
        d_2464_v0_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f26])
        hi21_ = len((d_2464_v0_) + (d_2464_v0_))
        for d_2465_i0_ in range(p1, hi21_):
            d_2466_v1_: _dafny.Array
            def lambda129_(d_2467_v0_, d_2468_i0_, d_2469_p1_, d_2470_p3_):
                def lambda130_(d_2471_i1_):
                    return (_dafny.SeqWithoutIsStrInference([len(d_2467_v0_), d_2468_i0_])) == (_dafny.SeqWithoutIsStrInference([d_2469_p1_, (self).f31, len(_dafny.SeqWithoutIsStrInference([d_2470_p3_ for d_2472_i2_ in range(default__.abs(734))]))]))

                return lambda130_

            init74_ = lambda129_(d_2464_v0_, d_2465_i0_, p1, p3)
            nw401_ = _dafny.Array(None, 7)
            for i0_74_ in range(nw401_.length(0)):
                nw401_[i0_74_] = init74_(i0_74_)
            d_2466_v1_ = nw401_
            index364_ = default__.safeIndex(908, (d_2466_v1_).length(0))
            (d_2466_v1_)[index364_] = (self.f30 if not((self).f26) else (self).f26)
            d_2473_v2_: _dafny.Map
            d_2473_v2_ = _dafny.Map({p1: not((self).f26)})
            index365_ = default__.safeIndex(908, (d_2466_v1_).length(0))
            (d_2466_v1_)[index365_] = ((d_2473_v2_)[p3] if (p3) in (d_2473_v2_) else True)
            d_2474_v3_: D1
            d_2474_v3_ = D1_DC3(_dafny.SeqWithoutIsStrInference([not(self.f30)]))
            d_2475_v5_: _dafny.Seq
            def iife222_():
                coll74_ = _dafny.Set()
                compr_74_: int
                for compr_74_ in _dafny.IntegerRange(435, 179):
                    d_2476_v4_: int = compr_74_
                    if ((435) <= (d_2476_v4_)) and ((d_2476_v4_) < (179)):
                        coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_2476_v4_, len(p0))]))
                return _dafny.Set(coll74_)
            d_2475_v5_ = _dafny.SeqWithoutIsStrInference([iife222_()
            ])
            d_2477_v6_: _dafny.Set
            d_2477_v6_ = _dafny.Set({545, -539, (self).fm1(globalState), 747, p3})
            d_2478_v7_: _dafny.Set
            d_2478_v7_ = _dafny.Set({(self).f26, True, (self).fm2((d_2466_v1_)[default__.safeIndex(908, (d_2466_v1_).length(0))], d_2465_i0_, globalState)})
            d_2479_v8_: D6
            d_2479_v8_ = D6_DC17(len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2466_v1_)[default__.safeIndex(908, (d_2466_v1_).length(0))], (self).f27]): (self).f31})), (self).f31, p3, -624, p1)
            index366_ = default__.safeIndex(908, (d_2466_v1_).length(0))
            rhs376_ = (default__.fm55(globalState)) != ((d_2475_v5_).set(default__.safeIndex(p1, len(d_2475_v5_)), d_2477_v6_))
            rhs377_ = ((len(d_2478_v7_)) + (len(_dafny.SeqWithoutIsStrInference([(d_2466_v1_)[default__.safeIndex(908, (d_2466_v1_).length(0))]])))) < ((self).f31)
            rhs378_ = self.f30
            rhs379_ = (d_2479_v8_).cf34
            rhs380_ = d_2474_v3_
            lhs348_ = globalState
            lhs349_ = globalState
            lhs350_ = d_2466_v1_
            lhs351_ = default__.safeIndex(908, (d_2466_v1_).length(0))
            lhs352_ = globalState
            lhs348_.f24 = rhs376_
            lhs349_.f15 = rhs377_
            lhs350_[lhs351_] = rhs378_
            lhs352_.f12 = rhs379_
            d_2474_v3_ = rhs380_
            (globalState).f12 = default__.fm39(globalState)
            (globalState).f12 = d_2465_i0_
        d_2480_v9_: _dafny.Map
        d_2480_v9_ = _dafny.Map({p1: not((self).fm2(self.f30, (self).f31, globalState))})
        d_2481_v10_: _dafny.Map
        d_2481_v10_ = _dafny.Map({len(d_2480_v9_): not((self).f26)})
        d_2482_v11_: _dafny.Seq
        d_2482_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvcq"))
        d_2483_v12_: _dafny.Map
        d_2483_v12_ = _dafny.Map({(((d_2480_v9_)[-801] if (-801) in (d_2480_v9_) else (self).fm2(self.f30, p1, globalState))) or (self.f30): d_2482_v11_})
        (globalState).f6 = ((d_2483_v12_)[(self).f27] if ((self).f27) in (d_2483_v12_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qntndfgm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkyhfrdrq"))))
        d_2484_v13_: D1
        d_2484_v13_ = D1_DC3(d_2464_v0_)
        def lambda131_(source35_):
            if source35_.is_DC4:
                d_2485___mcc_h0_ = source35_.cf7
                d_2486___mcc_h1_ = source35_.cf8
                d_2487___mcc_h2_ = source35_.cf9
                d_2488___mcc_h3_ = source35_.cf10
                d_2489___mcc_h4_ = source35_.cf11
                d_2490_cf11_ = d_2489___mcc_h4_
                d_2491_cf10_ = d_2488___mcc_h3_
                d_2492_cf9_ = d_2487___mcc_h2_
                d_2493_cf8_ = d_2486___mcc_h1_
                d_2494_cf7_ = d_2485___mcc_h0_
                return ((self).f27) == (self.f30)
            elif source35_.is_DC5:
                d_2495___mcc_h5_ = source35_.cf12
                d_2496_cf12_ = d_2495___mcc_h5_
                return (self).f27
            elif True:
                d_2497___mcc_h6_ = source35_.cf6
                d_2498_cf6_ = d_2497___mcc_h6_
                return (self).f26

        if lambda131_(d_2484_v13_):
            if not(((p3) * (737)) > ((0) - (p1))):
                d_2499_v14_: _dafny.Array
                def lambda132_(d_2500_i3_):
                    return self.f30

                init75_ = lambda132_
                nw402_ = _dafny.Array(None, 25)
                for i0_75_ in range(nw402_.length(0)):
                    nw402_[i0_75_] = init75_(i0_75_)
                d_2499_v14_ = nw402_
                d_2501_v15_: _dafny.Array
                d_2501_v15_ = d_2499_v14_
                d_2501_v15_ = d_2501_v15_
                (globalState).f15 = not (((self).f26) and (self.f30)) or (((self).f27) and (self.f30))
                d_2502_v16_: _dafny.Map
                d_2502_v16_ = _dafny.Map({default__.safeDivisionInt(-67, p3): p2})
                d_2502_v16_ = (d_2502_v16_).set(p1, p2)
                d_2503_v17_: _dafny.MultiSet
                d_2503_v17_ = _dafny.MultiSet([default__.safeModuloInt(p3, (self).f31), p1])
                d_2503_v17_ = (self).f29
                rhs381_ = d_2482_v11_
                rhs382_ = d_2482_v11_
                rhs383_ = (self).f31
                rhs384_ = p2
                rhs385_ = p3
                lhs353_ = globalState
                lhs354_ = globalState
                lhs355_ = globalState
                lhs356_ = globalState
                lhs357_ = globalState
                lhs353_.f6 = rhs381_
                lhs354_.f6 = rhs382_
                lhs355_.f16 = rhs383_
                lhs356_.f19 = rhs384_
                lhs357_.f12 = rhs385_
            elif True:
                d_2504_v18_: _dafny.MultiSet
                d_2504_v18_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p1])])
                d_2505_v19_: C3
                nw403_ = C3()
                nw403_.ctor__(d_2504_v18_, (self).f26, self.f30)
                d_2505_v19_ = nw403_
                d_2506_v20_: _dafny.Map
                d_2506_v20_ = _dafny.Map({p3: d_2505_v19_})
                d_2507_v21_: _dafny.Set
                d_2507_v21_ = _dafny.Set({d_2506_v20_, d_2506_v20_, (_dafny.Map({p3: d_2505_v19_})) | (d_2506_v20_), d_2506_v20_})
                d_2508_v22_: _dafny.Array
                nw404_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_2508_v22_ = nw404_
                d_2509_v23_: _dafny.Array
                nw405_ = _dafny.Array(False, 7)
                d_2509_v23_ = nw405_
                index367_ = default__.safeIndex(780, (d_2508_v22_).length(0))
                (d_2508_v22_)[index367_] = d_2509_v23_
                d_2510_v24_: _dafny.Seq
                d_2510_v24_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_2511_v25_: D10
                d_2511_v25_ = D10_DC24(d_2482_v11_, d_2510_v24_, self.f30)
                index368_ = default__.safeIndex(780, (d_2508_v22_).length(0))
                rhs386_ = d_2507_v21_
                rhs387_ = (default__.safeDivisionInt((self).f31, (self).f31)) - ((p3) * (p1))
                rhs388_ = (p2) in ((d_2511_v25_).cf44)
                rhs389_ = (self).f26
                rhs390_ = d_2509_v23_
                lhs358_ = globalState
                lhs359_ = globalState
                lhs360_ = globalState
                lhs361_ = d_2508_v22_
                lhs362_ = default__.safeIndex(780, (d_2508_v22_).length(0))
                d_2507_v21_ = rhs386_
                lhs358_.f12 = rhs387_
                lhs359_.f15 = rhs388_
                lhs360_.f15 = rhs389_
                lhs361_[lhs362_] = rhs390_
                d_2512_v26_: _dafny.Map
                d_2512_v26_ = _dafny.Map({p2: (self).f27})
                d_2512_v26_ = (d_2512_v26_).set(p2, self.f30)
                index369_ = default__.safeIndex(780, (d_2508_v22_).length(0))
                rhs391_ = (self).f31
                rhs392_ = (self).f26
                rhs393_ = d_2509_v23_
                lhs363_ = globalState
                lhs364_ = globalState
                lhs365_ = d_2508_v22_
                lhs366_ = default__.safeIndex(780, (d_2508_v22_).length(0))
                lhs363_.f12 = rhs391_
                lhs364_.f2 = rhs392_
                lhs365_[lhs366_] = rhs393_
                (globalState).f15 = not(False)
                (globalState).f16 = p3
            d_2513_v27_: _dafny.Array
            nw406_ = _dafny.Array(False, 25)
            d_2513_v27_ = nw406_
            d_2513_v27_ = d_2513_v27_
            d_2514_v28_: _dafny.Map
            d_2514_v28_ = _dafny.Map({d_2482_v11_: _dafny.SeqWithoutIsStrInference([p2 for d_2515_i4_ in range(default__.abs(932))])})
            (globalState).f12 = len((d_2514_v28_) | ((d_2514_v28_).set(d_2482_v11_, d_2482_v11_)))
            index370_ = default__.safeIndex(926, (d_2513_v27_).length(0))
            (d_2513_v27_)[index370_] = False
            index371_ = default__.safeIndex(926, (d_2513_v27_).length(0))
            (d_2513_v27_)[index371_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioovj"))) != (d_2482_v11_)
            d_2516_v29_: _dafny.Seq
            d_2516_v29_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f31), (self).f31, p1])
            d_2517_v30_: _dafny.Map
            d_2517_v30_ = _dafny.Map({d_2516_v29_: p3})
            d_2518_v32_: _dafny.Seq
            d_2518_v32_ = _dafny.SeqWithoutIsStrInference([p1 for d_2519_i5_ in range(default__.abs(40))])
            d_2520_v33_: _dafny.Seq
            d_2520_v33_ = _dafny.SeqWithoutIsStrInference([d_2516_v29_, (d_2518_v32_)])
            def iife223_():
                coll75_ = _dafny.Map()
                compr_75_: _dafny.Seq
                for compr_75_ in (d_2520_v33_).Elements:
                    d_2521_v31_: _dafny.Seq = compr_75_
                    if (d_2521_v31_) in (d_2520_v33_):
                        coll75_[d_2521_v31_] = p3
                return _dafny.Map(coll75_)
            d_2517_v30_ = iife223_()
            
        elif True:
            (globalState).f24 = (self).f27
            d_2522_v34_: _dafny.MultiSet
            d_2522_v34_ = _dafny.MultiSet([self.f30])
            d_2523_v35_: _dafny.Seq
            d_2523_v35_ = _dafny.SeqWithoutIsStrInference([p1, 72])
            d_2524_v36_: _dafny.Map
            d_2524_v36_ = _dafny.Map({(d_2523_v35_)[default__.safeIndex(p3, len(d_2523_v35_))]: p0})
            d_2525_v37_: T0
            nw407_ = C7()
            nw407_.ctor__(True, d_2524_v36_)
            d_2525_v37_ = nw407_
            d_2526_v38_: _dafny.Set
            d_2526_v38_ = _dafny.Set({p1})
            rhs394_ = (d_2522_v34_).set((self).f26, default__.abs(len(d_2526_v38_)))
            rhs395_ = (d_2523_v35_) <= (d_2523_v35_)
            rhs396_ = d_2525_v37_
            lhs367_ = globalState
            d_2522_v34_ = rhs394_
            lhs367_.f22 = rhs395_
            d_2525_v37_ = rhs396_
            d_2527_v39_: _dafny.Seq
            d_2527_v39_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2528_v40_: _dafny.Map
            d_2528_v40_ = _dafny.Map({self.f30: ((0) - (p3)) - ((D21_DC52((self).f31, d_2527_v39_)).cf95)})
            d_2529_v41_: D23
            d_2529_v41_ = D23_DC57()
            d_2530_v42_: _dafny.Map
            d_2530_v42_ = _dafny.Map({D23_DC58(d_2529_v41_): (self).f27})
            d_2531_v43_: D23
            d_2531_v43_ = D23_DC58(d_2529_v41_)
            d_2532_v44_: D10
            d_2532_v44_ = D10_DC24(d_2482_v11_, d_2527_v39_, False)
            d_2528_v40_ = (d_2528_v40_).set(((self).f27 if (self).f26 else ((d_2530_v42_)[d_2531_v43_] if (d_2531_v43_) in (d_2530_v42_) else (self).f26)), (_dafny.MultiSet([len(d_2482_v11_), len(d_2482_v11_), len((d_2532_v44_).cf44), p1, len(d_2464_v0_)])).cardinality)
            (globalState).f2 = (d_2526_v38_).ispropersubset(_dafny.Set({(self).f31, -14, (self).f31, -503, (self).f31}))
            (globalState).f16 = (self).f31
        d_2533_v45_: C8
        nw408_ = C8()
        nw408_.ctor__(_dafny.MultiSet([p3]))
        d_2533_v45_ = nw408_
        d_2534_v46_: _dafny.Array
        def lambda133_(d_2535_i7_):
            return self.f30

        init76_ = lambda133_
        nw409_ = _dafny.Array(None, 13)
        for i0_76_ in range(nw409_.length(0)):
            nw409_[i0_76_] = init76_(i0_76_)
        d_2534_v46_ = nw409_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_2534_v46_).length(0)):
            d_2536_i6_: int = guard_loop_3_
            if (True) and (((0) <= (d_2536_i6_)) and ((d_2536_i6_) < ((d_2534_v46_).length(0)))):
                (d_2534_v46_)[(d_2536_i6_)] = (self.f30) and (((d_2481_v10_)[len(d_2482_v11_)] if (len(d_2482_v11_)) in (d_2481_v10_) else (self).f26))
        d_2537_v47_: _dafny.Seq
        d_2537_v47_ = _dafny.SeqWithoutIsStrInference([p3])
        d_2538_v48_: _dafny.Set
        d_2538_v48_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([922, (self).fm1(globalState)]), (d_2537_v47_) + (d_2537_v47_)})
        d_2538_v48_ = d_2538_v48_
        r0 = p3
        r1 = (self).fm1(globalState)
        r2 = not((p1) == (p3))
        return r0, r1, r2

    @property
    def f31(self):
        return self._f31
