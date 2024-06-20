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
    def fm2(globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([False, True]))

    @staticmethod
    def fm6(globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfyngxuaa"))

    @staticmethod
    def fm7(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: _dafny.Seq
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crncpbwhy")) for d_0_i0_ in range(default__.abs(155))])).Elements:
                d_1_v0_: _dafny.Seq = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "crncpbwhy")) for d_0_i0_ in range(default__.abs(155))])):
                    coll0_[d_1_v0_] = 275
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([754, -590, 382, len(_dafny.Map({False: _dafny.CodePoint('t')})), len(_dafny.Set({len(_dafny.Map({True: True}))}))]))])).Elements:
                d_2_v1_: int = compr_1_
                if (d_2_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([754, -590, 382, len(_dafny.Map({False: _dafny.CodePoint('t')})), len(_dafny.Set({len(_dafny.Map({True: True}))}))]))])):
                    coll1_[default__.safeDivisionInt(d_2_v1_, -199)] = False
            return _dafny.Map(coll1_)
        return len(_dafny.SeqWithoutIsStrInference([D6_DC24(D6_DC21(_dafny.MultiSet([198]))), D6_DC24(D6_DC21(_dafny.MultiSet([len(iife0_()
)]))), D6_DC24(D6_DC22(-556)), D6_DC24(D6_DC21(_dafny.MultiSet([245, len(_dafny.SeqWithoutIsStrInference([iife1_()
]))])))]))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: str
            for compr_2_ in (_dafny.Set({_dafny.CodePoint('x'), _dafny.CodePoint('u')})).Elements:
                d_3_v0_: str = compr_2_
                if (d_3_v0_) in (_dafny.Set({_dafny.CodePoint('x'), _dafny.CodePoint('u')})):
                    coll2_ = coll2_.union(_dafny.Set([d_3_v0_]))
            return _dafny.Set(coll2_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True])): True})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not(False), False])).cardinality])): True}))) | (_dafny.Map({(0) - (len(iife2_()
        )): True}))

    @staticmethod
    def fm9(globalState):
        if (not(True)) or (True):
            return _dafny.SeqWithoutIsStrInference([276])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, False])) for d_4_i0_ in range(default__.abs(142))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([44]))})) for d_5_i1_ in range(default__.abs(571))]))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({not(False): False}))) | ((_dafny.Map({False: False})) | (_dafny.Map({True: not(False)})))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return ((_dafny.Set({True})) | (_dafny.Set({not(False)}))).intersection(_dafny.Set({False, True, True}))

    @staticmethod
    def fm12(p0, p1, globalState):
        return (D10_DC34(_dafny.CodePoint('h'), False, False, True, False)).cf55

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([547])

    @staticmethod
    def fm15(p0, globalState):
        return False

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_6_i0_ in range(default__.abs(291))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luposje")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_7_i1_ in range(default__.abs(624))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yiwbhd"))))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_8_i0_ in range(default__.abs(990))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "em"))])).Elements:
                d_9_v0_: _dafny.Seq = compr_3_
                if (d_9_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_8_i0_ in range(default__.abs(990))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "em"))])):
                    coll3_[d_9_v0_] = 322
            return _dafny.Map(coll3_)
        return (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thuvag")): -296})) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njhdgt")): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgehlao"))))})) | (iife3_()
        ))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        source0_ = (D6_DC23(True, -860, not(True)) if True else D6_DC23(True, 663, True))
        if source0_.is_DC22:
            d_10___mcc_h0_ = source0_.cf36
            d_11_cf36_ = d_10___mcc_h0_
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(180, 854):
                    d_12_v0_: int = compr_4_
                    if ((180) <= (d_12_v0_)) and ((d_12_v0_) < (854)):
                        coll4_[default__.safeModuloInt(d_12_v0_, d_11_cf36_)] = False
                return _dafny.Map(coll4_)
            return (_dafny.Map({d_11_cf36_: True})) | (iife4_()
            )
        elif source0_.is_DC23:
            d_13___mcc_h1_ = source0_.cf37
            d_14___mcc_h2_ = source0_.cf38
            d_15___mcc_h3_ = source0_.cf39
            d_16_cf39_ = d_15___mcc_h3_
            d_17_cf38_ = d_14___mcc_h2_
            d_18_cf37_ = d_13___mcc_h1_
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in (_dafny.Map({d_17_cf38_: d_18_cf37_})).keys.Elements:
                    d_19_v1_: int = compr_5_
                    if (d_19_v1_) in (_dafny.Map({d_17_cf38_: d_18_cf37_})):
                        coll5_[(d_19_v1_) * (d_17_cf38_)] = d_18_cf37_
                return _dafny.Map(coll5_)
            return (iife5_()
            ) | (_dafny.Map({d_17_cf38_: d_18_cf37_}))
        elif source0_.is_DC21:
            d_20___mcc_h4_ = source0_.cf35
            d_21_cf35_ = d_20___mcc_h4_
            def iife6_():
                def iife8_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in (d_21_cf35_).Elements:
                        d_24_v3_: int = compr_8_
                        if (d_24_v3_) in (d_21_cf35_):
                            coll8_ = coll8_.union(_dafny.Set([(d_24_v3_) - ((_dafny.MultiSet([True, True, True])).cardinality)]))
                    return _dafny.Set(coll8_)
                coll6_ = _dafny.Map()
                def iife7_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in (d_21_cf35_).Elements:
                        d_22_v3_: int = compr_7_
                        if (d_22_v3_) in (d_21_cf35_):
                            coll7_ = coll7_.union(_dafny.Set([(d_22_v3_) - ((_dafny.MultiSet([True, True, True])).cardinality)]))
                    return _dafny.Set(coll7_)
                compr_6_: int
                for compr_6_ in (iife7_()
                ).Elements:
                    d_23_v2_: int = compr_6_
                    if (d_23_v2_) in (iife8_()
                    ):
                        coll6_[default__.safeModuloInt(d_23_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udnqqbuso"))))] = not(not(False))
                return _dafny.Map(coll6_)
            return iife6_()
            
        elif True:
            d_25___mcc_h5_ = source0_.cf40
            d_26_cf40_ = d_25___mcc_h5_
            return (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([True]))): True}))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('k')

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqwo"))): False}))})).Elements:
                d_27_v0_: int = compr_9_
                if (d_27_v0_) in (_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqwo"))): False}))})):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_27_v0_, len(_dafny.Map({296: len(_dafny.Set({276, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([157, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvxtdfslh"))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlkshnmc")))})), 968, 660, 940]))).cardinality}))})))]))
            return _dafny.Set(coll9_)
        return ((_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xunsdlrj"))))])).cardinality]))})).intersection(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))}))) | ((iife9_()
        ) | (_dafny.Set({723})))

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_28_i0_ in range(default__.abs(614))])))])) + (_dafny.SeqWithoutIsStrInference([743, 349, 228, -108]))) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gw"))), -30])).cardinality]))

    @staticmethod
    def fm25(p0, p1, globalState):
        def iife10_():
            def iife14_():
                def iife16_():
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])).Elements:
                        d_29_v1_: int = compr_16_
                        if (d_29_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])):
                            coll16_[default__.safeDivisionInt(d_29_v1_, 890)] = True
                    return _dafny.Map(coll16_)
                coll14_ = _dafny.Map()
                def iife15_():
                    coll15_ = _dafny.Map()
                    compr_15_: int
                    for compr_15_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])).Elements:
                        d_29_v1_: int = compr_15_
                        if (d_29_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])):
                            coll15_[default__.safeDivisionInt(d_29_v1_, 890)] = True
                    return _dafny.Map(coll15_)
                compr_14_: int
                for compr_14_ in (iife15_()
                ).keys.Elements:
                    d_30_v0_: int = compr_14_
                    if (d_30_v0_) in (iife16_()
                    ):
                        coll14_[default__.safeModuloInt(d_30_v0_, len(_dafny.Map({True: _dafny.Set({(0) - (len(_dafny.Map({_dafny.CodePoint('f'): False})))})})))] = len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))]))
                return _dafny.Map(coll14_)
            coll10_ = _dafny.Set()
            def iife11_():
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])).Elements:
                        d_29_v1_: int = compr_13_
                        if (d_29_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])):
                            coll13_[default__.safeDivisionInt(d_29_v1_, 890)] = True
                    return _dafny.Map(coll13_)
                coll11_ = _dafny.Map()
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: int
                    for compr_12_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])).Elements:
                        d_29_v1_: int = compr_12_
                        if (d_29_v1_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({67: 98}))), 587])):
                            coll12_[default__.safeDivisionInt(d_29_v1_, 890)] = True
                    return _dafny.Map(coll12_)
                compr_11_: int
                for compr_11_ in (iife12_()
                ).keys.Elements:
                    d_30_v0_: int = compr_11_
                    if (d_30_v0_) in (iife13_()
                    ):
                        coll11_[default__.safeModuloInt(d_30_v0_, len(_dafny.Map({True: _dafny.Set({(0) - (len(_dafny.Map({_dafny.CodePoint('f'): False})))})})))] = len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))]))
                return _dafny.Map(coll11_)
            compr_10_: D3
            for compr_10_ in (_dafny.Map({D3_DC10(iife11_()
): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_31_i0_ in range(default__.abs(992))]))})).keys.Elements:
                d_32_v2_: D3 = compr_10_
                if (d_32_v2_) in (_dafny.Map({D3_DC10(iife14_()
): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_31_i0_ in range(default__.abs(992))]))})):
                    coll10_ = coll10_.union(_dafny.Set([d_32_v2_]))
            return _dafny.Set(coll10_)
        return (iife10_()
        ) | (_dafny.Set({D3_DC10(_dafny.Map({575: 899}))}))

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([466, 203]), _dafny.SeqWithoutIsStrInference([107, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udohgqy")))]), (_dafny.SeqWithoutIsStrInference([591]) if False else _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_33_i0_ in range(default__.abs(326))])), len(_dafny.SeqWithoutIsStrInference([True])), (D6_DC23(False, 174, False)).cf38, 942])), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-517 for d_34_i1_ in range(default__.abs(239))])), len(_dafny.Set({False}))]), (_dafny.SeqWithoutIsStrInference([855 for d_35_i2_ in range(default__.abs(685))])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mogyijfrl"))))]))})

    @staticmethod
    def fm27(p0, globalState):
        return (_dafny.Map({False: 553})) | ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf")) for d_36_i0_ in range(default__.abs(755))]))})) | (_dafny.Map({not(True): -295})))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        source1_ = D5_DC17(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eql")): (0) - ((_dafny.MultiSet([True, True, False, True, not(True)])).cardinality)}))
        if source1_.is_DC18:
            d_37___mcc_h0_ = source1_.cf30
            d_38___mcc_h1_ = source1_.cf31
            d_39_cf31_ = d_38___mcc_h1_
            d_40_cf30_ = d_37___mcc_h0_
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(936, 244):
                    d_41_v0_: int = compr_17_
                    if ((936) <= (d_41_v0_)) and ((d_41_v0_) < (244)):
                        coll17_[(d_41_v0_) + ((0) - (d_40_cf30_))] = d_40_cf30_
                return _dafny.Map(coll17_)
            return iife17_()
            
        elif source1_.is_DC19:
            d_42___mcc_h2_ = source1_.cf32
            d_43___mcc_h3_ = source1_.cf33
            d_44_cf33_ = d_43___mcc_h3_
            d_45_cf32_ = d_42___mcc_h2_
            return d_45_cf32_
        elif source1_.is_DC17:
            d_46___mcc_h4_ = source1_.cf29
            d_47_cf29_ = d_46___mcc_h4_
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(213, 935):
                    d_48_v1_: int = compr_18_
                    if ((213) <= (d_48_v1_)) and ((d_48_v1_) < (935)):
                        coll18_[(d_48_v1_) - (-704)] = 766
                return _dafny.Map(coll18_)
            return (D5_DC19(iife18_()
, not(False))).cf32
        elif True:
            d_49___mcc_h5_ = source1_.cf34
            d_50_cf34_ = d_49___mcc_h5_
            def iife19_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(-975, 332):
                    d_51_v2_: int = compr_19_
                    if ((-975) <= (d_51_v2_)) and ((d_51_v2_) < (332)):
                        coll19_[default__.safeDivisionInt(d_51_v2_, -974)] = -753
                return _dafny.Map(coll19_)
            return iife19_()
            

    @staticmethod
    def fm29(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([626]) if True else _dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([False])))]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-947])))

    @staticmethod
    def fm30(globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_52_i0_ in range(default__.abs(287))]) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wuqb")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfp")))

    @staticmethod
    def fm31(p0, globalState):
        if True:
            return D4_DC15(len(_dafny.Map({True: (D12_DC42(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nekfj"))))).cf69})), 237)
        elif True:
            return D4_DC15((0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cpwd"))): True}))), len(_dafny.Map({True: len(_dafny.Set({not(True)}))})))

    @staticmethod
    def fm32(globalState):
        return D7_DC25((_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhjvyhhbk")))]), _dafny.SeqWithoutIsStrInference([270]), _dafny.SeqWithoutIsStrInference([281, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujoxgmcix")))]), _dafny.SeqWithoutIsStrInference([240]), _dafny.SeqWithoutIsStrInference([-985])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([815, (_dafny.MultiSet([439])).cardinality]), _dafny.SeqWithoutIsStrInference([381, len(_dafny.SeqWithoutIsStrInference([not(False)]))])})))

    @staticmethod
    def fm33(globalState):
        return _dafny.Set({_dafny.CodePoint('c')})

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_53_i0_ in range(default__.abs(53))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsupnh")))

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        return not ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujybiveij"))) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvlkj")))) or ((True if True else True))

    @staticmethod
    def fm38(p0, p1, globalState):
        return _dafny.CodePoint('x')

    @staticmethod
    def fm39(p0, p1, globalState):
        return (D19_DC59(_dafny.Map({True: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))}))).cf89

    @staticmethod
    def fm40(p0, p1, p2, globalState):
        if ((_dafny.MultiSet([(0) - (-647), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_54_i0_ in range(default__.abs(442))])), len(_dafny.Map({-834: False})), 258])).cardinality) > (len(_dafny.SeqWithoutIsStrInference([False]))):
            return _dafny.Map({337: True})
        elif True:
            return _dafny.Map({-765: True})

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm42(globalState):
        return D5_DC17(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icyaksiiv")): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hf")))}))

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        return D11_DC36((_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([-360]) for d_55_i0_ in range(default__.abs(876))]) if True else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([-909])])))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({len(_dafny.Map({78: not(True)})): _dafny.CodePoint('y')})) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsq")))): _dafny.CodePoint('c')}))) | (_dafny.Map({(0) - (-871): _dafny.CodePoint('d')}))

    @staticmethod
    def fm45(p0, p1, globalState):
        return D2_DC6((424) * (-232))

    @staticmethod
    def fm46(globalState):
        if (False) and (False):
            return _dafny.MultiSet([D6_DC22(-148)])
        elif True:
            return _dafny.MultiSet([D6_DC22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mko")))), D6_DC22(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_56_i0_ in range(default__.abs(858))]))), D6_DC22((_dafny.MultiSet([False, False, True])).cardinality), D6_DC22(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xh")))), D6_DC22(936)])

    @staticmethod
    def fm47(p0, p1, globalState):
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: int
            for compr_20_ in _dafny.IntegerRange(329, -646):
                d_57_v0_: int = compr_20_
                if ((329) <= (d_57_v0_)) and ((d_57_v0_) < (-646)):
                    coll20_[default__.safeDivisionInt(d_57_v0_, len(_dafny.Map({False: not(True)})))] = -100
            return _dafny.Map(coll20_)
        return (_dafny.Map({True: _dafny.Map({len(_dafny.SeqWithoutIsStrInference([21, 456])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aridp")))})})) | (_dafny.Map({False: iife20_()
        }))

    @staticmethod
    def fm48(globalState):
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: _dafny.Seq
            for compr_21_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwx"))])).Elements:
                d_58_v0_: _dafny.Seq = compr_21_
                if (d_58_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "me")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwx"))])):
                    coll21_[d_58_v0_] = True
            return _dafny.Map(coll21_)
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(-807, 704):
                d_59_v1_: int = compr_22_
                if ((-807) <= (d_59_v1_)) and ((d_59_v1_) < (704)):
                    coll22_[default__.safeDivisionInt(d_59_v1_, 861)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_60_i0_ in range(default__.abs(3))]))
            return _dafny.Map(coll22_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, False, True, False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(D2_DC7(_dafny.SeqWithoutIsStrInference([(D2_DC6((0) - (-604))).cf9, len((D20_DC61(iife21_()
)).cf93)]), not(False), not(not(not(not(False)))), (_dafny.MultiSet([len(_dafny.Set({True, not(False)})), -983, len(_dafny.Map({not(True): True})), 480, len(iife22_()
)])).cardinality)).cf11]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True])]))

    @staticmethod
    def fm49(p0, globalState):
        return _dafny.Map({_dafny.CodePoint('c'): True})

    @staticmethod
    def fm50(p0, globalState):
        def lambda0_(source2_):
            if source2_.is_DC13:
                d_61___mcc_h0_ = source2_.cf22
                d_62_cf22_ = d_61___mcc_h0_
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])
            elif source2_.is_DC14:
                d_63___mcc_h1_ = source2_.cf23
                d_64___mcc_h2_ = source2_.cf24
                d_65___mcc_h3_ = source2_.cf25
                d_66_cf25_ = d_65___mcc_h3_
                d_67_cf24_ = d_64___mcc_h2_
                d_68_cf23_ = d_63___mcc_h1_
                return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False])])
            elif source2_.is_DC15:
                d_69___mcc_h4_ = source2_.cf26
                d_70___mcc_h5_ = source2_.cf27
                d_71_cf27_ = d_70___mcc_h5_
                d_72_cf26_ = d_69___mcc_h4_
                return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, (D15_DC51((D6_DC23(False, d_71_cf27_, True)).cf39, not(False))).cf83]), _dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True, True]), _dafny.SeqWithoutIsStrInference([False, True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_73_i0_ in range(default__.abs(565))]))
            elif source2_.is_DC12:
                d_74___mcc_h6_ = source2_.cf21
                d_75_cf21_ = d_74___mcc_h6_
                return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(False), False, not(False), True, True]), _dafny.SeqWithoutIsStrInference([True, not(True)]), _dafny.SeqWithoutIsStrInference([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False)])]))
            elif True:
                d_76___mcc_h7_ = source2_.cf28
                d_77_cf28_ = d_76___mcc_h7_
                return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, False, False, False])]))

        return _dafny.MultiSet(lambda0_(D4_DC15(len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([473, len(_dafny.Map({False: len(_dafny.Map({True: -197}))}))]))]))).cardinality, len(_dafny.SeqWithoutIsStrInference([False]))])), 50)))

    @staticmethod
    def fm51(globalState):
        if (not(not(True))) or (False):
            return D4_DC12(_dafny.Map({len(_dafny.Set({False})): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([933]))}))
        elif True:
            return D4_DC12(_dafny.Map({-648: _dafny.MultiSet([len(_dafny.Map({True: True})), 132])}))

    @staticmethod
    def fm52(globalState):
        return (((D13_DC44(_dafny.SeqWithoutIsStrInference([True]))).cf71) + (_dafny.SeqWithoutIsStrInference([True]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True])))

    @staticmethod
    def fm53(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_78_i1_ in range(default__.abs(296))]) for d_79_i0_ in range(default__.abs(879))])

    @staticmethod
    def fm54(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vslhotx")))]) for d_80_i0_ in range(default__.abs(14))]) if False else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctbyue")))]), _dafny.SeqWithoutIsStrInference([767 for d_81_i1_ in range(default__.abs(554))])]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-507]) for d_82_i2_ in range(default__.abs(754))]) if not(False) else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvii")))]), _dafny.SeqWithoutIsStrInference([-259, len(_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_83_i3_ in range(default__.abs(803))]))})]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqnhvr")))])])))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        def iife23_():
            coll23_ = _dafny.Set()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(73, 703):
                d_84_v0_: int = compr_23_
                if ((73) <= (d_84_v0_)) and ((d_84_v0_) < (703)):
                    coll23_ = coll23_.union(_dafny.Set([(d_84_v0_) * (697)]))
            return _dafny.Set(coll23_)
        if (iife23_()
        ).issubset(_dafny.Set({952, 130, len(_dafny.Map({True: False}))})):
            return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnfrahgri")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_85_i0_ in range(default__.abs(-769))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "as")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkaix")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujnwkx"))})
        elif True:
            return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_86_i1_ in range(default__.abs(175))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_87_i2_ in range(default__.abs(751))])})

    @staticmethod
    def fm56(p0, globalState):
        return D0_DC1(default__.safeDivisionInt(-605, -71), True, _dafny.Map({not(False): _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sg")))): (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxkfrh"))))})}), len(_dafny.SeqWithoutIsStrInference([True, False])), _dafny.CodePoint('k'))

    @staticmethod
    def Main(noArgsParameter__):
        d_88_v0_: _dafny.MultiSet
        d_88_v0_ = _dafny.MultiSet([False])
        d_89_v1_: _dafny.MultiSet
        d_89_v1_ = _dafny.MultiSet([d_88_v0_, d_88_v0_])
        d_90_v2_: bool
        d_90_v2_ = True
        d_91_v3_: _dafny.Seq
        d_91_v3_ = _dafny.SeqWithoutIsStrInference([d_90_v2_, d_90_v2_])
        d_92_v4_: int
        d_92_v4_ = 166
        d_93_v5_: _dafny.Seq
        d_93_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
        d_94_v6_: _dafny.MultiSet
        d_94_v6_ = _dafny.MultiSet([len(d_93_v5_), 476])
        d_95_v7_: _dafny.MultiSet
        d_95_v7_ = _dafny.MultiSet([(d_94_v6_).cardinality])
        d_96_v8_: _dafny.Seq
        d_96_v8_ = _dafny.SeqWithoutIsStrInference([d_95_v7_])
        d_97_v9_: _dafny.Seq
        d_97_v9_ = _dafny.SeqWithoutIsStrInference([(d_96_v8_)[default__.safeIndex(d_92_v4_, len(d_96_v8_))]])
        d_98_v10_: _dafny.Array
        def lambda1_(d_99_i0_):
            return True

        init0_ = lambda1_
        nw0_ = _dafny.Array(None, 10)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_98_v10_ = nw0_
        d_100_v11_: _dafny.Array
        nw1_ = _dafny.Array(None, 7)
        nw1_[int(0)] = (0) - ((0) - (d_92_v4_))
        nw1_[int(1)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_101_i1_ in range(default__.abs(656))]))
        nw1_[int(2)] = d_92_v4_
        nw1_[int(3)] = d_92_v4_
        nw1_[int(4)] = d_92_v4_
        nw1_[int(5)] = d_92_v4_
        nw1_[int(6)] = d_92_v4_
        d_100_v11_ = nw1_
        d_102_globalState_: GlobalState
        nw2_ = GlobalState()
        nw2_.ctor__(True, d_89_v1_, _dafny.SeqWithoutIsStrInference([d_90_v2_]), (d_91_v3_) + (d_91_v3_), False, _dafny.SeqWithoutIsStrInference([d_90_v2_]), 301, True, 794, 616, False, 333, 349, _dafny.Set({d_92_v4_}), 578, (d_97_v9_)[default__.safeIndex(d_92_v4_, len(d_97_v9_))], 982, -565, d_98_v10_, d_100_v11_, d_95_v7_, d_98_v10_)
        d_102_globalState_ = nw2_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_98_v10_).length(0)):
            d_103_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_103_i2_)) and ((d_103_i2_) < ((d_98_v10_).length(0)))):
                (d_98_v10_)[(d_103_i2_)] = d_90_v2_
        d_104_v12_: _dafny.Seq
        d_104_v12_ = _dafny.SeqWithoutIsStrInference([d_98_v10_, d_98_v10_])
        d_105_v13_: C4
        nw3_ = C4()
        nw3_.ctor__(d_90_v2_, d_104_v12_)
        d_105_v13_ = nw3_
        d_106_v14_: D2
        d_106_v14_ = D2_DC8(D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufigqlgw"))), d_100_v11_, d_90_v2_)
        d_107_v15_: str
        d_107_v15_ = _dafny.CodePoint('m')
        d_108_v16_: _dafny.Map
        d_108_v16_ = _dafny.Map({d_107_v15_: d_90_v2_})
        d_109_v17_: _dafny.Set
        d_109_v17_ = _dafny.Set({(d_106_v14_).cf16, ((d_108_v16_)[d_107_v15_] if (d_107_v15_) in (d_108_v16_) else d_90_v2_), d_90_v2_})
        d_110_v18_: _dafny.Seq
        d_110_v18_ = _dafny.SeqWithoutIsStrInference([-999, d_92_v4_])
        d_111_v19_: _dafny.Map
        d_111_v19_ = _dafny.Map({(d_105_v13_.f28) in (d_109_v17_): (d_105_v13_).fm13((d_110_v18_).set(default__.safeIndex(d_92_v4_, len(d_110_v18_)), 731), d_90_v2_, len(d_110_v18_), d_102_globalState_)})
        d_112_v20_: _dafny.Array
        nw4_ = _dafny.Array(None, 13)
        d_112_v20_ = nw4_
        d_113_v21_: D11
        d_113_v21_ = D11_DC37(d_112_v20_)
        d_114_v22_: _dafny.Map
        d_114_v22_ = _dafny.Map({d_113_v21_: d_90_v2_})
        d_111_v19_ = (d_111_v19_).set(((d_114_v22_)[d_113_v21_] if (d_113_v21_) in (d_114_v22_) else d_90_v2_), not((d_105_v13_.f28) and (d_90_v2_)))
        d_115_v23_: _dafny.Set
        d_115_v23_ = _dafny.Set({len(default__.fm6(d_102_globalState_)), d_92_v4_, d_92_v4_, d_92_v4_, d_92_v4_})
        d_116_v24_: D17
        d_116_v24_ = D17_DC55(d_115_v23_)
        pat_let_tv0_ = d_92_v4_
        pat_let_tv1_ = d_105_v13_
        pat_let_tv2_ = d_93_v5_
        pat_let_tv3_ = d_92_v4_
        pat_let_tv4_ = d_92_v4_
        pat_let_tv5_ = d_92_v4_
        def lambda2_(source4_):
            if source4_.is_DC56:
                d_117___mcc_h6_ = source4_.cf87
                d_118_cf87_ = d_117___mcc_h6_
                return D5_DC20(D5_DC18(pat_let_tv0_, 280))
            elif True:
                d_119___mcc_h7_ = source4_.cf86
                d_120_cf86_ = d_119___mcc_h7_
                if pat_let_tv1_.f28:
                    return D5_DC20(D5_DC17(_dafny.Map({pat_let_tv2_: pat_let_tv3_})))
                elif True:
                    return D5_DC20(D5_DC18(pat_let_tv4_, pat_let_tv5_))

        source3_ = lambda2_(d_116_v24_)
        if source3_.is_DC18:
            d_121___mcc_h0_ = source3_.cf30
            d_122___mcc_h1_ = source3_.cf31
            d_123_cf31_ = d_122___mcc_h1_
            d_124_cf30_ = d_121___mcc_h0_
            d_125_v25_: _dafny.Map
            d_125_v25_ = _dafny.Map({d_90_v2_: d_124_cf30_})
            d_125_v25_ = d_125_v25_
            d_126_v26_: C4
            nw5_ = C4()
            nw5_.ctor__(d_90_v2_, d_104_v12_)
            d_126_v26_ = nw5_
            d_127_v27_: _dafny.Array
            nw6_ = _dafny.Array(_dafny.Map({}), 28)
            d_127_v27_ = nw6_
            d_128_v28_: _dafny.Map
            d_128_v28_ = _dafny.Map({d_105_v13_.f28: d_127_v27_})
            d_129_v29_: D10
            d_129_v29_ = D10_DC34(d_107_v15_, True, d_105_v13_.f28, not(d_105_v13_.f28), d_126_v26_.f28)
            d_130_v30_: _dafny.Array
            nw7_ = _dafny.Array(None, 10)
            nw7_[int(0)] = ((d_128_v28_)[(d_129_v29_).cf57] if ((d_129_v29_).cf57) in (d_128_v28_) else d_127_v27_)
            nw7_[int(1)] = d_127_v27_
            nw7_[int(2)] = d_127_v27_
            nw7_[int(3)] = d_127_v27_
            nw7_[int(4)] = d_127_v27_
            nw7_[int(5)] = d_127_v27_
            nw7_[int(6)] = (d_127_v27_ if d_105_v13_.f28 else d_127_v27_)
            nw7_[int(7)] = d_127_v27_
            nw7_[int(8)] = d_127_v27_
            nw7_[int(9)] = (d_127_v27_ if d_90_v2_ else d_127_v27_)
            d_130_v30_ = nw7_
            index0_ = default__.safeIndex(706, (d_130_v30_).length(0))
            (d_130_v30_)[index0_] = d_127_v27_
            index1_ = default__.safeIndex(706, (d_130_v30_).length(0))
            (d_130_v30_)[index1_] = d_127_v27_
            d_131_v31_: D6
            d_131_v31_ = D6_DC22(590)
            d_132_v32_: D6
            d_132_v32_ = D6_DC24(d_131_v31_)
            d_133_v33_: _dafny.Map
            d_133_v33_ = _dafny.Map({d_100_v11_: d_132_v32_})
            d_133_v33_ = d_133_v33_
        elif source3_.is_DC19:
            d_134___mcc_h2_ = source3_.cf32
            d_135___mcc_h3_ = source3_.cf33
            d_136_cf33_ = d_135___mcc_h3_
            d_137_cf32_ = d_134___mcc_h2_
            (d_102_globalState_).f11 = d_92_v4_
            d_138_v34_: C5
            nw8_ = C5()
            nw8_.ctor__(_dafny.Map({d_92_v4_: d_105_v13_.f28}), d_92_v4_)
            d_138_v34_ = nw8_
            d_139_v35_: _dafny.Seq
            d_139_v35_ = _dafny.SeqWithoutIsStrInference([d_115_v23_])
            (d_102_globalState_).f11 = (d_92_v4_) - ((0) - ((d_105_v13_).fm0(d_95_v7_, len(_dafny.Set({(d_138_v34_).f26})), d_139_v35_, d_105_v13_.f28, d_102_globalState_)))
            d_140_v36_: _dafny.Array
            def lambda3_(d_141_v13_, d_142_v34_):
                def lambda4_(d_143_i3_):
                    return _dafny.Map({d_141_v13_.f28: (d_142_v34_).f27})

                return lambda4_

            init1_ = lambda3_(d_105_v13_, d_138_v34_)
            nw9_ = _dafny.Array(None, 10)
            for i0_1_ in range(nw9_.length(0)):
                nw9_[i0_1_] = init1_(i0_1_)
            d_140_v36_ = nw9_
            d_140_v36_ = d_140_v36_
        elif source3_.is_DC17:
            d_144___mcc_h4_ = source3_.cf29
            d_145_cf29_ = d_144___mcc_h4_
            d_146_v37_: _dafny.Array
            nw10_ = _dafny.Array(None, 27)
            d_146_v37_ = nw10_
            (d_102_globalState_).f11 = ((d_92_v4_ if d_105_v13_.f28 else d_92_v4_)) + (len(_dafny.Map({d_105_v13_.f28: d_146_v37_})))
            d_147_v38_: _dafny.Array
            nw11_ = _dafny.Array(None, 1)
            nw11_[int(0)] = d_93_v5_
            d_147_v38_ = nw11_
            index2_ = default__.safeIndex(216, (d_147_v38_).length(0))
            (d_147_v38_)[index2_] = d_93_v5_
            index3_ = default__.safeIndex(216, (d_147_v38_).length(0))
            (d_147_v38_)[index3_] = d_93_v5_
            (d_105_v13_).f28 = d_105_v13_.f28
            d_148_v39_: bool
            d_149_v40_: _dafny.Array
            d_150_v41_: _dafny.Map
            d_151_v42_: int
            out0_: bool
            out1_: _dafny.Array
            out2_: _dafny.Map
            out3_: int
            out0_, out1_, out2_, out3_ = (d_105_v13_).m0(d_105_v13_.f28, d_102_globalState_)
            d_148_v39_ = out0_
            d_149_v40_ = out1_
            d_150_v41_ = out2_
            d_151_v42_ = out3_
        elif True:
            d_152___mcc_h5_ = source3_.cf34
            d_153_cf34_ = d_152___mcc_h5_
            d_154_v43_: _dafny.Map
            d_154_v43_ = _dafny.Map({d_98_v10_: d_92_v4_})
            d_155_v44_: D14
            d_155_v44_ = D14_DC48(d_92_v4_, d_154_v43_, d_105_v13_.f28)
            source5_ = d_155_v44_
            if source5_.is_DC48:
                d_156___mcc_h8_ = source5_.cf78
                d_157___mcc_h9_ = source5_.cf79
                d_158___mcc_h10_ = source5_.cf80
                d_159_cf80_ = d_158___mcc_h10_
                d_160_cf79_ = d_157___mcc_h9_
                d_161_cf78_ = d_156___mcc_h8_
                (d_105_v13_).f28 = not((d_105_v13_).fm13((d_110_v18_) + (d_110_v18_), d_105_v13_.f28, 201, d_102_globalState_))
                rhs0_ = d_105_v13_.f28
                rhs1_ = d_105_v13_.f28
                rhs2_ = (d_161_cf78_) == (((0) - (len(d_93_v5_)) if d_159_cf80_ else d_161_cf78_))
                lhs0_ = d_105_v13_
                lhs1_ = d_105_v13_
                lhs2_ = d_102_globalState_
                lhs0_.f28 = rhs0_
                lhs1_.f28 = rhs1_
                lhs2_.f10 = rhs2_
                d_93_v5_ = (d_93_v5_) + ((d_93_v5_).set(default__.safeIndex(len(d_115_v23_), len(d_93_v5_)), d_107_v15_))
                d_162_v45_: bool
                d_163_v46_: _dafny.Array
                d_164_v47_: _dafny.Map
                d_165_v48_: int
                out4_: bool
                out5_: _dafny.Array
                out6_: _dafny.Map
                out7_: int
                out4_, out5_, out6_, out7_ = (d_105_v13_).m0(d_105_v13_.f28, d_102_globalState_)
                d_162_v45_ = out4_
                d_163_v46_ = out5_
                d_164_v47_ = out6_
                d_165_v48_ = out7_
            elif True:
                d_166___mcc_h11_ = source5_.cf77
                d_167_cf77_ = d_166___mcc_h11_
                d_168_v49_: _dafny.Map
                d_168_v49_ = _dafny.Map({(D3_DC11((d_91_v3_)[default__.safeIndex(d_92_v4_, len(d_91_v3_))], d_105_v13_.f28)).cf19: (0) - (d_92_v4_)})
                d_169_v50_: _dafny.Seq
                d_169_v50_ = _dafny.SeqWithoutIsStrInference([d_115_v23_])
                d_170_v51_: _dafny.Seq
                d_170_v51_ = _dafny.SeqWithoutIsStrInference([d_169_v50_])
                d_168_v49_ = (d_168_v49_).set(True, (d_105_v13_).fm0(d_95_v7_, d_92_v4_, (d_170_v51_)[default__.safeIndex(616, len(d_170_v51_))], d_105_v13_.f28, d_102_globalState_))
                d_171_v52_: C5
                nw12_ = C5()
                nw12_.ctor__(_dafny.Map({d_92_v4_: d_105_v13_.f28}), d_92_v4_)
                d_171_v52_ = nw12_
                d_172_v53_: C2
                nw13_ = C2()
                nw13_.ctor__((D12_DC42(d_92_v4_)).cf69, d_104_v12_)
                d_172_v53_ = nw13_
                d_173_v54_: _dafny.Set
                d_173_v54_ = _dafny.Set({d_93_v5_, d_93_v5_})
                d_174_v56_: T0
                nw14_ = C7()
                nw14_.ctor__(d_115_v23_, _dafny.SeqWithoutIsStrInference([d_98_v10_, d_98_v10_, d_98_v10_, d_98_v10_, (d_104_v12_)[default__.safeIndex((d_167_cf77_).cardinality, len(d_104_v12_))]]))
                d_174_v56_ = nw14_
                d_175_v57_: _dafny.Map
                d_175_v57_ = _dafny.Map({d_174_v56_: d_173_v54_})
                d_176_v58_: D13
                d_176_v58_ = D13_DC46((d_171_v52_).fm5(not(d_105_v13_.f28), d_102_globalState_), d_174_v56_, d_92_v4_, _dafny.Set({d_100_v11_, d_100_v11_}))
                d_177_v59_: _dafny.Map
                d_177_v59_ = _dafny.Map({d_93_v5_: d_105_v13_.f28})
                d_178_v62_: _dafny.Seq
                d_178_v62_ = _dafny.SeqWithoutIsStrInference([d_173_v54_, d_173_v54_])
                d_179_v63_: _dafny.Array
                nw15_ = _dafny.Array(None, 15)
                nw15_[int(0)] = d_173_v54_
                nw15_[int(1)] = d_173_v54_
                def iife24_():
                    coll24_ = _dafny.Set()
                    compr_24_: _dafny.Seq
                    for compr_24_ in (d_173_v54_).Elements:
                        d_180_v55_: _dafny.Seq = compr_24_
                        if (d_180_v55_) in (d_173_v54_):
                            coll24_ = coll24_.union(_dafny.Set([d_180_v55_]))
                    return _dafny.Set(coll24_)
                nw15_[int(2)] = (iife24_()
                ) - (((d_175_v57_)[(d_176_v58_).cf74] if ((d_176_v58_).cf74) in (d_175_v57_) else d_173_v54_))
                nw15_[int(3)] = d_173_v54_
                nw15_[int(4)] = d_173_v54_
                def iife25_():
                    coll25_ = _dafny.Set()
                    compr_25_: _dafny.Seq
                    for compr_25_ in (d_177_v59_).keys.Elements:
                        d_181_v60_: _dafny.Seq = compr_25_
                        if (d_181_v60_) in (d_177_v59_):
                            coll25_ = coll25_.union(_dafny.Set([d_181_v60_]))
                    return _dafny.Set(coll25_)
                nw15_[int(5)] = iife25_()
                
                nw15_[int(6)] = d_173_v54_
                nw15_[int(7)] = (default__.fm55(d_105_v13_.f28, len(d_91_v3_), d_107_v15_, d_102_globalState_)) | (d_173_v54_)
                nw15_[int(8)] = _dafny.Set({d_93_v5_, d_93_v5_, d_93_v5_})
                nw15_[int(9)] = d_173_v54_
                nw15_[int(10)] = d_173_v54_
                nw15_[int(11)] = (d_173_v54_).intersection(d_173_v54_)
                nw15_[int(12)] = default__.fm55(d_105_v13_.f28, len(d_93_v5_), d_107_v15_, d_102_globalState_)
                def iife26_():
                    coll26_ = _dafny.Set()
                    compr_26_: _dafny.Seq
                    for compr_26_ in (d_173_v54_).Elements:
                        d_182_v61_: _dafny.Seq = compr_26_
                        if (d_182_v61_) in (d_173_v54_):
                            coll26_ = coll26_.union(_dafny.Set([d_182_v61_]))
                    return _dafny.Set(coll26_)
                nw15_[int(13)] = (iife26_()
                ) | ((d_178_v62_)[default__.safeIndex((d_172_v53_).f29, len(d_178_v62_))])
                nw15_[int(14)] = d_173_v54_
                d_179_v63_ = nw15_
                index4_ = default__.safeIndex(259, (d_179_v63_).length(0))
                (d_179_v63_)[index4_] = (d_173_v54_) | (d_173_v54_)
                index5_ = default__.safeIndex(259, (d_179_v63_).length(0))
                (d_179_v63_)[index5_] = (d_173_v54_) - (d_173_v54_)
            d_183_v64_: _dafny.Array
            nw16_ = _dafny.Array(None, 12)
            nw16_[int(0)] = default__.fm38(d_90_v2_, d_92_v4_, d_102_globalState_)
            nw16_[int(1)] = d_107_v15_
            nw16_[int(2)] = d_107_v15_
            nw16_[int(3)] = d_107_v15_
            nw16_[int(4)] = d_107_v15_
            nw16_[int(5)] = d_107_v15_
            nw16_[int(6)] = d_107_v15_
            nw16_[int(7)] = _dafny.CodePoint('n')
            nw16_[int(8)] = d_107_v15_
            nw16_[int(9)] = d_107_v15_
            nw16_[int(10)] = d_107_v15_
            nw16_[int(11)] = d_107_v15_
            d_183_v64_ = nw16_
            d_184_v65_: _dafny.Seq
            d_184_v65_ = _dafny.SeqWithoutIsStrInference([d_183_v64_])
            index6_ = default__.safeIndex(323, (d_98_v10_).length(0))
            (d_98_v10_)[index6_] = (d_184_v65_) != (d_184_v65_)
            index7_ = default__.safeIndex(323, (d_98_v10_).length(0))
            rhs3_ = (len(d_93_v5_)) == ((d_92_v4_) + (len(d_93_v5_)))
            rhs4_ = d_93_v5_
            rhs5_ = (d_109_v17_) == ((d_109_v17_) | (d_109_v17_))
            rhs6_ = d_92_v4_
            rhs7_ = d_107_v15_
            lhs3_ = d_102_globalState_
            lhs4_ = d_98_v10_
            lhs5_ = default__.safeIndex(323, (d_98_v10_).length(0))
            lhs6_ = d_102_globalState_
            lhs3_.f0 = rhs3_
            d_93_v5_ = rhs4_
            lhs4_[lhs5_] = rhs5_
            lhs6_.f16 = rhs6_
            d_107_v15_ = rhs7_
            d_185_v66_: C3
            nw17_ = C3()
            nw17_.ctor__(d_104_v12_)
            d_185_v66_ = nw17_
            d_185_v66_ = (d_185_v66_ if (d_98_v10_)[default__.safeIndex(323, (d_98_v10_).length(0))] else d_185_v66_)
            d_186_v67_: _dafny.Array
            def lambda5_(d_187_i4_):
                return D15_DC52()

            init2_ = lambda5_
            nw18_ = _dafny.Array(None, 3)
            for i0_2_ in range(nw18_.length(0)):
                nw18_[i0_2_] = init2_(i0_2_)
            d_186_v67_ = nw18_
            d_188_v68_: D15
            d_188_v68_ = D15_DC52()
            index8_ = default__.safeIndex(193, (d_186_v67_).length(0))
            (d_186_v67_)[index8_] = d_188_v68_
            d_189_v69_: _dafny.Array
            def lambda6_(d_190_v13_):
                def lambda7_(d_191_i5_):
                    return d_190_v13_.f28

                return lambda7_

            init3_ = lambda6_(d_105_v13_)
            nw19_ = _dafny.Array(None, 27)
            for i0_3_ in range(nw19_.length(0)):
                nw19_[i0_3_] = init3_(i0_3_)
            d_189_v69_ = nw19_
            d_192_v70_: C2
            nw20_ = C2()
            nw20_.ctor__((0) - (d_92_v4_), _dafny.SeqWithoutIsStrInference([d_189_v69_, d_189_v69_]))
            d_192_v70_ = nw20_
            index9_ = default__.safeIndex(193, (d_186_v67_).length(0))
            rhs8_ = ((d_88_v0_).cardinality) >= ((d_192_v70_).f29)
            rhs9_ = d_188_v68_
            rhs10_ = d_192_v70_
            lhs7_ = d_105_v13_
            lhs8_ = d_186_v67_
            lhs9_ = default__.safeIndex(193, (d_186_v67_).length(0))
            lhs7_.f28 = rhs8_
            lhs8_[lhs9_] = rhs9_
            d_192_v70_ = rhs10_
        d_193_v71_: D18
        d_193_v71_ = D18_DC57(d_105_v13_)
        d_194_v72_: _dafny.Map
        d_194_v72_ = _dafny.Map({d_90_v2_: d_105_v13_})
        d_195_v73_: _dafny.Array
        nw21_ = _dafny.Array(None, 13)
        nw21_[int(0)] = (d_193_v71_).cf88
        nw21_[int(1)] = d_105_v13_
        nw21_[int(2)] = ((d_194_v72_)[d_90_v2_] if (d_90_v2_) in (d_194_v72_) else d_105_v13_)
        nw21_[int(3)] = d_105_v13_
        nw21_[int(4)] = d_105_v13_
        nw21_[int(5)] = d_105_v13_
        nw21_[int(6)] = d_105_v13_
        nw21_[int(7)] = d_105_v13_
        nw21_[int(8)] = d_105_v13_
        nw21_[int(9)] = (d_193_v71_).cf88
        nw21_[int(10)] = d_105_v13_
        nw21_[int(11)] = d_105_v13_
        nw21_[int(12)] = d_105_v13_
        d_195_v73_ = nw21_
        index10_ = default__.safeIndex(306, (d_195_v73_).length(0))
        (d_195_v73_)[index10_] = d_105_v13_
        index11_ = default__.safeIndex(306, (d_195_v73_).length(0))
        (d_195_v73_)[index11_] = d_105_v13_
        if (d_92_v4_) == (d_92_v4_):
            d_196_v74_: _dafny.Map
            d_196_v74_ = _dafny.Map({d_105_v13_.f28: d_92_v4_})
            d_197_v75_: _dafny.Seq
            d_197_v75_ = _dafny.SeqWithoutIsStrInference([d_93_v5_])
            d_196_v74_ = (d_196_v74_).set(d_90_v2_, (len((d_197_v75_)[default__.safeIndex(d_92_v4_, len(d_197_v75_))])) + (d_92_v4_))
            d_198_v76_: C3
            nw22_ = C3()
            nw22_.ctor__(_dafny.SeqWithoutIsStrInference([d_98_v10_, d_98_v10_, d_98_v10_, d_98_v10_, d_98_v10_]))
            d_198_v76_ = nw22_
            d_199_v77_: _dafny.Map
            d_199_v77_ = _dafny.Map({len(d_115_v23_): d_98_v10_})
            d_200_v78_: _dafny.Map
            d_200_v78_ = _dafny.Map({((d_199_v77_)[d_92_v4_] if (d_92_v4_) in (d_199_v77_) else d_98_v10_): d_100_v11_})
            d_200_v78_ = (d_200_v78_).set(d_98_v10_, d_100_v11_)
            (d_102_globalState_).f16 = default__.safeDivisionInt((len(d_93_v5_)) * (len(_dafny.Map({d_92_v4_: d_90_v2_}))), d_92_v4_)
            d_201_v79_: _dafny.Seq
            d_201_v79_ = _dafny.SeqWithoutIsStrInference([d_109_v17_, d_109_v17_, _dafny.Set({d_105_v13_.f28, d_105_v13_.f28})])
            d_202_v80_: D12
            d_202_v80_ = D12_DC42(len(d_115_v23_))
            d_203_v81_: _dafny.Map
            d_203_v81_ = _dafny.Map({d_202_v80_: d_109_v17_})
            d_204_v82_: _dafny.Array
            nw23_ = _dafny.Array(None, 1)
            nw23_[int(0)] = (d_201_v79_) + (_dafny.SeqWithoutIsStrInference([((d_203_v81_)[D12_DC42(d_92_v4_)] if (D12_DC42(d_92_v4_)) in (d_203_v81_) else d_109_v17_), d_109_v17_]))
            d_204_v82_ = nw23_
            index12_ = default__.safeIndex(435, (d_204_v82_).length(0))
            (d_204_v82_)[index12_] = d_201_v79_
            index13_ = default__.safeIndex(435, (d_204_v82_).length(0))
            (d_204_v82_)[index13_] = d_201_v79_
        elif True:
            d_205_v83_: C0
            nw24_ = C0()
            nw24_.ctor__()
            d_205_v83_ = nw24_
            d_206_v84_: _dafny.MultiSet
            d_206_v84_ = _dafny.MultiSet([d_205_v83_])
            d_207_v85_: _dafny.Map
            d_207_v85_ = _dafny.Map({113: D12_DC41(d_206_v84_)})
            d_207_v85_ = d_207_v85_
            d_208_v86_: C3
            nw25_ = C3()
            nw25_.ctor__(d_104_v12_)
            d_208_v86_ = nw25_
            rhs11_ = (default__.safeDivisionInt(d_92_v4_, len((d_93_v5_).set(default__.safeIndex(-282, len(d_93_v5_)), d_107_v15_)))) * (208)
            rhs12_ = d_208_v86_
            d_92_v4_ = rhs11_
            d_208_v86_ = rhs12_
            d_93_v5_ = d_93_v5_
            d_209_v87_: _dafny.Map
            d_209_v87_ = _dafny.Map({d_91_v3_: d_92_v4_})
            d_210_v88_: _dafny.Seq
            d_210_v88_ = _dafny.SeqWithoutIsStrInference([d_209_v87_, d_209_v87_, d_209_v87_])
            (d_102_globalState_).f8 = (287) * (len((d_210_v88_)[default__.safeIndex(d_92_v4_, len(d_210_v88_))]))
            index14_ = default__.safeIndex(698, (d_98_v10_).length(0))
            (d_98_v10_)[index14_] = default__.fm15(d_90_v2_, d_102_globalState_)
            index15_ = default__.safeIndex(698, (d_98_v10_).length(0))
            (d_98_v10_)[index15_] = d_105_v13_.f28
        d_211_v89_: bool
        d_212_v90_: _dafny.Array
        d_213_v91_: _dafny.Map
        d_214_v92_: int
        out8_: bool
        out9_: _dafny.Array
        out10_: _dafny.Map
        out11_: int
        out8_, out9_, out10_, out11_ = (d_105_v13_).m0(d_90_v2_, d_102_globalState_)
        d_211_v89_ = out8_
        d_212_v90_ = out9_
        d_213_v91_ = out10_
        d_214_v92_ = out11_
        hi0_ = 465
        for d_215_i6_ in range((0) - (d_92_v4_), hi0_):
            d_216_v93_: D12
            d_216_v93_ = D12_DC42(d_92_v4_)
            def iife27_(_pat_let0_0):
                def iife28_(d_217_dt__update__tmp_h0_):
                    def iife29_(_pat_let1_0):
                        def iife30_(d_218_dt__update_hcf69_h0_):
                            return D12_DC42(d_218_dt__update_hcf69_h0_)
                        return iife30_(_pat_let1_0)
                    return iife29_(d_215_i6_)
                return iife28_(_pat_let0_0)
            d_216_v93_ = iife27_(d_216_v93_)
            if True:
                d_219_v94_: C1
                nw26_ = C1()
                nw26_.ctor__(d_214_v92_, d_104_v12_)
                d_219_v94_ = nw26_
                d_220_v95_: _dafny.Seq
                d_220_v95_ = _dafny.SeqWithoutIsStrInference([d_219_v94_, d_219_v94_])
                d_211_v89_ = ((0) - (d_92_v4_)) <= (default__.safeDivisionInt(len(d_220_v95_), d_92_v4_))
                d_221_v96_: C4
                nw27_ = C4()
                nw27_.ctor__(d_211_v89_, (d_104_v12_) + (d_104_v12_))
                d_221_v96_ = nw27_
                d_88_v0_ = d_88_v0_
                d_222_v97_: _dafny.Array
                nw28_ = _dafny.Array(None, 3)
                d_222_v97_ = nw28_
                d_223_v98_: _dafny.Map
                d_223_v98_ = _dafny.Map({d_215_i6_: True})
                d_224_v99_: T0
                nw29_ = C1()
                nw29_.ctor__(len(d_223_v98_), d_104_v12_)
                d_224_v99_ = nw29_
                index16_ = default__.safeIndex(277, (d_222_v97_).length(0))
                (d_222_v97_)[index16_] = d_224_v99_
                index17_ = default__.safeIndex(277, (d_222_v97_).length(0))
                (d_222_v97_)[index17_] = d_224_v99_
                d_225_v100_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_225_v100_ = nw30_
                d_226_v101_: _dafny.MultiSet
                d_226_v101_ = _dafny.MultiSet([d_225_v100_])
                index18_ = default__.safeIndex(463, (d_98_v10_).length(0))
                (d_98_v10_)[index18_] = (d_226_v101_).ispropersubset(_dafny.MultiSet([d_225_v100_, d_225_v100_]))
                index19_ = default__.safeIndex(463, (d_98_v10_).length(0))
                (d_98_v10_)[index19_] = not(d_105_v13_.f28)
            elif True:
                (d_102_globalState_).f0 = d_105_v13_.f28
                d_110_v18_ = d_110_v18_
                d_227_v102_: _dafny.Map
                d_227_v102_ = _dafny.Map({d_214_v92_: d_215_i6_})
                d_228_v103_: D5
                d_228_v103_ = D5_DC19(d_227_v102_, d_90_v2_)
                d_229_v104_: bool
                d_230_v105_: _dafny.Array
                d_231_v106_: _dafny.Map
                d_232_v107_: int
                out12_: bool
                out13_: _dafny.Array
                out14_: _dafny.Map
                out15_: int
                out12_, out13_, out14_, out15_ = (d_105_v13_).m0((not(True) if d_90_v2_ else (d_228_v103_).cf33), d_102_globalState_)
                d_229_v104_ = out12_
                d_230_v105_ = out13_
                d_231_v106_ = out14_
                d_232_v107_ = out15_
                d_233_v108_: _dafny.Map
                d_233_v108_ = _dafny.Map({d_232_v107_: d_90_v2_})
                d_234_v109_: C5
                nw31_ = C5()
                nw31_.ctor__(d_233_v108_, d_92_v4_)
                d_234_v109_ = nw31_
                d_235_v110_: bool
                d_236_v111_: _dafny.Array
                d_237_v112_: _dafny.Map
                d_238_v113_: int
                out16_: bool
                out17_: _dafny.Array
                out18_: _dafny.Map
                out19_: int
                out16_, out17_, out18_, out19_ = (d_105_v13_).m0((default__.fm56(d_232_v107_, d_102_globalState_)).cf1, d_102_globalState_)
                d_235_v110_ = out16_
                d_236_v111_ = out17_
                d_237_v112_ = out18_
                d_238_v113_ = out19_
            index20_ = default__.safeIndex(472, (d_100_v11_).length(0))
            (d_100_v11_)[index20_] = d_215_i6_
            index21_ = default__.safeIndex(472, (d_100_v11_).length(0))
            (d_100_v11_)[index21_] = default__.safeDivisionInt(628, (d_92_v4_) + (d_215_i6_))
            d_239_v114_: C7
            nw32_ = C7()
            nw32_.ctor__(d_115_v23_, _dafny.SeqWithoutIsStrInference([d_98_v10_]))
            d_239_v114_ = nw32_
        if not((d_91_v3_)[default__.safeIndex(254, len(d_91_v3_))]):
            d_93_v5_ = (d_93_v5_).set(default__.safeIndex((d_214_v92_) * (-876), len(d_93_v5_)), d_107_v15_)
            d_240_v115_: bool
            d_241_v116_: _dafny.Array
            d_242_v117_: _dafny.Map
            d_243_v118_: int
            out20_: bool
            out21_: _dafny.Array
            out22_: _dafny.Map
            out23_: int
            out20_, out21_, out22_, out23_ = (d_105_v13_).m0(d_211_v89_, d_102_globalState_)
            d_240_v115_ = out20_
            d_241_v116_ = out21_
            d_242_v117_ = out22_
            d_243_v118_ = out23_
            d_244_v119_: bool
            d_245_v120_: _dafny.Array
            d_246_v121_: _dafny.Map
            d_247_v122_: int
            out24_: bool
            out25_: _dafny.Array
            out26_: _dafny.Map
            out27_: int
            out24_, out25_, out26_, out27_ = (d_105_v13_).m0((d_95_v7_).ispropersubset(d_95_v7_), d_102_globalState_)
            d_244_v119_ = out24_
            d_245_v120_ = out25_
            d_246_v121_ = out26_
            d_247_v122_ = out27_
            d_248_v123_: _dafny.Set
            d_248_v123_ = _dafny.Set({d_91_v3_, d_91_v3_, d_91_v3_})
            d_240_v115_ = ((d_248_v123_) | (_dafny.Set({d_91_v3_, d_91_v3_}))) != (d_248_v123_)
            d_249_v124_: D17
            d_249_v124_ = D17_DC56(d_90_v2_)
            (d_102_globalState_).f0 = (((d_249_v124_).cf87) or (d_90_v2_)) or (d_105_v13_.f28)
        elif True:
            d_250_v125_: _dafny.Set
            d_250_v125_ = _dafny.Set({d_111_v19_, (d_111_v19_) | (d_111_v19_)})
            d_250_v125_ = d_250_v125_
            d_251_v126_: _dafny.Map
            d_251_v126_ = _dafny.Map({(d_105_v13_.f28 if not(d_105_v13_.f28) else d_211_v89_): default__.fm18(d_90_v2_, d_102_globalState_)})
            d_251_v126_ = d_251_v126_
            d_252_v127_: _dafny.Map
            d_252_v127_ = _dafny.Map({475: d_214_v92_})
            d_253_v128_: _dafny.Map
            d_253_v128_ = _dafny.Map({d_214_v92_: d_252_v127_})
            index22_ = default__.safeIndex(318, (d_100_v11_).length(0))
            (d_100_v11_)[index22_] = len((d_252_v127_) | (((d_253_v128_)[d_214_v92_] if (d_214_v92_) in (d_253_v128_) else d_252_v127_)))
            index23_ = default__.safeIndex(318, (d_100_v11_).length(0))
            (d_100_v11_)[index23_] = 170
            d_254_v129_: C4
            nw33_ = C4()
            nw33_.ctor__(d_211_v89_, (d_104_v12_) + (d_104_v12_))
            d_254_v129_ = nw33_
            (d_102_globalState_).f8 = default__.fm7(len(d_115_v23_), d_102_globalState_)
        d_255_v130_: bool
        d_256_v131_: _dafny.Array
        d_257_v132_: _dafny.Map
        d_258_v133_: int
        out28_: bool
        out29_: _dafny.Array
        out30_: _dafny.Map
        out31_: int
        out28_, out29_, out30_, out31_ = (d_105_v13_).m0(d_211_v89_, d_102_globalState_)
        d_255_v130_ = out28_
        d_256_v131_ = out29_
        d_257_v132_ = out30_
        d_258_v133_ = out31_
        d_93_v5_ = d_93_v5_
        d_259_v134_: C3
        nw34_ = C3()
        nw34_.ctor__(d_104_v12_)
        d_259_v134_ = nw34_
        d_90_v2_ = (d_92_v4_) >= ((d_214_v92_) * ((_dafny.MultiSet([d_259_v134_])).cardinality))
        d_260_i7_: int
        d_260_i7_ = 0
        with _dafny.label("0"):
            while not(d_105_v13_.f28):
                with _dafny.c_label("0"):
                    if (d_260_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_260_i7_ = (d_260_i7_) + (1)
                    index24_ = default__.safeIndex(845, (d_100_v11_).length(0))
                    (d_100_v11_)[index24_] = d_92_v4_
                    index25_ = default__.safeIndex(845, (d_100_v11_).length(0))
                    (d_100_v11_)[index25_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxxok")))) * (default__.safeModuloInt(d_258_v133_, d_258_v133_))
                    (d_102_globalState_).f11 = (d_100_v11_)[default__.safeIndex(845, (d_100_v11_).length(0))]
                    d_261_v135_: _dafny.Seq
                    d_261_v135_ = _dafny.SeqWithoutIsStrInference([d_115_v23_, d_115_v23_])
                    (d_102_globalState_).f0 = (318) == ((d_105_v13_).fm0(d_94_v6_, (d_100_v11_)[default__.safeIndex(845, (d_100_v11_).length(0))], d_261_v135_, d_105_v13_.f28, d_102_globalState_))
                    index26_ = default__.safeIndex(306, (d_195_v73_).length(0))
                    (d_195_v73_)[index26_] = (d_195_v73_)[default__.safeIndex(306, (d_195_v73_).length(0))]
                    pass
            pass
        d_262_v136_: _dafny.Array
        d_263_v137_: int
        out32_: _dafny.Array
        out33_: int
        out32_, out33_ = (d_259_v134_).m4(d_102_globalState_)
        d_262_v136_ = out32_
        d_263_v137_ = out33_
        d_264_v138_: _dafny.Map
        d_264_v138_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_107_v15_ for d_265_i8_ in range(default__.abs(188))])): d_211_v89_})
        d_266_v139_: C5
        nw35_ = C5()
        nw35_.ctor__(d_264_v138_, (d_92_v4_) * ((0) - (d_258_v133_)))
        d_266_v139_ = nw35_
        d_267_v140_: _dafny.MultiSet
        d_267_v140_ = _dafny.MultiSet([d_107_v15_, _dafny.CodePoint('d')])
        d_268_v141_: _dafny.Map
        d_268_v141_ = _dafny.Map({d_267_v140_: d_214_v92_})
        d_268_v141_ = (d_268_v141_).set((d_267_v140_).set(d_107_v15_, default__.abs((d_266_v139_).f27)), default__.safeDivisionInt(len(d_110_v18_), d_263_v137_))
        _dafny.print(_dafny.string_of((d_88_v0_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v1_) == (_dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_90_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_91_v3_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_92_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_93_v5_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_94_v6_) == (_dafny.MultiSet([1, 476]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_95_v7_) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_96_v8_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([2])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_97_v9_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([2])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v10_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v11_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_.f1) == (_dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f2) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f3) == (_dafny.SeqWithoutIsStrInference([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_.f5) == (_dafny.SeqWithoutIsStrInference([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_.f13) == (_dafny.Set({166}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f15) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_102_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f18)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f19)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f20) == (_dafny.MultiSet([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_globalState_).f21)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_104_v12_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_105_v13_.f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_106_v14_).cf14).cf6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_v14_).cf15)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v14_).cf16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_107_v15_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_108_v16_) == (_dafny.Map({_dafny.CodePoint('m'): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_109_v17_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_110_v18_) == (_dafny.SeqWithoutIsStrInference([-999, 166]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v19_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_114_v22_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_115_v23_) == (_dafny.Set({9, 166}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_v24_).cf86) == (_dafny.Set({9, 166}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v71_).cf88.f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_193_v71_).cf88).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_194_v72_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[0].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[0]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[1].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[1]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[2].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[2]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[3].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[3]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[4].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[4]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[5].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[5]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[6].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[6]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[7].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[7]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[8].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[8]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[9].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[9]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[10].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[10]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[11].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[11]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v73_)[12].f28))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_195_v73_)[12]).f22)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_211_v89_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v90_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v91_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([509, 508, 508, 508, 509]): _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_214_v92_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_255_v130_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_256_v131_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v132_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([509, 508, 508, 508, 509]): _dafny.CodePoint('l')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_258_v133_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_260_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v136_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_263_v137_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v138_) == (_dafny.Map({188: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_266_v139_).f26) == (_dafny.Map({188: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_266_v139_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v140_) == (_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('d')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v141_) == (_dafny.Map({_dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('d')]): 4, _dafny.MultiSet([_dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('m'), _dafny.CodePoint('d')]): 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({self.cf6.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D2_DC9)

class D2_DC6(D2, NamedTuple('DC6', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC11(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(D0.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D4_DC16)

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {self.cf25.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC16(D4, NamedTuple('DC16', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC16({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC16) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC18(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D5_DC19)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D5_DC20)

class D5_DC18(D5, NamedTuple('DC18', [('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC19(D5, NamedTuple('DC19', [('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC19({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC19) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC20(D5, NamedTuple('DC20', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC20({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC20) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC22(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D6_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D6_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D6_DC21)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D6_DC24)

class D6_DC22(D6, NamedTuple('DC22', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC22({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC22) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC23(D6, NamedTuple('DC23', [('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC23({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC23) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC21(D6, NamedTuple('DC21', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC21({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC21) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC24(D6, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC26(_dafny.Seq({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D7_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D7_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D7_DC27)

class D7_DC26(D7, NamedTuple('DC26', [('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC26({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC26) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC25(D7, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC27(D7, NamedTuple('DC27', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC27({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC27) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC29(False, False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D8_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D8_DC28)

class D8_DC29(D8, NamedTuple('DC29', [('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC29({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC29) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC28(D8, NamedTuple('DC28', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC28({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC28) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC31(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D9_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D9_DC30)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D9_DC32)

class D9_DC31(D9, NamedTuple('DC31', [('cf51', Any), ('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC31({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC31) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC30(D9, NamedTuple('DC30', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC30({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC30) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC32(D9, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC34(_dafny.CodePoint('D'), False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D10_DC34)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D10_DC35)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D10_DC33)

class D10_DC34(D10, NamedTuple('DC34', [('cf55', Any), ('cf56', Any), ('cf57', Any), ('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC34({_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC34) and self.cf55 == __o.cf55 and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC35(D10, NamedTuple('DC35', [('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC35({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC35) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC33(D10, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC33({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC37(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D11_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D11_DC38)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D11_DC39)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D11_DC36)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D11_DC40)

class D11_DC37(D11, NamedTuple('DC37', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC37({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC37) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC38(D11, NamedTuple('DC38', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC38({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC38) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC39(D11, NamedTuple('DC39', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC39({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC39) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC36(D11, NamedTuple('DC36', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC36({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC36) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC40(D11, NamedTuple('DC40', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC40({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC40) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC42(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D12_DC42)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D12_DC41)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D12_DC43)

class D12_DC42(D12, NamedTuple('DC42', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC42({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC42) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC41(D12, NamedTuple('DC41', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC41({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC41) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC43(D12, NamedTuple('DC43', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC43({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC43) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC45(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D13_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D13_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D13_DC44)

class D13_DC45(D13, NamedTuple('DC45', [('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC45({_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC45) and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC46(D13, NamedTuple('DC46', [('cf73', Any), ('cf74', Any), ('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC46({_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC46) and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC44(D13, NamedTuple('DC44', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC44({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC44) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC48(int(0), _dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D14_DC48)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D14_DC47)

class D14_DC48(D14, NamedTuple('DC48', [('cf78', Any), ('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC48({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC48) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC47(D14, NamedTuple('DC47', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC47({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC47) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC50()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D15_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D15_DC51)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D15_DC52)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D15_DC49)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D15_DC53)

class D15_DC50(D15, NamedTuple('DC50', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC50'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC50)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC51(D15, NamedTuple('DC51', [('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC51({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC51) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC52(D15, NamedTuple('DC52', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC52'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC52)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC49(D15, NamedTuple('DC49', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC49({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC49) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC53(D15, NamedTuple('DC53', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC53({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC53) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D16_DC54)

class D16_DC54(D16, NamedTuple('DC54', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC54({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC54) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC56(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D17_DC56)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D17_DC55)

class D17_DC56(D17, NamedTuple('DC56', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC56({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC56) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC55(D17, NamedTuple('DC55', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC55({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC55) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC58()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D18_DC58)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D18_DC57)

class D18_DC58(D18, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC57(D18, NamedTuple('DC57', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC57({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC57) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC60(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D19_DC60)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D19_DC59)

class D19_DC60(D19, NamedTuple('DC60', [('cf90', Any), ('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC60({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC60) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC59(D19, NamedTuple('DC59', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC59({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC59) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC62(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D20_DC62)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D20_DC61)

class D20_DC62(D20, NamedTuple('DC62', [('cf94', Any), ('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC62({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC62) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC61(D20, NamedTuple('DC61', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC61({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC61) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m0(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f1: _dafny.MultiSet = _dafny.MultiSet({})
        self.f5: _dafny.Seq = _dafny.Seq({})
        self.f6: int = int(0)
        self.f8: int = int(0)
        self.f10: bool = False
        self.f11: int = int(0)
        self.f13: _dafny.Set = _dafny.Set({})
        self.f16: int = int(0)
        self._f2: _dafny.Seq = _dafny.Seq({})
        self._f3: _dafny.Seq = _dafny.Seq({})
        self._f4: bool = False
        self._f7: bool = False
        self._f9: int = int(0)
        self._f12: int = int(0)
        self._f14: int = int(0)
        self._f15: _dafny.MultiSet = _dafny.MultiSet({})
        self._f17: int = int(0)
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f19: _dafny.Array = _dafny.Array(None, 0)
        self._f20: _dafny.MultiSet = _dafny.MultiSet({})
        self._f21: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self)._f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21

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
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9
    @property
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm16(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(False), False])) + ((_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([False])))

    def fm17(self, p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_269_i0_ in range(default__.abs(712))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdpncg")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujpfdfbpa")))


class C1(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self._f30: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f30, f22):
        (self)._f30 = f30
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return (self).f30

    def fm34(self, p0, globalState):
        return (259) * ((self).f30)

    def fm35(self, p0, p1, p2, p3, globalState):
        def iife31_():
            coll27_ = _dafny.Map()
            compr_27_: str
            for compr_27_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_270_i0_ in range(default__.abs(141))])).Elements:
                d_271_v0_: str = compr_27_
                if (d_271_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_270_i0_ in range(default__.abs(141))])):
                    coll27_[d_271_v0_] = (self).f30
            return _dafny.Map(coll27_)
        def iife32_():
            coll28_ = _dafny.Map()
            compr_28_: str
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('n')])).Elements:
                d_272_v1_: str = compr_28_
                if (d_272_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('n')])):
                    coll28_[d_272_v1_] = len(_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True, True, True]))).cardinality: False}))
            return _dafny.Map(coll28_)
        return len((_dafny.Map({_dafny.CodePoint('y'): (self).f30})) | ((iife31_()
        ) | (iife32_()
        )))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        (globalState).f11 = 664
        d_273_v0_: _dafny.Array
        nw36_ = _dafny.Array(False, 20)
        d_273_v0_ = nw36_
        index27_ = default__.safeIndex(722, (d_273_v0_).length(0))
        (d_273_v0_)[index27_] = True
        d_274_v1_: D3
        d_274_v1_ = D3_DC11(p0, p0)
        d_275_v2_: _dafny.Map
        d_275_v2_ = _dafny.Map({(self).f30: len(_dafny.Map({True: p0}))})
        d_276_v3_: _dafny.Map
        d_276_v3_ = _dafny.Map({p0: d_275_v2_})
        d_277_v4_: str
        d_277_v4_ = _dafny.CodePoint('o')
        d_278_v5_: D0
        d_278_v5_ = D0_DC1((self).f30, p0, d_276_v3_, (self).f30, d_277_v4_)
        d_279_v6_: _dafny.Seq
        d_279_v6_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        pat_let_tv6_ = p0
        index28_ = default__.safeIndex(722, (d_273_v0_).length(0))
        def lambda8_(source6_):
            if source6_.is_DC11:
                d_280___mcc_h0_ = source6_.cf19
                d_281___mcc_h1_ = source6_.cf20
                d_282_cf20_ = d_281___mcc_h1_
                d_283_cf19_ = d_280___mcc_h0_
                return d_282_cf20_
            elif True:
                d_284___mcc_h2_ = source6_.cf18
                d_285_cf18_ = d_284___mcc_h2_
                return pat_let_tv6_

        rhs13_ = not(p0)
        rhs14_ = lambda8_(d_274_v1_)
        rhs15_ = (((self).f30 if p0 else (d_278_v5_).cf3)) != (len(d_279_v6_))
        lhs10_ = globalState
        lhs11_ = globalState
        lhs12_ = d_273_v0_
        lhs13_ = default__.safeIndex(722, (d_273_v0_).length(0))
        lhs10_.f0 = rhs13_
        lhs11_.f0 = rhs14_
        lhs12_[lhs13_] = rhs15_
        index29_ = default__.safeIndex(722, (d_273_v0_).length(0))
        (d_273_v0_)[index29_] = (False if p0 else ((d_273_v0_)[default__.safeIndex(722, (d_273_v0_).length(0))]) and (p0))
        d_286_v7_: D8
        d_286_v7_ = D8_DC29((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jn")))) <= (len(d_275_v2_)), ((self).f30) <= ((self).f30), d_277_v4_)
        d_286_v7_ = (D8_DC29(p0, (d_273_v0_)[default__.safeIndex(722, (d_273_v0_).length(0))], d_277_v4_) if (d_273_v0_)[default__.safeIndex(722, (d_273_v0_).length(0))] else d_286_v7_)
        (globalState).f16 = (default__.safeModuloInt((self).f30, (self).f30)) * ((self).f30)
        d_287_v8_: _dafny.Set
        d_287_v8_ = _dafny.Set({(self).f30})
        d_288_v9_: _dafny.MultiSet
        d_288_v9_ = _dafny.MultiSet([d_273_v0_])
        (globalState).f11 = ((len(d_287_v8_)) + ((self).f30)) * (((d_288_v9_) - (d_288_v9_)).cardinality)
        r0 = (d_273_v0_)[default__.safeIndex(722, (d_273_v0_).length(0))]
        r1 = d_273_v0_
        d_289_v10_: _dafny.Seq
        d_289_v10_ = _dafny.SeqWithoutIsStrInference([(self).f30])
        d_290_v11_: _dafny.Seq
        d_290_v11_ = _dafny.SeqWithoutIsStrInference([d_279_v6_])
        d_291_v12_: _dafny.Seq
        d_291_v12_ = _dafny.SeqWithoutIsStrInference([len(d_289_v10_), (self).f30, len(d_290_v11_)])
        d_292_v13_: _dafny.Seq
        d_292_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aipcn"))
        d_293_v14_: _dafny.Map
        d_293_v14_ = _dafny.Map({(d_291_v12_).set(default__.safeIndex(len(d_292_v13_), len(d_291_v12_)), (self).f30): (d_277_v4_ if (d_273_v0_)[default__.safeIndex(722, (d_273_v0_).length(0))] else _dafny.CodePoint('y'))})
        r2 = d_293_v14_
        d_294_v15_: D6
        d_294_v15_ = D6_DC22((0) - ((self).f30))
        r3 = (self).fm35((self).f30, (0) - ((self).f30), p0, (0) - ((d_294_v15_).cf36), globalState)
        return r0, r1, r2, r3

    def m6(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        r2: int = int(0)
        r3: int = int(0)
        (globalState).f0 = (p1 if True else p1)
        d_295_v0_: _dafny.Map
        d_295_v0_ = _dafny.Map({(self).f30: (self).f30})
        d_296_v1_: _dafny.Array
        nw37_ = _dafny.Array(False, 20)
        d_296_v1_ = nw37_
        d_297_v2_: _dafny.MultiSet
        d_297_v2_ = _dafny.MultiSet([d_296_v1_, d_296_v1_])
        d_298_v3_: _dafny.Seq
        d_298_v3_ = _dafny.SeqWithoutIsStrInference([p1, False, p1, p1, p1])
        d_295_v0_ = (d_295_v0_).set(default__.safeDivisionInt((d_297_v2_).cardinality, (_dafny.MultiSet(d_298_v3_)).cardinality), (p2) - ((self).fm35(220, p2, False, (0) - (p0), globalState)))
        d_299_v4_: _dafny.MultiSet
        d_299_v4_ = _dafny.MultiSet([p2])
        (globalState).f0 = not(((d_299_v4_).issubset(d_299_v4_) if False else not((p1 if p1 else not(p1)))))
        index30_ = default__.safeIndex(191, (d_296_v1_).length(0))
        (d_296_v1_)[index30_] = p1
        index31_ = default__.safeIndex(191, (d_296_v1_).length(0))
        (d_296_v1_)[index31_] = p1
        d_300_v5_: _dafny.Map
        d_300_v5_ = _dafny.Map({d_296_v1_: False})
        d_301_v6_: _dafny.Set
        d_301_v6_ = _dafny.Set({d_298_v3_, _dafny.SeqWithoutIsStrInference([p1, not(p1), not(((d_300_v5_)[d_296_v1_] if (d_296_v1_) in (d_300_v5_) else (d_296_v1_)[default__.safeIndex(191, (d_296_v1_).length(0))]))]), d_298_v3_, d_298_v3_})
        d_301_v6_ = ((d_301_v6_) | (d_301_v6_)) - ((d_301_v6_) | (_dafny.Set({d_298_v3_, (d_298_v3_).set(default__.safeIndex((self).fm0(d_299_v4_, p3, _dafny.SeqWithoutIsStrInference([_dafny.Set({p0}) for d_302_i0_ in range(default__.abs(773))]), p1, globalState), len(d_298_v3_)), p1), d_298_v3_})))
        d_303_v7_: _dafny.Map
        d_303_v7_ = _dafny.Map({(d_296_v1_)[default__.safeIndex(191, (d_296_v1_).length(0))]: (self).f30})
        d_303_v7_ = (D9_DC30(d_303_v7_)).cf50
        r0 = -425
        d_304_v8_: _dafny.Set
        d_304_v8_ = _dafny.Set({(self).f30, p0, p0})
        r1 = d_304_v8_
        r2 = p0
        r3 = p2
        return r0, r1, r2, r3

    def m7(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: _dafny.Set = _dafny.Set({})
        (globalState).f8 = (self).f30
        r2 = p0
        d_305_v0_: _dafny.MultiSet
        d_305_v0_ = _dafny.MultiSet([214, (self).f30])
        d_306_v1_: _dafny.Map
        d_306_v1_ = _dafny.Map({(self).f30: d_305_v0_})
        d_307_v2_: D4
        d_307_v2_ = D4_DC12(d_306_v1_)
        d_308_v3_: _dafny.Map
        d_308_v3_ = _dafny.Map({p0: d_307_v2_})
        d_309_v4_: _dafny.Seq
        d_309_v4_ = _dafny.SeqWithoutIsStrInference([112, len(d_308_v3_), (self).f30])
        d_310_v5_: _dafny.Seq
        d_310_v5_ = _dafny.SeqWithoutIsStrInference([d_309_v4_, (d_309_v4_) + (_dafny.SeqWithoutIsStrInference([(self).f30 for d_311_i0_ in range(default__.abs(904))])), d_309_v4_])
        d_309_v4_ = (d_310_v5_)[default__.safeIndex((self).f30, len(d_310_v5_))]
        d_312_v6_: str
        d_312_v6_ = _dafny.CodePoint('i')
        d_313_v7_: _dafny.Map
        d_313_v7_ = _dafny.Map({p0: d_312_v6_})
        d_314_i1_: int
        d_314_i1_ = 0
        with _dafny.label("1"):
            while (len(_dafny.Map({((d_313_v7_)[p0] if (p0) in (d_313_v7_) else d_312_v6_): (self).f30}))) > (len(_dafny.Map({p0: p0}))):
                with _dafny.c_label("1"):
                    if (d_314_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_314_i1_ = (d_314_i1_) + (1)
                    if p0:
                        d_315_v8_: _dafny.Array
                        nw38_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
                        d_315_v8_ = nw38_
                        d_316_v9_: _dafny.Map
                        d_316_v9_ = _dafny.Map({d_312_v6_: d_315_v8_})
                        d_316_v9_ = (d_316_v9_).set(d_312_v6_, d_315_v8_)
                        d_312_v6_ = d_312_v6_
                        r0 = ((self).f30) >= ((0) - (len((d_309_v4_) + (d_309_v4_))))
                        d_317_v10_: _dafny.Seq
                        d_317_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frwriukby"))
                        d_317_v10_ = (d_317_v10_) + (default__.fm36(p0, _dafny.Set({(self).f30, (self).f30}), (self).f30, (self).f30, globalState))
                        d_318_v11_: _dafny.Array
                        nw39_ = _dafny.Array(False, 2)
                        d_318_v11_ = nw39_
                        rhs16_ = d_317_v10_
                        rhs17_ = d_312_v6_
                        rhs18_ = d_318_v11_
                        rhs19_ = d_317_v10_
                        d_317_v10_ = rhs16_
                        d_312_v6_ = rhs17_
                        d_318_v11_ = rhs18_
                        d_317_v10_ = rhs19_
                    elif True:
                        (globalState).f6 = (self).f30
                        d_319_v12_: D2
                        d_319_v12_ = D2_DC7((d_309_v4_).set(default__.safeIndex((self).f30, len(d_309_v4_)), (self).f30), p0, p0, 27)
                        (globalState).f0 = (d_319_v12_).cf11
                        d_312_v6_ = d_312_v6_
                        (globalState).f8 = ((self).f30) * ((self).f30)
                        d_320_v13_: _dafny.Seq
                        d_320_v13_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_321_v14_: _dafny.Map
                        d_321_v14_ = _dafny.Map({(self).f30: p0})
                        d_322_v15_: _dafny.Seq
                        d_322_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xw"))
                        (globalState).f10 = default__.fm37(d_320_v13_, d_321_v14_, (d_322_v15_) + (d_322_v15_), globalState)
                    (globalState).f16 = (d_309_v4_)[default__.safeIndex((self).f30, len(d_309_v4_))]
                    d_323_v16_: _dafny.Seq
                    d_323_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdib"))
                    d_323_v16_ = d_323_v16_
                    r2 = p0
                    pass
            pass
        d_324_v17_: D3
        d_324_v17_ = D3_DC11(p0, True)
        d_325_i2_: int
        d_325_i2_ = 0
        with _dafny.label("2"):
            while not((not (not(p0)) or (p0) if (p0) and (p0) else (d_324_v17_).cf20)):
                with _dafny.c_label("2"):
                    if (d_325_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_325_i2_ = (d_325_i2_) + (1)
                    (globalState).f6 = default__.safeModuloInt((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcye")))) - ((0) - ((self).f30)), ((self).f30) - ((self).f30))
                    d_326_v18_: _dafny.Array
                    nw40_ = _dafny.Array(_dafny.CodePoint('D'), 8)
                    d_326_v18_ = nw40_
                    index32_ = default__.safeIndex(792, (d_326_v18_).length(0))
                    (d_326_v18_)[index32_] = default__.fm38(p0, (self).f30, globalState)
                    index33_ = default__.safeIndex(792, (d_326_v18_).length(0))
                    (d_326_v18_)[index33_] = d_312_v6_
                    d_327_v19_: _dafny.Seq
                    d_327_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whqyh"))
                    d_327_v19_ = d_327_v19_
                    d_328_v20_: _dafny.MultiSet
                    d_328_v20_ = _dafny.MultiSet([p0])
                    d_329_v21_: _dafny.Set
                    d_329_v21_ = _dafny.Set({(d_328_v20_).cardinality})
                    d_330_v22_: _dafny.Map
                    d_330_v22_ = _dafny.Map({p0: len(d_327_v19_)})
                    d_331_v23_: _dafny.Array
                    nw41_ = _dafny.Array(None, 24)
                    nw41_[int(0)] = d_327_v19_
                    nw41_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mx"))
                    nw41_[int(2)] = d_327_v19_
                    nw41_[int(3)] = (d_327_v19_ if p0 else d_327_v19_)
                    nw41_[int(4)] = d_327_v19_
                    nw41_[int(5)] = d_327_v19_
                    nw41_[int(6)] = (d_327_v19_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foa")))
                    nw41_[int(7)] = default__.fm36(p0, d_329_v21_, len(d_327_v19_), (self).f30, globalState)
                    nw41_[int(8)] = d_327_v19_
                    nw41_[int(9)] = d_327_v19_
                    nw41_[int(10)] = d_327_v19_
                    nw41_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))).set(default__.safeIndex((self).f30, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))), (d_326_v18_)[default__.safeIndex(792, (d_326_v18_).length(0))])
                    nw41_[int(12)] = d_327_v19_
                    nw41_[int(13)] = d_327_v19_
                    nw41_[int(14)] = default__.fm36(p0, d_329_v21_, ((d_330_v22_)[p0] if (p0) in (d_330_v22_) else (self).f30), (self).f30, globalState)
                    nw41_[int(15)] = d_327_v19_
                    nw41_[int(16)] = d_327_v19_
                    nw41_[int(17)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohl")) if p0 else d_327_v19_)
                    nw41_[int(18)] = (d_327_v19_) + (d_327_v19_)
                    nw41_[int(19)] = d_327_v19_
                    nw41_[int(20)] = (d_327_v19_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdespalc"))).set(default__.safeIndex((0) - ((self).f30), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdespalc")))), (d_326_v18_)[default__.safeIndex(792, (d_326_v18_).length(0))]))
                    nw41_[int(21)] = d_327_v19_
                    nw41_[int(22)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))) + (_dafny.SeqWithoutIsStrInference([d_312_v6_ for d_332_i3_ in range(default__.abs(433))]))
                    nw41_[int(23)] = d_327_v19_
                    d_331_v23_ = nw41_
                    index34_ = default__.safeIndex(9, (d_331_v23_).length(0))
                    (d_331_v23_)[index34_] = d_327_v19_
                    index35_ = default__.safeIndex(9, (d_331_v23_).length(0))
                    (d_331_v23_)[index35_] = (d_327_v19_) + (d_327_v19_)
                    pass
            pass
        (globalState).f16 = (self).f30
        d_333_v24_: _dafny.Seq
        d_333_v24_ = _dafny.SeqWithoutIsStrInference([p0, False])
        d_334_v25_: _dafny.Map
        d_334_v25_ = _dafny.Map({(self).f30: p0})
        d_335_v26_: _dafny.Seq
        d_335_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gewwn"))
        r0 = not (default__.fm37(d_333_v24_, d_334_v25_, (d_335_v26_).set(default__.safeIndex(len(d_333_v24_), len(d_335_v26_)), d_312_v6_), globalState)) or (not(p0))
        r1 = (self).f30
        r2 = not(p0)
        r3 = (_dafny.Set({len(d_335_v26_), 169})).intersection(_dafny.Set({(self).f30}))
        return r0, r1, r2, r3

    @property
    def f30(self):
        return self._f30

class C2(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self._f29: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f29, f22):
        (self)._f29 = f29
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return (self).f29

    def fm21(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qomukech"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eevpwlb")))

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_336_v0_: D3
        d_336_v0_ = D3_DC11(p0, (p0 if not(p0) else p0))
        source7_ = d_336_v0_
        if source7_.is_DC11:
            d_337___mcc_h0_ = source7_.cf19
            d_338___mcc_h1_ = source7_.cf20
            d_339_cf20_ = d_338___mcc_h1_
            d_340_cf19_ = d_337___mcc_h0_
            d_341_v1_: C0
            nw42_ = C0()
            nw42_.ctor__()
            d_341_v1_ = nw42_
            d_342_v2_: _dafny.Map
            d_342_v2_ = _dafny.Map({False: p0})
            d_343_v3_: _dafny.Map
            d_343_v3_ = _dafny.Map({not(((d_342_v2_)[d_339_cf20_] if (d_339_cf20_) in (d_342_v2_) else d_340_cf19_)): p0})
            d_344_v4_: _dafny.Map
            d_344_v4_ = _dafny.Map({p0: _dafny.Map({(self).f29: (self).f29})})
            d_345_v5_: str
            d_345_v5_ = _dafny.CodePoint('h')
            d_346_v6_: D0
            d_346_v6_ = D0_DC1(len(d_342_v2_), d_340_cf19_, d_344_v4_, (self).f29, d_345_v5_)
            d_347_v7_: D4
            d_347_v7_ = D4_DC13(d_346_v6_)
            source8_ = d_347_v7_
            if source8_.is_DC13:
                d_348___mcc_h3_ = source8_.cf22
                d_349_cf22_ = d_348___mcc_h3_
                d_350_v8_: _dafny.Seq
                d_350_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgkie"))
                d_351_v9_: _dafny.Map
                d_351_v9_ = _dafny.Map({(self).f29: d_339_cf20_})
                d_352_v10_: _dafny.Map
                d_352_v10_ = _dafny.Map({default__.fm22(d_339_cf20_, not(d_339_cf20_), (self).f29, len(default__.fm23((self).f29, len(d_350_v8_), not(False), not(d_340_cf19_), globalState)), globalState): (((d_351_v9_)[(0) - ((self).f29)] if ((0) - ((self).f29)) in (d_351_v9_) else False)) and (True)})
                d_353_v11_: _dafny.Set
                d_353_v11_ = _dafny.Set({d_339_cf20_})
                rhs20_ = (_dafny.Set({False})).isdisjoint((_dafny.Set({d_339_cf20_})).intersection(d_353_v11_))
                rhs21_ = d_352_v10_
                rhs22_ = (self).f29
                lhs14_ = globalState
                d_340_cf19_ = rhs20_
                d_352_v10_ = rhs21_
                lhs14_.f6 = rhs22_
                d_354_v12_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 10)
                d_354_v12_ = nw43_
                d_355_v13_: _dafny.Seq
                d_355_v13_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                d_356_v14_: _dafny.MultiSet
                d_356_v14_ = _dafny.MultiSet([d_339_cf20_, d_340_cf19_])
                d_357_v15_: _dafny.Seq
                d_357_v15_ = _dafny.SeqWithoutIsStrInference([False, p0, p0, d_340_cf19_])
                d_358_v16_: _dafny.Seq
                d_358_v16_ = _dafny.SeqWithoutIsStrInference([d_355_v13_, default__.fm24(d_356_v14_, 948, (self).f29, (d_357_v15_)[default__.safeIndex((self).f29, len(d_357_v15_))], globalState), d_355_v13_])
                index36_ = default__.safeIndex(685, (d_354_v12_).length(0))
                (d_354_v12_)[index36_] = len(d_358_v16_)
                d_359_v17_: _dafny.Array
                def lambda9_(d_360_i0_):
                    return False

                init4_ = lambda9_
                nw44_ = _dafny.Array(None, 3)
                for i0_4_ in range(nw44_.length(0)):
                    nw44_[i0_4_] = init4_(i0_4_)
                d_359_v17_ = nw44_
                index37_ = default__.safeIndex(162, (d_359_v17_).length(0))
                (d_359_v17_)[index37_] = not(d_339_cf20_)
                d_361_v18_: D3
                d_361_v18_ = D3_DC10(_dafny.Map({(self).f29: (self).f29}))
                d_362_v20_: _dafny.Set
                def iife33_():
                    coll29_ = _dafny.Map()
                    compr_29_: int
                    for compr_29_ in _dafny.IntegerRange(5, 124):
                        d_363_v19_: int = compr_29_
                        if ((5) <= (d_363_v19_)) and ((d_363_v19_) < (124)):
                            coll29_[default__.safeModuloInt(d_363_v19_, (self).f29)] = (self).f29
                    return _dafny.Map(coll29_)
                d_362_v20_ = _dafny.Set({d_361_v18_, d_361_v18_, D3_DC10(iife33_()
)})
                d_364_v21_: _dafny.MultiSet
                d_364_v21_ = _dafny.MultiSet([(self).f29])
                d_365_v22_: D6
                d_365_v22_ = D6_DC21(d_364_v21_)
                d_366_v23_: D1
                d_366_v23_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hk")))
                d_367_v24_: D1
                d_367_v24_ = D1_DC4(d_366_v23_)
                d_368_v25_: _dafny.Map
                d_368_v25_ = _dafny.Map({d_367_v24_: (self).f29})
                d_369_v26_: _dafny.Map
                d_369_v26_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiu")): ((d_368_v25_)[D1_DC4(d_366_v23_)] if (D1_DC4(d_366_v23_)) in (d_368_v25_) else (0) - ((self).f29))})
                d_370_v27_: _dafny.Map
                d_370_v27_ = _dafny.Map({_dafny.CodePoint('g'): D5_DC17(d_369_v26_)})
                d_371_v28_: _dafny.Set
                d_371_v28_ = _dafny.Set({d_346_v6_})
                d_372_v29_: _dafny.Set
                d_372_v29_ = _dafny.Set({d_355_v13_, d_355_v13_, d_355_v13_})
                d_373_v30_: D7
                d_373_v30_ = D7_DC25(d_372_v29_)
                index38_ = default__.safeIndex(685, (d_354_v12_).length(0))
                index39_ = default__.safeIndex(162, (d_359_v17_).length(0))
                rhs23_ = ((d_355_v13_)[default__.safeIndex((self).f29, len(d_355_v13_))]) - (((d_356_v14_)[d_339_cf20_] if (d_339_cf20_) in (d_356_v14_) else (self).f29))
                rhs24_ = (default__.fm25((self).f29, (self).f29, globalState)).issubset(d_362_v20_)
                rhs25_ = ((default__.fm26(((d_365_v22_).cf35).cardinality, d_370_v27_, p0, d_340_cf19_, globalState)) - (default__.fm26((self).f29, d_370_v27_, d_340_cf19_, (self).fm21(d_371_v28_, (self).f29, globalState), globalState))).isdisjoint((d_373_v30_).cf41)
                rhs26_ = 860
                lhs15_ = d_354_v12_
                lhs16_ = default__.safeIndex(685, (d_354_v12_).length(0))
                lhs17_ = d_359_v17_
                lhs18_ = default__.safeIndex(162, (d_359_v17_).length(0))
                lhs19_ = globalState
                lhs15_[lhs16_] = rhs23_
                d_340_cf19_ = rhs24_
                lhs17_[lhs18_] = rhs25_
                lhs19_.f6 = rhs26_
                d_374_v31_: D1
                d_374_v31_ = D1_DC3(d_350_v8_)
                pat_let_tv7_ = d_350_v8_
                pat_let_tv8_ = d_350_v8_
                d_375_v32_: _dafny.Array
                nw45_ = _dafny.Array(None, 26)
                nw45_[int(0)] = (d_374_v31_ if p0 else d_374_v31_)
                nw45_[int(1)] = d_374_v31_
                nw45_[int(2)] = D1_DC3(d_350_v8_)
                nw45_[int(3)] = d_374_v31_
                nw45_[int(4)] = d_374_v31_
                nw45_[int(5)] = d_374_v31_
                nw45_[int(6)] = d_374_v31_
                nw45_[int(7)] = d_374_v31_
                nw45_[int(8)] = D1_DC3((_dafny.SeqWithoutIsStrInference([d_345_v5_ for d_376_i1_ in range(default__.abs(-589))])).set(default__.safeIndex((self).f29, len(_dafny.SeqWithoutIsStrInference([d_345_v5_ for d_377_i1_ in range(default__.abs(-589))]))), d_345_v5_))
                nw45_[int(9)] = D1_DC3((d_341_v1_).fm17(d_350_v8_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwe"))), d_340_cf19_, len(default__.fm27(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gu"))), globalState)), globalState))
                nw45_[int(10)] = d_374_v31_
                nw45_[int(11)] = d_374_v31_
                nw45_[int(12)] = d_374_v31_
                nw45_[int(13)] = d_374_v31_
                nw45_[int(14)] = d_374_v31_
                nw45_[int(15)] = d_374_v31_
                def iife34_(_pat_let2_0):
                    def iife35_(d_378_dt__update__tmp_h0_):
                        def iife36_(_pat_let3_0):
                            def iife37_(d_379_dt__update_hcf6_h0_):
                                return D1_DC3(d_379_dt__update_hcf6_h0_)
                            return iife37_(_pat_let3_0)
                        return iife36_(pat_let_tv7_)
                    return iife35_(_pat_let2_0)
                nw45_[int(16)] = iife34_(d_374_v31_)
                nw45_[int(17)] = d_374_v31_
                nw45_[int(18)] = d_374_v31_
                nw45_[int(19)] = d_374_v31_
                nw45_[int(20)] = d_374_v31_
                nw45_[int(21)] = d_374_v31_
                def iife38_(_pat_let4_0):
                    def iife39_(d_380_dt__update__tmp_h1_):
                        def iife40_(_pat_let5_0):
                            def iife41_(d_381_dt__update_hcf6_h1_):
                                return D1_DC3(d_381_dt__update_hcf6_h1_)
                            return iife41_(_pat_let5_0)
                        return iife40_(pat_let_tv8_)
                    return iife39_(_pat_let4_0)
                nw45_[int(22)] = iife38_(d_374_v31_)
                nw45_[int(23)] = d_374_v31_
                nw45_[int(24)] = d_374_v31_
                nw45_[int(25)] = d_374_v31_
                d_375_v32_ = nw45_
                index40_ = default__.safeIndex(453, (d_375_v32_).length(0))
                (d_375_v32_)[index40_] = d_374_v31_
                index41_ = default__.safeIndex(453, (d_375_v32_).length(0))
                (d_375_v32_)[index41_] = d_374_v31_
                d_382_v33_: C0
                nw46_ = C0()
                nw46_.ctor__()
                d_382_v33_ = nw46_
            elif source8_.is_DC14:
                d_383___mcc_h4_ = source8_.cf23
                d_384___mcc_h5_ = source8_.cf24
                d_385___mcc_h6_ = source8_.cf25
                d_386_cf25_ = d_385___mcc_h6_
                d_387_cf24_ = d_384___mcc_h5_
                d_388_cf23_ = d_383___mcc_h4_
                d_343_v3_ = (d_343_v3_).set((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_389_i2_ in range(default__.abs(975))])) < (d_386_cf25_), p0)
                d_390_v34_: _dafny.Map
                d_390_v34_ = _dafny.Map({(self).f29: (_dafny.SeqWithoutIsStrInference([d_388_cf23_ for d_391_i3_ in range(default__.abs(775))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeuvryom")))})
                d_392_v35_: _dafny.MultiSet
                d_392_v35_ = _dafny.MultiSet([not(p0), p0])
                d_390_v34_ = (d_390_v34_).set(((self).f29 if False else (self).f29), (d_392_v35_).ispropersubset(d_392_v35_))
                d_393_v36_: _dafny.MultiSet
                d_393_v36_ = _dafny.MultiSet([(self).f29, (self).f29, (self).f29, (self).f29])
                d_394_v37_: _dafny.Set
                d_394_v37_ = _dafny.Set({(self).f29})
                d_395_v38_: _dafny.Seq
                d_395_v38_ = _dafny.SeqWithoutIsStrInference([d_394_v37_, d_394_v37_, d_394_v37_, d_394_v37_])
                d_396_v39_: _dafny.Array
                nw47_ = _dafny.Array(None, 7)
                nw47_[int(0)] = len((_dafny.SeqWithoutIsStrInference([d_388_cf23_ for d_397_i4_ in range(default__.abs(991))])) + (d_386_cf25_))
                nw47_[int(1)] = default__.safeModuloInt((self).fm0(d_393_v36_, (self).f29, d_395_v38_, p0, globalState), (0) - ((self).f29))
                nw47_[int(2)] = (self).f29
                nw47_[int(3)] = -527
                nw47_[int(4)] = ((self).f29 if d_339_cf20_ else (self).f29)
                nw47_[int(5)] = (d_393_v36_).cardinality
                nw47_[int(6)] = (self).f29
                d_396_v39_ = nw47_
                d_398_v40_: _dafny.Map
                d_398_v40_ = _dafny.Map({d_339_cf20_: (self).f29})
                index42_ = default__.safeIndex(638, (d_396_v39_).length(0))
                (d_396_v39_)[index42_] = ((d_398_v40_)[d_340_cf19_] if (d_340_cf19_) in (d_398_v40_) else (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_386_cf25_ for d_399_i5_ in range(default__.abs(953))])).set(default__.safeIndex((self).f29, len(_dafny.SeqWithoutIsStrInference([d_386_cf25_ for d_400_i5_ in range(default__.abs(953))]))), d_386_cf25_))).cardinality)
                index43_ = default__.safeIndex(638, (d_396_v39_).length(0))
                rhs27_ = default__.safeDivisionInt((self).f29, (0) - ((self).f29))
                rhs28_ = d_339_cf20_
                rhs29_ = not (d_340_cf19_) or (p0)
                rhs30_ = (self).f29
                rhs31_ = d_387_cf24_
                lhs20_ = globalState
                lhs21_ = globalState
                lhs22_ = globalState
                lhs23_ = d_396_v39_
                lhs24_ = default__.safeIndex(638, (d_396_v39_).length(0))
                lhs20_.f11 = rhs27_
                lhs21_.f10 = rhs28_
                lhs22_.f10 = rhs29_
                lhs23_[lhs24_] = rhs30_
                d_387_cf24_ = rhs31_
                d_401_v41_: int
                d_402_v42_: bool
                d_403_v43_: bool
                d_404_v44_: int
                out34_: int
                out35_: bool
                out36_: bool
                out37_: int
                out34_, out35_, out36_, out37_ = (self).m5((0) - ((self).f29), (self).f29, globalState)
                d_401_v41_ = out34_
                d_402_v42_ = out35_
                d_403_v43_ = out36_
                d_404_v44_ = out37_
            elif source8_.is_DC15:
                d_405___mcc_h7_ = source8_.cf26
                d_406___mcc_h8_ = source8_.cf27
                d_407_cf27_ = d_406___mcc_h8_
                d_408_cf26_ = d_405___mcc_h7_
                d_409_v45_: _dafny.Map
                d_409_v45_ = _dafny.Map({(self).f29: (self).f29})
                d_410_v46_: _dafny.Map
                d_410_v46_ = _dafny.Map({d_407_cf27_: D6_DC23(d_339_cf20_, ((d_409_v45_)[42] if (42) in (d_409_v45_) else (self).f29), p0)})
                d_411_v47_: D6
                d_411_v47_ = D6_DC23(p0, d_407_cf27_, False)
                d_410_v46_ = (d_410_v46_).set(144, d_411_v47_)
                (globalState).f16 = d_408_cf26_
                (globalState).f6 = default__.safeDivisionInt(d_408_cf26_, default__.safeModuloInt(d_408_cf26_, d_408_cf26_))
                d_412_v48_: _dafny.MultiSet
                d_412_v48_ = _dafny.MultiSet([d_408_cf26_])
                d_413_v49_: _dafny.Array
                nw48_ = _dafny.Array(False, 22)
                d_413_v49_ = nw48_
                d_414_v50_: _dafny.Map
                d_414_v50_ = _dafny.Map({d_340_cf19_: (self).f29})
                d_415_v51_: D4
                d_415_v51_ = D4_DC14(d_345_v5_, d_413_v49_, ((d_341_v1_).fm17(_dafny.SeqWithoutIsStrInference([d_345_v5_ for d_416_i6_ in range(default__.abs(897))]), 496, d_339_cf20_, d_408_cf26_, globalState)).set(default__.safeIndex(len(d_414_v50_), len((d_341_v1_).fm17(_dafny.SeqWithoutIsStrInference([d_345_v5_ for d_417_i6_ in range(default__.abs(897))]), 496, d_339_cf20_, d_408_cf26_, globalState))), d_345_v5_))
                d_418_v52_: D4
                d_418_v52_ = D4_DC16(d_415_v51_)
                d_419_v53_: _dafny.Map
                d_419_v53_ = _dafny.Map({False: d_418_v52_})
                d_420_v55_: _dafny.Set
                d_420_v55_ = _dafny.Set({124})
                def iife42_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(742, -148):
                        d_421_v54_: int = compr_30_
                        if ((742) <= (d_421_v54_)) and ((d_421_v54_) < (-148)):
                            coll30_ = coll30_.union(_dafny.Set([default__.safeModuloInt(d_421_v54_, d_408_cf26_)]))
                    return _dafny.Set(coll30_)
                (globalState).f11 = (self).fm0(d_412_v48_, len(d_419_v53_), (_dafny.SeqWithoutIsStrInference([iife42_()
 for d_422_i7_ in range(default__.abs(576))])) + (_dafny.SeqWithoutIsStrInference([d_420_v55_, d_420_v55_])), p0, globalState)
            elif source8_.is_DC12:
                d_423___mcc_h9_ = source8_.cf21
                d_424_cf21_ = d_423___mcc_h9_
                (globalState).f10 = d_340_cf19_
                d_425_v56_: _dafny.Map
                d_425_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_426_i8_ in range(default__.abs(960))]): (self).f29})
                d_427_v57_: _dafny.Seq
                d_427_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "delahctx"))
                d_428_v58_: D5
                d_428_v58_ = D5_DC17((d_425_v56_).set(d_427_v57_, (self).f29))
                d_429_v59_: D5
                d_429_v59_ = D5_DC20(d_428_v58_)
                d_429_v59_ = d_429_v59_
                d_430_v60_: _dafny.Array
                nw49_ = _dafny.Array(False, 12)
                d_430_v60_ = nw49_
                index44_ = default__.safeIndex(296, (d_430_v60_).length(0))
                (d_430_v60_)[index44_] = False
                d_431_v61_: _dafny.MultiSet
                d_431_v61_ = _dafny.MultiSet([(self).f29, (self).f29])
                index45_ = default__.safeIndex(296, (d_430_v60_).length(0))
                (d_430_v60_)[index45_] = (d_431_v61_).isdisjoint(d_431_v61_)
                d_432_v62_: _dafny.Array
                nw50_ = _dafny.Array(D3.default()(), 12)
                d_432_v62_ = nw50_
                d_433_v63_: _dafny.Map
                d_433_v63_ = _dafny.Map({(self).f29: (self).f29})
                d_434_v64_: D3
                d_434_v64_ = D3_DC10(d_433_v63_)
                index46_ = default__.safeIndex(249, (d_432_v62_).length(0))
                (d_432_v62_)[index46_] = d_434_v64_
                d_435_v65_: _dafny.Set
                d_435_v65_ = _dafny.Set({d_346_v6_})
                d_436_v66_: _dafny.Map
                d_436_v66_ = _dafny.Map({d_427_v57_: (d_430_v60_)[default__.safeIndex(296, (d_430_v60_).length(0))]})
                d_437_v67_: _dafny.Set
                d_437_v67_ = _dafny.Set({len(d_436_v66_)})
                d_438_v68_: _dafny.Map
                d_438_v68_ = _dafny.Map({p0: (self).f29})
                d_439_v69_: _dafny.Seq
                d_439_v69_ = _dafny.SeqWithoutIsStrInference([p0, True])
                index47_ = default__.safeIndex(249, (d_432_v62_).length(0))
                rhs32_ = D3_DC10(default__.fm28(d_339_cf20_, d_340_cf19_, not((self).fm21(d_435_v65_, len(d_437_v67_), globalState)), d_438_v68_, globalState))
                rhs33_ = p0
                rhs34_ = not((d_439_v69_)[default__.safeIndex((self).f29, len(d_439_v69_))])
                lhs25_ = d_432_v62_
                lhs26_ = default__.safeIndex(249, (d_432_v62_).length(0))
                lhs27_ = globalState
                lhs25_[lhs26_] = rhs32_
                lhs27_.f0 = rhs33_
                d_340_cf19_ = rhs34_
            elif True:
                d_440___mcc_h10_ = source8_.cf28
                d_441_cf28_ = d_440___mcc_h10_
                d_442_v70_: _dafny.Array
                def lambda10_(d_443_i9_):
                    return (d_443_i9_) - (len(_dafny.SeqWithoutIsStrInference([(self).f29])))

                init5_ = lambda10_
                nw51_ = _dafny.Array(None, 2)
                for i0_5_ in range(nw51_.length(0)):
                    nw51_[i0_5_] = init5_(i0_5_)
                d_442_v70_ = nw51_
                d_442_v70_ = (d_442_v70_ if (d_345_v5_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "roqfhujdu"))) else d_442_v70_)
                d_444_v71_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Seq({}), 6)
                d_444_v71_ = nw52_
                d_445_v72_: _dafny.Map
                d_445_v72_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsdcwur")): d_444_v71_})
                d_445_v72_ = (d_445_v72_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlmwfrvcm")), d_444_v71_)
                (globalState).f8 = (self).f29
                d_446_v73_: _dafny.Seq
                d_446_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frayqgv"))
                rhs35_ = ((self).f29) < ((self).f29)
                rhs36_ = d_442_v70_
                rhs37_ = d_339_cf20_
                rhs38_ = (d_341_v1_).fm17(d_446_v73_, (self).f29, (d_446_v73_) <= (d_446_v73_), (self).f29, globalState)
                r0 = rhs35_
                d_442_v70_ = rhs36_
                d_340_cf19_ = rhs37_
                d_446_v73_ = rhs38_
            d_447_v74_: _dafny.Array
            nw53_ = _dafny.Array(_dafny.Array(None, 0), 13)
            d_447_v74_ = nw53_
            d_448_v75_: _dafny.Seq
            d_448_v75_ = _dafny.SeqWithoutIsStrInference([p0])
            d_449_v76_: _dafny.Map
            d_449_v76_ = _dafny.Map({d_447_v74_: not(not ((d_448_v75_)[default__.safeIndex((self).f29, len(d_448_v75_))]) or (p0))})
            d_449_v76_ = (d_449_v76_).set(d_447_v74_, p0)
            d_450_v77_: _dafny.Array
            nw54_ = _dafny.Array(False, 26)
            d_450_v77_ = nw54_
            index48_ = default__.safeIndex(762, (d_450_v77_).length(0))
            (d_450_v77_)[index48_] = (d_339_cf20_ if p0 else False)
            d_451_v78_: _dafny.Seq
            d_451_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xljlu"))
            index49_ = default__.safeIndex(762, (d_450_v77_).length(0))
            rhs39_ = len(d_451_v78_)
            rhs40_ = d_340_cf19_
            lhs28_ = globalState
            lhs29_ = d_450_v77_
            lhs30_ = default__.safeIndex(762, (d_450_v77_).length(0))
            lhs28_.f11 = rhs39_
            lhs29_[lhs30_] = rhs40_
        elif True:
            d_452___mcc_h2_ = source7_.cf18
            d_453_cf18_ = d_452___mcc_h2_
            d_454_v79_: D2
            d_454_v79_ = D2_DC6((self).f29)
            d_455_v80_: D2
            d_455_v80_ = D2_DC9(d_454_v79_)
            d_456_v81_: str
            d_456_v81_ = _dafny.CodePoint('j')
            d_457_v82_: _dafny.Array
            nw55_ = _dafny.Array(D2.default()(), 18)
            d_457_v82_ = nw55_
            d_458_v83_: _dafny.MultiSet
            d_458_v83_ = _dafny.MultiSet([d_457_v82_])
            d_459_v84_: _dafny.Map
            d_459_v84_ = _dafny.Map({d_456_v81_: d_458_v83_})
            d_460_v85_: _dafny.Map
            d_460_v85_ = _dafny.Map({d_455_v80_: ((d_459_v84_)[d_456_v81_] if (d_456_v81_) in (d_459_v84_) else (d_458_v83_).set(d_457_v82_, default__.abs((self).f29)))})
            d_461_v86_: _dafny.Seq
            d_461_v86_ = _dafny.SeqWithoutIsStrInference([d_458_v83_])
            d_460_v85_ = (d_460_v85_).set(d_455_v80_, ((d_461_v86_)[default__.safeIndex((self).f29, len(d_461_v86_))]).intersection(d_458_v83_))
            d_462_v87_: _dafny.Map
            d_462_v87_ = _dafny.Map({(self).f29: p0})
            d_463_v88_: _dafny.Seq
            d_463_v88_ = _dafny.SeqWithoutIsStrInference([((d_462_v87_)[(self).f29] if ((self).f29) in (d_462_v87_) else p0)])
            d_464_v89_: _dafny.Seq
            d_464_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lklrnmhq"))
            d_465_v90_: _dafny.Array
            nw56_ = _dafny.Array(None, 14)
            nw56_[int(0)] = p0
            nw56_[int(1)] = (d_463_v88_)[default__.safeIndex((_dafny.MultiSet([len(d_464_v89_)])).cardinality, len(d_463_v88_))]
            nw56_[int(2)] = not(p0)
            nw56_[int(3)] = True
            nw56_[int(4)] = p0
            nw56_[int(5)] = (d_463_v88_)[default__.safeIndex((self).f29, len(d_463_v88_))]
            nw56_[int(6)] = p0
            nw56_[int(7)] = p0
            nw56_[int(8)] = True
            nw56_[int(9)] = p0
            nw56_[int(10)] = p0
            nw56_[int(11)] = p0
            nw56_[int(12)] = p0
            nw56_[int(13)] = p0
            d_465_v90_ = nw56_
            d_466_v91_: _dafny.Array
            nw57_ = _dafny.Array(None, 6)
            nw57_[int(0)] = d_465_v90_
            nw57_[int(1)] = d_465_v90_
            nw57_[int(2)] = d_465_v90_
            nw57_[int(3)] = (d_465_v90_ if p0 else d_465_v90_)
            nw57_[int(4)] = d_465_v90_
            nw57_[int(5)] = d_465_v90_
            d_466_v91_ = nw57_
            index50_ = default__.safeIndex(37, (d_466_v91_).length(0))
            (d_466_v91_)[index50_] = d_465_v90_
            index51_ = default__.safeIndex(37, (d_466_v91_).length(0))
            def lambda11_(d_467_p0_):
                def lambda12_(d_468_i10_):
                    return d_467_p0_

                return lambda12_

            init6_ = lambda11_(p0)
            nw58_ = _dafny.Array(None, 14)
            for i0_6_ in range(nw58_.length(0)):
                nw58_[i0_6_] = init6_(i0_6_)
            (d_466_v91_)[index51_] = nw58_
            d_469_v92_: _dafny.MultiSet
            d_469_v92_ = _dafny.MultiSet([p0])
            d_470_v93_: int
            d_471_v94_: bool
            d_472_v95_: bool
            d_473_v96_: int
            out38_: int
            out39_: bool
            out40_: bool
            out41_: int
            out38_, out39_, out40_, out41_ = (self).m5((0) - ((self).f29), (0) - (default__.safeModuloInt(((d_453_cf18_)[(d_469_v92_).cardinality] if ((d_469_v92_).cardinality) in (d_453_cf18_) else (_dafny.MultiSet([p0])).cardinality), (self).f29)), globalState)
            d_470_v93_ = out38_
            d_471_v94_ = out39_
            d_472_v95_ = out40_
            d_473_v96_ = out41_
            d_474_v97_: _dafny.Array
            def lambda13_(d_475_v96_):
                def lambda14_(d_476_i11_):
                    return (d_476_i11_) * (d_475_v96_)

                return lambda14_

            init7_ = lambda13_(d_473_v96_)
            nw59_ = _dafny.Array(None, 23)
            for i0_7_ in range(nw59_.length(0)):
                nw59_[i0_7_] = init7_(i0_7_)
            d_474_v97_ = nw59_
            d_477_v98_: _dafny.Seq
            d_477_v98_ = _dafny.SeqWithoutIsStrInference([d_474_v97_, d_474_v97_, d_474_v97_])
            d_478_v99_: C0
            nw60_ = C0()
            nw60_.ctor__()
            d_478_v99_ = nw60_
            d_474_v97_ = (d_477_v98_)[default__.safeIndex(len(_dafny.Map({d_470_v93_: d_478_v99_})), len(d_477_v98_))]
        d_479_v100_: _dafny.Array
        nw61_ = _dafny.Array(_dafny.Array(None, 0), 17)
        d_479_v100_ = nw61_
        d_480_v101_: _dafny.Array
        nw62_ = _dafny.Array(False, 16)
        d_480_v101_ = nw62_
        nw63_ = _dafny.Array(None, 7)
        nw63_[int(0)] = d_480_v101_
        nw63_[int(1)] = d_480_v101_
        nw63_[int(2)] = d_480_v101_
        nw63_[int(3)] = d_480_v101_
        nw63_[int(4)] = ((self).f22)[default__.safeIndex((self).f29, len((self).f22))]
        nw63_[int(5)] = d_480_v101_
        nw63_[int(6)] = d_480_v101_
        d_479_v100_ = nw63_
        r3 = (default__.safeDivisionInt((self).f29, (default__.fm29(p0, 277, p0, (self).f29, globalState)).cardinality)) - (638)
        hi1_ = 388
        for d_481_i12_ in range(((self).f29 if p0 else (self).f29), hi1_):
            d_482_v102_: str
            d_482_v102_ = _dafny.CodePoint('u')
            d_483_v103_: _dafny.Seq
            d_483_v103_ = _dafny.SeqWithoutIsStrInference([d_482_v102_])
            d_483_v103_ = (d_483_v103_) + (d_483_v103_)
            (globalState).f16 = d_481_i12_
            (globalState).f6 = d_481_i12_
            d_484_v104_: _dafny.MultiSet
            d_484_v104_ = _dafny.MultiSet([(self).f29, 680])
            d_485_v106_: _dafny.Seq
            d_485_v106_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_486_v107_: _dafny.Seq
            d_486_v107_ = _dafny.SeqWithoutIsStrInference([len(d_485_v106_), (self).f29])
            d_487_v108_: D4
            d_487_v108_ = D4_DC14(d_482_v102_, d_480_v101_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxovnuk")))
            d_488_v109_: _dafny.MultiSet
            d_488_v109_ = _dafny.MultiSet([d_487_v108_])
            d_489_v110_: _dafny.Map
            d_489_v110_ = _dafny.Map({p0: (d_488_v109_).cardinality})
            d_490_v111_: _dafny.Set
            d_490_v111_ = _dafny.Set({12, d_481_i12_, (d_486_v107_)[default__.safeIndex(len(d_486_v107_), len(d_486_v107_))]})
            d_491_v112_: _dafny.Map
            d_491_v112_ = _dafny.Map({p0: d_490_v111_})
            d_492_v113_: _dafny.Seq
            d_492_v113_ = _dafny.SeqWithoutIsStrInference([((d_491_v112_)[p0] if (p0) in (d_491_v112_) else d_490_v111_), d_490_v111_])
            d_493_v114_: _dafny.Set
            d_493_v114_ = _dafny.Set({(self).fm0(_dafny.MultiSet(d_486_v107_), len(d_489_v110_), d_492_v113_, False, globalState)})
            d_494_v115_: _dafny.Seq
            def iife43_():
                coll31_ = _dafny.Set()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(984, 889):
                    d_495_v105_: int = compr_31_
                    if ((984) <= (d_495_v105_)) and ((d_495_v105_) < (889)):
                        coll31_ = coll31_.union(_dafny.Set([(d_495_v105_) + (d_481_i12_)]))
                return _dafny.Set(coll31_)
            d_494_v115_ = _dafny.SeqWithoutIsStrInference([iife43_()
            , d_493_v114_])
            d_496_v116_: D6
            d_496_v116_ = D6_DC22((self).fm0(d_484_v104_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekh"))), d_492_v113_, p0, globalState))
            d_497_v117_: _dafny.Map
            d_497_v117_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxitwh"))): (self).f29})
            d_498_v118_: _dafny.Map
            d_498_v118_ = _dafny.Map({p0: d_497_v117_})
            d_499_v119_: D0
            d_499_v119_ = D0_DC1((d_496_v116_).cf36, p0, d_498_v118_, d_481_i12_, _dafny.CodePoint('h'))
            d_500_v120_: D4
            d_500_v120_ = D4_DC13(d_499_v119_)
            d_500_v120_ = d_500_v120_
        d_501_v121_: _dafny.Set
        d_501_v121_ = _dafny.Set({p0})
        if (d_501_v121_).issubset(d_501_v121_):
            d_502_v122_: _dafny.Seq
            d_502_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bknlragrj"))
            d_503_v123_: _dafny.Map
            d_503_v123_ = _dafny.Map({p0: len(d_502_v122_)})
            (globalState).f6 = default__.safeDivisionInt(69, len(d_503_v123_))
            d_504_v124_: _dafny.Map
            d_504_v124_ = _dafny.Map({default__.fm30(globalState): False})
            (globalState).f0 = (((d_503_v123_)[p0] if (p0) in (d_503_v123_) else (self).f29)) < ((0) - (len((d_504_v124_) | (d_504_v124_))))
            d_505_v125_: _dafny.Array
            nw64_ = _dafny.Array(int(0), 17)
            d_505_v125_ = nw64_
            d_506_v127_: _dafny.Map
            def iife44_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(-266, 398):
                    d_507_v126_: int = compr_32_
                    if ((-266) <= (d_507_v126_)) and ((d_507_v126_) < (398)):
                        coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_507_v126_, (0) - ((self).f29))]))
                return _dafny.Set(coll32_)
            d_506_v127_ = _dafny.Map({(iife44_()
            ) in (_dafny.Set({_dafny.Set({(self).f29})})): d_505_v125_})
            d_505_v125_ = ((d_506_v127_)[p0] if (p0) in (d_506_v127_) else d_505_v125_)
            d_508_v128_: D2
            d_508_v128_ = D2_DC5(d_505_v125_)
            pat_let_tv9_ = d_505_v125_
            d_509_v129_: _dafny.Map
            def iife45_(_pat_let6_0):
                def iife46_(d_510_dt__update__tmp_h2_):
                    def iife47_(_pat_let7_0):
                        def iife48_(d_511_dt__update_hcf8_h0_):
                            return D2_DC5(d_511_dt__update_hcf8_h0_)
                        return iife48_(_pat_let7_0)
                    return iife47_(pat_let_tv9_)
                return iife46_(_pat_let6_0)
            d_509_v129_ = _dafny.Map({d_480_v101_: (iife45_(d_508_v128_)).cf8})
            (globalState).f6 = (0) - (len((d_509_v129_) | ((_dafny.Map({d_480_v101_: d_505_v125_})).set(d_480_v101_, d_505_v125_))))
            if not (p0) or (p0):
                d_512_v130_: _dafny.MultiSet
                d_512_v130_ = _dafny.MultiSet([(self).f29])
                (globalState).f10 = ((d_512_v130_).isdisjoint(d_512_v130_)) == (p0)
                d_513_v131_: _dafny.Set
                d_513_v131_ = _dafny.Set({d_480_v101_, d_480_v101_, d_480_v101_})
                d_513_v131_ = (d_513_v131_ if p0 else (_dafny.Set({d_480_v101_})).intersection(d_513_v131_))
                index52_ = default__.safeIndex(720, (d_505_v125_).length(0))
                (d_505_v125_)[index52_] = (self).f29
                index53_ = default__.safeIndex(569, (d_480_v101_).length(0))
                (d_480_v101_)[index53_] = p0
                index54_ = default__.safeIndex(748, (d_505_v125_).length(0))
                (d_505_v125_)[index54_] = (self).f29
                d_514_v132_: _dafny.Seq
                d_514_v132_ = _dafny.SeqWithoutIsStrInference([p0])
                d_515_v133_: _dafny.Set
                d_515_v133_ = _dafny.Set({686, 316})
                d_516_v134_: _dafny.Seq
                d_516_v134_ = _dafny.SeqWithoutIsStrInference([(self).fm0(d_512_v130_, len(d_514_v132_), _dafny.SeqWithoutIsStrInference([d_515_v133_]), p0, globalState)])
                d_517_v135_: D6
                d_517_v135_ = D6_DC22((self).f29)
                index55_ = default__.safeIndex(720, (d_505_v125_).length(0))
                index56_ = default__.safeIndex(569, (d_480_v101_).length(0))
                index57_ = default__.safeIndex(748, (d_505_v125_).length(0))
                rhs41_ = len(d_516_v134_)
                rhs42_ = (default__.fm31(False, globalState)).cf27
                rhs43_ = p0
                rhs44_ = ((d_517_v135_ if p0 else d_517_v135_)).cf36
                lhs31_ = d_505_v125_
                lhs32_ = default__.safeIndex(720, (d_505_v125_).length(0))
                lhs33_ = globalState
                lhs34_ = d_480_v101_
                lhs35_ = default__.safeIndex(569, (d_480_v101_).length(0))
                lhs36_ = d_505_v125_
                lhs37_ = default__.safeIndex(748, (d_505_v125_).length(0))
                lhs31_[lhs32_] = rhs41_
                lhs33_.f11 = rhs42_
                lhs34_[lhs35_] = rhs43_
                lhs36_[lhs37_] = rhs44_
                d_518_v136_: _dafny.Seq
                d_518_v136_ = _dafny.SeqWithoutIsStrInference([(d_515_v133_) | (d_515_v133_)])
                d_518_v136_ = (d_518_v136_) + (d_518_v136_)
                d_519_v137_: _dafny.Array
                nw65_ = _dafny.Array(int(0), 17)
                d_519_v137_ = nw65_
                d_520_v138_: _dafny.Seq
                d_520_v138_ = _dafny.SeqWithoutIsStrInference([d_519_v137_])
                d_521_v139_: D8
                d_521_v139_ = D8_DC28(d_520_v138_)
                d_522_v140_: D8
                d_522_v140_ = D8_DC28((d_521_v139_).cf46)
                d_523_v141_: _dafny.Array
                nw66_ = _dafny.Array(None, 21)
                nw66_[int(0)] = d_520_v138_
                nw66_[int(1)] = d_520_v138_
                nw66_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_519_v137_]) if p0 else _dafny.SeqWithoutIsStrInference([d_519_v137_]))
                nw66_[int(3)] = d_520_v138_
                nw66_[int(4)] = _dafny.SeqWithoutIsStrInference([d_519_v137_, d_519_v137_, d_519_v137_])
                nw66_[int(5)] = d_520_v138_
                nw66_[int(6)] = d_520_v138_
                nw66_[int(7)] = (d_520_v138_) + (d_520_v138_)
                nw66_[int(8)] = d_520_v138_
                nw66_[int(9)] = (d_522_v140_).cf46
                nw66_[int(10)] = d_520_v138_
                nw66_[int(11)] = (d_520_v138_) + (d_520_v138_)
                nw66_[int(12)] = d_520_v138_
                nw66_[int(13)] = _dafny.SeqWithoutIsStrInference([d_519_v137_])
                nw66_[int(14)] = (d_520_v138_) + (_dafny.SeqWithoutIsStrInference([d_519_v137_, d_519_v137_]))
                nw66_[int(15)] = (d_520_v138_) + (d_520_v138_)
                nw66_[int(16)] = (d_520_v138_) + (d_520_v138_)
                nw66_[int(17)] = d_520_v138_
                nw66_[int(18)] = d_520_v138_
                nw66_[int(19)] = d_520_v138_
                nw66_[int(20)] = (d_520_v138_) + (d_520_v138_)
                d_523_v141_ = nw66_
                index58_ = default__.safeIndex(303, (d_523_v141_).length(0))
                (d_523_v141_)[index58_] = d_520_v138_
                index59_ = default__.safeIndex(303, (d_523_v141_).length(0))
                (d_523_v141_)[index59_] = ((d_520_v138_) + (_dafny.SeqWithoutIsStrInference([d_519_v137_, d_519_v137_, d_519_v137_, d_519_v137_, d_519_v137_]))) + (_dafny.SeqWithoutIsStrInference([d_519_v137_, d_519_v137_, d_519_v137_, d_519_v137_, ((d_506_v127_)[p0] if (p0) in (d_506_v127_) else d_519_v137_)]))
            elif True:
                (globalState).f8 = (self).f29
                d_524_v142_: _dafny.Array
                nw67_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_524_v142_ = nw67_
                d_525_v143_: _dafny.Seq
                d_525_v143_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])
                d_526_v144_: _dafny.Set
                d_526_v144_ = _dafny.Set({d_525_v143_})
                d_527_v145_: D7
                d_527_v145_ = D7_DC25(d_526_v144_)
                pat_let_tv10_ = d_525_v143_
                pat_let_tv11_ = d_525_v143_
                d_528_v146_: _dafny.Array
                nw68_ = _dafny.Array(None, 24)
                nw68_[int(0)] = d_527_v145_
                nw68_[int(1)] = d_527_v145_
                nw68_[int(2)] = default__.fm32(globalState)
                nw68_[int(3)] = D7_DC25(_dafny.Set({d_525_v143_}))
                nw68_[int(4)] = d_527_v145_
                nw68_[int(5)] = d_527_v145_
                nw68_[int(6)] = D7_DC25(d_526_v144_)
                nw68_[int(7)] = d_527_v145_
                nw68_[int(8)] = D7_DC25(d_526_v144_)
                nw68_[int(9)] = d_527_v145_
                nw68_[int(10)] = d_527_v145_
                nw68_[int(11)] = d_527_v145_
                nw68_[int(12)] = D7_DC25(d_526_v144_)
                nw68_[int(13)] = d_527_v145_
                nw68_[int(14)] = d_527_v145_
                nw68_[int(15)] = d_527_v145_
                def iife49_(_pat_let8_0):
                    def iife50_(d_529_dt__update__tmp_h3_):
                        def iife51_(_pat_let9_0):
                            def iife52_(d_530_dt__update_hcf41_h0_):
                                return D7_DC25(d_530_dt__update_hcf41_h0_)
                            return iife52_(_pat_let9_0)
                        return iife51_(_dafny.Set({pat_let_tv10_, pat_let_tv11_}))
                    return iife50_(_pat_let8_0)
                nw68_[int(16)] = iife49_(d_527_v145_)
                nw68_[int(17)] = d_527_v145_
                nw68_[int(18)] = d_527_v145_
                nw68_[int(19)] = d_527_v145_
                nw68_[int(20)] = d_527_v145_
                nw68_[int(21)] = d_527_v145_
                nw68_[int(22)] = D7_DC25(d_526_v144_)
                nw68_[int(23)] = d_527_v145_
                d_528_v146_ = nw68_
                index60_ = default__.safeIndex(662, (d_524_v142_).length(0))
                (d_524_v142_)[index60_] = d_528_v146_
                d_531_v147_: _dafny.Seq
                d_531_v147_ = _dafny.SeqWithoutIsStrInference([d_528_v146_, d_528_v146_])
                index61_ = default__.safeIndex(662, (d_524_v142_).length(0))
                (d_524_v142_)[index61_] = (d_531_v147_)[default__.safeIndex(((d_503_v123_)[p0] if (p0) in (d_503_v123_) else -113), len(d_531_v147_))]
                d_532_v148_: D2
                d_532_v148_ = D2_DC6((self).f29)
                d_533_v149_: _dafny.Map
                d_533_v149_ = _dafny.Map({((self).f29) * (-50): ((d_532_v148_).cf9) * ((self).f29)})
                d_533_v149_ = (d_533_v149_).set(default__.safeModuloInt((self).f29, len(d_502_v122_)), (self).f29)
                index62_ = default__.safeIndex(152, (d_480_v101_).length(0))
                (d_480_v101_)[index62_] = p0
                d_534_v150_: _dafny.MultiSet
                d_534_v150_ = _dafny.MultiSet([(self).f29, (self).f29])
                d_535_v151_: _dafny.Seq
                d_535_v151_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({len(d_502_v122_)})])
                index63_ = default__.safeIndex(152, (d_480_v101_).length(0))
                (d_480_v101_)[index63_] = ((len(d_502_v122_)) >= ((self).fm0(d_534_v150_, len(d_502_v122_), (d_535_v151_).set(default__.safeIndex((self).f29, len(d_535_v151_)), _dafny.Set({len(default__.fm33(globalState))})), not(p0), globalState))) == (True)
                d_536_v152_: _dafny.Array
                nw69_ = _dafny.Array(False, 27)
                d_536_v152_ = nw69_
                d_537_v153_: T0
                nw70_ = C1()
                nw70_.ctor__((0) - ((self).f29), _dafny.SeqWithoutIsStrInference([d_536_v152_]))
                d_537_v153_ = nw70_
                d_538_v154_: _dafny.Array
                nw71_ = _dafny.Array(None, 16)
                nw71_[int(0)] = d_537_v153_
                nw71_[int(1)] = d_537_v153_
                nw71_[int(2)] = d_537_v153_
                nw71_[int(3)] = d_537_v153_
                nw71_[int(4)] = d_537_v153_
                nw71_[int(5)] = d_537_v153_
                nw71_[int(6)] = d_537_v153_
                nw71_[int(7)] = d_537_v153_
                nw71_[int(8)] = d_537_v153_
                nw71_[int(9)] = d_537_v153_
                nw71_[int(10)] = d_537_v153_
                nw71_[int(11)] = d_537_v153_
                nw71_[int(12)] = d_537_v153_
                nw71_[int(13)] = d_537_v153_
                nw71_[int(14)] = d_537_v153_
                nw71_[int(15)] = d_537_v153_
                d_538_v154_ = nw71_
                index64_ = default__.safeIndex(441, (d_538_v154_).length(0))
                (d_538_v154_)[index64_] = d_537_v153_
                index65_ = default__.safeIndex(441, (d_538_v154_).length(0))
                (d_538_v154_)[index65_] = d_537_v153_
        elif True:
            d_539_v155_: _dafny.Array
            nw72_ = _dafny.Array(int(0), 13)
            d_539_v155_ = nw72_
            d_540_v156_: _dafny.Seq
            d_540_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rmsqxsyv"))
            d_541_v157_: _dafny.Seq
            d_541_v157_ = _dafny.SeqWithoutIsStrInference([d_540_v156_, d_540_v156_])
            index66_ = default__.safeIndex(779, (d_539_v155_).length(0))
            (d_539_v155_)[index66_] = (0) - (len((d_541_v157_)[default__.safeIndex((self).f29, len(d_541_v157_))]))
            d_542_v158_: _dafny.Map
            d_542_v158_ = _dafny.Map({(not(p0) if p0 else p0): (self).f29})
            d_543_v159_: _dafny.Map
            d_543_v159_ = _dafny.Map({(self).f29: (self).f29})
            d_544_v160_: _dafny.Set
            d_544_v160_ = _dafny.Set({(self).f29})
            d_545_v161_: _dafny.Seq
            d_545_v161_ = _dafny.SeqWithoutIsStrInference([len(d_544_v160_)])
            index67_ = default__.safeIndex(779, (d_539_v155_).length(0))
            (d_539_v155_)[index67_] = ((d_542_v158_)[p0] if (p0) in (d_542_v158_) else (0) - (default__.safeDivisionInt(((d_543_v159_)[(self).f29] if ((self).f29) in (d_543_v159_) else len(d_545_v161_)), (self).f29)))
            d_546_v162_: _dafny.Map
            d_546_v162_ = _dafny.Map({_dafny.Map({p0: (d_545_v161_).set(default__.safeIndex((d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))], len(d_545_v161_)), (d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))])}): d_545_v161_})
            d_547_v163_: _dafny.MultiSet
            d_547_v163_ = _dafny.MultiSet([False])
            d_548_v164_: _dafny.Seq
            d_548_v164_ = _dafny.SeqWithoutIsStrInference([default__.fm24(d_547_v163_, (d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))]])), p0, globalState), d_545_v161_, _dafny.SeqWithoutIsStrInference([(self).f29]), d_545_v161_, d_545_v161_])
            d_549_v165_: _dafny.Map
            d_549_v165_ = _dafny.Map({p0: (d_548_v164_)[default__.safeIndex((d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))], len(d_548_v164_))]})
            d_550_v166_: _dafny.Map
            d_550_v166_ = _dafny.Map({(self).f29: d_549_v165_})
            d_545_v161_ = ((d_546_v162_)[(((d_550_v166_)[(self).f29] if ((self).f29) in (d_550_v166_) else d_549_v165_)) | (d_549_v165_)] if ((((d_550_v166_)[(self).f29] if ((self).f29) in (d_550_v166_) else d_549_v165_)) | (d_549_v165_)) in (d_546_v162_) else (d_545_v161_) + (_dafny.SeqWithoutIsStrInference([(self).f29 for d_551_i13_ in range(default__.abs(244))])))
            d_552_v167_: _dafny.Array
            def lambda15_(d_553_v159_):
                def lambda16_(d_554_i14_):
                    return D3_DC10(d_553_v159_)

                return lambda16_

            init8_ = lambda15_(d_543_v159_)
            nw73_ = _dafny.Array(None, 7)
            for i0_8_ in range(nw73_.length(0)):
                nw73_[i0_8_] = init8_(i0_8_)
            d_552_v167_ = nw73_
            d_555_v168_: D3
            d_555_v168_ = D3_DC10(d_543_v159_)
            index68_ = default__.safeIndex(730, (d_552_v167_).length(0))
            (d_552_v167_)[index68_] = d_555_v168_
            d_556_v170_: _dafny.MultiSet
            d_556_v170_ = _dafny.MultiSet([len(d_540_v156_)])
            index69_ = default__.safeIndex(730, (d_552_v167_).length(0))
            def iife53_():
                coll33_ = _dafny.Map()
                compr_33_: int
                for compr_33_ in (d_556_v170_).Elements:
                    d_557_v169_: int = compr_33_
                    if (d_557_v169_) in (d_556_v170_):
                        coll33_[default__.safeDivisionInt(d_557_v169_, (d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))])] = (self).f29
                return _dafny.Map(coll33_)
            (d_552_v167_)[index69_] = D3_DC10((d_543_v159_) | (iife53_()
))
            d_558_v171_: str
            d_558_v171_ = _dafny.CodePoint('s')
            d_559_v172_: _dafny.Map
            d_559_v172_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_558_v171_ for d_560_i15_ in range(default__.abs(-994))])})
            d_561_v173_: _dafny.Array
            nw74_ = _dafny.Array(None, 6)
            nw74_[int(0)] = (d_540_v156_).set(default__.safeIndex(len(d_540_v156_), len(d_540_v156_)), d_558_v171_)
            nw74_[int(1)] = d_540_v156_
            nw74_[int(2)] = ((d_559_v172_)[p0] if (p0) in (d_559_v172_) else d_540_v156_)
            nw74_[int(3)] = d_540_v156_
            nw74_[int(4)] = d_540_v156_
            nw74_[int(5)] = (d_540_v156_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwtvlaj")))
            d_561_v173_ = nw74_
            index70_ = default__.safeIndex(639, (d_561_v173_).length(0))
            (d_561_v173_)[index70_] = d_540_v156_
            d_562_v174_: _dafny.Map
            d_562_v174_ = _dafny.Map({p0: d_543_v159_})
            d_563_v175_: _dafny.Map
            d_563_v175_ = _dafny.Map({(self).f29: p0})
            d_564_v176_: D0
            d_564_v176_ = D0_DC1((d_539_v155_)[default__.safeIndex(779, (d_539_v155_).length(0))], p0, d_562_v174_, len(d_563_v175_), d_558_v171_)
            d_565_v177_: D4
            d_565_v177_ = D4_DC13(d_564_v176_)
            d_566_v178_: D4
            d_566_v178_ = D4_DC16(d_565_v177_)
            d_567_v179_: _dafny.Map
            d_567_v179_ = _dafny.Map({True: p0})
            d_568_v180_: _dafny.Map
            d_568_v180_ = _dafny.Map({d_566_v178_: ((d_567_v179_)[False] if (False) in (d_567_v179_) else p0)})
            d_569_v181_: D10
            d_569_v181_ = D10_DC33(d_568_v180_)
            index71_ = default__.safeIndex(639, (d_561_v173_).length(0))
            (d_561_v173_)[index71_] = (default__.fm30(globalState)).set(default__.safeIndex((0) - (len((d_569_v181_).cf54)), len(default__.fm30(globalState))), d_558_v171_)
            r3 = ((len(_dafny.Set({d_539_v155_}))) - ((self).f29)) + ((0) - (((self).f29) - ((self).f29)))
        d_570_v182_: _dafny.Map
        d_570_v182_ = _dafny.Map({(self).f29: p0})
        d_571_v183_: _dafny.Seq
        d_571_v183_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kdkvyfx"))
        d_572_v184_: _dafny.Map
        d_572_v184_ = _dafny.Map({default__.fm37(_dafny.SeqWithoutIsStrInference([p0]), d_570_v182_, d_571_v183_, globalState): len(d_571_v183_)})
        d_573_v185_: _dafny.Seq
        d_573_v185_ = _dafny.SeqWithoutIsStrInference([d_572_v184_, d_572_v184_])
        d_574_v186_: D9
        d_574_v186_ = D9_DC30((d_573_v185_)[default__.safeIndex((self).f29, len(d_573_v185_))])
        d_575_v187_: D9
        d_575_v187_ = D9_DC32(d_574_v186_)
        source9_ = d_575_v187_
        if source9_.is_DC31:
            d_576___mcc_h11_ = source9_.cf51
            d_577___mcc_h12_ = source9_.cf52
            d_578_cf52_ = d_577___mcc_h12_
            d_579_cf51_ = d_576___mcc_h11_
            d_580_v188_: _dafny.MultiSet
            d_580_v188_ = _dafny.MultiSet([len(d_501_v121_)])
            d_581_v189_: _dafny.Seq
            d_581_v189_ = _dafny.SeqWithoutIsStrInference([d_579_cf51_])
            d_582_v190_: _dafny.Set
            d_582_v190_ = _dafny.Set({len(d_581_v189_)})
            d_583_v191_: _dafny.Map
            d_583_v191_ = _dafny.Map({p0: d_582_v190_})
            d_584_v193_: _dafny.Seq
            def iife54_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in (d_582_v190_).Elements:
                    d_585_v192_: int = compr_34_
                    if (d_585_v192_) in (d_582_v190_):
                        coll34_ = coll34_.union(_dafny.Set([(d_585_v192_) + (492)]))
                return _dafny.Set(coll34_)
            d_584_v193_ = _dafny.SeqWithoutIsStrInference([d_582_v190_, d_582_v190_, d_582_v190_, d_582_v190_, ((d_583_v191_)[p0] if (p0) in (d_583_v191_) else iife54_()
            )])
            d_586_v194_: _dafny.Seq
            d_586_v194_ = _dafny.SeqWithoutIsStrInference([d_579_cf51_, (self).fm0(d_580_v188_, (self).f29, d_584_v193_, p0, globalState)])
            d_587_v195_: D2
            d_587_v195_ = D2_DC6(len(d_586_v194_))
            d_588_v196_: D2
            d_588_v196_ = D2_DC9(d_587_v195_)
            d_589_v197_: _dafny.Set
            d_589_v197_ = _dafny.Set({d_588_v196_, d_588_v196_})
            d_589_v197_ = d_589_v197_
            d_590_v198_: _dafny.Map
            d_590_v198_ = _dafny.Map({d_501_v121_: _dafny.CodePoint('u')})
            d_591_v199_: str
            d_591_v199_ = _dafny.CodePoint('t')
            d_590_v198_ = (d_590_v198_).set(d_501_v121_, d_591_v199_)
            d_592_v200_: _dafny.Map
            d_592_v200_ = _dafny.Map({d_501_v121_: p0})
            d_593_v201_: D8
            d_593_v201_ = D8_DC29(p0, ((d_592_v200_)[d_501_v121_] if (d_501_v121_) in (d_592_v200_) else p0), (d_571_v183_)[default__.safeIndex((self).f29, len(d_571_v183_))])
            d_593_v201_ = d_593_v201_
            d_586_v194_ = d_586_v194_
        elif source9_.is_DC30:
            d_594___mcc_h13_ = source9_.cf50
            d_595_cf50_ = d_594___mcc_h13_
            r0 = p0
            d_596_v202_: int
            d_597_v203_: bool
            d_598_v204_: bool
            d_599_v205_: int
            out42_: int
            out43_: bool
            out44_: bool
            out45_: int
            out42_, out43_, out44_, out45_ = (self).m5((self).f29, (self).f29, globalState)
            d_596_v202_ = out42_
            d_597_v203_ = out43_
            d_598_v204_ = out44_
            d_599_v205_ = out45_
            d_600_v206_: _dafny.MultiSet
            d_600_v206_ = _dafny.MultiSet([d_597_v203_])
            d_601_v207_: _dafny.Map
            d_601_v207_ = _dafny.Map({d_597_v203_: d_600_v206_})
            d_602_v208_: _dafny.Map
            d_602_v208_ = _dafny.Map({993: d_601_v207_})
            d_603_v209_: _dafny.Seq
            d_603_v209_ = _dafny.SeqWithoutIsStrInference([d_597_v203_])
            index72_ = default__.safeIndex(80, (d_480_v101_).length(0))
            (d_480_v101_)[index72_] = (((d_602_v208_)[d_596_v202_] if (d_596_v202_) in (d_602_v208_) else d_601_v207_)) != (default__.fm39((self).f29, d_603_v209_, globalState))
            index73_ = default__.safeIndex(80, (d_480_v101_).length(0))
            rhs45_ = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('t') if p0 else _dafny.CodePoint('l')) for d_604_i16_ in range(default__.abs(961))])
            rhs46_ = ((_dafny.MultiSet([737, 722, (self).f29, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b')]))])).set(d_596_v202_, default__.abs(-639))).cardinality
            rhs47_ = d_501_v121_
            rhs48_ = not(d_598_v204_)
            rhs49_ = d_571_v183_
            lhs38_ = globalState
            lhs39_ = d_480_v101_
            lhs40_ = default__.safeIndex(80, (d_480_v101_).length(0))
            d_571_v183_ = rhs45_
            lhs38_.f11 = rhs46_
            d_501_v121_ = rhs47_
            lhs39_[lhs40_] = rhs48_
            d_571_v183_ = rhs49_
            if default__.fm37(d_603_v209_, d_570_v182_, d_571_v183_, globalState):
                index74_ = default__.safeIndex(80, (d_480_v101_).length(0))
                (d_480_v101_)[index74_] = p0
                d_605_v210_: _dafny.Set
                d_605_v210_ = _dafny.Set({-516, (self).f29})
                d_606_v212_: _dafny.Seq
                d_606_v212_ = _dafny.SeqWithoutIsStrInference([len(d_571_v183_)])
                d_607_v213_: _dafny.Array
                nw75_ = _dafny.Array(None, 22)
                nw75_[int(0)] = d_605_v210_
                nw75_[int(1)] = (d_605_v210_) - (_dafny.Set({len(d_571_v183_)}))
                nw75_[int(2)] = d_605_v210_
                nw75_[int(3)] = d_605_v210_
                nw75_[int(4)] = _dafny.Set({(self).f29, d_599_v205_, 60, len(d_571_v183_), d_596_v202_})
                nw75_[int(5)] = d_605_v210_
                nw75_[int(6)] = _dafny.Set({d_596_v202_})
                def iife55_():
                    coll35_ = _dafny.Set()
                    compr_35_: int
                    for compr_35_ in _dafny.IntegerRange(-907, -883):
                        d_608_v211_: int = compr_35_
                        if ((-907) <= (d_608_v211_)) and ((d_608_v211_) < (-883)):
                            coll35_ = coll35_.union(_dafny.Set([(d_608_v211_) * (len(d_501_v121_))]))
                    return _dafny.Set(coll35_)
                nw75_[int(7)] = iife55_()
                
                nw75_[int(8)] = (d_605_v210_) - (d_605_v210_)
                nw75_[int(9)] = (_dafny.Set({d_599_v205_})) | (d_605_v210_)
                nw75_[int(10)] = d_605_v210_
                nw75_[int(11)] = (d_605_v210_) - (d_605_v210_)
                nw75_[int(12)] = d_605_v210_
                nw75_[int(13)] = default__.fm23(d_596_v202_, d_596_v202_, p0, (d_480_v101_)[default__.safeIndex(80, (d_480_v101_).length(0))], globalState)
                nw75_[int(14)] = d_605_v210_
                nw75_[int(15)] = (d_605_v210_) - (d_605_v210_)
                nw75_[int(16)] = d_605_v210_
                nw75_[int(17)] = d_605_v210_
                nw75_[int(18)] = (default__.fm23(len(d_605_v210_), (d_606_v212_)[default__.safeIndex((self).f29, len(d_606_v212_))], True, d_598_v204_, globalState)).intersection(d_605_v210_)
                nw75_[int(19)] = d_605_v210_
                nw75_[int(20)] = d_605_v210_
                nw75_[int(21)] = d_605_v210_
                d_607_v213_ = nw75_
                d_609_v214_: _dafny.Map
                d_609_v214_ = _dafny.Map({d_571_v183_: (self).f29})
                d_610_v215_: _dafny.Map
                d_610_v215_ = _dafny.Map({d_596_v202_: d_609_v214_})
                d_611_v216_: D5
                d_611_v216_ = D5_DC17(((d_610_v215_)[d_599_v205_] if (d_599_v205_) in (d_610_v215_) else _dafny.Map({d_571_v183_: (self).f29})))
                d_612_v217_: D5
                d_612_v217_ = D5_DC20(d_611_v216_)
                d_613_v218_: str
                d_613_v218_ = _dafny.CodePoint('h')
                d_614_v219_: D10
                d_614_v219_ = D10_DC34(d_613_v218_, (d_480_v101_)[default__.safeIndex(80, (d_480_v101_).length(0))], d_598_v204_, p0, False)
                index75_ = default__.safeIndex(80, (d_480_v101_).length(0))
                rhs50_ = d_607_v213_
                rhs51_ = default__.safeModuloInt(d_599_v205_, d_596_v202_)
                rhs52_ = d_612_v217_
                rhs53_ = (d_614_v219_).cf57
                lhs41_ = globalState
                lhs42_ = d_480_v101_
                lhs43_ = default__.safeIndex(80, (d_480_v101_).length(0))
                d_607_v213_ = rhs50_
                lhs41_.f6 = rhs51_
                d_612_v217_ = rhs52_
                lhs42_[lhs43_] = rhs53_
                index76_ = default__.safeIndex(80, (d_480_v101_).length(0))
                (d_480_v101_)[index76_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) != (d_571_v183_)
                d_615_v220_: _dafny.Map
                d_615_v220_ = _dafny.Map({d_598_v204_: (d_480_v101_)[default__.safeIndex(80, (d_480_v101_).length(0))]})
                (globalState).f10 = (d_596_v202_) < (len(d_615_v220_))
                (globalState).f8 = default__.safeModuloInt(default__.safeModuloInt(len(d_606_v212_), d_596_v202_), ((d_595_cf50_)[p0] if (p0) in (d_595_cf50_) else (self).f29))
            elif True:
                d_571_v183_ = d_571_v183_
                d_597_v203_ = d_598_v204_
                (globalState).f0 = not (True) or ((d_480_v101_)[default__.safeIndex(80, (d_480_v101_).length(0))])
                d_616_v221_: C0
                nw76_ = C0()
                nw76_.ctor__()
                d_616_v221_ = nw76_
                d_617_v222_: _dafny.Array
                nw77_ = _dafny.Array(False, 15)
                d_617_v222_ = nw77_
                d_618_v223_: _dafny.Map
                d_618_v223_ = _dafny.Map({d_596_v202_: _dafny.SeqWithoutIsStrInference([d_617_v222_])})
                d_619_v224_: C1
                nw78_ = C1()
                nw78_.ctor__((self).f29, ((self).f22) + (((d_618_v223_)[d_599_v205_] if (d_599_v205_) in (d_618_v223_) else _dafny.SeqWithoutIsStrInference([d_617_v222_, d_617_v222_, d_617_v222_, d_617_v222_]))))
                d_619_v224_ = nw78_
        elif True:
            d_620___mcc_h14_ = source9_.cf53
            d_621_cf53_ = d_620___mcc_h14_
            (globalState).f16 = ((self).f29 if not(p0) else (_dafny.MultiSet([not(p0)])).cardinality)
            (globalState).f11 = (self).f29
            d_622_v225_: _dafny.Map
            d_622_v225_ = _dafny.Map({p0: True})
            d_623_v226_: _dafny.Map
            d_623_v226_ = _dafny.Map({(d_622_v225_) | (d_622_v225_): p0})
            d_623_v226_ = _dafny.Map({(d_622_v225_ if p0 else d_622_v225_): ((self).f29) >= (len(d_571_v183_))})
            d_624_v227_: D9
            d_624_v227_ = D9_DC31((self).f29, d_480_v101_)
            d_625_v228_: _dafny.Map
            d_625_v228_ = _dafny.Map({((d_622_v225_)[p0] if (p0) in (d_622_v225_) else p0): d_624_v227_})
            d_626_v229_: _dafny.MultiSet
            d_626_v229_ = _dafny.MultiSet([len(d_625_v228_)])
            d_627_v230_: _dafny.Map
            d_627_v230_ = _dafny.Map({d_480_v101_: d_626_v229_})
            d_628_v231_: str
            d_628_v231_ = _dafny.CodePoint('q')
            d_629_v232_: _dafny.Seq
            d_629_v232_ = _dafny.SeqWithoutIsStrInference([d_626_v229_])
            d_630_v233_: D11
            d_630_v233_ = D11_DC36(d_629_v232_)
            d_631_v234_: _dafny.Seq
            d_631_v234_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            rhs54_ = len((((d_630_v233_).cf62) + (d_629_v232_)) + (d_629_v232_))
            rhs55_ = (d_627_v230_).set(d_480_v101_, (d_626_v229_).set((self).f29, default__.abs((self).f29)))
            rhs56_ = p0
            rhs57_ = default__.fm22(p0, (d_631_v234_)[default__.safeIndex((self).f29, len(d_631_v234_))], 667, (0) - (default__.safeDivisionInt((self).f29, len(_dafny.SeqWithoutIsStrInference([p0, p0])))), globalState)
            rhs58_ = not (not(p0)) or (p0)
            lhs44_ = globalState
            lhs45_ = globalState
            lhs46_ = globalState
            lhs44_.f6 = rhs54_
            d_627_v230_ = rhs55_
            lhs45_.f0 = rhs56_
            d_628_v231_ = rhs57_
            lhs46_.f10 = rhs58_
        r0 = p0
        r1 = (d_480_v101_ if not(p0) else d_480_v101_)
        d_632_v235_: T0
        nw79_ = C1()
        nw79_.ctor__((self).f29, (self).f22)
        d_632_v235_ = nw79_
        d_633_v236_: _dafny.Map
        d_633_v236_ = _dafny.Map({p0: d_632_v235_})
        d_634_v237_: _dafny.Array
        nw80_ = _dafny.Array(int(0), 22)
        d_634_v237_ = nw80_
        d_635_v238_: _dafny.Map
        d_635_v238_ = _dafny.Map({not(p0): d_634_v237_})
        d_636_v239_: _dafny.Seq
        d_636_v239_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f29), len(d_633_v236_), len(d_635_v238_), (self).f29])
        d_637_v240_: str
        d_637_v240_ = _dafny.CodePoint('y')
        d_638_v241_: _dafny.Map
        d_638_v241_ = _dafny.Map({d_636_v239_: d_637_v240_})
        r2 = d_638_v241_
        d_639_v242_: _dafny.MultiSet
        d_639_v242_ = _dafny.MultiSet([(self).f29, (self).f29, -818])
        d_640_v243_: _dafny.Map
        d_640_v243_ = _dafny.Map({p0: p0})
        d_641_v244_: _dafny.Set
        d_641_v244_ = _dafny.Set({(self).f29, (self).f29})
        d_642_v245_: _dafny.Seq
        d_642_v245_ = _dafny.SeqWithoutIsStrInference([d_641_v244_])
        r3 = default__.safeDivisionInt((self).f29, (d_632_v235_).fm0(d_639_v242_, len(d_640_v243_), d_642_v245_, p0, globalState))
        return r0, r1, r2, r3

    def m5(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_643_v0_: bool
        d_643_v0_ = True
        d_644_v1_: _dafny.MultiSet
        d_644_v1_ = _dafny.MultiSet([d_643_v0_])
        d_645_v2_: _dafny.Seq
        d_645_v2_ = _dafny.SeqWithoutIsStrInference([d_644_v1_])
        r3 = ((d_645_v2_)[default__.safeIndex(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrquxsrl"))), p0), len(d_645_v2_))]).cardinality
        d_646_v3_: _dafny.MultiSet
        d_646_v3_ = _dafny.MultiSet([p1])
        d_647_v4_: _dafny.Set
        d_647_v4_ = _dafny.Set({len(_dafny.Set({804})), p1})
        d_648_v5_: _dafny.Seq
        d_648_v5_ = _dafny.SeqWithoutIsStrInference([d_647_v4_, _dafny.Set({(self).f29}), d_647_v4_])
        d_649_v6_: _dafny.Seq
        d_649_v6_ = _dafny.SeqWithoutIsStrInference([(self).fm0(d_646_v3_, (self).f29, d_648_v5_, d_643_v0_, globalState)])
        d_650_v7_: _dafny.MultiSet
        d_650_v7_ = _dafny.MultiSet([_dafny.MultiSet([p1]), _dafny.MultiSet(d_649_v6_)])
        d_651_v8_: _dafny.Array
        nw81_ = _dafny.Array(False, 6)
        d_651_v8_ = nw81_
        d_652_v9_: D9
        d_652_v9_ = D9_DC31(len(_dafny.Map({(d_650_v7_).cardinality: 927})), d_651_v8_)
        d_653_v10_: D9
        d_653_v10_ = D9_DC32(d_652_v9_)
        d_653_v10_ = d_653_v10_
        d_654_v11_: _dafny.Seq
        d_654_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcwffodmy"))
        rhs59_ = d_654_v11_
        rhs60_ = d_646_v3_
        d_654_v11_ = rhs59_
        d_646_v3_ = rhs60_
        if d_643_v0_:
            d_655_v13_: _dafny.Seq
            d_655_v13_ = _dafny.SeqWithoutIsStrInference([d_643_v0_, True])
            d_656_v14_: _dafny.Seq
            d_656_v14_ = _dafny.SeqWithoutIsStrInference([d_655_v13_, (_dafny.SeqWithoutIsStrInference([d_643_v0_, d_643_v0_, d_643_v0_, d_643_v0_, d_643_v0_])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_643_v0_, d_643_v0_, d_643_v0_, d_643_v0_, d_643_v0_]))), d_643_v0_), d_655_v13_, d_655_v13_])
            d_657_v15_: _dafny.Map
            d_657_v15_ = _dafny.Map({d_646_v3_: d_643_v0_})
            d_658_v16_: _dafny.Array
            nw82_ = _dafny.Array(None, 20)
            def iife56_():
                coll36_ = _dafny.Map()
                compr_36_: int
                for compr_36_ in (d_647_v4_).Elements:
                    d_659_v12_: int = compr_36_
                    if (d_659_v12_) in (d_647_v4_):
                        coll36_[(d_659_v12_) * (p0)] = (d_654_v11_)[default__.safeIndex((self).f29, len(d_654_v11_))]
                return _dafny.Map(coll36_)
            nw82_[int(0)] = len(iife56_()
            )
            nw82_[int(1)] = (275) + (p0)
            nw82_[int(2)] = len((d_656_v14_) + (_dafny.SeqWithoutIsStrInference([d_655_v13_])))
            nw82_[int(3)] = ((d_644_v1_)[d_643_v0_] if (d_643_v0_) in (d_644_v1_) else (d_644_v1_).cardinality)
            nw82_[int(4)] = (p1) + (p1)
            nw82_[int(5)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wudmpyivk")))) + ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_660_i0_ in range(default__.abs(269))]))))
            nw82_[int(6)] = p1
            nw82_[int(7)] = p1
            nw82_[int(8)] = p0
            nw82_[int(9)] = (p1 if d_643_v0_ else p0)
            nw82_[int(10)] = (self).f29
            nw82_[int(11)] = (d_646_v3_).cardinality
            nw82_[int(12)] = p0
            nw82_[int(13)] = p1
            nw82_[int(14)] = 983
            nw82_[int(15)] = p1
            nw82_[int(16)] = p1
            nw82_[int(17)] = len(d_657_v15_)
            nw82_[int(18)] = (self).f29
            nw82_[int(19)] = (100) - (p1)
            d_658_v16_ = nw82_
            index77_ = default__.safeIndex(898, (d_658_v16_).length(0))
            (d_658_v16_)[index77_] = (self).f29
            index78_ = default__.safeIndex(898, (d_658_v16_).length(0))
            (d_658_v16_)[index78_] = len(_dafny.SeqWithoutIsStrInference([((d_654_v11_)[default__.safeIndex(341, len(d_654_v11_))] if not(d_643_v0_) else _dafny.CodePoint('h')) for d_661_i1_ in range(default__.abs(350))]))
            d_662_v17_: T0
            nw83_ = C1()
            nw83_.ctor__((d_658_v16_)[default__.safeIndex(898, (d_658_v16_).length(0))], (self).f22)
            d_662_v17_ = nw83_
            d_663_v18_: _dafny.MultiSet
            d_663_v18_ = _dafny.MultiSet([d_662_v17_])
            d_663_v18_ = d_663_v18_
            d_664_v19_: _dafny.MultiSet
            d_664_v19_ = _dafny.MultiSet([d_651_v8_])
            d_664_v19_ = ((d_664_v19_) | (d_664_v19_)).intersection(_dafny.MultiSet([d_651_v8_]))
            (globalState).f6 = len(((d_654_v11_ if d_643_v0_ else d_654_v11_)) + (d_654_v11_))
            d_654_v11_ = d_654_v11_
        elif True:
            d_665_v20_: _dafny.Seq
            d_665_v20_ = _dafny.SeqWithoutIsStrInference([d_643_v0_, not(d_643_v0_), d_643_v0_])
            d_666_v21_: _dafny.Map
            d_666_v21_ = _dafny.Map({p0: default__.safeModuloInt(len(d_647_v4_), len(d_665_v20_))})
            index79_ = default__.safeIndex(591, (d_651_v8_).length(0))
            (d_651_v8_)[index79_] = d_643_v0_
            d_667_v23_: _dafny.Map
            d_667_v23_ = _dafny.Map({(d_646_v3_).cardinality: d_643_v0_})
            index80_ = default__.safeIndex(591, (d_651_v8_).length(0))
            def iife57_():
                coll37_ = _dafny.Map()
                compr_37_: int
                for compr_37_ in _dafny.IntegerRange(-244, -210):
                    d_668_v22_: int = compr_37_
                    if ((-244) <= (d_668_v22_)) and ((d_668_v22_) < (-210)):
                        coll37_[default__.safeModuloInt(d_668_v22_, (self).f29)] = len(d_666_v21_)
                return _dafny.Map(coll37_)
            rhs61_ = 499
            rhs62_ = iife57_()

            rhs63_ = ((d_665_v20_)[default__.safeIndex(825, len(d_665_v20_))] if d_643_v0_ else (len(d_647_v4_)) > (len(d_665_v20_)))
            rhs64_ = ((d_667_v23_)[(self).f29] if ((self).f29) in (d_667_v23_) else not(d_643_v0_))
            lhs47_ = globalState
            lhs48_ = d_651_v8_
            lhs49_ = default__.safeIndex(591, (d_651_v8_).length(0))
            lhs47_.f6 = rhs61_
            d_666_v21_ = rhs62_
            r1 = rhs63_
            lhs48_[lhs49_] = rhs64_
            (globalState).f16 = 614
            if ((self).f29) in ((default__.fm40((self).f29, p0, (d_654_v11_)[default__.safeIndex(p0, len(d_654_v11_))], globalState) if not(False) else d_667_v23_)):
                (globalState).f0 = d_643_v0_
                d_669_v24_: _dafny.MultiSet
                d_669_v24_ = _dafny.MultiSet([d_649_v6_, _dafny.SeqWithoutIsStrInference([(self).f29])])
                d_670_v25_: _dafny.Set
                d_670_v25_ = _dafny.Set({(d_651_v8_)[default__.safeIndex(591, (d_651_v8_).length(0))], d_643_v0_})
                rhs65_ = d_669_v24_
                rhs66_ = d_644_v1_
                rhs67_ = default__.fm41(p0, d_654_v11_, d_670_v25_, globalState)
                d_669_v24_ = rhs65_
                d_644_v1_ = rhs66_
                d_644_v1_ = rhs67_
                d_671_v26_: _dafny.Array
                nw84_ = _dafny.Array(None, 4)
                nw84_[int(0)] = p1
                nw84_[int(1)] = p0
                nw84_[int(2)] = p1
                nw84_[int(3)] = p0
                d_671_v26_ = nw84_
                def lambda17_(d_672_p1_):
                    def lambda18_(d_673_i2_):
                        return (d_673_i2_) - (d_672_p1_)

                    return lambda18_

                init9_ = lambda17_(p1)
                nw85_ = _dafny.Array(None, 16)
                for i0_9_ in range(nw85_.length(0)):
                    nw85_[i0_9_] = init9_(i0_9_)
                d_671_v26_ = nw85_
                d_674_v27_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Seq({}), 23)
                d_674_v27_ = nw86_
                rhs68_ = d_674_v27_
                rhs69_ = (d_651_v8_)[default__.safeIndex(591, (d_651_v8_).length(0))]
                lhs50_ = globalState
                d_674_v27_ = rhs68_
                lhs50_.f10 = rhs69_
                index81_ = default__.safeIndex(591, (d_651_v8_).length(0))
                (d_651_v8_)[index81_] = d_643_v0_
            elif True:
                d_675_v28_: _dafny.Array
                def lambda19_(d_676_v0_):
                    def lambda20_(d_677_i3_):
                        return d_676_v0_

                    return lambda20_

                init10_ = lambda19_(d_643_v0_)
                nw87_ = _dafny.Array(None, 23)
                for i0_10_ in range(nw87_.length(0)):
                    nw87_[i0_10_] = init10_(i0_10_)
                d_675_v28_ = nw87_
                d_678_v29_: _dafny.Map
                d_678_v29_ = _dafny.Map({d_675_v28_: p1})
                d_679_v30_: _dafny.Map
                d_679_v30_ = _dafny.Map({(d_649_v6_)[default__.safeIndex(len(d_665_v20_), len(d_649_v6_))]: d_678_v29_})
                d_678_v29_ = ((d_679_v30_)[(p1) * (613)] if ((p1) * (613)) in (d_679_v30_) else _dafny.Map({d_675_v28_: p0}))
                (globalState).f16 = ((self).f29) + ((self).f29)
                d_680_v31_: C1
                nw88_ = C1()
                nw88_.ctor__(default__.safeModuloInt(p0, (self).f29), _dafny.SeqWithoutIsStrInference([d_675_v28_, d_675_v28_, d_675_v28_]))
                d_680_v31_ = nw88_
                d_681_v32_: _dafny.Array
                nw89_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                d_681_v32_ = nw89_
                d_681_v32_ = d_681_v32_
                d_682_v33_: _dafny.Array
                def lambda21_(d_683_p1_):
                    def lambda22_(d_684_i4_):
                        return (d_684_i4_) - (d_683_p1_)

                    return lambda22_

                init11_ = lambda21_(p1)
                nw90_ = _dafny.Array(None, 3)
                for i0_11_ in range(nw90_.length(0)):
                    nw90_[i0_11_] = init11_(i0_11_)
                d_682_v33_ = nw90_
                index82_ = default__.safeIndex(666, (d_682_v33_).length(0))
                (d_682_v33_)[index82_] = (p0) - (p0)
                index83_ = default__.safeIndex(666, (d_682_v33_).length(0))
                (d_682_v33_)[index83_] = 917
            d_685_v34_: C0
            nw91_ = C0()
            nw91_.ctor__()
            d_685_v34_ = nw91_
            d_686_v35_: _dafny.MultiSet
            d_686_v35_ = _dafny.MultiSet([d_685_v34_])
            d_687_v36_: D12
            d_687_v36_ = D12_DC41(d_686_v35_)
            d_688_v37_: C1
            nw92_ = C1()
            nw92_.ctor__(((self).fm0(d_646_v3_, ((d_687_v36_).cf68).cardinality, d_648_v5_, False, globalState) if (d_651_v8_)[default__.safeIndex(591, (d_651_v8_).length(0))] else p1), (self).f22)
            d_688_v37_ = nw92_
            d_643_v0_ = ((d_667_v23_)[(p0) - ((self).f29)] if ((p0) - ((self).f29)) in (d_667_v23_) else (d_651_v8_)[default__.safeIndex(591, (d_651_v8_).length(0))])
        d_689_i5_: int
        d_689_i5_ = 0
        with _dafny.label("3"):
            while not(default__.fm37(_dafny.SeqWithoutIsStrInference([d_643_v0_]), _dafny.Map({p1: not(d_643_v0_)}), d_654_v11_, globalState)):
                with _dafny.c_label("3"):
                    if (d_689_i5_) >= (100):
                        raise _dafny.Break("3")
                    d_689_i5_ = (d_689_i5_) + (1)
                    d_690_v38_: D12
                    d_690_v38_ = D12_DC42((self).f29)
                    d_691_v39_: _dafny.Seq
                    d_691_v39_ = _dafny.SeqWithoutIsStrInference([not(True), d_643_v0_, d_643_v0_])
                    d_692_v40_: _dafny.Set
                    d_692_v40_ = _dafny.Set({d_690_v38_, d_690_v38_, (d_690_v38_ if (d_691_v39_)[default__.safeIndex(len(d_654_v11_), len(d_691_v39_))] else D12_DC42(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agnqcr")))))})
                    d_692_v40_ = d_692_v40_
                    (globalState).f0 = not ((d_691_v39_)[default__.safeIndex(p1, len(d_691_v39_))]) or (not((d_643_v0_ if d_643_v0_ else d_643_v0_)))
                    d_693_v41_: _dafny.Map
                    d_693_v41_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_694_i6_ in range(default__.abs(732))]): p0})
                    d_695_v42_: D5
                    d_695_v42_ = D5_DC17((d_693_v41_) | (d_693_v41_))
                    index84_ = default__.safeIndex(313, (d_651_v8_).length(0))
                    (d_651_v8_)[index84_] = d_643_v0_
                    d_696_v43_: _dafny.Map
                    d_696_v43_ = _dafny.Map({p1: d_643_v0_})
                    d_697_v44_: _dafny.Map
                    d_697_v44_ = _dafny.Map({len(d_696_v43_): d_643_v0_})
                    index85_ = default__.safeIndex(313, (d_651_v8_).length(0))
                    rhs70_ = default__.fm42(globalState)
                    rhs71_ = not (((0) - (p1)) < (p1)) or (default__.fm37(_dafny.SeqWithoutIsStrInference([d_643_v0_]), d_697_v44_, d_654_v11_, globalState))
                    rhs72_ = (d_649_v6_) + (_dafny.SeqWithoutIsStrInference([696]))
                    rhs73_ = not(d_643_v0_)
                    lhs51_ = globalState
                    lhs52_ = d_651_v8_
                    lhs53_ = default__.safeIndex(313, (d_651_v8_).length(0))
                    d_695_v42_ = rhs70_
                    lhs51_.f10 = rhs71_
                    d_649_v6_ = rhs72_
                    lhs52_[lhs53_] = rhs73_
                    (globalState).f11 = (-412) - (621)
                    pass
            pass
        d_698_v45_: _dafny.Seq
        d_698_v45_ = _dafny.SeqWithoutIsStrInference([d_643_v0_])
        (globalState).f5 = (((d_698_v45_).set(default__.safeIndex((self).f29, len(d_698_v45_)), d_643_v0_)) + (d_698_v45_)) + (d_698_v45_)
        d_699_v46_: str
        d_699_v46_ = _dafny.CodePoint('c')
        d_700_v47_: _dafny.Map
        d_700_v47_ = _dafny.Map({(self).f29: d_699_v46_})
        r0 = ((d_649_v6_)[default__.safeIndex(p1, len(d_649_v6_))]) * ((len(d_700_v47_)) - (p0))
        r1 = d_643_v0_
        r2 = ((d_644_v1_).intersection(d_644_v1_)) in (_dafny.SeqWithoutIsStrInference([d_644_v1_ for d_701_i7_ in range(default__.abs(499))]))
        r3 = ((0) - ((0) - (p0)) if d_643_v0_ else (self).f29)
        return r0, r1, r2, r3

    @property
    def f29(self):
        return self._f29

class C3(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f22):
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return (default__.safeModuloInt(-883, 290)) - (-935)

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_702_v0_: int
        d_702_v0_ = -487
        d_703_v1_: _dafny.Seq
        d_703_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
        d_704_v2_: _dafny.Seq
        d_704_v2_ = _dafny.SeqWithoutIsStrInference([d_702_v0_, len(d_703_v1_)])
        d_705_v3_: _dafny.MultiSet
        d_705_v3_ = _dafny.MultiSet([default__.fm14(-870, d_702_v0_, globalState), d_704_v2_])
        d_706_v4_: _dafny.Seq
        d_706_v4_ = _dafny.SeqWithoutIsStrInference([(d_705_v3_).ispropersubset(d_705_v3_), p0])
        d_707_v5_: _dafny.Set
        d_707_v5_ = _dafny.Set({d_702_v0_, d_702_v0_})
        d_708_v6_: _dafny.Seq
        d_708_v6_ = _dafny.SeqWithoutIsStrInference([d_707_v5_, d_707_v5_, d_707_v5_])
        (globalState).f10 = (d_706_v4_)[default__.safeIndex((self).fm0((_dafny.MultiSet([d_702_v0_])).set(d_702_v0_, default__.abs(d_702_v0_)), d_702_v0_, d_708_v6_, p0, globalState), len(d_706_v4_))]
        d_709_i0_: int
        d_709_i0_ = 0
        with _dafny.label("4"):
            while default__.fm15((d_702_v0_) >= (d_702_v0_), globalState):
                with _dafny.c_label("4"):
                    if (d_709_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_709_i0_ = (d_709_i0_) + (1)
                    r0 = (d_702_v0_) >= ((433) + (-409))
                    (globalState).f10 = not((d_706_v4_)[default__.safeIndex(292, len(d_706_v4_))])
                    d_710_v7_: _dafny.MultiSet
                    d_710_v7_ = _dafny.MultiSet([100])
                    (globalState).f11 = (((d_710_v7_).cardinality) - (d_702_v0_)) + (d_702_v0_)
                    d_711_v8_: _dafny.Map
                    d_711_v8_ = _dafny.Map({p0: False})
                    d_711_v8_ = d_711_v8_
                    pass
            pass
        d_712_v9_: _dafny.Set
        d_712_v9_ = _dafny.Set({p0})
        (globalState).f0 = (d_712_v9_).ispropersubset((d_712_v9_).intersection(d_712_v9_))
        d_713_i1_: int
        d_713_i1_ = 0
        with _dafny.label("5"):
            while not((d_706_v4_)[default__.safeIndex(443, len(d_706_v4_))]):
                with _dafny.c_label("5"):
                    if (d_713_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_713_i1_ = (d_713_i1_) + (1)
                    d_714_v10_: _dafny.Map
                    d_714_v10_ = _dafny.Map({(d_704_v2_)[default__.safeIndex(d_702_v0_, len(d_704_v2_))]: d_706_v4_})
                    (globalState).f11 = len((_dafny.Map({d_702_v0_: d_706_v4_})) | ((_dafny.Map({65: d_706_v4_})) | (d_714_v10_)))
                    d_715_v11_: str
                    d_715_v11_ = _dafny.CodePoint('s')
                    d_716_v12_: D4
                    d_716_v12_ = D4_DC14(d_715_v11_, ((self).f22)[default__.safeIndex(d_702_v0_, len((self).f22))], d_703_v1_)
                    d_717_v13_: D4
                    d_717_v13_ = D4_DC16(d_716_v12_)
                    d_718_v14_: D4
                    d_718_v14_ = D4_DC16(d_716_v12_)
                    d_718_v14_ = d_718_v14_
                    (globalState).f10 = p0
                    if ((p0 if p0 else p0)) and ((d_706_v4_)[default__.safeIndex(120, len(d_706_v4_))]):
                        d_703_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdmeaxjx"))
                        r3 = 284
                        d_719_v15_: _dafny.Array
                        nw93_ = _dafny.Array(D2.default()(), 21)
                        d_719_v15_ = nw93_
                        d_720_v16_: _dafny.Seq
                        d_720_v16_ = _dafny.SeqWithoutIsStrInference([d_719_v15_])
                        d_719_v15_ = (d_720_v16_)[default__.safeIndex((0) - (d_702_v0_), len(d_720_v16_))]
                        d_721_v17_: _dafny.Array
                        nw94_ = _dafny.Array(None, 3)
                        nw94_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swkonbn"))) + (d_703_v1_)
                        nw94_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldxmpoy"))
                        nw94_[int(2)] = d_703_v1_
                        d_721_v17_ = nw94_
                        index86_ = default__.safeIndex(423, (d_721_v17_).length(0))
                        (d_721_v17_)[index86_] = d_703_v1_
                        index87_ = default__.safeIndex(423, (d_721_v17_).length(0))
                        (d_721_v17_)[index87_] = d_703_v1_
                        d_722_v18_: _dafny.Map
                        d_722_v18_ = _dafny.Map({len(d_706_v4_): p0})
                        d_722_v18_ = (d_722_v18_).set(d_702_v0_, not(p0))
                    elif True:
                        r0 = (p0) in ((d_706_v4_) + (d_706_v4_))
                        d_723_v19_: C0
                        nw95_ = C0()
                        nw95_.ctor__()
                        d_723_v19_ = nw95_
                        d_724_v20_: _dafny.Array
                        def lambda23_(d_725_i2_):
                            return default__.safeDivisionInt(d_725_i2_, 144)

                        init12_ = lambda23_
                        nw96_ = _dafny.Array(None, 26)
                        for i0_12_ in range(nw96_.length(0)):
                            nw96_[i0_12_] = init12_(i0_12_)
                        d_724_v20_ = nw96_
                        d_726_v21_: _dafny.Array
                        nw97_ = _dafny.Array(False, 6)
                        d_726_v21_ = nw97_
                        d_727_v22_: _dafny.Map
                        d_727_v22_ = _dafny.Map({p0: p0})
                        index88_ = default__.safeIndex(541, (d_726_v21_).length(0))
                        (d_726_v21_)[index88_] = ((d_727_v22_)[p0] if (p0) in (d_727_v22_) else False)
                        d_728_v23_: _dafny.Array
                        nw98_ = _dafny.Array(_dafny.MultiSet({}), 25)
                        d_728_v23_ = nw98_
                        d_729_v24_: _dafny.MultiSet
                        d_729_v24_ = _dafny.MultiSet([d_703_v1_, d_703_v1_])
                        index89_ = default__.safeIndex(873, (d_728_v23_).length(0))
                        (d_728_v23_)[index89_] = d_729_v24_
                        d_730_v25_: _dafny.Map
                        d_730_v25_ = _dafny.Map({False: d_724_v20_})
                        d_731_v26_: _dafny.Map
                        d_731_v26_ = _dafny.Map({(348) + (d_702_v0_): ((d_730_v25_)[p0] if (p0) in (d_730_v25_) else d_724_v20_)})
                        d_732_v27_: _dafny.Map
                        d_732_v27_ = _dafny.Map({d_702_v0_: d_702_v0_})
                        d_733_v28_: _dafny.Map
                        d_733_v28_ = _dafny.Map({d_732_v27_: (d_707_v5_).issubset(d_707_v5_)})
                        index90_ = default__.safeIndex(541, (d_726_v21_).length(0))
                        index91_ = default__.safeIndex(873, (d_728_v23_).length(0))
                        rhs74_ = ((d_731_v26_)[default__.safeModuloInt(d_702_v0_, len(d_704_v2_))] if (default__.safeModuloInt(d_702_v0_, len(d_704_v2_))) in (d_731_v26_) else d_724_v20_)
                        rhs75_ = ((d_733_v28_)[_dafny.Map({d_702_v0_: d_702_v0_})] if (_dafny.Map({d_702_v0_: d_702_v0_})) in (d_733_v28_) else p0)
                        rhs76_ = d_729_v24_
                        lhs54_ = d_726_v21_
                        lhs55_ = default__.safeIndex(541, (d_726_v21_).length(0))
                        lhs56_ = d_728_v23_
                        lhs57_ = default__.safeIndex(873, (d_728_v23_).length(0))
                        d_724_v20_ = rhs74_
                        lhs54_[lhs55_] = rhs75_
                        lhs56_[lhs57_] = rhs76_
                        d_734_v29_: _dafny.MultiSet
                        d_734_v29_ = _dafny.MultiSet([p0, p0])
                        d_735_v30_: _dafny.Seq
                        d_735_v30_ = _dafny.SeqWithoutIsStrInference([d_734_v29_])
                        d_736_v31_: D4
                        d_736_v31_ = D4_DC14(d_715_v11_, d_726_v21_, ((_dafny.SeqWithoutIsStrInference([d_715_v11_ for d_737_i3_ in range(default__.abs(736))])).set(default__.safeIndex(d_702_v0_, len(_dafny.SeqWithoutIsStrInference([d_715_v11_ for d_738_i3_ in range(default__.abs(736))]))), d_715_v11_)).set(default__.safeIndex(688, len((_dafny.SeqWithoutIsStrInference([d_715_v11_ for d_739_i3_ in range(default__.abs(736))])).set(default__.safeIndex(d_702_v0_, len(_dafny.SeqWithoutIsStrInference([d_715_v11_ for d_740_i3_ in range(default__.abs(736))]))), d_715_v11_))), d_715_v11_))
                        d_741_v32_: _dafny.Map
                        d_741_v32_ = _dafny.Map({d_736_v31_: (0) - (d_702_v0_)})
                        index92_ = default__.safeIndex(541, (d_726_v21_).length(0))
                        (d_726_v21_)[index92_] = ((_dafny.MultiSet([p0, (d_726_v21_)[default__.safeIndex(541, (d_726_v21_).length(0))], (d_726_v21_)[default__.safeIndex(541, (d_726_v21_).length(0))]]) if not(default__.fm15(p0, globalState)) else d_734_v29_)).isdisjoint((d_735_v30_)[default__.safeIndex(len(d_741_v32_), len(d_735_v30_))])
                        d_742_v33_: _dafny.Map
                        d_742_v33_ = _dafny.Map({d_723_v19_: (d_726_v21_)[default__.safeIndex(541, (d_726_v21_).length(0))]})
                        d_742_v33_ = (d_742_v33_) | (d_742_v33_)
                    pass
            pass
        d_743_v34_: D1
        d_743_v34_ = D1_DC3(d_703_v1_)
        d_744_v35_: _dafny.Array
        nw99_ = _dafny.Array(int(0), 10)
        d_744_v35_ = nw99_
        d_745_v36_: D2
        d_745_v36_ = D2_DC8(d_743_v34_, d_744_v35_, p0)
        if ((D2_DC8(d_743_v34_, d_744_v35_, p0) if p0 else d_745_v36_)).cf16:
            index93_ = default__.safeIndex(440, (d_744_v35_).length(0))
            (d_744_v35_)[index93_] = d_702_v0_
            index94_ = default__.safeIndex(440, (d_744_v35_).length(0))
            (d_744_v35_)[index94_] = d_702_v0_
            d_746_v37_: str
            d_746_v37_ = _dafny.CodePoint('q')
            d_747_v38_: _dafny.Map
            d_747_v38_ = _dafny.Map({len(d_703_v1_): False})
            index95_ = default__.safeIndex(440, (d_744_v35_).length(0))
            rhs77_ = d_702_v0_
            rhs78_ = (d_746_v37_ if p0 else d_746_v37_)
            rhs79_ = ((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]) <= (len(d_747_v38_))
            rhs80_ = d_702_v0_
            lhs58_ = d_744_v35_
            lhs59_ = default__.safeIndex(440, (d_744_v35_).length(0))
            lhs60_ = globalState
            lhs61_ = globalState
            lhs58_[lhs59_] = rhs77_
            d_746_v37_ = rhs78_
            lhs60_.f0 = rhs79_
            lhs61_.f11 = rhs80_
            d_748_v39_: D2
            d_748_v39_ = D2_DC6(default__.safeModuloInt(d_702_v0_, d_702_v0_))
            d_748_v39_ = d_748_v39_
            if (d_706_v4_)[default__.safeIndex(len(d_704_v2_), len(d_706_v4_))]:
                rhs81_ = ((0) - (d_702_v0_)) not in (d_707_v5_)
                rhs82_ = default__.fm18(p0, globalState)
                rhs83_ = default__.fm15(((0) - ((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))])) > (d_702_v0_), globalState)
                rhs84_ = p0
                lhs62_ = globalState
                lhs63_ = globalState
                lhs62_.f10 = rhs81_
                d_703_v1_ = rhs82_
                lhs63_.f10 = rhs83_
                r0 = rhs84_
                (globalState).f16 = (0) - ((d_702_v0_) - ((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]))
                d_749_v40_: _dafny.Array
                nw100_ = _dafny.Array(None, 15)
                nw100_[int(0)] = d_746_v37_
                nw100_[int(1)] = _dafny.CodePoint('e')
                nw100_[int(2)] = d_746_v37_
                nw100_[int(3)] = d_746_v37_
                nw100_[int(4)] = d_746_v37_
                nw100_[int(5)] = d_746_v37_
                nw100_[int(6)] = (d_746_v37_ if True else d_746_v37_)
                nw100_[int(7)] = (d_746_v37_ if p0 else d_746_v37_)
                nw100_[int(8)] = d_746_v37_
                nw100_[int(9)] = (d_703_v1_)[default__.safeIndex(d_702_v0_, len(d_703_v1_))]
                nw100_[int(10)] = d_746_v37_
                nw100_[int(11)] = d_746_v37_
                nw100_[int(12)] = d_746_v37_
                nw100_[int(13)] = _dafny.CodePoint('g')
                nw100_[int(14)] = d_746_v37_
                d_749_v40_ = nw100_
                index96_ = default__.safeIndex(455, (d_749_v40_).length(0))
                (d_749_v40_)[index96_] = d_746_v37_
                index97_ = default__.safeIndex(455, (d_749_v40_).length(0))
                (d_749_v40_)[index97_] = d_746_v37_
                d_750_v41_: _dafny.Map
                d_750_v41_ = _dafny.Map({not(p0): p0})
                r3 = (len(d_750_v41_)) * (d_702_v0_)
                d_751_v42_: _dafny.Map
                d_751_v42_ = _dafny.Map({d_703_v1_: (d_706_v4_) <= (d_706_v4_)})
                d_751_v42_ = (d_751_v42_).set(d_703_v1_, (d_706_v4_) <= (d_706_v4_))
            elif True:
                d_746_v37_ = _dafny.CodePoint('l')
                index98_ = default__.safeIndex(440, (d_744_v35_).length(0))
                (d_744_v35_)[index98_] = len(default__.fm18(p0, globalState))
                d_752_v43_: _dafny.Map
                d_752_v43_ = _dafny.Map({d_703_v1_: (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]})
                d_753_v45_: _dafny.Set
                d_753_v45_ = _dafny.Set({d_703_v1_})
                d_754_v46_: D5
                d_754_v46_ = D5_DC17(d_752_v43_)
                d_755_v47_: _dafny.Array
                nw101_ = _dafny.Array(None, 20)
                nw101_[int(0)] = d_752_v43_
                nw101_[int(1)] = d_752_v43_
                nw101_[int(2)] = _dafny.Map({d_703_v1_: d_702_v0_})
                nw101_[int(3)] = d_752_v43_
                nw101_[int(4)] = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucufv"))).set(default__.safeIndex((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucufv")))), d_746_v37_): d_702_v0_})
                nw101_[int(5)] = d_752_v43_
                def iife58_():
                    coll38_ = _dafny.Map()
                    compr_38_: _dafny.Seq
                    for compr_38_ in (d_753_v45_).Elements:
                        d_756_v44_: _dafny.Seq = compr_38_
                        if (d_756_v44_) in (d_753_v45_):
                            coll38_[d_756_v44_] = (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]
                    return _dafny.Map(coll38_)
                nw101_[int(6)] = ((iife58_()
                ).set(d_703_v1_, len(d_704_v2_)) if (d_706_v4_)[default__.safeIndex(d_702_v0_, len(d_706_v4_))] else d_752_v43_)
                nw101_[int(7)] = default__.fm19(p0, d_702_v0_, (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))], True, globalState)
                nw101_[int(8)] = d_752_v43_
                nw101_[int(9)] = d_752_v43_
                nw101_[int(10)] = d_752_v43_
                nw101_[int(11)] = d_752_v43_
                nw101_[int(12)] = (d_752_v43_) | (_dafny.Map({d_703_v1_: len(d_706_v4_)}))
                nw101_[int(13)] = d_752_v43_
                nw101_[int(14)] = d_752_v43_
                nw101_[int(15)] = _dafny.Map({d_703_v1_: d_702_v0_})
                nw101_[int(16)] = d_752_v43_
                nw101_[int(17)] = d_752_v43_
                nw101_[int(18)] = (d_754_v46_).cf29
                nw101_[int(19)] = (d_752_v43_ if p0 else _dafny.Map({d_703_v1_: d_702_v0_}))
                d_755_v47_ = nw101_
                index99_ = default__.safeIndex(885, (d_755_v47_).length(0))
                (d_755_v47_)[index99_] = d_752_v43_
                d_757_v48_: _dafny.MultiSet
                d_757_v48_ = _dafny.MultiSet([p0, False])
                d_758_v50_: _dafny.Map
                def iife59_():
                    coll39_ = _dafny.Map()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(733, -844):
                        d_759_v49_: int = compr_39_
                        if ((733) <= (d_759_v49_)) and ((d_759_v49_) < (-844)):
                            coll39_[(d_759_v49_) * ((d_704_v2_)[default__.safeIndex(d_702_v0_, len(d_704_v2_))])] = 859
                    return _dafny.Map(coll39_)
                d_758_v50_ = _dafny.Map({p0: iife59_()
                })
                d_760_v51_: D0
                d_760_v51_ = D0_DC1((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))], p0, d_758_v50_, d_702_v0_, d_746_v37_)
                index100_ = default__.safeIndex(885, (d_755_v47_).length(0))
                index101_ = default__.safeIndex(440, (d_744_v35_).length(0))
                rhs85_ = default__.fm19(((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]) > (((d_757_v48_)[p0] if (p0) in (d_757_v48_) else (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))])), (d_702_v0_) * (d_702_v0_), 829, (d_760_v51_).cf1, globalState)
                rhs86_ = default__.safeDivisionInt((d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))], default__.safeDivisionInt(557, (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]))
                rhs87_ = -230
                lhs64_ = d_755_v47_
                lhs65_ = default__.safeIndex(885, (d_755_v47_).length(0))
                lhs66_ = d_744_v35_
                lhs67_ = default__.safeIndex(440, (d_744_v35_).length(0))
                lhs64_[lhs65_] = rhs85_
                lhs66_[lhs67_] = rhs86_
                r3 = rhs87_
                d_761_v52_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.Map({}), 8)
                d_761_v52_ = nw102_
                d_762_v53_: _dafny.Array
                nw103_ = _dafny.Array(None, 9)
                nw103_[int(0)] = (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]
                nw103_[int(1)] = (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]
                nw103_[int(2)] = d_702_v0_
                nw103_[int(3)] = len(default__.fm20(p0, d_746_v37_, d_702_v0_, p0, globalState))
                nw103_[int(4)] = d_702_v0_
                nw103_[int(5)] = d_702_v0_
                nw103_[int(6)] = d_702_v0_
                nw103_[int(7)] = (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]
                nw103_[int(8)] = (d_744_v35_)[default__.safeIndex(440, (d_744_v35_).length(0))]
                d_762_v53_ = nw103_
                d_763_v54_: _dafny.Map
                d_763_v54_ = _dafny.Map({d_762_v53_: p0})
                index102_ = default__.safeIndex(183, (d_761_v52_).length(0))
                (d_761_v52_)[index102_] = d_763_v54_
                d_764_v55_: C0
                nw104_ = C0()
                nw104_.ctor__()
                d_764_v55_ = nw104_
                index103_ = default__.safeIndex(183, (d_761_v52_).length(0))
                rhs88_ = d_763_v54_
                rhs89_ = d_764_v55_
                lhs68_ = d_761_v52_
                lhs69_ = default__.safeIndex(183, (d_761_v52_).length(0))
                lhs68_[lhs69_] = rhs88_
                d_764_v55_ = rhs89_
                d_765_v56_: _dafny.Map
                d_765_v56_ = _dafny.Map({p0: p0})
                (globalState).f0 = (d_706_v4_)[default__.safeIndex((0) - (default__.safeDivisionInt(len(d_765_v56_), (0) - (d_702_v0_))), len(d_706_v4_))]
            index104_ = default__.safeIndex(440, (d_744_v35_).length(0))
            rhs90_ = d_702_v0_
            rhs91_ = p0
            lhs70_ = d_744_v35_
            lhs71_ = default__.safeIndex(440, (d_744_v35_).length(0))
            lhs72_ = globalState
            lhs70_[lhs71_] = rhs90_
            lhs72_.f0 = rhs91_
        elif True:
            (globalState).f10 = default__.fm15(p0, globalState)
            (globalState).f0 = p0
            d_766_v57_: _dafny.Array
            def lambda24_(d_767_p0_):
                def lambda25_(d_768_i4_):
                    return d_767_p0_

                return lambda25_

            init13_ = lambda24_(p0)
            nw105_ = _dafny.Array(None, 26)
            for i0_13_ in range(nw105_.length(0)):
                nw105_[i0_13_] = init13_(i0_13_)
            d_766_v57_ = nw105_
            index105_ = default__.safeIndex(873, (d_766_v57_).length(0))
            (d_766_v57_)[index105_] = False
            d_769_v58_: _dafny.Set
            d_769_v58_ = _dafny.Set({_dafny.Map({d_702_v0_: p0})})
            d_770_v59_: _dafny.Map
            d_770_v59_ = _dafny.Map({205: default__.fm15(False, globalState)})
            index106_ = default__.safeIndex(873, (d_766_v57_).length(0))
            rhs92_ = (_dafny.Set({d_770_v59_})).ispropersubset(d_769_v58_)
            rhs93_ = not ((d_702_v0_) in (d_707_v5_)) or (default__.fm15(p0, globalState))
            lhs73_ = d_766_v57_
            lhs74_ = default__.safeIndex(873, (d_766_v57_).length(0))
            r0 = rhs92_
            lhs73_[lhs74_] = rhs93_
            d_771_v60_: C0
            nw106_ = C0()
            nw106_.ctor__()
            d_771_v60_ = nw106_
            d_772_v61_: D2
            d_772_v61_ = D2_DC9(D2_DC6(d_702_v0_))
            d_773_v62_: D2
            d_773_v62_ = D2_DC9(d_772_v61_)
            d_774_v63_: _dafny.Set
            d_774_v63_ = _dafny.Set({d_773_v62_})
            d_775_v64_: _dafny.MultiSet
            d_775_v64_ = _dafny.MultiSet([d_766_v57_])
            d_776_v65_: _dafny.MultiSet
            d_776_v65_ = _dafny.MultiSet([(d_775_v64_).cardinality, 715])
            d_777_v66_: _dafny.Map
            d_777_v66_ = _dafny.Map({d_702_v0_: d_776_v65_})
            d_778_v67_: D4
            d_778_v67_ = D4_DC12(d_777_v66_)
            d_779_v68_: D4
            d_779_v68_ = D4_DC16(d_778_v67_)
            d_780_v69_: D4
            d_780_v69_ = D4_DC16((d_779_v68_).cf28)
            d_781_v70_: D4
            d_781_v70_ = D4_DC16(d_778_v67_)
            d_782_v71_: _dafny.Seq
            d_782_v71_ = _dafny.SeqWithoutIsStrInference([d_779_v68_, d_779_v68_, d_779_v68_])
            d_783_v72_: _dafny.Map
            d_783_v72_ = _dafny.Map({(_dafny.Set({d_773_v62_})).intersection(d_774_v63_): d_782_v71_})
            d_783_v72_ = (d_783_v72_).set(d_774_v63_, d_782_v71_)
        d_784_v73_: str
        d_784_v73_ = _dafny.CodePoint('o')
        d_785_v74_: _dafny.Array
        nw107_ = _dafny.Array(False, 9)
        d_785_v74_ = nw107_
        d_786_v75_: D4
        d_786_v75_ = D4_DC14(d_784_v73_, d_785_v74_, d_703_v1_)
        source10_ = d_786_v75_
        if source10_.is_DC13:
            d_787___mcc_h0_ = source10_.cf22
            d_788_cf22_ = d_787___mcc_h0_
            if p0:
                (globalState).f0 = (d_706_v4_)[default__.safeIndex(len(d_712_v9_), len(d_706_v4_))]
                d_789_v76_: _dafny.Array
                def lambda26_(d_790_v2_):
                    def lambda27_(d_791_i5_):
                        return d_790_v2_

                    return lambda27_

                init14_ = lambda26_(d_704_v2_)
                nw108_ = _dafny.Array(None, 16)
                for i0_14_ in range(nw108_.length(0)):
                    nw108_[i0_14_] = init14_(i0_14_)
                d_789_v76_ = nw108_
                index107_ = default__.safeIndex(192, (d_789_v76_).length(0))
                (d_789_v76_)[index107_] = d_704_v2_
                index108_ = default__.safeIndex(326, (d_744_v35_).length(0))
                (d_744_v35_)[index108_] = d_702_v0_
                index109_ = default__.safeIndex(192, (d_789_v76_).length(0))
                index110_ = default__.safeIndex(326, (d_744_v35_).length(0))
                rhs94_ = ((_dafny.SeqWithoutIsStrInference([(0) - (d_702_v0_) for d_792_i6_ in range(default__.abs(60))])) + ((d_704_v2_) + (d_704_v2_))).set(default__.safeIndex((0) - (d_702_v0_), len((_dafny.SeqWithoutIsStrInference([(0) - (d_702_v0_) for d_793_i6_ in range(default__.abs(60))])) + ((d_704_v2_) + (d_704_v2_)))), d_702_v0_)
                rhs95_ = (d_702_v0_) + ((d_702_v0_) - (d_702_v0_))
                lhs75_ = d_789_v76_
                lhs76_ = default__.safeIndex(192, (d_789_v76_).length(0))
                lhs77_ = d_744_v35_
                lhs78_ = default__.safeIndex(326, (d_744_v35_).length(0))
                lhs75_[lhs76_] = rhs94_
                lhs77_[lhs78_] = rhs95_
                d_794_v77_: _dafny.Array
                d_795_v78_: int
                out46_: _dafny.Array
                out47_: int
                out46_, out47_ = (self).m4(globalState)
                d_794_v77_ = out46_
                d_795_v78_ = out47_
                (globalState).f5 = ((_dafny.SeqWithoutIsStrInference([False, p0])) + (d_706_v4_)) + (d_706_v4_)
                index111_ = default__.safeIndex(974, (d_785_v74_).length(0))
                (d_785_v74_)[index111_] = not(not(((d_744_v35_)[default__.safeIndex(326, (d_744_v35_).length(0))]) < (d_702_v0_)))
                index112_ = default__.safeIndex(974, (d_785_v74_).length(0))
                (d_785_v74_)[index112_] = default__.fm15(True, globalState)
            elif True:
                d_796_v79_: _dafny.Map
                d_796_v79_ = _dafny.Map({p0: d_785_v74_})
                (globalState).f6 = len(d_796_v79_)
                d_744_v35_ = d_744_v35_
                d_784_v73_ = d_784_v73_
                d_797_v80_: C0
                nw109_ = C0()
                nw109_.ctor__()
                d_797_v80_ = nw109_
                rhs96_ = d_744_v35_
                rhs97_ = d_702_v0_
                lhs79_ = globalState
                d_744_v35_ = rhs96_
                lhs79_.f11 = rhs97_
            index113_ = default__.safeIndex(749, (d_744_v35_).length(0))
            (d_744_v35_)[index113_] = d_702_v0_
            d_798_v81_: _dafny.MultiSet
            d_798_v81_ = _dafny.MultiSet([d_702_v0_])
            d_799_v82_: _dafny.Map
            d_799_v82_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgi")): p0})
            index114_ = default__.safeIndex(749, (d_744_v35_).length(0))
            def iife60_():
                coll40_ = _dafny.Set()
                compr_40_: int
                for compr_40_ in _dafny.IntegerRange(623, 366):
                    d_800_v83_: int = compr_40_
                    if ((623) <= (d_800_v83_)) and ((d_800_v83_) < (366)):
                        coll40_ = coll40_.union(_dafny.Set([default__.safeDivisionInt(d_800_v83_, d_702_v0_)]))
                return _dafny.Set(coll40_)
            (d_744_v35_)[index114_] = (self).fm0(d_798_v81_, (0) - (len(d_799_v82_)), _dafny.SeqWithoutIsStrInference([iife60_()
            ]), p0, globalState)
            d_801_v84_: _dafny.Map
            d_801_v84_ = _dafny.Map({not(p0): p0})
            d_802_v85_: _dafny.Seq
            d_802_v85_ = _dafny.SeqWithoutIsStrInference([d_801_v84_])
            (globalState).f6 = (0) - (len(d_802_v85_))
            if p0:
                (globalState).f10 = (p0 if (d_784_v73_) not in (d_703_v1_) else p0)
                (globalState).f0 = (-252) < (d_702_v0_)
                d_803_v86_: T0
                nw110_ = C1()
                nw110_.ctor__(d_702_v0_, (self).f22)
                d_803_v86_ = nw110_
                d_804_v87_: _dafny.Seq
                d_804_v87_ = _dafny.SeqWithoutIsStrInference([d_803_v86_])
                d_805_v88_: _dafny.Set
                d_805_v88_ = _dafny.Set({d_804_v87_, _dafny.SeqWithoutIsStrInference([d_803_v86_])})
                d_806_v89_: _dafny.Map
                d_806_v89_ = _dafny.Map({(d_706_v4_)[default__.safeIndex((self).fm0(_dafny.MultiSet([(d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))], d_702_v0_, d_702_v0_]), (d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))], d_708_v6_, p0, globalState), len(d_706_v4_))]: d_805_v88_})
                rhs98_ = (len((d_805_v88_) - (((d_806_v89_)[p0] if (p0) in (d_806_v89_) else _dafny.Set({(_dafny.SeqWithoutIsStrInference([d_803_v86_])).set(default__.safeIndex(len(_dafny.Map({(d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))]: (d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))]})), len(_dafny.SeqWithoutIsStrInference([d_803_v86_]))), d_803_v86_)}))))) + ((d_702_v0_) * ((d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))]))
                rhs99_ = default__.safeDivisionInt(d_702_v0_, d_702_v0_)
                lhs80_ = globalState
                lhs80_.f11 = rhs98_
                r3 = rhs99_
                d_807_v90_: _dafny.MultiSet
                d_807_v90_ = _dafny.MultiSet([p0])
                d_808_v91_: _dafny.Map
                d_808_v91_ = _dafny.Map({(d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))]: p0})
                d_809_v92_: _dafny.Map
                d_809_v92_ = _dafny.Map({((d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))] if p0 else ((d_807_v90_)[((d_808_v91_)[d_702_v0_] if (d_702_v0_) in (d_808_v91_) else p0)] if (((d_808_v91_)[d_702_v0_] if (d_702_v0_) in (d_808_v91_) else p0)) in (d_807_v90_) else d_702_v0_)): default__.safeDivisionInt((d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))], (d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))])})
                rhs100_ = d_784_v73_
                rhs101_ = ((d_809_v92_)[len(d_703_v1_)] if (len(d_703_v1_)) in (d_809_v92_) else (_dafny.MultiSet([d_703_v1_])).cardinality)
                rhs102_ = p0
                rhs103_ = p0
                rhs104_ = d_803_v86_
                lhs81_ = globalState
                lhs82_ = globalState
                lhs83_ = globalState
                d_784_v73_ = rhs100_
                lhs81_.f6 = rhs101_
                lhs82_.f0 = rhs102_
                lhs83_.f10 = rhs103_
                d_803_v86_ = rhs104_
                d_810_v93_: _dafny.MultiSet
                d_810_v93_ = _dafny.MultiSet([(d_745_v36_).cf15])
                d_811_v94_: C1
                nw111_ = C1()
                nw111_.ctor__((d_810_v93_).cardinality, (d_803_v86_).f22)
                d_811_v94_ = nw111_
            elif True:
                d_812_v95_: D1
                d_812_v95_ = D1_DC2(d_703_v1_)
                d_813_v96_: _dafny.Set
                d_813_v96_ = _dafny.Set({d_812_v95_, d_812_v95_})
                (globalState).f10 = not((_dafny.Set({d_812_v95_, d_812_v95_})).isdisjoint(d_813_v96_))
                d_814_v98_: _dafny.Map
                d_814_v98_ = _dafny.Map({(d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))]: p0})
                d_815_v100_: _dafny.Map
                def iife61_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in (d_814_v98_).keys.Elements:
                        d_816_v97_: int = compr_41_
                        if (d_816_v97_) in (d_814_v98_):
                            coll41_[default__.safeDivisionInt(d_816_v97_, (d_744_v35_)[default__.safeIndex(749, (d_744_v35_).length(0))])] = False
                    return _dafny.Map(coll41_)
                def iife62_():
                    coll42_ = _dafny.Map()
                    compr_42_: int
                    for compr_42_ in (d_814_v98_).keys.Elements:
                        d_817_v99_: int = compr_42_
                        if (d_817_v99_) in (d_814_v98_):
                            coll42_[(d_817_v99_) - (d_702_v0_)] = p0
                    return _dafny.Map(coll42_)
                d_815_v100_ = _dafny.Map({p0: (iife61_()
                ) | ((iife62_()
                ).set(d_702_v0_, True))})
                index115_ = default__.safeIndex(749, (d_744_v35_).length(0))
                (d_744_v35_)[index115_] = len(((d_815_v100_)[not((True) and (p0))] if (not((True) and (p0))) in (d_815_v100_) else (d_814_v98_) | (d_814_v98_)))
                d_814_v98_ = ((d_814_v98_) | (d_814_v98_)) | (d_814_v98_)
                d_818_v101_: _dafny.Array
                def lambda28_(d_819_v1_, d_820_v95_):
                    def lambda29_(d_821_i7_):
                        return (d_819_v1_) + ((d_820_v95_).cf5)

                    return lambda29_

                init15_ = lambda28_(d_703_v1_, d_812_v95_)
                nw112_ = _dafny.Array(None, 20)
                for i0_15_ in range(nw112_.length(0)):
                    nw112_[i0_15_] = init15_(i0_15_)
                d_818_v101_ = nw112_
                index116_ = default__.safeIndex(386, (d_818_v101_).length(0))
                (d_818_v101_)[index116_] = d_703_v1_
                d_822_v102_: D12
                d_822_v102_ = D12_DC42(d_702_v0_)
                d_823_v103_: _dafny.Map
                d_823_v103_ = _dafny.Map({p0: _dafny.CodePoint('f')})
                index117_ = default__.safeIndex(386, (d_818_v101_).length(0))
                rhs105_ = p0
                rhs106_ = (0) - ((d_822_v102_).cf69)
                rhs107_ = (d_703_v1_) + (d_703_v1_)
                rhs108_ = (d_703_v1_).set(default__.safeIndex((0) - (d_702_v0_), len(d_703_v1_)), ((d_823_v103_)[p0] if (p0) in (d_823_v103_) else d_784_v73_))
                rhs109_ = p0
                lhs84_ = globalState
                lhs85_ = globalState
                lhs86_ = d_818_v101_
                lhs87_ = default__.safeIndex(386, (d_818_v101_).length(0))
                lhs88_ = globalState
                lhs84_.f0 = rhs105_
                lhs85_.f11 = rhs106_
                d_703_v1_ = rhs107_
                lhs86_[lhs87_] = rhs108_
                lhs88_.f10 = rhs109_
                d_824_v104_: _dafny.Array
                nw113_ = _dafny.Array(int(0), 7)
                d_824_v104_ = nw113_
                d_825_v105_: _dafny.Array
                nw114_ = _dafny.Array(None, 21)
                nw114_[int(0)] = d_824_v104_
                nw114_[int(1)] = d_824_v104_
                nw114_[int(2)] = d_824_v104_
                nw114_[int(3)] = d_824_v104_
                nw114_[int(4)] = d_824_v104_
                nw114_[int(5)] = d_824_v104_
                nw114_[int(6)] = d_824_v104_
                nw114_[int(7)] = d_824_v104_
                nw114_[int(8)] = d_824_v104_
                nw114_[int(9)] = d_824_v104_
                nw114_[int(10)] = d_824_v104_
                nw114_[int(11)] = d_824_v104_
                nw114_[int(12)] = d_824_v104_
                nw114_[int(13)] = d_824_v104_
                nw114_[int(14)] = d_824_v104_
                nw114_[int(15)] = d_824_v104_
                nw114_[int(16)] = d_824_v104_
                nw114_[int(17)] = d_824_v104_
                nw114_[int(18)] = d_824_v104_
                nw114_[int(19)] = d_824_v104_
                nw114_[int(20)] = d_824_v104_
                d_825_v105_ = nw114_
                d_825_v105_ = d_825_v105_
        elif source10_.is_DC14:
            d_826___mcc_h1_ = source10_.cf23
            d_827___mcc_h2_ = source10_.cf24
            d_828___mcc_h3_ = source10_.cf25
            d_829_cf25_ = d_828___mcc_h3_
            d_830_cf24_ = d_827___mcc_h2_
            d_831_cf23_ = d_826___mcc_h1_
            d_832_v106_: _dafny.Array
            d_833_v107_: int
            out48_: _dafny.Array
            out49_: int
            out48_, out49_ = (self).m4(globalState)
            d_832_v106_ = out48_
            d_833_v107_ = out49_
            d_834_v109_: _dafny.Map
            def iife63_():
                coll43_ = _dafny.Map()
                compr_43_: str
                for compr_43_ in (d_703_v1_).Elements:
                    d_835_v108_: str = compr_43_
                    if (d_835_v108_) in (d_703_v1_):
                        coll43_[d_835_v108_] = len(d_707_v5_)
                return _dafny.Map(coll43_)
            d_834_v109_ = _dafny.Map({iife63_()
            : d_833_v107_})
            d_836_v110_: _dafny.Map
            d_836_v110_ = _dafny.Map({d_831_cf23_: len(d_707_v5_)})
            (globalState).f11 = (0) - (len(((d_834_v109_) | ((d_834_v109_).set(d_836_v110_, d_702_v0_))) | (_dafny.Map({d_836_v110_: d_702_v0_}))))
            if not((d_702_v0_) <= (default__.safeDivisionInt(d_702_v0_, d_833_v107_))):
                (globalState).f0 = p0
                d_703_v1_ = ((d_829_cf25_) + (d_703_v1_)) + (d_829_cf25_)
                d_837_v111_: _dafny.Map
                d_837_v111_ = _dafny.Map({960: p0})
                d_838_v112_: _dafny.Map
                d_838_v112_ = _dafny.Map({p0: d_702_v0_})
                d_839_v113_: D9
                d_839_v113_ = D9_DC30(d_838_v112_)
                d_840_v114_: D9
                d_840_v114_ = D9_DC32(d_839_v113_)
                d_841_v115_: D9
                d_841_v115_ = D9_DC32((d_839_v113_ if ((d_837_v111_)[d_702_v0_] if (d_702_v0_) in (d_837_v111_) else p0) else d_840_v114_))
                pat_let_tv12_ = d_839_v113_
                def iife64_(_pat_let10_0):
                    def iife65_(d_842_dt__update__tmp_h0_):
                        def iife66_(_pat_let11_0):
                            def iife67_(d_843_dt__update_hcf53_h0_):
                                return D9_DC32(d_843_dt__update_hcf53_h0_)
                            return iife67_(_pat_let11_0)
                        return iife66_(pat_let_tv12_)
                    return iife65_(_pat_let10_0)
                d_841_v115_ = iife64_(d_841_v115_)
                r0 = False
                d_844_v116_: _dafny.Seq
                d_844_v116_ = _dafny.SeqWithoutIsStrInference([default__.fm29(p0, len(_dafny.SeqWithoutIsStrInference([d_784_v73_ for d_845_i8_ in range(default__.abs(220))])), p0, d_702_v0_, globalState)])
                d_846_v117_: D11
                d_846_v117_ = D11_DC36(d_844_v116_)
                index118_ = default__.safeIndex(250, (d_785_v74_).length(0))
                (d_785_v74_)[index118_] = p0
                d_847_v118_: _dafny.MultiSet
                d_847_v118_ = _dafny.MultiSet([d_833_v107_, d_833_v107_, d_833_v107_])
                d_848_v119_: _dafny.Map
                d_848_v119_ = _dafny.Map({False: d_847_v118_})
                d_849_v120_: _dafny.Map
                d_849_v120_ = _dafny.Map({p0: d_829_cf25_})
                d_850_v121_: _dafny.Map
                d_850_v121_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_851_i9_ in range(default__.abs(-940))])): len(d_849_v120_)})
                d_852_v122_: _dafny.Map
                d_852_v122_ = _dafny.Map({p0: d_850_v121_})
                index119_ = default__.safeIndex(250, (d_785_v74_).length(0))
                rhs110_ = (self).fm0(((d_848_v119_)[p0] if (p0) in (d_848_v119_) else d_847_v118_), (_dafny.MultiSet([not(p0), p0, p0])).cardinality, _dafny.SeqWithoutIsStrInference([_dafny.Set({d_833_v107_}), _dafny.Set({d_702_v0_, d_833_v107_})]), (d_706_v4_)[default__.safeIndex(d_702_v0_, len(d_706_v4_))], globalState)
                rhs111_ = default__.fm43(_dafny.CodePoint('v'), D0_DC1(d_702_v0_, p0, d_852_v122_, d_702_v0_, d_831_cf23_), p0, (_dafny.MultiSet([p0])).set(p0, default__.abs(d_833_v107_)), globalState)
                rhs112_ = p0
                lhs89_ = globalState
                lhs90_ = d_785_v74_
                lhs91_ = default__.safeIndex(250, (d_785_v74_).length(0))
                lhs89_.f6 = rhs110_
                d_846_v117_ = rhs111_
                lhs90_[lhs91_] = rhs112_
            elif True:
                (globalState).f0 = (p0 if (d_706_v4_) < (_dafny.SeqWithoutIsStrInference([p0])) else p0)
                d_853_v123_: D9
                d_853_v123_ = D9_DC31(d_702_v0_, d_832_v106_)
                d_854_v124_: _dafny.Map
                d_854_v124_ = _dafny.Map({p0: d_702_v0_})
                d_855_v125_: _dafny.Map
                d_855_v125_ = _dafny.Map({d_702_v0_: d_854_v124_})
                d_856_v126_: _dafny.Array
                nw115_ = _dafny.Array(None, 8)
                nw115_[int(0)] = _dafny.SeqWithoutIsStrInference([(d_853_v123_).cf52])
                nw115_[int(1)] = (self).f22
                nw115_[int(2)] = (self).f22
                nw115_[int(3)] = _dafny.SeqWithoutIsStrInference([d_830_cf24_])
                nw115_[int(4)] = (self).f22
                nw115_[int(5)] = (self).f22
                nw115_[int(6)] = ((self).f22) + ((self).f22)
                nw115_[int(7)] = (((self).f22).set(default__.safeIndex(len(d_855_v125_), len((self).f22)), d_830_cf24_)) + ((self).f22)
                d_856_v126_ = nw115_
                index120_ = default__.safeIndex(20, (d_856_v126_).length(0))
                (d_856_v126_)[index120_] = _dafny.SeqWithoutIsStrInference([d_832_v106_])
                index121_ = default__.safeIndex(20, (d_856_v126_).length(0))
                (d_856_v126_)[index121_] = (self).f22
                d_857_v127_: C1
                nw116_ = C1()
                nw116_.ctor__((0) - (d_702_v0_), ((self).f22) + ((d_856_v126_)[default__.safeIndex(20, (d_856_v126_).length(0))]))
                d_857_v127_ = nw116_
                d_858_v128_: _dafny.Array
                nw117_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_858_v128_ = nw117_
                d_858_v128_ = (d_858_v128_ if p0 else d_858_v128_)
                (globalState).f8 = default__.safeModuloInt(d_702_v0_, d_702_v0_)
            d_859_v129_: _dafny.Array
            nw118_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_859_v129_ = nw118_
            rhs113_ = d_859_v129_
            rhs114_ = d_833_v107_
            lhs92_ = globalState
            d_859_v129_ = rhs113_
            lhs92_.f6 = rhs114_
        elif source10_.is_DC15:
            d_860___mcc_h4_ = source10_.cf26
            d_861___mcc_h5_ = source10_.cf27
            d_862_cf27_ = d_861___mcc_h5_
            d_863_cf26_ = d_860___mcc_h4_
            r1 = d_785_v74_
            (globalState).f16 = len(d_703_v1_)
            (globalState).f0 = (d_702_v0_) < (d_702_v0_)
            d_784_v73_ = (D10_DC34(_dafny.CodePoint('b'), p0, p0, p0, p0)).cf55
        elif source10_.is_DC12:
            d_864___mcc_h6_ = source10_.cf21
            d_865_cf21_ = d_864___mcc_h6_
            (globalState).f5 = d_706_v4_
            d_866_v130_: _dafny.MultiSet
            d_866_v130_ = _dafny.MultiSet([True, not(p0), p0])
            d_866_v130_ = _dafny.MultiSet([(d_706_v4_)[default__.safeIndex(d_702_v0_, len(d_706_v4_))], p0])
            d_867_v131_: D7
            d_867_v131_ = D7_DC26(d_704_v2_, d_702_v0_, d_702_v0_)
            pat_let_tv13_ = d_704_v2_
            def iife68_(_pat_let12_0):
                def iife69_(d_868_dt__update__tmp_h1_):
                    def iife70_(_pat_let13_0):
                        def iife71_(d_869_dt__update_hcf42_h0_):
                            return D7_DC26(d_869_dt__update_hcf42_h0_, (d_868_dt__update__tmp_h1_).cf43, (d_868_dt__update__tmp_h1_).cf44)
                        return iife71_(_pat_let13_0)
                    return iife70_(pat_let_tv13_)
                return iife69_(_pat_let12_0)
            source11_ = iife68_(d_867_v131_)
            if source11_.is_DC26:
                d_870___mcc_h8_ = source11_.cf42
                d_871___mcc_h9_ = source11_.cf43
                d_872___mcc_h10_ = source11_.cf44
                d_873_cf44_ = d_872___mcc_h10_
                d_874_cf43_ = d_871___mcc_h9_
                d_875_cf42_ = d_870___mcc_h8_
                (globalState).f6 = d_702_v0_
                d_876_v132_: C0
                nw119_ = C0()
                nw119_.ctor__()
                d_876_v132_ = nw119_
                rhs115_ = not(p0)
                rhs116_ = (d_706_v4_) == (d_706_v4_)
                lhs93_ = globalState
                lhs94_ = globalState
                lhs93_.f10 = rhs115_
                lhs94_.f10 = rhs116_
                (globalState).f10 = default__.fm15(not(p0), globalState)
            elif source11_.is_DC25:
                d_877___mcc_h11_ = source11_.cf41
                d_878_cf41_ = d_877___mcc_h11_
                (globalState).f6 = default__.safeDivisionInt((0) - ((0) - (d_702_v0_)), d_702_v0_)
                d_879_v133_: D9
                d_879_v133_ = D9_DC31(default__.safeDivisionInt(d_702_v0_, d_702_v0_), d_785_v74_)
                d_879_v133_ = d_879_v133_
                (globalState).f11 = d_702_v0_
                (globalState).f6 = default__.safeModuloInt((0) - (len(d_712_v9_)), default__.safeModuloInt(d_702_v0_, d_702_v0_))
            elif True:
                d_880___mcc_h12_ = source11_.cf45
                d_881_cf45_ = d_880___mcc_h12_
                index122_ = default__.safeIndex(99, (d_785_v74_).length(0))
                (d_785_v74_)[index122_] = p0
                d_882_v134_: _dafny.MultiSet
                d_882_v134_ = _dafny.MultiSet([d_785_v74_, d_785_v74_])
                index123_ = default__.safeIndex(99, (d_785_v74_).length(0))
                (d_785_v74_)[index123_] = not(((d_882_v134_).intersection(d_882_v134_)).issubset((d_882_v134_) | (d_882_v134_)))
                (globalState).f11 = d_702_v0_
                d_883_v135_: C2
                nw120_ = C2()
                nw120_.ctor__(d_702_v0_, (self).f22)
                d_883_v135_ = nw120_
                d_884_v136_: _dafny.Map
                d_884_v136_ = _dafny.Map({(d_883_v135_).f29: d_702_v0_})
                d_885_v137_: D0
                d_885_v137_ = D0_DC1((d_883_v135_).f29, (d_785_v74_)[default__.safeIndex(99, (d_785_v74_).length(0))], _dafny.Map({not((d_785_v74_)[default__.safeIndex(99, (d_785_v74_).length(0))]): d_884_v136_}), d_702_v0_, d_784_v73_)
                d_886_v138_: _dafny.Set
                d_886_v138_ = _dafny.Set({d_885_v137_})
                index124_ = default__.safeIndex(99, (d_785_v74_).length(0))
                rhs117_ = (len(((_dafny.SeqWithoutIsStrInference([d_784_v73_, d_784_v73_])) + (d_703_v1_)).set(default__.safeIndex(d_702_v0_, len((_dafny.SeqWithoutIsStrInference([d_784_v73_, d_784_v73_])) + (d_703_v1_))), d_784_v73_))) + (d_702_v0_)
                rhs118_ = (d_706_v4_)[default__.safeIndex((d_883_v135_).f29, len(d_706_v4_))]
                rhs119_ = d_883_v135_
                rhs120_ = (d_883_v135_).fm21(d_886_v138_, d_702_v0_, globalState)
                rhs121_ = p0
                lhs95_ = globalState
                lhs96_ = globalState
                lhs97_ = d_785_v74_
                lhs98_ = default__.safeIndex(99, (d_785_v74_).length(0))
                lhs99_ = globalState
                lhs95_.f11 = rhs117_
                lhs96_.f10 = rhs118_
                d_883_v135_ = rhs119_
                lhs97_[lhs98_] = rhs120_
                lhs99_.f10 = rhs121_
                index125_ = default__.safeIndex(99, (d_785_v74_).length(0))
                (d_785_v74_)[index125_] = p0
            d_887_v139_: _dafny.Map
            d_887_v139_ = _dafny.Map({d_702_v0_: d_784_v73_})
            d_888_v140_: _dafny.Map
            d_888_v140_ = _dafny.Map({d_744_v35_: d_702_v0_})
            d_889_v141_: _dafny.Map
            d_889_v141_ = _dafny.Map({d_888_v140_: d_887_v139_})
            d_890_v142_: _dafny.Seq
            d_890_v142_ = _dafny.SeqWithoutIsStrInference([d_744_v35_])
            d_891_v146_: _dafny.Map
            d_891_v146_ = _dafny.Map({False: d_887_v139_})
            d_892_v147_: _dafny.Array
            nw121_ = _dafny.Array(None, 26)
            nw121_[int(0)] = d_887_v139_
            nw121_[int(1)] = d_887_v139_
            nw121_[int(2)] = ((d_889_v141_)[d_888_v140_] if (d_888_v140_) in (d_889_v141_) else d_887_v139_)
            nw121_[int(3)] = default__.fm44(p0, not(not(True)), p0, d_703_v1_, globalState)
            nw121_[int(4)] = (_dafny.Map({len(d_890_v142_): _dafny.CodePoint('i')})) | (d_887_v139_)
            nw121_[int(5)] = d_887_v139_
            def iife72_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(47, 90):
                    d_893_v143_: int = compr_44_
                    if ((47) <= (d_893_v143_)) and ((d_893_v143_) < (90)):
                        coll44_[(d_893_v143_) * (d_702_v0_)] = d_784_v73_
                return _dafny.Map(coll44_)
            nw121_[int(6)] = iife72_()
            
            nw121_[int(7)] = _dafny.Map({175: d_784_v73_})
            nw121_[int(8)] = d_887_v139_
            nw121_[int(9)] = d_887_v139_
            def iife73_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(-152, 201):
                    d_894_v144_: int = compr_45_
                    if ((-152) <= (d_894_v144_)) and ((d_894_v144_) < (201)):
                        coll45_[(d_894_v144_) + (d_702_v0_)] = ((d_887_v139_)[d_702_v0_] if (d_702_v0_) in (d_887_v139_) else d_784_v73_)
                return _dafny.Map(coll45_)
            nw121_[int(10)] = (iife73_()
            ) | ((d_887_v139_).set((0) - (d_702_v0_), d_784_v73_))
            def iife74_():
                coll46_ = _dafny.Map()
                compr_46_: int
                for compr_46_ in (_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([978]))).cardinality, d_702_v0_, 314})).Elements:
                    d_895_v145_: int = compr_46_
                    if (d_895_v145_) in (_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([978]))).cardinality, d_702_v0_, 314})):
                        coll46_[default__.safeModuloInt(d_895_v145_, d_702_v0_)] = d_784_v73_
                return _dafny.Map(coll46_)
            nw121_[int(11)] = (iife74_()
            ) | (d_887_v139_)
            nw121_[int(12)] = d_887_v139_
            nw121_[int(13)] = d_887_v139_
            nw121_[int(14)] = d_887_v139_
            nw121_[int(15)] = (d_887_v139_) | (d_887_v139_)
            nw121_[int(16)] = (_dafny.Map({d_702_v0_: d_784_v73_})).set(len(_dafny.Map({d_702_v0_: 242})), d_784_v73_)
            nw121_[int(17)] = (d_887_v139_).set(780, d_784_v73_)
            nw121_[int(18)] = (d_887_v139_) | (d_887_v139_)
            nw121_[int(19)] = d_887_v139_
            nw121_[int(20)] = (d_887_v139_) | (d_887_v139_)
            nw121_[int(21)] = d_887_v139_
            nw121_[int(22)] = ((d_891_v146_)[True] if (True) in (d_891_v146_) else (d_887_v139_).set(d_702_v0_, d_784_v73_))
            nw121_[int(23)] = (d_887_v139_) | (d_887_v139_)
            nw121_[int(24)] = (d_887_v139_).set(d_702_v0_, d_784_v73_)
            nw121_[int(25)] = d_887_v139_
            d_892_v147_ = nw121_
            index126_ = default__.safeIndex(268, (d_892_v147_).length(0))
            (d_892_v147_)[index126_] = _dafny.Map({d_702_v0_: d_784_v73_})
            index127_ = default__.safeIndex(268, (d_892_v147_).length(0))
            (d_892_v147_)[index127_] = (d_887_v139_) | (d_887_v139_)
        elif True:
            d_896___mcc_h7_ = source10_.cf28
            d_897_cf28_ = d_896___mcc_h7_
            index128_ = default__.safeIndex(80, (d_744_v35_).length(0))
            (d_744_v35_)[index128_] = d_702_v0_
            index129_ = default__.safeIndex(80, (d_744_v35_).length(0))
            (d_744_v35_)[index129_] = (d_702_v0_) + ((d_702_v0_ if p0 else d_702_v0_))
            (globalState).f0 = p0
            d_898_v148_: _dafny.MultiSet
            d_898_v148_ = _dafny.MultiSet([d_702_v0_, default__.safeDivisionInt(d_702_v0_, d_702_v0_)])
            d_898_v148_ = d_898_v148_
            d_899_v149_: _dafny.Array
            nw122_ = _dafny.Array(D12.default()(), 14)
            d_899_v149_ = nw122_
            d_900_v150_: C0
            nw123_ = C0()
            nw123_.ctor__()
            d_900_v150_ = nw123_
            d_901_v151_: _dafny.MultiSet
            d_901_v151_ = _dafny.MultiSet([d_900_v150_, d_900_v150_])
            d_902_v152_: D12
            d_902_v152_ = D12_DC41(d_901_v151_)
            index130_ = default__.safeIndex(646, (d_899_v149_).length(0))
            (d_899_v149_)[index130_] = d_902_v152_
            index131_ = default__.safeIndex(80, (d_744_v35_).length(0))
            index132_ = default__.safeIndex(646, (d_899_v149_).length(0))
            index133_ = default__.safeIndex(80, (d_744_v35_).length(0))
            rhs122_ = (d_744_v35_)[default__.safeIndex(80, (d_744_v35_).length(0))]
            rhs123_ = d_902_v152_
            rhs124_ = p0
            rhs125_ = p0
            rhs126_ = (self).fm0((d_898_v148_).intersection(d_898_v148_), len((d_703_v1_) + (d_703_v1_)), d_708_v6_, p0, globalState)
            lhs100_ = d_744_v35_
            lhs101_ = default__.safeIndex(80, (d_744_v35_).length(0))
            lhs102_ = d_899_v149_
            lhs103_ = default__.safeIndex(646, (d_899_v149_).length(0))
            lhs104_ = globalState
            lhs105_ = globalState
            lhs106_ = d_744_v35_
            lhs107_ = default__.safeIndex(80, (d_744_v35_).length(0))
            lhs100_[lhs101_] = rhs122_
            lhs102_[lhs103_] = rhs123_
            lhs104_.f0 = rhs124_
            lhs105_.f0 = rhs125_
            lhs106_[lhs107_] = rhs126_
        r0 = not((p0 if not(not(p0)) else False))
        def lambda30_(d_903_p0_):
            def lambda31_(d_904_i10_):
                return (d_903_p0_) == (d_903_p0_)

            return lambda31_

        init16_ = lambda30_(p0)
        nw124_ = _dafny.Array(None, 25)
        for i0_16_ in range(nw124_.length(0)):
            nw124_[i0_16_] = init16_(i0_16_)
        r1 = nw124_
        d_905_v153_: _dafny.Map
        d_905_v153_ = _dafny.Map({d_702_v0_: d_702_v0_})
        d_906_v154_: D5
        d_906_v154_ = D5_DC19(d_905_v153_, p0)
        pat_let_tv14_ = d_704_v2_
        pat_let_tv15_ = d_784_v73_
        pat_let_tv16_ = d_704_v2_
        pat_let_tv17_ = d_784_v73_
        pat_let_tv18_ = d_702_v0_
        pat_let_tv19_ = p0
        pat_let_tv20_ = d_784_v73_
        pat_let_tv21_ = d_704_v2_
        pat_let_tv22_ = d_784_v73_
        pat_let_tv23_ = d_704_v2_
        pat_let_tv24_ = d_784_v73_
        pat_let_tv25_ = d_704_v2_
        pat_let_tv26_ = d_784_v73_
        pat_let_tv27_ = d_704_v2_
        pat_let_tv28_ = d_784_v73_
        pat_let_tv29_ = d_704_v2_
        pat_let_tv30_ = d_784_v73_
        pat_let_tv31_ = d_704_v2_
        pat_let_tv32_ = d_702_v0_
        pat_let_tv33_ = d_704_v2_
        pat_let_tv34_ = d_702_v0_
        pat_let_tv35_ = d_784_v73_
        def lambda32_(source12_):
            if source12_.is_DC18:
                d_907___mcc_h13_ = source12_.cf30
                d_908___mcc_h14_ = source12_.cf31
                d_909_cf31_ = d_908___mcc_h14_
                d_910_cf30_ = d_907___mcc_h13_
                return (_dafny.Map({pat_let_tv14_: pat_let_tv15_})) | ((_dafny.Map({pat_let_tv16_: pat_let_tv17_})).set(_dafny.SeqWithoutIsStrInference([pat_let_tv18_ for d_911_i11_ in range(default__.abs(549))]), _dafny.CodePoint('t')))
            elif source12_.is_DC19:
                d_912___mcc_h15_ = source12_.cf32
                d_913___mcc_h16_ = source12_.cf33
                d_914_cf33_ = d_913___mcc_h16_
                d_915_cf32_ = d_912___mcc_h15_
                if pat_let_tv19_:
                    return _dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))) for d_916_i12_ in range(default__.abs(851))]): pat_let_tv20_})
                elif True:
                    return _dafny.Map({pat_let_tv21_: pat_let_tv22_})
            elif source12_.is_DC17:
                d_917___mcc_h17_ = source12_.cf29
                d_918_cf29_ = d_917___mcc_h17_
                return (_dafny.Map({pat_let_tv23_: pat_let_tv24_})) | (_dafny.Map({pat_let_tv25_: pat_let_tv26_}))
            elif True:
                d_919___mcc_h18_ = source12_.cf34
                d_920_cf34_ = d_919___mcc_h18_
                return ((_dafny.Map({pat_let_tv27_: pat_let_tv28_})).set(pat_let_tv29_, pat_let_tv30_)) | (_dafny.Map({(pat_let_tv31_).set(default__.safeIndex(pat_let_tv32_, len(pat_let_tv33_)), pat_let_tv34_): pat_let_tv35_}))

        r2 = lambda32_(d_906_v154_)
        r3 = d_702_v0_
        return r0, r1, r2, r3

    def m4(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        d_921_v0_: bool
        d_921_v0_ = True
        d_922_v1_: _dafny.Set
        d_922_v1_ = _dafny.Set({default__.fm15(d_921_v0_, globalState)})
        d_923_v2_: int
        d_923_v2_ = 54
        d_924_v3_: _dafny.MultiSet
        d_924_v3_ = _dafny.MultiSet([d_923_v2_, d_923_v2_, d_923_v2_])
        d_925_v4_: _dafny.MultiSet
        d_925_v4_ = _dafny.MultiSet([d_924_v3_])
        d_926_v5_: _dafny.Seq
        d_926_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftnyfkc"))
        d_927_v6_: _dafny.Set
        d_927_v6_ = _dafny.Set({len(d_926_v5_)})
        d_928_v7_: _dafny.Map
        d_928_v7_ = _dafny.Map({d_921_v0_: d_923_v2_})
        d_929_v8_: _dafny.Seq
        d_929_v8_ = _dafny.SeqWithoutIsStrInference([d_921_v0_])
        d_930_v9_: _dafny.Seq
        d_930_v9_ = _dafny.SeqWithoutIsStrInference([d_927_v6_])
        d_931_v10_: _dafny.Array
        def lambda33_(d_932_v0_):
            def lambda34_(d_933_i0_):
                return d_932_v0_

            return lambda34_

        init17_ = lambda33_(d_921_v0_)
        nw125_ = _dafny.Array(None, 7)
        for i0_17_ in range(nw125_.length(0)):
            nw125_[i0_17_] = init17_(i0_17_)
        d_931_v10_ = nw125_
        d_934_v11_: _dafny.Map
        d_934_v11_ = _dafny.Map({d_921_v0_: d_931_v10_})
        d_935_v12_: _dafny.Array
        nw126_ = _dafny.Array(None, 12)
        nw126_[int(0)] = ((_dafny.MultiSet([d_924_v3_])).intersection(d_925_v4_)).cardinality
        nw126_[int(1)] = len(d_927_v6_)
        nw126_[int(2)] = ((d_928_v7_)[d_921_v0_] if (d_921_v0_) in (d_928_v7_) else len(d_929_v8_))
        nw126_[int(3)] = (len(d_929_v8_) if not(d_921_v0_) else d_923_v2_)
        nw126_[int(4)] = (self).fm0(d_924_v3_, len(d_926_v5_), d_930_v9_, d_921_v0_, globalState)
        nw126_[int(5)] = (self).fm0(_dafny.MultiSet([len(d_929_v8_)]), -673, d_930_v9_, False, globalState)
        nw126_[int(6)] = default__.safeModuloInt(d_923_v2_, d_923_v2_)
        nw126_[int(7)] = len((d_934_v11_ if d_921_v0_ else d_934_v11_))
        nw126_[int(8)] = default__.safeModuloInt(d_923_v2_, d_923_v2_)
        nw126_[int(9)] = len(d_926_v5_)
        nw126_[int(10)] = d_923_v2_
        nw126_[int(11)] = d_923_v2_
        d_935_v12_ = nw126_
        index134_ = default__.safeIndex(67, (d_935_v12_).length(0))
        (d_935_v12_)[index134_] = (d_923_v2_) * (d_923_v2_)
        d_936_v13_: _dafny.MultiSet
        d_936_v13_ = _dafny.MultiSet([d_921_v0_])
        index135_ = default__.safeIndex(85, (d_931_v10_).length(0))
        (d_931_v10_)[index135_] = (d_936_v13_).ispropersubset(d_936_v13_)
        d_937_v14_: _dafny.Map
        d_937_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dkfdq")), d_926_v5_]): d_921_v0_})
        d_938_v15_: _dafny.Seq
        d_938_v15_ = _dafny.SeqWithoutIsStrInference([d_926_v5_])
        d_939_v16_: str
        d_939_v16_ = _dafny.CodePoint('i')
        d_940_v17_: D8
        d_940_v17_ = D8_DC29(d_921_v0_, d_921_v0_, d_939_v16_)
        d_941_v18_: D13
        d_941_v18_ = D13_DC44(d_929_v8_)
        d_942_v19_: _dafny.Map
        d_942_v19_ = _dafny.Map({d_923_v2_: d_921_v0_})
        pat_let_tv36_ = d_921_v0_
        index136_ = default__.safeIndex(67, (d_935_v12_).length(0))
        index137_ = default__.safeIndex(85, (d_931_v10_).length(0))
        def iife75_(_pat_let14_0):
            def iife76_(d_943_dt__update__tmp_h0_):
                def iife77_(_pat_let15_0):
                    def iife78_(d_944_dt__update_hcf47_h0_):
                        return D8_DC29(d_944_dt__update_hcf47_h0_, (d_943_dt__update__tmp_h0_).cf48, (d_943_dt__update__tmp_h0_).cf49)
                    return iife78_(_pat_let15_0)
                return iife77_(pat_let_tv36_)
            return iife76_(_pat_let14_0)
        rhs127_ = _dafny.Set({((d_937_v14_)[d_938_v15_] if (d_938_v15_) in (d_937_v14_) else d_921_v0_)})
        rhs128_ = 971
        rhs129_ = ((d_923_v2_) <= (len(_dafny.Set({d_940_v17_, iife75_(d_940_v17_), d_940_v17_})))) in (d_929_v8_)
        rhs130_ = (d_926_v5_) == (default__.fm18(d_921_v0_, globalState))
        rhs131_ = default__.fm37((d_941_v18_).cf71, (d_942_v19_) | (default__.fm40(d_923_v2_, (0) - (d_923_v2_), default__.fm38(True, d_923_v2_, globalState), globalState)), d_926_v5_, globalState)
        lhs108_ = d_935_v12_
        lhs109_ = default__.safeIndex(67, (d_935_v12_).length(0))
        lhs110_ = d_931_v10_
        lhs111_ = default__.safeIndex(85, (d_931_v10_).length(0))
        lhs112_ = globalState
        d_922_v1_ = rhs127_
        lhs108_[lhs109_] = rhs128_
        lhs110_[lhs111_] = rhs129_
        lhs112_.f0 = rhs130_
        d_921_v0_ = rhs131_
        d_945_v20_: _dafny.MultiSet
        d_945_v20_ = _dafny.MultiSet([d_940_v17_])
        d_946_v21_: D11
        d_946_v21_ = D11_DC36(_dafny.SeqWithoutIsStrInference([d_924_v3_]))
        d_947_v22_: _dafny.Seq
        d_947_v22_ = _dafny.SeqWithoutIsStrInference([(d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]])
        d_948_v23_: D2
        d_948_v23_ = D2_DC7(d_947_v22_, (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], d_923_v2_)
        pat_let_tv37_ = d_921_v0_
        pat_let_tv38_ = d_931_v10_
        pat_let_tv39_ = d_931_v10_
        pat_let_tv40_ = d_921_v0_
        pat_let_tv41_ = d_931_v10_
        pat_let_tv42_ = d_931_v10_
        pat_let_tv43_ = d_921_v0_
        pat_let_tv44_ = d_931_v10_
        pat_let_tv45_ = d_931_v10_
        pat_let_tv46_ = d_931_v10_
        pat_let_tv47_ = d_931_v10_
        index138_ = default__.safeIndex(85, (d_931_v10_).length(0))
        def lambda35_(source13_):
            if source13_.is_DC37:
                d_949___mcc_h0_ = source13_.cf63
                d_950_cf63_ = d_949___mcc_h0_
                return (True) or (pat_let_tv37_)
            elif source13_.is_DC38:
                d_951___mcc_h1_ = source13_.cf64
                d_952___mcc_h2_ = source13_.cf65
                d_953_cf65_ = d_952___mcc_h2_
                d_954_cf64_ = d_951___mcc_h1_
                if (pat_let_tv39_)[default__.safeIndex(85, (pat_let_tv38_).length(0))]:
                    return pat_let_tv40_
                elif True:
                    return (pat_let_tv42_)[default__.safeIndex(85, (pat_let_tv41_).length(0))]
            elif source13_.is_DC39:
                d_955___mcc_h3_ = source13_.cf66
                d_956_cf66_ = d_955___mcc_h3_
                return pat_let_tv43_
            elif source13_.is_DC36:
                d_957___mcc_h4_ = source13_.cf62
                d_958_cf62_ = d_957___mcc_h4_
                return not (False) or ((pat_let_tv45_)[default__.safeIndex(85, (pat_let_tv44_).length(0))])
            elif True:
                d_959___mcc_h5_ = source13_.cf67
                d_960_cf67_ = d_959___mcc_h5_
                return (pat_let_tv47_)[default__.safeIndex(85, (pat_let_tv46_).length(0))]

        rhs132_ = not(lambda35_(d_946_v21_))
        rhs133_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
        rhs134_ = default__.fm22((d_923_v2_) < ((d_948_v23_).cf13), (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], d_923_v2_, d_923_v2_, globalState)
        rhs135_ = d_945_v20_
        lhs113_ = d_931_v10_
        lhs114_ = default__.safeIndex(85, (d_931_v10_).length(0))
        lhs113_[lhs114_] = rhs132_
        d_923_v2_ = rhs133_
        d_939_v16_ = rhs134_
        d_945_v20_ = rhs135_
        d_961_v24_: D9
        d_961_v24_ = D9_DC31(618, d_931_v10_)
        d_962_v25_: D9
        d_962_v25_ = D9_DC32(d_961_v24_)
        source14_ = d_962_v25_
        if source14_.is_DC31:
            d_963___mcc_h6_ = source14_.cf51
            d_964___mcc_h7_ = source14_.cf52
            d_965_cf52_ = d_964___mcc_h7_
            d_966_cf51_ = d_963___mcc_h6_
            rhs136_ = not(d_921_v0_)
            rhs137_ = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
            rhs138_ = 837
            lhs115_ = globalState
            lhs116_ = globalState
            lhs115_.f0 = rhs136_
            lhs116_.f0 = rhs137_
            r1 = rhs138_
            index139_ = default__.safeIndex(85, (d_931_v10_).length(0))
            (d_931_v10_)[index139_] = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
            d_967_v26_: _dafny.Map
            d_967_v26_ = _dafny.Map({d_928_v7_: d_921_v0_})
            d_967_v26_ = (d_967_v26_).set(d_928_v7_, d_921_v0_)
            d_968_v27_: _dafny.Map
            d_968_v27_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")): (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]})
            if ((d_968_v27_)[d_926_v5_] if (d_926_v5_) in (d_968_v27_) else (not(not((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))])) if (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))] else d_921_v0_)):
                d_929_v8_ = d_929_v8_
                d_969_v28_: _dafny.Map
                d_969_v28_ = _dafny.Map({(d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]: (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]})
                d_970_v29_: D5
                d_970_v29_ = D5_DC19(d_969_v28_, d_921_v0_)
                d_971_v30_: _dafny.Seq
                d_971_v30_ = _dafny.SeqWithoutIsStrInference([d_970_v29_, d_970_v29_, d_970_v29_])
                d_972_v31_: _dafny.Map
                d_972_v31_ = _dafny.Map({_dafny.Set({d_921_v0_, (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]}): d_971_v30_})
                rhs139_ = d_935_v12_
                rhs140_ = d_931_v10_
                rhs141_ = ((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]) or (d_921_v0_)
                rhs142_ = (_dafny.Map({d_922_v1_: d_971_v30_})) | (d_972_v31_)
                d_935_v12_ = rhs139_
                r0 = rhs140_
                d_921_v0_ = rhs141_
                d_972_v31_ = rhs142_
                index140_ = default__.safeIndex(85, (d_931_v10_).length(0))
                (d_931_v10_)[index140_] = d_921_v0_
                d_973_v32_: _dafny.Array
                nw127_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_973_v32_ = nw127_
                d_974_v33_: _dafny.Map
                d_974_v33_ = _dafny.Map({d_921_v0_: d_926_v5_})
                index141_ = default__.safeIndex(153, (d_973_v32_).length(0))
                (d_973_v32_)[index141_] = ((d_974_v33_)[d_921_v0_] if (d_921_v0_) in (d_974_v33_) else d_926_v5_)
                index142_ = default__.safeIndex(153, (d_973_v32_).length(0))
                (d_973_v32_)[index142_] = ((d_926_v5_) + (d_926_v5_)) + ((d_926_v5_).set(default__.safeIndex(d_923_v2_, len(d_926_v5_)), d_939_v16_))
                index143_ = default__.safeIndex(153, (d_973_v32_).length(0))
                (d_973_v32_)[index143_] = default__.fm18(d_921_v0_, globalState)
            elif True:
                (globalState).f0 = (d_921_v0_) or (True)
                (globalState).f10 = d_921_v0_
                d_975_v34_: _dafny.Map
                d_975_v34_ = _dafny.Map({d_947_v22_: d_966_cf51_})
                d_976_v35_: _dafny.Map
                d_976_v35_ = _dafny.Map({d_921_v0_: d_942_v19_})
                d_977_v36_: _dafny.Map
                d_977_v36_ = _dafny.Map({d_926_v5_: (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]})
                index144_ = default__.safeIndex(85, (d_931_v10_).length(0))
                index145_ = default__.safeIndex(67, (d_935_v12_).length(0))
                rhs143_ = ((_dafny.Map({d_947_v22_: len((d_926_v5_).set(default__.safeIndex((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], len(d_926_v5_)), _dafny.CodePoint('b')))})) == (d_975_v34_)) in (d_976_v35_)
                rhs144_ = d_921_v0_
                rhs145_ = default__.safeModuloInt(((d_977_v36_)[_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_978_i1_ in range(default__.abs(644))])] if (_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_979_i1_ in range(default__.abs(644))])) in (d_977_v36_) else (0) - (d_923_v2_)), (d_966_cf51_ if d_921_v0_ else d_966_cf51_))
                lhs117_ = d_931_v10_
                lhs118_ = default__.safeIndex(85, (d_931_v10_).length(0))
                lhs119_ = d_935_v12_
                lhs120_ = default__.safeIndex(67, (d_935_v12_).length(0))
                d_921_v0_ = rhs143_
                lhs117_[lhs118_] = rhs144_
                lhs119_[lhs120_] = rhs145_
                d_980_v37_: D1
                d_980_v37_ = D1_DC3(d_926_v5_)
                d_981_v38_: D2
                d_981_v38_ = D2_DC8(d_980_v37_, d_935_v12_, (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))])
                d_982_v39_: _dafny.Seq
                d_982_v39_ = _dafny.SeqWithoutIsStrInference([d_935_v12_, d_935_v12_, d_935_v12_, d_935_v12_])
                d_983_v40_: _dafny.Array
                nw128_ = _dafny.Array(None, 20)
                nw128_[int(0)] = d_935_v12_
                nw128_[int(1)] = d_935_v12_
                nw128_[int(2)] = d_935_v12_
                nw128_[int(3)] = d_935_v12_
                nw128_[int(4)] = d_935_v12_
                nw128_[int(5)] = d_935_v12_
                nw128_[int(6)] = d_935_v12_
                nw128_[int(7)] = (d_981_v38_).cf15
                nw128_[int(8)] = d_935_v12_
                nw128_[int(9)] = d_935_v12_
                nw128_[int(10)] = d_935_v12_
                nw128_[int(11)] = d_935_v12_
                nw128_[int(12)] = d_935_v12_
                nw128_[int(13)] = d_935_v12_
                nw128_[int(14)] = (d_982_v39_)[default__.safeIndex((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], len(d_982_v39_))]
                nw128_[int(15)] = d_935_v12_
                nw128_[int(16)] = d_935_v12_
                nw128_[int(17)] = d_935_v12_
                nw128_[int(18)] = d_935_v12_
                nw128_[int(19)] = d_935_v12_
                d_983_v40_ = nw128_
                index146_ = default__.safeIndex(903, (d_983_v40_).length(0))
                (d_983_v40_)[index146_] = (d_981_v38_).cf15
                d_984_v41_: _dafny.Array
                nw129_ = _dafny.Array(_dafny.Map({}), 15)
                d_984_v41_ = nw129_
                d_985_v42_: _dafny.Map
                d_985_v42_ = _dafny.Map({(0) - (d_966_cf51_): d_935_v12_})
                index147_ = default__.safeIndex(588, (d_984_v41_).length(0))
                (d_984_v41_)[index147_] = d_985_v42_
                d_986_v43_: _dafny.Map
                d_986_v43_ = _dafny.Map({d_921_v0_: _dafny.Map({d_966_cf51_: (d_982_v39_)[default__.safeIndex(d_966_cf51_, len(d_982_v39_))]})})
                d_987_v44_: _dafny.Map
                d_987_v44_ = _dafny.Map({d_921_v0_: (d_929_v8_)[default__.safeIndex(d_966_cf51_, len(d_929_v8_))]})
                index148_ = default__.safeIndex(903, (d_983_v40_).length(0))
                index149_ = default__.safeIndex(588, (d_984_v41_).length(0))
                rhs146_ = d_935_v12_
                rhs147_ = d_921_v0_
                rhs148_ = (d_985_v42_) | ((((d_986_v43_)[(d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]] if ((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]) in (d_986_v43_) else _dafny.Map({d_966_cf51_: d_935_v12_}))) | (d_985_v42_))
                rhs149_ = (0) - ((0) - (default__.safeDivisionInt(default__.safeDivisionInt(d_966_cf51_, 47), len(d_987_v44_))))
                lhs121_ = d_983_v40_
                lhs122_ = default__.safeIndex(903, (d_983_v40_).length(0))
                lhs123_ = globalState
                lhs124_ = d_984_v41_
                lhs125_ = default__.safeIndex(588, (d_984_v41_).length(0))
                lhs121_[lhs122_] = rhs146_
                lhs123_.f0 = rhs147_
                lhs124_[lhs125_] = rhs148_
                d_966_cf51_ = rhs149_
                (globalState).f8 = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
        elif source14_.is_DC30:
            d_988___mcc_h8_ = source14_.cf50
            d_989_cf50_ = d_988___mcc_h8_
            r1 = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
            (globalState).f0 = (d_929_v8_)[default__.safeIndex(d_923_v2_, len(d_929_v8_))]
            d_990_v45_: _dafny.Map
            d_990_v45_ = _dafny.Map({d_923_v2_: d_924_v3_})
            d_991_v46_: D4
            d_991_v46_ = D4_DC12(d_990_v45_)
            d_991_v46_ = d_991_v46_
            d_992_v47_: _dafny.Array
            nw130_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_992_v47_ = nw130_
            index150_ = default__.safeIndex(969, (d_992_v47_).length(0))
            (d_992_v47_)[index150_] = d_931_v10_
            index151_ = default__.safeIndex(969, (d_992_v47_).length(0))
            (d_992_v47_)[index151_] = d_931_v10_
        elif True:
            d_993___mcc_h9_ = source14_.cf53
            d_994_cf53_ = d_993___mcc_h9_
            d_923_v2_ = d_923_v2_
            index152_ = default__.safeIndex(85, (d_931_v10_).length(0))
            rhs150_ = (len(d_927_v6_)) - (len(d_947_v22_))
            rhs151_ = d_921_v0_
            lhs126_ = globalState
            lhs127_ = d_931_v10_
            lhs128_ = default__.safeIndex(85, (d_931_v10_).length(0))
            lhs126_.f16 = rhs150_
            lhs127_[lhs128_] = rhs151_
            index153_ = default__.safeIndex(67, (d_935_v12_).length(0))
            rhs152_ = d_921_v0_
            rhs153_ = ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) * ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
            lhs129_ = globalState
            lhs130_ = d_935_v12_
            lhs131_ = default__.safeIndex(67, (d_935_v12_).length(0))
            lhs129_.f10 = rhs152_
            lhs130_[lhs131_] = rhs153_
            d_926_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia"))
        d_995_v48_: _dafny.Map
        d_995_v48_ = _dafny.Map({(d_923_v2_) * (d_923_v2_): default__.safeModuloInt((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], d_923_v2_)})
        d_995_v48_ = _dafny.Map({d_923_v2_: (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]})
        if (d_923_v2_) >= (((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) + (-544)):
            d_996_v49_: _dafny.Array
            nw131_ = _dafny.Array(None, 6)
            d_996_v49_ = nw131_
            d_997_v50_: D11
            d_997_v50_ = D11_DC37(d_996_v49_)
            source15_ = d_997_v50_
            if source15_.is_DC37:
                d_998___mcc_h10_ = source15_.cf63
                d_999_cf63_ = d_998___mcc_h10_
                rhs154_ = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                rhs155_ = d_921_v0_
                rhs156_ = (212) * (d_923_v2_)
                rhs157_ = ((d_924_v3_)[(self).fm0(d_924_v3_, 467, _dafny.SeqWithoutIsStrInference([d_927_v6_ for d_1000_i2_ in range(default__.abs(-839))]), not(d_921_v0_), globalState)] if ((self).fm0(d_924_v3_, 467, _dafny.SeqWithoutIsStrInference([d_927_v6_ for d_1001_i2_ in range(default__.abs(-839))]), not(d_921_v0_), globalState)) in (d_924_v3_) else 750)
                rhs158_ = (d_923_v2_) - ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                lhs132_ = globalState
                lhs133_ = globalState
                lhs134_ = globalState
                lhs135_ = globalState
                lhs136_ = globalState
                lhs132_.f0 = rhs154_
                lhs133_.f0 = rhs155_
                lhs134_.f6 = rhs156_
                lhs135_.f6 = rhs157_
                lhs136_.f11 = rhs158_
                index154_ = default__.safeIndex(85, (d_931_v10_).length(0))
                rhs159_ = d_923_v2_
                rhs160_ = ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) == (d_923_v2_)
                rhs161_ = ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) * ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                lhs137_ = d_931_v10_
                lhs138_ = default__.safeIndex(85, (d_931_v10_).length(0))
                lhs139_ = globalState
                r1 = rhs159_
                lhs137_[lhs138_] = rhs160_
                lhs139_.f16 = rhs161_
                index155_ = default__.safeIndex(85, (d_931_v10_).length(0))
                (d_931_v10_)[index155_] = default__.fm37(d_929_v8_, (d_942_v19_).set(d_923_v2_, d_921_v0_), d_926_v5_, globalState)
                index156_ = default__.safeIndex(67, (d_935_v12_).length(0))
                index157_ = default__.safeIndex(67, (d_935_v12_).length(0))
                rhs162_ = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                rhs163_ = d_923_v2_
                rhs164_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                lhs140_ = globalState
                lhs141_ = d_935_v12_
                lhs142_ = default__.safeIndex(67, (d_935_v12_).length(0))
                lhs143_ = d_935_v12_
                lhs144_ = default__.safeIndex(67, (d_935_v12_).length(0))
                lhs140_.f10 = rhs162_
                lhs141_[lhs142_] = rhs163_
                lhs143_[lhs144_] = rhs164_
            elif source15_.is_DC38:
                d_1002___mcc_h11_ = source15_.cf64
                d_1003___mcc_h12_ = source15_.cf65
                d_1004_cf65_ = d_1003___mcc_h12_
                d_1005_cf64_ = d_1002___mcc_h11_
                def iife79_():
                    coll47_ = _dafny.Map()
                    compr_47_: int
                    for compr_47_ in (_dafny.SeqWithoutIsStrInference([925, d_1004_cf65_])).Elements:
                        d_1006_v51_: int = compr_47_
                        if (d_1006_v51_) in (_dafny.SeqWithoutIsStrInference([925, d_1004_cf65_])):
                            coll47_[(d_1006_v51_) + ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])] = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                    return _dafny.Map(coll47_)
                d_942_v19_ = iife79_()
                
                index158_ = default__.safeIndex(85, (d_931_v10_).length(0))
                (d_931_v10_)[index158_] = d_921_v0_
                d_934_v11_ = (d_934_v11_).set((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], d_931_v10_)
                index159_ = default__.safeIndex(67, (d_935_v12_).length(0))
                rhs165_ = (d_923_v2_) > (d_1005_cf64_)
                rhs166_ = d_1004_cf65_
                rhs167_ = ((d_995_v48_)[d_1005_cf64_] if (d_1005_cf64_) in (d_995_v48_) else ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) - (d_1005_cf64_))
                rhs168_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                lhs145_ = globalState
                lhs146_ = globalState
                lhs147_ = d_935_v12_
                lhs148_ = default__.safeIndex(67, (d_935_v12_).length(0))
                lhs149_ = globalState
                lhs145_.f10 = rhs165_
                lhs146_.f11 = rhs166_
                lhs147_[lhs148_] = rhs167_
                lhs149_.f6 = rhs168_
            elif source15_.is_DC39:
                d_1007___mcc_h13_ = source15_.cf66
                d_1008_cf66_ = d_1007___mcc_h13_
                d_1009_v52_: _dafny.Map
                d_1009_v52_ = _dafny.Map({d_1008_cf66_: d_921_v0_})
                d_1010_v53_: _dafny.Seq
                d_1010_v53_ = _dafny.SeqWithoutIsStrInference([d_929_v8_])
                d_1011_v54_: _dafny.Map
                d_1011_v54_ = _dafny.Map({d_1009_v52_: (_dafny.SeqWithoutIsStrInference([d_929_v8_])) != (d_1010_v53_)})
                d_1012_v55_: _dafny.Seq
                d_1012_v55_ = _dafny.SeqWithoutIsStrInference([D6_DC24(D6_DC21(_dafny.MultiSet([(d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], len(d_929_v8_), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]])))])
                d_1013_v56_: T0
                nw132_ = C2()
                nw132_.ctor__(len(d_926_v5_), (self).f22)
                d_1013_v56_ = nw132_
                d_1014_v57_: D6
                d_1014_v57_ = D6_DC23(d_1008_cf66_, len(_dafny.Map({35: d_1013_v56_})), d_921_v0_)
                d_1015_v58_: D6
                d_1015_v58_ = D6_DC24(d_1014_v57_)
                def iife80_():
                    coll48_ = _dafny.Map()
                    compr_48_: _dafny.Map
                    for compr_48_ in (d_1011_v54_).keys.Elements:
                        d_1016_v59_: _dafny.Map = compr_48_
                        if (d_1016_v59_) in (d_1011_v54_):
                            coll48_[d_1016_v59_] = d_921_v0_
                    return _dafny.Map(coll48_)
                rhs169_ = (d_1012_v55_) == (((d_1012_v55_) + (d_1012_v55_)).set(default__.safeIndex((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], len((d_1012_v55_) + (d_1012_v55_))), d_1015_v58_))
                rhs170_ = ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))] if ((d_942_v19_)[d_923_v2_] if (d_923_v2_) in (d_942_v19_) else (d_929_v8_)[default__.safeIndex(d_923_v2_, len(d_929_v8_))]) else d_923_v2_)
                rhs171_ = d_947_v22_
                rhs172_ = iife80_()

                rhs173_ = d_926_v5_
                lhs150_ = globalState
                lhs151_ = globalState
                lhs150_.f0 = rhs169_
                lhs151_.f8 = rhs170_
                d_947_v22_ = rhs171_
                d_1011_v54_ = rhs172_
                d_926_v5_ = rhs173_
                d_1009_v52_ = (d_1009_v52_).set((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], not ((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]) or ((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]))
                d_1017_v60_: _dafny.Array
                nw133_ = _dafny.Array(_dafny.MultiSet({}), 28)
                d_1017_v60_ = nw133_
                d_1018_v61_: _dafny.Seq
                d_1018_v61_ = _dafny.SeqWithoutIsStrInference([d_1017_v60_, d_1017_v60_, d_1017_v60_])
                d_1017_v60_ = (d_1018_v61_)[default__.safeIndex(default__.safeDivisionInt(d_923_v2_, (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]), len(d_1018_v61_))]
                d_1019_v62_: D5
                d_1019_v62_ = D5_DC18(d_923_v2_, (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                d_1020_v63_: _dafny.Seq
                d_1020_v63_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1019_v62_, d_1019_v62_, d_1019_v62_])])
                (globalState).f0 = not((d_1020_v63_) <= (d_1020_v63_))
            elif source15_.is_DC36:
                d_1021___mcc_h14_ = source15_.cf62
                d_1022_cf62_ = d_1021___mcc_h14_
                r1 = ((0) - (d_923_v2_) if d_921_v0_ else (self).fm0(_dafny.MultiSet([(d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], d_923_v2_]), (0) - (len(_dafny.SeqWithoutIsStrInference([d_921_v0_, False]))), _dafny.SeqWithoutIsStrInference([d_927_v6_ for d_1023_i3_ in range(default__.abs(-26))]), d_921_v0_, globalState))
                (globalState).f6 = (0) - (default__.safeDivisionInt(d_923_v2_, d_923_v2_))
                r0 = d_931_v10_
                index160_ = default__.safeIndex(67, (d_935_v12_).length(0))
                (d_935_v12_)[index160_] = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
            elif True:
                d_1024___mcc_h15_ = source15_.cf67
                d_1025_cf67_ = d_1024___mcc_h15_
                index161_ = default__.safeIndex(85, (d_931_v10_).length(0))
                (d_931_v10_)[index161_] = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                (globalState).f0 = d_921_v0_
                index162_ = default__.safeIndex(67, (d_935_v12_).length(0))
                (d_935_v12_)[index162_] = 411
                index163_ = default__.safeIndex(85, (d_931_v10_).length(0))
                rhs174_ = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                rhs175_ = d_923_v2_
                rhs176_ = (d_922_v1_ if (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))] else (_dafny.Set({(d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]}) if d_921_v0_ else d_922_v1_))
                rhs177_ = _dafny.SeqWithoutIsStrInference([((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))] if d_921_v0_ else d_923_v2_), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]])
                rhs178_ = (d_929_v8_)[default__.safeIndex(-182, len(d_929_v8_))]
                lhs152_ = d_931_v10_
                lhs153_ = default__.safeIndex(85, (d_931_v10_).length(0))
                lhs154_ = globalState
                lhs155_ = globalState
                lhs152_[lhs153_] = rhs174_
                lhs154_.f16 = rhs175_
                d_922_v1_ = rhs176_
                d_947_v22_ = rhs177_
                lhs155_.f0 = rhs178_
            if ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) >= (d_923_v2_):
                (globalState).f11 = (d_923_v2_) - (((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) + ((self).fm0(d_924_v3_, len(d_929_v8_), d_930_v9_, default__.fm15(d_921_v0_, globalState), globalState)))
                d_1026_v64_: C2
                nw134_ = C2()
                nw134_.ctor__((0) - ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]), ((self).f22) + (_dafny.SeqWithoutIsStrInference([d_931_v10_])))
                d_1026_v64_ = nw134_
                d_926_v5_ = d_926_v5_
                d_1027_v65_: _dafny.Array
                nw135_ = _dafny.Array(_dafny.CodePoint('D'), 26)
                d_1027_v65_ = nw135_
                d_1027_v65_ = d_1027_v65_
                d_1028_v66_: C2
                nw136_ = C2()
                nw136_.ctor__(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1029_i4_ in range(default__.abs(543))])), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]), (self).f22)
                d_1028_v66_ = nw136_
            elif True:
                d_926_v5_ = d_926_v5_
                (globalState).f11 = d_923_v2_
                (globalState).f16 = -93
                (globalState).f0 = (_dafny.Set({606})).issubset(d_927_v6_)
                d_1030_v67_: _dafny.Map
                d_1030_v67_ = _dafny.Map({d_921_v0_: d_922_v1_})
                d_1030_v67_ = d_1030_v67_
            (globalState).f0 = d_921_v0_
            d_923_v2_ = 484
            d_926_v5_ = (d_926_v5_) + ((d_938_v15_)[default__.safeIndex((0) - (d_923_v2_), len(d_938_v15_))])
        elif True:
            if ((d_923_v2_) > ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) if ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) < (len(d_926_v5_)) else (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]):
                (globalState).f6 = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                index164_ = default__.safeIndex(67, (d_935_v12_).length(0))
                (d_935_v12_)[index164_] = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                index165_ = default__.safeIndex(67, (d_935_v12_).length(0))
                (d_935_v12_)[index165_] = default__.safeDivisionInt((self).fm0(d_924_v3_, d_923_v2_, d_930_v9_, True, globalState), 511)
                d_1031_v68_: _dafny.Array
                def lambda36_(d_1032_v16_):
                    def lambda37_(d_1033_i5_):
                        return d_1032_v16_

                    return lambda37_

                init18_ = lambda36_(d_939_v16_)
                nw137_ = _dafny.Array(None, 21)
                for i0_18_ in range(nw137_.length(0)):
                    nw137_[i0_18_] = init18_(i0_18_)
                d_1031_v68_ = nw137_
                d_1034_v69_: _dafny.MultiSet
                d_1034_v69_ = _dafny.MultiSet([d_1031_v68_, d_1031_v68_])
                rhs179_ = d_921_v0_
                rhs180_ = d_921_v0_
                rhs181_ = (0) - ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                rhs182_ = ((d_1034_v69_ if (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))] else ((d_1034_v69_).set(d_1031_v68_, default__.abs(990))).set(d_1031_v68_, default__.abs((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])))).isdisjoint((d_1034_v69_).set(d_1031_v68_, default__.abs((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])))
                lhs156_ = globalState
                lhs157_ = globalState
                lhs158_ = globalState
                lhs159_ = globalState
                lhs156_.f10 = rhs179_
                lhs157_.f0 = rhs180_
                lhs158_.f8 = rhs181_
                lhs159_.f10 = rhs182_
                d_942_v19_ = (d_942_v19_).set((d_924_v3_).cardinality, d_921_v0_)
            elif True:
                d_1035_v70_: D2
                d_1035_v70_ = D2_DC6(d_923_v2_)
                d_1036_v71_: _dafny.Map
                d_1036_v71_ = _dafny.Map({(d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]: _dafny.MultiSet([(d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], 928])})
                d_1037_v72_: _dafny.Map
                d_1037_v72_ = _dafny.Map({(d_936_v13_).set((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], default__.abs((self).fm0(((d_1036_v71_)[d_921_v0_] if (d_921_v0_) in (d_1036_v71_) else _dafny.MultiSet(d_947_v22_)), d_923_v2_, _dafny.SeqWithoutIsStrInference([d_927_v6_ for d_1038_i6_ in range(default__.abs(818))]), (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], globalState))): True})
                pat_let_tv48_ = d_947_v22_
                def iife81_(_pat_let16_0):
                    def iife82_(d_1039_dt__update__tmp_h1_):
                        def iife83_(_pat_let17_0):
                            def iife84_(d_1040_dt__update_hcf9_h0_):
                                return D2_DC6(d_1040_dt__update_hcf9_h0_)
                            return iife84_(_pat_let17_0)
                        return iife83_(len(pat_let_tv48_))
                    return iife82_(_pat_let16_0)
                rhs183_ = 130
                rhs184_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                rhs185_ = len(d_1037_v72_)
                rhs186_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
                rhs187_ = iife81_(default__.fm45((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], globalState))
                lhs160_ = globalState
                lhs161_ = globalState
                lhs162_ = globalState
                lhs163_ = globalState
                lhs160_.f11 = rhs183_
                lhs161_.f16 = rhs184_
                lhs162_.f11 = rhs185_
                lhs163_.f6 = rhs186_
                d_1035_v70_ = rhs187_
                d_936_v13_ = (_dafny.MultiSet([True])).intersection(d_936_v13_)
                d_1041_v73_: C1
                nw138_ = C1()
                nw138_.ctor__((len(d_995_v48_)) + (len(d_926_v5_)), (self).f22)
                d_1041_v73_ = nw138_
                d_1042_v74_: C0
                nw139_ = C0()
                nw139_.ctor__()
                d_1042_v74_ = nw139_
                d_1043_v75_: C0
                nw140_ = C0()
                nw140_.ctor__()
                d_1043_v75_ = nw140_
            d_1044_v76_: D5
            d_1044_v76_ = D5_DC17(_dafny.Map({d_926_v5_: (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]}))
            source16_ = d_1044_v76_
            if source16_.is_DC18:
                d_1045___mcc_h16_ = source16_.cf30
                d_1046___mcc_h17_ = source16_.cf31
                d_1047_cf31_ = d_1046___mcc_h17_
                d_1048_cf30_ = d_1045___mcc_h16_
                (globalState).f8 = d_1048_cf30_
                d_1049_v77_: C2
                nw141_ = C2()
                nw141_.ctor__(d_923_v2_, (self).f22)
                d_1049_v77_ = nw141_
                (globalState).f16 = d_923_v2_
                d_1050_v78_: _dafny.Map
                d_1050_v78_ = _dafny.Map({d_935_v12_: not (True) or (d_921_v0_)})
                d_1050_v78_ = (d_1050_v78_).set(d_935_v12_, ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) < (417))
            elif source16_.is_DC19:
                d_1051___mcc_h18_ = source16_.cf32
                d_1052___mcc_h19_ = source16_.cf33
                d_1053_cf33_ = d_1052___mcc_h19_
                d_1054_cf32_ = d_1051___mcc_h18_
                d_922_v1_ = d_922_v1_
                d_1055_v79_: C0
                nw142_ = C0()
                nw142_.ctor__()
                d_1055_v79_ = nw142_
                (globalState).f10 = d_1053_cf33_
                d_923_v2_ = (len((d_926_v5_) + (d_926_v5_))) - ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
            elif source16_.is_DC17:
                d_1056___mcc_h20_ = source16_.cf29
                d_1057_cf29_ = d_1056___mcc_h20_
                (globalState).f0 = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                d_921_v0_ = False
                (globalState).f0 = d_921_v0_
                d_1058_v80_: _dafny.Set
                d_1058_v80_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbtvyab"))})
                index166_ = default__.safeIndex(85, (d_931_v10_).length(0))
                (d_931_v10_)[index166_] = (d_1058_v80_).ispropersubset(d_1058_v80_)
            elif True:
                d_1059___mcc_h21_ = source16_.cf34
                d_1060_cf34_ = d_1059___mcc_h21_
                (globalState).f11 = (0) - ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                d_1061_v81_: D10
                d_1061_v81_ = D10_DC34(d_939_v16_, d_921_v0_, d_921_v0_, False, d_921_v0_)
                d_1062_v82_: _dafny.Map
                d_1062_v82_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1063_i7_ in range(default__.abs(164))]): (d_1061_v81_).cf56})
                (globalState).f6 = default__.safeModuloInt(len(d_1062_v82_), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
                d_926_v5_ = ((_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1064_i8_ in range(default__.abs(878))])).set(default__.safeIndex(((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) + ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1065_i8_ in range(default__.abs(878))]))), d_939_v16_)).set(default__.safeIndex((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], len((_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1066_i8_ in range(default__.abs(878))])).set(default__.safeIndex(((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) + ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_939_v16_ for d_1067_i8_ in range(default__.abs(878))]))), d_939_v16_))), d_939_v16_)
                index167_ = default__.safeIndex(85, (d_931_v10_).length(0))
                rhs188_ = d_935_v12_
                rhs189_ = not (not((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))])) or (d_921_v0_)
                lhs164_ = d_931_v10_
                lhs165_ = default__.safeIndex(85, (d_931_v10_).length(0))
                d_935_v12_ = rhs188_
                lhs164_[lhs165_] = rhs189_
            d_1068_v83_: _dafny.Map
            d_1068_v83_ = _dafny.Map({len(d_929_v8_): d_928_v7_})
            (globalState).f8 = len(((d_1068_v83_).set(len(d_926_v5_), d_928_v7_)) | (d_1068_v83_))
            d_1069_v84_: D6
            d_1069_v84_ = D6_DC23(not((d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))])
            d_1069_v84_ = d_1069_v84_
            d_1070_v85_: _dafny.Seq
            d_1070_v85_ = _dafny.SeqWithoutIsStrInference([d_935_v12_, d_935_v12_])
            d_1071_v86_: D8
            d_1071_v86_ = D8_DC28((d_1070_v85_) + (_dafny.SeqWithoutIsStrInference([d_935_v12_])))
            index168_ = default__.safeIndex(67, (d_935_v12_).length(0))
            rhs190_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
            rhs191_ = (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]
            rhs192_ = d_1071_v86_
            rhs193_ = ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) * ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))])
            lhs166_ = globalState
            lhs167_ = globalState
            lhs168_ = d_935_v12_
            lhs169_ = default__.safeIndex(67, (d_935_v12_).length(0))
            lhs166_.f16 = rhs190_
            lhs167_.f16 = rhs191_
            d_1071_v86_ = rhs192_
            lhs168_[lhs169_] = rhs193_
        d_1072_i9_: int
        d_1072_i9_ = 0
        with _dafny.label("6"):
            while default__.fm37(d_929_v8_, d_942_v19_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1076_i10_ in range(default__.abs(964))])).set(default__.safeIndex(168, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1077_i10_ in range(default__.abs(964))]))), d_939_v16_), globalState):
                with _dafny.c_label("6"):
                    if (d_1072_i9_) >= (100):
                        raise _dafny.Break("6")
                    d_1072_i9_ = (d_1072_i9_) + (1)
                    d_1073_v87_: _dafny.Array
                    nw143_ = _dafny.Array(None, 2)
                    nw143_[int(0)] = (d_926_v5_) + (d_926_v5_)
                    nw143_[int(1)] = d_926_v5_
                    d_1073_v87_ = nw143_
                    d_1073_v87_ = d_1073_v87_
                    (globalState).f0 = True
                    (globalState).f10 = (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]
                    d_1074_v88_: _dafny.Map
                    d_1074_v88_ = _dafny.Map({d_921_v0_: (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))]})
                    d_1075_v89_: D12
                    d_1075_v89_ = D12_DC42(d_923_v2_)
                    d_1074_v88_ = _dafny.Map({d_921_v0_: ((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) > ((self).fm0(_dafny.MultiSet([(d_1075_v89_).cf69]), (d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))], d_930_v9_, (d_931_v10_)[default__.safeIndex(85, (d_931_v10_).length(0))], globalState))})
                    pass
            pass
        r0 = d_931_v10_
        r1 = (((d_935_v12_)[default__.safeIndex(67, (d_935_v12_).length(0))]) - (65)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlauvedc"))))
        return r0, r1


class C4(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self.f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f28, f22):
        (self).f28 = f28
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return 955

    def fm13(self, p0, p1, p2, globalState):
        return False

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_1078_v0_: _dafny.Seq
        d_1078_v0_ = _dafny.SeqWithoutIsStrInference([not(p0)])
        d_1079_v1_: int
        d_1079_v1_ = 508
        d_1080_v2_: _dafny.MultiSet
        d_1080_v2_ = _dafny.MultiSet([p0])
        (globalState).f5 = ((d_1078_v0_) + ((d_1078_v0_).set(default__.safeIndex(d_1079_v1_, len(d_1078_v0_)), (self).fm13(_dafny.SeqWithoutIsStrInference([d_1079_v1_ for d_1081_i0_ in range(default__.abs(170))]), p0, (d_1080_v2_).cardinality, globalState)))) + ((d_1078_v0_ if self.f28 else _dafny.SeqWithoutIsStrInference([self.f28])))
        d_1082_v3_: _dafny.Array
        def lambda38_(d_1083_i2_):
            return not(False)

        init19_ = lambda38_
        nw144_ = _dafny.Array(None, 21)
        for i0_19_ in range(nw144_.length(0)):
            nw144_[i0_19_] = init19_(i0_19_)
        d_1082_v3_ = nw144_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1082_v3_).length(0)):
            d_1084_i1_: int = guard_loop_1_
            if (True) and (((0) <= (d_1084_i1_)) and ((d_1084_i1_) < ((d_1082_v3_).length(0)))):
                (d_1082_v3_)[(d_1084_i1_)] = p0
        d_1080_v2_ = d_1080_v2_
        d_1085_v4_: C1
        nw145_ = C1()
        nw145_.ctor__((len(d_1078_v0_)) - ((0) - (d_1079_v1_)), (self).f22)
        d_1085_v4_ = nw145_
        d_1086_v5_: _dafny.Seq
        d_1086_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqfjidb"))
        d_1086_v5_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1087_i3_ in range(default__.abs(-549))])) + (d_1086_v5_)
        d_1082_v3_ = (d_1082_v3_ if p0 else d_1082_v3_)
        r0 = self.f28
        r1 = d_1082_v3_
        d_1088_v6_: _dafny.Seq
        d_1088_v6_ = _dafny.SeqWithoutIsStrInference([d_1079_v1_, d_1079_v1_, (d_1085_v4_).f30])
        d_1089_v7_: _dafny.Seq
        d_1089_v7_ = _dafny.SeqWithoutIsStrInference([(d_1085_v4_).f30, len(d_1088_v6_), (d_1085_v4_).f30, (d_1085_v4_).f30, d_1079_v1_])
        d_1090_v8_: str
        d_1090_v8_ = _dafny.CodePoint('l')
        d_1091_v9_: D10
        d_1091_v9_ = D10_DC34(d_1090_v8_, self.f28, p0, self.f28, p0)
        d_1092_v10_: _dafny.Map
        d_1092_v10_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(d_1085_v4_).f30, d_1079_v1_])) + (d_1088_v6_): (d_1091_v9_).cf55})
        r2 = d_1092_v10_
        r3 = (d_1085_v4_).fm35((d_1085_v4_).f30, (d_1085_v4_).f30, p0, (d_1085_v4_).f30, globalState)
        return r0, r1, r2, r3


class C5:
    def  __init__(self):
        self._f26: _dafny.Map = _dafny.Map({})
        self._f27: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f26, f27):
        (self)._f26 = f26
        (self)._f27 = f27

    def fm4(self, p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([not(False)]))

    def fm5(self, p0, globalState):
        return not (((self).f27) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([511])))) or (((self).f27) != (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f27})), (self).f27, (self).f27, (self).f27, len((self).f26)]))))

    def m2(self, p0, p1, p2, globalState):
        d_1093_v0_: D0
        d_1093_v0_ = D0_DC0()
        source17_ = d_1093_v0_
        if source17_.is_DC0:
            d_1094_v1_: bool
            d_1094_v1_ = True
            if (self).fm5(not(not (True) or (d_1094_v1_)), globalState):
                d_1095_v2_: str
                d_1095_v2_ = _dafny.CodePoint('s')
                d_1096_v3_: _dafny.Seq
                d_1096_v3_ = _dafny.SeqWithoutIsStrInference([d_1094_v1_, d_1094_v1_, d_1094_v1_])
                d_1097_v4_: _dafny.Map
                d_1097_v4_ = _dafny.Map({d_1095_v2_: d_1096_v3_})
                rhs194_ = ((d_1096_v3_)[default__.safeIndex(p0, len(d_1096_v3_))]) or (d_1094_v1_)
                rhs195_ = (_dafny.Map({d_1095_v2_: d_1096_v3_})) | (d_1097_v4_)
                rhs196_ = p2
                lhs170_ = globalState
                lhs171_ = globalState
                lhs170_.f0 = rhs194_
                d_1097_v4_ = rhs195_
                lhs171_.f8 = rhs196_
                d_1098_v5_: _dafny.Array
                nw146_ = _dafny.Array(int(0), 5)
                d_1098_v5_ = nw146_
                d_1099_v6_: _dafny.Map
                d_1099_v6_ = _dafny.Map({d_1098_v5_: _dafny.CodePoint('q')})
                d_1100_v7_: _dafny.Array
                nw147_ = _dafny.Array(None, 17)
                nw147_[int(0)] = d_1099_v6_
                nw147_[int(1)] = d_1099_v6_
                nw147_[int(2)] = _dafny.Map({d_1098_v5_: d_1095_v2_})
                nw147_[int(3)] = d_1099_v6_
                nw147_[int(4)] = (d_1099_v6_) | (_dafny.Map({d_1098_v5_: d_1095_v2_}))
                nw147_[int(5)] = (d_1099_v6_) | (d_1099_v6_)
                nw147_[int(6)] = (d_1099_v6_) | (d_1099_v6_)
                nw147_[int(7)] = d_1099_v6_
                nw147_[int(8)] = d_1099_v6_
                nw147_[int(9)] = d_1099_v6_
                nw147_[int(10)] = d_1099_v6_
                nw147_[int(11)] = (d_1099_v6_).set(d_1098_v5_, d_1095_v2_)
                nw147_[int(12)] = d_1099_v6_
                nw147_[int(13)] = _dafny.Map({d_1098_v5_: d_1095_v2_})
                nw147_[int(14)] = (d_1099_v6_) | (_dafny.Map({d_1098_v5_: d_1095_v2_}))
                nw147_[int(15)] = (d_1099_v6_) | (_dafny.Map({d_1098_v5_: d_1095_v2_}))
                nw147_[int(16)] = d_1099_v6_
                d_1100_v7_ = nw147_
                index169_ = default__.safeIndex(141, (d_1100_v7_).length(0))
                (d_1100_v7_)[index169_] = d_1099_v6_
                d_1101_v8_: _dafny.Map
                d_1101_v8_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 636})
                d_1102_v9_: _dafny.Map
                d_1102_v9_ = _dafny.Map({d_1094_v1_: d_1101_v8_})
                d_1103_v10_: _dafny.Map
                d_1103_v10_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([988])})
                d_1104_v11_: D0
                d_1104_v11_ = D0_DC1((0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_1095_v2_ for d_1105_i0_ in range(default__.abs(496))])), p2)), (d_1094_v1_ if d_1094_v1_ else d_1094_v1_), (d_1102_v9_) | (d_1102_v9_), len(d_1103_v10_), d_1095_v2_)
                d_1106_v12_: _dafny.Array
                nw148_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_1106_v12_ = nw148_
                d_1107_v13_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Array(None, 0), 7)
                d_1107_v13_ = nw149_
                index170_ = default__.safeIndex(814, (d_1106_v12_).length(0))
                (d_1106_v12_)[index170_] = d_1107_v13_
                pat_let_tv49_ = p2
                pat_let_tv50_ = d_1094_v1_
                pat_let_tv51_ = p0
                pat_let_tv52_ = d_1094_v1_
                pat_let_tv53_ = d_1101_v8_
                index171_ = default__.safeIndex(141, (d_1100_v7_).length(0))
                index172_ = default__.safeIndex(814, (d_1106_v12_).length(0))
                def iife85_(_pat_let18_0):
                    def iife86_(d_1108_dt__update__tmp_h0_):
                        def iife87_(_pat_let19_0):
                            def iife88_(d_1109_dt__update_hcf3_h0_):
                                def iife89_(_pat_let20_0):
                                    def iife90_(d_1110_dt__update_hcf0_h0_):
                                        def iife91_(_pat_let21_0):
                                            def iife92_(d_1111_dt__update_hcf2_h0_):
                                                return D0_DC1(d_1110_dt__update_hcf0_h0_, (d_1108_dt__update__tmp_h0_).cf1, d_1111_dt__update_hcf2_h0_, d_1109_dt__update_hcf3_h0_, (d_1108_dt__update__tmp_h0_).cf4)
                                            return iife92_(_pat_let21_0)
                                        return iife91_(_dafny.Map({pat_let_tv52_: pat_let_tv53_}))
                                    return iife90_(_pat_let20_0)
                                return iife89_(pat_let_tv51_)
                            return iife88_(_pat_let19_0)
                        return iife87_((pat_let_tv49_ if pat_let_tv50_ else -302))
                    return iife86_(_pat_let18_0)
                rhs197_ = (default__.safeDivisionInt(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))) != (p2)
                rhs198_ = d_1099_v6_
                rhs199_ = iife85_(d_1104_v11_)
                rhs200_ = d_1107_v13_
                lhs172_ = globalState
                lhs173_ = d_1100_v7_
                lhs174_ = default__.safeIndex(141, (d_1100_v7_).length(0))
                lhs175_ = d_1106_v12_
                lhs176_ = default__.safeIndex(814, (d_1106_v12_).length(0))
                lhs172_.f10 = rhs197_
                lhs173_[lhs174_] = rhs198_
                d_1104_v11_ = rhs199_
                lhs175_[lhs176_] = rhs200_
                d_1112_v14_: _dafny.Seq
                d_1112_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hcvvsgggf"))
                d_1113_v15_: _dafny.MultiSet
                d_1113_v15_ = _dafny.MultiSet([(self).f27, p0])
                rhs201_ = d_1112_v14_
                rhs202_ = (default__.safeModuloInt((d_1113_v15_).cardinality, p0)) * (p2)
                rhs203_ = d_1094_v1_
                lhs177_ = globalState
                lhs178_ = globalState
                d_1112_v14_ = rhs201_
                lhs177_.f8 = rhs202_
                lhs178_.f0 = rhs203_
                d_1114_v16_: int
                d_1115_v17_: D0
                out50_: int
                out51_: D0
                out50_, out51_ = (self).m3(_dafny.SeqWithoutIsStrInference([d_1095_v2_ for d_1116_i1_ in range(default__.abs(882))]), not(not (d_1094_v1_) or (d_1094_v1_)), globalState)
                d_1114_v16_ = out50_
                d_1115_v17_ = out51_
                d_1117_v18_: _dafny.Map
                d_1117_v18_ = _dafny.Map({not (not(d_1094_v1_)) or (True): d_1098_v5_})
                d_1117_v18_ = (d_1117_v18_).set(d_1094_v1_, d_1098_v5_)
            elif True:
                d_1118_v19_: _dafny.Seq
                d_1118_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                d_1118_v19_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tggp")) if False else (p1) + (p1))
                d_1094_v1_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkfl"))) + (d_1118_v19_)) == (default__.fm6(globalState))
                d_1119_v20_: str
                d_1119_v20_ = _dafny.CodePoint('e')
                d_1120_v21_: _dafny.Set
                d_1120_v21_ = _dafny.Set({p1, (d_1118_v19_).set(default__.safeIndex(p2, len(d_1118_v19_)), d_1119_v20_)})
                d_1121_v22_: _dafny.Seq
                d_1121_v22_ = _dafny.SeqWithoutIsStrInference([d_1118_v19_])
                def iife93_():
                    coll49_ = _dafny.Set()
                    compr_49_: _dafny.Seq
                    for compr_49_ in (d_1121_v22_).Elements:
                        d_1122_v23_: _dafny.Seq = compr_49_
                        if (d_1122_v23_) in (d_1121_v22_):
                            coll49_ = coll49_.union(_dafny.Set([d_1122_v23_]))
                    return _dafny.Set(coll49_)
                rhs204_ = (iife93_()
                ) | (d_1120_v21_)
                rhs205_ = ((self).f27) * (893)
                rhs206_ = d_1094_v1_
                rhs207_ = ((self).f27) - ((0) - (p0))
                lhs179_ = globalState
                lhs180_ = globalState
                d_1120_v21_ = rhs204_
                lhs179_.f11 = rhs205_
                d_1094_v1_ = rhs206_
                lhs180_.f11 = rhs207_
                d_1123_v24_: _dafny.Seq
                d_1123_v24_ = _dafny.SeqWithoutIsStrInference([d_1094_v1_, True])
                d_1124_v25_: _dafny.MultiSet
                d_1124_v25_ = _dafny.MultiSet([d_1094_v1_, d_1094_v1_])
                d_1125_v26_: _dafny.Map
                d_1125_v26_ = _dafny.Map({((d_1124_v25_)[d_1094_v1_] if (d_1094_v1_) in (d_1124_v25_) else len(d_1118_v19_)): p0})
                d_1126_v27_: _dafny.Map
                d_1126_v27_ = _dafny.Map({not((d_1123_v24_)[default__.safeIndex(p0, len(d_1123_v24_))]): d_1125_v26_})
                d_1127_v28_: _dafny.Map
                d_1127_v28_ = _dafny.Map({d_1118_v19_: d_1094_v1_})
                d_1128_v29_: D0
                d_1128_v29_ = D0_DC1(len((self).f26), False, d_1126_v27_, len(d_1127_v28_), d_1119_v20_)
                (globalState).f0 = (not((self).fm5(d_1094_v1_, globalState)) if (d_1128_v29_).cf1 else d_1094_v1_)
                d_1094_v1_ = not((d_1124_v25_).issubset(d_1124_v25_))
            if d_1094_v1_:
                d_1093_v0_ = d_1093_v0_
                d_1129_v30_: _dafny.Array
                nw150_ = _dafny.Array(_dafny.Array(None, 0), 20)
                d_1129_v30_ = nw150_
                d_1130_v31_: _dafny.Array
                def lambda39_(d_1131_v1_):
                    def lambda40_(d_1132_i2_):
                        return d_1131_v1_

                    return lambda40_

                init20_ = lambda39_(d_1094_v1_)
                nw151_ = _dafny.Array(None, 7)
                for i0_20_ in range(nw151_.length(0)):
                    nw151_[i0_20_] = init20_(i0_20_)
                d_1130_v31_ = nw151_
                index173_ = default__.safeIndex(866, (d_1129_v30_).length(0))
                (d_1129_v30_)[index173_] = d_1130_v31_
                index174_ = default__.safeIndex(866, (d_1129_v30_).length(0))
                nw152_ = _dafny.Array(None, 7)
                nw152_[int(0)] = d_1094_v1_
                nw152_[int(1)] = False
                nw152_[int(2)] = (default__.fm7(p0, globalState)) == ((self).f27)
                nw152_[int(3)] = d_1094_v1_
                nw152_[int(4)] = d_1094_v1_
                nw152_[int(5)] = d_1094_v1_
                nw152_[int(6)] = d_1094_v1_
                (d_1129_v30_)[index174_] = nw152_
                d_1133_v32_: _dafny.Seq
                d_1133_v32_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26, (self).f26])
                d_1134_v35_: _dafny.Seq
                d_1134_v35_ = _dafny.SeqWithoutIsStrInference([738])
                d_1135_v36_: _dafny.Array
                nw153_ = _dafny.Array(None, 20)
                nw153_[int(0)] = (self).f26
                nw153_[int(1)] = (self).f26
                nw153_[int(2)] = (d_1133_v32_)[default__.safeIndex(p2, len(d_1133_v32_))]
                nw153_[int(3)] = (self).f26
                nw153_[int(4)] = (((self).f26).set(p0, False)) | ((self).f26)
                nw153_[int(5)] = (self).f26
                nw153_[int(6)] = (self).f26
                nw153_[int(7)] = ((self).f26) | (_dafny.Map({(_dafny.MultiSet([d_1094_v1_, not(d_1094_v1_)])).cardinality: d_1094_v1_}))
                nw153_[int(8)] = _dafny.Map({(self).f27: d_1094_v1_})
                nw153_[int(9)] = (self).f26
                nw153_[int(10)] = (self).f26
                nw153_[int(11)] = ((d_1133_v32_)[default__.safeIndex(p0, len(d_1133_v32_))] if d_1094_v1_ else _dafny.Map({(self).f27: d_1094_v1_}))
                nw153_[int(12)] = ((self).f26) | ((self).f26)
                nw153_[int(13)] = (self).f26
                nw153_[int(14)] = (self).f26
                def iife94_():
                    coll50_ = _dafny.Map()
                    compr_50_: int
                    for compr_50_ in _dafny.IntegerRange(243, 987):
                        d_1136_v33_: int = compr_50_
                        if ((243) <= (d_1136_v33_)) and ((d_1136_v33_) < (987)):
                            coll50_[(d_1136_v33_) - (-492)] = d_1094_v1_
                    return _dafny.Map(coll50_)
                nw153_[int(15)] = iife94_()
                
                nw153_[int(16)] = (self).f26
                def iife95_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in (d_1134_v35_).Elements:
                        d_1137_v34_: int = compr_51_
                        if (d_1137_v34_) in (d_1134_v35_):
                            coll51_[(d_1137_v34_) + ((self).f27)] = d_1094_v1_
                    return _dafny.Map(coll51_)
                nw153_[int(17)] = (iife95_()
                ) | (default__.fm8(p1, d_1094_v1_, p2, globalState))
                nw153_[int(18)] = _dafny.Map({653: d_1094_v1_})
                nw153_[int(19)] = (self).f26
                d_1135_v36_ = nw153_
                d_1135_v36_ = d_1135_v36_
                d_1138_v37_: _dafny.Set
                d_1138_v37_ = _dafny.Set({d_1094_v1_})
                d_1139_v38_: _dafny.MultiSet
                d_1139_v38_ = _dafny.MultiSet([d_1138_v37_])
                (globalState).f0 = ((d_1134_v35_) == (d_1134_v35_)) == ((d_1139_v38_).ispropersubset(d_1139_v38_))
                d_1140_v39_: _dafny.Array
                nw154_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_1140_v39_ = nw154_
                d_1141_v40_: _dafny.MultiSet
                d_1141_v40_ = _dafny.MultiSet([len(p1)])
                d_1142_v41_: _dafny.Array
                nw155_ = _dafny.Array(None, 8)
                nw155_[int(0)] = 740
                nw155_[int(1)] = p0
                nw155_[int(2)] = p2
                nw155_[int(3)] = (self).f27
                nw155_[int(4)] = p0
                nw155_[int(5)] = 23
                nw155_[int(6)] = len(p1)
                nw155_[int(7)] = (d_1141_v40_).cardinality
                d_1142_v41_ = nw155_
                index175_ = default__.safeIndex(853, (d_1140_v39_).length(0))
                (d_1140_v39_)[index175_] = d_1142_v41_
                index176_ = default__.safeIndex(853, (d_1140_v39_).length(0))
                (d_1140_v39_)[index176_] = d_1142_v41_
            elif True:
                d_1143_v42_: _dafny.Array
                def lambda41_(d_1144_p2_):
                    def lambda42_(d_1145_i3_):
                        return (d_1145_i3_) + (d_1144_p2_)

                    return lambda42_

                init21_ = lambda41_(p2)
                nw156_ = _dafny.Array(None, 1)
                for i0_21_ in range(nw156_.length(0)):
                    nw156_[i0_21_] = init21_(i0_21_)
                d_1143_v42_ = nw156_
                d_1143_v42_ = d_1143_v42_
                d_1146_v43_: _dafny.Set
                d_1146_v43_ = _dafny.Set({d_1094_v1_})
                d_1147_v44_: _dafny.Map
                d_1147_v44_ = _dafny.Map({p2: -518})
                d_1148_v45_: str
                d_1148_v45_ = _dafny.CodePoint('l')
                d_1149_v46_: D0
                d_1149_v46_ = D0_DC1((len(d_1146_v43_)) + ((self).f27), d_1094_v1_, _dafny.Map({not(d_1094_v1_): d_1147_v44_}), (self).f27, (d_1148_v45_ if d_1094_v1_ else d_1148_v45_))
                d_1150_v47_: _dafny.Seq
                d_1150_v47_ = _dafny.SeqWithoutIsStrInference([default__.fm7((self).f27, globalState), -995, len(p1)])
                d_1151_v48_: _dafny.MultiSet
                d_1151_v48_ = _dafny.MultiSet([p0, (_dafny.MultiSet([(self).f26])).cardinality, p2, len(d_1150_v47_), (self).f27])
                rhs208_ = d_1149_v46_
                rhs209_ = (_dafny.MultiSet([p0])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p2])))
                rhs210_ = (p2) + (p2)
                lhs181_ = globalState
                d_1149_v46_ = rhs208_
                d_1151_v48_ = rhs209_
                lhs181_.f8 = rhs210_
                d_1148_v45_ = d_1148_v45_
                (globalState).f0 = d_1094_v1_
                d_1152_v49_: _dafny.Seq
                d_1152_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lx"))
                d_1153_v50_: D1
                d_1153_v50_ = D1_DC3(p1)
                d_1152_v49_ = ((p1) + ((d_1153_v50_).cf6)) + (p1)
            d_1154_v51_: _dafny.Map
            d_1154_v51_ = _dafny.Map({d_1094_v1_: _dafny.Map({p2: p0})})
            d_1155_v52_: _dafny.Map
            d_1155_v52_ = _dafny.Map({d_1094_v1_: d_1094_v1_})
            d_1156_v53_: str
            d_1156_v53_ = _dafny.CodePoint('k')
            d_1157_v54_: D0
            d_1157_v54_ = D0_DC1(((self).f27) + ((self).f27), d_1094_v1_, (d_1154_v51_) | (d_1154_v51_), default__.fm7(len(d_1155_v52_), globalState), d_1156_v53_)
            source18_ = d_1157_v54_
            if source18_.is_DC0:
                (globalState).f0 = (d_1156_v53_) not in ((_dafny.SeqWithoutIsStrInference([d_1156_v53_ for d_1158_i4_ in range(default__.abs(-337))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qewbkgcot"))))
                (globalState).f11 = default__.fm7(93, globalState)
                d_1159_v55_: _dafny.Seq
                d_1159_v55_ = _dafny.SeqWithoutIsStrInference([d_1094_v1_])
                d_1160_v56_: _dafny.Seq
                d_1160_v56_ = _dafny.SeqWithoutIsStrInference([len(p1), p2])
                d_1161_v57_: _dafny.Seq
                d_1161_v57_ = _dafny.SeqWithoutIsStrInference([d_1159_v55_, (d_1159_v55_) + ((self).fm4(d_1160_v56_, p1, _dafny.Set({not(d_1094_v1_)}), p2, globalState))])
                d_1161_v57_ = d_1161_v57_
                d_1162_v58_: _dafny.Array
                def lambda43_(d_1163_v1_):
                    def lambda44_(d_1164_i5_):
                        return (d_1163_v1_ if d_1163_v1_ else d_1163_v1_)

                    return lambda44_

                init22_ = lambda43_(d_1094_v1_)
                nw157_ = _dafny.Array(None, 18)
                for i0_22_ in range(nw157_.length(0)):
                    nw157_[i0_22_] = init22_(i0_22_)
                d_1162_v58_ = nw157_
                d_1162_v58_ = d_1162_v58_
            elif True:
                d_1165___mcc_h5_ = source18_.cf0
                d_1166___mcc_h6_ = source18_.cf1
                d_1167___mcc_h7_ = source18_.cf2
                d_1168___mcc_h8_ = source18_.cf3
                d_1169___mcc_h9_ = source18_.cf4
                d_1170_cf4_ = d_1169___mcc_h9_
                d_1171_cf3_ = d_1168___mcc_h8_
                d_1172_cf2_ = d_1167___mcc_h7_
                d_1173_cf1_ = d_1166___mcc_h6_
                d_1174_cf0_ = d_1165___mcc_h5_
                d_1175_v59_: int
                d_1176_v60_: D0
                out52_: int
                out53_: D0
                out52_, out53_ = (self).m3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bubrnhav")), d_1094_v1_, globalState)
                d_1175_v59_ = out52_
                d_1176_v60_ = out53_
                d_1177_v61_: _dafny.Array
                def lambda45_(d_1178_cf0_):
                    def lambda46_(d_1179_i6_):
                        return (d_1179_i6_) - (d_1178_cf0_)

                    return lambda46_

                init23_ = lambda45_(d_1174_cf0_)
                nw158_ = _dafny.Array(None, 21)
                for i0_23_ in range(nw158_.length(0)):
                    nw158_[i0_23_] = init23_(i0_23_)
                d_1177_v61_ = nw158_
                d_1177_v61_ = d_1177_v61_
                d_1180_v62_: _dafny.Seq
                d_1180_v62_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kgfnnee"))
                d_1180_v62_ = (p1).set(default__.safeIndex((_dafny.MultiSet(default__.fm9(globalState))).cardinality, len(p1)), d_1170_cf4_)
                d_1181_v63_: _dafny.Map
                d_1181_v63_ = _dafny.Map({d_1173_cf1_: d_1170_cf4_})
                d_1182_v64_: _dafny.Map
                d_1182_v64_ = _dafny.Map({(self).f27: d_1181_v63_})
                d_1183_v65_: _dafny.Array
                nw159_ = _dafny.Array(None, 8)
                nw159_[int(0)] = (d_1181_v63_) | (d_1181_v63_)
                nw159_[int(1)] = ((d_1182_v64_)[d_1174_cf0_] if (d_1174_cf0_) in (d_1182_v64_) else d_1181_v63_)
                nw159_[int(2)] = d_1181_v63_
                nw159_[int(3)] = _dafny.Map({d_1173_cf1_: d_1156_v53_})
                nw159_[int(4)] = (d_1181_v63_).set(d_1173_cf1_, d_1156_v53_)
                nw159_[int(5)] = d_1181_v63_
                nw159_[int(6)] = (d_1181_v63_) | (d_1181_v63_)
                nw159_[int(7)] = (d_1181_v63_) | (_dafny.Map({d_1173_cf1_: d_1170_cf4_}))
                d_1183_v65_ = nw159_
                index177_ = default__.safeIndex(461, (d_1183_v65_).length(0))
                (d_1183_v65_)[index177_] = (d_1181_v63_).set(d_1094_v1_, _dafny.CodePoint('e'))
                d_1184_v66_: _dafny.Array
                nw160_ = _dafny.Array(_dafny.Seq({}), 28)
                d_1184_v66_ = nw160_
                d_1185_v67_: _dafny.Seq
                d_1185_v67_ = _dafny.SeqWithoutIsStrInference([d_1094_v1_])
                index178_ = default__.safeIndex(73, (d_1184_v66_).length(0))
                (d_1184_v66_)[index178_] = d_1185_v67_
                d_1186_v68_: _dafny.Seq
                d_1186_v68_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_1187_v69_: _dafny.MultiSet
                d_1187_v69_ = _dafny.MultiSet([len(d_1186_v68_), d_1174_cf0_])
                index179_ = default__.safeIndex(461, (d_1183_v65_).length(0))
                index180_ = default__.safeIndex(73, (d_1184_v66_).length(0))
                rhs211_ = d_1094_v1_
                rhs212_ = d_1094_v1_
                rhs213_ = (d_1181_v63_) | (d_1181_v63_)
                rhs214_ = (_dafny.SeqWithoutIsStrInference([d_1094_v1_])) + (_dafny.SeqWithoutIsStrInference([(D0_DC1(((d_1187_v69_)[d_1171_cf3_] if (d_1171_cf3_) in (d_1187_v69_) else len(default__.fm10(len(d_1180_v62_), d_1094_v1_, d_1171_cf3_, globalState))), False, d_1172_cf2_, default__.fm7(p0, globalState), d_1156_v53_)).cf1, d_1173_cf1_, d_1094_v1_]))
                lhs182_ = globalState
                lhs183_ = d_1183_v65_
                lhs184_ = default__.safeIndex(461, (d_1183_v65_).length(0))
                lhs185_ = d_1184_v66_
                lhs186_ = default__.safeIndex(73, (d_1184_v66_).length(0))
                lhs182_.f10 = rhs211_
                d_1173_cf1_ = rhs212_
                lhs183_[lhs184_] = rhs213_
                lhs185_[lhs186_] = rhs214_
            d_1156_v53_ = d_1156_v53_
        elif True:
            d_1188___mcc_h0_ = source17_.cf0
            d_1189___mcc_h1_ = source17_.cf1
            d_1190___mcc_h2_ = source17_.cf2
            d_1191___mcc_h3_ = source17_.cf3
            d_1192___mcc_h4_ = source17_.cf4
            d_1193_cf4_ = d_1192___mcc_h4_
            d_1194_cf3_ = d_1191___mcc_h3_
            d_1195_cf2_ = d_1190___mcc_h2_
            d_1196_cf1_ = d_1189___mcc_h1_
            d_1197_cf0_ = d_1188___mcc_h0_
            d_1198_v70_: _dafny.Map
            d_1198_v70_ = _dafny.Map({p0: len(default__.fm11(d_1197_cf0_, d_1193_cf4_, d_1196_cf1_, globalState))})
            d_1199_v71_: _dafny.Seq
            d_1199_v71_ = _dafny.SeqWithoutIsStrInference([d_1197_cf0_, -250, ((d_1198_v70_)[(self).f27] if ((self).f27) in (d_1198_v70_) else (self).f27), (0) - ((self).f27), p2])
            d_1200_v72_: _dafny.Map
            d_1200_v72_ = _dafny.Map({d_1196_cf1_: d_1197_cf0_})
            d_1201_v73_: _dafny.Set
            d_1201_v73_ = _dafny.Set({(self).f27, ((d_1200_v72_)[d_1196_cf1_] if (d_1196_cf1_) in (d_1200_v72_) else d_1194_cf3_)})
            d_1202_v74_: _dafny.Map
            d_1202_v74_ = _dafny.Map({d_1196_cf1_: d_1201_v73_})
            d_1203_v75_: D0
            d_1203_v75_ = D0_DC1(len(d_1199_v71_), d_1196_cf1_, _dafny.Map({d_1196_cf1_: d_1198_v70_}), p2, d_1193_cf4_)
            d_1204_v76_: _dafny.Map
            d_1204_v76_ = _dafny.Map({d_1201_v73_: d_1197_cf0_})
            d_1205_v77_: _dafny.Array
            nw161_ = _dafny.Array(None, 25)
            nw161_[int(0)] = len(d_1199_v71_)
            nw161_[int(1)] = (0) - (len(d_1202_v74_))
            nw161_[int(2)] = -82
            nw161_[int(3)] = d_1194_cf3_
            nw161_[int(4)] = 191
            nw161_[int(5)] = d_1194_cf3_
            nw161_[int(6)] = d_1194_cf3_
            nw161_[int(7)] = p2
            nw161_[int(8)] = p0
            nw161_[int(9)] = p2
            nw161_[int(10)] = (d_1199_v71_)[default__.safeIndex(322, len(d_1199_v71_))]
            nw161_[int(11)] = d_1197_cf0_
            nw161_[int(12)] = p0
            nw161_[int(13)] = p0
            nw161_[int(14)] = d_1197_cf0_
            nw161_[int(15)] = (self).f27
            nw161_[int(16)] = (self).f27
            nw161_[int(17)] = 65
            nw161_[int(18)] = d_1194_cf3_
            nw161_[int(19)] = 750
            nw161_[int(20)] = (d_1203_v75_).cf3
            nw161_[int(21)] = (self).f27
            nw161_[int(22)] = ((d_1204_v76_)[d_1201_v73_] if (d_1201_v73_) in (d_1204_v76_) else p2)
            nw161_[int(23)] = len(d_1200_v72_)
            nw161_[int(24)] = p0
            d_1205_v77_ = nw161_
            d_1206_v78_: D2
            d_1206_v78_ = D2_DC5(d_1205_v77_)
            d_1207_v79_: _dafny.Array
            nw162_ = _dafny.Array(None, 14)
            nw162_[int(0)] = d_1205_v77_
            nw162_[int(1)] = d_1205_v77_
            nw162_[int(2)] = d_1205_v77_
            nw162_[int(3)] = (d_1205_v77_ if d_1196_cf1_ else d_1205_v77_)
            nw162_[int(4)] = d_1205_v77_
            nw162_[int(5)] = d_1205_v77_
            nw162_[int(6)] = d_1205_v77_
            nw162_[int(7)] = d_1205_v77_
            nw162_[int(8)] = d_1205_v77_
            nw162_[int(9)] = d_1205_v77_
            nw162_[int(10)] = (d_1205_v77_ if d_1196_cf1_ else (d_1206_v78_).cf8)
            nw162_[int(11)] = d_1205_v77_
            nw162_[int(12)] = d_1205_v77_
            nw162_[int(13)] = d_1205_v77_
            d_1207_v79_ = nw162_
            index181_ = default__.safeIndex(273, (d_1207_v79_).length(0))
            (d_1207_v79_)[index181_] = d_1205_v77_
            index182_ = default__.safeIndex(273, (d_1207_v79_).length(0))
            rhs215_ = d_1205_v77_
            rhs216_ = (len(default__.fm6(globalState))) + ((0) - (d_1194_cf3_))
            rhs217_ = d_1193_cf4_
            rhs218_ = (self).f27
            lhs187_ = d_1207_v79_
            lhs188_ = default__.safeIndex(273, (d_1207_v79_).length(0))
            lhs189_ = globalState
            lhs190_ = globalState
            lhs187_[lhs188_] = rhs215_
            lhs189_.f16 = rhs216_
            d_1193_cf4_ = rhs217_
            lhs190_.f16 = rhs218_
            d_1208_v80_: _dafny.Array
            def lambda47_(d_1209_v75_):
                def lambda48_(d_1210_i7_):
                    return (d_1209_v75_).cf1

                return lambda48_

            init24_ = lambda47_(d_1203_v75_)
            nw163_ = _dafny.Array(None, 12)
            for i0_24_ in range(nw163_.length(0)):
                nw163_[i0_24_] = init24_(i0_24_)
            d_1208_v80_ = nw163_
            d_1208_v80_ = d_1208_v80_
            (globalState).f10 = False
            d_1196_cf1_ = False
        hi2_ = p0
        for d_1211_i8_ in range((0) - (len(p1)), hi2_):
            d_1212_v81_: _dafny.Seq
            d_1212_v81_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(617, default__.fm7(d_1211_i8_, globalState))])
            d_1213_v82_: _dafny.Map
            d_1213_v82_ = _dafny.Map({(p2) + ((self).f27): (_dafny.SeqWithoutIsStrInference([p0, (0) - ((self).f27), d_1211_i8_, (self).f27])) + (d_1212_v81_)})
            d_1212_v81_ = (((d_1213_v82_)[(0) - (p2)] if ((0) - (p2)) in (d_1213_v82_) else d_1212_v81_)).set(default__.safeIndex(p2, len(((d_1213_v82_)[(0) - (p2)] if ((0) - (p2)) in (d_1213_v82_) else d_1212_v81_))), (p0) - (p2))
            d_1214_v83_: _dafny.Array
            nw164_ = _dafny.Array(_dafny.Map({}), 12)
            d_1214_v83_ = nw164_
            d_1214_v83_ = d_1214_v83_
            d_1215_v84_: bool
            d_1215_v84_ = True
            if ((self).fm5(d_1215_v84_, globalState)) == (d_1215_v84_):
                (globalState).f10 = (d_1215_v84_) or (d_1215_v84_)
                d_1216_v85_: _dafny.Array
                nw165_ = _dafny.Array(int(0), 12)
                d_1216_v85_ = nw165_
                nw166_ = _dafny.Array(int(0), 3)
                d_1216_v85_ = nw166_
                (globalState).f0 = d_1215_v84_
                d_1217_v86_: _dafny.Array
                nw167_ = _dafny.Array(D2.default()(), 16)
                d_1217_v86_ = nw167_
                index183_ = default__.safeIndex(169, (d_1217_v86_).length(0))
                (d_1217_v86_)[index183_] = D2_DC5(d_1216_v85_)
                pat_let_tv54_ = d_1216_v85_
                index184_ = default__.safeIndex(169, (d_1217_v86_).length(0))
                def iife96_(_pat_let22_0):
                    def iife97_(d_1218_dt__update__tmp_h1_):
                        def iife98_(_pat_let23_0):
                            def iife99_(d_1219_dt__update_hcf8_h0_):
                                return D2_DC5(d_1219_dt__update_hcf8_h0_)
                            return iife99_(_pat_let23_0)
                        return iife98_(pat_let_tv54_)
                    return iife97_(_pat_let22_0)
                (d_1217_v86_)[index184_] = iife96_(D2_DC5(d_1216_v85_))
                (globalState).f11 = default__.fm7((self).f27, globalState)
            elif True:
                d_1220_v87_: str
                d_1220_v87_ = _dafny.CodePoint('h')
                d_1220_v87_ = default__.fm12((d_1215_v84_) in (_dafny.Map({not(d_1215_v84_): (d_1212_v81_)[default__.safeIndex(p0, len(d_1212_v81_))]})), 287, globalState)
                d_1221_v88_: int
                d_1222_v89_: D0
                out54_: int
                out55_: D0
                out54_, out55_ = (self).m3(p1, d_1215_v84_, globalState)
                d_1221_v88_ = out54_
                d_1222_v89_ = out55_
                d_1223_v90_: _dafny.Map
                d_1223_v90_ = _dafny.Map({default__.fm6(globalState): len(p1)})
                d_1223_v90_ = (d_1223_v90_).set(p1, p2)
                d_1224_v91_: _dafny.Array
                def lambda49_(d_1225_p1_, d_1226_p2_):
                    def lambda50_(d_1227_i9_):
                        return (len((d_1225_p1_).set(default__.safeIndex(d_1226_p2_, len(d_1225_p1_)), _dafny.CodePoint('u')))) == (d_1226_p2_)

                    return lambda50_

                init25_ = lambda49_(p1, p2)
                nw168_ = _dafny.Array(None, 17)
                for i0_25_ in range(nw168_.length(0)):
                    nw168_[i0_25_] = init25_(i0_25_)
                d_1224_v91_ = nw168_
                index185_ = default__.safeIndex(651, (d_1224_v91_).length(0))
                (d_1224_v91_)[index185_] = d_1215_v84_
                index186_ = default__.safeIndex(481, (d_1224_v91_).length(0))
                (d_1224_v91_)[index186_] = d_1215_v84_
                d_1228_v92_: _dafny.Seq
                d_1228_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
                d_1229_v93_: _dafny.Map
                d_1229_v93_ = _dafny.Map({d_1215_v84_: d_1215_v84_})
                d_1230_v94_: D2
                d_1230_v94_ = D2_DC7(d_1212_v81_, d_1215_v84_, d_1215_v84_, len(d_1229_v93_))
                d_1231_v95_: _dafny.MultiSet
                d_1231_v95_ = _dafny.MultiSet([d_1230_v94_])
                d_1232_v96_: _dafny.Map
                d_1232_v96_ = _dafny.Map({d_1211_i8_: (self).f27})
                d_1233_v97_: D3
                d_1233_v97_ = D3_DC10(d_1232_v96_)
                d_1234_v98_: _dafny.Map
                d_1234_v98_ = _dafny.Map({d_1215_v84_: (d_1233_v97_).cf18})
                d_1235_v99_: D0
                d_1235_v99_ = D0_DC1(p0, d_1215_v84_, d_1234_v98_, p0, d_1220_v87_)
                index187_ = default__.safeIndex(651, (d_1224_v91_).length(0))
                index188_ = default__.safeIndex(481, (d_1224_v91_).length(0))
                rhs219_ = (d_1231_v95_).ispropersubset(d_1231_v95_)
                rhs220_ = (d_1235_v99_).cf1
                rhs221_ = 29
                rhs222_ = (p1)[default__.safeIndex((self).f27, len(p1))]
                rhs223_ = p1
                lhs191_ = d_1224_v91_
                lhs192_ = default__.safeIndex(651, (d_1224_v91_).length(0))
                lhs193_ = d_1224_v91_
                lhs194_ = default__.safeIndex(481, (d_1224_v91_).length(0))
                lhs195_ = globalState
                lhs191_[lhs192_] = rhs219_
                lhs193_[lhs194_] = rhs220_
                lhs195_.f11 = rhs221_
                d_1220_v87_ = rhs222_
                d_1228_v92_ = rhs223_
                d_1236_v100_: _dafny.Map
                d_1236_v100_ = _dafny.Map({d_1220_v87_: d_1221_v88_})
                d_1236_v100_ = (d_1236_v100_).set(default__.fm12((d_1224_v91_)[default__.safeIndex(651, (d_1224_v91_).length(0))], len(d_1228_v92_), globalState), d_1211_i8_)
            d_1237_v101_: _dafny.Map
            d_1237_v101_ = _dafny.Map({d_1211_i8_: p0})
            d_1238_v102_: D3
            d_1238_v102_ = D3_DC10(d_1237_v101_)
            source19_ = d_1238_v102_
            if source19_.is_DC11:
                d_1239___mcc_h10_ = source19_.cf19
                d_1240___mcc_h11_ = source19_.cf20
                d_1241_cf20_ = d_1240___mcc_h11_
                d_1242_cf19_ = d_1239___mcc_h10_
                d_1242_cf19_ = d_1215_v84_
                (globalState).f0 = d_1215_v84_
                (globalState).f0 = not(d_1242_cf19_)
                d_1243_v103_: _dafny.Seq
                d_1243_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkhdpyee"))
                d_1244_v104_: _dafny.Seq
                d_1244_v104_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1245_v105_: _dafny.Map
                d_1245_v105_ = _dafny.Map({d_1242_cf19_: d_1237_v101_})
                d_1246_v106_: D0
                d_1246_v106_ = D0_DC1(d_1211_i8_, d_1242_cf19_, d_1245_v105_, (0) - (p2), _dafny.CodePoint('e'))
                d_1243_v103_ = ((d_1244_v104_)[default__.safeIndex((d_1246_v106_).cf0, len(d_1244_v104_))]) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_1247_i10_ in range(default__.abs(-122))]))
            elif True:
                d_1248___mcc_h12_ = source19_.cf18
                d_1249_cf18_ = d_1248___mcc_h12_
                d_1250_v107_: _dafny.Seq
                d_1250_v107_ = _dafny.SeqWithoutIsStrInference([d_1215_v84_])
                d_1251_v108_: D3
                d_1251_v108_ = D3_DC11(d_1215_v84_, (d_1250_v107_)[default__.safeIndex((self).f27, len(d_1250_v107_))])
                (globalState).f10 = not((d_1251_v108_).cf19)
                d_1252_v109_: str
                d_1252_v109_ = _dafny.CodePoint('h')
                d_1253_v110_: _dafny.Map
                d_1253_v110_ = _dafny.Map({not(False): _dafny.SeqWithoutIsStrInference([d_1252_v109_])})
                d_1254_v111_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 24)
                d_1254_v111_ = nw169_
                d_1255_v112_: _dafny.MultiSet
                d_1255_v112_ = _dafny.MultiSet([d_1254_v111_, d_1254_v111_])
                d_1256_v113_: _dafny.Map
                d_1256_v113_ = _dafny.Map({(0) - (default__.safeDivisionInt(len(((d_1253_v110_)[d_1215_v84_] if (d_1215_v84_) in (d_1253_v110_) else p1)), len(_dafny.Set({((d_1255_v112_)[d_1254_v111_] if (d_1254_v111_) in (d_1255_v112_) else d_1211_i8_)})))): True})
                d_1256_v113_ = (self).f26
                d_1254_v111_ = d_1254_v111_
                d_1257_v114_: _dafny.Set
                d_1257_v114_ = _dafny.Set({d_1252_v109_, d_1252_v109_})
                d_1258_v115_: _dafny.Seq
                d_1258_v115_ = _dafny.SeqWithoutIsStrInference([d_1256_v113_, (d_1256_v113_) | (d_1256_v113_), (self).f26, (_dafny.Map({(0) - (p0): d_1215_v84_})).set(len(d_1257_v114_), d_1215_v84_)])
                d_1259_v116_: _dafny.MultiSet
                d_1259_v116_ = _dafny.MultiSet([default__.fm7((self).f27, globalState), 122, len(d_1250_v107_), 107])
                d_1260_v117_: _dafny.Map
                d_1260_v117_ = _dafny.Map({(d_1259_v116_).ispropersubset(d_1259_v116_): not(not(d_1215_v84_))})
                rhs224_ = (p2) < (p0)
                rhs225_ = d_1215_v84_
                rhs226_ = d_1211_i8_
                rhs227_ = d_1258_v115_
                rhs228_ = d_1260_v117_
                lhs196_ = globalState
                lhs197_ = globalState
                lhs198_ = globalState
                lhs196_.f0 = rhs224_
                lhs197_.f0 = rhs225_
                lhs198_.f8 = rhs226_
                d_1258_v115_ = rhs227_
                d_1260_v117_ = rhs228_
        d_1261_v119_: _dafny.Map
        d_1261_v119_ = _dafny.Map({(self).f27: p2})
        d_1262_v120_: bool
        d_1262_v120_ = False
        d_1263_v121_: _dafny.Map
        d_1263_v121_ = _dafny.Map({d_1262_v120_: p2})
        d_1264_v122_: _dafny.MultiSet
        d_1264_v122_ = _dafny.MultiSet([len(d_1263_v121_), (self).f27])
        d_1265_v123_: _dafny.Map
        d_1265_v123_ = _dafny.Map({len(p1): d_1264_v122_})
        d_1266_v124_: _dafny.Map
        def iife100_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in ((d_1261_v119_).set((self).f27, p0)).keys.Elements:
                d_1267_v118_: int = compr_52_
                if (d_1267_v118_) in ((d_1261_v119_).set((self).f27, p0)):
                    coll52_[(d_1267_v118_) * ((self).f27)] = _dafny.MultiSet([(self).f27])
            return _dafny.Map(coll52_)
        d_1266_v124_ = _dafny.Map({(iife100_()
        ) | ((D4_DC12(d_1265_v123_)).cf21): d_1262_v120_})
        d_1268_v125_: _dafny.Seq
        d_1268_v125_ = _dafny.SeqWithoutIsStrInference([d_1265_v123_])
        d_1266_v124_ = (d_1266_v124_).set((d_1268_v125_)[default__.safeIndex(p0, len(d_1268_v125_))], d_1262_v120_)
        if d_1262_v120_:
            (globalState).f6 = len(p1)
            d_1269_v126_: _dafny.Map
            d_1269_v126_ = _dafny.Map({d_1262_v120_: d_1262_v120_})
            d_1270_v127_: _dafny.Array
            def lambda51_(d_1271_v120_):
                def lambda52_(d_1272_i11_):
                    return d_1271_v120_

                return lambda52_

            init26_ = lambda51_(d_1262_v120_)
            nw170_ = _dafny.Array(None, 28)
            for i0_26_ in range(nw170_.length(0)):
                nw170_[i0_26_] = init26_(i0_26_)
            d_1270_v127_ = nw170_
            d_1273_v128_: _dafny.Seq
            d_1273_v128_ = _dafny.SeqWithoutIsStrInference([d_1270_v127_])
            d_1274_v129_: C2
            nw171_ = C2()
            nw171_.ctor__(len((d_1269_v126_ if d_1262_v120_ else d_1269_v126_)), (_dafny.SeqWithoutIsStrInference([d_1270_v127_, d_1270_v127_])) + (d_1273_v128_))
            d_1274_v129_ = nw171_
            d_1275_v130_: _dafny.Seq
            d_1275_v130_ = _dafny.SeqWithoutIsStrInference([d_1264_v122_])
            d_1276_v131_: D11
            d_1276_v131_ = D11_DC36(d_1275_v130_)
            rhs229_ = d_1262_v120_
            rhs230_ = d_1276_v131_
            lhs199_ = globalState
            lhs199_.f10 = rhs229_
            d_1276_v131_ = rhs230_
            d_1263_v121_ = (d_1263_v121_).set(d_1262_v120_, len(p1))
            source20_ = d_1093_v0_
            if source20_.is_DC0:
                d_1277_v132_: C3
                nw172_ = C3()
                nw172_.ctor__((d_1273_v128_) + (d_1273_v128_))
                d_1277_v132_ = nw172_
                (globalState).f8 = (self).f27
                d_1278_v133_: _dafny.MultiSet
                d_1278_v133_ = _dafny.MultiSet([d_1262_v120_, d_1262_v120_])
                d_1279_v134_: _dafny.MultiSet
                d_1279_v134_ = _dafny.MultiSet([d_1278_v133_, d_1278_v133_, d_1278_v133_])
                (globalState).f1 = (d_1279_v134_) - (d_1279_v134_)
                index189_ = default__.safeIndex(27, (d_1270_v127_).length(0))
                (d_1270_v127_)[index189_] = False
                index190_ = default__.safeIndex(27, (d_1270_v127_).length(0))
                (d_1270_v127_)[index190_] = not(not (d_1262_v120_) or (True))
            elif True:
                d_1280___mcc_h13_ = source20_.cf0
                d_1281___mcc_h14_ = source20_.cf1
                d_1282___mcc_h15_ = source20_.cf2
                d_1283___mcc_h16_ = source20_.cf3
                d_1284___mcc_h17_ = source20_.cf4
                d_1285_cf4_ = d_1284___mcc_h17_
                d_1286_cf3_ = d_1283___mcc_h16_
                d_1287_cf2_ = d_1282___mcc_h15_
                d_1288_cf1_ = d_1281___mcc_h14_
                d_1289_cf0_ = d_1280___mcc_h13_
                d_1290_v135_: _dafny.Seq
                d_1290_v135_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "riagx"))
                rhs231_ = p1
                rhs232_ = d_1286_cf3_
                d_1290_v135_ = rhs231_
                d_1286_cf3_ = rhs232_
                d_1291_v136_: C0
                nw173_ = C0()
                nw173_.ctor__()
                d_1291_v136_ = nw173_
                d_1292_v137_: C4
                nw174_ = C4()
                nw174_.ctor__(not(not(d_1262_v120_)), d_1273_v128_)
                d_1292_v137_ = nw174_
                d_1293_v138_: _dafny.Set
                d_1293_v138_ = _dafny.Set({d_1292_v137_})
                (globalState).f8 = ((p2) * (len(d_1293_v138_))) - (len(p1))
                d_1294_v139_: _dafny.Seq
                d_1294_v139_ = _dafny.SeqWithoutIsStrInference([(d_1274_v129_).f29, (d_1274_v129_).f29])
                d_1295_v140_: D2
                d_1295_v140_ = D2_DC6(len(d_1294_v139_))
                (globalState).f0 = (((self).f26)[(d_1289_cf0_) + (len(default__.fm14(p2, (d_1295_v140_).cf9, globalState)))] if ((d_1289_cf0_) + (len(default__.fm14(p2, (d_1295_v140_).cf9, globalState)))) in ((self).f26) else ((self).f27) <= ((d_1274_v129_).f29))
        elif True:
            d_1296_v141_: _dafny.Map
            d_1296_v141_ = _dafny.Map({p1: p0})
            d_1297_v142_: _dafny.Seq
            d_1297_v142_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            d_1296_v141_ = (d_1296_v141_).set(((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kceepx")))).set(default__.safeIndex(p0, len((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kceepx"))))), _dafny.CodePoint('j')), (_dafny.MultiSet([p2, len(d_1297_v142_), p0])).cardinality)
            d_1298_v143_: D4
            d_1298_v143_ = D4_DC15(p2, len(d_1263_v121_))
            (globalState).f10 = not((default__.fm7((d_1298_v143_).cf26, globalState)) == (p2))
            d_1299_v144_: _dafny.Seq
            d_1299_v144_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxtwiau"))
            d_1299_v144_ = d_1299_v144_
            d_1300_v145_: _dafny.Map
            d_1300_v145_ = _dafny.Map({p0: d_1299_v144_})
            d_1301_v146_: _dafny.Set
            d_1301_v146_ = _dafny.Set({d_1262_v120_})
            d_1302_v147_: _dafny.Set
            d_1302_v147_ = _dafny.Set({len(d_1301_v146_), p0})
            d_1303_v148_: _dafny.Map
            d_1303_v148_ = _dafny.Map({d_1302_v147_: d_1262_v120_})
            d_1300_v145_ = (d_1300_v145_).set((p2) + (len(d_1303_v148_)), (p1 if not(d_1262_v120_) else p1))
            (globalState).f0 = d_1262_v120_
        d_1304_v149_: str
        d_1304_v149_ = _dafny.CodePoint('q')
        d_1305_v150_: _dafny.Array
        nw175_ = _dafny.Array(False, 25)
        d_1305_v150_ = nw175_
        pat_let_tv55_ = d_1304_v149_
        d_1306_v151_: C2
        nw176_ = C2()
        def iife101_(_pat_let24_0):
            def iife102_(d_1307_dt__update__tmp_h2_):
                def iife103_(_pat_let25_0):
                    def iife104_(d_1308_dt__update_hcf23_h0_):
                        return D4_DC14(d_1308_dt__update_hcf23_h0_, (d_1307_dt__update__tmp_h2_).cf24, (d_1307_dt__update__tmp_h2_).cf25)
                    return iife104_(_pat_let25_0)
                return iife103_(pat_let_tv55_)
            return iife102_(_pat_let24_0)
        nw176_.ctor__(p2, _dafny.SeqWithoutIsStrInference([(iife101_(D4_DC14(d_1304_v149_, d_1305_v150_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jixw"))))).cf24]))
        d_1306_v151_ = nw176_
        d_1309_v152_: _dafny.MultiSet
        d_1309_v152_ = _dafny.MultiSet([False, d_1262_v120_])
        d_1310_v153_: _dafny.Seq
        d_1310_v153_ = _dafny.SeqWithoutIsStrInference([d_1262_v120_])
        (globalState).f11 = len((default__.fm24((d_1309_v152_) - (d_1309_v152_), p2, (len(d_1310_v153_)) - (p2), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrtm"))) != (p1), globalState)).set(default__.safeIndex((0) - ((d_1306_v151_).f29), len(default__.fm24((d_1309_v152_) - (d_1309_v152_), p2, (len(d_1310_v153_)) - (p2), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrtm"))) != (p1), globalState))), p2))

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: D0 = D0.default()()
        d_1311_v0_: _dafny.Array
        nw177_ = _dafny.Array(False, 16)
        d_1311_v0_ = nw177_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1311_v0_).length(0)):
            d_1312_i0_: int = guard_loop_2_
            if (True) and (((0) <= (d_1312_i0_)) and ((d_1312_i0_) < ((d_1311_v0_).length(0)))):
                (d_1311_v0_)[(d_1312_i0_)] = not(not((default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtnfgu"))), (self).f27)) > ((self).f27)))
        d_1313_v1_: _dafny.Seq
        d_1313_v1_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        d_1314_v2_: _dafny.Seq
        d_1314_v2_ = _dafny.SeqWithoutIsStrInference([d_1313_v1_])
        d_1315_v4_: D7
        def iife105_():
            coll53_ = _dafny.Set()
            compr_53_: _dafny.Seq
            for compr_53_ in (d_1314_v2_).Elements:
                d_1316_v3_: _dafny.Seq = compr_53_
                if (d_1316_v3_) in (d_1314_v2_):
                    coll53_ = coll53_.union(_dafny.Set([d_1316_v3_]))
            return _dafny.Set(coll53_)
        d_1315_v4_ = D7_DC25(iife105_()
)
        d_1317_v5_: _dafny.Map
        d_1317_v5_ = _dafny.Map({d_1315_v4_: (self).f27})
        d_1318_v7_: _dafny.Map
        def iife106_():
            coll54_ = _dafny.Map()
            compr_54_: int
            for compr_54_ in _dafny.IntegerRange(-409, 24):
                d_1319_v6_: int = compr_54_
                if ((-409) <= (d_1319_v6_)) and ((d_1319_v6_) < (24)):
                    coll54_[(d_1319_v6_) * ((self).f27)] = len(p0)
            return _dafny.Map(coll54_)
        d_1318_v7_ = _dafny.Map({p1: iife106_()
        })
        d_1320_v8_: _dafny.Seq
        d_1320_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_1321_v9_: str
        d_1321_v9_ = _dafny.CodePoint('u')
        d_1322_v10_: D0
        d_1322_v10_ = D0_DC1(len(d_1317_v5_), p1, d_1318_v7_, len(_dafny.SeqWithoutIsStrInference([len(d_1320_v8_)])), d_1321_v9_)
        d_1323_v11_: _dafny.MultiSet
        d_1323_v11_ = _dafny.MultiSet([d_1322_v10_, d_1322_v10_])
        d_1324_v12_: _dafny.Array
        nw178_ = _dafny.Array(None, 1)
        nw178_[int(0)] = (_dafny.MultiSet([d_1322_v10_])) - (d_1323_v11_)
        d_1324_v12_ = nw178_
        index191_ = default__.safeIndex(879, (d_1324_v12_).length(0))
        (d_1324_v12_)[index191_] = d_1323_v11_
        index192_ = default__.safeIndex(879, (d_1324_v12_).length(0))
        (d_1324_v12_)[index192_] = d_1323_v11_
        d_1325_v13_: D6
        d_1325_v13_ = D6_DC22((self).f27)
        hi3_ = default__.safeDivisionInt((self).f27, (d_1325_v13_).cf36)
        for d_1326_i1_ in range((self).f27, hi3_):
            d_1327_v14_: C0
            nw179_ = C0()
            nw179_.ctor__()
            d_1327_v14_ = nw179_
            d_1328_v15_: _dafny.Seq
            d_1328_v15_ = _dafny.SeqWithoutIsStrInference([d_1327_v14_, d_1327_v14_])
            d_1329_v16_: _dafny.Array
            nw180_ = _dafny.Array(None, 21)
            nw180_[int(0)] = d_1327_v14_
            nw180_[int(1)] = d_1327_v14_
            nw180_[int(2)] = d_1327_v14_
            nw180_[int(3)] = d_1327_v14_
            nw180_[int(4)] = d_1327_v14_
            nw180_[int(5)] = d_1327_v14_
            nw180_[int(6)] = (d_1328_v15_)[default__.safeIndex(d_1326_i1_, len(d_1328_v15_))]
            nw180_[int(7)] = d_1327_v14_
            nw180_[int(8)] = d_1327_v14_
            nw180_[int(9)] = d_1327_v14_
            nw180_[int(10)] = d_1327_v14_
            nw180_[int(11)] = d_1327_v14_
            nw180_[int(12)] = d_1327_v14_
            nw180_[int(13)] = d_1327_v14_
            nw180_[int(14)] = d_1327_v14_
            nw180_[int(15)] = d_1327_v14_
            nw180_[int(16)] = d_1327_v14_
            nw180_[int(17)] = d_1327_v14_
            nw180_[int(18)] = d_1327_v14_
            nw180_[int(19)] = d_1327_v14_
            nw180_[int(20)] = d_1327_v14_
            d_1329_v16_ = nw180_
            d_1330_v17_: _dafny.Array
            nw181_ = _dafny.Array(None, 4)
            nw181_[int(0)] = d_1329_v16_
            nw181_[int(1)] = d_1329_v16_
            nw181_[int(2)] = d_1329_v16_
            nw181_[int(3)] = d_1329_v16_
            d_1330_v17_ = nw181_
            d_1330_v17_ = d_1330_v17_
            d_1331_v18_: _dafny.Seq
            d_1331_v18_ = _dafny.SeqWithoutIsStrInference([d_1311_v0_, d_1311_v0_, d_1311_v0_, d_1311_v0_])
            d_1332_v19_: C4
            nw182_ = C4()
            nw182_.ctor__(p1, d_1331_v18_)
            d_1332_v19_ = nw182_
            (d_1332_v19_).f28 = p1
            d_1333_v20_: _dafny.Seq
            d_1333_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "myxqtmm"))
            d_1333_v20_ = d_1333_v20_
        d_1334_v21_: _dafny.Array
        nw183_ = _dafny.Array(int(0), 27)
        d_1334_v21_ = nw183_
        index193_ = default__.safeIndex(259, (d_1334_v21_).length(0))
        (d_1334_v21_)[index193_] = 361
        index194_ = default__.safeIndex(259, (d_1334_v21_).length(0))
        (d_1334_v21_)[index194_] = (self).f27
        d_1335_v22_: D2
        d_1335_v22_ = D2_DC6((d_1334_v21_)[default__.safeIndex(259, (d_1334_v21_).length(0))])
        d_1336_v23_: _dafny.Map
        d_1336_v23_ = _dafny.Map({(d_1335_v22_).cf9: (self).f27})
        d_1336_v23_ = (d_1336_v23_).set(default__.safeDivisionInt((self).f27, (self).f27), (self).f27)
        rhs233_ = D6_DC22((self).f27)
        rhs234_ = p1
        lhs200_ = globalState
        d_1325_v13_ = rhs233_
        lhs200_.f10 = rhs234_
        r0 = default__.safeDivisionInt(len(d_1320_v8_), (0) - ((len(p0) if p1 else (d_1334_v21_)[default__.safeIndex(259, (d_1334_v21_).length(0))])))
        r1 = D0_DC0()
        return r0, r1

    @property
    def f26(self):
        return self._f26
    @property
    def f27(self):
        return self._f27

class C6(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self._f24: bool = False
        self._f25: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f24, f25, f22):
        (self)._f24 = f24
        (self)._f25 = f25
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return (default__.safeModuloInt(378, -343)) + (-858)

    def fm3(self, p0, p1, p2, globalState):
        return (self).f24

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_1337_v0_: _dafny.Array
        nw184_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
        d_1337_v0_ = nw184_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1337_v0_).length(0)):
            d_1338_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1338_i0_)) and ((d_1338_i0_) < ((d_1337_v0_).length(0)))):
                (d_1337_v0_)[(d_1338_i0_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tri")) if p0 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1339_i1_ in range(default__.abs(268))]))
        (globalState).f0 = (self).f24
        d_1340_v1_: D0
        d_1340_v1_ = D0_DC0()
        source21_ = d_1340_v1_
        if source21_.is_DC0:
            d_1341_v2_: int
            d_1341_v2_ = 468
            (globalState).f16 = d_1341_v2_
            d_1342_v3_: C3
            nw185_ = C3()
            nw185_.ctor__(((self).f22) + ((self).f22))
            d_1342_v3_ = nw185_
            if p0:
                d_1343_v4_: C2
                nw186_ = C2()
                nw186_.ctor__(d_1341_v2_, ((self).f22) + ((self).f22))
                d_1343_v4_ = nw186_
                d_1344_v5_: _dafny.Array
                def lambda53_(d_1345_i2_):
                    return D11_DC39((self).f24)

                init27_ = lambda53_
                nw187_ = _dafny.Array(None, 17)
                for i0_27_ in range(nw187_.length(0)):
                    nw187_[i0_27_] = init27_(i0_27_)
                d_1344_v5_ = nw187_
                d_1346_v6_: D11
                d_1346_v6_ = D11_DC39((self).f24)
                index195_ = default__.safeIndex(817, (d_1344_v5_).length(0))
                (d_1344_v5_)[index195_] = d_1346_v6_
                index196_ = default__.safeIndex(817, (d_1344_v5_).length(0))
                (d_1344_v5_)[index196_] = d_1346_v6_
                d_1347_v7_: _dafny.Seq
                d_1347_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfs"))
                d_1348_v8_: _dafny.Array
                def lambda54_(d_1349_v4_):
                    def lambda55_(d_1350_i3_):
                        return default__.safeDivisionInt(d_1350_i3_, (d_1349_v4_).f29)

                    return lambda55_

                init28_ = lambda54_(d_1343_v4_)
                nw188_ = _dafny.Array(None, 17)
                for i0_28_ in range(nw188_.length(0)):
                    nw188_[i0_28_] = init28_(i0_28_)
                d_1348_v8_ = nw188_
                d_1351_v9_: _dafny.Map
                d_1351_v9_ = _dafny.Map({len(d_1347_v7_): d_1348_v8_})
                d_1352_v10_: _dafny.Seq
                d_1352_v10_ = _dafny.SeqWithoutIsStrInference([len(d_1347_v7_), (d_1343_v4_).f29])
                d_1353_v11_: _dafny.Map
                d_1353_v11_ = _dafny.Map({(d_1343_v4_).f29: len(d_1352_v10_)})
                d_1354_v12_: _dafny.Seq
                d_1354_v12_ = _dafny.SeqWithoutIsStrInference([((d_1353_v11_)[len(d_1352_v10_)] if (len(d_1352_v10_)) in (d_1353_v11_) else default__.fm7(d_1341_v2_, globalState))])
                d_1355_v13_: D2
                d_1355_v13_ = D2_DC5(d_1348_v8_)
                d_1351_v9_ = (d_1351_v9_).set(len(_dafny.Set({d_1341_v2_, d_1341_v2_, (d_1354_v12_)[default__.safeIndex(d_1341_v2_, len(d_1354_v12_))], len(default__.fm33(globalState))})), (d_1355_v13_).cf8)
                d_1356_v14_: str
                d_1356_v14_ = _dafny.CodePoint('q')
                d_1357_v15_: _dafny.MultiSet
                d_1357_v15_ = _dafny.MultiSet([(self).f24])
                d_1358_v16_: _dafny.Map
                d_1358_v16_ = _dafny.Map({not (p0) or ((self).f24): (_dafny.MultiSet([502, len(_dafny.Map({d_1356_v14_: (d_1343_v4_).f29}))])).set((d_1357_v15_).cardinality, default__.abs(d_1341_v2_))})
                d_1358_v16_ = (d_1358_v16_).set((self).f24, (_dafny.MultiSet(d_1352_v10_)).intersection(_dafny.MultiSet([(d_1343_v4_).f29])))
                d_1359_v17_: _dafny.MultiSet
                d_1359_v17_ = _dafny.MultiSet([(d_1343_v4_).f29, d_1341_v2_, len(d_1347_v7_), d_1341_v2_])
                d_1360_v21_: _dafny.Seq
                def iife107_():
                    def iife109_():
                        coll57_ = _dafny.Map()
                        compr_57_: int
                        for compr_57_ in _dafny.IntegerRange(325, -867):
                            d_1361_v18_: int = compr_57_
                            if ((325) <= (d_1361_v18_)) and ((d_1361_v18_) < (-867)):
                                coll57_[(d_1361_v18_) * (d_1341_v2_)] = p0
                        return _dafny.Map(coll57_)
                    coll55_ = _dafny.Set()
                    def iife108_():
                        coll56_ = _dafny.Map()
                        compr_56_: int
                        for compr_56_ in _dafny.IntegerRange(325, -867):
                            d_1361_v18_: int = compr_56_
                            if ((325) <= (d_1361_v18_)) and ((d_1361_v18_) < (-867)):
                                coll56_[(d_1361_v18_) * (d_1341_v2_)] = p0
                        return _dafny.Map(coll56_)
                    compr_55_: int
                    for compr_55_ in (iife108_()
                    ).keys.Elements:
                        d_1362_v19_: int = compr_55_
                        if (d_1362_v19_) in (iife109_()
                        ):
                            def iife110_():
                                coll58_ = _dafny.Map()
                                compr_58_: int
                                for compr_58_ in (_dafny.Map({219: len(_dafny.Set({not(False), False, False, True, False}))})).keys.Elements:
                                    d_1363_v20_: int = compr_58_
                                    if (d_1363_v20_) in (_dafny.Map({219: len(_dafny.Set({not(False), False, False, True, False}))})):
                                        coll58_[default__.safeDivisionInt(d_1363_v20_, len(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), True]))).cardinality})))] = _dafny.MultiSet([len(_dafny.Map({True: (0) - (-324)})), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqmipstkb"))): _dafny.CodePoint('g')}))])
                                return _dafny.Map(coll58_)
                            coll55_ = coll55_.union(_dafny.Set([(d_1362_v19_) * (len(iife110_()
))]))
                    return _dafny.Set(coll55_)
                d_1360_v21_ = _dafny.SeqWithoutIsStrInference([iife107_()
                ])
                d_1364_v22_: _dafny.Set
                d_1364_v22_ = _dafny.Set({d_1341_v2_})
                (globalState).f11 = (d_1343_v4_).fm0(d_1359_v17_, 574, ((d_1360_v21_).set(default__.safeIndex(816, len(d_1360_v21_)), _dafny.Set({(d_1343_v4_).f29, len(d_1354_v12_), d_1341_v2_, (d_1343_v4_).f29}))).set(default__.safeIndex((0) - ((d_1343_v4_).f29), len((d_1360_v21_).set(default__.safeIndex(816, len(d_1360_v21_)), _dafny.Set({(d_1343_v4_).f29, len(d_1354_v12_), d_1341_v2_, (d_1343_v4_).f29})))), d_1364_v22_), p0, globalState)
            elif True:
                d_1365_v23_: _dafny.Seq
                d_1365_v23_ = _dafny.SeqWithoutIsStrInference([d_1341_v2_])
                d_1366_v24_: _dafny.Seq
                d_1366_v24_ = _dafny.SeqWithoutIsStrInference([d_1365_v23_])
                d_1367_v25_: _dafny.Set
                d_1367_v25_ = _dafny.Set({d_1341_v2_, d_1341_v2_, d_1341_v2_, d_1341_v2_})
                r0 = not((d_1365_v23_) == ((d_1366_v24_)[default__.safeIndex(len(d_1367_v25_), len(d_1366_v24_))]))
                d_1337_v0_ = d_1337_v0_
                (globalState).f6 = d_1341_v2_
                r3 = (d_1341_v2_ if p0 else d_1341_v2_)
                (globalState).f0 = (self).f24
            d_1368_v26_: _dafny.Map
            d_1368_v26_ = _dafny.Map({_dafny.CodePoint('s'): (self).f24})
            d_1369_v27_: str
            d_1369_v27_ = _dafny.CodePoint('m')
            d_1370_v28_: _dafny.Array
            nw189_ = _dafny.Array(None, 9)
            nw189_[int(0)] = (self).f24
            nw189_[int(1)] = p0
            nw189_[int(2)] = not(((d_1368_v26_)[d_1369_v27_] if (d_1369_v27_) in (d_1368_v26_) else (self).f24))
            nw189_[int(3)] = p0
            nw189_[int(4)] = (self).f24
            nw189_[int(5)] = p0
            nw189_[int(6)] = (self).f24
            nw189_[int(7)] = p0
            nw189_[int(8)] = p0
            d_1370_v28_ = nw189_
            d_1371_v29_: C3
            nw190_ = C3()
            nw190_.ctor__(_dafny.SeqWithoutIsStrInference([d_1370_v28_, d_1370_v28_, d_1370_v28_, d_1370_v28_, d_1370_v28_]))
            d_1371_v29_ = nw190_
        elif True:
            d_1372___mcc_h0_ = source21_.cf0
            d_1373___mcc_h1_ = source21_.cf1
            d_1374___mcc_h2_ = source21_.cf2
            d_1375___mcc_h3_ = source21_.cf3
            d_1376___mcc_h4_ = source21_.cf4
            d_1377_cf4_ = d_1376___mcc_h4_
            d_1378_cf3_ = d_1375___mcc_h3_
            d_1379_cf2_ = d_1374___mcc_h2_
            d_1380_cf1_ = d_1373___mcc_h1_
            d_1381_cf0_ = d_1372___mcc_h0_
            d_1382_v30_: bool
            d_1383_v31_: int
            out56_: bool
            out57_: int
            out56_, out57_ = (self).m1(globalState)
            d_1382_v30_ = out56_
            d_1383_v31_ = out57_
            d_1384_v32_: _dafny.Array
            def lambda56_(d_1385_cf0_):
                def lambda57_(d_1386_i4_):
                    return (d_1386_i4_) * (d_1385_cf0_)

                return lambda57_

            init29_ = lambda56_(d_1381_cf0_)
            nw191_ = _dafny.Array(None, 1)
            for i0_29_ in range(nw191_.length(0)):
                nw191_[i0_29_] = init29_(i0_29_)
            d_1384_v32_ = nw191_
            index197_ = default__.safeIndex(784, (d_1384_v32_).length(0))
            (d_1384_v32_)[index197_] = d_1378_cf3_
            index198_ = default__.safeIndex(784, (d_1384_v32_).length(0))
            (d_1384_v32_)[index198_] = ((0) - (d_1381_cf0_)) - (d_1383_v31_)
            if d_1382_v30_:
                d_1387_v33_: _dafny.Array
                nw192_ = _dafny.Array(_dafny.Map({}), 12)
                d_1387_v33_ = nw192_
                d_1387_v33_ = d_1387_v33_
                d_1388_v34_: C0
                nw193_ = C0()
                nw193_.ctor__()
                d_1388_v34_ = nw193_
                (globalState).f6 = default__.safeModuloInt(d_1381_cf0_, d_1381_cf0_)
                d_1389_v35_: _dafny.Array
                def lambda58_(d_1390_cf0_):
                    def lambda59_(d_1391_i5_):
                        return _dafny.Set({d_1390_cf0_, d_1390_cf0_})

                    return lambda59_

                init30_ = lambda58_(d_1381_cf0_)
                nw194_ = _dafny.Array(None, 28)
                for i0_30_ in range(nw194_.length(0)):
                    nw194_[i0_30_] = init30_(i0_30_)
                d_1389_v35_ = nw194_
                d_1392_v36_: _dafny.Map
                d_1392_v36_ = _dafny.Map({d_1378_cf3_: d_1383_v31_})
                d_1393_v37_: _dafny.Seq
                d_1393_v37_ = _dafny.SeqWithoutIsStrInference([d_1392_v36_, d_1392_v36_])
                d_1394_v38_: _dafny.Map
                d_1394_v38_ = _dafny.Map({p0: d_1383_v31_})
                d_1395_v40_: _dafny.Set
                def iife111_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(-383, -369):
                        d_1396_v39_: int = compr_59_
                        if ((-383) <= (d_1396_v39_)) and ((d_1396_v39_) < (-369)):
                            coll59_[default__.safeModuloInt(d_1396_v39_, (d_1384_v32_)[default__.safeIndex(784, (d_1384_v32_).length(0))])] = d_1381_cf0_
                    return _dafny.Map(coll59_)
                d_1395_v40_ = _dafny.Set({len((d_1393_v37_).set(default__.safeIndex(len(d_1394_v38_), len(d_1393_v37_)), iife111_()
                )), (d_1384_v32_)[default__.safeIndex(784, (d_1384_v32_).length(0))], d_1381_cf0_, (0) - (default__.fm7(d_1381_cf0_, globalState)), d_1378_cf3_})
                index199_ = default__.safeIndex(773, (d_1389_v35_).length(0))
                (d_1389_v35_)[index199_] = d_1395_v40_
                index200_ = default__.safeIndex(773, (d_1389_v35_).length(0))
                def iife112_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(460, -440):
                        d_1397_v41_: int = compr_60_
                        if ((460) <= (d_1397_v41_)) and ((d_1397_v41_) < (-440)):
                            coll60_ = coll60_.union(_dafny.Set([(d_1397_v41_) * ((d_1384_v32_)[default__.safeIndex(784, (d_1384_v32_).length(0))])]))
                    return _dafny.Set(coll60_)
                (d_1389_v35_)[index200_] = (iife112_()
                ) | (d_1395_v40_)
                (globalState).f13 = (d_1389_v35_)[default__.safeIndex(773, (d_1389_v35_).length(0))]
            elif True:
                (globalState).f6 = (0) - (d_1381_cf0_)
                d_1398_v42_: _dafny.Set
                d_1398_v42_ = _dafny.Set({d_1377_cf4_, _dafny.CodePoint('w'), d_1377_cf4_, d_1377_cf4_})
                (globalState).f10 = (_dafny.Set({d_1377_cf4_})).ispropersubset(d_1398_v42_)
                d_1399_v43_: _dafny.Seq
                d_1399_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkn"))
                (globalState).f11 = len((d_1399_v43_ if not((self).f24) else (d_1399_v43_) + (d_1399_v43_)))
                d_1400_v44_: _dafny.Seq
                d_1400_v44_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1401_v45_: _dafny.Array
                nw195_ = _dafny.Array(None, 4)
                nw195_[int(0)] = default__.fm37(d_1400_v44_, _dafny.Map({d_1381_cf0_: d_1382_v30_}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqe")), globalState)
                nw195_[int(1)] = p0
                nw195_[int(2)] = p0
                nw195_[int(3)] = p0
                d_1401_v45_ = nw195_
                d_1402_v46_: C4
                nw196_ = C4()
                nw196_.ctor__(d_1380_cf1_, _dafny.SeqWithoutIsStrInference([d_1401_v45_, d_1401_v45_, d_1401_v45_]))
                d_1402_v46_ = nw196_
                d_1403_v47_: _dafny.Seq
                d_1403_v47_ = _dafny.SeqWithoutIsStrInference([d_1337_v0_])
                d_1404_v48_: _dafny.Array
                nw197_ = _dafny.Array(None, 20)
                nw197_[int(0)] = d_1337_v0_
                nw197_[int(1)] = (d_1403_v47_)[default__.safeIndex(d_1378_cf3_, len(d_1403_v47_))]
                nw197_[int(2)] = d_1337_v0_
                nw197_[int(3)] = d_1337_v0_
                nw197_[int(4)] = (d_1403_v47_)[default__.safeIndex((d_1384_v32_)[default__.safeIndex(784, (d_1384_v32_).length(0))], len(d_1403_v47_))]
                nw197_[int(5)] = d_1337_v0_
                nw197_[int(6)] = d_1337_v0_
                nw197_[int(7)] = d_1337_v0_
                nw197_[int(8)] = d_1337_v0_
                nw197_[int(9)] = d_1337_v0_
                nw197_[int(10)] = d_1337_v0_
                nw197_[int(11)] = d_1337_v0_
                nw197_[int(12)] = d_1337_v0_
                nw197_[int(13)] = d_1337_v0_
                nw197_[int(14)] = d_1337_v0_
                nw197_[int(15)] = d_1337_v0_
                nw197_[int(16)] = d_1337_v0_
                nw197_[int(17)] = d_1337_v0_
                nw197_[int(18)] = d_1337_v0_
                nw197_[int(19)] = d_1337_v0_
                d_1404_v48_ = nw197_
                index201_ = default__.safeIndex(567, (d_1404_v48_).length(0))
                (d_1404_v48_)[index201_] = d_1337_v0_
                index202_ = default__.safeIndex(567, (d_1404_v48_).length(0))
                (d_1404_v48_)[index202_] = (d_1337_v0_ if ((self).f24 if d_1382_v30_ else d_1382_v30_) else d_1337_v0_)
            d_1405_v49_: _dafny.Seq
            d_1405_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_1377_cf4_ = (d_1405_v49_)[default__.safeIndex(d_1378_cf3_, len(d_1405_v49_))]
        d_1406_v50_: int
        d_1406_v50_ = -250
        if default__.fm15((_dafny.MultiSet([d_1406_v50_])).issubset(_dafny.MultiSet([d_1406_v50_])), globalState):
            d_1407_v51_: _dafny.Array
            def lambda60_(d_1408_p0_):
                def lambda61_(d_1409_i6_):
                    return _dafny.MultiSet([_dafny.MultiSet([(self).f24, (self).f24]), _dafny.MultiSet([d_1408_p0_]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f24]))])

                return lambda61_

            init31_ = lambda60_(p0)
            nw198_ = _dafny.Array(None, 8)
            for i0_31_ in range(nw198_.length(0)):
                nw198_[i0_31_] = init31_(i0_31_)
            d_1407_v51_ = nw198_
            d_1410_v52_: _dafny.Seq
            d_1410_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnkxhr"))
            d_1411_v53_: _dafny.Map
            d_1411_v53_ = _dafny.Map({(self).fm3(d_1410_v52_, (self).f24, False, globalState): (self).f24})
            d_1412_v54_: _dafny.Set
            d_1412_v54_ = _dafny.Set({((d_1411_v53_)[False] if (False) in (d_1411_v53_) else (self).f24), p0, (self).f24, (self).f24, (self).f24})
            d_1413_v55_: _dafny.MultiSet
            d_1413_v55_ = _dafny.MultiSet([(self).f24, (self).f24])
            d_1414_v56_: _dafny.MultiSet
            d_1414_v56_ = _dafny.MultiSet([default__.fm41(d_1406_v50_, d_1410_v52_, d_1412_v54_, globalState), d_1413_v55_, d_1413_v55_])
            index203_ = default__.safeIndex(633, (d_1407_v51_).length(0))
            (d_1407_v51_)[index203_] = d_1414_v56_
            index204_ = default__.safeIndex(633, (d_1407_v51_).length(0))
            (d_1407_v51_)[index204_] = (d_1414_v56_) - (d_1414_v56_)
            d_1415_v57_: _dafny.Array
            def lambda62_(d_1416_p0_):
                def lambda63_(d_1417_i7_):
                    return d_1416_p0_

                return lambda63_

            init32_ = lambda62_(p0)
            nw199_ = _dafny.Array(None, 3)
            for i0_32_ in range(nw199_.length(0)):
                nw199_[i0_32_] = init32_(i0_32_)
            d_1415_v57_ = nw199_
            index205_ = default__.safeIndex(378, (d_1415_v57_).length(0))
            (d_1415_v57_)[index205_] = not(p0)
            d_1418_v58_: str
            d_1418_v58_ = _dafny.CodePoint('i')
            d_1419_v59_: C1
            nw200_ = C1()
            nw200_.ctor__(d_1406_v50_, (self).f22)
            d_1419_v59_ = nw200_
            d_1420_v60_: _dafny.Map
            d_1420_v60_ = _dafny.Map({p0: d_1419_v59_})
            d_1421_v61_: _dafny.Map
            d_1421_v61_ = _dafny.Map({d_1418_v58_: d_1420_v60_})
            index206_ = default__.safeIndex(378, (d_1415_v57_).length(0))
            (d_1415_v57_)[index206_] = (not ((self).f24) or (False) if (False if False else (self).f24) else (d_1406_v50_) < (len(d_1421_v61_)))
            d_1422_v62_: D4
            d_1422_v62_ = D4_DC15(((d_1419_v59_).f30) * (len(d_1410_v52_)), (d_1419_v59_).f30)
            source22_ = d_1422_v62_
            if source22_.is_DC13:
                d_1423___mcc_h5_ = source22_.cf22
                d_1424_cf22_ = d_1423___mcc_h5_
                index207_ = default__.safeIndex(378, (d_1415_v57_).length(0))
                (d_1415_v57_)[index207_] = (self).f24
                d_1418_v58_ = d_1418_v58_
                d_1419_v59_ = d_1419_v59_
                d_1425_v63_: _dafny.Array
                nw201_ = _dafny.Array(_dafny.MultiSet({}), 6)
                d_1425_v63_ = nw201_
                d_1426_v64_: _dafny.MultiSet
                d_1426_v64_ = _dafny.MultiSet([D6_DC22((d_1419_v59_).f30)])
                index208_ = default__.safeIndex(668, (d_1425_v63_).length(0))
                (d_1425_v63_)[index208_] = d_1426_v64_
                index209_ = default__.safeIndex(668, (d_1425_v63_).length(0))
                (d_1425_v63_)[index209_] = default__.fm46(globalState)
            elif source22_.is_DC14:
                d_1427___mcc_h6_ = source22_.cf23
                d_1428___mcc_h7_ = source22_.cf24
                d_1429___mcc_h8_ = source22_.cf25
                d_1430_cf25_ = d_1429___mcc_h8_
                d_1431_cf24_ = d_1428___mcc_h7_
                d_1432_cf23_ = d_1427___mcc_h6_
                d_1432_cf23_ = d_1432_cf23_
                index210_ = default__.safeIndex(378, (d_1415_v57_).length(0))
                (d_1415_v57_)[index210_] = p0
                d_1433_v65_: _dafny.Array
                nw202_ = _dafny.Array(int(0), 5)
                d_1433_v65_ = nw202_
                rhs235_ = d_1433_v65_
                rhs236_ = (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]
                lhs201_ = globalState
                d_1433_v65_ = rhs235_
                lhs201_.f10 = rhs236_
                index211_ = default__.safeIndex(367, (d_1433_v65_).length(0))
                (d_1433_v65_)[index211_] = default__.safeModuloInt((d_1419_v59_).f30, (0) - ((d_1419_v59_).f30))
                index212_ = default__.safeIndex(367, (d_1433_v65_).length(0))
                (d_1433_v65_)[index212_] = (d_1419_v59_).f30
            elif source22_.is_DC15:
                d_1434___mcc_h9_ = source22_.cf26
                d_1435___mcc_h10_ = source22_.cf27
                d_1436_cf27_ = d_1435___mcc_h10_
                d_1437_cf26_ = d_1434___mcc_h9_
                d_1438_v66_: bool
                d_1439_v67_: int
                out58_: bool
                out59_: int
                out58_, out59_ = (self).m1(globalState)
                d_1438_v66_ = out58_
                d_1439_v67_ = out59_
                d_1440_v68_: C2
                nw203_ = C2()
                nw203_.ctor__((d_1419_v59_).f30, (self).f22)
                d_1440_v68_ = nw203_
                (globalState).f6 = d_1406_v50_
                d_1441_v69_: _dafny.Array
                def lambda64_(d_1442_v59_):
                    def lambda65_(d_1443_i8_):
                        return (d_1443_i8_) - ((d_1442_v59_).f30)

                    return lambda65_

                init33_ = lambda64_(d_1419_v59_)
                nw204_ = _dafny.Array(None, 14)
                for i0_33_ in range(nw204_.length(0)):
                    nw204_[i0_33_] = init33_(i0_33_)
                d_1441_v69_ = nw204_
                d_1444_v70_: _dafny.Seq
                d_1444_v70_ = _dafny.SeqWithoutIsStrInference([-441])
                index213_ = default__.safeIndex(285, (d_1441_v69_).length(0))
                (d_1441_v69_)[index213_] = default__.safeDivisionInt(d_1406_v50_, (d_1444_v70_)[default__.safeIndex((d_1440_v68_).f29, len(d_1444_v70_))])
                index214_ = default__.safeIndex(285, (d_1441_v69_).length(0))
                (d_1441_v69_)[index214_] = (d_1444_v70_)[default__.safeIndex((d_1419_v59_).f30, len(d_1444_v70_))]
            elif source22_.is_DC12:
                d_1445___mcc_h11_ = source22_.cf21
                d_1446_cf21_ = d_1445___mcc_h11_
                (globalState).f16 = (d_1419_v59_).f30
                d_1447_v71_: C3
                nw205_ = C3()
                nw205_.ctor__((self).f22)
                d_1447_v71_ = nw205_
                d_1448_v72_: _dafny.Array
                nw206_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_1448_v72_ = nw206_
                d_1449_v73_: _dafny.Array
                nw207_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_1449_v73_ = nw207_
                index215_ = default__.safeIndex(394, (d_1448_v72_).length(0))
                (d_1448_v72_)[index215_] = d_1449_v73_
                index216_ = default__.safeIndex(394, (d_1448_v72_).length(0))
                (d_1448_v72_)[index216_] = d_1449_v73_
                d_1450_v74_: D4
                d_1450_v74_ = D4_DC14(d_1418_v58_, d_1415_v57_, _dafny.SeqWithoutIsStrInference([d_1418_v58_ for d_1451_i9_ in range(default__.abs(778))]))
                pat_let_tv56_ = d_1410_v52_
                def iife113_(_pat_let26_0):
                    def iife114_(d_1452_dt__update__tmp_h0_):
                        def iife115_(_pat_let27_0):
                            def iife116_(d_1453_dt__update_hcf25_h0_):
                                return D4_DC14((d_1452_dt__update__tmp_h0_).cf23, (d_1452_dt__update__tmp_h0_).cf24, d_1453_dt__update_hcf25_h0_)
                            return iife116_(_pat_let27_0)
                        return iife115_(pat_let_tv56_)
                    return iife114_(_pat_let26_0)
                rhs237_ = d_1406_v50_
                rhs238_ = (iife113_(d_1450_v74_)).cf24
                rhs239_ = (d_1419_v59_).f30
                lhs202_ = globalState
                r3 = rhs237_
                r1 = rhs238_
                lhs202_.f11 = rhs239_
            elif True:
                d_1454___mcc_h12_ = source22_.cf28
                d_1455_cf28_ = d_1454___mcc_h12_
                (globalState).f6 = (0) - ((d_1419_v59_).f30)
                d_1456_v75_: D14
                d_1456_v75_ = D14_DC47(d_1413_v55_)
                (globalState).f0 = ((d_1456_v75_).cf77).issubset((d_1413_v55_) - (d_1413_v55_))
                d_1457_v76_: _dafny.Array
                nw208_ = _dafny.Array(int(0), 15)
                d_1457_v76_ = nw208_
                d_1458_v77_: _dafny.Map
                d_1458_v77_ = _dafny.Map({d_1457_v76_: (d_1419_v59_).f30})
                d_1459_v78_: _dafny.Map
                d_1459_v78_ = _dafny.Map({p0: d_1406_v50_})
                d_1460_v79_: _dafny.Map
                d_1460_v79_ = _dafny.Map({d_1406_v50_: not((d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))])})
                d_1461_v80_: _dafny.Map
                d_1461_v80_ = _dafny.Map({d_1460_v79_: not((d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))])})
                d_1462_v81_: _dafny.Map
                d_1462_v81_ = _dafny.Map({p0: d_1413_v55_})
                d_1463_v82_: _dafny.Seq
                d_1463_v82_ = _dafny.SeqWithoutIsStrInference([d_1406_v50_, (d_1419_v59_).f30, (d_1419_v59_).f30, (d_1419_v59_).f30, len(d_1462_v81_)])
                d_1464_v83_: _dafny.Seq
                d_1464_v83_ = _dafny.SeqWithoutIsStrInference([d_1460_v79_])
                d_1465_v84_: _dafny.Array
                nw209_ = _dafny.Array(None, 24)
                nw209_[int(0)] = (d_1419_v59_).f30
                nw209_[int(1)] = (d_1419_v59_).fm35((0) - ((d_1419_v59_).f30), len(d_1458_v77_), p0, 665, globalState)
                nw209_[int(2)] = len(d_1459_v78_)
                nw209_[int(3)] = (d_1419_v59_).f30
                nw209_[int(4)] = (d_1419_v59_).f30
                nw209_[int(5)] = (d_1419_v59_).f30
                nw209_[int(6)] = (d_1419_v59_).f30
                nw209_[int(7)] = (d_1419_v59_).f30
                nw209_[int(8)] = (d_1419_v59_).f30
                nw209_[int(9)] = (d_1419_v59_).f30
                nw209_[int(10)] = (d_1419_v59_).f30
                nw209_[int(11)] = (d_1419_v59_).f30
                nw209_[int(12)] = d_1406_v50_
                nw209_[int(13)] = -563
                nw209_[int(14)] = (d_1419_v59_).f30
                nw209_[int(15)] = len(d_1461_v80_)
                nw209_[int(16)] = (d_1419_v59_).f30
                nw209_[int(17)] = (d_1463_v82_)[default__.safeIndex((0) - ((0) - ((d_1419_v59_).f30)), len(d_1463_v82_))]
                nw209_[int(18)] = (d_1419_v59_).f30
                nw209_[int(19)] = (d_1419_v59_).f30
                nw209_[int(20)] = (_dafny.MultiSet(d_1463_v82_)).cardinality
                nw209_[int(21)] = len(d_1464_v83_)
                nw209_[int(22)] = 624
                nw209_[int(23)] = (d_1419_v59_).f30
                d_1465_v84_ = nw209_
                d_1466_v85_: _dafny.Map
                d_1466_v85_ = _dafny.Map({d_1465_v84_: (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]})
                d_1467_v86_: D10
                d_1467_v86_ = D10_DC34(d_1418_v58_, (self).f24, (D8_DC29(not(p0), p0, _dafny.CodePoint('l'))).cf48, (self).f24, True)
                d_1468_v87_: T0
                nw210_ = C4()
                nw210_.ctor__((d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))], (self).f22)
                d_1468_v87_ = nw210_
                d_1469_v88_: _dafny.Set
                d_1469_v88_ = _dafny.Set({d_1457_v76_})
                d_1470_v89_: D13
                d_1470_v89_ = D13_DC46((self).f24, d_1468_v87_, d_1406_v50_, d_1469_v88_)
                d_1471_v90_: _dafny.Seq
                d_1471_v90_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1465_v84_: (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]}), d_1466_v85_])
                d_1472_v91_: _dafny.Array
                nw211_ = _dafny.Array(None, 17)
                nw211_[int(0)] = d_1466_v85_
                nw211_[int(1)] = _dafny.Map({d_1457_v76_: False})
                nw211_[int(2)] = (_dafny.Map({d_1457_v76_: default__.fm15((d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))], globalState)})) | (d_1466_v85_)
                nw211_[int(3)] = d_1466_v85_
                nw211_[int(4)] = d_1466_v85_
                nw211_[int(5)] = d_1466_v85_
                nw211_[int(6)] = (_dafny.Map({d_1457_v76_: (d_1467_v86_).cf57})) | (d_1466_v85_)
                nw211_[int(7)] = _dafny.Map({d_1457_v76_: (self).f24})
                nw211_[int(8)] = d_1466_v85_
                nw211_[int(9)] = d_1466_v85_
                nw211_[int(10)] = d_1466_v85_
                nw211_[int(11)] = (_dafny.Map({d_1465_v84_: (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]}) if (d_1470_v89_).cf73 else (d_1471_v90_)[default__.safeIndex(len(d_1463_v82_), len(d_1471_v90_))])
                nw211_[int(12)] = (d_1466_v85_) | (d_1466_v85_)
                nw211_[int(13)] = d_1466_v85_
                nw211_[int(14)] = d_1466_v85_
                nw211_[int(15)] = d_1466_v85_
                nw211_[int(16)] = (d_1466_v85_).set(d_1457_v76_, p0)
                d_1472_v91_ = nw211_
                index217_ = default__.safeIndex(235, (d_1472_v91_).length(0))
                (d_1472_v91_)[index217_] = _dafny.Map({d_1465_v84_: (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]})
                index218_ = default__.safeIndex(235, (d_1472_v91_).length(0))
                (d_1472_v91_)[index218_] = d_1466_v85_
                d_1473_v92_: _dafny.Seq
                d_1473_v92_ = _dafny.SeqWithoutIsStrInference([d_1457_v76_])
                d_1457_v76_ = (d_1473_v92_)[default__.safeIndex(d_1406_v50_, len(d_1473_v92_))]
            d_1474_v93_: _dafny.Seq
            d_1474_v93_ = _dafny.SeqWithoutIsStrInference([(d_1419_v59_).f30])
            d_1411_v53_ = (d_1411_v53_).set(not(((d_1419_v59_).f30) >= ((d_1474_v93_)[default__.safeIndex(len(d_1474_v93_), len(d_1474_v93_))])), (d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))])
            d_1475_v94_: _dafny.Map
            d_1475_v94_ = _dafny.Map({(d_1415_v57_)[default__.safeIndex(378, (d_1415_v57_).length(0))]: (d_1419_v59_).f30})
            d_1476_v95_: D9
            d_1476_v95_ = D9_DC30((D9_DC30(d_1475_v94_)).cf50)
            d_1476_v95_ = d_1476_v95_
        elif True:
            d_1477_v96_: _dafny.Seq
            d_1477_v96_ = _dafny.SeqWithoutIsStrInference([(d_1406_v50_) != (d_1406_v50_)])
            if (d_1477_v96_)[default__.safeIndex(d_1406_v50_, len(d_1477_v96_))]:
                def lambda66_(d_1478_i10_):
                    return (self).f24

                init34_ = lambda66_
                nw212_ = _dafny.Array(None, 16)
                for i0_34_ in range(nw212_.length(0)):
                    nw212_[i0_34_] = init34_(i0_34_)
                r1 = nw212_
                d_1479_v97_: C1
                nw213_ = C1()
                nw213_.ctor__(d_1406_v50_, (self).f22)
                d_1479_v97_ = nw213_
                d_1480_v98_: _dafny.Array
                def lambda67_(d_1481_v97_):
                    def lambda68_(d_1482_i11_):
                        return (d_1482_i11_) + ((d_1481_v97_).f30)

                    return lambda68_

                init35_ = lambda67_(d_1479_v97_)
                nw214_ = _dafny.Array(None, 24)
                for i0_35_ in range(nw214_.length(0)):
                    nw214_[i0_35_] = init35_(i0_35_)
                d_1480_v98_ = nw214_
                index219_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                (d_1480_v98_)[index219_] = (d_1479_v97_).f30
                index220_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                (d_1480_v98_)[index220_] = default__.fm7((d_1479_v97_).f30, globalState)
                d_1483_v99_: str
                d_1483_v99_ = _dafny.CodePoint('y')
                index221_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                index222_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                rhs240_ = (d_1479_v97_).f30
                rhs241_ = default__.fm38(True, d_1406_v50_, globalState)
                rhs242_ = default__.safeModuloInt(default__.safeModuloInt(d_1406_v50_, len(d_1477_v96_)), (d_1406_v50_) * ((d_1479_v97_).f30))
                rhs243_ = ((d_1479_v97_).f30) * (d_1406_v50_)
                rhs244_ = (d_1480_v98_)[default__.safeIndex(446, (d_1480_v98_).length(0))]
                lhs203_ = d_1480_v98_
                lhs204_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                lhs205_ = globalState
                lhs206_ = d_1480_v98_
                lhs207_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                lhs203_[lhs204_] = rhs240_
                d_1483_v99_ = rhs241_
                lhs205_.f6 = rhs242_
                r3 = rhs243_
                lhs206_[lhs207_] = rhs244_
                index223_ = default__.safeIndex(446, (d_1480_v98_).length(0))
                (d_1480_v98_)[index223_] = d_1406_v50_
            elif True:
                d_1484_v100_: D2
                d_1484_v100_ = D2_DC7(_dafny.SeqWithoutIsStrInference([d_1406_v50_]), (self).f24, not((self).f24), d_1406_v50_)
                d_1485_v101_: _dafny.Set
                d_1485_v101_ = _dafny.Set({d_1484_v100_, d_1484_v100_, d_1484_v100_})
                (globalState).f10 = (d_1485_v101_) == (d_1485_v101_)
                d_1486_v102_: _dafny.Array
                def lambda69_(d_1487_p0_):
                    def lambda70_(d_1488_i12_):
                        return d_1487_p0_

                    return lambda70_

                init36_ = lambda69_(p0)
                nw215_ = _dafny.Array(None, 7)
                for i0_36_ in range(nw215_.length(0)):
                    nw215_[i0_36_] = init36_(i0_36_)
                d_1486_v102_ = nw215_
                d_1489_v103_: T0
                nw216_ = C2()
                nw216_.ctor__(d_1406_v50_, _dafny.SeqWithoutIsStrInference([d_1486_v102_]))
                d_1489_v103_ = nw216_
                d_1490_v104_: _dafny.Map
                d_1490_v104_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([d_1406_v50_ for d_1491_i13_ in range(default__.abs(584))]))})
                d_1492_v105_: _dafny.Array
                nw217_ = _dafny.Array(int(0), 9)
                d_1492_v105_ = nw217_
                d_1493_v106_: D13
                d_1493_v106_ = D13_DC46(p0, d_1489_v103_, len(d_1490_v104_), _dafny.Set({d_1492_v105_}))
                d_1494_v107_: D13
                d_1494_v107_ = D13_DC46((self).f24, (d_1493_v106_).cf74, d_1406_v50_, _dafny.Set({d_1492_v105_, d_1492_v105_, d_1492_v105_, d_1492_v105_}))
                d_1495_v108_: _dafny.MultiSet
                d_1495_v108_ = _dafny.MultiSet([d_1406_v50_])
                d_1496_v111_: _dafny.Set
                def iife117_():
                    def iife119_():
                        coll63_ = _dafny.Set()
                        compr_63_: int
                        for compr_63_ in (d_1495_v108_).Elements:
                            d_1499_v109_: int = compr_63_
                            if (d_1499_v109_) in (d_1495_v108_):
                                coll63_ = coll63_.union(_dafny.Set([(d_1499_v109_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eelqwwl"))))]))
                        return _dafny.Set(coll63_)
                    coll61_ = _dafny.Set()
                    def iife118_():
                        coll62_ = _dafny.Set()
                        compr_62_: int
                        for compr_62_ in (d_1495_v108_).Elements:
                            d_1497_v109_: int = compr_62_
                            if (d_1497_v109_) in (d_1495_v108_):
                                coll62_ = coll62_.union(_dafny.Set([(d_1497_v109_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eelqwwl"))))]))
                        return _dafny.Set(coll62_)
                    compr_61_: int
                    for compr_61_ in (iife118_()
                    ).Elements:
                        d_1498_v110_: int = compr_61_
                        if (d_1498_v110_) in (iife119_()
                        ):
                            coll61_ = coll61_.union(_dafny.Set([(d_1498_v110_) * (274)]))
                    return _dafny.Set(coll61_)
                d_1496_v111_ = _dafny.Set({iife117_()
                })
                d_1500_v112_: _dafny.Set
                d_1500_v112_ = _dafny.Set({d_1492_v105_, d_1492_v105_, d_1492_v105_})
                d_1501_v113_: _dafny.Map
                d_1501_v113_ = _dafny.Map({D13_DC46(p0, d_1489_v103_, len(_dafny.Map({(self).f24: d_1406_v50_})), d_1500_v112_): 718})
                d_1502_v114_: _dafny.Seq
                d_1502_v114_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1494_v107_: len(d_1496_v111_)}), (d_1501_v113_).set(d_1494_v107_, -263)])
                (globalState).f11 = len(d_1502_v114_)
                d_1503_v115_: _dafny.Seq
                d_1503_v115_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phpv"))
                d_1503_v115_ = d_1503_v115_
                d_1504_v116_: str
                d_1504_v116_ = _dafny.CodePoint('y')
                d_1505_v117_: _dafny.Map
                d_1505_v117_ = _dafny.Map({d_1406_v50_: d_1504_v116_})
                d_1506_v119_: _dafny.Map
                def iife120_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(667, 816):
                        d_1507_v118_: int = compr_64_
                        if ((667) <= (d_1507_v118_)) and ((d_1507_v118_) < (816)):
                            coll64_[(d_1507_v118_) * (len(d_1490_v104_))] = d_1504_v116_
                    return _dafny.Map(coll64_)
                d_1506_v119_ = _dafny.Map({d_1406_v50_: (d_1505_v117_) != ((iife120_()
                ).set(d_1406_v50_, _dafny.CodePoint('n')))})
                d_1508_v120_: _dafny.Seq
                d_1508_v120_ = _dafny.SeqWithoutIsStrInference([d_1406_v50_])
                d_1509_v121_: D13
                d_1509_v121_ = D13_DC45(len(d_1508_v120_))
                pat_let_tv57_ = d_1406_v50_
                def iife121_(_pat_let28_0):
                    def iife122_(d_1510_dt__update__tmp_h1_):
                        def iife123_(_pat_let29_0):
                            def iife124_(d_1511_dt__update_hcf72_h0_):
                                return D13_DC45(d_1511_dt__update_hcf72_h0_)
                            return iife124_(_pat_let29_0)
                        return iife123_(pat_let_tv57_)
                    return iife122_(_pat_let28_0)
                d_1506_v119_ = (d_1506_v119_).set(d_1406_v50_, (d_1406_v50_) == ((iife121_(d_1509_v121_)).cf72))
                (globalState).f16 = d_1406_v50_
            d_1512_v122_: D10
            d_1512_v122_ = D10_DC35((0) - (d_1406_v50_), d_1406_v50_)
            source23_ = d_1512_v122_
            if source23_.is_DC34:
                d_1513___mcc_h13_ = source23_.cf55
                d_1514___mcc_h14_ = source23_.cf56
                d_1515___mcc_h15_ = source23_.cf57
                d_1516___mcc_h16_ = source23_.cf58
                d_1517___mcc_h17_ = source23_.cf59
                d_1518_cf59_ = d_1517___mcc_h17_
                d_1519_cf58_ = d_1516___mcc_h16_
                d_1520_cf57_ = d_1515___mcc_h15_
                d_1521_cf56_ = d_1514___mcc_h14_
                d_1522_cf55_ = d_1513___mcc_h13_
                d_1523_v123_: _dafny.Seq
                d_1523_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
                d_1524_v124_: _dafny.Map
                d_1524_v124_ = _dafny.Map({((d_1477_v96_).set(default__.safeIndex(d_1406_v50_, len(d_1477_v96_)), d_1518_cf59_)) != (d_1477_v96_): (d_1406_v50_) + ((0) - (len(d_1523_v123_)))})
                rhs245_ = (0) - ((0) - ((0) - (default__.safeDivisionInt((0) - (d_1406_v50_), d_1406_v50_))))
                rhs246_ = not((d_1523_v123_) <= (d_1523_v123_))
                rhs247_ = ((d_1524_v124_) | (d_1524_v124_)) | (d_1524_v124_)
                lhs208_ = globalState
                lhs208_.f8 = rhs245_
                d_1519_cf58_ = rhs246_
                d_1524_v124_ = rhs247_
                d_1525_v125_: _dafny.Array
                def lambda71_(d_1526_v50_):
                    def lambda72_(d_1527_i14_):
                        return default__.safeModuloInt(d_1527_i14_, d_1526_v50_)

                    return lambda72_

                init37_ = lambda71_(d_1406_v50_)
                nw218_ = _dafny.Array(None, 7)
                for i0_37_ in range(nw218_.length(0)):
                    nw218_[i0_37_] = init37_(i0_37_)
                d_1525_v125_ = nw218_
                nw219_ = _dafny.Array(None, 1)
                nw219_[int(0)] = -204
                d_1525_v125_ = nw219_
                (globalState).f16 = d_1406_v50_
                r3 = d_1406_v50_
            elif source23_.is_DC35:
                d_1528___mcc_h18_ = source23_.cf60
                d_1529___mcc_h19_ = source23_.cf61
                d_1530_cf61_ = d_1529___mcc_h19_
                d_1531_cf60_ = d_1528___mcc_h18_
                d_1532_v126_: str
                d_1532_v126_ = _dafny.CodePoint('l')
                d_1533_v127_: _dafny.Seq
                d_1533_v127_ = _dafny.SeqWithoutIsStrInference([d_1406_v50_])
                d_1534_v128_: D0
                d_1534_v128_ = D0_DC1(d_1530_cf61_, p0, default__.fm47(d_1532_v126_, False, globalState), len(d_1533_v127_), d_1532_v126_)
                d_1535_v129_: D4
                d_1535_v129_ = D4_DC13(d_1534_v128_)
                d_1536_v130_: _dafny.Seq
                d_1536_v130_ = _dafny.SeqWithoutIsStrInference([d_1535_v129_])
                d_1537_v131_: _dafny.Map
                d_1537_v131_ = _dafny.Map({len(d_1536_v130_): p0})
                d_1538_v132_: _dafny.Seq
                d_1538_v132_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmopkrvet"))
                rhs248_ = (not(p0)) or (default__.fm15((self).f24, globalState))
                rhs249_ = default__.fm37(_dafny.SeqWithoutIsStrInference([p0, p0, p0]), d_1537_v131_, d_1538_v132_, globalState)
                rhs250_ = d_1530_cf61_
                lhs209_ = globalState
                lhs210_ = globalState
                r0 = rhs248_
                lhs209_.f0 = rhs249_
                lhs210_.f16 = rhs250_
                d_1539_v133_: _dafny.Map
                d_1539_v133_ = _dafny.Map({p0: len(d_1538_v132_)})
                d_1540_v134_: D9
                d_1540_v134_ = D9_DC30(d_1539_v133_)
                d_1541_v135_: _dafny.Seq
                d_1541_v135_ = _dafny.SeqWithoutIsStrInference([d_1539_v133_, _dafny.Map({(self).f24: len(d_1533_v127_)})])
                pat_let_tv58_ = d_1541_v135_
                pat_let_tv59_ = d_1530_cf61_
                pat_let_tv60_ = d_1541_v135_
                d_1542_v136_: _dafny.Map
                def iife125_(_pat_let30_0):
                    def iife126_(d_1543_dt__update__tmp_h2_):
                        def iife127_(_pat_let31_0):
                            def iife128_(d_1544_dt__update_hcf50_h0_):
                                return D9_DC30(d_1544_dt__update_hcf50_h0_)
                            return iife128_(_pat_let31_0)
                        return iife127_((pat_let_tv58_)[default__.safeIndex(pat_let_tv59_, len(pat_let_tv60_))])
                    return iife126_(_pat_let30_0)
                d_1542_v136_ = _dafny.Map({d_1538_v132_: iife125_(d_1540_v134_)})
                d_1542_v136_ = d_1542_v136_
                d_1545_v137_: _dafny.Array
                nw220_ = _dafny.Array(False, 11)
                d_1545_v137_ = nw220_
                r1 = d_1545_v137_
                d_1546_v138_: _dafny.Array
                def lambda73_(d_1547_cf61_):
                    def lambda74_(d_1548_i15_):
                        return (d_1548_i15_) * ((0) - (len(_dafny.Map({d_1547_cf61_: 43}))))

                    return lambda74_

                init38_ = lambda73_(d_1530_cf61_)
                nw221_ = _dafny.Array(None, 3)
                for i0_38_ in range(nw221_.length(0)):
                    nw221_[i0_38_] = init38_(i0_38_)
                d_1546_v138_ = nw221_
                index224_ = default__.safeIndex(516, (d_1546_v138_).length(0))
                (d_1546_v138_)[index224_] = -623
                d_1549_v139_: _dafny.MultiSet
                d_1549_v139_ = _dafny.MultiSet([d_1406_v50_])
                d_1550_v140_: _dafny.MultiSet
                d_1550_v140_ = _dafny.MultiSet([(self).f24, (self).f24, (self).f24, p0, p0])
                d_1551_v141_: _dafny.Set
                d_1551_v141_ = _dafny.Set({726, -786})
                d_1552_v142_: _dafny.Seq
                d_1552_v142_ = _dafny.SeqWithoutIsStrInference([default__.fm23(d_1406_v50_, d_1406_v50_, p0, p0, globalState), d_1551_v141_])
                index225_ = default__.safeIndex(516, (d_1546_v138_).length(0))
                rhs251_ = ((d_1538_v132_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))) + ((d_1538_v132_) + (d_1538_v132_))
                rhs252_ = (self).f24
                rhs253_ = (self).fm0(d_1549_v139_, (d_1406_v50_ if p0 else (0) - ((d_1550_v140_).cardinality)), d_1552_v142_, (self).f24, globalState)
                rhs254_ = (d_1533_v127_)[default__.safeIndex((0) - (d_1530_cf61_), len(d_1533_v127_))]
                lhs211_ = globalState
                lhs212_ = d_1546_v138_
                lhs213_ = default__.safeIndex(516, (d_1546_v138_).length(0))
                d_1538_v132_ = rhs251_
                lhs211_.f10 = rhs252_
                r3 = rhs253_
                lhs212_[lhs213_] = rhs254_
            elif True:
                d_1553___mcc_h20_ = source23_.cf54
                d_1554_cf54_ = d_1553___mcc_h20_
                d_1555_v143_: _dafny.Map
                d_1555_v143_ = _dafny.Map({p0: p0})
                d_1556_v144_: _dafny.Map
                d_1556_v144_ = _dafny.Map({d_1340_v1_: d_1555_v143_})
                d_1557_v145_: _dafny.Seq
                d_1557_v145_ = _dafny.SeqWithoutIsStrInference([d_1556_v144_])
                d_1558_v146_: _dafny.Seq
                d_1558_v146_ = _dafny.SeqWithoutIsStrInference([d_1406_v50_, 306])
                d_1556_v144_ = (d_1557_v145_)[default__.safeIndex((d_1558_v146_)[default__.safeIndex((0) - (d_1406_v50_), len(d_1558_v146_))], len(d_1557_v145_))]
                (globalState).f11 = d_1406_v50_
                d_1559_v147_: _dafny.MultiSet
                d_1559_v147_ = _dafny.MultiSet([d_1406_v50_])
                (globalState).f0 = (_dafny.MultiSet([d_1406_v50_])).ispropersubset(d_1559_v147_)
                d_1560_v148_: _dafny.Array
                nw222_ = _dafny.Array(False, 2)
                d_1560_v148_ = nw222_
                d_1561_v149_: str
                d_1561_v149_ = _dafny.CodePoint('s')
                d_1562_v150_: _dafny.Map
                d_1562_v150_ = _dafny.Map({d_1560_v148_: d_1561_v149_})
                d_1562_v150_ = (d_1562_v150_).set(d_1560_v148_, d_1561_v149_)
            d_1563_v151_: _dafny.Array
            nw223_ = _dafny.Array(_dafny.Map({}), 3)
            d_1563_v151_ = nw223_
            d_1564_v152_: _dafny.Array
            nw224_ = _dafny.Array(_dafny.Set({}), 10)
            d_1564_v152_ = nw224_
            d_1565_v153_: _dafny.Map
            d_1565_v153_ = _dafny.Map({d_1406_v50_: p0})
            d_1566_v154_: _dafny.Seq
            d_1566_v154_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lntn"))
            d_1567_v155_: _dafny.Map
            d_1567_v155_ = _dafny.Map({d_1406_v50_: default__.fm37(d_1477_v96_, d_1565_v153_, d_1566_v154_, globalState)})
            d_1568_v156_: _dafny.Seq
            d_1568_v156_ = _dafny.SeqWithoutIsStrInference([d_1567_v155_])
            rhs255_ = not(p0)
            rhs256_ = d_1563_v151_
            rhs257_ = d_1564_v152_
            rhs258_ = d_1568_v156_
            rhs259_ = p0
            lhs214_ = globalState
            lhs215_ = globalState
            lhs214_.f10 = rhs255_
            d_1563_v151_ = rhs256_
            d_1564_v152_ = rhs257_
            d_1568_v156_ = rhs258_
            lhs215_.f10 = rhs259_
            d_1569_v157_: _dafny.MultiSet
            d_1569_v157_ = _dafny.MultiSet([d_1566_v154_])
            d_1570_v158_: D10
            d_1570_v158_ = D10_DC34(default__.fm38(p0, d_1406_v50_, globalState), p0, p0, not((d_1569_v157_).issubset(d_1569_v157_)), p0)
            source24_ = d_1570_v158_
            if source24_.is_DC34:
                d_1571___mcc_h21_ = source24_.cf55
                d_1572___mcc_h22_ = source24_.cf56
                d_1573___mcc_h23_ = source24_.cf57
                d_1574___mcc_h24_ = source24_.cf58
                d_1575___mcc_h25_ = source24_.cf59
                d_1576_cf59_ = d_1575___mcc_h25_
                d_1577_cf58_ = d_1574___mcc_h24_
                d_1578_cf57_ = d_1573___mcc_h23_
                d_1579_cf56_ = d_1572___mcc_h22_
                d_1580_cf55_ = d_1571___mcc_h21_
                d_1581_v159_: C0
                nw225_ = C0()
                nw225_.ctor__()
                d_1581_v159_ = nw225_
                d_1582_v160_: C3
                nw226_ = C3()
                nw226_.ctor__((self).f22)
                d_1582_v160_ = nw226_
                d_1583_v161_: _dafny.MultiSet
                d_1583_v161_ = _dafny.MultiSet([d_1582_v160_])
                d_1584_v162_: _dafny.Map
                d_1584_v162_ = _dafny.Map({d_1406_v50_: d_1583_v161_})
                d_1585_v163_: C1
                nw227_ = C1()
                nw227_.ctor__(len(d_1584_v162_), ((self).f22) + ((self).f22))
                d_1585_v163_ = nw227_
                d_1586_v164_: _dafny.Map
                d_1586_v164_ = _dafny.Map({d_1576_cf59_: -861})
                d_1587_v165_: _dafny.Map
                d_1587_v165_ = _dafny.Map({d_1576_cf59_: d_1582_v160_})
                d_1588_v166_: _dafny.Set
                d_1588_v166_ = _dafny.Set({d_1587_v165_, d_1587_v165_, d_1587_v165_, (d_1587_v165_).set(p0, d_1582_v160_), (d_1587_v165_).set(d_1576_cf59_, d_1582_v160_)})
                d_1589_v167_: _dafny.Seq
                d_1589_v167_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1576_cf59_])])
                rhs260_ = (0) - (d_1406_v50_)
                rhs261_ = (d_1585_v163_).f30
                rhs262_ = d_1586_v164_
                rhs263_ = _dafny.Set({(d_1587_v165_) | (d_1587_v165_), d_1587_v165_, d_1587_v165_})
                rhs264_ = d_1589_v167_
                lhs216_ = globalState
                d_1406_v50_ = rhs260_
                lhs216_.f16 = rhs261_
                d_1586_v164_ = rhs262_
                d_1588_v166_ = rhs263_
                d_1589_v167_ = rhs264_
                (globalState).f16 = (0) - ((0) - ((d_1585_v163_).fm34(True, globalState)))
            elif source24_.is_DC35:
                d_1590___mcc_h26_ = source24_.cf60
                d_1591___mcc_h27_ = source24_.cf61
                d_1592_cf61_ = d_1591___mcc_h27_
                d_1593_cf60_ = d_1590___mcc_h26_
                d_1594_v168_: _dafny.Array
                nw228_ = _dafny.Array(int(0), 7)
                d_1594_v168_ = nw228_
                d_1595_v169_: _dafny.Array
                nw229_ = _dafny.Array(False, 17)
                d_1595_v169_ = nw229_
                index226_ = default__.safeIndex(157, (d_1595_v169_).length(0))
                (d_1595_v169_)[index226_] = (self).f24
                d_1596_v170_: _dafny.Set
                d_1596_v170_ = _dafny.Set({p0})
                d_1597_v171_: _dafny.MultiSet
                d_1597_v171_ = _dafny.MultiSet([d_1596_v170_, d_1596_v170_])
                index227_ = default__.safeIndex(157, (d_1595_v169_).length(0))
                rhs265_ = d_1594_v168_
                rhs266_ = d_1406_v50_
                rhs267_ = (d_1597_v171_).issubset(d_1597_v171_)
                rhs268_ = not((self).f24)
                lhs217_ = globalState
                lhs218_ = globalState
                lhs219_ = d_1595_v169_
                lhs220_ = default__.safeIndex(157, (d_1595_v169_).length(0))
                d_1594_v168_ = rhs265_
                lhs217_.f16 = rhs266_
                lhs218_.f10 = rhs267_
                lhs219_[lhs220_] = rhs268_
                (globalState).f5 = ((d_1477_v96_) + (d_1477_v96_)) + ((_dafny.SeqWithoutIsStrInference([(d_1595_v169_)[default__.safeIndex(157, (d_1595_v169_).length(0))], p0, (d_1595_v169_)[default__.safeIndex(157, (d_1595_v169_).length(0))], (self).f24, (self).f24])) + (d_1477_v96_))
                d_1598_v172_: _dafny.Seq
                d_1598_v172_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([default__.fm15((self).f24, globalState)])])
                d_1599_v173_: _dafny.MultiSet
                d_1599_v173_ = _dafny.MultiSet([p0])
                d_1598_v172_ = (((d_1598_v172_) + (default__.fm48(globalState))) + (_dafny.SeqWithoutIsStrInference([d_1477_v96_]))).set(default__.safeIndex((d_1406_v50_) * ((d_1599_v173_).cardinality), len(((d_1598_v172_) + (default__.fm48(globalState))) + (_dafny.SeqWithoutIsStrInference([d_1477_v96_])))), d_1477_v96_)
                d_1600_v174_: _dafny.Seq
                d_1600_v174_ = _dafny.SeqWithoutIsStrInference([d_1592_cf61_])
                d_1600_v174_ = d_1600_v174_
            elif True:
                d_1601___mcc_h28_ = source24_.cf54
                d_1602_cf54_ = d_1601___mcc_h28_
                rhs269_ = not ((p0 if p0 else p0)) or ((self).fm3(d_1566_v154_, p0, (d_1477_v96_)[default__.safeIndex(124, len(d_1477_v96_))], globalState))
                rhs270_ = (d_1477_v96_) <= (d_1477_v96_)
                lhs221_ = globalState
                r0 = rhs269_
                lhs221_.f10 = rhs270_
                d_1603_v176_: _dafny.Array
                def lambda75_(d_1604_v50_):
                    def lambda76_(d_1605_i16_):
                        def iife129_():
                            coll65_ = _dafny.Map()
                            compr_65_: int
                            for compr_65_ in _dafny.IntegerRange(205, 829):
                                d_1606_v175_: int = compr_65_
                                if ((205) <= (d_1606_v175_)) and ((d_1606_v175_) < (829)):
                                    coll65_[(d_1606_v175_) * (d_1604_v50_)] = d_1604_v50_
                            return _dafny.Map(coll65_)
                        return (D5_DC19(iife129_()
, True)).cf33

                    return lambda76_

                init39_ = lambda75_(d_1406_v50_)
                nw230_ = _dafny.Array(None, 17)
                for i0_39_ in range(nw230_.length(0)):
                    nw230_[i0_39_] = init39_(i0_39_)
                d_1603_v176_ = nw230_
                d_1607_v177_: _dafny.Map
                d_1607_v177_ = _dafny.Map({d_1406_v50_: d_1603_v176_})
                d_1608_v178_: _dafny.Seq
                d_1608_v178_ = _dafny.SeqWithoutIsStrInference([len(d_1607_v177_)])
                d_1609_v179_: _dafny.Map
                d_1609_v179_ = _dafny.Map({p0: d_1406_v50_})
                d_1610_v180_: D2
                d_1610_v180_ = D2_DC7(d_1608_v178_, (self).f24, True, len(d_1609_v179_))
                index228_ = default__.safeIndex(360, (d_1603_v176_).length(0))
                (d_1603_v176_)[index228_] = (d_1610_v180_).cf11
                index229_ = default__.safeIndex(360, (d_1603_v176_).length(0))
                (d_1603_v176_)[index229_] = (d_1477_v96_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "on")) if False else d_1566_v154_)), len(d_1477_v96_))]
                d_1566_v154_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrcsc"))
                d_1611_v181_: _dafny.Map
                d_1611_v181_ = _dafny.Map({len(d_1566_v154_): -637})
                d_1612_v182_: _dafny.Array
                nw231_ = _dafny.Array(None, 20)
                nw231_[int(0)] = d_1406_v50_
                nw231_[int(1)] = d_1406_v50_
                nw231_[int(2)] = len(d_1566_v154_)
                nw231_[int(3)] = d_1406_v50_
                nw231_[int(4)] = d_1406_v50_
                nw231_[int(5)] = len(d_1566_v154_)
                nw231_[int(6)] = d_1406_v50_
                nw231_[int(7)] = d_1406_v50_
                nw231_[int(8)] = d_1406_v50_
                nw231_[int(9)] = d_1406_v50_
                nw231_[int(10)] = 605
                nw231_[int(11)] = 413
                nw231_[int(12)] = ((d_1611_v181_)[default__.fm7(d_1406_v50_, globalState)] if (default__.fm7(d_1406_v50_, globalState)) in (d_1611_v181_) else d_1406_v50_)
                nw231_[int(13)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbq")))
                nw231_[int(14)] = d_1406_v50_
                nw231_[int(15)] = d_1406_v50_
                nw231_[int(16)] = d_1406_v50_
                nw231_[int(17)] = len(d_1566_v154_)
                nw231_[int(18)] = 402
                nw231_[int(19)] = -262
                d_1612_v182_ = nw231_
                d_1613_v183_: _dafny.Seq
                d_1613_v183_ = _dafny.SeqWithoutIsStrInference([d_1612_v182_])
                d_1614_v184_: _dafny.Array
                nw232_ = _dafny.Array(None, 29)
                nw232_[int(0)] = (d_1613_v183_)[default__.safeIndex((0) - (d_1406_v50_), len(d_1613_v183_))]
                nw232_[int(1)] = d_1612_v182_
                nw232_[int(2)] = (d_1612_v182_ if (d_1603_v176_)[default__.safeIndex(360, (d_1603_v176_).length(0))] else d_1612_v182_)
                nw232_[int(3)] = d_1612_v182_
                nw232_[int(4)] = d_1612_v182_
                nw232_[int(5)] = d_1612_v182_
                nw232_[int(6)] = d_1612_v182_
                nw232_[int(7)] = d_1612_v182_
                nw232_[int(8)] = d_1612_v182_
                nw232_[int(9)] = d_1612_v182_
                nw232_[int(10)] = d_1612_v182_
                nw232_[int(11)] = d_1612_v182_
                nw232_[int(12)] = d_1612_v182_
                nw232_[int(13)] = d_1612_v182_
                nw232_[int(14)] = d_1612_v182_
                nw232_[int(15)] = d_1612_v182_
                nw232_[int(16)] = d_1612_v182_
                nw232_[int(17)] = d_1612_v182_
                nw232_[int(18)] = d_1612_v182_
                nw232_[int(19)] = d_1612_v182_
                nw232_[int(20)] = d_1612_v182_
                nw232_[int(21)] = d_1612_v182_
                nw232_[int(22)] = d_1612_v182_
                nw232_[int(23)] = d_1612_v182_
                nw232_[int(24)] = d_1612_v182_
                nw232_[int(25)] = d_1612_v182_
                nw232_[int(26)] = d_1612_v182_
                nw232_[int(27)] = d_1612_v182_
                nw232_[int(28)] = d_1612_v182_
                d_1614_v184_ = nw232_
                index230_ = default__.safeIndex(50, (d_1614_v184_).length(0))
                (d_1614_v184_)[index230_] = d_1612_v182_
                index231_ = default__.safeIndex(50, (d_1614_v184_).length(0))
                (d_1614_v184_)[index231_] = d_1612_v182_
            d_1615_v185_: str
            d_1615_v185_ = _dafny.CodePoint('h')
            d_1616_v186_: _dafny.Map
            d_1616_v186_ = _dafny.Map({d_1615_v185_: p0})
            rhs271_ = ((default__.fm49(p0, globalState)).set(d_1615_v185_, False)) | (d_1616_v186_)
            rhs272_ = ((d_1477_v96_) + (d_1477_v96_)) <= (d_1477_v96_)
            lhs222_ = globalState
            d_1616_v186_ = rhs271_
            lhs222_.f0 = rhs272_
        d_1617_v187_: _dafny.Seq
        d_1617_v187_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emlcy"))
        d_1617_v187_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usnjcg"))
        if not(not((d_1617_v187_) != (d_1617_v187_))):
            d_1617_v187_ = default__.fm30(globalState)
            d_1618_v188_: _dafny.MultiSet
            d_1618_v188_ = _dafny.MultiSet([len(d_1617_v187_)])
            d_1618_v188_ = (_dafny.MultiSet([d_1406_v50_, d_1406_v50_, (0) - (default__.fm7(d_1406_v50_, globalState)), d_1406_v50_, d_1406_v50_])) - ((d_1618_v188_).intersection(d_1618_v188_))
            d_1619_v189_: C4
            nw233_ = C4()
            nw233_.ctor__(p0, (self).f22)
            d_1619_v189_ = nw233_
            r0 = (d_1617_v187_) < ((d_1617_v187_) + (d_1617_v187_))
            d_1620_v190_: _dafny.Map
            d_1620_v190_ = _dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('x')]))})
            d_1621_v191_: _dafny.Set
            d_1621_v191_ = _dafny.Set({not(p0)})
            d_1622_v192_: _dafny.Set
            d_1622_v192_ = _dafny.Set({d_1406_v50_, len(d_1620_v190_), (0) - (len(d_1621_v191_))})
            d_1623_v194_: _dafny.MultiSet
            def iife130_():
                coll66_ = _dafny.Set()
                compr_66_: int
                for compr_66_ in _dafny.IntegerRange(965, 848):
                    d_1624_v193_: int = compr_66_
                    if ((965) <= (d_1624_v193_)) and ((d_1624_v193_) < (848)):
                        coll66_ = coll66_.union(_dafny.Set([default__.safeModuloInt(d_1624_v193_, d_1406_v50_)]))
                return _dafny.Set(coll66_)
            d_1623_v194_ = _dafny.MultiSet([(d_1622_v192_ if p0 else iife130_()
            ), d_1622_v192_, d_1622_v192_, (d_1622_v192_ if p0 else d_1622_v192_)])
            index232_ = default__.safeIndex(383, (d_1337_v0_).length(0))
            (d_1337_v0_)[index232_] = d_1617_v187_
            d_1625_v196_: _dafny.Map
            def iife131_():
                coll67_ = _dafny.Set()
                compr_67_: int
                for compr_67_ in _dafny.IntegerRange(377, -477):
                    d_1626_v195_: int = compr_67_
                    if ((377) <= (d_1626_v195_)) and ((d_1626_v195_) < (-477)):
                        coll67_ = coll67_.union(_dafny.Set([default__.safeDivisionInt(d_1626_v195_, d_1406_v50_)]))
                return _dafny.Set(coll67_)
            d_1625_v196_ = _dafny.Map({d_1617_v187_: len(iife131_()
            )})
            index233_ = default__.safeIndex(383, (d_1337_v0_).length(0))
            rhs273_ = d_1623_v194_
            rhs274_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1627_i17_ in range(default__.abs(610))])
            rhs275_ = d_1619_v189_.f28
            rhs276_ = d_1617_v187_
            rhs277_ = (len((d_1625_v196_) | (d_1625_v196_)) if (self).f24 else (d_1406_v50_) * (841))
            lhs223_ = d_1337_v0_
            lhs224_ = default__.safeIndex(383, (d_1337_v0_).length(0))
            lhs225_ = globalState
            lhs226_ = globalState
            d_1623_v194_ = rhs273_
            lhs223_[lhs224_] = rhs274_
            lhs225_.f0 = rhs275_
            d_1617_v187_ = rhs276_
            lhs226_.f16 = rhs277_
        elif True:
            d_1628_v197_: str
            d_1628_v197_ = _dafny.CodePoint('d')
            if not((d_1628_v197_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "emuwsan")))):
                index234_ = default__.safeIndex(65, (d_1337_v0_).length(0))
                (d_1337_v0_)[index234_] = d_1617_v187_
                index235_ = default__.safeIndex(65, (d_1337_v0_).length(0))
                (d_1337_v0_)[index235_] = d_1617_v187_
                d_1629_v198_: _dafny.Array
                def lambda77_(d_1630_v50_):
                    def lambda78_(d_1631_i18_):
                        return default__.safeModuloInt(d_1631_i18_, d_1630_v50_)

                    return lambda78_

                init40_ = lambda77_(d_1406_v50_)
                nw234_ = _dafny.Array(None, 3)
                for i0_40_ in range(nw234_.length(0)):
                    nw234_[i0_40_] = init40_(i0_40_)
                d_1629_v198_ = nw234_
                d_1632_v199_: _dafny.Set
                d_1632_v199_ = _dafny.Set({-974})
                d_1633_v200_: _dafny.Map
                d_1633_v200_ = _dafny.Map({True: (0) - (len(d_1632_v199_))})
                index236_ = default__.safeIndex(409, (d_1629_v198_).length(0))
                (d_1629_v198_)[index236_] = len((d_1633_v200_) | (_dafny.Map({p0: -823})))
                d_1634_v201_: _dafny.Set
                d_1634_v201_ = _dafny.Set({p0, True})
                d_1635_v202_: _dafny.Seq
                d_1635_v202_ = _dafny.SeqWithoutIsStrInference([d_1406_v50_, len((d_1634_v201_) - (d_1634_v201_)), d_1406_v50_])
                d_1636_v203_: _dafny.Array
                def lambda79_(d_1637_v50_):
                    def lambda80_(d_1638_i19_):
                        return (417) >= (d_1637_v50_)

                    return lambda80_

                init41_ = lambda79_(d_1406_v50_)
                nw235_ = _dafny.Array(None, 26)
                for i0_41_ in range(nw235_.length(0)):
                    nw235_[i0_41_] = init41_(i0_41_)
                d_1636_v203_ = nw235_
                index237_ = default__.safeIndex(409, (d_1629_v198_).length(0))
                rhs278_ = (0) - (len(d_1635_v202_))
                rhs279_ = d_1636_v203_
                rhs280_ = (d_1617_v187_) <= (d_1617_v187_)
                lhs227_ = d_1629_v198_
                lhs228_ = default__.safeIndex(409, (d_1629_v198_).length(0))
                lhs229_ = globalState
                lhs227_[lhs228_] = rhs278_
                r1 = rhs279_
                lhs229_.f0 = rhs280_
                d_1639_v204_: _dafny.Seq
                d_1639_v204_ = _dafny.SeqWithoutIsStrInference([p0, (self).f24])
                d_1640_v205_: _dafny.Seq
                d_1640_v205_ = _dafny.SeqWithoutIsStrInference([(d_1639_v204_) + (d_1639_v204_)])
                d_1640_v205_ = d_1640_v205_
                d_1641_v206_: _dafny.Array
                nw236_ = _dafny.Array(_dafny.MultiSet({}), 29)
                d_1641_v206_ = nw236_
                d_1642_v207_: _dafny.MultiSet
                d_1642_v207_ = _dafny.MultiSet([(d_1629_v198_)[default__.safeIndex(409, (d_1629_v198_).length(0))]])
                index238_ = default__.safeIndex(899, (d_1641_v206_).length(0))
                (d_1641_v206_)[index238_] = d_1642_v207_
                index239_ = default__.safeIndex(899, (d_1641_v206_).length(0))
                (d_1641_v206_)[index239_] = d_1642_v207_
                d_1643_v208_: D2
                d_1643_v208_ = D2_DC6((d_1629_v198_)[default__.safeIndex(409, (d_1629_v198_).length(0))])
                d_1644_v209_: D2
                d_1644_v209_ = D2_DC9(D2_DC9(d_1643_v208_))
                d_1645_v210_: _dafny.MultiSet
                d_1645_v210_ = _dafny.MultiSet([d_1644_v209_])
                d_1646_v211_: C4
                nw237_ = C4()
                nw237_.ctor__((d_1645_v210_) != (d_1645_v210_), (self).f22)
                d_1646_v211_ = nw237_
            elif True:
                d_1647_v212_: D11
                d_1647_v212_ = D11_DC39((self).f24)
                d_1648_v213_: _dafny.Array
                nw238_ = _dafny.Array(None, 1)
                nw238_[int(0)] = d_1647_v212_
                d_1648_v213_ = nw238_
                index240_ = default__.safeIndex(683, (d_1648_v213_).length(0))
                (d_1648_v213_)[index240_] = D11_DC39((self).f24)
                d_1649_v214_: D1
                d_1649_v214_ = D1_DC3(_dafny.SeqWithoutIsStrInference([d_1628_v197_ for d_1650_i20_ in range(default__.abs(-352))]))
                d_1651_v215_: _dafny.Seq
                def iife132_(_pat_let32_0):
                    def iife133_(d_1652_dt__update__tmp_h3_):
                        def iife134_(_pat_let33_0):
                            def iife135_(d_1653_dt__update_hcf6_h0_):
                                return D1_DC3(d_1653_dt__update_hcf6_h0_)
                            return iife135_(_pat_let33_0)
                        return iife134_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujsmsmjoj")))
                    return iife133_(_pat_let32_0)
                d_1651_v215_ = _dafny.SeqWithoutIsStrInference([d_1649_v214_, iife132_(D1_DC3(d_1617_v187_))])
                d_1654_v216_: _dafny.Array
                nw239_ = _dafny.Array(None, 4)
                nw239_[int(0)] = _dafny.CodePoint('i')
                nw239_[int(1)] = d_1628_v197_
                nw239_[int(2)] = (d_1617_v187_)[default__.safeIndex(d_1406_v50_, len(d_1617_v187_))]
                nw239_[int(3)] = d_1628_v197_
                d_1654_v216_ = nw239_
                d_1655_v217_: _dafny.Seq
                d_1655_v217_ = _dafny.SeqWithoutIsStrInference([p0])
                index241_ = default__.safeIndex(8, (d_1654_v216_).length(0))
                (d_1654_v216_)[index241_] = (d_1617_v187_)[default__.safeIndex((0) - (len(d_1655_v217_)), len(d_1617_v187_))]
                d_1656_v218_: _dafny.Map
                d_1656_v218_ = _dafny.Map({(self).f24: d_1406_v50_})
                d_1657_v219_: _dafny.Map
                d_1657_v219_ = _dafny.Map({len(d_1656_v218_): 4})
                d_1658_v220_: D0
                d_1658_v220_ = D0_DC1(d_1406_v50_, (p0) or (p0), _dafny.Map({(self).f24: d_1657_v219_}), (d_1406_v50_ if p0 else 186), d_1628_v197_)
                index242_ = default__.safeIndex(683, (d_1648_v213_).length(0))
                index243_ = default__.safeIndex(8, (d_1654_v216_).length(0))
                rhs281_ = d_1406_v50_
                rhs282_ = d_1647_v212_
                rhs283_ = _dafny.SeqWithoutIsStrInference([d_1649_v214_ for d_1659_i21_ in range(default__.abs(90))])
                rhs284_ = d_1628_v197_
                rhs285_ = d_1658_v220_
                lhs230_ = d_1648_v213_
                lhs231_ = default__.safeIndex(683, (d_1648_v213_).length(0))
                lhs232_ = d_1654_v216_
                lhs233_ = default__.safeIndex(8, (d_1654_v216_).length(0))
                r3 = rhs281_
                lhs230_[lhs231_] = rhs282_
                d_1651_v215_ = rhs283_
                lhs232_[lhs233_] = rhs284_
                d_1658_v220_ = rhs285_
                rhs286_ = d_1406_v50_
                rhs287_ = (d_1654_v216_)[default__.safeIndex(8, (d_1654_v216_).length(0))]
                lhs234_ = globalState
                lhs234_.f11 = rhs286_
                d_1628_v197_ = rhs287_
                d_1660_v221_: _dafny.Array
                nw240_ = _dafny.Array(_dafny.Map({}), 13)
                d_1660_v221_ = nw240_
                d_1661_v222_: T0
                nw241_ = C1()
                nw241_.ctor__(d_1406_v50_, (self).f22)
                d_1661_v222_ = nw241_
                d_1662_v223_: _dafny.MultiSet
                d_1662_v223_ = _dafny.MultiSet([p0])
                d_1663_v224_: D14
                d_1663_v224_ = D14_DC47(d_1662_v223_)
                d_1664_v225_: _dafny.Map
                d_1664_v225_ = _dafny.Map({d_1661_v222_: (d_1663_v224_).cf77})
                index244_ = default__.safeIndex(650, (d_1660_v221_).length(0))
                (d_1660_v221_)[index244_] = d_1664_v225_
                index245_ = default__.safeIndex(650, (d_1660_v221_).length(0))
                (d_1660_v221_)[index245_] = ((d_1664_v225_) | ((d_1664_v225_).set(d_1661_v222_, d_1662_v223_))) | ((_dafny.Map({d_1661_v222_: d_1662_v223_})) | (d_1664_v225_))
                d_1665_v226_: _dafny.Array
                nw242_ = _dafny.Array(None, 7)
                d_1665_v226_ = nw242_
                d_1666_v227_: D11
                d_1666_v227_ = D11_DC37(d_1665_v226_)
                d_1667_v228_: _dafny.Array
                nw243_ = _dafny.Array(None, 10)
                nw243_[int(0)] = d_1406_v50_
                nw243_[int(1)] = len(d_1617_v187_)
                nw243_[int(2)] = d_1406_v50_
                nw243_[int(3)] = d_1406_v50_
                nw243_[int(4)] = len((d_1617_v187_).set(default__.safeIndex(len(d_1617_v187_), len(d_1617_v187_)), (d_1654_v216_)[default__.safeIndex(8, (d_1654_v216_).length(0))]))
                nw243_[int(5)] = len((d_1655_v217_).set(default__.safeIndex((d_1662_v223_).cardinality, len(d_1655_v217_)), (self).f24))
                nw243_[int(6)] = d_1406_v50_
                nw243_[int(7)] = d_1406_v50_
                nw243_[int(8)] = 380
                nw243_[int(9)] = d_1406_v50_
                d_1667_v228_ = nw243_
                d_1668_v229_: D2
                d_1668_v229_ = D2_DC8(d_1649_v214_, d_1667_v228_, p0)
                d_1669_v230_: _dafny.Map
                d_1669_v230_ = _dafny.Map({d_1666_v227_: d_1668_v229_})
                d_1670_v231_: _dafny.Map
                d_1670_v231_ = _dafny.Map({d_1406_v50_: d_1665_v226_})
                d_1671_v232_: _dafny.Seq
                d_1671_v232_ = _dafny.SeqWithoutIsStrInference([d_1667_v228_])
                pat_let_tv61_ = d_1670_v231_
                pat_let_tv62_ = d_1671_v232_
                pat_let_tv63_ = d_1671_v232_
                pat_let_tv64_ = d_1670_v231_
                pat_let_tv65_ = d_1665_v226_
                def iife136_(_pat_let34_0):
                    def iife137_(d_1672_dt__update__tmp_h4_):
                        def iife138_(_pat_let35_0):
                            def iife139_(d_1673_dt__update_hcf63_h0_):
                                return D11_DC37(d_1673_dt__update_hcf63_h0_)
                            return iife139_(_pat_let35_0)
                        return iife138_(((pat_let_tv61_)[len(pat_let_tv62_)] if (len(pat_let_tv63_)) in (pat_let_tv64_) else pat_let_tv65_))
                    return iife137_(_pat_let34_0)
                d_1669_v230_ = (d_1669_v230_).set(iife136_(d_1666_v227_), d_1668_v229_)
                d_1674_v233_: _dafny.Array
                nw244_ = _dafny.Array(False, 9)
                d_1674_v233_ = nw244_
                d_1675_v234_: _dafny.MultiSet
                d_1675_v234_ = _dafny.MultiSet([d_1674_v233_, d_1674_v233_, d_1674_v233_])
                d_1675_v234_ = ((d_1675_v234_) - (_dafny.MultiSet([d_1674_v233_, d_1674_v233_]))).intersection(d_1675_v234_)
            d_1676_v235_: _dafny.Array
            nw245_ = _dafny.Array(False, 27)
            d_1676_v235_ = nw245_
            index246_ = default__.safeIndex(46, (d_1676_v235_).length(0))
            (d_1676_v235_)[index246_] = p0
            index247_ = default__.safeIndex(46, (d_1676_v235_).length(0))
            (d_1676_v235_)[index247_] = p0
            d_1677_v236_: T0
            nw246_ = C4()
            nw246_.ctor__((d_1676_v235_)[default__.safeIndex(46, (d_1676_v235_).length(0))], (self).f22)
            d_1677_v236_ = nw246_
            d_1678_v237_: _dafny.Set
            d_1678_v237_ = _dafny.Set({d_1406_v50_})
            d_1679_v238_: _dafny.Seq
            d_1679_v238_ = _dafny.SeqWithoutIsStrInference([d_1678_v237_])
            d_1680_v239_: _dafny.Set
            d_1680_v239_ = _dafny.Set({True, (self).f24, p0})
            d_1681_v240_: _dafny.Array
            nw247_ = _dafny.Array(None, 9)
            nw247_[int(0)] = -876
            nw247_[int(1)] = (0) - (len(d_1617_v187_))
            nw247_[int(2)] = (0) - (d_1406_v50_)
            nw247_[int(3)] = d_1406_v50_
            nw247_[int(4)] = d_1406_v50_
            nw247_[int(5)] = d_1406_v50_
            nw247_[int(6)] = d_1406_v50_
            nw247_[int(7)] = (d_1677_v236_).fm0(((self).f25)[default__.safeIndex(d_1406_v50_, len((self).f25))], d_1406_v50_, d_1679_v238_, (self).f24, globalState)
            nw247_[int(8)] = len(d_1680_v239_)
            d_1681_v240_ = nw247_
            d_1682_v241_: _dafny.Set
            d_1682_v241_ = _dafny.Set({d_1681_v240_})
            (globalState).f10 = (D13_DC46((self).f24, d_1677_v236_, d_1406_v50_, d_1682_v241_)).cf73
            d_1683_v242_: _dafny.Seq
            d_1683_v242_ = _dafny.SeqWithoutIsStrInference([(d_1676_v235_)[default__.safeIndex(46, (d_1676_v235_).length(0))]])
            d_1684_v243_: _dafny.MultiSet
            d_1684_v243_ = _dafny.MultiSet([(d_1683_v242_) + (_dafny.SeqWithoutIsStrInference([(self).f24]))])
            d_1685_v244_: _dafny.MultiSet
            d_1685_v244_ = _dafny.MultiSet([(d_1676_v235_)[default__.safeIndex(46, (d_1676_v235_).length(0))], (self).f24])
            d_1684_v243_ = default__.fm50(((d_1685_v244_).cardinality) * (d_1406_v50_), globalState)
            (globalState).f0 = (d_1676_v235_)[default__.safeIndex(46, (d_1676_v235_).length(0))]
        d_1686_v245_: str
        d_1686_v245_ = _dafny.CodePoint('k')
        d_1687_v246_: _dafny.Set
        d_1687_v246_ = _dafny.Set({d_1686_v245_})
        d_1688_v247_: _dafny.Map
        d_1688_v247_ = _dafny.Map({d_1686_v245_: _dafny.Set({d_1687_v246_, default__.fm33(globalState), d_1687_v246_, d_1687_v246_})})
        d_1689_v248_: _dafny.Map
        d_1689_v248_ = _dafny.Map({_dafny.Set({d_1686_v245_}): len(_dafny.SeqWithoutIsStrInference([d_1406_v50_]))})
        def iife140_():
            coll68_ = _dafny.Set()
            compr_68_: _dafny.Set
            for compr_68_ in (d_1689_v248_).keys.Elements:
                d_1690_v249_: _dafny.Set = compr_68_
                if (d_1690_v249_) in (d_1689_v248_):
                    coll68_ = coll68_.union(_dafny.Set([d_1690_v249_]))
            return _dafny.Set(coll68_)
        r0 = (default__.fm7(d_1406_v50_, globalState)) < (len(((d_1688_v247_)[d_1686_v245_] if (d_1686_v245_) in (d_1688_v247_) else iife140_()
        )))
        nw248_ = _dafny.Array(False, 21)
        r1 = nw248_
        def iife141_():
            coll69_ = _dafny.Map()
            compr_69_: _dafny.Seq
            for compr_69_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfkcni"))).set(default__.safeIndex(d_1406_v50_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfkcni")))), d_1686_v245_)) for d_1691_i23_ in range(default__.abs(822))]) for d_1692_i22_ in range(default__.abs(915))])).Elements:
                d_1693_v250_: _dafny.Seq = compr_69_
                if (d_1693_v250_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfkcni"))).set(default__.safeIndex(d_1406_v50_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfkcni")))), d_1686_v245_)) for d_1691_i23_ in range(default__.abs(822))]) for d_1692_i22_ in range(default__.abs(915))])):
                    coll69_[d_1693_v250_] = d_1686_v245_
            return _dafny.Map(coll69_)
        r2 = iife141_()
        
        r3 = default__.safeModuloInt(d_1406_v50_, d_1406_v50_)
        return r0, r1, r2, r3

    def m1(self, globalState):
        r0: bool = False
        r1: int = int(0)
        d_1694_v0_: int
        d_1694_v0_ = 217
        r1 = d_1694_v0_
        if (self).f24:
            d_1695_v1_: D15
            d_1695_v1_ = D15_DC49(_dafny.Set({(self).f24}))
            d_1696_v2_: _dafny.Set
            d_1696_v2_ = _dafny.Set({(self).f24})
            (globalState).f0 = (((d_1695_v1_).cf81) - (d_1696_v2_)) == (d_1696_v2_)
            d_1697_v3_: _dafny.Array
            def lambda81_(d_1698_v0_):
                def lambda82_(d_1699_i0_):
                    return default__.safeModuloInt(d_1699_i0_, len(_dafny.Set({d_1698_v0_, len(_dafny.Map({(self).f24: d_1698_v0_}))})))

                return lambda82_

            init42_ = lambda81_(d_1694_v0_)
            nw249_ = _dafny.Array(None, 23)
            for i0_42_ in range(nw249_.length(0)):
                nw249_[i0_42_] = init42_(i0_42_)
            d_1697_v3_ = nw249_
            d_1700_v4_: _dafny.Set
            d_1700_v4_ = _dafny.Set({d_1697_v3_, d_1697_v3_, d_1697_v3_, d_1697_v3_})
            d_1701_v5_: _dafny.Seq
            d_1701_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxel"))
            d_1702_v6_: D15
            d_1702_v6_ = D15_DC51((self).f24, (self).f24)
            d_1703_v7_: _dafny.Seq
            d_1703_v7_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, (self).f24, (self).f24])
            d_1704_v8_: _dafny.Map
            d_1704_v8_ = _dafny.Map({d_1694_v0_: not((self).f24)})
            d_1705_v9_: _dafny.Array
            nw250_ = _dafny.Array(None, 24)
            nw250_[int(0)] = (self).f24
            nw250_[int(1)] = (self).f24
            nw250_[int(2)] = (self).f24
            nw250_[int(3)] = (self).f24
            nw250_[int(4)] = (self).f24
            nw250_[int(5)] = (self).f24
            nw250_[int(6)] = (self).f24
            nw250_[int(7)] = not((self).f24)
            nw250_[int(8)] = (self).f24
            nw250_[int(9)] = (self).f24
            nw250_[int(10)] = default__.fm15((self).f24, globalState)
            nw250_[int(11)] = not((self).f24)
            nw250_[int(12)] = (self).f24
            nw250_[int(13)] = (self).f24
            nw250_[int(14)] = (d_1702_v6_).cf82
            nw250_[int(15)] = (self).f24
            nw250_[int(16)] = (self).f24
            nw250_[int(17)] = (self).f24
            nw250_[int(18)] = default__.fm37(d_1703_v7_, d_1704_v8_, d_1701_v5_, globalState)
            nw250_[int(19)] = (self).f24
            nw250_[int(20)] = (self).f24
            nw250_[int(21)] = (self).f24
            nw250_[int(22)] = default__.fm15((self).f24, globalState)
            nw250_[int(23)] = True
            d_1705_v9_ = nw250_
            d_1706_v10_: D9
            d_1706_v10_ = D9_DC31((len(d_1700_v4_)) + (len(d_1701_v5_)), d_1705_v9_)
            index248_ = default__.safeIndex(731, (d_1697_v3_).length(0))
            (d_1697_v3_)[index248_] = d_1694_v0_
            index249_ = default__.safeIndex(731, (d_1697_v3_).length(0))
            rhs288_ = d_1706_v10_
            rhs289_ = d_1697_v3_
            rhs290_ = d_1694_v0_
            rhs291_ = (((d_1703_v7_).set(default__.safeIndex(d_1694_v0_, len(d_1703_v7_)), (self).f24)).set(default__.safeIndex((0) - (d_1694_v0_), len((d_1703_v7_).set(default__.safeIndex(d_1694_v0_, len(d_1703_v7_)), (self).f24))), (self).f24)) + (_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, (self).f24]))
            rhs292_ = d_1694_v0_
            lhs235_ = d_1697_v3_
            lhs236_ = default__.safeIndex(731, (d_1697_v3_).length(0))
            lhs237_ = globalState
            d_1706_v10_ = rhs288_
            d_1697_v3_ = rhs289_
            lhs235_[lhs236_] = rhs290_
            d_1703_v7_ = rhs291_
            lhs237_.f8 = rhs292_
            d_1707_v11_: _dafny.Map
            d_1707_v11_ = _dafny.Map({(d_1697_v3_)[default__.safeIndex(731, (d_1697_v3_).length(0))]: default__.fm7((0) - (len(d_1696_v2_)), globalState)})
            d_1708_v12_: _dafny.Map
            d_1708_v12_ = _dafny.Map({(self).f24: (self).f24})
            d_1707_v11_ = (d_1707_v11_).set(default__.safeDivisionInt(d_1694_v0_, (d_1697_v3_)[default__.safeIndex(731, (d_1697_v3_).length(0))]), (0) - ((len(_dafny.SeqWithoutIsStrInference([d_1708_v12_, d_1708_v12_]))) - ((d_1697_v3_)[default__.safeIndex(731, (d_1697_v3_).length(0))])))
            d_1709_v13_: _dafny.Map
            d_1709_v13_ = _dafny.Map({d_1701_v5_: (self).f24})
            d_1710_v14_: _dafny.Map
            d_1710_v14_ = _dafny.Map({not((self).f24): d_1709_v13_})
            rhs293_ = (d_1709_v13_) | (((d_1710_v14_)[(self).f24] if ((self).f24) in (d_1710_v14_) else d_1709_v13_))
            rhs294_ = (len(d_1701_v5_)) == ((d_1697_v3_)[default__.safeIndex(731, (d_1697_v3_).length(0))])
            lhs238_ = globalState
            d_1709_v13_ = rhs293_
            lhs238_.f0 = rhs294_
            d_1694_v0_ = d_1694_v0_
        elif True:
            d_1711_v15_: _dafny.Array
            def lambda83_(d_1712_i1_):
                return (self).f24

            init43_ = lambda83_
            nw251_ = _dafny.Array(None, 2)
            for i0_43_ in range(nw251_.length(0)):
                nw251_[i0_43_] = init43_(i0_43_)
            d_1711_v15_ = nw251_
            d_1713_v16_: C1
            nw252_ = C1()
            nw252_.ctor__(default__.safeDivisionInt(d_1694_v0_, 817), (_dafny.SeqWithoutIsStrInference([d_1711_v15_, d_1711_v15_, d_1711_v15_, d_1711_v15_, d_1711_v15_])) + ((self).f22))
            d_1713_v16_ = nw252_
            if default__.fm15(True, globalState):
                d_1714_v17_: int
                d_1715_v18_: _dafny.Set
                d_1716_v19_: int
                d_1717_v20_: int
                out60_: int
                out61_: _dafny.Set
                out62_: int
                out63_: int
                out60_, out61_, out62_, out63_ = (d_1713_v16_).m6(d_1694_v0_, ((d_1713_v16_).f30) == ((d_1713_v16_).f30), default__.safeModuloInt(d_1694_v0_, d_1694_v0_), (d_1713_v16_).f30, globalState)
                d_1714_v17_ = out60_
                d_1715_v18_ = out61_
                d_1716_v19_ = out62_
                d_1717_v20_ = out63_
                d_1718_v21_: str
                d_1718_v21_ = _dafny.CodePoint('q')
                d_1718_v21_ = d_1718_v21_
                d_1719_v22_: _dafny.Set
                d_1719_v22_ = _dafny.Set({(self).f24, (self).f24, False})
                d_1720_v23_: _dafny.Seq
                d_1720_v23_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1721_v24_: _dafny.Array
                nw253_ = _dafny.Array(None, 23)
                nw253_[int(0)] = (d_1719_v22_).intersection(d_1719_v22_)
                nw253_[int(1)] = d_1719_v22_
                nw253_[int(2)] = (d_1719_v22_ if (self).f24 else d_1719_v22_)
                nw253_[int(3)] = d_1719_v22_
                nw253_[int(4)] = (d_1719_v22_).intersection(d_1719_v22_)
                nw253_[int(5)] = _dafny.Set({(self).f24})
                nw253_[int(6)] = d_1719_v22_
                nw253_[int(7)] = d_1719_v22_
                nw253_[int(8)] = d_1719_v22_
                nw253_[int(9)] = d_1719_v22_
                nw253_[int(10)] = (d_1719_v22_).intersection(d_1719_v22_)
                nw253_[int(11)] = _dafny.Set({(self).f24, (self).f24, (self).f24})
                nw253_[int(12)] = d_1719_v22_
                nw253_[int(13)] = d_1719_v22_
                nw253_[int(14)] = (_dafny.Set({(d_1720_v23_)[default__.safeIndex((d_1713_v16_).f30, len(d_1720_v23_))]})) - (d_1719_v22_)
                nw253_[int(15)] = d_1719_v22_
                nw253_[int(16)] = (d_1719_v22_) | (d_1719_v22_)
                nw253_[int(17)] = d_1719_v22_
                nw253_[int(18)] = d_1719_v22_
                nw253_[int(19)] = d_1719_v22_
                nw253_[int(20)] = d_1719_v22_
                nw253_[int(21)] = (default__.fm11(d_1717_v20_, d_1718_v21_, (self).f24, globalState)) | (d_1719_v22_)
                nw253_[int(22)] = _dafny.Set({not((self).f24), False})
                d_1721_v24_ = nw253_
                index250_ = default__.safeIndex(226, (d_1721_v24_).length(0))
                (d_1721_v24_)[index250_] = d_1719_v22_
                index251_ = default__.safeIndex(226, (d_1721_v24_).length(0))
                (d_1721_v24_)[index251_] = d_1719_v22_
                d_1722_v25_: D13
                d_1722_v25_ = D13_DC45(d_1716_v19_)
                pat_let_tv66_ = d_1714_v17_
                def iife142_(_pat_let36_0):
                    def iife143_(d_1723_dt__update__tmp_h0_):
                        def iife144_(_pat_let37_0):
                            def iife145_(d_1724_dt__update_hcf72_h0_):
                                return D13_DC45(d_1724_dt__update_hcf72_h0_)
                            return iife145_(_pat_let37_0)
                        return iife144_(pat_let_tv66_)
                    return iife143_(_pat_let36_0)
                d_1722_v25_ = iife142_(d_1722_v25_)
                d_1725_v26_: C2
                nw254_ = C2()
                nw254_.ctor__(d_1717_v20_, _dafny.SeqWithoutIsStrInference([d_1711_v15_, d_1711_v15_, d_1711_v15_, d_1711_v15_]))
                d_1725_v26_ = nw254_
                d_1726_v27_: _dafny.Map
                d_1726_v27_ = _dafny.Map({296: d_1714_v17_})
                d_1727_v28_: _dafny.Array
                nw255_ = _dafny.Array(None, 11)
                nw255_[int(0)] = (d_1713_v16_).f30
                nw255_[int(1)] = -755
                nw255_[int(2)] = len(d_1726_v27_)
                nw255_[int(3)] = 978
                nw255_[int(4)] = (d_1713_v16_).f30
                nw255_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_1718_v21_ for d_1728_i2_ in range(default__.abs(692))]))
                nw255_[int(6)] = d_1716_v19_
                nw255_[int(7)] = (d_1725_v26_).f29
                nw255_[int(8)] = d_1716_v19_
                nw255_[int(9)] = (d_1713_v16_).f30
                nw255_[int(10)] = d_1694_v0_
                d_1727_v28_ = nw255_
                d_1729_v29_: _dafny.Map
                d_1729_v29_ = _dafny.Map({(self).f24: _dafny.Map({d_1725_v26_: d_1727_v28_})})
                d_1730_v30_: _dafny.Map
                d_1730_v30_ = _dafny.Map({d_1725_v26_: d_1727_v28_})
                d_1729_v29_ = (d_1729_v29_).set(not((self).f24), (d_1730_v30_) | (d_1730_v30_))
            elif True:
                d_1731_v31_: _dafny.Seq
                d_1731_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iyldgbb"))
                d_1732_v32_: str
                d_1732_v32_ = _dafny.CodePoint('i')
                d_1733_v33_: _dafny.Seq
                d_1733_v33_ = _dafny.SeqWithoutIsStrInference([len(((d_1731_v31_).set(default__.safeIndex(346, len(d_1731_v31_)), d_1732_v32_)) + (d_1731_v31_))])
                d_1734_v34_: D2
                d_1734_v34_ = D2_DC7(d_1733_v33_, (self).f24, (self).f24, (d_1713_v16_).f30)
                d_1735_v35_: _dafny.Map
                d_1735_v35_ = _dafny.Map({-395: (d_1734_v34_).cf10})
                d_1733_v33_ = ((d_1733_v33_).set(default__.safeIndex((d_1713_v16_).fm34(False, globalState), len(d_1733_v33_)), len(d_1731_v31_))) + (((d_1735_v35_)[(d_1713_v16_).f30] if ((d_1713_v16_).f30) in (d_1735_v35_) else d_1733_v33_))
                d_1736_v36_: C1
                nw256_ = C1()
                nw256_.ctor__((d_1713_v16_).f30, (((self).f22).set(default__.safeIndex((d_1713_v16_).f30, len((self).f22)), d_1711_v15_)) + (_dafny.SeqWithoutIsStrInference([d_1711_v15_])))
                d_1736_v36_ = nw256_
                d_1737_v37_: _dafny.MultiSet
                d_1737_v37_ = _dafny.MultiSet([(self).f24, (self).f24, (self).f24, (self).f24])
                d_1738_v38_: _dafny.MultiSet
                d_1738_v38_ = _dafny.MultiSet([d_1694_v0_])
                d_1739_v39_: _dafny.Set
                d_1739_v39_ = _dafny.Set({(d_1713_v16_).f30})
                d_1740_v40_: _dafny.Seq
                d_1740_v40_ = _dafny.SeqWithoutIsStrInference([d_1739_v39_])
                d_1741_v41_: _dafny.Map
                d_1741_v41_ = _dafny.Map({(d_1713_v16_).fm0(d_1738_v38_, d_1694_v0_, d_1740_v40_, (self).f24, globalState): (self).f24})
                (globalState).f16 = (((d_1737_v37_)[(self).f24] if ((self).f24) in (d_1737_v37_) else len(d_1741_v41_))) * (((d_1736_v36_).f30) + ((d_1736_v36_).f30))
                r0 = (self).fm3(d_1731_v31_, ((self).f24 if (self).f24 else (self).f24), (self).f24, globalState)
                d_1742_v42_: _dafny.Map
                d_1742_v42_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcllqkg")): d_1733_v33_})
                d_1743_v44_: _dafny.Map
                d_1743_v44_ = _dafny.Map({d_1731_v31_: (d_1736_v36_).f30})
                def iife146_():
                    coll70_ = _dafny.Map()
                    compr_70_: _dafny.Seq
                    for compr_70_ in (d_1743_v44_).keys.Elements:
                        d_1744_v43_: _dafny.Seq = compr_70_
                        if (d_1744_v43_) in (d_1743_v44_):
                            coll70_[d_1744_v43_] = (d_1733_v33_) + (d_1733_v33_)
                    return _dafny.Map(coll70_)
                d_1742_v42_ = iife146_()
                
            r1 = d_1694_v0_
            d_1745_v45_: _dafny.Seq
            d_1745_v45_ = _dafny.SeqWithoutIsStrInference([default__.fm51(globalState)])
            d_1746_v46_: _dafny.Array
            nw257_ = _dafny.Array(int(0), 15)
            d_1746_v46_ = nw257_
            d_1747_v47_: _dafny.Seq
            d_1747_v47_ = _dafny.SeqWithoutIsStrInference([d_1746_v46_])
            d_1748_v48_: _dafny.Map
            d_1748_v48_ = _dafny.Map({573: d_1746_v46_})
            d_1749_v49_: _dafny.Set
            d_1749_v49_ = _dafny.Set({(d_1713_v16_).f30})
            d_1750_v50_: _dafny.Array
            nw258_ = _dafny.Array(None, 19)
            nw258_[int(0)] = d_1746_v46_
            nw258_[int(1)] = d_1746_v46_
            nw258_[int(2)] = d_1746_v46_
            nw258_[int(3)] = (d_1747_v47_)[default__.safeIndex(908, len(d_1747_v47_))]
            nw258_[int(4)] = d_1746_v46_
            nw258_[int(5)] = d_1746_v46_
            nw258_[int(6)] = d_1746_v46_
            nw258_[int(7)] = d_1746_v46_
            nw258_[int(8)] = (d_1746_v46_ if (self).f24 else d_1746_v46_)
            nw258_[int(9)] = d_1746_v46_
            nw258_[int(10)] = d_1746_v46_
            nw258_[int(11)] = d_1746_v46_
            nw258_[int(12)] = d_1746_v46_
            nw258_[int(13)] = d_1746_v46_
            nw258_[int(14)] = (d_1746_v46_ if (self).f24 else d_1746_v46_)
            nw258_[int(15)] = d_1746_v46_
            nw258_[int(16)] = ((d_1748_v48_)[(d_1713_v16_).f30] if ((d_1713_v16_).f30) in (d_1748_v48_) else d_1746_v46_)
            nw258_[int(17)] = (((d_1748_v48_)[len(d_1749_v49_)] if (len(d_1749_v49_)) in (d_1748_v48_) else d_1746_v46_) if (self).f24 else d_1746_v46_)
            nw258_[int(18)] = d_1746_v46_
            d_1750_v50_ = nw258_
            index252_ = default__.safeIndex(494, (d_1750_v50_).length(0))
            (d_1750_v50_)[index252_] = d_1746_v46_
            d_1751_v51_: D4
            d_1751_v51_ = D4_DC15(238, d_1694_v0_)
            index253_ = default__.safeIndex(494, (d_1750_v50_).length(0))
            rhs295_ = (d_1745_v45_) + (d_1745_v45_)
            rhs296_ = d_1694_v0_
            rhs297_ = d_1746_v46_
            rhs298_ = (((d_1751_v51_).cf26) >= (d_1694_v0_)) == ((self).f24)
            lhs239_ = globalState
            lhs240_ = d_1750_v50_
            lhs241_ = default__.safeIndex(494, (d_1750_v50_).length(0))
            lhs242_ = globalState
            d_1745_v45_ = rhs295_
            lhs239_.f6 = rhs296_
            lhs240_[lhs241_] = rhs297_
            lhs242_.f10 = rhs298_
            d_1752_v52_: C3
            nw259_ = C3()
            nw259_.ctor__(((self).f22) + ((self).f22))
            d_1752_v52_ = nw259_
        d_1753_i3_: int
        d_1753_i3_ = 0
        with _dafny.label("7"):
            while (self).f24:
                with _dafny.c_label("7"):
                    if (d_1753_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_1753_i3_ = (d_1753_i3_) + (1)
                    d_1754_v53_: _dafny.Array
                    def lambda84_(d_1755_i4_):
                        return (_dafny.Map({(self).f24: (self).f24})) != (_dafny.Map({(self).f24: (self).f24}))

                    init44_ = lambda84_
                    nw260_ = _dafny.Array(None, 4)
                    for i0_44_ in range(nw260_.length(0)):
                        nw260_[i0_44_] = init44_(i0_44_)
                    d_1754_v53_ = nw260_
                    index254_ = default__.safeIndex(80, (d_1754_v53_).length(0))
                    (d_1754_v53_)[index254_] = not((self).f24)
                    d_1756_v54_: _dafny.Seq
                    d_1756_v54_ = _dafny.SeqWithoutIsStrInference([d_1694_v0_])
                    d_1757_v55_: _dafny.Seq
                    d_1757_v55_ = _dafny.SeqWithoutIsStrInference([d_1756_v54_, (d_1756_v54_).set(default__.safeIndex(d_1694_v0_, len(d_1756_v54_)), d_1694_v0_), d_1756_v54_])
                    index255_ = default__.safeIndex(80, (d_1754_v53_).length(0))
                    (d_1754_v53_)[index255_] = (d_1756_v54_) not in (d_1757_v55_)
                    d_1758_v56_: _dafny.Seq
                    d_1758_v56_ = _dafny.SeqWithoutIsStrInference([(d_1754_v53_)[default__.safeIndex(80, (d_1754_v53_).length(0))]])
                    d_1759_v57_: _dafny.MultiSet
                    d_1759_v57_ = _dafny.MultiSet([d_1694_v0_, d_1694_v0_, d_1694_v0_, len(d_1758_v56_)])
                    (globalState).f0 = (d_1694_v0_) not in ((d_1759_v57_).intersection(d_1759_v57_))
                    d_1760_v58_: _dafny.Set
                    d_1760_v58_ = _dafny.Set({(d_1754_v53_)[default__.safeIndex(80, (d_1754_v53_).length(0))], (d_1754_v53_)[default__.safeIndex(80, (d_1754_v53_).length(0))], (self).f24})
                    d_1761_v59_: _dafny.Set
                    d_1761_v59_ = _dafny.Set({526, len(d_1760_v58_), 272, d_1694_v0_})
                    (globalState).f13 = d_1761_v59_
                    d_1762_v60_: _dafny.Map
                    d_1762_v60_ = _dafny.Map({(d_1754_v53_)[default__.safeIndex(80, (d_1754_v53_).length(0))]: d_1694_v0_})
                    d_1763_v61_: _dafny.Map
                    d_1763_v61_ = _dafny.Map({d_1694_v0_: len(d_1762_v60_)})
                    d_1764_v62_: _dafny.Map
                    d_1764_v62_ = _dafny.Map({not((self).f24): d_1763_v61_})
                    d_1764_v62_ = (d_1764_v62_).set((d_1754_v53_)[default__.safeIndex(80, (d_1754_v53_).length(0))], (d_1763_v61_).set(d_1694_v0_, d_1694_v0_))
                    pass
            pass
        if not((self).f24):
            d_1765_v63_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
            d_1765_v63_ = nw261_
            d_1766_v64_: _dafny.Seq
            d_1766_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlluhdhw"))
            index256_ = default__.safeIndex(492, (d_1765_v63_).length(0))
            (d_1765_v63_)[index256_] = d_1766_v64_
            d_1767_v65_: _dafny.Set
            d_1767_v65_ = _dafny.Set({d_1694_v0_})
            d_1768_v67_: _dafny.MultiSet
            d_1768_v67_ = _dafny.MultiSet([True])
            index257_ = default__.safeIndex(492, (d_1765_v63_).length(0))
            def iife147_():
                coll71_ = _dafny.Set()
                compr_71_: int
                for compr_71_ in _dafny.IntegerRange(9, 495):
                    d_1769_v66_: int = compr_71_
                    if ((9) <= (d_1769_v66_)) and ((d_1769_v66_) < (495)):
                        coll71_ = coll71_.union(_dafny.Set([default__.safeModuloInt(d_1769_v66_, d_1694_v0_)]))
                return _dafny.Set(coll71_)
            (d_1765_v63_)[index257_] = default__.fm36((self).f24, (d_1767_v65_) - (iife147_()
            ), d_1694_v0_, (d_1768_v67_).cardinality, globalState)
            d_1770_v68_: _dafny.Array
            def lambda85_(d_1771_i5_):
                return (self).f24

            init45_ = lambda85_
            nw262_ = _dafny.Array(None, 1)
            for i0_45_ in range(nw262_.length(0)):
                nw262_[i0_45_] = init45_(i0_45_)
            d_1770_v68_ = nw262_
            d_1772_v69_: _dafny.Map
            d_1772_v69_ = _dafny.Map({d_1770_v68_: (self).f24})
            d_1772_v69_ = ((d_1772_v69_).set(d_1770_v68_, (self).f24)) | (d_1772_v69_)
            d_1773_v70_: str
            d_1773_v70_ = _dafny.CodePoint('b')
            d_1774_v71_: _dafny.Map
            d_1774_v71_ = _dafny.Map({_dafny.Set({(self).f24}): d_1773_v70_})
            d_1775_v72_: _dafny.Seq
            d_1775_v72_ = _dafny.SeqWithoutIsStrInference([False])
            d_1776_v73_: _dafny.Map
            d_1776_v73_ = _dafny.Map({d_1694_v0_: (self).f24})
            d_1777_v74_: _dafny.MultiSet
            d_1777_v74_ = _dafny.MultiSet([618, len(d_1776_v73_), (0) - (len((d_1765_v63_)[default__.safeIndex(492, (d_1765_v63_).length(0))]))])
            d_1778_v75_: _dafny.Map
            d_1778_v75_ = _dafny.Map({(self).f24: (self).f24})
            d_1779_v76_: _dafny.Array
            nw263_ = _dafny.Array(None, 13)
            nw263_[int(0)] = len(d_1774_v71_)
            nw263_[int(1)] = len((d_1775_v72_) + (_dafny.SeqWithoutIsStrInference([(self).f24, (self).f24, (self).f24, False])))
            nw263_[int(2)] = (len(d_1767_v65_)) - (983)
            nw263_[int(3)] = (0) - (d_1694_v0_)
            nw263_[int(4)] = d_1694_v0_
            nw263_[int(5)] = len(d_1766_v64_)
            nw263_[int(6)] = default__.safeDivisionInt(d_1694_v0_, d_1694_v0_)
            nw263_[int(7)] = (d_1694_v0_) * (((d_1777_v74_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrhauoaod")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrhauoaod")))) in (d_1777_v74_) else d_1694_v0_))
            nw263_[int(8)] = d_1694_v0_
            nw263_[int(9)] = -508
            nw263_[int(10)] = len((d_1778_v75_) | (d_1778_v75_))
            nw263_[int(11)] = d_1694_v0_
            nw263_[int(12)] = d_1694_v0_
            d_1779_v76_ = nw263_
            d_1780_v77_: D2
            d_1780_v77_ = D2_DC6(d_1694_v0_)
            index258_ = default__.safeIndex(711, (d_1779_v76_).length(0))
            (d_1779_v76_)[index258_] = (d_1780_v77_).cf9
            index259_ = default__.safeIndex(711, (d_1779_v76_).length(0))
            rhs299_ = (d_1694_v0_) + (len((d_1765_v63_)[default__.safeIndex(492, (d_1765_v63_).length(0))]))
            rhs300_ = ((self).f24) or (not((self).f24))
            lhs243_ = d_1779_v76_
            lhs244_ = default__.safeIndex(711, (d_1779_v76_).length(0))
            lhs243_[lhs244_] = rhs299_
            r0 = rhs300_
            index260_ = default__.safeIndex(283, (d_1770_v68_).length(0))
            (d_1770_v68_)[index260_] = (not((self).f24)) or ((self).f24)
            index261_ = default__.safeIndex(283, (d_1770_v68_).length(0))
            rhs301_ = ((d_1775_v72_)[default__.safeIndex(d_1694_v0_, len(d_1775_v72_))]) in (((default__.fm52(globalState)).set(default__.safeIndex(len(d_1776_v73_), len(default__.fm52(globalState))), (self).f24)) + (d_1775_v72_))
            rhs302_ = ((d_1779_v76_)[default__.safeIndex(711, (d_1779_v76_).length(0))]) - (d_1694_v0_)
            rhs303_ = d_1694_v0_
            lhs245_ = d_1770_v68_
            lhs246_ = default__.safeIndex(283, (d_1770_v68_).length(0))
            lhs247_ = globalState
            lhs245_[lhs246_] = rhs301_
            r1 = rhs302_
            lhs247_.f8 = rhs303_
            d_1770_v68_ = d_1770_v68_
        elif True:
            r0 = (self).f24
            if (self).f24:
                d_1781_v78_: _dafny.Array
                nw264_ = _dafny.Array(int(0), 6)
                d_1781_v78_ = nw264_
                d_1782_v79_: _dafny.Map
                d_1782_v79_ = _dafny.Map({d_1781_v78_: (d_1694_v0_) == (d_1694_v0_)})
                d_1782_v79_ = (_dafny.Map({d_1781_v78_: False})) | (d_1782_v79_)
                (globalState).f10 = (self).f24
                d_1783_v80_: _dafny.Seq
                d_1783_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxmowc"))
                d_1783_v80_ = (d_1783_v80_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))) + (default__.fm30(globalState)))
                d_1784_v81_: C5
                nw265_ = C5()
                nw265_.ctor__(_dafny.Map({d_1694_v0_: (self).f24}), 895)
                d_1784_v81_ = nw265_
                d_1785_v82_: C2
                nw266_ = C2()
                nw266_.ctor__(default__.safeDivisionInt(d_1694_v0_, d_1694_v0_), (self).f22)
                d_1785_v82_ = nw266_
            elif True:
                d_1786_v83_: _dafny.Array
                nw267_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_1786_v83_ = nw267_
                d_1787_v85_: _dafny.Array
                def lambda86_(d_1788_v0_):
                    def lambda87_(d_1789_i6_):
                        def iife148_():
                            coll72_ = _dafny.Map()
                            compr_72_: int
                            for compr_72_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1788_v0_: (self).f24})), d_1788_v0_, (0) - (-811), len(_dafny.SeqWithoutIsStrInference([(0) - (d_1788_v0_)])), d_1788_v0_])).Elements:
                                d_1790_v84_: int = compr_72_
                                if (d_1790_v84_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_1788_v0_: (self).f24})), d_1788_v0_, (0) - (-811), len(_dafny.SeqWithoutIsStrInference([(0) - (d_1788_v0_)])), d_1788_v0_])):
                                    coll72_[default__.safeModuloInt(d_1790_v84_, d_1788_v0_)] = d_1788_v0_
                            return _dafny.Map(coll72_)
                        return iife148_()
                        

                    return lambda87_

                init46_ = lambda86_(d_1694_v0_)
                nw268_ = _dafny.Array(None, 18)
                for i0_46_ in range(nw268_.length(0)):
                    nw268_[i0_46_] = init46_(i0_46_)
                d_1787_v85_ = nw268_
                index262_ = default__.safeIndex(156, (d_1786_v83_).length(0))
                (d_1786_v83_)[index262_] = d_1787_v85_
                d_1791_v86_: str
                d_1791_v86_ = _dafny.CodePoint('i')
                d_1792_v87_: _dafny.Array
                nw269_ = _dafny.Array(False, 28)
                d_1792_v87_ = nw269_
                index263_ = default__.safeIndex(171, (d_1792_v87_).length(0))
                (d_1792_v87_)[index263_] = (self).f24
                index264_ = default__.safeIndex(156, (d_1786_v83_).length(0))
                index265_ = default__.safeIndex(171, (d_1792_v87_).length(0))
                rhs304_ = d_1787_v85_
                rhs305_ = default__.fm22((self).f24, not((self).f24), len((default__.fm52(globalState)) + (default__.fm52(globalState))), (684) * (d_1694_v0_), globalState)
                rhs306_ = (self).f24
                lhs248_ = d_1786_v83_
                lhs249_ = default__.safeIndex(156, (d_1786_v83_).length(0))
                lhs250_ = d_1792_v87_
                lhs251_ = default__.safeIndex(171, (d_1792_v87_).length(0))
                lhs248_[lhs249_] = rhs304_
                d_1791_v86_ = rhs305_
                lhs250_[lhs251_] = rhs306_
                d_1793_v88_: _dafny.Set
                d_1793_v88_ = _dafny.Set({d_1694_v0_, d_1694_v0_})
                d_1794_v89_: _dafny.MultiSet
                d_1794_v89_ = _dafny.MultiSet([d_1694_v0_, d_1694_v0_])
                d_1795_v90_: D6
                d_1795_v90_ = D6_DC23((self).f24, d_1694_v0_, (d_1792_v87_)[default__.safeIndex(171, (d_1792_v87_).length(0))])
                d_1796_v91_: _dafny.Seq
                d_1796_v91_ = _dafny.SeqWithoutIsStrInference([d_1793_v88_])
                d_1797_v92_: _dafny.Map
                d_1797_v92_ = _dafny.Map({default__.fm36((self).f24, d_1793_v88_, d_1694_v0_, (0) - ((self).fm0(d_1794_v89_, (d_1795_v90_).cf38, d_1796_v91_, (d_1792_v87_)[default__.safeIndex(171, (d_1792_v87_).length(0))], globalState)), globalState): (d_1694_v0_) + (d_1694_v0_)})
                d_1798_v93_: _dafny.Seq
                d_1798_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uji"))
                d_1797_v92_ = (d_1797_v92_).set(d_1798_v93_, (0) - (d_1694_v0_))
                d_1799_v94_: _dafny.Array
                def lambda88_(d_1800_v87_):
                    def lambda89_(d_1801_i7_):
                        return default__.safeModuloInt(d_1801_i7_, len(_dafny.Map({(d_1800_v87_)[default__.safeIndex(171, (d_1800_v87_).length(0))]: True})))

                    return lambda89_

                init47_ = lambda88_(d_1792_v87_)
                nw270_ = _dafny.Array(None, 16)
                for i0_47_ in range(nw270_.length(0)):
                    nw270_[i0_47_] = init47_(i0_47_)
                d_1799_v94_ = nw270_
                index266_ = default__.safeIndex(528, (d_1799_v94_).length(0))
                (d_1799_v94_)[index266_] = (713) + (len(default__.fm27(d_1694_v0_, globalState)))
                index267_ = default__.safeIndex(528, (d_1799_v94_).length(0))
                (d_1799_v94_)[index267_] = (d_1694_v0_) * (d_1694_v0_)
                d_1799_v94_ = d_1799_v94_
                d_1802_v95_: _dafny.Map
                d_1802_v95_ = _dafny.Map({(d_1792_v87_)[default__.safeIndex(171, (d_1792_v87_).length(0))]: len(d_1798_v93_)})
                d_1803_v96_: _dafny.Set
                d_1803_v96_ = _dafny.Set({default__.fm15(False, globalState)})
                d_1804_v97_: _dafny.Map
                d_1804_v97_ = _dafny.Map({len(d_1802_v95_): d_1803_v96_})
                rhs307_ = not((self).f24)
                rhs308_ = (d_1798_v93_) <= (d_1798_v93_)
                rhs309_ = ((d_1804_v97_) | (d_1804_v97_)).set((d_1799_v94_)[default__.safeIndex(528, (d_1799_v94_).length(0))], d_1803_v96_)
                rhs310_ = (d_1798_v93_) + (_dafny.SeqWithoutIsStrInference([d_1791_v86_ for d_1805_i8_ in range(default__.abs(-748))]))
                lhs252_ = globalState
                r0 = rhs307_
                lhs252_.f10 = rhs308_
                d_1804_v97_ = rhs309_
                d_1798_v93_ = rhs310_
            (globalState).f0 = (self).f24
            d_1806_v98_: _dafny.Array
            def lambda90_(d_1807_i9_):
                return _dafny.Set({not((self).f24), not(not((self).f24))})

            init48_ = lambda90_
            nw271_ = _dafny.Array(None, 26)
            for i0_48_ in range(nw271_.length(0)):
                nw271_[i0_48_] = init48_(i0_48_)
            d_1806_v98_ = nw271_
            d_1808_v99_: _dafny.Set
            d_1808_v99_ = _dafny.Set({(self).f24})
            d_1809_v100_: _dafny.Map
            d_1809_v100_ = _dafny.Map({(self).f24: d_1808_v99_})
            index268_ = default__.safeIndex(42, (d_1806_v98_).length(0))
            (d_1806_v98_)[index268_] = ((d_1809_v100_)[(self).f24] if ((self).f24) in (d_1809_v100_) else d_1808_v99_)
            index269_ = default__.safeIndex(42, (d_1806_v98_).length(0))
            (d_1806_v98_)[index269_] = d_1808_v99_
            d_1810_v101_: _dafny.MultiSet
            d_1810_v101_ = _dafny.MultiSet([d_1694_v0_])
            d_1811_v102_: _dafny.Set
            d_1811_v102_ = _dafny.Set({d_1810_v101_})
            rhs311_ = d_1694_v0_
            rhs312_ = ((self).f24) or ((d_1811_v102_).isdisjoint(d_1811_v102_))
            lhs253_ = globalState
            lhs254_ = globalState
            lhs253_.f6 = rhs311_
            lhs254_.f10 = rhs312_
        d_1812_v103_: _dafny.Set
        d_1812_v103_ = _dafny.Set({((self).f24) == ((self).f24), (self).f24, (self).f24})
        d_1812_v103_ = d_1812_v103_
        d_1813_v104_: D15
        d_1813_v104_ = D15_DC51((self).f24, (self).f24)
        source25_ = (d_1813_v104_ if (self).f24 else d_1813_v104_)
        if source25_.is_DC50:
            (globalState).f6 = d_1694_v0_
            d_1814_v105_: _dafny.Seq
            d_1814_v105_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            d_1815_v106_: _dafny.MultiSet
            d_1815_v106_ = _dafny.MultiSet([False, (self).f24, (self).f24, (self).f24, (self).f24])
            d_1816_v107_: str
            d_1816_v107_ = _dafny.CodePoint('x')
            d_1817_v108_: D10
            d_1817_v108_ = D10_DC34(d_1816_v107_, (self).f24, (self).f24, (self).f24, (self).f24)
            d_1818_v109_: D11
            d_1818_v109_ = D11_DC39((self).f24)
            d_1819_v110_: _dafny.Set
            d_1819_v110_ = _dafny.Set({(_dafny.MultiSet(d_1814_v105_)).cardinality})
            d_1820_v111_: _dafny.Map
            d_1820_v111_ = _dafny.Map({len(d_1819_v110_): (self).f24})
            d_1821_v112_: _dafny.Map
            d_1821_v112_ = d_1820_v111_
            d_1822_v113_: _dafny.Seq
            d_1822_v113_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsxnhwfp"))
            d_1823_v114_: _dafny.Array
            nw272_ = _dafny.Array(None, 22)
            nw272_[int(0)] = (d_1814_v105_)[default__.safeIndex(((d_1815_v106_)[(self).f24] if ((self).f24) in (d_1815_v106_) else d_1694_v0_), len(d_1814_v105_))]
            nw272_[int(1)] = (self).f24
            nw272_[int(2)] = (self).f24
            nw272_[int(3)] = (self).f24
            nw272_[int(4)] = ((self).f24 if (self).f24 else (self).f24)
            nw272_[int(5)] = ((self).f24 if (self).f24 else (self).f24)
            nw272_[int(6)] = True
            nw272_[int(7)] = (self).f24
            nw272_[int(8)] = (self).f24
            nw272_[int(9)] = (self).f24
            nw272_[int(10)] = (self).f24
            nw272_[int(11)] = (self).f24
            nw272_[int(12)] = not((d_1817_v108_).cf57)
            nw272_[int(13)] = (self).f24
            nw272_[int(14)] = (d_1694_v0_) < (d_1694_v0_)
            nw272_[int(15)] = (d_1818_v109_).cf66
            nw272_[int(16)] = not((self).f24)
            nw272_[int(17)] = (d_1694_v0_) <= (-636)
            nw272_[int(18)] = ((self).f24 if (self).f24 else (self).f24)
            nw272_[int(19)] = True
            nw272_[int(20)] = default__.fm37(d_1814_v105_, (d_1821_v112_), d_1822_v113_, globalState)
            nw272_[int(21)] = True
            d_1823_v114_ = nw272_
            index270_ = default__.safeIndex(184, (d_1823_v114_).length(0))
            (d_1823_v114_)[index270_] = not((self).f24)
            index271_ = default__.safeIndex(184, (d_1823_v114_).length(0))
            (d_1823_v114_)[index271_] = (d_1819_v110_).issubset(d_1819_v110_)
            d_1824_v115_: C3
            nw273_ = C3()
            nw273_.ctor__((self).f22)
            d_1824_v115_ = nw273_
            d_1824_v115_ = d_1824_v115_
            d_1694_v0_ = 861
        elif source25_.is_DC51:
            d_1825___mcc_h0_ = source25_.cf82
            d_1826___mcc_h1_ = source25_.cf83
            d_1827_cf83_ = d_1826___mcc_h1_
            d_1828_cf82_ = d_1825___mcc_h0_
            d_1829_v116_: _dafny.Map
            d_1829_v116_ = _dafny.Map({d_1827_cf83_: False})
            d_1830_v117_: _dafny.Array
            def lambda91_(d_1831_i10_):
                return (d_1831_i10_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvvl"))))

            init49_ = lambda91_
            nw274_ = _dafny.Array(None, 4)
            for i0_49_ in range(nw274_.length(0)):
                nw274_[i0_49_] = init49_(i0_49_)
            d_1830_v117_ = nw274_
            d_1832_v118_: _dafny.Map
            d_1832_v118_ = _dafny.Map({d_1830_v117_: d_1827_cf83_})
            d_1829_v116_ = (d_1829_v116_).set(((d_1832_v118_)[d_1830_v117_] if (d_1830_v117_) in (d_1832_v118_) else (self).f24), d_1827_cf83_)
            d_1833_v119_: T0
            nw275_ = C1()
            nw275_.ctor__(779, (self).f22)
            d_1833_v119_ = nw275_
            d_1834_v120_: _dafny.Map
            d_1834_v120_ = _dafny.Map({d_1694_v0_: d_1833_v119_})
            d_1835_v121_: _dafny.Seq
            d_1835_v121_ = _dafny.SeqWithoutIsStrInference([True, not((self).f24), False])
            d_1836_v122_: _dafny.Array
            nw276_ = _dafny.Array(None, 23)
            nw276_[int(0)] = d_1833_v119_
            nw276_[int(1)] = d_1833_v119_
            nw276_[int(2)] = d_1833_v119_
            nw276_[int(3)] = d_1833_v119_
            nw276_[int(4)] = d_1833_v119_
            nw276_[int(5)] = d_1833_v119_
            nw276_[int(6)] = d_1833_v119_
            nw276_[int(7)] = ((d_1834_v120_)[len(d_1835_v121_)] if (len(d_1835_v121_)) in (d_1834_v120_) else d_1833_v119_)
            nw276_[int(8)] = d_1833_v119_
            nw276_[int(9)] = d_1833_v119_
            nw276_[int(10)] = d_1833_v119_
            nw276_[int(11)] = d_1833_v119_
            nw276_[int(12)] = d_1833_v119_
            nw276_[int(13)] = d_1833_v119_
            nw276_[int(14)] = (d_1833_v119_ if (self).f24 else d_1833_v119_)
            nw276_[int(15)] = d_1833_v119_
            nw276_[int(16)] = d_1833_v119_
            nw276_[int(17)] = d_1833_v119_
            nw276_[int(18)] = d_1833_v119_
            nw276_[int(19)] = d_1833_v119_
            nw276_[int(20)] = d_1833_v119_
            nw276_[int(21)] = d_1833_v119_
            nw276_[int(22)] = d_1833_v119_
            d_1836_v122_ = nw276_
            index272_ = default__.safeIndex(380, (d_1836_v122_).length(0))
            (d_1836_v122_)[index272_] = d_1833_v119_
            index273_ = default__.safeIndex(380, (d_1836_v122_).length(0))
            (d_1836_v122_)[index273_] = d_1833_v119_
            d_1828_cf82_ = not (d_1827_cf83_) or (False)
            d_1837_v123_: C2
            nw277_ = C2()
            nw277_.ctor__(d_1694_v0_, (self).f22)
            d_1837_v123_ = nw277_
            d_1837_v123_ = d_1837_v123_
        elif source25_.is_DC52:
            d_1838_v124_: _dafny.Array
            nw278_ = _dafny.Array(False, 15)
            d_1838_v124_ = nw278_
            index274_ = default__.safeIndex(939, (d_1838_v124_).length(0))
            (d_1838_v124_)[index274_] = True
            index275_ = default__.safeIndex(939, (d_1838_v124_).length(0))
            (d_1838_v124_)[index275_] = (312) < (d_1694_v0_)
            d_1839_v125_: _dafny.Seq
            d_1839_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "elp"))
            d_1840_v126_: str
            d_1840_v126_ = _dafny.CodePoint('d')
            d_1841_v127_: _dafny.Map
            d_1841_v127_ = _dafny.Map({(d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]: d_1840_v126_})
            d_1842_v128_: _dafny.Map
            d_1842_v128_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([486, (0) - (len(d_1839_v125_))]): (d_1841_v127_ if False else d_1841_v127_)})
            d_1843_v129_: _dafny.MultiSet
            d_1843_v129_ = _dafny.MultiSet([d_1812_v103_])
            d_1844_v130_: _dafny.Seq
            d_1844_v130_ = _dafny.SeqWithoutIsStrInference([(d_1843_v129_).cardinality, (0) - (d_1694_v0_)])
            d_1845_v131_: _dafny.Seq
            d_1845_v131_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({not((d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]): d_1840_v126_})])
            d_1846_v132_: _dafny.MultiSet
            d_1846_v132_ = _dafny.MultiSet([d_1694_v0_, d_1694_v0_])
            d_1842_v128_ = (d_1842_v128_).set(d_1844_v130_, (d_1845_v131_)[default__.safeIndex((d_1846_v132_).cardinality, len(d_1845_v131_))])
            if (self).f24:
                d_1847_v133_: C3
                nw279_ = C3()
                nw279_.ctor__(((self).f22).set(default__.safeIndex(d_1694_v0_, len((self).f22)), d_1838_v124_))
                d_1847_v133_ = nw279_
                d_1848_v134_: _dafny.Array
                nw280_ = _dafny.Array(int(0), 3)
                d_1848_v134_ = nw280_
                index276_ = default__.safeIndex(199, (d_1848_v134_).length(0))
                (d_1848_v134_)[index276_] = (d_1694_v0_) - (d_1694_v0_)
                d_1849_v135_: _dafny.Seq
                d_1849_v135_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                d_1850_v136_: _dafny.Map
                d_1850_v136_ = _dafny.Map({d_1694_v0_: (d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]})
                index277_ = default__.safeIndex(199, (d_1848_v134_).length(0))
                rhs313_ = _dafny.SeqWithoutIsStrInference([((d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))] if False else default__.fm37(d_1849_v135_, d_1850_v136_, d_1839_v125_, globalState)), ((self).f24) or ((self).f24)])
                rhs314_ = (d_1694_v0_) - (d_1694_v0_)
                rhs315_ = d_1694_v0_
                rhs316_ = d_1838_v124_
                rhs317_ = d_1694_v0_
                lhs255_ = globalState
                lhs256_ = d_1848_v134_
                lhs257_ = default__.safeIndex(199, (d_1848_v134_).length(0))
                lhs258_ = globalState
                lhs255_.f5 = rhs313_
                lhs256_[lhs257_] = rhs314_
                lhs258_.f8 = rhs315_
                d_1838_v124_ = rhs316_
                d_1694_v0_ = rhs317_
                d_1840_v126_ = d_1840_v126_
                r1 = (d_1844_v130_)[default__.safeIndex((d_1848_v134_)[default__.safeIndex(199, (d_1848_v134_).length(0))], len(d_1844_v130_))]
                d_1840_v126_ = d_1840_v126_
            elif True:
                d_1851_v137_: _dafny.Seq
                d_1851_v137_ = _dafny.SeqWithoutIsStrInference([(d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]])
                d_1852_v138_: _dafny.Map
                d_1852_v138_ = _dafny.Map({(0) - ((-419) - ((0) - (d_1694_v0_))): (d_1851_v137_) + (d_1851_v137_)})
                d_1853_v139_: _dafny.Array
                nw281_ = _dafny.Array(int(0), 8)
                d_1853_v139_ = nw281_
                index278_ = default__.safeIndex(802, (d_1853_v139_).length(0))
                (d_1853_v139_)[index278_] = d_1694_v0_
                d_1854_v140_: _dafny.Map
                d_1854_v140_ = _dafny.Map({d_1840_v126_: d_1694_v0_})
                d_1855_v141_: _dafny.Seq
                d_1855_v141_ = _dafny.SeqWithoutIsStrInference([d_1854_v140_])
                d_1856_v142_: _dafny.Map
                d_1856_v142_ = _dafny.Map({367: (d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]})
                index279_ = default__.safeIndex(802, (d_1853_v139_).length(0))
                rhs318_ = d_1851_v137_
                rhs319_ = (d_1852_v138_) | (_dafny.Map({d_1694_v0_: d_1851_v137_}))
                rhs320_ = default__.safeModuloInt(len((d_1855_v141_)[default__.safeIndex(d_1694_v0_, len(d_1855_v141_))]), len(d_1856_v142_))
                rhs321_ = ((d_1846_v132_).cardinality) != (default__.safeDivisionInt(len(d_1839_v125_), d_1694_v0_))
                lhs259_ = globalState
                lhs260_ = d_1853_v139_
                lhs261_ = default__.safeIndex(802, (d_1853_v139_).length(0))
                lhs259_.f5 = rhs318_
                d_1852_v138_ = rhs319_
                lhs260_[lhs261_] = rhs320_
                r0 = rhs321_
                d_1857_v143_: _dafny.Seq
                d_1857_v143_ = _dafny.SeqWithoutIsStrInference([d_1839_v125_, d_1839_v125_])
                d_1858_v144_: _dafny.Map
                d_1858_v144_ = _dafny.Map({(d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]: (self).f24})
                d_1859_v145_: D9
                d_1859_v145_ = D9_DC31(len(d_1858_v144_), d_1838_v124_)
                d_1860_v146_: D9
                d_1860_v146_ = D9_DC32(d_1859_v145_)
                d_1861_v147_: _dafny.Map
                d_1861_v147_ = _dafny.Map({d_1857_v143_: D9_DC32(d_1860_v146_)})
                d_1862_v148_: D9
                d_1862_v148_ = D9_DC32(((d_1861_v147_)[d_1857_v143_] if (d_1857_v143_) in (d_1861_v147_) else d_1860_v146_))
                d_1863_v149_: D9
                d_1863_v149_ = D9_DC32(d_1859_v145_)
                d_1864_v150_: _dafny.Map
                d_1864_v150_ = _dafny.Map({d_1863_v149_: (self).f24})
                d_1865_v151_: _dafny.Seq
                d_1865_v151_ = _dafny.SeqWithoutIsStrInference([d_1864_v150_, _dafny.Map({d_1863_v149_: (self).f24}), (d_1864_v150_) | (d_1864_v150_), d_1864_v150_])
                pat_let_tv67_ = d_1860_v146_
                def iife149_(_pat_let38_0):
                    def iife150_(d_1866_dt__update__tmp_h1_):
                        def iife151_(_pat_let39_0):
                            def iife152_(d_1867_dt__update_hcf53_h0_):
                                return D9_DC32(d_1867_dt__update_hcf53_h0_)
                            return iife152_(_pat_let39_0)
                        return iife151_(pat_let_tv67_)
                    return iife150_(_pat_let38_0)
                rhs322_ = _dafny.SeqWithoutIsStrInference([(_dafny.Map({iife149_(D9_DC32(d_1860_v146_)): True})) | (d_1864_v150_)])
                rhs323_ = (self).f24
                rhs324_ = _dafny.CodePoint('m')
                lhs262_ = globalState
                d_1865_v151_ = rhs322_
                lhs262_.f10 = rhs323_
                d_1840_v126_ = rhs324_
                d_1868_v152_: _dafny.MultiSet
                d_1868_v152_ = _dafny.MultiSet([True])
                d_1840_v126_ = (d_1839_v125_)[default__.safeIndex(((d_1868_v152_)[(d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]] if ((d_1838_v124_)[default__.safeIndex(939, (d_1838_v124_).length(0))]) in (d_1868_v152_) else (d_1853_v139_)[default__.safeIndex(802, (d_1853_v139_).length(0))]), len(d_1839_v125_))]
                d_1839_v125_ = d_1839_v125_
                d_1869_v153_: _dafny.Set
                d_1869_v153_ = _dafny.Set({d_1694_v0_})
                (globalState).f13 = d_1869_v153_
            d_1870_v154_: _dafny.Array
            def lambda92_(d_1871_i11_):
                return _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])

            init50_ = lambda92_
            nw282_ = _dafny.Array(None, 16)
            for i0_50_ in range(nw282_.length(0)):
                nw282_[i0_50_] = init50_(i0_50_)
            d_1870_v154_ = nw282_
            d_1872_v155_: _dafny.Map
            d_1872_v155_ = _dafny.Map({d_1844_v130_: d_1870_v154_})
            d_1873_v156_: _dafny.Array
            nw283_ = _dafny.Array(None, 26)
            nw283_[int(0)] = d_1870_v154_
            nw283_[int(1)] = d_1870_v154_
            nw283_[int(2)] = d_1870_v154_
            nw283_[int(3)] = d_1870_v154_
            nw283_[int(4)] = d_1870_v154_
            nw283_[int(5)] = d_1870_v154_
            nw283_[int(6)] = d_1870_v154_
            nw283_[int(7)] = d_1870_v154_
            nw283_[int(8)] = d_1870_v154_
            nw283_[int(9)] = d_1870_v154_
            nw283_[int(10)] = d_1870_v154_
            nw283_[int(11)] = d_1870_v154_
            nw283_[int(12)] = d_1870_v154_
            nw283_[int(13)] = d_1870_v154_
            nw283_[int(14)] = (d_1870_v154_ if (self).f24 else ((d_1872_v155_)[d_1844_v130_] if (d_1844_v130_) in (d_1872_v155_) else d_1870_v154_))
            nw283_[int(15)] = d_1870_v154_
            nw283_[int(16)] = d_1870_v154_
            nw283_[int(17)] = d_1870_v154_
            nw283_[int(18)] = d_1870_v154_
            nw283_[int(19)] = d_1870_v154_
            nw283_[int(20)] = d_1870_v154_
            nw283_[int(21)] = d_1870_v154_
            nw283_[int(22)] = d_1870_v154_
            nw283_[int(23)] = d_1870_v154_
            nw283_[int(24)] = d_1870_v154_
            nw283_[int(25)] = d_1870_v154_
            d_1873_v156_ = nw283_
            index280_ = default__.safeIndex(738, (d_1873_v156_).length(0))
            (d_1873_v156_)[index280_] = d_1870_v154_
            index281_ = default__.safeIndex(738, (d_1873_v156_).length(0))
            (d_1873_v156_)[index281_] = d_1870_v154_
        elif source25_.is_DC49:
            d_1874___mcc_h2_ = source25_.cf81
            d_1875_cf81_ = d_1874___mcc_h2_
            d_1876_v157_: C3
            nw284_ = C3()
            nw284_.ctor__(((self).f22) + ((self).f22))
            d_1876_v157_ = nw284_
            d_1877_v158_: _dafny.Array
            nw285_ = _dafny.Array(_dafny.Map({}), 26)
            d_1877_v158_ = nw285_
            d_1877_v158_ = d_1877_v158_
            rhs325_ = (self).f24
            rhs326_ = ((910) - (218)) + (d_1694_v0_)
            lhs263_ = globalState
            lhs264_ = globalState
            lhs263_.f0 = rhs325_
            lhs264_.f16 = rhs326_
            d_1878_v159_: _dafny.MultiSet
            d_1878_v159_ = _dafny.MultiSet([d_1694_v0_, d_1694_v0_, d_1694_v0_, d_1694_v0_])
            d_1879_v160_: _dafny.Seq
            d_1879_v160_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1880_i13_ in range(default__.abs(974))])])
            (globalState).f6 = (default__.safeDivisionInt(d_1694_v0_, (d_1876_v157_).fm0(d_1878_v159_, d_1694_v0_, _dafny.SeqWithoutIsStrInference([_dafny.Set({d_1694_v0_, d_1694_v0_}) for d_1881_i12_ in range(default__.abs(861))]), (self).f24, globalState))) - (len(d_1879_v160_))
        elif True:
            d_1882___mcc_h3_ = source25_.cf84
            d_1883_cf84_ = d_1882___mcc_h3_
            d_1884_v161_: C1
            nw286_ = C1()
            nw286_.ctor__(d_1694_v0_, (self).f22)
            d_1884_v161_ = nw286_
            d_1885_v162_: _dafny.Seq
            d_1885_v162_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrcpc"))
            d_1886_v163_: D1
            d_1886_v163_ = D1_DC3(d_1885_v162_)
            source26_ = d_1886_v163_
            if source26_.is_DC3:
                d_1887___mcc_h4_ = source26_.cf6
                d_1888_cf6_ = d_1887___mcc_h4_
                d_1889_v164_: D4
                d_1889_v164_ = D4_DC15(len(d_1885_v162_), (0) - ((d_1884_v161_).fm34((self).f24, globalState)))
                d_1889_v164_ = d_1889_v164_
                d_1885_v162_ = d_1888_cf6_
                (globalState).f0 = (self).f24
                d_1890_v165_: _dafny.Map
                d_1890_v165_ = _dafny.Map({d_1888_cf6_: True})
                d_1890_v165_ = (d_1890_v165_).set(d_1888_cf6_, (self).f24)
            elif source26_.is_DC2:
                d_1891___mcc_h5_ = source26_.cf5
                d_1892_cf5_ = d_1891___mcc_h5_
                d_1892_cf5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqyii"))
                d_1893_v166_: T0
                nw287_ = C3()
                nw287_.ctor__((self).f22)
                d_1893_v166_ = nw287_
                d_1894_v167_: _dafny.Array
                nw288_ = _dafny.Array(int(0), 9)
                d_1894_v167_ = nw288_
                d_1895_v168_: _dafny.Set
                d_1895_v168_ = _dafny.Set({d_1894_v167_})
                d_1896_v169_: D13
                d_1896_v169_ = D13_DC46((self).f24, d_1893_v166_, (d_1884_v161_).f30, d_1895_v168_)
                d_1897_v170_: _dafny.Map
                d_1897_v170_ = _dafny.Map({d_1892_cf5_: (d_1884_v161_).f30})
                d_1898_v171_: _dafny.Map
                d_1898_v171_ = _dafny.Map({d_1694_v0_: (d_1884_v161_).f30})
                d_1899_v172_: _dafny.Seq
                d_1899_v172_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1898_v171_)), (d_1884_v161_).f30, (d_1884_v161_).f30])
                d_1900_v173_: _dafny.Array
                nw289_ = _dafny.Array(None, 10)
                nw289_[int(0)] = (self).f24
                nw289_[int(1)] = (d_1896_v169_).cf73
                nw289_[int(2)] = (_dafny.SeqWithoutIsStrInference([((d_1897_v170_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "os"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "os"))) in (d_1897_v170_) else (d_1884_v161_).f30)])) < (d_1899_v172_)
                nw289_[int(3)] = True
                nw289_[int(4)] = False
                nw289_[int(5)] = (self).f24
                nw289_[int(6)] = (self).f24
                nw289_[int(7)] = (self).f24
                nw289_[int(8)] = not ((self).f24) or ((self).f24)
                nw289_[int(9)] = default__.fm15((self).f24, globalState)
                d_1900_v173_ = nw289_
                d_1901_v174_: _dafny.Map
                d_1901_v174_ = _dafny.Map({(self).f24: True})
                d_1902_v175_: _dafny.Map
                d_1902_v175_ = _dafny.Map({len(d_1901_v174_): (self).f24})
                index282_ = default__.safeIndex(697, (d_1900_v173_).length(0))
                (d_1900_v173_)[index282_] = not (default__.fm37(_dafny.SeqWithoutIsStrInference([(self).f24, False, (self).f24]), d_1902_v175_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a")), globalState)) or ((self).f24)
                index283_ = default__.safeIndex(697, (d_1900_v173_).length(0))
                (d_1900_v173_)[index283_] = (self).f24
                d_1903_v176_: _dafny.Map
                d_1903_v176_ = _dafny.Map({d_1902_v175_: ((d_1900_v173_)[default__.safeIndex(697, (d_1900_v173_).length(0))] if (self).f24 else default__.fm15(False, globalState))})
                d_1903_v176_ = (d_1903_v176_).set(d_1902_v175_, (d_1900_v173_)[default__.safeIndex(697, (d_1900_v173_).length(0))])
                index284_ = default__.safeIndex(689, (d_1894_v167_).length(0))
                (d_1894_v167_)[index284_] = (d_1884_v161_).f30
                d_1904_v177_: _dafny.MultiSet
                d_1904_v177_ = _dafny.MultiSet([d_1694_v0_])
                d_1905_v178_: _dafny.Seq
                d_1905_v178_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_1884_v161_).f30 for d_1906_i14_ in range(default__.abs(-361))])])
                d_1907_v179_: _dafny.Set
                d_1907_v179_ = _dafny.Set({(d_1884_v161_).f30})
                d_1908_v180_: _dafny.Seq
                d_1908_v180_ = _dafny.SeqWithoutIsStrInference([d_1907_v179_])
                d_1909_v181_: str
                d_1909_v181_ = _dafny.CodePoint('k')
                index285_ = default__.safeIndex(689, (d_1894_v167_).length(0))
                index286_ = default__.safeIndex(697, (d_1900_v173_).length(0))
                rhs327_ = ((self).f24) or ((d_1900_v173_)[default__.safeIndex(697, (d_1900_v173_).length(0))])
                rhs328_ = (0) - ((d_1893_v166_).fm0(d_1904_v177_, len(d_1905_v178_), d_1908_v180_, (self).f24, globalState))
                rhs329_ = d_1694_v0_
                rhs330_ = d_1694_v0_
                rhs331_ = (d_1909_v181_) in (d_1892_cf5_)
                lhs265_ = globalState
                lhs266_ = d_1894_v167_
                lhs267_ = default__.safeIndex(689, (d_1894_v167_).length(0))
                lhs268_ = globalState
                lhs269_ = globalState
                lhs270_ = d_1900_v173_
                lhs271_ = default__.safeIndex(697, (d_1900_v173_).length(0))
                lhs265_.f0 = rhs327_
                lhs266_[lhs267_] = rhs328_
                lhs268_.f11 = rhs329_
                lhs269_.f6 = rhs330_
                lhs270_[lhs271_] = rhs331_
            elif True:
                d_1910___mcc_h6_ = source26_.cf7
                d_1911_cf7_ = d_1910___mcc_h6_
                (globalState).f8 = d_1694_v0_
                d_1912_v182_: D1
                d_1912_v182_ = D1_DC3(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cokvtw")))
                d_1913_v183_: D1
                d_1913_v183_ = D1_DC4(d_1912_v182_)
                d_1914_v184_: D1
                d_1914_v184_ = D1_DC4(d_1912_v182_)
                d_1915_v185_: D1
                d_1915_v185_ = D1_DC4(d_1913_v183_)
                d_1916_v186_: _dafny.Array
                nw290_ = _dafny.Array(None, 1)
                nw290_[int(0)] = (d_1885_v162_) + (d_1885_v162_)
                d_1916_v186_ = nw290_
                d_1917_v187_: str
                d_1917_v187_ = _dafny.CodePoint('y')
                index287_ = default__.safeIndex(293, (d_1916_v186_).length(0))
                (d_1916_v186_)[index287_] = ((d_1885_v162_).set(default__.safeIndex((d_1884_v161_).f30, len(d_1885_v162_)), d_1917_v187_)).set(default__.safeIndex(d_1694_v0_, len((d_1885_v162_).set(default__.safeIndex((d_1884_v161_).f30, len(d_1885_v162_)), d_1917_v187_))), d_1917_v187_)
                d_1918_v188_: _dafny.Array
                def lambda93_(d_1919_v161_):
                    def lambda94_(d_1920_i15_):
                        return (d_1920_i15_) - ((d_1919_v161_).f30)

                    return lambda94_

                init51_ = lambda93_(d_1884_v161_)
                nw291_ = _dafny.Array(None, 16)
                for i0_51_ in range(nw291_.length(0)):
                    nw291_[i0_51_] = init51_(i0_51_)
                d_1918_v188_ = nw291_
                index288_ = default__.safeIndex(345, (d_1918_v188_).length(0))
                (d_1918_v188_)[index288_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krpsp")))
                d_1921_v189_: _dafny.MultiSet
                d_1921_v189_ = _dafny.MultiSet([len(d_1885_v162_), 999, d_1694_v0_])
                d_1922_v190_: _dafny.Map
                d_1922_v190_ = _dafny.Map({(self).f24: len(_dafny.SeqWithoutIsStrInference([d_1917_v187_ for d_1923_i16_ in range(default__.abs(391))]))})
                d_1924_v191_: _dafny.Map
                d_1924_v191_ = _dafny.Map({((d_1921_v189_)[232] if (232) in (d_1921_v189_) else (d_1884_v161_).f30): d_1922_v190_})
                index289_ = default__.safeIndex(293, (d_1916_v186_).length(0))
                index290_ = default__.safeIndex(345, (d_1918_v188_).length(0))
                def iife153_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(249, 61):
                        d_1925_v192_: int = compr_73_
                        if ((249) <= (d_1925_v192_)) and ((d_1925_v192_) < (61)):
                            coll73_[default__.safeDivisionInt(d_1925_v192_, (d_1884_v161_).f30)] = ((d_1922_v190_).set((self).f24, d_1694_v0_)) | (d_1922_v190_)
                    return _dafny.Map(coll73_)
                rhs332_ = default__.safeDivisionInt((d_1884_v161_).f30, (d_1884_v161_).f30)
                rhs333_ = D1_DC4(d_1913_v183_)
                rhs334_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbjw"))
                rhs335_ = (d_1884_v161_).f30
                rhs336_ = iife153_()

                lhs272_ = globalState
                lhs273_ = d_1916_v186_
                lhs274_ = default__.safeIndex(293, (d_1916_v186_).length(0))
                lhs275_ = d_1918_v188_
                lhs276_ = default__.safeIndex(345, (d_1918_v188_).length(0))
                lhs272_.f16 = rhs332_
                d_1915_v185_ = rhs333_
                lhs273_[lhs274_] = rhs334_
                lhs275_[lhs276_] = rhs335_
                d_1924_v191_ = rhs336_
                d_1926_v193_: D11
                d_1926_v193_ = D11_DC36((self).f25)
                d_1927_v194_: D11
                d_1927_v194_ = D11_DC40(d_1926_v193_)
                d_1928_v195_: _dafny.Seq
                d_1928_v195_ = _dafny.SeqWithoutIsStrInference([d_1927_v194_])
                d_1929_v196_: _dafny.Map
                d_1929_v196_ = _dafny.Map({(d_1928_v195_) + (d_1928_v195_): d_1694_v0_})
                d_1929_v196_ = (d_1929_v196_).set((d_1928_v195_) + (_dafny.SeqWithoutIsStrInference([d_1927_v194_])), (d_1884_v161_).f30)
                d_1930_v197_: int
                d_1931_v198_: _dafny.Set
                d_1932_v199_: int
                d_1933_v200_: int
                out64_: int
                out65_: _dafny.Set
                out66_: int
                out67_: int
                out64_, out65_, out66_, out67_ = (d_1884_v161_).m6(104, (self).f24, ((d_1918_v188_)[default__.safeIndex(345, (d_1918_v188_).length(0))]) * ((d_1884_v161_).f30), (d_1884_v161_).f30, globalState)
                d_1930_v197_ = out64_
                d_1931_v198_ = out65_
                d_1932_v199_ = out66_
                d_1933_v200_ = out67_
            d_1934_v201_: _dafny.Array
            nw292_ = _dafny.Array(None, 2)
            d_1934_v201_ = nw292_
            d_1935_v202_: D11
            d_1935_v202_ = D11_DC37(d_1934_v201_)
            d_1936_v203_: D11
            d_1936_v203_ = D11_DC40(d_1935_v202_)
            pat_let_tv68_ = d_1935_v202_
            d_1937_v204_: _dafny.Seq
            def iife154_(_pat_let40_0):
                def iife155_(d_1938_dt__update__tmp_h2_):
                    def iife156_(_pat_let41_0):
                        def iife157_(d_1939_dt__update_hcf67_h0_):
                            return D11_DC40(d_1939_dt__update_hcf67_h0_)
                        return iife157_(_pat_let41_0)
                    return iife156_(pat_let_tv68_)
                return iife155_(_pat_let40_0)
            def iife158_(_pat_let42_0):
                def iife159_(d_1940_dt__update__tmp_h3_):
                    def iife160_(_pat_let43_0):
                        def iife161_(d_1941_dt__update_hcf67_h1_):
                            return D11_DC40(d_1941_dt__update_hcf67_h1_)
                        return iife161_(_pat_let43_0)
                    return iife160_(D11_DC39(not((self).f24)))
                return iife159_(_pat_let42_0)
            d_1937_v204_ = _dafny.SeqWithoutIsStrInference([d_1936_v203_, iife154_(d_1936_v203_), d_1936_v203_, d_1936_v203_, iife158_(D11_DC40(d_1935_v202_))])
            d_1937_v204_ = (d_1937_v204_) + (d_1937_v204_)
            d_1942_v205_: _dafny.Map
            d_1942_v205_ = _dafny.Map({D15_DC50(): (self).f24})
            d_1943_v206_: D15
            d_1943_v206_ = D15_DC50()
            d_1942_v205_ = (d_1942_v205_).set(d_1943_v206_, (self).f24)
        r0 = (self).f24
        r1 = d_1694_v0_
        return r0, r1

    @property
    def f24(self):
        return self._f24
    @property
    def f25(self):
        return self._f25

class C7(T0):
    def  __init__(self):
        self._f22: _dafny.Seq = _dafny.Seq({})
        self._f23: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f22(self):
        return self._f22
    def ctor__(self, f23, f22):
        (self)._f23 = f23
        (self)._f22 = f22

    def fm0(self, p0, p1, p2, p3, globalState):
        return (0) - (((len((self).f23)) * (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([44])})))) * (len((_dafny.SeqWithoutIsStrInference([len((self).f23), len(_dafny.Map({134: False})), (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1944_i0_ in range(default__.abs(-324))])), len(_dafny.Map({False: not(True)}))])))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([176 for d_1945_i1_ in range(default__.abs(313))]))])))))

    def fm1(self, p0, p1, globalState):
        return _dafny.CodePoint('q')

    def m0(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: int = int(0)
        d_1946_v0_: _dafny.Seq
        d_1946_v0_ = _dafny.SeqWithoutIsStrInference([True, p0])
        d_1947_v1_: int
        d_1947_v1_ = 840
        d_1948_v2_: _dafny.Map
        d_1948_v2_ = _dafny.Map({d_1946_v0_: _dafny.Set({d_1947_v1_})})
        if (((d_1948_v2_)[default__.fm2(globalState)] if (default__.fm2(globalState)) in (d_1948_v2_) else _dafny.Set({d_1947_v1_, d_1947_v1_}))) != ((self).f23):
            d_1949_v3_: _dafny.MultiSet
            d_1949_v3_ = _dafny.MultiSet([d_1947_v1_])
            d_1950_v4_: _dafny.Seq
            d_1950_v4_ = _dafny.SeqWithoutIsStrInference([d_1949_v3_, d_1949_v3_, _dafny.MultiSet([d_1947_v1_, 285]), (d_1949_v3_).set(d_1947_v1_, default__.abs(d_1947_v1_))])
            d_1951_v5_: C6
            nw293_ = C6()
            nw293_.ctor__(p0, d_1950_v4_, (self).f22)
            d_1951_v5_ = nw293_
            d_1952_v6_: _dafny.MultiSet
            d_1952_v6_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1953_i0_ in range(default__.abs(-446))])])
            d_1954_v7_: _dafny.Seq
            d_1954_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wp"))
            d_1955_v8_: _dafny.Array
            def lambda95_(d_1956_p0_):
                def lambda96_(d_1957_i1_):
                    return d_1956_p0_

                return lambda96_

            init52_ = lambda95_(p0)
            nw294_ = _dafny.Array(None, 1)
            for i0_52_ in range(nw294_.length(0)):
                nw294_[i0_52_] = init52_(i0_52_)
            d_1955_v8_ = nw294_
            d_1958_v9_: _dafny.Map
            d_1958_v9_ = _dafny.Map({d_1947_v1_: d_1955_v8_})
            d_1959_v10_: _dafny.Array
            nw295_ = _dafny.Array(None, 13)
            nw295_[int(0)] = d_1947_v1_
            nw295_[int(1)] = d_1947_v1_
            nw295_[int(2)] = d_1947_v1_
            nw295_[int(3)] = ((d_1952_v6_)[d_1954_v7_] if (d_1954_v7_) in (d_1952_v6_) else d_1947_v1_)
            nw295_[int(4)] = d_1947_v1_
            nw295_[int(5)] = d_1947_v1_
            nw295_[int(6)] = default__.fm7(d_1947_v1_, globalState)
            nw295_[int(7)] = len((d_1958_v9_).set(d_1947_v1_, d_1955_v8_))
            nw295_[int(8)] = len(d_1954_v7_)
            nw295_[int(9)] = 484
            nw295_[int(10)] = d_1947_v1_
            nw295_[int(11)] = d_1947_v1_
            nw295_[int(12)] = 571
            d_1959_v10_ = nw295_
            d_1960_v11_: _dafny.Set
            d_1960_v11_ = _dafny.Set({d_1959_v10_, d_1959_v10_, d_1959_v10_})
            d_1960_v11_ = ((d_1960_v11_).intersection(_dafny.Set({d_1959_v10_, d_1959_v10_, d_1959_v10_}))).intersection(d_1960_v11_)
            d_1961_v12_: _dafny.Seq
            d_1961_v12_ = _dafny.SeqWithoutIsStrInference([d_1947_v1_])
            d_1962_v13_: _dafny.Map
            d_1962_v13_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1951_v5_).f24, default__.fm15(p0, globalState)]): d_1961_v12_})
            d_1963_v14_: D10
            d_1963_v14_ = D10_DC34((self).fm1(p0, d_1962_v13_, globalState), (d_1951_v5_).f24, False, not(p0), (d_1951_v5_).f24)
            d_1964_v15_: _dafny.Map
            d_1964_v15_ = _dafny.Map({d_1947_v1_: (0) - (d_1947_v1_)})
            d_1965_v16_: _dafny.Map
            d_1965_v16_ = _dafny.Map({d_1963_v14_: ((d_1964_v15_)[d_1947_v1_] if (d_1947_v1_) in (d_1964_v15_) else 490)})
            d_1965_v16_ = (d_1965_v16_).set(d_1963_v14_, d_1947_v1_)
            d_1966_v17_: _dafny.Seq
            d_1966_v17_ = _dafny.SeqWithoutIsStrInference([d_1954_v7_])
            d_1967_v18_: _dafny.Map
            d_1967_v18_ = _dafny.Map({d_1947_v1_: d_1966_v17_})
            d_1968_v19_: D12
            d_1968_v19_ = D12_DC42(len(d_1961_v12_))
            d_1969_v20_: _dafny.Seq
            d_1969_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1954_v7_ for d_1970_i2_ in range(default__.abs(-56))]), d_1966_v17_, ((d_1967_v18_)[(d_1968_v19_).cf69] if ((d_1968_v19_).cf69) in (d_1967_v18_) else d_1966_v17_), _dafny.SeqWithoutIsStrInference([d_1954_v7_ for d_1971_i3_ in range(default__.abs(42))]), d_1966_v17_])
            d_1972_v21_: _dafny.Array
            nw296_ = _dafny.Array(None, 8)
            nw296_[int(0)] = d_1966_v17_
            nw296_[int(1)] = d_1966_v17_
            nw296_[int(2)] = ((d_1967_v18_)[d_1947_v1_] if (d_1947_v1_) in (d_1967_v18_) else d_1966_v17_)
            nw296_[int(3)] = ((d_1969_v20_)[default__.safeIndex(len((self).f23), len(d_1969_v20_))]) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1973_i4_ in range(default__.abs(432))]), d_1954_v7_]))
            nw296_[int(4)] = d_1966_v17_
            nw296_[int(5)] = (d_1966_v17_) + (d_1966_v17_)
            nw296_[int(6)] = (d_1966_v17_) + (d_1966_v17_)
            nw296_[int(7)] = d_1966_v17_
            d_1972_v21_ = nw296_
            index291_ = default__.safeIndex(216, (d_1972_v21_).length(0))
            (d_1972_v21_)[index291_] = d_1966_v17_
            index292_ = default__.safeIndex(216, (d_1972_v21_).length(0))
            rhs337_ = (d_1961_v12_ if p0 else d_1961_v12_)
            rhs338_ = default__.fm53(globalState)
            lhs277_ = d_1972_v21_
            lhs278_ = default__.safeIndex(216, (d_1972_v21_).length(0))
            d_1961_v12_ = rhs337_
            lhs277_[lhs278_] = rhs338_
            (globalState).f16 = (0) - (((d_1964_v15_)[d_1947_v1_] if (d_1947_v1_) in (d_1964_v15_) else (0) - (default__.fm7((_dafny.MultiSet([(d_1951_v5_).f24])).cardinality, globalState))))
        elif True:
            (globalState).f10 = not(p0)
            d_1974_v23_: _dafny.Seq
            def iife162_():
                coll74_ = _dafny.Set()
                compr_74_: int
                for compr_74_ in _dafny.IntegerRange(-127, 193):
                    d_1975_v22_: int = compr_74_
                    if ((-127) <= (d_1975_v22_)) and ((d_1975_v22_) < (193)):
                        coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_1975_v22_, (0) - (d_1947_v1_))]))
                return _dafny.Set(coll74_)
            d_1974_v23_ = _dafny.SeqWithoutIsStrInference([len(iife162_()
            )])
            d_1976_v24_: _dafny.Seq
            d_1976_v24_ = _dafny.SeqWithoutIsStrInference([d_1974_v23_])
            (globalState).f16 = (d_1947_v1_) - (len(d_1976_v24_))
            (globalState).f16 = default__.safeModuloInt(d_1947_v1_, 471)
            (globalState).f11 = d_1947_v1_
            d_1977_v25_: _dafny.Array
            def lambda97_(d_1978_i5_):
                return True

            init53_ = lambda97_
            nw297_ = _dafny.Array(None, 16)
            for i0_53_ in range(nw297_.length(0)):
                nw297_[i0_53_] = init53_(i0_53_)
            d_1977_v25_ = nw297_
            index293_ = default__.safeIndex(894, (d_1977_v25_).length(0))
            (d_1977_v25_)[index293_] = True
            index294_ = default__.safeIndex(894, (d_1977_v25_).length(0))
            (d_1977_v25_)[index294_] = (d_1946_v0_)[default__.safeIndex(d_1947_v1_, len(d_1946_v0_))]
        d_1979_v26_: D17
        d_1979_v26_ = D17_DC55((self).f23)
        rhs339_ = len(((d_1979_v26_).cf86) | (((self).f23).intersection((self).f23)))
        rhs340_ = (self).f23
        rhs341_ = d_1947_v1_
        lhs279_ = globalState
        lhs280_ = globalState
        d_1947_v1_ = rhs339_
        lhs279_.f13 = rhs340_
        lhs280_.f8 = rhs341_
        d_1980_v27_: _dafny.Seq
        d_1980_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfcba"))
        d_1981_v28_: _dafny.Seq
        d_1981_v28_ = _dafny.SeqWithoutIsStrInference([len(d_1980_v27_)])
        d_1982_v29_: D2
        d_1982_v29_ = D2_DC7(d_1981_v28_, p0, p0, (0) - (d_1947_v1_))
        hi4_ = (d_1947_v1_) + ((d_1982_v29_).cf13)
        for d_1983_i6_ in range(d_1947_v1_, hi4_):
            if p0:
                d_1984_v30_: _dafny.Array
                nw298_ = _dafny.Array(False, 20)
                d_1984_v30_ = nw298_
                d_1985_v31_: _dafny.MultiSet
                d_1985_v31_ = _dafny.MultiSet([p0, p0])
                index295_ = default__.safeIndex(976, (d_1984_v30_).length(0))
                (d_1984_v30_)[index295_] = (d_1985_v31_).ispropersubset(d_1985_v31_)
                d_1986_v33_: str
                d_1986_v33_ = _dafny.CodePoint('c')
                d_1987_v34_: _dafny.Map
                d_1987_v34_ = _dafny.Map({d_1986_v33_: d_1983_i6_})
                d_1988_v35_: _dafny.MultiSet
                d_1988_v35_ = _dafny.MultiSet([len(d_1987_v34_), d_1983_i6_, (0) - (d_1983_i6_), d_1947_v1_])
                index296_ = default__.safeIndex(976, (d_1984_v30_).length(0))
                def iife163_():
                    coll75_ = _dafny.Set()
                    compr_75_: int
                    for compr_75_ in ((d_1981_v28_).set(default__.safeIndex(d_1947_v1_, len(d_1981_v28_)), d_1947_v1_)).Elements:
                        d_1989_v32_: int = compr_75_
                        if (d_1989_v32_) in ((d_1981_v28_).set(default__.safeIndex(d_1947_v1_, len(d_1981_v28_)), d_1947_v1_)):
                            coll75_ = coll75_.union(_dafny.Set([default__.safeModuloInt(d_1989_v32_, (0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkgbxwxpm")))}))))]))
                    return _dafny.Set(coll75_)
                rhs342_ = (False) == (p0)
                rhs343_ = ((self).f23).issubset(iife163_()
                )
                rhs344_ = (601 if (d_1988_v35_).issubset(_dafny.MultiSet([d_1983_i6_])) else d_1947_v1_)
                rhs345_ = (d_1981_v28_) + ((d_1981_v28_) + (d_1981_v28_))
                lhs281_ = globalState
                lhs282_ = d_1984_v30_
                lhs283_ = default__.safeIndex(976, (d_1984_v30_).length(0))
                lhs284_ = globalState
                lhs281_.f0 = rhs342_
                lhs282_[lhs283_] = rhs343_
                lhs284_.f6 = rhs344_
                d_1981_v28_ = rhs345_
                d_1990_v36_: _dafny.Map
                d_1990_v36_ = _dafny.Map({p0: len(d_1946_v0_)})
                d_1991_v37_: C1
                nw299_ = C1()
                nw299_.ctor__((len(d_1990_v36_)) * (d_1947_v1_), (self).f22)
                d_1991_v37_ = nw299_
                d_1991_v37_ = d_1991_v37_
                d_1992_v38_: _dafny.Map
                d_1992_v38_ = _dafny.Map({(d_1984_v30_)[default__.safeIndex(976, (d_1984_v30_).length(0))]: (d_1946_v0_)[default__.safeIndex((d_1991_v37_).f30, len(d_1946_v0_))]})
                d_1992_v38_ = d_1992_v38_
                d_1993_v39_: _dafny.MultiSet
                d_1993_v39_ = _dafny.MultiSet([d_1984_v30_])
                d_1994_v40_: _dafny.Array
                def lambda98_(d_1995_v33_):
                    def lambda99_(d_1996_i7_):
                        return (_dafny.MultiSet([d_1995_v33_])).intersection(_dafny.MultiSet([d_1995_v33_]))

                    return lambda99_

                init54_ = lambda98_(d_1986_v33_)
                nw300_ = _dafny.Array(None, 24)
                for i0_54_ in range(nw300_.length(0)):
                    nw300_[i0_54_] = init54_(i0_54_)
                d_1994_v40_ = nw300_
                index297_ = default__.safeIndex(6, (d_1994_v40_).length(0))
                (d_1994_v40_)[index297_] = _dafny.MultiSet([d_1986_v33_, d_1986_v33_])
                d_1997_v41_: _dafny.Array
                nw301_ = _dafny.Array(int(0), 23)
                d_1997_v41_ = nw301_
                index298_ = default__.safeIndex(720, (d_1997_v41_).length(0))
                (d_1997_v41_)[index298_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))
                d_1998_v42_: T0
                nw302_ = C2()
                nw302_.ctor__((d_1988_v35_).cardinality, (self).f22)
                d_1998_v42_ = nw302_
                d_1999_v43_: _dafny.Seq
                d_1999_v43_ = _dafny.SeqWithoutIsStrInference([d_1981_v28_, _dafny.SeqWithoutIsStrInference([(d_1991_v37_).f30 for d_2000_i8_ in range(default__.abs(-755))])])
                d_2001_v44_: _dafny.Set
                d_2001_v44_ = _dafny.Set({(d_1984_v30_)[default__.safeIndex(976, (d_1984_v30_).length(0))]})
                d_2002_v45_: _dafny.Map
                d_2002_v45_ = _dafny.Map({d_1998_v42_: default__.safeModuloInt(d_1947_v1_, len(((d_1999_v43_)[default__.safeIndex(len(d_2001_v44_), len(d_1999_v43_))]).set(default__.safeIndex(d_1983_i6_, len((d_1999_v43_)[default__.safeIndex(len(d_2001_v44_), len(d_1999_v43_))])), (d_1991_v37_).f30)))})
                d_2003_v46_: _dafny.Map
                d_2003_v46_ = _dafny.Map({len(default__.fm23(len(d_1981_v28_), d_1947_v1_, p0, (d_1984_v30_)[default__.safeIndex(976, (d_1984_v30_).length(0))], globalState)): (d_1984_v30_)[default__.safeIndex(976, (d_1984_v30_).length(0))]})
                d_2004_v47_: _dafny.MultiSet
                d_2004_v47_ = _dafny.MultiSet([d_1986_v33_, _dafny.CodePoint('d'), d_1986_v33_])
                d_2005_v48_: D14
                d_2005_v48_ = D14_DC48(d_1983_i6_, _dafny.Map({d_1984_v30_: -950}), p0)
                index299_ = default__.safeIndex(6, (d_1994_v40_).length(0))
                index300_ = default__.safeIndex(720, (d_1997_v41_).length(0))
                rhs346_ = (d_1993_v39_).set(d_1984_v30_, default__.abs((0) - ((d_1991_v37_).fm34(((d_2003_v46_)[d_1947_v1_] if (d_1947_v1_) in (d_2003_v46_) else p0), globalState))))
                rhs347_ = d_2004_v47_
                rhs348_ = (0) - (d_1947_v1_)
                rhs349_ = ((d_2002_v45_ if (d_1984_v30_)[default__.safeIndex(976, (d_1984_v30_).length(0))] else d_2002_v45_)) | (d_2002_v45_)
                rhs350_ = ((d_1985_v31_)[(d_1985_v31_).ispropersubset(d_1985_v31_)] if ((d_1985_v31_).ispropersubset(d_1985_v31_)) in (d_1985_v31_) else default__.safeDivisionInt((d_2005_v48_).cf78, (0) - (d_1947_v1_)))
                lhs285_ = d_1994_v40_
                lhs286_ = default__.safeIndex(6, (d_1994_v40_).length(0))
                lhs287_ = d_1997_v41_
                lhs288_ = default__.safeIndex(720, (d_1997_v41_).length(0))
                lhs289_ = globalState
                d_1993_v39_ = rhs346_
                lhs285_[lhs286_] = rhs347_
                lhs287_[lhs288_] = rhs348_
                d_2002_v45_ = rhs349_
                lhs289_.f16 = rhs350_
                d_2006_v49_: _dafny.Array
                nw303_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_2006_v49_ = nw303_
                index301_ = default__.safeIndex(715, (d_2006_v49_).length(0))
                (d_2006_v49_)[index301_] = d_1980_v27_
                index302_ = default__.safeIndex(715, (d_2006_v49_).length(0))
                (d_2006_v49_)[index302_] = d_1980_v27_
            elif True:
                (globalState).f16 = default__.safeModuloInt(d_1983_i6_, 737)
                (globalState).f8 = (d_1983_i6_) - (d_1947_v1_)
                d_2007_v50_: C0
                nw304_ = C0()
                nw304_.ctor__()
                d_2007_v50_ = nw304_
                (globalState).f5 = (d_1946_v0_ if p0 else d_1946_v0_)
                d_2008_v51_: C4
                nw305_ = C4()
                nw305_.ctor__(p0, (self).f22)
                d_2008_v51_ = nw305_
            d_2009_v52_: _dafny.Map
            d_2009_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): _dafny.Map({d_1983_i6_: p0})})
            d_2010_v53_: _dafny.Map
            d_2010_v53_ = _dafny.Map({d_1947_v1_: False})
            d_2011_v54_: _dafny.Seq
            d_2011_v54_ = _dafny.SeqWithoutIsStrInference([((d_2009_v52_)[d_1946_v0_] if (d_1946_v0_) in (d_2009_v52_) else d_2010_v53_)])
            d_2011_v54_ = (_dafny.SeqWithoutIsStrInference([d_2010_v53_ for d_2012_i9_ in range(default__.abs(25))])) + (d_2011_v54_)
            if p0:
                (globalState).f10 = p0
                d_2013_v55_: _dafny.Map
                d_2013_v55_ = _dafny.Map({p0: default__.safeModuloInt((0) - (len(d_1980_v27_)), d_1983_i6_)})
                d_2013_v55_ = (d_2013_v55_).set(p0, len(d_1980_v27_))
                (globalState).f11 = (d_1981_v28_)[default__.safeIndex(d_1947_v1_, len(d_1981_v28_))]
                d_2014_v56_: _dafny.Array
                nw306_ = _dafny.Array(None, 3)
                nw306_[int(0)] = p0
                nw306_[int(1)] = (p0) and (p0)
                nw306_[int(2)] = p0
                d_2014_v56_ = nw306_
                index303_ = default__.safeIndex(862, (d_2014_v56_).length(0))
                (d_2014_v56_)[index303_] = True
                d_2015_v57_: T0
                nw307_ = C3()
                nw307_.ctor__((self).f22)
                d_2015_v57_ = nw307_
                d_2016_v58_: _dafny.Map
                d_2016_v58_ = _dafny.Map({(d_1947_v1_) == (419): p0})
                index304_ = default__.safeIndex(862, (d_2014_v56_).length(0))
                rhs351_ = (default__.fm7(d_1947_v1_, globalState)) < ((_dafny.MultiSet([d_2014_v56_])).cardinality)
                rhs352_ = ((d_2016_v58_)[p0] if (p0) in (d_2016_v58_) else p0)
                rhs353_ = d_2015_v57_
                rhs354_ = not (p0) or (p0)
                rhs355_ = (d_1946_v0_) + (d_1946_v0_)
                lhs290_ = globalState
                lhs291_ = d_2014_v56_
                lhs292_ = default__.safeIndex(862, (d_2014_v56_).length(0))
                lhs293_ = globalState
                lhs290_.f0 = rhs351_
                lhs291_[lhs292_] = rhs352_
                d_2015_v57_ = rhs353_
                r0 = rhs354_
                lhs293_.f5 = rhs355_
                d_2017_v59_: C4
                nw308_ = C4()
                nw308_.ctor__(p0, (self).f22)
                d_2017_v59_ = nw308_
            elif True:
                d_2018_v60_: _dafny.Map
                d_2018_v60_ = _dafny.Map({p0: (0) - ((0) - (d_1983_i6_))})
                d_2019_v61_: _dafny.Array
                nw309_ = _dafny.Array(False, 6)
                d_2019_v61_ = nw309_
                d_2020_v62_: C4
                nw310_ = C4()
                nw310_.ctor__((d_1981_v28_) != (_dafny.SeqWithoutIsStrInference([len(d_2018_v60_), d_1983_i6_])), (_dafny.SeqWithoutIsStrInference([d_2019_v61_])) + (((self).f22).set(default__.safeIndex(d_1947_v1_, len((self).f22)), d_2019_v61_)))
                d_2020_v62_ = nw310_
                d_2020_v62_ = d_2020_v62_
                d_2021_v63_: D10
                d_2021_v63_ = D10_DC35(d_1983_i6_, d_1983_i6_)
                def iife164_(_pat_let44_0):
                    def iife165_(d_2022_dt__update__tmp_h0_):
                        def iife166_(_pat_let45_0):
                            def iife167_(d_2023_dt__update_hcf61_h0_):
                                return D10_DC35((d_2022_dt__update__tmp_h0_).cf60, d_2023_dt__update_hcf61_h0_)
                            return iife167_(_pat_let45_0)
                        return iife166_(d_1983_i6_)
                    return iife165_(_pat_let44_0)
                (globalState).f8 = (iife164_(d_2021_v63_)).cf60
                (globalState).f10 = p0
                d_2024_v64_: _dafny.Array
                def lambda100_(d_2025_p0_):
                    def lambda101_(d_2026_i10_):
                        return _dafny.Map({d_2025_p0_: False})

                    return lambda101_

                init55_ = lambda100_(p0)
                nw311_ = _dafny.Array(None, 3)
                for i0_55_ in range(nw311_.length(0)):
                    nw311_[i0_55_] = init55_(i0_55_)
                d_2024_v64_ = nw311_
                d_2027_v65_: _dafny.Array
                nw312_ = _dafny.Array(None, 3)
                nw312_[int(0)] = d_2024_v64_
                nw312_[int(1)] = d_2024_v64_
                nw312_[int(2)] = d_2024_v64_
                d_2027_v65_ = nw312_
                index305_ = default__.safeIndex(424, (d_2027_v65_).length(0))
                (d_2027_v65_)[index305_] = (d_2024_v64_ if d_2020_v62_.f28 else d_2024_v64_)
                d_2028_v66_: _dafny.MultiSet
                d_2028_v66_ = _dafny.MultiSet([d_2020_v62_.f28, not(d_2020_v62_.f28), p0])
                index306_ = default__.safeIndex(424, (d_2027_v65_).length(0))
                rhs356_ = (802) * (default__.safeModuloInt(d_1947_v1_, (d_2028_v66_).cardinality))
                rhs357_ = (d_2028_v66_).cardinality
                rhs358_ = (d_2024_v64_ if d_2020_v62_.f28 else d_2024_v64_)
                rhs359_ = p0
                rhs360_ = p0
                lhs294_ = globalState
                lhs295_ = globalState
                lhs296_ = d_2027_v65_
                lhs297_ = default__.safeIndex(424, (d_2027_v65_).length(0))
                lhs298_ = globalState
                lhs299_ = globalState
                lhs294_.f16 = rhs356_
                lhs295_.f16 = rhs357_
                lhs296_[lhs297_] = rhs358_
                lhs298_.f10 = rhs359_
                lhs299_.f0 = rhs360_
                d_1981_v28_ = (default__.fm14(d_1983_i6_, d_1947_v1_, globalState)) + (d_1981_v28_)
            (globalState).f0 = p0
        d_2029_v67_: _dafny.MultiSet
        d_2029_v67_ = _dafny.MultiSet([p0, p0, p0])
        hi5_ = ((0) - (-607)) + (((d_2029_v67_)[p0] if (p0) in (d_2029_v67_) else d_1947_v1_))
        for d_2030_i11_ in range(621, hi5_):
            d_2031_v68_: _dafny.Map
            d_2031_v68_ = _dafny.Map({p0: len(d_1980_v27_)})
            d_2032_v69_: _dafny.Map
            d_2032_v69_ = _dafny.Map({d_2030_i11_: d_1947_v1_})
            (globalState).f8 = ((d_2031_v68_)[p0] if (p0) in (d_2031_v68_) else ((d_2032_v69_)[d_2030_i11_] if (d_2030_i11_) in (d_2032_v69_) else len(d_1981_v28_)))
            d_2033_v70_: _dafny.Array
            def lambda102_(d_2034_v1_, d_2035_p0_, d_2036_i11_):
                def lambda103_(d_2037_i12_):
                    return _dafny.SeqWithoutIsStrInference([d_2034_v1_, d_2034_v1_, len(_dafny.SeqWithoutIsStrInference([d_2035_p0_, d_2035_p0_])), len(_dafny.Set({d_2036_i11_, d_2036_i11_, d_2034_v1_}))])

                return lambda103_

            init56_ = lambda102_(d_1947_v1_, p0, d_2030_i11_)
            nw313_ = _dafny.Array(None, 10)
            for i0_56_ in range(nw313_.length(0)):
                nw313_[i0_56_] = init56_(i0_56_)
            d_2033_v70_ = nw313_
            index307_ = default__.safeIndex(231, (d_2033_v70_).length(0))
            (d_2033_v70_)[index307_] = d_1981_v28_
            index308_ = default__.safeIndex(231, (d_2033_v70_).length(0))
            (d_2033_v70_)[index308_] = ((d_1981_v28_).set(default__.safeIndex(default__.fm7(d_1947_v1_, globalState), len(d_1981_v28_)), d_2030_i11_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([((d_2032_v69_)[len(d_2032_v69_)] if (len(d_2032_v69_)) in (d_2032_v69_) else d_1947_v1_) for d_2038_i13_ in range(default__.abs(781))])), len((d_1981_v28_).set(default__.safeIndex(default__.fm7(d_1947_v1_, globalState), len(d_1981_v28_)), d_2030_i11_))), (0) - (d_1947_v1_))
            if p0:
                r3 = d_1947_v1_
                index309_ = default__.safeIndex(231, (d_2033_v70_).length(0))
                (d_2033_v70_)[index309_] = d_1981_v28_
                (globalState).f11 = d_2030_i11_
                r0 = True
                (globalState).f10 = p0
            elif True:
                d_2039_v71_: _dafny.Map
                d_2039_v71_ = _dafny.Map({d_2030_i11_: p0})
                d_2040_v72_: _dafny.Array
                nw314_ = _dafny.Array(None, 26)
                nw314_[int(0)] = p0
                nw314_[int(1)] = p0
                nw314_[int(2)] = p0
                nw314_[int(3)] = p0
                nw314_[int(4)] = p0
                nw314_[int(5)] = default__.fm37(d_1946_v0_, d_2039_v71_, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2041_i14_ in range(default__.abs(928))])).set(default__.safeIndex(d_2030_i11_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2042_i14_ in range(default__.abs(928))]))), _dafny.CodePoint('n')), globalState)
                nw314_[int(6)] = (not(p0)) == (False)
                nw314_[int(7)] = p0
                nw314_[int(8)] = (d_1947_v1_) <= (d_1947_v1_)
                nw314_[int(9)] = (d_1947_v1_) == (d_1947_v1_)
                nw314_[int(10)] = p0
                nw314_[int(11)] = p0
                nw314_[int(12)] = p0
                nw314_[int(13)] = (p0 if (d_1946_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awkmh"))), len(d_1946_v0_))] else True)
                nw314_[int(14)] = (p0) or (p0)
                nw314_[int(15)] = (d_1947_v1_) == (len(d_1946_v0_))
                nw314_[int(16)] = (p0 if p0 else p0)
                nw314_[int(17)] = (p0) and (p0)
                nw314_[int(18)] = p0
                nw314_[int(19)] = not(not(p0))
                nw314_[int(20)] = p0
                nw314_[int(21)] = (d_1947_v1_) >= (d_2030_i11_)
                nw314_[int(22)] = (d_1947_v1_) > (d_2030_i11_)
                nw314_[int(23)] = p0
                nw314_[int(24)] = p0
                nw314_[int(25)] = p0
                d_2040_v72_ = nw314_
                index310_ = default__.safeIndex(392, (d_2040_v72_).length(0))
                (d_2040_v72_)[index310_] = not (p0) or (p0)
                d_2043_v73_: _dafny.Map
                d_2043_v73_ = _dafny.Map({d_2029_v67_: d_1946_v0_})
                d_2044_v75_: _dafny.Set
                d_2044_v75_ = _dafny.Set({d_2029_v67_})
                index311_ = default__.safeIndex(392, (d_2040_v72_).length(0))
                def iife168_():
                    coll76_ = _dafny.Set()
                    compr_76_: _dafny.MultiSet
                    for compr_76_ in (d_2043_v73_).keys.Elements:
                        d_2045_v74_: _dafny.MultiSet = compr_76_
                        if (d_2045_v74_) in (d_2043_v73_):
                            coll76_ = coll76_.union(_dafny.Set([d_2045_v74_]))
                    return _dafny.Set(coll76_)
                (d_2040_v72_)[index311_] = (iife168_()
                ).isdisjoint((d_2044_v75_) | (d_2044_v75_))
                (globalState).f8 = d_2030_i11_
                d_2046_v76_: C4
                nw315_ = C4()
                nw315_.ctor__(p0, (self).f22)
                d_2046_v76_ = nw315_
                d_2047_v77_: str
                d_2047_v77_ = _dafny.CodePoint('u')
                d_2047_v77_ = d_2047_v77_
                (globalState).f16 = d_2030_i11_
            d_2048_v78_: _dafny.Array
            nw316_ = _dafny.Array(False, 13)
            d_2048_v78_ = nw316_
            index312_ = default__.safeIndex(292, (d_2048_v78_).length(0))
            (d_2048_v78_)[index312_] = p0
            index313_ = default__.safeIndex(292, (d_2048_v78_).length(0))
            (d_2048_v78_)[index313_] = p0
        d_2049_v79_: _dafny.Map
        d_2049_v79_ = _dafny.Map({p0: d_1980_v27_})
        hi6_ = (len(d_1946_v0_)) * (len(d_2049_v79_))
        for d_2050_i15_ in range(d_1947_v1_, hi6_):
            (globalState).f10 = p0
            d_2051_v80_: _dafny.Array
            nw317_ = _dafny.Array(False, 8)
            d_2051_v80_ = nw317_
            index314_ = default__.safeIndex(200, (d_2051_v80_).length(0))
            (d_2051_v80_)[index314_] = (not(p0)) or (False)
            d_2052_v81_: _dafny.MultiSet
            d_2052_v81_ = _dafny.MultiSet([(d_2050_i15_) - (d_2050_i15_)])
            d_2053_v82_: _dafny.Map
            d_2053_v82_ = _dafny.Map({p0: d_1947_v1_})
            index315_ = default__.safeIndex(200, (d_2051_v80_).length(0))
            rhs361_ = not(True)
            rhs362_ = True
            rhs363_ = (0) - ((d_2052_v81_).cardinality)
            rhs364_ = (0) - (len((_dafny.SeqWithoutIsStrInference([d_1946_v0_, d_1946_v0_, d_1946_v0_, d_1946_v0_, (d_1946_v0_).set(default__.safeIndex(len(d_2053_v82_), len(d_1946_v0_)), p0)])).set(default__.safeIndex(d_1947_v1_, len(_dafny.SeqWithoutIsStrInference([d_1946_v0_, d_1946_v0_, d_1946_v0_, d_1946_v0_, (d_1946_v0_).set(default__.safeIndex(len(d_2053_v82_), len(d_1946_v0_)), p0)]))), d_1946_v0_)))
            lhs300_ = globalState
            lhs301_ = d_2051_v80_
            lhs302_ = default__.safeIndex(200, (d_2051_v80_).length(0))
            lhs303_ = globalState
            lhs304_ = globalState
            lhs300_.f0 = rhs361_
            lhs301_[lhs302_] = rhs362_
            lhs303_.f11 = rhs363_
            lhs304_.f6 = rhs364_
            if False:
                d_2054_v83_: _dafny.Map
                d_2054_v83_ = _dafny.Map({p0: p0})
                d_2054_v83_ = (d_2054_v83_).set((d_2051_v80_)[default__.safeIndex(200, (d_2051_v80_).length(0))], (d_2051_v80_)[default__.safeIndex(200, (d_2051_v80_).length(0))])
                r1 = d_2051_v80_
                d_2055_v84_: D9
                d_2055_v84_ = D9_DC31(d_2050_i15_, d_2051_v80_)
                (globalState).f6 = (d_2055_v84_).cf51
                index316_ = default__.safeIndex(200, (d_2051_v80_).length(0))
                (d_2051_v80_)[index316_] = p0
                d_2056_v85_: _dafny.Map
                d_2056_v85_ = _dafny.Map({d_2050_i15_: False})
                d_2057_v86_: C4
                nw318_ = C4()
                nw318_.ctor__(default__.fm37(d_1946_v0_, d_2056_v85_, d_1980_v27_, globalState), (self).f22)
                d_2057_v86_ = nw318_
            elif True:
                d_2058_v87_: _dafny.Array
                nw319_ = _dafny.Array(int(0), 29)
                d_2058_v87_ = nw319_
                index317_ = default__.safeIndex(334, (d_2058_v87_).length(0))
                (d_2058_v87_)[index317_] = default__.safeModuloInt(d_1947_v1_, len(d_2053_v82_))
                index318_ = default__.safeIndex(334, (d_2058_v87_).length(0))
                (d_2058_v87_)[index318_] = d_1947_v1_
                d_2058_v87_ = d_2058_v87_
                d_2059_v88_: C1
                nw320_ = C1()
                nw320_.ctor__((d_2058_v87_)[default__.safeIndex(334, (d_2058_v87_).length(0))], (self).f22)
                d_2059_v88_ = nw320_
                d_2060_v89_: _dafny.Array
                nw321_ = _dafny.Array(None, 27)
                nw321_[int(0)] = d_2059_v88_
                nw321_[int(1)] = d_2059_v88_
                nw321_[int(2)] = d_2059_v88_
                nw321_[int(3)] = d_2059_v88_
                nw321_[int(4)] = d_2059_v88_
                nw321_[int(5)] = d_2059_v88_
                nw321_[int(6)] = d_2059_v88_
                nw321_[int(7)] = d_2059_v88_
                nw321_[int(8)] = d_2059_v88_
                nw321_[int(9)] = d_2059_v88_
                nw321_[int(10)] = d_2059_v88_
                nw321_[int(11)] = d_2059_v88_
                nw321_[int(12)] = d_2059_v88_
                nw321_[int(13)] = d_2059_v88_
                nw321_[int(14)] = d_2059_v88_
                nw321_[int(15)] = d_2059_v88_
                nw321_[int(16)] = d_2059_v88_
                nw321_[int(17)] = d_2059_v88_
                nw321_[int(18)] = d_2059_v88_
                nw321_[int(19)] = d_2059_v88_
                nw321_[int(20)] = d_2059_v88_
                nw321_[int(21)] = d_2059_v88_
                nw321_[int(22)] = d_2059_v88_
                nw321_[int(23)] = d_2059_v88_
                nw321_[int(24)] = d_2059_v88_
                nw321_[int(25)] = d_2059_v88_
                nw321_[int(26)] = d_2059_v88_
                d_2060_v89_ = nw321_
                d_2061_v90_: D11
                d_2061_v90_ = D11_DC37(d_2060_v89_)
                d_2062_v91_: D11
                d_2062_v91_ = D11_DC40(d_2061_v90_)
                d_2063_v92_: _dafny.Seq
                d_2063_v92_ = _dafny.SeqWithoutIsStrInference([d_2061_v90_])
                d_2064_v93_: _dafny.Map
                d_2064_v93_ = _dafny.Map({(d_2058_v87_)[default__.safeIndex(334, (d_2058_v87_).length(0))]: (d_2051_v80_)[default__.safeIndex(200, (d_2051_v80_).length(0))]})
                d_2065_v94_: C5
                nw322_ = C5()
                nw322_.ctor__(d_2064_v93_, d_1947_v1_)
                d_2065_v94_ = nw322_
                pat_let_tv69_ = d_2061_v90_
                pat_let_tv70_ = d_2061_v90_
                pat_let_tv71_ = d_2061_v90_
                d_2066_v95_: _dafny.Array
                nw323_ = _dafny.Array(None, 19)
                def iife169_(_pat_let46_0):
                    def iife170_(d_2067_dt__update__tmp_h1_):
                        def iife171_(_pat_let47_0):
                            def iife172_(d_2068_dt__update_hcf67_h0_):
                                return D11_DC40(d_2068_dt__update_hcf67_h0_)
                            return iife172_(_pat_let47_0)
                        return iife171_(pat_let_tv69_)
                    return iife170_(_pat_let46_0)
                nw323_[int(0)] = iife169_(d_2062_v91_)
                nw323_[int(1)] = d_2062_v91_
                nw323_[int(2)] = D11_DC40(d_2061_v90_)
                nw323_[int(3)] = d_2062_v91_
                nw323_[int(4)] = d_2062_v91_
                nw323_[int(5)] = d_2062_v91_
                nw323_[int(6)] = d_2062_v91_
                nw323_[int(7)] = d_2062_v91_
                nw323_[int(8)] = d_2062_v91_
                def iife173_(_pat_let48_0):
                    def iife174_(d_2069_dt__update__tmp_h2_):
                        def iife175_(_pat_let49_0):
                            def iife176_(d_2070_dt__update_hcf67_h1_):
                                return D11_DC40(d_2070_dt__update_hcf67_h1_)
                            return iife176_(_pat_let49_0)
                        return iife175_(pat_let_tv70_)
                    return iife174_(_pat_let48_0)
                nw323_[int(9)] = iife173_(d_2062_v91_)
                nw323_[int(10)] = d_2062_v91_
                nw323_[int(11)] = d_2062_v91_
                nw323_[int(12)] = d_2062_v91_
                def iife177_(_pat_let50_0):
                    def iife178_(d_2071_dt__update__tmp_h3_):
                        def iife179_(_pat_let51_0):
                            def iife180_(d_2072_dt__update_hcf67_h2_):
                                return D11_DC40(d_2072_dt__update_hcf67_h2_)
                            return iife180_(_pat_let51_0)
                        return iife179_(pat_let_tv71_)
                    return iife178_(_pat_let50_0)
                nw323_[int(13)] = iife177_(d_2062_v91_)
                nw323_[int(14)] = (d_2062_v91_ if (d_2051_v80_)[default__.safeIndex(200, (d_2051_v80_).length(0))] else d_2062_v91_)
                nw323_[int(15)] = D11_DC40((d_2063_v92_)[default__.safeIndex(len(_dafny.Set({d_2065_v94_})), len(d_2063_v92_))])
                nw323_[int(16)] = d_2062_v91_
                nw323_[int(17)] = D11_DC40(d_2061_v90_)
                nw323_[int(18)] = d_2062_v91_
                d_2066_v95_ = nw323_
                index319_ = default__.safeIndex(573, (d_2066_v95_).length(0))
                (d_2066_v95_)[index319_] = d_2062_v91_
                d_2073_v96_: _dafny.Set
                d_2073_v96_ = _dafny.Set({p0})
                index320_ = default__.safeIndex(573, (d_2066_v95_).length(0))
                rhs365_ = not ((d_2051_v80_)[default__.safeIndex(200, (d_2051_v80_).length(0))]) or ((d_2073_v96_).ispropersubset(d_2073_v96_))
                rhs366_ = d_2062_v91_
                rhs367_ = d_2051_v80_
                lhs305_ = globalState
                lhs306_ = d_2066_v95_
                lhs307_ = default__.safeIndex(573, (d_2066_v95_).length(0))
                lhs305_.f10 = rhs365_
                lhs306_[lhs307_] = rhs366_
                r1 = rhs367_
                index321_ = default__.safeIndex(334, (d_2058_v87_).length(0))
                (d_2058_v87_)[index321_] = 625
                (globalState).f16 = (d_2065_v94_).f27
            d_2074_v97_: str
            d_2074_v97_ = _dafny.CodePoint('j')
            d_2074_v97_ = d_2074_v97_
        d_2075_v98_: _dafny.Array
        def lambda104_(d_2076_v28_):
            def lambda105_(d_2077_i16_):
                return _dafny.SeqWithoutIsStrInference([d_2076_v28_ for d_2078_i17_ in range(default__.abs(93))])

            return lambda105_

        init57_ = lambda104_(d_1981_v28_)
        nw324_ = _dafny.Array(None, 7)
        for i0_57_ in range(nw324_.length(0)):
            nw324_[i0_57_] = init57_(i0_57_)
        d_2075_v98_ = nw324_
        d_2079_v99_: _dafny.Map
        d_2079_v99_ = _dafny.Map({p0: d_1947_v1_})
        d_2080_v100_: _dafny.Seq
        d_2080_v100_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1947_v1_, len(d_2079_v99_)])])
        index322_ = default__.safeIndex(770, (d_2075_v98_).length(0))
        (d_2075_v98_)[index322_] = (d_2080_v100_) + (d_2080_v100_)
        d_2081_v101_: D15
        d_2081_v101_ = D15_DC51(p0, False)
        d_2082_v102_: str
        d_2082_v102_ = _dafny.CodePoint('q')
        d_2083_v104_: _dafny.Map
        def iife181_():
            coll77_ = _dafny.Map()
            compr_77_: int
            for compr_77_ in (_dafny.SeqWithoutIsStrInference([-952])).Elements:
                d_2084_v103_: int = compr_77_
                if (d_2084_v103_) in (_dafny.SeqWithoutIsStrInference([-952])):
                    coll77_[(d_2084_v103_) * (d_1947_v1_)] = d_1947_v1_
            return _dafny.Map(coll77_)
        d_2083_v104_ = _dafny.Map({p0: iife181_()
        })
        d_2085_v105_: D0
        d_2085_v105_ = D0_DC1(d_1947_v1_, p0, d_2083_v104_, (d_2029_v67_).cardinality, d_2082_v102_)
        d_2086_v106_: _dafny.Seq
        d_2086_v106_ = _dafny.SeqWithoutIsStrInference([d_2085_v105_, d_2085_v105_])
        index323_ = default__.safeIndex(770, (d_2075_v98_).length(0))
        (d_2075_v98_)[index323_] = (((default__.fm54(d_2081_v101_, d_1947_v1_, d_2082_v102_, len(d_2086_v106_), globalState)) + (_dafny.SeqWithoutIsStrInference([d_1981_v28_]))).set(default__.safeIndex(64, len((default__.fm54(d_2081_v101_, d_1947_v1_, d_2082_v102_, len(d_2086_v106_), globalState)) + (_dafny.SeqWithoutIsStrInference([d_1981_v28_])))), d_1981_v28_)) + (d_2080_v100_)
        r0 = p0
        def lambda106_(d_2087_p0_):
            def lambda107_(d_2088_i18_):
                return d_2087_p0_

            return lambda107_

        init58_ = lambda106_(p0)
        nw325_ = _dafny.Array(None, 4)
        for i0_58_ in range(nw325_.length(0)):
            nw325_[i0_58_] = init58_(i0_58_)
        r1 = nw325_
        d_2089_v107_: _dafny.Map
        d_2089_v107_ = _dafny.Map({d_1981_v28_: d_2082_v102_})
        r2 = (d_2089_v107_) | (d_2089_v107_)
        r3 = d_1947_v1_
        return r0, r1, r2, r3

    @property
    def f23(self):
        return self._f23
