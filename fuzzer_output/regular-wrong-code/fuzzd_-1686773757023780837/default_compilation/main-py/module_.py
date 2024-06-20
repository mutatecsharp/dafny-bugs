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
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_0_i1_ in range(default__.abs(452))]) for d_1_i0_ in range(default__.abs(-157))])).Elements:
                d_2_v0_: _dafny.Seq = compr_0_
                if (d_2_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_0_i1_ in range(default__.abs(452))]) for d_1_i0_ in range(default__.abs(-157))])):
                    coll0_ = coll0_.union(_dafny.Set([d_2_v0_]))
            return _dafny.Set(coll0_)
        return default__.safeDivisionInt(len(iife0_()
        ), 114)

    @staticmethod
    def fm1(globalState):
        return _dafny.CodePoint('m')

    @staticmethod
    def fm2(p0, p1, globalState):
        return False

    @staticmethod
    def fm6(p0, p1, globalState):
        return D2_DC3((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrfdnr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhmutw"))))

    @staticmethod
    def fm7(p0, globalState):
        return (_dafny.MultiSet([True])) | (_dafny.MultiSet([True, True]))

    @staticmethod
    def fm8(p0, p1, globalState):
        return (_dafny.MultiSet([False])) - (_dafny.MultiSet([False]))

    @staticmethod
    def fm9(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([D23_DC63(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fntkqwquv"))), D23_DC63(not(not(not(True))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbbsoqdy"))), D23_DC63(not(True), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_3_i0_ in range(default__.abs(339))])), D23_DC63(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfuqyuqw"))), D23_DC63(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftyox")))])).ispropersubset(_dafny.MultiSet([D23_DC63(not(True), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_4_i1_ in range(default__.abs(-705))]))])), True, (False if not(True) else True), False])

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])).Elements:
                d_5_v0_: int = compr_1_
                if (d_5_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])):
                    coll1_ = coll1_.union(_dafny.Set([(d_5_v0_) * (len(_dafny.SeqWithoutIsStrInference([False, True, False])))]))
            return _dafny.Set(coll1_)
        return D0_DC1(_dafny.CodePoint('e'), (len(iife1_()
) if not(False) else 45), (_dafny.Set({not(True), False}) if True else _dafny.Set({not(True)})), (False) or (not(True)), (733) not in (_dafny.Set({348, 546})))

    @staticmethod
    def fm11(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxfurkjc")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_6_i0_ in range(default__.abs(590))]))

    @staticmethod
    def fm14(globalState):
        return _dafny.MultiSet([False, True])

    @staticmethod
    def fm16(p0, globalState):
        return (_dafny.Set({True, True, True, not(False)})).intersection(_dafny.Set({not(False)}))

    @staticmethod
    def fm17(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(71) * (-987)])

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: not(not(True))}))]))).intersection(_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j')]))).cardinality]))) - (_dafny.MultiSet([993]))

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.Map({not(False): False})) | (_dafny.Map({not(not(not(True))): True}))

    @staticmethod
    def fm20(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.Map({66: False})})).keys.Elements:
                d_7_v0_: _dafny.Seq = compr_2_
                if (d_7_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.Map({66: False})})):
                    coll2_[d_7_v0_] = False
            return _dafny.Map(coll2_)
        source0_ = D23_DC62(iife2_()
)
        if source0_.is_DC63:
            d_8___mcc_h0_ = source0_.cf121
            d_9___mcc_h1_ = source0_.cf122
            d_10_cf122_ = d_9___mcc_h1_
            d_11_cf121_ = d_8___mcc_h0_
            return D2_DC3(d_10_cf122_)
        elif source0_.is_DC64:
            d_12___mcc_h2_ = source0_.cf123
            d_13_cf123_ = d_12___mcc_h2_
            return D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))
        elif source0_.is_DC62:
            d_14___mcc_h3_ = source0_.cf120
            d_15_cf120_ = d_14___mcc_h3_
            return D2_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_16_i0_ in range(default__.abs(200))]))
        elif True:
            d_17___mcc_h4_ = source0_.cf124
            d_18_cf124_ = d_17___mcc_h4_
            return D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcvjkoxqd")))

    @staticmethod
    def fm21(globalState):
        return (_dafny.MultiSet([False])) | (_dafny.MultiSet([False, False, True, (D3_DC8(True)).cf16]))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([True, True])

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: _dafny.Seq
                for compr_5_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krscvahhm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "il"))})).Elements:
                    d_19_v1_: _dafny.Seq = compr_5_
                    if (d_19_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krscvahhm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "il"))})):
                        coll5_[d_19_v1_] = 341
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: _dafny.Seq
                for compr_4_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krscvahhm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "il"))})).Elements:
                    d_19_v1_: _dafny.Seq = compr_4_
                    if (d_19_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krscvahhm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "il"))})):
                        coll4_[d_19_v1_] = 341
                return _dafny.Map(coll4_)
            compr_3_: _dafny.Seq
            for compr_3_ in (iife4_()
            ).keys.Elements:
                d_20_v0_: _dafny.Seq = compr_3_
                if (d_20_v0_) in (iife5_()
                ):
                    coll3_[d_20_v0_] = 594
            return _dafny.Map(coll3_)
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: _dafny.Seq
            for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_22_i2_ in range(default__.abs(-304))]) for d_23_i1_ in range(default__.abs(-421))])).Elements:
                d_24_v2_: _dafny.Seq = compr_6_
                if (d_24_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_22_i2_ in range(default__.abs(-304))]) for d_23_i1_ in range(default__.abs(-421))])):
                    coll6_[d_24_v2_] = (0) - ((D6_DC18(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))): 728})), -594, _dafny.SeqWithoutIsStrInference([False]), False, len(_dafny.Map({True: len(_dafny.Map({False: False}))})))).cf31)
            return _dafny.Map(coll6_)
        return ((iife3_()
        ) | ((D29_DC76(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkkpy")): (_dafny.MultiSet([(D21_DC58(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_21_i0_ in range(default__.abs(898))]))).cf108, not(False), not(not(True))])).cardinality}))).cf140)) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u")): -237})) | (iife6_()
        ))

    @staticmethod
    def fm26(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([53, 864, 365])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(not(False))])).cardinality, (_dafny.MultiSet([True, False])).cardinality]))) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlc")))]))

    @staticmethod
    def fm27(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))

    @staticmethod
    def fm28(globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Set
            for compr_7_ in (_dafny.Set({(D28_DC74(True, _dafny.Set({395}), (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yltdmqefg"))})))).cf137})).Elements:
                d_25_v0_: _dafny.Set = compr_7_
                if (d_25_v0_) in (_dafny.Set({(D28_DC74(True, _dafny.Set({395}), (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yltdmqefg"))})))).cf137})):
                    coll7_[d_25_v0_] = -865
            return _dafny.Map(coll7_)
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(393, -937):
                d_26_v1_: int = compr_8_
                if ((393) <= (d_26_v1_)) and ((d_26_v1_) < (-937)):
                    coll8_ = coll8_.union(_dafny.Set([(d_26_v1_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_27_i0_ in range(default__.abs(512))])))]))
            return _dafny.Set(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(230, 503):
                d_28_v2_: int = compr_9_
                if ((230) <= (d_28_v2_)) and ((d_28_v2_) < (503)):
                    coll9_ = coll9_.union(_dafny.Set([(d_28_v2_) - (258)]))
            return _dafny.Set(coll9_)
        return (((D30_DC78(iife7_()
)).cf144) | (_dafny.Map({iife8_()
        : 950}))) | ((_dafny.Map({_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}): 635})) | (_dafny.Map({iife9_()
        : -864})))

    @staticmethod
    def fm29(globalState):
        return _dafny.SeqWithoutIsStrInference([False])

    @staticmethod
    def fm30(p0, p1, globalState):
        return _dafny.MultiSet([(D0_DC1(_dafny.CodePoint('x'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvxisawj"))), _dafny.Set({True}), False, not(not(True)))).cf5])

    @staticmethod
    def fm31(p0, p1, globalState):
        if False:
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(587, 411):
                    d_29_v0_: int = compr_10_
                    if ((587) <= (d_29_v0_)) and ((d_29_v0_) < (411)):
                        coll10_[(d_29_v0_) * (len(_dafny.Map({False: True})))] = -415
                return _dafny.Map(coll10_)
            return iife10_()
            
        elif True:
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyckjo"))): (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])))})

    @staticmethod
    def fm32(p0, p1, globalState):
        return D2_DC4(604, (False if False else not(True)), (len(_dafny.Map({True: _dafny.CodePoint('d')}))) != (len(_dafny.SeqWithoutIsStrInference([False]))), not((not(True)) and (False)))

    @staticmethod
    def fm33(globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjrixvybf"))), -245])).intersection((_dafny.MultiSet([-529])).intersection(_dafny.MultiSet([132])))

    @staticmethod
    def fm34(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(D29_DC77(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdosto"))), 499, _dafny.CodePoint('m'))).cf143 for d_30_i0_ in range(default__.abs(618))])

    @staticmethod
    def fm35(p0, globalState):
        return D0_DC1(_dafny.CodePoint('c'), 94, (_dafny.Set({True}) if False else _dafny.Set({not(not(not((D6_DC18(346, len(_dafny.Set({not(False)})), _dafny.SeqWithoutIsStrInference([not(True)]), True, 830)).cf33))), not(not(True))})), True, ((0) - (-289)) == (747))

    @staticmethod
    def fm36(p0, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(-750, 406):
                d_31_v0_: int = compr_11_
                if ((-750) <= (d_31_v0_)) and ((d_31_v0_) < (406)):
                    coll11_[default__.safeModuloInt(d_31_v0_, len(_dafny.SeqWithoutIsStrInference([True, True])))] = False
            return _dafny.Map(coll11_)
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(28, 390):
                d_32_v1_: int = compr_12_
                if ((28) <= (d_32_v1_)) and ((d_32_v1_) < (390)):
                    coll12_[default__.safeDivisionInt(d_32_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vw"))))] = True
            return _dafny.Map(coll12_)
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(358, -313):
                d_33_v2_: int = compr_13_
                if ((358) <= (d_33_v2_)) and ((d_33_v2_) < (-313)):
                    coll13_[(d_33_v2_) + (-589)] = not(True)
            return _dafny.Map(coll13_)
        return (_dafny.Set({iife11_()
        })).intersection((_dafny.Set({iife12_()
        , _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sa"))): True})})) - (_dafny.Set({iife13_()
        })))

    @staticmethod
    def fm37(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([-213])) + (_dafny.SeqWithoutIsStrInference([18]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True, True])) for d_34_i0_ in range(default__.abs(536))])) + (_dafny.SeqWithoutIsStrInference([819, len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhdv"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utx"))), 234])))

    @staticmethod
    def fm38(p0, globalState):
        return D2_DC6(D2_DC5(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnd")))))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([_dafny.CodePoint('u')])

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in _dafny.IntegerRange(592, 521):
                d_35_v0_: int = compr_14_
                if ((592) <= (d_35_v0_)) and ((d_35_v0_) < (521)):
                    coll14_[(d_35_v0_) - (508)] = len(_dafny.SeqWithoutIsStrInference([382, len(_dafny.Set({False, False})), 220]))
            return _dafny.Map(coll14_)
        return D6_DC18(-464, len((_dafny.Map({len(_dafny.Map({False: 288})): 889})) | (iife14_()
)), (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(False), False])), False, len(_dafny.Set({True})))

    @staticmethod
    def fm42(p0, p1, globalState):
        if (D28_DC74(False, _dafny.Set({len(_dafny.SeqWithoutIsStrInference([408]))}), _dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbprfgut"))}))).cf136:
            return D2_DC4((_dafny.MultiSet([False, not(False), False, True])).cardinality, True, False, True)
        elif True:
            return D2_DC4(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))), False, not(True), True)

    @staticmethod
    def fm43(p0, p1, globalState):
        return ((_dafny.Map({not(True): D2_DC4(910, True, True, True)})) | (_dafny.Map({not(False): D2_DC4(len(_dafny.Set({-688})), True, False, True)}))) | ((_dafny.Map({False: D2_DC4(254, not(False), False, False)}) if False else _dafny.Map({False: D2_DC4((_dafny.MultiSet([not(True)])).cardinality, True, True, False)})))

    @staticmethod
    def fm44(p0, globalState):
        return D3_DC8((not(True) if True else True))

    @staticmethod
    def fm45(globalState):
        return ((_dafny.SeqWithoutIsStrInference([528])) + (_dafny.SeqWithoutIsStrInference([527 for d_36_i0_ in range(default__.abs(312))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_37_i1_ in range(default__.abs(718))]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_38_i2_ in range(default__.abs(-190))]))])))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyfk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jch")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_39_i0_ in range(default__.abs(745))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rtgapgxx"))))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality])) for d_40_i0_ in range(default__.abs(522))])) + ((D10_DC26(_dafny.SeqWithoutIsStrInference([(0) - (-814)]))).cf49)

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        return D0_DC1(_dafny.CodePoint('q'), default__.safeDivisionInt(-457, len(_dafny.Map({926: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqxqbisor")))]))}))), (_dafny.Set({False, not(False), True, False, True})) | (_dafny.Set({False, True})), not((_dafny.CodePoint('n')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvrugbft")))), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gucrheabq")))) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iud"))), len(_dafny.Set({True, not(False)}))])))

    @staticmethod
    def fm49(p0, globalState):
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))): _dafny.CodePoint('b')}))) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwf")))): _dafny.CodePoint('n')}))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return (D6_DC18(845, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([51])): True})), _dafny.SeqWithoutIsStrInference([not(False)]), True, len(_dafny.Set({False})))).cf32

    @staticmethod
    def fm51(globalState):
        return (_dafny.MultiSet([True, True, False, False])) - (_dafny.MultiSet([not(False), True]))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return _dafny.Set({not (False) or (not(False)), (_dafny.Set({True, not(True)})).issubset(_dafny.Set({True, True})), (len(_dafny.SeqWithoutIsStrInference([43]))) >= (-119)})

    @staticmethod
    def fm53(p0, globalState):
        return D10_DC28((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsqyq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wl"))), (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: not(True)})) for d_41_i0_ in range(default__.abs(246))])) if False else len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True})) for d_42_i1_ in range(default__.abs(-632))]))), False, False, (not(False)) and (True))

    @staticmethod
    def fm54(p0, globalState):
        return _dafny.MultiSet([_dafny.MultiSet([True]), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))).intersection(_dafny.MultiSet([not(not(False))]))])

    @staticmethod
    def fm55(globalState):
        if (286) <= (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([not(False)]))}))):
            return D3_DC8(False)
        elif True:
            return D3_DC8(True)

    @staticmethod
    def fm56(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([38, -732])) - ((_dafny.MultiSet([588])) - (_dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('k'): False})), -834, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecwfcmqt")))), 678, len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_43_i0_ in range(default__.abs(161))]): not(False)}))])))

    @staticmethod
    def fm57(p0, globalState):
        return D0_DC0((True) == (not(not(not(False)))))

    @staticmethod
    def fm58(p0, p1, p2, globalState):
        def iife15_():
            coll15_ = _dafny.Set()
            compr_15_: D5
            for compr_15_ in (_dafny.SeqWithoutIsStrInference([D5_DC15(D5_DC13(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({not(False): False})]))).cardinality: True}))) for d_44_i0_ in range(default__.abs(-636))])).Elements:
                d_45_v0_: D5 = compr_15_
                if (d_45_v0_) in (_dafny.SeqWithoutIsStrInference([D5_DC15(D5_DC13(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({not(False): False})]))).cardinality: True}))) for d_44_i0_ in range(default__.abs(-636))])):
                    coll15_ = coll15_.union(_dafny.Set([d_45_v0_]))
            return _dafny.Set(coll15_)
        return D16_DC44((iife15_()
 if (D17_DC48(True, False)).cf93 else _dafny.Set({D5_DC15(D5_DC13(_dafny.Map({653: True}))), D5_DC15(D5_DC14(False))})))

    @staticmethod
    def fm59(p0, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in ((_dafny.SeqWithoutIsStrInference([651]) if True else _dafny.SeqWithoutIsStrInference([192 for d_46_i0_ in range(default__.abs(-68))]))).Elements:
                d_47_v0_: int = compr_16_
                if (d_47_v0_) in ((_dafny.SeqWithoutIsStrInference([651]) if True else _dafny.SeqWithoutIsStrInference([192 for d_46_i0_ in range(default__.abs(-68))]))):
                    coll16_[(d_47_v0_) + (436)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir")))])) < (_dafny.SeqWithoutIsStrInference([756, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyrgjf")))}))]))
            return _dafny.Map(coll16_)
        return iife16_()
        

    @staticmethod
    def fm60(p0, globalState):
        return (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eu"))), (0) - ((_dafny.MultiSet([True])).cardinality), (0) - (len(_dafny.SeqWithoutIsStrInference([-595 for d_48_i0_ in range(default__.abs(33))])))})) | ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")))})) | (_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqtegsv")))})))

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        if (True) or (True):
            return D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrcu")))
        elif True:
            return D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjbvqmaj")))

    @staticmethod
    def fm62(p0, p1, p2, globalState):
        return (_dafny.Map({_dafny.MultiSet([472, 134]): _dafny.CodePoint('i')})) | (_dafny.Map({_dafny.MultiSet([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality]): _dafny.CodePoint('p')}))

    @staticmethod
    def fm63(p0, p1, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: _dafny.Map
            for compr_17_ in (_dafny.Map({_dafny.Map({671: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([966]))])).cardinality}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir"))})).keys.Elements:
                d_50_v0_: _dafny.Map = compr_17_
                if (d_50_v0_) in (_dafny.Map({_dafny.Map({671: (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([966]))])).cardinality}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ir"))})):
                    coll17_[d_50_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rq"))
            return _dafny.Map(coll17_)
        return ((_dafny.Map({(D32_DC82(_dafny.Map({788: -165}))).cf146: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsyfjkxk"))})) | (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True, True])): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_49_i0_ in range(default__.abs(821))]))}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qj"))}))) | ((iife17_()
        ) | (_dafny.Map({_dafny.Map({294: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baaaf")))}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vskrpe"))})))

    @staticmethod
    def fm64(p0, p1, p2, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_51_i0_ in range(default__.abs(62))]))).Elements:
                d_52_v0_: int = compr_18_
                if (d_52_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({True})) for d_51_i0_ in range(default__.abs(62))]))):
                    coll18_[(d_52_v0_) * (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_53_i1_ in range(default__.abs(-480))])): 885})))] = True
            return _dafny.Map(coll18_)
        return _dafny.SeqWithoutIsStrInference([iife18_()
        ])

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([-349, 542])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({-382: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_54_i0_ in range(default__.abs(299))]))}))])), (_dafny.SeqWithoutIsStrInference([-509])) + (_dafny.SeqWithoutIsStrInference([854, -356]))])

    @staticmethod
    def fm66(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([((D33_DC87(_dafny.Map({not(True): not(False)}))).cf155) | (_dafny.Map({False: False}))])

    @staticmethod
    def fm67(p0, p1, p2, p3, globalState):
        source1_ = D7_DC20((0) - (len(_dafny.Set({True, not(True), not(True)}))), False)
        if source1_.is_DC20:
            d_55___mcc_h0_ = source1_.cf36
            d_56___mcc_h1_ = source1_.cf37
            d_57_cf37_ = d_56___mcc_h1_
            d_58_cf36_ = d_55___mcc_h0_
            def iife19_():
                coll19_ = _dafny.Set()
                compr_19_: _dafny.Seq
                for compr_19_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_58_cf36_]): _dafny.CodePoint('d')})).keys.Elements:
                    d_59_v0_: _dafny.Seq = compr_19_
                    if (d_59_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([d_58_cf36_]): _dafny.CodePoint('d')})):
                        coll19_ = coll19_.union(_dafny.Set([d_59_v0_]))
                return _dafny.Set(coll19_)
            return D12_DC33(iife19_()
, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oeyndabq"))))
        elif source1_.is_DC21:
            d_60___mcc_h2_ = source1_.cf38
            d_61___mcc_h3_ = source1_.cf39
            d_62___mcc_h4_ = source1_.cf40
            d_63___mcc_h5_ = source1_.cf41
            d_64___mcc_h6_ = source1_.cf42
            d_65_cf42_ = d_64___mcc_h6_
            d_66_cf41_ = d_63___mcc_h5_
            d_67_cf40_ = d_62___mcc_h4_
            d_68_cf39_ = d_61___mcc_h3_
            d_69_cf38_ = d_60___mcc_h2_
            return D12_DC33(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_69_cf38_])}), d_69_cf38_)
        elif True:
            d_70___mcc_h7_ = source1_.cf35
            d_71_cf35_ = d_70___mcc_h7_
            return D12_DC33(_dafny.Set({_dafny.SeqWithoutIsStrInference([844]), _dafny.SeqWithoutIsStrInference([811])}), len(_dafny.SeqWithoutIsStrInference([not(not(True)), True, False])))

    @staticmethod
    def fm68(p0, p1, p2, p3, globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: D10
            for compr_20_ in (_dafny.MultiSet([D10_DC26(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_72_i1_ in range(default__.abs(-305))])) for d_73_i0_ in range(default__.abs(266))]))])).Elements:
                d_74_v0_: D10 = compr_20_
                if (d_74_v0_) in (_dafny.MultiSet([D10_DC26(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_72_i1_ in range(default__.abs(-305))])) for d_73_i0_ in range(default__.abs(266))]))])):
                    coll20_[d_74_v0_] = False
            return _dafny.Map(coll20_)
        return iife20_()
        

    @staticmethod
    def fm69(p0, globalState):
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: _dafny.Map
            for compr_21_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({624: -205}) for d_75_i0_ in range(default__.abs(679))])).Elements:
                d_76_v0_: _dafny.Map = compr_21_
                if (d_76_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({624: -205}) for d_75_i0_ in range(default__.abs(679))])):
                    coll21_ = coll21_.union(_dafny.Set([d_76_v0_]))
            return _dafny.Set(coll21_)
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(-458, 907):
                d_77_v1_: int = compr_22_
                if ((-458) <= (d_77_v1_)) and ((d_77_v1_) < (907)):
                    coll22_ = coll22_.union(_dafny.Set([(d_77_v1_) - (len(_dafny.Map({_dafny.CodePoint('f'): 118})))]))
            return _dafny.Set(coll22_)
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(197, 485):
                d_78_v2_: int = compr_23_
                if ((197) <= (d_78_v2_)) and ((d_78_v2_) < (485)):
                    coll23_[default__.safeModuloInt(d_78_v2_, 965)] = 199
            return _dafny.Map(coll23_)
        def iife24_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(323, 926):
                d_79_v3_: int = compr_24_
                if ((323) <= (d_79_v3_)) and ((d_79_v3_) < (926)):
                    coll24_[default__.safeDivisionInt(d_79_v3_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lksgqv"))))] = 917
            return _dafny.Map(coll24_)
        return (iife21_()
        ) - (_dafny.Set({_dafny.Map({len(_dafny.Map({False: False})): len(iife22_()
        )}), iife23_()
        , iife24_()
        , _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): -253}), _dafny.Map({len(_dafny.Map({960: True})): 735})}))

    @staticmethod
    def fm70(p0, p1, globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in (_dafny.MultiSet([len(_dafny.Set({not(not(not(True)))}))])).Elements:
                d_80_v0_: int = compr_25_
                if (d_80_v0_) in (_dafny.MultiSet([len(_dafny.Set({not(not(not(True)))}))])):
                    coll25_[(d_80_v0_) + (846)] = True
            return _dafny.Map(coll25_)
        return D5_DC13(iife25_()
)

    @staticmethod
    def fm71(globalState):
        source2_ = D13_DC38()
        if source2_.is_DC36:
            d_81___mcc_h0_ = source2_.cf68
            d_82_cf68_ = d_81___mcc_h0_
            if False:
                return D16_DC45(len(_dafny.Set({not(True), True})), True, d_82_cf68_)
            elif True:
                return D16_DC45(d_82_cf68_, True, 605)
        elif source2_.is_DC37:
            d_83___mcc_h1_ = source2_.cf69
            d_84___mcc_h2_ = source2_.cf70
            d_85___mcc_h3_ = source2_.cf71
            d_86_cf71_ = d_85___mcc_h3_
            d_87_cf70_ = d_84___mcc_h2_
            d_88_cf69_ = d_83___mcc_h1_
            return D16_DC45(d_88_cf69_, not(True), d_87_cf70_)
        elif source2_.is_DC38:
            return D16_DC45(len(_dafny.SeqWithoutIsStrInference([not(False), not(True), False, not(True), False])), False, -544)
        elif True:
            d_89___mcc_h4_ = source2_.cf67
            d_90_cf67_ = d_89___mcc_h4_
            return D16_DC45(516, False, 559)

    @staticmethod
    def fm72(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nv"))) + ((D8_DC23(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fk")))).cf44), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjhhjaemp"))])

    @staticmethod
    def fm73(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: not(True)})) | ((_dafny.Map({True: True})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm74(p0, globalState):
        return D4_DC11((not(True) if False else True), False, 85, (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True, False])), 778, 278, 520, 677])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eciqgjccx"))))))

    @staticmethod
    def fm75(p0, globalState):
        def iife26_():
            coll26_ = _dafny.Map()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(-100, 449):
                d_91_v0_: int = compr_26_
                if ((-100) <= (d_91_v0_)) and ((d_91_v0_) < (449)):
                    coll26_[(d_91_v0_) - (350)] = _dafny.CodePoint('o')
            return _dafny.Map(coll26_)
        return ((_dafny.Set({_dafny.Map({(0) - ((0) - (-784)): _dafny.CodePoint('b')}), _dafny.Map({181: _dafny.CodePoint('q')})})).intersection(_dafny.Set({iife26_()
        }))) - (_dafny.Set({_dafny.Map({(0) - (-461): _dafny.CodePoint('u')})}))

    @staticmethod
    def fm76(p0, globalState):
        return ((_dafny.Map({False: -95})) | (_dafny.Map({not(True): (_dafny.MultiSet([len(_dafny.Set({True})), (0) - ((_dafny.MultiSet([_dafny.Map({not(False): 141})])).cardinality)])).cardinality}))) | ((_dafny.Map({False: 721})) | (_dafny.Map({(D21_DC58(True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdylanfep")))).cf108: (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxd")))))})))

    @staticmethod
    def fm77(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 872}) for d_92_i0_ in range(default__.abs(101))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: -881}) for d_93_i1_ in range(default__.abs(-966))]))

    @staticmethod
    def fm78(globalState):
        return (_dafny.Map({993: _dafny.MultiSet([509])})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnnqqvf"))): _dafny.MultiSet([-284])}))

    @staticmethod
    def fm79(p0, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([_dafny.Set({False, False}) for d_94_i0_ in range(default__.abs(489))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({True, False}), _dafny.Set({not(False), True}), _dafny.Set({True, True})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Set({False})])))

    @staticmethod
    def fm80(p0, p1, globalState):
        if not(False):
            return D13_DC35(_dafny.Set({len(_dafny.Map({False: (0) - (-110)}))}))
        elif False:
            def iife27_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True])]))])).Elements:
                    d_95_v0_: int = compr_27_
                    if (d_95_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, True])]))])):
                        coll27_ = coll27_.union(_dafny.Set([default__.safeModuloInt(d_95_v0_, len(_dafny.SeqWithoutIsStrInference([True, False, False])))]))
                return _dafny.Set(coll27_)
            return D13_DC35(iife27_()
)
        elif True:
            def iife28_():
                coll28_ = _dafny.Set()
                compr_28_: int
                for compr_28_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eo"))) for d_96_i0_ in range(default__.abs(846))])).Elements:
                    d_97_v1_: int = compr_28_
                    if (d_97_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eo"))) for d_96_i0_ in range(default__.abs(846))])):
                        coll28_ = coll28_.union(_dafny.Set([default__.safeModuloInt(d_97_v1_, -79)]))
                return _dafny.Set(coll28_)
            return D13_DC35((D13_DC35(_dafny.Set({(0) - (len(iife28_()
))}))).cf67)

    @staticmethod
    def fm81(p0, globalState):
        def iife29_():
            coll29_ = _dafny.Set()
            compr_29_: _dafny.Map
            for compr_29_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))})])).Elements:
                d_98_v0_: _dafny.Map = compr_29_
                if (d_98_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))})])):
                    coll29_ = coll29_.union(_dafny.Set([d_98_v0_]))
            return _dafny.Set(coll29_)
        source3_ = D12_DC32(iife29_()
)
        if source3_.is_DC33:
            d_99___mcc_h0_ = source3_.cf62
            d_100___mcc_h1_ = source3_.cf63
            d_101_cf63_ = d_100___mcc_h1_
            d_102_cf62_ = d_99___mcc_h0_
            return D23_DC62(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): True}))
        elif source3_.is_DC34:
            d_103___mcc_h2_ = source3_.cf64
            d_104___mcc_h3_ = source3_.cf65
            d_105___mcc_h4_ = source3_.cf66
            d_106_cf66_ = d_105___mcc_h4_
            d_107_cf65_ = d_104___mcc_h3_
            d_108_cf64_ = d_103___mcc_h2_
            def iife30_():
                coll30_ = _dafny.Map()
                compr_30_: _dafny.Seq
                for compr_30_ in ((D34_DC90(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_106_cf66_, d_106_cf66_])}))).cf158).Elements:
                    d_109_v1_: _dafny.Seq = compr_30_
                    if (d_109_v1_) in ((D34_DC90(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_106_cf66_, d_106_cf66_])}))).cf158):
                        coll30_[d_109_v1_] = d_108_cf64_
                return _dafny.Map(coll30_)
            return D23_DC62(iife30_()
)
        elif True:
            d_110___mcc_h5_ = source3_.cf61
            d_111_cf61_ = d_110___mcc_h5_
            return D23_DC62(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): not(not(not(True)))}))

    @staticmethod
    def fm82(globalState):
        if False:
            return D22_DC60((0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (-101) for d_112_i0_ in range(default__.abs(-334))]))).cardinality), _dafny.Map({470: False}), -905, -953)
        elif True:
            return D22_DC60(len(_dafny.SeqWithoutIsStrInference([620 for d_113_i1_ in range(default__.abs(-391))])), _dafny.Map({394: False}), (_dafny.MultiSet([_dafny.CodePoint('u')])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o'), _dafny.CodePoint('n')])))

    @staticmethod
    def fm83(p0, p1, p2, globalState):
        return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_114_i0_ in range(default__.abs(28))])})) | ((_dafny.Map({(D18_DC50(not(True))).cf96: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jffwgpbnf"))}) if False else _dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcdflbvs"))})))

    @staticmethod
    def fm84(p0, p1, globalState):
        return D13_DC36(default__.safeModuloInt(len(_dafny.Map({_dafny.CodePoint('o'): False})), 641))

    @staticmethod
    def m13(p0, p1, p2, p3, globalState):
        d_115_v0_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 6)
        d_115_v0_ = nw0_
        d_116_v1_: D7
        d_116_v1_ = D7_DC21(default__.safeModuloInt(p3, p3), (p3) > (p3), not(True), d_115_v0_, (p2 if True else p2))
        d_117_v2_: _dafny.Array
        nw1_ = _dafny.Array(False, 1)
        d_117_v2_ = nw1_
        rhs0_ = len(_dafny.SeqWithoutIsStrInference([d_117_v2_]))
        rhs1_ = p2
        rhs2_ = d_116_v1_
        lhs0_ = globalState
        lhs1_ = globalState
        lhs0_.f18 = rhs0_
        lhs1_.f6 = rhs1_
        d_116_v1_ = rhs2_
        if not(p2):
            d_118_v3_: str
            d_118_v3_ = _dafny.CodePoint('m')
            d_119_v4_: _dafny.MultiSet
            d_119_v4_ = _dafny.MultiSet([d_118_v3_])
            d_120_v5_: _dafny.Seq
            d_120_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtfn"))
            d_121_v6_: _dafny.Set
            d_121_v6_ = _dafny.Set({p2})
            d_122_v7_: _dafny.Seq
            d_122_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p2}), d_121_v6_])
            d_123_v8_: _dafny.MultiSet
            d_123_v8_ = _dafny.MultiSet([p1])
            d_124_v9_: _dafny.Map
            d_124_v9_ = _dafny.Map({d_123_v8_: default__.fm52(p3, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rynnglrqe")), p2, p2, globalState)})
            d_125_v10_: D26
            d_125_v10_ = D26_DC70(d_118_v3_, len(d_120_v5_), p1, d_121_v6_)
            pat_let_tv0_ = d_121_v6_
            d_126_v11_: _dafny.Array
            nw2_ = _dafny.Array(None, 29)
            nw2_[int(0)] = (_dafny.Set({p1})).intersection(_dafny.Set({default__.fm2(d_119_v4_, len(d_120_v5_), globalState)}))
            nw2_[int(1)] = d_121_v6_
            nw2_[int(2)] = d_121_v6_
            nw2_[int(3)] = (d_122_v7_)[default__.safeIndex(p3, len(d_122_v7_))]
            nw2_[int(4)] = d_121_v6_
            nw2_[int(5)] = (d_121_v6_) - (d_121_v6_)
            nw2_[int(6)] = d_121_v6_
            nw2_[int(7)] = d_121_v6_
            nw2_[int(8)] = d_121_v6_
            nw2_[int(9)] = d_121_v6_
            nw2_[int(10)] = d_121_v6_
            nw2_[int(11)] = (d_121_v6_) - ((d_122_v7_)[default__.safeIndex(p3, len(d_122_v7_))])
            nw2_[int(12)] = d_121_v6_
            nw2_[int(13)] = (d_121_v6_).intersection(d_121_v6_)
            nw2_[int(14)] = d_121_v6_
            nw2_[int(15)] = (d_121_v6_) | (d_121_v6_)
            nw2_[int(16)] = d_121_v6_
            nw2_[int(17)] = _dafny.Set({p2})
            nw2_[int(18)] = d_121_v6_
            nw2_[int(19)] = ((d_124_v9_)[d_123_v8_] if (d_123_v8_) in (d_124_v9_) else d_121_v6_)
            nw2_[int(20)] = (d_121_v6_) - (d_121_v6_)
            nw2_[int(21)] = d_121_v6_
            nw2_[int(22)] = (d_121_v6_).intersection(d_121_v6_)
            nw2_[int(23)] = (d_121_v6_) | (d_121_v6_)
            nw2_[int(24)] = d_121_v6_
            nw2_[int(25)] = (d_121_v6_).intersection(d_121_v6_)
            def iife31_(_pat_let0_0):
                def iife32_(d_127_dt__update__tmp_h0_):
                    def iife33_(_pat_let1_0):
                        def iife34_(d_128_dt__update_hcf132_h0_):
                            return D26_DC70((d_127_dt__update__tmp_h0_).cf129, (d_127_dt__update__tmp_h0_).cf130, (d_127_dt__update__tmp_h0_).cf131, d_128_dt__update_hcf132_h0_)
                        return iife34_(_pat_let1_0)
                    return iife33_(pat_let_tv0_)
                return iife32_(_pat_let0_0)
            nw2_[int(26)] = (default__.fm16(d_120_v5_, globalState)).intersection((iife31_(d_125_v10_)).cf132)
            nw2_[int(27)] = d_121_v6_
            nw2_[int(28)] = d_121_v6_
            d_126_v11_ = nw2_
            index0_ = default__.safeIndex(73, (d_126_v11_).length(0))
            (d_126_v11_)[index0_] = d_121_v6_
            index1_ = default__.safeIndex(280, (d_117_v2_).length(0))
            (d_117_v2_)[index1_] = p2
            d_129_v12_: D8
            d_129_v12_ = D8_DC23(d_120_v5_)
            pat_let_tv1_ = d_118_v3_
            d_130_v13_: _dafny.Seq
            def iife35_(_pat_let2_0):
                def iife36_(d_131_dt__update__tmp_h1_):
                    def iife37_(_pat_let3_0):
                        def iife38_(d_133_dt__update_hcf44_h0_):
                            return D8_DC23(d_133_dt__update_hcf44_h0_)
                        return iife38_(_pat_let3_0)
                    return iife37_(_dafny.SeqWithoutIsStrInference([pat_let_tv1_ for d_132_i0_ in range(default__.abs(211))]))
                return iife36_(_pat_let2_0)
            d_130_v13_ = _dafny.SeqWithoutIsStrInference([D8_DC23(d_120_v5_), d_129_v12_, d_129_v12_, d_129_v12_, iife35_(d_129_v12_)])
            d_134_v14_: _dafny.Seq
            d_134_v14_ = _dafny.SeqWithoutIsStrInference([True])
            d_135_v15_: D2
            d_135_v15_ = D2_DC5(p1, len(d_134_v14_))
            d_136_v16_: D4
            d_136_v16_ = D4_DC11(p2, p2, p3, p3)
            d_137_v17_: _dafny.Seq
            d_137_v17_ = _dafny.SeqWithoutIsStrInference([d_130_v13_])
            index2_ = default__.safeIndex(73, (d_126_v11_).length(0))
            index3_ = default__.safeIndex(280, (d_117_v2_).length(0))
            rhs3_ = p1
            rhs4_ = (d_135_v15_).cf13
            rhs5_ = (d_121_v6_ if (p1 if (d_136_v16_).cf19 else p1) else (d_121_v6_) - (d_121_v6_))
            rhs6_ = p1
            rhs7_ = ((d_130_v13_) + ((d_137_v17_)[default__.safeIndex(p3, len(d_137_v17_))])).set(default__.safeIndex((0) - ((0) - (p3)), len((d_130_v13_) + ((d_137_v17_)[default__.safeIndex(p3, len(d_137_v17_))]))), d_129_v12_)
            lhs2_ = globalState
            lhs3_ = globalState
            lhs4_ = d_126_v11_
            lhs5_ = default__.safeIndex(73, (d_126_v11_).length(0))
            lhs6_ = d_117_v2_
            lhs7_ = default__.safeIndex(280, (d_117_v2_).length(0))
            lhs2_.f5 = rhs3_
            lhs3_.f18 = rhs4_
            lhs4_[lhs5_] = rhs5_
            lhs6_[lhs7_] = rhs6_
            d_130_v13_ = rhs7_
            d_138_v18_: D0
            d_138_v18_ = D0_DC0((d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))])
            d_139_v19_: C7
            nw3_ = C7()
            nw3_.ctor__((d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))], d_138_v18_)
            d_139_v19_ = nw3_
            d_140_v20_: D26
            d_140_v20_ = D26_DC69(d_139_v19_)
            source4_ = d_140_v20_
            if source4_.is_DC70:
                d_141___mcc_h0_ = source4_.cf129
                d_142___mcc_h1_ = source4_.cf130
                d_143___mcc_h2_ = source4_.cf131
                d_144___mcc_h3_ = source4_.cf132
                d_145_cf132_ = d_144___mcc_h3_
                d_146_cf131_ = d_143___mcc_h2_
                d_147_cf130_ = d_142___mcc_h1_
                d_148_cf129_ = d_141___mcc_h0_
                (globalState).f6 = (d_139_v19_).f26
                d_149_v21_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.Map({}), 16)
                d_149_v21_ = nw4_
                d_150_v22_: D0
                d_150_v22_ = D0_DC1(d_118_v3_, d_147_cf130_, _dafny.Set({p1}), p2, (d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))])
                pat_let_tv2_ = d_122_v7_
                pat_let_tv3_ = d_147_cf130_
                pat_let_tv4_ = d_122_v7_
                d_151_v23_: _dafny.Map
                def iife39_(_pat_let4_0):
                    def iife40_(d_152_dt__update__tmp_h2_):
                        def iife41_(_pat_let5_0):
                            def iife42_(d_153_dt__update_hcf3_h0_):
                                return D0_DC1((d_152_dt__update__tmp_h2_).cf1, (d_152_dt__update__tmp_h2_).cf2, d_153_dt__update_hcf3_h0_, (d_152_dt__update__tmp_h2_).cf4, (d_152_dt__update__tmp_h2_).cf5)
                            return iife42_(_pat_let5_0)
                        return iife41_((pat_let_tv2_)[default__.safeIndex(pat_let_tv3_, len(pat_let_tv4_))])
                    return iife40_(_pat_let4_0)
                d_151_v23_ = _dafny.Map({d_149_v21_: default__.fm0(831, iife39_(d_150_v22_), globalState)})
                d_151_v23_ = (d_151_v23_).set(d_149_v21_, p3)
                index4_ = default__.safeIndex(280, (d_117_v2_).length(0))
                (d_117_v2_)[index4_] = p1
                (globalState).f18 = p3
            elif source4_.is_DC69:
                d_154___mcc_h4_ = source4_.cf128
                d_155_cf128_ = d_154___mcc_h4_
                d_156_v24_: _dafny.Seq
                d_156_v24_ = _dafny.SeqWithoutIsStrInference([p3])
                (globalState).f5 = not(((_dafny.MultiSet([p3, len(d_156_v24_)])).set(p3, default__.abs(p3))).ispropersubset(_dafny.MultiSet([p3])))
                d_157_v25_: _dafny.Map
                d_157_v25_ = _dafny.Map({p3: p3})
                d_120_v5_ = default__.fm46(len(d_157_v25_), (d_139_v19_).f26, (d_139_v19_).f26, globalState)
                d_157_v25_ = (d_157_v25_).set((p3) - (p3), p3)
                (globalState).f16 = d_120_v5_
            elif True:
                d_158___mcc_h5_ = source4_.cf133
                d_159_cf133_ = d_158___mcc_h5_
                (globalState).f5 = not((d_139_v19_).f26)
                (globalState).f18 = p3
                d_160_v26_: T1
                nw5_ = C1()
                nw5_.ctor__(len(d_120_v5_), d_138_v18_)
                d_160_v26_ = nw5_
                index5_ = default__.safeIndex(918, (p0).length(0))
                (p0)[index5_] = d_160_v26_
                index6_ = default__.safeIndex(918, (p0).length(0))
                rhs8_ = ((d_120_v5_) + (_dafny.SeqWithoutIsStrInference([d_118_v3_ for d_161_i1_ in range(default__.abs(-852))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "akajnsajn"))) + (d_120_v5_))
                rhs9_ = p3
                rhs10_ = d_160_v26_
                lhs8_ = globalState
                lhs9_ = p0
                lhs10_ = default__.safeIndex(918, (p0).length(0))
                d_120_v5_ = rhs8_
                lhs8_.f17 = rhs9_
                lhs9_[lhs10_] = rhs10_
                d_162_v27_: T0
                nw6_ = C7()
                nw6_.ctor__(True, d_138_v18_)
                d_162_v27_ = nw6_
                d_163_v28_: _dafny.Map
                d_163_v28_ = _dafny.Map({p3: d_162_v27_})
                d_164_v29_: _dafny.Map
                d_164_v29_ = _dafny.Map({(d_139_v19_).f26: d_163_v28_})
                d_165_v30_: _dafny.Array
                nw7_ = _dafny.Array(None, 5)
                nw7_[int(0)] = ((d_164_v29_)[(d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))]] if ((d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))]) in (d_164_v29_) else d_163_v28_)
                nw7_[int(1)] = d_163_v28_
                nw7_[int(2)] = d_163_v28_
                nw7_[int(3)] = d_163_v28_
                nw7_[int(4)] = d_163_v28_
                d_165_v30_ = nw7_
                d_166_v31_: _dafny.Array
                nw8_ = _dafny.Array(None, 15)
                nw8_[int(0)] = d_165_v30_
                nw8_[int(1)] = d_165_v30_
                nw8_[int(2)] = d_165_v30_
                nw8_[int(3)] = (d_165_v30_ if p2 else d_165_v30_)
                nw8_[int(4)] = d_165_v30_
                nw8_[int(5)] = d_165_v30_
                nw8_[int(6)] = d_165_v30_
                nw8_[int(7)] = d_165_v30_
                nw8_[int(8)] = d_165_v30_
                nw8_[int(9)] = d_165_v30_
                nw8_[int(10)] = d_165_v30_
                nw8_[int(11)] = d_165_v30_
                nw8_[int(12)] = d_165_v30_
                nw8_[int(13)] = d_165_v30_
                nw8_[int(14)] = d_165_v30_
                d_166_v31_ = nw8_
                index7_ = default__.safeIndex(781, (d_166_v31_).length(0))
                (d_166_v31_)[index7_] = (d_165_v30_ if p1 else d_165_v30_)
                index8_ = default__.safeIndex(781, (d_166_v31_).length(0))
                rhs11_ = p3
                rhs12_ = (D28_DC73(d_165_v30_)).cf135
                lhs11_ = globalState
                lhs12_ = d_166_v31_
                lhs13_ = default__.safeIndex(781, (d_166_v31_).length(0))
                lhs11_.f17 = rhs11_
                lhs12_[lhs13_] = rhs12_
            d_167_v32_: C3
            nw9_ = C3()
            nw9_.ctor__(d_138_v18_)
            d_167_v32_ = nw9_
            nw10_ = C3()
            nw10_.ctor__(d_138_v18_)
            d_167_v32_ = nw10_
            d_168_v33_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_168_v33_ = nw11_
            d_169_v34_: _dafny.Array
            def lambda0_(d_170_v5_):
                def lambda1_(d_171_i2_):
                    return d_170_v5_

                return lambda1_

            init0_ = lambda0_(d_120_v5_)
            nw12_ = _dafny.Array(None, 21)
            for i0_0_ in range(nw12_.length(0)):
                nw12_[i0_0_] = init0_(i0_0_)
            d_169_v34_ = nw12_
            index9_ = default__.safeIndex(383, (d_168_v33_).length(0))
            (d_168_v33_)[index9_] = d_169_v34_
            d_172_v35_: _dafny.Seq
            d_172_v35_ = _dafny.SeqWithoutIsStrInference([d_120_v5_, d_120_v5_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxnae")), d_120_v5_])
            d_173_v36_: _dafny.Map
            d_173_v36_ = _dafny.Map({(d_117_v2_)[default__.safeIndex(280, (d_117_v2_).length(0))]: d_120_v5_})
            index10_ = default__.safeIndex(383, (d_168_v33_).length(0))
            nw13_ = _dafny.Array(None, 19)
            nw13_[int(0)] = _dafny.SeqWithoutIsStrInference([d_118_v3_ for d_174_i3_ in range(default__.abs(428))])
            nw13_[int(1)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vi"))).set(default__.safeIndex(len(d_120_v5_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vi")))), d_118_v3_)) + (d_120_v5_)
            nw13_[int(2)] = (d_120_v5_) + (d_120_v5_)
            nw13_[int(3)] = d_120_v5_
            nw13_[int(4)] = d_120_v5_
            nw13_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tugpqptu"))
            nw13_[int(6)] = d_120_v5_
            nw13_[int(7)] = (d_172_v35_)[default__.safeIndex(p3, len(d_172_v35_))]
            nw13_[int(8)] = ((d_120_v5_) + (d_120_v5_)).set(default__.safeIndex(p3, len((d_120_v5_) + (d_120_v5_))), default__.fm1(globalState))
            nw13_[int(9)] = d_120_v5_
            nw13_[int(10)] = (d_120_v5_) + (d_120_v5_)
            nw13_[int(11)] = d_120_v5_
            nw13_[int(12)] = d_120_v5_
            nw13_[int(13)] = ((d_173_v36_)[p1] if (p1) in (d_173_v36_) else _dafny.SeqWithoutIsStrInference([d_118_v3_ for d_175_i4_ in range(default__.abs(668))]))
            nw13_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmprer"))
            nw13_[int(15)] = d_120_v5_
            nw13_[int(16)] = (d_120_v5_).set(default__.safeIndex(p3, len(d_120_v5_)), (d_120_v5_)[default__.safeIndex(849, len(d_120_v5_))])
            nw13_[int(17)] = (d_120_v5_).set(default__.safeIndex(p3, len(d_120_v5_)), d_118_v3_)
            nw13_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_118_v3_ for d_176_i5_ in range(default__.abs(227))])) + (d_120_v5_)
            (d_168_v33_)[index10_] = nw13_
            d_177_v37_: _dafny.Map
            d_177_v37_ = _dafny.Map({(d_139_v19_).f26: (d_139_v19_).f26})
            d_178_v38_: _dafny.Seq
            d_178_v38_ = _dafny.SeqWithoutIsStrInference([(d_177_v37_) | (d_177_v37_), d_177_v37_, (_dafny.Map({p1: (d_139_v19_).f26})) | (d_177_v37_)])
            d_178_v38_ = (d_178_v38_) + (d_178_v38_)
        elif True:
            (globalState).f18 = p3
            d_179_v39_: _dafny.Map
            d_179_v39_ = _dafny.Map({p2: (default__.fm74((0) - (p3), globalState)).cf21})
            d_180_v40_: _dafny.Set
            d_180_v40_ = _dafny.Set({p3})
            d_179_v39_ = (d_179_v39_).set((d_180_v40_).ispropersubset(d_180_v40_), default__.safeDivisionInt(p3, p3))
            d_181_v41_: _dafny.Seq
            d_181_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufsfmdr"))
            (globalState).f17 = (0) - (default__.safeDivisionInt(len(d_181_v41_), p3))
            (globalState).f18 = (p3) - (p3)
            d_182_v42_: _dafny.Seq
            d_182_v42_ = _dafny.SeqWithoutIsStrInference([p2])
            d_183_v43_: D6
            d_183_v43_ = D6_DC18((513) - (p3), p3, d_182_v42_, p1, p3)
            source5_ = d_183_v43_
            if source5_.is_DC17:
                d_184___mcc_h6_ = source5_.cf28
                d_185___mcc_h7_ = source5_.cf29
                d_186_cf29_ = d_185___mcc_h7_
                d_187_cf28_ = d_184___mcc_h6_
                (globalState).f6 = p2
                d_188_v44_: str
                d_188_v44_ = _dafny.CodePoint('k')
                d_189_v45_: _dafny.Set
                d_189_v45_ = _dafny.Set({p1, p1})
                d_190_v46_: D0
                d_190_v46_ = D0_DC1(d_188_v44_, p3, d_189_v45_, p2, p1)
                d_179_v39_ = (d_179_v39_).set(p2, default__.fm0(p3, d_190_v46_, globalState))
                d_191_v47_: D0
                d_191_v47_ = D0_DC0(p1)
                d_192_v48_: C10
                nw14_ = C10()
                nw14_.ctor__(p3, d_191_v47_)
                d_192_v48_ = nw14_
                d_193_v49_: _dafny.Map
                d_193_v49_ = _dafny.Map({d_115_v0_: d_192_v48_})
                d_194_v50_: _dafny.Seq
                d_194_v50_ = _dafny.SeqWithoutIsStrInference([d_181_v41_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfsx")), _dafny.SeqWithoutIsStrInference([d_188_v44_ for d_195_i6_ in range(default__.abs(-261))]), d_181_v41_, d_181_v41_])
                d_196_v51_: _dafny.Seq
                d_196_v51_ = _dafny.SeqWithoutIsStrInference([p3, 978])
                rhs13_ = (d_192_v48_).fm5(p1, globalState)
                rhs14_ = (_dafny.Map({d_115_v0_: d_192_v48_}) if (len((d_194_v50_)[default__.safeIndex(len(d_189_v45_), len(d_194_v50_))])) in (d_196_v51_) else d_193_v49_)
                rhs15_ = (p1) or (p1)
                lhs14_ = globalState
                lhs15_ = globalState
                lhs14_.f18 = rhs13_
                d_193_v49_ = rhs14_
                lhs15_.f6 = rhs15_
                index11_ = default__.safeIndex(793, (d_117_v2_).length(0))
                (d_117_v2_)[index11_] = p1
                index12_ = default__.safeIndex(793, (d_117_v2_).length(0))
                (d_117_v2_)[index12_] = False
            elif source5_.is_DC18:
                d_197___mcc_h8_ = source5_.cf30
                d_198___mcc_h9_ = source5_.cf31
                d_199___mcc_h10_ = source5_.cf32
                d_200___mcc_h11_ = source5_.cf33
                d_201___mcc_h12_ = source5_.cf34
                d_202_cf34_ = d_201___mcc_h12_
                d_203_cf33_ = d_200___mcc_h11_
                d_204_cf32_ = d_199___mcc_h10_
                d_205_cf31_ = d_198___mcc_h9_
                d_206_cf30_ = d_197___mcc_h8_
                d_207_v52_: C8
                nw15_ = C8()
                nw15_.ctor__(d_203_cf33_, D0_DC0(p1))
                d_207_v52_ = nw15_
                d_208_v53_: _dafny.MultiSet
                d_208_v53_ = _dafny.MultiSet([d_203_cf33_])
                d_209_v54_: _dafny.Map
                d_209_v54_ = _dafny.Map({len(d_182_v42_): (d_208_v53_).cardinality})
                d_209_v54_ = (d_209_v54_).set((_dafny.MultiSet([-809, p3])).cardinality, d_206_cf30_)
                d_179_v39_ = _dafny.Map({d_203_cf33_: 173})
                d_210_v55_: C6
                nw16_ = C6()
                nw16_.ctor__(D0_DC0(p2))
                d_210_v55_ = nw16_
                index13_ = default__.safeIndex(853, (d_117_v2_).length(0))
                (d_117_v2_)[index13_] = p1
                index14_ = default__.safeIndex(853, (d_117_v2_).length(0))
                rhs16_ = d_117_v2_
                rhs17_ = p3
                rhs18_ = d_117_v2_
                rhs19_ = d_210_v55_
                rhs20_ = p1
                lhs16_ = globalState
                lhs17_ = globalState
                lhs18_ = globalState
                lhs19_ = d_117_v2_
                lhs20_ = default__.safeIndex(853, (d_117_v2_).length(0))
                lhs16_.f4 = rhs16_
                lhs17_.f17 = rhs17_
                lhs18_.f4 = rhs18_
                d_210_v55_ = rhs19_
                lhs19_[lhs20_] = rhs20_
            elif True:
                d_211___mcc_h13_ = source5_.cf27
                d_212_cf27_ = d_211___mcc_h13_
                (globalState).f17 = len((d_181_v41_) + (d_181_v41_))
                d_213_v56_: _dafny.Array
                nw17_ = _dafny.Array(_dafny.Seq({}), 5)
                d_213_v56_ = nw17_
                d_214_v57_: _dafny.Seq
                d_214_v57_ = _dafny.SeqWithoutIsStrInference([d_117_v2_])
                index15_ = default__.safeIndex(622, (d_213_v56_).length(0))
                (d_213_v56_)[index15_] = (d_214_v57_) + (d_214_v57_)
                index16_ = default__.safeIndex(622, (d_213_v56_).length(0))
                (d_213_v56_)[index16_] = d_214_v57_
                d_215_v58_: _dafny.Array
                nw18_ = _dafny.Array(None, 6)
                d_215_v58_ = nw18_
                d_216_v59_: D22
                d_216_v59_ = D22_DC59(d_215_v58_)
                pat_let_tv5_ = d_215_v58_
                def iife43_(_pat_let6_0):
                    def iife44_(d_217_dt__update__tmp_h3_):
                        def iife45_(_pat_let7_0):
                            def iife46_(d_218_dt__update_hcf110_h0_):
                                return D22_DC59(d_218_dt__update_hcf110_h0_)
                            return iife46_(_pat_let7_0)
                        return iife45_(pat_let_tv5_)
                    return iife44_(_pat_let6_0)
                d_216_v59_ = iife43_(d_216_v59_)
                d_219_v60_: _dafny.Seq
                d_219_v60_ = _dafny.SeqWithoutIsStrInference([p3])
                d_220_v61_: C1
                nw19_ = C1()
                nw19_.ctor__((d_219_v60_)[default__.safeIndex(p3, len(d_219_v60_))], D0_DC0(True))
                d_220_v61_ = nw19_
                rhs21_ = d_220_v61_
                rhs22_ = p2
                lhs21_ = globalState
                d_220_v61_ = rhs21_
                lhs21_.f5 = rhs22_
        hi0_ = (0) - (p3)
        for d_221_i7_ in range(p3, hi0_):
            d_222_v62_: _dafny.Array
            nw20_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_222_v62_ = nw20_
            d_223_v63_: str
            d_223_v63_ = _dafny.CodePoint('k')
            d_224_v64_: _dafny.Set
            d_224_v64_ = _dafny.Set({d_223_v63_})
            d_225_v65_: _dafny.Map
            d_225_v65_ = _dafny.Map({d_224_v64_: d_115_v0_})
            index17_ = default__.safeIndex(625, (d_222_v62_).length(0))
            (d_222_v62_)[index17_] = ((d_225_v65_)[d_224_v64_] if (d_224_v64_) in (d_225_v65_) else d_115_v0_)
            index18_ = default__.safeIndex(625, (d_222_v62_).length(0))
            (d_222_v62_)[index18_] = d_115_v0_
            if not ((not(p2)) and (False)) or ((d_221_i7_) == ((0) - (p3))):
                d_226_v66_: _dafny.Seq
                d_226_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xyrcmgb"))
                (globalState).f16 = d_226_v66_
                d_223_v63_ = (d_226_v66_)[default__.safeIndex(d_221_i7_, len(d_226_v66_))]
                (globalState).f6 = p2
                d_227_v67_: _dafny.Seq
                d_227_v67_ = _dafny.SeqWithoutIsStrInference([d_221_i7_])
                d_228_v68_: _dafny.Seq
                d_228_v68_ = _dafny.SeqWithoutIsStrInference([not(False), p2, False])
                d_229_v69_: C2
                nw21_ = C2()
                nw21_.ctor__((0) - ((0) - (default__.safeDivisionInt((d_227_v67_)[default__.safeIndex(-906, len(d_227_v67_))], len(d_228_v68_)))))
                d_229_v69_ = nw21_
                d_230_v70_: _dafny.Map
                d_230_v70_ = _dafny.Map({d_223_v63_: p1})
                d_230_v70_ = (d_230_v70_).set(d_223_v63_, p1)
            elif True:
                d_231_v71_: _dafny.Seq
                d_231_v71_ = _dafny.SeqWithoutIsStrInference([p1, p2])
                d_232_v72_: _dafny.Set
                d_232_v72_ = _dafny.Set({p3})
                d_233_v73_: _dafny.Set
                d_233_v73_ = _dafny.Set({d_231_v71_, default__.fm9(d_232_v72_, p1, globalState)})
                d_234_v74_: _dafny.MultiSet
                d_234_v74_ = _dafny.MultiSet([d_231_v71_])
                def iife47_():
                    coll31_ = _dafny.Set()
                    compr_31_: _dafny.Seq
                    for compr_31_ in (d_234_v74_).Elements:
                        d_235_v75_: _dafny.Seq = compr_31_
                        if (d_235_v75_) in (d_234_v74_):
                            coll31_ = coll31_.union(_dafny.Set([d_235_v75_]))
                    return _dafny.Set(coll31_)
                d_233_v73_ = ((d_233_v73_) - (d_233_v73_)) | (iife47_()
                )
                index19_ = default__.safeIndex(168, (d_115_v0_).length(0))
                (d_115_v0_)[index19_] = d_221_i7_
                index20_ = default__.safeIndex(168, (d_115_v0_).length(0))
                (d_115_v0_)[index20_] = p3
                (globalState).f6 = p1
                d_236_v76_: D0
                d_236_v76_ = D0_DC0(p2)
                d_237_v77_: C1
                nw22_ = C1()
                nw22_.ctor__(d_221_i7_, d_236_v76_)
                d_237_v77_ = nw22_
                d_238_v78_: _dafny.MultiSet
                d_238_v78_ = _dafny.MultiSet([d_223_v63_])
                nw23_ = C1()
                nw23_.ctor__(default__.fm0(len(_dafny.SeqWithoutIsStrInference([p3, (d_115_v0_)[default__.safeIndex(168, (d_115_v0_).length(0))], p3, (d_115_v0_)[default__.safeIndex(168, (d_115_v0_).length(0))]])), D0_DC1(d_223_v63_, p3, _dafny.Set({default__.fm2(d_238_v78_, 52, globalState)}), p2, True), globalState), d_236_v76_)
                d_237_v77_ = nw23_
                d_239_v79_: _dafny.Seq
                d_239_v79_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yguwc"))
                (globalState).f18 = default__.safeModuloInt(len(d_239_v79_), -597)
            d_240_v80_: D0
            d_240_v80_ = D0_DC0(p2)
            pat_let_tv6_ = p2
            d_241_v81_: C9
            nw24_ = C9()
            def iife48_(_pat_let8_0):
                def iife49_(d_242_dt__update__tmp_h4_):
                    def iife50_(_pat_let9_0):
                        def iife51_(d_243_dt__update_hcf0_h0_):
                            return D0_DC0(d_243_dt__update_hcf0_h0_)
                        return iife51_(_pat_let9_0)
                    return iife50_(pat_let_tv6_)
                return iife49_(_pat_let8_0)
            nw24_.ctor__(iife48_(d_240_v80_))
            d_241_v81_ = nw24_
            d_244_v82_: int
            out0_: int
            out0_ = (d_241_v81_).m3(globalState)
            d_244_v82_ = out0_
        d_245_v83_: _dafny.Set
        d_245_v83_ = _dafny.Set({p2, p1, True, p1, p1})
        d_246_v84_: _dafny.Map
        d_246_v84_ = _dafny.Map({p3: d_245_v83_})
        (globalState).f17 = (len(((d_246_v84_)[p3] if (p3) in (d_246_v84_) else _dafny.Set({p2})))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obsm"))))
        d_247_v85_: D3
        d_247_v85_ = D3_DC8(p1)
        pat_let_tv7_ = p1
        pat_let_tv8_ = p2
        pat_let_tv9_ = p2
        def lambda2_(source6_):
            if source6_.is_DC8:
                d_248___mcc_h14_ = source6_.cf16
                d_249_cf16_ = d_248___mcc_h14_
                return pat_let_tv7_
            elif source6_.is_DC7:
                d_250___mcc_h15_ = source6_.cf15
                d_251_cf15_ = d_250___mcc_h15_
                return pat_let_tv8_
            elif True:
                d_252___mcc_h16_ = source6_.cf17
                d_253_cf17_ = d_252___mcc_h16_
                return pat_let_tv9_

        if lambda2_(d_247_v85_):
            d_254_v86_: _dafny.Seq
            d_254_v86_ = _dafny.SeqWithoutIsStrInference([p3])
            d_255_v87_: D2
            d_255_v87_ = D2_DC4((d_254_v86_)[default__.safeIndex(p3, len(d_254_v86_))], p2, p1, p2)
            d_256_v88_: str
            d_256_v88_ = _dafny.CodePoint('c')
            d_257_v89_: _dafny.Seq
            d_257_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtddddrnp"))
            d_258_v90_: _dafny.Map
            d_258_v90_ = _dafny.Map({default__.fm2(_dafny.MultiSet([d_256_v88_]), len(d_257_v89_), globalState): p3})
            rhs23_ = (d_255_v87_).cf10
            rhs24_ = len(d_258_v90_)
            lhs22_ = globalState
            lhs23_ = globalState
            lhs22_.f6 = rhs23_
            lhs23_.f18 = rhs24_
            index21_ = default__.safeIndex(996, (d_115_v0_).length(0))
            (d_115_v0_)[index21_] = p3
            index22_ = default__.safeIndex(996, (d_115_v0_).length(0))
            (d_115_v0_)[index22_] = p3
            (globalState).f18 = p3
            d_259_v91_: _dafny.MultiSet
            d_259_v91_ = _dafny.MultiSet([p3])
            d_259_v91_ = (_dafny.MultiSet([p3, p3])) | (d_259_v91_)
            (globalState).f18 = (0) - (p3)
        elif True:
            d_260_v92_: _dafny.Seq
            d_260_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            (globalState).f16 = (d_260_v92_) + (d_260_v92_)
            d_261_v93_: str
            d_261_v93_ = _dafny.CodePoint('h')
            d_262_v94_: _dafny.MultiSet
            d_262_v94_ = _dafny.MultiSet([d_261_v93_, d_261_v93_, d_261_v93_, d_261_v93_, _dafny.CodePoint('t')])
            (globalState).f6 = default__.fm2((d_262_v94_).intersection(_dafny.MultiSet(d_260_v92_)), p3, globalState)
            d_263_v95_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Map({}), 4)
            d_263_v95_ = nw25_
            d_264_v96_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_264_v96_ = nw26_
            d_265_v97_: D20
            d_265_v97_ = D20_DC53(d_264_v96_)
            d_266_v98_: D20
            d_266_v98_ = D20_DC55(d_265_v97_)
            pat_let_tv10_ = d_264_v96_
            index23_ = default__.safeIndex(597, (d_263_v95_).length(0))
            def iife52_(_pat_let10_0):
                def iife53_(d_267_dt__update__tmp_h5_):
                    def iife54_(_pat_let11_0):
                        def iife55_(d_268_dt__update_hcf105_h0_):
                            return D20_DC55(d_268_dt__update_hcf105_h0_)
                        return iife55_(_pat_let11_0)
                    return iife54_(D20_DC53(pat_let_tv10_))
                return iife53_(_pat_let10_0)
            (d_263_v95_)[index23_] = _dafny.Map({iife52_(D20_DC55(d_265_v97_)): p3})
            d_269_v99_: D20
            d_269_v99_ = D20_DC55(d_265_v97_)
            d_270_v100_: _dafny.Map
            d_270_v100_ = _dafny.Map({d_269_v99_: 345})
            index24_ = default__.safeIndex(597, (d_263_v95_).length(0))
            (d_263_v95_)[index24_] = (d_270_v100_) | (d_270_v100_)
            d_271_v101_: _dafny.Seq
            d_271_v101_ = _dafny.SeqWithoutIsStrInference([p3])
            if default__.fm2((d_262_v94_) - (_dafny.MultiSet((d_260_v92_).set(default__.safeIndex(p3, len(d_260_v92_)), d_261_v93_))), len((_dafny.SeqWithoutIsStrInference([p3])) + (d_271_v101_)), globalState):
                index25_ = default__.safeIndex(676, (d_117_v2_).length(0))
                (d_117_v2_)[index25_] = (p2 if p1 else p2)
                index26_ = default__.safeIndex(676, (d_117_v2_).length(0))
                (d_117_v2_)[index26_] = ((d_260_v92_) < ((d_260_v92_).set(default__.safeIndex(p3, len(d_260_v92_)), d_261_v93_))) and ((_dafny.Set({p2, p2})).issubset(d_245_v83_))
                d_272_v102_: _dafny.Map
                d_272_v102_ = _dafny.Map({p3: p1})
                d_273_v103_: _dafny.Seq
                d_273_v103_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_274_v104_: D22
                d_274_v104_ = D22_DC60(p3, d_272_v102_, len(d_273_v103_), p3)
                d_275_v105_: _dafny.MultiSet
                d_275_v105_ = _dafny.MultiSet([p3, p3, 274, (d_274_v104_).cf111, p3])
                d_276_v106_: _dafny.Seq
                d_276_v106_ = _dafny.SeqWithoutIsStrInference([d_275_v105_])
                d_277_v107_: _dafny.Map
                d_277_v107_ = _dafny.Map({(d_117_v2_)[default__.safeIndex(676, (d_117_v2_).length(0))]: _dafny.Set({p3})})
                (globalState).f6 = (p2 if ((d_276_v106_)[default__.safeIndex(p3, len(d_276_v106_))]).issubset(d_275_v105_) else (len(d_277_v107_)) > ((0) - (p3)))
                d_278_v108_: D0
                d_278_v108_ = D0_DC0((d_117_v2_)[default__.safeIndex(676, (d_117_v2_).length(0))])
                d_278_v108_ = d_278_v108_
                d_279_v109_: _dafny.Array
                nw27_ = _dafny.Array(_dafny.Map({}), 16)
                d_279_v109_ = nw27_
                d_280_v110_: _dafny.Array
                nw28_ = _dafny.Array(None, 1)
                nw28_[int(0)] = d_273_v103_
                d_280_v110_ = nw28_
                d_281_v111_: D15
                d_281_v111_ = D15_DC43(d_280_v110_, p1, p3, p1, d_247_v85_)
                d_282_v112_: _dafny.Array
                def lambda3_(d_283_p1_):
                    def lambda4_(d_284_i8_):
                        return d_283_p1_

                    return lambda4_

                init1_ = lambda3_(p1)
                nw29_ = _dafny.Array(None, 8)
                for i0_1_ in range(nw29_.length(0)):
                    nw29_[i0_1_] = init1_(i0_1_)
                d_282_v112_ = nw29_
                d_285_v113_: C1
                nw30_ = C1()
                nw30_.ctor__((0) - (len(default__.fm16(d_260_v92_, globalState))), d_278_v108_)
                d_285_v113_ = nw30_
                d_286_v114_: _dafny.Map
                d_286_v114_ = _dafny.Map({d_282_v112_: d_285_v113_})
                index27_ = default__.safeIndex(74, (d_279_v109_).length(0))
                (d_279_v109_)[index27_] = (_dafny.Map({d_282_v112_: d_285_v113_}) if not((d_281_v111_).cf85) else d_286_v114_)
                index28_ = default__.safeIndex(74, (d_279_v109_).length(0))
                (d_279_v109_)[index28_] = _dafny.Map({d_282_v112_: d_285_v113_})
                d_287_v115_: _dafny.Map
                d_287_v115_ = _dafny.Map({not((d_117_v2_)[default__.safeIndex(676, (d_117_v2_).length(0))]): d_115_v0_})
                d_288_v116_: D13
                d_288_v116_ = D13_DC36(len(d_287_v115_))
                d_289_v117_: C4
                nw31_ = C4()
                nw31_.ctor__(d_278_v108_)
                d_289_v117_ = nw31_
                d_290_v118_: _dafny.Map
                d_290_v118_ = _dafny.Map({278: d_289_v117_})
                d_291_v119_: D7
                d_291_v119_ = D7_DC20(p3, p2)
                pat_let_tv11_ = d_271_v101_
                pat_let_tv12_ = p3
                pat_let_tv13_ = p3
                pat_let_tv14_ = p3
                d_292_v120_: _dafny.Array
                nw32_ = _dafny.Array(None, 23)
                nw32_[int(0)] = d_288_v116_
                nw32_[int(1)] = d_288_v116_
                nw32_[int(2)] = d_288_v116_
                nw32_[int(3)] = d_288_v116_
                nw32_[int(4)] = d_288_v116_
                nw32_[int(5)] = d_288_v116_
                nw32_[int(6)] = d_288_v116_
                nw32_[int(7)] = d_288_v116_
                def iife56_(_pat_let12_0):
                    def iife57_(d_293_dt__update__tmp_h6_):
                        def iife58_(_pat_let13_0):
                            def iife59_(d_294_dt__update_hcf68_h0_):
                                return D13_DC36(d_294_dt__update_hcf68_h0_)
                            return iife59_(_pat_let13_0)
                        return iife58_(174)
                    return iife57_(_pat_let12_0)
                nw32_[int(8)] = iife56_(d_288_v116_)
                nw32_[int(9)] = D13_DC36(len(d_290_v118_))
                def iife60_(_pat_let14_0):
                    def iife61_(d_295_dt__update__tmp_h7_):
                        def iife62_(_pat_let15_0):
                            def iife63_(d_296_dt__update_hcf68_h1_):
                                return D13_DC36(d_296_dt__update_hcf68_h1_)
                            return iife63_(_pat_let15_0)
                        return iife62_((_dafny.MultiSet(pat_let_tv11_)).cardinality)
                    return iife61_(_pat_let14_0)
                nw32_[int(10)] = (d_288_v116_ if p2 else iife60_(d_288_v116_))
                nw32_[int(11)] = D13_DC36(p3)
                nw32_[int(12)] = (d_288_v116_ if (d_291_v119_).cf37 else d_288_v116_)
                nw32_[int(13)] = D13_DC36(p3)
                nw32_[int(14)] = d_288_v116_
                nw32_[int(15)] = d_288_v116_
                def iife64_(_pat_let16_0):
                    def iife65_(d_297_dt__update__tmp_h8_):
                        def iife66_(_pat_let17_0):
                            def iife67_(d_298_dt__update_hcf68_h2_):
                                return D13_DC36(d_298_dt__update_hcf68_h2_)
                            return iife67_(_pat_let17_0)
                        return iife66_(pat_let_tv12_)
                    return iife65_(_pat_let16_0)
                nw32_[int(16)] = iife64_(d_288_v116_)
                def iife69_(_pat_let19_0):
                    def iife70_(d_299_dt__update__tmp_h9_):
                        def iife71_(_pat_let20_0):
                            def iife72_(d_300_dt__update_hcf68_h3_):
                                return D13_DC36(d_300_dt__update_hcf68_h3_)
                            return iife72_(_pat_let20_0)
                        return iife71_((0) - (pat_let_tv13_))
                    return iife70_(_pat_let19_0)
                def iife68_(_pat_let18_0):
                    def iife73_(d_301_dt__update__tmp_h10_):
                        def iife74_(_pat_let21_0):
                            def iife75_(d_302_dt__update_hcf68_h4_):
                                return D13_DC36(d_302_dt__update_hcf68_h4_)
                            return iife75_(_pat_let21_0)
                        return iife74_(pat_let_tv14_)
                    return iife73_(_pat_let18_0)
                nw32_[int(17)] = (iife68_(iife69_(D13_DC36(p3))) if p2 else d_288_v116_)
                nw32_[int(18)] = default__.fm84(p3, p1, globalState)
                nw32_[int(19)] = d_288_v116_
                nw32_[int(20)] = d_288_v116_
                nw32_[int(21)] = d_288_v116_
                nw32_[int(22)] = d_288_v116_
                d_292_v120_ = nw32_
                d_292_v120_ = d_292_v120_
            elif True:
                (globalState).f16 = d_260_v92_
                d_303_v121_: D0
                d_303_v121_ = D0_DC0(p2)
                d_304_v122_: C9
                nw33_ = C9()
                nw33_.ctor__(d_303_v121_)
                d_304_v122_ = nw33_
                d_305_v123_: _dafny.MultiSet
                d_305_v123_ = _dafny.MultiSet([p2, p1])
                d_306_v124_: _dafny.Seq
                d_306_v124_ = _dafny.SeqWithoutIsStrInference([d_305_v123_, d_305_v123_, d_305_v123_])
                d_307_v125_: _dafny.Map
                d_307_v125_ = _dafny.Map({p1: p2})
                d_308_v126_: _dafny.Array
                nw34_ = _dafny.Array(None, 19)
                nw34_[int(0)] = (d_305_v123_) - ((d_305_v123_).set(p2, default__.abs(457)))
                nw34_[int(1)] = d_305_v123_
                nw34_[int(2)] = d_305_v123_
                nw34_[int(3)] = d_305_v123_
                nw34_[int(4)] = d_305_v123_
                nw34_[int(5)] = (d_306_v124_)[default__.safeIndex(p3, len(d_306_v124_))]
                nw34_[int(6)] = d_305_v123_
                nw34_[int(7)] = ((d_306_v124_)[default__.safeIndex(p3, len(d_306_v124_))]) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([default__.fm2(d_262_v94_, p3, globalState)])))
                nw34_[int(8)] = d_305_v123_
                nw34_[int(9)] = default__.fm30(p1, d_260_v92_, globalState)
                nw34_[int(10)] = d_305_v123_
                nw34_[int(11)] = d_305_v123_
                nw34_[int(12)] = d_305_v123_
                nw34_[int(13)] = (_dafny.MultiSet([False, p2, not(((d_307_v125_)[p2] if (p2) in (d_307_v125_) else default__.fm2(_dafny.MultiSet([d_261_v93_, d_261_v93_, d_261_v93_, d_261_v93_]), p3, globalState)))])) - (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex(451, len(_dafny.SeqWithoutIsStrInference([p2]))), p2)))
                nw34_[int(14)] = d_305_v123_
                nw34_[int(15)] = d_305_v123_
                nw34_[int(16)] = d_305_v123_
                nw34_[int(17)] = (d_305_v123_).intersection(d_305_v123_)
                nw34_[int(18)] = d_305_v123_
                d_308_v126_ = nw34_
                d_309_v127_: _dafny.Seq
                d_309_v127_ = _dafny.SeqWithoutIsStrInference([p2])
                d_310_v128_: _dafny.Map
                d_310_v128_ = _dafny.Map({(d_260_v92_).set(default__.safeIndex(len(_dafny.Map({p2: p3})), len(d_260_v92_)), d_261_v93_): _dafny.MultiSet(d_309_v127_)})
                index29_ = default__.safeIndex(491, (d_308_v126_).length(0))
                (d_308_v126_)[index29_] = ((d_310_v128_)[d_260_v92_] if (d_260_v92_) in (d_310_v128_) else d_305_v123_)
                index30_ = default__.safeIndex(491, (d_308_v126_).length(0))
                (d_308_v126_)[index30_] = d_305_v123_
                d_311_v129_: _dafny.Map
                d_311_v129_ = _dafny.Map({p3: p1})
                d_311_v129_ = (d_311_v129_).set(p3, p2)
                d_312_v130_: C3
                nw35_ = C3()
                nw35_.ctor__(default__.fm57(394, globalState))
                d_312_v130_ = nw35_
            d_313_v131_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Array(None, 0), 26)
            d_313_v131_ = nw36_
            index31_ = default__.safeIndex(34, (d_313_v131_).length(0))
            (d_313_v131_)[index31_] = d_117_v2_
            index32_ = default__.safeIndex(34, (d_313_v131_).length(0))
            (d_313_v131_)[index32_] = d_117_v2_
        d_314_v132_: _dafny.Seq
        d_314_v132_ = _dafny.SeqWithoutIsStrInference([p2])
        d_315_v133_: C0
        nw37_ = C0()
        nw37_.ctor__(d_117_v2_, d_314_v132_)
        d_315_v133_ = nw37_
        d_316_v134_: _dafny.Seq
        d_316_v134_ = _dafny.SeqWithoutIsStrInference([p3, p3])
        d_317_v135_: _dafny.MultiSet
        d_317_v135_ = _dafny.MultiSet([p1])
        d_318_v136_: D6
        d_318_v136_ = D6_DC18((0) - (p3), p3, _dafny.SeqWithoutIsStrInference([p1, p1]), p2, (d_317_v135_).cardinality)
        d_319_v137_: str
        d_319_v137_ = _dafny.CodePoint('s')
        d_320_v138_: _dafny.Seq
        d_320_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnts"))
        d_321_v139_: T1
        nw38_ = C1()
        nw38_.ctor__(p3, default__.fm57(len(d_320_v138_), globalState))
        d_321_v139_ = nw38_
        d_322_v140_: _dafny.Map
        d_322_v140_ = _dafny.Map({d_321_v139_: not(p2)})
        d_323_v141_: _dafny.Map
        d_323_v141_ = _dafny.Map({d_319_v137_: ((d_322_v140_)[d_321_v139_] if (d_321_v139_) in (d_322_v140_) else not(p2))})
        d_324_v142_: D6
        d_324_v142_ = D6_DC18((d_316_v134_)[default__.safeIndex(p3, len(d_316_v134_))], p3, (d_318_v136_).cf32, not(p1), len(d_323_v141_))
        nw39_ = C0()
        nw39_.ctor__(d_117_v2_, (d_324_v142_).cf32)
        d_315_v133_ = nw39_

    @staticmethod
    def Main(noArgsParameter__):
        d_325_v0_: bool
        d_325_v0_ = False
        d_326_v1_: D0
        d_326_v1_ = D0_DC0(d_325_v0_)
        d_327_v2_: _dafny.Array
        nw40_ = _dafny.Array(None, 14)
        nw40_[int(0)] = d_325_v0_
        nw40_[int(1)] = d_325_v0_
        nw40_[int(2)] = False
        nw40_[int(3)] = d_325_v0_
        nw40_[int(4)] = d_325_v0_
        nw40_[int(5)] = d_325_v0_
        nw40_[int(6)] = (d_326_v1_).cf0
        nw40_[int(7)] = False
        nw40_[int(8)] = d_325_v0_
        nw40_[int(9)] = d_325_v0_
        nw40_[int(10)] = d_325_v0_
        nw40_[int(11)] = not(d_325_v0_)
        nw40_[int(12)] = d_325_v0_
        nw40_[int(13)] = False
        d_327_v2_ = nw40_
        d_328_v3_: _dafny.Seq
        d_328_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntc"))
        d_329_v4_: int
        d_329_v4_ = 40
        d_330_v5_: _dafny.Seq
        d_330_v5_ = _dafny.SeqWithoutIsStrInference([d_328_v3_, d_328_v3_])
        d_331_v6_: _dafny.Map
        d_331_v6_ = _dafny.Map({d_329_v4_: (d_330_v5_)[default__.safeIndex(d_329_v4_, len(d_330_v5_))]})
        d_332_v7_: _dafny.Array
        nw41_ = _dafny.Array(_dafny.CodePoint('D'), 28)
        d_332_v7_ = nw41_
        d_333_v8_: _dafny.Array
        nw42_ = _dafny.Array(int(0), 15)
        d_333_v8_ = nw42_
        d_334_v9_: _dafny.Array
        nw43_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_334_v9_ = nw43_
        d_335_globalState_: GlobalState
        nw44_ = GlobalState()
        nw44_.ctor__(False, True, True, False, d_327_v2_, False, False, True, -478, d_328_v3_, d_331_v6_, d_332_v7_, 559, 539, d_333_v8_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_336_i0_ in range(default__.abs(834))]), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_337_i1_ in range(default__.abs(263))])) + (d_328_v3_), 764, 282, d_334_v9_)
        d_335_globalState_ = nw44_
        (d_335_globalState_).f6 = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_338_i2_ in range(default__.abs(7))]))) == ((0) - (d_329_v4_))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_333_v8_).length(0)):
            d_339_i3_: int = guard_loop_0_
            if (True) and (((0) <= (d_339_i3_)) and ((d_339_i3_) < ((d_333_v8_).length(0)))):
                (d_333_v8_)[(d_339_i3_)] = (d_339_i3_) * (d_329_v4_)
        index33_ = default__.safeIndex(334, (d_333_v8_).length(0))
        (d_333_v8_)[index33_] = d_329_v4_
        d_340_v10_: str
        d_340_v10_ = _dafny.CodePoint('v')
        d_341_v11_: _dafny.Set
        d_341_v11_ = _dafny.Set({d_325_v0_})
        d_342_v12_: D0
        d_342_v12_ = D0_DC1(d_340_v10_, 15, d_341_v11_, d_325_v0_, d_325_v0_)
        d_343_v13_: _dafny.Seq
        d_343_v13_ = _dafny.SeqWithoutIsStrInference([146, d_329_v4_])
        pat_let_tv15_ = d_325_v0_
        pat_let_tv16_ = d_341_v11_
        pat_let_tv17_ = d_335_globalState_
        pat_let_tv18_ = d_325_v0_
        index34_ = default__.safeIndex(334, (d_333_v8_).length(0))
        def iife76_(_pat_let22_0):
            def iife77_(d_344_dt__update__tmp_h0_):
                def iife78_(_pat_let23_0):
                    def iife79_(d_345_dt__update_hcf4_h0_):
                        def iife80_(_pat_let24_0):
                            def iife81_(d_346_dt__update_hcf3_h0_):
                                def iife82_(_pat_let25_0):
                                    def iife83_(d_347_dt__update_hcf1_h0_):
                                        def iife84_(_pat_let26_0):
                                            def iife85_(d_348_dt__update_hcf5_h0_):
                                                return D0_DC1(d_347_dt__update_hcf1_h0_, (d_344_dt__update__tmp_h0_).cf2, d_346_dt__update_hcf3_h0_, d_345_dt__update_hcf4_h0_, d_348_dt__update_hcf5_h0_)
                                            return iife85_(_pat_let26_0)
                                        return iife84_(pat_let_tv18_)
                                    return iife83_(_pat_let25_0)
                                return iife82_(default__.fm1(pat_let_tv17_))
                            return iife81_(_pat_let24_0)
                        return iife80_(pat_let_tv16_)
                    return iife79_(_pat_let23_0)
                return iife78_(pat_let_tv15_)
            return iife77_(_pat_let22_0)
        rhs25_ = (0) - ((d_329_v4_) - (d_329_v4_))
        rhs26_ = (0) - (default__.fm0(d_329_v4_, iife76_(d_342_v12_), d_335_globalState_))
        rhs27_ = (d_343_v13_)[default__.safeIndex(default__.safeDivisionInt(d_329_v4_, d_329_v4_), len(d_343_v13_))]
        rhs28_ = 215
        rhs29_ = (len((d_328_v3_) + (d_328_v3_)) if (d_326_v1_).cf0 else len((d_343_v13_ if d_325_v0_ else d_343_v13_)))
        lhs24_ = d_335_globalState_
        lhs25_ = d_335_globalState_
        lhs26_ = d_333_v8_
        lhs27_ = default__.safeIndex(334, (d_333_v8_).length(0))
        lhs28_ = d_335_globalState_
        lhs24_.f18 = rhs25_
        lhs25_.f18 = rhs26_
        lhs26_[lhs27_] = rhs27_
        lhs28_.f17 = rhs28_
        d_329_v4_ = rhs29_
        (d_335_globalState_).f17 = default__.safeDivisionInt(d_329_v4_, (0) - (d_329_v4_))
        d_349_v14_: _dafny.Map
        d_349_v14_ = _dafny.Map({d_325_v0_: (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
        d_349_v14_ = (d_349_v14_).set(((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) >= (d_329_v4_), default__.safeModuloInt((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]))
        source7_ = D0_DC0(d_325_v0_)
        if source7_.is_DC1:
            d_350___mcc_h0_ = source7_.cf1
            d_351___mcc_h1_ = source7_.cf2
            d_352___mcc_h2_ = source7_.cf3
            d_353___mcc_h3_ = source7_.cf4
            d_354___mcc_h4_ = source7_.cf5
            d_355_cf5_ = d_354___mcc_h4_
            d_356_cf4_ = d_353___mcc_h3_
            d_357_cf3_ = d_352___mcc_h2_
            d_358_cf2_ = d_351___mcc_h1_
            d_359_cf1_ = d_350___mcc_h0_
            d_360_v15_: _dafny.MultiSet
            d_360_v15_ = _dafny.MultiSet([758, d_358_cf2_])
            d_361_v16_: _dafny.Seq
            d_361_v16_ = _dafny.SeqWithoutIsStrInference([d_325_v0_, d_325_v0_])
            index35_ = default__.safeIndex(105, (d_327_v2_).length(0))
            (d_327_v2_)[index35_] = (d_360_v15_).issubset((d_360_v15_).set((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], default__.abs(len(d_361_v16_))))
            d_362_v17_: _dafny.MultiSet
            d_362_v17_ = _dafny.MultiSet([d_325_v0_])
            index36_ = default__.safeIndex(105, (d_327_v2_).length(0))
            (d_327_v2_)[index36_] = ((_dafny.MultiSet(d_361_v16_)) - ((d_362_v17_))).ispropersubset(d_362_v17_)
            if (d_355_cf5_ if (193) > ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) else (d_342_v12_).cf4):
                d_355_cf5_ = (d_355_cf5_ if d_356_cf4_ else not((d_327_v2_)[default__.safeIndex(105, (d_327_v2_).length(0))]))
                index37_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index37_] = (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]
                d_363_v18_: _dafny.Map
                d_363_v18_ = _dafny.Map({d_333_v8_: d_359_cf1_})
                d_340_v10_ = ((d_363_v18_)[d_333_v8_] if (d_333_v8_) in (d_363_v18_) else d_359_cf1_)
                index38_ = default__.safeIndex(105, (d_327_v2_).length(0))
                (d_327_v2_)[index38_] = default__.fm2(_dafny.MultiSet([d_340_v10_]), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
                d_364_v19_: C1
                nw45_ = C1()
                nw45_.ctor__((len(d_328_v3_)) * (d_358_cf2_), D0_DC0(d_325_v0_))
                d_364_v19_ = nw45_
            elif True:
                d_365_v20_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Seq({}), 1)
                d_365_v20_ = nw46_
                index39_ = default__.safeIndex(946, (d_365_v20_).length(0))
                (d_365_v20_)[index39_] = d_361_v16_
                index40_ = default__.safeIndex(946, (d_365_v20_).length(0))
                (d_365_v20_)[index40_] = d_361_v16_
                d_366_v21_: _dafny.MultiSet
                d_366_v21_ = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([True])).set(default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(_dafny.SeqWithoutIsStrInference([True]))), (d_327_v2_)[default__.safeIndex(105, (d_327_v2_).length(0))]))
                d_367_v22_: _dafny.Array
                nw47_ = _dafny.Array(_dafny.Seq({}), 22)
                d_367_v22_ = nw47_
                d_368_v23_: D3
                d_368_v23_ = D3_DC7(d_367_v22_)
                d_369_v24_: C5
                nw48_ = C5()
                nw48_.ctor__(d_366_v21_, d_368_v23_, d_326_v1_)
                d_369_v24_ = nw48_
                d_356_cf4_ = not(d_355_cf5_)
                (d_335_globalState_).f17 = (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]
                d_370_v25_: _dafny.Array
                def lambda5_(d_371_cf5_):
                    def lambda6_(d_372_i4_):
                        return d_371_cf5_

                    return lambda6_

                init2_ = lambda5_(d_355_cf5_)
                nw49_ = _dafny.Array(None, 6)
                for i0_2_ in range(nw49_.length(0)):
                    nw49_[i0_2_] = init2_(i0_2_)
                d_370_v25_ = nw49_
                d_373_v26_: _dafny.MultiSet
                d_373_v26_ = _dafny.MultiSet([d_370_v25_])
                (d_335_globalState_).f6 = not((d_373_v26_).issubset(d_373_v26_))
            d_374_v27_: _dafny.Set
            d_374_v27_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference([d_359_cf1_ for d_375_i5_ in range(default__.abs(150))])).set(default__.safeIndex(d_358_cf2_, len(_dafny.SeqWithoutIsStrInference([d_359_cf1_ for d_376_i5_ in range(default__.abs(150))]))), (d_328_v3_)[default__.safeIndex(d_358_cf2_, len(d_328_v3_))]))})
            d_377_v28_: _dafny.Map
            d_377_v28_ = _dafny.Map({(d_328_v3_) + (d_328_v3_): d_374_v27_})
            d_377_v28_ = (d_377_v28_).set(d_328_v3_, (d_374_v27_ if default__.fm2(_dafny.MultiSet([(d_328_v3_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_328_v3_))]]), d_358_cf2_, d_335_globalState_) else d_374_v27_))
            index41_ = default__.safeIndex(105, (d_327_v2_).length(0))
            (d_327_v2_)[index41_] = d_356_cf4_
        elif True:
            d_378___mcc_h5_ = source7_.cf0
            d_379_cf0_ = d_378___mcc_h5_
            d_380_v29_: _dafny.MultiSet
            d_380_v29_ = _dafny.MultiSet([d_379_cf0_])
            index42_ = default__.safeIndex(50, (d_327_v2_).length(0))
            (d_327_v2_)[index42_] = ((d_380_v29_).cardinality) > (d_329_v4_)
            d_381_v30_: _dafny.Map
            d_381_v30_ = _dafny.Map({d_327_v2_: len((default__.fm81(d_325_v0_, d_335_globalState_)).cf120)})
            d_382_v31_: _dafny.Array
            def lambda7_(d_383_cf0_, d_384_v0_):
                def lambda8_(d_385_i6_):
                    return _dafny.SeqWithoutIsStrInference([d_383_cf0_, d_384_v0_])

                return lambda8_

            init3_ = lambda7_(d_379_cf0_, d_325_v0_)
            nw50_ = _dafny.Array(None, 26)
            for i0_3_ in range(nw50_.length(0)):
                nw50_[i0_3_] = init3_(i0_3_)
            d_382_v31_ = nw50_
            d_386_v32_: D2
            d_386_v32_ = D2_DC5(d_325_v0_, 356)
            d_387_v33_: D13
            d_387_v33_ = D13_DC37(len(d_341_v11_), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])
            d_388_v34_: D3
            d_388_v34_ = D3_DC8(True)
            d_389_v35_: D15
            d_389_v35_ = D15_DC43(d_382_v31_, (d_386_v32_).cf12, (d_387_v33_).cf70, d_325_v0_, d_388_v34_)
            index43_ = default__.safeIndex(50, (d_327_v2_).length(0))
            rhs30_ = not(not((d_389_v35_).cf83))
            rhs31_ = d_379_cf0_
            rhs32_ = (d_381_v30_) | ((d_381_v30_).set(d_327_v2_, (0) - (d_329_v4_)))
            lhs29_ = d_327_v2_
            lhs30_ = default__.safeIndex(50, (d_327_v2_).length(0))
            lhs31_ = d_335_globalState_
            lhs29_[lhs30_] = rhs30_
            lhs31_.f6 = rhs31_
            d_381_v30_ = rhs32_
            d_390_v36_: C3
            nw51_ = C3()
            nw51_.ctor__(d_326_v1_)
            d_390_v36_ = nw51_
            d_391_v37_: _dafny.MultiSet
            d_391_v37_ = _dafny.MultiSet([_dafny.Map({d_329_v4_: (d_380_v29_).cardinality})])
            d_391_v37_ = d_391_v37_
            d_343_v13_ = d_343_v13_
        if d_325_v0_:
            d_392_v38_: _dafny.Seq
            d_392_v38_ = _dafny.SeqWithoutIsStrInference([d_325_v0_, d_325_v0_, d_325_v0_])
            d_393_v39_: C0
            nw52_ = C0()
            nw52_.ctor__(d_327_v2_, (d_392_v38_).set(default__.safeIndex(667, len(d_392_v38_)), d_325_v0_))
            d_393_v39_ = nw52_
            d_394_v40_: _dafny.Map
            d_394_v40_ = _dafny.Map({d_393_v39_: d_325_v0_})
            d_395_v41_: _dafny.MultiSet
            d_395_v41_ = _dafny.MultiSet([(len(d_394_v40_)) * (d_329_v4_), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_329_v4_])
            d_395_v41_ = ((d_395_v41_) - ((d_395_v41_).set(d_329_v4_, default__.abs((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])))).intersection((d_395_v41_ if d_325_v0_ else _dafny.MultiSet([d_329_v4_])))
            d_333_v8_ = d_333_v8_
            d_396_v42_: C1
            nw53_ = C1()
            nw53_.ctor__(d_329_v4_, D0_DC0(d_325_v0_))
            d_396_v42_ = nw53_
            d_397_v43_: _dafny.Map
            d_397_v43_ = _dafny.Map({d_325_v0_: _dafny.SeqWithoutIsStrInference([d_396_v42_])})
            d_398_v44_: _dafny.Seq
            d_398_v44_ = _dafny.SeqWithoutIsStrInference([d_396_v42_])
            d_399_v45_: _dafny.Seq
            d_399_v45_ = _dafny.SeqWithoutIsStrInference([d_396_v42_, (d_398_v44_)[default__.safeIndex((default__.fm30(d_325_v0_, d_328_v3_, d_335_globalState_)).cardinality, len(d_398_v44_))]])
            d_397_v43_ = (((d_397_v43_).set(d_325_v0_, _dafny.SeqWithoutIsStrInference([d_396_v42_]))).set(False, d_399_v45_)) | (d_397_v43_)
            d_400_v46_: _dafny.Array
            nw54_ = _dafny.Array(None, 11)
            nw54_[int(0)] = d_396_v42_
            nw54_[int(1)] = d_396_v42_
            nw54_[int(2)] = d_396_v42_
            nw54_[int(3)] = d_396_v42_
            nw54_[int(4)] = d_396_v42_
            nw54_[int(5)] = d_396_v42_
            nw54_[int(6)] = d_396_v42_
            nw54_[int(7)] = d_396_v42_
            nw54_[int(8)] = d_396_v42_
            nw54_[int(9)] = d_396_v42_
            nw54_[int(10)] = d_396_v42_
            d_400_v46_ = nw54_
            d_401_v47_: _dafny.Seq
            d_401_v47_ = _dafny.SeqWithoutIsStrInference([d_400_v46_, d_400_v46_])
            d_402_v48_: _dafny.Array
            nw55_ = _dafny.Array(None, 11)
            nw55_[int(0)] = d_400_v46_
            nw55_[int(1)] = d_400_v46_
            nw55_[int(2)] = d_400_v46_
            nw55_[int(3)] = d_400_v46_
            nw55_[int(4)] = d_400_v46_
            nw55_[int(5)] = d_400_v46_
            nw55_[int(6)] = (d_401_v47_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_401_v47_))]
            nw55_[int(7)] = d_400_v46_
            nw55_[int(8)] = d_400_v46_
            nw55_[int(9)] = d_400_v46_
            nw55_[int(10)] = d_400_v46_
            d_402_v48_ = nw55_
            index44_ = default__.safeIndex(135, (d_402_v48_).length(0))
            (d_402_v48_)[index44_] = d_400_v46_
            index45_ = default__.safeIndex(135, (d_402_v48_).length(0))
            (d_402_v48_)[index45_] = d_400_v46_
            d_349_v14_ = (d_349_v14_).set(not((815) < (d_329_v4_)), default__.safeDivisionInt(d_329_v4_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]))
        elif True:
            (d_335_globalState_).f17 = (d_329_v4_) - ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])
            d_403_v49_: _dafny.Map
            d_403_v49_ = _dafny.Map({(0) - (d_329_v4_): (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
            (d_335_globalState_).f17 = len(((d_403_v49_).set((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (0) - (d_329_v4_))) | ((d_403_v49_) | (d_403_v49_)))
            d_329_v4_ = ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))] if not (d_325_v0_) or (d_325_v0_) else (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])
            if not (d_325_v0_) or (((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) == ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])):
                index46_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index46_] = (d_329_v4_) - (len(d_331_v6_))
                (d_335_globalState_).f6 = d_325_v0_
                (d_335_globalState_).f6 = d_325_v0_
                d_333_v8_ = d_333_v8_
                d_404_v50_: _dafny.MultiSet
                d_404_v50_ = _dafny.MultiSet([(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]])
                (d_335_globalState_).f6 = (_dafny.MultiSet(d_343_v13_)).issubset((d_404_v50_).set(d_329_v4_, default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emuwpm"))))))
            elif True:
                d_325_v0_ = d_325_v0_
                index47_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index47_] = d_329_v4_
                d_405_v51_: D14
                d_405_v51_ = D14_DC39(d_327_v2_)
                d_406_v52_: _dafny.Map
                d_406_v52_ = _dafny.Map({d_325_v0_: d_327_v2_})
                d_407_v53_: _dafny.Seq
                d_407_v53_ = _dafny.SeqWithoutIsStrInference([d_327_v2_, d_327_v2_, ((d_406_v52_)[d_325_v0_] if (d_325_v0_) in (d_406_v52_) else d_327_v2_), d_327_v2_])
                d_408_v54_: _dafny.MultiSet
                d_408_v54_ = _dafny.MultiSet([(d_327_v2_ if d_325_v0_ else (d_405_v51_).cf72), d_327_v2_, (d_407_v53_)[default__.safeIndex(default__.fm0(d_329_v4_, d_342_v12_, d_335_globalState_), len(d_407_v53_))], d_327_v2_, d_327_v2_])
                d_409_v55_: _dafny.MultiSet
                d_409_v55_ = _dafny.MultiSet([d_327_v2_, d_327_v2_, d_327_v2_])
                rhs33_ = (d_409_v55_)
                rhs34_ = d_333_v8_
                d_408_v54_ = rhs33_
                d_333_v8_ = rhs34_
                (d_335_globalState_).f17 = default__.safeDivisionInt(default__.safeDivisionInt(d_329_v4_, d_329_v4_), (d_329_v4_) * ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]))
                d_410_v56_: _dafny.Map
                d_410_v56_ = _dafny.Map({d_328_v3_: default__.fm39(_dafny.SeqWithoutIsStrInference([d_329_v4_, d_329_v4_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]]), d_325_v0_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (0) - ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]), d_335_globalState_)})
                d_411_v57_: _dafny.MultiSet
                d_411_v57_ = _dafny.MultiSet([d_340_v10_, default__.fm1(d_335_globalState_), d_340_v10_, d_340_v10_])
                (d_335_globalState_).f5 = default__.fm2(((d_410_v56_)[d_328_v3_] if (d_328_v3_) in (d_410_v56_) else d_411_v57_), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
            d_412_v58_: _dafny.Array
            nw56_ = _dafny.Array(_dafny.Map({}), 29)
            d_412_v58_ = nw56_
            d_413_v59_: _dafny.Map
            d_413_v59_ = _dafny.Map({d_325_v0_: not(d_325_v0_)})
            index48_ = default__.safeIndex(398, (d_412_v58_).length(0))
            (d_412_v58_)[index48_] = d_413_v59_
            index49_ = default__.safeIndex(398, (d_412_v58_).length(0))
            (d_412_v58_)[index49_] = d_413_v59_
        d_414_v60_: _dafny.Array
        nw57_ = _dafny.Array(None, 17)
        d_414_v60_ = nw57_
        d_415_v61_: _dafny.Seq
        d_415_v61_ = _dafny.SeqWithoutIsStrInference([d_414_v60_])
        hi1_ = d_329_v4_
        for d_416_i7_ in range(len(d_415_v61_), hi1_):
            d_417_v62_: _dafny.MultiSet
            d_417_v62_ = _dafny.MultiSet([d_340_v10_, d_340_v10_])
            d_418_v63_: _dafny.Seq
            d_418_v63_ = _dafny.SeqWithoutIsStrInference([d_325_v0_, not(d_325_v0_), not(False), default__.fm2(d_417_v62_, d_329_v4_, d_335_globalState_)])
            d_419_v64_: _dafny.Map
            d_419_v64_ = _dafny.Map({(d_418_v63_) + (d_418_v63_): (d_325_v0_) and (d_325_v0_)})
            d_419_v64_ = d_419_v64_
            d_420_v65_: T1
            nw58_ = C1()
            nw58_.ctor__(d_329_v4_, d_326_v1_)
            d_420_v65_ = nw58_
            d_421_v66_: _dafny.Array
            nw59_ = _dafny.Array(None, 1)
            nw59_[int(0)] = d_420_v65_
            d_421_v66_ = nw59_
            d_422_v67_: D18
            d_422_v67_ = D18_DC49(d_332_v7_)
            default__.m13(d_421_v66_, d_325_v0_, (_dafny.SeqWithoutIsStrInference([d_416_i7_ for d_423_i8_ in range(default__.abs(669))])) < (d_343_v13_), (((_dafny.MultiSet([d_422_v67_, d_422_v67_, d_422_v67_])).set(d_422_v67_, default__.abs(d_416_i7_))).cardinality) - (len(d_418_v63_)), d_335_globalState_)
            (d_335_globalState_).f6 = d_325_v0_
            d_424_v68_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.Seq({}), 10)
            d_424_v68_ = nw60_
            d_424_v68_ = d_424_v68_
        (d_335_globalState_).f18 = (d_329_v4_) * (d_329_v4_)
        d_425_v69_: _dafny.Array
        nw61_ = _dafny.Array(None, 23)
        d_425_v69_ = nw61_
        d_426_v70_: _dafny.Seq
        d_426_v70_ = _dafny.SeqWithoutIsStrInference([d_325_v0_, False, d_325_v0_])
        default__.m13(d_425_v69_, d_325_v0_, (False if d_325_v0_ else (d_426_v70_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_426_v70_))]), d_329_v4_, d_335_globalState_)
        if (d_325_v0_) in (d_426_v70_):
            d_427_v71_: T0
            nw62_ = C6()
            nw62_.ctor__(d_326_v1_)
            d_427_v71_ = nw62_
            d_428_v72_: _dafny.Map
            d_428_v72_ = _dafny.Map({d_427_v71_: d_329_v4_})
            d_429_v73_: _dafny.Map
            d_429_v73_ = _dafny.Map({d_329_v4_: d_428_v72_})
            d_430_v74_: _dafny.Map
            d_430_v74_ = _dafny.Map({(d_325_v0_ if d_325_v0_ else d_325_v0_): ((d_429_v73_)[(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]] if ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) in (d_429_v73_) else d_428_v72_)})
            d_430_v74_ = d_430_v74_
            d_431_v75_: C4
            nw63_ = C4()
            nw63_.ctor__((d_427_v71_).f20)
            d_431_v75_ = nw63_
            d_432_v76_: _dafny.Seq
            d_432_v76_ = _dafny.SeqWithoutIsStrInference([d_333_v8_])
            d_433_v77_: _dafny.MultiSet
            d_433_v77_ = _dafny.MultiSet([(d_432_v76_)[default__.safeIndex(len(d_343_v13_), len(d_432_v76_))]])
            d_434_v78_: _dafny.Map
            d_434_v78_ = _dafny.Map({(d_333_v8_) in (d_433_v77_): d_325_v0_})
            d_434_v78_ = (d_434_v78_).set(d_325_v0_, d_325_v0_)
            d_435_v79_: _dafny.Seq
            d_435_v79_ = _dafny.SeqWithoutIsStrInference([d_327_v2_])
            d_436_v80_: _dafny.Map
            d_436_v80_ = _dafny.Map({d_329_v4_: (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
            (d_335_globalState_).f5 = (len(d_435_v79_)) in ((default__.fm31(d_325_v0_, d_343_v13_, d_335_globalState_)) | (d_436_v80_))
            d_437_v81_: _dafny.Array
            def lambda9_(d_438_v13_):
                def lambda10_(d_439_i9_):
                    return d_438_v13_

                return lambda10_

            init4_ = lambda9_(d_343_v13_)
            nw64_ = _dafny.Array(None, 24)
            for i0_4_ in range(nw64_.length(0)):
                nw64_[i0_4_] = init4_(i0_4_)
            d_437_v81_ = nw64_
            d_437_v81_ = d_437_v81_
        elif True:
            d_440_v82_: _dafny.MultiSet
            d_440_v82_ = _dafny.MultiSet([d_329_v4_])
            index50_ = default__.safeIndex(334, (d_333_v8_).length(0))
            (d_333_v8_)[index50_] = (((0) - ((0) - ((d_440_v82_).cardinality))) - (d_329_v4_) if not(d_325_v0_) else (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])
            d_441_v83_: D21
            d_441_v83_ = D21_DC58(d_325_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))
            d_442_v84_: _dafny.Map
            d_442_v84_ = _dafny.Map({(d_441_v83_).cf108: d_327_v2_})
            d_443_v85_: D9
            d_443_v85_ = D9_DC24(d_442_v84_)
            source8_ = d_443_v85_
            if source8_.is_DC25:
                d_444___mcc_h6_ = source8_.cf46
                d_445___mcc_h7_ = source8_.cf47
                d_446___mcc_h8_ = source8_.cf48
                d_447_cf48_ = d_446___mcc_h8_
                d_448_cf47_ = d_445___mcc_h7_
                d_449_cf46_ = d_444___mcc_h6_
                d_450_v86_: C10
                nw65_ = C10()
                nw65_.ctor__((0) - (d_448_cf47_), d_326_v1_)
                d_450_v86_ = nw65_
                index51_ = default__.safeIndex(260, (d_327_v2_).length(0))
                (d_327_v2_)[index51_] = False
                index52_ = default__.safeIndex(260, (d_327_v2_).length(0))
                (d_327_v2_)[index52_] = (d_448_cf47_) > (len(d_426_v70_))
                d_451_v87_: _dafny.Array
                nw66_ = _dafny.Array(False, 12)
                d_451_v87_ = nw66_
                (d_335_globalState_).f4 = d_451_v87_
                d_452_v88_: D13
                d_452_v88_ = D13_DC36(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrxkqy"))))
                d_452_v88_ = d_452_v88_
            elif True:
                d_453___mcc_h9_ = source8_.cf45
                d_454_cf45_ = d_453___mcc_h9_
                index53_ = default__.safeIndex(791, (d_327_v2_).length(0))
                (d_327_v2_)[index53_] = (_dafny.MultiSet([(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]])).ispropersubset(_dafny.MultiSet([(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]]))
                d_455_v89_: _dafny.MultiSet
                d_455_v89_ = _dafny.MultiSet([d_325_v0_, d_325_v0_])
                index54_ = default__.safeIndex(791, (d_327_v2_).length(0))
                (d_327_v2_)[index54_] = not((d_329_v4_) <= ((((d_455_v89_)[d_325_v0_] if (d_325_v0_) in (d_455_v89_) else (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) if (d_426_v70_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_329_v4_, d_329_v4_, d_329_v4_})])), len(d_426_v70_))] else (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])))
                d_456_v90_: _dafny.MultiSet
                d_456_v90_ = _dafny.MultiSet([d_340_v10_, _dafny.CodePoint('q'), d_340_v10_])
                default__.m13((d_425_v69_ if default__.fm2(d_456_v90_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_) else d_425_v69_), (d_327_v2_)[default__.safeIndex(791, (d_327_v2_).length(0))], (d_327_v2_)[default__.safeIndex(791, (d_327_v2_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
                d_328_v3_ = d_328_v3_
                index55_ = default__.safeIndex(817, (d_334_v9_).length(0))
                (d_334_v9_)[index55_] = d_328_v3_
                d_457_v91_: _dafny.Set
                d_457_v91_ = _dafny.Set({(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
                d_458_v92_: D13
                d_458_v92_ = D13_DC35(d_457_v91_)
                d_459_v93_: D22
                d_459_v93_ = D22_DC61(default__.fm0((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_342_v12_, d_335_globalState_), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_325_v0_, d_458_v92_, d_329_v4_)
                d_460_v94_: _dafny.Map
                d_460_v94_ = _dafny.Map({_dafny.Set({543}): len(d_328_v3_)})
                pat_let_tv19_ = d_458_v92_
                pat_let_tv20_ = d_325_v0_
                pat_let_tv21_ = d_333_v8_
                pat_let_tv22_ = d_333_v8_
                index56_ = default__.safeIndex(817, (d_334_v9_).length(0))
                def iife86_(_pat_let27_0):
                    def iife87_(d_461_dt__update__tmp_h2_):
                        def iife88_(_pat_let28_0):
                            def iife89_(d_462_dt__update_hcf118_h0_):
                                def iife90_(_pat_let29_0):
                                    def iife91_(d_463_dt__update_hcf117_h0_):
                                        def iife92_(_pat_let30_0):
                                            def iife93_(d_464_dt__update_hcf119_h0_):
                                                return D22_DC61((d_461_dt__update__tmp_h2_).cf115, (d_461_dt__update__tmp_h2_).cf116, d_463_dt__update_hcf117_h0_, d_462_dt__update_hcf118_h0_, d_464_dt__update_hcf119_h0_)
                                            return iife93_(_pat_let30_0)
                                        return iife92_((pat_let_tv22_)[default__.safeIndex(334, (pat_let_tv21_).length(0))])
                                    return iife91_(_pat_let29_0)
                                return iife90_(pat_let_tv20_)
                            return iife89_(_pat_let28_0)
                        return iife88_(pat_let_tv19_)
                    return iife87_(_pat_let27_0)
                (d_334_v9_)[index56_] = default__.fm27(default__.safeModuloInt(d_329_v4_, (iife86_(d_459_v93_)).cf116), default__.fm33(d_335_globalState_), default__.fm0(len(d_426_v70_), d_342_v12_, d_335_globalState_), d_460_v94_, d_335_globalState_)
            index57_ = default__.safeIndex(582, (d_327_v2_).length(0))
            (d_327_v2_)[index57_] = d_325_v0_
            index58_ = default__.safeIndex(582, (d_327_v2_).length(0))
            (d_327_v2_)[index58_] = default__.fm2((_dafny.MultiSet([d_340_v10_])) | (_dafny.MultiSet([d_340_v10_])), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
            d_333_v8_ = d_333_v8_
            default__.m13(d_425_v69_, d_325_v0_, (d_327_v2_)[default__.safeIndex(582, (d_327_v2_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_334_v9_).length(0)):
            d_465_i10_: int = guard_loop_1_
            if (True) and (((0) <= (d_465_i10_)) and ((d_465_i10_) < ((d_334_v9_).length(0)))):
                _ingredientsd_0.append((d_334_v9_, int(d_465_i10_), ((d_328_v3_ if not(d_325_v0_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmm")))) + ((d_330_v5_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_330_v5_))])))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        if d_325_v0_:
            (d_335_globalState_).f18 = (len(d_328_v3_)) - (d_329_v4_)
            index59_ = default__.safeIndex(334, (d_333_v8_).length(0))
            (d_333_v8_)[index59_] = (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]
            d_466_v95_: _dafny.Set
            d_466_v95_ = _dafny.Set({636, default__.fm0(325, D0_DC1(d_340_v10_, len(d_341_v11_), d_341_v11_, d_325_v0_, d_325_v0_), d_335_globalState_), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
            d_467_v96_: _dafny.MultiSet
            d_467_v96_ = _dafny.MultiSet([d_325_v0_, d_325_v0_, d_325_v0_])
            d_468_v97_: _dafny.MultiSet
            d_468_v97_ = d_467_v96_
            d_469_v98_: _dafny.Array
            nw67_ = _dafny.Array(_dafny.Seq({}), 14)
            d_469_v98_ = nw67_
            d_470_v99_: D3
            d_470_v99_ = D3_DC7(d_469_v98_)
            d_471_v100_: C5
            nw68_ = C5()
            nw68_.ctor__(d_468_v97_, d_470_v99_, d_326_v1_)
            d_471_v100_ = nw68_
            d_472_v101_: _dafny.Map
            d_472_v101_ = _dafny.Map({d_471_v100_: d_325_v0_})
            rhs35_ = (d_325_v0_) in (default__.fm9(d_466_v95_, not(d_325_v0_), d_335_globalState_))
            rhs36_ = d_329_v4_
            rhs37_ = default__.fm0((0) - (((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) * ((0) - ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]))), (d_342_v12_ if d_325_v0_ else d_342_v12_), d_335_globalState_)
            rhs38_ = (len((d_472_v101_) | (_dafny.Map({d_471_v100_: True})))) <= (851)
            lhs32_ = d_335_globalState_
            lhs33_ = d_335_globalState_
            lhs34_ = d_335_globalState_
            lhs32_.f6 = rhs35_
            lhs33_.f18 = rhs36_
            d_329_v4_ = rhs37_
            lhs34_.f6 = rhs38_
            (d_335_globalState_).f5 = d_325_v0_
            if d_325_v0_:
                d_473_v102_: C4
                nw69_ = C4()
                nw69_.ctor__(d_326_v1_)
                d_473_v102_ = nw69_
                d_474_v103_: _dafny.Array
                nw70_ = _dafny.Array(None, 6)
                nw70_[int(0)] = d_341_v11_
                nw70_[int(1)] = d_341_v11_
                nw70_[int(2)] = (d_341_v11_) - (d_341_v11_)
                nw70_[int(3)] = d_341_v11_
                nw70_[int(4)] = (d_341_v11_) - (d_341_v11_)
                nw70_[int(5)] = d_341_v11_
                d_474_v103_ = nw70_
                nw71_ = _dafny.Array(_dafny.Set({}), 14)
                d_474_v103_ = nw71_
                index60_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index60_] = (d_343_v13_)[default__.safeIndex(d_329_v4_, len(d_343_v13_))]
                (d_335_globalState_).f4 = d_327_v2_
                (d_335_globalState_).f17 = d_329_v4_
            elif True:
                d_334_v9_ = d_334_v9_
                (d_335_globalState_).f5 = (d_466_v95_).isdisjoint(_dafny.Set({len(default__.fm37(d_325_v0_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]}))
                d_475_v105_: _dafny.Array
                def lambda11_(d_476_v13_, d_477_v0_):
                    def lambda12_(d_478_i11_):
                        def iife94_():
                            coll32_ = _dafny.Map()
                            compr_32_: int
                            for compr_32_ in (d_476_v13_).Elements:
                                d_479_v104_: int = compr_32_
                                if (d_479_v104_) in (d_476_v13_):
                                    coll32_[default__.safeDivisionInt(d_479_v104_, 83)] = d_477_v0_
                            return _dafny.Map(coll32_)
                        return iife94_()
                        

                    return lambda12_

                init5_ = lambda11_(d_343_v13_, d_325_v0_)
                nw72_ = _dafny.Array(None, 23)
                for i0_5_ in range(nw72_.length(0)):
                    nw72_[i0_5_] = init5_(i0_5_)
                d_475_v105_ = nw72_
                d_475_v105_ = d_475_v105_
                (d_335_globalState_).f4 = d_327_v2_
                d_325_v0_ = d_325_v0_
        elif True:
            default__.m13(d_425_v69_, (d_329_v4_) >= (len(d_328_v3_)), d_325_v0_, d_329_v4_, d_335_globalState_)
            d_480_v106_: _dafny.Map
            d_480_v106_ = _dafny.Map({d_329_v4_: len(d_343_v13_)})
            d_481_v107_: C10
            nw73_ = C10()
            nw73_.ctor__(len(d_328_v3_), d_326_v1_)
            d_481_v107_ = nw73_
            d_482_v108_: _dafny.MultiSet
            d_482_v108_ = _dafny.MultiSet([d_481_v107_, d_481_v107_])
            default__.m13(d_425_v69_, False, (d_341_v11_).isdisjoint(d_341_v11_), ((d_480_v106_)[((d_482_v108_)[d_481_v107_] if (d_481_v107_) in (d_482_v108_) else 900)] if (((d_482_v108_)[d_481_v107_] if (d_481_v107_) in (d_482_v108_) else 900)) in (d_480_v106_) else d_329_v4_), d_335_globalState_)
            d_483_v109_: _dafny.Map
            d_483_v109_ = _dafny.Map({d_325_v0_: d_325_v0_})
            d_484_v110_: _dafny.Seq
            d_484_v110_ = _dafny.SeqWithoutIsStrInference([d_483_v109_])
            d_485_v111_: _dafny.MultiSet
            d_485_v111_ = _dafny.MultiSet([((d_484_v110_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_484_v110_))]) | (d_483_v109_), d_483_v109_, d_483_v109_])
            d_485_v111_ = d_485_v111_
            index61_ = default__.safeIndex(714, (d_327_v2_).length(0))
            (d_327_v2_)[index61_] = not (d_325_v0_) or (d_325_v0_)
            index62_ = default__.safeIndex(714, (d_327_v2_).length(0))
            (d_327_v2_)[index62_] = True
            d_486_v112_: _dafny.MultiSet
            d_486_v112_ = _dafny.MultiSet([(d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))]])
            d_487_v113_: _dafny.MultiSet
            d_487_v113_ = d_486_v112_
            d_488_v114_: _dafny.Array
            nw74_ = _dafny.Array(_dafny.Seq({}), 17)
            d_488_v114_ = nw74_
            d_489_v115_: D3
            d_489_v115_ = D3_DC7(d_488_v114_)
            d_490_v116_: C5
            nw75_ = C5()
            nw75_.ctor__(d_487_v113_, d_489_v115_, d_326_v1_)
            d_490_v116_ = nw75_
            d_491_v117_: _dafny.Map
            d_491_v117_ = _dafny.Map({(d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))]: d_490_v116_})
            index63_ = default__.safeIndex(714, (d_327_v2_).length(0))
            index64_ = default__.safeIndex(334, (d_333_v8_).length(0))
            rhs39_ = (d_426_v70_)[default__.safeIndex(d_329_v4_, len(d_426_v70_))]
            rhs40_ = False
            rhs41_ = (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]
            rhs42_ = ((d_491_v117_)[((d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))]) and ((d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))])] if (((d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))]) and ((d_327_v2_)[default__.safeIndex(714, (d_327_v2_).length(0))])) in (d_491_v117_) else d_490_v116_)
            lhs35_ = d_327_v2_
            lhs36_ = default__.safeIndex(714, (d_327_v2_).length(0))
            lhs37_ = d_335_globalState_
            lhs38_ = d_333_v8_
            lhs39_ = default__.safeIndex(334, (d_333_v8_).length(0))
            lhs35_[lhs36_] = rhs39_
            lhs37_.f6 = rhs40_
            lhs38_[lhs39_] = rhs41_
            d_490_v116_ = rhs42_
        d_492_v118_: _dafny.Set
        d_492_v118_ = _dafny.Set({(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]})
        d_493_v119_: C2
        nw76_ = C2()
        nw76_.ctor__(d_329_v4_)
        d_493_v119_ = nw76_
        d_494_v120_: _dafny.Map
        d_494_v120_ = _dafny.Map({d_493_v119_: d_325_v0_})
        d_495_v121_: _dafny.MultiSet
        d_495_v121_ = _dafny.MultiSet([d_340_v10_])
        d_496_v122_: _dafny.Array
        nw77_ = _dafny.Array(None, 16)
        nw77_[int(0)] = d_426_v70_
        nw77_[int(1)] = d_426_v70_
        nw77_[int(2)] = d_426_v70_
        nw77_[int(3)] = d_426_v70_
        nw77_[int(4)] = d_426_v70_
        nw77_[int(5)] = _dafny.SeqWithoutIsStrInference([False, d_325_v0_, d_325_v0_, d_325_v0_, d_325_v0_])
        nw77_[int(6)] = d_426_v70_
        nw77_[int(7)] = d_426_v70_
        nw77_[int(8)] = default__.fm9(d_492_v118_, ((d_494_v120_)[d_493_v119_] if (d_493_v119_) in (d_494_v120_) else d_325_v0_), d_335_globalState_)
        nw77_[int(9)] = d_426_v70_
        nw77_[int(10)] = d_426_v70_
        nw77_[int(11)] = d_426_v70_
        nw77_[int(12)] = _dafny.SeqWithoutIsStrInference([default__.fm2(d_495_v121_, d_329_v4_, d_335_globalState_)])
        nw77_[int(13)] = d_426_v70_
        nw77_[int(14)] = _dafny.SeqWithoutIsStrInference([d_325_v0_, (D7_DC21((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_325_v0_, d_325_v0_, d_333_v8_, d_325_v0_)).cf39, d_325_v0_, d_325_v0_, d_325_v0_])
        nw77_[int(15)] = _dafny.SeqWithoutIsStrInference([d_325_v0_])
        d_496_v122_ = nw77_
        d_497_v123_: D3
        d_497_v123_ = D3_DC8(d_325_v0_)
        d_498_v124_: D15
        d_498_v124_ = D15_DC43(d_496_v122_, True, (d_493_v119_).f22, (d_329_v4_) >= (d_329_v4_), d_497_v123_)
        source9_ = d_498_v124_
        if source9_.is_DC43:
            d_499___mcc_h10_ = source9_.cf82
            d_500___mcc_h11_ = source9_.cf83
            d_501___mcc_h12_ = source9_.cf84
            d_502___mcc_h13_ = source9_.cf85
            d_503___mcc_h14_ = source9_.cf86
            d_504_cf86_ = d_503___mcc_h14_
            d_505_cf85_ = d_502___mcc_h13_
            d_506_cf84_ = d_501___mcc_h12_
            d_507_cf83_ = d_500___mcc_h11_
            d_508_cf82_ = d_499___mcc_h10_
            d_509_v125_: T0
            nw78_ = C3()
            nw78_.ctor__(D0_DC0(d_505_cf85_))
            d_509_v125_ = nw78_
            d_509_v125_ = d_509_v125_
            d_510_v126_: _dafny.Array
            nw79_ = _dafny.Array(_dafny.Array(None, 0), 8)
            d_510_v126_ = nw79_
            d_511_v127_: D10
            d_511_v127_ = D10_DC27(d_510_v126_, d_329_v4_)
            d_512_v128_: _dafny.MultiSet
            d_512_v128_ = _dafny.MultiSet([d_506_cf84_, (d_511_v127_).cf51])
            d_513_v129_: _dafny.Map
            d_513_v129_ = _dafny.Map({(d_493_v119_).f22: d_325_v0_})
            d_512_v128_ = ((d_512_v128_) - (d_512_v128_)).set(d_506_cf84_, default__.abs((len(d_513_v129_)) * ((d_493_v119_).f22)))
            d_514_v130_: _dafny.Map
            d_514_v130_ = _dafny.Map({(d_493_v119_).f22: 721})
            d_515_v131_: C9
            nw80_ = C9()
            nw80_.ctor__(d_326_v1_)
            d_515_v131_ = nw80_
            d_516_v132_: _dafny.Map
            d_516_v132_ = _dafny.Map({d_515_v131_: d_329_v4_})
            if (((d_514_v130_)[(d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]] if ((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) in (d_514_v130_) else (d_493_v119_).f22)) < ((d_506_cf84_) * (len(d_516_v132_))):
                default__.m13(d_425_v69_, d_505_cf85_, (d_426_v70_)[default__.safeIndex((d_493_v119_).f22, len(d_426_v70_))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
                d_517_v133_: _dafny.MultiSet
                d_517_v133_ = _dafny.MultiSet([(d_329_v4_) > (220), d_505_cf85_])
                d_517_v133_ = d_517_v133_
                (d_335_globalState_).f6 = d_325_v0_
                d_518_v134_: _dafny.MultiSet
                d_518_v134_ = _dafny.MultiSet([d_512_v128_])
                d_519_v135_: _dafny.Seq
                d_519_v135_ = _dafny.SeqWithoutIsStrInference([d_509_v125_, d_509_v125_, d_509_v125_, d_509_v125_, d_509_v125_])
                d_520_v136_: _dafny.Map
                d_520_v136_ = _dafny.Map({d_518_v134_: d_519_v135_})
                (d_335_globalState_).f17 = (0) - (len(((d_520_v136_)[d_518_v134_] if (d_518_v134_) in (d_520_v136_) else (d_519_v135_ if d_325_v0_ else _dafny.SeqWithoutIsStrInference([d_509_v125_, d_509_v125_, d_509_v125_, d_509_v125_, d_509_v125_])))))
                d_521_v137_: bool
                d_522_v138_: _dafny.Array
                d_523_v139_: D2
                out1_: bool
                out2_: _dafny.Array
                out3_: D2
                out1_, out2_, out3_ = (d_493_v119_).m2(((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]) == ((default__.fm82(d_335_globalState_)).cf113), (d_493_v119_).f22, d_325_v0_, len(default__.fm45(d_335_globalState_)), d_335_globalState_)
                d_521_v137_ = out1_
                d_522_v138_ = out2_
                d_523_v139_ = out3_
            elif True:
                d_340_v10_ = _dafny.CodePoint('n')
                d_524_v140_: _dafny.MultiSet
                d_524_v140_ = _dafny.MultiSet([d_507_cf83_, d_325_v0_])
                d_525_v141_: _dafny.MultiSet
                d_525_v141_ = d_524_v140_
                index65_ = default__.safeIndex(906, (d_496_v122_).length(0))
                (d_496_v122_)[index65_] = default__.fm50(True, d_507_cf83_, d_525_v141_, d_335_globalState_)
                d_526_v142_: _dafny.Seq
                d_526_v142_ = _dafny.SeqWithoutIsStrInference([d_514_v130_])
                index66_ = default__.safeIndex(906, (d_496_v122_).length(0))
                (d_496_v122_)[index66_] = _dafny.SeqWithoutIsStrInference([not(((d_493_v119_).f22) not in ((d_526_v142_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_526_v142_))]))])
                d_527_v143_: C1
                nw81_ = C1()
                nw81_.ctor__(d_329_v4_, (d_509_v125_).f20)
                d_527_v143_ = nw81_
                d_528_v144_: D5
                d_528_v144_ = D5_DC13(d_513_v129_)
                d_529_v145_: _dafny.Array
                nw82_ = _dafny.Array(None, 10)
                nw82_[int(0)] = d_528_v144_
                nw82_[int(1)] = d_528_v144_
                nw82_[int(2)] = d_528_v144_
                nw82_[int(3)] = d_528_v144_
                nw82_[int(4)] = d_528_v144_
                nw82_[int(5)] = d_528_v144_
                nw82_[int(6)] = d_528_v144_
                nw82_[int(7)] = d_528_v144_
                nw82_[int(8)] = d_528_v144_
                nw82_[int(9)] = d_528_v144_
                d_529_v145_ = nw82_
                d_530_v146_: _dafny.MultiSet
                d_530_v146_ = _dafny.MultiSet([d_529_v145_, d_529_v145_])
                d_531_v147_: _dafny.MultiSet
                d_531_v147_ = _dafny.MultiSet([d_530_v146_])
                index67_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index67_] = (((d_531_v147_)[d_530_v146_] if (d_530_v146_) in (d_531_v147_) else (d_493_v119_).f22)) + (d_506_cf84_)
                d_532_v148_: _dafny.Map
                d_532_v148_ = _dafny.Map({d_527_v143_: (len((d_496_v122_)[default__.safeIndex(906, (d_496_v122_).length(0))]) if d_507_cf83_ else len(d_426_v70_))})
                d_532_v148_ = (d_532_v148_).set(d_527_v143_, (d_493_v119_).f22)
            if ((0) - (len(d_426_v70_))) != ((d_493_v119_).f22):
                d_533_v149_: C4
                nw83_ = C4()
                nw83_.ctor__((d_509_v125_).f20)
                d_533_v149_ = nw83_
                d_341_v11_ = d_341_v11_
                d_534_v150_: C0
                nw84_ = C0()
                nw84_.ctor__(d_327_v2_, d_426_v70_)
                d_534_v150_ = nw84_
                d_535_v151_: _dafny.Map
                d_535_v151_ = _dafny.Map({(d_534_v150_ if d_325_v0_ else d_534_v150_): d_327_v2_})
                d_536_v152_: _dafny.Map
                d_536_v152_ = _dafny.Map({not(d_507_cf83_): d_328_v3_})
                d_537_v153_: _dafny.Map
                d_537_v153_ = _dafny.Map({d_505_cf85_: d_536_v152_})
                d_538_v154_: C7
                nw85_ = C7()
                nw85_.ctor__(d_507_cf83_, (d_509_v125_).f20)
                d_538_v154_ = nw85_
                d_539_v155_: _dafny.Seq
                d_539_v155_ = _dafny.SeqWithoutIsStrInference([d_538_v154_, d_538_v154_])
                d_540_v156_: D25
                d_540_v156_ = D25_DC67(d_515_v131_)
                d_541_v157_: _dafny.MultiSet
                d_541_v157_ = _dafny.MultiSet([(d_540_v156_).cf126])
                d_542_v158_: _dafny.Map
                d_542_v158_ = _dafny.Map({(d_539_v155_)[default__.safeIndex((d_541_v157_).cardinality, len(d_539_v155_))]: d_536_v152_})
                d_543_v159_: _dafny.Seq
                d_543_v159_ = _dafny.SeqWithoutIsStrInference([d_536_v152_, _dafny.Map({d_325_v0_: d_328_v3_})])
                d_544_v160_: _dafny.Seq
                d_544_v160_ = _dafny.SeqWithoutIsStrInference([d_333_v8_, d_333_v8_])
                d_545_v161_: _dafny.Map
                d_545_v161_ = d_536_v152_
                d_546_v162_: _dafny.Array
                nw86_ = _dafny.Array(None, 29)
                nw86_[int(0)] = default__.fm83((d_493_v119_).f22, (d_493_v119_).f22, len(d_328_v3_), d_335_globalState_)
                nw86_[int(1)] = d_536_v152_
                nw86_[int(2)] = (_dafny.Map({d_325_v0_: d_328_v3_}) if d_505_cf85_ else ((d_537_v153_)[False] if (False) in (d_537_v153_) else ((d_542_v158_)[(D26_DC69(d_538_v154_)).cf128] if ((D26_DC69(d_538_v154_)).cf128) in (d_542_v158_) else d_536_v152_)))
                nw86_[int(3)] = d_536_v152_
                nw86_[int(4)] = (_dafny.Map({(d_534_v150_).fm13((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_): d_328_v3_})) | (d_536_v152_)
                nw86_[int(5)] = (_dafny.Map({(d_538_v154_).f26: d_328_v3_})).set(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "np")))
                nw86_[int(6)] = default__.fm83((d_493_v119_).f22, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
                nw86_[int(7)] = ((d_543_v159_)[default__.safeIndex((0) - ((d_493_v119_).f22), len(d_543_v159_))]) | (d_536_v152_)
                nw86_[int(8)] = d_536_v152_
                nw86_[int(9)] = _dafny.Map({d_507_cf83_: d_328_v3_})
                nw86_[int(10)] = (d_536_v152_) | (((d_536_v152_).set(d_505_cf85_, d_328_v3_)).set(False, d_328_v3_))
                nw86_[int(11)] = (d_536_v152_).set(True, d_328_v3_)
                nw86_[int(12)] = default__.fm83((0) - (len(d_328_v3_)), len(d_544_v160_), (d_493_v119_).f22, d_335_globalState_)
                nw86_[int(13)] = d_536_v152_
                nw86_[int(14)] = (d_536_v152_).set(d_505_cf85_, d_328_v3_)
                nw86_[int(15)] = (d_536_v152_) | (_dafny.Map({d_325_v0_: d_328_v3_}))
                nw86_[int(16)] = d_536_v152_
                nw86_[int(17)] = d_536_v152_
                nw86_[int(18)] = _dafny.Map({(d_538_v154_).f26: d_328_v3_})
                nw86_[int(19)] = (d_536_v152_) | (d_536_v152_)
                nw86_[int(20)] = d_536_v152_
                nw86_[int(21)] = d_536_v152_
                nw86_[int(22)] = (d_536_v152_) | (d_536_v152_)
                nw86_[int(23)] = (d_536_v152_) | ((_dafny.Map({(d_538_v154_).f26: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qg"))})).set((d_538_v154_).f26, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oytvdn"))))
                nw86_[int(24)] = _dafny.Map({d_325_v0_: d_328_v3_})
                nw86_[int(25)] = (d_545_v161_)
                nw86_[int(26)] = d_536_v152_
                nw86_[int(27)] = (d_536_v152_) | (_dafny.Map({d_325_v0_: d_328_v3_}))
                nw86_[int(28)] = (d_536_v152_) | (d_536_v152_)
                d_546_v162_ = nw86_
                index68_ = default__.safeIndex(484, (d_546_v162_).length(0))
                (d_546_v162_)[index68_] = d_536_v152_
                index69_ = default__.safeIndex(484, (d_546_v162_).length(0))
                rhs43_ = d_535_v151_
                rhs44_ = d_536_v152_
                rhs45_ = d_534_v150_.f23
                rhs46_ = d_340_v10_
                lhs40_ = d_546_v162_
                lhs41_ = default__.safeIndex(484, (d_546_v162_).length(0))
                lhs42_ = d_534_v150_
                d_535_v151_ = rhs43_
                lhs40_[lhs41_] = rhs44_
                lhs42_.f23 = rhs45_
                d_340_v10_ = rhs46_
                d_547_v163_: _dafny.Map
                d_547_v163_ = _dafny.Map({((d_493_v119_).f22) * (len(d_492_v118_)): d_534_v150_.f23})
                rhs47_ = d_507_cf83_
                rhs48_ = d_547_v163_
                lhs43_ = d_335_globalState_
                lhs43_.f6 = rhs47_
                d_547_v163_ = rhs48_
                (d_335_globalState_).f5 = d_325_v0_
            elif True:
                d_548_v164_: _dafny.MultiSet
                d_548_v164_ = _dafny.MultiSet([D12_DC34(not(d_505_cf85_), d_505_cf85_, False)])
                d_549_v165_: D12
                d_549_v165_ = D12_DC34(d_505_cf85_, d_325_v0_, d_505_cf85_)
                d_550_v166_: T1
                nw87_ = C1()
                nw87_.ctor__(d_329_v4_, d_326_v1_)
                d_550_v166_ = nw87_
                d_551_v167_: _dafny.Map
                d_551_v167_ = _dafny.Map({d_550_v166_: _dafny.MultiSet([d_549_v165_])})
                d_548_v164_ = ((d_548_v164_) - (_dafny.MultiSet([d_549_v165_, d_549_v165_]))) - (((d_551_v167_)[d_550_v166_] if (d_550_v166_) in (d_551_v167_) else d_548_v164_))
                index70_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index70_] = -403
                d_506_cf84_ = len(d_328_v3_)
                d_552_v168_: _dafny.MultiSet
                d_552_v168_ = _dafny.MultiSet([d_426_v70_, d_426_v70_, (d_426_v70_) + (d_426_v70_)])
                index71_ = default__.safeIndex(334, (d_333_v8_).length(0))
                (d_333_v8_)[index71_] = ((d_552_v168_)[default__.fm40(d_325_v0_, d_506_cf84_, d_514_v130_, d_335_globalState_)] if (default__.fm40(d_325_v0_, d_506_cf84_, d_514_v130_, d_335_globalState_)) in (d_552_v168_) else d_506_cf84_)
                (d_335_globalState_).f6 = not(d_325_v0_)
        elif True:
            d_553___mcc_h15_ = source9_.cf81
            d_554_cf81_ = d_553___mcc_h15_
            if d_325_v0_:
                d_555_v169_: _dafny.Map
                d_555_v169_ = _dafny.Map({(D7_DC21((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], False, d_325_v0_, d_333_v8_, d_325_v0_)).cf38: not(False)})
                d_556_v170_: _dafny.Set
                d_556_v170_ = _dafny.Set({_dafny.Map({(d_493_v119_).f22: False}), d_555_v169_})
                d_557_v171_: _dafny.Seq
                d_557_v171_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_325_v0_: len(d_556_v170_)})])
                d_349_v14_ = ((d_349_v14_) | (d_349_v14_)) | (((d_557_v171_)[default__.safeIndex(677, len(d_557_v171_))]).set(d_325_v0_, len(_dafny.Map({d_325_v0_: (d_493_v119_).f22}))))
                (d_335_globalState_).f18 = 828
                d_558_v172_: D13
                d_558_v172_ = D13_DC35(d_492_v118_)
                d_559_v173_: D22
                d_559_v173_ = D22_DC61((d_493_v119_).f22, 239, False, d_558_v172_, d_329_v4_)
                d_560_v174_: _dafny.Map
                d_560_v174_ = _dafny.Map({d_325_v0_: False})
                d_349_v14_ = (d_349_v14_).set((d_559_v173_).cf117, len((d_560_v174_) | (d_560_v174_)))
                (d_335_globalState_).f18 = (0) - ((d_493_v119_).f22)
                d_561_v175_: C1
                nw88_ = C1()
                nw88_.ctor__((d_493_v119_).f22, D0_DC0(d_325_v0_))
                d_561_v175_ = nw88_
            elif True:
                d_562_v176_: C4
                nw89_ = C4()
                nw89_.ctor__(d_326_v1_)
                d_562_v176_ = nw89_
                d_563_v177_: D11
                d_563_v177_ = D11_DC31(_dafny.Map({d_328_v3_: d_327_v2_}), (d_493_v119_).f22)
                d_564_v178_: _dafny.Map
                d_564_v178_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbnewsawb")): d_327_v2_})
                d_565_v179_: T1
                nw90_ = C10()
                nw90_.ctor__(653, d_326_v1_)
                d_565_v179_ = nw90_
                d_566_v180_: _dafny.MultiSet
                d_566_v180_ = _dafny.MultiSet([d_565_v179_, d_565_v179_])
                pat_let_tv23_ = d_329_v4_
                pat_let_tv24_ = d_329_v4_
                d_567_v181_: _dafny.Array
                nw91_ = _dafny.Array(None, 26)
                nw91_[int(0)] = d_563_v177_
                nw91_[int(1)] = d_563_v177_
                nw91_[int(2)] = d_563_v177_
                nw91_[int(3)] = d_563_v177_
                nw91_[int(4)] = d_563_v177_
                def iife95_(_pat_let31_0):
                    def iife96_(d_568_dt__update__tmp_h3_):
                        def iife97_(_pat_let32_0):
                            def iife98_(d_569_dt__update_hcf60_h0_):
                                return D11_DC31((d_568_dt__update__tmp_h3_).cf59, d_569_dt__update_hcf60_h0_)
                            return iife98_(_pat_let32_0)
                        return iife97_(pat_let_tv23_)
                    return iife96_(_pat_let31_0)
                nw91_[int(5)] = iife95_(d_563_v177_)
                nw91_[int(6)] = d_563_v177_
                nw91_[int(7)] = D11_DC31(d_564_v178_, 983)
                nw91_[int(8)] = d_563_v177_
                nw91_[int(9)] = d_563_v177_
                nw91_[int(10)] = d_563_v177_
                nw91_[int(11)] = d_563_v177_
                nw91_[int(12)] = d_563_v177_
                nw91_[int(13)] = D11_DC31(_dafny.Map({d_328_v3_: d_327_v2_}), default__.fm0((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], D0_DC1(d_340_v10_, (d_566_v180_).cardinality, _dafny.Set({True}), d_325_v0_, d_325_v0_), d_335_globalState_))
                nw91_[int(14)] = d_563_v177_
                nw91_[int(15)] = D11_DC31(d_564_v178_, (d_493_v119_).f22)
                nw91_[int(16)] = d_563_v177_
                nw91_[int(17)] = D11_DC31(d_564_v178_, d_329_v4_)
                nw91_[int(18)] = d_563_v177_
                nw91_[int(19)] = D11_DC31(d_564_v178_, (d_343_v13_)[default__.safeIndex(464, len(d_343_v13_))])
                nw91_[int(20)] = D11_DC31(d_564_v178_, len(d_343_v13_))
                def iife99_(_pat_let33_0):
                    def iife100_(d_570_dt__update__tmp_h4_):
                        def iife101_(_pat_let34_0):
                            def iife102_(d_571_dt__update_hcf60_h1_):
                                return D11_DC31((d_570_dt__update__tmp_h4_).cf59, d_571_dt__update_hcf60_h1_)
                            return iife102_(_pat_let34_0)
                        return iife101_(pat_let_tv24_)
                    return iife100_(_pat_let33_0)
                nw91_[int(21)] = iife99_(d_563_v177_)
                nw91_[int(22)] = d_563_v177_
                nw91_[int(23)] = d_563_v177_
                nw91_[int(24)] = d_563_v177_
                nw91_[int(25)] = d_563_v177_
                d_567_v181_ = nw91_
                d_572_v182_: _dafny.Map
                d_572_v182_ = _dafny.Map({d_567_v181_: ((d_328_v3_).set(default__.safeIndex(301, len(d_328_v3_)), d_340_v10_)) < (d_328_v3_)})
                rhs49_ = ((d_572_v182_)[d_567_v181_] if (d_567_v181_) in (d_572_v182_) else (True) == (d_325_v0_))
                rhs50_ = default__.safeDivisionInt(default__.safeDivisionInt(d_329_v4_, 74), len(_dafny.SeqWithoutIsStrInference([True, d_325_v0_])))
                lhs44_ = d_335_globalState_
                lhs45_ = d_335_globalState_
                lhs44_.f5 = rhs49_
                lhs45_.f17 = rhs50_
                (d_335_globalState_).f6 = d_325_v0_
                (d_335_globalState_).f6 = (d_426_v70_)[default__.safeIndex((d_493_v119_).f22, len(d_426_v70_))]
                (d_335_globalState_).f17 = (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]
            (d_335_globalState_).f17 = default__.fm0((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_342_v12_, d_335_globalState_)
            (d_335_globalState_).f6 = (d_426_v70_)[default__.safeIndex((d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], len(d_426_v70_))]
            d_573_v183_: _dafny.Map
            d_573_v183_ = _dafny.Map({d_343_v13_: d_341_v11_})
            d_574_v184_: _dafny.Map
            d_574_v184_ = _dafny.Map({d_349_v14_: d_341_v11_})
            d_575_v185_: _dafny.Array
            nw92_ = _dafny.Array(None, 2)
            nw92_[int(0)] = ((d_573_v183_)[d_343_v13_] if (d_343_v13_) in (d_573_v183_) else d_341_v11_)
            nw92_[int(1)] = (((d_574_v184_)[d_349_v14_] if (d_349_v14_) in (d_574_v184_) else d_341_v11_) if not(not(d_325_v0_)) else _dafny.Set({d_325_v0_}))
            d_575_v185_ = nw92_
            index72_ = default__.safeIndex(887, (d_575_v185_).length(0))
            (d_575_v185_)[index72_] = d_341_v11_
            index73_ = default__.safeIndex(887, (d_575_v185_).length(0))
            (d_575_v185_)[index73_] = default__.fm16(d_328_v3_, d_335_globalState_)
        d_576_i12_: int
        d_576_i12_ = 0
        with _dafny.label("0"):
            while d_325_v0_:
                with _dafny.c_label("0"):
                    if (d_576_i12_) >= (100):
                        raise _dafny.Break("0")
                    d_576_i12_ = (d_576_i12_) + (1)
                    d_577_v186_: C6
                    nw93_ = C6()
                    nw93_.ctor__(d_326_v1_)
                    d_577_v186_ = nw93_
                    d_578_v187_: _dafny.Array
                    nw94_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_578_v187_ = nw94_
                    index74_ = default__.safeIndex(726, (d_578_v187_).length(0))
                    (d_578_v187_)[index74_] = (d_333_v8_ if d_325_v0_ else d_333_v8_)
                    index75_ = default__.safeIndex(726, (d_578_v187_).length(0))
                    (d_578_v187_)[index75_] = d_333_v8_
                    (d_335_globalState_).f18 = default__.fm0((805) + ((d_493_v119_).f22), d_342_v12_, d_335_globalState_)
                    d_579_v188_: _dafny.Map
                    d_579_v188_ = _dafny.Map({(default__.fm60((d_493_v119_).f22, d_335_globalState_)) - (_dafny.Set({(d_493_v119_).f22, (d_493_v119_).f22})): (d_493_v119_).f22})
                    d_579_v188_ = (d_579_v188_).set(d_492_v118_, d_329_v4_)
                    pass
            pass
        d_580_i13_: int
        d_580_i13_ = 0
        with _dafny.label("1"):
            while d_325_v0_:
                with _dafny.c_label("1"):
                    if (d_580_i13_) >= (100):
                        raise _dafny.Break("1")
                    d_580_i13_ = (d_580_i13_) + (1)
                    if d_325_v0_:
                        index76_ = default__.safeIndex(811, (d_327_v2_).length(0))
                        (d_327_v2_)[index76_] = d_325_v0_
                        index77_ = default__.safeIndex(811, (d_327_v2_).length(0))
                        rhs51_ = (default__.fm0(len(d_343_v13_), d_342_v12_, d_335_globalState_)) > ((d_493_v119_).f22)
                        rhs52_ = d_325_v0_
                        rhs53_ = default__.fm2((d_495_v121_) | (_dafny.MultiSet(d_328_v3_)), (0) - (d_329_v4_), d_335_globalState_)
                        lhs46_ = d_335_globalState_
                        lhs47_ = d_327_v2_
                        lhs48_ = default__.safeIndex(811, (d_327_v2_).length(0))
                        lhs49_ = d_335_globalState_
                        lhs46_.f6 = rhs51_
                        lhs47_[lhs48_] = rhs52_
                        lhs49_.f6 = rhs53_
                        (d_335_globalState_).f6 = d_325_v0_
                        d_581_v189_: _dafny.Array
                        nw95_ = _dafny.Array(_dafny.Seq({}), 9)
                        d_581_v189_ = nw95_
                        index78_ = default__.safeIndex(967, (d_581_v189_).length(0))
                        (d_581_v189_)[index78_] = (d_343_v13_).set(default__.safeIndex(len(d_328_v3_), len(d_343_v13_)), (d_493_v119_).f22)
                        index79_ = default__.safeIndex(967, (d_581_v189_).length(0))
                        (d_581_v189_)[index79_] = d_343_v13_
                        d_582_v190_: _dafny.MultiSet
                        d_582_v190_ = _dafny.MultiSet([(d_327_v2_)[default__.safeIndex(811, (d_327_v2_).length(0))]])
                        d_583_v191_: _dafny.Set
                        d_583_v191_ = _dafny.Set({default__.fm47(len(_dafny.Map({(d_327_v2_)[default__.safeIndex(811, (d_327_v2_).length(0))]: d_325_v0_})), (d_327_v2_)[default__.safeIndex(811, (d_327_v2_).length(0))], True, d_335_globalState_), d_343_v13_, _dafny.SeqWithoutIsStrInference([(d_582_v190_).cardinality, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]])})
                        d_584_v192_: D12
                        d_584_v192_ = D12_DC33(d_583_v191_, d_329_v4_)
                        d_585_v193_: _dafny.Map
                        d_585_v193_ = _dafny.Map({d_584_v192_: not((len(d_328_v3_)) > ((d_493_v119_).f22))})
                        d_585_v193_ = (d_585_v193_).set(d_584_v192_, not((d_327_v2_)[default__.safeIndex(811, (d_327_v2_).length(0))]))
                        d_325_v0_ = not((d_327_v2_)[default__.safeIndex(811, (d_327_v2_).length(0))])
                    elif True:
                        d_586_v194_: _dafny.Seq
                        d_586_v194_ = _dafny.SeqWithoutIsStrInference([d_333_v8_])
                        d_587_v195_: _dafny.Seq
                        d_587_v195_ = _dafny.SeqWithoutIsStrInference([(d_586_v194_)[default__.safeIndex(d_329_v4_, len(d_586_v194_))], d_333_v8_, d_333_v8_, d_333_v8_])
                        d_588_v196_: _dafny.Array
                        nw96_ = _dafny.Array(None, 19)
                        nw96_[int(0)] = d_333_v8_
                        nw96_[int(1)] = d_333_v8_
                        nw96_[int(2)] = d_333_v8_
                        nw96_[int(3)] = d_333_v8_
                        nw96_[int(4)] = d_333_v8_
                        nw96_[int(5)] = d_333_v8_
                        nw96_[int(6)] = d_333_v8_
                        nw96_[int(7)] = d_333_v8_
                        nw96_[int(8)] = d_333_v8_
                        nw96_[int(9)] = d_333_v8_
                        nw96_[int(10)] = (d_586_v194_)[default__.safeIndex((0) - ((d_493_v119_).f22), len(d_586_v194_))]
                        nw96_[int(11)] = d_333_v8_
                        nw96_[int(12)] = d_333_v8_
                        nw96_[int(13)] = d_333_v8_
                        nw96_[int(14)] = d_333_v8_
                        nw96_[int(15)] = d_333_v8_
                        nw96_[int(16)] = d_333_v8_
                        nw96_[int(17)] = d_333_v8_
                        nw96_[int(18)] = d_333_v8_
                        d_588_v196_ = nw96_
                        index80_ = default__.safeIndex(732, (d_588_v196_).length(0))
                        (d_588_v196_)[index80_] = d_333_v8_
                        index81_ = default__.safeIndex(732, (d_588_v196_).length(0))
                        (d_588_v196_)[index81_] = d_333_v8_
                        rhs54_ = not (d_325_v0_) or (d_325_v0_)
                        rhs55_ = d_340_v10_
                        rhs56_ = d_588_v196_
                        rhs57_ = (_dafny.Map({d_325_v0_: (d_493_v119_).f22})) == (_dafny.Map({d_325_v0_: d_329_v4_}))
                        lhs50_ = d_335_globalState_
                        lhs51_ = d_335_globalState_
                        lhs50_.f5 = rhs54_
                        d_340_v10_ = rhs55_
                        d_588_v196_ = rhs56_
                        lhs51_.f5 = rhs57_
                        d_589_v197_: _dafny.MultiSet
                        d_589_v197_ = _dafny.MultiSet([d_325_v0_])
                        d_590_v198_: _dafny.MultiSet
                        d_590_v198_ = d_589_v197_
                        d_591_v199_: _dafny.Array
                        nw97_ = _dafny.Array(_dafny.Seq({}), 14)
                        d_591_v199_ = nw97_
                        d_592_v200_: C5
                        nw98_ = C5()
                        nw98_.ctor__(d_590_v198_, D3_DC7(d_591_v199_), default__.fm57((d_493_v119_).f22, d_335_globalState_))
                        d_592_v200_ = nw98_
                        d_593_v201_: C4
                        nw99_ = C4()
                        nw99_.ctor__(d_326_v1_)
                        d_593_v201_ = nw99_
                        d_594_v202_: _dafny.Map
                        d_594_v202_ = _dafny.Map({(d_493_v119_).f22: d_593_v201_})
                        d_595_v203_: _dafny.MultiSet
                        d_595_v203_ = _dafny.MultiSet([len(d_328_v3_), len(d_594_v202_), (d_493_v119_).f22, 901, (d_493_v119_).f22])
                        d_596_v204_: bool
                        d_597_v205_: _dafny.Array
                        d_598_v206_: D2
                        out4_: bool
                        out5_: _dafny.Array
                        out6_: D2
                        out4_, out5_, out6_ = (d_493_v119_).m2(d_325_v0_, d_329_v4_, ((d_493_v119_).f22) >= ((d_595_v203_).cardinality), (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))], d_335_globalState_)
                        d_596_v204_ = out4_
                        d_597_v205_ = out5_
                        d_598_v206_ = out6_
                        index82_ = default__.safeIndex(334, (d_333_v8_).length(0))
                        (d_333_v8_)[index82_] = ((d_493_v119_).f22) - (817)
                    d_599_v207_: _dafny.MultiSet
                    d_599_v207_ = _dafny.MultiSet([d_325_v0_, d_325_v0_])
                    (d_335_globalState_).f6 = ((d_426_v70_)[default__.safeIndex((d_493_v119_).f22, len(d_426_v70_))]) and ((len(_dafny.Map({d_325_v0_: (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))]}))) != (((d_599_v207_)[d_325_v0_] if (d_325_v0_) in (d_599_v207_) else (d_493_v119_).f22)))
                    (d_335_globalState_).f4 = d_327_v2_
                    d_600_v208_: D6
                    d_600_v208_ = D6_DC18((d_493_v119_).f22, d_329_v4_, _dafny.SeqWithoutIsStrInference([d_325_v0_]), d_325_v0_, (d_333_v8_)[default__.safeIndex(334, (d_333_v8_).length(0))])
                    pat_let_tv25_ = d_328_v3_
                    def iife103_(_pat_let35_0):
                        def iife104_(d_601_dt__update__tmp_h5_):
                            def iife105_(_pat_let36_0):
                                def iife106_(d_602_dt__update_hcf34_h0_):
                                    return D6_DC18((d_601_dt__update__tmp_h5_).cf30, (d_601_dt__update__tmp_h5_).cf31, (d_601_dt__update__tmp_h5_).cf32, (d_601_dt__update__tmp_h5_).cf33, d_602_dt__update_hcf34_h0_)
                                return iife106_(_pat_let36_0)
                            return iife105_(len(pat_let_tv25_))
                        return iife104_(_pat_let35_0)
                    rhs58_ = _dafny.MultiSet((((iife103_(d_600_v208_)).cf32) + (d_426_v70_)) + (d_426_v70_))
                    rhs59_ = d_325_v0_
                    lhs52_ = d_335_globalState_
                    d_599_v207_ = rhs58_
                    lhs52_.f6 = rhs59_
                    pass
            pass
        _dafny.print(_dafny.string_of(d_325_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v1_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_327_v2_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_328_v3_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_329_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_330_v5_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntc"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_331_v6_) == (_dafny.Map({40: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntc"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_333_v8_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[23]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_334_v9_)[27]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_.f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_335_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_335_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_335_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f10) == (_dafny.Map({40: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntc"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_335_globalState_).f14)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_335_globalState_.f15) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_335_globalState_.f16).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_335_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_335_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[23]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_335_globalState_).f19)[27]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_340_v10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v11_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v12_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v12_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_342_v12_).cf3) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v12_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_342_v12_).cf5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v13_) == (_dafny.SeqWithoutIsStrInference([146, 40]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_349_v14_) == (_dafny.Map({False: 40, True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_415_v61_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_426_v70_) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_492_v118_) == (_dafny.Set({39}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_493_v119_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_494_v120_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_495_v121_) == (_dafny.MultiSet([_dafny.CodePoint('v')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[0]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[1]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[2]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[3]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[4]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[5]) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[6]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[7]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[8]) == (_dafny.SeqWithoutIsStrInference([False, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[9]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[11]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[12]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[13]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[14]) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_496_v122_)[15]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_497_v123_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[0]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[1]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[2]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[3]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[4]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[5]) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[6]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[7]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[8]) == (_dafny.SeqWithoutIsStrInference([False, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[9]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[11]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[12]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[13]) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[14]) == (_dafny.SeqWithoutIsStrInference([False, False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_498_v124_).cf82)[15]) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_498_v124_).cf83))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_498_v124_).cf84))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_498_v124_).cf85))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_498_v124_).cf86).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_576_i12_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_580_i13_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.CodePoint('D'), int(0), _dafny.Set({}), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
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
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC2(D1, NamedTuple('DC2', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(int(0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC4(D2, NamedTuple('DC4', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({self.cf7.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(False, False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D5_DC14(False)
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

class D5_DC14(D5, NamedTuple('DC14', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(_dafny.MultiSet({}), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC18(D6, NamedTuple('DC18', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC20(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)

class D7_DC20(D7, NamedTuple('DC20', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC19(D7, NamedTuple('DC19', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({self.cf44.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(int(0), int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)

class D9_DC25(D9, NamedTuple('DC25', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC27(_dafny.Array(None, 0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D10_DC26)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)

class D10_DC27(D10, NamedTuple('DC27', [('cf50', Any), ('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf50 == __o.cf50 and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC28(D10, NamedTuple('DC28', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({self.cf52.VerbatimString(True)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC26(D10, NamedTuple('DC26', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC26({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC26) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC31(_dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D11_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC31(D11, NamedTuple('DC31', [('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC31({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC31) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC30(D11, NamedTuple('DC30', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC33(_dafny.Set({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)

class D12_DC33(D12, NamedTuple('DC33', [('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [('cf64', Any), ('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC36(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC36(D13, NamedTuple('DC36', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC37(D13, NamedTuple('DC37', [('cf69', Any), ('cf70', Any), ('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38)
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC40(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D14_DC39)

class D14_DC40(D14, NamedTuple('DC40', [('cf73', Any), ('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf73)}, {self.cf74.VerbatimString(True)}, {self.cf75.VerbatimString(True)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC41(D14, NamedTuple('DC41', [('cf77', Any), ('cf78', Any), ('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC39(D14, NamedTuple('DC39', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC39({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC39) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(_dafny.Array(None, 0), False, int(0), False, D3.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)

class D15_DC43(D15, NamedTuple('DC43', [('cf82', Any), ('cf83', Any), ('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)}, {_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83 and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC45(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D16_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D16_DC44)

class D16_DC45(D16, NamedTuple('DC45', [('cf88', Any), ('cf89', Any), ('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC45({_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)}, {_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC45) and self.cf88 == __o.cf88 and self.cf89 == __o.cf89 and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC44(D16, NamedTuple('DC44', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC44({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC44) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC47(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D17_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D17_DC48)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D17_DC46)

class D17_DC47(D17, NamedTuple('DC47', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC47({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC47) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC48(D17, NamedTuple('DC48', [('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC48({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC48) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC46(D17, NamedTuple('DC46', [('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC46({_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC46) and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC50(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D18_DC50)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D18_DC49)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D18_DC51)

class D18_DC50(D18, NamedTuple('DC50', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC50({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC50) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC49(D18, NamedTuple('DC49', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC49({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC49) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC51(D18, NamedTuple('DC51', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC51({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC51) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D19_DC52)

class D19_DC52(D19, NamedTuple('DC52', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC52({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC52) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC54(int(0), _dafny.Array(None, 0), int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D20_DC54)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D20_DC55)

class D20_DC54(D20, NamedTuple('DC54', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC54({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC54) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC53(D20, NamedTuple('DC53', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC55(D20, NamedTuple('DC55', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC55({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC55) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC57(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D21_DC58)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)

class D21_DC57(D21, NamedTuple('DC57', [('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC58(D21, NamedTuple('DC58', [('cf108', Any), ('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC58({_dafny.string_of(self.cf108)}, {self.cf109.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC58) and self.cf108 == __o.cf108 and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [('cf106', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf106)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf106 == __o.cf106
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC60(int(0), _dafny.Map({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D22_DC61)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)

class D22_DC60(D22, NamedTuple('DC60', [('cf111', Any), ('cf112', Any), ('cf113', Any), ('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf111 == __o.cf111 and self.cf112 == __o.cf112 and self.cf113 == __o.cf113 and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC61(D22, NamedTuple('DC61', [('cf115', Any), ('cf116', Any), ('cf117', Any), ('cf118', Any), ('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC61({_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)}, {_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC61) and self.cf115 == __o.cf115 and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118 and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf110', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf110)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf110 == __o.cf110
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC63(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D23_DC63)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D23_DC64)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D23_DC65)

class D23_DC63(D23, NamedTuple('DC63', [('cf121', Any), ('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC63({_dafny.string_of(self.cf121)}, {self.cf122.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC63) and self.cf121 == __o.cf121 and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC64(D23, NamedTuple('DC64', [('cf123', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC64({_dafny.string_of(self.cf123)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC64) and self.cf123 == __o.cf123
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC62(D23, NamedTuple('DC62', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC65(D23, NamedTuple('DC65', [('cf124', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC65({_dafny.string_of(self.cf124)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC65) and self.cf124 == __o.cf124
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D24_DC66)

class D24_DC66(D24, NamedTuple('DC66', [('cf125', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC66({_dafny.string_of(self.cf125)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC66) and self.cf125 == __o.cf125
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC68(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D25_DC68)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D25_DC67)

class D25_DC68(D25, NamedTuple('DC68', [('cf127', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC68({_dafny.string_of(self.cf127)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC68) and self.cf127 == __o.cf127
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC67(D25, NamedTuple('DC67', [('cf126', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC67({_dafny.string_of(self.cf126)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC67) and self.cf126 == __o.cf126
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC70(_dafny.CodePoint('D'), int(0), False, _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D26_DC70)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D26_DC69)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)

class D26_DC70(D26, NamedTuple('DC70', [('cf129', Any), ('cf130', Any), ('cf131', Any), ('cf132', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC70({_dafny.string_of(self.cf129)}, {_dafny.string_of(self.cf130)}, {_dafny.string_of(self.cf131)}, {_dafny.string_of(self.cf132)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC70) and self.cf129 == __o.cf129 and self.cf130 == __o.cf130 and self.cf131 == __o.cf131 and self.cf132 == __o.cf132
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC69(D26, NamedTuple('DC69', [('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC69({_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC69) and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC71(D26, NamedTuple('DC71', [('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D27_DC72)

class D27_DC72(D27, NamedTuple('DC72', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC72({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC72) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: D28_DC74(False, _dafny.Set({}), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D28_DC74)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D28_DC73)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D28_DC75)

class D28_DC74(D28, NamedTuple('DC74', [('cf136', Any), ('cf137', Any), ('cf138', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC74({_dafny.string_of(self.cf136)}, {_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC74) and self.cf136 == __o.cf136 and self.cf137 == __o.cf137 and self.cf138 == __o.cf138
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC73(D28, NamedTuple('DC73', [('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC73({_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC73) and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()

class D28_DC75(D28, NamedTuple('DC75', [('cf139', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC75({_dafny.string_of(self.cf139)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC75) and self.cf139 == __o.cf139
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC77(int(0), int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D29_DC77)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D29_DC76)

class D29_DC77(D29, NamedTuple('DC77', [('cf141', Any), ('cf142', Any), ('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC77({_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC77) and self.cf141 == __o.cf141 and self.cf142 == __o.cf142 and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC76(D29, NamedTuple('DC76', [('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC76({_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC76) and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC79()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D30_DC79)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D30_DC80)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D30_DC78)

class D30_DC79(D30, NamedTuple('DC79', [])):
    def __dafnystr__(self) -> str:
        return f'D30.DC79'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC79)
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC80(D30, NamedTuple('DC80', [])):
    def __dafnystr__(self) -> str:
        return f'D30.DC80'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC80)
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC78(D30, NamedTuple('DC78', [('cf144', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC78({_dafny.string_of(self.cf144)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC78) and self.cf144 == __o.cf144
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D31_DC81)

class D31_DC81(D31, NamedTuple('DC81', [('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC81({_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC81) and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC83(int(0), int(0), False, int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D32_DC83)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D32_DC84)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D32_DC85)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D32_DC82)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D32_DC86)

class D32_DC83(D32, NamedTuple('DC83', [('cf147', Any), ('cf148', Any), ('cf149', Any), ('cf150', Any), ('cf151', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC83({_dafny.string_of(self.cf147)}, {_dafny.string_of(self.cf148)}, {_dafny.string_of(self.cf149)}, {_dafny.string_of(self.cf150)}, {_dafny.string_of(self.cf151)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC83) and self.cf147 == __o.cf147 and self.cf148 == __o.cf148 and self.cf149 == __o.cf149 and self.cf150 == __o.cf150 and self.cf151 == __o.cf151
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC84(D32, NamedTuple('DC84', [('cf152', Any), ('cf153', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC84({_dafny.string_of(self.cf152)}, {_dafny.string_of(self.cf153)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC84) and self.cf152 == __o.cf152 and self.cf153 == __o.cf153
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC85(D32, NamedTuple('DC85', [])):
    def __dafnystr__(self) -> str:
        return f'D32.DC85'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC85)
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC82(D32, NamedTuple('DC82', [('cf146', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC82({_dafny.string_of(self.cf146)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC82) and self.cf146 == __o.cf146
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC86(D32, NamedTuple('DC86', [('cf154', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC86({_dafny.string_of(self.cf154)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC86) and self.cf154 == __o.cf154
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: D33_DC88(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D33_DC88)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D33_DC87)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D33_DC89)

class D33_DC88(D33, NamedTuple('DC88', [('cf156', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC88({_dafny.string_of(self.cf156)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC88) and self.cf156 == __o.cf156
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC87(D33, NamedTuple('DC87', [('cf155', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC87({_dafny.string_of(self.cf155)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC87) and self.cf155 == __o.cf155
    def __hash__(self) -> int:
        return super().__hash__()

class D33_DC89(D33, NamedTuple('DC89', [('cf157', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC89({_dafny.string_of(self.cf157)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC89) and self.cf157 == __o.cf157
    def __hash__(self) -> int:
        return super().__hash__()


class D34:
    @classmethod
    def default(cls, ):
        return lambda: D34_DC91(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC91(self) -> bool:
        return isinstance(self, D34_DC91)
    @property
    def is_DC90(self) -> bool:
        return isinstance(self, D34_DC90)

class D34_DC91(D34, NamedTuple('DC91', [('cf159', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC91({_dafny.string_of(self.cf159)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC91) and self.cf159 == __o.cf159
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC90(D34, NamedTuple('DC90', [('cf158', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC90({_dafny.string_of(self.cf158)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC90) and self.cf158 == __o.cf158
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, p1, globalState):
        pass


class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f4: _dafny.Array = _dafny.Array(None, 0)
        self.f5: bool = False
        self.f6: bool = False
        self.f15: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f17: int = int(0)
        self.f18: int = int(0)
        self._f0: bool = False
        self._f1: bool = False
        self._f2: bool = False
        self._f3: bool = False
        self._f7: bool = False
        self._f8: int = int(0)
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f10: _dafny.Map = _dafny.Map({})
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f12: int = int(0)
        self._f13: int = int(0)
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self).f17 = f17
        (self).f18 = f18
        (self)._f19 = f19

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
    @property
    def f14(self):
        return self._f14
    @property
    def f19(self):
        return self._f19

class C0:
    def  __init__(self):
        self.f23: _dafny.Array = _dafny.Array(None, 0)
        self.f24: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f23, f24):
        (self).f23 = f23
        (self).f24 = f24

    def fm12(self, p0, p1, globalState):
        return _dafny.CodePoint('r')

    def fm13(self, p0, globalState):
        if (not(True) if False else True):
            return not (False) or (False)
        elif True:
            return not(True)


class C1(T1, T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f21, f20):
        (self)._f21 = f21
        (self)._f20 = f20

    def fm3(self, p0, p1, p2, globalState):
        def iife107_():
            coll33_ = _dafny.Map()
            compr_33_: int
            for compr_33_ in (_dafny.SeqWithoutIsStrInference([706 for d_603_i0_ in range(default__.abs(213))])).Elements:
                d_604_v0_: int = compr_33_
                if (d_604_v0_) in (_dafny.SeqWithoutIsStrInference([706 for d_603_i0_ in range(default__.abs(213))])):
                    coll33_[(d_604_v0_) + ((self).f21)] = False
            return _dafny.Map(coll33_)
        return (iife107_()
        ) | (_dafny.Map({(self).f21: not(False)}))

    def fm4(self, p0, p1, globalState):
        d_605_dt__update__tmp_h0_ = (self).f20
        d_606_dt__update_hcf0_h0_ = True
        return D0_DC0(d_606_dt__update_hcf0_h0_)

    def fm15(self, globalState):
        return True

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        d_607_v0_: _dafny.Seq
        d_607_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmttija"))
        d_608_v1_: _dafny.Map
        d_608_v1_ = _dafny.Map({p0: p1})
        d_609_v2_: _dafny.Seq
        d_609_v2_ = _dafny.SeqWithoutIsStrInference([len(d_607_v0_), len(d_608_v1_)])
        d_610_v3_: _dafny.Map
        d_610_v3_ = _dafny.Map({p0: len(d_609_v2_)})
        d_611_v4_: str
        d_611_v4_ = _dafny.CodePoint('b')
        d_612_v5_: _dafny.Set
        d_612_v5_ = _dafny.Set({p0})
        hi2_ = p1
        for d_613_i0_ in range(default__.fm0(((d_610_v3_)[p0] if (p0) in (d_610_v3_) else p1), D0_DC1(d_611_v4_, p1, d_612_v5_, p0, p0), globalState), hi2_):
            if p0:
                (globalState).f5 = p0
                d_614_v6_: _dafny.Array
                nw100_ = _dafny.Array(int(0), 7)
                d_614_v6_ = nw100_
                index83_ = default__.safeIndex(811, (d_614_v6_).length(0))
                (d_614_v6_)[index83_] = -15
                index84_ = default__.safeIndex(811, (d_614_v6_).length(0))
                (d_614_v6_)[index84_] = (0) - (d_613_i0_)
                d_615_v7_: _dafny.Array
                nw101_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_615_v7_ = nw101_
                d_616_v8_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.Set({}), 16)
                d_616_v8_ = nw102_
                d_617_v9_: _dafny.Set
                d_617_v9_ = _dafny.Set({(self).f21})
                index85_ = default__.safeIndex(702, (d_616_v8_).length(0))
                (d_616_v8_)[index85_] = d_617_v9_
                d_618_v10_: _dafny.Seq
                d_618_v10_ = _dafny.SeqWithoutIsStrInference([d_617_v9_])
                index86_ = default__.safeIndex(702, (d_616_v8_).length(0))
                rhs60_ = ((0) - (p1)) * (default__.safeDivisionInt(len(d_608_v1_), len(d_610_v3_)))
                rhs61_ = d_615_v7_
                rhs62_ = 500
                rhs63_ = (d_617_v9_).intersection((d_618_v10_)[default__.safeIndex((self).f21, len(d_618_v10_))])
                lhs53_ = globalState
                lhs54_ = globalState
                lhs55_ = d_616_v8_
                lhs56_ = default__.safeIndex(702, (d_616_v8_).length(0))
                lhs53_.f17 = rhs60_
                d_615_v7_ = rhs61_
                lhs54_.f18 = rhs62_
                lhs55_[lhs56_] = rhs63_
                (globalState).f17 = (d_614_v6_)[default__.safeIndex(811, (d_614_v6_).length(0))]
                d_609_v2_ = d_609_v2_
            elif True:
                (globalState).f6 = p0
                d_619_v11_: _dafny.Array
                nw103_ = _dafny.Array(int(0), 7)
                d_619_v11_ = nw103_
                d_620_v12_: _dafny.MultiSet
                d_620_v12_ = _dafny.MultiSet([p0])
                d_621_v13_: _dafny.Map
                d_621_v13_ = _dafny.Map({d_619_v11_: d_620_v12_})
                d_621_v13_ = (d_621_v13_).set(d_619_v11_, d_620_v12_)
                d_620_v12_ = d_620_v12_
                def iife108_():
                    coll34_ = _dafny.Map()
                    compr_34_: int
                    for compr_34_ in (d_609_v2_).Elements:
                        d_622_v14_: int = compr_34_
                        if (d_622_v14_) in (d_609_v2_):
                            coll34_[default__.safeModuloInt(d_622_v14_, (self).f21)] = d_607_v0_
                    return _dafny.Map(coll34_)
                (globalState).f6 = ((self).f21) in (iife108_()
                )
                (globalState).f5 = not(p0)
            (globalState).f18 = 174
            d_623_v15_: _dafny.Array
            def lambda13_(d_624_i1_):
                return False

            init6_ = lambda13_
            nw104_ = _dafny.Array(None, 2)
            for i0_6_ in range(nw104_.length(0)):
                nw104_[i0_6_] = init6_(i0_6_)
            d_623_v15_ = nw104_
            (globalState).f4 = (d_623_v15_ if False else d_623_v15_)
            d_625_v16_: _dafny.Seq
            d_625_v16_ = _dafny.SeqWithoutIsStrInference([p0])
            d_626_v17_: _dafny.Map
            d_626_v17_ = _dafny.Map({p0: (len(d_625_v16_)) == (p1)})
            d_626_v17_ = d_626_v17_
        d_627_v18_: _dafny.Map
        d_627_v18_ = _dafny.Map({D2_DC6(D2_DC6(D2_DC3(d_607_v0_))): p1})
        d_628_v19_: _dafny.MultiSet
        d_628_v19_ = _dafny.MultiSet([d_611_v4_])
        d_629_v20_: _dafny.MultiSet
        d_629_v20_ = _dafny.MultiSet([p0, default__.fm2(d_628_v19_, p1, globalState)])
        d_630_v21_: D2
        d_630_v21_ = D2_DC6(D2_DC5((self).fm15(globalState), (d_629_v20_).cardinality))
        d_627_v18_ = (d_627_v18_).set(d_630_v21_, (0) - (default__.safeModuloInt((self).f21, p1)))
        d_631_v22_: _dafny.Array
        def lambda14_(d_632_p0_):
            def lambda15_(d_633_i2_):
                return d_632_p0_

            return lambda15_

        init7_ = lambda14_(p0)
        nw105_ = _dafny.Array(None, 6)
        for i0_7_ in range(nw105_.length(0)):
            nw105_[i0_7_] = init7_(i0_7_)
        d_631_v22_ = nw105_
        d_634_v23_: C0
        nw106_ = C0()
        nw106_.ctor__(d_631_v22_, _dafny.SeqWithoutIsStrInference([p0, p0]))
        d_634_v23_ = nw106_
        d_635_v24_: _dafny.Array
        nw107_ = _dafny.Array(_dafny.Seq({}), 23)
        d_635_v24_ = nw107_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_635_v24_).length(0)):
            d_636_i3_: int = guard_loop_2_
            if (True) and (((0) <= (d_636_i3_)) and ((d_636_i3_) < ((d_635_v24_).length(0)))):
                def iife109_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in (_dafny.Map({p1: (_dafny.MultiSet([(0) - (p1), (self).f21])).cardinality})).keys.Elements:
                        d_637_v25_: int = compr_35_
                        if (d_637_v25_) in (_dafny.Map({p1: (_dafny.MultiSet([(0) - (p1), (self).f21])).cardinality})):
                            coll35_[(d_637_v25_) - (p1)] = _dafny.Map({False: p0})
                    return _dafny.Map(coll35_)
                (d_635_v24_)[(d_636_i3_)] = ((_dafny.SeqWithoutIsStrInference([len(iife109_()
) for d_638_i4_ in range(default__.abs(512))])) + (d_609_v2_)) + ((d_609_v2_ if p0 else d_609_v2_))
        d_639_i5_: int
        d_639_i5_ = 0
        with _dafny.label("2"):
            while (self).fm15(globalState):
                with _dafny.c_label("2"):
                    if (d_639_i5_) >= (100):
                        raise _dafny.Break("2")
                    d_639_i5_ = (d_639_i5_) + (1)
                    d_640_v26_: D0
                    d_640_v26_ = D0_DC1(d_611_v4_, (self).f21, default__.fm16(d_607_v0_, globalState), p0, p0)
                    (globalState).f17 = (default__.fm0((self).f21, d_640_v26_, globalState)) - ((self).f21)
                    d_641_v27_: _dafny.Array
                    def lambda16_(d_642_p0_):
                        def lambda17_(d_643_i6_):
                            return D2_DC4((self).f21, ((self).f20).cf0, d_642_p0_, d_642_p0_)

                        return lambda17_

                    init8_ = lambda16_(p0)
                    nw108_ = _dafny.Array(None, 7)
                    for i0_8_ in range(nw108_.length(0)):
                        nw108_[i0_8_] = init8_(i0_8_)
                    d_641_v27_ = nw108_
                    index87_ = default__.safeIndex(787, (d_641_v27_).length(0))
                    (d_641_v27_)[index87_] = D2_DC4((self).f21, p0, p0, p0)
                    d_644_v28_: D2
                    d_644_v28_ = D2_DC4((self).f21, (p0 if (self).fm15(globalState) else p0), ((self).f21) < (len(d_607_v0_)), p0)
                    index88_ = default__.safeIndex(787, (d_641_v27_).length(0))
                    (d_641_v27_)[index88_] = d_644_v28_
                    d_645_v29_: _dafny.Map
                    d_645_v29_ = _dafny.Map({(self).f21: d_634_v23_})
                    d_646_v30_: _dafny.Map
                    d_646_v30_ = _dafny.Map({p1: _dafny.CodePoint('l')})
                    d_647_v31_: _dafny.Map
                    d_647_v31_ = _dafny.Map({((self).f21 if p0 else len(d_646_v30_)): d_645_v29_})
                    d_648_v32_: _dafny.Seq
                    d_648_v32_ = _dafny.SeqWithoutIsStrInference([d_629_v20_])
                    d_645_v29_ = ((d_647_v31_)[((d_648_v32_)[default__.safeIndex((0) - ((self).f21), len(d_648_v32_))]).cardinality] if (((d_648_v32_)[default__.safeIndex((0) - ((self).f21), len(d_648_v32_))]).cardinality) in (d_647_v31_) else d_645_v29_)
                    d_649_v33_: D2
                    d_649_v33_ = D2_DC5(p0, (self).f21)
                    d_650_v34_: _dafny.Seq
                    d_650_v34_ = _dafny.SeqWithoutIsStrInference([default__.fm17(d_649_v33_, p0, globalState)])
                    d_651_v35_: _dafny.Array
                    nw109_ = _dafny.Array(None, 19)
                    nw109_[int(0)] = (p1) + ((self).f21)
                    nw109_[int(1)] = len(d_607_v0_)
                    nw109_[int(2)] = len(d_607_v0_)
                    nw109_[int(3)] = (d_609_v2_)[default__.safeIndex(679, len(d_609_v2_))]
                    nw109_[int(4)] = default__.safeModuloInt(179, (d_629_v20_).cardinality)
                    nw109_[int(5)] = 102
                    nw109_[int(6)] = (d_649_v33_).cf13
                    nw109_[int(7)] = len(d_650_v34_)
                    nw109_[int(8)] = ((d_629_v20_)[p0] if (p0) in (d_629_v20_) else (self).f21)
                    nw109_[int(9)] = p1
                    nw109_[int(10)] = p1
                    nw109_[int(11)] = (self).f21
                    nw109_[int(12)] = (self).f21
                    nw109_[int(13)] = p1
                    nw109_[int(14)] = p1
                    nw109_[int(15)] = default__.safeDivisionInt(p1, (self).f21)
                    nw109_[int(16)] = (635) * (p1)
                    nw109_[int(17)] = (self).f21
                    nw109_[int(18)] = (default__.fm18(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbilqwqgm")), False, p0, not(p0), globalState)).cardinality
                    d_651_v35_ = nw109_
                    pat_let_tv26_ = d_611_v4_
                    pat_let_tv27_ = p0
                    pat_let_tv28_ = p1
                    index89_ = default__.safeIndex(557, (d_651_v35_).length(0))
                    def iife110_(_pat_let37_0):
                        def iife111_(d_652_dt__update__tmp_h0_):
                            def iife112_(_pat_let38_0):
                                def iife113_(d_653_dt__update_hcf1_h0_):
                                    def iife114_(_pat_let39_0):
                                        def iife115_(d_654_dt__update_hcf4_h0_):
                                            def iife116_(_pat_let40_0):
                                                def iife117_(d_655_dt__update_hcf2_h0_):
                                                    return D0_DC1(d_653_dt__update_hcf1_h0_, d_655_dt__update_hcf2_h0_, (d_652_dt__update__tmp_h0_).cf3, d_654_dt__update_hcf4_h0_, (d_652_dt__update__tmp_h0_).cf5)
                                                return iife117_(_pat_let40_0)
                                            return iife116_(pat_let_tv28_)
                                        return iife115_(_pat_let39_0)
                                    return iife114_(pat_let_tv27_)
                                return iife113_(_pat_let38_0)
                            return iife112_(pat_let_tv26_)
                        return iife111_(_pat_let37_0)
                    (d_651_v35_)[index89_] = default__.fm0(p1, iife110_(d_640_v26_), globalState)
                    index90_ = default__.safeIndex(557, (d_651_v35_).length(0))
                    (d_651_v35_)[index90_] = len((_dafny.Set({p0})) | (d_612_v5_))
                    pass
            pass
        d_656_v36_: _dafny.Array
        nw110_ = _dafny.Array(_dafny.Map({}), 2)
        d_656_v36_ = nw110_
        d_657_v37_: _dafny.Map
        d_657_v37_ = _dafny.Map({618: (self).f21})
        d_658_v38_: _dafny.Map
        d_658_v38_ = _dafny.Map({((d_657_v37_)[679] if (679) in (d_657_v37_) else (self).f21): p0})
        index91_ = default__.safeIndex(530, (d_656_v36_).length(0))
        (d_656_v36_)[index91_] = d_658_v38_
        index92_ = default__.safeIndex(530, (d_656_v36_).length(0))
        (d_656_v36_)[index92_] = (d_658_v38_) | (d_658_v38_)
        r0 = (d_629_v20_) | (d_629_v20_)
        r1 = (((self).f21) + ((self).f21)) - (len(d_612_v5_))
        r2 = _dafny.CodePoint('q')
        return r0, r1, r2


class C2:
    def  __init__(self):
        self._f22: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f22):
        (self)._f22 = f22

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: D2 = D2.default()()
        d_659_v0_: _dafny.Array
        nw111_ = _dafny.Array(False, 3)
        d_659_v0_ = nw111_
        d_660_v1_: _dafny.Seq
        d_660_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epjj"))
        index93_ = default__.safeIndex(301, (d_659_v0_).length(0))
        (d_659_v0_)[index93_] = not((d_660_v1_) <= (d_660_v1_))
        index94_ = default__.safeIndex(301, (d_659_v0_).length(0))
        (d_659_v0_)[index94_] = p2
        index95_ = default__.safeIndex(301, (d_659_v0_).length(0))
        (d_659_v0_)[index95_] = not(p0)
        index96_ = default__.safeIndex(301, (d_659_v0_).length(0))
        (d_659_v0_)[index96_] = True
        d_661_v2_: str
        d_661_v2_ = _dafny.CodePoint('e')
        d_662_v3_: _dafny.MultiSet
        d_662_v3_ = _dafny.MultiSet([d_661_v2_])
        d_663_v4_: _dafny.Seq
        d_663_v4_ = _dafny.SeqWithoutIsStrInference([p0, p2, default__.fm2(d_662_v3_, p3, globalState)])
        if (d_663_v4_)[default__.safeIndex(p1, len(d_663_v4_))]:
            (globalState).f18 = p3
            d_664_v5_: _dafny.Map
            d_664_v5_ = _dafny.Map({p2: (self).f22})
            d_664_v5_ = _dafny.Map({p2: p3})
            (globalState).f17 = (self).f22
            d_665_v6_: _dafny.Set
            d_665_v6_ = _dafny.Set({(d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], p0, p0})
            d_666_v7_: D0
            d_666_v7_ = D0_DC1(d_661_v2_, p3, d_665_v6_, (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], False)
            d_667_v8_: _dafny.Map
            d_667_v8_ = _dafny.Map({(d_661_v2_ if p0 else (d_666_v7_).cf1): default__.fm8(True, (self).f22, globalState)})
            d_667_v8_ = d_667_v8_
            d_668_v9_: _dafny.Set
            d_668_v9_ = _dafny.Set({p3, p3, default__.fm0(p3, default__.fm10((self).f22, p2, (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], globalState), globalState), p3})
            r0 = (d_663_v4_) <= (default__.fm9(d_668_v9_, (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], globalState))
        elif True:
            d_669_v10_: _dafny.Map
            d_669_v10_ = _dafny.Map({_dafny.MultiSet([not(p0)]): p0})
            d_670_v11_: _dafny.MultiSet
            d_670_v11_ = _dafny.MultiSet([(d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], p0, True])
            d_659_v0_ = (d_659_v0_ if ((d_669_v10_)[d_670_v11_] if (d_670_v11_) in (d_669_v10_) else p0) else d_659_v0_)
            index97_ = default__.safeIndex(301, (d_659_v0_).length(0))
            rhs64_ = p2
            rhs65_ = (p2 if p2 else (793) not in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p3}))])))
            lhs57_ = d_659_v0_
            lhs58_ = default__.safeIndex(301, (d_659_v0_).length(0))
            lhs59_ = globalState
            lhs57_[lhs58_] = rhs64_
            lhs59_.f5 = rhs65_
            (globalState).f16 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csug"))
            (globalState).f5 = p0
            d_671_v12_: _dafny.Map
            d_671_v12_ = _dafny.Map({d_663_v4_: default__.fm11((d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], globalState)})
            d_672_v13_: _dafny.MultiSet
            d_672_v13_ = _dafny.MultiSet([p1, (self).f22, p1, len(((d_671_v12_)[d_663_v4_] if (d_663_v4_) in (d_671_v12_) else _dafny.SeqWithoutIsStrInference([d_661_v2_ for d_673_i0_ in range(default__.abs(657))])))])
            index98_ = default__.safeIndex(301, (d_659_v0_).length(0))
            rhs66_ = p2
            rhs67_ = d_672_v13_
            rhs68_ = d_660_v1_
            lhs60_ = d_659_v0_
            lhs61_ = default__.safeIndex(301, (d_659_v0_).length(0))
            lhs60_[lhs61_] = rhs66_
            d_672_v13_ = rhs67_
            d_660_v1_ = rhs68_
        d_674_v14_: C0
        nw112_ = C0()
        nw112_.ctor__(d_659_v0_, d_663_v4_)
        d_674_v14_ = nw112_
        d_675_v15_: D2
        d_675_v15_ = D2_DC4(p1, (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], p2, ((d_674_v14_.f24)[default__.safeIndex(p3, len(d_674_v14_.f24))]) or (p0))
        source10_ = d_675_v15_
        if source10_.is_DC4:
            d_676___mcc_h0_ = source10_.cf8
            d_677___mcc_h1_ = source10_.cf9
            d_678___mcc_h2_ = source10_.cf10
            d_679___mcc_h3_ = source10_.cf11
            d_680_cf11_ = d_679___mcc_h3_
            d_681_cf10_ = d_678___mcc_h2_
            d_682_cf9_ = d_677___mcc_h1_
            d_683_cf8_ = d_676___mcc_h0_
            d_684_v16_: _dafny.Array
            nw113_ = _dafny.Array(int(0), 19)
            d_684_v16_ = nw113_
            index99_ = default__.safeIndex(967, (d_684_v16_).length(0))
            (d_684_v16_)[index99_] = p3
            index100_ = default__.safeIndex(967, (d_684_v16_).length(0))
            (d_684_v16_)[index100_] = (default__.safeDivisionInt(92, d_683_cf8_)) + (len(d_663_v4_))
            d_683_cf8_ = p1
            d_680_cf11_ = ((self).f22) <= (((self).f22) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljajpoj")))))
            (globalState).f16 = default__.fm11((d_681_cf10_) and (d_680_cf11_), globalState)
        elif source10_.is_DC5:
            d_685___mcc_h4_ = source10_.cf12
            d_686___mcc_h5_ = source10_.cf13
            d_687_cf13_ = d_686___mcc_h5_
            d_688_cf12_ = d_685___mcc_h4_
            (globalState).f6 = True
            pat_let_tv29_ = d_674_v14_
            pat_let_tv30_ = d_674_v14_
            def iife118_(_pat_let41_0):
                def iife119_(d_690_dt__update__tmp_h0_):
                    def iife120_(_pat_let42_0):
                        def iife121_(d_691_dt__update_hcf12_h0_):
                            return D2_DC5(d_691_dt__update_hcf12_h0_, (d_690_dt__update__tmp_h0_).cf13)
                        return iife121_(_pat_let42_0)
                    return iife120_((pat_let_tv29_.f24)[default__.safeIndex((self).f22, len(pat_let_tv30_.f24))])
                return iife119_(_pat_let41_0)
            (globalState).f5 = (iife118_(D2_DC5(p0, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_689_i1_ in range(default__.abs(-192))]))))).cf12
            d_692_v17_: _dafny.MultiSet
            d_692_v17_ = _dafny.MultiSet([p2, p2])
            d_693_v18_: _dafny.MultiSet
            d_693_v18_ = d_692_v17_
            d_694_v19_: _dafny.Map
            d_694_v19_ = _dafny.Map({d_693_v18_: p1})
            d_695_v20_: _dafny.Seq
            d_695_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({default__.fm14(globalState): (self).f22}), d_694_v19_, d_694_v19_, d_694_v19_, d_694_v19_])
            d_696_v21_: _dafny.Map
            d_696_v21_ = _dafny.Map({False: (d_695_v20_) + (d_695_v20_)})
            d_696_v21_ = (d_696_v21_).set(d_688_cf12_, (d_695_v20_) + (d_695_v20_))
            d_697_v22_: _dafny.Map
            d_697_v22_ = _dafny.Map({(self).f22: p0})
            d_698_v23_: _dafny.Map
            d_698_v23_ = _dafny.Map({d_687_cf13_: p1})
            d_699_v24_: _dafny.Map
            d_699_v24_ = _dafny.Map({p2: p2})
            d_700_v25_: _dafny.Seq
            d_700_v25_ = _dafny.SeqWithoutIsStrInference([(((d_698_v23_).set(p1, len(d_699_v24_))).set(p3, d_687_cf13_)).set(d_687_cf13_, p1)])
            d_701_v26_: _dafny.Seq
            d_701_v26_ = _dafny.SeqWithoutIsStrInference([p3, len(d_700_v25_)])
            d_702_v27_: _dafny.Seq
            d_702_v27_ = _dafny.SeqWithoutIsStrInference([656, p1, len((_dafny.SeqWithoutIsStrInference([d_687_cf13_, len(d_697_v22_)])) + (d_701_v26_)), (self).f22, (0) - ((self).f22)])
            d_703_v28_: _dafny.Seq
            d_703_v28_ = _dafny.SeqWithoutIsStrInference([d_702_v27_])
            d_702_v27_ = ((d_703_v28_)[default__.safeIndex(d_687_cf13_, len(d_703_v28_))]) + (_dafny.SeqWithoutIsStrInference([p3]))
        elif source10_.is_DC3:
            d_704___mcc_h6_ = source10_.cf7
            d_705_cf7_ = d_704___mcc_h6_
            d_706_v29_: D0
            d_706_v29_ = D0_DC0(p2)
            d_707_v30_: T1
            nw114_ = C1()
            nw114_.ctor__((0) - (p1), d_706_v29_)
            d_707_v30_ = nw114_
            d_708_v31_: _dafny.MultiSet
            d_708_v31_ = _dafny.MultiSet([d_707_v30_])
            (globalState).f17 = ((d_708_v31_)[d_707_v30_] if (d_707_v30_) in (d_708_v31_) else p1)
            d_709_v32_: _dafny.Set
            d_709_v32_ = _dafny.Set({_dafny.CodePoint('s')})
            index101_ = default__.safeIndex(301, (d_659_v0_).length(0))
            (d_659_v0_)[index101_] = ((d_709_v32_).intersection(d_709_v32_)).ispropersubset(d_709_v32_)
            (globalState).f6 = p2
            d_710_v33_: C1
            nw115_ = C1()
            nw115_.ctor__(p1, d_706_v29_)
            d_710_v33_ = nw115_
        elif True:
            d_711___mcc_h7_ = source10_.cf14
            d_712_cf14_ = d_711___mcc_h7_
            (globalState).f18 = (p1 if ((d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))]) == ((d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))]) else (self).f22)
            d_713_v34_: _dafny.Array
            nw116_ = _dafny.Array(_dafny.MultiSet({}), 22)
            d_713_v34_ = nw116_
            d_714_v35_: _dafny.MultiSet
            d_714_v35_ = _dafny.MultiSet([(d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))], (d_674_v14_).fm13((self).f22, globalState)])
            d_715_v36_: _dafny.MultiSet
            d_715_v36_ = d_714_v35_
            index102_ = default__.safeIndex(344, (d_713_v34_).length(0))
            (d_713_v34_)[index102_] = d_715_v36_
            index103_ = default__.safeIndex(344, (d_713_v34_).length(0))
            (d_713_v34_)[index103_] = d_715_v36_
            d_716_v37_: _dafny.Array
            def lambda18_(d_717_v1_):
                def lambda19_(d_718_i2_):
                    return (d_717_v1_) + (d_717_v1_)

                return lambda19_

            init9_ = lambda18_(d_660_v1_)
            nw117_ = _dafny.Array(None, 24)
            for i0_9_ in range(nw117_.length(0)):
                nw117_[i0_9_] = init9_(i0_9_)
            d_716_v37_ = nw117_
            index104_ = default__.safeIndex(521, (d_716_v37_).length(0))
            (d_716_v37_)[index104_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drbyfkijt"))
            index105_ = default__.safeIndex(521, (d_716_v37_).length(0))
            (d_716_v37_)[index105_] = ((d_660_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brsiadvol")))) + (d_660_v1_)
            d_719_v38_: _dafny.Set
            d_719_v38_ = _dafny.Set({p1})
            d_720_v39_: _dafny.Map
            d_720_v39_ = _dafny.Map({(self).f22: d_719_v38_})
            d_721_v40_: _dafny.Set
            d_721_v40_ = _dafny.Set({d_719_v38_, d_719_v38_, d_719_v38_, d_719_v38_, ((d_720_v39_)[len((d_716_v37_)[default__.safeIndex(521, (d_716_v37_).length(0))])] if (len((d_716_v37_)[default__.safeIndex(521, (d_716_v37_).length(0))])) in (d_720_v39_) else d_719_v38_)})
            index106_ = default__.safeIndex(301, (d_659_v0_).length(0))
            (d_659_v0_)[index106_] = (p0 if (d_721_v40_).ispropersubset(d_721_v40_) else (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))])
        d_722_v41_: _dafny.MultiSet
        d_722_v41_ = _dafny.MultiSet([p0])
        r0 = (_dafny.MultiSet([p0])).isdisjoint(d_722_v41_)
        d_723_v42_: D0
        d_723_v42_ = D0_DC0(p0)
        d_724_v43_: _dafny.Set
        d_724_v43_ = _dafny.Set({p2, default__.fm2(d_662_v3_, p1, globalState)})
        d_725_v44_: D0
        d_725_v44_ = D0_DC1(d_661_v2_, p3, d_724_v43_, p2, (d_659_v0_)[default__.safeIndex(301, (d_659_v0_).length(0))])
        d_726_v45_: _dafny.Map
        d_726_v45_ = _dafny.Map({(0) - ((self).f22): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))})
        nw118_ = _dafny.Array(None, 16)
        nw118_[int(0)] = 403
        nw118_[int(1)] = p1
        nw118_[int(2)] = ((self).f22 if (d_723_v42_).cf0 else p1)
        nw118_[int(3)] = (0) - (len((d_660_v1_) + (d_660_v1_)))
        nw118_[int(4)] = len(d_660_v1_)
        nw118_[int(5)] = (p3) - (p1)
        nw118_[int(6)] = -825
        nw118_[int(7)] = ((0) - ((0) - ((self).f22))) * (p3)
        nw118_[int(8)] = default__.safeDivisionInt(p1, (d_722_v41_).cardinality)
        nw118_[int(9)] = 405
        nw118_[int(10)] = len(default__.fm19(default__.fm0(p3, d_725_v44_, globalState), d_661_v2_, globalState))
        nw118_[int(11)] = p3
        nw118_[int(12)] = ((self).f22) + (len(((d_726_v45_)[p1] if (p1) in (d_726_v45_) else d_660_v1_)))
        nw118_[int(13)] = p3
        nw118_[int(14)] = (d_722_v41_).cardinality
        nw118_[int(15)] = p3
        r1 = nw118_
        r2 = default__.fm20(p2, d_725_v44_, globalState)
        return r0, r1, r2

    @property
    def f22(self):
        return self._f22

class C3(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f20):
        (self)._f20 = f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        hi3_ = p1
        for d_727_i0_ in range(p1, hi3_):
            d_728_v0_: _dafny.Array
            def lambda20_(d_729_p0_):
                def lambda21_(d_730_i1_):
                    return d_729_p0_

                return lambda21_

            init10_ = lambda20_(p0)
            nw119_ = _dafny.Array(None, 27)
            for i0_10_ in range(nw119_.length(0)):
                nw119_[i0_10_] = init10_(i0_10_)
            d_728_v0_ = nw119_
            d_731_v1_: _dafny.Set
            d_731_v1_ = _dafny.Set({p0})
            index107_ = default__.safeIndex(199, (d_728_v0_).length(0))
            (d_728_v0_)[index107_] = (d_731_v1_).issubset(d_731_v1_)
            d_732_v2_: _dafny.Map
            d_732_v2_ = _dafny.Map({(0) - (d_727_i0_): (p1) <= (d_727_i0_)})
            d_733_v3_: _dafny.MultiSet
            d_733_v3_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
            d_734_v4_: _dafny.Array
            nw120_ = _dafny.Array(None, 10)
            nw120_[int(0)] = d_733_v3_
            nw120_[int(1)] = d_733_v3_
            nw120_[int(2)] = d_733_v3_
            nw120_[int(3)] = d_733_v3_
            nw120_[int(4)] = _dafny.MultiSet([p0, True])
            nw120_[int(5)] = d_733_v3_
            nw120_[int(6)] = d_733_v3_
            nw120_[int(7)] = d_733_v3_
            nw120_[int(8)] = d_733_v3_
            nw120_[int(9)] = d_733_v3_
            d_734_v4_ = nw120_
            d_735_v5_: _dafny.Map
            d_735_v5_ = _dafny.Map({d_734_v4_: p0})
            index108_ = default__.safeIndex(199, (d_728_v0_).length(0))
            (d_728_v0_)[index108_] = ((d_732_v2_)[d_727_i0_] if (d_727_i0_) in (d_732_v2_) else ((d_735_v5_)[d_734_v4_] if (d_734_v4_) in (d_735_v5_) else False))
            d_736_v6_: str
            d_736_v6_ = _dafny.CodePoint('o')
            d_737_v7_: _dafny.Seq
            d_737_v7_ = _dafny.SeqWithoutIsStrInference([d_736_v6_])
            d_738_v8_: D10
            d_738_v8_ = D10_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuile")), p1, not((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]), True, ((d_737_v7_).set(default__.safeIndex((d_733_v3_).cardinality, len(d_737_v7_)), d_736_v6_)) <= ((d_737_v7_).set(default__.safeIndex(d_727_i0_, len(d_737_v7_)), d_736_v6_)))
            source11_ = d_738_v8_
            if source11_.is_DC27:
                d_739___mcc_h0_ = source11_.cf50
                d_740___mcc_h1_ = source11_.cf51
                d_741_cf51_ = d_740___mcc_h1_
                d_742_cf50_ = d_739___mcc_h0_
                index109_ = default__.safeIndex(199, (d_728_v0_).length(0))
                (d_728_v0_)[index109_] = not((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))])
                d_743_v9_: _dafny.Array
                nw121_ = _dafny.Array(int(0), 24)
                d_743_v9_ = nw121_
                index110_ = default__.safeIndex(145, (d_743_v9_).length(0))
                (d_743_v9_)[index110_] = d_727_i0_
                d_744_v10_: _dafny.Seq
                d_744_v10_ = _dafny.SeqWithoutIsStrInference([(d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]])
                d_745_v11_: _dafny.Map
                d_745_v11_ = _dafny.Map({(d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]: len(d_744_v10_)})
                index111_ = default__.safeIndex(145, (d_743_v9_).length(0))
                rhs69_ = ((d_737_v7_) + (d_737_v7_)) + (default__.fm46(155, (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))], (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))], globalState))
                rhs70_ = ((d_745_v11_)[(d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]] if ((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]) in (d_745_v11_) else d_727_i0_)
                lhs62_ = globalState
                lhs63_ = d_743_v9_
                lhs64_ = default__.safeIndex(145, (d_743_v9_).length(0))
                lhs62_.f15 = rhs69_
                lhs63_[lhs64_] = rhs70_
                index112_ = default__.safeIndex(145, (d_743_v9_).length(0))
                (d_743_v9_)[index112_] = len(_dafny.SeqWithoutIsStrInference([d_736_v6_ for d_746_i2_ in range(default__.abs(411))]))
                (globalState).f5 = p0
            elif source11_.is_DC28:
                d_747___mcc_h2_ = source11_.cf52
                d_748___mcc_h3_ = source11_.cf53
                d_749___mcc_h4_ = source11_.cf54
                d_750___mcc_h5_ = source11_.cf55
                d_751___mcc_h6_ = source11_.cf56
                d_752_cf56_ = d_751___mcc_h6_
                d_753_cf55_ = d_750___mcc_h5_
                d_754_cf54_ = d_749___mcc_h4_
                d_755_cf53_ = d_748___mcc_h3_
                d_756_cf52_ = d_747___mcc_h2_
                d_757_v12_: _dafny.Array
                nw122_ = _dafny.Array(D4.default()(), 25)
                d_757_v12_ = nw122_
                d_757_v12_ = d_757_v12_
                (globalState).f4 = (d_728_v0_ if d_754_cf54_ else d_728_v0_)
                (globalState).f6 = p0
                d_758_v13_: _dafny.Array
                nw123_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_758_v13_ = nw123_
                d_758_v13_ = (d_758_v13_ if d_752_cf56_ else d_758_v13_)
            elif source11_.is_DC26:
                d_759___mcc_h7_ = source11_.cf49
                d_760_cf49_ = d_759___mcc_h7_
                d_761_v14_: _dafny.Seq
                d_761_v14_ = _dafny.SeqWithoutIsStrInference([False, p0])
                d_762_v15_: _dafny.Set
                d_762_v15_ = _dafny.Set({d_761_v14_, (_dafny.SeqWithoutIsStrInference([not(p0)])) + (d_761_v14_), ((d_761_v14_).set(default__.safeIndex(p1, len(d_761_v14_)), p0)) + (_dafny.SeqWithoutIsStrInference([(d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))], p0]))})
                rhs71_ = (d_762_v15_).intersection(d_762_v15_)
                rhs72_ = p0
                rhs73_ = (d_727_i0_) - (p1)
                rhs74_ = (p1 if ((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))] if (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))] else (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]) else d_727_i0_)
                lhs65_ = globalState
                lhs66_ = globalState
                lhs67_ = globalState
                d_762_v15_ = rhs71_
                lhs65_.f6 = rhs72_
                lhs66_.f18 = rhs73_
                lhs67_.f18 = rhs74_
                d_763_v16_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_763_v16_ = nw124_
                index113_ = default__.safeIndex(706, (d_763_v16_).length(0))
                (d_763_v16_)[index113_] = d_737_v7_
                d_764_v17_: _dafny.MultiSet
                d_764_v17_ = _dafny.MultiSet([(0) - ((0) - (default__.safeDivisionInt(p1, d_727_i0_))), 307])
                index114_ = default__.safeIndex(706, (d_763_v16_).length(0))
                rhs75_ = ((d_737_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nydoft")))) + (d_737_v7_)
                rhs76_ = (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]
                rhs77_ = (d_764_v17_).intersection(d_764_v17_)
                lhs68_ = d_763_v16_
                lhs69_ = default__.safeIndex(706, (d_763_v16_).length(0))
                lhs70_ = globalState
                lhs68_[lhs69_] = rhs75_
                lhs70_.f5 = rhs76_
                d_764_v17_ = rhs77_
                d_765_v18_: _dafny.Array
                nw125_ = _dafny.Array(int(0), 16)
                d_765_v18_ = nw125_
                index115_ = default__.safeIndex(390, (d_765_v18_).length(0))
                (d_765_v18_)[index115_] = p1
                d_766_v19_: _dafny.MultiSet
                d_766_v19_ = d_733_v3_
                d_767_v20_: _dafny.Map
                d_767_v20_ = _dafny.Map({d_766_v19_: d_727_i0_})
                index116_ = default__.safeIndex(390, (d_765_v18_).length(0))
                (d_765_v18_)[index116_] = len(((d_760_cf49_) + (d_760_cf49_)).set(default__.safeIndex(default__.safeModuloInt(((d_767_v20_)[d_766_v19_] if (d_766_v19_) in (d_767_v20_) else p1), d_727_i0_), len((d_760_cf49_) + (d_760_cf49_))), (0) - (d_727_i0_)))
                d_768_v21_: _dafny.Map
                d_768_v21_ = _dafny.Map({not (p0) or ((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]): p0})
                d_769_v22_: _dafny.MultiSet
                d_769_v22_ = _dafny.MultiSet([_dafny.CodePoint('n')])
                d_770_v23_: _dafny.Map
                d_770_v23_ = _dafny.Map({d_765_v18_: d_765_v18_})
                d_771_v24_: _dafny.Map
                d_771_v24_ = _dafny.Map({len(d_770_v23_): (d_765_v18_)[default__.safeIndex(390, (d_765_v18_).length(0))]})
                d_768_v21_ = (d_768_v21_).set(default__.fm2(d_769_v22_, p1, globalState), (len(d_771_v24_)) != (len(d_760_cf49_)))
            elif True:
                d_772___mcc_h8_ = source11_.cf57
                d_773_cf57_ = d_772___mcc_h8_
                index117_ = default__.safeIndex(199, (d_728_v0_).length(0))
                (d_728_v0_)[index117_] = ((d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))]) or (p0)
                d_732_v2_ = (d_732_v2_).set(p1, (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))])
                d_774_v25_: _dafny.MultiSet
                d_774_v25_ = d_733_v3_
                d_774_v25_ = d_733_v3_
                (globalState).f4 = d_728_v0_
            d_775_v26_: _dafny.MultiSet
            d_775_v26_ = _dafny.MultiSet([(0) - (p1)])
            d_776_v27_: _dafny.Map
            d_776_v27_ = _dafny.Map({p0: (_dafny.MultiSet([d_727_i0_]) if (d_728_v0_)[default__.safeIndex(199, (d_728_v0_).length(0))] else d_775_v26_)})
            (globalState).f17 = (((d_776_v27_)[p0] if (p0) in (d_776_v27_) else _dafny.MultiSet([d_727_i0_]))).cardinality
            (globalState).f6 = p0
        d_777_v28_: _dafny.Set
        d_777_v28_ = _dafny.Set({p0})
        (globalState).f6 = (d_777_v28_).issubset((d_777_v28_) | (d_777_v28_))
        if not(p0):
            (globalState).f17 = (p1) + ((-844) + (p1))
            if not((True) or (p0)):
                d_778_v29_: _dafny.Array
                nw126_ = _dafny.Array(int(0), 2)
                d_778_v29_ = nw126_
                d_779_v30_: _dafny.Map
                d_779_v30_ = _dafny.Map({d_778_v29_: p1})
                d_780_v31_: _dafny.MultiSet
                d_780_v31_ = _dafny.MultiSet([p1, 202, p1])
                d_781_v32_: _dafny.Seq
                d_781_v32_ = _dafny.SeqWithoutIsStrInference([p1, ((d_780_v31_)[p1] if (p1) in (d_780_v31_) else p1), p1])
                d_782_v33_: _dafny.Seq
                d_782_v33_ = _dafny.SeqWithoutIsStrInference([p0])
                d_783_v34_: _dafny.Array
                nw127_ = _dafny.Array(None, 4)
                nw127_[int(0)] = _dafny.SeqWithoutIsStrInference([len(d_779_v30_), -699, p1])
                nw127_[int(1)] = (d_781_v32_).set(default__.safeIndex(len(d_782_v33_), len(d_781_v32_)), p1)
                nw127_[int(2)] = d_781_v32_
                nw127_[int(3)] = d_781_v32_
                d_783_v34_ = nw127_
                index118_ = default__.safeIndex(607, (d_783_v34_).length(0))
                (d_783_v34_)[index118_] = d_781_v32_
                d_784_v35_: str
                d_784_v35_ = _dafny.CodePoint('x')
                d_785_v36_: D0
                d_785_v36_ = D0_DC1(d_784_v35_, p1, d_777_v28_, not(p0), p0)
                d_786_v37_: _dafny.MultiSet
                d_786_v37_ = _dafny.MultiSet([p0, p0, p0, default__.fm2(_dafny.MultiSet([d_784_v35_, d_784_v35_]), len(_dafny.Set({p1, len(d_781_v32_)})), globalState)])
                index119_ = default__.safeIndex(607, (d_783_v34_).length(0))
                rhs78_ = p0
                rhs79_ = p0
                rhs80_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p1, d_785_v36_, globalState), p1, p1, (d_786_v37_).cardinality])
                rhs81_ = (0) - (p1)
                lhs71_ = globalState
                lhs72_ = globalState
                lhs73_ = d_783_v34_
                lhs74_ = default__.safeIndex(607, (d_783_v34_).length(0))
                lhs75_ = globalState
                lhs71_.f5 = rhs78_
                lhs72_.f6 = rhs79_
                lhs73_[lhs74_] = rhs80_
                lhs75_.f18 = rhs81_
                d_787_v38_: _dafny.Set
                d_787_v38_ = _dafny.Set({(d_783_v34_)[default__.safeIndex(607, (d_783_v34_).length(0))], d_781_v32_})
                d_788_v39_: _dafny.Seq
                d_788_v39_ = _dafny.SeqWithoutIsStrInference([default__.fm47(p1, p0, p0, globalState), (d_783_v34_)[default__.safeIndex(607, (d_783_v34_).length(0))], _dafny.SeqWithoutIsStrInference([p1])])
                d_789_v41_: _dafny.Map
                d_789_v41_ = _dafny.Map({p0: p0})
                d_790_v42_: _dafny.Set
                d_790_v42_ = _dafny.Set({d_789_v41_})
                nw128_ = _dafny.Array(None, 21)
                nw128_[int(0)] = default__.safeModuloInt(p1, p1)
                def iife122_():
                    coll36_ = _dafny.Set()
                    compr_36_: _dafny.Seq
                    for compr_36_ in (d_788_v39_).Elements:
                        d_791_v40_: _dafny.Seq = compr_36_
                        if (d_791_v40_) in (d_788_v39_):
                            coll36_ = coll36_.union(_dafny.Set([d_791_v40_]))
                    return _dafny.Set(coll36_)
                nw128_[int(1)] = len((d_787_v38_).intersection(iife122_()
                ))
                nw128_[int(2)] = default__.fm0(p1, default__.fm48(len(_dafny.SeqWithoutIsStrInference([d_784_v35_ for d_792_i3_ in range(default__.abs(406))])), p1, p1, globalState), globalState)
                nw128_[int(3)] = default__.safeModuloInt(p1, p1)
                nw128_[int(4)] = p1
                nw128_[int(5)] = p1
                nw128_[int(6)] = p1
                nw128_[int(7)] = (0) - (p1)
                nw128_[int(8)] = p1
                nw128_[int(9)] = len(d_790_v42_)
                nw128_[int(10)] = (p1) * (p1)
                nw128_[int(11)] = -75
                nw128_[int(12)] = p1
                nw128_[int(13)] = p1
                nw128_[int(14)] = 727
                nw128_[int(15)] = default__.fm0(p1, default__.fm48(p1, p1, -287, globalState), globalState)
                nw128_[int(16)] = len(_dafny.SeqWithoutIsStrInference([p0, p0, p0, p0]))
                nw128_[int(17)] = p1
                nw128_[int(18)] = p1
                nw128_[int(19)] = p1
                nw128_[int(20)] = (p1) - ((0) - (p1))
                d_778_v29_ = nw128_
                d_793_v43_: _dafny.Map
                d_793_v43_ = _dafny.Map({p1: p1})
                (globalState).f18 = len(d_793_v43_)
                d_794_v44_: _dafny.Map
                d_794_v44_ = _dafny.Map({p1: p0})
                d_795_v45_: _dafny.Seq
                d_795_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqvqoqi"))
                d_794_v44_ = (d_794_v44_).set(p1, (d_795_v45_) != (d_795_v45_))
                (globalState).f18 = default__.safeDivisionInt(p1, (0) - (p1))
            elif True:
                d_796_v46_: _dafny.Seq
                d_796_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqkcsf"))
                d_797_v47_: _dafny.Set
                d_797_v47_ = _dafny.Set({p1, p1, p1})
                d_798_v48_: str
                d_798_v48_ = _dafny.CodePoint('s')
                d_799_v49_: _dafny.Map
                d_799_v49_ = _dafny.Map({default__.fm49(d_796_v46_, globalState): d_798_v48_})
                d_800_v50_: _dafny.Map
                d_800_v50_ = _dafny.Map({p0: p1})
                d_801_v51_: D0
                d_801_v51_ = D0_DC1(d_798_v48_, p1, d_777_v28_, not(p0), p0)
                d_802_v52_: _dafny.Seq
                d_802_v52_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, p0, p0]))).cardinality])
                d_803_v53_: _dafny.Map
                d_803_v53_ = _dafny.Map({len(d_802_v52_): p1})
                d_804_v54_: _dafny.Seq
                d_804_v54_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
                d_805_v55_: _dafny.MultiSet
                d_805_v55_ = _dafny.MultiSet([p1])
                pat_let_tv31_ = p1
                d_806_v56_: _dafny.Array
                nw129_ = _dafny.Array(None, 19)
                nw129_[int(0)] = (p1) - (len((d_796_v46_).set(default__.safeIndex(423, len(d_796_v46_)), default__.fm1(globalState))))
                nw129_[int(1)] = len(d_797_v47_)
                nw129_[int(2)] = len(d_799_v49_)
                nw129_[int(3)] = -884
                nw129_[int(4)] = (585 if p0 else len(d_800_v50_))
                nw129_[int(5)] = 658
                def iife123_(_pat_let43_0):
                    def iife124_(d_807_dt__update__tmp_h1_):
                        def iife125_(_pat_let44_0):
                            def iife126_(d_808_dt__update_hcf2_h0_):
                                return D0_DC1((d_807_dt__update__tmp_h1_).cf1, d_808_dt__update_hcf2_h0_, (d_807_dt__update__tmp_h1_).cf3, (d_807_dt__update__tmp_h1_).cf4, (d_807_dt__update__tmp_h1_).cf5)
                            return iife126_(_pat_let44_0)
                        return iife125_(pat_let_tv31_)
                    return iife124_(_pat_let43_0)
                nw129_[int(6)] = default__.fm0(p1, iife123_(d_801_v51_), globalState)
                nw129_[int(7)] = len((d_796_v46_) + (_dafny.SeqWithoutIsStrInference([d_798_v48_ for d_809_i4_ in range(default__.abs(402))])))
                nw129_[int(8)] = ((d_803_v53_)[p1] if (p1) in (d_803_v53_) else 723)
                nw129_[int(9)] = ((d_800_v50_)[p0] if (p0) in (d_800_v50_) else p1)
                nw129_[int(10)] = len(((d_804_v54_) + (d_804_v54_)).set(default__.safeIndex(p1, len((d_804_v54_) + (d_804_v54_))), p0))
                nw129_[int(11)] = (0) - (p1)
                nw129_[int(12)] = p1
                nw129_[int(13)] = 353
                nw129_[int(14)] = p1
                nw129_[int(15)] = (p1) - (594)
                nw129_[int(16)] = (d_805_v55_).cardinality
                nw129_[int(17)] = default__.safeModuloInt(len(d_803_v53_), p1)
                nw129_[int(18)] = p1
                d_806_v56_ = nw129_
                index120_ = default__.safeIndex(357, (d_806_v56_).length(0))
                (d_806_v56_)[index120_] = 657
                index121_ = default__.safeIndex(448, (d_806_v56_).length(0))
                (d_806_v56_)[index121_] = (len(d_796_v46_)) + (p1)
                d_810_v57_: _dafny.Array
                nw130_ = _dafny.Array(False, 19)
                d_810_v57_ = nw130_
                d_811_v58_: _dafny.Map
                d_811_v58_ = _dafny.Map({p1: p0})
                d_812_v59_: D5
                d_812_v59_ = D5_DC13((d_811_v58_).set(p1, not(p0)))
                pat_let_tv32_ = d_811_v58_
                d_813_v60_: _dafny.Seq
                def iife127_(_pat_let45_0):
                    def iife128_(d_814_dt__update__tmp_h2_):
                        def iife129_(_pat_let46_0):
                            def iife130_(d_815_dt__update_hcf24_h0_):
                                return D5_DC13(d_815_dt__update_hcf24_h0_)
                            return iife130_(_pat_let46_0)
                        return iife129_((D5_DC13(pat_let_tv32_)).cf24)
                    return iife128_(_pat_let45_0)
                d_813_v60_ = _dafny.SeqWithoutIsStrInference([iife127_(d_812_v59_)])
                d_816_v61_: _dafny.Map
                d_816_v61_ = _dafny.Map({p1: (d_813_v60_ if p0 else d_813_v60_)})
                d_817_v62_: _dafny.MultiSet
                d_817_v62_ = _dafny.MultiSet([d_798_v48_, d_798_v48_, d_798_v48_])
                index122_ = default__.safeIndex(357, (d_806_v56_).length(0))
                index123_ = default__.safeIndex(448, (d_806_v56_).length(0))
                rhs82_ = d_810_v57_
                rhs83_ = default__.safeModuloInt((0) - ((739) - (p1)), 476)
                rhs84_ = len(d_816_v61_)
                rhs85_ = not(((len(d_800_v50_)) - (p1)) >= ((d_817_v62_).cardinality))
                rhs86_ = (0) - (default__.fm0(p1, d_801_v51_, globalState))
                lhs76_ = globalState
                lhs77_ = globalState
                lhs78_ = d_806_v56_
                lhs79_ = default__.safeIndex(357, (d_806_v56_).length(0))
                lhs80_ = globalState
                lhs81_ = d_806_v56_
                lhs82_ = default__.safeIndex(448, (d_806_v56_).length(0))
                lhs76_.f4 = rhs82_
                lhs77_.f18 = rhs83_
                lhs78_[lhs79_] = rhs84_
                lhs80_.f6 = rhs85_
                lhs81_[lhs82_] = rhs86_
                (globalState).f6 = p0
                d_818_v63_: _dafny.Array
                def lambda22_(d_819_v48_):
                    def lambda23_(d_820_i5_):
                        return d_819_v48_

                    return lambda23_

                init11_ = lambda22_(d_798_v48_)
                nw131_ = _dafny.Array(None, 13)
                for i0_11_ in range(nw131_.length(0)):
                    nw131_[i0_11_] = init11_(i0_11_)
                d_818_v63_ = nw131_
                d_821_v64_: _dafny.MultiSet
                d_821_v64_ = _dafny.MultiSet([d_818_v63_, d_818_v63_, d_818_v63_])
                (globalState).f6 = not((d_821_v64_).ispropersubset(d_821_v64_))
                d_822_v65_: _dafny.MultiSet
                d_822_v65_ = _dafny.MultiSet([p0])
                d_823_v66_: _dafny.Map
                d_823_v66_ = _dafny.Map({d_818_v63_: (0) - (((d_822_v65_).set(p0, default__.abs(46))).cardinality)})
                d_823_v66_ = (d_823_v66_).set(d_818_v63_, p1)
                nw132_ = _dafny.Array(None, 11)
                nw132_[int(0)] = default__.fm0((0) - ((d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]), d_801_v51_, globalState)
                nw132_[int(1)] = (d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]
                nw132_[int(2)] = (d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]
                nw132_[int(3)] = len(d_804_v54_)
                nw132_[int(4)] = p1
                nw132_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nkihb")))
                nw132_[int(6)] = (d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]
                nw132_[int(7)] = len(d_797_v47_)
                nw132_[int(8)] = ((d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]) * ((d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))])
                nw132_[int(9)] = (d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))]
                nw132_[int(10)] = (len(d_796_v46_)) - ((d_806_v56_)[default__.safeIndex(357, (d_806_v56_).length(0))])
                d_806_v56_ = nw132_
            if not(p0):
                d_824_v67_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_824_v67_ = nw133_
                d_825_v68_: _dafny.Array
                nw134_ = _dafny.Array(False, 19)
                d_825_v68_ = nw134_
                index124_ = default__.safeIndex(252, (d_824_v67_).length(0))
                (d_824_v67_)[index124_] = d_825_v68_
                d_826_v69_: _dafny.Map
                d_826_v69_ = _dafny.Map({p1: p0})
                d_827_v70_: _dafny.Set
                d_827_v70_ = _dafny.Set({p1, p1, default__.safeDivisionInt(len(_dafny.Set({p0})), p1), p1, (len(d_826_v69_)) - (len(_dafny.Set({p1, p1})))})
                index125_ = default__.safeIndex(252, (d_824_v67_).length(0))
                rhs87_ = p1
                rhs88_ = d_825_v68_
                rhs89_ = len(d_827_v70_)
                lhs83_ = globalState
                lhs84_ = d_824_v67_
                lhs85_ = default__.safeIndex(252, (d_824_v67_).length(0))
                lhs86_ = globalState
                lhs83_.f17 = rhs87_
                lhs84_[lhs85_] = rhs88_
                lhs86_.f18 = rhs89_
                d_828_v71_: _dafny.Array
                nw135_ = _dafny.Array(int(0), 20)
                d_828_v71_ = nw135_
                d_828_v71_ = (d_828_v71_ if ((d_826_v69_)[p1] if (p1) in (d_826_v69_) else p0) else d_828_v71_)
                d_829_v72_: _dafny.Map
                d_829_v72_ = _dafny.Map({True: p1})
                d_830_v73_: D12
                d_830_v73_ = D12_DC32(_dafny.Set({d_829_v72_, d_829_v72_}))
                (globalState).f6 = ((d_830_v73_).cf61).ispropersubset(_dafny.Set({_dafny.Map({True: p1}), d_829_v72_, d_829_v72_}))
                d_831_v74_: str
                d_831_v74_ = _dafny.CodePoint('f')
                r2 = d_831_v74_
                d_832_v75_: _dafny.MultiSet
                d_832_v75_ = default__.fm51(globalState)
                d_833_v76_: C0
                nw136_ = C0()
                nw136_.ctor__(d_825_v68_, default__.fm50(True, not(p0), d_832_v75_, globalState))
                d_833_v76_ = nw136_
            elif True:
                r2 = default__.fm1(globalState)
                d_834_v77_: _dafny.Seq
                d_834_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnuxmwxl"))
                (globalState).f16 = (d_834_v77_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hql")))
                (globalState).f18 = (0) - (p1)
                d_835_v78_: _dafny.Array
                nw137_ = _dafny.Array(False, 1)
                d_835_v78_ = nw137_
                d_836_v79_: _dafny.Seq
                d_836_v79_ = _dafny.SeqWithoutIsStrInference([p0])
                index126_ = default__.safeIndex(107, (d_835_v78_).length(0))
                (d_835_v78_)[index126_] = (d_836_v79_)[default__.safeIndex(p1, len(d_836_v79_))]
                d_837_v80_: _dafny.Array
                nw138_ = _dafny.Array(int(0), 12)
                d_837_v80_ = nw138_
                index127_ = default__.safeIndex(138, (d_837_v80_).length(0))
                (d_837_v80_)[index127_] = (p1) + (p1)
                d_838_v82_: str
                d_838_v82_ = _dafny.CodePoint('s')
                d_839_v83_: _dafny.MultiSet
                d_839_v83_ = _dafny.MultiSet([-787])
                d_840_v84_: D0
                d_840_v84_ = D0_DC1(d_838_v82_, (d_839_v83_).cardinality, d_777_v28_, p0, p0)
                index128_ = default__.safeIndex(107, (d_835_v78_).length(0))
                index129_ = default__.safeIndex(138, (d_837_v80_).length(0))
                def iife131_():
                    coll37_ = _dafny.Map()
                    compr_37_: str
                    for compr_37_ in (d_834_v77_).Elements:
                        d_841_v81_: str = compr_37_
                        if (d_841_v81_) in (d_834_v77_):
                            coll37_[d_841_v81_] = p0
                    return _dafny.Map(coll37_)
                rhs90_ = False
                rhs91_ = p1
                rhs92_ = default__.fm46((0) - (len((iife131_()
                ) | (_dafny.Map({d_838_v82_: p0})))), (d_836_v79_) < (d_836_v79_), (d_777_v28_).isdisjoint(d_777_v28_), globalState)
                rhs93_ = p1
                rhs94_ = default__.fm0((p1) * (p1), d_840_v84_, globalState)
                lhs87_ = d_835_v78_
                lhs88_ = default__.safeIndex(107, (d_835_v78_).length(0))
                lhs89_ = globalState
                lhs90_ = globalState
                lhs91_ = d_837_v80_
                lhs92_ = default__.safeIndex(138, (d_837_v80_).length(0))
                lhs87_[lhs88_] = rhs90_
                lhs89_.f18 = rhs91_
                d_834_v77_ = rhs92_
                lhs90_.f17 = rhs93_
                lhs91_[lhs92_] = rhs94_
                d_842_v85_: _dafny.Map
                d_842_v85_ = _dafny.Map({p1: p1})
                d_843_v86_: _dafny.Seq
                d_843_v86_ = _dafny.SeqWithoutIsStrInference([default__.fm0(717, d_840_v84_, globalState)])
                d_844_v87_: _dafny.MultiSet
                d_844_v87_ = _dafny.MultiSet([d_837_v80_])
                r1 = (((d_842_v85_)[(d_837_v80_)[default__.safeIndex(138, (d_837_v80_).length(0))]] if ((d_837_v80_)[default__.safeIndex(138, (d_837_v80_).length(0))]) in (d_842_v85_) else (d_843_v86_)[default__.safeIndex(p1, len(d_843_v86_))])) - ((d_844_v87_).cardinality)
            d_845_v88_: _dafny.Map
            d_845_v88_ = _dafny.Map({p0: p0})
            d_846_v89_: C1
            nw139_ = C1()
            nw139_.ctor__((p1 if p0 else len(d_845_v88_)), D0_DC0(p0))
            d_846_v89_ = nw139_
            d_847_v90_: _dafny.Array
            nw140_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_847_v90_ = nw140_
            d_848_v91_: _dafny.Array
            nw141_ = _dafny.Array(int(0), 11)
            d_848_v91_ = nw141_
            index130_ = default__.safeIndex(714, (d_847_v90_).length(0))
            (d_847_v90_)[index130_] = d_848_v91_
            d_849_v92_: str
            d_849_v92_ = _dafny.CodePoint('c')
            d_850_v93_: _dafny.MultiSet
            d_850_v93_ = _dafny.MultiSet([d_849_v92_])
            d_851_v94_: D7
            d_851_v94_ = D7_DC21(723, p0, True, d_848_v91_, p0)
            index131_ = default__.safeIndex(714, (d_847_v90_).length(0))
            (d_847_v90_)[index131_] = (((d_851_v94_).cf41 if default__.fm2(d_850_v93_, p1, globalState) else d_848_v91_) if (p0 if p0 else p0) else d_848_v91_)
        elif True:
            d_852_v95_: _dafny.Seq
            d_852_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caabamdwi"))
            d_853_v96_: _dafny.Array
            def lambda24_(d_854_p0_):
                def lambda25_(d_855_i6_):
                    return d_854_p0_

                return lambda25_

            init12_ = lambda24_(p0)
            nw142_ = _dafny.Array(None, 12)
            for i0_12_ in range(nw142_.length(0)):
                nw142_[i0_12_] = init12_(i0_12_)
            d_853_v96_ = nw142_
            d_856_v97_: _dafny.Map
            d_856_v97_ = _dafny.Map({d_852_v95_: d_853_v96_})
            d_857_v98_: D11
            d_857_v98_ = D11_DC31(d_856_v97_, p1)
            if not((p1) >= ((d_857_v98_).cf60)):
                d_858_v99_: _dafny.Map
                d_858_v99_ = _dafny.Map({p0: p0})
                d_859_v100_: _dafny.MultiSet
                d_859_v100_ = _dafny.MultiSet([d_858_v99_, d_858_v99_])
                d_860_v101_: _dafny.Set
                d_860_v101_ = _dafny.Set({(0) - (((d_859_v100_)[d_858_v99_] if (d_858_v99_) in (d_859_v100_) else len(d_852_v95_)))})
                index132_ = default__.safeIndex(249, (d_853_v96_).length(0))
                (d_853_v96_)[index132_] = p0
                d_861_v102_: _dafny.Map
                d_861_v102_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybjxva"))})
                d_862_v103_: _dafny.Map
                d_862_v103_ = _dafny.Map({p1: p1})
                d_863_v104_: str
                d_863_v104_ = _dafny.CodePoint('s')
                d_864_v105_: _dafny.Map
                d_864_v105_ = _dafny.Map({p0: p1})
                d_865_v106_: D0
                d_865_v106_ = D0_DC1(d_863_v104_, len(d_864_v105_), default__.fm52(p1, d_852_v95_, True, p0, globalState), p0, p0)
                d_866_v107_: _dafny.Array
                nw143_ = _dafny.Array(None, 19)
                nw143_[int(0)] = len(((d_861_v102_)[p0] if (p0) in (d_861_v102_) else d_852_v95_))
                nw143_[int(1)] = p1
                nw143_[int(2)] = len(d_862_v103_)
                nw143_[int(3)] = p1
                nw143_[int(4)] = (0) - (p1)
                nw143_[int(5)] = p1
                nw143_[int(6)] = p1
                nw143_[int(7)] = p1
                nw143_[int(8)] = p1
                nw143_[int(9)] = 95
                nw143_[int(10)] = default__.fm0(p1, d_865_v106_, globalState)
                nw143_[int(11)] = p1
                nw143_[int(12)] = p1
                nw143_[int(13)] = p1
                nw143_[int(14)] = p1
                nw143_[int(15)] = p1
                nw143_[int(16)] = p1
                nw143_[int(17)] = p1
                nw143_[int(18)] = p1
                d_866_v107_ = nw143_
                d_867_v108_: D7
                d_867_v108_ = D7_DC21(p1, p0, p0, d_866_v107_, p0)
                d_868_v109_: _dafny.Map
                d_868_v109_ = _dafny.Map({(d_867_v108_).cf39: _dafny.Set({len(d_852_v95_)})})
                d_869_v110_: _dafny.Seq
                d_869_v110_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_870_v111_: _dafny.MultiSet
                d_870_v111_ = _dafny.MultiSet([d_866_v107_])
                index133_ = default__.safeIndex(249, (d_853_v96_).length(0))
                rhs95_ = (((d_868_v109_)[p0] if (p0) in (d_868_v109_) else d_860_v101_) if p0 else d_860_v101_)
                rhs96_ = not((d_869_v110_) == (d_869_v110_))
                rhs97_ = (d_870_v111_).cardinality
                rhs98_ = p0
                rhs99_ = p0
                lhs93_ = globalState
                lhs94_ = globalState
                lhs95_ = globalState
                lhs96_ = d_853_v96_
                lhs97_ = default__.safeIndex(249, (d_853_v96_).length(0))
                d_860_v101_ = rhs95_
                lhs93_.f5 = rhs96_
                lhs94_.f17 = rhs97_
                lhs95_.f5 = rhs98_
                lhs96_[lhs97_] = rhs99_
                d_871_v112_: D5
                d_871_v112_ = D5_DC14(p0)
                d_872_v113_: _dafny.MultiSet
                d_872_v113_ = _dafny.MultiSet([(d_871_v112_).cf25])
                d_873_v114_: _dafny.Array
                nw144_ = _dafny.Array(None, 11)
                nw144_[int(0)] = d_872_v113_
                nw144_[int(1)] = d_872_v113_
                nw144_[int(2)] = (d_872_v113_) | (d_872_v113_)
                nw144_[int(3)] = (default__.fm51(globalState)).intersection(_dafny.MultiSet([True]))
                nw144_[int(4)] = (d_872_v113_) - (d_872_v113_)
                nw144_[int(5)] = _dafny.MultiSet([p0])
                nw144_[int(6)] = d_872_v113_
                nw144_[int(7)] = d_872_v113_
                nw144_[int(8)] = d_872_v113_
                nw144_[int(9)] = d_872_v113_
                nw144_[int(10)] = (_dafny.MultiSet(d_869_v110_)).set(p0, default__.abs(p1))
                d_873_v114_ = nw144_
                index134_ = default__.safeIndex(127, (d_873_v114_).length(0))
                (d_873_v114_)[index134_] = d_872_v113_
                d_874_v115_: _dafny.Map
                d_874_v115_ = _dafny.Map({p1: (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]})
                d_875_v116_: _dafny.Map
                d_875_v116_ = _dafny.Map({d_874_v115_: (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]})
                d_876_v117_: _dafny.Array
                nw145_ = _dafny.Array(None, 14)
                nw145_[int(0)] = p0
                nw145_[int(1)] = p0
                nw145_[int(2)] = ((d_875_v116_)[_dafny.Map({p1: p0})] if (_dafny.Map({p1: p0})) in (d_875_v116_) else True)
                nw145_[int(3)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                nw145_[int(4)] = default__.fm2(_dafny.MultiSet([d_863_v104_]), p1, globalState)
                nw145_[int(5)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                nw145_[int(6)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                nw145_[int(7)] = p0
                nw145_[int(8)] = p0
                nw145_[int(9)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                nw145_[int(10)] = p0
                nw145_[int(11)] = p0
                nw145_[int(12)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                nw145_[int(13)] = (d_853_v96_)[default__.safeIndex(249, (d_853_v96_).length(0))]
                d_876_v117_ = nw145_
                d_877_v118_: _dafny.MultiSet
                d_877_v118_ = _dafny.MultiSet([d_876_v117_, d_876_v117_, d_876_v117_, d_876_v117_, d_876_v117_])
                index135_ = default__.safeIndex(127, (d_873_v114_).length(0))
                rhs100_ = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))).intersection(d_872_v113_)).intersection(d_872_v113_)
                rhs101_ = not((d_877_v118_).issubset((d_877_v118_) - (d_877_v118_)))
                lhs98_ = d_873_v114_
                lhs99_ = default__.safeIndex(127, (d_873_v114_).length(0))
                lhs100_ = globalState
                lhs98_[lhs99_] = rhs100_
                lhs100_.f6 = rhs101_
                d_852_v95_ = ((d_852_v95_ if (d_869_v110_)[default__.safeIndex(p1, len(d_869_v110_))] else d_852_v95_)) + (d_852_v95_)
                (globalState).f6 = (d_869_v110_)[default__.safeIndex((0) - (default__.safeModuloInt(p1, (0) - (p1))), len(d_869_v110_))]
                (globalState).f6 = p0
            elif True:
                d_878_v119_: _dafny.Map
                d_878_v119_ = _dafny.Map({p0: p1})
                (globalState).f17 = ((d_878_v119_)[p0] if (p0) in (d_878_v119_) else p1)
                r2 = _dafny.CodePoint('o')
                index136_ = default__.safeIndex(143, (d_853_v96_).length(0))
                (d_853_v96_)[index136_] = p0
                index137_ = default__.safeIndex(143, (d_853_v96_).length(0))
                (d_853_v96_)[index137_] = p0
                (globalState).f16 = ((d_852_v95_) + (d_852_v95_)) + (d_852_v95_)
                d_879_v120_: _dafny.Array
                def lambda26_(d_880_p1_):
                    def lambda27_(d_881_i7_):
                        return default__.safeModuloInt(d_881_i7_, d_880_p1_)

                    return lambda27_

                init13_ = lambda26_(p1)
                nw146_ = _dafny.Array(None, 16)
                for i0_13_ in range(nw146_.length(0)):
                    nw146_[i0_13_] = init13_(i0_13_)
                d_879_v120_ = nw146_
                d_882_v121_: T1
                nw147_ = C1()
                nw147_.ctor__(p1, (self).f20)
                d_882_v121_ = nw147_
                d_883_v122_: _dafny.Array
                nw148_ = _dafny.Array(False, 18)
                d_883_v122_ = nw148_
                d_884_v123_: _dafny.MultiSet
                d_884_v123_ = _dafny.MultiSet([d_883_v122_, d_883_v122_, d_883_v122_])
                rhs102_ = d_879_v120_
                rhs103_ = d_882_v121_
                rhs104_ = (default__.safeModuloInt(p1, (d_882_v121_).f21)) <= ((len(d_852_v95_)) + ((d_882_v121_).f21))
                rhs105_ = (0) - (p1)
                rhs106_ = ((d_884_v123_ if (d_853_v96_)[default__.safeIndex(143, (d_853_v96_).length(0))] else d_884_v123_)) - (d_884_v123_)
                lhs101_ = globalState
                d_879_v120_ = rhs102_
                d_882_v121_ = rhs103_
                lhs101_.f6 = rhs104_
                r1 = rhs105_
                d_884_v123_ = rhs106_
            d_885_v124_: _dafny.MultiSet
            d_885_v124_ = _dafny.MultiSet([p0])
            d_886_v125_: str
            d_886_v125_ = _dafny.CodePoint('a')
            d_887_v126_: _dafny.Seq
            d_887_v126_ = _dafny.SeqWithoutIsStrInference([p0])
            d_888_v127_: _dafny.Set
            d_888_v127_ = _dafny.Set({len(d_887_v126_), p1})
            d_889_v128_: D0
            d_889_v128_ = D0_DC1(d_886_v125_, len(d_888_v127_), d_777_v28_, p0, p0)
            d_890_v129_: C1
            nw149_ = C1()
            nw149_.ctor__((default__.fm0((0) - (((d_885_v124_)[p0] if (p0) in (d_885_v124_) else p1)), d_889_v128_, globalState)) * (p1), (self).f20)
            d_890_v129_ = nw149_
            d_891_v130_: _dafny.Array
            nw150_ = _dafny.Array(int(0), 14)
            d_891_v130_ = nw150_
            index138_ = default__.safeIndex(860, (d_891_v130_).length(0))
            (d_891_v130_)[index138_] = p1
            index139_ = default__.safeIndex(860, (d_891_v130_).length(0))
            (d_891_v130_)[index139_] = p1
            index140_ = default__.safeIndex(860, (d_891_v130_).length(0))
            (d_891_v130_)[index140_] = p1
            d_891_v130_ = d_891_v130_
        d_892_v131_: _dafny.Seq
        d_892_v131_ = _dafny.SeqWithoutIsStrInference([p0, True])
        d_893_v132_: _dafny.Set
        d_893_v132_ = _dafny.Set({p1})
        d_894_v133_: D13
        d_894_v133_ = D13_DC35(d_893_v132_)
        d_895_v134_: _dafny.Seq
        d_895_v134_ = _dafny.SeqWithoutIsStrInference([(d_894_v133_).cf67])
        d_896_v135_: _dafny.MultiSet
        d_896_v135_ = _dafny.MultiSet([not(not(p0))])
        d_897_v136_: _dafny.Seq
        d_897_v136_ = _dafny.SeqWithoutIsStrInference([d_896_v135_])
        d_898_v137_: _dafny.MultiSet
        d_898_v137_ = (d_897_v136_)[default__.safeIndex(p1, len(d_897_v136_))]
        hi4_ = p1
        for d_899_i8_ in range(len((_dafny.Map({p1: d_892_v131_}) if p0 else _dafny.Map({len(d_895_v134_): default__.fm50(p0, p0, d_898_v137_, globalState)}))), hi4_):
            d_900_v138_: _dafny.Array
            nw151_ = _dafny.Array(D10.default()(), 17)
            d_900_v138_ = nw151_
            d_901_v139_: _dafny.Array
            nw152_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_901_v139_ = nw152_
            d_902_v140_: _dafny.Seq
            d_902_v140_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfk"))
            d_903_v141_: D10
            d_903_v141_ = D10_DC27(d_901_v139_, len(d_902_v140_))
            index141_ = default__.safeIndex(726, (d_900_v138_).length(0))
            (d_900_v138_)[index141_] = d_903_v141_
            d_904_v142_: str
            d_904_v142_ = _dafny.CodePoint('a')
            index142_ = default__.safeIndex(726, (d_900_v138_).length(0))
            rhs107_ = d_903_v141_
            rhs108_ = d_899_i8_
            rhs109_ = ((default__.fm46(p1, p0, False, globalState)) + (d_902_v140_)) + (((d_902_v140_) + (d_902_v140_)).set(default__.safeIndex(p1, len((d_902_v140_) + (d_902_v140_))), d_904_v142_))
            rhs110_ = not(p0)
            lhs102_ = d_900_v138_
            lhs103_ = default__.safeIndex(726, (d_900_v138_).length(0))
            lhs104_ = globalState
            lhs105_ = globalState
            lhs106_ = globalState
            lhs102_[lhs103_] = rhs107_
            lhs104_.f17 = rhs108_
            lhs105_.f15 = rhs109_
            lhs106_.f6 = rhs110_
            d_905_v143_: _dafny.MultiSet
            d_905_v143_ = _dafny.MultiSet([d_899_i8_])
            d_906_v144_: _dafny.MultiSet
            d_906_v144_ = _dafny.MultiSet([d_904_v142_, d_904_v142_])
            d_907_v145_: D0
            d_907_v145_ = D0_DC1(_dafny.CodePoint('c'), p1, _dafny.Set({p0, p0}), p0, not(p0))
            d_908_v146_: _dafny.Seq
            d_908_v146_ = _dafny.SeqWithoutIsStrInference([d_907_v145_])
            d_909_v147_: _dafny.Array
            nw153_ = _dafny.Array(None, 14)
            nw153_[int(0)] = (p0) or (p0)
            nw153_[int(1)] = p0
            nw153_[int(2)] = (492) > (430)
            nw153_[int(3)] = (_dafny.MultiSet([519])).isdisjoint(d_905_v143_)
            nw153_[int(4)] = default__.fm2(d_906_v144_, d_899_i8_, globalState)
            nw153_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_907_v145_])) < (d_908_v146_)
            nw153_[int(6)] = True
            nw153_[int(7)] = p0
            nw153_[int(8)] = p0
            nw153_[int(9)] = default__.fm2(d_906_v144_, p1, globalState)
            nw153_[int(10)] = p0
            nw153_[int(11)] = p0
            nw153_[int(12)] = p0
            nw153_[int(13)] = (p0) and (p0)
            d_909_v147_ = nw153_
            index143_ = default__.safeIndex(631, (d_909_v147_).length(0))
            (d_909_v147_)[index143_] = (default__.fm2(d_906_v144_, d_899_i8_, globalState) if p0 else p0)
            index144_ = default__.safeIndex(516, (d_909_v147_).length(0))
            (d_909_v147_)[index144_] = p0
            index145_ = default__.safeIndex(631, (d_909_v147_).length(0))
            index146_ = default__.safeIndex(516, (d_909_v147_).length(0))
            rhs111_ = p0
            rhs112_ = ((d_899_i8_) >= ((0) - ((d_905_v143_).cardinality)) if p0 else p0)
            rhs113_ = p0
            rhs114_ = p1
            lhs107_ = globalState
            lhs108_ = d_909_v147_
            lhs109_ = default__.safeIndex(631, (d_909_v147_).length(0))
            lhs110_ = d_909_v147_
            lhs111_ = default__.safeIndex(516, (d_909_v147_).length(0))
            lhs112_ = globalState
            lhs107_.f6 = rhs111_
            lhs108_[lhs109_] = rhs112_
            lhs110_[lhs111_] = rhs113_
            lhs112_.f17 = rhs114_
            d_910_v148_: _dafny.Map
            d_910_v148_ = _dafny.Map({p0: d_909_v147_})
            d_911_v149_: D9
            d_911_v149_ = D9_DC24(d_910_v148_)
            pat_let_tv33_ = d_910_v148_
            def iife132_(_pat_let47_0):
                def iife133_(d_912_dt__update__tmp_h3_):
                    def iife134_(_pat_let48_0):
                        def iife135_(d_913_dt__update_hcf45_h0_):
                            return D9_DC24(d_913_dt__update_hcf45_h0_)
                        return iife135_(_pat_let48_0)
                    return iife134_(pat_let_tv33_)
                return iife133_(_pat_let47_0)
            source12_ = (d_911_v149_ if not(p0) else iife132_(D9_DC24(d_910_v148_)))
            if source12_.is_DC25:
                d_914___mcc_h9_ = source12_.cf46
                d_915___mcc_h10_ = source12_.cf47
                d_916___mcc_h11_ = source12_.cf48
                d_917_cf48_ = d_916___mcc_h11_
                d_918_cf47_ = d_915___mcc_h10_
                d_919_cf46_ = d_914___mcc_h9_
                d_920_v150_: _dafny.Array
                nw154_ = _dafny.Array(int(0), 14)
                d_920_v150_ = nw154_
                index147_ = default__.safeIndex(627, (d_920_v150_).length(0))
                (d_920_v150_)[index147_] = p1
                d_921_v151_: _dafny.Map
                d_921_v151_ = _dafny.Map({(d_909_v147_)[default__.safeIndex(631, (d_909_v147_).length(0))]: d_904_v142_})
                d_922_v152_: _dafny.Seq
                d_922_v152_ = _dafny.SeqWithoutIsStrInference([len(d_921_v151_)])
                d_923_v153_: _dafny.MultiSet
                d_923_v153_ = _dafny.MultiSet([_dafny.MultiSet([(d_909_v147_)[default__.safeIndex(631, (d_909_v147_).length(0))]])])
                index148_ = default__.safeIndex(627, (d_920_v150_).length(0))
                index149_ = default__.safeIndex(631, (d_909_v147_).length(0))
                rhs115_ = default__.safeModuloInt(len(d_922_v152_), 695)
                rhs116_ = len(d_895_v134_)
                rhs117_ = d_919_cf46_
                rhs118_ = not(not(((d_923_v153_) | (_dafny.MultiSet([d_896_v135_, d_896_v135_]))).issubset(_dafny.MultiSet([_dafny.MultiSet([p0]), _dafny.MultiSet(d_892_v131_), d_896_v135_]))))
                lhs113_ = globalState
                lhs114_ = d_920_v150_
                lhs115_ = default__.safeIndex(627, (d_920_v150_).length(0))
                lhs116_ = d_909_v147_
                lhs117_ = default__.safeIndex(631, (d_909_v147_).length(0))
                d_919_cf46_ = rhs115_
                lhs113_.f17 = rhs116_
                lhs114_[lhs115_] = rhs117_
                lhs116_[lhs117_] = rhs118_
                (globalState).f6 = not((d_909_v147_)[default__.safeIndex(631, (d_909_v147_).length(0))])
                d_924_v154_: C1
                nw155_ = C1()
                nw155_.ctor__(d_919_cf46_, (self).f20)
                d_924_v154_ = nw155_
                index150_ = default__.safeIndex(631, (d_909_v147_).length(0))
                index151_ = default__.safeIndex(631, (d_909_v147_).length(0))
                index152_ = default__.safeIndex(631, (d_909_v147_).length(0))
                index153_ = default__.safeIndex(631, (d_909_v147_).length(0))
                rhs119_ = (d_896_v135_).issubset((_dafny.MultiSet(d_892_v131_)) - (d_896_v135_))
                rhs120_ = p0
                rhs121_ = not ((p1) not in (_dafny.Map({len(_dafny.Map({d_924_v154_: d_918_cf47_})): d_898_v137_}))) or (True)
                rhs122_ = d_902_v140_
                rhs123_ = p0
                lhs118_ = d_909_v147_
                lhs119_ = default__.safeIndex(631, (d_909_v147_).length(0))
                lhs120_ = d_909_v147_
                lhs121_ = default__.safeIndex(631, (d_909_v147_).length(0))
                lhs122_ = d_909_v147_
                lhs123_ = default__.safeIndex(631, (d_909_v147_).length(0))
                lhs124_ = globalState
                lhs125_ = d_909_v147_
                lhs126_ = default__.safeIndex(631, (d_909_v147_).length(0))
                lhs118_[lhs119_] = rhs119_
                lhs120_[lhs121_] = rhs120_
                lhs122_[lhs123_] = rhs121_
                lhs124_.f16 = rhs122_
                lhs125_[lhs126_] = rhs123_
                d_925_v155_: _dafny.Set
                d_925_v155_ = _dafny.Set({d_909_v147_, d_909_v147_})
                (globalState).f18 = len(d_925_v155_)
            elif True:
                d_926___mcc_h12_ = source12_.cf45
                d_927_cf45_ = d_926___mcc_h12_
                d_928_v156_: D10
                d_928_v156_ = D10_DC28(d_902_v140_, p1, True, (d_905_v143_).issubset(d_905_v143_), (d_909_v147_)[default__.safeIndex(631, (d_909_v147_).length(0))])
                d_928_v156_ = default__.fm53(default__.fm2(d_906_v144_, d_899_i8_, globalState), globalState)
                d_929_v157_: C1
                nw156_ = C1()
                nw156_.ctor__(p1, (self).f20)
                d_929_v157_ = nw156_
                r1 = d_899_i8_
                d_930_v158_: C1
                nw157_ = C1()
                nw157_.ctor__(p1, (self).f20)
                d_930_v158_ = nw157_
            (globalState).f18 = len(d_902_v140_)
        d_931_i9_: int
        d_931_i9_ = 0
        with _dafny.label("3"):
            while p0:
                with _dafny.c_label("3"):
                    if (d_931_i9_) >= (100):
                        raise _dafny.Break("3")
                    d_931_i9_ = (d_931_i9_) + (1)
                    d_932_v159_: str
                    d_932_v159_ = _dafny.CodePoint('q')
                    d_933_v160_: _dafny.MultiSet
                    d_933_v160_ = _dafny.MultiSet([d_932_v159_, d_932_v159_])
                    d_892_v131_ = (((d_892_v131_) + (d_892_v131_)).set(default__.safeIndex(p1, len((d_892_v131_) + (d_892_v131_))), not(default__.fm2(d_933_v160_, ((d_896_v135_)[p0] if (p0) in (d_896_v135_) else p1), globalState)))) + (d_892_v131_)
                    if (d_893_v132_).ispropersubset((d_893_v132_) - (d_893_v132_)):
                        d_934_v161_: _dafny.Map
                        d_934_v161_ = _dafny.Map({p0: 276})
                        d_934_v161_ = (d_934_v161_).set(p0, -457)
                        d_932_v159_ = default__.fm1(globalState)
                        d_935_v162_: C1
                        nw158_ = C1()
                        nw158_.ctor__((default__.fm54(default__.fm2(d_933_v160_, (0) - (-75), globalState), globalState)).cardinality, (self).f20)
                        d_935_v162_ = nw158_
                        r1 = p1
                        d_936_v163_: _dafny.Seq
                        d_936_v163_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyiksls"))
                        d_937_v164_: _dafny.Set
                        d_937_v164_ = _dafny.Set({d_936_v163_})
                        (globalState).f6 = (len(d_937_v164_)) <= (p1)
                    elif True:
                        (globalState).f17 = p1
                        d_938_v165_: D3
                        d_938_v165_ = D3_DC8(default__.fm2(d_933_v160_, p1, globalState))
                        d_939_v166_: _dafny.MultiSet
                        d_939_v166_ = _dafny.MultiSet([d_938_v165_, default__.fm55(globalState), D3_DC8(p0)])
                        (globalState).f17 = (d_939_v166_).cardinality
                        d_940_v167_: _dafny.Map
                        d_940_v167_ = _dafny.Map({default__.fm2(_dafny.MultiSet([d_932_v159_]), p1, globalState): p1})
                        d_941_v168_: _dafny.Seq
                        d_941_v168_ = _dafny.SeqWithoutIsStrInference([d_940_v167_])
                        (globalState).f5 = (_dafny.Map({default__.fm2(d_933_v160_, p1, globalState): p1})) != ((d_941_v168_)[default__.safeIndex(p1, len(d_941_v168_))])
                        d_942_v169_: _dafny.Array
                        def lambda28_(d_943_p1_):
                            def lambda29_(d_944_i10_):
                                return (d_944_i10_) - ((0) - (d_943_p1_))

                            return lambda29_

                        init14_ = lambda28_(p1)
                        nw159_ = _dafny.Array(None, 3)
                        for i0_14_ in range(nw159_.length(0)):
                            nw159_[i0_14_] = init14_(i0_14_)
                        d_942_v169_ = nw159_
                        d_942_v169_ = d_942_v169_
                        (globalState).f5 = p0
                    d_945_v170_: _dafny.MultiSet
                    d_945_v170_ = _dafny.MultiSet([p1])
                    d_946_v171_: _dafny.Map
                    d_946_v171_ = _dafny.Map({p0: d_945_v170_})
                    (globalState).f6 = ((((d_946_v171_)[p0] if (p0) in (d_946_v171_) else d_945_v170_)) - (d_945_v170_)).issubset(((default__.fm56(p1, p0, False, (d_945_v170_).cardinality, globalState)).set(p1, default__.abs(p1))).intersection(d_945_v170_))
                    if p0:
                        d_947_v172_: _dafny.Array
                        def lambda30_(d_948_p0_):
                            def lambda31_(d_949_i11_):
                                return not(d_948_p0_)

                            return lambda31_

                        init15_ = lambda30_(p0)
                        nw160_ = _dafny.Array(None, 15)
                        for i0_15_ in range(nw160_.length(0)):
                            nw160_[i0_15_] = init15_(i0_15_)
                        d_947_v172_ = nw160_
                        (globalState).f4 = d_947_v172_
                        d_950_v173_: _dafny.Seq
                        d_950_v173_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-101 for d_951_i12_ in range(default__.abs(852))]))])
                        d_950_v173_ = d_950_v173_
                        d_952_v174_: _dafny.Seq
                        d_952_v174_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umw"))
                        d_953_v175_: _dafny.Array
                        nw161_ = _dafny.Array(None, 10)
                        nw161_[int(0)] = (d_952_v174_) + (d_952_v174_)
                        nw161_[int(1)] = d_952_v174_
                        nw161_[int(2)] = d_952_v174_
                        nw161_[int(3)] = (d_952_v174_) + (d_952_v174_)
                        nw161_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ads"))
                        nw161_[int(5)] = d_952_v174_
                        nw161_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_932_v159_ for d_954_i13_ in range(default__.abs(677))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
                        nw161_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aoxrrsm"))
                        nw161_[int(8)] = _dafny.SeqWithoutIsStrInference([d_932_v159_ for d_955_i14_ in range(default__.abs(24))])
                        nw161_[int(9)] = (d_952_v174_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mn")))
                        d_953_v175_ = nw161_
                        d_953_v175_ = (d_953_v175_ if p0 else d_953_v175_)
                        d_956_v176_: _dafny.Array
                        def lambda32_(d_957_v131_):
                            def lambda33_(d_958_i15_):
                                return (d_958_i15_) - (len(d_957_v131_))

                            return lambda33_

                        init16_ = lambda32_(d_892_v131_)
                        nw162_ = _dafny.Array(None, 16)
                        for i0_16_ in range(nw162_.length(0)):
                            nw162_[i0_16_] = init16_(i0_16_)
                        d_956_v176_ = nw162_
                        d_959_v177_: _dafny.Array
                        nw163_ = _dafny.Array(_dafny.Seq({}), 24)
                        d_959_v177_ = nw163_
                        d_960_v178_: D8
                        d_960_v178_ = D8_DC23(d_952_v174_)
                        d_961_v179_: _dafny.Seq
                        d_961_v179_ = _dafny.SeqWithoutIsStrInference([D8_DC23(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggjfws"))), d_960_v178_, d_960_v178_])
                        index154_ = default__.safeIndex(386, (d_959_v177_).length(0))
                        (d_959_v177_)[index154_] = d_961_v179_
                        d_962_v180_: _dafny.Seq
                        d_962_v180_ = _dafny.SeqWithoutIsStrInference([d_952_v174_])
                        pat_let_tv34_ = d_952_v174_
                        index155_ = default__.safeIndex(386, (d_959_v177_).length(0))
                        def iife136_(_pat_let49_0):
                            def iife137_(d_963_dt__update__tmp_h4_):
                                def iife138_(_pat_let50_0):
                                    def iife139_(d_964_dt__update_hcf44_h0_):
                                        return D8_DC23(d_964_dt__update_hcf44_h0_)
                                    return iife139_(_pat_let50_0)
                                return iife138_(pat_let_tv34_)
                            return iife137_(_pat_let49_0)
                        rhs124_ = d_956_v176_
                        rhs125_ = (d_947_v172_ if p0 else d_947_v172_)
                        rhs126_ = _dafny.SeqWithoutIsStrInference([D8_DC23(d_952_v174_), iife136_(D8_DC23(default__.fm46(p1, p0, False, globalState))), D8_DC23(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbnrtear"))), D8_DC23(default__.fm46(len(d_962_v180_), p0, not(p0), globalState))])
                        rhs127_ = p1
                        lhs127_ = globalState
                        lhs128_ = d_959_v177_
                        lhs129_ = default__.safeIndex(386, (d_959_v177_).length(0))
                        lhs130_ = globalState
                        d_956_v176_ = rhs124_
                        lhs127_.f4 = rhs125_
                        lhs128_[lhs129_] = rhs126_
                        lhs130_.f17 = rhs127_
                        index156_ = default__.safeIndex(879, (d_947_v172_).length(0))
                        (d_947_v172_)[index156_] = not (p0) or (True)
                        index157_ = default__.safeIndex(879, (d_947_v172_).length(0))
                        (d_947_v172_)[index157_] = p0
                    elif True:
                        r1 = ((0) - (p1) if p0 else default__.safeDivisionInt(-374, p1))
                        d_965_v181_: D2
                        d_965_v181_ = D2_DC4((0) - (p1), p0, p0, p0)
                        d_966_v182_: _dafny.Seq
                        d_966_v182_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])
                        d_967_v183_: D0
                        d_967_v183_ = D0_DC1((d_966_v182_)[default__.safeIndex(p1, len(d_966_v182_))], p1, d_777_v28_, not(True), p0)
                        d_968_v184_: _dafny.Array
                        nw164_ = _dafny.Array(None, 10)
                        nw164_[int(0)] = p1
                        nw164_[int(1)] = (d_965_v181_).cf8
                        nw164_[int(2)] = p1
                        nw164_[int(3)] = (0) - (p1)
                        nw164_[int(4)] = default__.fm0(p1, d_967_v183_, globalState)
                        nw164_[int(5)] = 324
                        nw164_[int(6)] = p1
                        nw164_[int(7)] = p1
                        nw164_[int(8)] = 208
                        nw164_[int(9)] = (p1) + (p1)
                        d_968_v184_ = nw164_
                        index158_ = default__.safeIndex(21, (d_968_v184_).length(0))
                        (d_968_v184_)[index158_] = p1
                        d_969_v185_: _dafny.Set
                        d_969_v185_ = _dafny.Set({_dafny.Set({p0, p0, p0}), d_777_v28_, d_777_v28_})
                        d_970_v186_: D6
                        d_970_v186_ = D6_DC17(d_945_v170_, d_969_v185_)
                        d_971_v187_: _dafny.Seq
                        d_971_v187_ = _dafny.SeqWithoutIsStrInference([(d_968_v184_ if p0 else d_968_v184_), d_968_v184_])
                        d_972_v188_: _dafny.Seq
                        d_972_v188_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_973_v189_: _dafny.Map
                        d_973_v189_ = _dafny.Map({len(d_972_v188_): p0})
                        index159_ = default__.safeIndex(21, (d_968_v184_).length(0))
                        rhs128_ = (0) - (p1)
                        rhs129_ = (d_971_v187_)[default__.safeIndex((p1) * (len(d_973_v189_)), len(d_971_v187_))]
                        rhs130_ = (D7_DC20(len(d_892_v131_), p0)).cf36
                        rhs131_ = p0
                        rhs132_ = d_970_v186_
                        lhs131_ = globalState
                        lhs132_ = d_968_v184_
                        lhs133_ = default__.safeIndex(21, (d_968_v184_).length(0))
                        lhs134_ = globalState
                        lhs131_.f17 = rhs128_
                        d_968_v184_ = rhs129_
                        lhs132_[lhs133_] = rhs130_
                        lhs134_.f5 = rhs131_
                        d_970_v186_ = rhs132_
                        (globalState).f18 = default__.safeModuloInt(p1, (d_968_v184_)[default__.safeIndex(21, (d_968_v184_).length(0))])
                        d_974_v190_: _dafny.Array
                        nw165_ = _dafny.Array(False, 21)
                        d_974_v190_ = nw165_
                        d_975_v191_: C0
                        nw166_ = C0()
                        nw166_.ctor__(d_974_v190_, (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([p0])))
                        d_975_v191_ = nw166_
                        r2 = _dafny.CodePoint('i')
                    pass
            pass
        d_976_v192_: _dafny.Map
        d_976_v192_ = _dafny.Map({not(False): p1})
        hi5_ = p1
        for d_977_i16_ in range(len(d_976_v192_), hi5_):
            (globalState).f5 = p0
            d_978_v193_: str
            d_978_v193_ = _dafny.CodePoint('w')
            d_979_v194_: _dafny.Seq
            d_979_v194_ = _dafny.SeqWithoutIsStrInference([d_978_v193_])
            d_980_v195_: _dafny.Set
            d_980_v195_ = _dafny.Set({d_979_v194_, d_979_v194_, d_979_v194_})
            d_980_v195_ = d_980_v195_
            d_981_v196_: _dafny.Map
            d_981_v196_ = _dafny.Map({p0: p0})
            (globalState).f18 = len(d_981_v196_)
            (globalState).f18 = d_977_i16_
        r0 = (d_896_v135_) - (d_896_v135_)
        d_982_v197_: str
        d_982_v197_ = _dafny.CodePoint('x')
        d_983_v198_: _dafny.Seq
        d_983_v198_ = _dafny.SeqWithoutIsStrInference([d_982_v197_, d_982_v197_, d_982_v197_])
        d_984_v199_: _dafny.Map
        d_984_v199_ = _dafny.Map({d_983_v198_: d_982_v197_})
        r1 = (len(d_984_v199_) if p0 else 655)
        r2 = d_982_v197_
        return r0, r1, r2


class C4(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f20):
        (self)._f20 = f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        if p0:
            d_985_v0_: _dafny.Seq
            d_985_v0_ = _dafny.SeqWithoutIsStrInference([p0])
            d_986_v1_: str
            d_986_v1_ = _dafny.CodePoint('e')
            d_987_v2_: _dafny.MultiSet
            d_987_v2_ = _dafny.MultiSet([d_986_v1_])
            d_988_v3_: _dafny.Seq
            d_988_v3_ = _dafny.SeqWithoutIsStrInference([p1])
            d_989_v4_: _dafny.Array
            nw167_ = _dafny.Array(False, 29)
            d_989_v4_ = nw167_
            d_990_v5_: _dafny.Map
            d_990_v5_ = _dafny.Map({True: d_989_v4_})
            rhs133_ = (d_985_v0_) + (d_985_v0_)
            rhs134_ = default__.fm2((d_987_v2_).set(d_986_v1_, default__.abs(p1)), (d_988_v3_)[default__.safeIndex(len(d_990_v5_), len(d_988_v3_))], globalState)
            lhs135_ = globalState
            d_985_v0_ = rhs133_
            lhs135_.f6 = rhs134_
            d_991_v6_: _dafny.Array
            def lambda34_(d_992_v1_):
                def lambda35_(d_993_i0_):
                    return d_992_v1_

                return lambda35_

            init17_ = lambda34_(d_986_v1_)
            nw168_ = _dafny.Array(None, 10)
            for i0_17_ in range(nw168_.length(0)):
                nw168_[i0_17_] = init17_(i0_17_)
            d_991_v6_ = nw168_
            d_994_v7_: _dafny.Map
            d_994_v7_ = _dafny.Map({d_991_v6_: _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])})
            (globalState).f17 = len(d_994_v7_)
            d_995_v8_: _dafny.Seq
            d_995_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfiodfgc"))
            d_996_v9_: _dafny.Map
            d_996_v9_ = _dafny.Map({len((d_995_v8_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdn")))): (d_985_v0_) == (_dafny.SeqWithoutIsStrInference([p0, p0, p0]))})
            d_996_v9_ = d_996_v9_
            source13_ = default__.fm44(p0, globalState)
            if source13_.is_DC8:
                d_997___mcc_h0_ = source13_.cf16
                d_998_cf16_ = d_997___mcc_h0_
                d_999_v10_: _dafny.Array
                nw169_ = _dafny.Array(_dafny.Set({}), 20)
                d_999_v10_ = nw169_
                index160_ = default__.safeIndex(720, (d_999_v10_).length(0))
                def iife140_():
                    coll38_ = _dafny.Set()
                    compr_38_: int
                    for compr_38_ in _dafny.IntegerRange(48, 933):
                        d_1000_v11_: int = compr_38_
                        if ((48) <= (d_1000_v11_)) and ((d_1000_v11_) < (933)):
                            coll38_ = coll38_.union(_dafny.Set([(d_1000_v11_) * (p1)]))
                    return _dafny.Set(coll38_)
                (d_999_v10_)[index160_] = iife140_()
                
                d_1001_v12_: _dafny.Set
                d_1001_v12_ = _dafny.Set({p1})
                index161_ = default__.safeIndex(720, (d_999_v10_).length(0))
                (d_999_v10_)[index161_] = d_1001_v12_
                d_1002_v13_: C2
                nw170_ = C2()
                nw170_.ctor__(p1)
                d_1002_v13_ = nw170_
                d_1003_v14_: _dafny.Array
                d_1004_v15_: bool
                d_1005_v16_: _dafny.Set
                out7_: _dafny.Array
                out8_: bool
                out9_: _dafny.Set
                out7_, out8_, out9_ = (self).m12(globalState)
                d_1003_v14_ = out7_
                d_1004_v15_ = out8_
                d_1005_v16_ = out9_
                d_1006_v17_: _dafny.Set
                d_1006_v17_ = _dafny.Set({d_998_cf16_, d_998_cf16_})
                d_1006_v17_ = d_1006_v17_
            elif source13_.is_DC7:
                d_1007___mcc_h1_ = source13_.cf15
                d_1008_cf15_ = d_1007___mcc_h1_
                d_1009_v18_: _dafny.Array
                def lambda36_(d_1010_p1_):
                    def lambda37_(d_1011_i1_):
                        return (d_1011_i1_) + (d_1010_p1_)

                    return lambda37_

                init18_ = lambda36_(p1)
                nw171_ = _dafny.Array(None, 7)
                for i0_18_ in range(nw171_.length(0)):
                    nw171_[i0_18_] = init18_(i0_18_)
                d_1009_v18_ = nw171_
                rhs135_ = d_1009_v18_
                rhs136_ = (d_1008_cf15_ if p0 else d_1008_cf15_)
                d_1009_v18_ = rhs135_
                d_1008_cf15_ = rhs136_
                d_1012_v19_: _dafny.Map
                d_1012_v19_ = _dafny.Map({p1: p1})
                d_1013_v20_: _dafny.MultiSet
                d_1013_v20_ = _dafny.MultiSet([p1])
                d_1014_v21_: _dafny.Map
                d_1014_v21_ = _dafny.Map({p0: d_1009_v18_})
                d_1015_v22_: _dafny.Set
                d_1015_v22_ = _dafny.Set({p1, p1, len((d_1012_v19_).set((d_1013_v20_).cardinality, len(d_1014_v21_))), p1})
                d_1015_v22_ = d_1015_v22_
                (globalState).f18 = p1
                (globalState).f6 = p0
            elif True:
                d_1016___mcc_h2_ = source13_.cf17
                d_1017_cf17_ = d_1016___mcc_h2_
                (globalState).f5 = p0
                d_1018_v23_: _dafny.Map
                d_1018_v23_ = _dafny.Map({p1: p1})
                d_1018_v23_ = d_1018_v23_
                d_1019_v24_: C1
                nw172_ = C1()
                nw172_.ctor__(p1, (self).f20)
                d_1019_v24_ = nw172_
                d_1020_v25_: _dafny.Set
                d_1020_v25_ = _dafny.Set({p0, p0})
                d_1021_v26_: _dafny.Map
                d_1021_v26_ = _dafny.Map({len(d_995_v8_): d_1020_v25_})
                d_1022_v27_: _dafny.Array
                nw173_ = _dafny.Array(None, 18)
                nw173_[int(0)] = d_989_v4_
                nw173_[int(1)] = d_989_v4_
                nw173_[int(2)] = d_989_v4_
                nw173_[int(3)] = d_989_v4_
                nw173_[int(4)] = d_989_v4_
                nw173_[int(5)] = d_989_v4_
                nw173_[int(6)] = d_989_v4_
                nw173_[int(7)] = d_989_v4_
                nw173_[int(8)] = d_989_v4_
                nw173_[int(9)] = d_989_v4_
                nw173_[int(10)] = d_989_v4_
                nw173_[int(11)] = d_989_v4_
                nw173_[int(12)] = d_989_v4_
                nw173_[int(13)] = d_989_v4_
                nw173_[int(14)] = d_989_v4_
                nw173_[int(15)] = d_989_v4_
                nw173_[int(16)] = d_989_v4_
                nw173_[int(17)] = d_989_v4_
                d_1022_v27_ = nw173_
                d_1023_v28_: D10
                d_1023_v28_ = D10_DC27(d_1022_v27_, len(d_988_v3_))
                d_1024_v29_: _dafny.Map
                d_1024_v29_ = _dafny.Map({d_1023_v28_: p1})
                d_1025_v30_: _dafny.Map
                d_1025_v30_ = _dafny.Map({d_1024_v29_: p0})
                d_1026_v31_: _dafny.Seq
                d_1026_v31_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), d_986_v1_, d_986_v1_, d_986_v1_, d_986_v1_])) + (d_995_v8_), d_995_v8_, (D10_DC28(d_995_v8_, p1, default__.fm2(d_987_v2_, len(_dafny.Map({p1: ((d_1021_v26_)[len(d_1025_v30_)] if (len(d_1025_v30_)) in (d_1021_v26_) else d_1020_v25_)})), globalState), p0, not(p0))).cf52, (d_995_v8_) + (_dafny.SeqWithoutIsStrInference([d_986_v1_])), d_995_v8_])
                (globalState).f16 = (d_1026_v31_)[default__.safeIndex(p1, len(d_1026_v31_))]
            (globalState).f18 = default__.safeModuloInt(765, p1)
        elif True:
            if p0:
                d_1027_v32_: _dafny.Set
                d_1027_v32_ = _dafny.Set({p1, (0) - (p1), p1})
                d_1028_v33_: _dafny.Map
                d_1028_v33_ = _dafny.Map({p0: len(d_1027_v32_)})
                d_1029_v34_: _dafny.Set
                d_1029_v34_ = _dafny.Set({p0, p0})
                d_1028_v33_ = (d_1028_v33_).set(p0, default__.safeModuloInt(len(d_1029_v34_), p1))
                (globalState).f17 = (p1) - (p1)
                d_1030_v35_: _dafny.Array
                nw174_ = _dafny.Array(_dafny.Seq({}), 3)
                d_1030_v35_ = nw174_
                d_1031_v36_: _dafny.Seq
                d_1031_v36_ = _dafny.SeqWithoutIsStrInference([p1])
                index162_ = default__.safeIndex(10, (d_1030_v35_).length(0))
                (d_1030_v35_)[index162_] = d_1031_v36_
                index163_ = default__.safeIndex(10, (d_1030_v35_).length(0))
                (d_1030_v35_)[index163_] = (d_1031_v36_) + (_dafny.SeqWithoutIsStrInference([p1 for d_1032_i2_ in range(default__.abs(136))]))
                d_1033_v37_: _dafny.Seq
                d_1033_v37_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1034_v38_: _dafny.Seq
                d_1034_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                d_1035_v39_: _dafny.Array
                nw175_ = _dafny.Array(None, 6)
                nw175_[int(0)] = p1
                nw175_[int(1)] = p1
                nw175_[int(2)] = p1
                nw175_[int(3)] = len(d_1033_v37_)
                nw175_[int(4)] = len(d_1034_v38_)
                nw175_[int(5)] = -731
                d_1035_v39_ = nw175_
                d_1036_v40_: _dafny.Seq
                d_1036_v40_ = _dafny.SeqWithoutIsStrInference([d_1035_v39_, d_1035_v39_, d_1035_v39_, d_1035_v39_])
                d_1037_v41_: _dafny.Array
                nw176_ = _dafny.Array(None, 6)
                nw176_[int(0)] = (d_1034_v38_) + (d_1034_v38_)
                nw176_[int(1)] = d_1034_v38_
                nw176_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oawn"))
                nw176_[int(3)] = (d_1034_v38_) + (d_1034_v38_)
                nw176_[int(4)] = d_1034_v38_
                nw176_[int(5)] = d_1034_v38_
                d_1037_v41_ = nw176_
                index164_ = default__.safeIndex(169, (d_1037_v41_).length(0))
                (d_1037_v41_)[index164_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cedjhgqk"))) + (d_1034_v38_)
                d_1038_v42_: D4
                d_1038_v42_ = D4_DC11(p0, True, p1, len(d_1034_v38_))
                d_1039_v43_: _dafny.Map
                d_1039_v43_ = _dafny.Map({p0: not(not((d_1038_v42_).cf19))})
                d_1040_v44_: _dafny.Map
                d_1040_v44_ = _dafny.Map({p1: p0})
                d_1041_v45_: _dafny.Map
                d_1041_v45_ = _dafny.Map({d_1039_v43_: len(d_1040_v44_)})
                index165_ = default__.safeIndex(169, (d_1037_v41_).length(0))
                rhs137_ = (d_1036_v40_) + (d_1036_v40_)
                rhs138_ = default__.safeModuloInt(857, len(d_1041_v45_))
                rhs139_ = (d_1034_v38_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixbopwr")))
                lhs136_ = globalState
                lhs137_ = d_1037_v41_
                lhs138_ = default__.safeIndex(169, (d_1037_v41_).length(0))
                d_1036_v40_ = rhs137_
                lhs136_.f17 = rhs138_
                lhs137_[lhs138_] = rhs139_
                d_1042_v46_: _dafny.Map
                d_1042_v46_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byt"))})
                d_1043_v47_: _dafny.Map
                d_1043_v47_ = _dafny.Map({d_1035_v39_: _dafny.CodePoint('a')})
                (globalState).f6 = (((d_1042_v46_)[len(d_1043_v47_)] if (len(d_1043_v47_)) in (d_1042_v46_) else (d_1037_v41_)[default__.safeIndex(169, (d_1037_v41_).length(0))])) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))
            elif True:
                (globalState).f5 = not (p0) or (p0)
                (globalState).f17 = (0) - (p1)
                d_1044_v48_: _dafny.Seq
                d_1044_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gs"))
                d_1045_v49_: str
                d_1045_v49_ = _dafny.CodePoint('y')
                d_1046_v50_: _dafny.Set
                d_1046_v50_ = _dafny.Set({not(p0)})
                d_1047_v51_: D0
                d_1047_v51_ = D0_DC1(d_1045_v49_, p1, d_1046_v50_, p0, p0)
                (globalState).f18 = (0) - (default__.fm0(len(d_1044_v48_), d_1047_v51_, globalState))
                d_1048_v52_: _dafny.Array
                d_1049_v53_: bool
                d_1050_v54_: _dafny.Set
                out10_: _dafny.Array
                out11_: bool
                out12_: _dafny.Set
                out10_, out11_, out12_ = (self).m12(globalState)
                d_1048_v52_ = out10_
                d_1049_v53_ = out11_
                d_1050_v54_ = out12_
                d_1051_v55_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Set({}), 5)
                d_1051_v55_ = nw177_
                d_1052_v56_: _dafny.Map
                d_1052_v56_ = _dafny.Map({d_1051_v55_: d_1044_v48_})
                d_1052_v56_ = (d_1052_v56_).set((d_1051_v55_ if d_1049_v53_ else d_1051_v55_), d_1044_v48_)
            d_1053_v57_: _dafny.Seq
            d_1053_v57_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            (globalState).f17 = len(((d_1053_v57_) + (d_1053_v57_) if p0 else d_1053_v57_))
            d_1054_v58_: C2
            nw178_ = C2()
            nw178_.ctor__(962)
            d_1054_v58_ = nw178_
            rhs140_ = p0
            rhs141_ = d_1054_v58_
            lhs139_ = globalState
            lhs139_.f6 = rhs140_
            d_1054_v58_ = rhs141_
            d_1055_v59_: C1
            nw179_ = C1()
            nw179_.ctor__(default__.safeModuloInt(p1, (0) - ((0) - ((d_1054_v58_).f22))), (self).f20)
            d_1055_v59_ = nw179_
            (globalState).f17 = ((d_1054_v58_).f22 if True else (d_1054_v58_).f22)
        (globalState).f5 = not(p0)
        d_1056_v60_: _dafny.Array
        def lambda38_(d_1057_p0_):
            def lambda39_(d_1058_i3_):
                return not(not(d_1057_p0_))

            return lambda39_

        init19_ = lambda38_(p0)
        nw180_ = _dafny.Array(None, 25)
        for i0_19_ in range(nw180_.length(0)):
            nw180_[i0_19_] = init19_(i0_19_)
        d_1056_v60_ = nw180_
        d_1059_v61_: str
        d_1059_v61_ = _dafny.CodePoint('x')
        d_1060_v62_: _dafny.MultiSet
        d_1060_v62_ = _dafny.MultiSet([d_1059_v61_, d_1059_v61_])
        index166_ = default__.safeIndex(288, (d_1056_v60_).length(0))
        (d_1056_v60_)[index166_] = default__.fm2(d_1060_v62_, p1, globalState)
        d_1061_v63_: _dafny.Set
        d_1061_v63_ = _dafny.Set({p1})
        d_1062_v64_: _dafny.Seq
        d_1062_v64_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, (d_1061_v63_) != (d_1061_v63_), p0])
        index167_ = default__.safeIndex(288, (d_1056_v60_).length(0))
        (d_1056_v60_)[index167_] = (d_1062_v64_)[default__.safeIndex(p1, len(d_1062_v64_))]
        d_1063_v65_: _dafny.Array
        def lambda40_(d_1064_p1_):
            def lambda41_(d_1065_i4_):
                return (d_1065_i4_) - (d_1064_p1_)

            return lambda41_

        init20_ = lambda40_(p1)
        nw181_ = _dafny.Array(None, 9)
        for i0_20_ in range(nw181_.length(0)):
            nw181_[i0_20_] = init20_(i0_20_)
        d_1063_v65_ = nw181_
        d_1063_v65_ = d_1063_v65_
        d_1066_i5_: int
        d_1066_i5_ = 0
        with _dafny.label("4"):
            while (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]:
                with _dafny.c_label("4"):
                    if (d_1066_i5_) >= (100):
                        raise _dafny.Break("4")
                    d_1066_i5_ = (d_1066_i5_) + (1)
                    d_1067_v66_: _dafny.Set
                    d_1067_v66_ = _dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                    d_1068_v68_: _dafny.Map
                    d_1068_v68_ = _dafny.Map({d_1059_v61_: (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                    d_1069_v69_: _dafny.Map
                    def iife141_():
                        coll39_ = _dafny.Map()
                        compr_39_: str
                        for compr_39_ in (d_1068_v68_).keys.Elements:
                            d_1070_v67_: str = compr_39_
                            if (d_1070_v67_) in (d_1068_v68_):
                                coll39_[d_1070_v67_] = p1
                        return _dafny.Map(coll39_)
                    d_1069_v69_ = _dafny.Map({d_1067_v66_: (iife141_()
                    ).set(d_1059_v61_, p1)})
                    d_1071_v70_: _dafny.Map
                    d_1071_v70_ = _dafny.Map({d_1059_v61_: 32})
                    d_1072_v71_: _dafny.Seq
                    d_1072_v71_ = _dafny.SeqWithoutIsStrInference([len(d_1061_v63_)])
                    d_1073_v72_: _dafny.Map
                    d_1073_v72_ = _dafny.Map({len(((d_1069_v69_)[d_1067_v66_] if (d_1067_v66_) in (d_1069_v69_) else d_1071_v70_)): (default__.fm45(globalState)) + (d_1072_v71_)})
                    d_1073_v72_ = (d_1073_v72_).set((p1) * (p1), (d_1072_v71_ if p0 else d_1072_v71_))
                    d_1074_v73_: _dafny.Seq
                    d_1074_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgmrq"))
                    d_1075_v74_: C2
                    nw182_ = C2()
                    nw182_.ctor__(p1)
                    d_1075_v74_ = nw182_
                    d_1076_v75_: _dafny.Map
                    d_1076_v75_ = _dafny.Map({len(d_1074_v73_): len(_dafny.Map({d_1075_v74_: p0}))})
                    d_1077_v76_: D3
                    d_1077_v76_ = D3_DC8((p1) in ((d_1076_v75_).set(len(d_1072_v71_), p1)))
                    source14_ = d_1077_v76_
                    if source14_.is_DC8:
                        d_1078___mcc_h3_ = source14_.cf16
                        d_1079_cf16_ = d_1078___mcc_h3_
                        (globalState).f5 = (p0) and (p0)
                        d_1080_v77_: _dafny.Array
                        def lambda42_(d_1081_i6_):
                            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkmipcp"))

                        init21_ = lambda42_
                        nw183_ = _dafny.Array(None, 26)
                        for i0_21_ in range(nw183_.length(0)):
                            nw183_[i0_21_] = init21_(i0_21_)
                        d_1080_v77_ = nw183_
                        d_1080_v77_ = d_1080_v77_
                        d_1082_v78_: _dafny.MultiSet
                        d_1082_v78_ = _dafny.MultiSet([(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]])
                        index168_ = default__.safeIndex(986, (d_1063_v65_).length(0))
                        (d_1063_v65_)[index168_] = (0) - ((((d_1082_v78_)[p0] if (p0) in (d_1082_v78_) else len((d_1074_v73_).set(default__.safeIndex((d_1075_v74_).f22, len(d_1074_v73_)), d_1059_v61_)))) * (p1))
                        index169_ = default__.safeIndex(986, (d_1063_v65_).length(0))
                        (d_1063_v65_)[index169_] = (d_1075_v74_).f22
                        d_1083_v80_: _dafny.Map
                        d_1083_v80_ = _dafny.Map({d_1079_cf16_: (0) - ((d_1075_v74_).f22)})
                        def iife142_():
                            coll40_ = _dafny.Map()
                            compr_40_: _dafny.Map
                            for compr_40_ in (_dafny.Set({d_1083_v80_, _dafny.Map({d_1079_cf16_: (d_1075_v74_).f22}), _dafny.Map({d_1079_cf16_: (0) - (p1)}), d_1083_v80_})).Elements:
                                d_1084_v79_: _dafny.Map = compr_40_
                                if (d_1084_v79_) in (_dafny.Set({d_1083_v80_, _dafny.Map({d_1079_cf16_: (d_1075_v74_).f22}), _dafny.Map({d_1079_cf16_: (0) - (p1)}), d_1083_v80_})):
                                    coll40_[d_1084_v79_] = (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]
                            return _dafny.Map(coll40_)
                        (globalState).f18 = (((0) - (len(iife142_()
                        ))) + (len(d_1067_v66_))) - ((d_1075_v74_).f22)
                    elif source14_.is_DC7:
                        d_1085___mcc_h4_ = source14_.cf15
                        d_1086_cf15_ = d_1085___mcc_h4_
                        d_1087_v81_: _dafny.MultiSet
                        d_1087_v81_ = _dafny.MultiSet([len(d_1074_v73_), (d_1075_v74_).f22, (d_1075_v74_).f22])
                        index170_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        (d_1056_v60_)[index170_] = not((d_1087_v81_).issubset(_dafny.MultiSet([p1, 38])))
                        d_1088_v82_: _dafny.Map
                        d_1088_v82_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: d_1056_v60_})
                        d_1089_v83_: D9
                        d_1089_v83_ = D9_DC24(d_1088_v82_)
                        d_1090_v84_: _dafny.Map
                        d_1090_v84_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: d_1089_v83_})
                        d_1091_v85_: _dafny.Set
                        d_1091_v85_ = _dafny.Set({d_1067_v66_})
                        d_1092_v86_: D6
                        d_1092_v86_ = D6_DC17(d_1087_v81_, d_1091_v85_)
                        d_1093_v87_: _dafny.Map
                        d_1093_v87_ = _dafny.Map({d_1092_v86_: (not(not((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]))) == ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])})
                        index171_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        rhs142_ = d_1090_v84_
                        rhs143_ = ((d_1093_v87_)[d_1092_v86_] if (d_1092_v86_) in (d_1093_v87_) else (D0_DC0(p0)).cf0)
                        rhs144_ = p0
                        lhs140_ = d_1056_v60_
                        lhs141_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        lhs142_ = globalState
                        d_1090_v84_ = rhs142_
                        lhs140_[lhs141_] = rhs143_
                        lhs142_.f5 = rhs144_
                        d_1094_v88_: D10
                        d_1094_v88_ = D10_DC26(d_1072_v71_)
                        d_1094_v88_ = d_1094_v88_
                        d_1095_v89_: C1
                        nw184_ = C1()
                        nw184_.ctor__(-237, (self).f20)
                        d_1095_v89_ = nw184_
                        d_1096_v90_: D11
                        d_1096_v90_ = D11_DC30(d_1095_v89_)
                        d_1095_v89_ = (d_1096_v90_).cf58
                    elif True:
                        d_1097___mcc_h5_ = source14_.cf17
                        d_1098_cf17_ = d_1097___mcc_h5_
                        index172_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        (d_1056_v60_)[index172_] = ((d_1075_v74_).f22) > (((d_1075_v74_).f22) + (p1))
                        r1 = len(d_1072_v71_)
                        d_1099_v91_: D0
                        d_1099_v91_ = D0_DC1(d_1059_v61_, p1, _dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]}), p0, (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
                        d_1100_v92_: C2
                        nw185_ = C2()
                        nw185_.ctor__(default__.fm0(p1, d_1099_v91_, globalState))
                        d_1100_v92_ = nw185_
                        d_1101_v93_: _dafny.Seq
                        d_1101_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1075_v74_).f22])])
                        d_1102_v94_: _dafny.Map
                        d_1102_v94_ = _dafny.Map({(d_1101_v93_)[default__.safeIndex(p1, len(d_1101_v93_))]: (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                        d_1103_v95_: D6
                        d_1103_v95_ = D6_DC16(d_1102_v94_)
                        d_1103_v95_ = d_1103_v95_
                    d_1104_v96_: _dafny.Map
                    d_1104_v96_ = _dafny.Map({len(d_1076_v75_): d_1067_v66_})
                    if default__.fm2((d_1060_v62_) | ((d_1060_v62_).set(d_1059_v61_, default__.abs(p1))), len(d_1104_v96_), globalState):
                        index173_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        (d_1056_v60_)[index173_] = ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]) or ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
                        d_1105_v97_: _dafny.Array
                        nw186_ = _dafny.Array(_dafny.Map({}), 10)
                        d_1105_v97_ = nw186_
                        d_1106_v98_: _dafny.Map
                        d_1106_v98_ = _dafny.Map({(d_1075_v74_).f22: (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                        index174_ = default__.safeIndex(922, (d_1105_v97_).length(0))
                        (d_1105_v97_)[index174_] = (d_1106_v98_) | (d_1106_v98_)
                        d_1107_v99_: _dafny.MultiSet
                        d_1107_v99_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fshrx")), d_1074_v73_])
                        d_1108_v100_: _dafny.Seq
                        d_1108_v100_ = _dafny.SeqWithoutIsStrInference([d_1074_v73_])
                        index175_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        index176_ = default__.safeIndex(922, (d_1105_v97_).length(0))
                        rhs145_ = ((_dafny.MultiSet(d_1108_v100_)) - (d_1107_v99_)).ispropersubset(d_1107_v99_)
                        rhs146_ = (d_1106_v98_).set(((d_1075_v74_).f22) * ((d_1075_v74_).f22), ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]) and (p0))
                        rhs147_ = (d_1075_v74_).f22
                        rhs148_ = default__.fm1(globalState)
                        rhs149_ = d_1072_v71_
                        lhs143_ = d_1056_v60_
                        lhs144_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                        lhs145_ = d_1105_v97_
                        lhs146_ = default__.safeIndex(922, (d_1105_v97_).length(0))
                        lhs147_ = globalState
                        lhs143_[lhs144_] = rhs145_
                        lhs145_[lhs146_] = rhs146_
                        lhs147_.f17 = rhs147_
                        r2 = rhs148_
                        d_1072_v71_ = rhs149_
                        d_1109_v101_: D10
                        d_1109_v101_ = D10_DC26(_dafny.SeqWithoutIsStrInference([-267]))
                        d_1110_v102_: T0
                        nw187_ = C3()
                        nw187_.ctor__(default__.fm57((d_1075_v74_).f22, globalState))
                        d_1110_v102_ = nw187_
                        d_1111_v103_: _dafny.Seq
                        d_1111_v103_ = _dafny.SeqWithoutIsStrInference([d_1110_v102_])
                        d_1112_v104_: _dafny.Map
                        d_1112_v104_ = _dafny.Map({d_1109_v101_: len(d_1111_v103_)})
                        d_1112_v104_ = (d_1112_v104_).set(d_1109_v101_, default__.safeModuloInt((d_1075_v74_).f22, (d_1075_v74_).f22))
                        d_1113_v105_: _dafny.Array
                        def lambda43_(d_1114_v61_, d_1115_p1_, d_1116_v66_, d_1117_p0_, d_1118_v60_):
                            def lambda44_(d_1119_i7_):
                                return D0_DC1(d_1114_v61_, d_1115_p1_, d_1116_v66_, d_1117_p0_, not((d_1118_v60_)[default__.safeIndex(288, (d_1118_v60_).length(0))]))

                            return lambda44_

                        init22_ = lambda43_(d_1059_v61_, p1, d_1067_v66_, p0, d_1056_v60_)
                        nw188_ = _dafny.Array(None, 14)
                        for i0_22_ in range(nw188_.length(0)):
                            nw188_[i0_22_] = init22_(i0_22_)
                        d_1113_v105_ = nw188_
                        d_1120_v107_: D0
                        def iife143_():
                            coll41_ = _dafny.Map()
                            compr_41_: int
                            for compr_41_ in _dafny.IntegerRange(895, -363):
                                d_1121_v106_: int = compr_41_
                                if ((895) <= (d_1121_v106_)) and ((d_1121_v106_) < (-363)):
                                    coll41_[(d_1121_v106_) * ((0) - ((_dafny.MultiSet([(d_1075_v74_).f22])).cardinality))] = p0
                            return _dafny.Map(coll41_)
                        d_1120_v107_ = D0_DC1(d_1059_v61_, len(iife143_()
), d_1067_v66_, p0, p0)
                        index177_ = default__.safeIndex(89, (d_1113_v105_).length(0))
                        (d_1113_v105_)[index177_] = d_1120_v107_
                        index178_ = default__.safeIndex(89, (d_1113_v105_).length(0))
                        (d_1113_v105_)[index178_] = d_1120_v107_
                        (globalState).f18 = (0) - (len((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqptffvha"))).set(default__.safeIndex((0) - ((0) - ((d_1075_v74_).f22)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqptffvha")))), d_1059_v61_)) + (d_1074_v73_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwxdlibi")))))
                    elif True:
                        d_1122_v108_: D0
                        d_1122_v108_ = D0_DC1(d_1059_v61_, p1, d_1067_v66_, (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
                        d_1123_v109_: _dafny.Map
                        d_1123_v109_ = _dafny.Map({p0: _dafny.Set({len(_dafny.SeqWithoutIsStrInference([default__.fm0(22, d_1122_v108_, globalState), (d_1075_v74_).f22, (d_1075_v74_).f22])), p1, len(d_1074_v73_)})})
                        d_1124_v110_: _dafny.Map
                        d_1124_v110_ = _dafny.Map({d_1123_v109_: ((d_1075_v74_).f22) - (-705)})
                        d_1124_v110_ = (d_1124_v110_).set(d_1123_v109_, (p1) * ((d_1075_v74_).f22))
                        d_1125_v111_: C2
                        nw189_ = C2()
                        nw189_.ctor__((0) - ((d_1075_v74_).f22))
                        d_1125_v111_ = nw189_
                        d_1126_v112_: _dafny.MultiSet
                        d_1126_v112_ = _dafny.MultiSet([len(d_1074_v73_)])
                        (globalState).f6 = (d_1126_v112_).isdisjoint(_dafny.MultiSet([(d_1125_v111_).f22, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcsp"))), len(d_1061_v63_), (d_1075_v74_).f22, p1]))
                        d_1127_v113_: _dafny.Array
                        nw190_ = _dafny.Array(_dafny.Set({}), 29)
                        d_1127_v113_ = nw190_
                        d_1128_v114_: _dafny.MultiSet
                        d_1128_v114_ = _dafny.MultiSet([d_1062_v64_, _dafny.SeqWithoutIsStrInference([p0]), d_1062_v64_])
                        d_1129_v115_: D7
                        d_1129_v115_ = D7_DC20(p1, p0)
                        rhs150_ = (d_1128_v114_).isdisjoint(d_1128_v114_)
                        rhs151_ = d_1127_v113_
                        rhs152_ = (d_1129_v115_).cf37
                        rhs153_ = (d_1061_v63_) - ((d_1061_v63_).intersection(d_1061_v63_))
                        lhs148_ = globalState
                        lhs149_ = globalState
                        lhs148_.f5 = rhs150_
                        d_1127_v113_ = rhs151_
                        lhs149_.f6 = rhs152_
                        d_1061_v63_ = rhs153_
                        r2 = d_1059_v61_
                    (globalState).f5 = (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]
                    pass
            pass
        d_1130_v116_: _dafny.MultiSet
        d_1130_v116_ = _dafny.MultiSet([False, False])
        d_1131_v117_: _dafny.Seq
        d_1131_v117_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbis"))
        if ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]) == (((d_1130_v116_).set((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], default__.abs(len(d_1131_v117_)))).issubset(_dafny.MultiSet([True]))):
            (globalState).f18 = ((0) - (p1)) * (p1)
            r2 = d_1059_v61_
            if not ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]) or (not (p0) or ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])):
                (globalState).f18 = 859
                d_1059_v61_ = d_1059_v61_
                d_1131_v117_ = (d_1131_v117_) + ((_dafny.SeqWithoutIsStrInference([d_1059_v61_ for d_1132_i8_ in range(default__.abs(108))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_1059_v61_ for d_1133_i8_ in range(default__.abs(108))]))), d_1059_v61_))
                (globalState).f18 = p1
                d_1134_v118_: _dafny.MultiSet
                d_1134_v118_ = _dafny.MultiSet([(0) - (p1)])
                d_1135_v119_: _dafny.Map
                d_1135_v119_ = _dafny.Map({d_1134_v118_: p1})
                d_1136_v120_: _dafny.Seq
                d_1136_v120_ = _dafny.SeqWithoutIsStrInference([303])
                (globalState).f6 = not(not(not((D6_DC18(((d_1135_v119_)[_dafny.MultiSet((d_1136_v120_).set(default__.safeIndex(len(d_1131_v117_), len(d_1136_v120_)), p1))] if (_dafny.MultiSet((d_1136_v120_).set(default__.safeIndex(len(d_1131_v117_), len(d_1136_v120_)), p1))) in (d_1135_v119_) else p1), p1, d_1062_v64_, p0, p1)).cf33)))
            elif True:
                (globalState).f17 = p1
                d_1137_v121_: _dafny.Map
                d_1137_v121_ = _dafny.Map({(d_1130_v116_) | ((d_1130_v116_).set(p0, default__.abs(len(_dafny.SeqWithoutIsStrInference([p1]))))): (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                d_1138_v122_: _dafny.Seq
                d_1138_v122_ = _dafny.SeqWithoutIsStrInference([d_1130_v116_, d_1130_v116_, d_1130_v116_])
                d_1137_v121_ = (_dafny.Map({_dafny.MultiSet(d_1062_v64_): (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})) | (_dafny.Map({(d_1138_v122_)[default__.safeIndex(p1, len(d_1138_v122_))]: p0}))
                index179_ = default__.safeIndex(874, (d_1063_v65_).length(0))
                (d_1063_v65_)[index179_] = p1
                d_1139_v123_: _dafny.Map
                d_1139_v123_ = _dafny.Map({668: (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                d_1140_v124_: _dafny.Set
                d_1140_v124_ = _dafny.Set({p0})
                d_1141_v125_: _dafny.Map
                d_1141_v125_ = _dafny.Map({_dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], ((d_1139_v123_)[p1] if (p1) in (d_1139_v123_) else default__.fm2(_dafny.MultiSet(d_1131_v117_), p1, globalState)), (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]}): d_1140_v124_})
                index180_ = default__.safeIndex(874, (d_1063_v65_).length(0))
                (d_1063_v65_)[index180_] = len(d_1141_v125_)
                index181_ = default__.safeIndex(874, (d_1063_v65_).length(0))
                (d_1063_v65_)[index181_] = ((0) - ((d_1063_v65_)[default__.safeIndex(874, (d_1063_v65_).length(0))]) if (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))] else (0) - (default__.safeDivisionInt(p1, len(d_1140_v124_))))
                (globalState).f18 = (d_1063_v65_)[default__.safeIndex(874, (d_1063_v65_).length(0))]
            if (d_1062_v64_)[default__.safeIndex(p1, len(d_1062_v64_))]:
                d_1142_v126_: D0
                d_1142_v126_ = D0_DC1(_dafny.CodePoint('g'), len(_dafny.SeqWithoutIsStrInference([p0])), _dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]}), p0, default__.fm2(d_1060_v62_, p1, globalState))
                (globalState).f17 = default__.fm0(p1, d_1142_v126_, globalState)
                d_1143_v127_: D14
                d_1143_v127_ = D14_DC39(d_1056_v60_)
                d_1144_v128_: _dafny.Seq
                d_1144_v128_ = _dafny.SeqWithoutIsStrInference([d_1056_v60_, d_1056_v60_, d_1056_v60_])
                d_1145_v129_: _dafny.MultiSet
                d_1145_v129_ = _dafny.MultiSet([p1])
                d_1146_v130_: _dafny.Map
                d_1146_v130_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: d_1056_v60_})
                d_1147_v131_: _dafny.Array
                nw191_ = _dafny.Array(None, 15)
                nw191_[int(0)] = (d_1143_v127_).cf72
                nw191_[int(1)] = d_1056_v60_
                nw191_[int(2)] = d_1056_v60_
                nw191_[int(3)] = d_1056_v60_
                nw191_[int(4)] = (d_1144_v128_)[default__.safeIndex(458, len(d_1144_v128_))]
                nw191_[int(5)] = d_1056_v60_
                nw191_[int(6)] = d_1056_v60_
                nw191_[int(7)] = d_1056_v60_
                nw191_[int(8)] = d_1056_v60_
                nw191_[int(9)] = d_1056_v60_
                nw191_[int(10)] = (d_1144_v128_)[default__.safeIndex((d_1145_v129_).cardinality, len(d_1144_v128_))]
                nw191_[int(11)] = d_1056_v60_
                nw191_[int(12)] = d_1056_v60_
                nw191_[int(13)] = d_1056_v60_
                nw191_[int(14)] = ((d_1146_v130_)[not(p0)] if (not(p0)) in (d_1146_v130_) else d_1056_v60_)
                d_1147_v131_ = nw191_
                index182_ = default__.safeIndex(763, (d_1147_v131_).length(0))
                (d_1147_v131_)[index182_] = d_1056_v60_
                index183_ = default__.safeIndex(763, (d_1147_v131_).length(0))
                def lambda45_(d_1148_p0_):
                    def lambda46_(d_1149_i9_):
                        return d_1148_p0_

                    return lambda46_

                init23_ = lambda45_(p0)
                nw192_ = _dafny.Array(None, 27)
                for i0_23_ in range(nw192_.length(0)):
                    nw192_[i0_23_] = init23_(i0_23_)
                (d_1147_v131_)[index183_] = nw192_
                d_1150_v132_: _dafny.Array
                nw193_ = _dafny.Array(None, 14)
                d_1150_v132_ = nw193_
                d_1151_v133_: T0
                nw194_ = C3()
                nw194_.ctor__((self).f20)
                d_1151_v133_ = nw194_
                index184_ = default__.safeIndex(382, (d_1150_v132_).length(0))
                (d_1150_v132_)[index184_] = d_1151_v133_
                d_1152_v134_: D15
                d_1152_v134_ = D15_DC42(d_1151_v133_)
                index185_ = default__.safeIndex(382, (d_1150_v132_).length(0))
                (d_1150_v132_)[index185_] = (d_1151_v133_ if False else (d_1152_v134_).cf81)
                (globalState).f6 = p0
                (globalState).f5 = (True) or ((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
            elif True:
                (globalState).f6 = not((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
                d_1153_v135_: _dafny.Map
                d_1153_v135_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: p1})
                d_1154_v137_: _dafny.Map
                d_1154_v137_ = _dafny.Map({p0: p0})
                d_1155_v138_: _dafny.MultiSet
                d_1155_v138_ = _dafny.MultiSet([d_1154_v137_])
                d_1156_v139_: _dafny.MultiSet
                def iife144_():
                    coll42_ = _dafny.Map()
                    compr_42_: _dafny.Map
                    for compr_42_ in (d_1155_v138_).Elements:
                        d_1157_v136_: _dafny.Map = compr_42_
                        if (d_1157_v136_) in (d_1155_v138_):
                            coll42_[d_1157_v136_] = p1
                    return _dafny.Map(coll42_)
                d_1156_v139_ = _dafny.MultiSet([(p1) + (p1), (p1) - (p1), p1, default__.safeModuloInt(default__.fm0(((d_1153_v135_)[p0] if (p0) in (d_1153_v135_) else p1), D0_DC1(d_1059_v61_, len(iife144_()
), default__.fm52(p1, d_1131_v117_, p0, p0, globalState), p0, default__.fm2(d_1060_v62_, p1, globalState)), globalState), p1)])
                d_1158_v140_: _dafny.Set
                d_1158_v140_ = _dafny.Set({p0})
                d_1159_v141_: D0
                d_1159_v141_ = D0_DC1(d_1059_v61_, len(d_1153_v135_), d_1158_v140_, False, p0)
                d_1160_v142_: D0
                d_1160_v142_ = D0_DC1(_dafny.CodePoint('y'), default__.fm0(p1, d_1159_v141_, globalState), d_1158_v140_, p0, p0)
                d_1161_v143_: _dafny.Map
                d_1161_v143_ = _dafny.Map({p1: default__.fm0(p1, d_1159_v141_, globalState)})
                index186_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                rhs154_ = default__.safeDivisionInt(len(d_1131_v117_), p1)
                rhs155_ = False
                rhs156_ = ((d_1156_v139_) - (d_1156_v139_)) - (default__.fm56(p1, (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], p0, len(d_1161_v143_), globalState))
                lhs150_ = globalState
                lhs151_ = d_1056_v60_
                lhs152_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                lhs150_.f17 = rhs154_
                lhs151_[lhs152_] = rhs155_
                d_1156_v139_ = rhs156_
                d_1162_v144_: C2
                nw195_ = C2()
                nw195_.ctor__((0) - (len(d_1131_v117_)))
                d_1162_v144_ = nw195_
                d_1163_v145_: _dafny.Seq
                d_1163_v145_ = _dafny.SeqWithoutIsStrInference([p1, (d_1162_v144_).f22])
                d_1164_v146_: _dafny.Map
                d_1164_v146_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: (d_1163_v145_ if False else _dafny.SeqWithoutIsStrInference([(d_1162_v144_).f22 for d_1165_i10_ in range(default__.abs(751))]))})
                (globalState).f17 = len(d_1164_v146_)
                index187_ = default__.safeIndex(805, (d_1063_v65_).length(0))
                (d_1063_v65_)[index187_] = default__.safeDivisionInt(p1, len(d_1061_v63_))
                index188_ = default__.safeIndex(805, (d_1063_v65_).length(0))
                (d_1063_v65_)[index188_] = (d_1162_v144_).f22
            d_1166_v148_: _dafny.Map
            d_1166_v148_ = _dafny.Map({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]: (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
            d_1167_v149_: _dafny.Map
            def iife145_():
                coll43_ = _dafny.Map()
                compr_43_: _dafny.Map
                for compr_43_ in (_dafny.SeqWithoutIsStrInference([d_1166_v148_])).Elements:
                    d_1168_v147_: _dafny.Map = compr_43_
                    if (d_1168_v147_) in (_dafny.SeqWithoutIsStrInference([d_1166_v148_])):
                        coll43_[d_1168_v147_] = p0
                return _dafny.Map(coll43_)
            d_1167_v149_ = _dafny.Map({default__.safeDivisionInt(len(iife145_()
            ), 604): (0) - (p1)})
            d_1167_v149_ = (d_1167_v149_).set(len(d_1131_v117_), p1)
        elif True:
            d_1169_v150_: D3
            d_1169_v150_ = D3_DC8(False)
            index189_ = default__.safeIndex(253, (d_1063_v65_).length(0))
            (d_1063_v65_)[index189_] = p1
            index190_ = default__.safeIndex(253, (d_1063_v65_).length(0))
            rhs157_ = (p1 if p0 else default__.safeModuloInt(701, p1))
            rhs158_ = d_1063_v65_
            rhs159_ = d_1169_v150_
            rhs160_ = default__.safeModuloInt(p1, p1)
            rhs161_ = (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "llgtgcnr")))) + ((p1) * (p1)))
            lhs153_ = globalState
            lhs154_ = d_1063_v65_
            lhs155_ = default__.safeIndex(253, (d_1063_v65_).length(0))
            lhs156_ = globalState
            lhs153_.f17 = rhs157_
            d_1063_v65_ = rhs158_
            d_1169_v150_ = rhs159_
            lhs154_[lhs155_] = rhs160_
            lhs156_.f18 = rhs161_
            d_1170_v151_: _dafny.Array
            def lambda47_(d_1171_i11_):
                return _dafny.CodePoint('j')

            init24_ = lambda47_
            nw196_ = _dafny.Array(None, 21)
            for i0_24_ in range(nw196_.length(0)):
                nw196_[i0_24_] = init24_(i0_24_)
            d_1170_v151_ = nw196_
            d_1172_v152_: _dafny.MultiSet
            d_1172_v152_ = _dafny.MultiSet([d_1170_v151_, d_1170_v151_])
            (globalState).f5 = (d_1172_v152_).issubset(d_1172_v152_)
            (globalState).f17 = p1
            index191_ = default__.safeIndex(253, (d_1063_v65_).length(0))
            (d_1063_v65_)[index191_] = (d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))]
            if ((self).f20).cf0:
                (globalState).f18 = (((0) - ((d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))])) + ((d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))])) * ((d_1130_v116_).cardinality)
                d_1173_v153_: _dafny.Array
                def lambda48_(d_1174_v61_, d_1175_p1_, d_1176_v60_):
                    def lambda49_(d_1177_i12_):
                        return (D0_DC1(d_1174_v61_, d_1175_p1_, _dafny.Set({False}), (d_1176_v60_)[default__.safeIndex(288, (d_1176_v60_).length(0))], not((d_1176_v60_)[default__.safeIndex(288, (d_1176_v60_).length(0))]))).cf3

                    return lambda49_

                init25_ = lambda48_(d_1059_v61_, p1, d_1056_v60_)
                nw197_ = _dafny.Array(None, 9)
                for i0_25_ in range(nw197_.length(0)):
                    nw197_[i0_25_] = init25_(i0_25_)
                d_1173_v153_ = nw197_
                d_1178_v154_: _dafny.Set
                d_1178_v154_ = _dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                index192_ = default__.safeIndex(977, (d_1173_v153_).length(0))
                (d_1173_v153_)[index192_] = d_1178_v154_
                index193_ = default__.safeIndex(977, (d_1173_v153_).length(0))
                (d_1173_v153_)[index193_] = (d_1178_v154_ if p0 else d_1178_v154_)
                r1 = (0) - ((len(d_1061_v63_)) - (p1))
                (globalState).f5 = not((d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
                (globalState).f17 = p1
            elif True:
                d_1179_v155_: _dafny.Array
                nw198_ = _dafny.Array(int(0), 20)
                d_1179_v155_ = nw198_
                d_1180_v156_: _dafny.Array
                nw199_ = _dafny.Array(None, 1)
                nw199_[int(0)] = d_1179_v155_
                d_1180_v156_ = nw199_
                d_1180_v156_ = d_1180_v156_
                d_1181_v157_: _dafny.Map
                d_1181_v157_ = _dafny.Map({p1: True})
                (globalState).f17 = default__.safeModuloInt(len(d_1062_v64_), len(d_1181_v157_))
                (globalState).f5 = p0
                index194_ = default__.safeIndex(288, (d_1056_v60_).length(0))
                (d_1056_v60_)[index194_] = p0
                d_1182_v158_: _dafny.Map
                d_1182_v158_ = _dafny.Map({p0: (0) - (p1)})
                d_1183_v159_: _dafny.MultiSet
                d_1183_v159_ = _dafny.MultiSet([p1, p1])
                d_1184_v160_: _dafny.Set
                d_1184_v160_ = _dafny.Set({(d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))], False, (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))]})
                d_1185_v161_: D0
                d_1185_v161_ = D0_DC1(default__.fm1(globalState), p1, d_1184_v160_, p0, (d_1062_v64_)[default__.safeIndex((d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))], len(d_1062_v64_))])
                d_1182_v158_ = (d_1182_v158_).set(default__.fm2((d_1060_v62_).set(_dafny.CodePoint('e'), default__.abs((d_1183_v159_).cardinality)), default__.fm0((d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))], d_1185_v161_, globalState), globalState), ((d_1063_v65_)[default__.safeIndex(253, (d_1063_v65_).length(0))]) * (len(d_1184_v160_)))
        d_1186_v162_: _dafny.Set
        d_1186_v162_ = _dafny.Set({p0, p0})
        d_1187_v163_: D0
        d_1187_v163_ = D0_DC1(d_1059_v61_, p1, d_1186_v162_, p0, (d_1056_v60_)[default__.safeIndex(288, (d_1056_v60_).length(0))])
        d_1188_v164_: D6
        d_1188_v164_ = D6_DC18(496, p1, (d_1062_v64_).set(default__.safeIndex(p1, len(d_1062_v64_)), not(p0)), p0, default__.fm0(p1, d_1187_v163_, globalState))
        r0 = _dafny.MultiSet(((d_1188_v164_).cf32).set(default__.safeIndex(len(d_1131_v117_), len((d_1188_v164_).cf32)), default__.fm2(d_1060_v62_, p1, globalState)))
        r1 = (p1) + (p1)
        r2 = d_1059_v61_
        return r0, r1, r2

    def m12(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: _dafny.Set = _dafny.Set({})
        d_1189_v0_: bool
        d_1189_v0_ = False
        d_1190_v1_: int
        d_1190_v1_ = 631
        d_1191_v2_: _dafny.Seq
        d_1191_v2_ = _dafny.SeqWithoutIsStrInference([d_1190_v1_, d_1190_v1_])
        d_1192_v3_: _dafny.Seq
        d_1192_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qorehxlli"))
        d_1193_v4_: D13
        d_1193_v4_ = D13_DC36((d_1191_v2_)[default__.safeIndex(len(d_1192_v3_), len(d_1191_v2_))])
        d_1194_v5_: str
        d_1194_v5_ = _dafny.CodePoint('f')
        pat_let_tv35_ = d_1192_v3_
        pat_let_tv36_ = d_1191_v2_
        pat_let_tv37_ = d_1192_v3_
        pat_let_tv38_ = d_1194_v5_
        pat_let_tv39_ = d_1189_v0_
        pat_let_tv40_ = d_1190_v1_
        def iife146_(_pat_let51_0):
            def iife147_(d_1195_dt__update__tmp_h0_):
                def iife148_(_pat_let52_0):
                    def iife149_(d_1196_dt__update_hcf68_h0_):
                        return D13_DC36(d_1196_dt__update_hcf68_h0_)
                    return iife149_(_pat_let52_0)
                return iife148_((len((pat_let_tv35_).set(default__.safeIndex(len(pat_let_tv36_), len(pat_let_tv37_)), pat_let_tv38_)) if pat_let_tv39_ else pat_let_tv40_))
            return iife147_(_pat_let51_0)
        source15_ = iife146_((d_1193_v4_ if d_1189_v0_ else d_1193_v4_))
        if source15_.is_DC36:
            d_1197___mcc_h0_ = source15_.cf68
            d_1198_cf68_ = d_1197___mcc_h0_
            d_1199_v6_: _dafny.Array
            def lambda50_(d_1200_v0_, d_1201_v1_):
                def lambda51_(d_1202_i0_):
                    return default__.safeModuloInt(d_1202_i0_, len(_dafny.Map({d_1200_v0_: d_1201_v1_})))

                return lambda51_

            init26_ = lambda50_(d_1189_v0_, d_1190_v1_)
            nw200_ = _dafny.Array(None, 1)
            for i0_26_ in range(nw200_.length(0)):
                nw200_[i0_26_] = init26_(i0_26_)
            d_1199_v6_ = nw200_
            index195_ = default__.safeIndex(266, (d_1199_v6_).length(0))
            (d_1199_v6_)[index195_] = d_1190_v1_
            index196_ = default__.safeIndex(266, (d_1199_v6_).length(0))
            (d_1199_v6_)[index196_] = d_1198_cf68_
            (globalState).f5 = d_1189_v0_
            d_1203_v7_: _dafny.MultiSet
            d_1203_v7_ = _dafny.MultiSet([d_1189_v0_, False, True])
            d_1203_v7_ = (d_1203_v7_) - ((d_1203_v7_) - (d_1203_v7_))
            (globalState).f5 = d_1189_v0_
        elif source15_.is_DC37:
            d_1204___mcc_h1_ = source15_.cf69
            d_1205___mcc_h2_ = source15_.cf70
            d_1206___mcc_h3_ = source15_.cf71
            d_1207_cf71_ = d_1206___mcc_h3_
            d_1208_cf70_ = d_1205___mcc_h2_
            d_1209_cf69_ = d_1204___mcc_h1_
            (globalState).f5 = d_1189_v0_
            (globalState).f6 = d_1189_v0_
            if True:
                d_1194_v5_ = d_1194_v5_
                (globalState).f6 = d_1189_v0_
                d_1210_v8_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.Map({}), 25)
                d_1210_v8_ = nw201_
                d_1211_v9_: _dafny.Map
                d_1211_v9_ = _dafny.Map({d_1209_cf69_: d_1189_v0_})
                d_1212_v10_: _dafny.Map
                d_1212_v10_ = _dafny.Map({((d_1211_v9_)[d_1209_cf69_] if (d_1209_cf69_) in (d_1211_v9_) else d_1189_v0_): d_1192_v3_})
                index197_ = default__.safeIndex(710, (d_1210_v8_).length(0))
                (d_1210_v8_)[index197_] = (d_1212_v10_) | (_dafny.Map({d_1189_v0_: d_1192_v3_}))
                d_1213_v11_: _dafny.Array
                def lambda52_(d_1214_v0_):
                    def lambda53_(d_1215_i1_):
                        return _dafny.Set({D5_DC15(D5_DC14(not(True))), D5_DC15(D5_DC14(True)), D5_DC15(D5_DC14(d_1214_v0_)), D5_DC15(D5_DC14(d_1214_v0_)), D5_DC15(D5_DC14(d_1214_v0_))})

                    return lambda53_

                init27_ = lambda52_(d_1189_v0_)
                nw202_ = _dafny.Array(None, 8)
                for i0_27_ in range(nw202_.length(0)):
                    nw202_[i0_27_] = init27_(i0_27_)
                d_1213_v11_ = nw202_
                d_1216_v12_: D0
                d_1216_v12_ = D0_DC1(d_1194_v5_, d_1209_cf69_, _dafny.Set({d_1189_v0_}), d_1189_v0_, d_1189_v0_)
                d_1217_v13_: D5
                d_1217_v13_ = D5_DC13(d_1211_v9_)
                d_1218_v14_: D5
                d_1218_v14_ = D5_DC15(d_1217_v13_)
                d_1219_v15_: _dafny.Set
                d_1219_v15_ = _dafny.Set({D5_DC15(d_1217_v13_), d_1218_v14_})
                index198_ = default__.safeIndex(70, (d_1213_v11_).length(0))
                (d_1213_v11_)[index198_] = (_dafny.Set({d_1218_v14_, d_1218_v14_, d_1218_v14_, d_1218_v14_, d_1218_v14_}) if (d_1216_v12_).cf5 else d_1219_v15_)
                d_1220_v16_: _dafny.Seq
                d_1220_v16_ = _dafny.SeqWithoutIsStrInference([d_1192_v3_])
                d_1221_v17_: _dafny.Seq
                d_1221_v17_ = _dafny.SeqWithoutIsStrInference([d_1218_v14_, d_1218_v14_, d_1218_v14_])
                d_1222_v19_: _dafny.Map
                def iife150_():
                    coll44_ = _dafny.Set()
                    compr_44_: D5
                    for compr_44_ in (d_1221_v17_).Elements:
                        d_1223_v18_: D5 = compr_44_
                        if (d_1223_v18_) in (d_1221_v17_):
                            coll44_ = coll44_.union(_dafny.Set([d_1223_v18_]))
                    return _dafny.Set(coll44_)
                d_1222_v19_ = _dafny.Map({d_1189_v0_: iife150_()
                })
                index199_ = default__.safeIndex(710, (d_1210_v8_).length(0))
                index200_ = default__.safeIndex(70, (d_1213_v11_).length(0))
                rhs162_ = ((0) - ((0) - (d_1190_v1_))) - (len(d_1192_v3_))
                rhs163_ = _dafny.Map({d_1189_v0_: (d_1220_v16_)[default__.safeIndex(d_1209_cf69_, len(d_1220_v16_))]})
                rhs164_ = (((d_1222_v19_)[d_1189_v0_] if (d_1189_v0_) in (d_1222_v19_) else d_1219_v15_)) - ((default__.fm58(d_1189_v0_, d_1190_v1_, _dafny.CodePoint('b'), globalState)).cf87)
                lhs157_ = globalState
                lhs158_ = d_1210_v8_
                lhs159_ = default__.safeIndex(710, (d_1210_v8_).length(0))
                lhs160_ = d_1213_v11_
                lhs161_ = default__.safeIndex(70, (d_1213_v11_).length(0))
                lhs157_.f17 = rhs162_
                lhs158_[lhs159_] = rhs163_
                lhs160_[lhs161_] = rhs164_
                d_1190_v1_ = d_1190_v1_
                (globalState).f18 = d_1207_cf71_
            elif True:
                d_1224_v20_: _dafny.Seq
                d_1224_v20_ = _dafny.SeqWithoutIsStrInference([d_1189_v0_, True])
                d_1225_v21_: _dafny.MultiSet
                d_1225_v21_ = _dafny.MultiSet([len(_dafny.Map({d_1224_v20_: d_1189_v0_}))])
                (globalState).f6 = (d_1189_v0_) == (((d_1225_v21_).set(len(d_1224_v20_), default__.abs(d_1190_v1_))).isdisjoint(default__.fm56(len(d_1192_v3_), d_1189_v0_, d_1189_v0_, d_1209_cf69_, globalState)))
                d_1226_v22_: D10
                d_1226_v22_ = D10_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoox")), d_1209_cf69_, not(d_1189_v0_), d_1189_v0_, d_1189_v0_)
                (globalState).f5 = (d_1226_v22_).cf56
                d_1227_v23_: C1
                nw203_ = C1()
                nw203_.ctor__(d_1208_cf70_, (self).f20)
                d_1227_v23_ = nw203_
                d_1228_v24_: _dafny.Seq
                d_1228_v24_ = _dafny.SeqWithoutIsStrInference([d_1224_v20_])
                d_1229_v25_: _dafny.Array
                nw204_ = _dafny.Array(None, 29)
                nw204_[int(0)] = d_1224_v20_
                nw204_[int(1)] = d_1224_v20_
                nw204_[int(2)] = d_1224_v20_
                nw204_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1189_v0_])
                nw204_[int(4)] = d_1224_v20_
                nw204_[int(5)] = d_1224_v20_
                nw204_[int(6)] = _dafny.SeqWithoutIsStrInference([not(d_1189_v0_)])
                nw204_[int(7)] = d_1224_v20_
                nw204_[int(8)] = d_1224_v20_
                nw204_[int(9)] = d_1224_v20_
                nw204_[int(10)] = d_1224_v20_
                nw204_[int(11)] = d_1224_v20_
                nw204_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1189_v0_, d_1189_v0_])
                nw204_[int(13)] = d_1224_v20_
                nw204_[int(14)] = _dafny.SeqWithoutIsStrInference([d_1189_v0_])
                nw204_[int(15)] = d_1224_v20_
                nw204_[int(16)] = d_1224_v20_
                nw204_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1189_v0_])
                nw204_[int(18)] = d_1224_v20_
                nw204_[int(19)] = d_1224_v20_
                nw204_[int(20)] = d_1224_v20_
                nw204_[int(21)] = d_1224_v20_
                nw204_[int(22)] = d_1224_v20_
                nw204_[int(23)] = _dafny.SeqWithoutIsStrInference([d_1189_v0_])
                nw204_[int(24)] = d_1224_v20_
                nw204_[int(25)] = d_1224_v20_
                nw204_[int(26)] = (d_1228_v24_)[default__.safeIndex(d_1208_cf70_, len(d_1228_v24_))]
                nw204_[int(27)] = d_1224_v20_
                nw204_[int(28)] = d_1224_v20_
                d_1229_v25_ = nw204_
                d_1230_v26_: D3
                d_1230_v26_ = D3_DC8(d_1189_v0_)
                d_1231_v27_: D15
                d_1231_v27_ = D15_DC43(d_1229_v25_, True, d_1209_cf69_, d_1189_v0_, d_1230_v26_)
                d_1232_v28_: D2
                d_1232_v28_ = D2_DC4((d_1225_v21_).cardinality, d_1189_v0_, True, (d_1231_v27_).cf85)
                (globalState).f5 = (d_1232_v28_).cf10
                d_1233_v29_: D2
                d_1233_v29_ = D2_DC3(default__.fm46(d_1207_cf71_, d_1189_v0_, d_1189_v0_, globalState))
                d_1234_v30_: _dafny.Seq
                d_1234_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1235_i2_ in range(default__.abs(602))])])
                d_1236_v31_: _dafny.Set
                d_1236_v31_ = _dafny.Set({((d_1233_v29_).cf7) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elscgnk"))), d_1189_v0_, (d_1192_v3_) not in (d_1234_v30_)})
                d_1236_v31_ = d_1236_v31_
            d_1237_v32_: T0
            nw205_ = C3()
            nw205_.ctor__((self).f20)
            d_1237_v32_ = nw205_
            d_1237_v32_ = d_1237_v32_
        elif source15_.is_DC38:
            d_1238_v33_: _dafny.MultiSet
            d_1238_v33_ = _dafny.MultiSet([d_1194_v5_, d_1194_v5_, (d_1192_v3_)[default__.safeIndex((0) - (d_1190_v1_), len(d_1192_v3_))]])
            rhs165_ = default__.fm2(d_1238_v33_, d_1190_v1_, globalState)
            rhs166_ = False
            lhs162_ = globalState
            lhs162_.f6 = rhs165_
            r1 = rhs166_
            d_1239_v34_: _dafny.Array
            def lambda54_(d_1240_v0_):
                def lambda55_(d_1241_i3_):
                    return (d_1241_i3_) * (len(_dafny.Map({len(_dafny.Map({d_1240_v0_: d_1240_v0_})): d_1240_v0_})))

                return lambda55_

            init28_ = lambda54_(d_1189_v0_)
            nw206_ = _dafny.Array(None, 7)
            for i0_28_ in range(nw206_.length(0)):
                nw206_[i0_28_] = init28_(i0_28_)
            d_1239_v34_ = nw206_
            d_1242_v35_: D0
            d_1242_v35_ = D0_DC1(default__.fm1(globalState), d_1190_v1_, _dafny.Set({d_1189_v0_}), d_1189_v0_, d_1189_v0_)
            index201_ = default__.safeIndex(432, (d_1239_v34_).length(0))
            (d_1239_v34_)[index201_] = default__.fm0(636, d_1242_v35_, globalState)
            d_1243_v36_: D2
            d_1243_v36_ = D2_DC4(d_1190_v1_, d_1189_v0_, d_1189_v0_, d_1189_v0_)
            index202_ = default__.safeIndex(432, (d_1239_v34_).length(0))
            (d_1239_v34_)[index202_] = (d_1243_v36_).cf8
            (globalState).f15 = (d_1192_v3_) + (_dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1244_i4_ in range(default__.abs(-677))]))
            d_1245_v37_: _dafny.Map
            d_1245_v37_ = _dafny.Map({(0) - ((d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]): d_1189_v0_})
            d_1246_v39_: _dafny.Map
            d_1246_v39_ = _dafny.Map({d_1190_v1_: d_1245_v37_})
            d_1247_v41_: _dafny.MultiSet
            d_1247_v41_ = _dafny.MultiSet([d_1190_v1_])
            d_1248_v42_: _dafny.Seq
            d_1248_v42_ = _dafny.SeqWithoutIsStrInference([d_1245_v37_])
            d_1249_v43_: _dafny.Map
            d_1249_v43_ = _dafny.Map({d_1192_v3_: d_1245_v37_})
            d_1250_v46_: _dafny.Set
            d_1250_v46_ = _dafny.Set({526, (d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]})
            d_1251_v47_: _dafny.Array
            nw207_ = _dafny.Array(None, 26)
            nw207_[int(0)] = d_1245_v37_
            nw207_[int(1)] = d_1245_v37_
            nw207_[int(2)] = (d_1245_v37_) | (d_1245_v37_)
            nw207_[int(3)] = (d_1245_v37_) | (d_1245_v37_)
            def iife151_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(82, 725):
                    d_1252_v38_: int = compr_45_
                    if ((82) <= (d_1252_v38_)) and ((d_1252_v38_) < (725)):
                        coll45_[default__.safeDivisionInt(d_1252_v38_, (d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))])] = d_1189_v0_
                return _dafny.Map(coll45_)
            nw207_[int(4)] = iife151_()
            
            nw207_[int(5)] = d_1245_v37_
            nw207_[int(6)] = d_1245_v37_
            nw207_[int(7)] = d_1245_v37_
            nw207_[int(8)] = d_1245_v37_
            nw207_[int(9)] = _dafny.Map({(0) - ((d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]): d_1189_v0_})
            nw207_[int(10)] = (d_1245_v37_) | (((d_1246_v39_)[len(d_1245_v37_)] if (len(d_1245_v37_)) in (d_1246_v39_) else d_1245_v37_))
            nw207_[int(11)] = d_1245_v37_
            def iife152_():
                coll46_ = _dafny.Map()
                compr_46_: int
                for compr_46_ in (d_1191_v2_).Elements:
                    d_1253_v40_: int = compr_46_
                    if (d_1253_v40_) in (d_1191_v2_):
                        coll46_[(d_1253_v40_) * ((d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))])] = d_1189_v0_
                return _dafny.Map(coll46_)
            nw207_[int(12)] = iife152_()
            
            nw207_[int(13)] = d_1245_v37_
            nw207_[int(14)] = (d_1245_v37_).set((d_1191_v2_)[default__.safeIndex((0) - ((d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]), len(d_1191_v2_))], d_1189_v0_)
            nw207_[int(15)] = _dafny.Map({(d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]: d_1189_v0_})
            nw207_[int(16)] = default__.fm59((d_1247_v41_).cardinality, globalState)
            nw207_[int(17)] = d_1245_v37_
            nw207_[int(18)] = d_1245_v37_
            nw207_[int(19)] = d_1245_v37_
            nw207_[int(20)] = (d_1248_v42_)[default__.safeIndex(len(_dafny.Map({(d_1239_v34_)[default__.safeIndex(432, (d_1239_v34_).length(0))]: d_1189_v0_})), len(d_1248_v42_))]
            nw207_[int(21)] = ((d_1249_v43_)[d_1192_v3_] if (d_1192_v3_) in (d_1249_v43_) else d_1245_v37_)
            nw207_[int(22)] = (d_1245_v37_) | (d_1245_v37_)
            def iife153_():
                def iife155_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (d_1250_v46_).Elements:
                        d_1254_v45_: int = compr_49_
                        if (d_1254_v45_) in (d_1250_v46_):
                            coll49_[(d_1254_v45_) + (d_1190_v1_)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1189_v0_, (d_1243_v36_).cf10])) for d_1255_i5_ in range(default__.abs(721))])
                    return _dafny.Map(coll49_)
                coll47_ = _dafny.Map()
                def iife154_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in (d_1250_v46_).Elements:
                        d_1254_v45_: int = compr_48_
                        if (d_1254_v45_) in (d_1250_v46_):
                            coll48_[(d_1254_v45_) + (d_1190_v1_)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1189_v0_, (d_1243_v36_).cf10])) for d_1255_i5_ in range(default__.abs(721))])
                    return _dafny.Map(coll48_)
                compr_47_: int
                for compr_47_ in (iife154_()
                ).keys.Elements:
                    d_1256_v44_: int = compr_47_
                    if (d_1256_v44_) in (iife155_()
                    ):
                        coll47_[(d_1256_v44_) - (d_1190_v1_)] = False
                return _dafny.Map(coll47_)
            nw207_[int(23)] = (d_1245_v37_) | (iife153_()
            )
            nw207_[int(24)] = default__.fm59(d_1190_v1_, globalState)
            nw207_[int(25)] = (d_1245_v37_) | (d_1245_v37_)
            d_1251_v47_ = nw207_
            index203_ = default__.safeIndex(681, (d_1251_v47_).length(0))
            (d_1251_v47_)[index203_] = default__.fm59(len(d_1191_v2_), globalState)
            index204_ = default__.safeIndex(681, (d_1251_v47_).length(0))
            (d_1251_v47_)[index204_] = d_1245_v37_
        elif True:
            d_1257___mcc_h4_ = source15_.cf67
            d_1258_cf67_ = d_1257___mcc_h4_
            d_1259_v48_: _dafny.Map
            d_1259_v48_ = _dafny.Map({d_1190_v1_: d_1189_v0_})
            d_1260_v49_: _dafny.Map
            d_1260_v49_ = _dafny.Map({(0) - (d_1190_v1_): d_1194_v5_})
            d_1261_v50_: _dafny.MultiSet
            d_1261_v50_ = _dafny.MultiSet([((d_1260_v49_)[d_1190_v1_] if (d_1190_v1_) in (d_1260_v49_) else d_1194_v5_), d_1194_v5_, d_1194_v5_, d_1194_v5_])
            d_1262_v51_: _dafny.Seq
            d_1262_v51_ = _dafny.SeqWithoutIsStrInference([d_1189_v0_, default__.fm2(d_1261_v50_, 250, globalState)])
            d_1263_v52_: _dafny.Array
            nw208_ = _dafny.Array(None, 14)
            nw208_[int(0)] = (d_1190_v1_) >= (d_1190_v1_)
            nw208_[int(1)] = (default__.fm60(239, globalState)).isdisjoint(d_1258_cf67_)
            nw208_[int(2)] = not((d_1189_v0_ if ((d_1259_v48_)[len(d_1262_v51_)] if (len(d_1262_v51_)) in (d_1259_v48_) else False) else d_1189_v0_))
            nw208_[int(3)] = (len(_dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1264_i6_ in range(default__.abs(-741))]))) not in (d_1258_cf67_)
            nw208_[int(4)] = d_1189_v0_
            nw208_[int(5)] = d_1189_v0_
            nw208_[int(6)] = not(True)
            nw208_[int(7)] = default__.fm2(d_1261_v50_, d_1190_v1_, globalState)
            nw208_[int(8)] = d_1189_v0_
            nw208_[int(9)] = not(d_1189_v0_)
            nw208_[int(10)] = not(d_1189_v0_)
            nw208_[int(11)] = d_1189_v0_
            nw208_[int(12)] = (d_1189_v0_) and (d_1189_v0_)
            nw208_[int(13)] = d_1189_v0_
            d_1263_v52_ = nw208_
            index205_ = default__.safeIndex(733, (d_1263_v52_).length(0))
            (d_1263_v52_)[index205_] = (d_1189_v0_) in ((d_1262_v51_).set(default__.safeIndex(d_1190_v1_, len(d_1262_v51_)), d_1189_v0_))
            index206_ = default__.safeIndex(733, (d_1263_v52_).length(0))
            (d_1263_v52_)[index206_] = not(d_1189_v0_)
            d_1265_v53_: _dafny.Set
            d_1265_v53_ = _dafny.Set({(d_1262_v51_) + (d_1262_v51_), d_1262_v51_})
            r2 = d_1265_v53_
            if (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]:
                d_1259_v48_ = (d_1259_v48_).set(710, (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))])
                d_1192_v3_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwf"))).set(default__.safeIndex(d_1190_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwf")))), d_1194_v5_)) + (d_1192_v3_)
                d_1266_v54_: D2
                d_1266_v54_ = D2_DC3(d_1192_v3_)
                rhs167_ = d_1190_v1_
                rhs168_ = len((d_1192_v3_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ti"))))
                rhs169_ = len((d_1266_v54_).cf7)
                lhs163_ = globalState
                lhs164_ = globalState
                lhs165_ = globalState
                lhs163_.f18 = rhs167_
                lhs164_.f18 = rhs168_
                lhs165_.f18 = rhs169_
                d_1267_v55_: _dafny.MultiSet
                d_1267_v55_ = _dafny.MultiSet([(d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]])
                d_1268_v56_: _dafny.Seq
                d_1268_v56_ = _dafny.SeqWithoutIsStrInference([d_1267_v55_])
                d_1269_v57_: D17
                d_1269_v57_ = D17_DC46(d_1268_v56_)
                d_1270_v58_: D8
                d_1270_v58_ = D8_DC23((d_1192_v3_).set(default__.safeIndex(len((d_1269_v57_).cf91), len(d_1192_v3_)), d_1194_v5_))
                d_1271_v59_: _dafny.Map
                d_1271_v59_ = _dafny.Map({(d_1270_v58_).cf44: d_1263_v52_})
                d_1272_v60_: D11
                d_1272_v60_ = D11_DC31(d_1271_v59_, d_1190_v1_)
                d_1273_v61_: _dafny.Set
                d_1273_v61_ = _dafny.Set({d_1272_v60_})
                d_1274_v62_: _dafny.Seq
                d_1274_v62_ = _dafny.SeqWithoutIsStrInference([d_1273_v61_, (_dafny.Set({d_1272_v60_})).intersection(d_1273_v61_), (d_1273_v61_).intersection(d_1273_v61_), d_1273_v61_])
                d_1274_v62_ = (_dafny.SeqWithoutIsStrInference([d_1273_v61_, d_1273_v61_])) + (d_1274_v62_)
                d_1275_v63_: _dafny.Array
                nw209_ = _dafny.Array(None, 3)
                nw209_[int(0)] = d_1194_v5_
                nw209_[int(1)] = (d_1192_v3_)[default__.safeIndex(len(d_1192_v3_), len(d_1192_v3_))]
                nw209_[int(2)] = d_1194_v5_
                d_1275_v63_ = nw209_
                d_1276_v64_: _dafny.Map
                d_1276_v64_ = _dafny.Map({d_1194_v5_: d_1190_v1_})
                d_1277_v65_: D18
                d_1277_v65_ = D18_DC49(d_1275_v63_)
                d_1275_v63_ = ((d_1277_v65_).cf95 if (((d_1276_v64_)[d_1194_v5_] if (d_1194_v5_) in (d_1276_v64_) else d_1190_v1_)) == (d_1190_v1_) else d_1275_v63_)
            elif True:
                d_1278_v66_: _dafny.MultiSet
                d_1278_v66_ = _dafny.MultiSet([len(d_1192_v3_)])
                rhs170_ = _dafny.MultiSet([default__.safeModuloInt(d_1190_v1_, d_1190_v1_), default__.safeDivisionInt(len(d_1192_v3_), d_1190_v1_), d_1190_v1_])
                rhs171_ = (d_1192_v3_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lioncxx"))) + (d_1192_v3_))
                rhs172_ = -984
                lhs166_ = globalState
                d_1278_v66_ = rhs170_
                lhs166_.f15 = rhs171_
                d_1190_v1_ = rhs172_
                d_1279_v67_: _dafny.Array
                def lambda56_(d_1280_v52_, d_1281_cf67_):
                    def lambda57_(d_1282_i7_):
                        return (d_1282_i7_) + (len(_dafny.Map({(d_1280_v52_)[default__.safeIndex(733, (d_1280_v52_).length(0))]: len(d_1281_cf67_)})))

                    return lambda57_

                init29_ = lambda56_(d_1263_v52_, d_1258_cf67_)
                nw210_ = _dafny.Array(None, 18)
                for i0_29_ in range(nw210_.length(0)):
                    nw210_[i0_29_] = init29_(i0_29_)
                d_1279_v67_ = nw210_
                r0 = d_1279_v67_
                d_1283_v68_: _dafny.Map
                d_1283_v68_ = _dafny.Map({d_1189_v0_: False})
                d_1284_v69_: _dafny.MultiSet
                d_1284_v69_ = _dafny.MultiSet([d_1189_v0_, (d_1283_v68_) == (_dafny.Map({not((d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]): default__.fm2(d_1261_v50_, d_1190_v1_, globalState)}))])
                d_1284_v69_ = (d_1284_v69_) | (d_1284_v69_)
                index207_ = default__.safeIndex(382, (d_1279_v67_).length(0))
                (d_1279_v67_)[index207_] = -909
                index208_ = default__.safeIndex(382, (d_1279_v67_).length(0))
                (d_1279_v67_)[index208_] = default__.safeModuloInt((d_1190_v1_) - (d_1190_v1_), ((d_1278_v66_)[d_1190_v1_] if (d_1190_v1_) in (d_1278_v66_) else d_1190_v1_))
                d_1285_v70_: _dafny.Map
                d_1285_v70_ = _dafny.Map({(d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]: d_1278_v66_})
                d_1286_v71_: T0
                nw211_ = C3()
                nw211_.ctor__((default__.fm57(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnmfxcv"))).set(default__.safeIndex(len(d_1285_v70_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnmfxcv")))), d_1194_v5_)), globalState) if (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))] else D0_DC0(d_1189_v0_)))
                d_1286_v71_ = nw211_
                d_1287_v72_: D15
                d_1287_v72_ = D15_DC42(d_1286_v71_)
                d_1286_v71_ = (d_1287_v72_).cf81
            if False:
                (globalState).f5 = (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]
                d_1288_v73_: _dafny.Array
                def lambda58_(d_1289_v2_):
                    def lambda59_(d_1290_i8_):
                        return d_1289_v2_

                    return lambda59_

                init30_ = lambda58_(d_1191_v2_)
                nw212_ = _dafny.Array(None, 11)
                for i0_30_ in range(nw212_.length(0)):
                    nw212_[i0_30_] = init30_(i0_30_)
                d_1288_v73_ = nw212_
                index209_ = default__.safeIndex(225, (d_1288_v73_).length(0))
                (d_1288_v73_)[index209_] = d_1191_v2_
                index210_ = default__.safeIndex(225, (d_1288_v73_).length(0))
                (d_1288_v73_)[index210_] = d_1191_v2_
                (globalState).f5 = ((d_1262_v51_) + ((_dafny.SeqWithoutIsStrInference([False, d_1189_v0_])).set(default__.safeIndex(len(d_1192_v3_), len(_dafny.SeqWithoutIsStrInference([False, d_1189_v0_]))), True))) != (d_1262_v51_)
                (globalState).f6 = (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]
                d_1291_v74_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1291_v74_ = nw213_
                d_1292_v75_: _dafny.Array
                nw214_ = _dafny.Array(int(0), 16)
                d_1292_v75_ = nw214_
                index211_ = default__.safeIndex(992, (d_1291_v74_).length(0))
                (d_1291_v74_)[index211_] = d_1292_v75_
                index212_ = default__.safeIndex(992, (d_1291_v74_).length(0))
                (d_1291_v74_)[index212_] = d_1292_v75_
            elif True:
                d_1293_v76_: _dafny.Set
                d_1293_v76_ = _dafny.Set({d_1189_v0_, True})
                d_1294_v77_: D0
                d_1294_v77_ = D0_DC1(_dafny.CodePoint('t'), d_1190_v1_, d_1293_v76_, (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))], d_1189_v0_)
                (globalState).f17 = (default__.safeDivisionInt(len(_dafny.Map({d_1192_v3_: d_1189_v0_})), d_1190_v1_)) + ((d_1190_v1_) + (default__.fm0(d_1190_v1_, d_1294_v77_, globalState)))
                d_1295_v78_: D17
                d_1295_v78_ = D17_DC47(d_1190_v1_)
                (globalState).f17 = ((d_1295_v78_).cf92) * (d_1190_v1_)
                (globalState).f6 = (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]
                d_1296_v79_: _dafny.Map
                d_1296_v79_ = _dafny.Map({d_1189_v0_: _dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1297_i9_ in range(default__.abs(317))])})
                d_1296_v79_ = (d_1296_v79_).set((d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))], d_1192_v3_)
                (globalState).f6 = (d_1263_v52_)[default__.safeIndex(733, (d_1263_v52_).length(0))]
        d_1298_v80_: _dafny.Array
        nw215_ = _dafny.Array(int(0), 11)
        d_1298_v80_ = nw215_
        index213_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        (d_1298_v80_)[index213_] = d_1190_v1_
        d_1299_v81_: _dafny.Map
        d_1299_v81_ = _dafny.Map({d_1194_v5_: d_1190_v1_})
        d_1300_v82_: _dafny.MultiSet
        d_1300_v82_ = _dafny.MultiSet([len(d_1192_v3_), d_1190_v1_])
        d_1301_v83_: _dafny.Seq
        d_1301_v83_ = _dafny.SeqWithoutIsStrInference([d_1300_v82_])
        index214_ = default__.safeIndex(495, (d_1298_v80_).length(0))
        (d_1298_v80_)[index214_] = (((d_1299_v81_)[d_1194_v5_] if (d_1194_v5_) in (d_1299_v81_) else len(d_1301_v83_))) - (d_1190_v1_)
        d_1302_v84_: _dafny.MultiSet
        d_1302_v84_ = _dafny.MultiSet([d_1189_v0_])
        d_1303_v85_: _dafny.Seq
        d_1303_v85_ = _dafny.SeqWithoutIsStrInference([False])
        d_1304_v86_: _dafny.Map
        d_1304_v86_ = _dafny.Map({d_1189_v0_: d_1298_v80_})
        d_1305_v87_: _dafny.Map
        d_1305_v87_ = _dafny.Map({((d_1304_v86_)[d_1189_v0_] if (d_1189_v0_) in (d_1304_v86_) else d_1298_v80_): d_1190_v1_})
        d_1306_v88_: _dafny.MultiSet
        d_1306_v88_ = (d_1302_v84_).set(True, default__.abs(d_1190_v1_))
        d_1307_v89_: _dafny.MultiSet
        d_1307_v89_ = _dafny.MultiSet([d_1194_v5_])
        index215_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        index216_ = default__.safeIndex(495, (d_1298_v80_).length(0))
        rhs173_ = d_1190_v1_
        rhs174_ = (D7_DC20((d_1302_v84_).cardinality, (d_1303_v85_)[default__.safeIndex(d_1190_v1_, len(d_1303_v85_))])).cf36
        rhs175_ = default__.safeModuloInt(len(default__.fm46(len(d_1305_v87_), d_1189_v0_, d_1189_v0_, globalState)), d_1190_v1_)
        rhs176_ = ((default__.fm50(d_1189_v0_, d_1189_v0_, d_1306_v88_, globalState)).set(default__.safeIndex(d_1190_v1_, len(default__.fm50(d_1189_v0_, d_1189_v0_, d_1306_v88_, globalState))), default__.fm2(d_1307_v89_, d_1190_v1_, globalState))) != (d_1303_v85_)
        rhs177_ = d_1302_v84_
        lhs167_ = d_1298_v80_
        lhs168_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        lhs169_ = d_1298_v80_
        lhs170_ = default__.safeIndex(495, (d_1298_v80_).length(0))
        d_1190_v1_ = rhs173_
        lhs167_[lhs168_] = rhs174_
        lhs169_[lhs170_] = rhs175_
        d_1189_v0_ = rhs176_
        d_1302_v84_ = rhs177_
        if (d_1190_v1_) < ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]):
            d_1308_v90_: T0
            nw216_ = C3()
            nw216_.ctor__((self).f20)
            d_1308_v90_ = nw216_
            d_1308_v90_ = d_1308_v90_
            (globalState).f5 = d_1189_v0_
            index217_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            (d_1298_v80_)[index217_] = 457
            index218_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            (d_1298_v80_)[index218_] = (0) - ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])
            r0 = d_1298_v80_
        elif True:
            rhs178_ = (d_1300_v82_) - (d_1300_v82_)
            rhs179_ = (d_1194_v5_) not in (d_1192_v3_)
            rhs180_ = d_1189_v0_
            rhs181_ = (((d_1300_v82_) | (d_1300_v82_)).cardinality) + (d_1190_v1_)
            lhs171_ = globalState
            d_1300_v82_ = rhs178_
            d_1189_v0_ = rhs179_
            r1 = rhs180_
            lhs171_.f17 = rhs181_
            (globalState).f6 = (default__.safeModuloInt((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])) != (d_1190_v1_)
            d_1309_v91_: D2
            d_1309_v91_ = D2_DC3(d_1192_v3_)
            d_1310_v92_: _dafny.Map
            d_1310_v92_ = _dafny.Map({d_1189_v0_: d_1190_v1_})
            pat_let_tv41_ = d_1192_v3_
            pat_let_tv42_ = d_1190_v1_
            pat_let_tv43_ = d_1192_v3_
            pat_let_tv44_ = d_1194_v5_
            pat_let_tv45_ = d_1194_v5_
            d_1311_v93_: _dafny.Array
            nw217_ = _dafny.Array(None, 25)
            nw217_[int(0)] = d_1309_v91_
            nw217_[int(1)] = D2_DC3((d_1192_v3_).set(default__.safeIndex((0) - (len(d_1310_v92_)), len(d_1192_v3_)), d_1194_v5_))
            nw217_[int(2)] = d_1309_v91_
            def iife156_(_pat_let53_0):
                def iife157_(d_1312_dt__update__tmp_h1_):
                    def iife158_(_pat_let54_0):
                        def iife159_(d_1313_dt__update_hcf7_h0_):
                            return D2_DC3(d_1313_dt__update_hcf7_h0_)
                        return iife159_(_pat_let54_0)
                    return iife158_((pat_let_tv41_).set(default__.safeIndex(pat_let_tv42_, len(pat_let_tv43_)), pat_let_tv44_))
                return iife157_(_pat_let53_0)
            nw217_[int(3)] = iife156_(d_1309_v91_)
            nw217_[int(4)] = D2_DC3(d_1192_v3_)
            nw217_[int(5)] = default__.fm61(d_1189_v0_, d_1189_v0_, len(d_1192_v3_), globalState)
            nw217_[int(6)] = D2_DC3(d_1192_v3_)
            nw217_[int(7)] = d_1309_v91_
            nw217_[int(8)] = d_1309_v91_
            nw217_[int(9)] = d_1309_v91_
            nw217_[int(10)] = D2_DC3(d_1192_v3_)
            nw217_[int(11)] = d_1309_v91_
            nw217_[int(12)] = D2_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjrkmdbrr")))
            nw217_[int(13)] = D2_DC3(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1314_i10_ in range(default__.abs(149))]))
            nw217_[int(14)] = d_1309_v91_
            nw217_[int(15)] = d_1309_v91_
            nw217_[int(16)] = d_1309_v91_
            nw217_[int(17)] = d_1309_v91_
            nw217_[int(18)] = d_1309_v91_
            def iife160_(_pat_let55_0):
                def iife161_(d_1315_dt__update__tmp_h2_):
                    def iife162_(_pat_let56_0):
                        def iife163_(d_1317_dt__update_hcf7_h1_):
                            return D2_DC3(d_1317_dt__update_hcf7_h1_)
                        return iife163_(_pat_let56_0)
                    return iife162_(_dafny.SeqWithoutIsStrInference([pat_let_tv45_ for d_1316_i11_ in range(default__.abs(921))]))
                return iife161_(_pat_let55_0)
            nw217_[int(19)] = iife160_(d_1309_v91_)
            nw217_[int(20)] = d_1309_v91_
            nw217_[int(21)] = d_1309_v91_
            nw217_[int(22)] = d_1309_v91_
            nw217_[int(23)] = d_1309_v91_
            nw217_[int(24)] = (default__.fm61(d_1189_v0_, False, 361, globalState) if not(d_1189_v0_) else d_1309_v91_)
            d_1311_v93_ = nw217_
            d_1318_v94_: _dafny.Set
            d_1318_v94_ = _dafny.Set({d_1190_v1_})
            d_1319_v95_: D13
            d_1319_v95_ = D13_DC35(d_1318_v94_)
            d_1320_v96_: _dafny.Array
            nw218_ = _dafny.Array(False, 12)
            d_1320_v96_ = nw218_
            index219_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            index220_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            rhs182_ = d_1190_v1_
            rhs183_ = d_1311_v93_
            rhs184_ = not(((d_1319_v95_).cf67).issubset(_dafny.Set({d_1190_v1_})))
            rhs185_ = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
            rhs186_ = d_1320_v96_
            lhs172_ = d_1298_v80_
            lhs173_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            lhs174_ = globalState
            lhs175_ = d_1298_v80_
            lhs176_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            lhs177_ = globalState
            lhs172_[lhs173_] = rhs182_
            d_1311_v93_ = rhs183_
            lhs174_.f6 = rhs184_
            lhs175_[lhs176_] = rhs185_
            lhs177_.f4 = rhs186_
            d_1190_v1_ = (d_1190_v1_) - ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])
            source16_ = d_1309_v91_
            if source16_.is_DC4:
                d_1321___mcc_h5_ = source16_.cf8
                d_1322___mcc_h6_ = source16_.cf9
                d_1323___mcc_h7_ = source16_.cf10
                d_1324___mcc_h8_ = source16_.cf11
                d_1325_cf11_ = d_1324___mcc_h8_
                d_1326_cf10_ = d_1323___mcc_h7_
                d_1327_cf9_ = d_1322___mcc_h6_
                d_1328_cf8_ = d_1321___mcc_h5_
                index221_ = default__.safeIndex(449, (d_1320_v96_).length(0))
                (d_1320_v96_)[index221_] = d_1327_cf9_
                index222_ = default__.safeIndex(109, (d_1320_v96_).length(0))
                (d_1320_v96_)[index222_] = d_1326_cf10_
                d_1329_v97_: _dafny.Map
                d_1329_v97_ = _dafny.Map({d_1326_cf10_: d_1327_cf9_})
                index223_ = default__.safeIndex(449, (d_1320_v96_).length(0))
                index224_ = default__.safeIndex(109, (d_1320_v96_).length(0))
                rhs187_ = (_dafny.MultiSet([len(d_1329_v97_)])).ispropersubset(_dafny.MultiSet([(d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]]))
                rhs188_ = False
                rhs189_ = d_1325_cf11_
                rhs190_ = d_1325_cf11_
                rhs191_ = d_1299_v81_
                lhs178_ = d_1320_v96_
                lhs179_ = default__.safeIndex(449, (d_1320_v96_).length(0))
                lhs180_ = d_1320_v96_
                lhs181_ = default__.safeIndex(109, (d_1320_v96_).length(0))
                lhs178_[lhs179_] = rhs187_
                d_1325_cf11_ = rhs188_
                d_1326_cf10_ = rhs189_
                lhs180_[lhs181_] = rhs190_
                d_1299_v81_ = rhs191_
                d_1326_cf10_ = (True) or (d_1325_cf11_)
                index225_ = default__.safeIndex(449, (d_1320_v96_).length(0))
                (d_1320_v96_)[index225_] = ((d_1329_v97_) | ((d_1329_v97_).set(d_1326_cf10_, d_1327_cf9_))) != (d_1329_v97_)
                d_1330_v98_: _dafny.Array
                nw219_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_1330_v98_ = nw219_
                d_1330_v98_ = d_1330_v98_
            elif source16_.is_DC5:
                d_1331___mcc_h9_ = source16_.cf12
                d_1332___mcc_h10_ = source16_.cf13
                d_1333_cf13_ = d_1332___mcc_h10_
                d_1334_cf12_ = d_1331___mcc_h9_
                d_1191_v2_ = d_1191_v2_
                d_1335_v99_: C2
                nw220_ = C2()
                nw220_.ctor__((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])
                d_1335_v99_ = nw220_
                d_1336_v100_: C0
                nw221_ = C0()
                nw221_.ctor__((D14_DC39(d_1320_v96_)).cf72, _dafny.SeqWithoutIsStrInference([False]))
                d_1336_v100_ = nw221_
                d_1337_v101_: D14
                d_1337_v101_ = D14_DC40(d_1189_v0_, d_1192_v3_, d_1192_v3_, d_1334_cf12_)
                (globalState).f6 = ((d_1337_v101_).cf76 if d_1334_cf12_ else (d_1336_v100_.f24) < (d_1303_v85_))
            elif source16_.is_DC3:
                d_1338___mcc_h11_ = source16_.cf7
                d_1339_cf7_ = d_1338___mcc_h11_
                d_1340_v102_: _dafny.Map
                d_1340_v102_ = _dafny.Map({d_1191_v2_: (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]})
                (globalState).f5 = (d_1191_v2_) not in (d_1340_v102_)
                index226_ = default__.safeIndex(694, (d_1298_v80_).length(0))
                (d_1298_v80_)[index226_] = (d_1191_v2_)[default__.safeIndex((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], len(d_1191_v2_))]
                d_1341_v103_: _dafny.Map
                d_1341_v103_ = _dafny.Map({(d_1300_v82_) - (d_1300_v82_): d_1194_v5_})
                d_1341_v103_ = ((default__.fm62(d_1189_v0_, d_1190_v1_, len(_dafny.Map({not(d_1189_v0_): d_1320_v96_})), globalState)) | ((d_1341_v103_).set(_dafny.MultiSet([d_1190_v1_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]]), d_1194_v5_))) | (_dafny.Map({d_1300_v82_: _dafny.CodePoint('f')}))
                d_1342_v104_: _dafny.Array
                nw222_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_1342_v104_ = nw222_
                d_1343_v105_: _dafny.Array
                nw223_ = _dafny.Array(_dafny.Map({}), 21)
                d_1343_v105_ = nw223_
                index227_ = default__.safeIndex(522, (d_1342_v104_).length(0))
                (d_1342_v104_)[index227_] = d_1343_v105_
                index228_ = default__.safeIndex(522, (d_1342_v104_).length(0))
                rhs192_ = ((d_1191_v2_) + (d_1191_v2_) if (d_1190_v1_) != (d_1190_v1_) else d_1191_v2_)
                rhs193_ = d_1343_v105_
                lhs182_ = d_1342_v104_
                lhs183_ = default__.safeIndex(522, (d_1342_v104_).length(0))
                d_1191_v2_ = rhs192_
                lhs182_[lhs183_] = rhs193_
            elif True:
                d_1344___mcc_h12_ = source16_.cf14
                d_1345_cf14_ = d_1344___mcc_h12_
                d_1346_v106_: _dafny.Array
                nw224_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_1346_v106_ = nw224_
                d_1347_v107_: _dafny.Set
                d_1347_v107_ = _dafny.Set({d_1189_v0_})
                d_1348_v108_: D0
                d_1348_v108_ = D0_DC1(d_1194_v5_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], d_1347_v107_, False, False)
                d_1349_v109_: D10
                d_1349_v109_ = D10_DC27(d_1346_v106_, default__.fm0((0) - (((d_1299_v81_)[d_1194_v5_] if (d_1194_v5_) in (d_1299_v81_) else d_1190_v1_)), d_1348_v108_, globalState))
                d_1349_v109_ = d_1349_v109_
                index229_ = default__.safeIndex(694, (d_1298_v80_).length(0))
                (d_1298_v80_)[index229_] = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
                d_1350_v110_: C1
                nw225_ = C1()
                nw225_.ctor__(d_1190_v1_, (self).f20)
                d_1350_v110_ = nw225_
                index230_ = default__.safeIndex(694, (d_1298_v80_).length(0))
                (d_1298_v80_)[index230_] = ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]) + (d_1190_v1_)
        d_1351_v111_: _dafny.Map
        d_1351_v111_ = _dafny.Map({d_1189_v0_: (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]})
        d_1352_v112_: _dafny.Set
        d_1352_v112_ = _dafny.Set({len(d_1351_v111_), (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]})
        if (d_1352_v112_).issubset(d_1352_v112_):
            d_1353_v113_: _dafny.Array
            nw226_ = _dafny.Array(False, 10)
            d_1353_v113_ = nw226_
            d_1354_v114_: D14
            d_1354_v114_ = D14_DC39(d_1353_v113_)
            d_1355_v115_: D14
            d_1355_v115_ = D14_DC39((d_1354_v114_).cf72)
            pat_let_tv46_ = d_1353_v113_
            d_1356_v116_: _dafny.Array
            nw227_ = _dafny.Array(None, 17)
            nw227_[int(0)] = d_1353_v113_
            nw227_[int(1)] = (d_1353_v113_ if d_1189_v0_ else d_1353_v113_)
            def iife164_(_pat_let57_0):
                def iife165_(d_1357_dt__update__tmp_h3_):
                    def iife166_(_pat_let58_0):
                        def iife167_(d_1358_dt__update_hcf72_h0_):
                            return D14_DC39(d_1358_dt__update_hcf72_h0_)
                        return iife167_(_pat_let58_0)
                    return iife166_(pat_let_tv46_)
                return iife165_(_pat_let57_0)
            nw227_[int(2)] = (iife164_(d_1355_v115_)).cf72
            nw227_[int(3)] = d_1353_v113_
            nw227_[int(4)] = d_1353_v113_
            nw227_[int(5)] = d_1353_v113_
            nw227_[int(6)] = d_1353_v113_
            nw227_[int(7)] = d_1353_v113_
            nw227_[int(8)] = d_1353_v113_
            nw227_[int(9)] = d_1353_v113_
            nw227_[int(10)] = d_1353_v113_
            nw227_[int(11)] = (d_1353_v113_ if default__.fm2(d_1307_v89_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], globalState) else d_1353_v113_)
            nw227_[int(12)] = d_1353_v113_
            nw227_[int(13)] = d_1353_v113_
            nw227_[int(14)] = d_1353_v113_
            nw227_[int(15)] = d_1353_v113_
            nw227_[int(16)] = d_1353_v113_
            d_1356_v116_ = nw227_
            d_1356_v116_ = (d_1356_v116_ if default__.fm2(d_1307_v89_, d_1190_v1_, globalState) else d_1356_v116_)
            if False:
                index231_ = default__.safeIndex(694, (d_1298_v80_).length(0))
                (d_1298_v80_)[index231_] = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
                d_1359_v117_: _dafny.Map
                d_1359_v117_ = _dafny.Map({d_1298_v80_: d_1353_v113_})
                d_1359_v117_ = (d_1359_v117_).set(d_1298_v80_, d_1353_v113_)
                d_1360_v118_: _dafny.Set
                d_1360_v118_ = _dafny.Set({False, d_1189_v0_})
                d_1361_v119_: D0
                d_1361_v119_ = D0_DC1(d_1194_v5_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], d_1360_v118_, True, True)
                pat_let_tv47_ = d_1360_v118_
                pat_let_tv48_ = d_1298_v80_
                pat_let_tv49_ = d_1298_v80_
                d_1362_v120_: D2
                def iife168_(_pat_let59_0):
                    def iife169_(d_1363_dt__update__tmp_h4_):
                        def iife170_(_pat_let60_0):
                            def iife171_(d_1364_dt__update_hcf3_h0_):
                                def iife172_(_pat_let61_0):
                                    def iife173_(d_1365_dt__update_hcf2_h0_):
                                        return D0_DC1((d_1363_dt__update__tmp_h4_).cf1, d_1365_dt__update_hcf2_h0_, d_1364_dt__update_hcf3_h0_, (d_1363_dt__update__tmp_h4_).cf4, (d_1363_dt__update__tmp_h4_).cf5)
                                    return iife173_(_pat_let61_0)
                                return iife172_((pat_let_tv49_)[default__.safeIndex(694, (pat_let_tv48_).length(0))])
                            return iife171_(_pat_let60_0)
                        return iife170_(pat_let_tv47_)
                    return iife169_(_pat_let59_0)
                d_1362_v120_ = D2_DC5(default__.fm2(d_1307_v89_, default__.fm0(len(d_1351_v111_), iife168_(d_1361_v119_), globalState), globalState), d_1190_v1_)
                (globalState).f18 = (d_1362_v120_).cf13
                (globalState).f5 = d_1189_v0_
                index232_ = default__.safeIndex(694, (d_1298_v80_).length(0))
                (d_1298_v80_)[index232_] = default__.safeDivisionInt((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], (d_1190_v1_) + (d_1190_v1_))
            elif True:
                d_1353_v113_ = d_1353_v113_
                d_1366_v121_: _dafny.Map
                d_1366_v121_ = _dafny.Map({d_1189_v0_: d_1189_v0_})
                d_1367_v122_: _dafny.Map
                d_1367_v122_ = _dafny.Map({len(d_1366_v121_): d_1189_v0_})
                d_1368_v123_: _dafny.Map
                d_1368_v123_ = _dafny.Map({d_1190_v1_: (_dafny.SeqWithoutIsStrInference([not(((d_1367_v122_)[(d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]] if ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]) in (d_1367_v122_) else True))])) + (d_1303_v85_)})
                (globalState).f17 = len(((d_1368_v123_)[default__.safeDivisionInt(d_1190_v1_, (0) - (len(d_1191_v2_)))] if (default__.safeDivisionInt(d_1190_v1_, (0) - (len(d_1191_v2_)))) in (d_1368_v123_) else (d_1303_v85_).set(default__.safeIndex(d_1190_v1_, len(d_1303_v85_)), d_1189_v0_)))
                index233_ = default__.safeIndex(470, (d_1353_v113_).length(0))
                (d_1353_v113_)[index233_] = d_1189_v0_
                index234_ = default__.safeIndex(470, (d_1353_v113_).length(0))
                (d_1353_v113_)[index234_] = True
                d_1369_v124_: _dafny.Set
                d_1369_v124_ = _dafny.Set({(d_1353_v113_)[default__.safeIndex(470, (d_1353_v113_).length(0))]})
                d_1370_v126_: _dafny.Seq
                def iife174_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in (d_1191_v2_).Elements:
                        d_1371_v125_: int = compr_50_
                        if (d_1371_v125_) in (d_1191_v2_):
                            coll50_[default__.safeDivisionInt(d_1371_v125_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])] = (d_1353_v113_)[default__.safeIndex(470, (d_1353_v113_).length(0))]
                    return _dafny.Map(coll50_)
                d_1370_v126_ = _dafny.SeqWithoutIsStrInference([iife174_()
                , d_1367_v122_, d_1367_v122_, d_1367_v122_, ((d_1367_v122_).set((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], (d_1353_v113_)[default__.safeIndex(470, (d_1353_v113_).length(0))])).set(d_1190_v1_, (d_1353_v113_)[default__.safeIndex(470, (d_1353_v113_).length(0))])])
                index235_ = default__.safeIndex(470, (d_1353_v113_).length(0))
                rhs194_ = (d_1302_v84_).set((-837) > (len(d_1369_v124_)), default__.abs(len(d_1370_v126_)))
                rhs195_ = default__.safeModuloInt(default__.safeDivisionInt((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], 217), d_1190_v1_)
                rhs196_ = not(d_1189_v0_)
                lhs184_ = globalState
                lhs185_ = d_1353_v113_
                lhs186_ = default__.safeIndex(470, (d_1353_v113_).length(0))
                d_1302_v84_ = rhs194_
                lhs184_.f18 = rhs195_
                lhs185_[lhs186_] = rhs196_
                d_1372_v127_: _dafny.Map
                d_1372_v127_ = _dafny.Map({d_1190_v1_: len(d_1192_v3_)})
                (globalState).f5 = (((d_1372_v127_)[d_1190_v1_] if (d_1190_v1_) in (d_1372_v127_) else (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])) == ((0) - ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]))
            d_1373_v128_: D13
            d_1373_v128_ = D13_DC35(d_1352_v112_)
            d_1373_v128_ = d_1373_v128_
            d_1374_v129_: C2
            nw228_ = C2()
            nw228_.ctor__((d_1193_v4_).cf68)
            d_1374_v129_ = nw228_
            (globalState).f5 = d_1189_v0_
        elif True:
            (globalState).f17 = ((d_1190_v1_ if not(d_1189_v0_) else (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))])) - (d_1190_v1_)
            index236_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            (d_1298_v80_)[index236_] = d_1190_v1_
            d_1189_v0_ = d_1189_v0_
            d_1303_v85_ = (d_1303_v85_).set(default__.safeIndex(d_1190_v1_, len(d_1303_v85_)), not (d_1189_v0_) or (d_1189_v0_))
            d_1375_v130_: _dafny.Map
            d_1375_v130_ = _dafny.Map({_dafny.Map({(d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]: d_1190_v1_}): default__.fm46((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], d_1189_v0_, d_1189_v0_, globalState)})
            d_1376_v131_: _dafny.Map
            d_1376_v131_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmgtb"))): len(d_1192_v3_)})
            d_1375_v130_ = (default__.fm63(d_1352_v112_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], globalState)).set(d_1376_v131_, (d_1192_v3_) + (d_1192_v3_))
        d_1377_v132_: D14
        d_1377_v132_ = D14_DC40(not(d_1189_v0_), _dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1378_i12_ in range(default__.abs(435))]), _dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1379_i13_ in range(default__.abs(247))]), True)
        pat_let_tv50_ = d_1189_v0_
        pat_let_tv51_ = d_1303_v85_
        pat_let_tv52_ = d_1303_v85_
        index237_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        index238_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        def lambda60_(source17_):
            if source17_.is_DC40:
                d_1380___mcc_h13_ = source17_.cf73
                d_1381___mcc_h14_ = source17_.cf74
                d_1382___mcc_h15_ = source17_.cf75
                d_1383___mcc_h16_ = source17_.cf76
                d_1384_cf76_ = d_1383___mcc_h16_
                d_1385_cf75_ = d_1382___mcc_h15_
                d_1386_cf74_ = d_1381___mcc_h14_
                d_1387_cf73_ = d_1380___mcc_h13_
                return d_1384_cf76_
            elif source17_.is_DC41:
                d_1388___mcc_h17_ = source17_.cf77
                d_1389___mcc_h18_ = source17_.cf78
                d_1390___mcc_h19_ = source17_.cf79
                d_1391___mcc_h20_ = source17_.cf80
                d_1392_cf80_ = d_1391___mcc_h20_
                d_1393_cf79_ = d_1390___mcc_h19_
                d_1394_cf78_ = d_1389___mcc_h18_
                d_1395_cf77_ = d_1388___mcc_h17_
                return pat_let_tv50_
            elif True:
                d_1396___mcc_h21_ = source17_.cf72
                d_1397_cf72_ = d_1396___mcc_h21_
                return (pat_let_tv51_)[default__.safeIndex(935, len(pat_let_tv52_))]

        rhs197_ = lambda60_(d_1377_v132_)
        rhs198_ = d_1189_v0_
        rhs199_ = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
        rhs200_ = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
        lhs187_ = globalState
        lhs188_ = globalState
        lhs189_ = d_1298_v80_
        lhs190_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        lhs191_ = d_1298_v80_
        lhs192_ = default__.safeIndex(694, (d_1298_v80_).length(0))
        lhs187_.f6 = rhs197_
        lhs188_.f6 = rhs198_
        lhs189_[lhs190_] = rhs199_
        lhs191_[lhs192_] = rhs200_
        d_1398_v133_: _dafny.Seq
        d_1398_v133_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1194_v5_ for d_1399_i14_ in range(default__.abs(309))])])
        d_1400_v134_: D4
        d_1400_v134_ = D4_DC10(d_1398_v133_)
        d_1401_v135_: D4
        d_1401_v135_ = D4_DC12(D4_DC12(d_1400_v134_))
        d_1402_v136_: D4
        d_1402_v136_ = D4_DC12(d_1400_v134_)
        pat_let_tv53_ = d_1189_v0_
        pat_let_tv54_ = d_1189_v0_
        def lambda61_(source18_):
            if source18_.is_DC11:
                d_1403___mcc_h22_ = source18_.cf19
                d_1404___mcc_h23_ = source18_.cf20
                d_1405___mcc_h24_ = source18_.cf21
                d_1406___mcc_h25_ = source18_.cf22
                d_1407_cf22_ = d_1406___mcc_h25_
                d_1408_cf21_ = d_1405___mcc_h24_
                d_1409_cf20_ = d_1404___mcc_h23_
                d_1410_cf19_ = d_1403___mcc_h22_
                return not(pat_let_tv53_)
            elif source18_.is_DC10:
                d_1411___mcc_h26_ = source18_.cf18
                d_1412_cf18_ = d_1411___mcc_h26_
                return True
            elif True:
                d_1413___mcc_h27_ = source18_.cf23
                d_1414_cf23_ = d_1413___mcc_h27_
                return pat_let_tv54_

        if lambda61_(d_1402_v136_):
            index239_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            (d_1298_v80_)[index239_] = ((d_1191_v2_)[default__.safeIndex(d_1190_v1_, len(d_1191_v2_))]) - (d_1190_v1_)
            d_1415_v137_: C2
            nw229_ = C2()
            nw229_.ctor__(len(d_1192_v3_))
            d_1415_v137_ = nw229_
            d_1299_v81_ = (d_1299_v81_).set(d_1194_v5_, (0) - (len(default__.fm46((d_1415_v137_).f22, d_1189_v0_, d_1189_v0_, globalState))))
            if ((d_1302_v84_).cardinality) != ((d_1190_v1_) * (d_1190_v1_)):
                d_1189_v0_ = d_1189_v0_
                (globalState).f18 = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
                (globalState).f6 = d_1189_v0_
                (globalState).f6 = d_1189_v0_
                d_1416_v138_: _dafny.Map
                d_1416_v138_ = _dafny.Map({d_1190_v1_: (d_1415_v137_).f22})
                d_1417_v139_: _dafny.Array
                nw230_ = _dafny.Array(None, 4)
                nw230_[int(0)] = d_1189_v0_
                nw230_[int(1)] = True
                nw230_[int(2)] = (566) in (d_1416_v138_)
                nw230_[int(3)] = False
                d_1417_v139_ = nw230_
                index240_ = default__.safeIndex(683, (d_1417_v139_).length(0))
                (d_1417_v139_)[index240_] = d_1189_v0_
                index241_ = default__.safeIndex(683, (d_1417_v139_).length(0))
                (d_1417_v139_)[index241_] = not(d_1189_v0_)
            elif True:
                d_1418_v140_: _dafny.Map
                d_1418_v140_ = _dafny.Map({d_1190_v1_: d_1189_v0_})
                d_1419_v141_: _dafny.Array
                nw231_ = _dafny.Array(None, 8)
                nw231_[int(0)] = d_1300_v82_
                nw231_[int(1)] = d_1300_v82_
                nw231_[int(2)] = d_1300_v82_
                nw231_[int(3)] = default__.fm56((d_1415_v137_).f22, ((d_1418_v140_)[(d_1415_v137_).f22] if ((d_1415_v137_).f22) in (d_1418_v140_) else d_1189_v0_), not(d_1189_v0_), (d_1415_v137_).f22, globalState)
                nw231_[int(4)] = d_1300_v82_
                nw231_[int(5)] = d_1300_v82_
                nw231_[int(6)] = d_1300_v82_
                nw231_[int(7)] = d_1300_v82_
                d_1419_v141_ = nw231_
                d_1420_v142_: _dafny.Map
                d_1420_v142_ = _dafny.Map({d_1419_v141_: d_1191_v2_})
                d_1191_v2_ = ((d_1420_v142_)[d_1419_v141_] if (d_1419_v141_) in (d_1420_v142_) else d_1191_v2_)
                d_1191_v2_ = d_1191_v2_
                d_1421_v143_: _dafny.Set
                d_1421_v143_ = _dafny.Set({_dafny.CodePoint('g'), d_1194_v5_})
                d_1422_v144_: _dafny.Array
                nw232_ = _dafny.Array(None, 19)
                nw232_[int(0)] = d_1189_v0_
                nw232_[int(1)] = d_1189_v0_
                nw232_[int(2)] = d_1189_v0_
                nw232_[int(3)] = not(((d_1415_v137_).f22) <= ((d_1415_v137_).f22))
                nw232_[int(4)] = d_1189_v0_
                nw232_[int(5)] = default__.fm2(d_1307_v89_, (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], globalState)
                nw232_[int(6)] = d_1189_v0_
                nw232_[int(7)] = not (d_1189_v0_) or (d_1189_v0_)
                nw232_[int(8)] = d_1189_v0_
                nw232_[int(9)] = not(not(d_1189_v0_))
                nw232_[int(10)] = d_1189_v0_
                nw232_[int(11)] = d_1189_v0_
                nw232_[int(12)] = (161) < ((d_1415_v137_).f22)
                nw232_[int(13)] = not(True)
                nw232_[int(14)] = d_1189_v0_
                nw232_[int(15)] = default__.fm2(d_1307_v89_, (d_1415_v137_).f22, globalState)
                nw232_[int(16)] = d_1189_v0_
                nw232_[int(17)] = (d_1421_v143_).ispropersubset(d_1421_v143_)
                nw232_[int(18)] = d_1189_v0_
                d_1422_v144_ = nw232_
                index242_ = default__.safeIndex(590, (d_1422_v144_).length(0))
                (d_1422_v144_)[index242_] = d_1189_v0_
                d_1423_v145_: _dafny.Set
                d_1423_v145_ = _dafny.Set({d_1189_v0_, d_1189_v0_, d_1189_v0_, not(d_1189_v0_)})
                index243_ = default__.safeIndex(590, (d_1422_v144_).length(0))
                (d_1422_v144_)[index243_] = ((d_1189_v0_ if d_1189_v0_ else d_1189_v0_) if d_1189_v0_ else ((d_1415_v137_).f22) >= (len(d_1423_v145_)))
                (globalState).f4 = (d_1422_v144_ if not(d_1189_v0_) else d_1422_v144_)
                d_1189_v0_ = (d_1422_v144_)[default__.safeIndex(590, (d_1422_v144_).length(0))]
            d_1190_v1_ = -319
        elif True:
            d_1424_v146_: C2
            nw233_ = C2()
            nw233_.ctor__(d_1190_v1_)
            d_1424_v146_ = nw233_
            index244_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            (d_1298_v80_)[index244_] = 738
            (globalState).f18 = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
            index245_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            rhs201_ = d_1189_v0_
            rhs202_ = d_1302_v84_
            rhs203_ = (d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]
            lhs193_ = globalState
            lhs194_ = d_1298_v80_
            lhs195_ = default__.safeIndex(694, (d_1298_v80_).length(0))
            lhs193_.f5 = rhs201_
            d_1302_v84_ = rhs202_
            lhs194_[lhs195_] = rhs203_
            if d_1189_v0_:
                (globalState).f6 = d_1189_v0_
                d_1425_v147_: C3
                nw234_ = C3()
                nw234_.ctor__((self).f20)
                d_1425_v147_ = nw234_
                d_1426_v148_: _dafny.Map
                d_1426_v148_ = _dafny.Map({((d_1300_v82_).intersection(_dafny.MultiSet([(d_1424_v146_).f22]))).cardinality: d_1189_v0_})
                def iife175_():
                    coll51_ = _dafny.Set()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(86, 170):
                        d_1427_v149_: int = compr_51_
                        if ((86) <= (d_1427_v149_)) and ((d_1427_v149_) < (170)):
                            coll51_ = coll51_.union(_dafny.Set([(d_1427_v149_) - (len(_dafny.Set({d_1189_v0_, d_1189_v0_, d_1189_v0_, d_1189_v0_})))]))
                    return _dafny.Set(coll51_)
                d_1426_v148_ = (d_1426_v148_).set((d_1424_v146_).f22, (len(iife175_()
                )) < (((d_1299_v81_)[d_1194_v5_] if (d_1194_v5_) in (d_1299_v81_) else (d_1424_v146_).f22)))
                d_1398_v133_ = ((D4_DC10(d_1398_v133_)).cf18) + (d_1398_v133_)
                (globalState).f5 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gltghxa"))) < (d_1192_v3_)
            elif True:
                d_1428_v150_: _dafny.Map
                d_1428_v150_ = _dafny.Map({d_1189_v0_: d_1189_v0_})
                d_1429_v151_: _dafny.Seq
                d_1429_v151_ = _dafny.SeqWithoutIsStrInference([d_1351_v111_])
                d_1430_v152_: _dafny.Map
                d_1430_v152_ = _dafny.Map({((d_1428_v150_)[d_1189_v0_] if (d_1189_v0_) in (d_1428_v150_) else d_1189_v0_): d_1429_v151_})
                d_1430_v152_ = (d_1430_v152_).set((d_1189_v0_) or (d_1189_v0_), d_1429_v151_)
                d_1431_v153_: C2
                nw235_ = C2()
                nw235_.ctor__((0) - ((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))]))
                d_1431_v153_ = nw235_
                d_1432_v154_: _dafny.Map
                d_1432_v154_ = _dafny.Map({(d_1424_v146_).f22: d_1189_v0_})
                d_1433_v155_: _dafny.Set
                d_1433_v155_ = _dafny.Set({d_1189_v0_, d_1189_v0_, d_1189_v0_})
                d_1434_v156_: _dafny.Seq
                d_1434_v156_ = _dafny.SeqWithoutIsStrInference([d_1433_v155_, d_1433_v155_, d_1433_v155_, d_1433_v155_])
                d_1432_v154_ = (d_1432_v154_).set(len(d_1434_v156_), not(not (d_1189_v0_) or (d_1189_v0_)))
                (globalState).f5 = not(not(d_1189_v0_))
                d_1435_v157_: C1
                nw236_ = C1()
                nw236_.ctor__((d_1424_v146_).f22, (self).f20)
                d_1435_v157_ = nw236_
        d_1436_v158_: _dafny.Seq
        d_1436_v158_ = _dafny.SeqWithoutIsStrInference([d_1298_v80_, d_1298_v80_, d_1298_v80_])
        r0 = (d_1436_v158_)[default__.safeIndex((d_1298_v80_)[default__.safeIndex(694, (d_1298_v80_).length(0))], len(d_1436_v158_))]
        r1 = d_1189_v0_
        d_1437_v159_: _dafny.Set
        d_1437_v159_ = _dafny.Set({d_1303_v85_, _dafny.SeqWithoutIsStrInference([d_1189_v0_, not(False)])})
        r2 = d_1437_v159_
        return r0, r1, r2


class C5(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        self._f27: _dafny.MultiSet = _dafny.MultiSet({})
        self._f28: D3 = D3.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f27, f28, f20):
        (self)._f27 = f27
        (self)._f28 = f28
        (self)._f20 = f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        d_1438_v0_: _dafny.Seq
        d_1438_v0_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), p1])
        d_1438_v0_ = (d_1438_v0_) + ((d_1438_v0_) + (d_1438_v0_))
        d_1439_i0_: int
        d_1439_i0_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_1439_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_1439_i0_ = (d_1439_i0_) + (1)
                    d_1440_v1_: str
                    d_1440_v1_ = _dafny.CodePoint('w')
                    r2 = d_1440_v1_
                    if not(not(True)):
                        d_1441_v3_: C1
                        nw237_ = C1()
                        nw237_.ctor__(p1, (self).f20)
                        d_1441_v3_ = nw237_
                        d_1442_v4_: _dafny.Map
                        d_1442_v4_ = _dafny.Map({p1: d_1441_v3_})
                        d_1443_v5_: _dafny.Map
                        def iife176_():
                            coll52_ = _dafny.Set()
                            compr_52_: int
                            for compr_52_ in _dafny.IntegerRange(87, 762):
                                d_1444_v2_: int = compr_52_
                                if ((87) <= (d_1444_v2_)) and ((d_1444_v2_) < (762)):
                                    coll52_ = coll52_.union(_dafny.Set([default__.safeModuloInt(d_1444_v2_, len(_dafny.Map({p0: p1})))]))
                            return _dafny.Set(coll52_)
                        d_1443_v5_ = _dafny.Map({(iife176_()
                        ) - (_dafny.Set({p1, p1})): d_1442_v4_})
                        d_1443_v5_ = ((d_1443_v5_) | (d_1443_v5_)) | (d_1443_v5_)
                        d_1445_v6_: _dafny.Seq
                        d_1445_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cimpundh"))
                        (globalState).f18 = (0) - (len(((default__.fm34(p0, globalState)) + (d_1445_v6_)) + (d_1445_v6_)))
                        d_1446_v7_: C1
                        nw238_ = C1()
                        nw238_.ctor__(default__.safeDivisionInt(p1, p1), (self).f20)
                        d_1446_v7_ = nw238_
                        d_1447_v8_: _dafny.Map
                        d_1447_v8_ = _dafny.Map({p0: p0})
                        d_1448_v9_: _dafny.Map
                        d_1448_v9_ = _dafny.Map({p1: d_1440_v1_})
                        d_1449_v10_: _dafny.Map
                        d_1449_v10_ = _dafny.Map({d_1447_v8_: (d_1448_v9_).set(len(d_1445_v6_), d_1440_v1_)})
                        d_1450_v11_: _dafny.Seq
                        d_1450_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1440_v1_ for d_1451_i1_ in range(default__.abs(-441))]), _dafny.SeqWithoutIsStrInference([d_1440_v1_, d_1440_v1_, d_1440_v1_])])
                        d_1452_v12_: _dafny.Map
                        d_1452_v12_ = _dafny.Map({d_1449_v10_: (0) - (default__.safeModuloInt(len((d_1450_v11_)[default__.safeIndex(p1, len(d_1450_v11_))]), p1))})
                        d_1452_v12_ = (d_1452_v12_).set(d_1449_v10_, (0) - (len(d_1447_v8_)))
                        d_1453_v13_: _dafny.Array
                        nw239_ = _dafny.Array(False, 5)
                        d_1453_v13_ = nw239_
                        d_1454_v14_: _dafny.MultiSet
                        d_1454_v14_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1440_v1_ for d_1455_i2_ in range(default__.abs(638))])), p1])
                        d_1456_v15_: _dafny.Array
                        nw240_ = _dafny.Array(None, 23)
                        nw240_[int(0)] = p1
                        nw240_[int(1)] = -356
                        nw240_[int(2)] = p1
                        nw240_[int(3)] = len(_dafny.Map({(d_1454_v14_).cardinality: p0}))
                        nw240_[int(4)] = p1
                        nw240_[int(5)] = p1
                        nw240_[int(6)] = p1
                        nw240_[int(7)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs")))
                        nw240_[int(8)] = p1
                        nw240_[int(9)] = -817
                        nw240_[int(10)] = p1
                        nw240_[int(11)] = p1
                        nw240_[int(12)] = p1
                        nw240_[int(13)] = p1
                        nw240_[int(14)] = p1
                        nw240_[int(15)] = p1
                        nw240_[int(16)] = p1
                        nw240_[int(17)] = p1
                        nw240_[int(18)] = p1
                        nw240_[int(19)] = p1
                        nw240_[int(20)] = p1
                        nw240_[int(21)] = p1
                        nw240_[int(22)] = (0) - (p1)
                        d_1456_v15_ = nw240_
                        d_1457_v16_: D7
                        d_1457_v16_ = D7_DC21(765, p0, p0, d_1456_v15_, p0)
                        d_1458_v17_: D8
                        d_1458_v17_ = D8_DC22(_dafny.Map({d_1457_v16_: len(d_1445_v6_)}))
                        index246_ = default__.safeIndex(652, (d_1453_v13_).length(0))
                        (d_1453_v13_)[index246_] = (_dafny.Map({d_1457_v16_: p1})) != ((d_1458_v17_).cf43)
                        d_1459_v18_: _dafny.Seq
                        d_1459_v18_ = _dafny.SeqWithoutIsStrInference([not(not(p0))])
                        d_1460_v19_: _dafny.Map
                        d_1460_v19_ = _dafny.Map({p1: d_1456_v15_})
                        d_1461_v20_: _dafny.Map
                        d_1461_v20_ = _dafny.Map({d_1460_v19_: d_1459_v18_})
                        index247_ = default__.safeIndex(652, (d_1453_v13_).length(0))
                        (d_1453_v13_)[index247_] = (d_1459_v18_) < (((d_1461_v20_)[_dafny.Map({p1: d_1456_v15_})] if (_dafny.Map({p1: d_1456_v15_})) in (d_1461_v20_) else d_1459_v18_))
                    elif True:
                        d_1462_v21_: _dafny.Array
                        def lambda62_(d_1463_i3_):
                            return default__.safeModuloInt(d_1463_i3_, 902)

                        init31_ = lambda62_
                        nw241_ = _dafny.Array(None, 6)
                        for i0_31_ in range(nw241_.length(0)):
                            nw241_[i0_31_] = init31_(i0_31_)
                        d_1462_v21_ = nw241_
                        d_1462_v21_ = d_1462_v21_
                        d_1464_v22_: _dafny.MultiSet
                        d_1464_v22_ = _dafny.MultiSet([d_1440_v1_, d_1440_v1_])
                        (globalState).f18 = default__.fm0(p1, default__.fm35(d_1464_v22_, globalState), globalState)
                        (globalState).f17 = (d_1438_v0_)[default__.safeIndex(default__.safeModuloInt((0) - (p1), p1), len(d_1438_v0_))]
                        d_1465_v23_: _dafny.Seq
                        d_1465_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uj"))
                        rhs204_ = ((d_1465_v23_) + (default__.fm34(p0, globalState))) + (d_1465_v23_)
                        rhs205_ = -759
                        lhs196_ = globalState
                        lhs197_ = globalState
                        lhs196_.f15 = rhs204_
                        lhs197_.f17 = rhs205_
                        d_1466_v24_: _dafny.Map
                        d_1466_v24_ = _dafny.Map({p0: p1})
                        d_1466_v24_ = d_1466_v24_
                    (globalState).f6 = p0
                    d_1467_v25_: _dafny.Map
                    d_1467_v25_ = _dafny.Map({(p1) * (p1): p0})
                    d_1467_v25_ = (d_1467_v25_).set(p1, p0)
                    pass
            pass
        d_1468_v26_: _dafny.Set
        d_1468_v26_ = _dafny.Set({(self).f27})
        d_1469_i4_: int
        d_1469_i4_ = 0
        with _dafny.label("6"):
            while (d_1468_v26_).issubset(d_1468_v26_):
                with _dafny.c_label("6"):
                    if (d_1469_i4_) >= (100):
                        raise _dafny.Break("6")
                    d_1469_i4_ = (d_1469_i4_) + (1)
                    d_1470_v27_: _dafny.Seq
                    d_1470_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0})])
                    d_1470_v27_ = ((d_1470_v27_ if p0 else d_1470_v27_) if p0 else d_1470_v27_)
                    d_1471_v28_: _dafny.Map
                    d_1471_v28_ = _dafny.Map({445: p0})
                    d_1472_v29_: _dafny.Set
                    d_1472_v29_ = _dafny.Set({d_1471_v28_})
                    d_1473_v30_: str
                    d_1473_v30_ = _dafny.CodePoint('j')
                    d_1474_v31_: _dafny.MultiSet
                    d_1474_v31_ = _dafny.MultiSet([_dafny.CodePoint('i'), d_1473_v30_])
                    d_1475_v33_: _dafny.Map
                    d_1475_v33_ = _dafny.Map({p0: d_1472_v29_})
                    d_1476_v35_: _dafny.Map
                    d_1476_v35_ = _dafny.Map({(0) - (p1): d_1472_v29_})
                    d_1477_v36_: _dafny.Map
                    d_1477_v36_ = _dafny.Map({_dafny.Map({p1: p0}): p0})
                    d_1478_v38_: _dafny.Array
                    nw242_ = _dafny.Array(None, 21)
                    nw242_[int(0)] = (default__.fm36(427, globalState)).intersection(d_1472_v29_)
                    def iife177_():
                        coll53_ = _dafny.Map()
                        compr_53_: int
                        for compr_53_ in _dafny.IntegerRange(430, 306):
                            d_1479_v32_: int = compr_53_
                            if ((430) <= (d_1479_v32_)) and ((d_1479_v32_) < (306)):
                                coll53_[default__.safeModuloInt(d_1479_v32_, p1)] = p0
                        return _dafny.Map(coll53_)
                    nw242_[int(1)] = (_dafny.Set({_dafny.Map({len(d_1438_v0_): p0}), d_1471_v28_, iife177_()
                    }) if default__.fm2(d_1474_v31_, 96, globalState) else d_1472_v29_)
                    nw242_[int(2)] = (d_1472_v29_).intersection(default__.fm36(p1, globalState))
                    nw242_[int(3)] = d_1472_v29_
                    nw242_[int(4)] = ((d_1475_v33_)[p0] if (p0) in (d_1475_v33_) else d_1472_v29_)
                    nw242_[int(5)] = d_1472_v29_
                    nw242_[int(6)] = d_1472_v29_
                    nw242_[int(7)] = _dafny.Set({d_1471_v28_, d_1471_v28_})
                    def iife178_():
                        coll54_ = _dafny.Map()
                        compr_54_: int
                        for compr_54_ in _dafny.IntegerRange(174, 56):
                            d_1480_v34_: int = compr_54_
                            if ((174) <= (d_1480_v34_)) and ((d_1480_v34_) < (56)):
                                coll54_[(d_1480_v34_) * (p1)] = p0
                        return _dafny.Map(coll54_)
                    nw242_[int(8)] = _dafny.Set({iife178_()
                    })
                    nw242_[int(9)] = d_1472_v29_
                    nw242_[int(10)] = d_1472_v29_
                    nw242_[int(11)] = (d_1472_v29_).intersection(d_1472_v29_)
                    nw242_[int(12)] = (d_1472_v29_) | (d_1472_v29_)
                    nw242_[int(13)] = d_1472_v29_
                    nw242_[int(14)] = d_1472_v29_
                    nw242_[int(15)] = ((d_1476_v35_)[p1] if (p1) in (d_1476_v35_) else d_1472_v29_)
                    nw242_[int(16)] = default__.fm36(p1, globalState)
                    nw242_[int(17)] = d_1472_v29_
                    def iife179_():
                        coll55_ = _dafny.Set()
                        compr_55_: _dafny.Map
                        for compr_55_ in (d_1477_v36_).keys.Elements:
                            d_1481_v37_: _dafny.Map = compr_55_
                            if (d_1481_v37_) in (d_1477_v36_):
                                coll55_ = coll55_.union(_dafny.Set([d_1481_v37_]))
                        return _dafny.Set(coll55_)
                    nw242_[int(18)] = iife179_()
                    
                    nw242_[int(19)] = d_1472_v29_
                    nw242_[int(20)] = _dafny.Set({d_1471_v28_, _dafny.Map({p1: p0})})
                    d_1478_v38_ = nw242_
                    index248_ = default__.safeIndex(496, (d_1478_v38_).length(0))
                    (d_1478_v38_)[index248_] = ((d_1476_v35_)[p1] if (p1) in (d_1476_v35_) else d_1472_v29_)
                    d_1482_v39_: _dafny.Map
                    d_1482_v39_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjqmvaink"))): p1})
                    d_1483_v41_: _dafny.Seq
                    def iife180_():
                        coll56_ = _dafny.Map()
                        compr_56_: int
                        for compr_56_ in _dafny.IntegerRange(555, -154):
                            d_1484_v40_: int = compr_56_
                            if ((555) <= (d_1484_v40_)) and ((d_1484_v40_) < (-154)):
                                coll56_[(d_1484_v40_) + (((d_1482_v39_)[p1] if (p1) in (d_1482_v39_) else p1))] = p1
                        return _dafny.Map(coll56_)
                    d_1483_v41_ = _dafny.SeqWithoutIsStrInference([d_1482_v39_, iife180_()
                    ])
                    d_1485_v42_: _dafny.Seq
                    d_1485_v42_ = _dafny.SeqWithoutIsStrInference([d_1471_v28_, d_1471_v28_, d_1471_v28_])
                    d_1486_v43_: _dafny.Seq
                    d_1486_v43_ = _dafny.SeqWithoutIsStrInference([d_1472_v29_])
                    index249_ = default__.safeIndex(496, (d_1478_v38_).length(0))
                    (d_1478_v38_)[index249_] = (_dafny.Set({d_1471_v28_, d_1471_v28_, _dafny.Map({len(d_1483_v41_): p0}), (d_1485_v42_)[default__.safeIndex(len(_dafny.Map({p1: False})), len(d_1485_v42_))], d_1471_v28_})) - ((d_1486_v43_)[default__.safeIndex(p1, len(d_1486_v43_))])
                    d_1487_v44_: _dafny.Seq
                    d_1487_v44_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_1488_v45_: _dafny.Seq
                    d_1488_v45_ = _dafny.SeqWithoutIsStrInference([d_1487_v44_])
                    d_1488_v45_ = d_1488_v45_
                    d_1489_v46_: _dafny.Array
                    def lambda63_(d_1490_v30_):
                        def lambda64_(d_1491_i5_):
                            return D4_DC10(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1490_v30_ for d_1492_i6_ in range(default__.abs(532))])]))

                        return lambda64_

                    init32_ = lambda63_(d_1473_v30_)
                    nw243_ = _dafny.Array(None, 14)
                    for i0_32_ in range(nw243_.length(0)):
                        nw243_[i0_32_] = init32_(i0_32_)
                    d_1489_v46_ = nw243_
                    d_1489_v46_ = d_1489_v46_
                    pass
            pass
        (globalState).f6 = False
        d_1493_v47_: C2
        nw244_ = C2()
        nw244_.ctor__(p1)
        d_1493_v47_ = nw244_
        d_1494_v48_: _dafny.Array
        nw245_ = _dafny.Array(False, 10)
        d_1494_v48_ = nw245_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1494_v48_).length(0)):
            d_1495_i7_: int = guard_loop_3_
            if (True) and (((0) <= (d_1495_i7_)) and ((d_1495_i7_) < ((d_1494_v48_).length(0)))):
                (d_1494_v48_)[(d_1495_i7_)] = p0
        d_1496_v49_: str
        d_1496_v49_ = _dafny.CodePoint('w')
        d_1497_v50_: _dafny.MultiSet
        d_1497_v50_ = _dafny.MultiSet([d_1496_v49_])
        d_1498_v51_: _dafny.MultiSet
        d_1498_v51_ = _dafny.MultiSet([default__.fm2(d_1497_v50_, (d_1493_v47_).f22, globalState)])
        r0 = d_1498_v51_
        r1 = p1
        d_1499_v52_: D4
        d_1499_v52_ = D4_DC11(not(not(p0)), p0, (d_1493_v47_).f22, p1)
        pat_let_tv55_ = d_1496_v49_
        pat_let_tv56_ = d_1496_v49_
        pat_let_tv57_ = d_1496_v49_
        def lambda65_(source19_):
            if source19_.is_DC11:
                d_1500___mcc_h0_ = source19_.cf19
                d_1501___mcc_h1_ = source19_.cf20
                d_1502___mcc_h2_ = source19_.cf21
                d_1503___mcc_h3_ = source19_.cf22
                d_1504_cf22_ = d_1503___mcc_h3_
                d_1505_cf21_ = d_1502___mcc_h2_
                d_1506_cf20_ = d_1501___mcc_h1_
                d_1507_cf19_ = d_1500___mcc_h0_
                return pat_let_tv55_
            elif source19_.is_DC10:
                d_1508___mcc_h4_ = source19_.cf18
                d_1509_cf18_ = d_1508___mcc_h4_
                return pat_let_tv56_
            elif True:
                d_1510___mcc_h5_ = source19_.cf23
                d_1511_cf23_ = d_1510___mcc_h5_
                return pat_let_tv57_

        r2 = lambda65_(d_1499_v52_)
        return r0, r1, r2

    def m10(self, p0, p1, globalState):
        d_1512_v0_: int
        d_1512_v0_ = -120
        (globalState).f18 = d_1512_v0_
        d_1513_v1_: _dafny.Seq
        d_1513_v1_ = _dafny.SeqWithoutIsStrInference([741])
        hi6_ = d_1512_v0_
        for d_1514_i0_ in range(len(d_1513_v1_), hi6_):
            (globalState).f18 = d_1512_v0_
            d_1515_v2_: _dafny.Array
            nw246_ = _dafny.Array(False, 4)
            d_1515_v2_ = nw246_
            (globalState).f4 = d_1515_v2_
            d_1516_v3_: bool
            d_1516_v3_ = True
            d_1517_v4_: _dafny.Map
            d_1517_v4_ = _dafny.Map({default__.fm37(d_1516_v3_, -999, globalState): False})
            d_1518_v5_: D6
            d_1518_v5_ = D6_DC16((_dafny.Map({d_1513_v1_: d_1516_v3_})) | (d_1517_v4_))
            d_1518_v5_ = d_1518_v5_
            (globalState).f17 = d_1514_i0_
        d_1519_v6_: bool
        d_1519_v6_ = True
        d_1520_v7_: _dafny.Map
        d_1520_v7_ = _dafny.Map({d_1512_v0_: p0})
        d_1521_v8_: str
        d_1521_v8_ = _dafny.CodePoint('e')
        d_1522_v9_: _dafny.MultiSet
        d_1522_v9_ = _dafny.MultiSet([d_1521_v8_])
        d_1523_v10_: _dafny.Map
        d_1523_v10_ = _dafny.Map({default__.fm2(d_1522_v9_, d_1512_v0_, globalState): p0})
        d_1524_v11_: _dafny.Map
        d_1524_v11_ = _dafny.Map({d_1519_v6_: len(((d_1520_v7_).set(d_1512_v0_, ((d_1523_v10_)[d_1519_v6_] if (d_1519_v6_) in (d_1523_v10_) else p0))).set(d_1512_v0_, p0))})
        d_1524_v11_ = d_1524_v11_
        d_1525_v12_: _dafny.Set
        d_1525_v12_ = _dafny.Set({d_1519_v6_})
        d_1526_v13_: D0
        d_1526_v13_ = D0_DC1(d_1521_v8_, d_1512_v0_, d_1525_v12_, d_1519_v6_, d_1519_v6_)
        d_1527_i1_: int
        d_1527_i1_ = 0
        with _dafny.label("7"):
            pat_let_tv63_ = p0
            def lambda74_(source20_):
                if source20_.is_DC1:
                    d_1587___mcc_h0_ = source20_.cf1
                    d_1588___mcc_h1_ = source20_.cf2
                    d_1589___mcc_h2_ = source20_.cf3
                    d_1590___mcc_h3_ = source20_.cf4
                    d_1591___mcc_h4_ = source20_.cf5
                    d_1592_cf5_ = d_1591___mcc_h4_
                    d_1593_cf4_ = d_1590___mcc_h3_
                    d_1594_cf3_ = d_1589___mcc_h2_
                    d_1595_cf2_ = d_1588___mcc_h1_
                    d_1596_cf1_ = d_1587___mcc_h0_
                    return (D2_DC5(True, len(pat_let_tv63_))).cf12
                elif True:
                    d_1597___mcc_h5_ = source20_.cf0
                    d_1598_cf0_ = d_1597___mcc_h5_
                    return d_1598_cf0_

            while lambda74_(d_1526_v13_):
                with _dafny.c_label("7"):
                    if (d_1527_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_1527_i1_ = (d_1527_i1_) + (1)
                    d_1528_v14_: _dafny.Seq
                    d_1528_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywdx"))
                    d_1529_v15_: D2
                    d_1529_v15_ = D2_DC3(d_1528_v14_)
                    d_1530_v16_: D2
                    d_1530_v16_ = D2_DC6(d_1529_v15_)
                    d_1531_v17_: D2
                    d_1531_v17_ = D2_DC6(d_1529_v15_)
                    d_1532_v18_: _dafny.Map
                    d_1532_v18_ = _dafny.Map({d_1531_v17_: d_1512_v0_})
                    d_1532_v18_ = (d_1532_v18_).set(default__.fm38(d_1512_v0_, globalState), d_1512_v0_)
                    d_1533_v19_: _dafny.Map
                    d_1533_v19_ = _dafny.Map({d_1521_v8_: d_1519_v6_})
                    d_1534_v20_: _dafny.Seq
                    d_1534_v20_ = _dafny.SeqWithoutIsStrInference([((d_1533_v19_)[_dafny.CodePoint('w')] if (_dafny.CodePoint('w')) in (d_1533_v19_) else d_1519_v6_), d_1519_v6_, default__.fm2(default__.fm39(d_1513_v1_, default__.fm2(d_1522_v9_, d_1512_v0_, globalState), len(p0), d_1512_v0_, globalState), default__.fm0(235, d_1526_v13_, globalState), globalState), False, d_1519_v6_])
                    d_1535_v21_: D8
                    d_1535_v21_ = D8_DC23(d_1528_v14_)
                    d_1536_v22_: _dafny.Map
                    d_1536_v22_ = _dafny.Map({d_1519_v6_: d_1519_v6_})
                    d_1537_v23_: _dafny.Array
                    nw247_ = _dafny.Array(int(0), 7)
                    d_1537_v23_ = nw247_
                    d_1538_v24_: _dafny.Map
                    d_1538_v24_ = _dafny.Map({d_1519_v6_: d_1537_v23_})
                    d_1539_v25_: _dafny.Map
                    d_1539_v25_ = _dafny.Map({(default__.fm41(d_1519_v6_, len(d_1513_v1_), d_1519_v6_, globalState)).cf31: d_1512_v0_})
                    d_1540_v26_: _dafny.MultiSet
                    d_1540_v26_ = _dafny.MultiSet([len(d_1525_v12_), d_1512_v0_])
                    d_1541_v27_: _dafny.Map
                    d_1541_v27_ = _dafny.Map({len(d_1539_v25_): (d_1540_v26_).cardinality})
                    pat_let_tv58_ = d_1528_v14_
                    pat_let_tv59_ = d_1528_v14_
                    def iife181_(_pat_let62_0):
                        def iife182_(d_1542_dt__update__tmp_h0_):
                            def iife183_(_pat_let63_0):
                                def iife184_(d_1543_dt__update_hcf44_h0_):
                                    return D8_DC23(d_1543_dt__update_hcf44_h0_)
                                return iife184_(_pat_let63_0)
                            return iife183_(pat_let_tv58_)
                        return iife182_(_pat_let62_0)
                    def iife185_(_pat_let64_0):
                        def iife186_(d_1544_dt__update__tmp_h1_):
                            def iife187_(_pat_let65_0):
                                def iife188_(d_1545_dt__update_hcf44_h1_):
                                    return D8_DC23(d_1545_dt__update_hcf44_h1_)
                                return iife188_(_pat_let65_0)
                            return iife187_(pat_let_tv59_)
                        return iife186_(_pat_let64_0)
                    rhs206_ = ((d_1512_v0_) + (len(_dafny.SeqWithoutIsStrInference([(d_1513_v1_)[default__.safeIndex(d_1512_v0_, len(d_1513_v1_))]])))) > ((D2_DC5(d_1519_v6_, len(_dafny.Set({len(((iife181_(d_1535_v21_)).cf44).set(default__.safeIndex(d_1512_v0_, len((iife185_(d_1535_v21_)).cf44)), d_1521_v8_)), d_1512_v0_, d_1512_v0_, len((d_1536_v22_).set(d_1519_v6_, d_1519_v6_)), len(d_1538_v24_)})))).cf13)
                    rhs207_ = (d_1534_v20_) + (default__.fm40(d_1519_v6_, (0) - (default__.fm0(d_1512_v0_, d_1526_v13_, globalState)), d_1539_v25_, globalState))
                    rhs208_ = not(d_1519_v6_)
                    lhs198_ = globalState
                    lhs199_ = globalState
                    lhs198_.f6 = rhs206_
                    d_1534_v20_ = rhs207_
                    lhs199_.f6 = rhs208_
                    d_1546_v28_: _dafny.Seq
                    d_1546_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1519_v6_]), d_1534_v20_])
                    d_1547_v29_: _dafny.Seq
                    d_1547_v29_ = _dafny.SeqWithoutIsStrInference([((d_1546_v28_)[default__.safeIndex((0) - (d_1512_v0_), len(d_1546_v28_))]) + (d_1534_v20_)])
                    d_1547_v29_ = (_dafny.SeqWithoutIsStrInference([p0, d_1534_v20_, d_1534_v20_, p0, p0])) + (d_1547_v29_)
                    if d_1519_v6_:
                        d_1548_v30_: _dafny.Array
                        def lambda66_(d_1549_v6_):
                            def lambda67_(d_1550_i2_):
                                return not(d_1549_v6_)

                            return lambda67_

                        init33_ = lambda66_(d_1519_v6_)
                        nw248_ = _dafny.Array(None, 20)
                        for i0_33_ in range(nw248_.length(0)):
                            nw248_[i0_33_] = init33_(i0_33_)
                        d_1548_v30_ = nw248_
                        d_1551_v31_: _dafny.Map
                        d_1551_v31_ = _dafny.Map({d_1548_v30_: d_1512_v0_})
                        d_1552_v32_: _dafny.MultiSet
                        d_1552_v32_ = _dafny.MultiSet([d_1548_v30_])
                        d_1553_v33_: _dafny.Map
                        d_1553_v33_ = _dafny.Map({d_1519_v6_: d_1552_v32_})
                        rhs209_ = (d_1551_v31_) | ((d_1551_v31_) | (d_1551_v31_))
                        rhs210_ = (d_1552_v32_) | (((d_1553_v33_)[d_1519_v6_] if (d_1519_v6_) in (d_1553_v33_) else d_1552_v32_))
                        rhs211_ = (0) - (-305)
                        lhs200_ = globalState
                        d_1551_v31_ = rhs209_
                        d_1552_v32_ = rhs210_
                        lhs200_.f18 = rhs211_
                        d_1519_v6_ = d_1519_v6_
                        d_1513_v1_ = (_dafny.SeqWithoutIsStrInference([d_1512_v0_ for d_1554_i3_ in range(default__.abs(-333))])) + (d_1513_v1_)
                        (globalState).f17 = d_1512_v0_
                        d_1555_v34_: _dafny.Array
                        nw249_ = _dafny.Array(_dafny.Set({}), 21)
                        d_1555_v34_ = nw249_
                        index250_ = default__.safeIndex(932, (d_1555_v34_).length(0))
                        (d_1555_v34_)[index250_] = d_1525_v12_
                        index251_ = default__.safeIndex(932, (d_1555_v34_).length(0))
                        (d_1555_v34_)[index251_] = _dafny.Set({d_1519_v6_})
                    elif True:
                        (globalState).f18 = (d_1512_v0_) - (d_1512_v0_)
                        d_1556_v35_: C2
                        nw250_ = C2()
                        nw250_.ctor__(len((d_1528_v14_) + (d_1528_v14_)))
                        d_1556_v35_ = nw250_
                        d_1557_v36_: _dafny.Map
                        d_1557_v36_ = _dafny.Map({d_1512_v0_: d_1519_v6_})
                        d_1558_v37_: D5
                        d_1558_v37_ = D5_DC13(d_1557_v36_)
                        d_1559_v38_: D5
                        d_1559_v38_ = D5_DC15(d_1558_v37_)
                        d_1560_v39_: _dafny.Map
                        d_1560_v39_ = _dafny.Map({d_1559_v38_: d_1519_v6_})
                        index252_ = default__.safeIndex(569, (d_1537_v23_).length(0))
                        (d_1537_v23_)[index252_] = (0) - ((len(d_1524_v11_)) - (len(d_1560_v39_)))
                        d_1561_v40_: _dafny.Array
                        def lambda68_(d_1562_v8_):
                            def lambda69_(d_1563_i4_):
                                return d_1562_v8_

                            return lambda69_

                        init34_ = lambda68_(d_1521_v8_)
                        nw251_ = _dafny.Array(None, 10)
                        for i0_34_ in range(nw251_.length(0)):
                            nw251_[i0_34_] = init34_(i0_34_)
                        d_1561_v40_ = nw251_
                        index253_ = default__.safeIndex(538, (d_1561_v40_).length(0))
                        (d_1561_v40_)[index253_] = d_1521_v8_
                        d_1564_v41_: _dafny.Array
                        def lambda70_(d_1565_v6_):
                            def lambda71_(d_1566_i5_):
                                return d_1565_v6_

                            return lambda71_

                        init35_ = lambda70_(d_1519_v6_)
                        nw252_ = _dafny.Array(None, 11)
                        for i0_35_ in range(nw252_.length(0)):
                            nw252_[i0_35_] = init35_(i0_35_)
                        d_1564_v41_ = nw252_
                        index254_ = default__.safeIndex(569, (d_1537_v23_).length(0))
                        index255_ = default__.safeIndex(538, (d_1561_v40_).length(0))
                        rhs212_ = 0
                        rhs213_ = default__.fm1(globalState)
                        rhs214_ = d_1519_v6_
                        rhs215_ = (d_1564_v41_) not in (_dafny.Map({d_1564_v41_: d_1512_v0_}))
                        lhs201_ = d_1537_v23_
                        lhs202_ = default__.safeIndex(569, (d_1537_v23_).length(0))
                        lhs203_ = d_1561_v40_
                        lhs204_ = default__.safeIndex(538, (d_1561_v40_).length(0))
                        lhs205_ = globalState
                        lhs206_ = globalState
                        lhs201_[lhs202_] = rhs212_
                        lhs203_[lhs204_] = rhs213_
                        lhs205_.f6 = rhs214_
                        lhs206_.f6 = rhs215_
                        index256_ = default__.safeIndex(928, (d_1564_v41_).length(0))
                        (d_1564_v41_)[index256_] = d_1519_v6_
                        d_1567_v42_: _dafny.Array
                        def lambda72_(d_1568_v1_):
                            def lambda73_(d_1569_i6_):
                                return D6_DC16(_dafny.Map({d_1568_v1_: not(False)}))

                            return lambda73_

                        init36_ = lambda72_(d_1513_v1_)
                        nw253_ = _dafny.Array(None, 5)
                        for i0_36_ in range(nw253_.length(0)):
                            nw253_[i0_36_] = init36_(i0_36_)
                        d_1567_v42_ = nw253_
                        d_1570_v43_: _dafny.Map
                        d_1570_v43_ = _dafny.Map({d_1513_v1_: d_1519_v6_})
                        d_1571_v44_: D6
                        d_1571_v44_ = D6_DC16(d_1570_v43_)
                        pat_let_tv60_ = d_1570_v43_
                        index257_ = default__.safeIndex(661, (d_1567_v42_).length(0))
                        def iife189_(_pat_let66_0):
                            def iife190_(d_1572_dt__update__tmp_h2_):
                                def iife191_(_pat_let67_0):
                                    def iife192_(d_1573_dt__update_hcf27_h0_):
                                        return D6_DC16(d_1573_dt__update_hcf27_h0_)
                                    return iife192_(_pat_let67_0)
                                return iife191_(pat_let_tv60_)
                            return iife190_(_pat_let66_0)
                        (d_1567_v42_)[index257_] = iife189_(d_1571_v44_)
                        d_1574_v45_: _dafny.Map
                        d_1574_v45_ = _dafny.Map({d_1519_v6_: d_1564_v41_})
                        d_1575_v46_: D9
                        d_1575_v46_ = D9_DC24(d_1574_v45_)
                        d_1576_v47_: _dafny.Array
                        nw254_ = _dafny.Array(None, 8)
                        nw254_[int(0)] = (d_1575_v46_).cf45
                        nw254_[int(1)] = _dafny.Map({d_1519_v6_: ((d_1574_v45_)[d_1519_v6_] if (d_1519_v6_) in (d_1574_v45_) else d_1564_v41_)})
                        nw254_[int(2)] = (d_1574_v45_) | (d_1574_v45_)
                        nw254_[int(3)] = ((d_1574_v45_).set(d_1519_v6_, d_1564_v41_)) | (_dafny.Map({d_1519_v6_: d_1564_v41_}))
                        nw254_[int(4)] = d_1574_v45_
                        nw254_[int(5)] = d_1574_v45_
                        nw254_[int(6)] = d_1574_v45_
                        nw254_[int(7)] = d_1574_v45_
                        d_1576_v47_ = nw254_
                        index258_ = default__.safeIndex(413, (d_1576_v47_).length(0))
                        (d_1576_v47_)[index258_] = (d_1574_v45_) | (_dafny.Map({d_1519_v6_: d_1564_v41_}))
                        pat_let_tv61_ = d_1521_v8_
                        pat_let_tv62_ = d_1519_v6_
                        index259_ = default__.safeIndex(928, (d_1564_v41_).length(0))
                        index260_ = default__.safeIndex(661, (d_1567_v42_).length(0))
                        index261_ = default__.safeIndex(569, (d_1537_v23_).length(0))
                        index262_ = default__.safeIndex(413, (d_1576_v47_).length(0))
                        def iife193_(_pat_let68_0):
                            def iife194_(d_1577_dt__update__tmp_h3_):
                                def iife195_(_pat_let69_0):
                                    def iife196_(d_1578_dt__update_hcf1_h0_):
                                        def iife197_(_pat_let70_0):
                                            def iife198_(d_1579_dt__update_hcf5_h0_):
                                                return D0_DC1(d_1578_dt__update_hcf1_h0_, (d_1577_dt__update__tmp_h3_).cf2, (d_1577_dt__update__tmp_h3_).cf3, (d_1577_dt__update__tmp_h3_).cf4, d_1579_dt__update_hcf5_h0_)
                                            return iife198_(_pat_let70_0)
                                        return iife197_(pat_let_tv62_)
                                    return iife196_(_pat_let69_0)
                                return iife195_(pat_let_tv61_)
                            return iife194_(_pat_let68_0)
                        rhs216_ = not((d_1519_v6_ if (len(_dafny.SeqWithoutIsStrInference([d_1519_v6_, True]))) == ((d_1556_v35_).f22) else d_1519_v6_))
                        rhs217_ = len((d_1539_v25_).set(default__.fm0(d_1512_v0_, iife193_(d_1526_v13_), globalState), d_1512_v0_))
                        rhs218_ = (d_1571_v44_ if d_1519_v6_ else d_1571_v44_)
                        rhs219_ = (d_1537_v23_)[default__.safeIndex(569, (d_1537_v23_).length(0))]
                        rhs220_ = (d_1574_v45_) | (d_1574_v45_)
                        lhs207_ = d_1564_v41_
                        lhs208_ = default__.safeIndex(928, (d_1564_v41_).length(0))
                        lhs209_ = globalState
                        lhs210_ = d_1567_v42_
                        lhs211_ = default__.safeIndex(661, (d_1567_v42_).length(0))
                        lhs212_ = d_1537_v23_
                        lhs213_ = default__.safeIndex(569, (d_1537_v23_).length(0))
                        lhs214_ = d_1576_v47_
                        lhs215_ = default__.safeIndex(413, (d_1576_v47_).length(0))
                        lhs207_[lhs208_] = rhs216_
                        lhs209_.f18 = rhs217_
                        lhs210_[lhs211_] = rhs218_
                        lhs212_[lhs213_] = rhs219_
                        lhs214_[lhs215_] = rhs220_
                        d_1580_v48_: C1
                        nw255_ = C1()
                        nw255_.ctor__((d_1556_v35_).f22, (self).f20)
                        d_1580_v48_ = nw255_
                        d_1581_v49_: _dafny.Array
                        nw256_ = _dafny.Array(_dafny.Array(None, 0), 1)
                        d_1581_v49_ = nw256_
                        d_1582_v50_: _dafny.Array
                        nw257_ = _dafny.Array(int(0), 24)
                        d_1582_v50_ = nw257_
                        d_1583_v51_: _dafny.Array
                        nw258_ = _dafny.Array(None, 13)
                        nw258_[int(0)] = d_1582_v50_
                        nw258_[int(1)] = d_1582_v50_
                        nw258_[int(2)] = d_1582_v50_
                        nw258_[int(3)] = d_1582_v50_
                        nw258_[int(4)] = d_1582_v50_
                        nw258_[int(5)] = d_1582_v50_
                        nw258_[int(6)] = d_1582_v50_
                        nw258_[int(7)] = d_1582_v50_
                        nw258_[int(8)] = d_1582_v50_
                        nw258_[int(9)] = d_1582_v50_
                        nw258_[int(10)] = d_1582_v50_
                        nw258_[int(11)] = d_1582_v50_
                        nw258_[int(12)] = d_1582_v50_
                        d_1583_v51_ = nw258_
                        index263_ = default__.safeIndex(122, (d_1581_v49_).length(0))
                        (d_1581_v49_)[index263_] = d_1583_v51_
                        d_1584_v52_: D10
                        d_1584_v52_ = D10_DC26(_dafny.SeqWithoutIsStrInference([len(d_1525_v12_) for d_1585_i7_ in range(default__.abs(-504))]))
                        d_1586_v53_: _dafny.Seq
                        d_1586_v53_ = _dafny.SeqWithoutIsStrInference([d_1580_v48_])
                        index264_ = default__.safeIndex(122, (d_1581_v49_).length(0))
                        rhs221_ = (d_1584_v52_).cf49
                        rhs222_ = (d_1586_v53_)[default__.safeIndex(len(d_1528_v14_), len(d_1586_v53_))]
                        rhs223_ = (d_1583_v51_ if (d_1564_v41_)[default__.safeIndex(928, (d_1564_v41_).length(0))] else d_1583_v51_)
                        lhs216_ = d_1581_v49_
                        lhs217_ = default__.safeIndex(122, (d_1581_v49_).length(0))
                        d_1513_v1_ = rhs221_
                        d_1580_v48_ = rhs222_
                        lhs216_[lhs217_] = rhs223_
                    pass
            pass
        d_1599_v54_: _dafny.Array
        nw259_ = _dafny.Array(False, 25)
        d_1599_v54_ = nw259_
        d_1600_v55_: C0
        nw260_ = C0()
        nw260_.ctor__(d_1599_v54_, p0)
        d_1600_v55_ = nw260_
        d_1601_v56_: _dafny.Seq
        d_1601_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vwp"))
        d_1602_v57_: D2
        d_1602_v57_ = D2_DC3(d_1601_v56_)
        d_1603_v58_: _dafny.Map
        d_1603_v58_ = _dafny.Map({(d_1519_v6_ if d_1519_v6_ else d_1519_v6_): default__.fm42(d_1602_v57_, not(d_1519_v6_), globalState)})
        d_1603_v58_ = default__.fm43((not(d_1519_v6_)) or (d_1519_v6_), (default__.fm1(globalState) if d_1519_v6_ else d_1521_v8_), globalState)

    def m11(self, p0, p1, globalState):
        r0: bool = False
        d_1604_v0_: int
        d_1604_v0_ = -441
        if (d_1604_v0_) >= (d_1604_v0_):
            if p0:
                d_1605_v1_: _dafny.Seq
                d_1605_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                (globalState).f5 = ((p1) in (d_1605_v1_)) == (p0)
                (globalState).f17 = d_1604_v0_
                d_1606_v2_: _dafny.Map
                d_1606_v2_ = _dafny.Map({len(d_1605_v1_): p0})
                d_1607_v3_: _dafny.Seq
                d_1607_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1608_v4_: D2
                d_1608_v4_ = D2_DC5(False, d_1604_v0_)
                d_1609_v5_: _dafny.Array
                nw261_ = _dafny.Array(None, 6)
                nw261_[int(0)] = p0
                nw261_[int(1)] = p0
                nw261_[int(2)] = True
                nw261_[int(3)] = True
                nw261_[int(4)] = ((d_1606_v2_)[len(d_1607_v3_)] if (len(d_1607_v3_)) in (d_1606_v2_) else (d_1608_v4_).cf12)
                nw261_[int(5)] = p0
                d_1609_v5_ = nw261_
                d_1610_v6_: C0
                nw262_ = C0()
                nw262_.ctor__(d_1609_v5_, (d_1607_v3_) + (_dafny.SeqWithoutIsStrInference([True])))
                d_1610_v6_ = nw262_
                d_1610_v6_ = d_1610_v6_
                d_1611_v7_: T0
                nw263_ = C3()
                nw263_.ctor__((self).f20)
                d_1611_v7_ = nw263_
                d_1611_v7_ = d_1611_v7_
                (globalState).f16 = (d_1605_v1_ if p0 else d_1605_v1_)
            elif True:
                d_1612_v8_: _dafny.Set
                d_1612_v8_ = _dafny.Set({p0, p0})
                d_1613_v9_: C1
                nw264_ = C1()
                nw264_.ctor__(len(d_1612_v8_), (self).f20)
                d_1613_v9_ = nw264_
                d_1614_v10_: D11
                d_1614_v10_ = D11_DC30(d_1613_v9_)
                d_1614_v10_ = d_1614_v10_
                d_1615_v11_: _dafny.Seq
                d_1615_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                d_1616_v12_: D14
                d_1616_v12_ = D14_DC40(p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")), d_1615_v11_, p0)
                d_1617_v13_: _dafny.Array
                nw265_ = _dafny.Array(None, 4)
                nw265_[int(0)] = 724
                nw265_[int(1)] = default__.fm0(d_1604_v0_, D0_DC1(p1, 890, d_1612_v8_, p0, p0), globalState)
                nw265_[int(2)] = d_1604_v0_
                nw265_[int(3)] = d_1604_v0_
                d_1617_v13_ = nw265_
                index265_ = default__.safeIndex(475, (d_1617_v13_).length(0))
                (d_1617_v13_)[index265_] = 767
                pat_let_tv64_ = d_1615_v11_
                pat_let_tv65_ = d_1615_v11_
                pat_let_tv66_ = d_1615_v11_
                index266_ = default__.safeIndex(475, (d_1617_v13_).length(0))
                def iife199_(_pat_let71_0):
                    def iife200_(d_1618_dt__update__tmp_h0_):
                        def iife201_(_pat_let72_0):
                            def iife202_(d_1619_dt__update_hcf74_h0_):
                                def iife203_(_pat_let73_0):
                                    def iife204_(d_1620_dt__update_hcf75_h0_):
                                        return D14_DC40((d_1618_dt__update__tmp_h0_).cf73, d_1619_dt__update_hcf74_h0_, d_1620_dt__update_hcf75_h0_, (d_1618_dt__update__tmp_h0_).cf76)
                                    return iife204_(_pat_let73_0)
                                return iife203_(pat_let_tv66_)
                            return iife202_(_pat_let72_0)
                        return iife201_((pat_let_tv64_) + (pat_let_tv65_))
                    return iife200_(_pat_let71_0)
                rhs224_ = iife199_(d_1616_v12_)
                rhs225_ = d_1604_v0_
                rhs226_ = d_1604_v0_
                rhs227_ = d_1613_v9_
                lhs218_ = d_1617_v13_
                lhs219_ = default__.safeIndex(475, (d_1617_v13_).length(0))
                lhs220_ = globalState
                d_1616_v12_ = rhs224_
                lhs218_[lhs219_] = rhs225_
                lhs220_.f18 = rhs226_
                d_1613_v9_ = rhs227_
                d_1612_v8_ = d_1612_v8_
                d_1621_v14_: str
                d_1621_v14_ = _dafny.CodePoint('x')
                d_1621_v14_ = p1
                (globalState).f6 = (p0) and (p0)
            d_1622_v15_: _dafny.Array
            nw266_ = _dafny.Array(int(0), 23)
            d_1622_v15_ = nw266_
            index267_ = default__.safeIndex(116, (d_1622_v15_).length(0))
            (d_1622_v15_)[index267_] = d_1604_v0_
            index268_ = default__.safeIndex(116, (d_1622_v15_).length(0))
            (d_1622_v15_)[index268_] = (0) - (d_1604_v0_)
            d_1623_v16_: _dafny.Seq
            d_1623_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldhia"))
            d_1604_v0_ = (len((d_1623_v16_) + (d_1623_v16_))) * (((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))]) - ((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))]))
            d_1624_v17_: D2
            d_1624_v17_ = D2_DC5(p0, d_1604_v0_)
            d_1625_v18_: _dafny.Array
            nw267_ = _dafny.Array(None, 9)
            nw267_[int(0)] = p0
            nw267_[int(1)] = False
            nw267_[int(2)] = p0
            nw267_[int(3)] = p0
            nw267_[int(4)] = ((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))]) >= ((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))])
            nw267_[int(5)] = p0
            nw267_[int(6)] = p0
            nw267_[int(7)] = p0
            nw267_[int(8)] = (d_1624_v17_).cf12
            d_1625_v18_ = nw267_
            index269_ = default__.safeIndex(307, (d_1625_v18_).length(0))
            (d_1625_v18_)[index269_] = (not(True)) == (p0)
            index270_ = default__.safeIndex(307, (d_1625_v18_).length(0))
            rhs228_ = ((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))]) + ((d_1622_v15_)[default__.safeIndex(116, (d_1622_v15_).length(0))])
            rhs229_ = False
            rhs230_ = p0
            lhs221_ = globalState
            lhs222_ = d_1625_v18_
            lhs223_ = default__.safeIndex(307, (d_1625_v18_).length(0))
            lhs224_ = globalState
            lhs221_.f17 = rhs228_
            lhs222_[lhs223_] = rhs229_
            lhs224_.f5 = rhs230_
            (globalState).f17 = d_1604_v0_
        elif True:
            (globalState).f6 = not(((d_1604_v0_) < (d_1604_v0_) if (d_1604_v0_) <= (d_1604_v0_) else p0))
            d_1626_v19_: _dafny.Array
            nw268_ = _dafny.Array(int(0), 17)
            d_1626_v19_ = nw268_
            index271_ = default__.safeIndex(384, (d_1626_v19_).length(0))
            (d_1626_v19_)[index271_] = (d_1604_v0_) - (d_1604_v0_)
            d_1627_v20_: _dafny.Set
            d_1627_v20_ = _dafny.Set({p0, False})
            index272_ = default__.safeIndex(384, (d_1626_v19_).length(0))
            (d_1626_v19_)[index272_] = (d_1604_v0_ if p0 else len((d_1627_v20_) - (d_1627_v20_)))
            (globalState).f5 = p0
            d_1628_v21_: _dafny.Map
            d_1628_v21_ = _dafny.Map({d_1604_v0_: p0})
            d_1629_v22_: _dafny.Seq
            d_1629_v22_ = _dafny.SeqWithoutIsStrInference([(d_1628_v21_) | (d_1628_v21_)])
            d_1630_v23_: _dafny.MultiSet
            d_1630_v23_ = _dafny.MultiSet([(d_1626_v19_)[default__.safeIndex(384, (d_1626_v19_).length(0))]])
            d_1631_v24_: _dafny.Seq
            d_1631_v24_ = _dafny.SeqWithoutIsStrInference([p0, False, p0])
            d_1629_v22_ = default__.fm64((d_1630_v23_) | (_dafny.MultiSet([(d_1626_v19_)[default__.safeIndex(384, (d_1626_v19_).length(0))], d_1604_v0_])), False, (d_1631_v24_) + (d_1631_v24_), globalState)
            (globalState).f5 = not(True)
        d_1632_v25_: _dafny.Array
        nw269_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_1632_v25_ = nw269_
        d_1633_v26_: _dafny.Map
        d_1633_v26_ = _dafny.Map({p0: p0})
        d_1634_v27_: _dafny.Array
        nw270_ = _dafny.Array(None, 6)
        d_1634_v27_ = nw270_
        index273_ = default__.safeIndex(587, (d_1632_v25_).length(0))
        (d_1632_v25_)[index273_] = (d_1634_v27_ if ((d_1633_v26_)[p0] if (p0) in (d_1633_v26_) else p0) else d_1634_v27_)
        d_1635_v28_: _dafny.Seq
        d_1635_v28_ = _dafny.SeqWithoutIsStrInference([d_1634_v27_])
        index274_ = default__.safeIndex(587, (d_1632_v25_).length(0))
        (d_1632_v25_)[index274_] = (d_1635_v28_)[default__.safeIndex(d_1604_v0_, len(d_1635_v28_))]
        d_1636_v29_: _dafny.MultiSet
        d_1636_v29_ = _dafny.MultiSet([p0])
        d_1637_v30_: _dafny.Seq
        d_1637_v30_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p0, p0]), d_1636_v29_, d_1636_v29_])
        d_1638_v31_: D17
        d_1638_v31_ = D17_DC46(d_1637_v30_)
        source21_ = d_1638_v31_
        if source21_.is_DC47:
            d_1639___mcc_h0_ = source21_.cf92
            d_1640_cf92_ = d_1639___mcc_h0_
            d_1641_v32_: _dafny.Seq
            d_1641_v32_ = _dafny.SeqWithoutIsStrInference([d_1604_v0_, d_1604_v0_, d_1604_v0_])
            d_1642_v33_: _dafny.Map
            d_1642_v33_ = _dafny.Map({default__.fm1(globalState): p0})
            rhs231_ = 577
            rhs232_ = (_dafny.SeqWithoutIsStrInference([d_1640_cf92_, len(_dafny.SeqWithoutIsStrInference([d_1640_cf92_, d_1604_v0_, d_1604_v0_, d_1604_v0_])), d_1640_cf92_, -957])).set(default__.safeIndex(d_1604_v0_, len(_dafny.SeqWithoutIsStrInference([d_1640_cf92_, len(_dafny.SeqWithoutIsStrInference([d_1640_cf92_, d_1604_v0_, d_1604_v0_, d_1604_v0_])), d_1640_cf92_, -957]))), (0) - (len(d_1642_v33_)))
            d_1640_cf92_ = rhs231_
            d_1641_v32_ = rhs232_
            d_1643_v34_: _dafny.Seq
            d_1643_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axwcrc"))
            (globalState).f16 = (d_1643_v34_) + (d_1643_v34_)
            d_1644_v35_: D3
            d_1644_v35_ = D3_DC8((D5_DC14(p0)).cf25)
            source22_ = D3_DC9(D3_DC9(d_1644_v35_))
            if source22_.is_DC8:
                d_1645___mcc_h4_ = source22_.cf16
                d_1646_cf16_ = d_1645___mcc_h4_
                d_1647_v36_: _dafny.Set
                d_1647_v36_ = _dafny.Set({d_1640_cf92_})
                d_1647_v36_ = d_1647_v36_
                d_1648_v37_: _dafny.Array
                def lambda75_(d_1649_cf92_):
                    def lambda76_(d_1650_i0_):
                        return default__.safeDivisionInt(d_1650_i0_, d_1649_cf92_)

                    return lambda76_

                init37_ = lambda75_(d_1640_cf92_)
                nw271_ = _dafny.Array(None, 13)
                for i0_37_ in range(nw271_.length(0)):
                    nw271_[i0_37_] = init37_(i0_37_)
                d_1648_v37_ = nw271_
                index275_ = default__.safeIndex(161, (d_1648_v37_).length(0))
                (d_1648_v37_)[index275_] = d_1640_cf92_
                index276_ = default__.safeIndex(161, (d_1648_v37_).length(0))
                (d_1648_v37_)[index276_] = d_1640_cf92_
                d_1651_v38_: _dafny.Seq
                d_1651_v38_ = _dafny.SeqWithoutIsStrInference([d_1646_cf16_, p0])
                (globalState).f6 = (d_1651_v38_) != (d_1651_v38_)
                d_1652_v39_: _dafny.Set
                d_1652_v39_ = _dafny.Set({d_1646_cf16_})
                d_1653_v40_: _dafny.Map
                d_1653_v40_ = _dafny.Map({p0: len(d_1652_v39_)})
                d_1653_v40_ = (d_1653_v40_).set(d_1646_cf16_, ((d_1653_v40_)[p0] if (p0) in (d_1653_v40_) else d_1640_cf92_))
            elif source22_.is_DC7:
                d_1654___mcc_h5_ = source22_.cf15
                d_1655_cf15_ = d_1654___mcc_h5_
                d_1656_v41_: D14
                d_1656_v41_ = D14_DC40(p0, d_1643_v34_, d_1643_v34_, not(p0))
                d_1657_v42_: _dafny.Array
                nw272_ = _dafny.Array(None, 22)
                nw272_[int(0)] = d_1643_v34_
                nw272_[int(1)] = (d_1643_v34_) + (d_1643_v34_)
                nw272_[int(2)] = d_1643_v34_
                nw272_[int(3)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajw"))) + (d_1643_v34_)
                nw272_[int(4)] = d_1643_v34_
                nw272_[int(5)] = (d_1643_v34_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbt")))
                nw272_[int(6)] = _dafny.SeqWithoutIsStrInference([p1 for d_1658_i1_ in range(default__.abs(629))])
                nw272_[int(7)] = d_1643_v34_
                nw272_[int(8)] = d_1643_v34_
                nw272_[int(9)] = (d_1643_v34_) + (d_1643_v34_)
                nw272_[int(10)] = d_1643_v34_
                nw272_[int(11)] = ((d_1643_v34_).set(default__.safeIndex(d_1604_v0_, len(d_1643_v34_)), p1)) + (d_1643_v34_)
                nw272_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
                nw272_[int(13)] = d_1643_v34_
                nw272_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mfpc"))
                nw272_[int(15)] = d_1643_v34_
                nw272_[int(16)] = (d_1656_v41_).cf75
                nw272_[int(17)] = default__.fm34(p0, globalState)
                nw272_[int(18)] = d_1643_v34_
                nw272_[int(19)] = _dafny.SeqWithoutIsStrInference([p1 for d_1659_i2_ in range(default__.abs(34))])
                nw272_[int(20)] = (d_1643_v34_ if p0 else d_1643_v34_)
                nw272_[int(21)] = (d_1643_v34_) + (default__.fm46(d_1640_cf92_, p0, p0, globalState))
                d_1657_v42_ = nw272_
                d_1657_v42_ = d_1657_v42_
                d_1660_v43_: _dafny.Set
                d_1660_v43_ = _dafny.Set({p0, p0, p0, not(False), p0})
                d_1661_v44_: D0
                d_1661_v44_ = D0_DC1(p1, d_1640_cf92_, d_1660_v43_, p0, p0)
                d_1640_cf92_ = default__.fm0(d_1604_v0_, d_1661_v44_, globalState)
                d_1662_v45_: _dafny.Array
                nw273_ = _dafny.Array(None, 12)
                nw273_[int(0)] = not(True)
                nw273_[int(1)] = p0
                nw273_[int(2)] = not(p0)
                nw273_[int(3)] = p0
                nw273_[int(4)] = p0
                nw273_[int(5)] = p0
                nw273_[int(6)] = p0
                nw273_[int(7)] = p0
                nw273_[int(8)] = p0
                nw273_[int(9)] = p0
                nw273_[int(10)] = p0
                nw273_[int(11)] = p0
                d_1662_v45_ = nw273_
                d_1663_v46_: _dafny.Map
                d_1663_v46_ = _dafny.Map({d_1604_v0_: d_1662_v45_})
                d_1664_v47_: _dafny.Set
                d_1664_v47_ = _dafny.Set({((d_1663_v46_)[d_1604_v0_] if (d_1604_v0_) in (d_1663_v46_) else d_1662_v45_), ((d_1663_v46_)[len(d_1643_v34_)] if (len(d_1643_v34_)) in (d_1663_v46_) else d_1662_v45_)})
                d_1664_v47_ = d_1664_v47_
                d_1665_v48_: _dafny.Map
                d_1665_v48_ = _dafny.Map({p0: default__.fm0(d_1640_cf92_, d_1661_v44_, globalState)})
                d_1665_v48_ = (d_1665_v48_).set(p0, d_1604_v0_)
            elif True:
                d_1666___mcc_h6_ = source22_.cf17
                d_1667_cf17_ = d_1666___mcc_h6_
                (globalState).f6 = p0
                d_1668_v49_: _dafny.Set
                d_1668_v49_ = _dafny.Set({p0})
                d_1669_v50_: _dafny.Map
                d_1669_v50_ = _dafny.Map({(D0_DC1(p1, d_1640_cf92_, d_1668_v49_, p0, p0)).cf2: (0) - (d_1604_v0_)})
                d_1670_v51_: _dafny.Seq
                d_1670_v51_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1671_v52_: _dafny.Set
                d_1671_v52_ = _dafny.Set({d_1670_v51_})
                d_1669_v50_ = (d_1669_v50_).set(d_1640_cf92_, len(d_1671_v52_))
                d_1672_v53_: _dafny.MultiSet
                d_1672_v53_ = _dafny.MultiSet([p1, p1, p1, p1, p1])
                d_1673_v54_: _dafny.Map
                d_1673_v54_ = _dafny.Map({p0: d_1643_v34_})
                d_1674_v55_: _dafny.Array
                nw274_ = _dafny.Array(None, 23)
                nw274_[int(0)] = (d_1643_v34_).set(default__.safeIndex((0) - (len(d_1670_v51_)), len(d_1643_v34_)), p1)
                nw274_[int(1)] = d_1643_v34_
                nw274_[int(2)] = (default__.fm46(d_1604_v0_, p0, default__.fm2(d_1672_v53_, (0) - (d_1604_v0_), globalState), globalState)) + (d_1643_v34_)
                nw274_[int(3)] = (d_1643_v34_) + (((d_1673_v54_)[p0] if (p0) in (d_1673_v54_) else d_1643_v34_))
                nw274_[int(4)] = d_1643_v34_
                nw274_[int(5)] = (_dafny.SeqWithoutIsStrInference([p1 for d_1675_i3_ in range(default__.abs(435))])) + (d_1643_v34_)
                nw274_[int(6)] = d_1643_v34_
                nw274_[int(7)] = d_1643_v34_
                nw274_[int(8)] = d_1643_v34_
                nw274_[int(9)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnlyxkoeu"))
                nw274_[int(10)] = d_1643_v34_
                nw274_[int(11)] = d_1643_v34_
                nw274_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ful"))).set(default__.safeIndex(d_1640_cf92_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ful")))), p1)
                nw274_[int(13)] = d_1643_v34_
                nw274_[int(14)] = d_1643_v34_
                nw274_[int(15)] = d_1643_v34_
                nw274_[int(16)] = ((d_1643_v34_).set(default__.safeIndex(d_1604_v0_, len(d_1643_v34_)), p1)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwhfwhvp")))
                nw274_[int(17)] = _dafny.SeqWithoutIsStrInference([p1 for d_1676_i4_ in range(default__.abs(562))])
                nw274_[int(18)] = d_1643_v34_
                nw274_[int(19)] = d_1643_v34_
                nw274_[int(20)] = default__.fm34(p0, globalState)
                nw274_[int(21)] = d_1643_v34_
                nw274_[int(22)] = (d_1643_v34_) + (d_1643_v34_)
                d_1674_v55_ = nw274_
                index277_ = default__.safeIndex(565, (d_1674_v55_).length(0))
                (d_1674_v55_)[index277_] = d_1643_v34_
                index278_ = default__.safeIndex(565, (d_1674_v55_).length(0))
                (d_1674_v55_)[index278_] = d_1643_v34_
                d_1677_v56_: _dafny.Array
                nw275_ = _dafny.Array(False, 2)
                d_1677_v56_ = nw275_
                d_1678_v57_: C0
                nw276_ = C0()
                nw276_.ctor__(d_1677_v56_, d_1670_v51_)
                d_1678_v57_ = nw276_
            d_1679_v58_: _dafny.Array
            def lambda77_(d_1680_v0_):
                def lambda78_(d_1681_i5_):
                    return (d_1680_v0_) == (d_1680_v0_)

                return lambda78_

            init38_ = lambda77_(d_1604_v0_)
            nw277_ = _dafny.Array(None, 1)
            for i0_38_ in range(nw277_.length(0)):
                nw277_[i0_38_] = init38_(i0_38_)
            d_1679_v58_ = nw277_
            index279_ = default__.safeIndex(907, (d_1679_v58_).length(0))
            (d_1679_v58_)[index279_] = p0
            index280_ = default__.safeIndex(907, (d_1679_v58_).length(0))
            (d_1679_v58_)[index280_] = p0
        elif source21_.is_DC48:
            d_1682___mcc_h1_ = source21_.cf93
            d_1683___mcc_h2_ = source21_.cf94
            d_1684_cf94_ = d_1683___mcc_h2_
            d_1685_cf93_ = d_1682___mcc_h1_
            d_1686_v59_: _dafny.Seq
            d_1686_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hspwjtm"))
            d_1687_v60_: _dafny.Map
            d_1687_v60_ = _dafny.Map({False: d_1686_v59_})
            d_1688_v61_: _dafny.Array
            nw278_ = _dafny.Array(False, 28)
            d_1688_v61_ = nw278_
            d_1689_v62_: _dafny.Map
            d_1689_v62_ = _dafny.Map({((d_1687_v60_)[d_1684_cf94_] if (d_1684_cf94_) in (d_1687_v60_) else d_1686_v59_): d_1688_v61_})
            d_1690_v63_: D11
            d_1690_v63_ = D11_DC31(d_1689_v62_, d_1604_v0_)
            (globalState).f18 = ((0) - (d_1604_v0_)) + ((d_1690_v63_).cf60)
            (globalState).f6 = d_1685_cf93_
            d_1691_v64_: _dafny.Seq
            d_1691_v64_ = _dafny.SeqWithoutIsStrInference([224])
            (globalState).f5 = (_dafny.CodePoint('y')) not in (((d_1686_v59_).set(default__.safeIndex((d_1691_v64_)[default__.safeIndex(d_1604_v0_, len(d_1691_v64_))], len(d_1686_v59_)), p1)) + (d_1686_v59_))
            r0 = p0
        elif True:
            d_1692___mcc_h3_ = source21_.cf91
            d_1693_cf91_ = d_1692___mcc_h3_
            d_1694_v65_: _dafny.Seq
            d_1694_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bng"))
            (globalState).f15 = (_dafny.SeqWithoutIsStrInference([p1 for d_1695_i6_ in range(default__.abs(577))])) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yupxlchu"))).set(default__.safeIndex(d_1604_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yupxlchu")))), p1)) + (d_1694_v65_))
            d_1696_v66_: C4
            nw279_ = C4()
            nw279_.ctor__((self).f20)
            d_1696_v66_ = nw279_
            if p0:
                d_1697_v67_: _dafny.Set
                d_1697_v67_ = _dafny.Set({p0})
                d_1698_v68_: D0
                d_1698_v68_ = D0_DC1(p1, d_1604_v0_, d_1697_v67_, p0, p0)
                (globalState).f17 = default__.fm0(d_1604_v0_, (D0_DC1(p1, d_1604_v0_, _dafny.Set({p0, p0}), p0, p0) if p0 else d_1698_v68_), globalState)
                (globalState).f18 = (default__.fm0(d_1604_v0_, d_1698_v68_, globalState)) - (len(_dafny.Map({d_1604_v0_: len(d_1694_v65_)})))
                d_1699_v69_: _dafny.Map
                d_1699_v69_ = _dafny.Map({p0: p1})
                d_1700_v70_: _dafny.Map
                d_1700_v70_ = _dafny.Map({p1: d_1699_v69_})
                d_1699_v69_ = ((d_1699_v69_) | (((d_1700_v70_)[p1] if (p1) in (d_1700_v70_) else d_1699_v69_)) if False else (d_1699_v69_) | (d_1699_v69_))
                d_1701_v71_: _dafny.MultiSet
                d_1701_v71_ = _dafny.MultiSet([p1])
                d_1702_v72_: D16
                d_1702_v72_ = D16_DC45(d_1604_v0_, p0, d_1604_v0_)
                d_1703_v73_: _dafny.Array
                nw280_ = _dafny.Array(None, 26)
                nw280_[int(0)] = p0
                nw280_[int(1)] = p0
                nw280_[int(2)] = not(p0)
                nw280_[int(3)] = p0
                nw280_[int(4)] = p0
                nw280_[int(5)] = p0
                nw280_[int(6)] = not(p0)
                nw280_[int(7)] = p0
                nw280_[int(8)] = p0
                nw280_[int(9)] = default__.fm2(d_1701_v71_, d_1604_v0_, globalState)
                nw280_[int(10)] = p0
                nw280_[int(11)] = p0
                nw280_[int(12)] = p0
                nw280_[int(13)] = p0
                nw280_[int(14)] = p0
                nw280_[int(15)] = p0
                nw280_[int(16)] = p0
                nw280_[int(17)] = p0
                nw280_[int(18)] = p0
                nw280_[int(19)] = p0
                nw280_[int(20)] = True
                nw280_[int(21)] = (d_1702_v72_).cf89
                nw280_[int(22)] = p0
                nw280_[int(23)] = p0
                nw280_[int(24)] = p0
                nw280_[int(25)] = not(True)
                d_1703_v73_ = nw280_
                d_1704_v74_: _dafny.Array
                nw281_ = _dafny.Array(None, 6)
                nw281_[int(0)] = d_1703_v73_
                nw281_[int(1)] = d_1703_v73_
                nw281_[int(2)] = d_1703_v73_
                nw281_[int(3)] = d_1703_v73_
                nw281_[int(4)] = d_1703_v73_
                nw281_[int(5)] = d_1703_v73_
                d_1704_v74_ = nw281_
                d_1704_v74_ = d_1704_v74_
                (globalState).f17 = d_1604_v0_
            elif True:
                d_1705_v75_: _dafny.Map
                d_1705_v75_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xofrdjmvx"))): not(p0)})
                d_1705_v75_ = (d_1705_v75_).set(len(_dafny.Map({(0) - ((d_1636_v29_).cardinality): p0})), p0)
                d_1706_v76_: _dafny.Array
                def lambda79_(d_1707_p0_):
                    def lambda80_(d_1708_i7_):
                        return d_1707_p0_

                    return lambda80_

                init39_ = lambda79_(p0)
                nw282_ = _dafny.Array(None, 25)
                for i0_39_ in range(nw282_.length(0)):
                    nw282_[i0_39_] = init39_(i0_39_)
                d_1706_v76_ = nw282_
                d_1709_v77_: _dafny.Seq
                d_1709_v77_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, p0])
                d_1710_v78_: C0
                nw283_ = C0()
                nw283_.ctor__(d_1706_v76_, d_1709_v77_)
                d_1710_v78_ = nw283_
                d_1711_v79_: _dafny.Map
                d_1711_v79_ = _dafny.Map({True: (d_1604_v0_) + (d_1604_v0_)})
                (globalState).f17 = ((d_1711_v79_)[p0] if (p0) in (d_1711_v79_) else d_1604_v0_)
                d_1712_v80_: _dafny.Set
                d_1712_v80_ = _dafny.Set({d_1604_v0_})
                (globalState).f18 = (default__.safeDivisionInt(d_1604_v0_, len(d_1712_v80_))) - (default__.safeModuloInt(d_1604_v0_, d_1604_v0_))
                d_1713_v81_: _dafny.Set
                d_1713_v81_ = _dafny.Set({p0, p0})
                d_1714_v82_: D0
                d_1714_v82_ = D0_DC1(p1, d_1604_v0_, d_1713_v81_, p0, not(p0))
                rhs233_ = ((((d_1694_v65_ if p0 else (d_1694_v65_).set(default__.safeIndex(d_1604_v0_, len(d_1694_v65_)), p1))) + (d_1694_v65_)).set(default__.safeIndex(default__.safeDivisionInt(len(d_1710_v78_.f24), (0) - (len(d_1694_v65_))), len(((d_1694_v65_ if p0 else (d_1694_v65_).set(default__.safeIndex(d_1604_v0_, len(d_1694_v65_)), p1))) + (d_1694_v65_))), p1)).set(default__.safeIndex(d_1604_v0_, len((((d_1694_v65_ if p0 else (d_1694_v65_).set(default__.safeIndex(d_1604_v0_, len(d_1694_v65_)), p1))) + (d_1694_v65_)).set(default__.safeIndex(default__.safeDivisionInt(len(d_1710_v78_.f24), (0) - (len(d_1694_v65_))), len(((d_1694_v65_ if p0 else (d_1694_v65_).set(default__.safeIndex(d_1604_v0_, len(d_1694_v65_)), p1))) + (d_1694_v65_))), p1))), p1)
                rhs234_ = default__.fm0(d_1604_v0_, d_1714_v82_, globalState)
                d_1694_v65_ = rhs233_
                d_1604_v0_ = rhs234_
            d_1715_v83_: C1
            nw284_ = C1()
            nw284_.ctor__(d_1604_v0_, (self).f20)
            d_1715_v83_ = nw284_
        d_1716_v84_: D5
        d_1716_v84_ = D5_DC14(p0)
        d_1717_v85_: D5
        d_1717_v85_ = D5_DC15(d_1716_v84_)
        pat_let_tv67_ = d_1604_v0_
        pat_let_tv68_ = d_1604_v0_
        pat_let_tv69_ = d_1604_v0_
        pat_let_tv70_ = d_1604_v0_
        pat_let_tv71_ = d_1604_v0_
        pat_let_tv72_ = d_1604_v0_
        pat_let_tv73_ = p1
        pat_let_tv74_ = d_1716_v84_
        def lambda81_(source24_):
            if source24_.is_DC14:
                d_1718___mcc_h16_ = source24_.cf25
                d_1719_cf25_ = d_1718___mcc_h16_
                return D10_DC26(_dafny.SeqWithoutIsStrInference([-795, pat_let_tv67_, (D2_DC5(d_1719_cf25_, pat_let_tv68_)).cf13, len(_dafny.SeqWithoutIsStrInference([pat_let_tv69_, pat_let_tv70_])), pat_let_tv71_]))
            elif source24_.is_DC13:
                d_1720___mcc_h17_ = source24_.cf24
                d_1721_cf24_ = d_1720___mcc_h17_
                return D10_DC26(_dafny.SeqWithoutIsStrInference([pat_let_tv72_, len(_dafny.SeqWithoutIsStrInference([pat_let_tv73_ for d_1722_i8_ in range(default__.abs(726))]))]))
            elif True:
                d_1723___mcc_h18_ = source24_.cf26
                d_1724_cf26_ = d_1723___mcc_h18_
                return D10_DC26(_dafny.SeqWithoutIsStrInference([702]))

        def iife205_(_pat_let74_0):
            def iife206_(d_1725_dt__update__tmp_h1_):
                def iife207_(_pat_let75_0):
                    def iife208_(d_1726_dt__update_hcf26_h0_):
                        return D5_DC15(d_1726_dt__update_hcf26_h0_)
                    return iife208_(_pat_let75_0)
                return iife207_(pat_let_tv74_)
            return iife206_(_pat_let74_0)
        source23_ = lambda81_((iife205_(d_1717_v85_) if p0 else d_1717_v85_))
        if source23_.is_DC27:
            d_1727___mcc_h7_ = source23_.cf50
            d_1728___mcc_h8_ = source23_.cf51
            d_1729_cf51_ = d_1728___mcc_h8_
            d_1730_cf50_ = d_1727___mcc_h7_
            (globalState).f5 = (p0) and (p0)
            d_1731_v86_: D0
            d_1731_v86_ = D0_DC1(p1, d_1729_cf51_, _dafny.Set({((d_1633_v26_)[True] if (True) in (d_1633_v26_) else p0)}), p0, True)
            d_1732_v87_: _dafny.Map
            d_1732_v87_ = _dafny.Map({default__.safeDivisionInt(d_1604_v0_, 714): (d_1731_v86_).cf1})
            d_1732_v87_ = d_1732_v87_
            d_1729_cf51_ = d_1729_cf51_
            pat_let_tv75_ = p0
            d_1733_v88_: C4
            nw285_ = C4()
            def iife209_(_pat_let76_0):
                def iife210_(d_1734_dt__update__tmp_h2_):
                    def iife211_(_pat_let77_0):
                        def iife212_(d_1735_dt__update_hcf0_h0_):
                            return D0_DC0(d_1735_dt__update_hcf0_h0_)
                        return iife212_(_pat_let77_0)
                    return iife211_(pat_let_tv75_)
                return iife210_(_pat_let76_0)
            nw285_.ctor__(iife209_((self).f20))
            d_1733_v88_ = nw285_
        elif source23_.is_DC28:
            d_1736___mcc_h9_ = source23_.cf52
            d_1737___mcc_h10_ = source23_.cf53
            d_1738___mcc_h11_ = source23_.cf54
            d_1739___mcc_h12_ = source23_.cf55
            d_1740___mcc_h13_ = source23_.cf56
            d_1741_cf56_ = d_1740___mcc_h13_
            d_1742_cf55_ = d_1739___mcc_h12_
            d_1743_cf54_ = d_1738___mcc_h11_
            d_1744_cf53_ = d_1737___mcc_h10_
            d_1745_cf52_ = d_1736___mcc_h9_
            d_1746_v89_: D5
            d_1746_v89_ = D5_DC14(d_1741_cf56_)
            d_1747_v90_: _dafny.Array
            def lambda82_(d_1748_cf55_):
                def lambda83_(d_1749_i9_):
                    return d_1748_cf55_

                return lambda83_

            init40_ = lambda82_(d_1742_cf55_)
            nw286_ = _dafny.Array(None, 9)
            for i0_40_ in range(nw286_.length(0)):
                nw286_[i0_40_] = init40_(i0_40_)
            d_1747_v90_ = nw286_
            d_1750_v91_: _dafny.Map
            d_1750_v91_ = _dafny.Map({p0: d_1745_cf52_})
            d_1751_v92_: _dafny.Set
            d_1751_v92_ = _dafny.Set({d_1604_v0_, d_1604_v0_})
            rhs235_ = (d_1747_v90_ if False else d_1747_v90_)
            rhs236_ = (((d_1750_v91_)[d_1742_cf55_] if (d_1742_cf55_) in (d_1750_v91_) else d_1745_cf52_)) < ((d_1745_cf52_) + (d_1745_cf52_))
            rhs237_ = D5_DC14(d_1742_cf55_)
            rhs238_ = (d_1751_v92_) == (d_1751_v92_)
            lhs225_ = globalState
            lhs226_ = globalState
            lhs225_.f4 = rhs235_
            d_1743_cf54_ = rhs236_
            d_1746_v89_ = rhs237_
            lhs226_.f5 = rhs238_
            d_1752_v93_: _dafny.MultiSet
            d_1752_v93_ = _dafny.MultiSet([_dafny.CodePoint('n')])
            index281_ = default__.safeIndex(985, (d_1747_v90_).length(0))
            (d_1747_v90_)[index281_] = (default__.fm2(d_1752_v93_, d_1744_cf53_, globalState) if d_1741_cf56_ else not(False))
            index282_ = default__.safeIndex(985, (d_1747_v90_).length(0))
            (d_1747_v90_)[index282_] = default__.fm2(d_1752_v93_, (0) - (default__.safeDivisionInt(d_1744_cf53_, d_1604_v0_)), globalState)
            d_1753_v94_: _dafny.Set
            d_1753_v94_ = _dafny.Set({d_1743_cf54_})
            d_1754_v95_: D0
            d_1754_v95_ = D0_DC1(p1, (0) - (d_1604_v0_), d_1753_v94_, False, d_1743_cf54_)
            rhs239_ = (d_1754_v95_).cf2
            rhs240_ = default__.safeModuloInt(d_1604_v0_, d_1604_v0_)
            lhs227_ = globalState
            d_1604_v0_ = rhs239_
            lhs227_.f17 = rhs240_
            d_1755_v96_: str
            d_1755_v96_ = _dafny.CodePoint('g')
            d_1755_v96_ = d_1755_v96_
        elif source23_.is_DC26:
            d_1756___mcc_h14_ = source23_.cf49
            d_1757_cf49_ = d_1756___mcc_h14_
            (globalState).f17 = d_1604_v0_
            d_1758_v97_: str
            d_1758_v97_ = _dafny.CodePoint('l')
            d_1759_v98_: _dafny.Array
            nw287_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_1759_v98_ = nw287_
            d_1760_v99_: _dafny.Seq
            d_1760_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlkvw"))
            d_1761_v100_: D8
            d_1761_v100_ = D8_DC23(d_1760_v99_)
            index283_ = default__.safeIndex(479, (d_1759_v98_).length(0))
            (d_1759_v98_)[index283_] = (d_1761_v100_).cf44
            d_1762_v101_: D4
            d_1762_v101_ = D4_DC11(p0, p0, d_1604_v0_, d_1604_v0_)
            d_1763_v102_: _dafny.Map
            d_1763_v102_ = _dafny.Map({len(_dafny.Set({False, p0})): d_1604_v0_})
            d_1764_v103_: D14
            d_1764_v103_ = D14_DC40(default__.fm2(_dafny.MultiSet([d_1758_v97_]), len(d_1760_v99_), globalState), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dntthc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsgu")), p0)
            index284_ = default__.safeIndex(479, (d_1759_v98_).length(0))
            rhs241_ = (((d_1762_v101_).cf21) + ((0) - ((0) - (len(d_1763_v102_))))) + (d_1604_v0_)
            rhs242_ = p0
            rhs243_ = p1
            rhs244_ = (d_1760_v99_) + (((d_1764_v103_).cf75).set(default__.safeIndex(d_1604_v0_, len((d_1764_v103_).cf75)), p1))
            lhs228_ = globalState
            lhs229_ = globalState
            lhs230_ = d_1759_v98_
            lhs231_ = default__.safeIndex(479, (d_1759_v98_).length(0))
            lhs228_.f18 = rhs241_
            lhs229_.f6 = rhs242_
            d_1758_v97_ = rhs243_
            lhs230_[lhs231_] = rhs244_
            (globalState).f16 = (d_1759_v98_)[default__.safeIndex(479, (d_1759_v98_).length(0))]
            d_1765_v104_: _dafny.Array
            def lambda84_(d_1766_v0_):
                def lambda85_(d_1767_i10_):
                    return (d_1767_i10_) + (d_1766_v0_)

                return lambda85_

            init41_ = lambda84_(d_1604_v0_)
            nw288_ = _dafny.Array(None, 26)
            for i0_41_ in range(nw288_.length(0)):
                nw288_[i0_41_] = init41_(i0_41_)
            d_1765_v104_ = nw288_
            d_1768_v105_: _dafny.MultiSet
            d_1768_v105_ = _dafny.MultiSet([d_1604_v0_])
            index285_ = default__.safeIndex(258, (d_1765_v104_).length(0))
            (d_1765_v104_)[index285_] = ((d_1768_v105_).intersection(default__.fm56(d_1604_v0_, p0, p0, d_1604_v0_, globalState))).cardinality
            d_1769_v106_: D10
            d_1769_v106_ = D10_DC28((d_1759_v98_)[default__.safeIndex(479, (d_1759_v98_).length(0))], ((d_1768_v105_)[d_1604_v0_] if (d_1604_v0_) in (d_1768_v105_) else d_1604_v0_), p0, p0, p0)
            index286_ = default__.safeIndex(258, (d_1765_v104_).length(0))
            index287_ = default__.safeIndex(479, (d_1759_v98_).length(0))
            rhs245_ = len(_dafny.Set({len((d_1757_cf49_).set(default__.safeIndex(d_1604_v0_, len(d_1757_cf49_)), ((d_1763_v102_)[d_1604_v0_] if (d_1604_v0_) in (d_1763_v102_) else d_1604_v0_))), d_1604_v0_, (d_1604_v0_) + ((d_1769_v106_).cf53), d_1604_v0_, (d_1604_v0_) * (len(d_1757_cf49_))}))
            rhs246_ = (((d_1759_v98_)[default__.safeIndex(479, (d_1759_v98_).length(0))]).set(default__.safeIndex(d_1604_v0_, len((d_1759_v98_)[default__.safeIndex(479, (d_1759_v98_).length(0))])), d_1758_v97_)) + ((d_1759_v98_)[default__.safeIndex(479, (d_1759_v98_).length(0))])
            rhs247_ = d_1604_v0_
            rhs248_ = d_1604_v0_
            rhs249_ = p0
            lhs232_ = d_1765_v104_
            lhs233_ = default__.safeIndex(258, (d_1765_v104_).length(0))
            lhs234_ = d_1759_v98_
            lhs235_ = default__.safeIndex(479, (d_1759_v98_).length(0))
            lhs236_ = globalState
            lhs237_ = globalState
            lhs238_ = globalState
            lhs232_[lhs233_] = rhs245_
            lhs234_[lhs235_] = rhs246_
            lhs236_.f17 = rhs247_
            lhs237_.f18 = rhs248_
            lhs238_.f5 = rhs249_
        elif True:
            d_1770___mcc_h15_ = source23_.cf57
            d_1771_cf57_ = d_1770___mcc_h15_
            (globalState).f18 = ((0) - (d_1604_v0_) if p0 else (0) - ((d_1604_v0_) + (d_1604_v0_)))
            (globalState).f6 = p0
            d_1772_v107_: _dafny.Set
            d_1772_v107_ = _dafny.Set({False, p0})
            (globalState).f6 = (p0) == ((d_1772_v107_) != (d_1772_v107_))
            d_1773_v108_: _dafny.Seq
            d_1773_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfvvh"))
            d_1774_v110_: _dafny.Map
            def iife213_():
                coll57_ = _dafny.Map()
                compr_57_: int
                for compr_57_ in _dafny.IntegerRange(710, 213):
                    d_1775_v109_: int = compr_57_
                    if ((710) <= (d_1775_v109_)) and ((d_1775_v109_) < (213)):
                        coll57_[default__.safeModuloInt(d_1775_v109_, d_1604_v0_)] = p0
                return _dafny.Map(coll57_)
            d_1774_v110_ = _dafny.Map({len(d_1773_v108_): iife213_()
            })
            d_1776_v111_: _dafny.Map
            d_1776_v111_ = _dafny.Map({d_1604_v0_: p0})
            (globalState).f6 = not(((d_1604_v0_) - (d_1604_v0_)) == (len((_dafny.Map({d_1604_v0_: p0})) | (((d_1774_v110_)[d_1604_v0_] if (d_1604_v0_) in (d_1774_v110_) else d_1776_v111_)))))
        d_1777_v112_: _dafny.Array
        nw289_ = _dafny.Array(False, 22)
        d_1777_v112_ = nw289_
        index288_ = default__.safeIndex(598, (d_1777_v112_).length(0))
        (d_1777_v112_)[index288_] = p0
        index289_ = default__.safeIndex(598, (d_1777_v112_).length(0))
        (d_1777_v112_)[index289_] = not(p0)
        d_1604_v0_ = d_1604_v0_
        r0 = p0
        return r0

    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28

class C6(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f20):
        (self)._f20 = f20

    def fm24(self, p0, p1, p2, p3, globalState):
        return not((True if (-661) > (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True])), len((D4_DC10(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1778_i1_ in range(default__.abs(14))]) for d_1779_i0_ in range(default__.abs(748))]))).cf18), len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qo")))})), len(_dafny.Set({_dafny.CodePoint('q')}))}))) else not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1780_i2_ in range(default__.abs(-184))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "flulwi"))))))

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        d_1781_v0_: _dafny.Map
        d_1781_v0_ = _dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1782_i1_ in range(default__.abs(525))]))})
        d_1783_v1_: _dafny.Seq
        d_1783_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivofje"))
        d_1784_v2_: _dafny.Seq
        d_1784_v2_ = _dafny.SeqWithoutIsStrInference([d_1783_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")), d_1783_v1_])
        d_1785_v3_: D4
        d_1785_v3_ = D4_DC10(d_1784_v2_)
        d_1786_v4_: _dafny.Set
        d_1786_v4_ = _dafny.Set({d_1785_v3_, d_1785_v3_})
        hi7_ = default__.safeModuloInt(p1, len(d_1786_v4_))
        for d_1787_i0_ in range((0) - (((d_1781_v0_)[p1] if (p1) in (d_1781_v0_) else p1)), hi7_):
            d_1788_v5_: _dafny.Map
            d_1788_v5_ = _dafny.Map({p0: p1})
            d_1789_v6_: _dafny.MultiSet
            d_1789_v6_ = _dafny.MultiSet([(0) - (((d_1788_v5_)[p0] if (p0) in (d_1788_v5_) else d_1787_i0_))])
            if (d_1789_v6_).isdisjoint(d_1789_v6_):
                (globalState).f5 = (p0) == (p0)
                d_1790_v7_: _dafny.Map
                d_1790_v7_ = _dafny.Map({d_1787_i0_: p0})
                d_1791_v8_: _dafny.Map
                d_1791_v8_ = _dafny.Map({d_1790_v7_: not((self).fm24(p0, p0, d_1787_i0_, p1, globalState))})
                d_1791_v8_ = (d_1791_v8_).set(_dafny.Map({d_1787_i0_: p0}), not(p0))
                (globalState).f6 = not (p0) or (p0)
                d_1792_v9_: _dafny.Seq
                d_1792_v9_ = _dafny.SeqWithoutIsStrInference([d_1789_v6_, d_1789_v6_, d_1789_v6_])
                d_1793_v10_: _dafny.Seq
                d_1793_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1787_i0_]), (d_1792_v9_)[default__.safeIndex(p1, len(d_1792_v9_))]])
                d_1794_v11_: C1
                nw290_ = C1()
                nw290_.ctor__(d_1787_i0_, (self).f20)
                d_1794_v11_ = nw290_
                d_1795_v12_: _dafny.MultiSet
                d_1795_v12_ = _dafny.MultiSet([d_1794_v11_])
                d_1796_v13_: str
                d_1796_v13_ = _dafny.CodePoint('h')
                d_1797_v14_: _dafny.Seq
                d_1797_v14_ = _dafny.SeqWithoutIsStrInference([len(default__.fm25((d_1789_v6_).cardinality, d_1792_v9_, (d_1795_v12_).cardinality, d_1796_v13_, globalState))])
                d_1798_v15_: _dafny.Array
                nw291_ = _dafny.Array(False, 11)
                d_1798_v15_ = nw291_
                index290_ = default__.safeIndex(818, (d_1798_v15_).length(0))
                (d_1798_v15_)[index290_] = p0
                d_1799_v16_: _dafny.Set
                d_1799_v16_ = _dafny.Set({811})
                d_1800_v17_: _dafny.Map
                d_1800_v17_ = _dafny.Map({d_1799_v16_: p1})
                index291_ = default__.safeIndex(818, (d_1798_v15_).length(0))
                rhs250_ = default__.fm26((p0 if p0 else p0), globalState)
                rhs251_ = p0
                rhs252_ = default__.fm27((0) - (p1), d_1789_v6_, (len(d_1799_v16_) if p0 else d_1787_i0_), (default__.fm28(globalState)) | (d_1800_v17_), globalState)
                rhs253_ = not(not((not(p0)) and (p0)))
                lhs239_ = globalState
                lhs240_ = globalState
                lhs241_ = d_1798_v15_
                lhs242_ = default__.safeIndex(818, (d_1798_v15_).length(0))
                d_1797_v14_ = rhs250_
                lhs239_.f5 = rhs251_
                lhs240_.f16 = rhs252_
                lhs241_[lhs242_] = rhs253_
                d_1801_v18_: D2
                d_1801_v18_ = D2_DC4(d_1787_i0_, p0, True, (d_1798_v15_)[default__.safeIndex(818, (d_1798_v15_).length(0))])
                (globalState).f17 = (d_1801_v18_).cf8
            elif True:
                d_1802_v19_: str
                d_1802_v19_ = _dafny.CodePoint('s')
                rhs254_ = (d_1783_v1_) + (d_1783_v1_)
                rhs255_ = d_1802_v19_
                lhs243_ = globalState
                lhs243_.f16 = rhs254_
                r2 = rhs255_
                d_1803_v20_: _dafny.Set
                d_1803_v20_ = _dafny.Set({p1, d_1787_i0_})
                d_1804_v21_: _dafny.Seq
                d_1804_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: len(d_1803_v20_)}), d_1781_v0_, d_1781_v0_])
                d_1804_v21_ = (d_1804_v21_) + (_dafny.SeqWithoutIsStrInference([d_1781_v0_]))
                (globalState).f17 = d_1787_i0_
                d_1805_v22_: _dafny.Array
                nw292_ = _dafny.Array(False, 19)
                d_1805_v22_ = nw292_
                (globalState).f4 = d_1805_v22_
                d_1806_v23_: _dafny.Set
                d_1806_v23_ = _dafny.Set({p0, p0})
                d_1806_v23_ = (d_1806_v23_) - ((_dafny.Set({p0})) - (d_1806_v23_))
            d_1807_v24_: _dafny.Array
            def lambda86_(d_1808_i2_):
                return _dafny.CodePoint('a')

            init42_ = lambda86_
            nw293_ = _dafny.Array(None, 3)
            for i0_42_ in range(nw293_.length(0)):
                nw293_[i0_42_] = init42_(i0_42_)
            d_1807_v24_ = nw293_
            d_1807_v24_ = d_1807_v24_
            d_1809_v25_: _dafny.Map
            d_1809_v25_ = _dafny.Map({d_1787_i0_: p0})
            (globalState).f17 = (len(d_1809_v25_)) + (p1)
            d_1810_v26_: _dafny.Map
            d_1810_v26_ = _dafny.Map({_dafny.Set({d_1787_i0_, p1}): p0})
            d_1811_v27_: _dafny.Set
            d_1811_v27_ = _dafny.Set({p1})
            d_1810_v26_ = (_dafny.Map({d_1811_v27_: p0})) | (d_1810_v26_)
        d_1812_v29_: _dafny.Set
        d_1812_v29_ = _dafny.Set({p1})
        def iife214_():
            coll58_ = _dafny.Set()
            compr_58_: int
            for compr_58_ in _dafny.IntegerRange(96, 218):
                d_1813_v28_: int = compr_58_
                if ((96) <= (d_1813_v28_)) and ((d_1813_v28_) < (218)):
                    coll58_ = coll58_.union(_dafny.Set([(d_1813_v28_) * ((_dafny.MultiSet([_dafny.CodePoint('x')])).cardinality)]))
            return _dafny.Set(coll58_)
        if (d_1812_v29_).ispropersubset(iife214_()
        ):
            d_1814_v30_: _dafny.Seq
            d_1814_v30_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1815_v31_: _dafny.Map
            d_1815_v31_ = _dafny.Map({p0: p0})
            d_1816_v32_: _dafny.Array
            nw294_ = _dafny.Array(None, 20)
            nw294_[int(0)] = (d_1814_v30_)[default__.safeIndex(p1, len(d_1814_v30_))]
            nw294_[int(1)] = ((d_1815_v31_)[p0] if (p0) in (d_1815_v31_) else p0)
            nw294_[int(2)] = p0
            nw294_[int(3)] = p0
            nw294_[int(4)] = p0
            nw294_[int(5)] = p0
            nw294_[int(6)] = p0
            nw294_[int(7)] = p0
            nw294_[int(8)] = p0
            nw294_[int(9)] = p0
            nw294_[int(10)] = p0
            nw294_[int(11)] = p0
            nw294_[int(12)] = False
            nw294_[int(13)] = p0
            nw294_[int(14)] = p0
            nw294_[int(15)] = p0
            nw294_[int(16)] = True
            nw294_[int(17)] = p0
            nw294_[int(18)] = p0
            nw294_[int(19)] = p0
            d_1816_v32_ = nw294_
            (globalState).f17 = default__.safeModuloInt(default__.safeDivisionInt(p1, p1), len(_dafny.SeqWithoutIsStrInference([d_1816_v32_])))
            d_1817_v33_: _dafny.Array
            nw295_ = _dafny.Array(_dafny.Map({}), 27)
            d_1817_v33_ = nw295_
            d_1818_v34_: _dafny.Map
            d_1818_v34_ = _dafny.Map({d_1783_v1_: d_1781_v0_})
            index292_ = default__.safeIndex(813, (d_1817_v33_).length(0))
            (d_1817_v33_)[index292_] = d_1818_v34_
            d_1819_v35_: str
            d_1819_v35_ = _dafny.CodePoint('r')
            d_1820_v36_: _dafny.Set
            d_1820_v36_ = _dafny.Set({p0, p0})
            index293_ = default__.safeIndex(813, (d_1817_v33_).length(0))
            (d_1817_v33_)[index293_] = (d_1818_v34_).set(d_1783_v1_, _dafny.Map({default__.fm0(p1, D0_DC1(d_1819_v35_, 355, d_1820_v36_, p0, not(p0)), globalState): 353}))
            d_1814_v30_ = d_1814_v30_
            (globalState).f6 = not(p0)
            d_1819_v35_ = d_1819_v35_
        elif True:
            d_1821_v37_: T1
            d_1822_v38_: bool
            d_1823_v39_: _dafny.Array
            out13_: T1
            out14_: bool
            out15_: _dafny.Array
            out13_, out14_, out15_ = (self).m8(p1, p0, p1, globalState)
            d_1821_v37_ = out13_
            d_1822_v38_ = out14_
            d_1823_v39_ = out15_
            d_1824_v40_: _dafny.Map
            d_1824_v40_ = _dafny.Map({(d_1821_v37_).f21: d_1822_v38_})
            d_1824_v40_ = (d_1824_v40_).set(default__.safeModuloInt((0) - (len(d_1783_v1_)), (d_1821_v37_).f21), not(p0))
            d_1825_v41_: _dafny.Array
            nw296_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1825_v41_ = nw296_
            d_1826_v42_: _dafny.Array
            nw297_ = _dafny.Array(_dafny.MultiSet({}), 16)
            d_1826_v42_ = nw297_
            index294_ = default__.safeIndex(565, (d_1825_v41_).length(0))
            (d_1825_v41_)[index294_] = d_1826_v42_
            d_1827_v43_: _dafny.Seq
            d_1827_v43_ = _dafny.SeqWithoutIsStrInference([d_1826_v42_, d_1826_v42_, d_1826_v42_, d_1826_v42_, d_1826_v42_])
            index295_ = default__.safeIndex(565, (d_1825_v41_).length(0))
            (d_1825_v41_)[index295_] = (d_1827_v43_)[default__.safeIndex((p1) + (77), len(d_1827_v43_))]
            d_1828_v44_: str
            d_1828_v44_ = _dafny.CodePoint('d')
            d_1829_v45_: _dafny.Map
            d_1829_v45_ = _dafny.Map({d_1822_v38_: d_1828_v44_})
            index296_ = default__.safeIndex(709, (d_1823_v39_).length(0))
            (d_1823_v39_)[index296_] = len(d_1829_v45_)
            index297_ = default__.safeIndex(709, (d_1823_v39_).length(0))
            (d_1823_v39_)[index297_] = ((d_1821_v37_).f21) + ((d_1821_v37_).f21)
            d_1830_v46_: _dafny.Seq
            d_1830_v46_ = _dafny.SeqWithoutIsStrInference([p0, not(((d_1821_v37_).f21) > ((d_1821_v37_).f21)), False, not (p0) or (d_1822_v38_)])
            d_1830_v46_ = default__.fm29(globalState)
        r1 = (p1) + (384)
        d_1831_v47_: _dafny.Seq
        d_1831_v47_ = _dafny.SeqWithoutIsStrInference([(p1 if p0 else -233), p1])
        d_1831_v47_ = d_1831_v47_
        d_1832_v48_: _dafny.Array
        nw298_ = _dafny.Array(False, 19)
        d_1832_v48_ = nw298_
        index298_ = default__.safeIndex(659, (d_1832_v48_).length(0))
        (d_1832_v48_)[index298_] = (471) != (p1)
        d_1833_v49_: _dafny.Set
        d_1833_v49_ = _dafny.Set({d_1831_v47_, (d_1831_v47_) + (d_1831_v47_)})
        d_1834_v50_: str
        d_1834_v50_ = _dafny.CodePoint('a')
        d_1835_v51_: _dafny.MultiSet
        d_1835_v51_ = _dafny.MultiSet([d_1834_v50_])
        d_1836_v52_: _dafny.Set
        d_1836_v52_ = _dafny.Set({p0, default__.fm2((d_1835_v51_).set(d_1834_v50_, default__.abs(966)), p1, globalState)})
        d_1837_v53_: D2
        d_1837_v53_ = D2_DC5(p0, default__.fm0(p1, D0_DC1(d_1834_v50_, (0) - (p1), d_1836_v52_, p0, not(p0)), globalState))
        d_1838_v54_: D2
        d_1838_v54_ = D2_DC6(d_1837_v53_)
        d_1839_v55_: _dafny.Map
        d_1839_v55_ = _dafny.Map({True: p0})
        d_1840_v56_: _dafny.Seq
        d_1840_v56_ = _dafny.SeqWithoutIsStrInference([((d_1839_v55_)[False] if (False) in (d_1839_v55_) else p0)])
        d_1841_v57_: _dafny.MultiSet
        d_1841_v57_ = _dafny.MultiSet([not(p0), (d_1840_v56_)[default__.safeIndex(p1, len(d_1840_v56_))]])
        pat_let_tv76_ = p0
        pat_let_tv77_ = p0
        pat_let_tv78_ = p0
        index299_ = default__.safeIndex(659, (d_1832_v48_).length(0))
        def lambda87_(source25_):
            if source25_.is_DC4:
                d_1842___mcc_h0_ = source25_.cf8
                d_1843___mcc_h1_ = source25_.cf9
                d_1844___mcc_h2_ = source25_.cf10
                d_1845___mcc_h3_ = source25_.cf11
                d_1846_cf11_ = d_1845___mcc_h3_
                d_1847_cf10_ = d_1844___mcc_h2_
                d_1848_cf9_ = d_1843___mcc_h1_
                d_1849_cf8_ = d_1842___mcc_h0_
                return True
            elif source25_.is_DC5:
                d_1850___mcc_h4_ = source25_.cf12
                d_1851___mcc_h5_ = source25_.cf13
                d_1852_cf13_ = d_1851___mcc_h5_
                d_1853_cf12_ = d_1850___mcc_h4_
                return pat_let_tv76_
            elif source25_.is_DC3:
                d_1854___mcc_h6_ = source25_.cf7
                d_1855_cf7_ = d_1854___mcc_h6_
                return pat_let_tv77_
            elif True:
                d_1856___mcc_h7_ = source25_.cf14
                d_1857_cf14_ = d_1856___mcc_h7_
                return pat_let_tv78_

        rhs256_ = (p1) * (p1)
        rhs257_ = (p1) < (p1)
        rhs258_ = lambda87_(d_1838_v54_)
        rhs259_ = d_1833_v49_
        rhs260_ = d_1841_v57_
        lhs244_ = globalState
        lhs245_ = d_1832_v48_
        lhs246_ = default__.safeIndex(659, (d_1832_v48_).length(0))
        r1 = rhs256_
        lhs244_.f5 = rhs257_
        lhs245_[lhs246_] = rhs258_
        d_1833_v49_ = rhs259_
        r0 = rhs260_
        d_1858_v58_: _dafny.Array
        nw299_ = _dafny.Array(_dafny.Set({}), 9)
        d_1858_v58_ = nw299_
        d_1859_v59_: _dafny.Set
        d_1859_v59_ = _dafny.Set({d_1832_v48_})
        index300_ = default__.safeIndex(439, (d_1858_v58_).length(0))
        (d_1858_v58_)[index300_] = d_1859_v59_
        index301_ = default__.safeIndex(439, (d_1858_v58_).length(0))
        (d_1858_v58_)[index301_] = _dafny.Set({d_1832_v48_, d_1832_v48_, (d_1832_v48_ if p0 else d_1832_v48_), d_1832_v48_})
        r0 = ((d_1841_v57_).intersection(d_1841_v57_)) | (default__.fm30(p0, d_1783_v1_, globalState))
        r1 = p1
        r2 = d_1834_v50_
        return r0, r1, r2

    def m8(self, p0, p1, p2, globalState):
        r0: T1 = None
        r1: bool = False
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1860_i0_: int
        d_1860_i0_ = 0
        with _dafny.label("8"):
            while p1:
                with _dafny.c_label("8"):
                    if (d_1860_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_1860_i0_ = (d_1860_i0_) + (1)
                    d_1861_v0_: _dafny.Seq
                    d_1861_v0_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                    d_1862_v1_: _dafny.MultiSet
                    d_1862_v1_ = _dafny.MultiSet([p1, (d_1861_v0_)[default__.safeIndex(p0, len(d_1861_v0_))]])
                    d_1862_v1_ = ((d_1862_v1_ if True else _dafny.MultiSet([p1]))) - (d_1862_v1_)
                    (globalState).f6 = (not(True)) and (p1)
                    d_1863_v2_: _dafny.Seq
                    d_1863_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bc"))
                    d_1864_v3_: _dafny.Map
                    d_1864_v3_ = _dafny.Map({d_1863_v2_: p1})
                    d_1865_v4_: str
                    d_1865_v4_ = _dafny.CodePoint('d')
                    d_1866_v5_: _dafny.Map
                    d_1866_v5_ = _dafny.Map({d_1865_v4_: p1})
                    d_1864_v3_ = (d_1864_v3_).set(d_1863_v2_, ((d_1866_v5_)[d_1865_v4_] if (d_1865_v4_) in (d_1866_v5_) else p1))
                    d_1867_v6_: _dafny.Map
                    d_1867_v6_ = _dafny.Map({p1: False})
                    d_1868_v7_: _dafny.Map
                    d_1868_v7_ = _dafny.Map({d_1867_v6_: d_1863_v2_})
                    (globalState).f18 = len(default__.fm31(default__.fm2(_dafny.MultiSet(((d_1868_v7_)[(d_1867_v6_).set(p1, p1)] if ((d_1867_v6_).set(p1, p1)) in (d_1868_v7_) else d_1863_v2_)), p2, globalState), (_dafny.SeqWithoutIsStrInference([p0])) + (_dafny.SeqWithoutIsStrInference([-810, p2, p2])), globalState))
                    pass
            pass
        (globalState).f17 = 691
        d_1869_v8_: _dafny.Set
        d_1869_v8_ = _dafny.Set({p1, not(p1)})
        (globalState).f17 = (0) - (((0) - (default__.safeDivisionInt(p2, len(d_1869_v8_)))) - (207))
        if p1:
            d_1870_v9_: _dafny.Array
            def lambda88_(d_1871_p0_):
                def lambda89_(d_1872_i1_):
                    return _dafny.SeqWithoutIsStrInference([d_1871_p0_])

                return lambda89_

            init43_ = lambda88_(p0)
            nw300_ = _dafny.Array(None, 21)
            for i0_43_ in range(nw300_.length(0)):
                nw300_[i0_43_] = init43_(i0_43_)
            d_1870_v9_ = nw300_
            d_1873_v10_: D3
            d_1873_v10_ = D3_DC7(d_1870_v9_)
            d_1874_v11_: _dafny.MultiSet
            d_1874_v11_ = _dafny.MultiSet([D3_DC7(d_1870_v9_), D3_DC7(d_1870_v9_), d_1873_v10_, D3_DC7((d_1873_v10_).cf15)])
            d_1875_v12_: str
            d_1875_v12_ = _dafny.CodePoint('o')
            d_1876_v13_: _dafny.Seq
            d_1876_v13_ = _dafny.SeqWithoutIsStrInference([p1, True])
            d_1877_v14_: D0
            d_1877_v14_ = D0_DC1(d_1875_v12_, p2, _dafny.Set({True, p1}), (d_1876_v13_)[default__.safeIndex(p2, len(d_1876_v13_))], not(p1))
            d_1878_v15_: _dafny.Map
            d_1878_v15_ = _dafny.Map({p2: p0})
            d_1879_v16_: _dafny.Array
            nw301_ = _dafny.Array(None, 16)
            nw301_[int(0)] = p2
            nw301_[int(1)] = p0
            nw301_[int(2)] = p0
            nw301_[int(3)] = p0
            nw301_[int(4)] = p2
            nw301_[int(5)] = (0) - (default__.fm0(p2, d_1877_v14_, globalState))
            nw301_[int(6)] = p2
            nw301_[int(7)] = p2
            nw301_[int(8)] = p2
            nw301_[int(9)] = p0
            nw301_[int(10)] = p2
            nw301_[int(11)] = 747
            nw301_[int(12)] = len(_dafny.SeqWithoutIsStrInference([len(d_1878_v15_)]))
            nw301_[int(13)] = (d_1877_v14_).cf2
            nw301_[int(14)] = p2
            nw301_[int(15)] = p0
            d_1879_v16_ = nw301_
            d_1880_v17_: _dafny.Map
            d_1880_v17_ = _dafny.Map({(p1 if p1 else p1): d_1879_v16_})
            d_1881_v18_: _dafny.Array
            def lambda90_(d_1882_v8_, d_1883_p0_):
                def lambda91_(d_1884_i2_):
                    return (len(_dafny.Set({d_1882_v8_, d_1882_v8_}))) == (d_1883_p0_)

                return lambda91_

            init44_ = lambda90_(d_1869_v8_, p0)
            nw302_ = _dafny.Array(None, 10)
            for i0_44_ in range(nw302_.length(0)):
                nw302_[i0_44_] = init44_(i0_44_)
            d_1881_v18_ = nw302_
            index302_ = default__.safeIndex(15, (d_1881_v18_).length(0))
            (d_1881_v18_)[index302_] = (d_1877_v14_).cf5
            d_1885_v19_: _dafny.Map
            d_1885_v19_ = _dafny.Map({p1: len(d_1876_v13_)})
            d_1886_v20_: _dafny.MultiSet
            d_1886_v20_ = _dafny.MultiSet([(d_1885_v19_).set(p1, p0)])
            index303_ = default__.safeIndex(15, (d_1881_v18_).length(0))
            rhs261_ = ((d_1874_v11_) | (d_1874_v11_)) | (d_1874_v11_)
            rhs262_ = (d_1886_v20_).cardinality
            rhs263_ = (d_1880_v17_ if p1 else (d_1880_v17_) | (d_1880_v17_))
            rhs264_ = p1
            rhs265_ = 875
            lhs247_ = globalState
            lhs248_ = d_1881_v18_
            lhs249_ = default__.safeIndex(15, (d_1881_v18_).length(0))
            lhs250_ = globalState
            d_1874_v11_ = rhs261_
            lhs247_.f17 = rhs262_
            d_1880_v17_ = rhs263_
            lhs248_[lhs249_] = rhs264_
            lhs250_.f18 = rhs265_
            d_1887_v21_: _dafny.Seq
            d_1887_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tjrergf"))
            d_1888_v22_: D2
            d_1888_v22_ = D2_DC3(d_1887_v21_)
            source26_ = d_1888_v22_
            if source26_.is_DC4:
                d_1889___mcc_h0_ = source26_.cf8
                d_1890___mcc_h1_ = source26_.cf9
                d_1891___mcc_h2_ = source26_.cf10
                d_1892___mcc_h3_ = source26_.cf11
                d_1893_cf11_ = d_1892___mcc_h3_
                d_1894_cf10_ = d_1891___mcc_h2_
                d_1895_cf9_ = d_1890___mcc_h1_
                d_1896_cf8_ = d_1889___mcc_h0_
                d_1885_v19_ = (d_1885_v19_).set(False, p0)
                d_1897_v23_: _dafny.MultiSet
                d_1897_v23_ = _dafny.MultiSet([d_1894_cf10_, not(p1)])
                d_1898_v24_: _dafny.Map
                d_1898_v24_ = _dafny.Map({d_1897_v23_: p2})
                d_1898_v24_ = (d_1898_v24_).set(d_1897_v23_, p0)
                d_1899_v25_: _dafny.Seq
                d_1899_v25_ = _dafny.SeqWithoutIsStrInference([p0, default__.safeModuloInt(p0, d_1896_cf8_)])
                d_1900_v26_: D2
                d_1900_v26_ = D2_DC4(p0, d_1893_cf11_, d_1895_cf9_, d_1893_cf11_)
                rhs266_ = len((d_1899_v25_) + (d_1899_v25_))
                rhs267_ = (d_1899_v25_) + (_dafny.SeqWithoutIsStrInference([(d_1900_v26_).cf8, p2]))
                rhs268_ = d_1875_v12_
                rhs269_ = default__.fm0(default__.fm0(default__.fm0(922, d_1877_v14_, globalState), d_1877_v14_, globalState), d_1877_v14_, globalState)
                rhs270_ = d_1881_v18_
                lhs251_ = globalState
                lhs252_ = globalState
                lhs253_ = globalState
                lhs251_.f17 = rhs266_
                d_1899_v25_ = rhs267_
                d_1875_v12_ = rhs268_
                lhs252_.f17 = rhs269_
                lhs253_.f4 = rhs270_
                d_1893_cf11_ = (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]
            elif source26_.is_DC5:
                d_1901___mcc_h4_ = source26_.cf12
                d_1902___mcc_h5_ = source26_.cf13
                d_1903_cf13_ = d_1902___mcc_h5_
                d_1904_cf12_ = d_1901___mcc_h4_
                d_1876_v13_ = d_1876_v13_
                nw303_ = _dafny.Array(None, 8)
                nw303_[int(0)] = not ((d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]) or (d_1904_cf12_)
                nw303_[int(1)] = (default__.fm2(_dafny.MultiSet(d_1887_v21_), d_1903_cf13_, globalState) if False else (self).fm24(p1, p1, p0, p0, globalState))
                nw303_[int(2)] = ((d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))] if (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))] else p1)
                nw303_[int(3)] = (d_1877_v14_).cf4
                nw303_[int(4)] = (d_1903_cf13_) != (default__.fm0(p0, D0_DC1(d_1875_v12_, d_1903_cf13_, d_1869_v8_, (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))], p1), globalState))
                nw303_[int(5)] = not(p1)
                nw303_[int(6)] = p1
                nw303_[int(7)] = (d_1904_cf12_) == (d_1904_cf12_)
                (globalState).f4 = nw303_
                d_1905_v27_: _dafny.Set
                d_1905_v27_ = _dafny.Set({len(d_1878_v15_), p0})
                rhs271_ = default__.safeModuloInt(len(d_1887_v21_), 989)
                rhs272_ = (d_1905_v27_).ispropersubset(_dafny.Set({d_1903_cf13_, p2, p0}))
                lhs254_ = globalState
                lhs255_ = globalState
                lhs254_.f17 = rhs271_
                lhs255_.f6 = rhs272_
                (globalState).f17 = p0
            elif source26_.is_DC3:
                d_1906___mcc_h6_ = source26_.cf7
                d_1907_cf7_ = d_1906___mcc_h6_
                index304_ = default__.safeIndex(617, (d_1879_v16_).length(0))
                (d_1879_v16_)[index304_] = len((d_1907_cf7_) + (d_1907_cf7_))
                index305_ = default__.safeIndex(617, (d_1879_v16_).length(0))
                (d_1879_v16_)[index305_] = p0
                d_1908_v28_: _dafny.MultiSet
                d_1908_v28_ = _dafny.MultiSet([d_1875_v12_])
                index306_ = default__.safeIndex(15, (d_1881_v18_).length(0))
                (d_1881_v18_)[index306_] = default__.fm2(d_1908_v28_, p0, globalState)
                d_1909_v29_: _dafny.Array
                nw304_ = _dafny.Array(int(0), 23)
                d_1909_v29_ = nw304_
                d_1910_v30_: _dafny.MultiSet
                d_1910_v30_ = _dafny.MultiSet([d_1909_v29_, d_1909_v29_, d_1909_v29_, d_1909_v29_])
                d_1910_v30_ = d_1910_v30_
                d_1911_v31_: D2
                d_1911_v31_ = D2_DC4((d_1879_v16_)[default__.safeIndex(617, (d_1879_v16_).length(0))], (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))], (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))], p1)
                d_1912_v32_: _dafny.Seq
                d_1912_v32_ = _dafny.SeqWithoutIsStrInference([d_1881_v18_, d_1881_v18_])
                d_1913_v33_: _dafny.Map
                d_1913_v33_ = _dafny.Map({p2: (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]})
                pat_let_tv79_ = d_1881_v18_
                pat_let_tv80_ = d_1881_v18_
                pat_let_tv81_ = p1
                pat_let_tv82_ = d_1881_v18_
                pat_let_tv83_ = d_1881_v18_
                pat_let_tv84_ = p1
                pat_let_tv85_ = p1
                pat_let_tv86_ = d_1912_v32_
                d_1914_v34_: _dafny.Array
                nw305_ = _dafny.Array(None, 21)
                nw305_[int(0)] = D2_DC4(default__.fm0(p0, d_1877_v14_, globalState), p1, p1, p1)
                nw305_[int(1)] = d_1911_v31_
                def iife215_(_pat_let78_0):
                    def iife216_(d_1915_dt__update__tmp_h0_):
                        def iife217_(_pat_let79_0):
                            def iife218_(d_1916_dt__update_hcf9_h0_):
                                def iife219_(_pat_let80_0):
                                    def iife220_(d_1917_dt__update_hcf10_h0_):
                                        return D2_DC4((d_1915_dt__update__tmp_h0_).cf8, d_1916_dt__update_hcf9_h0_, d_1917_dt__update_hcf10_h0_, (d_1915_dt__update__tmp_h0_).cf11)
                                    return iife220_(_pat_let80_0)
                                return iife219_(pat_let_tv81_)
                            return iife218_(_pat_let79_0)
                        return iife217_((pat_let_tv80_)[default__.safeIndex(15, (pat_let_tv79_).length(0))])
                    return iife216_(_pat_let78_0)
                nw305_[int(2)] = iife215_(d_1911_v31_)
                def iife221_(_pat_let81_0):
                    def iife222_(d_1918_dt__update__tmp_h1_):
                        def iife223_(_pat_let82_0):
                            def iife224_(d_1919_dt__update_hcf11_h0_):
                                return D2_DC4((d_1918_dt__update__tmp_h1_).cf8, (d_1918_dt__update__tmp_h1_).cf9, (d_1918_dt__update__tmp_h1_).cf10, d_1919_dt__update_hcf11_h0_)
                            return iife224_(_pat_let82_0)
                        return iife223_((pat_let_tv83_)[default__.safeIndex(15, (pat_let_tv82_).length(0))])
                    return iife222_(_pat_let81_0)
                nw305_[int(3)] = iife221_(d_1911_v31_)
                nw305_[int(4)] = d_1911_v31_
                nw305_[int(5)] = d_1911_v31_
                nw305_[int(6)] = D2_DC4(len(d_1907_cf7_), (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))], p1, (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))])
                nw305_[int(7)] = d_1911_v31_
                def iife226_(_pat_let84_0):
                    def iife227_(d_1920_dt__update__tmp_h2_):
                        def iife228_(_pat_let85_0):
                            def iife229_(d_1921_dt__update_hcf11_h1_):
                                def iife230_(_pat_let86_0):
                                    def iife231_(d_1922_dt__update_hcf10_h1_):
                                        return D2_DC4((d_1920_dt__update__tmp_h2_).cf8, (d_1920_dt__update__tmp_h2_).cf9, d_1922_dt__update_hcf10_h1_, d_1921_dt__update_hcf11_h1_)
                                    return iife231_(_pat_let86_0)
                                return iife230_(pat_let_tv85_)
                            return iife229_(_pat_let85_0)
                        return iife228_(not(pat_let_tv84_))
                    return iife227_(_pat_let84_0)
                def iife225_(_pat_let83_0):
                    def iife232_(d_1923_dt__update__tmp_h3_):
                        def iife233_(_pat_let87_0):
                            def iife234_(d_1924_dt__update_hcf8_h0_):
                                return D2_DC4(d_1924_dt__update_hcf8_h0_, (d_1923_dt__update__tmp_h3_).cf9, (d_1923_dt__update__tmp_h3_).cf10, (d_1923_dt__update__tmp_h3_).cf11)
                            return iife234_(_pat_let87_0)
                        return iife233_(len(pat_let_tv86_))
                    return iife232_(_pat_let83_0)
                nw305_[int(8)] = iife225_(iife226_(d_1911_v31_))
                nw305_[int(9)] = D2_DC4(p0, False, p1, p1)
                nw305_[int(10)] = d_1911_v31_
                nw305_[int(11)] = d_1911_v31_
                nw305_[int(12)] = d_1911_v31_
                nw305_[int(13)] = default__.fm32(p1, _dafny.Map({len((D5_DC13(d_1913_v33_)).cf24): p2}), globalState)
                nw305_[int(14)] = d_1911_v31_
                nw305_[int(15)] = d_1911_v31_
                nw305_[int(16)] = d_1911_v31_
                nw305_[int(17)] = d_1911_v31_
                nw305_[int(18)] = d_1911_v31_
                nw305_[int(19)] = d_1911_v31_
                nw305_[int(20)] = d_1911_v31_
                d_1914_v34_ = nw305_
                d_1925_v35_: _dafny.Map
                d_1925_v35_ = _dafny.Map({d_1914_v34_: ((d_1879_v16_)[default__.safeIndex(617, (d_1879_v16_).length(0))]) + (p2)})
                d_1925_v35_ = (d_1925_v35_).set(d_1914_v34_, p0)
            elif True:
                d_1926___mcc_h7_ = source26_.cf14
                d_1927_cf14_ = d_1926___mcc_h7_
                d_1928_v36_: D2
                d_1928_v36_ = D2_DC4(p2, (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))], p1, p1)
                d_1929_v37_: _dafny.Seq
                d_1929_v37_ = _dafny.SeqWithoutIsStrInference([d_1928_v36_, d_1928_v36_, d_1928_v36_])
                d_1930_v38_: _dafny.Seq
                d_1930_v38_ = _dafny.SeqWithoutIsStrInference([d_1881_v18_, d_1881_v18_])
                d_1931_v39_: _dafny.Map
                d_1931_v39_ = _dafny.Map({(d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]: p1})
                d_1932_v40_: _dafny.Map
                d_1932_v40_ = _dafny.Map({p2: d_1931_v39_})
                index307_ = default__.safeIndex(15, (d_1881_v18_).length(0))
                rhs273_ = p2
                rhs274_ = (0) - ((0) - (default__.safeModuloInt((default__.fm0(p0, d_1877_v14_, globalState)) * (len(d_1929_v37_)), (931) + ((0) - (p0)))))
                rhs275_ = (len(d_1887_v21_)) > (p2)
                rhs276_ = default__.safeModuloInt(p2, (p2) * (p0))
                rhs277_ = (len(((d_1930_v38_).set(default__.safeIndex(len(((d_1932_v40_)[p2] if (p2) in (d_1932_v40_) else d_1931_v39_)), len(d_1930_v38_)), d_1881_v18_)).set(default__.safeIndex(p0, len((d_1930_v38_).set(default__.safeIndex(len(((d_1932_v40_)[p2] if (p2) in (d_1932_v40_) else d_1931_v39_)), len(d_1930_v38_)), d_1881_v18_))), d_1881_v18_))) - (default__.fm0(p2, d_1877_v14_, globalState))
                lhs256_ = globalState
                lhs257_ = globalState
                lhs258_ = d_1881_v18_
                lhs259_ = default__.safeIndex(15, (d_1881_v18_).length(0))
                lhs260_ = globalState
                lhs261_ = globalState
                lhs256_.f18 = rhs273_
                lhs257_.f17 = rhs274_
                lhs258_[lhs259_] = rhs275_
                lhs260_.f18 = rhs276_
                lhs261_.f17 = rhs277_
                d_1933_v41_: _dafny.Map
                d_1933_v41_ = _dafny.Map({d_1887_v21_: False})
                d_1934_v42_: _dafny.MultiSet
                d_1934_v42_ = _dafny.MultiSet([d_1875_v12_])
                d_1935_v43_: _dafny.Seq
                d_1935_v43_ = _dafny.SeqWithoutIsStrInference([p2])
                rhs278_ = default__.safeDivisionInt((0) - (p0), (0) - (p2))
                rhs279_ = (d_1933_v41_) | (_dafny.Map({d_1887_v21_: (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]}))
                rhs280_ = default__.fm2(d_1934_v42_, (0) - (default__.safeDivisionInt(p2, len(d_1935_v43_))), globalState)
                rhs281_ = (p2) * (p2)
                lhs262_ = globalState
                lhs263_ = globalState
                lhs264_ = globalState
                lhs262_.f17 = rhs278_
                d_1933_v41_ = rhs279_
                lhs263_.f5 = rhs280_
                lhs264_.f18 = rhs281_
                (globalState).f5 = (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]
                d_1936_v44_: _dafny.Map
                d_1936_v44_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1, not((d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]), (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]])) for d_1937_i3_ in range(default__.abs(-31))]): (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]})
                d_1938_v45_: D6
                d_1938_v45_ = D6_DC16(d_1936_v44_)
                d_1936_v44_ = (d_1938_v45_).cf27
            d_1879_v16_ = d_1879_v16_
            index308_ = default__.safeIndex(15, (d_1881_v18_).length(0))
            (d_1881_v18_)[index308_] = (d_1876_v13_)[default__.safeIndex(default__.safeModuloInt(p2, p2), len(d_1876_v13_))]
            d_1939_v46_: _dafny.Array
            def lambda92_(d_1940_p2_, d_1941_p0_, d_1942_v13_, d_1943_v18_):
                def lambda93_(d_1944_i4_):
                    return D6_DC18(d_1940_p2_, d_1941_p0_, d_1942_v13_, (d_1943_v18_)[default__.safeIndex(15, (d_1943_v18_).length(0))], d_1940_p2_)

                return lambda93_

            init45_ = lambda92_(p2, p0, d_1876_v13_, d_1881_v18_)
            nw306_ = _dafny.Array(None, 17)
            for i0_45_ in range(nw306_.length(0)):
                nw306_[i0_45_] = init45_(i0_45_)
            d_1939_v46_ = nw306_
            d_1945_v47_: _dafny.Map
            d_1945_v47_ = _dafny.Map({len(d_1887_v21_): (d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]})
            d_1946_v48_: D6
            d_1946_v48_ = D6_DC18(p0, len(d_1876_v13_), d_1876_v13_, (self).fm24(p1, p1, p2, (0) - (len(d_1945_v47_)), globalState), p2)
            index309_ = default__.safeIndex(247, (d_1939_v46_).length(0))
            (d_1939_v46_)[index309_] = d_1946_v48_
            d_1947_v49_: _dafny.Map
            d_1947_v49_ = _dafny.Map({(d_1881_v18_)[default__.safeIndex(15, (d_1881_v18_).length(0))]: p1})
            d_1948_v50_: _dafny.MultiSet
            d_1948_v50_ = _dafny.MultiSet([p1, p1])
            d_1949_v51_: _dafny.Set
            d_1949_v51_ = _dafny.Set({333})
            d_1950_v52_: _dafny.Map
            d_1950_v52_ = _dafny.Map({d_1949_v51_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubacpbltn"))})
            pat_let_tv87_ = d_1947_v49_
            pat_let_tv88_ = d_1948_v50_
            index310_ = default__.safeIndex(247, (d_1939_v46_).length(0))
            def iife235_(_pat_let88_0):
                def iife236_(d_1951_dt__update__tmp_h4_):
                    def iife237_(_pat_let89_0):
                        def iife238_(d_1952_dt__update_hcf31_h0_):
                            return D6_DC18((d_1951_dt__update__tmp_h4_).cf30, d_1952_dt__update_hcf31_h0_, (d_1951_dt__update__tmp_h4_).cf32, (d_1951_dt__update__tmp_h4_).cf33, (d_1951_dt__update__tmp_h4_).cf34)
                        return iife238_(_pat_let89_0)
                    return iife237_(default__.safeDivisionInt(len(pat_let_tv87_), (pat_let_tv88_).cardinality))
                return iife236_(_pat_let88_0)
            rhs282_ = (default__.safeModuloInt(p2, (0) - (p2))) - (p2)
            rhs283_ = iife235_(d_1946_v48_)
            rhs284_ = p0
            rhs285_ = (((d_1950_v52_)[_dafny.Set({p2, p0})] if (_dafny.Set({p2, p0})) in (d_1950_v52_) else d_1887_v21_)) + (d_1887_v21_)
            lhs265_ = globalState
            lhs266_ = d_1939_v46_
            lhs267_ = default__.safeIndex(247, (d_1939_v46_).length(0))
            lhs268_ = globalState
            lhs265_.f17 = rhs282_
            lhs266_[lhs267_] = rhs283_
            lhs268_.f17 = rhs284_
            d_1887_v21_ = rhs285_
        elif True:
            d_1953_v53_: _dafny.Seq
            d_1953_v53_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            (globalState).f5 = not(not (p1) or ((d_1953_v53_) == (d_1953_v53_)))
            d_1954_v54_: _dafny.Seq
            d_1954_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bv"))
            d_1955_v55_: _dafny.MultiSet
            d_1955_v55_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1956_i5_ in range(default__.abs(760))])) + (d_1954_v54_), d_1954_v54_, d_1954_v54_])
            d_1957_v56_: _dafny.MultiSet
            d_1957_v56_ = _dafny.MultiSet([p2])
            d_1958_v57_: _dafny.Seq
            d_1958_v57_ = _dafny.SeqWithoutIsStrInference([p0, p2, (0) - (p0), p2])
            d_1959_v58_: _dafny.Set
            d_1959_v58_ = _dafny.Set({len(d_1958_v57_), p0})
            d_1960_v59_: _dafny.Map
            d_1960_v59_ = _dafny.Map({d_1959_v58_: p2})
            (globalState).f17 = ((d_1955_v55_)[(default__.fm27(p2, d_1957_v56_, len(d_1959_v58_), d_1960_v59_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkvl")))] if ((default__.fm27(p2, d_1957_v56_, len(d_1959_v58_), d_1960_v59_, globalState)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkvl")))) in (d_1955_v55_) else p2)
            d_1961_v60_: _dafny.Array
            nw307_ = _dafny.Array(False, 9)
            d_1961_v60_ = nw307_
            index311_ = default__.safeIndex(489, (d_1961_v60_).length(0))
            (d_1961_v60_)[index311_] = (p0) >= (p2)
            index312_ = default__.safeIndex(489, (d_1961_v60_).length(0))
            (d_1961_v60_)[index312_] = default__.fm2(_dafny.MultiSet([_dafny.CodePoint('k')]), p0, globalState)
            if (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]:
                d_1962_v61_: _dafny.Array
                nw308_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_1962_v61_ = nw308_
                d_1962_v61_ = d_1962_v61_
                d_1963_v62_: _dafny.MultiSet
                d_1963_v62_ = _dafny.MultiSet([False, (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p1, (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]])
                d_1964_v63_: _dafny.Seq
                d_1964_v63_ = _dafny.SeqWithoutIsStrInference([d_1963_v62_])
                d_1965_v64_: _dafny.Map
                d_1965_v64_ = _dafny.Map({(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]: d_1964_v63_})
                index313_ = default__.safeIndex(489, (d_1961_v60_).length(0))
                rhs286_ = (d_1953_v53_)[default__.safeIndex(p2, len(d_1953_v53_))]
                rhs287_ = (default__.fm33(globalState)).isdisjoint(_dafny.MultiSet(d_1958_v57_))
                rhs288_ = ((d_1963_v62_).cardinality) * (p0)
                rhs289_ = ((d_1965_v64_)[(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]] if ((d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]) in (d_1965_v64_) else (d_1964_v63_).set(default__.safeIndex(p2, len(d_1964_v63_)), d_1963_v62_))
                lhs269_ = globalState
                lhs270_ = d_1961_v60_
                lhs271_ = default__.safeIndex(489, (d_1961_v60_).length(0))
                lhs272_ = globalState
                lhs269_.f5 = rhs286_
                lhs270_[lhs271_] = rhs287_
                lhs272_.f18 = rhs288_
                d_1964_v63_ = rhs289_
                d_1966_v65_: _dafny.Map
                d_1966_v65_ = _dafny.Map({p1: p2})
                d_1967_v66_: str
                d_1967_v66_ = _dafny.CodePoint('p')
                d_1968_v67_: _dafny.Array
                nw309_ = _dafny.Array(int(0), 11)
                d_1968_v67_ = nw309_
                d_1969_v68_: D7
                d_1969_v68_ = D7_DC21((d_1958_v57_)[default__.safeIndex(default__.fm0((0) - (len(d_1966_v65_)), D0_DC1(d_1967_v66_, p2, _dafny.Set({(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]}), (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]), globalState), len(d_1958_v57_))], (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p1, d_1968_v67_, p1)
                d_1970_v69_: _dafny.Map
                d_1970_v69_ = _dafny.Map({(d_1969_v68_).cf41: len(d_1958_v57_)})
                d_1970_v69_ = (d_1970_v69_).set(d_1968_v67_, (d_1958_v57_)[default__.safeIndex(p0, len(d_1958_v57_))])
                r1 = p1
                d_1971_v70_: _dafny.Map
                d_1971_v70_ = _dafny.Map({len(d_1958_v57_): p1})
                d_1972_v71_: D2
                d_1972_v71_ = D2_DC4(p0, (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p1, p1)
                d_1973_v72_: _dafny.Map
                d_1973_v72_ = _dafny.Map({len(d_1971_v70_): D2_DC6(d_1972_v71_)})
                d_1974_v73_: D0
                d_1974_v73_ = D0_DC1(d_1967_v66_, p0, d_1869_v8_, (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p1)
                d_1975_v74_: D2
                d_1975_v74_ = D2_DC6(((d_1973_v72_)[(0) - (default__.fm0(p2, d_1974_v73_, globalState))] if ((0) - (default__.fm0(p2, d_1974_v73_, globalState))) in (d_1973_v72_) else d_1972_v71_))
                d_1976_v75_: _dafny.Map
                d_1976_v75_ = _dafny.Map({p0: d_1975_v74_})
                d_1976_v75_ = (d_1976_v75_) | ((d_1976_v75_).set(p0, d_1975_v74_))
            elif True:
                d_1869_v8_ = d_1869_v8_
                d_1977_v76_: C2
                nw310_ = C2()
                nw310_.ctor__(p0)
                d_1977_v76_ = nw310_
                d_1954_v54_ = d_1954_v54_
                d_1978_v77_: _dafny.MultiSet
                d_1978_v77_ = _dafny.MultiSet([(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]])
                d_1979_v78_: _dafny.MultiSet
                d_1979_v78_ = _dafny.MultiSet([(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]])
                d_1980_v79_: _dafny.Array
                nw311_ = _dafny.Array(None, 18)
                nw311_[int(0)] = d_1978_v77_
                nw311_[int(1)] = d_1978_v77_
                nw311_[int(2)] = d_1979_v78_
                nw311_[int(3)] = d_1978_v77_
                nw311_[int(4)] = d_1978_v77_
                nw311_[int(5)] = d_1978_v77_
                nw311_[int(6)] = d_1978_v77_
                nw311_[int(7)] = d_1978_v77_
                nw311_[int(8)] = d_1978_v77_
                nw311_[int(9)] = d_1978_v77_
                nw311_[int(10)] = d_1979_v78_
                nw311_[int(11)] = d_1978_v77_
                nw311_[int(12)] = d_1978_v77_
                nw311_[int(13)] = d_1979_v78_
                nw311_[int(14)] = d_1978_v77_
                nw311_[int(15)] = d_1978_v77_
                nw311_[int(16)] = d_1979_v78_
                nw311_[int(17)] = d_1978_v77_
                d_1980_v79_ = nw311_
                d_1981_v80_: _dafny.Map
                d_1981_v80_ = _dafny.Map({d_1980_v79_: p1})
                d_1981_v80_ = (d_1981_v80_).set(d_1980_v79_, p1)
                d_1982_v81_: str
                d_1982_v81_ = _dafny.CodePoint('d')
                d_1983_v82_: D0
                d_1983_v82_ = D0_DC1(d_1982_v81_, p0, d_1869_v8_, p1, p1)
                d_1984_v83_: _dafny.Array
                def lambda94_(d_1985_v76_):
                    def lambda95_(d_1986_i6_):
                        return (d_1986_i6_) * ((d_1985_v76_).f22)

                    return lambda95_

                init46_ = lambda94_(d_1977_v76_)
                nw312_ = _dafny.Array(None, 14)
                for i0_46_ in range(nw312_.length(0)):
                    nw312_[i0_46_] = init46_(i0_46_)
                d_1984_v83_ = nw312_
                d_1987_v84_: _dafny.Map
                d_1987_v84_ = _dafny.Map({default__.fm0(477, d_1983_v82_, globalState): d_1984_v83_})
                d_1987_v84_ = d_1987_v84_
            if (d_1958_v57_) != (d_1958_v57_):
                d_1988_v85_: _dafny.Map
                d_1988_v85_ = _dafny.Map({p0: p2})
                d_1989_v86_: C2
                nw313_ = C2()
                nw313_.ctor__(((d_1988_v85_)[p2] if (p2) in (d_1988_v85_) else p2))
                d_1989_v86_ = nw313_
                (globalState).f16 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txrd"))
                d_1990_v87_: _dafny.Array
                def lambda96_(d_1991_v86_):
                    def lambda97_(d_1992_i7_):
                        return default__.safeDivisionInt(d_1992_i7_, (d_1991_v86_).f22)

                    return lambda97_

                init47_ = lambda96_(d_1989_v86_)
                nw314_ = _dafny.Array(None, 27)
                for i0_47_ in range(nw314_.length(0)):
                    nw314_[i0_47_] = init47_(i0_47_)
                d_1990_v87_ = nw314_
                d_1993_v88_: _dafny.MultiSet
                d_1993_v88_ = _dafny.MultiSet([d_1990_v87_])
                d_1994_v89_: _dafny.Map
                d_1994_v89_ = _dafny.Map({d_1993_v88_: (self).fm24(False, (self).fm24(not((d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]), p1, p0, p2, globalState), p0, p0, globalState)})
                d_1994_v89_ = (d_1994_v89_).set(d_1993_v88_, p1)
                d_1995_v90_: bool
                d_1996_v91_: _dafny.Array
                d_1997_v92_: D2
                out16_: bool
                out17_: _dafny.Array
                out18_: D2
                out16_, out17_, out18_ = (d_1989_v86_).m2(p1, (0) - ((d_1989_v86_).f22), (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p2, globalState)
                d_1995_v90_ = out16_
                d_1996_v91_ = out17_
                d_1997_v92_ = out18_
                d_1998_v93_: _dafny.Seq
                d_1998_v93_ = _dafny.SeqWithoutIsStrInference([d_1869_v8_])
                d_1999_v94_: C2
                nw315_ = C2()
                nw315_.ctor__(len(d_1998_v93_))
                d_1999_v94_ = nw315_
            elif True:
                d_2000_v95_: _dafny.Array
                nw316_ = _dafny.Array(int(0), 12)
                d_2000_v95_ = nw316_
                d_2001_v96_: _dafny.Map
                d_2001_v96_ = _dafny.Map({(d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))]: d_2000_v95_})
                rhs290_ = p1
                rhs291_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gka"))
                rhs292_ = ((d_2001_v96_)[(p1) or ((d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))])] if ((p1) or ((d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))])) in (d_2001_v96_) else d_2000_v95_)
                lhs273_ = globalState
                lhs274_ = globalState
                lhs273_.f5 = rhs290_
                lhs274_.f15 = rhs291_
                r2 = rhs292_
                index314_ = default__.safeIndex(937, (d_2000_v95_).length(0))
                (d_2000_v95_)[index314_] = (p0 if (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))] else -875)
                d_2002_v97_: D3
                d_2002_v97_ = D3_DC8(False)
                d_2003_v98_: _dafny.Map
                d_2003_v98_ = _dafny.Map({d_2002_v97_: p1})
                d_2004_v99_: _dafny.Set
                d_2004_v99_ = _dafny.Set({d_2003_v98_, d_2003_v98_, d_2003_v98_, d_2003_v98_})
                index315_ = default__.safeIndex(937, (d_2000_v95_).length(0))
                (d_2000_v95_)[index315_] = len(d_2004_v99_)
                (self).m9((d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], p1, (d_1961_v60_)[default__.safeIndex(489, (d_1961_v60_).length(0))], len(d_1958_v57_), globalState)
                index316_ = default__.safeIndex(937, (d_2000_v95_).length(0))
                (d_2000_v95_)[index316_] = p2
                (globalState).f18 = default__.safeModuloInt(p2, (d_2000_v95_)[default__.safeIndex(937, (d_2000_v95_).length(0))])
        d_2005_v100_: _dafny.Seq
        d_2005_v100_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        d_2006_v101_: _dafny.Map
        d_2006_v101_ = _dafny.Map({p0: d_2005_v100_})
        d_2007_v102_: _dafny.Map
        d_2007_v102_ = _dafny.Map({default__.safeDivisionInt(p0, p2): d_2006_v101_})
        d_2007_v102_ = (d_2007_v102_).set(p0, _dafny.Map({(d_2005_v100_)[default__.safeIndex(642, len(d_2005_v100_))]: d_2005_v100_}))
        if p1:
            (globalState).f18 = p0
            d_2008_v103_: _dafny.Set
            d_2008_v103_ = _dafny.Set({805})
            (globalState).f6 = ((_dafny.Set({p2, p2})) - (d_2008_v103_)).ispropersubset((d_2008_v103_) - (d_2008_v103_))
            d_2009_v104_: T0
            nw317_ = C4()
            nw317_.ctor__(D0_DC0(False))
            d_2009_v104_ = nw317_
            d_2010_v105_: _dafny.Seq
            d_2010_v105_ = _dafny.SeqWithoutIsStrInference([(d_2009_v104_ if p1 else d_2009_v104_)])
            d_2010_v105_ = d_2010_v105_
            d_2011_v106_: _dafny.Array
            nw318_ = _dafny.Array(False, 29)
            d_2011_v106_ = nw318_
            d_2012_v107_: _dafny.Seq
            d_2012_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
            d_2013_v108_: _dafny.Map
            d_2013_v108_ = _dafny.Map({d_2011_v106_: d_2012_v107_})
            d_2013_v108_ = (d_2013_v108_).set(d_2011_v106_, (d_2012_v107_) + (d_2012_v107_))
            if p1:
                d_2014_v109_: _dafny.Seq
                d_2014_v109_ = _dafny.SeqWithoutIsStrInference([p1])
                d_2015_v110_: _dafny.MultiSet
                d_2015_v110_ = _dafny.MultiSet([(d_2014_v109_ if not((d_2014_v109_)[default__.safeIndex((0) - ((0) - (p0)), len(d_2014_v109_))]) else d_2014_v109_)])
                d_2016_v111_: str
                d_2016_v111_ = _dafny.CodePoint('w')
                d_2017_v112_: _dafny.Array
                def lambda98_(d_2018_v111_):
                    def lambda99_(d_2019_i8_):
                        return d_2018_v111_

                    return lambda99_

                init48_ = lambda98_(d_2016_v111_)
                nw319_ = _dafny.Array(None, 10)
                for i0_48_ in range(nw319_.length(0)):
                    nw319_[i0_48_] = init48_(i0_48_)
                d_2017_v112_ = nw319_
                d_2020_v113_: D18
                d_2020_v113_ = D18_DC49(d_2017_v112_)
                d_2021_v114_: _dafny.Set
                d_2021_v114_ = _dafny.Set({d_2020_v113_})
                rhs293_ = (d_2015_v110_) | (d_2015_v110_)
                rhs294_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dg"))
                rhs295_ = (d_2012_v107_).set(default__.safeIndex((p0) + (p0), len(d_2012_v107_)), d_2016_v111_)
                rhs296_ = (p2) - ((721) - (len(d_2021_v114_)))
                lhs275_ = globalState
                lhs276_ = globalState
                lhs277_ = globalState
                d_2015_v110_ = rhs293_
                lhs275_.f16 = rhs294_
                lhs276_.f15 = rhs295_
                lhs277_.f17 = rhs296_
                d_2022_v115_: C2
                nw320_ = C2()
                nw320_.ctor__((p2) - (p0))
                d_2022_v115_ = nw320_
                d_2023_v116_: _dafny.Map
                d_2023_v116_ = _dafny.Map({(0) - ((d_2022_v115_).f22): default__.fm65(d_2005_v100_, p1, (0) - (p0), globalState)})
                d_2024_v117_: _dafny.Map
                d_2024_v117_ = _dafny.Map({d_2012_v107_: (d_2022_v115_).f22})
                d_2023_v116_ = (d_2023_v116_).set(len(d_2024_v117_), (_dafny.MultiSet([d_2005_v100_, d_2005_v100_, d_2005_v100_])).intersection(_dafny.MultiSet([d_2005_v100_])))
                d_2025_v118_: _dafny.Array
                nw321_ = _dafny.Array(_dafny.Seq({}), 21)
                d_2025_v118_ = nw321_
                d_2026_v119_: _dafny.Map
                d_2026_v119_ = _dafny.Map({475: p1})
                d_2027_v120_: _dafny.Seq
                d_2027_v120_ = _dafny.SeqWithoutIsStrInference([d_2026_v119_, d_2026_v119_, d_2026_v119_])
                index317_ = default__.safeIndex(666, (d_2025_v118_).length(0))
                (d_2025_v118_)[index317_] = d_2027_v120_
                index318_ = default__.safeIndex(666, (d_2025_v118_).length(0))
                (d_2025_v118_)[index318_] = ((d_2027_v120_ if p1 else d_2027_v120_)) + (d_2027_v120_)
                d_2015_v110_ = ((d_2015_v110_).intersection(d_2015_v110_)).set(d_2014_v109_, default__.abs(p2))
            elif True:
                (globalState).f6 = ((d_2008_v103_).ispropersubset(_dafny.Set({368}))) and (p1)
                (globalState).f4 = d_2011_v106_
                d_2028_v121_: _dafny.Map
                d_2028_v121_ = _dafny.Map({p0: (p2) + (p0)})
                d_2028_v121_ = (d_2028_v121_).set((p2) * (p2), p0)
                d_2029_v122_: _dafny.Array
                def lambda100_(d_2030_p0_):
                    def lambda101_(d_2031_i9_):
                        return default__.safeModuloInt(d_2031_i9_, d_2030_p0_)

                    return lambda101_

                init49_ = lambda100_(p0)
                nw322_ = _dafny.Array(None, 1)
                for i0_49_ in range(nw322_.length(0)):
                    nw322_[i0_49_] = init49_(i0_49_)
                d_2029_v122_ = nw322_
                index319_ = default__.safeIndex(921, (d_2029_v122_).length(0))
                (d_2029_v122_)[index319_] = len(d_2008_v103_)
                index320_ = default__.safeIndex(406, (d_2011_v106_).length(0))
                (d_2011_v106_)[index320_] = p1
                d_2032_v123_: _dafny.MultiSet
                d_2032_v123_ = _dafny.MultiSet([p0])
                index321_ = default__.safeIndex(921, (d_2029_v122_).length(0))
                index322_ = default__.safeIndex(406, (d_2011_v106_).length(0))
                rhs297_ = p1
                rhs298_ = p2
                rhs299_ = p1
                rhs300_ = True
                rhs301_ = d_2032_v123_
                lhs278_ = d_2029_v122_
                lhs279_ = default__.safeIndex(921, (d_2029_v122_).length(0))
                lhs280_ = globalState
                lhs281_ = d_2011_v106_
                lhs282_ = default__.safeIndex(406, (d_2011_v106_).length(0))
                r1 = rhs297_
                lhs278_[lhs279_] = rhs298_
                lhs280_.f5 = rhs299_
                lhs281_[lhs282_] = rhs300_
                d_2032_v123_ = rhs301_
                index323_ = default__.safeIndex(921, (d_2029_v122_).length(0))
                (d_2029_v122_)[index323_] = len(d_1869_v8_)
        elif True:
            d_2033_v124_: _dafny.Seq
            d_2033_v124_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obe"))
            (globalState).f15 = d_2033_v124_
            (globalState).f17 = (0) - (len(_dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p1]) if p1 else _dafny.SeqWithoutIsStrInference([p1])) for d_2034_i10_ in range(default__.abs(135))])))
            d_2035_v125_: _dafny.Seq
            d_2035_v125_ = _dafny.SeqWithoutIsStrInference([p1, True])
            d_2036_v126_: _dafny.MultiSet
            d_2036_v126_ = _dafny.MultiSet([(0) - (len(d_2035_v125_))])
            (globalState).f17 = default__.safeModuloInt(((d_2036_v126_)[(0) - (len(d_2033_v124_))] if ((0) - (len(d_2033_v124_))) in (d_2036_v126_) else 498), p2)
            d_2037_v127_: str
            d_2037_v127_ = _dafny.CodePoint('m')
            rhs302_ = d_2037_v127_
            rhs303_ = not(p1)
            lhs283_ = globalState
            d_2037_v127_ = rhs302_
            lhs283_.f5 = rhs303_
            d_2038_v128_: _dafny.Map
            d_2038_v128_ = _dafny.Map({p1: p1})
            d_2039_v129_: _dafny.MultiSet
            d_2039_v129_ = _dafny.MultiSet([d_2038_v128_, d_2038_v128_])
            d_2040_v130_: _dafny.MultiSet
            d_2040_v130_ = _dafny.MultiSet([_dafny.CodePoint('q'), default__.fm1(globalState)])
            d_2041_v131_: D0
            d_2041_v131_ = D0_DC1(d_2037_v127_, p2, d_1869_v8_, p1, p1)
            d_2042_v132_: _dafny.Map
            d_2042_v132_ = _dafny.Map({d_1869_v8_: p1})
            pat_let_tv89_ = d_2037_v127_
            pat_let_tv90_ = d_2038_v128_
            pat_let_tv91_ = p1
            pat_let_tv92_ = p1
            pat_let_tv93_ = d_2038_v128_
            pat_let_tv94_ = p1
            def iife239_(_pat_let90_0):
                def iife240_(d_2045_dt__update__tmp_h8_):
                    def iife241_(_pat_let91_0):
                        def iife242_(d_2046_dt__update_hcf1_h0_):
                            def iife243_(_pat_let92_0):
                                def iife244_(d_2047_dt__update_hcf4_h0_):
                                    return D0_DC1(d_2046_dt__update_hcf1_h0_, (d_2045_dt__update__tmp_h8_).cf2, (d_2045_dt__update__tmp_h8_).cf3, d_2047_dt__update_hcf4_h0_, (d_2045_dt__update__tmp_h8_).cf5)
                                return iife244_(_pat_let92_0)
                            return iife243_(((pat_let_tv90_)[pat_let_tv91_] if (pat_let_tv92_) in (pat_let_tv93_) else pat_let_tv94_))
                        return iife242_(_pat_let91_0)
                    return iife241_(pat_let_tv89_)
                return iife240_(_pat_let90_0)
            rhs304_ = default__.safeDivisionInt((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_2037_v127_ for d_2043_i11_ in range(default__.abs(-120))])) + (_dafny.SeqWithoutIsStrInference([d_2037_v127_ for d_2044_i12_ in range(default__.abs(662))])))).cardinality, p0)
            rhs305_ = ((d_2039_v129_) | (d_2039_v129_)).issubset(default__.fm66(p0, _dafny.MultiSet([p0]), p1, p2, globalState))
            rhs306_ = (d_2033_v124_) + ((d_2033_v124_) + (d_2033_v124_))
            rhs307_ = default__.fm0(p0, iife239_(default__.fm35(d_2040_v130_, globalState)), globalState)
            rhs308_ = not(((default__.fm0(p2, d_2041_v131_, globalState) if p1 else len(_dafny.SeqWithoutIsStrInference([p0 for d_2048_i13_ in range(default__.abs(195))])))) != (len(d_2042_v132_)))
            lhs284_ = globalState
            lhs285_ = globalState
            lhs286_ = globalState
            lhs287_ = globalState
            lhs288_ = globalState
            lhs284_.f17 = rhs304_
            lhs285_.f6 = rhs305_
            lhs286_.f16 = rhs306_
            lhs287_.f18 = rhs307_
            lhs288_.f6 = rhs308_
        d_2049_v133_: T1
        nw323_ = C1()
        nw323_.ctor__(p2, (self).f20)
        d_2049_v133_ = nw323_
        r0 = d_2049_v133_
        r1 = True
        d_2050_v134_: _dafny.Array
        def lambda102_(d_2051_p2_):
            def lambda103_(d_2052_i14_):
                return (d_2052_i14_) + (d_2051_p2_)

            return lambda103_

        init50_ = lambda102_(p2)
        nw324_ = _dafny.Array(None, 7)
        for i0_50_ in range(nw324_.length(0)):
            nw324_[i0_50_] = init50_(i0_50_)
        d_2050_v134_ = nw324_
        r2 = (d_2050_v134_ if p1 else d_2050_v134_)
        return r0, r1, r2

    def m9(self, p0, p1, p2, p3, globalState):
        d_2053_v0_: _dafny.Array
        nw325_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
        d_2053_v0_ = nw325_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2053_v0_).length(0)):
            d_2054_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_2054_i0_)) and ((d_2054_i0_) < ((d_2053_v0_).length(0)))):
                (d_2053_v0_)[(d_2054_i0_)] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_2055_i1_ in range(default__.abs(-677))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfxx")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_2056_i2_ in range(default__.abs(-923))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lb"))))
        (globalState).f18 = (0) - (p3)
        d_2057_v1_: _dafny.Array
        def lambda104_(d_2058_p3_):
            def lambda105_(d_2059_i3_):
                return (d_2059_i3_) + (d_2058_p3_)

            return lambda105_

        init51_ = lambda104_(p3)
        nw326_ = _dafny.Array(None, 10)
        for i0_51_ in range(nw326_.length(0)):
            nw326_[i0_51_] = init51_(i0_51_)
        d_2057_v1_ = nw326_
        index324_ = default__.safeIndex(930, (d_2057_v1_).length(0))
        (d_2057_v1_)[index324_] = p3
        d_2060_v2_: D8
        d_2060_v2_ = D8_DC22(_dafny.Map({D7_DC21(p3, False, False, d_2057_v1_, p1): p3}))
        d_2061_v3_: D7
        d_2061_v3_ = D7_DC21(-149, p1, p1, d_2057_v1_, p2)
        d_2062_v4_: _dafny.Map
        d_2062_v4_ = _dafny.Map({d_2061_v3_: p3})
        d_2063_v5_: _dafny.Array
        nw327_ = _dafny.Array(None, 5)
        nw327_[int(0)] = d_2060_v2_
        nw327_[int(1)] = d_2060_v2_
        nw327_[int(2)] = d_2060_v2_
        nw327_[int(3)] = (d_2060_v2_ if p2 else d_2060_v2_)
        nw327_[int(4)] = D8_DC22(d_2062_v4_)
        d_2063_v5_ = nw327_
        index325_ = default__.safeIndex(760, (d_2057_v1_).length(0))
        (d_2057_v1_)[index325_] = (D13_DC37(p3, p3, p3)).cf71
        d_2064_v6_: str
        d_2064_v6_ = _dafny.CodePoint('b')
        d_2065_v7_: _dafny.Seq
        d_2065_v7_ = _dafny.SeqWithoutIsStrInference([d_2064_v6_, d_2064_v6_])
        d_2066_v8_: _dafny.Seq
        d_2066_v8_ = _dafny.SeqWithoutIsStrInference([p1, p0])
        d_2067_v9_: D10
        d_2067_v9_ = D10_DC28(d_2065_v7_, 475, p2, (d_2066_v8_)[default__.safeIndex((0) - (p3), len(d_2066_v8_))], p0)
        d_2068_v10_: D0
        d_2068_v10_ = D0_DC1(d_2064_v6_, -653, _dafny.Set({p0}), p2, p2)
        d_2069_v11_: _dafny.Map
        d_2069_v11_ = _dafny.Map({default__.fm29(globalState): len(d_2065_v7_)})
        index326_ = default__.safeIndex(930, (d_2057_v1_).length(0))
        index327_ = default__.safeIndex(760, (d_2057_v1_).length(0))
        rhs309_ = default__.safeDivisionInt(p3, p3)
        rhs310_ = (d_2063_v5_ if p2 else d_2063_v5_)
        rhs311_ = d_2065_v7_
        rhs312_ = (d_2067_v9_).cf55
        rhs313_ = (default__.safeDivisionInt(default__.fm0(528, d_2068_v10_, globalState), len(_dafny.Set({p0})))) * (((d_2069_v11_)[d_2066_v8_] if (d_2066_v8_) in (d_2069_v11_) else p3))
        lhs289_ = d_2057_v1_
        lhs290_ = default__.safeIndex(930, (d_2057_v1_).length(0))
        lhs291_ = globalState
        lhs292_ = globalState
        lhs293_ = d_2057_v1_
        lhs294_ = default__.safeIndex(760, (d_2057_v1_).length(0))
        lhs289_[lhs290_] = rhs309_
        d_2063_v5_ = rhs310_
        lhs291_.f15 = rhs311_
        lhs292_.f6 = rhs312_
        lhs293_[lhs294_] = rhs313_
        d_2070_v12_: _dafny.Array
        nw328_ = _dafny.Array(False, 20)
        d_2070_v12_ = nw328_
        index328_ = default__.safeIndex(181, (d_2070_v12_).length(0))
        (d_2070_v12_)[index328_] = p2
        d_2071_v13_: _dafny.Map
        d_2071_v13_ = _dafny.Map({p0: p0})
        d_2072_v14_: _dafny.Map
        d_2072_v14_ = _dafny.Map({p3: ((d_2071_v13_)[p2] if (p2) in (d_2071_v13_) else p0)})
        index329_ = default__.safeIndex(181, (d_2070_v12_).length(0))
        (d_2070_v12_)[index329_] = ((d_2072_v14_)[197] if (197) in (d_2072_v14_) else p0)
        if p2:
            (globalState).f18 = p3
            d_2073_v16_: _dafny.Map
            d_2073_v16_ = _dafny.Map({d_2064_v6_: True})
            d_2074_v17_: _dafny.Seq
            d_2074_v17_ = _dafny.SeqWithoutIsStrInference([d_2073_v16_])
            d_2075_v18_: _dafny.Map
            d_2075_v18_ = (d_2074_v17_)[default__.safeIndex(p3, len(d_2074_v17_))]
            d_2076_v19_: _dafny.Map
            d_2076_v19_ = _dafny.Map({_dafny.CodePoint('h'): d_2067_v9_})
            d_2077_v20_: _dafny.Set
            def iife245_():
                coll59_ = _dafny.Map()
                compr_59_: str
                for compr_59_ in ((d_2075_v18_)).keys.Elements:
                    d_2078_v15_: str = compr_59_
                    if (d_2078_v15_) in ((d_2075_v18_)):
                        coll59_[d_2078_v15_] = d_2067_v9_
                return _dafny.Map(coll59_)
            d_2077_v20_ = _dafny.Set({(iife245_()
             if p2 else d_2076_v19_)})
            d_2077_v20_ = d_2077_v20_
            d_2079_v21_: _dafny.Seq
            d_2079_v21_ = _dafny.SeqWithoutIsStrInference([p3])
            d_2080_v22_: _dafny.Set
            d_2080_v22_ = _dafny.Set({d_2079_v21_, d_2079_v21_})
            d_2081_v23_: D12
            d_2081_v23_ = D12_DC33(d_2080_v22_, p3)
            index330_ = default__.safeIndex(930, (d_2057_v1_).length(0))
            (d_2057_v1_)[index330_] = default__.fm0((d_2081_v23_).cf63, d_2068_v10_, globalState)
            d_2082_v24_: C2
            nw329_ = C2()
            nw329_.ctor__((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))])
            d_2082_v24_ = nw329_
            d_2082_v24_ = d_2082_v24_
            (globalState).f5 = p2
        elif True:
            d_2083_v25_: _dafny.Array
            def lambda106_(d_2084_p3_, d_2085_v1_):
                def lambda107_(d_2086_i4_):
                    return _dafny.Map({d_2084_p3_: (d_2085_v1_)[default__.safeIndex(930, (d_2085_v1_).length(0))]})

                return lambda107_

            init52_ = lambda106_(p3, d_2057_v1_)
            nw330_ = _dafny.Array(None, 26)
            for i0_52_ in range(nw330_.length(0)):
                nw330_[i0_52_] = init52_(i0_52_)
            d_2083_v25_ = nw330_
            d_2087_v26_: D9
            d_2087_v26_ = D9_DC25((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], p3, d_2083_v25_)
            d_2088_v27_: _dafny.MultiSet
            d_2088_v27_ = _dafny.MultiSet([(d_2087_v26_).cf47])
            d_2089_v28_: _dafny.Map
            d_2089_v28_ = _dafny.Map({(d_2088_v27_) != ((d_2088_v27_).set(p3, default__.abs(p3))): (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]})
            d_2089_v28_ = (d_2089_v28_).set((default__.fm2(_dafny.MultiSet([d_2064_v6_]), (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], globalState)) == ((d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))]), p3)
            if p0:
                d_2070_v12_ = (d_2070_v12_ if ((d_2071_v13_)[p1] if (p1) in (d_2071_v13_) else (d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))]) else d_2070_v12_)
                d_2090_v29_: _dafny.Map
                d_2090_v29_ = _dafny.Map({d_2065_v7_: d_2070_v12_})
                (globalState).f4 = ((d_2090_v29_)[d_2065_v7_] if (d_2065_v7_) in (d_2090_v29_) else d_2070_v12_)
                (globalState).f18 = (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]
                d_2091_v30_: C1
                nw331_ = C1()
                nw331_.ctor__((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], (self).f20)
                d_2091_v30_ = nw331_
                (globalState).f6 = p0
            elif True:
                (globalState).f17 = (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]
                (globalState).f5 = p2
                d_2092_v31_: T0
                nw332_ = C3()
                nw332_.ctor__((self).f20)
                d_2092_v31_ = nw332_
                d_2093_v32_: _dafny.MultiSet
                d_2093_v32_ = _dafny.MultiSet([d_2092_v31_])
                index331_ = default__.safeIndex(930, (d_2057_v1_).length(0))
                (d_2057_v1_)[index331_] = ((d_2093_v32_).cardinality) * (622)
                d_2094_v33_: _dafny.Array
                nw333_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_2094_v33_ = nw333_
                d_2095_v34_: _dafny.Array
                nw334_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                d_2095_v34_ = nw334_
                index332_ = default__.safeIndex(524, (d_2094_v33_).length(0))
                (d_2094_v33_)[index332_] = d_2095_v34_
                index333_ = default__.safeIndex(524, (d_2094_v33_).length(0))
                (d_2094_v33_)[index333_] = (d_2095_v34_ if p0 else d_2095_v34_)
                (globalState).f6 = (p2 if p2 else p1)
            d_2096_v35_: _dafny.Seq
            d_2096_v35_ = _dafny.SeqWithoutIsStrInference([(d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]])
            d_2096_v35_ = ((_dafny.SeqWithoutIsStrInference([(0) - ((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))])])) + (d_2096_v35_)).set(default__.safeIndex(((0) - (len(_dafny.Map({p3: (self).fm24(p1, p2, p3, p3, globalState)})))) + (len(d_2096_v35_)), len((_dafny.SeqWithoutIsStrInference([(0) - ((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))])])) + (d_2096_v35_))), -834)
            d_2097_v36_: _dafny.MultiSet
            d_2097_v36_ = _dafny.MultiSet([p0])
            d_2098_v37_: D18
            d_2098_v37_ = D18_DC50(True)
            d_2099_v38_: D18
            d_2099_v38_ = D18_DC51(d_2098_v37_)
            d_2100_v39_: _dafny.Seq
            d_2100_v39_ = _dafny.SeqWithoutIsStrInference([d_2098_v37_, d_2098_v37_])
            d_2101_v40_: _dafny.Seq
            d_2101_v40_ = _dafny.SeqWithoutIsStrInference([d_2099_v38_, D18_DC51(d_2098_v37_), D18_DC51((d_2100_v39_)[default__.safeIndex((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], len(d_2100_v39_))])])
            (globalState).f18 = ((d_2097_v36_)[(d_2101_v40_) < (d_2101_v40_)] if ((d_2101_v40_) < (d_2101_v40_)) in (d_2097_v36_) else default__.safeModuloInt((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]))
            source27_ = D2_DC5(p2, default__.safeDivisionInt(100, p3))
            if source27_.is_DC4:
                d_2102___mcc_h0_ = source27_.cf8
                d_2103___mcc_h1_ = source27_.cf9
                d_2104___mcc_h2_ = source27_.cf10
                d_2105___mcc_h3_ = source27_.cf11
                d_2106_cf11_ = d_2105___mcc_h3_
                d_2107_cf10_ = d_2104___mcc_h2_
                d_2108_cf9_ = d_2103___mcc_h1_
                d_2109_cf8_ = d_2102___mcc_h0_
                d_2065_v7_ = d_2065_v7_
                (globalState).f17 = d_2109_cf8_
                nw335_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
                d_2053_v0_ = nw335_
                d_2110_v41_: _dafny.Set
                d_2110_v41_ = _dafny.Set({(_dafny.MultiSet([d_2109_cf8_])).cardinality, p3})
                (globalState).f17 = (((0) - (len(_dafny.Set({(d_2066_v8_).set(default__.safeIndex((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], len(d_2066_v8_)), (d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))])})))) * (default__.fm0(len(d_2110_v41_), d_2068_v10_, globalState))) - (d_2109_cf8_)
            elif source27_.is_DC5:
                d_2111___mcc_h4_ = source27_.cf12
                d_2112___mcc_h5_ = source27_.cf13
                d_2113_cf13_ = d_2112___mcc_h5_
                d_2114_cf12_ = d_2111___mcc_h4_
                d_2115_v42_: _dafny.Map
                d_2115_v42_ = _dafny.Map({True: (d_2065_v7_) + (d_2065_v7_)})
                d_2115_v42_ = (d_2115_v42_).set((d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))], d_2065_v7_)
                d_2071_v13_ = (d_2071_v13_).set(d_2114_cf12_, p1)
                d_2116_v43_: T1
                nw336_ = C1()
                nw336_.ctor__(p3, (self).f20)
                d_2116_v43_ = nw336_
                d_2117_v44_: _dafny.Map
                d_2117_v44_ = _dafny.Map({d_2057_v1_: d_2116_v43_})
                d_2117_v44_ = (d_2117_v44_).set(d_2057_v1_, d_2116_v43_)
                (globalState).f17 = default__.fm0(default__.fm0(len(d_2065_v7_), d_2068_v10_, globalState), d_2068_v10_, globalState)
            elif source27_.is_DC3:
                d_2118___mcc_h6_ = source27_.cf7
                d_2119_cf7_ = d_2118___mcc_h6_
                (globalState).f5 = p1
                (globalState).f16 = d_2119_cf7_
                d_2120_v45_: _dafny.Set
                d_2120_v45_ = _dafny.Set({(d_2097_v36_).cardinality})
                d_2121_v46_: _dafny.Map
                d_2121_v46_ = _dafny.Map({d_2120_v45_: (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]})
                d_2119_cf7_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iemm"))) + ((default__.fm27(402, d_2088_v27_, p3, d_2121_v46_, globalState)).set(default__.safeIndex((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], len(default__.fm27(402, d_2088_v27_, p3, d_2121_v46_, globalState))), d_2064_v6_))).set(default__.safeIndex(default__.safeModuloInt((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iemm"))) + ((default__.fm27(402, d_2088_v27_, p3, d_2121_v46_, globalState)).set(default__.safeIndex((d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))], len(default__.fm27(402, d_2088_v27_, p3, d_2121_v46_, globalState))), d_2064_v6_)))), d_2064_v6_)
                d_2122_v47_: D5
                d_2122_v47_ = D5_DC14(p2)
                d_2123_v48_: D5
                d_2123_v48_ = D5_DC15(d_2122_v47_)
                d_2123_v48_ = d_2123_v48_
            elif True:
                d_2124___mcc_h7_ = source27_.cf14
                d_2125_cf14_ = d_2124___mcc_h7_
                d_2089_v28_ = d_2089_v28_
                d_2126_v49_: _dafny.Seq
                d_2126_v49_ = _dafny.SeqWithoutIsStrInference([d_2072_v14_])
                d_2127_v50_: _dafny.Map
                d_2127_v50_ = _dafny.Map({d_2126_v49_: (d_2065_v7_ if (d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))] else d_2065_v7_)})
                d_2127_v50_ = (d_2127_v50_).set(default__.fm64(d_2088_v27_, (d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))], d_2066_v8_, globalState), d_2065_v7_)
                d_2128_v51_: _dafny.Map
                d_2128_v51_ = _dafny.Map({(d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]: (d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]})
                d_2129_v52_: _dafny.Map
                d_2129_v52_ = _dafny.Map({d_2128_v51_: d_2064_v6_})
                (globalState).f18 = (default__.fm0(default__.fm0(p3, d_2068_v10_, globalState), d_2068_v10_, globalState)) - (len(d_2129_v52_))
                d_2130_v53_: C4
                nw337_ = C4()
                nw337_.ctor__((self).f20)
                d_2130_v53_ = nw337_
        d_2131_v54_: _dafny.Map
        d_2131_v54_ = _dafny.Map({(d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]: _dafny.Map({(d_2057_v1_)[default__.safeIndex(930, (d_2057_v1_).length(0))]: (d_2070_v12_)[default__.safeIndex(181, (d_2070_v12_).length(0))]})})
        d_2132_v55_: _dafny.Map
        d_2132_v55_ = _dafny.Map({p3: d_2131_v54_})
        index334_ = default__.safeIndex(181, (d_2070_v12_).length(0))
        (d_2070_v12_)[index334_] = (((d_2132_v55_)[len(d_2072_v14_)] if (len(d_2072_v14_)) in (d_2132_v55_) else d_2131_v54_)) == ((d_2131_v54_) | (d_2131_v54_))


class C7(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f26, f20):
        (self)._f26 = f26
        (self)._f20 = f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        if p0:
            d_2133_v0_: _dafny.Map
            d_2133_v0_ = _dafny.Map({p1: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxoqliqx")))})
            d_2133_v0_ = (d_2133_v0_).set((p1) * ((0) - (p1)), default__.safeDivisionInt(p1, p1))
            (globalState).f18 = (0) - (p1)
            if p0:
                d_2134_v1_: _dafny.Seq
                d_2134_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lghxhneol"))
                d_2135_v2_: D3
                d_2135_v2_ = D3_DC8(p0)
                d_2136_v3_: D3
                d_2136_v3_ = D3_DC9(d_2135_v2_)
                d_2137_v4_: D3
                d_2137_v4_ = D3_DC9(d_2136_v3_)
                pat_let_tv95_ = d_2136_v3_
                d_2138_v5_: _dafny.Map
                def iife246_(_pat_let93_0):
                    def iife247_(d_2139_dt__update__tmp_h0_):
                        def iife248_(_pat_let94_0):
                            def iife249_(d_2140_dt__update_hcf17_h0_):
                                return D3_DC9(d_2140_dt__update_hcf17_h0_)
                            return iife249_(_pat_let94_0)
                        return iife248_(pat_let_tv95_)
                    return iife247_(_pat_let93_0)
                d_2138_v5_ = _dafny.Map({(0) - (len((d_2134_v1_ if p0 else d_2134_v1_))): iife246_(d_2137_v4_)})
                d_2138_v5_ = (d_2138_v5_).set(p1, d_2137_v4_)
                d_2141_v6_: _dafny.MultiSet
                d_2141_v6_ = _dafny.MultiSet([p0])
                d_2142_v7_: _dafny.MultiSet
                d_2142_v7_ = d_2141_v6_
                d_2143_v8_: _dafny.Array
                def lambda108_(d_2144_p1_):
                    def lambda109_(d_2145_i0_):
                        return _dafny.SeqWithoutIsStrInference([d_2144_p1_, d_2144_p1_])

                    return lambda109_

                init53_ = lambda108_(p1)
                nw338_ = _dafny.Array(None, 9)
                for i0_53_ in range(nw338_.length(0)):
                    nw338_[i0_53_] = init53_(i0_53_)
                d_2143_v8_ = nw338_
                d_2146_v9_: D3
                d_2146_v9_ = D3_DC7(d_2143_v8_)
                d_2147_v10_: T0
                nw339_ = C5()
                nw339_.ctor__(d_2142_v7_, d_2146_v9_, (self).f20)
                d_2147_v10_ = nw339_
                d_2148_v11_: T1
                nw340_ = C1()
                nw340_.ctor__(p1, (d_2147_v10_).f20)
                d_2148_v11_ = nw340_
                d_2149_v12_: _dafny.Map
                d_2149_v12_ = _dafny.Map({d_2148_v11_: (d_2148_v11_).f21})
                d_2150_v13_: _dafny.Map
                d_2150_v13_ = _dafny.Map({d_2147_v10_: default__.fm46(len(d_2149_v12_), False, p0, globalState)})
                d_2151_v14_: _dafny.Set
                d_2151_v14_ = _dafny.Set({(self).f20, (d_2147_v10_).f20})
                d_2152_v15_: _dafny.MultiSet
                d_2152_v15_ = _dafny.MultiSet([p1, p1, len(d_2151_v14_), p1])
                d_2153_v16_: _dafny.Array
                nw341_ = _dafny.Array(False, 25)
                d_2153_v16_ = nw341_
                d_2154_v17_: _dafny.Set
                d_2154_v17_ = _dafny.Set({False})
                d_2155_v18_: D0
                d_2155_v18_ = D0_DC1(default__.fm1(globalState), (d_2148_v11_).f21, d_2154_v17_, p0, (self).f26)
                d_2156_v19_: _dafny.Map
                d_2156_v19_ = _dafny.Map({d_2153_v16_: default__.fm0((d_2148_v11_).f21, d_2155_v18_, globalState)})
                d_2157_v20_: _dafny.Map
                d_2157_v20_ = _dafny.Map({p0: (self).f26})
                d_2158_v21_: str
                d_2158_v21_ = _dafny.CodePoint('g')
                d_2159_v22_: _dafny.Array
                nw342_ = _dafny.Array(None, 13)
                nw342_[int(0)] = len(d_2150_v13_)
                nw342_[int(1)] = (d_2148_v11_).f21
                nw342_[int(2)] = p1
                nw342_[int(3)] = (d_2148_v11_).f21
                nw342_[int(4)] = p1
                nw342_[int(5)] = default__.safeModuloInt((d_2141_v6_).cardinality, p1)
                nw342_[int(6)] = ((d_2152_v15_)[(d_2148_v11_).f21] if ((d_2148_v11_).f21) in (d_2152_v15_) else p1)
                nw342_[int(7)] = default__.safeDivisionInt(len(d_2156_v19_), (d_2148_v11_).f21)
                nw342_[int(8)] = (0) - (default__.safeDivisionInt((0) - ((default__.fm67(len(d_2157_v20_), d_2158_v21_, (self).f26, not(p0), globalState)).cf63), (d_2148_v11_).f21))
                nw342_[int(9)] = (d_2148_v11_).f21
                nw342_[int(10)] = (d_2148_v11_).f21
                nw342_[int(11)] = (p1) + (740)
                nw342_[int(12)] = (d_2148_v11_).f21
                d_2159_v22_ = nw342_
                index335_ = default__.safeIndex(693, (d_2159_v22_).length(0))
                (d_2159_v22_)[index335_] = ((d_2148_v11_).f21) + ((d_2148_v11_).f21)
                d_2160_v23_: _dafny.MultiSet
                d_2160_v23_ = _dafny.MultiSet([(d_2152_v15_).set(-200, default__.abs((d_2148_v11_).f21)), d_2152_v15_, d_2152_v15_])
                d_2161_v24_: _dafny.Map
                d_2161_v24_ = _dafny.Map({(self).f26: _dafny.MultiSet([(d_2152_v15_).cardinality, len(d_2134_v1_)])})
                index336_ = default__.safeIndex(693, (d_2159_v22_).length(0))
                (d_2159_v22_)[index336_] = ((d_2160_v23_)[d_2152_v15_] if (d_2152_v15_) in (d_2160_v23_) else (p1) - ((0) - (len(d_2161_v24_))))
                d_2157_v20_ = d_2157_v20_
                (globalState).f17 = (d_2148_v11_).f21
                d_2162_v25_: _dafny.Set
                d_2162_v25_ = _dafny.Set({355})
                d_2163_v26_: _dafny.Map
                d_2163_v26_ = _dafny.Map({p0: d_2162_v25_})
                index337_ = default__.safeIndex(906, (d_2153_v16_).length(0))
                (d_2153_v16_)[index337_] = not((self).f26)
                index338_ = default__.safeIndex(906, (d_2153_v16_).length(0))
                rhs314_ = _dafny.Map({p0: _dafny.Set({(d_2148_v11_).f21, (0) - ((d_2159_v22_)[default__.safeIndex(693, (d_2159_v22_).length(0))]), (d_2148_v11_).f21})})
                rhs315_ = ((d_2159_v22_)[default__.safeIndex(693, (d_2159_v22_).length(0))]) + ((d_2159_v22_)[default__.safeIndex(693, (d_2159_v22_).length(0))])
                rhs316_ = True
                lhs295_ = globalState
                lhs296_ = d_2153_v16_
                lhs297_ = default__.safeIndex(906, (d_2153_v16_).length(0))
                d_2163_v26_ = rhs314_
                lhs295_.f18 = rhs315_
                lhs296_[lhs297_] = rhs316_
            elif True:
                d_2164_v27_: _dafny.Map
                d_2164_v27_ = _dafny.Map({(self).f26: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2165_i1_ in range(default__.abs(512))])})
                d_2166_v28_: _dafny.Seq
                d_2166_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ne"))
                d_2167_v29_: _dafny.Map
                d_2167_v29_ = _dafny.Map({(d_2164_v27_) | ((_dafny.Map({(self).f26: d_2166_v28_})).set(p0, d_2166_v28_)): p1})
                d_2167_v29_ = (d_2167_v29_).set((d_2164_v27_) | (d_2164_v27_), p1)
                d_2168_v30_: C6
                nw343_ = C6()
                nw343_.ctor__((self).f20)
                d_2168_v30_ = nw343_
                (globalState).f17 = p1
                d_2169_v31_: D4
                d_2169_v31_ = D4_DC11((self).f26, False, p1, p1)
                d_2170_v32_: _dafny.Seq
                d_2170_v32_ = _dafny.SeqWithoutIsStrInference([(d_2169_v31_).cf21, p1, (0) - (p1), 995])
                r1 = default__.safeModuloInt(p1, (d_2170_v32_)[default__.safeIndex(p1, len(d_2170_v32_))])
                d_2170_v32_ = default__.fm47(p1, (self).f26, (self).f26, globalState)
            (globalState).f6 = (p1) == (((0) - (p1)) * (p1))
            d_2171_v33_: _dafny.MultiSet
            d_2171_v33_ = _dafny.MultiSet([(self).f26])
            d_2172_v34_: _dafny.Seq
            d_2172_v34_ = _dafny.SeqWithoutIsStrInference([False, (self).f26])
            if (_dafny.MultiSet((d_2172_v34_ if (self).f26 else d_2172_v34_))).issubset(d_2171_v33_):
                d_2173_v35_: _dafny.Array
                nw344_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_2173_v35_ = nw344_
                d_2174_v36_: _dafny.MultiSet
                d_2174_v36_ = _dafny.MultiSet([d_2173_v35_])
                (globalState).f18 = default__.safeModuloInt((0) - (p1), ((d_2174_v36_)[d_2173_v35_] if (d_2173_v35_) in (d_2174_v36_) else p1))
                d_2175_v37_: str
                d_2175_v37_ = _dafny.CodePoint('t')
                d_2176_v38_: _dafny.MultiSet
                d_2176_v38_ = _dafny.MultiSet([d_2175_v37_])
                d_2177_v39_: _dafny.Set
                d_2177_v39_ = _dafny.Set({False})
                d_2178_v40_: _dafny.Seq
                d_2178_v40_ = _dafny.SeqWithoutIsStrInference([p1, default__.fm0(p1, D0_DC1(d_2175_v37_, p1, d_2177_v39_, (self).f26, p0), globalState), p1, 317])
                d_2179_v41_: _dafny.Map
                d_2179_v41_ = _dafny.Map({(self).f26: p0})
                (globalState).f6 = default__.fm2((d_2176_v38_) | ((d_2176_v38_).set(d_2175_v37_, default__.abs(len(d_2178_v40_)))), len(d_2179_v41_), globalState)
                d_2180_v42_: _dafny.MultiSet
                d_2180_v42_ = _dafny.MultiSet([p1])
                d_2181_v43_: _dafny.Array
                nw345_ = _dafny.Array(None, 7)
                nw345_[int(0)] = (d_2180_v42_).issubset(_dafny.MultiSet([507]))
                nw345_[int(1)] = (self).f26
                nw345_[int(2)] = p0
                nw345_[int(3)] = (d_2177_v39_).isdisjoint(d_2177_v39_)
                nw345_[int(4)] = (p0 if (self).f26 else p0)
                nw345_[int(5)] = (self).f26
                nw345_[int(6)] = (self).f26
                d_2181_v43_ = nw345_
                index339_ = default__.safeIndex(880, (d_2181_v43_).length(0))
                (d_2181_v43_)[index339_] = (d_2178_v40_) <= (d_2178_v40_)
                index340_ = default__.safeIndex(880, (d_2181_v43_).length(0))
                rhs317_ = (default__.safeDivisionInt(p1, p1)) * (p1)
                rhs318_ = not((self).f26)
                lhs298_ = globalState
                lhs299_ = d_2181_v43_
                lhs300_ = default__.safeIndex(880, (d_2181_v43_).length(0))
                lhs298_.f17 = rhs317_
                lhs299_[lhs300_] = rhs318_
                d_2182_v44_: _dafny.Map
                d_2182_v44_ = _dafny.Map({d_2181_v43_: d_2180_v42_})
                (globalState).f5 = not((d_2181_v43_) in (d_2182_v44_))
                d_2183_v45_: _dafny.Set
                d_2184_v46_: C0
                out19_: _dafny.Set
                out20_: C0
                out19_, out20_ = (self).m7(_dafny.CodePoint('g'), globalState)
                d_2183_v45_ = out19_
                d_2184_v46_ = out20_
            elif True:
                d_2185_v47_: str
                d_2185_v47_ = _dafny.CodePoint('a')
                d_2186_v48_: D0
                d_2186_v48_ = D0_DC1(d_2185_v47_, p1, _dafny.Set({(self).f26}), (self).f26, not((self).f26))
                r1 = default__.fm0(p1, d_2186_v48_, globalState)
                (globalState).f6 = p0
                d_2187_v49_: _dafny.Seq
                d_2187_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lateempw"))
                d_2188_v50_: _dafny.Set
                d_2188_v50_ = _dafny.Set({424, 28, p1, len(d_2187_v49_), p1})
                d_2189_v51_: _dafny.Seq
                d_2189_v51_ = _dafny.SeqWithoutIsStrInference([d_2188_v50_])
                d_2190_v52_: _dafny.Seq
                d_2190_v52_ = _dafny.SeqWithoutIsStrInference([-428])
                def iife250_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in (d_2190_v52_).Elements:
                        d_2191_v53_: int = compr_60_
                        if (d_2191_v53_) in (d_2190_v52_):
                            coll60_ = coll60_.union(_dafny.Set([(d_2191_v53_) * (-476)]))
                    return _dafny.Set(coll60_)
                rhs319_ = (((d_2189_v51_)[default__.safeIndex(p1, len(d_2189_v51_))]) | (d_2188_v50_)) | (d_2188_v50_)
                rhs320_ = (0) - (default__.safeModuloInt(len(iife250_()
                ), ((d_2171_v33_).intersection(d_2171_v33_)).cardinality))
                lhs301_ = globalState
                d_2188_v50_ = rhs319_
                lhs301_.f18 = rhs320_
                d_2192_v54_: D10
                d_2192_v54_ = D10_DC26(d_2190_v52_)
                rhs321_ = -781
                rhs322_ = ((d_2192_v54_).cf49).set(default__.safeIndex(((d_2171_v33_).cardinality) + (p1), len((d_2192_v54_).cf49)), p1)
                r1 = rhs321_
                d_2190_v52_ = rhs322_
                d_2193_v55_: _dafny.Array
                nw346_ = _dafny.Array(None, 16)
                nw346_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xysq"))) + (d_2187_v49_)
                nw346_[int(1)] = (d_2187_v49_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcrcl")))
                nw346_[int(2)] = d_2187_v49_
                nw346_[int(3)] = d_2187_v49_
                nw346_[int(4)] = d_2187_v49_
                nw346_[int(5)] = (d_2187_v49_) + (d_2187_v49_)
                nw346_[int(6)] = (d_2187_v49_) + (d_2187_v49_)
                nw346_[int(7)] = d_2187_v49_
                nw346_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uej"))) + (d_2187_v49_)
                nw346_[int(9)] = (d_2187_v49_).set(default__.safeIndex(len(d_2187_v49_), len(d_2187_v49_)), d_2185_v47_)
                nw346_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (_dafny.SeqWithoutIsStrInference([(d_2187_v49_)[default__.safeIndex(p1, len(d_2187_v49_))] for d_2194_i2_ in range(default__.abs(793))]))
                nw346_[int(11)] = d_2187_v49_
                nw346_[int(12)] = d_2187_v49_
                nw346_[int(13)] = default__.fm34(p0, globalState)
                nw346_[int(14)] = (d_2187_v49_) + (d_2187_v49_)
                nw346_[int(15)] = d_2187_v49_
                d_2193_v55_ = nw346_
                d_2195_v56_: _dafny.Seq
                d_2195_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gnpba")), d_2187_v49_, (d_2187_v49_).set(default__.safeIndex(p1, len(d_2187_v49_)), d_2185_v47_), d_2187_v49_, default__.fm34((self).f26, globalState)])
                index341_ = default__.safeIndex(622, (d_2193_v55_).length(0))
                (d_2193_v55_)[index341_] = (d_2195_v56_)[default__.safeIndex(len(d_2172_v34_), len(d_2195_v56_))]
                index342_ = default__.safeIndex(622, (d_2193_v55_).length(0))
                (d_2193_v55_)[index342_] = ((_dafny.SeqWithoutIsStrInference([d_2185_v47_ for d_2196_i3_ in range(default__.abs(407))])) + (d_2187_v49_)) + (d_2187_v49_)
        elif True:
            d_2197_v57_: _dafny.Array
            def lambda110_(d_2198_i4_):
                return (d_2198_i4_) - (305)

            init54_ = lambda110_
            nw347_ = _dafny.Array(None, 22)
            for i0_54_ in range(nw347_.length(0)):
                nw347_[i0_54_] = init54_(i0_54_)
            d_2197_v57_ = nw347_
            d_2199_v58_: _dafny.Seq
            d_2199_v58_ = _dafny.SeqWithoutIsStrInference([(self).f26])
            d_2200_v59_: _dafny.Map
            d_2200_v59_ = _dafny.Map({(d_2199_v58_)[default__.safeIndex(p1, len(d_2199_v58_))]: d_2197_v57_})
            d_2201_v60_: _dafny.Seq
            d_2201_v60_ = _dafny.SeqWithoutIsStrInference([d_2197_v57_])
            d_2202_v61_: _dafny.Map
            d_2202_v61_ = _dafny.Map({p1: d_2197_v57_})
            d_2203_v62_: _dafny.Array
            nw348_ = _dafny.Array(None, 11)
            nw348_[int(0)] = d_2197_v57_
            nw348_[int(1)] = d_2197_v57_
            nw348_[int(2)] = d_2197_v57_
            nw348_[int(3)] = ((d_2200_v59_)[p0] if (p0) in (d_2200_v59_) else d_2197_v57_)
            nw348_[int(4)] = d_2197_v57_
            nw348_[int(5)] = d_2197_v57_
            nw348_[int(6)] = d_2197_v57_
            nw348_[int(7)] = d_2197_v57_
            nw348_[int(8)] = (d_2201_v60_)[default__.safeIndex(p1, len(d_2201_v60_))]
            nw348_[int(9)] = ((d_2202_v61_)[len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2204_i5_ in range(default__.abs(-408))]))] if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2205_i5_ in range(default__.abs(-408))]))) in (d_2202_v61_) else d_2197_v57_)
            nw348_[int(10)] = d_2197_v57_
            d_2203_v62_ = nw348_
            index343_ = default__.safeIndex(635, (d_2203_v62_).length(0))
            (d_2203_v62_)[index343_] = d_2197_v57_
            index344_ = default__.safeIndex(635, (d_2203_v62_).length(0))
            (d_2203_v62_)[index344_] = d_2197_v57_
            d_2206_v63_: _dafny.Seq
            d_2206_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kajyo"))
            (globalState).f6 = (d_2206_v63_) != (d_2206_v63_)
            d_2207_v64_: str
            d_2207_v64_ = _dafny.CodePoint('j')
            d_2208_v65_: D8
            d_2208_v65_ = D8_DC23((d_2206_v63_).set(default__.safeIndex(p1, len(d_2206_v63_)), d_2207_v64_))
            rhs323_ = p1
            rhs324_ = d_2208_v65_
            lhs302_ = globalState
            lhs302_.f18 = rhs323_
            d_2208_v65_ = rhs324_
            (globalState).f5 = False
            d_2209_v66_: D17
            d_2209_v66_ = D17_DC48(not((d_2199_v58_)[default__.safeIndex((_dafny.MultiSet([len(default__.fm59(p1, globalState)), p1])).cardinality, len(d_2199_v58_))]), (p0) == ((self).f26))
            d_2209_v66_ = d_2209_v66_
        d_2210_v67_: _dafny.Map
        d_2210_v67_ = _dafny.Map({(self).f26: (self).f26})
        d_2211_v68_: str
        d_2211_v68_ = _dafny.CodePoint('x')
        d_2212_v69_: _dafny.Seq
        d_2212_v69_ = _dafny.SeqWithoutIsStrInference([d_2211_v68_, _dafny.CodePoint('q'), d_2211_v68_])
        d_2213_v70_: _dafny.MultiSet
        d_2213_v70_ = _dafny.MultiSet([d_2211_v68_, (d_2212_v69_)[default__.safeIndex(p1, len(d_2212_v69_))], d_2211_v68_])
        d_2210_v67_ = (d_2210_v67_).set((self).f26, (default__.fm2(d_2213_v70_, p1, globalState) if p0 else p0))
        rhs325_ = p1
        rhs326_ = (0) - (p1)
        lhs303_ = globalState
        lhs304_ = globalState
        lhs303_.f18 = rhs325_
        lhs304_.f18 = rhs326_
        d_2214_v71_: _dafny.Array
        def lambda111_(d_2215_p1_):
            def lambda112_(d_2216_i6_):
                return default__.safeModuloInt(d_2216_i6_, d_2215_p1_)

            return lambda112_

        init55_ = lambda111_(p1)
        nw349_ = _dafny.Array(None, 16)
        for i0_55_ in range(nw349_.length(0)):
            nw349_[i0_55_] = init55_(i0_55_)
        d_2214_v71_ = nw349_
        d_2214_v71_ = d_2214_v71_
        d_2217_i7_: int
        d_2217_i7_ = 0
        with _dafny.label("9"):
            while (d_2212_v69_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))):
                with _dafny.c_label("9"):
                    if (d_2217_i7_) >= (100):
                        raise _dafny.Break("9")
                    d_2217_i7_ = (d_2217_i7_) + (1)
                    (globalState).f5 = (self).f26
                    d_2218_v72_: _dafny.Set
                    d_2218_v72_ = _dafny.Set({p0, p0})
                    d_2219_v73_: D0
                    d_2219_v73_ = D0_DC1(default__.fm1(globalState), p1, d_2218_v72_, (self).f26, (self).f26)
                    r1 = default__.fm0((p1) + (p1), d_2219_v73_, globalState)
                    if p0:
                        (globalState).f6 = (p1) != (p1)
                        (globalState).f18 = 245
                        (globalState).f6 = (self).f26
                        d_2220_v74_: _dafny.MultiSet
                        d_2220_v74_ = _dafny.MultiSet([p1])
                        d_2221_v75_: _dafny.Set
                        d_2221_v75_ = _dafny.Set({_dafny.CodePoint('s'), default__.fm1(globalState)})
                        rhs327_ = ((d_2210_v67_) | (d_2210_v67_)).set(not ((self).f26) or ((self).f26), (d_2220_v74_).ispropersubset((d_2220_v74_).set(p1, default__.abs(p1))))
                        rhs328_ = not ((self).f26) or ((d_2221_v75_).issubset(d_2221_v75_))
                        lhs305_ = globalState
                        d_2210_v67_ = rhs327_
                        lhs305_.f6 = rhs328_
                        index345_ = default__.safeIndex(486, (d_2214_v71_).length(0))
                        (d_2214_v71_)[index345_] = 357
                        d_2222_v76_: _dafny.Map
                        d_2222_v76_ = _dafny.Map({p1: p1})
                        index346_ = default__.safeIndex(486, (d_2214_v71_).length(0))
                        (d_2214_v71_)[index346_] = ((d_2222_v76_)[(p1) + (p1)] if ((p1) + (p1)) in (d_2222_v76_) else len(d_2212_v69_))
                    elif True:
                        d_2223_v77_: _dafny.Seq
                        d_2223_v77_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])
                        rhs329_ = d_2212_v69_
                        rhs330_ = d_2223_v77_
                        rhs331_ = (self).f26
                        lhs306_ = globalState
                        lhs307_ = globalState
                        lhs306_.f16 = rhs329_
                        d_2223_v77_ = rhs330_
                        lhs307_.f5 = rhs331_
                        d_2224_v78_: _dafny.Set
                        d_2225_v79_: C0
                        out21_: _dafny.Set
                        out22_: C0
                        out21_, out22_ = (self).m7(d_2211_v68_, globalState)
                        d_2224_v78_ = out21_
                        d_2225_v79_ = out22_
                        d_2226_v80_: D4
                        d_2226_v80_ = D4_DC10(_dafny.SeqWithoutIsStrInference([d_2212_v69_]))
                        d_2227_v81_: D4
                        d_2227_v81_ = D4_DC12(d_2226_v80_)
                        d_2228_v82_: _dafny.Seq
                        d_2228_v82_ = _dafny.SeqWithoutIsStrInference([d_2212_v69_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm"))])
                        d_2229_v83_: _dafny.Array
                        nw350_ = _dafny.Array(None, 24)
                        nw350_[int(0)] = d_2212_v69_
                        nw350_[int(1)] = d_2212_v69_
                        nw350_[int(2)] = d_2212_v69_
                        nw350_[int(3)] = (d_2212_v69_) + (d_2212_v69_)
                        nw350_[int(4)] = d_2212_v69_
                        nw350_[int(5)] = d_2212_v69_
                        nw350_[int(6)] = (d_2212_v69_) + (d_2212_v69_)
                        nw350_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idylssf"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbvimle")))
                        nw350_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2230_i8_ in range(default__.abs(733))])
                        nw350_[int(9)] = d_2212_v69_
                        nw350_[int(10)] = ((d_2228_v82_)[default__.safeIndex((0) - (p1), len(d_2228_v82_))]) + (d_2212_v69_)
                        nw350_[int(11)] = d_2212_v69_
                        nw350_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pe"))) + (d_2212_v69_)
                        nw350_[int(13)] = (d_2212_v69_ if True else d_2212_v69_)
                        nw350_[int(14)] = d_2212_v69_
                        nw350_[int(15)] = d_2212_v69_
                        nw350_[int(16)] = ((d_2212_v69_).set(default__.safeIndex(p1, len(d_2212_v69_)), d_2211_v68_)).set(default__.safeIndex(p1, len((d_2212_v69_).set(default__.safeIndex(p1, len(d_2212_v69_)), d_2211_v68_))), _dafny.CodePoint('a'))
                        nw350_[int(17)] = d_2212_v69_
                        nw350_[int(18)] = d_2212_v69_
                        nw350_[int(19)] = (d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opvcfjgy")))
                        nw350_[int(20)] = (d_2212_v69_) + (d_2212_v69_)
                        nw350_[int(21)] = default__.fm34((self).f26, globalState)
                        nw350_[int(22)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qvhgvptst"))) + (d_2212_v69_)
                        nw350_[int(23)] = d_2212_v69_
                        d_2229_v83_ = nw350_
                        index347_ = default__.safeIndex(811, (d_2229_v83_).length(0))
                        (d_2229_v83_)[index347_] = default__.fm34(p0, globalState)
                        index348_ = default__.safeIndex(811, (d_2229_v83_).length(0))
                        rhs332_ = (((d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmjpw")))).set(default__.safeIndex(p1, len((d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmjpw"))))), _dafny.CodePoint('c'))) + ((d_2212_v69_) + (d_2212_v69_))
                        rhs333_ = d_2227_v81_
                        rhs334_ = default__.fm34(p0, globalState)
                        rhs335_ = d_2212_v69_
                        rhs336_ = (p0) or (not (p0) or ((self).f26))
                        lhs308_ = globalState
                        lhs309_ = d_2229_v83_
                        lhs310_ = default__.safeIndex(811, (d_2229_v83_).length(0))
                        lhs311_ = globalState
                        d_2212_v69_ = rhs332_
                        d_2227_v81_ = rhs333_
                        lhs308_.f16 = rhs334_
                        lhs309_[lhs310_] = rhs335_
                        lhs311_.f6 = rhs336_
                        arr0_ = d_2225_v79_.f23
                        index349_ = default__.safeIndex(23, (d_2225_v79_.f23).length(0))
                        arr0_[index349_] = (self).f26
                        arr1_ = d_2225_v79_.f23
                        index350_ = default__.safeIndex(23, (d_2225_v79_.f23).length(0))
                        arr1_[index350_] = ((((d_2213_v70_)[d_2211_v68_] if (d_2211_v68_) in (d_2213_v70_) else p1)) >= (p1)) or ((self).f26)
                        d_2231_v84_: _dafny.Array
                        nw351_ = _dafny.Array(_dafny.Array(None, 0), 7)
                        d_2231_v84_ = nw351_
                        index351_ = default__.safeIndex(922, (d_2231_v84_).length(0))
                        (d_2231_v84_)[index351_] = d_2214_v71_
                        d_2232_v85_: _dafny.MultiSet
                        d_2232_v85_ = _dafny.MultiSet([(_dafny.MultiSet(d_2223_v77_)).cardinality, p1])
                        d_2233_v86_: D17
                        d_2233_v86_ = D17_DC47(p1)
                        d_2234_v87_: _dafny.Set
                        d_2234_v87_ = _dafny.Set({p1})
                        index352_ = default__.safeIndex(922, (d_2231_v84_).length(0))
                        nw352_ = _dafny.Array(None, 13)
                        nw352_[int(0)] = p1
                        nw352_[int(1)] = (p1) + (default__.fm0(p1, D0_DC1(d_2211_v68_, p1, d_2218_v72_, p0, p0), globalState))
                        nw352_[int(2)] = (0) - (default__.safeModuloInt(p1, p1))
                        nw352_[int(3)] = p1
                        nw352_[int(4)] = (d_2232_v85_).cardinality
                        nw352_[int(5)] = p1
                        nw352_[int(6)] = p1
                        nw352_[int(7)] = default__.safeDivisionInt(p1, p1)
                        nw352_[int(8)] = p1
                        nw352_[int(9)] = (_dafny.MultiSet([p0, (self).f26])).cardinality
                        nw352_[int(10)] = p1
                        nw352_[int(11)] = (d_2233_v86_).cf92
                        nw352_[int(12)] = (671 if (self).f26 else ((d_2232_v85_)[p1] if (p1) in (d_2232_v85_) else len(d_2234_v87_)))
                        (d_2231_v84_)[index352_] = nw352_
                    (globalState).f18 = (len((d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hewnmccek")))) if p0 else default__.safeDivisionInt(p1, p1))
                    pass
            pass
        d_2235_i9_: int
        d_2235_i9_ = 0
        with _dafny.label("10"):
            while (self).f26:
                with _dafny.c_label("10"):
                    if (d_2235_i9_) >= (100):
                        raise _dafny.Break("10")
                    d_2235_i9_ = (d_2235_i9_) + (1)
                    d_2236_v88_: _dafny.Set
                    d_2236_v88_ = _dafny.Set({(d_2212_v69_) + (d_2212_v69_), (d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlfpup"))), (d_2212_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjmbtkash"))), d_2212_v69_})
                    d_2236_v88_ = d_2236_v88_
                    (globalState).f18 = default__.safeModuloInt(p1, (995) * (len(d_2212_v69_)))
                    d_2214_v71_ = d_2214_v71_
                    d_2237_v89_: D2
                    d_2237_v89_ = D2_DC3(d_2212_v69_)
                    d_2237_v89_ = d_2237_v89_
                    pass
            pass
        d_2238_v90_: _dafny.MultiSet
        d_2238_v90_ = _dafny.MultiSet([p0])
        d_2239_v91_: _dafny.MultiSet
        d_2239_v91_ = d_2238_v90_
        r0 = (d_2238_v90_).set(not ((self).f26) or ((self).f26), default__.abs(((d_2239_v91_)).cardinality))
        r1 = p1
        d_2240_v92_: D0
        d_2240_v92_ = D0_DC1(d_2211_v68_, p1, _dafny.Set({(self).f26, (self).f26}), p0, (self).f26)
        d_2241_v93_: _dafny.Set
        d_2241_v93_ = _dafny.Set({p0})
        d_2242_v94_: _dafny.MultiSet
        d_2242_v94_ = _dafny.MultiSet([default__.fm0(p1, d_2240_v92_, globalState), default__.fm0(p1, D0_DC1(d_2211_v68_, p1, d_2241_v93_, p0, (self).f26), globalState)])
        d_2243_v95_: _dafny.Map
        d_2243_v95_ = _dafny.Map({(d_2242_v94_) - (d_2242_v94_): _dafny.CodePoint('k')})
        r2 = ((d_2243_v95_)[_dafny.MultiSet([p1, p1])] if (_dafny.MultiSet([p1, p1])) in (d_2243_v95_) else d_2211_v68_)
        return r0, r1, r2

    def m6(self, globalState):
        r0: int = int(0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_2244_v1_: _dafny.Array
        def lambda113_(d_2245_i0_):
            def iife251_():
                coll61_ = _dafny.Set()
                compr_61_: int
                for compr_61_ in _dafny.IntegerRange(518, 662):
                    d_2246_v0_: int = compr_61_
                    if ((518) <= (d_2246_v0_)) and ((d_2246_v0_) < (662)):
                        coll61_ = coll61_.union(_dafny.Set([(d_2246_v0_) * (-439)]))
                return _dafny.Set(coll61_)
            return (_dafny.MultiSet([_dafny.Map({(0) - (-49): -44}), _dafny.Map({924: 86}), _dafny.Map({16: len(iife251_()
            )})])) | (_dafny.MultiSet([_dafny.Map({85: 242})]))

        init56_ = lambda113_
        nw353_ = _dafny.Array(None, 1)
        for i0_56_ in range(nw353_.length(0)):
            nw353_[i0_56_] = init56_(i0_56_)
        d_2244_v1_ = nw353_
        d_2247_v2_: int
        d_2247_v2_ = 938
        d_2248_v3_: _dafny.Map
        d_2248_v3_ = _dafny.Map({d_2247_v2_: d_2247_v2_})
        d_2249_v4_: _dafny.MultiSet
        d_2249_v4_ = _dafny.MultiSet([d_2248_v3_])
        index353_ = default__.safeIndex(704, (d_2244_v1_).length(0))
        (d_2244_v1_)[index353_] = d_2249_v4_
        index354_ = default__.safeIndex(704, (d_2244_v1_).length(0))
        (d_2244_v1_)[index354_] = d_2249_v4_
        d_2250_v5_: D8
        d_2250_v5_ = D8_DC23(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgaae")))
        d_2251_i1_: int
        d_2251_i1_ = 0
        with _dafny.label("11"):
            def lambda116_(source28_):
                if source28_.is_DC23:
                    d_2285___mcc_h0_ = source28_.cf44
                    d_2286_cf44_ = d_2285___mcc_h0_
                    return (self).f26
                elif True:
                    d_2287___mcc_h1_ = source28_.cf43
                    d_2288_cf43_ = d_2287___mcc_h1_
                    return (self).f26

            while lambda116_(d_2250_v5_):
                with _dafny.c_label("11"):
                    if (d_2251_i1_) >= (100):
                        raise _dafny.Break("11")
                    d_2251_i1_ = (d_2251_i1_) + (1)
                    d_2252_v6_: _dafny.MultiSet
                    d_2252_v6_ = _dafny.MultiSet([(self).f26])
                    (globalState).f5 = (d_2252_v6_).ispropersubset((_dafny.MultiSet([(self).f26, not((self).f26), not((self).f26)])) - (d_2252_v6_))
                    d_2253_v7_: _dafny.Array
                    nw354_ = _dafny.Array(_dafny.Array(None, 0), 23)
                    d_2253_v7_ = nw354_
                    d_2254_v8_: _dafny.Array
                    nw355_ = _dafny.Array(int(0), 1)
                    d_2254_v8_ = nw355_
                    index355_ = default__.safeIndex(948, (d_2253_v7_).length(0))
                    (d_2253_v7_)[index355_] = d_2254_v8_
                    d_2255_v9_: _dafny.MultiSet
                    d_2255_v9_ = _dafny.MultiSet([d_2247_v2_, d_2247_v2_])
                    d_2256_v10_: str
                    d_2256_v10_ = _dafny.CodePoint('r')
                    d_2257_v11_: _dafny.Map
                    d_2257_v11_ = _dafny.Map({d_2256_v10_: (self).f26})
                    d_2258_v12_: _dafny.MultiSet
                    d_2258_v12_ = _dafny.MultiSet([d_2256_v10_])
                    d_2259_v13_: _dafny.Seq
                    d_2259_v13_ = _dafny.SeqWithoutIsStrInference([not((self).f26), (self).f26, (self).f26])
                    d_2260_v14_: D18
                    d_2260_v14_ = D18_DC50((d_2259_v13_)[default__.safeIndex(d_2247_v2_, len(d_2259_v13_))])
                    d_2261_v15_: _dafny.Set
                    d_2261_v15_ = _dafny.Set({(self).f26})
                    d_2262_v17_: _dafny.Set
                    d_2262_v17_ = _dafny.Set({528})
                    d_2263_v18_: _dafny.Map
                    d_2263_v18_ = _dafny.Map({d_2262_v17_: d_2247_v2_})
                    d_2264_v19_: _dafny.Array
                    nw356_ = _dafny.Array(None, 17)
                    nw356_[int(0)] = (self).f26
                    nw356_[int(1)] = ((d_2257_v11_)[d_2256_v10_] if (d_2256_v10_) in (d_2257_v11_) else not((self).f26))
                    nw356_[int(2)] = True
                    nw356_[int(3)] = default__.fm2((d_2258_v12_).set(_dafny.CodePoint('x'), default__.abs(d_2247_v2_)), (0) - (d_2247_v2_), globalState)
                    nw356_[int(4)] = (self).f26
                    nw356_[int(5)] = (self).f26
                    nw356_[int(6)] = (self).f26
                    nw356_[int(7)] = (d_2260_v14_).cf96
                    nw356_[int(8)] = default__.fm2(_dafny.MultiSet([d_2256_v10_, d_2256_v10_]), len(d_2259_v13_), globalState)
                    nw356_[int(9)] = (False) == ((self).f26)
                    nw356_[int(10)] = (self).f26
                    nw356_[int(11)] = (self).f26
                    nw356_[int(12)] = (self).f26
                    nw356_[int(13)] = False
                    def iife252_():
                        coll62_ = _dafny.Set()
                        compr_62_: int
                        for compr_62_ in _dafny.IntegerRange(-109, 517):
                            d_2266_v16_: int = compr_62_
                            if ((-109) <= (d_2266_v16_)) and ((d_2266_v16_) < (517)):
                                coll62_ = coll62_.union(_dafny.Set([(d_2266_v16_) + (935)]))
                        return _dafny.Set(coll62_)
                    nw356_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_2256_v10_ for d_2265_i2_ in range(default__.abs(687))])) == (default__.fm27(default__.fm0(len(_dafny.Map({(self).f26: d_2247_v2_})), D0_DC1(d_2256_v10_, d_2247_v2_, d_2261_v15_, (self).f26, (self).f26), globalState), d_2255_v9_, len(iife252_()
                    ), d_2263_v18_, globalState))
                    nw356_[int(15)] = (self).f26
                    nw356_[int(16)] = (self).f26
                    d_2264_v19_ = nw356_
                    index356_ = default__.safeIndex(948, (d_2253_v7_).length(0))
                    rhs337_ = not((d_2255_v9_).ispropersubset(d_2255_v9_))
                    rhs338_ = d_2264_v19_
                    rhs339_ = (0) - ((0) - ((d_2247_v2_) * (d_2247_v2_)))
                    rhs340_ = (d_2254_v8_ if (self).f26 else d_2254_v8_)
                    lhs312_ = globalState
                    lhs313_ = d_2253_v7_
                    lhs314_ = default__.safeIndex(948, (d_2253_v7_).length(0))
                    lhs312_.f6 = rhs337_
                    r1 = rhs338_
                    d_2247_v2_ = rhs339_
                    lhs313_[lhs314_] = rhs340_
                    (globalState).f17 = ((d_2247_v2_) + ((0) - (d_2247_v2_)) if (self).f26 else d_2247_v2_)
                    if True:
                        (globalState).f6 = not ((self).f26) or ((self).f26)
                        arr2_ = (d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]
                        index357_ = default__.safeIndex(718, ((d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]).length(0))
                        arr2_[index357_] = d_2247_v2_
                        d_2267_v20_: _dafny.Seq
                        d_2267_v20_ = _dafny.SeqWithoutIsStrInference([d_2247_v2_])
                        d_2268_v21_: D10
                        d_2268_v21_ = D10_DC26(d_2267_v20_)
                        d_2269_v22_: _dafny.Map
                        d_2269_v22_ = _dafny.Map({d_2268_v21_: (self).f26})
                        d_2270_v23_: _dafny.Map
                        d_2270_v23_ = _dafny.Map({default__.fm2(d_2258_v12_, d_2247_v2_, globalState): (self).f26})
                        d_2271_v24_: _dafny.Seq
                        d_2271_v24_ = _dafny.SeqWithoutIsStrInference([d_2267_v20_])
                        d_2272_v25_: _dafny.Seq
                        d_2272_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdkxbsgd"))
                        arr3_ = (d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]
                        index358_ = default__.safeIndex(718, ((d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]).length(0))
                        rhs341_ = ((_dafny.MultiSet([(self).f26, (self).f26, (self).f26])).intersection(d_2252_v6_)).cardinality
                        rhs342_ = (_dafny.Map({D10_DC26(d_2267_v20_): False})) | (((default__.fm68(d_2247_v2_, d_2247_v2_, d_2272_v25_, (default__.fm51(globalState)).cardinality, globalState)).set(d_2268_v21_, not(False))) | (d_2269_v22_))
                        rhs343_ = (d_2270_v23_) | (d_2270_v23_)
                        rhs344_ = d_2271_v24_
                        rhs345_ = (self).f26
                        lhs315_ = (d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]
                        lhs316_ = default__.safeIndex(718, ((d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]).length(0))
                        lhs317_ = globalState
                        lhs315_[lhs316_] = rhs341_
                        d_2269_v22_ = rhs342_
                        d_2270_v23_ = rhs343_
                        d_2271_v24_ = rhs344_
                        lhs317_.f6 = rhs345_
                        d_2273_v26_: D7
                        d_2273_v26_ = D7_DC19(d_2254_v8_)
                        pat_let_tv96_ = d_2254_v8_
                        d_2274_v27_: _dafny.Array
                        nw357_ = _dafny.Array(None, 26)
                        nw357_[int(0)] = d_2273_v26_
                        nw357_[int(1)] = d_2273_v26_
                        nw357_[int(2)] = d_2273_v26_
                        nw357_[int(3)] = d_2273_v26_
                        nw357_[int(4)] = d_2273_v26_
                        nw357_[int(5)] = d_2273_v26_
                        nw357_[int(6)] = d_2273_v26_
                        nw357_[int(7)] = d_2273_v26_
                        nw357_[int(8)] = d_2273_v26_
                        nw357_[int(9)] = d_2273_v26_
                        nw357_[int(10)] = D7_DC19(d_2254_v8_)
                        nw357_[int(11)] = d_2273_v26_
                        nw357_[int(12)] = d_2273_v26_
                        nw357_[int(13)] = d_2273_v26_
                        def iife253_(_pat_let95_0):
                            def iife254_(d_2275_dt__update__tmp_h0_):
                                def iife255_(_pat_let96_0):
                                    def iife256_(d_2276_dt__update_hcf35_h0_):
                                        return D7_DC19(d_2276_dt__update_hcf35_h0_)
                                    return iife256_(_pat_let96_0)
                                return iife255_(pat_let_tv96_)
                            return iife254_(_pat_let95_0)
                        nw357_[int(14)] = iife253_(d_2273_v26_)
                        nw357_[int(15)] = D7_DC19(d_2254_v8_)
                        nw357_[int(16)] = d_2273_v26_
                        nw357_[int(17)] = d_2273_v26_
                        nw357_[int(18)] = d_2273_v26_
                        nw357_[int(19)] = d_2273_v26_
                        nw357_[int(20)] = D7_DC19(d_2254_v8_)
                        nw357_[int(21)] = d_2273_v26_
                        nw357_[int(22)] = d_2273_v26_
                        nw357_[int(23)] = D7_DC19(d_2254_v8_)
                        nw357_[int(24)] = d_2273_v26_
                        nw357_[int(25)] = d_2273_v26_
                        d_2274_v27_ = nw357_
                        d_2277_v28_: _dafny.Map
                        d_2277_v28_ = _dafny.Map({d_2272_v25_: d_2274_v27_})
                        d_2278_v29_: _dafny.Seq
                        d_2278_v29_ = _dafny.SeqWithoutIsStrInference([d_2274_v27_])
                        d_2277_v28_ = (d_2277_v28_).set((default__.fm34((self).f26, globalState)) + (d_2272_v25_), (d_2278_v29_)[default__.safeIndex(((d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))])[default__.safeIndex(718, ((d_2253_v7_)[default__.safeIndex(948, (d_2253_v7_).length(0))]).length(0))], len(d_2278_v29_))])
                        d_2279_v30_: _dafny.Seq
                        d_2279_v30_ = _dafny.SeqWithoutIsStrInference([d_2254_v8_, d_2254_v8_])
                        rhs346_ = (self).f26
                        rhs347_ = (d_2279_v30_) + (d_2279_v30_)
                        lhs318_ = globalState
                        lhs318_.f5 = rhs346_
                        d_2279_v30_ = rhs347_
                        d_2280_v31_: D13
                        d_2280_v31_ = D13_DC38()
                        rhs348_ = d_2247_v2_
                        rhs349_ = d_2280_v31_
                        lhs319_ = globalState
                        lhs319_.f18 = rhs348_
                        d_2280_v31_ = rhs349_
                    elif True:
                        d_2281_v32_: _dafny.Array
                        def lambda114_(d_2282_v10_):
                            def lambda115_(d_2283_i3_):
                                return d_2282_v10_

                            return lambda115_

                        init57_ = lambda114_(d_2256_v10_)
                        nw358_ = _dafny.Array(None, 19)
                        for i0_57_ in range(nw358_.length(0)):
                            nw358_[i0_57_] = init57_(i0_57_)
                        d_2281_v32_ = nw358_
                        index359_ = default__.safeIndex(623, (d_2281_v32_).length(0))
                        (d_2281_v32_)[index359_] = d_2256_v10_
                        index360_ = default__.safeIndex(623, (d_2281_v32_).length(0))
                        (d_2281_v32_)[index360_] = (d_2256_v10_ if (self).f26 else default__.fm1(globalState))
                        (globalState).f5 = (((self).f26) or ((self).f26)) in (d_2259_v13_)
                        (globalState).f17 = default__.safeModuloInt((d_2247_v2_) * (d_2247_v2_), d_2247_v2_)
                        (globalState).f5 = (self).f26
                        d_2284_v33_: _dafny.Seq
                        d_2284_v33_ = _dafny.SeqWithoutIsStrInference([((d_2255_v9_)[d_2247_v2_] if (d_2247_v2_) in (d_2255_v9_) else d_2247_v2_)])
                        index361_ = default__.safeIndex(238, (d_2254_v8_).length(0))
                        (d_2254_v8_)[index361_] = (d_2247_v2_ if (self).f26 else len(d_2284_v33_))
                        index362_ = default__.safeIndex(238, (d_2254_v8_).length(0))
                        (d_2254_v8_)[index362_] = default__.safeModuloInt((d_2247_v2_) + (len(d_2284_v33_)), default__.safeModuloInt(d_2247_v2_, len(d_2248_v3_)))
                    pass
            pass
        d_2289_v34_: _dafny.Seq
        d_2289_v34_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_2290_v35_: _dafny.Seq
        d_2290_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjslsen"))
        d_2291_v36_: str
        d_2291_v36_ = _dafny.CodePoint('h')
        d_2292_v37_: _dafny.MultiSet
        d_2292_v37_ = _dafny.MultiSet([d_2291_v36_, d_2291_v36_])
        d_2293_v38_: _dafny.MultiSet
        d_2293_v38_ = _dafny.MultiSet([(self).f26, False, (self).f26, (self).f26])
        d_2294_v39_: _dafny.Map
        d_2294_v39_ = _dafny.Map({(self).f26: d_2247_v2_})
        d_2295_v40_: _dafny.Seq
        d_2295_v40_ = _dafny.SeqWithoutIsStrInference([d_2294_v39_, d_2294_v39_])
        d_2296_v41_: _dafny.Set
        d_2296_v41_ = _dafny.Set({d_2290_v35_, d_2290_v35_})
        d_2297_v42_: _dafny.Array
        nw359_ = _dafny.Array(None, 24)
        nw359_[int(0)] = (self).f26
        nw359_[int(1)] = (self).f26
        nw359_[int(2)] = (d_2289_v34_)[default__.safeIndex(len(d_2290_v35_), len(d_2289_v34_))]
        nw359_[int(3)] = (self).f26
        nw359_[int(4)] = ((self).f26 if (self).f26 else True)
        nw359_[int(5)] = ((self).f26) == ((self).f26)
        nw359_[int(6)] = default__.fm2(d_2292_v37_, d_2247_v2_, globalState)
        nw359_[int(7)] = not ((self).f26) or ((self).f26)
        nw359_[int(8)] = (self).f26
        nw359_[int(9)] = (self).f26
        nw359_[int(10)] = (self).f26
        nw359_[int(11)] = (self).f26
        nw359_[int(12)] = (self).f26
        nw359_[int(13)] = (self).f26
        nw359_[int(14)] = (self).f26
        nw359_[int(15)] = (d_2289_v34_)[default__.safeIndex(d_2247_v2_, len(d_2289_v34_))]
        nw359_[int(16)] = default__.fm2(d_2292_v37_, (d_2293_v38_).cardinality, globalState)
        nw359_[int(17)] = (self).f26
        nw359_[int(18)] = (d_2294_v39_) not in (d_2295_v40_)
        nw359_[int(19)] = (d_2296_v41_).issubset(d_2296_v41_)
        nw359_[int(20)] = False
        nw359_[int(21)] = (self).f26
        nw359_[int(22)] = (self).f26
        nw359_[int(23)] = (self).f26
        d_2297_v42_ = nw359_
        index363_ = default__.safeIndex(979, (d_2297_v42_).length(0))
        (d_2297_v42_)[index363_] = (self).f26
        d_2298_v43_: _dafny.Set
        d_2298_v43_ = _dafny.Set({(self).f26})
        d_2299_v44_: D0
        d_2299_v44_ = D0_DC1(d_2291_v36_, d_2247_v2_, d_2298_v43_, (self).f26, False)
        d_2300_v45_: D7
        d_2300_v45_ = D7_DC20(default__.fm0(d_2247_v2_, d_2299_v44_, globalState), not((self).f26))
        index364_ = default__.safeIndex(979, (d_2297_v42_).length(0))
        (d_2297_v42_)[index364_] = (d_2300_v45_).cf37
        d_2247_v2_ = d_2247_v2_
        d_2301_v46_: _dafny.Map
        d_2301_v46_ = _dafny.Map({-82: (self).f26})
        hi8_ = len(d_2301_v46_)
        for d_2302_i4_ in range(d_2247_v2_, hi8_):
            index365_ = default__.safeIndex(979, (d_2297_v42_).length(0))
            rhs350_ = default__.safeDivisionInt(d_2302_i4_, d_2302_i4_)
            rhs351_ = not((d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))])
            rhs352_ = ((d_2291_v36_ if (self).f26 else d_2291_v36_)) in (d_2290_v35_)
            lhs320_ = globalState
            lhs321_ = globalState
            lhs322_ = d_2297_v42_
            lhs323_ = default__.safeIndex(979, (d_2297_v42_).length(0))
            lhs320_.f17 = rhs350_
            lhs321_.f5 = rhs351_
            lhs322_[lhs323_] = rhs352_
            d_2303_v47_: _dafny.MultiSet
            d_2303_v47_ = _dafny.MultiSet([d_2298_v43_, (d_2298_v43_).intersection(d_2298_v43_)])
            d_2304_v48_: _dafny.Seq
            d_2304_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))]}), d_2298_v43_])
            d_2303_v47_ = _dafny.MultiSet(d_2304_v48_)
            d_2305_v49_: _dafny.Map
            d_2305_v49_ = _dafny.Map({-805: d_2290_v35_})
            d_2306_v50_: _dafny.Map
            d_2306_v50_ = _dafny.Map({default__.safeModuloInt(d_2247_v2_, d_2302_i4_): ((d_2305_v49_)[d_2247_v2_] if (d_2247_v2_) in (d_2305_v49_) else d_2290_v35_)})
            d_2290_v35_ = ((d_2305_v49_)[d_2302_i4_] if (d_2302_i4_) in (d_2305_v49_) else ((d_2305_v49_)[d_2302_i4_] if (d_2302_i4_) in (d_2305_v49_) else _dafny.SeqWithoutIsStrInference([d_2291_v36_, d_2291_v36_])))
            d_2307_v52_: _dafny.Array
            def lambda117_(d_2308_v2_, d_2309_v3_):
                def lambda118_(d_2310_i5_):
                    def iife257_():
                        coll63_ = _dafny.Set()
                        compr_63_: int
                        for compr_63_ in (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_2308_v2_)), len(d_2309_v3_)])).Elements:
                            d_2311_v51_: int = compr_63_
                            if (d_2311_v51_) in (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_2308_v2_)), len(d_2309_v3_)])):
                                coll63_ = coll63_.union(_dafny.Set([(d_2311_v51_) * (43)]))
                        return _dafny.Set(coll63_)
                    return _dafny.SeqWithoutIsStrInference([len(iife257_()
                    )])

                return lambda118_

            init58_ = lambda117_(d_2247_v2_, d_2248_v3_)
            nw360_ = _dafny.Array(None, 1)
            for i0_58_ in range(nw360_.length(0)):
                nw360_[i0_58_] = init58_(i0_58_)
            d_2307_v52_ = nw360_
            d_2312_v53_: D3
            d_2312_v53_ = D3_DC7(d_2307_v52_)
            d_2313_v54_: D3
            d_2313_v54_ = D3_DC9(d_2312_v53_)
            source29_ = d_2313_v54_
            if source29_.is_DC8:
                d_2314___mcc_h2_ = source29_.cf16
                d_2315_cf16_ = d_2314___mcc_h2_
                (globalState).f6 = True
                d_2316_v55_: _dafny.Map
                d_2316_v55_ = _dafny.Map({default__.fm1(globalState): len((d_2290_v35_ if False else d_2290_v35_))})
                d_2316_v55_ = (d_2316_v55_) | (_dafny.Map({d_2291_v36_: 93}))
                (globalState).f6 = ((d_2247_v2_ if (d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))] else d_2302_i4_)) > (d_2247_v2_)
                d_2317_v56_: _dafny.Set
                d_2318_v57_: C0
                out23_: _dafny.Set
                out24_: C0
                out23_, out24_ = (self).m7(d_2291_v36_, globalState)
                d_2317_v56_ = out23_
                d_2318_v57_ = out24_
            elif source29_.is_DC7:
                d_2319___mcc_h3_ = source29_.cf15
                d_2320_cf15_ = d_2319___mcc_h3_
                (globalState).f17 = default__.safeModuloInt(d_2302_i4_, d_2247_v2_)
                r0 = default__.fm0(d_2247_v2_, d_2299_v44_, globalState)
                (globalState).f17 = (d_2302_i4_) * (d_2247_v2_)
                (globalState).f18 = 738
            elif True:
                d_2321___mcc_h4_ = source29_.cf17
                d_2322_cf17_ = d_2321___mcc_h4_
                d_2248_v3_ = (d_2248_v3_).set(823, (0) - (d_2302_i4_))
                d_2301_v46_ = d_2301_v46_
                d_2323_v58_: C1
                nw361_ = C1()
                nw361_.ctor__(d_2247_v2_, D0_DC0((d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))]))
                d_2323_v58_ = nw361_
                d_2324_v59_: D11
                d_2324_v59_ = D11_DC30(d_2323_v58_)
                d_2324_v59_ = d_2324_v59_
                (globalState).f5 = (default__.fm30((self).f26, d_2290_v35_, globalState)).issubset(_dafny.MultiSet([(d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))], (self).f26, not((d_2297_v42_)[default__.safeIndex(979, (d_2297_v42_).length(0))])]))
        (globalState).f17 = ((d_2247_v2_) * (d_2247_v2_)) - ((d_2247_v2_) - (d_2247_v2_))
        r0 = (d_2247_v2_) + (d_2247_v2_)
        r1 = (d_2297_v42_ if (self).f26 else d_2297_v42_)
        return r0, r1

    def m7(self, p0, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: C0 = None
        d_2325_v0_: _dafny.Array
        def lambda119_(d_2326_i0_):
            return (self).f26

        init59_ = lambda119_
        nw362_ = _dafny.Array(None, 25)
        for i0_59_ in range(nw362_.length(0)):
            nw362_[i0_59_] = init59_(i0_59_)
        d_2325_v0_ = nw362_
        d_2327_v1_: D14
        d_2327_v1_ = D14_DC39(d_2325_v0_)
        d_2328_v2_: _dafny.Array
        nw363_ = _dafny.Array(None, 26)
        nw363_[int(0)] = d_2325_v0_
        nw363_[int(1)] = d_2325_v0_
        nw363_[int(2)] = d_2325_v0_
        nw363_[int(3)] = d_2325_v0_
        nw363_[int(4)] = d_2325_v0_
        nw363_[int(5)] = d_2325_v0_
        nw363_[int(6)] = d_2325_v0_
        nw363_[int(7)] = d_2325_v0_
        nw363_[int(8)] = d_2325_v0_
        nw363_[int(9)] = d_2325_v0_
        nw363_[int(10)] = d_2325_v0_
        nw363_[int(11)] = d_2325_v0_
        nw363_[int(12)] = d_2325_v0_
        nw363_[int(13)] = d_2325_v0_
        nw363_[int(14)] = d_2325_v0_
        nw363_[int(15)] = d_2325_v0_
        nw363_[int(16)] = d_2325_v0_
        nw363_[int(17)] = d_2325_v0_
        nw363_[int(18)] = d_2325_v0_
        nw363_[int(19)] = (d_2327_v1_).cf72
        nw363_[int(20)] = d_2325_v0_
        nw363_[int(21)] = d_2325_v0_
        nw363_[int(22)] = d_2325_v0_
        nw363_[int(23)] = d_2325_v0_
        nw363_[int(24)] = d_2325_v0_
        nw363_[int(25)] = d_2325_v0_
        d_2328_v2_ = nw363_
        d_2329_v3_: D10
        d_2329_v3_ = D10_DC27(d_2328_v2_, len(_dafny.SeqWithoutIsStrInference([p0 for d_2330_i1_ in range(default__.abs(222))])))
        source30_ = (d_2329_v3_ if (self).f26 else d_2329_v3_)
        if source30_.is_DC27:
            d_2331___mcc_h0_ = source30_.cf50
            d_2332___mcc_h1_ = source30_.cf51
            d_2333_cf51_ = d_2332___mcc_h1_
            d_2334_cf50_ = d_2331___mcc_h0_
            d_2335_v4_: _dafny.Seq
            d_2335_v4_ = _dafny.SeqWithoutIsStrInference([d_2333_cf51_])
            d_2336_v5_: _dafny.Set
            d_2336_v5_ = _dafny.Set({(self).f26})
            d_2337_v6_: D0
            d_2337_v6_ = D0_DC1(p0, d_2333_cf51_, d_2336_v5_, (self).f26, (self).f26)
            d_2338_v7_: _dafny.Map
            d_2338_v7_ = _dafny.Map({default__.fm0((d_2335_v4_)[default__.safeIndex(d_2333_cf51_, len(d_2335_v4_))], d_2337_v6_, globalState): d_2328_v2_})
            d_2334_cf50_ = ((d_2338_v7_)[d_2333_cf51_] if (d_2333_cf51_) in (d_2338_v7_) else d_2334_cf50_)
            d_2339_v8_: _dafny.Seq
            d_2339_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rflqneoc"))
            d_2340_v9_: _dafny.Seq
            d_2340_v9_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26, (self).f26])
            d_2341_v10_: C0
            nw364_ = C0()
            nw364_.ctor__(d_2325_v0_, d_2340_v9_)
            d_2341_v10_ = nw364_
            d_2342_v11_: _dafny.Map
            d_2342_v11_ = _dafny.Map({d_2339_v8_: d_2341_v10_})
            d_2343_v12_: _dafny.Seq
            d_2343_v12_ = _dafny.SeqWithoutIsStrInference([d_2339_v8_])
            d_2344_v13_: _dafny.Map
            d_2344_v13_ = _dafny.Map({d_2333_cf51_: p0})
            d_2342_v11_ = (d_2342_v11_).set((d_2339_v8_) + ((d_2343_v12_)[default__.safeIndex(len(d_2344_v13_), len(d_2343_v12_))]), d_2341_v10_)
            arr4_ = d_2341_v10_.f23
            index366_ = default__.safeIndex(250, (d_2341_v10_.f23).length(0))
            arr4_[index366_] = (self).f26
            arr5_ = d_2341_v10_.f23
            index367_ = default__.safeIndex(250, (d_2341_v10_.f23).length(0))
            arr5_[index367_] = (self).f26
            (globalState).f6 = (d_2336_v5_).ispropersubset(_dafny.Set({True}))
        elif source30_.is_DC28:
            d_2345___mcc_h2_ = source30_.cf52
            d_2346___mcc_h3_ = source30_.cf53
            d_2347___mcc_h4_ = source30_.cf54
            d_2348___mcc_h5_ = source30_.cf55
            d_2349___mcc_h6_ = source30_.cf56
            d_2350_cf56_ = d_2349___mcc_h6_
            d_2351_cf55_ = d_2348___mcc_h5_
            d_2352_cf54_ = d_2347___mcc_h4_
            d_2353_cf53_ = d_2346___mcc_h3_
            d_2354_cf52_ = d_2345___mcc_h2_
            if (self).f26:
                d_2355_v14_: _dafny.Set
                d_2355_v14_ = _dafny.Set({d_2351_cf55_})
                d_2356_v15_: _dafny.Map
                d_2356_v15_ = _dafny.Map({(-653 if (self).f26 else d_2353_cf53_): d_2355_v14_})
                rhs353_ = d_2356_v15_
                rhs354_ = d_2353_cf53_
                rhs355_ = d_2353_cf53_
                lhs324_ = globalState
                lhs325_ = globalState
                d_2356_v15_ = rhs353_
                lhs324_.f17 = rhs354_
                lhs325_.f17 = rhs355_
                d_2357_v16_: D10
                d_2357_v16_ = D10_DC28(d_2354_cf52_, d_2353_cf53_, (self).f26, True, d_2350_cf56_)
                d_2358_v17_: _dafny.Map
                d_2358_v17_ = _dafny.Map({len(default__.fm46(d_2353_cf53_, (d_2357_v16_).cf55, True, globalState)): d_2353_cf53_})
                d_2359_v18_: _dafny.Map
                d_2359_v18_ = _dafny.Map({((d_2358_v17_)[d_2353_cf53_] if (d_2353_cf53_) in (d_2358_v17_) else 246): default__.fm0((0) - (d_2353_cf53_), D0_DC1(p0, d_2353_cf53_, d_2355_v14_, d_2352_cf54_, (self).f26), globalState)})
                d_2360_v19_: _dafny.Map
                d_2360_v19_ = _dafny.Map({d_2353_cf53_: d_2351_cf55_})
                d_2361_v20_: _dafny.Map
                d_2361_v20_ = _dafny.Map({d_2353_cf53_: d_2360_v19_})
                d_2358_v17_ = (d_2358_v17_).set(len((((d_2361_v20_)[d_2353_cf53_] if (d_2353_cf53_) in (d_2361_v20_) else d_2360_v19_)) | (d_2360_v19_)), default__.safeDivisionInt(d_2353_cf53_, d_2353_cf53_))
                d_2362_v21_: str
                d_2362_v21_ = _dafny.CodePoint('b')
                d_2362_v21_ = default__.fm1(globalState)
                d_2363_v22_: _dafny.Array
                nw365_ = _dafny.Array(int(0), 2)
                d_2363_v22_ = nw365_
                d_2364_v23_: _dafny.Map
                d_2364_v23_ = _dafny.Map({(self).f26: d_2363_v22_})
                d_2363_v22_ = ((d_2364_v23_)[d_2352_cf54_] if (d_2352_cf54_) in (d_2364_v23_) else d_2363_v22_)
                (globalState).f15 = d_2354_cf52_
            elif True:
                d_2365_v24_: _dafny.Map
                d_2365_v24_ = _dafny.Map({d_2353_cf53_: d_2353_cf53_})
                d_2366_v25_: _dafny.Map
                d_2366_v25_ = _dafny.Map({True: (0) - (d_2353_cf53_)})
                d_2367_v27_: _dafny.Set
                def iife258_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(358, -460):
                        d_2368_v26_: int = compr_64_
                        if ((358) <= (d_2368_v26_)) and ((d_2368_v26_) < (-460)):
                            coll64_[default__.safeModuloInt(d_2368_v26_, d_2353_cf53_)] = d_2353_cf53_
                    return _dafny.Map(coll64_)
                d_2367_v27_ = _dafny.Set({(d_2365_v24_).set(d_2353_cf53_, (0) - (d_2353_cf53_)), d_2365_v24_, d_2365_v24_, _dafny.Map({len(d_2366_v25_): d_2353_cf53_}), (iife258_()
                 if not(d_2351_cf55_) else d_2365_v24_)})
                d_2369_v28_: _dafny.Seq
                d_2369_v28_ = _dafny.SeqWithoutIsStrInference([d_2352_cf54_])
                rhs356_ = d_2353_cf53_
                rhs357_ = ((d_2367_v27_) | (d_2367_v27_)) - ((default__.fm69((d_2369_v28_)[default__.safeIndex(d_2353_cf53_, len(d_2369_v28_))], globalState)).intersection(d_2367_v27_))
                lhs326_ = globalState
                lhs326_.f18 = rhs356_
                d_2367_v27_ = rhs357_
                d_2370_v30_: _dafny.Map
                d_2370_v30_ = _dafny.Map({d_2353_cf53_: d_2369_v28_})
                d_2371_v31_: D5
                def iife259_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in (d_2370_v30_).keys.Elements:
                        d_2372_v29_: int = compr_65_
                        if (d_2372_v29_) in (d_2370_v30_):
                            coll65_[default__.safeDivisionInt(d_2372_v29_, d_2353_cf53_)] = (D3_DC8(d_2351_cf55_)).cf16
                    return _dafny.Map(coll65_)
                d_2371_v31_ = D5_DC13(iife259_()
)
                d_2373_v34_: _dafny.Map
                d_2373_v34_ = _dafny.Map({620: (self).f26})
                d_2374_v35_: _dafny.Array
                nw366_ = _dafny.Array(None, 19)
                nw366_[int(0)] = d_2371_v31_
                def iife260_():
                    def iife262_():
                        coll68_ = _dafny.Set()
                        compr_68_: int
                        for compr_68_ in _dafny.IntegerRange(130, -661):
                            d_2377_v33_: int = compr_68_
                            if ((130) <= (d_2377_v33_)) and ((d_2377_v33_) < (-661)):
                                coll68_ = coll68_.union(_dafny.Set([(d_2377_v33_) - (d_2353_cf53_)]))
                        return _dafny.Set(coll68_)
                    coll66_ = _dafny.Map()
                    def iife261_():
                        coll67_ = _dafny.Set()
                        compr_67_: int
                        for compr_67_ in _dafny.IntegerRange(130, -661):
                            d_2375_v33_: int = compr_67_
                            if ((130) <= (d_2375_v33_)) and ((d_2375_v33_) < (-661)):
                                coll67_ = coll67_.union(_dafny.Set([(d_2375_v33_) - (d_2353_cf53_)]))
                        return _dafny.Set(coll67_)
                    compr_66_: int
                    for compr_66_ in (iife261_()
                    ).Elements:
                        d_2376_v32_: int = compr_66_
                        if (d_2376_v32_) in (iife262_()
                        ):
                            coll66_[default__.safeModuloInt(d_2376_v32_, 30)] = (self).f26
                    return _dafny.Map(coll66_)
                nw366_[int(1)] = D5_DC13(iife260_()
)
                nw366_[int(2)] = d_2371_v31_
                nw366_[int(3)] = D5_DC13((d_2371_v31_).cf24)
                nw366_[int(4)] = d_2371_v31_
                nw366_[int(5)] = D5_DC13(d_2373_v34_)
                nw366_[int(6)] = d_2371_v31_
                nw366_[int(7)] = default__.fm70(d_2351_cf55_, False, globalState)
                nw366_[int(8)] = default__.fm70(d_2350_cf56_, d_2351_cf55_, globalState)
                nw366_[int(9)] = d_2371_v31_
                nw366_[int(10)] = d_2371_v31_
                nw366_[int(11)] = d_2371_v31_
                nw366_[int(12)] = d_2371_v31_
                nw366_[int(13)] = default__.fm70(d_2351_cf55_, d_2352_cf54_, globalState)
                nw366_[int(14)] = d_2371_v31_
                nw366_[int(15)] = d_2371_v31_
                nw366_[int(16)] = d_2371_v31_
                def iife263_(_pat_let97_0):
                    def iife264_(d_2378_dt__update__tmp_h0_):
                        def iife265_(_pat_let98_0):
                            def iife266_(d_2379_dt__update_hcf24_h0_):
                                return D5_DC13(d_2379_dt__update_hcf24_h0_)
                            return iife266_(_pat_let98_0)
                        return iife265_(_dafny.Map({-54: False}))
                    return iife264_(_pat_let97_0)
                nw366_[int(17)] = iife263_(d_2371_v31_)
                nw366_[int(18)] = d_2371_v31_
                d_2374_v35_ = nw366_
                index368_ = default__.safeIndex(519, (d_2374_v35_).length(0))
                (d_2374_v35_)[index368_] = d_2371_v31_
                index369_ = default__.safeIndex(519, (d_2374_v35_).length(0))
                (d_2374_v35_)[index369_] = d_2371_v31_
                d_2380_v36_: C1
                nw367_ = C1()
                nw367_.ctor__(d_2353_cf53_, D0_DC0((self).f26))
                d_2380_v36_ = nw367_
                d_2381_v37_: str
                d_2381_v37_ = _dafny.CodePoint('j')
                d_2382_v38_: _dafny.Set
                d_2382_v38_ = _dafny.Set({d_2353_cf53_})
                rhs358_ = (d_2382_v38_).ispropersubset((_dafny.Set({d_2353_cf53_})).intersection(d_2382_v38_))
                rhs359_ = d_2381_v37_
                rhs360_ = ((d_2369_v28_).set(default__.safeIndex((d_2353_cf53_) - (d_2353_cf53_), len(d_2369_v28_)), (not(d_2352_cf54_)) and (d_2350_cf56_))).set(default__.safeIndex((d_2353_cf53_) - (d_2353_cf53_), len((d_2369_v28_).set(default__.safeIndex((d_2353_cf53_) - (d_2353_cf53_), len(d_2369_v28_)), (not(d_2352_cf54_)) and (d_2350_cf56_)))), d_2352_cf54_)
                lhs327_ = globalState
                lhs328_ = r1
                lhs327_.f5 = rhs358_
                d_2381_v37_ = rhs359_
                lhs328_.f24 = rhs360_
                (globalState).f18 = d_2353_cf53_
            (globalState).f6 = d_2352_cf54_
            d_2383_v39_: D16
            d_2383_v39_ = D16_DC45(d_2353_cf53_, d_2350_cf56_, d_2353_cf53_)
            d_2384_v40_: _dafny.Seq
            d_2384_v40_ = _dafny.SeqWithoutIsStrInference([d_2383_v39_, d_2383_v39_, default__.fm71(globalState)])
            d_2385_v41_: _dafny.Seq
            d_2385_v41_ = _dafny.SeqWithoutIsStrInference([d_2325_v0_, d_2325_v0_])
            d_2386_v42_: _dafny.Seq
            d_2386_v42_ = _dafny.SeqWithoutIsStrInference([d_2353_cf53_])
            d_2387_v43_: _dafny.Map
            d_2387_v43_ = _dafny.Map({d_2353_cf53_: len(d_2386_v42_)})
            d_2388_v44_: _dafny.Map
            d_2388_v44_ = _dafny.Map({(d_2384_v40_) + (d_2384_v40_): (d_2385_v41_)[default__.safeIndex(((d_2387_v43_)[d_2353_cf53_] if (d_2353_cf53_) in (d_2387_v43_) else d_2353_cf53_), len(d_2385_v41_))]})
            d_2388_v44_ = d_2388_v44_
            if d_2351_cf55_:
                d_2389_v45_: _dafny.Map
                d_2389_v45_ = _dafny.Map({d_2386_v42_: d_2387_v43_})
                d_2389_v45_ = (d_2389_v45_).set(d_2386_v42_, (default__.fm31(d_2351_cf55_, d_2386_v42_, globalState)) | (default__.fm31(d_2352_cf54_, default__.fm37(d_2350_cf56_, d_2353_cf53_, globalState), globalState)))
                index370_ = default__.safeIndex(324, (d_2325_v0_).length(0))
                (d_2325_v0_)[index370_] = d_2350_cf56_
                index371_ = default__.safeIndex(324, (d_2325_v0_).length(0))
                rhs361_ = d_2350_cf56_
                rhs362_ = 358
                lhs329_ = d_2325_v0_
                lhs330_ = default__.safeIndex(324, (d_2325_v0_).length(0))
                lhs331_ = globalState
                lhs329_[lhs330_] = rhs361_
                lhs331_.f17 = rhs362_
                (globalState).f6 = (d_2351_cf55_ if d_2351_cf55_ else (d_2353_cf53_) >= (d_2353_cf53_))
                index372_ = default__.safeIndex(324, (d_2325_v0_).length(0))
                (d_2325_v0_)[index372_] = (d_2352_cf54_ if d_2351_cf55_ else d_2350_cf56_)
                (globalState).f5 = not((d_2353_cf53_) > (d_2353_cf53_))
            elif True:
                (globalState).f6 = d_2351_cf55_
                (globalState).f6 = not(not((self).f26))
                d_2390_v46_: _dafny.Array
                def lambda120_(d_2391_cf53_):
                    def lambda121_(d_2392_i2_):
                        return (d_2392_i2_) - (d_2391_cf53_)

                    return lambda121_

                init60_ = lambda120_(d_2353_cf53_)
                nw368_ = _dafny.Array(None, 10)
                for i0_60_ in range(nw368_.length(0)):
                    nw368_[i0_60_] = init60_(i0_60_)
                d_2390_v46_ = nw368_
                d_2393_v47_: _dafny.MultiSet
                d_2393_v47_ = _dafny.MultiSet([d_2353_cf53_])
                index373_ = default__.safeIndex(693, (d_2390_v46_).length(0))
                (d_2390_v46_)[index373_] = ((d_2393_v47_).intersection(d_2393_v47_)).cardinality
                index374_ = default__.safeIndex(693, (d_2390_v46_).length(0))
                (d_2390_v46_)[index374_] = d_2353_cf53_
                index375_ = default__.safeIndex(373, (d_2325_v0_).length(0))
                (d_2325_v0_)[index375_] = d_2352_cf54_
                index376_ = default__.safeIndex(373, (d_2325_v0_).length(0))
                (d_2325_v0_)[index376_] = d_2352_cf54_
                d_2394_v48_: _dafny.Seq
                d_2394_v48_ = _dafny.SeqWithoutIsStrInference([d_2350_cf56_])
                d_2395_v49_: _dafny.Map
                d_2395_v49_ = _dafny.Map({(d_2390_v46_)[default__.safeIndex(693, (d_2390_v46_).length(0))]: d_2350_cf56_})
                index377_ = default__.safeIndex(693, (d_2390_v46_).length(0))
                (d_2390_v46_)[index377_] = len((d_2394_v48_) + ((D6_DC18(d_2353_cf53_, -707, (d_2394_v48_).set(default__.safeIndex(d_2353_cf53_, len(d_2394_v48_)), d_2351_cf55_), (self).f26, len(d_2395_v49_))).cf32))
        elif source30_.is_DC26:
            d_2396___mcc_h7_ = source30_.cf49
            d_2397_cf49_ = d_2396___mcc_h7_
            d_2398_v50_: str
            d_2398_v50_ = _dafny.CodePoint('d')
            d_2398_v50_ = default__.fm1(globalState)
            if (self).f26:
                d_2399_v51_: int
                d_2399_v51_ = 37
                (globalState).f6 = (((0) - (d_2399_v51_) if (self).f26 else d_2399_v51_)) in (_dafny.Map({d_2399_v51_: d_2399_v51_}))
                d_2400_v52_: _dafny.Set
                d_2400_v52_ = _dafny.Set({d_2399_v51_, d_2399_v51_})
                d_2400_v52_ = d_2400_v52_
                index378_ = default__.safeIndex(858, (d_2325_v0_).length(0))
                (d_2325_v0_)[index378_] = (d_2399_v51_) <= (d_2399_v51_)
                index379_ = default__.safeIndex(858, (d_2325_v0_).length(0))
                (d_2325_v0_)[index379_] = (self).f26
                rhs363_ = d_2399_v51_
                rhs364_ = (d_2325_v0_)[default__.safeIndex(858, (d_2325_v0_).length(0))]
                lhs332_ = globalState
                lhs333_ = globalState
                lhs332_.f18 = rhs363_
                lhs333_.f5 = rhs364_
                d_2401_v53_: C6
                nw369_ = C6()
                nw369_.ctor__((self).f20)
                d_2401_v53_ = nw369_
            elif True:
                d_2402_v54_: _dafny.Map
                d_2402_v54_ = _dafny.Map({p0: (self).f26})
                d_2403_v55_: _dafny.Map
                d_2403_v55_ = d_2402_v54_
                d_2403_v55_ = (d_2402_v54_ if (self).f26 else (d_2402_v54_ if (self).f26 else d_2403_v55_))
                d_2404_v56_: int
                d_2404_v56_ = 205
                d_2405_v57_: _dafny.Seq
                d_2405_v57_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26])
                d_2406_v58_: _dafny.Set
                d_2406_v58_ = _dafny.Set({default__.safeDivisionInt(d_2404_v56_, d_2404_v56_), default__.safeModuloInt((0) - (len(d_2405_v57_)), d_2404_v56_), default__.safeDivisionInt(d_2404_v56_, d_2404_v56_)})
                d_2406_v58_ = d_2406_v58_
                d_2407_v59_: _dafny.Array
                def lambda122_(d_2408_v56_):
                    def lambda123_(d_2409_i3_):
                        return _dafny.Map({d_2408_v56_: d_2408_v56_})

                    return lambda123_

                init61_ = lambda122_(d_2404_v56_)
                nw370_ = _dafny.Array(None, 2)
                for i0_61_ in range(nw370_.length(0)):
                    nw370_[i0_61_] = init61_(i0_61_)
                d_2407_v59_ = nw370_
                d_2410_v60_: D9
                d_2410_v60_ = D9_DC25(d_2404_v56_, d_2404_v56_, d_2407_v59_)
                d_2411_v61_: _dafny.MultiSet
                d_2411_v61_ = _dafny.MultiSet([((d_2410_v60_).cf47) != (d_2404_v56_)])
                d_2411_v61_ = (d_2411_v61_) - (d_2411_v61_)
                d_2412_v62_: C0
                nw371_ = C0()
                nw371_.ctor__(d_2325_v0_, ((d_2405_v57_ if (self).f26 else d_2405_v57_)).set(default__.safeIndex(d_2404_v56_, len((d_2405_v57_ if (self).f26 else d_2405_v57_))), not((self).f26)))
                d_2412_v62_ = nw371_
                d_2413_v63_: _dafny.Array
                nw372_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_2413_v63_ = nw372_
                d_2414_v64_: D20
                d_2414_v64_ = D20_DC53(d_2413_v63_)
                d_2413_v63_ = (d_2414_v64_).cf99
            (globalState).f5 = ((self).f26) or ((self).f26)
            d_2415_v65_: _dafny.Array
            nw373_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_2415_v65_ = nw373_
            d_2416_v66_: _dafny.Array
            nw374_ = _dafny.Array(int(0), 10)
            d_2416_v66_ = nw374_
            index380_ = default__.safeIndex(492, (d_2415_v65_).length(0))
            (d_2415_v65_)[index380_] = d_2416_v66_
            index381_ = default__.safeIndex(492, (d_2415_v65_).length(0))
            (d_2415_v65_)[index381_] = d_2416_v66_
        elif True:
            d_2417___mcc_h8_ = source30_.cf57
            d_2418_cf57_ = d_2417___mcc_h8_
            d_2419_v67_: int
            d_2419_v67_ = -324
            d_2420_v68_: _dafny.Array
            def lambda124_(d_2421_v67_):
                def lambda125_(d_2422_i4_):
                    return (d_2422_i4_) + (d_2421_v67_)

                return lambda125_

            init62_ = lambda124_(d_2419_v67_)
            nw375_ = _dafny.Array(None, 24)
            for i0_62_ in range(nw375_.length(0)):
                nw375_[i0_62_] = init62_(i0_62_)
            d_2420_v68_ = nw375_
            d_2423_v69_: D7
            d_2423_v69_ = D7_DC21(d_2419_v67_, (self).f26, not((self).f26), d_2420_v68_, (self).f26)
            d_2424_v70_: _dafny.Map
            d_2424_v70_ = _dafny.Map({d_2423_v69_: d_2419_v67_})
            d_2425_v71_: D8
            d_2425_v71_ = D8_DC22(d_2424_v70_)
            d_2426_v72_: _dafny.Array
            nw376_ = _dafny.Array(None, 4)
            nw376_[int(0)] = d_2425_v71_
            nw376_[int(1)] = d_2425_v71_
            nw376_[int(2)] = d_2425_v71_
            nw376_[int(3)] = d_2425_v71_
            d_2426_v72_ = nw376_
            d_2426_v72_ = d_2426_v72_
            d_2427_v73_: C6
            nw377_ = C6()
            nw377_.ctor__((self).f20)
            d_2427_v73_ = nw377_
            d_2428_v74_: _dafny.Seq
            d_2428_v74_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
            d_2429_v75_: _dafny.Array
            nw378_ = _dafny.Array(None, 1)
            nw378_[int(0)] = d_2428_v74_
            d_2429_v75_ = nw378_
            d_2430_v76_: D15
            d_2430_v76_ = D15_DC43(d_2429_v75_, True, default__.safeModuloInt(d_2419_v67_, (0) - (d_2419_v67_)), (self).f26, D3_DC8((self).f26))
            d_2431_v77_: _dafny.Set
            d_2431_v77_ = _dafny.Set({(self).f26})
            d_2432_v78_: D0
            d_2432_v78_ = D0_DC1(p0, len(_dafny.Map({d_2419_v67_: d_2419_v67_})), d_2431_v77_, (self).f26, True)
            d_2433_v79_: D3
            d_2433_v79_ = D3_DC8((self).f26)
            pat_let_tv97_ = d_2419_v67_
            pat_let_tv98_ = d_2432_v78_
            pat_let_tv99_ = globalState
            pat_let_tv100_ = globalState
            def iife267_(_pat_let99_0):
                def iife268_(d_2434_dt__update__tmp_h2_):
                    def iife269_(_pat_let100_0):
                        def iife270_(d_2435_dt__update_hcf84_h0_):
                            def iife271_(_pat_let101_0):
                                def iife272_(d_2436_dt__update_hcf86_h0_):
                                    return D15_DC43((d_2434_dt__update__tmp_h2_).cf82, (d_2434_dt__update__tmp_h2_).cf83, d_2435_dt__update_hcf84_h0_, (d_2434_dt__update__tmp_h2_).cf85, d_2436_dt__update_hcf86_h0_)
                                return iife272_(_pat_let101_0)
                            return iife271_(default__.fm44(True, pat_let_tv100_))
                        return iife270_(_pat_let100_0)
                    return iife269_(default__.fm0(pat_let_tv97_, pat_let_tv98_, pat_let_tv99_))
                return iife268_(_pat_let99_0)
            d_2430_v76_ = iife267_(D15_DC43(d_2429_v75_, not((self).f26), default__.fm0((0) - (d_2419_v67_), d_2432_v78_, globalState), (self).f26, d_2433_v79_))
            d_2437_v80_: _dafny.Seq
            d_2437_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqufrkgvt"))
            d_2438_v81_: _dafny.Seq
            d_2438_v81_ = _dafny.SeqWithoutIsStrInference([(d_2437_v80_) + (d_2437_v80_), d_2437_v80_])
            (globalState).f15 = (d_2438_v81_)[default__.safeIndex(d_2419_v67_, len(d_2438_v81_))]
        index382_ = default__.safeIndex(394, (d_2325_v0_).length(0))
        (d_2325_v0_)[index382_] = (self).f26
        d_2439_v82_: _dafny.Array
        nw379_ = _dafny.Array(_dafny.Seq({}), 20)
        d_2439_v82_ = nw379_
        d_2440_v83_: int
        d_2440_v83_ = 134
        d_2441_v84_: _dafny.Seq
        d_2441_v84_ = _dafny.SeqWithoutIsStrInference([d_2440_v83_])
        d_2442_v85_: _dafny.Set
        d_2442_v85_ = _dafny.Set({not((self).f26)})
        index383_ = default__.safeIndex(568, (d_2439_v82_).length(0))
        (d_2439_v82_)[index383_] = (d_2441_v84_) + (_dafny.SeqWithoutIsStrInference([len(d_2442_v85_)]))
        d_2443_v86_: _dafny.Map
        d_2443_v86_ = _dafny.Map({d_2440_v83_: d_2325_v0_})
        d_2444_v87_: _dafny.Seq
        d_2444_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuechypmx"))
        d_2445_v88_: _dafny.Seq
        d_2445_v88_ = _dafny.SeqWithoutIsStrInference([(d_2443_v86_).set(len(d_2444_v87_), d_2325_v0_)])
        d_2446_v89_: D7
        d_2446_v89_ = D7_DC20(-754, (self).f26)
        index384_ = default__.safeIndex(394, (d_2325_v0_).length(0))
        index385_ = default__.safeIndex(568, (d_2439_v82_).length(0))
        rhs365_ = ((d_2445_v88_)[default__.safeIndex(608, len(d_2445_v88_))]) != (d_2443_v86_)
        rhs366_ = (d_2446_v89_).cf37
        rhs367_ = (d_2441_v84_ if (self).f26 else _dafny.SeqWithoutIsStrInference([d_2440_v83_]))
        rhs368_ = (self).f26
        lhs334_ = d_2325_v0_
        lhs335_ = default__.safeIndex(394, (d_2325_v0_).length(0))
        lhs336_ = globalState
        lhs337_ = d_2439_v82_
        lhs338_ = default__.safeIndex(568, (d_2439_v82_).length(0))
        lhs339_ = globalState
        lhs334_[lhs335_] = rhs365_
        lhs336_.f6 = rhs366_
        lhs337_[lhs338_] = rhs367_
        lhs339_.f6 = rhs368_
        d_2447_i5_: int
        d_2447_i5_ = 0
        with _dafny.label("12"):
            while (self).f26:
                with _dafny.c_label("12"):
                    if (d_2447_i5_) >= (100):
                        raise _dafny.Break("12")
                    d_2447_i5_ = (d_2447_i5_) + (1)
                    d_2448_v90_: D0
                    d_2448_v90_ = D0_DC1(p0, d_2440_v83_, d_2442_v85_, (self).f26, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))])
                    d_2449_v91_: _dafny.Array
                    nw380_ = _dafny.Array(None, 10)
                    nw380_[int(0)] = default__.fm0(d_2440_v83_, d_2448_v90_, globalState)
                    nw380_[int(1)] = d_2440_v83_
                    nw380_[int(2)] = (0) - (d_2440_v83_)
                    nw380_[int(3)] = len(d_2444_v87_)
                    nw380_[int(4)] = d_2440_v83_
                    nw380_[int(5)] = (d_2440_v83_) - (d_2440_v83_)
                    nw380_[int(6)] = default__.safeModuloInt((0) - (d_2440_v83_), d_2440_v83_)
                    nw380_[int(7)] = d_2440_v83_
                    nw380_[int(8)] = d_2440_v83_
                    nw380_[int(9)] = d_2440_v83_
                    d_2449_v91_ = nw380_
                    index386_ = default__.safeIndex(327, (d_2449_v91_).length(0))
                    (d_2449_v91_)[index386_] = d_2440_v83_
                    index387_ = default__.safeIndex(327, (d_2449_v91_).length(0))
                    (d_2449_v91_)[index387_] = d_2440_v83_
                    index388_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                    (d_2325_v0_)[index388_] = ((d_2449_v91_)[default__.safeIndex(327, (d_2449_v91_).length(0))]) <= (-507)
                    d_2450_v92_: _dafny.MultiSet
                    d_2450_v92_ = _dafny.MultiSet([(d_2449_v91_)[default__.safeIndex(327, (d_2449_v91_).length(0))]])
                    d_2451_v93_: _dafny.Seq
                    d_2451_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_2452_i6_ in range(default__.abs(-248))]), (d_2444_v87_).set(default__.safeIndex(730, len(d_2444_v87_)), p0)])
                    d_2453_v94_: _dafny.Seq
                    d_2453_v94_ = _dafny.SeqWithoutIsStrInference([(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], not((d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))])])
                    index389_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                    index390_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                    rhs369_ = not(((d_2451_v93_)[default__.safeIndex((0) - (d_2440_v83_), len(d_2451_v93_))]) < (d_2444_v87_))
                    rhs370_ = (self).f26
                    rhs371_ = (self).f26
                    rhs372_ = (d_2453_v94_)[default__.safeIndex(len(d_2453_v94_), len(d_2453_v94_))]
                    rhs373_ = d_2450_v92_
                    lhs340_ = d_2325_v0_
                    lhs341_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                    lhs342_ = globalState
                    lhs343_ = globalState
                    lhs344_ = d_2325_v0_
                    lhs345_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                    lhs340_[lhs341_] = rhs369_
                    lhs342_.f6 = rhs370_
                    lhs343_.f6 = rhs371_
                    lhs344_[lhs345_] = rhs372_
                    d_2450_v92_ = rhs373_
                    d_2454_v95_: _dafny.Map
                    d_2454_v95_ = _dafny.Map({d_2449_v91_: d_2327_v1_})
                    d_2454_v95_ = _dafny.Map({d_2449_v91_: d_2327_v1_})
                    pass
            pass
        d_2455_v96_: D0
        d_2455_v96_ = D0_DC1(p0, d_2440_v83_, d_2442_v85_, (self).f26, (self).f26)
        d_2456_i7_: int
        d_2456_i7_ = 0
        with _dafny.label("13"):
            while (d_2440_v83_) < (default__.fm0(d_2440_v83_, d_2455_v96_, globalState)):
                with _dafny.c_label("13"):
                    if (d_2456_i7_) >= (100):
                        raise _dafny.Break("13")
                    d_2456_i7_ = (d_2456_i7_) + (1)
                    d_2457_v97_: _dafny.Seq
                    d_2457_v97_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26, False])
                    d_2458_v98_: D12
                    d_2458_v98_ = D12_DC34((d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], not ((d_2457_v97_)[default__.safeIndex(d_2440_v83_, len(d_2457_v97_))]) or ((self).f26), (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))])
                    source31_ = d_2458_v98_
                    if source31_.is_DC33:
                        d_2459___mcc_h9_ = source31_.cf62
                        d_2460___mcc_h10_ = source31_.cf63
                        d_2461_cf63_ = d_2460___mcc_h10_
                        d_2462_cf62_ = d_2459___mcc_h9_
                        d_2463_v99_: _dafny.Set
                        d_2463_v99_ = _dafny.Set({d_2461_cf63_})
                        d_2464_v100_: _dafny.Map
                        d_2464_v100_ = _dafny.Map({(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]: d_2461_cf63_})
                        d_2465_v101_: _dafny.Map
                        d_2465_v101_ = _dafny.Map({d_2463_v99_: d_2461_cf63_})
                        d_2466_v102_: _dafny.Array
                        nw381_ = _dafny.Array(None, 10)
                        nw381_[int(0)] = len(default__.fm27(d_2461_cf63_, _dafny.MultiSet([d_2440_v83_]), len(d_2464_v100_), d_2465_v101_, globalState))
                        nw381_[int(1)] = 654
                        nw381_[int(2)] = d_2440_v83_
                        nw381_[int(3)] = d_2440_v83_
                        nw381_[int(4)] = ((d_2464_v100_)[(self).f26] if ((self).f26) in (d_2464_v100_) else d_2440_v83_)
                        nw381_[int(5)] = 76
                        nw381_[int(6)] = 311
                        nw381_[int(7)] = (0) - (d_2440_v83_)
                        nw381_[int(8)] = d_2440_v83_
                        nw381_[int(9)] = 975
                        d_2466_v102_ = nw381_
                        d_2467_v103_: _dafny.MultiSet
                        d_2467_v103_ = _dafny.MultiSet([d_2466_v102_, d_2466_v102_, d_2466_v102_])
                        d_2468_v104_: _dafny.Map
                        d_2468_v104_ = _dafny.Map({((self).f26) and ((self).f26): d_2467_v103_})
                        index391_ = default__.safeIndex(568, (d_2439_v82_).length(0))
                        rhs374_ = d_2463_v99_
                        rhs375_ = (self).f26
                        rhs376_ = ((d_2468_v104_)[(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]] if ((d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]) in (d_2468_v104_) else (((d_2468_v104_)[(self).f26] if ((self).f26) in (d_2468_v104_) else d_2467_v103_)).intersection(d_2467_v103_))
                        rhs377_ = ((d_2439_v82_)[default__.safeIndex(568, (d_2439_v82_).length(0))]) + ((d_2439_v82_)[default__.safeIndex(568, (d_2439_v82_).length(0))])
                        lhs346_ = globalState
                        lhs347_ = d_2439_v82_
                        lhs348_ = default__.safeIndex(568, (d_2439_v82_).length(0))
                        d_2463_v99_ = rhs374_
                        lhs346_.f5 = rhs375_
                        d_2467_v103_ = rhs376_
                        lhs347_[lhs348_] = rhs377_
                        (globalState).f6 = (self).f26
                        d_2469_v105_: _dafny.Array
                        nw382_ = _dafny.Array(None, 1)
                        nw382_[int(0)] = d_2457_v97_
                        d_2469_v105_ = nw382_
                        d_2470_v106_: _dafny.MultiSet
                        d_2470_v106_ = _dafny.MultiSet([d_2469_v105_])
                        d_2471_v107_: _dafny.Seq
                        d_2471_v107_ = _dafny.SeqWithoutIsStrInference([d_2444_v87_])
                        d_2472_v108_: _dafny.Seq
                        d_2472_v108_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2444_v87_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vitgtypm")), default__.fm46(len(_dafny.SeqWithoutIsStrInference([d_2461_cf63_ for d_2473_i8_ in range(default__.abs(948))])), (d_2457_v97_)[default__.safeIndex((0) - (d_2461_cf63_), len(d_2457_v97_))], (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], globalState)]), d_2471_v107_])
                        d_2474_v109_: _dafny.Array
                        nw383_ = _dafny.Array(None, 28)
                        nw383_[int(0)] = default__.fm72((self).f26, (self).f26, ((d_2470_v106_)[d_2469_v105_] if (d_2469_v105_) in (d_2470_v106_) else d_2440_v83_), globalState)
                        nw383_[int(1)] = d_2471_v107_
                        nw383_[int(2)] = d_2471_v107_
                        nw383_[int(3)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipce"))])
                        nw383_[int(4)] = (_dafny.SeqWithoutIsStrInference([default__.fm34((self).f26, globalState)])) + (d_2471_v107_)
                        nw383_[int(5)] = (d_2471_v107_) + (_dafny.SeqWithoutIsStrInference([d_2444_v87_]))
                        nw383_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2444_v87_])
                        nw383_[int(7)] = d_2471_v107_
                        nw383_[int(8)] = (d_2471_v107_) + (d_2471_v107_)
                        nw383_[int(9)] = ((d_2472_v108_)[default__.safeIndex(d_2440_v83_, len(d_2472_v108_))]) + (d_2471_v107_)
                        nw383_[int(10)] = d_2471_v107_
                        nw383_[int(11)] = ((d_2471_v107_).set(default__.safeIndex(d_2461_cf63_, len(d_2471_v107_)), d_2444_v87_)) + ((d_2471_v107_).set(default__.safeIndex(d_2440_v83_, len(d_2471_v107_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awjq"))))
                        nw383_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_2444_v87_])).set(default__.safeIndex(len(d_2442_v85_), len(_dafny.SeqWithoutIsStrInference([d_2444_v87_]))), d_2444_v87_)
                        nw383_[int(13)] = default__.fm72(default__.fm2(_dafny.MultiSet(d_2444_v87_), len((d_2441_v84_).set(default__.safeIndex(d_2440_v83_, len(d_2441_v84_)), d_2440_v83_)), globalState), (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], (d_2441_v84_)[default__.safeIndex((0) - (d_2440_v83_), len(d_2441_v84_))], globalState)
                        nw383_[int(14)] = (d_2471_v107_).set(default__.safeIndex(562, len(d_2471_v107_)), d_2444_v87_)
                        nw383_[int(15)] = d_2471_v107_
                        nw383_[int(16)] = d_2471_v107_
                        nw383_[int(17)] = (d_2471_v107_) + (d_2471_v107_)
                        nw383_[int(18)] = d_2471_v107_
                        nw383_[int(19)] = (d_2472_v108_)[default__.safeIndex(d_2461_cf63_, len(d_2472_v108_))]
                        nw383_[int(20)] = d_2471_v107_
                        nw383_[int(21)] = (d_2471_v107_) + (_dafny.SeqWithoutIsStrInference([d_2444_v87_ for d_2475_i9_ in range(default__.abs(557))]))
                        nw383_[int(22)] = (d_2471_v107_) + (_dafny.SeqWithoutIsStrInference([d_2444_v87_ for d_2476_i10_ in range(default__.abs(292))]))
                        nw383_[int(23)] = d_2471_v107_
                        nw383_[int(24)] = d_2471_v107_
                        nw383_[int(25)] = d_2471_v107_
                        nw383_[int(26)] = d_2471_v107_
                        nw383_[int(27)] = (_dafny.SeqWithoutIsStrInference([d_2444_v87_ for d_2477_i11_ in range(default__.abs(386))])) + (d_2471_v107_)
                        d_2474_v109_ = nw383_
                        index392_ = default__.safeIndex(102, (d_2474_v109_).length(0))
                        (d_2474_v109_)[index392_] = d_2471_v107_
                        index393_ = default__.safeIndex(102, (d_2474_v109_).length(0))
                        (d_2474_v109_)[index393_] = (d_2471_v107_).set(default__.safeIndex(d_2440_v83_, len(d_2471_v107_)), (d_2444_v87_).set(default__.safeIndex(d_2440_v83_, len(d_2444_v87_)), p0))
                        (globalState).f6 = (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]
                    elif source31_.is_DC34:
                        d_2478___mcc_h11_ = source31_.cf64
                        d_2479___mcc_h12_ = source31_.cf65
                        d_2480___mcc_h13_ = source31_.cf66
                        d_2481_cf66_ = d_2480___mcc_h13_
                        d_2482_cf65_ = d_2479___mcc_h12_
                        d_2483_cf64_ = d_2478___mcc_h11_
                        d_2484_v110_: _dafny.Map
                        d_2484_v110_ = _dafny.Map({d_2481_cf66_: (self).f26})
                        d_2485_v111_: _dafny.Seq
                        d_2485_v111_ = _dafny.SeqWithoutIsStrInference([d_2484_v110_, d_2484_v110_])
                        d_2486_v112_: _dafny.Map
                        d_2486_v112_ = _dafny.Map({d_2440_v83_: d_2485_v111_})
                        d_2485_v111_ = ((d_2486_v112_)[d_2440_v83_] if (d_2440_v83_) in (d_2486_v112_) else (d_2485_v111_).set(default__.safeIndex(d_2440_v83_, len(d_2485_v111_)), default__.fm73((d_2457_v97_)[default__.safeIndex(d_2440_v83_, len(d_2457_v97_))], True, d_2440_v83_, (d_2441_v84_)[default__.safeIndex(d_2440_v83_, len(d_2441_v84_))], globalState)))
                        d_2487_v113_: _dafny.Map
                        d_2487_v113_ = _dafny.Map({(d_2440_v83_) * ((0) - (len(d_2457_v97_))): default__.fm74((0) - (-198), globalState)})
                        d_2488_v114_: D4
                        d_2488_v114_ = D4_DC11(d_2482_cf65_, d_2482_cf65_, d_2440_v83_, d_2440_v83_)
                        pat_let_tv101_ = d_2481_cf66_
                        def iife273_(_pat_let102_0):
                            def iife274_(d_2489_dt__update__tmp_h3_):
                                def iife275_(_pat_let103_0):
                                    def iife276_(d_2490_dt__update_hcf19_h0_):
                                        return D4_DC11(d_2490_dt__update_hcf19_h0_, (d_2489_dt__update__tmp_h3_).cf20, (d_2489_dt__update__tmp_h3_).cf21, (d_2489_dt__update__tmp_h3_).cf22)
                                    return iife276_(_pat_let103_0)
                                return iife275_(pat_let_tv101_)
                            return iife274_(_pat_let102_0)
                        d_2487_v113_ = (d_2487_v113_).set(d_2440_v83_, iife273_(d_2488_v114_))
                        nw384_ = _dafny.Array(False, 4)
                        (r1).f23 = nw384_
                        d_2491_v115_: C1
                        nw385_ = C1()
                        nw385_.ctor__(len(d_2444_v87_), (self).f20)
                        d_2491_v115_ = nw385_
                    elif True:
                        d_2492___mcc_h14_ = source31_.cf61
                        d_2493_cf61_ = d_2492___mcc_h14_
                        index394_ = default__.safeIndex(394, (d_2325_v0_).length(0))
                        (d_2325_v0_)[index394_] = (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]
                        d_2494_v116_: _dafny.Seq
                        d_2494_v116_ = _dafny.SeqWithoutIsStrInference([d_2444_v87_, d_2444_v87_])
                        d_2495_v117_: _dafny.MultiSet
                        d_2495_v117_ = _dafny.MultiSet([len(d_2494_v116_)])
                        d_2496_v118_: C0
                        nw386_ = C0()
                        nw386_.ctor__(d_2325_v0_, ((d_2457_v97_).set(default__.safeIndex(d_2440_v83_, len(d_2457_v97_)), (self).f26)).set(default__.safeIndex(870, len((d_2457_v97_).set(default__.safeIndex(d_2440_v83_, len(d_2457_v97_)), (self).f26))), (d_2457_v97_)[default__.safeIndex((d_2495_v117_).cardinality, len(d_2457_v97_))]))
                        d_2496_v118_ = nw386_
                        (globalState).f17 = default__.safeModuloInt(d_2440_v83_, (d_2495_v117_).cardinality)
                        d_2497_v119_: _dafny.Array
                        nw387_ = _dafny.Array(_dafny.Array(None, 0), 20)
                        d_2497_v119_ = nw387_
                        d_2498_v120_: _dafny.MultiSet
                        d_2498_v120_ = _dafny.MultiSet([p0])
                        d_2499_v121_: _dafny.Map
                        d_2499_v121_ = _dafny.Map({d_2440_v83_: d_2442_v85_})
                        d_2500_v122_: _dafny.Array
                        nw388_ = _dafny.Array(None, 28)
                        nw388_[int(0)] = d_2442_v85_
                        nw388_[int(1)] = d_2442_v85_
                        nw388_[int(2)] = d_2442_v85_
                        nw388_[int(3)] = _dafny.Set({(self).f26, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]})
                        nw388_[int(4)] = d_2442_v85_
                        nw388_[int(5)] = d_2442_v85_
                        nw388_[int(6)] = d_2442_v85_
                        nw388_[int(7)] = d_2442_v85_
                        nw388_[int(8)] = d_2442_v85_
                        nw388_[int(9)] = d_2442_v85_
                        nw388_[int(10)] = d_2442_v85_
                        nw388_[int(11)] = d_2442_v85_
                        nw388_[int(12)] = d_2442_v85_
                        nw388_[int(13)] = default__.fm52((d_2498_v120_).cardinality, d_2444_v87_, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], False, globalState)
                        nw388_[int(14)] = d_2442_v85_
                        nw388_[int(15)] = ((d_2499_v121_)[d_2440_v83_] if (d_2440_v83_) in (d_2499_v121_) else _dafny.Set({(self).f26}))
                        nw388_[int(16)] = _dafny.Set({(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], (self).f26, (self).f26})
                        nw388_[int(17)] = d_2442_v85_
                        nw388_[int(18)] = ((d_2499_v121_)[d_2440_v83_] if (d_2440_v83_) in (d_2499_v121_) else _dafny.Set({(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]}))
                        nw388_[int(19)] = d_2442_v85_
                        nw388_[int(20)] = d_2442_v85_
                        nw388_[int(21)] = d_2442_v85_
                        nw388_[int(22)] = d_2442_v85_
                        nw388_[int(23)] = d_2442_v85_
                        nw388_[int(24)] = d_2442_v85_
                        nw388_[int(25)] = d_2442_v85_
                        nw388_[int(26)] = d_2442_v85_
                        nw388_[int(27)] = d_2442_v85_
                        d_2500_v122_ = nw388_
                        index395_ = default__.safeIndex(231, (d_2497_v119_).length(0))
                        (d_2497_v119_)[index395_] = d_2500_v122_
                        index396_ = default__.safeIndex(231, (d_2497_v119_).length(0))
                        nw389_ = _dafny.Array(None, 2)
                        nw389_[int(0)] = d_2442_v85_
                        nw389_[int(1)] = d_2442_v85_
                        (d_2497_v119_)[index396_] = nw389_
                    (globalState).f18 = d_2440_v83_
                    d_2501_v123_: _dafny.MultiSet
                    d_2501_v123_ = _dafny.MultiSet([(self).f26, not((self).f26), (self).f26, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]])
                    d_2440_v83_ = ((_dafny.MultiSet([not((self).f26), default__.fm2(_dafny.MultiSet([(d_2455_v96_).cf1]), (_dafny.MultiSet([(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], (self).f26, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]])).cardinality, globalState)])) - (d_2501_v123_)).cardinality
                    d_2502_v124_: _dafny.Array
                    nw390_ = _dafny.Array(_dafny.Seq({}), 3)
                    d_2502_v124_ = nw390_
                    d_2503_v125_: D20
                    d_2503_v125_ = D20_DC54((d_2440_v83_) - (d_2440_v83_), d_2502_v124_, d_2440_v83_, ((self).f26) == ((self).f26), len(d_2444_v87_))
                    d_2503_v125_ = d_2503_v125_
                    pass
            pass
        d_2504_v126_: _dafny.MultiSet
        d_2504_v126_ = _dafny.MultiSet([d_2440_v83_])
        d_2505_v127_: _dafny.Map
        d_2505_v127_ = _dafny.Map({d_2444_v87_: d_2504_v126_})
        d_2506_v128_: _dafny.Map
        d_2506_v128_ = _dafny.Map({(self).f26: (self).f26})
        d_2507_v129_: _dafny.Map
        d_2507_v129_ = _dafny.Map({d_2506_v128_: (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]})
        d_2508_v130_: _dafny.Array
        nw391_ = _dafny.Array(None, 18)
        nw391_[int(0)] = d_2504_v126_
        nw391_[int(1)] = (_dafny.MultiSet([len((d_2439_v82_)[default__.safeIndex(568, (d_2439_v82_).length(0))]), 285])) - (d_2504_v126_)
        nw391_[int(2)] = d_2504_v126_
        nw391_[int(3)] = d_2504_v126_
        nw391_[int(4)] = _dafny.MultiSet(d_2441_v84_)
        nw391_[int(5)] = (d_2504_v126_).set(d_2440_v83_, default__.abs(d_2440_v83_))
        nw391_[int(6)] = default__.fm33(globalState)
        nw391_[int(7)] = d_2504_v126_
        nw391_[int(8)] = d_2504_v126_
        nw391_[int(9)] = _dafny.MultiSet([d_2440_v83_])
        nw391_[int(10)] = d_2504_v126_
        nw391_[int(11)] = default__.fm56(d_2440_v83_, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], not((self).f26), d_2440_v83_, globalState)
        nw391_[int(12)] = d_2504_v126_
        nw391_[int(13)] = (d_2504_v126_) - (((d_2505_v127_)[(d_2444_v87_).set(default__.safeIndex(len(d_2444_v87_), len(d_2444_v87_)), p0)] if ((d_2444_v87_).set(default__.safeIndex(len(d_2444_v87_), len(d_2444_v87_)), p0)) in (d_2505_v127_) else _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([p0 for d_2509_i13_ in range(default__.abs(771))]))])))
        nw391_[int(14)] = _dafny.MultiSet([d_2440_v83_, len(d_2507_v129_)])
        nw391_[int(15)] = (d_2504_v126_).set((0) - (d_2440_v83_), default__.abs((0) - (d_2440_v83_)))
        nw391_[int(16)] = d_2504_v126_
        nw391_[int(17)] = d_2504_v126_
        d_2508_v130_ = nw391_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2508_v130_).length(0)):
            d_2510_i12_: int = guard_loop_5_
            if (True) and (((0) <= (d_2510_i12_)) and ((d_2510_i12_) < ((d_2508_v130_).length(0)))):
                (d_2508_v130_)[(d_2510_i12_)] = ((d_2504_v126_) - (d_2504_v126_)).intersection(_dafny.MultiSet(d_2441_v84_))
        d_2511_v131_: _dafny.MultiSet
        d_2511_v131_ = _dafny.MultiSet([p0, p0])
        d_2512_v132_: _dafny.Map
        d_2512_v132_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([(d_2439_v82_)[default__.safeIndex(568, (d_2439_v82_).length(0))] for d_2513_i15_ in range(default__.abs(-468))])): default__.fm2(d_2511_v131_, d_2440_v83_, globalState)})
        hi9_ = len(d_2512_v132_)
        for d_2514_i14_ in range(d_2440_v83_, hi9_):
            d_2515_v133_: C2
            nw392_ = C2()
            nw392_.ctor__(d_2514_i14_)
            d_2515_v133_ = nw392_
            (globalState).f18 = d_2440_v83_
            d_2516_v134_: _dafny.Set
            d_2516_v134_ = _dafny.Set({default__.fm0(-893, d_2455_v96_, globalState), (0) - ((0) - ((0) - (d_2514_i14_))), default__.fm0(d_2440_v83_, d_2455_v96_, globalState), d_2440_v83_, (d_2441_v84_)[default__.safeIndex(d_2440_v83_, len(d_2441_v84_))]})
            rhs378_ = d_2516_v134_
            rhs379_ = (d_2515_v133_).f22
            lhs349_ = globalState
            d_2516_v134_ = rhs378_
            lhs349_.f17 = rhs379_
            d_2517_v135_: _dafny.Map
            d_2517_v135_ = _dafny.Map({(d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]: (d_2515_v133_).f22})
            d_2518_v136_: _dafny.Array
            nw393_ = _dafny.Array(_dafny.Seq({}), 10)
            d_2518_v136_ = nw393_
            d_2519_v137_: _dafny.Seq
            d_2519_v137_ = _dafny.SeqWithoutIsStrInference([d_2444_v87_])
            d_2520_v138_: _dafny.Map
            d_2520_v138_ = _dafny.Map({D20_DC54((d_2515_v133_).f22, d_2518_v136_, d_2514_i14_, (self).f26, len((d_2519_v137_)[default__.safeIndex((d_2515_v133_).f22, len(d_2519_v137_))])): True})
            d_2521_v139_: _dafny.Array
            def lambda126_(d_2522_v133_):
                def lambda127_(d_2523_i16_):
                    return (d_2523_i16_) * ((d_2522_v133_).f22)

                return lambda127_

            init63_ = lambda126_(d_2515_v133_)
            nw394_ = _dafny.Array(None, 17)
            for i0_63_ in range(nw394_.length(0)):
                nw394_[i0_63_] = init63_(i0_63_)
            d_2521_v139_ = nw394_
            d_2524_v140_: _dafny.Map
            d_2524_v140_ = _dafny.Map({d_2521_v139_: (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))]})
            d_2525_v141_: _dafny.Array
            nw395_ = _dafny.Array(None, 22)
            nw395_[int(0)] = (0) - (d_2514_i14_)
            nw395_[int(1)] = d_2514_i14_
            nw395_[int(2)] = default__.safeModuloInt(len(_dafny.Map({(self).f26: not((self).f26)})), d_2440_v83_)
            nw395_[int(3)] = 590
            nw395_[int(4)] = (len(d_2517_v135_) if (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))] else default__.fm0(d_2514_i14_, d_2455_v96_, globalState))
            nw395_[int(5)] = (d_2515_v133_).f22
            nw395_[int(6)] = len(d_2444_v87_)
            nw395_[int(7)] = ((d_2515_v133_).f22) + (default__.fm0(d_2440_v83_, d_2455_v96_, globalState))
            nw395_[int(8)] = d_2440_v83_
            nw395_[int(9)] = d_2440_v83_
            nw395_[int(10)] = (d_2515_v133_).f22
            nw395_[int(11)] = ((d_2504_v126_)[(d_2515_v133_).f22] if ((d_2515_v133_).f22) in (d_2504_v126_) else d_2514_i14_)
            nw395_[int(12)] = ((d_2515_v133_).f22) + (len(d_2520_v138_))
            nw395_[int(13)] = 655
            nw395_[int(14)] = (len((d_2444_v87_).set(default__.safeIndex((d_2515_v133_).f22, len(d_2444_v87_)), p0))) + (len(d_2524_v140_))
            nw395_[int(15)] = default__.safeModuloInt(d_2440_v83_, (d_2515_v133_).f22)
            nw395_[int(16)] = d_2514_i14_
            nw395_[int(17)] = 633
            nw395_[int(18)] = (d_2514_i14_ if not((self).f26) else len(d_2506_v128_))
            nw395_[int(19)] = -202
            nw395_[int(20)] = default__.fm0((d_2515_v133_).f22, d_2455_v96_, globalState)
            nw395_[int(21)] = (d_2515_v133_).f22
            d_2525_v141_ = nw395_
            index397_ = default__.safeIndex(337, (d_2521_v139_).length(0))
            (d_2521_v139_)[index397_] = 157
            index398_ = default__.safeIndex(337, (d_2521_v139_).length(0))
            (d_2521_v139_)[index398_] = (len(d_2444_v87_)) * (694)
        d_2526_v142_: _dafny.Set
        d_2526_v142_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([len(d_2442_v85_)]), (d_2439_v82_)[default__.safeIndex(568, (d_2439_v82_).length(0))], d_2441_v84_})
        r0 = d_2526_v142_
        d_2527_v143_: _dafny.Seq
        d_2527_v143_ = _dafny.SeqWithoutIsStrInference([(self).f26])
        d_2528_v144_: _dafny.Map
        d_2528_v144_ = _dafny.Map({d_2440_v83_: d_2440_v83_})
        d_2529_v145_: _dafny.Map
        d_2529_v145_ = _dafny.Map({(self).f26: 64})
        d_2530_v146_: _dafny.Array
        nw396_ = _dafny.Array(None, 25)
        nw396_[int(0)] = d_2440_v83_
        nw396_[int(1)] = d_2440_v83_
        nw396_[int(2)] = d_2440_v83_
        nw396_[int(3)] = ((d_2528_v144_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxdcf")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxdcf")))) in (d_2528_v144_) else (0) - (d_2440_v83_))
        nw396_[int(4)] = d_2440_v83_
        nw396_[int(5)] = 742
        nw396_[int(6)] = d_2440_v83_
        nw396_[int(7)] = ((d_2529_v145_)[True] if (True) in (d_2529_v145_) else d_2440_v83_)
        nw396_[int(8)] = d_2440_v83_
        nw396_[int(9)] = -538
        nw396_[int(10)] = d_2440_v83_
        nw396_[int(11)] = d_2440_v83_
        nw396_[int(12)] = d_2440_v83_
        nw396_[int(13)] = d_2440_v83_
        nw396_[int(14)] = d_2440_v83_
        nw396_[int(15)] = 710
        nw396_[int(16)] = d_2440_v83_
        nw396_[int(17)] = d_2440_v83_
        nw396_[int(18)] = d_2440_v83_
        nw396_[int(19)] = len(d_2444_v87_)
        nw396_[int(20)] = d_2440_v83_
        nw396_[int(21)] = d_2440_v83_
        nw396_[int(22)] = d_2440_v83_
        nw396_[int(23)] = d_2440_v83_
        nw396_[int(24)] = len(d_2442_v85_)
        d_2530_v146_ = nw396_
        d_2531_v147_: D7
        d_2531_v147_ = D7_DC21(d_2440_v83_, (self).f26, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))], d_2530_v146_, (d_2325_v0_)[default__.safeIndex(394, (d_2325_v0_).length(0))])
        d_2532_v148_: C0
        nw397_ = C0()
        nw397_.ctor__(d_2325_v0_, ((d_2527_v143_) + (d_2527_v143_)).set(default__.safeIndex(d_2440_v83_, len((d_2527_v143_) + (d_2527_v143_))), (d_2531_v147_).cf39))
        d_2532_v148_ = nw397_
        r1 = d_2532_v148_
        return r0, r1

    @property
    def f26(self):
        return self._f26

class C8(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        self._f25: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f25, f20):
        (self)._f25 = f25
        (self)._f20 = f20

    def fm22(self, p0, p1, p2, globalState):
        return (self).f25

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        d_2533_v0_: _dafny.Array
        nw398_ = _dafny.Array(int(0), 28)
        d_2533_v0_ = nw398_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_2533_v0_).length(0)):
            d_2534_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_2534_i0_)) and ((d_2534_i0_) < ((d_2533_v0_).length(0)))):
                (d_2533_v0_)[(d_2534_i0_)] = default__.safeDivisionInt(d_2534_i0_, (len(_dafny.SeqWithoutIsStrInference([p1, p1]))) + (len(_dafny.SeqWithoutIsStrInference([D0_DC1(_dafny.CodePoint('a'), len(_dafny.Set({p1, p1})), _dafny.Set({(self).f25}), (self).f25, p0), D0_DC1(_dafny.CodePoint('d'), 625, _dafny.Set({(self).f25, p0}), p0, not(p0))]))))
        d_2535_v1_: _dafny.Array
        nw399_ = _dafny.Array(_dafny.Seq({}), 29)
        d_2535_v1_ = nw399_
        d_2536_v2_: _dafny.Seq
        d_2536_v2_ = _dafny.SeqWithoutIsStrInference([d_2533_v0_, d_2533_v0_, d_2533_v0_])
        index399_ = default__.safeIndex(276, (d_2535_v1_).length(0))
        (d_2535_v1_)[index399_] = (d_2536_v2_).set(default__.safeIndex(p1, len(d_2536_v2_)), d_2533_v0_)
        index400_ = default__.safeIndex(276, (d_2535_v1_).length(0))
        (d_2535_v1_)[index400_] = d_2536_v2_
        d_2537_v3_: _dafny.Array
        def lambda128_(d_2538_p0_):
            def lambda129_(d_2539_i1_):
                return ((self).f25 if (self).f25 else d_2538_p0_)

            return lambda129_

        init64_ = lambda128_(p0)
        nw400_ = _dafny.Array(None, 20)
        for i0_64_ in range(nw400_.length(0)):
            nw400_[i0_64_] = init64_(i0_64_)
        d_2537_v3_ = nw400_
        (globalState).f4 = d_2537_v3_
        d_2540_v4_: str
        d_2540_v4_ = _dafny.CodePoint('l')
        pat_let_tv102_ = p1
        pat_let_tv103_ = p0
        pat_let_tv104_ = p0
        def lambda130_(source33_):
            d_2541___mcc_h8_ = source33_
            d_2542_cf6_ = d_2541___mcc_h8_
            return D2_DC4(pat_let_tv102_, pat_let_tv103_, pat_let_tv104_, (self).f25)

        source32_ = lambda130_(default__.fm23(d_2540_v4_, False, p1, (self).f25, globalState))
        if source32_.is_DC4:
            d_2543___mcc_h0_ = source32_.cf8
            d_2544___mcc_h1_ = source32_.cf9
            d_2545___mcc_h2_ = source32_.cf10
            d_2546___mcc_h3_ = source32_.cf11
            d_2547_cf11_ = d_2546___mcc_h3_
            d_2548_cf10_ = d_2545___mcc_h2_
            d_2549_cf9_ = d_2544___mcc_h1_
            d_2550_cf8_ = d_2543___mcc_h0_
            source34_ = D0_DC0(((self).f25) == (p0))
            if source34_.is_DC1:
                d_2551___mcc_h9_ = source34_.cf1
                d_2552___mcc_h10_ = source34_.cf2
                d_2553___mcc_h11_ = source34_.cf3
                d_2554___mcc_h12_ = source34_.cf4
                d_2555___mcc_h13_ = source34_.cf5
                d_2556_cf5_ = d_2555___mcc_h13_
                d_2557_cf4_ = d_2554___mcc_h12_
                d_2558_cf3_ = d_2553___mcc_h11_
                d_2559_cf2_ = d_2552___mcc_h10_
                d_2560_cf1_ = d_2551___mcc_h9_
                d_2561_v5_: _dafny.Map
                d_2561_v5_ = _dafny.Map({d_2559_cf2_: d_2540_v4_})
                d_2561_v5_ = (d_2561_v5_).set(p1, d_2540_v4_)
                (globalState).f15 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukqmxahg"))
                (globalState).f6 = d_2556_cf5_
                d_2562_v6_: _dafny.MultiSet
                d_2562_v6_ = _dafny.MultiSet([d_2547_cf11_, p0])
                d_2562_v6_ = d_2562_v6_
            elif True:
                d_2563___mcc_h14_ = source34_.cf0
                d_2564_cf0_ = d_2563___mcc_h14_
                d_2565_v7_: _dafny.Array
                nw401_ = _dafny.Array(_dafny.Seq({}), 11)
                d_2565_v7_ = nw401_
                d_2566_v8_: D3
                d_2566_v8_ = D3_DC7(d_2565_v7_)
                d_2567_v9_: _dafny.Array
                nw402_ = _dafny.Array(None, 25)
                nw402_[int(0)] = d_2565_v7_
                nw402_[int(1)] = (d_2565_v7_ if d_2548_cf10_ else d_2565_v7_)
                nw402_[int(2)] = (d_2566_v8_).cf15
                nw402_[int(3)] = d_2565_v7_
                nw402_[int(4)] = d_2565_v7_
                nw402_[int(5)] = d_2565_v7_
                nw402_[int(6)] = d_2565_v7_
                nw402_[int(7)] = d_2565_v7_
                nw402_[int(8)] = d_2565_v7_
                nw402_[int(9)] = d_2565_v7_
                nw402_[int(10)] = d_2565_v7_
                nw402_[int(11)] = d_2565_v7_
                nw402_[int(12)] = d_2565_v7_
                nw402_[int(13)] = d_2565_v7_
                nw402_[int(14)] = d_2565_v7_
                nw402_[int(15)] = d_2565_v7_
                nw402_[int(16)] = d_2565_v7_
                nw402_[int(17)] = d_2565_v7_
                nw402_[int(18)] = d_2565_v7_
                nw402_[int(19)] = d_2565_v7_
                nw402_[int(20)] = d_2565_v7_
                nw402_[int(21)] = d_2565_v7_
                nw402_[int(22)] = d_2565_v7_
                nw402_[int(23)] = d_2565_v7_
                nw402_[int(24)] = d_2565_v7_
                d_2567_v9_ = nw402_
                d_2568_v10_: _dafny.Seq
                d_2568_v10_ = _dafny.SeqWithoutIsStrInference([d_2565_v7_, d_2565_v7_])
                index401_ = default__.safeIndex(955, (d_2567_v9_).length(0))
                (d_2567_v9_)[index401_] = (d_2568_v10_)[default__.safeIndex((0) - (d_2550_cf8_), len(d_2568_v10_))]
                index402_ = default__.safeIndex(955, (d_2567_v9_).length(0))
                (d_2567_v9_)[index402_] = d_2565_v7_
                d_2569_v11_: _dafny.MultiSet
                d_2569_v11_ = _dafny.MultiSet([d_2564_cf0_])
                d_2570_v12_: T0
                nw403_ = C4()
                nw403_.ctor__((self).f20)
                d_2570_v12_ = nw403_
                d_2571_v13_: _dafny.Map
                d_2571_v13_ = _dafny.Map({d_2569_v11_: d_2570_v12_})
                d_2572_v14_: _dafny.Map
                d_2572_v14_ = _dafny.Map({d_2548_cf10_: len(d_2571_v13_)})
                d_2573_v15_: D17
                d_2573_v15_ = D17_DC47(p1)
                d_2574_v16_: C1
                nw404_ = C1()
                nw404_.ctor__(len((d_2572_v14_) | (d_2572_v14_)), default__.fm57((0) - ((d_2573_v15_).cf92), globalState))
                d_2574_v16_ = nw404_
                d_2575_v17_: _dafny.MultiSet
                d_2575_v17_ = default__.fm30((self).f25, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kohhjbsd")), globalState)
                d_2576_v18_: C5
                nw405_ = C5()
                nw405_.ctor__(d_2569_v11_, d_2566_v8_, (self).f20)
                d_2576_v18_ = nw405_
                (globalState).f6 = d_2548_cf10_
            d_2577_v19_: _dafny.Seq
            d_2577_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjdicgyfc"))
            d_2578_v20_: _dafny.Map
            d_2578_v20_ = _dafny.Map({d_2577_v19_: d_2550_cf8_})
            d_2579_v21_: T0
            nw406_ = C7()
            nw406_.ctor__(False, (self).f20)
            d_2579_v21_ = nw406_
            d_2580_v22_: _dafny.Map
            d_2580_v22_ = _dafny.Map({d_2578_v20_: d_2579_v21_})
            d_2580_v22_ = (d_2580_v22_).set(d_2578_v20_, d_2579_v21_)
            d_2548_cf10_ = d_2548_cf10_
            d_2581_v23_: _dafny.Set
            d_2581_v23_ = _dafny.Set({d_2548_cf10_, d_2548_cf10_, (_dafny.Set({d_2549_cf9_})).ispropersubset(_dafny.Set({(self).f25}))})
            d_2581_v23_ = (d_2581_v23_ if d_2547_cf11_ else d_2581_v23_)
        elif source32_.is_DC5:
            d_2582___mcc_h4_ = source32_.cf12
            d_2583___mcc_h5_ = source32_.cf13
            d_2584_cf13_ = d_2583___mcc_h5_
            d_2585_cf12_ = d_2582___mcc_h4_
            d_2586_v24_: _dafny.Seq
            d_2586_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia"))
            d_2587_v25_: D2
            d_2587_v25_ = D2_DC3(d_2586_v24_)
            d_2588_v26_: _dafny.Array
            d_2589_v27_: int
            d_2590_v28_: int
            d_2591_v29_: int
            out25_: _dafny.Array
            out26_: int
            out27_: int
            out28_: int
            out25_, out26_, out27_, out28_ = (self).m5(d_2587_v25_, p1, p1, globalState)
            d_2588_v26_ = out25_
            d_2589_v27_ = out26_
            d_2590_v28_ = out27_
            d_2591_v29_ = out28_
            (globalState).f6 = True
            if d_2585_cf12_:
                d_2592_v30_: _dafny.MultiSet
                d_2592_v30_ = _dafny.MultiSet([d_2540_v4_])
                d_2593_v31_: _dafny.Set
                d_2593_v31_ = _dafny.Set({p0})
                d_2594_v32_: D0
                d_2594_v32_ = D0_DC1(d_2540_v4_, d_2584_cf13_, d_2593_v31_, (self).f25, (self).f25)
                (globalState).f6 = default__.fm2(d_2592_v30_, default__.fm0(p1, d_2594_v32_, globalState), globalState)
                d_2595_v33_: _dafny.MultiSet
                d_2595_v33_ = _dafny.MultiSet([d_2585_cf12_])
                d_2596_v34_: _dafny.Map
                d_2596_v34_ = _dafny.Map({False: d_2593_v31_})
                d_2597_v35_: _dafny.Seq
                d_2597_v35_ = _dafny.SeqWithoutIsStrInference([d_2590_v28_, (((d_2595_v33_)[p0] if (p0) in (d_2595_v33_) else 108)) - (len(d_2586_v24_)), len(((d_2596_v34_)[True] if (True) in (d_2596_v34_) else d_2593_v31_))])
                d_2598_v36_: _dafny.MultiSet
                d_2598_v36_ = _dafny.MultiSet([d_2589_v27_, d_2591_v29_, d_2590_v28_])
                (globalState).f17 = (d_2597_v35_)[default__.safeIndex((0) - ((d_2598_v36_).cardinality), len(d_2597_v35_))]
                d_2599_v37_: D16
                d_2599_v37_ = D16_DC45(p1, p0, d_2591_v29_)
                d_2600_v38_: _dafny.Seq
                d_2600_v38_ = _dafny.SeqWithoutIsStrInference([(self).f25, p0, (d_2599_v37_).cf89, (self).f25, (p1) > (d_2591_v29_)])
                d_2601_v39_: D6
                d_2601_v39_ = D6_DC18(len(d_2600_v38_), d_2591_v29_, d_2600_v38_, p0, 140)
                d_2600_v38_ = (d_2601_v39_).cf32
                (globalState).f5 = d_2585_cf12_
                d_2602_v40_: _dafny.Set
                d_2602_v40_ = _dafny.Set({d_2590_v28_})
                d_2603_v41_: _dafny.Set
                d_2603_v41_ = _dafny.Set({d_2589_v27_, 940, len((d_2586_v24_).set(default__.safeIndex(len(d_2602_v40_), len(d_2586_v24_)), d_2540_v4_))})
                (globalState).f6 = not((d_2602_v40_).isdisjoint(d_2602_v40_))
            elif True:
                d_2604_v42_: _dafny.Map
                d_2604_v42_ = _dafny.Map({d_2591_v29_: d_2590_v28_})
                index403_ = default__.safeIndex(663, (d_2533_v0_).length(0))
                (d_2533_v0_)[index403_] = ((d_2604_v42_)[d_2589_v27_] if (d_2589_v27_) in (d_2604_v42_) else d_2589_v27_)
                d_2605_v43_: _dafny.Map
                d_2605_v43_ = _dafny.Map({d_2585_cf12_: False})
                index404_ = default__.safeIndex(663, (d_2533_v0_).length(0))
                (d_2533_v0_)[index404_] = default__.safeDivisionInt(len((d_2605_v43_).set((self).f25, False)), d_2591_v29_)
                d_2590_v28_ = d_2589_v27_
                d_2606_v44_: C3
                nw407_ = C3()
                nw407_.ctor__(D0_DC0(not(True)))
                d_2606_v44_ = nw407_
                d_2607_v45_: _dafny.Array
                nw408_ = _dafny.Array(None, 15)
                nw408_[int(0)] = d_2606_v44_
                nw408_[int(1)] = d_2606_v44_
                nw408_[int(2)] = d_2606_v44_
                nw408_[int(3)] = d_2606_v44_
                nw408_[int(4)] = d_2606_v44_
                nw408_[int(5)] = d_2606_v44_
                nw408_[int(6)] = d_2606_v44_
                nw408_[int(7)] = d_2606_v44_
                nw408_[int(8)] = d_2606_v44_
                nw408_[int(9)] = d_2606_v44_
                nw408_[int(10)] = d_2606_v44_
                nw408_[int(11)] = d_2606_v44_
                nw408_[int(12)] = d_2606_v44_
                nw408_[int(13)] = d_2606_v44_
                nw408_[int(14)] = d_2606_v44_
                d_2607_v45_ = nw408_
                index405_ = default__.safeIndex(663, (d_2533_v0_).length(0))
                rhs380_ = len(_dafny.SeqWithoutIsStrInference([d_2607_v45_, d_2607_v45_]))
                rhs381_ = d_2585_cf12_
                lhs350_ = d_2533_v0_
                lhs351_ = default__.safeIndex(663, (d_2533_v0_).length(0))
                lhs352_ = globalState
                lhs350_[lhs351_] = rhs380_
                lhs352_.f5 = rhs381_
                (globalState).f6 = (self).f25
                index406_ = default__.safeIndex(312, (d_2537_v3_).length(0))
                (d_2537_v3_)[index406_] = True
                index407_ = default__.safeIndex(312, (d_2537_v3_).length(0))
                (d_2537_v3_)[index407_] = (self).f25
            if True:
                d_2608_v46_: _dafny.Set
                d_2608_v46_ = _dafny.Set({d_2586_v24_})
                d_2609_v47_: _dafny.Map
                d_2609_v47_ = _dafny.Map({len(d_2608_v46_): -407})
                d_2610_v48_: _dafny.Map
                d_2610_v48_ = _dafny.Map({d_2585_cf12_: not(d_2585_cf12_)})
                d_2611_v49_: _dafny.Seq
                d_2611_v49_ = _dafny.SeqWithoutIsStrInference([d_2585_cf12_, not((self).f25)])
                d_2612_v50_: _dafny.Seq
                d_2612_v50_ = _dafny.SeqWithoutIsStrInference([len(d_2611_v49_)])
                d_2609_v47_ = default__.fm31(((d_2610_v48_)[d_2585_cf12_] if (d_2585_cf12_) in (d_2610_v48_) else p0), (d_2612_v50_) + (_dafny.SeqWithoutIsStrInference([len(d_2609_v47_) for d_2613_i2_ in range(default__.abs(243))])), globalState)
                d_2614_v51_: _dafny.Set
                d_2614_v51_ = _dafny.Set({d_2584_cf13_, d_2584_cf13_, len(d_2586_v24_), len(d_2586_v24_)})
                d_2615_v52_: _dafny.Map
                d_2615_v52_ = _dafny.Map({d_2537_v3_: d_2614_v51_})
                d_2615_v52_ = (d_2615_v52_).set(d_2537_v3_, d_2614_v51_)
                d_2616_v53_: _dafny.Map
                d_2616_v53_ = _dafny.Map({default__.safeDivisionInt(d_2584_cf13_, d_2591_v29_): d_2588_v26_})
                d_2533_v0_ = ((d_2616_v53_)[len((d_2586_v24_) + (d_2586_v24_))] if (len((d_2586_v24_) + (d_2586_v24_))) in (d_2616_v53_) else d_2533_v0_)
                (globalState).f17 = (0) - (len(d_2611_v49_))
                d_2617_v54_: _dafny.Array
                nw409_ = _dafny.Array(_dafny.Set({}), 17)
                d_2617_v54_ = nw409_
                d_2618_v55_: _dafny.Set
                d_2618_v55_ = _dafny.Set({d_2611_v49_, default__.fm29(globalState), d_2611_v49_, d_2611_v49_})
                index408_ = default__.safeIndex(805, (d_2617_v54_).length(0))
                (d_2617_v54_)[index408_] = d_2618_v55_
                d_2619_v56_: _dafny.Map
                d_2619_v56_ = _dafny.Map({d_2586_v24_: _dafny.Set({d_2611_v49_, d_2611_v49_})})
                index409_ = default__.safeIndex(805, (d_2617_v54_).length(0))
                (d_2617_v54_)[index409_] = ((d_2619_v56_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idangmfg"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idangmfg"))) in (d_2619_v56_) else (d_2618_v55_ if False else d_2618_v55_))
            elif True:
                d_2589_v27_ = d_2591_v29_
                d_2620_v57_: C3
                nw410_ = C3()
                def iife277_(_pat_let104_0):
                    def iife278_(d_2621_dt__update__tmp_h1_):
                        def iife279_(_pat_let105_0):
                            def iife280_(d_2622_dt__update_hcf0_h0_):
                                return D0_DC0(d_2622_dt__update_hcf0_h0_)
                            return iife280_(_pat_let105_0)
                        return iife279_((self).f25)
                    return iife278_(_pat_let104_0)
                nw410_.ctor__(iife277_((self).f20))
                d_2620_v57_ = nw410_
                d_2623_v58_: _dafny.Seq
                d_2623_v58_ = _dafny.SeqWithoutIsStrInference([d_2589_v27_, p1])
                d_2624_v59_: _dafny.Set
                d_2624_v59_ = _dafny.Set({p0, (self).f25})
                d_2625_v60_: _dafny.Seq
                d_2625_v60_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(d_2586_v24_), D0_DC1(d_2540_v4_, len(d_2623_v58_), d_2624_v59_, d_2585_cf12_, p0), globalState)])
                d_2626_v61_: _dafny.Map
                d_2626_v61_ = _dafny.Map({d_2589_v27_: len(d_2625_v60_)})
                d_2627_v63_: _dafny.Map
                d_2627_v63_ = _dafny.Map({p1: d_2540_v4_})
                d_2628_v65_: _dafny.Set
                d_2628_v65_ = _dafny.Set({d_2584_cf13_})
                d_2629_v68_: _dafny.Map
                d_2629_v68_ = _dafny.Map({p0: d_2626_v61_})
                d_2630_v69_: _dafny.Map
                d_2630_v69_ = _dafny.Map({(self).f25: d_2585_cf12_})
                d_2631_v70_: _dafny.Seq
                d_2631_v70_ = _dafny.SeqWithoutIsStrInference([d_2630_v69_, _dafny.Map({d_2585_cf12_: p0})])
                d_2632_v71_: _dafny.Map
                d_2632_v71_ = _dafny.Map({True: d_2631_v70_})
                d_2633_v73_: _dafny.Array
                nw411_ = _dafny.Array(None, 27)
                def iife281_():
                    coll69_ = _dafny.Map()
                    compr_69_: int
                    for compr_69_ in (d_2627_v63_).keys.Elements:
                        d_2634_v62_: int = compr_69_
                        if (d_2634_v62_) in (d_2627_v63_):
                            coll69_[(d_2634_v62_) + (d_2591_v29_)] = d_2584_cf13_
                    return _dafny.Map(coll69_)
                nw411_[int(0)] = (d_2626_v61_ if True else iife281_()
                )
                nw411_[int(1)] = _dafny.Map({d_2591_v29_: d_2589_v27_})
                def iife282_():
                    coll70_ = _dafny.Map()
                    compr_70_: int
                    for compr_70_ in (d_2628_v65_).Elements:
                        d_2635_v64_: int = compr_70_
                        if (d_2635_v64_) in (d_2628_v65_):
                            coll70_[default__.safeModuloInt(d_2635_v64_, d_2591_v29_)] = len(_dafny.Map({p0: p0}))
                    return _dafny.Map(coll70_)
                nw411_[int(2)] = (d_2626_v61_) | (iife282_()
                )
                nw411_[int(3)] = d_2626_v61_
                nw411_[int(4)] = d_2626_v61_
                def iife283_():
                    coll71_ = _dafny.Map()
                    compr_71_: int
                    for compr_71_ in _dafny.IntegerRange(845, -190):
                        d_2636_v66_: int = compr_71_
                        if ((845) <= (d_2636_v66_)) and ((d_2636_v66_) < (-190)):
                            coll71_[default__.safeDivisionInt(d_2636_v66_, d_2590_v28_)] = d_2589_v27_
                    return _dafny.Map(coll71_)
                nw411_[int(5)] = iife283_()
                
                nw411_[int(6)] = d_2626_v61_
                nw411_[int(7)] = d_2626_v61_
                nw411_[int(8)] = d_2626_v61_
                nw411_[int(9)] = d_2626_v61_
                def iife284_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(885, 671):
                        d_2637_v67_: int = compr_72_
                        if ((885) <= (d_2637_v67_)) and ((d_2637_v67_) < (671)):
                            coll72_[default__.safeDivisionInt(d_2637_v67_, d_2591_v29_)] = p1
                    return _dafny.Map(coll72_)
                nw411_[int(10)] = (d_2626_v61_) | (iife284_()
                )
                nw411_[int(11)] = d_2626_v61_
                nw411_[int(12)] = d_2626_v61_
                nw411_[int(13)] = (((d_2629_v68_)[d_2585_cf12_] if (d_2585_cf12_) in (d_2629_v68_) else d_2626_v61_)) | (default__.fm31((self).f25, d_2625_v60_, globalState))
                nw411_[int(14)] = (d_2626_v61_) | (_dafny.Map({p1: p1}))
                nw411_[int(15)] = ((d_2626_v61_).set(d_2589_v27_, len(((d_2632_v71_)[d_2585_cf12_] if (d_2585_cf12_) in (d_2632_v71_) else _dafny.SeqWithoutIsStrInference([d_2630_v69_, d_2630_v69_]))))) | (d_2626_v61_)
                nw411_[int(16)] = d_2626_v61_
                nw411_[int(17)] = d_2626_v61_
                nw411_[int(18)] = d_2626_v61_
                nw411_[int(19)] = (d_2626_v61_).set(p1, d_2589_v27_)
                def iife285_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(206, 90):
                        d_2638_v72_: int = compr_73_
                        if ((206) <= (d_2638_v72_)) and ((d_2638_v72_) < (90)):
                            coll73_[(d_2638_v72_) + (d_2584_cf13_)] = len(d_2624_v59_)
                    return _dafny.Map(coll73_)
                nw411_[int(20)] = iife285_()
                
                nw411_[int(21)] = d_2626_v61_
                nw411_[int(22)] = default__.fm31(p0, _dafny.SeqWithoutIsStrInference([p1 for d_2639_i3_ in range(default__.abs(-762))]), globalState)
                nw411_[int(23)] = _dafny.Map({467: d_2584_cf13_})
                nw411_[int(24)] = default__.fm31((self).f25, (_dafny.SeqWithoutIsStrInference([-183, d_2589_v27_])).set(default__.safeIndex(d_2591_v29_, len(_dafny.SeqWithoutIsStrInference([-183, d_2589_v27_]))), p1), globalState)
                nw411_[int(25)] = default__.fm31(d_2585_cf12_, (d_2625_v60_).set(default__.safeIndex((0) - (d_2584_cf13_), len(d_2625_v60_)), p1), globalState)
                nw411_[int(26)] = d_2626_v61_
                d_2633_v73_ = nw411_
                d_2640_v75_: _dafny.Map
                d_2640_v75_ = _dafny.Map({len(d_2624_v59_): d_2585_cf12_})
                index410_ = default__.safeIndex(93, (d_2633_v73_).length(0))
                def iife286_():
                    coll74_ = _dafny.Map()
                    compr_74_: int
                    for compr_74_ in ((d_2640_v75_).set(924, d_2585_cf12_)).keys.Elements:
                        d_2641_v74_: int = compr_74_
                        if (d_2641_v74_) in ((d_2640_v75_).set(924, d_2585_cf12_)):
                            coll74_[(d_2641_v74_) + (d_2591_v29_)] = 826
                    return _dafny.Map(coll74_)
                (d_2633_v73_)[index410_] = iife286_()
                
                index411_ = default__.safeIndex(93, (d_2633_v73_).length(0))
                def iife287_():
                    coll75_ = _dafny.Map()
                    compr_75_: int
                    for compr_75_ in _dafny.IntegerRange(-787, 884):
                        d_2642_v76_: int = compr_75_
                        if ((-787) <= (d_2642_v76_)) and ((d_2642_v76_) < (884)):
                            coll75_[default__.safeDivisionInt(d_2642_v76_, len(d_2640_v75_))] = d_2584_cf13_
                    return _dafny.Map(coll75_)
                (d_2633_v73_)[index411_] = ((d_2629_v68_)[not(d_2585_cf12_)] if (not(d_2585_cf12_)) in (d_2629_v68_) else iife287_()
                )
                (globalState).f18 = d_2589_v27_
                d_2643_v77_: _dafny.Seq
                d_2643_v77_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2589_v27_: d_2585_cf12_}), d_2640_v75_, d_2640_v75_])
                d_2643_v77_ = (_dafny.SeqWithoutIsStrInference([d_2640_v75_])) + (d_2643_v77_)
        elif source32_.is_DC3:
            d_2644___mcc_h6_ = source32_.cf7
            d_2645_cf7_ = d_2644___mcc_h6_
            (globalState).f17 = (p1 if (p1) == (p1) else p1)
            (globalState).f5 = (self).f25
            d_2646_v78_: _dafny.MultiSet
            d_2646_v78_ = _dafny.MultiSet([d_2540_v4_, d_2540_v4_, d_2540_v4_])
            d_2646_v78_ = ((d_2646_v78_ if p0 else d_2646_v78_)) | ((d_2646_v78_) | (d_2646_v78_))
            d_2647_v79_: _dafny.Seq
            d_2647_v79_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            d_2648_v80_: C0
            nw412_ = C0()
            nw412_.ctor__(d_2537_v3_, d_2647_v79_)
            d_2648_v80_ = nw412_
        elif True:
            d_2649___mcc_h7_ = source32_.cf14
            d_2650_cf14_ = d_2649___mcc_h7_
            d_2651_v81_: _dafny.Seq
            d_2651_v81_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2651_v81_ = (d_2651_v81_) + (default__.fm29(globalState))
            d_2652_v82_: _dafny.Seq
            d_2652_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npwc"))
            d_2653_v83_: _dafny.Array
            nw413_ = _dafny.Array(None, 12)
            nw413_[int(0)] = d_2533_v0_
            nw413_[int(1)] = d_2533_v0_
            nw413_[int(2)] = d_2533_v0_
            nw413_[int(3)] = d_2533_v0_
            nw413_[int(4)] = d_2533_v0_
            nw413_[int(5)] = d_2533_v0_
            nw413_[int(6)] = (d_2536_v2_)[default__.safeIndex((0) - (len(d_2652_v82_)), len(d_2536_v2_))]
            nw413_[int(7)] = d_2533_v0_
            nw413_[int(8)] = d_2533_v0_
            nw413_[int(9)] = d_2533_v0_
            nw413_[int(10)] = d_2533_v0_
            nw413_[int(11)] = d_2533_v0_
            d_2653_v83_ = nw413_
            index412_ = default__.safeIndex(510, (d_2653_v83_).length(0))
            (d_2653_v83_)[index412_] = d_2533_v0_
            index413_ = default__.safeIndex(510, (d_2653_v83_).length(0))
            (d_2653_v83_)[index413_] = d_2533_v0_
            r1 = p1
            d_2654_v84_: _dafny.Set
            d_2654_v84_ = _dafny.Set({p0})
            pat_let_tv105_ = d_2654_v84_
            def iife288_(_pat_let106_0):
                def iife289_(d_2655_dt__update__tmp_h2_):
                    def iife290_(_pat_let107_0):
                        def iife291_(d_2656_dt__update_hcf3_h0_):
                            return D0_DC1((d_2655_dt__update__tmp_h2_).cf1, (d_2655_dt__update__tmp_h2_).cf2, d_2656_dt__update_hcf3_h0_, (d_2655_dt__update__tmp_h2_).cf4, (d_2655_dt__update__tmp_h2_).cf5)
                        return iife291_(_pat_let107_0)
                    return iife290_(pat_let_tv105_)
                return iife289_(_pat_let106_0)
            (globalState).f17 = (iife288_(D0_DC1(d_2540_v4_, p1, d_2654_v84_, True, p0))).cf2
        d_2657_v85_: _dafny.Seq
        d_2657_v85_ = _dafny.SeqWithoutIsStrInference([p1])
        (globalState).f5 = (_dafny.SeqWithoutIsStrInference([p1])) == (d_2657_v85_)
        d_2658_v86_: C6
        nw414_ = C6()
        nw414_.ctor__(D0_DC0(p0))
        d_2658_v86_ = nw414_
        d_2659_v87_: _dafny.MultiSet
        d_2659_v87_ = _dafny.MultiSet([p0, not((self).f25), (p0 if (self).f25 else (self).f25)])
        r0 = d_2659_v87_
        d_2660_v88_: _dafny.Seq
        d_2660_v88_ = _dafny.SeqWithoutIsStrInference([True])
        d_2661_v89_: _dafny.Set
        d_2661_v89_ = _dafny.Set({p0})
        d_2662_v90_: D0
        d_2662_v90_ = D0_DC1(d_2540_v4_, p1, d_2661_v89_, (self).f25, p0)
        r1 = (len((_dafny.SeqWithoutIsStrInference([(self).f25, p0])) + (d_2660_v88_))) - ((p1) - (default__.fm0(len(d_2657_v85_), d_2662_v90_, globalState)))
        r2 = (d_2540_v4_ if (self).f25 else d_2540_v4_)
        return r0, r1, r2

    def m5(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r2: int = int(0)
        r3: int = int(0)
        d_2663_v0_: _dafny.Array
        def lambda131_(d_2664_p1_, d_2665_p2_):
            def lambda132_(d_2666_i0_):
                return (_dafny.SeqWithoutIsStrInference([d_2664_p1_, d_2665_p2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('f'), _dafny.CodePoint('s'), _dafny.CodePoint('l'), _dafny.CodePoint('k')]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwb"))), d_2665_p2_]))

            return lambda132_

        init65_ = lambda131_(p1, p2)
        nw415_ = _dafny.Array(None, 13)
        for i0_65_ in range(nw415_.length(0)):
            nw415_[i0_65_] = init65_(i0_65_)
        d_2663_v0_ = nw415_
        d_2667_v1_: _dafny.Seq
        d_2667_v1_ = _dafny.SeqWithoutIsStrInference([True, (self).f25])
        d_2668_v2_: _dafny.Map
        d_2668_v2_ = _dafny.Map({p1: len(d_2667_v1_)})
        d_2669_v3_: _dafny.MultiSet
        d_2669_v3_ = _dafny.MultiSet([p2])
        d_2670_v4_: _dafny.Map
        d_2670_v4_ = _dafny.Map({p2: d_2669_v3_})
        d_2671_v5_: _dafny.Seq
        d_2671_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lilaf"))
        d_2672_v6_: _dafny.Map
        d_2672_v6_ = _dafny.Map({_dafny.CodePoint('h'): not (False) or ((self).f25)})
        d_2673_v7_: str
        d_2673_v7_ = _dafny.CodePoint('s')
        d_2674_v8_: _dafny.Set
        d_2674_v8_ = _dafny.Set({p1, p1})
        rhs382_ = default__.safeModuloInt(((d_2668_v2_)[len(d_2670_v4_)] if (len(d_2670_v4_)) in (d_2668_v2_) else p2), len(d_2671_v5_))
        rhs383_ = ((d_2672_v6_)[d_2673_v7_] if (d_2673_v7_) in (d_2672_v6_) else (d_2674_v8_).issubset(d_2674_v8_))
        rhs384_ = d_2663_v0_
        rhs385_ = (self).f25
        lhs353_ = globalState
        lhs354_ = globalState
        r3 = rhs382_
        lhs353_.f6 = rhs383_
        d_2663_v0_ = rhs384_
        lhs354_.f5 = rhs385_
        source35_ = (self).f20
        if source35_.is_DC1:
            d_2675___mcc_h0_ = source35_.cf1
            d_2676___mcc_h1_ = source35_.cf2
            d_2677___mcc_h2_ = source35_.cf3
            d_2678___mcc_h3_ = source35_.cf4
            d_2679___mcc_h4_ = source35_.cf5
            d_2680_cf5_ = d_2679___mcc_h4_
            d_2681_cf4_ = d_2678___mcc_h3_
            d_2682_cf3_ = d_2677___mcc_h2_
            d_2683_cf2_ = d_2676___mcc_h1_
            d_2684_cf1_ = d_2675___mcc_h0_
            r3 = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrt")) if (d_2680_cf5_ if not(d_2680_cf5_) else True) else d_2671_v5_))
            (globalState).f6 = (d_2671_v5_) <= (d_2671_v5_)
            d_2685_v9_: _dafny.MultiSet
            d_2685_v9_ = _dafny.MultiSet([d_2684_cf1_, d_2684_cf1_, d_2684_cf1_, _dafny.CodePoint('k')])
            d_2686_v10_: _dafny.Map
            d_2686_v10_ = _dafny.Map({(self).f25: (False) and (False)})
            d_2687_v11_: _dafny.MultiSet
            d_2687_v11_ = _dafny.MultiSet([d_2681_cf4_])
            d_2688_v12_: _dafny.Seq
            d_2688_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_2681_cf4_, (self).f25, d_2681_cf4_]), d_2687_v11_, (_dafny.MultiSet([d_2680_cf5_])).set((self).f25, default__.abs(p2)), d_2687_v11_])
            rhs386_ = (default__.fm2(d_2685_v9_, len(d_2682_cf3_), globalState)) == (not((self).f25))
            rhs387_ = ((d_2686_v10_)[((0) - (p2)) <= (p1)] if (((0) - (p2)) <= (p1)) in (d_2686_v10_) else d_2680_cf5_)
            rhs388_ = ((d_2687_v11_).isdisjoint(d_2687_v11_)) not in ((d_2688_v12_)[default__.safeIndex(p1, len(d_2688_v12_))])
            rhs389_ = (p2) - (len(_dafny.Map({d_2680_cf5_: d_2683_cf2_})))
            rhs390_ = 180
            lhs355_ = globalState
            lhs356_ = globalState
            d_2681_cf4_ = rhs386_
            lhs355_.f5 = rhs387_
            d_2680_cf5_ = rhs388_
            lhs356_.f18 = rhs389_
            r1 = rhs390_
            (globalState).f17 = default__.safeModuloInt(default__.safeModuloInt(p2, p1), default__.safeModuloInt(d_2683_cf2_, d_2683_cf2_))
        elif True:
            d_2689___mcc_h5_ = source35_.cf0
            d_2690_cf0_ = d_2689___mcc_h5_
            d_2691_v13_: _dafny.Map
            d_2691_v13_ = _dafny.Map({d_2671_v5_: p2})
            d_2691_v13_ = d_2691_v13_
            (globalState).f6 = d_2690_cf0_
            if (self).f25:
                d_2692_v14_: _dafny.Map
                d_2692_v14_ = _dafny.Map({d_2690_cf0_: d_2668_v2_})
                (globalState).f18 = len((d_2692_v14_).set(d_2690_cf0_, d_2668_v2_))
                d_2693_v15_: _dafny.MultiSet
                d_2693_v15_ = _dafny.MultiSet([d_2673_v7_, d_2673_v7_, d_2673_v7_])
                d_2694_v16_: _dafny.Set
                d_2694_v16_ = _dafny.Set({not(d_2690_cf0_), d_2690_cf0_})
                r2 = (0) - ((0) - (((d_2693_v15_)[d_2673_v7_] if (d_2673_v7_) in (d_2693_v15_) else (p1) + (len(d_2694_v16_)))))
                d_2695_v17_: D6
                d_2695_v17_ = D6_DC18(p2, p1, _dafny.SeqWithoutIsStrInference([(d_2667_v1_)[default__.safeIndex(p1, len(d_2667_v1_))]]), (self).f25, 350)
                (globalState).f18 = (d_2695_v17_).cf30
                d_2696_v18_: _dafny.Map
                d_2696_v18_ = _dafny.Map({p2: not((self).f25)})
                d_2697_v19_: _dafny.Map
                d_2697_v19_ = _dafny.Map({(d_2696_v18_) | (d_2696_v18_): p2})
                d_2697_v19_ = (d_2697_v19_).set((_dafny.Map({p1: (self).f25})) | (_dafny.Map({p1: d_2690_cf0_})), p2)
                d_2698_v20_: C4
                nw416_ = C4()
                nw416_.ctor__((self).f20)
                d_2698_v20_ = nw416_
                d_2699_v21_: _dafny.Map
                d_2699_v21_ = _dafny.Map({d_2698_v20_: p1})
                d_2699_v21_ = (d_2699_v21_).set(d_2698_v20_, (p1) - (p1))
            elif True:
                (globalState).f17 = p1
                d_2700_v22_: _dafny.Array
                nw417_ = _dafny.Array(False, 17)
                d_2700_v22_ = nw417_
                d_2701_v23_: _dafny.Set
                d_2701_v23_ = _dafny.Set({d_2690_cf0_})
                d_2702_v24_: D0
                d_2702_v24_ = D0_DC1(d_2673_v7_, (0) - (p1), d_2701_v23_, (self).f25, d_2690_cf0_)
                index414_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                (d_2700_v22_)[index414_] = (default__.fm0(p2, d_2702_v24_, globalState)) < (p1)
                d_2703_v25_: _dafny.Seq
                d_2703_v25_ = _dafny.SeqWithoutIsStrInference([304])
                index415_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                rhs391_ = ((d_2668_v2_) | (_dafny.Map({(d_2703_v25_)[default__.safeIndex(len(_dafny.Set({p2})), len(d_2703_v25_))]: p1}))) != (d_2668_v2_)
                rhs392_ = d_2690_cf0_
                rhs393_ = (self).f25
                lhs357_ = globalState
                lhs358_ = d_2700_v22_
                lhs359_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                lhs360_ = globalState
                lhs357_.f6 = rhs391_
                lhs358_[lhs359_] = rhs392_
                lhs360_.f5 = rhs393_
                d_2704_v26_: _dafny.Array
                nw418_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_2704_v26_ = nw418_
                index416_ = default__.safeIndex(861, (d_2704_v26_).length(0))
                (d_2704_v26_)[index416_] = _dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2705_i1_ in range(default__.abs(19))])
                index417_ = default__.safeIndex(861, (d_2704_v26_).length(0))
                (d_2704_v26_)[index417_] = _dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2706_i2_ in range(default__.abs(132))])
                d_2707_v27_: D7
                d_2707_v27_ = D7_DC20(492, (d_2700_v22_)[default__.safeIndex(428, (d_2700_v22_).length(0))])
                d_2708_v28_: _dafny.Map
                d_2708_v28_ = _dafny.Map({p2: not((self).f25)})
                d_2709_v29_: _dafny.Map
                d_2709_v29_ = _dafny.Map({(self).f25: p1})
                index418_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                rhs394_ = (d_2707_v27_ if (p2) in (d_2708_v28_) else d_2707_v27_)
                rhs395_ = (self).fm22(default__.safeDivisionInt(p2, p2), d_2690_cf0_, d_2709_v29_, globalState)
                rhs396_ = ((p1) * (p2) if d_2690_cf0_ else p2)
                lhs361_ = d_2700_v22_
                lhs362_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                d_2707_v27_ = rhs394_
                lhs361_[lhs362_] = rhs395_
                r3 = rhs396_
                index419_ = default__.safeIndex(428, (d_2700_v22_).length(0))
                (d_2700_v22_)[index419_] = (d_2700_v22_)[default__.safeIndex(428, (d_2700_v22_).length(0))]
            d_2710_v30_: _dafny.Array
            def lambda133_(d_2711_i3_):
                return True

            init66_ = lambda133_
            nw419_ = _dafny.Array(None, 16)
            for i0_66_ in range(nw419_.length(0)):
                nw419_[i0_66_] = init66_(i0_66_)
            d_2710_v30_ = nw419_
            (globalState).f4 = d_2710_v30_
        d_2712_v31_: _dafny.Seq
        d_2712_v31_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2713_v32_: _dafny.Set
        d_2713_v32_ = _dafny.Set({d_2712_v31_})
        d_2714_v33_: D12
        d_2714_v33_ = D12_DC33(d_2713_v32_, default__.safeDivisionInt(p1, p2))
        pat_let_tv106_ = d_2713_v32_
        def iife292_(_pat_let108_0):
            def iife293_(d_2715_dt__update__tmp_h0_):
                def iife294_(_pat_let109_0):
                    def iife295_(d_2716_dt__update_hcf62_h0_):
                        return D12_DC33(d_2716_dt__update_hcf62_h0_, (d_2715_dt__update__tmp_h0_).cf63)
                    return iife295_(_pat_let109_0)
                return iife294_(pat_let_tv106_)
            return iife293_(_pat_let108_0)
        d_2714_v33_ = iife292_(d_2714_v33_)
        d_2717_v34_: D6
        d_2717_v34_ = D6_DC18(p2, p2, d_2667_v1_, (self).f25, p2)
        pat_let_tv107_ = p1
        def lambda134_(source36_):
            if source36_.is_DC17:
                d_2718___mcc_h6_ = source36_.cf28
                d_2719___mcc_h7_ = source36_.cf29
                d_2720_cf29_ = d_2719___mcc_h7_
                d_2721_cf28_ = d_2718___mcc_h6_
                return (self).f25
            elif source36_.is_DC18:
                d_2722___mcc_h8_ = source36_.cf30
                d_2723___mcc_h9_ = source36_.cf31
                d_2724___mcc_h10_ = source36_.cf32
                d_2725___mcc_h11_ = source36_.cf33
                d_2726___mcc_h12_ = source36_.cf34
                d_2727_cf34_ = d_2726___mcc_h12_
                d_2728_cf33_ = d_2725___mcc_h11_
                d_2729_cf32_ = d_2724___mcc_h10_
                d_2730_cf31_ = d_2723___mcc_h9_
                d_2731_cf30_ = d_2722___mcc_h8_
                return (pat_let_tv107_) < (d_2731_cf30_)
            elif True:
                d_2732___mcc_h13_ = source36_.cf27
                d_2733_cf27_ = d_2732___mcc_h13_
                return False

        if lambda134_(d_2717_v34_):
            (globalState).f5 = ((self).f25) or (not((self).f25))
            d_2734_v35_: _dafny.Set
            d_2734_v35_ = _dafny.Set({(self).f25})
            d_2735_v36_: D0
            d_2735_v36_ = D0_DC1(d_2673_v7_, p1, d_2734_v35_, (self).f25, (d_2717_v34_).cf33)
            d_2736_v37_: _dafny.MultiSet
            d_2736_v37_ = _dafny.MultiSet([True, False])
            d_2737_v38_: _dafny.Seq
            d_2737_v38_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "meoipt")), _dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2738_i4_ in range(default__.abs(798))])])
            d_2739_v39_: _dafny.Array
            nw420_ = _dafny.Array(None, 9)
            nw420_[int(0)] = (default__.fm0(p1, d_2735_v36_, globalState)) - (p2)
            nw420_[int(1)] = (d_2669_v3_).cardinality
            nw420_[int(2)] = p2
            nw420_[int(3)] = len((d_2667_v1_) + (d_2667_v1_))
            nw420_[int(4)] = p1
            nw420_[int(5)] = ((d_2736_v37_)[(self).f25] if ((self).f25) in (d_2736_v37_) else len(_dafny.Map({p2: (d_2737_v38_)[default__.safeIndex((0) - (p2), len(d_2737_v38_))]})))
            nw420_[int(6)] = p1
            nw420_[int(7)] = (len(d_2671_v5_)) + (p1)
            nw420_[int(8)] = -597
            d_2739_v39_ = nw420_
            index420_ = default__.safeIndex(182, (d_2739_v39_).length(0))
            (d_2739_v39_)[index420_] = p1
            index421_ = default__.safeIndex(811, (d_2739_v39_).length(0))
            (d_2739_v39_)[index421_] = p1
            d_2740_v40_: C2
            nw421_ = C2()
            nw421_.ctor__(len(d_2671_v5_))
            d_2740_v40_ = nw421_
            d_2741_v41_: _dafny.Map
            d_2741_v41_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxsptslf"))): d_2673_v7_})
            d_2742_v42_: _dafny.Map
            d_2742_v42_ = _dafny.Map({d_2741_v41_: (self).f25})
            index422_ = default__.safeIndex(182, (d_2739_v39_).length(0))
            index423_ = default__.safeIndex(811, (d_2739_v39_).length(0))
            def iife296_():
                coll76_ = _dafny.Map()
                compr_76_: _dafny.Map
                for compr_76_ in (default__.fm75(p1, globalState)).Elements:
                    d_2743_v43_: _dafny.Map = compr_76_
                    if (d_2743_v43_) in (default__.fm75(p1, globalState)):
                        coll76_[d_2743_v43_] = (self).f25
                return _dafny.Map(coll76_)
            rhs397_ = p2
            rhs398_ = len(d_2734_v35_)
            rhs399_ = (len(_dafny.SeqWithoutIsStrInference([(self).f25, (self).f25, (self).f25, (self).f25, (self).f25]))) - (default__.safeModuloInt(((d_2669_v3_)[len(d_2671_v5_)] if (len(d_2671_v5_)) in (d_2669_v3_) else len(d_2674_v8_)), (d_2740_v40_).f22))
            rhs400_ = (d_2740_v40_ if (self).f25 else d_2740_v40_)
            rhs401_ = ((d_2742_v42_) | (iife296_()
            )) | (_dafny.Map({d_2741_v41_: (self).f25}))
            lhs363_ = d_2739_v39_
            lhs364_ = default__.safeIndex(182, (d_2739_v39_).length(0))
            lhs365_ = d_2739_v39_
            lhs366_ = default__.safeIndex(811, (d_2739_v39_).length(0))
            r1 = rhs397_
            lhs363_[lhs364_] = rhs398_
            lhs365_[lhs366_] = rhs399_
            d_2740_v40_ = rhs400_
            d_2742_v42_ = rhs401_
            d_2712_v31_ = d_2712_v31_
            (globalState).f5 = False
            (globalState).f6 = (self).f25
        elif True:
            r2 = len(d_2671_v5_)
            if (p1) > ((p2) - (p2)):
                d_2744_v44_: _dafny.Array
                def lambda135_(d_2745_i5_):
                    return (d_2745_i5_) * (len(_dafny.Map({(self).f25: (self).f25})))

                init67_ = lambda135_
                nw422_ = _dafny.Array(None, 23)
                for i0_67_ in range(nw422_.length(0)):
                    nw422_[i0_67_] = init67_(i0_67_)
                d_2744_v44_ = nw422_
                d_2746_v45_: _dafny.Seq
                d_2746_v45_ = _dafny.SeqWithoutIsStrInference([(d_2744_v44_ if (self).f25 else d_2744_v44_), d_2744_v44_, d_2744_v44_, d_2744_v44_, d_2744_v44_])
                d_2746_v45_ = d_2746_v45_
                (globalState).f5 = (self).f25
                (globalState).f6 = not((self).f25)
                r3 = (p2 if (self).f25 else p1)
                d_2747_v46_: _dafny.MultiSet
                d_2747_v46_ = _dafny.MultiSet([d_2673_v7_])
                d_2748_v47_: _dafny.Set
                d_2748_v47_ = _dafny.Set({(self).f25, True})
                d_2749_v48_: _dafny.Map
                d_2749_v48_ = _dafny.Map({(self).f25: len(d_2748_v47_)})
                d_2750_v49_: _dafny.MultiSet
                d_2750_v49_ = _dafny.MultiSet([(self).f25, default__.fm2(d_2747_v46_, p2, globalState), (self).fm22(p2, (self).f25, d_2749_v48_, globalState)])
                (globalState).f17 = ((d_2750_v49_)[(self).f25] if ((self).f25) in (d_2750_v49_) else default__.safeDivisionInt(p1, p1))
            elif True:
                (globalState).f18 = (0) - (((p2) * (p2)) - (p1))
                d_2751_v50_: _dafny.MultiSet
                d_2751_v50_ = _dafny.MultiSet([d_2673_v7_, d_2673_v7_])
                d_2752_v51_: T0
                nw423_ = C7()
                nw423_.ctor__(default__.fm2(d_2751_v50_, p1, globalState), (self).f20)
                d_2752_v51_ = nw423_
                d_2753_v52_: _dafny.Array
                nw424_ = _dafny.Array(None, 18)
                nw424_[int(0)] = d_2673_v7_
                nw424_[int(1)] = (d_2671_v5_)[default__.safeIndex(p2, len(d_2671_v5_))]
                nw424_[int(2)] = d_2673_v7_
                nw424_[int(3)] = d_2673_v7_
                nw424_[int(4)] = d_2673_v7_
                nw424_[int(5)] = d_2673_v7_
                nw424_[int(6)] = d_2673_v7_
                nw424_[int(7)] = d_2673_v7_
                nw424_[int(8)] = d_2673_v7_
                nw424_[int(9)] = (d_2671_v5_)[default__.safeIndex(p2, len(d_2671_v5_))]
                nw424_[int(10)] = d_2673_v7_
                nw424_[int(11)] = d_2673_v7_
                nw424_[int(12)] = d_2673_v7_
                nw424_[int(13)] = d_2673_v7_
                nw424_[int(14)] = d_2673_v7_
                nw424_[int(15)] = default__.fm1(globalState)
                nw424_[int(16)] = d_2673_v7_
                nw424_[int(17)] = _dafny.CodePoint('l')
                d_2753_v52_ = nw424_
                d_2754_v53_: _dafny.Map
                d_2754_v53_ = _dafny.Map({d_2752_v51_: d_2753_v52_})
                d_2755_v54_: _dafny.Seq
                d_2755_v54_ = _dafny.SeqWithoutIsStrInference([d_2753_v52_, d_2753_v52_])
                d_2754_v53_ = (d_2754_v53_).set(d_2752_v51_, (d_2755_v54_)[default__.safeIndex((_dafny.MultiSet(d_2712_v31_)).cardinality, len(d_2755_v54_))])
                d_2756_v55_: _dafny.Array
                def lambda136_(d_2757_v1_):
                    def lambda137_(d_2758_i6_):
                        return d_2757_v1_

                    return lambda137_

                init68_ = lambda136_(d_2667_v1_)
                nw425_ = _dafny.Array(None, 14)
                for i0_68_ in range(nw425_.length(0)):
                    nw425_[i0_68_] = init68_(i0_68_)
                d_2756_v55_ = nw425_
                index424_ = default__.safeIndex(454, (d_2756_v55_).length(0))
                (d_2756_v55_)[index424_] = d_2667_v1_
                index425_ = default__.safeIndex(454, (d_2756_v55_).length(0))
                (d_2756_v55_)[index425_] = ((d_2667_v1_).set(default__.safeIndex(p2, len(d_2667_v1_)), (self).f25)) + (d_2667_v1_)
                d_2759_v56_: _dafny.Map
                d_2759_v56_ = _dafny.Map({((d_2712_v31_)[default__.safeIndex(p2, len(d_2712_v31_))]) >= (p2): ((self).f25 if False else (self).f25)})
                d_2760_v57_: _dafny.MultiSet
                d_2760_v57_ = _dafny.MultiSet([(self).f25])
                d_2761_v58_: _dafny.Seq
                d_2761_v58_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f25]), d_2760_v57_])
                d_2759_v56_ = (d_2759_v56_).set(((self).f25) not in ((d_2761_v58_)[default__.safeIndex(p1, len(d_2761_v58_))]), (self).f25)
                d_2762_v59_: _dafny.Array
                def lambda138_(d_2763_i7_):
                    return (self).f25

                init69_ = lambda138_
                nw426_ = _dafny.Array(None, 9)
                for i0_69_ in range(nw426_.length(0)):
                    nw426_[i0_69_] = init69_(i0_69_)
                d_2762_v59_ = nw426_
                index426_ = default__.safeIndex(378, (d_2762_v59_).length(0))
                (d_2762_v59_)[index426_] = (self).f25
                d_2764_v60_: _dafny.Array
                def lambda139_(d_2765_v5_):
                    def lambda140_(d_2766_i8_):
                        return d_2765_v5_

                    return lambda140_

                init70_ = lambda139_(d_2671_v5_)
                nw427_ = _dafny.Array(None, 10)
                for i0_70_ in range(nw427_.length(0)):
                    nw427_[i0_70_] = init70_(i0_70_)
                d_2764_v60_ = nw427_
                index427_ = default__.safeIndex(794, (d_2764_v60_).length(0))
                (d_2764_v60_)[index427_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vxfdoy"))
                index428_ = default__.safeIndex(378, (d_2762_v59_).length(0))
                index429_ = default__.safeIndex(794, (d_2764_v60_).length(0))
                rhs402_ = (self).f25
                rhs403_ = (self).f25
                rhs404_ = (self).f25
                rhs405_ = not((d_2671_v5_) != ((d_2671_v5_).set(default__.safeIndex(p1, len(d_2671_v5_)), d_2673_v7_)))
                rhs406_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_2767_i9_ in range(default__.abs(413))])
                lhs367_ = globalState
                lhs368_ = globalState
                lhs369_ = globalState
                lhs370_ = d_2762_v59_
                lhs371_ = default__.safeIndex(378, (d_2762_v59_).length(0))
                lhs372_ = d_2764_v60_
                lhs373_ = default__.safeIndex(794, (d_2764_v60_).length(0))
                lhs367_.f5 = rhs402_
                lhs368_.f5 = rhs403_
                lhs369_.f6 = rhs404_
                lhs370_[lhs371_] = rhs405_
                lhs372_[lhs373_] = rhs406_
            if not((p2) == (p2)):
                d_2768_v61_: _dafny.Map
                d_2768_v61_ = _dafny.Map({(self).f25: p2})
                rhs407_ = ((self).fm22(p1, default__.fm2(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2673_v7_, _dafny.CodePoint('q')])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), globalState), d_2768_v61_, globalState) if (self).f25 else (self).f25)
                rhs408_ = p1
                rhs409_ = p1
                lhs374_ = globalState
                lhs374_.f6 = rhs407_
                r1 = rhs408_
                r3 = rhs409_
                d_2769_v62_: _dafny.Array
                nw428_ = _dafny.Array(False, 14)
                d_2769_v62_ = nw428_
                index430_ = default__.safeIndex(857, (d_2769_v62_).length(0))
                (d_2769_v62_)[index430_] = ((self).f25 if True else not(not(not((self).f25))))
                index431_ = default__.safeIndex(857, (d_2769_v62_).length(0))
                (d_2769_v62_)[index431_] = (self).f25
                index432_ = default__.safeIndex(857, (d_2769_v62_).length(0))
                (d_2769_v62_)[index432_] = (self).f25
                (globalState).f15 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                d_2770_v63_: _dafny.Set
                d_2770_v63_ = _dafny.Set({(self).f25})
                d_2771_v64_: _dafny.Map
                d_2771_v64_ = _dafny.Map({(d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))]: d_2669_v3_})
                d_2772_v65_: _dafny.Seq
                d_2772_v65_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([p2])])
                d_2773_v66_: _dafny.Map
                d_2773_v66_ = _dafny.Map({(((d_2771_v64_)[(d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))]] if ((d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))]) in (d_2771_v64_) else (d_2772_v65_)[default__.safeIndex(p1, len(d_2772_v65_))])).cardinality: (self).f25})
                d_2774_v68_: _dafny.Set
                def iife297_():
                    coll77_ = _dafny.Map()
                    compr_77_: int
                    for compr_77_ in (d_2674_v8_).Elements:
                        d_2775_v67_: int = compr_77_
                        if (d_2775_v67_) in (d_2674_v8_):
                            coll77_[(d_2775_v67_) - ((0) - (p1))] = (d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))]
                    return _dafny.Map(coll77_)
                d_2774_v68_ = _dafny.Set({d_2773_v66_, iife297_()
                })
                d_2776_v69_: _dafny.Array
                nw429_ = _dafny.Array(None, 17)
                nw429_[int(0)] = p1
                nw429_[int(1)] = len(_dafny.Map({p2: (d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))]}))
                nw429_[int(2)] = p1
                nw429_[int(3)] = p2
                nw429_[int(4)] = p1
                nw429_[int(5)] = len(_dafny.SeqWithoutIsStrInference([385 for d_2777_i10_ in range(default__.abs(-61))]))
                nw429_[int(6)] = len(d_2671_v5_)
                nw429_[int(7)] = p2
                nw429_[int(8)] = (0) - (len(d_2770_v63_))
                nw429_[int(9)] = len(d_2774_v68_)
                nw429_[int(10)] = p1
                nw429_[int(11)] = (_dafny.MultiSet([p1, p2])).cardinality
                nw429_[int(12)] = p1
                nw429_[int(13)] = p1
                nw429_[int(14)] = p2
                nw429_[int(15)] = p2
                nw429_[int(16)] = p1
                d_2776_v69_ = nw429_
                d_2778_v70_: _dafny.Seq
                d_2778_v70_ = _dafny.SeqWithoutIsStrInference([d_2776_v69_, d_2776_v69_])
                r3 = (0) - (len((_dafny.SeqWithoutIsStrInference([d_2776_v69_]) if (d_2769_v62_)[default__.safeIndex(857, (d_2769_v62_).length(0))] else d_2778_v70_)))
            elif True:
                d_2779_v71_: _dafny.Array
                def lambda141_(d_2780_i11_):
                    return (self).f25

                init71_ = lambda141_
                nw430_ = _dafny.Array(None, 1)
                for i0_71_ in range(nw430_.length(0)):
                    nw430_[i0_71_] = init71_(i0_71_)
                d_2779_v71_ = nw430_
                index433_ = default__.safeIndex(599, (d_2779_v71_).length(0))
                (d_2779_v71_)[index433_] = (d_2673_v7_) in (d_2671_v5_)
                index434_ = default__.safeIndex(599, (d_2779_v71_).length(0))
                (d_2779_v71_)[index434_] = ((_dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2781_i12_ in range(default__.abs(849))])) + ((_dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2782_i13_ in range(default__.abs(373))])).set(default__.safeIndex(803, len(_dafny.SeqWithoutIsStrInference([d_2673_v7_ for d_2783_i13_ in range(default__.abs(373))]))), default__.fm1(globalState)))) <= ((d_2671_v5_) + (d_2671_v5_))
                d_2784_v72_: C7
                nw431_ = C7()
                nw431_.ctor__(not((d_2779_v71_)[default__.safeIndex(599, (d_2779_v71_).length(0))]), (self).f20)
                d_2784_v72_ = nw431_
                nw432_ = _dafny.Array(None, 5)
                nw432_[int(0)] = p1
                nw432_[int(1)] = (p1) + (len(d_2671_v5_))
                nw432_[int(2)] = p2
                nw432_[int(3)] = p1
                nw432_[int(4)] = p1
                r0 = nw432_
                (globalState).f17 = p2
                d_2785_v73_: D3
                d_2785_v73_ = D3_DC7(d_2663_v0_)
                d_2786_v74_: D14
                d_2786_v74_ = D14_DC41(p1, p2, 770, d_2785_v73_)
                d_2787_v75_: D0
                d_2787_v75_ = D0_DC1(d_2673_v7_, p2, _dafny.Set({True}), not(not((self).f25)), (d_2779_v71_)[default__.safeIndex(599, (d_2779_v71_).length(0))])
                d_2788_v76_: _dafny.Map
                d_2788_v76_ = _dafny.Map({len(d_2671_v5_): (d_2779_v71_)[default__.safeIndex(599, (d_2779_v71_).length(0))]})
                d_2789_v77_: _dafny.Map
                d_2789_v77_ = _dafny.Map({d_2788_v76_: p2})
                d_2790_v78_: _dafny.Map
                d_2790_v78_ = _dafny.Map({False: (d_2784_v72_).f26})
                d_2791_v79_: _dafny.Array
                nw433_ = _dafny.Array(None, 22)
                nw433_[int(0)] = p1
                nw433_[int(1)] = (p1) - ((d_2786_v74_).cf78)
                nw433_[int(2)] = default__.safeModuloInt(default__.fm0(p2, d_2787_v75_, globalState), (0) - (p1))
                nw433_[int(3)] = default__.safeDivisionInt(len(d_2789_v77_), (d_2669_v3_).cardinality)
                nw433_[int(4)] = p1
                nw433_[int(5)] = (p2) * (default__.fm0(len(d_2788_v76_), d_2787_v75_, globalState))
                nw433_[int(6)] = ((d_2668_v2_)[(d_2669_v3_).cardinality] if ((d_2669_v3_).cardinality) in (d_2668_v2_) else p2)
                nw433_[int(7)] = len(d_2671_v5_)
                nw433_[int(8)] = (p1) - (p1)
                nw433_[int(9)] = p1
                nw433_[int(10)] = p2
                nw433_[int(11)] = ((_dafny.MultiSet(d_2712_v31_)) | (d_2669_v3_)).cardinality
                nw433_[int(12)] = 303
                nw433_[int(13)] = default__.safeModuloInt(p1, p1)
                nw433_[int(14)] = p1
                nw433_[int(15)] = p1
                nw433_[int(16)] = (0) - (p1)
                nw433_[int(17)] = default__.safeModuloInt(p1, p1)
                nw433_[int(18)] = p1
                nw433_[int(19)] = len(d_2790_v78_)
                nw433_[int(20)] = p2
                nw433_[int(21)] = 121
                d_2791_v79_ = nw433_
                index435_ = default__.safeIndex(430, (d_2791_v79_).length(0))
                (d_2791_v79_)[index435_] = p2
                index436_ = default__.safeIndex(430, (d_2791_v79_).length(0))
                (d_2791_v79_)[index436_] = ((p1) - (237)) + ((p1) + (p2))
            d_2792_v80_: _dafny.Set
            d_2792_v80_ = _dafny.Set({(self).f25, (self).f25})
            d_2793_v81_: D0
            d_2793_v81_ = D0_DC1(d_2673_v7_, p1, d_2792_v80_, (self).f25, (self).f25)
            d_2794_v82_: C2
            nw434_ = C2()
            nw434_.ctor__(default__.safeModuloInt(default__.fm0(p1, d_2793_v81_, globalState), p1))
            d_2794_v82_ = nw434_
            if not((self).f25):
                d_2795_v83_: C1
                nw435_ = C1()
                nw435_.ctor__(((d_2794_v82_).f22) - ((d_2793_v81_).cf2), D0_DC0((self).fm22((0) - (p1), (self).f25, default__.fm76((self).f25, globalState), globalState)))
                d_2795_v83_ = nw435_
                d_2712_v31_ = d_2712_v31_
                d_2796_v84_: _dafny.Array
                nw436_ = _dafny.Array(False, 8)
                d_2796_v84_ = nw436_
                d_2797_v85_: C0
                nw437_ = C0()
                nw437_.ctor__(d_2796_v84_, d_2667_v1_)
                d_2797_v85_ = nw437_
                d_2798_v86_: _dafny.Map
                d_2798_v86_ = _dafny.Map({d_2797_v85_: len(d_2712_v31_)})
                d_2799_v87_: _dafny.Array
                nw438_ = _dafny.Array(None, 11)
                nw438_[int(0)] = default__.safeModuloInt(-344, len(d_2798_v86_))
                nw438_[int(1)] = default__.safeDivisionInt((d_2794_v82_).f22, default__.fm0(len(d_2797_v85_.f24), d_2793_v81_, globalState))
                nw438_[int(2)] = (p2) - ((d_2794_v82_).f22)
                nw438_[int(3)] = default__.safeDivisionInt(469, len(_dafny.Set({d_2797_v85_.f23, d_2796_v84_})))
                nw438_[int(4)] = default__.safeModuloInt(p2, p2)
                nw438_[int(5)] = (d_2794_v82_).f22
                nw438_[int(6)] = (d_2794_v82_).f22
                nw438_[int(7)] = len(d_2712_v31_)
                nw438_[int(8)] = ((d_2794_v82_).f22) * ((d_2794_v82_).f22)
                nw438_[int(9)] = (d_2794_v82_).f22
                nw438_[int(10)] = (d_2794_v82_).f22
                d_2799_v87_ = nw438_
                index437_ = default__.safeIndex(636, (d_2799_v87_).length(0))
                (d_2799_v87_)[index437_] = p2
                index438_ = default__.safeIndex(636, (d_2799_v87_).length(0))
                (d_2799_v87_)[index438_] = ((0) - ((d_2794_v82_).f22)) - (p2)
                (globalState).f6 = (self).f25
                (globalState).f15 = (d_2671_v5_) + ((d_2671_v5_).set(default__.safeIndex(p2, len(d_2671_v5_)), _dafny.CodePoint('o')))
            elif True:
                d_2800_v88_: _dafny.MultiSet
                d_2800_v88_ = _dafny.MultiSet([not((self).f25)])
                d_2801_v89_: _dafny.MultiSet
                d_2801_v89_ = _dafny.MultiSet([not((self).f25), (self).f25, (self).f25, (self).f25, (self).f25])
                (globalState).f6 = ((d_2801_v89_)).issubset((_dafny.MultiSet(d_2667_v1_) if (self).f25 else d_2800_v88_))
                d_2802_v90_: C6
                nw439_ = C6()
                nw439_.ctor__((self).f20)
                d_2802_v90_ = nw439_
                (globalState).f18 = (0) - ((d_2794_v82_).f22)
                (globalState).f17 = ((d_2669_v3_).intersection(default__.fm56((0) - ((d_2794_v82_).f22), (self).f25, not(not((self).f25)), p1, globalState))).cardinality
                r3 = (d_2794_v82_).f22
        (globalState).f6 = (self).f25
        if True:
            d_2803_v91_: _dafny.MultiSet
            d_2803_v91_ = _dafny.MultiSet([d_2673_v7_, d_2673_v7_, d_2673_v7_])
            d_2804_v92_: C7
            nw440_ = C7()
            nw440_.ctor__(default__.fm2(d_2803_v91_, (_dafny.MultiSet([len((d_2671_v5_).set(default__.safeIndex(p2, len(d_2671_v5_)), d_2673_v7_))])).cardinality, globalState), (self).f20)
            d_2804_v92_ = nw440_
            (globalState).f6 = (self).f25
            d_2805_v93_: _dafny.Set
            d_2805_v93_ = _dafny.Set({False})
            d_2806_v94_: _dafny.Set
            d_2806_v94_ = _dafny.Set({d_2805_v93_, d_2805_v93_, d_2805_v93_})
            d_2807_v95_: D6
            d_2807_v95_ = D6_DC17(_dafny.MultiSet([p2, p1, 509]), d_2806_v94_)
            d_2669_v3_ = (d_2669_v3_ if (d_2804_v92_).f26 else (_dafny.MultiSet([p1])).intersection((d_2807_v95_).cf28))
            (globalState).f5 = (d_2674_v8_).ispropersubset(d_2674_v8_)
            d_2808_v96_: C2
            nw441_ = C2()
            nw441_.ctor__((p1) * (p2))
            d_2808_v96_ = nw441_
        elif True:
            d_2673_v7_ = (d_2673_v7_ if True else d_2673_v7_)
            d_2671_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdw"))
            if True:
                d_2809_v97_: _dafny.Array
                nw442_ = _dafny.Array(int(0), 3)
                d_2809_v97_ = nw442_
                d_2810_v98_: D7
                d_2810_v98_ = D7_DC19(d_2809_v97_)
                d_2810_v98_ = d_2810_v98_
                d_2811_v99_: C3
                nw443_ = C3()
                nw443_.ctor__((self).f20)
                d_2811_v99_ = nw443_
                d_2812_v100_: D3
                d_2812_v100_ = D3_DC8((self).f25)
                d_2813_v101_: _dafny.Map
                d_2813_v101_ = _dafny.Map({(p1) - (p2): (d_2812_v100_).cf16})
                (globalState).f5 = ((d_2813_v101_)[p2] if (p2) in (d_2813_v101_) else (self).f25)
                d_2814_v102_: _dafny.Set
                d_2814_v102_ = _dafny.Set({(self).f25, (self).f25})
                d_2815_v103_: _dafny.Map
                d_2815_v103_ = _dafny.Map({len(d_2671_v5_): d_2814_v102_})
                (globalState).f6 = ((d_2814_v102_) - (((d_2815_v103_)[len(d_2712_v31_)] if (len(d_2712_v31_)) in (d_2815_v103_) else d_2814_v102_))).issubset((d_2814_v102_ if (self).f25 else d_2814_v102_))
                d_2816_v104_: _dafny.Map
                d_2816_v104_ = _dafny.Map({(self).f25: len(d_2712_v31_)})
                d_2817_v105_: _dafny.MultiSet
                d_2817_v105_ = _dafny.MultiSet([(self).f25, (self).f25, (self).fm22(len(d_2671_v5_), (self).f25, d_2816_v104_, globalState)])
                d_2818_v106_: _dafny.Seq
                d_2818_v106_ = _dafny.SeqWithoutIsStrInference([d_2817_v105_])
                d_2817_v105_ = ((d_2817_v105_).intersection((d_2818_v106_)[default__.safeIndex(p2, len(d_2818_v106_))])).intersection(default__.fm51(globalState))
            elif True:
                (globalState).f15 = d_2671_v5_
                d_2819_v108_: _dafny.Map
                d_2819_v108_ = _dafny.Map({(self).f25: d_2673_v7_})
                d_2820_v109_: _dafny.Seq
                d_2820_v109_ = _dafny.SeqWithoutIsStrInference([d_2819_v108_])
                def iife298_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in _dafny.IntegerRange(-846, 117):
                        d_2821_v107_: int = compr_78_
                        if ((-846) <= (d_2821_v107_)) and ((d_2821_v107_) < (117)):
                            coll78_ = coll78_.union(_dafny.Set([(d_2821_v107_) - (582)]))
                    return _dafny.Set(coll78_)
                (globalState).f16 = (default__.fm27((0) - (p2), d_2669_v3_, p2, _dafny.Map({iife298_()
                : len(d_2820_v109_)}), globalState)) + (d_2671_v5_)
                (globalState).f6 = ((self).f25) and ((self).f25)
                (globalState).f18 = (d_2712_v31_)[default__.safeIndex(-93, len(d_2712_v31_))]
                d_2822_v110_: _dafny.Array
                nw444_ = _dafny.Array(D8.default()(), 23)
                d_2822_v110_ = nw444_
                d_2823_v111_: D8
                d_2823_v111_ = D8_DC23(d_2671_v5_)
                index439_ = default__.safeIndex(285, (d_2822_v110_).length(0))
                (d_2822_v110_)[index439_] = d_2823_v111_
                index440_ = default__.safeIndex(285, (d_2822_v110_).length(0))
                rhs410_ = d_2823_v111_
                rhs411_ = (self).f25
                lhs375_ = d_2822_v110_
                lhs376_ = default__.safeIndex(285, (d_2822_v110_).length(0))
                lhs377_ = globalState
                lhs375_[lhs376_] = rhs410_
                lhs377_.f6 = rhs411_
            r2 = len(d_2671_v5_)
            (globalState).f18 = p2
        d_2824_v112_: _dafny.Array
        def lambda142_(d_2825_p1_):
            def lambda143_(d_2826_i14_):
                return (d_2826_i14_) + ((0) - (len(_dafny.Map({(self).f25: d_2825_p1_}))))

            return lambda143_

        init72_ = lambda142_(p1)
        nw445_ = _dafny.Array(None, 2)
        for i0_72_ in range(nw445_.length(0)):
            nw445_[i0_72_] = init72_(i0_72_)
        d_2824_v112_ = nw445_
        r0 = (d_2824_v112_ if (self).f25 else d_2824_v112_)
        d_2827_v113_: _dafny.Map
        d_2827_v113_ = _dafny.Map({(self).f25: d_2673_v7_})
        d_2828_v114_: _dafny.Map
        d_2828_v114_ = _dafny.Map({((d_2827_v113_)[not((self).f25)] if (not((self).f25)) in (d_2827_v113_) else d_2673_v7_): p1})
        d_2829_v115_: _dafny.Map
        d_2829_v115_ = _dafny.Map({not((self).f25): p2})
        r1 = ((((d_2828_v114_)[default__.fm1(globalState)] if (default__.fm1(globalState)) in (d_2828_v114_) else p2) if (D2_DC4(p1, (self).f25, (self).f25, (self).f25)).cf11 else (0) - (p2))) * (((d_2829_v115_)[(self).f25] if ((self).f25) in (d_2829_v115_) else p2))
        r2 = ((0) - (p2)) + (799)
        r3 = 997
        return r0, r1, r2, r3

    @property
    def f25(self):
        return self._f25

class C9(T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f20(self):
        return self._f20
    def ctor__(self, f20):
        (self)._f20 = f20

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        (globalState).f6 = p0
        (globalState).f5 = False
        d_2830_v0_: D2
        d_2830_v0_ = D2_DC4(146, p0, p0, False)
        d_2831_v1_: str
        d_2831_v1_ = _dafny.CodePoint('n')
        d_2832_v2_: _dafny.Map
        d_2832_v2_ = _dafny.Map({d_2831_v1_: p0})
        d_2833_v3_: _dafny.Array
        def lambda144_(d_2834_p0_):
            def lambda145_(d_2835_i0_):
                return (d_2834_p0_ if d_2834_p0_ else d_2834_p0_)

            return lambda145_

        init73_ = lambda144_(p0)
        nw446_ = _dafny.Array(None, 7)
        for i0_73_ in range(nw446_.length(0)):
            nw446_[i0_73_] = init73_(i0_73_)
        d_2833_v3_ = nw446_
        rhs412_ = (d_2830_v0_).cf11
        rhs413_ = ((p1) * (len(d_2832_v2_))) > (p1)
        rhs414_ = p1
        rhs415_ = d_2833_v3_
        lhs378_ = globalState
        lhs379_ = globalState
        lhs380_ = globalState
        lhs381_ = globalState
        lhs378_.f6 = rhs412_
        lhs379_.f6 = rhs413_
        lhs380_.f18 = rhs414_
        lhs381_.f4 = rhs415_
        d_2836_v4_: _dafny.Seq
        d_2836_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhowjnjs"))
        d_2837_v5_: D2
        d_2837_v5_ = D2_DC3(d_2836_v4_)
        source37_ = d_2837_v5_
        if source37_.is_DC4:
            d_2838___mcc_h0_ = source37_.cf8
            d_2839___mcc_h1_ = source37_.cf9
            d_2840___mcc_h2_ = source37_.cf10
            d_2841___mcc_h3_ = source37_.cf11
            d_2842_cf11_ = d_2841___mcc_h3_
            d_2843_cf10_ = d_2840___mcc_h2_
            d_2844_cf9_ = d_2839___mcc_h1_
            d_2845_cf8_ = d_2838___mcc_h0_
            d_2845_cf8_ = (d_2845_cf8_) * (p1)
            (globalState).f15 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xl"))) + (d_2836_v4_)
            (globalState).f16 = (d_2836_v4_) + (d_2836_v4_)
            d_2846_v6_: _dafny.Seq
            d_2846_v6_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(535, -334), p1])
            d_2846_v6_ = (_dafny.SeqWithoutIsStrInference([815, 458, p1, p1, d_2845_cf8_])) + (d_2846_v6_)
        elif source37_.is_DC5:
            d_2847___mcc_h4_ = source37_.cf12
            d_2848___mcc_h5_ = source37_.cf13
            d_2849_cf13_ = d_2848___mcc_h5_
            d_2850_cf12_ = d_2847___mcc_h4_
            d_2851_v7_: _dafny.Array
            nw447_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_2851_v7_ = nw447_
            d_2851_v7_ = d_2851_v7_
            d_2852_v8_: _dafny.MultiSet
            d_2852_v8_ = _dafny.MultiSet([_dafny.CodePoint('g')])
            d_2853_v9_: _dafny.Array
            def lambda146_(d_2854_p1_):
                def lambda147_(d_2855_i1_):
                    return default__.safeModuloInt(d_2855_i1_, d_2854_p1_)

                return lambda147_

            init74_ = lambda146_(p1)
            nw448_ = _dafny.Array(None, 19)
            for i0_74_ in range(nw448_.length(0)):
                nw448_[i0_74_] = init74_(i0_74_)
            d_2853_v9_ = nw448_
            pat_let_tv108_ = d_2852_v8_
            pat_let_tv109_ = globalState
            d_2856_v10_: _dafny.Map
            def iife299_(_pat_let110_0):
                def iife300_(d_2857_dt__update__tmp_h0_):
                    def iife301_(_pat_let111_0):
                        def iife302_(d_2858_dt__update_hcf11_h0_):
                            return D2_DC4((d_2857_dt__update__tmp_h0_).cf8, (d_2857_dt__update__tmp_h0_).cf9, (d_2857_dt__update__tmp_h0_).cf10, d_2858_dt__update_hcf11_h0_)
                        return iife302_(_pat_let111_0)
                    return iife301_(default__.fm2(pat_let_tv108_, 23, pat_let_tv109_))
                return iife300_(_pat_let110_0)
            d_2856_v10_ = _dafny.Map({iife299_(d_2830_v0_): d_2853_v9_})
            d_2856_v10_ = (d_2856_v10_) | ((d_2856_v10_) | (d_2856_v10_))
            d_2859_v11_: _dafny.Map
            d_2859_v11_ = _dafny.Map({d_2831_v1_: d_2836_v4_})
            (globalState).f16 = (d_2836_v4_) + (((d_2859_v11_)[d_2831_v1_] if (d_2831_v1_) in (d_2859_v11_) else d_2836_v4_))
            if p0:
                d_2860_v12_: _dafny.MultiSet
                d_2860_v12_ = _dafny.MultiSet([p0])
                r0 = (default__.fm21(globalState)).intersection(d_2860_v12_)
                index441_ = default__.safeIndex(840, (d_2853_v9_).length(0))
                (d_2853_v9_)[index441_] = p1
                d_2861_v13_: _dafny.Set
                d_2861_v13_ = _dafny.Set({p1})
                d_2862_v14_: _dafny.Seq
                d_2862_v14_ = _dafny.SeqWithoutIsStrInference([d_2850_cf12_, d_2850_cf12_, False])
                d_2863_v15_: _dafny.Map
                d_2863_v15_ = _dafny.Map({p1: d_2831_v1_})
                index442_ = default__.safeIndex(840, (d_2853_v9_).length(0))
                rhs416_ = (0) - (len(d_2861_v13_))
                rhs417_ = default__.safeModuloInt(len(d_2862_v14_), len(d_2863_v15_))
                rhs418_ = p0
                lhs382_ = d_2853_v9_
                lhs383_ = default__.safeIndex(840, (d_2853_v9_).length(0))
                lhs384_ = globalState
                lhs382_[lhs383_] = rhs416_
                lhs384_.f17 = rhs417_
                d_2850_cf12_ = rhs418_
                d_2864_v16_: C0
                nw449_ = C0()
                nw449_.ctor__(d_2833_v3_, d_2862_v14_)
                d_2864_v16_ = nw449_
                d_2865_v17_: _dafny.Map
                d_2865_v17_ = _dafny.Map({not(d_2850_cf12_): d_2864_v16_})
                d_2866_v18_: _dafny.Map
                d_2866_v18_ = _dafny.Map({d_2833_v3_: d_2864_v16_})
                d_2867_v19_: _dafny.Seq
                d_2867_v19_ = _dafny.SeqWithoutIsStrInference([d_2864_v16_, d_2864_v16_, d_2864_v16_, d_2864_v16_, ((d_2866_v18_)[d_2864_v16_.f23] if (d_2864_v16_.f23) in (d_2866_v18_) else d_2864_v16_)])
                d_2868_v20_: T0
                nw450_ = C8()
                nw450_.ctor__(d_2850_cf12_, (self).f20)
                d_2868_v20_ = nw450_
                d_2869_v21_: _dafny.Seq
                d_2869_v21_ = _dafny.SeqWithoutIsStrInference([(0) - (p1), len(_dafny.SeqWithoutIsStrInference([d_2868_v20_, d_2868_v20_])), (0) - ((d_2853_v9_)[default__.safeIndex(840, (d_2853_v9_).length(0))]), len(d_2861_v13_)])
                d_2870_v22_: _dafny.Map
                d_2870_v22_ = _dafny.Map({d_2850_cf12_: d_2869_v21_})
                d_2871_v23_: _dafny.MultiSet
                d_2871_v23_ = _dafny.MultiSet([((d_2870_v22_)[p0] if (p0) in (d_2870_v22_) else _dafny.SeqWithoutIsStrInference([len(d_2836_v4_), p1, len(_dafny.SeqWithoutIsStrInference([p1, (d_2853_v9_)[default__.safeIndex(840, (d_2853_v9_).length(0))], d_2849_cf13_]))])), _dafny.SeqWithoutIsStrInference([d_2849_cf13_ for d_2872_i2_ in range(default__.abs(53))]), d_2869_v21_, d_2869_v21_, _dafny.SeqWithoutIsStrInference([(d_2853_v9_)[default__.safeIndex(840, (d_2853_v9_).length(0))], (d_2853_v9_)[default__.safeIndex(840, (d_2853_v9_).length(0))]])])
                d_2865_v17_ = (d_2865_v17_).set(p0, (d_2867_v19_)[default__.safeIndex((d_2871_v23_).cardinality, len(d_2867_v19_))])
                d_2873_v24_: _dafny.Map
                d_2873_v24_ = _dafny.Map({p0: d_2849_cf13_})
                d_2873_v24_ = (d_2873_v24_).set(d_2850_cf12_, default__.safeDivisionInt(p1, (d_2853_v9_)[default__.safeIndex(840, (d_2853_v9_).length(0))]))
                d_2874_v25_: _dafny.Map
                d_2874_v25_ = _dafny.Map({d_2849_cf13_: default__.fm46(d_2849_cf13_, p0, d_2850_cf12_, globalState)})
                d_2875_v27_: _dafny.Set
                d_2875_v27_ = _dafny.Set({d_2864_v16_.f23, d_2833_v3_})
                def iife303_():
                    coll79_ = _dafny.Map()
                    compr_79_: int
                    for compr_79_ in (d_2861_v13_).Elements:
                        d_2876_v26_: int = compr_79_
                        if (d_2876_v26_) in (d_2861_v13_):
                            coll79_[(d_2876_v26_) * (p1)] = d_2836_v4_
                    return _dafny.Map(coll79_)
                d_2874_v25_ = (iife303_()
                ).set(len(d_2875_v27_), d_2836_v4_)
            elif True:
                d_2877_v28_: _dafny.Seq
                d_2877_v28_ = _dafny.SeqWithoutIsStrInference([d_2853_v9_, d_2853_v9_, d_2853_v9_, d_2853_v9_])
                d_2878_v29_: _dafny.Seq
                d_2878_v29_ = _dafny.SeqWithoutIsStrInference([d_2877_v28_, d_2877_v28_, (_dafny.SeqWithoutIsStrInference([d_2853_v9_, d_2853_v9_])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2831_v1_ for d_2879_i3_ in range(default__.abs(937))])), len(_dafny.SeqWithoutIsStrInference([d_2853_v9_, d_2853_v9_]))), d_2853_v9_), d_2877_v28_])
                d_2880_v30_: _dafny.MultiSet
                d_2880_v30_ = _dafny.MultiSet([d_2849_cf13_, d_2849_cf13_])
                d_2881_v31_: _dafny.Map
                d_2881_v31_ = _dafny.Map({d_2850_cf12_: (d_2878_v29_)[default__.safeIndex((d_2880_v30_).cardinality, len(d_2878_v29_))]})
                d_2882_v32_: _dafny.Seq
                d_2882_v32_ = _dafny.SeqWithoutIsStrInference([d_2849_cf13_])
                d_2883_v33_: D7
                d_2883_v33_ = D7_DC21(len(d_2882_v32_), p0, d_2850_cf12_, d_2853_v9_, p0)
                pat_let_tv110_ = d_2850_cf12_
                pat_let_tv111_ = p0
                d_2884_v34_: _dafny.Array
                nw451_ = _dafny.Array(None, 22)
                nw451_[int(0)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2853_v9_, d_2853_v9_])
                nw451_[int(2)] = (d_2877_v28_) + (_dafny.SeqWithoutIsStrInference([d_2853_v9_]))
                nw451_[int(3)] = ((d_2881_v31_)[d_2850_cf12_] if (d_2850_cf12_) in (d_2881_v31_) else d_2877_v28_)
                nw451_[int(4)] = d_2877_v28_
                nw451_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_2853_v9_]) if p0 else d_2877_v28_)
                nw451_[int(6)] = d_2877_v28_
                nw451_[int(7)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(8)] = d_2877_v28_
                nw451_[int(9)] = ((d_2877_v28_) + (d_2877_v28_)).set(default__.safeIndex(d_2849_cf13_, len((d_2877_v28_) + (d_2877_v28_))), d_2853_v9_)
                nw451_[int(10)] = d_2877_v28_
                nw451_[int(11)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(12)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(13)] = d_2877_v28_
                nw451_[int(14)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(15)] = d_2877_v28_
                nw451_[int(16)] = (d_2877_v28_) + (d_2877_v28_)
                nw451_[int(17)] = _dafny.SeqWithoutIsStrInference([d_2853_v9_])
                nw451_[int(18)] = d_2877_v28_
                nw451_[int(19)] = d_2877_v28_
                def iife304_(_pat_let112_0):
                    def iife305_(d_2885_dt__update__tmp_h1_):
                        def iife306_(_pat_let113_0):
                            def iife307_(d_2886_dt__update_hcf39_h0_):
                                def iife308_(_pat_let114_0):
                                    def iife309_(d_2887_dt__update_hcf40_h0_):
                                        return D7_DC21((d_2885_dt__update__tmp_h1_).cf38, d_2886_dt__update_hcf39_h0_, d_2887_dt__update_hcf40_h0_, (d_2885_dt__update__tmp_h1_).cf41, (d_2885_dt__update__tmp_h1_).cf42)
                                    return iife309_(_pat_let114_0)
                                return iife308_(pat_let_tv111_)
                            return iife307_(_pat_let113_0)
                        return iife306_(pat_let_tv110_)
                    return iife305_(_pat_let112_0)
                nw451_[int(20)] = _dafny.SeqWithoutIsStrInference([d_2853_v9_, d_2853_v9_, (iife304_(d_2883_v33_)).cf41, d_2853_v9_, d_2853_v9_])
                nw451_[int(21)] = d_2877_v28_
                d_2884_v34_ = nw451_
                index443_ = default__.safeIndex(109, (d_2884_v34_).length(0))
                (d_2884_v34_)[index443_] = d_2877_v28_
                index444_ = default__.safeIndex(109, (d_2884_v34_).length(0))
                (d_2884_v34_)[index444_] = (d_2877_v28_) + (d_2877_v28_)
                d_2888_v35_: _dafny.Map
                d_2888_v35_ = _dafny.Map({d_2850_cf12_: d_2853_v9_})
                d_2889_v36_: _dafny.Seq
                d_2889_v36_ = _dafny.SeqWithoutIsStrInference([d_2850_cf12_])
                d_2890_v37_: _dafny.MultiSet
                d_2890_v37_ = _dafny.MultiSet([p0])
                d_2891_v38_: _dafny.MultiSet
                d_2891_v38_ = d_2890_v37_
                d_2892_v39_: _dafny.Map
                d_2892_v39_ = _dafny.Map({d_2850_cf12_: False})
                d_2893_v40_: _dafny.Array
                nw452_ = _dafny.Array(None, 23)
                nw452_[int(0)] = d_2889_v36_
                nw452_[int(1)] = _dafny.SeqWithoutIsStrInference([d_2850_cf12_])
                nw452_[int(2)] = d_2889_v36_
                nw452_[int(3)] = d_2889_v36_
                nw452_[int(4)] = d_2889_v36_
                nw452_[int(5)] = d_2889_v36_
                nw452_[int(6)] = d_2889_v36_
                nw452_[int(7)] = d_2889_v36_
                nw452_[int(8)] = _dafny.SeqWithoutIsStrInference([d_2850_cf12_, default__.fm2(d_2852_v8_, (0) - (-566), globalState), True, p0])
                nw452_[int(9)] = d_2889_v36_
                nw452_[int(10)] = d_2889_v36_
                nw452_[int(11)] = d_2889_v36_
                nw452_[int(12)] = d_2889_v36_
                nw452_[int(13)] = d_2889_v36_
                nw452_[int(14)] = d_2889_v36_
                nw452_[int(15)] = d_2889_v36_
                nw452_[int(16)] = d_2889_v36_
                nw452_[int(17)] = d_2889_v36_
                nw452_[int(18)] = d_2889_v36_
                nw452_[int(19)] = default__.fm50(d_2850_cf12_, d_2850_cf12_, d_2891_v38_, globalState)
                nw452_[int(20)] = (d_2889_v36_).set(default__.safeIndex(len(d_2892_v39_), len(d_2889_v36_)), d_2850_cf12_)
                nw452_[int(21)] = d_2889_v36_
                nw452_[int(22)] = d_2889_v36_
                d_2893_v40_ = nw452_
                d_2894_v41_: D3
                d_2894_v41_ = D3_DC8(d_2850_cf12_)
                d_2895_v42_: D15
                d_2895_v42_ = D15_DC43(d_2893_v40_, p0, p1, d_2850_cf12_, d_2894_v41_)
                d_2896_v43_: _dafny.Seq
                d_2896_v43_ = _dafny.SeqWithoutIsStrInference([d_2888_v35_, d_2888_v35_, (d_2888_v35_).set((d_2895_v42_).cf83, d_2853_v9_)])
                d_2897_v44_: _dafny.Map
                d_2897_v44_ = _dafny.Map({(d_2896_v43_)[default__.safeIndex(d_2849_cf13_, len(d_2896_v43_))]: d_2850_cf12_})
                (globalState).f6 = not(((d_2897_v44_)[_dafny.Map({d_2850_cf12_: d_2853_v9_})] if (_dafny.Map({d_2850_cf12_: d_2853_v9_})) in (d_2897_v44_) else True))
                d_2849_cf13_ = len(d_2836_v4_)
                (globalState).f18 = (p1) + (d_2849_cf13_)
                (globalState).f6 = p0
        elif source37_.is_DC3:
            d_2898___mcc_h6_ = source37_.cf7
            d_2899_cf7_ = d_2898___mcc_h6_
            d_2900_v45_: _dafny.Seq
            d_2900_v45_ = _dafny.SeqWithoutIsStrInference([p1, (0) - (p1)])
            d_2901_v46_: _dafny.Seq
            d_2901_v46_ = _dafny.SeqWithoutIsStrInference([d_2900_v45_, d_2900_v45_, d_2900_v45_, _dafny.SeqWithoutIsStrInference([p1, p1]), d_2900_v45_])
            if not((d_2900_v45_) not in (d_2901_v46_)):
                d_2902_v47_: _dafny.Array
                nw453_ = _dafny.Array(None, 3)
                nw453_[int(0)] = d_2833_v3_
                nw453_[int(1)] = d_2833_v3_
                nw453_[int(2)] = d_2833_v3_
                d_2902_v47_ = nw453_
                index445_ = default__.safeIndex(35, (d_2902_v47_).length(0))
                (d_2902_v47_)[index445_] = d_2833_v3_
                index446_ = default__.safeIndex(35, (d_2902_v47_).length(0))
                nw454_ = _dafny.Array(False, 9)
                (d_2902_v47_)[index446_] = nw454_
                (globalState).f17 = p1
                (globalState).f18 = p1
                d_2903_v48_: _dafny.Set
                d_2903_v48_ = _dafny.Set({(len(d_2899_cf7_)) < (-542), p0, p0, False, (not(not(p0))) or (p0)})
                d_2903_v48_ = (_dafny.Set({p0})) | ((d_2903_v48_) - (_dafny.Set({p0, p0})))
                d_2904_v49_: _dafny.Array
                nw455_ = _dafny.Array(_dafny.Seq({}), 13)
                d_2904_v49_ = nw455_
                d_2905_v50_: D0
                d_2905_v50_ = D0_DC1(d_2831_v1_, len(_dafny.SeqWithoutIsStrInference([p0, p0, p0])), _dafny.Set({p0, p0, p0, p0, False}), p0, p0)
                d_2906_v51_: _dafny.MultiSet
                d_2906_v51_ = _dafny.MultiSet([d_2831_v1_])
                d_2907_v52_: D20
                d_2907_v52_ = D20_DC54(len(d_2899_cf7_), d_2904_v49_, default__.fm0(len(d_2836_v4_), d_2905_v50_, globalState), not(default__.fm2(d_2906_v51_, p1, globalState)), p1)
                d_2908_v53_: D20
                d_2908_v53_ = D20_DC54(p1, (d_2907_v52_).cf101, p1, p0, (p1) + (p1))
                index447_ = default__.safeIndex(385, (d_2833_v3_).length(0))
                (d_2833_v3_)[index447_] = p0
                d_2909_v54_: D4
                d_2909_v54_ = D4_DC11(p0, p0, p1, p1)
                d_2910_v55_: D4
                d_2910_v55_ = D4_DC12(d_2909_v54_)
                d_2911_v56_: D4
                d_2911_v56_ = D4_DC12(d_2910_v55_)
                d_2912_v57_: _dafny.Set
                d_2912_v57_ = _dafny.Set({d_2836_v4_})
                index448_ = default__.safeIndex(385, (d_2833_v3_).length(0))
                rhs419_ = p1
                rhs420_ = d_2908_v53_
                rhs421_ = True
                rhs422_ = (d_2912_v57_).ispropersubset(_dafny.Set({d_2836_v4_}))
                rhs423_ = d_2911_v56_
                lhs385_ = globalState
                lhs386_ = globalState
                lhs387_ = d_2833_v3_
                lhs388_ = default__.safeIndex(385, (d_2833_v3_).length(0))
                lhs385_.f18 = rhs419_
                d_2907_v52_ = rhs420_
                lhs386_.f6 = rhs421_
                lhs387_[lhs388_] = rhs422_
                d_2911_v56_ = rhs423_
            elif True:
                index449_ = default__.safeIndex(364, (d_2833_v3_).length(0))
                (d_2833_v3_)[index449_] = (d_2899_cf7_) < (d_2899_cf7_)
                index450_ = default__.safeIndex(364, (d_2833_v3_).length(0))
                (d_2833_v3_)[index450_] = ((0) - (p1)) != (len(d_2900_v45_))
                d_2913_v58_: _dafny.Array
                def lambda148_(d_2914_cf7_):
                    def lambda149_(d_2915_i4_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwgw"))) + (d_2914_cf7_)

                    return lambda149_

                init75_ = lambda148_(d_2899_cf7_)
                nw456_ = _dafny.Array(None, 7)
                for i0_75_ in range(nw456_.length(0)):
                    nw456_[i0_75_] = init75_(i0_75_)
                d_2913_v58_ = nw456_
                rhs424_ = d_2913_v58_
                rhs425_ = default__.safeDivisionInt(p1, (0) - (p1))
                lhs389_ = globalState
                d_2913_v58_ = rhs424_
                lhs389_.f17 = rhs425_
                (globalState).f5 = (d_2833_v3_)[default__.safeIndex(364, (d_2833_v3_).length(0))]
                (globalState).f6 = ((0) - (p1)) < (p1)
                (globalState).f18 = p1
            if False:
                d_2916_v59_: _dafny.Seq
                d_2916_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0])])
                d_2917_v60_: _dafny.Map
                d_2917_v60_ = _dafny.Map({p0: len((d_2916_v59_)[default__.safeIndex(p1, len(d_2916_v59_))])})
                d_2918_v61_: _dafny.Seq
                d_2918_v61_ = _dafny.SeqWithoutIsStrInference([d_2917_v60_])
                d_2918_v61_ = default__.fm77(p0, p1, globalState)
                d_2919_v62_: _dafny.MultiSet
                d_2919_v62_ = _dafny.MultiSet([p0, p0])
                d_2920_v63_: _dafny.Map
                d_2920_v63_ = _dafny.Map({(_dafny.MultiSet([p0, p0])).issubset(d_2919_v62_): default__.fm1(globalState)})
                d_2920_v63_ = (d_2920_v63_).set(p0, d_2831_v1_)
                d_2921_v64_: _dafny.Map
                d_2921_v64_ = _dafny.Map({d_2831_v1_: default__.safeDivisionInt(p1, p1)})
                d_2921_v64_ = d_2921_v64_
                (globalState).f17 = default__.safeDivisionInt(default__.safeModuloInt(170, p1), len(d_2920_v63_))
                d_2922_v65_: _dafny.MultiSet
                d_2922_v65_ = _dafny.MultiSet([d_2831_v1_])
                d_2923_v66_: D21
                d_2923_v66_ = D21_DC56(d_2922_v65_)
                (globalState).f6 = ((d_2923_v66_).cf106).issubset((d_2923_v66_).cf106)
            elif True:
                d_2924_v67_: _dafny.Array
                nw457_ = _dafny.Array(_dafny.Map({}), 9)
                d_2924_v67_ = nw457_
                d_2925_v68_: _dafny.Map
                d_2925_v68_ = _dafny.Map({d_2833_v3_: d_2831_v1_})
                index451_ = default__.safeIndex(577, (d_2924_v67_).length(0))
                (d_2924_v67_)[index451_] = (d_2925_v68_) | (_dafny.Map({d_2833_v3_: d_2831_v1_}))
                index452_ = default__.safeIndex(577, (d_2924_v67_).length(0))
                (d_2924_v67_)[index452_] = ((d_2925_v68_ if p0 else d_2925_v68_)).set(d_2833_v3_, d_2831_v1_)
                (globalState).f6 = (p0 if p0 else p0)
                d_2926_v69_: _dafny.Map
                d_2926_v69_ = _dafny.Map({p0: d_2831_v1_})
                d_2926_v69_ = (d_2926_v69_) | (_dafny.Map({not(p0): d_2831_v1_}))
                d_2927_v70_: _dafny.Map
                d_2927_v70_ = _dafny.Map({(d_2836_v4_) + (d_2836_v4_): p0})
                d_2927_v70_ = (d_2927_v70_).set((d_2899_cf7_) + (d_2836_v4_), p0)
                d_2928_v71_: D14
                d_2928_v71_ = D14_DC40(p0, (d_2899_cf7_) + (d_2899_cf7_), d_2899_cf7_, p0)
                d_2928_v71_ = D14_DC40((True) not in (_dafny.SeqWithoutIsStrInference([p0])), d_2836_v4_, d_2836_v4_, not(p0))
            index453_ = default__.safeIndex(280, (d_2833_v3_).length(0))
            (d_2833_v3_)[index453_] = (69) > (766)
            index454_ = default__.safeIndex(280, (d_2833_v3_).length(0))
            (d_2833_v3_)[index454_] = p0
            index455_ = default__.safeIndex(280, (d_2833_v3_).length(0))
            (d_2833_v3_)[index455_] = (d_2833_v3_)[default__.safeIndex(280, (d_2833_v3_).length(0))]
        elif True:
            d_2929___mcc_h7_ = source37_.cf14
            d_2930_cf14_ = d_2929___mcc_h7_
            (globalState).f6 = p0
            d_2931_v72_: _dafny.Seq
            d_2931_v72_ = _dafny.SeqWithoutIsStrInference([False])
            r1 = ((0) - (len((_dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])) + (d_2931_v72_)))) + (p1)
            index456_ = default__.safeIndex(286, (d_2833_v3_).length(0))
            (d_2833_v3_)[index456_] = p0
            d_2932_v73_: _dafny.MultiSet
            d_2932_v73_ = _dafny.MultiSet([p0])
            index457_ = default__.safeIndex(286, (d_2833_v3_).length(0))
            (d_2833_v3_)[index457_] = ((d_2932_v73_).set(p0, default__.abs(520))).ispropersubset(d_2932_v73_)
            d_2933_v74_: _dafny.Array
            nw458_ = _dafny.Array(int(0), 6)
            d_2933_v74_ = nw458_
            d_2934_v75_: D0
            d_2934_v75_ = D0_DC1(d_2831_v1_, p1, default__.fm52(len(d_2836_v4_), _dafny.SeqWithoutIsStrInference([d_2831_v1_ for d_2935_i5_ in range(default__.abs(-931))]), p0, False, globalState), p0, (d_2833_v3_)[default__.safeIndex(286, (d_2833_v3_).length(0))])
            d_2936_v76_: C6
            nw459_ = C6()
            nw459_.ctor__((self).f20)
            d_2936_v76_ = nw459_
            d_2937_v77_: _dafny.Map
            d_2937_v77_ = _dafny.Map({default__.fm0(p1, d_2934_v75_, globalState): d_2936_v76_})
            d_2938_v78_: _dafny.Seq
            d_2938_v78_ = _dafny.SeqWithoutIsStrInference([p1, (0) - (len(d_2937_v77_)), p1])
            d_2939_v79_: _dafny.Seq
            d_2939_v79_ = _dafny.SeqWithoutIsStrInference([d_2931_v72_])
            d_2940_v80_: _dafny.Map
            d_2940_v80_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxrp")): len(d_2938_v78_)})
            d_2941_v81_: _dafny.Array
            nw460_ = _dafny.Array(None, 27)
            d_2941_v81_ = nw460_
            nw461_ = _dafny.Array(None, 12)
            nw461_[int(0)] = p1
            nw461_[int(1)] = p1
            nw461_[int(2)] = p1
            nw461_[int(3)] = len(d_2938_v78_)
            nw461_[int(4)] = (p1) + (p1)
            nw461_[int(5)] = (p1) * (p1)
            nw461_[int(6)] = p1
            nw461_[int(7)] = p1
            nw461_[int(8)] = default__.safeDivisionInt(len((d_2939_v79_)[default__.safeIndex(len(d_2940_v80_), len(d_2939_v79_))]), len(d_2836_v4_))
            nw461_[int(9)] = len(_dafny.Map({(d_2833_v3_)[default__.safeIndex(286, (d_2833_v3_).length(0))]: d_2941_v81_}))
            nw461_[int(10)] = (p1 if p0 else p1)
            nw461_[int(11)] = p1
            d_2933_v74_ = nw461_
        index458_ = default__.safeIndex(112, (d_2833_v3_).length(0))
        (d_2833_v3_)[index458_] = p0
        index459_ = default__.safeIndex(112, (d_2833_v3_).length(0))
        (d_2833_v3_)[index459_] = (_dafny.SeqWithoutIsStrInference([d_2831_v1_ for d_2942_i6_ in range(default__.abs(-129))])) < ((d_2836_v4_) + (d_2836_v4_))
        (globalState).f18 = (0) - ((p1) + (p1))
        d_2943_v82_: _dafny.MultiSet
        d_2943_v82_ = _dafny.MultiSet([d_2831_v1_])
        d_2944_v83_: _dafny.MultiSet
        d_2944_v83_ = _dafny.MultiSet([default__.fm2(d_2943_v82_, 45, globalState)])
        r0 = d_2944_v83_
        r1 = default__.safeDivisionInt(p1, p1)
        r2 = (d_2831_v1_ if not((not((d_2833_v3_)[default__.safeIndex(112, (d_2833_v3_).length(0))])) == ((d_2833_v3_)[default__.safeIndex(112, (d_2833_v3_).length(0))])) else d_2831_v1_)
        return r0, r1, r2

    def m3(self, globalState):
        r0: int = int(0)
        d_2945_v0_: int
        d_2945_v0_ = 224
        r0 = (0) - (d_2945_v0_)
        d_2946_v1_: bool
        d_2946_v1_ = True
        d_2947_v2_: _dafny.Map
        d_2947_v2_ = _dafny.Map({not(d_2946_v1_): (0) - (d_2945_v0_)})
        (globalState).f18 = len(d_2947_v2_)
        (globalState).f17 = d_2945_v0_
        d_2948_v3_: _dafny.Array
        nw462_ = _dafny.Array(_dafny.CodePoint('D'), 7)
        d_2948_v3_ = nw462_
        d_2949_v4_: _dafny.Seq
        d_2949_v4_ = _dafny.SeqWithoutIsStrInference([d_2948_v3_, d_2948_v3_])
        d_2950_i0_: int
        d_2950_i0_ = 0
        with _dafny.label("14"):
            while ((d_2949_v4_) + (d_2949_v4_)) <= (d_2949_v4_):
                with _dafny.c_label("14"):
                    if (d_2950_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_2950_i0_ = (d_2950_i0_) + (1)
                    (globalState).f5 = False
                    d_2951_v5_: _dafny.MultiSet
                    d_2951_v5_ = _dafny.MultiSet([True])
                    (globalState).f18 = ((d_2951_v5_)[d_2946_v1_] if (d_2946_v1_) in (d_2951_v5_) else d_2945_v0_)
                    d_2952_v6_: _dafny.Array
                    nw463_ = _dafny.Array(False, 25)
                    d_2952_v6_ = nw463_
                    index460_ = default__.safeIndex(947, (d_2952_v6_).length(0))
                    (d_2952_v6_)[index460_] = d_2946_v1_
                    d_2953_v7_: _dafny.Map
                    d_2953_v7_ = _dafny.Map({d_2946_v1_: not(d_2946_v1_)})
                    index461_ = default__.safeIndex(947, (d_2952_v6_).length(0))
                    rhs426_ = (d_2952_v6_ if d_2946_v1_ else d_2952_v6_)
                    rhs427_ = d_2946_v1_
                    rhs428_ = d_2945_v0_
                    rhs429_ = d_2945_v0_
                    rhs430_ = (((d_2953_v7_)[d_2946_v1_] if (d_2946_v1_) in (d_2953_v7_) else True) if d_2946_v1_ else d_2946_v1_)
                    lhs390_ = globalState
                    lhs391_ = d_2952_v6_
                    lhs392_ = default__.safeIndex(947, (d_2952_v6_).length(0))
                    lhs393_ = globalState
                    lhs394_ = globalState
                    lhs395_ = globalState
                    lhs390_.f4 = rhs426_
                    lhs391_[lhs392_] = rhs427_
                    lhs393_.f18 = rhs428_
                    lhs394_.f17 = rhs429_
                    lhs395_.f6 = rhs430_
                    (globalState).f17 = d_2945_v0_
                    pass
            pass
        d_2954_v8_: _dafny.Map
        d_2954_v8_ = _dafny.Map({d_2946_v1_: True})
        d_2954_v8_ = (d_2954_v8_).set(d_2946_v1_, d_2946_v1_)
        d_2955_i1_: int
        d_2955_i1_ = 0
        with _dafny.label("15"):
            while d_2946_v1_:
                with _dafny.c_label("15"):
                    if (d_2955_i1_) >= (100):
                        raise _dafny.Break("15")
                    d_2955_i1_ = (d_2955_i1_) + (1)
                    d_2946_v1_ = not(d_2946_v1_)
                    d_2956_v9_: _dafny.MultiSet
                    d_2956_v9_ = _dafny.MultiSet([d_2945_v0_])
                    d_2957_v10_: _dafny.Map
                    d_2957_v10_ = _dafny.Map({d_2946_v1_: _dafny.SeqWithoutIsStrInference([((d_2956_v9_)[(0) - (d_2945_v0_)] if ((0) - (d_2945_v0_)) in (d_2956_v9_) else d_2945_v0_)])})
                    d_2957_v10_ = d_2957_v10_
                    d_2958_v11_: T0
                    nw464_ = C3()
                    nw464_.ctor__((self).f20)
                    d_2958_v11_ = nw464_
                    rhs431_ = d_2958_v11_
                    rhs432_ = (d_2946_v1_) not in (d_2947_v2_)
                    lhs396_ = globalState
                    d_2958_v11_ = rhs431_
                    lhs396_.f6 = rhs432_
                    (globalState).f17 = d_2945_v0_
                    pass
            pass
        d_2959_v12_: str
        d_2959_v12_ = _dafny.CodePoint('a')
        d_2960_v13_: _dafny.Seq
        d_2960_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ru"))
        r0 = len((((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwjneywp"))).set(default__.safeIndex(d_2945_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwjneywp")))), d_2959_v12_)) + ((d_2960_v13_).set(default__.safeIndex(d_2945_v0_, len(d_2960_v13_)), d_2959_v12_)) if not((d_2945_v0_) == (len(default__.fm78(globalState)))) else d_2960_v13_))
        return r0

    def m4(self, p0, p1, p2, globalState):
        d_2961_v0_: C6
        nw465_ = C6()
        nw465_.ctor__((self).f20)
        d_2961_v0_ = nw465_
        d_2962_v1_: bool
        d_2962_v1_ = False
        d_2963_v2_: _dafny.Seq
        d_2963_v2_ = _dafny.SeqWithoutIsStrInference([d_2962_v1_])
        hi10_ = (p1) + (len(d_2963_v2_))
        for d_2964_i0_ in range(p1, hi10_):
            d_2962_v1_ = d_2962_v1_
            d_2965_v3_: _dafny.Map
            d_2965_v3_ = _dafny.Map({p1: -609})
            d_2966_v4_: _dafny.MultiSet
            d_2966_v4_ = _dafny.MultiSet([len(d_2965_v3_)])
            d_2967_v5_: _dafny.Map
            d_2967_v5_ = _dafny.Map({((d_2966_v4_)[d_2964_i0_] if (d_2964_i0_) in (d_2966_v4_) else d_2964_i0_): (0) - (p1)})
            d_2968_v6_: _dafny.Seq
            d_2968_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvxnkeuh"))
            d_2969_v7_: D0
            d_2969_v7_ = D0_DC1(_dafny.CodePoint('r'), d_2964_i0_, _dafny.Set({d_2962_v1_}), d_2962_v1_, False)
            d_2970_v8_: _dafny.Array
            nw466_ = _dafny.Array(None, 22)
            nw466_[int(0)] = (0) - (p1)
            nw466_[int(1)] = p1
            nw466_[int(2)] = d_2964_i0_
            nw466_[int(3)] = p1
            nw466_[int(4)] = p1
            nw466_[int(5)] = (d_2964_i0_) * (d_2964_i0_)
            nw466_[int(6)] = (0) - (p1)
            nw466_[int(7)] = p1
            nw466_[int(8)] = d_2964_i0_
            nw466_[int(9)] = d_2964_i0_
            nw466_[int(10)] = (p1) + (p1)
            nw466_[int(11)] = (0) - (p1)
            nw466_[int(12)] = d_2964_i0_
            nw466_[int(13)] = default__.safeModuloInt(-840, len(d_2967_v5_))
            nw466_[int(14)] = -62
            nw466_[int(15)] = d_2964_i0_
            nw466_[int(16)] = default__.safeModuloInt((d_2966_v4_).cardinality, p1)
            nw466_[int(17)] = len((d_2968_v6_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2971_i1_ in range(default__.abs(-54))])))
            nw466_[int(18)] = default__.fm0(d_2964_i0_, d_2969_v7_, globalState)
            nw466_[int(19)] = (d_2964_i0_) - (243)
            nw466_[int(20)] = len(d_2968_v6_)
            nw466_[int(21)] = d_2964_i0_
            d_2970_v8_ = nw466_
            index462_ = default__.safeIndex(128, (d_2970_v8_).length(0))
            (d_2970_v8_)[index462_] = 421
            d_2972_v9_: _dafny.Array
            nw467_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_2972_v9_ = nw467_
            d_2973_v10_: _dafny.Set
            d_2973_v10_ = _dafny.Set({p1})
            d_2974_v11_: _dafny.Seq
            d_2974_v11_ = _dafny.SeqWithoutIsStrInference([len(d_2973_v10_), p1])
            index463_ = default__.safeIndex(544, (d_2970_v8_).length(0))
            (d_2970_v8_)[index463_] = (0) - (((d_2966_v4_)[(d_2974_v11_)[default__.safeIndex(d_2964_i0_, len(d_2974_v11_))]] if ((d_2974_v11_)[default__.safeIndex(d_2964_i0_, len(d_2974_v11_))]) in (d_2966_v4_) else -528))
            d_2975_v12_: _dafny.Set
            d_2975_v12_ = _dafny.Set({d_2962_v1_})
            d_2976_v13_: _dafny.Array
            nw468_ = _dafny.Array(_dafny.Seq({}), 9)
            d_2976_v13_ = nw468_
            d_2977_v14_: D20
            d_2977_v14_ = D20_DC54((0) - (p1), d_2976_v13_, p1, False, d_2964_i0_)
            d_2978_v15_: _dafny.Set
            d_2978_v15_ = _dafny.Set({d_2975_v12_, default__.fm52(p1, d_2968_v6_, (d_2977_v14_).cf103, d_2962_v1_, globalState)})
            d_2979_v16_: D6
            d_2979_v16_ = D6_DC17(d_2966_v4_, (d_2978_v15_).intersection(d_2978_v15_))
            d_2980_v17_: C7
            nw469_ = C7()
            nw469_.ctor__(d_2962_v1_, (self).f20)
            d_2980_v17_ = nw469_
            d_2981_v18_: _dafny.Seq
            d_2981_v18_ = _dafny.SeqWithoutIsStrInference([d_2980_v17_])
            d_2982_v19_: _dafny.Array
            def lambda150_(d_2983_i2_):
                return _dafny.CodePoint('b')

            init76_ = lambda150_
            nw470_ = _dafny.Array(None, 25)
            for i0_76_ in range(nw470_.length(0)):
                nw470_[i0_76_] = init76_(i0_76_)
            d_2982_v19_ = nw470_
            d_2984_v20_: _dafny.Seq
            d_2984_v20_ = _dafny.SeqWithoutIsStrInference([d_2982_v19_, d_2982_v19_, d_2982_v19_, d_2982_v19_])
            index464_ = default__.safeIndex(128, (d_2970_v8_).length(0))
            index465_ = default__.safeIndex(544, (d_2970_v8_).length(0))
            rhs433_ = default__.safeModuloInt(d_2964_i0_, p1)
            rhs434_ = d_2972_v9_
            rhs435_ = (d_2981_v18_) < (d_2981_v18_)
            rhs436_ = (0) - (len(d_2984_v20_))
            rhs437_ = d_2979_v16_
            lhs397_ = d_2970_v8_
            lhs398_ = default__.safeIndex(128, (d_2970_v8_).length(0))
            lhs399_ = globalState
            lhs400_ = d_2970_v8_
            lhs401_ = default__.safeIndex(544, (d_2970_v8_).length(0))
            lhs397_[lhs398_] = rhs433_
            d_2972_v9_ = rhs434_
            lhs399_.f5 = rhs435_
            lhs400_[lhs401_] = rhs436_
            d_2979_v16_ = rhs437_
            d_2985_v21_: _dafny.Array
            nw471_ = _dafny.Array(False, 1)
            d_2985_v21_ = nw471_
            index466_ = default__.safeIndex(734, (d_2985_v21_).length(0))
            (d_2985_v21_)[index466_] = (d_2980_v17_).f26
            d_2986_v22_: str
            d_2986_v22_ = _dafny.CodePoint('b')
            index467_ = default__.safeIndex(734, (d_2985_v21_).length(0))
            index468_ = default__.safeIndex(128, (d_2970_v8_).length(0))
            rhs438_ = (not(d_2962_v1_) if d_2962_v1_ else d_2962_v1_)
            rhs439_ = (d_2970_v8_)[default__.safeIndex(128, (d_2970_v8_).length(0))]
            rhs440_ = (d_2962_v1_ if (d_2980_v17_).f26 else d_2962_v1_)
            rhs441_ = ((d_2986_v22_ if d_2962_v1_ else _dafny.CodePoint('j'))) in ((d_2968_v6_ if not((d_2963_v2_)[default__.safeIndex((d_2970_v8_)[default__.safeIndex(128, (d_2970_v8_).length(0))], len(d_2963_v2_))]) else default__.fm46(((d_2966_v4_)[len((d_2968_v6_).set(default__.safeIndex(d_2964_i0_, len(d_2968_v6_)), _dafny.CodePoint('h')))] if (len((d_2968_v6_).set(default__.safeIndex(d_2964_i0_, len(d_2968_v6_)), _dafny.CodePoint('h')))) in (d_2966_v4_) else p1), (d_2980_v17_).f26, (d_2980_v17_).f26, globalState)))
            rhs442_ = d_2964_i0_
            lhs402_ = d_2985_v21_
            lhs403_ = default__.safeIndex(734, (d_2985_v21_).length(0))
            lhs404_ = d_2970_v8_
            lhs405_ = default__.safeIndex(128, (d_2970_v8_).length(0))
            lhs406_ = globalState
            lhs402_[lhs403_] = rhs438_
            lhs404_[lhs405_] = rhs439_
            d_2962_v1_ = rhs440_
            d_2962_v1_ = rhs441_
            lhs406_.f17 = rhs442_
            (globalState).f6 = (((d_2970_v8_)[default__.safeIndex(128, (d_2970_v8_).length(0))] if not((d_2980_v17_).f26) else (d_2970_v8_)[default__.safeIndex(128, (d_2970_v8_).length(0))])) < (len((_dafny.SeqWithoutIsStrInference([d_2986_v22_ for d_2987_i3_ in range(default__.abs(872))])).set(default__.safeIndex((d_2970_v8_)[default__.safeIndex(128, (d_2970_v8_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_2986_v22_ for d_2988_i3_ in range(default__.abs(872))]))), d_2986_v22_)))
        (globalState).f17 = p1
        d_2989_v23_: str
        d_2989_v23_ = _dafny.CodePoint('t')
        d_2990_v24_: _dafny.Set
        d_2990_v24_ = _dafny.Set({d_2962_v1_, d_2962_v1_})
        d_2991_v25_: D0
        d_2991_v25_ = D0_DC1(d_2989_v23_, p1, d_2990_v24_, d_2962_v1_, d_2962_v1_)
        (globalState).f17 = default__.fm0((0) - (p1), d_2991_v25_, globalState)
        d_2992_v26_: _dafny.Map
        d_2992_v26_ = _dafny.Map({d_2962_v1_: d_2962_v1_})
        if (d_2962_v1_) == ((d_2962_v1_ if d_2962_v1_ else ((d_2992_v26_)[d_2962_v1_] if (d_2962_v1_) in (d_2992_v26_) else d_2962_v1_))):
            d_2993_v27_: _dafny.Seq
            d_2993_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evokl"))
            (globalState).f5 = not((d_2989_v23_) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "um"))) + (d_2993_v27_)))
            d_2994_v28_: _dafny.Array
            def lambda151_(d_2995_p1_):
                def lambda152_(d_2996_i4_):
                    return (d_2996_i4_) - ((0) - (d_2995_p1_))

                return lambda152_

            init77_ = lambda151_(p1)
            nw472_ = _dafny.Array(None, 7)
            for i0_77_ in range(nw472_.length(0)):
                nw472_[i0_77_] = init77_(i0_77_)
            d_2994_v28_ = nw472_
            d_2997_v29_: _dafny.Seq
            d_2997_v29_ = _dafny.SeqWithoutIsStrInference([p1])
            d_2998_v30_: _dafny.MultiSet
            d_2998_v30_ = _dafny.MultiSet([len(d_2997_v29_)])
            d_2999_v31_: _dafny.Set
            d_2999_v31_ = _dafny.Set({d_2990_v24_})
            d_3000_v32_: D6
            d_3000_v32_ = D6_DC17((d_2998_v30_) | (d_2998_v30_), (d_2999_v31_).intersection(d_2999_v31_))
            rhs443_ = d_2994_v28_
            rhs444_ = d_3000_v32_
            d_2994_v28_ = rhs443_
            d_3000_v32_ = rhs444_
            nw473_ = _dafny.Array(int(0), 1)
            d_2994_v28_ = nw473_
            (globalState).f15 = (d_2993_v27_).set(default__.safeIndex(p1, len(d_2993_v27_)), d_2989_v23_)
            if d_2962_v1_:
                index469_ = default__.safeIndex(619, (d_2994_v28_).length(0))
                (d_2994_v28_)[index469_] = p1
                index470_ = default__.safeIndex(619, (d_2994_v28_).length(0))
                (d_2994_v28_)[index470_] = default__.fm0(p1, d_2991_v25_, globalState)
                (globalState).f5 = False
                (globalState).f18 = default__.fm0((p1) - (default__.fm0(p1, d_2991_v25_, globalState)), default__.fm48(len(d_2990_v24_), (d_2994_v28_)[default__.safeIndex(619, (d_2994_v28_).length(0))], p1, globalState), globalState)
                d_3001_v33_: _dafny.Map
                d_3001_v33_ = _dafny.Map({d_2993_v27_: (p1 if d_2962_v1_ else (0) - ((0) - ((0) - ((d_2994_v28_)[default__.safeIndex(619, (d_2994_v28_).length(0))]))))})
                d_3002_v34_: _dafny.Map
                d_3002_v34_ = _dafny.Map({not(d_2962_v1_): d_2961_v0_})
                d_3001_v33_ = (d_3001_v33_).set((d_2993_v27_) + (_dafny.SeqWithoutIsStrInference([d_2989_v23_ for d_3003_i5_ in range(default__.abs(-221))])), (0) - ((len(d_3002_v34_)) + (p1)))
                d_3004_v35_: _dafny.Map
                d_3004_v35_ = _dafny.Map({(d_2994_v28_)[default__.safeIndex(619, (d_2994_v28_).length(0))]: d_2962_v1_})
                d_3005_v36_: D4
                d_3005_v36_ = D4_DC11(d_2962_v1_, True, (d_2994_v28_)[default__.safeIndex(619, (d_2994_v28_).length(0))], len(d_3004_v35_))
                d_3006_v37_: _dafny.Map
                d_3006_v37_ = _dafny.Map({p1: (d_3005_v36_).cf22})
                d_3006_v37_ = (d_3006_v37_).set((0) - ((d_2994_v28_)[default__.safeIndex(619, (d_2994_v28_).length(0))]), len(default__.fm45(globalState)))
            elif True:
                d_3007_v38_: _dafny.Map
                d_3007_v38_ = _dafny.Map({p1: d_2962_v1_})
                d_3008_v39_: _dafny.Map
                d_3008_v39_ = _dafny.Map({default__.safeModuloInt(p1, p1): not (d_2962_v1_) or (((d_3007_v38_)[len(_dafny.Set({d_2962_v1_}))] if (len(_dafny.Set({d_2962_v1_}))) in (d_3007_v38_) else d_2962_v1_))})
                d_3009_v40_: _dafny.Map
                d_3009_v40_ = _dafny.Map({d_2989_v23_: d_2997_v29_})
                d_3010_v41_: _dafny.Seq
                d_3010_v41_ = _dafny.SeqWithoutIsStrInference([d_3007_v38_])
                d_3011_v42_: _dafny.Set
                d_3011_v42_ = _dafny.Set({p1})
                d_3012_v43_: _dafny.Map
                d_3012_v43_ = _dafny.Map({d_2962_v1_: len(d_3011_v42_)})
                d_3013_v44_: _dafny.Map
                d_3013_v44_ = _dafny.Map({p0: d_2962_v1_})
                d_3014_v45_: _dafny.MultiSet
                d_3014_v45_ = _dafny.MultiSet([d_2989_v23_, d_2989_v23_])
                rhs445_ = len(((d_3009_v40_)[d_2989_v23_] if (d_2989_v23_) in (d_3009_v40_) else d_2997_v29_))
                rhs446_ = d_2962_v1_
                rhs447_ = ((d_3010_v41_)[default__.safeIndex(((d_3012_v43_)[d_2962_v1_] if (d_2962_v1_) in (d_3012_v43_) else p1), len(d_3010_v41_))]).set((p1) * (p1), ((d_3013_v44_)[p0] if (p0) in (d_3013_v44_) else default__.fm2(d_3014_v45_, p1, globalState)))
                lhs407_ = globalState
                lhs408_ = globalState
                lhs407_.f17 = rhs445_
                lhs408_.f5 = rhs446_
                d_3008_v39_ = rhs447_
                d_3015_v46_: _dafny.Array
                def lambda153_(d_3016_v27_):
                    def lambda154_(d_3017_i6_):
                        return d_3016_v27_

                    return lambda154_

                init78_ = lambda153_(d_2993_v27_)
                nw474_ = _dafny.Array(None, 15)
                for i0_78_ in range(nw474_.length(0)):
                    nw474_[i0_78_] = init78_(i0_78_)
                d_3015_v46_ = nw474_
                d_3018_v47_: D7
                d_3018_v47_ = D7_DC21(p1, d_2962_v1_, d_2962_v1_, d_2994_v28_, d_2962_v1_)
                d_3019_v48_: _dafny.Map
                d_3019_v48_ = _dafny.Map({d_3015_v46_: d_3018_v47_})
                d_3019_v48_ = (d_3019_v48_).set(d_3015_v46_, d_3018_v47_)
                d_3020_v49_: _dafny.Array
                nw475_ = _dafny.Array(None, 6)
                nw475_[int(0)] = d_2997_v29_
                nw475_[int(1)] = d_2997_v29_
                nw475_[int(2)] = d_2997_v29_
                nw475_[int(3)] = d_2997_v29_
                nw475_[int(4)] = _dafny.SeqWithoutIsStrInference([p1])
                nw475_[int(5)] = d_2997_v29_
                d_3020_v49_ = nw475_
                d_3021_v50_: D14
                d_3021_v50_ = D14_DC41(-117, 59, p1, D3_DC7(d_3020_v49_))
                d_3022_v51_: _dafny.Set
                d_3022_v51_ = _dafny.Set({d_3021_v50_})
                d_3023_v52_: _dafny.Map
                d_3023_v52_ = _dafny.Map({d_2962_v1_: d_2997_v29_})
                d_3024_v53_: D3
                d_3024_v53_ = D3_DC7(d_3020_v49_)
                d_3025_v54_: _dafny.Map
                d_3025_v54_ = _dafny.Map({d_2962_v1_: _dafny.Set({d_3021_v50_, d_3021_v50_, D14_DC41(len(d_2993_v27_), len(_dafny.SeqWithoutIsStrInference([d_2989_v23_ for d_3026_i7_ in range(default__.abs(548))])), len(((d_3023_v52_)[True] if (True) in (d_3023_v52_) else d_2997_v29_)), d_3024_v53_)})})
                d_3027_v55_: _dafny.Seq
                d_3027_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_3021_v50_}), d_3022_v51_, ((d_3025_v54_)[d_2962_v1_] if (d_2962_v1_) in (d_3025_v54_) else d_3022_v51_)])
                d_3028_v56_: C8
                nw476_ = C8()
                nw476_.ctor__((d_3022_v51_).isdisjoint((d_3027_v55_)[default__.safeIndex(280, len(d_3027_v55_))]), D0_DC0(d_2962_v1_))
                d_3028_v56_ = nw476_
                d_2992_v26_ = (d_2992_v26_).set((d_2998_v30_).ispropersubset(_dafny.MultiSet([p1, len(_dafny.SeqWithoutIsStrInference([d_2989_v23_ for d_3029_i8_ in range(default__.abs(830))])), p1])), d_2962_v1_)
                (globalState).f17 = p1
        elif True:
            d_3030_v57_: _dafny.MultiSet
            d_3031_v58_: int
            d_3032_v59_: str
            out29_: _dafny.MultiSet
            out30_: int
            out31_: str
            out29_, out30_, out31_ = (p0).m0(d_2962_v1_, (p1) * (p1), globalState)
            d_3030_v57_ = out29_
            d_3031_v58_ = out30_
            d_3032_v59_ = out31_
            d_3033_v60_: _dafny.Array
            nw477_ = _dafny.Array(int(0), 10)
            d_3033_v60_ = nw477_
            d_3034_v61_: _dafny.Seq
            d_3034_v61_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            index471_ = default__.safeIndex(775, (d_3033_v60_).length(0))
            (d_3033_v60_)[index471_] = (d_3034_v61_)[default__.safeIndex(162, len(d_3034_v61_))]
            d_3035_v62_: _dafny.Array
            nw478_ = _dafny.Array(False, 8)
            d_3035_v62_ = nw478_
            index472_ = default__.safeIndex(775, (d_3033_v60_).length(0))
            rhs448_ = p1
            rhs449_ = len((((_dafny.SeqWithoutIsStrInference([d_3031_v58_, 792])) + ((d_3034_v61_).set(default__.safeIndex(p1, len(d_3034_v61_)), 646))).set(default__.safeIndex(447, len((_dafny.SeqWithoutIsStrInference([d_3031_v58_, 792])) + ((d_3034_v61_).set(default__.safeIndex(p1, len(d_3034_v61_)), 646)))), p1)) + (d_3034_v61_))
            rhs450_ = (d_3035_v62_ if d_2962_v1_ else d_3035_v62_)
            rhs451_ = d_2962_v1_
            lhs409_ = d_3033_v60_
            lhs410_ = default__.safeIndex(775, (d_3033_v60_).length(0))
            lhs411_ = globalState
            lhs412_ = globalState
            lhs413_ = globalState
            lhs409_[lhs410_] = rhs448_
            lhs411_.f18 = rhs449_
            lhs412_.f4 = rhs450_
            lhs413_.f6 = rhs451_
            d_3036_v63_: _dafny.Set
            d_3036_v63_ = _dafny.Set({d_3031_v58_, (d_3034_v61_)[default__.safeIndex((_dafny.MultiSet([d_3031_v58_])).cardinality, len(d_3034_v61_))]})
            d_3036_v63_ = (d_3036_v63_) - ((_dafny.Set({658, (d_3033_v60_)[default__.safeIndex(775, (d_3033_v60_).length(0))], p1})) | (d_3036_v63_))
            d_3037_v64_: _dafny.Seq
            d_3037_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "befrf"))
            d_3038_v65_: _dafny.Map
            d_3038_v65_ = _dafny.Map({(d_3033_v60_)[default__.safeIndex(775, (d_3033_v60_).length(0))]: d_3037_v64_})
            d_3039_v66_: _dafny.Map
            d_3039_v66_ = _dafny.Map({d_2962_v1_: (d_3033_v60_)[default__.safeIndex(775, (d_3033_v60_).length(0))]})
            d_3040_v67_: _dafny.Map
            d_3040_v67_ = _dafny.Map({len(d_2990_v24_): len(d_3039_v66_)})
            d_3038_v65_ = (d_3038_v65_).set((len(d_3036_v63_)) * (len(d_3040_v67_)), (d_3037_v64_) + (d_3037_v64_))
            d_3041_v68_: D13
            d_3041_v68_ = D13_DC36((d_3033_v60_)[default__.safeIndex(775, (d_3033_v60_).length(0))])
            index473_ = default__.safeIndex(559, (d_3035_v62_).length(0))
            (d_3035_v62_)[index473_] = (p1) <= (len(d_3037_v64_))
            index474_ = default__.safeIndex(890, (d_3035_v62_).length(0))
            (d_3035_v62_)[index474_] = d_2962_v1_
            d_3042_v69_: _dafny.Seq
            d_3042_v69_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_3037_v64_)])
            pat_let_tv112_ = d_3039_v66_
            index475_ = default__.safeIndex(559, (d_3035_v62_).length(0))
            index476_ = default__.safeIndex(890, (d_3035_v62_).length(0))
            def iife310_(_pat_let115_0):
                def iife311_(d_3043_dt__update__tmp_h0_):
                    def iife312_(_pat_let116_0):
                        def iife313_(d_3044_dt__update_hcf68_h0_):
                            return D13_DC36(d_3044_dt__update_hcf68_h0_)
                        return iife313_(_pat_let116_0)
                    return iife312_(len(pat_let_tv112_))
                return iife311_(_pat_let115_0)
            rhs452_ = iife310_(d_3041_v68_)
            rhs453_ = d_2962_v1_
            rhs454_ = (d_3042_v69_)[default__.safeIndex(((d_2991_v25_).cf2) + ((d_3033_v60_)[default__.safeIndex(775, (d_3033_v60_).length(0))]), len(d_3042_v69_))]
            rhs455_ = d_2962_v1_
            lhs414_ = d_3035_v62_
            lhs415_ = default__.safeIndex(559, (d_3035_v62_).length(0))
            lhs416_ = globalState
            lhs417_ = d_3035_v62_
            lhs418_ = default__.safeIndex(890, (d_3035_v62_).length(0))
            d_3041_v68_ = rhs452_
            lhs414_[lhs415_] = rhs453_
            lhs416_.f16 = rhs454_
            lhs417_[lhs418_] = rhs455_
        d_3045_v70_: D7
        d_3045_v70_ = D7_DC20(len(d_2992_v26_), not(d_2962_v1_))
        d_3046_v71_: _dafny.Map
        d_3046_v71_ = _dafny.Map({d_3045_v70_: p1})
        d_3047_v72_: _dafny.Map
        d_3047_v72_ = _dafny.Map({((d_3046_v71_)[d_3045_v70_] if (d_3045_v70_) in (d_3046_v71_) else p1): d_2962_v1_})
        d_3048_v73_: D17
        d_3048_v73_ = D17_DC48(d_2962_v1_, d_2962_v1_)
        d_3049_v74_: D21
        d_3049_v74_ = D21_DC58(d_2962_v1_, _dafny.SeqWithoutIsStrInference([d_2989_v23_ for d_3050_i9_ in range(default__.abs(820))]))
        rhs456_ = (default__.fm59(p1, globalState)) | (default__.fm59(p1, globalState))
        rhs457_ = (0) - (p1)
        rhs458_ = D17_DC48((d_3049_v74_).cf108, (d_2963_v2_)[default__.safeIndex(p1, len(d_2963_v2_))])
        lhs419_ = globalState
        d_3047_v72_ = rhs456_
        lhs419_.f17 = rhs457_
        d_3048_v73_ = rhs458_


class C10(T1, T0):
    def  __init__(self):
        self._f20: D0 = D0.default()()
        self._f21: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21
    def ctor__(self, f21, f20):
        (self)._f21 = f21
        (self)._f20 = f20

    def fm3(self, p0, p1, p2, globalState):
        if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_3051_i0_ in range(default__.abs(137))])) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_3052_i1_ in range(default__.abs(952))])):
            return (_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbuyy"))): _dafny.CodePoint('l')})): True})) | (_dafny.Map({(self).f21: (D0_DC1(_dafny.CodePoint('e'), (self).f21, _dafny.Set({True}), False, False)).cf5}))
        elif True:
            return (_dafny.Map({(self).f21: True})) | (_dafny.Map({419: True}))

    def fm4(self, p0, p1, globalState):
        return (self).f20

    def fm5(self, p0, globalState):
        return (0) - ((self).f21)

    def m0(self, p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: str = _dafny.CodePoint('D')
        d_3053_v0_: str
        d_3053_v0_ = _dafny.CodePoint('k')
        d_3054_v1_: _dafny.MultiSet
        d_3054_v1_ = _dafny.MultiSet([_dafny.CodePoint('l'), d_3053_v0_])
        d_3055_v2_: _dafny.MultiSet
        d_3055_v2_ = _dafny.MultiSet([(self).f21])
        d_3056_v3_: _dafny.Seq
        d_3056_v3_ = _dafny.SeqWithoutIsStrInference([(self).fm5(not(p0), globalState)])
        if not (default__.fm2(d_3054_v1_, (d_3055_v2_).cardinality, globalState)) or (default__.fm2(_dafny.MultiSet([d_3053_v0_]), len(d_3056_v3_), globalState)):
            d_3057_v4_: _dafny.Seq
            d_3057_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkftoju"))
            d_3058_v5_: _dafny.Set
            d_3058_v5_ = _dafny.Set({p0, not(p0)})
            d_3059_v6_: _dafny.MultiSet
            d_3059_v6_ = _dafny.MultiSet([d_3058_v5_, d_3058_v5_, d_3058_v5_])
            d_3060_v7_: _dafny.Map
            d_3060_v7_ = _dafny.Map({(0) - (len(d_3057_v4_)): p0})
            d_3061_v8_: _dafny.Array
            nw479_ = _dafny.Array(None, 26)
            nw479_[int(0)] = p0
            nw479_[int(1)] = p0
            nw479_[int(2)] = (p1) == (p1)
            nw479_[int(3)] = (d_3057_v4_) == (_dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3062_i0_ in range(default__.abs(62))]))
            nw479_[int(4)] = p0
            nw479_[int(5)] = p0
            nw479_[int(6)] = p0
            nw479_[int(7)] = p0
            nw479_[int(8)] = p0
            nw479_[int(9)] = (d_3059_v6_) != (d_3059_v6_)
            nw479_[int(10)] = False
            nw479_[int(11)] = p0
            nw479_[int(12)] = p0
            nw479_[int(13)] = p0
            nw479_[int(14)] = p0
            nw479_[int(15)] = p0
            nw479_[int(16)] = p0
            nw479_[int(17)] = (p1) < (len(d_3060_v7_))
            nw479_[int(18)] = ((d_3060_v7_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjitwfms")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjitwfms")))) in (d_3060_v7_) else p0)
            nw479_[int(19)] = p0
            nw479_[int(20)] = (d_3053_v0_) not in (d_3057_v4_)
            nw479_[int(21)] = default__.fm2(d_3054_v1_, p1, globalState)
            nw479_[int(22)] = p0
            nw479_[int(23)] = p0
            nw479_[int(24)] = ((self).f21) <= (p1)
            nw479_[int(25)] = p0
            d_3061_v8_ = nw479_
            index477_ = default__.safeIndex(552, (d_3061_v8_).length(0))
            (d_3061_v8_)[index477_] = (p1) >= ((self).f21)
            d_3063_v9_: _dafny.Seq
            d_3063_v9_ = _dafny.SeqWithoutIsStrInference([p0])
            d_3064_v10_: D2
            d_3064_v10_ = D2_DC3(d_3057_v4_)
            d_3065_v11_: _dafny.MultiSet
            d_3065_v11_ = _dafny.MultiSet([d_3063_v9_])
            d_3066_v12_: _dafny.Map
            d_3066_v12_ = _dafny.Map({(self).f21: (d_3065_v11_).set(d_3063_v9_, default__.abs(42))})
            index478_ = default__.safeIndex(552, (d_3061_v8_).length(0))
            (d_3061_v8_)[index478_] = not(((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_3063_v9_, d_3063_v9_]))).set(d_3063_v9_, default__.abs(len((d_3064_v10_).cf7)))).isdisjoint(((d_3066_v12_)[len(_dafny.Map({p0: p0}))] if (len(_dafny.Map({p0: p0}))) in (d_3066_v12_) else d_3065_v11_)))
            d_3067_v13_: _dafny.Map
            d_3067_v13_ = _dafny.Map({(d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))]: (self).f21})
            d_3068_v14_: _dafny.Map
            d_3068_v14_ = _dafny.Map({p1: -460})
            d_3069_v16_: D0
            def iife314_():
                coll80_ = _dafny.Set()
                compr_80_: int
                for compr_80_ in _dafny.IntegerRange(687, 982):
                    d_3070_v15_: int = compr_80_
                    if ((687) <= (d_3070_v15_)) and ((d_3070_v15_) < (982)):
                        coll80_ = coll80_.union(_dafny.Set([default__.safeModuloInt(d_3070_v15_, len(d_3067_v13_))]))
                return _dafny.Set(coll80_)
            d_3069_v16_ = D0_DC1(default__.fm1(globalState), len(iife314_()
), d_3058_v5_, (d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))], p0)
            d_3071_v17_: _dafny.MultiSet
            d_3071_v17_ = _dafny.MultiSet([d_3068_v14_, d_3068_v14_])
            d_3072_v18_: _dafny.Array
            nw480_ = _dafny.Array(None, 15)
            nw480_[int(0)] = p1
            nw480_[int(1)] = (self).f21
            nw480_[int(2)] = (self).f21
            nw480_[int(3)] = (self).f21
            nw480_[int(4)] = ((d_3067_v13_)[(d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))]] if ((d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))]) in (d_3067_v13_) else (self).f21)
            nw480_[int(5)] = ((self).f21) * (((d_3068_v14_)[927] if (927) in (d_3068_v14_) else len(_dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3073_i1_ in range(default__.abs(322))]))))
            nw480_[int(6)] = (p1) - (default__.fm0((0) - ((self).f21), d_3069_v16_, globalState))
            nw480_[int(7)] = (0) - (p1)
            nw480_[int(8)] = (self).f21
            nw480_[int(9)] = (self).f21
            nw480_[int(10)] = (self).f21
            nw480_[int(11)] = default__.safeDivisionInt(697, (d_3071_v17_).cardinality)
            nw480_[int(12)] = ((self).f21) * (len(d_3067_v13_))
            nw480_[int(13)] = 614
            nw480_[int(14)] = (0) - ((self).f21)
            d_3072_v18_ = nw480_
            index479_ = default__.safeIndex(233, (d_3072_v18_).length(0))
            (d_3072_v18_)[index479_] = p1
            index480_ = default__.safeIndex(233, (d_3072_v18_).length(0))
            def iife315_():
                coll81_ = _dafny.Set()
                compr_81_: int
                for compr_81_ in _dafny.IntegerRange(100, 107):
                    d_3074_v19_: int = compr_81_
                    if ((100) <= (d_3074_v19_)) and ((d_3074_v19_) < (107)):
                        coll81_ = coll81_.union(_dafny.Set([default__.safeModuloInt(d_3074_v19_, len(_dafny.Map({(0) - (-596): d_3053_v0_})))]))
                return _dafny.Set(coll81_)
            rhs459_ = d_3063_v9_
            rhs460_ = default__.fm6(_dafny.Set({False, p0, (d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))]}), (p1) - ((self).f21), globalState)
            rhs461_ = ((_dafny.Set({p1})).ispropersubset(iife315_()
            ) if not((d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))]) else p0)
            rhs462_ = not((d_3063_v9_)[default__.safeIndex(p1, len(d_3063_v9_))])
            rhs463_ = (0) - (default__.safeDivisionInt((self).f21, 100))
            lhs420_ = globalState
            lhs421_ = globalState
            lhs422_ = d_3072_v18_
            lhs423_ = default__.safeIndex(233, (d_3072_v18_).length(0))
            d_3063_v9_ = rhs459_
            d_3064_v10_ = rhs460_
            lhs420_.f5 = rhs461_
            lhs421_.f6 = rhs462_
            lhs422_[lhs423_] = rhs463_
            index481_ = default__.safeIndex(233, (d_3072_v18_).length(0))
            index482_ = default__.safeIndex(552, (d_3061_v8_).length(0))
            rhs464_ = (self).f21
            rhs465_ = d_3061_v8_
            rhs466_ = ((d_3072_v18_)[default__.safeIndex(233, (d_3072_v18_).length(0))]) >= (p1)
            lhs424_ = d_3072_v18_
            lhs425_ = default__.safeIndex(233, (d_3072_v18_).length(0))
            lhs426_ = globalState
            lhs427_ = d_3061_v8_
            lhs428_ = default__.safeIndex(552, (d_3061_v8_).length(0))
            lhs424_[lhs425_] = rhs464_
            lhs426_.f4 = rhs465_
            lhs427_[lhs428_] = rhs466_
            r1 = ((default__.fm7((self).f21, globalState)).cardinality) + (len(d_3057_v4_))
            d_3075_v20_: _dafny.MultiSet
            d_3075_v20_ = _dafny.MultiSet([False])
            if (True) not in ((d_3075_v20_).set((d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))], default__.abs((d_3072_v18_)[default__.safeIndex(233, (d_3072_v18_).length(0))]))):
                d_3076_v21_: _dafny.Map
                d_3076_v21_ = _dafny.Map({not((len(_dafny.Set({not((d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))])}))) != (p1)): p0})
                d_3076_v21_ = (d_3076_v21_).set(not((p0) and ((d_3061_v8_)[default__.safeIndex(552, (d_3061_v8_).length(0))])), p0)
                d_3067_v13_ = (_dafny.Map({p0: (d_3072_v18_)[default__.safeIndex(233, (d_3072_v18_).length(0))]}) if ((self).f21) >= ((self).f21) else d_3067_v13_)
                index483_ = default__.safeIndex(233, (d_3072_v18_).length(0))
                (d_3072_v18_)[index483_] = p1
                pat_let_tv113_ = d_3053_v0_
                pat_let_tv114_ = d_3061_v8_
                pat_let_tv115_ = d_3061_v8_
                pat_let_tv116_ = d_3057_v4_
                def iife316_(_pat_let117_0):
                    def iife317_(d_3077_dt__update__tmp_h0_):
                        def iife318_(_pat_let118_0):
                            def iife319_(d_3079_dt__update_hcf7_h0_):
                                return D2_DC3(d_3079_dt__update_hcf7_h0_)
                            return iife319_(_pat_let118_0)
                        return iife318_((_dafny.SeqWithoutIsStrInference([pat_let_tv113_ for d_3078_i2_ in range(default__.abs(833))]) if (pat_let_tv115_)[default__.safeIndex(552, (pat_let_tv114_).length(0))] else pat_let_tv116_))
                    return iife317_(_pat_let117_0)
                d_3064_v10_ = iife316_(d_3064_v10_)
                index484_ = default__.safeIndex(552, (d_3061_v8_).length(0))
                rhs467_ = (d_3055_v2_).issubset(_dafny.MultiSet(d_3056_v3_))
                rhs468_ = (len((d_3056_v3_) + (_dafny.SeqWithoutIsStrInference([len(d_3057_v4_) for d_3080_i3_ in range(default__.abs(58))])))) <= (default__.safeModuloInt((self).f21, (0) - (p1)))
                lhs429_ = globalState
                lhs430_ = d_3061_v8_
                lhs431_ = default__.safeIndex(552, (d_3061_v8_).length(0))
                lhs429_.f6 = rhs467_
                lhs430_[lhs431_] = rhs468_
            elif True:
                d_3064_v10_ = d_3064_v10_
                index485_ = default__.safeIndex(233, (d_3072_v18_).length(0))
                (d_3072_v18_)[index485_] = (d_3072_v18_)[default__.safeIndex(233, (d_3072_v18_).length(0))]
                index486_ = default__.safeIndex(233, (d_3072_v18_).length(0))
                (d_3072_v18_)[index486_] = (d_3072_v18_)[default__.safeIndex(233, (d_3072_v18_).length(0))]
                d_3055_v2_ = (d_3055_v2_) - (d_3055_v2_)
                d_3081_v22_: C0
                nw481_ = C0()
                nw481_.ctor__(d_3061_v8_, d_3063_v9_)
                d_3081_v22_ = nw481_
        elif True:
            d_3082_v23_: _dafny.Set
            d_3082_v23_ = _dafny.Set({p0, p0})
            (globalState).f5 = (_dafny.Set({p0})).ispropersubset(d_3082_v23_)
            d_3083_v24_: _dafny.Seq
            d_3083_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_3082_v23_ = default__.fm16(d_3083_v24_, globalState)
            d_3084_v25_: int
            out32_: int
            out32_ = (self).m1(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcg")), p0, globalState)
            d_3084_v25_ = out32_
            rhs469_ = not(p0)
            rhs470_ = (0) - (p1)
            lhs432_ = globalState
            lhs432_.f5 = rhs469_
            d_3084_v25_ = rhs470_
            d_3085_v26_: D2
            d_3085_v26_ = D2_DC5(p0, p1)
            d_3086_v27_: D2
            d_3086_v27_ = D2_DC6(d_3085_v26_)
            d_3087_v28_: D2
            d_3087_v28_ = D2_DC6(d_3086_v27_)
            pat_let_tv117_ = d_3086_v27_
            def iife320_(_pat_let119_0):
                def iife321_(d_3088_dt__update__tmp_h1_):
                    def iife322_(_pat_let120_0):
                        def iife323_(d_3089_dt__update_hcf14_h0_):
                            return D2_DC6(d_3089_dt__update_hcf14_h0_)
                        return iife323_(_pat_let120_0)
                    return iife322_(pat_let_tv117_)
                return iife321_(_pat_let119_0)
            d_3087_v28_ = iife320_(d_3087_v28_)
        d_3090_v29_: D2
        d_3090_v29_ = D2_DC3(_dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3091_i4_ in range(default__.abs(700))]))
        source38_ = d_3090_v29_
        if source38_.is_DC4:
            d_3092___mcc_h0_ = source38_.cf8
            d_3093___mcc_h1_ = source38_.cf9
            d_3094___mcc_h2_ = source38_.cf10
            d_3095___mcc_h3_ = source38_.cf11
            d_3096_cf11_ = d_3095___mcc_h3_
            d_3097_cf10_ = d_3094___mcc_h2_
            d_3098_cf9_ = d_3093___mcc_h1_
            d_3099_cf8_ = d_3092___mcc_h0_
            d_3100_v30_: _dafny.Set
            d_3100_v30_ = _dafny.Set({d_3097_cf10_, d_3096_cf11_})
            d_3101_v31_: D0
            d_3101_v31_ = D0_DC1(d_3053_v0_, (self).f21, d_3100_v30_, False, default__.fm2(d_3054_v1_, d_3099_cf8_, globalState))
            pat_let_tv118_ = d_3096_cf11_
            pat_let_tv119_ = d_3053_v0_
            pat_let_tv120_ = d_3097_cf10_
            pat_let_tv121_ = d_3099_cf8_
            pat_let_tv122_ = d_3096_cf11_
            pat_let_tv123_ = d_3096_cf11_
            def iife325_(_pat_let122_0):
                def iife326_(d_3102_dt__update__tmp_h2_):
                    def iife327_(_pat_let123_0):
                        def iife328_(d_3103_dt__update_hcf4_h0_):
                            def iife329_(_pat_let124_0):
                                def iife330_(d_3104_dt__update_hcf1_h0_):
                                    def iife331_(_pat_let125_0):
                                        def iife332_(d_3105_dt__update_hcf5_h0_):
                                            def iife333_(_pat_let126_0):
                                                def iife334_(d_3106_dt__update_hcf2_h0_):
                                                    return D0_DC1(d_3104_dt__update_hcf1_h0_, d_3106_dt__update_hcf2_h0_, (d_3102_dt__update__tmp_h2_).cf3, d_3103_dt__update_hcf4_h0_, d_3105_dt__update_hcf5_h0_)
                                                return iife334_(_pat_let126_0)
                                            return iife333_(pat_let_tv121_)
                                        return iife332_(_pat_let125_0)
                                    return iife331_(pat_let_tv120_)
                                return iife330_(_pat_let124_0)
                            return iife329_(pat_let_tv119_)
                        return iife328_(_pat_let123_0)
                    return iife327_((D0_DC0(pat_let_tv118_)).cf0)
                return iife326_(_pat_let122_0)
            def iife324_(_pat_let121_0):
                def iife335_(d_3107_dt__update__tmp_h3_):
                    def iife336_(_pat_let127_0):
                        def iife337_(d_3108_dt__update_hcf5_h1_):
                            def iife338_(_pat_let128_0):
                                def iife339_(d_3109_dt__update_hcf4_h1_):
                                    return D0_DC1((d_3107_dt__update__tmp_h3_).cf1, (d_3107_dt__update__tmp_h3_).cf2, (d_3107_dt__update__tmp_h3_).cf3, d_3109_dt__update_hcf4_h1_, d_3108_dt__update_hcf5_h1_)
                                return iife339_(_pat_let128_0)
                            return iife338_(pat_let_tv123_)
                        return iife337_(_pat_let127_0)
                    return iife336_((False) and (pat_let_tv122_))
                return iife335_(_pat_let121_0)
            source39_ = iife324_(iife325_(d_3101_v31_))
            if source39_.is_DC1:
                d_3110___mcc_h8_ = source39_.cf1
                d_3111___mcc_h9_ = source39_.cf2
                d_3112___mcc_h10_ = source39_.cf3
                d_3113___mcc_h11_ = source39_.cf4
                d_3114___mcc_h12_ = source39_.cf5
                d_3115_cf5_ = d_3114___mcc_h12_
                d_3116_cf4_ = d_3113___mcc_h11_
                d_3117_cf3_ = d_3112___mcc_h10_
                d_3118_cf2_ = d_3111___mcc_h9_
                d_3119_cf1_ = d_3110___mcc_h8_
                d_3116_cf4_ = ((d_3117_cf3_).intersection(d_3100_v30_)).isdisjoint((d_3100_v30_).intersection(d_3117_cf3_))
                (globalState).f17 = ((0) - (default__.safeModuloInt((self).f21, len(_dafny.Map({len(d_3056_v3_): d_3097_cf10_}))))) - ((0) - (((self).f21) - (d_3099_cf8_)))
                d_3120_v32_: int
                out33_: int
                out33_ = (self).m1(d_3098_cf9_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), not(p0), globalState)
                d_3120_v32_ = out33_
                d_3098_cf9_ = d_3116_cf4_
            elif True:
                d_3121___mcc_h13_ = source39_.cf0
                d_3122_cf0_ = d_3121___mcc_h13_
                d_3123_v33_: _dafny.Seq
                d_3123_v33_ = _dafny.SeqWithoutIsStrInference([p0])
                d_3124_v34_: _dafny.Set
                d_3124_v34_ = _dafny.Set({len(d_3056_v3_), len(d_3123_v33_), 420})
                d_3125_v35_: _dafny.Seq
                d_3125_v35_ = _dafny.SeqWithoutIsStrInference([not(p0), d_3122_cf0_, (d_3124_v34_).isdisjoint(d_3124_v34_)])
                d_3125_v35_ = ((d_3123_v33_) + (d_3125_v35_)) + (((d_3123_v33_) + (d_3125_v35_)).set(default__.safeIndex(p1, len((d_3123_v33_) + (d_3125_v35_))), (d_3125_v35_)[default__.safeIndex((self).f21, len(d_3125_v35_))]))
                d_3126_v36_: _dafny.MultiSet
                d_3126_v36_ = _dafny.MultiSet([d_3096_cf11_])
                d_3127_v37_: _dafny.Map
                d_3127_v37_ = _dafny.Map({(d_3126_v36_).cardinality: d_3096_cf11_})
                d_3123_v33_ = _dafny.SeqWithoutIsStrInference([((d_3127_v37_)[(self).f21] if ((self).f21) in (d_3127_v37_) else d_3122_cf0_)])
                d_3128_v38_: T0
                nw482_ = C4()
                nw482_.ctor__((self).f20)
                d_3128_v38_ = nw482_
                d_3128_v38_ = d_3128_v38_
                (globalState).f18 = p1
            d_3129_v39_: _dafny.Map
            d_3129_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_3099_cf8_]): p0})
            d_3130_v40_: D6
            d_3130_v40_ = D6_DC16(d_3129_v39_)
            source40_ = d_3130_v40_
            if source40_.is_DC17:
                d_3131___mcc_h14_ = source40_.cf28
                d_3132___mcc_h15_ = source40_.cf29
                d_3133_cf29_ = d_3132___mcc_h15_
                d_3134_cf28_ = d_3131___mcc_h14_
                d_3055_v2_ = ((d_3134_cf28_) - (d_3134_cf28_)) - ((d_3134_cf28_) - (d_3055_v2_))
                d_3135_v41_: _dafny.Array
                nw483_ = _dafny.Array(False, 15)
                d_3135_v41_ = nw483_
                d_3136_v42_: _dafny.Map
                d_3136_v42_ = _dafny.Map({(p1) + ((self).f21): d_3135_v41_})
                d_3136_v42_ = (d_3136_v42_).set(len(d_3056_v3_), d_3135_v41_)
                d_3137_v43_: _dafny.Map
                d_3137_v43_ = _dafny.Map({d_3098_cf9_: p0})
                d_3138_v44_: _dafny.Seq
                d_3138_v44_ = _dafny.SeqWithoutIsStrInference([d_3053_v0_])
                d_3139_v45_: _dafny.Seq
                d_3139_v45_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_3138_v44_), d_3054_v1_])
                d_3140_v46_: _dafny.Map
                d_3140_v46_ = _dafny.Map({d_3137_v43_: default__.fm2((d_3139_v45_)[default__.safeIndex((self).f21, len(d_3139_v45_))], p1, globalState)})
                d_3140_v46_ = (d_3140_v46_).set((d_3137_v43_).set(p0, d_3097_cf10_), p0)
                d_3141_v47_: _dafny.Array
                nw484_ = _dafny.Array(int(0), 20)
                d_3141_v47_ = nw484_
                d_3141_v47_ = d_3141_v47_
            elif source40_.is_DC18:
                d_3142___mcc_h16_ = source40_.cf30
                d_3143___mcc_h17_ = source40_.cf31
                d_3144___mcc_h18_ = source40_.cf32
                d_3145___mcc_h19_ = source40_.cf33
                d_3146___mcc_h20_ = source40_.cf34
                d_3147_cf34_ = d_3146___mcc_h20_
                d_3148_cf33_ = d_3145___mcc_h19_
                d_3149_cf32_ = d_3144___mcc_h18_
                d_3150_cf31_ = d_3143___mcc_h17_
                d_3151_cf30_ = d_3142___mcc_h16_
                d_3152_v48_: _dafny.Array
                def lambda155_(d_3153_i5_):
                    return default__.safeDivisionInt(d_3153_i5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvwesnuh"))))

                init79_ = lambda155_
                nw485_ = _dafny.Array(None, 26)
                for i0_79_ in range(nw485_.length(0)):
                    nw485_[i0_79_] = init79_(i0_79_)
                d_3152_v48_ = nw485_
                d_3154_v49_: _dafny.Map
                d_3154_v49_ = _dafny.Map({d_3152_v48_: d_3053_v0_})
                d_3154_v49_ = ((d_3154_v49_) | (d_3154_v49_)).set(d_3152_v48_, d_3053_v0_)
                d_3155_v50_: _dafny.Seq
                d_3155_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                d_3156_v51_: _dafny.MultiSet
                d_3156_v51_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjxvjj"))).set(default__.safeIndex(d_3151_cf30_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjxvjj")))), d_3053_v0_)])
                (globalState).f17 = ((d_3156_v51_ if (d_3053_v0_) not in (d_3155_v50_) else _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_3157_i6_ in range(default__.abs(891))]), d_3155_v50_, (d_3155_v50_).set(default__.safeIndex(len(d_3056_v3_), len(d_3155_v50_)), d_3053_v0_), d_3155_v50_]))).cardinality
                d_3158_v52_: _dafny.Map
                d_3158_v52_ = _dafny.Map({d_3154_v49_: p0})
                d_3158_v52_ = (d_3158_v52_).set(d_3154_v49_, default__.fm2(d_3054_v1_, (d_3056_v3_)[default__.safeIndex((self).f21, len(d_3056_v3_))], globalState))
                d_3159_v53_: C3
                nw486_ = C3()
                nw486_.ctor__((self).f20)
                d_3159_v53_ = nw486_
            elif True:
                d_3160___mcc_h21_ = source40_.cf27
                d_3161_cf27_ = d_3160___mcc_h21_
                d_3098_cf9_ = d_3098_cf9_
                d_3162_v54_: _dafny.Array
                nw487_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_3162_v54_ = nw487_
                d_3163_v55_: _dafny.Map
                d_3163_v55_ = _dafny.Map({not(d_3098_cf9_): d_3162_v54_})
                d_3162_v54_ = (d_3162_v54_ if (d_3097_cf10_) in (_dafny.MultiSet([p0, d_3096_cf11_, d_3097_cf10_])) else ((d_3163_v55_)[d_3096_cf11_] if (d_3096_cf11_) in (d_3163_v55_) else d_3162_v54_))
                d_3164_v56_: _dafny.MultiSet
                d_3164_v56_ = _dafny.MultiSet([default__.fm2(d_3054_v1_, (self).f21, globalState), d_3096_cf11_])
                r0 = ((_dafny.MultiSet([p0, d_3097_cf10_, True, d_3097_cf10_, d_3097_cf10_])) - (d_3164_v56_) if p0 else default__.fm21(globalState))
                d_3165_v57_: D13
                d_3165_v57_ = D13_DC36(d_3099_cf8_)
                d_3166_v58_: _dafny.Array
                def lambda156_(d_3167_v3_):
                    def lambda157_(d_3168_i7_):
                        return d_3167_v3_

                    return lambda157_

                init80_ = lambda156_(d_3056_v3_)
                nw488_ = _dafny.Array(None, 15)
                for i0_80_ in range(nw488_.length(0)):
                    nw488_[i0_80_] = init80_(i0_80_)
                d_3166_v58_ = nw488_
                d_3169_v59_: D3
                d_3169_v59_ = D3_DC7(d_3166_v58_)
                d_3170_v60_: D3
                d_3170_v60_ = D3_DC9(d_3169_v59_)
                d_3171_v61_: _dafny.Seq
                d_3171_v61_ = _dafny.SeqWithoutIsStrInference([d_3169_v59_, D3_DC7(d_3166_v58_)])
                d_3172_v62_: D3
                d_3172_v62_ = D3_DC9((d_3171_v61_)[default__.safeIndex(d_3099_cf8_, len(d_3171_v61_))])
                d_3173_v63_: _dafny.Map
                d_3173_v63_ = _dafny.Map({(d_3165_v57_).cf68: d_3172_v62_})
                d_3174_v64_: _dafny.Set
                d_3174_v64_ = _dafny.Set({-775, (self).f21})
                d_3175_v65_: _dafny.Map
                d_3175_v65_ = _dafny.Map({p1: _dafny.Set({d_3099_cf8_})})
                d_3173_v63_ = (d_3173_v63_).set(len((_dafny.Map({d_3099_cf8_: d_3174_v64_})) | (d_3175_v65_)), d_3172_v62_)
            d_3176_v66_: _dafny.Array
            def lambda158_(d_3177_p1_):
                def lambda159_(d_3178_i8_):
                    return (d_3178_i8_) - (d_3177_p1_)

                return lambda159_

            init81_ = lambda158_(p1)
            nw489_ = _dafny.Array(None, 28)
            for i0_81_ in range(nw489_.length(0)):
                nw489_[i0_81_] = init81_(i0_81_)
            d_3176_v66_ = nw489_
            d_3179_v67_: _dafny.Array
            def lambda160_(d_3180_cf9_):
                def lambda161_(d_3181_i9_):
                    return D7_DC20(95, d_3180_cf9_)

                return lambda161_

            init82_ = lambda160_(d_3098_cf9_)
            nw490_ = _dafny.Array(None, 16)
            for i0_82_ in range(nw490_.length(0)):
                nw490_[i0_82_] = init82_(i0_82_)
            d_3179_v67_ = nw490_
            d_3182_v68_: D7
            d_3182_v68_ = D7_DC20(d_3099_cf8_, p0)
            index487_ = default__.safeIndex(346, (d_3179_v67_).length(0))
            (d_3179_v67_)[index487_] = d_3182_v68_
            d_3183_v69_: _dafny.Seq
            d_3183_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysryk"))
            d_3184_v70_: _dafny.Set
            d_3184_v70_ = _dafny.Set({645})
            d_3185_v71_: _dafny.Set
            d_3185_v71_ = _dafny.Set({d_3184_v70_, d_3184_v70_, d_3184_v70_})
            index488_ = default__.safeIndex(346, (d_3179_v67_).length(0))
            rhs471_ = d_3176_v66_
            rhs472_ = d_3182_v68_
            rhs473_ = ((D0_DC1(d_3053_v0_, p1, d_3100_v30_, d_3096_cf11_, d_3097_cf10_)).cf1) not in (d_3183_v69_)
            rhs474_ = default__.safeModuloInt((len(d_3185_v71_)) - ((0) - (p1)), d_3099_cf8_)
            lhs433_ = d_3179_v67_
            lhs434_ = default__.safeIndex(346, (d_3179_v67_).length(0))
            d_3176_v66_ = rhs471_
            lhs433_[lhs434_] = rhs472_
            d_3098_cf9_ = rhs473_
            r1 = rhs474_
            (globalState).f16 = (d_3183_v69_) + ((d_3183_v69_) + (d_3183_v69_))
        elif source38_.is_DC5:
            d_3186___mcc_h4_ = source38_.cf12
            d_3187___mcc_h5_ = source38_.cf13
            d_3188_cf13_ = d_3187___mcc_h5_
            d_3189_cf12_ = d_3186___mcc_h4_
            d_3190_v72_: _dafny.Array
            def lambda162_(d_3191_p1_):
                def lambda163_(d_3192_i10_):
                    return (d_3192_i10_) + ((0) - ((0) - (d_3191_p1_)))

                return lambda163_

            init83_ = lambda162_(p1)
            nw491_ = _dafny.Array(None, 11)
            for i0_83_ in range(nw491_.length(0)):
                nw491_[i0_83_] = init83_(i0_83_)
            d_3190_v72_ = nw491_
            d_3193_v73_: _dafny.Seq
            d_3193_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvco"))
            index489_ = default__.safeIndex(867, (d_3190_v72_).length(0))
            (d_3190_v72_)[index489_] = len(d_3193_v73_)
            d_3194_v74_: D10
            d_3194_v74_ = D10_DC28(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhgaqtm")), p1, p0, d_3189_cf12_, p0)
            d_3195_v75_: _dafny.Array
            nw492_ = _dafny.Array(None, 26)
            nw492_[int(0)] = (default__.fm46(p1, True, p0, globalState)) + (d_3193_v73_)
            nw492_[int(1)] = ((d_3194_v74_).cf52) + (d_3193_v73_)
            nw492_[int(2)] = d_3193_v73_
            nw492_[int(3)] = d_3193_v73_
            nw492_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upymoh"))) + ((d_3193_v73_).set(default__.safeIndex(d_3188_cf13_, len(d_3193_v73_)), d_3053_v0_))
            nw492_[int(5)] = d_3193_v73_
            nw492_[int(6)] = (d_3193_v73_).set(default__.safeIndex(len(d_3056_v3_), len(d_3193_v73_)), d_3053_v0_)
            nw492_[int(7)] = (d_3193_v73_) + (d_3193_v73_)
            nw492_[int(8)] = d_3193_v73_
            nw492_[int(9)] = _dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3196_i11_ in range(default__.abs(-52))])
            nw492_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vr"))) + (d_3193_v73_)
            nw492_[int(11)] = (d_3193_v73_) + (d_3193_v73_)
            nw492_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prlfknx"))
            nw492_[int(13)] = d_3193_v73_
            nw492_[int(14)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gotsytxmo"))) + (_dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3197_i12_ in range(default__.abs(-956))]))
            nw492_[int(15)] = d_3193_v73_
            nw492_[int(16)] = d_3193_v73_
            nw492_[int(17)] = d_3193_v73_
            nw492_[int(18)] = d_3193_v73_
            nw492_[int(19)] = d_3193_v73_
            nw492_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adw"))
            nw492_[int(21)] = d_3193_v73_
            nw492_[int(22)] = d_3193_v73_
            nw492_[int(23)] = default__.fm11(p0, globalState)
            nw492_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umd"))
            nw492_[int(25)] = _dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3198_i13_ in range(default__.abs(91))])
            d_3195_v75_ = nw492_
            d_3199_v76_: _dafny.Seq
            d_3199_v76_ = _dafny.SeqWithoutIsStrInference([d_3195_v75_])
            index490_ = default__.safeIndex(867, (d_3190_v72_).length(0))
            rhs475_ = (d_3056_v3_)[default__.safeIndex(p1, len(d_3056_v3_))]
            rhs476_ = (self).f21
            rhs477_ = (d_3199_v76_)[default__.safeIndex(p1, len(d_3199_v76_))]
            lhs435_ = d_3190_v72_
            lhs436_ = default__.safeIndex(867, (d_3190_v72_).length(0))
            lhs437_ = globalState
            lhs435_[lhs436_] = rhs475_
            lhs437_.f18 = rhs476_
            d_3195_v75_ = rhs477_
            d_3200_v77_: C9
            nw493_ = C9()
            nw493_.ctor__(((self).f20 if p0 else (self).f20))
            d_3200_v77_ = nw493_
            d_3201_v78_: _dafny.Array
            def lambda164_(d_3202_cf12_, d_3203_v72_, d_3204_p0_):
                def lambda165_(d_3205_i14_):
                    return (_dafny.Map({d_3202_cf12_: (d_3203_v72_)[default__.safeIndex(867, (d_3203_v72_).length(0))]})) | (_dafny.Map({d_3204_p0_: len(_dafny.Map({(self).f21: d_3204_p0_}))}))

                return lambda165_

            init84_ = lambda164_(d_3189_cf12_, d_3190_v72_, p0)
            nw494_ = _dafny.Array(None, 9)
            for i0_84_ in range(nw494_.length(0)):
                nw494_[i0_84_] = init84_(i0_84_)
            d_3201_v78_ = nw494_
            d_3206_v79_: _dafny.Seq
            d_3206_v79_ = _dafny.SeqWithoutIsStrInference([not(d_3189_cf12_)])
            rhs478_ = len((d_3056_v3_) + (_dafny.SeqWithoutIsStrInference([p1])))
            rhs479_ = (d_3206_v79_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trucpo"))), len(d_3206_v79_))]
            rhs480_ = d_3201_v78_
            rhs481_ = (default__.safeDivisionInt(d_3188_cf13_, d_3188_cf13_)) > ((d_3056_v3_)[default__.safeIndex((d_3190_v72_)[default__.safeIndex(867, (d_3190_v72_).length(0))], len(d_3056_v3_))])
            lhs438_ = globalState
            lhs439_ = globalState
            lhs438_.f18 = rhs478_
            lhs439_.f5 = rhs479_
            d_3201_v78_ = rhs480_
            d_3189_cf12_ = rhs481_
            (globalState).f18 = default__.safeDivisionInt(156, d_3188_cf13_)
        elif source38_.is_DC3:
            d_3207___mcc_h6_ = source38_.cf7
            d_3208_cf7_ = d_3207___mcc_h6_
            d_3209_v80_: _dafny.Map
            d_3209_v80_ = _dafny.Map({d_3053_v0_: p0})
            d_3210_v81_: _dafny.Map
            d_3210_v81_ = _dafny.Map({_dafny.Set({p1}): (d_3209_v80_) | ((d_3209_v80_).set(d_3053_v0_, p0))})
            def iife340_():
                coll82_ = _dafny.Set()
                compr_82_: int
                for compr_82_ in (d_3056_v3_).Elements:
                    d_3211_v82_: int = compr_82_
                    if (d_3211_v82_) in (d_3056_v3_):
                        def iife341_():
                            coll83_ = _dafny.Set()
                            compr_83_: str
                            for compr_83_ in (_dafny.MultiSet([_dafny.CodePoint('c')])).Elements:
                                d_3212_v83_: str = compr_83_
                                if (d_3212_v83_) in (_dafny.MultiSet([_dafny.CodePoint('c')])):
                                    coll83_ = coll83_.union(_dafny.Set([d_3212_v83_]))
                            return _dafny.Set(coll83_)
                        coll82_ = coll82_.union(_dafny.Set([default__.safeDivisionInt(d_3211_v82_, (0) - (len(iife341_()
)))]))
                return _dafny.Set(coll82_)
            d_3210_v81_ = (d_3210_v81_).set(iife340_()
            , d_3209_v80_)
            d_3213_v84_: C1
            nw495_ = C1()
            nw495_.ctor__(831, (self).f20)
            d_3213_v84_ = nw495_
            d_3214_v85_: _dafny.Array
            nw496_ = _dafny.Array(None, 27)
            nw496_[int(0)] = d_3213_v84_
            nw496_[int(1)] = d_3213_v84_
            nw496_[int(2)] = d_3213_v84_
            nw496_[int(3)] = (d_3213_v84_ if p0 else d_3213_v84_)
            nw496_[int(4)] = d_3213_v84_
            nw496_[int(5)] = d_3213_v84_
            nw496_[int(6)] = d_3213_v84_
            nw496_[int(7)] = d_3213_v84_
            nw496_[int(8)] = d_3213_v84_
            nw496_[int(9)] = d_3213_v84_
            nw496_[int(10)] = d_3213_v84_
            nw496_[int(11)] = d_3213_v84_
            nw496_[int(12)] = d_3213_v84_
            nw496_[int(13)] = d_3213_v84_
            nw496_[int(14)] = d_3213_v84_
            nw496_[int(15)] = d_3213_v84_
            nw496_[int(16)] = d_3213_v84_
            nw496_[int(17)] = d_3213_v84_
            nw496_[int(18)] = d_3213_v84_
            nw496_[int(19)] = d_3213_v84_
            nw496_[int(20)] = d_3213_v84_
            nw496_[int(21)] = d_3213_v84_
            nw496_[int(22)] = d_3213_v84_
            nw496_[int(23)] = d_3213_v84_
            nw496_[int(24)] = d_3213_v84_
            nw496_[int(25)] = d_3213_v84_
            nw496_[int(26)] = d_3213_v84_
            d_3214_v85_ = nw496_
            index491_ = default__.safeIndex(55, (d_3214_v85_).length(0))
            (d_3214_v85_)[index491_] = d_3213_v84_
            index492_ = default__.safeIndex(55, (d_3214_v85_).length(0))
            (d_3214_v85_)[index492_] = d_3213_v84_
            (globalState).f15 = default__.fm46(len(d_3208_cf7_), p0, p0, globalState)
            d_3215_v86_: _dafny.MultiSet
            d_3215_v86_ = _dafny.MultiSet([False, p0])
            (globalState).f6 = (d_3215_v86_).ispropersubset(d_3215_v86_)
        elif True:
            d_3216___mcc_h7_ = source38_.cf14
            d_3217_cf14_ = d_3216___mcc_h7_
            d_3218_v87_: _dafny.Map
            d_3218_v87_ = _dafny.Map({d_3055_v2_: (self).f21})
            d_3218_v87_ = (d_3218_v87_).set(_dafny.MultiSet([p1]), (self).f21)
            (globalState).f6 = p0
            d_3219_v88_: _dafny.Map
            d_3219_v88_ = _dafny.Map({(self).f21: d_3056_v3_})
            d_3220_v89_: _dafny.Set
            d_3220_v89_ = _dafny.Set({p1, (self).fm5(p0, globalState)})
            d_3056_v3_ = (((d_3219_v88_)[p1] if (p1) in (d_3219_v88_) else d_3056_v3_)) + (((d_3056_v3_).set(default__.safeIndex((self).f21, len(d_3056_v3_)), -913)) + ((d_3056_v3_).set(default__.safeIndex(len(d_3220_v89_), len(d_3056_v3_)), (self).f21)))
            rhs482_ = p0
            rhs483_ = ((self).f21) - (p1)
            rhs484_ = (0) - ((self).f21)
            rhs485_ = ((0) - (p1)) == ((0) - (p1))
            rhs486_ = ((0) - ((p1) + (((d_3055_v2_)[p1] if (p1) in (d_3055_v2_) else p1))) if default__.fm2(d_3054_v1_, 621, globalState) else ((self).f21) + ((self).f21))
            lhs440_ = globalState
            lhs441_ = globalState
            lhs442_ = globalState
            lhs443_ = globalState
            lhs440_.f5 = rhs482_
            lhs441_.f18 = rhs483_
            lhs442_.f18 = rhs484_
            lhs443_.f5 = rhs485_
            r1 = rhs486_
        d_3221_v90_: _dafny.MultiSet
        d_3221_v90_ = _dafny.MultiSet([p0, False])
        d_3222_v91_: _dafny.MultiSet
        d_3222_v91_ = d_3221_v90_
        d_3223_v92_: _dafny.Array
        nw497_ = _dafny.Array(_dafny.Seq({}), 11)
        d_3223_v92_ = nw497_
        d_3224_v93_: T0
        nw498_ = C5()
        nw498_.ctor__(d_3222_v91_, D3_DC7(d_3223_v92_), (self).f20)
        d_3224_v93_ = nw498_
        d_3225_v94_: _dafny.Seq
        d_3225_v94_ = _dafny.SeqWithoutIsStrInference([p0])
        d_3226_v95_: D6
        d_3226_v95_ = D6_DC18(len(_dafny.Set({p0, p0})), (self).f21, d_3225_v94_, p0, p1)
        (globalState).f17 = (default__.safeDivisionInt(len(_dafny.Map({(self).f21: d_3224_v93_})), (self).f21)) + ((len((d_3226_v95_).cf32)) - (p1))
        d_3227_i15_: int
        d_3227_i15_ = 0
        with _dafny.label("16"):
            while (d_3053_v0_) not in (_dafny.SeqWithoutIsStrInference([d_3053_v0_ for d_3234_i16_ in range(default__.abs(-272))])):
                with _dafny.c_label("16"):
                    if (d_3227_i15_) >= (100):
                        raise _dafny.Break("16")
                    d_3227_i15_ = (d_3227_i15_) + (1)
                    d_3228_v96_: _dafny.MultiSet
                    d_3228_v96_ = _dafny.MultiSet([_dafny.MultiSet([d_3053_v0_])])
                    pat_let_tv124_ = p0
                    d_3229_v97_: _dafny.Seq
                    def iife342_(_pat_let129_0):
                        def iife343_(d_3230_dt__update__tmp_h4_):
                            def iife344_(_pat_let130_0):
                                def iife345_(d_3231_dt__update_hcf0_h0_):
                                    return D0_DC0(d_3231_dt__update_hcf0_h0_)
                                return iife345_(_pat_let130_0)
                            return iife344_(pat_let_tv124_)
                        return iife343_(_pat_let129_0)
                    d_3229_v97_ = _dafny.SeqWithoutIsStrInference([(d_3224_v93_).f20, (self).f20, iife342_((d_3224_v93_).f20), (self).fm4(d_3053_v0_, d_3228_v96_, globalState), (self).fm4(d_3053_v0_, d_3228_v96_, globalState)])
                    d_3232_v98_: _dafny.Seq
                    d_3232_v98_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmdurf"))
                    rhs487_ = _dafny.SeqWithoutIsStrInference([D0_DC0(p0) for d_3233_i17_ in range(default__.abs(13))])
                    rhs488_ = p0
                    rhs489_ = p1
                    rhs490_ = (d_3056_v3_) == ((d_3056_v3_) + ((d_3056_v3_).set(default__.safeIndex((self).f21, len(d_3056_v3_)), (self).fm5(p0, globalState))))
                    rhs491_ = ((self).f21) * ((len(d_3232_v98_)) + (p1))
                    lhs444_ = globalState
                    lhs445_ = globalState
                    lhs446_ = globalState
                    lhs447_ = globalState
                    d_3229_v97_ = rhs487_
                    lhs444_.f6 = rhs488_
                    lhs445_.f18 = rhs489_
                    lhs446_.f5 = rhs490_
                    lhs447_.f17 = rhs491_
                    (globalState).f18 = default__.safeModuloInt(-960, (self).f21)
                    (globalState).f18 = ((self).f21) + (p1)
                    (globalState).f18 = (self).f21
                    pass
            pass
        d_3235_v99_: _dafny.Array
        def lambda166_(d_3236_p1_):
            def lambda167_(d_3237_i18_):
                return (d_3237_i18_) * (d_3236_p1_)

            return lambda167_

        init85_ = lambda166_(p1)
        nw499_ = _dafny.Array(None, 20)
        for i0_85_ in range(nw499_.length(0)):
            nw499_[i0_85_] = init85_(i0_85_)
        d_3235_v99_ = nw499_
        index493_ = default__.safeIndex(12, (d_3235_v99_).length(0))
        (d_3235_v99_)[index493_] = 908
        index494_ = default__.safeIndex(12, (d_3235_v99_).length(0))
        rhs492_ = ((self).f21) * (p1)
        rhs493_ = (self).fm5((d_3055_v2_) == (d_3055_v2_), globalState)
        rhs494_ = (d_3221_v90_).ispropersubset(d_3221_v90_)
        rhs495_ = p1
        rhs496_ = d_3221_v90_
        lhs448_ = globalState
        lhs449_ = globalState
        lhs450_ = globalState
        lhs451_ = d_3235_v99_
        lhs452_ = default__.safeIndex(12, (d_3235_v99_).length(0))
        lhs448_.f18 = rhs492_
        lhs449_.f17 = rhs493_
        lhs450_.f5 = rhs494_
        lhs451_[lhs452_] = rhs495_
        r0 = rhs496_
        d_3238_v100_: _dafny.Array
        def lambda168_(d_3239_p0_):
            def lambda169_(d_3240_i19_):
                return d_3239_p0_

            return lambda169_

        init86_ = lambda168_(p0)
        nw500_ = _dafny.Array(None, 6)
        for i0_86_ in range(nw500_.length(0)):
            nw500_[i0_86_] = init86_(i0_86_)
        d_3238_v100_ = nw500_
        d_3241_v101_: C0
        nw501_ = C0()
        nw501_.ctor__(d_3238_v100_, d_3225_v94_)
        d_3241_v101_ = nw501_
        d_3242_v102_: _dafny.Array
        nw502_ = _dafny.Array(None, 19)
        nw502_[int(0)] = d_3241_v101_
        nw502_[int(1)] = d_3241_v101_
        nw502_[int(2)] = d_3241_v101_
        nw502_[int(3)] = d_3241_v101_
        nw502_[int(4)] = d_3241_v101_
        nw502_[int(5)] = d_3241_v101_
        nw502_[int(6)] = d_3241_v101_
        nw502_[int(7)] = d_3241_v101_
        nw502_[int(8)] = d_3241_v101_
        nw502_[int(9)] = d_3241_v101_
        nw502_[int(10)] = d_3241_v101_
        nw502_[int(11)] = d_3241_v101_
        nw502_[int(12)] = d_3241_v101_
        nw502_[int(13)] = d_3241_v101_
        nw502_[int(14)] = d_3241_v101_
        nw502_[int(15)] = d_3241_v101_
        nw502_[int(16)] = d_3241_v101_
        nw502_[int(17)] = d_3241_v101_
        nw502_[int(18)] = d_3241_v101_
        d_3242_v102_ = nw502_
        d_3243_v103_: D22
        d_3243_v103_ = D22_DC59(d_3242_v102_)
        d_3244_v104_: _dafny.Array
        nw503_ = _dafny.Array(None, 26)
        nw503_[int(0)] = d_3242_v102_
        nw503_[int(1)] = d_3242_v102_
        nw503_[int(2)] = d_3242_v102_
        nw503_[int(3)] = d_3242_v102_
        nw503_[int(4)] = d_3242_v102_
        nw503_[int(5)] = d_3242_v102_
        nw503_[int(6)] = d_3242_v102_
        nw503_[int(7)] = d_3242_v102_
        nw503_[int(8)] = d_3242_v102_
        nw503_[int(9)] = (d_3243_v103_).cf110
        nw503_[int(10)] = d_3242_v102_
        nw503_[int(11)] = d_3242_v102_
        nw503_[int(12)] = d_3242_v102_
        nw503_[int(13)] = d_3242_v102_
        nw503_[int(14)] = d_3242_v102_
        nw503_[int(15)] = d_3242_v102_
        nw503_[int(16)] = d_3242_v102_
        nw503_[int(17)] = d_3242_v102_
        nw503_[int(18)] = d_3242_v102_
        nw503_[int(19)] = d_3242_v102_
        nw503_[int(20)] = d_3242_v102_
        nw503_[int(21)] = d_3242_v102_
        nw503_[int(22)] = d_3242_v102_
        nw503_[int(23)] = d_3242_v102_
        nw503_[int(24)] = d_3242_v102_
        nw503_[int(25)] = d_3242_v102_
        d_3244_v104_ = nw503_
        d_3245_v105_: D20
        d_3245_v105_ = D20_DC53(d_3244_v104_)
        d_3245_v105_ = d_3245_v105_
        r0 = d_3221_v90_
        d_3246_v106_: _dafny.Map
        d_3246_v106_ = _dafny.Map({(self).f21: p0})
        r1 = (default__.safeDivisionInt(len(d_3246_v106_), (d_3235_v99_)[default__.safeIndex(12, (d_3235_v99_).length(0))])) * ((self).f21)
        r2 = d_3053_v0_
        return r0, r1, r2

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        d_3247_i0_: int
        d_3247_i0_ = 0
        with _dafny.label("17"):
            while True:
                with _dafny.c_label("17"):
                    if (d_3247_i0_) >= (100):
                        raise _dafny.Break("17")
                    d_3247_i0_ = (d_3247_i0_) + (1)
                    (globalState).f18 = ((self).f21) + ((self).f21)
                    d_3248_v0_: C6
                    nw504_ = C6()
                    nw504_.ctor__((self).f20)
                    d_3248_v0_ = nw504_
                    d_3249_v1_: _dafny.Array
                    nw505_ = _dafny.Array(False, 27)
                    d_3249_v1_ = nw505_
                    index495_ = default__.safeIndex(635, (d_3249_v1_).length(0))
                    (d_3249_v1_)[index495_] = True
                    index496_ = default__.safeIndex(635, (d_3249_v1_).length(0))
                    rhs497_ = (self).fm5((p1) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_3250_i1_ in range(default__.abs(-564))])), globalState)
                    rhs498_ = p2
                    lhs453_ = globalState
                    lhs454_ = d_3249_v1_
                    lhs455_ = default__.safeIndex(635, (d_3249_v1_).length(0))
                    lhs453_.f17 = rhs497_
                    lhs454_[lhs455_] = rhs498_
                    d_3251_v2_: _dafny.Map
                    d_3251_v2_ = _dafny.Map({(self).f21: p1})
                    d_3252_v3_: _dafny.Set
                    d_3252_v3_ = _dafny.Set({not(p0)})
                    d_3253_v4_: _dafny.MultiSet
                    d_3253_v4_ = _dafny.MultiSet([d_3252_v3_, d_3252_v3_])
                    d_3254_v5_: _dafny.Seq
                    d_3254_v5_ = _dafny.SeqWithoutIsStrInference([d_3252_v3_, d_3252_v3_])
                    d_3255_v6_: str
                    d_3255_v6_ = _dafny.CodePoint('c')
                    d_3256_v7_: _dafny.Map
                    d_3256_v7_ = _dafny.Map({d_3255_v6_: d_3253_v4_})
                    d_3257_v8_: _dafny.Array
                    nw506_ = _dafny.Array(None, 13)
                    nw506_[int(0)] = d_3253_v4_
                    nw506_[int(1)] = d_3253_v4_
                    nw506_[int(2)] = default__.fm79(p2, globalState)
                    nw506_[int(3)] = (d_3253_v4_) - (d_3253_v4_)
                    nw506_[int(4)] = _dafny.MultiSet(d_3254_v5_)
                    nw506_[int(5)] = d_3253_v4_
                    nw506_[int(6)] = d_3253_v4_
                    nw506_[int(7)] = _dafny.MultiSet(d_3254_v5_)
                    nw506_[int(8)] = ((d_3256_v7_)[d_3255_v6_] if (d_3255_v6_) in (d_3256_v7_) else d_3253_v4_)
                    nw506_[int(9)] = d_3253_v4_
                    nw506_[int(10)] = d_3253_v4_
                    nw506_[int(11)] = d_3253_v4_
                    nw506_[int(12)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_3252_v3_]))).intersection(d_3253_v4_)
                    d_3257_v8_ = nw506_
                    index497_ = default__.safeIndex(986, (d_3257_v8_).length(0))
                    (d_3257_v8_)[index497_] = d_3253_v4_
                    d_3258_v9_: _dafny.Seq
                    d_3258_v9_ = _dafny.SeqWithoutIsStrInference([p0, True, p2, p2])
                    d_3259_v10_: _dafny.MultiSet
                    d_3259_v10_ = _dafny.MultiSet([(d_3249_v1_)[default__.safeIndex(635, (d_3249_v1_).length(0))]])
                    index498_ = default__.safeIndex(986, (d_3257_v8_).length(0))
                    rhs499_ = d_3251_v2_
                    rhs500_ = (default__.fm79(False, globalState)).intersection(d_3253_v4_)
                    rhs501_ = (not(p2)) or ((_dafny.MultiSet(d_3258_v9_)).isdisjoint(d_3259_v10_))
                    rhs502_ = d_3255_v6_
                    lhs456_ = d_3257_v8_
                    lhs457_ = default__.safeIndex(986, (d_3257_v8_).length(0))
                    lhs458_ = globalState
                    d_3251_v2_ = rhs499_
                    lhs456_[lhs457_] = rhs500_
                    lhs458_.f5 = rhs501_
                    d_3255_v6_ = rhs502_
                    pass
            pass
        (globalState).f17 = ((self).f21) + ((self).f21)
        d_3260_v11_: D13
        d_3260_v11_ = D13_DC35(_dafny.Set({(self).f21}))
        d_3260_v11_ = default__.fm80(p0, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuiaynak")), globalState)
        d_3261_v12_: C4
        nw507_ = C4()
        nw507_.ctor__((self).f20)
        d_3261_v12_ = nw507_
        d_3262_v13_: _dafny.Seq
        d_3262_v13_ = _dafny.SeqWithoutIsStrInference([p0])
        d_3262_v13_ = (d_3262_v13_) + ((d_3262_v13_) + (d_3262_v13_))
        (globalState).f18 = (self).f21
        d_3263_v14_: D2
        d_3263_v14_ = D2_DC5(p2, (self).f21)
        pat_let_tv125_ = p0
        pat_let_tv126_ = p2
        pat_let_tv127_ = p2
        pat_let_tv128_ = p1
        def lambda170_(source41_):
            if source41_.is_DC4:
                d_3264___mcc_h0_ = source41_.cf8
                d_3265___mcc_h1_ = source41_.cf9
                d_3266___mcc_h2_ = source41_.cf10
                d_3267___mcc_h3_ = source41_.cf11
                d_3268_cf11_ = d_3267___mcc_h3_
                d_3269_cf10_ = d_3266___mcc_h2_
                d_3270_cf9_ = d_3265___mcc_h1_
                d_3271_cf8_ = d_3264___mcc_h0_
                return (self).f21
            elif source41_.is_DC5:
                d_3272___mcc_h4_ = source41_.cf12
                d_3273___mcc_h5_ = source41_.cf13
                d_3274_cf13_ = d_3273___mcc_h5_
                d_3275_cf12_ = d_3272___mcc_h4_
                return (self).f21
            elif source41_.is_DC3:
                d_3276___mcc_h6_ = source41_.cf7
                d_3277_cf7_ = d_3276___mcc_h6_
                def iife346_():
                    coll84_ = _dafny.Map()
                    compr_84_: int
                    for compr_84_ in (_dafny.Set({-631, 119})).Elements:
                        d_3278_v15_: int = compr_84_
                        if (d_3278_v15_) in (_dafny.Set({-631, 119})):
                            coll84_[(d_3278_v15_) * ((self).f21)] = pat_let_tv125_
                    return _dafny.Map(coll84_)
                def iife347_():
                    def iife349_():
                        coll87_ = _dafny.Map()
                        compr_87_: int
                        for compr_87_ in (_dafny.SeqWithoutIsStrInference([(self).f21])).Elements:
                            d_3279_v17_: int = compr_87_
                            if (d_3279_v17_) in (_dafny.SeqWithoutIsStrInference([(self).f21])):
                                coll87_[(d_3279_v17_) + ((self).f21)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_3280_i2_ in range(default__.abs(154))]))
                        return _dafny.Map(coll87_)
                    coll85_ = _dafny.Map()
                    def iife348_():
                        coll86_ = _dafny.Map()
                        compr_86_: int
                        for compr_86_ in (_dafny.SeqWithoutIsStrInference([(self).f21])).Elements:
                            d_3279_v17_: int = compr_86_
                            if (d_3279_v17_) in (_dafny.SeqWithoutIsStrInference([(self).f21])):
                                coll86_[(d_3279_v17_) + ((self).f21)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_3280_i2_ in range(default__.abs(154))]))
                        return _dafny.Map(coll86_)
                    compr_85_: int
                    for compr_85_ in (iife348_()
                    ).keys.Elements:
                        d_3281_v16_: int = compr_85_
                        if (d_3281_v16_) in (iife349_()
                        ):
                            coll85_[(d_3281_v16_) - (len(pat_let_tv128_))] = pat_let_tv127_
                    return _dafny.Map(coll85_)
                return (0) - (len((iife346_()
                 if pat_let_tv126_ else iife347_()
                )))
            elif True:
                d_3282___mcc_h7_ = source41_.cf14
                d_3283_cf14_ = d_3282___mcc_h7_
                return (self).f21

        r0 = (0) - (lambda170_(d_3263_v14_))
        return r0

